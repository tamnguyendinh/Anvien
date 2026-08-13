package lbugload

import (
	"fmt"
	"path/filepath"
	"strings"
	"testing"

	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/lbugschema"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
)

func TestDefinitionCSVProjectionPreservesCorrectedFactsAcrossEveryTableBranch(t *testing.T) {
	definitionTables := p2bDefinitionTables()
	g := graph.New()
	fileID := "File:src/all.ts"
	g.AddNode(graph.Node{ID: fileID, Label: scopeir.NodeFile, Properties: graph.NodeProperties{"name": "all.ts", "filePath": "src/all.ts"}})

	wantSelection := map[string]bool{}
	wantIDs := map[string]string{}
	for index, table := range definitionTables {
		id := fmt.Sprintf("opaque-definition-%02d", index)
		qualifiedName := "scope." + strings.ToLower(table)
		properties := graph.NodeProperties{
			"name":          strings.ToLower(table),
			"filePath":      "src/all.ts",
			"qualifiedName": qualifiedName,
			"startLine":     index + 1,
			"startCol":      0,
			"endLine":       index + 2,
			"endCol":        index + 3,
		}
		if index%2 == 0 {
			properties["selectionStartLine"] = index + 1
			properties["selectionStartCol"] = 0
			properties["selectionEndLine"] = index + 1
			properties["selectionEndCol"] = index + 1
			wantSelection[table] = true
		}
		g.AddNode(graph.Node{ID: id, Label: scopeir.NodeLabel(table), Properties: properties})
		g.AddRelationship(graph.Relationship{
			ID:         "rel:defines:" + table,
			SourceID:   fileID,
			TargetID:   id,
			Type:       graph.RelDefines,
			Confidence: 1,
		})
		wantIDs[table] = id
	}

	export, err := ExportGraphCSVs(g, filepath.Join(t.TempDir(), "csv"))
	if err != nil {
		t.Fatalf("ExportGraphCSVs() error = %v", err)
	}
	if export.Metrics.SkippedNodes != 0 || export.Metrics.SkippedRelationships != 0 {
		t.Fatalf("export dropped corrected records: %#v", export.Metrics)
	}
	if export.RelationshipRows != len(definitionTables) {
		t.Fatalf("relationship rows = %d, want %d", export.RelationshipRows, len(definitionTables))
	}

	for index, table := range definitionTables {
		rows := readCSV(t, filepath.Join(export.CSVDir, strings.ToLower(table)+".csv"))
		if len(rows) != 2 {
			t.Fatalf("%s rows = %d, want header plus one record", table, len(rows))
		}
		header, row := rows[0], rows[1]
		wantHeader := p2bDefinitionColumns(table)
		if strings.Join(header, "\x00") != strings.Join(wantHeader, "\x00") {
			t.Fatalf("%s header = %#v, want exact schema order %#v", table, header, wantHeader)
		}
		if len(header) != len(row) {
			t.Fatalf("%s header/row width = %d/%d", table, len(header), len(row))
		}
		wantValues := map[string]string{
			"id":            wantIDs[table],
			"qualifiedName": "scope." + strings.ToLower(table),
			"startLine":     fmt.Sprint(index + 1),
			"startCol":      "0",
			"endLine":       fmt.Sprint(index + 2),
			"endCol":        fmt.Sprint(index + 3),
		}
		for column, want := range wantValues {
			if got := csvColumnValue(t, header, row, column); got != want {
				t.Fatalf("%s.%s = %q, want %q", table, column, got, want)
			}
		}
		selectionColumns := []string{"selectionStartLine", "selectionStartCol", "selectionEndLine", "selectionEndCol"}
		if wantSelection[table] {
			want := []string{fmt.Sprint(index + 1), "0", fmt.Sprint(index + 1), fmt.Sprint(index + 1)}
			for selectionIndex, column := range selectionColumns {
				if got := csvColumnValue(t, header, row, column); got != want[selectionIndex] {
					t.Fatalf("%s.%s = %q, want %q", table, column, got, want[selectionIndex])
				}
			}
		} else {
			for _, column := range selectionColumns {
				if got := csvColumnValue(t, header, row, column); got != "" {
					t.Fatalf("%s.%s = %q, want blank/NULL input", table, column, got)
				}
			}
		}
	}

	endpointRows := readCSV(t, export.RelationshipCSVPath)
	seenTargets := map[string]struct{}{}
	for _, row := range endpointRows[1:] {
		if row[0] != fileID || row[2] != string(graph.RelDefines) {
			t.Fatalf("DEFINES row drifted: %#v", row)
		}
		seenTargets[row[1]] = struct{}{}
	}
	for _, id := range wantIDs {
		if _, ok := seenTargets[id]; !ok {
			t.Fatalf("missing DEFINES target %q in relationship CSV", id)
		}
	}

	if len(export.RelationshipPairFiles) != len(definitionTables) {
		t.Fatalf("relationship pair files = %d, want %d", len(export.RelationshipPairFiles), len(definitionTables))
	}
	for _, pair := range export.RelationshipPairFiles {
		if pair.From != "File" || !pair.CopySupported || pair.Rows != 1 {
			t.Fatalf("DEFINES pair not load-closed: %#v", pair)
		}
	}

	runner := &recordingRunner{}
	loadResult, err := LoadCSVExport(runner, export)
	if err != nil {
		t.Fatalf("LoadCSVExport() error = %v", err)
	}
	if loadResult.NodeCopyCount != len(definitionTables)+1 || loadResult.RelationshipCopyCount != len(definitionTables) || loadResult.FallbackInsertCount != 0 || loadResult.FallbackInsertFailures != 0 || loadResult.SkippedRelationships != 0 {
		t.Fatalf("all-label load closure drifted: %#v", loadResult)
	}
	joinedQueries := strings.Join(runner.queries, "\n")
	for _, table := range definitionTables {
		wantCopy := "COPY " + lbugschema.FormatIdent(table) + "(" + strings.Join(p2bDefinitionColumns(table), ", ") + ")"
		if !strings.Contains(joinedQueries, wantCopy) {
			t.Fatalf("load queries missing exact %s node COPY %q:\n%s", table, wantCopy, joinedQueries)
		}
		if !strings.Contains(joinedQueries, `from="File", to="`+table+`"`) {
			t.Fatalf("load queries missing File -> %s DEFINES COPY:\n%s", table, joinedQueries)
		}
	}
}

