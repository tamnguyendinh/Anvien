# Benchmark Ledger

Date: 2026-06-04

Status: Approved - implementation pending

Companion files:

- Plan: [2026-06-04-anvien-skill-mirror-incremental-sync-plan.md](2026-06-04-anvien-skill-mirror-incremental-sync-plan.md)
- Evidence ledger: [2026-06-04-anvien-skill-mirror-incremental-sync-evidence.md](2026-06-04-anvien-skill-mirror-incremental-sync-evidence.md)

## Benchmark Rules

1. Record measured numbers only.
2. Build/test pass-fail belongs in the evidence ledger unless timing, count, size, or throughput is the measured target.
3. For this plan, benchmarkable data includes generated-output inventory, source package inventory, sync operation counts, output size, and manifest counts.
4. Record measurements as the matching implementation or validation phase completes.

## B0 - Planned Metrics

| Metric | Unit | Baseline | Latest | Final | Target | Notes |
| --- | --- | ---: | ---: | ---: | ---: | --- |
| Source skill packages | count | pending | pending | pending | current source count | Count immediate child package roots under `internal/aicontext/skills` |
| Source skill files | count | pending | pending | pending | current source payload count | Count files included in desired snapshot |
| Generated skill files | count | pending | pending | pending | equals source payload count | Excludes `.anvien-skill-manifest.json` |
| Manifest packages | count | pending | pending | pending | equals source package count | No stale deleted packages |
| Manifest files | count | pending | pending | pending | equals source payload count | Sum per-package file counts |
| Sync writes | count | pending | pending | pending | measured | New desired files written |
| Sync overwrites | count | pending | pending | pending | measured | Changed or tampered output repaired |
| Sync deletes | count | pending | pending | pending | measured | Actual files absent from desired snapshot |
| Sync skips | count | pending | pending | pending | measured | Unchanged files left untouched |
| Sync collisions | count | pending | pending | pending | 0 expected | Target paths or ownership conflicts encountered |
| Unsafe filesystem entries | count | pending | pending | pending | 0 expected | Non-regular or unsafe target entries encountered |
| Generated output size | bytes | pending | pending | pending | measured | Size of `.claude/skills/anvien/**` excluding manifest if measured separately |

## B1 - Baseline Inventory

Pending. Capture after plan approval and before implementation if needed.

## B2 - Final Inventory

Pending. Capture after implementation and real repo analyze validation.
