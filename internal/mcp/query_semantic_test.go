package mcp

import (
	"testing"

	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/graphhealth"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
	"github.com/tamnguyendinh/anvien/internal/semantic"
)

func TestQueryToolReturnsSemanticStatusAndResolutionGapSummaries(t *testing.T) {
	store, repoPath := newMCPQueryBenchmarkRepo(t)
	g := graph.New()
	g.AddNode(graph.Node{ID: "Process:resolution", Label: scopeir.NodeProcess, Properties: graph.NodeProperties{
		"heuristicLabel":                      "Resolution Reference Flow",
		"processType":                         "resolution",
		"stepCount":                           1,
		semantic.AppLayerProperty:             string(semantic.AppLayerBackend),
		semantic.AppLayerSourceProperty:       "process_member_inference",
		semantic.FunctionalAreaProperty:       string(semantic.FunctionalAreaResolution),
		semantic.FunctionalAreaSourceProperty: "process_member_inference",
	}})
	g.AddNode(graph.Node{ID: "Function:emitUnresolvedReference", Label: scopeir.NodeFunction, Properties: graph.NodeProperties{
		"name":                                "emitUnresolvedReference",
		"filePath":                            "internal/resolution/emit.go",
		"content":                             "emit unresolved reference diagnostic",
		"topologyStatus":                      string(graphhealth.TopologyConnected),
		"resolutionConfidence":                string(graphhealth.ResolutionConfidenceDegraded),
		"resolutionGapCount":                  2,
		semantic.AppLayerProperty:             string(semantic.AppLayerBackend),
		semantic.AppLayerSourceProperty:       "path_rule",
		semantic.FunctionalAreaProperty:       string(semantic.FunctionalAreaResolution),
		semantic.FunctionalAreaSourceProperty: "resolution_path",
	}})
	g.AddNode(graph.Node{ID: "ResolutionGap:site-stop", Label: scopeir.NodeResolutionGap, Properties: graph.NodeProperties{
		"name":                                "stop",
		"gapKind":                             graphhealth.ResolutionGapKindUnresolvedCall,
		"targetText":                          "stop",
		"classification":                      graphhealth.DiagnosticClassificationInRepoUnresolved,
		"actionability":                       graphhealth.DiagnosticActionabilityAnalyzerGap,
		"count":                               2,
		semantic.AppLayerProperty:             string(semantic.AppLayerBackend),
		semantic.AppLayerSourceProperty:       "source_node",
		semantic.FunctionalAreaProperty:       string(semantic.FunctionalAreaResolution),
		semantic.FunctionalAreaSourceProperty: "source_node",
	}})
	step := 1
	g.AddRelationship(graph.Relationship{
		ID:       "rel:step",
		SourceID: "Function:emitUnresolvedReference",
		TargetID: "Process:resolution",
		Type:     graph.RelStepInProcess,
		Step:     &step,
	})
	g.AddRelationship(graph.Relationship{
		ID:              "rel:gap",
		SourceID:        "Function:emitUnresolvedReference",
		TargetID:        "ResolutionGap:site-stop",
		Type:            graph.RelHasResolutionGap,
		SourceSiteCount: 2,
		TargetText:      "stop",
	})
	writeMCPGraphTB(t, repoPath, g)

	server := NewServer(Config{Store: store})
	payload, err := server.queryTool(map[string]any{"repo": "fixture", "query": "unresolved reference resolution", "limit": 5})
	if err != nil {
		t.Fatalf("queryTool: %v", err)
	}
	status, ok := payload["semanticStatus"].(semantic.GraphStatus)
	if !ok {
		t.Fatalf("query payload missing semanticStatus: %#v", payload)
	}
	if status.AppLayer.Status != semantic.StatusComplete || status.FunctionalArea.Status != semantic.StatusComplete {
		t.Fatalf("semantic status = %#v, want complete", status)
	}
	if _, ok := payload["semanticWarning"]; ok {
		t.Fatalf("fresh semantic payload should not warn: %#v", payload["semanticWarning"])
	}

	definitions := payload["definitions"].([]map[string]any)
	definition := findQueryRow(definitions, "Function:emitUnresolvedReference")
	if definition == nil {
		t.Fatalf("definitions missing semantic match: %#v", definitions)
	}
	if definition["type"] != string(scopeir.NodeFunction) ||
		definition["appLayer"] != string(semantic.AppLayerBackend) ||
		definition["functionalArea"] != string(semantic.FunctionalAreaResolution) ||
		definition["topologyStatus"] != string(graphhealth.TopologyConnected) ||
		definition["resolutionConfidence"] != string(graphhealth.ResolutionConfidenceDegraded) ||
		definition["resolutionGapCount"] != 2 {
		t.Fatalf("definition lost semantic fields: %#v", definition)
	}
	if kinds, ok := definition["resolutionGapKinds"].(map[string]int); !ok || kinds[graphhealth.ResolutionGapKindUnresolvedCall] != 2 {
		t.Fatalf("definition lost gap kinds: %#v", definition["resolutionGapKinds"])
	}
	if topTargets, ok := definition["resolutionGapTopTargets"].(map[string]int); !ok || topTargets["stop"] != 2 {
		t.Fatalf("definition lost gap targets: %#v", definition["resolutionGapTopTargets"])
	}

	symbols := payload["process_symbols"].([]map[string]any)
	symbol := findQueryRow(symbols, "Function:emitUnresolvedReference")
	if symbol == nil || symbol["appLayer"] != string(semantic.AppLayerBackend) || symbol["resolutionGapCount"] != 2 {
		t.Fatalf("process symbol lost semantic fields: %#v", symbols)
	}
}

