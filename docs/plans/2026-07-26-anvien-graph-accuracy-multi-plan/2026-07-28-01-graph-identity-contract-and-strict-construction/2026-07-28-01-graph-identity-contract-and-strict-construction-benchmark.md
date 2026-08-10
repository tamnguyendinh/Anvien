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
| P1 | bounded same-name declarations represented distinctly | represented/source declarations | `2/4` | pending implementation | pending | `4/4` | pending | `E1-P1E-TARGET1` |
| P1 | identity collision groups for the bounded oracle | groups | 2 | pending implementation | pending | 0 | pending | `E1-P1D-COLLISION1`, `E1-P1E-INTEGRITY1` |
| P1 | declaration occurrence conservation | output/input occurrences | pending P1-C denominator | pending | pending | `100%`, zero unexplained drops | pending | `E1-P1C-ORACLE1`, `E1-P1E-INTEGRITY1` |
| P1 | affected relationships with missing endpoints | relationships | pending P1-E baseline | pending | pending | 0 | pending | `E1-P1E-INTEGRITY1` |
| P1 | matched deterministic analyze results | equal runs/total runs | pending matched baseline | pending | pending | `5/5` | pending | `E1-P1E-DETERMINISM1` |
| P1 | graph size, if identity implementation changes it | bytes | pending matched baseline | pending | pending | measured with explained delta | pending | `E1-P1E-DETERMINISM1` |
| P1 | analyze duration median, if benchmarkable | seconds | pending matched baseline | pending | pending | measured and reviewed | pending | `E1-P1E-DETERMINISM1` |
| P1 | peak analyze RSS, if benchmarkable | bytes | pending matched baseline | pending | pending | measured and reviewed | pending | `E1-P1E-DETERMINISM1` |

## Non-Benchmarkable Notes

- P1-A is documentation-only and produced no new product/runtime benchmark. Its contract trace, document checks, Supervisor verdict, and commit are evidence gates.
- P1-B adds an optional in-memory ScopeIR selection range and records the existing TSJS coordinate contract; it does not change a benchmark-owned performance, capacity, package-size, graph-throughput, or target-accuracy metric. Build/test timings and the self-analyze inventory are validation evidence under `E1-P1B-BUILD1`, not benchmark results.
- Child 02 owns persistence/reader counts and field-parity measurements.
