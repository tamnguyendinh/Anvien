package filecontext

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
	"github.com/tamnguyendinh/anvien/internal/semantic"
)

func TestBuildFileContextDerivesTreeRelationshipsAndUnresolved(t *testing.T) {
	builder := NewBuilder(fileContextFixture(false))

	context, ok := builder.BuildFileContext(`.\src\app.go`, Options{RelationshipSamplesPerGroup: 1})
	if !ok {
		t.Fatalf("BuildFileContext() did not find file")
	}

	if context.Summary.Path != "src/app.go" {
		t.Fatalf("summary path = %q, want src/app.go", context.Summary.Path)
	}
	if context.Summary.FileRole != string(semantic.FileRoleUnknown) {
		t.Fatalf("summary fileRole = %q, want unknown", context.Summary.FileRole)
	}
	if context.Summary.SymbolCount != 3 {
		t.Fatalf("symbol count = %d, want 3", context.Summary.SymbolCount)
	}
	if context.Summary.ExportedSymbolCount != 1 {
		t.Fatalf("exported count = %d, want 1", context.Summary.ExportedSymbolCount)
	}
	if context.Summary.LocalRelationshipCount != 1 || context.Summary.OutboundRefCount != 2 || context.Summary.InboundRefCount != 1 {
		t.Fatalf("relationship summary = local %d outbound %d inbound %d, want 1/2/1",
			context.Summary.LocalRelationshipCount,
			context.Summary.OutboundRefCount,
			context.Summary.InboundRefCount,
		)
	}
	if context.Summary.Unresolved != 1 || context.Quality.UnresolvedCalls != 1 {
		t.Fatalf("unresolved summary = %d calls=%d, want 1/1", context.Summary.Unresolved, context.Quality.UnresolvedCalls)
	}

	if len(context.SymbolTree) != 2 {
		t.Fatalf("symbol tree roots = %d, want 2", len(context.SymbolTree))
	}
	if context.SymbolTree[0].ID != "Struct:src/app.go:Server" {
		t.Fatalf("first root = %q, want Server", context.SymbolTree[0].ID)
	}
	if len(context.SymbolTree[0].Children) != 1 || context.SymbolTree[0].Children[0].ID != "Method:src/app.go:Server.Start" {
		t.Fatalf("server children = %#v, want Start", context.SymbolTree[0].Children)
	}
	if context.SymbolTree[1].RelationshipCounts.Outbound != 2 || context.SymbolTree[1].RelationshipCounts.Inbound != 1 || context.SymbolTree[1].RelationshipCounts.Unresolved != 1 {
		t.Fatalf("NewServer counts = %#v, want outbound=2 inbound=1 unresolved=1", context.SymbolTree[1].RelationshipCounts)
	}

	if context.Relationships.Local.Total != 1 || len(context.Relationships.Local.Samples) != 1 {
		t.Fatalf("local relationships = total %d samples %d, want 1/1", context.Relationships.Local.Total, len(context.Relationships.Local.Samples))
	}
	if len(context.Relationships.OutboundByFile) != 1 {
		t.Fatalf("outbound groups = %d, want 1", len(context.Relationships.OutboundByFile))
	}
	outbound := context.Relationships.OutboundByFile[0]
	if outbound.File != "src/store.go" || outbound.Total != 2 || len(outbound.Samples) != 1 {
		t.Fatalf("outbound group = %#v, want src/store.go total 2 sample limit 1", outbound)
	}
	if len(context.Relationships.InboundByFile) != 1 || context.Relationships.InboundByFile[0].File != "src/app_test.go" {
		t.Fatalf("inbound groups = %#v, want src/app_test.go", context.Relationships.InboundByFile)
	}

	if context.Unresolved.Total != 1 || context.Unresolved.ByKind["unresolved_call"] != 1 {
		t.Fatalf("unresolved = %#v, want one unresolved_call", context.Unresolved)
	}
	if len(context.Unresolved.Groups) != 1 || context.Unresolved.Groups[0].SourceSymbol != "Function:src/app.go:NewServer" {
		t.Fatalf("unresolved groups = %#v, want NewServer group", context.Unresolved.Groups)
	}
	if context.Linked.Counts.Flows != 1 || context.Linked.Counts.Routes != 1 || context.Linked.Counts.MCPTools != 1 || context.Linked.Counts.Tests != 1 {
		t.Fatalf("linked counts = %#v, want 1 flow/route/tool/test", context.Linked.Counts)
	}
	if context.Summary.LinkedFlowCount != 1 || context.Summary.LinkedTestCount != 1 {
		t.Fatalf("summary linked counts = flows %d tests %d, want 1/1", context.Summary.LinkedFlowCount, context.Summary.LinkedTestCount)
	}
	if len(context.Linked.Flows) != 1 || context.Linked.Flows[0].Name != "MCP initialize" {
		t.Fatalf("linked flows = %#v, want MCP initialize", context.Linked.Flows)
	}
	if len(context.Linked.Routes) != 1 || context.Linked.Routes[0].Name != "GET /api/app" {
		t.Fatalf("linked routes = %#v, want GET /api/app", context.Linked.Routes)
	}
	if len(context.Linked.MCPTools) != 1 || context.Linked.MCPTools[0].Name != "context" {
		t.Fatalf("linked tools = %#v, want context", context.Linked.MCPTools)
	}
	if len(context.Linked.Tests) != 1 || context.Linked.Tests[0].Name != "src/app_test.go" {
		t.Fatalf("linked tests = %#v, want src/app_test.go", context.Linked.Tests)
	}
}

