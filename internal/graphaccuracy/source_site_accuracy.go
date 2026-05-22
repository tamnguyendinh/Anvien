package graphaccuracy

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	sourceSiteRelCalls    = "CALLS"
	sourceSiteRelAccesses = "ACCESSES"
	sourceSiteRelUses     = "USES"
	sourceSiteRelInherits = "INHERITS"

	sourceSiteDiagnosticUnresolvedReference = "unresolved_reference"
	sourceSiteDiagnosticPropertyKey         = "graphHealthDiagnostics"

	sourceSiteStatusResolved = "resolved"
	sourceSiteStatusUnknown  = "unknown"

	sourceSiteProofMissing             = "missing"
	sourceSiteProofGlobalFallbackLow   = "global-fallback-low-confidence"
	sourceSiteAccessPropertyTarget     = "Property"
	sourceSiteMissingTargetLabel       = "missing_target"
	sourceSiteGoldenValidationDisabled = "golden fixture validation is disabled unless --golden is provided; policy violation counts are computed from the graph."
)

type SourceSiteAccuracyOptions struct {
	GraphPath   string
	GoldenPath  string
	OutPath     string
	MaxExamples int
}

type SourceSiteAccuracyResult struct {
	GeneratedAt string `json:"generatedAt"`
	Inputs      struct {
		Graph  string `json:"graph"`
		Golden string `json:"golden,omitempty"`
	} `json:"inputs"`
	Totals                SourceSiteAccuracyTotals         `json:"totals"`
	ResolvedEdges         SourceSiteResolvedEdgeMetrics    `json:"resolvedEdges"`
	UnresolvedDiagnostics SourceSiteDiagnosticMetrics      `json:"unresolvedDiagnostics"`
	SourceSiteInventory   SourceSiteInventoryMetrics       `json:"sourceSiteInventory"`
	StatusCounts          map[string]int                   `json:"statusCounts"`
	ProofKindCounts       map[string]int                   `json:"proofKindCounts"`
	FactFamilyCounts      map[string]int                   `json:"factFamilyCounts"`
	TargetRoleCounts      map[string]int                   `json:"targetRoleCounts"`
	AccessesTargets       SourceSiteAccessTargetMetrics    `json:"accessesTargets"`
	Duplicates            SourceSiteDuplicateMetrics       `json:"duplicates"`
	PolicyViolations      SourceSitePolicyViolationMetrics `json:"policyViolations"`
	GoldenValidation      SourceSiteGoldenValidation       `json:"goldenValidation"`
	Examples              SourceSiteAccuracyExamples       `json:"examples,omitempty"`
	Notes                 []string                         `json:"notes"`
}

type SourceSiteAccuracyTotals struct {
	Nodes         int `json:"nodes"`
	Relationships int `json:"relationships"`
}

type SourceSiteResolvedEdgeMetrics struct {
	EdgesByType                      map[string]int `json:"edgesByType"`
	SourceSiteEdgesByType            map[string]int `json:"sourceSiteEdgesByType"`
	SourceSiteOccurrencesByType      map[string]int `json:"sourceSiteOccurrencesByType"`
	EdgesWithoutProof                int            `json:"edgesWithoutProof"`
	EdgesWithoutSourceSiteID         int            `json:"edgesWithoutSourceSiteId"`
	LowConfidenceGlobalFallbackEdges int            `json:"lowConfidenceGlobalFallbackEdges"`
	CoarseFileSourceEdges            int            `json:"coarseFileSourceEdges"`
	CoarseFileTargetEdges            int            `json:"coarseFileTargetEdges"`
}

type SourceSiteDiagnosticMetrics struct {
	Buckets                                int            `json:"buckets"`
	Occurrences                            int            `json:"occurrences"`
	SourceSiteBuckets                      int            `json:"sourceSiteBuckets"`
	SourceSiteOccurrences                  int            `json:"sourceSiteOccurrences"`
	LowConfidenceGlobalFallbackOccurrences int            `json:"lowConfidenceGlobalFallbackOccurrences"`
	StatusCounts                           map[string]int `json:"statusCounts"`
	ProofKindCounts                        map[string]int `json:"proofKindCounts"`
	FactFamilyCounts                       map[string]int `json:"factFamilyCounts"`
	ClassificationCounts                   map[string]int `json:"classificationCounts"`
	ActionabilityCounts                    map[string]int `json:"actionabilityCounts"`
}

