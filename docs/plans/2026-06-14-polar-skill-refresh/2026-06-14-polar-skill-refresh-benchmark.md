# Benchmark Ledger

## Metadata

- Title: `Polar Skill Refresh`
- Date: `2026-06-14`
- Plan: `docs/plans/2026-06-14-polar-skill-refresh/2026-06-14-polar-skill-refresh-plan.md`
- Evidence: `docs/plans/2026-06-14-polar-skill-refresh/2026-06-14-polar-skill-refresh-evidence.md`
- Benchmark: `docs/plans/2026-06-14-polar-skill-refresh/2026-06-14-polar-skill-refresh-benchmark.md`
- Actual status: `docs/plans/2026-06-14-polar-skill-refresh/2026-06-14-polar-skill-refresh-actual-status.md`

## Benchmark Rules

This benchmark file records measured inventory counts only. Build/test pass-fail evidence belongs in the evidence ledger unless the timing, count, size, or coverage is the measured target.

## B0 - P0 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P0 | Existing Polar reference files under `references/polar` | files | 8 | 8 | 8 baseline | 8 existing plus planned new references if accepted | 0 at baseline | `rg --files internal\aicontext\skills\payment-integration\references\polar`, `E0-P0A-SRC1..E0-P0A-SRC8` |
| P0 | Planned new Polar reference files | files | 0 | 4 planned | 4 planned | `usage-based-billing.md`, `customer-portal.md`, `customer-state.md`, and `orders-refunds-discounts.md` if P1-A proceeds | +4 | `E0-P0A-REPORT1`, `E0-P0A-DOCS1..E0-P0A-DOCS6` |
| P0 | Direct routing files inspected | files | 2 | 2 | 2 | update only if routing changes or new references are added | 0 | `E0-P0A-ROUTE1`, `E0-P0A-ROUTE2`, `E0-P0A-FD9`, `E0-P0A-FD10` |
| P0 | Existing target markdown docs with Anvien relationship count 0 | files | 11 | 11 | 11 | preserve low graph-risk touch mode | 0 | `E0-P0A-FD1..E0-P0A-FD11` |
| P0 | Inspect-only helper scripts with high graph risk | files | 1 | 1 | 1 | 0 edited unless user expands scope | 0 edited | `E0-P0A-SCRIPT1`, `E0-P0A-FD12` |
| P0 | Current-state finding groups imported from problem report | groups | 11 | 11 | 11 | all mapped to plan phases | 0 | `E0-P0A-REPORT1` |
| P0 | Official Polar source URLs listed in report | URLs | 24+ | 24+ | 24+ | re-check before implementation if stale | 0 | `E0-P0A-REPORT1` |
| P0 | Additional practical Polar docs reviewed after plan draft | docs | 0 | 16+ | 16+ | plan covers practical guide surfaces, not only API deltas | +16+ | `E0-P0A-DOCS2..E0-P0A-DOCS6` |
| P0 | Anvien graph inventory after refresh | nodes / relationships | unknown before refresh | 82486 / 120547 | 82486 / 120547 baseline | graph fresh before file-detail evidence | recorded | `E0-P0A-GRAPH2` |

## B1 - P1 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P1 | Polar surface routes documented in overview | surfaces | partial | 11 | 11 | products, checkout, webhooks, subscriptions, benefits, usage billing, customer portal, customer state, orders/refunds/discounts, SDK/adapters, fees | target met | `E1-P1A-IMPL1` |
| P1 | New reference discoverability | route refs | 0 for new files | 4 | 4 | entrypoint and workflow mention `usage-based-billing.md`, `customer-portal.md`, `customer-state.md`, and `orders-refunds-discounts.md` if added | +4 | `E1-P1A-IMPL2`, `E1-P1A-ROUTE1`, `E1-P1A-ROUTE2`, `E1-P1A-VAL1` |

