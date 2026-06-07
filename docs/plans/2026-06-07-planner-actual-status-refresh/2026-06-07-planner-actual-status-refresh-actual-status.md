# Planner Actual Status Refresh Actual Status

Title: Planner Actual Status Refresh
Date: 2026-06-07
Status: P0 Complete
Companion plan: `docs/plans/2026-06-07-planner-actual-status-refresh/2026-06-07-planner-actual-status-refresh-plan.md`
Companion evidence: `docs/plans/2026-06-07-planner-actual-status-refresh/2026-06-07-planner-actual-status-refresh-evidence.md`
Companion benchmark: `docs/plans/2026-06-07-planner-actual-status-refresh/2026-06-07-planner-actual-status-refresh-benchmark.md`

## Purpose

This file records the real current state before implementation.

Implementation must not start until the target scope has a completed status row, evidence IDs, and a downstream plan decision.

This file does not replace `evidence.md`. It classifies current state from evidence.

## Scope

Target scope:

- `internal/aicontext/skills/planner/SKILL.md`
- `internal/aicontext/skills/planner/templates/plan.template.md`
- `internal/aicontext/skills/planner/templates/actual-status.template.md`

Out of scope:

- Unrelated dirty worktree changes under `internal/aicontext/skills/databases`.
- Unrelated dirty worktree changes under `internal/aicontext/skills/skill-creator`.
- Generated `AGENTS.md` / `CLAUDE.md` content.

## Relationship / Impact Evidence

| Unit / File / Surface | File Detail Evidence | Related File Count | Relationship Summary | Impact Note |
|-----------------------|----------------------|--------------------|----------------------|-------------|
| `internal/aicontext/skills/planner/SKILL.md` | E0.2 | 0 | No inbound, outbound, local relationships, linked flows, or linked tests. | low scope warning |
| `internal/aicontext/skills/planner/templates/actual-status.template.md` | E0.3 | 0 | No inbound, outbound, local relationships, linked flows, or linked tests. | low scope warning |
| `internal/aicontext/skills/planner/templates/plan.template.md` | E0.7 | 0 | No inbound, outbound, local relationships, linked flows, or linked tests. | low scope warning |

## Status Rules

| Status | Meaning | Allowed next action |
|--------|---------|---------------------|
| `correct` | Already behaves as required. | Preserve. Add evidence or tests only if needed. |
| `partial` | Some required behavior exists, but gaps remain. | Change only the missing parts. Preserve correct parts. |
| `wrong` | Current behavior, source, or contract is incorrect. | Replace with required behavior. Record the exact reason. |
| `missing` | Required behavior, source, or contract does not exist. | Implement the missing piece only. |
| `unbound` | Surface exists but is not wired to the real source, flow, or contract. | Bind to the real source only. Preserve approved surface. |
| `fake-or-stub` | Prototype, demo, mock, fallback, or placeholder data is being used as real behavior. | Remove fake behavior or replace it with an approved truthful state. |
| `blocked` | Source, authority, contract, or required evidence is unclear. | Stop. Do not implement until resolved. |

## Current Status Matrix

| Unit | Current State | Required State | Status | Relationship Count | Evidence | Next Plan Decision |
|------|---------------|----------------|--------|--------------------|----------|--------------------|
| `internal/aicontext/skills/planner/SKILL.md` | Describes actual-status as P0 current-state evidence and says not to execute stale phases, but does not explicitly require ongoing refresh after each implementation slice. | Skill workflow must tell agents to keep actual-status current as repo reality changes during implementation. | partial | 0 related files | E0.2, E0.4 | edit P1-A |
| `internal/aicontext/skills/planner/templates/actual-status.template.md` | Provides P0 baseline structure, status matrix, rewrite decisions, and implementation gate, but lacks freshness rules and a refresh log. | Template must include living-status freshness rules and a status refresh log table. | partial | 0 related files | E0.3, E0.5 | edit P1-A |
| `internal/aicontext/skills/planner/templates/plan.template.md` | Contains broad "rewrite later phases" wording that can make agents change plan goals instead of updating stale phase state. | Template must describe updating later phase status assumptions, next actions, and work steps. | partial | 0 related files | E0.7, E0.8 | edit P1-A |

## Phase Touch Map

