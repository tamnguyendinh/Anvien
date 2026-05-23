# AVmatrix App Layer Resolution Gap Lens Benchmark Ledger

Date: 2026-05-22

Status: Phase 9 and Phase 10 closure validation complete

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

Status: complete for App Layer persistence metrics.

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

Status: complete for Phase 2A; graph-inventory, golden corpus, source-site accuracy golden fixture command, and source-site bridge metrics recorded.

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

Source-site to ResolutionGap input bridge slice notes:

- Source-backed bridge model: `internal/graphhealth.ResolutionGapInput`.
- Full source-backed extractor: `SourceBackedResolutionGapInputs`.
- Call/access-specific Phase 3 input extractor: `SourceBackedCallAccessResolutionGapInputs`.
- Fixture bridge test counts: `3` source-backed unresolved inputs from call/access/type-reference diagnostics; `2` call/access inputs; `1` diagnostic without `sourceSiteId` excluded from precise source-backed input.
- Fresh current-graph command artifact for this slice: `.tmp\2026-05-22-p2a-gap-input-source-site-accuracy.json`.
- Fresh current-graph source-site inventory for the bridge: relationship buckets `15578`, relationship occurrences `26395`, diagnostic buckets `57814`, diagnostic occurrences `58560`, all source-site occurrences `84955`, stable source-site ID occurrences `84955`, missing source-site ID occurrences `0`.
- Fresh unresolved diagnostic fact-family counts usable by Phase 3 input records: call `31091`, access `18129`, type-reference `9333`, heritage `7`.
- Fresh current-graph policy after the bridge: false resolved edge candidates `0`, resolved edges without proof `0`, resolved edges without source-site ID `0`, low-confidence fallback resolved edges `0`, coarse file call edges `0`, non-property ACCESSES targets `0`.
- Fresh post-edit analyze before detect-changes: scanned `739`, parsed `550`, unsupported `189`, failed `0`, graph nodes `22806`, graph relationships `52584`.
- Pre-commit detect-changes summary after staging new files: changed_count `56`, changed_files `5`, affected_count `0`, risk_level `low`.

## B6 - ResolutionGap Persistence Metrics

Status: complete for Phase 3.

Record after P3.

| Metric | Before | Target / expected direction | After |
| --- | ---: | --- | ---: |
| Persisted ResolutionGap/UnresolvedSymbol entities | 0 | source-backed unresolved buckets or accurate aggregate entities | 58495 |
| Persisted gap relationships/typed relations | 0 | nonzero where source evidence supports them | 58495 `HAS_RESOLUTION_GAP` |
| Gaps preserving source node ID | pending | all source-backed gaps | 58495/58495; missing 0 |
| Gaps preserving sourceSiteID | pending | all gaps originating from call/access source-site inventory | 58495/58495; missing 0 |
| Gaps preserving source-site status | pending | all gaps originating from call/access source-site inventory | 58495/58495; missing 0 |
| Gaps preserving proof kind | pending | all gaps originating from call/access source-site inventory | 58495/58495; missing 0 |
| Gaps preserving target text | pending | all source-backed gaps | 58495/58495; missing 0 |
| Gaps preserving source App Layer | pending | all source-backed gaps with classified source | 58495/58495; unknown/missing 0 |
| Gaps preserving source Functional Area | pending | all source-backed gaps with classified source | latest aggregate-policy graph retains source Functional Area distribution; unknown count pending next full field inventory |
| Gaps preserving actionability | pending | all modeled gaps | 58495/58495; missing 0 |
| Fake resolved target nodes created | pending | 0 | 0 in semantic fixtures; `ValidateResolutionGapPersistence` covers fake resolved target claims, fake topology edges, and resolved-edge source-site overlap |
| Aggregates preserving exact occurrence count | pending | all aggregates | 35411 aggregate buckets over 58495 gap nodes preserve 59242 exact occurrences |
| Aggregates preserving source samples | pending | all aggregates | representative samples are capped only by option; full sourceSiteID traceability remains uncapped |
| Evidence capped away for graph-size reasons | pending | 0 | 0 for persisted source-backed gaps; aggregate layer does not reduce persisted graph entities |

By fact family:

| Fact family | Count after |
| --- | ---: |
| unresolved call | 31393 |
| unresolved access | 18418 |
| unresolved type-reference | 8677 |
| unresolved heritage | 7 |
| external symbol | 158 external_library |
| builtin/stdlib/test reference | 24285 total; builtin 9927, standard_library 7134, test_framework 7224 |
| in-repo analyzer gap | 34052 |
| unknown/unclassified | 0 in current analyzed graph |

By target role:

| Target role | Count after |
| --- | ---: |
| callable | 31393 |
| member | 18418 |
| type | 8684 |
| external | 0 as target role; represented by classification `external_library=158` |
| builtin | 0 as target role in current fact-family-backed graph; represented by classification `builtin=9927` |
| test | 0 as target role in current fact-family-backed graph; represented by classification `test_framework=7224` |
| unknown | 0 |

By actionability:

| Actionability | Count after |
| --- | ---: |
| non_actionable | 24285 |
| review | 158 |
| analyzer_gap | 34052 |
| unknown | 0 |

ResolutionGap persistence slice notes:

- Fresh analyze artifact: `.tmp\2026-05-22-p3-resolution-gap-persist-analyze.json`.
- Fresh analyze inventory: scanned `739`, parsed `550`, unsupported `189`, failed `0`, graph nodes `80957`, graph relationships `110949`.
- Semantic persistence metrics from analyze: `resolutionGapInputs=58083`, `resolutionGapNodes=58083`, `resolutionGapRelationships=58083`.
- LadybugDB load metrics from analyze: node rows `80957`, relationship rows `110949`, skipped relationships `0`, fallback inserts `0`.
- Graph JSON inventory check: `ResolutionGap` nodes `58083`; `HAS_RESOLUTION_GAP` relationships `58083`; missing source node ID `0`, sourceSiteID `0`, source-site status `0`, proof kind `0`, target text `0`, source App Layer `0`, actionability `0`; unknown source Functional Area `7716`.
- Role-validation analyze artifact: `.tmp\2026-05-22-p3-role-validation-postedit-analyze.json`.
- Role-validation analyze inventory: scanned `741`, parsed `552`, unsupported `189`, failed `0`, graph nodes `81298`, graph relationships `111411`.
- Role-validation semantic metrics: `resolutionGapInputs=58350`, `resolutionGapNodes=58350`, `resolutionGapRelationships=58350`, semantic phase latency `609.862 ms`.
- Role-validation graph JSON inventory: `ResolutionGap` nodes `58350`; `HAS_RESOLUTION_GAP` relationships `58350`; missing source node ID `0`, sourceSiteID `0`, source-site status `0`, proof kind `0`, target text `0`, source App Layer `0`; unknown source Functional Area `7821`.
- Role-validation source-site accuracy artifact: `.tmp\2026-05-22-p3-role-validation-source-site-accuracy.json`.
- Role-validation source-site accuracy inventory: relationship source-site occurrences `85879`, diagnostic source-site occurrences `59097`, all source-site occurrences `144976`, stable source-site IDs `144976`, missing IDs `0`, resolved `CALLS=7733`, resolved `ACCESSES=3490`, unresolved diagnostics `59097`, low-confidence fallback diagnostics `2260`, non-property ACCESSES targets `0`, false resolved edge candidates `0`, resolved edges without proof `0`, resolved edges without sourceSiteID `0`, low-confidence fallback resolved edges `0`, coarse file call edges `0`.
- Aggregate-policy analyze artifact: `.tmp\2026-05-22-p3-aggregate-policy-postedit-analyze.json`.
- Aggregate-policy analyze inventory: scanned `743`, parsed `554`, unsupported `189`, failed `0`, graph nodes `81522`, graph relationships `111801`.
- Aggregate-policy semantic metrics: `resolutionGapInputs=58495`, `resolutionGapNodes=58495`, `resolutionGapRelationships=58495`, semantic phase latency `595.5799 ms`.
- Aggregate-policy source-site accuracy artifact: `.tmp\2026-05-22-p3-aggregate-policy-source-site-accuracy.json`.
- Aggregate-policy source-site accuracy inventory: relationship source-site occurrences `86148`, diagnostic source-site occurrences `59242`, all source-site occurrences `145390`, stable source-site IDs `145390`, missing IDs `0`, resolved `CALLS=7751`, resolved `ACCESSES=3548`, unresolved diagnostics `59242`, low-confidence fallback diagnostics `2266`, non-property ACCESSES targets `0`, false resolved edge candidates `0`, resolved edges without proof `0`, resolved edges without sourceSiteID `0`, low-confidence fallback resolved edges `0`, coarse file call edges `0`.
- Aggregate inventory over latest graph: source-backed gap nodes `58495`, aggregate buckets `35411`, exact occurrences `59242`, buckets with multiple source sites `10854`, max bucket source sites `99`, max bucket occurrences `99`.
- The current graph still persists each sourceSiteID-backed unresolved diagnostic as its own graph entity. The aggregate/dedupe policy is an inventory layer that keeps full sourceSiteID traceability and exact counts.

