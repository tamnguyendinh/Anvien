package aicontext

import (
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"testing"
	"testing/fstest"
)

func expectedSkillPackageIDs(t *testing.T) []string {
	t.Helper()
	entries, err := fs.ReadDir(skillSourceFS, "skills")
	if err != nil {
		t.Fatalf("read embedded skills: %v", err)
	}
	ids := make([]string, 0, len(entries))
	for _, entry := range entries {
		if entry.IsDir() {
			ids = append(ids, entry.Name())
		}
	}
	sort.Strings(ids)
	return ids
}

func skillPackageIDs(packages []SkillPackage) []string {
	ids := make([]string, 0, len(packages))
	for _, pkg := range packages {
		ids = append(ids, pkg.Name)
	}
	return ids
}

func findSkillPackage(t *testing.T, packages []SkillPackage, name string) SkillPackage {
	t.Helper()
	for _, pkg := range packages {
		if pkg.Name == name {
			return pkg
		}
	}
	t.Fatalf("missing skill package %s", name)
	return SkillPackage{}
}

func requireContains(t *testing.T, text string, want string) {
	t.Helper()
	if !strings.Contains(text, want) {
		t.Fatalf("missing %q:\n%s", want, text)
	}
}

func TestSkillPackageCatalogDiscoversTopLevelPackagesAndNestedEntries(t *testing.T) {
	packages, err := SkillPackages()
	if err != nil {
		t.Fatalf("SkillPackages: %v", err)
	}
	if got, want := strings.Join(skillPackageIDs(packages), ","), strings.Join(expectedSkillPackageIDs(t), ","); got != want {
		t.Fatalf("skill package ids mismatch:\n got: %s\nwant: %s", got, want)
	}

	debugging := findSkillPackage(t, packages, "debugging")
	if got, want := len(debugging.Entries), 5; got != want {
		t.Fatalf("debugging entries = %d, want %d", got, want)
	}
	for _, want := range []string{
		"debugging-parent-skill/SKILL.md",
		"defense-in-depth/SKILL.md",
		"root-cause-tracing/SKILL.md",
		"systematic-debugging/SKILL.md",
		"verification-before-completion/SKILL.md",
	} {
		if !packageHasEntry(debugging, want) {
			t.Fatalf("debugging package missing entry %s", want)
		}
	}
	if !packageHasFile(debugging, "root-cause-tracing/find-polluter.sh") {
		t.Fatalf("debugging package missing script payload")
	}

	documentSkills := findSkillPackage(t, packages, "document-skills")
	if got, want := len(documentSkills.Entries), 4; got != want {
		t.Fatalf("document-skills entries = %d, want %d", got, want)
	}
	if !packageHasFile(documentSkills, "docx/scripts/document.py") {
		t.Fatalf("document-skills package missing nested script payload")
	}

	uiStyling := findSkillPackage(t, packages, "ui-styling")
	if !packageHasFile(uiStyling, "scripts/shadcn_add.py") {
		t.Fatalf("ui-styling package missing script payload")
	}
	if !packageHasFile(uiStyling, "scripts/.coverage") {
		t.Fatalf("ui-styling package should include dotfile payload")
	}
	if !strings.HasPrefix(uiStyling.Hash, "sha256:") {
		t.Fatalf("ui-styling package hash missing sha256 prefix: %s", uiStyling.Hash)
	}

	files, err := BaseSkillFiles()
	if err != nil {
		t.Fatalf("BaseSkillFiles: %v", err)
	}
	entryCount := 0
	for _, pkg := range packages {
		entryCount += len(pkg.Entries)
		if strings.TrimSpace(pkg.Description) == "" {
			t.Fatalf("package %s missing description", pkg.Name)
		}
		for _, entry := range pkg.Entries {
			if strings.TrimSpace(entry.Name) == "" || strings.TrimSpace(entry.Description) == "" {
				t.Fatalf("entry missing metadata: %#v", entry)
			}
		}
	}
	if len(files) != entryCount {
		t.Fatalf("BaseSkillFiles returned %d entries, want %d", len(files), entryCount)
	}
}

func TestSkillPackageCatalogRejectsPackageWithoutSkillEntry(t *testing.T) {
	_, err := discoverSkillPackages(fstest.MapFS{
		"skills/empty-package/readme.md": {Data: []byte("# No skill entry\n")},
	})
	if err == nil || !strings.Contains(err.Error(), "has no SKILL.md entry") {
		t.Fatalf("expected missing SKILL.md validation error, got %v", err)
	}
}

