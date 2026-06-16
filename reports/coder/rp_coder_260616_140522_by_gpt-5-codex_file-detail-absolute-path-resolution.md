# Coder Report: File Detail Absolute Path Resolution

Status: READY FOR SUPERVISOR REVIEW

## Metadata

- Report file: `reports/coder/rp_coder_260616_140522_by_gpt-5-codex_file-detail-absolute-path-resolution.md`
- Time: `260616 140522 +07:00`
- Coder: `gpt-5-codex`
- Scope: implement `docs/plans/2026-06-16-file-detail-absolute-path-resolution/2026-06-16-file-detail-absolute-path-resolution-plan.md`
- Authority: user request, `AGENTS.md`, `.agents/skills/coder/SKILL.md`, plan/evidence/status ledgers
- Implementation checkpoints: `9a27930`, `8499041`, `783c809`, `09aa58c`

## Invariant Family Map

- Family name: repo-aware file target lookup for file-detail/file-context command surfaces.
- SSOT / authority source: graph file nodes remain repo-relative; command/API/MCP boundaries normalize user input against the resolved repo root.
- Sibling runtime surfaces checked: shared `filecontext.NormalizeRepoFilePath`, CLI `file-detail`, HTTP `/api/file-detail`, MCP `context` file target, MCP `impact` file target, MCP auto file candidate path.
- Forbidden fallback / alternate path: no fuzzy matching, basename guessing, repo guessing, or graph storage change to absolute paths.
- Stale tests/helpers/plans updated: helper, CLI, HTTP, and MCP focused tests; plan/evidence/benchmark/status ledgers.
- Verify matrix: helper unit tests, CLI command tests and manual binary smoke, HTTP endpoint tests, MCP JSON-RPC tool tests, applicable Go build/test suite, Anvien analyze/detect.

## Files Changed

- `internal/filecontext/context.go` and `internal/filecontext/context_test.go`: shared repo-aware normalization and helper coverage.
- `internal/cli/file_detail_command.go` and `internal/cli/file_detail_command_test.go`: CLI absolute path resolution and outside-repo failure.
- `internal/httpapi/file_context.go` and `internal/httpapi/file_context_test.go`: API absolute path resolution and outside-repo HTTP 400.
- `internal/mcp/target_dispatch.go`, `internal/mcp/context.go`, `internal/mcp/impact.go`, and `internal/mcp/server_test.go`: MCP context/impact file target resolution and explicit outside-repo payload errors.
- `docs/plans/2026-06-16-file-detail-absolute-path-resolution/*`: plan, evidence, benchmark, and actual-status ledgers.
- `reports/coder/rp_coder_260616_140522_by_gpt-5-codex_file-detail-absolute-path-resolution.md`: handoff report.
- `docs/notes_decisions_log/notes_decisions_log_20260616.md`: report link and verification summary.

## Verify Outputs

| Command | Result |
|---|---|
| `go build ./...` | Failed before behavior validation on known non-buildable analyzer fixtures under `anvien/test/fixtures` (`models`, `animal`, mixed package fixture, C `simple.c`). Not caused by this scope. |
| `go build ./cmd/... ./internal/...` | Pass. |
| `go test ./internal/filecontext -count=1` | Pass, `0.047s`. |
| `go test ./internal/cli -run TestFileDetailCommand -count=1` | Pass, `2.008s`. |
| `go test ./internal/httpapi -run TestFileDetailEndpoint -count=1` | Pass, `2.619s`. |
| `go test ./internal/mcp -run TestServeCallToolContextAndImpactResolveAbsoluteFileTargets -count=1` | Pass, `0.058s`. |
| `go test ./cmd/... ./internal/... -count=1` | Pass; applicable Go suite for product packages. |
| Manual CLI relative vs absolute smoke with `.tmp\anvien-p2.exe` | Pass; both inputs normalized to `internal/mcp/tools.go` with `symbolCount=169`. |
| Manual CLI outside-repo absolute smoke | Pass; command exited `1` with `file path is outside repository`. |
| `anvien analyze --force` | Pass after removing transient empty validation directory; final graph `1422` files, `82738` nodes, `120864` relationships, `16508` dependency edges, `419` unresolved. |
| `anvien detect-changes --repo Anvien --scope all` | Pass; final P2 state reported no changes, `risk_level=none`. |

## E2E Flow

- Trigger: user or agent calls `file-detail`, HTTP `/api/file-detail`, or MCP `context`/`impact` with a file path.
- Process: boundary resolves the selected repo, normalizes absolute in-repo input to a repo-relative lookup path, and calls `BuildFileContext` against existing repo-relative graph keys.
- Observable result: absolute in-repo paths resolve to the same file summary as repo-relative paths; absolute outside-repo paths fail clearly without fuzzy fallback.

## Closure Notes

- Residual unverified surfaces: none for helper, CLI, HTTP, and MCP file target lookup.
- Known caveat: repo-wide `go build ./...` remains blocked by intentionally non-buildable fixture input folders; applicable product build/test scope passed.
- Risks/open points: direct Supervisor acceptance is still required before this scope is DONE.
