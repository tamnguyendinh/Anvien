package resolution

import (
	"testing"

	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/scanner"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
)

func TestDefinitionGraphProjectionPreservesCorrectedFactsAndDefinesEndpoints(t *testing.T) {
	moduleScope := "scope:src/app.ts#module"
	selection := scopeir.Range{StartLine: 2, StartCol: 0, EndLine: 2, EndCol: 6}
	definitions := []scopeir.DefinitionFact{
		{
			ID:             "def:src/app.ts#2:0:Function:worker",
			FilePath:       "src/app.ts",
			Name:           "worker",
			QualifiedName:  "api.worker",
			Label:          scopeir.NodeFunction,
			Range:          scopeir.Range{StartLine: 2, StartCol: 0, EndLine: 4, EndCol: 1},
			SelectionRange: &selection,
		},
		{
			ID:            "def:src/app.ts#6:0:Variable:config",
			FilePath:      "src/app.ts",
			Name:          "config",
			QualifiedName: "config",
			Label:         scopeir.NodeVariable,
			Range:         scopeir.Range{StartLine: 6, StartCol: 0, EndLine: 6, EndCol: 24},
		},
	}
	ir := scopeir.ScopeIR{
		FilePath:    "src/app.ts",
		Language:    scanner.TypeScript,
		ModuleScope: moduleScope,
		Scopes: []scopeir.ScopeFact{{
			ID:          moduleScope,
			Kind:        scopeir.ScopeModule,
			FilePath:    "src/app.ts",
			OwnedDefIDs: []string{definitions[0].ID, definitions[1].ID},
		}},
		Definitions: definitions,
	}

	result, err := Resolve([]scopeir.ScopeIR{ir}, Options{})
	if err != nil {
		t.Fatalf("Resolve() error = %v", err)
	}
	workspace, err := buildWorkspace([]scopeir.ScopeIR{ir})
	if err != nil {
		t.Fatalf("buildWorkspace() error = %v", err)
	}

	definesTargets := map[string]struct{}{}
	for _, relationship := range result.Graph.Relationships {
		if _, ok := result.Graph.GetNode(relationship.SourceID); !ok {
			t.Fatalf("relationship %q missing source endpoint %q", relationship.ID, relationship.SourceID)
		}
		if _, ok := result.Graph.GetNode(relationship.TargetID); !ok {
			t.Fatalf("relationship %q missing target endpoint %q", relationship.ID, relationship.TargetID)
		}
		if relationship.Type == graph.RelDefines {
			definesTargets[relationship.TargetID] = struct{}{}
		}
	}

	for _, definition := range workspace.defsByFile["src/app.ts"] {
		node, ok := result.Graph.GetNode(definition.GraphID)
		if !ok {
			t.Fatalf("definition %q missing graph node %q", definition.Fact.ID, definition.GraphID)
		}
		if node.Label != definition.Fact.Label || node.Properties["name"] != definition.Fact.Name || node.Properties["qualifiedName"] != definition.Fact.QualifiedName {
			t.Fatalf("definition semantic facts drifted: node=%#v fact=%#v", node, definition.Fact)
		}
		assertIntProperty(t, node, "startLine", definition.Fact.Range.StartLine)
		assertIntProperty(t, node, "startCol", definition.Fact.Range.StartCol)
		assertIntProperty(t, node, "endLine", definition.Fact.Range.EndLine)
		assertIntProperty(t, node, "endCol", definition.Fact.Range.EndCol)
		if definition.Fact.SelectionRange == nil {
			for _, key := range []string{"selectionStartLine", "selectionStartCol", "selectionEndLine", "selectionEndCol"} {
				if _, exists := node.Properties[key]; exists {
					t.Fatalf("definition %q has unexpected absent selection property %q: %#v", definition.Fact.ID, key, node.Properties)
				}
			}
		} else {
			assertIntProperty(t, node, "selectionStartLine", definition.Fact.SelectionRange.StartLine)
			assertIntProperty(t, node, "selectionStartCol", definition.Fact.SelectionRange.StartCol)
			assertIntProperty(t, node, "selectionEndLine", definition.Fact.SelectionRange.EndLine)
			assertIntProperty(t, node, "selectionEndCol", definition.Fact.SelectionRange.EndCol)
		}
		if _, ok := definesTargets[definition.GraphID]; !ok {
			t.Fatalf("definition %q missing DEFINES target %q", definition.Fact.ID, definition.GraphID)
		}
	}
	if len(definesTargets) != len(definitions) {
		t.Fatalf("DEFINES target count = %d, want %d", len(definesTargets), len(definitions))
	}
}

func assertIntProperty(t *testing.T, node graph.Node, key string, want int) {
	t.Helper()
	got, ok := node.Properties[key].(int)
	if !ok || got != want {
		t.Fatalf("node %q property %s = %#v, want int(%d)", node.ID, key, node.Properties[key], want)
	}
}
