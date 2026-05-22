# AVmatrix App Layer Resolution Gap Lens Benchmark Ledger

Date: 2026-05-22

Status: in progress; Phase 0 closure audit complete; Phase 2 complete; Phase 2A low-confidence global CALLS fallback, source-site metadata persistence, source-site accuracy command, File-source CALLS gate, golden corpus, and source-site accuracy golden fixture command slices complete

Plan: [2026-05-22-avmatrix-app-layer-resolution-gap-lens-plan.md](2026-05-22-avmatrix-app-layer-resolution-gap-lens-plan.md)

Evidence: [2026-05-22-avmatrix-app-layer-resolution-gap-lens-evidence.md](2026-05-22-avmatrix-app-layer-resolution-gap-lens-evidence.md)

## B0 - Baseline Analyze Inventory

Status: complete for Phase 0 closure audit; later implementation slices continue to record changed metrics.

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
| Files scanned | 736 |
| Files parsed | 547 |
| Unsupported files | 189 |
| Failed files | 0 |
| Graph nodes | 22635 |
| Graph relationships | 52144 |
| Counted semantic relationships | 23437 |
| Execution flows | 645 |
| `unknown_connectivity` nodes | 0 |
| Graph timestamp/hash | 2026-05-22 13:44:46 +07:00 / `DB28BF1D99D0CFEEC860840AE3921A878DA7B20086481F9139C437D9112F9432` |
| Analyze semantic enrichment phase present | yes |
| Semantic enrichment latency | 58102700 ns |
| Semantic enrichment memory delta | start alloc 1195152 bytes; end alloc 151258008 bytes; max observed sys 536672504 bytes |
| Semantic enrichment graph-size delta | 22635 nodes / 52144 relationships after analyze |
| LadybugDB semantic export rows | 22635 node rows / 52144 relationship rows; fallback inserts 0; skipped relationships 0 |
| Raw unresolved facts captured before aggregation | 58195 unresolved diagnostics; 57449 buckets |
| Raw call/access source sites captured before resolved-edge emission | 43163 call source-site occurrences; 24222 access source-site occurrences |
| Semantic enrichment duplicate full-graph scans | one relationship scan recorded by semantic metrics: 52144 relationships scanned |

Phase 0 command:

```powershell
.\avmatrix-launcher\server-bundle\avmatrix.exe analyze --force --benchmark-json .tmp\2026-05-22-phase0-baseline.json --benchmark-label phase0-baseline
```

## B1 - Discussion Coverage Checklist

Status: complete for Phase 0 closure audit.

Record after P0-A.

| Discussion area | Plan task mapped | Evidence section mapped | Validation mapped |
| --- | --- | --- | --- |
| node type is insufficient | Problem, P1, P2, P3 | E1 | P0-A |
| graph/API must answer before UI | Problem, P1, P3, P6, P7 | E1, E3 | P0-A/P0-G |
| App Layer and BE/API/FE rings | P1, P7 | E1, E4 | P0-A/P0-E |
| non-overlapping mixed categories | P1-A/P1-B | E1, E4 | P0-A/P0-E |
| API first-class layer | P1, P6 | E1, E4 | P0-A/P0-E |
| Functional Area accuracy gate | P2 | E1, E3, E4 | P0-A/P0-G |
| Proof-based CALLS/ACCESSES and source-site inventory | Phase 2A | E1, E3, E4 | P0-A/P0-C/P0-D |
| persisted ResolutionGap/UnresolvedSymbol | P3 | E1, E3 | P0-A/P0-G |
| fine-grained gap relations | P3-C | E1, E4 | P0-A/P0-D |
| Resolution Health separate from Topology Health | P3/P4 | E1, E4 | P0-A/P0-D |
| query-health command | P5 | E1, E5 | P0-A/P0-F |
| semantic query/context/impact/detect-changes output | P6 | E1, E3, E5 | P0-A/P0-F/P0-G |
| API-specific MCP semantic output | P6 | E1, E3, E5 | P0-A/P0-F/P0-G |
| multi-ring layout and same-color islands | P7 | E1, E3, E5 | P0-A/P0-F/P0-G |
| no dead-code verdict from unresolved refs alone | Scope Boundary, P3/P4 | E1, E4 | P0-A/P0-D |
| no timeout/auto optimizer behavior | Rules of plan, P7 | E1, E3 | P0-A/P0-G |
| no stale graph fallback | Rules of plan, P1/P2/P3 | E1, E2 | P0-A/P0-B |
| no evidence loss for graph-size reasons | Rules of plan, P3/P5 | E1, E4 | P0-A/P0-D |
| user-facing naming consistency | P1-H, P3, P7 | E1, E3 | P0-A/P0-G |

