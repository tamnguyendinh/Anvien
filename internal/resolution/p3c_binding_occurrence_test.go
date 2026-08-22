//go:build resolution_parser_integration

package resolution

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
)

func TestP3CBindingOccurrencesProjectAndResolveLexically(t *testing.T) {
	source := []byte(`import { helper } from './dep';
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
`)
	ir := parseTypeScriptSource(t, "src/p3c.ts", source)
	dep := parseTypeScriptSource(t, "src/dep.ts", []byte(`export function helper(): void {}`))

	if len(ir.Imports) != 1 {
		t.Fatalf("import facts = %d, want exactly 1", len(ir.Imports))
	}
	valueDefinitions := p3cDefinitionsNamed(ir, "value")
	if len(valueDefinitions) != 2 {
		t.Fatalf("value definitions = %d, want outer/inner exactly 2: %#v", len(valueDefinitions), valueDefinitions)
	}
	assignmentWrites := 0
	for _, access := range ir.Accesses {
		if access.ExplicitReceiver == "" && access.Name == "value" && access.Kind == scopeir.AccessWrite {
			assignmentWrites++
		}
	}
	if assignmentWrites != 1 {
		t.Fatalf("assignment-form loop writes = %d, want exactly 1", assignmentWrites)
	}

	result, err := Resolve([]scopeir.ScopeIR{ir, dep}, Options{})
	if err != nil {
		t.Fatalf("Resolve() error = %v", err)
	}

	leafDefinitions := make(map[string]scopeir.DefinitionFact, len(ir.BindingLeaves))
	for _, leaf := range ir.BindingLeaves {
		definition := p3cDefinitionForLeaf(t, ir, leaf)
		if _, duplicate := leafDefinitions[definition.ID]; duplicate {
			t.Fatalf("accepted binding leaf projected duplicate definition %q", definition.ID)
		}
		leafDefinitions[definition.ID] = definition
		nodeID := p3cRequireDefinitionNodeID(t, result.Graph, definition)
		defines := 0
		for _, relationship := range result.Graph.Relationships {
			if relationship.Type == graph.RelDefines && relationship.TargetID == nodeID {
				defines++
			}
		}
		if defines != 1 {
			t.Fatalf("binding leaf %q graph DEFINES occurrences = %d, want exactly 1", definition.ID, defines)
		}
	}
	if len(leafDefinitions) != len(ir.BindingLeaves) {
		t.Fatalf("binding occurrence conservation = %d/%d", len(leafDefinitions), len(ir.BindingLeaves))
	}

	bindingReferences := 0
	memberReceiverReferences := 0
	for _, call := range ir.Calls {
		if call.CallForm != scopeir.CallMember || call.ExplicitReceiver == "" {
			continue
		}
		expectedDefID := p3cLexicalBindingDefID(t, ir, call.InScope, call.ExplicitReceiver)
		definition := p3cDefinitionByID(t, ir, expectedDefID)
		reference := p3cRequireReference(t, result.ReferenceIndex.ByTargetDef[expectedDefID], call.InScope, call.ExplicitReceiver, call.Range, ReferenceRead)
		p3cRequireBindingRelationship(t, result.Graph, p3cRequireDefinitionNodeID(t, result.Graph, definition), reference)
		bindingReferences++
		memberReceiverReferences++
	}
	if memberReceiverReferences != 5 {
		t.Fatalf("member receiver references = %d, want 5", memberReceiverReferences)
	}

	for _, access := range ir.Accesses {
		if access.ExplicitReceiver != "" || access.Kind != scopeir.AccessWrite {
			continue
		}
		expectedDefID := p3cLexicalBindingDefID(t, ir, access.InScope, access.Name)
		definition := p3cDefinitionByID(t, ir, expectedDefID)
		reference := p3cRequireReference(t, result.ReferenceIndex.ByTargetDef[expectedDefID], access.InScope, access.Name, access.Range, ReferenceWrite)
		p3cRequireBindingRelationship(t, result.Graph, p3cRequireDefinitionNodeID(t, result.Graph, definition), reference)
		bindingReferences++
	}
	if bindingReferences != 6 {
		t.Fatalf("resolved binding references = %d, want 5 reads + 1 write", bindingReferences)
	}

	for _, relationship := range result.Graph.Relationships {
		if _, ok := result.Graph.GetNode(relationship.SourceID); !ok {
			t.Fatalf("relationship %q missing source endpoint %q", relationship.ID, relationship.SourceID)
		}
		if _, ok := result.Graph.GetNode(relationship.TargetID); !ok {
			t.Fatalf("relationship %q missing target endpoint %q", relationship.ID, relationship.TargetID)
		}
	}
	if result.Metrics.UnresolvedReferences != 0 || result.Metrics.UnresolvedReferenceDiagnostics != 0 || result.Metrics.UnattributedUnresolvedReferences != 0 {
		t.Fatalf("binding-caused resolution gaps remain: %#v", result.Metrics)
	}
	if result.Metrics.ResolvedAccesses != 6 || result.Metrics.ResolvedReferences != 6 {
		t.Fatalf("binding reference metrics = accesses:%d references:%d, want 6/6", result.Metrics.ResolvedAccesses, result.Metrics.ResolvedReferences)
	}
	if result.Metrics.ImportsResolved != 1 || result.Metrics.ImportUsesEmitted != 1 || result.Metrics.FinalizedImportsEmitted != 1 {
		t.Fatalf("import pipeline delta: %#v", result.Metrics)
	}
}

