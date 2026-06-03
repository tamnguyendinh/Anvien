package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/tamnguyendinh/anvien/internal/filecontext"
	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/graphhealth"
)

type resolutionInventoryResult struct {
	GeneratedAt string                         `json:"generatedAt"`
	Inputs      resolutionInventoryInputs      `json:"inputs"`
	Totals      resolutionInventoryTotals      `json:"totals"`
	GraphHealth graphhealth.Summary            `json:"graphHealth"`
	FileGroups  []resolutionInventoryFileGroup `json:"fileGroups"`
}

type resolutionInventoryInputs struct {
	Graph string `json:"graph"`
}

type resolutionInventoryTotals struct {
	Nodes         int `json:"nodes"`
	Relationships int `json:"relationships"`
}

type resolutionInventoryFileGroup struct {
	Path                 string                         `json:"path"`
	Total                int                            `json:"total"`
	ByKind               map[string]int                 `json:"byKind,omitempty"`
	ByClassification     map[string]int                 `json:"byClassification,omitempty"`
	ByActionability      map[string]int                 `json:"byActionability,omitempty"`
	NearestSourceSymbols []string                       `json:"nearestSourceSymbols,omitempty"`
	Samples              []filecontext.UnresolvedSample `json:"samples,omitempty"`
}

func newResolutionInventoryCommand() *cobra.Command {
	var graphPath string
	var outPath string
	var jsonOutput bool

	cmd := &cobra.Command{
		Use:   "resolution-inventory",
		Short: "Report full persisted ResolutionGap and Resolution Health inventory",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			result, err := runResolutionInventory(graphPath)
			if err != nil {
				return err
			}
			if outPath != "" {
				if err := writeResolutionInventoryResult(outPath, result); err != nil {
					return err
				}
			}
			if jsonOutput || outPath == "" {
				raw, err := json.MarshalIndent(result, "", "  ")
				if err != nil {
					return fmt.Errorf("marshal resolution inventory report: %w", err)
				}
				_, err = fmt.Fprintf(cmd.OutOrStdout(), "%s\n", raw)
				return err
			}
			if _, err := fmt.Fprintf(cmd.OutOrStdout(), "wrote %s\n", outPath); err != nil {
				return err
			}
			for _, line := range resolutionInventorySummaryLines(result) {
				if _, err := fmt.Fprintln(cmd.OutOrStdout(), line); err != nil {
					return err
				}
			}
			return nil
		},
	}
	cmd.Flags().StringVar(&graphPath, "graph", filepath.Join(".anvien", "graph.json"), "Anvien graph snapshot JSON")
	cmd.Flags().StringVar(&outPath, "out", "", "write inventory JSON to this path")
	cmd.Flags().BoolVar(&jsonOutput, "json", false, "write full JSON report to stdout")
	return cmd
}

func runResolutionInventory(graphPath string) (resolutionInventoryResult, error) {
	if strings.TrimSpace(graphPath) == "" {
		return resolutionInventoryResult{}, fmt.Errorf("graph path is required")
	}
	g, err := readResolutionInventoryGraph(graphPath)
	if err != nil {
		return resolutionInventoryResult{}, err
	}
	summary := graphhealth.ComputeSummary(&g)
	return resolutionInventoryResult{
		GeneratedAt: time.Now().Format(time.RFC3339),
		Inputs: resolutionInventoryInputs{
			Graph: graphPath,
		},
		Totals: resolutionInventoryTotals{
			Nodes:         len(g.Nodes),
			Relationships: len(g.Relationships),
		},
		GraphHealth: summary,
		FileGroups:  resolutionInventoryFileGroups(&g, 20),
	}, nil
}

func readResolutionInventoryGraph(graphPath string) (graph.Graph, error) {
	raw, err := os.ReadFile(graphPath)
	if err != nil {
		return graph.Graph{}, fmt.Errorf("read graph %s: %w", graphPath, err)
	}
	var g graph.Graph
	if err := json.Unmarshal(raw, &g); err != nil {
		return graph.Graph{}, fmt.Errorf("decode graph %s: %w", graphPath, err)
	}
	return g, nil
}

func writeResolutionInventoryResult(path string, result resolutionInventoryResult) error {
	raw, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal resolution inventory report: %w", err)
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, append(raw, '\n'), 0o644)
}

