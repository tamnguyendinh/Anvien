# Anvien Graph Identity Resolution v2 Multi-Plan Authoring Benchmark Ledger

## Metadata

- Date: `2026-07-28`
- Plan: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/2026-07-28-00-multi-plan-authoring/2026-07-28-00-multi-plan-authoring-plan.md`
- Evidence: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/2026-07-28-00-multi-plan-authoring/2026-07-28-00-multi-plan-authoring-evidence.md`
- Benchmark: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/2026-07-28-00-multi-plan-authoring/2026-07-28-00-multi-plan-authoring-benchmark.md`
- Actual status: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/2026-07-28-00-multi-plan-authoring/2026-07-28-00-multi-plan-authoring-actual-status.md`

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
- `B3` corresponds to `P3`.
- Use item-level IDs such as `B-P1-A` when a checklist item needs separate benchmark evidence.
- Create a benchmark section only when the matching phase has benchmarkable measurements.
- Do not invent fixed metric categories; record the measurements required by the matching plan phase.

## B0 - P0 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P0 | Legacy campaign artifacts | files | 5 | 5 | 5 | 5 preserved | 0 | `E0-P0A-SRC2` |
| P0 | Legacy plan size | bytes | 567,584 | 567,584 | 567,584 | preserve until cutover marker | 0 | `E0-P0A-SRC1` |
| P0 | Legacy plan total lines | lines | 5,467 | 5,467 | 5,467 | source baseline | 0 | `E0-P0A-SRC1` |
| P0 | Legacy plan nonblank lines | lines | 5,291 | 5,291 | 5,291 | source baseline | 0 | `E0-P0A-SRC1` |
| P0 | Legacy implementation phases | phases | 7 | 7 | 7 | 7 child owners | 0 | `E0-P0A-STRUCT1` |
| P0 | Legacy closure phases | phases | 1 | 1 | 1 | distribute to child Pn | 0 | `E0-P0A-STRUCT1` |
| P0 | Legacy P1 slices | slices | 11 | 11 | 11 | 11 mapped | 0 | `E0-P0A-SLICE1` |
| P0 | Legacy P2 slices | slices | 42 | 42 | 42 | 42 mapped | 0 | `E0-P0A-SLICE2` |
| P0 | Legacy P3 slices | slices | 17 | 17 | 17 | 17 mapped | 0 | `E0-P0A-SLICE3` |
| P0 | Legacy P4 slices | slices | 15 | 15 | 15 | 15 mapped | 0 | `E0-P0A-SLICE4` |
| P0 | Legacy P5 slices | slices | 4 | 4 | 4 | 4 mapped | 0 | `E0-P0A-SLICE5` |
| P0 | Legacy P6 slices | slices | 6 | 6 | 6 | 6 mapped | 0 | `E0-P0A-SLICE6` |
| P0 | Legacy P7 slices | slices | 3 | 3 | 3 | 3 mapped | 0 | `E0-P0A-SLICE7` |
| P0 | Total legacy implementation slices | slices | 98 | 98 | 98 | 98 mapped exactly once | 0 | `E0-P0A-TOTAL1` |
| P0 | Legacy closure items | items | 3 | 3 | 3 | 3 closure roles per child | 0 | `E0-P0A-TOTAL1` |
| P0 | Existing campaign roadmaps | files | 0 | 0 | — | 1 | 0 | `E0-P0A-SRC2` |
| P0 | Existing numbered implementation children | plan sets | 0 | 0 | — | 7 | 0 | `E0-P0A-SRC2` |
| P0 | Existing standard child files | files | 0 | 0 | — | 28 | 0 | `E0-P0A-SRC2` |
| P0 | Existing mapped destination slices | slices | 0 | 0 | — | 98 | 0 | `E0-P0A-SRC2` |
| P0 | Fresh Anvien graph files | files | 1,506 | 1,506 | — | informational | 0 | `E0-P0A-GRAPH1` |
| P0 | Fresh Anvien graph nodes | nodes | 84,548 | 84,548 | — | informational | 0 | `E0-P0A-GRAPH1` |
| P0 | Fresh Anvien graph relationships | relationships | 123,388 | 123,388 | — | informational | 0 | `E0-P0A-GRAPH1` |

## B1 - P1 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P1-A | Frozen source snapshots | snapshots | 0 | 1 | 1 | 1 | +1 | `E1-P1A-SNAPSHOT1` |
| P1-A | Source phase ownership rows | rows | 0 | 7 | 7 | 7 | +7 | `E1-P1A-MAP1` |
| P1-A | Source slice crosswalk rows | rows | 0 | 98 | 98 | 98 | +98 | `E1-P1A-MAP1` |
| P1-A | Unmapped source slices | slices | 98 | 0 | 0 | 0 | -98 | `E1-P1A-MAP1` |
| P1-A | Duplicate source mappings | mappings | 0 | 0 | 0 | 0 | 0 | `E1-P1A-MAP2` |
| P1-B | Campaign roadmaps | files | 0 | 0 | — | 1 | 0 | `E1-P1B-ROADMAP1` |
| P1-B | Roadmap child records | rows | 0 | 0 | — | 7 | 0 | `E1-P1B-INVENTORY1` |
| P1-B | Roadmap planned standard child files | files | 0 | 0 | — | 28 | 0 | `E1-P1B-INVENTORY1` |
| P1-B | Broken roadmap links | links | 0 | 0 | — | 0 | 0 | `E1-P1B-LINK1` |

