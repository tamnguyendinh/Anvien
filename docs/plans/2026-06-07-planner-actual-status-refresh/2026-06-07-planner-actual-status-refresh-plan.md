# Planner Actual Status Refresh Plan

## Metadata

- Date: `2026-06-07`
- Status: `in-progress`
- Plan: `docs/plans/2026-06-07-planner-actual-status-refresh/2026-06-07-planner-actual-status-refresh-plan.md`
- Evidence: `docs/plans/2026-06-07-planner-actual-status-refresh/2026-06-07-planner-actual-status-refresh-evidence.md`
- Benchmark: `docs/plans/2026-06-07-planner-actual-status-refresh/2026-06-07-planner-actual-status-refresh-benchmark.md`
- Actual status: `docs/plans/2026-06-07-planner-actual-status-refresh/2026-06-07-planner-actual-status-refresh-actual-status.md`

## Goal

Update the planner skill so generated actual-status files are treated as living current-state records after implementation begins, not frozen P0 snapshots.

## Rules

- Complete P0 actual status before implementation work.
- Update each checklist item immediately when it is completed.
- Record evidence as work completes.
- Record benchmarkable counts or measurements when they are taken.
- Update later phase status assumptions, next actions, and work steps when actual-status evidence changes the repo state.

## Problem

The current planner instructions and actual-status template correctly require P0 baseline classification before implementation. They do not explicitly require agents to refresh actual-status after each implementation slice or before later phases when repo reality changes. That can leave later agents with stale context.

## Scope

- `internal/aicontext/skills/planner/SKILL.md`
- `internal/aicontext/skills/planner/templates/plan.template.md`
- `internal/aicontext/skills/planner/templates/actual-status.template.md`

## Non-Goals

- Do not modify unrelated skills.
- Do not change generated `AGENTS.md` content directly.
- Do not redesign the full plan template set beyond the actual-status freshness behavior.

## Requirements

- Add rules that define actual-status as a living status record.
- Require status matrix refresh after completed implementation slices and before later phases when repo state changed.
- Avoid wording that implies changing the plan goal; update only stale phase status, next-action, and work-step fields when needed.
- Update plan-template wording so new plans do not describe next-phase status updates as broad rewrites.
- Keep evidence details in `evidence.md`; actual-status should store classification, evidence IDs, and plan consequences.
- Add a markdown `Status Refresh Log` table to the actual-status template.
- Preserve the existing P0 gate and relationship/impact evidence discipline.

## Acceptance Criteria

- Planner skill instructions mention actual-status refresh requirements after implementation begins.
- `actual-status.template.md` contains freshness/refresh rules and a status refresh log.
- `plan.template.md` uses status/next-action/work-step update wording instead of broad "rewrite phase" wording.
- Existing status matrix and implementation gate still protect P0 baseline work.
- Validation and Anvien detect-changes evidence are recorded before commit.

## Checklist

- [x] P0-A: Complete actual status before implementation work.
  - Goal: establish the real current state.
  - Work Steps: inspect source-of-truth files, classify each surface, record blocked or missing pieces, and update later phase status assumptions, next actions, and work steps from evidence.
  - Implementation Gate: no implementation or editing starts until `docs/plans/2026-06-07-planner-actual-status-refresh/2026-06-07-planner-actual-status-refresh-actual-status.md` has a final P0 decision.
  - Acceptance: actual status identifies correct, partial, missing/unbound, fake-or-stub, and blocked surfaces for this scope.
- [ ] P1-A: Add living-status refresh behavior to the planner skill.
  - Goal: make future planner users update actual-status as repo reality changes during implementation.
  - Work Steps:
    1. Update `internal/aicontext/skills/planner/SKILL.md` with explicit refresh rules.
    2. Update `internal/aicontext/skills/planner/templates/actual-status.template.md` with a freshness section and status refresh log.
    3. Update `internal/aicontext/skills/planner/templates/plan.template.md` to avoid broad "rewrite phase" wording in generated plans.
    4. Keep evidence references lightweight and avoid duplicating `evidence.md`.
  - Implementation Gate: P0 actual-status decision is complete and target files are low-risk documentation files.
  - Acceptance: the changed skill and templates clearly state when and how to refresh actual-status and how to update stale next-phase state.
- [ ] Pn-A: Call supervisor for the implemented-plan acceptance loop.
  - Goal: verify the completed plan work against the accepted plan, actual-status decisions, evidence, benchmark, changed files, generated output, and validation results before closure.
  - Work Steps:
    1. Call the supervisor skill to review the full completed plan work.
    2. If supervisor fails the work, return to the responsible implementation workflow/skill for the failed scope only.
    3. Re-run supervisor review after the fix.
    4. Repeat until supervisor passes or records a blocker.
  - Implementation Gate: all planned implementation phases must be completed or explicitly blocked before this review.
  - Acceptance: supervisor review passes, or the plan records a blocker with evidence and no closure is performed.
- [ ] Pn-B: Remove dead work created during this plan.
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
    1. Run the required final validation for the accepted scope.
    2. Regenerate generated outputs if source-of-truth changes require it.
    3. Run Anvien detect-changes before commit when implementation work was performed.
    4. Record final validation, detect-changes, benchmark, and commit evidence.
    5. Commit the completed scope and verify the worktree state.
  - Implementation Gate: Pn-A and Pn-B must pass or record blockers.
  - Acceptance: final evidence is recorded, required commits exist, and the worktree state is known.

## Risk Notes

- Blast radius is LOW for all target markdown files.
- The worktree already has unrelated changes in other skill folders; this plan must not include or revert them.
