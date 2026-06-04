# QA Source Rules Analysis For Anvien Skill Design

Date: 2026-06-04

Scope: analyze the supplied QA documents and extract a complete QA operating model that can later be turned into an Anvien-compatible QA skill.

Boundary: this report is analysis only. It does not execute QA, does not create tests, and does not fix product code.

## Source Files Reviewed

| Source | Role in the QA rule set |
| --- | --- |
| `E:\Website\DOCS\plans\2026-05-29-qa-visible-browser-test-run-plan.md` | Strongest visible-browser QA execution plan. Adds full build, real runtime, visible PC browser, action ledger, evidence, benchmark, source-of-truth, and no-fix constraints. |
| `E:\owner-tool\clockwork\skills\QA\General\QA.md` | General QA runtime review skill. Defines modes, runtime-first doctrine, surface inventory, mounted runtime map, state/action matrix, classification, severity, report format, and artifact rules. |
| `E:\owner-tool\clockwork\skills\QA\General\QA-Route-Mount-And-Navigation.md` | Lane-specific QA for mounted shell routes, tabs, sub-tabs, launchers, redirects, default landings, breadcrumbs, and navigation blocked states. |
| `E:\owner-tool\clockwork\skills\QA\General\QA-Table-List-And-Row-Actions.md` | Lane-specific QA for tables, lists, card grids, row actions, toolbar actions, bulk actions, menus, selection, search, sort, filter, pagination. |
| `E:\owner-tool\clockwork\skills\QA\General\QA-Dialogs-Forms-And-Submits.md` | Lane-specific QA for dialogs, drawers, modals, forms, fields, validation, dirty state, submit/cancel/close/reopen behavior. |
| `E:\owner-tool\clockwork\skills\QA\General\QA-Displayed-Data-And-View-States.md` | Lane-specific QA for summaries, read-only values, statuses, counters, timestamps, loading, empty, error, success, disabled, blocked, stale, refresh states. |
| `E:\owner-tool\clockwork\skills\QA\General\QA-End-To-End-User-Flows-And-Usability-Friction.md` | Lane-specific QA for cross-surface user journeys, task continuity, terminal states, recovery branches, and usability friction. |
| `E:\owner-tool\clockwork\skills\QA\General\QA-Visual-Fidelity-And-Layout.md` | Lane-specific QA for mounted visual fidelity, layout, spacing, alignment, containment, hierarchy, text wrapping, clipping, responsive/resize behavior. |
| `E:\owner-tool\clockwork\skills\QA\General\QA-Context-Switch-And-Blocked-UX.md` | Lane-specific QA for owner/app/role/shift/session/subscription context changes and the blocked UX that follows those changes. |

## Executive Conclusion

The supplied files define QA as a runtime observation and reporting discipline, not as implementation work. QA must answer whether a real user can reach the mounted surface, perform the visible action, see the correct data/state, understand blocked/error outcomes, and recover from normal mistakes or context changes.

The documents are not primarily test-writing instructions. Tests, code inspection, screenshots, and graph analysis can support QA, but none of them replace mounted runtime interaction. The visible-browser plan is even stricter: final QA approval requires a visible Chrome or Edge window on the user's physical PC, a full build first, real runtime data, explicit inventories, action ledgers, evidence links, and no code fixes during the QA run.

For an Anvien-compatible QA skill, the right model is:

1. Use Anvien before execution to map mounted routes, entry points, APIs, components, data flow, and likely risk surfaces.
2. Treat Anvien output as scope and impact evidence, not as pass/fail proof.
3. Execute QA through the real mounted runtime path.
4. Record a verdict and evidence for every inventoried surface/action/state.
5. Report bugs and blockers only. Do not repair code in QA mode.

## Non-Negotiable Rules Extracted

### 1. QA Does Not Fix Code

The visible-browser plan explicitly says not to fix bugs during QA unless the user gives a separate fix order after reviewing QA results. Its non-goals also say not to edit production code, SPEC, UI/UX design docs, or silently patch test failures during the QA run.

The final QA skill must make this an iron rule:

- QA may inspect code, graph, routes, logs, DB rows, screenshots, traces, and runtime output.
- QA may report likely fix direction.
- QA must not edit product code, generated code, tests, docs/specs, or fixtures as part of QA.
- If a blocker prevents QA, QA records the blocker and stops or marks downstream scope as blocked.
- Fix work must be a separate user-approved implementation task.