func TestDefinitionCSVProjectionRejectsIncompleteSelectionRangeInEveryDispatchFamily(t *testing.T) {
	for _, table := range []string{"Function", "Method", "Variable"} {
		t.Run(table, func(t *testing.T) {
			g := graph.New()
			g.AddNode(graph.Node{ID: "opaque-" + strings.ToLower(table), Label: scopeir.NodeLabel(table), Properties: graph.NodeProperties{
				"name": "worker", "filePath": "src/app.ts", "qualifiedName": "worker",
				"startLine": 1, "startCol": 0, "endLine": 2, "endCol": 1,
				"selectionStartLine": 1,
			}})

			_, err := ExportGraphCSVs(g, filepath.Join(t.TempDir(), "csv"))
			if err == nil || !strings.Contains(err.Error(), "incomplete selection range: 1 of 4 coordinates present") {
				t.Fatalf("ExportGraphCSVs() error = %v, want family-scoped incomplete selection range", err)
			}
		})
	}
}

func TestSelectionRangeValidationDoesNotRejectUnrelatedNodeTables(t *testing.T) {
	g := graph.New()
	g.AddNode(graph.Node{ID: "File:src/app.ts", Label: scopeir.NodeFile, Properties: graph.NodeProperties{
		"name": "app.ts", "filePath": "src/app.ts", "selectionStartLine": 1,
	}})

	export, err := ExportGraphCSVs(g, filepath.Join(t.TempDir(), "csv"))
	if err != nil {
		t.Fatalf("ExportGraphCSVs(unrelated partial selection property) error = %v", err)
	}
	rows := readCSV(t, filepath.Join(export.CSVDir, "file.csv"))
	if len(rows) != 2 || len(rows[0]) != len(fileNodeColumns) {
		t.Fatalf("unrelated File projection drifted: %#v", rows)
	}
	for _, column := range rows[0] {
		if strings.HasPrefix(column, "selection") {
			t.Fatalf("unrelated File projection unexpectedly persisted %q: %#v", column, rows)
		}
	}
}

func TestDefinitionCSVProjectionCoversEveryDefaultSchemaTable(t *testing.T) {
	for _, table := range lbugschema.NodeTables {
		if _, explicit := nodeColumnLookup[table]; explicit {
			continue
		}
		columns, ok := nodeColumns(table)
		if !ok {
			t.Fatalf("default valid table %q has no CSV columns", table)
		}
		for _, want := range []string{"qualifiedName", "startLine", "startCol", "endLine", "endCol", "selectionStartLine", "selectionStartCol", "selectionEndLine", "selectionEndCol"} {
			if !containsColumn(columns, want) {
				t.Fatalf("default table %q missing corrected column %q: %#v", table, want, columns)
			}
		}
	}
}

func csvColumnValue(t *testing.T, header []string, row []string, column string) string {
	t.Helper()
	for index, candidate := range header {
		if candidate == column {
			return row[index]
		}
	}
	t.Fatalf("CSV header missing column %q: %#v", column, header)
	return ""
}

func p2bDefinitionTables() []string {
	return []string{
		"Function", "Class", "Interface", "CodeElement", "Method",
		"Package", "Struct", "Enum", "Macro", "Typedef", "Union", "Namespace", "Trait", "Impl",
		"TypeAlias", "Const", "Static", "Variable", "Property", "Record", "Delegate", "Annotation",
		"Constructor", "Template", "Module",
	}
}

func p2bDefinitionColumns(table string) []string {
	base := []string{
		"id", "name", "filePath", "qualifiedName",
		"startLine", "startCol", "endLine", "endCol",
		"selectionStartLine", "selectionStartCol", "selectionEndLine", "selectionEndCol",
	}
	switch table {
	case "Function", "Class", "Interface", "CodeElement":
		base = append(base, "isExported", "content", "description")
	case "Method":
		base = append(base, "isExported", "content", "description", "parameterCount", "returnType")
	default:
		base = append(base, "content", "description")
	}
	return append(base, "appLayer", "functionalArea")
}
