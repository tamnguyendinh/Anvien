# Actual Wiring Status: {{PROJECT_OR_SURFACE_NAME}} UI-BE Binding

## Metadata

| Field | Value |
|---|---|
| Title | {{PROJECT_OR_SURFACE_NAME}} UI-BE Binding Actual Wiring Status |
| Date | {{YYYY-MM-DD}} |
| Slug | {{slug-name}} |
| Status | Draft / P0 Audit Required / P0 Complete / Blocked |
| Owner | {{owner-or-agent-name}} |
| Companion plan | {{path/to/plan.md}} |
| Companion UI authority | {{active-plan-dir}}/{{YYYY-MM-DD}}-ui-authority.md |
| Companion slot map | {{active-plan-dir}}/{{YYYY-MM-DD}}-ui-slot-map.md |
| Companion state map | {{active-plan-dir}}/{{YYYY-MM-DD}}-state-map.md |
| Companion backend contract map | {{active-plan-dir}}/{{YYYY-MM-DD}}-backend-contract-map.md |
| Companion visible text lock | {{active-plan-dir}}/{{YYYY-MM-DD}}-visible-text.lock.json |

---

## Purpose

This artifact is the required preparation output before any UI-BE implementation slice.

It records the real current state of the codebase:

- which UI surfaces already read real backend truth;
- which surfaces are partially bound;
- which surfaces are still prototype/static;
- which surfaces incorrectly use sample/fake/demo data as production display;
- which surfaces have no real data and must render an approved truthful empty/not-configured/unavailable/denied state.

No UI-BE implementation phase may start until the target surface has a completed row in this file.

The next implementation phase must be written from this evidence. Do not guess code changes before this matrix is filled for the target surface.

---

## Authority Order

When this audit is used for implementation, authority order is:

1. User instruction
2. Project authority / repository rules
3. UI authority document
4. Approved prototype / approved screenshots
5. This actual wiring status file
6. UI slot map
7. State map
8. Backend contract map
9. Visible text lock
10. Existing architecture
11. Agent implementation judgment

If this file says a surface is already `bound-correct`, implementation must preserve it unless a higher authority explicitly says otherwise.

---

## Classification Rules

| Status | Meaning | Allowed next action |
|---|---|---|
| `bound-correct` | UI already reads the correct DB/admin-config/read-model/source and renders approved states when source is empty. | Do not rewrite. Add missing evidence/tests only. |
| `partial` | Some approved slots read real sources, but other approved slots are still static/prototype/fallback or missing provenance. | Bind only the missing slots. Preserve correct existing binding. |
| `wrong-fake-data` | UI renders invented/prototype/sample/demo data as production truth when real source is empty or absent. | Remove fake production display and bind to real source or approved empty/not-configured/unavailable/denied state. |
| `unbound` | UI is still static/prototype-only for dynamic slots. | Implement DTO/read-model/adapter/view-model binding from the real source. |
| `no-real-data` | The correct source currently has no rows/config/artifacts/invoices/entitlements. | Do not seed or invent data. Render the approved truthful empty/not-configured/unavailable/denied state. |
| `blocked` | Source, authority, slot ownership, or state contract is unclear. | Stop and record blocker. Do not redesign or fabricate data. |

---

## Source Truth Rules

Only these sources may be treated as production truth for visible UI display.

| Domain | Only valid production source for UI display |
|---|---|
| App catalog | Admin-created and published app catalog state. |
| Public content | Website-managed public content/admin content configuration. |
| Commercial/pricing | Admin commercial configuration and commercial read models. |
| Release/download | Admin-uploaded/published release artifacts and release distribution read models. |
| Billing/invoice | Billing truth, invoice read models, provider reconciliation-safe state. |
| Entitlement/usage | Backend entitlement, access-state, retained usage, metering, and tuple-safe authority. |
| Account | Account DB/read model and portal session state. |
| Admin | Permission-gated admin read models, mutation results, audit/support evidence. |
| Locale/text | Approved i18n catalogs, route locale helpers, visible text lock. |
| Media/assets | Approved uploaded media, configured media references, or approved fallback asset only. |
| Auth/session | Auth backend/session/account truth and approved auth notices. |

Prototype HTML sample values such as invented app names, sample prices, sample releases, sample invoices, sample entitlements, sample metrics, demo rows, fixture labels, or placeholder media are visual references only.

They are never production truth.

---

## Audit Matrix

Fill one row per UI surface before implementation.

| Surface | Current Code State | Real Source | Current UI Behavior | Correct? | Action | Evidence | Next Phase Decision |
|---|---|---|---|---|---|---|---|
| {{Surface name}} | {{What files/components/readers currently do}} | {{DB/read model/API/admin config/source of truth}} | {{What the user currently sees}} | `bound-correct` / `partial` / `wrong-fake-data` / `unbound` / `no-real-data` / `blocked` | {{Preserve / bind missing slots / remove fake display / implement source binding / render approved state / block}} | {{Evidence IDs}} | {{P1-A / preserve-only / exact partial-binding task / blocked}} |
| {{Surface name}} |  |  |  |  |  |  |  |

