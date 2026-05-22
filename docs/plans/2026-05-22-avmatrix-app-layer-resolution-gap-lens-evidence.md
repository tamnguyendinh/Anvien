# AVmatrix App Layer Resolution Gap Lens Evidence Ledger

Date: 2026-05-22

Status: in progress; Phase 0 closure audit complete; Phase 2 complete; Phase 2A proof-based CALLS/ACCESSES and source-site bridge slices complete; Phase 3 complete; Phase 4 complete; Phase 5 complete; Phase 6 complete

Plan: [2026-05-22-avmatrix-app-layer-resolution-gap-lens-plan.md](2026-05-22-avmatrix-app-layer-resolution-gap-lens-plan.md)

Benchmark: [2026-05-22-avmatrix-app-layer-resolution-gap-lens-benchmark.md](2026-05-22-avmatrix-app-layer-resolution-gap-lens-benchmark.md)

## E0 - Source Discussion Record

Status: recorded

Primary discussion record:

- [reports/problem/2026-05-22-avmatrix-app-layer-resolution-gap-discussion.md](../../reports/problem/2026-05-22-avmatrix-app-layer-resolution-gap-discussion.md)

Key decisions already captured:

- node type labels describe symbol shape, not product ownership;
- graph/API data must answer the semantic questions before Web UI can render them;
- App Layer must be one primary category, with mixed categories as separate values instead of overlapping labels;
- API is a first-class App Layer and may need separate API contract/client/shared categories;
- Functional Area requires an accuracy gate and must not be guessed from weak evidence;
- ResolutionGap/UnresolvedSymbol must persist into graph data;
- fine-grained gap relation types or typed metadata are preferred;
- Resolution Health and Topology Health must remain separate;
- query-health should become a command, not a one-off report;
- query/context/impact/detect-changes should surface the new semantic layers when available;
- API-specific MCP tools should surface the semantic layers where applicable because API is a first-class layer;
- resolved `CALLS` and `ACCESSES` must be proof-based; unproven source sites must become unresolved/ambiguous/external/dynamic/unsupported facts instead of guessed edges;
- every syntactic call/access source site needs inventory so unresolved facts are not silently lost;
- graph-based work starts from fresh analyze output; no stale graph fallback is a product path;
- accuracy is more important than minimizing graph size.
- analyze semantic enrichment must preserve both correctness and speed; neither may be traded away silently.
- aggregation or dedupe must not discard evidence only to make the graph smaller.

## E0A - Planning Codebase Audit Findings

Status: recorded

This audit was performed after the initial plan commit to make the plan match the current source tree. It is planning evidence only; facts not already superseded by recorded Phase 1/Phase 2 evidence must still be re-verified before the implementation slice that depends on them.

Source facts that must shape implementation:

- `avmatrix analyze --force` on 2026-05-22 scanned `728` files, parsed `539`, reported `189` unsupported, `0` failed, and produced `22095` nodes and `54772` relationships at `.avmatrix/graph.json`.
- A follow-up planning review with `avmatrix analyze --force` on 2026-05-22 scanned `732` files, parsed `543`, reported `189` unsupported, `0` failed, and produced `22361` nodes and `55352` relationships. Baseline counts must still be refreshed at implementation start because graph counts drift as code changes.
- `internal/resolution/emit.go` emits unresolved references through `emitUnresolvedReference`, then attaches them to source nodes with `graphhealth.AppendDiagnosticToNode`.
- `internal/resolution/resolve.go` emits unresolved heritage, call, access, and type-reference diagnostics.
- `internal/resolution/resolve.go` also emits resolved `CALLS` through fallback paths that include `resolveGlobalCallName` with confidence `0.5`. This is the code path that must stop becoming a resolved edge unless Phase 2A defines and proves an accepted binding.
- `internal/resolution/indexes.go` contains the resolver indexes and labels that must be audited for proof kinds: `resolveGlobalCallName`, `resolveSameFileName`, `resolveGoSamePackageFunction`, `resolveMember`, `resolveImportedMember`, `callableLabels`, `propertyLabels`, and dispatch-owner indexes.
- `internal/resolution/emit.go` dedupes semantic relationships with `semanticEdgeKey`; source-site inventory must be recorded before this dedupe can hide distinct occurrences.
- `internal/graph/types.go` `Relationship` has evidence, confidence, and resolution-source fields, but does not currently have sourceSiteID, source-site status, proof kind, or source range fields.
- `internal/scopeir/facts.go` defines `CallSiteFact` and `AccessFact`, and provider collectors/tests under `internal/providers/*/references.go`, `internal/providers/*/extract_test.go`, and `internal/providers/provider_parity_test.go` emit or validate those facts across languages. The source-site inventory must start from these provider facts rather than reconstructing call/access sites after the graph has already merged relationships.
- `internal/resolution/indexes.go` `propertyLabels()` and `internal/contracts/web_ui.go` provider fact coverage currently include `Property`, `Variable`, `Const`, and `Static` as ACCESSES-like labels. Phase 2A must either restrict resolved ACCESSES to proven property/field targets or split non-property uses into a separate relation/fact role.
- `internal/graphaccuracy/graphaccuracy.go` currently focuses on Go definitions/imports and a direct CALLS subset; it does not yet provide ACCESSES precision metrics, source-site inventory metrics, false resolved edge detection, or the cross-language `stop()` false-positive fixture.
- `internal/graphhealth/diagnostics.go` stores diagnostics under `graphHealthDiagnostics`; `sameDiagnosticBucket` currently does not include `TargetText`, so different unresolved target texts can collapse into one bucket.
- `internal/graphhealth/policy.go` and `internal/graphhealth/diagnostics.go` already define diagnostic classification/actionability for builtin, standard-library, test-framework, external-library, in-repo unresolved, unclassified, non-actionable, review, and analyzer-gap.
- `internal/httpapi/graph.go` calls graph-health summary computation when building graph responses, and its graph-health report candidate limit is capped. Full ResolutionGap inventory must not rely on that capped report path.
- `internal/analyze/analyze.go` currently runs resolution, MRO, communities, processes, graph compaction, LadybugDB load, and graph snapshot writing in that order. Semantic enrichment must run before LadybugDB load and graph snapshot writing, and after any upstream signal it depends on.
- The target flow is resolution with raw unresolved fact capture, then MRO, communities, processes, semantic enrichment, graph compaction, LadybugDB load, and graph snapshot. Raw unresolved facts should keep `sourceNodeID`, `factFamily`, `targetText`, `filePath`, range or line, `resolutionSource`, and note; call/access facts from Phase 2A must also keep `sourceSiteID`, source-site status, proof kind, and target role when known.
- Semantic enrichment should use already-produced facts and reusable indexes such as `nodeID -> node`, `filePath -> App Layer`, `nodeID -> process/community`, and `sourceNodeID -> raw gaps`; it should avoid file rescans, AST reparses, and nested whole-graph loops.
- `internal/lbugload` exports the in-memory graph into LadybugDB. New semantic fields, entities, or relationships need export/load coverage when query or Cypher consumers are expected to see them.
- `internal/mcp` already contains API-specific tools including `route_map`, `shape_check`, and `api_impact`; these need inclusion or explicit limitation because API is now a first-class App Layer.
- `internal/mcp/tools.go` currently ranks `query` results with simple process/step contains scoring, and definition matching is narrow enough that function/method-heavy intents can miss the expected files.
- `internal/contracts/web_ui.go` is the generated Web contract source; Web TypeScript contract shape should be generated from this source rather than hand-edited.
- `avmatrix-web/src/lib/graph-adapter.ts` currently applies deterministic filter-based clustered layout during graph conversion and already caps rendered node size at `3`.
- `avmatrix-web/src/hooks/useSigma.ts` invokes layout work from the manual optimizer path; the current plan should change deterministic initial placement, not add automatic optimizer execution.
- `avmatrix-web/src/hooks/app-state/graph.tsx`, `avmatrix-web/src/lib/graph-health-filters.ts`, `avmatrix-web/src/components/FileTreePanel.tsx`, and `avmatrix-web/src/components/GraphCanvas.tsx` are current Web state/filter/layout surfaces and do not yet have App Layer or Resolution Health filter state.

Query audit sample from the current codebase:

| Intent | Expected area | Observed top noise |
| --- | --- | --- |
| unresolved reference diagnostic graph health resolution gap | `internal/resolution/resolve.go`, `internal/resolution/emit.go`, `internal/graphhealth/diagnostics.go` | launcher `HiddenProcAttr`, Web `handleKeyDown`, backend-client process steps |
| query ranking process matching definitions | `internal/mcp/tools.go`, CLI query command surfaces | launcher `HiddenProcAttr`, Web `handleKeyDown`, generated contract process |
| CALLS ACCESSES resolver proof and source-site inventory | `internal/resolution/resolve.go`, `internal/resolution/emit.go`, `internal/resolution/indexes.go`, `internal/scopeir/facts.go`, `internal/graphaccuracy/graphaccuracy.go` | launcher/runtime and Web process matches instead of resolver/source-site files |

## E1 - Discussion-To-Plan Coverage

Status: complete for Phase 0 closure audit.

Recorded during P0-A from the discussion record and current plan.

| Discussion area | Plan task | Evidence result |
| --- | --- | --- |
| Node type is insufficient | Problem, Design Decisions, P1/P2/P3 | Covered: node labels are symbol shape, not product ownership. |
| Graph/API, not UI, must answer ownership | Problem, Scope Boundary, P1/P3/P6/P7 | Covered: Web UI must render persisted graph/API truth, not infer ownership client-side. |
| BE/API/FE/App Layer rings | Design Decisions, P1, P7 | Covered: App Layer is macro placement ring; API is first-class; Web rings remain Phase 7. |
| Non-overlapping mixed App Layer categories | Design Decisions, P1-A/P1-B | Covered: one primary category; mixed concerns become separate category values. |
| API as first-class layer | Design Decisions, P1/P6 | Covered: API, API contract, API shared contract, frontend API client, and API test categories exist. |
| Functional Area accuracy gate | Design Decisions, P2 | Covered: only high-confidence deterministic rules are accepted; ambiguous nodes remain `unknown`. |
| Proof-based CALLS/ACCESSES and source-site inventory | Problem, Design Decisions, Phase 2A | Covered: resolved edges need proof; source sites are inventoried even when unresolved. |
| Persisted ResolutionGap/UnresolvedSymbol | Design Decisions, P3 | Covered: persisted graph records are required; virtual UI-only data is out of scope. |
| Fine-grained gap relations | Acceptance Criteria, P3-C | Covered: call/access/type-reference/heritage/external/builtin/test/analyzer-gap distinctions are planned. |
| Resolution Health separate from Topology Health | Design Decisions, P3/P4 | Covered: resolution confidence does not overwrite topology status. |
| Query-health command | Design Decisions, P5 | Covered: query-health must be a repeatable command, not a one-off report. |
| Query/context/impact/detect-changes semantic output | Scope Boundary, P6 | Covered: command output must surface App Layer, Functional Area, and ResolutionGap when persisted. |
| API-specific MCP semantic output | Scope Boundary, P6 | Covered: route_map, shape_check, and api_impact are included because API is first-class. |
| Multi-ring layout and same-color islands | Design Decisions, P7 | Covered: App Layer controls macro rings; node type/gap kind controls islands inside each ring. |
| No dead-code verdict from unresolved refs alone | Scope Boundary | Covered: unresolved references are not treated as confirmed dead code. |
| No timeout/auto optimizer behavior | Rules of plan, Scope Boundary, P7 | Covered: no product/runtime timeouts, delayed refresh, elapsed-time budget, or automatic optimizer runs. |
| No stale graph fallback | Rules of plan, Acceptance Criteria | Covered: stale/missing semantic metadata is incomplete evidence, not a trigger for API/UI guessing. |
| No evidence loss for graph-size reasons | Rules of plan, Acceptance Criteria | Covered: aggregation/dedupe may only preserve counts, samples, and traceability. |
| User-facing naming consistency | P1-H/P3/P7 | Covered: naming registry exists for semantic terms; later phases extend it as new gap categories land. |

## E2 - Baseline Analyze Evidence

Status: complete for Phase 0 closure audit.

Recorded during P0-B.

Command used:

```powershell
.\avmatrix-launcher\server-bundle\avmatrix.exe analyze --force --benchmark-json .tmp\2026-05-22-phase0-baseline.json --benchmark-label phase0-baseline
```

Command output:

```text
analyzed E:\AVmatrix-GO
files: scanned=736 parsed=547 unsupported=189 failed=0
graph: nodes=22635 relationships=52144 path=E:\AVmatrix-GO\.avmatrix\graph.json
```

Index status after analyze:

```text
Repository: E:\AVmatrix-GO
Indexed: 5/22/2026, 1:44:46 PM
Indexed commit: 7c5b858
Current commit: 7c5b858
Status: up-to-date
```

Fresh graph evidence:

| Metric | Value |
| --- | ---: |
| Files scanned | 736 |
| Files parsed | 547 |
| Unsupported files | 189 |
| Failed files | 0 |
| Graph nodes | 22635 |
| Graph relationships | 52144 |
| Counted semantic relationships | 23437 |
| Execution flow nodes | 645 |
| `unknown_connectivity` nodes in graph snapshot | 0 |
| Graph path | `.avmatrix/graph.json` |
| Graph size | 108027212 bytes |
| Graph timestamp | 2026-05-22 13:44:46 +07:00 |
| Graph SHA-256 | `DB28BF1D99D0CFEEC860840AE3921A878DA7B20086481F9139C437D9112F9432` |

Comparison with discussion observations:

| Metric | Discussion observation | Fresh baseline | Note |
| --- | ---: | ---: | --- |
| Graph nodes | about 22010 | 22635 | Drift expected after implementation slices. |
| Counted semantic relationships | about 26906 | 23437 | Current counted policy excludes structural edges and reflects proof-based CALLS gates already implemented. |
| `unknown_connectivity` | 0 | 0 | Still separated from unresolved diagnostics. |
| Unresolved occurrences | about 51232 | 58195 | Increased after source-site inventory and low-confidence fallback demotion. |
| Unresolved buckets | about 8880 | 57449 | Buckets now preserve source-site identity; old collapsed bucket count is no longer comparable. |

## E3 - Source Trace Evidence

Status: complete for Phase 0 closure audit.

