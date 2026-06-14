# Plan

## Metadata

- Title: `Polar Skill Refresh`
- Date: `2026-06-14`
- Status: `Closure evidence recorded; final commit pending in same turn`
- Plan: `docs/plans/2026-06-14-polar-skill-refresh/2026-06-14-polar-skill-refresh-plan.md`
- Evidence: `docs/plans/2026-06-14-polar-skill-refresh/2026-06-14-polar-skill-refresh-evidence.md`
- Benchmark: `docs/plans/2026-06-14-polar-skill-refresh/2026-06-14-polar-skill-refresh-benchmark.md`
- Actual status: `docs/plans/2026-06-14-polar-skill-refresh/2026-06-14-polar-skill-refresh-actual-status.md`

## Goal

Refresh the Polar references in the `payment-integration` skill so future agents use current official Polar API, SDK, checkout, webhook, subscription, usage-based billing, customer portal, benefits, adapter, and fee guidance instead of stale product-price, rate-limit, webhook, subscription, and fee assumptions.

## Rules

- Complete P0 actual status before implementation work.
- Update each checklist item immediately when it is completed.
- Record evidence as work completes.
- Record benchmarkable counts or measurements when they are taken.
- Update later phase status assumptions, next actions, and work steps when actual-status evidence changes the repo state.
- After completing a phase or implementation slice and refreshing `actual-status.md`, update the next affected phase's work steps as needed to match the latest repo reality while preserving that phase's original goal, scope, acceptance criteria, and major phase order.
- Do not edit generated `.agents/**`, `.claude/**`, `AGENTS.md`, or `CLAUDE.md` as source of truth.
- Re-check official Polar docs before implementation if this plan is implemented after the current date.

## Problem

`reports/problem/2026-06-14-polar-skill-refresh-analysis.md` records that the current Polar references are directionally useful but stale in important implementation details. The main risks are wrong checkout creation fields, stale production rate limits, stale fee defaults, inconsistent webhook raw-body guidance, incomplete webhook event taxonomy, stale subscription update/proration fields, stale usage-based billing API examples, missing customer portal/customer session coverage, missing Feature Flag benefits, and adapter examples that do not reflect current official Polar guidance.

## Scope

- Primary Polar reference source:
  - `internal/aicontext/skills/payment-integration/references/polar/overview.md`
  - `internal/aicontext/skills/payment-integration/references/polar/products.md`
  - `internal/aicontext/skills/payment-integration/references/polar/checkouts.md`
  - `internal/aicontext/skills/payment-integration/references/polar/subscriptions.md`
  - `internal/aicontext/skills/payment-integration/references/polar/webhooks.md`
  - `internal/aicontext/skills/payment-integration/references/polar/benefits.md`
  - `internal/aicontext/skills/payment-integration/references/polar/sdk.md`
  - `internal/aicontext/skills/payment-integration/references/polar/best-practices.md`
- New Polar reference source if needed for separation:
  - `internal/aicontext/skills/payment-integration/references/polar/usage-based-billing.md`
  - `internal/aicontext/skills/payment-integration/references/polar/customer-portal.md`
  - `internal/aicontext/skills/payment-integration/references/polar/customer-state.md`
  - `internal/aicontext/skills/payment-integration/references/polar/orders-refunds-discounts.md`
- Minimal routing updates if new references are added or workflow order changes:
  - `internal/aicontext/skills/payment-integration/SKILL.md`
  - `internal/aicontext/skills/payment-integration/references/implementation-workflows.md`
- Skill overview cleanup when stale Polar examples are found during readback:
  - `internal/aicontext/skills/payment-integration/README.md`
- Inspect-only / follow-up unless user expands scope:
  - `internal/aicontext/skills/payment-integration/scripts/polar-webhook-verify.js`
  - `internal/aicontext/skills/payment-integration/scripts/test-scripts.js`
- Generated skill output validation:
  - generated `.agents/skills/payment-integration/**` and `.claude/skills/payment-integration/**` are validation outputs, not permanent source edits.

## Non-Goals

