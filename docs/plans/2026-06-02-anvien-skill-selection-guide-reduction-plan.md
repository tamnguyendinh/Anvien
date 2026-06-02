# Anvien Skill Selection Guide Reduction Plan

Date: 2026-06-02

Status: Active

Companion files:

- Evidence ledger: [2026-06-02-anvien-skill-selection-guide-reduction-evidence.md](2026-06-02-anvien-skill-selection-guide-reduction-evidence.md)
- Benchmark ledger: [2026-06-02-anvien-skill-selection-guide-reduction-benchmark.md](2026-06-02-anvien-skill-selection-guide-reduction-benchmark.md)

## Master Rules

1. Follow active workspace and repository instructions, including generated `AGENTS.md`; this plan records product work and validation, it does not replace repository rules.
2. Do not use Anvien for doc-only commits.
3. Use Anvien for implementation slices that inspect code ownership, graph impact, API surfaces, refactoring blast radius, or debugging paths.
4. Refresh the graph with `anvien analyze --force` before graph-based implementation evidence.
5. Run impact analysis before editing generator functions, exported/shared contracts, API handlers, analyzer/resolver code, or retained skill workflow owners.
6. Do not edit generated `AGENTS.md`, `CLAUDE.md`, or `.claude/skills/anvien/**` as the permanent source of truth.
7. Update `internal/aicontext/aicontext.go` and embedded source skills under `internal/aicontext/skills/*.md`; regenerate generated outputs through the normal analyze path.
8. Keep the generated `Command Selection Guide` and generated `Skill Selection Guide` separate.
9. Run the full build before testing. For this repo the full build gate is `powershell -ExecutionPolicy Bypass -File anvien-launcher\build.ps1`.
10. Record evidence as each evidenced task completes.
11. Record benchmarkable inventory counts as each benchmarkable task completes.
12. Run `anvien detect-changes --repo Anvien --scope all` before implementation commits.
13. Commit each completed implementation slice after evidence and benchmark ledgers are updated.

## Goal

Reduce the generated Anvien skill set to four mandatory domain workflow skills and add a separate generated `Skill Selection Guide` that routes agents to those skills by task. Concrete Anvien CLI/MCP command selection must remain in the generated `Command Selection Guide`.

Retained skills:

- `anvien-api-surface`
- `anvien-refactoring`
- `anvien-debugging`
- `anvien-planner`

## Problem

The generated Anvien skill set currently includes broad command-router skills such as `anvien-cli`, `anvien-exploring`, `anvien-graph-quality`, `anvien-guide`, `anvien-impact-analysis`, `anvien-cross-repo`, and `anvien-runtime-packaging`. These skills overlap with the generated command table and can cause agents to route ordinary command selection through a skill layer before calling the actual Anvien command.

That is the wrong model for Anvien command use. The command table should answer:

```text
When you need to... -> Use this Anvien CLI/MCP command
```

Skills should answer a different question:

```text
When you need this workflow... -> Use this Anvien workflow skill
```

Keeping many broad skills makes the skill layer look like a universal command router. Removing all skills would also be wrong because four workflows are useful as mandatory domain gates:

- API surface work needs a focused API route/tool/contract workflow.
- Refactoring work needs a focused rename/extract/move/split workflow.
- Debugging work needs a focused bug/failure/diagnostics workflow.
- Plan authoring needs a focused `docs/plans` plan/evidence/benchmark workflow.

The product fix is to generate both tables from source-owned AI-context generation:

- `Command Selection Guide` for direct Anvien commands.
- `Skill Selection Guide` for exactly the four retained workflow skills.

## Scope

In scope:

- `internal/aicontext/aicontext.go` base skill registry and generated guidance tables.
- Embedded source skills under `internal/aicontext/skills/*.md`.
- Generated root AI context outputs produced by the normal generation path, including `AGENTS.md` and `CLAUDE.md`.
- Generated `.claude/skills/anvien/**` output produced by the normal generation path.
- Tests that assert generated skill inventory, generated table content, source-vs-generated rules, setup/package inventories, or embedded skill content.
- Evidence and benchmark ledgers for source skill inventory, generated skill inventory, build/test results, regeneration, and detect-changes.

Out of scope:

- Removing or changing Anvien commands such as `query`, `context`, `impact`, `graph-health`, `query-health`, `resolution-inventory`, `source-site-accuracy`, `benchmark-compare`, or `detect-changes`.
- Changing graph-health, query-health, impact, refactoring, debugging, or API-surface command behavior.
- Editing generated `AGENTS.md`, `CLAUDE.md`, or `.claude/skills/anvien/**` directly as the source of truth.
- Adding new workflow skills beyond the four retained skills.
- Reworking non-Anvien plugin or Codex global skill behavior.

## Requirements

1. `internal/aicontext/aicontext.go` must register exactly four generated Anvien skills: `anvien-api-surface`, `anvien-refactoring`, `anvien-debugging`, and `anvien-planner`.
2. `internal/aicontext/skills` must keep exactly four Anvien embedded source skill files:
   - `anvien-api-surface.md`
   - `anvien-refactoring.md`
   - `anvien-debugging.md`
   - `anvien-planner.md`
3. The following embedded source skill files must be deleted:
   - `anvien-cli.md`
   - `anvien-cross-repo.md`
   - `anvien-exploring.md`
   - `anvien-graph-quality.md`
   - `anvien-guide.md`
   - `anvien-impact-analysis.md`
   - `anvien-runtime-packaging.md`
4. Generated `.claude/skills/anvien/**` output must contain only:
   - `.claude/skills/anvien/anvien-api-surface/SKILL.md`
   - `.claude/skills/anvien/anvien-refactoring/SKILL.md`
   - `.claude/skills/anvien/anvien-debugging/SKILL.md`
   - `.claude/skills/anvien/anvien-planner/SKILL.md`
5. Generated AI context must include a `Skill Selection Guide` table separate from the `Command Selection Guide`.
6. The `Skill Selection Guide` must contain these rows:

| When you need to... | Use |
|---|---|
| Inspect API routes, MCP tools, contracts, response shapes, or consumers | `.claude/skills/anvien/anvien-api-surface/SKILL.md` |
| Rename, extract, split, move, or restructure code | `.claude/skills/anvien/anvien-refactoring/SKILL.md` |
| Debug bugs, failures, diagnostics, or failure traces | `.claude/skills/anvien/anvien-debugging/SKILL.md` |
| Create or review `docs/plans` plan, evidence, or benchmark work | `.claude/skills/anvien/anvien-planner/SKILL.md` |

7. The generated `Command Selection Guide` must continue to route directly to CLI/MCP commands, not to skill files.
8. Retained skill descriptions must be domain-specific workflow triggers, not generic command routers.
9. After a retained skill is selected, concrete Anvien commands still come from the generated `Command Selection Guide`.
10. Generated root Skills tables, if present, must list only the four retained Anvien skills.
11. Tests must prove the source skill inventory, generated skill inventory, and generated table content cannot drift.
12. The normal generation path must recreate the final outputs without manual patching generated files.

## Invariants

1. The command table is canonical for command selection.
2. The skill table is canonical for retained workflow skill selection.
3. Skills do not replace direct CLI/MCP command execution.
4. Removed skills do not remove command capability.
5. Generated outputs are validation artifacts, not source of truth.
6. Embedded source skill Markdown and `aicontext.go` registry must remain in sync.
7. The four retained skills must remain useful as mandatory project workflow gates.
8. AI-context guidance must remain repo-agnostic and must not hard-code Anvien as the only possible indexed repo name.

## Technical Direction

Update the AI-context generator source so skill and command routing are separate concepts.

Expected generated section shape:

```text
## Command Selection Guide

Use Anvien by task, not by a fixed workflow. Pick the command surface that matches the job.

| When you need to... | Use |
|---|---|
| Find where a concept, behavior, or bug lives | `anvien query "<concept>" --repo <repo>` |
| Inspect one symbol deeply | `anvien context symbol "<symbol>" --repo <repo>` |
| Check symbol blast radius before editing | `anvien impact symbol "<symbol>" --repo <repo> --direction upstream` |

## Skill Selection Guide

Use Anvien workflow skills only for the retained domains below.

| When you need to... | Use |
|---|---|
| Inspect API routes, MCP tools, contracts, response shapes, or consumers | `.claude/skills/anvien/anvien-api-surface/SKILL.md` |
| Rename, extract, split, move, or restructure code | `.claude/skills/anvien/anvien-refactoring/SKILL.md` |
| Debug bugs, failures, diagnostics, or failure traces | `.claude/skills/anvien/anvien-debugging/SKILL.md` |
| Create or review `docs/plans` plan, evidence, or benchmark work | `.claude/skills/anvien/anvien-planner/SKILL.md` |
```

