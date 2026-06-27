# Spec-to-SVG Flow Map Skill Refresh Plan

## Metadata

- Date: `2026-06-27`
- Status: `active`
- Plan: `docs/plans/2026-06-27-spec-to-svg-flow-map-skill-refresh/2026-06-27-spec-to-svg-flow-map-skill-refresh-plan.md`
- Evidence: `docs/plans/2026-06-27-spec-to-svg-flow-map-skill-refresh/2026-06-27-spec-to-svg-flow-map-skill-refresh-evidence.md`
- Benchmark: `docs/plans/2026-06-27-spec-to-svg-flow-map-skill-refresh/2026-06-27-spec-to-svg-flow-map-skill-refresh-benchmark.md`
- Actual status: `docs/plans/2026-06-27-spec-to-svg-flow-map-skill-refresh/2026-06-27-spec-to-svg-flow-map-skill-refresh-actual-status.md`

## Goal

Refresh `internal/aicontext/skills/Spec-to-SVG-Flow-Map/SKILL.md` by adding valid skill frontmatter, standardizing generated output paths, and expanding the skill with detail-completeness, source-union inventory, flow-by-flow rendering, no-bulk-drawing, no-collapse, source-coverage metadata, and verification additions while preserving the semantic SVG and gap-detection contract.

## Rules

- Complete P0 actual status before implementation work.
- Update each checklist item immediately when it is completed.
- Record evidence as work completes.
- Record benchmarkable counts or measurements when they are taken.
- Update later phase status assumptions, next actions, and work steps when actual-status evidence changes the repo state.
- After completing a phase or implementation slice and refreshing `actual-status.md`, update the next affected phase's work steps as needed to match the latest repo reality, while preserving that phase's original goal, scope, acceptance criteria, and major phase order.
- Run Anvien detect-changes before every implementation-slice commit when implementation work was performed.
- For this docs-only skill refresh, runtime/UI, DB, Docker, and Playwright checks are not applicable.
- Keep the standard planner structure. These detail rules only make phase checklist items concrete enough to implement safely.
- Every implementation phase must be decomposed into slices. This plan has one implementation slice because the user requested one atomic text refresh to one skill file.

## Problem

The current skill needs a valid skill contract plus stricter completeness rules so generated SVG flow maps cannot be overview-only, less detailed than reference diagrams, or missing source inventory coverage.

## Scope

- Edit `internal/aicontext/skills/Spec-to-SVG-Flow-Map/SKILL.md`.
- Create this required plan set and record evidence.
- Validate the updated markdown and skill structure.
- Run Anvien detect-changes before commit.

## Non-Goals

- Do not implement app source code.
- Do not create generated SVG flow maps.
- Do not change unrelated skill files.
- Do not edit the pre-existing `.dockerignore` worktree change.

## Requirements

- Add valid YAML frontmatter with only `name` and `description`.
- Use one output directory contract: `docs/flow-maps/`.
- Keep required output files as `<feature-name>.flow.svg`, `<feature-name>.flow-map.md`, and `<feature-name>.flow-verification.md`.
- Keep the semantic SVG node/edge metadata requirements.
- Keep explicit gap, decision, junction, terminal state, and owner review rules.
- Add detail-completeness, source-union inventory, flow-by-flow rendering, no-bulk-drawing, no-collapse, minimum-detail, domain-detail, source-coverage metadata, and verification additions.
- Do not limit this special skill to fewer than 500 lines; preserve the full required instruction set.

## Acceptance Criteria

- `SKILL.md` begins with valid YAML frontmatter delimiters.
- All output path examples use `docs/flow-maps/`.
- The new detail-completeness and source-coverage contracts are present.
- The special skill is allowed to exceed 500 lines when needed to preserve required behavior.
- Validation evidence records source inspection and a parse/check result.
- Anvien detect-changes is run before commit.
- The final commit includes only the plan set and the skill refresh unless pre-existing unrelated work remains outside the commit.

## Checklist

- [x] P0-A: Complete actual status before implementation work.
  - Goal: establish the real current state for the target skill.
  - Work Steps: inspect the target skill, refresh Anvien graph, run file-detail, classify current state, and record the allowed next action.
  - Implementation Gate: no implementation or editing starts until the actual-status file has a final P0 decision.
  - Acceptance: actual status identifies the target as partial, low risk, and editable for the requested text refresh.

### P1: Refresh Skill Instructions

- Phase Goal: update the skill file to be valid, consistent, and shorter while preserving behavior.
- Phase Boundary:
  - In scope: the target `SKILL.md` text contract.
  - Out of scope: app source code, generated SVG outputs, unrelated docs, and unrelated worktree changes.
  - Dependencies: P0 actual status and user request.
- Phase Implementation Rule: do not implement `P1` directly. Implement `P1-A`, verify it, record evidence, refresh actual-status, commit when required, then close.
- Ordered Slice List:
  - P1-A: Refresh `SKILL.md` frontmatter, output paths, and detail-completeness coverage rules.