Recorded during P0-C and P0-G. Items marked for later phases are implementation work still covered by the plan, not missing Phase 0 discovery.

| Surface | Files/symbols found | Notes |
| --- | --- | --- |
| graph schema/snapshot | `internal/graph/types.go`, `internal/analyze/analyze.go`, `.avmatrix/graph.json` | App Layer is persisted on node properties as `appLayer`/`appLayerSource`; Functional Area is persisted as `functionalArea`/`functionalAreaSource`. Fresh graph snapshot from the locally built CLI has zero missing Functional Area fields across 22358 nodes. |
| analyze semantic enrichment flow | `internal/analyze/analyze.go`, `internal/semantic/app_layer.go`, `internal/semantic/functional_area.go` | The `semantic_enrichment` phase runs after processes and before graph compact, LadybugDB load, embeddings, and graph snapshot. Phase order from benchmark: scan, structure, documents, cobol, parse, routes, tools, orm, cross_file_binding, resolution, mro, communities, processes, semantic_enrichment, db_load. Functional Area assignment runs in the same enrichment pass as App Layer after process/community signals exist. |
| semantic enrichment input indexes and complexity | `internal/semantic/app_layer.go`, `internal/semantic/functional_area.go` | Enrichment builds App Layer and Functional Area path caches, `nodeID -> index`, `nodeID -> appLayer`, and `nodeID -> functionalArea` maps, then performs one relationship scan for Process/Community inference. It uses graph facts only; it does not rescan files or reparse ASTs. |
| LadybugDB export/load | `internal/lbugschema/schema.go`, `internal/lbugload/csv.go`, `internal/lbugload/load_test.go` | Node schemas and COPY CSV columns include both `appLayer` and `functionalArea`; benchmark DB load wrote 22358 node rows and 55349 relationship rows with zero fallback inserts. |
| resolved/unresolved call emission | `internal/resolution/resolve.go`, `internal/resolution/emit.go`, `internal/resolution/indexes.go`, `internal/resolution/source_site.go` | `resolveCall` emits source-backed unresolved diagnostics for missing source scope, file-level call source, unresolved target, and low-confidence global fallback. Proven calls emit `CALLS` with `sourceSiteID`, `sourceSiteStatus=resolved`, proof kind, target role, target text, file path, and range. |
| resolved/unresolved access emission | `internal/resolution/resolve.go`, `internal/resolution/emit.go`, `internal/resolution/indexes.go`, `internal/resolution/access_audit.go` | `resolveAccess` resolves receiver/imported members through `propertyLabels()` and emits unresolved access diagnostics with source-site metadata when property proof is absent. Current source-site report shows `ACCESSES` targets are `Property` only. |
| call/access source facts | `internal/scopeir/facts.go`, `internal/providers/*/references.go`, `internal/providers/*/extract_test.go`, `internal/providers/provider_parity_test.go` | Provider facts carry call/access name, receiver, scope, file, hash, range, and arity. Resolution converts those facts into source-site IDs before relationship dedupe, so the fresh graph has 84372 source-site occurrences and 0 missing source-site IDs. |
| relationship dedupe/proof metadata | `internal/resolution/emit.go`, `internal/graph/types.go`, `internal/lbugload/csv.go`, `internal/lbugload/queries.go` | `Relationship` now carries `sourceSiteId`, `sourceSiteIds`, `sourceSiteCount`, `sourceSiteStatus`, `proofKind`, `targetRole`, `targetText`, file path, and range. Dedupe preserves merged source-site IDs and occurrence counts. |
| graph accuracy command/report | `internal/graphaccuracy/source_site_accuracy.go`, `internal/cli/source_site_accuracy_command.go`, `internal/graphaccuracy/graphaccuracy.go` | `source-site-accuracy` reports source-site inventory, ACCESSES target labels, low-confidence fallback diagnostics, duplicate/merged relationship evidence, and false resolved edge candidates. Existing `graphaccuracy.go` still owns older Go graph accuracy checks. |
| unresolved type-reference emission | `internal/resolution/resolve.go`, `internal/resolution/emit.go` | `resolveTypeAnnotation` emits `type-reference` unresolved diagnostics with `sourceSiteID`, target text, source node, file hash, range, classification, and actionability when the type target is not resolved. |
| unresolved heritage emission | `internal/resolution/indexes.go`, `internal/resolution/resolve.go`, `internal/resolution/emit.go` | `resolveHeritage` indexes raw heritage facts; unresolved heritage calls `emitUnresolvedReference` with fact family `heritage`; resolved inheritance carries source-site metadata on `INHERITS`. |
| diagnostic attachment | `internal/resolution/emit.go`, `internal/graphhealth/diagnostics.go`, `internal/graphhealth/policy.go` | `emitUnresolvedReference` builds `graphhealth.Diagnostic` and attaches it with `AppendDiagnosticToNode`. Classification/actionability reuse the existing policy path. `sameDiagnosticBucket` uses `sourceSiteID` when present; fresh graph evidence has 0 missing diagnostic source-site IDs, so target identity is not lost in current data. |
| graph-health summary/report | `internal/graphhealth/compute.go`, `internal/graphhealth/policy.go`, `internal/httpapi/graph.go`, `internal/httpapi/server.go`, `internal/contracts/web_ui.go` | Topology status is computed by HTTP graph/report paths and typed in Web contracts; it is not persisted as a node property in `.avmatrix/graph.json`. Resolution diagnostics are persisted on source nodes under `graphHealthDiagnostics`. |
| HTTP graph payload | `internal/httpapi/graph.go`, `internal/semantic/metadata.go` | HTTP graph responses pass node `appLayer`/`appLayerSource` and `functionalArea`/`functionalAreaSource` properties through and include `semanticStatus`; NDJSON starts with `semantic_status`. Missing App Layer or Functional Area metadata is reported as stale/incomplete schema evidence and is not classified in the API loader. |
| generated Web contracts | `internal/contracts/web_ui.go`, `contracts/web-ui/avmatrix-web-contract.schema.json`, `avmatrix-web/src/generated/avmatrix-contracts.ts` | Contract manifest exposes App Layer and Functional Area enums/labels plus semantic status terms. Generated TypeScript exposes `FUNCTIONAL_AREAS`, `FUNCTIONAL_AREA_LABELS`, `FunctionalArea`, `GraphSemanticStatus.functionalArea`, `NodeProperties.functionalArea`, and `NodeProperties.functionalAreaSource` in addition to the App Layer fields. |
| query command | `internal/mcp/tools.go`, `internal/cli/tool_command.go`, `internal/httpapi/query.go` | Current query audit shows noisy process ranking; Phase 5/6 must improve the actual retrieval path and command output semantics. |
| context command | `internal/mcp/tools.go`, `internal/cli/tool_command.go`, `internal/httpapi/query.go`, `internal/lbugload/queries.go` | Context can read graph/DB data but does not yet present App Layer, Functional Area, and ResolutionGap as first-class semantic output. |
| impact command | `internal/mcp/tools.go`, `internal/cli/tool_command.go`, `internal/group/*`, `internal/lbugload/queries.go` | Impact surfaces graph blast radius; Phase 6 must add semantic layer/gap context without treating HIGH/CRITICAL as a blocker. |
| detect-changes command | `internal/mcp/tools.go`, `internal/cli/tool_command.go`, `internal/group/*`, `internal/repo/*` | Detect-changes remains required before commits; Phase 6 must include semantic layer/gap context in changed-scope output. |
| API-specific MCP tools | `internal/mcp/tools.go`, `internal/httpapi/*`, `internal/contracts/web_ui.go` | `route_map`, `shape_check`, and `api_impact` are API App Layer consumers and need semantic output coverage in later phases. |
| Web graph app state | `avmatrix-web/src/hooks/app-state/graph.tsx`, `avmatrix-web/src/services/backend-client.ts`, `avmatrix-web/src/generated/avmatrix-contracts.ts` | Web state receives graph nodes, relationships, and semantic status from API/stream contracts; it must not invent missing semantic truth client-side. |
| Web graph filters/detail/layout | `avmatrix-web/src/lib/graph-adapter.ts`, `avmatrix-web/src/lib/graph-health-filters.ts`, `avmatrix-web/src/components/FileTreePanel.tsx`, `avmatrix-web/src/components/GraphCanvas.tsx`, `avmatrix-web/src/hooks/useSigma.ts` | Current deterministic layout and filters do not yet expose App Layer rings/Resolution Health filters. Optimizer remains manual through `useSigma`; Phase 7 must not add auto optimizer execution. |

## E4 - Baseline Unresolved And App Layer Evidence

Status: complete for Phase 0 closure audit.

Recorded during P0-D and P0-E from fresh `.avmatrix/graph.json`, `.tmp\2026-05-22-phase0-baseline.json`, and `.tmp\2026-05-22-phase0-source-site-accuracy.json`.

Unresolved inventory:

| Metric | Value |
| --- | ---: |
| Unresolved diagnostic buckets | 57449 |
| Unresolved diagnostic occurrences | 58195 |
| Source-site diagnostic buckets | 57449 |
| Source-site diagnostic occurrences | 58195 |
| Low-confidence global fallback occurrences | 2159 |
| Missing diagnostic source-site IDs | 0 |

Unresolved fact-family counts:

| Fact family | Buckets | Occurrences |
| --- | ---: | ---: |
| call | 30946 | 30946 |
| access | 17935 | 17935 |
| type-reference | 8561 | 9307 |
| heritage | 7 | 7 |

Classification/actionability counts:

| Classification | Buckets | Occurrences | Actionability |
| --- | ---: | ---: | --- |
| in_repo_unresolved | 33428 | 33753 | analyzer_gap |
| builtin | 9775 | 9851 | non_actionable |
| standard_library | 7009 | 7339 | non_actionable |
| test_framework | 7079 | 7094 | non_actionable |
| external_library | 158 | 158 | review |

Top unresolved target texts:

| Rank | Fact family | Target text | Occurrences | Classification | Actionability | Source App Layer |
| ---: | --- | --- | ---: | --- | --- | --- |
| 1 | call | `t.Fatalf` | 3410 | test_framework | non_actionable | backend_test |
| 2 | type-reference | `testing.T` | 2452 | test_framework | non_actionable | backend_test |
| 3 | call | `len` | 1955 | builtin | non_actionable | cli_launcher |
| 4 | call | `string` | 1561 | builtin | non_actionable | cli_launcher |
| 5 | call | `append` | 1198 | builtin | non_actionable | cli_launcher |
| 6 | type-reference | `int` | 1075 | builtin | non_actionable | cli_launcher |
| 7 | call | `expect` | 870 | in_repo_unresolved | analyzer_gap | frontend_test |
| 8 | call | `make` | 652 | builtin | non_actionable | cli_launcher |
| 9 | call | `strings.Contains` | 587 | standard_library | non_actionable | cli_launcher |
| 10 | access | `result.Metrics` | 575 | in_repo_unresolved | analyzer_gap | backend |
| 11 | type-reference | `collector` | 481 | in_repo_unresolved | analyzer_gap | backend |
| 12 | call | `any` | 474 | builtin | non_actionable | backend_test |
| 13 | call | `filepath.Join` | 432 | standard_library | non_actionable | cli_launcher |
| 14 | call | `c.text` | 379 | in_repo_unresolved | analyzer_gap | backend |
| 15 | call | `t.Helper` | 345 | test_framework | non_actionable | backend_test |

Source buckets:

| Source dimension | Top values |
| --- | --- |
| Source labels | Function 44043, Method 7589, Variable 4128, Struct 926, File 418, Property 162, Package 59, TypeAlias 44, Const 31, Interface 24 |
| Source App Layers | backend_test 19155, backend 17229, frontend_test 5480, api_test 5041, frontend 4813, api 4116, api_contract 609, cli_launcher 525, frontend_api_client 371, config 88, generated_contract 22 |
| Source Functional Areas | providers 12593, unknown 7618, analyzer 7000, mcp 4640, api 4477, resolution 3777, cli 2923, web_graph_ui 2598, storage 2561, query 2475, graph_health 1835, embeddings 1646 |
| Source path buckets | internal/providers 11299, avmatrix-web/src 5227, internal/mcp 4775, internal/httpapi 4391, avmatrix-web/test 3914, internal/resolution 3800, internal/cli 2996, internal/group 2175, internal/analyze 2009, internal/embeddings 1701, avmatrix-web/e2e 1596 |

Topology-status note:

- The graph snapshot persists resolution diagnostics on source nodes, but it does not persist a topology-status node property. Topology Health is computed by graph-health/API report paths. Phase 3/4 must keep Resolution Health separate from this computed topology layer and must not use unresolved references alone as dead-code proof.

App Layer inventory from the fresh graph:

| App Layer | Node count |
| --- | ---: |
| backend | 9937 |
| backend_test | 4457 |
| api | 2010 |
| frontend | 1859 |
| docs | 1604 |
| api_test | 1107 |
| frontend_test | 620 |
| mixed | 364 |
| cli_launcher | 256 |
| frontend_api_client | 181 |
| api_contract | 161 |
| config | 37 |
| unknown | 26 |
| generated_contract | 16 |
| shared_contract | 0 |
| api_shared_contract | 0 |
| generated | 0 |

Seed-path rule notes:

| Seed path or pattern | Current classification behavior |
| --- | --- |
| `avmatrix-web/src/**` | frontend, except `avmatrix-web/src/services/backend-client.ts` as `frontend_api_client`. |
| `avmatrix-web/test/**`, `avmatrix-web/e2e/**` | frontend_test. |
| `internal/**` | backend by default, with stronger rules for API, contracts, providers, graph health, resolution, storage, and other Functional Areas. |
| `cmd/**` | backend/cli command surfaces; API contract generator paths classify into API contract/test categories when applicable. |
| `internal/httpapi/**`, `internal/mcp/**` | api or api_test depending on test/source path. |
| `contracts/**`, `internal/contracts/**`, `cmd/generate-web-contracts/**` | api_contract, generated_contract, or api_test depending on source/test/generated path. |
| `docs/**`, `reports/**`, `*.md` | docs/reporting categories. |
| `*_test.go`, Web `*.test.*` and `*.spec.*` | backend_test, api_test, or frontend_test based on owning path. |
| Config/package/build files | config when covered by explicit known config/build rules; unknown when evidence is insufficient. |

Unknown App Layer examples:

- `Dockerfile.cli`, `Dockerfile.web`, `avmatrix-web/vercel.json`, `baseline/phase-1-contract-freeze/*.json`, and several root/folder nodes remain `unknown` because the current rules do not have enough high-confidence ownership evidence.

## E5 - Baseline Query Evidence

Status: complete for Phase 0 closure audit.

Recorded during P0-F after fresh analyze. Raw query output files are `.tmp\phase0-query-*.json`.

| Intent | Expected files/symbols | Actual top results | Hit/miss | Noise reason |
| --- | --- | --- | --- | --- |
| unresolved reference diagnostic generation | `internal/resolution/resolve.go`, `internal/resolution/emit.go`, `internal/graphhealth/diagnostics.go`, `internal/graphhealth/policy.go` | `avmatrix-launcher/src/main.go` `main/resetRuntime/stopRuntime/...`; `internal/graphaccuracy/property_access.go`; `internal/contracts/web_ui.go`; `internal/cli/admin_command.go` | miss | Process ranking returns generic `Main -> HiddenProcAttr`, property-access audit, contract title-word, and clean-command flows instead of resolver/diagnostic code. |
| graph health unknown-connectivity separation | `internal/graphhealth/compute.go`, `internal/graphhealth/policy.go`, `internal/httpapi/graph.go`, `internal/contracts/web_ui.go` | Same generic launcher/property-access/contracts/admin top processes | miss | Query terms `graph health` and `unknown connectivity` do not pull the graph-health implementation; process contains scoring dominates. |
| App Layer and ResolutionGap layout | `internal/semantic/app_layer.go`, `internal/semantic/functional_area.go`, `avmatrix-web/src/lib/graph-adapter.ts`, `avmatrix-web/src/lib/graph-health-filters.ts`, `avmatrix-web/src/hooks/useSigma.ts` | Same generic launcher/property-access/contracts/admin top processes | miss | Layout/filter terms are not matching Web graph surfaces; current retrieval is not layer-aware. |
| API contract surfaces | `internal/contracts/web_ui.go`, `contracts/web-ui/avmatrix-web-contract.schema.json`, `avmatrix-web/src/generated/avmatrix-contracts.ts`, `avmatrix-web/src/services/backend-client.ts` | Launcher flow first, then property-access audit and contract title-word process | partial | Contract file appears through `WebUIContract*`, but top result is unrelated launcher noise and frontend client/generated schema are not surfaced. |
| frontend graph filter surfaces | `avmatrix-web/src/lib/graph-health-filters.ts`, `avmatrix-web/src/hooks/app-state/graph.tsx`, `avmatrix-web/src/components/FileTreePanel.tsx`, `avmatrix-web/src/components/GraphCanvas.tsx` | Same generic launcher/property-access/contracts/admin top processes | miss | Frontend graph/filter files are absent from top results; query is dominated by unrelated process labels. |
| runtime reset hidden-terminal behavior | `avmatrix-launcher/src/main.go` `resetRuntime`, `stopRuntime`, `hiddenCommand`, `hiddenProcAttr` | `avmatrix-launcher/src/main.go` `main/resetRuntime/stopRuntime/stopPID/waitForPIDExit/processAlive/tasklistCommand/hiddenCommand/hiddenProcAttr` | hit | This intent matches because expected terms overlap the dominant launcher process. |

## E6 - App Layer Implementation Evidence

Status: in progress; Phase 1 App Layer taxonomy, persistence, public status, and naming registry complete.

Record during P1.

Required evidence:

- final App Layer category registry;
- source evidence rules for every category;
- examples for backend, api, frontend, shared/API contract, frontend API client, test/doc/config/generated, mixed, and unknown;
- stale/incomplete schema behavior proving missing App Layer metadata is not load-time classified;
- enum keys, display labels, CLI labels, and Web labels for App Layer and initial ResolutionGap naming;
- schema/snapshot/API/contract fields changed;
- tests proving one primary category per node and no overlapping primary labels;
- before/after counts and generated contract output.

Implemented evidence:

- Category registry is defined in `internal/semantic/app_layer.go` as one primary `AppLayer` value per node: `backend`, `api`, `frontend`, `cli_launcher`, `shared_contract`, `api_contract`, `api_shared_contract`, `frontend_api_client`, `backend_test`, `frontend_test`, `api_test`, `generated_contract`, `docs`, `config`, `generated`, `mixed`, and `unknown`.
- Source rules implemented in `ClassifyAppLayer`: docs/report/markdown paths; Web test roots and `*.test.*`/`*.spec.*`; Go API tests under `internal/httpapi`, `internal/mcp`, `internal/contracts`, and `cmd/generate-web-contracts`; generated contract paths; API contract paths; `avmatrix-web/src/services/backend-client.ts`; config files; generated paths; API paths including `internal/httpapi`, `internal/mcp`, `app/api`, and `pages/api`; frontend roots; CLI launcher paths; backend paths; otherwise `unknown`.
- Mixed category inference is relationship-backed for Process and Community nodes: if membership/step relationships connect more than one non-unknown App Layer, the target node receives `mixed` rather than overlapping labels.
- Generated contract output was regenerated with `go run ./cmd/generate-web-contracts`.
- `internal/semantic/metadata.go` defines `semantic_app_layer_v1`, `GraphSemanticStatus`, `StatusComplete`, and `StatusStaleIncomplete`. `GraphSemanticStatus` only counts existing node metadata; it does not infer App Layer at API or UI load time.
- `internal/httpapi/graph.go` returns `semanticStatus` in JSON graph payloads and emits a first NDJSON record with type `semantic_status`. A graph with missing `appLayer` or `appLayerSource` is marked `stale_incomplete`; a graph with explicit `unknown` App Layer remains fresh semantic evidence.
- `avmatrix-web/src/services/backend-client.ts` accepts the streamed `semantic_status` record and returns it as `semanticStatus` with the graph payload.
- User-facing and machine-facing names are centralized in `internal/semantic/metadata.go`: App Layer category labels plus semantic terms for `app_layer`, `api_layer`, `api_contract`, `frontend_api_client`, `resolution_gap`, `unresolved_symbol`, `analyzer_gap`, `external_reference`, and `non_actionable_reference`. The Go contract exports those labels into the Web manifest and generated TypeScript.
- Fresh analyze with the locally built CLI produced these App Layer counts: backend 9554, api 2013, frontend 1862, cli_launcher 256, api_contract 155, frontend_api_client 182, backend_test 4415, frontend_test 620, api_test 1102, generated_contract 17, docs 1601, config 37, mixed 399, unknown 26, shared_contract 0, api_shared_contract 0, generated 0.
- Validation evidence: full build passed with `powershell -ExecutionPolicy Bypass -File .\avmatrix-launcher\build.ps1`; focused Go validation passed with `go test .\internal\semantic .\internal\httpapi .\internal\contracts`; wider Go validation passed with `go test .\internal\... .\cmd\...`; Web unit validation passed with `npm test -- --run` in `avmatrix-web`; whitespace validation passed with `git diff --check`.
- Test evidence includes `TestGraphSemanticStatusDistinguishesUnknownFromMissingMetadata`, `TestSemanticTermDefinitionsAreStableAndNonOverlapping`, `TestGraphReturnsJSONForRegisteredRepo`, `TestGraphPayloadMarksFreshSemanticMetadataComplete`, `TestGraphStreamingReturnsNDJSON`, `TestGraphStreamingKeepsRouteAndToolMetadata`, and `TestWebUIContractManifestUsesGoRuntimeConstants`.
- `go test ./...` still fails outside the implementation slice because fixture packages under `avmatrix/test/fixtures/...` are intentionally non-buildable as standalone Go packages; the prior real failure in `internal/analyze` is fixed.

## E7 - Functional Area Evidence

Status: complete for Phase 2.

Implemented evidence:

- Functional Area registry is defined in `internal/semantic/functional_area.go`: `resolution`, `graph_health`, `query`, `mcp`, `web_graph_ui`, `layout`, `contracts`, `providers`, `runtime`, `analyzer`, `session`, `launcher`, `cli`, `reporting`, `api`, `storage`, `embeddings`, `configuration`, `documentation`, `mixed`, and `unknown`.
- Accepted signals for direct node classification are deterministic high-confidence path/package rules. Examples from fresh graph output include `internal/resolution/access_audit.go` as `resolution`, `internal/graphhealth/**` and `internal/graphaccuracy/**` as `graph_health`, `internal/httpapi/**` as `api`, `avmatrix-web/src/lib/graph-adapter.ts` as `layout`, `internal/lbugload/**` and `internal/graph/**` as `storage`, and markdown/docs paths as `documentation`.
- Process and Community nodes use relationship-backed inference only after processes/communities exist: one non-unknown area becomes that area, multiple non-unknown areas become `mixed`, and no accepted evidence remains `unknown`.
- Rejected low-confidence signals for Phase 2 are import/call neighborhood ownership, community label text, process label text, AI-assisted labeling, and explicit semantic config that does not yet exist. These are not used to reduce unknown counts.
- Functional Area is persisted on graph node properties as `functionalArea` and `functionalAreaSource`; fresh analyze output has zero missing Functional Area fields. Unknown remains explicit evidence, not a stale-field fallback.
- `semantic_app_functional_v1` in `internal/semantic/metadata.go` marks the schema version where both App Layer and Functional Area are required semantic fields. `GraphSemanticStatus` now reports both `appLayer` and `functionalArea`; stale fixture tests prove missing fields are reported as stale/incomplete rather than guessed at load time.
- LadybugDB schema/export surfaces include `functionalArea` on node tables and CSV rows. This makes Functional Area available to DB-backed CLI/Cypher consumers immediately. Rich `query`, `context`, `impact`, and `detect-changes` semantic formatting remains Phase 6.
- Generated Web contracts expose Functional Area enum values, labels, semantic status, and node property fields. Visible Web filters/detail rendering remain Phase 7; this slice only ensures the graph/API/contract data exists for those UI surfaces.
- Fresh analyze benchmark using `.\avmatrix-launcher\server-bundle\avmatrix.exe analyze --force --benchmark-json .tmp\2026-05-22-functional-area-semantic-enrichment.json --benchmark-label functional-area-semantic-enrichment` produced 22358 nodes, 55349 relationships, 0 missing `functionalArea`, 0 missing `functionalAreaSource`, 20622 classified Functional Area nodes, and 1736 `unknown` nodes.
- Functional Area counts from `.avmatrix/graph.json`: analyzer 2426, api 1583, cli 943, configuration 37, contracts 248, documentation 1293, embeddings 701, graph_health 761, launcher 243, layout 280, mcp 1566, mixed 294, providers 4954, query 1154, reporting 308, resolution 1225, runtime 21, session 510, storage 1243, unknown 1736, web_graph_ui 832.
- Test evidence includes `TestClassifyFunctionalAreaUsesHighConfidencePathRules`, `TestApplyAnnotatesAppLayerAndFunctionalArea`, `TestApplyInfersSemanticMetadataForProcessAndCommunityNodes`, `TestGraphSemanticStatusDistinguishesUnknownFromMissingMetadata`, `TestNodeSchemaIncludesSemanticFields`, `TestNodeCSVRowIncludesSemanticFields`, `TestLoadGraphWritesSemanticFields`, `TestWebUIContractManifestUsesGoRuntimeConstants`, and `TestGraphPayloadMarksFreshSemanticMetadataComplete`.
- Validation evidence for this slice: full build passed with `powershell -ExecutionPolicy Bypass -File .\avmatrix-launcher\build.ps1`; focused Go tests passed with `go test .\internal\semantic .\internal\lbugschema .\internal\lbugload .\internal\contracts .\internal\httpapi .\internal\analyze`; wider Go tests passed with `go test .\internal\... .\cmd\...`; Web unit tests passed with `npm test -- --run` in `avmatrix-web`. Web e2e was not run for Phase 2 because this slice changed persisted graph/contract metadata but no visible Web UI behavior.
- AVmatrix change detection before commit used `.\avmatrix-launcher\server-bundle\avmatrix.exe detect-changes --repo AVmatrix --scope all`; it returned affected_count 24, changed_count 144, changed_files 18, risk_level `critical`. The critical scope is expected for this slice because it intentionally changes semantic enrichment, graph semantic status, LadybugDB schema/export, HTTP/Web contract surfaces, generated contracts, and plan/evidence ledgers.

## E7A - Proof-Based CALLS/ACCESSES Evidence

Status: complete for Phase 2A; Phase 3 can now consume source-backed unresolved inputs

Record during Phase 2A.

Required evidence:

- final resolved-edge contract for `CALLS` and `ACCESSES`;
- accepted proof kinds for calls and property/field accesses;
- rejected proof sources such as global simple-name fallback, cross-language same-name fallback, coarse file-level source edges, and selector/import references that resolve to functions;
- code trace proving low-confidence `resolveGlobalCallName` results are no longer emitted as resolved `CALLS` without accepted proof;
- source-site inventory schema with sourceSiteID, source node, file/range, target text, fact family, status, proof kind, target role, and linked resolved target when available;
- chosen persistence schema for source-site facts: extended `graph.Relationship` metadata, SourceSite entity/record, or both, plus the exact consumer contract for each;
- evidence that `semanticEdgeKey` or any replacement dedupe preserves exact source-site occurrences or exact occurrence counts;
- evidence that `propertyLabels()` and provider fact coverage now match the strict ACCESSES contract or split non-property uses into a separate relation/fact role;
- resolver trace for the Go `stop()` false-positive class, selector/import cases, receiver method calls, property reads/writes, local function variables, closures, imports, builtins, external packages, and TypeScript/React owner attribution;
- provider and scope-IR evidence for `CallSiteFact` and `AccessFact` preservation across supported languages;
- golden corpus cases proving expected resolved edges exist, known false edges do not exist, and unresolved source sites are persisted;
- graphaccuracy or dedicated CLI/report command output for B5A metrics;
- graph/API/DB/contract/CLI visibility for source-site status and proof metadata, or exact limitations when a surface is deferred;
- tests and benchmark output proving false resolved edges `0`, silent missing source sites `0`, source sites hidden by dedupe `0`, and non-property ACCESSES targets `0` in the golden corpus unless non-property uses have been split into a separate relation/fact role.

Implemented evidence for the low-confidence global CALLS fallback slice:

- `internal/resolution/indexes.go` now keeps `resolveName` behavior for existing non-call consumers but exposes `resolveScopedName` so call resolution can require local/import scope proof before trying narrower fallback paths.
- `internal/resolution/resolve.go` now uses `resolveScopedName` for constructor/free-call proof, keeps same-file and Go same-package fallback as explicit lower-confidence resolved edges, and records `resolveGlobalCallName` matches as unresolved source-backed diagnostics with note `call target matched low-confidence global fallback only`.
- `internal/resolution/resolution_test.go` updates the global fallback arity case to require no resolved edge and adds `TestResolveBareGoCallDoesNotFallbackToCrossLanguageMethod`, proving a bare Go `stop()` call does not emit `CALLS` to TypeScript `SSEListener.stop`.
- `internal/providers/php/imports.go` now emits PHP `use function` and `use const` imports, and `internal/providers/php/extract_test.go` proves those imports exist. This preserves proof-backed PHP imported function calls after global fallback is rejected.
- `internal/resolution/testdata/typescript_graph_signature.golden.json` records the expected confidence change for same-file function fallback from `1.000` to `0.950`.
- Impact context was checked with AVmatrix after `avmatrix analyze --force`: `resolveCall`, `resolveAccess`, `propertyLabels`, `semanticEdgeKey`, `workspace.resolveName`, and PHP import surfaces were inspected. CRITICAL/HIGH results were treated as blast-radius context requiring focused and wider validation, not as a prohibition on editing the required code.
- Full build passed with `powershell -ExecutionPolicy Bypass -File .\avmatrix-launcher\build.ps1`.
- Focused tests passed: `go test .\internal\resolution`, `go test .\internal\providers\php`, and `go test .\internal\graphaccuracy .\internal\analyze .\internal\cli`.
- Wider backend validation passed with `go test .\internal\... .\cmd\...`.
- AVmatrix change detection before commit used fresh `.\avmatrix-launcher\server-bundle\avmatrix.exe analyze --force` followed by `.\avmatrix-launcher\server-bundle\avmatrix.exe detect-changes --repo AVmatrix --scope all`; it reported changed_count `25`, changed_files `9`, affected_count `0`, and risk_level `low`.
- No Web UI behavior changed in this slice, so Web e2e was not required for this slice.

Implemented evidence for the source-site/proof metadata persistence slice:

- `internal/graph/types.go` extends `graph.Relationship` with source-site identity, status, proof kind, target role/text, file path, range, and occurrence fields: `sourceSiteId`, `sourceSiteIds`, `sourceSiteCount`, `sourceSiteStatus`, `proofKind`, `targetRole`, `targetText`, `filePath`, `startLine`, `startCol`, `endLine`, and `endCol`.
- `internal/graphhealth/policy.go` and `internal/graphhealth/diagnostics.go` extend unresolved-reference diagnostics with `sourceSiteId`, `sourceSiteStatus`, `proofKind`, `targetRole`, and full range fields. `sameDiagnosticBucket` now preserves target identity by source site when available instead of collapsing different unresolved targets from the same source/fact/file/note bucket.
- `internal/resolution/source_site.go` defines the source-site status/proof/target-role constants used by resolver output. `internal/resolution/resolve.go` records accepted proof kinds for resolved calls and accesses: `scope-binding`, `same-file`, `go-same-package`, `receiver-member`, and `import-member`; low-confidence global fallback remains an unresolved diagnostic with proof marker `global-fallback-low-confidence`.
- `internal/resolution/emit.go` persists source-site metadata onto resolved relationships and unresolved diagnostics. `mergeRelationship` now carries `sourceSiteIds` and `sourceSiteCount` so duplicate edge dedupe does not hide exact source-site occurrences.
- `internal/resolution/indexes.go` changes `propertyLabels()` to `Property` only. `internal/resolution/access_audit.go` adds the `non_property_target` bucket for selector targets that resolve to `Variable`, `Const`, or `Static` but are not valid proof for `ACCESSES`.
- `internal/lbugschema/schema.go`, `internal/lbugload/csv.go`, and `internal/lbugload/queries.go` propagate the new relationship fields through LadybugDB relation schema, CSV export, copy load, and fallback insert. The phase-1 LadybugDB graph contract freeze file was updated for the intentionally changed relation schema.
- `internal/contracts/web_ui.go` and generated `avmatrix-web/src/generated/avmatrix-contracts.ts` expose the new relationship and diagnostic fields to Web consumers without making the Web UI infer source-site truth.
- Tests updated or added for source-site metadata preservation, unresolved diagnostic metadata, duplicate source-site occurrence preservation, strict non-property ACCESSES rejection, diagnostic bucket identity, LadybugDB CSV/schema propagation, and generated contract exposure.
- Validation evidence for this slice: full build passed twice with `powershell -ExecutionPolicy Bypass -File .\avmatrix-launcher\build.ps1`; focused tests passed with `go test .\internal\resolution`, `go test .\internal\graphhealth`, `go test .\internal\lbugload .\internal\lbugschema`, and `go test .\internal\contracts`; wider backend validation passed with `go test .\internal\... .\cmd\...`.
- Fresh `.\avmatrix-launcher\server-bundle\avmatrix.exe analyze --force` passed after schema propagation and wrote `.avmatrix\graph.json` with scanned `733`, parsed `544`, unsupported `189`, failed `0`, graph nodes `22404`, and graph relationships `51521`.
- AVmatrix change detection before commit used fresh `.\avmatrix-launcher\server-bundle\avmatrix.exe analyze --force` followed by `.\avmatrix-launcher\server-bundle\avmatrix.exe detect-changes --repo AVmatrix --scope all`; it reported changed_count `100`, changed_files `21`, affected_count `33`, and risk_level `critical`. The critical scope is expected because this slice intentionally changes graph relationship schema, graphhealth diagnostic schema, resolver edge emission metadata, strict ACCESSES target policy, LadybugDB relation schema/export/load, generated Web contracts, and plan/evidence/benchmark ledgers.
- No visible Web UI behavior changed in this slice, so Web e2e was not required for this slice.

Implemented evidence for the source-site accuracy command graph-inventory slice:

- `internal/graphaccuracy/graphaccuracy.go` now decodes source-site/proof fields from graph relationship JSON so graphaccuracy consumers can inspect the relationship metadata persisted by the resolver.
- `internal/graphaccuracy/source_site_accuracy.go` adds `RunSourceSiteAccuracy`, `WriteSourceSiteAccuracyResult`, and `SourceSiteAccuracySummaryLines`. The report reads current graph JSON and records source-site relationship/diagnostic buckets and occurrences, status/proof/fact-family/target-role counts, resolved `CALLS`/`ACCESSES` counts, ACCESSES target label distribution, duplicate source-target pairs, merged source-site occurrence evidence, and graph-policy violation candidates.
- `internal/cli/source_site_accuracy_command.go` adds packaged command `avmatrix source-site-accuracy` with `--graph`, `--out`, `--json`, and `--max-examples`. This puts the report into the built `avmatrix.exe` path instead of leaving it as a standalone development command.
- The command deliberately reports golden validation as disabled in graph-inventory mode. P2A-F/P2A-G remain open for fixture-backed false-positive and silent-missing source-site checks; the command still exposes the fields so the later fixture mode can report them without changing the output shape.
- Fresh local build/analyze/report command: `.\avmatrix-launcher\server-bundle\avmatrix.exe analyze --force` produced scanned `736`, parsed `547`, unsupported `189`, failed `0`, graph nodes `22628`, graph relationships `52161`; `.\avmatrix-launcher\server-bundle\avmatrix.exe source-site-accuracy --graph .avmatrix\graph.json --out .tmp\2026-05-22-source-site-accuracy.json --max-examples 20` produced `84355` source-site occurrences with `84355` stable IDs and `0` missing IDs.
- Source-site accuracy command output on the fresh graph: relationship source-site occurrences `26221` across `15494` buckets, diagnostic source-site occurrences `58134` across `57388` buckets, resolved `CALLS=7648`, resolved `ACCESSES=3298`, low-confidence fallback diagnostics `2159`, ACCESSES Property targets `3298/3298`, non-property ACCESSES targets `0`, duplicate source-target pairs `10`, max duplicate `2`, merged relationships `5174`, merged source-site occurrences `15901`, resolved edges without proof `0`, resolved edges without source-site ID `0`, and low-confidence fallback resolved edges `0`.
- The command found `16` graph-policy false-resolved-edge candidates, all from coarse `File -> Function` `CALLS` edges in Web unit test files such as `File:avmatrix-web/test/unit/constants.test.ts -> Function:avmatrix-web/src/lib/constants.ts:getNodeColor`. This keeps P2A-D and P2A-F open: coarse file-level call sources still need to be kept out of resolved symbol-level `CALLS` or split into a separate relation/fact role.
- Tests added: `TestRunSourceSiteAccuracyReportsProofInventory` and `TestSourceSiteAccuracyCommandOutputsJSON`.
- Validation evidence for this slice: full build passed with `powershell -ExecutionPolicy Bypass -File .\avmatrix-launcher\build.ps1`; focused tests passed with `go test .\internal\graphaccuracy` and `go test .\internal\cli`; wider backend validation passed with `go test .\internal\... .\cmd\...`.
- AVmatrix change detection before commit used fresh `.\avmatrix-launcher\server-bundle\avmatrix.exe analyze --force` followed by `.\avmatrix-launcher\server-bundle\avmatrix.exe detect-changes --repo AVmatrix --scope all`; it reported changed_count `35`, changed_files `6`, affected_count `11`, and risk_level `high`. The high scope is expected because this slice adds a packaged root CLI command and extends graphaccuracy graph relationship decoding/reporting.
- No visible Web UI behavior changed in this slice, so Web e2e was not required.

Implemented evidence for the File-source CALLS gate slice:

- `internal/resolution/resolve.go` now checks resolved call sources before target resolution. If the only caller owner is a `File` node, the resolver does not emit a resolved `CALLS` edge.
- The file-level call site is preserved as a source-backed unresolved/reference diagnostic with `sourceSiteStatus=unsupported_syntax`, target role `callable`, and note `call source is file-level; resolved edge not emitted`. This keeps the source-site inventory while avoiding a fake symbol-level topology edge.
- `internal/resolution/emit.go` maps that note to `unsupported_syntax`, and `internal/resolution/source_site.go` centralizes the note string.
- `TestResolveDoesNotEmitResolvedCallFromFileCaller` proves a module-scope call to a same-file function no longer emits `File:src/app.ts -> Function:src/app.ts:target` and still preserves the source-site diagnostic on the file node.
- Fresh local build/analyze/report command after the gate: `.\avmatrix-launcher\server-bundle\avmatrix.exe analyze --force` produced scanned `736`, parsed `547`, unsupported `189`, failed `0`, graph nodes `22635`, graph relationships `52144`; `.\avmatrix-launcher\server-bundle\avmatrix.exe source-site-accuracy --graph .avmatrix\graph.json --out .tmp\2026-05-22-source-site-accuracy-after-file-source-call-gate.json --max-examples 20` produced `policy.falseResolvedEdgeCandidates=0`, `coarseFileCallEdges=0`, resolved edges without proof `0`, resolved edges without sourceSiteID `0`, low-confidence fallback resolved edges `0`, and non-property ACCESSES targets `0`.
- Validation evidence for this slice: full build passed with `powershell -ExecutionPolicy Bypass -File .\avmatrix-launcher\build.ps1`; focused tests passed with `go test .\internal\resolution` and `go test .\internal\graphaccuracy`; wider backend validation passed with `go test .\internal\... .\cmd\...`.
- AVmatrix change detection before commit used fresh `.\avmatrix-launcher\server-bundle\avmatrix.exe analyze --force` followed by `.\avmatrix-launcher\server-bundle\avmatrix.exe detect-changes --repo AVmatrix --scope all`; it reported changed_count `15`, changed_files `7`, affected_count `0`, and risk_level `low`.

Implemented evidence for the golden corpus slice:

- `internal/resolution/proof_accuracy_golden_test.go` adds `TestProofBasedCallAccessGoldenCorpus`, a controlled ScopeIR fixture with twelve call/access source sites. It avoids parser ambiguity and asserts the resolver contract directly against graph relationships and source-backed diagnostics.
- Positive edge expectations: two `helper()` calls merge into one proof-backed `CALLS` relationship with `sourceSiteCount=2`; `user.save()` emits a `CALLS` edge with `proofKind=receiver-member`; `closure()` emits a `CALLS` edge with `proofKind=scope-binding`; `api.fetchUser()` emits a `CALLS` edge with `proofKind=import-member`; `user.id` emits an `ACCESSES` edge with `proofKind=receiver-member`.
- Negative edge expectations: `callback()` is a local function-variable call and must not emit `CALLS` to a `Variable`; bare `stop()` must not emit `CALLS` to TypeScript `SSEListener.stop`; a module/file-level `helper()` call must not emit `File -> Function` `CALLS`; `config.make` must not emit `ACCESSES` to a `Function`.
- Diagnostic expectations: unresolved `callback`, low-confidence `stop`, builtin `len`, external `cobra.Command`, non-property selector `config.make`, and file-level `helper` all keep source-site IDs, fact family, target text, status, proof kind, classification, actionability, and occurrence count.
- Golden metrics from the test: false resolved edges `0`; silent missing call/access source sites `0`; call/access source-site occurrences `12`; resolved call/access occurrences `6`; unresolved call/access diagnostics `6`; non-property resolved `ACCESSES` targets `0`; source sites hidden by helper-call dedupe `0` because both helper source-site IDs survive on the merged relationship.
- Validation evidence for this slice: full build passed with `powershell -ExecutionPolicy Bypass -File .\avmatrix-launcher\build.ps1`; focused resolver validation passed with `go test .\internal\resolution`; focused graph/CLI/contract surfaces passed with `go test .\internal\graphaccuracy .\internal\cli .\internal\contracts .\internal\lbugload .\internal\graphhealth`; wider backend validation passed with `go test .\internal\... .\cmd\...`.
- Fresh `.\avmatrix-launcher\server-bundle\avmatrix.exe analyze --force --benchmark-json .tmp\2026-05-22-p2a-golden-postedit-analyze.json --benchmark-label p2a-golden-postedit` passed after the golden corpus slice and produced scanned `737`, parsed `548`, unsupported `189`, failed `0`, graph nodes `22710`, and graph relationships `52293`.
- AVmatrix change detection before commit used fresh analyze output followed by `.\avmatrix-launcher\server-bundle\avmatrix.exe detect-changes --repo AVmatrix --scope all`; after staging the new test file and final ledger updates it reported changed_count `81`, changed_files `4`, affected_count `0`, and risk_level `low`.

