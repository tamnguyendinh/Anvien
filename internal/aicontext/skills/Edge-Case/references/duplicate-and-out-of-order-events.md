# Attack Guide: Duplicate & Out-of-Order Events

> This file is part of Edge-Case Review Skill v2. Use when: testing event streams, message queues, async relay, out-of-order logs, replay drift, or deduplication logic.

---

## Attack Surface Overview

In asynchronous and event-driven architectures, networks guarantee at-least-once delivery, not exactly-once. Events can arrive out of chronological order, be replayed repeatedly during retries, or arrive after a newer state update has already been applied.

---

## Targeted Invariants

- **Idempotency**: Processing the same event `N` times produces the exact same outcome as processing it once.
- **Monotonic Version Ordering**: Older event snapshots must never overwrite newer local or persisted state.
- **Deduplication Window**: Duplicate event deliveries within the deduplication window must be silently ignored or safely ACKed without reprocessing.
- **Replay Safety**: Full event replay from offset 0 must reconstruct the exact canonical projection state.

---

## Hostile Perturbations & Attack Scenarios

### 1. Duplicate Event Storm
- Relay the exact same event payload 10 times consecutively with identical event IDs and timestamps.
- *Check:* Does the recipient execute side effects (charges, emails, state increments) 10 times or exactly once?

### 2. Inverted Delivery Sequence (Out-of-Order)
- Given two events: `Event #1 (Status=CREATED, Seq=1)` and `Event #2 (Status=CANCELLED, Seq=2)`.
- Deliver `Event #2` first, then deliver `Event #1` 500ms later.
- *Check:* Does `Event #1` overwrite the status back to `CREATED`? Or does the system recognize sequence/timestamp monotonic ordering and reject the stale event?

### 3. Replay Old Snapshot Over Current State
- Record a state snapshot from 1 hour ago.
- Modify the live system to state `V_NEW`.
- Re-inject the old snapshot `V_OLD`.
- *Check:* Does the system drop `V_OLD` or incorrectly revert live state?

### 4. Duplicate Delivery with Modified Payload (ID Collision)
- Send an event with ID `evt-100` and payload A.
- Send another event with ID `evt-100` but modified payload B.
- *Check:* Does the deduplication layer detect the hash discrepancy or fail closed?

---

## Pass / Fail Checklist

- [ ] Duplicate event delivery causes zero duplicate side effects or balance mutations.
- [ ] Out-of-order event delivery preserves the latest monotonic state.
- [ ] Event deduplication store does not leak memory over long runtimes.
- [ ] Total event log replay produces identical deterministic state to live execution.
