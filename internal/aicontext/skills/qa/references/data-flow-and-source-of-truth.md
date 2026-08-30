# Data Flow And Source-Of-Truth

> This file is part of QA Skill v2. Read when: tracing data flow, DB readback, source-of-truth checks, or verifying field/button data paths.

## Mounted Runtime Map

For each inventoried item, record the runtime map:

- runtime entry point
- user trigger
- owner, app, tenant, or active scope
- persona and role
- shift state
- session/subscription state
- locale and viewport when relevant
- dataset or fixture
- data source, store, API, DB, or read-model
- expected visible actions
- expected blocked actions
- expected next destination

## Field And Button Data Flow

For each relevant field, button, submit, mutation, money-sensitive action, source-of-truth value, or DB-backed result, understand and record when applicable:

- where the data comes from
- which field receives input
- which endpoint, action, tool, or submit path receives it
- what frontend validation runs
- what backend/server validation runs
- which table, read-model, store, cache, or state is read
- which table, read-model, store, cache, or state is written
- where the app redirects or renders afterward
- what the next UI allows the user to do
- how public, account, admin, or user surfaces reflect the change

Use Anvien to map this data flow when useful. Verify it with runtime UI and source-of-truth evidence.

## Source-Of-Truth Data

Treat runtime data as production state, not demo decoration.

Rules:
- Public app data appears only when created through the correct upload/publish/data-entry/admin flow and exists as real runtime DB/source state.
- Do not inject DB rows directly just to make a website display app data unless the scope explicitly defines it as a fixture-only diagnostic.
- If no app is published, public app surfaces must show real empty/unavailable state, not mock cards.
- If no active price/commercial data exists, public/account surfaces must show no-price/not-configured state, not fake price.
- Every visible app name, summary, type, release cue, price, free-use window, billing breakdown, and commercial cue must trace to runtime source evidence when in scope.
- DB-backed write/readback checks need DB/source evidence or an explicit blocker.
- Cleanup should use product/admin actions when the scenario requires user-realistic cleanup. Do not delete directly from DB to hide effects unless environment rules explicitly allow and record it.
