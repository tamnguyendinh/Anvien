# Anvien Skill Selection Guide Reduction Evidence Ledger

Date: 2026-06-02

Status: Complete

Companion files:

- Plan: [2026-06-02-anvien-skill-selection-guide-reduction-plan.md](2026-06-02-anvien-skill-selection-guide-reduction-plan.md)
- Benchmark ledger: [2026-06-02-anvien-skill-selection-guide-reduction-benchmark.md](2026-06-02-anvien-skill-selection-guide-reduction-benchmark.md)

## Evidence Rules

1. Record Anvien command evidence for implementation slices.
2. Do not use Anvien for doc-only planning commits.
3. Keep quantitative inventory counts in the benchmark ledger.
4. Record impact/blast-radius before editing generator functions or retained workflow owners.
5. Record generated output checks after regeneration.
6. Record `anvien detect-changes --repo Anvien --scope all` before implementation commits.

## Evidence Template

Use this template for implementation phases:

```text
## E<n> - <Phase/Task>

Date:

Status:

Scope:

- ...

Source / command evidence:

| Check | Result |
|---|---|
| ... | ... |

Impact / blast radius:

| Target | Result |
|---|---|
| ... | ... |

Implementation evidence:

| File | Evidence |
|---|---|
| ... | ... |

Validation:

| Command | Result |
|---|---|
| ... | ... |

Failures / handling:

- ...

Detect changes:

| Command | Result |
|---|---|
| `anvien detect-changes --repo Anvien --scope all` | ... |

Commit:

- `<hash> <subject>`
```

## E0 - Plan Structure Review

Date: 2026-06-02

Status: recorded

Scope:

- Created plan, evidence ledger, and benchmark ledger using the existing `docs/plans` structure.
- No implementation source files changed in this planning step.
- No Anvien command was used because this is doc-only planning.

Source / command evidence:

| Check | Result |
|---|---|
| Reviewed `docs/plans/2026-06-01-anvien-analyze-file-classification-metrics-plan.md` | Confirmed plan convention: Date, Status, companion ledgers, Master Rules, Goal, Problem, Scope, Requirements, Invariants, Technical Direction, Definition Of Done, Phase Checklist, Risk Notes. |
| Reviewed `docs/plans/2026-05-23-anvien-skill-system-upgrade-plan.md` | Confirmed generated AI-context plans treat generated files as validation artifacts and source files as `internal/aicontext/aicontext.go` plus embedded Markdown. |
| `Get-ChildItem internal\aicontext\skills -Filter 'anvien-*.md'` | Current embedded Anvien skill inventory has 10 files. |
| `Get-ChildItem .claude\skills\anvien` | Current generated Anvien skill inventory has 10 directories. |

Impact / blast radius:

| Target | Result |
|---|---|
| Implementation symbols | Not run; no implementation edits in this planning step. |

Validation:

| Command | Result |
|---|---|
| `git status --short` | Shows only `docs/plans` doc-only planning files changed for this work. |

Failures / handling:

- None.

## E1 - P0-A Baseline And Owner Evidence

Date: 2026-06-02

Status: recorded

Scope:

- Established implementation owners before editing AI-context generation.
- Recorded current source and generated skill inventory.
- Recorded blast radius for the generator functions and source file that will change.

Source / command evidence:

| Check | Result |
|---|---|
| `anvien analyze --force` | Refreshed graph for `E:\Anvien`; scanned 821 files, parsed 598 code files, failed 0; graph path `.anvien/graph.json`; 96,254 nodes and 131,727 relationships. |
| `anvien query "AI context generated skills AGENTS CLAUDE baseSkills Skill Selection Guide" --repo E:\Anvien` | Found AI-context owners in `internal/aicontext/aicontext.go`: `GenerateAIContextFiles`, `renderAnvienBlock`, `installBaseSkills`, `BaseSkillFiles`, and `baseSkillContent`; tests in `internal/aicontext/aicontext_test.go` and CLI setup/analyze tests. |
| `rg --files internal\aicontext\skills .claude\skills\anvien` | Current embedded source inventory has 10 `internal/aicontext/skills/anvien-*.md` files; current generated inventory has 10 `.claude/skills/anvien/*/SKILL.md` files. |
| Source inspection | `internal/aicontext/aicontext.go` owns `baseSkills`, generated `Command Selection Guide`, generated root `## Skills` table, and `.claude/skills/anvien/**` installation through embedded Markdown. |

Impact / blast radius:

