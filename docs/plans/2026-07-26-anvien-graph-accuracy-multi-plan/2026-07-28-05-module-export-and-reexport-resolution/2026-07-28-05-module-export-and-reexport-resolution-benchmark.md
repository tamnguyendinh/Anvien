# Anvien Module Export And Re-Export Resolution Benchmark Ledger

## Metadata

- Date: `2026-07-28`
- Plan: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-05-module-export-and-reexport-resolution/2026-07-28-05-module-export-and-reexport-resolution-plan.md`
- Evidence: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-05-module-export-and-reexport-resolution/2026-07-28-05-module-export-and-reexport-resolution-evidence.md`
- Benchmark: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-05-module-export-and-reexport-resolution/2026-07-28-05-module-export-and-reexport-resolution-benchmark.md`
- Actual status: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-05-module-export-and-reexport-resolution/2026-07-28-05-module-export-and-reexport-resolution-actual-status.md`

## Benchmark Rules

- Record only Child 05 module/export/re-export measurements.
- Use the exact source sites, input corpus, configuration, built runtime, and command for every baseline/final comparison.
- The accepted investigation numbers are bounded baseline context, not a current final measurement.
- Record absolute physical path-resolution and syntactic `IMPORTS` counts before reporting their deltas.
- Aggregate graph counts do not replace the two terminal call-site or proof checks.
- Build/test pass-fail belongs in the evidence ledger. Record timing, size, throughput, or memory here only when Child 05 changes that measured system.
- Do not invent fixed topology or performance limits before P5-A/P5-C measurement establishes the relevant behavior and Owner accepts a limit.

## B0 - P0 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P0 | bounded source imports whose module path reaches the barrel | resolved/expected | 2/2 | 2/2 | pending | preserve 2/2 | 0 | `E0-P0A-VERIFY1` |
| P0 | bounded barrel calls with terminal `CALLS` | resolved/expected | 0/2 | 0/2 | pending | 2/2 | 0 | `E0-P0A-VERIFY1` |
| P0 | target import facts in accepted differential | count | 10 | 10 | pending | unchanged by export traversal | 0 | `E0-P0A-VERIFY1` |
| P0 | target import-use bindings in accepted differential | count | 1 actual / 3 direct-import control | same bounded baseline | pending | recover the 2 missing bindings at exact sites | pending | `E0-P0A-VERIFY1` |
| P0 | target resolved calls in accepted differential | count | 37 actual / 39 direct-import control | same bounded baseline | pending | recover the 2 exact calls; aggregate remeasured | pending | `E0-P0A-VERIFY1` |
| P0 | target unresolved references in accepted differential | count | 542 actual / 540 direct-import control | same bounded baseline | pending | remove the 2 exact false sites; aggregate remeasured | pending | `E0-P0A-VERIFY1` |

## B5 - P5 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P5-A | parsed-code corpus | files | 736 | 736 post-change | pending Supervisor | preserve the same source corpus | 0 | `E5-P5A-COUNT1` |
| P5-A | physical target-file resolutions (`resolution.ImportsResolved`) | absolute count | 5,072 | 5,072 post-change | pending Supervisor | preserve for the same input | 0 | `E5-P5A-COUNT1`, `E5-P5D-COUNT1` |
| P5-A | resolver-emitted syntactic `IMPORTS` (`resolution.FinalizedImportsEmitted`) | absolute count | 5,072 | 5,072 post-change | pending Supervisor | preserve for the same input | 0 | `E5-P5A-COUNT1`, `E5-P5D-COUNT1` |
| P5-A | final persisted graph-wide `IMPORTS` | absolute count | 5,088 | 5,088 post-change | pending Supervisor | preserve for the same input | 0 | `E5-P5A-COUNT1`, `E5-P5D-COUNT1` |
| P5-B | zero-physical barrel declarations | count | fixture not yet created | pending | pending | exactly 0 | pending | `E5-P5B-ZEROBARREL1` |
| P5-B | zero-physical barrel export entries | count | no export table | pending | pending | exact non-zero fixture expectation | pending | `E5-P5B-ZEROBARREL1` |
| P5-C | named traversal vectors | passed/total | no accepted implementation | pending | pending | 100% of accepted alias/star/namespace/meaning/cycle/ambiguity vectors | pending | `E5-P5C-PROOF1`, `E5-P5C-TEST1` |
| P5-C | explicit-import global-name rescues | count | current branch exists; exact fixture count pending | pending | pending | 0 | pending | `E5-P5C-NOGLOBAL1` |
| P5-D | direct/barrel terminal equality | equal/expected pairs | 0/2 | 0/2 | pending | 2/2 | pending | `E5-P5D-TARGET1`, `E5-P5D-ORACLE1` |
| P5-D | bounded barrel terminal calls | resolved/expected | 0/2 | 0/2 | pending | 2/2 | pending | `E5-P5D-TARGET1` |
| P5-D | matching false-gap sites | count | 2 bounded sites | 2 | pending | 0 | pending | `E5-P5D-TARGET1` |
| P5-D | complete barrel proof chains | complete/expected | 0/2 | 0/2 | pending | 2/2 | pending | `E5-P5D-TARGET1` |
| P5-D | physical target-file resolution delta | count | 5,072 P5-A baseline | pending | pending | 0 | pending | `E5-P5D-COUNT1` |
| P5-D | resolver-emitted syntactic `IMPORTS` delta | count | 5,072 P5-A baseline | pending | pending | 0 | pending | `E5-P5D-COUNT1` |
| P5-D | final persisted graph-wide `IMPORTS` delta | count | 5,088 P5-A baseline | pending | pending | 0 | pending | `E5-P5D-COUNT1` |
| P5-D | affected persistence/reader field differences | count | pending affected-surface inventory | pending | pending | 0 across every affected row | pending | `E5-P5D-PARITY1` |

## Conditional Product Measurements

Capture these only if the accepted P5 implementation changes the measured system. Record comparable baseline/final runs with the same corpus, configuration, build, machine, and cache policy.

| Metric | Unit | Baseline | Latest | Final | Target | Evidence |
|--------|------|----------|--------|-------|--------|----------|
| analyze duration | milliseconds | pending if applicable | pending | pending | measured; any acceptance threshold requires Owner decision | `E5-P5D-COUNT1` |
| peak analyze memory | bytes | pending if applicable | pending | pending | measured; any acceptance threshold requires Owner decision | `E5-P5D-COUNT1` |
| graph node/relationship size change attributable to P5 | counts | pending if applicable | pending | pending | explain exact expected delta; no unexplained loss | `E5-P5D-PARITY1` |

## Non-Benchmarkable Notes

- P0 authority classification, file ownership, Supervisor review, cleanup, detect-changes, and commits are evidence gates rather than benchmark measurements.
- Declaration-universe and unrelated command/reader metrics belong to their owning children, not Child 05.
