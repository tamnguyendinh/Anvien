package aicontext

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/tamnguyendinh/anvien/internal/analyze"
)

const (
	startMarker = "<!-- anvien:start -->"
	endMarker   = "<!-- anvien:end -->"
)

var managedSectionPattern = regexp.MustCompile(`(?is)<!--\s*[a-z0-9-]+:start\s*-->.*?#\s+[^\n]*Code Intelligence.*?<!--\s*[a-z0-9-]+:end\s*-->`)

type Options struct {
	NoStats bool
}

type Stats struct {
	Files       int
	Nodes       int
	Edges       int
	Communities int
	Processes   int
}

type Result struct {
	Files        []string
	BaseSkillIDs []string
}

func Generate(repoPath string, projectName string, run analyze.Result, options Options) (Result, error) {
	files, baseSkills, err := GenerateAIContextFiles(repoPath, projectName, statsFromRun(run), options)
	if err != nil {
		return Result{}, err
	}
	return Result{
		Files:        files,
		BaseSkillIDs: baseSkills,
	}, nil
}

func statsFromRun(run analyze.Result) Stats {
	stats := Stats{
		Files:       run.Metrics.Files.Scanned,
		Communities: run.Metrics.Communities.CommunitiesEmitted,
		Processes:   run.Metrics.Processes.ProcessesEmitted,
	}
	if run.Graph != nil {
		stats.Nodes = len(run.Graph.Nodes)
		stats.Edges = len(run.Graph.Relationships)
	}
	return stats
}

func GenerateAIContextFiles(repoPath string, projectName string, stats Stats, options Options) ([]string, []string, error) {
	packages, err := SkillPackages()
	if err != nil {
		return nil, nil, err
	}
	content := renderAnvienBlock(projectName, stats, options.NoStats, packages)
	created := make([]string, 0, 3)

	agentsResult, err := upsertSection(filepath.Join(repoPath, "AGENTS.md"), content)
	if err != nil {
		return nil, nil, err
	}
	claudeResult, err := upsertSection(filepath.Join(repoPath, "CLAUDE.md"), content)
	if err != nil {
		return nil, nil, err
	}
	created = append(created, "AGENTS.md ("+agentsResult+")", "CLAUDE.md ("+claudeResult+")")

	installResult, err := installBaseSkills(repoPath, packages)
	if err != nil {
		return nil, nil, err
	}
	baseSkills := installResult.PackageIDs
	if len(baseSkills) > 0 {
		created = append(created, fmt.Sprintf(".claude/skills/anvien/ (%s)", installResult.Summary()))
	}
	if err := removeGeneratedSkills(repoPath); err != nil {
		return nil, nil, err
	}
	return created, baseSkills, nil
}

