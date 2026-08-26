package graphhealth

import (
	"encoding/json"
	"sort"
	"strings"

	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/tsstdlib"
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
	diagnostic = normalizeDiagnosticMetadata(diagnostic)
	node.Properties[DiagnosticPropertyKey] = appendDiagnosticsFromProperties(node.Properties, diagnostic)
	g.AddNode(node)
	return true
}

type diagnosticAppender struct {
	graph             *graph.Graph
	diagnosticsByNode map[string][]Diagnostic
}

// NewDiagnosticAppender creates a write-through diagnostic appender whose
// retained state is scoped to the caller's graph run.
func NewDiagnosticAppender(g *graph.Graph) func(string, Diagnostic) bool {
	appender := &diagnosticAppender{
		graph:             g,
		diagnosticsByNode: make(map[string][]Diagnostic),
	}
	return appender.appendToNode
}

func (appender *diagnosticAppender) appendToNode(nodeID string, diagnostic Diagnostic) bool {
	if appender.graph == nil || nodeID == "" || diagnostic.Kind == "" {
		return false
	}
	node, ok := appender.graph.GetNode(nodeID)
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
	diagnostic = normalizeDiagnosticMetadata(diagnostic)

	diagnostics, touched := appender.diagnosticsByNode[nodeID]
	if !touched {
		diagnostics = diagnosticsFromProperties(node.Properties)
	}
	diagnostics = appendNormalizedDiagnostic(diagnostics, diagnostic)
	appender.diagnosticsByNode[nodeID] = diagnostics
	node.Properties[DiagnosticPropertyKey] = diagnostics
	appender.graph.AddNode(node)
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
	return appendNormalizedDiagnostic(
		diagnosticsFromProperties(properties),
		normalizeDiagnosticMetadata(diagnostic),
	)
}

