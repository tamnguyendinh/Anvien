# Anvien Exe-Served Start Screen Evidence Ledger

Date: 2026-05-21

Status: completed

Plan: [2026-05-21-anvien-exe-served-start-screen-plan.md](2026-05-21-anvien-exe-served-start-screen-plan.md)

Benchmark: [2026-05-21-anvien-exe-served-start-screen-benchmark.md](2026-05-21-anvien-exe-served-start-screen-benchmark.md)

## E0 - User Clarification Evidence

Status: recorded

User clarified the desired product behavior:

- The exe currently opens the inner Web UI, but the expected first screen is the start surface represented by the old HTML.
- The start surface is the tool's entry point and must remain.
- The issue is the loose `Start-Anvien.html` file flow, not the existence of a start screen.
- `AnvienLauncher.exe` is rebuilt frequently and is the packaged user entrypoint.
- The Back button is a feature and must remain.

Interpretation:

- Move the start screen into the exe-served Web UI.
- Remove the root HTML file flow.
- Keep one packaged entrypoint: the rebuilt launcher exe.

## E1 - Current HTML Start Surface

Status: recorded

Source inspected from `HEAD:Start-Anvien.html`.

Current start surface contains:

- title: `Anvien`;
- primary button: `Start Anvien`;
- secondary button: `RESET RUNTIME`;
- secondary button: `User Guide`;
- user guide panel;
- status text with `aria-live="polite"`.

Current button behavior:

- `Start Anvien` sets status to `Starting Anvien...` and navigates to `anvien://start`.
- `RESET RUNTIME` sets status to `Resetting Anvien runtime...` and navigates to `anvien://reset`.
- `User Guide` toggles a guide panel and tries to fetch `user_guide.md`.

Implementation implication:

- The Web UI start screen must preserve the product role and user-visible actions.
- The new Start action must transition in-app instead of using `anvien://start`.
- Reset can continue to target the launcher exe path if hidden reset behavior is preserved.

## E2 - Current Launcher/Web Facts

Status: recorded

Files inspected:

- `anvien-launcher/src/main.go`
- `anvien-launcher/build.ps1`
- `anvien-web/src/App.tsx`
- `anvien-web/src/components/Header.tsx`
- `anvien-web/vite.config.ts`
- `README.md`
- `RUNBOOK.md`
- `TESTING.md`

Relevant facts:

- `anvien-launcher/build.ps1` rebuilds `anvien-launcher\AnvienLauncher.exe`.
- The same build copies `anvien-web\dist\` into `anvien-launcher\web-dist\`.
- The launcher opens the browser to the local Web UI on `127.0.0.1:5228`.
- The launcher currently has a static-handler path for root `Start-Anvien.html`.
- `Header` currently routes Back to `/Start-Anvien.html`.
- `App` currently starts at the onboarding/repo landing flow, not at a launcher start surface.
- Active docs still instruct users to open `Start-Anvien.html`.

## E3 - Interrupted Wrong-Direction Worktree

Status: recorded

Observed `git status --short` during plan creation:

- `D Start-Anvien.html`
- `M anvien-launcher/src/main.go`
- `M anvien-launcher/src/main_test.go`
- `M anvien-web/e2e/shell-interactions.spec.ts`
- `M anvien-web/src/App.tsx`
- `M anvien-web/src/components/Header.tsx`
- `M anvien-web/test/unit/Branding.local-only.test.tsx`
- `M anvien-web/vite.config.ts`

These changes came from an interrupted implementation attempt that removed the file flow but also removed the Back feature and did not migrate the start surface into the Web UI.

Implementation implication:

- Do not treat the interrupted code edits as accepted completion.
- The implementation must reconcile these edits: preserve the valid direction of removing the loose file flow, but restore Back and add the in-app start screen.

## E4 - Initial Reference Scan

Status: recorded

Initial scan found active references to `Start-Anvien.html` in:

- `README.md`
- `RUNBOOK.md`
- `TESTING.md`
- `CHANGELOG.md`
- `anvien-launcher/src/main.go`
- `anvien-launcher/src/main_test.go`
- `anvien-web/src/components/Header.tsx`
- `anvien-web/test/unit/Branding.local-only.test.tsx`
- `anvien-web/e2e/shell-interactions.spec.ts`
- historical plan/evidence/benchmark ledgers.

Implementation implication:

- Active code/docs/tests must be updated away from root `Start-Anvien.html`.
- Historical ledgers can remain historical unless they confuse active instructions.

## E5 - Implementation Evidence

Status: recorded

Implemented slices:

- P0 reconciled the interrupted wrong-direction edits by keeping the useful removal of loose HTML serving while restoring the Back feature as in-app navigation.
- P1 added `anvien-web/src/components/LauncherStartScreen.tsx` with `Anvien`, `Start Anvien`, `RESET RUNTIME`, `User Guide`, and status feedback.
- P1 updated `ViewMode` so the Web UI starts on the new in-app start screen.
- P1 wired `Start Anvien` to `setViewMode("onboarding")` instead of `Start-Anvien.html` or `anvien://start`.
- P1 made the User Guide panel fail gracefully when `user_guide.md` is unavailable.
- P2 restored the Header Back button and changed it to call an in-app `onNavigateToStart` callback.
- P2 updated App Back handling to clear graph/project state, close the right panel, remove `server`/`project` URL params, and return to the in-app start screen.
- P3 removed launcher special serving of the root HTML start file.
- P3 removed Vite dev-server special handling for the root HTML start file.
- P3 deleted the tracked root HTML start file.
- P4 updated active packaged launcher docs in `README.md`, `RUNBOOK.md`, `TESTING.md`, and `CHANGELOG.md`.

