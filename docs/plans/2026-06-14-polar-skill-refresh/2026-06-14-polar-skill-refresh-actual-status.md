# Actual Status

Title: Polar Skill Refresh
Date: 2026-06-14
Status: Closure evidence recorded; final commit pending in same turn
Companion plan: `docs/plans/2026-06-14-polar-skill-refresh/2026-06-14-polar-skill-refresh-plan.md`
Companion evidence: `docs/plans/2026-06-14-polar-skill-refresh/2026-06-14-polar-skill-refresh-evidence.md`
Companion benchmark: `docs/plans/2026-06-14-polar-skill-refresh/2026-06-14-polar-skill-refresh-benchmark.md`

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

- `internal/aicontext/skills/payment-integration/references/polar/overview.md`
- `internal/aicontext/skills/payment-integration/references/polar/products.md`
- `internal/aicontext/skills/payment-integration/references/polar/checkouts.md`
- `internal/aicontext/skills/payment-integration/references/polar/subscriptions.md`
- `internal/aicontext/skills/payment-integration/references/polar/webhooks.md`
- `internal/aicontext/skills/payment-integration/references/polar/benefits.md`
- `internal/aicontext/skills/payment-integration/references/polar/sdk.md`
- `internal/aicontext/skills/payment-integration/references/polar/best-practices.md`
- `internal/aicontext/skills/payment-integration/references/polar/usage-based-billing.md` if added by P1-A
- `internal/aicontext/skills/payment-integration/references/polar/customer-portal.md` if added by P1-A
- `internal/aicontext/skills/payment-integration/references/polar/customer-state.md` if added by P1-A
- `internal/aicontext/skills/payment-integration/references/polar/orders-refunds-discounts.md` if added by P1-A
- minimal routing updates in `internal/aicontext/skills/payment-integration/SKILL.md` and `internal/aicontext/skills/payment-integration/references/implementation-workflows.md`
- source skill README cleanup in `internal/aicontext/skills/payment-integration/README.md` when stale Polar examples are found during readback

Out of scope:

- `internal/aicontext/skills/payment-integration/scripts/polar-webhook-verify.js` implementation changes unless the user explicitly expands scope.
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
| `internal/aicontext/skills/payment-integration/references/polar/overview.md` | `E0-P0A-FD1` | 0 related files | Markdown docs; no local, inbound, outbound, linked flow, or linked test relationships. | low scope warning |
| `internal/aicontext/skills/payment-integration/references/polar/checkouts.md` | `E0-P0A-FD2` | 0 related files | Markdown docs; no local, inbound, outbound, linked flow, or linked test relationships. | low scope warning |
| `internal/aicontext/skills/payment-integration/references/polar/webhooks.md` | `E0-P0A-FD3` | 0 related files | Markdown docs; no local, inbound, outbound, linked flow, or linked test relationships. | low scope warning |
| `internal/aicontext/skills/payment-integration/references/polar/sdk.md` | `E0-P0A-FD4` | 0 related files | Markdown docs; no local, inbound, outbound, linked flow, or linked test relationships. | low scope warning |
| `internal/aicontext/skills/payment-integration/references/polar/products.md` | `E0-P0A-FD5` | 0 related files | Markdown docs; no local, inbound, outbound, linked flow, or linked test relationships. | low scope warning |
| `internal/aicontext/skills/payment-integration/references/polar/subscriptions.md` | `E0-P0A-FD6` | 0 related files | Markdown docs; no local, inbound, outbound, linked flow, or linked test relationships. | low scope warning |
| `internal/aicontext/skills/payment-integration/references/polar/benefits.md` | `E0-P0A-FD7` | 0 related files | Markdown docs; no local, inbound, outbound, linked flow, or linked test relationships. | low scope warning |
| `internal/aicontext/skills/payment-integration/references/polar/best-practices.md` | `E0-P0A-FD8` | 0 related files | Markdown docs; no local, inbound, outbound, linked flow, or linked test relationships. | low scope warning |
| `internal/aicontext/skills/payment-integration/SKILL.md` | `E0-P0A-FD9` | 0 related files | Markdown docs; route surface for Polar reference discoverability. | low scope warning |
| `internal/aicontext/skills/payment-integration/references/implementation-workflows.md` | `E0-P0A-FD10` | 0 related files | Markdown docs; workflow route surface for Polar implementation guidance. | low scope warning |
| `reports/problem/2026-06-14-polar-skill-refresh-analysis.md` | `E0-P0A-FD11` | 0 related files | Markdown report used as planning authority. | inspect-only |
| `internal/aicontext/skills/payment-integration/scripts/polar-webhook-verify.js` | `E0-P0A-FD12` | 4 inbound refs plus 16 local relationships | JavaScript source; related test file `scripts/test-scripts.js`; 28 symbols and 91 unresolved entries. | high scope warning; inspect-only/follow-up |

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
| `overview.md` | Surface routing, auth, sandbox, rate limits, fees, Customer State, and next-reference map refreshed. | Current surface map for products, checkout, webhooks, subscriptions, benefits, usage billing, customer portal, SDK/adapters, fees, sandbox, and auth. | correct | 0 related files; post-change graph/file-detail refreshed in P7 | `E1-P1A-IMPL1`, `E6-P6A-VAL2`, `E7-P7A-FD2` | preserve; final staged detect/commit |
| `checkouts.md` | Checkout Sessions, Checkout Links, embedded checkout/payment method, localization, customer fields, discounts, ad-hoc prices, and `order.paid` verification refreshed; legacy price field only appears as a warning. | Current checkout sessions using `products: [productId]`, multiple products, ad-hoc prices, success URL, external customer ID, embedded checkout, and customer IP guidance. | correct | 0 related files; post-change graph/file-detail refreshed in P7 | `E2-P2A-IMPL1`, `E2-P2A-IMPL2`, `E6-P6A-VAL1`, `E6-P6A-VAL2`, `E7-P7A-FD2` | preserve; final staged detect/commit |
| `products.md` | Product/pricing model refreshed for intervals, interval counts, pricing types, metered prices, seats, currencies, tax behavior, media, benefits, and custom fields. | Current product model with billing intervals/counts, pricing types, metered prices, seats, currencies, tax behavior, media, visibility, and custom fields. | correct | 0 related files; post-change graph/file-detail refreshed in P7 | `E2-P2A-IMPL3`, `E2-P2A-IMPL4`, `E7-P7A-FD2` | preserve; final staged detect/commit |
| `webhooks.md` | Standard Webhooks, raw body, SDK/adapters, timeout/retry/disable behavior, event taxonomy, sequences, idempotency, and Customer State webhook coverage refreshed. | Current webhook verification, delivery, retry, endpoint-disable, event taxonomy, cancellation, renewal, and SDK/adapters guidance. | correct | 0 related files; post-change graph/file-detail refreshed in P7 | `E3-P3A-IMPL1`, `E3-P3A-IMPL2`, `E3-P3A-IMPL3`, `E7-P7A-FD2` | preserve; final staged detect/commit |
| `subscriptions.md` | Subscription fields, statuses, proration, discount/trial fields, cancel/revoke/uncancel, renewal/failure, portal, Customer State, and order/refund cross-links refreshed. | Current subscription management using `product_id`, `proration_behavior`, `trial_end`, cancel/revoke/uncancel, and Customer Portal boundaries. | correct | 0 related files; post-change graph/file-detail refreshed in P7 | `E4-P4A-IMPL1`, `E7-P7A-FD2` | preserve; final staged detect/commit |
| `customer-portal.md` | Dedicated Customer Portal/customer session reference created and filled. | Dedicated Customer Portal/customer sessions reference. | correct | new file; 0 related files after post-change file-detail | `E1-P1A-IMPL2`, `E4-P4A-IMPL2`, `E4-P4A-IMPL3`, `E7-P7A-FD1` | preserve; final staged detect/commit |
| `usage-based-billing.md` | Dedicated usage billing reference created and filled. | Dedicated usage billing, event ingestion, meters, credits, customer meters, and cost insight reference. | correct | new file; 0 related files after post-change file-detail | `E1-P1A-IMPL2`, `E5-P5A-IMPL1`, `E7-P7A-FD1` | preserve; final staged detect/commit |
| `customer-state.md` | Dedicated Customer State reference created and filled. | Dedicated Customer State entitlement-sync reference covering one-call state, external ID lookup, active subscriptions, granted benefits, active meters/current balance, and `customer.state_changed`. | correct | new file; 0 related files after post-change file-detail | `E1-P1A-IMPL2`, `E5-P5A-IMPL2`, `E7-P7A-FD1` | preserve; final staged detect/commit |
| `orders-refunds-discounts.md` | Dedicated commercial-operations reference created and filled. | Dedicated commercial-operations reference covering order statuses, `billing_reason`, invoices/receipts, `order.paid`, refunds, benefit revocation, subscription refund vs cancellation, and discount modes/restrictions. | correct | new file; 0 related files after post-change file-detail | `E1-P1A-IMPL2`, `E6-P6A-IMPL1`, `E7-P7A-FD1` | preserve; final staged detect/commit |
| `benefits.md` | Benefit reference refreshed for Credits, License Keys, Feature Flags, File Downloads, GitHub, Discord, Custom, grants, Customer State, and `benefit_id` scoping. | Current benefits reference including Credits, License Keys, Feature Flags, File Downloads, GitHub, Discord, Custom, grants, and Customer State. | correct | 0 related files; post-change graph/file-detail refreshed in P7 | `E5-P5A-IMPL3`, `E7-P7A-FD2` | preserve; final staged detect/commit |
| `sdk.md` | SDK/adapters refreshed for TypeScript SDK, camelCase/snake_case, Next.js, Express, BetterAuth, Laravel, checkout embed, portal, usage, products, and webhooks. | Current TypeScript SDK, language SDK, Next.js, Express, Laravel, BetterAuth, checkout/webhook/portal/usage adapter guidance. | correct | 0 related files; post-change graph/file-detail refreshed in P7 | `E6-P6A-IMPL2`, `E7-P7A-FD2` | preserve; final staged detect/commit |
| `best-practices.md` | Production patterns refreshed for fees, rate limits, envs, sandbox limitations, pagination, webhooks, idempotency, checkout, portal, usage, benefits, orders/refunds/discounts, and MoR launch checks. | Repo-agnostic current best practices for fees, rate limits, envs, webhooks, idempotency, checkout, portal, usage, benefits, and retries. | correct | 0 related files; post-change graph/file-detail refreshed in P7 | `E6-P6A-IMPL3`, `E6-P6A-VAL1`, `E6-P6A-VAL2`, `E7-P7A-FD2` | preserve; final staged detect/commit |
| `SKILL.md` Polar quick reference | Lists all 12 Polar references and removes stale rate highlight. | Route all current Polar refs and remove stale rate highlight. | correct | 0 related files; post-change graph/file-detail refreshed in P7 | `E1-P1A-ROUTE1`, `E6-P6A-VAL2`, `E7-P7A-FD3` | preserve; final staged detect/commit |
| `implementation-workflows.md` Polar workflow | Loads overview first, routes all new references, and prefers SDK/adapters for webhook verification. | Load overview first, then specific current refs; prefer SDK/adapters for webhooks; include usage and customer portal refs. | correct | 0 related files; post-change graph/file-detail refreshed in P7 | `E1-P1A-ROUTE2`, `E3-P3A-IMPL3`, `E6-P6A-VAL2`, `E7-P7A-FD3` | preserve; final staged detect/commit |
| `README.md` Polar overview | Source skill README lists new Polar references and no longer promotes the legacy Polar checkout helper as the production path. | Skill overview should not teach stale Polar checkout fields and should expose new reference files. | correct | 0 related files after post-change file-detail | `E6-P6A-IMPL4`, `E6-P6A-VAL1`, `E6-P6A-VAL2`, `E7-P7A-FD3` | preserve; final staged detect/commit |
| `scripts/polar-webhook-verify.js` | Custom verifier exists and has inbound tests; official docs prefer SDK/adapters. | Inspect-only for this content plan; update only if user expands scope and impact/test work is planned. | partial | 4 inbound refs plus 16 local relationships | `E0-P0A-SCRIPT1`, `E0-P0A-FD12` | inspect-only / do not edit |

