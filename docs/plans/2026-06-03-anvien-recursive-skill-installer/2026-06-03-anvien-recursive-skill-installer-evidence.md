# Evidence Ledger

Date: 2026-06-03

Status: Complete

Companion files:

- Plan: [2026-06-03-anvien-recursive-skill-installer-plan.md](2026-06-03-anvien-recursive-skill-installer-plan.md)
- Benchmark ledger: [2026-06-03-anvien-recursive-skill-installer-benchmark.md](2026-06-03-anvien-recursive-skill-installer-benchmark.md)
- Problem report: [2026-06-03-aicontext-skill-script-problems.md](../../../reports/problem/2026-06-03-aicontext-skill-script-problems.md)

## Evidence Rules

1. Record source facts and command evidence as phases complete.
2. Use Anvien for codebase ownership, impact, graph, and detect-changes evidence.
3. Keep validation evidence here; measured inventory/size/performance metrics belong in the benchmark ledger.
4. Do not treat generated `.claude/skills/anvien/**`, `AGENTS.md`, or `CLAUDE.md` edits as source-of-truth evidence unless they were produced by normal generation.
5. Record failures and residual risks instead of hiding them.

## E0 - Planning Baseline

Status: Complete

### User Decision

The user accepted the combined approach for problem 8:

```text
recursive discovery
preserve namespace path
per-package hash
incremental install
manifest ownership
preserve repo-local skills
```

The user then tightened the design decision: all skills must be treated as packages, with no separate flow for Markdown-only skills versus script/resource skills.

Accepted package rule:

```text
skill package root = immediate child directory under internal/aicontext/skills/
package declaration = top-level folder name
skill entries = any SKILL.md files inside that package tree
package payload = every included file under the package root
target package root = .claude/skills/anvien/<package name>
```

This means a simple top-level folder is a one-entry package, while a folder such as `debugging` can be a multi-entry package containing several child `SKILL.md` files. Those child skill entries are package contents, not independent package boundaries.

The user also clarified that package discovery must be connected back into `aicontext.go` generation itself. The catalog is not only an installer input; it must drive generated `AGENTS.md`/`CLAUDE.md` skill guidance instead of the old four-skill `baseSkills` registry.

The user explicitly rejected deleting all target skills that are missing from source. The accepted ownership rule is:

```text
Anvien can update/delete only skills it owns through the Anvien manifest.
Skills outside the Anvien manifest are repo-local/user-local and must be preserved.
```

### Problem Evidence

The related problem report is committed at:

```text
reports/problem/2026-06-03-aicontext-skill-script-problems.md
```

Relevant original problem entries for this plan:

```text
8. Some skills are nested deeper than `skills/<name>/SKILL.md`.
9. Some parent skill folders and child skill folders both contain `SKILL.md`, which creates ambiguity.
```

The clarified package model resolves problem 8 by copying the whole top-level package tree. Nested `SKILL.md` files can be as deep as needed inside the package and still travel with their scripts/resources.

The package model also covers the related script/resource problem: copied scripts only remain usable if they travel with the top-level package that explains when and how to run them, and if relative paths such as `scripts/`, `references/`, `templates/`, and `assets/` remain valid from the installed package root.

The earlier parent/child ambiguity is no longer treated as a package-boundary problem. Under the clarified model, parent/child `SKILL.md` files can coexist inside a single top-level package. A filesystem scan still found no old-style ambiguous package roots:

```text
ambiguousParentSkills=0
```

### Anvien Freshness Evidence

Command:

```powershell
anvien analyze --force
```

Result:

```text
analyzed E:\Anvien
files: scanned=1352 parsed_code=702 failed=0
indexed: documents=405 metadata=114 analyzers=0 scripts=7 static=3
graph: nodes=83837 relationships=122360
fileProjection: status=built files=1352 dependencyEdges=16701 unresolved=448 hotspots=5
```

The refreshed graph now includes skill script files as first-class analyzed files. Examples from the analyze output include:

