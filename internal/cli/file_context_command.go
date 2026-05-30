package cli

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tamnguyendinh/anvien/internal/filecontext"
	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/repo"
)

const (
	defaultHumanSymbolRoots        = 20
	defaultHumanRelationshipGroups = 10
	defaultHumanUnresolvedGroups   = 5
)

type fileProjectionCommandInputs struct {
	Repo          string
	RepoPath      string
	Graph         string
	IndexedCommit string
	CurrentCommit string
	IndexedAt     string
}

type fileHotspotsPayload struct {
	Repo     string                    `json:"repo,omitempty"`
	RepoPath string                    `json:"repoPath,omitempty"`
	Graph    filecontext.GraphInfo     `json:"graph"`
	Total    int                       `json:"total"`
	Offset   int                       `json:"offset"`
	Limit    int                       `json:"limit"`
	Sort     string                    `json:"sort"`
	Files    []filecontext.FileSummary `json:"files"`
}

func newFileContextCommand() *cobra.Command {
	var repoName string
	var jsonOutput bool
	var relationshipSamples int
	var unresolvedSamples int
	var linkedSamples int

	cmd := &cobra.Command{
		Use:   "file-context <path>",
		Short: "Show file-first graph context for one indexed file",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			inputs, g, err := loadFileProjectionGraph(repoName)
			if err != nil {
				return err
			}
			context, ok := filecontext.NewBuilder(g).BuildFileContext(args[0], filecontext.Options{
				RelationshipSamplesPerGroup: relationshipSamples,
				UnresolvedSamplesPerGroup:   unresolvedSamples,
				LinkedSamplesPerKind:        linkedSamples,
			})
			if !ok {
				return fmt.Errorf("file %q not found in repo %s", args[0], inputs.Repo)
			}
			attachFileProjectionMetadata(&context, inputs)
			if jsonOutput {
				return writeJSON(cmd, context)
			}
			return renderFileContext(cmd, context)
		},
	}
	cmd.Flags().StringVarP(&repoName, "repo", "r", "", "target repository")
	cmd.Flags().BoolVar(&jsonOutput, "json", false, "write full file context JSON")
	cmd.Flags().IntVar(&relationshipSamples, "relationships", 5, "relationship samples per group")
	cmd.Flags().IntVar(&unresolvedSamples, "unresolved", 5, "unresolved source-site samples per group")
	cmd.Flags().IntVar(&linkedSamples, "linked", 5, "linked overlay samples per kind")
	return cmd
}

