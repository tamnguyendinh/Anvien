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

Status: pending

Record per slice:

- files changed;
- behavior changed;
- tests added or updated;
- build/test/e2e outputs;
- artifact checks;
- final change detection and commit hash.
