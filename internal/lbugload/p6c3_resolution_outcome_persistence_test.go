package lbugload

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/resolution"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
	"github.com/tamnguyendinh/anvien/internal/semantic"
)

func TestP6C3ResolutionOutcomesPreserveGraphJSONAndLadybugParity(t *testing.T) {
	ir := p6c3PersistenceIR()
	result, err := resolution.Resolve([]scopeir.ScopeIR{ir}, resolution.Options{})
	if err != nil {
		t.Fatalf("Resolve() error = %v", err)
	}
	resolved := requireP6C3Outcome(t, result.ResolutionOutcomes, "helper", resolution.ResolutionResolvedInternal)
	unresolved := requireP6C3Outcome(t, result.ResolutionOutcomes, "missing", resolution.ResolutionUnresolved)
	resolvedRelationship := requireP6C3Relationship(t, result.Graph, resolved.SourceSiteID)

	if _, err := semantic.Apply(result.Graph); err != nil {
		t.Fatalf("semantic.Apply() error = %v", err)
	}
	gapNode := requireP6C3ResolutionGap(t, result.Graph, unresolved)

	raw, err := json.Marshal(result.Graph)
	if err != nil {
		t.Fatalf("marshal Graph JSON: %v", err)
	}
	var snapshot graph.Graph
	if err := json.Unmarshal(raw, &snapshot); err != nil {
		t.Fatalf("unmarshal Graph JSON: %v", err)
	}
	afterJSON := requireP6C3RelationshipByID(t, &snapshot, resolvedRelationship.ID)
	if !reflect.DeepEqual(resolvedRelationship.Evidence, afterJSON.Evidence) {
		t.Fatalf("resolved outcome Graph JSON evidence drift:\nwant=%#v\ngot=%#v", resolvedRelationship.Evidence, afterJSON.Evidence)
	}
	afterGap, ok := snapshot.GetNode(gapNode.ID)
	if !ok || afterGap.Properties["note"] != gapNode.Properties["note"] {
		t.Fatalf("unresolved outcome Graph JSON note drift: want=%#v got=%#v", gapNode, afterGap)
	}

	tempRoot := filepath.Join("..", "..", ".tmp", "p6c3-lbugload-tests")
	if err := os.MkdirAll(tempRoot, 0o755); err != nil {
		t.Fatalf("create repo-local P6-C3 temp root: %v", err)
	}
	tempDir, err := os.MkdirTemp(tempRoot, "outcome-")
	if err != nil {
		t.Fatalf("create repo-local P6-C3 temp dir: %v", err)
	}
	t.Cleanup(func() { _ = os.RemoveAll(tempDir) })

	exported, err := ExportGraphCSVs(&snapshot, filepath.Join(tempDir, "csv"))
	if err != nil {
		t.Fatalf("ExportGraphCSVs() error = %v", err)
	}
	if exported.Metrics.SkippedNodes != 0 || exported.Metrics.SkippedRelationships != 0 {
		t.Fatalf("Ladybug export skipped outcome carriers: %#v", exported.Metrics)
	}
	relationshipRows, err := ReadRelationshipCSVRows(exported.RelationshipCSVPath)
	if err != nil {
		t.Fatalf("ReadRelationshipCSVRows() error = %v", err)
	}
	foundResolved := false
	for _, row := range relationshipRows {
		if row.SourceSiteID != resolved.SourceSiteID {
			continue
		}
		var evidence []graph.Evidence
		if err := json.Unmarshal([]byte(row.Evidence), &evidence); err != nil {
			t.Fatalf("decode persisted resolved evidence: %v", err)
		}
		persisted := requireP6C3EvidenceOutcome(t, evidence, resolved.SourceSiteID)
		if !reflect.DeepEqual(persisted, resolved) {
			t.Fatalf("Ladybug resolved outcome drift: want=%#v got=%#v", resolved, persisted)
		}
		foundResolved = true
	}
	if !foundResolved {
		t.Fatalf("Ladybug relationship CSV missing resolved source site %q", resolved.SourceSiteID)
	}

	gapRows := readP6C3CSV(t, filepath.Join(exported.CSVDir, "resolutiongap.csv"))
	header := make(map[string]int, len(gapRows[0]))
	for index, column := range gapRows[0] {
		header[column] = index
	}
	foundGap := false
	for _, row := range gapRows[1:] {
		if row[header["sourceSiteId"]] != unresolved.SourceSiteID {
			continue
		}
		var persisted resolution.ResolutionOutcome
		if err := json.Unmarshal([]byte(row[header["note"]]), &persisted); err != nil {
			t.Fatalf("decode persisted ResolutionGap outcome: %v", err)
		}
		if !reflect.DeepEqual(persisted, unresolved) || row[header["sourceSiteStatus"]] != "unresolved_local_binding" {
			t.Fatalf("Ladybug ResolutionGap outcome drift: row=%#v want=%#v got=%#v", row, unresolved, persisted)
		}
		foundGap = true
	}
	if !foundGap {
		t.Fatalf("Ladybug ResolutionGap CSV missing source site %q", unresolved.SourceSiteID)
	}

	runner := &p6c3RecordingRunner{}
	loaded, err := LoadCSVExport(runner, exported)
	if err != nil {
		t.Fatalf("LoadCSVExport() error = %v", err)
	}
	if loaded.FallbackInsertCount != 0 || loaded.FallbackInsertFailures != 0 ||
		loaded.SkippedRelationships != 0 || len(loaded.Warnings) != 0 {
		t.Fatalf("Ladybug outcome load parity = %#v", loaded)
	}
}

