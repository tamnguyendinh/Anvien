package cli

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/tamnguyendinh/avmatrix-go/internal/graph"
	"github.com/tamnguyendinh/avmatrix-go/internal/graphhealth"
	"github.com/tamnguyendinh/avmatrix-go/internal/repo"
	"github.com/tamnguyendinh/avmatrix-go/internal/scopeir"
)

func TestAPICommandsUseLocalMCPRuntimeJSON(t *testing.T) {
	home := t.TempDir()
	t.Setenv(repo.HomeEnvName, home)
	repoPath := t.TempDir()
	registerDirectToolCommandRepo(t, repo.NewStore(home), repoPath, "fixture")
	writeCLIAPIParityGraph(t, repoPath)

	out, errOut, err := executeForTest(t, "api", "route-map", "/api/users", "--repo", "fixture", "--json")
	if err != nil {
		t.Fatalf("api route-map returned error: %v\nstdout:\n%s\nstderr:\n%s", err, out, errOut)
	}
	if errOut != "" {
		t.Fatalf("api route-map wrote stderr: %q", errOut)
	}
	var routePayload struct {
		Routes []struct {
			Route     string `json:"route"`
			Handler   string `json:"handler"`
			Consumers []struct {
				FilePath     string   `json:"filePath"`
				AccessedKeys []string `json:"accessedKeys"`
				FetchCount   int      `json:"fetchCount"`
			} `json:"consumers"`
			Flows []string `json:"flows"`
		} `json:"routes"`
		Total int `json:"total"`
	}
	if err := json.Unmarshal([]byte(out), &routePayload); err != nil {
		t.Fatalf("parse route-map JSON: %v\n%s", err, out)
	}
	if routePayload.Total != 1 || len(routePayload.Routes) != 1 {
		t.Fatalf("route-map payload = %#v", routePayload)
	}
	if routePayload.Routes[0].Route != "/api/users" || routePayload.Routes[0].Handler != "src/server.ts" || len(routePayload.Routes[0].Consumers) != 1 || routePayload.Routes[0].Consumers[0].FilePath != "src/app.ts" {
		t.Fatalf("route-map route = %#v", routePayload.Routes[0])
	}
	if len(routePayload.Routes[0].Flows) != 1 || routePayload.Routes[0].Flows[0] != "MainFlow" || routePayload.Routes[0].Consumers[0].FetchCount != 2 {
		t.Fatalf("route-map flow/consumer detail = %#v", routePayload.Routes[0])
	}

	out, errOut, err = executeForTest(t, "api", "tool-map", "--tool", "search", "--repo", "fixture", "--json")
	if err != nil {
		t.Fatalf("api tool-map returned error: %v\nstdout:\n%s\nstderr:\n%s", err, out, errOut)
	}
	if errOut != "" {
		t.Fatalf("api tool-map wrote stderr: %q", errOut)
	}
	var toolPayload struct {
		Tools []struct {
			Name        string   `json:"name"`
			FilePath    string   `json:"filePath"`
			Description string   `json:"description"`
			Flows       []string `json:"flows"`
		} `json:"tools"`
		Total int `json:"total"`
	}
	if err := json.Unmarshal([]byte(out), &toolPayload); err != nil {
		t.Fatalf("parse tool-map JSON: %v\n%s", err, out)
	}
	if toolPayload.Total != 1 || len(toolPayload.Tools) != 1 || toolPayload.Tools[0].Name != "search" || toolPayload.Tools[0].FilePath != "src/tools.ts" || toolPayload.Tools[0].Description != "Search code intelligence" {
		t.Fatalf("tool-map payload = %#v", toolPayload)
	}
	if len(toolPayload.Tools[0].Flows) != 1 || toolPayload.Tools[0].Flows[0] != "MainFlow" {
		t.Fatalf("tool-map flows = %#v", toolPayload.Tools[0])
	}

	out, errOut, err = executeForTest(t, "api", "shape-check", "--route", "/api/users", "--repo", "fixture", "--json")
	if err != nil {
		t.Fatalf("api shape-check returned error: %v\nstdout:\n%s\nstderr:\n%s", err, out, errOut)
	}
	if errOut != "" {
		t.Fatalf("api shape-check wrote stderr: %q", errOut)
	}
	var shapePayload struct {
		Routes []struct {
			Route     string `json:"route"`
			Status    string `json:"status"`
			Consumers []struct {
				Mismatched []string `json:"mismatched"`
			} `json:"consumers"`
		} `json:"routes"`
		Total      int `json:"total"`
		Mismatches int `json:"mismatches"`
	}
	if err := json.Unmarshal([]byte(out), &shapePayload); err != nil {
		t.Fatalf("parse shape-check JSON: %v\n%s", err, out)
	}
	if shapePayload.Total != 1 || shapePayload.Mismatches != 1 || len(shapePayload.Routes) != 1 || shapePayload.Routes[0].Status != "MISMATCH" {
		t.Fatalf("shape-check payload = %#v", shapePayload)
	}
	if len(shapePayload.Routes[0].Consumers) != 1 || !stringSliceContains(shapePayload.Routes[0].Consumers[0].Mismatched, "missing") {
		t.Fatalf("shape-check mismatches = %#v", shapePayload.Routes[0])
	}

	out, errOut, err = executeForTest(t, "api", "impact", "/api/users", "--repo", "fixture", "--json")
	if err != nil {
		t.Fatalf("api impact returned error: %v\nstdout:\n%s\nstderr:\n%s", err, out, errOut)
	}
	if errOut != "" {
		t.Fatalf("api impact wrote stderr: %q", errOut)
	}
	var impactPayload struct {
		Route         string `json:"route"`
		Handler       string `json:"handler"`
		ImpactSummary struct {
			DirectConsumers float64 `json:"directConsumers"`
			AffectedFlows   float64 `json:"affectedFlows"`
			RiskLevel       string  `json:"riskLevel"`
		} `json:"impactSummary"`
		Mismatches []struct {
			Field string `json:"field"`
		} `json:"mismatches"`
	}
	if err := json.Unmarshal([]byte(out), &impactPayload); err != nil {
		t.Fatalf("parse api impact JSON: %v\n%s", err, out)
	}
	if impactPayload.Route != "/api/users" || impactPayload.Handler != "src/server.ts" || impactPayload.ImpactSummary.DirectConsumers != 1 || impactPayload.ImpactSummary.AffectedFlows != 1 || impactPayload.ImpactSummary.RiskLevel != "MEDIUM" {
		t.Fatalf("api impact payload = %#v", impactPayload)
	}
	if len(impactPayload.Mismatches) != 1 || impactPayload.Mismatches[0].Field != "missing" {
		t.Fatalf("api impact mismatches = %#v", impactPayload.Mismatches)
	}
}

