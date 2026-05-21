# AVmatrix Web Local Folder Picker Stuck Evidence Ledger

Date: 2026-05-21

Status: initial evidence recorded

Plan: [2026-05-21-avmatrix-web-local-folder-picker-stuck-plan.md](2026-05-21-avmatrix-web-local-folder-picker-stuck-plan.md)

Benchmark: [2026-05-21-avmatrix-web-local-folder-picker-stuck-benchmark.md](2026-05-21-avmatrix-web-local-folder-picker-stuck-benchmark.md)

## E0 - User Report Evidence

Status: recorded

User-provided failing screenshot:

- `reports/problem/screenshot_1779328050.png`

Observed from screenshot:

- Screen: Analyze Repository.
- Existing repo cards are visible, so the app reached the local runtime landing/analyze surface.
- The repository chooser button is stuck on `Opening repository picker...`.
- The screenshot points at the stuck chooser button.

Initial classification:

- This is a local folder picker/analyze UI flow issue.
- It is not the graph clustered layout issue.
- It is not evidence of layout optimizer auto-running.

## E0A - Runtime Reset Window Flashing Report

Status: recorded

User-reported additional problem:

- The current `RESET RUNTIME` feature causes terminal-like windows to appear or flash repeatedly while it works.
- The expected behavior is fully hidden runtime reset. Visible terminal/helper windows are not acceptable for the user-facing packaged flow.

Initial classification:

- This is a local runtime launcher lifecycle issue.
- It belongs in this plan because it affects the same Start/local-runtime user flow as the folder picker issue.
- It must not be fixed with product timeout or delayed reset logic.

## E1 - Initial Code Inspection Evidence

Status: recorded

Files inspected:

- `avmatrix-web/src/components/RepoAnalyzer.tsx`
- `avmatrix-web/src/services/backend-client.ts`
- `internal/httpapi/local_folder_picker.go`
- `internal/httpapi/handlers_test.go`
- `avmatrix-web/e2e/onboarding.spec.ts`
- `Start-AVmatrix.html`
- `avmatrix-launcher/src/main.go`
- `avmatrix-launcher/server-wrapper/main.go`

Relevant facts:

- `RepoAnalyzer` sets `isPickingFolder=true` before awaiting `pickLocalFolder()`.
- `isPickingFolder` is reset only after the picker promise resolves or rejects.
- `canSubmit` requires `!isPickingFolder`, so a pending picker can also block the Analyze Repository button.
- `pickLocalFolder` sends `POST /api/local/folder-picker` without an abort signal.
- `fetchFromBackend` accepts `RequestInit.signal`, but wraps browser `AbortError` as `BackendError('Request aborted', status=0, code='network')`.
- Implementation must therefore avoid showing that wrapped abort as a false picker error.
- The backend endpoint calls `pickLocalFolderFunc()` without passing `r.Context()`.
- Picker commands use `exec.Command(...).Output()`, not `exec.CommandContext`.
- Windows picker uses a PowerShell `System.Windows.Forms.FolderBrowserDialog`.
- Existing e2e covers the successful mocked picker path, but not pending/cancel recovery.
- `Start-AVmatrix.html` sends reset through `window.location.href = 'avmatrix://reset'`.
- `avmatrix-launcher/src/main.go` owns `resetRuntime`, `stopRuntime`, `stopRuntimeProcessesByPath`, `stopPID`, `registerProtocol`, `openBrowser`, and `hiddenProcAttr`.
- `stopRuntimeProcessesByPath` already sets `hiddenProcAttr()` on its PowerShell process sweep.
- `stopPID` already sets `hiddenProcAttr()` on Windows `taskkill` commands.
- `processAlive` runs `tasklist` without `hiddenProcAttr()`.
- `waitForPIDExit` calls `processAlive` repeatedly. This is the concrete reset-window-flash candidate because reset/stop can spawn visible `tasklist` console windows in a loop.
- `registerProtocol` already sets `hiddenProcAttr()` on `reg` commands.
- `openBrowser` already sets `hiddenProcAttr()` for the Windows `rundll32` browser launch helper.
- `avmatrix-launcher/server-wrapper/main.go` starts the backend with `hiddenProcAttr()`.
- Because user still sees flashing terminal-like windows, implementation must audit the entire reset/protocol/package chain, not only individual child commands that already have hidden process attributes.

