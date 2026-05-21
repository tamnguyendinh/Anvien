package httpapi

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/tamnguyendinh/avmatrix-go/internal/graph"
	"github.com/tamnguyendinh/avmatrix-go/internal/graphhealth"
	"github.com/tamnguyendinh/avmatrix-go/internal/repo"
	"github.com/tamnguyendinh/avmatrix-go/internal/scopeir"
)

func TestInfoPreservesWebUICompatibilityShape(t *testing.T) {
	server := httptest.NewServer(NewHandler(Config{
		Version:        "9.9.9",
		LaunchContext:  LaunchContextLocal,
		RuntimeVersion: "go-test",
	}))
	defer server.Close()

	var payload map[string]string
	getJSON(t, server.URL+"/api/info", http.StatusOK, &payload)

	for _, key := range []string{"version", "launchContext", "nodeVersion"} {
		if payload[key] == "" {
			t.Fatalf("/api/info missing %q in %#v", key, payload)
		}
	}
	if payload["version"] != "9.9.9" || payload["launchContext"] != LaunchContextLocal {
		t.Fatalf("unexpected /api/info payload: %#v", payload)
	}
}

func TestHeartbeatKeepsSSEConnectionOpen(t *testing.T) {
	server := httptest.NewServer(NewHandler(Config{}))
	defer server.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, server.URL+"/api/heartbeat", nil)
	if err != nil {
		t.Fatalf("create heartbeat request: %v", err)
	}
	response, err := server.Client().Do(request)
	if err != nil {
		t.Fatalf("open heartbeat stream: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Fatalf("heartbeat status = %d", response.StatusCode)
	}
	if contentType := response.Header.Get("Content-Type"); !strings.Contains(contentType, "text/event-stream") {
		t.Fatalf("heartbeat content type = %q", contentType)
	}

	reader := bufio.NewReader(response.Body)
	if line, err := reader.ReadString('\n'); err != nil || line != ":ok\n" {
		t.Fatalf("initial heartbeat line = %q, err=%v", line, err)
	}
	if line, err := reader.ReadString('\n'); err != nil || line != "\n" {
		t.Fatalf("initial heartbeat separator = %q, err=%v", line, err)
	}

	earlyClose := make(chan error, 1)
	go func() {
		_, err := reader.ReadByte()
		earlyClose <- err
	}()

	select {
	case err := <-earlyClose:
		t.Fatalf("heartbeat stream closed before client cancellation: %v", err)
	case <-time.After(150 * time.Millisecond):
	}

	cancel()
	response.Body.Close()
	select {
	case <-earlyClose:
	case <-time.After(2 * time.Second):
		t.Fatal("heartbeat stream did not close after client cancellation")
	}
}

func TestReposReturnsRegistrySnapshot(t *testing.T) {
	server, _ := newRepoServer(t, []repoFixture{
		{name: "alpha", indexedAt: "2026-05-08T01:02:03Z", lastCommit: "abc123"},
	})
	defer server.Close()

	var payload []map[string]any
	getJSON(t, server.URL+"/api/repos", http.StatusOK, &payload)

	if len(payload) != 1 {
		t.Fatalf("repo count = %d", len(payload))
	}
	if payload[0]["name"] != "alpha" || payload[0]["path"] == "" || payload[0]["lastCommit"] != "abc123" {
		t.Fatalf("unexpected /api/repos payload: %#v", payload[0])
	}
}

func TestRepoInfoUsesMetaWhenAvailable(t *testing.T) {
	server, fixtures := newRepoServer(t, []repoFixture{
		{name: "alpha", indexedAt: "registry-time", metaIndexedAt: "meta-time"},
	})
	defer server.Close()

	var payload map[string]any
	getJSON(t, server.URL+"/api/repo?repo="+url.QueryEscape(fixtures[0].path), http.StatusOK, &payload)

	if payload["name"] != "alpha" || payload["repoPath"] != fixtures[0].path {
		t.Fatalf("unexpected /api/repo identity: %#v", payload)
	}
	if payload["indexedAt"] != "meta-time" {
		t.Fatalf("indexedAt = %q, want meta-time", payload["indexedAt"])
	}
	stats, ok := payload["stats"].(map[string]any)
	if !ok || stats["files"] != float64(7) {
		t.Fatalf("unexpected stats payload: %#v", payload["stats"])
	}
}

func TestRepoInfoDefaultsSingleRegisteredRepo(t *testing.T) {
	server, fixtures := newRepoServer(t, []repoFixture{{name: "alpha"}})
	defer server.Close()

	var payload map[string]any
	getJSON(t, server.URL+"/api/repo", http.StatusOK, &payload)

	if payload["repoPath"] != fixtures[0].path {
		t.Fatalf("default repo path = %q, want %q", payload["repoPath"], fixtures[0].path)
	}
}

func TestRepoResolutionPrefersAbsolutePathWhenNamesDuplicate(t *testing.T) {
	server, fixtures := newRepoServer(t, []repoFixture{
		{name: "shared"},
		{name: "shared"},
	})
	defer server.Close()

	var payload map[string]any
	getJSON(t, server.URL+"/api/repo?repo="+url.QueryEscape(fixtures[1].path), http.StatusOK, &payload)

	if payload["repoPath"] != fixtures[1].path {
		t.Fatalf("path-first repo path = %q, want %q", payload["repoPath"], fixtures[1].path)
	}

	var errPayload map[string]string
	getJSON(t, server.URL+"/api/repo?repo=shared", http.StatusBadRequest, &errPayload)
	if !strings.Contains(errPayload["error"], "ambiguous") {
		t.Fatalf("duplicate name error = %#v", errPayload)
	}
}

