# Anvien Ambient And External Resolution Benchmark Ledger

## Metadata

- Date: `2026-07-28`
- Plan: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-plan.md`
- Evidence: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-evidence.md`
- Benchmark: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-benchmark.md`
- Actual status: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-actual-status.md`

## Benchmark Rules

- Record only Child 06 declaration-authority, external-target, structured-outcome, graph-health, and target measurements.
- Use exact source sites, input corpus, TypeScript configuration, selected authority inputs, built runtime, and commands for baseline/final comparisons.
- P6-A defines mechanism-specific metrics after comparing feasible designs; this ledger records only measurements required by the accepted mechanism.
- The accepted investigation numbers are bounded baseline context, not a current final measurement.
- Build/test pass-fail belongs in the evidence ledger. Record timing, size, memory, throughput, and capacity only for systems the accepted P6 design changes.
- No performance or resource limit becomes an acceptance threshold without measured baseline and explicit Owner decision.

## B0 - P0 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P0 | selected TypeScript target sites accepted by compiler oracle | resolved/expected | 3/3 | 3/3 bounded oracle | pending | preserve oracle 3/3 | 0 | `E0-P0A-VERIFY1`, `E0-P0A-TARGET1` |
| P0 | selected target sites with correct Anvien external/capability outcome | correct/expected | 0/3 | 0/3 | pending | 3/3 | 0 | `E0-P0A-TARGET1` |
| P0 | selected target sites classified as in-repository analyzer gaps | count | 3 | 3 | pending | 0 | 0 | `E0-P0A-TARGET1` |
| P0 | TypeScript standard-library declaration authority inputs in current workspace | count | 0 | 0 | pending | exact P6-A-selected authority present | 0 | `E0-P0A-SRC1` |
| P0 | structured external-capability outcomes at selected sites | count | 0 | 0 | pending | exact accepted outcome for each unavailable case | 0 | `E0-P0A-SRC2` |

## B6 - P6 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P6-A | declaration-authority decision cases | decided/required cases | 0 / pending inventory | pending | pending | 100% of source/config/runtime/consumer cases classified | pending | `E6-P6A-DECISION1` |
| P6-A | TypeScript differential vectors | passed/total | bounded target 3/3; general denominator pending | pending | pending | 100% of accepted standard-library/unavailable/config/language-isolation vectors | pending | `E6-P6A-ORACLE1` |
| P6-A | affected graph/persistence/readers | exact rows | unknown | pending | pending | complete evidence-backed inventory; no unlisted affected consumer | pending | `E6-P6A-CONSUMER1` |
| P6-B | general standard-library lookup vectors | passed/total | no authority | pending | pending | 100% of P6-A accepted vectors | pending | `E6-P6B-TEST1` |
| P6-B | target names reached through general declaration behavior | correct/expected | 0/3 | pending | pending | 3/3 without name-specific branches | pending | `E6-P6B-TEST1` |
| P6-B | deterministic authority output | equal/total comparable runs | no authority | pending | pending | 100% under P6-A protocol | pending | `E6-P6B-BENCH1` |
| P6-C1 | active project/package declaration cases | passed/required | scope undecided | pending | pending | 100% of P6-A-required cases, or evidence-backed 0 required | pending | `E6-P6C1-SCOPE1`, `E6-P6C1-TEST1` |
| P6-C2 | external target representations | correct/expected referenced targets | none | pending | pending | exact equality for P6-A consumer inventory | pending | `E6-P6C2-TEST1` |
| P6-C2 | external targets misclassified as repository-owned facts | count | N/A before representation | pending | pending | 0 | pending | `E6-P6C2-PARITY1` |
| P6-C3 | affected source sites with all required structured outcome fields | complete/affected | 0/affected | pending | pending | 100% | pending | `E6-P6C3-STATUS1`, `E6-P6C3-PARITY1` |
| P6-C3 | source sites both resolved and unresolved | count | current contract does not enforce exclusivity | pending | pending | 0 | pending | `E6-P6C3-TEST1` |
| P6-D | `Promise` target outcome | correct/expected | 0/1 | pending | pending | 1/1 external or accepted capability outcome | pending | `E6-P6D-TARGET1` |
| P6-D | `Math.max` target outcome | correct/expected | 0/1 | pending | pending | 1/1 external or accepted capability outcome | pending | `E6-P6D-TARGET1` |
| P6-D | `Math.min` target outcome | correct/expected | 0/1 | pending | pending | 1/1 external or accepted capability outcome | pending | `E6-P6D-TARGET1` |
| P6-D | selected target in-repository analyzer gaps | count | 3 | 3 | pending | 0 | pending | `E6-P6D-TARGET1` |
| P6-D | graph-health/outcome mismatches | count | 3 bounded target mismatches | pending | pending | 0 across every affected row | pending | `E6-P6D-PARITY1` |
| P6-D | affected persistence/reader field differences | count | pending affected inventory | pending | pending | 0 across every affected row | pending | `E6-P6D-PARITY1` |

## Mechanism-Specific Measurements

P6-A must add only measurements required by its accepted mechanism. Examples include packaged size, initialization time, lookup latency, peak memory, declaration input count/bytes, or cache size when the chosen design actually changes those systems. Each added row must name its baseline protocol and Owner-approved target.

| Metric | Unit | Baseline | Latest | Final | Target | Evidence |
|--------|------|----------|--------|-------|--------|----------|
| P6-A-selected mechanism measurement set | exact measured rows | pending P6-A decision | pending | pending | defined from measured evidence and Owner decision | `E6-P6A-DECISION1`, `E6-P6B-BENCH1` |
| analyze duration attributable to P6 | milliseconds | pending if applicable | pending | pending | measured; threshold requires Owner decision | `E6-P6B-BENCH1` |
| peak analyze memory attributable to P6 | bytes | pending if applicable | pending | pending | measured; threshold requires Owner decision | `E6-P6B-BENCH1` |
| graph/Ladybug size attributable to external facts | bytes/counts | pending if applicable | pending | pending | explain exact expected delta; no unexplained loss | `E6-P6C2-PARITY1` |

## Non-Benchmarkable Notes

- P0 authority classification, P6-A design reasoning, file ownership, Supervisor review, cleanup, detect-changes, and commits are evidence gates rather than benchmark measurements.
- No all-reader, public external-option, or fixed declaration-source denominator belongs here unless P6-A and fresh impact prove it part of the accepted implementation.