### 2. Runtime Is The Source Of QA Truth

Every QA file repeats the same zero-trust model:

- Do not trust component tests by themselves.
- Do not trust screenshots without reproducing the flow.
- Do not trust "wired" claims without checking the mounted runtime path.
- Do not substitute code reading for runtime interaction.
- A route/dialog/component that exists in code but is not mounted/reachable through the real shell is not covered.

This means source code can identify expected surfaces, but runtime behavior decides QA pass/fail.

### 3. Tests Verify Behavior, They Do Not Define Behavior

The visible-browser plan's notes say code must first be correct as a real website/app, then tests verify that behavior. Tests must model full real user flows and cannot be written or edited to legitimize wrong UX.

For the final skill:

- Green tests never replace QA.
- Wrong tests are a finding or follow-up, not proof that behavior is correct.
- QA can mention test gaps, but must not treat tests as the source of truth.

### 3A. Tests Must Model The Full Real User Flow

Tests are only useful after the real behavior is understood and correct. They must simulate the real user journey with all required steps. They must not be shortened to skip meaningful behavior, and they must not be written to legitimize an incorrect runtime or UX result.

The final skill should enforce:

- tests must follow the same flow a real user follows;
- tests must include the real route, visible controls, inputs, submits, redirects, state changes, and expected post-action UI;
- tests must not mark a broken or incomplete UX as correct just because the current implementation behaves that way;
- tests are verification evidence only, never the authority for expected behavior.

### 3B. Wrong Tests Must Be Treated As Wrong

If an existing test encodes the wrong product behavior, QA must not accept that test as truth. The visible-browser plan notes specific examples: a test that treats `Account surface is locked` or `Admin surface is locked` as correct for valid user/admin flows must be reviewed.

Expected behavior model:

- signed-out users should have a clear login/return path;
- valid account users should reach the account surface;
- valid admins should reach the admin surface;
- wrong-role users should receive a clear denied state;
- invalid, expired, or rotated sessions should show the appropriate blocked/re-auth path;
- tests that disagree with this behavior are test defects or follow-up findings, not QA proof.

### 4. Build And Runtime Gate Before QA

The visible-browser plan requires a full build before QA testing and use of the real runtime assigned for QA. If Docker real runtime is used, it must be rebuilt and started with the intended environment. The user notes also require production/prod-like runtime checks on the URL/UI users actually use.

Final skill rule:

- Before QA execution, record build command/result, runtime command/result, runtime URL, health state, seed/fixture state, and browser launch state.
- If the build or runtime cannot be prepared, mark QA blocked.
- Do not approve a run based only on dev/test shortcuts when production-like runtime is required.

### 4A. Production Must Be Built And Checked On The Real User URL/UI

When the QA scope is production or production-like runtime, QA must not rely on a dev server, test-only route, stale container, or unrebuilt artifact.

The final skill should require:

- run the real production build command before QA when production behavior is under review;
- rebuild Docker or the equivalent production-like runtime when that runtime is the target;
- verify behavior on the URL and UI that users actually use;
- record the build/runtime command evidence;
- mark QA blocked if the real build/runtime cannot be produced;
- never approve production behavior from a dev/test-only shortcut.

### 5. Visible Browser Is Mandatory For Final Visible QA Approval

The visible-browser plan adds a stricter gate than the generic QA skills:

- Final visible QA must use installed Chrome or Edge.
- The browser window must be visible on the user's physical PC.
- Headless, hidden, minimized, offscreen, screenshot-only, sandbox-only, or CI-only browser execution is not acceptable for final approval.
- If the visible browser cannot be opened and observed, stop and record a blocker.

Final skill should distinguish two levels:

| QA level | Browser requirement |
| --- | --- |
| General runtime QA | Browser automation or manual runtime interaction may be used, but mounted runtime interaction is still required. |
| Final visible-browser QA | Installed Chrome/Edge visible on the user's physical PC is mandatory; headless output cannot approve the run. |

### 5A. Playwright Must Exercise The Real App In A User-Observable Browser When Visible QA Is Required

The QA files mention Playwright as the mandatory E2E runner for declared scopes, but the visible-browser plan adds an important constraint: Playwright execution is not enough by itself if it is headless, hidden, minimized, offscreen, CI-only, or detached from the real user-observable browser session.

