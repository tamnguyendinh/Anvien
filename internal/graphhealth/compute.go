package graphhealth

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/tamnguyendinh/avmatrix-go/internal/graph"
	"github.com/tamnguyendinh/avmatrix-go/internal/scopeir"
)

// Compute annotates the provided graph in-place with graph-health metadata
// under each node's Properties["graphHealth"] using the accepted policies.
//
// Ownership: core layer (see P1-H). Callers (httpapi, mcp, reports) invoke this
// after loading a full graph and before emitting to consumers.
func Compute(g *graph.Graph) {
	_ = ComputeSummary(g)
}

// ComputeSummary annotates the graph and returns graph-level inventory counts.
func ComputeSummary(g *graph.Graph) Summary {
	summary := newSummary()
	if g == nil {
		return summary
	}
	summary.NodeCount = len(g.Nodes)
	metadataUnresolved, metadataSourceBacked, metadataUnattributed, hasResolutionMetadata := resolutionMetadata(g)
	if hasResolutionMetadata {
		summary.UnresolvedReferenceCount = metadataUnresolved
		summary.SourceBackedUnresolvedReferenceCount = metadataSourceBacked
		summary.UnattributedUnresolvedReferenceCount = metadataUnattributed
	}

	// 1. Build counted degree maps using only IsCounted relationships.
	nodeIDs := make([]string, 0, len(g.Nodes))
	nodeByID := make(map[string]graph.Node, len(g.Nodes))
	for _, node := range g.Nodes {
		nodeIDs = append(nodeIDs, node.ID)
		nodeByID[node.ID] = node
	}
	sort.Strings(nodeIDs)

	countedIn := make(map[string]int, len(g.Nodes))
	countedOut := make(map[string]int, len(g.Nodes))
	excludedByNode := make(map[string]map[string]int, len(g.Nodes))
	directedOut := make(map[string][]string, len(g.Nodes))
	weakAdjacency := make(map[string][]string, len(g.Nodes))
	rootNodeIDs := make(map[string]bool, len(g.Nodes))
	for _, rel := range g.Relationships {
		if IsCounted(rel.Type) {
			countedIn[rel.TargetID]++
			countedOut[rel.SourceID]++
			summary.CountedRelationshipCount++
			if _, ok := nodeByID[rel.SourceID]; ok {
				if _, ok := nodeByID[rel.TargetID]; ok {
					directedOut[rel.SourceID] = append(directedOut[rel.SourceID], rel.TargetID)
					weakAdjacency[rel.SourceID] = append(weakAdjacency[rel.SourceID], rel.TargetID)
					weakAdjacency[rel.TargetID] = append(weakAdjacency[rel.TargetID], rel.SourceID)
				}
			}
			if rel.Type == graph.RelEntryPointOf || rel.Type == graph.RelHandlesRoute || rel.Type == graph.RelHandlesTool {
				if _, ok := nodeByID[rel.SourceID]; ok {
					rootNodeIDs[rel.SourceID] = true
				}
			}
			continue
		}
		category := excludedEdgeCategory(rel.Type)
		summary.ExcludedEdgeCounts[category]++
		incrementNestedCount(excludedByNode, rel.SourceID, category)
		incrementNestedCount(excludedByNode, rel.TargetID, category)
	}
	for _, targets := range directedOut {
		sort.Strings(targets)
	}
	for _, neighbors := range weakAdjacency {
		sort.Strings(neighbors)
	}

	// 2. Expected-isolation heuristic bridged from existing path/process policies.
	isExpected := func(n graph.Node) []string {
		reasons := []string{}
		fp := stringProperty(n, "filePath")
		lower := strings.ToLower(strings.ReplaceAll(fp, "\\", "/"))
		name := stringProperty(n, "name")
		label := n.Label

		if isTestPath(lower, name) {
			reasons = append(reasons, ReasonTest)
		}
		if strings.Contains(lower, "/fixtures/") || strings.Contains(lower, "/__snapshots__/") || strings.Contains(lower, "/snapshots/") {
			reasons = append(reasons, ReasonFixture)
		}
		if strings.Contains(lower, "/generated/") || strings.Contains(lower, ".generated.") {
			reasons = append(reasons, ReasonGenerated)
		}
		if strings.Contains(lower, "/vendor/") || strings.Contains(lower, "/node_modules/") {
			reasons = append(reasons, ReasonVendor)
		}
		if label == scopeir.NodeSection || strings.HasSuffix(lower, ".md") || strings.Contains(lower, "/docs/") {
			reasons = append(reasons, ReasonDocumentation)
		}
		if strings.Contains(lower, "/migrations/") || strings.HasSuffix(lower, ".sql") && strings.Contains(lower, "migrate") {
			reasons = append(reasons, ReasonMigration)
		}
		if boolProperty(n, "isExported") {
			reasons = append(reasons, ReasonExportedAPI) // modifier, not auto-hide
		}
		// CLI/MCP rough: path based for stub
		if strings.Contains(lower, "/cmd/") || strings.Contains(lower, "/internal/cli/") || strings.Contains(lower, "/internal/mcp/") {
			reasons = append(reasons, ReasonCLIMCP)
		}
		return reasons
	}

	for _, node := range g.Nodes {
		if isAcceptedRoot(node) {
			rootNodeIDs[node.ID] = true
		}
	}
	reachableFromRoot := reachableNodes(rootNodeIDs, directedOut)
	componentByNode, components := computeComponents(nodeIDs, weakAdjacency, rootNodeIDs, reachableFromRoot)
	summary.ComponentCount = len(components)
	summary.RootNodeCount = len(rootNodeIDs)
	for _, component := range components {
		if component.Detached {
			summary.DetachedComponentCount++
		}
	}
	summary.LargestDetachedComponents = largestDetachedComponents(components, 10)

	// 3. Assign topology + confidence per node.
	for i := range g.Nodes {
		n := &g.Nodes[i]
		in := countedIn[n.ID]
		out := countedOut[n.ID]
		reasons := isExpected(*n)
		if rootNodeIDs[n.ID] {
			reasons = appendReason(reasons, ReasonFrameworkEntry)
		}
		component := components[componentByNode[n.ID]]
		diagnostics := diagnosticsFromProperties(n.Properties)
		hasUnresolvedDiagnostic := hasDiagnosticKind(diagnostics, DiagnosticUnresolvedReference)

		var status TopologyStatus
		switch {
		case component.Detached:
			status = TopologyDetached
		case in > 0 && out > 0:
			status = TopologyConnected
		case in == 0 && out == 0:
			status = TopologyTrueIsolated
		case in == 0 && out > 0:
			status = TopologyNoIncoming
		case in > 0 && out == 0:
			status = TopologyNoOutgoing
		default:
			status = TopologyUnknown
		}

		conf := ConfidenceCandidate
		if hasAutomaticExpectedReason(reasons) {
			conf = ConfidenceExpected
		}
		if hasUnresolvedDiagnostic || status == TopologyUnknown {
			conf = ConfidenceUnknown
		}

		health := NodeHealth{
			TopologyStatus:             status,
			CountedIncoming:            in,
			CountedOutgoing:            out,
			ExcludedEdgeCounts:         cloneCounts(excludedByNode[n.ID]),
			ComponentID:                component.ID,
			ComponentSize:              component.NodeCount,
			ComponentReachableFromRoot: component.ReachableFromRoot,
			ExpectedIsolationReasons:   reasons,
			Diagnostics:                diagnostics,
			Confidence:                 conf,
		}
		// Attach for consumers (flat + structured for easy access)
		if n.Properties == nil {
			n.Properties = make(graph.NodeProperties)
		}
		n.Properties["topologyStatus"] = string(health.TopologyStatus)
		n.Properties["countedIncoming"] = health.CountedIncoming
		n.Properties["countedOutgoing"] = health.CountedOutgoing
		if len(health.ExcludedEdgeCounts) > 0 {
			n.Properties["excludedEdgeCounts"] = health.ExcludedEdgeCounts
		} else {
			delete(n.Properties, "excludedEdgeCounts")
		}
		n.Properties["componentId"] = health.ComponentID
		n.Properties["componentSize"] = health.ComponentSize
		n.Properties["componentReachableFromRoot"] = health.ComponentReachableFromRoot
		delete(n.Properties, "componentRootNodeIds")
		n.Properties["expectedIsolationReasons"] = health.ExpectedIsolationReasons
		if len(health.Diagnostics) > 0 {
			n.Properties["diagnostics"] = health.Diagnostics
		} else {
			delete(n.Properties, "diagnostics")
		}
		n.Properties["confidence"] = health.Confidence
		// Also embed full for typed consumers
		n.Properties["graphHealth"] = health
		addNodeHealthToSummary(&summary, health)
	}

	sourceBackedFromDiagnostics := summary.DiagnosticCounts[DiagnosticUnresolvedReference]
	if sourceBackedFromDiagnostics > summary.SourceBackedUnresolvedReferenceCount {
		summary.SourceBackedUnresolvedReferenceCount = sourceBackedFromDiagnostics
	}
	if summary.UnresolvedReferenceCount < summary.SourceBackedUnresolvedReferenceCount {
		summary.UnresolvedReferenceCount = summary.SourceBackedUnresolvedReferenceCount
	}
	if summary.UnresolvedReferenceCount >= summary.SourceBackedUnresolvedReferenceCount {
		summary.UnattributedUnresolvedReferenceCount = summary.UnresolvedReferenceCount - summary.SourceBackedUnresolvedReferenceCount
	}
	return summary
}

