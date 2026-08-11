# Anvien Graph Identity Contract and Strict Construction Benchmark Ledger

## Metadata

- Date: `2026-07-28`
- Plan: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-plan.md`
- Evidence: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-evidence.md`
- Benchmark: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-benchmark.md`
- Actual status: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-actual-status.md`
- Roadmap: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-26-anvien-graph-accuracy-roadmap.md`

## Benchmark Rules

- Record measured identity-scope values only.
- The bounded target source oracle defines the `2/4 -> 4/4` accuracy target.
- Use the same source, configuration, analyzer build, machine, and cache policy for matched baseline/final runtime measurements.
- Build/test pass-fail and Supervisor verdicts belong in the evidence ledger.
- Record graph size, analyze duration, and peak RSS only if Child 01 implementation changes those measurements or the owning slice establishes a matched baseline.
- Do not copy binding, export, barrel, ambient, reader, or campaign-wide metrics into this Child.

## B0 - P0 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P0 | bounded same-name declarations represented distinctly | represented/source declarations | `2/4` | `2/4` bounded finding | pending | `4/4` | pending | `E0-P0A-REPORT1`, `E0-P0A-VERIFY1` |
| P0 | bounded identity collision groups | groups | 2 (`time`, `now`) | 2 | pending | 0 | pending | `E0-P0A-REPORT1`, `E0-P0A-VERIFY1` |
| P0 | Anvien analyze inventory | scanned / parsed-code / failed files | `1,557 / 676 / 0` | `1,557 / 676 / 0` | pending implementation | preserve baseline before code | `0` | `E0-P0A-GRAPH1` |
| P0 | Anvien graph inventory | nodes / relationships | `85,110 / 123,978` (pre-ledger) | `85,114 / 123,982` (post-ledger) | pending implementation | record documentation-only delta; no production delta | `+4 / +4` | `E0-P0A-GRAPH1` |
| P0 | source candidate surfaces with fresh file-detail | files | 5 retained candidates | `15/15` audited | `15/15` P0 | every editable candidate has a current count | `+10` | `E0-P0A-FD1..E0-P0A-FD15` |
| P0 | production callers of `BuildDefinitionIndex` | callers | `0` | `0` | preserve | `0` | `0` | `E0-P0A-SOURCE3`, `E0-P0A-IMPACT2` |
| P0 | Child 01 implementation slices | slices | 5 planned | 5 planned | pending | 5 accepted | pending | plan checklist |

## B1 - P1 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P1 | bounded same-name declarations represented distinctly | represented/source declarations | `2/4` | `4/4` | `4/4` target measurement | `4/4` | `+2/4` represented | `E1-P1E-TARGET1` |
| P1 | identity collision groups for the bounded oracle | groups | 2 | 0 | 0 target measurement | 0 | `-2` | `E1-P1D-COLLISION1`, `E1-P1E-INTEGRITY1`, `E1-P1E-TARGET1` |
| P1 | declaration occurrence conservation | output/input occurrences | not measured before P1-C; denominator is ScopeIR definitions carried by production `defsByFile` | package fixture `100%`; five built-runtime runs each retain `10/10`; target retains `4/4` unique Variable IDs and `4/4` `DEFINES` | `100%` at local and bounded target boundaries | `100%`, zero unexplained drops | target met | `E1-P1C-TEST1`, `E1-P1C-ORACLE1`, `E1-P1E-INTEGRITY1`, `E1-P1E-TARGET1` |
| P1 | affected relationships with missing endpoints | relationships | pending P1-E baseline | `0/19` in each fixture run; `0/4` bounded target `DEFINES` | 0 at local and target boundaries | 0 | target met | `E1-P1E-INTEGRITY1`, `E1-P1E-TARGET1` |
| P1 | matched deterministic analyze results | equal runs/total runs | pending matched baseline | `5/5` | `5/5` local integration | `5/5` | target met | `E1-P1E-DETERMINISM1` |
| P1 | graph size for the matched identity fixture | bytes | no matched pre-P1-C artifact | `44,199` in every run | `44,199` local integration | measured with explained delta | stable `5/5`; pre-change delta N/A | `E1-P1E-DETERMINISM1` |
| P1 | analyze duration median | seconds | no matched pre-P1-C timing | `3.220` (`3.043`–`3.469`) | `3.220` local integration | measured and reviewed | pre-change delta N/A | `E1-P1E-DETERMINISM1` |
| P1 | peak analyze RSS | bytes | no matched pre-P1-C RSS | max `456,040,448`; median run `444,989,440` | `456,040,448` max local integration | measured and reviewed | pre-change delta N/A | `E1-P1E-DETERMINISM1` |
| P1-E | target analyze inventory | scanned / parsed-code / failed files | `1,359 / 887 / 0` current target source | `1,359 / 887 / 0` | `1,359 / 887 / 0` | zero failed; preserve target source set | `0 / 0 / 0` | `E1-P1E-TARGET1`, `E1-P1E-BOUNDARY1` |
| P1-E | target graph inventory | nodes / relationships | `84,807 / 114,125` pre-P1-E artifact | `86,164 / 115,888` | `86,164 / 115,888` | record accepted identity output | `+1,357 / +1,763` | `E1-P1E-PREFLIGHT1`, `E1-P1E-TARGET1` |
| P1-E | target Graph JSON size | bytes | `315,569,880` | `415,998,487` | `415,998,487` | record with target preservation proof | `+100,428,607` | `E1-P1E-PREFLIGHT1`, `E1-P1E-TARGET1`, `E1-P1E-BOUNDARY1` |
| P1-E | independently repeated target analyze duration | seconds | no matched pre-change timing | `41.350` | `41.350` | measured and reviewed | pre-change delta N/A | `E1-P1E-TARGET1` |
| P1-E | independently repeated target analyze peak RSS | bytes | no matched pre-change RSS | `1,263,955,968` | `1,263,955,968` | measured and reviewed | pre-change delta N/A | `E1-P1E-TARGET1` |
| P1-D | candidate graph file inventory | scanned / parsed-code / failed files | `1,570 / 679 / 0` pre-flight | `1,574 / 680 / 0` before detect | `1,574 / 680 / 0` | record final with `0` failed | `+4 / +1 / 0` | `E1-P1D-IMPACT1`, `E1-P1D-GRAPH2`, `E1-P1D-REVIEW1`, `E1-P1D-DETECT1` |
| P1-D | candidate graph inventory | nodes / relationships | `95,577 / 134,444` pre-flight | `95,739 / 134,670` before detect | `95,739 / 134,670` | record final with explained delta | `+162 / +226` | `E1-P1D-IMPACT1`, `E1-P1D-GRAPH2`, `E1-P1D-REVIEW1`, `E1-P1D-DETECT1` |

## Non-Benchmarkable Notes

- P1-A is documentation-only and produced no new product/runtime benchmark. Its contract trace, document checks, Supervisor verdict, and commit are evidence gates.
- P1-B adds an optional in-memory ScopeIR selection range and records the existing TSJS coordinate contract; it does not change a benchmark-owned performance, capacity, package-size, graph-throughput, or target-accuracy metric. Build/test timings and the self-analyze inventory are validation evidence under `E1-P1B-BUILD1`, not benchmark results.
- P1-C package QA and the built CLI fixture measure occurrence conservation at the owned identity boundary. They do not establish the P1-E target/integration final value; build/test timings remain validation evidence.
- P1-D graph inventory changes include the production/test/ledger worktree and are not collision-correctness proof. Full-build timing and pass/fail remain validation evidence; the inventory is recorded only as a measured graph count.
- Child 02 owns persistence/reader counts and field-parity measurements.
