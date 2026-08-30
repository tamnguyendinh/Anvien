# Attack Matrix & Extreme Chaos Engineering

> This file is part of Edge-Case Review Skill v2. Read when: designing hostile perturbation matrices, simulating human/system chaos, and building compounded failure scenarios.

---

## The Chaos Mandate

Do not stop at ordinary happy-path checks or simple user mistakes. The review bar is:
> **"Can this system state exist under real timing, network retry, crash, reconnect, stale-cache, or hostile-operator conditions?"**

---

## Human Chaos vs. System Chaos

Treat `human-chaos` and `system-chaos` as two distinct attack surfaces:

### 1. Human Chaos
- Rapid double-clicks / repeated fast submit spam.
- Reopen, refresh, and back-navigation loops.
- Multi-window / multi-tab simultaneous actions.
- Cross-role actions (unauthorized user attempting protected controls).
- Operations attempted in uninitialized or closed states (e.g. without an active shift).

### 2. System Chaos
- Reconnect during in-flight write, acknowledgment, or lock transition.
- Duplicate event relay after partial network success.
- Out-of-order event replay after restart or resubscribe.
- Stale store, stale permission cache, or stale shift state.
- Process crash between lock acquire, local write, projection apply, sync push, and lock release.
- Recovery from partial apply, partial sync, or corrupt local cache.

---

## Mandatory Perturbation Dimensions

For every runnable scope, the perturbation matrix must include at least:
1. **Timing perturbation**: Race condition, slow network delay, fast double-submit.
2. **Ordering perturbation**: Out-of-order message delivery, inverted event sequence.
3. **Stale-state perturbation**: Stale local cache, expired token, revoked role while dialog is open.
4. **Isolation / Auth perturbation**: Cross-tenant ID injection, missing authorization token.
5. **Crash / Recovery perturbation**: Process killed mid-write, lock holder crashed.

---

## Compounded Attack Scenarios

Prefer compounded scenarios over isolated single-input fuzzing:

| Compounded Scenario | Combined Perturbations | Invariant Tested |
|---|---|---|
| **Stale Store + Network Retry** | Reconnect during retry while client store holds outdated snapshot | Idempotency & state freshness |
| **Partial Apply + Duplicate Relay** | Duplicate event delivered after local write but before ACK persistence | Deduplication & replay safety |
| **Role Revocation + Open Modal** | User role revoked on backend while protected modal stays mounted | Fail-closed permission check |
| **Shift Closure + Payment Confirm** | POS shift closed by manager right as cashier confirms payment | Active-shift financial guard |
| **Orphaned Lock + Immediate Retry** | Lock holder crashes before unlock; secondary client attempts write | Lock lease expiration & fail-safe release |

---

## Pass / Fail Criteria Under Chaos

### Pass Criteria:
- System **fails closed** on all permission, shift, and lock boundaries.
- **Idempotent recovery** under duplicate or replayed events.
- **Zero duplicate financial movement** or unauthorized state mutation.
- **Zero cross-`owner_id` or cross-scope data leakage**.
- No silent partial success leaving corrupted state armed for subsequent actions.

### Fail Criteria:
- Any fail-open behavior, data leakage, duplicate write, or unhandled panic.
