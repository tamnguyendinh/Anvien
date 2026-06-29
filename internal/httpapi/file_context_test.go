package httpapi

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/tamnguyendinh/anvien/internal/filecontext"
	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/repo"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
)

func TestFileDetailEndpointReturnsProjectionForRegisteredRepo(t *testing.T) {
	server, fixtures := newRepoServer(t, []repoFixture{{name: "alpha"}})
	defer server.Close()
	writeHTTPFileProjectionGraph(t, fixtures[0].path)

	var payload filecontext.CompactFileContext
	getJSON(t, server.URL+"/api/file-detail?repo=alpha&path=src%2Fapp.go", http.StatusOK, &payload)

	if payload.Format != filecontext.CompactFileContextFormat {
		t.Fatalf("payload format = %q, want compact format", payload.Format)
	}
	if payload.Repo != "alpha" || payload.RepoPath != fixtures[0].path {
		t.Fatalf("payload repo binding = %q/%q", payload.Repo, payload.RepoPath)
	}
	if payload.Graph.Path == "" || payload.Graph.IndexedCommit != "commit" {
		t.Fatalf("payload graph metadata = %#v", payload.Graph)
	}
	if payload.Summary.Path != "src/app.go" || payload.Summary.SymbolCount != 2 ||
		payload.Summary.OutboundRefCount != 2 || payload.Summary.InboundRefCount != 1 ||
		payload.Summary.Unresolved != 1 {
		t.Fatalf("summary = %#v, want file counts from projection graph", payload.Summary)
	}
	if payload.Limits.RelationshipSamplesPerGroup != filecontext.FullDetailSampleLimit ||
		payload.Limits.UnresolvedSamplesPerGroup != filecontext.FullDetailSampleLimit ||
		payload.Limits.LinkedSamplesPerKind != filecontext.FullDetailSampleLimit {
		t.Fatalf("compact default limits = %#v, want full-detail sentinel", payload.Limits)
	}
	if payload.Target.NormalizedPath != "src/app.go" {
		t.Fatalf("normalized path = %q, want src/app.go", payload.Target.NormalizedPath)
	}
	if len(payload.Tables.Symbols) != 2 {
		t.Fatalf("compact symbol rows = %#v, want two symbols", payload.Tables.Symbols)
	}
	if len(payload.Tables.RelatedFiles) != 2 {
		t.Fatalf("compact related-file rows = %#v, want store and test files", payload.Tables.RelatedFiles)
	}
	if len(payload.Tables.Relationships.OutboundByFile) != 1 {
		t.Fatalf("compact outbound groups = %#v", payload.Tables.Relationships.OutboundByFile)
	}
	outboundRows := payload.Tables.Relationships.OutboundByFile[0].Rows
	if outboundRows.Total != 2 || outboundRows.Returned != 2 || outboundRows.Omitted != 0 {
		t.Fatalf("compact default outbound rows = %#v, want full two-row group", outboundRows)
	}
	if payload.Tables.Unresolved.Total != 1 || len(payload.Tables.Unresolved.Groups) != 1 ||
		payload.Tables.Unresolved.Groups[0].Rows.Returned != 1 {
		t.Fatalf("compact unresolved = %#v", payload.Tables.Unresolved)
	}

	absolutePath := filepath.Join(fixtures[0].path, "src", "app.go")
	var absolutePayload filecontext.CompactFileContext
	getJSON(t, server.URL+"/api/file-detail?repo=alpha&path="+url.QueryEscape(absolutePath), http.StatusOK, &absolutePayload)
	if absolutePayload.Format != filecontext.CompactFileContextFormat {
		t.Fatalf("absolute payload format = %q, want compact format", absolutePayload.Format)
	}
	if absolutePayload.Summary.Path != "src/app.go" || absolutePayload.Target.NormalizedPath != "src/app.go" {
		t.Fatalf("absolute payload path/normalized = %q/%q, want src/app.go/src/app.go", absolutePayload.Summary.Path, absolutePayload.Target.NormalizedPath)
	}
	if absolutePayload.Target.Input != absolutePath {
		t.Fatalf("absolute payload input = %q, want %q", absolutePayload.Target.Input, absolutePath)
	}

	var expandedPayload filecontext.FileContext
	getJSON(t, server.URL+"/api/file-detail?repo=alpha&path=src%2Fapp.go&format=expanded", http.StatusOK, &expandedPayload)
	if expandedPayload.Repo != "alpha" || expandedPayload.Summary.Path != "src/app.go" {
		t.Fatalf("expanded payload repo/path = %q/%q", expandedPayload.Repo, expandedPayload.Summary.Path)
	}
	if expandedPayload.Limits.RelationshipSamplesPerGroup != fileContextDefaultSampleLimit {
		t.Fatalf("expanded default relationship limit = %d, want %d", expandedPayload.Limits.RelationshipSamplesPerGroup, fileContextDefaultSampleLimit)
	}
	if len(expandedPayload.SymbolTree) != 2 || expandedPayload.SymbolTree[0].Name != "Server" || expandedPayload.SymbolTree[1].Name != "NewServer" {
		t.Fatalf("expanded symbol tree = %#v", expandedPayload.SymbolTree)
	}

	var limitedPayload filecontext.CompactFileContext
	getJSON(t, server.URL+"/api/file-detail?repo=alpha&path=src%2Fapp.go&relationships=1", http.StatusOK, &limitedPayload)
	if limitedPayload.Limits.RelationshipSamplesPerGroup != 1 {
		t.Fatalf("limited relationship limit = %d, want 1", limitedPayload.Limits.RelationshipSamplesPerGroup)
	}
	limitedRows := limitedPayload.Tables.Relationships.OutboundByFile[0].Rows
	if limitedRows.Total != 2 || limitedRows.Returned != 1 || limitedRows.Omitted != 1 {
		t.Fatalf("limited compact outbound rows = %#v, want total/returned/omitted 2/1/1", limitedRows)
	}

	var formatError map[string]string
	getJSON(t, server.URL+"/api/file-detail?repo=alpha&path=src%2Fapp.go&format=summary", http.StatusBadRequest, &formatError)
	if formatError["error"] != `Unsupported "format" query parameter` {
		t.Fatalf("unsupported format error = %#v", formatError)
	}
}