func TestSkillPackagesForRepoPrefersRuntimeFilesystemSource(t *testing.T) {
	dir := t.TempDir()
	skillPath := filepath.Join(dir, filepath.FromSlash("internal/aicontext/skills/runtime-only/SKILL.md"))
	payloadPath := filepath.Join(dir, filepath.FromSlash("internal/aicontext/skills/runtime-only/scripts/.marker"))
	if err := os.MkdirAll(filepath.Dir(skillPath), 0o755); err != nil {
		t.Fatalf("mkdir runtime skill: %v", err)
	}
	if err := os.MkdirAll(filepath.Dir(payloadPath), 0o755); err != nil {
		t.Fatalf("mkdir runtime payload: %v", err)
	}
	if err := os.WriteFile(skillPath, []byte("---\nname: runtime-only\ndescription: Runtime filesystem skill\n---\n# Runtime Skill\n"), 0o644); err != nil {
		t.Fatalf("write runtime skill: %v", err)
	}
	if err := os.WriteFile(payloadPath, []byte("one\n"), 0o644); err != nil {
		t.Fatalf("write runtime payload: %v", err)
	}

	packages, err := SkillPackagesForRepo(dir)
	if err != nil {
		t.Fatalf("SkillPackagesForRepo: %v", err)
	}
	if got := strings.Join(skillPackageIDs(packages), ","); got != "runtime-only" {
		t.Fatalf("runtime package ids mismatch: %s", got)
	}
	runtimeSkill := findSkillPackage(t, packages, "runtime-only")
	if runtimeSkill.SourceRoot != "internal/aicontext/skills/runtime-only" {
		t.Fatalf("unexpected runtime source root: %s", runtimeSkill.SourceRoot)
	}
	if !packageHasFile(runtimeSkill, "scripts/.marker") {
		t.Fatalf("runtime package missing dotfile payload")
	}
	firstHash := runtimeSkill.Hash

	files, err := BaseSkillFilesForRepo(dir)
	if err != nil {
		t.Fatalf("BaseSkillFilesForRepo: %v", err)
	}
	if len(files) != 1 || files[0].InstallPath != "runtime-only/SKILL.md" || !strings.Contains(files[0].Content, "Runtime Skill") {
		t.Fatalf("unexpected runtime base skill files: %#v", files)
	}

	if err := os.WriteFile(payloadPath, []byte("two\n"), 0o644); err != nil {
		t.Fatalf("update runtime payload: %v", err)
	}
	updatedPackages, err := SkillPackagesForRepo(dir)
	if err != nil {
		t.Fatalf("SkillPackagesForRepo after update: %v", err)
	}
	updatedSkill := findSkillPackage(t, updatedPackages, "runtime-only")
	if updatedSkill.Hash == firstHash {
		t.Fatalf("runtime package hash did not change after payload update: %s", firstHash)
	}

	_, installedPackages, err := GenerateAIContextFiles(dir, "RuntimeRepo", Stats{Nodes: 1}, Options{})
	if err != nil {
		t.Fatalf("GenerateAIContextFiles: %v", err)
	}
	if got := strings.Join(installedPackages, ","); got != "runtime-only" {
		t.Fatalf("installed runtime package ids mismatch: %s", got)
	}
	installedPayload := filepath.Join(dir, ".claude", "skills", "anvien", "runtime-only", "scripts", ".marker")
	raw, err := os.ReadFile(installedPayload)
	if err != nil {
		t.Fatalf("read installed runtime payload: %v", err)
	}
	if string(raw) != "two\n" {
		t.Fatalf("installed payload not updated from runtime source: %q", raw)
	}
	manifest, err := loadSkillManifest(filepath.Join(dir, ".claude", "skills", "anvien"))
	if err != nil {
		t.Fatalf("load runtime manifest: %v", err)
	}
	entry, ok := manifest.Skills["runtime-only"]
	if !ok || entry.SourceRoot != "internal/aicontext/skills/runtime-only" || entry.Hash != updatedSkill.Hash {
		t.Fatalf("unexpected runtime manifest entry: %#v", entry)
	}
}

