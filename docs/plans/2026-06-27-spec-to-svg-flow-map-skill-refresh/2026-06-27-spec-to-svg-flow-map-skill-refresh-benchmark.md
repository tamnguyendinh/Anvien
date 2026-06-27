# Spec-to-SVG Flow Map Skill Refresh Benchmark Ledger

## Metadata

- Date: `2026-06-27`
- Plan: `docs/plans/2026-06-27-spec-to-svg-flow-map-skill-refresh/2026-06-27-spec-to-svg-flow-map-skill-refresh-plan.md`
- Evidence: `docs/plans/2026-06-27-spec-to-svg-flow-map-skill-refresh/2026-06-27-spec-to-svg-flow-map-skill-refresh-evidence.md`
- Benchmark: `docs/plans/2026-06-27-spec-to-svg-flow-map-skill-refresh/2026-06-27-spec-to-svg-flow-map-skill-refresh-benchmark.md`
- Actual status: `docs/plans/2026-06-27-spec-to-svg-flow-map-skill-refresh/2026-06-27-spec-to-svg-flow-map-skill-refresh-actual-status.md`

## Benchmark Rules

The benchmark file records measured product/runtime performance, capacity, package/startup size, graph/DB throughput, graph inventory counts, or generated-output inventory counts when the phase is benchmarkable. Build/test/e2e pass-fail belongs in evidence unless the timing, count, or size is the measured target.

## B0 - P0 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P0 | Anvien analyze scanned files | files | N/A | 1433 | 1433 | record graph inventory | N/A | `E0-P0A-ANV1` |
| P0 | Target file local relationships | relationships | N/A | 0 | 0 | record blast-radius inventory | N/A | `E0-P0A-FD1` |

## B1 - P1 Benchmarks

No product/runtime benchmark is expected for the markdown skill refresh. If a line-count or path inventory is measured during validation, record it here.

## B2 - Closure Benchmarks

Pending until closure.

## Non-Benchmarkable Notes

This plan changes skill documentation only. There is no product performance, DB throughput, package size, or UI runtime benchmark target.
