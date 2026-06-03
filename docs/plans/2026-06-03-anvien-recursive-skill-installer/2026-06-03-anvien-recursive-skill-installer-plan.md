# Anvien Recursive Skill Package Installer Plan

Date: 2026-06-03

Status: Complete

Companion files:

- Evidence ledger: [2026-06-03-anvien-recursive-skill-installer-evidence.md](2026-06-03-anvien-recursive-skill-installer-evidence.md)
- Benchmark ledger: [2026-06-03-anvien-recursive-skill-installer-benchmark.md](2026-06-03-anvien-recursive-skill-installer-benchmark.md)
- Problem report: [2026-06-03-aicontext-skill-script-problems.md](../../../reports/problem/2026-06-03-aicontext-skill-script-problems.md)

## Master Rules

1. Follow active repository instructions and generated `AGENTS.md`.
2. Use Anvien for codebase analysis, impact checks, and implementation-slice evidence.
3. Run `anvien analyze --force` before graph-based Anvien commands.
4. Run impact analysis before editing `internal/aicontext/aicontext.go`, setup/install owners, generated context owners, package/setup integration, or shared skill data contracts.
5. Do not edit generated `AGENTS.md`, `CLAUDE.md`, or `.claude/skills/anvien/**` as source of truth.
6. Keep source of truth under `internal/aicontext/skills/**` and AI-context generator code.
7. Preserve user/repo-local skills that are not owned by Anvien manifest data.
8. Run the full build before tests. For this repo, the full build gate is `powershell -ExecutionPolicy Bypass -File anvien-launcher\build.ps1`.
9. Add e2e only if Web UI behavior changes.
10. Record evidence as each evidenced task completes.
11. Record benchmarkable inventory, size, and package/startup metrics as each benchmarkable task completes.
12. Run `anvien detect-changes --repo Anvien --scope all` before implementation commits.
13. Commit each completed implementation slice after evidence and benchmark ledgers are updated.

## Goal

Replace the current hard-coded/fixed-depth Anvien skill installation model with a recursive, namespace-preserving, incremental **skill package** installer that can handle top-level source packages under `internal/aicontext/skills/`, including nested `SKILL.md` entries inside those packages.

All skills use one model:

```text
skill package root = immediate child directory under internal/aicontext/skills/
package declaration = top-level folder name
skill entries = any SKILL.md files inside that package tree
package payload = every included file under the package root
target package root = .claude/skills/anvien/<package name>
```

A Markdown-only package is still a package; it is just a top-level folder with one `SKILL.md`. A package with child skills, scripts, references, templates, or assets uses the same install/hash/update flow and is copied as one package.

The installer must support packages that contain nested skill entries such as:

```text
internal/aicontext/skills/document-skills/docx/SKILL.md
internal/aicontext/skills/problem-solving/when-stuck/SKILL.md
internal/aicontext/skills/debugging/root-cause-tracing/SKILL.md
```

and install their containing top-level packages without splitting child skill entries into separate managed packages, flattening, collision, duplicate copy, or accidental deletion of repo-local skills.

## Problem

The current AI-context installer is still shaped around a small, fixed embedded base-skill set. That model was acceptable for four source-owned Markdown-only workflow skills, but it does not fit the current skill package tree.

The current source tree now includes:

- many skills;
- nested skills below category folders;
- skills with scripts, references, templates, and assets;
- large skill assets;
- repo-local customization risk in generated skill output.

Problem 8 is specifically that some skill entries are nested deeper than `skills/<name>/SKILL.md`. A fixed-depth scanner or hard-coded registry will miss those entries unless the whole top-level package tree is copied and declared.

The same package model also addresses the script/resource problem: scripts only have value when copied with the package that explains how to use them, and `SKILL.md` entries can only reference `scripts/`, `references/`, `templates/`, or `assets/` reliably if those paths stay relative to the installed package root.

The implementation must also respect the source tree rule:

```text
Top-level folder under internal/aicontext/skills/ is the package boundary.
Package folder must contain at least one SKILL.md somewhere inside it.
Nested SKILL.md files are skill entries/subskills inside that package, not package boundaries.
```

## Scope

In scope:

- `internal/aicontext/aicontext.go` embedded skill package loading, discovery, metadata, install, manifest, and generated guidance.
- Source skills under `internal/aicontext/skills/**`.
- `.claude/skills/anvien/**` generated/install target behavior.
- `anvien analyze` AI-context generation path.
- `anvien setup` editor skill installation path.
- Package fallback source behavior for embedded skill source.
- Tests for recursive discovery, namespace-preserving package install paths, full payload copy, manifest ownership, incremental update behavior, and generated guidance.
- Evidence and benchmark ledgers for skill package inventory, package/install size, build/test results, and detect-changes.

Out of scope:

- Running skill scripts during analyze/setup.
- Changing the behavior of the scripts inside imported skills.
- Web UI behavior changes.
- Flattening nested skill paths.
- Deleting repo-local custom skills not owned by Anvien.
- Solving all script documentation quality issues from the problem report.

## Requirements

1. Treat every immediate child directory under `internal/aicontext/skills/` as a `SkillPackage`; do not create separate install flows for Markdown-only and script/resource packages.
2. Embed or otherwise include the full `internal/aicontext/skills` tree needed by the runtime package installer.
3. Discover package declarations by enumerating top-level source folders, not by a hard-coded Go registry.
4. Discover `SKILL.md` entries recursively inside each package.
5. Preserve the top-level package folder name when installing.
6. Use package install root as the unique manifest key, not any `SKILL.md` frontmatter name.
7. Keep nested entry paths such as `.claude/skills/anvien/document-skills/docx/SKILL.md`.
8. Never split `document-skills/docx` or `debugging/root-cause-tracing` into separate managed packages.
9. Copy and update the whole top-level package payload, not only root or child `SKILL.md` files.
10. Compute a per-package content hash from included relative file paths plus file content.
11. Use a manifest under `.claude/skills/anvien` as the Anvien ownership boundary.
12. Update a managed package only when its package hash changed or managed target files are missing.
13. Skip managed packages when package hash and target files match.
14. Install new source packages and add them to the manifest.
15. Preserve folders/files in `.claude/skills/anvien` that are not recorded in the Anvien manifest.
16. Missing source packages that are still in the manifest must be marked stale or pruned only by explicit policy.
17. The default analyze/setup flow must not delete unmanifested repo-local skills.
18. Analyzer/setup must not execute scripts copied from skills.
19. Generated AI context must explain recursive package paths and script/resource path resolution at a high level.
20. Tests must fail if nested `SKILL.md` entries are missed or flattened.
21. Tests must fail if scripts/resources are omitted from an installed package.
22. Tests must fail if an unmanifested repo-local skill is deleted.
23. Tests must fail if a changed script/resource file does not change the package hash.
24. Hashing must be deterministic across platforms: sort included relative paths, normalize separators to `/`, and hash relative path plus bytes.
25. All discovered top-level source package folders are install candidates; do not keep an allowlist that only installs the four old Anvien workflow skills.
26. A source package whose install root already exists without manifest ownership must not be overwritten. It may be adopted only when the target content is identifiable as Anvien-managed legacy output or exactly matches the source package.
27. User-added files inside an Anvien-managed package directory are preserved unless they are recorded as Anvien-managed files in the manifest.
28. Generated guidance and reports must identify nested skill entries by package-relative path, not only by frontmatter name, so duplicate entry names do not collide.
29. Intended source package folders under `internal/aicontext/skills/` must be included in the implementation commit or explicitly recorded as excluded.
30. Build/test/detect-changes evidence must be recorded before each implementation commit.

## Invariants

1. Source of truth is `internal/aicontext/skills/**`, not generated `.claude/skills/anvien/**`.
2. Top-level folders under `internal/aicontext/skills/` define package roots.
3. All package roots are managed the same way, including one-file Markdown-only packages.
4. `SKILL.md` files inside a package are entry/subskill files, not package boundaries.
5. Package install paths preserve source namespace.
6. Anvien only owns manifest-recorded packages/files.
7. Repo-local custom skills are preserved.
8. Analyze/setup can copy or update script files, but never run them.
9. Manifest package hash is content-addressed and independent of filesystem mtime.
10. Generated output remains disposable only for files owned by Anvien.
11. Implementation preserves current command guidance and retained Anvien workflow skills.
12. Existing target content without manifest ownership is a collision to report, not content to overwrite.