- Do not implement a Polar payment integration in an application repo.
- Do not change unrelated payment provider references such as SePay, Stripe, Paddle, or Creem.io.
- Do not update `scripts/polar-webhook-verify.js` in this content refresh unless the user explicitly expands scope to script correction; route agents toward official SDK/adapters first.
- Do not make pricing or fee calculations authoritative without linking to official Polar fee docs, because plan pricing can change.
- Do not rewrite historical reports, old plans, or past evidence that mention stale Polar behavior.
- Do not edit generated home-level installed skills as the permanent source of truth.

## Requirements

- Use the problem report as the implementation authority and re-check official Polar docs before editing if implementation starts in a later session.
- Make checkout creation examples use `products: [productId]` and demote `product_price_id` to deprecated/legacy response context.
- Update API base/rate-limit guidance to production `500/min`, sandbox `100/min`, and unauthenticated license validation/activation/deactivation `3 req/sec`.
- Update fees so Starter/current paid plans are separate from Early Member `4% + 40c + 0.5% subscription fee`.
- Update webhook guidance around Standard Webhooks, raw body, SDK/adapters, endpoint timeout/retry/disable behavior, and event taxonomy.
- Update subscription guidance to use `product_id`, `proration_behavior`, current statuses, cancel-at-period-end, immediate revoke, and Customer Portal self-service boundaries.
- Add or clearly document usage-based billing with `events.ingest`, meters, metered prices, credits, customer meter balance, and server-side billable-event ingestion.
- Add or clearly document Customer Portal/customer sessions with `customer_portal_url`, `token`, `expires_at`, `return_url`, and Core API vs Customer Portal API boundaries.
- Add or clearly document Customer State as the quickest integration surface for entitlement decisions: customer data, active subscriptions, granted benefits, active meters/current balance, external customer ID lookup, and `customer.state_changed`.
- Update benefits to include Feature Flags, official Credits benefit, license-key `benefit_id` scoping, and Customer State / `customer.state_changed` entitlement checks.
- Add or clearly document Polar order operations: order status, `billing_reason`, invoices/receipts, `order.paid` as the canonical paid signal, full/partial refunds, refund-triggered benefit revocation, subscription refund vs subscription cancellation, and discount application modes.
- Add or clearly document checkout operational details from practical docs: Checkout Links vs Checkout Sessions, session URL expiry, link query parameters, checkout-link multi-product choice vs true bundling limitation, embedded checkout `embed_origin`, embedded payment-method sessions, and checkout localization beta limitations.
- Refresh SDK/adapters for TypeScript SDK camelCase, Next.js, Express, BetterAuth, Laravel, and framework-specific checkout/webhook/portal helpers.

## Acceptance Criteria

- Polar reference docs make the correct first decision between checkout sessions, checkout links, products/pricing, subscriptions, webhooks, benefits, customer portal, usage-based billing, SDK/adapters, and fee planning.
- No Polar reference presents `product_price_id` / `productPriceId` as the default checkout creation field for new integrations.
- Rate-limit and fee guidance matches the official docs reviewed for this plan or is explicitly marked "verify current official docs before implementation".
- Webhook guidance uses Standard Webhooks, raw body, SDK/adapters, current retry/timeout/disable behavior, and current event taxonomy including `order.paid`, `checkout.expired`, customer events, and subscription lifecycle events.
- Subscription guidance uses current update/proration fields and distinguishes cancel-at-period-end from immediate revoke.
- Usage-based billing and Customer Portal are discoverable from the skill entrypoint and workflow.
- Customer State is discoverable as an entitlement-sync shortcut, not hidden only inside benefits or webhooks.
- Orders, refunds, and discounts are discoverable enough that agents do not infer paid access from checkout redirects or omit refund/benefit consequences.
- Benefits guidance includes current benefit types and license-key scoping.
- Source skill changes are validated by readback, `git diff --check`, full build before final validation, generated skill smoke where applicable, and Anvien detect-changes before commit.

## Checklist

