# Cheapapp Graph Root-Cause Restart Investigation Benchmark Ledger

## Metadata

- Date: `2026-07-26`
- Plan: `docs/plans/2026-07-26-cheapapp-graph-root-cause-restart/2026-07-26-cheapapp-graph-root-cause-restart-plan.md`
- Evidence: `docs/plans/2026-07-26-cheapapp-graph-root-cause-restart/2026-07-26-cheapapp-graph-root-cause-restart-evidence.md`
- Benchmark: `docs/plans/2026-07-26-cheapapp-graph-root-cause-restart/2026-07-26-cheapapp-graph-root-cause-restart-benchmark.md`
- Actual status: `docs/plans/2026-07-26-cheapapp-graph-root-cause-restart/2026-07-26-cheapapp-graph-root-cause-restart-actual-status.md`

## Benchmark Rules

This ledger records measured values only. Timings, graph inventory, path counts, accuracy counts, multiplicity counts, and hashes are evidence metadata; they are not performance conclusions. Command pass/fail narrative belongs in the evidence ledger.

## B0 - P0 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P0 | Target path exists at requested location | boolean | not measured | false | not measured | true before analysis | not measured | `E0-P0A-TARGET1` |
| P0 | Corrected owner target path exists | boolean | false (initial path interpretation) | true (`E:\cheapapp.org`) | not measured | true | not measured | `E0-P0A-R2-OWNER1` |
| P0 | Alternate similarly named directories at E:\\ root | directories | not measured | 1 (`cheapapp.org`) | not measured | informational only | not measured | `E0-P0A-TARGET1` |
| P0 | Target tracked/file inventory | files | not measured | not measured (target absent) | not measured | capture exact value | not measured | `E0-P0A-TARGET1` |
| P0 | Fresh analyze total duration | seconds | not measured | 42.1059032 | not measured | record only | not measured | `E1-P1A-GRAPH2` |
| P0 | Graph node inventory | nodes | not measured | 84,807 | not measured | record only | not measured | `E1-P1A-GRAPH2` |
| P0 | Graph relationship inventory | relationships | not measured | 114,125 | not measured | record only | not measured | `E1-P1A-GRAPH2` |
| P0 | Target tracked-status delta after analysis | paths | 0 expected | 0 | not measured | 0 | 0 | `E0-P0A-R2-BOUNDARY1` |
| P0 | Ignored generated guidance files timestamp-touched by later analyze | files | not measured | 2 (`AGENTS.md`, `CLAUDE.md`) | not measured | informational | not measured | P1-B boundary observation; content delta unavailable |