## Technical Direction

Introduce a `SkillPackage` catalog model driven by top-level folders. There is no separate single-file skill model, and there is no hard-coded Go registry of package names.

Expected data shape:

```go
type SkillPackage struct {
    Name        string
    SourceRoot  string
    InstallRoot string
    PackageHash string
    Entries     []SkillEntry
    Files       []SkillPackageFile
}

type SkillEntry struct {
    Name        string
    Description string
    SourcePath  string
    InstallPath string
}

type SkillPackageFile struct {
    SourcePath  string
    InstallPath string
    Hash        string
    SizeBytes   int64
}
```

Example:

```json
{
  "name": "document-skills",
  "sourceRoot": "skills/document-skills",
  "installRoot": "document-skills",
  "packageHash": "sha256:...",
  "entries": [
    {
      "name": "docx",
      "sourcePath": "skills/document-skills/docx/SKILL.md",
      "installPath": "document-skills/docx/SKILL.md"
    },
    {
      "name": "pdf",
      "sourcePath": "skills/document-skills/pdf/SKILL.md",
      "installPath": "document-skills/pdf/SKILL.md"
    }
  ]
}
```

Expected manifest target:

```text
.claude/skills/anvien/.anvien-skill-manifest.json
```

Manifest shape:

```json
{
  "schemaVersion": 1,
  "managedBy": "anvien",
  "skills": {
    "document-skills": {
      "sourceRoot": "skills/document-skills",
      "installRoot": "document-skills",
      "packageHash": "sha256:...",
      "skillEntryCount": 4,
      "fileCount": 130,
      "files": {
        "docx/SKILL.md": "sha256:...",
        "docx/scripts/extract.py": "sha256:...",
        "pdf/SKILL.md": "sha256:..."
      }
    }
  }
}
```

Install flow:

```text
discover top-level source package folders
load manifest

for each source package:
  recursively discover SKILL.md entries inside the package
  compute package hash from included relative paths and bytes

  if manifest has installRoot and package hash same and managed target files exist:
    skip

  if manifest missing installRoot:
    install package payload
    add manifest entry

  if manifest has installRoot but package hash changed or managed files are missing:
    update that package only
    update manifest entry

for each existing folder in .claude/skills/anvien:
  if not in manifest:
    preserve

for each manifest entry missing from source:
  mark stale by default, and only prune under explicit policy
```

Collision, bootstrap, and overlay rules:

```text
if manifest entry exists:
  Anvien may update files recorded in that manifest entry
  Anvien may remove files recorded in the old manifest entry that no longer exist in the source package
  Anvien must preserve extra target files that are not recorded in the manifest entry

if manifest is missing and target package root exists:
  if target payload exactly matches source package payload:
    adopt by writing a manifest entry without rewriting files
  else if target is identifiable as legacy Anvien output from the pre-manifest installer:
    adopt or update according to package hash policy
  else:
    preserve target package root, report an unmanaged collision, and skip overwriting that source package

if manifest is missing and target package root does not exist:
  install source package and record manifest ownership
```

Generated status should expose useful counts:

```text
skill packages: discovered=35 installed=3 updated=1 skipped=31 stale=0 preserved=2 collisions=0 adopted=0
```

## Definition Of Done

1. Package discovery finds every immediate child directory under `internal/aicontext/skills/` as a `SkillPackage`.
2. Entry discovery finds every `SKILL.md` inside each package without turning child entries into separate managed packages.
3. One-file and nested-entry packages install to namespace-preserving top-level package paths.
4. Markdown-only and script/resource packages use the same package install/hash/update flow.
5. Installed packages include all included `SKILL.md` entries plus scripts, references, templates, assets, licenses, and dependency manifests.
6. The installer no longer depends on a hard-coded four-skill embed/read path.
7. The installer no longer removes the whole `.claude/skills/anvien` root on each analyze/setup.
8. Per-package content hash skip/update behavior is implemented and tested.
9. Manifest ownership prevents deletion of unmanifested repo-local skills.
10. Stale managed packages are reported or handled only by explicit policy.
11. Generated AI context documents top-level package location, nested entry location, package-relative entry identifiers, and path resolution.
12. Full build, focused tests, inventory benchmarks, and detect-changes pass.
13. Each implementation slice is committed after evidence and benchmark ledgers are updated.

