# Anvien Ambient/External Declaration Universe And Truthful Diagnostics Benchmark Ledger

## Metadata

- Date: `2026-07-28`
- Plan: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-plan.md`
- Evidence: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-evidence.md`
- Benchmark: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-benchmark.md`
- Actual status: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-actual-status.md`
- Source benchmark: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/2026-07-26-anvien-graph-identity-resolution-v2-benchmark.md`
- Multi-plan roadmap: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-26-anvien-graph-identity-resolution-v2-roadmap.md`

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

Benchmark sections follow plan phases:

- `B0` corresponds to `P0`.
- `B1` corresponds to `P1`.
- `B2` corresponds to `P2`.
- Use item-level IDs such as `B-P1-A` when a checklist item needs separate benchmark evidence.
- Create a benchmark section only when the matching phase has benchmarkable measurements.
- Do not invent fixed metric categories; record the measurements required by the matching plan phase.

For performance release gates, capture at least five comparable runs before v2 cutover and at least five final runs. Use median for analyze/DB-load and p95 for query latency. Baseline and final must use the same commit-bound corpus, configuration, build mode, machine, and warm/cold-cache policy.

## B0 - P0 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P0 | Anvien self graph Files | count | 1,505 | 1,505 | pending | preserve/measure final | 0 | `E0-P0A-GRAPH1`, `E0-P0A-GRAPH6` |
| P0 | Anvien self parsed code | count | 676 | 676 | pending | preserve/measure final | 0 | `E0-P0A-GRAPH1`, `E0-P0A-GRAPH6` |
| P0 | Anvien self failed/unsupported/unknown | count | 0 | 0 | pending | 0 | 0 | `E0-P0A-GRAPH1`, `E0-P0A-GRAPH6` |
| P0 | Anvien self graph Nodes | count | 84,558 | 84,540 | pending | measured identity-v2 expansion | -18 | `E0-P0A-GRAPH1`, `E0-P0A-GRAPH6` |
| P0 | Anvien self graph Relationships | count | 123,398 | 123,380 | pending | measured identity-v2 expansion | -18 | `E0-P0A-GRAPH1`, `E0-P0A-GRAPH6` |
| P0 | reader-matrix seed anchors | count | 195 | 195 | pending | seed only; P2-A must prove complete source scan with 0 unlisted readers | 0 | `E0-P0A-MATRIX4` |
| P0 | accepted target graph Files | count | 1,359 | 1,359 | pending | no new omissions | 0 | `E0-P0A-BOUNDARY1` |
| P0 | accepted target graph Nodes | count | 84,807 | 84,807 | pending | measured identity-v2 expansion | 0 | `E0-P0A-BOUNDARY1` |
| P0 | accepted target graph Relationships | count | 114,125 | 114,125 | pending | measured identity-v2 expansion | 0 | `E0-P0A-BOUNDARY1` |
| P0 | accepted target graph SHA-256 | hash | `DD3A3A987E95C91787E7C70EF68BA5D2F953D24A32D35DEFE38B8E8E606FF090` | same | pending | generation-recorded final hash | 0 | `E0-P0A-BOUNDARY1` |
| P0 | accepted target analyze duration (historical context only) | seconds | 42.1059032 | 42.1059032 | pending | replace with fresh five-run P2-G baseline | 0 | `E0-P0A-HISTORY2` |
| P0 | target eligible TS/JS paths | count | 895 | 895 | pending | preserve input inventory | 0 | `E0-P0A-REPORT1` |
| P0 | target TS/JS File nodes | count | 887 | 887 | pending | 887 unless separate scanner plan | 0 | `E0-P0A-REPORT1`, scanner excluded |
| P0 | target scanner omissions | count | 8 | 8 | pending | exactly same 8; 0 new omissions | 0 | `E0-P0A-BOUNDARY2` |
| P0 | eligible Scope Boundary blocks with explicit template fields | ratio | 47/99 | 99/99 | pending | 99/99 | +52 slices | `E0-P0A-TEMPLATE4` |
| P0 | eligible Acceptance blocks with explicit template fields | ratio | 49/99 | 99/99 | pending | 99/99 | +50 slices | `E0-P0A-TEMPLATE4` |
| P0 | bounded array bindings represented | ratio | 0/6 | 0/6 | pending | 6/6 | 0 | `E0-P0A-REPORT1` |
| P0 | bounded same-name declarations represented distinctly | ratio | 2/4 | 2/4 | pending | 4/4 | 0 | `E0-P0A-REPORT1` |
| P0 | bounded direct exports with correct export metadata | ratio | 0/21 | 0/21 | pending | 21/21 | 0 | `E0-P0A-REPORT1` |
| P0 | bounded barrel calls resolved to terminal ResolutionTarget | ratio | 0/2 | 0/2 | pending | 2/2 | 0 | `E0-P0A-REPORT1` |
| P0 | bounded ambient sites correctly resolved/classified | ratio | 0/3 | 0/3 | pending | 3/3 | 0 | `E0-P0A-REPORT1` |