## B7 - Resolution Health Inventory Metrics

Status: complete for Phase 4.

Measured after `.\avmatrix-launcher\server-bundle\avmatrix.exe analyze --force --benchmark-json .\.tmp\2026-05-22-p4-resolution-health-inventory-analyze.json --benchmark-label p4-resolution-health-inventory` and `.\avmatrix-launcher\server-bundle\avmatrix.exe resolution-inventory --graph .\.avmatrix\graph.json --out .\.tmp\2026-05-22-p4-resolution-inventory.json`.

Record after P4.

| Resolution Health bucket | Count after |
| --- | ---: |
| resolved references | 27195 |
| unresolved non-actionable | 24926 |
| external unresolved | 163 |
| in-repo analyzer gap | 34537 |
| unresolved call target | 31590 |
| unresolved access target | 18541 |
| unresolved type target | 9488 |
| unresolved heritage target | 7 |
| unclassified/unknown | 0 |

Topology plus resolution overlay:

| Topology status | Nodes with no gaps | Nodes with gaps | Nodes with degraded confidence |
| --- | ---: | ---: | ---: |
| connected | 595 | 2146 | 2146 |
| no_incoming | 175 | 1525 | 1525 |
| no_outgoing | 2649 | 567 | 567 |
| detached_component | 88 | 143 | 143 |
| true_isolated | 73982 | 192 | 192 |
| unknown_connectivity | 0 | 0 | 0 |

Resolution inventory totals:

| Metric | Count |
| --- | ---: |
| Graph nodes | 82062 |
| Graph relationships | 112614 |
| ResolutionGap nodes | 58879 |
| `HAS_RESOLUTION_GAP` relationships | 58879 |
| Gap occurrences | 59626 |
| Resolved references | 27195 |
| Resolution Confidence clear nodes | 370 |
| Resolution Confidence degraded nodes | 4573 |
| Resolution Confidence unknown nodes | 77119 |

ResolutionGap occurrence counts by App Layer:

| App Layer | Count |
| --- | ---: |
| backend_test | 20057 |
| backend | 18126 |
| frontend_test | 5510 |
| api_test | 5114 |
| frontend | 4846 |
| api | 4296 |
| api_contract | 627 |
| cli_launcher | 553 |
| frontend_api_client | 387 |
| config | 88 |
| generated_contract | 22 |

ResolutionGap occurrence counts by Functional Area:

| Functional Area | Count |
| --- | ---: |
| providers | 12674 |
| unknown | 7881 |
| analyzer | 7084 |
| mcp | 4716 |
| api | 4660 |
| resolution | 4060 |
| cli | 3116 |
| graph_health | 2590 |
| storage | 2639 |
| web_graph_ui | 2616 |
| query | 2493 |
| embeddings | 1701 |
| session | 1062 |
| contracts | 880 |
| launcher | 671 |
| layout | 665 |
| configuration | 88 |
| runtime | 30 |

Source-site accuracy after Phase 4:

| Metric | Count |
| --- | ---: |
| Relationship source-site occurrences | 86821 |
| Diagnostic source-site occurrences | 59626 |
| All source-site occurrences | 146447 |
| Missing source-site IDs | 0 |
| Resolved CALLS | 7802 |
| Resolved ACCESSES | 3639 |
| False resolved edge candidates | 0 |
| Non-property ACCESSES targets | 0 |
| Resolved edges without proof | 0 |
| Low-confidence fallback resolved edges | 0 |
| Coarse File CALLS | 0 |

## B8 - Query Health Metrics

Status: complete for P6-A query retrieval; rerun in P8 for final full-plan validation.

Record before P5 and after P5.

| Intent | Expected core files/symbols | Baseline hit@5 | Baseline hit@10 | Final hit@5 | Final hit@10 | Result |
| --- | --- | ---: | ---: | ---: | ---: | --- |
| unresolved reference diagnostic generation | `internal/resolution/resolve.go`, `internal/resolution/emit.go`, `internal/graphhealth/diagnostics.go`, `resolveCall`, `emitUnresolvedReference`, `AppendDiagnosticToNode` | 1/6 | 1/6 | 4/6 | 4/6 | pass threshold after P6-A ranking/output changes |
| graph health unknown-connectivity separation | `internal/graphhealth/compute.go`, `internal/graphhealth/policy.go`, `avmatrix-web/src/lib/graph-health-filters.ts`, `ComputeSummary`, `getNodeGraphHealth` | 0/5 | 0/5 | 4/5 | 5/5 | pass threshold after graph-health surface boost |
| App Layer and ResolutionGap layout | `avmatrix-web/src/lib/graph-adapter.ts`, `avmatrix-web/src/hooks/useSigma.ts`, `avmatrix-web/src/lib/graph-health-filters.ts`, `knowledgeGraphToGraphology`, `useSigma` | 0/5 | 0/5 | 5/5 | 5/5 | pass threshold after layout/frontend surface boost and per-file result diversity |
| runtime reset hidden-terminal behavior | `avmatrix-launcher/src/main.go`, `startRuntime`, `resetRuntime`, `stopRuntime`, `hiddenCommand` | 4/5 | 4/5 | 4/5 | 4/5 | pass threshold |
| API contract surfaces | `internal/httpapi/graph.go`, `internal/contracts/web_ui.go`, `avmatrix-web/src/generated/avmatrix-contracts.ts`, `WebUIContract`, `WebUIContractTypeScript` | 0/5 | 2/5 | 3/5 | 4/5 | pass threshold after API/contract semantic surface boost |
| query implementation surfaces | `internal/mcp/tools.go`, `internal/cli/tool_command.go`, `cmd/avmatrix/main.go`, `queryTool`, `rankedProcessMatches`, `matchingDefinitionRows`, `newQueryCommand` | 0/7 | 0/7 | 6/7 | 6/7 | pass threshold after MCP/query/CLI surface boost |
| frontend graph filter surfaces | `avmatrix-web/src/lib/graph-health-filters.ts`, `avmatrix-web/src/hooks/app-state/graph.tsx`, `avmatrix-web/src/components/FileTreePanel.tsx`, `avmatrix-web/src/components/GraphCanvas.tsx`, `getNodeGraphHealth`, `GraphCanvas` | 0/6 | 0/6 | 4/6 | 4/6 | pass threshold after frontend filter/detail surface boost |

## B9 - Semantic Command Surface Metrics

Status: complete for Phase 6.

Record after P6.

