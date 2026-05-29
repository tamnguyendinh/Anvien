package graphhealth

import (
	"sort"

	"github.com/tamnguyendinh/anvien/internal/graph"
)

const (
	ExplainSampleLimit        = 20
	ReportDefaultLimit        = 100
	ReportMaxLimit            = 1000
	TriageDimensionTopology   = "topology"
	TriageDimensionDiagnostic = "diagnostic"
)

type ReportOptions struct {
	Limit           int
	IncludeExpected bool
}

type ExplainResponse struct {
	Kind                         string                `json:"kind"`
	NodeID                       string                `json:"nodeId,omitempty"`
	ComponentID                  string                `json:"componentId,omitempty"`
	Node                         *graph.Node           `json:"node,omitempty"`
	Health                       *NodeHealth           `json:"health,omitempty"`
	Component                    *ComponentExplanation `json:"component,omitempty"`
	CountedIncomingRelationships []graph.Relationship  `json:"countedIncomingRelationships,omitempty"`
	CountedOutgoingRelationships []graph.Relationship  `json:"countedOutgoingRelationships,omitempty"`
	ExcludedRelationships        []graph.Relationship  `json:"excludedRelationships,omitempty"`
	SampleNodes                  []graph.Node          `json:"sampleNodes,omitempty"`
	CountedRelationshipSamples   []graph.Relationship  `json:"countedRelationshipSamples,omitempty"`
	ExcludedRelationshipSamples  []graph.Relationship  `json:"excludedRelationshipSamples,omitempty"`
	SampleLimit                  int                   `json:"sampleLimit,omitempty"`
}

type ComponentExplanation struct {
	ID                             string         `json:"id"`
	NodeCount                      int            `json:"nodeCount"`
	CountedEdgeCount               int            `json:"countedEdgeCount"`
	Detached                       bool           `json:"detached"`
	ReachableFromRoot              bool           `json:"reachableFromRoot"`
	RootNodeIDs                    []string       `json:"rootNodeIds,omitempty"`
	SampleNodeIDs                  []string       `json:"sampleNodeIds,omitempty"`
	TopologyStatusCounts           map[string]int `json:"topologyStatusCounts"`
	ExpectedIsolationReasonCounts  map[string]int `json:"expectedIsolationReasonCounts"`
	ConfidenceCounts               map[string]int `json:"confidenceCounts"`
	DiagnosticCounts               map[string]int `json:"diagnosticCounts"`
	DiagnosticClassificationCounts map[string]int `json:"diagnosticClassificationCounts"`
	DiagnosticActionabilityCounts  map[string]int `json:"diagnosticActionabilityCounts"`
	ResolutionGapCount             int            `json:"resolutionGapCount"`
	ResolutionHealthBucketCounts   map[string]int `json:"resolutionHealthBucketCounts"`
	ResolutionConfidenceCounts     map[string]int `json:"resolutionConfidenceCounts"`
}

type ReportResponse struct {
	ReportType         string            `json:"reportType"`
	VerdictPolicy      string            `json:"verdictPolicy"`
	Limit              int               `json:"limit"`
	IncludeExpected    bool              `json:"includeExpected"`
	Summary            Summary           `json:"summary"`
	TotalCandidates    int               `json:"totalCandidates"`
	ReturnedCandidates int               `json:"returnedCandidates"`
	Candidates         []ReportCandidate `json:"candidates"`
}