## B2 - Baseline Unresolved Inventory

Status: complete for Phase 0 closure audit.

Record during P0-D.

| Metric | Value |
| --- | ---: |
| Unresolved buckets | 57449 |
| Unresolved occurrences | 58195 |
| Discussion observed unresolved buckets | about 8880 |
| Discussion observed unresolved occurrences | about 51232 |
| Call gaps | 30946 |
| Access gaps | 17935 |
| Type-reference gaps | 9307 |
| Heritage gaps | 7 |
| Builtin/predeclared classified | 9851 |
| Standard-library classified | 7339 |
| Test-framework classified | 7094 |
| External classified | 158 |
| In-repo analyzer gaps | 33753 |
| Unclassified/unknown | 0 |

Top unresolved targets:

| Rank | Fact family | Target text | Count | Current classification | Source App Layer hypothesis |
| ---: | --- | --- | ---: | --- | --- |
| 1 | call | `t.Fatalf` | 3410 | test_framework / non_actionable | backend_test |
| 2 | type-reference | `testing.T` | 2452 | test_framework / non_actionable | backend_test |
| 3 | call | `len` | 1955 | builtin / non_actionable | cli_launcher |
| 4 | call | `string` | 1561 | builtin / non_actionable | cli_launcher |
| 5 | call | `append` | 1198 | builtin / non_actionable | cli_launcher |
| 6 | type-reference | `int` | 1075 | builtin / non_actionable | cli_launcher |
| 7 | call | `expect` | 870 | in_repo_unresolved / analyzer_gap | frontend_test |
| 8 | call | `make` | 652 | builtin / non_actionable | cli_launcher |
| 9 | call | `strings.Contains` | 587 | standard_library / non_actionable | cli_launcher |
| 10 | access | `result.Metrics` | 575 | in_repo_unresolved / analyzer_gap | backend |

## B3 - Provisional App Layer Inventory

Status: complete for Phase 0 closure audit.

Record during P0-E. The plan already implemented Phase 1 before this closure audit, so these are fresh persisted App Layer graph counts rather than pre-implementation sizing-only counts.

| Provisional App Layer | Node count | Evidence confidence | Notes |
| --- | ---: | --- | --- |
| backend | 9937 | high | Backend source and inferred backend-owned Process/Community nodes. |
| api | 2010 | high | `internal/httpapi`, `internal/mcp`, and API surfaces. |
| frontend | 1859 | high | `avmatrix-web/src/**` excluding more specific frontend API client/test rules. |
| cli_launcher | 256 | high | `avmatrix-launcher/**`. |
| shared_contract | 0 | high | No current nodes assigned to this standalone category in fresh graph. |
| api_contract | 161 | high | Contract source/generator surfaces. |
| api_shared_contract | 0 | high | No current nodes assigned to this mixed category in fresh graph. |
| frontend_api_client | 181 | high | `avmatrix-web/src/services/backend-client.ts` and owned symbols. |
| backend_test | 4457 | high | Go/backend tests. |
| frontend_test | 620 | high | Web unit/e2e/manual test roots. |
| api_test | 1107 | high | API/MCP/contract tests. |
| generated_contract | 16 | high | Generated contract artifacts. |
| docs | 1604 | high | Markdown/docs/reporting paths. |
| config | 37 | high | Known config/package/build files covered by rules. |
| generated | 0 | high | No current generic generated nodes assigned outside generated contracts. |
| mixed | 364 | relationship-backed | Process/Community nodes spanning multiple non-unknown App Layers. |
| unknown | 26 | explicit insufficient evidence | Examples include Dockerfiles, `avmatrix-web/vercel.json`, and baseline proof JSON files. |

## B4 - Target App Layer Persistence Metrics

Status: in progress; App Layer persistence slice measured after locally built CLI analyze.

Record after P1.

