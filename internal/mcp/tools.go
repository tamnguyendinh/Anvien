package mcp

import (
	"errors"
	"fmt"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/tamnguyendinh/anvien/internal/filecontext"
	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/lbugnative"
	"github.com/tamnguyendinh/anvien/internal/lbugruntime"
	"github.com/tamnguyendinh/anvien/internal/repo"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
	"github.com/tamnguyendinh/anvien/internal/semantic"
)

var (
	mcpProcessIDLiteralPattern = regexp.MustCompile(`Process\s*\{\s*id\s*:\s*'((?:[^']|'')*)'`)
	mcpProcessIDInPattern      = regexp.MustCompile(`p\.id\s+IN\s+\[([^\]]*)\]`)
	mcpFromIDInPattern         = regexp.MustCompile(`from\.id\s+IN\s+\[([^\]]*)\]`)
	mcpToIDInPattern           = regexp.MustCompile(`to\.id\s+IN\s+\[([^\]]*)\]`)
	mcpQuotedStringPattern     = regexp.MustCompile(`'((?:[^']|'')*)'`)
	mcpNodeLabelQueryPattern   = regexp.MustCompile(`MATCH\s+\(n:(Function|Class|Interface)\)`)
	mcpLimitPattern            = regexp.MustCompile(`(?i)\bLIMIT\s+(\d+)`)
)

var errUnsupportedMCPGraphQuery = errors.New("unsupported graph query in Go MCP graph adapter")

// QueryCapabilityLane describes a user-facing retrieval lane under the umbrella query command.
type QueryCapabilityLane struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Keywords    []string `json:"keywords"`
}

var queryCapabilityLaneDefinitions = []QueryCapabilityLane{
	{
		ID:          "owner_discovery",
		Name:        "Owner discovery",
		Description: "Find the file, symbol, command, resource, generated artifact, or package that owns a problem.",
		Keywords:    []string{"owner", "owns", "source", "file", "symbol", "function", "method", "class", "where", "surface", "implementation"},
	},
	{
		ID:          "concept_discovery",
		Name:        "Concept discovery",
		Description: "Find likely code areas from broad natural-language intent.",
		Keywords:    []string{"concept", "behavior", "feature", "flow", "logic", "how", "why", "work"},
	},
	{
		ID:          "execution_flow_discovery",
		Name:        "Execution-flow discovery",
		Description: "Find processes, flows, and process steps related to the intent.",
		Keywords:    []string{"process", "flow", "execution", "step", "trace", "call", "calls"},
	},
	{
		ID:          "api_surface_discovery",
		Name:        "API surface discovery",
		Description: "Find route/tool handlers, API contracts, generated API types, and consumers.",
		Keywords:    []string{"api", "route", "handler", "tool", "mcp", "contract", "shape", "impact", "consumer", "response"},
	},
	{
		ID:          "graph_quality_discovery",
		Name:        "Graph-quality discovery",
		Description: "Find query-health, source-site, resolution, ResolutionGap, graph-health, and accuracy surfaces.",
		Keywords:    []string{"graph", "quality", "health", "query", "resolution", "inventory", "source", "site", "accuracy", "benchmark", "gap"},
	},
	{
		ID:          "docs_setup_ai_context_discovery",
		Name:        "Docs/setup/AI-context discovery",
		Description: "Find generated guidance, skill source, setup, package, and AI context generation surfaces.",
		Keywords:    []string{"doc", "docs", "setup", "skill", "skills", "agent", "agents", "claude", "aicontext", "context", "generated", "package"},
	},
	{
		ID:          "command_surface_discovery",
		Name:        "Command-surface discovery",
		Description: "Find CLI, MCP, resource, Web/API, package, and runtime command owners.",
		Keywords:    []string{"command", "cli", "mcp", "resource", "resources", "prompt", "prompts", "web", "runtime", "surface", "help"},
	},
	{
		ID:          "cross_repo_discovery",
		Name:        "Cross-repo discovery",
		Description: "Find group and cross-repo query, contracts, sync, status, and multi-repo surfaces.",
		Keywords:    []string{"cross", "repo", "repository", "repositories", "group", "groups", "sync", "contracts", "status", "multi"},
	},
}

func QueryCapabilityLanes() []QueryCapabilityLane {
	out := make([]QueryCapabilityLane, len(queryCapabilityLaneDefinitions))
	copy(out, queryCapabilityLaneDefinitions)
	return out
}