| Unit / File / Surface | Relationship to Target | Related File Count | Plan Item | Touch Mode | Evidence | Constraint |
|-----------------------|------------------------|--------------------|-----------|------------|----------|------------|
| `internal/aicontext/skills/planner/SKILL.md` | source-of-truth skill instructions | 0 | P1-A | edit | E0.2, E0.4 | Change only actual-status refresh guidance. |
| `internal/aicontext/skills/planner/templates/actual-status.template.md` | source-of-truth generated actual-status template | 0 | P1-A | edit | E0.3, E0.5 | Add freshness rules and refresh log without weakening P0 gate. |
| `internal/aicontext/skills/planner/templates/plan.template.md` | source-of-truth generated plan template | 0 | P1-A | edit | E0.7, E0.8 | Replace broad rewrite wording with status/next-action/work-step update wording. |
| `internal/aicontext/skills/databases/*` | unrelated dirty worktree | not measured | none | do-not-touch | E0.6 | Do not revert or include in commit. |
| `internal/aicontext/skills/skill-creator/*` | unrelated dirty worktree | not measured | none | do-not-touch | E0.6 | Do not revert or include in commit. |

## Detailed Findings

### Planner SKILL.md

Current state:

The skill says actual-status records true current state before implementation and requires phase updates when evidence changes scope. It does not explicitly define actual-status as a living status record after implementation starts.

Required state:

```text
Planner workflow must require actual-status refresh after implementation slices and before later phases when repo state changes, while keeping detailed proof in evidence.md.
```

Evidence:

- E0.2: File-detail evidence for `internal/aicontext/skills/planner/SKILL.md`.
- E0.4: File-level impact evidence for `internal/aicontext/skills/planner/SKILL.md`.

Relationship and impact:

- Related file count: 0
- Relationship summary: no graph relationships or linked flows/tests.
- Impact note: LOW blast radius.

Classification:

partial

Allowed next action:

Edit P1-A to add refresh workflow requirements.

Forbidden next action:

Do not change unrelated skill behavior or generated agent rules.

### Actual Status Template

Current state:

The template provides a P0 baseline and implementation gate, but no explicit freshness section or status refresh log.

Required state:

```text
The template must state that actual-status is living current-state context, include triggers for refresh, and provide a markdown status refresh log.
```

Evidence:

- E0.3: File-detail evidence for `internal/aicontext/skills/planner/templates/actual-status.template.md`.
- E0.5: File-level impact evidence for `internal/aicontext/skills/planner/templates/actual-status.template.md`.

Relationship and impact:

- Related file count: 0
- Relationship summary: no graph relationships or linked flows/tests.
- Impact note: LOW blast radius.

Classification:

partial

Allowed next action:

Edit P1-A to add freshness rules and refresh log.

Forbidden next action:

Do not remove existing status matrix, relationship evidence, next-phase status decisions, or implementation gate.

### Plan Template

Current state:

The template includes broad "rewrite later phases" wording in the rules and P0 work steps.

Required state:

```text
The template must say to update later phase status assumptions, next actions, and work steps from actual-status evidence.
```

Evidence:

- E0.7: File-detail evidence for `internal/aicontext/skills/planner/templates/plan.template.md`.
- E0.8: File-level impact evidence for `internal/aicontext/skills/planner/templates/plan.template.md`.

Relationship and impact:

- Related file count: 0
- Relationship summary: no graph relationships or linked flows/tests.
- Impact note: LOW blast radius.

Classification:

partial

Allowed next action:

Edit P1-A to remove broad "rewrite phase" wording.

Forbidden next action:

Do not change plan goals or unrelated template sections.

## Next Phase Status Decisions

| Plan Item | Actual Status Finding | Required Status / Next-Action Update |
|-----------|-----------------------|--------------------------------------|
| P1-A | All target files are partial and low risk. | keep |

## Implementation Gate

- [x] Target scope is listed in Current Status Matrix.
- [x] Each target unit has a status.
- [x] Each status has evidence IDs.
- [x] Each target file has relationship count evidence from `file-detail` when applicable.
- [x] Phase Touch Map defines touch mode for every related unit that may be affected.
- [x] Correct parts are marked preserve-only.
- [x] Partial, missing, wrong, unbound, and fake-or-stub parts have exact next actions.
- [x] Blockers are recorded, if any.
- [x] Next phase status assumptions, next action, and work steps have been updated from this status file when needed.

## Final P0 Decision

Choose one:

- [ ] P0 actual-status incomplete. Implementation is blocked.
- [ ] P0 complete. Next phase can proceed unchanged.
- [x] P0 complete. Next phase status, next action, or work steps must be updated before implementation.
- [ ] P0 complete. Target scope is preserve-only.
- [ ] P0 complete. Implementation is blocked by missing authority or evidence.

Decision note:

P1-A should proceed with a narrow documentation/template edit for living actual-status refresh behavior.
