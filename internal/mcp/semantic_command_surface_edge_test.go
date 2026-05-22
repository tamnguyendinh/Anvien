package mcp

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/tamnguyendinh/avmatrix-go/internal/graph"
	"github.com/tamnguyendinh/avmatrix-go/internal/repo"
	"github.com/tamnguyendinh/avmatrix-go/internal/scopeir"
	"github.com/tamnguyendinh/avmatrix-go/internal/semantic"
)

func TestImpactToolWarnsForStaleIncompleteSemanticMetadata(t *testing.T) {
	g := graph.New()
	target := graph.Node{ID: "Function:Target", Label: scopeir.NodeFunction, Properties: graph.NodeProperties{
		"name": "Target", "filePath": "src/target.go",
	}}
	caller := graph.Node{ID: "Function:Caller", Label: scopeir.NodeFunction, Properties: graph.NodeProperties{
		"name": "Caller", "filePath": "src/caller.go",
	}}
	g.AddNode(target)
	g.AddNode(caller)
	g.AddRelationship(graph.Relationship{ID: "rel:caller-target", SourceID: caller.ID, TargetID: target.ID, Type: graph.RelCalls, Confidence: 1})

	payload, _ := runImpactBFSProfiled(g, target, impactOptions{
		Direction:     "upstream",
		MaxDepth:      1,
		RelationTypes: []string{string(graph.RelCalls)},
		IncludeTests:  true,
	}, false)

	status := payload["semanticStatus"].(semantic.GraphStatus)
	if status.AppLayer.Status != semantic.StatusStaleIncomplete || status.FunctionalArea.Status != semantic.StatusStaleIncomplete {
		t.Fatalf("semanticStatus = %#v, want stale incomplete", status)
	}
	if warning, ok := payload["semanticWarning"].(string); !ok || warning == "" {
		t.Fatalf("impact payload missing stale semantic warning: %#v", payload)
	}
	targetPayload := payload["target"].(map[string]any)
	if _, ok := targetPayload["appLayer"]; ok {
		t.Fatalf("impact target should not invent App Layer: %#v", targetPayload)
	}
}

func TestDetectChangesToolWarnsForStaleIncompleteSemanticMetadata(t *testing.T) {
	if _, err := exec.LookPath("git"); err != nil {
		t.Skip("git executable is not available")
	}

	store := repo.NewStore(t.TempDir())
	repoPath := t.TempDir()
	sourceDir := filepath.Join(repoPath, "src")
	if err := os.MkdirAll(sourceDir, 0o755); err != nil {
		t.Fatalf("mkdir source: %v", err)
	}
	sourcePath := filepath.Join(sourceDir, "app.ts")
	if err := os.WriteFile(sourcePath, []byte("function main() {\n  return 1\n}\n"), 0o644); err != nil {
		t.Fatalf("write source: %v", err)
	}
	runMCPTestGit(t, repoPath, "init")
	runMCPTestGit(t, repoPath, "config", "user.email", "test@example.com")
	runMCPTestGit(t, repoPath, "config", "user.name", "Test User")
	runMCPTestGit(t, repoPath, "add", "src/app.ts")
	runMCPTestGit(t, repoPath, "commit", "-m", "initial")

	meta := repo.Meta{RepoPath: repoPath, IndexedAt: "2026-05-22T00:00:00Z", LastCommit: "abc123", Stats: &repo.Stats{}}
	if err := repo.SaveMeta(repo.StoragePath(repoPath), meta); err != nil {
		t.Fatalf("save meta: %v", err)
	}
	if _, err := store.Register(repoPath, meta, repo.RegisterOptions{Name: "fixture"}); err != nil {
		t.Fatalf("register repo: %v", err)
	}
	g := graph.New()
	g.AddNode(graph.Node{ID: "Function:main", Label: scopeir.NodeFunction, Properties: graph.NodeProperties{
		"name": "main", "filePath": "src/app.ts", "startLine": 1, "endLine": 3,
	}})
	writeMCPGraphTB(t, repoPath, g)

	if err := os.WriteFile(sourcePath, []byte("function main() {\n  return 2\n}\n"), 0o644); err != nil {
		t.Fatalf("modify source: %v", err)
	}

	server := NewServer(Config{Store: store})
	payload, err := server.detectChangesTool(map[string]any{"repo": "fixture", "scope": "unstaged"})
	if err != nil {
		t.Fatalf("detectChangesTool() error = %v", err)
	}
	status := payload["semanticStatus"].(semantic.GraphStatus)
	if status.AppLayer.Status != semantic.StatusStaleIncomplete || status.FunctionalArea.Status != semantic.StatusStaleIncomplete {
		t.Fatalf("semanticStatus = %#v, want stale incomplete", status)
	}
	if warning, ok := payload["semanticWarning"].(string); !ok || warning == "" {
		t.Fatalf("detect-changes payload missing stale semantic warning: %#v", payload)
	}
	changed := payload["changed_symbols"].([]map[string]any)
	if _, ok := changed[0]["appLayer"]; ok {
		t.Fatalf("detect-changes should not invent App Layer: %#v", changed[0])
	}
}

