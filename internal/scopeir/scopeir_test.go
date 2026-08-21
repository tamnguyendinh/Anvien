package scopeir

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/tamnguyendinh/anvien/internal/scanner"
)

func TestMarshalDeterministicMatchesGolden(t *testing.T) {
	raw, err := sampleScopeIR().MarshalDeterministic()
	if err != nil {
		t.Fatalf("MarshalDeterministic failed: %v", err)
	}
	golden, err := os.ReadFile("testdata/scopeir.golden.json")
	if err != nil {
		t.Fatalf("read golden: %v", err)
	}
	if string(raw) != string(golden) {
		t.Fatalf("golden mismatch\nwant:\n%s\ngot:\n%s", golden, raw)
	}
}

func TestUnmarshalNormalizesScopeIR(t *testing.T) {
	raw, err := sampleScopeIR().MarshalDeterministic()
	if err != nil {
		t.Fatalf("MarshalDeterministic failed: %v", err)
	}
	decoded, err := Unmarshal(raw)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	roundTrip, err := decoded.MarshalDeterministic()
	if err != nil {
		t.Fatalf("round trip marshal failed: %v", err)
	}
	if string(roundTrip) != string(raw) {
		t.Fatalf("round trip changed JSON\nfirst:\n%s\nsecond:\n%s", raw, roundTrip)
	}
	if decoded.Version != Version {
		t.Fatalf("Version = %q, want %q", decoded.Version, Version)
	}
}

func TestNormalizedDoesNotMutateSource(t *testing.T) {
	ir := largeUnorderedScopeIR(3)
	firstScopeID := ir.Scopes[0].ID
	firstBinding := ir.Scopes[0].Bindings[0].Name

	normalized := ir.Normalized()

	if ir.Scopes[0].ID != firstScopeID {
		t.Fatalf("Normalized mutated scope order: got %q, want %q", ir.Scopes[0].ID, firstScopeID)
	}
	if ir.Scopes[0].Bindings[0].Name != firstBinding {
		t.Fatalf("Normalized mutated nested bindings: got %q, want %q", ir.Scopes[0].Bindings[0].Name, firstBinding)
	}
	if reflect.DeepEqual(ir, normalized) {
		t.Fatalf("test fixture should differ after normalization")
	}
}

func TestNormalizeInPlaceMatchesNormalized(t *testing.T) {
	ir := largeUnorderedScopeIR(12)

	normalized := ir.Normalized()
	inPlace := ir.NormalizeInPlace()

	if !reflect.DeepEqual(inPlace, normalized) {
		t.Fatalf("NormalizeInPlace() differed from Normalized()")
	}
}

func TestNormalizeOwnedMatchesNormalized(t *testing.T) {
	ir := largeUnorderedScopeIR(12)

	normalized := ir.Normalized()
	owned := ir.NormalizeOwned()

	if !reflect.DeepEqual(owned, normalized) {
		t.Fatalf("NormalizeOwned() differed from Normalized()")
	}
}

