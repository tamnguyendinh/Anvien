# UI-BE Binding Skill Pack

A strict skill pack for connecting backend data to an approved frontend/UI without destroying the approved design.

## What This Solves

AI coding agents often break polished UI during backend integration by adding:

- MVP/demo/sample text
- helper notes
- loading/empty/error panels
- backend/API explanations
- generic dashboard blocks
- unapproved badges/cards/sections
- layout and typography drift

This skill pack prevents that by forcing backend data to enter only approved UI slots.

## Files

- `SKILL.md` — main skill definition
- `docs/ui-authority.template.md` — UI authority template
- `docs/actual-wiring-status.template.md` — P0 current-wiring audit template
- `docs/ui-slot-map.template.md` — approved data binding slot map
- `docs/state-map.template.md` — approved loading/empty/error state map
- `docs/backend-contract-map.template.md` — endpoint-to-slot map
- `docs/visible-text.lock.example.json` — visible text lock example
- `prompts/codex-ui-be-binding-task.md` — Codex task prompt for backend binding
- `prompts/codex-ui-port-pixel-lock-task.md` — Codex task prompt for prototype-to-TS porting
- `tests/playwright/approved-ui.guard.spec.ts` — Playwright visual/text guard
- `tests/playwright/slot-binding.guard.spec.ts` — Playwright slot binding guard
- `checklists/pre-coding-checklist.md` — required pre-coding checklist
- `checklists/final-verification-checklist.md` — final verification checklist
- `reports/ui-be-binding-report.template.md` — final report template

## Recommended Usage

1. Copy this folder into your repo under `.claude/skills/ui-be-binding/`, `.codex/skills/ui-be-binding/`, or `docs/skills/ui-be-binding/`.
2. Create or use the active plan directory: `<active-plan-dir> = docs/plans/YYYY-MM-DD-<plan-name>`.
3. Copy the template docs into `<active-plan-dir>` and rename them:
   - `<active-plan-dir>/YYYY-MM-DD-ui-authority.md`
   - `<active-plan-dir>/YYYY-MM-DD-actual-wiring-status.md`
   - `<active-plan-dir>/YYYY-MM-DD-ui-slot-map.md`
   - `<active-plan-dir>/YYYY-MM-DD-state-map.md`
   - `<active-plan-dir>/YYYY-MM-DD-backend-contract-map.md`
   - `<active-plan-dir>/YYYY-MM-DD-visible-text.lock.json`
4. Fill `<active-plan-dir>/YYYY-MM-DD-actual-wiring-status.md` for the target surface before writing implementation steps.
5. Fill the slot map before backend integration.
6. Add the Playwright tests and adjust routes/selectors to match the app.
7. Give AI Agent the prompt in `prompts/codex-ui-be-binding-task.md`.
8. Reject any handoff that does not include current-wiring audit, build/test/screenshot/text-lock evidence.

## Command Discipline

Do not give the agent a broad task like:

```text
Connect backend to frontend.
```

Use:

```text
Use the UI-BE Binding Skill.
Bind only these approved slots: ...
Complete actual wiring status first.
No new visible text. No new states. No redesign.
Run build, Playwright screenshot, visible text snapshot, and forbidden text guard.
```

## Rule

Prototype is law. Backend is data. Agent is labor.