Implementation should prefer one shared table writer or one shared source data structure for retained skill rows so `AGENTS.md`, `CLAUDE.md`, generated Skills tables, and tests do not drift. If current generator code intentionally emits different sections per output, the implementation must document that difference in evidence and still validate each generated output that should contain the `Skill Selection Guide`.

The four retained skill files should keep normal skill frontmatter with `name` and `description`. Their body should describe workflow steps for that domain and then point back to direct Anvien commands from the generated command table when commands are needed. They must not contain broad "use this skill whenever using Anvien" language.

## Definition Of Done

The plan is complete when:

1. embedded source Anvien skill inventory is reduced from 10 existing broad skills to exactly 4 retained workflow skills;
2. generated `.claude/skills/anvien/**` inventory is reduced from 10 existing broad skills to exactly 4 retained workflow skills;
3. `anvien-graph-quality` and the other six removed skills are absent from source registration and generated output;
4. generated AI-context outputs include a separate `Skill Selection Guide`;
5. generated `Skill Selection Guide` rows point to the four retained generated skill paths;
6. generated `Command Selection Guide` still routes directly to commands;
7. retained skills are mandatory workflow triggers for their domains and not generic command routers;
8. build, focused tests, regeneration, source/generated inventory checks, and detect-changes evidence are recorded;
9. benchmark ledger records before/after embedded and generated skill inventory counts;
10. implementation work is committed after evidence and benchmark ledgers are updated.

## Phase Checklist

- [x] [P0-A] Establish baseline and owner evidence.
  - Goal: identify the AI-context generation owners and record the current embedded/generated skill inventory.
  - Work Steps: run `anvien analyze --force`; use direct Anvien commands from the command table to inspect AI-context generation owners; inspect `internal/aicontext/aicontext.go`, `internal/aicontext/skills`, and generated `.claude/skills/anvien`; record current counts and retained/removed skill lists.
  - Implementation Gate: no code edits in this phase.
  - Acceptance: evidence records owner files and current inventory; benchmark records baseline embedded/generated skill counts.

- [x] [P1-A] Reduce the embedded skill set.
  - Goal: remove broad command-router skills from the current embedded source set while keeping the three existing retained workflow skills.
  - Work Steps: run impact for `aicontext.go` registry/generator owners before editing; delete the seven removed embedded skill Markdown files; update `baseSkills` or equivalent registry to the retained workflow skill set.
  - Implementation Gate: do not edit generated `.claude/skills/anvien/**`, `AGENTS.md`, or `CLAUDE.md` directly.
  - Acceptance: source inventory contains only `anvien-api-surface.md`, `anvien-refactoring.md`, and `anvien-debugging.md` after removal; registry references only those three until P2-B adds `anvien-planner`.

- [x] [P1-B] Add the generated Skill Selection Guide.
  - Goal: route workflow skill selection through a separate generated table instead of mixing skills into command selection.
  - Work Steps: update `internal/aicontext/aicontext.go` table generation to emit `Skill Selection Guide`; keep `Command Selection Guide` command-only; add exact rows for the three existing retained workflow skills: API surface, refactoring, and debugging.
  - Implementation Gate: if generator emits different root contexts for different agents, validate each expected output instead of assuming one file covers all.
  - Acceptance: generated AI context can contain both guides, with commands in `Command Selection Guide` and skill paths in `Skill Selection Guide`.

- [x] [P1-C] Tighten retained skill wording.
  - Goal: keep the three existing retained skills useful as mandatory domain workflows without turning them into a generic Anvien router.
  - Work Steps: review and edit `anvien-api-surface.md`, `anvien-refactoring.md`, and `anvien-debugging.md`; remove broad command-router language; preserve direct-command guidance inside each workflow when commands are needed.
  - Implementation Gate: do not add extra retained skills while editing wording.
  - Acceptance: the three existing retained skills trigger only for their domain workflows and tell agents to use direct Anvien commands for execution.

