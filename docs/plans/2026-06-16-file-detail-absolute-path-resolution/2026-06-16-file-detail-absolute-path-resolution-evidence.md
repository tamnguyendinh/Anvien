# File Detail Absolute Path Resolution Evidence Ledger

## Metadata

- Date: `2026-06-16`
- Plan: `docs/plans/2026-06-16-file-detail-absolute-path-resolution/2026-06-16-file-detail-absolute-path-resolution-plan.md`
- Evidence: `docs/plans/2026-06-16-file-detail-absolute-path-resolution/2026-06-16-file-detail-absolute-path-resolution-evidence.md`
- Benchmark: `docs/plans/2026-06-16-file-detail-absolute-path-resolution/2026-06-16-file-detail-absolute-path-resolution-benchmark.md`
- Actual status: `docs/plans/2026-06-16-file-detail-absolute-path-resolution/2026-06-16-file-detail-absolute-path-resolution-actual-status.md`

## Evidence Rules

- Evidence IDs use `E<phase>-<item>-<kind><n>`.
- Record exact commands, source locations, and outcomes.
- Keep command pass/fail evidence here; measured counts belong in the benchmark ledger.
- Add implementation, validation, detect-changes, commit, and supervisor evidence as each slice completes.

## E0 - P0 Evidence

Matching plan item(s): `P0-A`

- `E0-P0A-ANALYZE1`: `anvien analyze --force` from `E:\Anvien` succeeded. Output included `files: scanned=1418 parsed_code=673 failed=0`, `graph: nodes=82576 relationships=120656 path=E:\Anvien\.anvien\graph.json`, and `fileProjection: status=built files=1418 dependencyEdges=16484 unresolved=419`.
- `E0-P0A-REPRO1`: `anvien file-detail internal/mcp/tools.go --repo Anvien --json --relationships 1 --unresolved 1 --linked 1 | ConvertFrom-Json` succeeded. Compact output: `Repo=Anvien`, `Input=internal/mcp/tools.go`, `NormalizedPath=internal/mcp/tools.go`, `Symbols=169`, `Risk=high`, `Stale=False`.
- `E0-P0A-REPRO2`: `anvien file-detail E:\Anvien\internal\mcp\tools.go --repo Anvien --json --relationships 1 --unresolved 1 --linked 1` failed with `file "E:\\Anvien\\internal\\mcp\\tools.go" not found in repo Anvien`.
- `E0-P0A-CYPHER1`: `anvien cypher "MATCH (f:File) WHERE f.filePath = 'internal/mcp/tools.go' RETURN f.filePath AS path LIMIT 1" --repo Anvien` returned one row: `internal/mcp/tools.go`.
- `E0-P0A-CYPHER2`: `anvien cypher "MATCH (f:File) WHERE f.filePath = 'E:/Anvien/internal/mcp/tools.go' RETURN f.filePath AS path LIMIT 1" --repo Anvien` returned `_No rows_`.
- `E0-P0A-SRC1`: `rg -n "BuildFileContext\(args\[0\]|func loadFileProjectionGraph|file .*not found in repo" internal/cli/file_detail_command.go` showed CLI passes `args[0]` directly to `BuildFileContext` at line 61 and reports missing file at line 67. `loadFileProjectionGraph` starts at line 167 and already has resolved repo data available.
- `E0-P0A-SRC2`: `rg -n "func \(b \*Builder\) BuildFileContext|normalizedPath := normalizePath\(path\)|b\.filesByPath\[normalizedPath\]|func normalizePath" internal/filecontext/context.go` showed `BuildFileContext` at line 375, path normalization at line 377, lookup against `filesByPath` at line 378, and `normalizePath` at line 1118.
- `E0-P0A-SRC3`: `rg -n "path := strings.TrimSpace|BuildFileContext\(path" internal/httpapi/file_context.go` showed HTTP reads query `path` at line 91 and passes it directly to `BuildFileContext` at line 102.
- `E0-P0A-SRC4`: `rg -n "func mcpBuildFileContext|BuildFileContext\(path" internal/mcp/target_dispatch.go internal/mcp/context.go` showed MCP file context dispatch calls `BuildFileContext(path, ...)` at `internal/mcp/target_dispatch.go:56`.
- `E0-P0A-FD1`: `anvien file-detail internal/filecontext/context.go --repo Anvien --json` compact summary: symbols 414, local 555, inbound 239, outbound 88, unresolved 529, linked flows 50, linked tests 7, risk high, stale false.
- `E0-P0A-FD2`: `anvien file-detail internal/cli/file_detail_command.go --repo Anvien --json` compact summary: symbols 77, local 54, inbound 9, outbound 66, unresolved 195, linked flows 10, linked tests 1, risk high, stale false.
- `E0-P0A-FD3`: `anvien file-detail internal/httpapi/file_context.go --repo Anvien --json` compact summary: symbols 70, local 55, inbound 10, outbound 37, unresolved 166, linked flows 7, linked tests 3, risk high, stale false.
- `E0-P0A-FD4`: `anvien file-detail internal/mcp/target_dispatch.go --repo Anvien --json` compact summary: symbols 62, local 7, inbound 51, outbound 76, unresolved 101, linked flows 1, linked tests 1, risk high, stale false.
- `E0-P0A-FD5`: `anvien file-detail internal/mcp/context.go --repo Anvien --json` compact summary: symbols 103, local 48, inbound 48, outbound 111, unresolved 253, linked flows 41, linked tests 2, risk high, stale false.
- `E0-P0A-FD6`: `anvien file-detail internal/filecontext/context_test.go --repo Anvien --json` compact summary: symbols 90, local 40, inbound 0, outbound 88, unresolved 0, linked flows 0, linked tests 0, risk low, stale false.
- `E0-P0A-FD7`: `anvien file-detail internal/cli/file_detail_command_test.go --repo Anvien --json` compact summary: symbols 19, local 3, inbound 2, outbound 48, unresolved 0, linked flows 0, linked tests 1, risk low, stale false.
- `E0-P0A-FD8`: `anvien file-detail internal/httpapi/file_context_test.go --repo Anvien --json` compact summary: symbols 17, local 3, inbound 0, outbound 59, unresolved 0, linked flows 0, linked tests 0, risk low, stale false.
- `E0-P0A-TEST1`: `go test ./internal/filecontext -run TestBuildFileContext -count=1` passed: `ok github.com/tamnguyendinh/anvien/internal/filecontext 0.034s`.
- `E0-P0A-TEST2`: `go test ./internal/cli -run TestFileDetailCommand -count=1` passed: `ok github.com/tamnguyendinh/anvien/internal/cli 1.461s`.
- `E0-P0A-TEST3`: `go test ./internal/httpapi -run TestFileDetailEndpoint -count=1` passed: `ok github.com/tamnguyendinh/anvien/internal/httpapi 0.249s`.

