package aicontext

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func expectedBaseSkillIDs() []string {
	return []string{
		"anvien-api-surface",
		"anvien-refactoring",
		"anvien-debugging",
		"anvien-planner",
	}
}

func staleGeneratedBaseSkillIDs() []string {
	return []string{
		"anvien-stale",
		"av" + "matrix-stale",
	}
}

func registeredBaseSkillIDs() []string {
	ids := make([]string, 0, len(baseSkills))
	for _, skill := range baseSkills {
		ids = append(ids, skill.Name)
	}
	return ids
}

func TestGenerateAIContextFilesCreatesAndUpdatesManagedContext(t *testing.T) {
	dir := t.TempDir()
	stats := Stats{Nodes: 50, Edges: 100, Processes: 5}
	oldLower := "av" + "matrix"
	oldDisplay := "AV" + "matrix"
	oldUpper := "AV" + "MATRIX"

	staleGenerated := filepath.Join(dir, ".claude", "skills", "generated", "legacy", "SKILL.md")
	if err := os.MkdirAll(filepath.Dir(staleGenerated), 0o755); err != nil {
		t.Fatalf("mkdir stale generated skill: %v", err)
	}
	if err := os.WriteFile(staleGenerated, []byte("# Legacy Generated Skill\n"), 0o644); err != nil {
		t.Fatalf("write stale generated skill: %v", err)
	}
	staleOldNamespace := filepath.Join(dir, ".claude", "skills", oldLower, oldLower+"-cli", "SKILL.md")
	if err := os.MkdirAll(filepath.Dir(staleOldNamespace), 0o755); err != nil {
		t.Fatalf("mkdir stale old skill namespace: %v", err)
	}
	if err := os.WriteFile(staleOldNamespace, []byte("# Old Skill\n"), 0o644); err != nil {
		t.Fatalf("write stale old skill namespace: %v", err)
	}
	for _, stale := range staleGeneratedBaseSkillIDs() {
		stalePath := filepath.Join(dir, ".claude", "skills", "anvien", stale, "SKILL.md")
		if err := os.MkdirAll(filepath.Dir(stalePath), 0o755); err != nil {
			t.Fatalf("mkdir stale base skill: %v", err)
		}
		if err := os.WriteFile(stalePath, []byte("---\nname: "+stale+"\ndescription: stale generated skill\n---\n# Stale\n"), 0o644); err != nil {
			t.Fatalf("write stale base skill: %v", err)
		}
	}

	files, installedBaseSkills, err := GenerateAIContextFiles(dir, "TestProject", stats, Options{})
	if err != nil {
		t.Fatalf("GenerateAIContextFiles: %v", err)
	}
	if len(files) == 0 {
		t.Fatalf("expected generated files")
	}
	if len(installedBaseSkills) == 0 {
		t.Fatalf("expected base skills to be installed")
	}
	if got, want := strings.Join(installedBaseSkills, ","), strings.Join(expectedBaseSkillIDs(), ","); got != want {
		t.Fatalf("installed base skill ids mismatch:\n got: %s\nwant: %s", got, want)
	}

	agentsPath := filepath.Join(dir, "AGENTS.md")
	content, err := os.ReadFile(agentsPath)
	if err != nil {
		t.Fatalf("read AGENTS.md: %v", err)
	}
	text := string(content)
	for _, want := range []string{
		startMarker,
		endMarker,
		"TestProject",
		"50 symbols, 100 relationships, 5 execution flows",
		"Anvien is repo-agnostic",
		"## Core Rule",
		"MCP tools are Anvien commands exposed to AI agents",
		"There is no single mandatory workflow",
		"## Always Do",
		"before using any Anvien CLI command, MCP tool, MCP resource, Web/API view",
		"MCP `route_map`/`tool_map`/`shape_check`/`api_impact`, CLI `api route-map`",
		"## Never Do",
		"## Command Selection Guide",
		"MCP `list_repos` or CLI `anvien list`",
		"MCP `query` or CLI `anvien query \"<concept>\" --repo <repo>`",
		"CLI `anvien query files \"<concept>\" --repo <repo>` or MCP `query` with `target_type=files`",
		"Prefer CLI `anvien file-context <path> --repo <repo> --json`; use `anvien context file <path> --repo <repo>` only when you want the context wrapper / human-oriented view",
		"MCP `context` or CLI `anvien context symbol \"<symbol>\" --repo <repo>`",
		"MCP `cypher` or CLI `anvien cypher \"<query>\" --repo <repo>`",
		"MCP `route_map` or CLI `anvien api route-map [route] --repo <repo>`",
		"MCP `tool_map` or CLI `anvien api tool-map [tool] --repo <repo>`",
		"MCP `shape_check` or CLI `anvien api shape-check [route] --repo <repo>`",
		"MCP `api_impact` or CLI `anvien api impact [route] --repo <repo>`",
		"MCP `impact` or CLI `anvien impact symbol \"<symbol>\" --repo <repo> --direction upstream`",
		"CLI `anvien impact file <path> --repo <repo> --direction upstream`",
		"MCP `detect_changes` or CLI `anvien detect-changes --repo <repo> --scope all`; use `detect-changes files`",
		"MCP `rename` or CLI `anvien rename <symbol> <newName> --repo <repo>`",
		"`anvien graph-health summary --repo <repo> --json`",
		"`anvien graph-health explain \"File:<path>\" --repo <repo> --json`",
		"`anvien graph-health files --repo <repo> --json` or `anvien file-hotspots --repo <repo> --json`",
		"`anvien query-health --repo <repo>`",
		"`anvien resolution-inventory --graph .anvien/graph.json`",
		"`anvien source-site-accuracy --graph .anvien/graph.json`",
		"`anvien doctor locks --repo <repo> --json`",
		"`anvien completion <shell>`",
		"MCP `group_query` or CLI `anvien group query <name> \"<query>\"`",
		"## Resources",
		"anvien://repos",
		"anvien://setup",
		"anvien://repo/<repo>/schema",
		"anvien://repo/<repo>/cluster/{name}",
		"## MCP Prompts",
		"`detect_impact`",
		"`generate_map`",
		"MCP prompts are agent templates, not CLI commands.",
		"## Skill Selection Guide",
		"Use Anvien workflow skills only for the retained domains below.",
		"choose concrete Anvien CLI/MCP commands from the Command Selection Guide",
		"| Inspect API routes, MCP tools, contracts, response shapes, or consumers | `.claude/skills/anvien/anvien-api-surface/SKILL.md` |",
		"| Rename, extract, split, move, or restructure code | `.claude/skills/anvien/anvien-refactoring/SKILL.md` |",
		"| Debug bugs, failures, diagnostics, or failure traces | `.claude/skills/anvien/anvien-debugging/SKILL.md` |",
		"| Create or review `docs/plans` plan, evidence, benchmark, or checklist mini-plan work | `.claude/skills/anvien/anvien-planner/SKILL.md` |",
		"anvien-api-surface/SKILL.md",
		"anvien-refactoring/SKILL.md",
		"anvien-debugging/SKILL.md",
		"anvien-planner/SKILL.md",
		"file-context",
		"file-hotspots",
	} {
		if !strings.Contains(text, want) {
			t.Fatalf("AGENTS.md missing %q:\n%s", want, text)
		}
	}
	if strings.Contains(text, ".claude/skills/generated/") {
		t.Fatalf("AGENTS.md should not reference generated skills:\n%s", text)
	}
	if _, err := os.Stat(filepath.Join(dir, ".claude", "skills", "generated")); !os.IsNotExist(err) {
		t.Fatalf("generated skills directory should not be created: %v", err)
	}
	if _, err := os.Stat(filepath.Join(dir, ".claude", "skills", oldLower)); !os.IsNotExist(err) {
		t.Fatalf("old skill namespace should be removed: %v", err)
	}
	for _, stale := range staleGeneratedBaseSkillIDs() {
		if _, err := os.Stat(filepath.Join(dir, ".claude", "skills", "anvien", stale, "SKILL.md")); !os.IsNotExist(err) {
			t.Fatalf("stale generated base skill %s should not be installed: %v", stale, err)
		}
	}
	for _, forbidden := range []string{
		oldDisplay,
		oldLower,
		oldUpper,
		"." + oldLower,
		oldLower + "://",
		oldLower + "-",
	} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("AGENTS.md contains old generated name %q:\n%s", forbidden, text)
		}
	}
	for _, retired := range []string{
		"## Tools Quick Reference",
		"## Impact Risk Levels",
		"## Self-Check Before Finishing",
		"## When Debugging",
		"## When Refactoring",
		"## Keeping the Index Fresh",
		"## MCP Tools",
		"## CLI",
		"## Practical Workflow",
		"Use the Anvien MCP tools to understand code",
		"before using graph/query/impact/context/change-detection/cypher/accuracy commands",
		"anvien://repo/TestProject/context",
		"`anvien detect-changes --repo TestProject --scope all`",
		"anvien_impact",
		"anvien_detect_changes",
		"anvien_query",
		"anvien_context",
	} {
		if strings.Contains(text, retired) {
			t.Fatalf("AGENTS.md contains retired content %q:\n%s", retired, text)
		}
	}
	forbiddenFlag := "--skip-" + "agents-md"
	if strings.Contains(text, forbiddenFlag) {
		t.Fatalf("AGENTS.md contains forbidden context bypass flag:\n%s", text)
	}

	if _, _, err := GenerateAIContextFiles(dir, "TestProject", Stats{Nodes: 10}, Options{}); err != nil {
		t.Fatalf("second GenerateAIContextFiles: %v", err)
	}
	updated, err := os.ReadFile(agentsPath)
	if err != nil {
		t.Fatalf("read updated AGENTS.md: %v", err)
	}
	if count := strings.Count(string(updated), "anvien:start"); count != 1 {
		t.Fatalf("expected one managed section after update, got %d:\n%s", count, updated)
	}

	for _, skill := range baseSkills {
		path := filepath.Join(dir, ".claude", "skills", "anvien", skill.Name, "SKILL.md")
		info, err := os.Stat(path)
		if err != nil {
			t.Fatalf("base skill %s was not installed: %v", skill.Name, err)
		}
		if info.Size() < 1000 {
			t.Fatalf("base skill %s is too small to be rich content: %d bytes", skill.Name, info.Size())
		}
		raw, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("read base skill %s: %v", skill.Name, err)
		}
		skillText := string(raw)
		if strings.Contains(skillText, "Use Anvien tools to accomplish this task.") {
			t.Fatalf("base skill %s fell back to placeholder content:\n%s", skill.Name, skillText)
		}
		if !strings.Contains(skillText, "##") {
			t.Fatalf("base skill %s missing rich sections:\n%s", skill.Name, skillText)
		}
	}

	plannerSkill, err := os.ReadFile(filepath.Join(dir, ".claude", "skills", "anvien", "anvien-planner", "SKILL.md"))
	if err != nil {
		t.Fatalf("read generated planner skill: %v", err)
	}
	for _, want := range []string{
		"Standard Plan Set",
		"YYYY-MM-DD-<slug>",
		"Evidence Ledger",
		"Benchmark Ledger",
		"Every checklist item must be a complete mini-plan by itself",
		"Do not write generic checklist items",
		"what to do, in what order",
	} {
		if !strings.Contains(string(plannerSkill), want) {
			t.Fatalf("generated planner skill missing %q:\n%s", want, plannerSkill)
		}
	}

	apiSkill, err := os.ReadFile(filepath.Join(dir, ".claude", "skills", "anvien", "anvien-api-surface", "SKILL.md"))
	if err != nil {
		t.Fatalf("read generated API skill: %v", err)
	}
	for _, want := range []string{"handlerFile", "anvien context file", "target_type=file"} {
		if !strings.Contains(string(apiSkill), want) {
			t.Fatalf("generated API skill missing %q:\n%s", want, apiSkill)
		}
	}
}