type ReportCandidate struct {
	NodeID                     string         `json:"nodeId"`
	Label                      string         `json:"label"`
	Name                       string         `json:"name,omitempty"`
	FilePath                   string         `json:"filePath,omitempty"`
	TriagePriority             string         `json:"triagePriority"`
	TriageDimension            string         `json:"triageDimension"`
	TopologyStatus             TopologyStatus `json:"topologyStatus"`
	Confidence                 string         `json:"confidence"`
	CountedIncoming            int            `json:"countedIncoming"`
	CountedOutgoing            int            `json:"countedOutgoing"`
	ExcludedEdgeCounts         map[string]int `json:"excludedEdgeCounts,omitempty"`
	ExpectedIsolationReasons   []string       `json:"expectedIsolationReasons,omitempty"`
	Diagnostics                []Diagnostic   `json:"diagnostics,omitempty"`
	ComponentID                string         `json:"componentId,omitempty"`
	ComponentSize              int            `json:"componentSize,omitempty"`
	ComponentReachableFromRoot bool           `json:"componentReachableFromRoot"`
	ResolutionHealthBuckets    map[string]int `json:"resolutionHealthBuckets,omitempty"`
	ResolutionGapCount         int            `json:"resolutionGapCount,omitempty"`
	ResolutionConfidence       string         `json:"resolutionConfidence"`
}

func BuildReport(g *graph.Graph, options ReportOptions) ReportResponse {
	summary := ComputeSummary(g)
	limit := NormalizeReportLimit(options.Limit)
	candidates := reportCandidatesFromComputed(g, options.IncludeExpected)
	totalCandidates := len(candidates)
	if len(candidates) > limit {
		candidates = candidates[:limit]
	}
	return ReportResponse{
		ReportType:         "graph_health_candidate_review",
		VerdictPolicy:      "candidate_not_confirmed",
		Limit:              limit,
		IncludeExpected:    options.IncludeExpected,
		Summary:            summary,
		TotalCandidates:    totalCandidates,
		ReturnedCandidates: len(candidates),
		Candidates:         candidates,
	}
}

func NormalizeReportLimit(limit int) int {
	if limit < 1 {
		return ReportDefaultLimit
	}
	if limit > ReportMaxLimit {
		return ReportMaxLimit
	}
	return limit
}

func ExplainNode(g *graph.Graph, nodeID string) (ExplainResponse, bool) {
	if g == nil {
		return ExplainResponse{}, false
	}
	Compute(g)
	node, ok := g.GetNode(nodeID)
	if !ok {
		return ExplainResponse{}, false
	}
	health, ok := nodeHealthFromNode(node)
	if !ok {
		return ExplainResponse{}, false
	}
	incoming, outgoing, excluded := nodeRelationships(g, nodeID)
	return ExplainResponse{
		Kind:                         "node",
		NodeID:                       nodeID,
		ComponentID:                  health.ComponentID,
		Node:                         &node,
		Health:                       &health,
		CountedIncomingRelationships: incoming,
		CountedOutgoingRelationships: outgoing,
		ExcludedRelationships:        excluded,
	}, true
}

func ExplainComponent(g *graph.Graph, componentID string) (ExplainResponse, bool) {
	if g == nil {
		return ExplainResponse{}, false
	}
	Compute(g)
	component, sampleNodes, countedSamples, excludedSamples, ok := componentExplanation(g, componentID, ExplainSampleLimit)
	if !ok {
		return ExplainResponse{}, false
	}
	return ExplainResponse{
		Kind:                        "component",
		ComponentID:                 componentID,
		Component:                   &component,
		SampleNodes:                 sampleNodes,
		CountedRelationshipSamples:  countedSamples,
		ExcludedRelationshipSamples: excludedSamples,
		SampleLimit:                 ExplainSampleLimit,
	}, true
}