Tests added or updated:

- `anvien-web/test/unit/LauncherStartScreen.local-only.test.tsx`
- `anvien-web/test/unit/Branding.local-only.test.tsx`
- `anvien-web/e2e/onboarding.spec.ts`
- `anvien-web/e2e/shell-interactions.spec.ts`
- `anvien-launcher/src/main_test.go`

## E6 - Validation Evidence

Status: recorded

Validation commands and results:

- Full packaged build first: `powershell -ExecutionPolicy Bypass -File anvien-launcher\build.ps1` passed. It rebuilt Web dist, backend CLI, launcher exe, server wrapper, packaged Web dist, and protocol registration.
- Focused launcher Go tests: `go test ./...` from `anvien-launcher\src` passed.
- Focused Web unit tests: `npm test -- LauncherStartScreen.local-only.test.tsx Branding.local-only.test.tsx` passed with 2 files and 10 tests.
- Full Web unit tests: `npm test` passed with 44 files and 353 tests.
- Full Web e2e on production preview: `npm run test:e2e` passed with 15 passed and 29 skipped. Skips were runtime/backend-dependent smoke specs without an indexed local backend.
- Focused mocked Back e2e: `npm run test:e2e -- graph-health-ui.spec.ts` passed with 2 tests, including `graph shell Back returns to the exe-served start screen`.
- Packaged exe-served start screen check: started `anvien-launcher\AnvienLauncher.exe`, ran `npm run test:e2e -- onboarding.spec.ts -g "shows the exe-served start screen first"`, passed with 1 test, then stopped launcher runtime.
- Broader Go tests: `go test ./cmd/... ./internal/...` passed.
- Artifact check: root `Start-Anvien.html`, `anvien-web\dist\Start-Anvien.html`, and `anvien-launcher\web-dist\Start-Anvien.html` are absent.
- Artifact check: `anvien-launcher\AnvienLauncher.exe`, `anvien-launcher\web-dist\index.html`, `anvien-launcher\server-bundle\anvien-server.exe`, and `anvien-launcher\server-bundle\anvien.exe` exist.
- Active reference scan: `README.md`, `RUNBOOK.md`, `TESTING.md`, `CHANGELOG.md`, `anvien-launcher`, and `anvien-web` only retain `Start-Anvien.html` in stale-path tests.

## E7 - Change Detection and Commit Evidence

Status: recorded

Final change detection:

- Command: `anvien analyze --force`, then `detect_changes(repo: "Anvien", scope: "all")`.
- Result: `changed_files=18`, `changed_count=43`, `affected_count=7`, `risk_level=high`.
- Affected processes were the expected launcher runtime serving flows through `startRuntime`:
  - `StartRuntime -> HiddenProcAttr`
  - `StartRuntime -> AttachLog`
  - `StartRuntime -> UrlReady`
  - `StartRuntime -> BackendProcess`
  - `StartRuntime -> Done`
  - `StartRuntime -> LifecycleCheckInterval`
  - `StartRuntime -> WebLifecycleMonitor`

Assessment:

- The high risk is expected because the launcher static handler is on the packaged runtime startup path.
- The implementation intentionally limits launcher code changes to removing the loose root HTML special case and serving the packaged Web UI normally.
- Build, unit, e2e, packaged exe-served start-screen validation, artifact checks, and reference scans passed before commit.

Commit:

- Recorded in the implementation commit containing this ledger update.

## E8 - Reopened Regression Evidence

Status: recorded

User-provided failing screenshot:

- `reports/problem/screenshot_1779366896.png`

Observed from screenshot:

- URL: `127.0.0.1:5228`.
- Screen title: `Start Anvien locally`.
- The screen instructs the user to copy/run `anvien serve`.
- It shows `Waiting for local bridge to start` and `Listening for local bridge...`.

User clarification:

- This manual bridge guide should not appear after clicking `Start Anvien`.
- The manual guide screen is unnecessary and should be removed from active Web UI behavior.
- In the packaged flow, the exe should already start the runtime/server bridge.
- Start should lead to repository chooser/analyze flow, not manual startup instructions.

Implementation implication:

- Reopen this plan.
- Remove `OnboardingGuide` from active Web UI routing.
- Replace backend-unavailable/manual bridge fallback with a neutral runtime connection/recovery state that does not mention `anvien serve`.
- Add visible-browser validation so the user can observe the tested behavior on the PC.

## E9 - Follow-Up Implementation Evidence

Status: recorded

Implemented reopened slices:

- Removed `OnboardingGuide` from active Web UI routing in `DropZone`.
- Deleted the retired manual guide component file and its unit test:
  - `anvien-web/src/components/OnboardingGuide.tsx`
  - `anvien-web/test/unit/OnboardingGuide.local-only.test.tsx`
- Changed backend-unavailable rendering to a neutral `Connecting to Anvien runtime...` state.
- Kept the existing runtime/backend probe flow; did not add product timeout, retry-loop, delayed navigation, or elapsed-time budget logic.
- Updated onboarding e2e coverage so backend-unavailable state asserts absence of `Start Anvien locally`, `anvien serve`, terminal copy guidance, and local bridge waiting copy.
- Added `DropZone.runtime-connection.test.tsx` to prove the active fallback does not render the retired manual guide.
- Added a packaged launcher smoke e2e gated by `PACKAGED_LAUNCHER_E2E=1` for visible browser validation against `AnvienLauncher.exe`.

## E10 - Follow-Up Validation Evidence

Status: recorded

Validation commands and results:

- Full packaged build first: `powershell -ExecutionPolicy Bypass -File anvien-launcher\build.ps1` passed.
- Focused Web unit tests: `npm test -- DropZone.runtime-connection.test.tsx DropZone.full-analyze-flow.test.tsx LauncherStartScreen.local-only.test.tsx Branding.local-only.test.tsx` passed with 4 files and 15 tests.
- Full Web unit tests: `npm test` passed with 44 files and 353 tests.
- Full Web e2e on production preview: `npm run test:e2e` passed with 11 passed and 30 skipped. Skips were backend/indexed-repo dependent specs, including the packaged launcher smoke test when `PACKAGED_LAUNCHER_E2E` is not set.
- Onboarding e2e file check on production preview: `npm run test:e2e -- e2e/onboarding.spec.ts` passed with 9 passed and 3 skipped.
- Visible packaged launcher validation: started `anvien-launcher\AnvienLauncher.exe`, then ran `PACKAGED_LAUNCHER_E2E=1 npm run test:e2e -- e2e/onboarding.spec.ts -g "packaged launcher Start reaches repo chooser or analyze without manual guide" --headed`; passed with 1 test.
- Visible validation screenshot: `anvien-web/test-results/onboarding-Flow-1-Start-sc-396b2-nalyze-without-manual-guide-chromium/packaged-start-target.png`.
- Artifact check: root `Start-Anvien.html`, `anvien-web\dist\Start-Anvien.html`, and `anvien-launcher\web-dist\Start-Anvien.html` are absent.
- Artifact check: `anvien-launcher\AnvienLauncher.exe` and `anvien-launcher\web-dist\index.html` exist.
- Production Go package tests: `go test ./cmd/... ./internal/...` passed.
- Launcher Go tests: `go test ./...` from `anvien-launcher\src` passed.