type SourceSiteInventoryMetrics struct {
	RelationshipBuckets     int `json:"relationshipBuckets"`
	RelationshipOccurrences int `json:"relationshipOccurrences"`
	DiagnosticBuckets       int `json:"diagnosticBuckets"`
	DiagnosticOccurrences   int `json:"diagnosticOccurrences"`
	AllBuckets              int `json:"allBuckets"`
	AllOccurrences          int `json:"allOccurrences"`
	StableIDOccurrences     int `json:"stableIdOccurrences"`
	MissingIDOccurrences    int `json:"missingIdOccurrences"`
	RelationshipsMissingID  int `json:"relationshipsMissingId"`
	DiagnosticsMissingID    int `json:"diagnosticsMissingId"`
}

type SourceSiteAccessTargetMetrics struct {
	TargetLabelCounts      map[string]int `json:"targetLabelCounts"`
	NonPropertyTargetCount int            `json:"nonPropertyTargetCount"`
	MissingTargetNodeCount int            `json:"missingTargetNodeCount"`
}

type SourceSiteDuplicateMetrics struct {
	DuplicatePairCount                           int                         `json:"duplicatePairCount"`
	DuplicateEdgeCount                           int                         `json:"duplicateEdgeCount"`
	MaxDuplicate                                 int                         `json:"maxDuplicate"`
	MergedRelationshipCount                      int                         `json:"mergedRelationshipCount"`
	MergedSourceSiteOccurrenceCount              int                         `json:"mergedSourceSiteOccurrenceCount"`
	MergedRelationshipsWithoutOccurrenceEvidence int                         `json:"mergedRelationshipsWithoutOccurrenceEvidence"`
	MergedRelationshipsWithIncompleteIDs         int                         `json:"mergedRelationshipsWithIncompleteIds"`
	Examples                                     []SourceSiteAccuracyExample `json:"examples,omitempty"`
}

type SourceSitePolicyViolationMetrics struct {
	FalseResolvedEdgeCandidates      int `json:"falseResolvedEdgeCandidates"`
	ResolvedEdgesWithoutProof        int `json:"resolvedEdgesWithoutProof"`
	ResolvedEdgesWithoutSourceSiteID int `json:"resolvedEdgesWithoutSourceSiteId"`
	LowConfidenceFallbackEdges       int `json:"lowConfidenceFallbackEdges"`
	NonPropertyAccessTargets         int `json:"nonPropertyAccessTargets"`
	CoarseFileSourceCallEdges        int `json:"coarseFileSourceCallEdges"`
	CoarseFileTargetCallEdges        int `json:"coarseFileTargetCallEdges"`
}

type SourceSiteGoldenValidation struct {
	Enabled                    bool                        `json:"enabled"`
	ExpectedSourceSites        int                         `json:"expectedSourceSites"`
	MatchedSourceSites         int                         `json:"matchedSourceSites"`
	SilentMissingSourceSites   int                         `json:"silentMissingSourceSites"`
	ExpectedFalseResolvedEdges int                         `json:"expectedFalseResolvedEdges"`
	FalseResolvedEdges         int                         `json:"falseResolvedEdges"`
	MissingSourceSiteIDs       []string                    `json:"missingSourceSiteIds,omitempty"`
	FalseResolvedEdgeExamples  []SourceSiteAccuracyExample `json:"falseResolvedEdgeExamples,omitempty"`
	MissingSourceSiteExamples  []SourceSiteAccuracyExample `json:"missingSourceSiteExamples,omitempty"`
	Note                       string                      `json:"note"`
}

type SourceSiteAccuracyExamples struct {
	NonPropertyAccessTargets       []SourceSiteAccuracyExample `json:"nonPropertyAccessTargets,omitempty"`
	ResolvedEdgesWithoutProof      []SourceSiteAccuracyExample `json:"resolvedEdgesWithoutProof,omitempty"`
	ResolvedEdgesWithoutSourceSite []SourceSiteAccuracyExample `json:"resolvedEdgesWithoutSourceSite,omitempty"`
	LowConfidenceFallbackEdges     []SourceSiteAccuracyExample `json:"lowConfidenceFallbackEdges,omitempty"`
	CoarseFileCallEdges            []SourceSiteAccuracyExample `json:"coarseFileCallEdges,omitempty"`
	UnresolvedDiagnostics          []SourceSiteAccuracyExample `json:"unresolvedDiagnostics,omitempty"`
}

type SourceSiteAccuracyExample struct {
	ID               string `json:"id,omitempty"`
	Type             string `json:"type,omitempty"`
	SourceID         string `json:"sourceId,omitempty"`
	SourceLabel      string `json:"sourceLabel,omitempty"`
	TargetID         string `json:"targetId,omitempty"`
	TargetLabel      string `json:"targetLabel,omitempty"`
	NodeID           string `json:"nodeId,omitempty"`
	FactFamily       string `json:"factFamily,omitempty"`
	SourceSiteID     string `json:"sourceSiteId,omitempty"`
	SourceSiteStatus string `json:"sourceSiteStatus,omitempty"`
	ProofKind        string `json:"proofKind,omitempty"`
	TargetRole       string `json:"targetRole,omitempty"`
	TargetText       string `json:"targetText,omitempty"`
	FilePath         string `json:"filePath,omitempty"`
	StartLine        int    `json:"startLine,omitempty"`
	Count            int    `json:"count,omitempty"`
	Reason           string `json:"reason,omitempty"`
}

