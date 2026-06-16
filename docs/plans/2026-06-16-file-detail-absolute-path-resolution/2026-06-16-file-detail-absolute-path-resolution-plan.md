# File Detail Absolute Path Resolution Plan

## Metadata

- Date: `2026-06-16`
- Status: `p1-d-complete`
- Plan: `docs/plans/2026-06-16-file-detail-absolute-path-resolution/2026-06-16-file-detail-absolute-path-resolution-plan.md`
- Evidence: `docs/plans/2026-06-16-file-detail-absolute-path-resolution/2026-06-16-file-detail-absolute-path-resolution-evidence.md`
- Benchmark: `docs/plans/2026-06-16-file-detail-absolute-path-resolution/2026-06-16-file-detail-absolute-path-resolution-benchmark.md`
- Actual status: `docs/plans/2026-06-16-file-detail-absolute-path-resolution/2026-06-16-file-detail-absolute-path-resolution-actual-status.md`

## Goal

Make `anvien file-detail` and matching Web/API/MCP file-detail surfaces resolve file targets consistently when users provide repo-relative paths, `./` paths, Windows absolute paths, or slash-normalized absolute paths that are inside the indexed repo.

## Rules

- Complete P0 actual status before implementation work.
- Update each checklist item immediately when completed.
- Code first, then update tests for the implemented behavior.
- Before editing code, run Anvien impact for the exact target file or symbol and record the evidence ID.
- HIGH or CRITICAL blast radius is a scope warning, not a prohibition.
- After each completed implementation slice, refresh actual status, run `anvien detect-changes --repo Anvien --scope all`, commit the slice, then continue.
- Run full build before validation. For this non-UI scope, full build starts with `go build ./...`; if package or web contract changes are introduced, add the repo-native generated-output validation needed by that change.
- Do not introduce hidden fallback. If an absolute path is outside the selected repo, return a clear error instead of silently guessing.

## Problem