## Status Refresh Log

| Refresh | Date | Repo Basis | Changed Scope | Status Changes | Evidence | Next Phase Update |
|---------|------|------------|----------------|----------------|----------|-------------------|
| R0 | 2026-06-14 | baseline before implementation | Polar reference files, route files, problem report, and webhook helper script | initial classification | `E0-P0A-GRAPH1`, `E0-P0A-GRAPH2`, `E0-P0A-REPORT1`, `E0-P0A-SRC1..E0-P0A-SRC8`, `E0-P0A-ROUTE1..E0-P0A-ROUTE2`, `E0-P0A-SCRIPT1`, `E0-P0A-FD1..E0-P0A-FD12` | P1-A creates routing and new refs; P2-A through P6-A keep goals but use latest status rows. |
| R1 | 2026-06-14 | plan review against additional official practical Polar docs | plan file, actual-status, evidence, benchmark | added missing planned refs `customer-state.md` and `orders-refunds-discounts.md`; expanded checkout, portal, usage, SDK, and best-practice phase requirements | `E0-P0A-DOCS1..E0-P0A-DOCS6`, `E0-P0A-IMPL1` | P1-A must add four new refs, not two; P2-A/P4-A/P5-A/P6-A must include practical guide details found in official docs. |
| R2 | 2026-06-14 | source-doc implementation before P7 validation | Polar references, route files, and skill README | `partial -> correct` for existing Polar docs and routing; `missing -> correct` for four new references; README stale checkout-helper example corrected; helper script remains inspect-only | `E1-P1A-IMPL1..E1-P1A-VAL1`, `E2-P2A-IMPL1..E2-P2A-IMPL4`, `E3-P3A-IMPL1..E3-P3A-IMPL3`, `E4-P4A-IMPL1..E4-P4A-IMPL3`, `E5-P5A-IMPL1..E5-P5A-IMPL3`, `E6-P6A-IMPL1..E6-P6A-VAL2` | P7-A can proceed with readback, diff check, full build, graph refresh, generated-output smoke, and detect-changes. |
| R3 | 2026-06-14 | docs-only validation and supervisor review | P7-A and Pn-A | graph refreshed; file-detail stayed low-risk docs-only; generated repo skill output smoke passed with helper-script caveat; supervisor PASS; full build not used because user narrowed scope to docs-only validation | `E7-P7A-VAL1..E7-P7A-DETECT1`, `E8-PNA-SUP1` | Run dead-work cleanup, final staged detect-changes, commit, and close. |
| R4 | 2026-06-14 | dead-work cleanup and final staged detect | Pn-B and Pn-C | cleanup passed; final staged diff check passed; final staged detect-changes passed with docs/reporting only and low risk; final commit executed after evidence recording | `E9-PNB-CLEAN1`, `E10-PNC-VAL1`, `E10-PNC-DETECT1`, `E10-PNC-COMMIT1` | plan complete after commit. |