func TestCompactFileContextFromExpandedInternsRowsAndPreservesCounts(t *testing.T) {
	builder := NewBuilder(fileContextFixture(false))
	context, ok := builder.BuildFileContext("src/app.go", Options{RelationshipSamplesPerGroup: 1})
	if !ok {
		t.Fatalf("BuildFileContext() did not find file")
	}
	AttachMetadata(&context, "fixture", "/repo", GraphInfo{Path: "graph.json", Stale: true})

	compact := CompactFileContextFromExpanded(context)
	if compact.Format != CompactFileContextFormat || compact.Version != CompactFileContextVersion {
		t.Fatalf("compact identity = %s/%d", compact.Format, compact.Version)
	}
	if compact.Repo != "fixture" || compact.RepoPath != "/repo" || !compact.Graph.Stale {
		t.Fatalf("compact metadata = repo %q repoPath %q graph %#v", compact.Repo, compact.RepoPath, compact.Graph)
	}
	if len(compact.Schema.RelationshipRow) == 0 || len(compact.Schema.SymbolRow) == 0 || len(compact.Schema.RangeTuple) != 4 {
		t.Fatalf("compact schema incomplete: %#v", compact.Schema)
	}
	if compact.Summary.Path != context.Summary.Path || compact.Quality != context.Quality || compact.Limits != context.Limits {
		t.Fatalf("compact top-level facts differ from expanded context")
	}

	serverRef := compactDictIndex(t, compact.Dict.Symbols, "Struct:src/app.go:Server")
	startRef := compactDictIndex(t, compact.Dict.Symbols, "Method:src/app.go:Server.Start")
	newServerRef := compactDictIndex(t, compact.Dict.Symbols, "Function:src/app.go:NewServer")
	if len(compact.Tables.Symbols) != 3 {
		t.Fatalf("compact symbol rows = %d, want 3", len(compact.Tables.Symbols))
	}
	if compactRowInt(t, compact.Tables.Symbols[0][0]) != serverRef {
		t.Fatalf("first symbol row = %#v, want server ref %d", compact.Tables.Symbols[0], serverRef)
	}
	if compactRowInt(t, compact.Tables.Symbols[1][0]) != startRef || compactRowInt(t, compact.Tables.Symbols[1][1]) != serverRef {
		t.Fatalf("child symbol row = %#v, want start ref parent server ref", compact.Tables.Symbols[1])
	}
	if compactRowInt(t, compact.Tables.Symbols[2][0]) != newServerRef {
		t.Fatalf("third symbol row = %#v, want NewServer ref", compact.Tables.Symbols[2])
	}

	if compact.Tables.Relationships.Counts != context.Relationships.Counts {
		t.Fatalf("relationship counts = %#v, want %#v", compact.Tables.Relationships.Counts, context.Relationships.Counts)
	}
	if compact.Tables.Relationships.Local.Rows.Total != 1 || compact.Tables.Relationships.Local.Rows.Returned != 1 || compact.Tables.Relationships.Local.Rows.Omitted != 0 {
		t.Fatalf("local relationship rows = %#v, want total/returned/omitted 1/1/0", compact.Tables.Relationships.Local.Rows)
	}
	if len(compact.Tables.Relationships.OutboundByFile) != 1 {
		t.Fatalf("outbound groups = %#v, want one", compact.Tables.Relationships.OutboundByFile)
	}
	outbound := compact.Tables.Relationships.OutboundByFile[0]
	if compact.Dict.Files[outbound.File] != "src/store.go" {
		t.Fatalf("outbound file ref = %d -> %q, want src/store.go", outbound.File, compact.Dict.Files[outbound.File])
	}
	if outbound.Total != 2 || outbound.Rows.Returned != 1 || outbound.Rows.Omitted != 1 {
		t.Fatalf("outbound row limit metadata = %#v total=%d, want total 2 returned 1 omitted 1", outbound.Rows, outbound.Total)
	}
	outboundRow := outbound.Rows.Items[0]
	if compact.Dict.Files[compactRowInt(t, outboundRow[0])] != "src/app.go" || compact.Dict.Files[compactRowInt(t, outboundRow[4])] != "src/store.go" {
		t.Fatalf("outbound row file refs = %#v files=%#v", outboundRow, compact.Dict.Files)
	}
	if compact.Dict.SourceSites[compactRowInt(t, outboundRow[7])] != "site:new-save" {
		t.Fatalf("outbound source site = %#v sites=%#v", outboundRow[7], compact.Dict.SourceSites)
	}

	if compact.Tables.Unresolved.Total != context.Unresolved.Total || len(compact.Tables.Unresolved.Groups) != 1 {
		t.Fatalf("unresolved compact = %#v, want total %d one group", compact.Tables.Unresolved, context.Unresolved.Total)
	}
	unresolved := compact.Tables.Unresolved.Groups[0]
	if compactRowInt(t, unresolved.SourceSymbol) != newServerRef || unresolved.Rows.Total != 1 || unresolved.Rows.Returned != 1 || unresolved.Rows.Omitted != 0 {
		t.Fatalf("unresolved group = %#v, want NewServer and 1/1/0 rows", unresolved)
	}

	if compact.Tables.Linked.Counts != context.Linked.Counts {
		t.Fatalf("linked counts = %#v, want %#v", compact.Tables.Linked.Counts, context.Linked.Counts)
	}
	if compact.Tables.Linked.Flows.Total != 1 || compact.Tables.Linked.Routes.Total != 1 || compact.Tables.Linked.MCPTools.Total != 1 || compact.Tables.Linked.Tests.Total != 1 {
		t.Fatalf("linked row totals = flows %#v routes %#v tools %#v tests %#v",
			compact.Tables.Linked.Flows,
			compact.Tables.Linked.Routes,
			compact.Tables.Linked.MCPTools,
			compact.Tables.Linked.Tests,
		)
	}

	if _, err := json.Marshal(compact); err != nil {
		t.Fatalf("marshal compact context: %v", err)
	}
}

func TestBuildFileContextReturnsFalseForMissingFile(t *testing.T) {
	_, ok := NewBuilder(fileContextFixture(false)).BuildFileContext("src/missing.go", Options{})
	if ok {
		t.Fatalf("BuildFileContext() found missing file")
	}
}

func compactDictIndex(t *testing.T, values []string, want string) int {
	t.Helper()
	for index, value := range values {
		if value == want {
			return index
		}
	}
	t.Fatalf("dictionary missing %q in %#v", want, values)
	return -1
}

func compactRowInt(t *testing.T, value any) int {
	t.Helper()
	got, ok := value.(int)
	if !ok {
		t.Fatalf("compact row value %T(%#v), want int", value, value)
	}
	return got
}

