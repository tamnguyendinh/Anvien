package graphhealth

import (
	"encoding/json"
	"testing"

	"github.com/tamnguyendinh/avmatrix-go/internal/graph"
	"github.com/tamnguyendinh/avmatrix-go/internal/scopeir"
)

func TestCompute_BasicConnectivity(t *testing.T) {
	g := graph.New()

	// Structural only: file defines func (should not count for connectivity)
	g.AddNode(graph.Node{ID: "File:main.go", Label: scopeir.NodeFile, Properties: map[string]any{"filePath": "main.go"}})
	g.AddNode(graph.Node{ID: "Function:entry", Label: scopeir.NodeFunction, Properties: map[string]any{"name": "Entry", "filePath": "main.go", "isExported": true}})
	g.AddRelationship(graph.Relationship{ID: "r1", SourceID: "File:main.go", TargetID: "Function:entry", Type: graph.RelDefines})

	// A call edge (counts)
	g.AddNode(graph.Node{ID: "Function:helper", Label: scopeir.NodeFunction, Properties: map[string]any{"name": "helper", "filePath": "main.go"}})
	g.AddRelationship(graph.Relationship{ID: "r2", SourceID: "Function:entry", TargetID: "Function:helper", Type: graph.RelCalls})
	g.AddNode(graph.Node{ID: "Process:entry", Label: scopeir.NodeProcess, Properties: map[string]any{"name": "entry process"}})
	g.AddRelationship(graph.Relationship{ID: "r3", SourceID: "Function:entry", TargetID: "Process:entry", Type: graph.RelStepInProcess})

	summary := ComputeSummary(g)

	// entry has outgoing counted (call), incoming 0 from counted -> no_incoming + exported modifier
	entry := findNode(g, "Function:entry")
	if entry == nil {
		t.Fatal("entry node not found after compute")
	}
	if ts, _ := entry.Properties["topologyStatus"].(string); ts != string(TopologyNoIncoming) {
		t.Errorf("entry topologyStatus=%s want no_incoming", ts)
	}
	reasons, _ := entry.Properties["expectedIsolationReasons"].([]string)
	if len(reasons) == 0 || reasons[0] != ReasonExportedAPI {
		t.Errorf("entry reasons=%v want exported_api modifier", reasons)
	}
	if conf, _ := entry.Properties["confidence"].(string); conf != ConfidenceCandidate {
		t.Errorf("entry confidence=%s want candidate for exported-only modifier", conf)
	}
	health, ok := entry.Properties["graphHealth"].(NodeHealth)
	if !ok {
		t.Fatalf("entry graphHealth missing or wrong type: %#v", entry.Properties["graphHealth"])
	}
	if health.ExcludedEdgeCounts[ExcludedEdgeStructural] != 1 {
		t.Errorf("entry structural excluded count=%d want 1", health.ExcludedEdgeCounts[ExcludedEdgeStructural])
	}

	// helper has incoming counted, no out -> no_outgoing (normal leaf)
	helper := findNode(g, "Function:helper")
	if ts, _ := helper.Properties["topologyStatus"].(string); ts != string(TopologyNoOutgoing) {
		t.Errorf("helper topologyStatus=%s want no_outgoing", ts)
	}

	// isolated node (no counted edges at all)
	g.AddNode(graph.Node{ID: "Function:dead", Label: scopeir.NodeFunction, Properties: map[string]any{"name": "dead", "filePath": "dead.go"}})
	Compute(g) // re-compute includes new
	dead := findNode(g, "Function:dead")
	if ts, _ := dead.Properties["topologyStatus"].(string); ts != string(TopologyTrueIsolated) {
		t.Errorf("dead topologyStatus=%s want true_isolated", ts)
	}
	if conf, _ := dead.Properties["confidence"].(string); conf != ConfidenceCandidate {
		t.Errorf("dead confidence=%s want candidate (no reasons)", conf)
	}

	if summary.PolicyVersion != PolicyVersion {
		t.Errorf("summary policy version=%q want %q", summary.PolicyVersion, PolicyVersion)
	}
	if summary.CountedRelationshipCount != 2 {
		t.Errorf("summary counted relationships=%d want 2", summary.CountedRelationshipCount)
	}
	if got := summary.TopologyStatusCounts[string(TopologyNoIncoming)]; got != 1 {
		t.Errorf("summary no_incoming=%d want 1", got)
	}
	if got := summary.ExcludedEdgeCounts[ExcludedEdgeStructural]; got != 1 {
		t.Errorf("summary structural excluded=%d want 1", got)
	}
}

