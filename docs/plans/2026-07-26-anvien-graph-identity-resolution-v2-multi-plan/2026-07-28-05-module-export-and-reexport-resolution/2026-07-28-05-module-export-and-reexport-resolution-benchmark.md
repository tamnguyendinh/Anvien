# Anvien Module Export Tables And Barrel/Re-Export Resolution Benchmark Ledger

## Metadata

- Date: `2026-07-28`
- Plan: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-05-module-export-and-reexport-resolution/2026-07-28-05-module-export-and-reexport-resolution-plan.md`
- Evidence: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-05-module-export-and-reexport-resolution/2026-07-28-05-module-export-and-reexport-resolution-evidence.md`
- Benchmark: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-05-module-export-and-reexport-resolution/2026-07-28-05-module-export-and-reexport-resolution-benchmark.md`
- Actual status: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-05-module-export-and-reexport-resolution/2026-07-28-05-module-export-and-reexport-resolution-actual-status.md`
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

## B5 - P5 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P5 | export-table star-edge ceiling (`2,000` per module) | max edges / over-limit cases | no bounded export-table implementation | pending | pending | in-budget `100%`; edge 2,001 returns explicit `budget_exceeded` with no partial table | pending | `E5-P5B-BENCH1` |
| P5 | re-export hop ceiling (`64`) | max hops / over-limit cases | no bounded traversal implementation | pending | pending | in-budget `100%`; explicit over-limit `budget_exceeded` | pending | `E5-P5C-BENCH1` |
| P5 | star-edge ceiling (`2,000` per module) | max edges / over-limit cases | no bounded traversal implementation | pending | pending | in-budget `100%`; explicit over-limit `budget_exceeded` | pending | `E5-P5C-BENCH1` |
| P5 | candidate ceiling (`10,000` per lookup) | max candidates / over-limit cases | no bounded traversal implementation | pending | pending | in-budget `100%`; explicit over-limit `budget_exceeded` | pending | `E5-P5C-BENCH1` |
| P5 | SCC ceiling (`2,000` modules) | max SCC / over-limit cases | no bounded traversal implementation | pending | pending | in-budget `100%`; explicit over-limit `budget_exceeded` | pending | `E5-P5C-BENCH1` |
| P5 | semantic vector manifest coverage | passed/total | inventory in P5-C vector manifest | pending | pending | 100% exact target/status/proof | pending | `E5-P5C-VECTOR1`, `E5-P5C-TEST1` |
| P5 | SCC/cycle cases | passed/total | inventory in P5-C | pending | pending | 100% | pending | `E5-P5C-TEST1` |
| P5 | bounded barrel calls | resolved/expected | 0/2 | pending | pending | 2/2 | pending | `E5-P5D-TARGET1` |
| P5 | corresponding false barrel gaps | count | 4 bounded gap rows | pending | pending | 0 | pending | `E5-P5D-PARITY1` |

## B5 Semantic Correction Metrics

| Metric | Unit | Baseline | Latest | Final | Target | Evidence |
|--------|------|----------|--------|-------|--------|----------|
| zero-physical-declaration barrel export surface | exposed/expected entries | 0/0 | pending | pending | exact non-empty syntax-derived surface | `E2-PNC-MODULE1D` |
| physical path-resolution count delta after export traversal | count delta | 0 | pending | pending | 0 | `E2-PNC-MODULE1E` |
| syntactic `IMPORTS` count delta after export traversal | count delta | 0 | pending | pending | 0 | `E2-PNC-MODULE1E` |
| directExportedDefinitionCount / resolvedExportEntryCount / publicApiSymbolCount ownership | separate metrics | conflated | pending | pending | three separately reported owners | `E2-PNC-MODULE1A`, `E2-PNC-MODULE1B`, `E2-PNC-MODULE1C` |
| directExportedDefinitionCount consumed from P4 | count | 0 | pending | pending | exact unchanged P4 count | `E2-PNC-MODULE1A` |
| resolvedExportEntryCount | count | 0 | pending | pending | exact terminal entries, P5-C owner | `E2-PNC-MODULE1B` |
| reachableThroughBarrel | reachable entries / expected | 0/0 | pending | pending | exact proof-backed reachability, P5-C owner | `E2-PNC-MODULE1B` |
| publicApiSymbolCount | count | 0 | pending | pending | exact terminal set, P5-D owner | `E2-PNC-MODULE1C` |
| zero-physical barrel physical declaration count | count | unknown | pending | pending | exactly 0 with `resolvedExportEntryCount > 0` | `E2-PNC-MODULE1D` |
| zero-physical barrel terminal proof hops | complete/expected | 0/0 | pending | pending | 100% complete proof for every exposed entry | `E2-PNC-MODULE1D` |
| physical path-resolution count before export traversal | count | unknown | pending | pending | record absolute baseline | `E2-PNC-MODULE1E` |
| physical path-resolution count after export traversal | count | unknown | pending | pending | equal to before count | `E2-PNC-MODULE1E` |
| syntactic `IMPORTS` count before export traversal | count | unknown | pending | pending | record absolute baseline | `E2-PNC-MODULE1E` |
| syntactic `IMPORTS` count after export traversal | count | unknown | pending | pending | equal to before count | `E2-PNC-MODULE1E` |
| P5 job ownership rows | exact rows / P5 jobs | 0/4 | pending | pending | 100% exact production/test/generated/fixture rows, with one write-owner for each shared coordinator | `E5-P5-DERIVED1` |

## Non-Benchmarkable Notes

- P1-A contract ratification and P8 Supervisor/cleanup/closure are evidence gates, not benchmarkable product/runtime work.
- Build/test/e2e pass/fail belongs in the evidence ledger; only measured duration, count, size, throughput, latency, memory, parity, and accuracy ratios are recorded here.
- Local child closure is allowed after this child's local benchmark/evidence gates and qualified handoff pass. Pending P7 performance rows block only campaign/release closure owned by Child 07. Any budget exception still requires measured baseline/final values and explicit owner acceptance; it cannot be replaced by an unmeasured waiver.
