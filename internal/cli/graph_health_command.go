package cli

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tamnguyendinh/avmatrix-go/internal/graph"
	"github.com/tamnguyendinh/avmatrix-go/internal/graphhealth"
	"github.com/tamnguyendinh/avmatrix-go/internal/repo"
)

type graphHealthCommandInputs struct {
	Repo          string `json:"repo"`
	RepoPath      string `json:"repoPath"`
	Graph         string `json:"graph"`
	IndexedCommit string `json:"indexedCommit,omitempty"`
	CurrentCommit string `json:"currentCommit,omitempty"`
}

type graphHealthCommandTotals struct {
	Nodes         int `json:"nodes"`
	Relationships int `json:"relationships"`
}

type graphHealthSummaryResult struct {
	Inputs  graphHealthCommandInputs `json:"inputs"`
	Totals  graphHealthCommandTotals `json:"totals"`
	Summary graphhealth.Summary      `json:"summary"`
}

type graphHealthReportResult struct {
	Inputs graphHealthCommandInputs `json:"inputs"`
	Totals graphHealthCommandTotals `json:"totals"`
	graphhealth.ReportResponse
}

type graphHealthComponentsResult struct {
	Inputs             graphHealthCommandInputs           `json:"inputs"`
	Totals             graphHealthCommandTotals           `json:"totals"`
	Limit              int                                `json:"limit"`
	Summary            graphhealth.Summary                `json:"summary"`
	TotalComponents    int                                `json:"totalComponents"`
	ReturnedComponents int                                `json:"returnedComponents"`
	Components         []graphhealth.ComponentExplanation `json:"components"`
}

type graphHealthExplainResult struct {
	Inputs graphHealthCommandInputs `json:"inputs"`
	Totals graphHealthCommandTotals `json:"totals"`
	graphhealth.ExplainResponse
}

func newGraphHealthCommand() *cobra.Command {
	var repoName string
	cmd := &cobra.Command{
		Use:   "graph-health",
		Short: "Audit graph topology health, diagnostics, and components",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
	cmd.PersistentFlags().StringVar(&repoName, "repo", "", "indexed repository name or absolute path")
	cmd.AddCommand(
		newGraphHealthSummaryCommand(&repoName),
		newGraphHealthReportCommand(&repoName),
		newGraphHealthComponentsCommand(&repoName),
		newGraphHealthExplainCommand(&repoName),
	)
	return cmd
}

func newGraphHealthSummaryCommand(repoName *string) *cobra.Command {
	var jsonOutput bool
	cmd := &cobra.Command{
		Use:   "summary",
		Short: "Print graph-health topology and diagnostic summary",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			inputs, g, err := loadGraphHealthGraph(*repoName)
			if err != nil {
				return err
			}
			result := graphHealthSummaryResult{
				Inputs:  inputs,
				Totals:  graphHealthTotals(g),
				Summary: graphhealth.ComputeSummary(g),
			}
			if jsonOutput {
				return writeGraphHealthJSON(cmd, result)
			}
			for _, line := range graphHealthSummaryLines(result) {
				if _, err := fmt.Fprintln(cmd.OutOrStdout(), line); err != nil {
					return err
				}
			}
			return nil
		},
	}
	cmd.Flags().BoolVar(&jsonOutput, "json", false, "write JSON graph-health summary")
	return cmd
}

func newGraphHealthReportCommand(repoName *string) *cobra.Command {
	var limit int
	var includeExpected bool
	var jsonOutput bool
	cmd := &cobra.Command{
		Use:   "report",
		Short: "Print graph-health triage candidates",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			inputs, g, err := loadGraphHealthGraph(*repoName)
			if err != nil {
				return err
			}
			result := graphHealthReportResult{
				Inputs: inputs,
				Totals: graphHealthTotals(g),
				ReportResponse: graphhealth.BuildReport(g, graphhealth.ReportOptions{
					Limit:           limit,
					IncludeExpected: includeExpected,
				}),
			}
			if jsonOutput {
				return writeGraphHealthJSON(cmd, result)
			}
			for _, line := range graphHealthReportLines(result.ReportResponse) {
				if _, err := fmt.Fprintln(cmd.OutOrStdout(), line); err != nil {
					return err
				}
			}
			return nil
		},
	}
	cmd.Flags().IntVar(&limit, "limit", graphhealth.ReportDefaultLimit, "maximum candidate rows")
	cmd.Flags().BoolVar(&includeExpected, "include-expected", false, "include expected-isolated nodes")
	cmd.Flags().BoolVar(&jsonOutput, "json", false, "write JSON graph-health report")
	return cmd
}

