# Anvien TypeScript Binding-Pattern Extraction Benchmark Ledger

## Metadata

- Date: `2026-07-28`
- Plan: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-03-typescript-binding-pattern-extraction/2026-07-28-03-typescript-binding-pattern-extraction-plan.md`
- Evidence: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-03-typescript-binding-pattern-extraction/2026-07-28-03-typescript-binding-pattern-extraction-evidence.md`
- Benchmark: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-03-typescript-binding-pattern-extraction/2026-07-28-03-typescript-binding-pattern-extraction-benchmark.md`
- Actual status: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-03-typescript-binding-pattern-extraction/2026-07-28-03-typescript-binding-pattern-extraction-actual-status.md`

## Benchmark Rules

- Record only Child 03 binding measurements. Identity, export, module-resolution, ambient, and reader metrics belong to their owning children.
- The accepted investigation supplies bounded baseline counts only. General syntax coverage remains pending until P0 inventories the current implementation and fixtures.
- Use the same source/config/analyzer basis for comparable before/after target measurements.
- Build/test pass-fail belongs in the evidence ledger. Record timing only if a Child 03 slice intentionally changes measured extraction/analyze performance.
- Update `Latest`, `Final`, `Delta`, and exact evidence IDs immediately after a valid measurement.

## B0 - P0 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P0 | bounded legal array bindings represented | bindings / expected | 0/6 | 0/6 accepted baseline | pending | 6/6 | pending | `E0-P0A-VERIFY1`, `E0-P0A-VERIFY2` |
| P0 | bounded downstream binding sites without binding-caused gap | sites / expected | 0/6 | 0/6 accepted baseline | pending | 6/6 | pending | `E0-P0A-VERIFY1` |
| P0 | bounded silently rejected pattern cases diagnosed | diagnosed / rejected case | 0/1 | 0/1 accepted baseline | pending | 1/1 | pending | `E0-P0A-VERIFY1` |

## B3 - P3 Benchmarks

| Plan Item | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-----------|--------|------|----------|--------|-------|--------|-------|----------|
| P3-A | supported general pattern cases | passed / inventoried cases | pending P0 | pending | pending | 100% | pending | `E3-P3A-TEST1` |
| P3-A | unsupported-pattern diagnostic coverage | diagnosed / unsupported fixtures | 0/1 bounded case | pending | pending | 100% | pending | `E3-P3A-BOUNDARY1` |
| P3-A | emitted binding facts per legal leaf | facts / leaves | pending P0 | pending | pending | 1.0 exactly | pending | `E3-P3A-TEST1` |
| P3-B | variable leaves emitted when type inference fails | emitted / eligible leaves | pending P0 | pending | pending | 100% | pending | `E3-P3B-TEST1` |
| P3-B | false declarations from assignment destructuring | declarations | pending P0 | pending | pending | 0 | pending | `E3-P3B-TEST1` |
| P3-B | import-binding count delta | bindings | pending P0 | pending | pending | 0 | pending | `E3-P3B-TEST1` |
| P3-B1 | parameter binding cases | passed / inventoried cases | pending P0 | pending | pending | 100% | pending | `E3-P3B1-TEST1` |
| P3-B2 | catch binding cases | passed / inventoried cases | pending P0 | pending | pending | 100% | pending | `E3-P3B2-TEST1` |
| P3-B2A | loop declaration binding cases | passed / inventoried cases | pending P0 | pending | pending | 100% | pending | `E3-P3B2A-TEST1` |
| P3-B2A | false declarations from assignment-form loops | declarations | pending P0 | pending | pending | 0 | pending | `E3-P3B2A-BOUNDARY1` |
| P3-C | graph binding occurrence conservation | graph occurrences / accepted leaves | pending P0 | pending | pending | 1.0 exactly | pending | `E3-P3C-TEST1` |
| P3-C | nested shadowing target correctness | correct / expected targets | pending P0 | pending | pending | 100% | pending | `E3-P3C-TEST1` |
| P3-C | affected persistence field differences | differing fields | pending affected-surface inventory | pending | pending | 0 | pending | `E3-P3C-BOUNDARY1` |
| P3-C2 | bounded target bindings | correct / expected | 0/6 | pending | pending | 6/6 | pending | `E3-P3C2-TARGET1` |
| P3-C2 | bounded target sites without binding-caused gap | correct / expected | 0/6 | pending | pending | 6/6 | pending | `E3-P3C2-TARGET1` |

## Non-Benchmarkable Notes

- Report/authority classification, file ownership, Supervisor verdicts, full-build results, detect-changes, and commits are evidence gates rather than benchmark metrics.
- No global TypeScript accuracy percentage is claimed by this child.
