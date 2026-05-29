package semantic

import (
	"testing"

	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/graphhealth"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
)

func TestClassifyAppLayerUsesPrimaryNonOverlappingCategories(t *testing.T) {
	tests := []struct {
		name string
		path string
		want AppLayer
	}{
		{name: "backend", path: "internal/analyze/analyze.go", want: AppLayerBackend},
		{name: "api", path: "internal/httpapi/graph.go", want: AppLayerAPI},
		{name: "mcp api", path: "internal/mcp/tools.go", want: AppLayerAPI},
		{name: "app route api", path: "app/api/users/route.ts", want: AppLayerAPI},
		{name: "pages route api", path: "pages/api/users.ts", want: AppLayerAPI},
		{name: "frontend", path: "anvien-web/src/components/GraphCanvas.tsx", want: AppLayerFrontend},
		{name: "frontend test", path: "anvien-web/e2e/graph.spec.ts", want: AppLayerFrontendTest},
		{name: "api test", path: "internal/httpapi/graph_test.go", want: AppLayerAPITest},
		{name: "backend test", path: "internal/analyze/analyze_test.go", want: AppLayerBackendTest},
		{name: "api contract", path: "internal/contracts/web_ui.go", want: AppLayerAPIContract},
		{name: "generated contract", path: "anvien-web/src/generated/anvien-contracts.ts", want: AppLayerGeneratedContract},
		{name: "frontend api client", path: "anvien-web/src/services/backend-client.ts", want: AppLayerFrontendAPIClient},
		{name: "cli launcher", path: "cmd/anvien/main.go", want: AppLayerCLILauncher},
		{name: "docs", path: "docs/plans/example.md", want: AppLayerDocs},
		{name: "config", path: "anvien-web/package.json", want: AppLayerConfig},
		{name: "unknown", path: "", want: AppLayerUnknown},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ClassifyAppLayer(tt.path)
			if got.Layer != tt.want {
				t.Fatalf("ClassifyAppLayer(%q) = %q source %q, want %q", tt.path, got.Layer, got.Source, tt.want)
			}
		})
	}
}

func TestClassifyFunctionalAreaUsesHighConfidencePathRules(t *testing.T) {
	tests := []struct {
		name string
		path string
		want FunctionalArea
	}{
		{name: "resolution", path: "internal/resolution/resolve.go", want: FunctionalAreaResolution},
		{name: "graph health", path: "internal/graphhealth/compute.go", want: FunctionalAreaGraphHealth},
		{name: "query", path: "internal/group/query.go", want: FunctionalAreaQuery},
		{name: "mcp", path: "internal/mcp/tools.go", want: FunctionalAreaMCP},
		{name: "api", path: "internal/httpapi/graph.go", want: FunctionalAreaAPI},
		{name: "contracts", path: "internal/contracts/web_ui.go", want: FunctionalAreaContracts},
		{name: "layout", path: "anvien-web/src/lib/graph-adapter.ts", want: FunctionalAreaLayout},
		{name: "web graph ui", path: "anvien-web/src/components/GraphCanvas.tsx", want: FunctionalAreaWebGraphUI},
		{name: "providers", path: "internal/providers/tsjs/extract.go", want: FunctionalAreaProviders},
		{name: "storage", path: "internal/lbugload/csv.go", want: FunctionalAreaStorage},
		{name: "embeddings", path: "internal/embeddings/pipeline.go", want: FunctionalAreaEmbeddings},
		{name: "session", path: "internal/session/controller.go", want: FunctionalAreaSession},
		{name: "cli", path: "cmd/anvien/main.go", want: FunctionalAreaCLI},
		{name: "launcher", path: "anvien-launcher/src/main.go", want: FunctionalAreaLauncher},
		{name: "reporting", path: "reports/problem/example.md", want: FunctionalAreaReporting},
		{name: "docs", path: "docs/plans/example.md", want: FunctionalAreaDocumentation},
		{name: "config", path: "go.mod", want: FunctionalAreaConfiguration},
		{name: "unknown", path: "scratch/example.tmp", want: FunctionalAreaUnknown},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ClassifyFunctionalArea(tt.path)
			if got.Area != tt.want {
				t.Fatalf("ClassifyFunctionalArea(%q) = %q source %q, want %q", tt.path, got.Area, got.Source, tt.want)
			}
		})
	}
}