func newGraphHealthComponentsCommand(repoName *string) *cobra.Command {
	var limit int
	var jsonOutput bool
	cmd := &cobra.Command{
		Use:   "components",
		Short: "Print graph-health component summaries",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			inputs, g, err := loadGraphHealthGraph(*repoName)
			if err != nil {
				return err
			}
			summary := graphhealth.ComputeSummary(g)
			components := graphhealth.ComponentSummaries(g)
			limit = graphhealth.NormalizeReportLimit(limit)
			total := len(components)
			if len(components) > limit {
				components = components[:limit]
			}
			result := graphHealthComponentsResult{
				Inputs:             inputs,
				Totals:             graphHealthTotals(g),
				Limit:              limit,
				Summary:            summary,
				TotalComponents:    total,
				ReturnedComponents: len(components),
				Components:         components,
			}
			if jsonOutput {
				return writeGraphHealthJSON(cmd, result)
			}
			for _, line := range graphHealthComponentLines(result) {
				if _, err := fmt.Fprintln(cmd.OutOrStdout(), line); err != nil {
					return err
				}
			}
			return nil
		},
	}
	cmd.Flags().IntVar(&limit, "limit", graphhealth.ReportDefaultLimit, "maximum component rows")
	cmd.Flags().BoolVar(&jsonOutput, "json", false, "write JSON graph-health components")
	return cmd
}

func newGraphHealthExplainCommand(repoName *string) *cobra.Command {
	var componentID string
	var jsonOutput bool
	cmd := &cobra.Command{
		Use:   "explain [node-id-or-name]",
		Short: "Explain one graph node or component",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			inputs, g, err := loadGraphHealthGraph(*repoName)
			if err != nil {
				return err
			}
			var response graphhealth.ExplainResponse
			var ok bool
			componentID = strings.TrimSpace(componentID)
			switch {
			case componentID != "" && len(args) > 0:
				return fmt.Errorf("provide either a node selector or --component, not both")
			case componentID != "":
				response, ok = graphhealth.ExplainComponent(g, componentID)
				if !ok {
					return fmt.Errorf("graph component not found: %s", componentID)
				}
			case len(args) == 1:
				nodeID, err := resolveGraphHealthNodeSelector(g, args[0])
				if err != nil {
					return err
				}
				response, ok = graphhealth.ExplainNode(g, nodeID)
				if !ok {
					return fmt.Errorf("graph node not found: %s", args[0])
				}
			default:
				return fmt.Errorf("graph-health explain requires a node selector or --component")
			}
			result := graphHealthExplainResult{
				Inputs:          inputs,
				Totals:          graphHealthTotals(g),
				ExplainResponse: response,
			}
			if jsonOutput {
				return writeGraphHealthJSON(cmd, result)
			}
			for _, line := range graphHealthExplainLines(response) {
				if _, err := fmt.Fprintln(cmd.OutOrStdout(), line); err != nil {
					return err
				}
			}
			return nil
		},
	}
	cmd.Flags().StringVar(&componentID, "component", "", "component id to explain")
	cmd.Flags().BoolVar(&jsonOutput, "json", false, "write JSON graph-health explanation")
	return cmd
}