## Phase Touch Map

Use this map to prevent accidental edits. A related file is not automatically editable.

`Plan-Relevant Relationship File` lists only a relationship file that can directly affect or be affected by the planned phase or slice. Do not copy a full relationship inventory into this map. Include only files whose relationship can affect the phase/slice decision, touch mode, or validation.

| Unit / File / Surface | Plan-Relevant Relationship File | Relationship to Target | Plan Item | Touch Mode | Evidence | Constraint |
|-----------------------|---------------------------------|------------------------|-----------|------------|----------|------------|
| `overview.md` | `reports/problem/2026-06-14-polar-skill-refresh-analysis.md` | authority report | P1-A | inspect-only | `E0-P0A-REPORT1` | Use report findings; re-check official docs if delayed. |
| `overview.md` | `internal/aicontext/skills/payment-integration/SKILL.md` | routing/discoverability | P1-A | edit | `E0-P0A-ROUTE1`, `E0-P0A-FD9` | Preserve other provider entries. |
| `overview.md` | `internal/aicontext/skills/payment-integration/references/implementation-workflows.md` | workflow route | P1-A | edit | `E0-P0A-ROUTE2`, `E0-P0A-FD10` | Update Polar workflow only. |
| `checkouts.md` | `products.md` | product IDs, pricing, ad-hoc prices | P2-A | coordinate/edit in same phase | `E0-P0A-SRC2`, `E0-P0A-SRC5` | Do not leave conflicting product-price examples. |
| `checkouts.md` | `usage-based-billing.md` | metered checkout/ad-hoc price cross-link | P2-A/P5-A | edit when file exists | `E0-P0A-REPORT1` | Avoid duplicating usage billing details in checkout doc. |
| `webhooks.md` | `scripts/polar-webhook-verify.js` | helper currently exposed by workflow | P3-A | inspect-only / follow-up | `E0-P0A-SCRIPT1`, `E0-P0A-FD12` | Do not edit script without expanded scope and impact/test plan. |
| `webhooks.md` | `internal/aicontext/skills/payment-integration/scripts/test-scripts.js` | related test file for verifier script | P3-A | do-not-touch unless script scope expands | `E0-P0A-FD12` | No script change means no test update. |
| `subscriptions.md` | `customer-portal.md` | customer self-service boundaries | P4-A | coordinate/edit in same phase | `E0-P0A-SRC6`, `E0-P0A-REPORT1` | Keep Core API vs Customer Portal API separated. |
| `subscriptions.md` | `customer-state.md` | compact entitlement and subscription-state source | P4-A, P5-A | edit when file exists | `E0-P0A-DOCS4` | Use Customer State for access decisions when full subscription/benefit/meter state is needed. |
| `benefits.md` | `usage-based-billing.md` | Credits benefit and customer meter balance | P5-A | coordinate/edit in same phase | `E0-P0A-SRC7`, `E0-P0A-REPORT1` | Avoid treating Credits as generic custom benefit. |
| `webhooks.md` | `customer-state.md` | `customer.state_changed` entitlement webhook | P3-A, P5-A | edit when file exists | `E0-P0A-DOCS4` | Document as entitlement sync shortcut, not as replacement for all event-specific workflows. |
| `checkouts.md` | `orders-refunds-discounts.md` | checkout success produces orders, discounts, refund consequences | P2-A, P6-A | edit when file exists | `E0-P0A-DOCS5`, `E0-P0A-DOCS6` | Do not treat success redirect or embedded success event as final payment authority. |
| `sdk.md` | `best-practices.md` | adapter and production examples | P6-A | coordinate/edit in same phase | `E0-P0A-SRC4`, `E0-P0A-SRC8` | Remove stale examples across both files. |
| generated `.agents/.claude` payment-integration output | internal source skill files | generated output | P7-A | validate-only / regenerate if commanded by repo tooling | `E0-P0A-GRAPH2` | Never edit generated output as source. |

