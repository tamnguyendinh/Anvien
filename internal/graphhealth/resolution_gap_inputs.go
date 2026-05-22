package graphhealth

import (
	"sort"
	"strings"

	"github.com/tamnguyendinh/avmatrix-go/internal/graph"
	"github.com/tamnguyendinh/avmatrix-go/internal/scopeir"
)

const (
	resolutionGapInputKindPrefix  = "ResolutionGapInput:"
	resolutionGapNodeKindPrefix   = "ResolutionGap:"
	resolutionGapRelKindPrefix    = "rel:has-resolution-gap:"
	resolutionGapFactFamilyCall   = "call"
	resolutionGapFactFamilyAccess = "access"

	ResolutionGapKindUnresolvedCall          = "unresolved_call"
	ResolutionGapKindUnresolvedAccess        = "unresolved_access"
	ResolutionGapKindUnresolvedTypeReference = "unresolved_type_reference"
	ResolutionGapKindUnresolvedHeritage      = "unresolved_heritage"
	ResolutionGapKindUnresolvedReference     = "unresolved_reference"
)

// ResolutionGapInput is the source-backed input record that later phases can
// persist as a ResolutionGap/UnresolvedSymbol without rereading graph-health
// summaries or inventing resolved target semantics.
type ResolutionGapInput struct {
	ID                   string `json:"id"`
	SourceSiteID         string `json:"sourceSiteId"`
	SourceNodeID         string `json:"sourceNodeId"`
	SourceNodeLabel      string `json:"sourceNodeLabel,omitempty"`
	SourceAppLayer       string `json:"sourceAppLayer,omitempty"`
	SourceFunctionalArea string `json:"sourceFunctionalArea,omitempty"`
	FactFamily           string `json:"factFamily"`
	TargetText           string `json:"targetText"`
	TargetRole           string `json:"targetRole,omitempty"`
	SourceSiteStatus     string `json:"sourceSiteStatus,omitempty"`
	ProofKind            string `json:"proofKind,omitempty"`
	Classification       string `json:"classification,omitempty"`
	Actionability        string `json:"actionability,omitempty"`
	ResolutionSource     string `json:"resolutionSource,omitempty"`
	Source               string `json:"source,omitempty"`
	FilePath             string `json:"filePath,omitempty"`
	FileHash             string `json:"fileHash,omitempty"`
	StartLine            int    `json:"startLine,omitempty"`
	StartCol             int    `json:"startCol,omitempty"`
	EndLine              int    `json:"endLine,omitempty"`
	EndCol               int    `json:"endCol,omitempty"`
	Count                int    `json:"count"`
	Note                 string `json:"note,omitempty"`
}

func (input ResolutionGapInput) ResolutionGapNodeID() string {
	return ResolutionGapNodeID(
		input.SourceSiteID,
		input.SourceNodeID,
		input.FactFamily,
		input.TargetText,
		input.TargetRole,
		input.SourceSiteStatus,
		input.ProofKind,
		input.Classification,
		input.Actionability,
	)
}

func ResolutionGapNodeID(sourceSiteID string, identityParts ...string) string {
	sourceSiteID = strings.TrimSpace(sourceSiteID)
	if sourceSiteID != "" {
		return resolutionGapNodeKindPrefix + sourceSiteID
	}
	cleanParts := make([]string, 0, len(identityParts))
	for _, part := range identityParts {
		if trimmed := strings.TrimSpace(part); trimmed != "" {
			cleanParts = append(cleanParts, trimmed)
		}
	}
	if len(cleanParts) == 0 {
		return resolutionGapNodeKindPrefix + "unknown"
	}
	return resolutionGapNodeKindPrefix + strings.Join(cleanParts, "|")
}

func (input ResolutionGapInput) ResolutionGapRelationshipID() string {
	return resolutionGapRelKindPrefix + strings.TrimSpace(input.SourceNodeID) + "->" + input.ResolutionGapNodeID()
}

