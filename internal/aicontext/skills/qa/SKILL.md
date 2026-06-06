---
name: qa
description: Use when the user asks to run QA without fixing code, including mounted runtime behavior, visible user flows, browser-visible app execution, source-of-truth checks, action/state coverage, route/control inventories, Playwright control sweeps, or QA report generation in repositories where Anvien can support
---

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
- For final visible QA, use a real app/runtime in a browser window visible to the user. Headless or hidden execution cannot approve final visible QA.

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

  ## Coverage Model

  QA coverage is scope-driven, not lane-driven.

  For the declared scope, build and execute these inventories:

  - Runtime surface inventory: pages, routes, tabs, dialogs, drawers, menus, mounted entry points.
  - Control inventory: every reachable user control by role/name/locator and expected outcome.
  - State matrix: visible, hidden, enabled, disabled, loading, empty, error, success, blocked, stale, validation, submitting.
  - Context matrix: persona, role, owner/app scope, session, permission, subscription, viewport, locale.
  - Navigation matrix: links, tabs, redirects, deep links, back/forward, reload, selected nav state.
  - Data/source-of-truth map: displayed values, mutations, APIs, DB/read-model/source checks when relevant.

  A declared scope is not complete until every required inventory item has a verdict or an explicit `Blocked`, `Out of scope`, or `Unverified` mark.

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

## Visible Browser And Playwright

For final visible QA:

- Use installed Chrome or Edge, or another explicitly approved real browser.
- The browser window must be visible, normal, and observable on the user's physical PC.
- Headless, hidden, minimized, offscreen, screenshot-only, sandbox-only, or CI-only execution cannot approve final visible QA.
- If a visible browser cannot be opened or observed, stop and report a blocker.

When Playwright is used for visible QA:

- Drive the real app/runtime URL, not a synthetic component harness unless the scope explicitly allows it.
- Attach to or launch the real visible browser session that the user can observe.
- Clicks, typing, submits, navigation, redirects, reloads, dialogs, and state changes must occur in that visible browser session.
- Screenshots, videos, traces, and Playwright reports are evidence of the visible run, not replacements for it.
- Headless Playwright may support diagnostics or preflight, but cannot be the approval source for visible QA.

## Inventory Before Verdict

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

## Mounted Runtime Map

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

## No-Step-Skipped Flow Rule

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

## State And Action Matrix

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

## Field And Button Data Flow

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

## Source-Of-Truth Data

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
Mode:
QA lanes:
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
