# Anvien Current Graph Persistence and Reader Consistency Benchmark Ledger

## Metadata

- Date: `2026-07-28`
- Plan: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-02-current-graph-persistence-and-reader-consistency/2026-07-28-02-current-graph-persistence-and-reader-consistency-plan.md`
- Evidence: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-02-current-graph-persistence-and-reader-consistency/2026-07-28-02-current-graph-persistence-and-reader-consistency-evidence.md`
- Benchmark: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-02-current-graph-persistence-and-reader-consistency/2026-07-28-02-current-graph-persistence-and-reader-consistency-benchmark.md`
- Actual status: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-02-current-graph-persistence-and-reader-consistency/2026-07-28-02-current-graph-persistence-and-reader-consistency-actual-status.md`
- Roadmap: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-26-anvien-graph-accuracy-roadmap.md`

## Benchmark Rules

- Record only persistence, affected-reader, repeated-analyze, and directly related capacity measurements.
- P2-A establishes the affected denominator from current source; no historical reader count is a baseline.
- Compare corrected fields and records, not only aggregate node/relationship counts.
- Use matched source, configuration, analyzer build, machine, and cache policy for repeated/runtime measurements.
- Build/test/QA pass-fail and Supervisor verdicts belong in evidence.
- Record analyze duration, Ladybug load/query duration, graph size, or memory only when Child 02 changes or measures that boundary.
- Do not copy identity target, binding, export, barrel, ambient, or campaign-wide metrics here.

## B0 - P0 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P0 | selected bounded raw/projection facts retaining one-to-one cardinality | facts with one-to-one cardinality / selected facts | `7/7` | `7/7` bounded observation | pending Child 02 comparison | preserve corrected records; field parity measured separately | pending | `E0-P0A-VERIFY1` |
| P0 | source-proven affected readers | readers | unknown before P2-A | unknown | pending | exact discovered count, zero unassigned | pending | `E2-P2A-INVENTORY1`, `E2-P2A-MATRIX1` |
| P0 | Child 02 implementation slices | slices | 5 planned | 5 planned | pending | 5 accepted | pending | plan checklist |

## B2 - P2 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P2 | affected persistence paths | paths | unknown before source inventory | pending | pending | exact count, zero unassigned | pending | `E2-P2A-INVENTORY1` |
| P2 | affected readers | readers | unknown before source inventory | pending | pending | exact count, zero unassigned | pending | `E2-P2A-MATRIX1` |
| P2 | Graph JSON/Ladybug corrected records with field differences | records | pending accepted Child 01 fields | pending | pending | 0 | pending | `E2-P2B-PARITY1`, `E2-P2E-PARITY1` |
| P2 | Graph JSON/Ladybug differing corrected fields | fields | pending accepted Child 01 fields | pending | pending | 0 | pending | `E2-P2B-PARITY1`, `E2-P2E-PARITY1` |
| P2 | corrected records silently dropped | records | pending P2-A denominator | pending | pending | 0 | pending | `E2-P2B-PARITY1`, `E2-P2E-PARITY1` |
| P2 | affected reader acceptance | passed/total affected readers | pending P2-A denominator | pending | pending | all/all | pending | `E2-P2C-RUNTIME1`, `E2-P2E-READERS1` |
| P2 | repeated normal analyze acceptance | passed/declared matched runs | pending P2-D protocol | pending | pending | all/all | pending | `E2-P2D-REPEAT1`, `E2-P2E-REPEAT1` |
| P2 | failed owned-boundary attempts returning a successful result from another artifact | attempts | pending fault denominator | pending | pending | 0 | pending | `E2-P2D-REPEAT1`, `E2-P2E-REPEAT1` |
| P2 | analyze duration, if persistence changes or measures it | seconds | pending matched baseline | pending | pending | measured and reviewed | pending | `E2-P2D-REPEAT1` |
| P2 | Ladybug load/query duration, if benchmarkable | seconds or milliseconds | pending matched baseline | pending | pending | measured and reviewed | pending | `E2-P2B-PARITY1`, `E2-P2D-REPEAT1` |
| P2 | graph artifact size, if persistence changes it | bytes | pending matched baseline | pending | pending | measured with explained delta | pending | `E2-P2B-PARITY1` |

## Non-Benchmarkable Notes

- P2-A source inventory, full-build pass/fail, affected-reader QA pass/fail, detect-changes, commits, and Supervisor verdicts are evidence gates.
- Child 07 owns campaign-wide performance acceptance across all semantic corrections.
