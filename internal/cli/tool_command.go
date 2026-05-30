package cli

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	mcpserver "github.com/tamnguyendinh/anvien/internal/mcp"
	"github.com/tamnguyendinh/anvien/internal/repo"
)

func newAugmentCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "augment <pattern>",
		Short: "Augment a search pattern with knowledge graph context",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			pattern := strings.TrimSpace(args[0])
			if len(pattern) < 3 {
				return nil
			}
			cwd, _ := os.Getwd()
			text, err := callLocalMCPTool("query", map[string]any{
				"query":       pattern,
				"limit":       3,
				"max_symbols": 5,
				"repo":        emptyToNil(cwd),
			})
			if err != nil {
				return nil
			}
			if strings.TrimSpace(text) == "" {
				return nil
			}
			_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Anvien graph context for %q:\n%s\n", pattern, text)
			return nil
		},
	}
}

func newQueryCommand() *cobra.Command {
	var repoName string
	var taskContext string
	var goal string
	var limit string
	var includeContent bool
	var showLanes bool
	var explain bool
	var jsonOutput bool

	cmd := &cobra.Command{
		Use:   "query <search_query>",
		Short: "Search the knowledge graph for code, owner, flow, and command-surface candidates",
		Long:  queryCommandLongDescription(),
		Args: func(cmd *cobra.Command, args []string) error {
			if showLanes {
				if len(args) > 0 {
					return fmt.Errorf("query --lanes does not accept a search query")
				}
				return nil
			}
			if len(args) != 1 {
				return fmt.Errorf("usage: anvien query <search_query>")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if showLanes {
				return printQueryCapabilityLanes(cmd, jsonOutput)
			}
			return runQueryToolCommand(cmd, args[0], "", repoName, taskContext, goal, limit, includeContent, explain, jsonOutput)
		},
	}
	cmd.PersistentFlags().StringVarP(&repoName, "repo", "r", "", "target repository (omit if only one indexed)")
	cmd.PersistentFlags().StringVarP(&taskContext, "context", "c", "", "task context to improve ranking")
	cmd.PersistentFlags().StringVarP(&goal, "goal", "g", "", "what you want to find")
	cmd.PersistentFlags().StringVarP(&limit, "limit", "l", "", "max results to return")
	cmd.PersistentFlags().BoolVar(&includeContent, "content", false, "include full symbol source code")
	cmd.PersistentFlags().BoolVar(&explain, "explain", false, "include lane, rank, and match evidence in query output")
	cmd.PersistentFlags().BoolVar(&jsonOutput, "json", false, "write only the JSON payload")
	cmd.Flags().BoolVar(&showLanes, "lanes", false, "list query capability lanes")
	cmd.AddCommand(
		newQueryTargetCommand("files", "Search files first and include matched symbols and file summaries", &repoName, &taskContext, &goal, &limit, &includeContent, &explain, &jsonOutput),
		newQueryTargetCommand("symbols", "Search symbols first and include containing file summaries", &repoName, &taskContext, &goal, &limit, &includeContent, &explain, &jsonOutput),
		newQueryTargetCommand("flows", "Search execution flows only", &repoName, &taskContext, &goal, &limit, &includeContent, &explain, &jsonOutput),
		newQueryTargetCommand("api", "Search API routes and MCP tools only", &repoName, &taskContext, &goal, &limit, &includeContent, &explain, &jsonOutput),
	)
	return cmd
}

func newQueryTargetCommand(targetType string, short string, repoName *string, taskContext *string, goal *string, limit *string, includeContent *bool, explain *bool, jsonOutput *bool) *cobra.Command {
	return &cobra.Command{
		Use:   targetType + " <search_query>",
		Short: short,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) > 1 {
				return fmt.Errorf("usage: anvien query %s <search_query>", targetType)
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return runQueryToolCommand(cmd, targetType, "", *repoName, *taskContext, *goal, *limit, *includeContent, *explain, *jsonOutput)
			}
			return runQueryToolCommand(cmd, args[0], targetType, *repoName, *taskContext, *goal, *limit, *includeContent, *explain, *jsonOutput)
		},
	}
}

