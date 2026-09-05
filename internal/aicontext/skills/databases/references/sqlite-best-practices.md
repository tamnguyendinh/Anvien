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