func TestImportRequestedMeaningsCanonicalizeCloneAndOrder(t *testing.T) {
	targetRaw := "./dep"
	source := ScopeIR{Imports: []ImportFact{{
		ID:                "canonical",
		FilePath:          "src/app.ts",
		Kind:              ImportNamed,
		LocalName:         "value",
		ImportedName:      "value",
		RequestedMeanings: []ExportMeaning{ExportMeaningValue, ExportMeaningType, ExportMeaningNamespace, ExportMeaningValue},
		TargetRaw:         &targetRaw,
	}}}
	wantMeanings := []ExportMeaning{ExportMeaningNamespace, ExportMeaningType, ExportMeaningValue}
	for name, normalized := range map[string]ScopeIR{
		"Normalized":     source.Normalized(),
		"NormalizeOwned": source.NormalizeOwned(),
	} {
		if !reflect.DeepEqual(normalized.Imports[0].RequestedMeanings, wantMeanings) {
			t.Fatalf("%s meanings = %#v, want canonical set %#v", name, normalized.Imports[0].RequestedMeanings, wantMeanings)
		}
		normalized.Imports[0].RequestedMeanings[0] = ExportMeaning("mutated")
		if source.Imports[0].RequestedMeanings[0] != ExportMeaningValue ||
			source.Imports[0].RequestedMeanings[1] != ExportMeaningType {
			t.Fatalf("%s retained requested-meaning alias: source=%#v normalized=%#v", name, source.Imports[0], normalized.Imports[0])
		}
	}

	importFact := func(id string, meanings []ExportMeaning, typeOnly bool) ImportFact {
		return ImportFact{
			ID:                id,
			FilePath:          "src/order.ts",
			Kind:              ImportNamed,
			LocalName:         "same",
			ImportedName:      "same",
			RequestedMeanings: meanings,
			TypeOnly:          typeOnly,
			TargetRaw:         &targetRaw,
		}
	}
	imports := []ImportFact{
		importFact("a-value", []ExportMeaning{ExportMeaningValue}, false),
		importFact("a-type-only", []ExportMeaning{ExportMeaningType}, true),
		importFact("z-type", []ExportMeaning{ExportMeaningType}, false),
		importFact("z-namespace", []ExportMeaning{ExportMeaningNamespace}, false),
	}
	left := ScopeIR{Imports: append([]ImportFact(nil), imports...)}
	right := ScopeIR{Imports: []ImportFact{imports[2], imports[0], imports[3], imports[1]}}
	leftRaw, err := left.MarshalDeterministic()
	if err != nil {
		t.Fatalf("marshal left imports: %v", err)
	}
	rightRaw, err := right.MarshalDeterministic()
	if err != nil {
		t.Fatalf("marshal right imports: %v", err)
	}
	if string(leftRaw) != string(rightRaw) {
		t.Fatalf("deterministic import marshal mismatch\nleft:\n%s\nright:\n%s", leftRaw, rightRaw)
	}

	normalized := left.Normalized()
	wantOrder := []string{"z-namespace", "z-type", "a-type-only", "a-value"}
	for index, wantID := range wantOrder {
		if normalized.Imports[index].ID != wantID {
			t.Fatalf("normalized import order[%d] = %q, want %q: %#v", index, normalized.Imports[index].ID, wantID, normalized.Imports)
		}
	}
	roundTrip, err := Unmarshal(leftRaw)
	if err != nil {
		t.Fatalf("unmarshal requested meanings: %v", err)
	}
	if !reflect.DeepEqual(roundTrip.Imports, normalized.Imports) {
		t.Fatalf("requested-meaning round trip changed imports: got=%#v want=%#v", roundTrip.Imports, normalized.Imports)
	}
}

func TestBindingCollectionsAreDeepCopiedByNormalization(t *testing.T) {
	index := 2
	selection := Range{StartLine: 3, StartCol: 10, EndLine: 3, EndCol: 14}
	ir := ScopeIR{
		BindingLeaves: []BindingLeafFact{{
			FilePath:       "src/pattern.ts",
			Name:           "leaf",
			Range:          Range{StartLine: 3, StartCol: 2, EndLine: 3, EndCol: 14},
			SelectionRange: &selection,
			Path: []BindingPathSegment{{
				Kind:        BindingPathArrayIndex,
				ArrayIndex:  &index,
				SourceRange: Range{StartLine: 3, StartCol: 2, EndLine: 3, EndCol: 14},
			}},
		}},
		ExtractionDiagnostics: []ExtractionDiagnosticFact{{
			Code:     DiagnosticUnsupportedBindingNode,
			FilePath: "src/pattern.ts",
			Range:    Range{StartLine: 4, StartCol: 2, EndLine: 4, EndCol: 9},
			NodeKind: "member_expression",
			Reason:   "unsupported",
			Path: []BindingPathSegment{{
				Kind:        BindingPathArrayIndex,
				ArrayIndex:  &index,
				SourceRange: Range{StartLine: 4, StartCol: 2, EndLine: 4, EndCol: 9},
			}},
		}},
	}

	for name, normalized := range map[string]ScopeIR{
		"Normalized":     ir.Normalized(),
		"NormalizeOwned": ir.NormalizeOwned(),
	} {
		normalized.BindingLeaves[0].SelectionRange.StartCol = 99
		*normalized.BindingLeaves[0].Path[0].ArrayIndex = 99
		*normalized.ExtractionDiagnostics[0].Path[0].ArrayIndex = 100
		if ir.BindingLeaves[0].SelectionRange.StartCol != 10 ||
			*ir.BindingLeaves[0].Path[0].ArrayIndex != 2 ||
			*ir.ExtractionDiagnostics[0].Path[0].ArrayIndex != 2 {
			t.Fatalf("%s retained nested aliases: source=%#v normalized=%#v", name, ir, normalized)
		}
	}
}

