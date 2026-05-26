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
	startMarker = "<!-- avmatrix:start -->"
	endMarker   = "<!-- avmatrix:end -->"
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
	content := renderAVmatrixBlock(projectName, stats, options.NoStats)
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
		created = append(created, fmt.Sprintf(".claude/skills/avmatrix/ (%d skills)", len(baseSkills)))
	}
	if err := removeGeneratedSkills(repoPath); err != nil {
		return nil, nil, err
	}
	return created, baseSkills, nil
}

func renderAVmatrixBlock(projectName string, stats Stats, noStats bool) string {
	statsText := ""
	if !noStats {
		statsText = fmt.Sprintf(" (%d symbols, %d relationships, %d execution flows)", stats.Nodes, stats.Edges, stats.Processes)
	}

	var builder strings.Builder
	builder.WriteString(startMarker + "\n")
	builder.WriteString("# AVmatrix - Code Intelligence\n\n")
	fmt.Fprintf(&builder, "This project is indexed by AVmatrix as **%s**%s. Use AVmatrix to understand code, assess impact, navigate, audit graph quality, inspect resolution gaps, run query/accuracy benchmarks, and manage local/indexed repository intelligence.\n\n", projectName, statsText)
	builder.WriteString("AVmatrix is repo-agnostic: the same command surface works for any indexed repository. Use the indexed repository name shown above, or a repository name/path from `avmatrix list`, wherever examples refer to `<repo>`.\n\n")
	builder.WriteString("> If any AVmatrix command or MCP tool warns that the index is stale, run `avmatrix analyze --force` from the repository root first.\n\n")
	builder.WriteString("## Core Rule\n\n")
	builder.WriteString("AVmatrix has multiple command surfaces, but they are one tool:\n\n")
	builder.WriteString("- MCP tools are AVmatrix commands exposed to AI agents.\n")
	builder.WriteString("- CLI commands are AVmatrix commands exposed through terminal.\n")
	builder.WriteString("- Web/API commands are AVmatrix runtime commands exposed through the local server.\n\n")
	builder.WriteString("Use the surface that fits the current environment. Do not treat MCP tools and CLI commands as separate capabilities.\n\n")
	builder.WriteString("There is no single mandatory workflow beyond freshness, impact-before-edit, and change detection before commit. AVmatrix commands are selected by task. Use the full command set when it gives better evidence.\n\n")
	builder.WriteString("## Always Do\n\n")
	builder.WriteString("- **MUST refresh the graph before graph-based work.** Run `avmatrix analyze --force` before using any AVmatrix CLI command, MCP tool, MCP resource, Web/API view, or accuracy/benchmark command that reads, queries, validates, mutates, or reports on the semantic graph. This includes `query`, `context`, `impact`, `detect-changes`, `cypher`, `rename`, MCP `route_map`/`tool_map`/`shape_check`/`api_impact`, CLI `api route-map`/`api tool-map`/`api shape-check`/`api impact`, `augment`, `graph-health`, `query-health`, `resolution-inventory`, `source-site-accuracy`, and `benchmark-compare`.\n")
	builder.WriteString("- **MUST run impact analysis before editing any function, class, method, exported symbol, API handler, graph builder, resolver, analyzer, or shared contract.**\n")
	builder.WriteString("- **MUST report blast radius.** HIGH or CRITICAL impact means warn clearly and proceed carefully; it is not an automatic ban on editing.\n")
	builder.WriteString("- **MUST run change detection before committing implementation work.** Use MCP `detect_changes` or CLI `avmatrix detect-changes --repo <repo> --scope all`.\n")
	builder.WriteString("- When exploring unfamiliar code, start with the AVmatrix command that matches the task instead of defaulting to grep.\n")
	builder.WriteString("- When a command has important flags, check `avmatrix <command> --help`.\n\n")
	builder.WriteString("## Never Do\n\n")
	builder.WriteString("- NEVER edit generated `AGENTS.md` / `CLAUDE.md` content as the permanent source of truth. Update the generator that writes these files, then regenerate.\n")
	builder.WriteString("- NEVER edit a function, class, or method without first running `impact` on it.\n")
	builder.WriteString("- NEVER ignore HIGH or CRITICAL impact warnings; explain them and keep the change scoped.\n")
	builder.WriteString("- NEVER rename symbols with find-and-replace; use graph-guided rename when available.\n")
	builder.WriteString("- NEVER commit implementation changes without running change detection.\n")
	builder.WriteString("- NEVER reduce graph evidence just to make output smaller; preserve counts, samples, meaning, and traceability.\n\n")
	builder.WriteString("## Command Selection Guide\n\n")
	builder.WriteString("Use AVmatrix by task, not by a fixed workflow. Pick the command surface that matches the job.\n\n")
	builder.WriteString("| When you need to... | Use |\n")
	builder.WriteString("|---------------------|-----|\n")
	writeCommandRow := func(when string, use string) {
		fmt.Fprintf(&builder, "| %s | %s |\n", when, use)
	}
	writeCommandRow("Refresh or rebuild graph evidence", "`avmatrix analyze --force`")
	writeCommandRow("Analyze a specific local repository path", "`avmatrix analyze <path> --force`")
	writeCommandRow("Analyze with embeddings enabled", "`avmatrix analyze --embeddings`")
	writeCommandRow("Record analyze performance/capacity metrics", "`avmatrix analyze --benchmark-json <file>`")
	writeCommandRow("Check whether the index is current", "`avmatrix status`")
	writeCommandRow("See indexed repositories and repo names", "MCP `list_repos` or CLI `avmatrix list`")
	writeCommandRow("Register an existing local index folder", "`avmatrix index [path...]`")
	writeCommandRow("Delete current repo index", "`avmatrix clean --force`")
	writeCommandRow("Delete all indexed repo data", "`avmatrix clean --all --force`")
	writeCommandRow("Find where a concept, behavior, or bug lives", "MCP `query` or CLI `avmatrix query \"<concept>\" --repo <repo>`")
	writeCommandRow("Inspect one symbol deeply", "MCP `context` or CLI `avmatrix context \"<symbol>\" --repo <repo>`")
	writeCommandRow("Ask structured graph questions", "MCP `cypher` or CLI `avmatrix cypher \"<query>\" --repo <repo>`")
	writeCommandRow("Inspect API route handlers, consumers, and linked flows", "MCP `route_map` or CLI `avmatrix api route-map [route] --repo <repo>`")
	writeCommandRow("Inspect tool handlers and tool-related graph paths", "MCP `tool_map` or CLI `avmatrix api tool-map [tool] --repo <repo>`")
	writeCommandRow("Check API response shape drift against consumers", "MCP `shape_check` or CLI `avmatrix api shape-check [route] --repo <repo>`")
	writeCommandRow("Check route/API impact before changing handlers or contracts", "MCP `api_impact` or CLI `avmatrix api impact [route] --repo <repo>`")
	writeCommandRow("Add graph context to a text search pattern", "`avmatrix augment \"<pattern>\"`")
	writeCommandRow("Check blast radius before editing", "MCP `impact` or CLI `avmatrix impact \"<symbol>\" --repo <repo> --direction upstream`")
	writeCommandRow("Check changed symbols and affected flows before commit", "MCP `detect_changes` or CLI `avmatrix detect-changes --repo <repo> --scope all`")
	writeCommandRow("Rename a symbol safely", "MCP `rename` or CLI `avmatrix rename <symbol> <newName> --repo <repo>`")
	writeCommandRow("Audit topology health, diagnostics, and components", "`avmatrix graph-health summary --repo <repo> --json`")
	writeCommandRow("Audit query retrieval quality", "`avmatrix query-health --repo <repo>`")
	writeCommandRow("Audit source-site and resolved-edge accuracy", "`avmatrix source-site-accuracy --graph .avmatrix/graph.json`")
	writeCommandRow("Inspect unresolved references / ResolutionGap inventory", "`avmatrix resolution-inventory --graph .avmatrix/graph.json`")
	writeCommandRow("Compare analyze benchmark outputs", "`avmatrix benchmark-compare <before> <after>`")
	writeCommandRow("Start local Web/API runtime", "`avmatrix serve --host 127.0.0.1 --port <port>`")
	writeCommandRow("Start the MCP server for agents", "`avmatrix mcp`")
	writeCommandRow("Configure editor/agent integrations", "`avmatrix setup`")
	writeCommandRow("Print AVmatrix version", "`avmatrix version`")
	writeCommandRow("Run Claude Code hook integration (hidden lifecycle helper)", "`avmatrix hook claude`")
	writeCommandRow("Create a repository group", "MCP `group_list` for discovery or CLI `avmatrix group create <name>`")
	writeCommandRow("Add an indexed repo to a group", "`avmatrix group add <group> <groupPath> <registryName>`")
	writeCommandRow("Remove repo from a group", "`avmatrix group remove <group> <path>`")
	writeCommandRow("List or inspect groups", "MCP `group_list` or CLI `avmatrix group list [name]`")
	writeCommandRow("Check staleness across group repos", "MCP `group_status` or CLI `avmatrix group status <name>`")
	writeCommandRow("Sync contract registry from indexed group repos", "MCP `group_sync` or CLI `avmatrix group sync <name>`")
	writeCommandRow("Search execution flows across group repos", "MCP `group_query` or CLI `avmatrix group query <name> \"<query>\"`")
	writeCommandRow("Inspect group contract registry", "MCP `group_contracts` or CLI `avmatrix group contracts <name>`")
	writeCommandRow("Build packaged runtime binary (hidden lifecycle helper)", "`avmatrix package build-runtime`")
	writeCommandRow("Prepare Go source for package fallback builds (hidden lifecycle helper)", "`avmatrix package prepare-go-source`")
	writeCommandRow("Verify packaged runtime for this platform (hidden lifecycle helper)", "`avmatrix package ensure-runtime`")
	writeCommandRow("Remove temporary package source output (hidden lifecycle helper)", "`avmatrix package clean-go-source`")
	writeCommandRow("Show wiki capability status", "`avmatrix wiki`")
	writeCommandRow("Show or set wiki capability mode", "`avmatrix wiki-mode [off|local]`")
	writeCommandRow("Check exact syntax and flags", "`avmatrix <command> --help`")
	builder.WriteString("\n")
	builder.WriteString("## Resources\n\n")
	builder.WriteString("| Resource | Use for |\n")
	builder.WriteString("|----------|---------|\n")
	builder.WriteString("| `avmatrix://repos` | All indexed repositories |\n")
	builder.WriteString("| `avmatrix://setup` | MCP setup and tool reference |\n")
	builder.WriteString("| `avmatrix://repo/<repo>/context` | Codebase overview and index freshness |\n")
	builder.WriteString("| `avmatrix://repo/<repo>/clusters` | Functional areas / communities |\n")
	builder.WriteString("| `avmatrix://repo/<repo>/processes` | Execution flows |\n")
	builder.WriteString("| `avmatrix://repo/<repo>/schema` | Graph schema for Cypher |\n")
	builder.WriteString("| `avmatrix://repo/<repo>/cluster/{name}` | Functional area details |\n")
	builder.WriteString("| `avmatrix://repo/<repo>/process/{name}` | Step-by-step execution trace |\n\n")
	builder.WriteString("## Skills\n\n")
	builder.WriteString("| Task | Read this skill file |\n")
	builder.WriteString("|------|---------------------|\n")
	builder.WriteString("| Understand architecture / \"How does X work?\" | `.claude/skills/avmatrix/avmatrix-exploring/SKILL.md` |\n")
	builder.WriteString("| Blast radius / \"What breaks if I change X?\" | `.claude/skills/avmatrix/avmatrix-impact-analysis/SKILL.md` |\n")
	builder.WriteString("| Trace bugs / \"Why is X failing?\" | `.claude/skills/avmatrix/avmatrix-debugging/SKILL.md` |\n")
	builder.WriteString("| Rename / extract / split / refactor | `.claude/skills/avmatrix/avmatrix-refactoring/SKILL.md` |\n")
	builder.WriteString("| Tools, resources, schema reference | `.claude/skills/avmatrix/avmatrix-guide/SKILL.md` |\n")
	builder.WriteString("| Index, status, clean, and wiki capability CLI commands | `.claude/skills/avmatrix/avmatrix-cli/SKILL.md` |")
	builder.WriteString("\n" + endMarker)
	return builder.String()
}

