# File Detail Absolute Path Resolution Actual Status

Title: File Detail Absolute Path Resolution
Date: 2026-06-16
Status: Closed
Companion plan: `docs/plans/2026-06-16-file-detail-absolute-path-resolution/2026-06-16-file-detail-absolute-path-resolution-plan.md`
Companion evidence: `docs/plans/2026-06-16-file-detail-absolute-path-resolution/2026-06-16-file-detail-absolute-path-resolution-evidence.md`
Companion benchmark: `docs/plans/2026-06-16-file-detail-absolute-path-resolution/2026-06-16-file-detail-absolute-path-resolution-benchmark.md`

## Purpose

This file records the real current state before implementation. Implementation must not start until the target scope has a completed status row, evidence IDs, and downstream plan decisions.

## Freshness / Refresh Rules

- Refresh after each completed implementation slice.
- Append to the Status Refresh Log; do not delete P0 baseline rows.
- If status changes alter the next work, update only the affected next-phase work steps before continuing.
- Keep detailed proof in the evidence ledger and store classification here.

## Scope

Target scope:

- CLI `anvien file-detail`.
- HTTP `/api/file-detail`.
- MCP file context dispatch.
- Shared file lookup normalization.
- Focused tests for these surfaces.

Out of scope:

- Analyzer graph schema or graph `filePath` storage.
- Fuzzy matching, basename matching, or repo guessing.
- UI redesign or unrelated Web state changes.

## Relationship / Impact Evidence

Relationship count is `localRelationshipCount + inboundRefCount + outboundRefCount` from `anvien file-detail`.

| Unit / File / Surface | File Detail Evidence | Related File Count | Relationship Summary | Impact Note |
|-----------------------|----------------------|--------------------|----------------------|-------------|
| `internal/filecontext/context.go` | `E0-P0A-FD1` | 882 | local 555, inbound 239, outbound 88, unresolved 529, linked flows 50, tests 7 | high scope warning |
| `internal/cli/file_detail_command.go` | `E0-P0A-FD2` | 129 | local 54, inbound 9, outbound 66, unresolved 195, linked flows 10, tests 1 | high scope warning |
| `internal/httpapi/file_context.go` | `E0-P0A-FD3` | 102 | local 55, inbound 10, outbound 37, unresolved 166, linked flows 7, tests 3 | high scope warning |
| `internal/mcp/target_dispatch.go` | `E0-P0A-FD4` | 134 | local 7, inbound 51, outbound 76, unresolved 101, linked flows 1, tests 1 | high scope warning |
| `internal/mcp/context.go` | `E0-P0A-FD5` | 207 | local 48, inbound 48, outbound 111, unresolved 253, linked flows 41, tests 2 | high scope warning |
| `internal/filecontext/context_test.go` | `E0-P0A-FD6` | 128 | local 40, inbound 0, outbound 88, unresolved 0 | low risk test target |
| `internal/cli/file_detail_command_test.go` | `E0-P0A-FD7` | 53 | local 3, inbound 2, outbound 48, unresolved 0, linked tests 1 | low risk test target |
| `internal/httpapi/file_context_test.go` | `E0-P0A-FD8` | 62 | local 3, inbound 0, outbound 59, unresolved 0 | low risk test target |

## Status Rules

| Status | Meaning | Allowed next action |
|--------|---------|---------------------|
| `correct` | Already behaves as required. | Preserve. |
| `partial` | Some required behavior exists, but gaps remain. | Change only missing parts. |
| `wrong` | Current behavior conflicts with requirement. | Replace with required behavior. |
| `missing` | Required behavior or coverage does not exist. | Implement missing piece only. |
| `blocked` | Required evidence or authority is missing. | Stop until resolved. |

## Current Status Matrix

