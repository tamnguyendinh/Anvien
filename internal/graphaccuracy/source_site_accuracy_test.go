package graphaccuracy

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunSourceSiteAccuracyReportsProofInventory(t *testing.T) {
	repo := t.TempDir()
	graphPath := writeGraphFixture(t, repo, "source-site.json", sourceSiteAccuracyFixtureGraph())
	outPath := filepath.Join(repo, "source-site-report.json")

	result, err := RunSourceSiteAccuracy(SourceSiteAccuracyOptions{
		GraphPath:   graphPath,
		OutPath:     outPath,
		MaxExamples: 5,
	})
	if err != nil {
		t.Fatalf("RunSourceSiteAccuracy() error = %v", err)
	}
	if _, err := os.Stat(outPath); err != nil {
		t.Fatalf("source-site accuracy output missing: %v", err)
	}

	if result.ResolvedEdges.EdgesByType["CALLS"] != 3 {
		t.Fatalf("CALLS count = %d, want 3", result.ResolvedEdges.EdgesByType["CALLS"])
	}
	if result.ResolvedEdges.EdgesByType["ACCESSES"] != 2 {
		t.Fatalf("ACCESSES count = %d, want 2", result.ResolvedEdges.EdgesByType["ACCESSES"])
	}
	if result.SourceSiteInventory.RelationshipOccurrences != 5 ||
		result.SourceSiteInventory.DiagnosticOccurrences != 2 ||
		result.SourceSiteInventory.AllOccurrences != 7 {
		t.Fatalf("source-site inventory = %#v", result.SourceSiteInventory)
	}
	if result.StatusCounts["resolved"] != 5 ||
		result.StatusCounts["unresolved_local_binding"] != 2 {
		t.Fatalf("status counts = %#v", result.StatusCounts)
	}
	if result.UnresolvedDiagnostics.LowConfidenceGlobalFallbackOccurrences != 2 {
		t.Fatalf("low-confidence diagnostics = %d, want 2", result.UnresolvedDiagnostics.LowConfidenceGlobalFallbackOccurrences)
	}
	if result.AccessesTargets.TargetLabelCounts["Property"] != 1 ||
		result.AccessesTargets.NonPropertyTargetCount != 1 {
		t.Fatalf("ACCESSES target metrics = %#v", result.AccessesTargets)
	}
	if result.Duplicates.DuplicatePairCount != 1 ||
		result.Duplicates.MaxDuplicate != 2 ||
		result.Duplicates.MergedRelationshipCount != 1 ||
		result.Duplicates.MergedSourceSiteOccurrenceCount != 2 {
		t.Fatalf("duplicate metrics = %#v", result.Duplicates)
	}
	if result.PolicyViolations.ResolvedEdgesWithoutProof != 1 ||
		result.PolicyViolations.ResolvedEdgesWithoutSourceSiteID != 1 ||
		result.PolicyViolations.NonPropertyAccessTargets != 1 ||
		result.PolicyViolations.CoarseFileSourceCallEdges != 1 {
		t.Fatalf("policy violations = %#v", result.PolicyViolations)
	}
	if result.PolicyViolations.FalseResolvedEdgeCandidates != 3 {
		t.Fatalf("false resolved edge candidates = %d, want 3", result.PolicyViolations.FalseResolvedEdgeCandidates)
	}
	if result.GoldenValidation.Enabled {
		t.Fatalf("graph inventory mode should not claim golden validation: %#v", result.GoldenValidation)
	}
}

func sourceSiteAccuracyFixtureGraph() GraphFile {
	sourceID := "Function:src/main.go:main"
	targetID := "Function:src/main.go:target"
	propertyID := "Property:src/main.go:value"
	constID := "Const:src/main.go:Limit"
	fileID := "File:src/main.go"

	nodes := []GraphNode{
		{
			ID:    fileID,
			Label: "File",
			Properties: map[string]any{
				"name":     "main.go",
				"filePath": "src/main.go",
			},
		},
		{
			ID:    sourceID,
			Label: "Function",
			Properties: map[string]any{
				"name":     "main",
				"filePath": "src/main.go",
				"graphHealthDiagnostics": []map[string]any{
					{
						"kind":             "unresolved_reference",
						"factFamily":       "call",
						"sourceNodeId":     sourceID,
						"targetText":       "stop",
						"filePath":         "src/main.go",
						"startLine":        8,
						"sourceSiteId":     "call:src/main.go:8:stop",
						"sourceSiteStatus": "unresolved_local_binding",
						"proofKind":        "global-fallback-low-confidence",
						"targetRole":       "callable",
						"classification":   "in_repo_unresolved",
						"actionability":    "analyzer_gap",
						"count":            2,
					},
				},
			},
		},
		{ID: targetID, Label: "Function", Properties: map[string]any{"name": "target", "filePath": "src/main.go"}},
		{ID: propertyID, Label: "Property", Properties: map[string]any{"name": "value", "filePath": "src/main.go"}},
		{ID: constID, Label: "Const", Properties: map[string]any{"name": "Limit", "filePath": "src/main.go"}},
	}
	relationships := []GraphRelationship{
		{
			ID:               "calls:main-target",
			Type:             "CALLS",
			SourceID:         sourceID,
			TargetID:         targetID,
			Reason:           "fixture",
			SourceSiteID:     "call:src/main.go:3:target",
			SourceSiteStatus: "resolved",
			ProofKind:        "scope-binding",
			TargetRole:       "callable",
			TargetText:       "target",
			FilePath:         "src/main.go",
			StartLine:        3,
		},
		{
			ID:               "calls:main-target-merged",
			Type:             "CALLS",
			SourceID:         sourceID,
			TargetID:         targetID,
			Reason:           "fixture",
			SourceSiteID:     "call:src/main.go:4:target",
			SourceSiteIDs:    []string{"call:src/main.go:4:target", "call:src/main.go:5:target"},
			SourceSiteCount:  2,
			SourceSiteStatus: "resolved",
			ProofKind:        "same-file",
			TargetRole:       "callable",
			TargetText:       "target",
			FilePath:         "src/main.go",
			StartLine:        4,
		},
		{
			ID:               "accesses:main-value",
			Type:             "ACCESSES",
			SourceID:         sourceID,
			TargetID:         propertyID,
			Reason:           "fixture",
			SourceSiteID:     "access:src/main.go:6:value",
			SourceSiteStatus: "resolved",
			ProofKind:        "receiver-member",
			TargetRole:       "member",
			TargetText:       "value",
			FilePath:         "src/main.go",
			StartLine:        6,
		},
		{
			ID:               "accesses:main-limit",
			Type:             "ACCESSES",
			SourceID:         sourceID,
			TargetID:         constID,
			Reason:           "fixture violation",
			SourceSiteID:     "access:src/main.go:7:Limit",
			SourceSiteStatus: "resolved",
			ProofKind:        "import-member",
			TargetRole:       "member",
			TargetText:       "Limit",
			FilePath:         "src/main.go",
			StartLine:        7,
		},
		{
			ID:       "calls:file-target",
			Type:     "CALLS",
			SourceID: fileID,
			TargetID: targetID,
			Reason:   "coarse fixture violation",
		},
	}
	return GraphFile{Nodes: nodes, Relationships: relationships}
}
