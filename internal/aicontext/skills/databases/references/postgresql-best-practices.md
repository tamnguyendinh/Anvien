# PostgreSQL Best Practices

High-impact PostgreSQL review checks for schema and query performance. Use this before adding indexes, reviewing slow queries, or changing relational access paths.

## How To Use
1. Identify the exact query or write path that is slow or risky.
2. Run `EXPLAIN (ANALYZE, BUFFERS)` on representative data.
3. Apply the smallest index or query change that supports the workload.
4. Re-run the plan and compare scanned rows, scan type, sorting, and execution time.
5. Account for write overhead, index storage, and maintenance cost before shipping.

Use `CREATE INDEX CONCURRENTLY` for production tables when avoiding write locks matters. Do not run it inside a transaction block.

## Rule 1: Index Foreign Key Columns
PostgreSQL creates indexes for primary keys and unique constraints, but not automatically for referencing foreign key columns.

Missing foreign key indexes commonly cause slow parent lookups, slow joins, slow cascade checks, and full table scans on large child tables.

Before:
```sql
CREATE TABLE orders (
  id BIGINT PRIMARY KEY,
  customer_id BIGINT REFERENCES customers(id),
  total NUMERIC(10, 2)
);

SELECT * FROM orders WHERE customer_id = 12345;
```

Plan smell:
```text
Seq Scan on orders
  Filter: (customer_id = 12345)
  Rows Removed by Filter: 199997
```

After:
```sql
CREATE INDEX idx_orders_customer_id ON orders(customer_id);
```

Expected direction:
```text
Bitmap Index Scan on idx_orders_customer_id
  Index Cond: (customer_id = 12345)
```

Source workload benchmark:
| Metric | Before | After | Improvement |
|---|---:|---:|---:|
| Rows scanned | 200000 | 3 | 66666x fewer |
| Execution time | 9.17 ms | 0.11 ms | about 80x faster |

## Rule 2: Support JOIN-Heavy Queries
If a query joins a small parent set to a large child table, the child-side join key usually needs an index.

Before:
```sql
SELECT c.name, COUNT(o.id), SUM(o.total)
FROM customers c
JOIN orders o ON o.customer_id = c.id
WHERE c.id IN (100, 200, 300, 400, 500)
GROUP BY c.id, c.name;
```

Plan smell:
```text
Parallel Seq Scan on orders
  Filter: customer_id IN (...)
```

After:
```sql
CREATE INDEX idx_orders_customer_id ON orders(customer_id);
```

Expected direction:
```text
Nested Loop
  -> Index Scan on customers_pkey
  -> Bitmap Index Scan on idx_orders_customer_id
```

Source workload benchmark:
| Metric | Before | After | Improvement |
|---|---:|---:|---:|
| Orders scanned | 200000 | 22 | about 9000x fewer |
| Execution time | 11.86 ms | 0.25 ms | about 47x faster |

Prefer a composite index when the join also filters or sorts by stable columns:
```sql
CREATE INDEX idx_orders_customer_status_created
ON orders(customer_id, status, created_at DESC);
```

Check column order against the query: equality filters first, then range or ordering columns.

## Rule 3: Use Partial Indexes For Filtered Workloads
Full indexes cover every row. Partial indexes cover only rows matching the predicate, which can reduce scan work and index size when most queries target a subset.

Good candidates: soft-delete filters, active records, queue states, and tenant-local hot paths combined with another selective key.

Before:
```sql
SELECT * FROM orders
WHERE status = 'pending'
ORDER BY created_at DESC
LIMIT 10;
```

Plan smell:
```text
Parallel Seq Scan on orders
  Filter: (status = 'pending')
Sort (top-N heapsort)
```

After:
```sql
CREATE INDEX idx_orders_pending_created
ON orders(created_at DESC)
WHERE status = 'pending';
```

Expected direction:
```text
Index Scan using idx_orders_pending_created
```

Source workload benchmark:
| Metric | Before | After | Improvement |
|---|---:|---:|---:|
| Rows processed | 200000 -> 40000 -> 10 | 10 direct | about 20000x fewer |
| Execution time | 10.39 ms | 0.10 ms | about 107x faster |

## Review Checklist
- Foreign key columns used in lookups, joins, or cascades have supporting indexes.
- Large-table joins have an index on the large-table join key.
- Composite indexes match equality, range, and order requirements in useful order.
- Filtered hot paths use partial indexes when the predicate is stable and selective.
- `EXPLAIN (ANALYZE, BUFFERS)` confirms fewer scanned rows or better scan type.
- Added indexes are justified against write overhead, storage, and maintenance cost.