func TestBindingCollectionsMarshalDeterministically(t *testing.T) {
	index0 := 0
	index2 := 2
	selectionA := Range{StartLine: 2, StartCol: 1, EndLine: 2, EndCol: 2}
	selectionB := Range{StartLine: 2, StartCol: 5, EndLine: 2, EndCol: 6}
	leaves := []BindingLeafFact{
		{
			FilePath:       "src/pattern.ts",
			FileHash:       "hash-b",
			Name:           "b",
			Range:          Range{StartLine: 2, StartCol: 5, EndLine: 2, EndCol: 10},
			SelectionRange: &selectionB,
			Path:           []BindingPathSegment{{Kind: BindingPathArrayIndex, ArrayIndex: &index2}},
			Default:        true,
			Provenance: BindingPatternProvenance{
				Context:        BindingContextVariable,
				ConstructRange: Range{StartLine: 2, EndLine: 2, EndCol: 20},
				PatternRange:   Range{StartLine: 2, StartCol: 1, EndLine: 2, EndCol: 11},
				PatternKind:    "array_pattern",
			},
		},
		{
			FilePath:       "src/pattern.ts",
			FileHash:       "hash-a",
			Name:           "a",
			Range:          selectionA,
			SelectionRange: &selectionA,
			Path:           []BindingPathSegment{{Kind: BindingPathArrayIndex, ArrayIndex: &index0}},
			Provenance: BindingPatternProvenance{
				Context:        BindingContextVariable,
				ConstructRange: Range{StartLine: 2, EndLine: 2, EndCol: 20},
				PatternRange:   Range{StartLine: 2, StartCol: 1, EndLine: 2, EndCol: 11},
				PatternKind:    "array_pattern",
			},
		},
	}
	diagnostics := []ExtractionDiagnosticFact{
		{Code: DiagnosticUnsupportedBindingNode, FilePath: "src/pattern.ts", Range: Range{StartLine: 4}, NodeKind: "z", Reason: "z"},
		{Code: DiagnosticMalformedBindingNode, FilePath: "src/pattern.ts", Range: Range{StartLine: 3}, NodeKind: "a", Reason: "a"},
	}
	left := ScopeIR{
		FilePath:              "src/pattern.ts",
		BindingLeaves:         cloneBindingLeavesForTest(leaves),
		ExtractionDiagnostics: cloneExtractionDiagnosticsForTest(diagnostics),
	}
	right := ScopeIR{
		FilePath:              "src/pattern.ts",
		BindingLeaves:         cloneBindingLeavesForTest([]BindingLeafFact{leaves[1], leaves[0]}),
		ExtractionDiagnostics: cloneExtractionDiagnosticsForTest([]ExtractionDiagnosticFact{diagnostics[1], diagnostics[0]}),
	}
	leftBefore := cloneBindingLeavesForTest(left.BindingLeaves)
	leftDiagnosticsBefore := cloneExtractionDiagnosticsForTest(left.ExtractionDiagnostics)

	leftRaw, err := left.MarshalDeterministic()
	if err != nil {
		t.Fatalf("marshal left: %v", err)
	}
	rightRaw, err := right.MarshalDeterministic()
	if err != nil {
		t.Fatalf("marshal right: %v", err)
	}
	if string(leftRaw) != string(rightRaw) {
		t.Fatalf("deterministic binding marshal mismatch\nleft:\n%s\nright:\n%s", leftRaw, rightRaw)
	}
	if left.BindingLeaves[0].Name != "b" || right.BindingLeaves[0].Name != "a" ||
		left.ExtractionDiagnostics[0].Code != DiagnosticUnsupportedBindingNode ||
		right.ExtractionDiagnostics[0].Code != DiagnosticMalformedBindingNode {
		t.Fatalf("MarshalDeterministic reordered source slices: left=%#v right=%#v", left, right)
	}
	if !reflect.DeepEqual(left.BindingLeaves, leftBefore) ||
		!reflect.DeepEqual(left.ExtractionDiagnostics, leftDiagnosticsBefore) {
		t.Fatalf("MarshalDeterministic mutated source slices: left=%#v", left)
	}
	left.BindingLeaves[0].SelectionRange.StartCol = 99
	*left.BindingLeaves[0].Path[0].ArrayIndex = 99
	if leaves[0].SelectionRange.StartCol != 5 || *leaves[0].Path[0].ArrayIndex != 2 ||
		right.BindingLeaves[0].SelectionRange.StartCol != 1 || *right.BindingLeaves[0].Path[0].ArrayIndex != 0 {
		t.Fatalf("comparison inputs retained aliases: source=%#v right=%#v", leaves, right)
	}
}