Implemented evidence for the source-site accuracy golden fixture command slice:

- `internal/graphaccuracy/source_site_accuracy.go` adds `GoldenPath` input, `ReadSourceSiteGoldenFixture`, fixture structs, and fixture-backed validation that compares expected source-site IDs plus known-false resolved edges against the current graph snapshot.
- `SourceSiteGoldenValidation` now reports `expectedSourceSites`, `matchedSourceSites`, `silentMissingSourceSites`, `expectedFalseResolvedEdges`, `falseResolvedEdges`, missing source-site IDs, and capped examples for missing sites and false edges. This keeps graph inventory and fixture-backed policy checks in the same report shape.
- `internal/cli/source_site_accuracy_command.go` adds packaged CLI flag `--golden`. Command help from the built executable shows `--golden string      source-site golden fixture JSON for false-edge and missing-site validation`.
- `TestRunSourceSiteAccuracyValidatesGoldenFixture` validates fixture mode with 7 expected source-site IDs, 6 matched IDs, 1 silent missing source site, 1 expected false resolved edge, and 1 false resolved edge found in the fixture graph.
- `TestSourceSiteAccuracyCommandOutputsJSON` validates CLI JSON visibility for `--golden`, including `"enabled": true`, `"silentMissingSourceSites": 1`, and `"falseResolvedEdges": 1`.
- Fresh command output from the built executable on the current graph without a fixture was written to `.tmp\2026-05-22-p2a-source-site-accuracy-command.json`: relationship source-site occurrences `26304`, diagnostic source-site occurrences `58455`, all source-site occurrences `84759`, stable source-site ID occurrences `84759`, missing source-site ID occurrences `0`, resolved `CALLS=7645`, resolved `ACCESSES=3314`, low-confidence fallback diagnostics `2217`, ACCESSES Property targets `3314/3314`, non-property ACCESSES targets `0`, false resolved edge candidates `0`, coarse file call edges `0`, and golden validation disabled because no fixture was supplied.
- Validation evidence for this slice: full build passed with `powershell -ExecutionPolicy Bypass -File .\avmatrix-launcher\build.ps1`; focused tests passed with `go test .\internal\graphaccuracy` and `go test .\internal\cli`; wider backend validation passed with `go test .\internal\...` and `go test .\cmd\...`.
- Fresh `.\avmatrix-launcher\server-bundle\avmatrix.exe analyze --force --benchmark-json .tmp\2026-05-22-p2a-source-site-accuracy-golden-command-postedit-analyze.json --benchmark-label p2a-source-site-accuracy-golden-command-postedit` passed before detect-changes and produced scanned `737`, parsed `548`, unsupported `189`, failed `0`, graph nodes `22751`, and graph relationships `52425`.
- AVmatrix change detection before commit used fresh analyze output followed by `.\avmatrix-launcher\server-bundle\avmatrix.exe detect-changes --repo AVmatrix --scope all`; it reported changed_count `74`, changed_files `7`, affected_count `12`, and risk_level `high`. The high scope is expected because this slice extends `RunSourceSiteAccuracy`, `SourceSiteAccuracySummaryLines`, `newSourceSiteAccuracyCommand`, and new graphaccuracy fixture-validation helpers in a packaged CLI/report flow.
- No visible Web UI behavior changed in this slice, so browser/e2e validation was not required for this slice.

Implemented evidence for the source-site to ResolutionGap input bridge slice:

- `internal/graphhealth/resolution_gap_inputs.go` adds `ResolutionGapInput`, the source-backed input record intended for Phase 3 ResolutionGap/UnresolvedSymbol persistence. It is an input model only; it does not synthesize fake in-repo target nodes, fake resolved semantic edges, or fake topology edges.
- `SourceBackedResolutionGapInputs` reads persisted `graphHealthDiagnostics` entries whose kind is `unresolved_reference` and whose `sourceSiteId` is present. It preserves `sourceSiteId`, source node ID/label, source App Layer, source Functional Area, fact family, target text, target role, source-site status, proof kind, classification, actionability, resolution source, source, file path, file hash, range, count, and note.
- `SourceBackedCallAccessResolutionGapInputs` filters those source-backed inputs to call/access fact families so Phase 3 can consume call/access resolution health from source-site records rather than graph-health summary counts.
- The existing diagnostic summaries remain unchanged. The new bridge reads the persisted per-source-site diagnostic records before summary/report aggregation and does not rely on capped graph-health report candidates.
- `TestSourceBackedResolutionGapInputsPreserveSourceSiteEvidence` proves call/access/type-reference source-backed inputs preserve App Layer, Functional Area, source-site ID/status, proof kind, target role, classification, actionability, file/range, count, and note; it also proves diagnostics without `sourceSiteId` are not used as precise source-backed inputs.
- Fresh source-site accuracy command after this slice wrote `.tmp\2026-05-22-p2a-gap-input-source-site-accuracy.json`: relationship source-site occurrences `26395`, diagnostic source-site occurrences `58560`, all source-site occurrences `84955`, stable source-site ID occurrences `84955`, missing source-site ID occurrences `0`, unresolved diagnostic fact-family counts `call=31091`, `access=18129`, `type-reference=9333`, `heritage=7`, resolved `CALLS=7662`, resolved `ACCESSES=3344`, false resolved edge candidates `0`, and non-property ACCESSES targets `0`.
- Validation evidence for this slice: full build passed twice with `powershell -ExecutionPolicy Bypass -File .\avmatrix-launcher\build.ps1` because the focused test was corrected after an ordering assertion failure; final focused tests passed with `go test .\internal\graphhealth` and `go test .\internal\resolution .\internal\graphaccuracy`; wider backend validation passed with `go test .\internal\...` and `go test .\cmd\...`.
- Fresh `.\avmatrix-launcher\server-bundle\avmatrix.exe analyze --force --benchmark-json .tmp\2026-05-22-p2a-gap-input-postedit-analyze.json --benchmark-label p2a-gap-input-postedit` passed before detect-changes and produced scanned `739`, parsed `550`, unsupported `189`, failed `0`, graph nodes `22806`, and graph relationships `52584`.
- AVmatrix change detection after staging the new source files used fresh analyze output followed by `.\avmatrix-launcher\server-bundle\avmatrix.exe detect-changes --repo AVmatrix --scope all`; it reported changed_count `56`, changed_files `5`, affected_count `0`, and risk_level `low`.
- No visible Web UI behavior changed in this slice, so browser/e2e validation was not required for this slice.

## E8 - ResolutionGap Persistence Evidence

Status: complete for Phase 3.

Record during P3.

Persisted schema and graph representation:

- Graph node label: `ResolutionGap`.
- Graph diagnostic relationship type: `HAS_RESOLUTION_GAP`.
- Node identity: `ResolutionGap:<sourceSiteID>` when source-site ID exists; fallback identity includes source node, fact family, target text, target role, source-site status, proof kind, classification, and actionability so different target/status/proof/actionability buckets are not silently merged.
- Relationship identity: `rel:has-resolution-gap:<sourceNodeID>-><ResolutionGapID>`.
- Persisted node fields: `gapKind`, `sourceSiteId`, `sourceNodeId`, `sourceNodeLabel`, `sourceAppLayer`, `sourceFunctionalArea`, `factFamily`, `targetText`, `targetRole`, `sourceSiteStatus`, `proofKind`, `classification`, `actionability`, `resolutionSource`, `source`, `filePath`, `fileHash`, `startLine`, `startCol`, `endLine`, `endCol`, `count`, `note`, `appLayer`, and `functionalArea`.
- LadybugDB schema/export/load: `ResolutionGap` node table added; `HAS_RESOLUTION_GAP` relationship type added; relationship pairs from all existing node tables to `ResolutionGap` added; CSV export preserves all persisted gap fields.
- Web contract/UI constants: generated schema and TypeScript include `ResolutionGap` and `HAS_RESOLUTION_GAP`; Web node color/size/filter icon and edge metadata now cover them.

Source-backed examples and typed metadata:

- `TestSourceBackedResolutionGapInputsPreserveSourceSiteEvidence` now proves a source-backed unresolved call input can materialize a `ResolutionGap` node and `HAS_RESOLUTION_GAP` relationship while preserving source-site ID/status, proof kind, target role, classification, actionability, App Layer, Functional Area, file/range, count, and note.
- `TestApplyPersistsSourceBackedResolutionGaps` proves semantic enrichment persists a call gap from source-backed diagnostics after App Layer/Functional Area classification and does not synthesize `Function:stop` or a proofless `CALLS` edge.
- `TestExportGraphCSVsWritesResolutionGapNodesAndPairs` proves LadybugDB CSV export writes one `ResolutionGap` row, one `Function -> ResolutionGap` pair file, and zero skipped graph facts for the fixture.
- `TestResolutionGapInputInfersTargetRole` proves target-role inference preserves explicit roles, maps call/access/type-reference/heritage fact families to callable/member/type roles, maps builtin/standard-library/test-framework/external classifications when fact family does not decide, and writes the inferred role to both persisted gap nodes and `HAS_RESOLUTION_GAP` relationships.
- `TestValidateResolutionGapPersistenceAcceptsSourceBackedGaps` and `TestValidateResolutionGapPersistenceFlagsFakeResolvedOrTopologyClaims` prove persisted gaps may exist as diagnostic entities while validator failures catch dangling gap relationships, gap relationships targeting real code nodes, gap nodes that claim resolved targets, and resolved semantic edges that reuse unresolved source-site IDs.
- `TestApplyPersistsResolutionGapRolesClassificationsAndOccurrences` covers call, access, type-reference, heritage, builtin/predeclared, standard-library, test-framework, external, in-repo analyzer-gap, unknown target-role fallback, repeated occurrence count `3`, and two different target texts from the same source/fact/file/note fixture. The test proves target text, sourceSiteID, source-site status, proof kind, classification, actionability, App Layer, Functional Area, and count survive persistence without synthetic target nodes or resolved semantic edges.
- `ResolutionGapAggregates` and `SourceBackedResolutionGapAggregates` implement the Phase 3 aggregate/dedupe policy as an inventory layer over source-backed inputs. The aggregate keeps target text in bucket identity, sums exact occurrence counts, keeps full `sourceSiteIds`, records App Layer/Functional Area/file distributions, and caps only representative samples when requested.
- `TestResolutionGapAggregatesPreserveCountsSamplesAndDistributions` proves repeated source sites with counts `3` and `2` aggregate to occurrence count `5`, keep both sourceSiteIDs, preserve App Layer and Functional Area occurrence distributions, cap samples to one representative sample without losing counts, and keep a different target text in a different bucket.
- `TestSourceBackedResolutionGapAggregatesUseGraphDiagnostics` proves graph diagnostics can feed aggregate buckets directly, preserving access/member role, source-site traceability, occurrence count `5`, and samples from source-backed diagnostics.
- Fresh analyze produced typed gap counts: `unresolved_call=31220`, `unresolved_access=18245`, `unresolved_type_reference=8611`, `unresolved_heritage=7`.
- Fresh analyze produced classification counts: `in_repo_unresolved=33825`, `builtin=9854`, `test_framework=7154`, `standard_library=7092`, `external_library=158`.
- Fresh analyze produced actionability counts: `analyzer_gap=33825`, `non_actionable=24100`, `review=158`.
- Fresh analyze produced target-role counts: `callable=31220`, `member=18245`, `type=8618`.
- Latest role-validation analyze produced typed gap counts: `unresolved_call=31307`, `unresolved_access=18383`, `unresolved_type_reference=8653`, `unresolved_heritage=7`.
- Latest role-validation analyze produced classification counts: `in_repo_unresolved=34009`, `builtin=9880`, `test_framework=7205`, `standard_library=7098`, `external_library=158`.
- Latest role-validation analyze produced actionability counts: `analyzer_gap=34009`, `non_actionable=24183`, `review=158`.
- Latest role-validation analyze produced target-role counts: `callable=31307`, `member=18383`, `type=8660`.

Before/after and validation evidence:

- Fresh analyze command: `.\avmatrix-launcher\server-bundle\avmatrix.exe analyze --force --benchmark-json .\.tmp\2026-05-22-p3-resolution-gap-persist-analyze.json --benchmark-label p3-resolution-gap-persist`.
- Fresh graph inventory after persistence: nodes `80957`, relationships `110949`, `ResolutionGap` nodes `58083`, `HAS_RESOLUTION_GAP` relationships `58083`.
- Semantic enrichment benchmark after persistence: `resolutionGapInputs=58083`, `resolutionGapNodes=58083`, `resolutionGapRelationships=58083`, semantic phase latency `627.8377 ms`.
- LadybugDB load evidence from the same analyze: node rows `80957`, relationship rows `110949`, node copy count `17`, relationship copy count `86`, fallback inserts `0`, skipped relationships `0`.
- Field preservation inventory from `.avmatrix\graph.json`: missing source node ID `0`, missing sourceSiteID `0`, missing source-site status `0`, missing proof kind `0`, missing target text `0`, missing source App Layer `0`, missing actionability `0`; `7716` gaps have unknown source Functional Area because the source node was not functionally classified yet.
- Validation commands passed after full build: `go test .\internal\graphhealth .\internal\semantic .\internal\lbugschema .\internal\lbugload .\internal\contracts .\internal\scopeir`; `go test .\internal\...`; `go test .\cmd\...`; `go run .\cmd\generate-web-contracts --check`; `npm --prefix .\avmatrix-web run test -- --run`; `npm --prefix .\avmatrix-web run test:e2e` with `14` passed and `30` skipped.
- Role-validation analyze command: `.\avmatrix-launcher\server-bundle\avmatrix.exe analyze --force --benchmark-json .\.tmp\2026-05-22-p3-role-validation-postedit-analyze.json --benchmark-label p3-role-validation-postedit`.
- Role-validation graph inventory after inference/validation fixture slice: files scanned `741`, parsed `552`, unsupported `189`, failed `0`, graph nodes `81298`, relationships `111411`, `ResolutionGap` nodes `58350`, and `HAS_RESOLUTION_GAP` relationships `58350`.
- Semantic enrichment benchmark after role-validation slice: `resolutionGapInputs=58350`, `resolutionGapNodes=58350`, `resolutionGapRelationships=58350`, semantic phase latency `609.862 ms`.
- LadybugDB load evidence from the same analyze: node rows `81298`, relationship rows `111411`, node copy count `17`, relationship copy count `86`, fallback inserts `0`, skipped relationships `0`.
- Field preservation inventory from the latest `.avmatrix\graph.json`: missing source node ID `0`, missing sourceSiteID `0`, missing source-site status `0`, missing proof kind `0`, missing target text `0`, missing source App Layer `0`; `7821` gaps have unknown source Functional Area because the source node was not functionally classified yet.
- Source-site accuracy command after role-validation slice wrote `.tmp\2026-05-22-p3-role-validation-source-site-accuracy.json`: relationship source-site occurrences `85879`, diagnostic source-site occurrences `59097`, all source-site occurrences `144976`, stable source-site IDs `144976`, missing IDs `0`, resolved `CALLS=7733`, resolved `ACCESSES=3490`, unresolved diagnostics `59097`, non-property ACCESSES targets `0`, false resolved edge candidates `0`, resolved edges without proof `0`, resolved edges without sourceSiteID `0`, low-confidence fallback resolved edges `0`, and coarse file call edges `0`.
- Validation commands passed after full build for role-validation slice: `powershell -ExecutionPolicy Bypass -File .\avmatrix-launcher\build.ps1`; `go test .\internal\graphhealth .\internal\semantic`; `go test .\internal\...`; `go test .\cmd\...`.
- AVmatrix change detection after staging the role-validation slice used fresh analyze output followed by `.\avmatrix-launcher\server-bundle\avmatrix.exe detect-changes --repo AVmatrix --scope all`; it reported changed_count `344`, changed_files `8`, affected_count `0`, and risk_level `low`.
- Aggregate-policy analyze command: `.\avmatrix-launcher\server-bundle\avmatrix.exe analyze --force --benchmark-json .\.tmp\2026-05-22-p3-aggregate-policy-postedit-analyze.json --benchmark-label p3-aggregate-policy-postedit`.
- Aggregate-policy graph inventory after the new aggregate code/tests: files scanned `743`, parsed `554`, unsupported `189`, failed `0`, graph nodes `81522`, relationships `111801`, `ResolutionGap` nodes `58495`, and `HAS_RESOLUTION_GAP` relationships `58495`.
- Semantic enrichment benchmark after aggregate-policy slice: `resolutionGapInputs=58495`, `resolutionGapNodes=58495`, `resolutionGapRelationships=58495`, semantic phase latency `595.5799 ms`.
- Aggregate inventory over the latest graph produced `58495` source-backed gap nodes, `35411` aggregate buckets, `59242` exact occurrences, `10854` buckets with multiple source sites, max bucket source sites `99`, and max bucket occurrences `99`. Top buckets by occurrences included `strings.HasSuffix` `99`, `strings.Contains` `97`, `any` `92`, `string` `63`, and `result.Metrics` access buckets `51` and `49`.
- Source-site accuracy command after aggregate-policy slice wrote `.tmp\2026-05-22-p3-aggregate-policy-source-site-accuracy.json`: relationship source-site occurrences `86148`, diagnostic source-site occurrences `59242`, all source-site occurrences `145390`, stable source-site IDs `145390`, missing IDs `0`, resolved `CALLS=7751`, resolved `ACCESSES=3548`, unresolved diagnostics `59242`, non-property ACCESSES targets `0`, false resolved edge candidates `0`, resolved edges without proof `0`, resolved edges without sourceSiteID `0`, low-confidence fallback resolved edges `0`, and coarse file call edges `0`.
- Validation commands passed after full build for aggregate-policy slice: `powershell -ExecutionPolicy Bypass -File .\avmatrix-launcher\build.ps1`; `go test .\internal\graphhealth`; `go test .\internal\...`; `go test .\cmd\...`.
- AVmatrix change detection after staging the aggregate-policy slice used fresh analyze output followed by `.\avmatrix-launcher\server-bundle\avmatrix.exe detect-changes --repo AVmatrix --scope all`; it reported changed_count `228`, changed_files `5`, affected_count `0`, and risk_level `low`.

Remaining P3 evidence gaps: none.

## E9 - Resolution Health And Inventory Evidence

Status: complete for Phase 4.

Record during P4.

Implemented Resolution Health model:

- `internal/graphhealth/policy.go` defines Resolution Health buckets separately from Topology Health: `resolved_references`, `unresolved_non_actionable`, `external_unresolved`, `in_repo_analyzer_gap`, `unresolved_call_target`, `unresolved_access_target`, `unresolved_type_target`, `unresolved_heritage_target`, and `unclassified_unknown`.
- `NodeHealth` now exposes `resolutionHealthBuckets`, `resolutionGapCount`, and `resolutionConfidence` alongside existing topology-only `topologyStatus`, counted degree, and graph-health diagnostic fields.
- `Summary` now exposes full inventory counts for ResolutionGap nodes, `HAS_RESOLUTION_GAP` relationships, gap occurrence counts, resolved references, Resolution Health buckets, Resolution Confidence, fact family, target role, classification, actionability, App Layer, Functional Area, topology status, and topology-plus-resolution overlay counts.
- `internal/httpapi/graph.go` includes the new Resolution Health fields in graph summary, graph-health explain component summaries, and graph-health report candidates.
- Generated Web contracts now include `GraphHealthResolutionHealthBucket`, `GraphHealthResolutionConfidence`, Resolution Health fields on `GraphHealthNodeMetadata`, `GraphHealthSummary`, `GraphHealthComponentExplanation`, and `GraphHealthReportCandidate`.
- `avmatrix-web/src/lib/graph-health-filters.ts` maps flat graph properties into the new generated `GraphHealthNodeMetadata` fields when structured graphHealth is not present.
- `avmatrix resolution-inventory` reads persisted `.avmatrix\graph.json`, calls the same `graphhealth.ComputeSummary` full-count path as API/Web graph responses, and writes JSON or summary lines without using capped `/api/graph/report` candidates.

Topology separation evidence:

- `TestCompute_ResolutionHealthOverlayPreservesTopology` proves a source node with one incoming counted `CALLS`, one outgoing counted `CALLS`, and an excluded `HAS_RESOLUTION_GAP` remains `connected` with counted degree `1/1`, while `resolutionConfidence=degraded`, `resolutionGapCount=3`, and Resolution Health buckets include `resolved_references=2`, `unresolved_call_target=3`, and `in_repo_analyzer_gap=3`.
- The same test proves `HAS_RESOLUTION_GAP` is counted as an excluded `other` edge and does not become a topology edge.
- `TestCompute_ResolutionHealthClassifiesExternalAndNonActionableGaps` covers non-actionable builtin gaps and external unresolved gaps as Resolution Health buckets without changing topology.
- Existing graph-health e2e coverage passed after Phase 4 with the frontend server started correctly: `14` passed and `30` skipped.

Fresh Phase 4 analyze and inventory artifacts:

- Analyze artifact: `.tmp\2026-05-22-p4-resolution-health-inventory-analyze.json`.
- CLI inventory artifact: `.tmp\2026-05-22-p4-resolution-inventory.json`.
- Source-site accuracy artifact: `.tmp\2026-05-22-p4-resolution-health-source-site-accuracy.json`.

Fresh analyze output:

```text
files: scanned=745 parsed=556 unsupported=189 failed=0
graph: nodes=82062 relationships=112614 path=E:\AVmatrix-GO\.avmatrix\graph.json
semantic: resolutionGapInputs=58879 resolutionGapNodes=58879 resolutionGapRelationships=58879
semantic_enrichment duration: 586.7456 ms
LadybugDB: nodeRows=82062 relationshipRows=112614 fallbackInsertCount=0 skippedRelationships=0
```

CLI inventory command output:

```powershell
.\avmatrix-launcher\server-bundle\avmatrix.exe resolution-inventory --graph .\.avmatrix\graph.json --out .\.tmp\2026-05-22-p4-resolution-inventory.json
```

```text
resolutionInventory.nodes=82062 relationships=112614 gapNodes=58879 gapRelationships=58879 gapOccurrences=59626 resolvedReferences=27195
resolutionHealth.resolvedReferences=27195 unresolvedNonActionable=24926 externalUnresolved=163 inRepoAnalyzerGap=34537 unclassifiedUnknown=0
resolutionHealth.targets.call=31590 access=18541 type=9488 heritage=7
resolutionConfidence.clear=370 degraded=4573 unknown=77119
resolutionGap.appLayers=api:4296,api_contract:627,api_test:5114,backend:18126,backend_test:20057,cli_launcher:553,config:88,frontend:4846,frontend_api_client:387,frontend_test:5510,generated_contract:22
resolutionGap.functionalAreas=analyzer:7084,api:4660,cli:3116,configuration:88,contracts:880,embeddings:1701,graph_health:2590,launcher:671,layout:665,mcp:4716,providers:12674,query:2493,resolution:4060,runtime:30,session:1062,storage:2639,unknown:7881,web_graph_ui:2616
resolutionGap.factFamilies=access:18541,call:31590,heritage:7,type-reference:9488
resolutionGap.targetRoles=callable:31590,member:18541,type:9495
resolutionGap.classifications=builtin:10149,external_library:163,in_repo_unresolved:34537,standard_library:7506,test_framework:7271
resolutionGap.actionability=analyzer_gap:34537,non_actionable:24926,review:163
resolutionGap.topology=connected:24054,detached_component:2659,no_incoming:28070,no_outgoing:2520,true_isolated:2323
```

Count table from CLI inventory JSON:

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

Resolution Health bucket counts:

| Bucket | Count |
| --- | ---: |
| resolved_references | 27195 |
| unresolved_non_actionable | 24926 |
| external_unresolved | 163 |
| in_repo_analyzer_gap | 34537 |
| unresolved_call_target | 31590 |
| unresolved_access_target | 18541 |
| unresolved_type_target | 9488 |
| unresolved_heritage_target | 7 |
| unclassified_unknown | 0 |

Topology overlay counts from gap occurrences:

| Topology status | Gap occurrences |
| --- | ---: |
| connected | 24054 |
| no_incoming | 28070 |
| no_outgoing | 2520 |
| detached_component | 2659 |
| true_isolated | 2323 |
| unknown_connectivity | 0 |

Cypher/LadybugDB verification:

- `MATCH (n:ResolutionGap) RETURN count(n) AS gapNodes` returned `58879`.
- `MATCH (a)-[r]->(b) WHERE r.type = 'HAS_RESOLUTION_GAP' RETURN count(r) AS gapRelationships` returned `58879`.
- `MATCH (n:ResolutionGap) RETURN n.gapKind AS gapKind, count(n) AS count ORDER BY count DESC LIMIT 10` returned call `31590`, access `18541`, type-reference node count `8741`, and heritage `7`. The type-reference node count is lower than occurrence count `9488` because CLI inventory counts source-site occurrences while Cypher count here counts gap nodes.
- `MATCH (n:ResolutionGap) RETURN n.sourceAppLayer AS appLayer, count(n) AS count ORDER BY count DESC LIMIT 5` returned top node-count layers `backend_test=19975`, `backend=17799`, `frontend_test=5480`, `api_test=5046`, and `frontend=4819`.
- `MATCH (n:ResolutionGap) RETURN n.functionalArea AS functionalArea, count(n) AS count ORDER BY count DESC LIMIT 5` returned `providers=12593`, `unknown=7821`, `analyzer=7000`, `mcp=4640`, and `api=4490`.

Validation evidence for Phase 4:

- Full build passed after contract regeneration: `powershell -ExecutionPolicy Bypass -File .\avmatrix-launcher\build.ps1`.
- Focused Go tests passed: `go test .\internal\graphhealth .\internal\cli .\internal\contracts .\internal\httpapi`.
- Contract check passed: `go run .\cmd\generate-web-contracts --check`.
- Wider Go tests passed: `go test .\internal\...` and `go test .\cmd\...`.
- Web unit tests passed: `npm --prefix .\avmatrix-web run test -- --run` with `44` files and `358` tests passed.
- Web e2e initially failed because no frontend server was listening on `127.0.0.1:5228`; rerun with a hidden Vite dev server passed: `14` passed and `30` skipped.
- Source-site accuracy after Phase 4 reported all source-site occurrences `146447`, missing IDs `0`, false resolved edge candidates `0`, non-property ACCESSES targets `0`, resolved edges without proof `0`, and coarse file call edges `0`.
- AVmatrix detect-changes after staging the Phase 4 slice used fresh analyze output followed by `.\avmatrix-launcher\server-bundle\avmatrix.exe detect-changes --repo AVmatrix --scope all`; it reported changed_count `605`, changed_files `15`, affected_count `41`, and risk_level `critical`. The critical blast radius is expected because this slice extends root CLI registration, graph-health summary computation, HTTP graph/report/explain payloads, generated Web contracts, and Web graph-health metadata fallback.

## E10 - Query Health Command Evidence

Status: complete for Phase 5.

Record during P5.

Command name and usage:

```powershell
.\avmatrix-launcher\server-bundle\avmatrix.exe query-health --suite .\docs\query-health\2026-05-22-avmatrix-app-layer-resolution-gap-suite.json --repo AVmatrix --out .\.tmp\2026-05-22-p5-query-health-baseline.json --limit 10
```

The command verifies the target repo is not stale by comparing indexed commit and current commit, then runs every suite intent through the same local MCP `query` path used by `avmatrix query`. It writes JSON with expected targets, actual top results, matched/missed targets, hit@5, hit@10, noise reason, pass/fail, and semantic fields if current query results return them. It also supports `--json` and `--fail-on-threshold`.

Suite fixture:

- `docs/query-health/2026-05-22-avmatrix-app-layer-resolution-gap-suite.json`
- Format: `schemaVersion`, `suite`, `description`, and `cases[]` with `id`, `intent`, `expectedFiles`, `expectedSymbols`, optional `expectedAppLayers`, optional `expectedFunctionalAreas`, `hitAt5Threshold`, and `hitAt10Threshold`.

The first suite includes expected files for:

- unresolved reference diagnostics: `internal/resolution/resolve.go`, `internal/resolution/emit.go`, `internal/graphhealth/diagnostics.go`;
- graph-health separation: `internal/graphhealth/compute.go`, `internal/graphhealth/policy.go`, `avmatrix-web/src/lib/graph-health-filters.ts`;
- layout and optimizer surfaces: `avmatrix-web/src/lib/graph-adapter.ts`, `avmatrix-web/src/hooks/useSigma.ts`;
- runtime reset and hidden terminal behavior: `avmatrix-launcher/src/main.go`;
- API/contract/query surfaces: `internal/httpapi/graph.go`, `internal/contracts/web_ui.go`, `internal/mcp/tools.go`, `internal/cli/tool_command.go`, `cmd/avmatrix/main.go`, `avmatrix-web/src/generated/avmatrix-contracts.ts`;
- frontend graph filter/detail/layout surfaces: `avmatrix-web/src/hooks/app-state/graph.tsx`, `avmatrix-web/src/components/FileTreePanel.tsx`, `avmatrix-web/src/components/GraphCanvas.tsx`.

Fresh Phase 5 analyze evidence:

```text
files: scanned=748 parsed=558 unsupported=190 failed=0
graph: nodes=82538 relationships=113396 path=E:\AVmatrix-GO\.avmatrix\graph.json
artifact: .tmp\2026-05-22-p5-query-health-command-analyze.json
```

Query-health baseline command output:

```text
queryHealth.suite=avmatrix-app-layer-resolution-gap-lens cases=7 passed=1 failed=6
queryHealth.case=unresolved-reference-diagnostic-generation status=FAIL hitAt5=1/2 hitAt10=1/3 expected=6
queryHealth.case=graph-health-unknown-connectivity-separation status=FAIL hitAt5=0/2 hitAt10=0/3 expected=5
queryHealth.case=app-layer-resolution-gap-layout status=FAIL hitAt5=0/2 hitAt10=0/3 expected=5
queryHealth.case=runtime-reset-hidden-terminal status=PASS hitAt5=4/2 hitAt10=4/3 expected=5
queryHealth.case=api-contract-surfaces status=FAIL hitAt5=0/2 hitAt10=2/3 expected=5
queryHealth.case=query-implementation-surfaces status=FAIL hitAt5=0/2 hitAt10=0/3 expected=7
queryHealth.case=frontend-graph-filter-surfaces status=FAIL hitAt5=0/2 hitAt10=0/3 expected=6
```

Noise examples:

- Most failing cases still return launcher `avmatrix-launcher/src/main.go`, `internal/graphhealth/diagnostics.go`, `internal/graphhealth/resolution_gap_inputs.go`, `internal/semantic/app_layer.go`, and `avmatrix-web/src/components/Header.tsx` as top files.
- `unresolved-reference-diagnostic-generation` hits `internal/graphhealth/diagnostics.go` but misses `resolve.go`, `emit.go`, `resolveCall`, `emitUnresolvedReference`, and `AppendDiagnosticToNode`.
- `query-implementation-surfaces` misses `internal/mcp/tools.go`, `internal/cli/tool_command.go`, `cmd/avmatrix/main.go`, `queryTool`, `rankedProcessMatches`, `matchingDefinitionRows`, and `newQueryCommand`.
- These failures are expected baseline evidence for Phase 6; Phase 5 added the repeatable measurement command, not retrieval-ranking changes.

Validation evidence:

- Full build passed before final tests: `powershell -ExecutionPolicy Bypass -File .\avmatrix-launcher\build.ps1`.
- Focused tests passed after the full build: `go test .\internal\cli` and `go test .\internal\mcp`.
- Tests cover suite parsing, scoring, missing expected targets, noisy results, semantic field output, JSON output, summary/table output, report writing, and threshold failure behavior.
- AVmatrix detect-changes after staging the Phase 5 slice used fresh analyze output followed by `.\avmatrix-launcher\server-bundle\avmatrix.exe detect-changes --repo AVmatrix --scope all`; it reported changed_count `493`, changed_files `8`, affected_count `22`, and risk_level `critical`. The critical scope is expected because this slice registers a root CLI command and adds new `runQueryHealth` command flows.

## E11 - Semantic Command Surface Evidence

Status: complete for Phase 6.

Record during P6.

| Command | Required evidence |
| --- | --- |
| `query` | node type, App Layer, Functional Area, Resolution Health, and related gap summary when available |
| `context` | symbol/node view with topology, resolution-health summary, and source/nearby gaps |
| `impact` | affected App Layers, affected Functional Areas, and resolution-health risks when supported |
| `detect-changes` | changed App Layers, changed Functional Areas, ResolutionGap changes, and resolution-health impact |

If a command cannot fully expose a semantic layer in this implementation, record the exact limitation and follow-up.

### P6-A - `query` Semantic Output And Retrieval Evidence

Status: complete for `query`.

Changed behavior:

- `query` output includes `semanticStatus` from the loaded graph and `semanticWarning` when App Layer or Functional Area metadata is stale/incomplete.
- `process_symbols` and `definitions` rows expose `type`, `appLayer`, `functionalArea`, `topologyStatus`, `resolutionConfidence`, `resolutionGapCount`, and related gap summary maps when persisted graph data provides them.
- Related ResolutionGap summaries are read from persisted `HAS_RESOLUTION_GAP` relationships and `ResolutionGap` nodes; unresolved targets are not converted into fake resolved symbols or topology edges.
- Retrieval ranking now tokenizes/stems query text, searches broader executable/file definitions, skips docs/tests unless requested, diversifies definition output by file, and boosts App Layer/Functional Area surfaces for graph-health, layout, query, API, and frontend graph-filter intents.

Focused test evidence:

```text
powershell -ExecutionPolicy Bypass -File .\avmatrix-launcher\build.ps1
go test .\internal\mcp .\internal\cli
```

Both commands passed after the P6-A implementation. New coverage in `internal/mcp/query_semantic_test.go` verifies:

- fresh semantic graphs return complete `semanticStatus` and no warning;
- stale/incomplete semantic graphs return a semantic warning;
- `query` definition and process-symbol rows preserve App Layer, Functional Area, topology status, resolution confidence, and ResolutionGap summaries.

Fresh analyze evidence after P6-A:

```text
files: scanned=749 parsed=559 unsupported=190 failed=0
graph: nodes=82935 relationships=113953 path=E:\AVmatrix-GO\.avmatrix\graph.json
artifact: .tmp\2026-05-22-p6a-query-semantic-output-final-analyze.json
```

Final query-health evidence on the refreshed graph:

```text
artifact: .tmp\2026-05-22-p6a-query-health-final.json
queryHealth.suite=avmatrix-app-layer-resolution-gap-lens cases=7 passed=7 failed=0
queryHealth.case=unresolved-reference-diagnostic-generation status=PASS hitAt5=4/2 hitAt10=4/3 expected=6
queryHealth.case=graph-health-unknown-connectivity-separation status=PASS hitAt5=4/2 hitAt10=5/3 expected=5
queryHealth.case=app-layer-resolution-gap-layout status=PASS hitAt5=5/2 hitAt10=5/3 expected=5
queryHealth.case=runtime-reset-hidden-terminal status=PASS hitAt5=4/2 hitAt10=4/3 expected=5
queryHealth.case=api-contract-surfaces status=PASS hitAt5=3/2 hitAt10=4/3 expected=5
queryHealth.case=query-implementation-surfaces status=PASS hitAt5=6/2 hitAt10=6/3 expected=7
queryHealth.case=frontend-graph-filter-surfaces status=PASS hitAt5=4/2 hitAt10=4/3 expected=6
```

Representative command output:

```powershell
.\avmatrix-launcher\server-bundle\avmatrix.exe query "query ranking process matching definitions CLI query command implementation" --repo AVmatrix --limit 10
```

The output includes `semanticStatus` with complete App Layer and Functional Area coverage for `82935` nodes, `definitions` rows for `internal/mcp/tools.go`, `internal/cli/tool_command.go`, and query-related symbols, and `process_symbols` rows such as `rankedProcessMatches`, `matchingDefinitionRows`, and `queryTool` with App Layer, Functional Area, and ResolutionGap summaries.

Pre-commit AVmatrix scope check:

```text
.\avmatrix-launcher\server-bundle\avmatrix.exe detect-changes --repo AVmatrix --scope all
summary: affected_count=17, changed_count=432, changed_files=5, risk_level=critical
```

The critical scope is expected for this slice because it intentionally changes the shared `queryTool`, `rankedProcessMatches`, and `matchingDefinitionRows` retrieval/output path and adds tests plus plan ledgers. The affected processes reported by AVmatrix are query-tool and matching/ranking flows such as `QueryTool -> MinInt`, `RankedProcessMatches -> QueryTokenStem`, and `MatchingDefinitionRows -> MinInt`.

### P6-B - `context` Semantic Output Evidence

Status: complete for `context`.

Changed behavior:

- `context` output now includes graph-level `semanticStatus` and `semanticWarning` when App Layer or Functional Area metadata is stale/incomplete.
- `symbol`, ambiguous `candidates`, incoming/outgoing reference rows, and process rows expose `type`, `appLayer`, `appLayerSource`, `functionalArea`, `functionalAreaSource`, `topologyStatus`, `resolutionConfidence`, `resolutionGapCount`, and `resolutionHealthBuckets` when persisted graph data provides them.
- Relationship rows now include source-site proof/status metadata such as `sourceSiteId`, `sourceSiteIds`, `sourceSiteCount`, `sourceSiteStatus`, `proofKind`, `targetRole`, `targetText`, relationship file/range, confidence, reason, and resolution source.
- Source-node gaps are exposed in `sourceResolutionGaps` with `resolutionRelation=source_node_gap` and `resolvedTarget=false`.
- When the selected node is a `ResolutionGap`, the symbol row is marked with `resolutionGapEntity=true`, `resolvedTarget=false`, and `resolutionRelation=resolution_gap_entity`; the real source nodes are exposed separately in `resolutionGapSources` with `resolutionRelation=gap_source_node`.

Focused test evidence:

```text
powershell -ExecutionPolicy Bypass -File .\avmatrix-launcher\build.ps1
go test .\internal\mcp .\internal\cli
go test .\internal\mcp -run "TestContextTool(ReturnsSemanticFieldsAndSourceResolutionGaps|DistinguishesResolutionGapEntityFromSourceNode|WarnsForStaleIncompleteSemanticMetadata)$" -count=1
```

All commands passed after the P6-B implementation. New coverage in `internal/mcp/context_semantic_test.go` verifies:

- fresh semantic graphs return complete `semanticStatus` and no stale warning;
- stale/incomplete semantic graphs return a semantic warning;
- symbol rows preserve node type, App Layer, Functional Area, topology status, resolution confidence, resolution-health buckets, and gap counts;
- outgoing references preserve semantic fields plus source-site proof/status metadata;
- source-node gaps and selected `ResolutionGap` entities are separate output concepts.

Fresh analyze evidence after P6-B:

```text
files: scanned=750 parsed=560 unsupported=190 failed=0
graph: nodes=83205 relationships=114335 path=E:\AVmatrix-GO\.avmatrix\graph.json
artifact: .tmp\2026-05-22-p6b-context-semantic-output-analyze.json
```

Representative command output artifacts:

```powershell
.\avmatrix-launcher\server-bundle\avmatrix.exe context --repo AVmatrix --uid "Function:internal/mcp/context.go:contextSymbolPayload#2" > .\.tmp\2026-05-22-p6b-context-symbol-output.txt
.\avmatrix-launcher\server-bundle\avmatrix.exe context --repo AVmatrix --uid "ResolutionGap:SourceSite:internal/mcp/context.go#call#any#197#12#204#2" > .\.tmp\2026-05-22-p6b-context-resolution-gap-output.txt
```

The symbol output includes complete `semanticStatus` for `83205` nodes, semantic fields on incoming/outgoing refs and process rows, relationship proof fields such as `sourceSiteStatus=resolved` and `proofKind`, and `sourceResolutionGaps` entries with `resolutionRelation=source_node_gap` and `resolvedTarget=false`.

The ResolutionGap output includes `resolutionGapEntity=true`, `resolvedTarget=false`, `sourceNodeId`, `targetText`, `semanticStatus`, and a separate `resolutionGapSources` array whose source row uses `resolutionRelation=gap_source_node`.

Pre-commit AVmatrix scope check:

```text
.\avmatrix-launcher\server-bundle\avmatrix.exe detect-changes --repo AVmatrix --scope all
summary: affected_count=26, changed_count=87, changed_files=5, risk_level=critical
```

The critical scope is expected for this slice because it intentionally changes shared `context` output paths: `contextToolInternal`, `contextSymbolPayload`, `contextRefPayload`, `contextNeighborhood`, and new semantic/gap helper functions. The affected flows are context neighborhood/symbol output flows plus shared ambiguous-candidate rows also used by impact/rename disambiguation output.

### P6-C - `impact` Semantic Output Evidence

Status: complete for `impact`.

Changed behavior:

- `impact` output now includes graph-level `semanticStatus` and `semanticWarning` when semantic metadata is stale/incomplete.
- target and `byDepth` impacted rows expose persisted `type`, `appLayer`, `appLayerSource`, `functionalArea`, `functionalAreaSource`, `topologyStatus`, `resolutionConfidence`, `resolutionGapCount`, and `resolutionHealthBuckets` when graph data provides them.
- relationship rows now include source-site proof/status metadata such as `sourceSiteId`, `sourceSiteIds`, `sourceSiteCount`, `sourceSiteStatus`, `proofKind`, `targetRole`, `targetText`, relationship file/range, confidence, reason, and resolution source.
- top-level `affectedAppLayers`, `affectedFunctionalAreas`, and `resolutionHealthRisks` summarize the impacted nodes.
- affected process/module rows carry semantic fields from their persisted graph nodes.
- HIGH/CRITICAL output now includes `workflowWarning` and `workflowWarningBlocksOutput=false`, making the blast-radius warning explicit workflow safety information while keeping inspection output available.

Focused test evidence:

```text
powershell -ExecutionPolicy Bypass -File .\avmatrix-launcher\build.ps1
go test .\internal\mcp .\internal\cli
go test .\internal\mcp -run "TestImpact(ToolReturnsSemanticAffectedLayerAndResolutionRiskSummary|HighCriticalRiskWarningKeepsInspectionOutput)$" -count=1
```

