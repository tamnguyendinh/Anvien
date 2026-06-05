# Benchmark Ledger

Title: Anvien Package Runtime Skill Subtree Staging
Date: 2026-06-06
Status: Completed
Plan: docs/plans/2026-06-06-anvien-package-runtime-skill-subtree-staging/2026-06-06-anvien-package-runtime-skill-subtree-staging-plan.md
Evidence: docs/plans/2026-06-06-anvien-package-runtime-skill-subtree-staging/2026-06-06-anvien-package-runtime-skill-subtree-staging-evidence.md
Benchmark: docs/plans/2026-06-06-anvien-package-runtime-skill-subtree-staging/2026-06-06-anvien-package-runtime-skill-subtree-staging-benchmark.md

## Benchmark Rules

- Record measured product/runtime performance, package size, startup size, graph throughput, DB throughput, graph inventory counts, or generated-output inventory when the phase changes those systems.
- Build/test pass-fail belongs in the evidence file unless timing or size is the measured target.
- Inventory counts that prove staging completeness belong here.

## B0 - Initial Graph Inventory

| Metric | Value | Source |
| --- | ---: | --- |
| Indexed files before plan creation | 1399 | `anvien analyze . --force` |
| Graph nodes before plan creation | 84382 | `anvien analyze . --force` |
| Graph relationships before plan creation | 122854 | `anvien analyze . --force` |
| Indexed files after plan creation | 1402 | `anvien analyze . --force` |
| Graph nodes after plan creation | 84408 | `anvien analyze . --force` |
| Graph relationships after plan creation | 122880 | `anvien analyze . --force` |
| File projection dependency edges after plan creation | 16569 | `anvien analyze . --force` |
| Unresolved items after plan creation | 430 | `anvien analyze . --force` |
| Indexed files before implementation | 1402 | `anvien analyze . --force` |
| Graph nodes before implementation | 84409 | `anvien analyze . --force` |
| Graph relationships before implementation | 122881 | `anvien analyze . --force` |
| File projection dependency edges before implementation | 16569 | `anvien analyze . --force` |
| Unresolved items before implementation | 430 | `anvien analyze . --force` |

## B1 - Package Runtime Staging Inventory

| Metric | Value | Source |
| --- | ---: | --- |
| Final indexed files | 1404 | final `anvien analyze . --force` before commit |
| Final graph nodes | 84448 | final `anvien analyze . --force` before commit |
| Final graph relationships | 122937 | final `anvien analyze . --force` before commit |
| Final file projection dependency edges | 16570 | final `anvien analyze . --force` before commit |
| Final unresolved items | 430 | final `anvien analyze . --force` before commit |
| Package runtime staged files | 929 | repo-local `go run ..\cmd\anvien package prepare-go-source` smoke |
| Staged skill subtree files | 638 | repo-local `go-src` inventory |
| Staged non-`.md` skill subtree files | 295 | repo-local `go-src` inventory |
| Manifest file count | 929 | `go-src/anvien-go-source.json` |
| Manifest skill path count | 638 | `go-src/anvien-go-source.json` |

Representative staged non-`.md` skill paths:

- `internal/aicontext/skills/better-auth/scripts/better_auth_init.py`
- `internal/aicontext/skills/better-auth/scripts/requirements.txt`
- `internal/aicontext/skills/better-auth/scripts/tests/.coverage`
- `internal/aicontext/skills/better-auth/scripts/tests/test_better_auth_init.py`
- `internal/aicontext/skills/databases/scripts/db_backup.py`
