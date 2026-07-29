# Anvien Cross-Surface Acceptance, Target Validation, And Performance Benchmark Ledger

## Metadata

- Date: `2026-07-28`
- Plan: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-07-cross-surface-acceptance-and-target-validation/2026-07-28-07-cross-surface-acceptance-and-target-validation-plan.md`
- Evidence: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-07-cross-surface-acceptance-and-target-validation/2026-07-28-07-cross-surface-acceptance-and-target-validation-evidence.md`
- Benchmark: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-07-cross-surface-acceptance-and-target-validation/2026-07-28-07-cross-surface-acceptance-and-target-validation-benchmark.md`
- Actual status: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-07-cross-surface-acceptance-and-target-validation/2026-07-28-07-cross-surface-acceptance-and-target-validation-actual-status.md`
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

## B7 - P7 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P7 | final canonical Node-ID set determinism | equal runs | P1 shadow 5-run baseline | pending | pending | 5/5 | pending | `E7-P7A-DETERMINISM1` |
| P7 | final canonical Relationship-ID set determinism | equal runs | P1 shadow 5-run baseline | pending | pending | 5/5 | pending | `E7-P7A-DETERMINISM1` |
| P7 | duplicate IDs | count | P1 shadow baseline | pending | pending | 0 | pending | `E7-P7A-DETERMINISM1` |
| P7 | missing endpoints/orphan refs | count | P1 shadow baseline | pending | pending | 0 | pending | `E7-P7A-DETERMINISM1` |
| P7 | version/fault matrix | passed/total | P2 inventory | pending | pending | 100% | pending | `E7-P7A-VERSION1`, `E7-P7A-FAULT1` |
| P7 | final target bounded defects | passed/total | 0/5 defect families | pending | pending | 5/5 | pending | `E7-P7B-TARGET1` |
| P7 | final target new scanner omissions | count | 0 beyond exact 8 | pending | pending | 0 | pending | `E7-P7B-BOUNDARY1` |
| P7 | final analyze median regression | percent | five-run pre-cutover median | pending | pending | <= +10% | pending | `E7-P7C-BENCH1` |
| P7 | final Ladybug load median regression | percent | five-run pre-cutover median | pending | pending | <= +10% | pending | `E7-P7C-BENCH1` |
| P7 | final native-Cypher query p95 regression | percent | five-run pre-cutover native p95 | pending | pending | <= +10% | pending | `E7-P7C-NATIVEBENCH1` |
| P7 | final fallback-Cypher query p95 regression | percent | five-run pre-cutover fallback p95 | pending | pending | <= +10% | pending | `E7-P7C-FALLBACKBENCH1` |
| P7 | final peak RSS regression | percent | five-run pre-cutover peak | pending | pending | <= +15% | pending | `E7-P7C-BENCH1` |
| P7 | `S0` Graph JSON canonical parity | differing records/fields; orphan refs | P2 parity baseline | pending | pending | 0 / 0 | pending | `E7-P7C-S0-PARITY1` |
| P7 | `S1` Ladybug native Cypher canonical parity | differing records/fields; orphan refs | P2 parity baseline | pending | pending | 0 / 0 | pending | `E7-P7C-S1-PARITY1` |
| P7 | `S2` fallback Cypher canonical parity | differing records/fields; orphan refs | P2 parity baseline | pending | pending | 0 / 0 | pending | `E7-P7C-S2-PARITY1` |
| P7 | `S3` CLI canonical parity | differing records/fields; orphan refs | P2 parity baseline | pending | pending | 0 / 0 | pending | `E7-P7C-S3-PARITY1` |
| P7 | `S4` MCP canonical parity | differing records/fields; orphan refs | P2 parity baseline | pending | pending | 0 / 0 | pending | `E7-P7C-S4-PARITY1` |
| P7 | `S5` HTTP canonical parity | differing records/fields; orphan refs | P2 parity baseline | pending | pending | 0 / 0 | pending | `E7-P7C-S5-PARITY1` |
| P7 | `S6` Web canonical parity | differing records/fields; orphan refs | P2 parity baseline | pending | pending | 0 / 0 for complete `S6` row union | pending | `E7-P7C-S6-PARITY1` |
| P7 | `S7` file-context cache generation/field parity | differing records/fields; stale hits | P2 cache baseline | pending | pending | 0 / 0 | pending | `E7-P7C-S7-PARITY1` |
| P7 | `S8` HTTP/MCP cache generation/field parity | differing records/fields; stale hits | P2 cache baseline | pending | pending | 0 / 0 | pending | `E7-P7C-S8-PARITY1` |
| P7 | `S9` embedding node/generation/dimension/hash parity | differing rows/fields; orphan refs | P2 embedding baseline | pending | pending | 0 / 0 | pending | `E7-P7C-S9-PARITY1` |
| P7 | `S10` registry/group contract/epoch parity | differing records/fields; orphan refs | P2 registry baseline | pending | pending | 0 / 0 | pending | `E7-P7C-S10-PARITY1` |
| P7 | `S11` process/community source-anchored parity | differing memberships/order; orphan refs | P2 derived baseline | pending | pending | 0 / 0 | pending | `E7-P7C-S11-PARITY1` |

## Non-Benchmarkable Notes

