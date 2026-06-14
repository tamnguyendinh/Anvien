# Benchmark Ledger

## Metadata

- Title: `SePay Skill Refresh`
- Date: `2026-06-14`
- Plan: `docs/plans/2026-06-14-sepay-skill-refresh/2026-06-14-sepay-skill-refresh-plan.md`
- Evidence: `docs/plans/2026-06-14-sepay-skill-refresh/2026-06-14-sepay-skill-refresh-evidence.md`
- Benchmark: `docs/plans/2026-06-14-sepay-skill-refresh/2026-06-14-sepay-skill-refresh-benchmark.md`
- Actual status: `docs/plans/2026-06-14-sepay-skill-refresh/2026-06-14-sepay-skill-refresh-actual-status.md`

## Benchmark Rules

This benchmark file records measured inventory counts only. Build/test pass-fail evidence belongs in the evidence ledger unless the timing, count, size, or coverage is the measured target.

## B0 - P0 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P0 | Existing SePay reference files under `references/sepay` | files | 6 | 7 | 7 | 6 existing plus planned new gateway reference if accepted | +1 | `rg --files internal\aicontext\skills\payment-integration\references\sepay`, E0-P0A-SRC1..E0-P0A-SRC6, E1-P1A-IMPL2 |
| P0 | Planned new SePay reference files | files | 0 | 1 created | 1 created | 1 `payment-gateway.md` if P1-A proceeds | +1 | E0-P0A-REPORT1, E0-P0A-ROUTE1, E1-P1A-IMPL2 |
| P0 | Direct routing files inspected | files | 2 | 2 | pending | 2 if new reference is added | pending | E0-P0A-ROUTE1, E0-P0A-FD7, E0-P0A-FD8 |
| P0 | Existing target markdown files with Anvien relationship count 0 | files | 6 | 6 | pending | preserve low graph-risk touch mode | pending | E0-P0A-FD1..E0-P0A-FD6 |
| P0 | Inspect-only helper scripts with high graph risk | files | 1 | 1 | pending | 0 edited unless user expands scope | pending | E0-P0A-SRC7, E0-P0A-FD9 |
| P0 | Current-state findings imported from problem report | findings | 16 | 16 | pending | all mapped to plan phases | pending | E0-P0A-REPORT1 |
| P0 | Official SePay source URLs listed in report | URLs | 20 | 20 | pending | re-check before implementation if stale | pending | E0-P0A-REPORT1 |

## B1 - P1 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P1 | SePay surface routes documented in overview | surfaces | pending | 6 | 6 | API v2, Webhooks, Payment Gateway/IPN, QR, Order VA, optional products | met | `E1-P1A-IMPL1`, `E1-P1A-SRC1` |
| P1 | Dedicated gateway/IPN reference discoverability | route refs | pending | 2 | 2 | entrypoint and workflow mention `payment-gateway.md` if added | met | `E1-P1A-ROUTE1`, `E1-P1A-ROUTE2` |

## B2 - P2 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P2 | API default state | classification | legacy v1 default | API v2 default, v1 legacy | API v2 default, v1 legacy | API v2 default, v1 legacy | met | `E2-P2A-SRC1`, `E2-P2A-SRC2` |
| P2 | Order VA bank-specific rule coverage | rule groups | partial | 3 | 3 | BIDV, Sacombank, Vietcombank constraints represented | met | `E2-P2A-SRC3` |

## B3 - P3 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P3 | Webhook auth methods documented | methods | 3 | 4 | 4 | None, API Key, HMAC-SHA256, OAuth2 | +1 | `E3-P3A-SRC1` |
| P3 | Webhook operations topics documented | topic groups | partial | 7 | 7 | payment code, retry, monitoring, replay, incidents, reconciliation, IP allowlist | met | `E3-P3A-SRC2` |

## B4 - P4 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P4 | Gateway/IPN separation | classification | mixed/under-documented | dedicated gateway/IPN doc | dedicated gateway/IPN doc | dedicated gateway/IPN guidance | met | `E4-P4A-IMPL1`, `E4-P4A-SRC1` |
| P4 | SDK/order lifecycle topics documented | topic groups | partial | 6 | 6 | install, payment methods, field order, order detail/status/cancel/void | met | `E4-P4A-SRC2` |

## B5 - P5 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P5 | QR optional/current parameter coverage | parameters | 6 current local params | 10 documented parameters | 10 documented parameters | includes missing official parameters from report | +4 | `E5-P5A-SRC1` |
| P5 | Project-specific best-practice examples retained as generic guidance | examples | many | 0 under `references/sepay` | 0 under `references/sepay` | 0 unframed project-specific examples | met | `E5-P5A-SRC3` |

## B6 - P6 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P6 | Generated skill output smoke targets | directories | pending | 2 repo dirs pass, 1 home dir stale | 2 repo dirs pass, 1 home dir stale | `.agents` and `.claude` payment-integration output checked if regenerated | repo generated output met; home install follow-up | `E6-P6A-GEN1`, `E6-P6A-GEN2`, `E6-P6A-GEN3` |
| P6 | Detect-changes changed files | files | pending | 12 | 12 | docs-only low-risk affected file set recorded | met | `E6-P6A-DETECT1` |
| P6 | Final graph inventory after analyze | graph nodes / relationships | 82421 / 120482 | 82463 / 120524 | 82463 / 120524 | graph refreshed after source changes and closure docs | +42 nodes / +42 relationships | `E6-P6A-GRAPH1`, `E9-PNC-GRAPH1` |

## Non-Benchmarkable Notes

- This is primarily a documentation/skill-source refresh. Runtime performance is not a goal.
- Build and test pass/fail belong in `evidence.md`; timings only become benchmark entries if the implementation intentionally changes build/test/runtime performance.