func runQueryToolCommand(cmd *cobra.Command, query string, targetType string, repoName string, taskContext string, goal string, limit string, includeContent bool, explain bool, jsonOutput bool) error {
	toolArgs := map[string]any{
		"query":           query,
		"repo":            emptyToNil(repoName),
		"task_context":    emptyToNil(taskContext),
		"goal":            emptyToNil(goal),
		"include_content": includeContent,
		"explain":         explain,
	}
	if targetType != "" {
		toolArgs["target_type"] = targetType
		toolArgs["dispatch_mode"] = "explicit"
	}
	if limit != "" {
		parsed, err := parsePositiveIntFlag("limit", limit)
		if err != nil {
			return err
		}
		toolArgs["limit"] = parsed
	}
	return printLocalMCPToolWithJSON(cmd, "query", toolArgs, jsonOutput)
}

func queryCommandLongDescription() string {
	var builder strings.Builder
	builder.WriteString("Search the knowledge graph for code, owner, flow, and command-surface candidates.\n\n")
	builder.WriteString("Query capability lanes:\n")
	for _, lane := range mcpserver.QueryCapabilityLanes() {
		fmt.Fprintf(&builder, "  - %s: %s\n", lane.ID, lane.Description)
	}
	builder.WriteString("\nUse --lanes to list lanes, and --explain to include lane, rank, and match evidence in query output.")
	return builder.String()
}

func printQueryCapabilityLanes(cmd *cobra.Command, jsonOutput bool) error {
	lanes := mcpserver.QueryCapabilityLanes()
	if jsonOutput {
		raw, err := json.MarshalIndent(map[string]any{"queryCapabilities": lanes}, "", "  ")
		if err != nil {
			return err
		}
		_, err = fmt.Fprintf(cmd.OutOrStdout(), "%s\n", raw)
		return err
	}
	for _, lane := range lanes {
		if _, err := fmt.Fprintf(cmd.OutOrStdout(), "%s\t%s\t%s\n", lane.ID, lane.Name, lane.Description); err != nil {
			return err
		}
	}
	return nil
}

func newContextCommand() *cobra.Command {
	var repoName string
	var uid string
	var filePath string
	var includeContent bool
	var jsonOutput bool

	cmd := &cobra.Command{
		Use:   "context [name]",
		Short: "360-degree view of a code symbol: callers, callees, processes",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := ""
			if len(args) > 0 {
				name = args[0]
			}
			if strings.TrimSpace(name) == "" && strings.TrimSpace(uid) == "" {
				return fmt.Errorf("usage: anvien context <symbol_name> [--uid <uid>] [--file <path>]")
			}
			targetType := "auto"
			dispatchMode := "smart"
			if strings.TrimSpace(uid) != "" {
				targetType = "symbol"
				dispatchMode = "explicit"
			}
			return printLocalMCPToolWithJSON(cmd, "context", map[string]any{
				"name":            emptyToNil(name),
				"uid":             emptyToNil(uid),
				"file_path":       emptyToNil(filePath),
				"target_type":     targetType,
				"dispatch_mode":   dispatchMode,
				"include_content": includeContent,
				"repo":            emptyToNil(repoName),
			}, jsonOutput)
		},
	}
	cmd.PersistentFlags().StringVarP(&repoName, "repo", "r", "", "target repository")
	cmd.PersistentFlags().StringVarP(&uid, "uid", "u", "", "direct symbol UID")
	cmd.PersistentFlags().StringVarP(&filePath, "file", "f", "", "file path to disambiguate common names")
	cmd.PersistentFlags().BoolVar(&includeContent, "content", false, "include full symbol source code")
	cmd.PersistentFlags().BoolVar(&jsonOutput, "json", false, "write only the JSON payload")
	cmd.AddCommand(
		&cobra.Command{
			Use:   "symbol <symbol>",
			Short: "Force symbol context and include containing file summary",
			Args:  cobra.MaximumNArgs(1),
			RunE: func(cmd *cobra.Command, args []string) error {
				if len(args) == 0 {
					return printLocalMCPToolWithJSON(cmd, "context", map[string]any{
						"name":            "symbol",
						"uid":             emptyToNil(uid),
						"file_path":       emptyToNil(filePath),
						"target_type":     "auto",
						"dispatch_mode":   "smart",
						"include_content": includeContent,
						"repo":            emptyToNil(repoName),
					}, jsonOutput)
				}
				return printLocalMCPToolWithJSON(cmd, "context", map[string]any{
					"name":            args[0],
					"uid":             emptyToNil(uid),
					"file_path":       emptyToNil(filePath),
					"target_type":     "symbol",
					"dispatch_mode":   "explicit",
					"include_content": includeContent,
					"repo":            emptyToNil(repoName),
				}, jsonOutput)
			},
		},
		&cobra.Command{
			Use:   "file <path>",
			Short: "Force file context and avoid symbol ambiguity",
			Args:  cobra.MaximumNArgs(1),
			RunE: func(cmd *cobra.Command, args []string) error {
				if len(args) == 0 {
					return printLocalMCPToolWithJSON(cmd, "context", map[string]any{
						"name":            "file",
						"uid":             emptyToNil(uid),
						"file_path":       emptyToNil(filePath),
						"target_type":     "auto",
						"dispatch_mode":   "smart",
						"include_content": includeContent,
						"repo":            emptyToNil(repoName),
					}, jsonOutput)
				}
				return printLocalMCPToolWithJSON(cmd, "context", map[string]any{
					"name":          args[0],
					"target_type":   "file",
					"dispatch_mode": "explicit",
					"repo":          emptyToNil(repoName),
				}, jsonOutput)
			},
		},
	)
	return cmd
}

