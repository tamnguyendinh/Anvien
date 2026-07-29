# Anvien Versioned Persistence, Opaque Consumers, Atomic Generation, And V2 Cutover Benchmark Ledger

## Metadata

- Date: `2026-07-28`
- Plan: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-02-versioned-persistence-and-v2-cutover/2026-07-28-02-versioned-persistence-and-v2-cutover-plan.md`
- Evidence: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-02-versioned-persistence-and-v2-cutover/2026-07-28-02-versioned-persistence-and-v2-cutover-evidence.md`
- Benchmark: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-02-versioned-persistence-and-v2-cutover/2026-07-28-02-versioned-persistence-and-v2-cutover-benchmark.md`
- Actual status: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-02-versioned-persistence-and-v2-cutover/2026-07-28-02-versioned-persistence-and-v2-cutover-actual-status.md`
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

## B2 - P2 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P2 | reader inventory classification | classified/total, unassigned, unlisted | seed in `E0-P0A-MATRIX3` | pending | pending | 100% classified / 0 unassigned / 0 unlisted | pending | `E2-P2A1-MATRIXREVIEW1` |
| P2 | guarded reader-row execution | passed/total and unlisted count | frozen P2-A1 matrix | pending | pending | 100% passed / 0 unlisted | pending | `E2-P2E2-MATRIX1` |
| P2 | old-binary/client rejection matrix | passed/total | inventory in P2-A1 | pending | pending | 100% with `INDEX_VERSION_MISMATCH` | pending | `E2-P2E2-MATRIX1`, `E2-P2A8-PLAY1` |
| P2 | protocol handshake body-open violations | count | no v2 handshake implementation | pending | pending | 0 | pending | `E2-P2A-TEST1`, `E2-P2E2-MATRIX1` |
| P2 | JSON/Ladybug canonical node-field parity | differing records/fields | pending v2 store | pending | pending | 0 | pending | `E2-P2B-TEST1`, `E2-P2B2-TEST1`, `E2-P2E2-S0BASE1`, `E2-P2E2-S1BASE1` |
| P2 | JSON/native/fallback Cypher canonical relationship-field parity | differing records/fields | pending v2 store | pending | pending | 0 | pending | `E2-P2B3-TEST1`, `E2-P2B4-TEST1`, `E2-P2E2-S2BASE1` |
| P2 | `S0` canonical baseline | differing records/orphan refs | pending source manifest | pending | pending | 0 / 0 | pending | `E2-P2E2-S0BASE1` |
| P2 | `S1` canonical baseline | differing records/orphan refs | pending source manifest | pending | pending | 0 / 0 | pending | `E2-P2E2-S1BASE1` |
| P2 | `S2` canonical baseline | differing records/orphan refs | pending source manifest | pending | pending | 0 / 0 | pending | `E2-P2E2-S2BASE1` |
| P2 | `S3` canonical baseline | differing records/orphan refs | pending source manifest | pending | pending | 0 / 0 | pending | `E2-P2E2-S3BASE1` |
| P2 | `S4` canonical baseline | differing records/orphan refs | pending source manifest | pending | pending | 0 / 0 | pending | `E2-P2E2-S4BASE1` |
| P2 | `S5` canonical baseline | differing records/orphan refs | pending source manifest | pending | pending | 0 / 0 | pending | `E2-P2E2-S5BASE1` |
| P2 | `S6` canonical baseline | differing records/orphan refs | pending source manifest | pending | pending | 0 / 0 | pending | `E2-P2E2-S6BASE1` |
| P2 | `S7` canonical baseline | differing records/orphan refs | pending source manifest | pending | pending | 0 / 0 | pending | `E2-P2E2-S7BASE1` |
| P2 | `S8` canonical baseline | differing records/orphan refs | pending source manifest | pending | pending | 0 / 0 | pending | `E2-P2E2-S8BASE1` |
| P2 | `S9` canonical baseline | differing records/orphan refs | pending source manifest | pending | pending | 0 / 0 | pending | `E2-P2E2-S9BASE1` |
| P2 | `S10` canonical baseline | differing records/orphan refs | pending source manifest | pending | pending | 0 / 0 | pending | `E2-P2E2-S10BASE1` |
| P2 | `S11` canonical baseline | differing records/orphan refs | pending source manifest | pending | pending | 0 / 0 | pending | `E2-P2E2-S11BASE1` |
| P2 | skipped duplicate nodes/relationships during load | count | current behavior can hide duplicates | pending | pending | 0 | pending | `E2-P2B2-TEST1` |
| P2 | orphan embeddings/cache/group/registry refs | count | inventories in P2-C3/C5/C6, P2-D, and P2-F2/F3/F4 | pending | pending | 0 | pending | `E2-P2E2-S7BASE1`, `E2-P2E2-S8BASE1`, `E2-P2E2-S9BASE1`, `E2-P2E2-S10BASE1` |
| P2 | global repo registry generation consistency | differing records | pending repo epoch runtime | pending | pending | 0 | pending | `E2-P2F3-TEST1`, `E2-P2F6-FAULT1` |
| P2 | group snapshot member-generation-vector consistency | differing records/vector entries | pending group epoch runtime | pending | pending | 0 | pending | `E2-P2F4-TEST1`, `E2-P2F6-FAULT1` |
| P2 | cache/embedding generation-key isolation | stale/mixed hits | pending epoch runtime | pending | pending | 0 | pending | `E2-P2F2-TEST1`, `E2-P2F6-FAULT1` |
| P2 | fault boundaries retaining old generation | passed/total | inventory in P2-F6 | pending | pending | 100% | pending | `E2-P2F6-FAULT1`, `E2-P2F6-RECOVERY1` |
| P2 | mixed-generation requests | count | pending generation runtime | pending | pending | 0 | pending | `E2-P2G-RUNTIME1` |
| P2 | deterministic active-v2 analyzes | equal runs | pending cutover | pending | pending | 5/5 | pending | `E2-P2G-RUNTIME1` |
| P2 | pre-cutover analyze median | seconds | five identical v1 runs required before cutover | pending | pending | recorded; final <= baseline +10% | pending | `E2-P2G-PREBASE1` |
| P2 | pre-cutover Ladybug-load median | seconds | five identical v1 runs required before cutover | pending | pending | recorded; final <= baseline +10% | pending | `E2-P2G-PREBASE1` |
| P2 | pre-cutover native-Cypher query p95 | milliseconds | five identical v1 runs required before cutover | pending | pending | recorded; final <= baseline +10% | pending | `E2-P2G-PREBASE1` |
| P2 | pre-cutover fallback-Cypher query p95 | milliseconds | five identical v1 runs required before cutover | pending | pending | recorded; final <= baseline +10% | pending | `E2-P2G-PREBASE1` |
| P2 | pre-cutover peak RSS | bytes | five identical v1 runs required before cutover | pending | pending | recorded; final <= baseline +15% | pending | `E2-P2G-PREBASE1` |
| P2 | graph size at pre-cutover baseline | bytes | five identical v1 runs required before cutover | pending | pending | recorded | pending | `E2-P2G-PREBASE1` |
| P2 | staged-v2 candidate analyze median | seconds | `E2-P2G-PREBASE1` v1 analyze median | pending | pending | <= baseline +10% before active CAS | pending | `E2-P2G-CANDIDATE1` |
| P2 | staged-v2 candidate Ladybug-load median | seconds | `E2-P2G-PREBASE1` v1 Ladybug-load median | pending | pending | <= baseline +10% before active CAS | pending | `E2-P2G-CANDIDATE1` |
| P2 | staged-v2 candidate native-Cypher query p95 | milliseconds | `E2-P2G-PREBASE1` v1 native-Cypher p95 | pending | pending | <= baseline +10% before active CAS | pending | `E2-P2G-CANDIDATE1` |
| P2 | staged-v2 candidate fallback-Cypher query p95 | milliseconds | `E2-P2G-PREBASE1` v1 fallback-Cypher p95 | pending | pending | <= baseline +10% before active CAS | pending | `E2-P2G-CANDIDATE1` |
| P2 | staged-v2 candidate peak RSS | bytes | `E2-P2G-PREBASE1` v1 peak RSS | pending | pending | <= baseline +15% before active CAS | pending | `E2-P2G-CANDIDATE1` |
| P2 | staged-v2 candidate graph size | bytes | `E2-P2G-PREBASE1` v1 graph size | pending | pending | recorded before active CAS; owner-approved exception required if release budget is set and exceeded | pending | `E2-P2G-CANDIDATE1` |
| P2 | fault-injected registry/group/cache publication boundaries | passed/total | inventory in P2-F6 | pending | pending | 100% retain prior epoch | pending | `E2-P2F6-FAULT1`, `E2-P2F6-RECOVERY1` |
| P2 | concurrent two-repo group CAS conflicts | passed/total interleavings | no group-vector protocol | pending | pending | 100% conflict fail-closed / prior snapshot queryable | pending | `E2-P2F4-TEST1`, `E2-P2F6-FAULT1` |