func TestP3CBindingOccurrenceProjectionFailsClosedOnOrphanOrDrift(t *testing.T) {
	source := []byte(`function broken(rows: any) { const [value] = rows; value.map(() => 1); }`)
	base := parseTypeScriptSource(t, "src/p3c-broken.ts", source)
	var leaf scopeir.BindingLeafFact
	for _, candidate := range base.BindingLeaves {
		if candidate.Name == "value" {
			leaf = candidate
			break
		}
	}
	if leaf.Name == "" {
		t.Fatal("missing value binding leaf in failure fixture")
	}
	definition := p3cDefinitionForLeaf(t, base, leaf)

	tests := []struct {
		name      string
		mutate    func(scopeir.ScopeIR) scopeir.ScopeIR
		wantError string
	}{
		{
			name: "missing definition",
			mutate: func(ir scopeir.ScopeIR) scopeir.ScopeIR {
				definitions := ir.Definitions[:0]
				for _, candidate := range ir.Definitions {
					if candidate.ID != definition.ID {
						definitions = append(definitions, candidate)
					}
				}
				ir.Definitions = definitions
				return ir
			},
			wantError: "has 0 matching variable definitions, want exactly 1",
		},
		{
			name: "missing local binding",
			mutate: func(ir scopeir.ScopeIR) scopeir.ScopeIR {
				for scopeIndex := range ir.Scopes {
					bindings := ir.Scopes[scopeIndex].Bindings[:0]
					for _, binding := range ir.Scopes[scopeIndex].Bindings {
						if binding.DefID != definition.ID {
							bindings = append(bindings, binding)
						}
					}
					ir.Scopes[scopeIndex].Bindings = bindings
				}
				return ir
			},
			wantError: "has 0 matching local bindings",
		},
		{
			name: "duplicate accepted leaf",
			mutate: func(ir scopeir.ScopeIR) scopeir.ScopeIR {
				ir.BindingLeaves = append(ir.BindingLeaves, leaf)
				return ir
			},
			wantError: "is projected more than once",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ir := test.mutate(base.Normalized())
			result, err := Resolve([]scopeir.ScopeIR{ir}, Options{})
			if err == nil || !strings.Contains(err.Error(), test.wantError) {
				t.Fatalf("Resolve() error = %v, want containing %q", err, test.wantError)
			}
			if result.Graph != nil {
				t.Fatalf("fail-closed projection returned a graph: %#v", result.Graph)
			}
		})
	}

	t.Run("coordinated wrong-file owner and local binding", func(t *testing.T) {
		drifted := base.Normalized()
		for scopeIndex := range drifted.Scopes {
			ownedDefIDs := make([]string, 0, len(drifted.Scopes[scopeIndex].OwnedDefIDs))
			for _, defID := range drifted.Scopes[scopeIndex].OwnedDefIDs {
				if defID != definition.ID {
					ownedDefIDs = append(ownedDefIDs, defID)
				}
			}
			drifted.Scopes[scopeIndex].OwnedDefIDs = ownedDefIDs

			bindings := make([]scopeir.BindingFact, 0, len(drifted.Scopes[scopeIndex].Bindings))
			for _, binding := range drifted.Scopes[scopeIndex].Bindings {
				if binding.DefID != definition.ID {
					bindings = append(bindings, binding)
				}
			}
			drifted.Scopes[scopeIndex].Bindings = bindings
		}

		wrongOwner := parseTypeScriptSource(t, "src/p3c-wrong-owner.ts", []byte(`export {};`)).Normalized()
		wrongOwnerFound := false
		for scopeIndex := range wrongOwner.Scopes {
			if wrongOwner.Scopes[scopeIndex].ID != wrongOwner.ModuleScope {
				continue
			}
			wrongOwner.Scopes[scopeIndex].OwnedDefIDs = append(wrongOwner.Scopes[scopeIndex].OwnedDefIDs, definition.ID)
			wrongOwner.Scopes[scopeIndex].Bindings = append(wrongOwner.Scopes[scopeIndex].Bindings, scopeir.BindingFact{
				Name:   definition.Name,
				DefID:  definition.ID,
				Origin: scopeir.BindingLocal,
			})
			wrongOwnerFound = true
			break
		}
		if !wrongOwnerFound {
			t.Fatal("missing module scope in wrong-owner fixture")
		}

		sentinel := graph.New()
		sentinel.Metadata = map[string]any{"sentinel": "p3c-owner-drift"}
		sentinel.AddNode(graph.Node{
			ID:    "sentinel:p3c-owner-drift",
			Label: scopeir.NodeFile,
			Properties: graph.NodeProperties{
				"filePath": "sentinel/p3c-owner-drift.ts",
				"name":     "sentinel",
			},
		})
		beforeBytes, err := json.Marshal(sentinel)
		if err != nil {
			t.Fatalf("marshal sentinel graph before resolution: %v", err)
		}
		beforeNodes := len(sentinel.Nodes)
		beforeRelationships := len(sentinel.Relationships)

		result, err := ResolveInto(sentinel, []scopeir.ScopeIR{drifted, wrongOwner}, Options{})
		if err == nil {
			t.Error("ResolveInto() error = nil, want coordinated wrong-owner drift rejection")
		} else if !strings.Contains(err.Error(), "has file drift") {
			t.Errorf("ResolveInto() error = %v, want lexical-owner file drift", err)
		}
		if result.Graph != nil {
			t.Errorf("fail-closed owner-drift projection returned a graph: %#v", result.Graph)
		}

		afterBytes, marshalErr := json.Marshal(sentinel)
		if marshalErr != nil {
			t.Fatalf("marshal sentinel graph after resolution: %v", marshalErr)
		}
		if len(sentinel.Nodes) != beforeNodes || len(sentinel.Relationships) != beforeRelationships {
			t.Errorf(
				"owner-drift resolution mutated sentinel graph counts: nodes %d -> %d, relationships %d -> %d",
				beforeNodes,
				len(sentinel.Nodes),
				beforeRelationships,
				len(sentinel.Relationships),
			)
		}
		if !bytes.Equal(afterBytes, beforeBytes) {
			t.Errorf("owner-drift resolution mutated sentinel graph bytes:\nbefore: %s\nafter:  %s", beforeBytes, afterBytes)
		}
	})
}