## B6 - P6 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P6 | stdlib catalog uncompressed size | bytes | establish in P6-B | pending | pending | `<= 33,554,432` bytes (`32 MiB`) | pending | `E6-P6B-BENCH1` |
| P6 | packaged binary size delta | bytes/% | capture immediately before P6-B | pending | pending | measured and owner-approved | pending | `E6-P6B-BENCH1` |
| P6 | catalog cold load | milliseconds | establish in P6-B | pending | pending | `<= 250 ms` on the bound benchmark machine | pending | `E6-P6B-BENCH1` |
| P6 | catalog warm lookup p95 | milliseconds | establish in P6-B | pending | pending | final query p95 <= pre-cutover +10% | pending | `E6-P6B-BENCH1` |
| P6 | declaration status-vector coverage | passed/total codes | no structured status matrix | pending | pending | 100% positive/negative codes | pending | `E6-P6C3-STATUS1` |
| P6 | package declaration file ceiling (`50,000`) | files / over-limit cases | no bounded declaration loader | pending | pending | in-budget `100%`; explicit over-limit `budget_exceeded` | pending | `E6-P6C2-BENCH1` |
| P6 | package declaration byte ceiling (`256 MiB`) | bytes / over-limit cases | no bounded declaration loader | pending | pending | in-budget `100%`; explicit over-limit `budget_exceeded` | pending | `E6-P6C2-BENCH1` |
| P6 | per-file declaration ceiling (`8 MiB`) | bytes / over-limit cases | no bounded declaration loader | pending | pending | in-budget `100%`; explicit over-limit `budget_exceeded` | pending | `E6-P6C2-BENCH1` |
| P6 | declaration reference-depth ceiling (`64`) | max reference depth / over-limit cases | no bounded declaration loader | pending | pending | depths `<=64` resolve exactly; depth 65 returns explicit `budget_exceeded` with no partial result | pending | `E6-P6C2-BENCH1` |
| P6 | package cache ceiling (`256 MiB`) | bytes / over-limit cases | no bounded package cache | pending | pending | in-budget `100%`; explicit over-limit `budget_exceeded` | pending | `E6-P6C2-BENCH1` |
| P6 | referenced external-node materialization | nodes / descriptors | no external materializer | pending | pending | exact equality; zero catalog File nodes | pending | `E6-P6C2-BENCH1` |
| P6 | bounded ambient sites | correct/expected | 0/3 | pending | pending | 3/3 | pending | `E6-P6D-TARGET1` |
| P6 | source sites both resolved and unresolved | count | inventory in P6-C1 | pending | pending | 0 | pending | `E6-P6D-TARGET1` |

## Non-Benchmarkable Notes

- P1-A contract ratification and P8 Supervisor/cleanup/closure are evidence gates, not benchmarkable product/runtime work.
- Build/test/e2e pass/fail belongs in the evidence ledger; only measured duration, count, size, throughput, latency, memory, parity, and accuracy ratios are recorded here.
- A final plan cannot close while any required P7 performance row remains pending. A budget exception must contain measured baseline/final values and explicit owner acceptance; it cannot be replaced by an unmeasured waiver.