func renderAnvienBlock(projectName string, stats Stats, noStats bool, packages []SkillPackage) string {
	statsText := ""
	if !noStats {
		statsText = fmt.Sprintf(" (%d symbols, %d relationships, %d execution flows)", stats.Nodes, stats.Edges, stats.Processes)
	}

	var builder strings.Builder
	builder.WriteString(startMarker + "\n")
	builder.WriteString("# Anvien - Code Intelligence\n\n")
	fmt.Fprintf(&builder, "This project is indexed by Anvien as **%s**%s. Use Anvien to understand code, assess impact, navigate, audit graph quality, inspect resolution gaps, run query/accuracy benchmarks, and manage local/indexed repository intelligence.\n\n", projectName, statsText)
	builder.WriteString("Anvien is repo-agnostic: the same command surface works for any indexed repository. Use the indexed repository name shown above, or a repository name/path from `anvien list`, wherever examples refer to `<repo>`.\n\n")
	builder.WriteString("> If any Anvien command or MCP tool warns that the index is stale, run `anvien analyze --force` from the repository root first.\n\n")
	builder.WriteString("## Core Rule\n\n")
	builder.WriteString("Anvien has multiple command surfaces, but they are one tool:\n\n")
	builder.WriteString("- MCP tools are Anvien commands exposed to AI agents.\n")
	builder.WriteString("- CLI commands are Anvien commands exposed through terminal.\n")
	builder.WriteString("- Web/API commands are Anvien runtime commands exposed through the local server.\n\n")
	builder.WriteString("Use the surface that fits the current environment. Do not treat MCP tools and CLI commands as separate capabilities.\n\n")
	builder.WriteString("There is no single mandatory workflow beyond freshness, impact-before-edit, and change detection before commit. Anvien commands are selected by task. Use the full command set when it gives better evidence.\n\n")
	builder.WriteString("## Always Do\n\n")
	builder.WriteString("- **MUST refresh the graph before graph-based work.** Run `anvien analyze --force` before using any Anvien CLI command, MCP tool, MCP resource, Web/API view, or accuracy/benchmark command that reads, queries, validates, mutates, or reports on the semantic graph. This includes `query`, `context`, `impact`, `detect-changes`, `cypher`, `rename`, `file-context`, `file-hotspots`, MCP `route_map`/`tool_map`/`shape_check`/`api_impact`, CLI `api route-map`/`api tool-map`/`api shape-check`/`api impact`, `augment`, `graph-health`, `query-health`, `resolution-inventory`, `source-site-accuracy`, and `benchmark-compare`.\n")
	builder.WriteString("- **MUST run impact analysis before editing any function, class, method, exported symbol, API handler, graph builder, resolver, analyzer, or shared contract.**\n")
	builder.WriteString("- **MUST report blast radius.** HIGH or CRITICAL impact means warn clearly and proceed carefully; it is not an automatic ban on editing.\n")
	builder.WriteString("- **MUST run change detection before committing implementation work.** Use MCP `detect_changes` or CLI `anvien detect-changes --repo <repo> --scope all`.\n")
	builder.WriteString("- When exploring unfamiliar code, start with the Anvien command that matches the task instead of defaulting to grep.\n")
	builder.WriteString("- When a command has important flags, check `anvien <command> --help`.\n\n")
	builder.WriteString("## Never Do\n\n")
	builder.WriteString("- NEVER edit generated `AGENTS.md` / `CLAUDE.md` content as the permanent source of truth. Update the generator that writes these files, then regenerate.\n")
	builder.WriteString("- NEVER edit a function, class, or method without first running `impact` on it.\n")
	builder.WriteString("- NEVER ignore HIGH or CRITICAL impact warnings; explain them and keep the change scoped.\n")
	builder.WriteString("- NEVER rename symbols with find-and-replace; use graph-guided rename when available.\n")
	builder.WriteString("- NEVER commit implementation changes without running change detection.\n")
	builder.WriteString("- NEVER reduce graph evidence just to make output smaller; preserve counts, samples, meaning, and traceability.\n\n")
	builder.WriteString("## Command Selection Guide\n\n")
	builder.WriteString("Use Anvien by task, not by a fixed workflow. Pick the command surface that matches the job.\n\n")
	builder.WriteString("| When you need to... | Use |\n")
	builder.WriteString("|---------------------|-----|\n")
	writeCommandRow := func(when string, use string) {
		fmt.Fprintf(&builder, "| %s | %s |\n", when, use)
	}
	writeCommandRow("Refresh or rebuild graph evidence", "`anvien analyze --force`")
	writeCommandRow("Analyze a specific local repository path", "`anvien analyze <path> --force`")
	writeCommandRow("Analyze with embeddings enabled", "`anvien analyze --embeddings`")
	writeCommandRow("Record analyze performance/capacity metrics", "`anvien analyze --benchmark-json <file>`")
	writeCommandRow("Check whether the index is current", "`anvien status`")
	writeCommandRow("See indexed repositories and repo names", "MCP `list_repos` or CLI `anvien list`")
	writeCommandRow("Register an existing local index folder", "`anvien index [path...]`")
	writeCommandRow("Delete current repo index", "`anvien clean --force`")
	writeCommandRow("Delete all indexed repo data", "`anvien clean --all --force`")
	writeCommandRow("Find where a concept, behavior, or bug lives", "MCP `query` or CLI `anvien query \"<concept>\" --repo <repo>`")
	writeCommandRow("Find files relevant to a concept", "CLI `anvien query files \"<concept>\" --repo <repo>` or MCP `query` with `target_type=files`")
	writeCommandRow("Inspect one symbol deeply", "MCP `context` or CLI `anvien context symbol \"<symbol>\" --repo <repo>`")
	writeCommandRow("Inspect one file deeply as structured file-level data", "Prefer CLI `anvien file-context <path> --repo <repo> --json`; use `anvien context file <path> --repo <repo>` only when you want the context wrapper / human-oriented view")
	writeCommandRow("Ask structured graph questions", "MCP `cypher` or CLI `anvien cypher \"<query>\" --repo <repo>`")
	writeCommandRow("Inspect API route handlers, consumers, and linked flows", "MCP `route_map` or CLI `anvien api route-map [route] --repo <repo>`")
	writeCommandRow("Inspect tool handlers and tool-related graph paths", "MCP `tool_map` or CLI `anvien api tool-map [tool] --repo <repo>`")
	writeCommandRow("Check API response shape drift against consumers", "MCP `shape_check` or CLI `anvien api shape-check [route] --repo <repo>`")
	writeCommandRow("Check route/API impact before changing handlers or contracts", "MCP `api_impact` or CLI `anvien api impact [route] --repo <repo>`")
	writeCommandRow("Add graph context to a text search pattern", "`anvien augment \"<pattern>\"`")
	writeCommandRow("Check symbol blast radius before editing", "MCP `impact` or CLI `anvien impact symbol \"<symbol>\" --repo <repo> --direction upstream`")
	writeCommandRow("Check file-level blast radius before editing a file", "CLI `anvien impact file <path> --repo <repo> --direction upstream`")
	writeCommandRow("Check changed symbols, files, and affected flows before commit", "MCP `detect_changes` or CLI `anvien detect-changes --repo <repo> --scope all`; use `detect-changes files` for grouped file risk")
	writeCommandRow("Rename a symbol safely", "MCP `rename` or CLI `anvien rename <symbol> <newName> --repo <repo>`")
	writeCommandRow("Audit topology health, diagnostics, and components", "`anvien graph-health summary --repo <repo> --json`")
	writeCommandRow("Explain graph-health for one node", "`anvien graph-health explain \"File:<path>\" --repo <repo> --json`")
	writeCommandRow("Audit file-level graph health", "`anvien graph-health files --repo <repo> --json` or `anvien file-hotspots --repo <repo> --json`")
	writeCommandRow("Audit query retrieval quality", "`anvien query-health --repo <repo>`")
	writeCommandRow("Audit source-site and resolved-edge accuracy", "`anvien source-site-accuracy --graph .anvien/graph.json`")
	writeCommandRow("Inspect unresolved references / ResolutionGap inventory", "`anvien resolution-inventory --graph .anvien/graph.json`")
	writeCommandRow("Compare analyze benchmark outputs", "`anvien benchmark-compare <before> <after>`")
	writeCommandRow("Start local Web/API runtime", "`anvien serve --host 127.0.0.1 --port <port>`")
	writeCommandRow("Start the MCP server for agents", "`anvien mcp`")
	writeCommandRow("Configure editor/agent integrations", "`anvien setup`")
	writeCommandRow("Inspect local runtime locks and processes", "`anvien doctor locks --repo <repo> --json` or `anvien doctor processes --json`")
	writeCommandRow("Print Anvien version", "`anvien version`")
	writeCommandRow("Generate shell completion scripts", "`anvien completion <shell>`")
	writeCommandRow("Run Claude Code hook integration (hidden lifecycle helper)", "`anvien hook claude`")
	writeCommandRow("Create a repository group", "MCP `group_list` for discovery or CLI `anvien group create <name>`")
	writeCommandRow("Add an indexed repo to a group", "`anvien group add <group> <groupPath> <registryName>`")
	writeCommandRow("Remove repo from a group", "`anvien group remove <group> <path>`")
	writeCommandRow("List or inspect groups", "MCP `group_list` or CLI `anvien group list [name]`")
	writeCommandRow("Check staleness across group repos", "MCP `group_status` or CLI `anvien group status <name>`")
	writeCommandRow("Sync contract registry from indexed group repos", "MCP `group_sync` or CLI `anvien group sync <name>`")
	writeCommandRow("Search execution flows across group repos", "MCP `group_query` or CLI `anvien group query <name> \"<query>\"`")
	writeCommandRow("Inspect group contract registry", "MCP `group_contracts` or CLI `anvien group contracts <name>`")
	writeCommandRow("Build packaged runtime binary (hidden lifecycle helper)", "`anvien package build-runtime`")
	writeCommandRow("Prepare Go source for package fallback builds (hidden lifecycle helper)", "`anvien package prepare-go-source`")
	writeCommandRow("Verify packaged runtime for this platform (hidden lifecycle helper)", "`anvien package ensure-runtime`")
	writeCommandRow("Remove temporary package source output (hidden lifecycle helper)", "`anvien package clean-go-source`")
	writeCommandRow("Show wiki capability status", "`anvien wiki`")
	writeCommandRow("Show or set wiki capability mode", "`anvien wiki-mode [off|local]`")
	writeCommandRow("Check exact syntax and flags", "`anvien <command> --help`")
	builder.WriteString("\n")
	builder.WriteString("## Resources\n\n")
	builder.WriteString("| Resource | Use for |\n")
	builder.WriteString("|----------|---------|\n")
	builder.WriteString("| `anvien://repos` | All indexed repositories |\n")
	builder.WriteString("| `anvien://setup` | MCP setup and tool reference |\n")
	builder.WriteString("| `anvien://repo/<repo>/context` | Codebase overview and index freshness |\n")
	builder.WriteString("| `anvien://repo/<repo>/clusters` | Functional areas / communities |\n")
	builder.WriteString("| `anvien://repo/<repo>/processes` | Execution flows |\n")
	builder.WriteString("| `anvien://repo/<repo>/schema` | Graph schema for Cypher |\n")
	builder.WriteString("| `anvien://repo/<repo>/cluster/{name}` | Functional area details |\n")
	builder.WriteString("| `anvien://repo/<repo>/process/{name}` | Step-by-step execution trace |\n\n")
	builder.WriteString("## MCP Prompts\n\n")
	builder.WriteString("| Prompt | Use |\n")
	builder.WriteString("|--------|-----|\n")
	builder.WriteString("| `detect_impact` | Pre-commit impact workflow using `detect_changes`, `context`, and `impact`; HIGH/CRITICAL are blast-radius warnings, not edit bans. |\n")
	builder.WriteString("| `generate_map` | Evidence-backed architecture map workflow; resolves the repo through `anvien://repos` when needed and uses only resources/tools/command output actually read. |\n\n")
	builder.WriteString("MCP prompts are agent templates, not CLI commands. They guide tool/resource use and must still follow repository rules for freshness, impact-before-edit, and detect-changes before commit.\n\n")
	builder.WriteString("## Skill Selection Guide\n\n")
	builder.WriteString("Anvien installs every top-level package discovered under its embedded `internal/aicontext/skills/` catalog. A package may contain one or more `SKILL.md` entries plus scripts, references, assets, or templates; resolve those files relative to the package root.\n\n")
	builder.WriteString("| Package | Entries | Use |\n")
	builder.WriteString("|---------|---------|-----|\n")
	for _, pkg := range packages {
		fmt.Fprintf(&builder, "| `%s` | %s | %s |\n", pkg.Name, skillGuideEntries(pkg), skillGuideUse(pkg))
	}
	builder.WriteString("\n" + endMarker)
	return builder.String()
}

