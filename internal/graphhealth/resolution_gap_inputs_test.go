package graphhealth

import (
	"testing"

	"github.com/tamnguyendinh/avmatrix-go/internal/graph"
	"github.com/tamnguyendinh/avmatrix-go/internal/scopeir"
)

func TestSourceBackedResolutionGapInputsPreserveSourceSiteEvidence(t *testing.T) {
	g := graph.New()
	g.AddNode(graph.Node{
		ID:    "Function:src/app.ts:run",
		Label: "Function",
		Properties: graph.NodeProperties{
			"appLayer":       "backend",
			"functionalArea": "resolution",
			DiagnosticPropertyKey: []Diagnostic{
				{
					Kind:             DiagnosticUnresolvedReference,
					FactFamily:       "call",
					SourceNodeID:     "Function:src/app.ts:run",
					TargetText:       "stop",
					ResolutionSource: "scope-resolution",
					Classification:   DiagnosticClassificationInRepoUnresolved,
					Actionability:    DiagnosticActionabilityAnalyzerGap,
					FilePath:         "src/app.ts",
					FileHash:         "hash-1",
					StartLine:        10,
					StartCol:         4,
					EndLine:          10,
					EndCol:           10,
					SourceSiteID:     "SourceSite:src/app.ts#call#stop#10#4#10#10",
					SourceSiteStatus: "unresolved_local_binding",
					ProofKind:        "global-fallback-low-confidence",
					TargetRole:       "callable",
					Count:            2,
					Note:             "call target matched low-confidence global fallback only",
					Source:           "scope-resolution",
				},
				{
					Kind:             DiagnosticUnresolvedReference,
					FactFamily:       "access",
					SourceNodeID:     "Function:src/app.ts:run",
					TargetText:       "config.make",
					Classification:   DiagnosticClassificationInRepoUnresolved,
					Actionability:    DiagnosticActionabilityAnalyzerGap,
					FilePath:         "src/app.ts",
					StartLine:        12,
					SourceSiteID:     "SourceSite:src/app.ts#access#config.make#12#2#12#13",
					SourceSiteStatus: "unresolved_local_binding",
					ProofKind:        "none",
					TargetRole:       "member",
				},
				{
					Kind:         DiagnosticUnresolvedReference,
					FactFamily:   "type-reference",
					TargetText:   "MissingType",
					FilePath:     "src/app.ts",
					StartLine:    14,
					SourceSiteID: "SourceSite:src/app.ts#type-reference#MissingType#14#1#14#12",
					TargetRole:   "type",
				},
				{
					Kind:       DiagnosticUnresolvedReference,
					FactFamily: "call",
					TargetText: "missingSourceSite",
					FilePath:   "src/app.ts",
					StartLine:  16,
				},
			},
		},
	})

	inputs := SourceBackedResolutionGapInputs(g)
	if len(inputs) != 3 {
		t.Fatalf("SourceBackedResolutionGapInputs() len = %d, want 3: %#v", len(inputs), inputs)
	}
	call := requireResolutionGapInput(t, inputs, "call", "stop")
	if call.ID != "ResolutionGapInput:SourceSite:src/app.ts#call#stop#10#4#10#10" ||
		call.SourceSiteID != "SourceSite:src/app.ts#call#stop#10#4#10#10" ||
		call.SourceNodeID != "Function:src/app.ts:run" ||
		call.SourceNodeLabel != "Function" ||
		call.SourceAppLayer != "backend" ||
		call.SourceFunctionalArea != "resolution" ||
		call.FactFamily != "call" ||
		call.TargetText != "stop" ||
		call.TargetRole != "callable" ||
		call.SourceSiteStatus != "unresolved_local_binding" ||
		call.ProofKind != "global-fallback-low-confidence" ||
		call.Classification != DiagnosticClassificationInRepoUnresolved ||
		call.Actionability != DiagnosticActionabilityAnalyzerGap ||
		call.ResolutionSource != "scope-resolution" ||
		call.FilePath != "src/app.ts" ||
		call.FileHash != "hash-1" ||
		call.StartLine != 10 ||
		call.StartCol != 4 ||
		call.EndLine != 10 ||
		call.EndCol != 10 ||
		call.Count != 2 ||
		call.Note != "call target matched low-confidence global fallback only" {
		t.Fatalf("call gap input lost source-site evidence: %#v", call)
	}
	gapNode := call.GraphNode()
	if gapNode.ID != "ResolutionGap:SourceSite:src/app.ts#call#stop#10#4#10#10" ||
		gapNode.Label != "ResolutionGap" ||
		gapNode.Properties["gapKind"] != ResolutionGapKindUnresolvedCall ||
		gapNode.Properties["sourceSiteId"] != call.SourceSiteID ||
		gapNode.Properties["sourceNodeId"] != call.SourceNodeID ||
		gapNode.Properties["targetText"] != "stop" ||
		gapNode.Properties["targetRole"] != "callable" ||
		gapNode.Properties["classification"] != DiagnosticClassificationInRepoUnresolved ||
		gapNode.Properties["actionability"] != DiagnosticActionabilityAnalyzerGap ||
		gapNode.Properties["appLayer"] != "backend" ||
		gapNode.Properties["functionalArea"] != "resolution" ||
		gapNode.Properties["count"] != 2 {
		t.Fatalf("GraphNode() did not preserve persisted gap fields: %#v", gapNode)
	}
	gapRel := call.GraphRelationship()
	if gapRel.ID != "rel:has-resolution-gap:Function:src/app.ts:run->ResolutionGap:SourceSite:src/app.ts#call#stop#10#4#10#10" ||
		gapRel.SourceID != call.SourceNodeID ||
		gapRel.TargetID != gapNode.ID ||
		gapRel.Type != graph.RelHasResolutionGap ||
		gapRel.SourceSiteID != call.SourceSiteID ||
		gapRel.SourceSiteCount != 2 ||
		gapRel.SourceSiteStatus != "unresolved_local_binding" ||
		gapRel.ProofKind != "global-fallback-low-confidence" ||
		gapRel.TargetRole != "callable" ||
		gapRel.TargetText != "stop" ||
		len(gapRel.Evidence) != 1 ||
		gapRel.Evidence[0].Kind != "resolution_gap" {
		t.Fatalf("GraphRelationship() did not preserve persisted gap evidence: %#v", gapRel)
	}

	callAccess := SourceBackedCallAccessResolutionGapInputs(g)
	if len(callAccess) != 2 {
		t.Fatalf("SourceBackedCallAccessResolutionGapInputs() len = %d, want 2: %#v", len(callAccess), callAccess)
	}
	requireResolutionGapInput(t, callAccess, "call", "stop")
	requireResolutionGapInput(t, callAccess, "access", "config.make")
}

