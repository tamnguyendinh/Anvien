# Source-Of-Truth Rules

> This file is part of QA Skill. Read when: verifying runtime data integrity, DB readback, preventing mock/fake data, or enforcing ground-truth evidence across FE and BE.

## Ground Truth Data Invariants

Treat runtime data as real production state, never as demo decoration.

Rules:
- Public and user-facing data appears only when created through the correct upload/publish/data-entry/API flow and exists as real runtime DB/source state.
- Do not inject DB rows directly just to make an interface or API display data unless the scope explicitly defines it as a fixture-only diagnostic.
- If no entity is published/configured, surfaces and APIs must show real empty/unavailable state, never mock cards or placeholder data.
- If no active commercial data exists, surfaces and responses must show unconfigured/no-price state, never fake prices.
- Every visible or returned value, status, type, release cue, price, and commercial cue must trace to runtime source evidence when in scope.
- DB-backed write/readback checks strictly require DB/source evidence or an explicit blocker.
- Cleanup must use product/admin/API actions when the scenario requires realistic cleanup. Do not delete directly from DB behind the scenes to hide side effects unless environment rules explicitly allow and record it.