func TestCompute_DetachedComponent(t *testing.T) {
	g := graph.New()
	g.AddNode(graph.Node{ID: "Function:main", Label: scopeir.NodeFunction, Properties: map[string]any{"name": "main", "filePath": "cmd/app/main.go", "isExported": true}})
	g.AddNode(graph.Node{ID: "Function:reachable", Label: scopeir.NodeFunction, Properties: map[string]any{"name": "reachable", "filePath": "cmd/app/main.go"}})
	g.AddRelationship(graph.Relationship{ID: "r-root", SourceID: "Function:main", TargetID: "Function:reachable", Type: graph.RelCalls})

	g.AddNode(graph.Node{ID: "Function:detachedA", Label: scopeir.NodeFunction, Properties: map[string]any{"name": "detachedA", "filePath": "pkg/detached.go"}})
	g.AddNode(graph.Node{ID: "Function:detachedB", Label: scopeir.NodeFunction, Properties: map[string]any{"name": "detachedB", "filePath": "pkg/detached.go"}})
	g.AddRelationship(graph.Relationship{ID: "r-detached", SourceID: "Function:detachedA", TargetID: "Function:detachedB", Type: graph.RelCalls})

	summary := ComputeSummary(g)

	detachedA := findNode(g, "Function:detachedA")
	detachedB := findNode(g, "Function:detachedB")
	for _, node := range []*graph.Node{detachedA, detachedB} {
		if node == nil {
			t.Fatal("detached test node not found")
		}
		if ts, _ := node.Properties["topologyStatus"].(string); ts != string(TopologyDetached) {
			t.Fatalf("%s topologyStatus=%s want detached_component", node.ID, ts)
		}
		if reachable, _ := node.Properties["componentReachableFromRoot"].(bool); reachable {
			t.Fatalf("%s componentReachableFromRoot=true want false", node.ID)
		}
		health, ok := node.Properties["graphHealth"].(NodeHealth)
		if !ok {
			t.Fatalf("%s graphHealth missing: %#v", node.ID, node.Properties["graphHealth"])
		}
		if health.ComponentSize != 2 || health.ComponentReachableFromRoot {
			t.Fatalf("%s component health = %#v", node.ID, health)
		}
	}
	if detachedA.Properties["componentId"] != detachedB.Properties["componentId"] {
		t.Fatalf("detached nodes should share component ID: %v vs %v", detachedA.Properties["componentId"], detachedB.Properties["componentId"])
	}

	reachable := findNode(g, "Function:reachable")
	if ts, _ := reachable.Properties["topologyStatus"].(string); ts == string(TopologyDetached) {
		t.Fatalf("reachable node should not be detached: %#v", reachable.Properties)
	}
	if got := summary.DetachedComponentCount; got != 1 {
		t.Fatalf("detached component count=%d want 1", got)
	}
	if got := summary.TopologyStatusCounts[string(TopologyDetached)]; got != 2 {
		t.Fatalf("detached topology node count=%d want 2", got)
	}
	if len(summary.LargestDetachedComponents) != 1 {
		t.Fatalf("largest detached components=%#v want one component", summary.LargestDetachedComponents)
	}
	if summary.LargestDetachedComponents[0].NodeCount != 2 || summary.LargestDetachedComponents[0].CountedEdgeCount != 1 {
		t.Fatalf("detached component summary=%#v", summary.LargestDetachedComponents[0])
	}
}