- [x] P0-A: Complete actual status before implementation work.
  - Goal: establish the real current state for the Polar reference refresh.
  - Work Steps:
    1. Read the Polar problem report and current Polar reference files.
    2. Run `anvien analyze --force`.
    3. Run `anvien file-detail <path> --repo Anvien --json` for each target reference file and direct routing/helper file.
    4. Classify each target as correct, partial, wrong, missing, blocked, or inspect-only.
    5. Update later phase status assumptions, next actions, and work steps from the P0 evidence.
  - Implementation Gate: no Polar reference editing starts until `2026-06-14-polar-skill-refresh-actual-status.md` has a final P0 decision.
  - Acceptance: actual status identifies stale, partial, missing, generated-output, and inspect-only surfaces for this scope.
- [x] P1-A: Update Polar surface map and reference routing.
  - Goal: make the skill route agents to the right Polar surface before they read detailed API examples.
  - Work Steps:
    1. Update `overview.md` to define current Polar surfaces: products/pricing, checkout sessions/links/embedded checkout, webhooks, subscriptions, benefits, usage-based billing, customer portal, SDK/adapters, fees, sandbox, and auth.
    2. Add a concise "choose this surface when..." table.
    3. Add `usage-based-billing.md`, `customer-portal.md`, `customer-state.md`, and `orders-refunds-discounts.md` under `references/polar` if P0 remains valid.
    4. Update `SKILL.md` Quick Reference and platform highlights so Polar no longer says `300 req/min` and new references are discoverable.
    5. Update `implementation-workflows.md` so Polar agents load overview first, then only the needed detailed references.
  - Implementation Gate: P0 confirms Polar docs and route surfaces are editable markdown with low relationship risk; generated output is not edited as source.
  - Acceptance: a future agent can choose checkout, subscriptions, webhooks, benefits, usage billing, customer portal, customer state, orders/refunds/discounts, SDK/adapters, or fees before reading implementation details.
- [x] P2-A: Refresh checkout and product/pricing guidance.
  - Goal: make checkout and product docs current for product IDs, ad-hoc prices, product models, currencies, tax behavior, and custom fields.
  - Work Steps:
    1. Update `checkouts.md` so checkout sessions use `products: [productId]`, success URL `checkout_id={CHECKOUT_ID}`, `external_customer_id`, customer prefill, `customer_ip_address`, embedded checkout, and ad-hoc `prices`.
    2. Mark `product_price_id` / `productPriceId` as deprecated/legacy response context, not the default create field.
    3. Document Checkout Links separately from Checkout Sessions: checkout links are persistent, each visit creates a short-lived session, session URLs should not be shared, link query parameters can prefill/override fields, and multiple products on a link are a choice rather than a bundled order.
    4. Add embedded checkout details: `@polar-sh/checkout`, `data-polar-checkout`, `PolarEmbedCheckout`, checkout events, `embed_origin`, theme, locale, and the rule that client-side success/confirmed events do not replace webhook/order verification.
    5. Add checkout localization guidance: `locale` parameter/querystring, beta status, supported-language caveats, and untranslated email/error limitations.
    6. Update `products.md` with current billing intervals, `recurring_interval_count`, pricing models, metered prices stacked on fixed base pricing, multiple currencies, tax behavior, product media/visibility, custom fields, and seat-based pricing constraints.
    7. Cross-link checkout/product docs to `usage-based-billing.md`, `customer-portal.md`, and `orders-refunds-discounts.md` only where the integration surface changes.
  - Implementation Gate: P1-A routing is in place or actual-status has been refreshed with an alternate route that preserves the phase goal.
  - Acceptance: `checkouts.md` and `products.md` no longer teach stale checkout/product API shapes as current defaults.
- [x] P3-A: Refresh webhook security, delivery, and event taxonomy.
  - Goal: make `webhooks.md` safe for production Polar event processing.
  - Work Steps:
    1. Update webhook setup around Standard Webhooks and official SDK/adapters.
    2. Require raw body for verification examples and prefer `validateEvent` or framework adapters.
    3. Correct delivery behavior: 10-second timeout, recommended 2-second response, up to 10 retries, endpoint disable after 10 consecutive failed deliveries, and non-2xx redirect handling.
    4. Add current event taxonomy and clarify `order.created` vs `order.paid`.
    5. Add cancellation and renewal event sequences.
    6. Add `customer.state_changed` and Customer State as the preferred compact entitlement-sync webhook when the app only needs customer access state.
    7. Keep `scripts/polar-webhook-verify.js` inspect-only unless the user explicitly expands scope; update docs/workflow so agents prefer SDK/adapters.
  - Implementation Gate: P2-A has not changed checkout/order assumptions; if it has, refresh actual-status and update this phase's stale work steps only.
  - Acceptance: `webhooks.md` documents current verification, delivery behavior, event handling, and the script helper is not treated as the preferred production path.