// --- local helpers (duplicated from processes/ignore for Phase 2 starter; later factor) ---

func newSummary() Summary {
	summary := Summary{
		PolicyVersion:                  PolicyVersion,
		TopologyStatusCounts:           map[string]int{},
		ExpectedIsolationReasonCounts:  map[string]int{},
		ConfidenceCounts:               map[string]int{},
		DiagnosticCounts:               map[string]int{},
		DiagnosticClassificationCounts: map[string]int{},
		DiagnosticActionabilityCounts:  map[string]int{},
		ExcludedEdgeCounts:             map[string]int{},
	}
	for _, status := range TopologyStatuses {
		summary.TopologyStatusCounts[string(status)] = 0
	}
	for _, confidence := range ConfidenceLevels {
		summary.ConfidenceCounts[confidence] = 0
	}
	for _, classification := range DiagnosticClassifications {
		summary.DiagnosticClassificationCounts[classification] = 0
	}
	for _, actionability := range DiagnosticActionabilities {
		summary.DiagnosticActionabilityCounts[actionability] = 0
	}
	return summary
}

func addNodeHealthToSummary(summary *Summary, health NodeHealth) {
	summary.TopologyStatusCounts[string(health.TopologyStatus)]++
	summary.ConfidenceCounts[health.Confidence]++
	for _, reason := range health.ExpectedIsolationReasons {
		summary.ExpectedIsolationReasonCounts[reason]++
	}
	for _, diagnostic := range health.Diagnostics {
		if diagnostic.Kind == "" {
			continue
		}
		summary.DiagnosticCounts[diagnostic.Kind] += diagnosticCount(diagnostic)
		if diagnostic.Classification != "" {
			summary.DiagnosticClassificationCounts[diagnostic.Classification] += diagnosticCount(diagnostic)
		}
		if diagnostic.Actionability != "" {
			summary.DiagnosticActionabilityCounts[diagnostic.Actionability] += diagnosticCount(diagnostic)
		}
	}
}

