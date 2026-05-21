# AVmatrix Exe-Served Start Screen Benchmark Ledger

Date: 2026-05-21

Status: planned

Plan: [2026-05-21-avmatrix-exe-served-start-screen-plan.md](2026-05-21-avmatrix-exe-served-start-screen-plan.md)

Evidence: [2026-05-21-avmatrix-exe-served-start-screen-evidence.md](2026-05-21-avmatrix-exe-served-start-screen-evidence.md)

## B0 - Required Baseline

Status: recorded

Baseline product facts:

| Metric | Current / observed value | Source |
| --- | --- | --- |
| Packaged exe is rebuilt by full build | yes | `avmatrix-launcher/build.ps1` |
| Packaged exe path | `avmatrix-launcher\AVmatrixLauncher.exe` | build script |
| Packaged Web dist path | `avmatrix-launcher\web-dist\` | build script |
| Root HTML start file exists in `HEAD` | yes | `HEAD:Start-AVmatrix.html` |
| Root HTML Start button uses `avmatrix://start` | yes | `HEAD:Start-AVmatrix.html` |
| Root HTML Reset button uses `avmatrix://reset` | yes | `HEAD:Start-AVmatrix.html` |
| Web app initial screen is in-app start surface | no | `avmatrix-web/src/App.tsx` inspection |
| Header Back targets root HTML path | yes | `avmatrix-web/src/components/Header.tsx` inspection |
| Active docs instruct opening root HTML | yes | `README.md`, `RUNBOOK.md`, `TESTING.md` |
| Interrupted worktree removed Back | yes, but not accepted | plan-time `git status` and code inspection |

## B1 - Target Product Benchmarks

Status: pending

Record after implementation:

| Metric | Expected |
| --- | --- |
| Running rebuilt `AVmatrixLauncher.exe` opens in-app start screen first | yes |
| Start screen is served from packaged Web UI, not root HTML | yes |
| Start screen includes `AVmatrix` title | yes |
| Start screen includes `Start AVmatrix` | yes |
| Start screen includes `RESET RUNTIME` | yes |
| Start screen includes `User Guide` | yes |
| Start screen has status feedback | yes |
| `Start AVmatrix` enters repo landing/analyze flow in-app | yes |
| `Start AVmatrix` uses `avmatrix://start` | no |
| Header Back button exists in graph shell | yes |
| Header Back returns to in-app start screen | yes |
| Header Back navigates to `/Start-AVmatrix.html` | no |
| Root `Start-AVmatrix.html` tracked file exists | no |
| Root `Start-AVmatrix.html` appears in `avmatrix-web\dist` | no |
| Root `Start-AVmatrix.html` appears in `avmatrix-launcher\web-dist` | no |
| Users are instructed to open root HTML in active docs | no |

## B2 - Artifact Benchmarks

Status: pending

Record after full build:

| Artifact check | Expected |
| --- | --- |
| `avmatrix-launcher\AVmatrixLauncher.exe` exists after build | yes |
| `avmatrix-launcher\web-dist\index.html` exists after build | yes |
| `Start-AVmatrix.html` absent from repo root | yes |
| `avmatrix-web\dist\Start-AVmatrix.html` absent | yes |
| `avmatrix-launcher\web-dist\Start-AVmatrix.html` absent | yes |
| packaged Web bundle contains start screen code | yes |

## B3 - Validation Benchmarks

Status: pending

Record commands and results after implementation:

| Validation | Expected |
| --- | --- |
| Full packaged build before tests | pass |
| Focused launcher Go tests | pass |
| Focused Web unit tests for start screen/Header | pass |
| Full Web unit tests | pass |
| Web e2e start/back flow | pass |
| Broader Go tests for touched packages | pass |
| Active reference scan for `Start-AVmatrix.html` and `avmatrix://start` | only stale-path tests or historical ledgers remain |
| Required change detection before commit | recorded |

## B4 - Runtime UX Benchmarks

Status: pending

Manual or automated packaged runtime checks:

| UX check | Expected |
| --- | --- |
| Browser opens to start screen after launching exe | yes |
| Start button reaches repo landing/analyze flow | yes |
| Graph shell Back returns to start screen | yes |
| Reset Runtime action does not open visible terminal/helper windows | yes |
| User Guide action does not dead-end if guide file is missing | yes |
| Reloading the app at root shows start screen unless a deliberate project/server URL is provided | yes |
