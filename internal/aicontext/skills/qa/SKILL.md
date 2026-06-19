---
name: qa
description: Use when the user asks to run QA without fixing code, including mounted runtime behavior, visible user flows, browser-visible app execution, source-of-truth checks, action/state coverage, route/control inventories, Playwright control sweeps, or QA report generation in repositories where Anvien can support
---
# (MUST) For Codex: When QA must use plugins
- Browser - Control the in-app browser
- Chrome - Control the user's real Chrome browser
- Computer Use - Control Windows apps or installed artifacts when QA requires real app interaction outside a browser.
- Playwright: use as an automation arm for browser actions, control sweeps, screenshots, videos, traces, and reports.

# (MUST) For Claude or other agents:
- Use the equivalent browser, Chrome/session, or computer-control capability exposed by that agent/runtime or Playwright-like capability available in that environment.

# QA Runtime Review With Anvien

Use this skill to verify real runtime behavior from the user's point of view. QA is a no-fix role: inspect, execute, record, classify, and report. Do not repair product code, tests, specs, fixtures, or generated files during QA.

Anvien is a discovery tool for QA. It helps map routes, mounted surfaces, handlers, APIs, stores, permissions, source files, and data flow. It does not prove that a user can reach a surface, click a control, see the right state, or complete a flow. Runtime evidence is the QA verdict source.

## Iron Laws

- QA MUST NOT fix code. Report bugs, blockers, friction, test gaps, and suggested fix direction only.
- Runtime behavior is the source of QA truth. Code presence, tests, screenshots, or graph output do not prove pass/fail by themselves.
- Tests verify already-correct behavior. Tests do not define expected behavior and must not legitimize wrong UX.
- Wrong tests are findings. A green test that encodes wrong behavior is not QA proof.
- Do not skip steps. Every in-scope flow step matters equally.
- Every in-scope surface, action, state, data path, and blocker needs an explicit verdict or out-of-scope mark.
- An incomplete inventory or action ledger means QA is incomplete.
- If a blocker prevents downstream verification, downstream scope is `Blocked`, not passed.
- For final visible QA, open the real built website, runtime, or generated app artifact the same way a user would open it, in a visible browser or app window on the user's physical PC. Headless, hidden, source-only, test-only, or screenshot-only execution cannot approve final visible QA.

## Source Of Expected Behavior

Use the first available applicable authority, in this order:

1. User's explicit QA scope and no-fix instruction.
2. Repository rules such as `AGENTS.md`.
3. Active execution plan, QA plan, or SPEC family.
4. Visible-browser plan when final visible QA is requested or the plan is in scope.
5. This skill's QA protocol.
6. Coverage model, inventory requirements, and control/state protocols below.
7. Anvien graph evidence as mapping support only.
8. Automated tests as verification evidence only.

If authorities conflict, state the conflict and follow the higher authority. Never use a test expectation to override real required product behavior.

## Anvien Use In QA

Use Anvien to avoid missing scope and to understand data flow. Do not use Anvien as a substitute for runtime interaction.

Before graph-based Anvien work, run:

```text
anvien analyze --force
```

Useful Anvien choices:

| QA Need | Use |
| --- | --- |
| Find route, page, component, behavior, or risk owner | `anvien query "<concept>" --repo <repo>` |
| Find candidate files for a behavior | `anvien query files "<concept>" --repo <repo>` |
| Inspect one file deeply | `anvien file-detail <path> --repo <repo> --json` |
| Inspect route handlers and consumers | `anvien api route-map [route] --repo <repo>` |
| Inspect response shape drift against consumers | `anvien api shape-check [route] --repo <repo>` |
| Inspect tool/IPC paths | `anvien api tool-map [tool] --repo <repo>` |
| Identify source ownership for suggested fix direction | `anvien impact file <path> --repo <repo> --direction upstream` or `anvien impact symbol "<symbol>" --repo <repo> --direction upstream` |
| Audit graph evidence if Anvien looks stale or incomplete | `anvien graph-health summary --repo <repo> --json` |

Record Anvien evidence as scope-mapping evidence:
- route/page candidates
- mounted entry candidates
- action/handler/API candidates
- data source/store/API/DB candidates
- permission/context gate candidates
- likely source files for suggested fix direction

Do not mark anything passed because Anvien found it. Pass/fail requires runtime evidence.

## Runtime, Build, And Visible Browser Rules

### Preflight Gates

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

### Visible Runtime And Automation Rules

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

### Websites

Website QA must run after the full build and against the latest freshly built Docker/VPS runtime or repo-defined production runtime.

