# Anvien Runtime Lock And Process Lifecycle Hardening Plan

Date: 2026-05-26

Status: Complete

Companion files:

- Evidence ledger: [2026-05-26-anvien-runtime-lock-process-lifecycle-hardening-evidence.md](2026-05-26-anvien-runtime-lock-process-lifecycle-hardening-evidence.md)
- Benchmark ledger: [2026-05-26-anvien-runtime-lock-process-lifecycle-hardening-benchmark.md](2026-05-26-anvien-runtime-lock-process-lifecycle-hardening-benchmark.md)

## Master rules

1. Use Anvien for codebase analysis and impact checks while working on implementation slices in this plan.
2. As each task is completed, update the corresponding checklist item immediately.
3. Run the full build gate before testing; include focused backend/CLI/setup/package validation for generated skill behavior, and include Web unit/e2e/browser screenshot validation for the graph labeling phase.
4. Record benchmark results as each benchmarkable task is completed. Benchmarkable means measured product/runtime performance, capacity, package/startup size, graph/DB throughput, query hit rate, command output inventory, graph inventory counts, source-site inventory counts, generated skill inventory counts, setup/package file inventories, or resolved-edge accuracy; build/test/e2e timings are validation evidence unless the slice changes those systems.
5. Record evidence as each evidenced task is completed.
6. For doc-only commits, do not use Anvien.
7. After each completed implementation slice, commit the work, then continue until the full plan is complete.

## Problem

Anvien can leave users blocked by stale repository lock files and by confusing long-running background processes.

The repository lock for analyze/embed work is currently implemented as `.anvien/analyze.lock`. The lock file records a PID and timestamp, but acquisition currently fails immediately when the file already exists. It does not read the PID, check whether the process is still alive, recover after crash/reboot, or report actionable metadata. A stale lock can therefore make `anvien analyze --force` fail even when no analyze process is running.

Long-running processes such as `anvien mcp` and `anvien serve` are valid in some contexts, but the user experience does not clearly distinguish them from an analyze process. Editor or agent setup can configure `anvien mcp` as a child process owned by Codex, Claude, Cursor, or another host. Closing a browser window does not stop such a process, because it is not browser-owned. Users who inspect the process list can reasonably conclude Anvien is stuck unless the tool exposes runtime ownership clearly.

This is a product reliability issue. A global developer tool must recover from stale locks, explain real live locks, and make process ownership discoverable without requiring users to inspect OS process tables manually.

## Scope

Implementation may touch:

- lock implementation under `internal/repo/lock.go`;
- lock tests under `internal/repo`;
- analyze/embed lock callers under `internal/analyze`, `internal/httpapi/analyze.go`, and `internal/httpapi/embed.go`;
- CLI diagnostics under `internal/cli` if adding `doctor`, `runtime`, or lock/process status commands;
- targeted CLI-facing setup/diagnostic text under `internal/cli` if needed to make the new lock/process behavior actionable;
- launcher lifecycle code under `anvien-launcher/src/main.go` only if process ownership or cleanup behavior must change;
- tests for analyze, HTTP analyze/embed, CLI diagnostics, setup text, and launcher process cleanup.

Out of scope unless source inspection proves it is required:

- changing graph analysis semantics;
- changing MCP protocol behavior;
- killing editor-owned MCP processes automatically;
- broad launcher refactors unrelated to runtime ownership;
- broad README, active docs, MCP setup resources, or generated AI context guidance unless source inspection proves the bug fix is unusable without them;
- changing generated `AGENTS.md`, `CLAUDE.md`, or `.claude/skills/anvien/**` directly as source of truth.

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

- A stale `.anvien/analyze.lock` whose PID is dead on the same host no longer blocks `anvien analyze --force`.
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

- [x] [P0-A] Refresh the Anvien graph before graph-based work and record graph counts in the evidence ledger.
- [x] [P0-B] Trace lock owner code and record files/symbols for `AcquireStorageLock`, `StorageLock.Release`, analyze callers, HTTP analyze checks, and embed lock acquisition.
- [x] [P0-C] Reproduce stale lock behavior with a lock file containing a dead PID. Record command output, lock file content, and expected recovery behavior.
- [x] [P0-D] Inventory current Anvien processes on Windows with parent process chains. Record which processes are editor-owned `mcp`, launcher-owned `serve`, one-shot CLI commands, or unknown.
- [x] [P0-E] Trace setup-generated MCP configuration for Codex, Claude, Cursor, and OpenCode. Record how `anvien mcp` is started and why browser shutdown does not own that lifecycle.
- [x] [P0-F] Run impact analysis before editing lock or lifecycle symbols. Record blast radius and risk level for `AcquireStorageLock`, `StorageLock.Release`, HTTP analyze/embed lock callers, and any launcher cleanup symbols touched.

