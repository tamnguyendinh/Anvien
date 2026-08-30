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
- Must have screenshot Evidence.
- Tests verify already-correct behavior. Tests do not define expected behavior and must not legitimize wrong UX.
- Wrong tests are findings. A green test that encodes wrong behavior is not QA proof.
- Do not skip steps. Every in-scope flow step matters equally.
- Every in-scope surface, action, state, data path, and blocker needs an explicit verdict or out-of-scope mark.
- An incomplete inventory or action ledger means QA is incomplete.
- If a blocker prevents downstream verification, downstream scope is `Blocked`, not passed.
- For final visible QA, open the real built website, runtime, or generated app artifact the same way a user would open it, in a visible browser or app window on the user's physical PC. Headless, hidden, source-only, test-only, or screenshot-only execution cannot approve final visible QA.
- QA must run on the real, officially built application (actual Electron artifact or release binary).
- Every action (clicking buttons, entering data, changing state) must capture clear PNG screenshots saved sequentially into the `screenshots/` directory.
- Visually inspect each captured image / visual artifact to detect UI and behavioral defects; never rely solely on in-code assertions or generated `.json`/`.md` summary files.
- Never accept a QA report lacking the actual screenshot evidence set for each interaction.

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

## Core Workflow

1. **Preflight** → build, runtime, browser (→ read `references/runtime-and-build-rules.md`)
2. **Inventory** → surfaces, controls, states (→ read `references/coverage-model.md`)
3. **Automation** → screenshots, control sweep (→ read `references/automation-and-evidence-capture.md`)
4. **Data flow** → trace source-of-truth (→ read `references/data-flow-and-source-of-truth.md`)
5. **Locale** → if i18n in scope (→ read `references/i18n-and-locale.md`)
6. **Evidence** → collect, classify blockers (→ read `references/evidence-and-blockers.md`)
7. **Report** → findings, verdict, handoff (→ read `references/reporting.md`)

## Quick Decision Tree

- Preparing build, runtime, or browser? → `references/runtime-and-build-rules.md`
- Running control sweep or capturing screenshots? → `references/automation-and-evidence-capture.md`
- Building inventory or checking state matrix? → `references/coverage-model.md`
- Tracing data flow or DB readback? → `references/data-flow-and-source-of-truth.md`
- Scope includes i18n / localization? → `references/i18n-and-locale.md`
- Recording evidence or handling blockers? → `references/evidence-and-blockers.md`
- Writing report or classifying findings? → `references/reporting.md`

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

## Reference Index

| Need | File |
|------|------|
| Build, runtime, visible browser/app rules | `references/runtime-and-build-rules.md` |
| Screenshot capture, control sweep | `references/automation-and-evidence-capture.md` |
| Inventory, verdict, flow steps, state matrix | `references/coverage-model.md` |
| Data tracing, DB, source-of-truth | `references/data-flow-and-source-of-truth.md` |
| i18n / locale scope | `references/i18n-and-locale.md` |
| Evidence sources, tests, blocker propagation | `references/evidence-and-blockers.md` |
| Finding classification, report template, handoff | `references/reporting.md` |
