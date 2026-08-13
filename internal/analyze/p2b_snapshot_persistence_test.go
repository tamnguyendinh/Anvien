package analyze

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
)

func TestGraphSnapshotPreservesCorrectedDefinitionFactsAndDefinesEndpoints(t *testing.T) {
	g := graph.New()
	g.AddNode(graph.Node{ID: "File:src/app.ts", Label: scopeir.NodeFile, Properties: graph.NodeProperties{
		"name": "app.ts", "filePath": "src/app.ts",
	}})
	g.AddNode(graph.Node{ID: "Function:opaque-worker", Label: scopeir.NodeFunction, Properties: graph.NodeProperties{
		"name": "worker", "qualifiedName": "api.worker", "filePath": "src/app.ts",
		"startLine": 2, "startCol": 0, "endLine": 4, "endCol": 1,
		"selectionStartLine": 2, "selectionStartCol": 16, "selectionEndLine": 2, "selectionEndCol": 22,
	}})
	g.AddNode(graph.Node{ID: "Variable:opaque-config", Label: scopeir.NodeVariable, Properties: graph.NodeProperties{
		"name": "config", "qualifiedName": "config", "filePath": "src/app.ts",
		"startLine": 6, "startCol": 0, "endLine": 6, "endCol": 24,
	}})
	g.AddRelationship(graph.Relationship{ID: "rel:file-worker", SourceID: "File:src/app.ts", TargetID: "Function:opaque-worker", Type: graph.RelDefines, Confidence: 1})
	g.AddRelationship(graph.Relationship{ID: "rel:file-config", SourceID: "File:src/app.ts", TargetID: "Variable:opaque-config", Type: graph.RelDefines, Confidence: 1})

	path := filepath.Join(t.TempDir(), "graph.json")
	if err := writeGraphSnapshot(path, g); err != nil {
		t.Fatalf("writeGraphSnapshot() error = %v", err)
	}
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile() error = %v", err)
	}
	var snapshot graph.Graph
	if err := json.Unmarshal(raw, &snapshot); err != nil {
		t.Fatalf("Unmarshal() error = %v", err)
	}
	if len(snapshot.Nodes) != len(g.Nodes) || len(snapshot.Relationships) != len(g.Relationships) {
		t.Fatalf("snapshot records = %d/%d, want %d/%d", len(snapshot.Nodes), len(snapshot.Relationships), len(g.Nodes), len(g.Relationships))
	}

	worker := snapshotNode(t, &snapshot, "Function:opaque-worker")
	for key, want := range map[string]float64{
		"startLine": 2, "startCol": 0, "endLine": 4, "endCol": 1,
		"selectionStartLine": 2, "selectionStartCol": 16, "selectionEndLine": 2, "selectionEndCol": 22,
	} {
		if got := worker.Properties[key]; got != want {
			t.Fatalf("snapshot worker %s = %#v, want %v", key, got, want)
		}
	}
	if worker.Properties["qualifiedName"] != "api.worker" {
		t.Fatalf("snapshot worker qualifiedName = %#v", worker.Properties["qualifiedName"])
	}
	config := snapshotNode(t, &snapshot, "Variable:opaque-config")
	for _, key := range []string{"selectionStartLine", "selectionStartCol", "selectionEndLine", "selectionEndCol"} {
		if _, exists := config.Properties[key]; exists {
			t.Fatalf("snapshot config has unexpected selection property %q: %#v", key, config.Properties)
		}
	}

	wantEndpoints := map[string]string{"rel:file-worker": "Function:opaque-worker", "rel:file-config": "Variable:opaque-config"}
	for _, relationship := range snapshot.Relationships {
		wantTarget, ok := wantEndpoints[relationship.ID]
		if !ok || relationship.SourceID != "File:src/app.ts" || relationship.TargetID != wantTarget || relationship.Type != graph.RelDefines {
			t.Fatalf("snapshot relationship drifted: %#v", relationship)
		}
		if _, ok := snapshot.GetNode(relationship.SourceID); !ok {
			t.Fatalf("snapshot relationship %q missing source %q", relationship.ID, relationship.SourceID)
		}
		if _, ok := snapshot.GetNode(relationship.TargetID); !ok {
			t.Fatalf("snapshot relationship %q missing target %q", relationship.ID, relationship.TargetID)
		}
	}
}

func snapshotNode(t *testing.T, snapshot *graph.Graph, id string) graph.Node {
	t.Helper()
	node, ok := snapshot.GetNode(id)
	if !ok {
		t.Fatalf("snapshot missing node %q", id)
	}
	return node
}
