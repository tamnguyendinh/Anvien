# Anvien File Group Classification Benchmark Ledger

Date: 2026-06-03

Status: completed

Companion files:

- Plan: [2026-06-03-anvien-file-role-classification-gap-plan.md](2026-06-03-anvien-file-role-classification-gap-plan.md)
- Evidence ledger: [2026-06-03-anvien-file-role-classification-gap-evidence.md](2026-06-03-anvien-file-role-classification-gap-evidence.md)

## Benchmark Rules

1. Record quantitative data only.
2. Put command interpretation in the evidence ledger.
3. Preserve current, target, and delta where useful.
4. Track analyze inventory counts after every implementation slice.
5. Track file-role and file-group coverage counts after every implementation slice.
6. Build/test pass/fail belongs in the evidence ledger unless timing/count/size is the measured target.

## B0 - Analyze And File Projection Baseline

Status: recorded

Source evidence: E0.

| Metric | Unit | Baseline | Latest | Target |
|---|---:|---:|---:|---:|
| Files scanned | files | 821 | 828 | record |
| Parsed code files | files | 601 | 605 | no unintended decrease |
| Failed files | files | 0 | 0 | 0 |
| Graph nodes | nodes | 60974 | 61340 | record |
| Graph relationships | relationships | 96624 | 97225 | record |
| File projection files | files | 821 | 828 | record |
| File projection dependency edges | edges | 15965 | 16099 | record |
| Default-visible unresolved files | files | 336 | 338 | preserve unless separate behavior changes |
| Raw unresolved files | files | 353 | 355 | preserve unless separate behavior changes |
| Raw-only unresolved files | files | 17 | 17 | 17 sample files recognized, 0 unknown role, 17 sample files assigned file group after corrected implementation |

## B1 - Raw-Only Source-Site Baseline

Status: recorded

Source evidence: E1.

| Metric | Unit | Baseline | Latest | Target |
|---|---:|---:|---:|---:|
| Raw-only files | files | 17 | 17 | 17 |
| Raw-only source sites | source sites | 376 | 376 | preserve unless separate resolution work changes raw semantics |
| Raw-only production source sites | source sites | 0 | 0 | 0 |
| Raw-only non-actionable source sites | source sites | 376 | 376 | preserve unless separate resolution work changes raw semantics |
| Raw-only unknown source sites | source sites | 0 | 0 | 0 |

## B2 - Raw-Only Classification Baseline

Status: recorded

Source evidence: E1.

| Classification | Unit | Baseline | Latest | Target |
|---|---:|---:|---:|---:|
| `builtin` | source sites | 119 | 119 | record |
| `standard_library` | source sites | 254 | 254 | record |
| `test_framework` | source sites | 3 | 3 | record |
| `in_repo_unresolved` | source sites | 0 | 0 | 0 for raw-only 17-file set |
| `unclassified` | source sites | 0 | 0 | 0 |

## B3 - Raw-Only File List Baseline

Status: recorded

| File | Raw sites | Current role status | Target role |
|---|---:|---|---|
| `internal/frameworks/frameworks.go` | 209 | no first-class file role | `analyzer_helper` |
| `internal/scopeir/sort_keys.go` | 63 | no first-class file role | `helper` |
| `internal/group/types.go` | 16 | no first-class file role | `contract_model` |
| `internal/repo/paths.go` | 13 | no first-class file role | `storage_helper` |
| `internal/testutil/path.go` | 12 | no first-class file role | `test_helper` |
| `internal/repo/settings.go` | 11 | no first-class file role | `config` |
| `internal/repo/runtime_config.go` | 10 | no first-class file role | `config` |
| `internal/cobol/copy_expander.go` | 9 | no first-class file role | `analyzer_helper` |
| `internal/parser/metrics.go` | 8 | no first-class file role | `parser_model` |
| `internal/session/error.go` | 6 | no first-class file role | `runtime_model` |
| `internal/resolution/source_site.go` | 4 | no first-class file role | `helper` |
| `internal/scopeir/facts.go` | 4 | no first-class file role | `parser_model` |
| `internal/scopeir/range.go` | 4 | no first-class file role | `parser_model` |
| `internal/session/types.go` | 3 | no first-class file role | `runtime_model` |
| `internal/cli/exit_error.go` | 2 | no first-class file role | `helper` |
| `internal/lbugnative/runner.go` | 1 | no first-class file role | `adapter` |
| `internal/lbugnative/runner_default.go` | 1 | no first-class file role | `fallback_adapter` |

