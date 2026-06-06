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
2. Copy the template docs into your project `docs/` folder and rename them:
   - `docs/ui-authority.md`
   - `docs/actual-wiring-status.md`
   - `docs/ui-slot-map.md`
   - `docs/state-map.md`
   - `docs/backend-contract-map.md`
   - `docs/visible-text.lock.json`
3. Fill `docs/actual-wiring-status.md` for the target surface before writing implementation steps.
4. Fill the slot map before backend integration.
5. Add the Playwright tests and adjust routes/selectors to match the app.
6. Give Codex the prompt in `prompts/codex-ui-be-binding-task.md`.
7. Reject any handoff that does not include current-wiring audit, build/test/screenshot/text-lock evidence.

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