func TestCompute_ExpectedReasonConfidenceAndSummary(t *testing.T) {
	g := graph.New()
	g.AddNode(graph.Node{ID: "Function:testHelper", Label: scopeir.NodeFunction, Properties: map[string]any{"name": "testHelper", "filePath": "pkg/foo_test.go", "isExported": true}})
	g.AddNode(graph.Node{ID: "Function:publicAPI", Label: scopeir.NodeFunction, Properties: map[string]any{"name": "PublicAPI", "filePath": "pkg/api.go", "isExported": true}})
	g.AddNode(graph.Node{ID: "Route:/api/users", Label: scopeir.NodeRoute, Properties: map[string]any{"name": "GET /api/users", "filePath": "internal/httpapi/users.go"}})

	summary := ComputeSummary(g)

	testHelper := findNode(g, "Function:testHelper")
	if conf, _ := testHelper.Properties["confidence"].(string); conf != ConfidenceExpected {
		t.Errorf("test helper confidence=%s want expected", conf)
	}
	publicAPI := findNode(g, "Function:publicAPI")
	if conf, _ := publicAPI.Properties["confidence"].(string); conf != ConfidenceCandidate {
		t.Errorf("exported-only confidence=%s want candidate", conf)
	}
	route := findNode(g, "Route:/api/users")
	if conf, _ := route.Properties["confidence"].(string); conf != ConfidenceExpected {
		t.Errorf("framework entry confidence=%s want expected", conf)
	}

	if got := summary.ExpectedIsolationReasonCounts[ReasonTest]; got != 1 {
		t.Errorf("test reason count=%d want 1", got)
	}
	if got := summary.ExpectedIsolationReasonCounts[ReasonExportedAPI]; got != 2 {
		t.Errorf("exported reason count=%d want 2", got)
	}
	if got := summary.ExpectedIsolationReasonCounts[ReasonFrameworkEntry]; got != 1 {
		t.Errorf("framework reason count=%d want 1", got)
	}
	if got := summary.ConfidenceCounts[ConfidenceExpected]; got != 2 {
		t.Errorf("expected confidence count=%d want 2", got)
	}
	if got := summary.ConfidenceCounts[ConfidenceCandidate]; got != 1 {
		t.Errorf("candidate confidence count=%d want 1", got)
	}
}

