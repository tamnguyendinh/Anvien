# Tree-sitter Go Module Drift Plan

## Metadata

- Date: `2026-07-16`
- Status: `implementation complete - pending supervisor/commit closure`
- Plan: `docs/plans/2026-07-16-tree-sitter-go-module-drift/2026-07-16-tree-sitter-go-module-drift-plan.md`
- Evidence: `docs/plans/2026-07-16-tree-sitter-go-module-drift/2026-07-16-tree-sitter-go-module-drift-evidence.md`
- Benchmark: `docs/plans/2026-07-16-tree-sitter-go-module-drift/2026-07-16-tree-sitter-go-module-drift-benchmark.md`
- Actual status: `docs/plans/2026-07-16-tree-sitter-go-module-drift/2026-07-16-tree-sitter-go-module-drift-actual-status.md`

## Goal

Repair the Tree-sitter dependency drift monitoring so it reflects Anvien's current Go module parser stack instead of the retired npm `tree-sitter` runtime model.

## Rules

- Complete P0 actual status before implementation work.
- Update each checklist item immediately when it is completed.
- Record evidence as work completes.
- Record benchmarkable counts or measurements when they are taken.
- Update later phase status assumptions, next actions, and work steps when actual-status evidence changes the repo state.
- After completing a phase or implementation slice and refreshing `actual-status.md`, update the next affected phase's work steps as needed to match the latest repo reality, while preserving that phase's original goal, scope, acceptance criteria, and major phase order.
- Run Anvien detect-changes before every implementation-slice commit when implementation work was performed.
- Every implementation phase must be decomposed into narrow slices. Do not implement a phase directly.
- Hidden fallback is forbidden. Prefer a visible failure over a fallback that hides a broken primary path.
- Tests must prove behavior: the checker must run successfully, inventory the real Go module parser stack, and distinguish real update drift from upstream core releases that are not yet published as Go module tags.
- This plan has no UI or DB behavior. UI flow, DB read/write, render location, Docker runtime, Playwright target, and persistent cleanup checks are `N/A` unless an implementation slice introduces such a surface.

## Problem

The current Tree-sitter readiness script still assumes the old npm parser dependency model. It reads `anvien/package.json["dependencies"]["tree-sitter"]`, but `anvien/package.json` no longer has a `dependencies` section. Running the script now crashes with `KeyError: 'dependencies'`.

The repository's current parser source of truth is `go.mod`, where Tree-sitter bindings and grammars are Go modules such as `github.com/tree-sitter/go-tree-sitter`, `github.com/tree-sitter/tree-sitter-go`, and `github.com/tree-sitter/tree-sitter-javascript`.

This is a maintenance/CI dependency monitoring bug, not a runtime parser failure. Package build and `anvien analyze --force` currently pass.

## Scope

- Rewrite `.github/scripts/check-tree-sitter-upgrade-readiness.py` to inventory Tree-sitter modules from Go module data, not npm package metadata.
- Update `.github/workflows/tree-sitter-upgrade-readiness.yml` names, issue text, and schedule behavior so the report describes Go module drift accurately.
- Update `.github/dependabot.yml` so Go module parser dependencies are monitored through a `gomod` entry, while stale npm Tree-sitter comments are removed or corrected.
- Add focused tests or fixture-driven self-checks for the script so the missing npm `dependencies` field cannot regress.
- Validate the corrected checker locally and through the relevant build/test path.

## Non-Goals

- Do not force Tree-sitter core to `v0.26.11` unless a matching Go module tag exists and tests prove compatibility.
- Do not change parser extraction semantics, provider behavior, graph schema, LadybugDB runtime handling, or code analysis output.
- Do not update LadybugDB; its auto resolver is working and is outside this bug.
- Do not edit generated `AGENTS.md` or `CLAUDE.md`.
- Do not add hidden fallback behavior that silently reports success when dependency inventory fails.

## Requirements

- The checker must read the current Tree-sitter parser stack from `go.mod` or `go list -m -json all`.
- The checker must use `go list -m -u` or equivalent Go module metadata to detect available Go module updates.
- The checker must separately report upstream Tree-sitter core latest release as informational context.
- The checker must classify at least:
  - `UP_TO_DATE`
  - `GO_MODULE_UPDATE_AVAILABLE`
  - `UPSTREAM_CORE_AHEAD_GO_BINDING`
  - `GRAMMAR_UPDATE_AVAILABLE`
  - `UNKNOWN_FETCH_FAILED`
