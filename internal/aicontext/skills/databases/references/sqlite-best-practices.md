# SQLite Production Best Practices

Authoritative guide for running SQLite in production across Embedded Systems, Desktop Applications, Local-First architectures, and Edge/Serverless services.

---

## 1. Authoritative References & Foundation

This guide compiles production invariants strictly sourced from authoritative engineering documentation and battle-tested production whitepapers:
- **SQLite Consortium (`sqlite.org`)**:
  - [Write-Ahead Logging (WAL)](https://www.sqlite.org/wal.html)
  - [File Locking And Concurrency In SQLite Version 3](https://www.sqlite.org/lockingv3.html)
  - [PRAGMA Statements & Compile/Runtime Lifecycle](https://www.sqlite.org/pragma.html)
  - [How To Corrupt An SQLite Database File](https://www.sqlite.org/howtocorrupt.html)
  - [STRICT Tables Specification](https://www.sqlite.org/stricttables.html)
  - [The SQLite Query Optimizer & EXPLAIN QUERY PLAN](https://www.sqlite.org/eqp.html)
  - [SQLite Online Backup API](https://www.sqlite.org/backup.html)
- **Fly.io & Litestream Architecture** (Ben Johnson - author of Litestream & LiteFS): Production multi-tenant SQLite, WAL streaming replication, and Connection Pool segregation.
- **Mozilla Firefox Storage & Android AOSP Room**: Thread safety, non-blocking UI query queues, and integrity checking.

---

## 2. Connection Initialization Baseline (Mandatory PRAGMAs)

By default, SQLite operates in legacy compatibility mode (Rollback Journal, no foreign key checks, 2MB cache). Every production database connection MUST configure the following PRAGMAs immediately upon opening:

```sql
-- 1. Enable Write-Ahead Logging (Persistent database setting)
-- Allows concurrent readers while a write transaction is in progress.
PRAGMA journal_mode = WAL;

-- 2. Set Synchronous Mode (Connection setting)
-- In WAL mode, 'NORMAL' is 100% crash-safe against application crashes and OS faults,
-- providing orders of magnitude faster writes than 'FULL' while maintaining structural integrity.
PRAGMA synchronous = NORMAL;

-- 3. Set Busy Timeout (Connection setting)
-- Prevents immediate SQLITE_BUSY errors. Waits up to 5000ms for locks to clear.
PRAGMA busy_timeout = 5000;

-- 4. Enable Foreign Key Enforcement (Connection setting)
-- CRITICAL: SQLite disables foreign key constraint checks by default for backward compatibility.
-- MUST be executed on every newly opened connection!
PRAGMA foreign_keys = ON;

-- 5. Increase Page Cache (Connection setting)
-- Negative value sets cache size in KiB (-64000 = ~64MB RAM cache, default is only 2MB).
PRAGMA cache_size = -64000;

-- 6. Store Temporary Tables and Indices in Memory (Connection setting)
-- Avoids disk I/O for temporary sorting and aggregation structures.
PRAGMA temp_store = MEMORY;

-- 7. Enable Memory-Mapped I/O (Connection setting, 64-bit systems)
-- Maps up to 256MB into process memory for near-zero latency reads (verify OS compatibility).
PRAGMA mmap_size = 268435456;
```

---

## 3. Concurrency, Locking & Deadlock Elimination

### The Concurrency Model in WAL Mode
- **Multi-Reader, Single-Writer**: Unlimited concurrent readers (`-shm` shared memory index), exactly **1 active writer** appending to the `-wal` file.
- Readers never block writers; writers never block readers.

### The Deadlock Trap: `BEGIN DEFERRED` vs `BEGIN IMMEDIATE`
- **The Pitfall**: A standard `BEGIN` or `BEGIN DEFERRED` starts a transaction in read-only mode. If two concurrent transactions read data and then both attempt to execute an `UPDATE`/`INSERT`, both try to upgrade their lock to `RESERVED`. Neither can proceed because each is waiting on the other's shared read lock, resulting in immediate `SQLITE_BUSY: database is locked` or deadlock.
- **The Production Invariant**: Any transaction that intends to write MUST explicitly start with **`BEGIN IMMEDIATE;`**.

```sql
-- DANGEROUS: Prone to upgrade deadlocks in multi-threaded environments
BEGIN;
SELECT balance FROM accounts WHERE id = 1;
UPDATE accounts SET balance = balance - 100 WHERE id = 1; -- Fails with SQLITE_BUSY!
COMMIT;

-- PRODUCTION STANDARD: Acquires write intent lock immediately at start
BEGIN IMMEDIATE;
SELECT balance FROM accounts WHERE id = 1;
UPDATE accounts SET balance = balance - 100 WHERE id = 1; -- Guaranteed safe from upgrade deadlocks
COMMIT;
```

### Connection Pool Architecture
In multi-threaded backends (Go, Node.js, Python, Rust, C#):
- **Writer Pool**: Dedicated **single-connection writer** (`MaxOpenConns = 1` or protected by an in-process mutex). Serializes write operations cleanly without lock contention.
- **Reader Pool**: Multi-connection reader pool (`MaxOpenConns = NumCPU * 2`) with `read_only = true` or `PRAGMA query_only = ON;`.

---

## 4. Schema Design & Data Integrity

### STRICT Tables (SQLite 3.37.0+)
SQLite normally uses "Type Affinity" (allows storing strings in integer columns). For production reliability, enforce strict typing using the `STRICT` keyword:

```sql
CREATE TABLE users (
    id TEXT PRIMARY KEY,
    username TEXT NOT NULL UNIQUE,
    age INTEGER NOT NULL CHECK (age >= 0),
    balance REAL NOT NULL DEFAULT 0.0,
    avatar BLOB,
    created_at TEXT NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%fZ', 'now'))
) STRICT;
```
*Allowed STRICT types: `INT`, `INTEGER`, `REAL`, `TEXT`, `BLOB`, `ANY`.*

### Primary Key Selection: Local vs Distributed
- **Single-node / Desktop Standalone**:
  Use `id INTEGER PRIMARY KEY` (alias for the 64-bit signed `rowid`). Fast B-Tree lookups and minimal storage.
  *Avoid `AUTOINCREMENT`* unless strictly necessary to prevent ID reuse; `AUTOINCREMENT` adds overhead by updating `sqlite_sequence`.
- **Distributed / Local-First / Sync**:
  Use `id TEXT PRIMARY KEY` storing client-generated **UUIDv4** or **ULID** (lexicographically sortable). Completely eliminates ID collisions during sync.

### Date & Time Standard
SQLite has no dedicated datetime type. Standardize across all models:
- **Option A (Recommended for human readability & SQL filtering)**: ISO-8601 UTC `TEXT` (`YYYY-MM-DDTHH:MM:SS.SSSZ`).
- **Option B (Recommended for ultra-high throughput / minimal disk)**: Unix epoch timestamp in milliseconds `INTEGER`.

### Boolean Representation
Store booleans as `INTEGER` with a check constraint:
```sql
is_active INTEGER NOT NULL DEFAULT 1 CHECK (is_active IN (0, 1))
```

---

## 5. Query Optimization & Indexing Patterns

### Analyzing Execution Plans (`EXPLAIN QUERY PLAN`)
Always verify complex queries using `EXPLAIN QUERY PLAN`:

```sql
EXPLAIN QUERY PLAN
SELECT u.username, o.total
FROM users u
JOIN orders o ON u.id = o.user_id
WHERE u.status = 'active' AND o.created_at >= '2026-01-01';
```

- ❌ **RED FLAG**: `SCAN TABLE <table>` (Full table scan).
- ✅ **ACCEPTABLE**: `SEARCH TABLE <table> USING INDEX <index_name> (<columns>)`.
- 🚀 **OPTIMAL**: `SEARCH TABLE <table> USING COVERING INDEX <index_name>` (SQLite reads exclusively from B-Tree index without accessing table data pages).

### Covering & Partial Indexes
```sql
-- Partial Index: Only indexes active rows, saving 90% index size
CREATE INDEX idx_orders_active_pending
ON orders (user_id, created_at)
WHERE status = 'pending';

-- Covering Index: Query satisfied entirely by index pages
CREATE INDEX idx_users_lookup
ON users (status, id, username);
```

### Full-Text Search (FTS5)
For high-performance text search inside client applications without external search engines:
```sql
CREATE VIRTUAL TABLE documents_fts USING fts5(
    title,
    content,
    content='documents',
    content_rowid='id'
);

-- Search query with ranking
SELECT d.title, d.content, bm25(documents_fts) AS rank
FROM documents_fts f
JOIN documents d ON f.rowid = d.id
WHERE documents_fts MATCH 'database optimization'
ORDER BY rank;
```

---

## 6. Operational Safety, Corruption Prevention & Maintenance

### Preventing Database Corruption (`sqlite.org/howtocorrupt.html`)
1. **NEVER host active SQLite databases on Network Filesystems (NFS, SMB, CIFS, GlusterFS)**:
   Network file locking protocols frequently fail or implement broken POSIX byte-range locking, causing catastrophic corruption under concurrent access.
2. **NEVER copy database files via raw shell commands (`cp`, `rsync`) while writers are active**:
   Copying while WAL pages are in flight produces torn or corrupt database snapshots.

### Zero-Downtime Hot Backups
Always use one of the following two safe backup strategies:
```sql
-- Method 1: Atomic Online Backup via SQL (SQLite 3.27.0+)
VACUUM INTO '/backups/backup_2026_09_05.db';
```
```bash
# Method 2: SQLite CLI backup command (uses official sqlite3_backup API)
sqlite3 /data/app.db ".backup /backups/backup_latest.db"
```
*Both methods lock safely without disrupting concurrent readers or writers.*

### Database Health Checks
Integrate integrity audits into startup or scheduled health checks:
```sql
-- Fast check (checks B-Tree structure and page linkages)
PRAGMA quick_check;
-- Result: 'ok'

-- Exhaustive check (verifies all data, indexes, and constraints)
PRAGMA integrity_check;
-- Result: 'ok'
```

### Maintenance & Query Planner Statistics (`PRAGMA optimize`)
SQLite Query Planner optimizes routes based on table statistics:
- Run `PRAGMA optimize;` immediately before closing database connections or hourly on long-lived connections.
- When tables grow significantly or after schema migrations, running `PRAGMA optimize;` ensures the planner doesn't pick suboptimal full table scans.

### Space Reclamation without Blocking (`INCREMENTAL VACUUM`)
Running a full `VACUUM;` rewrites the entire database file and requires exclusive locks, freezing applications with large databases. Use incremental vacuuming instead:
```sql
-- Run once when creating the database:
PRAGMA auto_vacuum = INCREMENTAL;

-- Periodically run during idle times to release free pages in chunks:
PRAGMA incremental_vacuum(1000); -- Frees up to 1,000 pages at a time
```

---

## 7. High-Throughput & Extreme Performance Tuning (10x - 1000x Speedup)

When dealing with large volumes of data, real-time telemetry, local sync pipelines, or high-concurrency desktop/edge workloads, standard query patterns create severe bottlenecks. The following patterns deliver orders-of-magnitude performance gains.

### 7.1. Bulk Insert Acceleration: From 50 to 250,000+ writes/sec

#### The Root Cause of Slow Inserts
Without an explicit transaction, SQLite operates in auto-commit mode: every individual `INSERT` statement is treated as a separate transaction requiring a full disk sync (`fsync`). This limits throughput to physical drive rotational/flush latency (**50 - 100 writes/second**).

#### Level 1: Explicit Transaction Encapsulation (~100x speedup: 50,000+ writes/sec)
Wrapping batches inside a single transaction groups disk flushes into one contiguous WAL write:

```sql
-- DANGEROUS: 1,000 separate disk syncs -> takes 15-20 seconds
INSERT INTO events (id, payload) VALUES ('id-1', 'data-1');
INSERT INTO events (id, payload) VALUES ('id-2', 'data-2');
-- ...

-- PRODUCTION STANDARD: 1 single disk sync -> takes 15 milliseconds
BEGIN TRANSACTION;
INSERT INTO events (id, payload) VALUES ('id-1', 'data-1');
INSERT INTO events (id, payload) VALUES ('id-2', 'data-2');
-- ... up to 10,000 records
COMMIT;
```

#### Level 2: Prepared Statement Reuse & Multi-Row Syntax (~300x speedup: 150,000+ writes/sec)
Eliminate repeated SQL parsing and query planning overhead:
1. **Prepare Once, Bind Many**: Compile SQL once (`sqlite3_prepare_v2`), bind parameters in a loop, call `sqlite3_step()`, then `sqlite3_reset()`.
2. **Multi-Row Batches**: Insert up to 500–1,000 tuples per statement:
```sql
INSERT INTO events (id, payload) VALUES
  ('id-1', 'data-1'),
  ('id-2', 'data-2'),
  ('id-3', 'data-3');
```

#### Level 3: Drop-and-Rebuild Index Pattern for Massive ETL (~1000x speedup)
When importing millions of records:
- Maintaining B-Tree indexes dynamically during massive ingestion forces constant re-balancing, node splits, and cache thrashing.
- **The Invariant**:
  1. Drop secondary indexes: `DROP INDEX IF EXISTS idx_events_timestamp;`
  2. Ingest raw data inside batched transactions.
  3. Recreate indexes in a single pass: `CREATE INDEX idx_events_timestamp ON events(timestamp);`
  *Rebuilding an index once from an existing table is 5x to 10x faster than updating it incrementally.*

### 7.2. Multi-Threaded Query Execution (`PRAGMA threads`)

By default, SQLite evaluates queries strictly on a single CPU thread. SQLite 3.8.7+ supports helper threads to parallelize sorting, hash joins, and index creation:

```sql
-- Allocate up to 4 worker threads for sorting and parallel index building
PRAGMA threads = 4;
```
*Note: SQLite threads do not parallelize simple B-Tree point lookups; they accelerate intensive CPU operations such as `ORDER BY` on large unindexed sets, window functions, and `CREATE INDEX`.*

### 7.3. Keyset Pagination (Cursor) vs The `OFFSET` Degeneration Trap

#### The `OFFSET` Anti-Pattern
```sql
-- TERRIBLE: As page number grows, SQLite must scan and discard 500,000 B-Tree entries
SELECT * FROM orders
ORDER BY created_at DESC, id DESC
LIMIT 20 OFFSET 500000; -- High disk read, high CPU, linear O(N) slowdown
```

#### The Production Invariant: Keyset (Seek) Pagination
Use indexed column comparisons to jump directly to the target B-Tree leaf node in $O(\log N)$ time:
```sql
-- Optimal Compound Index required:
CREATE INDEX idx_orders_pagination ON orders (created_at DESC, id DESC);

-- Keyset Query: Instantaneous execution (< 1ms) regardless of page depth
SELECT * FROM orders
WHERE (created_at, id) < (:last_seen_created_at, :last_seen_id)
ORDER BY created_at DESC, id DESC
LIMIT 20;
```

### 7.4. WAL Checkpoint Tuning (Eliminating I/O Freezes)

In high-write environments, default automatic checkpointing (every 1,000 pages / ~4MB) can cause micro-stalls when writes are continuous:

1. **Increase Autocheckpoint Threshold**:
   ```sql
   -- Smooth out high-burst ingestion by allowing WAL to grow up to ~40MB before triggering checkpoint
   PRAGMA wal_autocheckpoint = 10000;
   ```
2. **Non-Blocking Background Checkpointing**:
   Schedule idle background workers to merge WAL frames back into the main database without blocking active readers or writers:
   ```sql
   -- Checkpoints as many frames as possible without waiting for readers/writers
   PRAGMA wal_checkpoint(PASSIVE);
   ```

### 7.5. Page Size Hardware Alignment (`PRAGMA page_size`)

SQLite reads and writes data in units of pages. Aligning page size with your workload prevents write amplification and eliminates overflow overhead:

- **Default Standard (Recommended for 95% of OLTP workloads)**: Keep **`4096`** (4KB).
  *Rationale*: Matches the native 4KB cluster/block size of Windows NTFS, Linux ext4, and macOS APFS. Minimizes write amplification when updating small rows and maximizes RAM page cache capacity.
- **Document / JSON-Heavy Exception**: Set explicitly to **`8192`** (8KB).
  *Rationale*: When average row size exceeds 2KB (e.g., storing JSON payloads, rich text), 8KB prevents SQLite from spilling data into **Overflow Pages** (which require extra I/O seek hops), while avoiding the severe write amplification of 16KB.

```sql
-- MUST be executed BEFORE creating tables in a new database:
PRAGMA page_size = 8192; -- Only for JSON/Document-heavy databases
```
*Note: For existing databases, changing page size requires running `VACUUM;` immediately afterward.*

### 7.6. Query Predicate SARGability (Search Argument Able)

Ensure expressions in `WHERE` clauses preserve index utilization:
- ❌ **Breaks Index**: `WHERE lower(email) = 'user@example.com'` (Forces a full table scan `SCAN TABLE`).
- ✅ **Option 1 (Collation Index)**:
  ```sql
  CREATE TABLE users (email TEXT COLLATE NOCASE);
  CREATE INDEX idx_users_email ON users(email);
  SELECT * FROM users WHERE email = 'user@example.com'; -- Uses index
  ```
- ✅ **Option 2 (Expression Index)**:
  ```sql
  CREATE INDEX idx_users_lower_email ON users(lower(email));
  SELECT * FROM users WHERE lower(email) = 'user@example.com'; -- Uses index
  ```
