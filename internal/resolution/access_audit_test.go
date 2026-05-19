package resolution

import (
	"testing"

	"github.com/tamnguyendinh/avmatrix-go/internal/scanner"
	"github.com/tamnguyendinh/avmatrix-go/internal/scopeir"
)

func TestAuditAccessCandidatesClassifiesReasons(t *testing.T) {
	moduleScope := "scope:src/app.ts:module"
	functionScope := "scope:src/app.ts:start"
	userID := "def:User"
	nameID := "def:User.name"
	duplicateOneID := "def:User.dup1"
	duplicateTwoID := "def:User.dup2"
	emailID := "def:email"
	ir := scopeir.ScopeIR{
		FilePath:    "src/app.ts",
		FileHash:    "hash",
		Language:    scanner.TypeScript,
		ModuleScope: moduleScope,
		Scopes: []scopeir.ScopeFact{
			{ID: moduleScope, Kind: scopeir.ScopeModule, FilePath: "src/app.ts", TypeBindings: []scopeir.TypeBindingFact{
				{Name: "user", Type: scopeir.TypeRef{RawName: "User", Source: scopeir.TypeSourceAnnotation}},
				{Name: "external", Type: scopeir.TypeRef{RawName: "ExternalThing", Source: scopeir.TypeSourceAnnotation}},
			}},
			{ID: functionScope, Parent: scopeStringPtr(moduleScope), Kind: scopeir.ScopeFunction, FilePath: "src/app.ts", OwnedDefIDs: []string{"def:start"}},
		},
		Definitions: []scopeir.DefinitionFact{
			{ID: userID, FilePath: "src/app.ts", FileHash: "hash", Name: "User", Label: scopeir.NodeClass, Range: scopeir.Range{StartLine: 1, EndLine: 5}},
			{ID: "def:start", FilePath: "src/app.ts", FileHash: "hash", Name: "start", Label: scopeir.NodeFunction, Range: scopeir.Range{StartLine: 7, EndLine: 20}},
			{ID: nameID, FilePath: "src/app.ts", FileHash: "hash", Name: "name", Label: scopeir.NodeProperty, OwnerID: userID, Range: scopeir.Range{StartLine: 2, EndLine: 2}},
			{ID: duplicateOneID, FilePath: "src/app.ts", FileHash: "hash", Name: "dup", Label: scopeir.NodeProperty, OwnerID: userID, Range: scopeir.Range{StartLine: 3, EndLine: 3}},
			{ID: duplicateTwoID, FilePath: "src/app.ts", FileHash: "hash", Name: "dup", Label: scopeir.NodeProperty, OwnerID: userID, Range: scopeir.Range{StartLine: 4, EndLine: 4}},
			{ID: emailID, FilePath: "src/app.ts", FileHash: "hash", Name: "email", Label: scopeir.NodeProperty, Range: scopeir.Range{StartLine: 6, EndLine: 6}},
		},
		Accesses: []scopeir.AccessFact{
			{FilePath: "src/app.ts", FileHash: "hash", Name: "name", Kind: scopeir.AccessRead, ExplicitReceiver: "user", InScope: functionScope, Range: scopeir.Range{StartLine: 8}},
			{FilePath: "src/app.ts", FileHash: "hash", Name: "unknownType", Kind: scopeir.AccessRead, ExplicitReceiver: "missing", InScope: functionScope, Range: scopeir.Range{StartLine: 9}},
			{FilePath: "src/app.ts", FileHash: "hash", Name: "value", Kind: scopeir.AccessRead, ExplicitReceiver: "external", InScope: functionScope, Range: scopeir.Range{StartLine: 10}},
			{FilePath: "src/app.ts", FileHash: "hash", Name: "email", Kind: scopeir.AccessRead, ExplicitReceiver: "user", InScope: functionScope, Range: scopeir.Range{StartLine: 11}},
			{FilePath: "src/app.ts", FileHash: "hash", Name: "dup", Kind: scopeir.AccessRead, ExplicitReceiver: "user", InScope: functionScope, Range: scopeir.Range{StartLine: 12}},
			{FilePath: "src/app.ts", FileHash: "hash", Name: "absent", Kind: scopeir.AccessRead, ExplicitReceiver: "user", InScope: functionScope, Range: scopeir.Range{StartLine: 13}},
			{FilePath: "src/app.ts", FileHash: "hash", Name: "bad", Kind: scopeir.AccessRead, ExplicitReceiver: "user.items[0]", InScope: functionScope, Range: scopeir.Range{StartLine: 14}},
			{FilePath: "src/app.ts", FileHash: "hash", Name: "orphan", Kind: scopeir.AccessRead, ExplicitReceiver: "user", InScope: "missing-scope", Range: scopeir.Range{StartLine: 15}},
		},
	}

	result, err := AuditAccessCandidates([]scopeir.ScopeIR{ir}, AccessCandidateAuditOptions{MaxExamples: 1})
	if err != nil {
		t.Fatalf("AuditAccessCandidates() error = %v", err)
	}
	if result.Total != 8 || result.Resolved != 1 || result.Unresolved != 7 {
		t.Fatalf("totals = %#v", result)
	}
	assertAccessReason(t, result, "resolved", 1)
	assertAccessReason(t, result, "missing_receiver_type", 1)
	assertAccessReason(t, result, "external_library_type", 1)
	assertAccessReason(t, result, "missing_owner_link", 0)
	assertAccessReason(t, result, "ambiguous_owner", 1)
	assertAccessReason(t, result, "false_positive_candidate", 2)
	assertAccessReason(t, result, "unsupported_syntax", 1)
	assertAccessReason(t, result, "missing_caller", 1)
	if result.Languages["typescript"].Total != 8 || result.Languages["typescript"].Resolved != 1 {
		t.Fatalf("language stats = %#v", result.Languages)
	}
}