- [x] P1-A: Refresh `SKILL.md` frontmatter, output paths, and detail-completeness coverage rules.
  - Goal: produce a valid special-purpose skill file matching the requested expanded coverage rules, without imposing a 500-line cap.
  - Scope Boundary:
    - Editable: `internal/aicontext/skills/Spec-to-SVG-Flow-Map/SKILL.md`.
    - Inspect-only: this plan set, existing git status, Anvien file-detail evidence.
    - Preserve-only: `.dockerignore` and unrelated files.
    - Out of scope: app source code and generated SVG flow map artifacts.
  - Non-Goals: no runtime behavior changes, no generated SVG output creation, no broad skill redesign.
  - Pre-flight Questions:
    - Data source: user request and current skill file.
    - Display permission: N/A, no UI.
    - DB read flow: N/A, docs-only change.
    - DB write flow: N/A, docs-only change.
    - Render location: N/A, no runtime render target.
    - UI behavior flow: N/A, no UI.
    - Docker runtime: N/A, no app/runtime change.
    - Playwright target: N/A, no UI.
    - Behavior test: validate frontmatter/path/content contract by source inspection and available parser/check commands.
    - Cleanup/quarantine: no temp files outside repo; no generated dead artifacts.
    - External side effects: git commit only after validation.
    - N/A notes: this is a markdown skill refresh, not production code.
  - Work Steps:
    1. Rewrite the target skill with valid frontmatter, one output path contract, and the requested detail-completeness/source-union/flow-by-flow/no-collapse sections.
       - UI flow check: N/A, no UI.
       - DB/data flow check: N/A, markdown-only.
       - Render location check: N/A, markdown-only.
       - Mini QA for each completed implementation slice: validate by file read, path search, and frontmatter check.
       - Evidence target: `E1-P1A-SRC1`.
    2. Validate the updated file and record exact evidence.
       - UI flow check: N/A, no UI.
       - DB/data flow check: N/A, markdown-only.
       - Render location check: N/A, markdown-only.
       - Mini QA for each completed implementation slice: run focused validation commands.
       - Evidence target: `E1-P1A-VAL1`.
  - Implementation Gate:
    - Before editing target files, record Anvien file-detail evidence for the target file.
    - P0 actual-status marks the target editable and low risk.
  - Acceptance:
    - Source: `SKILL.md` has valid YAML delimiters, unified `docs/flow-maps/` outputs, and the requested detail-completeness/source-union/flow-by-flow/no-collapse contracts.
    - Runtime/UI: N/A, docs-only skill change.
    - DB/data: N/A, docs-only skill change.
    - Behavior test: validation commands pass.
    - Cleanup/quarantine: no dead temp or generated artifacts remain.
    - Evidence IDs: `E1-P1A-SRC1`, `E1-P1A-VAL1`, `E2-P2A-DETECT1`.
    - Actual-status rows refreshed: target skill row changes from `partial` to `correct`.
  - Evidence Targets: updated source content, validation result, detect-changes result, commit hash.
  - Actual-status Update: refresh the target row after validation.
  - Commit Boundary: commit after this slice when acceptance passes.

- [x] Pn-A: Call supervisor for the implemented-plan acceptance loop.
  - Goal: verify the completed plan work against the accepted plan, actual-status decisions, evidence, benchmark, changed files, generated output, and validation results before closure.
  - Work Steps:
    1. Review the completed plan work against the requested scope and evidence.
    2. If review fails, fix only the failed scope and re-review.
  - Implementation Gate: P1-A must be completed or explicitly blocked.
  - Acceptance: review passes, or the plan records a blocker with evidence and no closure is performed.
- [x] Pn-B: Remove dead work created during this plan.
  - Goal: ensure the final diff contains only artifacts that still serve the accepted plan.
  - Work Steps:
    1. Review files and sections created or modified during this plan.
    2. Remove or rewrite obsolete artifacts created by this plan.
    3. Verify final diff/status contains no dead plan-created artifacts.
  - Implementation Gate: only remove artifacts created by this plan unless the user explicitly approves broader cleanup.
  - Acceptance: final `git diff/status` contains no dead plan-created artifacts, and evidence records what was preserved.
- [x] Pn-C: Close the plan.
  - Goal: finish validation, evidence, benchmark, detect-changes, commit, and final status.
  - Work Steps:
    1. Record final validation. Full app build is N/A because the slice changes only markdown skill instructions.
    2. Run Anvien detect-changes before commit.
    3. Record final validation, detect-changes, benchmark, and commit evidence.
    4. Commit the completed scope and verify the worktree state.
  - Implementation Gate: Pn-A and Pn-B must pass or record blockers.
  - Acceptance: final evidence is recorded, required commit exists, and the worktree state is known.

## Risk Notes

- The target file is currently uncommitted/untracked in git status, so the commit must avoid unrelated `.dockerignore` changes.
- This skill controls future agent behavior, so concision must not remove semantic SVG, gap detection, or owner-approval constraints.
