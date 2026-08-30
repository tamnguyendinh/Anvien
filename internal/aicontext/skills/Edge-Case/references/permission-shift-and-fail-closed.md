# Attack Guide: Permission Shifts & Fail-Closed Enforcement

> This file is part of Edge-Case Review Skill v2. Use when: testing permission bypasses, role escalations, active-shift enforcement for financial operations, or fail-closed boundary security.

---

## Attack Surface Overview

Security and authorization bugs occur when guards check permissions too early, when UI buttons are disabled but direct API endpoints remain open, or when critical operational gates (like cashier shifts or MFA challenges) can be bypassed during network blips or unexpected input combinations.

---

## Targeted Invariants

- **Server-Side Fail-Closed**: Every protected backend endpoint must enforce authorization independently of UI disabled states.
- **Active-Shift Gate**: Financial operations (payments, refunds, discounts, cash drawer openings) strictly require an open, validated operator shift.
- **Immediate Role Revocation**: Revoking permissions or locking an account must immediately invalidate active bearer tokens and in-flight operations.
- **Least Privilege Default**: Unrecognized routes, missing tokens, or ambiguous role claims must fail closed (HTTP 401/403) with zero data leakage.

---

## Hostile Perturbations & Attack Scenarios

### 1. Direct API Invocation with Restricted Role
- Find a protected action where the button is hidden/disabled for `Viewer` or `Cashier` roles.
- Use `curl` or Postman to invoke the underlying REST/GraphQL endpoint directly using that restricted role's token.
- *Check:* Does the backend return HTTP 403 Forbidden, or does it execute the action?

### 2. Financial Operation Without Active Shift
- Attempt to execute a payment, refund, or cash payout when:
  - No shift has been opened.
  - The previous shift was closed 1 second prior.
  - The shift is owned by a different cashier/register.
- *Check:* Does the server reject the mutation immediately?

### 3. Mid-Flight Role Revocation
- Open a protected administrative dialog as `Admin`.
- In a separate session, demote the user to `Guest`.
- In the original dialog, click "Save / Delete / Update".
- *Check:* Does the backend block the update and immediately force session invalidation?

### 4. Missing or Corrupted Auth Headers
- Send requests with:
  - Missing `Authorization` header.
  - Header containing `Bearer null`, `Bearer undefined`, `Bearer ""`.
  - Header with expired JWT or JWT signed by an unauthorized key.
- *Check:* Does the system fail closed without panics or leaking stack traces?

---

## Pass / Fail Checklist

- [ ] Zero API endpoints rely purely on frontend UI visibility checks.
- [ ] Financial transactions are 100% blocked unless verified against an active, valid shift.
- [ ] Role changes take effect immediately across all active connections.
- [ ] Missing, malformed, or expired auth tokens fail closed with explicit 401/403 errors.
