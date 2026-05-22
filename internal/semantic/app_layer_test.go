package semantic

import (
	"testing"

	"github.com/tamnguyendinh/avmatrix-go/internal/graph"
	"github.com/tamnguyendinh/avmatrix-go/internal/scopeir"
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
		{name: "frontend", path: "avmatrix-web/src/components/GraphCanvas.tsx", want: AppLayerFrontend},
		{name: "frontend test", path: "avmatrix-web/e2e/graph.spec.ts", want: AppLayerFrontendTest},
		{name: "api test", path: "internal/httpapi/graph_test.go", want: AppLayerAPITest},
		{name: "backend test", path: "internal/analyze/analyze_test.go", want: AppLayerBackendTest},
		{name: "api contract", path: "internal/contracts/web_ui.go", want: AppLayerAPIContract},
		{name: "generated contract", path: "avmatrix-web/src/generated/avmatrix-contracts.ts", want: AppLayerGeneratedContract},
		{name: "frontend api client", path: "avmatrix-web/src/services/backend-client.ts", want: AppLayerFrontendAPIClient},
		{name: "cli launcher", path: "cmd/avmatrix/main.go", want: AppLayerCLILauncher},
		{name: "docs", path: "docs/plans/example.md", want: AppLayerDocs},
		{name: "config", path: "avmatrix-web/package.json", want: AppLayerConfig},
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
	if result.Metrics.NodesInferredFromRelationships != 1 {
		t.Fatalf("inferred count = %d, want 1", result.Metrics.NodesInferredFromRelationships)
	}
}

func TestApplyUsesMixedForRelationshipBackedMultiLayerNodes(t *testing.T) {
	g := graph.New()
	g.AddNode(graph.Node{ID: "Function:backend", Label: scopeir.NodeFunction, Properties: graph.NodeProperties{"name": "backend", "filePath": "internal/analyze/analyze.go"}})
	g.AddNode(graph.Node{ID: "Function:frontend", Label: scopeir.NodeFunction, Properties: graph.NodeProperties{"name": "frontend", "filePath": "avmatrix-web/src/lib/graph-adapter.ts"}})
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
}

func TestSemanticTermDefinitionsAreStableAndNonOverlapping(t *testing.T) {
	appLayerDefinitions := AppLayerDefinitions()
	if len(appLayerDefinitions) != len(AppLayers) {
		t.Fatalf("app layer definitions = %d, want %d", len(appLayerDefinitions), len(AppLayers))
	}
	requireUniqueTermKeys(t, appLayerDefinitions)
	terms := SemanticTermDefinitions()
	requireUniqueTermKeys(t, terms)
	requireTerm(t, terms, "resolution_gap", "Resolution Gap")
	requireTerm(t, terms, "non_actionable_reference", "Non-actionable Reference")
	requireTerm(t, appLayerDefinitions, string(AppLayerAPI), "API Layer")
	requireTerm(t, appLayerDefinitions, string(AppLayerFrontendAPIClient), "Frontend API Client")
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
