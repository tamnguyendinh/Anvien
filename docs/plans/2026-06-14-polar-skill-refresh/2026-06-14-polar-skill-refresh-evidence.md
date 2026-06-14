# Evidence Ledger

## Metadata

- Title: `Polar Skill Refresh`
- Date: `2026-06-14`
- Plan: `docs/plans/2026-06-14-polar-skill-refresh/2026-06-14-polar-skill-refresh-plan.md`
- Evidence: `docs/plans/2026-06-14-polar-skill-refresh/2026-06-14-polar-skill-refresh-evidence.md`
- Benchmark: `docs/plans/2026-06-14-polar-skill-refresh/2026-06-14-polar-skill-refresh-benchmark.md`
- Actual status: `docs/plans/2026-06-14-polar-skill-refresh/2026-06-14-polar-skill-refresh-actual-status.md`

## Evidence Rules

The evidence file explains why the work is known to be correct.

Evidence IDs use stable, phase-scoped names so `plan.md`, `actual-status.md`, `benchmark.md`, and later agents can reference exact proof without ambiguity.

Format:

```text
E<phase>-<item>-<kind><n>
```

Rules:

- `E<phase>` matches the plan phase number: `E0` for `P0`, `E1` for `P1`, `E2` for `P2`, and so on.
- `<item>` matches the checklist item without the dash: `P0A`, `P1A`, `P2A`.
- `<kind>` is plan-local. This plan uses `GRAPH`, `FD`, `SRC`, `REPORT`, `ROUTE`, `SCRIPT`, `VAL`, `IMPL`, `GEN`, `DETECT`, `SUP`, and `COMMIT` as needed.
- `<n>` is a 1-based sequence number within that phase item and kind.
- Do not reuse an evidence ID for different facts.
- Reference exact IDs from `actual-status.md` and `benchmark.md`.

## E0 - P0 Evidence

Matching plan item(s): `P0-A`

