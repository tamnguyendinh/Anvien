package aicontext

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/tamnguyendinh/avmatrix-go/internal/analyze"
	"github.com/tamnguyendinh/avmatrix-go/internal/graph"
	"github.com/tamnguyendinh/avmatrix-go/internal/scopeir"
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

type GeneratedSkillInfo struct {
	Name        string
	Label       string
	SymbolCount int
	FileCount   int
}

type Result struct {
	Files        []string
	Skills       []GeneratedSkillInfo
	SkillsPath   string
	BaseSkillIDs []string
}

func Generate(repoPath string, projectName string, run analyze.Result, options Options) (Result, error) {
	skills, skillsPath, err := GenerateSkillFiles(repoPath, projectName, run.Graph)
	if err != nil {
		return Result{}, err
	}
	files, baseSkills, err := GenerateAIContextFiles(repoPath, projectName, statsFromRun(run), skills, options)
	if err != nil {
		return Result{}, err
	}
	return Result{
		Files:        files,
		Skills:       skills,
		SkillsPath:   skillsPath,
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

func GenerateAIContextFiles(repoPath string, projectName string, stats Stats, skills []GeneratedSkillInfo, options Options) ([]string, []string, error) {
	content := renderAVmatrixBlock(projectName, stats, skills, options.NoStats)
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
	return created, baseSkills, nil
}

func renderAVmatrixBlock(projectName string, stats Stats, skills []GeneratedSkillInfo, noStats bool) string {
	generatedRows := strings.Builder{}
	for _, skill := range skills {
		generatedRows.WriteString(fmt.Sprintf(
			"\n| Work in the %s area (%d symbols) | `.claude/skills/generated/%s/SKILL.md` |",
			skill.Label,
			skill.SymbolCount,
			skill.Name,
		))
	}

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
	builder.WriteString("- **MUST refresh the graph before graph-based work.** Run `avmatrix analyze --force` before using graph/query/impact/context/change-detection/cypher/accuracy commands.\n")
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
	writeCommandRow("Inspect API route handlers, consumers, and linked flows", "MCP `route_map`")
	writeCommandRow("Inspect tool handlers and tool-related graph paths", "MCP `tool_map`")
	writeCommandRow("Check API response shape drift against consumers", "MCP `shape_check`")
	writeCommandRow("Check route/API impact before changing handlers or contracts", "MCP `api_impact`")
	writeCommandRow("Add graph context to a text search pattern", "`avmatrix augment \"<pattern>\"`")
	writeCommandRow("Check blast radius before editing", "MCP `impact` or CLI `avmatrix impact \"<symbol>\" --repo <repo> --direction upstream`")
	writeCommandRow("Check changed symbols and affected flows before commit", "MCP `detect_changes` or CLI `avmatrix detect-changes --repo <repo> --scope all`")
	writeCommandRow("Rename a symbol safely", "MCP `rename` when available")
	writeCommandRow("Audit query retrieval quality", "`avmatrix query-health --repo <repo>`")
	writeCommandRow("Audit source-site and resolved-edge accuracy", "`avmatrix source-site-accuracy --graph .avmatrix/graph.json`")
	writeCommandRow("Inspect unresolved references / ResolutionGap inventory", "`avmatrix resolution-inventory --graph .avmatrix/graph.json`")
	writeCommandRow("Compare analyze benchmark outputs", "`avmatrix benchmark-compare <before> <after>`")
	writeCommandRow("Start local Web/API runtime", "`avmatrix serve --host 127.0.0.1 --port <port>`")
	writeCommandRow("Start the MCP server for agents", "`avmatrix mcp`")
	writeCommandRow("Configure editor/agent integrations", "`avmatrix setup`")
	writeCommandRow("Print AVmatrix version", "`avmatrix version`")
	writeCommandRow("Run Claude Code hook integration", "`avmatrix hook claude`")
	writeCommandRow("Create a repository group", "MCP `group_list` for discovery or CLI `avmatrix group create <name>`")
	writeCommandRow("Add an indexed repo to a group", "`avmatrix group add <group> <groupPath> <registryName>`")
	writeCommandRow("Remove repo from a group", "`avmatrix group remove <group> <path>`")
	writeCommandRow("List or inspect groups", "MCP `group_list` or CLI `avmatrix group list [name]`")
	writeCommandRow("Check staleness across group repos", "MCP `group_status` or CLI `avmatrix group status <name>`")
	writeCommandRow("Sync contract registry from indexed group repos", "MCP `group_sync` or CLI `avmatrix group sync <name>`")
	writeCommandRow("Search execution flows across group repos", "MCP `group_query` or CLI `avmatrix group query <name> \"<query>\"`")
	writeCommandRow("Inspect group contract registry", "MCP `group_contracts` or CLI `avmatrix group contracts <name>`")
	writeCommandRow("Build packaged runtime binary", "`avmatrix package build-runtime`")
	writeCommandRow("Prepare Go source for package fallback builds", "`avmatrix package prepare-go-source`")
	writeCommandRow("Verify packaged runtime for this platform", "`avmatrix package ensure-runtime`")
	writeCommandRow("Remove temporary package source output", "`avmatrix package clean-go-source`")
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
	builder.WriteString(generatedRows.String())
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

type communityInfo struct {
	ID          string
	Label       string
	Cohesion    float64
	SymbolCount int
	Members     []graph.Node
}

func GenerateSkillFiles(repoPath string, projectName string, g *graph.Graph) ([]GeneratedSkillInfo, string, error) {
	outputDir := filepath.Join(repoPath, ".claude", "skills", "generated")
	if g == nil {
		return nil, outputDir, nil
	}
	communities := significantCommunities(g)
	if len(communities) == 0 {
		return nil, outputDir, nil
	}
	if err := os.RemoveAll(outputDir); err != nil {
		return nil, outputDir, err
	}
	if err := os.MkdirAll(outputDir, 0o755); err != nil {
		return nil, outputDir, err
	}

	usedNames := make(map[string]struct{})
	skills := make([]GeneratedSkillInfo, 0, len(communities))
	for _, community := range communities {
		files := filesForMembers(repoPath, community.Members)
		name := uniqueKebab(community.Label, usedNames)
		dir := filepath.Join(outputDir, name)
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return nil, outputDir, err
		}
		if err := os.WriteFile(filepath.Join(dir, "SKILL.md"), []byte(renderSkill(projectName, name, community, files, processesForCommunity(g, community.ID))), 0o644); err != nil {
			return nil, outputDir, err
		}
		skills = append(skills, GeneratedSkillInfo{
			Name:        name,
			Label:       community.Label,
			SymbolCount: community.SymbolCount,
			FileCount:   len(files),
		})
	}
	return skills, outputDir, nil
}

func significantCommunities(g *graph.Graph) []communityInfo {
	membersByCommunity := make(map[string][]graph.Node)
	for _, relationship := range g.Relationships {
		if relationship.Type != graph.RelMemberOf {
			continue
		}
		member, ok := g.GetNode(relationship.SourceID)
		if ok {
			membersByCommunity[relationship.TargetID] = append(membersByCommunity[relationship.TargetID], member)
		}
	}
	communities := make([]communityInfo, 0)
	for _, node := range g.Nodes {
		if node.Label != scopeir.NodeCommunity {
			continue
		}
		members := membersByCommunity[node.ID]
		symbolCount := intProperty(node.Properties, "symbolCount")
		if symbolCount == 0 {
			symbolCount = len(members)
		}
		if symbolCount < 3 {
			continue
		}
		communities = append(communities, communityInfo{
			ID:          node.ID,
			Label:       firstNonEmpty(stringProperty(node.Properties, "heuristicLabel"), stringProperty(node.Properties, "label"), "Cluster"),
			Cohesion:    floatProperty(node.Properties, "cohesion"),
			SymbolCount: symbolCount,
			Members:     members,
		})
	}
	sort.Slice(communities, func(i, j int) bool {
		if communities[i].SymbolCount != communities[j].SymbolCount {
			return communities[i].SymbolCount > communities[j].SymbolCount
		}
		return communities[i].Label < communities[j].Label
	})
	if len(communities) > 20 {
		return communities[:20]
	}
	return communities
}

type fileInfo struct {
	Path    string
	Symbols []string
}

func filesForMembers(repoPath string, members []graph.Node) []fileInfo {
	byFile := make(map[string][]string)
	for _, member := range members {
		filePath := displayPath(repoPath, stringProperty(member.Properties, "filePath"))
		if filePath == "" {
			continue
		}
		byFile[filePath] = append(byFile[filePath], firstNonEmpty(stringProperty(member.Properties, "name"), member.ID))
	}
	files := make([]fileInfo, 0, len(byFile))
	for path, symbols := range byFile {
		sort.Strings(symbols)
		files = append(files, fileInfo{Path: path, Symbols: symbols})
	}
	sort.Slice(files, func(i, j int) bool {
		if len(files[i].Symbols) != len(files[j].Symbols) {
			return len(files[i].Symbols) > len(files[j].Symbols)
		}
		return files[i].Path < files[j].Path
	})
	return files
}

func processesForCommunity(g *graph.Graph, communityID string) []graph.Node {
	processes := make([]graph.Node, 0)
	for _, node := range g.Nodes {
		if node.Label != scopeir.NodeProcess {
			continue
		}
		if containsString(stringSliceProperty(node.Properties, "communities"), communityID) {
			processes = append(processes, node)
		}
	}
	sort.Slice(processes, func(i, j int) bool {
		left := intProperty(processes[i].Properties, "stepCount")
		right := intProperty(processes[j].Properties, "stepCount")
		if left != right {
			return left > right
		}
		return stringProperty(processes[i].Properties, "heuristicLabel") < stringProperty(processes[j].Properties, "heuristicLabel")
	})
	return processes
}

func renderSkill(projectName string, name string, community communityInfo, files []fileInfo, processes []graph.Node) string {
	lines := []string{
		"---",
		"name: " + name,
		fmt.Sprintf("description: \"Skill for the %s area of %s. %d symbols across %d files.\"", community.Label, projectName, community.SymbolCount, len(files)),
		"---",
		"",
		"# " + community.Label,
		"",
		fmt.Sprintf("%d symbols | %d files | Cohesion: %d%%", community.SymbolCount, len(files), int(community.Cohesion*100)),
		"",
		"## When to Use",
		"",
		fmt.Sprintf("- Modifying %s-related functionality", strings.ToLower(community.Label)),
	}
	if len(files) > 0 {
		lines = append(lines, "- Working with code in `"+filepath.ToSlash(filepath.Dir(files[0].Path))+"/`")
	}
	lines = append(lines, "", "## Key Files", "", "| File | Symbols |", "|------|---------|")
	for _, file := range limitFiles(files, 10) {
		symbols := strings.Join(limitStrings(file.Symbols, 5), ", ")
		if len(file.Symbols) > 5 {
			symbols += fmt.Sprintf(" (+%d)", len(file.Symbols)-5)
		}
		lines = append(lines, fmt.Sprintf("| `%s` | %s |", file.Path, symbols))
	}
	lines = append(lines, "", "## Key Symbols", "", "| Symbol | Type | File | Line |", "|--------|------|------|------|")
	for _, member := range limitNodes(community.Members, 20) {
		lines = append(lines, fmt.Sprintf(
			"| `%s` | %s | `%s` | %d |",
			firstNonEmpty(stringProperty(member.Properties, "name"), member.ID),
			member.Label,
			displayPath("", stringProperty(member.Properties, "filePath")),
			intProperty(member.Properties, "startLine"),
		))
	}
	if len(processes) > 0 {
		lines = append(lines, "", "## Execution Flows", "", "| Flow | Type | Steps |", "|------|------|-------|")
		for _, process := range limitNodes(processes, 10) {
			lines = append(lines, fmt.Sprintf(
				"| `%s` | %s | %d |",
				firstNonEmpty(stringProperty(process.Properties, "heuristicLabel"), stringProperty(process.Properties, "label"), process.ID),
				stringProperty(process.Properties, "processType"),
				intProperty(process.Properties, "stepCount"),
			))
		}
	}
	lines = append(lines, "", "## How to Explore", "", fmt.Sprintf("1. `query({query: \"%s\"})`", strings.ToLower(community.Label)), "2. Read key files listed above for implementation details", "")
	return strings.Join(lines, "\n")
}

func displayPath(repoPath string, filePath string) string {
	if filePath == "" {
		return ""
	}
	filePath = filepath.Clean(filePath)
	if repoPath != "" && filepath.IsAbs(filePath) {
		if rel, err := filepath.Rel(repoPath, filePath); err == nil {
			return filepath.ToSlash(rel)
		}
	}
	return filepath.ToSlash(filePath)
}

func uniqueKebab(label string, used map[string]struct{}) string {
	builder := strings.Builder{}
	lastDash := false
	for _, char := range strings.ToLower(label) {
		valid := (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9')
		if valid {
			builder.WriteRune(char)
			lastDash = false
			continue
		}
		if !lastDash {
			builder.WriteByte('-')
			lastDash = true
		}
	}
	base := strings.Trim(builder.String(), "-")
	if base == "" {
		base = "skill"
	}
	if len(base) > 50 {
		base = strings.TrimRight(base[:50], "-")
	}
	candidate := base
	for index := 2; ; index++ {
		if _, ok := used[candidate]; !ok {
			used[candidate] = struct{}{}
			return candidate
		}
		candidate = base + "-" + strconv.Itoa(index)
	}
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if value != "" {
			return value
		}
	}
	return ""
}

func stringProperty(properties graph.NodeProperties, key string) string {
	value, ok := properties[key]
	if !ok || value == nil {
		return ""
	}
	text, ok := value.(string)
	if !ok {
		return ""
	}
	return text
}

func intProperty(properties graph.NodeProperties, key string) int {
	value, ok := properties[key]
	if !ok || value == nil {
		return 0
	}
	switch typed := value.(type) {
	case int:
		return typed
	case int64:
		return int(typed)
	case float64:
		return int(typed)
	default:
		return 0
	}
}

func floatProperty(properties graph.NodeProperties, key string) float64 {
	value, ok := properties[key]
	if !ok || value == nil {
		return 0
	}
	switch typed := value.(type) {
	case float64:
		return typed
	case float32:
		return float64(typed)
	case int:
		return float64(typed)
	default:
		return 0
	}
}

func stringSliceProperty(properties graph.NodeProperties, key string) []string {
	value, ok := properties[key]
	if !ok || value == nil {
		return nil
	}
	switch typed := value.(type) {
	case []string:
		return typed
	case []any:
		out := make([]string, 0, len(typed))
		for _, item := range typed {
			if text, ok := item.(string); ok {
				out = append(out, text)
			}
		}
		return out
	default:
		return nil
	}
}

func containsString(values []string, target string) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}
	return false
}

func limitStrings(values []string, limit int) []string {
	if len(values) <= limit {
		return values
	}
	return values[:limit]
}

func limitNodes(values []graph.Node, limit int) []graph.Node {
	if len(values) <= limit {
		return values
	}
	return values[:limit]
}

func limitFiles(values []fileInfo, limit int) []fileInfo {
	if len(values) <= limit {
		return values
	}
	return values[:limit]
}