func TestApplyPersistsAppLayerAndInfersProcessLayer(t *testing.T) {
	g := graph.New()
	g.AddNode(graph.Node{ID: "Function:api", Label: scopeir.NodeFunction, Properties: graph.NodeProperties{"name": "api", "filePath": "internal/httpapi/graph.go"}})
	g.AddNode(graph.Node{ID: "Process:api", Label: scopeir.NodeProcess, Properties: graph.NodeProperties{"name": "api process"}})
	g.AddRelationship(graph.Relationship{ID: "step", SourceID: "Function:api", TargetID: "Process:api", Type: graph.RelStepInProcess})

	result, err := Apply(g)
	if err != nil {
		t.Fatalf("Apply() error = %v", err)
	}
	apiNode, ok := g.GetNode("Function:api")
	if !ok {
		t.Fatal("missing api node")
	}
	if got := apiNode.Properties[AppLayerProperty]; got != string(AppLayerAPI) {
		t.Fatalf("api node app layer = %v, want %s", got, AppLayerAPI)
	}
	processNode, ok := g.GetNode("Process:api")
	if !ok {
		t.Fatal("missing process node")
	}
	if got := processNode.Properties[AppLayerProperty]; got != string(AppLayerAPI) {
		t.Fatalf("process app layer = %v, want %s", got, AppLayerAPI)
	}
	if got := apiNode.Properties[FunctionalAreaProperty]; got != string(FunctionalAreaAPI) {
		t.Fatalf("api node functional area = %v, want %s", got, FunctionalAreaAPI)
	}
	if got := processNode.Properties[FunctionalAreaProperty]; got != string(FunctionalAreaAPI) {
		t.Fatalf("process functional area = %v, want %s", got, FunctionalAreaAPI)
	}
	if result.Metrics.NodesInferredFromRelationships != 1 {
		t.Fatalf("inferred count = %d, want 1", result.Metrics.NodesInferredFromRelationships)
	}
}

func TestApplyUsesMixedForRelationshipBackedMultiLayerNodes(t *testing.T) {
	g := graph.New()
	g.AddNode(graph.Node{ID: "Function:backend", Label: scopeir.NodeFunction, Properties: graph.NodeProperties{"name": "backend", "filePath": "internal/analyze/analyze.go"}})
	g.AddNode(graph.Node{ID: "Function:frontend", Label: scopeir.NodeFunction, Properties: graph.NodeProperties{"name": "frontend", "filePath": "anvien-web/src/lib/graph-adapter.ts"}})
	g.AddNode(graph.Node{ID: "Community:mixed", Label: scopeir.NodeCommunity, Properties: graph.NodeProperties{"name": "mixed"}})
	g.AddRelationship(graph.Relationship{ID: "member-a", SourceID: "Function:backend", TargetID: "Community:mixed", Type: graph.RelMemberOf})
	g.AddRelationship(graph.Relationship{ID: "member-b", SourceID: "Function:frontend", TargetID: "Community:mixed", Type: graph.RelMemberOf})

	if _, err := Apply(g); err != nil {
		t.Fatalf("Apply() error = %v", err)
	}
	node, ok := g.GetNode("Community:mixed")
	if !ok {
		t.Fatal("missing community")
	}
	if got := node.Properties[AppLayerProperty]; got != string(AppLayerMixed) {
		t.Fatalf("community app layer = %v, want %s", got, AppLayerMixed)
	}
	if got := node.Properties[FunctionalAreaProperty]; got != string(FunctionalAreaMixed) {
		t.Fatalf("community functional area = %v, want %s", got, FunctionalAreaMixed)
	}
}

