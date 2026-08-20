package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/tamnguyendinh/anvien/internal/scopeir"
)

const syntheticTypeScript = `export async function inspectRows(load: () => Promise<[string[], number[]]>) {
  const [alphaRows, betaRows] = await load()
  const sourceBodySentinel = "DO_NOT_EMIT_SOURCE_BODY_SENTINEL"
  return betaRows.length + alphaRows.length + sourceBodySentinel.length
}
`

func TestRunProducesDeterministicMetadataInCallerOrder(t *testing.T) {
	filePath := writeRepoLocalTypeScript(t, syntheticTypeScript)
	args := []string{"-file", filePath, "-name", "betaRows", "-name", "alphaRows"}

	var first bytes.Buffer
	if err := run(args, &first, io.Discard); err != nil {
		t.Fatalf("first run: %v", err)
	}
	var second bytes.Buffer
	if err := run(args, &second, io.Discard); err != nil {
		t.Fatalf("second run: %v", err)
	}
	if !bytes.Equal(first.Bytes(), second.Bytes()) {
		t.Fatal("metadata JSON differs across identical runs")
	}

	var output probeOutput
	if err := json.Unmarshal(first.Bytes(), &output); err != nil {
		t.Fatalf("decode output: %v", err)
	}
	if output.SchemaVersion != probeSchemaVersion {
		t.Fatalf("schema version = %q, want %q", output.SchemaVersion, probeSchemaVersion)
	}
	if output.File.Language != "typescript" || output.File.Grammar != "typescript" || output.File.RootKind != "program" {
		t.Fatalf("unexpected parse metadata: %+v", output.File)
	}
	sum := sha256.Sum256([]byte(syntheticTypeScript))
	if output.File.SHA256 != hex.EncodeToString(sum[:]) || output.File.Bytes != len(syntheticTypeScript) {
		t.Fatalf("unexpected file identity: %+v", output.File)
	}

	wantOrder := []string{"betaRows", "alphaRows"}
	wantIndex := []int{1, 0}
	wantRange := []scopeir.Range{
		{StartLine: 2, StartCol: 20, EndLine: 2, EndCol: 28},
		{StartLine: 2, StartCol: 9, EndLine: 2, EndCol: 18},
	}
	if len(output.Results) != len(wantOrder) {
		t.Fatalf("result count = %d, want %d", len(output.Results), len(wantOrder))
	}
	for index, result := range output.Results {
		if result.Name != wantOrder[index] {
			t.Fatalf("result %d name = %q, want %q", index, result.Name, wantOrder[index])
		}
		if result.BindingLeaf.Provenance.Context != scopeir.BindingContextVariable {
			t.Fatalf("result %q context = %q", result.Name, result.BindingLeaf.Provenance.Context)
		}
		if result.BindingLeaf.Range != wantRange[index] || result.BindingLeaf.SelectionRange != wantRange[index] {
			t.Fatalf("result %q leaf ranges = %+v / %+v, want %+v", result.Name, result.BindingLeaf.Range, result.BindingLeaf.SelectionRange, wantRange[index])
		}
		if len(result.BindingLeaf.Path) != 1 ||
			result.BindingLeaf.Path[0].Kind != scopeir.BindingPathArrayIndex ||
			result.BindingLeaf.Path[0].ArrayIndex == nil ||
			*result.BindingLeaf.Path[0].ArrayIndex != wantIndex[index] {
			t.Fatalf("result %q path = %+v, want array-index(%d)", result.Name, result.BindingLeaf.Path, wantIndex[index])
		}
		if result.Definition.DefID == "" || result.Definition.Label != scopeir.NodeVariable {
			t.Fatalf("result %q definition = %+v", result.Name, result.Definition)
		}
		if result.Definition.Range != result.BindingLeaf.Range || result.Definition.SelectionRange != result.BindingLeaf.SelectionRange {
			t.Fatalf("result %q definition does not match leaf", result.Name)
		}
		if result.LexicalOwner.Kind != scopeir.ScopeFunction || !rangeContains(result.LexicalOwner.Range, result.Definition.Range) {
			t.Fatalf("result %q owner = %+v", result.Name, result.LexicalOwner)
		}
		if countString(result.LexicalOwner.OwnedDefIDs, result.Definition.DefID) != 1 || result.OwnedDefIDOccurrences != 1 {
			t.Fatalf("result %q owner DefID conservation failed", result.Name)
		}
		if result.OwnerLocalBinding.Name != result.Name ||
			result.OwnerLocalBinding.DefID != result.Definition.DefID ||
			result.OwnerLocalBinding.Origin != scopeir.BindingLocal ||
			result.LocalBindingOccurrences != 1 {
			t.Fatalf("result %q local binding = %+v", result.Name, result.OwnerLocalBinding)
		}
	}
}

