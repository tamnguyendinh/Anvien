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
	if len(packages) == 0 {
		t.Fatalf("expected at least one embedded skill package")
	}

	files, err := BaseSkillFiles()
	if err != nil {
		t.Fatalf("BaseSkillFiles: %v", err)
	}
	entryCount := 0
	for _, pkg := range packages {
		entryCount += len(pkg.Entries)
		if len(pkg.Files) == 0 {
			t.Fatalf("package %s has no payload files", pkg.Name)
		}
		if !strings.HasPrefix(pkg.Hash, "sha256:") {
			t.Fatalf("package %s hash missing sha256 prefix: %s", pkg.Name, pkg.Hash)
		}
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

func TestSkillPackageCatalogDiscoversSyntheticNestedEntriesAndPayloads(t *testing.T) {
	packages, err := discoverSkillPackages(fstest.MapFS{
		"skills/multi/SKILL.md":                 {Data: []byte("---\nname: multi\ndescription: Multi skill\n---\n# Multi\n")},
		"skills/multi/child/SKILL.md":           {Data: []byte("---\nname: child\ndescription: Child skill\n---\n# Child\n")},
		"skills/multi/scripts/.marker":          {Data: []byte("marker\n")},
		"skills/multi/references/example.md":    {Data: []byte("# Reference\n")},
		"skills/second/SKILL.md":                {Data: []byte("---\nname: second\ndescription: Second skill\n---\n# Second\n")},
		"skills/second/assets/template-file.md": {Data: []byte("# Template\n")},
	})
	if err != nil {
		t.Fatalf("discoverSkillPackages synthetic: %v", err)
	}
	if got, want := strings.Join(skillPackageIDs(packages), ","), "multi,second"; got != want {
		t.Fatalf("synthetic package ids mismatch:\n got: %s\nwant: %s", got, want)
	}
	multi := findSkillPackage(t, packages, "multi")
	if got, want := len(multi.Entries), 2; got != want {
		t.Fatalf("multi entries = %d, want %d", got, want)
	}
	if !packageHasEntry(multi, "child/SKILL.md") {
		t.Fatalf("multi package missing nested entry")
	}
	if !packageHasFile(multi, "scripts/.marker") || !packageHasFile(multi, "references/example.md") {
		t.Fatalf("multi package missing nested payloads: %#v", multi.Files)
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
	for _, root := range []string{
		filepath.Join(dir, ".agents", "skills"),
		filepath.Join(dir, ".claude", "skills"),
	} {
		installedPayload := filepath.Join(root, "runtime-only", "scripts", ".marker")
		raw, err := os.ReadFile(installedPayload)
		if err != nil {
			t.Fatalf("read installed runtime payload %s: %v", installedPayload, err)
		}
		if string(raw) != "two\n" {
			t.Fatalf("installed payload %s not updated from runtime source: %q", installedPayload, raw)
		}
	}
	manifest, _, err := loadSkillManifestFile(filepath.Join(dir, ".agents", "skills"), codexSkillManifestFileName)
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
	legacyAnvienRoot := filepath.Join(dir, ".claude", "skills", "anvien")
	legacyManagedSkill := filepath.Join(legacyAnvienRoot, "legacy-managed", "SKILL.md")
	if err := os.MkdirAll(filepath.Dir(legacyManagedSkill), 0o755); err != nil {
		t.Fatalf("mkdir legacy Anvien skill: %v", err)
	}
	if err := os.WriteFile(legacyManagedSkill, []byte("# Legacy Managed\n"), 0o644); err != nil {
		t.Fatalf("write legacy Anvien skill: %v", err)
	}
	legacyManifest := newSkillManifest()
	legacyManifest.Skills["legacy-managed"] = skillManifestEntry{
		InstallPath: "legacy-managed",
		SourceRoot:  "skills/legacy-managed",
		Hash:        "sha256:legacy",
		Managed:     true,
		EntryCount:  1,
		FileCount:   1,
	}
	if err := writeSkillManifest(legacyAnvienRoot, skillManifestFileName, legacyManifest); err != nil {
		t.Fatalf("write legacy Anvien manifest: %v", err)
	}
	codexCustomSkill := filepath.Join(dir, ".agents", "skills", "my-repo-custom-skill", "SKILL.md")
	claudeCustomSkill := filepath.Join(dir, ".claude", "skills", "my-repo-custom-skill", "SKILL.md")
	for _, customSkill := range []string{codexCustomSkill, claudeCustomSkill} {
		if err := os.MkdirAll(filepath.Dir(customSkill), 0o755); err != nil {
			t.Fatalf("mkdir repo-local skill: %v", err)
		}
		if err := os.WriteFile(customSkill, []byte("---\nname: my-repo-custom-skill\ndescription: repo local\n---\n# Custom\n"), 0o644); err != nil {
			t.Fatalf("write repo-local skill: %v", err)
		}
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
	packages, err := SkillPackages()
	if err != nil {
		t.Fatalf("SkillPackages: %v", err)
	}
	if len(packages) == 0 {
		t.Fatalf("expected at least one skill package")
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
		"AI agent chooses the skill that fits the work.",
		"| When you need to... | Use |",
		"|---------------------|-----|",
		"file-context",
		"file-hotspots",
	} {
		requireContains(t, text, want)
	}
	for _, forbidden := range []string{
		"TestProject",
		"This project is indexed by Anvien",
		"50 symbols",
		"100 relationships",
		"5 execution flows",
		".claude/skills/anvien/",
		".agents/skills/anvien/",
	} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("AGENTS.md contains forbidden fragment %q:\n%s", forbidden, text)
		}
	}
	requireContains(t, text, "`.agents/skills/"+packages[0].Entries[0].InstallPath+"`")
	foundProblemSolving := false
	for _, pkg := range packages {
		if pkg.Name != "problem-solving" {
			continue
		}
		foundProblemSolving = true
		entry := primarySkillEntry(pkg)
		requireContains(t, text, "Use when the user asks to solve a hard problem.")
		requireContains(t, text, "`.agents/skills/"+entry.InstallPath+"`")
		if strings.Contains(text, "`.agents/skills/problem-solving/collision-zone-thinking/SKILL.md`") {
			t.Fatalf("Skill Selection Guide should show only the primary problem-solving entry:\n%s", text)
		}
		break
	}
	if !foundProblemSolving {
		t.Fatalf("expected problem-solving skill package in catalog")
	}
	if strings.Contains(text, "Anvien installs every top-level package discovered") {
		t.Fatalf("Skill Selection Guide should not include generated namespace explanation:\n%s", text)
	}
	if strings.Contains(text, "Package `"+packages[0].Name+"`") {
		t.Fatalf("Skill Selection Guide should not include package labels in Use column:\n%s", text)
	}
	if strings.Contains(text, "`.agents/skills/"+packages[0].InstallRoot+"/`") {
		t.Fatalf("Skill Selection Guide should not include package root directories in Use column:\n%s", text)
	}
	if strings.Contains(text, "| Package | Entries | Use |") {
		t.Fatalf("Skill Selection Guide should use When/Use table form, got old package table:\n%s", text)
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
	if _, err := os.Stat(legacyAnvienRoot); !os.IsNotExist(err) {
		t.Fatalf("legacy managed Anvien namespace should be removed: %v", err)
	}
	for _, customSkill := range []string{codexCustomSkill, claudeCustomSkill} {
		if _, err := os.Stat(customSkill); err != nil {
			t.Fatalf("repo-local custom skill should be preserved at %s: %v", customSkill, err)
		}
	}
	for _, forbidden := range []string{oldDisplay, oldLower, oldUpper, "." + oldLower, oldLower + "://", oldLower + "-"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("AGENTS.md contains old generated name %q:\n%s", forbidden, text)
		}
	}

	claudeContent, err := os.ReadFile(filepath.Join(dir, "CLAUDE.md"))
	if err != nil {
		t.Fatalf("read CLAUDE.md: %v", err)
	}
	claudeText := string(claudeContent)
	requireContains(t, claudeText, "`.claude/skills/"+packages[0].Entries[0].InstallPath+"`")
	if strings.Contains(claudeText, ".agents/skills/") || strings.Contains(claudeText, ".claude/skills/anvien/") {
		t.Fatalf("CLAUDE.md uses the wrong skill surface:\n%s", claudeText)
	}
	if strings.Contains(claudeText, "This project is indexed by Anvien") ||
		strings.Contains(claudeText, "50 symbols") ||
		strings.Contains(claudeText, "100 relationships") ||
		strings.Contains(claudeText, "5 execution flows") {
		t.Fatalf("CLAUDE.md contains volatile indexed-project sentence:\n%s", claudeText)
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

	for _, pkg := range packages {
		for _, file := range pkg.Files {
			for _, root := range []string{
				filepath.Join(dir, ".agents", "skills"),
				filepath.Join(dir, ".claude", "skills"),
			} {
				want := filepath.Join(root, filepath.FromSlash(file.InstallPath))
				if info, err := os.Stat(want); err != nil || info.IsDir() {
					t.Fatalf("expected package payload file %s: %v", want, err)
				}
			}
		}
	}
	for _, manifestSpec := range []struct {
		root string
		name string
	}{
		{filepath.Join(dir, ".agents", "skills"), codexSkillManifestFileName},
		{filepath.Join(dir, ".claude", "skills"), claudeSkillManifestFileName},
	} {
		manifest, _, err := loadSkillManifestFile(manifestSpec.root, manifestSpec.name)
		if err != nil {
			t.Fatalf("load skill manifest: %v", err)
		}
		if manifest.ManagedBy != skillManifestOwner || len(manifest.Skills) != len(expectedSkillPackageIDs(t)) {
			t.Fatalf("unexpected manifest at %s: %#v", filepath.Join(manifestSpec.root, manifestSpec.name), manifest)
		}
	}
}

func TestInstallSkillPackagesToSyncsManagedRootsAndPreservesCustomRoots(t *testing.T) {
	dir := t.TempDir()
	customTarget := filepath.Join(dir, "custom-skill", "SKILL.md")
	if err := os.MkdirAll(filepath.Dir(customTarget), 0o755); err != nil {
		t.Fatalf("mkdir custom target: %v", err)
	}
	if err := os.WriteFile(customTarget, []byte("# Custom\n"), 0o644); err != nil {
		t.Fatalf("write custom target: %v", err)
	}

	alpha := testSkillPackage("alpha", map[string]string{
		"SKILL.md":         "---\nname: alpha\ndescription: Alpha skill\n---\n# Alpha\n",
		"scripts/note.txt": "one\n",
	})
	old := testSkillPackage("old", map[string]string{
		"SKILL.md": "---\nname: old\ndescription: Old managed skill\n---\n# Old\n",
	})
	result, err := installSkillPackagesTo(dir, []SkillPackage{alpha, old})
	if err != nil {
		t.Fatalf("installSkillPackagesTo first sync: %v", err)
	}
	if result.Installed != 2 || result.Written != 3 || result.Deleted != 0 || result.Preserved != 0 || result.Stale != 0 {
		t.Fatalf("unexpected first sync result: %#v", result)
	}
	assertFileContent(t, filepath.Join(dir, "alpha", "scripts", "note.txt"), "one\n")
	assertFileContent(t, customTarget, "# Custom\n")
	assertManifestPackageNames(t, dir, []string{"alpha", "old"})

	if err := os.Remove(filepath.Join(dir, "alpha", "scripts", "note.txt")); err != nil {
		t.Fatalf("remove generated payload: %v", err)
	}
	if err := os.WriteFile(filepath.Join(dir, "alpha", "SKILL.md"), []byte("tampered\n"), 0o644); err != nil {
		t.Fatalf("tamper generated skill: %v", err)
	}
	extraInsidePackage := filepath.Join(dir, "alpha", "local-note.md")
	if err := os.WriteFile(extraInsidePackage, []byte("repo local note\n"), 0o644); err != nil {
		t.Fatalf("write extra generated-root file: %v", err)
	}
	alpha = testSkillPackage("alpha", map[string]string{
		"SKILL.md":          "---\nname: alpha\ndescription: Alpha skill\n---\n# Alpha v2\n",
		"scripts/note.txt":  "two\n",
		"references/new.md": "new\n",
	})
	result, err = installSkillPackagesTo(dir, []SkillPackage{alpha})
	if err != nil {
		t.Fatalf("installSkillPackagesTo repair sync: %v", err)
	}
	if result.Updated != 1 || result.Written != 2 || result.Overwritten != 1 || result.Deleted != 2 || result.Stale != 0 || result.Preserved != 0 {
		t.Fatalf("unexpected repair sync result: %#v", result)
	}
	assertFileContent(t, filepath.Join(dir, "alpha", "SKILL.md"), "---\nname: alpha\ndescription: Alpha skill\n---\n# Alpha v2\n")
	assertFileContent(t, filepath.Join(dir, "alpha", "scripts", "note.txt"), "two\n")
	assertFileContent(t, filepath.Join(dir, "alpha", "references", "new.md"), "new\n")
	assertNotExists(t, extraInsidePackage)
	assertNotExists(t, filepath.Join(dir, "old", "SKILL.md"))
	assertFileContent(t, customTarget, "# Custom\n")
	assertManifestPackageNames(t, dir, []string{"alpha"})

	alpha = testSkillPackage("alpha", map[string]string{
		"SKILL.md":              "---\nname: alpha\ndescription: Alpha skill\n---\n# Alpha v2\n",
		"scripts/note.txt":      "two\n",
		"references/renamed.md": "renamed\n",
	})
	result, err = installSkillPackagesTo(dir, []SkillPackage{alpha})
	if err != nil {
		t.Fatalf("installSkillPackagesTo rename sync: %v", err)
	}
	if result.Updated != 1 || result.Written != 1 || result.Deleted != 1 || result.SkippedFiles != 2 {
		t.Fatalf("unexpected rename sync result: %#v", result)
	}
	assertNotExists(t, filepath.Join(dir, "alpha", "references", "new.md"))
	assertFileContent(t, filepath.Join(dir, "alpha", "references", "renamed.md"), "renamed\n")

	result, err = installSkillPackagesTo(dir, nil)
	if err != nil {
		t.Fatalf("installSkillPackagesTo delete package sync: %v", err)
	}
	if result.Discovered != 0 || result.Deleted != 3 || result.Stale != 0 || result.Preserved != 0 {
		t.Fatalf("unexpected delete package sync result: %#v", result)
	}
	assertNotExists(t, filepath.Join(dir, "alpha"))
	assertFileContent(t, customTarget, "# Custom\n")
	assertManifestPackageNames(t, dir, nil)
}

func TestInstallSkillPackagesToRejectsForeignSameNameRootAndAdoptsExactMatch(t *testing.T) {
	alpha := testSkillPackage("alpha", map[string]string{
		"SKILL.md":         "---\nname: alpha\ndescription: Alpha skill\n---\n# Alpha\n",
		"scripts/note.txt": "one\n",
	})

	collisionDir := t.TempDir()
	foreignSkill := filepath.Join(collisionDir, "alpha", "SKILL.md")
	if err := os.MkdirAll(filepath.Dir(foreignSkill), 0o755); err != nil {
		t.Fatalf("mkdir foreign skill: %v", err)
	}
	if err := os.WriteFile(foreignSkill, []byte("# Foreign Alpha\n"), 0o644); err != nil {
		t.Fatalf("write foreign skill: %v", err)
	}
	result, err := installSkillPackagesTo(collisionDir, []SkillPackage{alpha})
	if err == nil || !strings.Contains(err.Error(), "skill package root collision") {
		t.Fatalf("expected foreign root collision, got result=%#v err=%v", result, err)
	}
	if result.Collisions != 1 {
		t.Fatalf("collision count = %d, want 1", result.Collisions)
	}
	assertFileContent(t, foreignSkill, "# Foreign Alpha\n")

	adoptionDir := t.TempDir()
	writeSkillPackageFilesWithoutManifest(t, adoptionDir, alpha)
	result, err = installSkillPackagesTo(adoptionDir, []SkillPackage{alpha})
	if err != nil {
		t.Fatalf("installSkillPackagesTo exact adoption: %v", err)
	}
	if result.Adopted != 1 || result.Written != 0 || result.Overwritten != 0 || result.Deleted != 0 || result.Collisions != 0 {
		t.Fatalf("unexpected adoption result: %#v", result)
	}
	assertManifestPackageNames(t, adoptionDir, []string{"alpha"})

	extraDir := t.TempDir()
	writeSkillPackageFilesWithoutManifest(t, extraDir, alpha)
	extraFile := filepath.Join(extraDir, "alpha", "extra.md")
	if err := os.WriteFile(extraFile, []byte("extra\n"), 0o644); err != nil {
		t.Fatalf("write extra file: %v", err)
	}
	result, err = installSkillPackagesTo(extraDir, []SkillPackage{alpha})
	if err == nil || !strings.Contains(err.Error(), "skill package root collision") {
		t.Fatalf("expected extra-file collision, got result=%#v err=%v", result, err)
	}
	if result.Collisions != 1 {
		t.Fatalf("extra-file collision count = %d, want 1", result.Collisions)
	}
	assertFileContent(t, extraFile, "extra\n")
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
		"api-surface/SKILL.md": {
			"route_map",
			"anvien api route-map",
			"api_impact",
			"Do not invent CLI commands",
		},
		"refactoring/SKILL.md": {
			"behavior-preserving",
			"impact",
			"detect-changes",
			"graph-guided dry run",
		},
		"debugging/SKILL.md": {
			"MUST understand the root cause before fixing.",
			"MUST write a plan before fixing.",
			"MUST run a full build before testing when active repository rules require it.",
			"anvien file-context <path> --repo <repo> --json",
			"anvien graph-health explain \"File:<path>\" --repo <repo> --json",
			"anvien detect-changes --repo <repo> --scope all",
		},
		"planner/SKILL.md": {
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

func TestGenerateAIContextFilesAlwaysOmitsVolatileIndexedProjectSentence(t *testing.T) {
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
		t.Fatalf("AGENTS.md contains volatile stats:\n%s", text)
	}
	if strings.Contains(text, "This project is indexed by Anvien") || strings.Contains(text, "TestProject") {
		t.Fatalf("AGENTS.md contains obsolete indexed-project sentence:\n%s", text)
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

func testSkillPackage(name string, files map[string]string) SkillPackage {
	pkg := SkillPackage{
		Name:        name,
		Description: name + " test skill",
		SourceRoot:  path.Join("skills", name),
		InstallRoot: name,
	}
	packagePaths := make([]string, 0, len(files))
	for packagePath := range files {
		packagePaths = append(packagePaths, packagePath)
	}
	sort.Strings(packagePaths)
	for _, packagePath := range packagePaths {
		raw := []byte(files[packagePath])
		file := SkillPackageFile{
			SourcePath:  path.Join("skills", name, packagePath),
			PackagePath: packagePath,
			InstallPath: path.Join(name, packagePath),
			Hash:        hashBytes(raw),
			SizeBytes:   int64(len(raw)),
			Content:     raw,
		}
		pkg.Files = append(pkg.Files, file)
		if path.Base(packagePath) == "SKILL.md" {
			pkg.Entries = append(pkg.Entries, SkillEntry{
				Name:        name,
				Description: pkg.Description,
				SourcePath:  file.SourcePath,
				PackagePath: file.PackagePath,
				InstallPath: file.InstallPath,
			})
		}
	}
	pkg.Hash = packageHash(pkg.Files)
	return pkg
}

func assertFileContent(t *testing.T, filePath string, want string) {
	t.Helper()
	raw, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("read %s: %v", filePath, err)
	}
	if string(raw) != want {
		t.Fatalf("%s content mismatch:\n got: %q\nwant: %q", filePath, raw, want)
	}
}

func assertNotExists(t *testing.T, filePath string) {
	t.Helper()
	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		t.Fatalf("%s should not exist: %v", filePath, err)
	}
}

func assertManifestPackageNames(t *testing.T, targetDir string, want []string) {
	t.Helper()
	manifest, err := loadSkillManifest(targetDir)
	if err != nil {
		t.Fatalf("load manifest: %v", err)
	}
	got := make([]string, 0, len(manifest.Skills))
	for name, entry := range manifest.Skills {
		if entry.Stale {
			t.Fatalf("manifest package %s should not be stale: %#v", name, entry)
		}
		got = append(got, name)
	}
	sort.Strings(got)
	sort.Strings(want)
	if strings.Join(got, ",") != strings.Join(want, ",") {
		t.Fatalf("manifest package names mismatch:\n got: %s\nwant: %s", strings.Join(got, ","), strings.Join(want, ","))
	}
}

func writeSkillPackageFilesWithoutManifest(t *testing.T, targetDir string, pkg SkillPackage) {
	t.Helper()
	for _, file := range pkg.Files {
		target := filepath.Join(targetDir, filepath.FromSlash(file.InstallPath))
		if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
			t.Fatalf("mkdir package file %s: %v", target, err)
		}
		if err := os.WriteFile(target, file.Content, 0o644); err != nil {
			t.Fatalf("write package file %s: %v", target, err)
		}
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