```text
internal/aicontext/skills/document-skills/pptx/scripts/html2pptx.js
internal/aicontext/skills/document-skills/docx/scripts/document.py
internal/aicontext/skills/ui-styling/scripts/tailwind_config_gen.py
internal/aicontext/skills/payment-integration/scripts/checkout-helper.js
```

### Anvien Ownership Evidence

Command:

```powershell
anvien query "aicontext skill installer BaseSkillFiles InstallBaseSkillsTo recursive skills" --repo Anvien
```

Relevant result:

```text
rank 1 definition: internal/aicontext/aicontext.go:installBaseSkills
rank 3 file: internal/aicontext/aicontext.go
matched symbols include BaseSkillFiles, GenerateAIContextFiles, InstallBaseSkillsTo, activeBaseSkillNames, baseSkillContent
```

Command:

```powershell
anvien context symbol "InstallBaseSkillsTo" --repo Anvien
```

Relevant result:

```text
symbol: Function internal/aicontext/aicontext.go:InstallBaseSkillsTo
linked flows: 2
linked tests: 2
inbound files include internal/aicontext/aicontext_test.go, internal/cli/command_test.go, internal/cli/setup_command.go
next suggested impact target: InstallBaseSkillsTo
```

### Source Tree Evidence

Filesystem scan found 35 top-level package folders under `internal/aicontext/skills`.

Filesystem scan also found 48 `SKILL.md` files under those packages. These are skill entries/subskills, not 48 independent packages.

Maximum observed skill entry path depth within a package is currently 2, with examples:

```text
document-skills/docx/SKILL.md
document-skills/pdf/SKILL.md
document-skills/pptx/SKILL.md
document-skills/xlsx/SKILL.md
problem-solving/when-stuck/SKILL.md
debugging/root-cause-tracing/SKILL.md
```

Multi-entry package examples:

```text
debugging: 5 skill entries
document-skills: 4 skill entries
problem-solving: 7 skill entries
```

Package payload inventory from the same source tree:

```text
packages=35
skillEntries=48
packagesWithMultipleSkillEntries=3
packagesWithScripts=16
totalPackageFiles=607
totalPackageBytes=11311115
totalScriptFiles=150
maxSkillEntryDepthWithinPackage=2
```

Script-heavy package examples:

```text
chrome-devtools: files=26 scriptFiles=21
document-skills: files=130 scriptFiles=37
databases: files=19 scriptFiles=10
web-frameworks: files=18 scriptFiles=9
ui-styling: files=98 scriptFiles=8
```

### Dirty Worktree Note

At planning time, the worktree contains untracked imported skill directories under `internal/aicontext/skills/**`. There are also unrelated report/image changes in `reports/problem`. This plan creation must not revert or stage unrelated user changes.

## E1 - Implementation Evidence

Status: Complete

### Impact Evidence

Commands run before editing AI-context/setup owners:

```powershell
anvien impact symbol "BaseSkillFiles" --repo Anvien --direction upstream
anvien impact symbol "InstallBaseSkillsTo" --repo Anvien --direction upstream
anvien impact symbol "GenerateAIContextFiles" --repo Anvien --direction upstream
anvien impact symbol "renderAnvienBlock" --repo Anvien --direction upstream
anvien impact file internal/aicontext/aicontext_test.go --repo Anvien --direction upstream
anvien impact file internal/cli/command_test.go --repo Anvien --direction upstream
anvien impact file internal/cli/setup_command.go --repo Anvien --direction upstream
anvien impact file internal/cli/analyze_postrun.go --repo Anvien --direction upstream
```

Relevant results:

```text
BaseSkillFiles: LOW
InstallBaseSkillsTo: HIGH; affected files internal/aicontext/aicontext.go and internal/cli/setup_command.go
GenerateAIContextFiles: CRITICAL; affected analyze AI-context generation path
renderAnvienBlock: CRITICAL; affected generated AGENTS.md/CLAUDE.md guidance
internal/aicontext/aicontext_test.go: LOW
internal/cli/command_test.go: LOW
internal/cli/setup_command.go: MEDIUM file-level blast radius
internal/cli/analyze_postrun.go: HIGH file-level blast radius, linked to NewAnalyzeCommand flows
```

