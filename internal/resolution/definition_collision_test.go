package resolution

import (
	"reflect"
	"strings"
	"testing"

	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/scanner"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
)

func TestP1DDefinitionConflictFailsClearAtResolveBoundary(t *testing.T) {
	base := p1dDefinitionFact()
	rangeConflict := base
	rangeConflict.Range = scopeir.Range{StartLine: 2, StartCol: 0, EndLine: 2, EndCol: 24}
	withReturnType := base
	withReturnType.ReturnType = "string"

	tests := []struct {
		name        string
		definitions []scopeir.DefinitionFact
	}{
		{name: "range mismatch", definitions: []scopeir.DefinitionFact{base, rangeConflict}},
		{name: "optional canonical property removed", definitions: []scopeir.DefinitionFact{withReturnType, base}},
		{name: "optional canonical property added", definitions: []scopeir.DefinitionFact{base, withReturnType}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := Resolve([]scopeir.ScopeIR{p1dDefinitionIR(test.definitions...)}, Options{})
			if err == nil {
				t.Fatal("Resolve() error = nil, want definition identity collision")
			}
			wantID := graphIDForDef(test.definitions[0], p1dModuleScopeID)
			if !strings.Contains(err.Error(), "definition identity collision") || !strings.Contains(err.Error(), wantID) {
				t.Fatalf("Resolve() error = %q, want clear collision for %q", err, wantID)
			}
			if result.Graph != nil {
				t.Fatalf("Resolve() returned a partial graph after conflict: %#v", result.Graph)
			}
			if !result.Metrics.BindingAccumulatorDisposed {
				t.Fatalf("collision path did not dispose binding accumulator: %#v", result.Metrics)
			}
		})
	}
}

func TestP1DDefinitionConflictDoesNotReplaceExistingNode(t *testing.T) {
	definition := p1dDefinitionFact()
	existing := p1dDefinitionNode(definition)
	existing.Properties["startLine"] = 2
	existing.Properties["enrichedBy"] = "pre-existing-phase"
	base := graph.New()
	base.AddNode(existing)

	result, err := ResolveInto(base, []scopeir.ScopeIR{p1dDefinitionIR(definition)}, Options{})
	if err == nil {
		t.Fatal("ResolveInto() error = nil, want definition identity collision")
	}
	if result.Graph != nil {
		t.Fatalf("ResolveInto() returned a partial graph after conflict: %#v", result.Graph)
	}
	got, ok := base.GetNode(existing.ID)
	if !ok || !reflect.DeepEqual(got, existing) {
		t.Fatalf("conflicting definition changed existing node: got %#v, want %#v", got, existing)
	}
	matchingNodes := 0
	for _, node := range base.Nodes {
		if node.ID == existing.ID {
			matchingNodes++
		}
	}
	if matchingNodes != 1 {
		t.Fatalf("conflicting definition node count = %d, want 1", matchingNodes)
	}
}

func TestP1DDefinitionExactReinsertionIsIdempotentAndPreservesEnrichment(t *testing.T) {
	definition := p1dDefinitionFact()
	enriched := p1dDefinitionNode(definition)
	enriched.Properties["enrichedBy"] = "pre-existing-phase"
	base := graph.New()
	base.AddNode(enriched)

	result, err := ResolveInto(base, []scopeir.ScopeIR{p1dDefinitionIR(definition, definition)}, Options{})
	if err != nil {
		t.Fatalf("ResolveInto() compatible reinsertion error = %v", err)
	}
	got, ok := result.Graph.GetNode(enriched.ID)
	if !ok || !reflect.DeepEqual(got, enriched) {
		t.Fatalf("compatible reinsertion changed enriched node: got %#v, want %#v", got, enriched)
	}
	matchingNodes := 0
	for _, node := range result.Graph.Nodes {
		if node.ID == enriched.ID {
			matchingNodes++
		}
	}
	if matchingNodes != 1 {
		t.Fatalf("node count after compatible reinsertion = %d, want 1", matchingNodes)
	}
}