func TestCompute_UnresolvedDiagnosticPreservesTopologyAndMarksUnknownConfidence(t *testing.T) {
	g := graph.New()
	g.AddNode(graph.Node{ID: "Function:main", Label: scopeir.NodeFunction, Properties: map[string]any{"name": "main", "filePath": "cmd/app/main.go", "isExported": true}})
	g.AddNode(graph.Node{ID: "Function:source", Label: scopeir.NodeFunction, Properties: map[string]any{"name": "source", "filePath": "src/app.ts"}})
	g.AddNode(graph.Node{ID: "Function:target", Label: scopeir.NodeFunction, Properties: map[string]any{"name": "target", "filePath": "src/app.ts"}})
	g.AddNode(graph.Node{ID: "Function:isolated", Label: scopeir.NodeFunction, Properties: map[string]any{"name": "isolated", "filePath": "src/app.ts"}})
	g.AddNode(graph.Node{ID: "Function:detachedA", Label: scopeir.NodeFunction, Properties: map[string]any{"name": "detachedA", "filePath": "src/detached.ts"}})
	g.AddNode(graph.Node{ID: "Function:detachedB", Label: scopeir.NodeFunction, Properties: map[string]any{"name": "detachedB", "filePath": "src/detached.ts"}})
	g.AddRelationship(graph.Relationship{ID: "r-main-source", SourceID: "Function:main", TargetID: "Function:source", Type: graph.RelCalls})
	g.AddRelationship(graph.Relationship{ID: "r-call", SourceID: "Function:source", TargetID: "Function:target", Type: graph.RelCalls})
	g.AddRelationship(graph.Relationship{ID: "r-detached", SourceID: "Function:detachedA", TargetID: "Function:detachedB", Type: graph.RelCalls})
	for _, item := range []struct {
		nodeID     string
		targetText string
	}{
		{nodeID: "Function:main", targetText: "make"},
		{nodeID: "Function:target", targetText: "len"},
		{nodeID: "Function:isolated", targetText: "node.Kind"},
		{nodeID: "Function:detachedA", targetText: "time.Second"},
	} {
		if !AppendDiagnosticToNode(g, item.nodeID, Diagnostic{
			Kind:             DiagnosticUnresolvedReference,
			FactFamily:       "call",
			TargetText:       item.targetText,
			ResolutionSource: "scope-resolution",
			FilePath:         "src/app.ts",
			StartLine:        2,
		}) {
			t.Fatalf("failed to attach unresolved diagnostic to %s", item.nodeID)
		}
	}
	if !AppendDiagnosticToNode(g, "Function:source", Diagnostic{
		Kind:             DiagnosticUnresolvedReference,
		FactFamily:       "call",
		TargetText:       "missing",
		ResolutionSource: "scope-resolution",
		FilePath:         "src/app.ts",
		StartLine:        3,
	}) {
		t.Fatal("failed to attach unresolved diagnostic to source")
	}
	if !AppendDiagnosticToNode(g, "Function:source", Diagnostic{
		Kind:             DiagnosticUnresolvedReference,
		FactFamily:       "call",
		TargetText:       "otherMissing",
		ResolutionSource: "scope-resolution",
		FilePath:         "src/app.ts",
		StartLine:        4,
	}) {
		t.Fatal("failed to attach second unresolved diagnostic to source")
	}
	SetResolutionMetadata(g, 6, 5, 1)

	summary := ComputeSummary(g)

	cases := []struct {
		nodeID       string
		wantTopology TopologyStatus
	}{
		{nodeID: "Function:main", wantTopology: TopologyNoIncoming},
		{nodeID: "Function:source", wantTopology: TopologyConnected},
		{nodeID: "Function:target", wantTopology: TopologyNoOutgoing},
		{nodeID: "Function:isolated", wantTopology: TopologyTrueIsolated},
		{nodeID: "Function:detachedA", wantTopology: TopologyDetached},
	}
	for _, item := range cases {
		node := findNode(g, item.nodeID)
		if node == nil {
			t.Fatalf("%s node not found", item.nodeID)
		}
		health, ok := node.Properties["graphHealth"].(NodeHealth)
		if !ok {
			t.Fatalf("%s graphHealth missing: %#v", item.nodeID, node.Properties["graphHealth"])
		}
		if health.TopologyStatus != item.wantTopology {
			t.Fatalf("%s topologyStatus=%s want %s", item.nodeID, health.TopologyStatus, item.wantTopology)
		}
		if health.Confidence != ConfidenceUnknown {
			t.Fatalf("%s confidence=%s want unknown", item.nodeID, health.Confidence)
		}
	}

	source := findNode(g, "Function:source")
	health, ok := source.Properties["graphHealth"].(NodeHealth)
	if !ok || len(health.Diagnostics) != 2 {
		t.Fatalf("source graphHealth diagnostics = %#v", source.Properties["graphHealth"])
	}
	if health.Diagnostics[0].Count != 1 || health.Diagnostics[0].TargetText != "missing" ||
		health.Diagnostics[1].Count != 1 || health.Diagnostics[1].TargetText != "otherMissing" {
		t.Fatalf("source diagnostic aggregation = %#v", health.Diagnostics)
	}
	if health.Diagnostics[0].Classification != DiagnosticClassificationInRepoUnresolved ||
		health.Diagnostics[0].Actionability != DiagnosticActionabilityAnalyzerGap {
		t.Fatalf("source diagnostic classification = %#v", health.Diagnostics[0])
	}
	if got := summary.TopologyStatusCounts[string(TopologyUnknown)]; got != 0 {
		t.Fatalf("unknown topology count=%d want 0 for valid graph nodes", got)
	}
	if got := summary.DiagnosticCounts[DiagnosticUnresolvedReference]; got != 6 {
		t.Fatalf("unresolved diagnostic count=%d want 6", got)
	}
	if summary.DiagnosticClassificationCounts[DiagnosticClassificationBuiltin] != 2 ||
		summary.DiagnosticClassificationCounts[DiagnosticClassificationStandardLibrary] != 1 ||
		summary.DiagnosticClassificationCounts[DiagnosticClassificationInRepoUnresolved] != 3 {
		t.Fatalf("diagnostic classification counts = %#v", summary.DiagnosticClassificationCounts)
	}
	if summary.DiagnosticActionabilityCounts[DiagnosticActionabilityNonActionable] != 3 ||
		summary.DiagnosticActionabilityCounts[DiagnosticActionabilityAnalyzerGap] != 3 {
		t.Fatalf("diagnostic actionability counts = %#v", summary.DiagnosticActionabilityCounts)
	}
	if summary.UnresolvedReferenceCount != 6 ||
		summary.SourceBackedUnresolvedReferenceCount != 6 ||
		summary.UnattributedUnresolvedReferenceCount != 0 {
		t.Fatalf("unresolved summary counts = %#v", summary)
	}
}

