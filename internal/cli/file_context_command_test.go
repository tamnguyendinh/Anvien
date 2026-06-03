package cli

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/tamnguyendinh/anvien/internal/filecontext"
	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/repo"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
)

func TestFileContextCommandOutputsFileProjection(t *testing.T) {
	home := t.TempDir()
	t.Setenv(repo.HomeEnvName, home)
	repoPath := t.TempDir()
	registerDirectToolCommandRepo(t, repo.NewStore(home), repoPath, "fixture")
	writeFileProjectionCommandGraph(t, repoPath)

	out, errOut, err := executeForTest(t, "file-context", "src/app.go", "--repo", "fixture", "--json")
	if err != nil {
		t.Fatalf("file-context returned error: %v\nstdout:\n%s\nstderr:\n%s", err, out, errOut)
	}
	if errOut != "" {
		t.Fatalf("file-context wrote stderr: %q", errOut)
	}
	var payload filecontext.FileContext
	if err := json.Unmarshal([]byte(out), &payload); err != nil {
		t.Fatalf("parse file-context JSON: %v\n%s", err, out)
	}
	if payload.Repo != "fixture" || payload.Summary.Path != "src/app.go" {
		t.Fatalf("payload repo/path = %q/%q", payload.Repo, payload.Summary.Path)
	}
	if payload.Summary.SymbolCount != 2 || payload.Summary.OutboundRefCount != 1 || payload.Summary.UnresolvedSourceSiteCount != 1 {
		t.Fatalf("summary = %#v, want symbols=2 outbound=1 unresolved=1", payload.Summary)
	}
	if len(payload.SymbolTree) != 2 || payload.SymbolTree[0].Name != "Server" || payload.SymbolTree[1].Name != "NewServer" {
		t.Fatalf("symbol tree = %#v", payload.SymbolTree)
	}

	out, errOut, err = executeForTest(t, "file-context", "src/app.go", "--repo", "fixture")
	if err != nil {
		t.Fatalf("file-context human returned error: %v\nstdout:\n%s\nstderr:\n%s", err, out, errOut)
	}
	if errOut != "" {
		t.Fatalf("file-context human wrote stderr: %q", errOut)
	}
	for _, want := range []string{"File: src/app.go", "Symbols: 2 exported=1", "Outbound files:", "src/store.go", "Unresolved source sites: total=1"} {
		if !strings.Contains(out, want) {
			t.Fatalf("file-context human output missing %q:\n%s", want, out)
		}
	}
}

func TestFileHotspotsCommandOutputsSortedProjection(t *testing.T) {
	home := t.TempDir()
	t.Setenv(repo.HomeEnvName, home)
	repoPath := t.TempDir()
	registerDirectToolCommandRepo(t, repo.NewStore(home), repoPath, "fixture")
	writeFileProjectionCommandGraph(t, repoPath)

	out, errOut, err := executeForTest(t, "file-hotspots", "--repo", "fixture", "--sort", "unresolved", "--limit", "2", "--json")
	if err != nil {
		t.Fatalf("file-hotspots returned error: %v\nstdout:\n%s\nstderr:\n%s", err, out, errOut)
	}
	if errOut != "" {
		t.Fatalf("file-hotspots wrote stderr: %q", errOut)
	}
	var payload struct {
		Repo  string                    `json:"repo"`
		Total int                       `json:"total"`
		Files []filecontext.FileSummary `json:"files"`
	}
	if err := json.Unmarshal([]byte(out), &payload); err != nil {
		t.Fatalf("parse file-hotspots JSON: %v\n%s", err, out)
	}
	if payload.Repo != "fixture" || payload.Total != 3 || len(payload.Files) != 2 {
		t.Fatalf("hotspots payload = %#v", payload)
	}
	if payload.Files[0].Path != "src/app.go" ||
		payload.Files[0].DefaultVisibleUnresolvedSourceSiteCount != 1 ||
		payload.Files[0].RawUnresolvedSourceSiteCount != 1 ||
		payload.Files[0].Risk != "medium" {
		t.Fatalf("top hotspot = %#v, want src/app.go default/raw unresolved=1 risk=medium", payload.Files[0])
	}

	out, errOut, err = executeForTest(t, "file-hotspots", "--repo", "fixture", "--unresolved-only")
	if err != nil {
		t.Fatalf("file-hotspots human returned error: %v\nstdout:\n%s\nstderr:\n%s", err, out, errOut)
	}
	if errOut != "" {
		t.Fatalf("file-hotspots human wrote stderr: %q", errOut)
	}
	for _, want := range []string{"File hotspots for fixture: total=1", "Path\tGroup\tRole\tLayer\tArea", "src/app.go"} {
		if !strings.Contains(out, want) {
			t.Fatalf("file-hotspots human output missing %q:\n%s", want, out)
		}
	}
	if strings.Contains(out, "src/app_test.go") {
		t.Fatalf("default unresolved-only output exposed test unresolved file:\n%s", out)
	}

	out, errOut, err = executeForTest(t, "file-hotspots", "--repo", "fixture", "--sort", "test-unresolved", "--json")
	if err == nil {
		t.Fatalf("file-hotspots accepted retired test-unresolved sort\nstdout:\n%s\nstderr:\n%s", out, errOut)
	}
	if !strings.Contains(err.Error(), `unsupported sort "test-unresolved"`) {
		t.Fatalf("retired sort error = %v\nstdout:\n%s\nstderr:\n%s", err, out, errOut)
	}
}

func TestFileContextCommandReportsMissingFile(t *testing.T) {
	home := t.TempDir()
	t.Setenv(repo.HomeEnvName, home)
	repoPath := t.TempDir()
	registerDirectToolCommandRepo(t, repo.NewStore(home), repoPath, "fixture")
	writeFileProjectionCommandGraph(t, repoPath)

	out, errOut, err := executeForTest(t, "file-context", "src/missing.go", "--repo", "fixture")
	if err == nil {
		t.Fatalf("file-context missing file succeeded\nstdout:\n%s\nstderr:\n%s", out, errOut)
	}
	if !strings.Contains(err.Error(), `file "src/missing.go" not found`) {
		t.Fatalf("missing file error = %v", err)
	}
}

func writeFileProjectionCommandGraph(t *testing.T, repoPath string) {
	t.Helper()
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
		t.Fatalf("marshal graph: %v", err)
	}
	graphPath := filepath.Join(repo.StoragePath(repoPath), "graph.json")
	if err := os.MkdirAll(filepath.Dir(graphPath), 0o755); err != nil {
		t.Fatalf("mkdir graph dir: %v", err)
	}
	if err := os.WriteFile(graphPath, append(raw, '\n'), 0o644); err != nil {
		t.Fatalf("write graph: %v", err)
	}
}
