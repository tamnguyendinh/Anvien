# AVmatrix Runtime Lock And Process Lifecycle Hardening Plan

Date: 2026-05-26

Status: Planned

Companion files:

- Evidence ledger: [2026-05-26-avmatrix-runtime-lock-process-lifecycle-hardening-evidence.md](2026-05-26-avmatrix-runtime-lock-process-lifecycle-hardening-evidence.md)
- Benchmark ledger: [2026-05-26-avmatrix-runtime-lock-process-lifecycle-hardening-benchmark.md](2026-05-26-avmatrix-runtime-lock-process-lifecycle-hardening-benchmark.md)

## Master Rules

1. Use AVmatrix for codebase analysis and impact checks while working on implementation slices in this plan.
2. As each task is completed, update the corresponding checklist item immediately.
3. Run a full build before testing: `powershell -ExecutionPolicy Bypass -File avmatrix-launcher\build.ps1`.
4. Record benchmark results as each benchmarkable task is completed. Benchmarkable means measured runtime behavior, process counts, lock inventory counts, command output counts, recovery rates, or CLI/API response inventory. Build/test timings are validation evidence unless the slice changes those systems.
5. Record evidence as each evidenced task is completed.
6. For doc-only commits, do not use AVmatrix.
7. After each completed implementation slice, commit the work, then continue until the full plan is complete.

## Problem

AVmatrix can leave users blocked by stale repository lock files and by confusing long-running background processes.

The repository lock for analyze/embed work is currently implemented as `.avmatrix/analyze.lock`. The lock file records a PID and timestamp, but acquisition currently fails immediately when the file already exists. It does not read the PID, check whether the process is still alive, recover after crash/reboot, or report actionable metadata. A stale lock can therefore make `avmatrix analyze --force` fail even when no analyze process is running.

Long-running processes such as `avmatrix mcp` and `avmatrix serve` are valid in some contexts, but the user experience does not clearly distinguish them from an analyze process. Editor or agent setup can configure `avmatrix mcp` as a child process owned by Codex, Claude, Cursor, or another host. Closing a browser window does not stop such a process, because it is not browser-owned. Users who inspect the process list can reasonably conclude AVmatrix is stuck unless the tool exposes runtime ownership clearly.

This is a product reliability issue. A global developer tool must recover from stale locks, explain real live locks, and make process ownership discoverable without requiring users to inspect OS process tables manually.

## Scope

Implementation may touch:

- lock implementation under `internal/repo/lock.go`;
- lock tests under `internal/repo`;
- analyze/embed lock callers under `internal/analyze`, `internal/httpapi/analyze.go`, and `internal/httpapi/embed.go`;
- CLI diagnostics under `internal/cli` if adding `doctor`, `runtime`, or lock/process status commands;
- setup guidance under `internal/cli/setup_command.go`, generated AI context guidance, MCP setup resources, README, or docs if user-facing wording changes;
- launcher lifecycle code under `avmatrix-launcher/src/main.go` only if process ownership or cleanup behavior must change;
- tests for analyze, HTTP analyze/embed, CLI diagnostics, setup text, and launcher process cleanup.

Out of scope unless source inspection proves it is required:

- changing graph analysis semantics;
- changing MCP protocol behavior;
- killing editor-owned MCP processes automatically;
- broad launcher refactors unrelated to runtime ownership;
- changing generated `AGENTS.md`, `CLAUDE.md`, or `.claude/skills/avmatrix/**` directly as source of truth.

## Design Decisions

Repository locks must be self-healing after owner death. When a lock exists, acquisition must inspect metadata and decide whether it is a live lock or stale lock. A live lock remains exclusive. A stale lock is removed safely and acquisition retries.

Lock metadata should include at least:

| Field | Purpose |
|---|---|
| `version` | Allows future-compatible parsing. |
| `pid` | Identifies the process that acquired the lock. |
| `acquiredAt` | Allows age and stale diagnostics. |
| `host` | Prevents unsafe PID interpretation across hosts when shared filesystems are involved. |
| `command` | Explains whether the owner was analyze, embed, serve, or another process. |
| `token` | Prevents one process from releasing a lock acquired by a later process. |

PID liveness must be platform-aware and testable. Prefer an injected or package-local process checker so unit tests can simulate live and dead PIDs without spawning fragile OS processes.

Recovery must be conservative on shared filesystems. If the lock host is different from the current host, do not remove it solely because the local PID is not alive. Report the foreign-host lock clearly unless an explicit stale policy is implemented and documented.

Malformed lock recovery must be deliberate. A malformed lock should not create a permanent dead end, but it also should not be removed instantly if it may be actively written. Use a short grace period or a minimum age threshold before considering malformed metadata stale.

`Release` must be ownership-safe. A lock holder should remove only the lock that it acquired, preferably by comparing token metadata before deleting. If the lock file has been replaced, release should not delete the new owner's lock.

Runtime processes must be classified by command and owner:

| Process class | Expected lifecycle |
|---|---|
| `analyze` | One-shot; should exit after analysis and release lock. |
| `embed` | One-shot or job-scoped; should release lock after job/cancel. |
| `serve` | Long-running Web/API runtime; may be launcher-owned. |
| `mcp` | Long-running editor/agent server; usually editor-owned. |
| launcher support | Launcher-owned and safe to clean up by launcher ownership rules. |

The tool must not kill editor-owned MCP processes as part of browser or launcher cleanup. Diagnostics can report them, and explicit user commands may stop only clearly selected processes if such a command is added.

## Acceptance Criteria