type sourceSiteDiagnostic struct {
	Kind             string `json:"kind"`
	FactFamily       string `json:"factFamily,omitempty"`
	SourceNodeID     string `json:"sourceNodeId,omitempty"`
	TargetText       string `json:"targetText,omitempty"`
	ResolutionSource string `json:"resolutionSource,omitempty"`
	Classification   string `json:"classification,omitempty"`
	Actionability    string `json:"actionability,omitempty"`
	FilePath         string `json:"filePath,omitempty"`
	StartLine        int    `json:"startLine,omitempty"`
	StartCol         int    `json:"startCol,omitempty"`
	EndLine          int    `json:"endLine,omitempty"`
	EndCol           int    `json:"endCol,omitempty"`
	SourceSiteID     string `json:"sourceSiteId,omitempty"`
	SourceSiteStatus string `json:"sourceSiteStatus,omitempty"`
	ProofKind        string `json:"proofKind,omitempty"`
	TargetRole       string `json:"targetRole,omitempty"`
	Count            int    `json:"count,omitempty"`
	Note             string `json:"note,omitempty"`
}

type SourceSiteGoldenFixture struct {
	Name                  string                              `json:"name,omitempty"`
	ExpectedSourceSiteIDs []string                            `json:"expectedSourceSiteIds,omitempty"`
	FalseResolvedEdges    []SourceSiteGoldenFalseResolvedEdge `json:"falseResolvedEdges,omitempty"`
}

type SourceSiteGoldenFalseResolvedEdge struct {
	Type     string `json:"type"`
	SourceID string `json:"sourceId"`
	TargetID string `json:"targetId"`
	Reason   string `json:"reason,omitempty"`
}

func RunSourceSiteAccuracy(options SourceSiteAccuracyOptions) (SourceSiteAccuracyResult, error) {
	if options.GraphPath == "" {
		return SourceSiteAccuracyResult{}, fmt.Errorf("graph path is required")
	}
	if options.MaxExamples <= 0 {
		options.MaxExamples = 50
	}
	graphFile, err := ReadGraph(options.GraphPath)
	if err != nil {
		return SourceSiteAccuracyResult{}, err
	}
	result := buildSourceSiteAccuracy(options.GraphPath, graphFile, options.MaxExamples)
	if options.GoldenPath != "" {
		fixture, err := ReadSourceSiteGoldenFixture(options.GoldenPath)
		if err != nil {
			return SourceSiteAccuracyResult{}, err
		}
		applySourceSiteGoldenValidation(&result, graphFile, fixture, options.GoldenPath, options.MaxExamples)
	}
	if options.OutPath != "" {
		if err := WriteSourceSiteAccuracyResult(options.OutPath, result); err != nil {
			return SourceSiteAccuracyResult{}, err
		}
	}
	return result, nil
}

func WriteSourceSiteAccuracyResult(path string, result SourceSiteAccuracyResult) error {
	raw, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal source-site accuracy report: %w", err)
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, append(raw, '\n'), 0o644)
}

func ReadSourceSiteGoldenFixture(path string) (SourceSiteGoldenFixture, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return SourceSiteGoldenFixture{}, fmt.Errorf("read source-site golden fixture %s: %w", path, err)
	}
	var fixture SourceSiteGoldenFixture
	if err := json.Unmarshal(raw, &fixture); err != nil {
		return SourceSiteGoldenFixture{}, fmt.Errorf("decode source-site golden fixture %s: %w", path, err)
	}
	return fixture, nil
}

