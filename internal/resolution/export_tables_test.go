package resolution

import (
	"reflect"
	"testing"

	"github.com/tamnguyendinh/anvien/internal/scanner"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
)

func TestBuildExportTablesZeroPhysicalBarrelAndDeterministicEntries(t *testing.T) {
	barrelFile := "src/barrel.ts"
	reexportRaw := "./impl"
	namespaceRaw := "./impl"
	starRaw := "./other"
	barrelScope := "scope:src/barrel.ts#1:0-20:0:Module"

	reexportFact := scopeir.ExportFact{
		FilePath:           barrelFile,
		FileHash:           "hash-barrel",
		Kind:               scopeir.ExportReexport,
		ExportedName:       "run",
		TargetRaw:          &reexportRaw,
		TargetExportedName: "run",
		Meanings:           []scopeir.ExportMeaning{scopeir.ExportMeaningValue},
		Range:              scopeir.Range{StartLine: 1, StartCol: 9, EndLine: 1, EndCol: 12},
		SelectionRange:     &scopeir.Range{StartLine: 1, StartCol: 9, EndLine: 1, EndCol: 12},
		Provenance: scopeir.ExportProvenance{
			StatementRange: scopeir.Range{StartLine: 1, StartCol: 0, EndLine: 1, EndCol: 31},
			SiteKind:       "export_specifier",
		},
	}
	starFact := scopeir.ExportFact{
		FilePath:  barrelFile,
		FileHash:  "hash-barrel",
		Kind:      scopeir.ExportStar,
		TargetRaw: &starRaw,
		Meanings: []scopeir.ExportMeaning{
			scopeir.ExportMeaningValue,
		},
		Range: scopeir.Range{StartLine: 2, StartCol: 7, EndLine: 2, EndCol: 8},
		Provenance: scopeir.ExportProvenance{
			StatementRange: scopeir.Range{StartLine: 2, StartCol: 0, EndLine: 2, EndCol: 29},
			SiteKind:       "export_star",
		},
	}
	namespaceFact := scopeir.ExportFact{
		FilePath:     barrelFile,
		FileHash:     "hash-barrel",
		Kind:         scopeir.ExportNamespace,
		ExportedName: "ns",
		TargetRaw:    &namespaceRaw,
		Meanings:     []scopeir.ExportMeaning{scopeir.ExportMeaningNamespace},
		Range:        scopeir.Range{StartLine: 3, StartCol: 13, EndLine: 3, EndCol: 15},
		Provenance: scopeir.ExportProvenance{
			StatementRange: scopeir.Range{StartLine: 3, StartCol: 0, EndLine: 3, EndCol: 35},
			SiteKind:       "export_namespace",
		},
	}

	newImport := func(kind scopeir.ImportKind, localName, importedName, raw string, targetFiles ...string) resolvedImport {
		rawCopy := raw
		return resolvedImport{
			Fact: scopeir.ImportFact{
				FilePath:     barrelFile,
				Kind:         kind,
				LocalName:    localName,
				ImportedName: importedName,
				TargetRaw:    &rawCopy,
			},
			TargetFiles: targetFiles,
		}
	}
	imports := []resolvedImport{
		newImport(scopeir.ImportWildcard, "", "", starRaw, "src/other.ts"),
		newImport(scopeir.ImportWildcard, "", "", reexportRaw, "src/impl.ts"),
		newImport(scopeir.ImportReexport, "run", "run", reexportRaw, "src/impl.ts"),
	}
	input := scopeir.ScopeIR{
		FilePath:    barrelFile,
		FileHash:    "hash-barrel",
		Language:    scanner.TypeScript,
		ModuleScope: barrelScope,
		Exports:     []scopeir.ExportFact{namespaceFact, starFact, reexportFact},
	}

	tables := buildExportTables([]scopeir.ScopeIR{input}, imports)
	table, ok := tables[barrelFile]
	if !ok {
		t.Fatalf("zero-physical barrel table missing: %#v", tables)
	}
	if len(table.Explicit) != 2 {
		t.Fatalf("explicit export-name keys = %d, want 2: %#v", len(table.Explicit), table.Explicit)
	}
	if len(table.Explicit["run"]) != 1 || len(table.Explicit["ns"]) != 1 {
		t.Fatalf("explicit entries = %#v, want run and ns", table.Explicit)
	}
	if len(table.StarAdjacency) != 1 {
		t.Fatalf("star adjacency = %#v, want one edge", table.StarAdjacency)
	}

	runEntry := table.Explicit["run"][0]
	if runEntry.Fact.Kind != scopeir.ExportReexport || runEntry.Fact.TargetExportedName != "run" ||
		!reflect.DeepEqual(runEntry.Fact.Meanings, []scopeir.ExportMeaning{scopeir.ExportMeaningValue}) {
		t.Fatalf("run entry fact drifted: %#v", runEntry.Fact)
	}
	if !reflect.DeepEqual(runEntry.TargetFiles, []string{"src/impl.ts"}) {
		t.Fatalf("run target files = %#v, want src/impl.ts", runEntry.TargetFiles)
	}
	if table.Explicit["ns"][0].Fact.Kind != scopeir.ExportNamespace {
		t.Fatalf("namespace entry kind = %q, want namespace", table.Explicit["ns"][0].Fact.Kind)
	}
	if !reflect.DeepEqual(table.StarAdjacency[0].TargetFiles, []string{"src/other.ts"}) {
		t.Fatalf("star target files = %#v, want src/other.ts", table.StarAdjacency[0].TargetFiles)
	}
	if _, synthesizedDefault := table.Explicit["default"]; synthesizedDefault {
		t.Fatalf("star adjacency synthesized an implicit default export: %#v", table.Explicit)
	}

	// Table facts and target paths own their nested values; mutating the source
	// input after construction must not mutate the accepted table snapshot.
	input.Exports[0].Meanings[0] = scopeir.ExportMeaningValue
	*input.Exports[0].TargetRaw = "./mutated"
	if table.Explicit["ns"][0].Fact.Meanings[0] != scopeir.ExportMeaningNamespace ||
		*table.Explicit["ns"][0].Fact.TargetRaw != "./impl" {
		t.Fatalf("table fact aliases source input: %#v", table.Explicit["ns"][0].Fact)
	}

	permuted := input
	permutedNamespaceRaw := "./impl"
	permutedNamespaceFact := namespaceFact
	permutedNamespaceFact.TargetRaw = &permutedNamespaceRaw
	permutedNamespaceFact.Meanings = []scopeir.ExportMeaning{scopeir.ExportMeaningNamespace}
	permuted.Exports = []scopeir.ExportFact{reexportFact, permutedNamespaceFact, starFact}
	permutedTables := buildExportTables([]scopeir.ScopeIR{permuted}, []resolvedImport{imports[2], imports[1], imports[0]})
	if !reflect.DeepEqual(tables, permutedTables) {
		t.Fatalf("export tables are not deterministic across input order:\nfirst=%#v\nsecond=%#v", tables, permutedTables)
	}
}

