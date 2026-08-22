package resolution

import (
	"testing"

	"github.com/tamnguyendinh/anvien/internal/scanner"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
)

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
