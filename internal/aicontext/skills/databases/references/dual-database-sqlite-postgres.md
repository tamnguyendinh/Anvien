# Dual Database Architecture: SQLite & PostgreSQL (Pitfalls & Protocols)

> Read when: Designing, implementing, or reviewing systems running SQLite locally (offline desktop/CLI client) and PostgreSQL remotely (cloud server sync), or building an abstracted database layer supporting both engines across any programming language.

---

## 1. Architectural Model & Scope

In Local-First and Desktop-Cloud Hybrid Architectures:
- **Client / Local Edge (SQLite):** Embedded, single-file database providing offline capabilities, zero-latency local operations, and local data privacy.
- **Server / Central Cloud (PostgreSQL):** Multi-tenant, highly concurrent relational database on the server for data synchronization, team collaboration, and centralized backups.

**Core Danger:** Code written and verified against PostgreSQL frequently crashes or produces silent corruption when executed against SQLite (and vice versa) due to file-locking conflicts, native driver build hurdles, and SQL dialect divergence.

---

## 2. Driver Rule: Avoid Native Build Breakages Across Runtimes

When packaging cross-platform Desktop or CLI applications (Windows, macOS, Linux), SQLite drivers depending on native C code frequently fail during terminal compilation or CI/CD cross-compilation:

| Language | Risky / Build-Breaking Driver (Avoid) | Recommended Safe Driver |
| :--- | :--- | :--- |
| **Go** | `github.com/mattn/go-sqlite3` (requires gcc/Cgo) | **`modernc.org/sqlite`** (100% Pure-Go, compiles with `CGO_ENABLED=0`) |
| **Node.js / Electron** | `better-sqlite3` (prone to `node-gyp` ABI rebuild breakages across Electron versions) | Use official prebuilt binaries of **`sqlite3`** or **`@libsql/client`** (WASM/Rust native) |
| **Rust (Tauri)** | Unpinned `libsqlite3-sys` relying on host environment libraries | Use **`rusqlite` with the `bundled` feature** or **`sqlx`** (compiles embedded SQLite without host dependencies) |
| **Python** | External C-wrapper packages | Use standard library **`sqlite3`** included in Python runtime |
| **C# / .NET** | Legacy COM/Interop bindings | **`Microsoft.Data.Sqlite`** (uses `SQLitePCLRaw`, bundles precompiled binaries for all platforms) |

---

## 3. Write Lock Handling & Preventing `database is locked`

PostgreSQL utilizes MVCC to permit thousands of concurrent read/write connections. Conversely, SQLite is a **single physical file**; multiple concurrent write connections immediately throw `database is locked` or `busy` errors.

### Mandatory Initialization Commands Across All Runtimes:
1. **Enable Write-Ahead Logging (WAL):** Allows readers to execute concurrently without blocking the writer:
   `PRAGMA journal_mode = WAL;`
2. **Set Busy Timeout:** Prevents instant failure when the file is locked:
   `PRAGMA busy_timeout = 5000;` (wait at least 5000ms)
3. **Golden Connection Pool Rule:**
   * **Go:** `sqlDB.SetMaxOpenConns(1)` on the writer connection.
   * **Rust (SQLx):** `SqlitePoolOptions::new().max_connections(1)` on the writer pool.
   * **Python:** `sqlite3.connect(..., timeout=5.0)` with a single dedicated writer connection.
   * **Node.js:** Ensure write operations are serialized through a single database instance rather than opening uncoordinated multiple file handles.

---

## 4. Dialect Incompatibility Matrix

Planner and Coder must strictly adhere to this matrix to prevent runtime dialect failures:

| Feature / Concept | PostgreSQL | SQLite | Safe Cross-Engine Practice |
| :--- | :--- | :--- | :--- |
| **Parameter Placeholders** | `$1, $2, $3` | `?` (or `?1, ?2`) | Use Query Builders / ORMs to automatically translate parameter placeholders. |
| **Boolean Values** | `BOOLEAN` (`true`/`false`) | `INTEGER` (`1`/`0`) | Always use native language booleans in models; let the ORM/driver map `true` $\leftrightarrow$ `1`. |
| **DateTime & Timezone** | `TIMESTAMPTZ` (native) | `TEXT` (ISO string) / `INTEGER` (Unix epoch) | Always persist UTC ISO-8601 strings (`YYYY-MM-DDTHH:MM:SSZ`) or Unix timestamps. Avoid DB-specific date functions. |
| **Case-Insensitive Search** | `ILIKE` | `LIKE` (case-insensitive for ASCII only) | Avoid `ILIKE`. Use standard cross-platform SQL: `LOWER(column) LIKE LOWER(?)`. |
| **JSON Operations** | `->>`, `jsonb_extract_path` | `json_extract()` | Encapsulate JSON queries inside Repository methods; never expose raw engine-specific JSON syntax. |
| **Upsert (Insert or Update)** | `ON CONFLICT (...) DO UPDATE` | `ON CONFLICT (...) DO UPDATE` (since v3.24+) | Use standard SQL-92 `ON CONFLICT` syntax; ensure target SQLite runtime is $\ge 3.24$. |
| **Auto-Increment Primary Key** | `BIGSERIAL` or `UUID` | `INTEGER PRIMARY KEY AUTOINCREMENT` | **Mandatory client-generated UUIDv4 or ULID strings.** Completely eliminates ID collisions during local-to-cloud sync. |

---

## 5. Architectural Abstraction (Multi-Language Repository & ORM)

**STRICTLY FORBIDDEN:** Writing raw dialect-locked SQL strings inside application business services or API handlers.

Mandatory use of Query Builders or ORMs supporting multi-dialect translation:
* **Go:** **Bun (`uptrace/bun`)** (lightweight, native support for `pgdialect` and `sqlitedialect`) or GORM.
* **TypeScript / JS:** **Drizzle ORM** (modular switching between `drizzle-orm/pg-core` and `drizzle-orm/sqlite-core`), **Prisma**, or **Kysely**.
* **Rust:** **SeaORM** or **SQLx** (abstracting backend drivers across Postgres and SQLite via traits).
* **Python:** **SQLAlchemy** (engine switching via `sqlite:///` and `postgresql://` connection strings with identical model declarations).
* **C# / .NET:** **Entity Framework Core (EF Core)** (identical `DbContext` and Entity Models, toggling `UseSqlite()` and `UseNpgsql()`).

---

## 6. Mandatory Planner Checklist for Dual SQLite & PostgreSQL Architecture

When creating execution plans involving SQLite and PostgreSQL (or data sync):
- [ ] Explicitly identify the safe SQLite driver for the project runtime (preventing native Cgo/build breakages).
- [ ] Mandate SQLite initialization with `WAL` mode and `PRAGMA busy_timeout=5000`.
- [ ] Configure SQLite write connections as single-writer (`MaxOpenConns=1` or equivalent).
- [ ] Mandate client-generated UUID/ULID as primary keys instead of engine auto-increment IDs.
- [ ] Enforce the Repository pattern or multi-dialect query builders (Bun, Drizzle, SQLAlchemy, EF Core, SeaORM).
