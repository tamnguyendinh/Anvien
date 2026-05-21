# AVmatrix Web Local Folder Picker Stuck Plan

Date: 2026-05-21

Status: planned

Companion files:

- Benchmark ledger: [2026-05-21-avmatrix-web-local-folder-picker-stuck-benchmark.md](2026-05-21-avmatrix-web-local-folder-picker-stuck-benchmark.md)
- Evidence ledger: [2026-05-21-avmatrix-web-local-folder-picker-stuck-evidence.md](2026-05-21-avmatrix-web-local-folder-picker-stuck-evidence.md)

## Rules

1. Follow active workspace and repository instructions, including `AGENTS.md`, for AVmatrix and codebase workflow. This plan records product work and validation; it does not replace those rules.
2. Update this plan, the benchmark ledger, and the evidence ledger as each implementation slice is completed.
3. Do not add product/runtime timeout, timer reset, delayed reset, or elapsed-time budget to solve this issue.
4. Timeout is allowed only in test runners or assertions that bound test failure.
5. Existing bounded waits in launcher lifecycle code are not the fix mechanism for this issue. Do not expand this plan into an unrelated lifecycle-wait refactor unless implementation evidence proves a specific existing wait causes the reported behavior.
6. Keep manual path entry usable. The folder picker is optional convenience, not a required path to analysis.
7. After implementation, run focused frontend/backend tests, Web build, applicable Go tests, launcher tests/builds, and browser/e2e validation for the chooser/runtime flow.

## Problem

User-provided screenshot `reports/problem/screenshot_1779328050.png` shows the Analyze Repository screen stuck on:

`Opening repository picker...`

The UI is waiting for `POST /api/local/folder-picker` to finish. If the native OS picker does not appear, appears behind another window, cannot open in the packaged/runtime context, or is left pending, the user has no visible recovery path from the picker state.

This is not a graph layout issue. It is a local repository chooser/analyze flow issue.

## Runtime Reset Window Flashing Finding

The same local runtime surface also has a second reported UX failure: using `RESET RUNTIME` currently causes terminal-like windows to appear or flash repeatedly.

Expected behavior:

- reset runtime must run fully hidden;
- no terminal, console, PowerShell, cmd, taskkill, process sweep, backend, server wrapper, or launcher helper window should visibly appear;
- reset may update the Start page status text, but OS-level helper windows should not flash;
- this applies to the packaged/user-facing `Start-AVmatrix.html` flow, not only to the development server flow.

Current entry point:

- `Start-AVmatrix.html` calls `window.location.href = 'avmatrix://reset'`.

This issue is related to local runtime UX and lifecycle control. It is still separate from graph layout and repository indexing.

## Product Intent

The user must always be able to recover and continue:

- clicking `Choose Repository` may open the native OS picker;
- if the picker does not work, the user can cancel the picker state without reloading the app;
- pasting an absolute local path remains a first-class path;
- analysis can start from a pasted path even if a picker request was previously started;
- closing the picker or canceling from the UI must not leave a hanging request or child picker process;
- runtime reset is allowed to stop and restart local processes, but it must not open visible terminal/helper windows;
- all reset/runtime helper processes must run hidden in the packaged user-facing path;
- no timeout-based product logic is allowed.
- this plan must not add any new timeout/timer-based recovery path.

## Current Code Facts

- `avmatrix-web/src/components/RepoAnalyzer.tsx` owns `isPickingFolder`.
- `handleChooseRepository` sets `isPickingFolder=true`, awaits `pickLocalFolder()`, then resets in `finally`.
- `canSubmit` currently requires `!isPickingFolder`, so a stuck picker also blocks Analyze Repository even if the user pastes a valid path.
- `avmatrix-web/src/services/backend-client.ts` sends `POST /api/local/folder-picker` without an abort signal.
- `fetchFromBackend` already accepts `RequestInit.signal`, but wraps browser `AbortError` as `BackendError('Request aborted', status=0, code='network')`.
- `internal/httpapi/local_folder_picker.go` handles `/api/local/folder-picker`.
- Backend picker commands currently use `exec.Command(...).Output()` rather than request-context-aware execution.
- On Windows, the backend starts PowerShell with `System.Windows.Forms.FolderBrowserDialog`.
- `Start-AVmatrix.html` exposes `RESET RUNTIME` through the `avmatrix://reset` protocol URL.
- `avmatrix-launcher/src/main.go` owns `resetRuntime`, `stopRuntime`, `stopRuntimeProcessesByPath`, `stopPID`, protocol registration, browser opening, and `hiddenProcAttr`.
- `stopRuntimeProcessesByPath` runs a PowerShell process sweep and already assigns `hiddenProcAttr()`.
- `stopPID` runs `taskkill` and already assigns `hiddenProcAttr()`.
- `processAlive` runs Windows `tasklist` without assigning `hiddenProcAttr()`.
- `waitForPIDExit` calls `processAlive` repeatedly, so a visible `tasklist` console can flash repeatedly during reset/stop.
- `avmatrix-launcher/server-wrapper/main.go` starts the Go backend and already assigns `hiddenProcAttr()`.
- The current checked artifact has GUI subsystem for `avmatrix-launcher/AVmatrixLauncher.exe` and `avmatrix-launcher/server-bundle/avmatrix-server.exe`, but `avmatrix-launcher/server-bundle/avmatrix.exe` is a Console subsystem binary that must only be launched through a hidden wrapper path.
- Current `HKCU\Software\Classes\avmatrix\shell\open\command` points at `E:\AVmatrix-GO\avmatrix-launcher\AVmatrixLauncher.exe "%1"`.
- The user report means the implementation must audit the whole reset launch chain, including how protocol invocation starts the launcher process itself and whether built binaries use the correct Windows subsystem.