func newFileHotspotsCommand() *cobra.Command {
	var repoName string
	var jsonOutput bool
	var sortMode string
	var limit int
	var offset int
	var kinds []string
	var appLayer string
	var functionalArea string
	var apiOnly bool
	var unresolvedOnly bool
	var highFanInOnly bool
	var highFanOutOnly bool
	var highFanInThreshold int
	var highFanOutThreshold int

	cmd := &cobra.Command{
		Use:   "file-hotspots",
		Short: "List file-level graph hotspots by unresolved, fan-in, fan-out, or symbol count",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if limit < 0 {
				return fmt.Errorf("limit must be zero or positive")
			}
			if offset < 0 {
				return fmt.Errorf("offset must be zero or positive")
			}
			inputs, g, err := loadFileProjectionGraph(repoName)
			if err != nil {
				return err
			}
			list := filecontext.NewBuilder(g).BuildFileList(filecontext.FileListOptions{
				Sort:                sortMode,
				Limit:               limit,
				Offset:              offset,
				Kinds:               kinds,
				AppLayer:            appLayer,
				FunctionalArea:      functionalArea,
				APIOnly:             apiOnly,
				UnresolvedOnly:      unresolvedOnly,
				HighFanInOnly:       highFanInOnly,
				HighFanOutOnly:      highFanOutOnly,
				HighFanInThreshold:  highFanInThreshold,
				HighFanOutThreshold: highFanOutThreshold,
			})
			payload := fileHotspotsPayload{
				Repo:     inputs.Repo,
				RepoPath: inputs.RepoPath,
				Graph:    fileProjectionGraphInfo(inputs),
				Total:    list.Total,
				Offset:   list.Offset,
				Limit:    list.Limit,
				Sort:     list.Sort,
				Files:    list.Files,
			}
			if jsonOutput {
				return writeJSON(cmd, payload)
			}
			return renderFileHotspots(cmd, payload)
		},
	}
	cmd.Flags().StringVarP(&repoName, "repo", "r", "", "target repository")
	cmd.Flags().BoolVar(&jsonOutput, "json", false, "write file hotspot JSON")
	cmd.Flags().StringVar(&sortMode, "sort", "unresolved", "sort by path, unresolved, fan-in, fan-out, symbols, flows, or tests")
	cmd.Flags().IntVar(&limit, "limit", 20, "maximum files to show; 0 means all")
	cmd.Flags().IntVar(&offset, "offset", 0, "files to skip before returning results")
	cmd.Flags().StringSliceVar(&kinds, "kind", nil, "filter by file kind; repeat or comma-separate values")
	cmd.Flags().StringVar(&appLayer, "app-layer", "", "filter by app layer")
	cmd.Flags().StringVar(&functionalArea, "functional-area", "", "filter by functional area")
	cmd.Flags().BoolVar(&apiOnly, "api-only", false, "show only API-related files")
	cmd.Flags().BoolVar(&unresolvedOnly, "unresolved-only", false, "show only files with unresolved source sites")
	cmd.Flags().BoolVar(&highFanInOnly, "high-fan-in", false, "show only files at or above the fan-in threshold")
	cmd.Flags().BoolVar(&highFanOutOnly, "high-fan-out", false, "show only files at or above the fan-out threshold")
	cmd.Flags().IntVar(&highFanInThreshold, "high-fan-in-threshold", 10, "fan-in threshold for --high-fan-in")
	cmd.Flags().IntVar(&highFanOutThreshold, "high-fan-out-threshold", 10, "fan-out threshold for --high-fan-out")
	return cmd
}

func loadFileProjectionGraph(repoName string) (fileProjectionCommandInputs, *graph.Graph, error) {
	store := repo.NewEnvStore()
	entries, err := store.ListRegistered(false)
	if err != nil {
		return fileProjectionCommandInputs{}, nil, err
	}
	var entry repo.RegistryEntry
	if strings.TrimSpace(repoName) == "" {
		if len(entries) != 1 {
			return fileProjectionCommandInputs{}, nil, fmt.Errorf("file projection requires exactly one indexed repo or --repo; run anvien list")
		}
		entry = entries[0]
	} else {
		resolved, err := repo.ResolveEntry(entries, repoName)
		if err != nil {
			return fileProjectionCommandInputs{}, nil, err
		}
		entry = resolved
	}
	currentCommit := repo.CurrentCommit(entry.Path)
	if currentCommit != "" && entry.LastCommit != "" && currentCommit != entry.LastCommit {
		return fileProjectionCommandInputs{}, nil, fmt.Errorf("file projection requires fresh analyze output for %s: indexed commit %s current commit %s; run anvien analyze --force", entry.Name, shortCommit(entry.LastCommit), shortCommit(currentCommit))
	}
	storagePath := entry.StoragePath
	if storagePath == "" {
		storagePath = repo.StoragePath(entry.Path)
	}
	graphPath := filepath.Join(storagePath, "graph.json")
	raw, err := os.ReadFile(graphPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return fileProjectionCommandInputs{}, nil, fmt.Errorf("file projection requires fresh analyze output for %s: graph not found at %s; run anvien analyze --force", entry.Name, graphPath)
		}
		return fileProjectionCommandInputs{}, nil, fmt.Errorf("read graph %s: %w", graphPath, err)
	}
	var g graph.Graph
	if err := json.Unmarshal(raw, &g); err != nil {
		return fileProjectionCommandInputs{}, nil, fmt.Errorf("decode graph %s: %w", graphPath, err)
	}
	return fileProjectionCommandInputs{
		Repo:          entry.Name,
		RepoPath:      entry.Path,
		Graph:         graphPath,
		IndexedCommit: entry.LastCommit,
		CurrentCommit: currentCommit,
		IndexedAt:     entry.IndexedAt,
	}, &g, nil
}