func TestRunFailsClosedForMissingName(t *testing.T) {
	filePath := writeRepoLocalTypeScript(t, syntheticTypeScript)
	var stdout bytes.Buffer
	err := run([]string{"-file", filePath, "-name", "missingRows"}, &stdout, io.Discard)
	if err == nil {
		t.Fatal("missing name unexpectedly passed")
	}
	if !strings.Contains(err.Error(), `local "missingRows": expected exactly one variable-context BindingLeafFact, found 0`) {
		t.Fatalf("unexpected error: %v", err)
	}
	if stdout.Len() != 0 {
		t.Fatalf("failure emitted stdout: %q", stdout.String())
	}
}

func TestValidateBindingContractsFailsClosedOnDuplicateAndAmbiguousState(t *testing.T) {
	base := extractSyntheticIR(t, syntheticTypeScript)
	leafIndex, definitionIndex, ownerIndex, bindingIndex := requireSyntheticChain(t, base, "alphaRows")

	tests := []struct {
		name   string
		mutate func(*scopeir.ScopeIR)
		want   string
	}{
		{
			name: "duplicate leaf",
			mutate: func(ir *scopeir.ScopeIR) {
				ir.BindingLeaves = append(ir.BindingLeaves, ir.BindingLeaves[leafIndex])
			},
			want: "expected exactly one variable-context BindingLeafFact, found 2",
		},
		{
			name: "duplicate matching definition",
			mutate: func(ir *scopeir.ScopeIR) {
				ir.Definitions = append(ir.Definitions, ir.Definitions[definitionIndex])
			},
			want: "expected exactly one matching Variable DefinitionFact, found 2",
		},
		{
			name: "ambiguous owner",
			mutate: func(ir *scopeir.ScopeIR) {
				duplicate := ir.Scopes[ownerIndex]
				duplicate.ID += ":duplicate"
				duplicate.OwnedDefIDs = append([]string(nil), duplicate.OwnedDefIDs...)
				duplicate.Bindings = append([]scopeir.BindingFact(nil), duplicate.Bindings...)
				ir.Scopes = append(ir.Scopes, duplicate)
			},
			want: "expected exactly one lexical owner ScopeFact",
		},
		{
			name: "duplicate owner DefID",
			mutate: func(ir *scopeir.ScopeIR) {
				definitionID := ir.Definitions[definitionIndex].ID
				ir.Scopes[ownerIndex].OwnedDefIDs = append(ir.Scopes[ownerIndex].OwnedDefIDs, definitionID)
			},
			want: "times in OwnedDefIDs; expected exactly once",
		},
		{
			name: "duplicate owner-local binding",
			mutate: func(ir *scopeir.ScopeIR) {
				binding := ir.Scopes[ownerIndex].Bindings[bindingIndex]
				ir.Scopes[ownerIndex].Bindings = append(ir.Scopes[ownerIndex].Bindings, binding)
			},
			want: "owner-local BindingLocal mismatch",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ir := base.Normalized()
			test.mutate(&ir)
			_, err := validateBindingContracts(ir, []string{"alphaRows"})
			if err == nil {
				t.Fatal("invalid ScopeIR unexpectedly passed")
			}
			if !strings.Contains(err.Error(), test.want) {
				t.Fatalf("error = %q, want substring %q", err, test.want)
			}
		})
	}
}