func TestAPIMCPToolsWarnForStaleAndDoNotInventSemanticFields(t *testing.T) {
	store, repoPath := newMCPQueryBenchmarkRepo(t)
	g := graph.New()
	g.AddNode(graph.Node{ID: "File:src/app.ts", Label: scopeir.NodeFile, Properties: graph.NodeProperties{
		"name": "app.ts", "filePath": "src/app.ts",
	}})
	g.AddNode(graph.Node{ID: "Route:/api/users", Label: scopeir.NodeRoute, Properties: graph.NodeProperties{
		"name": "/api/users", "filePath": "src/server.ts", "responseKeys": []string{"data"},
	}})
	g.AddRelationship(graph.Relationship{
		ID: "rel:app-fetch-users", SourceID: "File:src/app.ts", TargetID: "Route:/api/users", Type: graph.RelFetches, Reason: "fetch-route|keys:missing|fetches:1",
	})
	writeMCPGraphTB(t, repoPath, g)

	server := NewServer(Config{Store: store})
	routePayload, err := server.routeMapTool(map[string]any{"repo": "fixture", "route": "/api/users"})
	if err != nil {
		t.Fatalf("routeMapTool() error = %v", err)
	}
	if warning, ok := routePayload["semanticWarning"].(string); !ok || warning == "" {
		t.Fatalf("route_map payload missing stale semantic warning: %#v", routePayload)
	}
	route := routePayload["routes"].([]mcpRouteMapItem)[0]
	if route.AppLayer != "" || route.Consumers[0].AppLayer != "" {
		t.Fatalf("route_map should not invent semantic fields: %#v", route)
	}

	shapePayload, err := server.shapeCheckTool(map[string]any{"repo": "fixture", "route": "/api/users"})
	if err != nil {
		t.Fatalf("shapeCheckTool() error = %v", err)
	}
	if warning, ok := shapePayload["semanticWarning"].(string); !ok || warning == "" {
		t.Fatalf("shape_check payload missing stale semantic warning: %#v", shapePayload)
	}
	shapeRoute := shapePayload["routes"].([]mcpShapeRoute)[0]
	if shapeRoute.AppLayer != "" || shapeRoute.Consumers[0].AppLayer != "" {
		t.Fatalf("shape_check should not invent semantic fields: %#v", shapeRoute)
	}

	apiPayload, err := server.apiImpactTool(map[string]any{"repo": "fixture", "route": "/api/users"})
	if err != nil {
		t.Fatalf("apiImpactTool() error = %v", err)
	}
	if warning, ok := apiPayload["semanticWarning"].(string); !ok || warning == "" {
		t.Fatalf("api_impact payload missing stale semantic warning: %#v", apiPayload)
	}
	if _, ok := apiPayload["appLayer"]; ok {
		t.Fatalf("api_impact should not invent route App Layer: %#v", apiPayload)
	}
}
