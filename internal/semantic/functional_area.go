package semantic

import "strings"

type FunctionalArea string

const (
	FunctionalAreaResolution    FunctionalArea = "resolution"
	FunctionalAreaGraphHealth   FunctionalArea = "graph_health"
	FunctionalAreaQuery         FunctionalArea = "query"
	FunctionalAreaMCP           FunctionalArea = "mcp"
	FunctionalAreaWebGraphUI    FunctionalArea = "web_graph_ui"
	FunctionalAreaLayout        FunctionalArea = "layout"
	FunctionalAreaContracts     FunctionalArea = "contracts"
	FunctionalAreaProviders     FunctionalArea = "providers"
	FunctionalAreaRuntime       FunctionalArea = "runtime"
	FunctionalAreaAnalyzer      FunctionalArea = "analyzer"
	FunctionalAreaSession       FunctionalArea = "session"
	FunctionalAreaLauncher      FunctionalArea = "launcher"
	FunctionalAreaCLI           FunctionalArea = "cli"
	FunctionalAreaReporting     FunctionalArea = "reporting"
	FunctionalAreaAPI           FunctionalArea = "api"
	FunctionalAreaStorage       FunctionalArea = "storage"
	FunctionalAreaEmbeddings    FunctionalArea = "embeddings"
	FunctionalAreaConfiguration FunctionalArea = "configuration"
	FunctionalAreaDocumentation FunctionalArea = "documentation"
	FunctionalAreaMixed         FunctionalArea = "mixed"
	FunctionalAreaUnknown       FunctionalArea = "unknown"
)

const (
	FunctionalAreaProperty       = "functionalArea"
	FunctionalAreaSourceProperty = "functionalAreaSource"
)

var FunctionalAreas = []FunctionalArea{
	FunctionalAreaResolution,
	FunctionalAreaGraphHealth,
	FunctionalAreaQuery,
	FunctionalAreaMCP,
	FunctionalAreaWebGraphUI,
	FunctionalAreaLayout,
	FunctionalAreaContracts,
	FunctionalAreaProviders,
	FunctionalAreaRuntime,
	FunctionalAreaAnalyzer,
	FunctionalAreaSession,
	FunctionalAreaLauncher,
	FunctionalAreaCLI,
	FunctionalAreaReporting,
	FunctionalAreaAPI,
	FunctionalAreaStorage,
	FunctionalAreaEmbeddings,
	FunctionalAreaConfiguration,
	FunctionalAreaDocumentation,
	FunctionalAreaMixed,
	FunctionalAreaUnknown,
}

type FunctionalAreaClassification struct {
	Area   FunctionalArea
	Source string
}

func FunctionalAreaStrings() []string {
	out := make([]string, 0, len(FunctionalAreas))
	for _, area := range FunctionalAreas {
		out = append(out, string(area))
	}
	return out
}

func ClassifyFunctionalArea(filePath string) FunctionalAreaClassification {
	path := normalizePath(filePath)
	if path == "" {
		return FunctionalAreaClassification{Area: FunctionalAreaUnknown, Source: "missing_file_path"}
	}
	base := baseName(path)

	if isReportingAreaPath(path) {
		return FunctionalAreaClassification{Area: FunctionalAreaReporting, Source: "reporting_path"}
	}
	if isDocsPath(path, base) {
		return FunctionalAreaClassification{Area: FunctionalAreaDocumentation, Source: "docs_path"}
	}
	if isConfigPath(path, base) {
		return FunctionalAreaClassification{Area: FunctionalAreaConfiguration, Source: "config_path"}
	}
	if isGeneratedContractPath(path) || isAPIContractPath(path) {
		return FunctionalAreaClassification{Area: FunctionalAreaContracts, Source: "contract_path"}
	}
	if isResolutionAreaPath(path) {
		return FunctionalAreaClassification{Area: FunctionalAreaResolution, Source: "resolution_path"}
	}
	if isGraphHealthAreaPath(path) {
		return FunctionalAreaClassification{Area: FunctionalAreaGraphHealth, Source: "graph_health_path"}
	}
	if isQueryAreaPath(path) {
		return FunctionalAreaClassification{Area: FunctionalAreaQuery, Source: "query_path"}
	}
	if isMCPAreaPath(path) {
		return FunctionalAreaClassification{Area: FunctionalAreaMCP, Source: "mcp_path"}
	}
	if isAPIPath(path) || isFrontendAPIClientPath(path) {
		return FunctionalAreaClassification{Area: FunctionalAreaAPI, Source: "api_path"}
	}
	if isLayoutAreaPath(path) {
		return FunctionalAreaClassification{Area: FunctionalAreaLayout, Source: "layout_path"}
	}
	if isWebGraphUIAreaPath(path) {
		return FunctionalAreaClassification{Area: FunctionalAreaWebGraphUI, Source: "web_graph_ui_path"}
	}
	if isProviderAreaPath(path) {
		return FunctionalAreaClassification{Area: FunctionalAreaProviders, Source: "providers_path"}
	}
	if isSessionAreaPath(path) {
		return FunctionalAreaClassification{Area: FunctionalAreaSession, Source: "session_path"}
	}
	if isStorageAreaPath(path) {
		return FunctionalAreaClassification{Area: FunctionalAreaStorage, Source: "storage_path"}
	}
	if isEmbeddingAreaPath(path) {
		return FunctionalAreaClassification{Area: FunctionalAreaEmbeddings, Source: "embeddings_path"}
	}
	if isRuntimeAreaPath(path) {
		return FunctionalAreaClassification{Area: FunctionalAreaRuntime, Source: "runtime_path"}
	}
	if isCLIAreaPath(path) {
		return FunctionalAreaClassification{Area: FunctionalAreaCLI, Source: "cli_path"}
	}
	if isLauncherAreaPath(path) {
		return FunctionalAreaClassification{Area: FunctionalAreaLauncher, Source: "launcher_path"}
	}
	if isAnalyzerAreaPath(path) {
		return FunctionalAreaClassification{Area: FunctionalAreaAnalyzer, Source: "analyzer_path"}
	}
	return FunctionalAreaClassification{Area: FunctionalAreaUnknown, Source: "unmatched_path"}
}

