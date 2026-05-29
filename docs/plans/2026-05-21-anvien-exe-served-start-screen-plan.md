# Anvien Exe-Served Start Screen Plan

Date: 2026-05-21

Status: completed

Companion files:

- Benchmark ledger: [2026-05-21-anvien-exe-served-start-screen-benchmark.md](2026-05-21-anvien-exe-served-start-screen-benchmark.md)
- Evidence ledger: [2026-05-21-anvien-exe-served-start-screen-evidence.md](2026-05-21-anvien-exe-served-start-screen-evidence.md)

## Rules

1. Follow active workspace and repository instructions, including `AGENTS.md`, for implementation workflow. This plan records product work and validation; it does not replace those rules.
2. Use Anvien according to the active repo instructions for implementation slices; do not use Anvien for doc-only commits.
3. Update this plan, the benchmark ledger, and the evidence ledger as each implementation slice is completed.
4. Run the full packaged build before tests. For this plan the full build gate is `powershell -ExecutionPolicy Bypass -File anvien-launcher\build.ps1`.
5. Because this changes Web UI behavior, validation must include focused unit tests and Web e2e tests after the full build passes.
6. Do not add product/runtime timeout, delayed navigation, or elapsed-time budget as the fix mechanism.
7. Keep the Back button as a product feature.
8. Remove the separate `Start-Anvien.html` file flow; do not replace it with another loose root HTML entrypoint.

## Problem

Before this plan, the packaged launcher had two user-facing start paths:

- double-clicking or running `anvien-launcher\AnvienLauncher.exe`;
- opening the root `Start-Anvien.html` file, which then calls `anvien://start` or `anvien://reset`.

These paths could show different UI because the HTML file was a separate entrypoint outside the Web app served by the launcher. The user clarified that the exe opened the inner Web UI directly, but the expected first screen was the start surface represented by the old HTML file.

The correct product direction is not to delete the start experience. The correct direction is to move that start experience into the Web UI served by the launcher exe, then remove the loose `Start-Anvien.html` file flow.

## Reopened Problem - Start Enters Manual Bridge Guide

User-provided screenshot `reports/problem/screenshot_1779366896.png` shows that clicking `Start Anvien` opens the `Start Anvien locally` / `anvien serve` guide.

That is wrong for the packaged exe flow:

- `AnvienLauncher.exe` is already responsible for starting the local backend/server bridge before opening the browser.
- The Start button must not ask the user to start `anvien serve` manually.
- After clicking Start in the packaged exe-served Web UI, the user should land in the repository chooser/analyze surface.
- If the packaged runtime is still connecting, the UI may show a packaged-runtime connecting state, but not terminal instructions.
- This must be validated in a browser that the user can actually observe, not only by headless Playwright.

Additional user clarification:

- The manual bridge guide screen itself is unnecessary and should be removed from active Web UI behavior.
- The product should not show a screen that teaches the user to run `anvien serve` manually after entering the Web UI.

## User Clarification

- `AnvienLauncher.exe` is a build artifact rebuilt frequently by `anvien-launcher\build.ps1`; it must not be treated as a fixed binary.
- The exe should be the single packaged user entrypoint.
- The initial user-facing surface should still be the start screen: `Anvien`, `Start Anvien`, `RESET RUNTIME`, `User Guide`, and status feedback.
- The Back button is a feature and must remain available from the graph shell.
- Back must return to the in-app start screen served by the launcher, not to `Start-Anvien.html`.

## Code Facts

Historical baseline from before implementation commit `212080d`:

- Root `Start-Anvien.html` contained the old start surface and called:
  - `window.location.href = 'anvien://start'` for Start;
  - `window.location.href = 'anvien://reset'` for Reset Runtime;
  - `fetch('user_guide.md')` for the User Guide panel.
- `anvien-launcher/src/main.go` served the packaged Web UI on `127.0.0.1:5228` and had special handling that could serve root `Start-Anvien.html` from the repo root.
- `anvien-web/src/App.tsx` started at the Web onboarding/repo landing flow rather than a launcher start surface.
- `anvien-web/src/components/Header.tsx` had a Back button that targeted `/Start-Anvien.html`.
- `anvien-web/vite.config.ts` had development-server knowledge of `/Start-Anvien.html`.
- `README.md`, `RUNBOOK.md`, and `TESTING.md` documented `Start-Anvien.html` as the user-facing launcher entry.
- The interrupted worktree edits from the previous wrong direction deleted the root HTML file and removed the Back button. Those edits were not accepted as a completed implementation of this plan.

Reopened pre-fix code facts after implementation commit `212080d`:

- Root `Start-Anvien.html` is deleted from tracked source and must stay removed.
- Launcher special serving for root `Start-Anvien.html` is removed.
- Vite dev-server special handling for root `Start-Anvien.html` is removed.
- Active docs now point users to `AnvienLauncher.exe`, not root `Start-Anvien.html`.
- `anvien-web/src/App.tsx` starts at the in-app `LauncherStartScreen`.
- The Header Back button remains and returns to the in-app start screen instead of `/Start-Anvien.html`.
- `Start Anvien` entered the existing `onboarding` view.
- `DropZone` rendered `OnboardingGuide` when the backend probe failed or was still unavailable.
- `OnboardingGuide` contained the manual `anvien serve` guidance. Per reopened clarification, it needed removal from active Web UI flow rather than preservation for dev mode.
- The launcher already injects lifecycle JavaScript into `index.html` when serving the packaged Web UI, so implementation can use existing launcher/runtime state without adding a loose HTML entrypoint.

Completed P6 code facts:

- `DropZone` no longer imports or renders `OnboardingGuide`.
- `OnboardingGuide.tsx` and its unit test are deleted.
- Backend-unavailable state renders the neutral `Connecting to Anvien runtime...` card.
- Start reaches repo landing when repos exist and Analyze Repository when no repos exist.
- The active Web UI no longer contains the manual `Start Anvien locally` / `anvien serve` guide.

Completed P8 code facts:

- The start screen User Guide button loads `/README.md`.
- `README.md` is served in Vite dev mode and copied into `anvien-web\dist\` during Web build.
- The packaged launcher serves `README.md` from `anvien-launcher\web-dist\`.

## Non-Goals

- Do not remove the start screen feature.
- Do not remove the Back button feature.
- Do not make users open a root HTML file.
- Do not keep `anvien://start` as the normal user start path.
- Do not introduce a new detached HTML launcher file.
- Do not change graph layout, node clustering, optimizer behavior, graph schema, repo indexing, or chat behavior in this plan.
- Do not solve entrypoint drift by adding timers, timeouts, delayed refresh, or retry loops.
- Do not show `anvien serve` manual bridge instructions after the user clicks Start.
- Do not preserve `OnboardingGuide` as an active fallback screen.
- Do not hide the real problem with a headless-only test that the user cannot observe when checking the PC manually.

## Target Design

### Packaged Entry

- The only packaged user entrypoint is `anvien-launcher\AnvienLauncher.exe`.
- The launcher still starts the backend and serves the Web UI from `anvien-launcher\web-dist\`.
- The launcher opens the browser to the exe-served Web UI.
- The first Web UI screen shown by that path is an in-app start screen, not the graph shell directly and not a root HTML file.

### In-App Start Screen

- Add a Web UI start surface that preserves the old start screen product role:
  - title: `Anvien`;
  - primary action: `Start Anvien`;
  - secondary action: `RESET RUNTIME`;
  - secondary action: `User Guide`;
  - status text for user feedback.
- `Start Anvien` transitions inside the Web app to the existing repo landing/analyze flow. It must not call `anvien://start`.
- `User Guide` remains available from this start surface and displays the repository `README.md`. If `README.md` is absent from a build, the UI must fail gracefully without a broken dead end.
- `RESET RUNTIME` must continue to reset the packaged runtime through the launcher exe path and must remain hidden/no terminal flash. If the implementation uses `anvien://reset`, it must be treated as an internal exe reset action, not a separate start entrypoint.

### Packaged Start Flow

- Clicking `Start Anvien` must bypass the manual bridge guide and go directly to the runtime/repo landing/analyze path.
- If the backend is not immediately reachable, show a neutral runtime connecting/recovery state driven by existing runtime state/probes; do not add timeout, retry-loop, delayed navigation, or elapsed-time budget logic.
- The runtime connecting/recovery state must not tell the user to copy or run `anvien serve`.
- The old manual bridge guide should not remain as an active dev fallback. If the backend is unavailable, use a neutral runtime connection state or a product-level recovery action, not terminal command instructions.

### Back Navigation

- Keep the top-bar Back button in the graph shell.
- Rename its behavior away from file navigation: it returns to the in-app start screen.
- Back must not navigate to `/Start-Anvien.html`.
- Back must not show a false reconnect banner while switching screens inside the SPA.

### Removed File Flow

- Delete the tracked root `Start-Anvien.html` file after the in-app start screen exists.
- Remove launcher special-serving of root `Start-Anvien.html`.
- Remove Vite dev-server behavior that treats `Start-Anvien.html` as a launcher asset.
- Update active docs so users are told to run `AnvienLauncher.exe`, not open `Start-Anvien.html`.
- Stale `/Start-Anvien.html` requests should not serve the old file. Prefer an explicit 404/410 or a documented SPA fallback that still lands on the in-app start screen, but do not keep a physical root HTML artifact.

## Implementation Slices

### P0 - Reconcile Interrupted Worktree

- [x] Inspect the current dirty worktree and separate intentional plan docs from interrupted code edits.
- [x] Do not accept the previous wrong direction as done: the Back button must be restored as an in-app navigation feature.
- [x] Ensure future commits stage only the intended slice files.

### P1 - Add In-App Start Screen