func TestInstallBaseSkillsToRemovesInactiveGeneratedAnvienSkills(t *testing.T) {
	dir := t.TempDir()
	for _, stale := range staleGeneratedBaseSkillIDs() {
		path := filepath.Join(dir, stale, "SKILL.md")
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			t.Fatalf("mkdir stale generated skill: %v", err)
		}
		if err := os.WriteFile(path, []byte("---\nname: "+stale+"\ndescription: stale generated skill\n---\n# Stale\n"), 0o644); err != nil {
			t.Fatalf("write stale generated skill: %v", err)
		}
	}
	customPath := filepath.Join(dir, "custom-skill", "SKILL.md")
	if err := os.MkdirAll(filepath.Dir(customPath), 0o755); err != nil {
		t.Fatalf("mkdir custom skill: %v", err)
	}
	if err := os.WriteFile(customPath, []byte("---\nname: custom-skill\ndescription: user skill\n---\n# Custom\n"), 0o644); err != nil {
		t.Fatalf("write custom skill: %v", err)
	}

	if _, err := InstallBaseSkillsTo(dir); err != nil {
		t.Fatalf("InstallBaseSkillsTo: %v", err)
	}
	for _, stale := range staleGeneratedBaseSkillIDs() {
		if _, err := os.Stat(filepath.Join(dir, stale, "SKILL.md")); !os.IsNotExist(err) {
			t.Fatalf("stale generated skill %s should be removed: %v", stale, err)
		}
	}
	if _, err := os.Stat(customPath); err != nil {
		t.Fatalf("custom non-Anvien skill should be preserved: %v", err)
	}
}

