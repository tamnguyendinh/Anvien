# Actual Status

Title: SePay Skill Refresh
Date: 2026-06-14
Status: P0 Complete
Companion plan: `docs/plans/2026-06-14-sepay-skill-refresh/2026-06-14-sepay-skill-refresh-plan.md`
Companion evidence: `docs/plans/2026-06-14-sepay-skill-refresh/2026-06-14-sepay-skill-refresh-evidence.md`
Companion benchmark: `docs/plans/2026-06-14-sepay-skill-refresh/2026-06-14-sepay-skill-refresh-benchmark.md`

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
- use explicit transitions such as `missing -> correct`, `partial -> correct`, `wrong -> correct`, or `unbound -> bound-correct`;
- append a Status Refresh Log row instead of deleting history;
- keep detailed proof in `evidence.md`; store only classifications, evidence IDs, touch mode, and plan consequences here.

## Scope

Target scope:

- `internal/aicontext/skills/payment-integration/references/sepay/overview.md`
- `internal/aicontext/skills/payment-integration/references/sepay/api.md`
- `internal/aicontext/skills/payment-integration/references/sepay/webhooks.md`
- `internal/aicontext/skills/payment-integration/references/sepay/sdk.md`
- `internal/aicontext/skills/payment-integration/references/sepay/qr-codes.md`
- `internal/aicontext/skills/payment-integration/references/sepay/best-practices.md`
- `internal/aicontext/skills/payment-integration/references/sepay/payment-gateway.md` if added by P1-A
- minimal routing updates in `internal/aicontext/skills/payment-integration/SKILL.md` and `internal/aicontext/skills/payment-integration/references/implementation-workflows.md` if `payment-gateway.md` is added

Out of scope:

- `internal/aicontext/skills/payment-integration/scripts/sepay-webhook-verify.js` implementation changes unless the user explicitly expands scope.
- Other payment provider references.
- Generated `.agents/**`, `.claude/**`, `AGENTS.md`, or `CLAUDE.md` as source edits.
- Historical reports/plans/evidence.

## Relationship / Impact Evidence

For each target file, prefer:

```text
anvien file-detail <path> --repo <repo> --json
```

Record how many files the target is related to before deciding touch mode. A file with many relationships may still be editable, but the plan must narrow the exact phase, touch mode, and validation needed.