func newImpactCommand() *cobra.Command {
	var direction string
	var repoName string
	var uid string
	var depth string
	var includeTests bool
	var jsonOutput bool

	cmd := &cobra.Command{
		Use:   "impact [target]",
		Short: "Blast radius analysis: what breaks if you change a symbol",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			target := ""
			if len(args) > 0 {
				target = args[0]
			}
			if strings.TrimSpace(target) == "" && strings.TrimSpace(uid) == "" {
				return fmt.Errorf("usage: anvien impact [symbol_name] [--uid <uid>] [--direction upstream|downstream]")
			}
			if direction != "upstream" && direction != "downstream" {
				return fmt.Errorf("direction must be upstream or downstream")
			}
			targetType := "auto"
			dispatchMode := "smart"
			if strings.TrimSpace(uid) != "" {
				targetType = "symbol"
				dispatchMode = "explicit"
			}
			toolArgs := map[string]any{
				"target":        emptyToNil(target),
				"target_uid":    emptyToNil(uid),
				"target_type":   targetType,
				"dispatch_mode": dispatchMode,
				"direction":     direction,
				"includeTests":  includeTests,
				"repo":          emptyToNil(repoName),
			}
			if depth != "" {
				parsed, err := parsePositiveIntFlag("depth", depth)
				if err != nil {
					return err
				}
				toolArgs["maxDepth"] = parsed
			}
			return printLocalMCPToolWithJSON(cmd, "impact", toolArgs, jsonOutput)
		},
	}
	cmd.PersistentFlags().StringVarP(&direction, "direction", "d", "upstream", "upstream or downstream")
	cmd.PersistentFlags().StringVarP(&repoName, "repo", "r", "", "target repository")
	cmd.PersistentFlags().StringVarP(&uid, "uid", "u", "", "direct symbol UID")
	cmd.PersistentFlags().StringVar(&depth, "depth", "", "max relationship depth")
	cmd.PersistentFlags().BoolVar(&includeTests, "include-tests", false, "include test files in results")
	cmd.PersistentFlags().BoolVar(&jsonOutput, "json", false, "write only the JSON payload")
	cmd.AddCommand(
		newImpactTargetCommand("symbol", "Force symbol blast radius with file-layer evidence", &direction, &repoName, &uid, &depth, &includeTests, &jsonOutput),
		newImpactTargetCommand("file", "Aggregate blast radius from all symbols in one file", &direction, &repoName, &uid, &depth, &includeTests, &jsonOutput),
		newImpactTargetCommand("route", "Inspect route handler, consumer, shape, and flow impact", &direction, &repoName, &uid, &depth, &includeTests, &jsonOutput),
		newImpactTargetCommand("tool", "Inspect MCP tool definition and linked flow impact", &direction, &repoName, &uid, &depth, &includeTests, &jsonOutput),
	)
	return cmd
}