## B2 Semantic Correction Metrics

| Metric | Unit | Baseline | Latest | Final | Target | Evidence |
|--------|------|----------|--------|-------|--------|----------|
| manifest/handshake required metadata coverage | fields present / required | 6/6 legacy fields | pending | pending | 15/15 manifest fields and 10/10 request fields, including all nine semantic-correction fields, exact spelling, and fingerprints | `E2-P2A-MANIFEST1`, `E2-P2A-HANDSHAKE1`, `E2-P2A-METADATA1` |
| missing-fingerprint fail-closed cases | passed/total | 0/0 | pending | pending | 100% before body open | `E2-P2A-HANDSHAKE1`, `E2-P2A-METADATA1` |
| persisted manifest metadata fields | present/required | 6/15 | pending | pending | 15/15, with typed non-empty values | `E2-P2A-MANIFEST1`, `E2-P2A-METADATA1` |
| handshake request metadata fields | present/required | 3/10 | pending | pending | 10/10, including analyzer, column/position encoding support, and both fingerprints | `E2-P2A-HANDSHAKE1`, `E2-P2A-METADATA1` |
| invalid/absent metadata body-open attempts | blocked/attempts | 0/0 | pending | pending | 100% blocked before every S0-S11 body/stream/cache open | `E2-P2A-HANDSHAKE1` |
| P2 job ownership rows | accepted exact rows / 42 jobs | 0/42 | pending | pending | 41/41 implementation/validation rows with exact production/test/generated/fixture paths plus 1/1 documentation-only P2-A1 matrix-owner row with reasoned N/A fields | `E2-P2A-OWNERSHIP1` |

## Non-Benchmarkable Notes

- P1-A contract ratification and P8 Supervisor/cleanup/closure are evidence gates, not benchmarkable product/runtime work.
- Build/test/e2e pass/fail belongs in the evidence ledger; only measured duration, count, size, throughput, latency, memory, parity, and accuracy ratios are recorded here.
- Local child closure is allowed after this child's local benchmark/evidence gates and qualified handoff pass. Pending P7 performance rows block only campaign/release closure owned by Child 07. Any budget exception still requires measured baseline/final values and explicit owner acceptance; it cannot be replaced by an unmeasured waiver.
