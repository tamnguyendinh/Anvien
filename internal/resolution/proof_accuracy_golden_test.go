package resolution

import (
	"testing"

	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/graphhealth"
	"github.com/tamnguyendinh/anvien/internal/scanner"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
)

func TestProofBasedCallAccessGoldenCorpus(t *testing.T) {
	const (
		sourceFile    = "src/golden.ts"
		sourceHash    = "hash-golden"
		moduleScope   = "scope:src/golden.ts#1:0-30:1:Module"
		functionScope = "scope:src/golden.ts#10:0-25:1:Function"
		apiFile       = "src/api.ts"
		listenerFile  = "electron/main/sync/sse-listener.ts"
	)
	apiRaw := "./api"

	runDef := scopeir.DefinitionFact{ID: "def:run", FilePath: sourceFile, FileHash: sourceHash, Name: "run", Label: scopeir.NodeFunction, QualifiedName: "run", Range: scopeir.Range{StartLine: 10, EndLine: 25}}
	helperDef := scopeir.DefinitionFact{ID: "def:helper", FilePath: sourceFile, FileHash: sourceHash, Name: "helper", Label: scopeir.NodeFunction, QualifiedName: "helper", Range: scopeir.Range{StartLine: 2, EndLine: 2}}
	callbackDef := scopeir.DefinitionFact{ID: "def:callback", FilePath: sourceFile, FileHash: sourceHash, Name: "callback", Label: scopeir.NodeVariable, QualifiedName: "callback", Range: scopeir.Range{StartLine: 3, EndLine: 3}}
	closureDef := scopeir.DefinitionFact{ID: "def:closure", FilePath: sourceFile, FileHash: sourceHash, Name: "closure", Label: scopeir.NodeFunction, QualifiedName: "run.closure", Range: scopeir.Range{StartLine: 9, EndLine: 9}}
	userClass := scopeir.DefinitionFact{ID: "def:User", FilePath: sourceFile, FileHash: sourceHash, Name: "User", Label: scopeir.NodeClass, QualifiedName: "User", Range: scopeir.Range{StartLine: 4, EndLine: 7}}
	userSave := scopeir.DefinitionFact{ID: "def:User.save", FilePath: sourceFile, FileHash: sourceHash, Name: "save", Label: scopeir.NodeMethod, QualifiedName: "User.save", OwnerID: userClass.ID, Range: scopeir.Range{StartLine: 5, EndLine: 5}}
	userID := scopeir.DefinitionFact{ID: "def:User.id", FilePath: sourceFile, FileHash: sourceHash, Name: "id", Label: scopeir.NodeProperty, QualifiedName: "User.id", OwnerID: userClass.ID, Range: scopeir.Range{StartLine: 6, EndLine: 6}}
	configClass := scopeir.DefinitionFact{ID: "def:Config", FilePath: sourceFile, FileHash: sourceHash, Name: "Config", Label: scopeir.NodeClass, QualifiedName: "Config", Range: scopeir.Range{StartLine: 7, EndLine: 8}}
	configMake := scopeir.DefinitionFact{ID: "def:Config.make", FilePath: sourceFile, FileHash: sourceHash, Name: "make", Label: scopeir.NodeFunction, QualifiedName: "Config.make", OwnerID: configClass.ID, Range: scopeir.Range{StartLine: 8, EndLine: 8}}
	apiFetch := scopeir.DefinitionFact{ID: "def:api.fetchUser", FilePath: apiFile, FileHash: "hash-api", Name: "fetchUser", Label: scopeir.NodeFunction, QualifiedName: "fetchUser", Range: scopeir.Range{StartLine: 1, EndLine: 1}}
	listenerClass := scopeir.DefinitionFact{ID: "def:SSEListener", FilePath: listenerFile, FileHash: "hash-listener", Name: "SSEListener", Label: scopeir.NodeClass, QualifiedName: "SSEListener", Range: scopeir.Range{StartLine: 1, EndLine: 3}}
	listenerStop := scopeir.DefinitionFact{ID: "def:SSEListener.stop", FilePath: listenerFile, FileHash: "hash-listener", Name: "stop", Label: scopeir.NodeMethod, QualifiedName: "SSEListener.stop", OwnerID: listenerClass.ID, Range: scopeir.Range{StartLine: 2, EndLine: 2}}

	helperFirst := scopeir.Range{StartLine: 11, StartCol: 2, EndLine: 11, EndCol: 10}
	helperSecond := scopeir.Range{StartLine: 12, StartCol: 2, EndLine: 12, EndCol: 10}
	saveCall := scopeir.Range{StartLine: 13, StartCol: 2, EndLine: 13, EndCol: 13}
	callbackCall := scopeir.Range{StartLine: 14, StartCol: 2, EndLine: 14, EndCol: 12}
	stopCall := scopeir.Range{StartLine: 15, StartCol: 2, EndLine: 15, EndCol: 8}
	closureCall := scopeir.Range{StartLine: 16, StartCol: 2, EndLine: 16, EndCol: 11}
	importedCall := scopeir.Range{StartLine: 17, StartCol: 2, EndLine: 17, EndCol: 17}
	builtinCall := scopeir.Range{StartLine: 18, StartCol: 2, EndLine: 18, EndCol: 8}
	externalCall := scopeir.Range{StartLine: 19, StartCol: 2, EndLine: 19, EndCol: 15}
	fileLevelCall := scopeir.Range{StartLine: 20, StartCol: 0, EndLine: 20, EndCol: 8}
	idAccess := scopeir.Range{StartLine: 21, StartCol: 2, EndLine: 21, EndCol: 9}
	functionSelectorAccess := scopeir.Range{StartLine: 22, StartCol: 2, EndLine: 22, EndCol: 13}

	ir := scopeir.ScopeIR{
		FilePath:    sourceFile,
		FileHash:    sourceHash,
		Language:    scanner.TypeScript,
		ModuleScope: moduleScope,
		Scopes: []scopeir.ScopeFact{
			{ID: moduleScope, Kind: scopeir.ScopeModule, FilePath: sourceFile, FileHash: sourceHash, Range: scopeir.Range{StartLine: 1, EndLine: 30}},
			{
				ID:       functionScope,
				Parent:   &[]string{moduleScope}[0],
				Kind:     scopeir.ScopeFunction,
				FilePath: sourceFile,
				FileHash: sourceHash,
				Range:    scopeir.Range{StartLine: 10, EndLine: 25},
				OwnedDefIDs: []string{
					runDef.ID,
					callbackDef.ID,
				},
				Bindings: []scopeir.BindingFact{
					{Name: "run", DefID: runDef.ID, Origin: scopeir.BindingLocal},
					{Name: "helper", DefID: helperDef.ID, Origin: scopeir.BindingLocal},
					{Name: "callback", DefID: callbackDef.ID, Origin: scopeir.BindingLocal},
					{Name: "closure", DefID: closureDef.ID, Origin: scopeir.BindingLocal},
				},
				TypeBindings: []scopeir.TypeBindingFact{
					{Name: "user", Type: scopeir.TypeRef{RawName: "User", DeclaredAtScope: functionScope, Source: scopeir.TypeSourceAssignment}},
					{Name: "config", Type: scopeir.TypeRef{RawName: "Config", DeclaredAtScope: functionScope, Source: scopeir.TypeSourceAssignment}},
				},
			},
		},
		Definitions: []scopeir.DefinitionFact{
			runDef,
			helperDef,
			callbackDef,
			closureDef,
			userClass,
			userSave,
			userID,
			configClass,
			configMake,
		},
		Imports: []scopeir.ImportFact{{
			FilePath:     sourceFile,
			FileHash:     sourceHash,
			Kind:         scopeir.ImportNamespace,
			LocalName:    "api",
			ImportedName: "api",
			TargetRaw:    &apiRaw,
		}},
		Calls: []scopeir.CallSiteFact{
			{FilePath: sourceFile, FileHash: sourceHash, Name: "helper", InScope: functionScope, CallForm: scopeir.CallFree, Range: helperFirst},
			{FilePath: sourceFile, FileHash: sourceHash, Name: "helper", InScope: functionScope, CallForm: scopeir.CallFree, Range: helperSecond},
			{FilePath: sourceFile, FileHash: sourceHash, Name: "save", ExplicitReceiver: "user", InScope: functionScope, CallForm: scopeir.CallMember, Range: saveCall},
			{FilePath: sourceFile, FileHash: sourceHash, Name: "callback", InScope: functionScope, CallForm: scopeir.CallFree, Range: callbackCall},
			{FilePath: sourceFile, FileHash: sourceHash, Name: "stop", InScope: functionScope, CallForm: scopeir.CallFree, Range: stopCall},
			{FilePath: sourceFile, FileHash: sourceHash, Name: "closure", InScope: functionScope, CallForm: scopeir.CallFree, Range: closureCall},
			{FilePath: sourceFile, FileHash: sourceHash, Name: "fetchUser", ExplicitReceiver: "api", InScope: functionScope, CallForm: scopeir.CallMember, Range: importedCall},
			{FilePath: sourceFile, FileHash: sourceHash, Name: "len", InScope: functionScope, CallForm: scopeir.CallFree, Range: builtinCall},
			{FilePath: sourceFile, FileHash: sourceHash, Name: "Command", ExplicitReceiver: "cobra", InScope: functionScope, CallForm: scopeir.CallMember, Range: externalCall},
			{FilePath: sourceFile, FileHash: sourceHash, Name: "helper", InScope: moduleScope, CallForm: scopeir.CallFree, Range: fileLevelCall},
		},
		Accesses: []scopeir.AccessFact{
			{FilePath: sourceFile, FileHash: sourceHash, Name: "id", ExplicitReceiver: "user", InScope: functionScope, Kind: scopeir.AccessRead, Range: idAccess},
			{FilePath: sourceFile, FileHash: sourceHash, Name: "make", ExplicitReceiver: "config", InScope: functionScope, Kind: scopeir.AccessRead, Range: functionSelectorAccess},
		},
	}
	apiIR := scopeir.ScopeIR{
		FilePath:    apiFile,
		FileHash:    "hash-api",
		Language:    scanner.TypeScript,
		ModuleScope: "scope:src/api.ts#1:0-1:1:Module",
		Scopes:      []scopeir.ScopeFact{{ID: "scope:src/api.ts#1:0-1:1:Module", Kind: scopeir.ScopeModule, FilePath: apiFile, FileHash: "hash-api", Range: scopeir.Range{StartLine: 1, EndLine: 1}}},
		Definitions: []scopeir.DefinitionFact{apiFetch},
	}
	listenerIR := scopeir.ScopeIR{
		FilePath:    listenerFile,
		FileHash:    "hash-listener",
		Language:    scanner.TypeScript,
		ModuleScope: "scope:electron/main/sync/sse-listener.ts#1:0-3:1:Module",
		Scopes: []scopeir.ScopeFact{
			{ID: "scope:electron/main/sync/sse-listener.ts#1:0-3:1:Module", Kind: scopeir.ScopeModule, FilePath: listenerFile, FileHash: "hash-listener", Range: scopeir.Range{StartLine: 1, EndLine: 3}, OwnedDefIDs: []string{listenerClass.ID}},
		},
		Definitions: []scopeir.DefinitionFact{listenerClass, listenerStop},
	}

	result, err := Resolve([]scopeir.ScopeIR{ir, apiIR, listenerIR}, Options{})
	if err != nil {
		t.Fatalf("resolve failed: %v", err)
	}

	run := requireNode(t, result.Graph, "Function", sourceFile, "run")
	helper := requireNode(t, result.Graph, "Function", sourceFile, "helper")
	closure := requireNode(t, result.Graph, "Function", sourceFile, "run.closure")
	save := requireNode(t, result.Graph, "Method", sourceFile, "User.save")
	id := requireNode(t, result.Graph, "Property", sourceFile, "User.id")
	callback := requireNode(t, result.Graph, "Variable", sourceFile, "callback")
	configMakeNode := requireNode(t, result.Graph, "Function", sourceFile, "Config.make")
	fetchUser := requireNode(t, result.Graph, "Function", apiFile, "fetchUser")
	stop := requireNode(t, result.Graph, "Method", listenerFile, "SSEListener.stop")

	helperCall := requireRelationship(t, result.Graph, graph.RelCalls, run.ID, helper.ID)
	requireGoldenRelationshipMetadata(t, helperCall, "helper", proofKindScopeBinding, targetRoleCallable, 2)
	requireGoldenSourceSiteIDs(t, helperCall,
		sourceSiteID("call", sourceFile, "helper", helperFirst),
		sourceSiteID("call", sourceFile, "helper", helperSecond),
	)
	requireGoldenRelationshipMetadata(t, requireRelationship(t, result.Graph, graph.RelCalls, run.ID, save.ID), "user.save", proofKindReceiverMember, targetRoleCallable, 1)
	requireGoldenRelationshipMetadata(t, requireRelationship(t, result.Graph, graph.RelCalls, run.ID, closure.ID), "closure", proofKindScopeBinding, targetRoleCallable, 1)
	requireGoldenRelationshipMetadata(t, requireRelationship(t, result.Graph, graph.RelCalls, run.ID, fetchUser.ID), "api.fetchUser", proofKindImportMember, targetRoleCallable, 1)
	requireGoldenRelationshipMetadata(t, requireRelationship(t, result.Graph, graph.RelAccesses, run.ID, id.ID), "user.id", proofKindReceiverMember, targetRoleMember, 1)

	requireNoRelationship(t, result.Graph, graph.RelCalls, run.ID, callback.ID)
	requireNoRelationship(t, result.Graph, graph.RelCalls, run.ID, stop.ID)
	requireNoRelationship(t, result.Graph, graph.RelCalls, "File:"+sourceFile, helper.ID)
	requireNoRelationship(t, result.Graph, graph.RelAccesses, run.ID, configMakeNode.ID)

	runDiagnostics := requireGoldenDiagnostics(t, run)
	requireGoldenDiagnostic(t, runDiagnostics, "call", "callback", sourceSiteStatusUnresolvedLocalBinding, proofKindNone, graphhealth.DiagnosticClassificationInRepoUnresolved, graphhealth.DiagnosticActionabilityAnalyzerGap)
	requireGoldenDiagnostic(t, runDiagnostics, "call", "stop", sourceSiteStatusUnresolvedLocalBinding, proofKindGlobalFallbackLowConfidence, graphhealth.DiagnosticClassificationInRepoUnresolved, graphhealth.DiagnosticActionabilityAnalyzerGap)
	requireGoldenDiagnostic(t, runDiagnostics, "call", "len", sourceSiteStatusUnresolvedLocalBinding, proofKindNone, graphhealth.DiagnosticClassificationBuiltin, graphhealth.DiagnosticActionabilityNonActionable)
	requireGoldenDiagnostic(t, runDiagnostics, "call", "cobra.Command", sourceSiteStatusUnresolvedLocalBinding, proofKindNone, graphhealth.DiagnosticClassificationExternalLibrary, graphhealth.DiagnosticActionabilityReview)
	requireGoldenDiagnostic(t, runDiagnostics, "access", "config.make", sourceSiteStatusUnresolvedLocalBinding, proofKindNone, graphhealth.DiagnosticClassificationInRepoUnresolved, graphhealth.DiagnosticActionabilityAnalyzerGap)

	fileNode, ok := result.Graph.GetNode("File:" + sourceFile)
	if !ok {
		t.Fatalf("missing source file node")
	}
	fileDiagnostics := requireGoldenDiagnostics(t, fileNode)
	requireGoldenDiagnostic(t, fileDiagnostics, "call", "helper", sourceSiteStatusUnsupportedSyntax, proofKindNone, graphhealth.DiagnosticClassificationInRepoUnresolved, graphhealth.DiagnosticActionabilityAnalyzerGap)

	wantSourceSites := []string{
		sourceSiteID("call", sourceFile, "helper", helperFirst),
		sourceSiteID("call", sourceFile, "helper", helperSecond),
		sourceSiteID("call", sourceFile, "user.save", saveCall),
		sourceSiteID("call", sourceFile, "callback", callbackCall),
		sourceSiteID("call", sourceFile, "stop", stopCall),
		sourceSiteID("call", sourceFile, "closure", closureCall),
		sourceSiteID("call", sourceFile, "api.fetchUser", importedCall),
		sourceSiteID("call", sourceFile, "len", builtinCall),
		sourceSiteID("call", sourceFile, "cobra.Command", externalCall),
		sourceSiteID("call", sourceFile, "helper", fileLevelCall),
		sourceSiteID("access", sourceFile, "user.id", idAccess),
		sourceSiteID("access", sourceFile, "config.make", functionSelectorAccess),
	}
	if missing := missingGoldenSourceSites(result.Graph, wantSourceSites); len(missing) > 0 {
		t.Fatalf("missing source sites: %#v", missing)
	}
	if got := countGoldenFalseResolvedEdges(result.Graph, []goldenFalseEdge{
		{relType: graph.RelCalls, sourceID: run.ID, targetID: callback.ID},
		{relType: graph.RelCalls, sourceID: run.ID, targetID: stop.ID},
		{relType: graph.RelCalls, sourceID: "File:" + sourceFile, targetID: helper.ID},
		{relType: graph.RelAccesses, sourceID: run.ID, targetID: configMakeNode.ID},
	}); got != 0 {
		t.Fatalf("golden false resolved edges = %d, want 0", got)
	}
	if got := countGoldenCallAccessSourceSiteOccurrences(result.Graph); got != len(wantSourceSites) {
		t.Fatalf("call/access source-site occurrences = %d, want %d", got, len(wantSourceSites))
	}
	if result.Metrics.ResolvedCalls != 5 ||
		result.Metrics.ResolvedAccesses != 1 ||
		result.Metrics.UnresolvedReferences != 6 {
		t.Fatalf("unexpected metrics: %#v", result.Metrics)
	}
}

