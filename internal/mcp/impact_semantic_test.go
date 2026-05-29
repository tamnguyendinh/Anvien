package mcp

import (
	"strings"
	"testing"

	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/graphhealth"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
	"github.com/tamnguyendinh/anvien/internal/semantic"
)

func TestImpactToolReturnsSemanticAffectedLayerAndResolutionRiskSummary(t *testing.T) {
	g := graph.New()
	target := graph.Node{ID: "Function:Target", Label: scopeir.NodeFunction, Properties: impactSemanticProps(
		"Target", "internal/resolution/target.go", string(semantic.AppLayerBackend), string(semantic.FunctionalAreaResolution), nil,
	)}
	callerBackend := graph.Node{ID: "Function:BackendCaller", Label: scopeir.NodeFunction, Properties: impactSemanticProps(
		"BackendCaller", "internal/resolution/caller.go", string(semantic.AppLayerBackend), string(semantic.FunctionalAreaResolution), graph.NodeProperties{
			"topologyStatus":          string(graphhealth.TopologyConnected),
			"resolutionConfidence":    graphhealth.ResolutionConfidenceDegraded,
			"resolutionGapCount":      3,
			"resolutionHealthBuckets": map[string]int{string(graphhealth.ResolutionHealthUnresolvedCallTarget): 3, string(graphhealth.ResolutionHealthInRepoAnalyzerGap): 3},
		},
	)}
	callerFrontend := graph.Node{ID: "Function:FrontendCaller", Label: scopeir.NodeFunction, Properties: impactSemanticProps(
		"FrontendCaller", "anvien-web/src/GraphCanvas.tsx", string(semantic.AppLayerFrontend), string(semantic.FunctionalAreaWebGraphUI), graph.NodeProperties{
			"topologyStatus":       string(graphhealth.TopologyConnected),
			"resolutionConfidence": graphhealth.ResolutionConfidenceClear,
		},
	)}
	process := graph.Node{ID: "Process:ImpactFlow", Label: scopeir.NodeProcess, Properties: impactSemanticProps(
		"ImpactFlow", "", string(semantic.AppLayerBackend), string(semantic.FunctionalAreaResolution), graph.NodeProperties{
			"heuristicLabel": "ImpactFlow",
			"processType":    "cross_community",
		},
	)}
	module := graph.Node{ID: "Community:Resolution", Label: scopeir.NodeCommunity, Properties: impactSemanticProps(
		"Resolution", "", string(semantic.AppLayerBackend), string(semantic.FunctionalAreaResolution), graph.NodeProperties{"heuristicLabel": "Resolution"},
	)}
	g.AddNode(target)
	g.AddNode(callerBackend)
	g.AddNode(callerFrontend)
	g.AddNode(process)
	g.AddNode(module)
	step := 2
	g.AddRelationship(graph.Relationship{
		ID:               "rel:backend-target",
		SourceID:         callerBackend.ID,
		TargetID:         target.ID,
		Type:             graph.RelCalls,
		Confidence:       1,
		SourceSiteID:     "site-backend-target",
		SourceSiteStatus: "resolved",
		ProofKind:        "local-binding",
		TargetRole:       "callable",
		TargetText:       "Target",
	})
	g.AddRelationship(graph.Relationship{ID: "rel:frontend-target", SourceID: callerFrontend.ID, TargetID: target.ID, Type: graph.RelCalls, Confidence: 1})
	g.AddRelationship(graph.Relationship{ID: "rel:backend-process", SourceID: callerBackend.ID, TargetID: process.ID, Type: graph.RelStepInProcess, Step: &step})
	g.AddRelationship(graph.Relationship{ID: "rel:backend-module", SourceID: callerBackend.ID, TargetID: module.ID, Type: graph.RelMemberOf})

	payload, _ := runImpactBFSProfiled(g, target, impactOptions{
		Direction:     "upstream",
		MaxDepth:      1,
		RelationTypes: []string{string(graph.RelCalls)},
		IncludeTests:  true,
	}, false)

	status := payload["semanticStatus"].(semantic.GraphStatus)
	if status.AppLayer.Status != semantic.StatusComplete || status.FunctionalArea.Status != semantic.StatusComplete {
		t.Fatalf("semanticStatus = %#v, want complete", status)
	}
	targetPayload := payload["target"].(map[string]any)
	if targetPayload["type"] != string(scopeir.NodeFunction) ||
		targetPayload["appLayer"] != string(semantic.AppLayerBackend) ||
		targetPayload["functionalArea"] != string(semantic.FunctionalAreaResolution) {
		t.Fatalf("target lost semantic fields: %#v", targetPayload)
	}
	appLayers := payload["affectedAppLayers"].(map[string]int)
	if appLayers[string(semantic.AppLayerBackend)] != 1 || appLayers[string(semantic.AppLayerFrontend)] != 1 {
		t.Fatalf("affectedAppLayers = %#v", appLayers)
	}
	areas := payload["affectedFunctionalAreas"].(map[string]int)
	if areas[string(semantic.FunctionalAreaResolution)] != 1 || areas[string(semantic.FunctionalAreaWebGraphUI)] != 1 {
		t.Fatalf("affectedFunctionalAreas = %#v", areas)
	}

	byDepth := payload["byDepth"].(map[string][]map[string]any)
	if len(byDepth["1"]) != 2 {
		t.Fatalf("byDepth[1] = %#v", byDepth["1"])
	}
	backendRow := findImpactRow(byDepth["1"], callerBackend.ID)
	if backendRow == nil ||
		backendRow["appLayer"] != string(semantic.AppLayerBackend) ||
		backendRow["functionalArea"] != string(semantic.FunctionalAreaResolution) ||
		backendRow["sourceSiteStatus"] != "resolved" ||
		backendRow["proofKind"] != "local-binding" {
		t.Fatalf("backend row lost semantic/proof fields: %#v", backendRow)
	}

	risks := payload["resolutionHealthRisks"].(map[string]any)
	if risks["nodesWithGaps"] != 1 || risks["degradedNodes"] != 1 || risks["totalResolutionGapCount"] != 3 {
		t.Fatalf("resolutionHealthRisks = %#v", risks)
	}
	buckets := risks["resolutionHealthBuckets"].(map[string]int)
	if buckets[string(graphhealth.ResolutionHealthUnresolvedCallTarget)] != 3 ||
		buckets[string(graphhealth.ResolutionHealthInRepoAnalyzerGap)] != 3 {
		t.Fatalf("risk buckets = %#v", buckets)
	}

	processes := payload["affected_processes"].([]map[string]any)
	if len(processes) != 1 ||
		processes[0]["type"] != string(scopeir.NodeProcess) ||
		processes[0]["processType"] != "cross_community" ||
		processes[0]["appLayer"] != string(semantic.AppLayerBackend) ||
		processes[0]["functionalArea"] != string(semantic.FunctionalAreaResolution) {
		t.Fatalf("affected processes lost semantic fields: %#v", processes)
	}
	modules := payload["affected_modules"].([]map[string]any)
	if len(modules) != 1 ||
		modules[0]["type"] != string(scopeir.NodeCommunity) ||
		modules[0]["appLayer"] != string(semantic.AppLayerBackend) ||
		modules[0]["functionalArea"] != string(semantic.FunctionalAreaResolution) {
		t.Fatalf("affected modules lost semantic fields: %#v", modules)
	}
}