func p3cDefinitionsNamed(ir scopeir.ScopeIR, name string) []scopeir.DefinitionFact {
	definitions := make([]scopeir.DefinitionFact, 0)
	for _, definition := range ir.Definitions {
		if definition.Label == scopeir.NodeVariable && definition.Name == name {
			definitions = append(definitions, definition)
		}
	}
	return definitions
}

func p3cDefinitionForLeaf(t *testing.T, ir scopeir.ScopeIR, leaf scopeir.BindingLeafFact) scopeir.DefinitionFact {
	t.Helper()
	matches := make([]scopeir.DefinitionFact, 0, 1)
	for _, definition := range ir.Definitions {
		if definition.Label != scopeir.NodeVariable ||
			definition.Name != leaf.Name ||
			cleanPath(definition.FilePath) != cleanPath(leaf.FilePath) ||
			definition.Range != leaf.Range ||
			!p3cSelectionRangesEqual(definition.SelectionRange, leaf.SelectionRange) {
			continue
		}
		matches = append(matches, definition)
	}
	if len(matches) != 1 {
		t.Fatalf("binding leaf %#v matching definitions = %d, want exactly 1: %#v", leaf, len(matches), matches)
	}
	return matches[0]
}

func p3cSelectionRangesEqual(left *scopeir.Range, right *scopeir.Range) bool {
	if left == nil || right == nil {
		return left == nil && right == nil
	}
	return *left == *right
}

func p3cDefinitionByID(t *testing.T, ir scopeir.ScopeIR, defID string) scopeir.DefinitionFact {
	t.Helper()
	for _, definition := range ir.Definitions {
		if definition.ID == defID {
			return definition
		}
	}
	t.Fatalf("missing definition %q", defID)
	return scopeir.DefinitionFact{}
}