func appendNormalizedDiagnostic(diagnostics []Diagnostic, diagnostic Diagnostic) []Diagnostic {
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
	sameBase := left.Kind == right.Kind &&
		left.FactFamily == right.FactFamily &&
		left.SourceNodeID == right.SourceNodeID &&
		left.ResolutionSource == right.ResolutionSource &&
		left.FilePath == right.FilePath &&
		left.Note == right.Note
	if !sameBase {
		return false
	}
	if left.SourceSiteID != "" || right.SourceSiteID != "" {
		return left.SourceSiteID == right.SourceSiteID
	}
	return left.TargetText == right.TargetText &&
		left.StartLine == right.StartLine &&
		left.StartCol == right.StartCol
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
		return normalizeDiagnosticSlice(typed)
	case Diagnostic:
		return []Diagnostic{normalizeDiagnosticMetadata(typed)}
	case []any:
		out := make([]Diagnostic, 0, len(typed))
		for _, item := range typed {
			if diagnostic, ok := normalizeDiagnostic(item); ok {
				out = append(out, normalizeDiagnosticMetadata(diagnostic))
			}
		}
		return out
	default:
		if diagnostic, ok := normalizeDiagnostic(typed); ok {
			return []Diagnostic{normalizeDiagnosticMetadata(diagnostic)}
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
			Classification:   stringMapValue(typed, "classification"),
			Actionability:    stringMapValue(typed, "actionability"),
			FilePath:         stringMapValue(typed, "filePath"),
			FileHash:         stringMapValue(typed, "fileHash"),
			StartLine:        intValue(typed["startLine"]),
			StartCol:         intValue(typed["startCol"]),
			EndLine:          intValue(typed["endLine"]),
			EndCol:           intValue(typed["endCol"]),
			SourceSiteID:     stringMapValue(typed, "sourceSiteId"),
			SourceSiteStatus: stringMapValue(typed, "sourceSiteStatus"),
			ProofKind:        stringMapValue(typed, "proofKind"),
			TargetRole:       stringMapValue(typed, "targetRole"),
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

func normalizeDiagnosticSlice(values []Diagnostic) []Diagnostic {
	out := make([]Diagnostic, len(values))
	for index, diagnostic := range values {
		out[index] = normalizeDiagnosticMetadata(diagnostic)
	}
	return out
}

func normalizeDiagnosticMetadata(diagnostic Diagnostic) Diagnostic {
	if classification, actionability, structured := structuredResolutionDiagnosticPolicy(diagnostic); structured {
		diagnostic.Classification = classification
		diagnostic.Actionability = actionability
		return diagnostic
	}
	if diagnostic.Classification == "" {
		diagnostic.Classification = classifyDiagnostic(diagnostic)
	}
	if diagnostic.Actionability == "" {
		diagnostic.Actionability = actionabilityForDiagnosticClassification(diagnostic.Classification)
	}
	return diagnostic
}

const (
	structuredResolutionOutcomeSchemaVersion = 1

	structuredResolutionStageRepository            = "repository"
	structuredResolutionStageTypeScriptStandardLib = "typescript_standard_library"

	structuredResolutionStatusResolvedInternal      = "resolved_internal"
	structuredResolutionStatusResolvedExternal      = "resolved_external"
	structuredResolutionStatusUnresolved            = "unresolved"
	structuredResolutionStatusCapabilityUnavailable = "capability_unavailable"
)

type structuredResolutionOutcome struct {
	SchemaVersion    int             `json:"schemaVersion"`
	SourceSiteID     string          `json:"sourceSiteId"`
	Status           string          `json:"status"`
	Stage            string          `json:"stage"`
	SiteKind         string          `json:"siteKind"`
	FilePath         string          `json:"filePath"`
	FileHash         string          `json:"fileHash"`
	Range            structuredRange `json:"range"`
	RequestedName    string          `json:"requestedName"`
	RequestedMeaning string          `json:"requestedMeaning"`
	Language         string          `json:"language"`
	Target           json.RawMessage `json:"target"`
	Reason           string          `json:"reason"`
	Proof            structuredProof `json:"proof"`
	Authority        json.RawMessage `json:"authority"`
}

type structuredRange struct {
	StartLine int `json:"startLine"`
	StartCol  int `json:"startCol"`
	EndLine   int `json:"endLine"`
	EndCol    int `json:"endCol"`
}

type structuredProof struct {
	Kind string `json:"kind"`
}

type structuredResolutionAuthority struct {
	SourceSiteID        string                     `json:"sourceSiteId"`
	Stage               string                     `json:"stage"`
	FilePath            string                     `json:"filePath"`
	FileHash            string                     `json:"fileHash"`
	Range               structuredRange            `json:"range"`
	SiteKind            string                     `json:"siteKind"`
	RequestedName       string                     `json:"requestedName"`
	RequestedMeaning    tsstdlib.Meaning           `json:"requestedMeaning"`
	Status              tsstdlib.LookupStatus      `json:"status"`
	Reason              tsstdlib.Reason            `json:"reason"`
	ResolvedSymbolID    string                     `json:"resolvedSymbolId"`
	DeclarationRanges   []tsstdlib.Declaration     `json:"declarationRanges"`
	AuthorityKind       string                     `json:"authorityKind"`
	CatalogProofState   tsstdlib.CatalogProofState `json:"catalogProofState"`
	AuthorityHash       string                     `json:"authorityHash"`
	TypeScriptVersion   string                     `json:"typescriptVersion"`
	CatalogHash         string                     `json:"catalogHash"`
	CatalogArtifactHash string                     `json:"catalogArtifactHash"`
	ProfileHash         string                     `json:"profileHash"`
	ConfigHash          string                     `json:"configHash"`
}

func structuredResolutionDiagnosticPolicy(diagnostic Diagnostic) (string, string, bool) {
	outcome, structured, valid := decodeStructuredResolutionOutcome(diagnostic)
	if !structured {
		return "", "", false
	}
	if !valid {
		return DiagnosticClassificationUnclassified, DiagnosticActionabilityReview, true
	}
	switch outcome.Status {
	case structuredResolutionStatusUnresolved:
		switch outcome.Stage {
		case structuredResolutionStageRepository:
			return DiagnosticClassificationInRepoUnresolved, DiagnosticActionabilityAnalyzerGap, true
		case structuredResolutionStageTypeScriptStandardLib:
			return DiagnosticClassificationStandardLibrary, DiagnosticActionabilityNonActionable, true
		}
	case structuredResolutionStatusCapabilityUnavailable:
		if outcome.Stage == structuredResolutionStageTypeScriptStandardLib {
			return DiagnosticClassificationStandardLibrary, DiagnosticActionabilityNonActionable, true
		}
	}
	return DiagnosticClassificationUnclassified, DiagnosticActionabilityReview, true
}

func decodeStructuredResolutionOutcome(diagnostic Diagnostic) (structuredResolutionOutcome, bool, bool) {
	statusMarker := isStructuredResolutionStatus(strings.TrimSpace(diagnostic.SourceSiteStatus))
	note := strings.TrimSpace(diagnostic.Note)
	if note == "" {
		return structuredResolutionOutcome{}, statusMarker, false
	}

	var envelope map[string]json.RawMessage
	if err := json.Unmarshal([]byte(note), &envelope); err != nil {
		return structuredResolutionOutcome{}, statusMarker, false
	}
	_, hasSchemaVersion := envelope["schemaVersion"]
	_, hasStatus := envelope["status"]
	structured := statusMarker || hasSchemaVersion || hasStatus
	if !structured {
		return structuredResolutionOutcome{}, false, false
	}

	var outcome structuredResolutionOutcome
	if err := json.Unmarshal([]byte(note), &outcome); err != nil {
		return structuredResolutionOutcome{}, true, false
	}
	return outcome, true, validStructuredResolutionDiagnostic(diagnostic, outcome)
}

func validStructuredResolutionDiagnostic(diagnostic Diagnostic, outcome structuredResolutionOutcome) bool {
	if diagnostic.Kind != DiagnosticUnresolvedReference ||
		outcome.SchemaVersion != structuredResolutionOutcomeSchemaVersion ||
		strings.TrimSpace(outcome.SourceSiteID) == "" ||
		strings.TrimSpace(outcome.SourceSiteID) != strings.TrimSpace(diagnostic.SourceSiteID) ||
		strings.TrimSpace(outcome.Status) != strings.TrimSpace(diagnostic.SourceSiteStatus) ||
		strings.TrimSpace(outcome.Stage) == "" ||
		strings.TrimSpace(outcome.Stage) != strings.TrimSpace(diagnostic.ResolutionSource) ||
		strings.TrimSpace(outcome.SiteKind) == "" ||
		strings.TrimSpace(outcome.SiteKind) != strings.TrimSpace(diagnostic.FactFamily) ||
		strings.TrimSpace(outcome.FilePath) == "" ||
		strings.TrimSpace(outcome.FilePath) != strings.TrimSpace(diagnostic.FilePath) ||
		strings.TrimSpace(outcome.FileHash) != strings.TrimSpace(diagnostic.FileHash) ||
		outcome.Range.StartLine <= 0 ||
		outcome.Range.EndLine <= 0 ||
		outcome.Range.StartLine != diagnostic.StartLine ||
		outcome.Range.StartCol != diagnostic.StartCol ||
		outcome.Range.EndLine != diagnostic.EndLine ||
		outcome.Range.EndCol != diagnostic.EndCol ||
		strings.TrimSpace(outcome.RequestedName) == "" ||
		strings.TrimSpace(outcome.RequestedName) != strings.TrimSpace(diagnostic.TargetText) ||
		strings.TrimSpace(outcome.RequestedMeaning) == "" ||
		strings.TrimSpace(outcome.Language) == "" ||
		strings.TrimSpace(outcome.Proof.Kind) == "" ||
		strings.TrimSpace(outcome.Proof.Kind) != strings.TrimSpace(diagnostic.ProofKind) ||
		strings.TrimSpace(outcome.Stage) != strings.TrimSpace(diagnostic.Source) {
		return false
	}

	hasTarget := rawJSONValuePresent(outcome.Target)
	if hasTarget && !rawJSONObjectPresent(outcome.Target) {
		return false
	}
	hasAuthority := rawJSONValuePresent(outcome.Authority)
	if hasAuthority && !validStructuredResolutionAuthority(outcome) {
		return false
	}
	switch outcome.Status {
	case structuredResolutionStatusResolvedInternal:
		return hasTarget && strings.TrimSpace(outcome.Reason) == "" && !hasAuthority &&
			(outcome.Stage == structuredResolutionStageRepository || outcome.Stage == "intrinsic")
	case structuredResolutionStatusResolvedExternal:
		return hasTarget && strings.TrimSpace(outcome.Reason) == "" && hasAuthority &&
			outcome.Stage == structuredResolutionStageTypeScriptStandardLib
	case structuredResolutionStatusUnresolved:
		if hasTarget || strings.TrimSpace(outcome.Reason) == "" {
			return false
		}
		return outcome.Stage == structuredResolutionStageRepository && !hasAuthority ||
			outcome.Stage == structuredResolutionStageTypeScriptStandardLib && hasAuthority
	case structuredResolutionStatusCapabilityUnavailable:
		return !hasTarget && strings.TrimSpace(outcome.Reason) != "" && hasAuthority &&
			outcome.Stage == structuredResolutionStageTypeScriptStandardLib
	default:
		return false
	}
}

func validStructuredResolutionAuthority(outcome structuredResolutionOutcome) bool {
	if !rawJSONObjectPresent(outcome.Authority) {
		return false
	}
	var authority structuredResolutionAuthority
	if err := json.Unmarshal(outcome.Authority, &authority); err != nil ||
		!validStructuredTypeScriptAuthorityResult(authority) ||
		authority.SourceSiteID != outcome.SourceSiteID ||
		authority.Stage != outcome.Stage ||
		authority.FilePath != outcome.FilePath ||
		authority.FileHash != outcome.FileHash ||
		authority.Range.StartLine != outcome.Range.StartLine ||
		authority.Range.StartCol != outcome.Range.StartCol ||
		authority.Range.EndLine != outcome.Range.EndLine ||
		authority.Range.EndCol != outcome.Range.EndCol ||
		authority.SiteKind != outcome.SiteKind ||
		authority.RequestedName != outcome.RequestedName ||
		string(authority.RequestedMeaning) != outcome.RequestedMeaning {
		return false
	}
	switch outcome.Status {
	case structuredResolutionStatusResolvedExternal:
		return authority.Status == tsstdlib.LookupResolved
	case structuredResolutionStatusUnresolved:
		return authority.Status == tsstdlib.LookupProfileExcluded ||
			authority.Status == tsstdlib.LookupMeaningMismatch
	case structuredResolutionStatusCapabilityUnavailable:
		return authority.Status == tsstdlib.LookupCapabilityUnavailable
	default:
		return false
	}
}

func validStructuredTypeScriptAuthorityResult(authority structuredResolutionAuthority) bool {
	if authority.SourceSiteID == "" ||
		authority.Stage != structuredResolutionStageTypeScriptStandardLib ||
		authority.FilePath == "" ||
		authority.Range.StartLine <= 0 ||
		authority.Range.EndLine <= 0 ||
		authority.SiteKind == "" ||
		authority.RequestedName == "" ||
		authority.AuthorityKind != tsstdlib.AuthorityKind ||
		authority.TypeScriptVersion != tsstdlib.TypeScriptVersion ||
		authority.ProfileHash == "" ||
		authority.ConfigHash == "" ||
		!validStructuredTypeScriptCatalogProof(authority) {
		return false
	}
	switch authority.RequestedMeaning {
	case tsstdlib.MeaningValue, tsstdlib.MeaningType, tsstdlib.MeaningNamespace:
	default:
		return false
	}
	switch authority.Status {
	case tsstdlib.LookupResolved:
		return authority.Reason == "" && authority.ResolvedSymbolID != "" && len(authority.DeclarationRanges) > 0
	case tsstdlib.LookupCapabilityUnavailable:
		return authority.ResolvedSymbolID == "" && len(authority.DeclarationRanges) == 0
	case tsstdlib.LookupProfileExcluded:
		return authority.Reason == tsstdlib.ReasonProfileExcludes && authority.ResolvedSymbolID == "" && len(authority.DeclarationRanges) == 0
	case tsstdlib.LookupMeaningMismatch:
		return authority.Reason == tsstdlib.ReasonMeaningMismatch && authority.ResolvedSymbolID == "" && len(authority.DeclarationRanges) == 0
	default:
		return false
	}
}

func validStructuredTypeScriptCatalogProof(authority structuredResolutionAuthority) bool {
	switch authority.CatalogProofState {
	case tsstdlib.CatalogProofReady:
		if authority.AuthorityHash == "" || authority.CatalogHash == "" || authority.CatalogArtifactHash == "" {
			return false
		}
		switch authority.Status {
		case tsstdlib.LookupResolved:
			return authority.Reason == ""
		case tsstdlib.LookupProfileExcluded:
			return authority.Reason == tsstdlib.ReasonProfileExcludes
		case tsstdlib.LookupMeaningMismatch:
			return authority.Reason == tsstdlib.ReasonMeaningMismatch
		case tsstdlib.LookupCapabilityUnavailable:
			switch authority.Reason {
			case tsstdlib.ReasonDisabledByNoLib,
				tsstdlib.ReasonConfigInvalid,
				tsstdlib.ReasonConfigTopology,
				tsstdlib.ReasonConfigUnreadable:
				return true
			default:
				return false
			}
		default:
			return false
		}
	case tsstdlib.CatalogProofMissing:
		return authority.Status == tsstdlib.LookupCapabilityUnavailable &&
			authority.Reason == tsstdlib.ReasonCatalogMissing &&
			authority.AuthorityHash == "" &&
			authority.CatalogHash == "" &&
			authority.CatalogArtifactHash == ""
	case tsstdlib.CatalogProofRejected:
		return authority.Status == tsstdlib.LookupCapabilityUnavailable &&
			isStructuredTypeScriptCatalogRejection(authority.Reason) &&
			authority.AuthorityHash == "" &&
			authority.CatalogHash == "" &&
			authority.CatalogArtifactHash != ""
	default:
		return false
	}
}

func isStructuredTypeScriptCatalogRejection(reason tsstdlib.Reason) bool {
	switch reason {
	case tsstdlib.ReasonCatalogSchema,
		tsstdlib.ReasonCatalogVersion,
		tsstdlib.ReasonCatalogHash,
		tsstdlib.ReasonCatalogInputManifest:
		return true
	default:
		return false
	}
}

func isStructuredResolutionStatus(status string) bool {
	switch status {
	case structuredResolutionStatusResolvedInternal,
		structuredResolutionStatusResolvedExternal,
		structuredResolutionStatusUnresolved,
		structuredResolutionStatusCapabilityUnavailable:
		return true
	default:
		return false
	}
}

func rawJSONObjectPresent(value json.RawMessage) bool {
	trimmed := strings.TrimSpace(string(value))
	if trimmed == "" || trimmed == "null" {
		return false
	}
	var object map[string]json.RawMessage
	return json.Unmarshal(value, &object) == nil && object != nil
}

func rawJSONValuePresent(value json.RawMessage) bool {
	trimmed := strings.TrimSpace(string(value))
	return trimmed != "" && trimmed != "null"
}

func classifyDiagnostic(diagnostic Diagnostic) string {
	if diagnostic.Kind != DiagnosticUnresolvedReference {
		return DiagnosticClassificationUnclassified
	}
	target := strings.TrimSpace(diagnostic.TargetText)
	if target == "" {
		return DiagnosticClassificationUnclassified
	}
	target = strings.TrimPrefix(target, "*")
	if isGoBuiltinOrPredeclared(target) || isGoCompositeTypeText(target) {
		return DiagnosticClassificationBuiltin
	}
	if isGoTestFrameworkReference(target) {
		return DiagnosticClassificationTestFramework
	}
	if qualifier, ok := diagnosticTargetQualifier(target); ok {
		if goStandardLibraryQualifiers[qualifier] {
			return DiagnosticClassificationStandardLibrary
		}
		if externalLibraryQualifiers[qualifier] {
			return DiagnosticClassificationExternalLibrary
		}
		return DiagnosticClassificationInRepoUnresolved
	}
	return DiagnosticClassificationInRepoUnresolved
}

func actionabilityForDiagnosticClassification(classification string) string {
	switch classification {
	case DiagnosticClassificationBuiltin,
		DiagnosticClassificationStandardLibrary,
		DiagnosticClassificationTestFramework:
		return DiagnosticActionabilityNonActionable
	case DiagnosticClassificationInRepoUnresolved:
		return DiagnosticActionabilityAnalyzerGap
	case DiagnosticClassificationExternalLibrary,
		DiagnosticClassificationUnclassified:
		return DiagnosticActionabilityReview
	default:
		return DiagnosticActionabilityReview
	}
}

func diagnosticTargetQualifier(target string) (string, bool) {
	parts := strings.SplitN(target, ".", 2)
	if len(parts) != 2 {
		return "", false
	}
	qualifier := strings.TrimSpace(parts[0])
	return qualifier, qualifier != ""
}

func isGoBuiltinOrPredeclared(target string) bool {
	if goBuiltinOrPredeclared[target] {
		return true
	}
	if qualifier, ok := diagnosticTargetQualifier(target); ok {
		return goBuiltinOrPredeclared[qualifier]
	}
	return false
}

func isGoCompositeTypeText(target string) bool {
	switch {
	case strings.HasPrefix(target, "[]"),
		strings.HasPrefix(target, "map["),
		strings.HasPrefix(target, "chan "),
		strings.HasPrefix(target, "<-chan "),
		strings.Contains(target, "]"):
		return true
	default:
		return false
	}
}

func isGoTestFrameworkReference(target string) bool {
	if target == "testing.T" || target == "testing.B" || target == "testing.M" {
		return true
	}
	qualifier, ok := diagnosticTargetQualifier(target)
	if !ok {
		return false
	}
	if qualifier == "testing" {
		return true
	}
	if qualifier != "t" && qualifier != "b" {
		return false
	}
	member := strings.TrimPrefix(target, qualifier+".")
	return goTestingHelperMembers[member]
}

var goBuiltinOrPredeclared = map[string]bool{
	"any":        true,
	"append":     true,
	"bool":       true,
	"byte":       true,
	"cap":        true,
	"clear":      true,
	"close":      true,
	"comparable": true,
	"complex":    true,
	"complex64":  true,
	"complex128": true,
	"copy":       true,
	"delete":     true,
	"error":      true,
	"false":      true,
	"float32":    true,
	"float64":    true,
	"imag":       true,
	"int":        true,
	"int8":       true,
	"int16":      true,
	"int32":      true,
	"int64":      true,
	"iota":       true,
	"len":        true,
	"make":       true,
	"new":        true,
	"nil":        true,
	"panic":      true,
	"print":      true,
	"println":    true,
	"real":       true,
	"recover":    true,
	"rune":       true,
	"string":     true,
	"true":       true,
	"uint":       true,
	"uint8":      true,
	"uint16":     true,
	"uint32":     true,
	"uint64":     true,
	"uintptr":    true,
}

var goStandardLibraryQualifiers = map[string]bool{
	"archive":  true,
	"bufio":    true,
	"bytes":    true,
	"cmp":      true,
	"compress": true,
	"context":  true,
	"crypto":   true,
	"csv":      true,
	"database": true,
	"debug":    true,
	"embed":    true,
	"encoding": true,
	"errors":   true,
	"expvar":   true,
	"flag":     true,
	"fmt":      true,
	"go":       true,
	"hash":     true,
	"heap":     true,
	"html":     true,
	"http":     true,
	"image":    true,
	"io":       true,
	"json":     true,
	"log":      true,
	"maps":     true,
	"math":     true,
	"mime":     true,
	"net":      true,
	"os":       true,
	"path":     true,
	"filepath": true,
	"rand":     true,
	"reflect":  true,
	"regexp":   true,
	"runtime":  true,
	"slices":   true,
	"sort":     true,
	"strconv":  true,
	"strings":  true,
	"sync":     true,
	"syscall":  true,
	"template": true,
	"testing":  true,
	"text":     true,
	"time":     true,
	"unicode":  true,
	"unsafe":   true,
	"url":      true,
	"utf8":     true,
	"xml":      true,
}

var goTestingHelperMembers = map[string]bool{
	"Cleanup":  true,
	"Error":    true,
	"Errorf":   true,
	"Fail":     true,
	"FailNow":  true,
	"Failed":   true,
	"Fatal":    true,
	"Fatalf":   true,
	"Helper":   true,
	"Log":      true,
	"Logf":     true,
	"Name":     true,
	"Parallel": true,
	"Run":      true,
	"Setenv":   true,
	"Skip":     true,
	"SkipNow":  true,
	"Skipf":    true,
	"Skipped":  true,
	"TempDir":  true,
}

var externalLibraryQualifiers = map[string]bool{
	"anthropic": true,
	"assert":    true,
	"cobra":     true,
	"gin":       true,
	"grpc":      true,
	"jwt":       true,
	"openai":    true,
	"require":   true,
	"uuid":      true,
	"yaml":      true,
	"zap":       true,
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
