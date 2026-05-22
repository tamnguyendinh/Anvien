package graphhealth

import (
	"sort"
	"strings"

	"github.com/tamnguyendinh/avmatrix-go/internal/graph"
)

const (
	resolutionGapInputKindPrefix  = "ResolutionGapInput:"
	resolutionGapFactFamilyCall   = "call"
	resolutionGapFactFamilyAccess = "access"
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