func TestApplyPersistsSourceBackedResolutionGaps(t *testing.T) {
	g := graph.New()
	g.AddNode(graph.Node{
		ID:    "Function:resolve",
		Label: scopeir.NodeFunction,
		Properties: graph.NodeProperties{
			"name":     "resolve",
			"filePath": "internal/resolution/resolve.go",
			graphhealth.DiagnosticPropertyKey: []graphhealth.Diagnostic{
				{
					Kind:             graphhealth.DiagnosticUnresolvedReference,
					FactFamily:       "call",
					SourceNodeID:     "Function:resolve",
					TargetText:       "stop",
					TargetRole:       "callable",
					SourceSiteID:     "SourceSite:resolve.go#call#stop#10#4#10#10",
					SourceSiteStatus: "unresolved_local_binding",
					ProofKind:        "none",
					Classification:   graphhealth.DiagnosticClassificationInRepoUnresolved,
					Actionability:    graphhealth.DiagnosticActionabilityAnalyzerGap,
					ResolutionSource: "scope-resolution",
					FilePath:         "internal/resolution/resolve.go",
					StartLine:        10,
					StartCol:         4,
					EndLine:          10,
					EndCol:           10,
					Count:            2,
					Note:             "bare call has no binding proof",
				},
			},
		},
	})

	result, err := Apply(g)
	if err != nil {
		t.Fatalf("Apply() error = %v", err)
	}
	if result.Metrics.ResolutionGapInputs != 1 ||
		result.Metrics.ResolutionGapNodes != 1 ||
		result.Metrics.ResolutionGapRelationships != 1 {
		t.Fatalf("resolution gap metrics = %#v", result.Metrics)
	}
	source, ok := g.GetNode("Function:resolve")
	if !ok {
		t.Fatal("missing source node")
	}
	if got := source.Properties[AppLayerProperty]; got != string(AppLayerBackend) {
		t.Fatalf("source app layer = %v, want %s", got, AppLayerBackend)
	}
	gapID := graphhealth.ResolutionGapNodeID("SourceSite:resolve.go#call#stop#10#4#10#10")
	gapNode, ok := g.GetNode(gapID)
	if !ok {
		t.Fatalf("missing persisted resolution gap %q", gapID)
	}
	if gapNode.Label != scopeir.NodeResolutionGap ||
		gapNode.Properties["gapKind"] != graphhealth.ResolutionGapKindUnresolvedCall ||
		gapNode.Properties["sourceNodeId"] != "Function:resolve" ||
		gapNode.Properties["targetText"] != "stop" ||
		gapNode.Properties["sourceSiteStatus"] != "unresolved_local_binding" ||
		gapNode.Properties["classification"] != graphhealth.DiagnosticClassificationInRepoUnresolved ||
		gapNode.Properties["actionability"] != graphhealth.DiagnosticActionabilityAnalyzerGap ||
		gapNode.Properties[AppLayerProperty] != string(AppLayerBackend) ||
		gapNode.Properties[FunctionalAreaProperty] != string(FunctionalAreaResolution) ||
		gapNode.Properties["count"] != 2 {
		t.Fatalf("persisted resolution gap lost metadata: %#v", gapNode)
	}
	foundGapRelationship := false
	for _, relationship := range g.Relationships {
		if relationship.Type == graph.RelHasResolutionGap {
			foundGapRelationship = true
			if relationship.SourceID != "Function:resolve" ||
				relationship.TargetID != gapID ||
				relationship.SourceSiteID != "SourceSite:resolve.go#call#stop#10#4#10#10" ||
				relationship.SourceSiteCount != 2 ||
				relationship.TargetText != "stop" {
				t.Fatalf("gap relationship lost source-site evidence: %#v", relationship)
			}
		}
		if relationship.Type == graph.RelCalls && relationship.TargetText == "stop" {
			t.Fatalf("unresolved call was incorrectly emitted as CALLS: %#v", relationship)
		}
	}
	if !foundGapRelationship {
		t.Fatal("missing HAS_RESOLUTION_GAP relationship")
	}
	if _, ok := g.GetNode("Function:stop"); ok {
		t.Fatal("unresolved target was incorrectly synthesized as an in-repo Function node")
	}
}

