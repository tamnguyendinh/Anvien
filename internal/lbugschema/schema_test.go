package lbugschema

import (
	"reflect"
	"strings"
	"testing"
)

func TestSchemaConstantsMatchFrozenContract(t *testing.T) {
	wantNodeTables := expectedFrozenNodeTables()
	if !reflect.DeepEqual(NodeTables, wantNodeTables) {
		t.Fatalf("NodeTables drift\nwant %#v\ngot  %#v", wantNodeTables, NodeTables)
	}
	wantRelationshipTypes := expectedFrozenRelationshipTypes()
	if !reflect.DeepEqual(RelationshipTypes, wantRelationshipTypes) {
		t.Fatalf("RelationshipTypes drift\nwant %#v\ngot  %#v", wantRelationshipTypes, RelationshipTypes)
	}
	if RelTableName != "CodeRelation" {
		t.Fatalf("RelTableName = %q, want CodeRelation", RelTableName)
	}

	gotColumns := map[string]string{}
	for _, col := range relationColumns() {
		gotColumns[col.Name] = col.Type
	}
	wantColumns := expectedFrozenRelationshipColumns()
	if !reflect.DeepEqual(gotColumns, wantColumns) {
		t.Fatalf("relation columns drift\nwant %#v\ngot  %#v", wantColumns, gotColumns)
	}
}

func TestSchemaQueriesPreserveDDLShape(t *testing.T) {
	queries, err := SchemaQueries(DefaultEmbeddingDims)
	if err != nil {
		t.Fatalf("SchemaQueries() error = %v", err)
	}
	if len(queries) != len(NodeTables)+2 {
		t.Fatalf("SchemaQueries() len = %d, want %d", len(queries), len(NodeTables)+2)
	}
	if len(RelationPairs) != 286 {
		t.Fatalf("RelationPairs len = %d, want 286", len(RelationPairs))
	}

	fileSchema := NodeSchema("File")
	for _, want := range []string{"CREATE NODE TABLE File", "content STRING", "appLayer STRING", "functionalArea STRING", "PRIMARY KEY (id)"} {
		if !strings.Contains(fileSchema, want) {
			t.Fatalf("File schema missing %q:\n%s", want, fileSchema)
		}
	}
	methodSchema := NodeSchema("Method")
	for _, want := range []string{"parameterCount INT32", "returnType STRING", "appLayer STRING", "functionalArea STRING"} {
		if !strings.Contains(methodSchema, want) {
			t.Fatalf("Method schema missing %q:\n%s", want, methodSchema)
		}
	}
	if !strings.Contains(NodeSchema("Struct"), "CREATE NODE TABLE `Struct`") {
		t.Fatalf("Struct schema must quote the table name:\n%s", NodeSchema("Struct"))
	}
	resolutionGapSchema := NodeSchema("ResolutionGap")
	for _, want := range []string{"CREATE NODE TABLE ResolutionGap", "sourceSiteId STRING", "sourceNodeId STRING", "gapKind STRING", "classification STRING", "actionability STRING", "count INT32", "appLayer STRING", "functionalArea STRING"} {
		if !strings.Contains(resolutionGapSchema, want) {
			t.Fatalf("ResolutionGap schema missing %q:\n%s", want, resolutionGapSchema)
		}
	}

	relationSchema := RelationSchema()
	for _, want := range []string{
		"CREATE REL TABLE CodeRelation",
		"FROM File TO File",
		"FROM File TO Package",
		"FROM File TO `Struct`",
		"FROM Function TO `Variable`",
		"FROM Method TO `Const`",
		"FROM Method TO `Static`",
		"FROM Method TO `Variable`",
		"FROM Package TO `Const`",
		"FROM Package TO `Property`",
		"FROM Package TO `Static`",
		"FROM Package TO `Variable`",
		"FROM `Const` TO `Property`",
		"FROM `Const` TO `Static`",
		"FROM `Const` TO `Variable`",
		"FROM `Variable` TO `Const`",
		"FROM `Variable` TO Function",
		"FROM `Variable` TO `Static`",
		"FROM `Variable` TO `Variable`",
		"FROM `Const` TO `Const`",
		"FROM `Const` TO Function",
		"FROM `Constructor` TO `Property`",
		"FROM CodeElement TO CodeElement",
		"FROM CodeElement TO `Module`",
		"FROM `Module` TO `Namespace`",
		"FROM `Namespace` TO Function",
		"FROM `TypeAlias` TO Method",
		"FROM `TypeAlias` TO `Property`",
		"FROM Function TO ResolutionGap",
		"FROM Method TO ResolutionGap",
		"type STRING",
		"confidence DOUBLE",
		"resolutionSource STRING",
		"evidence STRING",
		"fileHash STRING",
	} {
		if !strings.Contains(relationSchema, want) {
			t.Fatalf("relation schema missing %q", want)
		}
	}
}

