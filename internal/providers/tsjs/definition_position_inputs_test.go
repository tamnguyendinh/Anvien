package tsjs

import (
	"bytes"
	"strings"
	"testing"

	"github.com/tamnguyendinh/anvien/internal/scanner"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
)

func TestExtractDefinitionPositionInputs(t *testing.T) {
	const lfSource = "function outer() {\n/*😀*/\tconst café = 1;\n}\n"

	lfIR := parseAndExtract(
		t,
		"src/positions.ts",
		"hash-lf",
		scanner.TypeScript,
		[]byte(lfSource),
	)
	crlfSource := strings.ReplaceAll(lfSource, "\n", "\r\n")
	crlfIR := parseAndExtract(
		t,
		"src/positions.ts",
		"hash-crlf",
		scanner.TypeScript,
		[]byte(crlfSource),
	)

	outerLF := requireDefinition(t, lfIR, "outer", scopeir.NodeFunction)
	cafeLF := requireDefinition(t, lfIR, "café", scopeir.NodeVariable)
	outerCRLF := requireDefinition(t, crlfIR, "outer", scopeir.NodeFunction)
	cafeCRLF := requireDefinition(t, crlfIR, "café", scopeir.NodeVariable)

	requireDefinitionRanges(
		t,
		outerLF,
		scopeir.Range{StartLine: 1, StartCol: 0, EndLine: 3, EndCol: 1},
		scopeir.Range{StartLine: 1, StartCol: 9, EndLine: 1, EndCol: 14},
	)
	requireDefinitionRanges(
		t,
		cafeLF,
		scopeir.Range{StartLine: 2, StartCol: 15, EndLine: 2, EndCol: 24},
		scopeir.Range{StartLine: 2, StartCol: 15, EndLine: 2, EndCol: 20},
	)
	if got := textAtSingleLineRange(t, []byte(lfSource), *cafeLF.SelectionRange); got != "café" {
		t.Fatalf("selection text = %q, want café", got)
	}
	if outerLF.Range != outerCRLF.Range || *outerLF.SelectionRange != *outerCRLF.SelectionRange {
		t.Fatalf("outer LF/CRLF ranges differ: LF=%#v/%#v CRLF=%#v/%#v", outerLF.Range, outerLF.SelectionRange, outerCRLF.Range, outerCRLF.SelectionRange)
	}
	if cafeLF.Range != cafeCRLF.Range || *cafeLF.SelectionRange != *cafeCRLF.SelectionRange {
		t.Fatalf("café LF/CRLF ranges differ: LF=%#v/%#v CRLF=%#v/%#v", cafeLF.Range, cafeLF.SelectionRange, cafeCRLF.Range, cafeCRLF.SelectionRange)
	}

	raw, err := lfIR.MarshalDeterministic()
	if err != nil {
		t.Fatalf("marshal ScopeIR: %v", err)
	}
	roundTripped, err := scopeir.Unmarshal(raw)
	if err != nil {
		t.Fatalf("unmarshal ScopeIR: %v", err)
	}
	cafeRoundTrip := requireDefinition(t, roundTripped, "café", scopeir.NodeVariable)
	requireDefinitionRanges(t, cafeRoundTrip, cafeLF.Range, *cafeLF.SelectionRange)
}