func mcpTools() []toolDefinition {
	return []toolDefinition{
		{
			Name:        "list_repos",
			Description: "List all indexed repositories available to Anvien.",
			InputSchema: map[string]any{
				"type":       "object",
				"properties": map[string]any{},
				"required":   []string{},
			},
		},
		{
			Name:        "query",
			Description: "Query the code knowledge graph for execution flows related to a concept.",
			InputSchema: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"query":           map[string]any{"type": "string", "description": "Natural language or keyword search query"},
					"task_context":    map[string]any{"type": "string", "description": "What you are working on. Helps ranking where supported."},
					"goal":            map[string]any{"type": "string", "description": "What you want to find. Helps ranking where supported."},
					"limit":           map[string]any{"type": "number", "description": "Max processes to return", "default": 5},
					"max_symbols":     map[string]any{"type": "number", "description": "Max symbols per process", "default": 10},
					"target_type":     map[string]any{"type": "string", "description": "Optional explicit query view", "enum": []string{"files", "symbols", "flows", "api"}},
					"dispatch_mode":   map[string]any{"type": "string", "description": "parent smart dispatch or explicit child command"},
					"include_content": map[string]any{"type": "boolean", "description": "Include full symbol source code", "default": false},
					"repo":            map[string]any{"type": "string", "description": "Repository name or path"},
				},
				"required": []string{"query"},
			},
		},
		{
			Name:        "cypher",
			Description: "Execute a read-only Cypher query against the indexed graph.",
			InputSchema: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"query": map[string]any{"type": "string", "description": "Cypher query to execute"},
					"repo":  map[string]any{"type": "string", "description": "Repository name or path"},
				},
				"required": []string{"query"},
			},
		},
		{
			Name:        "context",
			Description: "360-degree view of a single code symbol with categorized refs and process participation.",
			InputSchema: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"name":            map[string]any{"type": "string", "description": "Symbol name"},
					"uid":             map[string]any{"type": "string", "description": "Direct symbol UID from prior tool results"},
					"file_path":       map[string]any{"type": "string", "description": "File path to disambiguate common names"},
					"kind":            map[string]any{"type": "string", "description": "Kind filter such as Function, Class, Method, Interface, or Constructor"},
					"target_type":     map[string]any{"type": "string", "description": "Target resolver mode", "enum": []string{"auto", "symbol", "file"}, "default": "symbol"},
					"dispatch_mode":   map[string]any{"type": "string", "description": "parent smart dispatch or explicit child command"},
					"include_content": map[string]any{"type": "boolean", "description": "Include symbol source content", "default": false},
					"repo":            map[string]any{"type": "string", "description": "Repository name or path"},
				},
				"required": []string{},
			},
		},
		{
			Name:        "detect_changes",
			Description: "Analyze uncommitted git changes and find affected execution flows.",
			InputSchema: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"scope":    map[string]any{"type": "string", "description": "unstaged, staged, all, or compare", "enum": []string{"unstaged", "staged", "all", "compare"}, "default": "unstaged"},
					"base_ref": map[string]any{"type": "string", "description": "Branch or commit for compare scope"},
					"target_type": map[string]any{
						"type":        "string",
						"description": "Optional explicit changed target view",
						"enum":        []string{"files", "symbols", "flows"},
					},
					"dispatch_mode": map[string]any{"type": "string", "description": "parent smart dispatch or explicit child command"},
					"repo":          map[string]any{"type": "string", "description": "Repository name or path"},
				},
				"required": []string{},
			},
		},
		{
			Name:        "rename",
			Description: "Multi-file coordinated rename using the knowledge graph.",
			InputSchema: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"symbol_name": map[string]any{"type": "string", "description": "Current symbol name to rename"},
					"symbol_uid":  map[string]any{"type": "string", "description": "Direct symbol UID from prior tool results"},
					"new_name":    map[string]any{"type": "string", "description": "The new name for the symbol"},
					"file_path":   map[string]any{"type": "string", "description": "File path to disambiguate common names"},
					"dry_run":     map[string]any{"type": "boolean", "description": "Preview edits without modifying files", "default": true},
					"repo":        map[string]any{"type": "string", "description": "Repository name or path"},
				},
				"required": []string{"new_name"},
			},
		},
		{
			Name:        "impact",
			Description: "Analyze the blast radius of changing a code symbol.",
			InputSchema: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"target":        map[string]any{"type": "string", "description": "Name of function, class, or file to analyze"},
					"target_uid":    map[string]any{"type": "string", "description": "Direct symbol UID from prior tool results"},
					"target_type":   map[string]any{"type": "string", "description": "Target resolver mode", "enum": []string{"auto", "symbol", "file", "route", "tool"}, "default": "symbol"},
					"dispatch_mode": map[string]any{"type": "string", "description": "parent smart dispatch or explicit child command"},
					"direction":     map[string]any{"type": "string", "description": "upstream or downstream", "enum": []string{"upstream", "downstream"}, "default": "upstream"},
					"file_path":     map[string]any{"type": "string", "description": "File path hint to disambiguate common names"},
					"kind":          map[string]any{"type": "string", "description": "Kind filter such as Function, Class, Method, Interface, or Constructor"},
					"maxDepth":      map[string]any{"type": "number", "description": "Max relationship depth", "default": 3},
					"relationTypes": map[string]any{"type": "array", "items": map[string]any{"type": "string", "enum": sortedImpactAllowedRelationTypes()}},
					"includeTests":  map[string]any{"type": "boolean", "description": "Include test files", "default": false},
					"minConfidence": map[string]any{"type": "number", "description": "Minimum confidence 0-1", "default": 0},
					"repo":          map[string]any{"type": "string", "description": "Repository name or path"},
				},
				"required": []string{"direction"},
				"oneOf":    []map[string]any{{"required": []string{"target"}}, {"required": []string{"target_uid"}}},
			},
		},
		{
			Name:        "route_map",
			Description: "Show API route mappings: which components/hooks fetch which API endpoints, and which handler files serve them.",
			InputSchema: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"route": map[string]any{"type": "string", "description": "Filter by route path"},
					"repo":  map[string]any{"type": "string", "description": "Repository name or path"},
				},
				"required": []string{},
			},
		},
		{
			Name:        "tool_map",
			Description: "Show MCP/RPC tool definitions: which tools are defined, where they're handled, and their descriptions.",
			InputSchema: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"tool": map[string]any{"type": "string", "description": "Filter by tool name"},
					"repo": map[string]any{"type": "string", "description": "Repository name or path"},
				},
				"required": []string{},
			},
		},
		{
			Name:        "shape_check",
			Description: "Check response shapes for API routes against their consumers' property accesses.",
			InputSchema: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"route": map[string]any{"type": "string", "description": "Check a specific route"},
					"repo":  map[string]any{"type": "string", "description": "Repository name or path"},
				},
				"required": []string{},
			},
		},
		{
			Name:        "api_impact",
			Description: "Pre-change impact report for an API route handler.",
			InputSchema: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"route": map[string]any{"type": "string", "description": "Route path"},
					"file":  map[string]any{"type": "string", "description": "Handler file path"},
					"repo":  map[string]any{"type": "string", "description": "Repository name or path"},
				},
				"required": []string{},
			},
		},
		{
			Name:        "group_list",
			Description: "List all configured repository groups, or return details for one group.",
			InputSchema: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"name": map[string]any{"type": "string", "description": "Group name"},
				},
				"required": []string{},
			},
		},
		{
			Name:        "group_sync",
			Description: "Extract group contracts from indexed graph snapshots, exact-match cross-links, and write contracts.json.",
			InputSchema: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"name":           map[string]any{"type": "string", "description": "Group name"},
					"allowStale":     map[string]any{"type": "boolean", "description": "Allow stale member indexes", "default": false},
					"verbose":        map[string]any{"type": "boolean", "description": "Include verbose sync behavior where supported", "default": false},
					"exactOnly":      map[string]any{"type": "boolean", "description": "Use exact matching only", "default": false},
					"skipEmbeddings": map[string]any{"type": "boolean", "description": "Skip embedding fallback", "default": false},
				},
				"required": []string{"name"},
			},
		},
		{
			Name:        "group_contracts",
			Description: "Inspect a group's Contract Registry with type, repo, and unmatched filters.",
			InputSchema: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"name":          map[string]any{"type": "string", "description": "Group name"},
					"type":          map[string]any{"type": "string", "description": "Contract type filter such as http, grpc, topic, lib, or custom"},
					"repo":          map[string]any{"type": "string", "description": "Group repo path filter"},
					"unmatchedOnly": map[string]any{"type": "boolean", "description": "Only return contracts without matching cross-links", "default": false},
				},
				"required": []string{"name"},
			},
		},
		{
			Name:        "group_query",
			Description: "Search execution flows across all indexed repos in a group and merge ranked results.",
			InputSchema: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"name":     map[string]any{"type": "string", "description": "Group name"},
					"query":    map[string]any{"type": "string", "description": "Natural language or keyword search query"},
					"subgroup": map[string]any{"type": "string", "description": "Optional group repo path prefix"},
					"limit":    map[string]any{"type": "number", "description": "Max merged results", "default": 5},
				},
				"required": []string{"name", "query"},
			},
		},
		{
			Name:        "group_status",
			Description: "Report index staleness and Contract Registry staleness for each repo in a group.",
			InputSchema: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"name": map[string]any{"type": "string", "description": "Group name"},
				},
				"required": []string{"name"},
			},
		},
	}
}

