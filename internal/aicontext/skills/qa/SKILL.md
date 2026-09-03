---
name: qa
description: Use when the user asks to run QA without fixing code, including mounted runtime behavior, visible user flows, browser-visible app execution, source-of-truth checks, action/state coverage, route/control inventories, Playwright control sweeps, or QA report generation in repositories.
---
# (MUST) For Codex: When QA must use plugins
- Browser - Control the in-app browser
- Chrome - Control the user's real Chrome browser
- Computer Use - Control Windows apps or installed artifacts when QA requires real app interaction outside a browser.
- Playwright: use as an automation arm for browser actions, control sweeps, screenshots, videos, traces, and reports.
- Backend/API tools: use HTTP clients (curl/fetch), database CLI tools, or test runners for non-UI scopes.

# (MUST) For Claude or other agents:
- Use the equivalent browser, Chrome/session, computer-control, or HTTP/database CLI capability exposed by that agent/runtime.

# QA Runtime Review

Use this skill to verify real runtime behavior from the user's point of view. QA is a no-fix role: inspect, execute, record, classify, and report. Do not repair product code, tests, specs, fixtures, or generated files during QA.

## Iron Laws

- QA MUST NOT fix code. Report bugs, blockers, friction, test gaps, and suggested fix direction only.
- Runtime behavior is the source of QA truth. Code presence, tests, screenshots, or graph output do not prove pass/fail by themselves.
- Must have screenshot Evidence for UI scopes. For non-UI / Backend / Database scopes, follow `references/backend-qa-protocol.md` (structured payload, DB state diff, and execution logs).
- Tests verify already-correct behavior. Tests do not define expected behavior and must not legitimize wrong UX.
- Wrong tests are findings. A green test that encodes wrong behavior is not QA proof.
- Do not skip steps. Every in-scope flow step matters equally.
- Every in-scope surface, action, state, data path, and blocker needs an explicit verdict or out-of-scope mark.
- An incomplete inventory or action ledger means QA is incomplete.
- If a blocker prevents downstream verification, downstream scope is `Blocked`, not passed.
- For final visible QA, open the real built website, runtime, or generated app artifact the same way a user would open it, in a visible browser or app window on the user's physical PC. Headless, hidden, source-only, test-only, or screenshot-only execution cannot approve final visible QA.
- QA must run on the real, officially built application (actual Electron artifact or release binary).
- For UI/visible scopes: Every action (clicking buttons, entering data, changing state) must capture clear PNG screenshots saved sequentially into the `screenshots/` directory.
- For UI/visible scopes: Visually inspect each captured image / visual artifact to detect UI and behavioral defects; never rely solely on in-code assertions or generated `.json`/`.md` summary files.
- Never accept a UI QA report lacking the actual screenshot evidence set for each interaction. For Backend / non-UI reports, require structured HTTP status, response payload, and DB state diff evidence per `references/backend-qa-protocol.md`.

## Source Of Expected Behavior

Use the first available applicable authority, in this order:

1. User's explicit QA scope and no-fix instruction.
2. Repository rules such as `AGENTS.md`.
3. Active execution plan, QA plan, or SPEC family.
4. Visible-browser plan when final visible QA is requested or the plan is in scope.
5. This skill's QA protocol.
6. Coverage model, inventory requirements, and control/state protocols below.
7. Automated tests as verification evidence only.

If authorities conflict, state the conflict and follow the higher authority. Never use a test expectation to override real required product behavior.

Record evidence and report findings, including:
- route/page candidates
- mounted entry candidates
- action/handler/API candidates
- data source/store/API/DB candidates
- permission/context gate candidates
- likely source files for suggested fix direction

Pass/fail requires runtime evidence.

## Core Workflow (Frontend / UI)

1. **Preflight** → build, runtime, browser (`references/frontend-runtime-and-build-rules.md`)
2. **Inventory** → surfaces, controls, states (`references/frontend-coverage-model.md`)
3. **Automation** → screenshots, control sweep (`references/frontend-automation-and-control-sweep.md`)
4. **Data flow** → field/button tracing, DOM state (`references/frontend-data-flow.md`)
5. **Source-of-truth** → verify against real DB state (`references/source-of-truth-rules.md`)
6. **Locale** → if i18n in scope (`references/frontend-i18n-and-locale.md`)
7. **Evidence & Report** → screenshots, visual blockers, report, handoff (`references/reporting.md`)

## Core Workflow (Backend / API & DB)

1. **Preflight & Protocol** → verify server runtime, DB connection, closed-loop API execution (`references/backend-qa-protocol.md`)
2. **Source-of-truth** → verify real DB state mutations, no mock data (`references/source-of-truth-rules.md`)
3. **Evidence & Blockers** → verify test bounds, blocker propagation (`references/evidence-and-blockers.md`)
4. **Report & Escalation** → HTTP status, DB diff, latency/RAM, handoff; escalate to `Edge-Case` if hostile chaos is required (`references/reporting.md`)

## Red Flags

Stop, report, or re-scope if:

- QA starts fixing code.
- QA runs only headless/hidden browser for final visible QA.
- QA skips inventory.
- QA skips a step in an in-scope flow.
- QA clicks only the happy path.
- QA ignores disabled, no-op, rejected, empty, error, stale, or blocked states.
- QA uses fake app/commercial data as if it were real production state.
- QA ignores DB/source readback for DB-backed behavior.
- QA ignores locale/session/permission changes when in scope.
- QA reports downstream scope as passed after an upstream blocker.
