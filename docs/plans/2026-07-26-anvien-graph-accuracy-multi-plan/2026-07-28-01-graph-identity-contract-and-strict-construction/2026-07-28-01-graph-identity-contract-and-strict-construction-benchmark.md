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

- P1-A contract review, full-build pass/fail, behavior-test pass/fail, detect-changes, commits, and Supervisor verdicts are evidence gates.
- Child 02 owns persistence/reader counts and field-parity measurements.
