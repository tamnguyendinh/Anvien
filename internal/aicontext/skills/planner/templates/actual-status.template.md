# {{TITLE}} Actual Status

Title: {{TITLE}}
Date: {{YYYY-MM-DD}}
Status: Draft / P0 Complete / Blocked
Companion plan: `{{PLAN_PATH}}`
Companion evidence: `{{EVIDENCE_PATH}}`
Companion benchmark: `{{BENCHMARK_PATH}}`

## Purpose

This file records the real current state before implementation.

Implementation must not start until the target scope has a completed status row, evidence IDs, and a downstream plan decision.

This file does not replace `evidence.md`. It classifies current state from evidence.

Use exact evidence IDs from `evidence.md`, such as `E0-P0A-SRC1`, not broad section IDs such as `E0` or `E1`.

## Freshness / Refresh Rules

This actual-status file is a living current-state record, not a one-time P0 snapshot.

P0 records the baseline before implementation. After implementation begins, keep the Current Status Matrix updated so the next agent can trust it as the latest repo reality.

Update this file:

- after each completed implementation slice;
- before starting the next phase if repo state changed;
- whenever evidence changes a current-state classification;
- whenever the next phase's status assumptions, next action, or work steps need updating because reality differs from the previous status.

When refreshing status:

- update only the rows affected by the completed work or new evidence;
- use explicit transitions such as `missing -> correct`, `partial -> correct`, `fake-or-stub -> removed`, or `unbound -> bound-correct`;
- append a Status Refresh Log row instead of deleting history;
- keep detailed proof in `evidence.md`; store only classifications, evidence IDs, touch mode, and plan consequences here.

## Scope

Target scope:

- {{TARGET_SCOPE}}

Out of scope:

- {{OUT_OF_SCOPE}}

## Relationship / Impact Evidence

For each target file, prefer:

```text
anvien file-detail <path> --repo <repo> --json
```

Record how many files the target is related to before deciding touch mode. A file with many relationships may still be editable, but the plan must narrow the exact phase, touch mode, and validation needed.

`Plan-Relevant Relationship Files` lists only relationship files that can directly affect or be affected by the planned phase or slice. Do not list every related file from `file-detail`. If `file-detail` shows 100 related files but only 5 can matter for the current plan slice, list only those 5 and point to the evidence ID for the full relationship inventory.

| Unit / File / Surface | File Detail Evidence | Related File Count | Relationship Summary | Plan-Relevant Relationship Files | Impact Note |
|-----------------------|----------------------|--------------------|----------------------|----------------------------------|-------------|
| {{UNIT}} | {{EVIDENCE_ID}} | {{RELATED_FILE_COUNT}} | {{RELATIONSHIP_SUMMARY}} | {{PLAN_RELEVANT_RELATIONSHIP_FILES}} | low / medium / high / critical scope warning |

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
| {{UNIT}} | {{CURRENT_STATE}} | {{REQUIRED_STATE}} | correct/partial/wrong/missing/unbound/fake-or-stub/blocked | {{RELATED_FILE_COUNT}} related files | {{EVIDENCE_IDS}} | preserve / edit P1-A / update P2-B status / block |

## Status Refresh Log

| Refresh | Date | Repo Basis | Changed Scope | Status Changes | Evidence | Next Phase Update |
|---------|------|------------|----------------|----------------|----------|-------------------|
| R0 | {{YYYY-MM-DD}} | baseline before P0 | {{TARGET_SCOPE}} | initial classification | {{EVIDENCE_IDS}} | {{NEXT_PHASE_STATUS_UPDATE}} |
| R1 | {{YYYY-MM-DD}} | after {{COMPLETED_PLAN_ITEM_OR_COMMIT}} | {{CHANGED_SCOPE}} | {{STATUS_TRANSITION}} | {{EVIDENCE_IDS}} | {{NEXT_PHASE_UPDATE}} |

## Phase Touch Map

Use this map to prevent accidental edits. A related file is not automatically editable.

| Unit / File / Surface | Relationship to Target | Related File Count | Plan Item | Touch Mode | Evidence | Constraint |
|-----------------------|------------------------|--------------------|-----------|------------|----------|------------|
| {{UNIT}} | source-of-truth / consumer / generated output / test / config / dependency | {{RELATED_FILE_COUNT}} | P1-A | preserve-only / inspect-only / edit / regenerate / validate-only / block / do-not-touch | {{EVIDENCE_IDS}} | {{CONSTRAINT}} |

## Detailed Findings

### {{UNIT_NAME}}

Current state:

{{CURRENT_STATE_FACTS}}

Required state:

```text
{{REQUIRED_STATE_TEXT}}
```

Evidence:

- {{EVIDENCE_ID_1}}: {{EVIDENCE_1}}
- {{EVIDENCE_ID_2}}: {{EVIDENCE_2}}

Relationship and impact:

- Related file count: {{RELATED_FILE_COUNT}}
- Relationship summary: {{RELATIONSHIP_SUMMARY}}
- Plan-relevant relationship files: {{PLAN_RELEVANT_RELATIONSHIP_FILES}}
- Impact note: {{IMPACT_NOTE}}

Classification:

{{STATUS}}

Allowed next action:

{{ALLOWED_NEXT_ACTION}}

Forbidden next action:

{{FORBIDDEN_NEXT_ACTION}}

## Next Phase Status Decisions

| Plan Item | Actual Status Finding | Required Status / Next-Action Update |
|-----------|-----------------------|--------------------------------------|
| P1-A | {{FINDING}} | keep / change / remove / block |

## Implementation Gate

- [ ] Target scope is listed in Current Status Matrix.
- [ ] Each target unit has a status.
- [ ] Each status has evidence IDs.
- [ ] Each target file has relationship count evidence from `file-detail` when applicable.
- [ ] Plan-relevant relationship files are listed for each target file when relationship evidence shows direct phase/slice impact.
- [ ] Phase Touch Map defines touch mode for every related unit that may be affected.
- [ ] Correct parts are marked preserve-only.
- [ ] Partial, missing, wrong, unbound, and fake-or-stub parts have exact next actions.
- [ ] Blockers are recorded, if any.
- [ ] Next phase status assumptions, next action, and work steps have been updated from this status file when needed.
- [ ] Status Refresh Log has an R0 baseline row.
- [ ] If implementation has started, affected Current Status Matrix rows have been refreshed from latest evidence.
- [ ] If refreshed statuses changed next work, only the stale next-phase status assumptions, next action, or work steps have been updated before the next phase.

## Final P0 Decision

Choose one:

- [ ] P0 actual-status incomplete. Implementation is blocked.
- [ ] P0 complete. Next phase can proceed unchanged.
- [ ] P0 complete. Next phase status, next action, or work steps must be updated before implementation.
- [ ] P0 complete. Target scope is preserve-only.
- [ ] P0 complete. Implementation is blocked by missing authority or evidence.

Decision note:

{{DECISION_NOTE}}
