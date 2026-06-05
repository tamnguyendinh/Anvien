# Codex Task Prompt: UI-BE Binding

Use the UI-BE Binding Skill.

## Task

Bind backend data into the approved frontend UI without visual, textual, layout, or interaction drift.

## Authority Files

Read before coding:

- `docs/ui-authority.md`
- `docs/ui-slot-map.md`
- `docs/state-map.md`
- `docs/visible-text.lock.json`
- Backend API contract/schema
- Existing frontend architecture
- Existing API client/service layer
- Existing Playwright tests

## Scope

Allowed:

- Add/update API client functions.
- Add/update DTO types.
- Add adapter/mapping layer from backend DTO to UI view model.
- Bind backend data only into approved UI slots.
- Preserve approved fallbacks.
- Add/update tests.

Forbidden:

- Do not redesign.
- Do not change layout, spacing, colors, typography, section order, header, footer, or brand text.
- Do not add visible text outside approved slots.
- Do not rewrite approved copy.
- Do not add helper text, demo labels, MVP labels, placeholder notes, backend/API explanations, or technical UI messages.
- Do not add loading, empty, error, setup, maintenance, beta, or coming-soon states unless approved in `state-map.md`.
- Do not create nested website UI inside website UI.

## Required Plan Before Coding

Create this plan first:

```md
## UI-BE Binding Plan

### Scope
- Page(s):
- Backend endpoint(s):
- Approved slot(s):

### Authority Files Read
- 

### Binding Table
| UI Slot | Existing Approved Text | Backend Source | Fallback | UI Change Allowed |
|---|---|---|---|---|

### Implementation Steps
1.
2.
3.

### Verification
- Full build:
- Unit/integration tests:
- Playwright screenshot:
- Visible text snapshot:
- Forbidden text guard:
```

Do not start coding until the binding table is complete.

## Implementation Order

1. Read authority files.
2. Identify approved slots.
3. Identify backend fields.
4. Create DTO/view model mapping.
5. Bind only into approved slots.
6. Preserve fallback text.
7. Run full build.
8. Run tests.
9. Run Playwright guards.
10. Produce final report.

## Done Criteria

DONE only when:

- Full build passes.
- Tests pass.
- Playwright screenshot comparison passes or approved diff is documented.
- Visible text snapshot passes.
- Forbidden text guard passes.
- No unapproved visible text was added.
- No unapproved UI state was added.
- Final report includes evidence.
