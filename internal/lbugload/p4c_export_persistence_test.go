package lbugload

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
)

func TestP4CExportCSVAndLoaderPreserveFactFields(t *testing.T) {
	tempRoot := filepath.Join("..", "..", ".tmp", "p4c-tests")
	if err := os.MkdirAll(tempRoot, 0o755); err != nil {
		t.Fatalf("create repo-local test temp root: %v", err)
	}
	tempDir, err := os.MkdirTemp(tempRoot, "export-")
	if err != nil {
		t.Fatalf("create repo-local test temp dir: %v", err)
	}
	t.Cleanup(func() { _ = os.RemoveAll(tempDir) })

	targetRaw := "./remote"
	g := graph.New()
	fileID := "File:src/exports.ts"
	directExportID := "Export:src/exports.ts:0"
	reexportID := "Export:src/exports.ts:1"
	definitionID := "Variable:src/exports.ts:runtime"
	g.AddNode(graph.Node{ID: fileID, Label: scopeir.NodeFile, Properties: graph.NodeProperties{"name": "exports.ts", "filePath": "src/exports.ts"}})
	g.AddNode(graph.Node{ID: definitionID, Label: scopeir.NodeVariable, Properties: graph.NodeProperties{"name": "runtime", "filePath": "src/exports.ts", "qualifiedName": "runtime", "startLine": 1, "endLine": 1, "isExported": true}})
	g.AddNode(graph.Node{ID: directExportID, Label: scopeir.NodeLabel("Export"), Properties: graph.NodeProperties{
		"name": "runtime", "filePath": "src/exports.ts", "fileHash": "hash-exports", "kind": "direct", "exportedName": "runtime", "localName": "runtime", "localDefId": "def:src/exports.ts#1:0:Variable:runtime", "localDefinitionNodeId": definitionID, "meanings": []string{"value"}, "typeOnly": false, "startLine": 1, "startCol": 0, "endLine": 1, "endCol": 14, "selectionStartLine": 1, "selectionStartCol": 7, "selectionEndLine": 1, "selectionEndCol": 14, "statementStartLine": 1, "statementStartCol": 0, "statementEndLine": 1, "statementEndCol": 14, "siteKind": "export_declaration",
	}})
	g.AddNode(graph.Node{ID: reexportID, Label: scopeir.NodeLabel("Export"), Properties: graph.NodeProperties{
		"name": "remote", "filePath": "src/exports.ts", "fileHash": "hash-exports", "kind": "reexport", "exportedName": "remote", "targetRaw": targetRaw, "targetExportedName": "remote", "meanings": []string{"value"}, "typeOnly": false, "startLine": 3, "startCol": 0, "endLine": 3, "endCol": 28, "statementStartLine": 3, "statementStartCol": 0, "statementEndLine": 3, "statementEndCol": 28, "siteKind": "export_specifier",
	}})
	g.AddRelationship(graph.Relationship{ID: "rel:file-direct-export", SourceID: fileID, TargetID: directExportID, Type: graph.RelContains, Confidence: 1, Reason: "source export fact"})
	g.AddRelationship(graph.Relationship{ID: "rel:file-reexport", SourceID: fileID, TargetID: reexportID, Type: graph.RelContains, Confidence: 1, Reason: "source export fact"})

	export, err := ExportGraphCSVs(g, filepath.Join(tempDir, "csv"))
	if err != nil {
		t.Fatalf("ExportGraphCSVs() error = %v", err)
	}
	if export.Metrics.SkippedNodes != 0 || export.Metrics.SkippedRelationships != 0 {
		t.Fatalf("export skipped fact data: %#v", export.Metrics)
	}
	if export.Metrics.RowsByTable["Export"] != 2 {
		t.Fatalf("Export CSV row count = %d, want 2", export.Metrics.RowsByTable["Export"])
	}
	rows, err := readP4CCSV(filepath.Join(export.CSVDir, "export.csv"))
	if err != nil {
		t.Fatalf("read Export CSV: %v", err)
	}
	if len(rows) != 3 || len(rows[0]) != len(rows[1]) || len(rows[0]) != len(rows[2]) {
		t.Fatalf("Export CSV rows = %#v, want header plus two equal-width rows", rows)
	}
	var directRow, reexportRow []string
	for _, row := range rows[1:] {
		switch p4cCSVValue(rows[0], row, "kind") {
		case "direct":
			directRow = row
		case "reexport":
			reexportRow = row
		}
	}
	if directRow == nil || reexportRow == nil {
		t.Fatalf("Export CSV rows missing direct/reexport records: %#v", rows)
	}
	for _, column := range []string{"kind", "exportedName", "localDefId", "localDefinitionNodeId", "meanings", "typeOnly", "statementStartLine", "siteKind"} {
		if p4cCSVValue(rows[0], directRow, column) == "" {
			t.Fatalf("Export CSV lost direct %s: %#v", column, rows)
		}
	}
	if got := p4cCSVValue(rows[0], reexportRow, "targetRaw"); got != targetRaw {
		t.Fatalf("targetRaw = %q, want %q", got, targetRaw)
	}
	if got := p4cCSVValue(rows[0], directRow, "meanings"); !strings.Contains(got, "value") {
		t.Fatalf("meanings = %q, want value lane", got)
	}

	runner := &recordingRunner{}
	loaded, err := LoadCSVExport(runner, export)
	if err != nil {
		t.Fatalf("LoadCSVExport() error = %v", err)
	}
	if loaded.NodeCopyCount != 3 || loaded.RelationshipCopyCount != 1 || loaded.FallbackInsertCount != 0 || loaded.SkippedRelationships != 0 {
		t.Fatalf("loader parity = %#v, want three node-table/two-row and one File->Export pair with no fallback", loaded)
	}
	joined := strings.Join(runner.queries, "\n")
	if !strings.Contains(joined, "COPY Export(") || !strings.Contains(joined, `from="File", to="Export"`) {
		t.Fatalf("loader did not use Export table and File -> Export pair:\n%s", joined)
	}
}

func readP4CCSV(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return csv.NewReader(file).ReadAll()
}

func p4cCSVValue(header, row []string, column string) string {
	for index, candidate := range header {
		if candidate == column && index < len(row) {
			return row[index]
		}
	}
	return ""
}