---

## Required Surface Audit Form

Use this form when auditing each surface before adding it to the matrix.

### Surface

`{{surface-name}}`

### Route / Entry Points

- Route/page: `{{path}}`
- Component(s): `{{path}}`
- Server/read model/API reader: `{{path}}`
- Adapter/view-model: `{{path}}`
- Tests: `{{path}}`

### Approved UI Authority

- UI authority file: `{{path}}`
- Approved prototype/screenshot: `{{path}}`
- Slot map section: `{{section}}`
- State map section: `{{section}}`
- Visible text lock section: `{{section}}`

### Current Binding Status

| Slot / Visible dynamic value | Current source | Valid production source? | Status | Required action |
|---|---|---|---|---|
| `{{slot.name}}` | `{{current source}}` | Yes / No / Unclear | `bound-correct` / `partial` / `wrong-fake-data` / `unbound` / `blocked` | {{action}} |

### Current State Handling

| State | Current behavior | Approved? | Required action |
|---|---|---|---|
| Default | {{behavior}} | Yes / No / Unclear | {{action}} |
| Loading | {{behavior}} | Yes / No / Unclear | {{action}} |
| Empty | {{behavior}} | Yes / No / Unclear | {{action}} |
| Error | {{behavior}} | Yes / No / Unclear | {{action}} |
| Denied / unauthenticated | {{behavior}} | Yes / No / Unclear | {{action}} |
| Not configured / unavailable | {{behavior}} | Yes / No / Unclear | {{action}} |

### Fake / Prototype / Demo Leakage Search

Record exact search commands and results.

```bash
{{rg command for demo/sample/prototype/fallback/mock strings}}
{{rg command for hardcoded visible text}}
{{rg command for fixture-only helpers}}
{{rg command for fake rows/buildDemo/buildSample/buildPrototype}}
```

Findings:

- `{{file:path:line}}` — {{finding}}
- `{{file:path:line}}` — {{finding}}

### Production Source Trace

Trace each approved dynamic value back to the real source.

| UI value | View model field | Adapter field | Backend/API/read-model field | DB/admin/source truth | Evidence |
|---|---|---|---|---|---|
| `{{visible value / slot}}` | `{{viewModel.field}}` | `{{adapter.field}}` | `{{backend.field}}` | `{{source truth}}` | `{{evidence ID}}` |

### Surface Classification

Final classification:

`{{bound-correct | partial | wrong-fake-data | unbound | no-real-data | blocked}}`

Reason:

```text
{{Short factual reason. Do not speculate.}}
```

Allowed next action:

```text
{{Exact allowed next action.}}
```

Forbidden next action:

```text
{{Exact forbidden next action, especially rewrite/redesign/fake data/state creation.}}
```

---

## Audit Evidence Log

Every matrix row must cite at least one evidence ID.

| ID | Source | Evidence | Notes |
|---|---|---|---|
| `P0-ANV-001` | `{{anvien command / code search / file-context}}` | {{What this proved}} | {{Notes}} |
| `P0-RAW-001` | `{{manual source inspection path/range}}` | {{What this proved}} | {{Notes}} |
| `P0-RG-001` | `{{rg/search command}}` | {{What this found or ruled out}} | {{Notes}} |
| `P0-TEST-001` | `{{test command / current test}}` | {{What this proved}} | {{Notes}} |

---

## Required Audit Procedure Per Surface

Before any implementation phase is written or executed:

1. Run project codebase analysis/indexing if required by project rules.
2. Use the approved codebase analysis tool to identify exact route/page/component/read-model/test files.
3. Inspect the exact files and source ranges relevant to the target surface.
4. Read the companion UI authority, slot map, state map, backend contract map, and visible text lock.
5. Trace every visible dynamic value back to a valid production source.
6. Search for sample/prototype/demo/mock/fallback/hardcoded visible values in relevant files.
7. Verify permission, auth, tuple, audit, and fail-closed behavior where applicable.
8. Classify the surface using this file's status rules.
9. Write the next action narrowly:
   - preserve-only;
   - bind missing slots only;
   - remove fake production display;
   - implement source binding;
   - render approved truthful state;
   - or block.
10. Only after this row is complete may implementation begin.

---

## Non-Negotiable Audit Outcomes

- If the real source has zero rows/config/artifacts, the correct output is an approved empty/not-configured/unavailable/denied state, not invented display data.
- If a current surface is already correctly bound, do not rewrite it.
- If a current surface is partially correct, preserve the correct parts and bind only missing slots.
- If a current surface uses prototype sample data as production truth, the next phase must remove that behavior.
- If the only evidence is visual prototype HTML, that is not enough to mark production binding correct.
- If a backend field exists but has no approved UI slot, do not render it.
- If a UI state is not approved in the state map, do not create it.
- If source authority conflicts are unclear, classify as `blocked` and stop.
- Do not add seed/demo/sample data to make the UI look populated.
- Do not add visible technical notes such as API status, backend source labels, fixture labels, or debug text.