func TestApplyPersistsResolutionGapRolesClassificationsAndOccurrences(t *testing.T) {
	g := graph.New()
	g.AddNode(graph.Node{
		ID:    "Function:resolveMany",
		Label: scopeir.NodeFunction,
		Properties: graph.NodeProperties{
			"name":     "resolveMany",
			"filePath": "internal/resolution/resolve.go",
			graphhealth.DiagnosticPropertyKey: []graphhealth.Diagnostic{
				resolutionGapDiagnostic("call", "len", "site-call-builtin-len", "unresolved_external", "builtin-reference", "", graphhealth.DiagnosticClassificationBuiltin, graphhealth.DiagnosticActionabilityNonActionable, 2),
				resolutionGapDiagnostic("access", "client.Do", "site-access-external-client-do", "unresolved_external", "external-member", "", graphhealth.DiagnosticClassificationExternalLibrary, graphhealth.DiagnosticActionabilityReview, 1),
				resolutionGapDiagnostic("type-reference", "context.Context", "site-type-stdlib-context", "unresolved_external", "stdlib-type", "", graphhealth.DiagnosticClassificationStandardLibrary, graphhealth.DiagnosticActionabilityNonActionable, 1),
				resolutionGapDiagnostic("heritage", "MissingBase", "site-heritage-inrepo-missing-base", "unresolved_local_binding", "none", "", graphhealth.DiagnosticClassificationInRepoUnresolved, graphhealth.DiagnosticActionabilityAnalyzerGap, 1),
				resolutionGapDiagnostic("reference", "cobra.Command", "site-reference-external-cobra", "unresolved_external", "none", "", graphhealth.DiagnosticClassificationExternalLibrary, graphhealth.DiagnosticActionabilityReview, 1),
				resolutionGapDiagnostic("reference", "t.Fatalf", "site-reference-test-fatalf", "unresolved_external", "none", "", graphhealth.DiagnosticClassificationTestFramework, graphhealth.DiagnosticActionabilityNonActionable, 1),
				resolutionGapDiagnostic("reference", "mystery", "site-reference-unknown-mystery", "unknown", "none", "", graphhealth.DiagnosticClassificationUnclassified, graphhealth.DiagnosticActionabilityAnalyzerGap, 1),
				resolutionGapDiagnostic("call", "missingOne", "site-call-missing-one", "unresolved_local_binding", "none", "", graphhealth.DiagnosticClassificationInRepoUnresolved, graphhealth.DiagnosticActionabilityAnalyzerGap, 3),
				resolutionGapDiagnostic("call", "missingTwo", "site-call-missing-two", "unresolved_local_binding", "none", "", graphhealth.DiagnosticClassificationInRepoUnresolved, graphhealth.DiagnosticActionabilityAnalyzerGap, 1),
			},
		},
	})

	result, err := Apply(g)
	if err != nil {
		t.Fatalf("Apply() error = %v", err)
	}
	if result.Metrics.ResolutionGapInputs != 9 ||
		result.Metrics.ResolutionGapNodes != 9 ||
		result.Metrics.ResolutionGapRelationships != 9 {
		t.Fatalf("resolution gap metrics = %#v, want 9 persisted gaps", result.Metrics)
	}
	validation := graphhealth.ValidateResolutionGapPersistence(g)
	if !validation.OK() {
		t.Fatalf("ValidateResolutionGapPersistence() = %#v, want OK", validation)
	}

	gaps := resolutionGapNodesByTargetText(g)
	requireResolutionGapNode(t, gaps, "len", "callable", graphhealth.ResolutionGapKindUnresolvedCall, graphhealth.DiagnosticClassificationBuiltin, graphhealth.DiagnosticActionabilityNonActionable, "site-call-builtin-len", 2)
	requireResolutionGapNode(t, gaps, "client.Do", "member", graphhealth.ResolutionGapKindUnresolvedAccess, graphhealth.DiagnosticClassificationExternalLibrary, graphhealth.DiagnosticActionabilityReview, "site-access-external-client-do", 1)
	requireResolutionGapNode(t, gaps, "context.Context", "type", graphhealth.ResolutionGapKindUnresolvedTypeReference, graphhealth.DiagnosticClassificationStandardLibrary, graphhealth.DiagnosticActionabilityNonActionable, "site-type-stdlib-context", 1)
	requireResolutionGapNode(t, gaps, "MissingBase", "type", graphhealth.ResolutionGapKindUnresolvedHeritage, graphhealth.DiagnosticClassificationInRepoUnresolved, graphhealth.DiagnosticActionabilityAnalyzerGap, "site-heritage-inrepo-missing-base", 1)
	requireResolutionGapNode(t, gaps, "cobra.Command", "external", graphhealth.ResolutionGapKindUnresolvedReference, graphhealth.DiagnosticClassificationExternalLibrary, graphhealth.DiagnosticActionabilityReview, "site-reference-external-cobra", 1)
	requireResolutionGapNode(t, gaps, "t.Fatalf", "test", graphhealth.ResolutionGapKindUnresolvedReference, graphhealth.DiagnosticClassificationTestFramework, graphhealth.DiagnosticActionabilityNonActionable, "site-reference-test-fatalf", 1)
	requireResolutionGapNode(t, gaps, "mystery", "unknown", graphhealth.ResolutionGapKindUnresolvedReference, graphhealth.DiagnosticClassificationUnclassified, graphhealth.DiagnosticActionabilityAnalyzerGap, "site-reference-unknown-mystery", 1)
	requireResolutionGapNode(t, gaps, "missingOne", "callable", graphhealth.ResolutionGapKindUnresolvedCall, graphhealth.DiagnosticClassificationInRepoUnresolved, graphhealth.DiagnosticActionabilityAnalyzerGap, "site-call-missing-one", 3)
	requireResolutionGapNode(t, gaps, "missingTwo", "callable", graphhealth.ResolutionGapKindUnresolvedCall, graphhealth.DiagnosticClassificationInRepoUnresolved, graphhealth.DiagnosticActionabilityAnalyzerGap, "site-call-missing-two", 1)

	for _, relationship := range g.Relationships {
		if relationship.Type == graph.RelCalls || relationship.Type == graph.RelAccesses || relationship.Type == graph.RelUses || relationship.Type == graph.RelInherits {
			t.Fatalf("unresolved diagnostics were incorrectly emitted as resolved semantic relationship: %#v", relationship)
		}
	}
	for _, fakeTarget := range []string{"Function:len", "Function:client.Do", "Type:context.Context", "Class:MissingBase", "Function:missingOne", "Function:missingTwo"} {
		if _, ok := g.GetNode(fakeTarget); ok {
			t.Fatalf("unresolved target was incorrectly synthesized as %s", fakeTarget)
		}
	}
}