func FormatCrossRepoGroupsSection(groupNames []string) string {
	if len(groupNames) == 0 {
		return ""
	}
	groupToolNames := []string{"`group_list`", "`group_status`", "`group_sync`", "`group_contracts`", "`group_query`"}
	cliCommands := []string{
		"`avmatrix group list`",
		"`avmatrix group status <name>`",
		"`avmatrix group sync <name>`",
		"`avmatrix group contracts <name>`",
		"`avmatrix group query <name> \"<query>\"`",
	}
	return fmt.Sprintf(
		"## Cross-Repo Groups\n\nThis repository is listed under AVmatrix **group(s): %s**. For cross-repo work, use MCP tools %s. From the terminal, use %s.\n\n",
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
}

var baseSkills = []baseSkill{
	{Name: "avmatrix-exploring", Description: `Use when the user asks how code works, wants to understand architecture, trace execution flows, or explore unfamiliar parts of the codebase.`},
	{Name: "avmatrix-debugging", Description: `Use when the user is debugging a bug, tracing an error, or asking why something fails.`},
	{Name: "avmatrix-impact-analysis", Description: `Use when the user wants to know what will break if they change something, or needs safety analysis before editing code.`},
	{Name: "avmatrix-refactoring", Description: `Use when the user wants to rename, extract, split, move, or restructure code safely.`},
	{Name: "avmatrix-guide", Description: `Use when the user asks about AVmatrix itself, available tools, MCP resources, graph schema, or workflow reference.`},
	{Name: "avmatrix-cli", Description: `Use when the user needs to run AVmatrix CLI commands like analyze/index a repo, check status, clean the index, inspect wiki capability mode, or list indexed repos.`},
}

func installBaseSkills(repoPath string) ([]string, error) {
	skillsDir := filepath.Join(repoPath, ".claude", "skills", "avmatrix")
	if err := os.RemoveAll(skillsDir); err != nil {
		return nil, err
	}
	installed := make([]string, 0, len(baseSkills))
	for _, skill := range baseSkills {
		dir := filepath.Join(skillsDir, skill.Name)
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return nil, err
		}
		content, err := baseSkillContent(skill)
		if err != nil {
			return nil, err
		}
		if err := os.WriteFile(filepath.Join(dir, "SKILL.md"), []byte(content), 0o644); err != nil {
			return nil, err
		}
		installed = append(installed, skill.Name)
	}
	return installed, nil
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
	return fmt.Sprintf("---\nname: %s\ndescription: \"%s\"\n---\n\n# %s\n\n%s\n\nUse AVmatrix tools to accomplish this task.\n", skill.Name, skill.Description, skill.Name, skill.Description)
}

func removeGeneratedSkills(repoPath string) error {
	return os.RemoveAll(filepath.Join(repoPath, ".claude", "skills", "generated"))
}
