# Database Selection

Use this reference when choosing between MongoDB and PostgreSQL, explaining tradeoffs, or orienting a repo that could plausibly use either document or relational storage.

## Choose MongoDB When

- Schema flexibility matters and document shapes change frequently.
- The domain is naturally document-centric or hierarchical.
- Reads usually fetch a complete aggregate/document rather than many normalized relationships.
- Horizontal sharding and high write throughput are first-order requirements.
- The team accepts denormalization and application-level consistency tradeoffs.

Common fits: content management, catalogs, IoT/time-series style ingestion, real-time analytics, mobile app data, user profiles.

## Choose PostgreSQL When

- Strong consistency, ACID transactions, constraints, and referential integrity are critical.
- The domain has complex relationships or many-to-many joins.
- SQL, reporting, BI tooling, or analytical query capability is important.
- The team needs strict schema validation and mature operational tooling.
- Window functions, CTEs, complex joins, and extensions are valuable.

Common fits: financial systems, e-commerce transactions, ERP, CRM, operational SaaS data, reporting, analytics.

## Both Can Support

- JSON/JSONB-style document storage and querying.
- Full-text search.
- Geospatial queries and indexing.
- Replication and high availability.
- ACID transactions, with different ergonomics and operational tradeoffs.
- Strong security features when configured correctly.

## Practical Selection Questions

- What are the top read and write access patterns?
- Does the app need cross-entity transactions or database-enforced relationships?
- Are query/reporting needs mostly SQL-shaped or document-shaped?
- How often will the schema change, and who owns migrations?
- What database expertise and operational tooling does the team already have?
- What failure mode is more acceptable: duplicated/denormalized data drift, or relational/migration complexity?

## Quick Difference Summary

| Feature | MongoDB | PostgreSQL |
|---------|---------|------------|
| Data model | Document BSON/JSON | Relational tables/rows |
| Schema | Flexible, application-enforced | Strict, database-enforced |
| Query language | MongoDB Query Language | SQL |
| Joins | `$lookup`, usually limited use | Native, optimized |
| Transactions | Supported, but document-first ergonomics | Native ACID default |
| Scaling | Built around sharding | Vertical by default; horizontal via architecture/extensions |
| Indexes | Single, compound, text, geo, hashed, wildcard | B-tree, hash, GiST, GIN, BRIN, partial, expression |

## Source Docs

- MongoDB: https://www.mongodb.com/docs/
- PostgreSQL: https://www.postgresql.org/docs/