func TestQueryToolWarnsForStaleIncompleteSemanticMetadata(t *testing.T) {
	store, repoPath := newMCPQueryBenchmarkRepo(t)
	g := graph.New()
	g.AddNode(graph.Node{ID: "Function:legacy", Label: scopeir.NodeFunction, Properties: graph.NodeProperties{
		"name":     "legacyQueryTarget",
		"filePath": "internal/mcp/tools.go",
	}})
	writeMCPGraphTB(t, repoPath, g)

	server := NewServer(Config{Store: store})
	payload, err := server.queryTool(map[string]any{"repo": "fixture", "query": "legacy query target", "limit": 5})
	if err != nil {
		t.Fatalf("queryTool: %v", err)
	}
	status, ok := payload["semanticStatus"].(semantic.GraphStatus)
	if !ok {
		t.Fatalf("query payload missing semanticStatus: %#v", payload)
	}
	if status.AppLayer.Status != semantic.StatusStaleIncomplete || status.FunctionalArea.Status != semantic.StatusStaleIncomplete {
		t.Fatalf("semantic status = %#v, want stale incomplete", status)
	}
	warning, ok := payload["semanticWarning"].(string)
	if !ok || warning == "" {
		t.Fatalf("query payload missing stale semantic warning: %#v", payload)
	}
}

