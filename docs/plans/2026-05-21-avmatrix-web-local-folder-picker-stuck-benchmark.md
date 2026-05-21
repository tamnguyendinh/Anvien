# AVmatrix Web Local Folder Picker Stuck Benchmark Ledger

Date: 2026-05-21

Status: completed

Plan: [2026-05-21-avmatrix-web-local-folder-picker-stuck-plan.md](2026-05-21-avmatrix-web-local-folder-picker-stuck-plan.md)

Evidence: [2026-05-21-avmatrix-web-local-folder-picker-stuck-evidence.md](2026-05-21-avmatrix-web-local-folder-picker-stuck-evidence.md)

## B0 - Required Baseline

Status: completed

Capture before implementation:

- failing screenshot path;
- UI label while stuck;
- whether manual path input remains editable;
- whether Analyze Repository is blocked while picker is pending;
- whether the frontend request to `/api/local/folder-picker` remains pending;
- whether any backend picker child process remains after cancel/unmount when tested;
- current unit/e2e coverage gap.
- current `Start-AVmatrix.html` artifact location and whether dist copies exist;
- selected build-first validation commands.

Baseline table:

| Metric | Current value | Evidence |
| --- | --- | --- |
| Failing screenshot | `reports/problem/screenshot_1779328050.png` | user-provided |
| Visible stuck text | `Opening repository picker...` | screenshot |
| Recovery action visible | no before fix; yes after fix | screenshot, unit/e2e |
| Analyze from pasted path while picker pending | blocked before fix; yes after fix | code inspection, unit/e2e |
| Backend request cancellation kills picker command | yes after fix | `exec.CommandContext`, backend test |
| Reset runtime flashes terminal/helper windows | reported | user report |
| Reset runtime helper processes run hidden | yes after fix | launcher tests, packaged reset validation |
| `processAlive` Windows `tasklist` hidden | no | code inspection |
| Launcher executable subsystem | GUI | PE header inspection |
| Server wrapper executable subsystem | GUI | PE header inspection |
| Backend executable subsystem | Console | PE header inspection |
| `Start-AVmatrix.html` root entry exists | yes | post-build artifact gate |
| `Start-AVmatrix.html` absent from `avmatrix-web\dist` | yes | post-build artifact gate |
| `Start-AVmatrix.html` absent from `avmatrix-launcher\web-dist` | yes | post-build artifact gate |
| Current protocol command | `E:\AVmatrix-GO\avmatrix-launcher\AVmatrixLauncher.exe "%1"` | registry inspection |

## B1 - Frontend Recovery Benchmark

Status: completed

After implementation, record:

| Metric | Expected |
| --- | --- |
| Pending picker shows cancel action | yes |
| Cancel action clears pending state | yes |
| Cancel action does not show false error | yes |
| Analyze button remains usable for valid pasted path while picker is pending | yes |
| Pasted valid path can be analyzed after pending picker starts | yes |
| Component unmount aborts picker request | yes |
| Product timeout added | no |

## B2 - Backend Cancellation Benchmark

Status: completed

After implementation, record:

| Metric | Expected |
| --- | --- |
| `handleLocalFolderPicker` uses request context | yes |
| OS picker command uses `exec.CommandContext` | yes |
| Explicit picker cancel still returns `{ path: null, cancelled: true }` | yes |
| Unsupported picker still returns not implemented | yes |
| Request cancellation can stop picker execution | yes |
| Product timeout added | no |

## B3 - Validation Benchmark

Status: completed

Record validation commands and results in this order:

| Command | Expected |
| --- | --- |
| `powershell -ExecutionPolicy Bypass -File avmatrix-launcher\build.ps1` | pass |
| Packaged artifact location gate | pass; root present and dist copies absent |
| `npm --prefix avmatrix-web test -- test/unit/RepoAnalyzer.local-only.test.tsx test/unit/analyze-contract.local-only.test.tsx` | pass; 11 tests |
| `go test ./internal/httpapi -run TestLocalFolderPicker -count=1` | pass |
| `Push-Location avmatrix-launcher\src; go test . -count=1; Pop-Location; Push-Location avmatrix-launcher\server-wrapper; go test . -count=1; Pop-Location` | pass |
| `npm --prefix avmatrix-web test` | pass; 43 files, 349 tests |
| `go test ./cmd/... ./internal/... -count=1` | pass |
| `npm --prefix avmatrix-web run test:e2e` | pass; 36 passed, 7 skipped |
| Packaged/manual reset validation | pass; reset exit `0`, runtime ports `4848`/`5228` not listening afterward |

## B4 - Runtime Reset Hidden Execution Benchmark

Status: completed

After implementation, record:

| Metric | Expected |
| --- | --- |
| `Start-AVmatrix.html` reset protocol still triggers reset | yes |
| `processAlive` `tasklist` command hidden on Windows | yes |
| Reset process sweep command hidden on Windows | yes |
| `taskkill` soft/force commands hidden on Windows | yes |
| Server wrapper backend launch hidden on Windows | yes |
| Packaged launcher/protocol entry avoids visible console window | yes |
| Backend console executable is only launched by hidden server-wrapper path | yes |
| `Start-AVmatrix.html` remains root-only user-facing entry | yes |
| Logs written to files/status UI instead of visible terminal windows | yes |
| Reset still allows runtime to start again | yes |
| New product timeout/timer recovery added | no |

## B5 - Final Interpretation

Status: completed

The final benchmark passes only if the app no longer leaves the user trapped behind a pending picker request, reset runtime does not flash visible terminal/helper windows in the packaged user-facing path, and the fix does not add timeout/timer-based product recovery behavior.

Final result: pass.
