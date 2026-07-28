# Graph Identity Contract and Strict Graph Construction Benchmark Ledger

## Metadata

- Date: `2026-07-28`
- Plan: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-plan.md`
- Evidence: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-evidence.md`
- Benchmark: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-benchmark.md`
- Actual status: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-actual-status.md`
- Roadmap: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/2026-07-26-anvien-graph-identity-resolution-v2-roadmap.md`
- Source phase: legacy `P1`

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
| P0 | Source implementation slices | slices | 11 | 11 | — | 11 preserved | 0 | `E0-P0A-SOURCE1` |
| P0 | bounded same-name declarations represented distinctly | ratio | 2/4 | 2/4 | pending | 4/4 | 0 | `E0-P0A-STATUS1` |
| P0 | Target eligible TS/JS paths | files | 895 | 895 | pending | preserve inventory | 0 | `E0-P0A-SCANNER1` |
| P0 | Target TS/JS File nodes | files | 887 | 887 | pending | 887 unless separate scanner plan | 0 | `E0-P0A-SCANNER1` |
| P0 | Quarantined scanner omissions | files | 8 | 8 | pending | exactly 8 / 0 new | 0 | `E0-P0A-SCANNER1` |

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

## Non-Benchmarkable Notes

Pn Supervisor judgment, cleanup reasoning, roadmap handoff, and commit authorization are evidence rather than measurements. Do not invent runtime values before the corresponding local slice runs.