| Command | App Layer shown | Functional Area shown | ResolutionGap shown | Resolution Health shown | Limitation noted |
| --- | --- | --- | --- | --- | --- |
| `query` | yes | yes | yes | yes | complete for P6-A; `semanticStatus` and stale semantic warning are included, and per-row semantic fields are emitted only from persisted graph data |
| `context` | yes | yes | yes | yes | complete for P6-B; includes `semanticStatus`, stale semantic warning, semantic fields on symbol/ref/candidate/process rows, source-site proof metadata on relationship rows, `sourceResolutionGaps` for source-node gaps, and `resolutionGapSources` for selected gap entities |
| `impact` | yes | yes | yes | yes | complete for P6-C; includes semanticStatus, affected App Layer/Functional Area counts, resolutionHealthRisks, semantic fields on target/impacted/process/module rows, source-site proof metadata, and non-blocking workflow warning for HIGH/CRITICAL risk |
| `detect-changes` | yes | yes | yes | yes | complete for P6-D; includes semanticStatus, changed/affected App Layer and Functional Area counts, semantic fields on changed symbols/processes/steps, ResolutionGap change summaries, and Resolution Health impact summaries when changed rows carry gap/degraded evidence |
| API MCP tools | yes | yes | yes | yes | complete for P6-E; route_map, shape_check, and api_impact include semanticStatus plus route, consumer, flow-detail, summary, and Resolution Health impact fields when persisted route graph data exists; current AVmatrix graph has no route nodes, so live artifacts show semanticStatus on empty/error API-tool payloads while fixture tests prove populated rows |
| stale/missing semantic metadata tests | yes | yes | yes | yes | complete for P6-F; focused MCP tests prove stale/incomplete semantic metadata emits warnings and command/API output does not invent App Layer fields for `impact`, `detect-changes`, `route_map`, `shape_check`, or `api_impact` |
| resolution inventory command | yes | yes | yes | yes | implemented in Phase 4; command exposes full Resolution Health inventory from persisted graph data |
| query-health command | yes | yes | yes | yes | P6-A final run passed 7/7 suite cases and captures semantic fields returned by `query` |

## B10 - Web UI Ring And Filter Metrics

Status: complete for Phase 7 Web filter, detail-panel, ring layout, island, and manual optimizer metrics.

Record after P7.

| Metric | Value |
| --- | ---: |
| Visible App Layer ring count | 14 |
| Backend ring node count | 28466 |
| API ring node count | 6907 |
| Frontend ring node count | 7585 |
| API ring placed between Backend and Frontend | true |
| Contract rings placed near API when present | api_contract 790 nodes; generated_contract 38 nodes; frontend_api_client 555 nodes near API/Frontend side |
| Shared/API contract ring count | api_contract 1; generated_contract 1; standalone shared/api_shared rings absent in current graph |
| Frontend API client ring count | 1 |
| Test ring/group count | 3 rings: backend_test 24683, frontend_test 6678, api_test 6948 |
| Docs/Config/Generated ring count | docs 1614, config 125, generated_contract 38 |
| Unknown/Mixed ring count | unknown 26, mixed 365 |
| Ring size/spacing policy recorded | deterministic App Layer angles, adaptive island radius, island gap, ring gap, docs center clearance, no fixed ring limit |
| Default visible ring/lens count | all App Layer rings remain data-visible by default; node visibility still follows existing node-type/App Layer/Resolution Health filters |
| Default hidden/collapsed ring/lens count | no ring-level collapse added; uncommon node types may remain hidden by existing default node-type policy |
| Node type islands visible | backend 16, api 16, frontend 16, docs 3; all visible rings report at least 2 islands |
| Same-color island violations | 0 |
| ResolutionGap visible count | covered by persisted graph/API filter counts from P7-B; browser ring metric groups ResolutionGap by gap kind when visible |
| App Layer filters available | 17 default contract values plus missing App Layer stale/incomplete graph toggle |
| Resolution Health filters available | 3 confidence values, 9 health buckets, 8 gap fact families, 7 target roles, diagnostic classifications/actionability values, source App Layer filters, top target text filters, and 11 required lens rows |
| Detail panel App Layer fields | implemented from persisted node `appLayer` |
| Detail panel Functional Area fields | implemented from persisted node `functionalArea` |
| Detail panel Resolution Health fields | implemented: confidence, gap count, bucket counts |
| Detail panel related ResolutionGap rows | implemented via persisted `HAS_RESOLUTION_GAP` relationships |
| Degraded resolution confidence dead-code separation | implemented and covered by focused test |
| Optimizer auto-run events after render/load/filter | 0 in browser diagnostics before manual click; manual optimizer invocation count increments only after `Optimize Layout` click |

P7-A/P7-A2/P7-B filter/lens counts:

| Metric | Value |
| --- | ---: |
| Default App Layer values | 17 |
| Missing App Layer toggle | 1 |
| Resolution Confidence values | 3 |
| Resolution Health bucket values | 9 |
| ResolutionGap fact-family values | 8 |
| ResolutionGap target-role values | 7 |
| Required Resolution Health lens rows implemented | 11 |
| Focused semantic/Web filter unit tests | 29 passed |
| Full Web unit tests after filter slice | 363 passed |
| Playwright e2e after hidden Vite server start | 14 passed / 30 skipped |

P7-C detail-panel counts:

| Metric | Value |
| --- | ---: |
| Detail panel semantic field groups implemented | 4: App Layer, Functional Area, Topology Health, Resolution Health |
| Related ResolutionGap lookup path | 1 persisted relationship type: `HAS_RESOLUTION_GAP` |
| Related ResolutionGap fields rendered | 8: fact family, target text, target role, classification, actionability, source-site status, proof kind, count |
| Focused detail-panel unit tests | 1 passed |
| Full Web unit tests after detail-panel slice | 363 passed |
| Playwright e2e after detail-panel slice | 14 passed / 30 skipped |

P7-D through P7-J ring/layout counts:

| Metric | Value |
| --- | ---: |
| Browser graph nodes | 85561 |
| Browser App Layer rings | 14 |
| Ring screenshot artifact size | 231660 bytes |
| Ring screenshot artifact | `avmatrix-web\test-results\server-connect-Graph-Dashb-1289f-ode-type-islands-in-browser-chromium\app-layer-rings-visible.png` |
| Focused layout/unit tests | 29 passed |
| Full Web unit tests after final Phase 7 edits | 368 passed |
| Full e2e per-spec total | 44 passed / 1 expected packaged-launcher skip |
| Focused screenshot rerun | 1 passed |
| Browser `sameColorIslandViolations` | 0 |
| Browser `apiBetweenBackendAndFrontend` | true |
| Browser `docsCentered` | true |

## B11 - Validation Outputs

Status: complete through P8-J.

Build/test/e2e timings are validation evidence, not product performance benchmarks unless this plan changes those systems.