## B2 - P2 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P2-A | Child 01 standard files | files | 0 | 0 | — | 4 | 0 | `E2-P2A-FILES1` |
| P2-A | Child 01 mapped slices | slices | 0 | 0 | — | 11 | 0 | `E2-P2A-MAP1` |
| P2-B | Child 02 standard files | files | 0 | 0 | — | 4 | 0 | `E2-P2B-FILES1` |
| P2-B | Child 02 mapped slices | slices | 0 | 0 | — | 42 | 0 | `E2-P2B-MAP1` |
| P2-B | `index-reader-matrix.md` mutation owners | owners | 0 assigned | 0 assigned | — | 1 | 0 | `E2-P2B-MATRIX1` |
| P2-C | Child 03 standard files | files | 0 | 0 | — | 4 | 0 | `E2-P2C-FILES1` |
| P2-C | Child 03 mapped slices | slices | 0 | 0 | — | 17 | 0 | `E2-P2C-MAP1` |
| P2-D | Child 04 standard files | files | 0 | 0 | — | 4 | 0 | `E2-P2D-FILES1` |
| P2-D | Child 04 mapped slices | slices | 0 | 0 | — | 15 | 0 | `E2-P2D-MAP1` |
| P2-E | Child 05 standard files | files | 0 | 0 | — | 4 | 0 | `E2-P2E-FILES1` |
| P2-E | Child 05 mapped slices | slices | 0 | 0 | — | 4 | 0 | `E2-P2E-MAP1` |
| P2-F | Child 06 standard files | files | 0 | 0 | — | 4 | 0 | `E2-P2F-FILES1` |
| P2-F | Child 06 mapped slices | slices | 0 | 0 | — | 6 | 0 | `E2-P2F-MAP1` |
| P2-G | Child 07 standard files | files | 0 | 0 | — | 4 | 0 | `E2-P2G-FILES1` |
| P2-G | Child 07 mapped slices | slices | 0 | 0 | — | 3 | 0 | `E2-P2G-MAP1` |
| P2 | Complete child plan sets | plan sets | 0 | 0 | — | 7 | 0 | `E2-P2G-FILES1` |
| P2 | Standard child files | files | 0 | 0 | — | 28 | 0 | `E2-P2G-FILES1` |
| P2 | Child P0 lifecycle sections | sections | 0 | 0 | — | 7 | 0 | `E2-P2G-STRUCT1` |
| P2 | Child local implementation phases | phases | 0 | 0 | — | 7 | 0 | `E2-P2G-STRUCT1` |
| P2 | Child Pn-A/Pn-B/Pn-C sets | closure sets | 0 | 0 | — | 7 | 0 | `E2-P2G-CLOSURE1` |
| P2 | Cumulative mapped destination slices | slices | 0 | 0 | — | 98 | 0 | `E2-P2G-MAP1` |

## B3 - P3 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P3-A | Candidate campaign output files | files | 0 | 0 | — | 29 | 0 | `E3-P3A-STRUCT1` |
| P3-A | Source slices checked | slices | 0 | 0 | — | 98 | 0 | `E3-P3A-MAP1` |
| P3-A | Destination slices checked | slices | 0 | 0 | — | 98 | 0 | `E3-P3A-MAP1` |
| P3-A | Missing source mappings | slices | 98 | 98 | — | 0 | 0 | `E3-P3A-MAP1` |
| P3-A | Duplicate source mappings | mappings | 0 | 0 | — | 0 | 0 | `E3-P3A-MAP1` |
| P3-A | Extra destination mappings | slices | 0 | 0 | — | 0 | 0 | `E3-P3A-MAP1` |
| P3-A | Required-field failures | failures | 0 | 0 | — | 0 | 0 | `E3-P3A-FIELDS1` |
| P3-A | Broken links | links | 0 | 0 | — | 0 | 0 | `E3-P3A-LINK1` |
| P3-A | Mutable artifacts with multiple owners | artifacts | 0 | 0 | — | 0 | 0 | `E3-P3A-OWNER1` |
| P3-B | Unconditional Supervisor PASS verdicts | verdicts | 0 | 0 | — | 1 | 0 | `E3-P3B-SUPERVISOR1` |
| P3-B | Active campaign authorities | authorities | 1 | 1 | — | 1 | 0 | `E3-P3B-CUTOVER1` |
| P3-B | Deleted legacy plan artifacts | files | 0 | 0 | — | 0 | 0 | `E3-P3B-CUTOVER1` |
| P3-B | Production/test/runtime files changed | files | 0 | 0 | — | 0 | 0 | `E3-P3B-DIFF1` |
| P3-B | Target-repository files changed | files | 0 | 0 | — | 0 | 0 | `E3-P3B-DIFF1` |
| P3-B | Repository-root plan/temp artifacts created | artifacts | 0 | 0 | — | 0 | 0 | `E3-P3B-DIFF1` |

## Non-Benchmarkable Notes

P0-P3 document inventories, hashes, counts, and mapping error counts are benchmarkable and are recorded above. Pn Supervisor judgment, dead-work reasoning, and commit authorization are evidence rather than measurements. Build, Docker, browser, Playwright, runtime latency, and product performance are not benchmarkable for this plan because its authorized scope is documentation-only; no value may be fabricated for those surfaces.