func TestBaseSkillRegistryAndSourceFrontmatter(t *testing.T) {
	if got, want := strings.Join(registeredBaseSkillIDs(), ","), strings.Join(expectedBaseSkillIDs(), ","); got != want {
		t.Fatalf("base skill registry mismatch:\n got: %s\nwant: %s", got, want)
	}
	files, err := BaseSkillFiles()
	if err != nil {
		t.Fatalf("BaseSkillFiles: %v", err)
	}
	if len(files) != len(expectedBaseSkillIDs()) {
		t.Fatalf("BaseSkillFiles returned %d files, want %d", len(files), len(expectedBaseSkillIDs()))
	}
	for _, skill := range baseSkills {
		if strings.TrimSpace(skill.Task) == "" {
			t.Fatalf("base skill %s has empty task", skill.Name)
		}
		content, err := baseSkillContent(skill)
		if err != nil {
			t.Fatalf("base skill %s content: %v", skill.Name, err)
		}
		trimmed := strings.TrimSpace(content)
		if len(trimmed) < 1000 {
			t.Fatalf("base skill %s content is too small: %d bytes", skill.Name, len(trimmed))
		}
		if !strings.HasPrefix(trimmed, "---\n") {
			t.Fatalf("base skill %s missing frontmatter start:\n%s", skill.Name, content)
		}
		if !strings.Contains(content, "\nname: "+skill.Name+"\n") {
			t.Fatalf("base skill %s frontmatter name mismatch:\n%s", skill.Name, content)
		}
		if !strings.Contains(content, "\ndescription: ") {
			t.Fatalf("base skill %s missing description frontmatter:\n%s", skill.Name, content)
		}
		if strings.Contains(content, "Use Anvien tools to accomplish this task.") {
			t.Fatalf("base skill %s uses fallback placeholder content:\n%s", skill.Name, content)
		}
	}
	for i, file := range files {
		if file.Name != expectedBaseSkillIDs()[i] {
			t.Fatalf("BaseSkillFiles[%d].Name = %q, want %q", i, file.Name, expectedBaseSkillIDs()[i])
		}
		if strings.TrimSpace(file.Content) == "" || strings.TrimSpace(file.Task) == "" || strings.TrimSpace(file.Description) == "" {
			t.Fatalf("BaseSkillFiles[%d] missing metadata/content: %#v", i, file)
		}
	}
}