func TestNormalizeRepoFilePath(t *testing.T) {
	tests := []struct {
		name      string
		inputPath string
		repoRoot  string
		want      string
		wantErr   bool
	}{
		{
			name:      "repo relative",
			inputPath: "src/app.go",
			repoRoot:  `E:\Anvien`,
			want:      "src/app.go",
		},
		{
			name:      "dot relative windows separators",
			inputPath: `.\src\app.go`,
			repoRoot:  `E:\Anvien`,
			want:      "src/app.go",
		},
		{
			name:      "windows absolute inside repo",
			inputPath: `E:\Anvien\src\app.go`,
			repoRoot:  `E:\Anvien`,
			want:      "src/app.go",
		},
		{
			name:      "slash absolute inside repo",
			inputPath: "E:/Anvien/src/app.go",
			repoRoot:  `E:\Anvien`,
			want:      "src/app.go",
		},
		{
			name:      "windows absolute uses case insensitive root",
			inputPath: `e:\anvien\src\app.go`,
			repoRoot:  `E:\Anvien`,
			want:      "src/app.go",
		},
		{
			name:      "absolute outside repo",
			inputPath: `E:\Other\src\app.go`,
			repoRoot:  `E:\Anvien`,
			wantErr:   true,
		},
		{
			name:      "absolute sibling prefix is outside repo",
			inputPath: `E:\AnvienOther\src\app.go`,
			repoRoot:  `E:\Anvien`,
			wantErr:   true,
		},
		{
			name:      "blank path",
			inputPath: "  ",
			repoRoot:  `E:\Anvien`,
			want:      "",
		},
		{
			name:      "absolute path without repo root preserves lookup",
			inputPath: `E:\Anvien\src\app.go`,
			repoRoot:  "",
			want:      "E:/Anvien/src/app.go",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := NormalizeRepoFilePath(test.inputPath, test.repoRoot)
			if test.wantErr {
				if !errors.Is(err, ErrFilePathOutsideRepo) {
					t.Fatalf("NormalizeRepoFilePath() err = %v, want ErrFilePathOutsideRepo", err)
				}
				return
			}
			if err != nil {
				t.Fatalf("NormalizeRepoFilePath() err = %v", err)
			}
			if got != test.want {
				t.Fatalf("NormalizeRepoFilePath() = %q, want %q", got, test.want)
			}
		})
	}
}

func TestFileListSortSupportRejectsUnsupportedSorts(t *testing.T) {
	if IsSupportedFileListSort("bad-sort") {
		t.Fatalf("unsupported sort must not be supported")
	}
	if NormalizeFileListSort("bad-sort") != "path" {
		t.Fatalf("unsupported sort normalization should fall back to path")
	}
}

func TestBuildFileContextIsDeterministicAcrossRelationshipOrder(t *testing.T) {
	first, ok := NewBuilder(fileContextFixture(false)).BuildFileContext("src/app.go", Options{})
	if !ok {
		t.Fatalf("first context missing")
	}
	second, ok := NewBuilder(fileContextFixture(true)).BuildFileContext("src/app.go", Options{})
	if !ok {
		t.Fatalf("second context missing")
	}

	firstJSON, err := json.Marshal(first)
	if err != nil {
		t.Fatalf("marshal first: %v", err)
	}
	secondJSON, err := json.Marshal(second)
	if err != nil {
		t.Fatalf("marshal second: %v", err)
	}
	if string(firstJSON) != string(secondJSON) {
		t.Fatalf("contexts differ\nfirst:  %s\nsecond: %s", firstJSON, secondJSON)
	}
}

func TestBuildFileListSortsFiltersAndPaginates(t *testing.T) {
	list := NewBuilder(fileContextFixture(false)).BuildFileList(FileListOptions{
		Sort:  "fan-out",
		Limit: 2,
	})
	if list.Total != 3 || len(list.Files) != 2 {
		t.Fatalf("list size = total %d files %d, want 3/2", list.Total, len(list.Files))
	}
	if list.Sort != "fan-out" {
		t.Fatalf("sort = %q, want fan-out", list.Sort)
	}
	if list.Files[0].Path != "src/app.go" || list.Files[0].OutboundRefCount != 2 {
		t.Fatalf("top fan-out file = %#v, want src/app.go outbound 2", list.Files[0])
	}
	if list.Files[0].LinkedFlowCount != 1 || list.Files[0].LinkedTestCount != 1 {
		t.Fatalf("top fan-out linked counts = %#v, want flow/test 1/1", list.Files[0])
	}

	tests := NewBuilder(fileContextFixture(false)).BuildFileList(FileListOptions{
		Kinds: []string{"test"},
	})
	if tests.Total != 1 || tests.Files[0].Path != "src/app_test.go" {
		t.Fatalf("test filter = %#v, want src/app_test.go", tests)
	}

	unresolved := NewBuilder(fileContextFixture(false)).BuildFileList(FileListOptions{
		Sort:           "unresolved",
		UnresolvedOnly: true,
	})
	if unresolved.Total != 1 || unresolved.Files[0].Path != "src/app.go" || unresolved.Files[0].Unresolved != 1 {
		t.Fatalf("unresolved filter = %#v, want only src/app.go", unresolved)
	}

	secondPage := NewBuilder(fileContextFixture(false)).BuildFileList(FileListOptions{
		Sort:   "path",
		Offset: 1,
		Limit:  1,
	})
	if secondPage.Total != 3 || len(secondPage.Files) != 1 || secondPage.Files[0].Path != "src/app_test.go" {
		t.Fatalf("second page = %#v, want src/app_test.go", secondPage)
	}
}

