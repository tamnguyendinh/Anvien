package graphhealth

import (
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

	// 1. Build counted degree maps using only IsCounted relationships.
	countedIn := make(map[string]int, len(g.Nodes))
	countedOut := make(map[string]int, len(g.Nodes))
	excludedByNode := make(map[string]map[string]int, len(g.Nodes))
	for _, rel := range g.Relationships {
		if IsCounted(rel.Type) {
			countedIn[rel.TargetID]++
			countedOut[rel.SourceID]++
			summary.CountedRelationshipCount++
			continue
		}
		category := excludedEdgeCategory(rel.Type)
		summary.ExcludedEdgeCounts[category]++
		incrementNestedCount(excludedByNode, rel.SourceID, category)
		incrementNestedCount(excludedByNode, rel.TargetID, category)
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
		if label == scopeir.NodeRoute || label == scopeir.NodeTool || strings.HasPrefix(strings.ToLower(name), "main") {
			reasons = append(reasons, ReasonFrameworkEntry)
		}
		// CLI/MCP rough: path based for stub
		if strings.Contains(lower, "/cmd/") || strings.Contains(lower, "/internal/cli/") || strings.Contains(lower, "/internal/mcp/") {
			reasons = append(reasons, ReasonCLIMCP)
		}
		return reasons
	}

	// 3. Assign topology + confidence per node (simple version; detached deferred to P2-D).
	for i := range g.Nodes {
		n := &g.Nodes[i]
		in := countedIn[n.ID]
		out := countedOut[n.ID]
		reasons := isExpected(*n)

		var status TopologyStatus
		switch {
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
		if status == TopologyUnknown {
			conf = ConfidenceUnknown
		}

		health := NodeHealth{
			TopologyStatus:           status,
			CountedIncoming:          in,
			CountedOutgoing:          out,
			ExcludedEdgeCounts:       cloneCounts(excludedByNode[n.ID]),
			ExpectedIsolationReasons: reasons,
			Confidence:               conf,
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
		n.Properties["expectedIsolationReasons"] = health.ExpectedIsolationReasons
		n.Properties["confidence"] = health.Confidence
		// Also embed full for typed consumers
		n.Properties["graphHealth"] = health
		addNodeHealthToSummary(&summary, health)
	}

	// P2-D detached_component and P2-E unresolved diagnostics are stubs here.
	// Full version will compute components and attach "detached" + root explanations.
	return summary
}

// --- local helpers (duplicated from processes/ignore for Phase 2 starter; later factor) ---

func newSummary() Summary {
	summary := Summary{
		PolicyVersion:                 PolicyVersion,
		TopologyStatusCounts:          map[string]int{},
		ExpectedIsolationReasonCounts: map[string]int{},
		ConfidenceCounts:              map[string]int{},
		DiagnosticCounts:              map[string]int{},
		ExcludedEdgeCounts:            map[string]int{},
	}
	for _, status := range TopologyStatuses {
		summary.TopologyStatusCounts[string(status)] = 0
	}
	for _, confidence := range ConfidenceLevels {
		summary.ConfidenceCounts[confidence] = 0
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
		summary.DiagnosticCounts[diagnostic.Kind]++
	}
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