func TestBuildWorkspaceWiresExportTablesWithoutPhysicalDefinitions(t *testing.T) {
	barrelFile := "src/barrel.ts"
	targetFile := "src/impl.ts"
	targetRaw := "./impl"
	barrelScope := "scope:src/barrel.ts#1:0-3:0:Module"
	targetScope := "scope:src/impl.ts#1:0-3:0:Module"

	barrel := scopeir.ScopeIR{
		FilePath:    barrelFile,
		Language:    scanner.TypeScript,
		ModuleScope: barrelScope,
		Scopes: []scopeir.ScopeFact{{
			ID:       barrelScope,
			Kind:     scopeir.ScopeModule,
			FilePath: barrelFile,
		}},
		Exports: []scopeir.ExportFact{{
			FilePath:           barrelFile,
			Kind:               scopeir.ExportReexport,
			ExportedName:       "run",
			TargetRaw:          &targetRaw,
			TargetExportedName: "run",
			Meanings:           []scopeir.ExportMeaning{scopeir.ExportMeaningValue},
		}},
		Imports: []scopeir.ImportFact{{
			FilePath:     barrelFile,
			Kind:         scopeir.ImportReexport,
			LocalName:    "run",
			ImportedName: "run",
			TargetRaw:    &targetRaw,
		}},
	}
	target := scopeir.ScopeIR{
		FilePath:    targetFile,
		Language:    scanner.TypeScript,
		ModuleScope: targetScope,
		Scopes: []scopeir.ScopeFact{{
			ID:       targetScope,
			Kind:     scopeir.ScopeModule,
			FilePath: targetFile,
		}},
	}

	workspace, err := buildWorkspace([]scopeir.ScopeIR{barrel, target})
	if err != nil {
		t.Fatalf("buildWorkspace() error = %v", err)
	}
	table, ok := workspace.exportTables[barrelFile]
	if !ok {
		t.Fatalf("buildWorkspace() did not wire barrel export table: %#v", workspace.exportTables)
	}
	entries := table.Explicit["run"]
	if len(entries) != 1 {
		t.Fatalf("wired run entries = %#v, want one", entries)
	}
	if !reflect.DeepEqual(entries[0].TargetFiles, []string{targetFile}) {
		t.Fatalf("wired target files = %#v, want %s", entries[0].TargetFiles, targetFile)
	}
	if len(target.Definitions) != 0 || len(barrel.Definitions) != 0 {
		t.Fatalf("test fixture unexpectedly contains physical definitions")
	}
}