func TestRenameCommandDryRunUsesLocalMCPRuntimeJSON(t *testing.T) {
	home := t.TempDir()
	t.Setenv(repo.HomeEnvName, home)
	repoPath := t.TempDir()
	registerDirectToolCommandRepo(t, repo.NewStore(home), repoPath, "fixture")
	writeCLIRenameFixture(t, repoPath)

	out, errOut, err := executeForTest(t, "rename", "helper", "renamedHelper", "--repo", "fixture", "--json")
	if err != nil {
		t.Fatalf("rename returned error: %v\nstdout:\n%s\nstderr:\n%s", err, out, errOut)
	}
	if errOut != "" {
		t.Fatalf("rename wrote stderr: %q", errOut)
	}
	if strings.Contains(out, "---") || strings.Contains(out, "detect_changes") {
		t.Fatalf("rename --json included MCP next-step text:\n%s", out)
	}
	var payload struct {
		Status        string `json:"status"`
		OldName       string `json:"old_name"`
		NewName       string `json:"new_name"`
		Applied       bool   `json:"applied"`
		FilesAffected int    `json:"files_affected"`
		TotalEdits    int    `json:"total_edits"`
		GraphEdits    int    `json:"graph_edits"`
		Changes       []struct {
			FilePath string `json:"file_path"`
			Edits    []struct {
				Line       int    `json:"line"`
				Confidence string `json:"confidence"`
			} `json:"edits"`
		} `json:"changes"`
	}
	if err := json.Unmarshal([]byte(out), &payload); err != nil {
		t.Fatalf("parse rename JSON: %v\n%s", err, out)
	}
	if payload.Status != "success" || payload.OldName != "helper" || payload.NewName != "renamedHelper" || payload.Applied {
		t.Fatalf("rename payload = %#v", payload)
	}
	if payload.FilesAffected != 2 || payload.TotalEdits != 2 || payload.GraphEdits != 2 || len(payload.Changes) != 2 {
		t.Fatalf("rename counts = %#v", payload)
	}
	sourceRaw, err := os.ReadFile(filepath.Join(repoPath, "src", "main.ts"))
	if err != nil {
		t.Fatalf("read source after dry run: %v", err)
	}
	if strings.Contains(string(sourceRaw), "renamedHelper") {
		t.Fatalf("rename dry run changed source:\n%s", sourceRaw)
	}
}

func TestAPICommandRejectsDuplicateSelector(t *testing.T) {
	out, errOut, err := executeForTest(t, "api", "route-map", "/api/users", "--route", "/api/orders")
	if err == nil || !strings.Contains(err.Error(), "provide route as either positional argument or --route, not both") {
		t.Fatalf("route-map duplicate selector err = %v\nstdout:\n%s\nstderr:\n%s", err, out, errOut)
	}
}