type p6c3RecordingRunner struct {
	queries []string
}

func (runner *p6c3RecordingRunner) Query(query string) error {
	runner.queries = append(runner.queries, query)
	return nil
}

func p6c3PersistenceIR() scopeir.ScopeIR {
	const (
		filePath      = "src/app.ts"
		fileHash      = "hash-app"
		moduleScope   = "scope:src/app.ts:module"
		functionScope = "scope:src/app.ts:caller"
	)
	moduleParent := moduleScope
	helper := scopeir.DefinitionFact{
		ID: "def:helper", FilePath: filePath, FileHash: fileHash, Name: "helper", QualifiedName: "helper",
		Label: scopeir.NodeFunction, Range: scopeir.Range{StartLine: 1, EndLine: 1},
	}
	caller := scopeir.DefinitionFact{
		ID: "def:caller", FilePath: filePath, FileHash: fileHash, Name: "caller", QualifiedName: "caller",
		Label: scopeir.NodeFunction, Range: scopeir.Range{StartLine: 2, EndLine: 6},
	}
	return scopeir.ScopeIR{
		FilePath: filePath, FileHash: fileHash, Language: "typescript", ModuleScope: moduleScope,
		Scopes: []scopeir.ScopeFact{
			{
				ID: moduleScope, Kind: scopeir.ScopeModule, FilePath: filePath, FileHash: fileHash,
				Range: scopeir.Range{StartLine: 1, EndLine: 6}, OwnedDefIDs: []string{helper.ID},
				Bindings: []scopeir.BindingFact{{Name: helper.Name, DefID: helper.ID, Origin: scopeir.BindingLocal}},
			},
			{
				ID: functionScope, Parent: &moduleParent, Kind: scopeir.ScopeFunction, FilePath: filePath, FileHash: fileHash,
				Range: scopeir.Range{StartLine: 2, EndLine: 6}, OwnedDefIDs: []string{caller.ID},
				Bindings: []scopeir.BindingFact{{Name: caller.Name, DefID: caller.ID, Origin: scopeir.BindingLocal}},
			},
		},
		Definitions: []scopeir.DefinitionFact{helper, caller},
		Calls: []scopeir.CallSiteFact{
			{FilePath: filePath, FileHash: fileHash, Name: "helper", Range: scopeir.Range{StartLine: 3, StartCol: 2, EndLine: 3, EndCol: 10}, InScope: functionScope, CallForm: scopeir.CallFree},
			{FilePath: filePath, FileHash: fileHash, Name: "missing", Range: scopeir.Range{StartLine: 4, StartCol: 2, EndLine: 4, EndCol: 11}, InScope: functionScope, CallForm: scopeir.CallFree},
		},
	}
}