- [x] Add a Web UI start screen component that preserves the old HTML surface's product actions and status feedback.
- [x] Add a new app state/view mode or equivalent state transition so the start screen is the initial packaged Web UI surface.
- [x] Wire `Start Anvien` to enter the existing repo landing/analyze flow without using `Start-Anvien.html` or `anvien://start`.
- [x] Preserve User Guide access or provide a graceful unavailable state.

### P2 - Preserve Back Feature

- [x] Keep the Header Back button visible in the graph shell.
- [x] Change the Back action to return to the in-app start screen.
- [x] Remove `/Start-Anvien.html` URL generation from Header.
- [x] Update unit and e2e tests to prove Back returns to the in-app start screen.

### P3 - Remove Loose HTML Entrypoint

- [x] Delete root `Start-Anvien.html` after P1/P2 are implemented.
- [x] Remove launcher static-handler special casing for root `Start-Anvien.html`.
- [x] Remove Vite dev-server special handling for root `Start-Anvien.html`.
- [x] Update launcher tests so stale `/Start-Anvien.html` requests do not serve a root file.

### P4 - Update Active Docs

- [x] Update `README.md` packaged launcher docs to make `AnvienLauncher.exe` the entrypoint.
- [x] Update `RUNBOOK.md` packaged launcher run/reset instructions.
- [x] Update `TESTING.md` user acceptance steps.
- [x] Update `CHANGELOG.md` with the exe-served start screen migration.

### P5 - Validation

- [x] Run full build first: `powershell -ExecutionPolicy Bypass -File anvien-launcher\build.ps1`.
- [x] Verify the build regenerates `anvien-launcher\AnvienLauncher.exe` and packaged `web-dist`.
- [x] Run focused launcher Go tests.
- [x] Run focused Web unit tests for start screen and Header Back.
- [x] Run full Web tests.
- [x] Run Web e2e tests covering initial start screen, Start action, Back action, and absence of root HTML navigation.
- [x] Run broader Go tests required by the touched packages.
- [x] Verify root `Start-Anvien.html` is absent and not copied into `anvien-web\dist\` or `anvien-launcher\web-dist\`.
- [x] Verify `rg -n "Start-Anvien.html|anvien://start"` has no active product/docs/test references except historical plan ledgers or explicit stale-path tests.
- [x] Run required change detection before commit according to active repo instructions.
- [x] Commit the completed implementation slice.

### P6 - Fix Packaged Start Runtime Path

- [x] Remove `OnboardingGuide` from active Web UI routing.
- [x] Change `Start Anvien` so it enters the runtime/repo chooser flow, not the manual `OnboardingGuide`.
- [x] Ensure runtime connection flow lands on repo landing when repos exist.
- [x] Ensure runtime connection flow lands on Analyze Repository when no repos exist.
- [x] Replace backend-unavailable behavior with a neutral runtime connecting/recovery state that does not mention `anvien serve`.
- [x] Delete or retire tests that assert `Start Anvien locally` is shown.
- [x] Add unit tests that prove Start never renders `Start Anvien locally` or `anvien serve`.
- [x] Add e2e coverage that clicks Start and verifies repo chooser/analyze surface, not the manual bridge guide.

### P7 - User-Observable Browser Validation

- [x] Validate with a visible browser on the user's PC, not only headless Playwright.
- [x] Prefer a headed Playwright run against the built packaged launcher runtime so the browser window is visible on the desktop.
- [x] Record the exact command/tool used, whether a visible browser was opened, and the observed screen after clicking Start.
- [x] Capture or reference a screenshot proving the post-Start screen is repo chooser/analyze, not the manual bridge guide.

### P8 - User Guide Reads README

- [x] Change the start screen User Guide loader from `user_guide.md` to root `README.md`.
- [x] Ensure Web build/dev serving exposes `README.md` at `/README.md`.
- [x] Ensure packaged launcher `web-dist` contains `README.md`.
- [x] Add unit coverage proving User Guide fetches and displays `README.md`.
- [x] Add e2e coverage proving User Guide displays README content from the served Web UI.
- [x] Validate the same behavior through `AnvienLauncher.exe`.

## Completion Criteria

- Running the rebuilt `AnvienLauncher.exe` opens the exe-served Web UI at the in-app start screen.
- The start screen has Start, Reset Runtime, User Guide, and status feedback.
- Start enters the existing repo landing/analyze flow inside the app.
- The graph shell Back button returns to the in-app start screen.
- Users are no longer instructed to open `Start-Anvien.html`.
- Root `Start-Anvien.html` is removed from tracked source and packaged artifacts.
- There is no separate root HTML launcher flow left in active code.
- Full build, focused tests, full tests, and Web e2e validation are recorded in evidence.
- Clicking `Start Anvien` from the packaged exe-served start screen never opens `Start Anvien locally` / `anvien serve`.
- `Start Anvien locally` / manual `anvien serve` guide is removed from active Web UI behavior.
- Browser validation includes a user-observable run, not only headless automation.
- The User Guide button displays repository `README.md` content from the exe-served Web UI.
