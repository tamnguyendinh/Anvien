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

Correction temporal note: the authorized manifest-repair analyze capture is `1,631/688/0`, `97,369` nodes, `136,643` relationships. The original accepted P0-A graph benchmark rows below remain unchanged and refer to the original graph capture used for `E0-P0A-GRAPH1`; the correction pass did not rerun or replace accepted file-detail/impact gates.

| Phase | Inventory metric | Unit | Current | Evidence |
|-------|------------------|------|--------:|----------|
| P0-A | current graph scanned files | files | 1,629 | `E0-P0A-GRAPH1` |
| P0-A | current graph parsed code files | files | 688 | `E0-P0A-GRAPH1` |
| P0-A | current graph failed files | files | 0 | `E0-P0A-GRAPH1` |
| P0-A | current graph nodes | nodes | 97,340 | `E0-P0A-GRAPH1` |
| P0-A | current graph relationships | relationships | 136,614 | `E0-P0A-GRAPH1` |
| P0-A | full-read / owner / excluded inventory | paths | 62 / 27 / 35 | `E0-P0A-SRC1` |
| P0-A | frozen owner assignment | assigned / expected | 27 / 27 | `E0-P0A-SRC1` |
| P0-A | file-detail coverage | owners / expected | 27 / 27 | `E0-P0A-FD1` |
| P0-A | production file-impact coverage | files / expected | 15 / 15 | `E0-P0A-IMPACT1` |
| P0-A | canonical symbol-impact coverage | symbols / expected | 39 / 39 | `E0-P0A-IMPACT1` |

| P0-A | correction graph refresh | scanned / parsed / failed files | 1,629 / 688 / 0 | 1,633 / 688 / 0 | 1,633 / 688 / 0 | informational | +4 / 0 / 0 | `E0-P0A-GRAPH1`, `E0-P0A-DETECT1` |
| P0-A | correction graph inventory | nodes / relationships | 97,340 / 136,614 | 97,388 / 136,662 | 97,388 / 136,662 | informational | +48 / +48 | `E0-P0A-GRAPH1`, `E0-P0A-DETECT1` |
| P0-A | documentation change impact | changed / affected files; changed sections | 0 / 0; 0 | 8 / 8; 62 | 8 / 8; 62 | docs-only | +8 / +8; +62 | `E0-P0A-DETECT1` |

These are graph/inventory counts at HEAD `181b8cb8`, not product runtime-performance measurements. P0-A ran no build, tests, runtime, or target gate, so none is fabricated here.

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
