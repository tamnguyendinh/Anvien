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
- After completing a phase or implementation slice and refreshing `actual-status.md`, update the next affected phase's work steps as needed to match the latest repo reality, while preserving that phase's original goal, scope, acceptance criteria, and major phase order.
- Run Anvien detect-changes before every implementation-slice commit when implementation work was performed.
- For public runtime or UI-facing changes, validate the real user-visible runtime with browser or Playwright evidence.
- For app/runtime validation, full build must include Docker image/container build. If Docker is missing or not run, full build is incomplete.
- Any Playwright validation must target the real built Docker/container runtime. Running Playwright against a host dev server, framework dev mode, mocked server, or source-run shortcut is not valid runtime evidence.
- If the Docker runtime cannot be built or started, the slice/plan is blocked; do not replace it with dev-server Playwright evidence.
- Playwright evidence must record the Docker build/run or compose command, container/service name, exposed URL, Playwright command, and screenshot/trace/result.
- Keep the standard planner structure. These detail rules only make phase checklist items concrete enough to implement safely.
- Every implementation phase must be decomposed into multiple implementation slices that are as small as practical. A phase is a grouping and ordering container; a slice is the executable implementation unit.
- Do not implement a phase directly. Work starts from a slice ID such as `P1-A`, `P1-B`, or `P2-C`.
- Prefer many narrow slices over one broad slice. A single-slice implementation phase is allowed only when the plan explicitly states why the phase cannot be split further without creating empty or non-executable slices.
- Each implementation slice must include:
  + Goal
  + Scope Boundary
  + Non-Goals when useful
  + Pre-flight Questions
  + Work Steps
  + Implementation Gate
  + Acceptance
  + Evidence Targets
  + Actual-status Update
  + Commit Boundary
- Split planned work into separate slices when it contains more than one primary user-visible behavior, user trigger, render location, permission or visibility rule, DB write target, DB state transition, API/CLI/MCP contract, async/event/webhook flow, external side effect, cleanup/quarantine domain, behavior test target, independent acceptance gate, or independent commit boundary.
-  Hidden fallback is forbidden. Prefer a visible failure over a fallback that hides a broken primary path.
- When touching DB-backed content, verify the full loop when applicable: UI input -> submit action -> DB write -> DB read after reload/new request -> correct UI render or omission. If there is no UI, replace UI steps with the real caller/consumer flow.
- Tests must prove product behavior. Delete or replace tests that only assert implementation details, helper output, static DOM existence, or mocked plumbing without proving trigger -> process -> observable result.
- If a planned item uses wording such as `and`, `also`, `then wire`, `plus update`, `both`, or `handle all`, check whether it is actually multiple slices.
- Do not write broad actionable items such as `Implement checkout, webhook, entitlement update, and billing UI`; split them into narrow slices such as `Create checkout session request`, `Persist checkout session state`, `Handle provider webhook`, `Update entitlement from webhook event`, and `Render billing status from entitlement`.
- Each slice work step must include UI flow, DB/data flow, render location, and evidence target checks. Use `N/A` with a reason when a check does not apply.
- If tests write DB rows, app state, files, queues, provider state, or other persistent data, the slice must define cleanup or quarantine before implementation.

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

### P1: {{PHASE_1_TITLE}}

- Phase Goal: {{PHASE_1_GOAL}}
- Phase Boundary:
  - In scope: {{PHASE_1_IN_SCOPE}}
  - Out of scope: {{PHASE_1_OUT_OF_SCOPE}}
  - Dependencies: {{PHASE_1_DEPENDENCIES}}
- Phase Implementation Rule: do not implement `P1` directly. Implement `P1-A`, verify it, record evidence, refresh actual-status, commit when required, then continue to `P1-B`.
- Ordered Slice List:
  - P1-A: {{SLICE_1_TITLE}}
  - P1-B: {{SLICE_2_TITLE_OR_REMOVE}}