| Validation | Command | Result | Notes |
| --- | --- | --- | --- |
| Full build before P7-A/P7-A2/P7-B Web filter tests | `powershell -ExecutionPolicy Bypass -File .\avmatrix-launcher\build.ps1` | passed | run after App Layer/Resolution Health Web filter implementation |
| Focused Web filter tests | `npm --prefix .\avmatrix-web run test -- --run test/unit/semantic-filters.test.ts test/unit/graph-adapter.edge-geometry.test.ts test/unit/FileTreePanel.dashboard-completeness.test.tsx` | passed | 3 files and 29 tests passed; covers semantic filter counts/predicates, graph adapter filter composition, dashboard App Layer and Resolution Health rendering |
| GraphCanvas semantic depth-filter guard | `npm --prefix .\avmatrix-web run test -- --run test/unit/GraphCanvas.selection-performance.test.tsx` | passed | 1 file and 3 tests passed; covers selection/depth filtering and no manual optimizer invocation during graph load |
| Web unit tests after P7-A/P7-A2/P7-B | `npm --prefix .\avmatrix-web run test -- --run` | passed | 45 files and 363 tests passed |
| Web e2e after P7-A/P7-A2/P7-B | `npm --prefix .\avmatrix-web run test:e2e` | passed after starting hidden Vite dev server | direct run without frontend server failed with `ERR_CONNECTION_REFUSED`; rerun with hidden Vite dev server on 127.0.0.1:5228 passed 14 and skipped 30 |
| Fresh analyze benchmark for P7-A/P7-A2/P7-B Web filters | `.\avmatrix-launcher\server-bundle\avmatrix.exe analyze --force --benchmark-json .\.tmp\2026-05-22-p7a-p7b-web-semantic-filters-analyze.json --benchmark-label p7a-p7b-web-semantic-filters` | passed | files scanned 755, parsed 565, unsupported 190, failed 0; graph nodes 84880 and relationships 116668 |
| AVmatrix detect-changes for P7-A/P7-A2/P7-B Web filters | `.\avmatrix-launcher\server-bundle\avmatrix.exe detect-changes --repo AVmatrix --scope all` | passed with expected critical scope | after staging all slice files: affected_count 24, changed_count 859, changed_files 17, risk_level critical; affected surfaces match shared Web graph state, graph adapter filtering, GraphCanvas, FileTreePanel, backend graph response handling, tests, and plan ledgers |
| Full build before P7-C detail-panel tests | `powershell -ExecutionPolicy Bypass -File .\avmatrix-launcher\build.ps1` | passed | run before tests after semantic detail-panel implementation |
| Focused Web detail-panel test | `npm --prefix .\avmatrix-web run test -- --run test/unit/CodeReferencesPanel.graph-health.test.tsx` | passed | 1 file and 1 test passed; covers App Layer, Functional Area, Topology Health, Resolution Health, degraded confidence copy, and related persisted ResolutionGap rendering |
| Web unit tests after P7-C | `npm --prefix .\avmatrix-web run test -- --run` | passed | 45 files and 363 tests passed |
| Web e2e after P7-C | `npm --prefix .\avmatrix-web run test:e2e` | passed | already-running hidden Vite dev server on 127.0.0.1:5228; 14 passed and 30 skipped |
| Fresh analyze benchmark for P7-C detail panel | `.\avmatrix-launcher\server-bundle\avmatrix.exe analyze --force --benchmark-json .\.tmp\2026-05-22-p7c-detail-lens-start-analyze.json --benchmark-label p7c-detail-lens-start` | passed | files scanned 755, parsed 565, unsupported 190, failed 0; graph nodes 84880 and relationships 116668 |
| AVmatrix detect-changes for P7-C detail panel | `.\avmatrix-launcher\server-bundle\avmatrix.exe detect-changes --repo AVmatrix --scope all` | passed with expected low scope | after staging all slice files: affected_count 0, changed_count 102, changed_files 5, risk_level low; changed app layers docs 8, frontend 68, frontend_test 26 |
| Full build before P7-D through P7-J layout tests | `powershell -ExecutionPolicy Bypass -File .\avmatrix-launcher\build.ps1` | passed | rerun after ring layout, diagnostics, and e2e test updates; temporary `avmatrix serve` process had to be stopped before build because it locked the generated exe |
| Focused P7 layout/unit tests | `npm --prefix .\avmatrix-web run test -- --run test/unit/graph-adapter.edge-geometry.test.ts test/unit/GraphCanvas.selection-performance.test.tsx test/unit/runtime-diagnostics.test.ts` | passed | 3 files and 29 tests passed; covers ring grouping, color/type islands, ResolutionGap islands, docs ring anchor fields, runtime diagnostics, and no auto optimizer invocation |
| Web unit tests after final Phase 7 edits | `npm --prefix .\avmatrix-web run test -- --run` | passed | 45 files and 368 tests passed |
| Server-connect e2e with backend for P7 layout | `npm --prefix .\avmatrix-web run test:e2e -- e2e/server-connect.spec.ts` | passed | 10 tests passed; includes large-graph stability without automatic optimizer, manual optimizer invocation only after user action, visual scale cap, and Backend/API/Frontend ring diagnostic |
| Graph Health e2e | `npm --prefix .\avmatrix-web run test:e2e -- e2e/graph-health-ui.spec.ts` | passed | 4 tests passed |
| Heartbeat reconnect e2e | `npm --prefix .\avmatrix-web run test:e2e -- e2e/heartbeat-reconnect.spec.ts` | passed | 2 tests passed after updating the e2e wait budget to match observed large-graph load; no product/runtime timeout behavior changed |
| Multi-repo scoping e2e | `npm --prefix .\avmatrix-web run test:e2e -- e2e/multi-repo-scoping.spec.ts` | passed | 3 tests passed |
| Repo switching e2e | `npm --prefix .\avmatrix-web run test:e2e -- e2e/repo-switching.spec.ts` | passed | 6 tests passed |
| Onboarding e2e | `npm --prefix .\avmatrix-web run test:e2e -- e2e/onboarding.spec.ts` | passed with expected skip | 12 tests passed; packaged launcher case skipped unless `PACKAGED_LAUNCHER_E2E=1` |
| Shell interactions e2e | `npm --prefix .\avmatrix-web run test:e2e -- e2e/shell-interactions.spec.ts` | passed | 7 tests passed after aligning select-all expectation with default hidden node-type policy for large graphs |
| Focused ring screenshot e2e rerun | `npm --prefix .\avmatrix-web run test:e2e -- e2e/server-connect.spec.ts -g "reports Backend API Frontend rings"` | passed | 1 test passed and wrote `app-layer-rings-visible.png` |
| Fresh analyze benchmark for P7-D through P7-J layout slice | `.\avmatrix-launcher\server-bundle\avmatrix.exe analyze --force --benchmark-json .\.tmp\2026-05-22-p7-layout-final2-precommit-analyze.json --benchmark-label p7-layout-final2-precommit` | passed | files scanned 755, parsed 565, unsupported 190, failed 0; graph nodes 85594 and relationships 117530 |
| AVmatrix detect-changes for P7-D through P7-J layout slice | `.\avmatrix-launcher\server-bundle\avmatrix.exe detect-changes --repo AVmatrix --scope all` | passed with expected high scope | after staging all slice files: affected_count 12, changed_count 761, changed_files 12, risk_level high; changed app layers docs 8, frontend 417, frontend_test 336; affected app layers frontend 9 and mixed 3 |
| P8-A full build gate | `powershell -ExecutionPolicy Bypass -File .\avmatrix-launcher\build.ps1` | passed | Go `go1.26.3`; rebuilt `avmatrix-web/dist`; produced launcher/server bundle artifacts before P8 tests |
| P8-B backend/cmd tests | `go test .\internal\... .\cmd\...` | passed | wide backend/cmd validation after the P8-A build gate |
| P8-C contract generation/check | `go run .\cmd\generate-web-contracts`; `go run .\cmd\generate-web-contracts --check` | passed | generated Web contract is in sync with Go contract source |
| P8-D Web unit tests | `npm --prefix .\avmatrix-web run test -- --run` | passed | 45 files and 368 tests passed; duration 35.84s |
| P8-E server-connect e2e | `npm --prefix .\avmatrix-web run test:e2e -- e2e/server-connect.spec.ts` | passed | 10 tests passed; includes multi-ring diagnostic, node size cap, graph controls, and manual optimizer guard |
| P8-E graph-health e2e | `npm --prefix .\avmatrix-web run test:e2e -- e2e/graph-health-ui.spec.ts` | passed | 4 tests passed |
| P8-E heartbeat reconnect e2e | `npm --prefix .\avmatrix-web run test:e2e -- e2e/heartbeat-reconnect.spec.ts` | passed | 2 tests passed |
| P8-E multi-repo scoping e2e | `npm --prefix .\avmatrix-web run test:e2e -- e2e/multi-repo-scoping.spec.ts` | passed | 3 tests passed |
| P8-E repo switching e2e | `npm --prefix .\avmatrix-web run test:e2e -- e2e/repo-switching.spec.ts` | passed | 6 tests passed |
| P8-E onboarding e2e | `npm --prefix .\avmatrix-web run test:e2e -- e2e/onboarding.spec.ts` | passed with expected skip | 12 tests passed; 1 packaged-launcher case skipped by flag |
| P8-E shell interactions e2e | `npm --prefix .\avmatrix-web run test:e2e -- e2e/shell-interactions.spec.ts` | passed | 7 tests passed |
| P8-F fresh analyze benchmark | `.\avmatrix-launcher\server-bundle\avmatrix.exe analyze --force --benchmark-json .\.tmp\2026-05-22-p8-final-analyze.json --benchmark-label p8-final` | passed | files scanned 755, parsed 565, unsupported 190, failed 0; graph nodes 85599 and relationships 117535 |
| P8-F query-health final | `.\avmatrix-launcher\server-bundle\avmatrix.exe query-health --suite .\docs\query-health\2026-05-22-avmatrix-app-layer-resolution-gap-suite.json --repo AVmatrix --out .\.tmp\2026-05-22-p8-query-health-final.json --limit 10` | passed | 7 cases passed, 0 failed; final hit@5/hit@10 thresholds met for every case |
| P8-G resolution inventory final | `.\avmatrix-launcher\server-bundle\avmatrix.exe resolution-inventory --graph .\.avmatrix\graph.json --out .\.tmp\2026-05-22-p8-resolution-inventory-final.json` | passed | gapNodes 61519, gapRelationships 61519, gapOccurrences 62305, resolvedReferences 28395, inRepoAnalyzerGap 36304, unresolvedNonActionable 25833 |
| P8-H command-example analyze | `.\avmatrix-launcher\server-bundle\avmatrix.exe analyze --force --benchmark-json .\.tmp\2026-05-22-p8-command-examples-analyze.json --benchmark-label p8-command-examples` | passed | files scanned 755, parsed 565, unsupported 190, failed 0; graph nodes 85601 and relationships 117537 |
| P8-H query semantic example | `.\avmatrix-launcher\server-bundle\avmatrix.exe query "layout optimizer app layer ring graph filters" --repo AVmatrix --limit 10` | passed | artifact `.tmp\2026-05-22-p8-query-output.txt`; returned frontend/layout definitions with semantic fields |
| P8-H context semantic example | `.\avmatrix-launcher\server-bundle\avmatrix.exe context "applyFilterBasedClusteredLayout" --repo AVmatrix` | passed | artifact `.tmp\2026-05-22-p8-context-output.txt`; found frontend/layout symbol with source-site proof/status metadata |
| P8-H impact semantic example | `.\avmatrix-launcher\server-bundle\avmatrix.exe impact "applyFilterBasedClusteredLayout" --repo AVmatrix --direction upstream` | passed with expected CRITICAL warning | artifact `.tmp\2026-05-22-p8-impact-output.txt`; impactedCount 6, affectedAppLayers frontend=6, affectedFunctionalAreas layout=3 and web_graph_ui=3, `workflowWarningBlocksOutput=false` |
| P8-H/P8-I detect-changes semantic example | `.\avmatrix-launcher\server-bundle\avmatrix.exe detect-changes --repo AVmatrix --scope all` | passed with low doc-only scope | artifact `.tmp\2026-05-22-p8-detect-output.txt`; affected_count 0, changed_count 13, changed_files 3, risk_level low |
| P8-J closure update | plan/evidence/benchmark ledger update | complete | all required validation evidence recorded; final commit is doc-only under repository rules |
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
| Full build before ResolutionGap persistence tests | `powershell -ExecutionPolicy Bypass -File .\avmatrix-launcher\build.ps1` | passed | first run exposed missing Web constants for new generated contract values; final run passed after adding `ResolutionGap` and `HAS_RESOLUTION_GAP` Web constants/icon coverage |
| Backend/contract focused tests | `go test .\internal\graphhealth .\internal\semantic .\internal\lbugschema .\internal\lbugload .\internal\contracts .\internal\scopeir` | passed | focused persisted ResolutionGap model, semantic enrichment, LadybugDB schema/export, contract, and node-label coverage |
| Backend tests | `go test .\internal\...` and `go test .\cmd\...` | passed | wider Go validation after persisted ResolutionGap slice |
| Contract generation/check | `go run .\cmd\generate-web-contracts` and `go run .\cmd\generate-web-contracts --check` | passed | generated Web schema and TypeScript contract include `ResolutionGap` and `HAS_RESOLUTION_GAP` |
| Web unit tests | `npm --prefix .\avmatrix-web run test -- --run` | passed | 44 test files, 358 tests after adding filter-panel icon coverage for `ResolutionGap` |
| Web e2e tests | `npm --prefix .\avmatrix-web run test:e2e` | passed | 44 tests collected; 14 passed and 30 skipped under local mocked-backend conditions |
| Fresh analyze benchmark for ResolutionGap persistence | `.\avmatrix-launcher\server-bundle\avmatrix.exe analyze --force --benchmark-json .\.tmp\2026-05-22-p3-resolution-gap-persist-analyze.json --benchmark-label p3-resolution-gap-persist` | passed | graph nodes 80957, relationships 110949, ResolutionGap nodes 58083, HAS_RESOLUTION_GAP relationships 58083, LadybugDB skipped relationships 0 |
| AVmatrix detect-changes for ResolutionGap persistence | `.\avmatrix-launcher\server-bundle\avmatrix.exe detect-changes --repo AVmatrix --scope all` | passed with expected critical scope | affected_count 20, changed_count 348, changed_files 22, risk_level critical; affected surfaces match semantic enrichment, graphhealth ResolutionGap model, LadybugDB schema/export, generated contracts, Web constants/filter icon coverage, and plan ledgers |
| Full build before ResolutionGap role-validation tests | `powershell -ExecutionPolicy Bypass -File .\avmatrix-launcher\build.ps1` | passed | run before focused and wide tests after target-role inference and validator changes |
| Backend focused tests for ResolutionGap role validation | `go test .\internal\graphhealth .\internal\semantic` | passed | covers `InferredTargetRole`, `ValidateResolutionGapPersistence`, and broad semantic persistence fixtures |
| Backend tests | `go test .\internal\...` and `go test .\cmd\...` | passed | wider Go validation after ResolutionGap role inference and validation fixtures |
| Fresh analyze benchmark for ResolutionGap role validation | `.\avmatrix-launcher\server-bundle\avmatrix.exe analyze --force --benchmark-json .\.tmp\2026-05-22-p3-role-validation-postedit-analyze.json --benchmark-label p3-role-validation-postedit` | passed | graph nodes 81298, relationships 111411, ResolutionGap nodes 58350, HAS_RESOLUTION_GAP relationships 58350, LadybugDB skipped relationships 0 |
| Source-site accuracy after ResolutionGap role validation | `.\avmatrix-launcher\server-bundle\avmatrix.exe source-site-accuracy --graph .\.avmatrix\graph.json --out .\.tmp\2026-05-22-p3-role-validation-source-site-accuracy.json --max-examples 20` | passed | all source-site occurrences 144976, missing IDs 0, false resolved edge candidates 0, resolved edges without proof 0, non-property ACCESSES targets 0, coarse file call edges 0 |
| AVmatrix detect-changes for ResolutionGap role validation | `.\avmatrix-launcher\server-bundle\avmatrix.exe detect-changes --repo AVmatrix --scope all` | passed | after staging all slice files: affected_count 0, changed_count 344, changed_files 8, risk_level low |
| Full build before ResolutionGap aggregate-policy tests | `powershell -ExecutionPolicy Bypass -File .\avmatrix-launcher\build.ps1` | passed | run before focused and wide tests after adding aggregate/dedupe policy |
| Backend focused tests for ResolutionGap aggregation | `go test .\internal\graphhealth` | passed | covers exact occurrence counts, sourceSiteID traceability, sample caps, App Layer/Functional Area distributions, and target-text bucket identity |
| Backend tests | `go test .\internal\...` and `go test .\cmd\...` | passed | wider Go validation after ResolutionGap aggregate policy |
| Fresh analyze benchmark for ResolutionGap aggregate policy | `.\avmatrix-launcher\server-bundle\avmatrix.exe analyze --force --benchmark-json .\.tmp\2026-05-22-p3-aggregate-policy-postedit-analyze.json --benchmark-label p3-aggregate-policy-postedit` | passed | graph nodes 81522, relationships 111801, ResolutionGap nodes 58495, HAS_RESOLUTION_GAP relationships 58495, LadybugDB skipped relationships 0 |
| Source-site accuracy after ResolutionGap aggregate policy | `.\avmatrix-launcher\server-bundle\avmatrix.exe source-site-accuracy --graph .\.avmatrix\graph.json --out .\.tmp\2026-05-22-p3-aggregate-policy-source-site-accuracy.json --max-examples 20` | passed | all source-site occurrences 145390, missing IDs 0, false resolved edge candidates 0, resolved edges without proof 0, non-property ACCESSES targets 0, coarse file call edges 0 |
| AVmatrix detect-changes for ResolutionGap aggregate policy | `.\avmatrix-launcher\server-bundle\avmatrix.exe detect-changes --repo AVmatrix --scope all` | passed | after staging all slice files: affected_count 0, changed_count 228, changed_files 5, risk_level low |
| Full build before Resolution Health tests | `powershell -ExecutionPolicy Bypass -File .\avmatrix-launcher\build.ps1` | passed | first post-contract build exposed a generated Web type consumer that needed `resolutionConfidence`; final full build passed after updating `getNodeGraphHealth` fallback fields |
| Backend/contract focused tests for Resolution Health | `go test .\internal\graphhealth .\internal\cli .\internal\contracts .\internal\httpapi` | passed | covers Resolution Health overlay, `resolution-inventory` command output, generated contract fields, and HTTP graph/report/explain payload structs |
| Contract generation/check | `go run .\cmd\generate-web-contracts` and `go run .\cmd\generate-web-contracts --check` | passed | generated schema and TypeScript contract expose Resolution Health buckets and confidence fields |
| Web unit tests after Resolution Health contract update | `npm --prefix .\avmatrix-web run test -- --run` | passed | 44 files and 358 tests passed |
| Backend tests after Resolution Health slice | `go test .\internal\...` and `go test .\cmd\...` | passed | wider Go validation after Resolution Health overlay and CLI command |
| Web e2e after starting Vite dev server | `npm --prefix .\avmatrix-web run test:e2e` | passed | direct run without frontend server failed with `ERR_CONNECTION_REFUSED`; rerun with hidden Vite dev server passed 14 and skipped 30 |
| Fresh analyze benchmark for Resolution Health inventory | `.\avmatrix-launcher\server-bundle\avmatrix.exe analyze --force --benchmark-json .\.tmp\2026-05-22-p4-resolution-health-inventory-analyze.json --benchmark-label p4-resolution-health-inventory` | passed | graph nodes 82062, relationships 112614, ResolutionGap nodes 58879, HAS_RESOLUTION_GAP relationships 58879, LadybugDB skipped relationships 0 |
| Resolution inventory command | `.\avmatrix-launcher\server-bundle\avmatrix.exe resolution-inventory --graph .\.avmatrix\graph.json --out .\.tmp\2026-05-22-p4-resolution-inventory.json` | passed | gap occurrences 59626, resolved references 27195, analyzer-gap occurrences 34537, non-actionable occurrences 24926 |
| Source-site accuracy after Resolution Health slice | `.\avmatrix-launcher\server-bundle\avmatrix.exe source-site-accuracy --graph .\.avmatrix\graph.json --out .\.tmp\2026-05-22-p4-resolution-health-source-site-accuracy.json --max-examples 20` | passed | all source-site occurrences 146447, missing IDs 0, false resolved edge candidates 0, resolved edges without proof 0, non-property ACCESSES targets 0, coarse file call edges 0 |
| Cypher ResolutionGap inventory | `cypher` queries over `ResolutionGap` and `HAS_RESOLUTION_GAP` | passed | `ResolutionGap` node count 58879 and `HAS_RESOLUTION_GAP` relationship count 58879 visible in LadybugDB |
| Full build before Query Health tests | `powershell -ExecutionPolicy Bypass -File .\avmatrix-launcher\build.ps1` | passed | run before final focused tests after adding `query-health` command and suite |
| Backend focused tests for Query Health | `go test .\internal\cli` and `go test .\internal\mcp` | passed | covers query-health suite parsing/scoring/output/threshold behavior and confirms existing MCP query path still passes |
| Fresh analyze benchmark for Query Health command | `.\avmatrix-launcher\server-bundle\avmatrix.exe analyze --force --benchmark-json .\.tmp\2026-05-22-p5-query-health-command-analyze.json --benchmark-label p5-query-health-command` | passed | graph nodes 82538 and relationships 113396 after adding the command, tests, and suite |
| Query-health command | `.\avmatrix-launcher\server-bundle\avmatrix.exe query-health --suite .\docs\query-health\2026-05-22-avmatrix-app-layer-resolution-gap-suite.json --repo AVmatrix --out .\.tmp\2026-05-22-p5-query-health-baseline.json --limit 10` | passed | command completed and wrote baseline report; 7 cases, 1 passed, 6 failed |
| AVmatrix detect-changes for Query Health command | `.\avmatrix-launcher\server-bundle\avmatrix.exe detect-changes --repo AVmatrix --scope all` | passed with expected critical scope | affected_count 22, changed_count 493, changed_files 8, risk_level critical; affected surfaces match root CLI registration and new `runQueryHealth` flows |
| Resolution inventory command | `.\avmatrix-launcher\server-bundle\avmatrix.exe resolution-inventory --graph .\.avmatrix\graph.json --out .\.tmp\2026-05-22-p4-resolution-inventory.json` | passed | implemented in Phase 4; duplicate retained here until P6 command-surface metrics are finalized |
| `query` semantic output | `.\avmatrix-launcher\server-bundle\avmatrix.exe query "query ranking process matching definitions CLI query command implementation" --repo AVmatrix --limit 10` | passed | output includes complete `semanticStatus`, App Layer, Functional Area, node type, ResolutionGap summaries, and Resolution Health fields where available |
| Full build before P6-A query tests | `powershell -ExecutionPolicy Bypass -File .\avmatrix-launcher\build.ps1` | passed | run before focused MCP/CLI tests after query ranking/output changes |
| Backend focused tests for P6-A query output | `go test .\internal\mcp .\internal\cli` | passed | covers query semantic fields, stale semantic warning, existing query path, and query-health scoring/output |
| Fresh analyze benchmark for P6-A query output | `.\avmatrix-launcher\server-bundle\avmatrix.exe analyze --force --benchmark-json .\.tmp\2026-05-22-p6a-query-semantic-output-final-analyze.json --benchmark-label p6a-query-semantic-output-final` | passed | graph nodes 82935 and relationships 113953 after query ranking/output changes |
| Query-health after P6-A query output | `.\avmatrix-launcher\server-bundle\avmatrix.exe query-health --suite .\docs\query-health\2026-05-22-avmatrix-app-layer-resolution-gap-suite.json --repo AVmatrix --out .\.tmp\2026-05-22-p6a-query-health-final.json --limit 10` | passed | 7 cases, 7 passed, 0 failed |
| AVmatrix detect-changes for P6-A query output | `.\avmatrix-launcher\server-bundle\avmatrix.exe detect-changes --repo AVmatrix --scope all` | passed with expected critical scope | affected_count 17, changed_count 432, changed_files 5, risk_level critical; affected surfaces match queryTool, rankedProcessMatches, matchingDefinitionRows, and related query retrieval flows |
| `context` semantic output | `.\avmatrix-launcher\server-bundle\avmatrix.exe context --repo AVmatrix --uid "Function:internal/mcp/context.go:contextSymbolPayload#2"` and `.\avmatrix-launcher\server-bundle\avmatrix.exe context --repo AVmatrix --uid "ResolutionGap:SourceSite:internal/mcp/context.go#call#any#197#12#204#2"` | passed | output artifacts `.tmp\2026-05-22-p6b-context-symbol-output.txt` and `.tmp\2026-05-22-p6b-context-resolution-gap-output.txt` show semanticStatus, App Layer, Functional Area, source-site proof/status fields, sourceResolutionGaps, and separate ResolutionGap entity/source rows |
| AVmatrix detect-changes for P6-B context output | `.\avmatrix-launcher\server-bundle\avmatrix.exe detect-changes --repo AVmatrix --scope all` | passed with expected critical scope | affected_count 26, changed_count 87, changed_files 5, risk_level critical; affected surfaces match contextToolInternal, contextSymbolPayload, contextRefPayload, contextNeighborhood, semantic helper rows, and context candidate output flows |
| `impact` semantic output | `.\avmatrix-launcher\server-bundle\avmatrix.exe impact --repo AVmatrix --uid "Function:internal/mcp/context.go:contextSymbolPayload#2" --direction upstream` and `.\avmatrix-launcher\server-bundle\avmatrix.exe impact --repo AVmatrix --uid "Function:internal/mcp/impact.go:runImpactBFSProfiled#4" --direction upstream` | passed | output artifacts `.tmp\2026-05-22-p6c-impact-context-symbol-output.txt` and `.tmp\2026-05-22-p6c-impact-critical-warning-output.txt` show affectedAppLayers, affectedFunctionalAreas, resolutionHealthRisks, semantic fields, and `workflowWarningBlocksOutput=false` for CRITICAL output |
| AVmatrix detect-changes for P6-C impact output | `.\avmatrix-launcher\server-bundle\avmatrix.exe detect-changes --repo AVmatrix --scope all` | passed with expected high scope | affected_count 9, changed_count 127, changed_files 4, risk_level high; affected surfaces match runImpactBFSProfiled, impactItemPayload, impactAffectedProcesses, impactAffectedModules, and semantic impact summary helper output flows |
| `detect-changes` semantic output | `.\avmatrix-launcher\server-bundle\avmatrix.exe detect-changes --repo AVmatrix --scope all` | passed | output artifact `.tmp\2026-05-22-p6d-detect-changes-semantic-output.txt` shows semanticStatus, changedAppLayers, changedFunctionalAreas, affectedAppLayers, affectedFunctionalAreas, changedStepAppLayers, resolutionGapEntity rows, resolutionGapChanges, and resolutionHealthImpact; implementation-only P6-D diff reported changed_count 156, affected_count 15, changed_files 1, risk_level high, changedGapEntities 94, and changedGapOccurrenceCount 95 |
| AVmatrix detect-changes for P6-D detect-changes output | `.\avmatrix-launcher\server-bundle\avmatrix.exe detect-changes --repo AVmatrix --scope all` | passed with expected high scope | after ledger updates: affected_count 15, changed_count 164, changed_files 4, risk_level high; affected surfaces match detectChangesTool, detectChangedSymbols, detectAffectedProcesses, semantic changed-symbol/process summary helpers, and plan ledgers |
| API-specific MCP semantic output | MCP `tools/call` for `route_map`, `shape_check`, and `api_impact` against repo `AVmatrix` | passed | artifacts `.tmp\2026-05-22-p6e-route-map-semantic-output.txt`, `.tmp\2026-05-22-p6e-shape-check-semantic-output.txt`, and `.tmp\2026-05-22-p6e-api-impact-semantic-output.txt` show semanticStatus on current empty/error API-tool payloads; focused fixture test proves populated route, consumer, flowDetail, executionFlowDetails, and impactSummary semantic fields |
| AVmatrix detect-changes for P6-E API MCP output | `.\avmatrix-launcher\server-bundle\avmatrix.exe detect-changes --repo AVmatrix --scope all` | passed with expected critical scope | affected_count 29, changed_count 273, changed_files 6, risk_level critical; affected surfaces match route index building, route consumers, flow details, API impact shaping, semantic helper rows, route fixture tests, and plan ledgers |
| Full build before P6-F semantic edge tests | `powershell -ExecutionPolicy Bypass -File .\avmatrix-launcher\build.ps1` | passed | run before focused MCP edge tests after adding stale/missing semantic metadata coverage |
| Focused P6-F semantic edge tests | `go test .\internal\mcp -run "Test(ImpactToolWarnsForStaleIncompleteSemanticMetadata|DetectChangesToolWarnsForStaleIncompleteSemanticMetadata|APIMCPToolsWarnForStaleAndDoNotInventSemanticFields)$" -count=1` | passed | covers stale/incomplete warnings and no invented App Layer fields for impact, detect-changes, route_map, shape_check, and api_impact |
| Backend focused tests for P6-F command surfaces | `go test .\internal\mcp .\internal\cli -count=1` | passed | verifies existing MCP/CLI semantic command-surface tests continue to pass with edge coverage |
| AVmatrix detect-changes for P6-F command edge tests | `.\avmatrix-launcher\server-bundle\avmatrix.exe detect-changes --repo AVmatrix --scope all` | passed with expected low scope | after staging the test and ledger files: affected_count 0, changed_count 125, changed_files 4, risk_level low; changed_app_layers api_test 117 and docs 8; changedGapEntities 92 from test-fixture graph evidence |
| AVmatrix detect-changes for implementation commits | `.\avmatrix-launcher\server-bundle\avmatrix.exe detect-changes --repo AVmatrix --scope all` | passed with expected critical scope | affected_count 24, changed_count 144, changed_files 18; affected surfaces match Functional Area semantic/schema/API/contract/export slice |

