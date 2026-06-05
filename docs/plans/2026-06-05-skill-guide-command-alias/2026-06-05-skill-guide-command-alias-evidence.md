# Evidence Ledger

Title: Skill Guide Command Alias Column

Date: 2026-06-05

Status: Planned

Companion files:

- Plan: [2026-06-05-skill-guide-command-alias-plan.md](2026-06-05-skill-guide-command-alias-plan.md)
- Benchmark ledger: [2026-06-05-skill-guide-command-alias-benchmark.md](2026-06-05-skill-guide-command-alias-benchmark.md)

## Evidence Rules

1. This file records planning and future implementation evidence.
2. Generated files are validation artifacts, not permanent source of truth.
3. Tests are validation evidence only, not source of truth for desired behavior.
4. Impact and change-detection evidence must be recorded before implementation commit.
5. Benchmarkable inventory counts belong in the benchmark ledger.

## E0 - User Requirement

User requested a concrete implementation plan for adding a generated command alias column to the generated `Skill Selection Guide`.

Required target shape:

```text
When you need to... | command | Use
use when user ask to review spec | /architect-review | .agents/skills/Architect-review/SKILL.md
```

Required source-of-truth rule:

- command alias must be generated from the same skill source data as the row;
- if a skill disappears from `internal/aicontext/skills/**`, the alias disappears;
- if a skill `name` changes, the alias changes;
- no manual alias registry.

## E1 - Codebase Facts

Source inspection:

- `internal/aicontext/aicontext.go` calls `renderAnvienBlock(".agents/skills/", packages)` for `AGENTS.md`.
- `internal/aicontext/aicontext.go` calls `renderAnvienBlock(".claude/skills/", packages)` for `CLAUDE.md`.
- `renderAnvienBlock` currently writes:

```text
| When you need to... | Use |
|---------------------|-----|
```

- `renderAnvienBlock` currently formats rows with `skillGuideNeed(pkg)` and `skillGuideUse(pkg, skillPathPrefix)`.
- `internal/aicontext/skill_packages.go` defines `primarySkillEntry(pkg)`.
- `skillGuideUse(pkg, prefix)` already uses `primarySkillEntry(pkg).InstallPath`.
- `readSkillPackage` sets package description from the primary entry description.
- `internal/aicontext/aicontext_test.go` already asserts generated `Skill Selection Guide` shape and primary-only `problem-solving` behavior.

## E2 - Planning Commands

Commands run during planning:

```text
anvien analyze --force
Select-String internal/aicontext/aicontext.go internal/aicontext/skill_packages.go internal/aicontext/aicontext_test.go -Pattern "Skill Selection Guide|skillGuideNeed|skillGuideUse|primarySkillEntry|renderAnvienBlock|SkillPackageCatalog"
anvien impact symbol "renderAnvienBlock" --repo Anvien --direction upstream
anvien impact symbol "skillGuideNeed" --repo Anvien --direction upstream
anvien impact symbol "skillGuideUse" --repo Anvien --direction upstream
anvien context symbol "primarySkillEntry" --repo Anvien
```

Graph refresh result:

- files scanned: 1382
- parsed code files: 682
- parse failures: 0
- graph nodes: 84142
- graph relationships: 122608

## E3 - Impact Evidence

Planning impact results:

| Target | Risk | Affected files/processes |
|---|---|---|
| `renderAnvienBlock` | CRITICAL | affects `internal/aicontext/aicontext.go`, `internal/cli/analyze_postrun.go`, `internal/cli/command.go`, generated analyze flows |
| `skillGuideNeed` | HIGH | affects generated skill guide through `renderAnvienBlock`, analyze flows, focused aicontext tests |
| `skillGuideUse` | HIGH | affects generated skill guide through `renderAnvienBlock`, analyze flows, focused aicontext tests |
| `primarySkillEntry` | central source | used by `readSkillPackage`, `skillGuideUse`, and existing generated guide tests |

Interpretation:

- HIGH/CRITICAL blast radius is a scope warning, not a prohibition.
- Implementation must stay scoped to generated guide rendering and helpers.
- Tests must cover generated `AGENTS.md` and `CLAUDE.md` behavior because analyze flows consume the rendered block.

## E4 - Current Generated Output Evidence