func requireGoldenRelationshipMetadata(t *testing.T, relationship graph.Relationship, targetText string, proofKind string, targetRole string, sourceSiteCount int) {
	t.Helper()
	if relationship.SourceSiteStatus != sourceSiteStatusResolved ||
		relationship.ProofKind != proofKind ||
		relationship.TargetRole != targetRole ||
		relationship.TargetText != targetText ||
		relationship.SourceSiteID == "" ||
		relationship.SourceSiteCount != sourceSiteCount {
		t.Fatalf("relationship source-site metadata = %#v", relationship)
	}
}

func requireGoldenSourceSiteIDs(t *testing.T, relationship graph.Relationship, want ...string) {
	t.Helper()
	got := make(map[string]bool, len(relationship.SourceSiteIDs)+1)
	if relationship.SourceSiteID != "" {
		got[relationship.SourceSiteID] = true
	}
	for _, id := range relationship.SourceSiteIDs {
		got[id] = true
	}
	for _, id := range want {
		if !got[id] {
			t.Fatalf("relationship %s missing source site %s in %#v", relationship.ID, id, relationship.SourceSiteIDs)
		}
	}
}

func requireGoldenDiagnostics(t *testing.T, node graph.Node) []graphhealth.Diagnostic {
	t.Helper()
	diagnostics, ok := node.Properties[graphhealth.DiagnosticPropertyKey].([]graphhealth.Diagnostic)
	if !ok || len(diagnostics) == 0 {
		t.Fatalf("node %s diagnostics = %#v", node.ID, node.Properties[graphhealth.DiagnosticPropertyKey])
	}
	return diagnostics
}

