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
| P0 | source-proven affected readers | readers | unknown | not a P0 acceptance metric; accepted P2-A denominator is `8` | `8` accepted for later P2-C validation | exact discovered count, zero unassigned | N/A at P0 | `E2-P2A-INVENTORY2`, `E2-P2A-MATRIX2`, `E2-P2A-REVIEW1` |
| P0 | Child 02 implementation slices | slices | 5 planned | 5 planned | pending | 5 accepted | pending | plan checklist |

## B2 - P2 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P2 | affected persistence paths | unique owners | unknown before P2-A source inventory | `8` accepted (`5` future edit, `3` validate-only) | `8` accepted | exact count, zero unassigned | `+8` from unknown | `E2-P2A-INVENTORY2`, `E2-P2A-MATRIX2`, `E2-P2A-REVIEW1` |
| P2 | affected readers | unique readers | unknown before P2-A source inventory | `8` accepted (`4` future edit, `4` validate-only) | `8` accepted | exact count, zero unassigned | `+8` from unknown | `E2-P2A-INVENTORY2`, `E2-P2A-MATRIX2`, `E2-P2A-REVIEW1` |
| P2 | repeated-analyze/failure boundaries | unique boundaries | unknown before P2-A source inventory | `2` accepted (`analyze` shared with persistence, plus `Server.runCypherRead`) | `2` accepted | exact owned boundaries | `+2` from unknown | `E2-P2A-INVENTORY2`, `E2-P2A-MATRIX2`, `E2-P2A-REVIEW1` |
| P2 | unique classified frozen owner rows | rows | unknown before P2-A source inventory | `19` accepted, including `2` explicit out-of-campaign audit/probe exclusions | `19` accepted | every frozen row classified exactly once | `+19` from unknown | `E2-P2A-MATRIX2`, `E2-P2A-REVIEW1` |
| P2 | duplicate / unassigned / unclassified affected rows | rows | unknown before P2-A source inventory | `0 / 0 / 0` accepted | `0 / 0 / 0` accepted | `0 / 0 / 0` | zero | `E2-P2A-MATRIX2`, `E2-P2A-REVIEW1` |
| P2 | Graph JSON/Ladybug corrected records with field differences | records | pending accepted Child 01 fields | `0` across `36,611` matched Definition records | `0` across `36,611` | 0 | target met at P2-E | `E2-P2B-PARITY1`, `E2-P2E-PARITY1`, `E2-P2E-REVIEW1` |
| P2 | Graph JSON/Ladybug differing corrected fields | fields | pending accepted Child 01 fields | `0` for identity, label, name, path, qualified name, construct and optional selection coordinates | `0` | 0 | target met at P2-E | `E2-P2B-PARITY1`, `E2-P2E-PARITY1`, `E2-P2E-REVIEW1` |
| P2 | corrected records silently dropped | records | pending P2-A denominator | `0` (`36,611/36,611` matched) | `0` | 0 | target met at P2-E | `E2-P2B-PARITY1`, `E2-P2E-PARITY1`, `E2-P2E-REVIEW1` |
| P2 | optional SelectionRange integrity | present / absent / partial records | pending accepted Child 01 fields | `4,941 / 31,670 / 0` | `4,941 / 31,670 / 0` | all-or-none; absence NULL | target met; `5,650` real zero `startCol` retained | `E2-P2E-PARITY1`, `E2-P2E-REVIEW1` |
| P2 | exact `DEFINES` endpoint parity | Graph JSON pairs / Ladybug pairs / missing endpoints | pending accepted Child 01 facts | `36,611 / 36,611 / 0` | `36,611 / 36,611 / 0` | exact pair parity; zero missing endpoints | target met at P2-E | `E2-P2E-PARITY1`, `E2-P2E-REVIEW1` |
| P2 | affected reader acceptance | passed/total affected readers | pending P2-A denominator | `8/8` independently revalidated at P2-E | `8/8` | all/all | `+8` accepted readers from pending | `E2-P2C-RUNTIME1`, `E2-P2C-REVIEW1`, `E2-P2E-READERS1`, `E2-P2E-REVIEW1` |
| P2 | complete P2-E owner-row matrix | passed/total rows | pending accepted P2-A denominator | `17/17` C01-C17 | `17/17` | all/all | target met at P2-E | `E2-P2E-PARITY1`, `E2-P2E-READERS1`, `E2-P2E-REPEAT1`, `E2-P2E-REVIEW1` |
| P2 | repeated normal analyze acceptance | passed/declared matched runs | pending P2-D protocol | `4/4` successful runs: unchanged `2/2`, changed `1/1`, recovery `1/1`; full behavior matrix `7/7` | `4/4` successful; `7/7` full matrix | all/all | target met at P2-E | `E2-P2D-REPEAT1`, `E2-P2D-REVIEW1`, `E2-P2D-COMMIT1`, `E2-P2E-REPEAT1`, `E2-P2E-REVIEW1` |
| P2 | failed owned-boundary attempts returning a successful result from another artifact | attempts | pending fault denominator | `0/2` across analyze-storage and no-readable-backend faults | `0/2` | 0 | target met at P2-E | `E2-P2D-REPEAT1`, `E2-P2D-REVIEW1`, `E2-P2E-REPEAT1`, `E2-P2E-REVIEW1` |
| P2 | analyze duration, if persistence changes or measures it | seconds | pending matched baseline | unchanged `5.092s` / `3.271s`; changed `3.365s`; recovery `3.215s` | same P2-E measurements | measured and reviewed | measured; no pre-P2-D baseline delta claimed | `E2-P2E-REPEAT1`, `E2-P2E-REVIEW1` |
| P2 | Ladybug load/query duration, if benchmarkable | seconds | pending matched baseline | native current read `0.374s`; no-backend non-success `0.061s`; recovery read `0.354s` | same P2-E measurements | measured and reviewed when the boundary is benchmarked | measured; no pre-P2-D baseline delta claimed | `E2-P2E-REPEAT1`, `E2-P2E-REVIEW1` |
| P2 | graph artifact size, if persistence changes it | bytes | pending matched baseline | Graph JSON `391,303,083`; Ladybug `145,162,240` | same accepted P2-E artifacts | measured with explained delta | current artifacts recorded; no matched baseline delta claimed | `E2-P2E-PARITY1`, `E2-P2E-REVIEW1` |

## Non-Benchmarkable Notes

- P2-A source inventory, full-build pass/fail, affected-reader QA pass/fail, detect-changes, commits, and Supervisor verdicts are evidence gates.
- Child 07 owns campaign-wide performance acceptance across all semantic corrections.
