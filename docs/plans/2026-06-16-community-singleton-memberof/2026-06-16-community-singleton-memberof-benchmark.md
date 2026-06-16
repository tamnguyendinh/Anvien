# Community Singleton MEMBER_OF Benchmark Ledger

## Metadata

- Date: `2026-06-16`
- Plan: `docs/plans/2026-06-16-community-singleton-memberof/2026-06-16-community-singleton-memberof-plan.md`
- Evidence: `docs/plans/2026-06-16-community-singleton-memberof/2026-06-16-community-singleton-memberof-evidence.md`
- Benchmark: `docs/plans/2026-06-16-community-singleton-memberof/2026-06-16-community-singleton-memberof-benchmark.md`
- Actual status: `docs/plans/2026-06-16-community-singleton-memberof/2026-06-16-community-singleton-memberof-actual-status.md`

## B0 - P0 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P0 | Skipped DB relationships on failing HEAD | relationships | 506 | 506 | 0 | 0 | -506 | `E0-P0A-REPRO1`, `E0-P0A-DIAG1`, `E2-P2A-ANALYZE2` |
| P0 | Missing community target relationships | relationships | 506 | 506 | 0 | 0 | -506 | `E0-P0A-DIAG1`, `E2-P2A-ANALYZE2` |

## B1 - P1 Benchmarks

No separate P1 benchmark; P1 changed emission behavior and is covered by P2 graph/load counts.

## B2 - P2 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P2 | Global analyze node rows | rows | N/A | 272114 | 272114 | load succeeds | N/A | `E2-P2A-ANALYZE2` |
| P2 | Global analyze relationship rows | rows | N/A | 321335 | 321335 | load succeeds | N/A | `E2-P2A-ANALYZE2` |
| P2 | Global analyze duration | seconds | N/A | 158.875 | 158.875 | informational | N/A | `E2-P2A-ANALYZE2` |
| P2 | Communities emitted | communities | N/A | 1675 | 1675 | informational | N/A | `E2-P2A-ANALYZE2` |
| P2 | Memberships emitted | memberships | N/A | 5948 | 5948 | no dangling singleton memberships | N/A | `E2-P2A-ANALYZE2` |

## Non-Benchmarkable Notes

Targeted unit test pass/fail and build pass/fail will be recorded in evidence, not benchmark rows.
