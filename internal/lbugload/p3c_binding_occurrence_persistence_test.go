package lbugload

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/parser"
	"github.com/tamnguyendinh/anvien/internal/providers/tsjs"
	"github.com/tamnguyendinh/anvien/internal/resolution"
	"github.com/tamnguyendinh/anvien/internal/scanner"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
)

func TestP3CBindingOccurrencesPreserveGraphJSONAndLadybugLoadParity(t *testing.T) {
	ir := p3cParseTypeScript(t, "src/p3c.ts", []byte(`import { helper } from './dep';
function outer(rows: any) {
  let [value] = rows;
  value.map(() => 1);
  for ([value] of rows) {}
  function inner(rows: any) {
    const [value] = rows;
    value.map(() => 1);
  }
  try {} catch ([caught]) { caught.map(() => 1); }
  for (const [loop] of rows) { loop.map(() => 1); }
}
function parameter([arg]: any) { arg.map(() => 1); }
`))
	dep := p3cParseTypeScript(t, "src/dep.ts", []byte(`export function helper(): void {}`))

	result, err := resolution.Resolve([]scopeir.ScopeIR{ir, dep}, resolution.Options{})
	if err != nil {
		t.Fatalf("Resolve() error = %v", err)
	}
	wantRelationships := p3cBindingRelationships(t, result.Graph)
	if len(wantRelationships) != 6 {
		t.Fatalf("binding relationships before persistence = %d, want 5 reads + 1 write", len(wantRelationships))
	}

	raw, err := json.Marshal(result.Graph)
	if err != nil {
		t.Fatalf("marshal graph JSON: %v", err)
	}
	var persisted graph.Graph
	if err := json.Unmarshal(raw, &persisted); err != nil {
		t.Fatalf("unmarshal graph JSON: %v", err)
	}
	gotRelationships := p3cBindingRelationships(t, &persisted)
	if !reflect.DeepEqual(gotRelationships, wantRelationships) {
		t.Fatalf("binding relationship JSON parity drifted:\nwant=%#v\ngot=%#v", wantRelationships, gotRelationships)
	}

	wantBindingNodeIDs := make(map[string]struct{}, len(ir.BindingLeaves))
	for _, leaf := range ir.BindingLeaves {
		definition := p3cPersistenceDefinitionForLeaf(t, ir, leaf)
		nodeID := p3cPersistenceNodeID(t, result.Graph, definition)
		if _, duplicate := wantBindingNodeIDs[nodeID]; duplicate {
			t.Fatalf("binding leaf %q reused persisted node %q", definition.ID, nodeID)
		}
		wantBindingNodeIDs[nodeID] = struct{}{}
		if _, ok := persisted.GetNode(nodeID); !ok {
			t.Fatalf("binding leaf %q missing after graph JSON round-trip", definition.ID)
		}
		defines := 0
		for _, relationship := range persisted.Relationships {
			if relationship.Type == graph.RelDefines && relationship.TargetID == nodeID {
				defines++
			}
		}
		if defines != 1 {
			t.Fatalf("binding leaf %q persisted DEFINES count = %d, want exactly 1", definition.ID, defines)
		}
	}
	if len(wantBindingNodeIDs) != 7 {
		t.Fatalf("persisted binding occurrence conservation = %d, want exactly 7 accepted leaves", len(wantBindingNodeIDs))
	}

	export, err := ExportGraphCSVs(&persisted, t.TempDir())
	if err != nil {
		t.Fatalf("ExportGraphCSVs() error = %v", err)
	}
	if export.Metrics.SkippedNodes != 0 || export.Metrics.SkippedRelationships != 0 {
		t.Fatalf("Ladybug export skipped binding graph records: %#v", export.Metrics)
	}

	variableIDs := make(map[string]struct{})
	for _, nodeFile := range export.NodeFiles {
		if nodeFile.Table != string(scopeir.NodeVariable) {
			continue
		}
		rows := readCSV(t, nodeFile.CSVPath)
		for _, row := range rows[1:] {
			variableIDs[row[0]] = struct{}{}
		}
	}
	for nodeID := range wantBindingNodeIDs {
		if _, ok := variableIDs[nodeID]; !ok {
			t.Fatalf("Ladybug Variable CSV missing binding node %q", nodeID)
		}
	}

	csvRelationships, err := ReadRelationshipCSVRows(export.RelationshipCSVPath)
	if err != nil {
		t.Fatalf("ReadRelationshipCSVRows() error = %v", err)
	}
	seenCSV := make(map[string]struct{}, len(wantRelationships))
	for _, row := range csvRelationships {
		if row.Type != string(graph.RelAccesses) || row.TargetRole != "binding" || row.ProofKind != "scope-binding" {
			continue
		}
		want, ok := wantRelationships[row.SourceSiteID]
		if !ok {
			t.Fatalf("Ladybug CSV contains unexpected binding source site %q: %#v", row.SourceSiteID, row)
		}
		wantRow := p3cRelationshipCSVRow(t, want)
		if !reflect.DeepEqual(row, wantRow) {
			t.Fatalf("Ladybug relationship CSV parity drifted for %q:\nwant=%#v\ngot=%#v", row.SourceSiteID, wantRow, row)
		}
		seenCSV[row.SourceSiteID] = struct{}{}
	}
	if len(seenCSV) != len(wantRelationships) {
		t.Fatalf("Ladybug binding relationship rows = %d, want %d", len(seenCSV), len(wantRelationships))
	}

	runner := &recordingRunner{}
	loadResult, err := LoadCSVExport(runner, export)
	if err != nil {
		t.Fatalf("LoadCSVExport() error = %v", err)
	}
	if loadResult.NodeCopyCount != len(export.NodeFiles) ||
		loadResult.RelationshipCopyCount != len(export.RelationshipPairFiles) ||
		loadResult.FallbackInsertCount != 0 ||
		loadResult.FallbackInsertFailures != 0 ||
		loadResult.SkippedRelationships != 0 ||
		len(loadResult.Warnings) != 0 {
		t.Fatalf("Ladybug load parity drifted: %#v", loadResult)
	}
	if !strings.Contains(strings.Join(runner.queries, "\n"), "COPY CodeRelation FROM") {
		t.Fatalf("Ladybug load did not prepare CodeRelation COPY: %#v", runner.queries)
	}
}