---

## Downstream Phase Rewrite Rules

After this audit is filled, update the main plan before implementation.

| Audit status | How downstream phase must change |
|---|---|
| `bound-correct` | Rewrite the matching phase as preserve-only. Remove code-edit steps. Keep only missing evidence, test, or guard work. |
| `partial` | Rewrite the matching phase to list only missing approved slots. Explicitly name already-correct files/symbols that must not be rewritten. |
| `wrong-fake-data` | Rewrite the matching phase to remove fake/prototype production display and replace it with valid source binding or approved truthful state. |
| `unbound` | Rewrite the matching phase with exact route/page files, read models, adapter/view-model fields, tests, and impact targets. |
| `no-real-data` | Rewrite the matching phase so the implementation target is the approved empty/not-configured/unavailable/denied state. Do not add seed data. |
| `blocked` | Rewrite the matching phase as blocked with the exact missing authority/source evidence. Do not schedule code edits. |

Any later implementation phase is conditional until this rewrite is complete.

---

## Implementation Gate

Implementation may start only when all required fields below are checked.

- [ ] Target surface exists in the Audit Matrix.
- [ ] Target surface has a classification.
- [ ] Target surface has evidence IDs.
- [ ] Real production source is identified.
- [ ] Approved UI slots are identified.
- [ ] Approved states are identified.
- [ ] Missing slots are listed, if any.
- [ ] Already-correct files/symbols are listed as preserve-only, if any.
- [ ] Fake/prototype/demo leakage findings are recorded, if any.
- [ ] Blockers are recorded, if any.
- [ ] Main plan has been rewritten according to the audit status.

---

## Preserve-Only Declaration

Use this when a surface is already `bound-correct`.

```text
Surface: {{surface-name}}
Status: bound-correct

Preserve-only files/symbols:
- {{path/symbol}}
- {{path/symbol}}

Allowed work:
- Add missing evidence.
- Add missing tests.
- Add guard tests.
- Improve test coverage without changing production behavior.

Forbidden work:
- Do not rewrite the reader.
- Do not rewrite the adapter.
- Do not rewrite the component.
- Do not change layout/copy/state.
- Do not change data source.
- Do not add fallback/demo/sample values.
```

---

## Partial-Binding Declaration

Use this when a surface is `partial`.

```text
Surface: {{surface-name}}
Status: partial

Already correct. Preserve:
- {{slot/source/file/symbol}}

Missing approved binding:
- {{slot}} must bind from {{real source}} through {{adapter/view-model}}.

Forbidden:
- Do not rewrite already-correct slots.
- Do not add new slots.
- Do not add unapproved state.
- Do not change visible copy outside approved slots.
- Do not add sample/demo/prototype display.
```

---

## Wrong-Fake-Data Declaration

Use this when a surface is `wrong-fake-data`.

```text
Surface: {{surface-name}}
Status: wrong-fake-data

Fake/prototype/demo production display found:
- {{path:line}} — {{fake value / builder / fixture / fallback}}

Required correction:
- Remove fake production display.
- Bind to real source if approved slot exists.
- Otherwise render approved truthful empty/not-configured/unavailable/denied state.

Forbidden:
- Do not replace fake data with different fake data.
- Do not seed data.
- Do not hide fake data behind another fallback.
- Do not create unapproved UI state.
```

---

## Blocked Declaration

Use this when a surface is `blocked`.

```text
Surface: {{surface-name}}
Status: blocked

Blocker:
- {{exact missing authority/source/contract/evidence}}

Required next step:
- {{who/what must provide the missing decision}}

Forbidden:
- Do not code.
- Do not redesign.
- Do not infer source authority.
- Do not fabricate data.
```

---

## P0 Summary

| Classification | Count | Surfaces |
|---|---:|---|
| `bound-correct` | 0 |  |
| `partial` | 0 |  |
| `wrong-fake-data` | 0 |  |
| `unbound` | 0 |  |
| `no-real-data` | 0 |  |
| `blocked` | 0 |  |

---

## Final P0 Decision

Choose one:

- [ ] P0 audit incomplete. Implementation is blocked.
- [ ] P0 audit complete. Main implementation plan must be rewritten from the matrix before coding.
- [ ] P0 audit complete. Target surface is preserve-only.
- [ ] P0 audit complete. Target surface has exact partial-binding work.
- [ ] P0 audit complete. Target surface requires fake-data removal.
- [ ] P0 audit complete. Target surface is blocked by missing authority/source evidence.

Decision note:

```text
{{final decision note}}
```