func TestExportCollectionsAreDeepCopiedByNormalization(t *testing.T) {
	targetRaw := "./source"
	selection := Range{StartLine: 3, StartCol: 7, EndLine: 3, EndCol: 12}
	ir := ScopeIR{
		Definitions: []DefinitionFact{{
			ID:         "def:src/export.ts#3:0:Class:Thing",
			FilePath:   "src/export.ts",
			Name:       "Thing",
			Label:      NodeClass,
			Range:      Range{StartLine: 3, EndLine: 3, EndCol: 17},
			Visibility: "private",
		}},
		Exports: []ExportFact{{
			FilePath:       "src/export.ts",
			FileHash:       "hash-export",
			Kind:           ExportDirect,
			ExportedName:   "Thing",
			LocalName:      "Thing",
			LocalDefID:     "def:src/export.ts#3:0:Class:Thing",
			Meanings:       []ExportMeaning{ExportMeaningValue, ExportMeaningType, ExportMeaningValue},
			Range:          Range{StartLine: 3, EndLine: 3, EndCol: 17},
			SelectionRange: &selection,
			Provenance: ExportProvenance{
				StatementRange: Range{StartLine: 3, EndLine: 3, EndCol: 17},
				SiteKind:       "export_declaration",
			},
		}, {
			FilePath:           "src/export.ts",
			FileHash:           "hash-export",
			Kind:               ExportReexport,
			ExportedName:       "SourceThing",
			TargetRaw:          &targetRaw,
			TargetExportedName: "Thing",
			Meanings:           []ExportMeaning{ExportMeaningValue},
			Range:              Range{StartLine: 4, EndLine: 4, EndCol: 31},
			Provenance: ExportProvenance{
				StatementRange: Range{StartLine: 4, EndLine: 4, EndCol: 31},
				SiteKind:       "export_specifier",
			},
		}},
		ExportDiagnostics: []ExportDiagnosticFact{{
			Code:     ExportDiagnosticUnsupportedSyntax,
			FilePath: "src/export.ts",
			Range:    Range{StartLine: 8, StartCol: 0, EndLine: 8, EndCol: 12},
			NodeKind: "export_statement",
			Reason:   "unsupported export form",
			Provenance: ExportProvenance{
				StatementRange: Range{StartLine: 8, EndLine: 8, EndCol: 12},
				SiteKind:       "export_statement",
			},
		}},
	}

	for name, normalized := range map[string]ScopeIR{
		"Normalized":     ir.Normalized(),
		"NormalizeOwned": ir.NormalizeOwned(),
	} {
		wantMeanings := []ExportMeaning{ExportMeaningType, ExportMeaningValue}
		if !reflect.DeepEqual(normalized.Exports[0].Meanings, wantMeanings) {
			t.Fatalf("%s meanings = %#v, want canonical set %#v", name, normalized.Exports[0].Meanings, wantMeanings)
		}
		normalized.Exports[0].Meanings[0] = ExportMeaningNamespace
		normalized.Exports[0].SelectionRange.StartCol = 99
		*normalized.Exports[1].TargetRaw = "./mutated"
		if ir.Exports[0].Meanings[0] != ExportMeaningValue ||
			ir.Exports[0].SelectionRange.StartCol != 7 ||
			*ir.Exports[1].TargetRaw != "./source" {
			t.Fatalf("%s retained nested export aliases: source=%#v normalized=%#v", name, ir.Exports[0], normalized.Exports[0])
		}
		if normalized.Definitions[0].Visibility != "private" {
			t.Fatalf("%s changed access visibility while normalizing exports: %q", name, normalized.Definitions[0].Visibility)
		}
		if len(normalized.ExportDiagnostics) != 1 || normalized.ExportDiagnostics[0].Code != ExportDiagnosticUnsupportedSyntax {
			t.Fatalf("%s lost structured export diagnostic: %#v", name, normalized.ExportDiagnostics)
		}
	}
}

