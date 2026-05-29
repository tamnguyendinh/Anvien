package aicontext

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/tamnguyendinh/avmatrix-go/internal/analyze"
)

const (
	startMarker = "<!-- anvien:start -->"
	endMarker   = "<!-- anvien:end -->"
)

var managedSectionPattern = regexp.MustCompile(`(?is)<!--\s*[a-z0-9-]+:start\s*-->.*?#\s+[^\n]*Code Intelligence.*?<!--\s*[a-z0-9-]+:end\s*-->`)

//go:embed skills/*.md
var baseSkillFiles embed.FS

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

type BaseSkillFile struct {
	Name        string
	Description string
	Task        string
	Content     string
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
	content := renderAnvienBlock(projectName, stats, options.NoStats)
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

	baseSkills, err := installBaseSkills(repoPath)
	if err != nil {
		return nil, nil, err
	}
	if len(baseSkills) > 0 {
		created = append(created, fmt.Sprintf(".claude/skills/anvien/ (%d skills)", len(baseSkills)))
	}
	if err := removeGeneratedSkills(repoPath); err != nil {
		return nil, nil, err
	}
	return created, baseSkills, nil
}

func renderAnvienBlock(projectName string, stats Stats, noStats bool) string {
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
	builder.WriteString("- **MUST refresh the graph before graph-based work.** Run `anvien analyze --force` before using any Anvien CLI command, MCP tool, MCP resource, Web/API view, or accuracy/benchmark command that reads, queries, validates, mutates, or reports on the semantic graph. This includes `query`, `context`, `impact`, `detect-changes`, `cypher`, `rename`, MCP `route_map`/`tool_map`/`shape_check`/`api_impact`, CLI `api route-map`/`api tool-map`/`api shape-check`/`api impact`, `augment`, `graph-health`, `query-health`, `resolution-inventory`, `source-site-accuracy`, and `benchmark-compare`.\n")
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
	writeCommandRow("Inspect one symbol deeply", "MCP `context` or CLI `anvien context \"<symbol>\" --repo <repo>`")
	writeCommandRow("Ask structured graph questions", "MCP `cypher` or CLI `anvien cypher \"<query>\" --repo <repo>`")
	writeCommandRow("Inspect API route handlers, consumers, and linked flows", "MCP `route_map` or CLI `anvien api route-map [route] --repo <repo>`")
	writeCommandRow("Inspect tool handlers and tool-related graph paths", "MCP `tool_map` or CLI `anvien api tool-map [tool] --repo <repo>`")
	writeCommandRow("Check API response shape drift against consumers", "MCP `shape_check` or CLI `anvien api shape-check [route] --repo <repo>`")
	writeCommandRow("Check route/API impact before changing handlers or contracts", "MCP `api_impact` or CLI `anvien api impact [route] --repo <repo>`")
	writeCommandRow("Add graph context to a text search pattern", "`anvien augment \"<pattern>\"`")
	writeCommandRow("Check blast radius before editing", "MCP `impact` or CLI `anvien impact \"<symbol>\" --repo <repo> --direction upstream`")
	writeCommandRow("Check changed symbols and affected flows before commit", "MCP `detect_changes` or CLI `anvien detect-changes --repo <repo> --scope all`")
	writeCommandRow("Rename a symbol safely", "MCP `rename` or CLI `anvien rename <symbol> <newName> --repo <repo>`")
	writeCommandRow("Audit topology health, diagnostics, and components", "`anvien graph-health summary --repo <repo> --json`")
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
	builder.WriteString("## Skills\n\n")
	builder.WriteString("| Task | Read this skill file |\n")
	builder.WriteString("|------|---------------------|\n")
	for _, skill := range baseSkills {
		fmt.Fprintf(&builder, "| %s | `.claude/skills/anvien/%s/SKILL.md` |\n", skill.Task, skill.Name)
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

type baseSkill struct {
	Name        string
	Description string
	Task        string
}

var baseSkills = []baseSkill{
	{Name: "anvien-exploring", Description: `Use when the user asks how code works, wants to understand architecture, trace execution flows, or explore unfamiliar parts of the codebase.`, Task: `Understand architecture, ownership, and execution flows`},
	{Name: "anvien-impact-analysis", Description: `Use when the user wants to know what will break if they change something, or needs safety analysis before editing code.`, Task: `Blast radius, HIGH/CRITICAL warnings, and changed-scope checks`},
	{Name: "anvien-debugging", Description: `Use when the user is debugging a bug, tracing an error, or asking why something fails.`, Task: `Trace bugs, failures, diagnostics, and graph-quality evidence`},
	{Name: "anvien-refactoring", Description: `Use when the user wants to rename, extract, split, move, or restructure code safely.`, Task: `Rename, extract, split, move, or restructure code safely`},
	{Name: "anvien-guide", Description: `Use when the user asks about Anvien itself, available tools, MCP resources, graph schema, prompts, or workflow reference.`, Task: `Unified CLI, MCP, resource, prompt, and Web/API reference`},
	{Name: "anvien-cli", Description: `Use when the user needs to run Anvien CLI commands for analysis, query, graph quality, API parity, groups, setup, runtime diagnostics, completion, package, wiki, hook, or version workflows.`, Task: `Terminal command guide for Anvien CLI surfaces`},
	{Name: "anvien-graph-quality", Description: `Use when the user needs graph-health, query-health, resolution inventory, source-site accuracy, or benchmark comparison evidence.`, Task: `Graph health, query health, resolution inventory, and accuracy audits`},
	{Name: "anvien-api-surface", Description: `Use when the user needs to inspect API routes, MCP tools, contract shape drift, generated Web contracts, handlers, consumers, or route/tool impact.`, Task: `API routes, MCP tools, shape checks, contracts, and consumers`},
	{Name: "anvien-cross-repo", Description: `Use when the user works across indexed repository groups, cross-repo query, contracts, status, sync, or multi-repo ownership.`, Task: `Repository groups, cross-repo query, contracts, status, and sync`},
	{Name: "anvien-runtime-packaging", Description: `Use when the user needs serve, mcp, setup, doctor diagnostics, launcher, package runtime, canonical executable, startup, or process lifecycle validation.`, Task: `Runtime, setup, launcher, package, and canonical executable workflows`},
	{Name: "anvien-ai-context", Description: `Use when the user changes generated AGENTS.md, CLAUDE.md, embedded Anvien skills, AI context generation, or source-vs-generated validation.`, Task: `Generated AGENTS.md, CLAUDE.md, embedded skills, and AI context validation`},
}

var retiredBaseSkillNames = []string{
	"anvien-pr-review",
	"av" + "matrix-pr-review",
	"av" + "matrix-exploring",
	"av" + "matrix-impact-analysis",
	"av" + "matrix-debugging",
	"av" + "matrix-refactoring",
	"av" + "matrix-guide",
	"av" + "matrix-cli",
	"av" + "matrix-graph-quality",
	"av" + "matrix-api-surface",
	"av" + "matrix-cross-repo",
	"av" + "matrix-runtime-packaging",
	"av" + "matrix-ai-context",
}

func BaseSkillFiles() ([]BaseSkillFile, error) {
	files := make([]BaseSkillFile, 0, len(baseSkills))
	for _, skill := range baseSkills {
		content, err := baseSkillContent(skill)
		if err != nil {
			return nil, err
		}
		files = append(files, BaseSkillFile{
			Name:        skill.Name,
			Description: skill.Description,
			Task:        skill.Task,
			Content:     content,
		})
	}
	return files, nil
}

func InstallBaseSkillsTo(targetDir string) ([]string, error) {
	files, err := BaseSkillFiles()
	if err != nil {
		return nil, err
	}
	if err := os.MkdirAll(targetDir, 0o755); err != nil {
		return nil, err
	}
	for _, retired := range retiredBaseSkillNames {
		if err := os.RemoveAll(filepath.Join(targetDir, retired)); err != nil {
			return nil, err
		}
	}
	installed := make([]string, 0, len(files))
	for _, file := range files {
		dir := filepath.Join(targetDir, file.Name)
		if err := os.RemoveAll(dir); err != nil {
			return installed, err
		}
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return installed, err
		}
		if err := os.WriteFile(filepath.Join(dir, "SKILL.md"), []byte(file.Content), 0o644); err != nil {
			return installed, err
		}
		installed = append(installed, file.Name)
	}
	return installed, nil
}

func installBaseSkills(repoPath string) ([]string, error) {
	skillsRoot := filepath.Join(repoPath, ".claude", "skills")
	skillsDir := filepath.Join(skillsRoot, "anvien")
	legacySkillsDir := filepath.Join(skillsRoot, "av"+"matrix")
	for _, dir := range []string{skillsDir, legacySkillsDir} {
		if err := os.RemoveAll(dir); err != nil {
			return nil, err
		}
	}
	return InstallBaseSkillsTo(skillsDir)
}

func baseSkillContent(skill baseSkill) (string, error) {
	content, err := baseSkillFiles.ReadFile("skills/" + skill.Name + ".md")
	if err != nil {
		return "", err
	}
	if strings.TrimSpace(string(content)) == "" {
		return fallbackBaseSkillContent(skill), nil
	}
	return string(content), nil
}

func fallbackBaseSkillContent(skill baseSkill) string {
	return fmt.Sprintf("---\nname: %s\ndescription: \"%s\"\n---\n\n# %s\n\n%s\n\nUse Anvien tools to accomplish this task.\n", skill.Name, skill.Description, skill.Name, skill.Description)
}

func removeGeneratedSkills(repoPath string) error {
	return os.RemoveAll(filepath.Join(repoPath, ".claude", "skills", "generated"))
}
