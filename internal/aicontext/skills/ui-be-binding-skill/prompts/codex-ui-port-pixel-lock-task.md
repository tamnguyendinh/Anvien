# Codex Task Prompt: UI Port / Pixel Lock

Use this before UI-BE Binding if the approved HTML prototype has not yet been ported into the TypeScript app.

## Task

Port the approved static HTML prototype into the existing TypeScript frontend with pixel-locked fidelity.

This is not a redesign task. This is a mechanical preservation task.

## Scope

Allowed:

- Convert HTML to TSX/JSX.
- Preserve CSS as closely as possible.
- Extract components only if the rendered output remains identical.
- Add Playwright screenshot and text guards.

Forbidden:

- Do not connect backend.
- Do not add loading, empty, error, demo, sample, MVP, helper, or placeholder UI.
- Do not rewrite copy.
- Do not change layout, spacing, colors, typography, section order, header, footer, or brand text.
- Do not replace the approved UI with a generic component library layout.

## Required Verification

- Full build passes.
- Playwright screenshot comparison passes.
- Visible text snapshot passes.
- Forbidden text guard passes.

## Done Rule

Do not proceed to backend binding until the ported UI is locked by tests.
