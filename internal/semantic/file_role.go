package semantic

import "strings"

const (
	FileRoleProperty = "fileRole"
)

type FileRole string

const (
	FileRoleModel           FileRole = "model"
	FileRoleContractModel   FileRole = "contract_model"
	FileRoleHelper          FileRole = "helper"
	FileRoleStorageHelper   FileRole = "storage_helper"
	FileRoleConfig          FileRole = "config"
	FileRoleAdapter         FileRole = "adapter"
	FileRoleFallbackAdapter FileRole = "fallback_adapter"
	FileRoleTestHelper      FileRole = "test_helper"
	FileRoleAnalyzerHelper  FileRole = "analyzer_helper"
	FileRoleParserModel     FileRole = "parser_model"
	FileRoleRuntimeModel    FileRole = "runtime_model"
	FileRoleUnknown         FileRole = "unknown"
)

var FileRoles = []FileRole{
	FileRoleModel,
	FileRoleContractModel,
	FileRoleHelper,
	FileRoleStorageHelper,
	FileRoleConfig,
	FileRoleAdapter,
	FileRoleFallbackAdapter,
	FileRoleTestHelper,
	FileRoleAnalyzerHelper,
	FileRoleParserModel,
	FileRoleRuntimeModel,
	FileRoleUnknown,
}

type FileRoleClassification struct {
	Role   FileRole
	Source string
}

func FileRoleStrings() []string {
	out := make([]string, 0, len(FileRoles))
	for _, role := range FileRoles {
		out = append(out, string(role))
	}
	return out
}

func ClassifyFileRole(filePath string, kind string, appLayer string, functionalArea string) FileRoleClassification {
	path := normalizePath(filePath)
	if path == "" {
		return FileRoleClassification{Role: FileRoleUnknown, Source: "missing_file_path"}
	}

	base := baseName(path)
	layer := appLayerFromString(appLayer)
	area := functionalAreaFromString(functionalArea)

	if isTestFileRolePath(path, base, kind, layer) {
		return FileRoleClassification{Role: FileRoleTestHelper, Source: "test_helper_path"}
	}
	if isFileRoleConfigPath(path, base, area) {
		return FileRoleClassification{Role: FileRoleConfig, Source: "config_path"}
	}
	if isFallbackAdapterRolePath(path, base) {
		return FileRoleClassification{Role: FileRoleFallbackAdapter, Source: "fallback_adapter_path"}
	}
	if isAdapterRolePath(path) {
		return FileRoleClassification{Role: FileRoleAdapter, Source: "adapter_path"}
	}
	if isStorageHelperRolePath(path, base, area) {
		return FileRoleClassification{Role: FileRoleStorageHelper, Source: "storage_helper_path"}
	}
	if isAnalyzerHelperRolePath(path, area) {
		return FileRoleClassification{Role: FileRoleAnalyzerHelper, Source: "analyzer_helper_path"}
	}
	if isContractModelRolePath(path, base, area) {
		return FileRoleClassification{Role: FileRoleContractModel, Source: "contract_model_path"}
	}
	if isParserModelRolePath(path, base, area) {
		return FileRoleClassification{Role: FileRoleParserModel, Source: "parser_model_path"}
	}
	if isRuntimeModelRolePath(path, base, area) {
		return FileRoleClassification{Role: FileRoleRuntimeModel, Source: "runtime_model_path"}
	}
	if isModelRolePath(base) {
		return FileRoleClassification{Role: FileRoleModel, Source: "model_filename"}
	}
	if isHelperRolePath(path, base) {
		return FileRoleClassification{Role: FileRoleHelper, Source: "helper_path"}
	}

	return FileRoleClassification{Role: FileRoleUnknown, Source: "unmatched_path"}
}