| Unit | Current State | Required State | Status | Relationship Count | Evidence | Next Plan Decision |
|------|---------------|----------------|--------|--------------------|----------|--------------------|
| Graph file node storage | Graph contains repo-relative `File:internal/mcp/tools.go`; absolute graph path returns no row. | Preserve repo-relative graph storage. | correct | N/A | `E0-P0A-CYPHER1`, `E0-P0A-CYPHER2` | preserve |
| `BuildFileContext` lookup and shared helper | `BuildFileContext` still looks up repo-relative graph paths. Shared `NormalizeRepoFilePath` is now wired at CLI, HTTP, and MCP boundaries. | Keep repo-relative lookup and use the shared helper from command boundaries. | correct | 882 on owner file | `E0-P0A-SRC2`, `E0-P0A-FD1`, `E1-P1A-SRC1`, `E1-P1A-TEST1..E1-P1A-TEST3`, `E1-P1B-SRC1`, `E1-P1C-SRC1`, `E1-P1D-SRC1..E1-P1D-SRC2` | preserve in P2 |
| CLI `file-detail` | Relative and absolute in-repo paths now resolve to the same repo-relative graph file; outside-repo absolute paths fail before lookup. | Normalize `args[0]` against resolved `RepoPath` before lookup; reject outside-repo absolute paths clearly. | correct | 130 after P1-B analyze | `E0-P0A-REPRO1`, `E0-P0A-REPRO2`, `E0-P0A-SRC1`, `E0-P0A-FD2`, `E1-P1B-SRC1`, `E1-P1B-TEST1..E1-P1B-TEST2` | preserve in P1-C/P1-D |
| HTTP `/api/file-detail` | Query `path` now normalizes against `projection.repoPath`; absolute in-repo paths resolve and outside-repo absolute paths return HTTP 400 JSON error. | Normalize query `path` against resolved `projection.repoPath` before lookup. | correct | 103 after P1-C analyze | `E0-P0A-SRC3`, `E0-P0A-FD3`, `E1-P1C-SRC1`, `E1-P1C-TEST1..E1-P1C-TEST2` | preserve in P1-D |
| MCP file context dispatch | MCP `context` and `impact` file targets now normalize against the resolved repo path and reject outside-repo absolute paths with explicit payload errors. Internal repo-relative summary/layer callers still use the compatibility wrapper. | Thread repo root into file context lookup and apply shared helper. | correct | 138 target dispatch, 209 context, 290 impact after P1-D analyze | `E0-P0A-SRC4`, `E0-P0A-FD4`, `E0-P0A-FD5`, `E1-P1D-SRC1..E1-P1D-SRC2`, `E1-P1D-TEST1..E1-P1D-TEST3` | preserve in P2 |
| Existing focused tests | Helper, CLI, HTTP, and MCP now cover absolute in-repo and outside-repo behavior at command/API/MCP boundaries. | Preserve existing tests and add final validation. | correct | 128 helper, 60 CLI test after P1-B, 68 HTTP test after P1-C, 342 MCP test after P1-D | `E0-P0A-TEST1`, `E0-P0A-TEST2`, `E0-P0A-TEST3`, `E0-P0A-FD6..E0-P0A-FD8`, `E1-P1A-TEST1..E1-P1A-TEST3`, `E1-P1B-TEST1..E1-P1B-TEST2`, `E1-P1C-TEST1..E1-P1C-TEST2`, `E1-P1D-TEST1..E1-P1D-TEST3` | proceed to P2 |

## Status Refresh Log

