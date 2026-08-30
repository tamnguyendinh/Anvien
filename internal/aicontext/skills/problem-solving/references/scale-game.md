# Scale Game

> use when: uncertain about architectural scalability, when edge cases are ambiguous, or when validating production volume limits.

---

## Overview

Test your proposed approach at extreme scales (1000x larger, 1000x smaller, instantaneous, or over years) to discover what breaks and what fundamentally survives.

**Core Principle:** Extremes expose architectural truths that normal operating scales hide.

---

## Quick Reference

| Scale Dimension | Test At Extremes | What It Reveals |
|---|---|---|
| **Data Volume** | 1 item vs. 1 billion items | Algorithmic bottlenecks, memory ceiling, indexing needs |
| **Execution Speed** | 1 millisecond vs. 1 year | Async requirements, deadlocks, stale state risks |
| **Concurrency** | 1 single user vs. 1 million concurrent users | Race conditions, lock contention, isolation boundaries |
| **Lifecycle Duration** | Minutes vs. decades | Memory leaks, database table bloat, migration feasibility |
| **Failure Frequency** | 0% failure vs. 100% constant failure | Fail-closed guarantees, fallback robustness, recovery paths |

---

## Process

1. **Pick dimension**: Choose a variable dimension (volume, concurrency, duration, latency, failure rate).
2. **Test minimum**: What happens if this is 1000x smaller, faster, or fewer?
3. **Test maximum**: What happens if this is 1000x larger, slower, or more frequent?
4. **Note failure boundaries**: Where do assumptions, memory, or throughput collapse?
5. **Extract resilient cores**: What principles and structures remain valid across all scales?

---

## Real-World Examples

### Example 1: In-Memory State
- **Normal scale:** Holding active sessions in application memory works seamlessly for days.
- **At 1-year scale:** Unbounded memory growth, server restarts wipe user state.
- **Reveals:** Persistence, garbage collection, and explicit session eviction are non-negotiable architectural requirements.

### Example 2: Synchronous IPC / API Calls
- **Normal scale:** Direct synchronous HTTP / RPC calls succeed quickly.
- **At 100,000 requests/sec:** Cascading thread exhaustion and timeouts bring down the whole cluster.
- **Reveals:** Asynchronous event queues and backpressure buffering are survival mandates, not optional optimizations.

---

## Red Flags You Need the Scale Game

- "It works fine in local development" (unverified under real production load).
- System limits and breaking points are unknown.
- "It should scale fine" without quantitative or stress evidence.
- Surprises when encountering burst traffic or large payloads.

---

## Key Invariant

> What works at small scale often catastrophically fails at extreme scale.  
> Stress-testing dimensions mentally early in design prevents expensive rewrites later.