## Non-Goals

- Do not add product timeout around the picker request.
- Do not auto-close the picker after elapsed time.
- Do not add a delayed reset or timer-based retry to hide the picker/reset symptoms.
- Do not broaden the implementation into a cleanup of existing bounded waits such as launcher URL/PID waits unless one is proven to be the direct cause.
- Do not hide the manual path input while the picker is pending.
- Do not change the analyze job contract unrelated to picker cancellation.
- Do not change graph layout, graph loading, graph schema, or repo indexing behavior in this plan.
- Do not accept terminal/window flashing as a normal reset side effect.
- Do not solve reset flashing by asking users to run from a terminal manually.

## Target Design

Use explicit user cancellation and request context propagation.

Frontend:

- Track the in-flight picker request with an `AbortController`.
- While picker is pending, show an explicit cancel action instead of a dead-end disabled button.
- If the user starts analysis from a valid pasted path while the picker is pending, abort the picker request first and proceed with analysis.
- On component unmount, abort any in-flight picker request and any in-flight SSE analysis stream.
- Treat picker abort as user cancellation, not as an error banner. Because `fetchFromBackend` wraps aborts, the UI must detect cancellation through the controller signal and/or the wrapped `BackendError('Request aborted')` shape.

Backend:

- Thread `r.Context()` from `handleLocalFolderPicker` into the picker implementation.
- Use `exec.CommandContext` for OS picker commands so client cancellation can terminate the picker process.
- Preserve existing explicit picker cancel behavior: OK returns `path`, Cancel returns `{ path: null, cancelled: true }`.
- Return unsupported picker errors as the current not-implemented response.

Launcher/runtime reset:

- Audit every process spawned during `avmatrix://reset`.
- Ensure direct child processes use hidden Windows process attributes where applicable.
- Introduce or reuse one helper for Windows external commands that always applies `hiddenProcAttr()` before `Run`, `Output`, or `CombinedOutput`.
- Apply the hidden command path to `tasklist` inside `processAlive`; this is the concrete repeated-console-flash candidate.
- Ensure the packaged launcher/server-wrapper path does not visibly allocate a console window.
- Verify protocol registration points to the intended packaged launcher executable.
- Verify launcher/server-wrapper build outputs use Windows GUI subsystem, and verify the backend console binary is only started through `avmatrix-server.exe` with hidden process attributes.
- Keep logs in files, not visible terminal windows.
- Preserve reset semantics: stale backend/launcher processes are stopped and runtime can start again.

## Implementation Phases

### P0 - Reproduce And Classify

- [ ] [P0-A] Record `reports/problem/screenshot_1779328050.png` as the failing visual evidence.
- [ ] [P0-B] Verify the stuck state is caused by the folder picker request remaining pending, not by analyze job startup.
- [ ] [P0-C] Confirm manual path entry should remain available and recoverable without page reload.
- [ ] [P0-D] Record current behavior in the benchmark ledger.

### P1 - Frontend Cancellation Semantics

- [ ] [P1-A] Update `pickLocalFolder` to accept an optional `AbortSignal`.
- [ ] [P1-B] Store the picker `AbortController` in `RepoAnalyzer`.
- [ ] [P1-C] Replace the stuck pending button state with an explicit cancel picker action.
- [ ] [P1-D] Do not show an error banner when the picker is aborted by the user.
- [ ] [P1-E] Allow valid manual path analysis to proceed by aborting any pending picker first.
- [ ] [P1-F] Abort any pending picker request when `RepoAnalyzer` unmounts.
- [ ] [P1-G] Add a local cancellation helper so wrapped `BackendError('Request aborted')` is handled as picker cancel only for the picker request.

### P2 - Backend Request-Context Cancellation