| Unit / File / Surface | File Detail Evidence | Related File Count | Relationship Summary | Impact Note |
|-----------------------|----------------------|--------------------|----------------------|-------------|
| `internal/aicontext/skills/payment-integration/references/sepay/overview.md` | `E0-P0A-FD1` | 0 related files | Markdown docs; no local, inbound, outbound, linked flow, or linked test relationships. | low scope warning |
| `internal/aicontext/skills/payment-integration/references/sepay/api.md` | `E0-P0A-FD2` | 0 related files | Markdown docs; no local, inbound, outbound, linked flow, or linked test relationships. | low scope warning |
| `internal/aicontext/skills/payment-integration/references/sepay/webhooks.md` | `E0-P0A-FD3` | 0 related files | Markdown docs; no local, inbound, outbound, linked flow, or linked test relationships. | low scope warning |
| `internal/aicontext/skills/payment-integration/references/sepay/sdk.md` | `E0-P0A-FD4` | 0 related files | Markdown docs; no local, inbound, outbound, linked flow, or linked test relationships. | low scope warning |
| `internal/aicontext/skills/payment-integration/references/sepay/qr-codes.md` | `E0-P0A-FD5` | 0 related files | Markdown docs; no local, inbound, outbound, linked flow, or linked test relationships. | low scope warning |
| `internal/aicontext/skills/payment-integration/references/sepay/best-practices.md` | `E0-P0A-FD6` | 0 related files | Markdown docs; no local, inbound, outbound, linked flow, or linked test relationships. | low scope warning |
| `internal/aicontext/skills/payment-integration/SKILL.md` | `E0-P0A-FD7` | 0 related files | Markdown docs; route surface for SePay reference discoverability. | low scope warning |
| `internal/aicontext/skills/payment-integration/references/implementation-workflows.md` | `E0-P0A-FD8` | 0 related files | Markdown docs; workflow route surface for SePay implementation guidance. | low scope warning |
| `internal/aicontext/skills/payment-integration/scripts/sepay-webhook-verify.js` | `E0-P0A-FD9` | 1 inbound related file plus local symbol relationships | JavaScript source; related test file `scripts/test-scripts.js`; 15 local symbol relationships and 4 inbound references. | high scope warning; inspect-only/follow-up |

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
| `overview.md` | Defines SePay broadly but mixes product surfaces and includes stale gateway endpoints and rate limit. | Surface router with current API v2, Webhook, Payment Gateway/IPN, QR, Order VA, Test Mode, and official-doc link guidance. | wrong | 0 related files | `E0-P0A-SRC1`, `E0-P0A-FD1`, `E0-P0A-REPORT1` | edit P1-A |
| `api.md` | Uses legacy `https://my.sepay.vn/userapi/` and v1 transaction/account examples as default. | API v2 default, v1 legacy/migration, current rate limit, response envelope, pagination, and Order VA bank rules. | wrong | 0 related files | `E0-P0A-SRC2`, `E0-P0A-FD2`, `E0-P0A-REPORT1` | edit P2-A |
| `webhooks.md` | Covers basic webhook setup, payload, API Key/OAuth2/none, 200 success, and duplicate prevention, but misses HMAC/payment code/monitoring/replay and has stale retry timing. | Current bank-account webhook guidance with HMAC, raw body, retry correction, payment code, monitoring, replay, IP allowlist, and explicit IPN separation. | wrong | 0 related files | `E0-P0A-SRC3`, `E0-P0A-FD3`, `E0-P0A-REPORT1` | edit P3-A |
| `sdk.md` | Uses old Node install source, stale gateway endpoints, limited SDK/gateway lifecycle details, and ambiguous Laravel officialness. | Current Node/PHP SDK guidance, gateway base URLs/auth, payment methods, form field order, order detail/status/cancel/void, and accurate Laravel classification. | wrong | 0 related files | `E0-P0A-SRC4`, `E0-P0A-FD4`, `E0-P0A-REPORT1` | edit P4-A |
| `qr-codes.md` | Provides working basic QR parameter guidance and examples, but misses newer parameters and bank-specific VA/memo rules. | Current QR utility guidance with missing parameters, bank identifiers, official VA, memo-based VA, no VA, and VietinBank `SEVQR` examples. | partial | 0 related files | `E0-P0A-SRC5`, `E0-P0A-FD5`, `E0-P0A-REPORT1` | edit P5-A |
| `best-practices.md` | Contains many project-specific patterns (`CLAUDEKIT`, Polar, GitHub, discounts, product pricing) as if they are general SePay guidance. | Repo-agnostic SePay production patterns for surface choice, HMAC, payment code, idempotency, replay, reconciliation, Test Mode/Live, IP allowlisting, and amount policy. | wrong | 0 related files | `E0-P0A-SRC6`, `E0-P0A-FD6`, `E0-P0A-REPORT1` | edit P5-A |
| `payment-gateway.md` | Missing. Hosted checkout/IPN details are underrepresented and mixed into overview/sdk/webhook context. | Dedicated Payment Gateway/IPN reference under `references/sepay`. | missing | not applicable until file exists | `E0-P0A-REPORT1`, `E0-P0A-ROUTE1` | create P1-A, fill P4-A |
| `SKILL.md` SePay quick reference | Lists six SePay reference files, no dedicated gateway/IPN reference. | If `payment-gateway.md` is added, list it so agents can discover hosted checkout/IPN guidance. | partial | 0 related files | `E0-P0A-ROUTE1`, `E0-P0A-FD7` | edit P1-A only for routing |
| `implementation-workflows.md` SePay workflow | Routes SePay agents to overview, API/SDK, webhooks, verifier script, and best practices; no gateway/IPN decision point. | If `payment-gateway.md` is added, route hosted checkout/IPN work to it before SDK/webhook details. | partial | 0 related files | `E0-P0A-ROUTE1`, `E0-P0A-FD8` | edit P1-A only for routing |
| `scripts/sepay-webhook-verify.js` | Helper supports API Key/OAuth2/none and payload validation, but no HMAC raw-body signature path. | Out of current implementation scope; should remain inspect-only/follow-up unless user expands scope. | partial | 1 inbound related file plus local symbol relationships | `E0-P0A-SRC7`, `E0-P0A-FD9` | inspect-only / do not edit |