func TestBuildFileListUsesAppLayerAsTestClassificationTruth(t *testing.T) {
	g := graph.New()
	for _, node := range []graph.Node{
		fileNode("File:internal/analyze/analyze.go", "internal/analyze/analyze.go", "go", "backend", "analyzer"),
		fileNode("File:internal/mcp/server_test.go", "internal/mcp/server_test.go", "go", "api_test", "mcp"),
		fileNode("File:internal/resolution/resolution_test.go", "internal/resolution/resolution_test.go", "go", "backend_test", "resolution"),
		fileNode("File:anvien-web/e2e/graph.spec.ts", "anvien-web/e2e/graph.spec.ts", "typescript", "frontend_test", "web_graph_ui"),
	} {
		g.AddNode(node)
	}

	list := NewBuilder(g).BuildFileList(FileListOptions{Sort: "path", Limit: 0})
	byPath := map[string]FileSummary{}
	for _, file := range list.Files {
		byPath[file.Path] = file
	}

	for _, path := range []string{
		"anvien-web/e2e/graph.spec.ts",
		"internal/mcp/server_test.go",
		"internal/resolution/resolution_test.go",
	} {
		if byPath[path].Kind != "test" {
			t.Fatalf("summary %s kind = %q appLayer=%q, want test", path, byPath[path].Kind, byPath[path].AppLayer)
		}
	}
	if byPath["internal/analyze/analyze.go"].Kind != "source" {
		t.Fatalf("source summary kind = %q, want source", byPath["internal/analyze/analyze.go"].Kind)
	}

	tests := NewBuilder(g).BuildFileList(FileListOptions{Kinds: []string{"test"}, Sort: "path", Limit: 0})
	if tests.Total != 3 || len(tests.Files) != 3 {
		t.Fatalf("test file filter total/files = %d/%d, want 3/3: %#v", tests.Total, len(tests.Files), tests.Files)
	}
	for _, file := range tests.Files {
		if file.Kind != "test" || file.AppLayer == "" {
			t.Fatalf("test file summary = %#v, want kind=test with appLayer", file)
		}
	}
}

func TestBuildFileListClassifiesRawOnlySupportFileRoles(t *testing.T) {
	g := graph.New()
	fixtures := []struct {
		path           string
		appLayer       string
		functionalArea string
		want           semantic.FileRole
		rawSites       int
	}{
		{"internal/frameworks/frameworks.go", "backend", "analyzer", semantic.FileRoleAnalyzerHelper, 209},
		{"internal/scopeir/sort_keys.go", "backend", "providers", semantic.FileRoleHelper, 63},
		{"internal/group/types.go", "backend", "query", semantic.FileRoleContractModel, 16},
		{"internal/repo/paths.go", "backend", "storage", semantic.FileRoleStorageHelper, 13},
		{"internal/testutil/path.go", "backend", "unknown", semantic.FileRoleTestHelper, 12},
		{"internal/repo/settings.go", "backend", "storage", semantic.FileRoleConfig, 11},
		{"internal/repo/runtime_config.go", "backend", "storage", semantic.FileRoleConfig, 10},
		{"internal/cobol/copy_expander.go", "backend", "analyzer", semantic.FileRoleAnalyzerHelper, 9},
		{"internal/parser/metrics.go", "backend", "providers", semantic.FileRoleParserModel, 8},
		{"internal/session/error.go", "backend", "session", semantic.FileRoleRuntimeModel, 6},
		{"internal/resolution/source_site.go", "backend", "resolution", semantic.FileRoleHelper, 4},
		{"internal/scopeir/facts.go", "backend", "providers", semantic.FileRoleParserModel, 4},
		{"internal/scopeir/range.go", "backend", "providers", semantic.FileRoleParserModel, 4},
		{"internal/session/types.go", "backend", "session", semantic.FileRoleRuntimeModel, 3},
		{"internal/cli/exit_error.go", "backend", "cli", semantic.FileRoleHelper, 2},
		{"internal/lbugnative/runner.go", "backend", "storage", semantic.FileRoleAdapter, 1},
		{"internal/lbugnative/runner_default.go", "backend", "storage", semantic.FileRoleFallbackAdapter, 1},
	}
	for _, fixture := range fixtures {
		g.AddNode(fileNode("File:"+fixture.path, fixture.path, "go", fixture.appLayer, fixture.functionalArea))
		for i := 0; i < fixture.rawSites; i++ {
			g.AddNode(resolutionGap(
				"ResolutionGap:"+fixture.path+":raw:"+itoa(i),
				fixture.path,
				"File:"+fixture.path,
				"builtin",
				"unresolved_call",
				"builtin",
				"non_actionable",
				i+1,
			))
		}
	}

	list := NewBuilder(g).BuildFileList(FileListOptions{Sort: "path", Limit: 0})
	byPath := map[string]FileSummary{}
	for _, file := range list.Files {
		byPath[file.Path] = file
	}
	for _, fixture := range fixtures {
		summary, ok := byPath[fixture.path]
		if !ok {
			t.Fatalf("missing summary for %s", fixture.path)
		}
		if summary.FileRole != string(fixture.want) {
			t.Fatalf("%s fileRole = %q, want %q", fixture.path, summary.FileRole, fixture.want)
		}
		if summary.FileGroup != string(semantic.FileGroupBackendSupportModelHelper) {
			t.Fatalf("%s fileGroup = %q, want %q", fixture.path, summary.FileGroup, semantic.FileGroupBackendSupportModelHelper)
		}
	}

	detail, ok := NewBuilder(g).BuildFileContext("internal/repo/runtime_config.go", Options{})
	if !ok {
		t.Fatalf("BuildFileContext() did not find runtime_config.go")
	}
	if detail.Summary.FileRole != string(semantic.FileRoleConfig) {
		t.Fatalf("detail fileRole = %q, want %q", detail.Summary.FileRole, semantic.FileRoleConfig)
	}
	if detail.Summary.FileGroup != string(semantic.FileGroupBackendSupportModelHelper) {
		t.Fatalf("detail fileGroup = %q, want %q", detail.Summary.FileGroup, semantic.FileGroupBackendSupportModelHelper)
	}

	if len(list.FileGroups) != 1 {
		t.Fatalf("file groups = %#v, want one backend support group", list.FileGroups)
	}
	group := list.FileGroups[0]
	if group.Key != string(semantic.FileGroupBackendSupportModelHelper) || group.Label != "Backend support/model/helper files" {
		t.Fatalf("file group identity = %#v, want backend support/model/helper", group)
	}
	if group.Files != 17 || group.Unresolved != 376 {
		t.Fatalf("file group counts = files %d unresolved %d, want 17/376", group.Files, group.Unresolved)
	}
	wantRoles := map[string]int{
		string(semantic.FileRoleAnalyzerHelper):  2,
		string(semantic.FileRoleHelper):          3,
		string(semantic.FileRoleContractModel):   1,
		string(semantic.FileRoleStorageHelper):   1,
		string(semantic.FileRoleTestHelper):      1,
		string(semantic.FileRoleConfig):          2,
		string(semantic.FileRoleParserModel):     3,
		string(semantic.FileRoleRuntimeModel):    2,
		string(semantic.FileRoleAdapter):         1,
		string(semantic.FileRoleFallbackAdapter): 1,
	}
	for role, want := range wantRoles {
		if group.Roles[role] != want {
			t.Fatalf("file group role %s = %d, want %d in %#v", role, group.Roles[role], want, group.Roles)
		}
	}
}