func TestCompute_NormalizesDecodedDiagnostics(t *testing.T) {
	raw := []byte(`{
  "nodes": [
    {
      "id": "Function:source",
      "label": "Function",
      "properties": {
        "name": "source",
        "filePath": "src/app.ts",
        "graphHealthDiagnostics": [
          {
            "kind": "unresolved_reference",
            "factFamily": "type-reference",
            "sourceNodeId": "Function:source",
            "targetText": "MissingType",
            "resolutionSource": "scope-resolution",
            "filePath": "src/app.ts",
            "startLine": 7
          }
        ]
      }
    }
  ],
  "relationships": [],
  "metadata": {
    "resolution": {
      "unresolvedReferences": 1,
      "sourceBackedUnresolvedReferences": 1,
      "unattributedUnresolvedReferences": 0
    }
  }
}`)
	var g graph.Graph
	if err := json.Unmarshal(raw, &g); err != nil {
		t.Fatalf("unmarshal graph: %v", err)
	}

	summary := ComputeSummary(&g)

	source := findNode(&g, "Function:source")
	if source == nil {
		t.Fatal("source node missing")
	}
	health, ok := source.Properties["graphHealth"].(NodeHealth)
	if !ok || len(health.Diagnostics) != 1 {
		t.Fatalf("decoded graphHealth diagnostics = %#v", source.Properties["graphHealth"])
	}
	if health.Diagnostics[0].FactFamily != "type-reference" || health.Diagnostics[0].StartLine != 7 {
		t.Fatalf("decoded diagnostic not normalized: %#v", health.Diagnostics[0])
	}
	if health.Diagnostics[0].Classification != DiagnosticClassificationInRepoUnresolved ||
		health.Diagnostics[0].Actionability != DiagnosticActionabilityAnalyzerGap {
		t.Fatalf("decoded diagnostic classification = %#v", health.Diagnostics[0])
	}
	if summary.UnresolvedReferenceCount != 1 || summary.DiagnosticCounts[DiagnosticUnresolvedReference] != 1 {
		t.Fatalf("decoded summary = %#v", summary)
	}
}

func TestDiagnostics_ClassifiesObservedUnresolvedTargets(t *testing.T) {
	cases := []struct {
		target             string
		wantClassification string
		wantActionability  string
	}{
		{target: "testing.T", wantClassification: DiagnosticClassificationTestFramework, wantActionability: DiagnosticActionabilityNonActionable},
		{target: "testing.B", wantClassification: DiagnosticClassificationTestFramework, wantActionability: DiagnosticActionabilityNonActionable},
		{target: "make", wantClassification: DiagnosticClassificationBuiltin, wantActionability: DiagnosticActionabilityNonActionable},
		{target: "len", wantClassification: DiagnosticClassificationBuiltin, wantActionability: DiagnosticActionabilityNonActionable},
		{target: "append", wantClassification: DiagnosticClassificationBuiltin, wantActionability: DiagnosticActionabilityNonActionable},
		{target: "string", wantClassification: DiagnosticClassificationBuiltin, wantActionability: DiagnosticActionabilityNonActionable},
		{target: "int", wantClassification: DiagnosticClassificationBuiltin, wantActionability: DiagnosticActionabilityNonActionable},
		{target: "map[string]any", wantClassification: DiagnosticClassificationBuiltin, wantActionability: DiagnosticActionabilityNonActionable},
		{target: "fmt.Errorf", wantClassification: DiagnosticClassificationStandardLibrary, wantActionability: DiagnosticActionabilityNonActionable},
		{target: "time.Second", wantClassification: DiagnosticClassificationStandardLibrary, wantActionability: DiagnosticActionabilityNonActionable},
		{target: "strings.TrimSpace", wantClassification: DiagnosticClassificationStandardLibrary, wantActionability: DiagnosticActionabilityNonActionable},
		{target: "context.Context", wantClassification: DiagnosticClassificationStandardLibrary, wantActionability: DiagnosticActionabilityNonActionable},
		{target: "filepath.Join", wantClassification: DiagnosticClassificationStandardLibrary, wantActionability: DiagnosticActionabilityNonActionable},
		{target: "t.Helper", wantClassification: DiagnosticClassificationTestFramework, wantActionability: DiagnosticActionabilityNonActionable},
		{target: "t.TempDir", wantClassification: DiagnosticClassificationTestFramework, wantActionability: DiagnosticActionabilityNonActionable},
		{target: "t.Fatalf", wantClassification: DiagnosticClassificationTestFramework, wantActionability: DiagnosticActionabilityNonActionable},
		{target: "node.Kind", wantClassification: DiagnosticClassificationInRepoUnresolved, wantActionability: DiagnosticActionabilityAnalyzerGap},
		{target: "c.text", wantClassification: DiagnosticClassificationInRepoUnresolved, wantActionability: DiagnosticActionabilityAnalyzerGap},
		{target: "uuid.New", wantClassification: DiagnosticClassificationExternalLibrary, wantActionability: DiagnosticActionabilityReview},
	}
	for _, item := range cases {
		diagnostic := normalizeDiagnosticMetadata(Diagnostic{
			Kind:       DiagnosticUnresolvedReference,
			TargetText: item.target,
		})
		if diagnostic.Classification != item.wantClassification || diagnostic.Actionability != item.wantActionability {
			t.Fatalf("%s classification/actionability = %s/%s, want %s/%s",
				item.target,
				diagnostic.Classification,
				diagnostic.Actionability,
				item.wantClassification,
				item.wantActionability,
			)
		}
	}
}