func ComponentSummaries(g *graph.Graph) []ComponentExplanation {
	if g == nil {
		return nil
	}
	Compute(g)
	components := map[string]*ComponentExplanation{}
	nodeComponent := map[string]string{}
	for _, node := range g.Nodes {
		health, ok := nodeHealthFromNode(node)
		if !ok || health.ComponentID == "" {
			continue
		}
		component := components[health.ComponentID]
		if component == nil {
			next := newComponentExplanation(health.ComponentID)
			component = &next
			components[health.ComponentID] = component
		}
		nodeComponent[node.ID] = health.ComponentID
		addNodeToComponent(component, node.ID, health)
	}
	for _, relationship := range g.Relationships {
		if !IsCounted(relationship.Type) {
			continue
		}
		sourceComponent := nodeComponent[relationship.SourceID]
		if sourceComponent == "" || sourceComponent != nodeComponent[relationship.TargetID] {
			continue
		}
		components[sourceComponent].CountedEdgeCount++
	}
	out := make([]ComponentExplanation, 0, len(components))
	for _, component := range components {
		finalizeComponent(component)
		out = append(out, *component)
	}
	sort.SliceStable(out, func(i, j int) bool {
		if out[i].Detached != out[j].Detached {
			return out[i].Detached
		}
		if out[i].NodeCount != out[j].NodeCount {
			return out[i].NodeCount > out[j].NodeCount
		}
		return out[i].ID < out[j].ID
	})
	return out
}

func ReportCandidates(g *graph.Graph, includeExpected bool) []ReportCandidate {
	Compute(g)
	return reportCandidatesFromComputed(g, includeExpected)
}

func reportCandidatesFromComputed(g *graph.Graph, includeExpected bool) []ReportCandidate {
	if g == nil {
		return nil
	}
	candidates := make([]ReportCandidate, 0)
	for _, node := range g.Nodes {
		health, ok := nodeHealthFromNode(node)
		if !ok {
			continue
		}
		priority, dimension, rank := ReportPriority(health)
		if rank == 0 {
			continue
		}
		if !includeExpected && health.Confidence == ConfidenceExpected {
			continue
		}
		candidates = append(candidates, ReportCandidate{
			NodeID:                     node.ID,
			Label:                      string(node.Label),
			Name:                       stringProperty(node, "name"),
			FilePath:                   stringProperty(node, "filePath"),
			TriagePriority:             priority,
			TriageDimension:            dimension,
			TopologyStatus:             health.TopologyStatus,
			Confidence:                 health.Confidence,
			CountedIncoming:            health.CountedIncoming,
			CountedOutgoing:            health.CountedOutgoing,
			ExcludedEdgeCounts:         health.ExcludedEdgeCounts,
			ExpectedIsolationReasons:   health.ExpectedIsolationReasons,
			Diagnostics:                health.Diagnostics,
			ComponentID:                health.ComponentID,
			ComponentSize:              health.ComponentSize,
			ComponentReachableFromRoot: health.ComponentReachableFromRoot,
			ResolutionHealthBuckets:    health.ResolutionHealthBuckets,
			ResolutionGapCount:         health.ResolutionGapCount,
			ResolutionConfidence:       health.ResolutionConfidence,
		})
	}
	sort.SliceStable(candidates, func(i, j int) bool {
		leftRank := ReportPriorityRank(candidates[i].TriagePriority, candidates[i].TriageDimension)
		rightRank := ReportPriorityRank(candidates[j].TriagePriority, candidates[j].TriageDimension)
		if leftRank != rightRank {
			return leftRank < rightRank
		}
		if candidates[i].Confidence != candidates[j].Confidence {
			return candidates[i].Confidence < candidates[j].Confidence
		}
		if candidates[i].FilePath != candidates[j].FilePath {
			return candidates[i].FilePath < candidates[j].FilePath
		}
		return candidates[i].NodeID < candidates[j].NodeID
	})
	return candidates
}

func ReportPriority(health NodeHealth) (string, string, int) {
	switch health.TopologyStatus {
	case TopologyNoIncoming:
		return string(TopologyNoIncoming), TriageDimensionTopology, 1
	case TopologyDetached:
		return string(TopologyDetached), TriageDimensionTopology, 2
	case TopologyTrueIsolated:
		return string(TopologyTrueIsolated), TriageDimensionTopology, 4
	case TopologyNoOutgoing:
		return string(TopologyNoOutgoing), TriageDimensionTopology, 5
	case TopologyUnknown:
		return string(TopologyUnknown), TriageDimensionTopology, 6
	}
	if hasDiagnosticKind(health.Diagnostics, DiagnosticUnresolvedReference) {
		return DiagnosticUnresolvedReference, TriageDimensionDiagnostic, 30
	}
	return "", "", 0
}