func TestFileHotspotsEndpointReturnsSortedProjection(t *testing.T) {
	server, fixtures := newRepoServer(t, []repoFixture{{name: "workspace"}})
	defer server.Close()
	writeHTTPFileProjectionGraph(t, fixtures[0].path)

	var payload fileHotspotsResponse
	getJSON(t, server.URL+"/api/file-hotspots?repo=workspace&sort=unresolved&limit=2", http.StatusOK, &payload)

	if payload.Repo != "workspace" || payload.Total != 3 || len(payload.Files) != 2 {
		t.Fatalf("hotspots payload = %#v", payload)
	}
	if payload.Files[0].Path != "src/app.go" ||
		payload.Files[0].Unresolved != 1 ||
		payload.Files[0].Risk != "medium" {
		t.Fatalf("top hotspot = %#v, want src/app.go unresolved=1 risk=medium", payload.Files[0])
	}
	if len(payload.FileGroups) != 1 ||
		payload.FileGroups[0].Key != "backend_support_model_helper" ||
		payload.FileGroups[0].Label != "Backend support/model/helper files" ||
		payload.FileGroups[0].Files != 1 ||
		payload.FileGroups[0].Roles["storage_helper"] != 1 ||
		len(payload.FileGroups[0].SampleFiles) != 1 ||
		payload.FileGroups[0].SampleFiles[0] != "src/store.go" {
		t.Fatalf("fileGroups = %#v, want backend support/model/helper summary for src/store.go", payload.FileGroups)
	}

	getJSON(t, server.URL+"/api/file-hotspots?repo=workspace&unresolvedOnly=true", http.StatusOK, &payload)
	if payload.Total != 1 || len(payload.Files) != 1 || payload.Files[0].Path != "src/app.go" {
		t.Fatalf("unresolved-only hotspots = %#v", payload)
	}

	var errorPayload map[string]string
	getJSON(t, server.URL+"/api/file-hotspots?repo=workspace&sort=bad-sort", http.StatusBadRequest, &errorPayload)
	if errorPayload["error"] != `Unsupported "sort" query parameter` {
		t.Fatalf("retired sort error = %#v", errorPayload)
	}
}