- The checker must not crash when `anvien/package.json` has no `dependencies`.
- The workflow and tracking issue title/body must not say the repo is waiting for npm `tree-sitter@0.25`.
- Dependabot config must cover root Go modules and avoid stale comments that describe the retired npm grammar stack.

## Acceptance Criteria

- `python .github/scripts/check-tree-sitter-upgrade-readiness.py` exits cleanly and emits a Markdown report for the Go module parser stack.
- The report includes `github.com/tree-sitter/go-tree-sitter` and the Tree-sitter grammar modules from `go.mod`.
- The report distinguishes a real Go module update from an upstream Tree-sitter core release that has no corresponding Go module tag.
- `.github/workflows/tree-sitter-upgrade-readiness.yml` creates or updates a tracking issue whose title/body reflects Go module drift.
- `.github/dependabot.yml` contains a root `gomod` update entry or an explicit recorded no-change decision with evidence.
- Focused tests or script self-checks cover missing `anvien/package.json.dependencies`, Go module inventory parsing, update classification, and report formatting.
- Full build passes before final validation.
- `anvien detect-changes --repo Anvien --scope all` runs before the implementation commit.

## Checklist

- [x] P0-A: Complete actual status before implementation work.
  - Goal: establish the real current state of Tree-sitter dependency monitoring.
  - Work Steps: refresh graph, inspect current checker/workflow/dependabot/go.mod, reproduce the script crash, verify runtime build/analyze still works, and record the status matrix.
  - Implementation Gate: no implementation or editing starts until `docs/plans/2026-07-16-tree-sitter-go-module-drift/2026-07-16-tree-sitter-go-module-drift-actual-status.md` has a final P0 decision.
  - Acceptance: actual status identifies correct, partial, wrong, missing, and preserve-only surfaces for this scope.

### P1: Go Module Drift Checker

- Phase Goal: replace the stale npm readiness logic with a Go-module-aware Tree-sitter drift checker.
- Phase Boundary:
  - In scope: `.github/scripts/check-tree-sitter-upgrade-readiness.py` and focused test/fixture files for that script.
  - Out of scope: workflow issue text, Dependabot config, parser runtime upgrades.
  - Dependencies: P0 actual status and fresh source inspection.
- Phase Implementation Rule: do not implement `P1` directly. Implement `P1-A`, verify it, record evidence, refresh actual-status, commit when required, then continue.
- Ordered Slice List:
  - P1-A: Replace npm runtime inventory with Go module inventory.
  - P1-B: Add update classification and failure semantics.
  - P1-C: Add script tests or self-check fixtures.

- [x] P1-A: Replace npm runtime inventory with Go module inventory.
  - Goal: make the checker discover the current Tree-sitter parser stack from `go.mod` or `go list -m -json all`.
  - Scope Boundary:
    - Editable: `.github/scripts/check-tree-sitter-upgrade-readiness.py`.
    - Inspect-only: `go.mod`, parser registry files, existing workflow.
    - Preserve-only: LadybugDB runtime scripts and parser provider implementation.
    - Out of scope: changing Tree-sitter module versions.
  - Non-Goals: do not add npm parser logic as a fallback.
  - Pre-flight Questions:
    - Data source: `go.mod` and `go list -m -json all`.
    - Display permission: N/A.
    - DB read flow: N/A.
    - DB write flow: N/A.
    - Render location: Markdown report stdout.
    - UI behavior flow: N/A.
    - Docker runtime: N/A.
    - Playwright target: N/A.
    - Behavior test: script run and focused fixture/self-check.
    - Cleanup/quarantine: no persistent files outside repo.
    - External side effects: Go module metadata lookup and optional upstream release metadata fetch.
    - N/A notes: CI maintenance script only.
  - Work Steps:
    1. Inspect current parser module list from `go.mod` and decide the exact module-name filters.
       - UI flow check: N/A.
       - DB/data flow check: N/A.
       - Render location check: stdout report.
       - Evidence target: `E1-P1A-SRC1`.
    2. Replace `read_current_runtime()` and npm grammar inventory with Go module inventory.
       - UI flow check: N/A.
       - DB/data flow check: N/A.
       - Render location check: report still has stable Markdown tables.
       - Evidence target: `E1-P1A-SRC2`.
  - Implementation Gate:
    - Before editing, run `anvien file-detail go.mod --repo Anvien --json` and record evidence.
    - If editing Python logic only, record that `.github` files are not available through current file-detail indexing and use direct source inspection as evidence.
  - Acceptance:
    - Source: checker no longer reads `anvien/package.json["dependencies"]["tree-sitter"]`.
    - Runtime/UI: N/A.
    - DB/data: N/A.
    - Behavior test: script can inventory Go Tree-sitter modules without crashing.
    - Cleanup/quarantine: no temp files outside repo.
    - Evidence IDs: `E1-P1A-SRC1`, `E1-P1A-SRC2`, `E1-P1A-RUN1`.
    - Actual-status rows refreshed: checker script.
  - Evidence Targets: source diff, script run, status update.
  - Actual-status Update: checker script `wrong -> partial`.
  - Commit Boundary: commit after P1-C, not after this slice, because P1-A without tests is incomplete.