func (input ResolutionGapInput) GapKind() string {
	switch strings.TrimSpace(input.FactFamily) {
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

func (input ResolutionGapInput) GraphNode() graph.Node {
	appLayer := strings.TrimSpace(input.SourceAppLayer)
	if appLayer == "" {
		appLayer = "unknown"
	}
	functionalArea := strings.TrimSpace(input.SourceFunctionalArea)
	if functionalArea == "" {
		functionalArea = "unknown"
	}
	count := input.Count
	if count <= 0 {
		count = 1
	}
	name := strings.TrimSpace(input.TargetText)
	if name == "" {
		name = input.GapKind()
	}
	return graph.Node{
		ID:    input.ResolutionGapNodeID(),
		Label: scopeir.NodeResolutionGap,
		Properties: graph.NodeProperties{
			"name":                 name,
			"gapKind":              input.GapKind(),
			"sourceSiteId":         strings.TrimSpace(input.SourceSiteID),
			"sourceNodeId":         strings.TrimSpace(input.SourceNodeID),
			"sourceNodeLabel":      strings.TrimSpace(input.SourceNodeLabel),
			"sourceAppLayer":       appLayer,
			"sourceFunctionalArea": functionalArea,
			"factFamily":           strings.TrimSpace(input.FactFamily),
			"targetText":           strings.TrimSpace(input.TargetText),
			"targetRole":           strings.TrimSpace(input.TargetRole),
			"sourceSiteStatus":     strings.TrimSpace(input.SourceSiteStatus),
			"proofKind":            strings.TrimSpace(input.ProofKind),
			"classification":       strings.TrimSpace(input.Classification),
			"actionability":        strings.TrimSpace(input.Actionability),
			"resolutionSource":     strings.TrimSpace(input.ResolutionSource),
			"source":               strings.TrimSpace(input.Source),
			"filePath":             strings.TrimSpace(input.FilePath),
			"fileHash":             strings.TrimSpace(input.FileHash),
			"startLine":            input.StartLine,
			"startCol":             input.StartCol,
			"endLine":              input.EndLine,
			"endCol":               input.EndCol,
			"count":                count,
			"note":                 strings.TrimSpace(input.Note),
			"appLayer":             appLayer,
			"appLayerSource":       "resolution_gap_source_node",
			"functionalArea":       functionalArea,
			"functionalAreaSource": "resolution_gap_source_node",
		},
	}
}

func (input ResolutionGapInput) GraphRelationship() graph.Relationship {
	count := input.Count
	if count <= 0 {
		count = 1
	}
	sourceSiteID := strings.TrimSpace(input.SourceSiteID)
	sourceSiteIDs := []string(nil)
	if sourceSiteID != "" {
		sourceSiteIDs = []string{sourceSiteID}
	}
	note := strings.TrimSpace(input.Note)
	if note == "" {
		note = "source-backed unresolved reference"
	}
	return graph.Relationship{
		ID:               input.ResolutionGapRelationshipID(),
		SourceID:         strings.TrimSpace(input.SourceNodeID),
		TargetID:         input.ResolutionGapNodeID(),
		Type:             graph.RelHasResolutionGap,
		Confidence:       1,
		Reason:           "source-backed unresolved reference",
		ResolutionSource: strings.TrimSpace(input.ResolutionSource),
		FileHash:         strings.TrimSpace(input.FileHash),
		SourceSiteID:     sourceSiteID,
		SourceSiteIDs:    sourceSiteIDs,
		SourceSiteCount:  count,
		SourceSiteStatus: strings.TrimSpace(input.SourceSiteStatus),
		ProofKind:        strings.TrimSpace(input.ProofKind),
		TargetRole:       strings.TrimSpace(input.TargetRole),
		TargetText:       strings.TrimSpace(input.TargetText),
		FilePath:         strings.TrimSpace(input.FilePath),
		StartLine:        input.StartLine,
		StartCol:         input.StartCol,
		EndLine:          input.EndLine,
		EndCol:           input.EndCol,
		Evidence: []graph.Evidence{{
			Kind:   "resolution_gap",
			Weight: 1,
			Note:   note,
		}},
	}
}

func SourceBackedResolutionGapInputs(g *graph.Graph) []ResolutionGapInput {
	if g == nil {
		return nil
	}
	var inputs []ResolutionGapInput
	for _, node := range g.Nodes {
		for _, diagnostic := range diagnosticsFromProperties(node.Properties) {
			if diagnostic.Kind != DiagnosticUnresolvedReference || strings.TrimSpace(diagnostic.SourceSiteID) == "" {
				continue
			}
			inputs = append(inputs, resolutionGapInputFromDiagnostic(node, diagnostic))
		}
	}
	sort.SliceStable(inputs, func(i int, j int) bool {
		left := inputs[i]
		right := inputs[j]
		if left.SourceNodeID != right.SourceNodeID {
			return left.SourceNodeID < right.SourceNodeID
		}
		if left.FactFamily != right.FactFamily {
			return left.FactFamily < right.FactFamily
		}
		if left.FilePath != right.FilePath {
			return left.FilePath < right.FilePath
		}
		if left.StartLine != right.StartLine {
			return left.StartLine < right.StartLine
		}
		if left.SourceSiteID != right.SourceSiteID {
			return left.SourceSiteID < right.SourceSiteID
		}
		return left.TargetText < right.TargetText
	})
	return inputs
}

func SourceBackedCallAccessResolutionGapInputs(g *graph.Graph) []ResolutionGapInput {
	inputs := SourceBackedResolutionGapInputs(g)
	filtered := inputs[:0]
	for _, input := range inputs {
		if input.FactFamily == resolutionGapFactFamilyCall || input.FactFamily == resolutionGapFactFamilyAccess {
			filtered = append(filtered, input)
		}
	}
	return filtered
}

func resolutionGapInputFromDiagnostic(node graph.Node, diagnostic Diagnostic) ResolutionGapInput {
	sourceNodeID := strings.TrimSpace(diagnostic.SourceNodeID)
	if sourceNodeID == "" {
		sourceNodeID = node.ID
	}
	return ResolutionGapInput{
		ID:                   resolutionGapInputKindPrefix + strings.TrimSpace(diagnostic.SourceSiteID),
		SourceSiteID:         strings.TrimSpace(diagnostic.SourceSiteID),
		SourceNodeID:         sourceNodeID,
		SourceNodeLabel:      string(node.Label),
		SourceAppLayer:       stringNodeProperty(node.Properties, "appLayer"),
		SourceFunctionalArea: stringNodeProperty(node.Properties, "functionalArea"),
		FactFamily:           strings.TrimSpace(diagnostic.FactFamily),
		TargetText:           strings.TrimSpace(diagnostic.TargetText),
		TargetRole:           strings.TrimSpace(diagnostic.TargetRole),
		SourceSiteStatus:     strings.TrimSpace(diagnostic.SourceSiteStatus),
		ProofKind:            strings.TrimSpace(diagnostic.ProofKind),
		Classification:       strings.TrimSpace(diagnostic.Classification),
		Actionability:        strings.TrimSpace(diagnostic.Actionability),
		ResolutionSource:     strings.TrimSpace(diagnostic.ResolutionSource),
		Source:               strings.TrimSpace(diagnostic.Source),
		FilePath:             strings.TrimSpace(diagnostic.FilePath),
		FileHash:             strings.TrimSpace(diagnostic.FileHash),
		StartLine:            diagnostic.StartLine,
		StartCol:             diagnostic.StartCol,
		EndLine:              diagnostic.EndLine,
		EndCol:               diagnostic.EndCol,
		Count:                diagnosticCount(diagnostic),
		Note:                 strings.TrimSpace(diagnostic.Note),
	}
}

func stringNodeProperty(properties graph.NodeProperties, key string) string {
	if properties == nil {
		return ""
	}
	value, ok := properties[key].(string)
	if !ok {
		return ""
	}
	return strings.TrimSpace(value)
}