func FunctionalAreaDefinitions() []TermDefinition {
	return []TermDefinition{
		{Key: string(FunctionalAreaResolution), DisplayLabel: "Resolution", CLILabel: "resolution", WebLabel: "Resolution", Description: "Reference binding and unresolved-reference detection."},
		{Key: string(FunctionalAreaGraphHealth), DisplayLabel: "Graph Health", CLILabel: "graph-health", WebLabel: "Graph Health", Description: "Topology health, diagnostic confidence, and graph quality policy."},
		{Key: string(FunctionalAreaQuery), DisplayLabel: "Query", CLILabel: "query", WebLabel: "Query", Description: "Graph query, retrieval, and benchmark surfaces."},
		{Key: string(FunctionalAreaMCP), DisplayLabel: "MCP", CLILabel: "mcp", WebLabel: "MCP", Description: "Model Context Protocol tools, resources, and integrations."},
		{Key: string(FunctionalAreaWebGraphUI), DisplayLabel: "Web Graph UI", CLILabel: "web-graph-ui", WebLabel: "Web Graph UI", Description: "Frontend graph panels, filters, and visual controls."},
		{Key: string(FunctionalAreaLayout), DisplayLabel: "Layout", CLILabel: "layout", WebLabel: "Layout", Description: "Graph placement, rendering geometry, and layout controls."},
		{Key: string(FunctionalAreaContracts), DisplayLabel: "Contracts", CLILabel: "contracts", WebLabel: "Contracts", Description: "Generated and shared cross-surface contracts."},
		{Key: string(FunctionalAreaProviders), DisplayLabel: "Providers", CLILabel: "providers", WebLabel: "Providers", Description: "Language provider extraction and ScopeIR conversion."},
		{Key: string(FunctionalAreaRuntime), DisplayLabel: "Runtime", CLILabel: "runtime", WebLabel: "Runtime", Description: "Runtime process, local server, and execution management."},
		{Key: string(FunctionalAreaAnalyzer), DisplayLabel: "Analyzer", CLILabel: "analyzer", WebLabel: "Analyzer", Description: "Analyze pipeline phases and graph enrichment."},
		{Key: string(FunctionalAreaSession), DisplayLabel: "Session", CLILabel: "session", WebLabel: "Session", Description: "Interactive session, job, and chat runtime services."},
		{Key: string(FunctionalAreaLauncher), DisplayLabel: "Launcher", CLILabel: "launcher", WebLabel: "Launcher", Description: "Desktop launcher and startup flow."},
		{Key: string(FunctionalAreaCLI), DisplayLabel: "CLI", CLILabel: "cli", WebLabel: "CLI", Description: "Command-line entrypoints and command behavior."},
		{Key: string(FunctionalAreaReporting), DisplayLabel: "Reporting", CLILabel: "reporting", WebLabel: "Reporting", Description: "Reports, ledgers, and problem records."},
		{Key: string(FunctionalAreaAPI), DisplayLabel: "API", CLILabel: "api", WebLabel: "API", Description: "HTTP API and API client behavior."},
		{Key: string(FunctionalAreaStorage), DisplayLabel: "Storage", CLILabel: "storage", WebLabel: "Storage", Description: "Graph storage, LadybugDB export/load, and repository metadata."},
		{Key: string(FunctionalAreaEmbeddings), DisplayLabel: "Embeddings", CLILabel: "embeddings", WebLabel: "Embeddings", Description: "Embedding generation, indexing, and semantic search support."},
		{Key: string(FunctionalAreaConfiguration), DisplayLabel: "Configuration", CLILabel: "configuration", WebLabel: "Configuration", Description: "Project and build configuration."},
		{Key: string(FunctionalAreaDocumentation), DisplayLabel: "Documentation", CLILabel: "documentation", WebLabel: "Documentation", Description: "Documentation content."},
		{Key: string(FunctionalAreaMixed), DisplayLabel: "Mixed", CLILabel: "mixed", WebLabel: "Mixed", Description: "Process or community evidence spanning more than one Functional Area."},
		{Key: string(FunctionalAreaUnknown), DisplayLabel: "Unknown", CLILabel: "unknown", WebLabel: "Unknown", Description: "Insufficient evidence for a stable Functional Area."},
	}
}