func SourceSiteAccuracySummaryLines(result SourceSiteAccuracyResult) []string {
	lines := []string{
		fmt.Sprintf("sourceSites.relationships=%d buckets=%d diagnostics=%d buckets=%d all=%d stableIDs=%d missingIDs=%d",
			result.SourceSiteInventory.RelationshipOccurrences,
			result.SourceSiteInventory.RelationshipBuckets,
			result.SourceSiteInventory.DiagnosticOccurrences,
			result.SourceSiteInventory.DiagnosticBuckets,
			result.SourceSiteInventory.AllOccurrences,
			result.SourceSiteInventory.StableIDOccurrences,
			result.SourceSiteInventory.MissingIDOccurrences,
		),
		fmt.Sprintf("resolved.calls=%d accesses=%d unresolvedDiagnostics=%d lowConfidenceFallbackDiagnostics=%d",
			result.ResolvedEdges.EdgesByType[sourceSiteRelCalls],
			result.ResolvedEdges.EdgesByType[sourceSiteRelAccesses],
			result.UnresolvedDiagnostics.Occurrences,
			result.UnresolvedDiagnostics.LowConfidenceGlobalFallbackOccurrences,
		),
		fmt.Sprintf("accesses.propertyTargets=%d nonPropertyTargets=%d missingTargets=%d",
			result.AccessesTargets.TargetLabelCounts[sourceSiteAccessPropertyTarget],
			result.AccessesTargets.NonPropertyTargetCount,
			result.AccessesTargets.MissingTargetNodeCount,
		),
		fmt.Sprintf("duplicates.pairs=%d duplicateEdges=%d maxDuplicate=%d mergedRelationships=%d mergedOccurrences=%d",
			result.Duplicates.DuplicatePairCount,
			result.Duplicates.DuplicateEdgeCount,
			result.Duplicates.MaxDuplicate,
			result.Duplicates.MergedRelationshipCount,
			result.Duplicates.MergedSourceSiteOccurrenceCount,
		),
		fmt.Sprintf("policy.falseResolvedEdgeCandidates=%d noProof=%d noSourceSiteID=%d lowConfidenceEdges=%d coarseFileCallEdges=%d",
			result.PolicyViolations.FalseResolvedEdgeCandidates,
			result.PolicyViolations.ResolvedEdgesWithoutProof,
			result.PolicyViolations.ResolvedEdgesWithoutSourceSiteID,
			result.PolicyViolations.LowConfidenceFallbackEdges,
			result.PolicyViolations.CoarseFileSourceCallEdges+result.PolicyViolations.CoarseFileTargetCallEdges,
		),
		fmt.Sprintf("golden.enabled=%t expectedSourceSites=%d matchedSourceSites=%d silentMissingSourceSites=%d expectedFalseResolvedEdges=%d falseResolvedEdges=%d",
			result.GoldenValidation.Enabled,
			result.GoldenValidation.ExpectedSourceSites,
			result.GoldenValidation.MatchedSourceSites,
			result.GoldenValidation.SilentMissingSourceSites,
			result.GoldenValidation.ExpectedFalseResolvedEdges,
			result.GoldenValidation.FalseResolvedEdges,
		),
	}
	return lines
}

