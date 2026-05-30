package filecontext

import (
	"encoding/json"
	"testing"

	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
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
	if context.Summary.UnresolvedSourceSiteCount != 1 || context.Quality.UnresolvedCalls != 1 {
		t.Fatalf("unresolved summary = %d calls=%d, want 1/1", context.Summary.UnresolvedSourceSiteCount, context.Quality.UnresolvedCalls)
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

func TestBuildFileContextReturnsFalseForMissingFile(t *testing.T) {
	_, ok := NewBuilder(fileContextFixture(false)).BuildFileContext("src/missing.go", Options{})
	if ok {
		t.Fatalf("BuildFileContext() found missing file")
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
	if unresolved.Total != 1 || unresolved.Files[0].Path != "src/app.go" || unresolved.Files[0].UnresolvedSourceSiteCount != 1 {
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
