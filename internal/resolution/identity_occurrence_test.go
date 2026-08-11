package resolution

import (
	"os"
	"reflect"
	"slices"
	"testing"

	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/scanner"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
)

func TestP1CIdentityConservesSameNameOccurrencesAndEndpoints(t *testing.T) {
	ir := p1cIdentityFixtureIR(t)
	workspace, err := buildWorkspace([]scopeir.ScopeIR{ir})
	if err != nil {
		t.Fatalf("buildWorkspace() error = %v", err)
	}

	filePath := cleanPath(ir.FilePath)
	definitions := workspace.defsByFile[filePath]
	if len(definitions) != len(ir.Definitions) {
		t.Fatalf("defsByFile occurrence denominator = %d, want %d", len(definitions), len(ir.Definitions))
	}
	graphIDs := make(map[string]struct{}, len(definitions))
	for _, definition := range definitions {
		graphIDs[definition.GraphID] = struct{}{}
	}
	if len(graphIDs) != len(definitions) {
		t.Fatalf("unique graph identities = %d, want %d occurrences", len(graphIDs), len(definitions))
	}

	for _, name := range []string{"time", "now"} {
		refs := exactNamedDefinitions(workspace.defsByName[name], name)
		if len(refs) != 2 {
			t.Fatalf("%s occurrences = %d, want 2: %#v", name, len(refs), refs)
		}
		occurrences := map[string]struct{}{}
		scopes := map[string]struct{}{}
		ids := map[string]struct{}{}
		for _, ref := range refs {
			occurrences[ref.Fact.ID] = struct{}{}
			scopes[workspace.scopeByDef[ref.Fact.ID]] = struct{}{}
			ids[ref.GraphID] = struct{}{}
		}
		if len(occurrences) != 2 || len(scopes) != 2 || len(ids) != 2 {
			t.Fatalf("%s identity inputs collapsed: occurrences=%v scopes=%v graphIDs=%v", name, occurrences, scopes, ids)
		}
	}

	shared := exactNamedDefinitions(workspace.defsByName["Shared"], "Shared")
	if len(shared) != 2 {
		t.Fatalf("Shared meaning-lane occurrences = %d, want 2: %#v", len(shared), shared)
	}
	labels := map[scopeir.NodeLabel]struct{}{}
	sharedIDs := map[string]struct{}{}
	for _, ref := range shared {
		labels[ref.Fact.Label] = struct{}{}
		sharedIDs[ref.GraphID] = struct{}{}
	}
	if _, ok := labels[scopeir.NodeInterface]; !ok {
		t.Fatalf("Shared labels = %v, missing Interface", labels)
	}
	if _, ok := labels[scopeir.NodeFunction]; !ok {
		t.Fatalf("Shared labels = %v, missing Function", labels)
	}
	if len(sharedIDs) != 2 {
		t.Fatalf("Shared meaning lanes share graph identity: %v", sharedIDs)
	}

	result, err := Resolve([]scopeir.ScopeIR{ir}, Options{})
	if err != nil {
		t.Fatalf("Resolve() error = %v", err)
	}
	definedTargets := map[string]struct{}{}
	for _, relationship := range result.Graph.Relationships {
		if _, ok := result.Graph.GetNode(relationship.SourceID); !ok {
			t.Fatalf("relationship %s missing source endpoint %s", relationship.ID, relationship.SourceID)
		}
		if _, ok := result.Graph.GetNode(relationship.TargetID); !ok {
			t.Fatalf("relationship %s missing target endpoint %s", relationship.ID, relationship.TargetID)
		}
		if relationship.Type == graph.RelDefines {
			definedTargets[relationship.TargetID] = struct{}{}
		}
	}
	for _, definition := range definitions {
		if _, ok := result.Graph.GetNode(definition.GraphID); !ok {
			t.Fatalf("definition occurrence %s missing graph node %s", definition.Fact.ID, definition.GraphID)
		}
		if _, ok := definedTargets[definition.GraphID]; !ok {
			t.Fatalf("definition occurrence %s missing DEFINES endpoint %s", definition.Fact.ID, definition.GraphID)
		}
	}
}

func TestP1CIdentityIsDeterministicAcrossDefinitionOrder(t *testing.T) {
	first := p1cIdentityFixtureIR(t)
	second := p1cIdentityFixtureIR(t)
	slices.Reverse(second.Definitions)
	slices.Reverse(second.Scopes)
	for index := range second.Scopes {
		slices.Reverse(second.Scopes[index].OwnedDefIDs)
		slices.Reverse(second.Scopes[index].Bindings)
	}

	firstWorkspace, err := buildWorkspace([]scopeir.ScopeIR{first})
	if err != nil {
		t.Fatalf("first buildWorkspace() error = %v", err)
	}
	secondWorkspace, err := buildWorkspace([]scopeir.ScopeIR{second})
	if err != nil {
		t.Fatalf("second buildWorkspace() error = %v", err)
	}
	firstIDs := graphIDsByOccurrence(firstWorkspace)
	secondIDs := graphIDsByOccurrence(secondWorkspace)
	if !reflect.DeepEqual(firstIDs, secondIDs) {
		t.Fatalf("reordered identity set changed\nfirst=%v\nsecond=%v", firstIDs, secondIDs)
	}
}

