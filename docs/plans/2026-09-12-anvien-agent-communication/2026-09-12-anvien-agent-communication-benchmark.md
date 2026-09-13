# Anvien Agent Communication Benchmark Ledger

## Metadata

- Date: `2026-09-12`
- Plan: `docs/plans/2026-09-12-anvien-agent-communication/2026-09-12-anvien-agent-communication-plan.md`
- Rules: `docs/plans/2026-09-12-anvien-agent-communication/2026-09-12-anvien-agent-communication-plan.md#rules`
- Evidence: `docs/plans/2026-09-12-anvien-agent-communication/2026-09-12-anvien-agent-communication-evidence.md`
- Benchmark: `docs/plans/2026-09-12-anvien-agent-communication/2026-09-12-anvien-agent-communication-benchmark.md`
- Actual status: `docs/plans/2026-09-12-anvien-agent-communication/2026-09-12-anvien-agent-communication-actual-status.md`

## Benchmark Rules

The benchmark file records measurements.

It should contain:

- metadata and companion files;
- benchmark rules;
- benchmark sections such as `B0`, `B1`, or sections by phase/task;
- metric tables with unit, baseline, latest, final, target, and delta when needed;
- inventory count;
- runtime or performance metric;
- graph, coverage, or accuracy metric;
- package, bundle, file size, or hash metric;
- before/after numbers;
- UI, layout, or browser metric when the plan involves UI;
- command-surface or generated-output inventory when the plan involves generated documentation.

Benchmark records measured numbers only. Do not put command logs, design decisions, or validation narrative here. Build/test/e2e pass-fail belongs in evidence unless the timing, count, or size is the measured target.

Benchmark sections must follow the plan phases:

- `B0` corresponds to `P0`.
- `B1` corresponds to `P1`.
- `B2` corresponds to `P2`.
- Use item-level IDs such as `B-P1-A` when a checklist item needs separate benchmark evidence.
- Create a benchmark section only when the matching phase has benchmarkable measurements.
- Do not invent fixed metric categories; record the measurements required by the matching plan phase.

## B0 - P0 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P0 | Semantic graph nodes | nodes | 170912 | 170912 | Chưa đo | Không đặt trong khảo sát | N/A | E0-P0A-GRAPH1 |

## B1 - P1 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| Chưa triển khai | Chưa đo | — | — | — | — | — | — | — |

## B2 - P2 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| Chưa triển khai | Chưa đo | — | — | — | — | — | — | — |

## Non-Benchmarkable Notes

Use this section only for phases that are not benchmarkable. Do not invent metrics.

Chưa có benchmark Coordinator hoặc provider. Các hàng P1/P2 là chưa đo, không phải baseline mới.
