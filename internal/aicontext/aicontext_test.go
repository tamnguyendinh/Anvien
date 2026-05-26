package aicontext

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGenerateAIContextFilesCreatesAndUpdatesManagedContext(t *testing.T) {
	dir := t.TempDir()
	stats := Stats{Nodes: 50, Edges: 100, Processes: 5}
	staleGenerated := filepath.Join(dir, ".claude", "skills", "generated", "legacy", "SKILL.md")
	if err := os.MkdirAll(filepath.Dir(staleGenerated), 0o755); err != nil {
		t.Fatalf("mkdir stale generated skill: %v", err)
	}
	if err := os.WriteFile(staleGenerated, []byte("# Legacy Generated Skill\n"), 0o644); err != nil {
		t.Fatalf("write stale generated skill: %v", err)
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
		"AVmatrix is repo-agnostic",
		"## Core Rule",
		"MCP tools are AVmatrix commands exposed to AI agents",
		"There is no single mandatory workflow",
		"## Always Do",
		"before using any AVmatrix CLI command, MCP tool, MCP resource, Web/API view",
		"MCP `route_map`/`tool_map`/`shape_check`/`api_impact`, CLI `api route-map`",
		"## Never Do",
		"## Command Selection Guide",
		"MCP `list_repos` or CLI `avmatrix list`",
		"MCP `query` or CLI `avmatrix query \"<concept>\" --repo <repo>`",
		"MCP `context` or CLI `avmatrix context \"<symbol>\" --repo <repo>`",
		"MCP `cypher` or CLI `avmatrix cypher \"<query>\" --repo <repo>`",
		"MCP `route_map` or CLI `avmatrix api route-map [route] --repo <repo>`",
		"MCP `tool_map` or CLI `avmatrix api tool-map [tool] --repo <repo>`",
		"MCP `shape_check` or CLI `avmatrix api shape-check [route] --repo <repo>`",
		"MCP `api_impact` or CLI `avmatrix api impact [route] --repo <repo>`",
		"MCP `impact` or CLI `avmatrix impact \"<symbol>\" --repo <repo> --direction upstream`",
		"MCP `detect_changes` or CLI `avmatrix detect-changes --repo <repo> --scope all`",
		"MCP `rename` or CLI `avmatrix rename <symbol> <newName> --repo <repo>`",
		"`avmatrix graph-health summary --repo <repo> --json`",
		"`avmatrix query-health --repo <repo>`",
		"`avmatrix resolution-inventory --graph .avmatrix/graph.json`",
		"`avmatrix source-site-accuracy --graph .avmatrix/graph.json`",
		"MCP `group_query` or CLI `avmatrix group query <name> \"<query>\"`",
		"## Resources",
		"avmatrix://repos",
		"avmatrix://setup",
		"avmatrix://repo/<repo>/schema",
		"avmatrix://repo/<repo>/cluster/{name}",
		"## Skills",
		"avmatrix-impact-analysis/SKILL.md",
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
		"Use the AVmatrix MCP tools to understand code",
		"before using graph/query/impact/context/change-detection/cypher/accuracy commands",
		"avmatrix://repo/TestProject/context",
		"`avmatrix detect-changes --repo TestProject --scope all`",
		"avmatrix_impact",
		"avmatrix_detect_changes",
		"avmatrix_query",
		"avmatrix_context",
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
	if count := strings.Count(string(updated), "avmatrix:start"); count != 1 {
		t.Fatalf("expected one managed section after update, got %d:\n%s", count, updated)
	}

	for _, skill := range baseSkills {
		path := filepath.Join(dir, ".claude", "skills", "avmatrix", skill.Name, "SKILL.md")
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
		if strings.Contains(skillText, "Use AVmatrix tools to accomplish this task.") {
			t.Fatalf("base skill %s fell back to placeholder content:\n%s", skill.Name, skillText)
		}
		if !strings.Contains(skillText, "##") {
			t.Fatalf("base skill %s missing rich sections:\n%s", skill.Name, skillText)
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
	if !strings.Contains(text, "This project is indexed by AVmatrix as **TestProject**.") {
		t.Fatalf("AGENTS.md missing no-stats project sentence:\n%s", text)
	}
}

func TestGenerateAIContextFilesReplacesEmptyAndLegacyManagedContext(t *testing.T) {
	dir := t.TempDir()
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

	legacy := "# Manual\n\n<!-- avmatrix:start -->\n# AVmatrix - Code Intelligence\nold legacy body\n<!-- avmatrix:end -->\n\n# Tail\n"
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
	for _, want := range []string{"# Manual", startMarker, "# AVmatrix - Code Intelligence", "# Tail"} {
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
		"`avmatrix group list`",
		"`avmatrix group status <name>`",
		"`avmatrix group sync <name>`",
		"`avmatrix group contracts <name>`",
		"`avmatrix group query <name> \"<query>\"`",
	} {
		if !strings.Contains(section, want) {
			t.Fatalf("cross-repo section missing %q:\n%s", want, section)
		}
	}
	if strings.Contains(section, "group_impact") || strings.Contains(section, "avmatrix group impact") {
		t.Fatalf("cross-repo section mentions retired group impact:\n%s", section)
	}
	if got := FormatCrossRepoGroupsSection(nil); got != "" {
		t.Fatalf("nil groups should render empty section, got %q", got)
	}
}
