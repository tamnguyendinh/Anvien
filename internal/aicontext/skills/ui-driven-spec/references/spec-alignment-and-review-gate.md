# Spec Alignment And Review Gate

> This file is part of UI-Driven Spec Workflow. Read when: drafting/reviewing the authoritative SPEC, aligning extracted contract artifacts with SPEC, or preparing the handoff package for Architect approval.

---

## 1. SPEC Ingestion & Initial Drafting

**Purpose:** Draft or review the SPEC based on the existing UI/FE.
This SPEC serves as the **supreme benchmark to align** all artifacts.
- Draft SPEC according to architect / owner requirements; do not make self-directed assumptions.
- Review and mark items as `[CONFIRMED]` / `[NEEDS_REVIEW]` / `[MISSING]` compared to actual UI/FE.
- Do not override anything in this stage — record deltas only.

**Output:** `docs/SPEC.md` (or an updated version if one already exists).

**Checklist:**
- [ ] Every flow in prototype has a corresponding entry in SPEC
- [ ] Business rules observed from UI are recorded in SPEC
- [ ] Unclear points are marked `[TBD]`, never assumed or self-filled

---

## 2. Authority / Spec Alignment of Extracted Artifacts

**Purpose:** Use SPEC as the supreme standard to **align all extracted contract artifacts**:
`api-contract.md`, `slot-map.md`, `state-map.md`, `backend-contract-map.md`.

The Spec is not rewritten from artifacts — the flow is strictly opposite:
**Spec adjudicates artifacts; artifacts must conform to Spec.**

**Alignment Protocol:**

```text
For each extracted artifact:

  1. Compare each item with SPEC:
     → [OK]               — matches Spec, keep unchanged
     → [CONFLICT]         — conflicts with Spec → modify artifact to match Spec
     → [MISSING_IN_SPEC]  — present in artifact but unmentioned in Spec
                          → escalate, do not self-decide

  2. For [CONFLICT]: explicitly document: "Spec §X.Y states A, artifact currently states B → change to A"

  3. For [MISSING_IN_SPEC]:
     → If clearly direct UI evidence (observed directly from FE) → flag for addition to Spec
     → If uncertain → mark [TBD], include in open questions list for Owner Review

  4. Update artifacts — record version and modification date
```

**Output:** Updated/aligned artifact files + `docs/alignment-notes.md` (recording all resolved conflicts and remaining `[TBD]` items).

**Checklist before proceeding to Owner Review Gate:**
- [ ] Zero unresolved conflicts between Spec and any artifact
- [ ] Every `[MISSING_IN_SPEC]` is explicitly flagged
- [ ] Every `[TBD]` is documented in `alignment-notes.md`
- [ ] Nothing outside the Spec has been self-decided

---

## 3. Handoff To Architect 

**⛔ MANDATORY GATE — AI agent must not self-proceed past this gate without explicit approval.**

**Handoff Package:**
```text
docs/
├── SPEC.md                  ← Current authoritative Spec
├── api-contract.md          ← Aligned with Spec
├── slot-map.md              ← Aligned with Spec
├── state-map.md             ← Aligned with Spec
├── backend-contract-map.md  ← Aligned with Spec
└── alignment-notes.md       ← Resolved conflicts + remaining [TBD] items
```

**AI Agent Handoff Notes:**
- Clearly list all `[TBD]` items requiring approver decisions
- Explicitly list risks or unclear architectural points if any exist
- Propose backend module implementation sequence (derived from `backend-contract-map` dependency graph)

**After Handoff → AI Agent Stops Completely:**
- Do NOT self-proceed
- Do NOT interpret silence as approval
- If rejected → return to the exact designated stage; do NOT rewrite the entire codebase/spec