func loadGraphHealthGraph(repoName string) (graphHealthCommandInputs, *graph.Graph, error) {
	store := repo.NewEnvStore()
	entries, err := store.ListRegistered(false)
	if err != nil {
		return graphHealthCommandInputs{}, nil, err
	}
	var entry repo.RegistryEntry
	if strings.TrimSpace(repoName) == "" {
		if len(entries) != 1 {
			return graphHealthCommandInputs{}, nil, fmt.Errorf("Repository not found. Run: avmatrix analyze --force")
		}
		entry = entries[0]
	} else {
		resolved, err := repo.ResolveEntry(entries, repoName)
		if err != nil {
			return graphHealthCommandInputs{}, nil, err
		}
		entry = resolved
	}
	currentCommit := repo.CurrentCommit(entry.Path)
	if currentCommit != "" && entry.LastCommit != "" && currentCommit != entry.LastCommit {
		return graphHealthCommandInputs{}, nil, fmt.Errorf("graph-health requires fresh analyze output for %s: indexed commit %s current commit %s; run avmatrix analyze --force", entry.Name, shortCommit(entry.LastCommit), shortCommit(currentCommit))
	}
	storagePath := entry.StoragePath
	if storagePath == "" {
		storagePath = repo.StoragePath(entry.Path)
	}
	graphPath := filepath.Join(storagePath, "graph.json")
	raw, err := os.ReadFile(graphPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return graphHealthCommandInputs{}, nil, fmt.Errorf("graph-health requires fresh analyze output for %s: graph not found at %s; run avmatrix analyze --force", entry.Name, graphPath)
		}
		return graphHealthCommandInputs{}, nil, fmt.Errorf("read graph %s: %w", graphPath, err)
	}
	var g graph.Graph
	if err := json.Unmarshal(raw, &g); err != nil {
		return graphHealthCommandInputs{}, nil, fmt.Errorf("decode graph %s: %w", graphPath, err)
	}
	return graphHealthCommandInputs{
		Repo:          entry.Name,
		RepoPath:      entry.Path,
		Graph:         graphPath,
		IndexedCommit: entry.LastCommit,
		CurrentCommit: currentCommit,
	}, &g, nil
}

func graphHealthTotals(g *graph.Graph) graphHealthCommandTotals {
	if g == nil {
		return graphHealthCommandTotals{}
	}
	return graphHealthCommandTotals{Nodes: len(g.Nodes), Relationships: len(g.Relationships)}
}

func resolveGraphHealthNodeSelector(g *graph.Graph, selector string) (string, error) {
	selector = strings.TrimSpace(selector)
	if selector == "" {
		return "", fmt.Errorf("graph node selector is required")
	}
	if _, ok := g.GetNode(selector); ok {
		return selector, nil
	}
	matches := make([]string, 0)
	for _, node := range g.Nodes {
		if graphHealthNodePropertyString(node, "name") == selector {
			matches = append(matches, node.ID)
		}
	}
	sort.Strings(matches)
	switch len(matches) {
	case 0:
		return "", fmt.Errorf("graph node not found: %s", selector)
	case 1:
		return matches[0], nil
	default:
		return "", fmt.Errorf("graph node name %q is ambiguous across %d nodes: %s", selector, len(matches), strings.Join(matches, ", "))
	}
}

func graphHealthSummaryLines(result graphHealthSummaryResult) []string {
	summary := result.Summary
	return []string{
		fmt.Sprintf("graphHealth.repo=%q graph=%q nodes=%d relationships=%d countedRelationships=%d components=%d detachedComponents=%d rootNodes=%d",
			result.Inputs.Repo,
			result.Inputs.Graph,
			result.Totals.Nodes,
			result.Totals.Relationships,
			summary.CountedRelationshipCount,
			summary.ComponentCount,
			summary.DetachedComponentCount,
			summary.RootNodeCount,
		),
		"graphHealth.topology=" + formatCountMap(summary.TopologyStatusCounts),
		"graphHealth.confidence=" + formatCountMap(summary.ConfidenceCounts),
		"graphHealth.diagnostics=" + formatCountMap(summary.DiagnosticCounts),
		"graphHealth.resolutionConfidence=" + formatCountMap(summary.ResolutionConfidenceCounts),
		fmt.Sprintf("graphHealth.resolutionGaps nodes=%d relationships=%d occurrences=%d sourceBackedUnresolved=%d unattributedUnresolved=%d",
			summary.ResolutionGapNodeCount,
			summary.HasResolutionGapRelationshipCount,
			summary.ResolutionGapCount,
			summary.SourceBackedUnresolvedReferenceCount,
			summary.UnattributedUnresolvedReferenceCount,
		),
	}
}