func TestFileDetailEndpointReportsMissingFile(t *testing.T) {
	server, fixtures := newRepoServer(t, []repoFixture{{name: "beta"}})
	defer server.Close()
	writeHTTPFileProjectionGraph(t, fixtures[0].path)

	var payload map[string]string
	getJSON(t, server.URL+"/api/file-detail?repo="+url.QueryEscape("beta")+"&path=src%2Fmissing.go", http.StatusNotFound, &payload)

	if payload["error"] != "File not found in graph" {
		t.Fatalf("missing file error = %#v", payload)
	}
}

func TestFileDetailEndpointRejectsOutsideRepoAbsolutePath(t *testing.T) {
	server, fixtures := newRepoServer(t, []repoFixture{{name: "outside"}})
	defer server.Close()
	writeHTTPFileProjectionGraph(t, fixtures[0].path)

	outsidePath := filepath.Join(filepath.Dir(fixtures[0].path), "other-repo", "src", "app.go")
	var payload map[string]string
	getJSON(t, server.URL+"/api/file-detail?repo=outside&path="+url.QueryEscape(outsidePath), http.StatusBadRequest, &payload)

	if !strings.Contains(payload["error"], filecontext.ErrFilePathOutsideRepo.Error()) {
		t.Fatalf("outside repo error = %#v", payload)
	}
}