- [x] P1-B: Add update classification and failure semantics.
  - Goal: classify Go module update drift and upstream core lag without false runtime failure claims.
  - Scope Boundary:
    - Editable: `.github/scripts/check-tree-sitter-upgrade-readiness.py`.
    - Inspect-only: `go.mod`, `go list -m -u` output, official Tree-sitter release metadata fetch.
    - Preserve-only: existing GitHub token handling and no-external-deps standard-library implementation.
    - Out of scope: Dependabot config.
  - Non-Goals: do not fail the checker just because Tree-sitter core has a newer release than the latest Go binding tag.
  - Pre-flight Questions:
    - Data source: `go list -m -u`, `go list -m -versions` where needed, upstream release metadata.
    - Display permission: N/A.
    - DB read flow: N/A.
    - DB write flow: N/A.
    - Render location: Markdown report stdout.
    - UI behavior flow: N/A.
    - Docker runtime: N/A.
    - Playwright target: N/A.
    - Behavior test: fixture/self-check for each classification.
    - Cleanup/quarantine: no persistent files outside repo.
    - External side effects: network metadata lookup may fail and must be visible.
    - N/A notes: CI maintenance script only.
  - Work Steps:
    1. Implement status mapping for Go module updates, upstream core ahead, and fetch failures.
       - UI flow check: N/A.
       - DB/data flow check: N/A.
       - Render location check: report summary and table.
       - Evidence target: `E1-P1B-SRC1`.
    2. Define exit behavior: fail for checker/internal inventory failure or actionable Go module drift; report upstream core lag as informational unless policy says otherwise.
       - UI flow check: N/A.
       - DB/data flow check: N/A.
       - Render location check: report summary.
       - Evidence target: `E1-P1B-RUN1`.
  - Implementation Gate:
    - Classification names must match the requirements section.
  - Acceptance:
    - Source: report uses explicit statuses and no npm peer-dep language.
    - Runtime/UI: N/A.
    - DB/data: N/A.
    - Behavior test: controlled cases produce expected statuses.
    - Cleanup/quarantine: no temp files outside repo.
    - Evidence IDs: `E1-P1B-SRC1`, `E1-P1B-RUN1`.
    - Actual-status rows refreshed: checker classification.
  - Evidence Targets: source diff, focused command output.
  - Actual-status Update: checker script `partial -> partial` until tests land.
  - Commit Boundary: commit after P1-C.