## B4 - Role Coverage Foundation

Status: met as taxonomy foundation, not file-group closure

Source evidence: E8.

| Metric | Unit | Baseline | Latest | Target after implementation |
|---|---:|---:|---:|---:|
| Raw-only files with known first-class role | files | 0 | 17 | 17 |
| Raw-only files with unknown first-class role | files | 17 | 0 | 0 |
| Raw-only files preserving raw/default unresolved separation | files | 17 | 17 | 17 |
| Raw-only files incorrectly counted as production unresolved | files | 0 | 0 | 0 |
| Raw-only raw source sites | source sites | 376 | 376 | preserve |
| Raw-only non-actionable source sites | source sites | 376 | 376 | preserve |

## B5 - Web UI Consumer Coverage Target

Status: met for role and group display

Source evidence: E5.

| Web surface | Unit | Baseline role support | Latest | Target after implementation |
|---|---:|---:|---:|---:|
| `FileMapPanel` role display | covered surface | 0 | 1 | 1 |
| `FileDetailPanel` role display | covered surface | 0 | 1 | 1 |
| `FileTreePanel` inspected for role impact | inspected surface | 0 | 1 | 1 |
| Generated Web `FileSummary` includes role field | generated contract | 0 | 1 | 1 |
| Generated Web `FileSummary` includes group field | generated contract | 0 | 1 | 1 |
| Web path-pattern role inference sites | sites | 0 | 0 | 0 |
| Web path-pattern group inference sites | sites | 0 | 0 | 0 |
| Web tests for missing/unknown role behavior | test cases | 0 | 2 | at least 1 |
| Web/e2e test coverage when visible UI changes | e2e test record | 0 | 1 | at least 1 |
| Browser or screenshot validation for visible UI changes | validation record | 0 | 1 | supplemental |

## B6 - Backend Support/Model/Helper File Group Target

Status: measured after corrected implementation and refreshed after post-review fix

Source evidence: E10, E11, E12.

| Metric | Unit | Baseline before corrected group implementation | E11 closure-time latest | E12 current latest | Target after corrected implementation |
|---|---:|---:|---:|---:|---:|
| `backend_support_model_helper` full group files | files | 0 | 42 | 47 | measured from backend identity rules |
| `backend_support_model_helper` current file-projection unresolved source sites | source sites | 0 | n/a | 2531 | measured from current backend identity rules |
| `backend_support_model_helper` closure-time default unresolved source sites | source sites | 0 | 1073 | historical | measured at E11 closure time |
| `backend_support_model_helper` closure-time raw unresolved source sites | source sites | 0 | 2087 | historical | measured at E11 closure time |
| Anchor sample files assigned `backend_support_model_helper` | files | 0 | 17 | 17 | 17 |
| Anchor sample files missing `backend_support_model_helper` | files | 17 | 0 | 0 | 0 |
| Anchor sample default unresolved source sites | source sites | 0 | 0 | historical | 0 |
| Anchor sample raw unresolved source sites | source sites | 376 | 376 | historical | 376 |
| Analyze direct group line coverage | output line | 0 | 1 | 1 | 1 |
| CLI/API group summary coverage | covered surface | 0 | 3 | 3 | at least 1 |
| Web group label coverage | covered surface | 0 | 2 | 2 | at least 2 |
| Graph File-node `fileGroup` coverage for `internal/repo/runtime_config.go` | graph property | 0 | 0 | 1 | 1 |
| Self-contained e2e command coverage | e2e command | 0 | 0 | 1 | 1 |
| Web path-pattern group inference sites | sites | 0 | 0 | 0 | 0 |

Anchor sample role breakdown target:

| File role | Unit | Baseline assigned to group | Latest assigned to group | Target assigned to group |
|---|---:|---:|---:|---:|
| `analyzer_helper` | files | 0 | 2 | 2 |
| `helper` | files | 0 | 3 | 3 |
| `contract_model` | files | 0 | 1 | 1 |
| `storage_helper` | files | 0 | 1 | 1 |
| `test_helper` | files | 0 | 1 | 1 |
| `config` | files | 0 | 2 | 2 |
| `parser_model` | files | 0 | 3 | 3 |
| `runtime_model` | files | 0 | 2 | 2 |
| `adapter` | files | 0 | 1 | 1 |
| `fallback_adapter` | files | 0 | 1 | 1 |