## Status Refresh Log

| Refresh | Date | Repo Basis | Changed Scope | Status Changes | Evidence | Next Phase Update |
|---------|------|------------|----------------|----------------|----------|-------------------|
| R0 | 2026-06-14 | baseline before implementation | SePay reference files and direct routing/helper surfaces | initial classification | `E0-P0A-GRAPH1`, `E0-P0A-REPORT1`, `E0-P0A-SRC1..E0-P0A-SRC7`, `E0-P0A-FD1..E0-P0A-FD9` | P1-A creates surface map and gateway route; P2-A through P5-A keep goals but use latest status rows. |

## Phase Touch Map

Use this map to prevent accidental edits. A related file is not automatically editable.

`Plan-Relevant Relationship File` lists only a relationship file that can directly affect or be affected by the planned phase or slice. Do not copy a full relationship inventory into this map. Include only files whose relationship can affect the phase/slice decision, touch mode, or validation.

| Unit / File / Surface | Plan-Relevant Relationship File | Relationship to Target | Plan Item | Touch Mode | Evidence | Constraint |
|-----------------------|---------------------------------|------------------------|-----------|------------|----------|------------|
| `overview.md` | `reports/problem/2026-06-14-sepay-skill-refresh-analysis.md` | authority report | P1-A | inspect-only | `E0-P0A-REPORT1` | Use report findings; re-check official docs if delayed. |
| `overview.md` | `internal/aicontext/skills/payment-integration/SKILL.md` | routing/discoverability | P1-A | edit only if adding `payment-gateway.md` | `E0-P0A-ROUTE1`, `E0-P0A-FD7` | Preserve other provider entries. |
| `overview.md` | `internal/aicontext/skills/payment-integration/references/implementation-workflows.md` | workflow route | P1-A | edit only if adding `payment-gateway.md` | `E0-P0A-ROUTE1`, `E0-P0A-FD8` | Update SePay workflow only. |
| `api.md` | `webhooks.md` | API v2 reconciliation after webhook misses | P2-A, P3-A | coordinate/edit in own phase | `E0-P0A-SRC2`, `E0-P0A-SRC3` | Do not duplicate webhook auth details in API doc. |
| `api.md` | `qr-codes.md` | Order VA and QR account/memo rules intersect | P2-A, P5-A | coordinate/edit in own phase | `E0-P0A-SRC2`, `E0-P0A-SRC5` | Keep bank-specific QR rules in QR doc, not only API doc. |
| `webhooks.md` | `internal/aicontext/skills/payment-integration/scripts/sepay-webhook-verify.js` | helper currently weaker than target webhook guidance | P3-A | inspect-only / follow-up | `E0-P0A-SRC7`, `E0-P0A-FD9` | Do not edit script in this plan without user scope expansion. |
| `webhooks.md` | `internal/aicontext/skills/payment-integration/scripts/test-scripts.js` | related test file for verifier script | P3-A | do-not-touch unless script scope expands | `E0-P0A-FD9` | No script change means no test update. |
| `sdk.md` | `payment-gateway.md` | gateway SDK depends on hosted checkout/IPN contract | P4-A | edit | `E0-P0A-SRC4`, `E0-P0A-REPORT1` | Keep SDK examples separate from webhook reconciliation. |
| `best-practices.md` | all SePay reference docs | cross-surface production summary | P5-A | edit after P2-P4 content exists | `E0-P0A-SRC1..E0-P0A-SRC6` | Best practices must not introduce facts missing from detailed docs. |
| generated `.agents/.claude` payment-integration output | internal source skill files | generated output | P6-A | validate-only / regenerate if commanded by repo tooling | `E0-P0A-GRAPH1` | Never edit generated output as source. |

