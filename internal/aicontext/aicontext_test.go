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
	skills := []GeneratedSkillInfo{{Name: "core", Label: "Core", SymbolCount: 12, FileCount: 3}}

	files, installedBaseSkills, err := GenerateAIContextFiles(dir, "TestProject", stats, skills, Options{})
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
		"avmatrix://repo/TestProject/context",
		".claude/skills/generated/core/SKILL.md",
		"## Always Do",
		"## Never Do",
		"## MCP Tools",
		"## Resources",
		"avmatrix://repos",
		"avmatrix://setup",
		"avmatrix://repo/TestProject/schema",
		"avmatrix://repo/TestProject/cluster/{name}",
		"## Skills",
		"## CLI",
		"`avmatrix detect-changes --repo TestProject --scope all`",
		"avmatrix-impact-analysis/SKILL.md",
	} {
		if !strings.Contains(text, want) {
			t.Fatalf("AGENTS.md missing %q:\n%s", want, text)
		}
	}
	for _, retired := range []string{
		"## Tools Quick Reference",
		"## Impact Risk Levels",
		"## Self-Check Before Finishing",
		"## When Debugging",
		"## When Refactoring",
		"## Keeping the Index Fresh",
		"`avmatrix analyze --force --skip-agents-md`",
		"avmatrix_impact",
		"avmatrix_detect_changes",
		"avmatrix_query",
		"avmatrix_context",
	} {
		if strings.Contains(text, retired) {
			t.Fatalf("AGENTS.md contains retired content %q:\n%s", retired, text)
		}
	}

	if _, _, err := GenerateAIContextFiles(dir, "TestProject", Stats{Nodes: 10}, nil, Options{}); err != nil {
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
	if _, _, err := GenerateAIContextFiles(dir, "TestProject", Stats{Nodes: 50, Edges: 100, Processes: 5}, nil, Options{NoStats: true}); err != nil {
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

	if _, _, err := GenerateAIContextFiles(dir, "TestProject", Stats{Nodes: 1}, nil, Options{}); err != nil {
		t.Fatalf("GenerateAIContextFiles empty files: %v", err)
	}
	agents, err := os.ReadFile(filepath.Join(dir, "AGENTS.md"))
	if err != nil {
		t.Fatalf("read AGENTS.md: %v", err)
	}
	if strings.HasPrefix(string(agents), "\n") || !strings.Contains(string(agents), startMarker) {
		t.Fatalf("empty AGENTS.md was not replaced cleanly:\n%s", agents)
	}

	legacy := "# Manual\n\n<!-- gitnexus:start -->\n# GitNexus - Code Intelligence\nold\n<!-- gitnexus:end -->\n\n# Tail\n"
	if err := os.WriteFile(filepath.Join(dir, "CLAUDE.md"), []byte(legacy), 0o644); err != nil {
		t.Fatalf("write legacy CLAUDE.md: %v", err)
	}
	if _, _, err := GenerateAIContextFiles(dir, "TestProject", Stats{Nodes: 2}, nil, Options{}); err != nil {
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
	if strings.Contains(text, "gitnexus:start") || strings.Contains(text, "old") {
		t.Fatalf("legacy managed section was not replaced:\n%s", text)
	}
}

func TestGenerateAIContextFilesSkipAgentsPreservesManualFiles(t *testing.T) {
	dir := t.TempDir()
	agentsPath := filepath.Join(dir, "AGENTS.md")
	claudePath := filepath.Join(dir, "CLAUDE.md")
	agentsContent := "# AGENTS\n\nCustom manual instructions only\n"
	claudeContent := "# CLAUDE\n\nCustom manual instructions only\n"
	if err := os.WriteFile(agentsPath, []byte(agentsContent), 0o644); err != nil {
		t.Fatalf("write AGENTS.md: %v", err)
	}
	if err := os.WriteFile(claudePath, []byte(claudeContent), 0o644); err != nil {
		t.Fatalf("write CLAUDE.md: %v", err)
	}

	files, baseSkills, err := GenerateAIContextFiles(dir, "TestProject", Stats{Nodes: 42}, nil, Options{SkipAgentsMD: true})
	if err != nil {
		t.Fatalf("GenerateAIContextFiles: %v", err)
	}
	if len(baseSkills) == 0 {
		t.Fatalf("expected base skills to be installed")
	}
	joined := strings.Join(files, "\n")
	for _, want := range []string{"AGENTS.md (skipped via --skip-agents-md)", "CLAUDE.md (skipped via --skip-agents-md)"} {
		if !strings.Contains(joined, want) {
			t.Fatalf("generated file list missing %q: %v", want, files)
		}
	}
	if got, err := os.ReadFile(agentsPath); err != nil || string(got) != agentsContent {
		t.Fatalf("AGENTS.md changed: err=%v content=%q", err, got)
	}
	if got, err := os.ReadFile(claudePath); err != nil || string(got) != claudeContent {
		t.Fatalf("CLAUDE.md changed: err=%v content=%q", err, got)
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
