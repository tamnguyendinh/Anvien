package mcp

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

type promptDefinition struct {
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Arguments   []promptArgument `json:"arguments"`
}

type promptArgument struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Required    bool   `json:"required"`
}

type promptMessage struct {
	Role    string        `json:"role"`
	Content promptContent `json:"content"`
}

type promptContent struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func promptDefinitions() []promptDefinition {
	return []promptDefinition{
		{
			Name:        "detect_impact",
			Description: "Analyze current changes before committing using Anvien change detection, context, impact, and repository freshness rules.",
			Arguments: []promptArgument{
				{Name: "scope", Description: "What to analyze: unstaged, staged, all, or compare", Required: false},
				{Name: "base_ref", Description: "Branch/commit for compare scope", Required: false},
			},
		},
		{
			Name:        "generate_map",
			Description: "Generate evidence-backed architecture documentation from Anvien resources and graph facts.",
			Arguments: []promptArgument{
				{Name: "repo", Description: "Repository name (omit if only one indexed)", Required: false},
			},
		},
	}
}

func getPrompt(raw json.RawMessage) (map[string]any, error) {
	var params struct {
		Name      string         `json:"name"`
		Arguments map[string]any `json:"arguments"`
	}
	if err := json.Unmarshal(raw, &params); err != nil {
		return nil, fmt.Errorf("Invalid prompt params: %w", err)
	}

	switch params.Name {
	case "detect_impact":
		return map[string]any{"messages": []promptMessage{detectImpactPrompt(params.Arguments)}}, nil
	case "generate_map":
		return map[string]any{"messages": []promptMessage{generateMapPrompt(params.Arguments)}}, nil
	default:
		return nil, fmt.Errorf("Unknown prompt: %s", params.Name)
	}
}

func detectImpactPrompt(args map[string]any) promptMessage {
	scope := stringArg(args, "scope", "all")
	baseRef := stringArg(args, "base_ref", "")
	payload := map[string]any{"scope": scope}
	if baseRef != "" {
		payload["base_ref"] = baseRef
	}
	raw, _ := json.Marshal(payload)

	return promptMessage{
		Role: "user",
		Content: promptContent{
			Type: "text",
			Text: "Analyze the impact of my current code changes before committing.\n\n" +
				"Follow these steps:\n" +
				"1. Obey the active repository rules first. Before graph-based work, verify the Anvien index is fresh; if it is stale, run `anvien analyze --force` from the repository root before trusting graph facts.\n" +
				"2. Run MCP `detect_changes(" + string(raw) + ")` to find changed symbols, affected processes, App Layers, Functional Areas, and resolution-health impact. If MCP is unavailable, use CLI `anvien detect-changes --scope " + scope + cliBaseRefSuffix(baseRef) + " --repo <repo>`.\n" +
				"3. For changed symbols in important or risky flows, run MCP `context({name: \"<symbol>\"})` or CLI `anvien context \"<symbol>\" --repo <repo>` and inspect direct callers/callees before drawing conclusions.\n" +
				"4. For high-risk items, run MCP `impact({target: \"<symbol>\", direction: \"upstream\"})` or CLI `anvien impact \"<symbol>\" --repo <repo> --direction upstream` for blast radius. HIGH or CRITICAL is a blast-radius warning to report and account for; it is not an edit ban by itself.\n" +
				"5. If implementation work follows, run impact before editing functions/classes/methods/shared contracts and run detect-changes again before committing.\n" +
				"6. Summarize changed files/symbols, affected flows, risk level, direct callers, uncertainty, and recommended validation.\n\n" +
				"Present the analysis as a clear risk report backed only by command output you actually read.",
		},
	}
}

func generateMapPrompt(args map[string]any) promptMessage {
	repo := strings.TrimSpace(stringArg(args, "repo", ""))
	repoURI := url.PathEscape(repo)
	text := "Generate architecture documentation for this codebase using Anvien knowledge graph evidence.\n\n"
	if repo == "" {
		text += "Repo selection:\n" +
			"1. Because no `repo` argument was supplied, first READ `anvien://repos`.\n" +
			"2. If exactly one repo is listed, use that repo name.\n" +
			"3. If multiple repos are listed, select a repo only when its indexed path clearly matches the current workspace. If no match is provable, stop and ask the user which repo to map.\n" +
			"4. After selecting the repo, URL-escape the exact repo name for resource URIs and use that escaped value wherever `<resolved-repo-uri>` appears below. Never continue with a placeholder resource URI.\n\n" +
			generateMapSteps("<resolved-repo-uri>", "the resolved repo name")
	} else {
		text += "Repo selection:\n" +
			"1. Use the supplied repo name exactly as `" + repo + "`.\n" +
			"2. Use URL-escaped resource URI segment `" + repoURI + "` for Anvien resources.\n\n" +
			generateMapSteps(repoURI, repo)
	}

	return promptMessage{
		Role: "user",
		Content: promptContent{
			Type: "text",
			Text: text,
		},
	}
}

func generateMapSteps(repoURI string, repoName string) string {
	return "Evidence workflow:\n" +
		"1. READ `" + canonicalResourceScheme + "://repo/" + repoURI + "/context` for stats, index metadata, and staleness hints for " + repoName + ".\n" +
		"2. If the context resource or active repository rules indicate stale graph data, run `anvien analyze --force` from the repository root, then reread the context resource. If you cannot refresh, state that limitation in the output.\n" +
		"3. READ `" + canonicalResourceScheme + "://repo/" + repoURI + "/clusters` for functional areas and `" + canonicalResourceScheme + "://repo/" + repoURI + "/processes` for execution flows.\n" +
		"4. Choose representative clusters by explicit user request first, then user-facing runtime/API/tool relevance, then largest symbol count/cohesion from the resource output. Record the reason each cluster was selected.\n" +
		"5. Choose up to 5 representative processes by explicit user request first, then user-facing runtime/API/tool flows, process type such as route/tool/API flow, higher step count, or graph centrality when such evidence is available. Record the reason each process was selected; do not claim they are the top processes without stating the selection rule.\n" +
		"6. For each selected process, URL-escape the exact process name and READ `" + canonicalResourceScheme + "://repo/" + repoURI + "/process/<process-uri>` for step-by-step traces.\n" +
		"7. Use `context`, `impact`, `anvien api route-map`, `anvien api tool-map`, `anvien api shape-check`, `graph-health`, `query-health`, `resolution-inventory`, or `cypher` only when the architecture question needs evidence beyond the resources already read.\n\n" +
		"Output rules:\n" +
		"- Write architecture claims only from resources, MCP tools, CLI commands, or Web/API output you actually read.\n" +
		"- Do not invent Mermaid nodes, edges, dependencies, layers, or ownership. Mermaid edges must come from process steps, route/tool/consumer links, imports/calls, or other graph evidence you can cite.\n" +
		"- Include uncertainty notes when clusters, processes, graph-health, query-health, resolution inventory, or unresolved-reference evidence is incomplete or was not read.\n" +
		"- Produce an ARCHITECTURE.md-ready draft with overview, selected functional areas, selected execution flows, Mermaid diagram, evidence notes, and limitations. Edit or create repository files only if the user explicitly asked for file changes.\n"
}

func cliBaseRefSuffix(baseRef string) string {
	if baseRef == "" {
		return ""
	}
	return " --base-ref " + baseRef
}

func stringArg(args map[string]any, name string, fallback string) string {
	if args == nil {
		return fallback
	}
	value, ok := args[name].(string)
	if !ok || value == "" {
		return fallback
	}
	return value
}