| Metric | Before | Target / expected direction | After |
| --- | ---: | --- | ---: |
| Nodes with persisted App Layer | 0 | all classifiable graph nodes | 22239 |
| Nodes with exactly one primary App Layer | 0 | all graph nodes with field | 22239 |
| Nodes classified as `api` | pending | nonzero if API code exists | 2013 |
| Nodes classified as API contract/client/shared categories | pending | nonzero if such code exists | api_contract 155, frontend_api_client 182, shared_contract 0, api_shared_contract 0 |
| Nodes classified as test/doc/config/generated/mixed categories | pending | nonzero where source exists | backend_test 4415, frontend_test 620, api_test 1102, docs 1601, config 37, generated_contract 17, generated 0, mixed 399 |
| Nodes left `unknown` | pending | only where evidence is insufficient | 26 |
| Generated contract exposes App Layer enum/fields | pending | yes | yes |
| Missing metadata treated as stale/incomplete schema | pending | yes | yes |
| Load-time App Layer heuristic fallback count | pending | 0 | 0 |
| User-facing naming labels defined | pending | yes | yes |

App Layer count snapshot from `.avmatrix/graph.json` after `.\avmatrix-launcher\server-bundle\avmatrix.exe analyze --force --benchmark-json .tmp\2026-05-22-app-layer-semantic-enrichment.json --benchmark-label app-layer-semantic-enrichment`:

| App Layer | Node count |
| --- | ---: |
| backend | 9554 |
| api | 2013 |
| frontend | 1862 |
| cli_launcher | 256 |
| shared_contract | 0 |
| api_contract | 155 |
| api_shared_contract | 0 |
| frontend_api_client | 182 |
| backend_test | 4415 |
| frontend_test | 620 |
| api_test | 1102 |
| generated_contract | 17 |
| docs | 1601 |
| config | 37 |
| generated | 0 |
| mixed | 399 |
| unknown | 26 |

Public metadata/status snapshot after P1-E/P1-H:

| Metric | Value |
| --- | --- |
| Semantic schema version | `semantic_app_functional_v1` |
| JSON graph payload semantic field | `semanticStatus` |
| NDJSON semantic record | `semantic_status` |
| Stale fixture graph behavior | `stale_incomplete`, 5 missing `appLayer` nodes in `TestGraphReturnsJSONForRegisteredRepo` |
| Fresh explicit unknown behavior | explicit `unknown` is counted as unknown semantic evidence, not stale metadata |
| Generated Web status/types | `SEMANTIC_STATUS_VALUES`, `SEMANTIC_SCHEMA_VERSION`, `GraphSemanticStatus`, `GraphResponse.semanticStatus` |
| Generated Web naming constants | `APP_LAYER_LABELS`, `SEMANTIC_TERMS` |

## B5 - Functional Area Metrics

Status: complete for Phase 2.

Measured after `.\avmatrix-launcher\server-bundle\avmatrix.exe analyze --force --benchmark-json .tmp\2026-05-22-functional-area-semantic-enrichment.json --benchmark-label functional-area-semantic-enrichment`.

| Functional Area | Node count | Evidence rule | Unknown/rejected notes |
| --- | ---: | --- | --- |
| analyzer | 2426 | `internal/analyze` and analyzer pipeline paths | high-confidence path rule |
| api | 1583 | `internal/httpapi`, frontend backend-client API surface | high-confidence path rule |
| cli | 943 | `cmd/**` and CLI command paths | high-confidence path rule |
| configuration | 37 | known project/build config filenames | high-confidence path rule |
| contracts | 248 | `contracts/**`, `internal/contracts/**`, generator paths | high-confidence path rule |
| documentation | 1293 | docs and markdown content paths | high-confidence path rule |
| embeddings | 701 | `internal/embeddings` and embedding API paths | high-confidence path rule |
| graph_health | 761 | `internal/graphhealth`, `internal/graphaccuracy` | high-confidence path rule |
| launcher | 243 | `avmatrix-launcher/**` | high-confidence path rule |
| layout | 280 | graph adapter, edge style, links, Sigma/layout paths | high-confidence path rule |
| mcp | 1566 | `internal/mcp/**` | high-confidence path rule |
| mixed | 294 | Process/Community membership spans more than one non-unknown area | relationship-backed inference only |
| providers | 4954 | parser/provider/ScopeIR extraction paths | high-confidence path rule |
| query | 1154 | query/group/tool-command/query-profile paths | high-confidence path rule |
| reporting | 308 | `reports/**` | high-confidence path rule |
| resolution | 1225 | `internal/resolution/**` | high-confidence path rule |
| runtime | 21 | logging/version/runtime management paths | high-confidence path rule |
| session | 510 | `internal/session`, Web chat/session runtime paths | high-confidence path rule |
| storage | 1243 | graph, repo, LadybugDB schema/load paths | high-confidence path rule |
| unknown | 1736 | missing path or no accepted rule | intentionally not guessed |
| web_graph_ui | 832 | Web App/components/state/filter graph UI paths | high-confidence path rule |