## Phase Checklist

- [x] [P0-A] Establish baseline and owner evidence.
  - Goal: capture the current skill tree shape and AI-context installer owner surfaces before implementation.
  - Work Steps: run `anvien analyze --force`; query Anvien for AI-context skill installer ownership; inspect `InstallBaseSkillsTo`; count top-level package folders, nested `SKILL.md` entry files, entry depth, and package payload inventories.
  - Implementation Gate: no code edits in this phase.
  - Acceptance: evidence ledger records Anvien owner evidence and baseline skill inventory; benchmark ledger records current graph and skill counts.

- [x] [P1-A] Add top-level skill package catalog model.
  - Goal: replace fixed-depth/hard-coded source skill declaration with a `SkillPackage` catalog discovered from top-level folders under `internal/aicontext/skills/`.
  - Work Steps: run impact on `BaseSkillFiles`, `baseSkillContent`, `InstallBaseSkillsTo`, `installBaseSkills`, and `renderAnvienBlock`; replace the four-entry `baseSkills` registry as source of truth; introduce a `SkillPackage` model with package root, entry list, and payload file fields; walk embedded `skills` top-level directories; recursively parse frontmatter from all package `SKILL.md` entries; derive `InstallRoot` from the top-level package folder; keep existing four Anvien workflow skills as one-entry packages discovered from folders.
  - Implementation Gate: do not change removal/update semantics in this phase; only change discovery/model.
  - Acceptance: tests prove top-level package folders are discovered, nested `SKILL.md` entries are recorded as package entries, and child entries are not declared as independent managed packages.

- [x] [P1-B] Add package tree validation.
  - Goal: make source package shape explicit so future source trees do not create empty or ambiguous package output.
  - Work Steps: validate that each top-level package folder contains at least one `SKILL.md` entry somewhere inside it; allow multiple nested `SKILL.md` entries in the same package; validate deterministic package and entry paths; test one-entry packages, multi-entry packages such as `debugging`, and an empty top-level package folder.
  - Implementation Gate: validation must not reject the current source tree, where `debugging`, `document-skills`, and `problem-solving` are valid multi-entry packages.
  - Acceptance: tests fail for a top-level source package folder with zero `SKILL.md` entries and pass for top-level packages containing nested child `SKILL.md` entries.

- [x] [P1-C] Wire skill package catalog into AI-context generation.
  - Goal: make `aicontext.go` use the discovered `SkillPackage` catalog as the data source for generated `AGENTS.md`/`CLAUDE.md` skill guidance instead of the old four-skill registry.
  - Work Steps: run impact on `renderAnvienBlock`, `GenerateAIContextFiles`, `Generate`, `Result`, and any tests that assert `BaseSkillIDs`; thread the discovered package catalog into the render path; replace Skill Selection Guide rows derived from `baseSkills` with rows derived from package catalog entries; include package-relative entry paths for nested `SKILL.md` files; keep command guidance unchanged; rename report/result fields if needed so they no longer imply only four base skills.
  - Implementation Gate: this phase must not depend on installed `.claude/skills/anvien/**` output as source of truth; generated guidance must come from `internal/aicontext/skills/**` catalog data.
  - Acceptance: generator tests prove `AGENTS.md`/`CLAUDE.md` list or summarize discovered top-level packages, include nested entry paths for multi-entry packages such as `debugging`, and no longer depend on a hard-coded four-skill `baseSkills` list.

- [x] [P2-A] Introduce manifest ownership and package hash model.
  - Goal: make installed package updates incremental and bounded by Anvien ownership.
  - Work Steps: define `.anvien-skill-manifest.json` schema; compute per-package SHA-256 from included relative file paths and file bytes; store per-file hashes for managed files; load/write manifest atomically; add tests for same-hash skip, changed-hash update, new-source install, and manifest update.
  - Implementation Gate: manifest must use `installRoot` as key and must not rely on frontmatter name uniqueness.
  - Acceptance: unchanged managed packages are skipped, changed managed packages update, new source packages install, and manifest entries record package hash/skill entry count/file count/source/install roots.