func TestBuildFileSummariesUseCanonicalUnresolved(t *testing.T) {
	g := graph.New()
	for _, node := range []graph.Node{
		fileNode("File:src/app.go", "src/app.go", "go", "backend", "mcp"),
		fileNode("File:src/app_test.go", "src/app_test.go", "go", "backend_test", "mcp"),
		symbolNode("Function:src/app.go:Run", scopeir.NodeFunction, "Run", "src/app.go", 2, 1, 10, 1, ""),
		symbolNode("Function:src/app_test.go:TestRun", scopeir.NodeFunction, "TestRun", "src/app_test.go", 2, 1, 8, 1, ""),
		resolutionGap("ResolutionGap:src/app.go:call", "src/app.go", "Function:src/app.go:Run", "missingCall", "unresolved_call", "in_repo_unresolved", "analyzer_gap", 4),
		resolutionGap("ResolutionGap:src/app.go:builtin", "src/app.go", "Function:src/app.go:Run", "println", "unresolved_call", "builtin", "non_actionable", 5),
		resolutionGap("ResolutionGap:src/app.go:unknown", "src/app.go", "Function:src/app.go:Run", "mystery", "unresolved_reference", "", "", 6),
	} {
		g.AddNode(node)
	}
	g.AddRelationship(defines("File:src/app.go", "Function:src/app.go:Run"))
	g.AddRelationship(defines("File:src/app_test.go", "Function:src/app_test.go:TestRun"))

	builder := NewBuilder(g)
	list := builder.BuildFileList(FileListOptions{Sort: "path", Limit: 0})
	byPath := map[string]FileSummary{}
	for _, file := range list.Files {
		byPath[file.Path] = file
	}

	source := byPath["src/app.go"]
	if source.Unresolved != 3 {
		t.Fatalf("source unresolved = %d, want 3", source.Unresolved)
	}
	if source.Risk != "medium" {
		t.Fatalf("source risk = %q, want medium", source.Risk)
	}

	testFile := byPath["src/app_test.go"]
	if testFile.Kind != "test" || testFile.Unresolved != 0 {
		t.Fatalf("test file summary = %#v, want kind=test unresolved=0", testFile)
	}
	if testFile.Risk != "low" {
		t.Fatalf("test risk = %q, want low", testFile.Risk)
	}

	unresolved := builder.BuildFileList(FileListOptions{Sort: "unresolved", UnresolvedOnly: true, Limit: 0})
	if unresolved.Total != 1 || unresolved.Files[0].Path != "src/app.go" {
		t.Fatalf("unresolved filter = %#v, want only source file", unresolved)
	}

	sourceContext, ok := builder.BuildFileContext("src/app.go", Options{})
	if !ok {
		t.Fatalf("BuildFileContext() did not find source file")
	}
	if sourceContext.Summary.Unresolved != source.Unresolved {
		t.Fatalf("source detail unresolved = %#v, want list summary %#v", sourceContext.Summary, source)
	}

	testContext, ok := builder.BuildFileContext("src/app_test.go", Options{})
	if !ok {
		t.Fatalf("BuildFileContext() did not find test file")
	}
	if testContext.Summary.Unresolved != testFile.Unresolved {
		t.Fatalf("test detail unresolved = %#v, want list summary %#v", testContext.Summary, testFile)
	}
}

func TestBuildFileListChangedFileFilter(t *testing.T) {
	builder := NewBuilder(fileContextFixture(false))

	changed := builder.BuildFileList(FileListOptions{
		ChangedOnly: true,
		ChangedPaths: map[string]struct{}{
			"src/app.go": {},
		},
		Stale: true,
	})
	if changed.Total != 1 || len(changed.Files) != 1 || changed.Files[0].Path != "src/app.go" {
		t.Fatalf("changed filter = %#v, want only src/app.go", changed)
	}
	if !changed.Files[0].ChangedSinceAnalyze || !changed.Files[0].Stale {
		t.Fatalf("changed quality = stale %v changed %v, want both true", changed.Files[0].Stale, changed.Files[0].ChangedSinceAnalyze)
	}

	all := builder.BuildFileList(FileListOptions{
		ChangedPaths: map[string]struct{}{
			"src/app.go": {},
		},
	})
	if all.Total != 3 {
		t.Fatalf("all changed metadata total = %d, want 3", all.Total)
	}
	changedCount := 0
	for _, file := range all.Files {
		if file.ChangedSinceAnalyze {
			changedCount++
		}
	}
	if changedCount != 1 {
		t.Fatalf("changed metadata count = %d, want 1", changedCount)
	}
}

func TestBuildFileListHighFanFilters(t *testing.T) {
	builder := NewBuilder(fileContextFixture(false))

	highFanIn := builder.BuildFileList(FileListOptions{
		HighFanInOnly:      true,
		HighFanInThreshold: 1,
	})
	if highFanIn.Total != 2 {
		t.Fatalf("high fan-in total = %d, want 2", highFanIn.Total)
	}
	if highFanIn.Files[0].Path != "src/app.go" && highFanIn.Files[1].Path != "src/app.go" {
		t.Fatalf("high fan-in files = %#v, want src/app.go included", highFanIn.Files)
	}

	highFanOut := builder.BuildFileList(FileListOptions{
		HighFanOutOnly:      true,
		HighFanOutThreshold: 2,
	})
	if highFanOut.Total != 1 || highFanOut.Files[0].Path != "src/app.go" {
		t.Fatalf("high fan-out = %#v, want src/app.go", highFanOut)
	}
}

