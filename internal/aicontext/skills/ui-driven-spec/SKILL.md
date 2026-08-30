---
name: ui-driven-spec
description: UI-first software development workflow for AI agents. Use when building apps UI-first, extracting specifications from existing prototypes, preparing backend implementation handoff from completed frontend components, or when asked to do "UI-driven development", "prototype-before-spec", "FE before BE", "extract contracts from UI", "slot map", "state map", or "backend contract map".
---

# UI-Driven Spec — FE-First Development Workflow

## Core Philosophy

> **The UI is the user interaction map, not the supreme authority.**  
> **The supreme authority remains the Spec and System Architecture.**  
> **Human / Architect approval is mandatory before a single line of backend code is written.**

- **Traditional flow:** `Spec → Backend → Frontend` (Gaps discovered late, expensive patches).
- **UI-driven flow:** `UI Prototyping → Contract Extraction → Spec Alignment & Gate → Module-by-Module Backend Implementation & QA`.

---

## Iron Laws

1. **Extraction over invention**: Contract extraction reads from FE code and mock types without guessing. Missing information must be flagged as `[TBD]`, never self-filled.
2. **Spec judges artifacts, not the reverse**: If an artifact conflicts with the Spec, conform the artifact to the Spec. Do not alter the Spec during extraction.
3. **Owner Review Gate is non-negotiable**: No backend planning or implementation may begin without explicit owner sign-off. Silence is never approval.
4. **Strict single-module planning**: Never plan multiple backend modules concurrently. Plan exactly one module, implement, verify with runtime QA, obtain approval, then proceed to the next module.
5. **Spec is the supreme law**: If UI/FE behavior conflicts with Spec requirements, stop and escalate — never arbitrarily decide which side wins.
6. **No uncontracted backend work**: Backend never implements anything outside `backend-contract-map.md`. Do not add unsolicited endpoints.
7. **Real UI verification mandatory**: A module is never done without testing with real API integration on the actual UI.

---

## File-to-File Workflow Pipeline

```text
[1] references/ui-prototyping-and-components.md
    │   • Build static HTML prototype (`prototype/`)
    │   • Build FE components with mock types (`src/types/`, `src/mocks/`)
    ▼
[2] references/contract-maps-and-payloads.md
    │   • Extract Data Source Map & API Payload contracts
    │   • Extract UI Slot Map & State Transitions
    │   • Aggregate Backend Contract Map (`docs/backend-contract-map.md`)
    ▼
[3] references/spec-alignment-and-review-gate.md
    │   • Draft / ingest authoritative SPEC (`docs/SPEC.md`)
    │   • Adjudicate and align all contract maps against SPEC
    │   • Package handoff bundle & await Owner Approval
    │
    ▼ ⛔ [MANDATORY OWNER APPROVAL GATE] ⛔
    │
[4] references/module-implementation-and-qa.md
    │   • Plan single module (`Implementation Plan: Module [Name]`)
    │   • Implement backend services & wire live API to FE
    │   • Verify runtime QA on real UI (`docs/qa-log.md`)
    │   • Iterate module-by-module until complete
```

### Pipeline Execution Rules:
- Execute sequentially from `[1]` to `[4]`. Do not skip or invert the file sequence.
- **Stop at `[3] references/spec-alignment-and-review-gate.md`**: The AI agent must stop completely and await explicit owner approval before opening `[4]`.
- **Iterate inside `[4] references/module-implementation-and-qa.md`**: Plan, implement, and QA exactly one module at a time.

---

## Quick Decision Tree

- Building static HTML prototype or mock FE components? → `references/ui-prototyping-and-components.md`
- Documenting API payloads, slot maps, state maps, or backend contracts? → `references/contract-maps-and-payloads.md`
- Drafting initial Spec, aligning artifacts with Spec, or preparing the handoff package? → `references/spec-alignment-and-review-gate.md`
- Planning a single backend module or conducting post-module runtime QA? → `references/module-implementation-and-qa.md`

---

## Reference Index

| Pipeline Stage | Target Reference | Primary Deliverables |
|---|---|---|
| **1. UI Prototyping & Components** | `references/ui-prototyping-and-components.md` | HTML screens, Component hierarchy, TypeScript interfaces, Mock data |
| **2. Contract Maps & Payloads** | `references/contract-maps-and-payloads.md` | `api-payload.md`, `slot-map.md`, `state-map.md`, `backend-contract-map.md` |
| **3. Spec Alignment & Review Gate** | `references/spec-alignment-and-review-gate.md` | `SPEC.md`, `alignment-notes.md`, Handoff package & Owner Sign-off |
| **4. Module Implementation & QA** | `references/module-implementation-and-qa.md` | Single-module plan, Backend code, Real UI Runtime QA (`qa-log.md`) |