//go:build ladybugdb

package lbugnative

import (
	"fmt"
	"path/filepath"
	"strings"
	"testing"

	"github.com/tamnguyendinh/anvien/internal/embeddings"
	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/lbugload"
	"github.com/tamnguyendinh/anvien/internal/lbugruntime"
	"github.com/tamnguyendinh/anvien/internal/lbugschema"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
)

func TestNativeLadybugPersistenceReadbackAndStream(t *testing.T) {
	dbPath := filepath.Join(t.TempDir(), "anvien-test.lbug")

	db, err := openNativeDatabase(dbPath, false)
	if err != nil {
		t.Fatalf("open writable db: %v", err)
	}
	conn, err := db.OpenConnection()
	if err != nil {
		db.Close()
		t.Fatalf("open writable connection: %v", err)
	}

	writeRunner := nativeWriteRunner{conn: conn}
	for _, query := range schemaQueries(t) {
		if err := writeRunner.Query(query); err != nil {
			conn.Close()
			db.Close()
			t.Fatalf("schema query failed:\n%s\n%v", query, err)
		}
	}

	export, err := lbugload.ExportGraphCSVs(nativeFixtureGraph(), filepath.Join(t.TempDir(), "csv"))
	if err != nil {
		conn.Close()
		db.Close()
		t.Fatalf("ExportGraphCSVs() error = %v", err)
	}
	loadResult, err := lbugload.LoadCSVExport(writeRunner, export)
	if err != nil {
		conn.Close()
		db.Close()
		t.Fatalf("LoadCSVExport() error = %v", err)
	}
	if loadResult.NodeCopyCount != len(nativeDefinitionTables())+1 || loadResult.RelationshipCopyCount != len(nativeDefinitionTables()) || loadResult.FallbackInsertCount != 0 || loadResult.FallbackInsertFailures != 0 || loadResult.SkippedRelationships != 0 {
		conn.Close()
		db.Close()
		t.Fatalf("unexpected load result: %#v", loadResult)
	}
	if err := writeRunner.Query(embeddings.CreateEmbeddingQuery(embeddings.EmbeddingUpdate{
		NodeID:      "opaque-definition-00",
		Label:       scopeir.NodeFunction,
		ChunkIndex:  0,
		StartLine:   1,
		EndLine:     1,
		Embedding:   []float32{1, 2, 3},
		ContentHash: "native-content-hash",
	})); err != nil {
		conn.Close()
		db.Close()
		t.Fatalf("create embedding row: %v", err)
	}
	conn.Close()
	db.Close()

	readDB, err := openNativeDatabase(dbPath, true)
	if err != nil {
		t.Fatalf("reopen read-only db: %v", err)
	}
	defer readDB.Close()
	readConn, err := readDB.OpenConnection()
	if err != nil {
		t.Fatalf("open read-only connection: %v", err)
	}
	defer readConn.Close()

	readRunner := nativeReadRunner{conn: readConn, silencer: &lbugruntime.StdioSilencer{}}
	rows, err := readRunner.Query("MATCH (f:File) RETURN f.id AS id, f.name AS name, f.content AS content")
	if err != nil {
		t.Fatalf("read file row: %v", err)
	}
	if len(rows) != 1 || rows[0]["id"] != "File:src/app.ts" || rows[0]["name"] != "app.ts" {
		t.Fatalf("unexpected file rows: %#v", rows)
	}
	for index, table := range nativeDefinitionTables() {
		query := "MATCH (n:" + lbugschema.FormatIdent(table) + ") RETURN " +
			"n.id AS id, n.qualifiedName AS qualifiedName, " +
			"n.startLine AS startLine, n.startCol AS startCol, n.endLine AS endLine, n.endCol AS endCol, " +
			"n.selectionStartLine AS selectionStartLine, n.selectionStartCol AS selectionStartCol, " +
			"n.selectionEndLine AS selectionEndLine, n.selectionEndCol AS selectionEndCol, " +
			"n.selectionStartLine IS NULL AS selectionStartLineMissing, " +
			"n.selectionStartCol IS NULL AS selectionStartColMissing, " +
			"n.selectionEndLine IS NULL AS selectionEndLineMissing, " +
			"n.selectionEndCol IS NULL AS selectionEndColMissing"
		definitionRows, err := readRunner.Query(query)
		if err != nil {
			t.Fatalf("read %s definition row: %v", table, err)
		}
		wantID := fmt.Sprintf("opaque-definition-%02d", index)
		if len(definitionRows) != 1 {
			t.Fatalf("%s definition rows = %#v, want one", table, definitionRows)
		}
		row := definitionRows[0]
		if row["id"] != wantID || row["qualifiedName"] != "scope."+strings.ToLower(table) || row["startLine"] != fmt.Sprint(index+1) || row["startCol"] != "0" || row["endLine"] != fmt.Sprint(index+2) || row["endCol"] != fmt.Sprint(index+3) {
			t.Fatalf("%s corrected definition fields drifted: %#v", table, row)
		}
		selectionColumns := []string{"selectionStartLine", "selectionStartCol", "selectionEndLine", "selectionEndCol"}
		if index%2 == 0 {
			wantSelection := []string{fmt.Sprint(index + 1), "0", fmt.Sprint(index + 1), fmt.Sprint(index + 1)}
			for selectionIndex, column := range selectionColumns {
				if row[column] != wantSelection[selectionIndex] || nativeBoolTrue(row[column+"Missing"]) {
					t.Fatalf("%s.%s = %#v (missing=%#v), want %q and present", table, column, row[column], row[column+"Missing"], wantSelection[selectionIndex])
				}
			}
		} else {
			for _, column := range selectionColumns {
				if !nativeBoolTrue(row[column+"Missing"]) {
					t.Fatalf("%s.%s persisted a half/default range: %#v", table, column, row)
				}
			}
		}
	}
	embeddingRows, err := readRunner.Query("MATCH (e:CodeEmbedding) RETURN e.nodeId AS nodeId, e.label AS label, e.chunkIndex AS chunkIndex, e.contentHash AS contentHash")
	if err != nil {
		t.Fatalf("read embedding rows: %v", err)
	}
	if len(embeddingRows) != 1 || embeddingRows[0]["nodeId"] != "opaque-definition-00" || embeddingRows[0]["label"] != "Function" || embeddingRows[0]["chunkIndex"] != "0" || embeddingRows[0]["contentHash"] != "native-content-hash" {
		t.Fatalf("unexpected embedding rows: %#v", embeddingRows)
	}
	missingRows, err := readRunner.Query("MATCH (n:Function) WHERE n.id = '__nonexistent_id__' RETURN n.id AS id")
	if err != nil {
		t.Fatalf("read missing function rows: %v", err)
	}
	if len(missingRows) != 0 {
		t.Fatalf("missing function rows = %#v, want empty", missingRows)
	}
	if _, err := readRunner.Query("MATCH RETURN 1"); err == nil {
		t.Fatalf("malformed Cypher query error = nil, want error")
	}
	if _, err := readRunner.Query("MATCH (n:GhostTable) RETURN n"); err == nil {
		t.Fatalf("non-existent node label query error = nil, want error")
	}
	if _, err := readRunner.Query("CREATE (n:File {id: 'new'})"); err == nil {
		t.Fatalf("read runner write query error = nil, want read-only rejection")
	}

	var streamed []lbugruntime.Row
	count, err := readRunner.Stream("MATCH (a:File)-[r:CodeRelation]->(b:Function) RETURN a.id AS fromId, b.id AS toId, r.type AS type, r.fileHash AS fileHash", func(row lbugruntime.Row) error {
		streamed = append(streamed, row)
		return nil
	})
	if err != nil {
		t.Fatalf("stream relationship rows: %v", err)
	}
	if count != 1 || streamed[0]["fromId"] != "File:src/app.ts" || streamed[0]["toId"] != "opaque-definition-00" || streamed[0]["type"] != "DEFINES" || streamed[0]["fileHash"] != "hash-native-function" {
		t.Fatalf("unexpected streamed rows: count=%d rows=%#v", count, streamed)
	}
	for index, table := range nativeDefinitionTables() {
		rows, err := readRunner.Query("MATCH (a:File)-[r:CodeRelation]->(b:" + lbugschema.FormatIdent(table) + ") RETURN a.id AS fromId, b.id AS toId, r.type AS type")
		if err != nil {
			t.Fatalf("read %s DEFINES closure: %v", table, err)
		}
		wantID := fmt.Sprintf("opaque-definition-%02d", index)
		if len(rows) != 1 || rows[0]["fromId"] != "File:src/app.ts" || rows[0]["toId"] != wantID || rows[0]["type"] != "DEFINES" {
			t.Fatalf("unexpected %s DEFINES rows: %#v", table, rows)
		}
	}
}