- P1-A contract ratification and P8 Supervisor/cleanup/closure are evidence gates, not benchmarkable product/runtime work.
- Build/test/e2e pass/fail belongs in the evidence ledger; only measured duration, count, size, throughput, latency, memory, parity, and accuracy ratios are recorded here.
- A final plan cannot close while any required P7 performance row remains pending. A budget exception must contain measured baseline/final values and explicit owner acceptance; it cannot be replaced by an unmeasured waiver.

## B7 Semantic Correction Metrics

| Metric | Unit | Baseline | Latest | Final | Target | Evidence |
|--------|------|----------|--------|-------|--------|----------|
| default external isolation across context/impact/rename/process/groups | leaked records | unknown | pending | pending | 0 | `E2-PNC-FINALORACLE1A`, `E2-PNC-FINALORACLE1B`, `E2-PNC-FINALORACLE1C`, `E2-PNC-FINALORACLE1D`, `E2-PNC-FINALORACLE1E` |
| explicit `include_external` opt-in result | included/expected | 0/0 | pending | pending | exact opt-in set | `E2-PNC-FINALORACLE1A`, `E2-PNC-FINALORACLE1B`, `E2-PNC-FINALORACLE1C`, `E2-PNC-FINALORACLE1D`, `E2-PNC-FINALORACLE1E`, `E2-PNC-FINALORACLE1F` |
| zero-physical-declaration barrel surface | exposed/expected entries | 0/0 | pending | pending | exact syntax-derived surface | `E2-PNC-FINALORACLE1G` |
| physical path-resolution and `IMPORTS` count delta | count delta | 0 | pending | pending | 0 | `E2-PNC-FINALORACLE1H` |
| Promise/Math invalid intrinsic outcomes | count | unknown | pending | pending | 0 | `E2-PNC-FINALORACLE1I` |
| context default external leakage | records | unknown | pending | pending | 0 | `E2-PNC-FINALORACLE1A` |
| context explicit opt-in set | included/expected | 0/0 | pending | pending | exact | `E2-PNC-FINALORACLE1A` |
| impact default external leakage | records | unknown | pending | pending | 0 | `E2-PNC-FINALORACLE1B` |
| impact explicit opt-in set | included/expected | 0/0 | pending | pending | exact | `E2-PNC-FINALORACLE1B` |
| rename default external candidates/edits | candidates/edits | unknown | pending | pending | 0 edits and no candidates by default | `E2-PNC-FINALORACLE1C` |
| rename explicit opt-in candidates | included/expected | 0/0 | pending | pending | exact, non-editable unless separately authorized | `E2-PNC-FINALORACLE1C` |
| process default external members | members | unknown | pending | pending | 0 | `E2-PNC-FINALORACLE1D` |
| process explicit opt-in members | included/expected | 0/0 | pending | pending | exact | `E2-PNC-FINALORACLE1D` |
| groups default external members/contracts | members/contracts | unknown | pending | pending | 0 | `E2-PNC-FINALORACLE1E` |
| groups explicit opt-in members/contracts | included/expected | 0/0 | pending | pending | exact | `E2-PNC-FINALORACLE1E` |
| option propagation consistency | surfaces with same typed option / 5 | 0/5 | pending | pending | 5/5 | `E2-PNC-FINALORACLE1F` |
| graph-generation capability descriptors | descriptors / graph generations | 0/1 | pending | pending | exactly 1/1 with accepted mode/confidence/completeness and generation/config/catalog binding | `E2-PNC-FINALORACLE1K` |
| capability descriptor/outcome consistency | non-stronger consistent outcomes / outcomes | 0/0 | pending | pending | 100% | `E2-PNC-FINALORACLE1K` |
| capability descriptor S0-S11 differences | differing descriptors / applicable surfaces | unknown | pending | pending | 0 and no mixed-generation references | `E2-PNC-FINALORACLE1K` |
| zero-barrel physical declarations | count | unknown | pending | pending | 0 | `E2-PNC-FINALORACLE1G` |
| zero-barrel resolved export entries | count | 0 | pending | pending | >0 and exact expected set | `E2-PNC-FINALORACLE1G` |
| physical path-resolution count before/after | before,after | unknown/unknown | pending | pending | equal absolute values | `E2-PNC-FINALORACLE1H` |
| syntactic IMPORTS count before/after | before,after | unknown/unknown | pending | pending | equal absolute values | `E2-PNC-FINALORACLE1H` |
| Promise target outcome | allowed/observed target sites | 0/1 | pending | pending | `resolved_external` (target denominator 1/1) | `E2-PNC-FINALORACLE1I` |
| Math.max target outcome | allowed/observed target sites | 0/1 | pending | pending | `resolved_external` (target denominator 1/1) | `E2-PNC-FINALORACLE1I` |
| Math.min target outcome | allowed/observed target sites | 0/1 | pending | pending | `resolved_external` (target denominator 1/1) | `E2-PNC-FINALORACLE1I` |
| negative external-capability fixture outcomes | allowed/observed negative sites | 0/0 | pending | pending | explicit external-capability failure only; never part of target 3/3 | `E2-PNC-FINALORACLE1J` |
| P7 job ownership rows | exact rows / P7 jobs | 0/3 | pending | pending | 100% exact production/test/generated/fixture/evidence rows | `E7-P7-OWNERSHIP1` |

Child 07 is the campaign/release closure owner. Its P7 performance rows are mandatory for terminal closure; no earlier child is held open by these rows.