func TestBuildExportTablesRetainsExplicitFactKindsAndMeanings(t *testing.T) {
	filePath := "src/forms.ts"
	facts := []scopeir.ExportFact{
		{
			FilePath: filePath, Kind: scopeir.ExportDirect, ExportedName: "value", LocalName: "value",
			LocalDefID: "def:value", Meanings: []scopeir.ExportMeaning{scopeir.ExportMeaningValue},
		},
		{
			FilePath: filePath, Kind: scopeir.ExportDefault, ExportedName: "default", LocalName: "makeDefault",
			LocalDefID: "def:default", Meanings: []scopeir.ExportMeaning{scopeir.ExportMeaningValue},
		},
		{
			FilePath: filePath, Kind: scopeir.ExportAlias, ExportedName: "renamed", LocalName: "value",
			LocalDefID: "def:value", Meanings: []scopeir.ExportMeaning{scopeir.ExportMeaningValue},
		},
		{
			FilePath: filePath, Kind: scopeir.ExportNamed, ExportedName: "OnlyType", LocalName: "OnlyType",
			LocalDefID: "def:type", Meanings: []scopeir.ExportMeaning{scopeir.ExportMeaningType}, TypeOnly: true,
		},
	}
	table := buildExportTables([]scopeir.ScopeIR{{FilePath: filePath, Exports: facts}}, nil)[filePath]
	checks := []struct {
		name     string
		kind     scopeir.ExportKind
		local    string
		defID    string
		meanings []scopeir.ExportMeaning
		typeOnly bool
	}{
		{"value", scopeir.ExportDirect, "value", "def:value", []scopeir.ExportMeaning{scopeir.ExportMeaningValue}, false},
		{"default", scopeir.ExportDefault, "makeDefault", "def:default", []scopeir.ExportMeaning{scopeir.ExportMeaningValue}, false},
		{"renamed", scopeir.ExportAlias, "value", "def:value", []scopeir.ExportMeaning{scopeir.ExportMeaningValue}, false},
		{"OnlyType", scopeir.ExportNamed, "OnlyType", "def:type", []scopeir.ExportMeaning{scopeir.ExportMeaningType}, true},
	}
	for _, check := range checks {
		entries := table.Explicit[check.name]
		if len(entries) != 1 {
			t.Fatalf("explicit entries for %q = %#v, want one", check.name, entries)
		}
		fact := entries[0].Fact
		if fact.Kind != check.kind || fact.LocalName != check.local || fact.LocalDefID != check.defID ||
			fact.TypeOnly != check.typeOnly || !reflect.DeepEqual(fact.Meanings, check.meanings) {
			t.Fatalf("explicit fact %q = %#v, want kind/local/def/meaning/type-only %#v/%q/%q/%#v/%t", check.name, fact, check.kind, check.local, check.defID, check.meanings, check.typeOnly)
		}
	}
}
