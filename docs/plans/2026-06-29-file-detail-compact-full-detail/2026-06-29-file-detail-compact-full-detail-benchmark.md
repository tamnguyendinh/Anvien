# File Detail Compact Full Detail Benchmark Ledger

## Metadata

- Date: `2026-06-29`
- Plan: `docs/plans/2026-06-29-file-detail-compact-full-detail/2026-06-29-file-detail-compact-full-detail-plan.md`
- Evidence: `docs/plans/2026-06-29-file-detail-compact-full-detail/2026-06-29-file-detail-compact-full-detail-evidence.md`
- Benchmark: `docs/plans/2026-06-29-file-detail-compact-full-detail/2026-06-29-file-detail-compact-full-detail-benchmark.md`
- Actual status: `docs/plans/2026-06-29-file-detail-compact-full-detail/2026-06-29-file-detail-compact-full-detail-actual-status.md`

## Benchmark Rules

The benchmark file records measurements.

It should contain:

- metadata and companion files;
- benchmark rules;
- benchmark sections such as `B0`, `B1`, or sections by phase/task;
- metric tables with unit, baseline, latest, final, target, and delta when needed;
- inventory count;
- runtime or performance metric;
- graph, coverage, or accuracy metric;
- package, bundle, file size, or hash metric;
- before/after numbers;
- UI, layout, or browser metric when the plan involves UI;
- command-surface or generated-output inventory when the plan involves generated documentation.

Benchmark records measured numbers only. Do not put command logs, design decisions, or validation narrative here. Build/test/e2e pass-fail belongs in evidence unless the timing, count, or size is the measured target.

Benchmark sections must follow the plan phases:

- `B0` corresponds to `P0`.
- `B1` corresponds to `P1`.
- `B2` corresponds to `P2`.
- Use item-level IDs such as `B-P1-A` when a checklist item needs separate benchmark evidence.
- Create a benchmark section only when the matching phase has benchmarkable measurements.
- Do not invent fixed metric categories; record the measurements required by the matching plan phase.

## B0 - P0 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P0 | Fresh analyze scanned files | files | 1447 | 1452 | pending | record | +5 | `E0-P0A-GRAPH2` |
| P0 | Fresh analyze graph nodes | nodes | 83186 | 83252 | pending | record | +66 | `E0-P0A-GRAPH2` |
| P0 | Fresh analyze graph relationships | relationships | 121455 | 121521 | pending | record | +66 | `E0-P0A-GRAPH2` |
| P0 | Current expanded file-detail size for `internal/filecontext/context.go` with `--relationships 1 --unresolved 1 --linked 1` | characters | 253168 | 253168 | pending | compact full detail materially smaller | pending | `E0-P0A-FD1` |
| P0 | Current unique related files for `internal/filecontext/context.go` | files | 42 | 42 | pending | compact `relatedFiles` rows = 42 | pending | `E0-P0A-FD1` |
| P0 | Current symbols for `internal/filecontext/context.go` | symbols | 429 | 429 | pending | compact symbol rows preserve 429 symbols | pending | `E0-P0A-FD1` |
| P0 | Current relationship facts for `internal/filecontext/context.go` | relationships | 905 | 905 | pending | compact rows preserve requested relationship facts | pending | `E0-P0A-FD1` |
| P0 | Current unresolved sites for `internal/filecontext/context.go` | sites | 542 | 542 | pending | compact rows preserve unresolved facts | pending | `E0-P0A-FD1` |
| P0 | Current expanded file-detail size for `internal/mcp/tools.go` with `--relationships 1 --unresolved 1 --linked 1` | characters | 135040 | 135040 | pending | compact full detail materially smaller | pending | prior measurement in discussion |
| P0 | Current expanded file-detail size for `internal/aicontext/aicontext.go` with `--relationships 1 --unresolved 1 --linked 1` | characters | 38230 | 38230 | pending | compact full detail materially smaller | pending | prior measurement in discussion |
| P0 | Current expanded file-detail size for `internal/mcp/target_dispatch.go` with `--relationships 1 --unresolved 1 --linked 1` | characters | 40602 | 40602 | pending | MCP surface compatibility decided and tested | pending | `E0-P0A-FD8` |
| P0 | Current unique related files for `internal/mcp/target_dispatch.go` | files | 21 | 21 | pending | MCP plan row preserves or exposes related facts intentionally | pending | `E0-P0A-FD8` |
| P0 | Current expanded file-detail size for `internal/mcp/context.go` with `--relationships 1 --unresolved 1 --linked 1` | characters | 55736 | 55736 | pending | MCP context payload remains compatible or explicit | pending | `E0-P0A-FD9` |
| P0 | Current expanded file-detail size for `internal/mcp/impact.go` with `--relationships 1 --unresolved 1 --linked 1` | characters | 83000 | 83000 | pending | file-impact flow remains compatible or explicit | pending | `E0-P0A-FD10` |

## B1 - P1 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P1 | Compact model size for `internal/filecontext/context.go` versus expanded full-detail equivalent | characters | 849155 | 281704 | pending | materially smaller while preserving facts | -567451 | `E1-P1C-MEASURE1` |
| P1 | Compact related-file rows for `internal/filecontext/context.go` | rows | 0 | 43 | pending | 43 current graph rows | +43 | `E1-P1C-MEASURE1` |
| P1 | Fixture compact related-file rows for `src/app.go` | rows | 0 | 2 | pending | outbound and inbound fixture files represented | +2 | `E1-P1B-TEST1` |
| P1 | Compact symbol row count for `internal/filecontext/context.go` | rows | 0 | 429 | pending | 429 | +429 | `E1-P1C-MEASURE1` |
| P1 | Compact unresolved row count for `internal/filecontext/context.go` | rows | 0 | 542 | pending | 542 | +542 | `E1-P1C-MEASURE1` |
| P1 | Limited compact output omitted-count fields | fields | 0 | 3 | pending | total, returned, and omitted counts represented for limited sections | +3 | `E1-P1C-LIMIT1` |

## B2 - P2 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P2 | CLI compact JSON size for `internal/filecontext/context.go` versus expanded full-detail equivalent | characters | 1282361 | 282588 | pending | materially smaller than expanded equivalent | -999773 | `E2-P2A-SMOKE1` |
| P2 | HTTP compact JSON size for `internal/filecontext/context.go` versus max-limited expanded response | characters | 615355 | 282949 | pending | compact full-detail response smaller than expanded bounded response | -332406 | `E2-P2B-SMOKE1` |
| P2 | Contract file-detail response variants | variants | 1 | pending | pending | compact and expanded represented | pending | `E2-P2C-CONTRACT1` |
| P2 | MCP file-context response behavior variants | variants | 1 | pending | pending | expanded preserved or compact/expanded explicitly represented | pending | `E2-P2D-SRC1` |
| P2 | README/RUNBOOK file-detail docs decision | doc surfaces | 2 | pending | pending | updated or no-change decision recorded for both | pending | `E2-P2E-DOC2` |

## B3 - P3 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P3 | Web file-detail related-file rows rendered in fixture | rows | 0 | pending | pending | fixture-specific expected count | pending | `E3-P3B-TEST1` |
| P3 | Runtime compact endpoint response size for selected smoke file | characters | pending | pending | pending | record final | pending | `E3-P3C-RUNTIME1` |

## Non-Benchmarkable Notes

- Build/test pass-fail belongs in evidence, not this benchmark ledger, unless timing or size is explicitly measured.
- Supervisor review and dead-work cleanup are acceptance gates, not benchmarkable metrics.