func buildSourceSiteAccuracy(graphPath string, graphFile GraphFile, maxExamples int) SourceSiteAccuracyResult {
	nodesByID := make(map[string]GraphNode, len(graphFile.Nodes))
	for _, node := range graphFile.Nodes {
		nodesByID[node.ID] = node
	}

	var result SourceSiteAccuracyResult
	result.GeneratedAt = time.Now().Format(time.RFC3339)
	result.Inputs.Graph = graphPath
	result.Totals.Nodes = len(graphFile.Nodes)
	result.Totals.Relationships = len(graphFile.Relationships)
	result.ResolvedEdges.EdgesByType = map[string]int{}
	result.ResolvedEdges.SourceSiteEdgesByType = map[string]int{}
	result.ResolvedEdges.SourceSiteOccurrencesByType = map[string]int{}
	result.UnresolvedDiagnostics.StatusCounts = map[string]int{}
	result.UnresolvedDiagnostics.ProofKindCounts = map[string]int{}
	result.UnresolvedDiagnostics.FactFamilyCounts = map[string]int{}
	result.UnresolvedDiagnostics.ClassificationCounts = map[string]int{}
	result.UnresolvedDiagnostics.ActionabilityCounts = map[string]int{}
	result.StatusCounts = map[string]int{}
	result.ProofKindCounts = map[string]int{}
	result.FactFamilyCounts = map[string]int{}
	result.TargetRoleCounts = map[string]int{}
	result.AccessesTargets.TargetLabelCounts = map[string]int{}
	result.GoldenValidation.Note = sourceSiteGoldenValidationDisabled

	duplicatePairs := map[string]int{}
	duplicateExamples := map[string]SourceSiteAccuracyExample{}

	for _, relationship := range graphFile.Relationships {
		if isResolvedAccuracyRelationship(relationship.Type) {
			result.ResolvedEdges.EdgesByType[relationship.Type]++
			key := relationship.Type + "\x00" + relationship.SourceID + "\x00" + relationship.TargetID
			duplicatePairs[key]++
			if _, ok := duplicateExamples[key]; !ok {
				duplicateExamples[key] = relationshipSourceSiteExample(relationship, nodesByID, "duplicate source-target relationship pair")
			}
			checkResolvedRelationshipPolicy(relationship, nodesByID, &result, maxExamples)
		}

		occurrences := relationshipSourceSiteOccurrenceCount(relationship)
		if occurrences <= 0 {
			continue
		}
		result.SourceSiteInventory.RelationshipBuckets++
		result.SourceSiteInventory.RelationshipOccurrences += occurrences
		result.ResolvedEdges.SourceSiteEdgesByType[relationship.Type]++
		result.ResolvedEdges.SourceSiteOccurrencesByType[relationship.Type] += occurrences
		addCount(result.StatusCounts, normalizedSourceSiteStatus(relationship.SourceSiteStatus, sourceSiteStatusResolved), occurrences)
		addCount(result.ProofKindCounts, normalizedProofKind(relationship.ProofKind), occurrences)
		addCount(result.FactFamilyCounts, relationshipFactFamily(relationship.Type), occurrences)
		addCount(result.TargetRoleCounts, normalizedMapKey(relationship.TargetRole), occurrences)
		if relationshipHasStableSourceSiteID(relationship) {
			result.SourceSiteInventory.StableIDOccurrences += occurrences
		} else {
			result.SourceSiteInventory.MissingIDOccurrences += occurrences
			result.SourceSiteInventory.RelationshipsMissingID++
		}
		if occurrences > 1 {
			result.Duplicates.MergedRelationshipCount++
			result.Duplicates.MergedSourceSiteOccurrenceCount += occurrences
			if relationship.SourceSiteCount <= 1 && len(relationship.SourceSiteIDs) <= 1 {
				result.Duplicates.MergedRelationshipsWithoutOccurrenceEvidence++
			}
			if relationship.SourceSiteCount > 0 && len(relationship.SourceSiteIDs) > 0 && len(relationship.SourceSiteIDs) < relationship.SourceSiteCount {
				result.Duplicates.MergedRelationshipsWithIncompleteIDs++
			}
		}
	}

	for key, count := range duplicatePairs {
		if count <= 1 {
			continue
		}
		result.Duplicates.DuplicatePairCount++
		result.Duplicates.DuplicateEdgeCount += count
		if count > result.Duplicates.MaxDuplicate {
			result.Duplicates.MaxDuplicate = count
		}
		addSourceSiteExample(&result.Duplicates.Examples, duplicateExamples[key], maxExamples)
	}

	for _, node := range graphFile.Nodes {
		for _, diagnostic := range sourceSiteDiagnosticsFromNode(node) {
			if diagnostic.Kind != sourceSiteDiagnosticUnresolvedReference {
				continue
			}
			occurrences := sourceSiteDiagnosticOccurrenceCount(diagnostic)
			result.UnresolvedDiagnostics.Buckets++
			result.UnresolvedDiagnostics.Occurrences += occurrences
			addCount(result.UnresolvedDiagnostics.StatusCounts, normalizedSourceSiteStatus(diagnostic.SourceSiteStatus, sourceSiteStatusUnknown), occurrences)
			addCount(result.UnresolvedDiagnostics.ProofKindCounts, normalizedProofKind(diagnostic.ProofKind), occurrences)
			addCount(result.UnresolvedDiagnostics.FactFamilyCounts, normalizedMapKey(diagnostic.FactFamily), occurrences)
			addCount(result.UnresolvedDiagnostics.ClassificationCounts, normalizedMapKey(diagnostic.Classification), occurrences)
			addCount(result.UnresolvedDiagnostics.ActionabilityCounts, normalizedMapKey(diagnostic.Actionability), occurrences)
			addCount(result.StatusCounts, normalizedSourceSiteStatus(diagnostic.SourceSiteStatus, sourceSiteStatusUnknown), occurrences)
			addCount(result.ProofKindCounts, normalizedProofKind(diagnostic.ProofKind), occurrences)
			addCount(result.FactFamilyCounts, normalizedMapKey(diagnostic.FactFamily), occurrences)
			addCount(result.TargetRoleCounts, normalizedMapKey(diagnostic.TargetRole), occurrences)
			if diagnostic.SourceSiteID == "" {
				result.SourceSiteInventory.MissingIDOccurrences += occurrences
				result.SourceSiteInventory.DiagnosticsMissingID++
			} else {
				result.SourceSiteInventory.StableIDOccurrences += occurrences
				result.UnresolvedDiagnostics.SourceSiteBuckets++
				result.UnresolvedDiagnostics.SourceSiteOccurrences += occurrences
			}
			result.SourceSiteInventory.DiagnosticBuckets++
			result.SourceSiteInventory.DiagnosticOccurrences += occurrences
			if diagnostic.ProofKind == sourceSiteProofGlobalFallbackLow {
				result.UnresolvedDiagnostics.LowConfidenceGlobalFallbackOccurrences += occurrences
			}
			addSourceSiteExample(&result.Examples.UnresolvedDiagnostics, diagnosticSourceSiteExample(node, diagnostic), maxExamples)
		}
	}

	result.SourceSiteInventory.AllBuckets = result.SourceSiteInventory.RelationshipBuckets + result.SourceSiteInventory.DiagnosticBuckets
	result.SourceSiteInventory.AllOccurrences = result.SourceSiteInventory.RelationshipOccurrences + result.SourceSiteInventory.DiagnosticOccurrences
	result.PolicyViolations.FalseResolvedEdgeCandidates =
		result.PolicyViolations.ResolvedEdgesWithoutProof +
			result.PolicyViolations.LowConfidenceFallbackEdges +
			result.PolicyViolations.NonPropertyAccessTargets +
			result.PolicyViolations.CoarseFileSourceCallEdges +
			result.PolicyViolations.CoarseFileTargetCallEdges
	result.GoldenValidation.FalseResolvedEdges = 0
	result.GoldenValidation.SilentMissingSourceSites = 0
	result.Notes = []string{
		"Resolved CALLS/ACCESSES counts are read from final graph relationships.",
		"Source-site relationship occurrences use sourceSiteCount when present, then sourceSiteIds length, then sourceSiteId.",
		"Unresolved source-site occurrences are read from graphHealthDiagnostics entries with kind unresolved_reference.",
		"Policy violation counts are graph-inventory checks, not a replacement for golden corpus tests.",
		"Golden validation is disabled unless --golden is provided.",
	}
	return result
}