func resolutionInventorySummaryLines(result resolutionInventoryResult) []string {
	summary := result.GraphHealth
	lines := []string{
		fmt.Sprintf("resolutionInventory.nodes=%d relationships=%d gapNodes=%d gapRelationships=%d gapOccurrences=%d resolvedReferences=%d",
			result.Totals.Nodes,
			result.Totals.Relationships,
			summary.ResolutionGapNodeCount,
			summary.HasResolutionGapRelationshipCount,
			summary.ResolutionGapCount,
			summary.ResolvedReferenceCount,
		),
		fmt.Sprintf("resolutionHealth.resolvedReferences=%d unresolvedNonActionable=%d externalUnresolved=%d inRepoAnalyzerGap=%d unclassifiedUnknown=%d",
			summary.ResolutionHealthBucketCounts[string(graphhealth.ResolutionHealthResolvedReferences)],
			summary.ResolutionHealthBucketCounts[string(graphhealth.ResolutionHealthUnresolvedNonActionable)],
			summary.ResolutionHealthBucketCounts[string(graphhealth.ResolutionHealthExternalUnresolved)],
			summary.ResolutionHealthBucketCounts[string(graphhealth.ResolutionHealthInRepoAnalyzerGap)],
			summary.ResolutionHealthBucketCounts[string(graphhealth.ResolutionHealthUnclassifiedUnknown)],
		),
		fmt.Sprintf("resolutionHealth.unresolvedNonActionableBreakdown=builtin:%d,standard_library:%d,test_framework:%d",
			summary.ResolutionGapClassificationCounts[graphhealth.DiagnosticClassificationBuiltin],
			summary.ResolutionGapClassificationCounts[graphhealth.DiagnosticClassificationStandardLibrary],
			summary.ResolutionGapClassificationCounts[graphhealth.DiagnosticClassificationTestFramework],
		),
		fmt.Sprintf("resolutionHealth.targets.call=%d access=%d type=%d heritage=%d",
			summary.ResolutionHealthBucketCounts[string(graphhealth.ResolutionHealthUnresolvedCallTarget)],
			summary.ResolutionHealthBucketCounts[string(graphhealth.ResolutionHealthUnresolvedAccessTarget)],
			summary.ResolutionHealthBucketCounts[string(graphhealth.ResolutionHealthUnresolvedTypeTarget)],
			summary.ResolutionHealthBucketCounts[string(graphhealth.ResolutionHealthUnresolvedHeritageTarget)],
		),
		fmt.Sprintf("resolutionConfidence.clear=%d degraded=%d unknown=%d",
			summary.ResolutionConfidenceCounts[graphhealth.ResolutionConfidenceClear],
			summary.ResolutionConfidenceCounts[graphhealth.ResolutionConfidenceDegraded],
			summary.ResolutionConfidenceCounts[graphhealth.ResolutionConfidenceUnknown],
		),
		"resolutionGap.appLayers=" + formatCountMap(summary.ResolutionGapAppLayerCounts),
		"resolutionGap.functionalAreas=" + formatCountMap(summary.ResolutionGapFunctionalAreaCounts),
		"resolutionGap.factFamilies=" + formatCountMap(summary.ResolutionGapFactFamilyCounts),
		"resolutionGap.targetRoles=" + formatCountMap(summary.ResolutionGapTargetRoleCounts),
		"resolutionGap.classifications=" + formatCountMap(summary.ResolutionGapClassificationCounts),
		"resolutionGap.actionability=" + formatCountMap(summary.ResolutionGapActionabilityCounts),
		"resolutionGap.topology=" + formatCountMap(summary.ResolutionGapTopologyStatusCounts),
	}
	lines = append(lines, fmt.Sprintf("resolutionGap.fileGroups=%d", len(result.FileGroups)))
	for _, group := range result.FileGroups[:minInt(len(result.FileGroups), 5)] {
		lines = append(lines, fmt.Sprintf("resolutionGap.file path=%q total=%d kinds=%s classifications=%s actionability=%s nearestSymbols=%s",
			group.Path,
			group.Total,
			formatCountMap(group.ByKind),
			formatCountMap(group.ByClassification),
			formatCountMap(group.ByActionability),
			strings.Join(group.NearestSourceSymbols, ","),
		))
	}
	return lines
}

func resolutionInventoryFileGroups(g *graph.Graph, limit int) []resolutionInventoryFileGroup {
	builder := filecontext.NewBuilder(g)
	list := builder.BuildFileList(filecontext.FileListOptions{Sort: "unresolved", Limit: limit, UnresolvedOnly: true})
	groups := make([]resolutionInventoryFileGroup, 0, len(list.Files))
	for _, summary := range list.Files {
		context, ok := builder.BuildFileContext(summary.Path, filecontext.Options{
			UnresolvedSamplesPerGroup: 5,
		})
		if !ok || context.Unresolved.Total == 0 {
			continue
		}
		group := resolutionInventoryFileGroup{
			Path:             summary.Path,
			Total:            context.Unresolved.Total,
			ByKind:           context.Unresolved.ByKind,
			ByClassification: context.Unresolved.ByClassification,
			ByActionability:  context.Unresolved.ByActionability,
		}
		seenSymbols := map[string]bool{}
		for _, unresolvedGroup := range context.Unresolved.Groups {
			if unresolvedGroup.SourceSymbol != "" && !seenSymbols[unresolvedGroup.SourceSymbol] {
				seenSymbols[unresolvedGroup.SourceSymbol] = true
				group.NearestSourceSymbols = append(group.NearestSourceSymbols, unresolvedGroup.SourceSymbol)
			}
			for _, sample := range unresolvedGroup.Samples {
				if len(group.Samples) >= 10 {
					break
				}
				group.Samples = append(group.Samples, sample)
			}
		}
		if len(group.NearestSourceSymbols) > 5 {
			group.NearestSourceSymbols = group.NearestSourceSymbols[:5]
		}
		groups = append(groups, group)
	}
	return groups
}

func formatCountMap(counts map[string]int) string {
	if len(counts) == 0 {
		return ""
	}
	keys := make([]string, 0, len(counts))
	for key, count := range counts {
		if count == 0 {
			continue
		}
		keys = append(keys, key)
	}
	sort.Strings(keys)
	parts := make([]string, 0, len(keys))
	for _, key := range keys {
		parts = append(parts, fmt.Sprintf("%s:%d", key, counts[key]))
	}
	return strings.Join(parts, ",")
}
