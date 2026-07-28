# Versioned Persistence and Identity v2 Cutover Benchmark Ledger

## Metadata

- Date: `2026-07-28`
- Plan: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-02-versioned-persistence-and-v2-cutover/2026-07-28-02-versioned-persistence-and-v2-cutover-plan.md`
- Evidence: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-02-versioned-persistence-and-v2-cutover/2026-07-28-02-versioned-persistence-and-v2-cutover-evidence.md`
- Benchmark: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-02-versioned-persistence-and-v2-cutover/2026-07-28-02-versioned-persistence-and-v2-cutover-benchmark.md`
- Actual status: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-02-versioned-persistence-and-v2-cutover/2026-07-28-02-versioned-persistence-and-v2-cutover-actual-status.md`
- Roadmap: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-26-anvien-graph-identity-resolution-v2-roadmap.md`
- Source phase: legacy `P2`

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
| P0 | Source implementation slices | slices | 42 | 42 | — | 42 preserved | 0 | `E0-P0A-SOURCE1` |
| P0 | reader inventory classification | ratio | seed 195 anchors | seed 195 anchors | pending | 100% classified / 0 unassigned / 0 unlisted | 0 | `E0-P0A-STATUS1` |
| P0 | Target eligible TS/JS paths | files | 895 | 895 | pending | preserve inventory | 0 | `E0-P0A-SCANNER1` |
| P0 | Target TS/JS File nodes | files | 887 | 887 | pending | 887 unless separate scanner plan | 0 | `E0-P0A-SCANNER1` |
| P0 | Quarantined scanner omissions | files | 8 | 8 | pending | exactly 8 / 0 new | 0 | `E0-P0A-SCANNER1` |

## B1 - P1 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P1 | reader inventory classification | classified/total, unassigned, unlisted | seed in `2026-07-26-anvien-graph-identity-resolution-v2::E0-P0A-MATRIX3` | pending | pending | 100% classified / 0 unassigned / 0 unlisted | pending | `E1-P1A1-MATRIXREVIEW1` |
| P1 | guarded reader-row execution | passed/total and unlisted count | frozen P1-A1 matrix | pending | pending | 100% passed / 0 unlisted | pending | `E1-P1E2-MATRIX1` |
| P1 | old-binary/client rejection matrix | passed/total | inventory in P1-A1 | pending | pending | 100% with `INDEX_VERSION_MISMATCH` | pending | `E1-P1E2-MATRIX1`, `E1-P1A8-PLAY1` |
| P1 | protocol handshake body-open violations | count | no v2 handshake implementation | pending | pending | 0 | pending | `E1-P1A-TEST1`, `E1-P1E2-MATRIX1` |
| P1 | JSON/Ladybug canonical node-field parity | differing records/fields | pending v2 store | pending | pending | 0 | pending | `E1-P1B-TEST1`, `E1-P1B2-TEST1`, `E1-P1E2-S0BASE1`, `E1-P1E2-S1BASE1` |
| P1 | JSON/native/fallback Cypher canonical relationship-field parity | differing records/fields | pending v2 store | pending | pending | 0 | pending | `E1-P1B3-TEST1`, `E1-P1B4-TEST1`, `E1-P1E2-S2BASE1` |
| P1 | `S0` canonical baseline | differing records/orphan refs | pending source manifest | pending | pending | 0 / 0 | pending | `E1-P1E2-S0BASE1` |
| P1 | `S1` canonical baseline | differing records/orphan refs | pending source manifest | pending | pending | 0 / 0 | pending | `E1-P1E2-S1BASE1` |
| P1 | `S2` canonical baseline | differing records/orphan refs | pending source manifest | pending | pending | 0 / 0 | pending | `E1-P1E2-S2BASE1` |
| P1 | `S3` canonical baseline | differing records/orphan refs | pending source manifest | pending | pending | 0 / 0 | pending | `E1-P1E2-S3BASE1` |
| P1 | `S4` canonical baseline | differing records/orphan refs | pending source manifest | pending | pending | 0 / 0 | pending | `E1-P1E2-S4BASE1` |
| P1 | `S5` canonical baseline | differing records/orphan refs | pending source manifest | pending | pending | 0 / 0 | pending | `E1-P1E2-S5BASE1` |
| P1 | `S6` canonical baseline | differing records/orphan refs | pending source manifest | pending | pending | 0 / 0 | pending | `E1-P1E2-S6BASE1` |
| P1 | `S7` canonical baseline | differing records/orphan refs | pending source manifest | pending | pending | 0 / 0 | pending | `E1-P1E2-S7BASE1` |
| P1 | `S8` canonical baseline | differing records/orphan refs | pending source manifest | pending | pending | 0 / 0 | pending | `E1-P1E2-S8BASE1` |
| P1 | `S9` canonical baseline | differing records/orphan refs | pending source manifest | pending | pending | 0 / 0 | pending | `E1-P1E2-S9BASE1` |
| P1 | `S10` canonical baseline | differing records/orphan refs | pending source manifest | pending | pending | 0 / 0 | pending | `E1-P1E2-S10BASE1` |
| P1 | `S11` canonical baseline | differing records/orphan refs | pending source manifest | pending | pending | 0 / 0 | pending | `E1-P1E2-S11BASE1` |
| P1 | skipped duplicate nodes/relationships during load | count | current behavior can hide duplicates | pending | pending | 0 | pending | `E1-P1B2-TEST1` |
| P1 | orphan embeddings/cache/group/registry refs | count | inventories in P1-C3/C5/C6, P1-D, and P1-F2/F3/F4 | pending | pending | 0 | pending | `E1-P1E2-S7BASE1`, `E1-P1E2-S8BASE1`, `E1-P1E2-S9BASE1`, `E1-P1E2-S10BASE1` |
| P1 | global repo registry generation consistency | differing records | pending repo epoch runtime | pending | pending | 0 | pending | `E1-P1F3-TEST1`, `E1-P1F6-FAULT1` |
| P1 | group snapshot member-generation-vector consistency | differing records/vector entries | pending group epoch runtime | pending | pending | 0 | pending | `E1-P1F4-TEST1`, `E1-P1F6-FAULT1` |
| P1 | cache/embedding generation-key isolation | stale/mixed hits | pending epoch runtime | pending | pending | 0 | pending | `E1-P1F2-TEST1`, `E1-P1F6-FAULT1` |
| P1 | fault boundaries retaining old generation | passed/total | inventory in P1-F6 | pending | pending | 100% | pending | `E1-P1F6-FAULT1`, `E1-P1F6-RECOVERY1` |
| P1 | mixed-generation requests | count | pending generation runtime | pending | pending | 0 | pending | `E1-P1G-RUNTIME1` |
| P1 | deterministic active-v2 analyzes | equal runs | pending cutover | pending | pending | 5/5 | pending | `E1-P1G-RUNTIME1` |
| P1 | pre-cutover analyze median | seconds | five identical v1 runs required before cutover | pending | pending | recorded; final <= baseline +10% | pending | `E1-P1G-PREBASE1` |
| P1 | pre-cutover Ladybug-load median | seconds | five identical v1 runs required before cutover | pending | pending | recorded; final <= baseline +10% | pending | `E1-P1G-PREBASE1` |
| P1 | pre-cutover native-Cypher query p95 | milliseconds | five identical v1 runs required before cutover | pending | pending | recorded; final <= baseline +10% | pending | `E1-P1G-PREBASE1` |
| P1 | pre-cutover fallback-Cypher query p95 | milliseconds | five identical v1 runs required before cutover | pending | pending | recorded; final <= baseline +10% | pending | `E1-P1G-PREBASE1` |
| P1 | pre-cutover peak RSS | bytes | five identical v1 runs required before cutover | pending | pending | recorded; final <= baseline +15% | pending | `E1-P1G-PREBASE1` |
| P1 | graph size at pre-cutover baseline | bytes | five identical v1 runs required before cutover | pending | pending | recorded | pending | `E1-P1G-PREBASE1` |
| P1 | staged-v2 candidate analyze median | seconds | `E1-P1G-PREBASE1` v1 analyze median | pending | pending | <= baseline +10% before active CAS | pending | `E1-P1G-CANDIDATE1` |
| P1 | staged-v2 candidate Ladybug-load median | seconds | `E1-P1G-PREBASE1` v1 Ladybug-load median | pending | pending | <= baseline +10% before active CAS | pending | `E1-P1G-CANDIDATE1` |
| P1 | staged-v2 candidate native-Cypher query p95 | milliseconds | `E1-P1G-PREBASE1` v1 native-Cypher p95 | pending | pending | <= baseline +10% before active CAS | pending | `E1-P1G-CANDIDATE1` |
| P1 | staged-v2 candidate fallback-Cypher query p95 | milliseconds | `E1-P1G-PREBASE1` v1 fallback-Cypher p95 | pending | pending | <= baseline +10% before active CAS | pending | `E1-P1G-CANDIDATE1` |
| P1 | staged-v2 candidate peak RSS | bytes | `E1-P1G-PREBASE1` v1 peak RSS | pending | pending | <= baseline +15% before active CAS | pending | `E1-P1G-CANDIDATE1` |
| P1 | staged-v2 candidate graph size | bytes | `E1-P1G-PREBASE1` v1 graph size | pending | pending | recorded before active CAS; owner-approved exception required if release budget is set and exceeded | pending | `E1-P1G-CANDIDATE1` |
| P1 | fault-injected registry/group/cache publication boundaries | passed/total | inventory in P1-F6 | pending | pending | 100% retain prior epoch | pending | `E1-P1F6-FAULT1`, `E1-P1F6-RECOVERY1` |
| P1 | concurrent two-repo group CAS conflicts | passed/total interleavings | no group-vector protocol | pending | pending | 100% conflict fail-closed / prior snapshot queryable | pending | `E1-P1F4-TEST1`, `E1-P1F6-FAULT1` |

## Non-Benchmarkable Notes

Pn Supervisor judgment, cleanup reasoning, roadmap handoff, and commit authorization are evidence rather than measurements. Do not invent runtime values before the corresponding local slice runs.