func requireGoldenDiagnostic(t *testing.T, diagnostics []graphhealth.Diagnostic, factFamily string, targetText string, status string, proofKind string, classification string, actionability string) graphhealth.Diagnostic {
	t.Helper()
	for _, diagnostic := range diagnostics {
		if diagnostic.FactFamily != factFamily || diagnostic.TargetText != targetText {
			continue
		}
		if diagnostic.SourceSiteID == "" ||
			diagnostic.SourceSiteStatus != status ||
			diagnostic.ProofKind != proofKind ||
			diagnostic.Classification != classification ||
			diagnostic.Actionability != actionability ||
			diagnostic.Count != 1 {
			t.Fatalf("diagnostic %s/%s metadata = %#v", factFamily, targetText, diagnostic)
		}
		return diagnostic
	}
	t.Fatalf("missing diagnostic %s/%s in %#v", factFamily, targetText, diagnostics)
	return graphhealth.Diagnostic{}
}

func missingGoldenSourceSites(g *graph.Graph, want []string) []string {
	seen := make(map[string]bool, len(want))
	for _, relationship := range g.Relationships {
		if relationship.Type != graph.RelCalls && relationship.Type != graph.RelAccesses {
			continue
		}
		if relationship.SourceSiteID != "" {
			seen[relationship.SourceSiteID] = true
		}
		for _, id := range relationship.SourceSiteIDs {
			seen[id] = true
		}
	}
	for _, node := range g.Nodes {
		diagnostics, ok := node.Properties[graphhealth.DiagnosticPropertyKey].([]graphhealth.Diagnostic)
		if !ok {
			continue
		}
		for _, diagnostic := range diagnostics {
			if diagnostic.FactFamily != "call" && diagnostic.FactFamily != "access" {
				continue
			}
			if diagnostic.SourceSiteID != "" {
				seen[diagnostic.SourceSiteID] = true
			}
		}
	}
	var missing []string
	for _, id := range want {
		if !seen[id] {
			missing = append(missing, id)
		}
	}
	return missing
}

