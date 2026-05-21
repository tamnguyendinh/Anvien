package httpapi

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/tamnguyendinh/avmatrix-go/internal/graph"
	"github.com/tamnguyendinh/avmatrix-go/internal/graphhealth"
)

type graphResponse struct {
	Nodes         []graph.Node         `json:"nodes"`
	Relationships []graph.Relationship `json:"relationships"`
	GraphHealth   graphhealth.Summary  `json:"graphHealth"`
}

type graphHealthExplainResponse struct {
	Kind                         string                           `json:"kind"`
	NodeID                       string                           `json:"nodeId,omitempty"`
	ComponentID                  string                           `json:"componentId,omitempty"`
	Node                         *graph.Node                      `json:"node,omitempty"`
	Health                       *graphhealth.NodeHealth          `json:"health,omitempty"`
	Component                    *graphHealthComponentExplanation `json:"component,omitempty"`
	CountedIncomingRelationships []graph.Relationship             `json:"countedIncomingRelationships,omitempty"`
	CountedOutgoingRelationships []graph.Relationship             `json:"countedOutgoingRelationships,omitempty"`
	ExcludedRelationships        []graph.Relationship             `json:"excludedRelationships,omitempty"`
	SampleNodes                  []graph.Node                     `json:"sampleNodes,omitempty"`
	CountedRelationshipSamples   []graph.Relationship             `json:"countedRelationshipSamples,omitempty"`
	ExcludedRelationshipSamples  []graph.Relationship             `json:"excludedRelationshipSamples,omitempty"`
	SampleLimit                  int                              `json:"sampleLimit,omitempty"`
}

type graphHealthComponentExplanation struct {
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
}

type graphHealthReportResponse struct {
	ReportType         string                       `json:"reportType"`
	VerdictPolicy      string                       `json:"verdictPolicy"`
	Limit              int                          `json:"limit"`
	IncludeExpected    bool                         `json:"includeExpected"`
	Summary            graphhealth.Summary          `json:"summary"`
	TotalCandidates    int                          `json:"totalCandidates"`
	ReturnedCandidates int                          `json:"returnedCandidates"`
	Candidates         []graphHealthReportCandidate `json:"candidates"`
}

type graphHealthReportCandidate struct {
	NodeID                     string                     `json:"nodeId"`
	Label                      string                     `json:"label"`
	Name                       string                     `json:"name,omitempty"`
	FilePath                   string                     `json:"filePath,omitempty"`
	TriagePriority             string                     `json:"triagePriority"`
	TriageDimension            string                     `json:"triageDimension"`
	TopologyStatus             graphhealth.TopologyStatus `json:"topologyStatus"`
	Confidence                 string                     `json:"confidence"`
	CountedIncoming            int                        `json:"countedIncoming"`
	CountedOutgoing            int                        `json:"countedOutgoing"`
	ExcludedEdgeCounts         map[string]int             `json:"excludedEdgeCounts,omitempty"`
	ExpectedIsolationReasons   []string                   `json:"expectedIsolationReasons,omitempty"`
	Diagnostics                []graphhealth.Diagnostic   `json:"diagnostics,omitempty"`
	ComponentID                string                     `json:"componentId,omitempty"`
	ComponentSize              int                        `json:"componentSize,omitempty"`
	ComponentReachableFromRoot bool                       `json:"componentReachableFromRoot"`
}

const graphNDJSONFlushInterval = 512
const graphHealthExplainSampleLimit = 20
const graphHealthReportDefaultLimit = 100
const graphHealthReportMaxLimit = 1000
const graphHealthTriageDimensionTopology = "topology"
const graphHealthTriageDimensionDiagnostic = "diagnostic"