- [x] P4-A: Refresh subscriptions and Customer Portal guidance.
  - Goal: make subscription lifecycle, plan changes, proration, cancellation, revoke, and customer self-service guidance current.
  - Work Steps:
    1. Update `subscriptions.md` to use `product_id`, `proration_behavior`, current proration values, current statuses, `discount_id`, `trial_end`, and current cancel/revoke/uncancel behavior.
    2. Add or fill `customer-portal.md` with customer sessions, `customer_portal_url`, token expiry, `return_url`, customer ID vs external customer ID creation, Customer Portal API scope boundaries, hosted portal URL, transactional email portal links, payment methods, orders, subscriptions, benefits, seats, and portal settings.
    3. Document hosted Customer Portal constraints: portal cannot be disabled, hosted portal is required for PCI-safe default payment method recovery, and pre-authenticated portal links should be generated fresh at click time rather than stored.
    4. Add embedded payment-method guidance: create a short-lived customer session server-side, use `@polar-sh/checkout/payment-method`, support modal/inline/React flows, handle redirect-based payment-method results, and keep OATs off the client.
    5. Update workflow routing so self-service portal work does not get implemented as privileged Core API logic.
    6. Cross-link subscription docs to webhook event sequences, Customer Portal self-service actions, Customer State, and order/refund consequences.
  - Implementation Gate: P1-A has created/confirmed the customer portal reference location.
  - Acceptance: subscription and customer portal docs reflect current official API fields and customer-vs-merchant responsibility boundaries.
- [x] P5-A: Refresh usage-based billing and benefits guidance.
  - Goal: make usage metering, credits, customer meter balance, and entitlement guidance current and discoverable.
  - Work Steps:
    1. Add or fill `usage-based-billing.md` with `polar.events.ingest`, `name`, `customerId` / `externalCustomerId`, `metadata`, meters, aggregation, metered prices, credits, customer meters, and server-side ingestion for billable usage.
    2. Add ingestion strategy coverage for LLM usage, S3 operations, streams, and delta-time usage only as strategy guidance, not as required implementation for every app.
    3. Add cost-insights `_cost` metadata only as an optional cost/profit tracking note, not as a billing requirement.
    4. Add or fill `customer-state.md` with the single-call/single-webhook entitlement model: customer data, active subscriptions, granted benefits, active meters/current balance, external ID lookup, and `customer.state_changed`.
    5. Update `benefits.md` with Credits, License Keys, Feature Flags, File Downloads, GitHub, Discord, Custom benefits, grant lifecycle, Customer State, and `customer.state_changed`.
    6. Add license-key guidance to validate `benefit_id` when multiple license-key benefits exist in one organization.
  - Implementation Gate: P2-A product pricing docs are current enough to cross-link metered prices and credits without duplicating stale field names.
  - Acceptance: agents can implement usage billing, customer-state entitlement sync, and benefits from current official shapes without relying on stale `events.create`, `event_name`, or `properties` examples.