func TestSchemaSurfaceCoversLegacyCoreAndModernNodeTypes(t *testing.T) {
	for _, table := range []string{
		"File", "Folder", "Function", "Class", "Interface", "Method", "CodeElement", "Community", "Process",
		"Package", "Section", "Struct", "Enum", "Macro", "Typedef", "Union", "Namespace", "Trait", "Impl",
		"TypeAlias", "Const", "Static", "Variable", "Property", "Record", "Delegate", "Annotation", "Constructor",
		"Template", "Module", "Route", "Tool",
		"ResolutionGap",
	} {
		if !containsString(NodeTables, table) {
			t.Fatalf("NodeTables missing %q", table)
		}
	}
	if len(NodeTables) != 33 {
		t.Fatalf("NodeTables length = %d, want 33", len(NodeTables))
	}

	for _, relationType := range []string{
		"CONTAINS", "DEFINES", "IMPORTS", "CALLS", "USES", "INHERITS", "EXTENDS", "IMPLEMENTS",
		"HAS_METHOD", "HAS_PROPERTY", "ACCESSES", "METHOD_OVERRIDES", "OVERRIDES", "METHOD_IMPLEMENTS",
		"MEMBER_OF", "STEP_IN_PROCESS", "HANDLES_ROUTE", "FETCHES", "HANDLES_TOOL", "ENTRY_POINT_OF",
		"WRAPS", "QUERIES", "HAS_RESOLUTION_GAP",
	} {
		if !containsString(RelationshipTypes, relationType) {
			t.Fatalf("RelationshipTypes missing %q", relationType)
		}
	}

	for _, want := range []string{"startLine INT64", "endLine INT64", "isExported BOOLEAN"} {
		if !strings.Contains(NodeSchema("Function"), want) {
			t.Fatalf("Function schema missing %q:\n%s", want, NodeSchema("Function"))
		}
	}
	for _, want := range []string{"heuristicLabel STRING", "cohesion DOUBLE"} {
		if !strings.Contains(NodeSchema("Community"), want) {
			t.Fatalf("Community schema missing %q:\n%s", want, NodeSchema("Community"))
		}
	}
	for _, want := range []string{"processType STRING", "stepCount INT32"} {
		if !strings.Contains(NodeSchema("Process"), want) {
			t.Fatalf("Process schema missing %q:\n%s", want, NodeSchema("Process"))
		}
	}

	relationSchema := RelationSchema()
	for _, want := range []string{
		"FROM Class TO Method",
		"FROM Class TO `Constructor`",
		"FROM `Struct` TO Method",
		"FROM `Trait` TO `Constructor`",
		"FROM `Impl` TO `Property`",
		"FROM `Record` TO `Property`",
		"FROM `Property` TO Method",
		"FROM `TypeAlias` TO `Property`",
	} {
		if !strings.Contains(relationSchema, want) {
			t.Fatalf("relation schema missing %q", want)
		}
	}

	queries, err := SchemaQueries(DefaultEmbeddingDims)
	if err != nil {
		t.Fatalf("SchemaQueries() error = %v", err)
	}
	if got := len(queries); got != len(NodeTables)+2 {
		t.Fatalf("SchemaQueries() length = %d, want %d", got, len(NodeTables)+2)
	}
	joinedNodes := strings.Join(NodeSchemaQueries(), "\n")
	relationIndex := indexOfQuery(queries, RelationSchema())
	if relationIndex <= 0 || !strings.Contains(joinedNodes, queries[relationIndex-1]) {
		t.Fatalf("relation schema must come after node schemas")
	}
}