func (s Server) handleGraph(w http.ResponseWriter, r *http.Request) {
	if !methodAllowed(w, r, http.MethodGet) {
		return
	}

	entry, status, message, err := s.resolveRequestedRepo(r)
	if err != nil {
		if status == http.StatusNotFound {
			message = "Repository not found"
		}
		writeError(w, status, message)
		return
	}

	g, err := loadGraphSnapshot(filepath.Join(storagePathFor(entry), "graph.json"))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			writeError(w, http.StatusNotFound, "Graph not found. Run: avmatrix analyze")
			return
		}
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	includeContent := r.URL.Query().Get("includeContent") == "true"

	if r.URL.Query().Get("stream") == "true" {
		header := w.Header()
		header.Set("Content-Type", "application/x-ndjson; charset=utf-8")
		header.Set("Cache-Control", "no-cache")
		header.Set("X-Accel-Buffering", "no")
		w.WriteHeader(http.StatusOK)
		streamGraphNDJSON(w, g, includeContent)
		return
	}

	writeJSON(w, http.StatusOK, graphPayload(g, includeContent))
}

func (s Server) handleGraphHealthExplain(w http.ResponseWriter, r *http.Request) {
	if !methodAllowed(w, r, http.MethodGet) {
		return
	}

	nodeID := strings.TrimSpace(r.URL.Query().Get("nodeId"))
	componentID := strings.TrimSpace(r.URL.Query().Get("componentId"))
	if (nodeID == "" && componentID == "") || (nodeID != "" && componentID != "") {
		writeError(w, http.StatusBadRequest, `Provide exactly one of "nodeId" or "componentId"`)
		return
	}

	entry, status, message, err := s.resolveRequestedRepo(r)
	if err != nil {
		if status == http.StatusNotFound {
			message = "Repository not found"
		}
		writeError(w, status, message)
		return
	}

	g, err := loadGraphSnapshot(filepath.Join(storagePathFor(entry), "graph.json"))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			writeError(w, http.StatusNotFound, "Graph not found. Run: avmatrix analyze")
			return
		}
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	includeContent := r.URL.Query().Get("includeContent") == "true"
	graphhealth.Compute(g)

	if nodeID != "" {
		response, ok := graphHealthNodeExplain(g, nodeID, includeContent)
		if !ok {
			writeError(w, http.StatusNotFound, "Graph node not found")
			return
		}
		writeJSON(w, http.StatusOK, response)
		return
	}

	response, ok := graphHealthComponentExplain(g, componentID, includeContent)
	if !ok {
		writeError(w, http.StatusNotFound, "Graph component not found")
		return
	}
	writeJSON(w, http.StatusOK, response)
}

func (s Server) handleGraphHealthReport(w http.ResponseWriter, r *http.Request) {
	if !methodAllowed(w, r, http.MethodGet) {
		return
	}

	entry, status, message, err := s.resolveRequestedRepo(r)
	if err != nil {
		if status == http.StatusNotFound {
			message = "Repository not found"
		}
		writeError(w, status, message)
		return
	}

	g, err := loadGraphSnapshot(filepath.Join(storagePathFor(entry), "graph.json"))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			writeError(w, http.StatusNotFound, "Graph not found. Run: avmatrix analyze")
			return
		}
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	limit := graphHealthReportLimit(r)
	includeExpected := r.URL.Query().Get("includeExpected") == "true"
	summary := graphhealth.ComputeSummary(g)
	candidates := graphHealthReportCandidates(g, includeExpected)
	totalCandidates := len(candidates)
	if len(candidates) > limit {
		candidates = candidates[:limit]
	}
	writeJSON(w, http.StatusOK, graphHealthReportResponse{
		ReportType:         "graph_health_candidate_review",
		VerdictPolicy:      "candidate_not_confirmed",
		Limit:              limit,
		IncludeExpected:    includeExpected,
		Summary:            summary,
		TotalCandidates:    totalCandidates,
		ReturnedCandidates: len(candidates),
		Candidates:         candidates,
	})
}

func graphPayload(g *graph.Graph, includeContent bool) graphResponse {
	summary := graphhealth.ComputeSummary(g)
	nodes := make([]graph.Node, 0, len(g.Nodes))
	for _, node := range g.Nodes {
		nodes = append(nodes, graphNodeForResponse(node, includeContent))
	}
	return graphResponse{Nodes: nodes, Relationships: g.Relationships, GraphHealth: summary}
}