For visible QA, Playwright must be used to drive the real app/runtime in a real browser window that the user can see on their machine. Acceptable execution means:

- the real app/runtime URL is opened, not a synthetic component harness unless the scope explicitly allows that;
- the browser is installed Chrome or Edge, or another explicitly approved real browser;
- the browser window is visible, normal, and observable on the user's physical PC;
- clicks, typing, submits, navigation, redirects, reloads, dialogs, and state changes happen in that visible browser session;
- screenshots, videos, traces, and Playwright reports are evidence of that visible run, not a replacement for it;
- if Playwright cannot attach to or open a visible browser that the user can observe, final visible QA must stop and report a blocker.

Headless Playwright can still be useful for diagnostics, preflight, or supporting automation when the user did not request final visible-browser QA. It cannot be used as the approval source for a visible QA run.

### 6. Mode Dispatch Controls Scope

The QA skill files define three modes:

| Mode | When used | Main anchor |
| --- | --- | --- |
| Mode 1 - Phase/Job QA Review | Default when any phase/job in `Docs/execution/*` still lacks required checks. | Declared phase/job plus exact `Docs/SPEC/*` family. |
| Mode 2 - Post-Completion QA Review | Only when backlog review path is exhausted; covers bug hunts, follow-ups, rejects/resubmits, current worktree, post-completion scope. | Exact `Docs/SPEC/*` family plus mounted runtime path for current scope. |
| Mode 3 - Post-Completion QA With Supplement Plan | Only when backlog is exhausted and `reports\QA\QA+EDGE_CASE_TEST_PLAN.md` exists with usable content. | Supplement plan cluster/part scope plus exact SPEC and mounted runtime path. |

The important rule is not "read everything". The rule is choose exactly one mode and do only that mode's workflow. Do not mix mode workflows.

### 7. Fresh Independent Rerun

Most QA lanes require every QA turn to be a fresh independent rerun on the current head. Prior reports must not become checklists, hints, seeds, templates, or tie-breakers.

Mode 3 allows reading the most recent previous QA report only to locate sequential progress marks. Its content still cannot be used as QA context.

Final skill implication:

- QA should rediscover still-live bugs naturally.
- If an old bug was fixed, QA continues to look for current broken behavior.
- QA conclusion must come from current rules, current spec/scope, current runtime, and current evidence.

### 8. Inventory Before Verdict

The documents repeatedly forbid claiming coverage without a complete inventory. Each lane defines its own surface inventory, but the common rule is:

- Enumerate every mounted or mountable runtime surface in the declared scope.
- Each surface/action/state must have an explicit verdict.
- `100% coverage` means `100% of declared scope`, not the whole app unless the declared scope is the whole app.
- Excluded surfaces must be explicitly `Out of scope`.
- Unreachable surfaces due to blockers are `Blocked`, not covered.
- If the ledger is incomplete, the QA pass is incomplete.

The visible-browser plan adds:

- `visible-surface inventory` and `interactive inventory` are deliverables.
- Unaccounted visible surface target is `0`.
- Untested interactive element target is `0`.
- Every visible clickable/input control needs an action-ledger row, including disabled, no-op, no-result, and rejected-submit controls.

### 9. Mounted Runtime Map

For each inventoried surface, the QA files require a mounted runtime map. Common fields:

- runtime entry point
- user trigger
- owner
- active scope
- role
- shift state
- session/subscription state, if relevant
- dataset or fixture
- data source or store used by the mounted path
- expected visible actions
- expected blocked actions

This is where Anvien can help strongly: graph analysis can reveal likely routes, components, handlers, stores, APIs, and data source paths before the visible run. But QA still has to verify the mounted runtime result.

### 10. State And Action Matrix

Every lane requires more than the happy path. QA must check states and actions that apply to each surface.

Common states:

- visible / hidden
- enabled / disabled
- loading
- empty
- error
- success
- permission-gated
- shift-gated
- blocked
- stale-context after owner/scope/role/shift/session/subscription changes

Common action checks:

- user can actually reach the action
- action triggers the correct mounted behavior
- blocked actions explain why they are blocked
- displayed data/state updates after action
- retry, back, reopen, refresh, context switch, and resubmit do not leave stale UI behind

### 10A. No Step May Be Skipped In A Flow

The visible-browser plan's notes add a stronger step discipline: in each flow, every step matters equally. QA must not jump over steps just because a later state looks correct.