## B12 - Semantic Enrichment Flow Metrics

Status: complete; final semantic enrichment metrics recorded from P8 final analyze.

Record after the analyze semantic enrichment phase is introduced and after each phase extends it.

| Metric | Baseline | After App Layer | After Functional Area | After Source-Site Inventory | After ResolutionGap | After Role Validation | After Aggregate Policy | After Resolution Health | Final |
| --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| Semantic enrichment phase latency | pending | 33.9997 ms | 68.8814 ms | pending | 627.8377 ms | 609.862 ms | 595.5799 ms | 586.7456 ms | 763.8928 ms |
| Semantic enrichment memory delta | pending | pending | overall analyze heap +89525848 bytes; not phase-isolated | pending | overall analyze heap +256162968 bytes; not phase-isolated | overall analyze heap +256081936 bytes; not phase-isolated | overall analyze heap +241109064 bytes; not phase-isolated | overall analyze heap +423902888 bytes; not phase-isolated | overall analyze heap +431057664 bytes; max observed sys 701792504 bytes; not phase-isolated |
| Graph nodes before enrichment | pending | 22239 | 22358 | pending | 22874 inferred from after minus ResolutionGap nodes | 22948 inferred from after minus ResolutionGap nodes | 23027 inferred from after minus ResolutionGap nodes | 23183 inferred from after minus ResolutionGap nodes | 24080 inferred from after minus ResolutionGap nodes |
| Graph nodes after enrichment | pending | 22239 | 22358 | pending | 80957 | 81298 | 81522 | 82062 | 85599 |
| Graph relationships before enrichment | pending | 55006 | 55349 | pending | 52866 inferred from after minus HAS_RESOLUTION_GAP relationships | 53061 inferred from after minus HAS_RESOLUTION_GAP relationships | 53306 inferred from after minus HAS_RESOLUTION_GAP relationships | 53735 inferred from after minus HAS_RESOLUTION_GAP relationships | 56016 inferred from after minus HAS_RESOLUTION_GAP relationships |
| Graph relationships after enrichment | pending | 55006 | 55349 | pending | 110949 | 111411 | 111801 | 112614 | 117535 |
| Graph JSON size | pending | 45739916 bytes | 47953614 bytes | pending | 274996080 bytes | 276339686 bytes | 277198593 bytes | 279176550 bytes | 291894554 bytes |
| LadybugDB node rows | pending | 22239 | 22358 | pending | 80957 | 81298 | 81522 | 82062 | 85599 |
| LadybugDB relationship rows | pending | 55006 | 55349 | pending | 110949 | 111411 | 111801 | 112614 | 117535 |
| Duplicate graph traversals introduced | pending | 0 nested whole-graph loops; one node pass and one relationship pass | 0 nested whole-graph loops; one node pass and one relationship pass shared by App Layer and Functional Area | pending | 0 nested loops; adds one source-backed diagnostic node pass and per-gap AddNode/AddRelationship finalization | 0 nested loops; target-role inference is per gap and validator runs only in tests/explicit checks | 0 analyze traversals; aggregation is explicit inventory work outside analyze semantic enrichment | 0 analyze traversals; Resolution Health inventory runs in `ComputeSummary`/CLI/API from persisted graph data | 0 nested loops; semantic phase scanned 117535 relationships and finalized 61519 gap nodes/relationships from existing graph facts |
| File rescans introduced | pending | 0 | 0 | pending | 0 | 0 | 0 | 0 | 0 |
| AST reparses introduced | pending | 0 | 0 | pending | 0 | 0 | 0 | 0 | 0 |
| Raw call/access source-site count | pending | pending | pending | pending | call 31220, access 18245 persisted gap nodes | call 31307, access 18383 persisted gap nodes | call 31393, access 18418 persisted gap nodes | call 31590, access 18541 persisted gap nodes | call 32900, access 19348 persisted gap nodes |
| Raw unresolved fact count | pending | pending | pending | pending | 58083 persisted sourceSiteID-backed gaps | 58350 persisted sourceSiteID-backed gaps | 58495 persisted sourceSiteID-backed gaps | 58879 persisted sourceSiteID-backed gaps | 61519 persisted sourceSiteID-backed gaps |

