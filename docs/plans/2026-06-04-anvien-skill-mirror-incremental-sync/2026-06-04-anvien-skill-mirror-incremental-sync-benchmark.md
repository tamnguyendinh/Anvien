# Benchmark Ledger

Date: 2026-06-04

Status: Implementation validation complete

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
| Source skill packages | count | pending | 34 | 34 | current source count | Count immediate child package roots under `internal/aicontext/skills` |
| Source skill files | count | pending | 593 | 593 | current source payload count | Count files included in desired snapshot |
| Generated skill files | count | pending | 593 | 593 | equals source payload count | Excludes `.anvien-skill-manifest.json` |
| Manifest packages | count | pending | 34 | 34 | equals source package count | No stale deleted packages |
| Manifest files | count | pending | 593 | 593 | equals source payload count | Sum per-package file counts |
| Sync writes | count | pending | 0 | 0 | measured | Post-sync hash diff after real repo analyze |
| Sync overwrites | count | pending | 0 | 0 | measured | Post-sync hash diff after real repo analyze |
| Sync deletes | count | pending | 0 | 0 | measured | Post-sync hash diff after real repo analyze |
| Sync skips | count | pending | 593 | 593 | measured | Post-sync hash diff after real repo analyze |
| Sync collisions | count | pending | 0 | 0 | 0 expected | No duplicate desired target paths measured |
| Unsafe filesystem entries | count | pending | 0 | 0 | 0 expected | No non-regular target payload entries measured |
| Generated output size | bytes | pending | 11169926 | 11169926 | measured | Size of `.claude/skills/anvien/**` excluding manifest |

## B1 - Baseline Inventory

Pending. Capture after plan approval and before implementation if needed.

## B2 - Final Inventory

Measured after implementation and `anvien analyze --force` real repo validation:

```text
source_packages=34 manifest_packages=34 source_payload_files=593 target_payload_files=593 missing=0 extra=0 mismatch=0 debugging_source=False debugging_target=False
generated_output_size_bytes=11169926
```

Interpretation:

- generated target payload exactly matches the current source payload by path and SHA-256 hash;
- `.claude/skills/anvien/debugging` is absent because `internal/aicontext/skills/debugging` is absent;
- post-sync hash diff has 0 writes, 0 overwrites, 0 deletes, and 593 unchanged payload files.