func graphHealthNodeExplain(g *graph.Graph, nodeID string, includeContent bool) (graphHealthExplainResponse, bool) {
	node, ok := g.GetNode(nodeID)
	if !ok {
		return graphHealthExplainResponse{}, false
	}
	health, ok := nodeHealthFromNode(node)
	if !ok {
		return graphHealthExplainResponse{}, false
	}
	incoming, outgoing, excluded := graphHealthNodeRelationships(g, nodeID)
	publicNode := graphNodeForResponse(node, includeContent)
	return graphHealthExplainResponse{
		Kind:                         "node",
		NodeID:                       nodeID,
		ComponentID:                  health.ComponentID,
		Node:                         &publicNode,
		Health:                       &health,
		CountedIncomingRelationships: incoming,
		CountedOutgoingRelationships: outgoing,
		ExcludedRelationships:        excluded,
	}, true
}

func graphHealthComponentExplain(g *graph.Graph, componentID string, includeContent bool) (graphHealthExplainResponse, bool) {
	component := graphHealthComponentExplanation{
		ID:                             componentID,
		TopologyStatusCounts:           map[string]int{},
		ExpectedIsolationReasonCounts:  map[string]int{},
		ConfidenceCounts:               map[string]int{},
		DiagnosticCounts:               map[string]int{},
		DiagnosticClassificationCounts: map[string]int{},
		DiagnosticActionabilityCounts:  map[string]int{},
	}
	componentNodeIDs := map[string]bool{}
	sampleNodes := make([]graph.Node, 0, graphHealthExplainSampleLimit)
	for _, node := range g.Nodes {
		health, ok := nodeHealthFromNode(node)
		if !ok || health.ComponentID != componentID {
			continue
		}
		componentNodeIDs[node.ID] = true
		component.NodeCount++
		component.ReachableFromRoot = component.ReachableFromRoot || health.ComponentReachableFromRoot
		component.TopologyStatusCounts[string(health.TopologyStatus)]++
		component.ConfidenceCounts[health.Confidence]++
		for _, reason := range health.ExpectedIsolationReasons {
			component.ExpectedIsolationReasonCounts[reason]++
			if reason == graphhealth.ReasonFrameworkEntry {
				component.RootNodeIDs = append(component.RootNodeIDs, node.ID)
			}
		}
		for _, diagnostic := range health.Diagnostics {
			if diagnostic.Kind == "" {
				continue
			}
			component.DiagnosticCounts[diagnostic.Kind] += graphHealthDiagnosticCount(diagnostic)
			if diagnostic.Classification != "" {
				component.DiagnosticClassificationCounts[diagnostic.Classification] += graphHealthDiagnosticCount(diagnostic)
			}
			if diagnostic.Actionability != "" {
				component.DiagnosticActionabilityCounts[diagnostic.Actionability] += graphHealthDiagnosticCount(diagnostic)
			}
		}
		if len(sampleNodes) < graphHealthExplainSampleLimit {
			publicNode := graphNodeForResponse(node, includeContent)
			sampleNodes = append(sampleNodes, publicNode)
			component.SampleNodeIDs = append(component.SampleNodeIDs, node.ID)
		}
	}
	if component.NodeCount == 0 {
		return graphHealthExplainResponse{}, false
	}
	sort.Strings(component.RootNodeIDs)
	sort.Strings(component.SampleNodeIDs)
	sort.Slice(sampleNodes, func(i, j int) bool {
		return sampleNodes[i].ID < sampleNodes[j].ID
	})
	countedSamples, excludedSamples := graphHealthComponentRelationshipSamples(g, componentNodeIDs)
	component.CountedEdgeCount = countComponentCountedEdges(g, componentNodeIDs)
	component.Detached = component.CountedEdgeCount > 0 && !component.ReachableFromRoot
	return graphHealthExplainResponse{
		Kind:                        "component",
		ComponentID:                 componentID,
		Component:                   &component,
		SampleNodes:                 sampleNodes,
		CountedRelationshipSamples:  countedSamples,
		ExcludedRelationshipSamples: excludedSamples,
		SampleLimit:                 graphHealthExplainSampleLimit,
	}, true
}