Each in-scope flow must account for:

- opening the page or app;
- initial state;
- every tab;
- every link;
- every button;
- every field;
- every form;
- every submit;
- every redirect;
- every refresh, back action, or route change;
- every locale switch;
- every cookie or session change;
- every DB or read-model read/write when relevant;
- UI after each step;
- where the next step leads.

If any of these steps exists in the declared scope but was not observed or ledgered, the flow is not fully covered.

### 10B. Each Field And Button Must Have Data-Flow Understanding

QA must understand what each field and button does to data. This is not optional for forms, submits, money-sensitive behavior, source-of-truth checks, or any user-visible state that depends on runtime data.

For each relevant field/button/action, QA should know and record when applicable:

- where the data comes from;
- which field receives the input;
- which endpoint, action, tool, or submit path receives it;
- what frontend validation runs;
- what backend/server validation runs;
- which table, read-model, store, cache, or state is read;
- which table, read-model, store, cache, or state is written;
- where the app redirects or renders after the action;
- what the next UI allows the user to do;
- how public, account, admin, or user surfaces reflect the change.

Anvien can help map this data flow, but the runtime UI and source-of-truth evidence must still verify it.

### 11. Evidence Is Required, Not Optional

Required evidence sources across the documents:

- screenshots, videos, and traces from runtime execution
- visible-browser proof for visible QA
- Playwright report when used
- route/control inventory
- visible-surface inventory
- interactive inventory
- click/input action ledger
- no-result/no-op report
- console/network findings
- DB evidence for DB-backed writes/readbacks
- seed or fixture ledger
- source-of-truth mismatch report for public app/catalog/release/commercial data
- benchmark counts and runtime metrics when benchmarkable

Screenshots alone are not sufficient if the plan requires visible browser observation.

### 12. Source-Of-Truth Data Must Be Real

The visible-browser plan is strict about app/catalog/release/commercial data:

- Runtime data must be treated as production state, not demo decoration.
- If admin has not created/published an app, public app surfaces must show real empty/unavailable state, not mock app cards.
- If active price/commercial data is not configured, public/account surfaces must show no-price/not-configured state, not fake prices.
- Visible public app name, summary, type, release cue, price, free-use window, billing breakdown, and commercial cue must trace to runtime DB/source evidence.
- Data should be created through the admin/user flow when the QA scenario requires it, not by direct DB injection to make the UI look populated.

### 12A. App Data Must Not Be Fabricated

Public app data must appear only when it exists as real runtime state created through the correct product flow. QA must not inject DB rows directly just to make the website display app data.

The final skill should require:

- public app surfaces show apps only after app upload, publish, or data-entry flow creates valid runtime data;
- public app data must exist in DB/source state as a real app;
- direct DB injection is not acceptable as the way to make the UI look populated unless the scope explicitly defines it as a fixture-only diagnostic;
- test cleanup must use the correct admin/product actions when the scenario requires user-realistic cleanup;
- QA must not delete straight from DB to hide effects unless the test environment rules explicitly allow that cleanup and it is recorded.

### 12B. i18n Must Be Correct Across All Languages

Locale behavior is part of QA coverage when locale is in scope. QA must not approve a surface that only works in one language while another language has mixed text, wrong fallback, broken route behavior, or incorrect session/permission behavior.

The final skill should check:

- every supported locale in the declared scope;
- text content is in the correct language;
- no unintended mixed-language UI;
- no wrong fallback copy;
- locale switch preserves the correct flow;
- locale switch preserves route, UI state, session behavior, and permission behavior;
- forms, validation, errors, empty states, blocked states, and commercial copy are correct per locale.

### 13. Blocker Propagation

If a blocker prevents downstream QA:

- stop or mark downstream flows as `Blocked`
- do not imply coverage for unreached screens
- do not mark downstream scope as passed
- do not invent confidence
- include why it was blocked and what remains unverified

Example model: if the shell never mounts, page-level flows behind the shell are blocked, not verified.

### 14. Findings Must Be Classified

Finding classifications:

| Classification | Meaning |
| --- | --- |
| Confirmed runtime bug | Reproduced on a mounted path. |
| Likely runtime bug | Strong signal, but missing one runtime confirmation. |
| Usability friction | Technically works, but confuses, slows, or misleads a real user. |
| Spec drift | Runtime and docs disagree, but user-visible harm is not fully confirmed. |
| Test gap | Coverage is insufficient to conclude behavior safely. |