func TestGenerateAIContextFilesCreatesManagedContextAndSkillPackages(t *testing.T) {
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
	repoLocalSkill := filepath.Join(dir, ".claude", "skills", "anvien", "my-repo-custom-skill", "SKILL.md")
	if err := os.MkdirAll(filepath.Dir(repoLocalSkill), 0o755); err != nil {
		t.Fatalf("mkdir repo-local skill: %v", err)
	}
	if err := os.WriteFile(repoLocalSkill, []byte("---\nname: my-repo-custom-skill\ndescription: repo local\n---\n# Custom\n"), 0o644); err != nil {
		t.Fatalf("write repo-local skill: %v", err)
	}

	files, installedPackages, err := GenerateAIContextFiles(dir, "TestProject", stats, Options{})
	if err != nil {
		t.Fatalf("GenerateAIContextFiles: %v", err)
	}
	if len(files) == 0 {
		t.Fatalf("expected generated files")
	}
	if got, want := strings.Join(installedPackages, ","), strings.Join(expectedSkillPackageIDs(t), ","); got != want {
		t.Fatalf("installed package ids mismatch:\n got: %s\nwant: %s", got, want)
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
		"## Always Do",
		"before using any Anvien CLI command, MCP tool, MCP resource, Web/API view",
		"## Never Do",
		"## Command Selection Guide",
		"Prefer CLI `anvien file-context <path> --repo <repo> --json`; use `anvien context file <path> --repo <repo>` only when you want the context wrapper / human-oriented view",
		"`anvien graph-health explain \"File:<path>\" --repo <repo> --json`",
		"## Resources",
		"anvien://repo/<repo>/schema",
		"## MCP Prompts",
		"`detect_impact`",
		"## Skill Selection Guide",
		"Anvien installs every top-level package discovered under `internal/aicontext/skills/` when that source folder exists for the repo; otherwise it uses the embedded skill catalog.",
		"| `debugging` | `.claude/skills/anvien/debugging/debugging-parent-skill/SKILL.md`<br>`.claude/skills/anvien/debugging/defense-in-depth/SKILL.md`",
		"root-cause-tracing/SKILL.md",
		"`.claude/skills/anvien/debugging/`",
		"| `ui-styling` | `.claude/skills/anvien/ui-styling/SKILL.md` |",
		"`.claude/skills/anvien/ui-styling/`",
		"file-context",
		"file-hotspots",
	} {
		requireContains(t, text, want)
	}
	if strings.Contains(text, "Use Anvien workflow skills only for the retained domains below") {
		t.Fatalf("AGENTS.md contains obsolete four-skill wording:\n%s", text)
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
	if _, err := os.Stat(repoLocalSkill); err != nil {
		t.Fatalf("repo-local unmanifested skill should be preserved: %v", err)
	}
	for _, forbidden := range []string{oldDisplay, oldLower, oldUpper, "." + oldLower, oldLower + "://", oldLower + "-"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("AGENTS.md contains old generated name %q:\n%s", forbidden, text)
		}
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

	for _, want := range []string{
		filepath.Join(dir, ".claude", "skills", "anvien", "debugging", "root-cause-tracing", "find-polluter.sh"),
		filepath.Join(dir, ".claude", "skills", "anvien", "ui-styling", "scripts", "shadcn_add.py"),
		filepath.Join(dir, ".claude", "skills", "anvien", "ui-styling", "scripts", ".coverage"),
		filepath.Join(dir, ".claude", "skills", "anvien", "ui-styling", "canvas-fonts", "ArsenalSC-Regular.ttf"),
	} {
		if info, err := os.Stat(want); err != nil || info.IsDir() {
			t.Fatalf("expected package payload file %s: %v", want, err)
		}
	}
	manifest, err := loadSkillManifest(filepath.Join(dir, ".claude", "skills", "anvien"))
	if err != nil {
		t.Fatalf("load skill manifest: %v", err)
	}
	if manifest.ManagedBy != skillManifestOwner || len(manifest.Skills) != len(expectedSkillPackageIDs(t)) {
		t.Fatalf("unexpected manifest: %#v", manifest)
	}
}

func TestInstallBaseSkillsToUsesManifestBoundary(t *testing.T) {
	dir := t.TempDir()
	repoLocalSkill := filepath.Join(dir, "custom-skill", "SKILL.md")
	if err := os.MkdirAll(filepath.Dir(repoLocalSkill), 0o755); err != nil {
		t.Fatalf("mkdir custom skill: %v", err)
	}
	if err := os.WriteFile(repoLocalSkill, []byte("---\nname: custom-skill\ndescription: user skill\n---\n# Custom\n"), 0o644); err != nil {
		t.Fatalf("write custom skill: %v", err)
	}

	result, err := InstallSkillPackagesTo(dir)
	if err != nil {
		t.Fatalf("InstallSkillPackagesTo: %v", err)
	}
	if result.Installed != len(expectedSkillPackageIDs(t)) || result.Preserved != 1 || result.Skipped != 0 {
		t.Fatalf("unexpected first install result: %#v", result)
	}
	if _, err := os.Stat(repoLocalSkill); err != nil {
		t.Fatalf("custom non-Anvien skill should be preserved: %v", err)
	}

	localFileInsideManagedPackage := filepath.Join(dir, "debugging", "local-note.md")
	if err := os.WriteFile(localFileInsideManagedPackage, []byte("repo local note\n"), 0o644); err != nil {
		t.Fatalf("write local file inside managed package: %v", err)
	}
	tamperedSkill := filepath.Join(dir, "ui-styling", "SKILL.md")
	if err := os.WriteFile(tamperedSkill, []byte("tampered\n"), 0o644); err != nil {
		t.Fatalf("tamper managed skill: %v", err)
	}

	manifest, err := loadSkillManifest(dir)
	if err != nil {
		t.Fatalf("load manifest: %v", err)
	}
	oldManaged := filepath.Join(dir, "old-managed", "SKILL.md")
	if err := os.MkdirAll(filepath.Dir(oldManaged), 0o755); err != nil {
		t.Fatalf("mkdir old managed: %v", err)
	}
	if err := os.WriteFile(oldManaged, []byte("# Old Managed\n"), 0o644); err != nil {
		t.Fatalf("write old managed: %v", err)
	}
	manifest.Skills["old-managed"] = skillManifestEntry{
		InstallPath: "old-managed",
		SourceRoot:  "skills/old-managed",
		Hash:        "sha256:old",
		Managed:     true,
		EntryCount:  1,
		FileCount:   1,
		Files:       map[string]string{"SKILL.md": "sha256:old"},
	}
	if err := writeSkillManifest(dir, manifest); err != nil {
		t.Fatalf("write manifest with stale entry: %v", err)
	}

	result, err = InstallSkillPackagesTo(dir)
	if err != nil {
		t.Fatalf("second InstallSkillPackagesTo: %v", err)
	}
	if result.Updated != 1 || result.Stale != 1 || result.Preserved != 1 {
		t.Fatalf("unexpected second install result: %#v", result)
	}
	if _, err := os.Stat(localFileInsideManagedPackage); err != nil {
		t.Fatalf("local file inside managed package should be preserved: %v", err)
	}
	if raw, err := os.ReadFile(tamperedSkill); err != nil || strings.Contains(string(raw), "tampered") {
		t.Fatalf("managed package file was not repaired: %v\n%s", err, raw)
	}
	if _, err := os.Stat(oldManaged); err != nil {
		t.Fatalf("stale managed package payload should be preserved unless explicitly pruned: %v", err)
	}
	manifest, err = loadSkillManifest(dir)
	if err != nil {
		t.Fatalf("reload manifest: %v", err)
	}
	if entry, ok := manifest.Skills["old-managed"]; !ok || !entry.Stale {
		t.Fatalf("old managed package should be marked stale, got %#v", manifest.Skills["old-managed"])
	}
}

func TestInstallBaseSkillsToRejectsUnmanagedNameCollision(t *testing.T) {
	dir := t.TempDir()
	collision := filepath.Join(dir, "ui-styling", "SKILL.md")
	if err := os.MkdirAll(filepath.Dir(collision), 0o755); err != nil {
		t.Fatalf("mkdir collision: %v", err)
	}
	if err := os.WriteFile(collision, []byte("---\nname: local-ui-styling\ndescription: local\n---\n# Local\n"), 0o644); err != nil {
		t.Fatalf("write collision: %v", err)
	}
	result, err := InstallSkillPackagesTo(dir)
	if err == nil || !strings.Contains(err.Error(), "already exists and is not managed by Anvien") {
		t.Fatalf("expected unmanaged collision error, got %v", err)
	}
	if result.Collisions != 1 {
		t.Fatalf("expected one collision, got %#v", result)
	}
	if raw, err := os.ReadFile(collision); err != nil || !strings.Contains(string(raw), "local-ui-styling") {
		t.Fatalf("collision file should be preserved: %v\n%s", err, raw)
	}
}

func TestInstallBaseSkillsToAdoptsLegacyGeneratedAnvienPackage(t *testing.T) {
	dir := t.TempDir()
	legacy := filepath.Join(dir, "anvien-planner", "SKILL.md")
	if err := os.MkdirAll(filepath.Dir(legacy), 0o755); err != nil {
		t.Fatalf("mkdir legacy skill: %v", err)
	}
	if err := os.WriteFile(legacy, []byte("---\nname: anvien-planner\ndescription: old generated planner\n---\n# Legacy Planner\n"), 0o644); err != nil {
		t.Fatalf("write legacy skill: %v", err)
	}

	result, err := InstallSkillPackagesTo(dir)
	if err != nil {
		t.Fatalf("InstallSkillPackagesTo: %v", err)
	}
	if result.Adopted != 1 {
		t.Fatalf("expected one adopted legacy package, got %#v", result)
	}
	raw, err := os.ReadFile(legacy)
	if err != nil {
		t.Fatalf("read adopted legacy skill: %v", err)
	}
	if strings.Contains(string(raw), "Legacy Planner") || !strings.Contains(string(raw), "Standard Plan Set") {
		t.Fatalf("legacy skill was not updated from embedded source:\n%s", raw)
	}
}

func TestSkillGuidanceProtectsExpandedCommandSurface(t *testing.T) {
	files, err := BaseSkillFiles()
	if err != nil {
		t.Fatalf("BaseSkillFiles: %v", err)
	}
	contentByInstallPath := make(map[string]string, len(files))
	var combined strings.Builder
	for _, file := range files {
		if strings.TrimSpace(file.Content) == "" || strings.TrimSpace(file.Description) == "" {
			t.Fatalf("skill entry missing content or description: %#v", file)
		}
		if strings.Contains(file.Content, "Use Anvien tools to accomplish this task.") {
			t.Fatalf("skill entry %s uses fallback placeholder content", file.InstallPath)
		}
		contentByInstallPath[file.InstallPath] = file.Content
		combined.WriteString(file.Content)
		combined.WriteString("\n")
	}

	checks := map[string][]string{
		"anvien-api-surface/SKILL.md": {
			"route_map",
			"anvien api route-map",
			"api_impact",
			"Do not invent CLI commands",
		},
		"anvien-refactoring/SKILL.md": {
			"behavior-preserving",
			"impact",
			"detect-changes",
			"graph-guided dry run",
		},
		"anvien-debugging/SKILL.md": {
			"Reproduce or capture the symptom",
			"graph-health",
			"resolution-inventory",
			"Query Reliability Rule",
		},
		"anvien-planner/SKILL.md": {
			"Standard Plan Set",
			"YYYY-MM-DD-<slug>",
			"Evidence Ledger",
			"Benchmark Ledger",
			"Every checklist item must be a complete mini-plan by itself",
		},
		"ui-styling/SKILL.md": {
			"scripts/shadcn_add.py",
			"references/shadcn-components.md",
		},
		"debugging/root-cause-tracing/SKILL.md": {
			"find-polluter.sh",
		},
	}
	for installPath, fragments := range checks {
		content := contentByInstallPath[installPath]
		if content == "" {
			t.Fatalf("missing skill content for %s", installPath)
		}
		for _, fragment := range fragments {
			if !strings.Contains(content, fragment) {
				t.Fatalf("%s missing guidance fragment %q:\n%s", installPath, fragment, content)
			}
		}
	}

	allGuidance := combined.String()
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

func packageHasEntry(pkg SkillPackage, packagePath string) bool {
	for _, entry := range pkg.Entries {
		if entry.PackagePath == packagePath {
			return true
		}
	}
	return false
}

func packageHasFile(pkg SkillPackage, packagePath string) bool {
	clean := path.Clean(packagePath)
	for _, file := range pkg.Files {
		if file.PackagePath == clean {
			return true
		}
	}
	return false
}