## B1 - P1 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P1 | Source File paths compared | paths | not measured | not measured | not measured | complete bidirectional set | not measured | `E1-P1A-CMP1` |
| P1 | Source TS/JS paths compared | paths | not measured | 895 | not measured | bounded exact set | not measured | `E1-P1A-FILE1` |
| P1 | Graph TS/JS File nodes | nodes | not measured | 887 | not measured | bounded exact set | not measured | `E1-P1A-FILE1` |
| P1 | Missing TS/JS File nodes | paths | not measured | 8 | not measured | classify each | not measured | `E1-P1A-CMP1` |
| P1 | Graph-only TS/JS File nodes | paths | not measured | 0 | not measured | classify each | not measured | `E1-P1A-CMP1` |
| P1 | Missing/extra File nodes | paths | not measured | 8 missing / 0 graph-only | not measured | classify every case | bounded discrepancy | `E1-P1A-CMP1` |
| P1-B | Source declaration/binding sites compared | sites | 0 | 29 bounded facts (6 pattern bindings + 4 locals + 21 exports) | not measured | exact bounded set | +29 | `E1-P1B-CMP1` |
| P1-B | Destructured bindings present in source / graph | bindings | 6 / 0 | 6 / 0 | not measured | 6 / 6 | graph deficit | `E1-P1B-CMP1` |
| P1-B | Same-name `time`/`now` source declarations / graph nodes | locals | 4 / 2 | 4 / 2 | not measured | preserve 4 / 4 | graph deficit | `E1-P1B-CMP1` |
| P1-B | Source exports / graph definitions with export metadata | exports | 21 / 0 | 21 / 0 | not measured | 21 / 21 | metadata deficit | `E1-P1B-CMP1` |
| P1-C | Resolution/module sites compared | sites | 0 | 5 | not measured | exact bounded set | +5 | `E1-P1C-CMP1` |
| P1-C | Compiler-bound fixed sites | sites | 0 | 5 | not measured | 5 | +5 | `E1-P1C-ORACLE1`, `E1-P1C-ORACLE2` |
| P1-C | Compiler diagnostics at fixed sites | diagnostics | 0 | 0 | not measured | 0 | 0 | `E1-P1C-ORACLE1`, `E1-P1C-ORACLE2` |
| P1-C | Fixed sites exposed unresolved by Anvien | sites | 0 | 5 | not measured | 0 | +5 wrong | `E1-P1C-ANVIEN1`, `E1-P1C-ANVIEN2` |
| P1-C | Expected barrel-call `CALLS` edges | edges | 0 | 2 expected / 0 observed | not measured | 2 observed | -2 | `E1-P1C-CMP1` |
| P1-D | Raw/projection canonical facts compared | facts | 0 | 7 | not measured | 7 | +7 | `E1-P1D-CMP1` |
| P1-D | Exact duplicate raw nodes/edges in selected facts | facts | 0 | 0 | not measured | 0 | 0 | `E1-P1D-RAW1` |
| P1-D | Materialized SourceSite nodes in selected facts | nodes | 0 | 0 | not measured | 0 | 0 | `E1-P1D-RAW1` |
| P1-E | Command/derived cases compared | cases | 0 | 3 | not measured | exact bounded set | +3 | `E1-P1E-CMP1` |
| P1-E | Independent wrapper call sites versus symbol-level context callers | sites | 2 | 0 direct calls observed | not measured | reconcile | -2 | `E1-P1E-SRC1`, `E1-P1E-CMD1` |

