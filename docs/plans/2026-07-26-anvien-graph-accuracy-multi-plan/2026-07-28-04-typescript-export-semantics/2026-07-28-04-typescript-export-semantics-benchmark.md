# Anvien First-Class TypeScript Export Semantics Benchmark Ledger

## Metadata

- Date: `2026-07-28`
- Plan: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-04-typescript-export-semantics/2026-07-28-04-typescript-export-semantics-plan.md`
- Evidence: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-04-typescript-export-semantics/2026-07-28-04-typescript-export-semantics-evidence.md`
- Benchmark: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-04-typescript-export-semantics/2026-07-28-04-typescript-export-semantics-benchmark.md`
- Actual status: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-04-typescript-export-semantics/2026-07-28-04-typescript-export-semantics-actual-status.md`

## Benchmark Rules

- Record only Child 04 export-syntax and direct-projection measurements. Terminal module resolution, public API, ambient, and reader metrics belong to their owning children.
- The accepted investigation supplies the bounded `0/21` target baseline. P0-A additionally records the current first-class contract count, syntax-path coverage, graph property population, and affected persistence dialects without treating source inspection as post-implementation success.
- Use the same source/config/analyzer basis for comparable before/after target measurements.
- Build/test pass-fail belongs in the evidence ledger. Record timing only if a Child 04 slice intentionally changes measured extraction/analyze performance.
- Update `Latest`, `Final`, `Delta`, and exact evidence IDs immediately after a valid measurement.

## B0 - P0 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P0 | bounded exported definitions represented as definitions | definitions / expected | 21/21 | 21/21 accepted baseline | 21/21 P0 baseline | 21/21 preserved | 0 | `E0-P0A-VERIFY1`, `E0-P0A-VERIFY2` |
| P0 | bounded direct exports with export metadata | correct / expected | 0/21 | 0/21 accepted baseline | 0/21 P0 baseline | 21/21 | 0 | `E0-P0A-VERIFY1`, `E0-P0A-VERIFY2` |
| P0 | excluded analyze inventory | scanned / parsed code / failed | N/A | 1,126 / 626 / 0 | 1,126 / 626 / 0 | current clean basis | N/A | `E0-P0A-GRAPH1` |
| P0 | excluded graph inventory | nodes / relationships | N/A | 80,908 / 120,167 | 80,908 / 120,167 | current clean basis | N/A | `E0-P0A-GRAPH1` |
| P0 | first-class export fact contracts | contracts | 0 | 0 | 0 | 1 canonical contract | 0 | `E0-P0A-SRC1` |
| P0 | `ScopeIR` export collections | collections | 0 | 0 | 0 | 1 deterministic collection | 0 | `E0-P0A-SRC1` |
| P0 | graph nodes carrying `visibility` | nodes | 0 | 0 | 0 | truthful affected projection after P4-C | 0 | `E0-P0A-SRC1` |
| P0 | graph nodes carrying `isExported` | true / false | 0 / 5,714 | 0 / 5,714 | 0 / 5,714 | derived parity after P4-C | 0 / 0 | `E0-P0A-SRC1` |

## B4 - P4 Benchmarks

