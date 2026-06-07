# {{TITLE}} Plan

## Metadata

- Date: `{{YYYY-MM-DD}}`
- Status: `draft`
- Plan: `{{PLAN_PATH}}`
- Evidence: `{{EVIDENCE_PATH}}`
- Benchmark: `{{BENCHMARK_PATH}}`
- Actual status: `{{ACTUAL_STATUS_PATH}}`

## Goal

{{GOAL}}

## Rules

- Complete P0 actual status before implementation work.
- Update each checklist item immediately when it is completed.
- Record evidence as work completes.
- Record benchmarkable counts or measurements when they are taken.
- Update later phase status assumptions, next actions, and work steps when actual-status evidence changes the repo state.
 - After completing a phase or implementation slice and refreshing `actual- status.md`, update the next affected phase's work steps as needed to match the latest repo reality, while preserving that phase's original goal, scope,  acceptance criteria, and major phase order.

## Problem

{{PROBLEM}}

## Scope

{{SCOPE}}

## Non-Goals

{{NON_GOALS}}

## Requirements

{{REQUIREMENTS}}

## Acceptance Criteria

{{ACCEPTANCE_CRITERIA}}

## Checklist

- [ ] P0-A: Complete actual status before implementation work.
  - Goal: establish the real current state.
  - Work Steps: inspect source-of-truth files, classify each surface, record blocked or missing pieces, and update later phase status assumptions, next actions, and work steps from evidence.
  - Implementation Gate: no implementation or editing starts until `{{ACTUAL_STATUS_PATH}}` has a final P0 decision.
  - Acceptance: actual status identifies correct, partial, missing/unbound, fake-or-stub, and blocked surfaces for this scope.
 - [ ] P1-A: {{PHASE_1_TITLE}}
  - Goal: {{PHASE_1_GOAL}}
  - Work Steps: {{PHASE_1_WORK_STEPS}}
  - Implementation Gate: {{PHASE_1_GATE}}
  - Acceptance: {{PHASE_1_ACCEPTANCE}}
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

{{RISK_NOTES}}
