# AVmatrix Exe-Served Start Screen Evidence Ledger

Date: 2026-05-21

Status: planned

Plan: [2026-05-21-avmatrix-exe-served-start-screen-plan.md](2026-05-21-avmatrix-exe-served-start-screen-plan.md)

Benchmark: [2026-05-21-avmatrix-exe-served-start-screen-benchmark.md](2026-05-21-avmatrix-exe-served-start-screen-benchmark.md)

## E0 - User Clarification Evidence

Status: recorded

User clarified the desired product behavior:

- The exe currently opens the inner Web UI, but the expected first screen is the start surface represented by the old HTML.
- The start surface is the tool's entry point and must remain.
- The issue is the loose `Start-AVmatrix.html` file flow, not the existence of a start screen.
- `AVmatrixLauncher.exe` is rebuilt frequently and is the packaged user entrypoint.
- The Back button is a feature and must remain.

Interpretation:

- Move the start screen into the exe-served Web UI.
- Remove the root HTML file flow.
- Keep one packaged entrypoint: the rebuilt launcher exe.

## E1 - Current HTML Start Surface

Status: recorded

Source inspected from `HEAD:Start-AVmatrix.html`.

Current start surface contains:

- title: `AVmatrix`;
- primary button: `Start AVmatrix`;
- secondary button: `RESET RUNTIME`;
- secondary button: `User Guide`;
- user guide panel;
- status text with `aria-live="polite"`.

Current button behavior:

- `Start AVmatrix` sets status to `Starting AVmatrix...` and navigates to `avmatrix://start`.
- `RESET RUNTIME` sets status to `Resetting AVmatrix runtime...` and navigates to `avmatrix://reset`.
- `User Guide` toggles a guide panel and tries to fetch `user_guide.md`.

Implementation implication:

- The Web UI start screen must preserve the product role and user-visible actions.
- The new Start action must transition in-app instead of using `avmatrix://start`.
- Reset can continue to target the launcher exe path if hidden reset behavior is preserved.

## E2 - Current Launcher/Web Facts

Status: recorded

Files inspected:

- `avmatrix-launcher/src/main.go`
- `avmatrix-launcher/build.ps1`
- `avmatrix-web/src/App.tsx`
- `avmatrix-web/src/components/Header.tsx`
- `avmatrix-web/vite.config.ts`
- `README.md`
- `RUNBOOK.md`
- `TESTING.md`

Relevant facts:

- `avmatrix-launcher/build.ps1` rebuilds `avmatrix-launcher\AVmatrixLauncher.exe`.
- The same build copies `avmatrix-web\dist\` into `avmatrix-launcher\web-dist\`.
- The launcher opens the browser to the local Web UI on `127.0.0.1:5228`.
- The launcher currently has a static-handler path for root `Start-AVmatrix.html`.
- `Header` currently routes Back to `/Start-AVmatrix.html`.
- `App` currently starts at the onboarding/repo landing flow, not at a launcher start surface.
- Active docs still instruct users to open `Start-AVmatrix.html`.

## E3 - Interrupted Wrong-Direction Worktree

Status: recorded

Observed `git status --short` during plan creation:

- `D Start-AVmatrix.html`
- `M avmatrix-launcher/src/main.go`
- `M avmatrix-launcher/src/main_test.go`
- `M avmatrix-web/e2e/shell-interactions.spec.ts`
- `M avmatrix-web/src/App.tsx`
- `M avmatrix-web/src/components/Header.tsx`
- `M avmatrix-web/test/unit/Branding.local-only.test.tsx`
- `M avmatrix-web/vite.config.ts`

These changes came from an interrupted implementation attempt that removed the file flow but also removed the Back feature and did not migrate the start surface into the Web UI.

Implementation implication:

- Do not treat the interrupted code edits as accepted completion.
- The implementation must reconcile these edits: preserve the valid direction of removing the loose file flow, but restore Back and add the in-app start screen.

## E4 - Initial Reference Scan

Status: recorded

Initial scan found active references to `Start-AVmatrix.html` in:

- `README.md`
- `RUNBOOK.md`
- `TESTING.md`
- `CHANGELOG.md`
- `avmatrix-launcher/src/main.go`
- `avmatrix-launcher/src/main_test.go`
- `avmatrix-web/src/components/Header.tsx`
- `avmatrix-web/test/unit/Branding.local-only.test.tsx`
- `avmatrix-web/e2e/shell-interactions.spec.ts`
- historical plan/evidence/benchmark ledgers.

Implementation implication:

- Active code/docs/tests must be updated away from root `Start-AVmatrix.html`.
- Historical ledgers can remain historical unless they confuse active instructions.

## E5 - Implementation Evidence

Status: recorded

Implemented slices:

- P0 reconciled the interrupted wrong-direction edits by keeping the useful removal of loose HTML serving while restoring the Back feature as in-app navigation.
- P1 added `avmatrix-web/src/components/LauncherStartScreen.tsx` with `AVmatrix`, `Start AVmatrix`, `RESET RUNTIME`, `User Guide`, and status feedback.
- P1 updated `ViewMode` so the Web UI starts on the new in-app start screen.
- P1 wired `Start AVmatrix` to `setViewMode("onboarding")` instead of `Start-AVmatrix.html` or `avmatrix://start`.
- P1 made the User Guide panel fail gracefully when `user_guide.md` is unavailable.
- P2 restored the Header Back button and changed it to call an in-app `onNavigateToStart` callback.
- P2 updated App Back handling to clear graph/project state, close the right panel, remove `server`/`project` URL params, and return to the in-app start screen.
- P3 removed launcher special serving of the root HTML start file.
- P3 removed Vite dev-server special handling for the root HTML start file.
- P3 deleted the tracked root HTML start file.
- P4 updated active packaged launcher docs in `README.md`, `RUNBOOK.md`, `TESTING.md`, and `CHANGELOG.md`.