func TestFileContextEndpointIsNotRegistered(t *testing.T) {
	server, _ := newRepoServer(t, []repoFixture{{name: "legacy"}})
	defer server.Close()

	response, err := http.Get(server.URL + "/api/file-context?repo=legacy&path=src%2Fapp.go")
	if err != nil {
		t.Fatalf("GET legacy file-context endpoint: %v", err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusNotFound {
		t.Fatalf("legacy file-context endpoint status = %d, want 404", response.StatusCode)
	}
}

func writeHTTPFileProjectionGraph(t *testing.T, repoPath string) {
	t.Helper()
	for _, relPath := range []string{"src/app.go", "src/store.go", "src/app_test.go"} {
		fullPath := filepath.Join(repoPath, filepath.FromSlash(relPath))
		if err := os.MkdirAll(filepath.Dir(fullPath), 0o755); err != nil {
			t.Fatalf("mkdir source fixture: %v", err)
		}
		if err := os.WriteFile(fullPath, []byte("package src\n"), 0o644); err != nil {
			t.Fatalf("write source fixture: %v", err)
		}
	}

	g := graph.New()
	g.AddNode(graph.Node{ID: "File:src/app.go", Label: scopeir.NodeFile, Properties: graph.NodeProperties{
		"name": "app.go", "filePath": "src/app.go", "language": "go", "appLayer": "backend", "functionalArea": "mcp",
	}})
	g.AddNode(graph.Node{ID: "File:src/store.go", Label: scopeir.NodeFile, Properties: graph.NodeProperties{
		"name": "store.go", "filePath": "src/store.go", "language": "go", "appLayer": "backend", "functionalArea": "storage",
	}})
	g.AddNode(graph.Node{ID: "File:src/app_test.go", Label: scopeir.NodeFile, Properties: graph.NodeProperties{
		"name": "app_test.go", "filePath": "src/app_test.go", "language": "go", "appLayer": "backend_test", "functionalArea": "mcp",
	}})
	g.AddNode(graph.Node{ID: "Struct:src/app.go:Server", Label: scopeir.NodeStruct, Properties: graph.NodeProperties{
		"name": "Server", "filePath": "src/app.go", "startLine": 3, "endLine": 8, "visibility": "public",
	}})
	g.AddNode(graph.Node{ID: "Function:src/app.go:NewServer", Label: scopeir.NodeFunction, Properties: graph.NodeProperties{
		"name": "NewServer", "filePath": "src/app.go", "startLine": 10, "endLine": 16,
	}})
	g.AddNode(graph.Node{ID: "Function:src/store.go:Save", Label: scopeir.NodeFunction, Properties: graph.NodeProperties{
		"name": "Save", "filePath": "src/store.go", "startLine": 5, "endLine": 9,
	}})
	g.AddNode(graph.Node{ID: "Function:src/app_test.go:TestNewServer", Label: scopeir.NodeFunction, Properties: graph.NodeProperties{
		"name": "TestNewServer", "filePath": "src/app_test.go", "startLine": 4, "endLine": 12,
	}})
	g.AddNode(graph.Node{ID: "ResolutionGap:src/app.go:missing", Label: scopeir.NodeResolutionGap, Properties: graph.NodeProperties{
		"name": "missingCall", "filePath": "src/app.go", "sourceNodeId": "Function:src/app.go:NewServer",
		"targetText": "missingCall", "gapKind": "unresolved_call", "classification": "in_repo_unresolved",
		"actionability": "analyzer_gap", "sourceSiteId": "SourceSite:src/app.go#call#missingCall#12#2#12#13",
		"sourceSiteStatus": "unresolved_local_binding", "startLine": 12, "startCol": 2,
	}})
	g.AddRelationship(graph.Relationship{ID: "rel:def-app-server", SourceID: "File:src/app.go", TargetID: "Struct:src/app.go:Server", Type: graph.RelDefines})
	g.AddRelationship(graph.Relationship{ID: "rel:def-app-new", SourceID: "File:src/app.go", TargetID: "Function:src/app.go:NewServer", Type: graph.RelDefines})
	g.AddRelationship(graph.Relationship{ID: "rel:def-store-save", SourceID: "File:src/store.go", TargetID: "Function:src/store.go:Save", Type: graph.RelDefines})
	g.AddRelationship(graph.Relationship{ID: "rel:def-test", SourceID: "File:src/app_test.go", TargetID: "Function:src/app_test.go:TestNewServer", Type: graph.RelDefines})
	g.AddRelationship(graph.Relationship{
		ID: "rel:call-new-save", SourceID: "Function:src/app.go:NewServer", TargetID: "Function:src/store.go:Save", Type: graph.RelCalls,
		FilePath: "src/app.go", SourceSiteID: "SourceSite:src/app.go#call#Save#13#2#13#8", SourceSiteStatus: "resolved", ProofKind: "scope-binding",
	})
	g.AddRelationship(graph.Relationship{
		ID: "rel:call-server-save", SourceID: "Struct:src/app.go:Server", TargetID: "Function:src/store.go:Save", Type: graph.RelCalls,
		FilePath: "src/app.go", SourceSiteID: "SourceSite:src/app.go#call#Save#7#2#7#8", SourceSiteStatus: "resolved", ProofKind: "scope-binding",
	})
	g.AddRelationship(graph.Relationship{
		ID: "rel:call-test-new", SourceID: "Function:src/app_test.go:TestNewServer", TargetID: "Function:src/app.go:NewServer", Type: graph.RelCalls,
		FilePath: "src/app_test.go", SourceSiteID: "SourceSite:src/app_test.go#call#NewServer#6#2#6#13", SourceSiteStatus: "resolved", ProofKind: "scope-binding",
	})

	raw, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		t.Fatalf("marshal graph fixture: %v", err)
	}
	graphPath := filepath.Join(repo.StoragePath(repoPath), "graph.json")
	if err := os.MkdirAll(filepath.Dir(graphPath), 0o755); err != nil {
		t.Fatalf("mkdir graph dir: %v", err)
	}
	if err := os.WriteFile(graphPath, append(raw, '\n'), 0o644); err != nil {
		t.Fatalf("write graph fixture: %v", err)
	}
}