func TestExportCollectionsMarshalDeterministically(t *testing.T) {
	emptyTarget := ""
	exportWithTarget := func(target *string) ExportFact {
		return ExportFact{
			FilePath:           "src/export.ts",
			Kind:               ExportReexport,
			ExportedName:       "value",
			TargetRaw:          target,
			TargetExportedName: "value",
			Meanings:           []ExportMeaning{ExportMeaningValue, ExportMeaningValue},
			Range:              Range{StartLine: 2, EndLine: 2, EndCol: 25},
			Provenance: ExportProvenance{
				StatementRange: Range{StartLine: 2, EndLine: 2, EndCol: 25},
				SiteKind:       "export_specifier",
			},
		}
	}
	diagnostics := []ExportDiagnosticFact{
		{
			Code:       ExportDiagnosticMalformedSyntax,
			FilePath:   "src/export.ts",
			Range:      Range{StartLine: 6},
			NodeKind:   "export_statement",
			Reason:     "malformed",
			Provenance: ExportProvenance{StatementRange: Range{StartLine: 6}, SiteKind: "export_statement"},
		},
		{
			Code:       ExportDiagnosticUnsupportedSyntax,
			FilePath:   "src/export.ts",
			Range:      Range{StartLine: 5},
			NodeKind:   "export_statement",
			Reason:     "unsupported",
			Provenance: ExportProvenance{StatementRange: Range{StartLine: 5}, SiteKind: "export_statement"},
		},
	}
	left := ScopeIR{
		FilePath: "src/export.ts",
		Exports: []ExportFact{
			exportWithTarget(&emptyTarget),
			exportWithTarget(nil),
			{FilePath: "src/export.ts", Kind: ExportNamespace, ExportedName: "ns", Meanings: []ExportMeaning{ExportMeaningNamespace}, Range: Range{StartLine: 1}, Provenance: ExportProvenance{StatementRange: Range{StartLine: 1}, SiteKind: "export_namespace"}},
		},
		ExportDiagnostics: []ExportDiagnosticFact{diagnostics[0], diagnostics[1]},
	}
	right := ScopeIR{
		FilePath:          "src/export.ts",
		Exports:           []ExportFact{left.Exports[2], left.Exports[1], left.Exports[0]},
		ExportDiagnostics: []ExportDiagnosticFact{diagnostics[1], diagnostics[0]},
	}

	leftBefore := append([]ExportFact(nil), left.Exports...)
	leftDiagnosticsBefore := append([]ExportDiagnosticFact(nil), left.ExportDiagnostics...)
	leftRaw, err := left.MarshalDeterministic()
	if err != nil {
		t.Fatalf("marshal left: %v", err)
	}
	rightRaw, err := right.MarshalDeterministic()
	if err != nil {
		t.Fatalf("marshal right: %v", err)
	}
	if string(leftRaw) != string(rightRaw) {
		t.Fatalf("deterministic export marshal mismatch\nleft:\n%s\nright:\n%s", leftRaw, rightRaw)
	}
	if !reflect.DeepEqual(left.Exports, leftBefore) || !reflect.DeepEqual(left.ExportDiagnostics, leftDiagnosticsBefore) {
		t.Fatalf("MarshalDeterministic mutated source export slices: %#v %#v", left.Exports, left.ExportDiagnostics)
	}
	if !strings.Contains(string(leftRaw), `"targetRaw": ""`) {
		t.Fatalf("nil-versus-empty target provenance was not preserved: %s", leftRaw)
	}
	for _, forbidden := range []string{`"targetDefId"`, `"targetFile"`, `"targetModuleScope"`, `"linkStatus"`, `"transitiveVia"`, `"reachableThroughBarrel"`, `"publicApi"`} {
		if strings.Contains(string(leftRaw), forbidden) {
			t.Fatalf("export JSON introduced terminal/resolution state %q: %s", forbidden, leftRaw)
		}
	}
	decoded, err := Unmarshal(leftRaw)
	if err != nil {
		t.Fatalf("unmarshal exports: %v", err)
	}
	if len(decoded.Exports) != 3 || len(decoded.ExportDiagnostics) != 2 {
		t.Fatalf("round trip counts = exports %d diagnostics %d, want 3/2", len(decoded.Exports), len(decoded.ExportDiagnostics))
	}
	if !reflect.DeepEqual(decoded.Exports[0].Meanings, []ExportMeaning{ExportMeaningNamespace}) {
		t.Fatalf("round trip changed namespace meaning: %#v", decoded.Exports[0].Meanings)
	}
	if decoded.ExportDiagnostics[0].Code != ExportDiagnosticUnsupportedSyntax {
		t.Fatalf("diagnostics were not sorted by source range: %#v", decoded.ExportDiagnostics)
	}
}

func TestExportKindsAndMeaningsRepresentSourceForms(t *testing.T) {
	kinds := []ExportKind{ExportDirect, ExportNamed, ExportDefault, ExportAlias, ExportReexport, ExportStar, ExportNamespace}
	meanings := []ExportMeaning{ExportMeaningValue, ExportMeaningType, ExportMeaningNamespace}
	exports := make([]ExportFact, 0, len(kinds))
	for index, kind := range kinds {
		fact := ExportFact{
			FilePath:     "src/forms.ts",
			Kind:         kind,
			ExportedName: fmt.Sprintf("name%d", index),
			Meanings:     []ExportMeaning{meanings[index%len(meanings)]},
			TypeOnly:     index == 4,
			Range:        Range{StartLine: index + 1},
			Provenance:   ExportProvenance{StatementRange: Range{StartLine: index + 1}, SiteKind: "export_statement"},
		}
		if index == 0 {
			fact.Meanings = []ExportMeaning{ExportMeaningValue, ExportMeaningType}
		}
		exports = append(exports, fact)
	}
	decoded, err := Unmarshal(mustMarshalScopeIRForTest(t, ScopeIR{FilePath: "src/forms.ts", Exports: exports}))
	if err != nil {
		t.Fatalf("unmarshal source-form matrix: %v", err)
	}
	if len(decoded.Exports) != len(kinds) {
		t.Fatalf("source-form fact count = %d, want %d", len(decoded.Exports), len(kinds))
	}
	seen := make(map[ExportKind]ExportFact, len(decoded.Exports))
	for _, fact := range decoded.Exports {
		seen[fact.Kind] = fact
	}
	for _, kind := range kinds {
		if _, ok := seen[kind]; !ok {
			t.Fatalf("source-form kind %q was not representable", kind)
		}
	}
	if !reflect.DeepEqual(seen[ExportDirect].Meanings, []ExportMeaning{ExportMeaningType, ExportMeaningValue}) {
		t.Fatalf("dual value/type meaning was not preserved: %#v", seen[ExportDirect].Meanings)
	}
	if !seen[ExportReexport].TypeOnly {
		t.Fatalf("type-only re-export state was not preserved")
	}
}

