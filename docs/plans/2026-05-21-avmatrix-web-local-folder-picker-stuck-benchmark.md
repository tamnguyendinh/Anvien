# AVmatrix Web Local Folder Picker Stuck Benchmark Ledger

Date: 2026-05-21

Status: planned

Plan: [2026-05-21-avmatrix-web-local-folder-picker-stuck-plan.md](2026-05-21-avmatrix-web-local-folder-picker-stuck-plan.md)

Evidence: [2026-05-21-avmatrix-web-local-folder-picker-stuck-evidence.md](2026-05-21-avmatrix-web-local-folder-picker-stuck-evidence.md)

## B0 - Required Baseline

Status: planned

Capture before implementation:

- failing screenshot path;
- UI label while stuck;
- whether manual path input remains editable;
- whether Analyze Repository is blocked while picker is pending;
- whether the frontend request to `/api/local/folder-picker` remains pending;
- whether any backend picker child process remains after cancel/unmount when tested;
- current unit/e2e coverage gap.

Baseline table:

| Metric | Current value | Evidence |
| --- | --- | --- |
| Failing screenshot | `reports/problem/screenshot_1779328050.png` | user-provided |
| Visible stuck text | `Opening repository picker...` | screenshot |
| Recovery action visible | TBD | inspect/reproduce |
| Analyze from pasted path while picker pending | TBD | inspect/reproduce |
| Backend request cancellation kills picker command | TBD | test after implementation |
| Reset runtime flashes terminal/helper windows | reported | user report |
| Reset runtime helper processes run hidden | TBD | inspect/reproduce |
| `processAlive` Windows `tasklist` hidden | no | code inspection |
| Launcher executable subsystem | GUI | PE header inspection |
| Server wrapper executable subsystem | GUI | PE header inspection |
| Backend executable subsystem | Console | PE header inspection |
| Current protocol command | `E:\AVmatrix-GO\avmatrix-launcher\AVmatrixLauncher.exe "%1"` | registry inspection |

## B1 - Frontend Recovery Benchmark

Status: planned

After implementation, record:

| Metric | Expected |
| --- | --- |
| Pending picker shows cancel action | yes |
| Cancel action clears pending state | yes |
| Cancel action does not show false error | yes |
| Pasted valid path can be analyzed after pending picker starts | yes |
| Component unmount aborts picker request | yes |
| Product timeout added | no |

## B2 - Backend Cancellation Benchmark

Status: planned

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

Status: planned

Record validation commands and results:

| Command | Expected |
| --- | --- |
| Focused `RepoAnalyzer` tests | pass |
| Focused `local_folder_picker` Go tests | pass |
| Focused launcher/server-wrapper hidden process tests | pass |
| `npm --prefix avmatrix-web run build` | pass |
| Applicable Web unit tests | pass |
| Applicable Go tests | pass |
| Launcher/server-wrapper builds | pass |
| Applicable e2e chooser flow | pass |
| Packaged/manual reset validation | no visible terminal/helper window flashing |

## B4 - Runtime Reset Hidden Execution Benchmark

Status: planned

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
| Logs written to files/status UI instead of visible terminal windows | yes |
| Reset still allows runtime to start again | yes |
| New product timeout/timer recovery added | no |

## B5 - Final Interpretation

Status: planned

The final benchmark passes only if the app no longer leaves the user trapped behind a pending picker request, reset runtime does not flash visible terminal/helper windows in the packaged user-facing path, and the fix does not add timeout/timer-based product recovery behavior.