- Build and run the real runtime lifecycle target: Docker/container/VPS runtime when the repo or deployment uses it, otherwise the repo-defined production server or served static artifact.
- Do not use a dev server such as Vite dev, Next dev, `npm run dev`, or any equivalent development runtime as final QA evidence.
- Do not use dev test, component preview, fixture-only routes, or Playwright test-server output as final QA evidence.
- Do not click or inspect built files directly as website QA. The built website must be served by the expected runtime and opened through a real browser URL.
- The QA flow must start by opening a visible browser on the user's PC, entering the target URL in the address bar, loading the page, and then interacting with the loaded page.
- Browser automation may drive clicks, typing, submits, navigation, reloads, screenshots, traces, and reports only after the real visible browser has loaded the real target URL.
- All user actions must occur in the real visible browser session.

### Apps

App QA must run against the post-full-build built/generated artifact clicked or opened on this visible PC the same way a user would open it.

- Before QA, run the full build, build the app artifact, build any required server/runtime, and build a fresh Docker/container/VPS runtime when the app uses one.
- If the app uses Docker/container/VPS runtime, run the freshly built runtime during QA; do not QA against a stale container or previously running runtime.
- If the app ships as an installer, executable, desktop bundle, or packaged artifact, QA that post-full-build built artifact by clicking or opening it on this visible PC instead of source code, a dev runner, or dev test output.
- Use Computer Use or an equivalent desktop-control capability when the app is outside the browser.
- Use Chrome, Browser, or Playwright only for browser-based app surfaces.
- Automation is only the interaction and evidence layer; the built artifact/runtime remains the QA source of truth.

### Screenshot Evidence Rule

- When browser or desktop automation is used for QA, screenshots must be captured.
- Do not use the final screenshot after failure as the main evidence.
- Capture screenshots at each small action step: before entering data, after each field, before clicking, after clicking, and after the UI responds or settles.
- After the run finishes, open and visually inspect the screenshots to determine exactly which step first introduced the issue.
- Bugs are not necessarily blockers. If a bug is found, report it in the report/evidence section, but continue testing if a valid path remains.
- Review the whole screen for additional issues, including controls that do not respond, inputs that reject data, incorrect results, overlapping text/cards, broken fonts, overflow, disappearing elements, or layout shifts.

### Automated Control Sweep

When browser or desktop automation is used for QA, build a per-page/per-tab/per-locale control inventory before verdict.

For every in-scope page, tab, dialog, drawer, dropdown, menu, form, table row action, and navigation surface:

- Inventory every reachable user control by role/name/locator, visible state, enabled/disabled state, locale, route, tab, persona/context, and expected outcome.
- Exercise every reachable enabled control through real user interaction in the real visible browser or app window: click, type, select, submit, keyboard navigation, close, back, refresh, redirect, retry, and reopen paths when applicable.
- Do not force-click hidden or disabled controls as pass/fail proof.
- For hidden or disabled controls, verify they are correctly unreachable or disabled, expose the expected reason/state, and do not trigger forbidden behavior.
- For dropdowns, menus, listboxes, comboboxes, and selects, open the control, verify options, select applicable options, test close/escape/outside-click behavior, and verify resulting state/navigation.
- For forms, cover pristine, dirty, invalid, valid, submitting, success, error, cancel, reopen, and validation-message states.
- For navigation controls, cover links, tabs, sub-tabs, deep links, redirects, browser back/forward, reload, selected nav state, and destination correctness.
- For i18n scope, repeat the inventory and action ledger for every supported locale in scope.
- Record the sweep in the Action Ledger; a control/state combination without a ledger row is not covered.
- If the full cross-product cannot be completed, mark missing combinations explicitly as `Blocked`, `Out of scope`, or `Unverified`.

### Inventory Before Verdict

Before claiming coverage, build an inventory for the declared scope.

Inventory types:
- visible-surface inventory
- interactive inventory
- route/control inventory
- state inventory
- action ledger template
- data/source-of-truth map
- persona/context matrix

Each inventory item needs:
- surface/action/state name
- route or flow
- runtime entry point
- user trigger
- persona/context
- expected result
- data source when relevant
- evidence slot
- verdict slot

Verdicts:
- `Pass`
- `Fail`
- `Blocked`
- `N/A`
- `Out of scope`

Rules:
- `100% coverage` means `100% of declared scope`, not the whole app unless the declared scope is the whole app.
- A route, dialog, component, or action that exists in code but cannot be reached through the real mounted runtime path is not covered.
- A screen is not covered if child dialogs, menus, row actions, forms, or state variants were never inventoried.
- If the ledger is incomplete, QA is incomplete.

### No-Step-Skipped Flow Rule

In each in-scope flow, account for every step:

