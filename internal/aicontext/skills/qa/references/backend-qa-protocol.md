# Backend QA Protocol (API & Database)

> This file is part of QA Skill. Read when: verifying Backend runtime behavior, API endpoints, database mutations, transaction integrity, or non-UI services. Never open browsers or capture screenshots.

---

### 1. Test Execution (Action Loop)
* **Pre-API Execution:** Query the database to capture the initial baseline state snapshot of the target entity.
* **API Execution:** Send requests covering the full variation matrix (valid happy path, invalid schema, unauthorized/missing auth, boundary limits).
* **Post-API Execution:** Query the database to verify the actual State Diff on disk (confirming correct schema adherence, zero accidental overwrites, and preservation of foreign key integrity).

### 2. Evidence Recording (`evidence.md`)
* Record exact HTTP status codes and structured response payloads.
* Record executed SQL / DB query logs and execution duration (detecting unindexed sequential scans and N+1 query antipatterns).
* Record explicit transaction boundary confirmation logs (`COMMIT` or clean `ROLLBACK`).

### 3. Metric Benchmarking (`benchmark.md`)
* Record response latency distribution: Latency p50, p95 (ms).
* Record memory footprint (Heap / RSS) and active connections within the Connection Pool.
* Record automated test suite pass rates (unit tests, integration tests, contract tests).

### 4. Technical Trap Probes
* **Rollback Verification:** Intentionally trigger a terminal failure at the final step of a multi-step transaction $\rightarrow$ Query the database to prove zero partial writes and zero orphaned records.
* **Race Condition Probe:** Send concurrent requests competing for the exact same resource ID $\rightarrow$ Verify that atomic locks prevent Lost Updates or data constraint violations.
* **Resource Leak Probe:** Send consecutive burst error requests $\rightarrow$ Verify that connection pool clients are cleanly released (preventing pool starvation/leaks) and RAM returns to baseline.

### 5. Escalation to `Edge-Case`
* When the verified scope involves high-contention resources, atomic counters, distributed locks, or complex multi-table ACID transactions:
  * Complete baseline functional verification and record standard PASS.
  * Record an explicit recommendation in the report for CEO to dispatch an `Edge-Case` lane to perform hostile chaos attacks (microsecond timing collisions, out-of-order execution, process crash mid-flight).
