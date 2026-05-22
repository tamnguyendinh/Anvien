package semantic

import (
	"path/filepath"
	"strings"

	"github.com/tamnguyendinh/avmatrix-go/internal/graph"
	"github.com/tamnguyendinh/avmatrix-go/internal/graphhealth"
	"github.com/tamnguyendinh/avmatrix-go/internal/scopeir"
)

const (
	AppLayerProperty       = "appLayer"
	AppLayerSourceProperty = "appLayerSource"
)

type AppLayer string

const (
	AppLayerBackend           AppLayer = "backend"
	AppLayerAPI               AppLayer = "api"
	AppLayerFrontend          AppLayer = "frontend"
	AppLayerCLILauncher       AppLayer = "cli_launcher"
	AppLayerSharedContract    AppLayer = "shared_contract"
	AppLayerAPIContract       AppLayer = "api_contract"
	AppLayerAPISharedContract AppLayer = "api_shared_contract"
	AppLayerFrontendAPIClient AppLayer = "frontend_api_client"
	AppLayerBackendTest       AppLayer = "backend_test"
	AppLayerFrontendTest      AppLayer = "frontend_test"
	AppLayerAPITest           AppLayer = "api_test"
	AppLayerGeneratedContract AppLayer = "generated_contract"
	AppLayerDocs              AppLayer = "docs"
	AppLayerConfig            AppLayer = "config"
	AppLayerGenerated         AppLayer = "generated"
	AppLayerMixed             AppLayer = "mixed"
	AppLayerUnknown           AppLayer = "unknown"
)

var AppLayers = []AppLayer{
	AppLayerBackend,
	AppLayerAPI,
	AppLayerFrontend,
	AppLayerCLILauncher,
	AppLayerSharedContract,
	AppLayerAPIContract,
	AppLayerAPISharedContract,
	AppLayerFrontendAPIClient,
	AppLayerBackendTest,
	AppLayerFrontendTest,
	AppLayerAPITest,
	AppLayerGeneratedContract,
	AppLayerDocs,
	AppLayerConfig,
	AppLayerGenerated,
	AppLayerMixed,
	AppLayerUnknown,
}

type Classification struct {
	Layer  AppLayer
	Source string
}

type Result struct {
	Metrics Metrics `json:"metrics"`
}

type Metrics struct {
	NodesVisited                   int            `json:"nodesVisited"`
	NodesWithFilePath              int            `json:"nodesWithFilePath"`
	NodesClassified                int            `json:"nodesClassified"`
	NodesUnknown                   int            `json:"nodesUnknown"`
	NodesInferredFromRelationships int            `json:"nodesInferredFromRelationships"`
	FilePathCacheEntries           int            `json:"filePathCacheEntries"`
	FunctionalPathCacheEntries     int            `json:"functionalPathCacheEntries"`
	RelationshipsScanned           int            `json:"relationshipsScanned"`
	AppLayerCounts                 map[string]int `json:"appLayerCounts,omitempty"`
	FunctionalNodesClassified      int            `json:"functionalNodesClassified"`
	FunctionalNodesUnknown         int            `json:"functionalNodesUnknown"`
	FunctionalAreaCounts           map[string]int `json:"functionalAreaCounts,omitempty"`
	ResolutionGapInputs            int            `json:"resolutionGapInputs"`
	ResolutionGapNodes             int            `json:"resolutionGapNodes"`
	ResolutionGapRelationships     int            `json:"resolutionGapRelationships"`
}

func AppLayerStrings() []string {
	out := make([]string, 0, len(AppLayers))
	for _, layer := range AppLayers {
		out = append(out, string(layer))
	}
	return out
}