- `E0-P0A-GRAPH1`: Ran `anvien --help` from `E:\Anvien` to confirm the local Anvien command surface before plan work. Command listed `analyze`, `file-detail`, `detect-changes`, and related graph commands.
- `E0-P0A-GRAPH2`: Ran `anvien analyze --force` from `E:\Anvien`. Result: analyzed `E:\Anvien`; files scanned 1397; parsed code 673; failed 0; indexed documents 487; graph nodes 82486; relationships 120547; graph path `.anvien/graph.json`.
- `E0-P0A-REPORT1`: Read `reports/problem/2026-06-14-polar-skill-refresh-analysis.md`. The report records stale or missing Polar coverage for checkout `products`, rate limits, fees, webhooks, event taxonomy, subscriptions, usage-based billing, benefits, customer portal, SDK/adapters, and product model details.
- `E0-P0A-SRC1`: Read `internal/aicontext/skills/payment-integration/references/polar/overview.md`. Current content has useful MoR/auth/sandbox concepts but stale production rate-limit guidance and insufficient routing for customer portal and usage billing.
- `E0-P0A-SRC2`: Read `internal/aicontext/skills/payment-integration/references/polar/checkouts.md`. Current content still uses `product_price_id` as a primary checkout creation shape and needs current `products`/ad-hoc price guidance.
- `E0-P0A-SRC3`: Read `internal/aicontext/skills/payment-integration/references/polar/webhooks.md`. Current content has important webhook concepts but needs raw-body/Standard Webhooks delivery updates and current event taxonomy.
- `E0-P0A-SRC4`: Read `internal/aicontext/skills/payment-integration/references/polar/sdk.md`. Current content needs current SDK/adapters, camelCase guidance, and framework helper refresh.
- `E0-P0A-SRC5`: Read `internal/aicontext/skills/payment-integration/references/polar/products.md`. Current content has useful product concepts but needs current product/pricing model fields, metered pricing, currencies, tax behavior, and deprecation cleanup.
- `E0-P0A-SRC6`: Read `internal/aicontext/skills/payment-integration/references/polar/subscriptions.md`. Current content uses stale subscription update/proration examples and needs current status/cancel/revoke guidance.
- `E0-P0A-SRC7`: Read `internal/aicontext/skills/payment-integration/references/polar/benefits.md`. Current content covers several benefits but misses Feature Flags and official Credits benefit details.
- `E0-P0A-SRC8`: Read `internal/aicontext/skills/payment-integration/references/polar/best-practices.md`. Current content includes stale fee/rate examples and stale checkout field examples.
- `E0-P0A-ROUTE1`: Read `internal/aicontext/skills/payment-integration/SKILL.md`. Polar quick reference lists eight Polar files and platform highlights still say `300 req/min`.
- `E0-P0A-ROUTE2`: Read `internal/aicontext/skills/payment-integration/references/implementation-workflows.md`. Polar workflow does not route usage-based billing or customer portal as first-class references.
- `E0-P0A-SCRIPT1`: Read `internal/aicontext/skills/payment-integration/scripts/polar-webhook-verify.js`. Script implements a custom Standard Webhooks-style verifier and is referenced by workflow, but this plan treats it as inspect-only because official docs prefer SDK/adapters and the script has source-code blast radius.
- `E0-P0A-FD1`: `anvien file-detail internal/aicontext/skills/payment-integration/references/polar/overview.md --repo Anvien --json` returned markdown docs, parsed, 0 local/outbound/inbound relationships, 0 linked flows/tests, low risk, stale false.
- `E0-P0A-FD2`: `anvien file-detail internal/aicontext/skills/payment-integration/references/polar/checkouts.md --repo Anvien --json` returned markdown docs, parsed, 0 local/outbound/inbound relationships, 0 linked flows/tests, low risk, stale false.
- `E0-P0A-FD3`: `anvien file-detail internal/aicontext/skills/payment-integration/references/polar/webhooks.md --repo Anvien --json` returned markdown docs, parsed, 0 local/outbound/inbound relationships, 0 linked flows/tests, low risk, stale false.
- `E0-P0A-FD4`: `anvien file-detail internal/aicontext/skills/payment-integration/references/polar/sdk.md --repo Anvien --json` returned markdown docs, parsed, 0 local/outbound/inbound relationships, 0 linked flows/tests, low risk, stale false.
- `E0-P0A-FD5`: `anvien file-detail internal/aicontext/skills/payment-integration/references/polar/products.md --repo Anvien --json` returned markdown docs, parsed, 0 local/outbound/inbound relationships, 0 linked flows/tests, low risk, stale false.
- `E0-P0A-FD6`: `anvien file-detail internal/aicontext/skills/payment-integration/references/polar/subscriptions.md --repo Anvien --json` returned markdown docs, parsed, 0 local/outbound/inbound relationships, 0 linked flows/tests, low risk, stale false.
- `E0-P0A-FD7`: `anvien file-detail internal/aicontext/skills/payment-integration/references/polar/benefits.md --repo Anvien --json` returned markdown docs, parsed, 0 local/outbound/inbound relationships, 0 linked flows/tests, low risk, stale false.
- `E0-P0A-FD8`: `anvien file-detail internal/aicontext/skills/payment-integration/references/polar/best-practices.md --repo Anvien --json` returned markdown docs, parsed, 0 local/outbound/inbound relationships, 0 linked flows/tests, low risk, stale false.
- `E0-P0A-FD9`: `anvien file-detail internal/aicontext/skills/payment-integration/SKILL.md --repo Anvien --json` returned markdown docs, parsed, 0 local/outbound/inbound relationships, 0 linked flows/tests, low risk, stale false.
- `E0-P0A-FD10`: `anvien file-detail internal/aicontext/skills/payment-integration/references/implementation-workflows.md --repo Anvien --json` returned markdown docs, parsed, 0 local/outbound/inbound relationships, 0 linked flows/tests, low risk, stale false.
- `E0-P0A-FD11`: `anvien file-detail reports/problem/2026-06-14-polar-skill-refresh-analysis.md --repo Anvien --json` returned markdown docs, parsed, 0 local/outbound/inbound relationships, 0 linked flows/tests, low risk, stale false.
- `E0-P0A-FD12`: `anvien file-detail internal/aicontext/skills/payment-integration/scripts/polar-webhook-verify.js --repo Anvien --json` returned JavaScript source, parsed, 28 symbols, 16 local relationships, 4 inbound references from `scripts/test-scripts.js`, 91 unresolved entries, high risk, stale false.
- `E0-P0A-VAL1`: `Test-Path docs/plans/2026-06-14-polar-skill-refresh` returned false before plan creation, so the selected standard plan directory did not collide with an existing plan.
- `E0-P0A-DOCS1`: Re-read the standard plan set after user requested a second pass against practical Polar docs. The existing plan covered core checkout/products/webhooks/subscriptions/usage/portal/benefits/SDK/fees, but did not expose Customer State or orders/refunds/discounts as first-class references.
- `E0-P0A-DOCS2`: Reviewed official Polar documentation index at `https://polar.sh/docs/llms.txt`. The index exposes additional practical surfaces relevant to the plan: Customer State, Checkout Links, Embedded Checkout, Embedded Payment Method, Checkout Localization, Orders, Refunds, Discounts, Custom Fields, Tax Inclusive Pricing, Customer Portal navigation/settings, and framework guides.
- `E0-P0A-DOCS3`: Reviewed official Next.js and Laravel integration guides. Next.js uses `@polar-sh/nextjs` `Checkout`, `CustomerPortal`, and `Webhooks`; Laravel guide demonstrates direct HTTP checkout with `products`, confirmation polling, Standard Webhooks verification through PHP packages, queue processing, and subscription-active/revoked handling.
- `E0-P0A-DOCS4`: Reviewed official Customer State and Customer Portal docs. Customer State is the compact entitlement surface containing customer data, active subscriptions, granted benefits, active meters/current balances, and `customer.state_changed`; Customer Portal has default hosted URLs, pre-authenticated links, email links, settings, payment-method recovery, and hosted-portal constraints.
- `E0-P0A-DOCS5`: Reviewed official checkout practical docs. Checkout Links are persistent and create short-lived sessions on visit; multiple products on a Checkout Link are a customer choice, not a bundle; Embedded Checkout needs `@polar-sh/checkout` or data attributes plus `embed_origin`; Embedded Payment Method uses a short-lived customer session token; localization is beta with limitations.
- `E0-P0A-DOCS6`: Reviewed official Orders, Refunds, Discounts, Custom Fields, Tax Inclusive Pricing, Sandbox, Authentication, and API Overview docs. The plan needs explicit coverage for `billing_reason`, order status, invoices/receipts, refund benefit consequences, discount restrictions/application modes, custom field types and `custom_field_data`, tax behavior, sandbox email limitations, Core API vs Customer Portal API, pagination, and token leak/security rules.
- `E0-P0A-IMPL1`: Updated the plan, actual-status, and benchmark ledgers to add planned first-class coverage for `customer-state.md` and `orders-refunds-discounts.md`, and to add practical docs requirements for checkout links, embedded checkout, embedded payment method, localization, Customer Portal constraints, orders/refunds/discounts, API pagination, sandbox limitations, and MoR operational caveats.

