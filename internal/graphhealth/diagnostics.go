package graphhealth

import (
	"encoding/json"
	"sort"

	"github.com/tamnguyendinh/avmatrix-go/internal/graph"
)

const (
	resolutionMetadataKey                    = "resolution"
	resolutionMetadataUnresolvedReferences   = "unresolvedReferences"
	resolutionMetadataSourceBackedUnresolved = "sourceBackedUnresolvedReferences"
	resolutionMetadataUnattributedUnresolved = "unattributedUnresolvedReferences"
)

// AppendDiagnosticToNode records analyzer evidence on a real graph node.
// It returns false when the source node is absent, so callers can keep the
// unresolved aggregate count without inventing node-level evidence.
func AppendDiagnosticToNode(g *graph.Graph, nodeID string, diagnostic Diagnostic) bool {
	if g == nil || nodeID == "" || diagnostic.Kind == "" {
		return false
	}
	node, ok := g.GetNode(nodeID)
	if !ok {
		return false
	}
	if node.Properties == nil {
		node.Properties = graph.NodeProperties{}
	}
	if diagnostic.SourceNodeID == "" {
		diagnostic.SourceNodeID = nodeID
	}
	if diagnostic.Count <= 0 {
		diagnostic.Count = 1
	}
	node.Properties[DiagnosticPropertyKey] = appendDiagnosticsFromProperties(node.Properties, diagnostic)
	g.AddNode(node)
	return true
}

func SetResolutionMetadata(g *graph.Graph, unresolvedReferences int, sourceBacked int, unattributed int) {
	if g == nil {
		return
	}
	if unresolvedReferences == 0 && sourceBacked == 0 && unattributed == 0 {
		return
	}
	if g.Metadata == nil {
		g.Metadata = map[string]any{}
	}
	g.Metadata[resolutionMetadataKey] = map[string]any{
		resolutionMetadataUnresolvedReferences:   unresolvedReferences,
		resolutionMetadataSourceBackedUnresolved: sourceBacked,
		resolutionMetadataUnattributedUnresolved: unattributed,
	}
}

func resolutionMetadata(g *graph.Graph) (unresolvedReferences int, sourceBacked int, unattributed int, ok bool) {
	if g == nil || len(g.Metadata) == 0 {
		return 0, 0, 0, false
	}
	value, exists := g.Metadata[resolutionMetadataKey]
	if !exists {
		return 0, 0, 0, false
	}
	metadata, exists := value.(map[string]any)
	if !exists {
		if raw, err := json.Marshal(value); err == nil {
			_ = json.Unmarshal(raw, &metadata)
		}
	}
	if metadata == nil {
		return 0, 0, 0, false
	}
	return intValue(metadata[resolutionMetadataUnresolvedReferences]),
		intValue(metadata[resolutionMetadataSourceBackedUnresolved]),
		intValue(metadata[resolutionMetadataUnattributedUnresolved]),
		true
}

func diagnosticsFromProperties(properties graph.NodeProperties) []Diagnostic {
	if len(properties) == 0 {
		return nil
	}
	return normalizeDiagnostics(properties[DiagnosticPropertyKey])
}

func appendDiagnosticsFromProperties(properties graph.NodeProperties, diagnostic Diagnostic) []Diagnostic {
	diagnostics := diagnosticsFromProperties(properties)
	for index := range diagnostics {
		if sameDiagnosticBucket(diagnostics[index], diagnostic) {
			diagnostics[index].Count += diagnosticCount(diagnostic)
			if diagnostics[index].TargetText == "" {
				diagnostics[index].TargetText = diagnostic.TargetText
			}
			if diagnostics[index].StartLine == 0 || (diagnostic.StartLine > 0 && diagnostic.StartLine < diagnostics[index].StartLine) {
				diagnostics[index].StartLine = diagnostic.StartLine
			}
			return diagnostics
		}
	}
	diagnostics = append(diagnostics, diagnostic)
	sort.SliceStable(diagnostics, func(i int, j int) bool {
		if diagnostics[i].Kind != diagnostics[j].Kind {
			return diagnostics[i].Kind < diagnostics[j].Kind
		}
		if diagnostics[i].FactFamily != diagnostics[j].FactFamily {
			return diagnostics[i].FactFamily < diagnostics[j].FactFamily
		}
		if diagnostics[i].FilePath != diagnostics[j].FilePath {
			return diagnostics[i].FilePath < diagnostics[j].FilePath
		}
		if diagnostics[i].StartLine != diagnostics[j].StartLine {
			return diagnostics[i].StartLine < diagnostics[j].StartLine
		}
		return diagnostics[i].TargetText < diagnostics[j].TargetText
	})
	return diagnostics
}