func (s Server) queryTool(args map[string]any) (map[string]any, error) {
	query := strings.TrimSpace(stringArg(args, "query", ""))
	if query == "" {
		return nil, errors.New(`Missing "query" argument`)
	}
	g, err := s.graphForResource(stringArg(args, "repo", ""))
	if err != nil {
		return nil, err
	}
	limit := intArg(args, "limit", 5, 1, 50)
	targetType := normalizedTargetType(args, "", queryTargetTypeAllowed())
	dispatchMode := normalizedDispatchMode(args, targetType)
	if targetType != "" {
		return queryTargetPayload(g, query, targetType, dispatchMode, limit), nil
	}
	processSteps := resourceProcessStepsByProcess(g)
	matches := rankedProcessMatches(g, query, limit, processSteps)
	nodeByID := resourceGraphNodesByID(g)
	gapSummaries := queryResolutionGapSummaries(g, nodeByID)
	semanticStatus := semantic.GraphSemanticStatus(g)
	tokens := querySearchTokens(query)
	symbols := make([]map[string]any, 0)
	seenSymbolIDs := map[string]bool{}
	for processIndex, process := range matches {
		for stepIndex, step := range processSteps[process.ID] {
			if seenSymbolIDs[step.ID] {
				continue
			}
			seenSymbolIDs[step.ID] = true
			row := map[string]any{
				"id":          step.ID,
				"process":     process.Label,
				"processRank": processIndex + 1,
				"sourceRank":  stepIndex + 1,
				"step":        step.Step,
				"name":        step.Name,
				"filePath":    step.FilePath,
			}
			if node, ok := nodeByID[step.ID]; ok {
				addQueryNodeSemanticFields(row, node, gapSummaries[step.ID])
				addQueryMatchEvidence(row, node, step.Name, step.FilePath, tokens)
			}
			symbols = append(symbols, row)
		}
	}
	payload := map[string]any{
		"query":             query,
		"queryCapabilities": queryCapabilityEvidence(tokens),
		"semanticStatus":    semanticStatus,
		"processes":         matches,
		"process_symbols":   symbols,
		"definitions":       matchingDefinitionRows(g, query, limit, gapSummaries),
	}
	if boolArg(args, "explain", false) {
		payload["explain"] = map[string]any{
			"rankFields": []string{"processRank", "sourceRank", "rank", "score"},
			"evidenceFields": []string{
				"queryCapabilities",
				"queryLanes",
				"matchReasons",
			},
		}
	}
	if warning := querySemanticWarning(semanticStatus); warning != "" {
		payload["semanticWarning"] = warning
	}
	return payload, nil
}

func queryTargetPayload(g *graph.Graph, query string, targetType string, dispatchMode string, limit int) map[string]any {
	tokens := querySearchTokens(query)
	semanticStatus := semantic.GraphSemanticStatus(g)
	payload := map[string]any{
		"query":             query,
		"targetType":        targetType,
		"dispatchMode":      dispatchMode,
		"queryCapabilities": queryCapabilityEvidence(tokens),
		"semanticStatus":    semanticStatus,
	}
	switch targetType {
	case targetTypeFiles:
		files := queryFileRows(g, query, limit)
		payload["files"] = files
		payload["total"] = len(files)
	case targetTypeSymbols:
		gapSummaries := queryResolutionGapSummaries(g, resourceGraphNodesByID(g))
		symbols := querySymbolRows(g, query, limit, gapSummaries)
		payload["symbols"] = symbols
		payload["total"] = len(symbols)
	case targetTypeFlows:
		processSteps := resourceProcessStepsByProcess(g)
		flows := rankedProcessMatches(g, query, limit, processSteps)
		payload["flows"] = flows
		payload["processes"] = flows
		payload["total"] = len(flows)
	case targetTypeAPI:
		routes := mcpRouteMapItems(g, query)
		tools := mcpToolMapItems(g, query)
		if len(routes) > limit {
			routes = routes[:limit]
		}
		if len(tools) > limit {
			tools = tools[:limit]
		}
		payload["routes"] = routes
		payload["tools"] = tools
		payload["total"] = len(routes) + len(tools)
	}
	if warning := querySemanticWarning(semanticStatus); warning != "" {
		payload["semanticWarning"] = warning
	}
	return payload
}

func queryFileRows(g *graph.Graph, query string, limit int) []map[string]any {
	tokens := querySearchTokens(query)
	list := filecontext.NewBuilder(g).BuildFileList(filecontext.FileListOptions{Sort: "path", Limit: 0})
	type scoredFile struct {
		summary filecontext.FileSummary
		score   int
		symbols []map[string]any
	}
	byPath := make(map[string]*scoredFile, len(list.Files))
	for _, summary := range list.Files {
		score := queryTextScore(summary.Path, tokens)*6 +
			queryTextScore(summary.AppLayer, tokens)*3 +
			queryTextScore(summary.FunctionalArea, tokens)*3 +
			queryTextScore(summary.Kind, tokens)
		item := &scoredFile{summary: summary, score: score}
		byPath[normalizeContextPath(summary.Path)] = item
	}
	for _, node := range g.Nodes {
		path := normalizeContextPath(resourceNodeString(node, "filePath"))
		item := byPath[path]
		if item == nil || node.Label == scopeir.NodeFile {
			continue
		}
		name := firstResourceNodeString(node, "name", "label", "heuristicLabel")
		score := queryTextScore(name, tokens)*4 + queryTextScore(string(node.Label), tokens)
		if score <= 0 {
			continue
		}
		item.score += score
		if len(item.symbols) < 5 {
			item.symbols = append(item.symbols, map[string]any{
				"id":   node.ID,
				"name": name,
				"type": string(node.Label),
			})
		}
	}
	items := make([]scoredFile, 0, len(byPath))
	for _, item := range byPath {
		if item.score > 0 {
			items = append(items, *item)
		}
	}
	if len(items) == 0 {
		for _, summary := range list.Files[:minInt(len(list.Files), limit)] {
			items = append(items, scoredFile{summary: summary})
		}
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].score != items[j].score {
			return items[i].score > items[j].score
		}
		return items[i].summary.Path < items[j].summary.Path
	})
	out := make([]map[string]any, 0, minInt(len(items), limit))
	for index, item := range items[:minInt(len(items), limit)] {
		row := map[string]any{
			"rank":           index + 1,
			"score":          item.score,
			"summary":        item.summary,
			"path":           item.summary.Path,
			"matchedSymbols": item.symbols,
		}
		out = append(out, row)
	}
	return out
}

func querySymbolRows(g *graph.Graph, query string, limit int, gapSummaries map[string]queryResolutionGapSummary) []map[string]any {
	rows := matchingDefinitionRows(g, query, limit*3, gapSummaries)
	out := make([]map[string]any, 0, minInt(len(rows), limit))
	for _, row := range rows {
		label := scopeir.NodeLabel(fmt.Sprint(row["type"]))
		if !graphDefinitionLabelsForSymbols(label) {
			continue
		}
		if summary, ok := mcpFileSummaryForPath(g, fmt.Sprint(row["filePath"])); ok {
			row["fileSummary"] = summary
		}
		out = append(out, row)
		if len(out) >= limit {
			break
		}
	}
	return out
}

