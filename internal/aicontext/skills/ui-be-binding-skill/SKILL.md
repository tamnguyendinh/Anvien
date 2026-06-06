---
name: ui-be-binding-skill
description: Use when binding backend/API data into an already-approved frontend UI without changing approved layout, copy, visual design, component hierarchy, or interaction states.
---

## Mission

Bind backend data into an already-approved frontend/UI without changing the approved visual output, layout, copy, component hierarchy, interaction contract, or brand presentation.

This skill exists to prevent AI coding agents from turning a polished approved UI into a cheap MVP-looking interface during backend integration.

The approved prototype is the visual authority. Backend integration is only allowed to supply data to approved UI slots.

---

## Core Principle

Backend integration must not become UI redesign.

The agent must treat the frontend as a locked product surface. The backend may provide values, URLs, lists, flags, and actions, but it may not cause new visible UI text, new layout blocks, new explanation panels, new helper notes, new MVP/demo labels, or unapproved states to appear.

---

## Execution Rules

1. Write the UI-BE Binding Plan before coding.
2. Use Anvien as the primary code-intelligence tool when working in an Anvien-indexed repo.
   - Choose the Anvien command that fits the task: query, context, file-detail, impact, route-map, tool-map, shape-check, api impact, graph-health, file-hotspots, detect-changes, or other relevant commands.
   - Minimum required use: refresh/analyze before graph-based work, impact before editing shared/runtime surfaces, and detect-changes before handoff or commit.
3. Code only after the approved UI slots, backend fields, fallback behavior, and forbidden states are mapped.

---

## Authority Order

In this skill, `<active-plan-dir>` means `docs/plans/YYYY-MM-DD-<plan-name>`.
Plan-owned authority artifacts live in that directory and use `YYYY-MM-DD-*` names.

When instructions conflict, follow this order:

1. Direct user instruction for the current task
2. `<active-plan-dir>/YYYY-MM-DD-ui-authority.md`
3. Approved prototype HTML / approved screenshot / approved design source
4. `<active-plan-dir>/YYYY-MM-DD-actual-wiring-status.md`
5. `<active-plan-dir>/YYYY-MM-DD-ui-slot-map.md`
6. `<active-plan-dir>/YYYY-MM-DD-state-map.md`
7. `<active-plan-dir>/YYYY-MM-DD-visible-text.lock.json`
8. Backend API contract / OpenAPI / schema
9. Existing frontend architecture
10. Existing backend architecture
11. Agent implementation judgment

Agent judgment is last. It is not allowed to override UI authority.

---

## Trigger Conditions

Use this skill when the task includes any of these:

- Connect backend to frontend
- Bind API data into UI
- Replace static prototype data with live data
- Convert static UI into data-driven UI
- Wire FE to BE
- Connect pricing, app list, payment, auth, admin data, dashboard data
- Add API client calls to an approved page
- Integrate backend without changing UI

Do not use this skill for original UI design. Use it only after a UI/prototype is already approved.

---

## Non-Goals

This skill is not for:

- Designing new UI
- Improving visual design
- Rewriting marketing copy
- Creating new UX flows
- Adding unapproved empty/loading/error states
- Replacing approved layout with a component-library default
- Creating a generic admin/dashboard UI
- Adding explanatory text to make the UI “clearer”
- Adding badges, notices, alerts, cards, panels, banners, or helper descriptions unless explicitly approved

---

## Mandatory Inputs

Before coding, read and cite internally in the work plan:

- `<active-plan-dir>/YYYY-MM-DD-ui-authority.md`
- `<active-plan-dir>/YYYY-MM-DD-actual-wiring-status.md`
- `<active-plan-dir>/YYYY-MM-DD-ui-slot-map.md`
- `<active-plan-dir>/YYYY-MM-DD-state-map.md`
- `<active-plan-dir>/YYYY-MM-DD-visible-text.lock.json`
- `<active-plan-dir>/YYYY-MM-DD-backend-contract-map.md` or backend API contract/schema, if available
- Existing frontend page/component files
- Existing API client/service layer
- Existing Playwright tests

If one of these project files is missing, create a minimal version in `<active-plan-dir>` from this skill pack's `docs/*.template.md` files before implementation, unless the current scope explicitly forbids adding docs.

---

## Skill Pack Assets

Resolve these files relative to this `SKILL.md`. Load only the asset needed for the current step.

