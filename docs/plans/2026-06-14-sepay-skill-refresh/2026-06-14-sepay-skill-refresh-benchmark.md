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
| P0 | Existing SePay reference files under `references/sepay` | files | 6 | 6 | pending | 6 existing plus planned new gateway reference if accepted | pending | `rg --files internal\aicontext\skills\payment-integration\references\sepay`, E0-P0A-SRC1..E0-P0A-SRC6 |
| P0 | Planned new SePay reference files | files | 0 | 1 planned | pending | 1 `payment-gateway.md` if P1-A proceeds | pending | E0-P0A-REPORT1, E0-P0A-ROUTE1 |
| P0 | Direct routing files inspected | files | 2 | 2 | pending | 2 if new reference is added | pending | E0-P0A-ROUTE1, E0-P0A-FD7, E0-P0A-FD8 |
| P0 | Existing target markdown files with Anvien relationship count 0 | files | 6 | 6 | pending | preserve low graph-risk touch mode | pending | E0-P0A-FD1..E0-P0A-FD6 |
| P0 | Inspect-only helper scripts with high graph risk | files | 1 | 1 | pending | 0 edited unless user expands scope | pending | E0-P0A-SRC7, E0-P0A-FD9 |
| P0 | Current-state findings imported from problem report | findings | 16 | 16 | pending | all mapped to plan phases | pending | E0-P0A-REPORT1 |
| P0 | Official SePay source URLs listed in report | URLs | 20 | 20 | pending | re-check before implementation if stale | pending | E0-P0A-REPORT1 |

## B1 - P1 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P1 | SePay surface routes documented in overview | surfaces | pending | pending | pending | API v2, Webhooks, Payment Gateway/IPN, QR, Order VA, optional products | pending | pending |
| P1 | Dedicated gateway/IPN reference discoverability | route refs | pending | pending | pending | entrypoint and workflow mention `payment-gateway.md` if added | pending | pending |

## B2 - P2 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P2 | API default state | classification | legacy v1 default | pending | pending | API v2 default, v1 legacy | pending | pending |
| P2 | Order VA bank-specific rule coverage | rule groups | partial | pending | pending | BIDV, Sacombank, Vietcombank constraints represented | pending | pending |

## B3 - P3 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P3 | Webhook auth methods documented | methods | 3 | pending | pending | None, API Key, HMAC-SHA256, OAuth2 | pending | pending |
| P3 | Webhook operations topics documented | topic groups | partial | pending | pending | payment code, retry, monitoring, replay, incidents, reconciliation, IP allowlist | pending | pending |

## B4 - P4 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P4 | Gateway/IPN separation | classification | mixed/under-documented | pending | pending | dedicated gateway/IPN guidance | pending | pending |
| P4 | SDK/order lifecycle topics documented | topic groups | partial | pending | pending | install, payment methods, field order, order detail/status/cancel/void | pending | pending |

## B5 - P5 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P5 | QR optional/current parameter coverage | parameters | 6 current local params | pending | pending | includes missing official parameters from report | pending | pending |
| P5 | Project-specific best-practice examples retained as generic guidance | examples | many | pending | pending | 0 unframed project-specific examples | pending | pending |

## B6 - P6 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P6 | Generated skill output smoke targets | directories | pending | pending | pending | `.agents` and `.claude` payment-integration output checked if regenerated | pending | pending |

## Non-Benchmarkable Notes

- This is primarily a documentation/skill-source refresh. Runtime performance is not a goal.
- Build and test pass/fail belong in `evidence.md`; timings only become benchmark entries if the implementation intentionally changes build/test/runtime performance.