func baseName(path string) string {
	if slash := strings.LastIndex(path, "/"); slash >= 0 {
		return path[slash+1:]
	}
	return path
}

func isResolutionAreaPath(path string) bool {
	return hasPathPrefix(path, "internal/resolution")
}

func isGraphHealthAreaPath(path string) bool {
	return hasPathPrefix(path, "internal/graphhealth") ||
		hasPathPrefix(path, "internal/graphaccuracy")
}

func isQueryAreaPath(path string) bool {
	return hasPathPrefix(path, "internal/group") ||
		strings.Contains(path, "query_profile") ||
		path == "internal/httpapi/query.go" ||
		path == "internal/httpapi/query_test.go" ||
		path == "internal/cli/tool_command.go"
}

func isMCPAreaPath(path string) bool {
	return hasPathPrefix(path, "internal/mcp")
}

func isLayoutAreaPath(path string) bool {
	return path == "avmatrix-web/src/lib/graph-adapter.ts" ||
		path == "avmatrix-web/src/hooks/usesigma.ts" ||
		strings.Contains(path, "avmatrix-web/src/lib/graph-edge") ||
		strings.Contains(path, "avmatrix-web/src/lib/graph-links") ||
		strings.Contains(path, "avmatrix-web/src/lib/selected-graph-context")
}

func isWebGraphUIAreaPath(path string) bool {
	return hasPathPrefix(path, "avmatrix-web/src/components") ||
		hasPathPrefix(path, "avmatrix-web/src/hooks/app-state") ||
		path == "avmatrix-web/src/app.tsx" ||
		path == "avmatrix-web/src/main.tsx" ||
		path == "avmatrix-web/src/lib/graph-health-filters.ts"
}

func isProviderAreaPath(path string) bool {
	return hasPathPrefix(path, "internal/providers") ||
		hasPathPrefix(path, "internal/parser") ||
		hasPathPrefix(path, "internal/scopeir")
}

func isSessionAreaPath(path string) bool {
	return hasPathPrefix(path, "internal/session") ||
		pathContains(path, "avmatrix-web/src/hooks/chat-runtime") ||
		path == "avmatrix-web/src/core/llm/session-client.ts"
}

func isStorageAreaPath(path string) bool {
	return strings.HasPrefix(path, "internal/lbug") ||
		hasPathPrefix(path, "internal/repo") ||
		hasPathPrefix(path, "internal/graph")
}

func isEmbeddingAreaPath(path string) bool {
	return hasPathPrefix(path, "internal/embeddings") ||
		path == "internal/httpapi/embed.go" ||
		path == "internal/httpapi/embed_test.go"
}

func isRuntimeAreaPath(path string) bool {
	return hasPathPrefix(path, "internal/httpapi") ||
		hasPathPrefix(path, "internal/logging") ||
		hasPathPrefix(path, "internal/version")
}

func isCLIAreaPath(path string) bool {
	return hasPathPrefix(path, "internal/cli") ||
		hasPathPrefix(path, "cmd/avmatrix") ||
		hasPathPrefix(path, "cmd/access-candidate-audit") ||
		hasPathPrefix(path, "cmd/graph-accuracy-probe") ||
		hasPathPrefix(path, "cmd/property-access-audit")
}

func isLauncherAreaPath(path string) bool {
	return hasPathPrefix(path, "avmatrix-launcher") ||
		path == "start-avmatrix.html"
}

func isReportingAreaPath(path string) bool {
	return hasPathPrefix(path, "reports")
}

func isAnalyzerAreaPath(path string) bool {
	return hasPathPrefix(path, "internal/analyze") ||
		hasPathPrefix(path, "internal/scanner") ||
		hasPathPrefix(path, "internal/structure") ||
		hasPathPrefix(path, "internal/routes") ||
		hasPathPrefix(path, "internal/tools") ||
		hasPathPrefix(path, "internal/orm") ||
		hasPathPrefix(path, "internal/mro") ||
		hasPathPrefix(path, "internal/communities") ||
		hasPathPrefix(path, "internal/processes") ||
		hasPathPrefix(path, "internal/cobol") ||
		hasPathPrefix(path, "internal/documents") ||
		hasPathPrefix(path, "internal/frameworks") ||
		hasPathPrefix(path, "internal/ignore")
}