Rejected candidate signals:

| Signal | Rejection reason | Example |
| --- | --- | --- |
| Import/call neighborhood | Too easy to contaminate ownership through callers/callees; needs separate evidence before use | a provider called by query code should not become query-owned automatically |
| Community label text | Heuristic labels are useful context but not stable enough as direct classification evidence | community names may be inferred from mixed symbols |
| Process label text | Process membership is accepted for relationship-backed Process/Community inference, but label text alone is not enough | process label wording can drift with ranking/extraction |
| AI-assisted labels | Not reproducible or verifiable in the analyze pipeline today | no persisted deterministic label source exists |
| Explicit semantic config | No repository-owned config exists yet for Functional Area overrides | future phase may add one if needed |

## B5A - Proof-Based CALLS/ACCESSES Accuracy Metrics

Status: in progress; graph-inventory, golden corpus, and source-site accuracy golden fixture command metrics recorded.

Record during Phase 2A.

Discussion benchmark baseline that motivated this gate:

| Metric | AVmatrix | GitNexus | Notes |
| --- | ---: | ---: | --- |
| Raw ACCESSES | 7670 | 27559 | GitNexus over-emits heavily |
| Unique ACCESSES source-target | 7659 | 18236 | GitNexus duplicate count is high |
| Duplicate ACCESSES edges | 11 | 9323 | AVmatrix duplicate behavior is much cleaner |
| Max ACCESSES duplicate | 2 | 60 | GitNexus duplicate outliers are severe |
| ACCESSES target is Property | 7670 / 7670 | 19102 / 27559 | AVmatrix must preserve this semantic contract |
| Raw CALLS | 15109 | 20051 | More edges is not automatically better |
| Unique CALLS source-target | 15109 | 19147 | GitNexus has duplicates and coarse edges |
| Duplicate CALLS edges | 0 | 904 | AVmatrix currently cleaner on duplicate CALLS |
| Duplicate CALLS pairs | 0 | 513 | AVmatrix currently cleaner on duplicate CALLS pairs |
| Max CALLS duplicate | 1 | 20 | GitNexus duplicate outliers are severe |

Phase 2A target metrics:

| Metric | Baseline | Target | After |
| --- | ---: | ---: | ---: |
| Raw call source sites inventoried | pending | all syntactic call sites in supported files | 43163 current graph records: 12217 resolved `CALLS` occurrences + 30946 unresolved call diagnostics |
| Raw access source sites inventoried | pending | all syntactic access sites in supported files | 24222 current graph records: 6287 resolved `ACCESSES` occurrences + 17935 unresolved access diagnostics |
| Source-site records with stable sourceSiteID | pending | all inventoried call/access sites | 84372 all relationship/diagnostic source-site records including call/access/type-reference/heritage; missing sourceSiteID count 0 |
| Resolved `CALLS` edges | 15109 discussion sample | proof-backed only | 7632 current graph edges; 12217 source-site occurrences |
| Resolved `ACCESSES` edges | 7670 discussion sample | property/field proof-backed only | 3297 current graph edges; 6287 source-site occurrences |
| Resolved edges from low-confidence/global fallback | pending | 0 unless accepted proof is present | 0 in focused resolver golden fixtures |
| Low-confidence/global fallback source sites inventoried | pending | explicit count and status distribution | 2159 current graph diagnostics with `proofKind=global-fallback-low-confidence`; 2 focused resolver fixtures recorded as unresolved source-backed diagnostics |
| Unresolved local-binding call sites | pending | explicit count | 30896 current graph unresolved call diagnostics; finer external/ambiguous/dynamic split pending P3/P4 |
| Unresolved external call/access sites | pending | explicit count | 158 current `external_library` diagnostics; golden corpus covers `cobra.Command` as external/review |
| Ambiguous call/access sites | pending | explicit count | 0 explicit `ambiguous` statuses in current graph and golden corpus; richer split remains Phase 3/4 |
| Dynamic call/access sites | pending | explicit count | 0 explicit `dynamic` statuses in current graph and golden corpus; richer split remains Phase 3/4 |
| Unsupported syntax sites | pending | explicit count | 398 current graph diagnostics; golden corpus covers 1 file-level call source as `unsupported_syntax` |
| False resolved edges in golden corpus | pending | 0 | 0 in `TestProofBasedCallAccessGoldenCorpus` across callback variable, `stop()` cross-name, file-source call, and non-property ACCESSES cases |
| Silent missing source sites in golden corpus | pending | 0 | 0 in `TestProofBasedCallAccessGoldenCorpus`; 12 expected call/access source sites are found |
| Source sites hidden by relationship dedupe without occurrence evidence | pending | 0 | 0 in focused duplicate-edge resolver fixture and golden corpus; current graph has 5161 merged relationships with 15862 source-site occurrences preserved |
| Resolved ACCESSES targets with label `Property` | pending | all resolved ACCESSES unless split relation says otherwise | 3297 / 3297 current graph |
| Resolved ACCESSES targets with labels `Variable`/`Const`/`Static` | pending | 0 or moved to separate non-ACCESSES relation/fact role | 0 current graph; access-candidate audit bucket `non_property_target` covers rejected selector targets |
| Resolved ACCESSES targets with labels `Function`/`Method`/other | pending | 0 | 0 current graph |
| Non-property resolved ACCESSES targets in golden corpus | pending | 0 | 0 in `TestProofBasedCallAccessGoldenCorpus` |
| Duplicate resolved CALLS pairs | 0 discussion sample | 0 unless source-site evidence proves separate occurrences | 10 duplicate source-target pairs across current `CALLS`/`ACCESSES`; max duplicate 2; exact source-site occurrence evidence preserved |
| Duplicate resolved ACCESSES pairs | 11 discussion sample | expected duplicates documented by source-site occurrence count | included in 10 duplicate source-target pairs from current source-site report |
| `stop()` false-positive edge to `SSEListener.stop` | observed | absent; source site unresolved/ambiguous unless proven | absent in `TestResolveBareGoCallDoesNotFallbackToCrossLanguageMethod` |
| Selector/import function references emitted as ACCESSES | pending | 0 | current graph has 0 non-property ACCESSES targets |
| Coarse File-source CALLS | 16 observed by source-site accuracy command | 0 | 0 after File-source CALLS gate |
| Proof-based graphaccuracy/CLI report command | pending | emits all B5A metrics | `avmatrix source-site-accuracy` implemented for graph-inventory and `--golden` fixture modes |

Low-confidence global CALLS fallback slice notes:

- Same-file `CALLS` fallback remains resolved but now records confidence `0.950` in the TypeScript graph signature fixture instead of being promoted through global `resolveName` at confidence `1.000`.
- PHP `use function` import evidence now keeps imported PHP function calls resolved through import binding rather than global fallback; the wider backend suite passed after this change.
- Full source-site inventory now persists in graph JSON and LadybugDB relationship rows for resolved relationships, and in unresolved diagnostics for unresolved sites. The first P2A-I report command now exists as `avmatrix source-site-accuracy`; older values above the fresh command snapshot came from temporary PowerShell inventory before the command was available.

Source-site metadata persistence and accuracy command slice notes:

- Fresh analyze after the schema update reported files scanned `733`, parsed `544`, unsupported `189`, failed `0`, graph nodes `22404`, and graph relationships `51521`.
- Current source-site relationship counts by type: `CALLS=7611`, `ACCESSES=3136`, `USES=4512`, `INHERITS=1`.
- Current source-site relationship proof counts: `scope-binding=4889`, `same-file=3783`, `go-same-package=2192`, `receiver-member=3736`, `import-member=660`.
- Current unresolved diagnostic buckets by fact family: `call=30789`, `access=17846`, `type-reference=8471`, `heritage=7`.
- Current unresolved diagnostic occurrences by fact family: `call=30789`, `access=17846`, `type-reference=9217`, `heritage=7`.
- Fresh graph inventory after the source-site accuracy command slice used local build output: `.\avmatrix-launcher\server-bundle\avmatrix.exe analyze --force` scanned `736`, parsed `547`, unsupported `189`, failed `0`, graph nodes `22628`, and graph relationships `52161`.
- Command output artifact: `.tmp\2026-05-22-source-site-accuracy.json` from `.\avmatrix-launcher\server-bundle\avmatrix.exe source-site-accuracy --graph .avmatrix\graph.json --out .tmp\2026-05-22-source-site-accuracy.json --max-examples 20`.
- Current source-site inventory from the command: relationship buckets `15494`, relationship occurrences `26221`, diagnostic buckets `57388`, diagnostic occurrences `58134`, all source-site occurrences `84355`, stable source-site ID occurrences `84355`, missing source-site ID occurrences `0`.
- Current source-site relationship occurrences by type: `CALLS=12261`, `ACCESSES=6286`, `USES=7673`, `INHERITS=1`.
- Current unresolved diagnostic occurrences by fact family from the command: `call=30896`, `access=17924`, `type-reference=9307`, `heritage=7`.
- Current source-site proof counts from the command: `scope-binding=8413`, `same-file=6706`, `go-same-package=3035`, `receiver-member=7255`, `import-member=812`, `global-fallback-low-confidence=2159`, `none=55975`.
- Current status counts from the command: `resolved=26221`, `unresolved_local_binding=58134`. More precise external/ambiguous/dynamic/unsupported status splitting remains pending in later Phase 2A/Phase 3 work.
- Current graph-policy violation candidates from the command: false resolved edge candidates `16`, all from coarse `File -> Function` `CALLS` edges; resolved edges without proof `0`, resolved edges without source-site ID `0`, low-confidence fallback resolved edges `0`, non-property ACCESSES targets `0`.
- Fresh graph inventory after the File-source CALLS gate used local build output: `.\avmatrix-launcher\server-bundle\avmatrix.exe analyze --force` scanned `736`, parsed `547`, unsupported `189`, failed `0`, graph nodes `22635`, and graph relationships `52144`.
- Command output artifact after the gate: `.tmp\2026-05-22-source-site-accuracy-after-file-source-call-gate.json` from `.\avmatrix-launcher\server-bundle\avmatrix.exe source-site-accuracy --graph .avmatrix\graph.json --out .tmp\2026-05-22-source-site-accuracy-after-file-source-call-gate.json --max-examples 20`.
- Current source-site inventory after the gate: relationship buckets `15476`, relationship occurrences `26177`, diagnostic buckets `57449`, diagnostic occurrences `58195`, all source-site occurrences `84372`, stable source-site ID occurrences `84372`, missing source-site ID occurrences `0`.
- Current source-site accuracy policy after the gate: false resolved edge candidates `0`, resolved edges without proof `0`, resolved edges without source-site ID `0`, low-confidence fallback resolved edges `0`, coarse file call edges `0`, non-property ACCESSES targets `0`.

Golden corpus slice notes:

- `TestProofBasedCallAccessGoldenCorpus` fixture size: 10 call source sites and 2 access source sites.
- Golden resolved source-site occurrences: 5 `CALLS` occurrences and 1 `ACCESSES` occurrence.
- Golden unresolved source-site diagnostics: 5 call diagnostics and 1 access diagnostic.
- Golden duplicate/merge check: two `helper()` calls merge into one `CALLS` edge with both source-site IDs and `sourceSiteCount=2`.
- Golden false resolved edges: `0`.
- Golden silent missing source sites: `0`.
- Golden non-property resolved `ACCESSES` targets: `0`.

Source-site accuracy golden fixture command slice notes:

- Built command help exposes `--golden string` for source-site golden fixture JSON.
- Fixture-mode unit metrics from `TestRunSourceSiteAccuracyValidatesGoldenFixture`: expected source-site IDs `7`, matched source-site IDs `6`, silent missing source sites `1`, expected false resolved edges `1`, false resolved edges found `1`.
- CLI fixture-mode JSON visibility from `TestSourceSiteAccuracyCommandOutputsJSON`: golden validation enabled, silent missing source sites `1`, and false resolved edges `1`.
- Fresh current-graph command artifact without fixture: `.tmp\2026-05-22-p2a-source-site-accuracy-command.json`.
- Fresh current-graph command inventory: relationship buckets `15520`, relationship occurrences `26304`, diagnostic buckets `57709`, diagnostic occurrences `58455`, all source-site occurrences `84759`, stable source-site ID occurrences `84759`, missing source-site ID occurrences `0`.
- Fresh current-graph resolved edges: `CALLS=7645`, `ACCESSES=3314`; low-confidence fallback diagnostics `2217`; ACCESSES Property targets `3314/3314`; non-property ACCESSES targets `0`.
- Fresh current-graph policy: false resolved edge candidates `0`, resolved edges without proof `0`, resolved edges without source-site ID `0`, low-confidence fallback resolved edges `0`, coarse file call edges `0`.
- Fresh current-graph golden validation: disabled when no fixture is supplied, with expected source sites `0`, matched source sites `0`, silent missing source sites `0`, expected false resolved edges `0`, and false resolved edges `0`.
- Fresh post-edit analyze before detect-changes: scanned `737`, parsed `548`, unsupported `189`, failed `0`, graph nodes `22751`, graph relationships `52425`.
- Pre-commit detect-changes summary: changed_count `74`, changed_files `7`, affected_count `12`, risk_level `high`.

## B6 - ResolutionGap Persistence Metrics

Status: in progress; validation recorded through Phase 2A golden corpus slice.

Record after P3.

| Metric | Before | Target / expected direction | After |
| --- | ---: | --- | ---: |
| Persisted ResolutionGap/UnresolvedSymbol entities | pending | source-backed unresolved buckets or accurate aggregate entities | pending |
| Persisted gap relationships/typed relations | pending | nonzero where source evidence supports them | pending |
| Gaps preserving source node ID | pending | all source-backed gaps | pending |
| Gaps preserving sourceSiteID | pending | all gaps originating from call/access source-site inventory | pending |
| Gaps preserving source-site status | pending | all gaps originating from call/access source-site inventory | pending |
| Gaps preserving proof kind | pending | all gaps originating from call/access source-site inventory | pending |
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
| Full build | `powershell -ExecutionPolicy Bypass -File .\avmatrix-launcher\build.ps1` | passed | run before tests after Functional Area code changes |
| Backend tests | `go test .\internal\semantic .\internal\lbugschema .\internal\lbugload .\internal\contracts .\internal\httpapi .\internal\analyze` | passed | focused Functional Area, schema/export, contract, API status coverage |
| Backend tests | `go test .\internal\... .\cmd\...` | passed | wider Go validation excluding intentionally non-buildable fixture packages |
| Contract generation/check | `go run .\cmd\generate-web-contracts` | passed | regenerated schema and TypeScript contract from Go source |
| Web unit tests | `npm test -- --run` in `avmatrix-web` | passed | 44 test files, 357 tests |
| Web e2e tests | not run for Phase 2 | not applicable to this slice | Functional Area is persisted graph/contract metadata; visible Web UI behavior remains Phase 7 |
| Source-site accuracy command | `.\avmatrix-launcher\server-bundle\avmatrix.exe source-site-accuracy --graph .avmatrix\graph.json --out .tmp\2026-05-22-source-site-accuracy.json --max-examples 20` | passed | graph-inventory mode reported 84355 source-site occurrences, 0 missing source-site IDs, 0 non-property ACCESSES targets, and 16 coarse File-source CALLS candidates |
| Source-site accuracy command after File-source CALLS gate | `.\avmatrix-launcher\server-bundle\avmatrix.exe source-site-accuracy --graph .avmatrix\graph.json --out .tmp\2026-05-22-source-site-accuracy-after-file-source-call-gate.json --max-examples 20` | passed | graph-inventory mode reported 84372 source-site occurrences, 0 missing source-site IDs, 0 non-property ACCESSES targets, and 0 false resolved edge candidates |
| Backend tests | `go test .\internal\graphaccuracy` and `go test .\internal\cli` | passed | focused source-site accuracy report and packaged CLI command coverage |
| Backend tests | `go test .\internal\... .\cmd\...` | passed | wider Go validation after source-site accuracy command implementation |
| Backend tests | `go test .\internal\resolution` and `go test .\internal\graphaccuracy` | passed | focused File-source CALLS gate and source-site accuracy coverage |
| Backend tests | `go test .\internal\... .\cmd\...` | passed | wider Go validation after File-source CALLS gate |
| AVmatrix detect-changes for File-source CALLS gate | `.\avmatrix-launcher\server-bundle\avmatrix.exe detect-changes --repo AVmatrix --scope all` | passed | affected_count 0, changed_count 15, changed_files 7, risk_level low |
| Full build before golden corpus tests | `powershell -ExecutionPolicy Bypass -File .\avmatrix-launcher\build.ps1` | passed | run before tests for `TestProofBasedCallAccessGoldenCorpus` slice |
| Backend tests | `go test .\internal\resolution` | passed | focused golden corpus coverage for proof-based CALLS/ACCESSES |
| Backend tests | `go test .\internal\graphaccuracy .\internal\cli .\internal\contracts .\internal\lbugload .\internal\graphhealth` | passed | focused graphaccuracy/CLI/contract/export/diagnostic surfaces after golden corpus slice |
| Backend tests | `go test .\internal\... .\cmd\...` | passed | wider Go validation after golden corpus slice |
| AVmatrix detect-changes for golden corpus | `.\avmatrix-launcher\server-bundle\avmatrix.exe detect-changes --repo AVmatrix --scope all` | passed | after staging new test file and final ledger updates: affected_count 0, changed_count 81, changed_files 4, risk_level low |
| AVmatrix detect-changes for source-site accuracy command | `.\avmatrix-launcher\server-bundle\avmatrix.exe detect-changes --repo AVmatrix --scope all` | passed with expected high scope | affected_count 11, changed_count 35, changed_files 6; affected surfaces match root CLI command and graphaccuracy report/schema decoding changes |
| Query-health command | pending | pending | pending |
| Resolution inventory command | pending | pending | pending |
| `query` semantic output | pending | pending | pending |
| `context` semantic output | pending | pending | pending |
| `impact` semantic output | pending | pending | pending |
| `detect-changes` semantic output | pending | pending | pending |
| API-specific MCP semantic output | pending | pending | pending |
| AVmatrix detect-changes for implementation commits | `.\avmatrix-launcher\server-bundle\avmatrix.exe detect-changes --repo AVmatrix --scope all` | passed with expected critical scope | affected_count 24, changed_count 144, changed_files 18; affected surfaces match Functional Area semantic/schema/API/contract/export slice |