Artifact and protocol inspection:

- `HKCU\Software\Classes\avmatrix\shell\open\command` currently points to `"E:\AVmatrix-GO\avmatrix-launcher\AVmatrixLauncher.exe" "%1"`.
- PE subsystem inspection of current artifacts:
  - `avmatrix-launcher/AVmatrixLauncher.exe`: Windows GUI.
  - `avmatrix-launcher/server-bundle/avmatrix-server.exe`: Windows GUI.
  - `avmatrix-launcher/server-bundle/avmatrix.exe`: Windows CUI/Console.
- Interpretation: the backend console executable can be acceptable only when it is started through `avmatrix-server.exe` with hidden process attributes. The repeated visible-window report is more directly explained by unhidden `tasklist` calls inside `processAlive`.

## E2 - Initial AVmatrix Tool Evidence

Status: recorded

This section records tool output already gathered during initial inspection. It is historical evidence, not a standalone replacement for the active `AGENTS.md` workflow.

AVmatrix refresh:

- `avmatrix analyze --force`
  - Passed.
  - `files: scanned=715 parsed=538 unsupported=177 failed=0`
  - `graph: nodes=21821 relationships=54366`

Impact checks:

- `avmatrix impact handleLocalFolderPicker --repo "E:\AVmatrix-GO" --direction upstream --depth 2 --include-tests`
  - Risk: LOW.
  - Impacted count: 0.
- `avmatrix impact pickCommandFolder --repo "E:\AVmatrix-GO" --direction upstream --depth 2 --include-tests`
  - Risk: LOW.
  - Direct impacted symbols: `pickLocalFolder`, `pickWindowsFolder`.
- `avmatrix impact -u "Function:avmatrix-web/src/services/backend-client.ts:pickLocalFolder" --repo "E:\AVmatrix-GO" --direction upstream --depth 2 --include-tests`
  - Risk: LOW.
  - Impacted scope includes `RepoAnalyzer.tsx`, `handleChooseRepository`, `AnalyzeOnboarding.tsx`, `Header.tsx`, `RepoLanding.tsx`, and `RepoAnalyzer.local-only.test.tsx`.
- `avmatrix impact RepoAnalyzer --repo "E:\AVmatrix-GO" --direction upstream --depth 2 --include-tests`
  - Risk: LOW.
  - Impacted scope includes `AnalyzeOnboarding.tsx`, `Header.tsx`, `RepoLanding.tsx`, and related tests.

## E3 - Constraint Evidence

Status: recorded

The fix must not use product/runtime timeout logic.

Scope clarification:

- The implementation must not add new timeout, timer reset, delayed reset, or elapsed-time recovery behavior.
- Existing bounded waits in launcher lifecycle code are not the selected fix mechanism.
- This plan should not become a broad lifecycle-wait refactor unless a specific existing wait is proven to cause the reported picker or reset-window behavior.

Allowed:

- request cancellation through `AbortController`;
- request-context propagation in Go;
- `exec.CommandContext`;
- hidden process attributes for spawned helper processes;
- a shared hidden-command helper for launcher child processes;
- Windows launcher/server-wrapper packaging/build settings that avoid visible console windows;
- test assertion timeouts.

Not allowed:

- product timeout around folder picker;
- delayed UI reset after elapsed time;
- timer-based auto-close;
- treating a large repo or slow machine as a reason to add a timeout.
- accepting flashing terminal/helper windows as normal reset behavior.
- solving reset by instructing users to run commands manually in a terminal.

## E4 - Implementation Evidence

Status: pending

Record changed files and implementation details after code changes.

## E5 - Validation Evidence

Status: pending

Record focused tests, build, e2e/browser validation, and required pre-commit scope evidence after implementation.