func TestP1DDefinitionExactDuplicateKeepsOneNodeAndValidEndpoints(t *testing.T) {
	definition := p1dDefinitionFact()
	result, err := Resolve([]scopeir.ScopeIR{p1dDefinitionIR(definition, definition)}, Options{})
	if err != nil {
		t.Fatalf("Resolve() exact duplicate error = %v", err)
	}
	wantID := graphIDForDef(definition, p1dModuleScopeID)
	matchingNodes := 0
	for _, node := range result.Graph.Nodes {
		if node.ID == wantID {
			matchingNodes++
		}
	}
	if matchingNodes != 1 {
		t.Fatalf("exact duplicate graph nodes = %d, want 1", matchingNodes)
	}
	fileID := graph.GenerateID(string(scopeir.NodeFile), p1dFilePath)
	definesRelationships := 0
	for _, relationship := range result.Graph.Relationships {
		if relationship.Type == graph.RelDefines && relationship.SourceID == fileID && relationship.TargetID == wantID {
			definesRelationships++
		}
		if _, ok := result.Graph.GetNode(relationship.SourceID); !ok {
			t.Fatalf("relationship %s missing source endpoint %s", relationship.ID, relationship.SourceID)
		}
		if _, ok := result.Graph.GetNode(relationship.TargetID); !ok {
			t.Fatalf("relationship %s missing target endpoint %s", relationship.ID, relationship.TargetID)
		}
	}
	if definesRelationships != 1 {
		t.Fatalf("exact duplicate DEFINES relationships = %d, want 1", definesRelationships)
	}
}

const (
	p1dFilePath      = "src/collision.ts"
	p1dModuleScopeID = "scope:src/collision.ts#1:0-3:0:Module"
)

func p1dDefinitionFact() scopeir.DefinitionFact {
	return scopeir.DefinitionFact{
		ID:            "def:src/collision.ts#1:0:Function:run",
		FilePath:      p1dFilePath,
		Name:          "run",
		QualifiedName: "run",
		Label:         scopeir.NodeFunction,
		Range:         scopeir.Range{StartLine: 1, StartCol: 0, EndLine: 1, EndCol: 24},
	}
}

func p1dDefinitionNode(definition scopeir.DefinitionFact) graph.Node {
	return graph.Node{
		ID:    graphIDForDef(definition, p1dModuleScopeID),
		Label: definition.Label,
		Properties: graph.NodeProperties{
			"name":          definition.Name,
			"filePath":      definition.FilePath,
			"qualifiedName": definition.QualifiedName,
			"startLine":     definition.Range.StartLine,
			"endLine":       definition.Range.EndLine,
			"language":      string(scanner.TypeScript),
		},
	}
}

func p1dDefinitionIR(definitions ...scopeir.DefinitionFact) scopeir.ScopeIR {
	ownedDefIDs := make([]string, 0, len(definitions))
	bindings := make([]scopeir.BindingFact, 0, len(definitions))
	seen := map[string]struct{}{}
	for _, definition := range definitions {
		if _, ok := seen[definition.ID]; ok {
			continue
		}
		seen[definition.ID] = struct{}{}
		ownedDefIDs = append(ownedDefIDs, definition.ID)
		bindings = append(bindings, scopeir.BindingFact{
			Name:   definition.Name,
			DefID:  definition.ID,
			Origin: scopeir.BindingLocal,
		})
	}
	return scopeir.ScopeIR{
		FilePath:    p1dFilePath,
		Language:    scanner.TypeScript,
		ModuleScope: p1dModuleScopeID,
		Scopes: []scopeir.ScopeFact{{
			ID:          p1dModuleScopeID,
			Kind:        scopeir.ScopeModule,
			FilePath:    p1dFilePath,
			OwnedDefIDs: ownedDefIDs,
			Bindings:    bindings,
		}},
		Definitions: definitions,
	}
}