func Apply(g *graph.Graph) (Result, error) {
	var result Result
	result.Metrics.AppLayerCounts = map[string]int{}
	result.Metrics.FunctionalAreaCounts = map[string]int{}
	if g == nil {
		return result, nil
	}

	pathCache := map[string]Classification{}
	functionalPathCache := map[string]FunctionalAreaClassification{}
	nodeIndex := make(map[string]int, len(g.Nodes))
	nodeLayers := make(map[string]AppLayer, len(g.Nodes))
	nodeAreas := make(map[string]FunctionalArea, len(g.Nodes))

	for index := range g.Nodes {
		node := &g.Nodes[index]
		nodeIndex[node.ID] = index
		result.Metrics.NodesVisited++
		if node.Properties == nil {
			node.Properties = graph.NodeProperties{}
		}

		filePath := stringProperty(node.Properties, "filePath")
		if filePath != "" {
			result.Metrics.NodesWithFilePath++
		}
		classification, ok := pathCache[filePath]
		if !ok {
			classification = ClassifyAppLayer(filePath)
			pathCache[filePath] = classification
		}
		setAppLayer(node, classification)
		nodeLayers[node.ID] = classification.Layer

		functionalClassification, ok := functionalPathCache[filePath]
		if !ok {
			functionalClassification = ClassifyFunctionalArea(filePath)
			functionalPathCache[filePath] = functionalClassification
		}
		setFunctionalArea(node, functionalClassification)
		nodeAreas[node.ID] = functionalClassification.Area
	}

	result.Metrics.FilePathCacheEntries = len(pathCache)
	result.Metrics.FunctionalPathCacheEntries = len(functionalPathCache)
	inferred := inferRelationBackedLayers(g, nodeIndex, nodeLayers)
	for nodeID, classification := range inferred {
		index := nodeIndex[nodeID]
		node := &g.Nodes[index]
		if nodeLayers[nodeID] != AppLayerUnknown {
			continue
		}
		setAppLayer(node, classification)
		nodeLayers[nodeID] = classification.Layer
		result.Metrics.NodesInferredFromRelationships++
	}
	inferredAreas := inferRelationBackedFunctionalAreas(g, nodeIndex, nodeAreas)
	for nodeID, classification := range inferredAreas {
		index := nodeIndex[nodeID]
		node := &g.Nodes[index]
		if nodeAreas[nodeID] != FunctionalAreaUnknown {
			continue
		}
		setFunctionalArea(node, classification)
		nodeAreas[nodeID] = classification.Area
	}

	gapMetrics := persistResolutionGaps(g, nodeLayers, nodeAreas)
	result.Metrics.ResolutionGapInputs = gapMetrics.Inputs
	result.Metrics.ResolutionGapNodes = gapMetrics.Nodes
	result.Metrics.ResolutionGapRelationships = gapMetrics.Relationships
	result.Metrics.NodesVisited += gapMetrics.Nodes
	result.Metrics.NodesWithFilePath += gapMetrics.NodesWithFilePath
	result.Metrics.RelationshipsScanned = len(g.Relationships)

	for _, layer := range nodeLayers {
		result.Metrics.AppLayerCounts[string(layer)]++
		if layer == AppLayerUnknown {
			result.Metrics.NodesUnknown++
			continue
		}
		result.Metrics.NodesClassified++
	}
	for _, area := range nodeAreas {
		result.Metrics.FunctionalAreaCounts[string(area)]++
		if area == FunctionalAreaUnknown {
			result.Metrics.FunctionalNodesUnknown++
			continue
		}
		result.Metrics.FunctionalNodesClassified++
	}

	return result, nil
}

type resolutionGapPersistMetrics struct {
	Inputs            int
	Nodes             int
	Relationships     int
	NodesWithFilePath int
}