func applySourceSiteGoldenValidation(result *SourceSiteAccuracyResult, graphFile GraphFile, fixture SourceSiteGoldenFixture, goldenPath string, maxExamples int) {
	result.Inputs.Golden = goldenPath
	result.GoldenValidation.Enabled = true
	result.GoldenValidation.Note = fmt.Sprintf("golden fixture validation read from %s", filepath.ToSlash(goldenPath))

	sourceSiteIDs := graphSourceSiteIDs(graphFile)
	for _, id := range fixture.ExpectedSourceSiteIDs {
		id = strings.TrimSpace(id)
		if id == "" {
			continue
		}
		result.GoldenValidation.ExpectedSourceSites++
		if sourceSiteIDs[id] {
			result.GoldenValidation.MatchedSourceSites++
			continue
		}
		result.GoldenValidation.SilentMissingSourceSites++
		if maxExamples <= 0 || len(result.GoldenValidation.MissingSourceSiteIDs) < maxExamples {
			result.GoldenValidation.MissingSourceSiteIDs = append(result.GoldenValidation.MissingSourceSiteIDs, id)
		}
		addSourceSiteExample(&result.GoldenValidation.MissingSourceSiteExamples, SourceSiteAccuracyExample{
			SourceSiteID: id,
			Reason:       "expected source-site id is absent from graph relationships and unresolved diagnostics",
		}, maxExamples)
	}

	relationshipsByKey := graphRelationshipsByGoldenKey(graphFile)
	nodesByID := make(map[string]GraphNode, len(graphFile.Nodes))
	for _, node := range graphFile.Nodes {
		nodesByID[node.ID] = node
	}
	for _, expectedEdge := range fixture.FalseResolvedEdges {
		key := sourceSiteGoldenFalseEdgeKey(expectedEdge.Type, expectedEdge.SourceID, expectedEdge.TargetID)
		if key == "" {
			continue
		}
		result.GoldenValidation.ExpectedFalseResolvedEdges++
		relationship, ok := relationshipsByKey[key]
		if !ok {
			continue
		}
		result.GoldenValidation.FalseResolvedEdges++
		reason := strings.TrimSpace(expectedEdge.Reason)
		if reason == "" {
			reason = "golden fixture marks this resolved edge as false"
		}
		addSourceSiteExample(&result.GoldenValidation.FalseResolvedEdgeExamples, relationshipSourceSiteExample(relationship, nodesByID, reason), maxExamples)
	}

	result.Notes = append(result.Notes, "Golden validation compares expected source-site IDs and known-false resolved edges against the graph snapshot.")
}

