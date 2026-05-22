package cli

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/tamnguyendinh/avmatrix-go/internal/graphaccuracy"
)

func newSourceSiteAccuracyCommand() *cobra.Command {
	var graphPath string
	var outPath string
	var maxExamples int
	var jsonOutput bool

	cmd := &cobra.Command{
		Use:   "source-site-accuracy",
		Short: "Report proof-based source-site and resolved-edge accuracy metrics",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			result, err := graphaccuracy.RunSourceSiteAccuracy(graphaccuracy.SourceSiteAccuracyOptions{
				GraphPath:   graphPath,
				OutPath:     outPath,
				MaxExamples: maxExamples,
			})
			if err != nil {
				return err
			}
			if jsonOutput || outPath == "" {
				raw, err := json.MarshalIndent(result, "", "  ")
				if err != nil {
					return fmt.Errorf("marshal source-site accuracy report: %w", err)
				}
				_, err = fmt.Fprintf(cmd.OutOrStdout(), "%s\n", raw)
				return err
			}
			if _, err := fmt.Fprintf(cmd.OutOrStdout(), "wrote %s\n", outPath); err != nil {
				return err
			}
			for _, line := range graphaccuracy.SourceSiteAccuracySummaryLines(result) {
				if _, err := fmt.Fprintln(cmd.OutOrStdout(), line); err != nil {
					return err
				}
			}
			return nil
		},
	}
	cmd.Flags().StringVar(&graphPath, "graph", filepath.Join(".avmatrix", "graph.json"), "AVmatrix graph snapshot JSON")
	cmd.Flags().StringVar(&outPath, "out", "", "write report JSON to this path")
	cmd.Flags().IntVar(&maxExamples, "max-examples", 50, "maximum examples to include per bucket")
	cmd.Flags().BoolVar(&jsonOutput, "json", false, "write full JSON report to stdout")
	return cmd
}