func TestBuildFileContextQualitySignalsAndUnresolvedBuckets(t *testing.T) {
	builder := NewBuilder(qualitySignalFixture())

	context, ok := builder.BuildFileContext("src/app.go", Options{})
	if !ok {
		t.Fatalf("BuildFileContext() did not find source file")
	}
	AttachMetadata(&context, "fixture", "/repo", GraphInfo{Path: "graph.json", Stale: true})

	if !context.Quality.Stale || !context.Quality.ChangedSinceAnalyze {
		t.Fatalf("quality stale fields = stale %v changed %v, want both true", context.Quality.Stale, context.Quality.ChangedSinceAnalyze)
	}
	if context.Quality.UnresolvedCalls != 1 || context.Quality.UnresolvedImports != 1 || context.Quality.UnresolvedRefs != 1 {
		t.Fatalf("quality unresolved counts = calls %d imports %d refs %d, want 1/1/1",
			context.Quality.UnresolvedCalls,
			context.Quality.UnresolvedImports,
			context.Quality.UnresolvedRefs,
		)
	}
	if context.Unresolved.ByClassification["external_library"] != 1 || context.Unresolved.ByClassification["in_repo_unresolved"] != 2 {
		t.Fatalf("classification counts = %#v, want external=1 in_repo=2", context.Unresolved.ByClassification)
	}
	if context.Unresolved.ByActionability["review"] != 1 || context.Unresolved.ByActionability["analyzer_gap"] != 2 {
		t.Fatalf("actionability counts = %#v, want review=1 analyzer_gap=2", context.Unresolved.ByActionability)
	}

	generated, ok := builder.BuildFileContext("gen/client.ts", Options{})
	if !ok {
		t.Fatalf("BuildFileContext() did not find generated file")
	}
	if !generated.Quality.Generated || generated.Summary.Kind != "generated" {
		t.Fatalf("generated quality = %#v summary=%#v, want generated kind and flag", generated.Quality, generated.Summary)
	}

	testFile, ok := builder.BuildFileContext("src/app_test.go", Options{})
	if !ok {
		t.Fatalf("BuildFileContext() did not find test file")
	}
	if testFile.Summary.Kind != "test" || testFile.Quality.Generated {
		t.Fatalf("test file quality = %#v summary=%#v, want test kind and not generated", testFile.Quality, testFile.Summary)
	}
}

func TestBuildFileContextLimitsLinkedItemsAndPreservesCounts(t *testing.T) {
	g := fileContextFixture(false)
	g.AddNode(graph.Node{ID: "Process:secondary", Label: scopeir.NodeProcess, Properties: graph.NodeProperties{"name": "Secondary flow"}})
	g.AddRelationship(graph.Relationship{
		ID:         "rel:step:secondary",
		SourceID:   "Function:src/app.go:NewServer",
		TargetID:   "Process:secondary",
		Type:       graph.RelStepInProcess,
		Confidence: 0.7,
	})

	context, ok := NewBuilder(g).BuildFileContext("src/app.go", Options{LinkedSamplesPerKind: 1})
	if !ok {
		t.Fatalf("BuildFileContext() did not find file")
	}
	if context.Linked.Counts.Flows != 2 || len(context.Linked.Flows) != 1 {
		t.Fatalf("linked flow count/samples = %d/%d, want 2/1", context.Linked.Counts.Flows, len(context.Linked.Flows))
	}
	if context.Summary.LinkedFlowCount != 2 {
		t.Fatalf("summary linked flow count = %d, want 2", context.Summary.LinkedFlowCount)
	}
}

func BenchmarkBuildFileListCurrentScale(b *testing.B) {
	builder := NewBuilder(fileListBenchmarkGraph(821, 126000))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list := builder.BuildFileList(FileListOptions{Sort: "fan-out", Limit: 50})
		if len(list.Files) == 0 {
			b.Fatalf("empty file list")
		}
	}
}

func TestBuilderCacheHitsAndInvalidatesOnGraphChange(t *testing.T) {
	cache := NewBuilderCache()
	key := CacheKey{Repo: "Anvien", RepoPath: `E:\Anvien`, GraphPath: `.anvien\graph.json`}
	g := fileContextFixture(false)

	first, hit := cache.Get(key, g)
	if hit {
		t.Fatalf("first cache lookup hit")
	}
	second, hit := cache.Get(key, g)
	if !hit || first != second {
		t.Fatalf("second cache lookup = hit %v same %v, want hit same builder", hit, first == second)
	}

	g.AddNode(fileNode("File:src/new.go", "src/new.go", "go", "backend", "mcp"))
	third, hit := cache.Get(key, g)
	if hit || third == first {
		t.Fatalf("graph-change cache lookup = hit %v same %v, want miss with new builder", hit, third == first)
	}
	if cache.Len() != 2 {
		t.Fatalf("cache length = %d, want 2 fingerprinted builders", cache.Len())
	}

	cache.Invalidate(CacheKey{Repo: "Anvien", RepoPath: `E:\Anvien`, GraphPath: `.anvien\graph.json`})
	if cache.Len() != 0 {
		t.Fatalf("cache length after invalidate = %d, want 0", cache.Len())
	}
}

func TestBuilderCacheIsolatesReposAndExplicitGraphHashes(t *testing.T) {
	cache := NewBuilderCache()
	g := fileContextFixture(false)
	base := CacheKey{RepoPath: "repo", GraphPath: "graph.json", GraphHash: "v1"}

	first, hit := cache.Get(CacheKey{Repo: "one", RepoPath: base.RepoPath, GraphPath: base.GraphPath, GraphHash: base.GraphHash}, g)
	if hit {
		t.Fatalf("first repo lookup hit")
	}
	second, hit := cache.Get(CacheKey{Repo: "one", RepoPath: base.RepoPath, GraphPath: base.GraphPath, GraphHash: base.GraphHash}, g)
	if !hit || second != first {
		t.Fatalf("same repo/hash lookup = hit %v same %v, want hit same", hit, second == first)
	}
	otherRepo, hit := cache.Get(CacheKey{Repo: "two", RepoPath: base.RepoPath, GraphPath: base.GraphPath, GraphHash: base.GraphHash}, g)
	if hit || otherRepo == first {
		t.Fatalf("other repo lookup = hit %v same %v, want miss", hit, otherRepo == first)
	}
	otherHash, hit := cache.Get(CacheKey{Repo: "one", RepoPath: base.RepoPath, GraphPath: base.GraphPath, GraphHash: "v2"}, g)
	if hit || otherHash == first {
		t.Fatalf("other hash lookup = hit %v same %v, want miss", hit, otherHash == first)
	}
}