func checkResolvedRelationshipPolicy(relationship GraphRelationship, nodesByID map[string]GraphNode, result *SourceSiteAccuracyResult, maxExamples int) {
	source := nodesByID[relationship.SourceID]
	target, targetOK := nodesByID[relationship.TargetID]
	if relationship.Type == sourceSiteRelAccesses {
		targetLabel := sourceSiteMissingTargetLabel
		if targetOK {
			targetLabel = target.Label
		} else {
			result.AccessesTargets.MissingTargetNodeCount++
		}
		result.AccessesTargets.TargetLabelCounts[targetLabel]++
		if targetLabel != sourceSiteAccessPropertyTarget {
			result.AccessesTargets.NonPropertyTargetCount++
			result.PolicyViolations.NonPropertyAccessTargets++
			addSourceSiteExample(&result.Examples.NonPropertyAccessTargets, relationshipSourceSiteExample(relationship, nodesByID, "ACCESSES target is not Property"), maxExamples)
		}
	}
	if normalizedProofKind(relationship.ProofKind) == sourceSiteProofMissing {
		result.ResolvedEdges.EdgesWithoutProof++
		result.PolicyViolations.ResolvedEdgesWithoutProof++
		addSourceSiteExample(&result.Examples.ResolvedEdgesWithoutProof, relationshipSourceSiteExample(relationship, nodesByID, "resolved relationship has no proofKind"), maxExamples)
	}
	if !relationshipHasStableSourceSiteID(relationship) {
		result.ResolvedEdges.EdgesWithoutSourceSiteID++
		result.PolicyViolations.ResolvedEdgesWithoutSourceSiteID++
		addSourceSiteExample(&result.Examples.ResolvedEdgesWithoutSourceSite, relationshipSourceSiteExample(relationship, nodesByID, "resolved relationship has no sourceSiteId/sourceSiteIds"), maxExamples)
	}
	if relationship.ProofKind == sourceSiteProofGlobalFallbackLow {
		result.ResolvedEdges.LowConfidenceGlobalFallbackEdges++
		result.PolicyViolations.LowConfidenceFallbackEdges++
		addSourceSiteExample(&result.Examples.LowConfidenceFallbackEdges, relationshipSourceSiteExample(relationship, nodesByID, "resolved relationship uses low-confidence global fallback proof"), maxExamples)
	}
	if relationship.Type == sourceSiteRelCalls {
		if source.Label == "File" {
			result.ResolvedEdges.CoarseFileSourceEdges++
			result.PolicyViolations.CoarseFileSourceCallEdges++
			addSourceSiteExample(&result.Examples.CoarseFileCallEdges, relationshipSourceSiteExample(relationship, nodesByID, "CALLS source is coarse File node"), maxExamples)
		}
		if target.Label == "File" {
			result.ResolvedEdges.CoarseFileTargetEdges++
			result.PolicyViolations.CoarseFileTargetCallEdges++
			addSourceSiteExample(&result.Examples.CoarseFileCallEdges, relationshipSourceSiteExample(relationship, nodesByID, "CALLS target is coarse File node"), maxExamples)
		}
	}
}

func sourceSiteDiagnosticsFromNode(node GraphNode) []sourceSiteDiagnostic {
	if node.Properties == nil {
		return nil
	}
	value, ok := node.Properties[sourceSiteDiagnosticPropertyKey]
	if !ok || value == nil {
		return nil
	}
	raw, err := json.Marshal(value)
	if err != nil {
		return nil
	}
	var diagnostics []sourceSiteDiagnostic
	if err := json.Unmarshal(raw, &diagnostics); err == nil {
		return diagnostics
	}
	var diagnostic sourceSiteDiagnostic
	if err := json.Unmarshal(raw, &diagnostic); err == nil && diagnostic.Kind != "" {
		return []sourceSiteDiagnostic{diagnostic}
	}
	return nil
}

func graphSourceSiteIDs(graphFile GraphFile) map[string]bool {
	ids := map[string]bool{}
	for _, relationship := range graphFile.Relationships {
		for _, id := range relationshipSourceSiteIDs(relationship) {
			ids[id] = true
		}
	}
	for _, node := range graphFile.Nodes {
		for _, diagnostic := range sourceSiteDiagnosticsFromNode(node) {
			id := strings.TrimSpace(diagnostic.SourceSiteID)
			if id != "" {
				ids[id] = true
			}
		}
	}
	return ids
}

func graphRelationshipsByGoldenKey(graphFile GraphFile) map[string]GraphRelationship {
	relationshipsByKey := map[string]GraphRelationship{}
	for _, relationship := range graphFile.Relationships {
		key := sourceSiteGoldenFalseEdgeKey(relationship.Type, relationship.SourceID, relationship.TargetID)
		if key == "" {
			continue
		}
		if _, exists := relationshipsByKey[key]; !exists {
			relationshipsByKey[key] = relationship
		}
	}
	return relationshipsByKey
}

func sourceSiteGoldenFalseEdgeKey(relationshipType string, sourceID string, targetID string) string {
	relationshipType = strings.TrimSpace(relationshipType)
	sourceID = strings.TrimSpace(sourceID)
	targetID = strings.TrimSpace(targetID)
	if relationshipType == "" || sourceID == "" || targetID == "" {
		return ""
	}
	return relationshipType + "\x00" + sourceID + "\x00" + targetID
}

