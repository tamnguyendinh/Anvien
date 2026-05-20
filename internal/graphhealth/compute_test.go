package graphhealth

import (
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
	if summary.CountedRelationshipCount != 1 {
		t.Errorf("summary counted relationships=%d want 1", summary.CountedRelationshipCount)
	}
	if got := summary.TopologyStatusCounts[string(TopologyNoIncoming)]; got != 1 {
		t.Errorf("summary no_incoming=%d want 1", got)
	}
	if got := summary.ExcludedEdgeCounts[ExcludedEdgeStructural]; got != 1 {
		t.Errorf("summary structural excluded=%d want 1", got)
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

func findNode(g *graph.Graph, id string) *graph.Node {
	for i := range g.Nodes {
		if g.Nodes[i].ID == id {
			return &g.Nodes[i]
		}
	}
	return nil
}