## B13 - Phase 9 Non-Actionable Breakdown And Diagnostic Shape

Status: complete after P9-I full e2e closure validation.

Starting graph refresh:

```powershell
.\avmatrix-launcher\server-bundle\avmatrix.exe analyze --force --benchmark-json .\.tmp\2026-05-23-p9-non-actionable-breakdown-start-analyze.json --benchmark-label p9-non-actionable-breakdown-start
```

| Metric | Value |
| --- | ---: |
| Files scanned | 755 |
| Files parsed | 565 |
| Unsupported files | 190 |
| Failed files | 0 |
| Graph nodes | 85604 |
| Graph relationships | 117540 |

Phase 8 non-actionable baseline to preserve and clarify:

| Metric | Value |
| --- | ---: |
| `unresolvedNonActionable` total | 25833 |
| `builtin` classification | 10723 |
| `standard_library` classification | 7674 |
| `test_framework` classification | 7436 |

Final graph refresh after implementation:

```powershell
.\avmatrix-launcher\server-bundle\avmatrix.exe analyze --force --benchmark-json .\.tmp\2026-05-23-p9-non-actionable-breakdown-final-analyze.json --benchmark-label p9-non-actionable-breakdown-final
```

| Metric | Value |
| --- | ---: |
| Files scanned | 756 |
| Files parsed | 566 |
| Unsupported files | 190 |
| Failed files | 0 |
| Graph nodes | 85732 |
| Graph relationships | 117678 |