- [ ] [P2-A] Change `handleLocalFolderPicker` to pass `r.Context()` into picker execution.
- [ ] [P2-B] Change picker functions to accept `context.Context`.
- [ ] [P2-C] Use `exec.CommandContext` in `pickCommandFolder`.
- [ ] [P2-D] Preserve current OK, cancelled, and unsupported response contracts.
- [ ] [P2-E] Do not add product timeout, elapsed-time kill logic, delayed reset, or timer-based recovery.

### P3 - Tests

- [ ] [P3-A] Add frontend unit coverage for pending picker state showing a cancel action.
- [ ] [P3-B] Add frontend unit coverage proving cancel resets `isPickingFolder` without an error banner.
- [ ] [P3-C] Add frontend unit coverage proving pasted-path analysis can proceed after aborting a pending picker.
- [ ] [P3-D] Add frontend unit coverage proving unmount aborts a pending picker request.
- [ ] [P3-E] Add frontend unit coverage proving wrapped picker abort errors do not show a false error.
- [ ] [P3-F] Add backend test coverage proving request cancellation is propagated to picker execution.
- [ ] [P3-G] Add or update e2e coverage for chooser request success and pending/cancel recovery.

### P4 - Runtime Reset Hidden Execution

- [ ] [P4-A] Record the reset-runtime window flashing report in the evidence ledger.
- [ ] [P4-B] Audit `Start-AVmatrix.html` reset protocol flow and packaged launcher protocol registration.
- [ ] [P4-C] Audit `avmatrix-launcher/src` reset, stop, process sweep, taskkill, and browser-launch paths for visible windows.
- [ ] [P4-D] Audit `avmatrix-launcher/server-wrapper` backend launch path for visible windows.
- [ ] [P4-E] Fix `processAlive` so Windows `tasklist` runs with hidden process attributes.
- [ ] [P4-F] Ensure every reset/runtime helper process that can be hidden on Windows uses hidden process attributes through one consistent helper.
- [ ] [P4-G] Ensure packaged Windows launcher/server-wrapper binaries are built or wrapped so reset does not create a visible console window.
- [ ] [P4-H] Verify the backend console binary is not directly launched by protocol/reset; it must remain behind the hidden server-wrapper path.
- [ ] [P4-I] Add launcher/server-wrapper tests or build assertions covering hidden process attributes, including `tasklist`, and Windows subsystem expectations where practical.
- [ ] [P4-J] Preserve reset behavior: stale runtime processes are stopped, state is cleaned, and the user can start AVmatrix again.

### P5 - Validation

- [ ] [P5-A] Run focused frontend tests for `RepoAnalyzer`.
- [ ] [P5-B] Run focused backend tests for `local_folder_picker`.
- [ ] [P5-C] Run focused launcher/server-wrapper tests for hidden runtime process behavior.
- [ ] [P5-D] Run `npm --prefix avmatrix-web run build`.
- [ ] [P5-E] Run applicable Web unit tests.
- [ ] [P5-F] Run applicable Go tests.
- [ ] [P5-G] Run launcher and server-wrapper builds.
- [ ] [P5-H] Run browser/e2e chooser validation.
- [ ] [P5-I] Run packaged/manual reset validation or record why it cannot be automated in the current environment.

## Acceptance Criteria

- [ ] The UI cannot remain trapped in `Opening repository picker...` without a visible recovery action.
- [ ] User can cancel a pending repository picker from the Web UI.
- [ ] User can paste an absolute repository path and start analysis after a picker was started.
- [ ] Canceling the picker does not show a false analysis error.
- [ ] Unmounting the analyzer aborts any pending picker request.
- [ ] Backend picker execution is tied to request context cancellation.
- [ ] Product code does not introduce timeout, timer reset, delayed reset, or elapsed-time budget logic.
- [ ] Existing successful picker flow still fills the repository path.
- [ ] Existing picker Cancel behavior still returns a cancelled result.
- [ ] Existing analyze job flow remains unchanged after a valid path is submitted.
- [ ] `RESET RUNTIME` does not display flashing terminal/helper windows in the packaged user-facing path.
- [ ] Reset helper commands, process sweep, taskkill, backend launch, and server-wrapper launch run hidden where Windows supports hiding.
- [ ] Reset still stops stale runtime processes and allows AVmatrix to start again.
- [ ] Reset logging goes to files/status UI, not visible terminal windows.
- [ ] No new product timeout, timer reset, delayed reset, or elapsed-time recovery path is introduced.

## Completion Definition

This plan can be marked complete when the Analyze Repository screen provides explicit recovery from a stuck native picker, manual path analysis works without reload, the backend can cancel picker child processes through request context, `RESET RUNTIME` runs without visible terminal/helper window flashing in the packaged user-facing path, validation passes, benchmark/evidence ledgers are updated, no new timeout/timer-based product recovery path is introduced, and all active repo/tooling instructions have been satisfied.