## B2 - P2 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P2 | Scanner/extractor/identity causal slices | slices | 0 | 1 bounded P2-A slice (2 causal reports) | not measured | 1 bounded slice | +1 | `E2-P2A-REPORT1`, `E2-P2A-REPORT2`, `E2-P2A-REVIEW1` |
| P2-A | Missing TS/JS paths explained by scanner pruning | paths | 0 | 8 | not measured | explain all bounded paths | +8 | `E2-P2A-CAUSE1` |
| P2-A | Exact missing paths reproduced as matcher ignored | paths | 0 | 8 | not measured | 8 | +8 | `E2-P2A-CAUSE1` |
| P2-A | WalkRepositoryPaths affected processes | processes | not measured | 39 | not measured | record only | not measured | `E2-P2A-IMPACT2` |
| P2-A | Selected array binding definitions in source / ScopeIR | bindings | 6 / 0 | 6 / 0 | not measured | preserve 6 / 6 | -6 in IR | `E2-P2A-TSIR1` |
| P2-A | Selected same-name local definitions in ScopeIR / graph | definitions | 4 / 2 | 4 / 2 | not measured | preserve 4 / 4 | -2 in graph | `E2-P2A-TSIR1`, `E2-P2A-IDENTITY1` |
| P2-A | Selected-file DefinitionFacts with non-empty visibility | definitions | 0 | 0 | not measured | record bounded defect | 0 | `E2-P2A-VIS1` |
| P2-A | `graphIDForDef` affected modules / processes | modules / processes | not measured | 3 / 18 | not measured | record only | not measured | `E2-P2A-IMPACT3` |
| P2-A | `Graph.AddNode` direct symbols / modules / processes | symbols / modules / processes | not measured | 16 / 6 / 82 | not measured | record only | not measured | `E2-P2A-IMPACT3` |
| P2 | Resolver/module causal slices | slices | 0 | 1 bounded P2-B slice | not measured | 1 bounded slice | +1 | `E2-P2B-REPORT1`, `E2-P2B-REVIEW1` |
| P2-B | Fixed resolver/module sites independently probed | sites | 0 | 5 | not measured | 5 | +5 | `E2-P2B-ORACLE1`, `E2-P2B-ORACLE2`, `E2-P2B-TARGET1` |
| P2-B | Actual barrel-chain imports resolved | imports | 0 | 10 | not measured | record only | not measured | `E2-P2B-RESOLVE1` |
| P2-B | Actual barrel-chain import-use edges emitted | edges | 0 | 1 | not measured | record only | not measured | `E2-P2B-RESOLVE1` |
| P2-B | Direct-import control import-use edges emitted | edges | 0 | 3 | not measured | record only | not measured | `E2-P2B-RESOLVE1` |
| P2-B | Actual vs direct-control resolved calls | calls | 0 | 37 vs 39 | not measured | record only | +2 control delta | `E2-P2B-RESOLVE1` |
| P2-B | Actual vs direct-control unresolved references | references | 0 | 542 vs 540 | not measured | record only | -2 control delta | `E2-P2B-RESOLVE1` |
| P2-B | Target graph `.d.ts` File nodes | nodes | 0 | 0 | not measured | record only | 0 | `E2-P2B-BOUNDARY1`, `E2-P2B-CAUSE-A1` |
| P2 | Projection/consumer causal slices | slices | 0 | 1 bounded P2-C slice | not measured | 1 bounded slice | +1 | `E2-P2C-REPORT1`, `E2-P2C-REVIEW2` |
| P2-C | Selected canonical raw/Cypher/file-detail facts | facts | 0 | 7 / 7 / 7 | not measured | preserve exact cardinality | 0 missing / 0 duplicate | `E2-P2C-PROJ1` |
| P2-C | Target Process nodes / step edges | nodes / edges | 0 | 662 / 2,761 | not measured | record only | not measured | `E2-P2C-TARGET1` |
| P2-C | Selected symbols/files with raw process membership | symbols/files | 0 | 0 / 5 | not measured | bounded observation only | 0 | `E2-P2C-TARGET1` |
| P2-C | Actual vs +2-call control call edges considered | calls | 0 | 3,771 vs 3,773 | not measured | record only | +2 control delta | `E2-P2C-DERIVED1` |
| P2-C | Actual vs +2-call control processes emitted | processes | 0 | 662 vs 662 | not measured | record only | 0 control delta | `E2-P2C-DERIVED1` |
| P2-C | Actual vs +2-call control step edges emitted | edges | 0 | 2,761 vs 2,761 | not measured | record only | 0 control delta | `E2-P2C-DERIVED1` |
| P2-C | Selected raw edges without `step` / Cypher rows with `step=0` | edges / rows | 0 | 7 / 7 | not measured | preserve representation distinction | nil-to-zero | `E2-P2C-PROJ1`, `E2-P2C-SRC2` |
| P2-C | Owner symbols with CRITICAL impact warning | symbols | 0 | 7 | not measured | record only | +7 | `E2-P2C-IMPACT1` |

## B3 - P3 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P3 | Causal matrix rows | rows | 0 | 10 bounded rows (C1-C10) | not measured | every included row evidence-backed | +10 | `E3-P3A-SYN1` |
| P3 | Explicit unresolved-boundary ledger | ledgers | 0 | 1 | not measured | retain all unresolved/authority-blocked classes | +1 | `E3-P3A-SYN2` |

## Non-Benchmarkable Notes

- No speed or optimization conclusion is allowed from these measurements.
- UI, browser, Docker, package, and build metrics are not applicable to the read-only source/graph investigation unless a later approved slice adds that boundary.
- No value marked `not measured` may be used as evidence until the corresponding command or source comparison is run and recorded.
- `Latest` is the authoritative measurement for this accepted-bounded read-only investigation. `Final` remains `not measured` where no implementation before/after comparison was authorized; this is an intentional non-remediation boundary, not missing evidence for the bounded causal rows.
