package semantic

import "testing"

func TestClassifyFileRoleCoversRawOnlySupportFiles(t *testing.T) {
	tests := []struct {
		path           string
		kind           string
		appLayer       string
		functionalArea string
		want           FileRole
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
			got := ClassifyFileRole(tt.path, tt.kind, tt.appLayer, tt.functionalArea)
			if got.Role != tt.want {
				t.Fatalf("ClassifyFileRole(%q) = %q source=%q, want %q", tt.path, got.Role, got.Source, tt.want)
			}
			if got.Source == "" {
				t.Fatalf("ClassifyFileRole(%q) source is empty", tt.path)
			}
		})
	}
}

func TestClassifyFileRoleFallbacks(t *testing.T) {
	tests := []struct {
		name           string
		path           string
		kind           string
		appLayer       string
		functionalArea string
		want           FileRole
	}{
		{name: "missing path", path: "", want: FileRoleUnknown},
		{name: "analyzer source is not helper by area alone", path: "internal/analyze/analyze.go", kind: "source", appLayer: "backend", functionalArea: "analyzer", want: FileRoleUnknown},
		{name: "test kind", path: "src/app_test.go", kind: "test", appLayer: "backend_test", functionalArea: "mcp", want: FileRoleTestHelper},
		{name: "unmatched frontend source", path: "anvien-web/src/lib/theme.ts", kind: "source", appLayer: "frontend", functionalArea: "unknown", want: FileRoleUnknown},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ClassifyFileRole(tt.path, tt.kind, tt.appLayer, tt.functionalArea)
			if got.Role != tt.want {
				t.Fatalf("ClassifyFileRole(%q) = %q, want %q", tt.path, got.Role, tt.want)
			}
		})
	}
}

func TestFileRoleDefinitionsAreStable(t *testing.T) {
	definitions := FileRoleDefinitions()
	if len(definitions) != len(FileRoles) {
		t.Fatalf("file role definitions = %d, want %d", len(definitions), len(FileRoles))
	}
	requireUniqueTermKeys(t, definitions)
	requireTerm(t, definitions, string(FileRoleAnalyzerHelper), "Analyzer Helper")
	requireTerm(t, definitions, string(FileRoleFallbackAdapter), "Fallback Adapter")
	requireTerm(t, definitions, string(FileRoleUnknown), "Unknown")
}
