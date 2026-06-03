package semantic

import "testing"

func TestFileGroupDefinitionsAreStable(t *testing.T) {
	if got := string(FileGroupBackendSupportModelHelper); got != "backend_support_model_helper" {
		t.Fatalf("file group key = %q, want backend_support_model_helper", got)
	}
	if got := FileGroupLabel(string(FileGroupBackendSupportModelHelper)); got != "Backend support/model/helper files" {
		t.Fatalf("file group label = %q, want Backend support/model/helper files", got)
	}
	definitions := FileGroupDefinitions()
	if len(definitions) != len(FileGroups) {
		t.Fatalf("file group definitions = %d, want %d", len(definitions), len(FileGroups))
	}
	requireUniqueTermKeys(t, definitions)
	requireTerm(t, definitions, string(FileGroupBackendSupportModelHelper), "Backend support/model/helper files")
}

func TestClassifyFileGroupCoversBackendSupportModelHelperSample(t *testing.T) {
	tests := []struct {
		path           string
		kind           string
		appLayer       string
		functionalArea string
		role           FileRole
	}{
		{"internal/frameworks/frameworks.go", "source", "backend", "analyzer", FileRoleAnalyzerHelper},
		{"internal/scopeir/sort_keys.go", "source", "backend", "providers", FileRoleHelper},
		{"internal/group/types.go", "source", "backend", "query", FileRoleContractModel},
		{"internal/repo/paths.go", "source", "backend", "storage", FileRoleStorageHelper},
		{"internal/testutil/path.go", "source", "backend", "unknown", FileRoleTestHelper},
		{"internal/repo/settings.go", "source", "backend", "storage", FileRoleConfig},
		{"internal/repo/runtime_config.go", "source", "backend", "storage", FileRoleConfig},
		{"internal/cobol/copy_expander.go", "source", "backend", "analyzer", FileRoleAnalyzerHelper},
		{"internal/parser/metrics.go", "source", "backend", "providers", FileRoleParserModel},
		{"internal/session/error.go", "source", "backend", "session", FileRoleRuntimeModel},
		{"internal/resolution/source_site.go", "source", "backend", "resolution", FileRoleHelper},
		{"internal/scopeir/facts.go", "source", "backend", "providers", FileRoleParserModel},
		{"internal/scopeir/range.go", "source", "backend", "providers", FileRoleParserModel},
		{"internal/session/types.go", "source", "backend", "session", FileRoleRuntimeModel},
		{"internal/cli/exit_error.go", "source", "backend", "cli", FileRoleHelper},
		{"internal/lbugnative/runner.go", "source", "backend", "storage", FileRoleAdapter},
		{"internal/lbugnative/runner_default.go", "source", "backend", "storage", FileRoleFallbackAdapter},
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			role := ClassifyFileRole(tt.path, tt.kind, tt.appLayer, tt.functionalArea)
			if role.Role != tt.role {
				t.Fatalf("ClassifyFileRole(%q) = %q, want %q", tt.path, role.Role, tt.role)
			}
			group := ClassifyFileGroup(tt.path, tt.kind, tt.appLayer, string(role.Role))
			if group.Group != FileGroupBackendSupportModelHelper {
				t.Fatalf("ClassifyFileGroup(%q) = %q source=%q, want %q", tt.path, group.Group, group.Source, FileGroupBackendSupportModelHelper)
			}
			if group.Source == "" {
				t.Fatalf("ClassifyFileGroup(%q) source is empty", tt.path)
			}
		})
	}
}

func TestClassifyFileGroupBoundaries(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		kind     string
		appLayer string
		role     FileRole
	}{
		{name: "unknown role", path: "internal/analyze/analyze.go", kind: "source", appLayer: "backend", role: FileRoleUnknown},
		{name: "frontend source", path: "anvien-web/src/lib/theme.ts", kind: "source", appLayer: "frontend", role: FileRoleHelper},
		{name: "backend test kind", path: "internal/mcp/server_test.go", kind: "test", appLayer: "backend_test", role: FileRoleTestHelper},
		{name: "docs kind", path: "docs/plans/example.md", kind: "docs", appLayer: "docs", role: FileRoleHelper},
		{name: "config kind", path: ".anvien/config.json", kind: "config", appLayer: "config", role: FileRoleConfig},
		{name: "generated kind", path: "anvien-web/src/generated/anvien-contracts.ts", kind: "generated", appLayer: "generated_contract", role: FileRoleContractModel},
		{name: "missing path", path: "", kind: "source", appLayer: "backend", role: FileRoleHelper},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			group := ClassifyFileGroup(tt.path, tt.kind, tt.appLayer, string(tt.role))
			if group.Group != "" {
				t.Fatalf("ClassifyFileGroup(%q) = %q source=%q, want empty group", tt.path, group.Group, group.Source)
			}
			if group.Source == "" {
				t.Fatalf("ClassifyFileGroup(%q) source is empty", tt.path)
			}
		})
	}
}