## E1 - P1 Evidence

Matching plan item(s): `P1-A`

- `E1-P1A-IMPL1`: Updated `internal/aicontext/skills/payment-integration/references/polar/overview.md` with a current Polar surface-selection map covering products/pricing, checkout sessions/links/embedded checkout, webhooks, subscriptions, benefits, usage billing, customer portal, Customer State, orders/refunds/discounts, SDK/adapters, fees, sandbox, auth, and rate limits.
- `E1-P1A-IMPL2`: Added four first-class Polar references under `references/polar`: `usage-based-billing.md`, `customer-portal.md`, `customer-state.md`, and `orders-refunds-discounts.md`.
- `E1-P1A-ROUTE1`: Updated `internal/aicontext/skills/payment-integration/SKILL.md` so the Polar quick reference lists 12 Polar references and no longer presents the stale `300 req/min` platform highlight.
- `E1-P1A-ROUTE2`: Updated `internal/aicontext/skills/payment-integration/references/implementation-workflows.md` so Polar work loads `overview.md` first, then routes to the specific current references for checkout, subscriptions, webhooks, usage billing, customer portal, Customer State, orders/refunds/discounts, SDK/adapters, and best practices.
- `E1-P1A-VAL1`: `rg --files internal\aicontext\skills\payment-integration\references\polar | Measure-Object` returned 12 files after implementation, matching the eight existing Polar references plus the four planned new references.

## E2 - P2 Evidence

Matching plan item(s): `P2-A`

- `E2-P2A-IMPL1`: Rewrote `references/polar/checkouts.md` so new checkout session examples use `products: [productId]`, success URLs with `checkout_id={CHECKOUT_ID}`, customer prefill/identity fields, `customer_ip_address`, ad-hoc `prices`, discounts, Checkout Links, embedded checkout, embedded payment method cross-links, localization, and `order.paid` verification rules.
- `E2-P2A-IMPL2`: Marked `product_price_id` / `productPriceId` in `checkouts.md` as deprecated or legacy response context, not the default create field for new Polar work.
- `E2-P2A-IMPL3`: Rewrote `references/polar/products.md` with current product/pricing guidance for one-time, recurring, metered, custom, free, and seat-based pricing; billing intervals and interval counts; stacked fixed plus metered pricing; multiple currencies; tax behavior; product media/visibility; benefits; custom fields; and update constraints.
- `E2-P2A-IMPL4`: Cross-linked checkout/product docs to usage-based billing, customer portal, Customer State, and orders/refunds/discounts where those surfaces change implementation behavior.