The graph stores file nodes by repo-relative `filePath` values such as `internal/mcp/tools.go`. `file-detail` currently passes the user input path directly to `BuildFileContext`, whose internal `normalizePath` only trims whitespace, changes `\` to `/`, and removes `./`. Therefore an input such as `E:\Anvien\internal\mcp\tools.go` normalizes to `E:/Anvien/internal/mcp/tools.go` and cannot match the graph key `internal/mcp/tools.go`, even though the graph just scanned the file.

## Scope

- CLI `anvien file-detail <path> --repo <repo>`.
- HTTP `/api/file-detail?repo=<repo>&path=<path>`.
- MCP file target paths that call `mcpBuildFileContext`.
- Shared file lookup normalization that can be reused by those surfaces.
- Focused tests for repo-relative, dot-relative, Windows absolute, slash absolute, outside-repo absolute, and missing-file paths.

## Non-Goals

- Do not change graph storage format from repo-relative paths.
- Do not change analyzer scan behavior.
- Do not add fuzzy file matching or basename guessing.
- Do not rename commands or API routes.
- Do not change unrelated file projection summaries, graph health, query, impact, or detect-changes behavior.

## Requirements

- Repo-relative and `./` paths keep current behavior.
- Absolute paths under the selected repo resolve to the same normalized repo-relative graph path.
- Absolute paths outside the selected repo fail clearly at the command/API/MCP boundary.
- Missing repo-relative files still report missing file, not outside-repo.
- The normalized path reported in successful payloads remains repo-relative.
- CLI, HTTP, and MCP use one consistent normalization rule.
- Tests prove behavior at the command/API/MCP boundary, not only helper internals.

## Acceptance Criteria

- `anvien file-detail E:\Anvien\internal\mcp\tools.go --repo Anvien --json` returns the same `target.normalizedPath` and summary as `anvien file-detail internal/mcp/tools.go --repo Anvien --json`.
- HTTP `/api/file-detail` resolves an absolute in-repo path and rejects an absolute outside-repo path with a clear structured error.
- MCP file context resolves an absolute in-repo path or rejects an outside-repo path consistently with CLI/API.
- Focused Go tests cover helper, CLI, HTTP, and MCP behavior.
- Full build and targeted validation pass after implementation.
- `anvien detect-changes --repo Anvien --scope all` is recorded before each implementation commit.

## Checklist

- [x] P0-A: Complete actual status before implementation work.
  - Goal: establish current graph, source, reproduction, coverage, and touch map before implementation.
  - Work Steps: refresh graph, reproduce relative success and absolute failure, inspect owner source lines, collect file-detail relationship counts, run focused baseline tests, and update P1 work from the findings.
  - Implementation Gate: no code edits before `actual-status.md` records the P0 decision.
  - Acceptance: actual status classifies CLI/API/MCP as partial or wrong, identifies the root cause, and names P1 slices.

### P1: Repo-Aware File Target Resolution

- Phase Goal: implement a shared repo-aware file target normalization path and wire it into CLI, HTTP, and MCP file-detail surfaces.
- Phase Boundary:
  - In scope: file lookup normalization and file-detail/context entrypoints.
  - Out of scope: graph schema, analyzer output, unrelated query or impact target dispatch behavior.
  - Dependencies: P0 evidence, impact evidence for each edited file or symbol.
- Phase Implementation Rule: do not implement `P1` directly. Implement slices in order, record evidence, refresh actual status, run detect-changes, and commit each completed slice.
- Ordered Slice List:
  - P1-A: Add shared repo-aware file lookup normalization.
  - P1-B: Wire CLI `file-detail`.
  - P1-C: Wire HTTP `/api/file-detail`.
  - P1-D: Wire MCP file context dispatch.

- [x] P1-A: Add shared repo-aware file lookup normalization.
  - Goal: create one source-of-truth helper that converts user input plus repo root into the repo-relative graph lookup path.
  - Scope Boundary:
    - Editable: `internal/filecontext/context.go`, `internal/filecontext/context_test.go`, or a new narrow file in `internal/filecontext`.
    - Inspect-only: `internal/cli/file_detail_command.go`, `internal/httpapi/file_context.go`, `internal/mcp/target_dispatch.go`.
    - Preserve-only: graph node `filePath` semantics and existing `BuildFileContext` repo-relative lookup behavior.
    - Out of scope: changing analyzer graph output.
  - Non-Goals: do not change `BuildFileContext` payload shape or graph key format.
  - Pre-flight Questions:
    - Data source: graph file nodes keyed by repo-relative `filePath`.
    - Display permission: N/A, non-UI.
    - DB read flow: N/A, graph JSON read only.
    - DB write flow: N/A.
    - Render location: N/A.
    - UI behavior flow: N/A.
    - Docker runtime: N/A for this non-UI slice.
    - Playwright target: N/A.
    - Behavior test: helper unit tests in filecontext.
    - Cleanup/quarantine: no persistent data should be created.
    - External side effects: none.
    - N/A notes: file lookup is local command/API behavior.
  - Work Steps:
    1. Run impact on the filecontext target before editing and record evidence.
       - UI flow check: N/A.
       - DB/data flow check: N/A.
       - Render location check: N/A.
       - Evidence target: `E1-P1A-IMPACT1`.
    2. Implement a helper that accepts `inputPath` and `repoRoot`, normalizes separators, strips the repo root only when the absolute path is inside it, preserves relative paths, and returns an explicit outside-repo error.
       - UI flow check: N/A.
       - DB/data flow check: helper is pure and does not read or write graph data.
       - Render location check: N/A.
       - Evidence target: `E1-P1A-SRC1`.
    3. Add helper tests for repo-relative, dot-relative, Windows absolute in repo, slash absolute in repo, absolute outside repo, and blank path.
       - UI flow check: N/A.
       - DB/data flow check: N/A.
       - Render location check: N/A.
       - Evidence target: `E1-P1A-TEST1`.
  - Implementation Gate:
    - Impact evidence is recorded before editing.
    - The helper API keeps graph lookup path repo-relative.
  - Acceptance:
    - Source: helper has one consistent rule and no fuzzy fallback.
    - Runtime/UI: N/A.
    - DB/data: N/A.
    - Behavior test: filecontext helper tests pass.
    - Cleanup/quarantine: no temp files outside repo.
    - Evidence IDs: `E1-P1A-IMPACT1`, `E1-P1A-SRC1`, `E1-P1A-TEST1`.
    - Actual-status rows refreshed: filecontext rows transition `partial -> correct` for helper capability.
  - Evidence Targets: impact, source diff, focused test.
  - Actual-status Update: refresh filecontext row after tests pass.
  - Commit Boundary: commit after this slice when acceptance passes.

- [x] P1-B: Wire CLI `file-detail`.
  - Goal: make `anvien file-detail <absolute-in-repo-path> --repo Anvien` resolve to the same graph file as the repo-relative path.
  - Scope Boundary:
    - Editable: `internal/cli/file_detail_command.go`, `internal/cli/file_detail_command_test.go`.
    - Inspect-only: `internal/filecontext`.
    - Preserve-only: `file-hotspots` and repo resolution logic except where used to get `RepoPath`.
    - Out of scope: command rename or output contract changes beyond clearer errors.
  - Non-Goals: do not make `--repo` optional when multiple repos are registered.
  - Pre-flight Questions:
    - Data source: `loadFileProjectionGraph` already returns `RepoPath`.
    - Display permission: N/A.
    - DB read flow: graph JSON read only.
    - DB write flow: N/A.
    - Render location: CLI stdout/stderr.
    - UI behavior flow: N/A.
    - Docker runtime: N/A.
    - Playwright target: N/A.
    - Behavior test: CLI tests using fixture repo path.
    - Cleanup/quarantine: test temp dirs only.
    - External side effects: none.
    - N/A notes: command-only behavior.
  - Work Steps:
    1. Run impact on CLI target file before editing and record evidence.
       - UI flow check: N/A.
       - DB/data flow check: N/A.
       - Render location check: CLI command result.
       - Evidence target: `E1-P1B-IMPACT1`.
    2. Normalize `args[0]` with the shared helper using `inputs.RepoPath` before calling `BuildFileContext`.
       - UI flow check: N/A.
       - DB/data flow check: graph lookup uses repo-relative path.
       - Render location check: successful JSON keeps `target.normalizedPath` repo-relative.
       - Evidence target: `E1-P1B-SRC1`.
    3. Add CLI tests for absolute in-repo success and outside-repo failure.
       - UI flow check: N/A.
       - DB/data flow check: N/A.
       - Render location check: CLI output/error contract.
       - Evidence target: `E1-P1B-TEST1`.
  - Implementation Gate:
    - Shared helper from P1-A is complete.
    - Impact evidence is recorded before editing.
  - Acceptance:
    - Source: CLI uses shared helper and does not duplicate path logic.
    - Runtime/UI: N/A.
    - DB/data: N/A.
    - Behavior test: CLI focused tests pass.
    - Cleanup/quarantine: no temp files outside repo.
    - Evidence IDs: `E1-P1B-IMPACT1`, `E1-P1B-SRC1`, `E1-P1B-TEST1`.
    - Actual-status rows refreshed: CLI row transitions `wrong -> correct`.
  - Evidence Targets: impact, source diff, focused CLI test, manual CLI command.
  - Actual-status Update: refresh CLI row.
  - Commit Boundary: commit after this slice when acceptance passes.

- [x] P1-C: Wire HTTP `/api/file-detail`.
  - Goal: make the Web/API runtime resolve absolute in-repo file paths consistently with CLI.
  - Scope Boundary:
    - Editable: `internal/httpapi/file_context.go`, `internal/httpapi/file_context_test.go`.
    - Inspect-only: `anvien-web/src/services/backend-client.ts` if API error handling needs shape confirmation.
    - Preserve-only: `/api/file-hotspots`, cache key behavior, graph loading.
    - Out of scope: frontend visual changes.
  - Non-Goals: do not change API route path or successful JSON contract except `target.input` may remain the original user input.
  - Pre-flight Questions:
    - Data source: `projection.repoPath` and graph JSON.
    - Display permission: N/A.
    - DB read flow: graph JSON read only.
    - DB write flow: N/A.
    - Render location: HTTP JSON response.
    - UI behavior flow: frontend caller passes `path` unchanged.
    - Docker runtime: N/A unless broader runtime validation is required by later changes.
    - Playwright target: N/A.
    - Behavior test: HTTP endpoint tests.
    - Cleanup/quarantine: test temp dirs only.
    - External side effects: none.
    - N/A notes: route-level contract only.
  - Work Steps:
    1. Run impact on HTTP target file before editing and record evidence.
       - UI flow check: N/A.
       - DB/data flow check: graph lookup uses repo-relative path.
       - Render location check: HTTP status and JSON/error body.
       - Evidence target: `E1-P1C-IMPACT1`.
    2. Normalize request `path` against `projection.repoPath` before calling `BuildFileContext`.
       - UI flow check: N/A.
       - DB/data flow check: no graph mutation.
       - Render location check: API response reports repo-relative normalized path.
       - Evidence target: `E1-P1C-SRC1`.
    3. Add endpoint tests for absolute in-repo success and outside-repo failure.
       - UI flow check: N/A.
       - DB/data flow check: N/A.
       - Render location check: route returns expected status.
       - Evidence target: `E1-P1C-TEST1`.
  - Implementation Gate:
    - P1-A helper is complete.
    - Impact evidence is recorded before editing.
  - Acceptance:
    - Source: HTTP route uses shared helper with `projection.repoPath`.
    - Runtime/UI: HTTP endpoint test proves behavior.
    - DB/data: N/A.
    - Behavior test: HTTP focused tests pass.
    - Cleanup/quarantine: no temp files outside repo.
    - Evidence IDs: `E1-P1C-IMPACT1`, `E1-P1C-SRC1`, `E1-P1C-TEST1`.
    - Actual-status rows refreshed: HTTP row transitions `wrong -> correct`.
  - Evidence Targets: impact, source diff, focused HTTP tests.
  - Actual-status Update: refresh HTTP row.
  - Commit Boundary: commit after this slice when acceptance passes.

- [x] P1-D: Wire MCP file context dispatch.
  - Goal: make MCP file target payloads resolve absolute in-repo paths with the same helper and reject outside-repo absolute paths clearly.
  - Scope Boundary:
    - Editable: `internal/mcp/target_dispatch.go`, direct callers that must pass repo path, and focused MCP tests.
    - Inspect-only: `internal/mcp/context.go`, `internal/mcp/impact.go`, route/tool map files that call file target dispatch.
    - Preserve-only: symbol target dispatch, query ranking, process and route payload shapes.
    - Out of scope: changing MCP tool names or target type semantics.
  - Non-Goals: do not add fuzzy file matching for MCP.
  - Pre-flight Questions:
    - Data source: repo entry path from MCP graph resource resolution.
    - Display permission: N/A.
    - DB read flow: graph JSON read only.
    - DB write flow: N/A.
    - Render location: MCP payload JSON.
    - UI behavior flow: N/A.
    - Docker runtime: N/A.
    - Playwright target: N/A.
    - Behavior test: MCP target dispatch/context tests.
    - Cleanup/quarantine: no persistent files.
    - External side effects: none.
    - N/A notes: MCP command surface only.
  - Work Steps:
    1. Run impact on MCP target dispatch and direct callers before editing and record evidence.
       - UI flow check: N/A.
       - DB/data flow check: graph lookup uses repo-relative path.
       - Render location check: MCP payload target normalized path.
       - Evidence target: `E1-P1D-IMPACT1`.
    2. Thread repo root into MCP file context lookup without changing non-file target behavior.
       - UI flow check: N/A.
       - DB/data flow check: no graph mutation.
       - Render location check: file payload stays repo-relative on success.
       - Evidence target: `E1-P1D-SRC1`.
    3. Add MCP tests for absolute in-repo success and outside-repo failure.
       - UI flow check: N/A.
       - DB/data flow check: N/A.
       - Render location check: MCP response has found/not_found or explicit error as designed.
       - Evidence target: `E1-P1D-TEST1`.
  - Implementation Gate:
    - P1-A helper is complete.
    - Impact evidence is recorded before editing.
  - Acceptance:
    - Source: MCP file lookup uses shared helper and preserves non-file target dispatch.
    - Runtime/UI: N/A.
    - DB/data: N/A.
    - Behavior test: focused MCP tests pass.
    - Cleanup/quarantine: no temp files outside repo.
    - Evidence IDs: `E1-P1D-IMPACT1`, `E1-P1D-SRC1`, `E1-P1D-TEST1`.
    - Actual-status rows refreshed: MCP row transitions `wrong -> correct`.
  - Evidence Targets: impact, source diff, focused MCP tests.
  - Actual-status Update: refresh MCP row.
  - Commit Boundary: commit after this slice when acceptance passes.

### P2: Validation, Review, And Closure

- Phase Goal: prove the full invariant across command surfaces, remove dead work, and close the plan.
- Phase Boundary:
  - In scope: full build, targeted tests, runtime command/API/MCP validation, detect-changes, supervisor review, plan ledger closure.
  - Out of scope: unrelated refactors or UI redesign.
  - Dependencies: all P1 slices complete or explicitly blocked.
- Phase Implementation Rule: validation must use the real command surfaces and fresh graph evidence.
- Ordered Slice List:
  - P2-A: Full validation and detect-changes.
  - P2-B: Supervisor review and dead-work cleanup.
  - P2-C: Close plan.

- [ ] P2-A: Full validation and detect-changes.
  - Goal: prove the implemented behavior at the nearest real boundaries.
  - Work Steps:
    1. Run `anvien analyze --force` after implementation if graph-related source changed and record evidence.
    2. Run full build before validation: `go build ./...`.
    3. Run focused tests for filecontext, CLI, HTTP, and MCP.
    4. Run manual CLI validation for repo-relative and absolute paths against `Anvien`.
    5. If HTTP/MCP runtime tests are not sufficient, start the local Anvien runtime and validate `/api/file-detail` or MCP behavior with real requests.
    6. Run `anvien detect-changes --repo Anvien --scope all`.
  - Implementation Gate: P1-A through P1-D are complete or blocked with evidence.
  - Acceptance: build, tests, real command checks, and detect-changes evidence are recorded.

- [ ] P2-B: Supervisor review and dead-work cleanup.
  - Goal: verify the completed plan work and remove artifacts created by failed or obsolete approaches.
  - Work Steps:
    1. Run supervisor review over plan, actual-status, evidence, benchmark, diff, tests, and command results.
    2. If supervisor rejects, fix only the rejected scope and rerun the relevant evidence.
    3. Review final diff for dead work created by this plan and remove it.
    4. Rerun supervisor if cleanup changed source or tests.
  - Implementation Gate: P2-A evidence exists.
  - Acceptance: supervisor passes or records a blocker; no dead plan-created artifacts remain.

- [ ] P2-C: Close plan.
  - Goal: finish ledgers, commit state, and closure notes.
  - Work Steps:
    1. Update evidence, benchmark, and actual-status with final validation and status transitions.
    2. Record final commits for completed implementation slices.
    3. Verify `git status --short`.
    4. Mark checklist items complete only after evidence exists.
  - Implementation Gate: P2-B passes or blocks with evidence.
  - Acceptance: plan files reflect final state and worktree status is known.

## Risk Notes

- `internal/filecontext/context.go` is high blast radius by relationship count, so the shared helper must be narrow and tested before wiring surfaces.
- MCP file context helpers are shared by context, impact, and summary paths; preserve symbol and smart target dispatch.
- Changing error behavior can affect API consumers; keep successful response shape stable and make outside-repo errors explicit.
- `--repo` may be a repo name or repo path; path normalization must use the resolved registry `RepoPath`, not the raw `--repo` string.