func TestAuditAccessCandidatesResolvesTypeAliasMembers(t *testing.T) {
	moduleScope := "scope:src/app.ts:module"
	functionScope := "scope:src/app.ts:start"
	resultID := "def:ReadResult"
	modelID := "def:ReadResult.model"
	ir := scopeir.ScopeIR{
		FilePath:    "src/app.ts",
		FileHash:    "hash",
		Language:    scanner.TypeScript,
		ModuleScope: moduleScope,
		Scopes: []scopeir.ScopeFact{
			{ID: moduleScope, Kind: scopeir.ScopeModule, FilePath: "src/app.ts"},
			{ID: functionScope, Parent: scopeStringPtr(moduleScope), Kind: scopeir.ScopeFunction, FilePath: "src/app.ts", OwnedDefIDs: []string{"def:start"}, TypeBindings: []scopeir.TypeBindingFact{
				{Name: "result", Type: scopeir.TypeRef{RawName: "ReadResult", Source: scopeir.TypeSourceReturn}},
			}},
		},
		Definitions: []scopeir.DefinitionFact{
			{ID: resultID, FilePath: "src/app.ts", FileHash: "hash", Name: "ReadResult", Label: scopeir.NodeTypeAlias, Range: scopeir.Range{StartLine: 1, EndLine: 3}},
			{ID: modelID, FilePath: "src/app.ts", FileHash: "hash", Name: "model", Label: scopeir.NodeProperty, OwnerID: resultID, DeclaredType: "InvoiceModel", Range: scopeir.Range{StartLine: 2, EndLine: 2}},
			{ID: "def:start", FilePath: "src/app.ts", FileHash: "hash", Name: "start", Label: scopeir.NodeFunction, Range: scopeir.Range{StartLine: 5, EndLine: 7}},
		},
		Accesses: []scopeir.AccessFact{
			{FilePath: "src/app.ts", FileHash: "hash", Name: "model", Kind: scopeir.AccessRead, ExplicitReceiver: "result", InScope: functionScope, Range: scopeir.Range{StartLine: 6}},
		},
	}

	result, err := AuditAccessCandidates([]scopeir.ScopeIR{ir}, AccessCandidateAuditOptions{MaxExamples: 1})
	if err != nil {
		t.Fatalf("AuditAccessCandidates() error = %v", err)
	}
	if result.Total != 1 || result.Resolved != 1 || result.Unresolved != 0 {
		t.Fatalf("totals = %#v", result)
	}
	assertAccessReason(t, result, "resolved", 1)
}