func resolutionGapDiagnostic(factFamily, targetText, sourceSiteID, status, proofKind, targetRole, classification, actionability string, count int) graphhealth.Diagnostic {
	return graphhealth.Diagnostic{
		Kind:             graphhealth.DiagnosticUnresolvedReference,
		FactFamily:       factFamily,
		SourceNodeID:     "Function:resolveMany",
		TargetText:       targetText,
		TargetRole:       targetRole,
		SourceSiteID:     sourceSiteID,
		SourceSiteStatus: status,
		ProofKind:        proofKind,
		Classification:   classification,
		Actionability:    actionability,
		ResolutionSource: "scope-resolution",
		FilePath:         "internal/resolution/resolve.go",
		FileHash:         "hash-resolution",
		StartLine:        10,
		StartCol:         2,
		EndLine:          10,
		EndCol:           12,
		Count:            count,
		Note:             "source-backed unresolved reference fixture",
	}
}

func resolutionGapNodesByTargetText(g *graph.Graph) map[string]graph.Node {
	gaps := map[string]graph.Node{}
	for _, node := range g.Nodes {
		if node.Label != scopeir.NodeResolutionGap {
			continue
		}
		targetText, _ := node.Properties["targetText"].(string)
		gaps[targetText] = node
	}
	return gaps
}

func requireResolutionGapNode(t *testing.T, gaps map[string]graph.Node, targetText, targetRole, gapKind, classification, actionability, sourceSiteID string, count int) {
	t.Helper()
	node, ok := gaps[targetText]
	if !ok {
		t.Fatalf("missing ResolutionGap targetText=%q in %#v", targetText, gaps)
	}
	if node.Properties["targetRole"] != targetRole ||
		node.Properties["gapKind"] != gapKind ||
		node.Properties["classification"] != classification ||
		node.Properties["actionability"] != actionability ||
		node.Properties["sourceSiteId"] != sourceSiteID ||
		node.Properties["sourceSiteStatus"] == "" ||
		node.Properties["proofKind"] == "" ||
		node.Properties["sourceAppLayer"] != string(AppLayerBackend) ||
		node.Properties["sourceFunctionalArea"] != string(FunctionalAreaResolution) ||
		node.Properties[AppLayerProperty] != string(AppLayerBackend) ||
		node.Properties[FunctionalAreaProperty] != string(FunctionalAreaResolution) ||
		node.Properties["count"] != count {
		t.Fatalf("ResolutionGap %q lost persisted metadata: %#v", targetText, node)
	}
}