func writeCLIAPIParityGraph(t *testing.T, repoPath string) {
	t.Helper()
	g := graph.New()
	g.AddNode(graph.Node{ID: "File:src/app.ts", Label: scopeir.NodeFile, Properties: graph.NodeProperties{
		"name": "app.ts", "filePath": "src/app.ts",
		"appLayer": "frontend", "appLayerSource": "test_fixture", "functionalArea": "web_graph_ui", "functionalAreaSource": "test_fixture",
		"topologyStatus": string(graphhealth.TopologyConnected), "resolutionConfidence": graphhealth.ResolutionConfidenceDegraded, "resolutionGapCount": 1,
		"resolutionHealthBuckets": map[string]int{string(graphhealth.ResolutionHealthUnresolvedCallTarget): 1},
	}})
	g.AddNode(graph.Node{ID: "Route:/api/users", Label: scopeir.NodeRoute, Properties: graph.NodeProperties{
		"name": "/api/users", "filePath": "src/server.ts", "responseKeys": []string{"data", "total"}, "errorKeys": []string{"error"}, "middleware": []string{"withAuth"},
		"appLayer": "api", "appLayerSource": "test_fixture", "functionalArea": "api", "functionalAreaSource": "test_fixture",
		"topologyStatus": string(graphhealth.TopologyConnected), "resolutionConfidence": graphhealth.ResolutionConfidenceClear,
		"resolutionHealthBuckets": map[string]int{string(graphhealth.ResolutionHealthResolvedReferences): 2},
	}})
	g.AddNode(graph.Node{ID: "Tool:search", Label: scopeir.NodeTool, Properties: graph.NodeProperties{
		"name": "search", "filePath": "src/tools.ts", "description": "Search code intelligence",
	}})
	g.AddNode(graph.Node{ID: "Process:main", Label: scopeir.NodeProcess, Properties: graph.NodeProperties{
		"name": "MainFlow", "label": "MainFlow", "heuristicLabel": "MainFlow", "processType": "route", "stepCount": 1,
		"appLayer": "api", "appLayerSource": "test_fixture", "functionalArea": "api", "functionalAreaSource": "test_fixture",
	}})
	g.AddRelationship(graph.Relationship{ID: "rel:app-fetch-users", SourceID: "File:src/app.ts", TargetID: "Route:/api/users", Type: graph.RelFetches, Confidence: 0.9, Reason: "fetch-route|keys:data,total,missing|fetches:2"})
	g.AddRelationship(graph.Relationship{ID: "rel:route-main-flow", SourceID: "Route:/api/users", TargetID: "Process:main", Type: graph.RelEntryPointOf, Confidence: 0.8, Reason: "fixture"})
	g.AddRelationship(graph.Relationship{ID: "rel:tool-main-flow", SourceID: "Tool:search", TargetID: "Process:main", Type: graph.RelEntryPointOf, Confidence: 0.8, Reason: "fixture"})
	writeGroupCommandGraph(t, repoPath, g)
}

func writeCLIRenameFixture(t *testing.T, repoPath string) {
	t.Helper()
	sourceDir := filepath.Join(repoPath, "src")
	if err := os.MkdirAll(sourceDir, 0o755); err != nil {
		t.Fatalf("mkdir source: %v", err)
	}
	if err := os.WriteFile(filepath.Join(sourceDir, "helper.ts"), []byte("export function helper() {\n  return 1\n}\n"), 0o644); err != nil {
		t.Fatalf("write helper source: %v", err)
	}
	if err := os.WriteFile(filepath.Join(sourceDir, "main.ts"), []byte("function main() {\n  return helper()\n}\n"), 0o644); err != nil {
		t.Fatalf("write main source: %v", err)
	}

	g := graph.New()
	g.AddNode(graph.Node{ID: "Function:helper", Label: scopeir.NodeFunction, Properties: graph.NodeProperties{
		"name": "helper", "filePath": "src/helper.ts", "startLine": 1, "endLine": 3,
	}})
	g.AddNode(graph.Node{ID: "Function:main", Label: scopeir.NodeFunction, Properties: graph.NodeProperties{
		"name": "main", "filePath": "src/main.ts", "startLine": 1, "endLine": 3,
	}})
	g.AddRelationship(graph.Relationship{
		ID:               "rel:CALLS:Function:main->Function:helper:2:10",
		SourceID:         "Function:main",
		TargetID:         "Function:helper",
		Type:             graph.RelCalls,
		Confidence:       0.95,
		Reason:           "scope-resolution: call",
		ResolutionSource: "scope-resolution",
	})
	writeGroupCommandGraph(t, repoPath, g)
}

func stringSliceContains(values []string, target string) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}
	return false
}