func FormatCrossRepoGroupsSection(groupNames []string) string {
	if len(groupNames) == 0 {
		return ""
	}
	groupToolNames := []string{"`group_list`", "`group_status`", "`group_sync`", "`group_contracts`", "`group_query`"}
	cliCommands := []string{
		"`anvien group list`",
		"`anvien group status <name>`",
		"`anvien group sync <name>`",
		"`anvien group contracts <name>`",
		"`anvien group query <name> \"<query>\"`",
	}
	return fmt.Sprintf(
		"## Cross-Repo Groups\n\nThis repository is listed under Anvien **group(s): %s**. For cross-repo work, use MCP tools %s. From the terminal, use %s.\n\n",
		strings.Join(groupNames, ", "),
		strings.Join(groupToolNames, ", "),
		strings.Join(cliCommands, ", "),
	)
}

func upsertSection(path string, content string) (string, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.WriteFile(path, []byte(content+"\n"), 0o644); err != nil {
				return "", err
			}
			return "created", nil
		}
		return "", err
	}
	text := string(raw)
	if strings.TrimSpace(text) == "" {
		if err := os.WriteFile(path, []byte(content+"\n"), 0o644); err != nil {
			return "", err
		}
		return "created", nil
	}
	start := strings.Index(text, startMarker)
	end := strings.Index(text, endMarker)
	if start >= 0 && end >= start {
		end += len(endMarker)
		next := strings.TrimSpace(text[:start]) + "\n\n" + content + "\n\n" + strings.TrimSpace(text[end:])
		if err := os.WriteFile(path, []byte(strings.TrimSpace(next)+"\n"), 0o644); err != nil {
			return "", err
		}
		return "updated", nil
	}
	if legacy := managedSectionPattern.FindStringIndex(text); legacy != nil {
		next := strings.TrimSpace(text[:legacy[0]]) + "\n\n" + content + "\n\n" + strings.TrimSpace(text[legacy[1]:])
		if err := os.WriteFile(path, []byte(strings.TrimSpace(next)+"\n"), 0o644); err != nil {
			return "", err
		}
		return "updated", nil
	}
	next := strings.TrimSpace(text) + "\n\n" + content + "\n"
	if err := os.WriteFile(path, []byte(next), 0o644); err != nil {
		return "", err
	}
	return "appended", nil
}

func installBaseSkills(repoPath string, packages []SkillPackage) (SkillInstallResult, error) {
	skillsRoot := filepath.Join(repoPath, ".claude", "skills")
	skillsDir := filepath.Join(skillsRoot, "anvien")
	legacySkillsDir := filepath.Join(skillsRoot, "av"+"matrix")
	if err := os.RemoveAll(legacySkillsDir); err != nil {
		return SkillInstallResult{}, err
	}
	return installSkillPackagesTo(skillsDir, packages)
}

func removeGeneratedSkills(repoPath string) error {
	return os.RemoveAll(filepath.Join(repoPath, ".claude", "skills", "generated"))
}