Validation and final metrics:

| Metric | Result |
| --- | --- |
| CLI breakdown output | focused fixture expects `resolutionHealth.unresolvedNonActionableBreakdown=builtin:2,standard_library:4,test_framework:5` |
| Web lens row count delta | `getResolutionLensRows` changed from `11` rows to `13` rows by replacing one collapsed non-actionable row with three classification rows |
| ResolutionGap square-size unit coverage | `graph-adapter.edge-geometry.test.ts` asserts `nodeType=ResolutionGap`, `type=square`, and `size=1` |
| Focused test output | `go test .\internal\cli -run "TestResolutionInventoryCommandOutputsJSON" -count=1` passed; Web focused tests passed `3` files / `33` tests |
| Full build output | `powershell -ExecutionPolicy Bypass -File .\avmatrix-launcher\build.ps1` passed before tests |
| Backend/cmd validation output | `go test .\internal\... .\cmd\...` passed |
| Contract validation output | `go run .\cmd\generate-web-contracts --check` passed with no generated contract drift |
| Full Web unit output | `npm --prefix .\avmatrix-web run test -- --run` passed `45` files / `369` tests |
| Web/browser validation output | focused browser/e2e passed: `server-connect.spec.ts -g "reports Backend API Frontend rings"` passed `1` test; `graph-health-ui.spec.ts` passed `4` tests. Full e2e and full `server-connect.spec.ts` timed out in validation and are not recorded as passed. |
| P9-I fresh full build gate | passed before browser validation |
| P9-I all-spec e2e command | attempted `npm --prefix .\avmatrix-web run test:e2e`; timed out at `1804044 ms` and is not counted as passing evidence |
| P9-I full `server-connect.spec.ts` | `10/10` passed in `11.1m` |
| P9-I `graph-health-ui.spec.ts` | `4/4` passed in `20.9s` |
| P9-I `heartbeat-reconnect.spec.ts` | `2/2` passed in `1.9m` |
| P9-I `multi-repo-scoping.spec.ts` | `3/3` passed in `3.6m` |
| P9-I `repo-switching.spec.ts` | `6/6` passed in `8.8m` |
| P9-I `onboarding.spec.ts` | `12` passed, `1` expected packaged-launcher skip in `2.4m` |
| P9-I `shell-interactions.spec.ts` | `7/7` passed in `6.3m` |
| P9-I full e2e closure total | `44` passed, `0` failed, `1` expected skip |
| Final resolution inventory | `nodes=85732`, `relationships=117678`, `gapNodes=61625`, `gapRelationships=61625`, `gapOccurrences=62411`, `resolvedReferences=28419`, `unresolvedNonActionable=25841`, breakdown `builtin=10725`, `standard_library=7677`, `test_framework=7439` |
| Pre-commit detect-changes | after staging the implementation slice, `detect-changes --repo AVmatrix --scope all` reported `changed_count=163`, `changed_files=13`, `affected_count=4`, `risk_level=medium`, affected app layer `frontend=4`, changed app layers `backend=10`, `backend_test=17`, `docs=10`, `frontend=53`, `frontend_test=73` |