func (s Server) cypherTool(args map[string]any) (map[string]any, error) {
	query := strings.TrimSpace(stringArg(args, "query", ""))
	if query == "" {
		query = strings.TrimSpace(stringArg(args, "cypher", ""))
	}
	if query == "" {
		return nil, errors.New(`Missing "query" argument`)
	}
	if err := lbugruntime.ValidateReadQuery(query); err != nil {
		return nil, err
	}
	entry, err := s.resolveResourceRepo(stringArg(args, "repo", ""))
	if err != nil {
		return nil, err
	}
	rows, err := s.runCypherRead(entry, query)
	if err != nil {
		return nil, err
	}
	return map[string]any{
		"row_count": len(rows),
		"rows":      rows,
		"markdown":  markdownRows(rows),
	}, nil
}

func (s Server) runCypherRead(entry repo.RegistryEntry, query string) ([]map[string]any, error) {
	if s.openReadRunner != nil {
		runner, err := s.openReadRunner(filepath.Join(storagePathForEntry(entry), "lbug"))
		if err == nil {
			defer runner.Close()
			rows, err := runner.QueryRows(query)
			if err != nil {
				return nil, err
			}
			return mcpRowsFromLbugRows(rows), nil
		}
		if !errors.Is(err, lbugnative.ErrUnavailable) {
			return nil, err
		}
	}

	g, err := loadResourceGraphSnapshot(filepath.Join(storagePathForEntry(entry), "graph.json"))
	if err != nil {
		return nil, err
	}
	return runMCPGraphQuery(g, query)
}

func mcpRowsFromLbugRows(rows []lbugruntime.Row) []map[string]any {
	out := make([]map[string]any, 0, len(rows))
	for _, row := range rows {
		copied := make(map[string]any, len(row))
		for key, value := range row {
			copied[key] = value
		}
		out = append(out, copied)
	}
	return out
}

func rankedProcessMatches(g *graph.Graph, query string, limit int, processSteps map[string][]resourceProcessStep) []resourceProcess {
	tokens := querySearchTokens(query)
	processes := resourceProcessItems(g)
	nodeByID := resourceGraphNodesByID(g)
	type scored struct {
		process resourceProcess
		score   int
	}
	scoredItems := make([]scored, 0, len(processes))
	for _, process := range processes {
		score := queryTextScore(process.Label+" "+process.ProcessType, tokens)
		for _, step := range processSteps[process.ID] {
			score += queryTextScore(step.Name, tokens) * 3
			score += queryTextScore(step.FilePath, tokens) * 2
			if node, ok := nodeByID[step.ID]; ok {
				score += querySemanticSurfaceBoost(node, step.FilePath, tokens) / 2
			}
		}
		if score > 0 {
			scoredItems = append(scoredItems, scored{process: process, score: score})
		}
	}
	if len(scoredItems) == 0 {
		for _, process := range processes[:minInt(len(processes), limit)] {
			scoredItems = append(scoredItems, scored{process: process, score: 0})
		}
	}
	sort.Slice(scoredItems, func(i, j int) bool {
		if scoredItems[i].score != scoredItems[j].score {
			return scoredItems[i].score > scoredItems[j].score
		}
		if scoredItems[i].process.StepCount != scoredItems[j].process.StepCount {
			return scoredItems[i].process.StepCount > scoredItems[j].process.StepCount
		}
		return scoredItems[i].process.Label < scoredItems[j].process.Label
	})
	out := make([]resourceProcess, 0, minInt(len(scoredItems), limit))
	for _, item := range scoredItems[:minInt(len(scoredItems), limit)] {
		out = append(out, item.process)
	}
	return out
}

func matchingDefinitionRows(g *graph.Graph, query string, limit int, gapSummaries map[string]queryResolutionGapSummary) []map[string]any {
	tokens := querySearchTokens(query)
	type scored struct {
		node  graph.Node
		name  string
		path  string
		score int
	}
	items := make([]scored, 0)
	for _, node := range g.Nodes {
		if !queryDefinitionLabel(node.Label) {
			continue
		}
		name := firstResourceNodeString(node, "name", "label", "heuristicLabel")
		filePath := resourceNodeString(node, "filePath")
		if queryDefinitionSkip(node, filePath, tokens) {
			continue
		}
		score := queryTextScore(name, tokens)*5 +
			queryTextScore(node.ID, tokens)*4 +
			queryTextScore(filePath, tokens)*4 +
			queryTextScore(string(node.Label), tokens)*2 +
			queryTextScore(resourceNodeString(node, "appLayer"), tokens)*4 +
			queryTextScore(resourceNodeString(node, "functionalArea"), tokens)*4 +
			queryTextScore(resourceNodeString(node, "content"), tokens)
		if score == 0 {
			continue
		}
		score += queryDefinitionLabelPriority(node.Label)
		score += querySemanticSurfaceBoost(node, filePath, tokens)
		score -= queryDefinitionPenalty(node, filePath, tokens)
		if score <= 0 {
			continue
		}
		items = append(items, scored{node: node, name: name, path: filePath, score: score})
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].score != items[j].score {
			return items[i].score > items[j].score
		}
		if items[i].path != items[j].path {
			return items[i].path < items[j].path
		}
		if items[i].name != items[j].name {
			return items[i].name < items[j].name
		}
		return items[i].node.ID < items[j].node.ID
	})
	selected := make([]scored, 0, minInt(len(items), limit))
	selectedIDs := map[string]bool{}
	fileCounts := map[string]int{}
	for _, item := range items {
		key := item.path
		if key == "" {
			key = item.node.ID
		}
		if fileCounts[key] >= 3 {
			continue
		}
		selected = append(selected, item)
		selectedIDs[item.node.ID] = true
		fileCounts[key]++
		if len(selected) >= limit {
			break
		}
	}
	if len(selected) < limit {
		for _, item := range items {
			if selectedIDs[item.node.ID] {
				continue
			}
			selected = append(selected, item)
			selectedIDs[item.node.ID] = true
			if len(selected) >= limit {
				break
			}
		}
	}
	rows := make([]map[string]any, 0, len(selected))
	for index, item := range selected {
		row := map[string]any{
			"id":       item.node.ID,
			"name":     item.name,
			"type":     string(item.node.Label),
			"filePath": item.path,
			"rank":     index + 1,
			"score":    item.score,
		}
		addQueryNodeSemanticFields(row, item.node, gapSummaries[item.node.ID])
		addQueryMatchEvidence(row, item.node, item.name, item.path, tokens)
		rows = append(rows, row)
	}
	return rows
}

func queryDefinitionLabel(label scopeir.NodeLabel) bool {
	switch label {
	case scopeir.NodeFile,
		scopeir.NodeClass,
		scopeir.NodeFunction,
		scopeir.NodeMethod,
		scopeir.NodeInterface,
		scopeir.NodeStruct,
		scopeir.NodeType,
		scopeir.NodeTypeAlias,
		scopeir.NodeConstructor,
		scopeir.NodeEnum,
		scopeir.NodeRoute,
		scopeir.NodeTool:
		return true
	default:
		return false
	}
}