func p3cRequireDefinitionNodeID(t *testing.T, g *graph.Graph, definition scopeir.DefinitionFact) string {
	t.Helper()
	matches := make([]graph.Node, 0, 1)
	for _, node := range g.Nodes {
		if node.Label != definition.Label ||
			node.Properties["filePath"] != cleanPath(definition.FilePath) ||
			node.Properties["name"] != definition.Name ||
			node.Properties["startLine"] != definition.Range.StartLine ||
			node.Properties["startCol"] != definition.Range.StartCol ||
			node.Properties["endLine"] != definition.Range.EndLine ||
			node.Properties["endCol"] != definition.Range.EndCol {
			continue
		}
		if selection := definition.SelectionRange; selection != nil &&
			(node.Properties["selectionStartLine"] != selection.StartLine ||
				node.Properties["selectionStartCol"] != selection.StartCol ||
				node.Properties["selectionEndLine"] != selection.EndLine ||
				node.Properties["selectionEndCol"] != selection.EndCol) {
			continue
		}
		matches = append(matches, node)
	}
	if len(matches) != 1 {
		t.Fatalf("definition occurrence %q at %#v graph nodes = %d, want exactly 1: %#v", definition.ID, definition.Range, len(matches), matches)
	}
	return matches[0].ID
}

func p3cLexicalBindingDefID(t *testing.T, ir scopeir.ScopeIR, startScope string, name string) string {
	t.Helper()
	scopes := make(map[string]scopeir.ScopeFact, len(ir.Scopes))
	for _, scope := range ir.Scopes {
		scopes[scope.ID] = scope
	}
	for scopeID := startScope; scopeID != ""; {
		scope, ok := scopes[scopeID]
		if !ok {
			t.Fatalf("missing scope %q while resolving %q", scopeID, name)
		}
		matches := make([]scopeir.BindingFact, 0, 1)
		for _, binding := range scope.Bindings {
			if binding.Name == name && binding.Origin == scopeir.BindingLocal {
				matches = append(matches, binding)
			}
		}
		if len(matches) == 1 {
			return matches[0].DefID
		}
		if len(matches) > 1 {
			t.Fatalf("ambiguous lexical binding %q in scope %q: %#v", name, scopeID, matches)
		}
		if scope.Parent == nil {
			break
		}
		scopeID = *scope.Parent
	}
	t.Fatalf("missing lexical binding %q from scope %q", name, startScope)
	return ""
}

func p3cRequireReference(t *testing.T, references []Reference, fromScope string, targetText string, factRange scopeir.Range, kind ReferenceKind) Reference {
	t.Helper()
	matches := make([]Reference, 0, 1)
	for _, reference := range references {
		if reference.FromScope == fromScope && reference.TargetText == targetText && reference.Range == factRange && reference.Kind == kind {
			matches = append(matches, reference)
		}
	}
	if len(matches) != 1 {
		t.Fatalf("references for %s %q in %q at %#v = %d, want exactly 1: %#v", kind, targetText, fromScope, factRange, len(matches), references)
	}
	return matches[0]
}

func p3cRequireBindingRelationship(t *testing.T, g *graph.Graph, targetNodeID string, reference Reference) {
	t.Helper()
	matches := make([]graph.Relationship, 0, 1)
	candidates := make([]graph.Relationship, 0, 1)
	for _, relationship := range g.Relationships {
		if relationship.Type != graph.RelAccesses || relationship.TargetID != targetNodeID {
			continue
		}
		candidates = append(candidates, relationship)
		if relationship.SourceSiteID == reference.SourceSiteID {
			matches = append(matches, relationship)
		}
	}
	if len(matches) != 1 {
		t.Fatalf("binding relationship target=%q sourceSite=%q count=%d, want exactly 1; candidates=%#v", targetNodeID, reference.SourceSiteID, len(matches), candidates)
	}
	relationship := matches[0]
	wantStep := 1
	if reference.Kind == ReferenceWrite {
		wantStep = 2
	}
	if relationship.TargetRole != bindingOccurrenceTargetRole ||
		relationship.ProofKind != proofKindScopeBinding ||
		relationship.SourceSiteStatus != sourceSiteStatusResolved ||
		relationship.TargetText != reference.TargetText ||
		relationship.Step == nil || *relationship.Step != wantStep {
		t.Fatalf("binding relationship metadata drifted: %#v", relationship)
	}
}