## Detailed Findings

### Polar reference set

Current state:

The Polar folder has eight markdown references. They cover the main concepts, but the report, source readback, and second-pass official-doc review show stale implementation fields and missing first-class coverage for usage-based billing, Customer Portal, Customer State, and orders/refunds/discounts.

Required state:

```text
Polar references should route agents to the correct surface and document current official API, SDK, checkout, checkout links, embedded checkout/payment methods, webhook, subscription, usage billing, customer portal, customer state, orders/refunds/discounts, benefits, adapter, fee, sandbox, and MoR operational behavior.
```

Evidence:

- `E0-P0A-REPORT1`
- `E0-P0A-SRC1..E0-P0A-SRC8`
- `E0-P0A-FD1..E0-P0A-FD8`
- `E0-P0A-DOCS1..E0-P0A-DOCS6`

Relationship and impact:

- Related file count: 0 for each existing Polar markdown reference.
- Relationship summary: docs-only markdown surfaces with no linked graph flows/tests.
- Impact note: low graph risk, but high product-implementation risk if stale facts remain.

Classification:

`partial`

Allowed next action:

Edit only the targeted Polar docs and preserve correct MoR/auth/sandbox/idempotency concepts.

Forbidden next action:

Do not edit unrelated payment providers or generated output as source.