func sameDiagnosticBucket(left Diagnostic, right Diagnostic) bool {
	return left.Kind == right.Kind &&
		left.FactFamily == right.FactFamily &&
		left.SourceNodeID == right.SourceNodeID &&
		left.ResolutionSource == right.ResolutionSource &&
		left.FilePath == right.FilePath &&
		left.Note == right.Note
}

func diagnosticCount(diagnostic Diagnostic) int {
	if diagnostic.Count > 0 {
		return diagnostic.Count
	}
	return 1
}

func normalizeDiagnostics(value any) []Diagnostic {
	switch typed := value.(type) {
	case nil:
		return nil
	case []Diagnostic:
		return append([]Diagnostic(nil), typed...)
	case Diagnostic:
		return []Diagnostic{typed}
	case []any:
		out := make([]Diagnostic, 0, len(typed))
		for _, item := range typed {
			if diagnostic, ok := normalizeDiagnostic(item); ok {
				out = append(out, diagnostic)
			}
		}
		return out
	default:
		if diagnostic, ok := normalizeDiagnostic(typed); ok {
			return []Diagnostic{diagnostic}
		}
		return nil
	}
}

func normalizeDiagnostic(value any) (Diagnostic, bool) {
	switch typed := value.(type) {
	case Diagnostic:
		return typed, typed.Kind != ""
	case map[string]any:
		diagnostic := Diagnostic{
			Kind:             stringMapValue(typed, "kind"),
			FactFamily:       stringMapValue(typed, "factFamily"),
			SourceNodeID:     stringMapValue(typed, "sourceNodeId"),
			TargetText:       stringMapValue(typed, "targetText"),
			ResolutionSource: stringMapValue(typed, "resolutionSource"),
			FilePath:         stringMapValue(typed, "filePath"),
			FileHash:         stringMapValue(typed, "fileHash"),
			StartLine:        intValue(typed["startLine"]),
			Count:            intValue(typed["count"]),
			Note:             stringMapValue(typed, "note"),
			Source:           stringMapValue(typed, "source"),
		}
		return diagnostic, diagnostic.Kind != ""
	default:
		raw, err := json.Marshal(typed)
		if err != nil {
			return Diagnostic{}, false
		}
		var diagnostic Diagnostic
		if err := json.Unmarshal(raw, &diagnostic); err != nil {
			return Diagnostic{}, false
		}
		return diagnostic, diagnostic.Kind != ""
	}
}

func stringMapValue(values map[string]any, key string) string {
	value, ok := values[key]
	if !ok {
		return ""
	}
	text, ok := value.(string)
	if !ok {
		return ""
	}
	return text
}

func intValue(value any) int {
	switch typed := value.(type) {
	case int:
		return typed
	case int8:
		return int(typed)
	case int16:
		return int(typed)
	case int32:
		return int(typed)
	case int64:
		return int(typed)
	case uint:
		return int(typed)
	case uint8:
		return int(typed)
	case uint16:
		return int(typed)
	case uint32:
		return int(typed)
	case uint64:
		return int(typed)
	case float64:
		return int(typed)
	case float32:
		return int(typed)
	case json.Number:
		value, _ := typed.Int64()
		return int(value)
	default:
		return 0
	}
}