The HIGH/CRITICAL blast radius is expected because this slice changes generated AI-context and setup/analyze skill installation. No Web UI files were changed, so no e2e test is required by the repository rule.

### Implementation Summary

Changed source files:

```text
internal/aicontext/aicontext.go
internal/aicontext/skill_packages.go
internal/aicontext/aicontext_test.go
internal/cli/setup_command.go
internal/cli/command_test.go
docs/plans/2026-06-03-anvien-recursive-skill-installer/*
```

Behavior implemented:

```text
Skill package root = immediate child of internal/aicontext/skills/
Skill entries = every SKILL.md under that top-level package
Install root = .claude/skills/anvien/<package>
Manifest = .claude/skills/anvien/.anvien-skill-manifest.json
Hash = sha256 over sorted package-relative file paths and per-file hashes
```

`aicontext.go` now discovers packages from the embedded `skills` tree before rendering generated `AGENTS.md`/`CLAUDE.md`. The generated Skill Selection Guide is no longer sourced from a four-item `baseSkills` registry; it lists discovered package roots and full installed entry paths.

`skill_packages.go` now owns recursive discovery, frontmatter parsing, fallback Markdown description derivation for entries without YAML frontmatter, deterministic package hashing, manifest load/write, incremental install, legacy adoption, and collision/preservation rules.

Setup now calls the same package installer through `aicontext.InstallSkillPackagesTo` and reports detailed counts with `SkillInstallResult.Summary()`. The old `InstallBaseSkillsTo` wrapper remains for compatibility and returns package IDs.

### Ownership And Preservation Evidence

Implemented manifest boundary:

```json
{
  "schemaVersion": 1,
  "managedBy": "anvien",
  "skills": {
    "ui-styling": {
      "installPath": "ui-styling",
      "sourceRoot": "skills/ui-styling",
      "hash": "sha256:...",
      "managed": true,
      "entryCount": 1,
      "fileCount": 98,
      "files": {
        "SKILL.md": "sha256:..."
      }
    }
  }
}
```

Preservation behavior covered by tests:

```text
unmanifested repo-local skill folder is preserved
unmanifested same-name collision is rejected and preserved
user-added file inside a managed package is preserved
manifest-owned missing source package is marked stale, not deleted
legacy four Anvien workflow packages can be adopted from pre-manifest output
```

### Payload Evidence

Tests verify full package payloads, not only Markdown entries:

```text
debugging/root-cause-tracing/find-polluter.sh
document-skills/docx/scripts/document.py
ui-styling/scripts/shadcn_add.py
ui-styling/canvas-fonts/ArsenalSC-Regular.ttf
```

The installer copies scripts/assets as files only; analyze/setup do not execute skill scripts.

Package payload is copied as a whole folder. Dotfiles and package-local artifacts are included and hashed like any other package file; the installer does not classify package children by name.

```text
payloadFiles=607
payloadBytes=11311115
dotfileCoverageArtifacts=10
```

### Validation Evidence

Focused validation run:

```powershell
go test ./internal/aicontext
go test ./internal/aicontext ./internal/cli
```

Results:

```text
ok github.com/tamnguyendinh/anvien/internal/aicontext
ok github.com/tamnguyendinh/anvien/internal/cli
```

Full build gate:

```powershell
.\anvien-launcher\build.ps1
```

Result:

```text
pass
Go: go1.26.3 windows/amd64
Web: tsc -b && vite build completed
Notes: Vite reported existing chunk-size/dynamic-import warnings; build exit code was 0.
```

Broad Go test attempt:

```powershell
go test ./...
```

Result:

```text
fail outside this slice
```

Failure causes were unrelated fixture/baseline packages that are not valid normal build targets:

```text
anvien/test/fixtures/lang-resolution/go-map-range imports package models as a local fixture
anvien/test/fixtures/lang-resolution/go-method-enrichment mixes fixture packages/imports
anvien/test/fixtures/sample-code contains C source files without cgo
anvien/test/fixtures/lang-resolution/go-make-builtin has an intentional pointer-method compile error fixture
anvien/test/fixtures/lang-resolution/go-type-assertion has an intentional impossible assertion fixture
internal/lbugschema expects missing baseline/phase-1-contract-freeze/ladybugdb-graph-contract.json
```

Clean focused post-build validation:

```powershell
go test ./internal/aicontext ./internal/cli ./internal/mcp ./internal/gitignore ./internal/contracts
```

Result:

```text
ok github.com/tamnguyendinh/anvien/internal/aicontext
ok github.com/tamnguyendinh/anvien/internal/cli
ok github.com/tamnguyendinh/anvien/internal/mcp
ok github.com/tamnguyendinh/anvien/internal/gitignore
ok github.com/tamnguyendinh/anvien/internal/contracts
```

Observed and handled failure during implementation:

```text
google-adk-python/SKILL.md had no YAML frontmatter.
Resolution: discovery now derives entry description from Markdown content when frontmatter description is absent, so all source packages remain usable instead of being filtered out.
```

### Remaining Validation

Post-change Anvien refresh:

```powershell
anvien analyze --force
```

Result:

```text
analyzed E:\Anvien
files: scanned=1353 parsed_code=703 failed=0
indexed: documents=405 metadata=114 analyzers=0 scripts=7 static=3
graph: nodes=84170 relationships=122898
fileProjection: files=1353 dependencyEdges=16728 unresolved=449 hotspots=5
```

Detect changes:

```powershell
anvien detect-changes --repo Anvien --scope all
```

Relevant summary:

```text
changed_files=543
changed_count=22192
affected_files=384
affected_count=25
risk_level=critical
changed app layers: backend=14944, backend_test=49, docs=7199
affected app layer: backend=25
changed functional areas: cli=12, documentation=7199, unknown=14981
affected functional areas: cli=2, unknown=23
resolution gap changes: changedGapEntities=11655, changedGapOccurrenceCount=11665
```

High/critical blast-radius drivers reported by detect-changes:

```text
internal/aicontext/aicontext.go
internal/cli/setup_command.go
new internal/aicontext/skills/** package payload files
```

This matches the intended blast radius: generated AI-context output, setup skill installation, and the newly added package payload tree. The critical risk level is expected because this slice intentionally stages hundreds of skill payload files, including scripts and references, and the analyzer reports unresolved sites inside those copied payloads.

## E2 - Whole Package Payload Correction

Status: Complete

The package copy rule was corrected after review: a skill package is the whole top-level folder under `internal/aicontext/skills/`, and Anvien must not classify or exclude any package child path.

Implementation facts:

```text
source package boundary: internal/aicontext/skills/<package>/
embedded source: //go:embed all:skills
package payload: every regular file under the package root
package hash input: sorted package-relative file paths and file hashes
generated target: .claude/skills/anvien/<package>/
```

Validation added:

```text
ui-styling/scripts/.coverage is cataloged as package payload
ui-styling/scripts/.coverage is installed into .claude/skills/anvien/ui-styling/scripts/.coverage
manifest file counts include dotfile payloads
```

Measured source payload after including all package files:

```text
payloadFiles=607
payloadBytes=11311115
dotfileCoverageArtifacts=10
scriptFiles=150
```

Commands:

```powershell
.\anvien-launcher\build.ps1
go test ./internal/aicontext
go test ./internal/aicontext ./internal/cli ./internal/mcp ./internal/gitignore ./internal/contracts
anvien analyze --force
anvien detect-changes --repo Anvien --scope all
```

Results:

```text
full build: pass
focused tests: pass
analyze: files=1353 parsed_code=703 failed=0 nodes=84166 relationships=122890 dependencyEdges=16727 unresolved=449 hotspots=5
detect-changes: changed_files=8 changed_count=12 affected_files=4 affected_count=0 risk_level=low
```

Commit:

```text
Created after this ledger update; final response records the commit hash.
```