func hasDiagnosticKind(diagnostics []Diagnostic, kind string) bool {
	for _, diagnostic := range diagnostics {
		if diagnostic.Kind == kind {
			return true
		}
	}
	return false
}

type componentInfo struct {
	ID                string
	NodeIDs           []string
	NodeCount         int
	CountedEdgeCount  int
	RootNodeIDs       []string
	ReachableFromRoot bool
	Detached          bool
}

func computeComponents(nodeIDs []string, weakAdjacency map[string][]string, rootNodeIDs map[string]bool, reachableFromRoot map[string]bool) (map[string]int, []componentInfo) {
	seen := make(map[string]bool, len(nodeIDs))
	componentByNode := make(map[string]int, len(nodeIDs))
	components := make([]componentInfo, 0)
	for _, nodeID := range nodeIDs {
		if seen[nodeID] {
			continue
		}
		queue := []string{nodeID}
		seen[nodeID] = true
		nodes := make([]string, 0)
		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]
			nodes = append(nodes, current)
			for _, next := range weakAdjacency[current] {
				if seen[next] {
					continue
				}
				seen[next] = true
				queue = append(queue, next)
			}
		}
		sort.Strings(nodes)
		component := componentInfo{
			ID:        fmt.Sprintf("component_%06d", len(components)+1),
			NodeIDs:   nodes,
			NodeCount: len(nodes),
		}
		for _, id := range nodes {
			componentByNode[id] = len(components)
			component.CountedEdgeCount += countedNeighborEdges(id, weakAdjacency)
			if rootNodeIDs[id] {
				component.RootNodeIDs = append(component.RootNodeIDs, id)
			}
			if reachableFromRoot[id] {
				component.ReachableFromRoot = true
			}
		}
		component.CountedEdgeCount /= 2
		sort.Strings(component.RootNodeIDs)
		if len(component.RootNodeIDs) > 0 {
			component.ReachableFromRoot = true
		}
		component.Detached = component.CountedEdgeCount > 0 && !component.ReachableFromRoot
		components = append(components, component)
	}
	return componentByNode, components
}

func countedNeighborEdges(nodeID string, weakAdjacency map[string][]string) int {
	return len(weakAdjacency[nodeID])
}

func reachableNodes(rootNodeIDs map[string]bool, directedOut map[string][]string) map[string]bool {
	reachable := make(map[string]bool, len(rootNodeIDs))
	queue := make([]string, 0, len(rootNodeIDs))
	for rootID := range rootNodeIDs {
		reachable[rootID] = true
		queue = append(queue, rootID)
	}
	sort.Strings(queue)
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		for _, next := range directedOut[current] {
			if reachable[next] {
				continue
			}
			reachable[next] = true
			queue = append(queue, next)
		}
	}
	return reachable
}

