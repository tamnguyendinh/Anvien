# Anvien Cross-Surface Acceptance and Target Validation Benchmark Ledger

## Metadata

- Date: `2026-07-28`
- Last revised: `2026-08-24`
- Plan: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-07-cross-surface-acceptance-and-target-validation/2026-07-28-07-cross-surface-acceptance-and-target-validation-plan.md`
- Evidence: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-07-cross-surface-acceptance-and-target-validation/2026-07-28-07-cross-surface-acceptance-and-target-validation-evidence.md`
- Benchmark: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-07-cross-surface-acceptance-and-target-validation/2026-07-28-07-cross-surface-acceptance-and-target-validation-benchmark.md`
- Actual status: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-07-cross-surface-acceptance-and-target-validation/2026-07-28-07-cross-surface-acceptance-and-target-validation-actual-status.md`
- Predecessor: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy/2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy-plan.md`

## Benchmark Rules

- Record measured values only. Design decisions, command logs, build results, and Supervisor verdicts belong in the evidence ledger.
- Bind every comparison to the same corpus, source commit, configuration, analyzer build, machine, and cache policy.
- The five target denominators come from the problem-origin report and cannot be enlarged, reduced, or combined to hide a failing family.
- Create affected-reader metrics only after `E7-P7C-INVENTORY1` proves the concrete denominator.
- Use a performance value as a release gate only when its baseline method and budget were recorded before the final measurement.
- Do not average unrelated readers or native boundaries into one favorable aggregate.
- P7 performance rows consume the accepted Child 06A initial/final/equivalence/resource handoff. Current Child 06A truth is no detailed timing map, ordered bottleneck list, production attempt, accepted speedup, or closure commit; no P7 performance row may open from Child 06 directly.

## B0 - P0 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P0 | Anvien scanned files | files | 1,556 | 1,556 | pending terminal refresh | record final with explanation | 0 | `E0-P0A-GRAPH1` |
| P0 | Anvien parsed code files | files | 676 | 676 | pending terminal refresh | no unexplained loss | 0 | `E0-P0A-GRAPH1` |
| P0 | Anvien failed files | files | 0 | 0 | pending terminal refresh | 0 | 0 | `E0-P0A-GRAPH1` |
| P0 | Anvien graph nodes | nodes | 85,101 | 85,101 | pending terminal refresh | measured final | 0 | `E0-P0A-GRAPH1` |
| P0 | Anvien graph relationships | relationships | 123,969 | 123,969 | pending terminal refresh | measured final | 0 | `E0-P0A-GRAPH1` |
| P0 | Child 07 production files editable | files | 0 | 0 | 0 | 0 | 0 | `E0-P0A-DIFF1` |

## B7 - P7 Benchmarks

### B7-A — Determinism and integrity

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P7-A | successful identical-input analyze runs | runs | 0 | pending | pending | at least 5 | pending | `E7-P7A-ANALYZE1` |
| P7-A | equal canonical node fact sets | equal runs / accepted runs | pending | pending | pending | 100% | pending | `E7-P7A-DETERMINISM1` |
| P7-A | equal canonical relationship fact sets | equal runs / accepted runs | pending | pending | pending | 100% | pending | `E7-P7A-DETERMINISM1` |
| P7-A | unexplained lost source occurrences | count | pending accepted-child baseline | pending | pending | 0 | pending | `E7-P7A-INTEGRITY1` |
| P7-A | missing relationship endpoints | count | pending | pending | pending | 0 | pending | `E7-P7A-INTEGRITY1` |
| P7-A | orphan references in affected projections | count | pending | pending | pending | 0 | pending | `E7-P7A-INTEGRITY1` |

### B7-B — Terminal target accuracy

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P7-B | same-name identity | correct sites / total | 2/4 | pending | pending | 4/4 | pending | `E7-P7B-ORACLE1` |
| P7-B | binding patterns | correct sites / total | 0/6 | pending | pending | 6/6 | pending | `E7-P7B-ORACLE1` |
| P7-B | direct exports | correct sites / total | 0/21 | pending | pending | 21/21 | pending | `E7-P7B-ORACLE1` |
| P7-B | barrel calls | correct sites / total | 0/2 | pending | pending | 2/2 | pending | `E7-P7B-ORACLE1` |
| P7-B | ambient sites | correct external/capability outcomes / total | 0/3 | pending | pending | 3/3 | pending | `E7-P7B-ORACLE1` |

### B7-C — Affected surfaces and performance

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P7-C | affected reader inventory | classified readers / discovered affected readers | 0/unknown | pending | pending | 100% with no unresolved candidate | pending | `E7-P7C-INVENTORY1` |
| P7-C | affected projection drift | differing records or fields | pending | pending | pending | 0 | pending | `E7-P7C-PARITY1` |
| P7-C | affected UI rows | passed / included UI rows | 0/unknown | pending | pending | 100% or zero rows with exclusion evidence | pending | `E7-P7C-PLAY1` |
| P7-C | analyze duration | milliseconds; accepted aggregation | missing until accepted Child 06A handoff | pending | pending | accepted Child 06A final measured value/equivalence basis | pending | Child 06A `E3-P3C-HANDOFF1`; `E7-P7C-BENCH1` |
| P7-C | graph size | bytes | missing until accepted Child 06A handoff | pending | pending | accepted Child 06A final measured value/equivalence basis | pending | Child 06A `E3-P3C-HANDOFF1`; `E7-P7C-BENCH1` |
| P7-C | peak analyze memory | bytes | missing until accepted Child 06A handoff | pending | pending | accepted Child 06A final measured value/resource evidence when recorded | pending | Child 06A `E3-P3C-HANDOFF1`; `E7-P7C-BENCH1` |
| P7-C | additional affected performance paths | one row per concrete path | pending inventory | pending | pending | accepted pre-recorded budget | pending | `E7-P7C-INVENTORY1`, `E7-P7C-BENCH1` |

## Non-Benchmarkable Notes

- Full-build PASS/FAIL, real-boundary behavior, visual inspection, target worktree state, detect-changes, commits, and Supervisor verdicts are evidence, not benchmarks.
- Scanner omissions are outside this child and are not a terminal graph-accuracy metric.
- No performance row can pass with a retrospective or method-mismatched baseline.
