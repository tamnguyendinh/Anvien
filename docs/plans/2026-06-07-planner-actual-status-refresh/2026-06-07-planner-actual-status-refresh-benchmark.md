# Planner Actual Status Refresh Benchmark Ledger

## Metadata

- Date: `2026-06-07`
- Plan: `docs/plans/2026-06-07-planner-actual-status-refresh/2026-06-07-planner-actual-status-refresh-plan.md`
- Evidence: `docs/plans/2026-06-07-planner-actual-status-refresh/2026-06-07-planner-actual-status-refresh-evidence.md`
- Benchmark: `docs/plans/2026-06-07-planner-actual-status-refresh/2026-06-07-planner-actual-status-refresh-benchmark.md`
- Actual status: `docs/plans/2026-06-07-planner-actual-status-refresh/2026-06-07-planner-actual-status-refresh-actual-status.md`

## Benchmark Rules

Benchmark records measured numbers only. Build/test pass-fail belongs in evidence unless timing, count, or size is the measured target.

## B0 - P0 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P0 | Analyze scanned files | files | 1384 | 1384 | 1384 | record | 0 | E0.1 |
| P0 | Analyze graph nodes | nodes | 82343 | 82343 | 82343 | record | 0 | E0.1 |
| P0 | Analyze graph relationships | relationships | 120404 | 120404 | 120404 | record | 0 | E0.1 |
| P0 | Planner SKILL file relationships | relationships | 0 | 0 | 0 | low-risk target | 0 | E0.2 |
| P0 | Actual-status template file relationships | relationships | 0 | 0 | 0 | low-risk target | 0 | E0.3 |
| P0 | Plan template file relationships | relationships | 0 | 0 | 0 | low-risk target | 0 | E0.7 |

## B1 - P1 Benchmarks

No benchmarkable runtime or capacity metric is expected for this documentation/template edit. Record validation evidence in the evidence ledger.

## Non-Benchmarkable Notes

P1 changes are procedural markdown instructions and a markdown template update. No product runtime, package size, UI layout, database, API throughput, or graph-builder behavior is changed.