func queryDefinitionLabelPriority(label scopeir.NodeLabel) int {
	switch label {
	case scopeir.NodeFunction, scopeir.NodeMethod, scopeir.NodeConstructor:
		return 12
	case scopeir.NodeFile:
		return 10
	case scopeir.NodeRoute, scopeir.NodeTool:
		return 9
	case scopeir.NodeClass, scopeir.NodeStruct, scopeir.NodeInterface, scopeir.NodeType, scopeir.NodeTypeAlias, scopeir.NodeEnum:
		return 5
	default:
		return 0
	}
}

func queryDefinitionSkip(node graph.Node, filePath string, tokens []string) bool {
	appLayer := resourceNodeString(node, "appLayer")
	functionalArea := resourceNodeString(node, "functionalArea")
	normalizedPath := strings.ReplaceAll(strings.ToLower(filePath), "\\", "/")
	if !queryTokensContainAny(tokens, "test", "fixture", "benchmark", "e2e") {
		if strings.Contains(appLayer, "test") || strings.Contains(normalizedPath, "_test.") ||
			strings.Contains(normalizedPath, "/test/") || strings.Contains(normalizedPath, "/tests/") ||
			strings.Contains(normalizedPath, "/e2e/") {
			return true
		}
	}
	if !queryTokensContainAny(tokens, "doc", "docs", "documentation", "report", "plan", "evidence", "benchmark") {
		if appLayer == "docs" || functionalArea == "documentation" || functionalArea == "reporting" ||
			strings.HasPrefix(normalizedPath, "docs/") || strings.HasPrefix(normalizedPath, "reports/") {
			return true
		}
	}
	return false
}

func queryCapabilityEvidence(tokens []string) []map[string]any {
	out := make([]map[string]any, 0, len(queryCapabilityLaneDefinitions))
	for _, lane := range queryCapabilityLaneDefinitions {
		matched := queryMatchedLaneTokens(tokens, lane)
		if len(matched) == 0 && lane.ID != "concept_discovery" {
			continue
		}
		out = append(out, map[string]any{
			"id":            lane.ID,
			"name":          lane.Name,
			"description":   lane.Description,
			"matchedTokens": matched,
		})
	}
	return out
}

func queryMatchedLaneTokens(tokens []string, lane QueryCapabilityLane) []string {
	matched := make([]string, 0)
	seen := map[string]bool{}
	for _, token := range tokens {
		for _, keyword := range lane.Keywords {
			if token != keyword || seen[token] {
				continue
			}
			seen[token] = true
			matched = append(matched, token)
		}
	}
	return matched
}

func addQueryMatchEvidence(row map[string]any, node graph.Node, name string, filePath string, tokens []string) {
	lanes := queryNodeCapabilityLanes(node, name, filePath, tokens)
	if len(lanes) > 0 {
		row["queryLanes"] = lanes
	}
	reasons := queryMatchReasons(node, name, filePath, tokens)
	if len(reasons) > 0 {
		row["matchReasons"] = reasons
	}
}

func queryNodeCapabilityLanes(node graph.Node, name string, filePath string, tokens []string) []string {
	normalizedPath := strings.ReplaceAll(strings.ToLower(filePath), "\\", "/")
	appLayer := resourceNodeString(node, "appLayer")
	functionalArea := resourceNodeString(node, "functionalArea")
	out := make([]string, 0)
	for _, lane := range queryCapabilityLaneDefinitions {
		if len(queryMatchedLaneTokens(tokens, lane)) > 0 {
			out = append(out, lane.ID)
			continue
		}
		switch lane.ID {
		case "api_surface_discovery":
			if appLayer == "api" || appLayer == "api_contract" || functionalArea == "api" || functionalArea == "mcp" || strings.Contains(normalizedPath, "internal/mcp/") {
				out = append(out, lane.ID)
			}
		case "graph_quality_discovery":
			if functionalArea == "graph_health" || strings.Contains(normalizedPath, "query_health") || strings.Contains(normalizedPath, "resolution_inventory") || strings.Contains(normalizedPath, "source_site_accuracy") {
				out = append(out, lane.ID)
			}
		case "docs_setup_ai_context_discovery":
			if strings.Contains(normalizedPath, "internal/aicontext/") || strings.Contains(normalizedPath, "internal/cli/setup_command.go") || strings.Contains(normalizedPath, "internal/cli/analyze_postrun.go") {
				out = append(out, lane.ID)
			}
		case "command_surface_discovery":
			if functionalArea == "cli" || strings.Contains(normalizedPath, "internal/cli/") || strings.Contains(normalizedPath, "internal/mcp/resources.go") || strings.Contains(normalizedPath, "internal/mcp/prompts.go") {
				out = append(out, lane.ID)
			}
		case "cross_repo_discovery":
			if strings.Contains(normalizedPath, "internal/group/") || strings.Contains(normalizedPath, "group_command.go") || strings.Contains(normalizedPath, "group_tools.go") {
				out = append(out, lane.ID)
			}
		case "owner_discovery":
			if queryTextScore(name+" "+filePath, tokens) > 0 {
				out = append(out, lane.ID)
			}
		}
	}
	return uniqueQueryStrings(out)
}

func queryMatchReasons(node graph.Node, name string, filePath string, tokens []string) []string {
	reasons := make([]string, 0, 6)
	if queryTextScore(name, tokens) > 0 {
		reasons = append(reasons, "name")
	}
	if queryTextScore(node.ID, tokens) > 0 {
		reasons = append(reasons, "id")
	}
	if queryTextScore(filePath, tokens) > 0 {
		reasons = append(reasons, "filePath")
	}
	if queryTextScore(resourceNodeString(node, "appLayer"), tokens) > 0 {
		reasons = append(reasons, "appLayer")
	}
	if queryTextScore(resourceNodeString(node, "functionalArea"), tokens) > 0 {
		reasons = append(reasons, "functionalArea")
	}
	if querySemanticSurfaceBoost(node, filePath, tokens) > 0 {
		reasons = append(reasons, "semanticSurface")
	}
	return uniqueQueryStrings(reasons)
}

func uniqueQueryStrings(values []string) []string {
	if len(values) == 0 {
		return nil
	}
	out := make([]string, 0, len(values))
	seen := map[string]bool{}
	for _, value := range values {
		value = strings.TrimSpace(value)
		if value == "" || seen[value] {
			continue
		}
		seen[value] = true
		out = append(out, value)
	}
	return out
}

