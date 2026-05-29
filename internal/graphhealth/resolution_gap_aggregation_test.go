package graphhealth

import (
	"testing"

	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
)

func TestResolutionGapAggregatesPreserveCountsSamplesAndDistributions(t *testing.T) {
	inputs := []ResolutionGapInput{
		{
			SourceSiteID:         "site-one",
			SourceNodeID:         "Function:src/run",
			SourceNodeLabel:      string(scopeir.NodeFunction),
			SourceAppLayer:       "backend",
			SourceFunctionalArea: "resolution",
			FactFamily:           "call",
			TargetText:           "missing",
			SourceSiteStatus:     "unresolved_local_binding",
			ProofKind:            "none",
			Classification:       DiagnosticClassificationInRepoUnresolved,
			Actionability:        DiagnosticActionabilityAnalyzerGap,
			ResolutionSource:     "scope-resolution",
			FilePath:             "src/run.go",
			FileHash:             "hash-one",
			StartLine:            10,
			StartCol:             2,
			EndLine:              10,
			EndCol:               9,
			Count:                3,
			Note:                 "same source/fact/file/note bucket",
		},
		{
			SourceSiteID:         "site-two",
			SourceNodeID:         "Function:src/run",
			SourceNodeLabel:      string(scopeir.NodeFunction),
			SourceAppLayer:       "api",
			SourceFunctionalArea: "api",
			FactFamily:           "call",
			TargetText:           "missing",
			SourceSiteStatus:     "unresolved_local_binding",
			ProofKind:            "none",
			Classification:       DiagnosticClassificationInRepoUnresolved,
			Actionability:        DiagnosticActionabilityAnalyzerGap,
			ResolutionSource:     "scope-resolution",
			FilePath:             "src/run.go",
			FileHash:             "hash-two",
			StartLine:            12,
			StartCol:             2,
			EndLine:              12,
			EndCol:               9,
			Count:                2,
			Note:                 "same source/fact/file/note bucket",
		},
		{
			SourceSiteID:         "site-three",
			SourceNodeID:         "Function:src/run",
			SourceNodeLabel:      string(scopeir.NodeFunction),
			SourceAppLayer:       "backend",
			SourceFunctionalArea: "resolution",
			FactFamily:           "call",
			TargetText:           "otherMissing",
			SourceSiteStatus:     "unresolved_local_binding",
			ProofKind:            "none",
			Classification:       DiagnosticClassificationInRepoUnresolved,
			Actionability:        DiagnosticActionabilityAnalyzerGap,
			ResolutionSource:     "scope-resolution",
			FilePath:             "src/run.go",
			Count:                1,
			Note:                 "same source/fact/file/note bucket",
		},
	}

	aggregates := ResolutionGapAggregates(inputs, ResolutionGapAggregationOptions{MaxSamples: 1})
	if len(aggregates) != 2 {
		t.Fatalf("ResolutionGapAggregates() len = %d, want 2: %#v", len(aggregates), aggregates)
	}
	missing := requireResolutionGapAggregate(t, aggregates, "missing")
	if missing.GapKind != ResolutionGapKindUnresolvedCall ||
		missing.TargetRole != "callable" ||
		missing.InputCount != 2 ||
		missing.OccurrenceCount != 5 ||
		missing.SourceSiteCount != 2 {
		t.Fatalf("missing aggregate counts/roles = %#v", missing)
	}
	if len(missing.SourceSiteIDs) != 2 ||
		missing.SourceSiteIDs[0] != "site-one" ||
		missing.SourceSiteIDs[1] != "site-two" {
		t.Fatalf("source site traceability = %#v, want site-one/site-two", missing.SourceSiteIDs)
	}
	if len(missing.Samples) != 1 ||
		missing.Samples[0].SourceSiteID != "site-one" ||
		missing.Samples[0].Count != 3 {
		t.Fatalf("samples = %#v, want capped representative sample without count loss", missing.Samples)
	}
	if missing.AppLayerCounts["backend"] != 3 || missing.AppLayerCounts["api"] != 2 {
		t.Fatalf("app layer distribution = %#v, want backend=3 api=2", missing.AppLayerCounts)
	}
	if missing.FunctionalAreaCounts["resolution"] != 3 || missing.FunctionalAreaCounts["api"] != 2 {
		t.Fatalf("functional area distribution = %#v, want resolution=3 api=2", missing.FunctionalAreaCounts)
	}
	if missing.FilePathCounts["src/run.go"] != 5 {
		t.Fatalf("file path distribution = %#v, want src/run.go=5", missing.FilePathCounts)
	}
	other := requireResolutionGapAggregate(t, aggregates, "otherMissing")
	if other.InputCount != 1 || other.OccurrenceCount != 1 || other.SourceSiteIDs[0] != "site-three" {
		t.Fatalf("target text identity collapsed into wrong aggregate: %#v", other)
	}
	if missing.BucketKey == other.BucketKey {
		t.Fatalf("different target texts produced identical bucket key %q", missing.BucketKey)
	}
}

func TestSourceBackedResolutionGapAggregatesUseGraphDiagnostics(t *testing.T) {
	g := graph.New()
	g.AddNode(graph.Node{
		ID:    "Function:src/run",
		Label: scopeir.NodeFunction,
		Properties: graph.NodeProperties{
			"appLayer":       "backend",
			"functionalArea": "resolution",
			DiagnosticPropertyKey: []Diagnostic{
				{
					Kind:             DiagnosticUnresolvedReference,
					FactFamily:       "access",
					SourceNodeID:     "Function:src/run",
					TargetText:       "config.value",
					SourceSiteID:     "site-access-one",
					SourceSiteStatus: "unresolved_local_binding",
					ProofKind:        "none",
					Classification:   DiagnosticClassificationInRepoUnresolved,
					Actionability:    DiagnosticActionabilityAnalyzerGap,
					FilePath:         "src/run.go",
					Count:            4,
					Note:             "unresolved member access",
				},
				{
					Kind:             DiagnosticUnresolvedReference,
					FactFamily:       "access",
					SourceNodeID:     "Function:src/run",
					TargetText:       "config.value",
					SourceSiteID:     "site-access-two",
					SourceSiteStatus: "unresolved_local_binding",
					ProofKind:        "none",
					Classification:   DiagnosticClassificationInRepoUnresolved,
					Actionability:    DiagnosticActionabilityAnalyzerGap,
					FilePath:         "src/run.go",
					Count:            1,
					Note:             "unresolved member access",
				},
			},
		},
	})

	aggregates := SourceBackedResolutionGapAggregates(g, ResolutionGapAggregationOptions{MaxSamples: 0})
	if len(aggregates) != 1 {
		t.Fatalf("SourceBackedResolutionGapAggregates() len = %d, want 1: %#v", len(aggregates), aggregates)
	}
	aggregate := aggregates[0]
	if aggregate.TargetText != "config.value" ||
		aggregate.TargetRole != "member" ||
		aggregate.GapKind != ResolutionGapKindUnresolvedAccess ||
		aggregate.OccurrenceCount != 5 ||
		aggregate.SourceSiteCount != 2 ||
		len(aggregate.Samples) != 2 {
		t.Fatalf("aggregate lost graph diagnostic evidence: %#v", aggregate)
	}
}

func requireResolutionGapAggregate(t *testing.T, aggregates []ResolutionGapAggregate, targetText string) ResolutionGapAggregate {
	t.Helper()
	for _, aggregate := range aggregates {
		if aggregate.TargetText == targetText {
			return aggregate
		}
	}
	t.Fatalf("missing aggregate targetText=%q in %#v", targetText, aggregates)
	return ResolutionGapAggregate{}
}