func mustMarshalScopeIRForTest(t *testing.T, ir ScopeIR) []byte {
	t.Helper()
	raw, err := ir.MarshalDeterministic()
	if err != nil {
		t.Fatalf("marshal ScopeIR fixture: %v", err)
	}
	return raw
}

func cloneBindingLeavesForTest(input []BindingLeafFact) []BindingLeafFact {
	output := make([]BindingLeafFact, len(input))
	copy(output, input)
	for index := range output {
		if input[index].SelectionRange != nil {
			selection := *input[index].SelectionRange
			output[index].SelectionRange = &selection
		}
		output[index].Path = append([]BindingPathSegment(nil), input[index].Path...)
		for pathIndex := range output[index].Path {
			if input[index].Path[pathIndex].ArrayIndex == nil {
				continue
			}
			arrayIndex := *input[index].Path[pathIndex].ArrayIndex
			output[index].Path[pathIndex].ArrayIndex = &arrayIndex
		}
	}
	return output
}

func cloneExtractionDiagnosticsForTest(input []ExtractionDiagnosticFact) []ExtractionDiagnosticFact {
	output := make([]ExtractionDiagnosticFact, len(input))
	copy(output, input)
	for index := range output {
		output[index].Path = append([]BindingPathSegment(nil), input[index].Path...)
		for pathIndex := range output[index].Path {
			if input[index].Path[pathIndex].ArrayIndex == nil {
				continue
			}
			arrayIndex := *input[index].Path[pathIndex].ArrayIndex
			output[index].Path[pathIndex].ArrayIndex = &arrayIndex
		}
	}
	return output
}

func BenchmarkScopeIRSerialization(b *testing.B) {
	ir := sampleScopeIR()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		raw, err := ir.MarshalDeterministic()
		if err != nil {
			b.Fatalf("marshal: %v", err)
		}
		if _, err := Unmarshal(raw); err != nil {
			b.Fatalf("unmarshal: %v", err)
		}
	}
}

func BenchmarkScopeIRNormalizedLargeSort(b *testing.B) {
	ir := largeUnorderedScopeIR(2000)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		normalized := ir.Normalized()
		if len(normalized.Calls) != len(ir.Calls) {
			b.Fatalf("Calls = %d, want %d", len(normalized.Calls), len(ir.Calls))
		}
	}
}