func querySemanticSurfaceBoost(node graph.Node, filePath string, tokens []string) int {
	appLayer := resourceNodeString(node, "appLayer")
	functionalArea := resourceNodeString(node, "functionalArea")
	name := firstResourceNodeString(node, "name", "label", "heuristicLabel")
	normalizedPath := strings.ReplaceAll(strings.ToLower(filePath), "\\", "/")
	boost := 0
	boost += queryPrimaryFileSymbolBoost(name, filePath)
	normalizedName := normalizeQuerySearchText(name)
	if queryTokensContainAny(tokens, "agent", "agents", "claude", "skill", "skills", "aicontext", "generated", "setup", "package") {
		if strings.Contains(normalizedPath, "internal/aicontext/") {
			boost += 180
		}
		if strings.Contains(normalizedPath, "internal/aicontext/skills/") {
			boost += 80
		}
		if strings.Contains(normalizedPath, "internal/cli/analyze_postrun.go") {
			boost += 140
		}
		if strings.Contains(normalizedPath, "internal/cli/setup_command.go") {
			boost += 130
		}
		if strings.Contains(normalizedPath, "internal/cli/package_command.go") ||
			strings.Contains(normalizedPath, "internal/cli/package_runtime.go") {
			boost += 80
		}
		if normalizedName == "generate ai context files" ||
			normalizedName == "install base skills" ||
			normalizedName == "base skill content" ||
			normalizedName == "setup install skills to" ||
			normalizedName == "setup install editor skills" ||
			normalizedName == "setup skill target name" {
			boost += 180
		}
	}
	if queryTokensContainAny(tokens, "prompt", "prompts", "resource", "resources", "setup", "mcp") {
		if strings.Contains(normalizedPath, "internal/mcp/prompts.go") {
			boost += 150
		}
		if strings.Contains(normalizedPath, "internal/mcp/resources.go") {
			boost += 120
		}
		if normalizedName == "prompt definitions" ||
			normalizedName == "generate map prompt" ||
			normalizedName == "setup resource" ||
			normalizedName == "mcp tools" {
			boost += 140
		}
	}
	if queryTokensContainAny(tokens, "unknown", "connectivity", "topology", "resolution", "health", "separation") {
		if functionalArea == "graph_health" {
			boost += 40
		}
		if strings.Contains(normalizedPath, "internal/graphhealth/compute.go") ||
			strings.Contains(normalizedPath, "internal/graphhealth/policy.go") {
			boost += 80
		}
		if strings.Contains(normalizedPath, "graph-health-filters") {
			boost += 45
		}
		if normalizeQuerySearchText(name) == "get node graph health" {
			boost += 80
		}
		if strings.Contains(normalizedPath, "internal/cli/query_health_command.go") ||
			strings.Contains(normalizedPath, "internal/cli/resolution_inventory_command.go") ||
			strings.Contains(normalizedPath, "internal/cli/source_site_accuracy_command.go") {
			boost += 170
		}
		if normalizedName == "new query health command" ||
			normalizedName == "new resolution inventory command" ||
			normalizedName == "new source site accuracy command" {
			boost += 180
		}
	}
	if queryTokensContainAny(tokens, "layout", "ring", "island", "optimizer", "manual", "visibility", "filter", "detail", "panel") {
		if appLayer == "frontend" {
			boost += 25
		}
		if functionalArea == "layout" {
			boost += 70
		}
		if functionalArea == "web_graph_ui" {
			boost += 35
		}
		if strings.Contains(normalizedPath, "anvien-web/src/lib/graph-adapter.ts") ||
			strings.Contains(normalizedPath, "anvien-web/src/hooks/usesigma.ts") ||
			strings.Contains(normalizedPath, "anvien-web/src/hooks/useSigma.ts") {
			boost += 90
		}
		if normalizeQuerySearchText(name) == "knowledge graph to graphology" {
			boost += 100
		}
	}
	if queryTokensContainAny(tokens, "frontend", "graph", "filter", "detail", "panel", "visibility") {
		if strings.Contains(normalizedPath, "anvien-web/src/lib/graph-health-filters.ts") ||
			strings.Contains(normalizedPath, "anvien-web/src/hooks/app-state/graph.tsx") ||
			strings.Contains(normalizedPath, "anvien-web/src/components/filetreepanel.tsx") ||
			strings.Contains(normalizedPath, "anvien-web/src/components/graphcanvas.tsx") {
			boost += 130
		}
	}
	if queryTokensContainAny(tokens, "query", "rank", "ranking", "match", "matching", "definition", "definitions", "command", "implementation") {
		if functionalArea == "mcp" {
			boost += 55
		}
		if functionalArea == "query" || functionalArea == "cli" {
			boost += 25
		}
		if strings.Contains(normalizedPath, "internal/mcp/tools.go") {
			boost += 90
		}
		if strings.Contains(normalizedPath, "internal/cli/tool_command.go") ||
			strings.Contains(normalizedPath, "cmd/anvien/main.go") {
			boost += 55
		}
		if strings.Contains(normalizedPath, "internal/cli/query_health_command.go") {
			boost += 120
		}
	}
	if queryTokensContainAny(tokens, "api", "contract", "generated", "http", "response") {
		if appLayer == "api" || appLayer == "api_contract" || appLayer == "generated_contract" {
			boost += 35
		}
		if functionalArea == "api" || functionalArea == "contracts" {
			boost += 35
		}
	}
	if queryTokensContainAny(tokens, "api", "route", "tool", "tools", "shape", "impact", "handler", "handlers") {
		if strings.Contains(normalizedPath, "internal/mcp/route_tool_map.go") ||
			strings.Contains(normalizedPath, "internal/mcp/route_shape_impact.go") {
			boost += 150
		}
		if normalizedName == "route map tool" ||
			normalizedName == "tool map tool" ||
			normalizedName == "shape check tool" ||
			normalizedName == "api impact tool" {
			boost += 180
		}
	}
	if queryTokensContainAny(tokens, "group", "groups", "cross", "repo", "repository", "contracts", "sync", "status") {
		if strings.Contains(normalizedPath, "internal/cli/group_command.go") ||
			strings.Contains(normalizedPath, "internal/mcp/group_tools.go") ||
			strings.Contains(normalizedPath, "internal/group/") {
			boost += 150
		}
		if normalizedName == "new group command" ||
			normalizedName == "group query tool" ||
			normalizedName == "query" ||
			normalizedName == "sync" {
			boost += 120
		}
	}
	return boost
}

func queryPrimaryFileSymbolBoost(name string, filePath string) int {
	if name == "" || filePath == "" {
		return 0
	}
	normalizedPath := strings.ReplaceAll(filePath, "\\", "/")
	slash := strings.LastIndex(normalizedPath, "/")
	base := normalizedPath
	if slash >= 0 {
		base = normalizedPath[slash+1:]
	}
	if dot := strings.LastIndex(base, "."); dot >= 0 {
		base = base[:dot]
	}
	if normalizeQuerySearchText(name) == normalizeQuerySearchText(base) {
		return 90
	}
	return 0
}