func persistResolutionGaps(g *graph.Graph, nodeLayers map[string]AppLayer, nodeAreas map[string]FunctionalArea) resolutionGapPersistMetrics {
	var metrics resolutionGapPersistMetrics
	inputs := graphhealth.SourceBackedResolutionGapInputs(g)
	metrics.Inputs = len(inputs)
	for _, input := range inputs {
		if strings.TrimSpace(input.SourceNodeID) == "" {
			continue
		}
		if _, ok := g.GetNode(input.SourceNodeID); !ok {
			continue
		}
		if layer := nodeLayers[input.SourceNodeID]; layer != "" {
			input.SourceAppLayer = string(layer)
		}
		if area := nodeAreas[input.SourceNodeID]; area != "" {
			input.SourceFunctionalArea = string(area)
		}
		gapNode := input.GraphNode()
		_, existingNode := g.GetNode(gapNode.ID)
		g.AddNode(gapNode)
		if !existingNode {
			metrics.Nodes++
			if stringProperty(gapNode.Properties, "filePath") != "" {
				metrics.NodesWithFilePath++
			}
		}
		nodeLayers[gapNode.ID] = appLayerFromString(input.SourceAppLayer)
		nodeAreas[gapNode.ID] = functionalAreaFromString(input.SourceFunctionalArea)

		gapRelationship := input.GraphRelationship()
		_, existingRelationship := g.GetRelationship(gapRelationship.ID)
		g.AddRelationship(gapRelationship)
		if !existingRelationship {
			metrics.Relationships++
		}
	}
	return metrics
}

func appLayerFromString(value string) AppLayer {
	value = strings.TrimSpace(value)
	if value == "" {
		return AppLayerUnknown
	}
	return AppLayer(value)
}

func functionalAreaFromString(value string) FunctionalArea {
	value = strings.TrimSpace(value)
	if value == "" {
		return FunctionalAreaUnknown
	}
	return FunctionalArea(value)
}

func ClassifyAppLayer(filePath string) Classification {
	path := normalizePath(filePath)
	if path == "" {
		return Classification{Layer: AppLayerUnknown, Source: "missing_file_path"}
	}

	base := filepath.Base(path)
	if isDocsPath(path, base) {
		return Classification{Layer: AppLayerDocs, Source: "docs_path"}
	}
	if isWebTestPath(path, base) {
		return Classification{Layer: AppLayerFrontendTest, Source: "frontend_test_path"}
	}
	if isAPITestPath(path, base) {
		return Classification{Layer: AppLayerAPITest, Source: "api_test_path"}
	}
	if isBackendTestPath(path, base) {
		return Classification{Layer: AppLayerBackendTest, Source: "backend_test_path"}
	}
	if isGeneratedContractPath(path) {
		return Classification{Layer: AppLayerGeneratedContract, Source: "generated_contract_path"}
	}
	if isAPIContractPath(path) {
		return Classification{Layer: AppLayerAPIContract, Source: "api_contract_path"}
	}
	if isFrontendAPIClientPath(path) {
		return Classification{Layer: AppLayerFrontendAPIClient, Source: "frontend_api_client_path"}
	}
	if isConfigPath(path, base) {
		return Classification{Layer: AppLayerConfig, Source: "config_path"}
	}
	if isGeneratedPath(path) {
		return Classification{Layer: AppLayerGenerated, Source: "generated_path"}
	}
	if isAPIPath(path) {
		return Classification{Layer: AppLayerAPI, Source: "api_path"}
	}
	if isFrontendPath(path) {
		return Classification{Layer: AppLayerFrontend, Source: "frontend_path"}
	}
	if isCLILauncherPath(path, base) {
		return Classification{Layer: AppLayerCLILauncher, Source: "cli_launcher_path"}
	}
	if isBackendPath(path) {
		return Classification{Layer: AppLayerBackend, Source: "backend_path"}
	}

	return Classification{Layer: AppLayerUnknown, Source: "unmatched_path"}
}

