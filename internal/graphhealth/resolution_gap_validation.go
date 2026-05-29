package graphhealth

import (
	"strings"

	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
)

type ResolutionGapValidation struct {
	ResolutionGapNodes               int `json:"resolutionGapNodes"`
	HasResolutionGapRelationships    int `json:"hasResolutionGapRelationships"`
	DanglingGapRelationshipCount     int `json:"danglingGapRelationshipCount"`
	NonGapTargetRelationshipCount    int `json:"nonGapTargetRelationshipCount"`
	CountedGapRelationshipCount      int `json:"countedGapRelationshipCount"`
	GapResolvedClaimCount            int `json:"gapResolvedClaimCount"`
	ProoflessResolvedSourceSiteEdges int `json:"prooflessResolvedSourceSiteEdges"`
}

func (validation ResolutionGapValidation) OK() bool {
	return validation.DanglingGapRelationshipCount == 0 &&
		validation.NonGapTargetRelationshipCount == 0 &&
		validation.CountedGapRelationshipCount == 0 &&
		validation.GapResolvedClaimCount == 0 &&
		validation.ProoflessResolvedSourceSiteEdges == 0
}

func ValidateResolutionGapPersistence(g *graph.Graph) ResolutionGapValidation {
	var validation ResolutionGapValidation
	if g == nil {
		return validation
	}
	nodeByID := make(map[string]graph.Node, len(g.Nodes))
	gapSourceSites := map[string]struct{}{}
	for _, node := range g.Nodes {
		nodeByID[node.ID] = node
		if node.Label != scopeir.NodeResolutionGap {
			continue
		}
		validation.ResolutionGapNodes++
		if sourceSiteID := stringNodeProperty(node.Properties, "sourceSiteId"); sourceSiteID != "" {
			gapSourceSites[sourceSiteID] = struct{}{}
		}
		if gapClaimsResolvedTarget(node.Properties) {
			validation.GapResolvedClaimCount++
		}
	}
	for _, relationship := range g.Relationships {
		if relationship.Type == graph.RelHasResolutionGap {
			validation.HasResolutionGapRelationships++
			sourceMissing := false
			if _, ok := nodeByID[relationship.SourceID]; !ok {
				sourceMissing = true
			}
			target, targetOK := nodeByID[relationship.TargetID]
			if sourceMissing || !targetOK {
				validation.DanglingGapRelationshipCount++
			}
			if targetOK && target.Label != scopeir.NodeResolutionGap {
				validation.NonGapTargetRelationshipCount++
			}
			if IsCounted(relationship.Type) {
				validation.CountedGapRelationshipCount++
			}
			continue
		}
		if isResolvedSemanticRelationship(relationship.Type) {
			if relationshipUsesGapSourceSite(relationship, gapSourceSites) {
				validation.ProoflessResolvedSourceSiteEdges++
			}
		}
	}
	return validation
}

func relationshipUsesGapSourceSite(relationship graph.Relationship, gapSourceSites map[string]struct{}) bool {
	if len(gapSourceSites) == 0 {
		return false
	}
	if sourceSiteID := strings.TrimSpace(relationship.SourceSiteID); sourceSiteID != "" {
		if _, ok := gapSourceSites[sourceSiteID]; ok {
			return true
		}
	}
	for _, sourceSiteID := range relationship.SourceSiteIDs {
		sourceSiteID = strings.TrimSpace(sourceSiteID)
		if sourceSiteID == "" {
			continue
		}
		if _, ok := gapSourceSites[sourceSiteID]; ok {
			return true
		}
	}
	return false
}

func gapClaimsResolvedTarget(properties graph.NodeProperties) bool {
	if len(properties) == 0 {
		return false
	}
	for _, key := range []string{"resolvedTargetId", "resolvedTargetID", "resolvedTargetLabel"} {
		if strings.TrimSpace(stringNodeProperty(properties, key)) != "" {
			return true
		}
	}
	return strings.EqualFold(strings.TrimSpace(stringNodeProperty(properties, "resolutionStatus")), "resolved")
}

func isResolvedSemanticRelationship(relationshipType graph.RelationshipType) bool {
	switch relationshipType {
	case graph.RelCalls,
		graph.RelAccesses,
		graph.RelUses,
		graph.RelInherits,
		graph.RelExtends,
		graph.RelImplements,
		graph.RelMethodOverrides,
		graph.RelMethodImplements:
		return true
	default:
		return false
	}
}