## Detailed Findings

### Overview

Current state:

`overview.md` contains useful high-level SePay context, but its environment and rate-limit guidance are stale and it does not clearly route agents between API v2, Webhooks, Payment Gateway/IPN, QR utility, and Order VA.

Required state:

```text
Overview must be a surface-selection map. It should prevent agents from starting with a stale endpoint or the wrong SePay product surface.
```

Evidence:

- `E0-P0A-SRC1`: local source read.
- `E0-P0A-FD1`: file-detail evidence.
- `E0-P0A-REPORT1`: report findings 1, 2, 5, 6, 7, 13, 14.

Relationship and impact:

- Related file count: 0
- Relationship summary: no graph relationships, but plan-relevant routing files affect discoverability.
- Impact note: low graph risk, high product-guidance risk.

Classification:

wrong

Allowed next action:

Edit P1-A.

Forbidden next action:

Do not bury Payment Gateway/IPN under webhook wording.

### API

Current state:

`api.md` teaches legacy v1 `userapi/*` as default and has stale rate-limit behavior.

Required state:

```text
API v2 is default for new work. v1 is legacy/migration-only. Order VA is bank-specific, not one generic flow.
```

Evidence:

- `E0-P0A-SRC2`
- `E0-P0A-FD2`
- `E0-P0A-REPORT1`

Relationship and impact:

- Related file count: 0
- Relationship summary: no graph relationships; content must coordinate with webhooks and QR docs.
- Impact note: low graph risk, high correctness risk.

Classification:

wrong

Allowed next action:

Edit P2-A.

Forbidden next action:

Do not retain legacy v1 as the first/default API path.

### Webhooks

Current state:

`webhooks.md` has basic webhook setup and idempotency ideas but lacks current HMAC, payment-code, retry, replay, monitoring, incident, and allowlist details. It also needs stronger IPN separation.

Required state:

```text
Bank-account webhooks must be documented as their own surface with current auth, response, retry, payment-code, monitoring, replay, and reconciliation behavior.
```

Evidence:

- `E0-P0A-SRC3`
- `E0-P0A-FD3`
- `E0-P0A-REPORT1`

Relationship and impact:

- Related file count: 0
- Relationship summary: script verifier is plan-relevant but inspect-only.
- Impact note: low graph risk for docs; high money-handling guidance risk.

Classification:

wrong

Allowed next action:

Edit P3-A.

Forbidden next action:

Do not edit `scripts/sepay-webhook-verify.js` unless the user expands scope.

### SDK And Payment Gateway

Current state:

`sdk.md` includes stale install/endpoint details and does not fully describe current gateway/IPN/order lifecycle behavior. A dedicated gateway reference is missing.

Required state:

```text
Hosted checkout/IPN should be discoverable as a separate Payment Gateway surface, with SDK details kept current and connected to that surface.
```

Evidence:

- `E0-P0A-SRC4`
- `E0-P0A-FD4`
- `E0-P0A-ROUTE1`
- `E0-P0A-REPORT1`

Relationship and impact:

- Related file count: 0 for `sdk.md`; new `payment-gateway.md` missing.
- Relationship summary: route files need minimal updates if new file is added.
- Impact note: low graph risk, high product-surface risk.

Classification:

`sdk.md`: wrong

