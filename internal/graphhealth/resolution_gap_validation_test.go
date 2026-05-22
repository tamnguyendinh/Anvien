package graphhealth

import (
	"testing"

	"github.com/tamnguyendinh/avmatrix-go/internal/graph"
	"github.com/tamnguyendinh/avmatrix-go/internal/scopeir"
)

func TestValidateResolutionGapPersistenceAcceptsSourceBackedGaps(t *testing.T) {
	g := graph.New()
	source := graph.Node{
		ID:    "Function:src/run",
		Label: scopeir.NodeFunction,
		Properties: graph.NodeProperties{
			"name":           "run",
			"appLayer":       "backend",
			"functionalArea": "resolution",
		},
	}
	g.AddNode(source)
	input := ResolutionGapInput{
		SourceSiteID:     "SourceSite:src/run.go#call#stop#10#4#10#10",
		SourceNodeID:     source.ID,
		SourceNodeLabel:  string(source.Label),
		SourceAppLayer:   "backend",
		FactFamily:       "call",
		TargetText:       "stop",
		SourceSiteStatus: "unresolved_local_binding",
		ProofKind:        "global-fallback-low-confidence",
		Classification:   DiagnosticClassificationInRepoUnresolved,
		Actionability:    DiagnosticActionabilityAnalyzerGap,
		FilePath:         "src/run.go",
		Count:            1,
	}
	g.AddNode(input.GraphNode())
	g.AddRelationship(input.GraphRelationship())

	validation := ValidateResolutionGapPersistence(g)
	if !validation.OK() {
		t.Fatalf("ValidateResolutionGapPersistence() = %#v, want OK", validation)
	}
	if validation.ResolutionGapNodes != 1 ||
		validation.HasResolutionGapRelationships != 1 ||
		validation.CountedGapRelationshipCount != 0 {
		t.Fatalf("validation counts = %#v, want one non-counted gap relationship", validation)
	}
}

func TestValidateResolutionGapPersistenceFlagsFakeResolvedOrTopologyClaims(t *testing.T) {
	tests := []struct {
		name      string
		build     func() *graph.Graph
		assertion func(t *testing.T, validation ResolutionGapValidation)
	}{
		{
			name: "dangling gap relationship",
			build: func() *graph.Graph {
				g := graph.New()
				g.AddNode(graph.Node{ID: "Function:src/run", Label: scopeir.NodeFunction})
				g.AddRelationship(graph.Relationship{
					ID:       "rel:missing-gap",
					SourceID: "Function:src/run",
					TargetID: "ResolutionGap:missing",
					Type:     graph.RelHasResolutionGap,
				})
				return g
			},
			assertion: func(t *testing.T, validation ResolutionGapValidation) {
				t.Helper()
				if validation.DanglingGapRelationshipCount != 1 {
					t.Fatalf("dangling count = %d, want 1: %#v", validation.DanglingGapRelationshipCount, validation)
				}
			},
		},
		{
			name: "gap relationship target is not a gap node",
			build: func() *graph.Graph {
				g := graph.New()
				g.AddNode(graph.Node{ID: "Function:src/run", Label: scopeir.NodeFunction})
				g.AddNode(graph.Node{ID: "Function:src/stop", Label: scopeir.NodeFunction})
				g.AddRelationship(graph.Relationship{
					ID:       "rel:fake-gap",
					SourceID: "Function:src/run",
					TargetID: "Function:src/stop",
					Type:     graph.RelHasResolutionGap,
				})
				return g
			},
			assertion: func(t *testing.T, validation ResolutionGapValidation) {
				t.Helper()
				if validation.NonGapTargetRelationshipCount != 1 {
					t.Fatalf("non-gap target count = %d, want 1: %#v", validation.NonGapTargetRelationshipCount, validation)
				}
			},
		},
		{
			name: "gap node claims resolved target",
			build: func() *graph.Graph {
				g := graph.New()
				g.AddNode(graph.Node{ID: "Function:src/run", Label: scopeir.NodeFunction})
				g.AddNode(graph.Node{
					ID:    "ResolutionGap:site-call",
					Label: scopeir.NodeResolutionGap,
					Properties: graph.NodeProperties{
						"sourceSiteId":     "site-call",
						"resolutionStatus": "resolved",
						"resolvedTargetId": "Function:src/stop",
					},
				})
				g.AddRelationship(graph.Relationship{
					ID:           "rel:gap",
					SourceID:     "Function:src/run",
					TargetID:     "ResolutionGap:site-call",
					Type:         graph.RelHasResolutionGap,
					SourceSiteID: "site-call",
				})
				return g
			},
			assertion: func(t *testing.T, validation ResolutionGapValidation) {
				t.Helper()
				if validation.GapResolvedClaimCount != 1 {
					t.Fatalf("resolved claim count = %d, want 1: %#v", validation.GapResolvedClaimCount, validation)
				}
			},
		},
		{
			name: "same source site cannot be unresolved gap and resolved CALLS edge",
			build: func() *graph.Graph {
				g := graph.New()
				g.AddNode(graph.Node{ID: "Function:src/run", Label: scopeir.NodeFunction})
				g.AddNode(graph.Node{ID: "Function:src/stop", Label: scopeir.NodeFunction})
				g.AddNode(graph.Node{
					ID:    "ResolutionGap:site-call",
					Label: scopeir.NodeResolutionGap,
					Properties: graph.NodeProperties{
						"sourceSiteId": "site-call",
					},
				})
				g.AddRelationship(graph.Relationship{
					ID:           "rel:gap",
					SourceID:     "Function:src/run",
					TargetID:     "ResolutionGap:site-call",
					Type:         graph.RelHasResolutionGap,
					SourceSiteID: "site-call",
				})
				g.AddRelationship(graph.Relationship{
					ID:            "rel:calls",
					SourceID:      "Function:src/run",
					TargetID:      "Function:src/stop",
					Type:          graph.RelCalls,
					SourceSiteIDs: []string{"site-call"},
					ProofKind:     "scope-binding",
				})
				return g
			},
			assertion: func(t *testing.T, validation ResolutionGapValidation) {
				t.Helper()
				if validation.ProoflessResolvedSourceSiteEdges != 1 {
					t.Fatalf("resolved source-site overlap count = %d, want 1: %#v", validation.ProoflessResolvedSourceSiteEdges, validation)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validation := ValidateResolutionGapPersistence(tt.build())
			if validation.OK() {
				t.Fatalf("ValidateResolutionGapPersistence() = %#v, want violation", validation)
			}
			tt.assertion(t, validation)
		})
	}
}