- [x] P1-C: Add script tests or self-check fixtures.
  - Goal: prevent regression to the retired npm dependency model.
  - Scope Boundary:
    - Editable: new focused script test file or a fixture/self-check mode under `.github/scripts` or an approved test location.
    - Inspect-only: existing CI test conventions.
    - Preserve-only: production checker behavior.
    - Out of scope: broad CI refactor.
  - Non-Goals: do not require network access for core parsing/classification unit coverage.
  - Pre-flight Questions:
    - Data source: local fixture text/JSON for Go module inventory and update states.
    - Display permission: N/A.
    - DB read flow: N/A.
    - DB write flow: N/A.
    - Render location: test output.
    - UI behavior flow: N/A.
    - Docker runtime: N/A.
    - Playwright target: N/A.
    - Behavior test: Python unittest/pytest-compatible stdlib test or script self-check.
    - Cleanup/quarantine: no persistent state.
    - External side effects: none for fixture tests.
    - N/A notes: maintenance script tests only.
  - Work Steps:
    1. Add fixture coverage for a package with no npm `dependencies` section.
       - UI flow check: N/A.
       - DB/data flow check: N/A.
       - Render location check: test output.
       - Evidence target: `E1-P1C-TEST1`.
    2. Add fixture coverage for `UP_TO_DATE`, `GO_MODULE_UPDATE_AVAILABLE`, and `UPSTREAM_CORE_AHEAD_GO_BINDING`.
       - UI flow check: N/A.
       - DB/data flow check: N/A.
       - Render location check: test output.
       - Evidence target: `E1-P1C-TEST2`.
  - Implementation Gate:
    - Tests must be runnable locally without GitHub credentials.
  - Acceptance:
    - Source: focused tests or self-check fixture exist.
    - Runtime/UI: N/A.
    - DB/data: N/A.
    - Behavior test: tests pass and the live script runs without crashing.
    - Cleanup/quarantine: no temp files outside repo.
    - Evidence IDs: `E1-P1C-TEST1`, `E1-P1C-TEST2`, `E1-P1C-RUN1`.
    - Actual-status rows refreshed: checker tests.
  - Evidence Targets: test command output, live script output.
  - Actual-status Update: checker script `partial -> correct`.
  - Commit Boundary: commit after this slice when P1 acceptance passes.

### P2: CI and Dependency Configuration

- Phase Goal: align workflow and Dependabot configuration with the Go module checker.
- Phase Boundary:
  - In scope: `.github/workflows/tree-sitter-upgrade-readiness.yml`, `.github/dependabot.yml`.
  - Out of scope: script logic already handled in P1.
  - Dependencies: P1 complete or blocked with a stable report contract.
- Phase Implementation Rule: do not implement `P2` directly. Implement `P2-A`, verify it, record evidence, refresh actual-status, then continue.
- Ordered Slice List:
  - P2-A: Update workflow title, issue text, and parser for new report.
  - P2-B: Add or correct Dependabot Go module monitoring.

- [x] P2-A: Update workflow title, issue text, and parser for new report.
  - Goal: make scheduled CI reports and GitHub issues describe Tree-sitter Go module drift accurately.
  - Scope Boundary:
    - Editable: `.github/workflows/tree-sitter-upgrade-readiness.yml`.
    - Inspect-only: checker report output from P1.
    - Preserve-only: schedule, permissions, and pinned action SHAs unless the change is necessary.
    - Out of scope: changing the checker logic.
  - Non-Goals: do not rename unrelated workflows.
  - Pre-flight Questions:
    - Data source: P1 report output.
    - Display permission: GitHub Actions issue body.
    - DB read flow: N/A.
    - DB write flow: N/A.
    - Render location: workflow logs and tracking issue.
    - UI behavior flow: N/A.
    - Docker runtime: N/A.
    - Playwright target: N/A.
    - Behavior test: workflow YAML/source inspection and local checker output shape.
    - Cleanup/quarantine: no persistent state.
    - External side effects: scheduled GitHub issue update when workflow runs in CI.
    - N/A notes: CI text/config only.
  - Work Steps:
    1. Replace `Tree-sitter 0.25 upgrade readiness` naming with Go module drift naming.
       - UI flow check: N/A.
       - DB/data flow check: N/A.
       - Render location check: workflow name and issue title.
       - Evidence target: `E2-P2A-SRC1`.
    2. Update issue body/comment parsing to match the new report's module status rows.
       - UI flow check: N/A.
       - DB/data flow check: N/A.
       - Render location check: GitHub issue summary.
       - Evidence target: `E2-P2A-SRC2`.
  - Implementation Gate:
    - P1 report format must be known before editing issue parsing.
  - Acceptance:
    - Source: workflow no longer mentions npm `tree-sitter@0.25` as the current blocker model.
    - Runtime/UI: N/A.
    - DB/data: N/A.
    - Behavior test: report text can be embedded in workflow outputs and issue body logic remains coherent.
    - Cleanup/quarantine: no temp files outside repo.
    - Evidence IDs: `E2-P2A-SRC1`, `E2-P2A-SRC2`.
    - Actual-status rows refreshed: workflow.
  - Evidence Targets: workflow diff and optional YAML parse.
  - Actual-status Update: workflow `wrong -> correct`.
  - Commit Boundary: commit after P2-B when configuration acceptance passes.