func graphHealthReportCandidates(g *graph.Graph, includeExpected bool) []graphHealthReportCandidate {
	candidates := make([]graphHealthReportCandidate, 0)
	for _, node := range g.Nodes {
		health, ok := nodeHealthFromNode(node)
		if !ok {
			continue
		}
		priority, dimension, rank := graphHealthReportPriority(health)
		if rank == 0 {
			continue
		}
		if !includeExpected && health.Confidence == graphhealth.ConfidenceExpected {
			continue
		}
		candidates = append(candidates, graphHealthReportCandidate{
			NodeID:                     node.ID,
			Label:                      string(node.Label),
			Name:                       graphNodeStringProperty(node, "name"),
			FilePath:                   graphNodeStringProperty(node, "filePath"),
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
		})
	}
	sort.SliceStable(candidates, func(i, j int) bool {
		leftRank := graphHealthReportPriorityRank(candidates[i].TriagePriority, candidates[i].TriageDimension)
		rightRank := graphHealthReportPriorityRank(candidates[j].TriagePriority, candidates[j].TriageDimension)
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

func graphHealthReportPriority(health graphhealth.NodeHealth) (string, string, int) {
	switch health.TopologyStatus {
	case graphhealth.TopologyNoIncoming:
		return string(graphhealth.TopologyNoIncoming), graphHealthTriageDimensionTopology, 1
	case graphhealth.TopologyDetached:
		return string(graphhealth.TopologyDetached), graphHealthTriageDimensionTopology, 2
	case graphhealth.TopologyTrueIsolated:
		return string(graphhealth.TopologyTrueIsolated), graphHealthTriageDimensionTopology, 4
	case graphhealth.TopologyNoOutgoing:
		return string(graphhealth.TopologyNoOutgoing), graphHealthTriageDimensionTopology, 5
	case graphhealth.TopologyUnknown:
		return string(graphhealth.TopologyUnknown), graphHealthTriageDimensionTopology, 6
	}
	if hasDiagnosticKind(health.Diagnostics, graphhealth.DiagnosticUnresolvedReference) {
		return graphhealth.DiagnosticUnresolvedReference, graphHealthTriageDimensionDiagnostic, 30
	}
	return "", "", 0
}

func graphHealthReportPriorityRank(priority string, dimension string) int {
	if dimension == graphHealthTriageDimensionDiagnostic {
		switch priority {
		case graphhealth.DiagnosticUnresolvedReference:
			return 30
		default:
			return 90
		}
	}
	switch priority {
	case string(graphhealth.TopologyNoIncoming):
		return 1
	case string(graphhealth.TopologyDetached):
		return 2
	case string(graphhealth.TopologyTrueIsolated):
		return 4
	case string(graphhealth.TopologyNoOutgoing):
		return 5
	case string(graphhealth.TopologyUnknown):
		return 6
	default:
		return 99
	}
}

func streamGraphNDJSON(w http.ResponseWriter, g *graph.Graph, includeContent bool) {
	graphhealth.Compute(g)
	encoder := json.NewEncoder(w)
	flusher, _ := w.(http.Flusher)
	written := 0
	flush := func(force bool) {
		if flusher == nil {
			return
		}
		if force || written%graphNDJSONFlushInterval == 0 {
			flusher.Flush()
		}
	}
	for _, node := range g.Nodes {
		_ = encoder.Encode(map[string]any{
			"type": "node",
			"data": graphNodeForResponse(node, includeContent),
		})
		written++
		flush(false)
	}
	for _, relationship := range g.Relationships {
		_ = encoder.Encode(map[string]any{
			"type": "relationship",
			"data": relationship,
		})
		written++
		flush(false)
	}
	flush(true)
}

func graphNodeForResponse(node graph.Node, includeContent bool) graph.Node {
	if len(node.Properties) == 0 {
		return node
	}
	stripKeys := map[string]bool{
		graphhealth.DiagnosticPropertyKey: true,
	}
	if !includeContent {
		stripKeys["content"] = true
	}
	needsStrip := false
	for key := range stripKeys {
		if _, ok := node.Properties[key]; ok {
			needsStrip = true
			break
		}
	}
	if !needsStrip {
		return node
	}
	properties := make(graph.NodeProperties, len(node.Properties))
	for key, value := range node.Properties {
		if stripKeys[key] {
			continue
		}
		properties[key] = value
	}
	node.Properties = properties
	return node
}

func graphHealthReportLimit(r *http.Request) int {
	raw := strings.TrimSpace(r.URL.Query().Get("limit"))
	if raw == "" {
		return graphHealthReportDefaultLimit
	}
	limit, err := strconv.Atoi(raw)
	if err != nil || limit < 1 {
		return graphHealthReportDefaultLimit
	}
	if limit > graphHealthReportMaxLimit {
		return graphHealthReportMaxLimit
	}
	return limit
}

func graphNodeStringProperty(node graph.Node, key string) string {
	if node.Properties == nil {
		return ""
	}
	value, ok := node.Properties[key]
	if !ok {
		return ""
	}
	text, _ := value.(string)
	return text
}

func hasDiagnosticKind(diagnostics []graphhealth.Diagnostic, kind string) bool {
	for _, diagnostic := range diagnostics {
		if diagnostic.Kind == kind {
			return true
		}
	}
	return false
}

func nodeHealthFromNode(node graph.Node) (graphhealth.NodeHealth, bool) {
	if node.Properties == nil {
		return graphhealth.NodeHealth{}, false
	}
	health, ok := node.Properties["graphHealth"].(graphhealth.NodeHealth)
	return health, ok
}

func graphHealthNodeRelationships(g *graph.Graph, nodeID string) ([]graph.Relationship, []graph.Relationship, []graph.Relationship) {
	incoming := make([]graph.Relationship, 0)
	outgoing := make([]graph.Relationship, 0)
	excluded := make([]graph.Relationship, 0)
	for _, relationship := range g.Relationships {
		touchesNode := relationship.SourceID == nodeID || relationship.TargetID == nodeID
		if !touchesNode {
			continue
		}
		if !graphhealth.IsCounted(relationship.Type) {
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

func graphHealthComponentRelationshipSamples(g *graph.Graph, componentNodeIDs map[string]bool) ([]graph.Relationship, []graph.Relationship) {
	counted := make([]graph.Relationship, 0)
	excluded := make([]graph.Relationship, 0)
	for _, relationship := range g.Relationships {
		sourceInComponent := componentNodeIDs[relationship.SourceID]
		targetInComponent := componentNodeIDs[relationship.TargetID]
		if !sourceInComponent && !targetInComponent {
			continue
		}
		if graphhealth.IsCounted(relationship.Type) && sourceInComponent && targetInComponent {
			counted = append(counted, relationship)
			continue
		}
		if !graphhealth.IsCounted(relationship.Type) {
			excluded = append(excluded, relationship)
		}
	}
	sortRelationships(counted)
	sortRelationships(excluded)
	return firstRelationships(counted, graphHealthExplainSampleLimit), firstRelationships(excluded, graphHealthExplainSampleLimit)
}

func countComponentCountedEdges(g *graph.Graph, componentNodeIDs map[string]bool) int {
	count := 0
	for _, relationship := range g.Relationships {
		if !graphhealth.IsCounted(relationship.Type) {
			continue
		}
		if componentNodeIDs[relationship.SourceID] && componentNodeIDs[relationship.TargetID] {
			count++
		}
	}
	return count
}

func graphHealthDiagnosticCount(diagnostic graphhealth.Diagnostic) int {
	if diagnostic.Count > 0 {
		return diagnostic.Count
	}
	return 1
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
