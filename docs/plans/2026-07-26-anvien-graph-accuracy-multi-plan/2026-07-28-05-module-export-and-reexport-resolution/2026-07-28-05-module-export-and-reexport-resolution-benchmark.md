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
| P5-A | parsed-code corpus | files | 736 | 736 post-change | 736 accepted | preserve the same source corpus for fixed-corpus comparisons | 0 | `E5-P5A-COUNT1`, `E5-P5B-COUNT1` |
| P5-A | physical target-file resolutions (`resolution.ImportsResolved`) | absolute count | 5,072 | 5,072 post-change | 5,072 accepted | preserve for the same input | 0 | `E5-P5A-COUNT1`, `E5-P5B-COUNT1`, `E5-P5D-COUNT1` |
| P5-A | resolver-emitted syntactic `IMPORTS` (`resolution.FinalizedImportsEmitted`) | absolute count | 5,072 | 5,072 post-change | 5,072 accepted | preserve for the same input | 0 | `E5-P5A-COUNT1`, `E5-P5B-COUNT1`, `E5-P5D-COUNT1` |
| P5-A | final persisted graph-wide `IMPORTS` | absolute count | 5,088 | 5,088 post-change | 5,088 accepted | preserve for the same input | 0 | `E5-P5A-COUNT1`, `E5-P5B-COUNT1`, `E5-P5D-COUNT1` |
| P5-B | full graph inventory | nodes / relationships | 114,852 / 157,754 pre-implementation | 115,099 / 158,118 post-build | 115,099 / 158,118 Supervisor-accepted | no unexplained loss; new source/test/report corpus is retained | +247 / +364 | `E5-P5B-IMPACT1`, `E5-P5B-BUILD1`, `E5-P5B-REVIEW1`, `E5-P5B-DETECT1`, `E5-P5B-COMMIT1` |
| P5-B | zero-physical barrel declarations | count | fixture not yet created | 0 | 0 | exactly 0 | 0 | `E5-P5B-ZEROBARREL1` |
| P5-B | zero-physical barrel export surface | explicit keys / star adjacency | no export table | 2 / 1 | 2 / 1 | exact non-zero fixture expectation without implicit default | +2 / +1 | `E5-P5B-ZEROBARREL1` |
| P5-C | graph inventory | nodes / relationships | 115,134 / 158,153 pre-implementation at HEAD `cd35b48f` | 115,902 / 159,581 final reject-repair build | 115,902 / 159,581 Supervisor-accepted and committed candidate | retain source/test/report growth without unexplained loss | +768 / +1,428 candidate growth | `E5-P5C-IMPACT1`, `E5-P5C-BUILD1`, `E5-P5C-REVIEW1`, `E5-P5C-COMMIT1` |
| P5-C | named traversal vectors | passed/total | no accepted implementation | 11/11 final named tests/subtests | 11/11 Supervisor-accepted | 100% of accepted alias/star/namespace/meaning/cycle/ambiguity vectors, including ambiguous-owner member composition | +11 accepted vectors | `E5-P5C-PROOF1`, `E5-P5C-TEST1`, `E5-P5C-REVIEW1` |
| P5-C | explicit-import global-name rescues | count | current repository-global fallback exists | 0 in focused final fixture | 0 Supervisor-accepted | 0 | target reached | `E5-P5C-NOGLOBAL1`, `E5-P5C-REVIEW1` |
| P5-C | fixed-corpus physical target-file resolutions | absolute count | 5,072 on 736 parsed-code files | 5,072 after subtracting exact 89-edge P5-B/P5-C corpus growth from 5,161 | 5,072 Supervisor-accepted | preserve 5,072 | 0 | `E5-P5C-PROOF1`, `E5-P5C-BUILD1`, `E5-P5C-REVIEW1` |
| P5-C | fixed-corpus resolver-emitted syntactic `IMPORTS` | absolute count | 5,072 on 736 parsed-code files | 5,072 after the same exact corpus decomposition | 5,072 Supervisor-accepted | preserve 5,072 | 0 | `E5-P5C-PROOF1`, `E5-P5C-BUILD1`, `E5-P5C-REVIEW1` |
| P5-C | fixed-corpus persisted graph-wide `IMPORTS` | absolute count | 5,088 on 736 parsed-code files | 5,177 current minus 89 exact new-file edges = 5,088 | 5,088 Supervisor-accepted | preserve 5,088 | 0 | `E5-P5C-PROOF1`, `E5-P5C-BUILD1`, `E5-P5C-REVIEW1` |
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
- `E5-P5B-DETECT1` is recorded as an evidence-only gate: the fresh graph had `115,134` nodes / `158,153` relationships; detect-changes reported `18` changed symbols across `5` changed files and `1` affected process, with `risk_level=medium` and `fileLayer.changedFileRisk=high`. These figures are not substituted for the fixed-corpus or full-build measurements above.
- `E5-P5B-COMMIT1` is an evidence-only closure gate for the isolated commit `c1559df953a277b099009f8489576d00ed25aa58`; it does not add a performance or capacity measurement.
- `E5-P5C-REPAIR1` and `E5-P5C-REVIEW1` are evidence-only repair/acceptance gates. The final graph and fixed-corpus counts above are benchmark measurements; REJECT/PASS status and report identities remain in the evidence ledger.
- `E5-P5C-DETECT1` is an evidence-only pre-commit gate. Its single fresh graph contained `115,947` nodes / `159,626` relationships, and detect reported `55` changed symbols across `6` changed files with `0` affected processes; these current-worktree inventory figures do not replace the fixed-corpus or Supervisor-accepted full-build measurements above.
- `E5-P5C-COMMIT1` is an evidence-only closure gate for isolated commit `76899d45a21fce55f6328b4cb30a6a5cb8719a81`; it adds no separate product performance or capacity measurement.
- Declaration-universe and unrelated command/reader metrics belong to their owning children, not Child 05.