## B2 - P2 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P2 | Checkout default create field | classification | stale `product_price_id` examples present | `products: [productId]` default | `products: [productId]` default | `products: [productId]` default, price ID deprecated/legacy | corrected | `E2-P2A-IMPL1`, `E2-P2A-IMPL2`, `E6-P6A-VAL1`, `E6-P6A-VAL2` |
| P2 | Product model coverage | topic groups | partial | 8 | 8 | intervals, pricing types, metered, seats, currencies, tax, media, custom fields | target met | `E2-P2A-IMPL3` |
| P2 | Checkout practical surface coverage | topic groups | partial | 8 | 8 | sessions, links, embedded checkout, embedded payment method cross-link, localization, query params, success/return URLs | target met | `E2-P2A-IMPL1`, `E2-P2A-IMPL4` |

## B3 - P3 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P3 | Webhook delivery behavior topics documented | topic groups | partial | 6 | 6 | raw body, SDK/adapters, timeout, retries, endpoint disable, event taxonomy | target met | `E3-P3A-IMPL1` |
| P3 | Event taxonomy coverage | event groups | partial | 8 | 8 | checkout, customer, subscription, order, refund, benefit, product, organization | target met | `E3-P3A-IMPL2` |

## B4 - P4 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P4 | Subscription update field state | classification | stale price/proration examples present | `product_id` and `proration_behavior` current | `product_id` and `proration_behavior` current | `product_id` and `proration_behavior` current | corrected | `E4-P4A-IMPL1` |
| P4 | Customer portal dedicated coverage | files | 0 | 1 | 1 | 1 dedicated reference if P1-A proceeds | +1 | `E4-P4A-IMPL2`, `E1-P1A-VAL1` |
| P4 | Embedded payment-method coverage | topic groups | missing | 4 | 4 | customer session token, modal/inline/React, redirect-result handling, OAT server-only | target met | `E4-P4A-IMPL3` |

## B5 - P5 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P5 | Usage-based billing dedicated coverage | files | 0 | 1 | 1 | 1 dedicated reference if P1-A proceeds | +1 | `E5-P5A-IMPL1`, `E1-P1A-VAL1` |
| P5 | Customer State dedicated coverage | files | 0 | 1 | 1 | 1 dedicated reference if P1-A proceeds | +1 | `E5-P5A-IMPL2`, `E1-P1A-VAL1` |
| P5 | Benefit type coverage | types | partial | 7 | 7 | Credits, License Keys, Feature Flags, File Downloads, GitHub, Discord, Custom | target met | `E5-P5A-IMPL3` |

## B6 - P6 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P6 | Current fee/rate assumptions remaining | stale assumptions | at least `300 req/min`, old Early Member fee as default | 0 unqualified stale assumptions found in Polar references/README | 0 unqualified stale assumptions found in Polar references/README | 0 stale unqualified assumptions | corrected | `E6-P6A-IMPL3`, `E6-P6A-IMPL4`, `E6-P6A-VAL1`, `E6-P6A-VAL2` |
| P6 | Adapter coverage | adapters | partial | 5 | 5 | TypeScript SDK, Next.js, Express, Laravel, BetterAuth | target met | `E6-P6A-IMPL2` |
| P6 | Orders/refunds/discounts dedicated coverage | files | 0 | 1 | 1 | 1 dedicated reference if P1-A proceeds | +1 | `E6-P6A-IMPL1`, `E1-P1A-VAL1` |

## B7 - P7 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P7 | Generated skill output smoke targets | directories | 2 expected | 2 checked | 2 checked | repo `.agents` and `.claude` payment-integration output checked if regenerated | target met | `E7-P7A-GEN1`, `E7-P7A-GEN2` |
| P7 | Detect-changes changed files | files | not measured | 11 tracked changed docs before staging new files | 21 changed files, 20 affected files | docs-only or documented changed set recorded | +21 changed files | `E7-P7A-DETECT1`, `E10-PNC-DETECT1` |

## Non-Benchmarkable Notes

- This is primarily a documentation/skill-source refresh. Runtime performance is not a goal.
- Build and test pass/fail belong in `evidence.md`; timings only become benchmark entries if the implementation intentionally changes build/test/runtime performance.