## E1 - P1 Evidence

Matching plan item(s): `P1-A`, `P1-B`, `P1-C`, `P1-D`

### P1-A - Shared repo-aware file lookup normalization

- `E1-P1A-IMPACT1`: `anvien impact file internal/filecontext/context.go --repo Anvien --direction upstream` completed before editing. Blast radius was CRITICAL/HIGH at file level, including CLI, HTTP, MCP, web contract, and FileDetail/FileMap consumers. The implementation was intentionally limited to a pure helper and tests. `anvien impact file internal/filecontext/context_test.go --repo Anvien --direction upstream` reported LOW test-file risk.
- `E1-P1A-SRC1`: Added `ErrFilePathOutsideRepo` at `internal/filecontext/context.go:21` and `NormalizeRepoFilePath(inputPath, repoRoot)` at `internal/filecontext/context.go:1128`. The helper preserves repo-relative inputs, converts in-repo absolute paths to repo-relative lookup paths, and returns a wrapped outside-repo error instead of guessing.
- `E1-P1A-TEST1`: Added `TestNormalizeRepoFilePath` at `internal/filecontext/context_test.go:104`, covering repo-relative, dot-relative, Windows absolute in repo, slash absolute in repo, case-insensitive Windows root, outside-repo absolute, sibling-prefix outside-repo absolute, blank path, and no-repo-root preservation.
- `E1-P1A-TEST2`: `go test ./internal/filecontext -run TestNormalizeRepoFilePath -count=1` passed: `ok github.com/tamnguyendinh/anvien/internal/filecontext 0.021s`.
- `E1-P1A-TEST3`: `go test ./internal/filecontext -count=1` passed: `ok github.com/tamnguyendinh/anvien/internal/filecontext 0.020s`.
- `E1-P1A-ANALYZE1`: `anvien analyze --force` after P1-A implementation succeeded. Output included `files: scanned=1422 parsed_code=673 failed=0`, `graph: nodes=82684 relationships=120762`, and `fileProjection: status=built files=1422 dependencyEdges=16485 unresolved=419`.
- `E1-P1A-DETECT1`: `anvien detect-changes --repo Anvien --scope all` after P1-A implementation succeeded. It reported 5 changed files: the plan, evidence, actual-status docs, `internal/filecontext/context.go` with high file risk, and `internal/filecontext/context_test.go` with low file risk. Summary included `changed_count=48`, `changed_files=5`, `affected_files=5`, `risk_level=low`, and no affected processes.

### P1-B - CLI file-detail wiring

- Pending.

### P1-C - HTTP file-detail wiring

- Pending.

### P1-D - MCP file context dispatch

- Pending.

## E2 - P2 Evidence

Matching plan item(s): `P2-A`, `P2-B`, `P2-C`

- Pending. Record full build, final validation, supervisor review, cleanup, detect-changes, final commits, and worktree status.

## Closure Evidence

- Pending until implementation and validation complete.