func BenchmarkBuilderCacheHit(b *testing.B) {
	cache := NewBuilderCache()
	g := fileListBenchmarkGraph(821, 126000)
	key := CacheKey{Repo: "bench", RepoPath: "repo", GraphPath: "graph.json", GraphHash: "fixture-v1"}
	if _, hit := cache.Get(key, g); hit {
		b.Fatalf("first cache lookup hit")
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		builder, hit := cache.Get(key, g)
		if !hit || builder == nil {
			b.Fatalf("cache miss")
		}
	}
}

func BenchmarkBuilderCacheColdBuild(b *testing.B) {
	g := fileListBenchmarkGraph(821, 126000)
	key := CacheKey{Repo: "bench", RepoPath: "repo", GraphPath: "graph.json", GraphHash: "fixture-v1"}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache := NewBuilderCache()
		builder, hit := cache.Get(key, g)
		if hit || builder == nil {
			b.Fatalf("cache cold build = hit %v builder nil %v, want miss with builder", hit, builder == nil)
		}
	}
}

func qualitySignalFixture() *graph.Graph {
	g := graph.New()
	for _, node := range []graph.Node{
		fileNode("File:src/app.go", "src/app.go", "go", "backend", "mcp"),
		fileNode("File:gen/client.ts", "gen/client.ts", "typescript", "generated_contract", "contracts"),
		fileNode("File:src/app_test.go", "src/app_test.go", "go", "backend_test", "mcp"),
		symbolNode("Function:src/app.go:Run", scopeir.NodeFunction, "Run", "src/app.go", 2, 1, 10, 1, ""),
		symbolNode("Function:gen/client.ts:GeneratedClient", scopeir.NodeFunction, "GeneratedClient", "gen/client.ts", 1, 1, 3, 1, ""),
		symbolNode("Function:src/app_test.go:TestRun", scopeir.NodeFunction, "TestRun", "src/app_test.go", 2, 1, 8, 1, ""),
		resolutionGap("ResolutionGap:src/app.go:call", "src/app.go", "Function:src/app.go:Run", "missingCall", "unresolved_call", "in_repo_unresolved", "analyzer_gap", 4),
		resolutionGap("ResolutionGap:src/app.go:import", "src/app.go", "Function:src/app.go:Run", "external/pkg", "unresolved_import", "external_library", "review", 5),
		resolutionGap("ResolutionGap:src/app.go:type", "src/app.go", "Function:src/app.go:Run", "MissingType", "unresolved_type_reference", "in_repo_unresolved", "analyzer_gap", 6),
	} {
		g.AddNode(node)
	}
	g.AddRelationship(defines("File:src/app.go", "Function:src/app.go:Run"))
	g.AddRelationship(defines("File:gen/client.ts", "Function:gen/client.ts:GeneratedClient"))
	g.AddRelationship(defines("File:src/app_test.go", "Function:src/app_test.go:TestRun"))
	return g
}

func fileContextFixture(reverseRelationships bool) *graph.Graph {
	g := graph.New()
	for _, node := range []graph.Node{
		fileNode("File:src/app.go", "src/app.go", "go", "backend", "mcp"),
		fileNode("File:src/store.go", "src/store.go", "go", "backend", "storage"),
		fileNode("File:src/app_test.go", "src/app_test.go", "go", "backend_test", "mcp"),
		symbolNode("Struct:src/app.go:Server", scopeir.NodeStruct, "Server", "src/app.go", 5, 1, 20, 1, "public"),
		symbolNode("Method:src/app.go:Server.Start", scopeir.NodeMethod, "Start", "src/app.go", 10, 2, 15, 2, ""),
		symbolNode("Function:src/app.go:NewServer", scopeir.NodeFunction, "NewServer", "src/app.go", 22, 1, 32, 1, ""),
		symbolNode("Function:src/store.go:Save", scopeir.NodeFunction, "Save", "src/store.go", 4, 1, 8, 1, ""),
		symbolNode("Function:src/store.go:Load", scopeir.NodeFunction, "Load", "src/store.go", 10, 1, 14, 1, ""),
		symbolNode("Function:src/app_test.go:TestNewServer", scopeir.NodeFunction, "TestNewServer", "src/app_test.go", 9, 1, 20, 1, ""),
		resolutionGapNode(),
		{ID: "Process:mcp-initialize", Label: scopeir.NodeProcess, Properties: graph.NodeProperties{"name": "MCP initialize"}},
		{ID: "Route:GET /api/app", Label: scopeir.NodeRoute, Properties: graph.NodeProperties{"name": "GET /api/app"}},
		{ID: "Tool:context", Label: scopeir.NodeTool, Properties: graph.NodeProperties{"name": "context"}},
	} {
		g.AddNode(node)
	}

	relationships := []graph.Relationship{
		defines("File:src/app.go", "Struct:src/app.go:Server"),
		defines("File:src/app.go", "Method:src/app.go:Server.Start"),
		defines("File:src/app.go", "Function:src/app.go:NewServer"),
		defines("File:src/store.go", "Function:src/store.go:Save"),
		defines("File:src/store.go", "Function:src/store.go:Load"),
		defines("File:src/app_test.go", "Function:src/app_test.go:TestNewServer"),
		{
			ID:       "rel:contains:server-start",
			SourceID: "Struct:src/app.go:Server",
			TargetID: "Method:src/app.go:Server.Start",
			Type:     graph.RelContains,
		},
		call("rel:call:new-start", "Function:src/app.go:NewServer", "Method:src/app.go:Server.Start", "src/app.go", 24, "site:new-start"),
		call("rel:call:new-save", "Function:src/app.go:NewServer", "Function:src/store.go:Save", "src/app.go", 25, "site:new-save"),
		call("rel:call:new-load", "Function:src/app.go:NewServer", "Function:src/store.go:Load", "src/app.go", 26, "site:new-load"),
		call("rel:call:test-new", "Function:src/app_test.go:TestNewServer", "Function:src/app.go:NewServer", "src/app_test.go", 12, "site:test-new"),
		{ID: "rel:step:mcp", SourceID: "Function:src/app.go:NewServer", TargetID: "Process:mcp-initialize", Type: graph.RelStepInProcess, Confidence: 0.9},
		{ID: "rel:route:app", SourceID: "Function:src/app.go:NewServer", TargetID: "Route:GET /api/app", Type: graph.RelHandlesRoute, Confidence: 0.8},
		{ID: "rel:tool:context", SourceID: "Function:src/app.go:NewServer", TargetID: "Tool:context", Type: graph.RelHandlesTool, Confidence: 0.85},
	}
	if reverseRelationships {
		for left, right := 0, len(relationships)-1; left < right; left, right = left+1, right-1 {
			relationships[left], relationships[right] = relationships[right], relationships[left]
		}
	}
	for _, relationship := range relationships {
		g.AddRelationship(relationship)
	}
	return g
}