func TestSkillGuidanceProtectsExpandedCommandSurface(t *testing.T) {
	contentBySkill := make(map[string]string, len(baseSkills))
	var combined strings.Builder
	for _, skill := range baseSkills {
		content, err := baseSkillContent(skill)
		if err != nil {
			t.Fatalf("base skill %s content: %v", skill.Name, err)
		}
		contentBySkill[skill.Name] = content
		combined.WriteString(content)
		combined.WriteString("\n")
	}

	checks := map[string][]string{
		"anvien-api-surface": {
			"route_map",
			"anvien api route-map",
			"api_impact",
			"Do not invent CLI commands",
		},
		"anvien-refactoring": {
			"behavior-preserving",
			"impact",
			"detect-changes",
			"graph-guided dry run",
		},
		"anvien-debugging": {
			"Reproduce or capture the symptom",
			"graph-health",
			"resolution-inventory",
			"Query Reliability Rule",
		},
		"anvien-planner": {
			"Standard Plan Set",
			"YYYY-MM-DD-<slug>",
			"Evidence Ledger",
			"Benchmark Ledger",
			"Every checklist item must be a complete mini-plan by itself",
			"Do not write generic checklist items",
		},
	}
	for skillName, fragments := range checks {
		content := contentBySkill[skillName]
		for _, fragment := range fragments {
			if !strings.Contains(content, fragment) {
				t.Fatalf("%s missing guidance fragment %q:\n%s", skillName, fragment, content)
			}
		}
	}

	allGuidance := combined.String()
	for _, want := range []string{
		"not a command router",
		"generated Command Selection Guide",
		"`route_map`",
		"`anvien api route-map",
		"`anvien detect-changes --repo <repo> --scope all`",
		"docs/plans/YYYY-MM-DD-<slug>/",
	} {
		if !strings.Contains(allGuidance, want) {
			t.Fatalf("combined skill guidance missing %q", want)
		}
	}
	for _, forbidden := range []string{
		"anvien route_map",
		"anvien tool_map",
		"anvien shape_check",
		"anvien api_impact",
		"anvien query_health",
	} {
		if strings.Contains(allGuidance, forbidden) {
			t.Fatalf("combined skill guidance contains invented CLI spelling %q", forbidden)
		}
	}
}

