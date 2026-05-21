# AVmatrix Exe-Served Start Screen Plan

Date: 2026-05-21

Status: completed

Companion files:

- Benchmark ledger: [2026-05-21-avmatrix-exe-served-start-screen-benchmark.md](2026-05-21-avmatrix-exe-served-start-screen-benchmark.md)
- Evidence ledger: [2026-05-21-avmatrix-exe-served-start-screen-evidence.md](2026-05-21-avmatrix-exe-served-start-screen-evidence.md)

## Rules

1. Follow active workspace and repository instructions, including `AGENTS.md`, for implementation workflow. This plan records product work and validation; it does not replace those rules.
2. Use AVmatrix according to the active repo instructions for implementation slices; do not use AVmatrix for doc-only commits.
3. Update this plan, the benchmark ledger, and the evidence ledger as each implementation slice is completed.
4. Run the full packaged build before tests. For this plan the full build gate is `powershell -ExecutionPolicy Bypass -File avmatrix-launcher\build.ps1`.
5. Because this changes Web UI behavior, validation must include focused unit tests and Web e2e tests after the full build passes.
6. Do not add product/runtime timeout, delayed navigation, or elapsed-time budget as the fix mechanism.
7. Keep the Back button as a product feature.
8. Remove the separate `Start-AVmatrix.html` file flow; do not replace it with another loose root HTML entrypoint.

## Problem

The packaged launcher currently has two user-facing start paths:

- double-clicking or running `avmatrix-launcher\AVmatrixLauncher.exe`;
- opening the root `Start-AVmatrix.html` file, which then calls `avmatrix://start` or `avmatrix://reset`.

These paths can show different UI because the HTML file is a separate entrypoint outside the Web app served by the launcher. The user clarified that the exe currently opens the inner Web UI directly, but the expected first screen is the start surface represented by the old HTML file.

The correct product direction is not to delete the start experience. The correct direction is to move that start experience into the Web UI served by the launcher exe, then remove the loose `Start-AVmatrix.html` file flow.

## User Clarification

- `AVmatrixLauncher.exe` is a build artifact rebuilt frequently by `avmatrix-launcher\build.ps1`; it must not be treated as a fixed binary.
- The exe should be the single packaged user entrypoint.
- The initial user-facing surface should still be the start screen: `AVmatrix`, `Start AVmatrix`, `RESET RUNTIME`, `User Guide`, and status feedback.
- The Back button is a feature and must remain available from the graph shell.
- Back must return to the in-app start screen served by the launcher, not to `Start-AVmatrix.html`.

## Current Code Facts

- Root `Start-AVmatrix.html` contains the current start surface and calls:
  - `window.location.href = 'avmatrix://start'` for Start;
  - `window.location.href = 'avmatrix://reset'` for Reset Runtime;
  - `fetch('user_guide.md')` for the User Guide panel.
- `avmatrix-launcher/src/main.go` serves the packaged Web UI on `127.0.0.1:5228`.
- The launcher currently has special handling that can serve root `Start-AVmatrix.html` from the repo root.
- `avmatrix-web/src/App.tsx` currently starts at the Web onboarding/repo landing flow rather than a launcher start surface.
- `avmatrix-web/src/components/Header.tsx` currently has a Back button that targets `/Start-AVmatrix.html`.
- `avmatrix-web/vite.config.ts` has development-server knowledge of `/Start-AVmatrix.html`.
- `README.md`, `RUNBOOK.md`, and `TESTING.md` still document `Start-AVmatrix.html` as the user-facing launcher entry.
- Current interrupted worktree edits from the previous wrong direction deleted the root HTML file and removed the Back button. Those edits must not be considered a completed implementation of this plan.

## Non-Goals

- Do not remove the start screen feature.
- Do not remove the Back button feature.
- Do not make users open a root HTML file.
- Do not keep `avmatrix://start` as the normal user start path.
- Do not introduce a new detached HTML launcher file.
- Do not change graph layout, node clustering, optimizer behavior, graph schema, repo indexing, or chat behavior in this plan.
- Do not solve entrypoint drift by adding timers, timeouts, delayed refresh, or retry loops.

## Target Design

### Packaged Entry

- The only packaged user entrypoint is `avmatrix-launcher\AVmatrixLauncher.exe`.
- The launcher still starts the backend and serves the Web UI from `avmatrix-launcher\web-dist\`.
- The launcher opens the browser to the exe-served Web UI.
- The first Web UI screen shown by that path is an in-app start screen, not the graph shell directly and not a root HTML file.

### In-App Start Screen

- Add a Web UI start surface that preserves the old start screen product role:
  - title: `AVmatrix`;
  - primary action: `Start AVmatrix`;
  - secondary action: `RESET RUNTIME`;
  - secondary action: `User Guide`;
  - status text for user feedback.
- `Start AVmatrix` transitions inside the Web app to the existing repo landing/analyze flow. It must not call `avmatrix://start`.
- `User Guide` remains available from this start surface. If `user_guide.md` is still absent, the UI must fail gracefully without a broken dead end.
- `RESET RUNTIME` must continue to reset the packaged runtime through the launcher exe path and must remain hidden/no terminal flash. If the implementation uses `avmatrix://reset`, it must be treated as an internal exe reset action, not a separate start entrypoint.

