# AVmatrix Graph Health Unknown Connectivity Separation Benchmark Ledger

Date: 2026-05-21

Status: draft

Plan: [2026-05-21-avmatrix-graph-health-unknown-connectivity-separation-plan.md](2026-05-21-avmatrix-graph-health-unknown-connectivity-separation-plan.md)

Evidence: [2026-05-21-avmatrix-graph-health-unknown-connectivity-separation-evidence.md](2026-05-21-avmatrix-graph-health-unknown-connectivity-separation-evidence.md)

## B0 - Current Baseline

Status: recorded

Source graph:

| Metric | Value |
| --- | ---: |
| Repository | `E:\AVmatrix-GO` |
| Analyze command | `avmatrix analyze --force` |
| Files scanned | 721 |
| Files parsed | 539 |
| Unsupported files | 182 |
| Failed files | 0 |
| Graph nodes | 21941 |
| Graph relationships | 54489 |
| Counted relationships under Graph Health policy | 26841 |
| Nodes with `graphHealthDiagnostics` | 4345 |

Current diagnostic nodes by counted-edge bucket:

| Bucket | Count | Interpretation |
| --- | ---: | --- |
| `connected_in_out` | 2036 | Connected nodes currently forced into `Unknown` by diagnostics |
| `no_incoming` | 1581 | Possible unwired/dead-code candidates, but still need overlays |
| `no_outgoing` | 559 | Often leaf behavior, not a defect by default |
| `true_isolated_by_counted_edges` | 169 | Stronger triage candidate, still not a deletion verdict |

Current diagnostic nodes by semantic label:

| Label | Count |
| --- | ---: |
| Function | 3057 |
| Method | 738 |
| Struct | 292 |
| Variable | 166 |
| File | 22 |
| Property | 21 |
| Package | 18 |
| Const | 10 |
| Interface | 7 |
| Constructor | 6 |
| Class | 4 |
| TypeAlias | 4 |

Current diagnostic nodes by path bucket:

| Path bucket | Count |
| --- | ---: |
| Go source | 2577 |
| Test | 1627 |
| Web source | 141 |

## B1 - Current Top Unresolved Targets

Status: recorded

| Rank | Target | Count | Initial classification hypothesis |
| ---: | --- | ---: | --- |
| 1 | `type-reference:testing.T` | 1123 | Go stdlib/test framework |
| 2 | `type-reference:collector` | 421 | needs investigation |
| 3 | `call:t.Helper` | 342 | Go stdlib/test framework method |
| 4 | `type-reference:int` | 336 | Go predeclared type |
| 5 | `call:make` | 293 | Go builtin |
| 6 | `call:string` | 196 | Go predeclared conversion/type |
| 7 | `call:t.TempDir` | 191 | Go stdlib/test framework method |
| 8 | `call:len` | 180 | Go builtin |
| 9 | `call:node.Kind` | 132 | member access/call resolution gap candidate |
| 10 | `call:c.text` | 111 | member access/call resolution gap candidate |
| 11 | `call:append` | 108 | Go builtin |
| 12 | `access:time.Second` | 104 | Go stdlib |
| 13 | `call:t.Fatalf` | 100 | Go stdlib/test framework method |
| 14 | `call:strings.TrimSpace` | 89 | Go stdlib |
| 15 | `call:uint` | 88 | Go predeclared conversion/type |
| 16 | `type-reference:Server` | 78 | needs investigation |
| 17 | `type-reference:context.Context` | 72 | Go stdlib |
| 18 | `type-reference:map[string]any` | 51 | Go composite/predeclared type |
| 19 | `call:int` | 49 | Go predeclared conversion/type |
| 20 | `type-reference:byte` | 48 | Go predeclared alias |
| 21 | `type-reference:testing.B` | 48 | Go stdlib/test framework |
| 22 | `call:filepath.Join` | 47 | Go stdlib |
| 23 | `type-reference:float64` | 45 | Go predeclared type |
| 24 | `access:result.Graph` | 44 | member access resolution gap candidate |
| 25 | `access:result.Metrics` | 42 | member access resolution gap candidate |