- [x] P6-A: Refresh SDK/adapters, best practices, fees, and environment guidance.
  - Goal: make implementation examples and production patterns current across Polar SDK, adapters, fees, sandbox, and environment variables.
  - Work Steps:
    1. Add or fill `orders-refunds-discounts.md` with order creation contexts, `billing_reason`, statuses, invoices/receipts, `order.paid` as the canonical paid signal, partial/full refunds, benefit revocation rules, refund-vs-cancel distinction for subscriptions, and discount types/restrictions/application modes.
    2. Update `sdk.md` with TypeScript SDK install/quickstart, camelCase-vs-snake_case note, official Python/Go/PHP links as applicable, and framework adapters.
    3. Refresh Next.js, Express, Laravel, and BetterAuth examples with current checkout, webhook, portal, usage, customer state, and products parameters; distinguish framework adapter helpers from long-form framework guides when the guide uses direct HTTP or third-party webhook packages.
    4. Update `best-practices.md` around plan-aware fees, sandbox/production isolation, sandbox email limitations, Core API vs Customer Portal API boundaries, pagination, rate limits, retries, server-only access tokens, checkout/webhook idempotency, event-to-database mapping, orders/refunds/discounts, customer portal, customer state, usage billing, and generated-output caveats.
    5. Add MoR operational caveats to overview/best practices: supported countries, acceptable use, account reviews, tax behavior, and fee verification must be checked against official docs before launch.
    6. Remove stale `productPriceId`, old fee defaults, stale `300 req/min`, and examples that imply success redirects are sufficient payment proof.
    7. Update `SKILL.md` and `implementation-workflows.md` again if final SDK/best-practice routing differs from P1 assumptions.
  - Implementation Gate: P3-A through P5-A have established current webhook, subscription, customer portal, usage, and benefits terms.
  - Acceptance: SDK and best-practice docs are current, repo-agnostic, and route agents away from stale API fields and stale fee/rate assumptions.
- [x] P7-A: Validate source skill, generated outputs, and graph change evidence.
  - Goal: prove the Polar skill refresh is internally consistent and generated agent surfaces can consume it.
  - Work Steps:
    1. Read back every changed Polar reference file and route file.
    2. Run `git diff --check`.
    3. Run a full build before accepting final validation evidence for source skill changes unless the user explicitly narrows validation to docs-only evidence.
    4. Run `anvien analyze --force` after source changes to refresh graph evidence.
    5. Smoke-check generated `.agents/skills/payment-integration` and `.claude/skills/payment-integration` output if those directories are regenerated in the working tree.
    6. Run `anvien detect-changes --repo Anvien --scope all` before commit.
  - Implementation Gate: implementation phases P1-A through P6-A are complete or explicitly blocked.
  - Acceptance: validation evidence records what each command proves, generated output is not used as source of truth, and detect-changes is recorded before commit.
- [x] Pn-A: Call supervisor for the implemented-plan acceptance loop.
  - Goal: verify the completed plan work against the accepted plan, actual-status decisions, evidence, benchmark, changed files, generated output, and validation results before closure.
  - Work Steps:
    1. Call the supervisor skill to review the full completed plan work.
    2. If supervisor fails the work, return to the responsible implementation workflow/skill for the failed scope only.
    3. Re-run supervisor review after the fix.
    4. Repeat until supervisor passes or records a blocker.
  - Implementation Gate: all planned implementation phases must be completed or explicitly blocked before this review.
  - Acceptance: supervisor review passes, or the plan records a blocker with evidence and no closure is performed.
- [x] Pn-B: Remove dead work created during this plan.
  - Goal: ensure the final diff contains only artifacts that still serve the accepted plan.
  - Work Steps:
    1. Review files, sections, generated output, tests, temp files, and plan artifacts created or modified during this plan.
    2. Remove or rewrite any artifact made obsolete by actual-status findings, user corrections, failed approaches, or phase status updates.
    3. Verify no rejected approach, stale placeholder, unused generated output, or dead helper artifact remains in the final diff.
    4. Call supervisor to review the dead-work cleanup.
    5. If supervisor fails the cleanup, return to the responsible implementation workflow/skill for the failed cleanup scope only, then re-run supervisor review.
  - Implementation Gate: only remove artifacts created by this plan unless the user explicitly approves broader cleanup.
  - Acceptance: final `git diff/status` contains no dead plan-created artifacts, supervisor passes the cleanup, and evidence records what was removed or preserved.
- [x] Pn-C: Close the plan.
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

- This is a docs/skill-source refresh, but stale content can still cause broken payment implementations.
- `polar-webhook-verify.js` is source code with high graph risk and inbound tests; this content plan keeps it inspect-only unless the user expands scope.
- Fee and rate-limit facts are external and time-sensitive; implementation must re-check official docs if delayed.
- Generated skill outputs are validation artifacts; edit `internal/aicontext/**` source first, then regenerate through repo tooling if needed.