func largestDetachedComponents(components []componentInfo, limit int) []ComponentSummary {
	out := make([]ComponentSummary, 0)
	for _, component := range components {
		if !component.Detached {
			continue
		}
		out = append(out, componentSummary(component))
	}
	sort.Slice(out, func(i int, j int) bool {
		if out[i].NodeCount != out[j].NodeCount {
			return out[i].NodeCount > out[j].NodeCount
		}
		return out[i].ID < out[j].ID
	})
	if len(out) > limit {
		out = out[:limit]
	}
	return out
}

func componentSummary(component componentInfo) ComponentSummary {
	return ComponentSummary{
		ID:                component.ID,
		NodeCount:         component.NodeCount,
		CountedEdgeCount:  component.CountedEdgeCount,
		Detached:          component.Detached,
		ReachableFromRoot: component.ReachableFromRoot,
		RootNodeIDs:       cloneStrings(component.RootNodeIDs),
		SampleNodeIDs:     firstStrings(component.NodeIDs, 5),
	}
}

func firstStrings(values []string, limit int) []string {
	if len(values) == 0 {
		return nil
	}
	if len(values) < limit {
		limit = len(values)
	}
	out := make([]string, limit)
	copy(out, values[:limit])
	return out
}

func incrementNestedCount(counts map[string]map[string]int, nodeID string, category string) {
	if nodeID == "" {
		return
	}
	if counts[nodeID] == nil {
		counts[nodeID] = map[string]int{}
	}
	counts[nodeID][category]++
}

func cloneCounts(counts map[string]int) map[string]int {
	if len(counts) == 0 {
		return nil
	}
	out := make(map[string]int, len(counts))
	for key, value := range counts {
		out[key] = value
	}
	return out
}

func cloneStrings(values []string) []string {
	if len(values) == 0 {
		return nil
	}
	out := make([]string, len(values))
	copy(out, values)
	return out
}

func excludedEdgeCategory(t graph.RelationshipType) string {
	if StructuralEdgeTypes[t] {
		return ExcludedEdgeStructural
	}
	return ExcludedEdgeOther
}

func hasAutomaticExpectedReason(reasons []string) bool {
	for _, reason := range reasons {
		if reason != ReasonExportedAPI {
			return true
		}
	}
	return false
}

func appendReason(reasons []string, reason string) []string {
	for _, existing := range reasons {
		if existing == reason {
			return reasons
		}
	}
	return append(reasons, reason)
}

func isAcceptedRoot(n graph.Node) bool {
	switch n.Label {
	case scopeir.NodeProcess, scopeir.NodeRoute, scopeir.NodeTool:
		return true
	case scopeir.NodeFunction, scopeir.NodeMethod:
		name := strings.ToLower(strings.TrimSpace(stringProperty(n, "name")))
		if !isMainLikeName(name) {
			return false
		}
		return boolProperty(n, "isExported") || floatProperty(n, "astFrameworkMultiplier") > 1
	default:
		return false
	}
}

func isMainLikeName(name string) bool {
	switch name {
	case "main", "init", "run", "start", "bootstrap":
		return true
	default:
		return false
	}
}

func stringProperty(n graph.Node, key string) string {
	if n.Properties == nil {
		return ""
	}
	if v, ok := n.Properties[key]; ok {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

func boolProperty(n graph.Node, key string) bool {
	if n.Properties == nil {
		return false
	}
	if v, ok := n.Properties[key]; ok {
		if b, ok := v.(bool); ok {
			return b
		}
	}
	return false
}

func floatProperty(n graph.Node, key string) float64 {
	if n.Properties == nil {
		return 0
	}
	if v, ok := n.Properties[key]; ok {
		switch value := v.(type) {
		case float64:
			return value
		case float32:
			return float64(value)
		case int:
			return float64(value)
		case int64:
			return float64(value)
		case json.Number:
			out, _ := value.Float64()
			return out
		}
	}
	return 0
}

func isTestPath(lowerPath, name string) bool {
	return strings.Contains(lowerPath, ".test.") ||
		strings.Contains(lowerPath, ".spec.") ||
		strings.Contains(lowerPath, "/test/") ||
		strings.Contains(lowerPath, "/tests/") ||
		strings.Contains(lowerPath, "/__tests__/") ||
		strings.HasPrefix(lowerPath, "test/") ||
		strings.HasPrefix(lowerPath, "tests/") ||
		strings.HasSuffix(lowerPath, "_test.go") ||
		strings.HasSuffix(lowerPath, "_test.py")
}