## E3 - P3 Evidence

Matching plan item(s): `P3-A`

- `E3-P3A-IMPL1`: Rewrote `references/polar/webhooks.md` around Standard Webhooks, raw-body verification, official SDK/adapters, Next.js/Express examples, idempotent processing, delivery timeout/retry/disable behavior, redirects as failed delivery, and low-latency acknowledgement patterns.
- `E3-P3A-IMPL2`: Added current event taxonomy coverage in `webhooks.md`, including checkout, customer, subscription, order, refund, benefit, product, and organization groups, with `order.created` vs `order.paid`, cancellation, renewal, and `customer.state_changed` entitlement-sync sequences.
- `E3-P3A-IMPL3`: Preserved `scripts/polar-webhook-verify.js` as inspect-only and adjusted docs/workflow so future agents prefer official SDK/adapters for production webhook verification.

## E4 - P4 Evidence

Matching plan item(s): `P4-A`

- `E4-P4A-IMPL1`: Rewrote `references/polar/subscriptions.md` with current lifecycle/status guidance, `product_id`, `proration_behavior`, `discount_id`, `trial_end`, cancel-at-period-end, uncancel, immediate revoke, failed-payment handling, renewal/order event sequencing, and Customer Portal boundaries.
- `E4-P4A-IMPL2`: Added `references/polar/customer-portal.md` covering customer sessions, `customer_portal_url`, `token`, `expires_at`, `return_url`, customer ID vs external customer ID creation, Core API vs Customer Portal API scope, hosted portal links, transactional email links, portal settings, payment methods, orders, subscriptions, benefits, seats, and fresh pre-authenticated portal link generation.
- `E4-P4A-IMPL3`: Added embedded payment-method guidance in `customer-portal.md` and checkout cross-links: create a short-lived customer session server-side, keep organization access tokens off the client, use modal/inline/React flows, and handle redirect-based payment method results.

## E5 - P5 Evidence

Matching plan item(s): `P5-A`

- `E5-P5A-IMPL1`: Added `references/polar/usage-based-billing.md` with `events.ingest`, `name`, `customerId` / `externalCustomerId`, `metadata`, meters, aggregation, metered prices, credits, customer meter balances, server-side billable-event ingestion, event immutability/backdating, ingestion strategies, and optional `_cost` metadata guidance.
- `E5-P5A-IMPL2`: Added `references/polar/customer-state.md` with the single-call/single-webhook entitlement model covering customer data, active subscriptions, granted benefits, active meters/current balance, external customer ID lookup, local entitlement snapshot strategy, and `customer.state_changed`.
- `E5-P5A-IMPL3`: Rewrote `references/polar/benefits.md` with current benefit coverage for Credits, License Keys, Feature Flags, File Downloads, GitHub, Discord, and Custom benefits; grant lifecycle; Customer State; `customer.state_changed`; and license-key `benefit_id` scoping.

## E6 - P6 Evidence

Matching plan item(s): `P6-A`

- `E6-P6A-IMPL1`: Added `references/polar/orders-refunds-discounts.md` with order states, `billing_reason`, invoices/receipts, `order.paid` as the canonical paid signal, partial/full refunds, refund-triggered benefit consequences, subscription refund vs cancellation boundaries, and discount modes/restrictions.
- `E6-P6A-IMPL2`: Rewrote `references/polar/sdk.md` with current TypeScript SDK quickstart, camelCase vs REST snake_case guidance, official language SDK boundaries, and Next.js, Express, BetterAuth, Laravel, checkout embed, portal, usage, products, and webhook patterns.
- `E6-P6A-IMPL3`: Rewrote `references/polar/best-practices.md` with plan-aware fee checks, sandbox/production isolation, sandbox email limitations, Core API vs Customer Portal API boundaries, pagination, server-only access tokens, lazy clients, rate limits, retries, idempotency, checkout/webhook database mapping, orders/refunds/discounts, customer portal, Customer State, usage billing, MoR launch checks, and generated-output caveats.
- `E6-P6A-IMPL4`: Updated `internal/aicontext/skills/payment-integration/README.md` so the source skill overview lists the new Polar references, removes the stale Polar checkout helper example, and directs new Polar checkout work to `references/polar/checkouts.md` with `products: [productId]`; the legacy helper remains inspect-only until a separate script-change scope.
- `E6-P6A-VAL1`: Ran `rg -n "product_price_id|productPriceId|300 req/min|300 requests|events\.create|event_name|price_xxx|product_price" internal\aicontext\skills\payment-integration` after implementation. Remaining `productPriceId` hits are in inspect-only helper script/tests; Polar reference hits are negative/deprecated warnings; Stripe `price_xxx` examples are unrelated provider docs.
- `E6-P6A-VAL2`: Ran `rg -n "usage-based-billing|customer-portal|customer-state|orders-refunds-discounts|products: \[productId\]|order\.paid|customer\.state_changed|500/min|100/min|3 req/sec" internal\aicontext\skills\payment-integration` after implementation. Results show new references discoverable from `SKILL.md`, README, workflow, overview, and detailed Polar docs; checkout examples use `products: [productId]`; paid access guidance uses `order.paid`; entitlement sync uses `customer.state_changed`.