- [x] P2-B: Add or correct Dependabot Go module monitoring.
  - Goal: make automated dependency PRs cover root Go modules that include Tree-sitter parser dependencies.
  - Scope Boundary:
    - Editable: `.github/dependabot.yml`.
    - Inspect-only: current `go.mod`, npm entries for `/anvien` and `/anvien-web`.
    - Preserve-only: GitHub Actions and Web npm update entries unless stale comments require correction.
    - Out of scope: applying dependency upgrades.
  - Non-Goals: do not remove non-Tree-sitter dependency monitoring unless evidence proves it is obsolete.
  - Pre-flight Questions:
    - Data source: `go.mod`, current Dependabot config.
    - Display permission: GitHub PR labels/messages.
    - DB read flow: N/A.
    - DB write flow: N/A.
    - Render location: Dependabot config.
    - UI behavior flow: N/A.
    - Docker runtime: N/A.
    - Playwright target: N/A.
    - Behavior test: YAML/source inspection; optional Dependabot schema cannot be fully run locally.
    - Cleanup/quarantine: no persistent state.
    - External side effects: future Dependabot PRs.
    - N/A notes: configuration only.
  - Work Steps:
    1. Add a root `gomod` Dependabot entry or record a no-change decision if an equivalent exists.
       - UI flow check: N/A.
       - DB/data flow check: N/A.
       - Render location check: `.github/dependabot.yml`.
       - Evidence target: `E2-P2B-SRC1`.
    2. Remove stale npm Tree-sitter runtime comments that claim the runtime is pinned in `/anvien/package.json`.
       - UI flow check: N/A.
       - DB/data flow check: N/A.
       - Render location check: `.github/dependabot.yml`.
       - Evidence target: `E2-P2B-SRC2`.
  - Implementation Gate:
    - Do not remove `/anvien` or `/anvien-web` npm update entries unless unrelated evidence proves they are obsolete.
  - Acceptance:
    - Source: root Go module dependency updates are monitored.
    - Runtime/UI: N/A.
    - DB/data: N/A.
    - Behavior test: config remains valid YAML and comments match current dependency ownership.
    - Cleanup/quarantine: no temp files outside repo.
    - Evidence IDs: `E2-P2B-SRC1`, `E2-P2B-SRC2`, `E2-P2B-VALIDATE1`.
    - Actual-status rows refreshed: Dependabot config.
  - Evidence Targets: config diff and validation command.
  - Actual-status Update: Dependabot `partial -> correct`.
  - Commit Boundary: commit after this slice when P2 acceptance passes.

### P3: Final Validation and Closure

- Phase Goal: prove the corrected maintenance path works and close the plan cleanly.
- Phase Boundary:
  - In scope: final script command, focused tests, full build, Anvien detect-changes, supervisor review.
  - Out of scope: unrelated dependency upgrades and parser runtime changes.
  - Dependencies: P1 and P2 complete.
- Phase Implementation Rule: do not implement `P3` directly. Execute `P3-A`, then close through `Pn-*`.
- Ordered Slice List:
  - P3-A: Validate script, tests, build, and graph behavior.

