package lbugload

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/resolution"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
)

func TestP5DExportProofPreservesGraphJSONAndLadybugRelationshipParity(t *testing.T) {
	consumer := p3cParseTypeScript(t, "src/app.ts", []byte(`
import { barrelRun } from './barrel';
export function caller(): void {
  barrelRun();
}
`))
	barrel := p3cParseTypeScript(t, "src/barrel.ts", []byte(`
export { run as barrelRun } from './impl';
`))
	impl := p3cParseTypeScript(t, "src/impl.ts", []byte(`
export function run(): void {}
`))

	result, err := resolution.Resolve([]scopeir.ScopeIR{consumer, barrel, impl}, resolution.Options{})
	if err != nil {
		t.Fatalf("Resolve() error = %v", err)
	}
	before := p5dProofRelationship(t, result.Graph)
	if before.Type != graph.RelCalls || before.SourceSiteID == "" || len(before.Evidence) < 3 {
		t.Fatalf("P5-D relationship before persistence = %#v", before)
	}

	raw, err := json.Marshal(result.Graph)
	if err != nil {
		t.Fatalf("marshal Graph JSON: %v", err)
	}
	var snapshot graph.Graph
	if err := json.Unmarshal(raw, &snapshot); err != nil {
		t.Fatalf("unmarshal Graph JSON: %v", err)
	}
	afterJSON := p5dRelationshipByID(t, &snapshot, before.ID)
	if before.SourceID != afterJSON.SourceID ||
		before.TargetID != afterJSON.TargetID ||
		before.SourceSiteID != afterJSON.SourceSiteID ||
		!reflect.DeepEqual(before.SourceSiteIDs, afterJSON.SourceSiteIDs) ||
		!reflect.DeepEqual(before.Evidence, afterJSON.Evidence) {
		t.Fatalf("Graph JSON endpoint/Evidence parity drifted:\nbefore=%#v\nafter=%#v", before, afterJSON)
	}

	tempRoot := filepath.Join("..", "..", ".tmp", "p5d-lbugload-tests")
	if err := os.MkdirAll(tempRoot, 0o755); err != nil {
		t.Fatalf("create repo-local P5-D temp root: %v", err)
	}
	tempDir, err := os.MkdirTemp(tempRoot, "proof-")
	if err != nil {
		t.Fatalf("create repo-local P5-D temp dir: %v", err)
	}
	t.Cleanup(func() { _ = os.RemoveAll(tempDir) })

	exported, err := ExportGraphCSVs(&snapshot, filepath.Join(tempDir, "csv"))
	if err != nil {
		t.Fatalf("ExportGraphCSVs() error = %v", err)
	}
	if exported.Metrics.SkippedNodes != 0 || exported.Metrics.SkippedRelationships != 0 {
		t.Fatalf("Ladybug export skipped P5-D graph records: %#v", exported.Metrics)
	}
	rows, err := ReadRelationshipCSVRows(exported.RelationshipCSVPath)
	if err != nil {
		t.Fatalf("ReadRelationshipCSVRows() error = %v", err)
	}
	var persistedRow *RelationshipCSVRow
	for index := range rows {
		row := &rows[index]
		if row.FromID == before.SourceID &&
			row.ToID == before.TargetID &&
			row.Type == string(before.Type) &&
			row.SourceSiteID == before.SourceSiteID {
			persistedRow = row
			break
		}
	}
	if persistedRow == nil {
		t.Fatalf("Ladybug relationship CSV missing endpoint/source site for %#v", before)
	}
	var persistedEvidence []graph.Evidence
	if err := json.Unmarshal([]byte(persistedRow.Evidence), &persistedEvidence); err != nil {
		t.Fatalf("decode Ladybug Evidence JSON: %v", err)
	}
	if !reflect.DeepEqual(persistedEvidence, before.Evidence) {
		t.Fatalf("Ladybug Evidence parity drifted:\nwant=%#v\ngot=%#v", before.Evidence, persistedEvidence)
	}
	if persistedRow.ProofKind != before.ProofKind ||
		persistedRow.TargetRole != before.TargetRole ||
		persistedRow.TargetText != before.TargetText {
		t.Fatalf("Ladybug relationship semantic columns drifted: %#v", persistedRow)
	}

	runner := &recordingRunner{}
	loaded, err := LoadCSVExport(runner, exported)
	if err != nil {
		t.Fatalf("LoadCSVExport() error = %v", err)
	}
	if loaded.FallbackInsertCount != 0 ||
		loaded.FallbackInsertFailures != 0 ||
		loaded.SkippedRelationships != 0 ||
		len(loaded.Warnings) != 0 {
		t.Fatalf("Ladybug load parity = %#v", loaded)
	}
	if !strings.Contains(strings.Join(runner.queries, "\n"), "COPY CodeRelation FROM") {
		t.Fatalf("Ladybug loader did not retain relationship COPY path: %#v", runner.queries)
	}
}

func p5dProofRelationship(t *testing.T, g *graph.Graph) graph.Relationship {
	t.Helper()
	var matches []graph.Relationship
	for _, relationship := range g.Relationships {
		hasTerminal := false
		for _, evidence := range relationship.Evidence {
			if evidence.Kind == "export-terminal-v1" {
				hasTerminal = true
				break
			}
		}
		if relationship.Type == graph.RelCalls && hasTerminal {
			matches = append(matches, relationship)
		}
	}
	if len(matches) != 1 {
		t.Fatalf("P5-D proof relationships = %d, want exactly 1: %#v", len(matches), matches)
	}
	return matches[0]
}

func p5dRelationshipByID(t *testing.T, g *graph.Graph, relationshipID string) graph.Relationship {
	t.Helper()
	for _, relationship := range g.Relationships {
		if relationship.ID == relationshipID {
			return relationship
		}
	}
	t.Fatalf("relationship %q missing after Graph JSON round-trip", relationshipID)
	return graph.Relationship{}
}