func ReportPriorityRank(priority string, dimension string) int {
	if dimension == TriageDimensionDiagnostic {
		switch priority {
		case DiagnosticUnresolvedReference:
			return 30
		default:
			return 90
		}
	}
	switch priority {
	case string(TopologyNoIncoming):
		return 1
	case string(TopologyDetached):
		return 2
	case string(TopologyTrueIsolated):
		return 4
	case string(TopologyNoOutgoing):
		return 5
	case string(TopologyUnknown):
		return 6
	default:
		return 99
	}
}

func componentExplanation(g *graph.Graph, componentID string, sampleLimit int) (ComponentExplanation, []graph.Node, []graph.Relationship, []graph.Relationship, bool) {
	component := newComponentExplanation(componentID)
	componentNodeIDs := map[string]bool{}
	sampleNodes := make([]graph.Node, 0, sampleLimit)
	for _, node := range g.Nodes {
		health, ok := nodeHealthFromNode(node)
		if !ok || health.ComponentID != componentID {
			continue
		}
		componentNodeIDs[node.ID] = true
		addNodeToComponent(&component, node.ID, health)
		if len(sampleNodes) < sampleLimit {
			sampleNodes = append(sampleNodes, node)
		}
	}
	if component.NodeCount == 0 {
		return ComponentExplanation{}, nil, nil, nil, false
	}
	sort.Slice(sampleNodes, func(i, j int) bool {
		return sampleNodes[i].ID < sampleNodes[j].ID
	})
	countedSamples, excludedSamples := componentRelationshipSamples(g, componentNodeIDs, sampleLimit)
	component.CountedEdgeCount = countComponentCountedEdges(g, componentNodeIDs)
	finalizeComponent(&component)
	return component, sampleNodes, countedSamples, excludedSamples, true
}

func newComponentExplanation(componentID string) ComponentExplanation {
	return ComponentExplanation{
		ID:                             componentID,
		TopologyStatusCounts:           map[string]int{},
		ExpectedIsolationReasonCounts:  map[string]int{},
		ConfidenceCounts:               map[string]int{},
		DiagnosticCounts:               map[string]int{},
		DiagnosticClassificationCounts: map[string]int{},
		DiagnosticActionabilityCounts:  map[string]int{},
		ResolutionHealthBucketCounts:   map[string]int{},
		ResolutionConfidenceCounts:     map[string]int{},
	}
}

func addNodeToComponent(component *ComponentExplanation, nodeID string, health NodeHealth) {
	component.NodeCount++
	component.ReachableFromRoot = component.ReachableFromRoot || health.ComponentReachableFromRoot
	component.TopologyStatusCounts[string(health.TopologyStatus)]++
	component.ConfidenceCounts[health.Confidence]++
	component.ResolutionGapCount += health.ResolutionGapCount
	component.ResolutionConfidenceCounts[health.ResolutionConfidence]++
	for bucket, count := range health.ResolutionHealthBuckets {
		component.ResolutionHealthBucketCounts[bucket] += count
	}
	for _, reason := range health.ExpectedIsolationReasons {
		component.ExpectedIsolationReasonCounts[reason]++
		if reason == ReasonFrameworkEntry {
			component.RootNodeIDs = append(component.RootNodeIDs, nodeID)
		}
	}
	for _, diagnostic := range health.Diagnostics {
		if diagnostic.Kind == "" {
			continue
		}
		component.DiagnosticCounts[diagnostic.Kind] += diagnosticCount(diagnostic)
		if diagnostic.Classification != "" {
			component.DiagnosticClassificationCounts[diagnostic.Classification] += diagnosticCount(diagnostic)
		}
		if diagnostic.Actionability != "" {
			component.DiagnosticActionabilityCounts[diagnostic.Actionability] += diagnosticCount(diagnostic)
		}
	}
	if len(component.SampleNodeIDs) < 5 {
		component.SampleNodeIDs = append(component.SampleNodeIDs, nodeID)
	}
}