## B2 - Target Post-P1 Metrics

Status: pending

After topology and diagnostics are separated, record:

| Metric | Before | Target / expected direction | After |
| --- | ---: | --- | ---: |
| `unknown_connectivity` topology count | 4345 effective observed bucket | should drop sharply; only truly unclassifiable topology remains | pending |
| Nodes with diagnostics | 4345 | may remain high before P2; diagnostics are not hidden | pending |
| Connected nodes with diagnostics | 2036 | should be `connected` with `confidence: unknown` | pending |
| No-incoming nodes with diagnostics | 1581 | should be `no_incoming` with `confidence: unknown` | pending |
| No-outgoing nodes with diagnostics | 559 | should be `no_outgoing` with `confidence: unknown` | pending |
| True-isolated nodes with diagnostics | 169 | should be `true_isolated` with `confidence: unknown` | pending |
| Valid graph nodes emitted as `unknown_connectivity` | 4345 effective observed bucket | `0` after P1 unless malformed/incomplete input is intentionally tested | pending |

## B3 - Target Post-P2 Diagnostic Classification Metrics

Status: pending

Record classification counts after builtin/stdlib/external/test references are classified.

Required diagnostic metadata:

| Field | Allowed values |
| --- | --- |
| `classification` | `builtin`, `standard_library`, `test_framework`, `external_library`, `in_repo_unresolved`, `unclassified` |
| `actionability` | `non_actionable`, `review`, `analyzer_gap` |

| Diagnostic classification | Before | Expected actionability | After |
| --- | ---: | --- | ---: |
| builtin | pending exact classification | `non_actionable` | pending |
| standard_library | pending exact classification | `non_actionable` | pending |
| test_framework | pending exact classification | `non_actionable` | pending |
| external_library | pending exact classification | `review` | pending |
| in_repo_unresolved | pending exact classification | `analyzer_gap` | pending |
| unclassified | pending exact classification | `review` | pending |

Required before/after target rows:

| Target | Before | After |
| --- | ---: | ---: |
| `testing.T` | 1123 | pending |
| `make` | 293 | pending |
| `string` | 196 | pending |
| `t.TempDir` | 191 | pending |
| `len` | 180 | pending |
| `append` | 108 | pending |
| `time.Second` | 104 | pending |
| `t.Fatalf` | 100 | pending |
| `strings.TrimSpace` | 89 | pending |
| `context.Context` | 72 | pending |
| `map[string]any` | 51 | pending |
| `filepath.Join` | 47 | pending |

## B4 - Web UI Benchmarks

Status: pending

Record after Web UI validation:

| UX check | Expected | Result |
| --- | --- | --- |
| `Unknown` topology count no longer represents all diagnostics | yes | pending |
| Hiding `Unknown` topology keeps connected diagnostic nodes visible | yes | pending |
| Diagnostic filter remains separate from topology filter | yes | pending |
| Node detail shows topology and diagnostics together | yes | pending |
| Graph Health wording does not imply `Unknown` means dead code | yes | pending |
| Diagnostic classification/actionability is visible where needed for triage | yes | pending |
| Connected diagnostic nodes are not presented as topology defects | yes | pending |

## B5 - Validation Benchmarks

Status: pending

Record commands and results:

| Validation | Expected | Result |
| --- | --- | --- |
| Full packaged build before tests | pass | pending |
| Focused graphhealth tests | pass | pending |
| Focused resolution diagnostic tests | pass | pending |
| Focused HTTP graph/report/explain tests | pass | pending |
| Generated contract tests if touched | pass | pending |
| Full relevant Go tests | pass | pending |
| Focused Web unit tests | pass | pending |
| Full Web unit tests | pass | pending |
| Web e2e Graph Health topology/diagnostic separation | pass | pending |
| Final graph refresh and inventory count | recorded | pending |
| Required change detection before commit | recorded | pending |