func TestP1CIdentityUsesLexicalMeaningOwnerAndArityInputs(t *testing.T) {
	base := scopeir.DefinitionFact{
		ID:            "def:src/model.ts#3:2:Property:value",
		FilePath:      "src/model.ts",
		Name:          "value",
		QualifiedName: "Box.value",
		Label:         scopeir.NodeProperty,
		OwnerID:       "def:src/model.ts#1:0:Class:Box",
	}
	baseID := graphIDForDef(base, "scope:src/model.ts#1:0-5:1:Class")

	differentScopeID := graphIDForDef(base, "scope:src/model.ts#2:0-4:1:Function")
	if baseID == differentScopeID {
		t.Fatalf("lexical scope input did not change graph identity: %s", baseID)
	}
	differentOwner := base
	differentOwner.OwnerID = "def:src/model.ts#6:0:Class:Other"
	if baseID == graphIDForDef(differentOwner, "scope:src/model.ts#1:0-5:1:Class") {
		t.Fatalf("owner input did not change graph identity: %s", baseID)
	}
	differentMeaning := base
	differentMeaning.Label = scopeir.NodeVariable
	if baseID == graphIDForDef(differentMeaning, "scope:src/model.ts#1:0-5:1:Class") {
		t.Fatalf("meaning label did not change graph identity: %s", baseID)
	}

	callable := scopeir.DefinitionFact{
		ID:       "def:src/model.ts#8:0:Function:read",
		FilePath: "src/model.ts",
		Name:     "read",
		Label:    scopeir.NodeFunction,
	}
	zeroArity := callable
	zeroArity.ParameterCount = intPtr(0)
	if graphIDForDef(callable, "scope:module") == graphIDForDef(zeroArity, "scope:module") {
		t.Fatalf("nil and zero arity identities collapsed")
	}

	slashVariant := base
	slashVariant.FilePath = "src\\model.ts"
	if baseID != graphIDForDef(slashVariant, "scope:src/model.ts#1:0-5:1:Class") {
		t.Fatalf("equivalent repository-relative path separators changed graph identity")
	}
}

func TestP1CIdentityDoesNotMergeSameNameWithoutProviderEvidence(t *testing.T) {
	moduleScope := "scope:src/overloads.ts#1:0-3:1:Module"
	first := scopeir.DefinitionFact{
		ID:             "def:src/overloads.ts#1:0:Function:read",
		FilePath:       "src/overloads.ts",
		Name:           "read",
		QualifiedName:  "read",
		Label:          scopeir.NodeFunction,
		ParameterCount: intPtr(1),
		Range:          scopeir.Range{StartLine: 1, EndLine: 1},
	}
	second := first
	second.ID = "def:src/overloads.ts#2:0:Function:read"
	second.Range = scopeir.Range{StartLine: 2, EndLine: 2}
	ir := scopeir.ScopeIR{
		FilePath:    "src/overloads.ts",
		Language:    scanner.TypeScript,
		ModuleScope: moduleScope,
		Scopes: []scopeir.ScopeFact{{
			ID:          moduleScope,
			Kind:        scopeir.ScopeModule,
			FilePath:    "src/overloads.ts",
			OwnedDefIDs: []string{first.ID, second.ID},
			Bindings: []scopeir.BindingFact{
				{Name: "read", DefID: first.ID, Origin: scopeir.BindingLocal},
				{Name: "read", DefID: second.ID, Origin: scopeir.BindingLocal},
			},
		}},
		Definitions: []scopeir.DefinitionFact{first, second},
	}

	workspace, err := buildWorkspace([]scopeir.ScopeIR{ir})
	if err != nil {
		t.Fatalf("buildWorkspace() error = %v", err)
	}
	refs := exactNamedDefinitions(workspace.defsByName["read"], "read")
	if len(refs) != 2 {
		t.Fatalf("same-name provider occurrences = %d, want 2: %#v", len(refs), refs)
	}
	if refs[0].GraphID == refs[1].GraphID {
		t.Fatalf("same-name occurrences merged without provider evidence: %s", refs[0].GraphID)
	}
}

func p1cIdentityFixtureIR(t *testing.T) scopeir.ScopeIR {
	t.Helper()
	source, err := os.ReadFile("testdata/p1c_identity_repo/src/identity.ts")
	if err != nil {
		t.Fatalf("read P1-C identity fixture: %v", err)
	}
	return parseTypeScriptSource(t, "src/identity.ts", source)
}

func exactNamedDefinitions(definitions []defRef, name string) []defRef {
	out := make([]defRef, 0, len(definitions))
	for _, definition := range definitions {
		if definition.Fact.Name == name {
			out = append(out, definition)
		}
	}
	return out
}

func graphIDsByOccurrence(workspace *workspace) map[string]string {
	out := map[string]string{}
	for _, definitions := range workspace.defsByFile {
		for _, definition := range definitions {
			out[definition.Fact.ID] = definition.GraphID
		}
	}
	return out
}
