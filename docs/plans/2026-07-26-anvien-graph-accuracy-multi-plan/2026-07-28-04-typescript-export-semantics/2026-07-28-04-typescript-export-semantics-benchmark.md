# Anvien First-Class TypeScript Export Semantics Benchmark Ledger

## Metadata

- Date: `2026-07-28`
- Plan: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-04-typescript-export-semantics/2026-07-28-04-typescript-export-semantics-plan.md`
- Evidence: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-04-typescript-export-semantics/2026-07-28-04-typescript-export-semantics-evidence.md`
- Benchmark: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-04-typescript-export-semantics/2026-07-28-04-typescript-export-semantics-benchmark.md`
- Actual status: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-04-typescript-export-semantics/2026-07-28-04-typescript-export-semantics-actual-status.md`

## Benchmark Rules

- Record only Child 04 export-syntax and direct-projection measurements. Terminal module resolution, public API, ambient, and reader metrics belong to their owning children.
- The accepted investigation supplies a bounded `0/21` baseline only. General syntax coverage and affected persistence consumers remain pending until P0 inventories current source.
- Use the same source/config/analyzer basis for comparable before/after target measurements.
- Build/test pass-fail belongs in the evidence ledger. Record timing only if a Child 04 slice intentionally changes measured extraction/analyze performance.
- Update `Latest`, `Final`, `Delta`, and exact evidence IDs immediately after a valid measurement.

## B0 - P0 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P0 | bounded exported definitions represented as definitions | definitions / expected | 21/21 | 21/21 accepted baseline | pending | 21/21 preserved | pending | `E0-P0A-VERIFY1`, `E0-P0A-VERIFY2` |
| P0 | bounded direct exports with export metadata | correct / expected | 0/21 | 0/21 accepted baseline | pending | 21/21 | pending | `E0-P0A-VERIFY1`, `E0-P0A-VERIFY2` |

## B4 - P4 Benchmarks

| Plan Item | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-----------|--------|------|----------|--------|-------|--------|-------|----------|
| P4-A | export fact serialization cases | passed / inventoried cases | pending P0 | pending | pending | 100% | pending | `E4-P4A-TEST1` |
| P4-A | export facts per eligible binding/specifier | facts / eligible sites | pending P0 | pending | pending | 1.0 exactly | pending | `E4-P4A-TEST1` |
| P4-A | unsupported-export diagnostic coverage | diagnosed / unsupported fixtures | pending P0 | pending | pending | 100% | pending | `E4-P4A-BOUNDARY1` |
| P4-A | access fields changed by export fact creation | fields | pending P0 | pending | pending | 0 | pending | `E4-P4A-TEST1` |
| P4-B | direct/default/alias/type-only syntax cases | passed / inventoried cases | pending P0 | pending | pending | 100% | pending | `E4-P4B-TEST1` |
| P4-B | false-positive direct exports in negative controls | definitions | pending P0 | pending | pending | 0 | pending | `E4-P4B-TEST1` |
| P4-B1 | named/star/namespace/type-only re-export syntax cases | passed / inventoried cases | pending P0 | pending | pending | 100% | pending | `E4-P4B1-TEST1` |
| P4-B1 | terminal-resolution fields populated by Child 04 | fields | pending P0 | pending | pending | 0 | pending | `E4-P4B1-BOUNDARY1` |
| P4-C | graph direct-export fact conservation | graph records / accepted facts | pending P0 | pending | pending | 1.0 exactly | pending | `E4-P4C-TEST1` |
| P4-C | compatibility-field drift from source export fact | differing records | pending P0 | pending | pending | 0 | pending | `E4-P4C-TEST1` |
| P4-C | affected persistence field differences | differing fields | pending affected-surface inventory | pending | pending | 0 | pending | `E4-P4C-BOUNDARY1` |
| P4-C | affected persistence orphan references | references | pending affected-surface inventory | pending | pending | 0 | pending | `E4-P4C-BOUNDARY1` |
| P4-C2 | bounded target direct exports | correct / expected | 0/21 | pending | pending | 21/21 | pending | `E4-P4C2-TARGET1` |
| P4-C2 | false-positive target negative controls | definitions | pending oracle | pending | pending | 0 | pending | `E4-P4C2-TARGET1` |
| P4-C2 | target access/export conflations | records | pending oracle | pending | pending | 0 | pending | `E4-P4C2-TARGET1` |

## Non-Benchmarkable Notes

- Report/authority classification, file ownership, Supervisor verdicts, full-build results, detect-changes, and commits are evidence gates rather than benchmark metrics.
- No complete module-resolution, public-API, or global TypeScript accuracy percentage is claimed by this child.