func TestExtractDefinitionOwnerMeaningAndLexicalInputs(t *testing.T) {
	source := []byte(`function outer() {
  const value = 1;
  function inner() {
    const value = 2;
  }
}
class Service {
  run() {}
}
interface Shared {}
const Shared = 3;
`)
	ir := parseAndExtract(t, "src/owners.ts", "hash-owners", scanner.TypeScript, source)

	values := definitionsNamed(ir, "value", scopeir.NodeVariable)
	if len(values) != 2 {
		t.Fatalf("value definitions = %d, want 2: %#v", len(values), values)
	}
	outerValueScope := requireSingleOwningScope(t, ir, values[0].ID)
	innerValueScope := requireSingleOwningScope(t, ir, values[1].ID)
	if outerValueScope.ID == innerValueScope.ID {
		t.Fatalf("nested value lexical owners both = %q, want distinct scopes", outerValueScope.ID)
	}
	if innerValueScope.Parent == nil || *innerValueScope.Parent != outerValueScope.ID {
		t.Fatalf("inner value scope parent = %#v, want %q", innerValueScope.Parent, outerValueScope.ID)
	}
	for _, definition := range values {
		if definition.SelectionRange == nil {
			t.Fatalf("value definition missing selection range: %#v", definition)
		}
		if got := textAtSingleLineRange(t, source, *definition.SelectionRange); got != "value" {
			t.Fatalf("value selection text = %q, want value", got)
		}
		owner := requireSingleOwningScope(t, ir, definition.ID)
		if !scopeBindsDefinition(owner, definition.Name, definition.ID) {
			t.Fatalf("scope %q does not bind %q to %q: %#v", owner.ID, definition.Name, definition.ID, owner.Bindings)
		}
	}

	service := requireDefinition(t, ir, "Service", scopeir.NodeClass)
	run := requireDefinition(t, ir, "run", scopeir.NodeMethod)
	if run.OwnerID != service.ID {
		t.Fatalf("run owner = %q, want %q", run.OwnerID, service.ID)
	}
	if run.SelectionRange == nil || textAtSingleLineRange(t, source, *run.SelectionRange) != "run" {
		t.Fatalf("run selection range = %#v, want the method token", run.SelectionRange)
	}

	sharedInterface := requireDefinition(t, ir, "Shared", scopeir.NodeInterface)
	sharedVariable := requireDefinition(t, ir, "Shared", scopeir.NodeVariable)
	if sharedInterface.ID == sharedVariable.ID {
		t.Fatalf("same-lexeme meaning lanes share provider ID %q", sharedInterface.ID)
	}
	for _, definition := range []scopeir.DefinitionFact{sharedInterface, sharedVariable} {
		if definition.SelectionRange == nil || textAtSingleLineRange(t, source, *definition.SelectionRange) != "Shared" {
			t.Fatalf("%s Shared selection range = %#v, want the declaring token", definition.Label, definition.SelectionRange)
		}
	}

	raw, err := ir.MarshalDeterministic()
	if err != nil {
		t.Fatalf("marshal owner/meaning ScopeIR: %v", err)
	}
	roundTripped, err := scopeir.Unmarshal(raw)
	if err != nil {
		t.Fatalf("unmarshal owner/meaning ScopeIR: %v", err)
	}
	runRoundTrip := requireDefinition(t, roundTripped, "run", scopeir.NodeMethod)
	if runRoundTrip.OwnerID != run.OwnerID || runRoundTrip.Label != run.Label || runRoundTrip.SelectionRange == nil || *runRoundTrip.SelectionRange != *run.SelectionRange {
		t.Fatalf("run inputs drifted on round trip: before=%#v after=%#v", run, runRoundTrip)
	}
	for _, definition := range []scopeir.DefinitionFact{sharedInterface, sharedVariable} {
		roundTripDefinition := requireDefinition(t, roundTripped, definition.Name, definition.Label)
		if roundTripDefinition.ID != definition.ID || roundTripDefinition.SelectionRange == nil || *roundTripDefinition.SelectionRange != *definition.SelectionRange {
			t.Fatalf("%s Shared inputs drifted on round trip: before=%#v after=%#v", definition.Label, definition, roundTripDefinition)
		}
	}
	roundTripValues := definitionsNamed(roundTripped, "value", scopeir.NodeVariable)
	if len(roundTripValues) != len(values) {
		t.Fatalf("round-trip value definitions = %d, want %d", len(roundTripValues), len(values))
	}
	for index, definition := range values {
		roundTripDefinition := roundTripValues[index]
		if roundTripDefinition.ID != definition.ID || roundTripDefinition.Label != definition.Label || roundTripDefinition.SelectionRange == nil || *roundTripDefinition.SelectionRange != *definition.SelectionRange {
			t.Fatalf("value inputs drifted on round trip: before=%#v after=%#v", definition, roundTripDefinition)
		}
		owner := requireSingleOwningScope(t, ir, definition.ID)
		roundTripOwner := requireSingleOwningScope(t, roundTripped, roundTripDefinition.ID)
		if roundTripOwner.ID != owner.ID || (roundTripOwner.Parent == nil) != (owner.Parent == nil) || (owner.Parent != nil && *roundTripOwner.Parent != *owner.Parent) || !scopeBindsDefinition(roundTripOwner, roundTripDefinition.Name, roundTripDefinition.ID) {
			t.Fatalf("value lexical input drifted on round trip: before=%#v after=%#v", owner, roundTripOwner)
		}
	}
}