func TestQueryToolRanksAIContextOwnersAndExplainsLanes(t *testing.T) {
	store, repoPath := newMCPQueryBenchmarkRepo(t)
	g := graph.New()
	g.AddNode(graph.Node{ID: "Function:contract-main", Label: scopeir.NodeFunction, Properties: graph.NodeProperties{
		"name":                                "main",
		"filePath":                            "cmd/generate-web-contracts/main.go",
		"content":                             "generated web contracts",
		semantic.AppLayerProperty:             "generated_contract",
		semantic.AppLayerSourceProperty:       "path_rule",
		semantic.FunctionalAreaProperty:       "contracts",
		semantic.FunctionalAreaSourceProperty: "path_rule",
	}})
	g.AddNode(graph.Node{ID: "Function:generate-ai-context", Label: scopeir.NodeFunction, Properties: graph.NodeProperties{
		"name":                                "GenerateAIContextFiles",
		"filePath":                            "internal/aicontext/aicontext.go",
		"content":                             "generate AGENTS.md CLAUDE.md and Anvien skills",
		semantic.AppLayerProperty:             string(semantic.AppLayerBackend),
		semantic.AppLayerSourceProperty:       "path_rule",
		semantic.FunctionalAreaProperty:       string(semantic.FunctionalAreaCLI),
		semantic.FunctionalAreaSourceProperty: "path_rule",
	}})
	g.AddNode(graph.Node{ID: "Function:install-base-skills", Label: scopeir.NodeFunction, Properties: graph.NodeProperties{
		"name":                                "installBaseSkills",
		"filePath":                            "internal/aicontext/aicontext.go",
		"content":                             "install generated .claude skills anvien skill files",
		semantic.AppLayerProperty:             string(semantic.AppLayerBackend),
		semantic.AppLayerSourceProperty:       "path_rule",
		semantic.FunctionalAreaProperty:       string(semantic.FunctionalAreaCLI),
		semantic.FunctionalAreaSourceProperty: "path_rule",
	}})
	g.AddNode(graph.Node{ID: "File:internal/aicontext/skills/anvien-planner.md", Label: scopeir.NodeFile, Properties: graph.NodeProperties{
		"name":                                "anvien-planner.md",
		"filePath":                            "internal/aicontext/skills/anvien-planner.md",
		semantic.AppLayerProperty:             string(semantic.AppLayerBackend),
		semantic.AppLayerSourceProperty:       "path_rule",
		semantic.FunctionalAreaProperty:       string(semantic.FunctionalAreaCLI),
		semantic.FunctionalAreaSourceProperty: "path_rule",
	}})
	writeMCPGraphTB(t, repoPath, g)

	server := NewServer(Config{Store: store})
	payload, err := server.queryTool(map[string]any{
		"repo":    "fixture",
		"query":   "generated Anvien skills AGENTS.md CLAUDE.md internal aicontext",
		"limit":   5,
		"explain": true,
	})
	if err != nil {
		t.Fatalf("queryTool: %v", err)
	}
	definitions := payload["definitions"].([]map[string]any)
	if len(definitions) < 2 {
		t.Fatalf("definitions = %#v", definitions)
	}
	first := definitions[0]
	if first["filePath"] != "internal/aicontext/aicontext.go" {
		t.Fatalf("first definition should be AI-context owner, got %#v", first)
	}
	if !queryTestStringSliceContains(first["queryLanes"], "docs_setup_ai_context_discovery") {
		t.Fatalf("first definition missing AI-context lane: %#v", first)
	}
	if !queryTestStringSliceContains(first["matchReasons"], "semanticSurface") {
		t.Fatalf("first definition missing semantic surface reason: %#v", first)
	}
	if !queryTestPayloadLane(payload["queryCapabilities"], "docs_setup_ai_context_discovery") {
		t.Fatalf("payload missing AI-context query capability: %#v", payload["queryCapabilities"])
	}
	if _, ok := payload["explain"].(map[string]any); !ok {
		t.Fatalf("explain payload missing: %#v", payload)
	}
}

func findQueryRow(rows []map[string]any, id string) map[string]any {
	for _, row := range rows {
		if row["id"] == id {
			return row
		}
	}
	return nil
}

func queryTestStringSliceContains(raw any, want string) bool {
	switch values := raw.(type) {
	case []string:
		for _, value := range values {
			if value == want {
				return true
			}
		}
	case []any:
		for _, value := range values {
			if text, ok := value.(string); ok && text == want {
				return true
			}
		}
	}
	return false
}

func queryTestPayloadLane(raw any, want string) bool {
	lanes, ok := raw.([]map[string]any)
	if !ok {
		return false
	}
	for _, lane := range lanes {
		if lane["id"] == want {
			return true
		}
	}
	return false
}