Severity guide:

| Severity | Meaning |
| --- | --- |
| CRITICAL | User cannot complete core flow, wrong scoped data is shown, or money/protected action is exposed incorrectly. |
| HIGH | Mounted runtime path is wrong, feature is unreachable, or displayed data is materially incorrect. |
| MEDIUM | Degraded but still usable. |
| LOW | Polish only. |

Every issue should include severity, repro steps, expected result, actual result, affected route/dialog/flow, evidence references, classification, confidence, reproducibility, workaround if any, and suggested fix direction if requested by the plan.

## Lane-By-Lane Analysis

### General Runtime QA

`QA.md` is the base protocol. It owns end-to-end checks, mounted route/dialog/page verification, observable runtime behavior, displayed data, empty/loading/error/disabled states, and user-visible regressions.

Important extracted behavior:

- QA must be from the user point of view.
- The runtime path matters more than isolated components.
- Surface inventory plus mounted runtime map must precede coverage claims.
- Every inventoried surface needs `Pass`, `Fail`, `Blocked`, or `N/A`.
- Every report needs coverage counts and explicit unverified surfaces.

This file should be the backbone of the final Anvien QA skill.

### Route Mount And Navigation

This lane narrows QA to mounted shell navigation:

- shell-backed routes/pages
- tabs and sub-tabs
- sidebar/top nav/breadcrumb/launcher entries
- redirects and default/fallback landing targets
- blocked navigation notices

Core risks:

- route points to wrong or legacy screen
- shell mounts but target page/tab never renders
- navigation exists in code but is unreachable in runtime
- wrong role can see navigation
- stale route after context switch

Anvien fit:

- Use graph route maps to enumerate route handlers, page components, and consumers.
- Then verify with visible runtime navigation, including back/forward/deep-link/selected-state behavior.

### Table, List, And Row Actions

This lane owns collection surfaces:

- tables, lists, card grids
- row actions, overflow/kebab/context menus
- toolbar and bulk actions
- search, sort, filter, pagination, selection, expansion
- selection count and summaries

Core risks:

- collection mounts but expected rows/cards do not render
- row action targets wrong record
- bulk/menu action visible to wrong role
- selection/search/sort/filter state is stale after context switch
- selected count diverges from visible selection

Anvien fit:

- Use graph queries to map list components to stores/API calls and row action handlers.
- QA must still click visible row actions and verify target row/selection at runtime.

### Dialogs, Forms, And Submits

This lane owns form lifecycle:

- dialogs, drawers, modals, inline forms, dedicated form pages
- open/reopen/close/cancel/back-out/discard/retry
- fields, toggles, selectors, steppers, validations
- primary submit and secondary actions
- default, pristine, dirty, validation, submitting, success, error, blocked states

Core risks:

- wrong dialog/form target opens
- required fields/actions never become interactable
- cancel/close/back silently loses changes or leaves stale dirty state
- submit enabled when required fields are missing or role/session/shift should block it
- validation appears late, on wrong field, or not at all
- reopen shows stale values or stale validation
- submit appears successful but hides real error or lands on wrong result

Anvien fit:

- Use file-context and API route-map/tool-map to identify submit path, validation path, API/mutation, and data writes.
- QA must still perform visible form actions and verify immediate post-submit runtime state.

### Displayed Data And View States

This lane owns user-visible data correctness:

- summary cards, stat tiles, counters, badges, chips
- read-only detail panels and fields
- table/list/card values
- labels, totals, amounts, statuses, timestamps, placeholders
- loading/empty/error/success/disabled/blocked/stale/refresh states

Core risks:

- stale values from previous owner/scope/role/entity/session
- wrong counters/totals/status chips for visible context
- loading never resolves or placeholder data is shown as real
- state appears with wrong explanation
- refresh/success state does not reflect actual visible change
- labels/placeholders/status mappings are inconsistent

Anvien fit:

- Use Anvien to trace displayed values to stores, selectors, APIs, and DB/read-models.
- QA must verify that runtime data belongs to the current context and, when required, prove it with DB/source evidence.

### End-To-End User Flows And Usability Friction

This lane owns whole tasks across surfaces:

- task entry, intermediate handoffs, terminal states
- route/list/detail/dialog/form/summary/result handoffs
- start, progress, submit, success, failure, block, cancel, back-out, retry, reopen, resume
- usability friction and recovery burden