`payment-gateway.md`: missing

Allowed next action:

Create/fill P1-A/P4-A and update routing minimally.

Forbidden next action:

Do not leave a new gateway doc undiscoverable from the skill entrypoint.

### QR And Best Practices

Current state:

`qr-codes.md` is partial and `best-practices.md` is wrong for a reusable skill because it teaches one application's implementation as general SePay practice.

Required state:

```text
QR docs must cover current parameters and bank rules. Best practices must be repo-agnostic and summarize correct patterns from the refreshed detailed docs.
```

Evidence:

- `E0-P0A-SRC5`
- `E0-P0A-SRC6`
- `E0-P0A-FD5`
- `E0-P0A-FD6`
- `E0-P0A-REPORT1`

Relationship and impact:

- Related file count: 0 for both docs.
- Relationship summary: best practices depends on the detailed doc phases being accurate.
- Impact note: low graph risk, high guidance consistency risk.

Classification:

`qr-codes.md`: partial

`best-practices.md`: wrong

Allowed next action:

Edit P5-A after P2-A through P4-A establish current detailed facts.

Forbidden next action:

Do not preserve project-specific examples as if they are general SePay defaults.

## Next Phase Status Decisions

| Plan Item | Actual Status Finding | Required Status / Next-Action Update |
|-----------|-----------------------|--------------------------------------|
| P1-A | `overview.md` is wrong and gateway/IPN reference is missing. | keep P1-A goal; create surface map and gateway route before detailed phases. |
| P2-A | `api.md` is wrong and v1-default. | keep P2-A goal; rewrite around API v2 default and legacy v1 section. |
| P3-A | `webhooks.md` is wrong and script verifier is partial/high-risk but out of current scope. | keep P3-A goal; update docs only and record script as follow-up unless scope expands. |
| P4-A | `sdk.md` is wrong and gateway/IPN doc is missing. | keep P4-A goal; fill dedicated gateway/IPN doc and refresh SDK. |
| P5-A | QR is partial and best practices are wrong/project-specific. | keep P5-A goal; update after earlier phases to avoid stale summary. |
| P6-A | Source skill changes may regenerate agent skill output. | keep P6-A goal; validate source, generated output, full build, and detect-changes. |

## Implementation Gate

- [x] Target scope is listed in Current Status Matrix.
- [x] Each target unit has a status.
- [x] Each status has evidence IDs.
- [x] Each target file has relationship count evidence from `file-detail` when applicable.
- [x] Phase Touch Map lists plan-relevant relationship files that can affect the current phase/slice.
- [x] Phase Touch Map defines touch mode for every plan-relevant relationship unit that may be affected.
- [x] Correct parts are marked preserve-only where applicable.
- [x] Partial, missing, wrong, unbound, and fake-or-stub parts have exact next actions.
- [x] Blockers are recorded, if any.
- [x] Next phase status assumptions, next action, and work steps have been updated from this status file when needed.
- [x] Status Refresh Log has an R0 baseline row.
- [x] Implementation has not started; post-implementation refresh checks are not yet applicable.
- [x] If implementation later changes status, affected Current Status Matrix rows must be refreshed from latest evidence.
- [x] If refreshed statuses later change next work, only the stale next-phase status assumptions, next action, or work steps must be updated before the next phase.

## Final P0 Decision

Choose one:

- [ ] P0 actual-status incomplete. Implementation is blocked.
- [x] P0 complete. Next phase can proceed unchanged.
- [ ] P0 complete. Next phase status, next action, or work steps must be updated before implementation.
- [ ] P0 complete. Target scope is preserve-only.
- [ ] P0 complete. Implementation is blocked by missing authority or evidence.

Decision note:

Implementation can proceed with P1-A because the plan work steps already reflect the P0 decisions: add a dedicated gateway/IPN reference and route it, keep the webhook verifier script inspect-only unless scope expands, and refresh this actual-status file after each completed slice before continuing.
