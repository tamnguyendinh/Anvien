# Attack Guide: Lock Contention & Concurrent Writes

> This file is part of Edge-Case Review Skill v2. Use when: testing lock acquisition/release races, concurrent write collisions, double-booking, or distributed locking.

---

## Attack Surface Overview

Distributed and multi-threaded systems often fail when two operations attempt to mutate the same resource simultaneously, or when locks are released prematurely or leaked.

---

## Targeted Invariants

- **Lock-Before-Write**: No write mutation may occur without holding the authoritative lock.
- **Mutual Exclusion**: Exactly one writer holds exclusive lock for a given resource/scope.
- **Safe Lease & Expiration**: Locks must time out safely if the holder dies, without allowing double-execution.
- **Deterministic Release**: Locks must be released in `finally` / `defer` blocks across all error paths.

---

## Hostile Perturbations & Attack Scenarios

### 1. High-Frequency Concurrent Race
- Fire `N=50` identical write requests simultaneously with zero jitter.
- *Check:* Does the system accept more than 1 operation? Does inventory or balance drop into negative numbers?

### 2. Lock Acquisition Timeout
- Hold a lock artificially while a secondary client requests it.
- *Check:* Does the secondary client fail gracefully with an explicit timeout or hang indefinitely?

### 3. Crash During Held Lock
- Acquire a lock, start mutation, and forcibly kill the process before release.
- *Check:* Does the lock lease expire cleanly? Can another client acquire the lock after TTL without orphan deadlocks?

### 4. Release After Lease Expiration
- Acquire lock with 5-second TTL, sleep for 7 seconds, then attempt to commit write and release lock.
- *Check:* Does the write fail because the lease was lost? Or does it overwrite subsequent client data?

---

## Repro Command Templates

```bash
# Concurrent request burst (PowerShell example)
1..20 | ForEach-Object -Parallel {
    curl.exe -s -X POST http://localhost:8080/api/v1/resource/claim -H "Content-Type: application/json" -d '{"id":"item-123"}'
}
```

---

## Pass / Fail Checklist

- [ ] Exactly one concurrent operation succeeds; all other parallel attempts receive conflict errors.
- [ ] No phantom reads or split-brain state mutations.
- [ ] Deadlock-free under high contention.
- [ ] Lock TTL expiration correctly invalidates in-flight writes from expired holders.
