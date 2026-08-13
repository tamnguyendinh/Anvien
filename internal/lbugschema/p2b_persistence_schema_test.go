package lbugschema

import (
	"reflect"
	"strings"
	"testing"
)

func TestDefinitionSchemasPreserveCorrectedFieldsAcrossEveryTableBranch(t *testing.T) {
	definitionTables := []string{
		"Function", "Class", "Interface", "CodeElement", "Method",
		"Package", "Struct", "Enum", "Macro", "Typedef", "Union", "Namespace", "Trait", "Impl",
		"TypeAlias", "Const", "Static", "Variable", "Property", "Record", "Delegate", "Annotation",
		"Constructor", "Template", "Module",
	}
	for _, table := range definitionTables {
		schema := NodeSchema(table)
		gotColumns := p2bSchemaColumns(schema)
		wantColumns := p2bDefinitionSchemaColumns(table)
		if !reflect.DeepEqual(gotColumns, wantColumns) {
			t.Fatalf("%s schema columns = %#v, want exact projection parity %#v\n%s", table, gotColumns, wantColumns, schema)
		}
	}
}

func TestEmbeddingSchemaPersistsExplicitLabel(t *testing.T) {
	schema, err := EmbeddingSchema(3)
	if err != nil {
		t.Fatalf("EmbeddingSchema(3) error = %v", err)
	}
	for _, want := range []string{"nodeId STRING", "label STRING", "chunkIndex INT32", "embedding FLOAT[3]"} {
		if !strings.Contains(schema, want) {
			t.Fatalf("embedding schema missing %q:\n%s", want, schema)
		}
	}
	if strings.Index(schema, "label STRING") <= strings.Index(schema, "nodeId STRING") {
		t.Fatalf("embedding label column is not adjacent to node identity:\n%s", schema)
	}
}

func TestDefinesSchemaCoversEveryDefinitionEndpointTable(t *testing.T) {
	wantTargets := map[string]bool{
		"Function": true, "Class": true, "Interface": true, "Method": true, "CodeElement": true,
		"Package": true, "Struct": true, "Enum": true, "Macro": true, "Typedef": true, "Union": true,
		"Namespace": true, "Trait": true, "Impl": true, "TypeAlias": true, "Const": true, "Static": true,
		"Variable": true, "Property": true, "Record": true, "Delegate": true, "Annotation": true,
		"Constructor": true, "Template": true, "Module": true,
	}
	seen := map[string]bool{}
	for _, pair := range RelationPairs {
		if pair.From == "File" && wantTargets[pair.To] {
			seen[pair.To] = true
		}
	}
	for target := range wantTargets {
		if !seen[target] {
			t.Fatalf("relation schema missing File -> %s endpoint pair", target)
		}
	}
}

func p2bSchemaColumns(schema string) []string {
	lines := strings.Split(schema, "\n")
	columns := make([]string, 0, len(lines))
	for _, line := range lines[1:] {
		line = strings.TrimSpace(strings.TrimSuffix(line, ","))
		if line == "" || line == ")" || strings.HasPrefix(line, "PRIMARY KEY") {
			continue
		}
		columns = append(columns, line)
	}
	return columns
}

func p2bDefinitionSchemaColumns(table string) []string {
	columns := []string{
		"id STRING", "name STRING", "filePath STRING", "qualifiedName STRING",
		"startLine INT64", "startCol INT64", "endLine INT64", "endCol INT64",
		"selectionStartLine INT64", "selectionStartCol INT64", "selectionEndLine INT64", "selectionEndCol INT64",
	}
	switch table {
	case "Function", "Class", "Interface", "CodeElement":
		columns = append(columns, "isExported BOOLEAN", "content STRING", "description STRING")
	case "Method":
		columns = append(columns, "isExported BOOLEAN", "content STRING", "description STRING", "parameterCount INT32", "returnType STRING")
	default:
		columns = append(columns, "content STRING", "description STRING")
	}
	return append(columns, "appLayer STRING", "functionalArea STRING")
}
