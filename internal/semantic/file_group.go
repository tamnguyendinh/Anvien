package semantic

import "strings"

const (
	FileGroupProperty = "fileGroup"
)

type FileGroup string

const (
	FileGroupBackendSupportModelHelper FileGroup = "backend_support_model_helper"
)

var FileGroups = []FileGroup{
	FileGroupBackendSupportModelHelper,
}

type FileGroupClassification struct {
	Group  FileGroup
	Source string
}

func FileGroupStrings() []string {
	out := make([]string, 0, len(FileGroups))
	for _, group := range FileGroups {
		out = append(out, string(group))
	}
	return out
}

func ClassifyFileGroup(filePath string, kind string, appLayer string, fileRole string) FileGroupClassification {
	if normalizePath(filePath) == "" {
		return FileGroupClassification{Source: "missing_file_path"}
	}
	if strings.ToLower(strings.TrimSpace(kind)) != "source" {
		return FileGroupClassification{Source: "non_source_file_kind"}
	}
	if appLayerFromString(appLayer) != AppLayerBackend {
		return FileGroupClassification{Source: "non_backend_app_layer"}
	}
	role := FileRole(strings.ToLower(strings.TrimSpace(fileRole)))
	if isBackendSupportModelHelperRole(role) {
		return FileGroupClassification{Group: FileGroupBackendSupportModelHelper, Source: "backend_support_model_helper_role"}
	}
	return FileGroupClassification{Source: "unmatched_file_role"}
}

func FileGroupDefinitions() []TermDefinition {
	return []TermDefinition{
		{
			Key:          string(FileGroupBackendSupportModelHelper),
			DisplayLabel: "Backend support/model/helper files",
			CLILabel:     "backend-support-model-helper",
			WebLabel:     "Backend support/model/helper files",
			Description:  "Backend source files whose role is support, model, helper, adapter, config, runtime, parser, contract, analyzer, storage, or test utility.",
		},
	}
}

func FileGroupLabel(value string) string {
	group := FileGroup(strings.ToLower(strings.TrimSpace(value)))
	for _, definition := range FileGroupDefinitions() {
		if definition.Key == string(group) {
			return definition.DisplayLabel
		}
	}
	return ""
}

func isBackendSupportModelHelperRole(role FileRole) bool {
	switch role {
	case FileRoleModel,
		FileRoleContractModel,
		FileRoleHelper,
		FileRoleStorageHelper,
		FileRoleConfig,
		FileRoleAdapter,
		FileRoleFallbackAdapter,
		FileRoleTestHelper,
		FileRoleAnalyzerHelper,
		FileRoleParserModel,
		FileRoleRuntimeModel:
		return true
	default:
		return false
	}
}
