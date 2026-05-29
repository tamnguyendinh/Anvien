# Anvien Graph Health Unknown Connectivity Separation Benchmark Ledger

Date: 2026-05-21

Status: implemented

Plan: [2026-05-21-anvien-graph-health-unknown-connectivity-separation-plan.md](2026-05-21-anvien-graph-health-unknown-connectivity-separation-plan.md)

Evidence: [2026-05-21-anvien-graph-health-unknown-connectivity-separation-evidence.md](2026-05-21-anvien-graph-health-unknown-connectivity-separation-evidence.md)

## B0 - Current Baseline

Status: recorded

Source graph:

| Metric | Value |
| --- | ---: |
| Repository | `E:\Anvien` |
| Analyze command | `anvien analyze --force` |
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

Status: recorded

After topology and diagnostics are separated, record:

| Metric | Before | Target / expected direction | After |
| --- | ---: | --- | ---: |
| `unknown_connectivity` topology count for valid graph nodes | 4345 effective observed bucket | `0` after P1; only malformed/incomplete inputs may use this status | `0` |
| Nodes with diagnostics | 4345 diagnostic-bearing raw nodes | diagnostics are not hidden | diagnostics retained and counted as `51232` occurrences |
| Connected nodes with diagnostics | 2036 | should be `connected` with `confidence: unknown` | `2633` connected diagnostic nodes |
| No-incoming nodes with diagnostics | 1581 | should be `no_incoming` with `confidence: unknown` | preserved by topology; no longer overwritten by diagnostics |
| No-outgoing nodes with diagnostics | 559 | should be `no_outgoing` with `confidence: unknown` | preserved by topology; no longer overwritten by diagnostics |
| True-isolated nodes with diagnostics | 169 | should be `true_isolated` with `confidence: unknown` | preserved by topology; no longer overwritten by diagnostics |
| Valid graph nodes emitted as `unknown_connectivity` | 4345 effective observed bucket | `0` after P1 unless malformed/incomplete input is intentionally tested | `0` |

## B3 - Target Post-P2 Diagnostic Classification Metrics

Status: recorded

Record classification counts after builtin/stdlib/external/test references are classified.

Required diagnostic metadata:

| Field | Allowed values |
| --- | --- |
| `classification` | `builtin`, `standard_library`, `test_framework`, `external_library`, `in_repo_unresolved`, `unclassified` |
| `actionability` | `non_actionable`, `review`, `analyzer_gap` |

| Diagnostic classification | Before | Expected actionability | After |
| --- | ---: | --- | ---: |
| builtin | pending exact classification | `non_actionable` | `8003` |
| standard_library | pending exact classification | `non_actionable` | `7206` |
| test_framework | pending exact classification | `non_actionable` | `8400` |
| external_library | pending exact classification | `review` | recorded in API payload, not top final bucket |
| in_repo_unresolved | pending exact classification | `analyzer_gap` | `27461` |
| unclassified | pending exact classification | `review` | recorded in API payload, not top final bucket |

Required before/after target rows:

| Target | Before | After |
| --- | ---: | ---: |
| `testing.T` | 1123 | classified as `test_framework`, `non_actionable` |
| `make` | 293 | classified as `builtin`, `non_actionable` |
| `string` | 196 | classified as `builtin`, `non_actionable` |
| `t.TempDir` | 191 | classified as `test_framework`, `non_actionable` |
| `len` | 180 | classified as `builtin`, `non_actionable` |
| `append` | 108 | classified as `builtin`, `non_actionable` |
| `time.Second` | 104 | classified as `standard_library`, `non_actionable` |
| `t.Fatalf` | 100 | classified as `test_framework`, `non_actionable` |
| `strings.TrimSpace` | 89 | classified as `standard_library`, `non_actionable` |
| `context.Context` | 72 | classified as `standard_library`, `non_actionable` |
| `map[string]any` | 51 | classified as `builtin`, `non_actionable` |
| `filepath.Join` | 47 | classified as `standard_library`, `non_actionable` |

## B4 - Web UI Benchmarks

Status: recorded

Record after Web UI validation:

| UX check | Expected | Result |
| --- | --- | --- |
| `Unknown` topology count no longer represents all diagnostics | yes | `0` final valid graph nodes |
| Hiding `Unknown` topology keeps connected diagnostic nodes visible | yes | e2e `hiding unknown topology keeps connected diagnostic nodes selectable` passed |
| Diagnostic filter remains separate from topology filter | yes | unit and e2e coverage passed |
| Node detail shows topology and diagnostics together | yes | unit and e2e coverage passed |
| Graph Health wording does not imply `Unknown` means dead code | yes | wording updated |
| Diagnostic classification/actionability is visible where needed for triage | yes | node detail shows classification/actionability |
| Connected diagnostic nodes are not presented as topology defects | yes | report test passed |
| Report candidates expose `triageDimension` | yes; `topology` or `diagnostic` | contract/API tests passed |
| Connected diagnostic report candidates use diagnostic dimension only | yes | `TestGraphHealthReportSeparatesTopologyAndDiagnosticTriage` passed |

## B5 - Validation Benchmarks

Status: recorded

Record commands and results:

| Validation | Expected | Result |
| --- | --- | --- |
| Full packaged build before tests | pass | passed |
| Focused graphhealth tests | pass | passed |
| Focused resolution diagnostic tests | pass | passed |
| Focused HTTP graph/report/explain tests | pass | passed |
| Generated contract tests if touched | pass | passed |
| Generated Web contract output is current if contract fields change | pass | passed |
| Full relevant Go tests | pass | `go test ./cmd/... ./internal/...` passed |
| Focused Web unit tests | pass | passed |
| Full Web unit tests | pass | `44` files, `357` tests passed |
| Web e2e Graph Health topology/diagnostic separation | pass | focused spec passed; full e2e `36` passed, `8` skipped |
| Final graph refresh and inventory count | recorded | `nodes=22010`, `relationships=54679`, `unknown_connectivity=0` |
| Required change detection before commit | recorded | `risk_level=critical`, `changed_files=17`, `affected_count=35`; affected paths match intended Graph Health/API/contract/Web work |