Current generated `AGENTS.md` `Skill Selection Guide` sample:

```text
## Skill Selection Guide

AI agent chooses the skill that fits the work.

| When you need to... | Use |
|---------------------|-----|
| use when user ask to review spec | `.agents/skills/Architect-review/SKILL.md` |
```

Observed behavior:

- two-column table exists today;
- `Use` path already points to `.agents/skills/...` in `AGENTS.md`;
- no command alias column exists.

## E5 - Worktree Note

Planning was written against the current working tree. Existing unrelated working tree changes were present before this plan work:

```text
D coder.md
M internal/aicontext/skills/coder/SKILL.md
```

This plan does not modify those files.

## E6 - Future Implementation Evidence Template

Append implementation evidence under this section as phases complete.

Required evidence entries:

- P0-A impact evidence completed.
- P1-A helper implementation files and tests.
- P2-A generated table render change.
- P3-A focused tests added/updated.
- P4-A build/test/analyze/detect-changes output.
- Commit hash for each completed implementation slice.

## E7 - Implementation Evidence

Implementation files changed:

- `README.md`
- `internal/aicontext/aicontext.go`
- `internal/aicontext/skill_packages.go`
- `internal/aicontext/aicontext_test.go`
- `internal/version/version.go`

Implementation summary:

- Added generated `Command` column to `Skill Selection Guide`.
- Added `skillGuideCommand(pkg SkillPackage)` and deterministic command-name normalization.
- Kept command aliases surface-independent; only the `Use` path prefix differs between AGENTS and CLAUDE.
- Added tests for command normalization, `/architect-review`, `/problem-solving`, and nested child command exclusion.
- Updated README skill documentation so it describes generated skill command aliases separately from Anvien CLI command selection.
- Corrected Go CLI source version from `1.2.4` to `1.2.5` so validation output matches `anvien/package.json`.

Impact evidence during implementation:

| Target | Result |
|---|---|
| `renderAnvienBlock` | CRITICAL blast-radius warning; scoped to generated guide table rendering |
| `skillGuideNeed` | HIGH blast-radius warning; behavior unchanged |
| `skillGuideUse` | HIGH blast-radius warning; behavior unchanged except table placement |
| `primarySkillEntry` | LOW direct impact; reused as command alias source |
| `internal/version/version.go:Version` | LOW impact; linked CLI/MCP version tests |

## E8 - Validation Evidence

Full build sequence requested by user:

```text
cd .\anvien
npm install
npm run build
npm install -g .
Get-Command anvien
anvien version
cd ..
powershell -ExecutionPolicy Bypass -File .\anvien-launcher\build.ps1
anvien version
anvien analyze . --force
```

Results:

| Command | Result |
|---|---|
| `npm install` in `anvien` | Passed after stopping stale Anvien MCP process that held `lbug_shared.dll` |
| `npm run build` in `anvien` | Passed; rebuilt `anvien\bin\anvien.exe` |
| `npm install -g .` in `anvien` | Passed |
| `Get-Command anvien` | Resolved global `anvien.ps1` |
| `anvien version` in `anvien` | `1.2.5` |
| `powershell -ExecutionPolicy Bypass -File .\anvien-launcher\build.ps1` | Passed; Vite reported existing chunk-size warnings only |
| `anvien version` at repo root | `1.2.5` |
| `anvien analyze . --force` | Passed: files scanned 1385, parsed code 682, failed 0, graph nodes 84214, relationships 122686 |

Tests:

| Command | Result |
|---|---|
| `go test ./internal/version ./internal/cli ./internal/mcp -count=1` | Passed |
| `go test ./internal/aicontext -count=1` | Passed |
| `go test ./cmd/... ./internal/... -count=1` | Passed |

Generated output inspection:

| File | Result |
|---|---|
| `AGENTS.md` | Contains `| When you need to... | Command | Use |` |
| `CLAUDE.md` | Contains `| When you need to... | Command | Use |` |

Change detection:

```text
anvien detect-changes --repo Anvien --scope all
```

Result summary:

- changed files: 6
- affected files: 6
- risk level: low
- version source file risk: low
- aicontext renderer/helper files appear as high file-risk because they are generated-output surfaces; implementation remained scoped.
