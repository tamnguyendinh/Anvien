package semantic

import "github.com/tamnguyendinh/anvien/internal/graph"

const (
	SchemaVersion = "semantic_app_functional_v1"

	StatusComplete        = "complete"
	StatusStaleIncomplete = "stale_incomplete"
)

type FieldStatus struct {
	Field                         string `json:"field"`
	Status                        string `json:"status"`
	Required                      bool   `json:"required"`
	TotalNodes                    int    `json:"totalNodes"`
	NodesWithField                int    `json:"nodesWithField"`
	MissingNodes                  int    `json:"missingNodes"`
	UnknownNodes                  int    `json:"unknownNodes"`
	NodesWithSource               int    `json:"nodesWithSource"`
	MissingSourceNodes            int    `json:"missingSourceNodes"`
	StaleIncompleteSchemaEvidence bool   `json:"staleIncompleteSchemaEvidence"`
	Message                       string `json:"message,omitempty"`
}

type GraphStatus struct {
	SchemaVersion  string      `json:"schemaVersion"`
	AppLayer       FieldStatus `json:"appLayer"`
	FunctionalArea FieldStatus `json:"functionalArea"`
}

type TermDefinition struct {
	Key          string `json:"key"`
	DisplayLabel string `json:"displayLabel"`
	CLILabel     string `json:"cliLabel"`
	WebLabel     string `json:"webLabel"`
	Description  string `json:"description"`
}

func StatusValues() []string {
	return []string{StatusComplete, StatusStaleIncomplete}
}

func GraphSemanticStatus(g *graph.Graph) GraphStatus {
	status := GraphStatus{
		SchemaVersion: SchemaVersion,
		AppLayer: FieldStatus{
			Field:    AppLayerProperty,
			Status:   StatusComplete,
			Required: true,
		},
		FunctionalArea: FieldStatus{
			Field:    FunctionalAreaProperty,
			Status:   StatusComplete,
			Required: true,
		},
	}
	if g == nil {
		return status
	}

	status.AppLayer = semanticFieldStatus(g, AppLayerProperty, AppLayerSourceProperty, string(AppLayerUnknown))
	status.FunctionalArea = semanticFieldStatus(g, FunctionalAreaProperty, FunctionalAreaSourceProperty, string(FunctionalAreaUnknown))
	return status
}

func semanticFieldStatus(g *graph.Graph, field string, sourceField string, unknownValue string) FieldStatus {
	status := FieldStatus{
		Field:      field,
		Status:     StatusComplete,
		Required:   true,
		TotalNodes: len(g.Nodes),
	}
	for _, node := range g.Nodes {
		if node.Properties == nil {
			status.MissingNodes++
			status.MissingSourceNodes++
			continue
		}
		value, hasValue := stringValue(node.Properties[field])
		if hasValue {
			status.NodesWithField++
			if value == unknownValue {
				status.UnknownNodes++
			}
		} else {
			status.MissingNodes++
		}
		if _, hasSource := stringValue(node.Properties[sourceField]); hasSource {
			status.NodesWithSource++
		} else {
			status.MissingSourceNodes++
		}
	}

	if status.MissingNodes > 0 || status.MissingSourceNodes > 0 {
		status.Status = StatusStaleIncomplete
		status.StaleIncompleteSchemaEvidence = true
		status.Message = "Graph semantic metadata is incomplete; run avmatrix analyze --force to refresh graph evidence."
	}
	return status
}

