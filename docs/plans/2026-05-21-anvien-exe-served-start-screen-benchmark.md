# Anvien Exe-Served Start Screen Benchmark Ledger

Date: 2026-05-21

Status: completed

Plan: [2026-05-21-anvien-exe-served-start-screen-plan.md](2026-05-21-anvien-exe-served-start-screen-plan.md)

Evidence: [2026-05-21-anvien-exe-served-start-screen-evidence.md](2026-05-21-anvien-exe-served-start-screen-evidence.md)

## B0 - Required Baseline

Status: recorded

Baseline product facts:

| Metric | Current / observed value | Source |
| --- | --- | --- |
| Packaged exe is rebuilt by full build | yes | `anvien-launcher/build.ps1` |
| Packaged exe path | `anvien-launcher\AnvienLauncher.exe` | build script |
| Packaged Web dist path | `anvien-launcher\web-dist\` | build script |
| Root HTML start file exists in `HEAD` | yes | `HEAD:Start-Anvien.html` |
| Root HTML Start button uses `anvien://start` | yes | `HEAD:Start-Anvien.html` |
| Root HTML Reset button uses `anvien://reset` | yes | `HEAD:Start-Anvien.html` |
| Web app initial screen is in-app start surface | no | `anvien-web/src/App.tsx` inspection |
| Header Back targets root HTML path | yes | `anvien-web/src/components/Header.tsx` inspection |
| Active docs instruct opening root HTML | yes | `README.md`, `RUNBOOK.md`, `TESTING.md` |
| Interrupted worktree removed Back | yes, but not accepted | plan-time `git status` and code inspection |

## B1 - Target Product Benchmarks

Status: recorded from implementation; reopened runtime Start validation pending in B5

Record after implementation:

| Metric | Expected |
| --- | --- |
| Running rebuilt `AnvienLauncher.exe` opens in-app start screen first | yes; validated before reopened Start-click regression |
| Start screen is served from packaged Web UI, not root HTML | yes; implemented as `LauncherStartScreen` |
| Start screen includes `Anvien` title | yes |
| Start screen includes `Start Anvien` | yes |
| Start screen includes `RESET RUNTIME` | yes |
| Start screen includes `User Guide` | yes |
| Start screen has status feedback | yes |
| `Start Anvien` enters repo landing/analyze flow in-app | yes |
| `Start Anvien` uses `anvien://start` | no |
| Header Back button exists in graph shell | yes |
| Header Back returns to in-app start screen | yes |
| Header Back navigates to `/Start-Anvien.html` | no |
| Root `Start-Anvien.html` tracked file exists | no |
| Root `Start-Anvien.html` appears in `anvien-web\dist` | no; verified in B2 |
| Root `Start-Anvien.html` appears in `anvien-launcher\web-dist` | no; verified in B2 |
| Users are instructed to open root HTML in active docs | no |

## B2 - Artifact Benchmarks

Status: completed

Record after full build:

| Artifact check | Expected |
| --- | --- |
| `anvien-launcher\AnvienLauncher.exe` exists after build | yes |
| `anvien-launcher\web-dist\index.html` exists after build | yes |
| `Start-Anvien.html` absent from repo root | yes |
| `anvien-web\dist\Start-Anvien.html` absent | yes |
| `anvien-launcher\web-dist\Start-Anvien.html` absent | yes |
| packaged Web bundle contains start screen code | yes; verified through packaged exe-served e2e |

## B3 - Validation Benchmarks

Status: completed

Record commands and results after implementation:

| Validation | Expected |
| --- | --- |
| Full packaged build before tests | pass |
| Focused launcher Go tests | pass |
| Focused Web unit tests for start screen/Header | pass |
| Full Web unit tests | pass |
| Web e2e start/back flow | pass |
| Broader Go tests for touched packages | pass |
| Active reference scan for `Start-Anvien.html` and `anvien://start` | only stale-path tests and historical ledgers remain |
| Required change detection before commit | recorded; high risk expected on launcher startup path |

## B4 - Runtime UX Benchmarks

Status: completed where automatable in this workspace

Manual or automated packaged runtime checks:

| UX check | Expected |
| --- | --- |
| Browser opens to start screen after launching exe | yes; packaged exe-served e2e passed |
| Start button reaches repo landing/analyze flow | yes; onboarding e2e passed |
| Graph shell Back returns to start screen | yes; mocked graph e2e passed |
| Reset Runtime action does not open visible terminal/helper windows | not changed in this plan; previous hidden-reset validation remains applicable |
| User Guide action does not dead-end if guide file is missing | yes; unit test passed |
| Reloading the app at root shows start screen unless a deliberate project/server URL is provided | yes; e2e root start and query-param graph flows passed |

## B5 - Reopened Runtime Start Benchmarks

Status: completed

Record after the follow-up implementation:

| UX check | Expected |
| --- | --- |
| Clicking `Start Anvien` shows `Start Anvien locally` | no; unit, e2e, and headed packaged validation passed |
| Clicking `Start Anvien` shows `anvien serve` command instructions | no; unit, e2e, and headed packaged validation passed |
| `OnboardingGuide` remains in active Web UI routing | no; component import removed and retired file deleted |
| Start reaches repo landing when repos exist | yes; mocked e2e and headed packaged launcher validation passed |
| Start reaches Analyze Repository when no repos exist | yes; mocked e2e zero-repo flow passed |
| Backend unavailable state mentions manual terminal commands | no; neutral runtime connection e2e passed |
| Visible browser validation was performed on the user's PC | yes; headed Playwright against `AnvienLauncher.exe` passed |
| Visible browser screenshot or observation recorded | yes; `anvien-web/test-results/onboarding-Flow-1-Start-sc-396b2-nalyze-without-manual-guide-chromium/packaged-start-target.png` |

## B6 - README User Guide Benchmarks

Status: completed

| UX/artifact check | Expected |
| --- | --- |
| User Guide fetch target | `/README.md` |
| User Guide displays README heading | yes; unit and e2e passed |
| `anvien-web\dist\README.md` exists after full build | yes |
| `anvien-launcher\web-dist\README.md` exists after full build | yes |
| Packaged launcher serves README-backed User Guide | yes; focused packaged e2e passed |