Tests added or updated:

- `avmatrix-web/test/unit/LauncherStartScreen.local-only.test.tsx`
- `avmatrix-web/test/unit/Branding.local-only.test.tsx`
- `avmatrix-web/e2e/onboarding.spec.ts`
- `avmatrix-web/e2e/shell-interactions.spec.ts`
- `avmatrix-launcher/src/main_test.go`

## E6 - Validation Evidence

Status: recorded

Validation commands and results:

- Full packaged build first: `powershell -ExecutionPolicy Bypass -File avmatrix-launcher\build.ps1` passed. It rebuilt Web dist, backend CLI, launcher exe, server wrapper, packaged Web dist, and protocol registration.
- Focused launcher Go tests: `go test ./...` from `avmatrix-launcher\src` passed.
- Focused Web unit tests: `npm test -- LauncherStartScreen.local-only.test.tsx Branding.local-only.test.tsx` passed with 2 files and 10 tests.
- Full Web unit tests: `npm test` passed with 44 files and 353 tests.
- Full Web e2e on production preview: `npm run test:e2e` passed with 15 passed and 29 skipped. Skips were runtime/backend-dependent smoke specs without an indexed local backend.
- Focused mocked Back e2e: `npm run test:e2e -- graph-health-ui.spec.ts` passed with 2 tests, including `graph shell Back returns to the exe-served start screen`.
- Packaged exe-served start screen check: started `avmatrix-launcher\AVmatrixLauncher.exe`, ran `npm run test:e2e -- onboarding.spec.ts -g "shows the exe-served start screen first"`, passed with 1 test, then stopped launcher runtime.
- Broader Go tests: `go test ./cmd/... ./internal/...` passed.
- Artifact check: root `Start-AVmatrix.html`, `avmatrix-web\dist\Start-AVmatrix.html`, and `avmatrix-launcher\web-dist\Start-AVmatrix.html` are absent.
- Artifact check: `avmatrix-launcher\AVmatrixLauncher.exe`, `avmatrix-launcher\web-dist\index.html`, `avmatrix-launcher\server-bundle\avmatrix-server.exe`, and `avmatrix-launcher\server-bundle\avmatrix.exe` exist.
- Active reference scan: `README.md`, `RUNBOOK.md`, `TESTING.md`, `CHANGELOG.md`, `avmatrix-launcher`, and `avmatrix-web` only retain `Start-AVmatrix.html` in stale-path tests.

## E7 - Change Detection and Commit Evidence

Status: recorded

Final change detection:

- Command: `avmatrix analyze --force`, then `detect_changes(repo: "AVmatrix", scope: "all")`.
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