func TestCompute_ExpectedIsolationReasonMatrix(t *testing.T) {
	g := graph.New()
	cases := []struct {
		id               string
		label            scopeir.NodeLabel
		properties       graph.NodeProperties
		reason           string
		wantConfidence   string
		wantTopology     TopologyStatus
		wantSummaryCount int
	}{
		{
			id:               "Function:testHelper",
			label:            scopeir.NodeFunction,
			properties:       graph.NodeProperties{"name": "testHelper", "filePath": "pkg/service_test.go"},
			reason:           ReasonTest,
			wantConfidence:   ConfidenceExpected,
			wantTopology:     TopologyTrueIsolated,
			wantSummaryCount: 1,
		},
		{
			id:               "Function:fixtureHelper",
			label:            scopeir.NodeFunction,
			properties:       graph.NodeProperties{"name": "fixtureHelper", "filePath": "pkg/fixtures/sample.go"},
			reason:           ReasonFixture,
			wantConfidence:   ConfidenceExpected,
			wantTopology:     TopologyTrueIsolated,
			wantSummaryCount: 1,
		},
		{
			id:               "Function:generatedHelper",
			label:            scopeir.NodeFunction,
			properties:       graph.NodeProperties{"name": "generatedHelper", "filePath": "pkg/generated/client.go"},
			reason:           ReasonGenerated,
			wantConfidence:   ConfidenceExpected,
			wantTopology:     TopologyTrueIsolated,
			wantSummaryCount: 1,
		},
		{
			id:               "Function:vendorHelper",
			label:            scopeir.NodeFunction,
			properties:       graph.NodeProperties{"name": "vendorHelper", "filePath": "third_party/vendor/module/file.go"},
			reason:           ReasonVendor,
			wantConfidence:   ConfidenceExpected,
			wantTopology:     TopologyTrueIsolated,
			wantSummaryCount: 1,
		},
		{
			id:               "Section:docs-guide",
			label:            scopeir.NodeSection,
			properties:       graph.NodeProperties{"name": "Guide", "filePath": "docs/guide.md"},
			reason:           ReasonDocumentation,
			wantConfidence:   ConfidenceExpected,
			wantTopology:     TopologyTrueIsolated,
			wantSummaryCount: 1,
		},
		{
			id:               "Function:migrationHelper",
			label:            scopeir.NodeFunction,
			properties:       graph.NodeProperties{"name": "migrationHelper", "filePath": "db/migrations/001_migrate.sql"},
			reason:           ReasonMigration,
			wantConfidence:   ConfidenceExpected,
			wantTopology:     TopologyTrueIsolated,
			wantSummaryCount: 1,
		},
		{
			id:               "Function:publicAPI",
			label:            scopeir.NodeFunction,
			properties:       graph.NodeProperties{"name": "PublicAPI", "filePath": "pkg/api.go", "isExported": true},
			reason:           ReasonExportedAPI,
			wantConfidence:   ConfidenceCandidate,
			wantTopology:     TopologyTrueIsolated,
			wantSummaryCount: 1,
		},
		{
			id:               "Route:/api/users",
			label:            scopeir.NodeRoute,
			properties:       graph.NodeProperties{"name": "GET /api/users", "filePath": "internal/httpapi/users.go"},
			reason:           ReasonFrameworkEntry,
			wantConfidence:   ConfidenceExpected,
			wantTopology:     TopologyTrueIsolated,
			wantSummaryCount: 1,
		},
		{
			id:               "Function:cliCommand",
			label:            scopeir.NodeFunction,
			properties:       graph.NodeProperties{"name": "cliCommand", "filePath": "pkg/internal/cli/run.go"},
			reason:           ReasonCLIMCP,
			wantConfidence:   ConfidenceExpected,
			wantTopology:     TopologyTrueIsolated,
			wantSummaryCount: 1,
		},
	}
	for _, item := range cases {
		g.AddNode(graph.Node{ID: item.id, Label: item.label, Properties: item.properties})
	}

	summary := ComputeSummary(g)

	for _, item := range cases {
		node := findNode(g, item.id)
		if node == nil {
			t.Fatalf("%s node not found", item.id)
		}
		health, ok := node.Properties["graphHealth"].(NodeHealth)
		if !ok {
			t.Fatalf("%s graphHealth missing: %#v", item.id, node.Properties["graphHealth"])
		}
		if health.TopologyStatus != item.wantTopology {
			t.Fatalf("%s topologyStatus=%s want %s", item.id, health.TopologyStatus, item.wantTopology)
		}
		if health.Confidence != item.wantConfidence {
			t.Fatalf("%s confidence=%s want %s", item.id, health.Confidence, item.wantConfidence)
		}
		if !containsString(health.ExpectedIsolationReasons, item.reason) {
			t.Fatalf("%s reasons=%v want %s", item.id, health.ExpectedIsolationReasons, item.reason)
		}
		if got := summary.ExpectedIsolationReasonCounts[item.reason]; got != item.wantSummaryCount {
			t.Fatalf("summary reason %s count=%d want %d", item.reason, got, item.wantSummaryCount)
		}
	}
}