func finalizeComponent(component *ComponentExplanation) {
	sort.Strings(component.RootNodeIDs)
	sort.Strings(component.SampleNodeIDs)
	component.Detached = component.CountedEdgeCount > 0 && !component.ReachableFromRoot
}

func nodeHealthFromNode(node graph.Node) (NodeHealth, bool) {
	if node.Properties == nil {
		return NodeHealth{}, false
	}
	health, ok := node.Properties["graphHealth"].(NodeHealth)
	return health, ok
}

func nodeRelationships(g *graph.Graph, nodeID string) ([]graph.Relationship, []graph.Relationship, []graph.Relationship) {
	incoming := make([]graph.Relationship, 0)
	outgoing := make([]graph.Relationship, 0)
	excluded := make([]graph.Relationship, 0)
	for _, relationship := range g.Relationships {
		touchesNode := relationship.SourceID == nodeID || relationship.TargetID == nodeID
		if !touchesNode {
			continue
		}
		if !IsCounted(relationship.Type) {
			excluded = append(excluded, relationship)
			continue
		}
		if relationship.TargetID == nodeID {
			incoming = append(incoming, relationship)
		}
		if relationship.SourceID == nodeID {
			outgoing = append(outgoing, relationship)
		}
	}
	sortRelationships(incoming)
	sortRelationships(outgoing)
	sortRelationships(excluded)
	return incoming, outgoing, excluded
}

func componentRelationshipSamples(g *graph.Graph, componentNodeIDs map[string]bool, sampleLimit int) ([]graph.Relationship, []graph.Relationship) {
	counted := make([]graph.Relationship, 0)
	excluded := make([]graph.Relationship, 0)
	for _, relationship := range g.Relationships {
		sourceInComponent := componentNodeIDs[relationship.SourceID]
		targetInComponent := componentNodeIDs[relationship.TargetID]
		if !sourceInComponent && !targetInComponent {
			continue
		}
		if IsCounted(relationship.Type) && sourceInComponent && targetInComponent {
			counted = append(counted, relationship)
			continue
		}
		if !IsCounted(relationship.Type) {
			excluded = append(excluded, relationship)
		}
	}
	sortRelationships(counted)
	sortRelationships(excluded)
	return firstRelationships(counted, sampleLimit), firstRelationships(excluded, sampleLimit)
}

func countComponentCountedEdges(g *graph.Graph, componentNodeIDs map[string]bool) int {
	count := 0
	for _, relationship := range g.Relationships {
		if !IsCounted(relationship.Type) {
			continue
		}
		if componentNodeIDs[relationship.SourceID] && componentNodeIDs[relationship.TargetID] {
			count++
		}
	}
	return count
}

func sortRelationships(relationships []graph.Relationship) {
	sort.Slice(relationships, func(i, j int) bool {
		if relationships[i].Type != relationships[j].Type {
			return relationships[i].Type < relationships[j].Type
		}
		if relationships[i].SourceID != relationships[j].SourceID {
			return relationships[i].SourceID < relationships[j].SourceID
		}
		if relationships[i].TargetID != relationships[j].TargetID {
			return relationships[i].TargetID < relationships[j].TargetID
		}
		return relationships[i].ID < relationships[j].ID
	})
}

func firstRelationships(relationships []graph.Relationship, limit int) []graph.Relationship {
	if len(relationships) <= limit {
		return relationships
	}
	out := make([]graph.Relationship, limit)
	copy(out, relationships[:limit])
	return out
}
