package mcp

import (
	"testing"

	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/graphhealth"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
	"github.com/tamnguyendinh/anvien/internal/semantic"
)

func TestContextToolReturnsSemanticFieldsAndSourceResolutionGaps(t *testing.T) {
	store, repoPath := newMCPQueryBenchmarkRepo(t)
	g := graph.New()
	step := 1
	g.AddNode(graph.Node{ID: "Process:resolution", Label: scopeir.NodeProcess, Properties: graph.NodeProperties{
		"heuristicLabel":                      "Resolution Flow",
		"stepCount":                           1,
		semantic.AppLayerProperty:             string(semantic.AppLayerBackend),
		semantic.AppLayerSourceProperty:       "process_member_inference",
		semantic.FunctionalAreaProperty:       string(semantic.FunctionalAreaResolution),
		semantic.FunctionalAreaSourceProperty: "process_member_inference",
	}})
	g.AddNode(graph.Node{ID: "Function:source", Label: scopeir.NodeFunction, Properties: graph.NodeProperties{
		"name":                                "resolveCall",
		"filePath":                            "internal/resolution/resolve.go",
		"startLine":                           10,
		"endLine":                             20,
		"topologyStatus":                      string(graphhealth.TopologyConnected),
		"resolutionConfidence":                graphhealth.ResolutionConfidenceDegraded,
		"resolutionGapCount":                  2,
		"resolutionHealthBuckets":             map[string]int{string(graphhealth.ResolutionHealthUnresolvedCallTarget): 2},
		semantic.AppLayerProperty:             string(semantic.AppLayerBackend),
		semantic.AppLayerSourceProperty:       "path_rule",
		semantic.FunctionalAreaProperty:       string(semantic.FunctionalAreaResolution),
		semantic.FunctionalAreaSourceProperty: "path_rule",
	}})
	g.AddNode(graph.Node{ID: "Function:helper", Label: scopeir.NodeFunction, Properties: graph.NodeProperties{
		"name":                                "emitUnresolvedReference",
		"filePath":                            "internal/resolution/emit.go",
		semantic.AppLayerProperty:             string(semantic.AppLayerBackend),
		semantic.AppLayerSourceProperty:       "path_rule",
		semantic.FunctionalAreaProperty:       string(semantic.FunctionalAreaResolution),
		semantic.FunctionalAreaSourceProperty: "path_rule",
	}})
	g.AddNode(graph.Node{ID: "ResolutionGap:site-stop", Label: scopeir.NodeResolutionGap, Properties: graph.NodeProperties{
		"name":                                "stop",
		"gapKind":                             graphhealth.ResolutionGapKindUnresolvedCall,
		"sourceSiteId":                        "site-stop",
		"sourceNodeId":                        "Function:source",
		"sourceNodeLabel":                     string(scopeir.NodeFunction),
		"sourceAppLayer":                      string(semantic.AppLayerBackend),
		"sourceFunctionalArea":                string(semantic.FunctionalAreaResolution),
		"factFamily":                          "call",
		"targetText":                          "stop",
		"targetRole":                          "callable",
		"sourceSiteStatus":                    "unresolved_local_binding",
		"proofKind":                           "none",
		"classification":                      graphhealth.DiagnosticClassificationInRepoUnresolved,
		"actionability":                       graphhealth.DiagnosticActionabilityAnalyzerGap,
		"resolutionSource":                    "scope-resolution",
		"filePath":                            "internal/resolution/resolve.go",
		"startLine":                           12,
		"startCol":                            3,
		"endLine":                             12,
		"endCol":                              9,
		"count":                               2,
		semantic.AppLayerProperty:             string(semantic.AppLayerBackend),
		semantic.AppLayerSourceProperty:       "resolution_gap_source_node",
		semantic.FunctionalAreaProperty:       string(semantic.FunctionalAreaResolution),
		semantic.FunctionalAreaSourceProperty: "resolution_gap_source_node",
	}})
	g.AddRelationship(graph.Relationship{
		ID:       "rel:step",
		SourceID: "Function:source",
		TargetID: "Process:resolution",
		Type:     graph.RelStepInProcess,
		Step:     &step,
	})
	g.AddRelationship(graph.Relationship{
		ID:               "rel:calls",
		SourceID:         "Function:source",
		TargetID:         "Function:helper",
		Type:             graph.RelCalls,
		Confidence:       0.95,
		ResolutionSource: "scope-resolution",
		SourceSiteID:     "site-helper",
		SourceSiteCount:  1,
		SourceSiteStatus: "resolved",
		ProofKind:        "local-binding",
		TargetRole:       "callable",
		TargetText:       "emitUnresolvedReference",
	})
	g.AddRelationship(graph.Relationship{
		ID:               "rel:gap",
		SourceID:         "Function:source",
		TargetID:         "ResolutionGap:site-stop",
		Type:             graph.RelHasResolutionGap,
		Confidence:       1,
		SourceSiteID:     "site-stop",
		SourceSiteCount:  2,
		SourceSiteStatus: "unresolved_local_binding",
		ProofKind:        "none",
		TargetRole:       "callable",
		TargetText:       "stop",
		FilePath:         "internal/resolution/resolve.go",
		StartLine:        12,
		StartCol:         3,
		EndLine:          12,
		EndCol:           9,
	})
	writeMCPGraphTB(t, repoPath, g)

	server := NewServer(Config{Store: store})
	payload, err := server.contextTool(map[string]any{"repo": "fixture", "uid": "Function:source"})
	if err != nil {
		t.Fatalf("contextTool: %v", err)
	}
	status, ok := payload["semanticStatus"].(semantic.GraphStatus)
	if !ok || status.AppLayer.Status != semantic.StatusComplete || status.FunctionalArea.Status != semantic.StatusComplete {
		t.Fatalf("context payload missing complete semanticStatus: %#v", payload["semanticStatus"])
	}
	symbol := payload["symbol"].(map[string]any)
	if symbol["type"] != string(scopeir.NodeFunction) ||
		symbol["appLayer"] != string(semantic.AppLayerBackend) ||
		symbol["functionalArea"] != string(semantic.FunctionalAreaResolution) ||
		symbol["topologyStatus"] != string(graphhealth.TopologyConnected) ||
		symbol["resolutionConfidence"] != graphhealth.ResolutionConfidenceDegraded ||
		symbol["resolutionGapCount"] != 2 {
		t.Fatalf("symbol lost semantic fields: %#v", symbol)
	}

	outgoing := payload["outgoing"].(map[string][]map[string]any)
	calls := outgoing["calls"]
	if len(calls) != 1 || calls[0]["uid"] != "Function:helper" ||
		calls[0]["appLayer"] != string(semantic.AppLayerBackend) ||
		calls[0]["sourceSiteStatus"] != "resolved" ||
		calls[0]["proofKind"] != "local-binding" {
		t.Fatalf("outgoing calls lost semantic/proof fields: %#v", calls)
	}

	gaps := payload["sourceResolutionGaps"].([]map[string]any)
	if len(gaps) != 1 {
		t.Fatalf("sourceResolutionGaps = %#v", gaps)
	}
	gap := gaps[0]
	if gap["uid"] != "ResolutionGap:site-stop" ||
		gap["resolutionRelation"] != "source_node_gap" ||
		gap["resolvedTarget"] != false ||
		gap["resolutionGapEntity"] != true ||
		gap["gapKind"] != graphhealth.ResolutionGapKindUnresolvedCall ||
		gap["targetText"] != "stop" ||
		gap["sourceSiteStatus"] != "unresolved_local_binding" ||
		gap["classification"] != graphhealth.DiagnosticClassificationInRepoUnresolved ||
		gap["actionability"] != graphhealth.DiagnosticActionabilityAnalyzerGap ||
		gap["sourceSiteCount"] != 2 ||
		gap["count"] != 2 {
		t.Fatalf("source gap lost semantic identity: %#v", gap)
	}

	processes := payload["processes"].([]map[string]any)
	if len(processes) != 1 ||
		processes[0]["type"] != string(scopeir.NodeProcess) ||
		processes[0]["appLayer"] != string(semantic.AppLayerBackend) ||
		processes[0]["functionalArea"] != string(semantic.FunctionalAreaResolution) {
		t.Fatalf("process row lost semantic fields: %#v", processes)
	}
}

