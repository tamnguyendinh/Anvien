# Anvien Web Local Folder Picker Stuck Evidence Ledger

Date: 2026-05-21

Status: completed

Plan: [2026-05-21-anvien-web-local-folder-picker-stuck-plan.md](2026-05-21-anvien-web-local-folder-picker-stuck-plan.md)

Benchmark: [2026-05-21-anvien-web-local-folder-picker-stuck-benchmark.md](2026-05-21-anvien-web-local-folder-picker-stuck-benchmark.md)

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

- `anvien-web/src/components/RepoAnalyzer.tsx`
- `anvien-web/src/services/backend-client.ts`
- `anvien-web/test/unit/RepoAnalyzer.local-only.test.tsx`
- `internal/httpapi/local_folder_picker.go`
- `internal/httpapi/handlers_test.go`
- `anvien-web/e2e/onboarding.spec.ts`
- `Start-Anvien.html`
- `anvien-launcher/src/main.go`
- `anvien-launcher/server-wrapper/main.go`
- `anvien-launcher/build.ps1`
- `TESTING.md`
- `RUNBOOK.md`

Relevant facts:

- `RepoAnalyzer` sets `isPickingFolder=true` before awaiting `pickLocalFolder()`.
- `isPickingFolder` is reset only after the picker promise resolves or rejects.
- `canSubmit` requires `!isPickingFolder`, so a pending picker can also block the Analyze Repository button.
- `pickLocalFolder` sends `POST /api/local/folder-picker` without an abort signal.
- `fetchFromBackend` accepts `RequestInit.signal`, but wraps browser `AbortError` as `BackendError('Request aborted', status=0, code='network')`.
- Implementation must therefore avoid showing that wrapped abort as a false picker error.
- `anvien-web/test/unit/RepoAnalyzer.local-only.test.tsx` exists and currently covers local input/submission/completion, but not pending picker cancel, unmount abort, or manual analysis while picker is pending.
- The backend endpoint calls `pickLocalFolderFunc()` without passing `r.Context()`.
- Current backend tests mock `pickLocalFolderFunc` as `func() (string, error)`, so the seam must be updated to assert request-context propagation.
- Picker commands use `exec.Command(...).Output()`, not `exec.CommandContext`.
- Windows picker uses a PowerShell `System.Windows.Forms.FolderBrowserDialog`.
- Existing e2e covers the successful mocked picker path, but not pending/cancel recovery.
- `Start-Anvien.html` sends reset through `window.location.href = 'anvien://reset'`.
- `anvien-launcher/src/main.go` owns `resetRuntime`, `stopRuntime`, `stopRuntimeProcessesByPath`, `stopPID`, `registerProtocol`, `openBrowser`, and `hiddenProcAttr`.
- `stopRuntimeProcessesByPath` already sets `hiddenProcAttr()` on its PowerShell process sweep.
- `stopPID` already sets `hiddenProcAttr()` on Windows `taskkill` commands.
- `processAlive` runs `tasklist` without `hiddenProcAttr()`.
- `waitForPIDExit` calls `processAlive` repeatedly. This is the concrete reset-window-flash candidate because reset/stop can spawn visible `tasklist` console windows in a loop.
- `registerProtocol` already sets `hiddenProcAttr()` on `reg` commands.
- `openBrowser` already sets `hiddenProcAttr()` for the Windows `rundll32` browser launch helper.
- `anvien-launcher/server-wrapper/main.go` starts the backend with `hiddenProcAttr()`.
- Because user still sees flashing terminal-like windows, implementation must audit the entire reset/protocol/package chain, not only individual child commands that already have hidden process attributes.
- `anvien-launcher/build.ps1` runs the Web build, builds `server-bundle\anvien.exe`, builds `AnvienLauncher.exe` with `-H=windowsgui`, builds `server-bundle\anvien-server.exe` with `-H=windowsgui`, copies `anvien-web\dist` to `anvien-launcher\web-dist`, and registers the protocol.
- `TESTING.md`/`RUNBOOK.md` define `Start-Anvien.html` as a repository-root user-facing entry that must remain absent from `anvien-web\dist\` and `anvien-launcher\web-dist\`.

Artifact and protocol inspection:

- `HKCU\Software\Classes\anvien\shell\open\command` currently points to `"E:\Anvien\anvien-launcher\AnvienLauncher.exe" "%1"`.
- PE subsystem inspection of current artifacts:
  - `anvien-launcher/AnvienLauncher.exe`: Windows GUI.
  - `anvien-launcher/server-bundle/anvien-server.exe`: Windows GUI.
  - `anvien-launcher/server-bundle/anvien.exe`: Windows CUI/Console.
- Interpretation: the backend console executable can be acceptable only when it is started through `anvien-server.exe` with hidden process attributes. The repeated visible-window report is more directly explained by unhidden `tasklist` calls inside `processAlive`.

## E2 - Initial Anvien Tool Evidence

Status: recorded

This section records tool output already gathered during initial inspection. It is historical evidence, not a standalone replacement for the active `AGENTS.md` workflow.

Anvien refresh:

- `anvien analyze --force`
  - Passed.
  - `files: scanned=715 parsed=538 unsupported=177 failed=0`
  - `graph: nodes=21821 relationships=54366`

Impact checks:

- `anvien impact handleLocalFolderPicker --repo "E:\Anvien" --direction upstream --depth 2 --include-tests`
  - Risk: LOW.
  - Impacted count: 0.
- `anvien impact pickCommandFolder --repo "E:\Anvien" --direction upstream --depth 2 --include-tests`
  - Risk: LOW.
  - Direct impacted symbols: `pickLocalFolder`, `pickWindowsFolder`.
- `anvien impact -u "Function:anvien-web/src/services/backend-client.ts:pickLocalFolder" --repo "E:\Anvien" --direction upstream --depth 2 --include-tests`
  - Risk: LOW.
  - Impacted scope includes `RepoAnalyzer.tsx`, `handleChooseRepository`, `AnalyzeOnboarding.tsx`, `Header.tsx`, `RepoLanding.tsx`, and `RepoAnalyzer.local-only.test.tsx`.
- `anvien impact RepoAnalyzer --repo "E:\Anvien" --direction upstream --depth 2 --include-tests`
  - Risk: LOW.
  - Impacted scope includes `AnalyzeOnboarding.tsx`, `Header.tsx`, `RepoLanding.tsx`, and related tests.

Implementation Anvien refresh and impact checks:

- `anvien analyze --force`
  - Passed.
  - `files: scanned=718 parsed=538 unsupported=180 failed=0`
  - `graph: nodes=21856 relationships=54407`
- `anvien impact RepoAnalyzer --repo "E:\Anvien" --direction upstream --depth 2 --include-tests`
  - Risk: LOW.
  - Impacted count: 9.
- `anvien impact -u "Function:anvien-web/src/services/backend-client.ts:pickLocalFolder" --repo "E:\Anvien" --direction upstream --depth 2 --include-tests`
  - Risk: LOW.
  - Impacted count: 6.
- `anvien impact handleLocalFolderPicker --repo "E:\Anvien" --direction upstream --depth 2 --include-tests`
  - Risk: LOW.
  - Impacted count: 0.
- `anvien impact -u "Function:internal/httpapi/local_folder_picker.go:pickLocalFolder#0" --repo "E:\Anvien" --direction upstream --depth 2 --include-tests`
  - Risk: LOW.
  - Impacted count: 0.
- `anvien impact pickWindowsFolder --repo "E:\Anvien" --direction upstream --depth 2 --include-tests`
  - Risk: LOW.
  - Impacted count: 1.
- `anvien impact pickCommandFolder --repo "E:\Anvien" --direction upstream --depth 2 --include-tests`
  - Risk: LOW.
  - Impacted count: 2.
- Launcher reset symbols were checked and warned as CRITICAL because they sit on the start/reset runtime process chain:
  - `processAlive`: CRITICAL, impacted count 5, affected reset/start runtime process flows.
  - `stopPID`: CRITICAL, impacted count 4, affected reset/start runtime process flows.
  - `stopRuntimeProcessesByPath`: CRITICAL, impacted count 3, affected reset process flows.
  - `registerProtocol`: CRITICAL, impacted count 1, affected protocol registration flow.
  - `openBrowser`: CRITICAL, impacted count 3, affected start runtime/open-browser flow.
  - Scope decision: changes were kept to hidden child process creation only; lifecycle waits and reset semantics were not changed.

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

Status: completed

Changed files and implementation details:

- `anvien-web/src/services/backend-client.ts`
  - `pickLocalFolder` now accepts an optional `AbortSignal` and passes it to `fetchFromBackend`.
- `anvien-web/src/components/RepoAnalyzer.tsx`
  - Added picker `AbortController` tracking.
  - Added explicit `Cancel Repository Picker` action while the picker request is pending.
  - Removed `isPickingFolder` from manual path submit eligibility.
  - Manual analysis aborts a pending picker before starting analyze.
  - Component unmount aborts pending picker and existing analysis SSE.
  - Wrapped `BackendError('Request aborted', status=0, code='network')` is treated as picker cancellation, not a user-facing error.
- `internal/httpapi/local_folder_picker.go`
  - `handleLocalFolderPicker` passes `r.Context()` into picker execution.
  - `pickLocalFolderFunc`, `pickLocalFolder`, `pickWindowsFolder`, and `pickCommandFolder` are context-aware.
  - OS picker commands use `exec.CommandContext`.
- `anvien-launcher/src/main.go`
  - Added `hiddenCommand` helper.
  - `tasklist` now goes through `tasklistCommand`, which uses hidden process attributes.
  - Windows process sweep, `taskkill`, protocol `reg`, and `rundll32` helpers now use the shared hidden command helper.
  - No launcher lifecycle wait or timeout semantics were changed.
- Tests added/updated:
  - `anvien-web/test/unit/RepoAnalyzer.local-only.test.tsx`
  - `internal/httpapi/handlers_test.go`
  - `anvien-launcher/src/main_test.go`
  - `anvien-web/e2e/onboarding.spec.ts`

Timeout check:

- `rg` over touched product files found no new picker/reset product timeout, timer reset, delayed reset, or elapsed-time recovery path.
- Existing launcher waits/timeouts remain unchanged.

## E5 - Validation Evidence

Status: completed

Validation after implementation in build-first order:

- Full packaged build:
  - Command: `powershell -ExecutionPolicy Bypass -File anvien-launcher\build.ps1`
  - Result: pass.
- Artifact gate:
  - `Start-Anvien.html=True`
  - `anvien-web\dist\Start-Anvien.html=False`
  - `anvien-launcher\web-dist\Start-Anvien.html=False`
  - Protocol command: `"E:\Anvien\anvien-launcher\AnvienLauncher.exe" "%1"`
- PE subsystem inspection after build:
  - `anvien-launcher\AnvienLauncher.exe=Windows GUI`
  - `anvien-launcher\server-bundle\anvien-server.exe=Windows GUI`
  - `anvien-launcher\server-bundle\anvien.exe=Windows CUI/Console`
- Focused frontend tests:
  - Command: `npm --prefix anvien-web test -- test/unit/RepoAnalyzer.local-only.test.tsx test/unit/analyze-contract.local-only.test.tsx`
  - Result: pass, 2 files, 11 tests.
- Focused backend picker tests:
  - Command: `go test ./internal/httpapi -run TestLocalFolderPicker -count=1`
  - Result: pass.
- Focused launcher/server-wrapper tests:
  - Command: `Push-Location anvien-launcher\src; go test . -count=1; Pop-Location; Push-Location anvien-launcher\server-wrapper; go test . -count=1; Pop-Location`
  - Result: pass.
- Web unit suite:
  - Command: `npm --prefix anvien-web test`
  - Result: pass, 43 files, 349 tests.
- Go runtime suite:
  - Command: `go test ./cmd/... ./internal/... -count=1`
  - Result: pass.
- Web e2e suite:
  - Setup: hidden `go run ./cmd/anvien serve --host 127.0.0.1 --port 4848` and hidden `npm --prefix anvien-web run dev`.
  - Command: `npm --prefix anvien-web run test:e2e`
  - Result: pass, 36 passed, 7 skipped.
  - Chooser coverage included `Flow 3: Analyze form > pending repository chooser can be cancelled before analyzing a pasted path`.
- Packaged reset validation:
  - Command: `Start-Process -FilePath .\anvien-launcher\AnvienLauncher.exe -ArgumentList reset -Wait -PassThru`
  - Result: `reset_process_exit=0`.
  - Post-reset runtime port check: `runtime_ports_listening=false` for ports `4848` and `5228`.
  - No terminal/helper window flashing was observed during this packaged reset validation run.
- Pre-commit Anvien detect changes:
  - Refresh: `anvien analyze --force`
    - Passed.
    - `files: scanned=718 parsed=538 unsupported=180 failed=0`
    - `graph: nodes=21896 relationships=54444`
  - Command: `anvien detect-changes --repo "E:\Anvien" --scope all`
  - Result: pass.
  - Summary: `changed_files=11`, `changed_count=69`, `affected_count=5`, `risk_level=medium`.
  - Affected process flows are expected for this plan: folder picker response paths and launcher start/reset hidden-process paths.