- `checklists/pre-coding-checklist.md`: use before writing the binding plan or editing code.
- `checklists/final-verification-checklist.md`: use before handoff.
- `docs/ui-authority.template.md`: create missing `<active-plan-dir>/YYYY-MM-DD-ui-authority.md`.
- `docs/actual-wiring-status.template.md`: create missing `<active-plan-dir>/YYYY-MM-DD-actual-wiring-status.md`; complete target surface audit before implementation.
- `docs/ui-slot-map.template.md`: create missing `<active-plan-dir>/YYYY-MM-DD-ui-slot-map.md`.
- `docs/state-map.template.md`: create missing `<active-plan-dir>/YYYY-MM-DD-state-map.md`.
- `docs/backend-contract-map.template.md`: create missing `<active-plan-dir>/YYYY-MM-DD-backend-contract-map.md`.
- `docs/visible-text.lock.example.json`: create missing `<active-plan-dir>/YYYY-MM-DD-visible-text.lock.json`.
- `tests/playwright/approved-ui.guard.spec.ts`: template for screenshot, visible-text, forbidden-text, and shell-preservation guards.
- `tests/playwright/slot-binding.guard.spec.ts`: template for approved-slot binding guards.
- `reports/ui-be-binding-report.template.md`: final report template.
- `prompts/*.md`: use only when preparing a task prompt for another agent/session.

---

## Hard Forbidden Actions

The agent must not:

- Redesign layout
- Change spacing
- Change colors
- Change typography
- Change section order
- Change header/footer structure
- Change approved brand name or slogan
- Rewrite visible copy outside approved slots
- Add new visible text outside approved slots
- Add UI comments, explanatory notes, or technical backend messages
- Add MVP/demo/sample/placeholder labels
- Add loading state unless approved
- Add empty state unless approved
- Add error state unless approved
- Add “No data available” unless approved
- Add “Failed to load” unless approved
- Add “Coming soon” unless approved
- Add “Example”, “Sample”, “Mock data”, “For testing only”, “Backend connected”, or similar visible text
- Add new cards, banners, alerts, badges, panels, sidebars, tables, or modals unless approved
- Replace static approved UI with a generic generated UI
- Introduce nested website UI inside website UI
- Change the approved route/page shell unless required by architecture and visually identical

---

## Allowed Actions

The agent may:

- Create or update API client functions
- Create DTO types, response schemas, and validators
- Create adapter functions from backend DTO to UI view model
- Bind backend fields into approved UI slots
- Preserve approved fallback text when data is missing
- Add non-visible code comments
- Add tests for API mapping and UI binding
- Add Playwright guards for screenshot, visible text, forbidden text, and approved slots
- Refactor only when rendered UI remains identical

---

## Binding Rule

Backend data may only replace values inside approved slots listed in `<active-plan-dir>/YYYY-MM-DD-ui-slot-map.md`.

Allowed:

```tsx
<span>{viewModel.priceLabel}</span>
```

Forbidden:

```tsx
<div>
  <span>{viewModel.priceLabel}</span>
  <p>This price is loaded from backend API.</p>
</div>
```

Forbidden:

```tsx
{isLoading ? <p>Loading pricing...</p> : <PriceCard />}
```

unless `Loading pricing...` and the loading state are explicitly approved in `<active-plan-dir>/YYYY-MM-DD-state-map.md`.

---

## Missing Data Rule

If backend data is missing, malformed, late, or unavailable:

1. Preserve approved prototype text or approved fallback from `<active-plan-dir>/YYYY-MM-DD-ui-slot-map.md`.
2. Do not show technical errors.
3. Do not show API status messages.
4. Do not create a new empty state.
5. Do not add explanation text.
6. Log internally if needed, but do not expose debug state to the UI.

---

## State Rule

Use `<active-plan-dir>/YYYY-MM-DD-state-map.md`.

If a state is not explicitly approved, it does not exist.

Default assumption:

- Loading state: not approved
- Empty state: not approved
- Error state: not approved
- Maintenance state: not approved
- Demo state: not approved
- Setup/configuration state: not approved

---

## Implementation Protocol

The agent must follow this order:

1. Run `checklists/pre-coding-checklist.md`.
2. Read authority files.
3. Create missing authority docs in `<active-plan-dir>` from this skill pack's `docs/*.template.md`.
4. Complete `<active-plan-dir>/YYYY-MM-DD-actual-wiring-status.md` for the target surface before writing implementation steps.
5. Classify the target surface as `bound-correct`, `partial`, `wrong-fake-data`, `unbound`, `no-real-data`, or `blocked`.
6. Rewrite the implementation phase from the audit result: preserve-only, bind missing slots only, remove fake display, implement source binding, render approved truthful state, or block.
7. Identify approved pages in scope.
8. Identify approved UI slots.
9. Identify backend fields required for those slots.
10. Create a mapping table before coding.
11. Reuse existing API/client architecture when available.
12. Add or update backend DTO types.
13. Add adapter from backend DTO to UI view model.
14. Bind view model into existing approved UI slots only.
15. Preserve fallback text from `<active-plan-dir>/YYYY-MM-DD-ui-slot-map.md`.
16. Do not alter JSX/HTML structure unless technically necessary.
17. Do not alter CSS unless technically necessary.
18. Run full build before tests.
19. Run unit/integration tests.
20. Add or adapt guard tests from `tests/playwright/*.spec.ts` when the project lacks equivalent guards.
21. Run Playwright behavior tests.
22. Run screenshot comparison.
23. Run visible text snapshot.
24. Run forbidden text guard.
25. Run `checklists/final-verification-checklist.md`.
26. Produce final report with `reports/ui-be-binding-report.template.md`.