- [x] [P2-A] Update tests for reduced inventory and generated guidance.
  - Goal: prevent source/generator/generated output drift.
  - Work Steps: update or add AI-context tests for the reduced three-skill source inventory after removed skills are deleted, exact registry inventory, generated skill file inventory, generated root Skills table, `Command Selection Guide`, and `Skill Selection Guide`.
  - Implementation Gate: tests must assert absence of removed skills, not only presence of retained skills.
  - Acceptance: focused tests fail on the old 10-skill inventory and pass with the reduced three-skill inventory plus separate skill guide; P2-B then extends the retained set to the final four-skill inventory.

- [x] [P2-B] Add the Anvien planner workflow skill.
  - Goal: make plan/evidence/benchmark authoring a first-class retained workflow skill so agents follow the repository `docs/plans` convention instead of inventing ad hoc plan files.
  - Work Steps: create embedded source skill `internal/aicontext/skills/anvien-planner.md` with valid frontmatter and the canonical `docs/plans/YYYY-MM-DD-<slug>/` three-file convention; update `internal/aicontext/aicontext.go` so `anvien-planner` is registered with the retained skill set; add a `Skill Selection Guide` row for creating or reviewing plan/evidence/benchmark work that points to `.claude/skills/anvien/anvien-planner/SKILL.md`; update tests and benchmark inventory expectations from the P2-A reduced three-skill state to the final four-skill retained state.
  - Implementation Gate: do not create `.claude/skills/anvien/anvien-planner/SKILL.md`, `AGENTS.md`, or `CLAUDE.md` manually; the generated planner skill and guide rows must come only from the normal AI-context generation path.
  - Acceptance: source inventory includes `anvien-planner.md`; generated output includes `.claude/skills/anvien/anvien-planner/SKILL.md`; generated `Skill Selection Guide` includes the planner row; tests prove the planner skill content is embedded and not fallback-generated; evidence records generated output checks and benchmark records the retained skill count change.

- [x] [P3-A] Regenerate and validate generated outputs.
  - Goal: prove the normal generation path creates the final AI context.
  - Work Steps: run full build; run focused tests; run `anvien analyze --force`; inspect regenerated `AGENTS.md`, `CLAUDE.md`, and `.claude/skills/anvien/**`; record inventory and table evidence.
  - Implementation Gate: if generation changes Web UI behavior, add Web validation; otherwise record that this is AI-context output only.
  - Acceptance: regenerated output contains only retained skill files and includes separate `Command Selection Guide` and `Skill Selection Guide`.

- [x] [P3-B] Read and update README guidance.
  - Goal: keep user-facing README guidance aligned with the new AI-context skill model without turning README into a duplicate implementation log.
  - Work Steps: read `README.md` end to end before editing; identify the README sections that explain AI context generation, generated skills, command guidance, setup, or agent workflow; update only the user-facing guidance needed to explain the separate `Command Selection Guide` and `Skill Selection Guide`, the four retained Anvien workflow skills, and the rule that concrete CLI/MCP commands are still selected directly from the command table.
  - Implementation Gate: do not edit README before generated behavior is implemented and validated; do not copy phase evidence, benchmark tables, or internal implementation details into README.
  - Acceptance: README accurately describes the final generated skill set and guide separation; evidence records the README sections reviewed and changed; README does not mention removed broad/router skills as active generated skills.

- [ ] [P3-C] Detect changes, record benchmark deltas, and commit.
  - Goal: close the implementation slice with synchronized plan/evidence/benchmark state.
  - Work Steps: update evidence with build/test/regeneration results; update benchmark with before/after inventory counts; run `anvien detect-changes --repo Anvien --scope all`; commit the completed slice.
  - Implementation Gate: do not commit until detect-changes and ledger updates are recorded.
  - Acceptance: commit hash is recorded in evidence; plan checklist reflects completed tasks.

## Risk Notes

- `internal/aicontext/aicontext.go` is the source of truth for multiple generated outputs. Editing only `AGENTS.md` behavior would be an incomplete fix.
- Removing skill files can break tests or setup/package flows that assume the old inventory count.
- A retained skill can still become a generic router if its description is too broad; tests should verify inventory, but human review must verify wording.
- Generated output paths may differ between repository-local analyze and setup/package install flows; inventory validation must cover any source-verified path in scope.
- The plan intentionally removes the `anvien-graph-quality` skill but keeps graph-health/query-health/resolution-inventory commands available through the `Command Selection Guide`.
