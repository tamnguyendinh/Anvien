package httpapi

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/graphhealth"
	"github.com/tamnguyendinh/anvien/internal/semantic"
)

type graphResponse struct {
	Nodes          []graph.Node         `json:"nodes"`
	Relationships  []graph.Relationship `json:"relationships"`
	GraphHealth    graphhealth.Summary  `json:"graphHealth"`
	SemanticStatus semantic.GraphStatus `json:"semanticStatus"`
}

type graphHealthExplainResponse = graphhealth.ExplainResponse
type graphHealthComponentExplanation = graphhealth.ComponentExplanation
type graphHealthReportResponse = graphhealth.ReportResponse
type graphHealthReportCandidate = graphhealth.ReportCandidate

const graphNDJSONFlushInterval = 512
const graphHealthExplainSampleLimit = graphhealth.ExplainSampleLimit
const graphHealthReportDefaultLimit = graphhealth.ReportDefaultLimit
const graphHealthReportMaxLimit = graphhealth.ReportMaxLimit
const graphHealthTriageDimensionTopology = graphhealth.TriageDimensionTopology
const graphHealthTriageDimensionDiagnostic = graphhealth.TriageDimensionDiagnostic

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
			writeError(w, http.StatusNotFound, "Graph not found. Run: anvien analyze")
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
			writeError(w, http.StatusNotFound, "Graph not found. Run: anvien analyze")
			return
		}
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	includeContent := r.URL.Query().Get("includeContent") == "true"

	if nodeID != "" {
		response, ok := graphhealth.ExplainNode(g, nodeID)
		if !ok {
			writeError(w, http.StatusNotFound, "Graph node not found")
			return
		}
		writeJSON(w, http.StatusOK, graphHealthExplainResponseForResponse(response, includeContent))
		return
	}

	response, ok := graphhealth.ExplainComponent(g, componentID)
	if !ok {
		writeError(w, http.StatusNotFound, "Graph component not found")
		return
	}
	writeJSON(w, http.StatusOK, graphHealthExplainResponseForResponse(response, includeContent))
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
			writeError(w, http.StatusNotFound, "Graph not found. Run: anvien analyze")
			return
		}
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	limit := graphHealthReportLimit(r)
	includeExpected := r.URL.Query().Get("includeExpected") == "true"
	writeJSON(w, http.StatusOK, graphhealth.BuildReport(g, graphhealth.ReportOptions{
		Limit:           limit,
		IncludeExpected: includeExpected,
	}))
}

func graphPayload(g *graph.Graph, includeContent bool) graphResponse {
	summary := graphhealth.ComputeSummary(g)
	semanticStatus := semantic.GraphSemanticStatus(g)
	nodes := make([]graph.Node, 0, len(g.Nodes))
	for _, node := range g.Nodes {
		nodes = append(nodes, graphNodeForResponse(node, includeContent))
	}
	return graphResponse{Nodes: nodes, Relationships: g.Relationships, GraphHealth: summary, SemanticStatus: semanticStatus}
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
	_ = encoder.Encode(map[string]any{
		"type": "semantic_status",
		"data": semantic.GraphSemanticStatus(g),
	})
	written++
	flush(false)
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

func graphHealthExplainResponseForResponse(response graphHealthExplainResponse, includeContent bool) graphHealthExplainResponse {
	if response.Node != nil {
		node := graphNodeForResponse(*response.Node, includeContent)
		response.Node = &node
	}
	if len(response.SampleNodes) > 0 {
		nodes := make([]graph.Node, 0, len(response.SampleNodes))
		for _, node := range response.SampleNodes {
			nodes = append(nodes, graphNodeForResponse(node, includeContent))
		}
		response.SampleNodes = nodes
	}
	return response
}

func graphHealthReportCandidates(g *graph.Graph, includeExpected bool) []graphHealthReportCandidate {
	return graphhealth.ReportCandidates(g, includeExpected)
}
