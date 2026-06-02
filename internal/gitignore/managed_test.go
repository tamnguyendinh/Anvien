package gitignore

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	ignorematcher "github.com/tamnguyendinh/anvien/internal/ignore"
)

func TestEnsureCreatesManagedGitignore(t *testing.T) {
	dir := t.TempDir()

	result, err := Ensure(dir)
	if err != nil {
		t.Fatalf("Ensure: %v", err)
	}
	if !result.Changed {
		t.Fatalf("Ensure Changed = false, want true")
	}
	if filepath.Base(result.Path) != ".gitignore" {
		t.Fatalf("Ensure path = %q, want .gitignore", result.Path)
	}

	text := readGitignore(t, dir)
	for _, want := range []string{
		StartMarker,
		EndMarker,
		".anvien/",
		"AGENTS.md",
		"CLAUDE.md",
		".claude/",
		".codex/",
		".agents/",
		"anvien-launcher/web-dist/",
		"anvien-launcher/server-bundle/",
	} {
		if !strings.Contains(text, want) {
			t.Fatalf(".gitignore missing %q:\n%s", want, text)
		}
	}
}

func TestEnsurePreservesUserContentAndReplacesManagedBlock(t *testing.T) {
	dir := t.TempDir()
	initial := strings.Join([]string{
		"# user rules",
		"custom-output/",
		"",
		StartMarker,
		".anvien/",
		EndMarker,
		"",
		"# keep me",
		"vendor-cache/",
		"",
	}, "\n")
	if err := os.WriteFile(filepath.Join(dir, ".gitignore"), []byte(initial), 0o644); err != nil {
		t.Fatalf("write initial .gitignore: %v", err)
	}

	result, err := Ensure(dir)
	if err != nil {
		t.Fatalf("Ensure: %v", err)
	}
	if !result.Changed {
		t.Fatalf("Ensure Changed = false, want true")
	}

	text := readGitignore(t, dir)
	if strings.Count(text, StartMarker) != 1 || strings.Count(text, EndMarker) != 1 {
		t.Fatalf("managed markers should appear once:\n%s", text)
	}
	for _, want := range []string{"custom-output/", "vendor-cache/", "AGENTS.md", "CLAUDE.md"} {
		if !strings.Contains(text, want) {
			t.Fatalf(".gitignore missing %q:\n%s", want, text)
		}
	}

	second, err := Ensure(dir)
	if err != nil {
		t.Fatalf("second Ensure: %v", err)
	}
	if second.Changed {
		t.Fatalf("second Ensure Changed = true, want false")
	}
}

func TestManagedGitignoreRulesAreLoadedByMatcher(t *testing.T) {
	dir := t.TempDir()
	if _, err := Ensure(dir); err != nil {
		t.Fatalf("Ensure: %v", err)
	}

	matcher, err := ignorematcher.Load(dir, ignorematcher.Options{})
	if err != nil {
		t.Fatalf("ignore Load: %v", err)
	}

	for _, path := range []string{
		".gitignore",
		".anvien/graph.json",
		"AGENTS.md",
		"CLAUDE.md",
		".claude/skills/anvien/anvien-planner/SKILL.md",
		".codex/config.toml",
		".agents/skills/anvien-ai-context/SKILL.md",
		"anvien-launcher/web-dist/assets/app.js",
		"anvien-launcher/server-bundle/anvien.exe",
		"anvien-launcher/logs/runtime.log",
	} {
		if !matcher.Ignored(path, false) {
			t.Fatalf("path %q was not ignored by managed .gitignore", path)
		}
	}
}

func readGitignore(t *testing.T, dir string) string {
	t.Helper()
	raw, err := os.ReadFile(filepath.Join(dir, ".gitignore"))
	if err != nil {
		t.Fatalf("read .gitignore: %v", err)
	}
	return string(raw)
}