func TestP1BPreservesBindingPatternBoundary(t *testing.T) {
	ir := parseAndExtract(
		t,
		"src/binding-pattern.ts",
		"hash-binding-pattern",
		scanner.TypeScript,
		[]byte("const { left } = source;\n"),
	)
	if definitions := definitionsNamed(ir, "left", scopeir.NodeVariable); len(definitions) != 0 {
		t.Fatalf("P1-B must preserve Child 03 binding-pattern ownership; got %#v", definitions)
	}
}

func requireDefinitionRanges(t *testing.T, definition scopeir.DefinitionFact, full scopeir.Range, selection scopeir.Range) {
	t.Helper()
	if definition.Range != full {
		t.Fatalf("%s full range = %#v, want %#v", definition.Name, definition.Range, full)
	}
	if definition.SelectionRange == nil || *definition.SelectionRange != selection {
		t.Fatalf("%s selection range = %#v, want %#v", definition.Name, definition.SelectionRange, selection)
	}
}

func definitionsNamed(ir scopeir.ScopeIR, name string, label scopeir.NodeLabel) []scopeir.DefinitionFact {
	definitions := make([]scopeir.DefinitionFact, 0)
	for _, definition := range ir.Definitions {
		if definition.Name == name && definition.Label == label {
			definitions = append(definitions, definition)
		}
	}
	return definitions
}

func requireSingleOwningScope(t *testing.T, ir scopeir.ScopeIR, definitionID string) scopeir.ScopeFact {
	t.Helper()
	owners := make([]scopeir.ScopeFact, 0, 1)
	for _, scope := range ir.Scopes {
		for _, ownedID := range scope.OwnedDefIDs {
			if ownedID == definitionID {
				owners = append(owners, scope)
			}
		}
	}
	if len(owners) != 1 {
		t.Fatalf("definition %q lexical owners = %d, want exactly 1: %#v", definitionID, len(owners), owners)
	}
	return owners[0]
}

func scopeBindsDefinition(scope scopeir.ScopeFact, name string, definitionID string) bool {
	for _, binding := range scope.Bindings {
		if binding.Name == name && binding.DefID == definitionID {
			return true
		}
	}
	return false
}

func textAtSingleLineRange(t *testing.T, source []byte, sourceRange scopeir.Range) string {
	t.Helper()
	if sourceRange.StartLine != sourceRange.EndLine || sourceRange.StartLine < 1 {
		t.Fatalf("range is not a valid single-line range: %#v", sourceRange)
	}
	lines := bytes.Split(source, []byte{'\n'})
	lineIndex := sourceRange.StartLine - 1
	if lineIndex >= len(lines) {
		t.Fatalf("range line %d exceeds %d source lines", sourceRange.StartLine, len(lines))
	}
	line := bytes.TrimSuffix(lines[lineIndex], []byte{'\r'})
	if sourceRange.StartCol < 0 || sourceRange.EndCol < sourceRange.StartCol || sourceRange.EndCol > len(line) {
		t.Fatalf("range columns are outside line bytes: range=%#v line=%q", sourceRange, line)
	}
	return string(line[sourceRange.StartCol:sourceRange.EndCol])
}