func TestContextToolDistinguishesResolutionGapEntityFromSourceNode(t *testing.T) {
	store, repoPath := newMCPQueryBenchmarkRepo(t)
	g := graph.New()
	g.AddNode(graph.Node{ID: "Function:source", Label: scopeir.NodeFunction, Properties: graph.NodeProperties{
		"name":                                "resolveCall",
		"filePath":                            "internal/resolution/resolve.go",
		semantic.AppLayerProperty:             string(semantic.AppLayerBackend),
		semantic.AppLayerSourceProperty:       "path_rule",
		semantic.FunctionalAreaProperty:       string(semantic.FunctionalAreaResolution),
		semantic.FunctionalAreaSourceProperty: "path_rule",
	}})
	g.AddNode(graph.Node{ID: "ResolutionGap:site-stop", Label: scopeir.NodeResolutionGap, Properties: graph.NodeProperties{
		"name":                                "stop",
		"gapKind":                             graphhealth.ResolutionGapKindUnresolvedCall,
		"sourceSiteId":                        "site-stop",
		"sourceNodeId":                        "Function:source",
		"sourceAppLayer":                      string(semantic.AppLayerBackend),
		"sourceFunctionalArea":                string(semantic.FunctionalAreaResolution),
		"factFamily":                          "call",
		"targetText":                          "stop",
		"targetRole":                          "callable",
		"sourceSiteStatus":                    "unresolved_local_binding",
		"proofKind":                           "none",
		"classification":                      graphhealth.DiagnosticClassificationInRepoUnresolved,
		"actionability":                       graphhealth.DiagnosticActionabilityAnalyzerGap,
		"count":                               1,
		semantic.AppLayerProperty:             string(semantic.AppLayerBackend),
		semantic.AppLayerSourceProperty:       "resolution_gap_source_node",
		semantic.FunctionalAreaProperty:       string(semantic.FunctionalAreaResolution),
		semantic.FunctionalAreaSourceProperty: "resolution_gap_source_node",
	}})
	g.AddRelationship(graph.Relationship{
		ID:               "rel:gap",
		SourceID:         "Function:source",
		TargetID:         "ResolutionGap:site-stop",
		Type:             graph.RelHasResolutionGap,
		SourceSiteID:     "site-stop",
		SourceSiteCount:  1,
		SourceSiteStatus: "unresolved_local_binding",
		ProofKind:        "none",
		TargetRole:       "callable",
		TargetText:       "stop",
	})
	writeMCPGraphTB(t, repoPath, g)

	server := NewServer(Config{Store: store})
	payload, err := server.contextTool(map[string]any{"repo": "fixture", "uid": "ResolutionGap:site-stop"})
	if err != nil {
		t.Fatalf("contextTool: %v", err)
	}
	symbol := payload["symbol"].(map[string]any)
	if symbol["type"] != string(scopeir.NodeResolutionGap) ||
		symbol["resolutionGapEntity"] != true ||
		symbol["resolvedTarget"] != false ||
		symbol["resolutionRelation"] != "resolution_gap_entity" ||
		symbol["targetText"] != "stop" ||
		symbol["sourceNodeId"] != "Function:source" {
		t.Fatalf("gap symbol lost unresolved-entity identity: %#v", symbol)
	}
	sources := payload["resolutionGapSources"].([]map[string]any)
	if len(sources) != 1 ||
		sources[0]["uid"] != "Function:source" ||
		sources[0]["resolutionRelation"] != "gap_source_node" ||
		sources[0]["resolvedSourceNode"] != true ||
		sources[0]["appLayer"] != string(semantic.AppLayerBackend) ||
		sources[0]["sourceSiteStatus"] != "unresolved_local_binding" {
		t.Fatalf("resolutionGapSources lost source-node relation: %#v", sources)
	}
}

func TestContextToolWarnsForStaleIncompleteSemanticMetadata(t *testing.T) {
	store, repoPath := newMCPQueryBenchmarkRepo(t)
	g := graph.New()
	g.AddNode(graph.Node{ID: "Function:legacy", Label: scopeir.NodeFunction, Properties: graph.NodeProperties{
		"name":     "legacyContextTarget",
		"filePath": "internal/mcp/context.go",
	}})
	writeMCPGraphTB(t, repoPath, g)

	server := NewServer(Config{Store: store})
	payload, err := server.contextTool(map[string]any{"repo": "fixture", "uid": "Function:legacy"})
	if err != nil {
		t.Fatalf("contextTool: %v", err)
	}
	status, ok := payload["semanticStatus"].(semantic.GraphStatus)
	if !ok {
		t.Fatalf("context payload missing semanticStatus: %#v", payload)
	}
	if status.AppLayer.Status != semantic.StatusStaleIncomplete || status.FunctionalArea.Status != semantic.StatusStaleIncomplete {
		t.Fatalf("semantic status = %#v, want stale incomplete", status)
	}
	warning, ok := payload["semanticWarning"].(string)
	if !ok || warning == "" {
		t.Fatalf("context payload missing stale semantic warning: %#v", payload)
	}
}