| Plan Item | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-----------|--------|------|----------|--------|-------|--------|-------|----------|
| P4-A | source-form export fact serialization cases | passed / inventoried cases | 0 first-class cases | 7/7 | 7/7 | 100% | +7 cases | `E4-P4A-TEST1`, `E4-P4A-BOUNDARY1` |
| P4-A | export facts per supplied eligible binding/specifier site at the contract boundary | facts / supplied sites | 0.0; no first-class fact contract | 7/7 = 1.0 | 7/7 = 1.0 | 1.0 exactly | +1.0 | `E4-P4A-TEST1`, `E4-P4A-BOUNDARY1` |
| P4-A | structured export diagnostic classes represented | represented / contract classes | 0/2 | 2/2 | 2/2 | 100% | +2 classes | `E4-P4A-TEST1`, `E4-P4A-BOUNDARY1` |
| P4-A | access fields changed by export fact creation | fields | 0; no export fact writer exists | 0 | 0 | 0 | 0 | `E4-P4A-SRC1`, `E4-P4A-TEST1` |
| P4-A | source-form kinds represented | kinds / required kinds | 0/7 | 7/7 | 7/7 | 7/7 | +7 kinds | `E4-P4A-TEST1`, `E4-P4A-BOUNDARY1` |
| P4-A | meaning lanes represented | lanes / required lanes | 0/3 | 3/3 | 3/3 | 3/3 | +3 lanes | `E4-P4A-TEST1`, `E4-P4A-BOUNDARY1` |
| P4-A | nested mutable export members deep-copied | members / mutable members | 0/3 | 3/3 | 3/3 | 3/3 | +3 members | `E4-P4A-TEST1`, `E4-P4A-BOUNDARY1` |
| P4-A | forbidden later-slice JSON fields present | fields / checked names | 0/7 | 0/7 | 0/7 | 0/7 | 0 | `E4-P4A-SRC1`, `E4-P4A-BOUNDARY1` |
| P4-A | closure excluded analyze inventory | scanned / parsed code / failed | 1,126 / 626 / 0 at P0 | 1,130 / 626 / 0 | 1,130 / 626 / 0 | current accepted closure basis | +4 / 0 / 0 | `E4-P4A-DETECT1` |
| P4-A | closure excluded graph inventory | nodes / relationships | 80,908 / 120,167 at P0 | 81,132 / 120,514 | 81,132 / 120,514 | current accepted closure basis | +224 / +347 | `E4-P4A-DETECT1` |
| P4-B | direct/default/alias/type-only syntax cases | first-class fact paths / inventoried classes | 0/4 | 4/4 | 4/4 | 4/4 | +4/4 | `E4-P4B-TEST1`, `E4-P4B-BOUNDARY1` |
| P4-B | false-positive direct exports in negative controls | definitions | not fixture-measured; P4-B oracle required | 0 | 0 | 0 | 0 | `E4-P4B-TEST1` |
| P4-B | focused direct/export diagnostic matrix | passed tests / selected tests | 0/0 | 4/4 | 4/4 | 4/4 | +4/4 | `E4-P4B-TEST1` |
| P4-B | nearest provider/ScopeIR boundary packages | passed packages / selected packages | 0/0 | 3/3 | 3/3 | 3/3 | +3/3 | `E4-P4B-BOUNDARY1` |
| P4-B | excluded graph inventory at review basis | scanned / parsed / failed | 1,130 / 626 / 0 | 1,136 / 626 / 0 | 1,136 / 626 / 0 | current closure basis | +6 / 0 / 0 | `E4-P4B-IMPACT1` |
| P4-B | excluded graph inventory at review basis | nodes / relationships | 81,132 / 120,514 | 81,772 / 121,285 | 81,772 / 121,285 | current closure basis | +640 / +771 | `E4-P4B-IMPACT1` |
| P4-B1 | named/star/namespace/type-only re-export syntax cases | first-class fact paths / inventoried classes | 0/4; two import-compatibility paths exist | 4/4 | 4/4 | 4/4 | +4/4 | `E4-P4B1-TEST1`, `E4-P4B1-BOUNDARY1` |
| P4-B1 | comment-bearing recovered valid siblings | valid facts / expected facts | 1/2 per legal comment placement | 2/2 per placement | 2/2 per placement | 2/2 | +1 per placement | `E4-P4B1-BOUNDARY1`, `E4-P4B1-REVIEW1` |
| P4-B1 | malformed recovered alias diagnostics | diagnostics / malformed sites | 0/2 exact comment cases | 1/1 per case | 1/1 per case | 1/1 | +1 per case | `E4-P4B1-TEST1`, `E4-P4B1-BOUNDARY1` |
| P4-B1 | derived compatibility imports | imports / valid facts | 1/2 per legal comment placement | 2/2 per placement | 2/2 per placement | 2/2 | +1 per placement | `E4-P4B1-BOUNDARY1` |
| P4-B1 | terminal-resolution fields populated by Child 04 | fields | 0 current Child 04 fields | 0/7 | 0/7 | 0 | 0 | `E4-P4B1-BOUNDARY1`, `E4-P4B1-REVIEW1` |
| P4-B1 | excluded graph inventory at final closure basis | scanned / parsed / failed | 1,136 / 626 / 0 | 1,146 / 626 / 0 | 1,146 / 626 / 0 | current closure basis | +10 / 0 / 0 | `E4-P4B1-DETECT1` |
| P4-B1 | excluded graph inventory at final closure basis | nodes / relationships | 81,772 / 121,285 | 82,079 / 121,780 | 82,079 / 121,780 | current closure basis | +307 / +495 | `E4-P4B1-DETECT1` |
| P4-C | graph direct-export fact conservation | graph records / accepted facts | 0/0; no source export collection | 414/414 = 1.0 | 414/414 = 1.0 | 1.0 exactly | +1.0 | `E4-P4C-TEST1`, `E4-P4C-BOUNDARY1`, `E4-P4C-REVIEW1`, `E4-P4C-COMMIT1` |
| P4-C | compatibility-field drift from source export fact | differing records | no source fact; current dialects are `visibility` and `isExported` | 0 | 0 | 0 | 0 | `E4-P4C-TEST1`, `E4-P4C-REVIEW1`, `E4-P4C-COMMIT1` |
| P4-C | affected persistence field differences | differing fields | at least one known dialect mismatch; Ladybug drops `visibility` | 0 across 11,592 normalized comparisons | 0 | 0 | 0 | `E4-P4C-BOUNDARY1`, `E4-P4C-REVIEW1`, `E4-P4C-COMMIT1` |
| P4-C | affected persistence orphan references | references | N/A before first-class export facts | 0 | 0 | 0 | 0 | `E4-P4C-TEST1`, `E4-P4C-BOUNDARY1`, `E4-P4C-REVIEW1`, `E4-P4C-COMMIT1` |
| P4-C | Graph Export inventory | nodes / File→Export relations | 0/0 | 414 / 414 | 414 / 414 | 414 / 414 | +414 / +414 | `E4-P4C-TEST1`, `E4-P4C-BOUNDARY1`, `E4-P4C-COMMIT1` |
| P4-C | normalized Graph JSON↔Ladybug field comparison | fields compared / differing fields | N/A before projection | 11,592 / 0 | 11,592 / 0 | 11,592 / 0 | +11,592 / 0 | `E4-P4C-REVIEW1`, `E4-P4C-COMMIT1` |
| P4-C | fresh closure analyze inventory | scanned / parsed code / failed | 1,147 / 626 / 0 at successor basis | 1,857 / 735 / 0 | 1,857 / 735 / 0 | current clean basis | +710 / +109 / 0 | `E4-P4C-DETECT1`, `E4-P4C-COMMIT1` |
| P4-C | fresh closure graph inventory | nodes / relationships | 82,087 / 121,788 at successor basis | 113,523 / 156,030 | 113,523 / 156,030 | current clean basis | +31,436 / +34,242 | `E4-P4C-DETECT1`, `E4-P4C-COMMIT1` |
| P4-C2 | bounded target direct exports against sealed source-only oracle | correct / expected | 0/21 accepted baseline | 21/21 | 21/21 REVIEW2 PASS | 21/21 | +21 overall rows | `E4-P4C2-ORACLE1`, `E4-P4C2-TARGET1`, `E4-P4C2-REVIEW2` |
| P4-C2 | owner-qualified target negative controls remaining non-exported | correct / expected; false positives | pending source-only oracle | 11/11; false positives 0 | 11/11; false positives 0 | 11/11; false positives 0 | N/A | `E4-P4C2-ORACLE1`, `E4-P4C2-TARGET1`, `E4-P4C2-REVIEW2` |
| P4-C2 | target access/export conflations and Child 05-derived claims | records | pending source-only oracle | 0 / 0 | 0 / 0 REVIEW2 PASS | 0 / 0 | N/A | `E4-P4C2-TARGET1`, `E4-P4C2-BOUNDARY1`, `E4-P4C2-REVIEW2` |
| P4-C2 | exported TypeAlias compatibility/affected-reader mismatches | definitions / expected TypeAlias rows | N/A at bounded baseline | 0/15 mismatched | 0/15 mismatched REVIEW2 PASS | 0/15 mismatched | -15 mismatches from REVIEW1 | `E4-P4C2-TARGET1`, `E4-P4C2-REVIEW2` |
| P4-C2 | target retry analyze inventory | scanned / parsed / failed | N/A | 1,359 / 887 / 0 | 1,359 / 887 / 0 | one valid retry basis | N/A | `E4-P4C2-TARGET1` |
| P4-C2 | target retry graph inventory | nodes / relationships | N/A | 94,422 / 125,299 | 94,422 / 125,299 | current retry basis | N/A | `E4-P4C2-TARGET1` |
| P4-C2 | target persistence field parity | compared / differing fields | N/A | 588 / 0 | 588 / 0 REVIEW2 PASS | 588 / 0 | N/A | `E4-P4C2-TARGET1`, `E4-P4C2-BOUNDARY1`, `E4-P4C2-REVIEW2` |
| P4-C2 | fresh target Graph JSON / Ladybug size | bytes / bytes | N/A | 432,026,884 / 150,355,968 | 432,026,884 / 150,355,968 | recorded capacity | N/A | `E4-P4C2-TARGET1` |
| P4-C2 | fresh pre-commit self analyze inventory | scanned / parsed / failed | 1,917 / 736 / 0 at accepted repair build | 1,939 / 736 / 0 | 1,939 / 736 / 0 | current closure basis | +22 / 0 / 0 | `E4-P4C2-DETECT1` |
| P4-C2 | fresh pre-commit self graph inventory | nodes / relationships | 114,546 / 157,361 at accepted repair build | 114,628 / 157,443 | 114,628 / 157,443 | current closure basis | +82 / +82 | `E4-P4C2-DETECT1` |

## Non-Benchmarkable Notes

- Report/authority classification, file ownership, Supervisor verdicts, full-build results, detect-changes, and commits are evidence gates rather than benchmark metrics.
- Oracle row count, provenance completeness, seal identity, durable-path compliance, and `.tmp` absence are evidence-integrity gates rather than product benchmarks; they close only through `E4-P4C2-ORACLE1` and later Supervisor review.
- No complete module-resolution, public-API, or global TypeScript accuracy percentage is claimed by this child.