- open page or app
- initial state
- every tab
- every link
- every button
- every field
- every form
- every submit
- every redirect
- every refresh, back action, or route change
- every locale switch
- every cookie or session change
- every DB or read-model read/write when relevant
- UI after each step
- where the next step leads

If a step exists in scope but was not observed or ledgered, that flow is not fully covered.

### State And Action Matrix

For each inventoried surface, test every applicable state:

- visible / hidden
- enabled / disabled
- loading
- empty
- error
- success
- blocked
- warning
- stale
- selected / unselected
- expanded / collapsed
- pristine / dirty
- validation-error
- submitting
- permission-gated
- shift-gated
- session-gated
- subscription-gated
- stale-context after owner, app scope, role, shift, session, subscription, or locale changes

For every visible action, verify:

- the user can actually reach it
- the action triggers the correct mounted behavior
- the action targets the correct row, item, entity, or selection
- blocked actions explain why they are blocked
- loading, empty, error, success, disabled, and blocked states explain themselves clearly
- displayed data/state updates correctly after the action
- reopen, retry, back, refresh, resubmit, and context switch do not leave stale UI behind

Exploratory clicks are additive. They never replace the required state/action matrix.

## Runtime Map, Data Flow, And Source-Of-Truth

### Mounted Runtime Map

For each inventoried item, record the runtime map:

- runtime entry point
- user trigger
- owner, app, tenant, or active scope
- persona and role
- shift state
- session/subscription state
- locale and viewport when relevant
- dataset or fixture
- data source, store, API, DB, or read-model
- expected visible actions
- expected blocked actions
- expected next destination

### Field And Button Data Flow

For each relevant field, button, submit, mutation, money-sensitive action, source-of-truth value, or DB-backed result, understand and record when applicable:

- where the data comes from
- which field receives input
- which endpoint, action, tool, or submit path receives it
- what frontend validation runs
- what backend/server validation runs
- which table, read-model, store, cache, or state is read
- which table, read-model, store, cache, or state is written
- where the app redirects or renders afterward
- what the next UI allows the user to do
- how public, account, admin, or user surfaces reflect the change

Use Anvien to map this data flow when useful. Verify it with runtime UI and source-of-truth evidence.

### Source-Of-Truth Data

Treat runtime data as production state, not demo decoration.

Rules:
- Public app data appears only when created through the correct upload/publish/data-entry/admin flow and exists as real runtime DB/source state.
- Do not inject DB rows directly just to make a website display app data unless the scope explicitly defines it as a fixture-only diagnostic.
- If no app is published, public app surfaces must show real empty/unavailable state, not mock cards.
- If no active price/commercial data exists, public/account surfaces must show no-price/not-configured state, not fake price.
- Every visible app name, summary, type, release cue, price, free-use window, billing breakdown, and commercial cue must trace to runtime source evidence when in scope.
- DB-backed write/readback checks need DB/source evidence or an explicit blocker.
- Cleanup should use product/admin actions when the scenario requires user-realistic cleanup. Do not delete directly from DB to hide effects unless environment rules explicitly allow and record it.

## i18n And Locale

When locale is in scope:

- Check every supported locale in the declared scope.
- Verify correct language content.
- Reject unintended mixed-language UI and wrong fallback copy.
- Locale switch must preserve flow, route, UI state, session behavior, and permission behavior.
- Forms, validation, errors, empty states, blocked states, and commercial copy must be correct per locale.

## Tests In QA

Tests are supporting evidence only.

Rules:
- Do not write or edit tests during QA unless the user gives a separate order.
- Do not patch test failures silently.
- Do not use tests to legitimize broken behavior.
- Tests must model real user flows with all required steps.
- If tests mark wrong behavior as expected, report a test defect or test gap.

Examples of wrong expectations:
- A valid account user seeing only `Account surface is locked`.
- A valid admin seeing only `Admin surface is locked`.
- A signed-out user with no clear login/return path.
- A wrong-role user without a clear denied state.

## Evidence

Record evidence as work completes.

Evidence sources may include:
- build output
- runtime health output
- visible-browser proof
- screenshots, videos, traces
- Playwright report from the visible run
- route/control inventory
- visible-surface inventory
- interactive inventory
- click/input action ledger
- no-result/no-op report
- console/network findings
- DB/source readback evidence
- seed or fixture ledger
- source-of-truth mismatch report
- benchmark counts and runtime metrics when benchmarkable

Do not approve if required evidence, benchmarks, inventories, or ledgers are missing.

## Blocker Propagation

If a blocker prevents reaching downstream flows:

- stop if the blocked gate is required for the current run
- report the downstream flows as `Blocked`
- do not mark them passed
- do not imply coverage
- do not invent confidence
- list exactly what remains unverified and why

Example: if the shell never mounts, page-level flows behind the shell are blocked, not verified.

## Finding Classification And Severity

Classify each finding:

| Classification | Meaning |
| --- | --- |
| Confirmed runtime bug | Reproduced on a mounted path. |
| Likely runtime bug | Strong signal, but missing one runtime confirmation. |
| Usability friction | Technically works, but confuses, slows, or misleads a real user. |
| Spec drift | Runtime and docs disagree, but user-visible harm is not fully confirmed. |
| Test gap | Coverage is insufficient to conclude behavior safely, or tests encode the wrong behavior. |

Severity:

| Severity | Meaning |
| --- | --- |
| CRITICAL | User cannot complete a core flow, wrong scoped data is shown, or money/protected action is exposed incorrectly. |
| HIGH | Mounted runtime path is wrong, feature is unreachable, or displayed data is materially incorrect. |
| MEDIUM | Degraded but usable. |
| LOW | Polish only. |

For every finding, include:
- user-visible symptom first
- technical path second, if known
- severity
- classification
- confidence: high / medium / low
- reproducibility: always / intermittent / once
- persona/context/viewport
- route/dialog/flow/control
- expected result
- actual result
- repro steps
- evidence references
- blocked downstream scope
- workaround if one exists
- suggested fix direction if useful, without fixing

## Required QA Report Shape

When writing a QA artifact, create a new file. Do not overwrite old QA reports.

Typical QA report path:

```text
reports/QA/rp_qa_<YYMMDD>_<HHMMSS>_by_<model_slug>_<scope>.md
```

Report template:

```text
# QA Report

Scope:
Runtime:
Browser:
Build evidence:
Anvien evidence:
No-fix boundary: QA reports only. No code fixes were made.

## Runtime Contexts
- persona:
- owner/app/scope:
- role:
- shift/session/subscription:
- dataset/fixture:
- locale:
- viewport:

## Inventory Summary
- visible surfaces:
- interactive elements:
- actions:
- data/source-of-truth checks:
- blocked:
- out of scope:
- unverified:

## Coverage Verdict
- Pass:
- Fail:
- Blocked:
- N/A:
- 100% of declared scope: yes/no

## Action Ledger
| Surface | Route/flow | Control/action | Context | Expected | Actual | Verdict | Evidence |
| --- | --- | --- | --- | --- | --- | --- | --- |

## Findings
### [SEVERITY] Title
Classification:
Confidence:
Reproducibility:
Persona/context/viewport:
Route/dialog/flow/control:
Expected:
Actual:
Repro steps:
Evidence:
Downstream blocked scope:
Suggested fix direction:

## Source-Of-Truth Checks
| Visible value/action | Runtime source/API/DB | Expected source state | Actual source state | Verdict | Evidence |
| --- | --- | --- | --- | --- | --- |

## Console/Network/Runtime Evidence

## Blockers And Unverified Scope

## Handoff
Handoff: <architect|supervisor|coder> - <reason>

## Final Decision
Pass / Fail / Blocked for declared scope only.
```

## Handoff rules:
- (MUST) End every QA report with exactly one concise handoff line.
- Handoff to `architect` when the next step needs architecture, source-of-truth, spec, system-flow, or rule decisions.
- Handoff to `supervisor` when the next step needs acceptance, priority, scope, coordination, or QA/fix order decisions.
- Handoff to `coder` when the issue is a confirmed implementation bug with enough evidence to start a fix.
- If there is no finding or blocker, use `Handoff to supervisor - QA scope completed; awaiting acceptance decision.'

## Artifact And Commit Rules

- QA artifacts may be written if requested by the user or required by the active QA plan.
- Commit QA artifacts only when active repo/user rules explicitly require it.
- If the user says not to commit, do not commit.
- Do not commit screenshots, Playwright artifacts, `.tmp/`, generated evidence, or unrelated files unless explicitly requested.
- QA must stop after reporting and wait for a separate user fix order.

## Red Flags

Stop, report, or re-scope if:

- QA starts fixing code.
- QA uses Anvien, source code, or tests as pass/fail proof.
- QA runs only headless/hidden browser for final visible QA.
- QA skips inventory.
- QA skips a step in an in-scope flow.
- QA clicks only the happy path.
- QA ignores disabled, no-op, rejected, empty, error, stale, or blocked states.
- QA uses fake app/commercial data as if it were real production state.
- QA ignores DB/source readback for DB-backed behavior.
- QA ignores locale/session/permission changes when in scope.
- QA reports downstream scope as passed after an upstream blocker.
