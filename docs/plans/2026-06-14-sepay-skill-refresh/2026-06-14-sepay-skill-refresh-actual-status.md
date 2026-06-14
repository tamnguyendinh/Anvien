# Actual Status

Title: SePay Skill Refresh
Date: 2026-06-14
Status: Complete; final commit executed in this turn
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
| `overview.md` | Surface map now separates API v2, bank-account Webhooks, Payment Gateway/IPN, QR utility, Order VAs, Test Mode, and production safety rules. | Surface router with current API v2, Webhook, Payment Gateway/IPN, QR, Order VA, Test Mode, and official-doc link guidance. | correct | 0 related files | `E1-P1A-IMPL1`, `E1-P1A-SRC1` | preserve; validate P6-A |
| `api.md` | API v2 is the default; legacy v1 appears only as migration guidance; Order VA bank rules are represented. | API v2 default, v1 legacy/migration, current rate limit, response envelope, pagination, and Order VA bank rules. | correct | 0 related files | `E2-P2A-IMPL1`, `E2-P2A-SRC1`, `E2-P2A-SRC2`, `E2-P2A-SRC3` | preserve; validate P6-A |
| `webhooks.md` | Bank-account webhook doc now includes HMAC, raw body, retry correction, payment code, monitoring, replay, IP allowlist, reconciliation, and explicit IPN separation. | Current bank-account webhook guidance with HMAC, raw body, retry correction, payment code, monitoring, replay, IP allowlist, and explicit IPN separation. | correct | 0 related files | `E3-P3A-IMPL1`, `E3-P3A-SRC1`, `E3-P3A-SRC2`, `E3-P3A-SRC3` | preserve; validate P6-A |
| `sdk.md` | SDK doc now uses official Node/PHP install paths, current gateway lifecycle topics, payment methods, field-order warning, and Laravel package caution. | Current Node/PHP SDK guidance, gateway base URLs/auth, payment methods, form field order, order detail/status/cancel/void, and accurate Laravel classification. | correct | 0 related files | `E4-P4A-IMPL2`, `E4-P4A-SRC2`, `E4-P4A-SRC3` | preserve; validate P6-A |
| `qr-codes.md` | QR doc now includes current parameters and bank-specific VA/memo rules, including `TKP` and `SEVQR`. | Current QR utility guidance with missing parameters, bank identifiers, official VA, memo-based VA, no VA, and VietinBank `SEVQR` examples. | correct | 0 related files | `E5-P5A-IMPL1`, `E5-P5A-SRC1` | preserve; validate P6-A |
| `best-practices.md` | Best practices are repo-agnostic and summarize surface selection, security, idempotency, reconciliation, Test Mode/Live, allowlisting, and amount policy. | Repo-agnostic SePay production patterns for surface choice, HMAC, payment code, idempotency, replay, reconciliation, Test Mode/Live, IP allowlisting, and amount policy. | correct | 0 related files | `E5-P5A-IMPL2`, `E5-P5A-SRC2`, `E5-P5A-SRC3` | preserve; validate P6-A |
| `payment-gateway.md` | Dedicated Payment Gateway/IPN reference exists and covers hosted checkout, gateway API, IPN, order APIs, idempotency, and reconciliation. | Dedicated Payment Gateway/IPN reference under `references/sepay`. | correct | 0 related files | `E1-P1A-IMPL2`, `E4-P4A-IMPL1`, `E4-P4A-SRC1`, `E6-P6A-FD1` | preserve |
| `SKILL.md` SePay quick reference | Lists seven SePay reference files, including dedicated Payment Gateway/IPN reference. | If `payment-gateway.md` is added, list it so agents can discover hosted checkout/IPN guidance. | correct | 0 related files | `E1-P1A-ROUTE1`, `E1-P1A-SRC2` | preserve; validate P6-A |
| `implementation-workflows.md` SePay workflow | Routes SePay agents to surface selection first, then API v2, Webhooks, Payment Gateway/IPN, SDK, QR, helper caveat, and best practices. | If `payment-gateway.md` is added, route hosted checkout/IPN work to it before SDK/webhook details. | correct | 0 related files | `E1-P1A-ROUTE2`, `E1-P1A-SRC2` | preserve; validate P6-A |
| `scripts/sepay-webhook-verify.js` | Helper supports API Key/OAuth2/none and payload validation, but no HMAC raw-body signature path. | Out of current implementation scope; should remain inspect-only/follow-up unless user expands scope. | partial | 1 inbound related file plus local symbol relationships | `E0-P0A-SRC7`, `E0-P0A-FD9` | inspect-only / do not edit |

