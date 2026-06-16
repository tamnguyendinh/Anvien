# Supervisor Report: File Detail Absolute Path Resolution

Verdict: PASS

## Metadata

- Report file: `reports/Supervisor/rp_supervisor_260616_140659_by_gpt-5-codex_file-detail-absolute-path-resolution.md`
- Review time: `260616 140659 +07:00`
- Reviewer: `gpt-5-codex`
- Repo/project: `Anvien`
- Scope reviewed: plan `docs/plans/2026-06-16-file-detail-absolute-path-resolution`, coder report `reports/coder/rp_coder_260616_140522_by_gpt-5-codex_file-detail-absolute-path-resolution.md`, and implementation commits `9a27930`, `8499041`, `783c809`, `09aa58c`
- Claim reviewed: file-detail/file-context surfaces resolve absolute in-repo file paths against repo-relative graph keys and reject outside-repo absolute paths clearly.
- Authority used: latest user request, `AGENTS.md`, plan acceptance criteria, source code, tests, manual command output, Anvien analyze/detect evidence.
- Related artifacts: coder report, plan/evidence/benchmark/status ledgers, implementation commits listed above.

## Executive Summary

- Problem: graph file nodes are repo-relative, but command surfaces previously passed absolute user input directly to `BuildFileContext`, causing newly scanned files to be unresolved by absolute path.
- Decision: PASS. Source inspection and behavior evidence prove the shared normalization invariant is closed across helper, CLI, HTTP, and MCP file target surfaces.
- Required outcome: accepted; proceed with final plan closure.

## Source-Level Clearance Notes

- `internal/filecontext/context.go`: clear. `NormalizeRepoFilePath` preserves relative paths, strips the selected repo root only for absolute in-repo inputs, and returns `ErrFilePathOutsideRepo` for absolute paths outside the repo.
- `internal/cli/file_detail_command.go`: clear. CLI normalizes `args[0]` with resolved `inputs.RepoPath` before `BuildFileContext` and preserves original input in `context.Target.Input`.
- `internal/httpapi/file_context.go`: clear. HTTP `/api/file-detail` normalizes query `path` against `projection.repoPath`, returns HTTP 400 for outside-repo absolute paths, and keeps success payloads repo-relative.
- `internal/mcp/target_dispatch.go`: clear. `mcpBuildRepoFileContext` is the repo-aware MCP helper; existing `mcpBuildFileContext` remains as the compatibility path for internal graph-relative callers.
- `internal/mcp/context.go` and `internal/mcp/impact.go`: clear. MCP `context` and `impact` file targets pass the resolved repo path into file context lookup, return explicit payload errors for outside-repo absolute paths, and leave symbol/route/tool dispatch intact.

## Evidence Checked

Passed:

- `go build ./cmd/... ./internal/...`: pass.
- Focused tests: `go test ./internal/filecontext -count=1`, `go test ./internal/cli -run TestFileDetailCommand -count=1`, `go test ./internal/httpapi -run TestFileDetailEndpoint -count=1`, and `go test ./internal/mcp -run TestServeCallToolContextAndImpactResolveAbsoluteFileTargets -count=1` all passed.
- Applicable suite: `go test ./cmd/... ./internal/... -count=1` passed.
- Manual CLI smoke: relative and absolute `file-detail` for `internal/mcp/tools.go` both returned `target.normalizedPath="internal/mcp/tools.go"` and `summary.symbolCount=169`.
- Manual outside-repo smoke: absolute outside-repo `file-detail` exited `1` with `file path is outside repository`.
- Final graph refresh: `anvien analyze --force` passed with `1422` files, `82738` nodes, `120864` relationships, `16508` dependency edges, and `419` unresolved.
- Final P2 detect before report/doc edits: `anvien detect-changes --repo Anvien --scope all` reported `changed_count=0`, `affected_count=0`, `risk_level=none`.

Failed:

- `go build ./...` fails on existing intentionally non-buildable fixture packages under `anvien/test/fixtures` (`models`, `animal`, mixed packages, and C `simple.c`). This blocker is known from prior repo plans and is not caused by this scope.

Not run:

- Browser/UI validation: not applicable; no UI behavior changed.
- Live HTTP server smoke: not run; route-level HTTP endpoint tests cover the changed API boundary.
- Live MCP process smoke: not run; MCP JSON-RPC `Serve` tests cover the changed tool boundary.

## Invariant Closure

- Affected invariant: file target lookup must convert absolute in-repo user input to repo-relative graph lookup paths without changing graph storage.
- Sibling surfaces checked: shared helper, CLI file-detail, HTTP file-detail, MCP context file target, MCP impact file target, MCP auto file candidate path, focused tests, manual CLI behavior, final graph/detect evidence.
- Forbidden fallback status: no fuzzy matching, basename guessing, repo guessing, or absolute graph storage was introduced.
- Residual unverified same-invariant surfaces: none for the reviewed scope.

## Overall Evaluation

The implementation is narrow, uses one shared normalization rule, preserves repo-relative graph keys, and covers the command/API/MCP boundaries where the bug could occur. The known repo-wide build caveat is unrelated fixture input and does not block acceptance of this scope because applicable product build/test validation passed.