func inferRelationBackedLayers(g *graph.Graph, nodeIndex map[string]int, nodeLayers map[string]AppLayer) map[string]Classification {
	layerSets := map[string]map[AppLayer]struct{}{}
	for _, rel := range g.Relationships {
		if rel.Type != graph.RelStepInProcess && rel.Type != graph.RelMemberOf {
			continue
		}
		targetIndex, ok := nodeIndex[rel.TargetID]
		if !ok {
			continue
		}
		target := g.Nodes[targetIndex]
		if target.Label != scopeir.NodeProcess && target.Label != scopeir.NodeCommunity {
			continue
		}
		sourceLayer := nodeLayers[rel.SourceID]
		if sourceLayer == "" || sourceLayer == AppLayerUnknown {
			continue
		}
		if _, ok := layerSets[rel.TargetID]; !ok {
			layerSets[rel.TargetID] = map[AppLayer]struct{}{}
		}
		layerSets[rel.TargetID][sourceLayer] = struct{}{}
	}

	out := map[string]Classification{}
	for nodeID, layers := range layerSets {
		if len(layers) == 0 {
			continue
		}
		if len(layers) == 1 {
			for layer := range layers {
				out[nodeID] = Classification{Layer: layer, Source: "relationship_membership"}
			}
			continue
		}
		out[nodeID] = Classification{Layer: AppLayerMixed, Source: "mixed_relationship_membership"}
	}
	return out
}

func setAppLayer(node *graph.Node, classification Classification) {
	if classification.Layer == "" {
		classification.Layer = AppLayerUnknown
	}
	if classification.Source == "" {
		classification.Source = "unspecified"
	}
	node.Properties[AppLayerProperty] = string(classification.Layer)
	node.Properties[AppLayerSourceProperty] = classification.Source
}

func setFunctionalArea(node *graph.Node, classification FunctionalAreaClassification) {
	if classification.Area == "" {
		classification.Area = FunctionalAreaUnknown
	}
	if classification.Source == "" {
		classification.Source = "unspecified"
	}
	node.Properties[FunctionalAreaProperty] = string(classification.Area)
	node.Properties[FunctionalAreaSourceProperty] = classification.Source
}

func inferRelationBackedFunctionalAreas(g *graph.Graph, nodeIndex map[string]int, nodeAreas map[string]FunctionalArea) map[string]FunctionalAreaClassification {
	areaSets := map[string]map[FunctionalArea]struct{}{}
	for _, rel := range g.Relationships {
		if rel.Type != graph.RelStepInProcess && rel.Type != graph.RelMemberOf {
			continue
		}
		targetIndex, ok := nodeIndex[rel.TargetID]
		if !ok {
			continue
		}
		target := g.Nodes[targetIndex]
		if target.Label != scopeir.NodeProcess && target.Label != scopeir.NodeCommunity {
			continue
		}
		sourceArea := nodeAreas[rel.SourceID]
		if sourceArea == "" || sourceArea == FunctionalAreaUnknown {
			continue
		}
		if _, ok := areaSets[rel.TargetID]; !ok {
			areaSets[rel.TargetID] = map[FunctionalArea]struct{}{}
		}
		areaSets[rel.TargetID][sourceArea] = struct{}{}
	}

	out := map[string]FunctionalAreaClassification{}
	for nodeID, areas := range areaSets {
		if len(areas) == 0 {
			continue
		}
		if len(areas) == 1 {
			for area := range areas {
				out[nodeID] = FunctionalAreaClassification{Area: area, Source: "relationship_membership"}
			}
			continue
		}
		out[nodeID] = FunctionalAreaClassification{Area: FunctionalAreaMixed, Source: "mixed_relationship_membership"}
	}
	return out
}

func normalizePath(filePath string) string {
	path := strings.TrimSpace(filePath)
	path = strings.ReplaceAll(path, "\\", "/")
	path = strings.TrimPrefix(path, "./")
	path = strings.Trim(path, "/")
	return strings.ToLower(path)
}

func hasPathPrefix(path string, prefix string) bool {
	prefix = strings.Trim(prefix, "/")
	return path == prefix || strings.HasPrefix(path, prefix+"/") || strings.Contains(path, "/"+prefix+"/")
}

func pathContains(path string, segment string) bool {
	segment = strings.Trim(segment, "/")
	return path == segment || strings.Contains(path, "/"+segment+"/") || strings.HasPrefix(path, segment+"/")
}

func isDocsPath(path string, base string) bool {
	return hasPathPrefix(path, "docs") ||
		hasPathPrefix(path, "reports") ||
		strings.HasSuffix(base, ".md") ||
		strings.HasSuffix(base, ".mdx") ||
		strings.HasSuffix(base, ".rst") ||
		strings.HasSuffix(base, ".txt")
}

