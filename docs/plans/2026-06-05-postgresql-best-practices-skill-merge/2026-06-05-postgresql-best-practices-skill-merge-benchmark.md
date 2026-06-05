# Benchmark Ledger

Title: PostgreSQL Best Practices Skill Merge

Date: 2026-06-05

Status: Planned

Companion files:

- Plan: [2026-06-05-postgresql-best-practices-skill-merge-plan.md](2026-06-05-postgresql-best-practices-skill-merge-plan.md)
- Evidence ledger: [2026-06-05-postgresql-best-practices-skill-merge-evidence.md](2026-06-05-postgresql-best-practices-skill-merge-evidence.md)

## Benchmark Rules

1. Record measured inventory, size, source benchmark values, and generated-output counts only.
2. Build/test pass-fail belongs in the evidence ledger unless timing or count is the measured target.
3. Update this file as each benchmarkable implementation slice completes.
4. Treat source workload speedups as example evidence, not universal performance guarantees.

## B0 - Baseline Skill Inventory

Measured before implementation on 2026-06-05.

| Metric | Baseline | Unit | Source |
|---|---:|---|---|
| Files under `internal/aicontext/skills/databases` | 18 | files | recursive file count |
| Reference files under `databases/references` | 8 | files | directory file count |
| PostgreSQL reference files | 4 | files | `postgresql-*.md` count |
| MongoDB reference files | 4 | files | inferred from reference inventory |
| `databases/SKILL.md` length | 173 | lines | `Measure-Object` |
| Clockwork source document length | 234 | lines | `Measure-Object` |

## B1 - Target Skill Inventory

Measured after P1 implementation on 2026-06-05.

| Metric | Baseline | Current | Unit |
|---|---:|---:|---|
| Files under `internal/aicontext/skills/databases` | 18 | 19 | files |
| Reference files under `databases/references` | 8 | 9 | files |
| PostgreSQL reference files | 4 | 5 | files |
| New PostgreSQL best-practice rule sections | 0 | 3 | sections |
| New reference length | 0 | 142 | physical lines |
| MongoDB reference files changed | 0 | 0 | files |
| Python utility script files changed | 0 | 0 | files |

## B2 - Source Benchmark Values To Preserve

These values come from the Clockwork source workload and should be preserved as example benchmark evidence in compact form.

| Rule | Source Before | Source After | Source Improvement |
|---|---:|---:|---:|
| Foreign key lookup index | 9.17 ms | 0.11 ms | about 80x faster |
| JOIN with supporting foreign key index | 11.86 ms | 0.25 ms | about 47x faster |
| Partial index for pending orders | 10.39 ms | 0.10 ms | about 107x faster |

## B3 - Validation Inventory To Record After Implementation

Record these after implementation completes.

| Metric | Final | Unit | Evidence Source |
|---|---:|---|---|
| Demo-only install strings in new reference | 0 | matches | content search |
| `SKILL.md` links to new reference | 1 | boolean | content search |
| New reference benchmark tables present | 3 | count | content search |
| `anvien detect-changes` changed files | 5 | files | detect-changes output |
| Commit hash recorded | pending | boolean | git output |

## B4 - Validation Analyze Inventory

Measured in clean validation worktree after applying this slice.

| Metric | Final | Unit |
|---|---:|---|
| Files scanned | 1369 | files |
| Parsed code files | 697 | files |
| Failed files | 0 | files |
| Graph nodes | 83582 | nodes |
| Graph relationships | 122338 | relationships |
