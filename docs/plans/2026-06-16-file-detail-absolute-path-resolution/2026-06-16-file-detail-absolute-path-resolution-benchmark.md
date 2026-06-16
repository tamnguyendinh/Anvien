# File Detail Absolute Path Resolution Benchmark Ledger

## Metadata

- Date: `2026-06-16`
- Plan: `docs/plans/2026-06-16-file-detail-absolute-path-resolution/2026-06-16-file-detail-absolute-path-resolution-plan.md`
- Evidence: `docs/plans/2026-06-16-file-detail-absolute-path-resolution/2026-06-16-file-detail-absolute-path-resolution-evidence.md`
- Benchmark: `docs/plans/2026-06-16-file-detail-absolute-path-resolution/2026-06-16-file-detail-absolute-path-resolution-benchmark.md`
- Actual status: `docs/plans/2026-06-16-file-detail-absolute-path-resolution/2026-06-16-file-detail-absolute-path-resolution-actual-status.md`

## Benchmark Rules

- Record measured counts only.
- Build/test pass-fail evidence belongs in the evidence ledger unless timing, size, or count is the target metric.
- Relationship counts here are baseline inventory counts used to size blast radius, not performance targets.

## B0 - P0 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P0 | Scanned files | files | 1418 | 1418 | Pending | No regression | Pending | `E0-P0A-ANALYZE1` |
| P0 | Parsed code files | files | 673 | 673 | Pending | No regression | Pending | `E0-P0A-ANALYZE1` |
| P0 | Analyze failures | files | 0 | 0 | Pending | 0 | Pending | `E0-P0A-ANALYZE1` |
| P0 | Graph nodes | nodes | 82576 | 82576 | Pending | No unexplained drop | Pending | `E0-P0A-ANALYZE1` |
| P0 | Graph relationships | relationships | 120656 | 120656 | Pending | No unexplained drop | Pending | `E0-P0A-ANALYZE1` |
| P0 | File projection files | files | 1418 | 1418 | Pending | No regression | Pending | `E0-P0A-ANALYZE1` |
| P0 | File projection dependency edges | edges | 16484 | 16484 | Pending | No unexplained drop | Pending | `E0-P0A-ANALYZE1` |
| P0 | `internal/filecontext/context.go` relationship count | local+inbound+outbound | 882 | 882 | Pending | Scoped edits only | Pending | `E0-P0A-FD1` |
| P0 | `internal/cli/file_detail_command.go` relationship count | local+inbound+outbound | 129 | 129 | Pending | Scoped edits only | Pending | `E0-P0A-FD2` |
| P0 | `internal/httpapi/file_context.go` relationship count | local+inbound+outbound | 102 | 102 | Pending | Scoped edits only | Pending | `E0-P0A-FD3` |
| P0 | `internal/mcp/target_dispatch.go` relationship count | local+inbound+outbound | 134 | 134 | Pending | Scoped edits only | Pending | `E0-P0A-FD4` |
| P0 | `internal/mcp/context.go` relationship count | local+inbound+outbound | 207 | 207 | Pending | Scoped edits only | Pending | `E0-P0A-FD5` |
| P0 | `internal/filecontext/context_test.go` relationship count | local+inbound+outbound | 128 | 128 | Pending | Focused coverage | Pending | `E0-P0A-FD6` |
| P0 | `internal/cli/file_detail_command_test.go` relationship count | local+inbound+outbound | 53 | 53 | Pending | Focused coverage | Pending | `E0-P0A-FD7` |
| P0 | `internal/httpapi/file_context_test.go` relationship count | local+inbound+outbound | 62 | 62 | Pending | Focused coverage | Pending | `E0-P0A-FD8` |

