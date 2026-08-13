//go:build ladybugdb

package embeddings

import (
	"context"
	"path/filepath"
	"strings"
	"testing"

	"github.com/tamnguyendinh/anvien/internal/lbugnative"
	"github.com/tamnguyendinh/anvien/internal/lbugruntime"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
)

func TestNativeLadybugSemanticSearchHydratesPersistedExplicitLabels(t *testing.T) {
	dbPath := filepath.Join(t.TempDir(), "semantic-search.lbug")
	writer, err := lbugnative.OpenWriteRunnerWithEmbeddingDims(dbPath, 3)
	if err != nil {
		t.Fatalf("OpenWriteRunnerWithEmbeddingDims() error = %v", err)
	}

	queries := []string{
		"CREATE (n:Function {id: 'opaque-function-a', name: 'sharedName', filePath: 'src/function-a.ts', startLine: 10, endLine: 12})",
		"CREATE (n:Function {id: 'opaque-function-b', name: 'sharedName', filePath: 'src/function-b.ts', startLine: 14, endLine: 16})",
		"CREATE (n:Method {id: 'opaque-method', name: 'sharedName', filePath: 'src/method.ts', startLine: 20, endLine: 22})",
		CreateEmbeddingQuery(EmbeddingUpdate{
			NodeID: "opaque-function-a", Label: scopeir.NodeFunction, ChunkIndex: 0,
			StartLine: 10, EndLine: 12, Embedding: []float32{1, 2, 3}, ContentHash: "function-a-hash",
		}),
		CreateEmbeddingQuery(EmbeddingUpdate{
			NodeID: "opaque-function-b", Label: scopeir.NodeFunction, ChunkIndex: 0,
			StartLine: 14, EndLine: 16, Embedding: []float32{1, 2, 3}, ContentHash: "function-b-hash",
		}),
		CreateEmbeddingQuery(EmbeddingUpdate{
			NodeID: "opaque-method", Label: scopeir.NodeMethod, ChunkIndex: 0,
			StartLine: 20, EndLine: 22, Embedding: []float32{1, 2, 3}, ContentHash: "method-hash",
		}),
	}
	for _, query := range queries {
		if err := writer.Query(query); err != nil {
			_ = writer.Close()
			t.Fatalf("native write query failed:\n%s\n%v", query, err)
		}
	}
	rows, err := writer.QueryRows(
		"MATCH (e:CodeEmbedding) RETURN e.nodeId AS nodeId, e.label AS label, e.chunkIndex AS chunkIndex, e.startLine AS startLine, e.endLine AS endLine, 0.1 AS distance",
	)
	if err != nil {
		_ = writer.Close()
		t.Fatalf("read persisted embedding rows: %v", err)
	}
	runner := &nativeSemanticSearchRunner{native: writer, vectorRows: rows}
	results, err := SemanticSearch(
		context.Background(),
		runner,
		&recordingEmbedder{dimensions: 3},
		"shared name",
		SearchOptions{Limit: 3, Dimensions: 3, MaxDistance: 0.5},
	)
	if err != nil {
		_ = writer.Close()
		t.Fatalf("SemanticSearch() error = %v", err)
	}
	if !runner.sawVectorQuery {
		_ = writer.Close()
		t.Fatalf("SemanticSearch() did not issue the vector index query")
	}
	if err := writer.Close(); err != nil {
		t.Fatalf("writer.Close() error = %v", err)
	}
	if len(results) != 3 {
		t.Fatalf("len(results) = %d, want 3: %#v", len(results), results)
	}

	byID := make(map[string]SearchResult, len(results))
	for _, result := range results {
		byID[result.NodeID] = result
	}
	functionA := byID["opaque-function-a"]
	if functionA.Label != "Function" || functionA.Name != "sharedName" || functionA.FilePath != "src/function-a.ts" || functionA.StartLine != 10 || functionA.EndLine != 12 {
		t.Fatalf("Function A hydration drifted: %#v", functionA)
	}
	functionB := byID["opaque-function-b"]
	if functionB.Label != "Function" || functionB.Name != "sharedName" || functionB.FilePath != "src/function-b.ts" || functionB.StartLine != 14 || functionB.EndLine != 16 {
		t.Fatalf("Function B hydration drifted: %#v", functionB)
	}
	method := byID["opaque-method"]
	if method.Label != "Method" || method.Name != "sharedName" || method.FilePath != "src/method.ts" || method.StartLine != 20 || method.EndLine != 22 {
		t.Fatalf("Method hydration drifted: %#v", method)
	}
}

type nativeSemanticSearchRunner struct {
	native         lbugnative.WriteRunner
	vectorRows     []lbugruntime.Row
	sawVectorQuery bool
}

func (r *nativeSemanticSearchRunner) QueryRows(query string) ([]lbugruntime.Row, error) {
	if strings.Contains(query, "QUERY_VECTOR_INDEX") {
		r.sawVectorQuery = true
		return r.vectorRows, nil
	}
	return r.native.QueryRows(query)
}