---

## Required Pre-Coding Plan Format

Before coding, produce a short plan with this structure:

```md
## UI-BE Binding Plan

### Scope
- Active plan dir:
- Page(s):
- Backend endpoint(s):
- Approved slot(s):

### Authority Files Read
- <active-plan-dir>/YYYY-MM-DD-ui-authority.md
- <active-plan-dir>/YYYY-MM-DD-actual-wiring-status.md
- <active-plan-dir>/YYYY-MM-DD-ui-slot-map.md
- <active-plan-dir>/YYYY-MM-DD-state-map.md
- <active-plan-dir>/YYYY-MM-DD-visible-text.lock.json

### Actual Wiring Status
- Status file:
- Target surface row:
- Classification:
- Real production source:
- Allowed next action:
- Preserve-only files/symbols:
- Fake/prototype/demo leakage findings:

### Binding Table
| UI Slot | Existing Approved Text | Backend Source | Fallback | UI Change Allowed |
|---|---|---|---|---|

### Forbidden During This Task
- No redesign
- No new visible text
- No unapproved state
- No helper/demo/MVP/placeholder labels

### Verification
- Full build
- Unit/integration tests
- Playwright screenshot
- Visible text snapshot
- Forbidden text guard
```

Do not start coding until the table is complete.

---

## Required Final Report Format

After implementation, produce:

```md
## UI-BE Binding Report

### Summary

### Actual Wiring Status
- Status file:
- Target surface classification:
- Evidence IDs:
- Decision:

### Files Changed
- 

### Backend Endpoints Used
- 

### UI Slots Bound
| UI Slot | Backend Field | Fallback Preserved | Notes |
|---|---|---|---|

### UI Preservation Evidence
- Screenshot comparison:
- Visible text snapshot:
- Forbidden text guard:
- Header/footer unchanged:
- No unapproved states added:

### Commands Run
- Build:
- Tests:
- Playwright:

### Result
DONE / NOT DONE

### Remaining Risks
- 
```

A scope is not DONE without evidence.

---

## Playwright Guard Requirements

Every page touched by this skill should have these guard layers:

1. Screenshot comparison
2. Visible text snapshot
3. Forbidden visible text check
4. Approved slot check
5. Header/footer preservation check when applicable

If the project does not yet have these tests, add them before or during the binding task.

---

## Forbidden Text Default List

Use this default list unless project authority provides a stricter list:

- MVP
- Demo
- Sample
- Placeholder
- Mock data
- Coming soon
- TODO
- No data available
- Failed to load
- Loading...
- Backend
- API
- For testing only
- This is a demo
- Please configure
- Example
- Test data
- Development only
- Work in progress
- Under construction
- Beta
- Alpha

A forbidden term may appear only if explicitly approved in `<active-plan-dir>/YYYY-MM-DD-visible-text.lock.json` or `<active-plan-dir>/YYYY-MM-DD-ui-slot-map.md`.

---

## Backend Integration Pattern

Preferred architecture:

```text
Backend DTO → Adapter/Mapper → UI ViewModel → Approved UI Slot
```

Avoid:

```text
Backend DTO → Direct UI rendering everywhere
```

Reason: direct DTO rendering causes field leakage and agent-invented UI.

---

## Adapter Requirements

Adapters must:

- Return only fields needed by approved UI slots
- Provide approved fallbacks
- Sanitize/null-check backend data
- Avoid exposing backend/internal field names directly to UI
- Avoid returning debug labels or technical messages

Example:

```ts
export type PricingViewModel = {
  priceLabel: string;
  trialLabel: string;
  paymentUrl: string;
};

export function mapPricingDtoToViewModel(dto: PricingDto | null): PricingViewModel {
  return {
    priceLabel: dto?.monthlyPriceLabel || "$3/month",
    trialLabel: dto?.trialLabel || "30-day free trial",
    paymentUrl: dto?.paymentUrl || "/pricing"
  };
}
```

---

## Emergency Stop Conditions

Stop and report instead of improvising if:

- A backend field does not map to any approved UI slot
- A required UI state is missing from `<active-plan-dir>/YYYY-MM-DD-state-map.md`
- The backend contract conflicts with approved copy
- The implementation requires adding visible text not in `<active-plan-dir>/YYYY-MM-DD-visible-text.lock.json`
- Screenshot diff becomes large after a binding change
- The agent cannot preserve layout while integrating backend

Do not “solve” these by redesigning.

---

## Done Definition

The task is DONE only when:

- Backend data is bound only to approved slots
- Actual wiring status is complete for the target surface
- No approved layout/copy/visual structure is changed outside slots
- Full build passes
- Tests pass
- Playwright screenshot comparison passes or diff is explicitly approved
- Visible text snapshot passes
- Forbidden text guard passes
- Final report includes evidence