func TestGraphSemanticStatusDistinguishesUnknownFromMissingMetadata(t *testing.T) {
	g := graph.New()
	g.AddNode(graph.Node{ID: "Function:fresh", Label: scopeir.NodeFunction, Properties: graph.NodeProperties{
		AppLayerProperty:       string(AppLayerUnknown),
		AppLayerSourceProperty: "unmatched_path",
	}})
	g.AddNode(graph.Node{ID: "Function:stale", Label: scopeir.NodeFunction, Properties: graph.NodeProperties{
		"name": "stale",
	}})

	status := GraphSemanticStatus(g)

	if status.SchemaVersion != SchemaVersion {
		t.Fatalf("schema version = %q, want %q", status.SchemaVersion, SchemaVersion)
	}
	if status.AppLayer.Status != StatusStaleIncomplete {
		t.Fatalf("app layer status = %q, want %q", status.AppLayer.Status, StatusStaleIncomplete)
	}
	if status.AppLayer.UnknownNodes != 1 {
		t.Fatalf("unknown nodes = %d, want 1", status.AppLayer.UnknownNodes)
	}
	if status.AppLayer.MissingNodes != 1 || status.AppLayer.MissingSourceNodes != 1 {
		t.Fatalf("missing app layer counts = %#v, want one stale node", status.AppLayer)
	}
	if status.FunctionalArea.Status != StatusStaleIncomplete {
		t.Fatalf("functional area status = %q, want %q", status.FunctionalArea.Status, StatusStaleIncomplete)
	}
}

func TestSemanticTermDefinitionsAreStableAndNonOverlapping(t *testing.T) {
	appLayerDefinitions := AppLayerDefinitions()
	if len(appLayerDefinitions) != len(AppLayers) {
		t.Fatalf("app layer definitions = %d, want %d", len(appLayerDefinitions), len(AppLayers))
	}
	requireUniqueTermKeys(t, appLayerDefinitions)
	functionalAreaDefinitions := FunctionalAreaDefinitions()
	if len(functionalAreaDefinitions) != len(FunctionalAreas) {
		t.Fatalf("functional area definitions = %d, want %d", len(functionalAreaDefinitions), len(FunctionalAreas))
	}
	requireUniqueTermKeys(t, functionalAreaDefinitions)
	terms := SemanticTermDefinitions()
	requireUniqueTermKeys(t, terms)
	requireTerm(t, terms, "resolution_gap", "Resolution Gap")
	requireTerm(t, terms, "non_actionable_reference", "Non-actionable Reference")
	requireTerm(t, appLayerDefinitions, string(AppLayerAPI), "API Layer")
	requireTerm(t, appLayerDefinitions, string(AppLayerFrontendAPIClient), "Frontend API Client")
	requireTerm(t, functionalAreaDefinitions, string(FunctionalAreaResolution), "Resolution")
	requireTerm(t, functionalAreaDefinitions, string(FunctionalAreaGraphHealth), "Graph Health")
}

func requireUniqueTermKeys(t *testing.T, terms []TermDefinition) {
	t.Helper()
	seen := map[string]bool{}
	for _, term := range terms {
		if term.Key == "" || term.DisplayLabel == "" || term.CLILabel == "" || term.WebLabel == "" {
			t.Fatalf("incomplete term definition: %#v", term)
		}
		if seen[term.Key] {
			t.Fatalf("duplicate term key %q", term.Key)
		}
		seen[term.Key] = true
	}
}

func requireTerm(t *testing.T, terms []TermDefinition, key string, displayLabel string) {
	t.Helper()
	for _, term := range terms {
		if term.Key == key {
			if term.DisplayLabel != displayLabel {
				t.Fatalf("%s display label = %q, want %q", key, term.DisplayLabel, displayLabel)
			}
			return
		}
	}
	t.Fatalf("missing term %q in %#v", key, terms)
}