Additional Go validation note:

- `go test ./...` from the repo root was attempted and failed only in fixture packages under `anvien/test/fixtures/...` that intentionally contain non-buildable sample code/imports. Production `cmd`/`internal` packages and launcher tests passed.

## E11 - Follow-Up Change Detection Evidence

Status: recorded

Pre-commit Anvien checks:

- Refreshed graph with `anvien analyze --force`; no `--skip-agents-md` flag was used.
- `impact` on `DropZone`, `OnboardingGuide`, `App`, `LoadingCard`, `SuccessCard`, and `handleAutoConnect` was LOW.
- `detect_changes(repo: "Anvien", scope: "all")` reported `changed_files=7`, `changed_count=30`, `affected_count=7`, `risk_level=high`.
- Affected processes were the expected `handleAutoConnect` flows:
  - `HandleAutoConnect -> BackendError`
  - `HandleAutoConnect -> AssertOk`
  - `HandleAutoConnect -> FetchFromBackend`
  - `HandleAutoConnect -> ProbeBackend`
  - `HandleAutoConnect -> SetBackendUrl`
  - `HandleAutoConnect -> FetchGraph`
  - `HandleAutoConnect -> FetchRepoInfo`

Assessment:

- The high detect risk is expected because this slice intentionally changes the backend-unavailable branch in the Start runtime path.
- The happy paths remain covered by mocked e2e zero-repo and repo-landing flows, plus headed packaged launcher validation.

## E12 - README User Guide Evidence

Status: recorded

User clarification:

- Clicking the start screen `User Guide` button should display the contents of repository `README.md`.

Implementation:

- `LauncherStartScreen` now fetches `/README.md` instead of `/user_guide.md`.
- `vite.config.ts` adds a root README plugin that:
  - serves root `README.md` at `/README.md` in Vite dev mode;
  - copies root `README.md` into `anvien-web\dist\README.md` during Web build.
- The packaged launcher build copies the Web dist to `anvien-launcher\web-dist\`, so `README.md` is available to `AnvienLauncher.exe`.

Validation:

- First full build attempt failed because a previous packaged runtime process held `anvien-launcher\server-bundle\anvien-server.exe` locked. The repo-local launcher/server-bundle processes were stopped and the full build was rerun.
- Full packaged build: `powershell -ExecutionPolicy Bypass -File anvien-launcher\build.ps1` passed.
- Artifact check: `anvien-web\dist\README.md` and `anvien-launcher\web-dist\README.md` exist, both matching root README size at validation time.
- Focused unit: `npm test -- LauncherStartScreen.local-only.test.tsx` passed with 1 file and 5 tests.
- Full Web unit: `npm test` passed with 44 files and 354 tests.
- Focused e2e on production preview: `npm run test:e2e -- e2e/onboarding.spec.ts -g "User Guide displays README.md content"` passed.
- Full Web e2e on production preview: `npm run test:e2e` passed with 12 passed and 30 skipped.
- Packaged launcher e2e: started `AnvienLauncher.exe` with browser auto-open suppressed, then ran the focused User Guide e2e against `127.0.0.1:5228`; passed.
- Pre-commit change detection: after `anvien analyze --force`, `detect_changes(repo: "Anvien", scope: "all")` reported `changed_files=7`, `changed_count=25`, `affected_count=0`, `risk_level=low`.
