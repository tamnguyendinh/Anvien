# Automation And Evidence Capture

> This file is part of QA Skill v2. Read when: capturing screenshots, running control sweeps, or driving browser/desktop automation during QA.

## Screenshot Evidence Rule

- When browser or desktop automation is used for QA, screenshots must be captured.
- Do not use the final screenshot after failure as the main evidence.
- Capture screenshots at each small action step: before entering data, after each field, before clicking, after clicking, and after the UI responds or settles.
- After the run finishes, open and visually inspect the screenshots to determine exactly which step first introduced the issue.
- Bugs are not necessarily blockers. If a bug is found, report it in the report/evidence section, but continue testing if a valid path remains.
- Review the whole screen for additional issues, including controls that do not respond, inputs that reject data, incorrect results, overlapping text/cards, broken fonts, overflow, disappearing elements, or layout shifts.

## Automated Control Sweep

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