func TestRunDoesNotEmitSourceBody(t *testing.T) {
	const secretLiteral = "SOURCE_BODY_LITERAL_839A5C01"
	const secretComment = "SOURCE_BODY_COMMENT_24B7F103"
	source := `function metadataOnly(input: [number, number]) {
  const [firstValue, secondValue] = input
  const hidden = "` + secretLiteral + `"
  // ` + secretComment + `
  return firstValue + secondValue + hidden.length
}
`
	filePath := writeRepoLocalTypeScript(t, source)
	var stdout bytes.Buffer
	if err := run([]string{"-file", filePath, "-name", "firstValue"}, &stdout, io.Discard); err != nil {
		t.Fatalf("run: %v", err)
	}
	for _, forbidden := range []string{secretLiteral, secretComment, source} {
		if strings.Contains(stdout.String(), forbidden) {
			t.Fatalf("metadata JSON contains forbidden source body %q", forbidden)
		}
	}
}

func TestCommandOptionsRejectDuplicateNamesAndOutputFlag(t *testing.T) {
	if _, err := parseCommandOptions(
		[]string{"-file", "source.ts", "-name", "value", "-name", "value"},
		io.Discard,
	); err == nil || !strings.Contains(err.Error(), `duplicate requested local name "value"`) {
		t.Fatalf("duplicate-name error = %v", err)
	}
	if _, err := parseCommandOptions(
		[]string{"-file", "source.ts", "-name", "value", "-out", "result.json"},
		io.Discard,
	); err == nil || !strings.Contains(err.Error(), "flag provided but not defined: -out") {
		t.Fatalf("-out error = %v", err)
	}
}

func extractSyntheticIR(t *testing.T, source string) scopeir.ScopeIR {
	t.Helper()
	sum := sha256.Sum256([]byte(source))
	ir, _, err := extractScopeIR(
		"E:/Anvien/.tmp/binding-contract-probe-tests/in-memory.ts",
		hex.EncodeToString(sum[:]),
		[]byte(source),
	)
	if err != nil {
		t.Fatalf("extract synthetic ScopeIR: %v", err)
	}
	return ir
}

func requireSyntheticChain(t *testing.T, ir scopeir.ScopeIR, name string) (int, int, int, int) {
	t.Helper()
	leafIndex := -1
	for index, leaf := range ir.BindingLeaves {
		if leaf.Name == name && leaf.Provenance.Context == scopeir.BindingContextVariable {
			leafIndex = index
			break
		}
	}
	if leafIndex < 0 {
		t.Fatalf("missing synthetic BindingLeafFact for %q", name)
	}
	leaf := ir.BindingLeaves[leafIndex]
	definitionIndex := -1
	for index, definition := range ir.Definitions {
		if definition.Name == name && definition.Label == scopeir.NodeVariable && definition.Range == leaf.Range {
			definitionIndex = index
			break
		}
	}
	if definitionIndex < 0 {
		t.Fatalf("missing synthetic DefinitionFact for %q", name)
	}
	definitionID := ir.Definitions[definitionIndex].ID
	ownerIndex := -1
	bindingIndex := -1
	for scopeIndex, owner := range ir.Scopes {
		if countString(owner.OwnedDefIDs, definitionID) != 1 {
			continue
		}
		ownerIndex = scopeIndex
		for index, binding := range owner.Bindings {
			if binding.Name == name && binding.DefID == definitionID && binding.Origin == scopeir.BindingLocal {
				bindingIndex = index
				break
			}
		}
		break
	}
	if ownerIndex < 0 || bindingIndex < 0 {
		t.Fatalf("missing synthetic owner/local binding for %q", name)
	}
	return leafIndex, definitionIndex, ownerIndex, bindingIndex
}

func writeRepoLocalTypeScript(t *testing.T, source string) string {
	t.Helper()
	_, testFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("resolve test file path")
	}
	repoRoot := filepath.Clean(filepath.Join(filepath.Dir(testFile), "..", ".."))
	base := filepath.Join(repoRoot, ".tmp", "binding-contract-probe-tests")
	if err := os.MkdirAll(base, 0o755); err != nil {
		t.Fatalf("create repo-local test temp base: %v", err)
	}
	directory, err := os.MkdirTemp(base, "case-")
	if err != nil {
		t.Fatalf("create repo-local test temp: %v", err)
	}
	t.Cleanup(func() {
		if err := os.RemoveAll(directory); err != nil {
			t.Errorf("remove repo-local test temp: %v", err)
		}
		_ = os.Remove(base)
	})
	filePath := filepath.Join(directory, "source.ts")
	if err := os.WriteFile(filePath, []byte(source), 0o600); err != nil {
		t.Fatalf("write synthetic TypeScript: %v", err)
	}
	return filePath
}