Core risks:

- user starts a task but gets stranded
- flow advances but lands on wrong destination or completion state
- cancel/back/retry/reopen/resume loses context or duplicates work
- user cannot tell if task completed, failed, blocked, or is pending
- flow works only if user knows hidden/non-obvious sequence
- blocked/error states provide no recovery path

Anvien fit:

- Use Anvien to identify execution flows and cross-surface call paths.
- QA must run the complete visible user journey and record each handoff, branch, and terminal outcome.

### Visual Fidelity And Layout

This lane owns visible layout quality:

- page/section/panel/card/table/list/dialog layout
- hierarchy, spacing, alignment, density, grouping, containment
- typography, labels, truncation, wrapping, clipping, icon placement, dividers
- visual state presentation for loading/empty/error/success/disabled/blocked/warning
- resize/reopen/refresh/context-update aftermath

Core risks:

- mounted surface hierarchy differs from intended design/spec
- spacing/alignment/grouping misleads or breaks structure
- text clips, wraps, truncates, or overflows incorrectly
- dialog/drawer/modal has wrong geometry or footer/header layout
- state presentation has wrong emphasis or missing affordance
- refresh/resize/context change leaves stale/collapsed/broken layout

Anvien fit:

- Anvien has limited direct visual proof value, but can identify owning components and routes.
- Final judgment must come from visible render evidence, screenshots, and viewport/resize checks.

### Context Switch And Blocked UX

This lane owns context-sensitive UI correctness:

- owner/app type/app scope/role/shift/session/subscription switchers
- selected-context labels, badges, banners, notices, tooltips
- disabled actions and blocked-state explanations
- stale permission/shift/session/subscription symptoms
- switch/back/retry/refresh/reopen behavior after context changes

Core risks:

- context label updates but mounted data/actions still belong to previous context
- role/shift/session/subscription change does not update enabled/disabled actions
- blocked action gives no reason, wrong reason, or stale reason
- partial update leaves labels, banners, or controls stale
- user can trigger action that should be blocked
- switching back and forth leaves stale disabled state or warning text

Anvien fit:

- Use Anvien to map context providers, gates, permission checks, session/shift/subscription code paths, and consumers.
- QA must switch context at runtime and verify labels, data, enabled state, blocked explanations, and stale-state absence.

## Derived Complete QA Workflow

This is the workflow a final Anvien-compatible QA skill should implement.

### P0 - Select Scope And Mode

Inputs:

- user request
- `AGENTS.md`
- active execution plan or SPEC family
- optional QA supplement plan
- current repo rules

Actions:

1. Determine whether this is Phase/Job QA, post-completion QA, or supplement-plan QA.
2. Select only one mode.
3. Declare exact scope, expected source documents, runtime entry assumptions, and lane(s) to run.
4. State explicitly that QA will not fix code.

Output:

- selected mode
- declared scope
- source-of-expected-behavior
- QA lanes selected
- no-fix boundary

### P1 - Anvien-Assisted Scope Map

Use Anvien only as code intelligence, not as QA proof.

Required first step before graph-based work:

```text
anvien analyze --force
```

Useful command mapping:

| When QA needs to... | Use |
| --- | --- |
| Refresh graph evidence | `anvien analyze --force` |
| Find route/page/component ownership | `anvien query "<concept>" --repo <repo>` |
| Find files relevant to a route or behavior | `anvien query files "<concept>" --repo <repo>` |
| Inspect one file deeply | `anvien file-context <path> --repo <repo> --json` |
| Inspect route handlers and consumers | `anvien api route-map [route] --repo <repo>` |
| Inspect API shape drift where UI consumers depend on it | `anvien api shape-check [route] --repo <repo>` |
| Inspect tool/IPC paths if the app uses tools | `anvien api tool-map [tool] --repo <repo>` |
| Understand likely blast radius for suggested fix direction | `anvien impact file <path> --repo <repo> --direction upstream` or `anvien impact symbol "<symbol>" --repo <repo> --direction upstream` |
| Audit graph health if Anvien evidence looks stale/broken | `anvien graph-health summary --repo <repo> --json` |

Expected Anvien output for QA:

- route/page inventory candidates
- mounted entry candidates
- action/handler/API candidates
- data source/store/API/DB candidates
- permission/context gate candidates
- source files for likely fix direction

Important limitation:

- Anvien evidence can tell what should be checked.
- Anvien evidence cannot mark a runtime surface as passed.

### P2 - Runtime Preflight

Actions:

1. Run full build required by repo/plan.
2. Start real runtime or production-like runtime.
3. Record runtime URL, command, health, environment, seed/fixture state.
4. If production-like runtime is the target, rebuild Docker or equivalent runtime before checking the user-facing URL/UI.
5. If final visible QA, open installed Chrome/Edge visibly on user's physical PC and prove it is visible.
6. If Playwright is used for visible QA, attach to or launch that real visible browser and drive the real app/runtime there.
7. If visible browser cannot be confirmed when required, stop and report blocker.

Output:

- build evidence
- runtime evidence
- browser evidence
- blocker if any

### P3 - Build Inventories

Inventory types:

- visible-surface inventory
- interactive inventory
- route/control inventory
- state inventory
- action ledger template
- data/source-of-truth map
- persona/context matrix

Each inventory item should include:

- surface/action/state name
- route or flow
- runtime entry
- user trigger
- persona/context
- expected behavior
- data source if relevant
- evidence slot
- verdict slot

No QA verdict can be final until every in-scope inventory item has a verdict or explicit out-of-scope/blocker status.

### P4 - Execute Runtime QA

Run visible user behavior:

- enter through the real shell or public runtime path
- when Playwright is used for visible QA, drive the real app in a real browser window visible to the user
- navigate through visible controls, not only direct URLs
- click every visible relevant action
- type/select/toggle every relevant input
- submit, cancel, close, back out, reopen, retry
- switch context where relevant
- test happy, loading, empty, error, blocked, stale, permission/shift/session/subscription states
- record screenshot/video/trace and action-ledger row
- do not skip any in-scope flow step: page/app open, initial state, tabs, links, buttons, fields, forms, submits, redirects, refresh/back/route changes, locale switches, cookie/session changes, DB/read-model read/write, UI after each step, and next destination

For final visible QA:

- do not approve with headless-only evidence
- screenshots support the visible run but do not replace it

### P5 - Validate Data Flow And Source Of Truth

For every money-sensitive, scope-sensitive, account-sensitive, public commercial, or DB-backed result:

- confirm current runtime context
- trace visible data to source/API/DB/read-model when required
- prove DB writes/readbacks when the QA plan requires it
- distinguish real empty/no-price/unavailable state from mock/hardcoded decoration
- record source-of-truth mismatch if visible UI disagrees with runtime source
- map each relevant field/button from input source through endpoint/action, frontend/backend validation, read/write table or state, redirect/render result, next UI affordance, and public/account/admin/user reflection
- ensure app data is created/published/configured through the correct product/admin flow when the scenario requires realistic data, not by direct DB injection
- verify locale-sensitive data, routes, session behavior, permission behavior, messages, and fallback behavior across every locale in scope

### P6 - Classify Findings

For each issue:

- user-visible symptom first
- technical path second if known
- severity
- classification
- confidence
- reproducibility
- route/dialog/flow
- persona/context/viewport/control
- expected vs actual
- evidence references
- blocked downstream scope
- suggested fix direction when requested, without fixing

### P7 - Write Artifact And Stop

Artifacts:

- QA report normally goes in `reports\QA`.
- Shared blockers may create `reports\problem\pb_qa_<timestamp>_<scope>.md`.
- This current analysis report is intentionally in `reports\problem` because the user requested analysis, not an executed QA run.

Rules:

- Create a new artifact, do not overwrite old QA reports.
- Stage/commit only if the active QA role or user explicitly requires artifact commit.
- Do not commit screenshots/traces/temp files unless explicitly requested.
- Stop after report and wait for the user's fix order.

## Proposed Final Skill Structure

Recommended skill name:

```text
anvien-qa-runtime-review
```

Recommended description:

```text
QA specialist for mounted runtime behavior, visible user flows, route/control reachability, displayed data correctness, state/action coverage, source-of-truth evidence, and Anvien-assisted scope mapping. Use when the user asks whether an app actually works from a user point of view. QA reports only and does not fix code.
```

Recommended sections:

1. Identity and no-fix boundary
2. Source-of-truth hierarchy
3. Mode dispatch
4. Anvien-assisted mapping rules
5. Build/runtime/browser preflight
6. Lane selection guide
7. Surface inventory rules
8. Mounted runtime map rules
9. State/action matrix rules
10. Data/source-of-truth rules
11. Evidence rules
12. Blocker propagation
13. Finding classification and severity
14. Required output format
15. Artifact and commit rules