func attachFileProjectionMetadata(context *filecontext.FileContext, inputs fileProjectionCommandInputs) {
	filecontext.AttachMetadata(context, inputs.Repo, inputs.RepoPath, fileProjectionGraphInfo(inputs))
}

func fileProjectionGraphInfo(inputs fileProjectionCommandInputs) filecontext.GraphInfo {
	return filecontext.GraphInfo{
		Path:          inputs.Graph,
		IndexedCommit: inputs.IndexedCommit,
		CurrentCommit: inputs.CurrentCommit,
		Stale:         inputs.CurrentCommit != "" && inputs.IndexedCommit != "" && inputs.CurrentCommit != inputs.IndexedCommit,
		AnalyzedAt:    inputs.IndexedAt,
	}
}

func renderFileContext(cmd *cobra.Command, context filecontext.FileContext) error {
	out := cmd.OutOrStdout()
	summary := context.Summary
	if _, err := fmt.Fprintf(out, "File: %s\n", summary.Path); err != nil {
		return err
	}
	if _, err := fmt.Fprintf(out, "Repo: %s\n", context.Repo); err != nil {
		return err
	}
	if _, err := fmt.Fprintf(out, "Language: %s  Kind: %s  Layer: %s  Area: %s  Risk: %s\n",
		defaultString(summary.Language, "unknown"),
		defaultString(summary.Kind, "unknown"),
		defaultString(summary.AppLayer, "unknown"),
		defaultString(summary.FunctionalArea, "unknown"),
		defaultString(summary.Risk, "unknown"),
	); err != nil {
		return err
	}
	if _, err := fmt.Fprintf(out, "Symbols: %d exported=%d\n", summary.SymbolCount, summary.ExportedSymbolCount); err != nil {
		return err
	}
	if _, err := fmt.Fprintf(out, "Relationships: local=%d inbound=%d outbound=%d unresolved=%d\n",
		summary.LocalRelationshipCount,
		summary.InboundRefCount,
		summary.OutboundRefCount,
		summary.UnresolvedSourceSiteCount,
	); err != nil {
		return err
	}

	if _, err := fmt.Fprintln(out, "\nSymbol tree:"); err != nil {
		return err
	}
	if len(context.SymbolTree) == 0 {
		if _, err := fmt.Fprintln(out, "  (none)"); err != nil {
			return err
		}
	} else {
		shown := minInt(len(context.SymbolTree), defaultHumanSymbolRoots)
		for _, node := range context.SymbolTree[:shown] {
			if err := renderSymbolTreeNode(out, node, "  "); err != nil {
				return err
			}
		}
		if len(context.SymbolTree) > shown {
			if _, err := fmt.Fprintf(out, "  ... %d more symbol roots (use --json for full tree)\n", len(context.SymbolTree)-shown); err != nil {
				return err
			}
		}
	}

	if err := renderFileRelationshipGroups(out, "Outbound files", context.Relationships.OutboundByFile, defaultHumanRelationshipGroups); err != nil {
		return err
	}
	if err := renderFileRelationshipGroups(out, "Inbound files", context.Relationships.InboundByFile, defaultHumanRelationshipGroups); err != nil {
		return err
	}
	return renderUnresolvedGroups(out, context.Unresolved, defaultHumanUnresolvedGroups)
}