func isWebTestPath(path string, base string) bool {
	if !hasPathPrefix(path, "avmatrix-web") {
		return false
	}
	return pathContains(path, "avmatrix-web/test") ||
		pathContains(path, "avmatrix-web/tests") ||
		pathContains(path, "avmatrix-web/e2e") ||
		strings.Contains(base, ".test.") ||
		strings.Contains(base, ".spec.")
}

func isAPITestPath(path string, base string) bool {
	if !strings.HasSuffix(base, "_test.go") {
		return false
	}
	return hasPathPrefix(path, "internal/httpapi") ||
		hasPathPrefix(path, "internal/mcp") ||
		hasPathPrefix(path, "internal/contracts") ||
		hasPathPrefix(path, "cmd/generate-web-contracts")
}

func isBackendTestPath(path string, base string) bool {
	return strings.HasSuffix(base, "_test.go") ||
		pathContains(path, "test/fixtures") ||
		pathContains(path, "tests/fixtures")
}

func isGeneratedContractPath(path string) bool {
	return hasPathPrefix(path, "cmd/generate-web-contracts") ||
		pathContains(path, "avmatrix-web/src/generated") ||
		pathContains(path, "contracts/web-ui")
}

func isAPIContractPath(path string) bool {
	return hasPathPrefix(path, "internal/contracts") ||
		hasPathPrefix(path, "contracts")
}

func isFrontendAPIClientPath(path string) bool {
	return path == "avmatrix-web/src/services/backend-client.ts" ||
		strings.HasSuffix(path, "/avmatrix-web/src/services/backend-client.ts")
}

func isConfigPath(path string, base string) bool {
	if strings.HasPrefix(base, ".env") {
		return true
	}
	switch base {
	case ".gitignore", ".npmrc", ".nvmrc", "dockerfile", "makefile", "go.mod", "go.sum",
		"package.json", "package-lock.json", "pnpm-lock.yaml", "yarn.lock", "bun.lockb":
		return true
	}
	if strings.HasPrefix(base, "tsconfig") ||
		strings.HasPrefix(base, "vite.config") ||
		strings.HasPrefix(base, "vitest.config") ||
		strings.HasPrefix(base, "playwright.config") ||
		strings.HasPrefix(base, "eslint.config") ||
		strings.HasPrefix(base, "postcss.config") ||
		strings.HasPrefix(base, "tailwind.config") {
		return true
	}
	return strings.HasSuffix(base, ".yaml") ||
		strings.HasSuffix(base, ".yml") ||
		strings.HasSuffix(base, ".toml") ||
		strings.HasSuffix(base, ".ini")
}

func isGeneratedPath(path string) bool {
	return pathContains(path, "generated") ||
		pathContains(path, "dist") ||
		pathContains(path, "build") ||
		pathContains(path, "out")
}

func isAPIPath(path string) bool {
	return hasPathPrefix(path, "internal/httpapi") ||
		hasPathPrefix(path, "internal/mcp") ||
		pathContains(path, "app/api") ||
		pathContains(path, "pages/api")
}

func isFrontendPath(path string) bool {
	return hasPathPrefix(path, "avmatrix-web/src") ||
		hasPathPrefix(path, "avmatrix-web/public") ||
		hasPathPrefix(path, "avmatrix-web/app") ||
		path == "avmatrix-web/index.html" ||
		strings.HasSuffix(path, "/avmatrix-web/index.html")
}

func isCLILauncherPath(path string, base string) bool {
	return hasPathPrefix(path, "cmd") ||
		hasPathPrefix(path, "avmatrix-launcher") ||
		base == "start-avmatrix.html"
}

func isBackendPath(path string) bool {
	return hasPathPrefix(path, "internal") ||
		hasPathPrefix(path, "pkg")
}

func stringProperty(props graph.NodeProperties, key string) string {
	value, ok := props[key]
	if !ok || value == nil {
		return ""
	}
	text, _ := value.(string)
	return text
}
