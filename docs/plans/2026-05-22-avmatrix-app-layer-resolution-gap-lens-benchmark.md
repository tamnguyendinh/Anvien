# AVmatrix App Layer Resolution Gap Lens Benchmark Ledger

Date: 2026-05-22

Status: planned

Plan: [2026-05-22-avmatrix-app-layer-resolution-gap-lens-plan.md](2026-05-22-avmatrix-app-layer-resolution-gap-lens-plan.md)

Evidence: [2026-05-22-avmatrix-app-layer-resolution-gap-lens-evidence.md](2026-05-22-avmatrix-app-layer-resolution-gap-lens-evidence.md)

## B0 - Baseline Analyze Inventory

Status: pending

Record after P0-B.

Planning audit snapshot, not P0 closure baseline:

| Metric | Planning audit value |
| --- | ---: |
| Files scanned | 728 |
| Files parsed | 539 |
| Unsupported files | 189 |
| Failed files | 0 |
| Graph nodes | 22095 |
| Graph relationships | 54772 |
| Graph path | `.avmatrix/graph.json` |

| Metric | Value |
| --- | ---: |
| Files scanned | pending |
| Files parsed | pending |
| Unsupported files | pending |
| Failed files | pending |
| Graph nodes | pending |
| Graph relationships | pending |
| Counted semantic relationships | pending |
| Execution flows | pending |
| `unknown_connectivity` nodes | pending |
| Graph timestamp/hash | pending |

## B1 - Discussion Coverage Checklist

Status: pending

Record after P0-A.

| Discussion area | Plan task mapped | Evidence section mapped | Validation mapped |
| --- | --- | --- | --- |
| node type is insufficient | pending | pending | pending |
| graph/API must answer before UI | pending | pending | pending |
| App Layer and BE/API/FE rings | pending | pending | pending |
| non-overlapping mixed categories | pending | pending | pending |
| API first-class layer | pending | pending | pending |
| Functional Area accuracy gate | pending | pending | pending |
| persisted ResolutionGap/UnresolvedSymbol | pending | pending | pending |
| fine-grained gap relations | pending | pending | pending |
| Resolution Health separate from Topology Health | pending | pending | pending |
| query-health command | pending | pending | pending |
| semantic query/context/impact/detect-changes output | pending | pending | pending |
| multi-ring layout and same-color islands | pending | pending | pending |
| no dead-code verdict from unresolved refs alone | pending | pending | pending |
| no timeout/auto optimizer behavior | pending | pending | pending |
| no stale graph fallback | pending | pending | pending |
| no evidence loss for graph-size reasons | pending | pending | pending |
| user-facing naming consistency | pending | pending | pending |

## B2 - Baseline Unresolved Inventory

Status: pending

Record during P0-D.

| Metric | Value |
| --- | ---: |
| Unresolved buckets | pending |
| Unresolved occurrences | pending |
| Discussion observed unresolved buckets | about 8880 |
| Discussion observed unresolved occurrences | about 51232 |
| Call gaps | pending |
| Access gaps | pending |
| Type-reference gaps | pending |
| Heritage gaps | pending |
| Builtin/predeclared classified | pending |
| Standard-library classified | pending |
| Test-framework classified | pending |
| External classified | pending |
| In-repo analyzer gaps | pending |
| Unclassified/unknown | pending |

Top unresolved targets:

| Rank | Fact family | Target text | Count | Current classification | Source App Layer hypothesis |
| ---: | --- | --- | ---: | --- | --- |
| 1 | pending | pending | pending | pending | pending |
| 2 | pending | pending | pending | pending | pending |
| 3 | pending | pending | pending | pending | pending |
| 4 | pending | pending | pending | pending | pending |
| 5 | pending | pending | pending | pending | pending |
| 6 | pending | pending | pending | pending | pending |
| 7 | pending | pending | pending | pending | pending |
| 8 | pending | pending | pending | pending | pending |
| 9 | pending | pending | pending | pending | pending |
| 10 | pending | pending | pending | pending | pending |

## B3 - Provisional App Layer Inventory

Status: pending

Record during P0-E. These are sizing measurements only until P1 persists product semantics.