func renderFileHotspots(cmd *cobra.Command, payload fileHotspotsPayload) error {
	out := cmd.OutOrStdout()
	if _, err := fmt.Fprintf(out, "File hotspots for %s: total=%d showing=%d sort=%s offset=%d limit=%d\n",
		payload.Repo,
		payload.Total,
		len(payload.Files),
		payload.Sort,
		payload.Offset,
		payload.Limit,
	); err != nil {
		return err
	}
	if _, err := fmt.Fprintln(out, "Path\tLayer\tArea\tSymbols\tIn\tOut\tLocal\tUnresolved\tRisk"); err != nil {
		return err
	}
	for _, file := range payload.Files {
		if _, err := fmt.Fprintf(out, "%s\t%s\t%s\t%d\t%d\t%d\t%d\t%d\t%s\n",
			file.Path,
			defaultString(file.AppLayer, "-"),
			defaultString(file.FunctionalArea, "-"),
			file.SymbolCount,
			file.InboundRefCount,
			file.OutboundRefCount,
			file.LocalRelationshipCount,
			file.UnresolvedSourceSiteCount,
			defaultString(file.Risk, "-"),
		); err != nil {
			return err
		}
	}
	return nil
}

func renderSymbolTreeNode(out io.Writer, node filecontext.SymbolTreeNode, indent string) error {
	lineRange := ""
	if node.Range.StartLine > 0 {
		lineRange = fmt.Sprintf(" lines=%d-%d", node.Range.StartLine, node.Range.EndLine)
	}
	exported := ""
	if node.Exported {
		exported = " exported"
	}
	if _, err := fmt.Fprintf(out, "%s- %s %s%s%s inbound=%d outbound=%d unresolved=%d\n",
		indent,
		node.Name,
		node.Kind,
		lineRange,
		exported,
		node.RelationshipCounts.Inbound,
		node.RelationshipCounts.Outbound,
		node.RelationshipCounts.Unresolved,
	); err != nil {
		return err
	}
	for _, child := range node.Children {
		if err := renderSymbolTreeNode(out, child, indent+"  "); err != nil {
			return err
		}
	}
	return nil
}

func renderFileRelationshipGroups(out io.Writer, title string, groups []filecontext.FileRelationshipGroup, limit int) error {
	if _, err := fmt.Fprintf(out, "\n%s:\n", title); err != nil {
		return err
	}
	if len(groups) == 0 {
		_, err := fmt.Fprintln(out, "  (none)")
		return err
	}
	shown := minInt(len(groups), limit)
	for _, group := range groups[:shown] {
		if _, err := fmt.Fprintf(out, "  %s total=%d %s\n", group.File, group.Total, formatCountMap(group.Counts)); err != nil {
			return err
		}
	}
	if len(groups) > shown {
		if _, err := fmt.Fprintf(out, "  ... %d more files (use --json for full groups)\n", len(groups)-shown); err != nil {
			return err
		}
	}
	return nil
}

func renderUnresolvedGroups(out io.Writer, unresolved filecontext.UnresolvedSummary, limit int) error {
	if _, err := fmt.Fprintf(out, "\nUnresolved source sites: total=%d\n", unresolved.Total); err != nil {
		return err
	}
	if unresolved.Total == 0 {
		_, err := fmt.Fprintln(out, "  (none)")
		return err
	}
	shown := minInt(len(unresolved.Groups), limit)
	for _, group := range unresolved.Groups[:shown] {
		if _, err := fmt.Fprintf(out, "  %s total=%d\n", defaultString(group.SourceSymbol, "(unknown symbol)"), group.Total); err != nil {
			return err
		}
		for _, sample := range group.Samples {
			if _, err := fmt.Fprintf(out, "    line %d:%d %s %s/%s\n",
				sample.Line,
				sample.Column,
				defaultString(sample.TargetText, "(unknown target)"),
				defaultString(sample.GapKind, "unknown_gap"),
				defaultString(sample.Actionability, "unknown"),
			); err != nil {
				return err
			}
		}
	}
	if len(unresolved.Groups) > shown {
		if _, err := fmt.Fprintf(out, "  ... %d more source symbols (use --json for full unresolved samples)\n", len(unresolved.Groups)-shown); err != nil {
			return err
		}
	}
	return nil
}

func writeJSON(cmd *cobra.Command, payload any) error {
	raw, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(cmd.OutOrStdout(), "%s\n", raw)
	return err
}

func defaultString(value string, fallback string) string {
	if strings.TrimSpace(value) == "" {
		return fallback
	}
	return value
}

func minInt(left int, right int) int {
	if left < right {
		return left
	}
	return right
}
