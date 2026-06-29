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

func TestFileDetailCommandOutputsFileProjection(t *testing.T) {
	home := t.TempDir()
	t.Setenv(repo.HomeEnvName, home)
	repoPath := t.TempDir()
	registerDirectToolCommandRepo(t, repo.NewStore(home), repoPath, "fixture")
	writeFileProjectionCommandGraph(t, repoPath)

	out, errOut, err := executeForTest(t, "file-detail", "src/app.go", "--repo", "fixture", "--json")
	if err != nil {
		t.Fatalf("file-detail returned error: %v\nstdout:\n%s\nstderr:\n%s", err, out, errOut)
	}
	if errOut != "" {
		t.Fatalf("file-detail wrote stderr: %q", errOut)
	}
	var payload filecontext.CompactFileContext
	if err := json.Unmarshal([]byte(out), &payload); err != nil {
		t.Fatalf("parse file-detail JSON: %v\n%s", err, out)
	}
	if payload.Format != filecontext.CompactFileContextFormat {
		t.Fatalf("payload format = %q, want compact format", payload.Format)
	}
	if payload.Repo != "fixture" || payload.Summary.Path != "src/app.go" {
		t.Fatalf("payload repo/path = %q/%q", payload.Repo, payload.Summary.Path)
	}
	if payload.Target.NormalizedPath != "src/app.go" {
		t.Fatalf("payload normalized path = %q, want src/app.go", payload.Target.NormalizedPath)
	}
	if payload.Summary.SymbolCount != 2 || payload.Summary.OutboundRefCount != 2 || payload.Summary.Unresolved != 1 {
		t.Fatalf("summary = %#v, want symbols=2 outbound=2 unresolved=1", payload.Summary)
	}
	if payload.Limits.RelationshipSamplesPerGroup != filecontext.FullDetailSampleLimit ||
		payload.Limits.UnresolvedSamplesPerGroup != filecontext.FullDetailSampleLimit ||
		payload.Limits.LinkedSamplesPerKind != filecontext.FullDetailSampleLimit {
		t.Fatalf("compact default limits = %#v, want full-detail sentinel", payload.Limits)
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

	absolutePath := filepath.Join(repoPath, "src", "app.go")
	out, errOut, err = executeForTest(t, "file-detail", absolutePath, "--repo", "fixture", "--json")
	if err != nil {
		t.Fatalf("file-detail absolute path returned error: %v\nstdout:\n%s\nstderr:\n%s", err, out, errOut)
	}
	if errOut != "" {
		t.Fatalf("file-detail absolute path wrote stderr: %q", errOut)
	}
	var absolutePayload filecontext.CompactFileContext
	if err := json.Unmarshal([]byte(out), &absolutePayload); err != nil {
		t.Fatalf("parse absolute file-detail JSON: %v\n%s", err, out)
	}
	if absolutePayload.Format != filecontext.CompactFileContextFormat {
		t.Fatalf("absolute payload format = %q, want compact format", absolutePayload.Format)
	}
	if absolutePayload.Summary.Path != "src/app.go" || absolutePayload.Target.NormalizedPath != "src/app.go" {
		t.Fatalf("absolute payload path/normalized = %q/%q, want src/app.go/src/app.go", absolutePayload.Summary.Path, absolutePayload.Target.NormalizedPath)
	}
	if absolutePayload.Target.Input != absolutePath {
		t.Fatalf("absolute payload input = %q, want %q", absolutePayload.Target.Input, absolutePath)
	}

	out, errOut, err = executeForTest(t, "file-detail", "src/app.go", "--repo", "fixture", "--json", "--format", "expanded")
	if err != nil {
		t.Fatalf("file-detail expanded JSON returned error: %v\nstdout:\n%s\nstderr:\n%s", err, out, errOut)
	}
	if errOut != "" {
		t.Fatalf("file-detail expanded JSON wrote stderr: %q", errOut)
	}
	var expandedPayload filecontext.FileContext
	if err := json.Unmarshal([]byte(out), &expandedPayload); err != nil {
		t.Fatalf("parse expanded file-detail JSON: %v\n%s", err, out)
	}
	if expandedPayload.Repo != "fixture" || expandedPayload.Summary.Path != "src/app.go" {
		t.Fatalf("expanded payload repo/path = %q/%q", expandedPayload.Repo, expandedPayload.Summary.Path)
	}
	if expandedPayload.Limits.RelationshipSamplesPerGroup != 5 {
		t.Fatalf("expanded default relationship limit = %d, want 5", expandedPayload.Limits.RelationshipSamplesPerGroup)
	}
	if len(expandedPayload.SymbolTree) != 2 || expandedPayload.SymbolTree[0].Name != "Server" || expandedPayload.SymbolTree[1].Name != "NewServer" {
		t.Fatalf("expanded symbol tree = %#v", expandedPayload.SymbolTree)
	}

	out, errOut, err = executeForTest(t, "file-detail", "src/app.go", "--repo", "fixture", "--json", "--relationships", "1")
	if err != nil {
		t.Fatalf("file-detail compact limited JSON returned error: %v\nstdout:\n%s\nstderr:\n%s", err, out, errOut)
	}
	var limitedPayload filecontext.CompactFileContext
	if err := json.Unmarshal([]byte(out), &limitedPayload); err != nil {
		t.Fatalf("parse limited compact file-detail JSON: %v\n%s", err, out)
	}
	if limitedPayload.Limits.RelationshipSamplesPerGroup != 1 {
		t.Fatalf("limited relationship limit = %d, want 1", limitedPayload.Limits.RelationshipSamplesPerGroup)
	}
	limitedRows := limitedPayload.Tables.Relationships.OutboundByFile[0].Rows
	if limitedRows.Total != 2 || limitedRows.Returned != 1 || limitedRows.Omitted != 1 {
		t.Fatalf("limited compact outbound rows = %#v, want total/returned/omitted 2/1/1", limitedRows)
	}

	out, errOut, err = executeForTest(t, "file-detail", "src/app.go", "--repo", "fixture")
	if err != nil {
		t.Fatalf("file-detail human returned error: %v\nstdout:\n%s\nstderr:\n%s", err, out, errOut)
	}
	if errOut != "" {
		t.Fatalf("file-detail human wrote stderr: %q", errOut)
	}
	for _, want := range []string{"File: src/app.go", "Symbols: 2 exported=1", "Outbound files:", "src/store.go", "Unresolved source sites: total=1"} {
		if !strings.Contains(out, want) {
			t.Fatalf("file-detail human output missing %q:\n%s", want, out)
		}
	}
}

func TestFileDetailCommandRejectsUnsupportedFormat(t *testing.T) {
	home := t.TempDir()
	t.Setenv(repo.HomeEnvName, home)
	repoPath := t.TempDir()
	registerDirectToolCommandRepo(t, repo.NewStore(home), repoPath, "fixture")
	writeFileProjectionCommandGraph(t, repoPath)

	out, errOut, err := executeForTest(t, "file-detail", "src/app.go", "--repo", "fixture", "--json", "--format", "summary")
	if err == nil {
		t.Fatalf("file-detail accepted unsupported format\nstdout:\n%s\nstderr:\n%s", out, errOut)
	}
	if !strings.Contains(err.Error(), `unsupported file-detail format "summary"`) {
		t.Fatalf("unsupported format error = %v\nstdout:\n%s\nstderr:\n%s", err, out, errOut)
	}

	out, errOut, err = executeForTest(t, "file-detail", "src/app.go", "--repo", "fixture", "--format", "compact")
	if err == nil {
		t.Fatalf("file-detail accepted --format without --json\nstdout:\n%s\nstderr:\n%s", out, errOut)
	}
	if !strings.Contains(err.Error(), "--format requires --json") {
		t.Fatalf("format without json error = %v\nstdout:\n%s\nstderr:\n%s", err, out, errOut)
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
		payload.Files[0].Unresolved != 1 ||
		payload.Files[0].Risk != "medium" {
		t.Fatalf("top hotspot = %#v, want src/app.go unresolved=1 risk=medium", payload.Files[0])
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

	out, errOut, err = executeForTest(t, "file-hotspots", "--repo", "fixture", "--sort", "bad-sort", "--json")
	if err == nil {
		t.Fatalf("file-hotspots accepted unsupported sort\nstdout:\n%s\nstderr:\n%s", out, errOut)
	}
	if !strings.Contains(err.Error(), `unsupported sort "bad-sort"`) {
		t.Fatalf("retired sort error = %v\nstdout:\n%s\nstderr:\n%s", err, out, errOut)
	}
}

func TestFileDetailCommandReportsMissingFile(t *testing.T) {
	home := t.TempDir()
	t.Setenv(repo.HomeEnvName, home)
	repoPath := t.TempDir()
	registerDirectToolCommandRepo(t, repo.NewStore(home), repoPath, "fixture")
	writeFileProjectionCommandGraph(t, repoPath)

	out, errOut, err := executeForTest(t, "file-detail", "src/missing.go", "--repo", "fixture")
	if err == nil {
		t.Fatalf("file-detail missing file succeeded\nstdout:\n%s\nstderr:\n%s", out, errOut)
	}
	if !strings.Contains(err.Error(), `file "src/missing.go" not found`) {
		t.Fatalf("missing file error = %v", err)
	}
}

func TestFileDetailCommandRejectsOutsideRepoAbsolutePath(t *testing.T) {
	home := t.TempDir()
	t.Setenv(repo.HomeEnvName, home)
	repoPath := t.TempDir()
	registerDirectToolCommandRepo(t, repo.NewStore(home), repoPath, "fixture")
	writeFileProjectionCommandGraph(t, repoPath)

	out, errOut, err := executeForTest(t, "file-detail", filepath.Join(filepath.Dir(repoPath), "outside", "src", "app.go"), "--repo", "fixture")
	if err == nil {
		t.Fatalf("file-detail outside absolute path succeeded\nstdout:\n%s\nstderr:\n%s", out, errOut)
	}
	if !strings.Contains(err.Error(), filecontext.ErrFilePathOutsideRepo.Error()) {
		t.Fatalf("outside repo error = %v", err)
	}
}

func TestFileContextCommandIsNotRegistered(t *testing.T) {
	out, errOut, err := executeForTest(t, "file-context", "--help")
	if err == nil {
		t.Fatalf("file-context alias unexpectedly succeeded\nstdout:\n%s\nstderr:\n%s", out, errOut)
	}
	if !strings.Contains(err.Error(), `unknown command "file-context"`) {
		t.Fatalf("file-context error = %v\nstdout:\n%s\nstderr:\n%s", err, out, errOut)
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
		ID: "rel:call-server-save", SourceID: "Struct:src/app.go:Server", TargetID: "Function:src/store.go:Save", Type: graph.RelCalls,
		FilePath: "src/app.go", SourceSiteID: "SourceSite:src/app.go#call#Save#7#2#7#8", SourceSiteStatus: "resolved", ProofKind: "scope-binding",
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