## E7 - P7 Evidence

Matching plan item(s): `P7-A`

- `E7-P7A-VAL1`: `git diff --check` from `E:\Anvien` passed with no whitespace errors.
- `E7-P7A-BUILD1`: `npm run full-build` was started because the repo AGENTS rule normally requires a full build before validation, then the user interrupted and corrected scope: this task is documentation-only and build does not provide useful runtime proof. Full build was not re-run and is not used as acceptance evidence for this docs-only validation.
- `E7-P7A-GRAPH1`: Ran `anvien analyze --force` after source changes. Result: analyzed `E:\Anvien`; files scanned 1405; parsed code 673; failed 0; indexed documents 495; graph nodes 82521; relationships 120582; graph path `.anvien/graph.json`.
- `E7-P7A-FD1`: Post-change `anvien file-detail` for the four new Polar references returned markdown docs, parsed, 0 local/outbound/inbound relationships, 0 linked flows/tests, low risk, stale false, changedSinceAnalyze false.
- `E7-P7A-FD2`: Post-change `anvien file-detail` for the eight existing changed Polar references returned markdown docs, parsed, 0 local/outbound/inbound relationships, 0 linked flows/tests, low risk, stale false, changedSinceAnalyze false.
- `E7-P7A-FD3`: Post-change `anvien file-detail` for `SKILL.md`, `references/implementation-workflows.md`, and `README.md` returned markdown docs, parsed, 0 local/outbound/inbound relationships, 0 linked flows/tests, low risk, stale false, changedSinceAnalyze false.
- `E7-P7A-VAL2`: Readback grep over source skill docs confirmed new references are routed from `SKILL.md`, README, workflow, overview, and detailed Polar docs; checkout examples use `products: [productId]`; paid access guidance uses `order.paid`; entitlement sync uses `customer.state_changed`; usage billing uses `events.ingest`.
- `E7-P7A-GEN1`: Generated repo output smoke confirmed both `.agents/skills/payment-integration` and `.claude/skills/payment-integration` exist and contain the new Polar references and routing.
- `E7-P7A-GEN2`: Generated repo output smoke also found `productPriceId` in generated checkout helper scripts/tests. This matches the source helper-script surface intentionally kept inspect-only by the plan; README/source references now route new Polar checkout work away from that helper.
- `E7-P7A-DETECT1`: `anvien detect-changes --repo Anvien --scope all` on the current tracked diff passed. Summary: changed files 11, affected files 11, changed app layer docs, changed functional area documentation, affected processes none, risk level low. Because new files were still untracked, final staged detect must be run before commit.

## Closure Evidence

- `E8-PNA-SUP1`: Supervisor PASS report written at `reports/Supervisor/rp_supervisor_260614_200034_by_gpt-5-codex_polar-skill-refresh-docs.md`. Verdict: PASS for docs-only Polar skill-refresh scope. The report clears source references/routing, generated repo output smoke, Anvien graph/file-detail evidence, helper-script inspect-only caveat, and user-approved build skip.
- `E9-PNB-CLEAN1`: Dead-work cleanup review found no plan-created temp files, rejected approaches, placeholder docs, or stale `Pending implementation` / `Pending validation` markers. `rg` over the active plan/source/report/supervisor scope found no `TODO`, `PLACEHOLDER`, `Draft`, `Pending implementation`, `Pending validation`, or `implementation not started` hits. Existing `.tmp` entries predate this plan and were not created by this scope.
- `E10-PNC-VAL1`: Final staged `git diff --cached --check` passed with no whitespace errors.
- `E10-PNC-DETECT1`: Final staged `anvien detect-changes --repo Anvien --scope all` passed after staging source docs, new Polar references, plan ledgers, problem report, and supervisor report. Summary: changed files 21, affected files 20, changed app layer docs, changed functional areas documentation and reporting, affected processes none, risk level low.
- `E10-PNC-COMMIT1`: Final commit executed after recording this evidence; commit hash is reported in the final assistant response because a commit cannot include its own final hash in the same tree.
