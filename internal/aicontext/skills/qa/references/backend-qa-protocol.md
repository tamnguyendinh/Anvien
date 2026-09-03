# Backend QA Protocol (API & Database)

> This file is part of QA Skill. Read when: verifying Backend runtime behavior, API endpoints, database mutations, transaction integrity, concurrency safety, or non-UI services. Never open browsers or capture screenshots.

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

### 4. Database & Transaction Traps
* **Rollback Verification:** Intentionally trigger a terminal failure at the final step of a multi-step transaction $\rightarrow$ Query the database to prove zero partial writes and zero orphaned records.
* **Race Condition Probe:** Send concurrent requests competing for the exact same resource ID $\rightarrow$ Verify that database atomic locks prevent Lost Updates or constraint violations.
* **Resource Leak Probe:** Send consecutive burst error requests $\rightarrow$ Verify that connection pool clients are cleanly released (preventing pool starvation/leaks) and RAM returns to baseline.

### 5. Concurrency, Threads & Shared Memory Safety
* **Scope Matrix by Runtime:**
  - **Go:** Goroutines, background channels, shared in-memory maps/slices.
  - **Rust:** OS threads (`std::thread`), Tokio async tasks, shared `Arc<Mutex<T>>`.
  - **Java / C#:** Platform Threads, Virtual Threads (Project Loom), Tasks, `ExecutorService` thread pools.
  - **Node.js:** Worker Threads (`worker_threads`), background cluster workers.
  - **Python / C++:** Threads, Pthreads, multiprocessing workers.
  - *(And equivalent concurrency primitives in other languages).*
* **Mandatory Race Instrumentation:**
  - **Go:** Mandatory `-race` flag (`go test -race` or runtime binary built with `-race`).
  - **Rust / C / C++:** Mandatory ThreadSanitizer (`cargo-tsan` or `-fsanitize=thread`).
  - **Java / C# / Node.js:** Execute stress-test concurrency probes and analyze thread dumps / unhandled thread exceptions.
* **Trap Probes:**
  - **Shared Memory Data Race:** Fire concurrent requests (10–20 parallel workers) against shared in-memory state $\rightarrow$ Immediate FAIL if any data race warning occurs (e.g. Go: `WARNING: DATA RACE`, `concurrent map writes`; Java: race-induced `ConcurrentModificationException`; C++/Rust: TSan race reports).
  - **Thread / Coroutine / Worker Leak:** Measure active concurrent units before and after execution load:
    - Go: `runtime.NumGoroutine()`
    - Java / C#: Active Thread Count / ThreadPool queue
    - Node.js: Active Worker Thread count
    $\rightarrow$ Verify active count cleanly returns to baseline; flag critical leak if workers hang indefinitely on deadlocks, unreleased locks, or blocked channels.
* **Required Evidence:** Runtime `stderr` logs or thread-dump artifacts proving zero race detector warnings, plus before/after active thread/goroutine count diff.

### 6. Escalation to `Edge-Case`
* When the verified scope involves high-contention resources, atomic counters, distributed locks, or complex multi-table ACID transactions:
  * Complete baseline functional verification and record standard PASS.
  * Record an explicit recommendation in the report for CEO to dispatch an `Edge-Case` lane to perform hostile chaos attacks (microsecond timing collisions, out-of-order execution, process crash mid-flight).