| Refresh | Date | Repo Basis | Changed Scope | Status Changes | Evidence | Next Phase Update |
|---------|------|------------|----------------|----------------|----------|-------------------|
| R0 | 2026-06-16 | baseline before implementation, graph refreshed with `anvien analyze --force` | file-detail absolute path resolution | initial classification: graph storage correct; CLI/API/MCP wrong for absolute path; tests partial | `E0-P0A-ANALYZE1`, `E0-P0A-REPRO1`, `E0-P0A-REPRO2`, `E0-P0A-SRC1..E0-P0A-SRC4`, `E0-P0A-FD1..E0-P0A-FD8`, `E0-P0A-TEST1..E0-P0A-TEST3` | P1 split into shared helper, CLI, HTTP, MCP slices |
| R1 | 2026-06-16 | after P1-A source, tests, graph refresh, and detect-changes | shared file path helper | shared helper capability `missing -> correct`; boundary surfaces remain wrong until wired | `E1-P1A-SRC1`, `E1-P1A-TEST1..E1-P1A-TEST3`, `E1-P1A-ANALYZE1`, `E1-P1A-DETECT1` | proceed to P1-B CLI wiring |
| R2 | 2026-06-16 | after P1-B source, CLI tests, graph refresh, and detect-changes | CLI `file-detail` boundary | CLI row `wrong -> correct`; HTTP and MCP remain wrong until wired | `E1-P1B-SRC1`, `E1-P1B-TEST1..E1-P1B-TEST2`, `E1-P1B-ANALYZE1`, `E1-P1B-DETECT1` | proceed to P1-C HTTP wiring |
| R3 | 2026-06-16 | after P1-C source, HTTP endpoint tests, graph refresh, and detect-changes | HTTP `/api/file-detail` boundary | HTTP row `wrong -> correct`; MCP remains wrong until wired | `E1-P1C-SRC1`, `E1-P1C-TEST1..E1-P1C-TEST2`, `E1-P1C-ANALYZE1`, `E1-P1C-DETECT1` | proceed to P1-D MCP wiring |
| R4 | 2026-06-16 | after P1-D source, MCP boundary tests, graph refresh, and detect-changes | MCP file context dispatch | MCP row `wrong -> correct`; all P1 command surfaces now use shared normalization | `E1-P1D-SRC1..E1-P1D-SRC2`, `E1-P1D-TEST1..E1-P1D-TEST3`, `E1-P1D-ANALYZE1`, `E1-P1D-DETECT1` | proceed to P2 validation |
| R5 | 2026-06-16 | after P2 build/test/manual CLI validation and clean detect-changes | full validation | applicable build/test suite and manual CLI checks pass; repo-wide `go build ./...` remains blocked by known fixture input packages | `E2-P2A-ANALYZE1`, `E2-P2A-BUILD1..E2-P2A-BUILD2`, `E2-P2A-TEST1..E2-P2A-TEST2`, `E2-P2A-CLI1..E2-P2A-CLI2`, `E2-P2A-DETECT1` | proceed to supervisor review |
| R6 | 2026-06-16 | after supervisor review and closure report creation | closure | supervisor PASS; plan ledgers and reports prepared for final closure commit | `E2-P2B-SUPERVISOR1`, `E2-P2C-REPORT1..E2-P2C-STATUS1` | close plan |

## Phase Touch Map

| Unit / File / Surface | Plan-Relevant Relationship File | Relationship to Target | Plan Item | Touch Mode | Evidence | Constraint |
|-----------------------|---------------------------------|------------------------|-----------|------------|----------|------------|
| `internal/filecontext/context.go` | `internal/filecontext/context_test.go` | helper owner and tests | P1-A | edit | `E0-P0A-FD1`, `E0-P0A-FD6` | keep graph lookup repo-relative |
| `internal/cli/file_detail_command.go` | `internal/cli/file_detail_command_test.go` | CLI boundary and tests | P1-B | edit | `E0-P0A-FD2`, `E0-P0A-FD7` | use resolved `RepoPath`, not raw `--repo` |
| `internal/httpapi/file_context.go` | `internal/httpapi/file_context_test.go` | HTTP boundary and tests | P1-C | edit | `E0-P0A-FD3`, `E0-P0A-FD8` | preserve API route and success payload shape |
| `internal/mcp/target_dispatch.go` | `internal/mcp/context.go`, `internal/mcp/impact.go`, `internal/mcp/server_test.go` | MCP helper, context/impact caller surfaces, and tests | P1-D | edit | `E0-P0A-FD4`, `E0-P0A-FD5`, `E1-P1D-IMPACT1` | preserve non-file target dispatch |
| Graph file nodes | `.anvien/graph.json` | source evidence only | P1 all | preserve-only | `E0-P0A-ANALYZE1`, `E0-P0A-CYPHER1` | do not edit generated graph as source of truth |

## Detailed Findings

### Graph File Node Storage

Current state:

Graph contains `internal/mcp/tools.go` as a repo-relative file path. It does not contain `E:/Anvien/internal/mcp/tools.go`.

