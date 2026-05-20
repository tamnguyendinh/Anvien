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
	g.AddNode(graph.Node{ID: "Function:main", Label: scopeir.NodeFunction, Properties: map[string]any{"name": "main", "filePath": "main.go", "isExported": true}})
	g.AddRelationship(graph.Relationship{ID: "r1", SourceID: "File:main.go", TargetID: "Function:main", Type: graph.RelDefines})

	// A call edge (counts)
	g.AddNode(graph.Node{ID: "Function:helper", Label: scopeir.NodeFunction, Properties: map[string]any{"name": "helper", "filePath": "main.go"}})
	g.AddRelationship(graph.Relationship{ID: "r2", SourceID: "Function:main", TargetID: "Function:helper", Type: graph.RelCalls})

	Compute(g)

	// main has outgoing counted (call), incoming 0 from counted -> no_incoming + exported modifier
	main := findNode(g, "Function:main")
	if main == nil {
		t.Fatal("main node not found after compute")
	}
	if ts, _ := main.Properties["topologyStatus"].(string); ts != string(TopologyNoIncoming) {
		t.Errorf("main topologyStatus=%s want no_incoming", ts)
	}
	reasons, _ := main.Properties["expectedIsolationReasons"].([]string)
	if len(reasons) == 0 || reasons[0] != ReasonExportedAPI {
		t.Errorf("main reasons=%v want exported_api modifier", reasons)
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
}

func findNode(g *graph.Graph, id string) *graph.Node {
	for i := range g.Nodes {
		if g.Nodes[i].ID == id {
			return &g.Nodes[i]
		}
	}
	return nil
}