func readP6C3CSV(t *testing.T, path string) [][]string {
	t.Helper()
	file, err := os.Open(path)
	if err != nil {
		t.Fatalf("open csv %s: %v", path, err)
	}
	defer file.Close()
	rows, err := csv.NewReader(file).ReadAll()
	if err != nil {
		t.Fatalf("read csv %s: %v", path, err)
	}
	return rows
}

func requireP6C3Outcome(t *testing.T, outcomes []resolution.ResolutionOutcome, requestedName string, status resolution.ResolutionStatus) resolution.ResolutionOutcome {
	t.Helper()
	var matches []resolution.ResolutionOutcome
	for _, outcome := range outcomes {
		if outcome.RequestedName == requestedName && outcome.Status == status {
			matches = append(matches, outcome)
		}
	}
	if len(matches) != 1 {
		t.Fatalf("outcomes for %s/%s = %d, want one: %#v", requestedName, status, len(matches), matches)
	}
	return matches[0]
}

func requireP6C3Relationship(t *testing.T, g *graph.Graph, sourceSiteID string) graph.Relationship {
	t.Helper()
	var matches []graph.Relationship
	for _, relationship := range g.Relationships {
		for _, evidence := range relationship.Evidence {
			if evidence.Kind != resolution.ResolutionOutcomeEvidenceKind {
				continue
			}
			var outcome resolution.ResolutionOutcome
			if json.Unmarshal([]byte(evidence.Note), &outcome) == nil && outcome.SourceSiteID == sourceSiteID {
				matches = append(matches, relationship)
			}
		}
	}
	if len(matches) != 1 {
		t.Fatalf("relationships for source site %q = %d, want one: %#v", sourceSiteID, len(matches), matches)
	}
	return matches[0]
}

func requireP6C3ResolutionGap(t *testing.T, g *graph.Graph, want resolution.ResolutionOutcome) graph.Node {
	t.Helper()
	var matches []graph.Node
	for _, node := range g.Nodes {
		if node.Label != scopeir.NodeResolutionGap {
			continue
		}
		note, _ := node.Properties["note"].(string)
		var outcome resolution.ResolutionOutcome
		if json.Unmarshal([]byte(note), &outcome) == nil && outcome.SourceSiteID == want.SourceSiteID {
			if !reflect.DeepEqual(outcome, want) {
				t.Fatalf("ResolutionGap outcome drift: want=%#v got=%#v", want, outcome)
			}
			matches = append(matches, node)
		}
	}
	if len(matches) != 1 {
		t.Fatalf("ResolutionGap nodes for source site %q = %d, want one: %#v", want.SourceSiteID, len(matches), matches)
	}
	return matches[0]
}

func requireP6C3RelationshipByID(t *testing.T, g *graph.Graph, relationshipID string) graph.Relationship {
	t.Helper()
	for _, relationship := range g.Relationships {
		if relationship.ID == relationshipID {
			return relationship
		}
	}
	t.Fatalf("relationship %q missing", relationshipID)
	return graph.Relationship{}
}

func requireP6C3EvidenceOutcome(t *testing.T, evidence []graph.Evidence, sourceSiteID string) resolution.ResolutionOutcome {
	t.Helper()
	var matches []resolution.ResolutionOutcome
	for _, item := range evidence {
		if item.Kind != resolution.ResolutionOutcomeEvidenceKind {
			continue
		}
		var outcome resolution.ResolutionOutcome
		if err := json.Unmarshal([]byte(item.Note), &outcome); err != nil {
			t.Fatalf("decode persisted outcome evidence: %v", err)
		}
		if outcome.SourceSiteID == sourceSiteID {
			matches = append(matches, outcome)
		}
	}
	if len(matches) != 1 {
		t.Fatalf("outcome evidence for source site %q = %d, want one: %#v", sourceSiteID, len(matches), matches)
	}
	return matches[0]
}