func relationshipSourceSiteIDs(relationship GraphRelationship) []string {
	seen := map[string]bool{}
	var ids []string
	add := func(id string) {
		id = strings.TrimSpace(id)
		if id == "" || seen[id] {
			return
		}
		seen[id] = true
		ids = append(ids, id)
	}
	add(relationship.SourceSiteID)
	for _, id := range relationship.SourceSiteIDs {
		add(id)
	}
	return ids
}

func isResolvedAccuracyRelationship(relationshipType string) bool {
	return relationshipType == sourceSiteRelCalls || relationshipType == sourceSiteRelAccesses
}

func relationshipFactFamily(relationshipType string) string {
	switch relationshipType {
	case sourceSiteRelCalls:
		return "call"
	case sourceSiteRelAccesses:
		return "access"
	case sourceSiteRelUses:
		return "type-reference"
	case sourceSiteRelInherits:
		return "heritage"
	default:
		return normalizedMapKey(relationshipType)
	}
}

func relationshipSourceSiteOccurrenceCount(relationship GraphRelationship) int {
	if relationship.SourceSiteCount > 0 {
		return relationship.SourceSiteCount
	}
	if len(relationship.SourceSiteIDs) > 0 {
		return len(relationship.SourceSiteIDs)
	}
	if relationship.SourceSiteID != "" {
		return 1
	}
	return 0
}

func sourceSiteDiagnosticOccurrenceCount(diagnostic sourceSiteDiagnostic) int {
	if diagnostic.Count > 0 {
		return diagnostic.Count
	}
	return 1
}

func relationshipHasStableSourceSiteID(relationship GraphRelationship) bool {
	return relationship.SourceSiteID != "" || len(relationship.SourceSiteIDs) > 0
}

func normalizedSourceSiteStatus(status string, fallback string) string {
	status = strings.TrimSpace(status)
	if status != "" {
		return status
	}
	return fallback
}

func normalizedProofKind(proofKind string) string {
	proofKind = strings.TrimSpace(proofKind)
	if proofKind == "" {
		return sourceSiteProofMissing
	}
	return proofKind
}

func normalizedMapKey(value string) string {
	value = strings.TrimSpace(value)
	if value == "" {
		return sourceSiteStatusUnknown
	}
	return value
}

func addCount(counts map[string]int, key string, count int) {
	if count <= 0 {
		count = 1
	}
	counts[normalizedMapKey(key)] += count
}

func addSourceSiteExample(examples *[]SourceSiteAccuracyExample, example SourceSiteAccuracyExample, maxExamples int) {
	if maxExamples <= 0 || len(*examples) >= maxExamples {
		return
	}
	*examples = append(*examples, example)
}

func relationshipSourceSiteExample(relationship GraphRelationship, nodesByID map[string]GraphNode, reason string) SourceSiteAccuracyExample {
	source := nodesByID[relationship.SourceID]
	target := nodesByID[relationship.TargetID]
	return SourceSiteAccuracyExample{
		ID:               relationship.ID,
		Type:             relationship.Type,
		SourceID:         relationship.SourceID,
		SourceLabel:      source.Label,
		TargetID:         relationship.TargetID,
		TargetLabel:      target.Label,
		SourceSiteID:     relationship.SourceSiteID,
		SourceSiteStatus: relationship.SourceSiteStatus,
		ProofKind:        relationship.ProofKind,
		TargetRole:       relationship.TargetRole,
		TargetText:       relationship.TargetText,
		FilePath:         filepath.ToSlash(relationship.FilePath),
		StartLine:        relationship.StartLine,
		Count:            relationshipSourceSiteOccurrenceCount(relationship),
		Reason:           reason,
	}
}

func diagnosticSourceSiteExample(node GraphNode, diagnostic sourceSiteDiagnostic) SourceSiteAccuracyExample {
	return SourceSiteAccuracyExample{
		ID:               diagnostic.Kind,
		NodeID:           node.ID,
		FactFamily:       diagnostic.FactFamily,
		SourceSiteID:     diagnostic.SourceSiteID,
		SourceSiteStatus: diagnostic.SourceSiteStatus,
		ProofKind:        diagnostic.ProofKind,
		TargetRole:       diagnostic.TargetRole,
		TargetText:       diagnostic.TargetText,
		FilePath:         filepath.ToSlash(diagnostic.FilePath),
		StartLine:        diagnostic.StartLine,
		Count:            sourceSiteDiagnosticOccurrenceCount(diagnostic),
		Reason:           diagnostic.Note,
	}
}
