# Frontend Data Flow

> This file is part of QA Skill. Read when: tracing UI input data flow, mounted runtime mapping, or verifying field/button data paths.

## Mounted Runtime Map

For each inventoried UI item, record the runtime map:

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

For each relevant field, button, submit, mutation, or visible result, understand and record when applicable:

- where the data comes from
- which UI field receives input
- which endpoint, action, or submit path receives it
- what frontend client validation runs
- what backend server validation runs
- which store, state, or database table is updated
- where the app redirects or renders afterward
- what the next UI allows the user to do
- how public, account, admin, or user surfaces reflect the change
- verify with runtime UI and source-of-truth evidence