func TestRepoInfoNotFoundKeepsErrorShape(t *testing.T) {
	server, _ := newRepoServer(t, nil)
	defer server.Close()

	var payload map[string]string
	getJSON(t, server.URL+"/api/repo?repo=missing", http.StatusNotFound, &payload)

	if payload["error"] != "Repository not found. Run: avmatrix analyze" {
		t.Fatalf("unexpected not found error: %#v", payload)
	}
}

func TestRepoDeleteRemovesIndexAndUnregisters(t *testing.T) {
	server, fixtures := newRepoServer(t, []repoFixture{{name: "alpha"}})
	defer server.Close()

	request, err := http.NewRequest(http.MethodDelete, server.URL+"/api/repo?repo=alpha", nil)
	if err != nil {
		t.Fatalf("new delete request: %v", err)
	}
	response, err := server.Client().Do(request)
	if err != nil {
		t.Fatalf("delete repo: %v", err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(response.Body)
		t.Fatalf("DELETE /api/repo status = %d; body=%s", response.StatusCode, body)
	}
	if _, err := os.Stat(repo.StoragePath(fixtures[0].path)); !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("storage path still exists or stat failed differently: %v", err)
	}

	var repos []map[string]any
	getJSON(t, server.URL+"/api/repos", http.StatusOK, &repos)
	if len(repos) != 0 {
		t.Fatalf("repos after delete = %#v", repos)
	}
}