### Back Navigation

- Keep the top-bar Back button in the graph shell.
- Rename its behavior away from file navigation: it returns to the in-app start screen.
- Back must not navigate to `/Start-AVmatrix.html`.
- Back must not show a false reconnect banner while switching screens inside the SPA.

### Removed File Flow

- Delete the tracked root `Start-AVmatrix.html` file after the in-app start screen exists.
- Remove launcher special-serving of root `Start-AVmatrix.html`.
- Remove Vite dev-server behavior that treats `Start-AVmatrix.html` as a launcher asset.
- Update active docs so users are told to run `AVmatrixLauncher.exe`, not open `Start-AVmatrix.html`.
- Stale `/Start-AVmatrix.html` requests should not serve the old file. Prefer an explicit 404/410 or a documented SPA fallback that still lands on the in-app start screen, but do not keep a physical root HTML artifact.

## Implementation Slices

### P0 - Reconcile Interrupted Worktree

- [x] Inspect the current dirty worktree and separate intentional plan docs from interrupted code edits.
- [x] Do not accept the previous wrong direction as done: the Back button must be restored as an in-app navigation feature.
- [x] Ensure future commits stage only the intended slice files.

### P1 - Add In-App Start Screen

- [x] Add a Web UI start screen component that preserves the old HTML surface's product actions and status feedback.
- [x] Add a new app state/view mode or equivalent state transition so the start screen is the initial packaged Web UI surface.
- [x] Wire `Start AVmatrix` to enter the existing repo landing/analyze flow without using `Start-AVmatrix.html` or `avmatrix://start`.
- [x] Preserve User Guide access or provide a graceful unavailable state.

### P2 - Preserve Back Feature

- [x] Keep the Header Back button visible in the graph shell.
- [x] Change the Back action to return to the in-app start screen.
- [x] Remove `/Start-AVmatrix.html` URL generation from Header.
- [x] Update unit and e2e tests to prove Back returns to the in-app start screen.

### P3 - Remove Loose HTML Entrypoint

- [x] Delete root `Start-AVmatrix.html` after P1/P2 are implemented.
- [x] Remove launcher static-handler special casing for root `Start-AVmatrix.html`.
- [x] Remove Vite dev-server special handling for root `Start-AVmatrix.html`.
- [x] Update launcher tests so stale `/Start-AVmatrix.html` requests do not serve a root file.

### P4 - Update Active Docs

- [x] Update `README.md` packaged launcher docs to make `AVmatrixLauncher.exe` the entrypoint.
- [x] Update `RUNBOOK.md` packaged launcher run/reset instructions.
- [x] Update `TESTING.md` user acceptance steps.
- [x] Update `CHANGELOG.md` with the exe-served start screen migration.

### P5 - Validation

- [x] Run full build first: `powershell -ExecutionPolicy Bypass -File avmatrix-launcher\build.ps1`.
- [x] Verify the build regenerates `avmatrix-launcher\AVmatrixLauncher.exe` and packaged `web-dist`.
- [x] Run focused launcher Go tests.
- [x] Run focused Web unit tests for start screen and Header Back.
- [x] Run full Web tests.
- [x] Run Web e2e tests covering initial start screen, Start action, Back action, and absence of root HTML navigation.
- [x] Run broader Go tests required by the touched packages.
- [x] Verify root `Start-AVmatrix.html` is absent and not copied into `avmatrix-web\dist\` or `avmatrix-launcher\web-dist\`.
- [x] Verify `rg -n "Start-AVmatrix.html|avmatrix://start"` has no active product/docs/test references except historical plan ledgers or explicit stale-path tests.
- [x] Run required change detection before commit according to active repo instructions.
- [x] Commit the completed implementation slice.

## Completion Criteria

- Running the rebuilt `AVmatrixLauncher.exe` opens the exe-served Web UI at the in-app start screen.
- The start screen has Start, Reset Runtime, User Guide, and status feedback.
- Start enters the existing repo landing/analyze flow inside the app.
- The graph shell Back button returns to the in-app start screen.
- Users are no longer instructed to open `Start-AVmatrix.html`.
- Root `Start-AVmatrix.html` is removed from tracked source and packaged artifacts.
- There is no separate root HTML launcher flow left in active code.
- Full build, focused tests, full tests, and Web e2e validation are recorded in evidence.