type goldenFalseEdge struct {
	relType  graph.RelationshipType
	sourceID string
	targetID string
}

func countGoldenFalseResolvedEdges(g *graph.Graph, edges []goldenFalseEdge) int {
	count := 0
	for _, edge := range edges {
		for _, relationship := range g.Relationships {
			if relationship.Type == edge.relType && relationship.SourceID == edge.sourceID && relationship.TargetID == edge.targetID {
				count++
			}
		}
	}
	return count
}

func countGoldenCallAccessSourceSiteOccurrences(g *graph.Graph) int {
	count := 0
	for _, relationship := range g.Relationships {
		if relationship.Type != graph.RelCalls && relationship.Type != graph.RelAccesses {
			continue
		}
		if relationship.SourceSiteCount > 0 {
			count += relationship.SourceSiteCount
		} else if relationship.SourceSiteID != "" {
			count++
		}
	}
	for _, node := range g.Nodes {
		diagnostics, ok := node.Properties[graphhealth.DiagnosticPropertyKey].([]graphhealth.Diagnostic)
		if !ok {
			continue
		}
		for _, diagnostic := range diagnostics {
			if diagnostic.FactFamily != "call" && diagnostic.FactFamily != "access" {
				continue
			}
			if diagnostic.Count > 0 {
				count += diagnostic.Count
			} else if diagnostic.SourceSiteID != "" {
				count++
			}
		}
	}
	return count
}
