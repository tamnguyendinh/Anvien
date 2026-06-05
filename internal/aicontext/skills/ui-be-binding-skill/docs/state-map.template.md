# Approved UI State Map

This file defines which UI states are approved.

If a state is not listed as approved, it must not be created.

---

## Global Default

Approved states:

- Default state only

Not approved unless explicitly listed:

- Loading state
- Empty state
- Error state
- Maintenance state
- Demo state
- Mock state
- Setup/configuration state
- Coming soon state
- Beta/alpha state
- Development-only state

If backend data is unavailable:

- Preserve approved prototype text.
- Preserve approved prototype layout.
- Do not show technical errors.
- Do not show API status messages.
- Do not show “No data available”.
- Do not show “Failed to load”.
- Do not show “Please configure backend”.

---

## Home Page

Approved states:

- Default

Not approved:

- Loading
- Empty
- Error
- Demo
- API status

Backend may update only approved slots in `ui-slot-map.md`.

---

## Pricing Page

Approved states:

- Default

Backend may update:

- price label
- trial label
- payment URL

Backend may not add:

- discount badge
- billing explanation
- loading message
- failed payment setup message
- beta note

---

## Login Page

Approved states:

- Default

Add approved validation/error states only if they are part of the approved design.

---

## Payment Page

Approved states:

- Default

Add payment-specific approved states here only after user/design approval.