func newImpactTargetCommand(targetType string, short string, direction *string, repoName *string, uid *string, depth *string, includeTests *bool, jsonOutput *bool) *cobra.Command {
	return &cobra.Command{
		Use:   targetType + " <target>",
		Short: short,
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return runImpactTargetCommand(cmd, targetType, "auto", *direction, *repoName, *uid, *depth, *includeTests, *jsonOutput)
			}
			return runImpactTargetCommand(cmd, args[0], targetType, *direction, *repoName, *uid, *depth, *includeTests, *jsonOutput)
		},
	}
}

func runImpactTargetCommand(cmd *cobra.Command, target string, targetType string, direction string, repoName string, uid string, depth string, includeTests bool, jsonOutput bool) error {
	toolArgs := map[string]any{
		"target":        target,
		"target_uid":    emptyToNil(uid),
		"target_type":   targetType,
		"dispatch_mode": "explicit",
		"direction":     direction,
		"includeTests":  includeTests,
		"repo":          emptyToNil(repoName),
	}
	if targetType == "auto" {
		toolArgs["dispatch_mode"] = "smart"
	}
	if depth != "" {
		parsed, err := parsePositiveIntFlag("depth", depth)
		if err != nil {
			return err
		}
		toolArgs["maxDepth"] = parsed
	}
	return printLocalMCPToolWithJSON(cmd, "impact", toolArgs, jsonOutput)
}

func newRenameCommand() *cobra.Command {
	var repoName string
	var uid string
	var filePath string
	var apply bool
	var jsonOutput bool

	cmd := &cobra.Command{
		Use:   "rename [symbol_name] <new_name>",
		Short: "Preview or apply a graph-guided symbol rename",
		Args: func(cmd *cobra.Command, args []string) error {
			if strings.TrimSpace(uid) != "" {
				if len(args) != 1 {
					return fmt.Errorf("usage: anvien rename --uid <symbol_uid> <new_name>")
				}
				return nil
			}
			if len(args) != 2 {
				return fmt.Errorf("usage: anvien rename <symbol_name> <new_name>")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			symbolName := ""
			newName := ""
			if strings.TrimSpace(uid) != "" {
				newName = args[0]
			} else {
				symbolName = args[0]
				newName = args[1]
			}
			return printLocalMCPToolWithJSON(cmd, "rename", map[string]any{
				"symbol_name": emptyToNil(symbolName),
				"symbol_uid":  emptyToNil(uid),
				"new_name":    newName,
				"file_path":   emptyToNil(filePath),
				"dry_run":     !apply,
				"repo":        emptyToNil(repoName),
			}, jsonOutput)
		},
	}
	cmd.Flags().StringVarP(&repoName, "repo", "r", "", "target repository")
	cmd.Flags().StringVarP(&uid, "uid", "u", "", "direct symbol UID")
	cmd.Flags().StringVarP(&filePath, "file", "f", "", "file path to disambiguate common names")
	cmd.Flags().BoolVar(&apply, "apply", false, "apply edits instead of running a dry run")
	cmd.Flags().BoolVar(&jsonOutput, "json", false, "write only the JSON payload")
	return cmd
}

func newCypherCommand() *cobra.Command {
	var repoName string
	cmd := &cobra.Command{
		Use:   "cypher <query>",
		Short: "Execute raw Cypher query against the knowledge graph",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return printLocalMCPTool(cmd, "cypher", map[string]any{
				"query": args[0],
				"repo":  emptyToNil(repoName),
			})
		},
	}
	cmd.Flags().StringVarP(&repoName, "repo", "r", "", "target repository")
	return cmd
}