func TestGenerateAIContextFilesNoStatsOmitsVolatileCounts(t *testing.T) {
	dir := t.TempDir()
	if _, _, err := GenerateAIContextFiles(dir, "TestProject", Stats{Nodes: 50, Edges: 100, Processes: 5}, Options{NoStats: true}); err != nil {
		t.Fatalf("GenerateAIContextFiles: %v", err)
	}
	content, err := os.ReadFile(filepath.Join(dir, "AGENTS.md"))
	if err != nil {
		t.Fatalf("read AGENTS.md: %v", err)
	}
	text := string(content)
	if strings.Contains(text, "50 symbols") || strings.Contains(text, "100 relationships") || strings.Contains(text, "5 execution flows") {
		t.Fatalf("AGENTS.md contains volatile stats despite --no-stats:\n%s", text)
	}
	if !strings.Contains(text, "This project is indexed by Anvien as **TestProject**.") {
		t.Fatalf("AGENTS.md missing no-stats project sentence:\n%s", text)
	}
}

func TestGenerateAIContextFilesReplacesEmptyAndLegacyManagedContext(t *testing.T) {
	dir := t.TempDir()
	oldLower := "av" + "matrix"
	oldDisplay := "AV" + "matrix"
	for _, name := range []string{"AGENTS.md", "CLAUDE.md"} {
		if err := os.WriteFile(filepath.Join(dir, name), nil, 0o644); err != nil {
			t.Fatalf("write empty %s: %v", name, err)
		}
	}

	if _, _, err := GenerateAIContextFiles(dir, "TestProject", Stats{Nodes: 1}, Options{}); err != nil {
		t.Fatalf("GenerateAIContextFiles empty files: %v", err)
	}
	agents, err := os.ReadFile(filepath.Join(dir, "AGENTS.md"))
	if err != nil {
		t.Fatalf("read AGENTS.md: %v", err)
	}
	if strings.HasPrefix(string(agents), "\n") || !strings.Contains(string(agents), startMarker) {
		t.Fatalf("empty AGENTS.md was not replaced cleanly:\n%s", agents)
	}

	legacy := "# Manual\n\n<!-- " + oldLower + ":start -->\n# " + oldDisplay + " - Code Intelligence\nold legacy body\n<!-- " + oldLower + ":end -->\n\n# Tail\n"
	if err := os.WriteFile(filepath.Join(dir, "CLAUDE.md"), []byte(legacy), 0o644); err != nil {
		t.Fatalf("write legacy CLAUDE.md: %v", err)
	}
	if _, _, err := GenerateAIContextFiles(dir, "TestProject", Stats{Nodes: 2}, Options{}); err != nil {
		t.Fatalf("GenerateAIContextFiles legacy files: %v", err)
	}
	claude, err := os.ReadFile(filepath.Join(dir, "CLAUDE.md"))
	if err != nil {
		t.Fatalf("read CLAUDE.md: %v", err)
	}
	text := string(claude)
	for _, want := range []string{"# Manual", startMarker, "# Anvien - Code Intelligence", "# Tail"} {
		if !strings.Contains(text, want) {
			t.Fatalf("legacy CLAUDE.md missing %q:\n%s", want, text)
		}
	}
	if strings.Contains(text, "\nold legacy body\n") {
		t.Fatalf("legacy managed section was not replaced:\n%s", text)
	}
}

func TestRenderCrossRepoGroupsSectionMentionsSupportedToolsAndCommands(t *testing.T) {
	section := FormatCrossRepoGroupsSection([]string{"CorePlatform"})
	for _, want := range []string{
		"CorePlatform",
		"`group_list`",
		"`group_status`",
		"`group_sync`",
		"`group_contracts`",
		"`group_query`",
		"`anvien group list`",
		"`anvien group status <name>`",
		"`anvien group sync <name>`",
		"`anvien group contracts <name>`",
		"`anvien group query <name> \"<query>\"`",
	} {
		if !strings.Contains(section, want) {
			t.Fatalf("cross-repo section missing %q:\n%s", want, section)
		}
	}
	if strings.Contains(section, "group_impact") || strings.Contains(section, "anvien group impact") {
		t.Fatalf("cross-repo section mentions retired group impact:\n%s", section)
	}
	if got := FormatCrossRepoGroupsSection(nil); got != "" {
		t.Fatalf("nil groups should render empty section, got %q", got)
	}
}
