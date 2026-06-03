# Anvien File Role Classification Gap Benchmark Ledger

Date: 2026-06-03

Status: Open

Companion files:

- Plan: [2026-06-03-anvien-file-role-classification-gap-plan.md](2026-06-03-anvien-file-role-classification-gap-plan.md)
- Evidence ledger: [2026-06-03-anvien-file-role-classification-gap-evidence.md](2026-06-03-anvien-file-role-classification-gap-evidence.md)

## Benchmark Rules

1. Record quantitative data only.
2. Put command interpretation in the evidence ledger.
3. Preserve current, target, and delta where useful.
4. Track analyze inventory counts after every implementation slice.
5. Track file-role coverage counts after every implementation slice.
6. Build/test pass/fail belongs in the evidence ledger unless timing/count/size is the measured target.

## B0 - Analyze And File Projection Baseline

Status: recorded

Source evidence: E0.

| Metric | Unit | Baseline | Latest | Target |
|---|---:|---:|---:|---:|
| Files scanned | files | 821 | 824 | record |
| Parsed code files | files | 601 | 601 | no unintended decrease |
| Failed files | files | 0 | 0 | 0 |
| Graph nodes | nodes | 60974 | 61008 | record |
| Graph relationships | relationships | 96624 | 96664 | record |
| File projection files | files | 821 | 824 | record |
| File projection dependency edges | edges | 15965 | 15971 | record |
| Default-visible unresolved files | files | 336 | 336 | preserve unless separate behavior changes |
| Raw unresolved files | files | 353 | 353 | preserve unless separate behavior changes |
| Raw-only unresolved files | files | 17 | 17 | 17 recognized, 0 unknown role |

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

## B4 - Role Coverage Target

Status: target defined

| Metric | Unit | Baseline | Target after implementation |
|---|---:|---:|---:|
| Raw-only files with known first-class role | files | 0 | 17 |
| Raw-only files with unknown first-class role | files | 17 | 0 |
| Raw-only files preserving raw/default unresolved separation | files | 17 | 17 |
| Raw-only files incorrectly counted as production unresolved | files | 0 | 0 |

## B5 - Web UI Consumer Coverage Target

Status: target defined

Source evidence: E5.

| Web surface | Unit | Baseline role support | Target after implementation |
|---|---:|---:|---:|
| `FileMapPanel` role display | covered surface | 0 | 1 |
| `FileDetailPanel` role display | covered surface | 0 | 1 |
| `FileTreePanel` inspected for role impact | inspected surface | 0 | 1 |
| Generated Web `FileSummary` includes role field | generated contract | 0 | 1 |
| Web path-pattern role inference sites | sites | 0 | 0 |
| Web tests for missing/unknown role behavior | test cases | 0 | at least 1 |
| Web/e2e test coverage when visible UI changes | e2e test record | 0 | at least 1 if UI changes |
| Browser or screenshot validation for visible UI changes | validation record | 0 | supplemental when UI changes |