func requireResolutionGapInput(t *testing.T, inputs []ResolutionGapInput, factFamily string, targetText string) ResolutionGapInput {
	t.Helper()
	for _, input := range inputs {
		if input.FactFamily == factFamily && input.TargetText == targetText {
			return input
		}
	}
	t.Fatalf("missing resolution gap input factFamily=%q targetText=%q in %#v", factFamily, targetText, inputs)
	return ResolutionGapInput{}
}

func TestResolutionGapInputInfersTargetRole(t *testing.T) {
	tests := []struct {
		name           string
		input          ResolutionGapInput
		wantTargetRole string
		wantGapKind    string
	}{
		{
			name: "explicit target role wins",
			input: ResolutionGapInput{
				SourceSiteID:   "site-explicit",
				SourceNodeID:   "Function:src/run",
				FactFamily:     "call",
				TargetText:     "custom",
				TargetRole:     "custom_role",
				Classification: DiagnosticClassificationBuiltin,
			},
			wantTargetRole: "custom_role",
			wantGapKind:    ResolutionGapKindUnresolvedCall,
		},
		{
			name: "call fact family is callable",
			input: ResolutionGapInput{
				SourceSiteID: "site-call",
				SourceNodeID: "Function:src/run",
				FactFamily:   "call",
				TargetText:   "run",
			},
			wantTargetRole: "callable",
			wantGapKind:    ResolutionGapKindUnresolvedCall,
		},
		{
			name: "access fact family is member",
			input: ResolutionGapInput{
				SourceSiteID: "site-access",
				SourceNodeID: "Function:src/run",
				FactFamily:   "access",
				TargetText:   "config.value",
			},
			wantTargetRole: "member",
			wantGapKind:    ResolutionGapKindUnresolvedAccess,
		},
		{
			name: "type reference fact family is type",
			input: ResolutionGapInput{
				SourceSiteID: "site-type",
				SourceNodeID: "Function:src/run",
				FactFamily:   "type-reference",
				TargetText:   "MissingType",
			},
			wantTargetRole: "type",
			wantGapKind:    ResolutionGapKindUnresolvedTypeReference,
		},
		{
			name: "heritage fact family is type",
			input: ResolutionGapInput{
				SourceSiteID: "site-heritage",
				SourceNodeID: "Class:src/child",
				FactFamily:   "heritage",
				TargetText:   "Base",
			},
			wantTargetRole: "type",
			wantGapKind:    ResolutionGapKindUnresolvedHeritage,
		},
		{
			name: "builtin classification is builtin when fact family does not decide",
			input: ResolutionGapInput{
				SourceSiteID:   "site-builtin",
				SourceNodeID:   "Function:src/run",
				FactFamily:     "reference",
				TargetText:     "len",
				Classification: DiagnosticClassificationBuiltin,
			},
			wantTargetRole: "builtin",
			wantGapKind:    ResolutionGapKindUnresolvedReference,
		},
		{
			name: "standard library classification is builtin-like when fact family does not decide",
			input: ResolutionGapInput{
				SourceSiteID:   "site-stdlib",
				SourceNodeID:   "Function:src/run",
				FactFamily:     "reference",
				TargetText:     "strings.Contains",
				Classification: DiagnosticClassificationStandardLibrary,
			},
			wantTargetRole: "builtin",
			wantGapKind:    ResolutionGapKindUnresolvedReference,
		},
		{
			name: "test framework classification is test when fact family does not decide",
			input: ResolutionGapInput{
				SourceSiteID:   "site-test",
				SourceNodeID:   "Function:src/run",
				FactFamily:     "reference",
				TargetText:     "t.Fatalf",
				Classification: DiagnosticClassificationTestFramework,
			},
			wantTargetRole: "test",
			wantGapKind:    ResolutionGapKindUnresolvedReference,
		},
		{
			name: "external classification is external when fact family does not decide",
			input: ResolutionGapInput{
				SourceSiteID:   "site-external",
				SourceNodeID:   "Function:src/run",
				FactFamily:     "reference",
				TargetText:     "cobra.Command",
				Classification: DiagnosticClassificationExternalLibrary,
			},
			wantTargetRole: "external",
			wantGapKind:    ResolutionGapKindUnresolvedReference,
		},
		{
			name: "unknown fallback",
			input: ResolutionGapInput{
				SourceSiteID: "site-unknown",
				SourceNodeID: "Function:src/run",
				FactFamily:   "reference",
				TargetText:   "mystery",
			},
			wantTargetRole: "unknown",
			wantGapKind:    ResolutionGapKindUnresolvedReference,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.InferredTargetRole(); got != tt.wantTargetRole {
				t.Fatalf("InferredTargetRole() = %q, want %q", got, tt.wantTargetRole)
			}
			node := tt.input.GraphNode()
			if node.Label != scopeir.NodeResolutionGap {
				t.Fatalf("GraphNode() label = %q, want %q", node.Label, scopeir.NodeResolutionGap)
			}
			if got := node.Properties["targetRole"]; got != tt.wantTargetRole {
				t.Fatalf("GraphNode targetRole = %v, want %q", got, tt.wantTargetRole)
			}
			if got := node.Properties["gapKind"]; got != tt.wantGapKind {
				t.Fatalf("GraphNode gapKind = %v, want %q", got, tt.wantGapKind)
			}
			relationship := tt.input.GraphRelationship()
			if relationship.TargetRole != tt.wantTargetRole {
				t.Fatalf("GraphRelationship targetRole = %q, want %q", relationship.TargetRole, tt.wantTargetRole)
			}
			if relationship.Type != graph.RelHasResolutionGap {
				t.Fatalf("GraphRelationship type = %q, want %q", relationship.Type, graph.RelHasResolutionGap)
			}
		})
	}
}