type nativeWriteRunner struct {
	conn *nativeConnection
}

func (r nativeWriteRunner) Query(query string) error {
	result, err := r.conn.Query(query)
	if result != nil {
		result.Close()
	}
	return err
}

type nativeReadRunner struct {
	conn     *nativeConnection
	silencer *lbugruntime.StdioSilencer
}

func (r nativeReadRunner) Query(query string) ([]lbugruntime.Row, error) {
	if err := lbugruntime.ValidateReadQuery(query); err != nil {
		return nil, err
	}
	var rows []lbugruntime.Row
	err := r.silencer.Run(func() error {
		result, err := r.conn.Query(query)
		if err != nil {
			return err
		}
		defer result.Close()
		rows, err = result.Rows()
		return err
	})
	return rows, err
}

func (r nativeReadRunner) Stream(query string, onRow func(lbugruntime.Row) error) (int, error) {
	if err := lbugruntime.ValidateReadQuery(query); err != nil {
		return 0, err
	}
	count := 0
	err := r.silencer.Run(func() error {
		result, err := r.conn.Query(query)
		if err != nil {
			return err
		}
		defer result.Close()
		rows, err := result.Rows()
		if err != nil {
			return err
		}
		for _, row := range rows {
			if err := onRow(row); err != nil {
				return err
			}
			count++
		}
		return nil
	})
	return count, err
}

