package resolution

import (
	"sort"
	"testing"

	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/graphhealth"
	"github.com/tamnguyendinh/anvien/internal/scanner"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
)

func TestResolveMemberAccessRejectsCrossLanguageGlobalOwner(t *testing.T) {
	moduleScope := "scope:src/app.ts:module"
	functionScope := "scope:src/app.ts:start"
	tsIR := scopeir.ScopeIR{
		FilePath:    "src/app.ts",
		FileHash:    "hash-ts",
		Language:    scanner.TypeScript,
		ModuleScope: moduleScope,
		Scopes: []scopeir.ScopeFact{
			{ID: moduleScope, Kind: scopeir.ScopeModule, FilePath: "src/app.ts"},
			{ID: functionScope, Parent: &[]string{moduleScope}[0], Kind: scopeir.ScopeFunction, FilePath: "src/app.ts", OwnedDefIDs: []string{"def:start"}, TypeBindings: []scopeir.TypeBindingFact{
				{Name: "node", Type: scopeir.TypeRef{RawName: "GraphNode", Source: scopeir.TypeSourceParameter}},
			}},
		},
		Definitions: []scopeir.DefinitionFact{
			{ID: "def:start", FilePath: "src/app.ts", FileHash: "hash-ts", Name: "start", Label: scopeir.NodeFunction, Range: scopeir.Range{StartLine: 1, EndLine: 3}},
		},
		Accesses: []scopeir.AccessFact{
			{FilePath: "src/app.ts", FileHash: "hash-ts", Name: "ID", Kind: scopeir.AccessRead, ExplicitReceiver: "node", InScope: functionScope, Range: scopeir.Range{StartLine: 2, EndLine: 2}},
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

	result, err := Resolve([]scopeir.ScopeIR{tsIR, goIR}, Options{})
	if err != nil {
		t.Fatalf("resolve failed: %v", err)
	}
	requireNoRelationship(t, result.Graph, graph.RelAccesses, "Function:src/app.ts:start", "Property:internal/graphaccuracy/graphaccuracy.go:GraphNode.ID")
	if result.Metrics.ResolvedAccesses != 0 || result.Metrics.UnresolvedReferences != 1 {
		t.Fatalf("unexpected metrics: %#v", result.Metrics)
	}
}

func TestResolveImportedConstSelectorDoesNotEmitAccess(t *testing.T) {
	moduleScope := "scope:cmd/app/main.go:module"
	functionScope := "scope:cmd/app/main.go:main"
	targetRaw := "github.com/tamnguyendinh/anvien/internal/pkg"
	source := scopeir.ScopeIR{
		FilePath:    "cmd/app/main.go",
		FileHash:    "hash-main",
		Language:    scanner.Go,
		ModuleScope: moduleScope,
		Imports: []scopeir.ImportFact{{
			FilePath:     "cmd/app/main.go",
			Kind:         scopeir.ImportNamed,
			LocalName:    "pkg",
			ImportedName: "pkg",
			TargetRaw:    &targetRaw,
		}},
		Scopes: []scopeir.ScopeFact{
			{ID: moduleScope, Kind: scopeir.ScopeModule, FilePath: "cmd/app/main.go"},
			{ID: functionScope, Parent: &[]string{moduleScope}[0], Kind: scopeir.ScopeFunction, FilePath: "cmd/app/main.go", OwnedDefIDs: []string{"def:main"}},
		},
		Definitions: []scopeir.DefinitionFact{
			{ID: "def:main", FilePath: "cmd/app/main.go", FileHash: "hash-main", Name: "main", QualifiedName: "main", Label: scopeir.NodeFunction, Range: scopeir.Range{StartLine: 3, EndLine: 5}},
		},
		Accesses: []scopeir.AccessFact{
			{FilePath: "cmd/app/main.go", FileHash: "hash-main", Name: "Mode", Kind: scopeir.AccessRead, ExplicitReceiver: "pkg", InScope: functionScope, Range: scopeir.Range{StartLine: 4, EndLine: 4}},
		},
	}
	target := scopeir.ScopeIR{
		FilePath: "internal/pkg/constants.go",
		FileHash: "hash-pkg",
		Language: scanner.Go,
		Definitions: []scopeir.DefinitionFact{
			{ID: "def:pkg-mode", FilePath: "internal/pkg/constants.go", FileHash: "hash-pkg", Name: "Mode", QualifiedName: "Mode", Label: scopeir.NodeConst, Range: scopeir.Range{StartLine: 1, EndLine: 1}},
		},
	}

	result, err := Resolve([]scopeir.ScopeIR{source, target}, Options{})
	if err != nil {
		t.Fatalf("resolve failed: %v", err)
	}
	mainFn := requireNode(t, result.Graph, "Function", "cmd/app/main.go", "main")
	mode := requireNode(t, result.Graph, "Const", "internal/pkg/constants.go", "Mode")
	requireNoRelationship(t, result.Graph, graph.RelAccesses, mainFn.ID, mode.ID)
	if result.Metrics.ResolvedAccesses != 0 || result.Metrics.UnresolvedReferences != 1 {
		t.Fatalf("unexpected metrics: %#v", result.Metrics)
	}
	diagnostics, ok := mainFn.Properties[graphhealth.DiagnosticPropertyKey].([]graphhealth.Diagnostic)
	if !ok || len(diagnostics) != 1 {
		t.Fatalf("main diagnostics = %#v", mainFn.Properties[graphhealth.DiagnosticPropertyKey])
	}
	if diagnostics[0].TargetText != "pkg.Mode" ||
		diagnostics[0].SourceSiteStatus != sourceSiteStatusUnresolvedLocalBinding ||
		diagnostics[0].TargetRole != targetRoleMember {
		t.Fatalf("unexpected imported const selector diagnostic: %#v", diagnostics[0])
	}
}

func TestResolveIntoPreservesExistingFileNodeMetadata(t *testing.T) {
	base := graph.New()
	base.AddNode(graph.Node{
		ID:    graph.GenerateID("File", "README.md"),
		Label: scopeir.NodeFile,
		Properties: graph.NodeProperties{
			"name":         "README.md",
			"filePath":     "README.md",
			"documentKind": "markdown",
		},
	})
	base.AddRelationship(graph.Relationship{
		ID:         graph.GenerateID(string(graph.RelContains), "Folder:docs->File:README.md"),
		SourceID:   "Folder:docs",
		TargetID:   graph.GenerateID("File", "README.md"),
		Type:       graph.RelContains,
		Confidence: 1,
		Reason:     "project structure",
	})
	ir := scopeir.ScopeIR{
		FilePath: "README.md",
		Language: scanner.Markdown,
	}

	result, err := ResolveInto(base, []scopeir.ScopeIR{ir}, Options{})
	if err != nil {
		t.Fatalf("ResolveInto() error = %v", err)
	}
	node, ok := result.Graph.GetNode(graph.GenerateID("File", "README.md"))
	if !ok {
		t.Fatalf("file node missing")
	}
	if node.Properties["documentKind"] != "markdown" || node.Properties["language"] != string(scanner.Markdown) {
		t.Fatalf("file properties = %#v", node.Properties)
	}
	if _, ok := result.Graph.GetRelationship(graph.GenerateID(string(graph.RelContains), "Folder:docs->File:README.md")); !ok {
		t.Fatalf("pre-existing relationship was not preserved")
	}
}

func TestResolveAnnotatesFrameworkHintProperties(t *testing.T) {
	moduleScope := "scope:app/api/users/route.ts#1:0-3:1:Module"
	functionScope := "scope:app/api/users/route.ts#1:0-3:1:Function"
	functionDef := scopeir.DefinitionFact{
		ID:            "def:app/api/users/route.ts#1:0:Function:GET",
		FilePath:      "app/api/users/route.ts",
		Name:          "GET",
		Label:         scopeir.NodeFunction,
		QualifiedName: "GET",
		Range:         scopeir.Range{StartLine: 1, EndLine: 3},
	}
	ir := scopeir.ScopeIR{
		FilePath:    "app/api/users/route.ts",
		FileHash:    "hash-route",
		Language:    scanner.TypeScript,
		ModuleScope: moduleScope,
		Scopes: []scopeir.ScopeFact{
			{ID: moduleScope, Kind: scopeir.ScopeModule, FilePath: "app/api/users/route.ts", Range: scopeir.Range{StartLine: 1, EndLine: 4}},
			{ID: functionScope, Parent: &[]string{moduleScope}[0], Kind: scopeir.ScopeFunction, FilePath: "app/api/users/route.ts", Range: scopeir.Range{StartLine: 1, EndLine: 3}, OwnedDefIDs: []string{functionDef.ID}, Bindings: []scopeir.BindingFact{{Name: "GET", DefID: functionDef.ID, Origin: scopeir.BindingLocal}}},
		},
		Definitions: []scopeir.DefinitionFact{functionDef},
	}

	result, err := Resolve([]scopeir.ScopeIR{ir}, Options{})
	if err != nil {
		t.Fatalf("resolve failed: %v", err)
	}
	fileNode, ok := result.Graph.GetNode(graph.GenerateID("File", "app/api/users/route.ts"))
	if !ok {
		t.Fatalf("file node missing")
	}
	functionNode := requireNode(t, result.Graph, "Function", "app/api/users/route.ts", "GET")
	for _, node := range []graph.Node{fileNode, functionNode} {
		if node.Properties["framework"] != "nextjs-api" || node.Properties["frameworkReason"] != "nextjs-api-route" {
			t.Fatalf("framework properties missing on %s: %#v", node.ID, node.Properties)
		}
		if node.Properties["astFrameworkMultiplier"] != 3.0 || node.Properties["astFrameworkReason"] != "nextjs-api-route" {
			t.Fatalf("process framework properties missing on %s: %#v", node.ID, node.Properties)
		}
	}
}

func TestResolveAppliesScopeIRFrameworkFacts(t *testing.T) {
	moduleScope := "scope:src/users.controller.ts#1:0-4:1:Module"
	classScope := "scope:src/users.controller.ts#2:0-4:1:Class"
	classDef := scopeir.DefinitionFact{
		ID:            "def:src/users.controller.ts#2:0:Class:UsersController",
		FilePath:      "src/users.controller.ts",
		Name:          "UsersController",
		Label:         scopeir.NodeClass,
		QualifiedName: "UsersController",
		Range:         scopeir.Range{StartLine: 2, EndLine: 4},
	}
	ir := scopeir.ScopeIR{
		FilePath:    "src/users.controller.ts",
		FileHash:    "hash-users-controller",
		Language:    scanner.TypeScript,
		ModuleScope: moduleScope,
		Scopes: []scopeir.ScopeFact{
			{ID: moduleScope, Kind: scopeir.ScopeModule, FilePath: "src/users.controller.ts", Range: scopeir.Range{StartLine: 1, EndLine: 4}, OwnedDefIDs: []string{classDef.ID}},
			{ID: classScope, Parent: &[]string{moduleScope}[0], Kind: scopeir.ScopeClass, FilePath: "src/users.controller.ts", Range: scopeir.Range{StartLine: 2, EndLine: 4}, OwnedDefIDs: []string{classDef.ID}},
		},
		Definitions: []scopeir.DefinitionFact{classDef},
		Frameworks: []scopeir.FrameworkFact{{
			DefID:                classDef.ID,
			FilePath:             classDef.FilePath,
			FileHash:             "hash-users-controller",
			Framework:            "nestjs",
			Reason:               "nestjs-decorator",
			EntryPointMultiplier: 3.2,
			Range:                classDef.Range,
		}},
	}

	result, err := Resolve([]scopeir.ScopeIR{ir}, Options{})
	if err != nil {
		t.Fatalf("resolve failed: %v", err)
	}
	classNode := requireNode(t, result.Graph, "Class", "src/users.controller.ts", "UsersController")
	if classNode.Properties["framework"] != "nestjs" || classNode.Properties["frameworkReason"] != "nestjs-decorator" {
		t.Fatalf("framework fact properties missing: %#v", classNode.Properties)
	}
	if classNode.Properties["astFrameworkMultiplier"] != 3.2 || classNode.Properties["astFrameworkReason"] != "nestjs-decorator" {
		t.Fatalf("AST framework fact properties missing: %#v", classNode.Properties)
	}
}

func TestResolveExpandsGoPackageImportsToPackageFiles(t *testing.T) {
	targetRaw := "github.com/tamnguyendinh/anvien/internal/pkg"
	source := scopeir.ScopeIR{
		FilePath:    "cmd/app/main.go",
		Language:    scanner.Go,
		ModuleScope: "scope:cmd/app/main.go#1:0-4:1:Module",
		Imports: []scopeir.ImportFact{{
			FilePath:     "cmd/app/main.go",
			Kind:         scopeir.ImportNamed,
			LocalName:    "pkg",
			ImportedName: "pkg",
			TargetRaw:    &targetRaw,
		}},
	}
	targetA := scopeir.ScopeIR{FilePath: "internal/pkg/a.go", Language: scanner.Go}
	targetB := scopeir.ScopeIR{FilePath: "internal/pkg/b.go", Language: scanner.Go}
	targetTest := scopeir.ScopeIR{FilePath: "internal/pkg/a_test.go", Language: scanner.Go}

	result, err := Resolve([]scopeir.ScopeIR{source, targetTest, targetB, targetA}, Options{})
	if err != nil {
		t.Fatalf("resolve failed: %v", err)
	}
	requireRelationship(t, result.Graph, graph.RelImports, "File:cmd/app/main.go", "File:internal/pkg/a.go")
	requireRelationship(t, result.Graph, graph.RelImports, "File:cmd/app/main.go", "File:internal/pkg/b.go")
	requireNoRelationship(t, result.Graph, graph.RelImports, "File:cmd/app/main.go", "File:internal/pkg/a_test.go")
	if result.Metrics.FinalizedImportsEmitted != 2 || result.Metrics.ImportsResolved != 2 {
		t.Fatalf("import metrics = %#v, want two package-file imports", result.Metrics)
	}
}

func TestResolveMergesDuplicateSemanticEdgesAndCountsUnresolved(t *testing.T) {
	moduleScope := "scope:src/app.ts#1:0-3:1:Module"
	functionScope := "scope:src/app.ts#1:0-3:1:Function"
	functionDef := scopeir.DefinitionFact{
		ID:            "def:src/app.ts#1:0:Function:start",
		FilePath:      "src/app.ts",
		Name:          "start",
		Label:         scopeir.NodeFunction,
		QualifiedName: "start",
		Range:         scopeir.Range{StartLine: 1, EndLine: 3},
	}
	targetDef := scopeir.DefinitionFact{
		ID:            "def:src/app.ts#4:0:Function:target",
		FilePath:      "src/app.ts",
		Name:          "target",
		Label:         scopeir.NodeFunction,
		QualifiedName: "target",
		Range:         scopeir.Range{StartLine: 4, EndLine: 4},
	}
	ir := scopeir.ScopeIR{
		FilePath:    "src/app.ts",
		FileHash:    "hash-app",
		Language:    scanner.TypeScript,
		ModuleScope: moduleScope,
		Scopes: []scopeir.ScopeFact{
			{ID: moduleScope, Kind: scopeir.ScopeModule, FilePath: "src/app.ts", Range: scopeir.Range{StartLine: 1, EndLine: 4}},
			{ID: functionScope, Parent: &[]string{moduleScope}[0], Kind: scopeir.ScopeFunction, FilePath: "src/app.ts", Range: scopeir.Range{StartLine: 1, EndLine: 3}, OwnedDefIDs: []string{functionDef.ID}, Bindings: []scopeir.BindingFact{{Name: "start", DefID: functionDef.ID, Origin: scopeir.BindingLocal}, {Name: "target", DefID: targetDef.ID, Origin: scopeir.BindingLocal}, {Name: "aliasTarget", DefID: targetDef.ID, Origin: scopeir.BindingLocal}}},
		},
		Definitions: []scopeir.DefinitionFact{functionDef, targetDef},
		Calls: []scopeir.CallSiteFact{
			{FilePath: "src/app.ts", FileHash: "hash-app", Name: "target", InScope: functionScope, CallForm: scopeir.CallFree, Range: scopeir.Range{StartLine: 2, StartCol: 2, EndLine: 2}},
			{FilePath: "src/app.ts", FileHash: "hash-app", Name: "target", InScope: functionScope, CallForm: scopeir.CallFree, Range: scopeir.Range{StartLine: 3, StartCol: 2, EndLine: 3}},
			{FilePath: "src/app.ts", FileHash: "hash-app", Name: "aliasTarget", InScope: functionScope, CallForm: scopeir.CallFree, Range: scopeir.Range{StartLine: 3, StartCol: 6, EndLine: 3}},
			{FilePath: "src/app.ts", FileHash: "hash-app", Name: "missing", InScope: functionScope, CallForm: scopeir.CallFree, Range: scopeir.Range{StartLine: 3, StartCol: 10, EndLine: 3}},
		},
	}
	result, err := Resolve([]scopeir.ScopeIR{ir}, Options{})
	if err != nil {
		t.Fatalf("resolve failed: %v", err)
	}
	counts := result.Graph.RelationshipCountsByType()
	if counts[graph.RelCalls] != 2 {
		t.Fatalf("expected same calledName duplicate to merge while alias call remains distinct, got counts=%v", counts)
	}
	if result.Metrics.DuplicateEdgesMerged != 1 || result.Metrics.UnresolvedReferences != 1 {
		t.Fatalf("unexpected duplicate/unresolved metrics: %#v", result.Metrics)
	}
	targetNode := requireNode(t, result.Graph, "Function", "src/app.ts", "target")
	var targetCall graph.Relationship
	for _, relationship := range result.Graph.Relationships {
		if relationship.Type == graph.RelCalls && relationship.TargetID == targetNode.ID && relationship.TargetText == "target" {
			targetCall = relationship
			break
		}
	}
	if targetCall.ID == "" {
		t.Fatalf("missing merged target call relationship")
	}
	if targetCall.SourceSiteStatus != sourceSiteStatusResolved ||
		targetCall.ProofKind != proofKindScopeBinding ||
		targetCall.TargetRole != targetRoleCallable ||
		targetCall.FilePath != "src/app.ts" ||
		targetCall.SourceSiteCount != 2 ||
		len(targetCall.SourceSiteIDs) != 2 {
		t.Fatalf("merged call source-site metadata = %#v", targetCall)
	}
}

func TestResolveAttachesSourceBackedUnresolvedDiagnostics(t *testing.T) {
	moduleScope := "scope:src/app.ts#1:0-4:1:Module"
	functionScope := "scope:src/app.ts#1:0-4:1:Function"
	functionDef := scopeir.DefinitionFact{
		ID:            "def:src/app.ts#1:0:Function:start",
		FilePath:      "src/app.ts",
		Name:          "start",
		Label:         scopeir.NodeFunction,
		QualifiedName: "start",
		Range:         scopeir.Range{StartLine: 1, EndLine: 4},
	}
	ir := scopeir.ScopeIR{
		FilePath:    "src/app.ts",
		FileHash:    "hash-app",
		Language:    scanner.TypeScript,
		ModuleScope: moduleScope,
		Scopes: []scopeir.ScopeFact{
			{ID: moduleScope, Kind: scopeir.ScopeModule, FilePath: "src/app.ts", Range: scopeir.Range{StartLine: 1, EndLine: 4}},
			{ID: functionScope, Parent: &[]string{moduleScope}[0], Kind: scopeir.ScopeFunction, FilePath: "src/app.ts", Range: scopeir.Range{StartLine: 1, EndLine: 4}, OwnedDefIDs: []string{functionDef.ID}, Bindings: []scopeir.BindingFact{{Name: "start", DefID: functionDef.ID, Origin: scopeir.BindingLocal}}},
		},
		Definitions: []scopeir.DefinitionFact{functionDef},
		Calls: []scopeir.CallSiteFact{
			{FilePath: "src/app.ts", FileHash: "hash-app", Name: "missing", InScope: functionScope, CallForm: scopeir.CallFree, Range: scopeir.Range{StartLine: 3, StartCol: 2, EndLine: 3}},
		},
	}

	result, err := Resolve([]scopeir.ScopeIR{ir}, Options{})
	if err != nil {
		t.Fatalf("resolve failed: %v", err)
	}
	source := requireNode(t, result.Graph, "Function", "src/app.ts", "start")
	diagnostics, ok := source.Properties[graphhealth.DiagnosticPropertyKey].([]graphhealth.Diagnostic)
	if !ok || len(diagnostics) != 1 {
		t.Fatalf("source diagnostics = %#v", source.Properties[graphhealth.DiagnosticPropertyKey])
	}
	diagnostic := diagnostics[0]
	if diagnostic.Kind != graphhealth.DiagnosticUnresolvedReference ||
		diagnostic.FactFamily != "call" ||
		diagnostic.SourceNodeID != source.ID ||
		diagnostic.TargetText != "missing" ||
		diagnostic.ResolutionSource != "scope-resolution" ||
		diagnostic.FilePath != "src/app.ts" ||
		diagnostic.StartLine != 3 ||
		diagnostic.StartCol != 2 ||
		diagnostic.SourceSiteID != "SourceSite:src/app.ts#call#missing#3#2#3#0" ||
		diagnostic.SourceSiteStatus != sourceSiteStatusUnresolvedLocalBinding ||
		diagnostic.ProofKind != proofKindNone ||
		diagnostic.TargetRole != targetRoleCallable ||
		diagnostic.Count != 1 {
		t.Fatalf("unexpected diagnostic: %#v", diagnostic)
	}
	if result.Metrics.UnresolvedReferences != 1 ||
		result.Metrics.UnresolvedReferenceDiagnostics != 1 ||
		result.Metrics.UnattributedUnresolvedReferences != 0 {
		t.Fatalf("unexpected unresolved metrics: %#v", result.Metrics)
	}
	resolutionMetadata, ok := result.Graph.Metadata["resolution"].(map[string]any)
	if !ok {
		t.Fatalf("missing resolution metadata: %#v", result.Graph.Metadata)
	}
	if resolutionMetadata["unresolvedReferences"] != 1 ||
		resolutionMetadata["sourceBackedUnresolvedReferences"] != 1 ||
		resolutionMetadata["unattributedUnresolvedReferences"] != 0 {
		t.Fatalf("unexpected resolution metadata: %#v", resolutionMetadata)
	}
}

func TestDefinitionLookupNamesDeduplicatesTrimmedNames(t *testing.T) {
	def := scopeir.DefinitionFact{
		Name:          " Service ",
		QualifiedName: "pkg.Service",
	}
	names := definitionLookupNames(def)
	want := []string{"Service", "pkg.Service"}
	if len(names) != len(want) {
		t.Fatalf("lookup name count = %d, want %d: %#v", len(names), len(want), names)
	}
	for index := range want {
		if names[index] != want[index] {
			t.Fatalf("lookup name %d = %q, want %q", index, names[index], want[index])
		}
	}
	if !definitionLookupNameMatches(def, "Service") || !definitionLookupNameMatches(def, "pkg.Service") {
		t.Fatalf("expected lookup-name matcher to find simple and qualified names")
	}
	if definitionLookupNameMatches(def, "Other") {
		t.Fatalf("unexpected lookup-name match")
	}
}

func TestUniqueDefsFastPathKeepsFirstCandidate(t *testing.T) {
	first := defRef{Fact: scopeir.DefinitionFact{ID: "same", Name: "first"}}
	duplicate := defRef{Fact: scopeir.DefinitionFact{ID: "same", Name: "duplicate"}}
	other := defRef{Fact: scopeir.DefinitionFact{ID: "other", Name: "other"}}

	one := uniqueDefs([]defRef{first, duplicate})
	if len(one) != 1 || one[0].Fact.Name != "first" {
		t.Fatalf("unique duplicate defs = %#v, want first only", one)
	}
	two := uniqueDefs([]defRef{first, other})
	if len(two) != 2 || two[0].Fact.Name != "first" || two[1].Fact.Name != "other" {
		t.Fatalf("unique distinct defs = %#v, want both in order", two)
	}
}

func TestResolveDoesNotEmitResolvedCallFromFileCaller(t *testing.T) {
	moduleScope := "scope:src/app.ts#1:0-3:1:Module"
	targetDef := scopeir.DefinitionFact{
		ID:            "def:src/app.ts#4:0:Function:target",
		FilePath:      "src/app.ts",
		Name:          "target",
		Label:         scopeir.NodeFunction,
		QualifiedName: "target",
		Range:         scopeir.Range{StartLine: 4, EndLine: 4},
	}
	otherTargetDef := scopeir.DefinitionFact{
		ID:            "def:src/other.ts#1:0:Function:target",
		FilePath:      "src/other.ts",
		Name:          "target",
		Label:         scopeir.NodeFunction,
		QualifiedName: "target",
		Range:         scopeir.Range{StartLine: 1, EndLine: 1},
	}
	ir := scopeir.ScopeIR{
		FilePath:    "src/app.ts",
		FileHash:    "hash-app",
		Language:    scanner.TypeScript,
		ModuleScope: moduleScope,
		Scopes: []scopeir.ScopeFact{
			{ID: moduleScope, Kind: scopeir.ScopeModule, FilePath: "src/app.ts", Range: scopeir.Range{StartLine: 1, EndLine: 4}},
		},
		Definitions: []scopeir.DefinitionFact{targetDef},
		Calls: []scopeir.CallSiteFact{
			{FilePath: "src/app.ts", FileHash: "hash-app", Name: "target", InScope: moduleScope, CallForm: scopeir.CallFree, Range: scopeir.Range{StartLine: 2, StartCol: 2, EndLine: 2}},
		},
	}
	otherIR := scopeir.ScopeIR{
		FilePath:    "src/other.ts",
		Language:    scanner.TypeScript,
		ModuleScope: "scope:src/other.ts#1:0-1:1:Module",
		Scopes: []scopeir.ScopeFact{
			{ID: "scope:src/other.ts#1:0-1:1:Module", Kind: scopeir.ScopeModule, FilePath: "src/other.ts", Range: scopeir.Range{StartLine: 1, EndLine: 1}},
		},
		Definitions: []scopeir.DefinitionFact{otherTargetDef},
	}

	result, err := Resolve([]scopeir.ScopeIR{ir, otherIR}, Options{})
	if err != nil {
		t.Fatalf("resolve failed: %v", err)
	}
	requireNoRelationship(t, result.Graph, graph.RelCalls, "File:src/app.ts", "Function:src/app.ts:target")
	fileNode, ok := result.Graph.GetNode("File:src/app.ts")
	if !ok {
		t.Fatalf("source file node missing")
	}
	diagnostics, ok := fileNode.Properties[graphhealth.DiagnosticPropertyKey].([]graphhealth.Diagnostic)
	if !ok || len(diagnostics) != 1 {
		t.Fatalf("file diagnostics = %#v", fileNode.Properties[graphhealth.DiagnosticPropertyKey])
	}
	diagnostic := diagnostics[0]
	if diagnostic.TargetText != "target" ||
		diagnostic.SourceSiteStatus != sourceSiteStatusUnsupportedSyntax ||
		p6c3DiagnosticReason(t, diagnostic) != unresolvedNoteCallSourceFileLevel {
		t.Fatalf("unexpected file-source call diagnostic: %#v", diagnostic)
	}
	if result.Metrics.ResolvedCalls != 0 || result.Metrics.UnresolvedReferences != 1 {
		t.Fatalf("unexpected metrics: %#v", result.Metrics)
	}
}

func TestResolveGoSamePackageDirectCallAcrossFilesBeforeGlobalAmbiguity(t *testing.T) {
	moduleScope := "scope:internal/pkg/app.go#1:0-5:1:Module"
	functionScope := "scope:internal/pkg/app.go#3:0-5:1:Function"
	runDef := scopeir.DefinitionFact{
		ID:            "def:internal/pkg/app.go#3:0:Function:Run",
		FilePath:      "internal/pkg/app.go",
		Name:          "Run",
		Label:         scopeir.NodeFunction,
		QualifiedName: "Run",
		Range:         scopeir.Range{StartLine: 3, EndLine: 5},
	}
	zeroParams := 0
	samePackageHelper := scopeir.DefinitionFact{
		ID:             "def:internal/pkg/helper.go#1:0:Function:helper",
		FilePath:       "internal/pkg/helper.go",
		Name:           "helper",
		Label:          scopeir.NodeFunction,
		QualifiedName:  "helper",
		Range:          scopeir.Range{StartLine: 1, EndLine: 1},
		ParameterCount: &zeroParams,
	}
	otherPackageHelper := scopeir.DefinitionFact{
		ID:            "def:internal/other/helper.go#1:0:Function:helper",
		FilePath:      "internal/other/helper.go",
		Name:          "helper",
		Label:         scopeir.NodeFunction,
		QualifiedName: "helper",
		Range:         scopeir.Range{StartLine: 1, EndLine: 1},
	}
	source := scopeir.ScopeIR{
		FilePath:    "internal/pkg/app.go",
		FileHash:    "hash-app",
		Language:    scanner.Go,
		ModuleScope: moduleScope,
		Scopes: []scopeir.ScopeFact{
			{ID: moduleScope, Kind: scopeir.ScopeModule, FilePath: "internal/pkg/app.go", Range: scopeir.Range{StartLine: 1, EndLine: 5}},
			{ID: functionScope, Parent: &[]string{moduleScope}[0], Kind: scopeir.ScopeFunction, FilePath: "internal/pkg/app.go", Range: scopeir.Range{StartLine: 3, EndLine: 5}, OwnedDefIDs: []string{runDef.ID}, Bindings: []scopeir.BindingFact{{Name: "Run", DefID: runDef.ID, Origin: scopeir.BindingLocal}}},
		},
		Definitions: []scopeir.DefinitionFact{runDef},
		Calls: []scopeir.CallSiteFact{
			{FilePath: "internal/pkg/app.go", FileHash: "hash-app", Name: "helper", InScope: functionScope, CallForm: scopeir.CallFree, Arity: intPtr(2), Range: scopeir.Range{StartLine: 4, StartCol: 2, EndLine: 4}},
		},
	}
	samePackage := scopeir.ScopeIR{
		FilePath: "internal/pkg/helper.go",
		Language: scanner.Go,
		Definitions: []scopeir.DefinitionFact{
			samePackageHelper,
		},
	}
	otherPackage := scopeir.ScopeIR{
		FilePath: "internal/other/helper.go",
		Language: scanner.Go,
		Definitions: []scopeir.DefinitionFact{
			otherPackageHelper,
		},
	}

	result, err := Resolve([]scopeir.ScopeIR{source, samePackage, otherPackage}, Options{})
	if err != nil {
		t.Fatalf("resolve failed: %v", err)
	}
	runID := requireDefinitionNodeID(t, result.Graph, runDef)
	samePackageHelperID := requireDefinitionNodeID(t, result.Graph, samePackageHelper)
	otherPackageHelperID := requireDefinitionNodeID(t, result.Graph, otherPackageHelper)
	relationship := requireRelationship(t, result.Graph, graph.RelCalls, runID, samePackageHelperID)
	if relationship.Confidence != 0.95 {
		t.Fatalf("same-package fallback confidence = %v, want 0.95", relationship.Confidence)
	}
	requireNoRelationship(t, result.Graph, graph.RelCalls, runID, otherPackageHelperID)
	if result.Metrics.ResolvedCalls != 1 || result.Metrics.UnresolvedReferences != 0 {
		t.Fatalf("unexpected metrics: %#v", result.Metrics)
	}
}

func TestResolveImportedPackageMemberCall(t *testing.T) {
	targetRaw := "github.com/tamnguyendinh/anvien/internal/pkg"
	moduleScope := "scope:cmd/app/main.go#1:0-5:1:Module"
	functionScope := "scope:cmd/app/main.go#3:0-5:1:Function"
	mainDef := scopeir.DefinitionFact{
		ID:            "def:cmd/app/main.go#3:0:Function:main",
		FilePath:      "cmd/app/main.go",
		Name:          "main",
		Label:         scopeir.NodeFunction,
		QualifiedName: "main",
		Range:         scopeir.Range{StartLine: 3, EndLine: 5},
	}
	helperDef := scopeir.DefinitionFact{
		ID:            "def:internal/pkg/helper.go#1:0:Function:Helper",
		FilePath:      "internal/pkg/helper.go",
		Name:          "Helper",
		Label:         scopeir.NodeFunction,
		QualifiedName: "Helper",
		Range:         scopeir.Range{StartLine: 1, EndLine: 1},
	}
	source := scopeir.ScopeIR{
		FilePath:    "cmd/app/main.go",
		FileHash:    "hash-main",
		Language:    scanner.Go,
		ModuleScope: moduleScope,
		Scopes: []scopeir.ScopeFact{
			{ID: moduleScope, Kind: scopeir.ScopeModule, FilePath: "cmd/app/main.go", Range: scopeir.Range{StartLine: 1, EndLine: 5}},
			{ID: functionScope, Parent: &[]string{moduleScope}[0], Kind: scopeir.ScopeFunction, FilePath: "cmd/app/main.go", Range: scopeir.Range{StartLine: 3, EndLine: 5}, OwnedDefIDs: []string{mainDef.ID}, Bindings: []scopeir.BindingFact{{Name: "main", DefID: mainDef.ID, Origin: scopeir.BindingLocal}}},
		},
		Definitions: []scopeir.DefinitionFact{mainDef},
		Imports: []scopeir.ImportFact{{
			FilePath:     "cmd/app/main.go",
			Kind:         scopeir.ImportNamed,
			LocalName:    "pkg",
			ImportedName: "pkg",
			TargetRaw:    &targetRaw,
		}},
		Calls: []scopeir.CallSiteFact{
			{FilePath: "cmd/app/main.go", FileHash: "hash-main", Name: "Helper", ExplicitReceiver: "pkg", InScope: functionScope, CallForm: scopeir.CallMember, Range: scopeir.Range{StartLine: 4, StartCol: 2, EndLine: 4}},
		},
	}
	target := scopeir.ScopeIR{
		FilePath:    "internal/pkg/helper.go",
		Language:    scanner.Go,
		ModuleScope: "scope:internal/pkg/helper.go#1:0-1:1:Module",
		Scopes: []scopeir.ScopeFact{
			{ID: "scope:internal/pkg/helper.go#1:0-1:1:Module", Kind: scopeir.ScopeModule, FilePath: "internal/pkg/helper.go", Range: scopeir.Range{StartLine: 1, EndLine: 1}},
		},
		Definitions: []scopeir.DefinitionFact{helperDef},
	}

	result, err := Resolve([]scopeir.ScopeIR{source, target}, Options{})
	if err != nil {
		t.Fatalf("resolve failed: %v", err)
	}
	mainID := requireDefinitionNodeID(t, result.Graph, mainDef)
	helperID := requireDefinitionNodeID(t, result.Graph, helperDef)
	relationship := requireRelationship(t, result.Graph, graph.RelCalls, mainID, helperID)
	if relationship.Confidence != 0.9 {
		t.Fatalf("imported package member confidence = %v, want 0.9", relationship.Confidence)
	}
	if result.Metrics.ResolvedCalls != 1 || result.Metrics.UnresolvedReferences != 0 {
		t.Fatalf("unexpected metrics: %#v", result.Metrics)
	}
}

func TestResolveMemberCallThroughImportedCallReturnBinding(t *testing.T) {
	targetRaw := "github.com/tamnguyendinh/anvien/internal/graph"
	moduleScope := "scope:cmd/app/main.go#1:0-5:1:Module"
	functionScope := "scope:cmd/app/main.go#2:0-5:1:Function"
	mainDef := scopeir.DefinitionFact{
		ID:             "def:cmd/app/main.go#2:0:Function:main",
		FilePath:       "cmd/app/main.go",
		Name:           "main",
		Label:          scopeir.NodeFunction,
		QualifiedName:  "main",
		Range:          scopeir.Range{StartLine: 2, EndLine: 5},
		ParameterCount: intPtr(0),
	}
	graphVar := scopeir.DefinitionFact{
		ID:            "def:cmd/app/main.go#3:2:Variable:g",
		FilePath:      "cmd/app/main.go",
		Name:          "g",
		Label:         scopeir.NodeVariable,
		QualifiedName: "g",
		Range:         scopeir.Range{StartLine: 3, StartCol: 2, EndLine: 3, EndCol: 18},
	}
	graphDef := scopeir.DefinitionFact{
		ID:            "def:internal/graph/types.go#1:0:Struct:Graph",
		FilePath:      "internal/graph/types.go",
		Name:          "Graph",
		Label:         scopeir.NodeStruct,
		QualifiedName: "Graph",
		Range:         scopeir.Range{StartLine: 1, EndLine: 1},
	}
	newDef := scopeir.DefinitionFact{
		ID:             "def:internal/graph/types.go#2:0:Function:New",
		FilePath:       "internal/graph/types.go",
		Name:           "New",
		Label:          scopeir.NodeFunction,
		QualifiedName:  "New",
		Range:          scopeir.Range{StartLine: 2, EndLine: 2},
		ParameterCount: intPtr(0),
		ReturnType:     "*Graph",
	}
	addNodeDef := scopeir.DefinitionFact{
		ID:             "def:internal/graph/types.go#3:0:Method:AddNode",
		FilePath:       "internal/graph/types.go",
		Name:           "AddNode",
		Label:          scopeir.NodeMethod,
		QualifiedName:  "Graph.AddNode",
		Range:          scopeir.Range{StartLine: 3, EndLine: 3},
		OwnerID:        graphDef.ID,
		ParameterCount: intPtr(1),
	}
	ir := scopeir.ScopeIR{
		FilePath:    "cmd/app/main.go",
		FileHash:    "hash-main",
		Language:    scanner.Go,
		ModuleScope: moduleScope,
		Scopes: []scopeir.ScopeFact{
			{ID: moduleScope, Kind: scopeir.ScopeModule, FilePath: "cmd/app/main.go", Range: scopeir.Range{StartLine: 1, EndLine: 5}},
			{ID: functionScope, Parent: &[]string{moduleScope}[0], Kind: scopeir.ScopeFunction, FilePath: "cmd/app/main.go", Range: scopeir.Range{StartLine: 2, EndLine: 5}, OwnedDefIDs: []string{mainDef.ID, graphVar.ID}, Bindings: []scopeir.BindingFact{{Name: "main", DefID: mainDef.ID, Origin: scopeir.BindingLocal}, {Name: "g", DefID: graphVar.ID, Origin: scopeir.BindingLocal}}},
		},
		Definitions: []scopeir.DefinitionFact{mainDef, graphVar},
		Imports: []scopeir.ImportFact{{
			FilePath:     "cmd/app/main.go",
			Kind:         scopeir.ImportNamed,
			LocalName:    "graph",
			ImportedName: "graph",
			TargetRaw:    &targetRaw,
		}},
		Calls: []scopeir.CallSiteFact{
			{FilePath: "cmd/app/main.go", FileHash: "hash-main", Name: "New", ExplicitReceiver: "graph", InScope: functionScope, CallForm: scopeir.CallMember, Arity: intPtr(0), Range: scopeir.Range{StartLine: 3, StartCol: 7, EndLine: 3, EndCol: 18}},
			{FilePath: "cmd/app/main.go", FileHash: "hash-main", Name: "AddNode", ExplicitReceiver: "g", InScope: functionScope, CallForm: scopeir.CallMember, Arity: intPtr(1), Range: scopeir.Range{StartLine: 4, StartCol: 2, EndLine: 4, EndCol: 16}},
		},
	}
	graphIR := scopeir.ScopeIR{
		FilePath: "internal/graph/types.go",
		Language: scanner.Go,
		Definitions: []scopeir.DefinitionFact{
			{ID: "def:internal/graph/types.go#1:0:Package:graph", FilePath: "internal/graph/types.go", Name: "graph", Label: scopeir.NodePackage, QualifiedName: "graph", Range: scopeir.Range{StartLine: 1, EndLine: 1}},
			graphDef,
			newDef,
			addNodeDef,
		},
	}

	result, err := Resolve([]scopeir.ScopeIR{ir, graphIR}, Options{})
	if err != nil {
		t.Fatalf("resolve failed: %v", err)
	}
	mainID := requireDefinitionNodeID(t, result.Graph, mainDef)
	newID := requireDefinitionNodeID(t, result.Graph, newDef)
	addNodeID := requireDefinitionNodeID(t, result.Graph, addNodeDef)
	requireRelationship(t, result.Graph, graph.RelCalls, mainID, newID)
	requireRelationship(t, result.Graph, graph.RelCalls, mainID, addNodeID)
	if result.Metrics.ResolvedCalls != 2 || result.Metrics.UnresolvedReferences != 0 {
		t.Fatalf("unexpected metrics: %#v", result.Metrics)
	}
}

func TestResolveGlobalCallFallbackDoesNotEmitResolvedEdge(t *testing.T) {
	moduleScope := "scope:src/app.ts#1:0-3:1:Module"
	functionScope := "scope:src/app.ts#1:0-3:1:Function"
	callerDef := scopeir.DefinitionFact{
		ID:            "def:src/app.ts#1:0:Function:start",
		FilePath:      "src/app.ts",
		Name:          "start",
		Label:         scopeir.NodeFunction,
		QualifiedName: "start",
		Range:         scopeir.Range{StartLine: 1, EndLine: 3},
	}
	oneParam := 1
	twoParams := 2
	targetOne := scopeir.DefinitionFact{
		ID:             "def:src/one.ts#1:0:Function:target",
		FilePath:       "src/one.ts",
		Name:           "target",
		Label:          scopeir.NodeFunction,
		QualifiedName:  "target",
		Range:          scopeir.Range{StartLine: 1, EndLine: 1},
		ParameterCount: &oneParam,
	}
	targetTwo := scopeir.DefinitionFact{
		ID:             "def:src/two.ts#1:0:Function:target",
		FilePath:       "src/two.ts",
		Name:           "target",
		Label:          scopeir.NodeFunction,
		QualifiedName:  "target",
		Range:          scopeir.Range{StartLine: 1, EndLine: 1},
		ParameterCount: &twoParams,
	}
	ir := scopeir.ScopeIR{
		FilePath:    "src/app.ts",
		FileHash:    "hash-app",
		Language:    scanner.TypeScript,
		ModuleScope: moduleScope,
		Scopes: []scopeir.ScopeFact{
			{ID: moduleScope, Kind: scopeir.ScopeModule, FilePath: "src/app.ts", Range: scopeir.Range{StartLine: 1, EndLine: 3}},
			{ID: functionScope, Parent: &[]string{moduleScope}[0], Kind: scopeir.ScopeFunction, FilePath: "src/app.ts", Range: scopeir.Range{StartLine: 1, EndLine: 3}, OwnedDefIDs: []string{callerDef.ID}, Bindings: []scopeir.BindingFact{{Name: "start", DefID: callerDef.ID, Origin: scopeir.BindingLocal}}},
		},
		Definitions: []scopeir.DefinitionFact{callerDef},
		Calls: []scopeir.CallSiteFact{
			{FilePath: "src/app.ts", FileHash: "hash-app", Name: "target", InScope: functionScope, CallForm: scopeir.CallFree, Arity: &oneParam, Range: scopeir.Range{StartLine: 2, StartCol: 2, EndLine: 2}},
		},
	}
	oneIR := scopeir.ScopeIR{FilePath: "src/one.ts", Language: scanner.TypeScript, Definitions: []scopeir.DefinitionFact{targetOne}}
	twoIR := scopeir.ScopeIR{FilePath: "src/two.ts", Language: scanner.TypeScript, Definitions: []scopeir.DefinitionFact{targetTwo}}

	result, err := Resolve([]scopeir.ScopeIR{ir, oneIR, twoIR}, Options{})
	if err != nil {
		t.Fatalf("resolve failed: %v", err)
	}
	requireNoRelationship(t, result.Graph, graph.RelCalls, "Function:src/app.ts:start", "Function:src/one.ts:target#1")
	requireNoRelationship(t, result.Graph, graph.RelCalls, "Function:src/app.ts:start", "Function:src/two.ts:target#2")
	source := requireNode(t, result.Graph, "Function", "src/app.ts", "start")
	diagnostics, ok := source.Properties[graphhealth.DiagnosticPropertyKey].([]graphhealth.Diagnostic)
	if !ok || len(diagnostics) != 1 {
		t.Fatalf("source diagnostics = %#v", source.Properties[graphhealth.DiagnosticPropertyKey])
	}
	diagnostic := diagnostics[0]
	if diagnostic.Kind != graphhealth.DiagnosticUnresolvedReference ||
		diagnostic.FactFamily != "call" ||
		diagnostic.SourceNodeID != source.ID ||
		diagnostic.TargetText != "target" ||
		p6c3DiagnosticReason(t, diagnostic) != "call target matched low-confidence global fallback only" ||
		diagnostic.Count != 1 {
		t.Fatalf("unexpected diagnostic: %#v", diagnostic)
	}
	if result.Metrics.ResolvedCalls != 0 || result.Metrics.UnresolvedReferences != 1 {
		t.Fatalf("unexpected metrics: %#v", result.Metrics)
	}
}

func TestResolveBareGoCallDoesNotFallbackToCrossLanguageMethod(t *testing.T) {
	moduleScope := "scope:backend/cmd/rms-backend/main.go#1:0-8:1:Module"
	functionScope := "scope:backend/cmd/rms-backend/main.go#3:0-8:1:Function"
	mainDef := scopeir.DefinitionFact{
		ID:            "def:backend/cmd/rms-backend/main.go#3:0:Function:main",
		FilePath:      "backend/cmd/rms-backend/main.go",
		Name:          "main",
		Label:         scopeir.NodeFunction,
		QualifiedName: "main",
		Range:         scopeir.Range{StartLine: 3, EndLine: 8},
	}
	goIR := scopeir.ScopeIR{
		FilePath:    "backend/cmd/rms-backend/main.go",
		FileHash:    "hash-main",
		Language:    scanner.Go,
		ModuleScope: moduleScope,
		Scopes: []scopeir.ScopeFact{
			{ID: moduleScope, Kind: scopeir.ScopeModule, FilePath: "backend/cmd/rms-backend/main.go", Range: scopeir.Range{StartLine: 1, EndLine: 8}},
			{ID: functionScope, Parent: &[]string{moduleScope}[0], Kind: scopeir.ScopeFunction, FilePath: "backend/cmd/rms-backend/main.go", Range: scopeir.Range{StartLine: 3, EndLine: 8}, OwnedDefIDs: []string{mainDef.ID}, Bindings: []scopeir.BindingFact{{Name: "main", DefID: mainDef.ID, Origin: scopeir.BindingLocal}}},
		},
		Definitions: []scopeir.DefinitionFact{mainDef},
		Calls: []scopeir.CallSiteFact{
			{FilePath: "backend/cmd/rms-backend/main.go", FileHash: "hash-main", Name: "stop", InScope: functionScope, CallForm: scopeir.CallFree, Range: scopeir.Range{StartLine: 6, StartCol: 2, EndLine: 6}},
		},
	}
	listenerClass := scopeir.DefinitionFact{
		ID:            "def:electron/main/sync/sse-listener.ts#1:0:Class:SSEListener",
		FilePath:      "electron/main/sync/sse-listener.ts",
		Name:          "SSEListener",
		Label:         scopeir.NodeClass,
		QualifiedName: "SSEListener",
		Range:         scopeir.Range{StartLine: 1, EndLine: 8},
	}
	stopMethod := scopeir.DefinitionFact{
		ID:             "def:electron/main/sync/sse-listener.ts#4:2:Method:SSEListener.stop",
		FilePath:       "electron/main/sync/sse-listener.ts",
		Name:           "stop",
		Label:          scopeir.NodeMethod,
		QualifiedName:  "SSEListener.stop",
		OwnerID:        listenerClass.ID,
		Range:          scopeir.Range{StartLine: 4, StartCol: 2, EndLine: 6, EndCol: 3},
		ParameterCount: intPtr(0),
	}
	tsIR := scopeir.ScopeIR{
		FilePath: "electron/main/sync/sse-listener.ts",
		Language: scanner.TypeScript,
		Definitions: []scopeir.DefinitionFact{
			listenerClass,
			stopMethod,
		},
	}

	result, err := Resolve([]scopeir.ScopeIR{goIR, tsIR}, Options{})
	if err != nil {
		t.Fatalf("resolve failed: %v", err)
	}
	mainNode := requireNode(t, result.Graph, "Function", "backend/cmd/rms-backend/main.go", "main")
	stopNode := requireNode(t, result.Graph, "Method", "electron/main/sync/sse-listener.ts", "SSEListener.stop")
	requireNoRelationship(t, result.Graph, graph.RelCalls, mainNode.ID, stopNode.ID)
	diagnostics, ok := mainNode.Properties[graphhealth.DiagnosticPropertyKey].([]graphhealth.Diagnostic)
	if !ok || len(diagnostics) != 1 {
		t.Fatalf("source diagnostics = %#v", mainNode.Properties[graphhealth.DiagnosticPropertyKey])
	}
	diagnostic := diagnostics[0]
	if diagnostic.FactFamily != "call" ||
		diagnostic.TargetText != "stop" ||
		p6c3DiagnosticReason(t, diagnostic) != "call target matched low-confidence global fallback only" ||
		diagnostic.SourceNodeID != mainNode.ID ||
		diagnostic.Count != 1 {
		t.Fatalf("unexpected diagnostic: %#v", diagnostic)
	}
	if result.Metrics.ResolvedCalls != 0 || result.Metrics.UnresolvedReferences != 1 {
		t.Fatalf("unexpected metrics: %#v", result.Metrics)
	}
}

func BenchmarkResolveImportedMemberManyImports(b *testing.B) {
	const importCount = 1200
	const sourceFile = "src/main.ts"
	moduleScope := "scope:src/main.ts#1:0-4:1:Module"
	functionScope := "scope:src/main.ts#2:0-4:1:Function"
	mainDef := scopeir.DefinitionFact{
		ID:            "def:src/main.ts#2:0:Function:main",
		FilePath:      sourceFile,
		Name:          "main",
		Label:         scopeir.NodeFunction,
		QualifiedName: "main",
		Range:         scopeir.Range{StartLine: 2, EndLine: 4},
	}
	source := scopeir.ScopeIR{
		FilePath:    sourceFile,
		Language:    scanner.TypeScript,
		ModuleScope: moduleScope,
		Scopes: []scopeir.ScopeFact{
			{ID: moduleScope, Kind: scopeir.ScopeModule, FilePath: sourceFile, Range: scopeir.Range{StartLine: 1, EndLine: 4}},
			{ID: functionScope, Parent: &[]string{moduleScope}[0], Kind: scopeir.ScopeFunction, FilePath: sourceFile, Range: scopeir.Range{StartLine: 2, EndLine: 4}, OwnedDefIDs: []string{mainDef.ID}, Bindings: []scopeir.BindingFact{{Name: "main", DefID: mainDef.ID, Origin: scopeir.BindingLocal}}},
		},
		Definitions: []scopeir.DefinitionFact{mainDef},
		Imports:     make([]scopeir.ImportFact, 0, importCount),
	}
	files := make([]scopeir.ScopeIR, 0, importCount+1)
	files = append(files, source)
	for index := 0; index < importCount; index++ {
		receiver := "pkg" + intString(index)
		targetRaw := "./" + receiver
		targetFile := "src/" + receiver + ".ts"
		files[0].Imports = append(files[0].Imports, scopeir.ImportFact{
			FilePath:     sourceFile,
			Kind:         scopeir.ImportNamespace,
			LocalName:    receiver,
			ImportedName: receiver,
			TargetRaw:    &targetRaw,
		})
		files = append(files, scopeir.ScopeIR{
			FilePath: targetFile,
			Language: scanner.TypeScript,
			Definitions: []scopeir.DefinitionFact{{
				ID:            "def:" + targetFile + "#1:0:Function:Helper",
				FilePath:      targetFile,
				Name:          "Helper",
				Label:         scopeir.NodeFunction,
				QualifiedName: "Helper",
				Range:         scopeir.Range{StartLine: 1, EndLine: 1},
			}},
		})
	}

	w, err := buildWorkspace(files)
	if err != nil {
		b.Fatalf("buildWorkspace failed: %v", err)
	}
	receiver := "pkg" + intString(importCount-1)
	labels := callableLabels()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		target, ok := w.resolveImportedMember(receiver, "Helper", functionScope, labels)
		if !ok || target.Fact.Name != "Helper" {
			b.Fatalf("resolveImportedMember() = %#v, %v", target, ok)
		}
	}
}

type graphSignature struct {
	Nodes         []string `json:"nodes"`
	Relationships []string `json:"relationships"`
}

func buildGraphSignature(g *graph.Graph) graphSignature {
	signature := graphSignature{}
	// Graph IDs are occurrence identities. This golden keeps asserting semantic
	// topology; identity injectivity is covered by identity_occurrence_test.go.
	semanticIDs := make(map[string]string, len(g.Nodes))
	for _, node := range g.Nodes {
		semanticID := node.ID
		filePath, _ := node.Properties["filePath"].(string)
		qualifiedName, _ := node.Properties["qualifiedName"].(string)
		if qualifiedName == "" {
			qualifiedName, _ = node.Properties["name"].(string)
		}
		if node.Label != scopeir.NodeFile && filePath != "" && qualifiedName != "" {
			semanticID = graph.GenerateID(string(node.Label), cleanPath(filePath)+":"+qualifiedName)
		}
		semanticIDs[node.ID] = semanticID
		signature.Nodes = append(signature.Nodes, string(node.Label)+":"+semanticID)
	}
	for _, relationship := range g.SortedRelationships() {
		sourceID := semanticIDs[relationship.SourceID]
		if sourceID == "" {
			sourceID = relationship.SourceID
		}
		targetID := semanticIDs[relationship.TargetID]
		if targetID == "" {
			targetID = relationship.TargetID
		}
		signature.Relationships = append(signature.Relationships,
			string(relationship.Type)+":"+sourceID+"->"+targetID+":"+relationship.Reason,
		)
	}
	sort.Strings(signature.Nodes)
	sort.Strings(signature.Relationships)
	return signature
}

func intPtr(value int) *int {
	return &value
}

func requireNode(t *testing.T, g *graph.Graph, label string, filePath string, qualifiedName string) graph.Node {
	t.Helper()
	for _, node := range g.Nodes {
		if string(node.Label) != label {
			continue
		}
		if node.Properties["filePath"] == filePath && node.Properties["qualifiedName"] == qualifiedName {
			return node
		}
	}
	t.Fatalf("missing node %s %s %s", label, filePath, qualifiedName)
	return graph.Node{}
}

func requireDefinitionNodeID(t *testing.T, g *graph.Graph, definition scopeir.DefinitionFact) string {
	t.Helper()
	qualifiedName := definition.QualifiedName
	if qualifiedName == "" {
		qualifiedName = definition.Name
	}
	return requireNode(t, g, string(definition.Label), cleanPath(definition.FilePath), qualifiedName).ID
}

func requireRelationship(t *testing.T, g *graph.Graph, relType graph.RelationshipType, sourceID string, targetID string) graph.Relationship {
	t.Helper()
	for _, relationship := range g.Relationships {
		if relationship.Type == relType && relationship.SourceID == sourceID && relationship.TargetID == targetID {
			return relationship
		}
	}
	t.Fatalf("missing relationship %s %s -> %s", relType, sourceID, targetID)
	return graph.Relationship{}
}

func requireNoRelationship(t *testing.T, g *graph.Graph, relType graph.RelationshipType, sourceID string, targetID string) {
	t.Helper()
	for _, relationship := range g.Relationships {
		if relationship.Type == relType && relationship.SourceID == sourceID && relationship.TargetID == targetID {
			t.Fatalf("unexpected relationship %s %s -> %s", relType, sourceID, targetID)
		}
	}
}

func TestP5DResolveEmitsCallAccessProofAndConservesCoalescedSites(t *testing.T) {
	run := p5cDefinition("src/impl.ts", "run", scopeir.NodeFunction)
	value := p5cDefinition("src/impl.ts", "value", scopeir.NodeProperty)
	impl := p5cModule("src/impl.ts", []scopeir.DefinitionFact{run, value}, []scopeir.ExportFact{
		p5cLocalExport(run, scopeir.ExportDirect, "run", []scopeir.ExportMeaning{scopeir.ExportMeaningValue}),
		p5cLocalExport(value, scopeir.ExportDirect, "value", []scopeir.ExportMeaning{scopeir.ExportMeaningValue}),
	})
	barrel := p5cModule("src/barrel.ts", nil, []scopeir.ExportFact{
		p5cNamespace("src/barrel.ts", "api", "./impl", []scopeir.ExportMeaning{scopeir.ExportMeaningNamespace}, false),
	})
	consumer := p5cConsumer("src/app.ts", scanner.TypeScript, []scopeir.ImportFact{
		p5cImport(scopeir.ImportNamed, "api", "api", "./barrel", p5cAllMeanings, false),
	}, []scopeir.CallSiteFact{
		{Name: "run", ExplicitReceiver: "api", CallForm: scopeir.CallMember},
		{Name: "run", ExplicitReceiver: "api", CallForm: scopeir.CallMember},
	})
	consumer.Accesses = []scopeir.AccessFact{{
		FilePath:         "src/app.ts",
		FileHash:         "hash-src/app.ts",
		Name:             "value",
		Kind:             scopeir.AccessRead,
		Range:            scopeir.Range{StartLine: 6, EndLine: 6},
		InScope:          p5cFunctionScope("src/app.ts"),
		ExplicitReceiver: "api",
	}}

	result, err := Resolve([]scopeir.ScopeIR{consumer, barrel, impl}, Options{})
	if err != nil {
		t.Fatalf("Resolve() error = %v", err)
	}
	caller := requireNode(t, result.Graph, "Function", "src/app.ts", "caller")
	runNode := requireNode(t, result.Graph, "Function", "src/impl.ts", "run")
	valueNode := requireNode(t, result.Graph, "Property", "src/impl.ts", "value")

	call := requireRelationship(t, result.Graph, graph.RelCalls, caller.ID, runNode.ID)
	if call.ProofKind != proofKindImportMember ||
		call.SourceSiteCount != 2 ||
		len(call.SourceSiteIDs) != 2 ||
		len(call.Evidence) == 0 ||
		call.Evidence[0].Kind != "type-binding" ||
		call.Evidence[0].Note != "run" ||
		relationshipCallName(call) != "run" {
		t.Fatalf("coalesced CALLS contract = %#v", call)
	}
	if terminals := p5dEvidenceOfKind(call.Evidence, exportBindingTerminalEvidenceKind); len(terminals) != 2 {
		t.Fatalf("coalesced terminal proofs = %d, want 2 source sites: %#v", len(terminals), call.Evidence)
	}
	if hops := p5dEvidenceOfKind(call.Evidence, exportBindingHopEvidenceKind); len(hops) != 4 {
		t.Fatalf("coalesced hop proofs = %d, want 2 paths x 2 hops: %#v", len(hops), call.Evidence)
	}
	if failures := p5dEvidenceOfKind(call.Evidence, exportBindingFailureEvidenceKind); len(failures) != 0 {
		t.Fatalf("resolved namespace call gained failures: %#v", failures)
	}

	access := requireRelationship(t, result.Graph, graph.RelAccesses, caller.ID, valueNode.ID)
	if access.ProofKind != proofKindImportMember ||
		access.SourceSiteCount != 1 ||
		len(access.Evidence) == 0 ||
		access.Evidence[0].Kind != "import-binding" ||
		access.Evidence[0].Note != "api.value" {
		t.Fatalf("ACCESSES contract = %#v", access)
	}
	if terminals := p5dEvidenceOfKind(access.Evidence, exportBindingTerminalEvidenceKind); len(terminals) != 1 {
		t.Fatalf("access terminal proofs = %d, want 1: %#v", len(terminals), access.Evidence)
	}
	if hops := p5dEvidenceOfKind(access.Evidence, exportBindingHopEvidenceKind); len(hops) != 2 {
		t.Fatalf("access hop proofs = %d, want namespace + direct: %#v", len(hops), access.Evidence)
	}
}