func queryDefinitionPenalty(node graph.Node, filePath string, tokens []string) int {
	appLayer := resourceNodeString(node, "appLayer")
	functionalArea := resourceNodeString(node, "functionalArea")
	normalizedPath := strings.ReplaceAll(strings.ToLower(filePath), "\\", "/")
	penalty := 0
	if !queryTokensContainAny(tokens, "test", "fixture", "benchmark") {
		if strings.Contains(appLayer, "test") || strings.Contains(normalizedPath, "_test.") ||
			strings.Contains(normalizedPath, "/test/") || strings.Contains(normalizedPath, "/tests/") ||
			strings.Contains(normalizedPath, "/e2e/") {
			penalty += 22
		}
	}
	if !queryTokensContainAny(tokens, "doc", "docs", "documentation", "report", "plan", "evidence", "benchmark") {
		if appLayer == "docs" || functionalArea == "documentation" || functionalArea == "reporting" ||
			strings.HasPrefix(normalizedPath, "docs/") || strings.HasPrefix(normalizedPath, "reports/") {
			penalty += 35
		}
	}
	if !queryTokensContainAny(tokens, "config", "configuration", "package") {
		if appLayer == "config" || functionalArea == "configuration" {
			penalty += 18
		}
	}
	return penalty
}

type queryResolutionGapSummary struct {
	Count           int
	Kinds           map[string]int
	Classifications map[string]int
	Actionability   map[string]int
	TopTargets      map[string]int
}

func queryResolutionGapSummaries(g *graph.Graph, nodeByID map[string]graph.Node) map[string]queryResolutionGapSummary {
	summaries := make(map[string]queryResolutionGapSummary)
	for _, relationship := range g.Relationships {
		if relationship.Type != graph.RelHasResolutionGap {
			continue
		}
		target, ok := nodeByID[relationship.TargetID]
		if !ok || target.Label != scopeir.NodeResolutionGap {
			continue
		}
		count := relationship.SourceSiteCount
		if count <= 0 {
			count = resourceNodeInt(target, "count")
		}
		if count <= 0 {
			count = 1
		}
		summary := summaries[relationship.SourceID]
		if summary.Kinds == nil {
			summary.Kinds = map[string]int{}
			summary.Classifications = map[string]int{}
			summary.Actionability = map[string]int{}
			summary.TopTargets = map[string]int{}
		}
		summary.Count += count
		incrementQueryCount(summary.Kinds, resourceNodeString(target, "gapKind"), count)
		incrementQueryCount(summary.Classifications, resourceNodeString(target, "classification"), count)
		incrementQueryCount(summary.Actionability, resourceNodeString(target, "actionability"), count)
		incrementQueryCount(summary.TopTargets, resourceNodeString(target, "targetText"), count)
		summaries[relationship.SourceID] = summary
	}
	return summaries
}

func addQueryNodeSemanticFields(row map[string]any, node graph.Node, gapSummary queryResolutionGapSummary) {
	row["type"] = string(node.Label)
	if value := resourceNodeString(node, "appLayer"); value != "" {
		row["appLayer"] = value
	}
	if value := resourceNodeString(node, "functionalArea"); value != "" {
		row["functionalArea"] = value
	}
	if value := resourceNodeString(node, "topologyStatus"); value != "" {
		row["topologyStatus"] = value
	}
	if value := resourceNodeString(node, "resolutionConfidence"); value != "" {
		row["resolutionConfidence"] = value
	}
	if count := resourceNodeInt(node, "resolutionGapCount"); count > 0 {
		row["resolutionGapCount"] = count
	}
	if gapSummary.Count > 0 {
		row["resolutionGapCount"] = gapSummary.Count
		row["resolutionGapKinds"] = cloneQueryCountMap(gapSummary.Kinds)
		row["resolutionGapClassifications"] = cloneQueryCountMap(gapSummary.Classifications)
		row["resolutionGapActionability"] = cloneQueryCountMap(gapSummary.Actionability)
		row["resolutionGapTopTargets"] = topQueryCountMap(gapSummary.TopTargets, 5)
	}
}

func querySemanticWarning(status semantic.GraphStatus) string {
	if status.AppLayer.Status == semantic.StatusStaleIncomplete || status.FunctionalArea.Status == semantic.StatusStaleIncomplete {
		return "Graph semantic metadata is incomplete; run anvien analyze --force to refresh graph evidence."
	}
	return ""
}

func incrementQueryCount(counts map[string]int, key string, count int) {
	if key == "" || count <= 0 {
		return
	}
	counts[key] += count
}

func cloneQueryCountMap(counts map[string]int) map[string]int {
	if len(counts) == 0 {
		return nil
	}
	out := make(map[string]int, len(counts))
	for key, count := range counts {
		if count > 0 {
			out[key] = count
		}
	}
	return out
}

func topQueryCountMap(counts map[string]int, limit int) map[string]int {
	if len(counts) == 0 || limit <= 0 {
		return nil
	}
	keys := make([]string, 0, len(counts))
	for key := range counts {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		if counts[keys[i]] != counts[keys[j]] {
			return counts[keys[i]] > counts[keys[j]]
		}
		return keys[i] < keys[j]
	})
	out := make(map[string]int, minInt(len(keys), limit))
	for _, key := range keys[:minInt(len(keys), limit)] {
		out[key] = counts[key]
	}
	return out
}

func querySearchTokens(query string) []string {
	normalized := normalizeQuerySearchText(query)
	if normalized == "" {
		return nil
	}
	parts := strings.Fields(normalized)
	out := make([]string, 0, len(parts))
	seen := map[string]bool{}
	for _, part := range parts {
		if len(part) < 2 {
			continue
		}
		for _, token := range queryTokenVariants(part) {
			if len(token) < 2 || seen[token] {
				continue
			}
			seen[token] = true
			out = append(out, token)
		}
	}
	return out
}

func queryTokenVariants(token string) []string {
	stem := queryTokenStem(token)
	if stem == token {
		return []string{token}
	}
	return []string{token, stem}
}

func queryTokensContainAny(tokens []string, expected ...string) bool {
	for _, token := range tokens {
		for _, item := range expected {
			if token == item {
				return true
			}
		}
	}
	return false
}

func queryTokenStem(token string) string {
	if len(token) > 5 && strings.HasSuffix(token, "ies") {
		return token[:len(token)-3] + "y"
	}
	if len(token) > 5 && strings.HasSuffix(token, "ing") {
		return token[:len(token)-3]
	}
	if len(token) > 4 && strings.HasSuffix(token, "ed") {
		return token[:len(token)-2]
	}
	if len(token) > 4 && strings.HasSuffix(token, "es") {
		return token[:len(token)-2]
	}
	if len(token) > 3 && strings.HasSuffix(token, "s") && !strings.HasSuffix(token, "ss") {
		return token[:len(token)-1]
	}
	return token
}

func queryTextScore(value string, tokens []string) int {
	if value == "" || len(tokens) == 0 {
		return 0
	}
	normalized := normalizeQuerySearchText(value)
	if normalized == "" {
		return 0
	}
	score := 0
	for _, token := range tokens {
		if normalized == token {
			score += 8
			continue
		}
		if strings.Contains(normalized, " "+token+" ") ||
			strings.HasPrefix(normalized, token+" ") ||
			strings.HasSuffix(normalized, " "+token) {
			score += 3
			continue
		}
		if strings.Contains(normalized, token) {
			score++
		}
	}
	return score
}

