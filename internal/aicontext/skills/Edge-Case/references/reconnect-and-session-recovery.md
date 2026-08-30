# Attack Guide: Reconnect & Session Recovery

> This file is part of Edge-Case Review Skill v2. Use when: testing network loss, token refresh, stale session continuity, resubscribe flows, or offline transitions.

---

## Attack Surface Overview

Applications that maintain persistent sessions (WebSocket, SSE, local stores, offline sync) often fail when network drops mid-flight, when tokens expire during disconnection, or when permissions are revoked while the client is disconnected.

---

## Targeted Invariants

- **Fail-Closed on Auth Drift**: Reconnecting clients must fully revalidate tokens and active permissions before resuming actions.
- **No Stale Scope Resubscription**: Clients must not resubscribe to restricted topics/rooms if access was revoked while offline.
- **In-Flight Action Idempotency**: Actions sent just as the network dropped must not be executed twice upon reconnect.
- **Offline Guard**: Offline continuity must never impersonate authoritative live permissions.

---

## Hostile Perturbations & Attack Scenarios

### 1. Drop Connection During In-Flight Mutation
- Dispatch a state-changing POST/WebSocket message and sever the network connection at `T+10ms` before receiving the ACK.
- Reconnect network.
- *Check:* Does the client resend with an idempotency key? Does the server prevent double-application?

### 2. Token Expired / Revoked While Offline
- Authenticate client, capture session token.
- Disconnect client network.
- On backend, revoke token or change user role to unauthorized.
- Reconnect client network.
- *Check:* Does the client fail closed and redirect to login? Or does local cached state allow triggering protected actions?

### 3. Rapid Reconnect Flapping
- Toggle network interface `ON/OFF` 10 times in 5 seconds during active sync.
- *Check:* Do socket handlers leak connections, duplicate listeners, or trigger infinite retry storms?

### 4. Background Sleep / Wakeup with Stale Store
- Freeze/suspend the process for 30 minutes, change server-side state, then unfreeze.
- *Check:* Does the client refresh its store from source of truth before allowing UI interactions?

---

## Pass / Fail Checklist

- [ ] Network loss during mutation results in clean idempotent retry or explicit error.
- [ ] Session restore re-authenticates against backend; revoked tokens force immediate logout.
- [ ] Reconnection storms do not crash the server or exhaust socket descriptors.
- [ ] Stale offline cache never allows bypass of live backend security guards.
