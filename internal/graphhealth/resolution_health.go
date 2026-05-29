package graphhealth

import (
	"strings"

	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
)

type resolutionHealthInventory struct {
	bucketsByNode            map[string]map[string]int
	gapCountsByNode          map[string]int
	resolvedReferenceByNode  map[string]int
	resolutionGapNodeCount   int
	hasResolutionGapRelCount int
	resolutionGapCount       int
	resolvedReferenceCount   int
	bucketCounts             map[string]int
	factFamilyCounts         map[string]int
	targetRoleCounts         map[string]int
	classificationCounts     map[string]int
	actionabilityCounts      map[string]int
	appLayerCounts           map[string]int
	functionalAreaCounts     map[string]int
}

func computeResolutionHealthInventory(g *graph.Graph, nodeByID map[string]graph.Node) resolutionHealthInventory {
	inventory := newResolutionHealthInventory()
	if g == nil {
		return inventory
	}
	for _, node := range g.Nodes {
		if node.Label == scopeir.NodeResolutionGap {
			inventory.resolutionGapNodeCount++
		}
	}
	for _, relationship := range g.Relationships {
		if relationship.Type == graph.RelHasResolutionGap {
			inventory.recordResolutionGapRelationship(relationship, nodeByID)
			continue
		}
		if isResolvedReferenceRelationship(relationship) {
			inventory.recordResolvedReference(relationship)
		}
	}
	return inventory
}

func newResolutionHealthInventory() resolutionHealthInventory {
	inventory := resolutionHealthInventory{
		bucketsByNode:           map[string]map[string]int{},
		gapCountsByNode:         map[string]int{},
		resolvedReferenceByNode: map[string]int{},
		bucketCounts:            map[string]int{},
		factFamilyCounts:        map[string]int{},
		targetRoleCounts:        map[string]int{},
		classificationCounts:    map[string]int{},
		actionabilityCounts:     map[string]int{},
		appLayerCounts:          map[string]int{},
		functionalAreaCounts:    map[string]int{},
	}
	for _, bucket := range ResolutionHealthBuckets {
		inventory.bucketCounts[string(bucket)] = 0
	}
	return inventory
}

func (inventory *resolutionHealthInventory) recordResolutionGapRelationship(relationship graph.Relationship, nodeByID map[string]graph.Node) {
	count := relationshipSourceSiteCount(relationship)
	if count <= 0 {
		count = 1
	}
	inventory.hasResolutionGapRelCount++
	inventory.resolutionGapCount += count
	inventory.gapCountsByNode[relationship.SourceID] += count

	gapNode := nodeByID[relationship.TargetID]
	sourceNode := nodeByID[relationship.SourceID]
	factFamily := firstNonEmpty(
		relationshipFactFamily(relationship),
		stringNodeProperty(gapNode.Properties, "factFamily"),
	)
	gapKind := firstNonEmpty(
		stringNodeProperty(gapNode.Properties, "gapKind"),
		gapKindForFactFamily(factFamily),
	)
	targetRole := firstNonEmpty(
		strings.TrimSpace(relationship.TargetRole),
		stringNodeProperty(gapNode.Properties, "targetRole"),
		"unknown",
	)
	classification := firstNonEmpty(
		stringNodeProperty(gapNode.Properties, "classification"),
		DiagnosticClassificationUnclassified,
	)
	actionability := firstNonEmpty(
		stringNodeProperty(gapNode.Properties, "actionability"),
		actionabilityForDiagnosticClassification(classification),
	)
	appLayer := firstNonEmpty(
		stringNodeProperty(gapNode.Properties, "sourceAppLayer"),
		stringNodeProperty(sourceNode.Properties, "appLayer"),
		stringNodeProperty(gapNode.Properties, "appLayer"),
		"unknown",
	)
	functionalArea := firstNonEmpty(
		stringNodeProperty(gapNode.Properties, "sourceFunctionalArea"),
		stringNodeProperty(sourceNode.Properties, "functionalArea"),
		stringNodeProperty(gapNode.Properties, "functionalArea"),
		"unknown",
	)

	incrementCount(inventory.factFamilyCounts, normalizedInventoryKey(factFamily), count)
	incrementCount(inventory.targetRoleCounts, normalizedInventoryKey(targetRole), count)
	incrementCount(inventory.classificationCounts, normalizedInventoryKey(classification), count)
	incrementCount(inventory.actionabilityCounts, normalizedInventoryKey(actionability), count)
	incrementCount(inventory.appLayerCounts, normalizedInventoryKey(appLayer), count)
	incrementCount(inventory.functionalAreaCounts, normalizedInventoryKey(functionalArea), count)
	for _, bucket := range resolutionHealthBucketsForGap(gapKind, factFamily, classification, actionability) {
		inventory.addNodeBucket(relationship.SourceID, bucket, count)
	}
}

func (inventory *resolutionHealthInventory) recordResolvedReference(relationship graph.Relationship) {
	count := relationshipSourceSiteCount(relationship)
	if count <= 0 {
		return
	}
	inventory.resolvedReferenceCount += count
	inventory.resolvedReferenceByNode[relationship.SourceID] += count
	inventory.addNodeBucket(relationship.SourceID, ResolutionHealthResolvedReferences, count)
}

func (inventory *resolutionHealthInventory) addNodeBucket(nodeID string, bucket ResolutionHealthBucket, count int) {
	if nodeID == "" {
		return
	}
	if count <= 0 {
		count = 1
	}
	bucketKey := string(bucket)
	if inventory.bucketsByNode[nodeID] == nil {
		inventory.bucketsByNode[nodeID] = map[string]int{}
	}
	inventory.bucketsByNode[nodeID][bucketKey] += count
	inventory.bucketCounts[bucketKey] += count
}