func graphHealthReportLines(report graphhealth.ReportResponse) []string {
	lines := []string{
		fmt.Sprintf("graphHealth.report totalCandidates=%d returnedCandidates=%d limit=%d includeExpected=%t",
			report.TotalCandidates,
			report.ReturnedCandidates,
			report.Limit,
			report.IncludeExpected,
		),
	}
	for _, candidate := range report.Candidates {
		lines = append(lines, fmt.Sprintf("graphHealth.candidate nodeId=%q priority=%s dimension=%s topology=%s confidence=%s incoming=%d outgoing=%d component=%q file=%q name=%q",
			candidate.NodeID,
			candidate.TriagePriority,
			candidate.TriageDimension,
			candidate.TopologyStatus,
			candidate.Confidence,
			candidate.CountedIncoming,
			candidate.CountedOutgoing,
			candidate.ComponentID,
			candidate.FilePath,
			candidate.Name,
		))
	}
	return lines
}

func graphHealthComponentLines(result graphHealthComponentsResult) []string {
	lines := []string{
		fmt.Sprintf("graphHealth.components totalComponents=%d returnedComponents=%d limit=%d",
			result.TotalComponents,
			result.ReturnedComponents,
			result.Limit,
		),
	}
	for _, component := range result.Components {
		lines = append(lines, fmt.Sprintf("graphHealth.component componentId=%q nodes=%d countedEdges=%d detached=%t reachableFromRoot=%t roots=%q sample=%q",
			component.ID,
			component.NodeCount,
			component.CountedEdgeCount,
			component.Detached,
			component.ReachableFromRoot,
			strings.Join(component.RootNodeIDs, ","),
			strings.Join(component.SampleNodeIDs, ","),
		))
	}
	return lines
}

func graphHealthExplainLines(response graphhealth.ExplainResponse) []string {
	if response.Kind == "component" && response.Component != nil {
		component := response.Component
		return []string{
			fmt.Sprintf("graphHealth.component componentId=%q nodes=%d countedEdges=%d detached=%t reachableFromRoot=%t sampleLimit=%d",
				component.ID,
				component.NodeCount,
				component.CountedEdgeCount,
				component.Detached,
				component.ReachableFromRoot,
				response.SampleLimit,
			),
			"graphHealth.componentTopology=" + formatCountMap(component.TopologyStatusCounts),
			"graphHealth.componentConfidence=" + formatCountMap(component.ConfidenceCounts),
		}
	}
	if response.Health == nil {
		return []string{"graphHealth.explain kind=unknown"}
	}
	health := response.Health
	excludedCount := len(response.ExcludedRelationships)
	return []string{
		fmt.Sprintf("graphHealth.node nodeId=%q topology=%s confidence=%s incoming=%d outgoing=%d excluded=%d component=%q resolutionConfidence=%s resolutionGaps=%d",
			response.NodeID,
			health.TopologyStatus,
			health.Confidence,
			health.CountedIncoming,
			health.CountedOutgoing,
			excludedCount,
			health.ComponentID,
			health.ResolutionConfidence,
			health.ResolutionGapCount,
		),
		"graphHealth.nodeDiagnostics=" + graphHealthDiagnosticKinds(health.Diagnostics),
		"graphHealth.nodeExcludedEdges=" + formatCountMap(health.ExcludedEdgeCounts),
	}
}

func graphHealthDiagnosticKinds(diagnostics []graphhealth.Diagnostic) string {
	if len(diagnostics) == 0 {
		return ""
	}
	counts := map[string]int{}
	for _, diagnostic := range diagnostics {
		if diagnostic.Kind == "" {
			continue
		}
		count := diagnostic.Count
		if count <= 0 {
			count = 1
		}
		counts[diagnostic.Kind] += count
	}
	return formatCountMap(counts)
}

func graphHealthNodePropertyString(node graph.Node, key string) string {
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

func writeGraphHealthJSON(cmd *cobra.Command, value any) error {
	raw, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal graph-health output: %w", err)
	}
	_, err = fmt.Fprintf(cmd.OutOrStdout(), "%s\n", raw)
	return err
}