func TestAuditAccessCandidatesRejectsCrossLanguageGlobalOwner(t *testing.T) {
	moduleScope := "scope:src/app.ts:module"
	functionScope := "scope:src/app.ts:start"
	ir := scopeir.ScopeIR{
		FilePath:    "src/app.ts",
		FileHash:    "hash-ts",
		Language:    scanner.TypeScript,
		ModuleScope: moduleScope,
		Scopes: []scopeir.ScopeFact{
			{ID: moduleScope, Kind: scopeir.ScopeModule, FilePath: "src/app.ts"},
			{ID: functionScope, Parent: scopeStringPtr(moduleScope), Kind: scopeir.ScopeFunction, FilePath: "src/app.ts", OwnedDefIDs: []string{"def:start"}, TypeBindings: []scopeir.TypeBindingFact{
				{Name: "node", Type: scopeir.TypeRef{RawName: "GraphNode", Source: scopeir.TypeSourceParameter}},
			}},
		},
		Definitions: []scopeir.DefinitionFact{
			{ID: "def:start", FilePath: "src/app.ts", FileHash: "hash-ts", Name: "start", Label: scopeir.NodeFunction, Range: scopeir.Range{StartLine: 2, EndLine: 4}},
			{ID: "def:standalone-id", FilePath: "src/app.ts", FileHash: "hash-ts", Name: "id", Label: scopeir.NodeProperty, Range: scopeir.Range{StartLine: 10, EndLine: 10}},
		},
		Accesses: []scopeir.AccessFact{
			{FilePath: "src/app.ts", FileHash: "hash-ts", Name: "id", Kind: scopeir.AccessRead, ExplicitReceiver: "node", InScope: functionScope, Range: scopeir.Range{StartLine: 3}},
		},
	}
	goIR := scopeir.ScopeIR{
		FilePath: "internal/graphaccuracy/graphaccuracy.go",
		Language: scanner.Go,
		Definitions: []scopeir.DefinitionFact{
			{ID: "def:go-graph-node", FilePath: "internal/graphaccuracy/graphaccuracy.go", Name: "GraphNode", Label: scopeir.NodeStruct, Range: scopeir.Range{StartLine: 1, EndLine: 4}},
			{ID: "def:go-graph-node-id", FilePath: "internal/graphaccuracy/graphaccuracy.go", Name: "ID", Label: scopeir.NodeProperty, OwnerID: "def:go-graph-node", Range: scopeir.Range{StartLine: 2, EndLine: 2}},
		},
	}

	result, err := AuditAccessCandidates([]scopeir.ScopeIR{ir, goIR}, AccessCandidateAuditOptions{MaxExamples: 1})
	if err != nil {
		t.Fatalf("AuditAccessCandidates() error = %v", err)
	}
	assertAccessReason(t, result, "external_library_type", 1)
	assertAccessReason(t, result, "missing_owner_link", 0)
	assertAccessReason(t, result, "false_positive_candidate", 0)
}

func assertAccessReason(t *testing.T, result AccessCandidateAudit, reason string, want int) {
	t.Helper()
	if got := result.Reasons[reason].Count; got != want {
		t.Fatalf("reason %s = %d, want %d (%#v)", reason, got, want, result.Reasons)
	}
}

func scopeStringPtr(value string) *string {
	return &value
}