func TestCompute_ConnectedTopologyAndOtherExclusion(t *testing.T) {
	g := graph.New()
	g.AddNode(graph.Node{ID: "Function:main", Label: scopeir.NodeFunction, Properties: graph.NodeProperties{"name": "main", "filePath": "cmd/app/main.go", "isExported": true}})
	g.AddNode(graph.Node{ID: "Function:middle", Label: scopeir.NodeFunction, Properties: graph.NodeProperties{"name": "middle", "filePath": "pkg/middle.go"}})
	g.AddNode(graph.Node{ID: "Function:leaf", Label: scopeir.NodeFunction, Properties: graph.NodeProperties{"name": "leaf", "filePath": "pkg/leaf.go"}})
	g.AddRelationship(graph.Relationship{ID: "call-main-middle", SourceID: "Function:main", TargetID: "Function:middle", Type: graph.RelCalls})
	g.AddRelationship(graph.Relationship{ID: "call-middle-leaf", SourceID: "Function:middle", TargetID: "Function:leaf", Type: graph.RelCalls})
	g.AddRelationship(graph.Relationship{ID: "display-middle-leaf", SourceID: "Function:middle", TargetID: "Function:leaf", Type: graph.RelationshipType("DISPLAY_ONLY")})

	summary := ComputeSummary(g)

	middle := findNode(g, "Function:middle")
	if middle == nil {
		t.Fatal("middle node not found")
	}
	health, ok := middle.Properties["graphHealth"].(NodeHealth)
	if !ok {
		t.Fatalf("middle graphHealth missing: %#v", middle.Properties["graphHealth"])
	}
	if health.TopologyStatus != TopologyConnected {
		t.Fatalf("middle topologyStatus=%s want connected", health.TopologyStatus)
	}
	if health.CountedIncoming != 1 || health.CountedOutgoing != 1 {
		t.Fatalf("middle counted degree = %d/%d want 1/1", health.CountedIncoming, health.CountedOutgoing)
	}
	if health.ExcludedEdgeCounts[ExcludedEdgeOther] != 1 {
		t.Fatalf("middle other exclusions=%d want 1", health.ExcludedEdgeCounts[ExcludedEdgeOther])
	}
	if got := summary.TopologyStatusCounts[string(TopologyConnected)]; got != 1 {
		t.Fatalf("connected count=%d want 1", got)
	}
	if got := summary.ExcludedEdgeCounts[ExcludedEdgeOther]; got != 1 {
		t.Fatalf("summary other exclusions=%d want 1", got)
	}
}