func FileRoleDefinitions() []TermDefinition {
	return []TermDefinition{
		{Key: string(FileRoleModel), DisplayLabel: "Model", CLILabel: "model", WebLabel: "Model", Description: "General data model or value object file."},
		{Key: string(FileRoleContractModel), DisplayLabel: "Contract Model", CLILabel: "contract-model", WebLabel: "Contract Model", Description: "Data model used to carry API, group, or cross-surface contracts."},
		{Key: string(FileRoleHelper), DisplayLabel: "Helper", CLILabel: "helper", WebLabel: "Helper", Description: "Support code with small reusable helpers."},
		{Key: string(FileRoleStorageHelper), DisplayLabel: "Storage Helper", CLILabel: "storage-helper", WebLabel: "Storage Helper", Description: "Repository, path, or storage support helper."},
		{Key: string(FileRoleConfig), DisplayLabel: "Config", CLILabel: "config", WebLabel: "Config", Description: "Runtime or repository configuration model."},
		{Key: string(FileRoleAdapter), DisplayLabel: "Adapter", CLILabel: "adapter", WebLabel: "Adapter", Description: "Adapter between Anvien and an external/runtime backend."},
		{Key: string(FileRoleFallbackAdapter), DisplayLabel: "Fallback Adapter", CLILabel: "fallback-adapter", WebLabel: "Fallback Adapter", Description: "Fallback adapter used when a primary runtime backend is unavailable."},
		{Key: string(FileRoleTestHelper), DisplayLabel: "Test Helper", CLILabel: "test-helper", WebLabel: "Test Helper", Description: "Test utility or fixture helper file."},
		{Key: string(FileRoleAnalyzerHelper), DisplayLabel: "Analyzer Helper", CLILabel: "analyzer-helper", WebLabel: "Analyzer Helper", Description: "Analyzer support helper for detection, expansion, or framework metadata."},
		{Key: string(FileRoleParserModel), DisplayLabel: "Parser Model", CLILabel: "parser-model", WebLabel: "Parser Model", Description: "Parser, ScopeIR, range, fact, or metric model file."},
		{Key: string(FileRoleRuntimeModel), DisplayLabel: "Runtime Model", CLILabel: "runtime-model", WebLabel: "Runtime Model", Description: "Runtime/session type, error, or state model file."},
		{Key: string(FileRoleUnknown), DisplayLabel: "Unknown", CLILabel: "unknown", WebLabel: "Unknown", Description: "Insufficient evidence for a stable file role."},
	}
}

func isTestFileRolePath(path string, base string, kind string, layer AppLayer) bool {
	return kind == "test" ||
		isTestAppLayer(layer) ||
		hasPathPrefix(path, "internal/testutil") ||
		strings.Contains(base, "_test_helper")
}

func isFileRoleConfigPath(path string, base string, area FunctionalArea) bool {
	if area == FunctionalAreaConfiguration {
		return true
	}
	if hasPathPrefix(path, "internal/repo") {
		return base == "settings.go" || base == "runtime_config.go"
	}
	return strings.Contains(base, "config") || strings.Contains(base, "settings")
}

func isFallbackAdapterRolePath(path string, base string) bool {
	if !isAdapterRolePath(path) {
		return false
	}
	return strings.Contains(base, "default") || strings.Contains(base, "fallback") ||
		strings.HasSuffix(path, "/runner_default.go")
}

func isAdapterRolePath(path string) bool {
	return hasPathPrefix(path, "internal/lbugnative") ||
		hasPathPrefix(path, "internal/lbug")
}

func isStorageHelperRolePath(path string, base string, area FunctionalArea) bool {
	if area != FunctionalAreaStorage {
		return false
	}
	return strings.Contains(base, "path") || strings.Contains(base, "repo") || strings.Contains(base, "store")
}

func isAnalyzerHelperRolePath(path string, area FunctionalArea) bool {
	if area != FunctionalAreaAnalyzer {
		return false
	}
	return hasPathPrefix(path, "internal/frameworks") ||
		hasPathPrefix(path, "internal/cobol")
}

func isContractModelRolePath(path string, base string, area FunctionalArea) bool {
	return base == "types.go" && (area == FunctionalAreaQuery || hasPathPrefix(path, "internal/group"))
}

func isParserModelRolePath(path string, base string, area FunctionalArea) bool {
	if area != FunctionalAreaProviders {
		return false
	}
	switch base {
	case "metrics.go", "facts.go", "range.go", "types.go":
		return true
	default:
		return false
	}
}

func isRuntimeModelRolePath(path string, base string, area FunctionalArea) bool {
	if area != FunctionalAreaSession && area != FunctionalAreaRuntime {
		return false
	}
	return base == "types.go" || base == "error.go" || strings.Contains(base, "state")
}

func isModelRolePath(base string) bool {
	switch base {
	case "types.go", "metrics.go", "facts.go", "range.go", "model.go", "models.go":
		return true
	default:
		return false
	}
}

func isHelperRolePath(path string, base string) bool {
	if hasPathPrefix(path, "internal/resolution") && base == "source_site.go" {
		return true
	}
	if hasPathPrefix(path, "internal/scopeir") && strings.Contains(base, "sort") {
		return true
	}
	if hasPathPrefix(path, "internal/cli") && strings.Contains(base, "error") {
		return true
	}
	return strings.Contains(base, "helper") || strings.Contains(base, "util")
}
