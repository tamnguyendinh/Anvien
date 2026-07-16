# Tree-sitter Go Module Drift Benchmark Ledger

## Metadata

- Date: `2026-07-16`
- Plan: `docs/plans/2026-07-16-tree-sitter-go-module-drift/2026-07-16-tree-sitter-go-module-drift-plan.md`
- Evidence: `docs/plans/2026-07-16-tree-sitter-go-module-drift/2026-07-16-tree-sitter-go-module-drift-evidence.md`
- Benchmark: `docs/plans/2026-07-16-tree-sitter-go-module-drift/2026-07-16-tree-sitter-go-module-drift-benchmark.md`
- Actual status: `docs/plans/2026-07-16-tree-sitter-go-module-drift/2026-07-16-tree-sitter-go-module-drift-actual-status.md`

## Benchmark Rules

The benchmark file records measurements only. Build/test pass-fail belongs in evidence unless timing, count, size, or inventory is the measured target.

## B0 - P0 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P0 | Analyze scanned files | files | N/A | 1464 | TBD | No regression after fix | TBD | `E0-P0A-GRAPH1` |
| P0 | Analyze parsed code files | files | N/A | 676 | TBD | No regression after fix | TBD | `E0-P0A-GRAPH1` |
| P0 | Analyze failed files | files | N/A | 0 | TBD | 0 | TBD | `E0-P0A-GRAPH1` |
| P0 | Graph nodes | nodes | N/A | 84074 | TBD | No unexpected graph loss | TBD | `E0-P0A-GRAPH1` |
| P0 | Graph relationships | relationships | N/A | 122914 | TBD | No unexpected graph loss | TBD | `E0-P0A-GRAPH1` |
| P0 | `go.mod` file-detail related files | files | N/A | 0 | TBD | 0 or explained | TBD | `E0-P0A-FD1` |
| P0 | Direct Tree-sitter Go module inventory in `go.mod` | modules | N/A | 15 | TBD | Inventory preserved or explained | TBD | `E0-P0A-SRC5` |

## B1 - P1 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P1 | Checker report Tree-sitter module rows | rows | TBD | 19 | 19 | Includes all current Tree-sitter Go modules or explains exclusions | TBD | `E1-P1B-RUN1` |
| P1 | Checker actionable drift rows | rows | TBD | 2 | 2 | Report actionable drift explicitly | TBD | `E1-P1B-RUN1` |
| P1 | Checker crash count for missing npm dependencies | crashes | 1 | 0 | 0 | 0 | -1 | `E0-P0A-RUN1`, `E1-P1C-RUN1`, `E3-P3A-RUN2` |

## B2 - P2 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P2 | Dependabot root `gomod` entries | entries | 0 | 1 | 1 | 1 or explicit no-change decision | +1 | `E2-P2B-SRC1`, `E2-P2B-VALIDATE1` |

## B3 - P3 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P3 | Analyze scanned files after fix | files | 1464 | 1468 | 1468 | No unexpected graph loss | +4 plan docs | `E0-P0A-GRAPH1`, `E3-P3A-GRAPH1` |
| P3 | Analyze parsed code files after fix | files | 676 | 676 | 676 | No unexpected parser loss | 0 | `E0-P0A-GRAPH1`, `E3-P3A-GRAPH1` |
| P3 | Analyze failed files after fix | files | 0 | 0 | 0 | 0 | 0 | `E0-P0A-GRAPH1`, `E3-P3A-GRAPH1` |

## Non-Benchmarkable Notes

- Workflow name/text correctness is validation evidence, not a benchmark.
- Full build pass/fail is validation evidence unless a later slice chooses to measure build duration or artifact size.