## Phase 1 - Stale Lock Recovery

- [x] [P1-A] Define lock metadata format and compatibility behavior for existing two-line locks containing only `pid` and `acquiredAt`.
- [x] [P1-B] Add lock metadata parsing, process liveness detection, host matching, malformed metadata handling, and age calculation.
- [x] [P1-C] Update `AcquireStorageLock` so same-host dead-PID locks are removed and acquisition retries.
- [x] [P1-D] Update conflict errors to carry structured lock metadata while preserving `errors.Is(err, ErrLockHeld)` compatibility.
- [x] [P1-E] Update `StorageLock.Release` so it removes only the lock token it owns.
- [x] [P1-F] Add unit tests for live lock exclusion, dead PID stale recovery, old-format lock recovery, malformed stale lock recovery, foreign-host safety, and token-safe release.

## Phase 2 - Analyze, Embed, And API Behavior

- [x] [P2-A] Verify CLI analyze uses the hardened lock path without extra caller-specific stale logic.
- [x] [P2-B] Verify HTTP analyze preflight and actual analyze job do not race after stale recovery.
- [x] [P2-C] Verify embed job lock acquisition and release preserve mutual exclusion while recovering stale locks.
- [x] [P2-D] Add integration tests for CLI analyze and HTTP analyze/embed stale-lock recovery.
- [x] [P2-E] Update user-facing error text for real live locks so it explains the owning process and next action. The message now includes lock path, PID, host, acquired time, age, command, live/stale reason, and a next action pointing to `anvien doctor locks --repo <repo>`.

## Phase 3 - Runtime And Process Diagnostics

- [x] [P3-A] Decide the user-facing diagnostics surface: `anvien doctor locks`, `anvien doctor processes`, `anvien runtime status`, or an equivalent command design.
- [x] [P3-B] Implement lock diagnostics that report lock path, repo, owner PID, host, age, command, liveness, and stale/recoverable status.
- [x] [P3-C] Implement process diagnostics that classify Anvien processes as analyze/embed/serve/mcp/launcher/unknown and identify likely owner from command line and parent process when available.
- [x] [P3-D] Add `--json` output for diagnostics so agents can consume the result.
- [x] [P3-E] Add tests for command registration, JSON shape, process classification, and missing-lock/no-process output.

## Phase 4 - Setup, Launcher, And Documentation UX

- [x] [P4-A] Update targeted setup output so users know `anvien mcp` is long-running and editor-owned. Setup-installed guidance is rescoped out because this lock bug does not require broad skill/documentation updates.
- [x] [P4-B] Verify launcher cleanup targets only launcher-owned runtime processes and does not kill editor-owned MCP processes.
- [x] [P4-C] Re-scope README and active docs out of this bug fix. Lock recovery is validated through code, diagnostics command output/help, tests, smoke checks, and this plan ledger.
- [x] [P4-D] Re-scope generated AI context guidance out of this bug fix. Do not update `internal/aicontext` or generated context files for this lock-lifecycle change.

## Phase 5 - Validation And Closure

- [x] [P5-A] Run the full build gate: `powershell -ExecutionPolicy Bypass -File anvien-launcher\build.ps1`.
- [x] [P5-B] Run focused tests: `go test .\internal\repo .\internal\analyze .\internal\httpapi .\internal\cli -count=1`.
- [x] [P5-C] Run launcher tests if launcher lifecycle code changed.
- [x] [P5-D] Smoke test stale lock recovery by writing a dead-PID lock and running `anvien analyze --force`.
- [x] [P5-E] Smoke test live lock conflict and confirm the error message includes actionable owner metadata.
- [x] [P5-F] Smoke test diagnostics command output in table and JSON mode.
- [x] [P5-G] Run `anvien detect-changes --repo Anvien --scope all` before committing implementation work and record changed scope.
- [x] [P5-H] Close the plan only after code, tests, diagnostics output, scope decisions, evidence, and benchmark ledgers agree on stale-lock recovery and process lifecycle behavior.