func p3cParseTypeScript(t *testing.T, filePath string, source []byte) scopeir.ScopeIR {
	t.Helper()
	pool := parser.NewPool(nil, parser.PoolOptions{ParseTimeout: time.Second})
	defer pool.Close()
	parsed, err := pool.Parse(context.Background(), parser.Request{
		FilePath: filePath,
		Language: scanner.TypeScript,
		Source:   source,
	})
	if err != nil {
		t.Fatalf("parse %s: %v", filePath, err)
	}
	defer parsed.Close()
	ir, err := tsjs.Extract(tsjs.Request{
		FilePath: filePath,
		FileHash: "hash-" + filePath,
		Language: scanner.TypeScript,
		Source:   source,
		Root:     parsed.Tree.RootNode(),
	})
	if err != nil {
		t.Fatalf("extract %s: %v", filePath, err)
	}
	return ir
}

func p3cBindingRelationships(t *testing.T, g *graph.Graph) map[string]graph.Relationship {
	t.Helper()
	relationships := make(map[string]graph.Relationship)
	for _, relationship := range g.Relationships {
		if relationship.Type != graph.RelAccesses || relationship.TargetRole != "binding" || relationship.ProofKind != "scope-binding" {
			continue
		}
		if relationship.SourceSiteID == "" {
			t.Fatalf("binding relationship has no source-site identity: %#v", relationship)
		}
		if _, duplicate := relationships[relationship.SourceSiteID]; duplicate {
			t.Fatalf("duplicate binding relationship source site %q", relationship.SourceSiteID)
		}
		if _, ok := g.GetNode(relationship.SourceID); !ok {
			t.Fatalf("binding relationship missing source endpoint %q", relationship.SourceID)
		}
		if _, ok := g.GetNode(relationship.TargetID); !ok {
			t.Fatalf("binding relationship missing target endpoint %q", relationship.TargetID)
		}
		relationships[relationship.SourceSiteID] = relationship
	}
	return relationships
}