func TestEmbeddingAndIndexQueries(t *testing.T) {
	if _, err := EmbeddingSchema(0); err == nil {
		t.Fatalf("EmbeddingSchema(0) expected error")
	}
	schema, err := EmbeddingSchema(768)
	if err != nil {
		t.Fatalf("EmbeddingSchema(768) error = %v", err)
	}
	for _, want := range []string{"CREATE NODE TABLE CodeEmbedding", "embedding FLOAT[768]", "contentHash STRING"} {
		if !strings.Contains(schema, want) {
			t.Fatalf("embedding schema missing %q:\n%s", want, schema)
		}
	}
	if got := CreateVectorIndexQuery(); got != "CALL CREATE_VECTOR_INDEX('CodeEmbedding', 'code_embedding_idx', 'embedding', metric := 'cosine')" {
		t.Fatalf("CreateVectorIndexQuery() = %q", got)
	}
	if got := FTSIndexQueries(); len(got) != 5 || !strings.Contains(got[0], "file_fts") {
		t.Fatalf("unexpected FTSIndexQueries(): %#v", got)
	}
}

func indexOfQuery(queries []string, query string) int {
	for index, candidate := range queries {
		if candidate == query {
			return index
		}
	}
	return -1
}

func containsString(values []string, want string) bool {
	for _, value := range values {
		if value == want {
			return true
		}
	}
	return false
}

func expectedFrozenNodeTables() []string {
	return []string{
		"File",
		"Folder",
		"Function",
		"Class",
		"Interface",
		"Method",
		"CodeElement",
		"Community",
		"Process",
		"Package",
		"Section",
		"Struct",
		"Enum",
		"Macro",
		"Typedef",
		"Union",
		"Namespace",
		"Trait",
		"Impl",
		"TypeAlias",
		"Const",
		"Static",
		"Variable",
		"Property",
		"Record",
		"Delegate",
		"Annotation",
		"Constructor",
		"Template",
		"Module",
		"Route",
		"Tool",
		"ResolutionGap",
	}
}

func expectedFrozenRelationshipTypes() []string {
	return []string{
		"CONTAINS",
		"DEFINES",
		"IMPORTS",
		"CALLS",
		"USES",
		"INHERITS",
		"EXTENDS",
		"IMPLEMENTS",
		"HAS_METHOD",
		"HAS_PROPERTY",
		"ACCESSES",
		"METHOD_OVERRIDES",
		"OVERRIDES",
		"METHOD_IMPLEMENTS",
		"MEMBER_OF",
		"STEP_IN_PROCESS",
		"HANDLES_ROUTE",
		"FETCHES",
		"HANDLES_TOOL",
		"ENTRY_POINT_OF",
		"WRAPS",
		"QUERIES",
		"HAS_RESOLUTION_GAP",
	}
}

func expectedFrozenRelationshipColumns() map[string]string {
	return map[string]string{
		"type":             "STRING",
		"confidence":       "DOUBLE",
		"reason":           "STRING",
		"step":             "INT32",
		"resolutionSource": "STRING",
		"evidence":         "STRING",
		"fileHash":         "STRING",
		"sourceSiteId":     "STRING",
		"sourceSiteIds":    "STRING",
		"sourceSiteCount":  "INT32",
		"sourceSiteStatus": "STRING",
		"proofKind":        "STRING",
		"targetRole":       "STRING",
		"targetText":       "STRING",
		"filePath":         "STRING",
		"startLine":        "INT64",
		"startCol":         "INT64",
		"endLine":          "INT64",
		"endCol":           "INT64",
	}
}
