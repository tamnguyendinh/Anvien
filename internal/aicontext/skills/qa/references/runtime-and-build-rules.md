# Runtime, Build, And Visible Browser Rules

> This file is part of QA Skill v2. Read when: preparing build, starting runtime, opening browser/app for QA.

## Preflight Gates

Before runtime QA:

1. Confirm scope, scope mode, required inventories, runtime target, and no-fix boundary.
2. Read active repo rules and the expected-behavior docs for the declared scope.
3. Run the full build required by repo or plan before QA testing.
4. Start the real runtime or production-like runtime.
5. If production-like runtime is the target, rebuild Docker or equivalent runtime before checking the user-facing URL/UI.
6. Record build command, runtime command, URL, health state, environment, seed/fixture state, and browser state.
7. If build/runtime/browser cannot be prepared, stop or mark scope blocked.

Production rule:
- Do not approve production behavior from a dev server, test-only route, stale container, or unrebuilt artifact.
- Verify on the URL and UI that users actually use.

## Visible Runtime And Automation Rules

Final visible QA must start from the real user entry point, not from source code, component previews, test harnesses, or direct internal URLs unless the scope explicitly defines that URL as the user entry point.

For browser-based website QA:

1. Run the repo-required full build first, then build the Docker/VPS runtime or repo-defined production runtime required for the website.
2. Start that freshly built runtime exactly as the repo or deployment plan defines it.
3. Open a real visible Chrome or Edge window on the user's physical PC.
4. Click the browser address bar.
5. Type or paste the target user URL into the address bar.
6. Press Enter and wait for the page to load.
7. Perform QA by interacting with the visible page: click controls, type into fields, submit forms, navigate links, reload, go back/forward, and observe UI state changes.
8. Use Browser, Chrome, Playwright, or equivalent automation only to drive and record those real visible-browser interactions; automation must not replace opening the real built runtime in a visible browser.

For app or desktop artifact QA:

1. Run the full build, build the distributable artifact, and build any Docker/container/VPS runtime required by the app first.
2. If the app uses Docker/container/VPS runtime, start that freshly built runtime exactly as the repo or deployment plan defines it.
3. Click or open that post-full-build artifact runtime on this visible PC the same way a user would open it: installer, executable, packaged desktop app, or documented launch entry.
4. Use the visible app window on the user's physical PC as the QA surface.
5. Use Computer Use or equivalent desktop-control capability to click, type, navigate, and capture evidence in that real app window.
6. Do not approve app QA from source code, a dev runner, a component harness, or an unbuilt workspace state.

For Codex runtime control, choose the capability by target:

- Chrome: use for final visible website QA, existing login/session/cookies, or user-observable browser state.
- Browser: use for in-app browser inspection or local runtime diagnostics when a user-visible real browser is not required by the QA scope.
- Computer Use: use for desktop apps, installers, generated artifacts, Windows UI, or flows outside a browser.
- Playwright: use as browser automation for actions, control sweeps, screenshots, videos, traces, and reports against the real visible browser/runtime.

For Claude or other agents:

- Use the equivalent visible browser, session-aware browser, desktop-control, or Playwright-like capability available in that environment.
- If no equivalent visible runtime-control capability is available, mark visible runtime QA as blocked instead of approving from source code, tests, screenshots, or headless-only checks.

## Websites

Website QA must run after the full build and against the latest freshly built Docker/VPS runtime or repo-defined production runtime.

- Build and run the real runtime lifecycle target: Docker/container/VPS runtime when the repo or deployment uses it, otherwise the repo-defined production server or served static artifact.
- Do not use a dev server such as Vite dev, Next dev, `npm run dev`, or any equivalent development runtime as final QA evidence.
- Do not use dev test, component preview, fixture-only routes, or Playwright test-server output as final QA evidence.
- Do not click or inspect built files directly as website QA. The built website must be served by the expected runtime and opened through a real browser URL.
- The QA flow must start by opening a visible browser on the user's PC, entering the target URL in the address bar, loading the page, and then interacting with the loaded page.
- Browser automation may drive clicks, typing, submits, navigation, reloads, screenshots, traces, and reports only after the real visible browser has loaded the real target URL.
- All user actions must occur in the real visible browser session.

## Apps

App QA must run against the post-full-build built/generated artifact clicked or opened on this visible PC the same way a user would open it.

- Before QA, run the full build, build the app artifact, build any required server/runtime, and build a fresh Docker/container/VPS runtime when the app uses one.
- If the app uses Docker/container/VPS runtime, run the freshly built runtime during QA; do not QA against a stale container or previously running runtime.
- If the app ships as an installer, executable, desktop bundle, or packaged artifact, QA that post-full-build built artifact by clicking or opening it on this visible PC instead of source code, a dev runner, or dev test output.
- Use Computer Use or an equivalent desktop-control capability when the app is outside the browser.
- Use Chrome, Browser, or Playwright only for browser-based app surfaces.
- Automation is only the interaction and evidence layer; the built artifact/runtime remains the QA source of truth.