func normalizeQuerySearchText(value string) string {
	var builder strings.Builder
	previousAlphaNum := false
	for _, r := range value {
		if r >= 'A' && r <= 'Z' {
			if previousAlphaNum {
				builder.WriteByte(' ')
			}
			builder.WriteRune(r + ('a' - 'A'))
			previousAlphaNum = true
			continue
		}
		if r >= 'a' && r <= 'z' || r >= '0' && r <= '9' {
			builder.WriteRune(r)
			previousAlphaNum = true
			continue
		}
		if previousAlphaNum {
			builder.WriteByte(' ')
		}
		previousAlphaNum = false
	}
	return strings.Join(strings.Fields(builder.String()), " ")
}

func runMCPGraphQuery(g *graph.Graph, cypher string) ([]map[string]any, error) {
	if strings.Contains(cypher, "STEP_IN_PROCESS") {
		return mcpProcessStepRows(g, cypher)
	}
	if strings.Contains(cypher, "CALLS") && strings.Contains(cypher, "from.id IN") && strings.Contains(cypher, "to.id IN") {
		return mcpCallEdgeRows(g, cypher)
	}
	if matches := mcpNodeLabelQueryPattern.FindStringSubmatch(cypher); len(matches) == 2 {
		if strings.Contains(strings.ToUpper(cypher), " WHERE ") {
			return nil, errUnsupportedMCPGraphQuery
		}
		return mcpNodeLabelRows(g, scopeir.NodeLabel(matches[1]), mcpQueryLimit(cypher, 50))
	}
	return nil, errUnsupportedMCPGraphQuery
}

func mcpProcessStepRows(g *graph.Graph, cypher string) ([]map[string]any, error) {
	processIDs := extractMCPProcessIDs(cypher)
	if len(processIDs) == 0 {
		return nil, errors.New("process query did not include process ids")
	}
	processSet := stringSet(processIDs)
	nodeByID := resourceGraphNodesByID(g)
	rows := make([]map[string]any, 0)
	for _, relationship := range g.Relationships {
		if relationship.Type != graph.RelStepInProcess || !processSet[relationship.TargetID] {
			continue
		}
		node, ok := nodeByID[relationship.SourceID]
		if !ok {
			continue
		}
		step := 0
		if relationship.Step != nil {
			step = *relationship.Step
		}
		rows = append(rows, map[string]any{
			"id":         node.ID,
			"name":       node.Properties["name"],
			"filePath":   node.Properties["filePath"],
			"stepNumber": step,
			"step":       step,
		})
	}
	sort.Slice(rows, func(i, j int) bool {
		left, _ := rows[i]["stepNumber"].(int)
		right, _ := rows[j]["stepNumber"].(int)
		if left != right {
			return left < right
		}
		return rows[i]["id"].(string) < rows[j]["id"].(string)
	})
	return rows, nil
}

func mcpCallEdgeRows(g *graph.Graph, cypher string) ([]map[string]any, error) {
	fromIDs := extractMCPQuotedList(cypher, mcpFromIDInPattern)
	toIDs := extractMCPQuotedList(cypher, mcpToIDInPattern)
	if len(fromIDs) == 0 || len(toIDs) == 0 {
		return nil, errors.New("CALLS query did not include from/to id lists")
	}
	fromSet := stringSet(fromIDs)
	toSet := stringSet(toIDs)
	rows := make([]map[string]any, 0)
	for _, relationship := range g.Relationships {
		if relationship.Type != graph.RelCalls || !fromSet[relationship.SourceID] || !toSet[relationship.TargetID] {
			continue
		}
		rows = append(rows, map[string]any{"fromId": relationship.SourceID, "toId": relationship.TargetID, "type": string(relationship.Type)})
	}
	return rows, nil
}

func mcpNodeLabelRows(g *graph.Graph, label scopeir.NodeLabel, limit int) ([]map[string]any, error) {
	rows := make([]map[string]any, 0)
	for _, node := range g.Nodes {
		if node.Label != label {
			continue
		}
		rows = append(rows, map[string]any{"id": node.ID, "name": node.Properties["name"], "filePath": node.Properties["filePath"], "path": node.Properties["filePath"]})
		if len(rows) >= limit {
			break
		}
	}
	return rows, nil
}

func extractMCPProcessIDs(cypher string) []string {
	if matches := mcpProcessIDLiteralPattern.FindStringSubmatch(cypher); len(matches) == 2 {
		return []string{unescapeMCPCypherString(matches[1])}
	}
	return extractMCPQuotedList(cypher, mcpProcessIDInPattern)
}

func extractMCPQuotedList(cypher string, listPattern *regexp.Regexp) []string {
	matches := listPattern.FindStringSubmatch(cypher)
	if len(matches) != 2 {
		return nil
	}
	values := make([]string, 0)
	for _, match := range mcpQuotedStringPattern.FindAllStringSubmatch(matches[1], -1) {
		if len(match) == 2 {
			values = append(values, unescapeMCPCypherString(match[1]))
		}
	}
	return values
}

func unescapeMCPCypherString(value string) string {
	return strings.ReplaceAll(value, "''", "'")
}

func mcpQueryLimit(cypher string, fallback int) int {
	matches := mcpLimitPattern.FindStringSubmatch(cypher)
	if len(matches) != 2 {
		return fallback
	}
	parsed, err := strconv.Atoi(matches[1])
	if err != nil || parsed < 1 {
		return fallback
	}
	return minInt(parsed, 500)
}

func markdownRows(rows []map[string]any) string {
	if len(rows) == 0 {
		return "_No rows_"
	}
	keys := make([]string, 0, len(rows[0]))
	for key := range rows[0] {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	lines := []string{"| " + strings.Join(keys, " | ") + " |", "| " + strings.Repeat("--- | ", len(keys))}
	for _, row := range rows {
		values := make([]string, 0, len(keys))
		for _, key := range keys {
			values = append(values, strings.ReplaceAll(fmt.Sprint(row[key]), "|", "\\|"))
		}
		lines = append(lines, "| "+strings.Join(values, " | ")+" |")
	}
	return strings.Join(lines, "\n")
}

func containsScore(value string, needle string) int {
	if value == "" || needle == "" {
		return 0
	}
	if strings.Contains(strings.ToLower(value), needle) {
		return 1
	}
	return 0
}

func intArg(args map[string]any, name string, fallback int, minValue int, maxValue int) int {
	if args == nil {
		return fallback
	}
	var value int
	switch raw := args[name].(type) {
	case int:
		value = raw
	case float64:
		value = int(raw)
	default:
		return fallback
	}
	if value < minValue {
		return minValue
	}
	if value > maxValue {
		return maxValue
	}
	return value
}

func stringSet(values []string) map[string]bool {
	out := make(map[string]bool, len(values))
	for _, value := range values {
		out[value] = true
	}
	return out
}