func schemaQueries(t *testing.T) []string {
	t.Helper()
	queries, err := lbugschema.SchemaQueries(3)
	if err != nil {
		t.Fatalf("SchemaQueries() error = %v", err)
	}
	return queries
}

func nativeFixtureGraph() *graph.Graph {
	g := graph.New()
	g.AddNode(graph.Node{ID: "File:src/app.ts", Label: scopeir.NodeFile, Properties: graph.NodeProperties{
		"name": "app.ts", "filePath": "src/app.ts", "content": "export function doWork() {}",
	}})
	for index, table := range nativeDefinitionTables() {
		id := fmt.Sprintf("opaque-definition-%02d", index)
		properties := graph.NodeProperties{
			"name":          strings.ToLower(table),
			"filePath":      "src/app.ts",
			"qualifiedName": "scope." + strings.ToLower(table),
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
		}
		g.AddNode(graph.Node{ID: id, Label: scopeir.NodeLabel(table), Properties: properties})
		g.AddRelationship(graph.Relationship{
			ID:               "rel:file-" + strings.ToLower(table),
			SourceID:         "File:src/app.ts",
			TargetID:         id,
			Type:             graph.RelDefines,
			Confidence:       1,
			Reason:           "native persistence fixture",
			ResolutionSource: "native-test",
			FileHash:         "hash-native-" + strings.ToLower(table),
		})
	}
	return g
}

func nativeDefinitionTables() []string {
	return []string{
		"Function", "Class", "Interface", "CodeElement", "Method",
		"Package", "Struct", "Enum", "Macro", "Typedef", "Union", "Namespace", "Trait", "Impl",
		"TypeAlias", "Const", "Static", "Variable", "Property", "Record", "Delegate", "Annotation",
		"Constructor", "Template", "Module",
	}
}

func nativeBoolTrue(value any) bool {
	text := strings.ToLower(fmt.Sprint(value))
	return text == "true" || text == "1"
}