- [ ] P1-A: {{SLICE_1_TITLE}}
  - Goal: {{SLICE_1_GOAL}}
  - Scope Boundary:
    - Editable: {{SLICE_1_EDITABLE_SURFACES}}
    - Inspect-only: {{SLICE_1_INSPECT_ONLY_SURFACES}}
    - Preserve-only: {{SLICE_1_PRESERVE_ONLY_SURFACES}}
    - Out of scope: {{SLICE_1_OUT_OF_SCOPE}}
  - Non-Goals: {{SLICE_1_NON_GOALS}}
  - Pre-flight Questions:
    - Data source: {{SLICE_1_DATA_SOURCE}}
    - Display permission: {{SLICE_1_DISPLAY_PERMISSION}}
    - DB read flow: {{SLICE_1_DB_READ_FLOW}}
    - DB write flow: {{SLICE_1_DB_WRITE_FLOW}}
    - Render location: {{SLICE_1_RENDER_LOCATION}}
    - UI behavior flow: {{SLICE_1_UI_BEHAVIOR_FLOW}}
    - Docker runtime: {{SLICE_1_DOCKER_RUNTIME}}
    - Playwright target: {{SLICE_1_PLAYWRIGHT_TARGET}}
    - Behavior test: {{SLICE_1_BEHAVIOR_TEST}}
    - Cleanup/quarantine: {{SLICE_1_CLEANUP_OR_QUARANTINE}}
    - External side effects: {{SLICE_1_EXTERNAL_SIDE_EFFECTS}}
    - N/A notes: {{SLICE_1_NA_NOTES}}
  - Work Steps:
    1. {{SLICE_1_WORK_STEP_1}}
       - UI flow check: {{SLICE_1_STEP_1_UI_FLOW_CHECK}}
       - DB/data flow check: {{SLICE_1_STEP_1_DB_FLOW_CHECK}}
       - Render location check: {{SLICE_1_STEP_1_RENDER_LOCATION_CHECK}}
       - Mini QA for each completed implementation slice (MUST) 
          For Codex: When Mini QA must use plugins
            - Browser: Control the in-app browser
            - Chrome: Control the user's real Chrome browser
            - Computer Use: Control Windows apps or installed artifacts when mini QA requires real app interaction outside a browser.
            - Playwright: use as an automation arm for browser actions, control sweeps, screenshots, videos, traces, and reports.
          For Claude or other agents:
            - Use the equivalent browser, Chrome/session, or computer-control capability exposed by that agent/runtime or Playwright-like capability available in that environment.
       - Evidence target: {{SLICE_1_STEP_1_EVIDENCE_TARGET}}
    2. {{SLICE_1_WORK_STEP_2}}
       - UI flow check: {{SLICE_1_STEP_2_UI_FLOW_CHECK}}
       - DB/data flow check: {{SLICE_1_STEP_2_DB_FLOW_CHECK}}
       - Render location check: {{SLICE_1_STEP_2_RENDER_LOCATION_CHECK}}
       - Mini QA for each completed implementation slice (MUST) 
          For Codex: When Mini QA must use plugins
            - Browser: Control the in-app browser
            - Chrome: Control the user's real Chrome browser
            - Computer Use: Control Windows apps or installed artifacts when mini QA requires real app interaction outside a browser.
            - Playwright: use as an automation arm for browser actions, control sweeps, screenshots, videos, traces, and reports.
          For Claude or other agents:
            - Use the equivalent browser, Chrome/session, or computer-control capability exposed by that agent/runtime or Playwright-like capability available in that environment.
       - Evidence target: {{SLICE_1_STEP_2_EVIDENCE_TARGET}}
  - Implementation Gate:
    - Before editing target files, run the relevant Anvien impact/file-detail command for files, symbols, routes, tools, or contracts touched by this slice, and record the evidence IDs.
    - {{SLICE_1_GATE}}
  - Acceptance:
    - Source: {{SLICE_1_ACCEPTANCE_SOURCE}}
    - Runtime/UI: {{SLICE_1_ACCEPTANCE_RUNTIME_UI}}
    - DB/data: {{SLICE_1_ACCEPTANCE_DB_DATA}}
    - Behavior test: {{SLICE_1_ACCEPTANCE_BEHAVIOR_TEST}}
    - Cleanup/quarantine: {{SLICE_1_ACCEPTANCE_CLEANUP_OR_QUARANTINE}}
    - Evidence IDs: {{SLICE_1_ACCEPTANCE_EVIDENCE_IDS}}
    - Actual-status rows refreshed: {{SLICE_1_ACCEPTANCE_ACTUAL_STATUS_ROWS}}
  - Evidence Targets: {{SLICE_1_EVIDENCE_TARGETS}}
  - Actual-status Update: {{SLICE_1_ACTUAL_STATUS_UPDATE}}
  - Commit Boundary: commit after this slice when acceptance passes.
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
    1. Run the required final validation for the accepted scope, including full build before final runtime validation. For app/runtime scopes, full build must include Docker image/container build.
    2. Start the real built Docker/container runtime for app/runtime validation. If Docker cannot be built or started, record the blocker and do not substitute a host dev server.
    3. Validate public runtime or UI-facing changes with browser or Playwright evidence against the real built Docker/container runtime. Playwright evidence must include Docker build/run or compose command, container/service name, exposed URL, Playwright command, and screenshot/trace/result.
    4. Regenerate generated outputs if source-of-truth changes require it.
    5. Run Anvien detect-changes before commit when implementation work was performed.
    6. Record final validation, detect-changes, benchmark, and commit evidence.
    7. Commit the completed scope and verify the worktree state.
  - Implementation Gate: Pn-A and Pn-B must pass or record blockers.
  - Acceptance: final evidence is recorded, required commits exist, and the worktree state is known.

## Risk Notes

{{RISK_NOTES}}