func AppLayerDefinitions() []TermDefinition {
	return []TermDefinition{
		{Key: string(AppLayerBackend), DisplayLabel: "Backend", CLILabel: "backend", WebLabel: "Backend", Description: "Backend analyzer, storage, and runtime code."},
		{Key: string(AppLayerAPI), DisplayLabel: "API Layer", CLILabel: "api", WebLabel: "API", Description: "Server-side API handler and protocol surface."},
		{Key: string(AppLayerFrontend), DisplayLabel: "Frontend", CLILabel: "frontend", WebLabel: "Frontend", Description: "Browser UI source code."},
		{Key: string(AppLayerCLILauncher), DisplayLabel: "CLI Launcher", CLILabel: "cli-launcher", WebLabel: "CLI Launcher", Description: "CLI, launcher, and startup surfaces."},
		{Key: string(AppLayerSharedContract), DisplayLabel: "Shared Contract", CLILabel: "shared-contract", WebLabel: "Shared Contract", Description: "Contract data shared across product surfaces."},
		{Key: string(AppLayerAPIContract), DisplayLabel: "API Contract", CLILabel: "api-contract", WebLabel: "API Contract", Description: "API schema, generated contract, or contract source."},
		{Key: string(AppLayerAPISharedContract), DisplayLabel: "API Shared Contract", CLILabel: "api-shared-contract", WebLabel: "API Shared Contract", Description: "Contract surface that is both API-owned and shared."},
		{Key: string(AppLayerFrontendAPIClient), DisplayLabel: "Frontend API Client", CLILabel: "frontend-api-client", WebLabel: "Frontend API Client", Description: "Frontend client code that talks to backend APIs."},
		{Key: string(AppLayerBackendTest), DisplayLabel: "Backend Test", CLILabel: "backend-test", WebLabel: "Backend Test", Description: "Backend test and fixture code."},
		{Key: string(AppLayerFrontendTest), DisplayLabel: "Frontend Test", CLILabel: "frontend-test", WebLabel: "Frontend Test", Description: "Frontend unit, integration, or e2e test code."},
		{Key: string(AppLayerAPITest), DisplayLabel: "API Test", CLILabel: "api-test", WebLabel: "API Test", Description: "API or contract test code."},
		{Key: string(AppLayerGeneratedContract), DisplayLabel: "Generated Contract", CLILabel: "generated-contract", WebLabel: "Generated Contract", Description: "Generated cross-surface contract artifacts."},
		{Key: string(AppLayerDocs), DisplayLabel: "Docs", CLILabel: "docs", WebLabel: "Docs", Description: "Documentation and report files."},
		{Key: string(AppLayerConfig), DisplayLabel: "Config", CLILabel: "config", WebLabel: "Config", Description: "Configuration files."},
		{Key: string(AppLayerGenerated), DisplayLabel: "Generated", CLILabel: "generated", WebLabel: "Generated", Description: "Generated build or output files outside contract surfaces."},
		{Key: string(AppLayerMixed), DisplayLabel: "Mixed", CLILabel: "mixed", WebLabel: "Mixed", Description: "Process or community evidence spanning more than one App Layer."},
		{Key: string(AppLayerUnknown), DisplayLabel: "Unknown", CLILabel: "unknown", WebLabel: "Unknown", Description: "Insufficient evidence for a stable App Layer."},
	}
}

func SemanticTermDefinitions() []TermDefinition {
	return []TermDefinition{
		{Key: "app_layer", DisplayLabel: "App Layer", CLILabel: "app-layer", WebLabel: "App Layer", Description: "Primary product layer assigned to each graph node."},
		{Key: "functional_area", DisplayLabel: "Functional Area", CLILabel: "functional-area", WebLabel: "Functional Area", Description: "High-confidence functional ownership under an App Layer."},
		{Key: "api_layer", DisplayLabel: "API Layer", CLILabel: "api-layer", WebLabel: "API Layer", Description: "Server-facing API handlers and API protocol surface."},
		{Key: "api_contract", DisplayLabel: "API Contract", CLILabel: "api-contract", WebLabel: "API Contract", Description: "Schema or generated contract shared by API consumers."},
		{Key: "frontend_api_client", DisplayLabel: "Frontend API Client", CLILabel: "frontend-api-client", WebLabel: "Frontend API Client", Description: "Frontend client surface that calls backend APIs."},
		{Key: "resolution_gap", DisplayLabel: "Resolution Gap", CLILabel: "resolution-gap", WebLabel: "Resolution Gap", Description: "Persisted evidence that a source reference could not be resolved."},
		{Key: "unresolved_symbol", DisplayLabel: "Unresolved Symbol", CLILabel: "unresolved-symbol", WebLabel: "Unresolved Symbol", Description: "Unresolved target text retained without claiming a resolved in-repo symbol."},
		{Key: "analyzer_gap", DisplayLabel: "Analyzer Gap", CLILabel: "analyzer-gap", WebLabel: "Analyzer Gap", Description: "In-repo evidence that the analyzer could not resolve confidently."},
		{Key: "external_reference", DisplayLabel: "External Reference", CLILabel: "external-reference", WebLabel: "External Reference", Description: "Reference expected to resolve outside the indexed repository."},
		{Key: "non_actionable_reference", DisplayLabel: "Non-actionable Reference", CLILabel: "non-actionable-reference", WebLabel: "Non-actionable Reference", Description: "Builtin, standard library, fixture, or otherwise non-actionable unresolved reference."},
	}
}

func stringValue(value any) (string, bool) {
	text, ok := value.(string)
	if !ok || text == "" {
		return "", false
	}
	return text, true
}