- [x] P3-A: Validate script, tests, build, and graph behavior.
  - Goal: provide final evidence that the checker is repaired and Anvien runtime remains healthy.
  - Scope Boundary:
    - Editable: validation-only unless a failure reveals a bug in P1/P2 work.
    - Inspect-only: build scripts, package runtime, Anvien graph output.
    - Preserve-only: LadybugDB native resolver and parser runtime behavior.
    - Out of scope: dependency bump PRs.
  - Non-Goals: do not substitute a successful build for the required checker behavior test.
  - Pre-flight Questions:
    - Data source: live repo files and Go module metadata.
    - Display permission: CLI/stdout.
    - DB read flow: N/A.
    - DB write flow: N/A.
    - Render location: command output and evidence ledger.
    - UI behavior flow: N/A.
    - Docker runtime: N/A unless full build script requires it.
    - Playwright target: N/A.
    - Behavior test: checker command, focused script tests, full build, `anvien analyze --force`.
    - Cleanup/quarantine: remove or record any plan-created temp output.
    - External side effects: metadata fetches only.
    - N/A notes: maintenance validation only.
  - Work Steps:
    1. Run the focused checker tests and live checker command.
       - UI flow check: N/A.
       - DB/data flow check: N/A.
       - Render location check: report output.
       - Evidence target: `E3-P3A-RUN1`.
    2. Run full build and `anvien analyze --force` to prove runtime remains healthy.
       - UI flow check: N/A.
       - DB/data flow check: graph build succeeds.
       - Render location check: CLI output.
       - Evidence target: `E3-P3A-BUILD1`.
    3. Run `anvien detect-changes --repo Anvien --scope all` before commit.
       - UI flow check: N/A.
       - DB/data flow check: change inventory.
       - Render location check: CLI output.
       - Evidence target: `E3-P3A-DETECT1`.
  - Implementation Gate:
    - P1 and P2 acceptance must be complete or explicitly blocked.
  - Acceptance:
    - Source: no unresolved stale npm Tree-sitter assumptions remain in planned files.
    - Runtime/UI: N/A.
    - DB/data: Anvien analyze succeeds after the change.
    - Behavior test: checker command and focused tests pass.
    - Cleanup/quarantine: no dead artifacts remain.
    - Evidence IDs: `E3-P3A-RUN1`, `E3-P3A-BUILD1`, `E3-P3A-DETECT1`.
    - Actual-status rows refreshed: all target rows.
  - Evidence Targets: script/test/build/analyze/detect-changes output.
  - Actual-status Update: all planned rows transition to `correct` or blocked with evidence.
  - Commit Boundary: commit after this slice when acceptance passes.

- [x] Pn-A: Call supervisor for the implemented-plan acceptance loop.
  - Goal: verify the completed plan work against this plan, actual-status decisions, evidence, benchmark, changed files, and validation results before closure.
  - Work Steps:
    1. Call the supervisor skill to review the full completed plan work.
    2. If supervisor fails the work, return to the responsible implementation workflow/skill for the failed scope only.
    3. Re-run supervisor review after the fix.
    4. Repeat until supervisor passes or records a blocker.
  - Implementation Gate: all planned implementation phases must be completed or explicitly blocked before this review.
  - Acceptance: supervisor review passes, or the plan records a blocker with evidence and no closure is performed.
- [x] Pn-B: Remove dead work created during this plan.
  - Goal: ensure the final diff contains only artifacts that still serve the accepted plan.
  - Work Steps:
    1. Review files, sections, generated output, tests, temp files, and plan artifacts created or modified during this plan.
    2. Remove or rewrite any artifact made obsolete by actual-status findings, user corrections, failed approaches, or phase status updates.
    3. Verify no rejected approach, stale placeholder, unused generated output, or dead helper artifact remains in the final diff.
    4. Call supervisor to review the dead-work cleanup.
    5. If supervisor fails the cleanup, return to the responsible implementation workflow/skill for the failed cleanup scope only, then re-run supervisor review.
  - Implementation Gate: only remove artifacts created by this plan unless the user explicitly approves broader cleanup.
  - Acceptance: final `git diff/status` contains no dead plan-created artifacts, supervisor passes the cleanup, and evidence records what was removed or preserved.
- [ ] Pn-C: Close the plan.
  - Goal: finish validation, evidence, benchmark, detect-changes, commit, and final status.
  - Work Steps:
    1. Confirm final validation evidence exists for checker tests, live checker run, full build, analyze, and detect-changes.
    2. Regenerate generated outputs only if source-of-truth changes require it.
    3. Record final validation, detect-changes, benchmark, and commit evidence.
    4. Commit the completed scope and verify the worktree state.
  - Implementation Gate: Pn-A and Pn-B must pass or record blockers.
  - Acceptance: final evidence is recorded, required commits exist, and the worktree state is known.

## Risk Notes

- The checker touches CI maintenance, not runtime parsing. Treat runtime changes as out of scope unless validation exposes a real parser bug.
- `go-tree-sitter v0.25.0` can be latest as a Go module while upstream Tree-sitter core is `v0.26.11`; the report must not call that a runtime failure.
- Network metadata fetches can be flaky. The report must distinguish fetch failure from dependency drift.
- `.github` files are not available through current `anvien file-detail` lookup, so direct source inspection is required for those files and must be recorded as such.