func sampleScopeIR() ScopeIR {
	moduleScope := "scope:src/app.ts#1:0-6:0:Module"
	functionScope := "scope:src/app.ts#3:0-5:1:Function"
	userDef := "def:src/app.ts#2:0:Class:User"
	runDef := "def:src/app.ts#3:0:Function:run"
	targetRaw := "./user"
	exportTargetRaw := "./user"
	exportSelection := Range{StartLine: 2, StartCol: 7, EndLine: 2, EndCol: 11}
	arity := 1
	parameterCount := 1
	requiredParameterCount := 1
	async := true

	return ScopeIR{
		FilePath:    "src/app.ts",
		FileHash:    "hash-app",
		Language:    scanner.TypeScript,
		ModuleScope: moduleScope,
		Scopes: []ScopeFact{
			{
				ID:       functionScope,
				Parent:   &moduleScope,
				Kind:     ScopeFunction,
				Range:    Range{StartLine: 3, StartCol: 0, EndLine: 5, EndCol: 1},
				FilePath: "src/app.ts",
				Bindings: []BindingFact{
					{Name: "run", DefID: runDef, Origin: BindingLocal},
				},
				TypeBindings: []TypeBindingFact{
					{
						Name: "user",
						Type: TypeRef{
							RawName:         "User",
							DeclaredAtScope: functionScope,
							Source:          TypeSourceParameter,
						},
					},
				},
			},
			{
				ID:       moduleScope,
				Kind:     ScopeModule,
				Range:    Range{StartLine: 1, StartCol: 0, EndLine: 6, EndCol: 0},
				FilePath: "src/app.ts",
				Bindings: []BindingFact{
					{Name: "run", DefID: runDef, Origin: BindingLocal},
					{Name: "User", DefID: userDef, Origin: BindingLocal},
				},
				OwnedDefIDs: []string{runDef, userDef},
			},
		},
		Definitions: []DefinitionFact{
			{
				ID:                     runDef,
				FilePath:               "src/app.ts",
				Name:                   "run",
				Label:                  NodeFunction,
				Range:                  Range{StartLine: 3, StartCol: 0, EndLine: 5, EndCol: 1},
				QualifiedName:          "run",
				ParameterCount:         &parameterCount,
				RequiredParameterCount: &requiredParameterCount,
				ParameterTypes:         []string{"User"},
				ReturnType:             "Promise<void>",
				Async:                  &async,
			},
			{
				ID:            userDef,
				FilePath:      "src/app.ts",
				Name:          "User",
				Label:         NodeClass,
				Range:         Range{StartLine: 2, StartCol: 0, EndLine: 2, EndCol: 18},
				QualifiedName: "User",
			},
		},
		Exports: []ExportFact{
			{
				FilePath:       "src/app.ts",
				FileHash:       "hash-app",
				Kind:           ExportDirect,
				ExportedName:   "User",
				LocalName:      "User",
				LocalDefID:     userDef,
				Meanings:       []ExportMeaning{ExportMeaningValue, ExportMeaningType, ExportMeaningValue},
				Range:          Range{StartLine: 2, StartCol: 0, EndLine: 2, EndCol: 18},
				SelectionRange: &exportSelection,
				Provenance:     ExportProvenance{StatementRange: Range{StartLine: 2, EndLine: 2, EndCol: 18}, SiteKind: "export_declaration"},
			},
			{
				FilePath:           "src/app.ts",
				FileHash:           "hash-app",
				Kind:               ExportReexport,
				ExportedName:       "User",
				TargetRaw:          &exportTargetRaw,
				TargetExportedName: "User",
				Meanings:           []ExportMeaning{ExportMeaningValue},
				Range:              Range{StartLine: 7, StartCol: 0, EndLine: 7, EndCol: 24},
				Provenance:         ExportProvenance{StatementRange: Range{StartLine: 7, EndLine: 7, EndCol: 24}, SiteKind: "export_specifier"},
			},
		},
		ExportDiagnostics: []ExportDiagnosticFact{
			{
				Code:       ExportDiagnosticUnsupportedSyntax,
				FilePath:   "src/app.ts",
				FileHash:   "hash-app",
				Range:      Range{StartLine: 9, StartCol: 0, EndLine: 9, EndCol: 14},
				NodeKind:   "export_statement",
				Reason:     "unsupported export form",
				Provenance: ExportProvenance{StatementRange: Range{StartLine: 9, EndLine: 9, EndCol: 14}, SiteKind: "export_statement"},
			},
		},
		Imports: []ImportFact{
			{
				FilePath:     "src/app.ts",
				Kind:         ImportNamed,
				LocalName:    "User",
				ImportedName: "User",
				TargetRaw:    &targetRaw,
			},
		},
		Calls: []CallSiteFact{
			{
				FilePath:         "src/app.ts",
				Name:             "save",
				Range:            Range{StartLine: 4, StartCol: 7, EndLine: 4, EndCol: 11},
				InScope:          functionScope,
				CallForm:         CallMember,
				ExplicitReceiver: "user",
				Arity:            &arity,
			},
		},
		Accesses: []AccessFact{
			{
				FilePath:         "src/app.ts",
				Name:             "id",
				Kind:             AccessRead,
				Range:            Range{StartLine: 4, StartCol: 12, EndLine: 4, EndCol: 14},
				InScope:          functionScope,
				ExplicitReceiver: "user",
			},
		},
		Heritage: []HeritageFact{
			{
				FilePath: "src/app.ts",
				Name:     "BaseUser",
				Kind:     HeritageExtends,
				Range:    Range{StartLine: 2, StartCol: 19, EndLine: 2, EndCol: 27},
				InScope:  moduleScope,
			},
		},
		TypeAnnotations: []TypeAnnotationFact{
			{
				FilePath: "src/app.ts",
				Name:     "user",
				Range:    Range{StartLine: 3, StartCol: 19, EndLine: 3, EndCol: 23},
				InScope:  functionScope,
				Type: TypeRef{
					RawName:         "User",
					DeclaredAtScope: functionScope,
					Source:          TypeSourceParameter,
				},
			},
		},
		ReturnTypes: []ReturnTypeFact{
			{
				DefID:    runDef,
				FilePath: "src/app.ts",
				Range:    Range{StartLine: 3, StartCol: 26, EndLine: 3, EndCol: 39},
				Type: TypeRef{
					RawName:         "Promise<void>",
					DeclaredAtScope: functionScope,
					Source:          TypeSourceReturn,
				},
			},
		},
		Frameworks: []FrameworkFact{
			{
				DefID:                runDef,
				FilePath:             "src/app.ts",
				Framework:            "express",
				Reason:               "decorator",
				EntryPointMultiplier: 2.5,
				Range:                Range{StartLine: 3, StartCol: 0, EndLine: 5, EndCol: 1},
			},
		},
		Domains: []DomainFact{
			{
				DefID:    userDef,
				FilePath: "src/app.ts",
				Domain:   "identity",
				Role:     "model",
				Reason:   "path",
				Range:    Range{StartLine: 2, StartCol: 0, EndLine: 2, EndCol: 18},
			},
		},
	}
}

