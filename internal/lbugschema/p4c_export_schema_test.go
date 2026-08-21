package lbugschema

import (
	"strings"
	"testing"
)

func TestP4CExportSchemaPreservesFactAndSourceProvenanceFields(t *testing.T) {
	wantColumns := []string{
		"id STRING", "name STRING", "filePath STRING", "fileHash STRING", "kind STRING", "exportedName STRING",
		"localName STRING", "localDefId STRING", "localDefinitionNodeId STRING", "targetRaw STRING", "targetExportedName STRING",
		"meanings STRING[]", "typeOnly BOOLEAN", "startLine INT64", "startCol INT64", "endLine INT64", "endCol INT64",
		"selectionStartLine INT64", "selectionStartCol INT64", "selectionEndLine INT64", "selectionEndCol INT64",
		"statementStartLine INT64", "statementStartCol INT64", "statementEndLine INT64", "statementEndCol INT64", "siteKind STRING",
		"appLayer STRING", "functionalArea STRING",
	}
	schema := NodeSchema("Export")
	for _, want := range wantColumns {
		if !strings.Contains(schema, want) {
			t.Fatalf("Export schema missing %q:\n%s", want, schema)
		}
	}
	if !containsString(NodeTables, "Export") {
		t.Fatalf("NodeTables missing Export: %#v", NodeTables)
	}
	foundPair := false
	for _, pair := range RelationPairs {
		if pair.From == "File" && pair.To == "Export" {
			foundPair = true
			break
		}
	}
	if !foundPair {
		t.Fatal("RelationPairs missing File -> Export containment pair")
	}
}
