# Coverage Model

> This file is part of QA Skill v2. Read when: building inventory before verdict, checking flow steps, or testing state/action matrix.

## Inventory Before Verdict

Before claiming coverage, build an inventory for the declared scope.

Inventory types:
- visible-surface inventory
- interactive inventory
- route/control inventory
- state inventory
- action ledger template
- data/source-of-truth map
- persona/context matrix

Each inventory item needs:
- surface/action/state name
- route or flow
- runtime entry point
- user trigger
- persona/context
- expected result
- data source when relevant
- evidence slot
- verdict slot

Verdicts:
- `Pass`
- `Fail`
- `Blocked`
- `N/A`
- `Out of scope`

Rules:
- `100% coverage` means `100% of declared scope`, not the whole app unless the declared scope is the whole app.
- A route, dialog, component, or action that exists in code but cannot be reached through the real mounted runtime path is not covered.
- A screen is not covered if child dialogs, menus, row actions, forms, or state variants were never inventoried.
- If the ledger is incomplete, QA is incomplete.

## No-Step-Skipped Flow Rule

In each in-scope flow, account for every step:

- open page or app
- initial state
- every tab
- every link
- every button
- every field
- every form
- every submit
- every redirect
- every refresh, back action, or route change
- every locale switch
- every cookie or session change
- every DB or read-model read/write when relevant
- UI after each step
- where the next step leads

If a step exists in scope but was not observed or ledgered, that flow is not fully covered.

## State And Action Matrix

For each inventoried surface, test every applicable state:

- visible / hidden
- enabled / disabled
- loading
- empty
- error
- success
- blocked
- warning
- stale
- selected / unselected
- expanded / collapsed
- pristine / dirty
- validation-error
- submitting
- permission-gated
- shift-gated
- session-gated
- subscription-gated
- stale-context after owner, app scope, role, shift, session, subscription, or locale changes

For every visible action, verify:

- the user can actually reach it
- the action triggers the correct mounted behavior
- the action targets the correct row, item, entity, or selection
- blocked actions explain why they are blocked
- loading, empty, error, success, disabled, and blocked states explain themselves clearly
- displayed data/state updates correctly after the action
- reopen, retry, back, refresh, resubmit, and context switch do not leave stale UI behind

Exploratory clicks are additive. They never replace the required state/action matrix.
