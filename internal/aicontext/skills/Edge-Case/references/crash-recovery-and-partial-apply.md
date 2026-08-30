# Attack Guide: Crash Recovery & Partial Apply

> This file is part of Edge-Case Review Skill v2. Use when: testing crashes between lock/write/sync/ack steps, orphaned locks, partial local writes, restart safety, or recovery handlers.

---

## Attack Surface Overview

Complex operations involve multi-step pipelines (e.g. `Acquire Lock -> Write DB -> Update Projection -> Send Sync Event -> Acknowledge -> Release Lock`). If the process crashes or receives `SIGKILL` at any intermediate step, the system must recover deterministically without data corruption or ghost locks.

---

## Targeted Invariants

- **Atomic Transactions**: Multi-step state changes must either completely apply or completely roll back.
- **No Orphaned Locks / State**: Crashing mid-pipeline must not permanently lock resources or leave pending flags stuck.
- **Deterministic Startup Recovery**: On process restart, journal/WAL/reconciliation must repair partial writes.
- **Fail-Closed on Recovery Failure**: If recovery cannot resolve state ambiguity, fail closed and alert.

---

## Hostile Perturbations & Attack Scenarios

### 1. Kill Process Immediately After Local Write (Before Sync/ACK)
- Trigger write transaction.
- Intercept the process immediately after database commit but before writing the sync log or returning HTTP ACK (send `kill -9`).
- Restart the service.
- *Check:* Does the restart reconciliation process detect the un-synced commit and replay/publish the event? Or is the state orphaned?

### 2. Kill Process Between Lock Acquisition and Write
- Acquire distributed lock.
- Send `kill -9` before executing the protected write logic.
- Restart service and attempt write from a different worker.
- *Check:* Does the lock lease clear automatically via TTL or recovery?

### 3. Partial Batch Apply
- Send a batch of 100 items where item #50 triggers a fatal constraint violation or memory panic.
- *Check:* Are items 1–49 rolled back cleanly if atomicity was required? Or does the system leave a half-applied corrupt batch?

### 4. Corrupted WAL / Temp State on Startup
- Inject truncated JSON / corrupted binary data into the local cache/WAL file and launch the service.
- *Check:* Does the service handle corrupted files gracefully (re-fetching from canonical server) or crash in an infinite boot loop?

---

## Pass / Fail Checklist

- [ ] Partial writes roll back cleanly or are reconciled on restart.
- [ ] Process crashes never leave permanent deadlock / orphaned locks.
- [ ] Corrupted local temporary state is detected and quarantined without crashing the daemon.
- [ ] Multi-entity updates remain strictly consistent across restarts.