| Provisional App Layer | Node count | Evidence confidence | Notes |
| --- | ---: | --- | --- |
| backend | pending | pending | pending |
| api | pending | pending | pending |
| frontend | pending | pending | pending |
| cli_launcher | pending | pending | pending |
| shared_contract | pending | pending | pending |
| api_contract | pending | pending | pending |
| api_shared_contract | pending | pending | pending |
| frontend_api_client | pending | pending | pending |
| backend_test | pending | pending | pending |
| frontend_test | pending | pending | pending |
| api_test | pending | pending | pending |
| generated_contract | pending | pending | pending |
| docs | pending | pending | pending |
| config | pending | pending | pending |
| generated | pending | pending | pending |
| mixed | pending | pending | pending |
| unknown | pending | pending | pending |

## B4 - Target App Layer Persistence Metrics

Status: pending

Record after P1.

| Metric | Before | Target / expected direction | After |
| --- | ---: | --- | ---: |
| Nodes with persisted App Layer | pending | all classifiable graph nodes | pending |
| Nodes with exactly one primary App Layer | pending | all graph nodes with field | pending |
| Nodes classified as `api` | pending | nonzero if API code exists | pending |
| Nodes classified as API contract/client/shared categories | pending | nonzero if such code exists | pending |
| Nodes classified as test/doc/config/generated/mixed categories | pending | nonzero where source exists | pending |
| Nodes left `unknown` | pending | only where evidence is insufficient | pending |
| Generated contract exposes App Layer enum/fields | pending | yes | pending |
| Missing metadata treated as stale/incomplete schema | pending | yes | pending |
| Load-time App Layer heuristic fallback count | pending | 0 | pending |
| User-facing naming labels defined | pending | yes | pending |

## B5 - Functional Area Metrics

Status: pending

Record after P2.

| Functional Area | Node count | Evidence rule | Unknown/rejected notes |
| --- | ---: | --- | --- |
| resolution | pending | pending | pending |
| graph_health | pending | pending | pending |
| query | pending | pending | pending |
| mcp | pending | pending | pending |
| web_graph_ui | pending | pending | pending |
| layout | pending | pending | pending |
| contracts | pending | pending | pending |
| providers | pending | pending | pending |
| runtime | pending | pending | pending |
| analyzer | pending | pending | pending |
| session | pending | pending | pending |
| launcher | pending | pending | pending |
| cli | pending | pending | pending |
| reporting | pending | pending | pending |
| unknown | pending | pending | pending |

Rejected candidate signals:

| Signal | Rejection reason | Example |
| --- | --- | --- |
| pending | pending | pending |

## B6 - ResolutionGap Persistence Metrics

Status: pending

Record after P3.

| Metric | Before | Target / expected direction | After |
| --- | ---: | --- | ---: |
| Persisted ResolutionGap/UnresolvedSymbol entities | pending | source-backed unresolved buckets or accurate aggregate entities | pending |
| Persisted gap relationships/typed relations | pending | nonzero where source evidence supports them | pending |
| Gaps preserving source node ID | pending | all source-backed gaps | pending |
| Gaps preserving target text | pending | all source-backed gaps | pending |
| Gaps preserving source App Layer | pending | all source-backed gaps with classified source | pending |
| Gaps preserving source Functional Area | pending | all source-backed gaps with classified source | pending |
| Gaps preserving actionability | pending | all modeled gaps | pending |
| Fake resolved target nodes created | pending | 0 | pending |
| Aggregates preserving exact occurrence count | pending | all aggregates | pending |
| Aggregates preserving source samples | pending | all aggregates | pending |
| Evidence capped away for graph-size reasons | pending | 0 | pending |

By fact family:

| Fact family | Count after |
| --- | ---: |
| unresolved call | pending |
| unresolved access | pending |
| unresolved type-reference | pending |
| unresolved heritage | pending |
| external symbol | pending |
| builtin/stdlib/test reference | pending |
| in-repo analyzer gap | pending |
| unknown/unclassified | pending |

By target role:

| Target role | Count after |
| --- | ---: |
| callable | pending |
| member | pending |
| type | pending |
| external | pending |
| builtin | pending |
| test | pending |
| unknown | pending |

By actionability:

| Actionability | Count after |
| --- | ---: |
| non_actionable | pending |
| review | pending |
| analyzer_gap | pending |
| unknown | pending |

## B7 - Resolution Health Inventory Metrics

Status: pending

Record after P4.

