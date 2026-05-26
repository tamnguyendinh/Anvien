package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func newAPICommand() *cobra.Command {
	var repoName string
	cmd := &cobra.Command{
		Use:   "api",
		Short: "Inspect API route, tool, shape, and impact surfaces",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
	cmd.PersistentFlags().StringVar(&repoName, "repo", "", "target repository")
	cmd.AddCommand(
		newAPIRouteMapCommand(&repoName),
		newAPIToolMapCommand(&repoName),
		newAPIShapeCheckCommand(&repoName),
		newAPIImpactCommand(&repoName),
	)
	return cmd
}

func newAPIRouteMapCommand(repoName *string) *cobra.Command {
	var route string
	var jsonOutput bool
	cmd := &cobra.Command{
		Use:   "route-map [route]",
		Short: "Show API route handlers, consumers, and linked flows",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			selector, err := oneArgOrFlag(args, route, "route")
			if err != nil {
				return err
			}
			return printLocalMCPToolWithJSON(cmd, "route_map", map[string]any{
				"route": emptyToNil(selector),
				"repo":  emptyToNil(*repoName),
			}, jsonOutput)
		},
	}
	cmd.Flags().StringVar(&route, "route", "", "route path filter")
	cmd.Flags().BoolVar(&jsonOutput, "json", false, "write only the JSON payload")
	return cmd
}

func newAPIToolMapCommand(repoName *string) *cobra.Command {
	var tool string
	var jsonOutput bool
	cmd := &cobra.Command{
		Use:   "tool-map [tool]",
		Short: "Show MCP/RPC tool definitions and linked flows",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			selector, err := oneArgOrFlag(args, tool, "tool")
			if err != nil {
				return err
			}
			return printLocalMCPToolWithJSON(cmd, "tool_map", map[string]any{
				"tool": emptyToNil(selector),
				"repo": emptyToNil(*repoName),
			}, jsonOutput)
		},
	}
	cmd.Flags().StringVar(&tool, "tool", "", "tool name filter")
	cmd.Flags().BoolVar(&jsonOutput, "json", false, "write only the JSON payload")
	return cmd
}

func newAPIShapeCheckCommand(repoName *string) *cobra.Command {
	var route string
	var jsonOutput bool
	cmd := &cobra.Command{
		Use:   "shape-check [route]",
		Short: "Check API response shapes against consumers",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			selector, err := oneArgOrFlag(args, route, "route")
			if err != nil {
				return err
			}
			return printLocalMCPToolWithJSON(cmd, "shape_check", map[string]any{
				"route": emptyToNil(selector),
				"repo":  emptyToNil(*repoName),
			}, jsonOutput)
		},
	}
	cmd.Flags().StringVar(&route, "route", "", "route path filter")
	cmd.Flags().BoolVar(&jsonOutput, "json", false, "write only the JSON payload")
	return cmd
}

func newAPIImpactCommand(repoName *string) *cobra.Command {
	var route string
	var filePath string
	var jsonOutput bool
	cmd := &cobra.Command{
		Use:   "impact [route]",
		Short: "Pre-change impact report for an API route handler",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			selector, err := oneArgOrFlag(args, route, "route")
			if err != nil {
				return err
			}
			return printLocalMCPToolWithJSON(cmd, "api_impact", map[string]any{
				"route": emptyToNil(selector),
				"file":  emptyToNil(filePath),
				"repo":  emptyToNil(*repoName),
			}, jsonOutput)
		},
	}
	cmd.Flags().StringVar(&route, "route", "", "route path filter")
	cmd.Flags().StringVar(&filePath, "file", "", "handler file filter")
	cmd.Flags().BoolVar(&jsonOutput, "json", false, "write only the JSON payload")
	return cmd
}

func oneArgOrFlag(args []string, flagValue string, flagName string) (string, error) {
	flagValue = strings.TrimSpace(flagValue)
	if len(args) == 0 {
		return flagValue, nil
	}
	if flagValue != "" {
		return "", fmt.Errorf("provide %s as either positional argument or --%s, not both", flagName, flagName)
	}
	return args[0], nil
}
