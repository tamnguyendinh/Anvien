---
name: databases
description: Use when creating, changing, or reviewing MongoDB/PostgreSQL database work in a real repository, including schema, models, migrations, queries, indexes, transactions, connection config, backup/restore assumptions, or performance; follow repo-native tooling and meet the production database bar unless the user explicitly asks for a prototype.
license: MIT
---

# Databases Skill

Repo-first guidance for creating, changing, and reviewing MongoDB or PostgreSQL database work in real repositories.

## Purpose

Use this skill when the task involves database schema, models, migrations, queries, indexes, transactions, seed data, connection config, backup/restore assumptions, or database performance.

The goal is not to teach basic database syntax. The goal is to help the agent produce database work that is correct, production-ready, and consistent with the target repository.

## Core Rule

Follow the user's requested action.

If the user asks for review, review. If the user asks for a plan, plan. If the user asks for implementation, implementation is allowed.

In every mode, database work must meet the production database bar unless the user explicitly asks for a throwaway prototype.

## Repo-First Rule

Before proposing or editing database work, inspect the target repository first:
- Database engine and version
- ORM, query builder, or data-access layer
- Migration framework and migration conventions
- Existing schema, model, index, and query patterns
- Connection config and environment conventions
- Tests, seed data, fixtures, and deployment assumptions

Follow repo-native tooling and patterns unless there is a clear reason not to.

## Production Database Bar

Production-ready database work must account for:
- Correctness: constraints, nullability, uniqueness, referential rules, validation, and ownership boundaries.
- Migration safety: forward path, rollback or mitigation path, deploy order, locking risk, data volume, and backfill behavior.
- Performance: expected queries, indexes, cardinality, pagination, aggregation cost, and query plans where relevant.
- Concurrency: transaction boundaries, isolation assumptions, duplicate writes, retries, and out-of-order events.
- Security: least privilege, secret handling, PII, encryption assumptions, and auditability.
- Operations: backup/restore assumptions, monitoring, slow-query visibility, retention, and failure recovery.
- Repo fit: existing ORM, migration framework, naming, config, and test patterns.

Do not produce demo-only database work unless the user explicitly requests it.

## Implementation Rule

When implementation is requested, create the production-safe version directly.

Do not leave integrity, migration safety, indexing, transaction behavior, or operational checks as vague future hardening unless the user explicitly scopes them out.

Keep edits scoped to the requested database behavior and the repository's established database layer.

## Evidence Rule

Database recommendations need evidence appropriate to the risk.

For performance work, prefer real query evidence: `EXPLAIN (ANALYZE, BUFFERS)`, `pg_stat_statements`, MongoDB explain/profiler output, row counts, index usage, and before/after measurements.

For migrations, identify affected data, lock risk, deployment order, rollback or mitigation path, and verification queries.

For schema changes, trace callers and consumers before changing shared data shapes.

## Tooling Boundary

This skill defines the production-quality bar for database work.

It does not replace the repository's ORM, migration framework, backup system, restore process, deployment process, or observability stack.

Use the repository's own tooling first.

## Prototype Exception

If the user explicitly asks for a prototype, demo, spike, or throwaway implementation, the work may be lighter.

Even then, clearly mark which production concerns are intentionally omitted.

## Task Routing

- Database selection: evaluate domain shape, consistency needs, query workload, operations, and team tooling.
- Schema/model change: inspect models, migrations, constraints, indexes, serializers, API consumers, and tests.
- Migration: use repo-native migration tooling; check data volume, locks, rollback/mitigation, and deploy order.
- Query/index performance: inspect query shape, current indexes, table/collection size, cardinality, and plan evidence.
- Backup/restore: inspect existing operational procedures; do not invent a generic production backup process.
- Analytics/reporting SQL: if the task is BI/reporting/dashboard analysis rather than application database engineering, prefer a data-analysis skill.

## Reference Navigation

- [database-selection.md](references/database-selection.md) - MongoDB vs PostgreSQL selection and tradeoffs.
- [mongodb-crud.md](references/mongodb-crud.md) - MongoDB CRUD, query operators, atomic updates.
- [mongodb-aggregation.md](references/mongodb-aggregation.md) - MongoDB aggregation pipeline patterns.
- [mongodb-indexing.md](references/mongodb-indexing.md) - MongoDB index design and optimization.
- [mongodb-atlas.md](references/mongodb-atlas.md) - MongoDB Atlas setup, monitoring, and operational reference.
- [postgresql-best-practices.md](references/postgresql-best-practices.md) - PostgreSQL FK indexes, join indexes, partial indexes.
- [postgresql-queries.md](references/postgresql-queries.md) - PostgreSQL query patterns.
- [postgresql-performance.md](references/postgresql-performance.md) - PostgreSQL EXPLAIN and performance workflow.
- [postgresql-psql-cli.md](references/postgresql-psql-cli.md) - psql commands and scripting reference.
- [postgresql-administration.md](references/postgresql-administration.md) - PostgreSQL administration, backup, replication, and maintenance.
