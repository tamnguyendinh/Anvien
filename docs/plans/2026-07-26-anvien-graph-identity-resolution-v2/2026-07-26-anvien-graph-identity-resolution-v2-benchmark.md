# Anvien Graph Identity and TypeScript Resolution Correctness v2 Benchmark Ledger

## Metadata

- Date: `2026-07-26`
- Plan: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/2026-07-26-anvien-graph-identity-resolution-v2-plan.md`
- Evidence: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/2026-07-26-anvien-graph-identity-resolution-v2-evidence.md`
- Benchmark: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/2026-07-26-anvien-graph-identity-resolution-v2-benchmark.md`
- Actual status: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/2026-07-26-anvien-graph-identity-resolution-v2-actual-status.md`

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

## B1 - P1 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P1 | shadow v2 duplicate Node IDs | count | pending first shadow | pending | pending | 0 | pending | `E1-P1E-BENCH1` |
| P1 | shadow v2 duplicate Relationship IDs | count | pending first shadow | pending | pending | 0 | pending | `E1-P1E-BENCH1` |
| P1 | shadow v2 missing endpoints/orphan refs | count | pending first shadow | pending | pending | 0 | pending | `E1-P1E-BENCH1` |
| P1 | shadow v2 Declaration occurrence conservation | input/output occurrences | pending first shadow | pending | pending | 100% and `0` silently dropped | pending | `E1-P1E-BENCH1` |
| P1 | shadow v2 source-site RelationshipID conservation | input/output source sites | pending first shadow | pending | pending | 100% and `0` provenance loss | pending | `E1-P1E-BENCH1` |
| P1 | target v2-shadow `time`/`now` oracle coverage | occurrences | v1 known `2/4` | pending | pending | `4/4` in-memory v2 shadow only | pending | `E1-P1E-TARGET1`, `E1-P1E-ORACLE1` |
| P1 | five-run canonical Node-ID set equality | equal runs | pending first shadow | pending | pending | 5/5 | pending | `E1-P1E-BENCH1` |
| P1 | five-run canonical Relationship-ID set equality | equal runs | pending first shadow | pending | pending | 5/5 | pending | `E1-P1E-BENCH1` |
| P1 | Declaration-node expansion | nodes | 0 dedicated Declaration nodes | pending | pending | record and approve before cutover | pending | `E1-P1E-BENCH1` |
| P1 | graph JSON size | bytes | capture pre-change in P1-E | pending | pending | record and approve before cutover | pending | `E1-P1E-BENCH1` |
| P1 | peak analyze RSS | bytes | five-run pre-cutover baseline required | pending | pending | final <= baseline +15% | pending | `E1-P1E-BENCH1` |

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

## B4 - P4 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P4 | export syntax cases | passed/total | inventory in P4-B | pending | pending | 100% | pending | `E4-P4B-TEST1` |
| P4 | bounded target direct exports | correct/expected | 0/21 | pending | pending | 21/21 | pending | `E4-P4C2-TARGET1` |
| P4 | bounded export projection surfaces agreeing | agreeing/total | 0/12 exact surfaces | pending | pending | 12/12 (`S0`–`S11`) | pending | `E4-P4C2-TARGET1` |
| P4 | access-visibility/export conflations | count | current contract conflates/omits | pending | pending | 0 | pending | `E4-P4C2-TARGET1` |
| P4 | export persistence native/fallback field parity | differing records/fields | no v2 export projection | pending | pending | 0 per `S0`/`S1`/`S2` row | pending | `E4-P4C1-TEST1` |
| P4 | export CLI parity | differing records/fields; orphan refs | no v2 CLI projection | pending | pending | 0 / 0 for complete `S3` row union | pending | `E4-P4C1A-TEST1` |
| P4 | export MCP parity | differing records/fields; orphan refs | no v2 MCP projection | pending | pending | 0 / 0 for complete `S4` row union | pending | `E4-P4C1B-TEST1` |
| P4 | export file-context-cache parity | differing records/fields; stale hits | no v2 `S7` projection | pending | pending | 0 / 0 | pending | `E4-P4C1C-TEST1` |
| P4 | export HTTP parity | differing records/fields; orphan refs | no v2 HTTP projection | pending | pending | 0 / 0 for complete `S5` row union | pending | `E4-P4C1D-TEST1` |
| P4 | export HTTP/MCP resource-cache parity | differing records/fields; stale hits | no v2 `S8` projection | pending | pending | 0 / 0 | pending | `E4-P4C1E-TEST1` |
| P4 | export Web parity | differing records/fields; orphan refs | no v2 Web projection | pending | pending | 0 / 0 for complete `S6` row union | pending | `E4-P4C1F-PLAY1` |
| P4 | export embedding parity | differing rows/fields; orphan refs | no v2 `S9` projection | pending | pending | 0 / 0 | pending | `E4-P4C1G-TEST1` |
| P4 | export registry/group parity | differing records/fields; orphan refs | no v2 `S10` projection | pending | pending | 0 / 0 | pending | `E4-P4C1H-TEST1` |
| P4 | export process/community parity | differing memberships/order; orphan refs | no v2 `S11` projection | pending | pending | 0 / 0 | pending | `E4-P4C1I-TEST1` |

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