func resolutionGap(id string, filePath string, sourceNodeID string, targetText string, gapKind string, classification string, actionability string, line int) graph.Node {
	return graph.Node{
		ID:    id,
		Label: scopeir.NodeResolutionGap,
		Properties: graph.NodeProperties{
			"name":             targetText,
			"filePath":         filePath,
			"sourceNodeId":     sourceNodeID,
			"targetText":       targetText,
			"gapKind":          gapKind,
			"classification":   classification,
			"actionability":    actionability,
			"proofKind":        "none",
			"sourceSiteId":     "SourceSite:" + id,
			"sourceSiteStatus": "unresolved_local_binding",
			"startLine":        line,
			"startCol":         2,
			"endLine":          line,
			"endCol":           12,
		},
	}
}

func fileListBenchmarkGraph(fileCount int, relationshipCount int) *graph.Graph {
	g := graph.New()
	for i := 0; i < fileCount; i++ {
		filePath := benchmarkFilePath(i)
		fileID := "File:" + filePath
		symbolID := benchmarkSymbolID(i)
		g.AddNode(fileNode(fileID, filePath, "go", "backend", "benchmark"))
		g.AddNode(symbolNode(symbolID, scopeir.NodeFunction, "fn", filePath, 1, 1, 2, 1, ""))
		g.AddRelationship(defines(fileID, symbolID))
		if i%7 == 0 {
			g.AddNode(graph.Node{
				ID:    "ResolutionGap:bench:" + symbolID,
				Label: scopeir.NodeResolutionGap,
				Properties: graph.NodeProperties{
					"name":             "missing",
					"filePath":         filePath,
					"sourceNodeId":     symbolID,
					"targetText":       "missing",
					"gapKind":          "unresolved_call",
					"classification":   "in_repo_unresolved",
					"actionability":    "analyzer_gap",
					"sourceSiteId":     "SourceSite:" + filePath,
					"sourceSiteStatus": "unresolved_local_binding",
					"startLine":        2,
					"startCol":         1,
				},
			})
		}
	}
	for i := 0; i < relationshipCount; i++ {
		source := i % fileCount
		target := (i*17 + 3) % fileCount
		g.AddRelationship(call(
			"rel:bench:"+itoa(i),
			benchmarkSymbolID(source),
			benchmarkSymbolID(target),
			benchmarkFilePath(source),
			i%100+1,
			"site:bench",
		))
	}
	return g
}

func benchmarkFilePath(index int) string {
	return "src/bench/file" + itoa(index) + ".go"
}

func benchmarkSymbolID(index int) string {
	return "Function:" + benchmarkFilePath(index) + ":fn"
}

func itoa(value int) string {
	if value == 0 {
		return "0"
	}
	digits := [20]byte{}
	index := len(digits)
	for value > 0 {
		index--
		digits[index] = byte('0' + value%10)
		value /= 10
	}
	return string(digits[index:])
}

func fileNode(id string, filePath string, language string, appLayer string, functionalArea string) graph.Node {
	return graph.Node{
		ID:    id,
		Label: scopeir.NodeFile,
		Properties: graph.NodeProperties{
			"name":           filePath,
			"filePath":       filePath,
			"language":       language,
			"appLayer":       appLayer,
			"functionalArea": functionalArea,
		},
	}
}

func symbolNode(id string, label scopeir.NodeLabel, name string, filePath string, startLine int, startCol int, endLine int, endCol int, visibility string) graph.Node {
	properties := graph.NodeProperties{
		"name":      name,
		"filePath":  filePath,
		"startLine": startLine,
		"startCol":  startCol,
		"endLine":   endLine,
		"endCol":    endCol,
	}
	if visibility != "" {
		properties["visibility"] = visibility
	}
	return graph.Node{
		ID:         id,
		Label:      label,
		Properties: properties,
	}
}

func resolutionGapNode() graph.Node {
	return graph.Node{
		ID:    "ResolutionGap:site-dynamic",
		Label: scopeir.NodeResolutionGap,
		Properties: graph.NodeProperties{
			"name":             "dynamicHandler",
			"filePath":         "src/app.go",
			"sourceNodeId":     "Function:src/app.go:NewServer",
			"targetText":       "dynamicHandler",
			"gapKind":          "unresolved_call",
			"factFamily":       "call",
			"classification":   "in_repo_unresolved",
			"actionability":    "analyzer_gap",
			"proofKind":        "none",
			"sourceSiteId":     "SourceSite:src/app.go#call#dynamicHandler#30#4#30#18",
			"sourceSiteStatus": "unresolved_local_binding",
			"startLine":        30,
			"startCol":         4,
			"endLine":          30,
			"endCol":           18,
		},
	}
}

func defines(sourceID string, targetID string) graph.Relationship {
	return graph.Relationship{
		ID:       "rel:defines:" + sourceID + ":" + targetID,
		SourceID: sourceID,
		TargetID: targetID,
		Type:     graph.RelDefines,
	}
}

func call(id string, sourceID string, targetID string, filePath string, line int, sourceSiteID string) graph.Relationship {
	return graph.Relationship{
		ID:               id,
		SourceID:         sourceID,
		TargetID:         targetID,
		Type:             graph.RelCalls,
		Confidence:       1,
		FilePath:         filePath,
		SourceSiteID:     sourceSiteID,
		SourceSiteIDs:    []string{sourceSiteID},
		SourceSiteCount:  1,
		SourceSiteStatus: "resolved",
		ProofKind:        "scope-binding",
		StartLine:        line,
		StartCol:         2,
		EndLine:          line,
		EndCol:           20,
	}
}