func newDetectChangesCommand() *cobra.Command {
	var scope string
	var baseRef string
	var repoName string
	var jsonOutput bool

	cmd := &cobra.Command{
		Use:   "detect-changes",
		Short: "Analyze uncommitted git changes and find affected execution flows",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return printLocalMCPToolWithJSON(cmd, "detect_changes", map[string]any{
				"scope":    scope,
				"base_ref": emptyToNil(baseRef),
				"repo":     emptyToNil(repoName),
			}, jsonOutput)
		},
	}
	cmd.PersistentFlags().StringVarP(&scope, "scope", "s", "unstaged", "unstaged, staged, all, or compare")
	cmd.PersistentFlags().StringVar(&baseRef, "base-ref", "", "branch/commit for compare scope")
	cmd.PersistentFlags().StringVarP(&repoName, "repo", "r", "", "target repository")
	cmd.PersistentFlags().BoolVar(&jsonOutput, "json", false, "write only the JSON payload")
	cmd.AddCommand(
		newDetectChangesTargetCommand("files", "Show changed files only", &scope, &baseRef, &repoName, &jsonOutput),
		newDetectChangesTargetCommand("symbols", "Show changed symbols only", &scope, &baseRef, &repoName, &jsonOutput),
		newDetectChangesTargetCommand("flows", "Show affected flows only", &scope, &baseRef, &repoName, &jsonOutput),
	)
	return cmd
}

func newDetectChangesTargetCommand(targetType string, short string, scope *string, baseRef *string, repoName *string, jsonOutput *bool) *cobra.Command {
	return &cobra.Command{
		Use:   targetType,
		Short: short,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return printLocalMCPToolWithJSON(cmd, "detect_changes", map[string]any{
				"scope":         *scope,
				"base_ref":      emptyToNil(*baseRef),
				"target_type":   targetType,
				"dispatch_mode": "explicit",
				"repo":          emptyToNil(*repoName),
			}, *jsonOutput)
		},
	}
}

func printLocalMCPTool(cmd *cobra.Command, toolName string, args map[string]any) error {
	return printLocalMCPToolWithJSON(cmd, toolName, args, false)
}

func printLocalMCPToolWithJSON(cmd *cobra.Command, toolName string, args map[string]any, jsonOutput bool) error {
	text, err := callLocalMCPTool(toolName, args)
	if err != nil {
		return err
	}
	if jsonOutput {
		text = primaryMCPToolPayload(text)
	}
	_, err = fmt.Fprintln(cmd.OutOrStdout(), text)
	return err
}

func primaryMCPToolPayload(text string) string {
	if before, _, ok := strings.Cut(text, "\n\n---\n"); ok {
		return strings.TrimSpace(before)
	}
	return strings.TrimSpace(text)
}

func callLocalMCPTool(toolName string, args map[string]any) (string, error) {
	request := map[string]any{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  "tools/call",
		"params": map[string]any{
			"name":      toolName,
			"arguments": stripNilValues(args),
		},
	}
	raw, err := json.Marshal(request)
	if err != nil {
		return "", err
	}
	server := mcpserver.NewServer(mcpserver.Config{Store: repo.NewEnvStore()})
	responseRaw, ok, err := server.HandleJSONRPC(raw)
	if err != nil {
		return "", err
	}
	if !ok {
		return "", fmt.Errorf("tool %q did not return a response", toolName)
	}
	var response struct {
		Result struct {
			Content []struct {
				Type string `json:"type"`
				Text string `json:"text"`
			} `json:"content"`
		} `json:"result"`
		Error *struct {
			Message string `json:"message"`
		} `json:"error"`
	}
	if err := json.Unmarshal(responseRaw, &response); err != nil {
		return "", err
	}
	if response.Error != nil {
		return "", errors.New(response.Error.Message)
	}
	if len(response.Result.Content) == 0 {
		return "", nil
	}
	return response.Result.Content[0].Text, nil
}

func emptyToNil(value string) any {
	if strings.TrimSpace(value) == "" {
		return nil
	}
	return value
}

func stripNilValues(input map[string]any) map[string]any {
	out := make(map[string]any, len(input))
	for key, value := range input {
		if value != nil {
			out[key] = value
		}
	}
	return out
}

func parsePositiveIntFlag(name string, raw string) (int, error) {
	value, err := strconv.Atoi(raw)
	if err != nil || value <= 0 {
		return 0, fmt.Errorf("%s must be a positive integer", name)
	}
	return value, nil
}