- A stale `.avmatrix/analyze.lock` whose PID is dead on the same host no longer blocks `avmatrix analyze --force`.
- A live lock still blocks concurrent analyze/embed writers.
- Lock conflict errors include actionable metadata: lock path, PID, age, command when known, and whether the owner appears alive.
- `Release` does not remove a lock created by a different process after the original lock was recovered or replaced.
- Malformed stale locks are recoverable under a documented policy.
- HTTP analyze/embed endpoints preserve real concurrency protection while recovering stale locks.
- Users can distinguish `analyze`, `embed`, `serve`, and `mcp` process classes through diagnostics, status output, help text, or setup guidance.
- Browser/launcher shutdown does not kill editor-owned MCP processes.
- Tests cover live lock exclusion, stale lock recovery, malformed stale locks, token-safe release, and CLI/API error behavior.
- Full build and focused tests pass before closure.

## Phase 0 - Baseline And Reproduction

- [ ] [P0-A] Refresh the AVmatrix graph before graph-based work and record graph counts in the evidence ledger.
- [ ] [P0-B] Trace lock owner code and record files/symbols for `AcquireStorageLock`, `StorageLock.Release`, analyze callers, HTTP analyze checks, and embed lock acquisition.
- [ ] [P0-C] Reproduce stale lock behavior with a lock file containing a dead PID. Record command output, lock file content, and expected recovery behavior.
- [ ] [P0-D] Inventory current AVmatrix processes on Windows with parent process chains. Record which processes are editor-owned `mcp`, launcher-owned `serve`, one-shot CLI commands, or unknown.
- [ ] [P0-E] Trace setup-generated MCP configuration for Codex, Claude, Cursor, and OpenCode. Record how `avmatrix mcp` is started and why browser shutdown does not own that lifecycle.
- [ ] [P0-F] Run impact analysis before editing lock or lifecycle symbols. Record blast radius and risk level for `AcquireStorageLock`, `StorageLock.Release`, HTTP analyze/embed lock callers, and any launcher cleanup symbols touched.

## Phase 1 - Stale Lock Recovery

- [ ] [P1-A] Define lock metadata format and compatibility behavior for existing two-line locks containing only `pid` and `acquiredAt`.
- [ ] [P1-B] Add lock metadata parsing, process liveness detection, host matching, malformed metadata handling, and age calculation.
- [ ] [P1-C] Update `AcquireStorageLock` so same-host dead-PID locks are removed and acquisition retries.
- [ ] [P1-D] Update conflict errors to carry structured lock metadata while preserving `errors.Is(err, ErrLockHeld)` compatibility.
- [ ] [P1-E] Update `StorageLock.Release` so it removes only the lock token it owns.
- [ ] [P1-F] Add unit tests for live lock exclusion, dead PID stale recovery, old-format lock recovery, malformed stale lock recovery, foreign-host safety, and token-safe release.

## Phase 2 - Analyze, Embed, And API Behavior

- [ ] [P2-A] Verify CLI analyze uses the hardened lock path without extra caller-specific stale logic.
- [ ] [P2-B] Verify HTTP analyze preflight and actual analyze job do not race after stale recovery.
- [ ] [P2-C] Verify embed job lock acquisition and release preserve mutual exclusion while recovering stale locks.
- [ ] [P2-D] Add integration tests for CLI analyze and HTTP analyze/embed stale-lock recovery.
- [ ] [P2-E] Update user-facing error text for real live locks so it explains the owning process and next action.

## Phase 3 - Runtime And Process Diagnostics

- [ ] [P3-A] Decide the user-facing diagnostics surface: `avmatrix doctor locks`, `avmatrix doctor processes`, `avmatrix runtime status`, or an equivalent command design.
- [ ] [P3-B] Implement lock diagnostics that report lock path, repo, owner PID, host, age, command, liveness, and stale/recoverable status.
- [ ] [P3-C] Implement process diagnostics that classify AVmatrix processes as analyze/embed/serve/mcp/launcher/unknown and identify likely owner from command line and parent process when available.
- [ ] [P3-D] Add `--json` output for diagnostics so agents can consume the result.
- [ ] [P3-E] Add tests for command registration, JSON shape, process classification, and missing-lock/no-process output.

## Phase 4 - Setup, Launcher, And Documentation UX

- [ ] [P4-A] Update setup output and setup-installed guidance so users know `avmatrix mcp` is long-running and editor-owned.
- [ ] [P4-B] Verify launcher cleanup targets only launcher-owned runtime processes and does not kill editor-owned MCP processes.
- [ ] [P4-C] Update README and active docs to explain lock recovery, live-lock diagnostics, MCP process ownership, and safe cleanup commands.
- [ ] [P4-D] Update generated AI context guidance if final diagnostics commands become part of the AVmatrix command surface.

## Phase 5 - Validation And Closure

- [ ] [P5-A] Run the full build gate: `powershell -ExecutionPolicy Bypass -File avmatrix-launcher\build.ps1`.
- [ ] [P5-B] Run focused tests: `go test .\internal\repo .\internal\analyze .\internal\httpapi .\internal\cli -count=1`.
- [ ] [P5-C] Run launcher tests if launcher lifecycle code changed.
- [ ] [P5-D] Smoke test stale lock recovery by writing a dead-PID lock and running `avmatrix analyze --force`.
- [ ] [P5-E] Smoke test live lock conflict and confirm the error message includes actionable owner metadata.
- [ ] [P5-F] Smoke test diagnostics command output in table and JSON mode.
- [ ] [P5-G] Run `avmatrix detect-changes --repo AVmatrix --scope all` before committing implementation work and record changed scope.
- [ ] [P5-H] Close the plan only after code, tests, diagnostics output, docs, evidence, and benchmark ledgers agree on stale-lock recovery and process lifecycle behavior.