func TestRepoDeleteRequiresRepoQuery(t *testing.T) {
	server, _ := newRepoServer(t, []repoFixture{{name: "alpha"}})
	defer server.Close()

	request, err := http.NewRequest(http.MethodDelete, server.URL+"/api/repo", nil)
	if err != nil {
		t.Fatalf("new delete request: %v", err)
	}
	response, err := server.Client().Do(request)
	if err != nil {
		t.Fatalf("delete repo: %v", err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusBadRequest {
		t.Fatalf("DELETE /api/repo status = %d, want %d", response.StatusCode, http.StatusBadRequest)
	}
}

func TestLocalFolderPickerEndpointReturnsPickedPath(t *testing.T) {
	previous := pickLocalFolderFunc
	pickLocalFolderFunc = func(context.Context) (string, error) {
		return `C:\work\repo`, nil
	}
	t.Cleanup(func() {
		pickLocalFolderFunc = previous
	})

	server := httptest.NewServer(NewHandler(Config{}))
	defer server.Close()

	var payload map[string]any
	postJSON(t, server.URL+"/api/local/folder-picker", `{}`, http.StatusOK, &payload)
	if payload["path"] != `C:\work\repo` || payload["cancelled"] != false {
		t.Fatalf("folder picker payload = %#v", payload)
	}
}

func TestLocalFolderPickerEndpointReportsUnsupported(t *testing.T) {
	previous := pickLocalFolderFunc
	pickLocalFolderFunc = func(context.Context) (string, error) {
		return "", errFolderPickerUnsupported
	}
	t.Cleanup(func() {
		pickLocalFolderFunc = previous
	})

	server := httptest.NewServer(NewHandler(Config{}))
	defer server.Close()

	var payload map[string]any
	postJSON(t, server.URL+"/api/local/folder-picker", `{}`, http.StatusNotImplemented, &payload)
	if !strings.Contains(fmt.Sprint(payload["error"]), "folder picker") {
		t.Fatalf("unsupported picker payload = %#v", payload)
	}
}

func TestLocalFolderPickerEndpointPassesRequestContext(t *testing.T) {
	previous := pickLocalFolderFunc
	var pickerContext context.Context
	pickLocalFolderFunc = func(ctx context.Context) (string, error) {
		pickerContext = ctx
		return "", ctx.Err()
	}
	t.Cleanup(func() {
		pickLocalFolderFunc = previous
	})

	handler := NewHandler(Config{})
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	request := httptest.NewRequest(http.MethodPost, "/api/local/folder-picker", nil).WithContext(ctx)
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if pickerContext == nil {
		t.Fatalf("picker did not receive request context")
	}
	if !errors.Is(pickerContext.Err(), context.Canceled) {
		t.Fatalf("picker context err = %v, want %v", pickerContext.Err(), context.Canceled)
	}
}

func TestGraphReturnsJSONForRegisteredRepo(t *testing.T) {
	server, fixtures := newRepoServer(t, []repoFixture{{name: "alpha"}})
	defer server.Close()

	var payload graphResponse
	getJSON(t, server.URL+"/api/graph?repo="+url.QueryEscape(fixtures[0].path), http.StatusOK, &payload)

	if len(payload.Nodes) != 5 || len(payload.Relationships) != 5 {
		t.Fatalf("unexpected graph payload: %#v", payload)
	}
	if payload.GraphHealth.PolicyVersion != graphhealth.PolicyVersion {
		t.Fatalf("graph JSON missing graph health summary: %#v", payload.GraphHealth)
	}
	if payload.GraphHealth.NodeCount != 5 || payload.GraphHealth.CountedRelationshipCount != 3 {
		t.Fatalf("unexpected graph health summary counts: %#v", payload.GraphHealth)
	}
	if got := payload.GraphHealth.TopologyStatusCounts[string(graphhealth.TopologyNoIncoming)]; got != 1 {
		t.Fatalf("graph health no_incoming count = %d, want 1", got)
	}
	if _, ok := payload.Nodes[0].Properties["content"]; ok {
		t.Fatalf("graph JSON should strip content by default: %#v", payload.Nodes[0].Properties)
	}
	if _, ok := payload.Nodes[0].Properties["graphHealth"].(map[string]any); !ok {
		t.Fatalf("graph JSON should include per-node graph health metadata: %#v", payload.Nodes[0].Properties)
	}
	call := payload.Relationships[2]
	if call.ResolutionSource != "scope-resolution" || call.FileHash != "hash-graph" || len(call.Evidence) != 1 {
		t.Fatalf("graph JSON lost relationship audit metadata: %#v", call)
	}
}

func TestGraphStreamingReturnsNDJSON(t *testing.T) {
	server, fixtures := newRepoServer(t, []repoFixture{{name: "alpha"}})
	defer server.Close()

	resp, err := http.Get(server.URL + "/api/graph?stream=true&repo=" + url.QueryEscape(fixtures[0].path))
	if err != nil {
		t.Fatalf("GET /api/graph stream failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("stream status = %d", resp.StatusCode)
	}
	if got := resp.Header.Get("Content-Type"); got != "application/x-ndjson; charset=utf-8" {
		t.Fatalf("stream content-type = %q", got)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("read stream body: %v", err)
	}
	text := string(body)
	if !strings.Contains(text, `"type":"node"`) || !strings.Contains(text, `"type":"relationship"`) {
		t.Fatalf("stream body = %q", string(body))
	}
	if !strings.Contains(text, `"graphHealth"`) {
		t.Fatalf("stream body missing per-node graph health metadata: %q", string(body))
	}
}

func TestGraphStreamingBatchesFlushes(t *testing.T) {
	g := graph.New()
	for i := 0; i < graphNDJSONFlushInterval+1; i++ {
		g.AddNode(graph.Node{ID: fmt.Sprintf("Function:n%d", i), Label: scopeir.NodeFunction, Properties: graph.NodeProperties{"name": "n"}})
	}
	recorder := &flushRecorder{header: http.Header{}}

	streamGraphNDJSON(recorder, g, false)

	if recorder.flushes != 2 {
		t.Fatalf("flushes = %d, want interval flush plus final flush", recorder.flushes)
	}
	if lines := strings.Count(recorder.String(), "\n"); lines != graphNDJSONFlushInterval+1 {
		t.Fatalf("stream lines = %d", lines)
	}
}

func TestGraphNodeForResponseStripsInternalDiagnostics(t *testing.T) {
	node := graph.Node{
		ID:    "Function:source",
		Label: scopeir.NodeFunction,
		Properties: graph.NodeProperties{
			"name":                            "source",
			"content":                         "export function source() {}",
			graphhealth.DiagnosticPropertyKey: []graphhealth.Diagnostic{{Kind: graphhealth.DiagnosticUnresolvedReference}},
			"diagnostics":                     []graphhealth.Diagnostic{{Kind: graphhealth.DiagnosticUnresolvedReference}},
		},
	}

	stripped := graphNodeForResponse(node, true)

	if _, ok := stripped.Properties["content"]; !ok {
		t.Fatalf("includeContent=true should keep content: %#v", stripped.Properties)
	}
	if _, ok := stripped.Properties[graphhealth.DiagnosticPropertyKey]; ok {
		t.Fatalf("response should strip internal diagnostics: %#v", stripped.Properties)
	}
	if _, ok := stripped.Properties["diagnostics"]; !ok {
		t.Fatalf("response should keep public diagnostics: %#v", stripped.Properties)
	}
}

func TestGraphHealthExplainNode(t *testing.T) {
	server, fixtures := newRepoServer(t, []repoFixture{{name: "alpha"}})
	defer server.Close()

	var payload graphHealthExplainResponse
	getJSON(t, server.URL+"/api/graph/explain?repo="+url.QueryEscape(fixtures[0].path)+"&nodeId="+url.QueryEscape("Function:main"), http.StatusOK, &payload)

	if payload.Kind != "node" || payload.NodeID != "Function:main" || payload.Node == nil || payload.Node.ID != "Function:main" {
		t.Fatalf("unexpected node explain identity: %#v", payload)
	}
	if payload.Health == nil || payload.Health.TopologyStatus != graphhealth.TopologyConnected {
		t.Fatalf("unexpected node health explain: %#v", payload.Health)
	}
	if payload.Health.CountedIncoming != 1 || payload.Health.CountedOutgoing != 1 {
		t.Fatalf("unexpected counted degree: %#v", payload.Health)
	}
	if len(payload.CountedIncomingRelationships) != 1 || payload.CountedIncomingRelationships[0].ID != "rel:helper-calls-main" {
		t.Fatalf("unexpected incoming evidence: %#v", payload.CountedIncomingRelationships)
	}
	if len(payload.CountedOutgoingRelationships) != 1 || payload.CountedOutgoingRelationships[0].ID != "rel:main-process" {
		t.Fatalf("unexpected outgoing evidence: %#v", payload.CountedOutgoingRelationships)
	}
	if len(payload.ExcludedRelationships) != 1 || payload.ExcludedRelationships[0].ID != "rel:main-member" {
		t.Fatalf("unexpected excluded evidence: %#v", payload.ExcludedRelationships)
	}
	if _, ok := payload.Node.Properties["content"]; ok {
		t.Fatalf("node explain should strip content by default: %#v", payload.Node.Properties)
	}
}

func TestGraphHealthExplainComponent(t *testing.T) {
	server, fixtures := newRepoServer(t, []repoFixture{{name: "alpha"}})
	defer server.Close()

	var nodePayload graphHealthExplainResponse
	getJSON(t, server.URL+"/api/graph/explain?repo="+url.QueryEscape(fixtures[0].path)+"&nodeId="+url.QueryEscape("Function:main"), http.StatusOK, &nodePayload)
	if nodePayload.ComponentID == "" {
		t.Fatalf("node explain missing component id: %#v", nodePayload)
	}

	var payload graphHealthExplainResponse
	getJSON(t, server.URL+"/api/graph/explain?repo="+url.QueryEscape(fixtures[0].path)+"&componentId="+url.QueryEscape(nodePayload.ComponentID), http.StatusOK, &payload)

	if payload.Kind != "component" || payload.ComponentID != nodePayload.ComponentID || payload.Component == nil {
		t.Fatalf("unexpected component explain identity: %#v", payload)
	}
	if payload.Component.NodeCount != 3 || payload.Component.CountedEdgeCount != 3 || !payload.Component.ReachableFromRoot || payload.Component.Detached {
		t.Fatalf("unexpected component summary: %#v", payload.Component)
	}
	if len(payload.Component.RootNodeIDs) != 1 || payload.Component.RootNodeIDs[0] != "Process:main" {
		t.Fatalf("unexpected component roots: %#v", payload.Component.RootNodeIDs)
	}
	if got := payload.Component.TopologyStatusCounts[string(graphhealth.TopologyConnected)]; got != 1 {
		t.Fatalf("component connected count=%d want 1", got)
	}
	if len(payload.SampleNodes) != 3 {
		t.Fatalf("component sample nodes=%d want 3: %#v", len(payload.SampleNodes), payload.SampleNodes)
	}
	if len(payload.CountedRelationshipSamples) != 3 {
		t.Fatalf("component counted relationship samples=%d want 3: %#v", len(payload.CountedRelationshipSamples), payload.CountedRelationshipSamples)
	}
	if len(payload.ExcludedRelationshipSamples) != 2 {
		t.Fatalf("component excluded relationship samples=%d want 2: %#v", len(payload.ExcludedRelationshipSamples), payload.ExcludedRelationshipSamples)
	}
}

func TestGraphHealthExplainRequiresOneTarget(t *testing.T) {
	server, fixtures := newRepoServer(t, []repoFixture{{name: "alpha"}})
	defer server.Close()

	var payload map[string]string
	getJSON(t, server.URL+"/api/graph/explain?repo="+url.QueryEscape(fixtures[0].path), http.StatusBadRequest, &payload)
	if !strings.Contains(payload["error"], "exactly one") {
		t.Fatalf("unexpected missing target error: %#v", payload)
	}
}

func TestGraphHealthReportReturnsCandidateExport(t *testing.T) {
	server, fixtures := newRepoServer(t, []repoFixture{{name: "alpha"}})
	defer server.Close()

	var payload graphHealthReportResponse
	getJSON(t, server.URL+"/api/graph/report?repo="+url.QueryEscape(fixtures[0].path)+"&limit=1", http.StatusOK, &payload)

	if payload.ReportType != "graph_health_candidate_review" || payload.VerdictPolicy != "candidate_not_confirmed" {
		t.Fatalf("unexpected report identity: %#v", payload)
	}
	if payload.Limit != 1 || payload.ReturnedCandidates != 1 || payload.TotalCandidates < 2 {
		t.Fatalf("unexpected report candidate counts: %#v", payload)
	}
	first := payload.Candidates[0]
	if first.NodeID != "Function:helper" || first.TriagePriority != string(graphhealth.TopologyNoIncoming) {
		t.Fatalf("unexpected first report candidate: %#v", first)
	}
	if first.TriageDimension != graphHealthTriageDimensionTopology {
		t.Fatalf("first triage dimension=%q want topology: %#v", first.TriageDimension, first)
	}
	if first.TopologyStatus != graphhealth.TopologyNoIncoming || first.Confidence != graphhealth.ConfidenceCandidate {
		t.Fatalf("unexpected first candidate health: %#v", first)
	}
	for _, candidate := range payload.Candidates {
		if candidate.NodeID == "Process:main" {
			t.Fatalf("expected framework node should be hidden by default: %#v", payload.Candidates)
		}
	}
	if payload.Summary.PolicyVersion != graphhealth.PolicyVersion {
		t.Fatalf("report missing graph health summary: %#v", payload.Summary)
	}
}

func TestGraphHealthReportCanIncludeExpected(t *testing.T) {
	server, fixtures := newRepoServer(t, []repoFixture{{name: "alpha"}})
	defer server.Close()

	var payload graphHealthReportResponse
	getJSON(t, server.URL+"/api/graph/report?repo="+url.QueryEscape(fixtures[0].path)+"&includeExpected=true", http.StatusOK, &payload)

	if !payload.IncludeExpected {
		t.Fatalf("includeExpected flag not reflected: %#v", payload)
	}
	foundProcess := false
	for _, candidate := range payload.Candidates {
		if candidate.NodeID == "Process:main" {
			foundProcess = true
			if candidate.Confidence != graphhealth.ConfidenceExpected ||
				candidate.TriagePriority != string(graphhealth.TopologyNoOutgoing) ||
				candidate.TriageDimension != graphHealthTriageDimensionTopology {
				t.Fatalf("unexpected expected candidate: %#v", candidate)
			}
		}
	}
	if !foundProcess {
		t.Fatalf("expected Process:main in includeExpected report: %#v", payload.Candidates)
	}
}

func TestGraphHealthReportSeparatesTopologyAndDiagnosticTriage(t *testing.T) {
	g := graph.New()
	g.AddNode(graph.Node{ID: "Function:main", Label: scopeir.NodeFunction, Properties: graph.NodeProperties{"name": "main", "filePath": "cmd/app/main.go", "isExported": true}})
	g.AddNode(graph.Node{ID: "Function:connected", Label: scopeir.NodeFunction, Properties: graph.NodeProperties{"name": "connected", "filePath": "src/app.ts"}})
	g.AddNode(graph.Node{ID: "Function:leaf", Label: scopeir.NodeFunction, Properties: graph.NodeProperties{"name": "leaf", "filePath": "src/app.ts"}})
	g.AddNode(graph.Node{ID: "Function:unwiredRoot", Label: scopeir.NodeFunction, Properties: graph.NodeProperties{"name": "unwiredRoot", "filePath": "src/unwired.ts"}})
	g.AddNode(graph.Node{ID: "Function:unwiredLeaf", Label: scopeir.NodeFunction, Properties: graph.NodeProperties{"name": "unwiredLeaf", "filePath": "src/unwired.ts"}})
	g.AddRelationship(graph.Relationship{ID: "main-connected", SourceID: "Function:main", TargetID: "Function:connected", Type: graph.RelCalls})
	g.AddRelationship(graph.Relationship{ID: "connected-leaf", SourceID: "Function:connected", TargetID: "Function:leaf", Type: graph.RelCalls})
	g.AddRelationship(graph.Relationship{ID: "unwired-leaf", SourceID: "Function:unwiredRoot", TargetID: "Function:unwiredLeaf", Type: graph.RelCalls})
	graphhealth.AppendDiagnosticToNode(g, "Function:connected", graphhealth.Diagnostic{
		Kind:       graphhealth.DiagnosticUnresolvedReference,
		FactFamily: "call",
		TargetText: "missingCall",
		FilePath:   "src/app.ts",
	})
	graphhealth.AppendDiagnosticToNode(g, "Function:unwiredRoot", graphhealth.Diagnostic{
		Kind:       graphhealth.DiagnosticUnresolvedReference,
		FactFamily: "call",
		TargetText: "missingRootCall",
		FilePath:   "src/unwired.ts",
	})
	graphhealth.ComputeSummary(g)

	candidates := graphHealthReportCandidates(g, false)
	connected := findReportCandidate(candidates, "Function:connected")
	if connected == nil {
		t.Fatalf("connected diagnostic node missing from diagnostic triage: %#v", candidates)
	}
	if connected.TopologyStatus != graphhealth.TopologyConnected ||
		connected.TriagePriority != graphhealth.DiagnosticUnresolvedReference ||
		connected.TriageDimension != graphHealthTriageDimensionDiagnostic {
		t.Fatalf("connected diagnostic candidate = %#v", connected)
	}
	unwired := findReportCandidate(candidates, "Function:unwiredRoot")
	if unwired == nil {
		t.Fatalf("unwired diagnostic node missing from topology triage: %#v", candidates)
	}
	if unwired.TopologyStatus != graphhealth.TopologyDetached ||
		unwired.TriagePriority != string(graphhealth.TopologyDetached) ||
		unwired.TriageDimension != graphHealthTriageDimensionTopology {
		t.Fatalf("unwired topology candidate = %#v", unwired)
	}
}

func TestGraphStreamingKeepsRouteAndToolMetadata(t *testing.T) {
	g := graph.New()
	g.AddNode(graph.Node{ID: "Route:/api/graph:GET", Label: scopeir.NodeRoute, Properties: graph.NodeProperties{
		"name":         "GET /api/graph",
		"filePath":     "internal/httpapi/graph.go",
		"responseKeys": []string{"nodes", "relationships"},
		"errorKeys":    []string{"error"},
		"middleware":   []string{"cors"},
	}})
	g.AddNode(graph.Node{ID: "Tool:query", Label: scopeir.NodeTool, Properties: graph.NodeProperties{
		"name":        "query",
		"filePath":    "internal/mcp/tools.go",
		"description": "Query the code graph",
	}})
	recorder := &flushRecorder{header: http.Header{}}

	streamGraphNDJSON(recorder, g, false)

	lines := strings.Split(strings.TrimSpace(recorder.String()), "\n")
	if len(lines) != 2 {
		t.Fatalf("stream lines = %d, want 2: %q", len(lines), recorder.String())
	}
	var records []map[string]any
	for _, line := range lines {
		var record map[string]any
		if err := json.Unmarshal([]byte(line), &record); err != nil {
			t.Fatalf("unmarshal stream line %q: %v", line, err)
		}
		records = append(records, record)
	}

	route := records[0]["data"].(map[string]any)
	routeProps := route["properties"].(map[string]any)
	if route["label"] != "Route" || routeProps["responseKeys"] == nil || routeProps["errorKeys"] == nil || routeProps["middleware"] == nil {
		t.Fatalf("route stream record = %#v", records[0])
	}
	tool := records[1]["data"].(map[string]any)
	toolProps := tool["properties"].(map[string]any)
	if tool["label"] != "Tool" || toolProps["description"] != "Query the code graph" {
		t.Fatalf("tool stream record = %#v", records[1])
	}
}

type flushRecorder struct {
	bytes.Buffer
	header  http.Header
	flushes int
}

func (r *flushRecorder) Header() http.Header {
	return r.header
}

func (r *flushRecorder) WriteHeader(statusCode int) {}

func (r *flushRecorder) Flush() {
	r.flushes++
}

func TestFileEndpointReadsRegisteredRepoFileRange(t *testing.T) {
	server, fixtures := newRepoServer(t, []repoFixture{{name: "alpha"}})
	defer server.Close()

	var payload fileResponse
	getJSON(t, server.URL+"/api/file?repo="+url.QueryEscape(fixtures[0].path)+"&path=src%2Fapp.ts&startLine=1&endLine=2", http.StatusOK, &payload)

	if payload.Content != "export function helper() {\n  return main()" {
		t.Fatalf("file content = %q", payload.Content)
	}
	if payload.StartLine != 1 || payload.EndLine != 2 || payload.TotalLines != 4 {
		t.Fatalf("unexpected file range metadata: %#v", payload)
	}
}

func TestFileEndpointRejectsRepositoryEscape(t *testing.T) {
	server, fixtures := newRepoServer(t, []repoFixture{{name: "alpha"}})
	defer server.Close()

	var payload map[string]string
	getJSON(t, server.URL+"/api/file?repo="+url.QueryEscape(fixtures[0].path)+"&path=..%2Fsecret.txt", http.StatusBadRequest, &payload)

	if !strings.Contains(payload["error"], "relative to the repository") {
		t.Fatalf("unexpected file escape error: %#v", payload)
	}
}

func TestQueryEndpointReturnsProcessStepsAndCallEdges(t *testing.T) {
	server, fixtures := newRepoServer(t, []repoFixture{{name: "alpha"}})
	defer server.Close()

	stepsBody := `{"repo":"` + jsonEscape(fixtures[0].path) + `","cypher":"MATCH (s)-[r:CodeRelation {type: 'STEP_IN_PROCESS'}]->(p:Process {id: 'Process:main'}) RETURN s.id AS id, s.name AS name, s.filePath AS filePath, r.step AS stepNumber ORDER BY r.step"}`
	var steps queryResponse
	postJSON(t, server.URL+"/api/query", stepsBody, http.StatusOK, &steps)

	if len(steps.Result) != 2 {
		t.Fatalf("step rows = %d, want 2: %#v", len(steps.Result), steps.Result)
	}
	if steps.Result[0]["id"] != "Function:main" || steps.Result[0]["stepNumber"] != float64(1) {
		t.Fatalf("unexpected first step row: %#v", steps.Result[0])
	}

	edgesBody := `{"repo":"` + jsonEscape(fixtures[0].path) + `","cypher":"MATCH (from)-[r:CodeRelation {type: 'CALLS'}]->(to) WHERE from.id IN ['Function:main','Function:helper'] AND to.id IN ['Function:main','Function:helper'] RETURN from.id AS fromId, to.id AS toId, r.type AS type"}`
	var edges queryResponse
	postJSON(t, server.URL+"/api/query", edgesBody, http.StatusOK, &edges)

	if len(edges.Result) != 1 {
		t.Fatalf("edge rows = %d, want 1: %#v", len(edges.Result), edges.Result)
	}
	if edges.Result[0]["fromId"] != "Function:helper" || edges.Result[0]["toId"] != "Function:main" || edges.Result[0]["type"] != string(graph.RelCalls) {
		t.Fatalf("unexpected edge row: %#v", edges.Result[0])
	}
}

func TestQueryEndpointSupportsQueryFABExamples(t *testing.T) {
	server, fixtures := newRepoServer(t, []repoFixture{{name: "alpha"}})
	defer server.Close()

	functionsBody := `{"repo":"` + jsonEscape(fixtures[0].path) + `","cypher":"MATCH (n:Function) RETURN n.id AS id, n.name AS name, n.filePath AS path LIMIT 50"}`
	var functions queryResponse
	postJSON(t, server.URL+"/api/query", functionsBody, http.StatusOK, &functions)

	if len(functions.Result) != 2 {
		t.Fatalf("function rows = %d, want 2: %#v", len(functions.Result), functions.Result)
	}
	if functions.Result[0]["id"] != "Function:main" || functions.Result[0]["path"] != "src/app.ts" {
		t.Fatalf("unexpected function row: %#v", functions.Result[0])
	}

	callsBody := `{"repo":"` + jsonEscape(fixtures[0].path) + `","cypher":"MATCH (a:File)-[r:CodeRelation {type: 'CALLS'}]->(b:Function) RETURN a.id AS id, a.name AS caller, b.name AS callee LIMIT 50"}`
	var calls queryResponse
	postJSON(t, server.URL+"/api/query", callsBody, http.StatusOK, &calls)

	if len(calls.Result) != 1 {
		t.Fatalf("call rows = %d, want 1: %#v", len(calls.Result), calls.Result)
	}
	if calls.Result[0]["caller"] != "helper" || calls.Result[0]["callee"] != "main" {
		t.Fatalf("unexpected call row: %#v", calls.Result[0])
	}
}

func TestGrepEndpointSearchesIndexedRepoFiles(t *testing.T) {
	server, fixtures := newRepoServer(t, []repoFixture{{name: "alpha"}})
	defer server.Close()

	var payload grepResponse
	getJSON(t, server.URL+"/api/grep?repo="+url.QueryEscape(fixtures[0].path)+"&pattern=return%5Cs%2Bmain&limit=5", http.StatusOK, &payload)

	if len(payload.Results) != 1 {
		t.Fatalf("grep results = %d, want 1: %#v", len(payload.Results), payload.Results)
	}
	if payload.Results[0].FilePath != "src/app.ts" || payload.Results[0].Line != 3 || payload.Results[0].Text != "return main()" {
		t.Fatalf("unexpected grep result: %#v", payload.Results[0])
	}
}

func TestProcessesAndProcessDetailEndpointsUseGraphSnapshot(t *testing.T) {
	server, fixtures := newRepoServer(t, []repoFixture{{name: "alpha"}})
	defer server.Close()

	var list processesResponse
	getJSON(t, server.URL+"/api/processes?repo="+url.QueryEscape(fixtures[0].path), http.StatusOK, &list)

	if len(list.Processes) != 1 || list.Processes[0].ID != "Process:main" || list.Processes[0].StepCount != 2 {
		t.Fatalf("unexpected process list: %#v", list)
	}

	var detail processDetailResponse
	getJSON(t, server.URL+"/api/process?repo="+url.QueryEscape(fixtures[0].path)+"&name=main", http.StatusOK, &detail)

	if detail.Process.ID != "Process:main" || len(detail.Steps) != 2 {
		t.Fatalf("unexpected process detail: %#v", detail)
	}
	if detail.Steps[0].Name != "main" || detail.Steps[0].Step != 1 {
		t.Fatalf("unexpected first process step: %#v", detail.Steps[0])
	}
}

func TestClustersAndClusterDetailEndpointsUseGraphSnapshot(t *testing.T) {
	server, fixtures := newRepoServer(t, []repoFixture{{name: "alpha"}})
	defer server.Close()

	var list clustersResponse
	getJSON(t, server.URL+"/api/clusters?repo="+url.QueryEscape(fixtures[0].path), http.StatusOK, &list)

	if len(list.Clusters) != 1 || list.Clusters[0].Label != "Api" || list.Clusters[0].SymbolCount != 5 {
		t.Fatalf("unexpected cluster list: %#v", list)
	}

	var detail clusterDetailResponse
	getJSON(t, server.URL+"/api/cluster?repo="+url.QueryEscape(fixtures[0].path)+"&name=Api", http.StatusOK, &detail)

	if detail.Cluster.ID != "comm_api" || len(detail.Members) != 2 {
		t.Fatalf("unexpected cluster detail: %#v", detail)
	}
	if detail.Members[0].Name != "helper" || detail.Members[1].Name != "main" {
		t.Fatalf("unexpected cluster members: %#v", detail.Members)
	}
}

type repoFixture struct {
	name          string
	path          string
	indexedAt     string
	metaIndexedAt string
	lastCommit    string
}

func newRepoServer(t *testing.T, fixtures []repoFixture) (*httptest.Server, []repoFixture) {
	return newRepoServerWithConfig(t, fixtures, nil)
}

func newRepoServerWithConfig(t *testing.T, fixtures []repoFixture, configure func(*Config)) (*httptest.Server, []repoFixture) {
	t.Helper()

	homeDir := filepath.Join(t.TempDir(), "home")
	store := repo.NewStore(homeDir)
	entries := make([]repo.RegistryEntry, 0, len(fixtures))
	for index := range fixtures {
		if fixtures[index].path == "" {
			fixtures[index].path = filepath.Join(t.TempDir(), fixtures[index].name)
		}
		if fixtures[index].indexedAt == "" {
			fixtures[index].indexedAt = "2026-05-08T00:00:00Z"
		}
		if fixtures[index].metaIndexedAt == "" {
			fixtures[index].metaIndexedAt = fixtures[index].indexedAt
		}
		if fixtures[index].lastCommit == "" {
			fixtures[index].lastCommit = "commit"
		}
		if err := repo.SaveMeta(repo.StoragePath(fixtures[index].path), repo.Meta{
			RepoPath:   fixtures[index].path,
			LastCommit: fixtures[index].lastCommit,
			IndexedAt:  fixtures[index].metaIndexedAt,
			Stats:      intStats(7, 11, 13),
		}); err != nil {
			t.Fatalf("save meta: %v", err)
		}
		writeGraphFixture(t, fixtures[index].path)
		entries = append(entries, repo.RegistryEntry{
			Name:        fixtures[index].name,
			Path:        fixtures[index].path,
			StoragePath: repo.StoragePath(fixtures[index].path),
			IndexedAt:   fixtures[index].indexedAt,
			LastCommit:  fixtures[index].lastCommit,
			Stats:       intStats(1, 2, 3),
		})
	}
	if err := store.WriteRegistry(entries); err != nil {
		t.Fatalf("write registry: %v", err)
	}

	config := Config{
		Store:          store,
		Version:        "test-version",
		LaunchContext:  LaunchContextLocal,
		RuntimeVersion: "go-test",
	}
	if configure != nil {
		configure(&config)
	}

	return httptest.NewServer(NewHandler(config)), fixtures
}

func writeGraphFixture(t *testing.T, repoPath string) {
	t.Helper()
	sourcePath := filepath.Join(repoPath, "src", "app.ts")
	if err := os.MkdirAll(filepath.Dir(sourcePath), 0o755); err != nil {
		t.Fatalf("mkdir source fixture: %v", err)
	}
	sourceContent := strings.Join([]string{
		"export function main() {",
		"export function helper() {",
		"  return main()",
		"}",
	}, "\n")
	if err := os.WriteFile(sourcePath, []byte(sourceContent), 0o644); err != nil {
		t.Fatalf("write source fixture: %v", err)
	}

	g := graph.New()
	g.AddNode(graph.Node{ID: "File:src/app.ts", Label: scopeir.NodeFile, Properties: graph.NodeProperties{
		"name": "app.ts", "filePath": "src/app.ts", "content": sourceContent,
	}})
	g.AddNode(graph.Node{ID: "Function:main", Label: scopeir.NodeFunction, Properties: graph.NodeProperties{
		"name": "main", "filePath": "src/app.ts", "startLine": 0, "endLine": 0,
	}})
	g.AddNode(graph.Node{ID: "Function:helper", Label: scopeir.NodeFunction, Properties: graph.NodeProperties{
		"name": "helper", "filePath": "src/app.ts", "startLine": 1, "endLine": 3,
	}})
	g.AddNode(graph.Node{ID: "Process:main", Label: scopeir.NodeProcess, Properties: graph.NodeProperties{
		"name": "main", "label": "main", "heuristicLabel": "main", "processType": "intra_community", "stepCount": 2,
	}})
	g.AddNode(graph.Node{ID: "comm_api", Label: scopeir.NodeCommunity, Properties: graph.NodeProperties{
		"name": "Api", "label": "Api", "heuristicLabel": "Api", "cohesion": 0.8, "symbolCount": 5,
	}})
	mainStep := 1
	g.AddRelationship(graph.Relationship{
		ID:         "rel:main-process",
		SourceID:   "Function:main",
		TargetID:   "Process:main",
		Type:       graph.RelStepInProcess,
		Step:       &mainStep,
		Reason:     "fixture",
		Confidence: 1,
	})
	helperStep := 2
	g.AddRelationship(graph.Relationship{
		ID:         "rel:helper-process",
		SourceID:   "Function:helper",
		TargetID:   "Process:main",
		Type:       graph.RelStepInProcess,
		Step:       &helperStep,
		Reason:     "fixture",
		Confidence: 1,
	})
	g.AddRelationship(graph.Relationship{
		ID:               "rel:helper-calls-main",
		SourceID:         "Function:helper",
		TargetID:         "Function:main",
		Type:             graph.RelCalls,
		Reason:           "fixture",
		Confidence:       1,
		ResolutionSource: "scope-resolution",
		FileHash:         "hash-graph",
		Evidence:         []graph.Evidence{{Kind: "type-binding", Weight: 0.35, Note: "receiver User"}},
	})
	g.AddRelationship(graph.Relationship{
		ID:         "rel:main-member",
		SourceID:   "Function:main",
		TargetID:   "comm_api",
		Type:       graph.RelMemberOf,
		Reason:     "fixture",
		Confidence: 1,
	})
	g.AddRelationship(graph.Relationship{
		ID:         "rel:helper-member",
		SourceID:   "Function:helper",
		TargetID:   "comm_api",
		Type:       graph.RelMemberOf,
		Reason:     "fixture",
		Confidence: 1,
	})
	raw, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		t.Fatalf("marshal graph fixture: %v", err)
	}
	graphPath := filepath.Join(repo.StoragePath(repoPath), "graph.json")
	if err := os.MkdirAll(filepath.Dir(graphPath), 0o755); err != nil {
		t.Fatalf("mkdir graph fixture: %v", err)
	}
	if err := os.WriteFile(graphPath, append(raw, '\n'), 0o644); err != nil {
		t.Fatalf("write graph fixture: %v", err)
	}
}

func getJSON(t *testing.T, url string, wantStatus int, target any) {
	t.Helper()

	resp, err := http.Get(url)
	if err != nil {
		t.Fatalf("GET %s failed: %v", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != wantStatus {
		body, _ := io.ReadAll(resp.Body)
		t.Fatalf("GET %s status = %d, want %d; body=%s", url, resp.StatusCode, wantStatus, body)
	}
	if err := json.NewDecoder(resp.Body).Decode(target); err != nil {
		t.Fatalf("decode %s JSON: %v", url, err)
	}
}

func intStats(files, nodes, edges int) *repo.Stats {
	return &repo.Stats{
		Files: &files,
		Nodes: &nodes,
		Edges: &edges,
	}
}

func findReportCandidate(candidates []graphHealthReportCandidate, nodeID string) *graphHealthReportCandidate {
	for index := range candidates {
		if candidates[index].NodeID == nodeID {
			return &candidates[index]
		}
	}
	return nil
}
