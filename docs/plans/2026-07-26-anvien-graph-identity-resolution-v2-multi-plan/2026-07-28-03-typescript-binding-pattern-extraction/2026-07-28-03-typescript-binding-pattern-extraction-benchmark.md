# Anvien TypeScript Binding-Pattern Extraction Benchmark Ledger

## Metadata

- Date: `2026-07-28`
- Plan: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-03-typescript-binding-pattern-extraction/2026-07-28-03-typescript-binding-pattern-extraction-plan.md`
- Evidence: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-03-typescript-binding-pattern-extraction/2026-07-28-03-typescript-binding-pattern-extraction-evidence.md`
- Benchmark: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-03-typescript-binding-pattern-extraction/2026-07-28-03-typescript-binding-pattern-extraction-benchmark.md`
- Actual status: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-03-typescript-binding-pattern-extraction/2026-07-28-03-typescript-binding-pattern-extraction-actual-status.md`
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

## B3 - P3 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P3 | recursive binding-pattern syntax cases | passed/total | inventory in P3-A | pending | pending | 100% | pending | `E3-P3A-TEST1` |
| P3 | declaration-context cases | passed/total | inventory in P3-B | pending | pending | 100% | pending | `E3-P3B-TEST1` |
| P3 | bounded target array bindings | represented/expected | 0/6 | pending | pending | 6/6 | pending | `E3-P3C2-TARGET1` |
| P3 | new target File-node omissions | count | 0 beyond quarantined 8 | pending | pending | 0 | pending | `E3-P3C2-TARGET1` |
| P3 | binding persistence native/fallback field parity | differing records/fields | no v2 binding projection | pending | pending | 0 per `S0`/`S1`/`S2` row | pending | `E3-P3C1-TEST1` |
| P3 | binding CLI parity | differing records/fields; orphan refs | no v2 CLI projection | pending | pending | 0 / 0 for complete `S3` row union | pending | `E3-P3C1A-TEST1` |
| P3 | binding MCP parity | differing records/fields; orphan refs | no v2 MCP projection | pending | pending | 0 / 0 for complete `S4` row union | pending | `E3-P3C1B-TEST1` |
| P3 | binding file-context-cache parity | differing records/fields; stale hits | no v2 `S7` projection | pending | pending | 0 / 0 | pending | `E3-P3C1C-TEST1` |
| P3 | binding HTTP parity | differing records/fields; orphan refs | no v2 HTTP projection | pending | pending | 0 / 0 for complete `S5` row union | pending | `E3-P3C1D-TEST1` |
| P3 | binding HTTP/MCP resource-cache parity | differing records/fields; stale hits | no v2 `S8` projection | pending | pending | 0 / 0 | pending | `E3-P3C1E-TEST1` |
| P3 | binding Web parity | differing records/fields; orphan refs | no v2 Web projection | pending | pending | 0 / 0 for complete `S6` row union | pending | `E3-P3C1F-PLAY1` |
| P3 | binding embedding parity | differing rows/fields; orphan refs | no v2 `S9` projection | pending | pending | 0 / 0 | pending | `E3-P3C1G-TEST1` |
| P3 | binding registry/group parity | differing records/fields; orphan refs | no v2 `S10` projection | pending | pending | 0 / 0 | pending | `E3-P3C1H-TEST1` |
| P3 | binding process/community parity | differing memberships/order; orphan refs | no v2 `S11` projection | pending | pending | 0 / 0 | pending | `E3-P3C1I-TEST1` |

## B3 Semantic Correction Metrics

| Metric | Unit | Baseline | Latest | Final | Target | Evidence |
|--------|------|----------|--------|-------|--------|----------|
| unsupported-pattern extraction diagnostic coverage | diagnosed/unsupported fixtures | 0/0 | pending | pending | 100% with structured code and count | `E2-PNC-BINDING1A` |
| declaration emission independent of type inference | emitted/eligible bindings | 0/0 | pending | pending | 100% | `E2-PNC-BINDING1B` |
| named `.map()` sites with ResolutionGap | sites / gaps | 6 / unknown | pending | pending | 6 / 0 | `E2-PNC-BINDING1C` |
| nested shadowing SymbolID correctness | correct/expected | 0/0 | pending | pending | 100% | `E2-PNC-BINDING1D` |
| import double-count delta | count delta | unknown | pending | pending | 0 | `E2-PNC-BINDING1E` |
| P3-B2A predecessor gate | accepted/required | 0/1 | pending | pending | 1/1 accepted and committed before P3-C/P3-C2 | `E3-P3C-B2A-GATE1` |
| P3 job ownership rows | accepted exact rows / 17 implementation jobs | 0/17 | pending | pending | 17/17 exact production/test/generated/fixture rows before job open | `E3-P3-OWNERSHIP1` |

## Non-Benchmarkable Notes

- P1-A contract ratification and P8 Supervisor/cleanup/closure are evidence gates, not benchmarkable product/runtime work.
- Build/test/e2e pass/fail belongs in the evidence ledger; only measured duration, count, size, throughput, latency, memory, parity, and accuracy ratios are recorded here.
- Local child closure is allowed after this child's local benchmark/evidence gates and qualified handoff pass. Pending P7 performance rows block only campaign/release closure owned by Child 07. Any budget exception still requires measured baseline/final values and explicit owner acceptance; it cannot be replaced by an unmeasured waiver.