## B1 - P1 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P1-A | Scanned files after helper | files | 1422 | 1422 | Pending | No regression | 0 from P0 post-plan baseline | `E1-P1A-ANALYZE1` |
| P1-A | Parsed code files after helper | files | 673 | 673 | Pending | No regression | 0 | `E1-P1A-ANALYZE1` |
| P1-A | Analyze failures after helper | files | 0 | 0 | Pending | 0 | 0 | `E1-P1A-ANALYZE1` |
| P1-A | Graph nodes after helper | nodes | 82623 | 82684 | Pending | Expected source/test symbol increase only | +61 | `E1-P1A-ANALYZE1` |
| P1-A | Graph relationships after helper | relationships | 120703 | 120762 | Pending | Expected source/test relationship increase only | +59 | `E1-P1A-ANALYZE1` |
| P1-A | File projection dependency edges after helper | edges | 16484 | 16485 | Pending | No unexplained drop | +1 | `E1-P1A-ANALYZE1` |
| P1-B | Scanned files after CLI wiring | files | 1422 | 1422 | Pending | No regression | 0 from P1-A | `E1-P1B-ANALYZE1` |
| P1-B | Parsed code files after CLI wiring | files | 673 | 673 | Pending | No regression | 0 | `E1-P1B-ANALYZE1` |
| P1-B | Analyze failures after CLI wiring | files | 0 | 0 | Pending | 0 | 0 | `E1-P1B-ANALYZE1` |
| P1-B | Graph nodes after CLI wiring | nodes | 82684 | 82682 | Pending | No unexplained drop | -2 | `E1-P1B-ANALYZE1` |
| P1-B | Graph relationships after CLI wiring | relationships | 120762 | 120782 | Pending | Expected CLI/test relationship update only | +20 | `E1-P1B-ANALYZE1` |
| P1-B | File projection dependency edges after CLI wiring | edges | 16485 | 16492 | Pending | No unexplained drop | +7 | `E1-P1B-ANALYZE1` |
| P1-C | Scanned files after HTTP wiring | files | 1422 | 1422 | Pending | No regression | 0 from P1-B | `E1-P1C-ANALYZE1` |
| P1-C | Parsed code files after HTTP wiring | files | 673 | 673 | Pending | No regression | 0 | `E1-P1C-ANALYZE1` |
| P1-C | Analyze failures after HTTP wiring | files | 0 | 0 | Pending | 0 | 0 | `E1-P1C-ANALYZE1` |
| P1-C | Graph nodes after HTTP wiring | nodes | 82682 | 82694 | Pending | Expected HTTP/test symbol increase only | +12 | `E1-P1C-ANALYZE1` |
| P1-C | Graph relationships after HTTP wiring | relationships | 120782 | 120808 | Pending | Expected HTTP/test relationship update only | +26 | `E1-P1C-ANALYZE1` |
| P1-C | File projection dependency edges after HTTP wiring | edges | 16492 | 16498 | Pending | No unexplained drop | +6 | `E1-P1C-ANALYZE1` |
| P1-D | Scanned files after MCP wiring | files | 1422 | 1422 | Pending | No regression | 0 from P1-C | `E1-P1D-ANALYZE1` |
| P1-D | Parsed code files after MCP wiring | files | 673 | 673 | Pending | No regression | 0 | `E1-P1D-ANALYZE1` |
| P1-D | Analyze failures after MCP wiring | files | 0 | 0 | Pending | 0 | 0 | `E1-P1D-ANALYZE1` |
| P1-D | Graph nodes after MCP wiring | nodes | 82694 | 82738 | Pending | Expected MCP/test symbol increase only | +44 | `E1-P1D-ANALYZE1` |
| P1-D | Graph relationships after MCP wiring | relationships | 120808 | 120864 | Pending | Expected MCP/test relationship update only | +56 | `E1-P1D-ANALYZE1` |
| P1-D | File projection dependency edges after MCP wiring | edges | 16498 | 16508 | Pending | No unexplained drop | +10 | `E1-P1D-ANALYZE1` |

## B2 - P2 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P2-A | Validation scanned files before closure reports | files | 1422 | 1422 | 1422 | No regression | 0 from P1-D | `E2-P2A-ANALYZE1` |
| P2-A | Validation parsed code files before closure reports | files | 673 | 673 | 673 | No regression | 0 | `E2-P2A-ANALYZE1` |
| P2-A | Validation analyze failures before closure reports | files | 0 | 0 | 0 | 0 | 0 | `E2-P2A-ANALYZE1` |
| P2-A | Validation graph nodes before closure reports | nodes | 82738 | 82738 | 82738 | No unexplained drop | 0 from P1-D | `E2-P2A-ANALYZE1` |
| P2-A | Validation graph relationships before closure reports | relationships | 120864 | 120864 | 120864 | No unexplained drop | 0 | `E2-P2A-ANALYZE1` |
| P2-A | Validation file projection dependency edges before closure reports | edges | 16508 | 16508 | 16508 | No unexplained drop | 0 | `E2-P2A-ANALYZE1` |
| P2-C | Closure scanned files after reports | files | 1422 | 1425 | 1425 | Docs/report increase only | +3 from P2-A | `E2-P2C-ANALYZE1` |
| P2-C | Closure parsed code files after reports | files | 673 | 673 | 673 | No code-count regression | 0 | `E2-P2C-ANALYZE1` |
| P2-C | Closure analyze failures after reports | files | 0 | 0 | 0 | 0 | 0 | `E2-P2C-ANALYZE1` |
| P2-C | Closure graph nodes after reports | nodes | 82738 | 82760 | 82760 | Docs/report increase only | +22 from P2-A | `E2-P2C-ANALYZE1` |
| P2-C | Closure graph relationships after reports | relationships | 120864 | 120886 | 120886 | Docs/report increase only | +22 from P2-A | `E2-P2C-ANALYZE1` |
| P2-C | Closure file projection dependency edges after reports | edges | 16508 | 16508 | 16508 | No code dependency-edge change | 0 | `E2-P2C-ANALYZE1` |

## Non-Benchmarkable Notes

- Focused test pass/fail and full build pass/fail are validation evidence, not benchmark data for this plan.
- The core bug is path resolution correctness, not performance.