func largeUnorderedScopeIR(count int) ScopeIR {
	ir := ScopeIR{
		FilePath:    "src/large.ts",
		FileHash:    "hash-large",
		Language:    scanner.TypeScript,
		ModuleScope: "scope:src/large.ts#1:0-9999:0:Module",
	}
	for i := count - 1; i >= 0; i-- {
		scopeID := fmt.Sprintf("scope:src/large.ts#%d:0-%d:0:Function", i+2, i+3)
		defID := fmt.Sprintf("def:src/large.ts#%d:0:Function:item%d", i+2, i)
		ir.Scopes = append(ir.Scopes, ScopeFact{
			ID:       scopeID,
			Kind:     ScopeFunction,
			Range:    Range{StartLine: i + 2, EndLine: i + 3},
			FilePath: "src/large.ts",
			Bindings: []BindingFact{
				{Name: fmt.Sprintf("local%d", i), DefID: defID, Origin: BindingLocal},
				{Name: fmt.Sprintf("arg%d", i), DefID: defID, Origin: BindingLocal},
			},
			OwnedDefIDs: []string{defID},
			TypeBindings: []TypeBindingFact{{
				Name: fmt.Sprintf("arg%d", i),
				Type: TypeRef{
					RawName:         fmt.Sprintf("Type%d", i),
					DeclaredAtScope: scopeID,
					Source:          TypeSourceParameter,
				},
			}},
		})
		ir.Definitions = append(ir.Definitions, DefinitionFact{
			ID:             defID,
			FilePath:       "src/large.ts",
			Name:           fmt.Sprintf("item%d", i),
			Label:          NodeFunction,
			Range:          Range{StartLine: i + 2, EndLine: i + 3},
			QualifiedName:  fmt.Sprintf("item%d", i),
			ParameterTypes: []string{fmt.Sprintf("Type%d", i)},
			Annotations:    []string{fmt.Sprintf("@route%d", i)},
		})
		ir.Calls = append(ir.Calls, CallSiteFact{
			FilePath: "src/large.ts",
			Name:     fmt.Sprintf("call%d", i),
			Range:    Range{StartLine: i + 2, StartCol: 2, EndLine: i + 2, EndCol: 20},
			InScope:  scopeID,
			ArgTypes: []string{fmt.Sprintf("Type%d", i)},
		})
		ir.Accesses = append(ir.Accesses, AccessFact{
			FilePath: "src/large.ts",
			Name:     fmt.Sprintf("field%d", i),
			Kind:     AccessRead,
			Range:    Range{StartLine: i + 2, StartCol: 21, EndLine: i + 2, EndCol: 30},
			InScope:  scopeID,
		})
		ir.Heritage = append(ir.Heritage, HeritageFact{
			FilePath: "src/large.ts",
			Name:     fmt.Sprintf("Base%d", i),
			Kind:     HeritageExtends,
			Range:    Range{StartLine: i + 2, StartCol: 31, EndLine: i + 2, EndCol: 40},
			InScope:  scopeID,
		})
		ir.TypeAnnotations = append(ir.TypeAnnotations, TypeAnnotationFact{
			FilePath: "src/large.ts",
			Name:     fmt.Sprintf("arg%d", i),
			Range:    Range{StartLine: i + 2, StartCol: 41, EndLine: i + 2, EndCol: 50},
			InScope:  scopeID,
			Type: TypeRef{
				RawName:         fmt.Sprintf("Type%d", i),
				DeclaredAtScope: scopeID,
				Source:          TypeSourceParameter,
			},
		})
		ir.ReturnTypes = append(ir.ReturnTypes, ReturnTypeFact{
			DefID:    defID,
			FilePath: "src/large.ts",
			Range:    Range{StartLine: i + 2, StartCol: 51, EndLine: i + 2, EndCol: 60},
			Type: TypeRef{
				RawName:         fmt.Sprintf("Result%d", i),
				DeclaredAtScope: scopeID,
				Source:          TypeSourceReturn,
			},
		})
		ir.Frameworks = append(ir.Frameworks, FrameworkFact{
			DefID:                defID,
			FilePath:             "src/large.ts",
			Framework:            "synthetic",
			Reason:               fmt.Sprintf("reason%d", i),
			EntryPointMultiplier: 1,
			Range:                Range{StartLine: i + 2, EndLine: i + 3},
		})
		ir.Domains = append(ir.Domains, DomainFact{
			DefID:    defID,
			FilePath: "src/large.ts",
			Domain:   "synthetic",
			Role:     fmt.Sprintf("role%d", i),
			Range:    Range{StartLine: i + 2, EndLine: i + 3},
		})
	}
	return ir
}
