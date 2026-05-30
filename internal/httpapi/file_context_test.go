package httpapi

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"testing"

	"github.com/tamnguyendinh/anvien/internal/filecontext"
	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/repo"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
)

func TestFileContextEndpointReturnsProjectionForRegisteredRepo(t *testing.T) {
	server, fixtures := newRepoServer(t, []repoFixture{{name: "alpha"}})
	defer server.Close()
	writeHTTPFileProjectionGraph(t, fixtures[0].path)

	var payload filecontext.FileContext
	getJSON(t, server.URL+"/api/file-context?repo=alpha&path=src%2Fapp.go", http.StatusOK, &payload)

	if payload.Repo != "alpha" || payload.RepoPath != fixtures[0].path {
		t.Fatalf("payload repo binding = %q/%q", payload.Repo, payload.RepoPath)
	}
	if payload.Graph.Path == "" || payload.Graph.IndexedCommit != "commit" {
		t.Fatalf("payload graph metadata = %#v", payload.Graph)
	}
	if payload.Summary.Path != "src/app.go" || payload.Summary.SymbolCount != 2 ||
		payload.Summary.OutboundRefCount != 1 || payload.Summary.InboundRefCount != 1 ||
		payload.Summary.UnresolvedSourceSiteCount != 1 {
		t.Fatalf("summary = %#v, want file counts from projection graph", payload.Summary)
	}
	if len(payload.SymbolTree) != 2 || payload.SymbolTree[0].Name != "Server" || payload.SymbolTree[1].Name != "NewServer" {
		t.Fatalf("symbol tree = %#v", payload.SymbolTree)
	}
	if len(payload.Relationships.OutboundByFile) != 1 || payload.Relationships.OutboundByFile[0].File != "src/store.go" {
		t.Fatalf("outbound relationships = %#v", payload.Relationships.OutboundByFile)
	}
	if payload.Unresolved.Total != 1 || len(payload.Unresolved.Groups) != 1 ||
		payload.Unresolved.Groups[0].Samples[0].TargetText != "missingCall" {
		t.Fatalf("unresolved = %#v", payload.Unresolved)
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
	if payload.Files[0].Path != "src/app.go" || payload.Files[0].UnresolvedSourceSiteCount != 1 {
		t.Fatalf("top hotspot = %#v, want src/app.go unresolved=1", payload.Files[0])
	}

	getJSON(t, server.URL+"/api/file-hotspots?repo=workspace&unresolvedOnly=true", http.StatusOK, &payload)
	if payload.Total != 1 || len(payload.Files) != 1 || payload.Files[0].Path != "src/app.go" {
		t.Fatalf("unresolved-only hotspots = %#v", payload)
	}
}

func TestFileContextEndpointReportsMissingFile(t *testing.T) {
	server, fixtures := newRepoServer(t, []repoFixture{{name: "beta"}})
	defer server.Close()
	writeHTTPFileProjectionGraph(t, fixtures[0].path)

	var payload map[string]string
	getJSON(t, server.URL+"/api/file-context?repo="+url.QueryEscape("beta")+"&path=src%2Fmissing.go", http.StatusNotFound, &payload)

	if payload["error"] != "File not found in graph" {
		t.Fatalf("missing file error = %#v", payload)
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