Required state:

```text
Preserve repo-relative graph file paths. Fix input normalization before lookup.
```

Evidence:

- `E0-P0A-CYPHER1`: repo-relative graph lookup returned one row.
- `E0-P0A-CYPHER2`: absolute graph lookup returned no rows.

Classification:

`correct`

Allowed next action:

Preserve graph storage.

Forbidden next action:

Do not change analyzer output to store absolute paths.

### File Detail Input Normalization

Current state:

`BuildFileContext` only normalizes separators and `./`, then looks up `filesByPath`. Command surfaces pass raw user paths into this lookup.

Required state:

```text
Boundary input should be normalized against the resolved repo root before graph lookup.
Successful lookup paths should remain repo-relative.
Outside-repo absolute paths should fail explicitly.
```

Evidence:

- `E0-P0A-SRC1`: CLI passes `args[0]` directly.
- `E0-P0A-SRC2`: builder normalizes lightly and looks up `filesByPath`.
- `E0-P0A-SRC3`: HTTP route passes query `path` directly.
- `E0-P0A-SRC4`: MCP helper passes `path` directly.
- `E0-P0A-REPRO2`: absolute in-repo path fails even though graph contains the file.

Relationship and impact:

- Related file count: 882 for `internal/filecontext/context.go`, 129 CLI, 102 HTTP, 134/207 MCP.
- Impact note: high scope warning; keep edits narrow and test each surface.

Classification:

`wrong` for CLI, HTTP, and MCP; `partial` for shared lookup support.

Allowed next action:

Implement P1-A through P1-D in order.

Forbidden next action:

Do not add fuzzy matching, basename fallback, or repo guessing.

## Next Phase Status Decisions

| Plan Item | Actual Status Finding | Required Status / Next-Action Update |
|-----------|-----------------------|--------------------------------------|
| P1-A | Shared normalization is partial/missing for repo-aware absolute paths. | Implement helper first so CLI/API/MCP do not duplicate path logic. |
| P1-B | Shared helper is wired into CLI and focused CLI boundary tests pass. | Preserve CLI behavior while wiring HTTP/MCP siblings. |
| P1-C | HTTP endpoint now uses the shared helper and focused endpoint tests pass. | Preserve HTTP behavior while wiring MCP sibling. |
| P1-D | MCP file target dispatch now resolves absolute in-repo paths and rejects outside-repo absolute paths. | Preserve behavior during P2 validation. |
| P2-A | Focused behavior tests now cover helper, CLI, HTTP, and MCP. | Run full build, focused tests, and real CLI boundary commands. |

## Implementation Gate

- [x] Target scope is listed in Current Status Matrix.
- [x] Each target unit has a status.
- [x] Each status has evidence IDs.
- [x] Each target file has relationship count evidence from `file-detail` when applicable.
- [x] Phase Touch Map lists plan-relevant relationship files that can affect the current phase/slice.
- [x] Phase Touch Map defines touch mode for every plan-relevant relationship unit that may be affected.
- [x] Correct parts are marked preserve-only.
- [x] Partial and wrong parts have exact next actions.
- [x] Blockers are recorded, if any.
- [x] Next phase status assumptions, next action, and work steps have been updated from this status file.
- [x] Status Refresh Log has an R0 baseline row.
- [x] If implementation has started, affected rows will be refreshed from latest evidence.
- [x] If refreshed statuses change next work, only stale next-phase status assumptions, next action, or work steps will be updated before the next phase.

## Final P0 Decision

- [ ] P0 actual-status incomplete. Implementation is blocked.
- [ ] P0 complete. Next phase can proceed unchanged.
- [x] P0 complete. Next phase status, next action, or work steps must be updated before implementation.
- [ ] P0 complete. Target scope is preserve-only.
- [ ] P0 complete. Implementation is blocked by missing authority or evidence.

Decision note:

P0 is complete. Implementation should start with P1-A shared helper, then wire CLI, HTTP, and MCP separately. The plan was updated to avoid duplicated path normalization and to preserve repo-relative graph storage.