func resolutionHealthBucketsForGap(gapKind string, factFamily string, classification string, actionability string) []ResolutionHealthBucket {
	buckets := []ResolutionHealthBucket{}
	switch strings.TrimSpace(gapKind) {
	case ResolutionGapKindUnresolvedCall:
		buckets = append(buckets, ResolutionHealthUnresolvedCallTarget)
	case ResolutionGapKindUnresolvedAccess:
		buckets = append(buckets, ResolutionHealthUnresolvedAccessTarget)
	case ResolutionGapKindUnresolvedTypeReference:
		buckets = append(buckets, ResolutionHealthUnresolvedTypeTarget)
	case ResolutionGapKindUnresolvedHeritage:
		buckets = append(buckets, ResolutionHealthUnresolvedHeritageTarget)
	default:
		switch strings.TrimSpace(factFamily) {
		case "call":
			buckets = append(buckets, ResolutionHealthUnresolvedCallTarget)
		case "access":
			buckets = append(buckets, ResolutionHealthUnresolvedAccessTarget)
		case "type-reference", "type_reference", "type":
			buckets = append(buckets, ResolutionHealthUnresolvedTypeTarget)
		case "heritage", "inheritance", "inherits", "extends", "implements":
			buckets = append(buckets, ResolutionHealthUnresolvedHeritageTarget)
		default:
			buckets = append(buckets, ResolutionHealthUnclassifiedUnknown)
		}
	}

	switch strings.TrimSpace(classification) {
	case DiagnosticClassificationBuiltin, DiagnosticClassificationStandardLibrary, DiagnosticClassificationTestFramework:
		buckets = append(buckets, ResolutionHealthUnresolvedNonActionable)
	case DiagnosticClassificationExternalLibrary:
		buckets = append(buckets, ResolutionHealthExternalUnresolved)
	case DiagnosticClassificationInRepoUnresolved:
		if strings.TrimSpace(actionability) == DiagnosticActionabilityAnalyzerGap {
			buckets = append(buckets, ResolutionHealthInRepoAnalyzerGap)
		} else {
			buckets = append(buckets, ResolutionHealthUnclassifiedUnknown)
		}
	default:
		if strings.TrimSpace(actionability) == DiagnosticActionabilityNonActionable {
			buckets = append(buckets, ResolutionHealthUnresolvedNonActionable)
		} else {
			buckets = append(buckets, ResolutionHealthUnclassifiedUnknown)
		}
	}
	return uniqueResolutionHealthBuckets(buckets)
}

func resolutionConfidenceForNode(buckets map[string]int, gapCount int) string {
	if gapCount > 0 {
		return ResolutionConfidenceDegraded
	}
	if buckets[string(ResolutionHealthResolvedReferences)] > 0 {
		return ResolutionConfidenceClear
	}
	return ResolutionConfidenceUnknown
}

func isResolvedReferenceRelationship(relationship graph.Relationship) bool {
	if relationship.SourceID == "" || relationship.TargetID == "" {
		return false
	}
	if relationship.SourceSiteStatus != "" && relationship.SourceSiteStatus != "resolved" {
		return false
	}
	switch relationship.Type {
	case graph.RelCalls,
		graph.RelAccesses,
		graph.RelUses,
		graph.RelInherits,
		graph.RelImplements,
		graph.RelExtends,
		graph.RelMethodOverrides,
		graph.RelMethodImplements:
		return true
	default:
		return false
	}
}

func relationshipSourceSiteCount(relationship graph.Relationship) int {
	if relationship.SourceSiteCount > 0 {
		return relationship.SourceSiteCount
	}
	if len(relationship.SourceSiteIDs) > 0 {
		return len(relationship.SourceSiteIDs)
	}
	if strings.TrimSpace(relationship.SourceSiteID) != "" {
		return 1
	}
	return 0
}

func relationshipFactFamily(relationship graph.Relationship) string {
	switch relationship.Type {
	case graph.RelCalls:
		return "call"
	case graph.RelAccesses:
		return "access"
	case graph.RelUses:
		return "type-reference"
	case graph.RelInherits, graph.RelImplements, graph.RelExtends:
		return "heritage"
	default:
		return ""
	}
}

func gapKindForFactFamily(factFamily string) string {
	switch strings.TrimSpace(factFamily) {
	case "call":
		return ResolutionGapKindUnresolvedCall
	case "access":
		return ResolutionGapKindUnresolvedAccess
	case "type-reference", "type_reference", "type":
		return ResolutionGapKindUnresolvedTypeReference
	case "heritage", "inheritance", "inherits", "extends", "implements":
		return ResolutionGapKindUnresolvedHeritage
	default:
		return ResolutionGapKindUnresolvedReference
	}
}

func uniqueResolutionHealthBuckets(buckets []ResolutionHealthBucket) []ResolutionHealthBucket {
	out := buckets[:0]
	seen := map[ResolutionHealthBucket]bool{}
	for _, bucket := range buckets {
		if seen[bucket] {
			continue
		}
		seen[bucket] = true
		out = append(out, bucket)
	}
	return out
}

func incrementCount(counts map[string]int, key string, count int) {
	if key == "" {
		key = "unknown"
	}
	if count <= 0 {
		count = 1
	}
	counts[key] += count
}

func normalizedInventoryKey(value string) string {
	value = strings.TrimSpace(value)
	if value == "" {
		return "unknown"
	}
	return value
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if trimmed := strings.TrimSpace(value); trimmed != "" {
			return trimmed
		}
	}
	return ""
}