func p3cPersistenceDefinitionForLeaf(t *testing.T, ir scopeir.ScopeIR, leaf scopeir.BindingLeafFact) scopeir.DefinitionFact {
	t.Helper()
	matches := make([]scopeir.DefinitionFact, 0, 1)
	for _, definition := range ir.Definitions {
		if definition.Label != scopeir.NodeVariable ||
			definition.Name != leaf.Name ||
			definition.FilePath != leaf.FilePath ||
			definition.Range != leaf.Range ||
			!p3cPersistenceRangesEqual(definition.SelectionRange, leaf.SelectionRange) {
			continue
		}
		matches = append(matches, definition)
	}
	if len(matches) != 1 {
		t.Fatalf("binding leaf %#v matching definitions = %d, want exactly 1", leaf, len(matches))
	}
	return matches[0]
}

func p3cPersistenceRangesEqual(left *scopeir.Range, right *scopeir.Range) bool {
	if left == nil || right == nil {
		return left == nil && right == nil
	}
	return *left == *right
}

func p3cPersistenceNodeID(t *testing.T, g *graph.Graph, definition scopeir.DefinitionFact) string {
	t.Helper()
	matches := make([]graph.Node, 0, 1)
	for _, node := range g.Nodes {
		if node.Label == definition.Label &&
			node.Properties["filePath"] == definition.FilePath &&
			node.Properties["name"] == definition.Name &&
			node.Properties["startLine"] == definition.Range.StartLine &&
			node.Properties["startCol"] == definition.Range.StartCol &&
			node.Properties["endLine"] == definition.Range.EndLine &&
			node.Properties["endCol"] == definition.Range.EndCol {
			matches = append(matches, node)
		}
	}
	if len(matches) != 1 {
		t.Fatalf("binding definition %q graph nodes = %d, want exactly 1", definition.ID, len(matches))
	}
	return matches[0].ID
}

func p3cRelationshipCSVRow(t *testing.T, relationship graph.Relationship) RelationshipCSVRow {
	t.Helper()
	evidence, err := json.Marshal(relationship.Evidence)
	if err != nil {
		t.Fatalf("marshal relationship evidence: %v", err)
	}
	sourceSiteIDs, err := json.Marshal(relationship.SourceSiteIDs)
	if err != nil {
		t.Fatalf("marshal relationship source-site IDs: %v", err)
	}
	step := 0
	if relationship.Step != nil {
		step = *relationship.Step
	}
	return RelationshipCSVRow{
		FromID:           relationship.SourceID,
		ToID:             relationship.TargetID,
		Type:             string(relationship.Type),
		Confidence:       fmt.Sprint(relationship.Confidence),
		Reason:           relationship.Reason,
		Step:             fmt.Sprint(step),
		ResolutionSource: relationship.ResolutionSource,
		Evidence:         string(evidence),
		FileHash:         relationship.FileHash,
		SourceSiteID:     relationship.SourceSiteID,
		SourceSiteIDs:    string(sourceSiteIDs),
		SourceSiteCount:  fmt.Sprint(relationship.SourceSiteCount),
		SourceSiteStatus: relationship.SourceSiteStatus,
		ProofKind:        relationship.ProofKind,
		TargetRole:       relationship.TargetRole,
		TargetText:       relationship.TargetText,
		FilePath:         relationship.FilePath,
		StartLine:        fmt.Sprint(relationship.StartLine),
		StartCol:         fmt.Sprint(relationship.StartCol),
		EndLine:          fmt.Sprint(relationship.EndLine),
		EndCol:           fmt.Sprint(relationship.EndCol),
	}
}
