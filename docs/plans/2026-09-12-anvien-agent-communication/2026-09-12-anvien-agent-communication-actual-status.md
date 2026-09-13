# Anvien Agent Communication Actual Status

Title: Anvien Agent Communication
Date: 2026-09-12
Status: Draft / P0 Complete / Blocked
Companion plan: `docs/plans/2026-09-12-anvien-agent-communication/2026-09-12-anvien-agent-communication-plan.md`
Companion rules: `docs/plans/2026-09-12-anvien-agent-communication/2026-09-12-anvien-agent-communication-plan.md#rules`
Companion evidence: `docs/plans/2026-09-12-anvien-agent-communication/2026-09-12-anvien-agent-communication-evidence.md`
Companion benchmark: `docs/plans/2026-09-12-anvien-agent-communication/2026-09-12-anvien-agent-communication-benchmark.md`

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

- Shared conversation và activation module tích hợp Anvien

Out of scope:

- Production edits, external provider runs, SPEC edits

## Relationship / Impact Evidence

For each target file, prefer:

```text
anvien file-detail <path> --repo <repo> --json
```

Record how many files the target is related to before deciding touch mode. A file with many relationships may still be editable, but the plan must narrow the exact phase, touch mode, and validation needed.

| Unit / File / Surface | File Detail Evidence | Related File Count | Relationship Summary | Impact Note |
|-----------------------|----------------------|--------------------|----------------------|-------------|
| Session/chat integration | E0-P0A-FD1 | Chưa xác minh đầy đủ | HTTP server -> session controller -> Codex adapter; web client -> session API | low / medium / high / critical scope warning |

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
| Session/chat integration | Single adapter, auto-cancel per repo, cancel on browser disconnect | Multi-agent addressed durable communication with independent UI/activation | partial | Chưa xác minh đầy đủ related files | E0-P0A-SRC2, E0-P0A-SRC3, E0-P0A-IMPACT1 | Inspect-only; hoàn thiện P0 trước khi mở slice |

## Status Refresh Log

| Refresh | Date | Repo Basis | Changed Scope | Status Changes | Evidence | Next Phase Update |
|---------|------|------------|----------------|----------------|----------|-------------------|
| R0 | 2026-09-12 | baseline before P0 | Shared conversation và activation module tích hợp Anvien | initial classification | E0-P0A-SRC2, E0-P0A-SRC3, E0-P0A-IMPACT1 | Không mở implementation; chốt contract và provider protocol |

## Phase Touch Map

Use this map to prevent accidental edits. A related file is not automatically editable.

`Plan-Relevant Relationship File` lists only a relationship file that can directly affect or be affected by the planned phase or slice. Do not copy the full `file-detail` relationship inventory into this map. Include only files whose relationship can affect the phase/slice decision, touch mode, or validation.

| Unit / File / Surface | Plan-Relevant Relationship File | Relationship to Target | Plan Item | Touch Mode | Evidence | Constraint |
|-----------------------|---------------------------------|------------------------|-----------|------------|----------|------------|
| internal/session/controller.go | internal/httpapi/server.go; internal/httpapi/session.go; internal/session/codex.go | caller / API / adapter | P0-A | inspect-only | E0-P0A-SRC2, E0-P0A-IMPACT1 | Không sửa lifecycle chat cũ trong khảo sát |
| Web chat | anvien-web/src/hooks/chat-runtime/ChatRuntimeContext.tsx; anvien-web/src/components/ChatPanel.tsx | consumer / UI | P0-A | inspect-only | E0-P0A-SRC3 | Chỉ xác định ứng viên tái sử dụng, không đổi UI |

## Detailed Findings

### Session/controller

Current state:

Source cho thấy lifecycle chat cũ không phải multi-agent mailbox.

Required state:

```text
Agent identity độc lập role/provider; persistent messages + activation; không auto-cancel theo repo.
```

Evidence:

- E0-P0A-SRC2: Backend source inspected
- E0-P0A-SRC3: Frontend source inspected

Relationship and impact:

- Related file count: Chưa xác minh đầy đủ
- Relationship summary: HTTP server -> session controller -> Codex adapter; web client -> session API
- Impact note: CRITICAL; 8 affected files; new communication module preferred

Classification:

partial đối với mục tiêu mới

Allowed next action:

Hoàn thiện P0, không sửa production

Forbidden next action:

Coi phase roadmap là slice ready-to-code

## Next Phase Status Decisions

| Plan Item | Actual Status Finding | Required Status / Next-Action Update |
|-----------|-----------------------|--------------------------------------|
| P1-A | Protocol và persistence chưa chốt | keep / change / remove / block |

## Implementation Gate

- [ ] Target scope is listed in Current Status Matrix.
- [ ] Each target unit has a status.
- [ ] Each status has evidence IDs.
- [ ] Each target file has relationship count evidence from `file-detail` when applicable.
- [ ] Phase Touch Map lists plan-relevant relationship files that can affect the current phase/slice.
- [ ] Phase Touch Map defines touch mode for every plan-relevant relationship unit that may be affected.
- [ ] Correct parts are marked preserve-only.
- [ ] Partial, missing, wrong, unbound, and fake-or-stub parts have exact next actions.
- [ ] Blockers are recorded, if any.
- [ ] Next phase status assumptions, next action, and work steps have been updated from this status file when needed.
- [ ] Status Refresh Log has an R0 baseline row.
- [ ] If implementation has started, affected Current Status Matrix rows have been refreshed from latest evidence.
- [ ] If refreshed statuses changed next work, only the stale next-phase status assumptions, next action, or work steps have been updated before the next phase.

## Final P0 Decision

Choose one:

- [x] P0 actual-status incomplete. Implementation is blocked.
- [ ] P0 complete. Next phase can proceed unchanged.
- [ ] P0 complete. Next phase status, next action, or work steps must be updated before implementation.
- [ ] P0 complete. Target scope is preserve-only.
- [ ] P0 complete. Implementation is blocked by missing authority or evidence.

Decision note:

P0 chưa hoàn tất. Có đủ căn cứ đề xuất hướng, chưa đủ giao coder implement.
