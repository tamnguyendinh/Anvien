# Anvien Graph Identity Resolution v2 Multi-Plan Execution Benchmark Ledger

## Metadata

- Date: `2026-08-09`
- Plan: `docs/plans/2026-08-09-anvien-graph-identity-resolution-v2-multi-plan-execution/2026-08-09-anvien-graph-identity-resolution-v2-multi-plan-execution-plan.md`
- Evidence: `docs/plans/2026-08-09-anvien-graph-identity-resolution-v2-multi-plan-execution/2026-08-09-anvien-graph-identity-resolution-v2-multi-plan-execution-evidence.md`
- Benchmark: `docs/plans/2026-08-09-anvien-graph-identity-resolution-v2-multi-plan-execution/2026-08-09-anvien-graph-identity-resolution-v2-multi-plan-execution-benchmark.md`
- Actual status: `docs/plans/2026-08-09-anvien-graph-identity-resolution-v2-multi-plan-execution/2026-08-09-anvien-graph-identity-resolution-v2-multi-plan-execution-actual-status.md`

## Benchmark Rules

Record measured values only. Validation pass/fail belongs in the evidence ledger. Keep baselines and deltas tied to exact commands and evidence IDs.

## B0 - P0 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P0 | scanned files | files | 1,558 | 1,564 | pending | no unexplained loss | +6 docs | E0-P0A-GRAPH1, E0-P0A-GRAPH2 |
| P0 | parsed code files | files | 676 | 676 | pending | no unexplained loss | 0 | E0-P0A-GRAPH1 |
| P0 | graph nodes | nodes | 85,203 | 85,282 | pending | record final | +79 docs nodes | E0-P0A-GRAPH1, E0-P0A-GRAPH2 |
| P0 | graph relationships | relationships | 124,071 | 124,150 | pending | record final | +79 docs relationships | E0-P0A-GRAPH1, E0-P0A-GRAPH2 |

## B1 - Child 01 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P1 | identity collision count | conflicts | pending | pending | pending | 0 | pending | pending |
| P1 | shadow declaration conservation | ratio | pending | pending | pending | 100% | pending | pending |

## B2 - Child 02 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P2 | reader parity rows | rows | pending | pending | pending | 100% | pending | pending |

## B3 - Child 03 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P3 | binding oracle coverage | bindings | pending | pending | pending | 6/6 | pending | pending |

## B4 - Child 04 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P4 | direct export facts | facts | pending | pending | pending | target-specific | pending | pending |

## B5 - Child 05 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P5 | terminal barrel resolution | cases | pending | pending | pending | target-specific | pending | pending |

## B6 - Child 06 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P6 | external capability classification | cases | pending | pending | pending | exact/structural/degraded | pending | pending |

## B7 - Child 07 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P7 | analyze duration | seconds | pending | pending | pending | record | pending | pending |
| P7 | peak RSS | bytes | pending | pending | pending | record | pending | pending |

## Non-Benchmarkable Notes

No additional benchmarkable value is claimed until the matching child slice has produced fresh runtime evidence.