| Resolution Health bucket | Count after |
| --- | ---: |
| resolved references | pending |
| unresolved non-actionable | pending |
| external unresolved | pending |
| in-repo analyzer gap | pending |
| unresolved call target | pending |
| unresolved access target | pending |
| unresolved type target | pending |
| unresolved heritage target | pending |
| unclassified/unknown | pending |

Topology plus resolution overlay:

| Topology status | Nodes with no gaps | Nodes with gaps | Nodes with degraded confidence |
| --- | ---: | ---: | ---: |
| connected | pending | pending | pending |
| no_incoming | pending | pending | pending |
| no_outgoing | pending | pending | pending |
| detached_component | pending | pending | pending |
| true_isolated | pending | pending | pending |
| unknown_connectivity | pending | pending | pending |

## B8 - Query Health Metrics

Status: pending

Record before P5 and after P5.

| Intent | Expected core files/symbols | Baseline hit@5 | Baseline hit@10 | Final hit@5 | Final hit@10 | Result |
| --- | --- | ---: | ---: | ---: | ---: | --- |
| unresolved reference diagnostic generation | `internal/resolution/resolve.go`, `internal/resolution/emit.go`, `internal/graphhealth/diagnostics.go` | pending | pending | pending | pending | pending |
| graph health unknown-connectivity separation | `internal/graphhealth/compute.go`, `internal/graphhealth/policy.go`, `avmatrix-web/src/lib/graph-health-filters.ts` | pending | pending | pending | pending | pending |
| App Layer and ResolutionGap layout | Web graph layout code, layout optimizer code, graph filter/lens code | pending | pending | pending | pending | pending |
| runtime reset hidden-terminal behavior | `avmatrix-launcher/src/main.go`, runtime start/reset/stop code | pending | pending | pending | pending | pending |
| API contract surfaces | HTTP API code, contract generation code, generated Web contract code | pending | pending | pending | pending | pending |
| frontend graph filter surfaces | Web graph filter, detail, and layout code | pending | pending | pending | pending | pending |

## B9 - Semantic Command Surface Metrics

Status: pending

Record after P6.

| Command | App Layer shown | Functional Area shown | ResolutionGap shown | Resolution Health shown | Limitation noted |
| --- | --- | --- | --- | --- | --- |
| `query` | pending | pending | pending | pending | pending |
| `context` | pending | pending | pending | pending | pending |
| `impact` | pending | pending | pending | pending | pending |
| `detect-changes` | pending | pending | pending | pending | pending |
| resolution inventory command | pending | pending | pending | pending | pending |
| query-health command | pending | pending | pending | pending | pending |

## B10 - Web UI Ring And Filter Metrics

Status: pending

Record after P7.

| Metric | Value |
| --- | ---: |
| Visible App Layer ring count | pending |
| Backend ring node count | pending |
| API ring node count | pending |
| Frontend ring node count | pending |
| API ring placed between Backend and Frontend | pending |
| Contract rings placed near API when present | pending |
| Shared/API contract ring count | pending |
| Frontend API client ring count | pending |
| Test ring/group count | pending |
| Docs/Config/Generated ring count | pending |
| Unknown/Mixed ring count | pending |
| Ring size/spacing policy recorded | pending |
| Default visible ring/lens count | pending |
| Default hidden/collapsed ring/lens count | pending |
| Node type islands visible | pending |
| Same-color island violations | pending |
| ResolutionGap visible count | pending |
| App Layer filters available | pending |
| Resolution Health filters available | pending |
| Optimizer auto-run events after render/load/filter | pending |

## B11 - Validation Outputs

Status: pending

Build/test/e2e timings are validation evidence, not product performance benchmarks unless this plan changes those systems.

| Validation | Command | Result | Notes |
| --- | --- | --- | --- |
| Full build | `powershell -ExecutionPolicy Bypass -File avmatrix-launcher\build.ps1` | pending | pending |
| Backend tests | pending | pending | pending |
| Contract generation/check | pending | pending | pending |
| Web unit tests | pending | pending | pending |
| Web e2e tests | pending | pending | pending |
| Query-health command | pending | pending | pending |
| Resolution inventory command | pending | pending | pending |
| `query` semantic output | pending | pending | pending |
| `context` semantic output | pending | pending | pending |
| `impact` semantic output | pending | pending | pending |
| `detect-changes` semantic output | pending | pending | pending |
| AVmatrix detect-changes for implementation commits | pending | pending | pending |