func TestImpactHighCriticalRiskWarningKeepsInspectionOutput(t *testing.T) {
	g := graph.New()
	target := graph.Node{ID: "Function:Target", Label: scopeir.NodeFunction, Properties: impactSemanticProps(
		"Target", "src/target.go", string(semantic.AppLayerBackend), string(semantic.FunctionalAreaResolution), nil,
	)}
	g.AddNode(target)
	for i := 0; i < 5; i++ {
		callerID := "Function:Caller" + string(rune('A'+i))
		processID := "Process:Flow" + string(rune('A'+i))
		caller := graph.Node{ID: callerID, Label: scopeir.NodeFunction, Properties: impactSemanticProps(
			callerID, "src/caller.go", string(semantic.AppLayerBackend), string(semantic.FunctionalAreaResolution), nil,
		)}
		process := graph.Node{ID: processID, Label: scopeir.NodeProcess, Properties: impactSemanticProps(
			processID, "", string(semantic.AppLayerBackend), string(semantic.FunctionalAreaResolution), graph.NodeProperties{"heuristicLabel": processID},
		)}
		g.AddNode(caller)
		g.AddNode(process)
		step := i + 1
		g.AddRelationship(graph.Relationship{ID: "rel:call:" + callerID, SourceID: callerID, TargetID: target.ID, Type: graph.RelCalls, Confidence: 1})
		g.AddRelationship(graph.Relationship{ID: "rel:process:" + callerID, SourceID: callerID, TargetID: processID, Type: graph.RelStepInProcess, Step: &step})
	}

	payload, _ := runImpactBFSProfiled(g, target, impactOptions{
		Direction:     "upstream",
		MaxDepth:      1,
		RelationTypes: []string{string(graph.RelCalls)},
		IncludeTests:  true,
	}, false)
	if payload["risk"] != "CRITICAL" {
		t.Fatalf("risk = %#v, payload %#v", payload["risk"], payload)
	}
	warning, ok := payload["workflowWarning"].(string)
	if !ok || !strings.Contains(warning, "workflow safety information, not a blocker") {
		t.Fatalf("workflowWarning = %#v", payload["workflowWarning"])
	}
	if payload["workflowWarningBlocksOutput"] != false {
		t.Fatalf("workflow warning should not block output: %#v", payload["workflowWarningBlocksOutput"])
	}
	if payload["impactedCount"] != 5 || len(payload["byDepth"].(map[string][]map[string]any)["1"]) != 5 {
		t.Fatalf("critical warning lost inspection output: %#v", payload)
	}
}

func impactSemanticProps(name string, filePath string, appLayer string, functionalArea string, extra graph.NodeProperties) graph.NodeProperties {
	props := graph.NodeProperties{
		"name":                                name,
		"filePath":                            filePath,
		semantic.AppLayerProperty:             appLayer,
		semantic.AppLayerSourceProperty:       "test_fixture",
		semantic.FunctionalAreaProperty:       functionalArea,
		semantic.FunctionalAreaSourceProperty: "test_fixture",
	}
	for key, value := range extra {
		props[key] = value
	}
	return props
}

func findImpactRow(rows []map[string]any, id string) map[string]any {
	for _, row := range rows {
		if row["id"] == id {
			return row
		}
	}
	return nil
}
