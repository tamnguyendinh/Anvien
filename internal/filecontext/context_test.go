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
