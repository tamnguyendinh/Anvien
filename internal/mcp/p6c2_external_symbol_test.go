package mcp

import (
	"testing"

	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
	"github.com/tamnguyendinh/anvien/internal/semantic"
)

func TestP6C2ContextImpactAndRenamePreserveExternalNonEditableSemantics(t *testing.T) {
	store, repoPath := newMCPQueryBenchmarkRepo(t)
	externalID := "ExternalSymbol:tsstdlib:math-max"
	g := graph.New()
	g.AddNode(graph.Node{ID: "Function:run", Label: scopeir.NodeFunction, Properties: graph.NodeProperties{
		"name":                                "run",
		"filePath":                            "src/app.ts",
		semantic.AppLayerProperty:             string(semantic.AppLayerBackend),
		semantic.AppLayerSourceProperty:       "test_fixture",
		semantic.FunctionalAreaProperty:       string(semantic.FunctionalAreaResolution),
		semantic.FunctionalAreaSourceProperty: "test_fixture",
	}})
	external := graph.Node{ID: externalID, Label: scopeir.NodeExternalSymbol, Properties: graph.NodeProperties{
		"name":                                "Math.max",
		"qualifiedName":                       "Math.max",
		"requestedNames":                      []string{"Math.max"},
		"requestedTargetTexts":                []string{"Math.max"},
		"meaning":                             "value",
		"meanings":                            []string{"value"},
		"semanticSymbolId":                    "tsstdlib:math-max",
		"semanticOwnerId":                     "tsstdlib:math",
		"authorityKind":                       "typescript_standard_library",
		"typeScriptVersion":                   "5.9.3",
		"catalogProofState":                   "ready",
		"authorityHash":                       "authority-hash",
		"catalogHash":                         "catalog-hash",
		"catalogArtifactHash":                 "artifact-hash",
		"profileHashes":                       []string{"profile-hash"},
		"configHashes":                        []string{"config-hash"},
		"declarationLibraries":                []string{"lib.es2015.core.d.ts"},
		"declarationRanges":                   []map[string]any{{"library": "lib.es2015.core.d.ts", "startLine": 120, "startCol": 5, "endLine": 120, "endCol": 48}},
		"sourceSiteIds":                       []string{"SourceSite:src/app.ts#call#Math.max#7#2#7#10"},
		"sourceSiteCount":                     1,
		"origin":                              "typescript_standard_library",
		"external":                            true,
		"editable":                            false,
		"repositoryOwned":                     false,
		semantic.AppLayerProperty:             string(semantic.AppLayerUnknown),
		semantic.AppLayerSourceProperty:       "external_authority",
		semantic.FunctionalAreaProperty:       string(semantic.FunctionalAreaResolution),
		semantic.FunctionalAreaSourceProperty: "external_authority",
	}}
	g.AddNode(external)
	g.AddRelationship(graph.Relationship{
		ID: "rel:CALLS:run->Math.max", SourceID: "Function:run", TargetID: externalID,
		Type: graph.RelCalls, Confidence: 1, ResolutionSource: "typescript_standard_library",
		SourceSiteID: "SourceSite:src/app.ts#call#Math.max#7#2#7#10", SourceSiteCount: 1,
		SourceSiteStatus: "resolved", ProofKind: "typescript-standard-library-authority",
		TargetRole: "callable", TargetText: "Math.max",
	})
	writeMCPGraphTB(t, repoPath, g)

	server := NewServer(Config{Store: store})
	contextPayload, err := server.contextTool(map[string]any{"repo": "fixture", "uid": externalID})
	if err != nil {
		t.Fatalf("context external symbol: %v", err)
	}
	symbol := contextPayload["symbol"].(map[string]any)
	if symbol["type"] != string(scopeir.NodeExternalSymbol) || symbol["externalSymbol"] != true ||
		symbol["external"] != true || symbol["editable"] != false || symbol["repositoryOwned"] != false ||
		symbol["semanticSymbolId"] != "tsstdlib:math-max" || symbol["semanticOwnerId"] != "tsstdlib:math" ||
		symbol["authorityKind"] != "typescript_standard_library" || symbol["typeScriptVersion"] != "5.9.3" ||
		symbol["catalogProofState"] != "ready" || symbol["filePath"] != "" {
		t.Fatalf("context lost explicit external provenance/non-editable semantics: %#v", symbol)
	}
	if declarations, ok := symbol["declarationRanges"].([]any); !ok || len(declarations) != 1 {
		t.Fatalf("context lost external declaration ranges: %#v", symbol["declarationRanges"])
	}

	impactPayload, _ := runImpactBFSProfiled(g, external, impactOptions{
		Direction: "upstream", MaxDepth: 1, RelationTypes: []string{string(graph.RelCalls)}, IncludeTests: true,
	}, false)
	impactTarget := impactPayload["target"].(map[string]any)
	if impactTarget["externalSymbol"] != true || impactTarget["external"] != true ||
		impactTarget["editable"] != false || impactTarget["repositoryOwned"] != false ||
		impactTarget["semanticSymbolId"] != "tsstdlib:math-max" {
		t.Fatalf("impact lost external target semantics: %#v", impactTarget)
	}

	renamePayload, err := server.renameTool(map[string]any{
		"repo": "fixture", "symbol_uid": externalID, "new_name": "maximum", "dry_run": true,
	})
	if err != nil {
		t.Fatalf("rename external symbol: %v", err)
	}
	if renamePayload["status"] != "rejected" || renamePayload["reason"] != "external_symbol_non_editable" ||
		renamePayload["symbol_uid"] != externalID {
		t.Fatalf("rename did not reject external target before edit collection: %#v", renamePayload)
	}
}