### Routing surfaces

Current state:

`SKILL.md` and `implementation-workflows.md` route Polar work through the existing eight references. They do not expose usage-based billing, customer portal, Customer State, or orders/refunds/discounts as first-class references, and the quick reference still has stale rate-limit wording.

Required state:

```text
Polar routing should load overview first, then direct agents to the exact reference needed for checkout, products, subscriptions, webhooks, benefits, usage billing, customer portal, customer state, orders/refunds/discounts, SDK/adapters, or best practices.
```

Evidence:

- `E0-P0A-ROUTE1`
- `E0-P0A-ROUTE2`
- `E0-P0A-FD9`
- `E0-P0A-FD10`

Relationship and impact:

- Related file count: 0 for both routing markdown files.
- Impact note: low graph risk, but routing mistakes affect future agent behavior.

Classification:

`partial`

Allowed next action:

Edit only the Polar routing sections.

Forbidden next action:

Do not rewrite unrelated SePay, Stripe, Paddle, or Creem.io routing.

### Polar webhook helper script

Current state:

`scripts/polar-webhook-verify.js` exists and is referenced by the workflow. It is JavaScript source with inbound tests. Official docs prefer SDK/adapters for webhook validation, so this plan should not rely on the helper as the production-default path.

Required state:

```text
For this content refresh, keep the helper inspect-only and route agents toward official SDK/adapters. Any script correction requires expanded scope, impact analysis, code edits, and tests.
```

Evidence:

- `E0-P0A-SCRIPT1`
- `E0-P0A-FD12`

Relationship and impact:

- Related file count: 4 inbound refs plus 16 local relationships.
- Relationship summary: related test file `internal/aicontext/skills/payment-integration/scripts/test-scripts.js`.
- Impact note: high scope warning; not a prohibition, but code/script work needs separate implementation rigor.

Classification:

`partial`

Allowed next action:

Inspect-only, update docs/workflow wording so SDK/adapters are preferred.

Forbidden next action:

Do not edit script or tests under this docs-content plan unless the user expands scope.

## Next Phase Status Decisions

| Plan Item | Actual Status Finding | Required Status / Next-Action Update |
|-----------|-----------------------|--------------------------------------|
| P1-A | Overview, routing, and four new references are implemented. | Completed; validate in P7-A. |
| P2-A | Checkout and products are refreshed with product IDs and practical checkout surfaces. | Completed; validate in P7-A. |
| P3-A | Webhook docs are refreshed and helper script remains inspect-only. | Completed; validate in P7-A. |
| P4-A | Subscription and Customer Portal guidance are refreshed. | Completed; validate in P7-A. |
| P5-A | Usage billing, Customer State, and benefits guidance are refreshed. | Completed; validate in P7-A. |
| P6-A | SDK/adapters, best practices, fees/rates, README cleanup, and orders/refunds/discounts are refreshed. | Completed; validate in P7-A. |
| P7-A | Docs-only validation, supervisor PASS, dead-work cleanup, final staged diff check, and final staged detect are complete. | Commit final staged scope and close. |

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
- [x] If implementation has started, affected Current Status Matrix rows have been refreshed from latest evidence.
- [x] If refreshed statuses changed next work, only the stale next-phase status assumptions, next action, or work steps have been updated before the next phase.

## Final P0 Decision

Choose one:

- [ ] P0 actual-status incomplete. Implementation is blocked.
- [ ] P0 complete. Next phase can proceed unchanged.
- [x] P0 complete. Next phase status, next action, or work steps must be updated before implementation.
- [ ] P0 complete. Target scope is preserve-only.
- [ ] P0 complete. Implementation is blocked by missing authority or evidence.

Decision note:

P0 is complete for plan creation and second-pass plan review. Implementation may proceed from P1-A using the updated phase scope in this plan: add the four missing Polar references, refresh current stale docs in phases, keep the webhook helper script inspect-only unless the user expands scope, and validate through full build, graph refresh, generated-output smoke, and detect-changes before commit.