## B14 - Query Health Threshold And Exact Coverage

Status: complete.

Metric contract:

| Metric | Meaning |
| --- | --- |
| Threshold pass | hit@5 and hit@10 meet the suite thresholds; this is the usable-retrieval result. |
| Exact pass | every expected file/symbol target is found; this is complete expected-target coverage. |
| Legacy `passed` | compatibility field equal to threshold pass. |
| Matched targets | expected file/symbol targets matched by top query results. |
| Missed targets | expected file/symbol targets not matched by top query results. |

Validation and live command metrics:

| Metric | Value |
| --- | ---: |
| Full build before tests | passed |
| Focused Go tests | passed |
| Live query-health threshold-passed cases | 7 |
| Live query-health threshold-failed cases | 0 |
| Live query-health exact-passed cases | 2 |
| Live query-health exact-failed cases | 5 |
| Live query-health matched targets | 33 |
| Live query-health expected targets | 39 |
| Live query-health missed targets | 6 |
| Query-health artifact | `.tmp\2026-05-23-p10-query-health-threshold-exact.json` |
| Fresh analyze before commit | `nodes=85822`, `relationships=117824` |
| Pre-commit detect-changes risk | `high` |
| Pre-commit detect-changes changed count | 152 |
| Pre-commit detect-changes changed files | 12 |
| Pre-commit detect-changes affected count | 12 |
| Pre-commit detect-changes artifact | `.tmp\2026-05-23-p10-detect-changes.txt` |

Exact-miss inventory:

| Case | Threshold | Exact | Missed targets |
| --- | --- | --- | --- |
| unresolved-reference-diagnostic-generation | PASS | FAIL | `resolveCall`, `AppendDiagnosticToNode` |
| graph-health-unknown-connectivity-separation | PASS | PASS | none |
| app-layer-resolution-gap-layout | PASS | PASS | none |
| runtime-reset-hidden-terminal | PASS | FAIL | `startRuntime` |
| api-contract-surfaces | PASS | FAIL | `avmatrix-web/src/generated/avmatrix-contracts.ts` |
| query-implementation-surfaces | PASS | FAIL | `cmd/avmatrix/main.go` |
| frontend-graph-filter-surfaces | PASS | FAIL | `GraphCanvas` |