## B12 - Semantic Enrichment Flow Metrics

Status: in progress; App Layer and Functional Area semantic enrichment measured.

Record after the analyze semantic enrichment phase is introduced and after each phase extends it.

| Metric | Baseline | After App Layer | After Functional Area | After Source-Site Inventory | After ResolutionGap | Final |
| --- | ---: | ---: | ---: | ---: | ---: | ---: |
| Semantic enrichment phase latency | pending | 33.9997 ms | 68.8814 ms | pending | pending | pending |
| Semantic enrichment memory delta | pending | pending | overall analyze heap +89525848 bytes; not phase-isolated | pending | pending | pending |
| Graph nodes before enrichment | pending | 22239 | 22358 | pending | pending | pending |
| Graph nodes after enrichment | pending | 22239 | 22358 | pending | pending | pending |
| Graph relationships before enrichment | pending | 55006 | 55349 | pending | pending | pending |
| Graph relationships after enrichment | pending | 55006 | 55349 | pending | pending | pending |
| Graph JSON size | pending | 45739916 bytes | 47953614 bytes | pending | pending | pending |
| LadybugDB node rows | pending | 22239 | 22358 | pending | pending | pending |
| LadybugDB relationship rows | pending | 55006 | 55349 | pending | pending | pending |
| Duplicate graph traversals introduced | pending | 0 nested whole-graph loops; one node pass and one relationship pass | 0 nested whole-graph loops; one node pass and one relationship pass shared by App Layer and Functional Area | pending | pending | pending |
| File rescans introduced | pending | 0 | 0 | pending | pending | pending |
| AST reparses introduced | pending | 0 | 0 | pending | pending | pending |
| Raw call/access source-site count | pending | pending | pending | pending | pending | pending |
| Raw unresolved fact count | pending | pending | pending | pending | pending | pending |