All commands passed after the P6-C implementation. New coverage in `internal/mcp/impact_semantic_test.go` verifies:

- affected App Layer and Functional Area counts are emitted from impacted rows;
- resolution-health risks count degraded nodes, nodes with gaps, total gap count, Resolution Health buckets, and risk node examples;
- target, impacted rows, affected processes, and affected modules preserve semantic fields;
- source-site proof/status metadata remains visible on impacted relationship rows;
- CRITICAL risk output includes a non-blocking workflow warning and still returns `byDepth`.

Fresh analyze evidence after P6-C:

```text
files: scanned=751 parsed=561 unsupported=190 failed=0
graph: nodes=83519 relationships=114735 path=E:\AVmatrix-GO\.avmatrix\graph.json
artifact: .tmp\2026-05-22-p6c-impact-semantic-output-analyze.json
```

Representative command output artifacts:

```powershell
.\avmatrix-launcher\server-bundle\avmatrix.exe impact --repo AVmatrix --uid "Function:internal/mcp/context.go:contextSymbolPayload#2" --direction upstream > .\.tmp\2026-05-22-p6c-impact-context-symbol-output.txt
.\avmatrix-launcher\server-bundle\avmatrix.exe impact --repo AVmatrix --uid "Function:internal/mcp/impact.go:runImpactBFSProfiled#4" --direction upstream > .\.tmp\2026-05-22-p6c-impact-critical-warning-output.txt
```

The `contextSymbolPayload` impact output includes `affectedAppLayers`, `affectedFunctionalAreas`, semantic fields on affected process rows, and `resolutionHealthRisks`.

The `runImpactBFSProfiled` impact output has `risk=CRITICAL`, retains `byDepth`, `affected_processes`, and `affected_modules`, and includes `workflowWarningBlocksOutput=false` with wording that the warning is workflow safety information, not a blocker.

Pre-commit AVmatrix scope check:

```text
.\avmatrix-launcher\server-bundle\avmatrix.exe detect-changes --repo AVmatrix --scope all
summary: affected_count=9, changed_count=127, changed_files=4, risk_level=high
```

The high scope is expected for this slice because it intentionally changes the shared `impact` output path: `runImpactBFSProfiled`, `impactItemPayload`, `impactAffectedProcesses`, `impactAffectedModules`, and semantic summary helpers. The affected flows are impact output flows and shared resource helpers used to render impacted symbols, process rows, and module rows.

### P6-D - `detect-changes` Semantic Output Evidence

Status: complete for `detect-changes`.

Changed behavior:

- `detect-changes` output now includes graph-level `semanticStatus` and `semanticWarning` when semantic metadata is stale/incomplete.
- changed symbol rows now include persisted `type`, `appLayer`, `appLayerSource`, `functionalArea`, `functionalAreaSource`, `topologyStatus`, `resolutionConfidence`, `resolutionGapCount`, and `resolutionHealthBuckets` when graph data provides them.
- changed `ResolutionGap` entity rows include `resolutionGapEntity=true`, `resolvedTarget=false`, source node/layer/area fields, `gapKind`, `factFamily`, `targetText`, `targetRole`, `sourceSiteStatus`, `proofKind`, `classification`, `actionability`, range, count, and note fields from persisted graph data.
- affected process rows now include persisted semantic fields plus `changedStepAppLayers`, `changedStepFunctionalAreas`, changed-step semantic fields, and per-process `resolutionHealthImpact` when changed steps carry gap/degraded evidence.
- top-level and summary output now include `changedAppLayers`, `changedFunctionalAreas`, `affectedAppLayers`, `affectedFunctionalAreas`, `resolutionGapChanges`, and `resolutionHealthImpact` from persisted graph evidence.

Focused test evidence:

```text
powershell -ExecutionPolicy Bypass -File .\avmatrix-launcher\build.ps1
go test .\internal\mcp .\internal\cli
go test .\internal\mcp -run "TestDetectChangesToolReturnsSemanticSummaries|TestServeCallToolDetectChanges|TestServeCallToolDetectChangesReportsDeletedSymbols" -count=1
```

All commands passed after the P6-D implementation. New coverage in `internal/mcp/detect_changes_semantic_test.go` verifies:

- changed App Layer and Functional Area counts are emitted from changed rows;
- affected App Layer and Functional Area counts are emitted from affected process rows;
- changed `ResolutionGap` rows preserve target text, actionability, and entity marker fields;
- source changed rows preserve Resolution Health fields;
- `resolutionGapChanges` counts changed gap entities, gap occurrence counts, source nodes with gaps, source resolution gap counts, and top targets;
- `resolutionHealthImpact` counts degraded nodes, nodes with gaps, total gap count, and Resolution Health buckets;
- affected process rows expose semantic fields and changed-step App Layer summaries.

Fresh analyze evidence after P6-D:

```text
files: scanned=752 parsed=562 unsupported=190 failed=0
graph: nodes=83783 relationships=115138 path=E:\AVmatrix-GO\.avmatrix\graph.json
artifact: .tmp\2026-05-22-p6d-detect-semantic-output-analyze.json
```

Representative command output artifact:

```powershell
.\avmatrix-launcher\server-bundle\avmatrix.exe detect-changes --repo AVmatrix --scope all > .\.tmp\2026-05-22-p6d-detect-changes-semantic-output.txt
```

The command output includes `semanticStatus`, `changedAppLayers`, `changedFunctionalAreas`, `affectedAppLayers`, `affectedFunctionalAreas`, `changedStepAppLayers`, `resolutionGapEntity`, `resolutionGapChanges`, and `resolutionHealthImpact`. On the implementation-only P6-D diff before ledger updates it reported `changed_count=156`, `affected_count=15`, `changed_files=1`, `risk_level=high`, `changed_app_layers.api=156`, `affected_app_layers.api=15`, `changedGapEntities=94`, and `changedGapOccurrenceCount=95`.

Pre-commit AVmatrix scope check:

```text
.\avmatrix-launcher\server-bundle\avmatrix.exe detect-changes --repo AVmatrix --scope all
summary: affected_count=15, changed_count=164, changed_files=4, risk_level=high
artifact: .tmp\2026-05-22-p6d-detect-changes-precommit-output.txt
```

The high scope is expected for this slice because it intentionally changes the shared `detect-changes` output path: `detectChangesTool`, `detectChangedSymbols`, `detectAffectedProcesses`, and new semantic summary helpers. The final pre-commit diff also includes the plan, evidence, and benchmark ledger updates. The affected flows are detect-changes output flows and shared semantic helper paths used to render changed symbols, affected process rows, changed-step summaries, ResolutionGap changes, and Resolution Health impact.

### P6-E - API MCP Semantic Output Evidence

Status: complete for `route_map`, `shape_check`, and `api_impact`.

Changed behavior:

- `route_map`, `shape_check`, and `api_impact` now include graph-level `semanticStatus` and `semanticWarning` when semantic metadata is stale/incomplete.
- route rows now carry persisted App Layer, Functional Area, Topology Health, Resolution Confidence, ResolutionGap count, and Resolution Health bucket fields.
- route consumers now carry the same semantic fields from their persisted consumer graph nodes.
- linked execution flows are still exposed as the existing `flows`/`executionFlows` string list, and now also appear as `flowDetails`/`executionFlowDetails` rows with process semantic fields.
- `api_impact` summary now includes route semantic fields, consumer App Layer/Functional Area counts, flow App Layer/Functional Area counts, and `resolutionHealthImpact` when route, consumer, or flow rows carry gap/degraded evidence.

Focused test evidence:

```text
powershell -ExecutionPolicy Bypass -File .\avmatrix-launcher\build.ps1
go test .\internal\mcp .\internal\cli -count=1
go test .\internal\mcp -run "TestServeCallToolRouteAndToolMaps" -count=1
```

All commands passed after the P6-E implementation. `TestServeCallToolRouteAndToolMaps` now verifies:

- `route_map` route rows expose `appLayer=api`, `functionalArea=api`, and Resolution Confidence from the route node;
- `route_map` consumer rows expose `appLayer=frontend`, `functionalArea=web_graph_ui`, and `resolutionGapCount=1` from the consumer node;
- `route_map` `flowDetails` exposes process App Layer and flow name while preserving the existing `flows` list;
- `shape_check` route and consumer rows preserve semantic fields while still reporting mismatches;
- `api_impact` route, consumer, `executionFlowDetails`, and `impactSummary.consumerAppLayers` preserve semantic fields.

Fresh analyze evidence after P6-E:

```text
files: scanned=752 parsed=562 unsupported=190 failed=0
graph: nodes=83982 relationships=115558 path=E:\AVmatrix-GO\.avmatrix\graph.json
artifact: .tmp\2026-05-22-p6e-api-tools-semantic-output-analyze.json
```

Representative MCP output artifacts:

```powershell
route_map artifact: .tmp\2026-05-22-p6e-route-map-semantic-output.txt
shape_check artifact: .tmp\2026-05-22-p6e-shape-check-semantic-output.txt
api_impact artifact: .tmp\2026-05-22-p6e-api-impact-semantic-output.txt
```

The current AVmatrix graph has no persisted route nodes, so the live MCP artifacts return empty/error route payloads. They still include `semanticStatus` from the persisted graph, proving the API-specific tools surface graph semantic state even when no route rows match. Populated route/consumer/flow semantic rows are covered by the focused fixture test above.

Pre-commit AVmatrix scope check:

```text
.\avmatrix-launcher\server-bundle\avmatrix.exe detect-changes --repo AVmatrix --scope all
summary: affected_count=29, changed_count=273, changed_files=6, risk_level=critical
artifact: .tmp\2026-05-22-p6e-api-tools-semantic-precommit-output.txt
```

The critical scope is expected for this slice because it intentionally changes the shared API MCP route index and output structs used by `route_map`, `shape_check`, `api_impact`, and `tool_map` flow-name reuse. The affected surfaces match `buildMCPRouteIndex`, route consumer and flow indexing, API impact record shaping, semantic field helpers, route fixture tests, and plan ledgers.

### P6-F / P6-G - Semantic Command Edge Tests And Ledger Closure

Status: complete for Phase 6 command-surface edge coverage and ledger records.

Changed test coverage:

- `impact` now has focused stale/incomplete semantic metadata coverage proving the payload emits `semanticStatus`, emits `semanticWarning`, and does not invent `appLayer` fields for symbols whose graph nodes do not carry persisted App Layer metadata.
- `detect-changes` now has focused stale/incomplete semantic metadata coverage proving changed symbol rows preserve the missing-data state rather than guessing App Layer at command-output time.
- API MCP command surfaces now have focused stale/missing-field coverage for `route_map`, `shape_check`, and `api_impact`; route, consumer, and route-impact payloads warn about stale semantic evidence and do not synthesize App Layer fields when the persisted graph does not provide them.
- The P6-A through P6-E sections above already record changed output fields, representative command artifacts, live limitations, and populated fixture behavior for `query`, `context`, `impact`, `detect-changes`, `route_map`, `shape_check`, and `api_impact`.

Validation evidence:

```text
powershell -ExecutionPolicy Bypass -File .\avmatrix-launcher\build.ps1
go test .\internal\mcp -run "Test(ImpactToolWarnsForStaleIncompleteSemanticMetadata|DetectChangesToolWarnsForStaleIncompleteSemanticMetadata|APIMCPToolsWarnForStaleAndDoNotInventSemanticFields)$" -count=1
go test .\internal\mcp .\internal\cli -count=1
```

All commands passed after adding `internal/mcp/semantic_command_surface_edge_test.go`.

Pre-commit AVmatrix scope check after staging the P6-F/P6-G slice:

```text
.\avmatrix-launcher\server-bundle\avmatrix.exe detect-changes --repo AVmatrix --scope all
summary: affected_count=0, changed_count=125, changed_files=4, risk_level=low
changed_app_layers: api_test=117, docs=8
changed_functional_areas: mcp=117, documentation=8
changedGapEntities=92, changedGapOccurrenceCount=92
artifact: .tmp\2026-05-22-p6f-semantic-command-edge-tests-staged-output.txt
```

The low scope is expected for this slice because it adds focused MCP test coverage and updates the Phase 6 plan ledgers. No affected execution flows were reported.

Command-surface limitations recorded for Phase 6:

- Semantic command surfaces expose App Layer, Functional Area, ResolutionGap, and Resolution Health only from persisted graph data.
- When graph metadata is stale or incomplete, command output returns `semanticStatus`/`semanticWarning` rather than classifying nodes at command/API/UI load time.
- The current AVmatrix graph has no persisted route nodes, so live `route_map`, `shape_check`, and `api_impact` artifacts show semantic status on empty/error route payloads; fixture tests prove populated route, consumer, flow-detail, and impact-summary semantic fields.

## E12 - Web UI Evidence

Status: pending

Record during P7.

Required evidence:

- screenshots or Playwright screenshots for Backend/API/Frontend rings;
- proof that API is placed between Backend and Frontend, and contract rings sit near API when present;
- screenshots or Playwright screenshots for additional rings when present;
- proof that same node type/color islands remain grouped inside rings;
- ring size, spacing, ordering, and default visibility behavior;
- App Layer filter behavior;
- Resolution Health filter behavior;
- explicit lens rows for Backend unresolved calls, API unresolved handlers/contracts, Frontend unresolved type refs, Shared contract analyzer gaps, External unresolved symbols, Builtin/Test/Stdlib non-actionable references, In-repo analyzer gaps, Resolution gaps by functional area, Top app layers by analyzer gap count, Top functional areas by unresolved count, and Top unresolved target text;
- node detail fields for App Layer, Functional Area, Topology Health, Resolution Health, and gaps;
- proof that optimizer is manual-only and not auto-run after render/load/filter changes;
- Web unit/e2e test names/results.

## E13 - Full Validation Evidence

Status: pending

Record during P8.

Required evidence:

- full build command/output;
- backend test command/output;
- contract generation/check command/output;
- proof-based CALLS/ACCESSES golden corpus command/output;
- source-site inventory and resolved-edge accuracy benchmark output;
- Web unit command/output;
- Web e2e command/output;
- query-health benchmark output;
- resolution inventory output;
- `query`, `context`, `impact`, and `detect-changes` examples;
- AVmatrix detect-changes output for implementation commits;
- final commit hashes;
- residual risks and follow-up plan if any validation fails.