## Lane Selection Guide For The Future Skill

| When you need to QA... | Use lane |
| --- | --- |
| Whether users can reach the correct screen through shell routes, tabs, sub-tabs, breadcrumbs, launchers, redirects, or default landing | Route Mount And Navigation |
| Tables, lists, cards, row actions, toolbar actions, bulk actions, contextual menus, selection, search, sort, filter, pagination | Table List And Row Actions |
| Dialogs, drawers, modals, forms, fields, validation, dirty state, submit, cancel, close, discard, reopen | Dialogs Forms And Submits |
| Summary cards, counters, read-only values, statuses, labels, timestamps, placeholders, loading/empty/error/success/blocked/stale states | Displayed Data And View States |
| Full tasks crossing multiple surfaces, task continuity, terminal states, recovery paths, and usability friction | End-To-End User Flows And Usability Friction |
| Layout, spacing, alignment, hierarchy, text wrapping/clipping, containment, visible state presentation, responsive/resize aftermath | Visual Fidelity And Layout |
| Owner/app/role/shift/session/subscription switching, stale context, disabled/blocked actions, blocked explanations | Context Switch And Blocked UX |
| Whole app or broad user-visible behavior | General Runtime QA plus relevant lanes above |

## Output Template Recommended For The Skill

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

## Final Decision
Pass / Fail / Blocked for declared scope only.
```

## Rule Precedence Proposed For The Skill

1. User's explicit QA scope and no-fix order.
2. Repo `AGENTS.md`.
3. Active QA plan or SPEC family.
4. Visible-browser plan, if final visible QA is requested or the plan is in scope.
5. General QA protocol.
6. Lane-specific QA protocol.
7. Anvien graph evidence as support only.
8. Automated test output as verification evidence only, never source of truth.

## Important Design Decisions

### Use One Skill With Lanes, Not Many Separate Skills

The current source files duplicate the same protocol with lane-specific substitutions. A final Anvien skill should avoid duplication by using one core workflow plus lane modules. This prevents contradictory drift between route, table, form, data, visual, e2e, and context rules.

### Visible Browser Should Be A Gate, Not Always The Only Browser Mode

The visible-browser plan is clearly mandatory for final visible QA approval. For smaller QA asks, the skill can use browser automation or headed Playwright if the user did not ask final visible QA. But when the source plan or user asks visible final QA, the skill must stop if the visible PC browser cannot be confirmed.

### Anvien Should Drive Discovery, Not Verdicts

Anvien is valuable for:

- finding mounted paths and likely hidden surfaces
- reducing missed route/action/data-flow coverage
- understanding source ownership for suggested fix direction
- mapping API and data contracts before runtime checks

Anvien is not enough for:

- confirming that users can reach a surface
- confirming visual layout
- confirming click/input result
- approving coverage
- replacing browser evidence

### QA Reports Bugs; Implementation Fixes Start Later

The skill must explicitly stop after reporting. If the user orders fixes later, the agent changes role from QA to implementation and must then follow code-edit rules, including Anvien freshness, impact-before-edit, full build, tests, detect-changes, and commits as required.

## Open Items Before Turning This Into A Skill

1. Decide whether the new skill should always write QA artifacts to `reports\QA` or allow user-selected report directories.
2. Decide whether QA artifact commits should be automatic or only when repo/user rules explicitly require them.
3. Decide whether final visible-browser QA requires physical PC observation in all repos or only when a visible-browser plan/scope requests it.
4. Decide whether Anvien graph inventory counts should be mandatory evidence for broad QA, or optional support for narrow QA.
5. Decide whether the future skill should include exact command examples for Chrome/Edge visible launch per repo, or keep that repo-specific.

## Bottom Line

The QA files define a strict, evidence-led runtime review process:

- understand scope with docs and Anvien
- build and run the real app
- verify mounted runtime paths through visible user interaction
- inventory every surface/action/state before claiming coverage
- prove data with runtime source evidence when relevant
- classify findings with severity and reproducibility
- report only
- do not fix code during QA

This is enough to create a complete Anvien-compatible QA skill, as long as the final skill preserves the no-fix boundary and treats Anvien/tests as support tools instead of QA truth.