func TestCompute_UnattributedResolutionMetadataDoesNotMarkUnknown(t *testing.T) {
	g := graph.New()
	g.AddNode(graph.Node{ID: "Function:source", Label: scopeir.NodeFunction, Properties: graph.NodeProperties{"name": "source", "filePath": "src/app.ts"}})
	SetResolutionMetadata(g, 5, 0, 5)

	summary := ComputeSummary(g)

	source := findNode(g, "Function:source")
	if source == nil {
		t.Fatal("source node not found")
	}
	health, ok := source.Properties["graphHealth"].(NodeHealth)
	if !ok {
		t.Fatalf("source graphHealth missing: %#v", source.Properties["graphHealth"])
	}
	if health.TopologyStatus != TopologyTrueIsolated {
		t.Fatalf("source topologyStatus=%s want true_isolated", health.TopologyStatus)
	}
	if health.Confidence != ConfidenceCandidate {
		t.Fatalf("source confidence=%s want candidate", health.Confidence)
	}
	if len(health.Diagnostics) != 0 {
		t.Fatalf("source diagnostics=%#v want none", health.Diagnostics)
	}
	if summary.UnresolvedReferenceCount != 5 ||
		summary.SourceBackedUnresolvedReferenceCount != 0 ||
		summary.UnattributedUnresolvedReferenceCount != 5 {
		t.Fatalf("unresolved summary counts = %#v", summary)
	}
	if got := summary.TopologyStatusCounts[string(TopologyUnknown)]; got != 0 {
		t.Fatalf("unknown topology count=%d want 0", got)
	}
	if got := summary.DiagnosticCounts[DiagnosticUnresolvedReference]; got != 0 {
		t.Fatalf("diagnostic count=%d want 0", got)
	}
}

func TestEdgePolicyCountedAndStructuralSets(t *testing.T) {
	counted := []graph.RelationshipType{
		graph.RelCalls,
		graph.RelAccesses,
		graph.RelInherits,
		graph.RelImplements,
		graph.RelExtends,
		graph.RelMethodOverrides,
		graph.RelMethodImplements,
		graph.RelImports,
		graph.RelUses,
		graph.RelDecorates,
		graph.RelWraps,
		graph.RelQueries,
		graph.RelFetches,
		graph.RelStepInProcess,
		graph.RelHandlesRoute,
		graph.RelHandlesTool,
		graph.RelEntryPointOf,
	}
	if len(CountedEdgeTypes) != len(counted) {
		t.Fatalf("counted policy size=%d want %d", len(CountedEdgeTypes), len(counted))
	}
	for _, relType := range counted {
		if !IsCounted(relType) {
			t.Fatalf("%s should be counted", relType)
		}
	}

	structural := []graph.RelationshipType{
		graph.RelContains,
		graph.RelDefines,
		graph.RelHasMethod,
		graph.RelHasProperty,
		graph.RelMemberOf,
	}
	if len(StructuralEdgeTypes) != len(structural) {
		t.Fatalf("structural policy size=%d want %d", len(StructuralEdgeTypes), len(structural))
	}
	for _, relType := range structural {
		if IsCounted(relType) {
			t.Fatalf("%s should not be counted", relType)
		}
		if excludedEdgeCategory(relType) != ExcludedEdgeStructural {
			t.Fatalf("%s exclusion category=%s want structural", relType, excludedEdgeCategory(relType))
		}
	}
	if excludedEdgeCategory(graph.RelationshipType("DISPLAY_ONLY")) != ExcludedEdgeOther {
		t.Fatalf("custom display edge should be categorized as other exclusion")
	}
}

func findNode(g *graph.Graph, id string) *graph.Node {
	for i := range g.Nodes {
		if g.Nodes[i].ID == id {
			return &g.Nodes[i]
		}
	}
	return nil
}

func containsString(values []string, target string) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}
	return false
}