| Target | Result |
|---|---|
| `renderAnvienBlock` | CRITICAL; affects `internal/aicontext/aicontext.go`, `internal/cli/analyze_postrun.go`, and `internal/cli/command.go`; linked flows include `NewAnalyzeCommand -> RenderAnvienBlock`; linked tests include `internal/aicontext/aicontext_test.go` and `internal/cli/command_test.go`. |
| `GenerateAIContextFiles` | CRITICAL; affects analyze AI-context generation through `generateAnalyzeAIContext` and `newAnalyzeCommand`; linked tests include `internal/aicontext/aicontext_test.go` and `internal/cli/command_test.go`. |
| `BaseSkillFiles` | HIGH; affects `InstallBaseSkillsTo`, `installBaseSkills`, and setup skill installation through `internal/cli/setup_command.go`. |
| `internal/aicontext/aicontext.go` | File-level impact touches AI-context generation, setup skill installation, and CLI analyze flows; proceed with narrow edits only. |

Implementation evidence:

| File | Evidence |
|---|---|
| `docs/plans/2026-06-02-anvien-skill-selection-guide-reduction-plan.md` | Marked P0-A complete after owner and baseline evidence were recorded. |

Validation:

| Command | Result |
|---|---|
| `git status --short` | Clean before implementation source edits except the P0-A ledger update. |

Failures / handling:

- Initial impact commands without explicit symbol UID were ambiguous; reran impact with the exact symbol UIDs returned by Anvien.

## E2 - P1/P2 Source Skill Reduction, Skill Guide, And Planner

Date: 2026-06-02

Status: recorded

Scope:

- Reduced embedded source skill registry to the final four retained workflow skills.
- Added generated `Skill Selection Guide` source logic.
- Added planner embedded skill source.
- Updated tests to assert reduced inventory, generated guide rows, generated file absence for removed skills, and embedded planner content.

Source / command evidence:

| Check | Result |
|---|---|
| `rg --files internal\aicontext\skills` | Source inventory now has 4 files: `anvien-api-surface.md`, `anvien-refactoring.md`, `anvien-debugging.md`, `anvien-planner.md`. |
| `rg -n "anvien-(cli|cross-repo|exploring|graph-quality|guide|impact-analysis|runtime-packaging)|## Skills" internal\aicontext\skills internal\aicontext\aicontext.go` | Removed skill names remain only in `retiredBaseSkillNames` cleanup, not embedded source files or generated guide rows; `## Skills` is no longer generated by `renderAnvienBlock`. |
| Source inspection | `renderAnvienBlock` now emits `## Skill Selection Guide` after MCP prompts; rows are generated from `baseSkills` and point to `.claude/skills/anvien/<skill>/SKILL.md`. |

Impact / blast radius:

| Target | Result |
|---|---|
| AI-context generator | Existing E1 HIGH/CRITICAL blast radius applies; implementation was limited to `internal/aicontext/aicontext.go`, embedded skill Markdown, and tests. |

Implementation evidence:

| File | Evidence |
|---|---|
| `internal/aicontext/aicontext.go` | `baseSkills` now registers exactly `anvien-api-surface`, `anvien-refactoring`, `anvien-debugging`, and `anvien-planner`; removed seven broad/router skills were added to `retiredBaseSkillNames`; generated root skill table changed to `Skill Selection Guide`. |
| `internal/aicontext/skills/anvien-planner.md` | Added planner workflow skill with `docs/plans/YYYY-MM-DD-<slug>/` three-file convention and plan/evidence/benchmark rules. |
| `internal/aicontext/skills/anvien-api-surface.md` | Tightened description and added command-table source rule. |
| `internal/aicontext/skills/anvien-refactoring.md` | Tightened description and added command-table source rule. |
| `internal/aicontext/skills/anvien-debugging.md` | Tightened description and added command-table source rule. |
| `internal/aicontext/skills/anvien-cli.md`, `anvien-cross-repo.md`, `anvien-exploring.md`, `anvien-graph-quality.md`, `anvien-guide.md`, `anvien-impact-analysis.md`, `anvien-runtime-packaging.md` | Deleted from embedded source inventory. |
| `internal/aicontext/aicontext_test.go` | Updated registry, generated guide, removed-skill absence, and planner content assertions. |
| `internal/cli/command_test.go` | Updated analyze smoke to expect planner skill path and reject generated `anvien-cli` path; setup install assertion checks removed broad/router skills are absent. |
| `internal/cli/package_command_test.go` | Updated package-source fixture from `anvien-cli.md` to `anvien-planner.md`. |
| `internal/mcp/query_semantic_test.go` | Updated synthetic AI-context skill fixture from `anvien-cli.md` to `anvien-planner.md`. |

Validation:

| Command | Result |
|---|---|
| `gofmt -w internal\aicontext\aicontext.go internal\aicontext\aicontext_test.go internal\cli\command_test.go internal\cli\package_command_test.go internal\mcp\query_semantic_test.go` | Pass. |
| `go test ./internal/aicontext ./internal/cli ./internal/mcp` | Pass: `internal/aicontext` 0.796s, `internal/cli` 14.451s, `internal/mcp` 7.482s. |

Failures / handling:

- None in focused source/test validation.

## E3 - P3-A/P3-B Regeneration, Validation, And README

Date: 2026-06-02

Status: recorded

Scope:

- Ran full build before post-build testing.
- Regenerated AI context through the normal `anvien analyze --force` path.
- Validated generated root guides and generated skill inventory.
- Read `README.md` end to end and updated the AI context/skills guidance.

Source / command evidence:

| Check | Result |
|---|---|
| `powershell -ExecutionPolicy Bypass -File anvien-launcher\build.ps1` | First run failed at Go binary copy because `anvien\bin\anvien.exe` was locked by two `anvien.exe mcp` processes. |
| `Get-CimInstance Win32_Process` | Found PID 5916 and PID 2468 running `anvien.exe mcp` from the npm global Anvien package path. |
| `Stop-Process -Id 5916,2468 -Force` | Stopped the two MCP processes that were holding the binary. |
| `powershell -ExecutionPolicy Bypass -File anvien-launcher\build.ps1` | Pass after stopping the locked processes; Web/Vite build completed and Go runtime binary was rebuilt. |
| `go test ./...` | Not usable as repo validation; it includes intentionally invalid fixture packages and failed on missing `baseline/phase-1-contract-freeze/ladybugdb-graph-contract.json`. Relevant internal packages still ran and passed where applicable. |
| `go test ./internal/aicontext ./internal/cli ./internal/mcp` | Pass after full build. |
| `anvien analyze --force` | Pass; scanned 815 files, parsed 598 code files, failed 0; graph has 96,204 nodes and 131,703 relationships. |
| `Get-ChildItem .claude\skills\anvien -Directory` | Generated inventory has exactly 4 directories: `anvien-api-surface`, `anvien-refactoring`, `anvien-debugging`, and `anvien-planner`. |
| `rg -n "## Command Selection Guide|## Skill Selection Guide|anvien-planner|..." AGENTS.md CLAUDE.md .claude\skills\anvien internal\aicontext\skills` | `AGENTS.md` and `CLAUDE.md` contain both generated guide headings and the four expected skill rows; generated skill files contain only the four retained skill names. |
| `rg -n "anvien-(cli|cross-repo|exploring|graph-quality|guide|impact-analysis|runtime-packaging)/SKILL.md|## Skills" AGENTS.md CLAUDE.md .claude\skills\anvien` | No matches. |
| `Get-Content README.md` | README was read end to end before editing. |
| `rg -n "Command Selection Guide|Skill Selection Guide|anvien-api-surface|anvien-refactoring|anvien-debugging|anvien-planner|..." README.md` | README now documents guide separation and the four retained generated workflow skills. |

Impact / blast radius:

| Target | Result |
|---|---|
| Generated AI context output | Changed by normal analyze path only; no manual generated-output patching. |
| README | User-facing documentation only; no runtime behavior. |

Implementation evidence:

| File | Evidence |
|---|---|
| `AGENTS.md` / `CLAUDE.md` | Regenerated managed Anvien block includes `Command Selection Guide` and `Skill Selection Guide`; old broad/router skill paths are absent. |
| `.claude/skills/anvien/**` | Regenerated output contains exactly four retained workflow skill directories. |
| `README.md` | Added AI context guidance explaining direct command selection, separate retained workflow skills, and the four generated skill paths. |

Validation:

| Command | Result |
|---|---|
| Full build | Pass after resolving locked MCP processes. |
| Focused tests | Pass: `go test ./internal/aicontext ./internal/cli ./internal/mcp`. |
| Regeneration | Pass: `anvien analyze --force`. |
| Generated inventory checks | Pass: source 4 files, generated 4 dirs, 4 `Skill Selection Guide` rows, no old generated skill paths. |

Failures / handling:

- Full build initially failed because the existing Anvien MCP processes held the runtime binary; stopped only those two `anvien.exe mcp` processes and reran the full build successfully.
- `go test ./...` is not a valid repo-wide command in the current tree because it includes intentionally invalid fixture packages and a missing frozen contract baseline path; focused affected package tests passed after the full build.

## E4 - P3-C Detect Changes Before Commit

Date: 2026-06-02

Status: recorded

Scope:

- Ran pre-commit changed-scope analysis after implementation, validation, regeneration, and README update.

Detect changes:

| Command | Result |
|---|---|
| `anvien detect-changes --repo Anvien --scope all` | Pass; summary reported 19 changed files, 13 affected files, risk level medium, and one affected process: `NewAnalyzeCommand -> RenderAnvienBlock`. |

Impact / blast radius:

| Target | Result |
|---|---|
| `internal/aicontext/aicontext.go` | Changed file risk high; linked to 3 flows and 2 tests; affected process is generated AI-context rendering. |
| Tests and docs | Changed test files are high risk by file projection due test fan-out/unresolved counts; focused affected package tests passed. |
| Resolution health | `degradedNodes=0`; `totalResolutionGapCount=0` in resolution health impact. |

Failures / handling:

- None from detect-changes.

Commit:

- `6ed6fb4 feat: reduce generated Anvien workflow skills`
