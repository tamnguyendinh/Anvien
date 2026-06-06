# UI Authority

## Source of Truth

The approved prototype/design is the source of truth for visual output.

Approved sources:

- Prototype HTML:
- Approved screenshots:
- Approved design file:
- Approved route/page:

## Global Rules

- Do not redesign.
- Do not change layout, spacing, colors, typography, section order, header, footer, or brand text.
- Do not add visible text unless explicitly approved.
- Do not rewrite copy unless explicitly approved.
- Do not add helper text, demo labels, MVP labels, placeholder notes, backend/API explanations, or technical UI messages.
- Do not add loading, empty, error, setup, maintenance, beta, or coming-soon states unless approved in the active plan's `YYYY-MM-DD-state-map.md`.
- Do not create nested website UI inside website UI.

## Allowed During UI-BE Binding

- Bind backend values into approved UI slots.
- Preserve approved fallback values when backend data is absent.
- Add API clients, DTOs, adapters, and non-visible code.
- Add tests and guards.

## Brand Lock

Brand name:

```text
AV Cheap Apps
```

Slogan / core text:

```text
Always have what you need
```

Do not change brand wording unless explicitly instructed by the user.

## Header/Footer Lock

Header must remain shared and visually consistent across public pages.

Admin pages may have a different admin shell only if explicitly approved.

## Design Preservation

The rendered page after backend binding must visually match the approved page, except for approved slot values.