- [x] [P2-B] Preserve repo-local packages and handle target collisions.
  - Goal: prevent analyze/setup from deleting or overwriting user or repo-specific skill folders that are not owned by Anvien.
  - Work Steps: replace whole-root `RemoveAll(.claude/skills/anvien)` behavior with per-managed-package updates; create tests with an unmanifested `.claude/skills/anvien/my-repo-custom-skill/SKILL.md`; create tests where a source package install root collides with an unmanifested target package; verify both survive install/update; preserve user-added files inside managed package directories; define stale managed package default behavior as report/mark, not auto-delete.
  - Implementation Gate: no delete path can target a folder/file unless it is in the Anvien manifest and the prune policy explicitly allows it.
  - Acceptance: tests prove unmanifested repo-local packages, unmanaged collisions, and unmanifested overlay files inside managed packages survive analyze/setup package installation.

- [x] [P2-C] Bootstrap legacy installs into manifest ownership.
  - Goal: make existing pre-manifest Anvien installs upgrade cleanly without overwriting unrelated local skills.
  - Work Steps: handle missing-manifest first run; adopt target packages that exactly match source payloads; adopt or update known legacy Anvien workflow skill outputs when identifiable; report unmanaged collisions for same install roots that cannot be proven Anvien-owned; test each branch.
  - Implementation Gate: bootstrap must prefer preserve/report over overwrite when ownership is uncertain.
  - Acceptance: legacy Anvien-installed skills become manifest-owned, exact matches are adopted without rewrites, and unknown same-path targets are preserved with a collision report.

- [x] [P3-A] Sync full skill package payload with artifact policy.
  - Goal: copy each package root's required `SKILL.md`, scripts, references, templates, assets, licenses, and dependency manifests without copying obvious generated artifacts.
  - Work Steps: define payload traversal; include source files under each top-level package root; exclude known generated/cache artifacts such as `.coverage`, `__pycache__`, `.pytest_cache`, `node_modules`, `dist`, and `build`; test with `ui-styling` and the multi-entry `document-skills` package scripts/references and artifact exclusions.
  - Implementation Gate: script files are copied but never executed during analyze/setup.
  - Acceptance: installed package payload includes expected script/reference/template files and excludes known generated artifacts.

- [x] [P3-B] Update generated AI-context guidance.
  - Goal: teach agents where recursive skill packages are installed and how to resolve script/reference paths without listing every skill inline.
  - Work Steps: update generated `AGENTS.md`/`CLAUDE.md` Anvien block source to state that Anvien skills live as top-level packages under `.claude/skills/anvien/<package>/`; state that `SKILL.md` files may be nested inside a package; state that all discovered top-level packages are eligible, not only the four old workflow skills; state that scripts/references/templates/assets resolve relative to each package directory; preserve direct Anvien command table guidance.
  - Implementation Gate: do not manually edit generated root files as source.
  - Acceptance: generator tests prove the recursive skill guidance is present, nested entries are identified by package-relative path, and command guidance remains command-first.

- [x] [P4-A] Integrate analyze/setup reporting and package fallback.
  - Goal: make analyze/setup report skill package install activity and ensure package/fallback source includes recursive skill payloads.
  - Work Steps: thread install result counts through analyze/setup reporting; update package source tests from flat `skills/*.md` assumptions to recursive package payload assumptions; record package/source size benchmarks.
  - Implementation Gate: package changes must not silently omit nested package payloads.
  - Acceptance: analyze/setup can report discovered/installed/updated/skipped/stale/preserved counts; package tests prove recursive package sources are included.

- [x] [P5-A] Validate full implementation and commit.
  - Goal: close the recursive installer implementation slice with build, tests, detect-changes, evidence, benchmarks, and commit.
  - Work Steps: verify intended source package folders under `internal/aicontext/skills/` are tracked or explicitly excluded; run full build; run focused Go tests for `internal/aicontext`, `internal/cli`, `internal/mcp`, and package tests touched by the implementation; run any new installer tests; refresh graph; run `anvien detect-changes --repo Anvien --scope all`; update evidence and benchmark ledgers; commit.
  - Implementation Gate: no commit before full build, tests, detect-changes, evidence, and benchmark entries are complete.
  - Acceptance: commit exists for the completed implementation slice and final evidence identifies any residual risk.