## Status Refresh Log

| Refresh | Date | Repo Basis | Changed Scope | Status Changes | Evidence | Next Phase Update |
|---------|------|------------|----------------|----------------|----------|-------------------|
| R0 | 2026-06-14 | baseline before implementation | SePay reference files and direct routing/helper surfaces | initial classification | `E0-P0A-GRAPH1`, `E0-P0A-REPORT1`, `E0-P0A-SRC1..E0-P0A-SRC7`, `E0-P0A-FD1..E0-P0A-FD9` | P1-A creates surface map and gateway route; P2-A through P5-A keep goals but use latest status rows. |
| R1 | 2026-06-14 | source-doc implementation before final validation | SePay reference docs and direct routing files | `overview.md`, `api.md`, `webhooks.md`, `sdk.md`, `qr-codes.md`, `best-practices.md`, `payment-gateway.md`, `SKILL.md`, and `implementation-workflows.md` moved to correct; script remains inspect-only | `E1-P1A-*`, `E2-P2A-*`, `E3-P3A-*`, `E4-P4A-*`, `E5-P5A-*` | P6-A validation can proceed; no P1-P5 source-doc work remains. |
| R2 | 2026-06-14 | post-build validation and graph refresh | P6 validation | full build passed; graph refreshed; new gateway file has 0 related files and low risk; repo-generated `.agents`/`.claude` payment-integration outputs smoke-pass; home-level installed skill not regenerated; `checkout-helper.js` stale URLs recorded as out-of-scope follow-up | `E6-P6A-VAL1`, `E6-P6A-VAL2`, `E6-P6A-GRAPH1`, `E6-P6A-FD1`, `E6-P6A-GEN1`, `E6-P6A-GEN2`, `E6-P6A-GEN3`, `E6-P6A-DETECT1`, `E6-P6A-SCOPE1` | Pn-A supervisor review can proceed. |
| R3 | 2026-06-14 | supervisor and dead-work cleanup review | Pn-A/Pn-B closure | supervisor PASS recorded; no plan-created dead work found; stale helper script remains documented out-of-scope follow-up | `E7-PNA-SUP1`, `E8-PNB-CLEAN1`, `E8-PNB-SUP1` | Pn-C final validation/detect/commit can proceed. |
| R4 | 2026-06-14 | final graph refresh and detect-changes | Pn-C closure | final analyze and detect-changes passed; final commit executed after evidence recording | `E9-PNC-GRAPH1`, `E9-PNC-DETECT1`, `E9-PNC-COMMIT1` | plan complete. |

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
| `internal/aicontext/skills/payment-integration/scripts/checkout-helper.js` | gateway docs now supersede old helper endpoint guidance | script helper outside docs-only scope | follow-up | do-not-touch in this plan | `E6-P6A-SCOPE1` | Stale gateway URLs found during validation; requires a separate code/script plan if user expands scope. |

## Detailed Findings

### Overview

Current state:

`overview.md` now acts as a surface-selection map. It separates API v2, bank-account Webhooks, Payment Gateway/IPN, QR utility, Order VAs, optional SePay products, Test Mode, and production safety guidance.

Required state:

```text
Overview must be a surface-selection map. It should prevent agents from starting with a stale endpoint or the wrong SePay product surface.
```

Evidence:

- `E1-P1A-IMPL1`
- `E1-P1A-SRC1`
- `E1-P1A-OFFICIAL1`

Relationship and impact:

- Related file count: 0
- Relationship summary: no graph relationships, but plan-relevant routing files affect discoverability.
- Impact note: low graph risk, high product-guidance risk.

Classification:

correct

Allowed next action:

Preserve and validate in P6-A.

Forbidden next action:

Do not bury Payment Gateway/IPN under webhook wording.

### API

Current state:

`api.md` now teaches API v2 as the default, keeps legacy v1 only under migration notes, documents API v2 rate limits/response shape/pagination, and adds bank-specific Order VA guidance.

Required state:

```text
API v2 is default for new work. v1 is legacy/migration-only. Order VA is bank-specific, not one generic flow.
```

Evidence:

- `E2-P2A-IMPL1`
- `E2-P2A-SRC1`
- `E2-P2A-SRC2`
- `E2-P2A-SRC3`

Relationship and impact:

- Related file count: 0
- Relationship summary: no graph relationships; content must coordinate with webhooks and QR docs.
- Impact note: low graph risk, high correctness risk.

Classification:

correct

Allowed next action:

Preserve and validate in P6-A.

Forbidden next action:

Do not retain legacy v1 as the first/default API path.

### Webhooks

Current state:

`webhooks.md` now documents bank-account Webhooks with HMAC-SHA256, raw-body signature verification, payment-code recognition, retry/replay behavior, monitoring, incidents, reconciliation, IP allowlisting, and explicit Payment Gateway IPN separation.

Required state:

```text
Bank-account webhooks must be documented as their own surface with current auth, response, retry, payment-code, monitoring, replay, and reconciliation behavior.
```

Evidence:

- `E3-P3A-IMPL1`
- `E3-P3A-SRC1`
- `E3-P3A-SRC2`
- `E3-P3A-SRC3`
- `E3-P3A-SCOPE1`

Relationship and impact:

- Related file count: 0
- Relationship summary: script verifier is plan-relevant but inspect-only.
- Impact note: low graph risk for docs; high money-handling guidance risk.

Classification:

correct

Allowed next action:

Preserve and validate in P6-A.

Forbidden next action:

Do not edit `scripts/sepay-webhook-verify.js` unless the user expands scope.

### SDK And Payment Gateway

Current state:

`sdk.md` now documents official Node/PHP Payment Gateway SDK use, payment methods, field-order signature warning, order lifecycle methods, and Laravel package caution. `payment-gateway.md` now exists and covers hosted checkout, gateway API, IPN, idempotency, and reconciliation.

Required state:

```text
Hosted checkout/IPN should be discoverable as a separate Payment Gateway surface, with SDK details kept current and connected to that surface.
```

Evidence:

- `E1-P1A-IMPL2`
- `E4-P4A-IMPL1`
- `E4-P4A-IMPL2`
- `E4-P4A-SRC1`
- `E4-P4A-SRC2`
- `E4-P4A-SRC3`

Relationship and impact:

- Related file count: 0 for `sdk.md`; new `payment-gateway.md` will receive graph relationship evidence after P6-A analyze refresh.
- Relationship summary: route files now mention the new gateway reference.
- Impact note: low graph risk, high product-surface risk.

Classification:

`sdk.md`: correct

`payment-gateway.md`: correct

Allowed next action:

Preserve and validate in P6-A.

Forbidden next action:

Do not leave a new gateway doc undiscoverable from the skill entrypoint.

### QR And Best Practices

Current state:

`qr-codes.md` now covers current QR parameters and bank-specific VA/memo rules. `best-practices.md` is repo-agnostic and no longer teaches one application model as general SePay practice.

Required state:

```text
QR docs must cover current parameters and bank rules. Best practices must be repo-agnostic and summarize correct patterns from the refreshed detailed docs.
```

Evidence:

- `E5-P5A-IMPL1`
- `E5-P5A-IMPL2`
- `E5-P5A-SRC1`
- `E5-P5A-SRC2`
- `E5-P5A-SRC3`

Relationship and impact:

- Related file count: 0 for both docs.
- Relationship summary: best practices depends on the detailed doc phases being accurate.
- Impact note: low graph risk, high guidance consistency risk.

Classification:

`qr-codes.md`: correct

`best-practices.md`: correct

Allowed next action:

Preserve and validate in P6-A.

Forbidden next action:

Do not preserve project-specific examples as if they are general SePay defaults.

## Next Phase Status Decisions

| Plan Item | Actual Status Finding | Required Status / Next-Action Update |
|-----------|-----------------------|--------------------------------------|
| P1-A | Surface map and gateway routing are complete. | preserve; P6-A readback and generated-output smoke must confirm discoverability. |
| P2-A | API v2 guidance is default and v1 is legacy-only. | preserve; P6-A validation must confirm no stale v1 default remains. |
| P3-A | Bank-account webhook guidance is current and IPN is separated; verifier script remains out of scope. | preserve; record script/helper stale endpoint observations only as follow-up unless scope expands. |
| P4-A | Dedicated Payment Gateway/IPN guidance and SDK guidance are complete. | preserve; P6-A validation must include the new file and routing surfaces. |
| P5-A | QR and best practices are refreshed and repo-agnostic. | preserve; P6-A validation must confirm project-specific SePay reference examples are gone. |
| P6-A | Source skill docs changed and plan files were refreshed; validation completed. | preserve evidence; proceed to supervisor/dead-work/closure steps. |

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
