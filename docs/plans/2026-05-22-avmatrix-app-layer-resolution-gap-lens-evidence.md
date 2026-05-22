# AVmatrix App Layer Resolution Gap Lens Evidence Ledger

Date: 2026-05-22

Status: in progress; Phase 2 complete; Phase 2A proof-based CALLS/ACCESSES gate added before Phase 3

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

Status: pending

Record during P0-A.

| Discussion area | Plan task | Evidence result |
| --- | --- | --- |
| Node type is insufficient | pending | pending |
| Graph/API, not UI, must answer ownership | pending | pending |
| BE/API/FE/App Layer rings | pending | pending |
| Non-overlapping mixed App Layer categories | pending | pending |
| API as first-class layer | pending | pending |
| Functional Area accuracy gate | pending | pending |
| Proof-based CALLS/ACCESSES and source-site inventory | pending | pending |
| Persisted ResolutionGap/UnresolvedSymbol | pending | pending |
| Fine-grained gap relations | pending | pending |
| Resolution Health separate from Topology Health | pending | pending |
| Query-health command | pending | pending |
| Query/context/impact/detect-changes semantic output | pending | pending |
| API-specific MCP semantic output | pending | pending |
| Multi-ring layout and same-color islands | pending | pending |
| No dead-code verdict from unresolved refs alone | pending | pending |
| No timeout/auto optimizer behavior | pending | pending |
| No stale graph fallback | pending | pending |
| No evidence loss for graph-size reasons | pending | pending |
| User-facing naming consistency | pending | pending |

## E2 - Baseline Analyze Evidence

Status: pending

Record during P0-B.

Required evidence:

- command used;
- command output or reproducible notes;
- graph timestamp/hash if available;
- scanned/parsed/unsupported/failed counts;
- node/relationship/counted-semantic-relationship/execution-flow counts;
- `unknown_connectivity` count;
- comparison with discussion observations: about `22010` nodes, `26906` counted semantic relationships, `0` `unknown_connectivity`, `51232` unresolved occurrences, and `8880` unresolved buckets;
- repo path and analyzer version/commit if available.

## E3 - Source Trace Evidence

Status: pending

Record during P0-C and P0-G.

| Surface | Files/symbols found | Notes |
| --- | --- | --- |
| graph schema/snapshot | `internal/graph/types.go`, `internal/analyze/analyze.go`, `.avmatrix/graph.json` | App Layer is persisted on node properties as `appLayer`/`appLayerSource`; Functional Area is persisted as `functionalArea`/`functionalAreaSource`. Fresh graph snapshot from the locally built CLI has zero missing Functional Area fields across 22358 nodes. |
| analyze semantic enrichment flow | `internal/analyze/analyze.go`, `internal/semantic/app_layer.go`, `internal/semantic/functional_area.go` | The `semantic_enrichment` phase runs after processes and before graph compact, LadybugDB load, embeddings, and graph snapshot. Phase order from benchmark: scan, structure, documents, cobol, parse, routes, tools, orm, cross_file_binding, resolution, mro, communities, processes, semantic_enrichment, db_load. Functional Area assignment runs in the same enrichment pass as App Layer after process/community signals exist. |
| semantic enrichment input indexes and complexity | `internal/semantic/app_layer.go`, `internal/semantic/functional_area.go` | Enrichment builds App Layer and Functional Area path caches, `nodeID -> index`, `nodeID -> appLayer`, and `nodeID -> functionalArea` maps, then performs one relationship scan for Process/Community inference. It uses graph facts only; it does not rescan files or reparse ASTs. |
| LadybugDB export/load | `internal/lbugschema/schema.go`, `internal/lbugload/csv.go`, `internal/lbugload/load_test.go` | Node schemas and COPY CSV columns include both `appLayer` and `functionalArea`; benchmark DB load wrote 22358 node rows and 55349 relationship rows with zero fallback inserts. |
| resolved/unresolved call emission | `internal/resolution/resolve.go`, `internal/resolution/emit.go`, `internal/resolution/indexes.go` | Must trace fallback order, confidence values, proof kinds, and whether `resolveGlobalCallName` can emit a resolved `CALLS` edge. |
| resolved/unresolved access emission | `internal/resolution/resolve.go`, `internal/resolution/emit.go`, `internal/resolution/indexes.go` | Must trace property/member resolution, `propertyLabels()`, and whether non-property labels are emitted as ACCESSES. |
| call/access source facts | `internal/scopeir/facts.go`, `internal/providers/*/references.go`, `internal/providers/*/extract_test.go`, `internal/providers/provider_parity_test.go` | Must prove sourceSiteID/status/proof metadata starts from raw provider facts and survives relationship dedupe. |
| relationship dedupe/proof metadata | `internal/resolution/emit.go`, `internal/graph/types.go` | `semanticEdgeKey` and `Relationship` currently do not preserve all source-site/proof fields. |
| graph accuracy command/report | `internal/graphaccuracy/graphaccuracy.go` | Existing accuracy checks need extension for ACCESSES, source-site inventory, false resolved edges, and low-confidence fallback counts. |
| unresolved type-reference emission | pending | pending |
| unresolved heritage emission | pending | pending |
| diagnostic attachment | pending | pending |
| graph-health summary/report | pending | pending |
| HTTP graph payload | `internal/httpapi/graph.go`, `internal/semantic/metadata.go` | HTTP graph responses pass node `appLayer`/`appLayerSource` and `functionalArea`/`functionalAreaSource` properties through and include `semanticStatus`; NDJSON starts with `semantic_status`. Missing App Layer or Functional Area metadata is reported as stale/incomplete schema evidence and is not classified in the API loader. |
| generated Web contracts | `internal/contracts/web_ui.go`, `contracts/web-ui/avmatrix-web-contract.schema.json`, `avmatrix-web/src/generated/avmatrix-contracts.ts` | Contract manifest exposes App Layer and Functional Area enums/labels plus semantic status terms. Generated TypeScript exposes `FUNCTIONAL_AREAS`, `FUNCTIONAL_AREA_LABELS`, `FunctionalArea`, `GraphSemanticStatus.functionalArea`, `NodeProperties.functionalArea`, and `NodeProperties.functionalAreaSource` in addition to the App Layer fields. |
| query command | pending | pending |
| context command | pending | pending |
| impact command | pending | pending |
| detect-changes command | pending | pending |
| API-specific MCP tools | pending | pending |
| Web graph app state | pending | pending |
| Web graph filters/detail/layout | pending | pending |

## E4 - Baseline Unresolved And App Layer Evidence

Status: pending

Record during P0-D and P0-E.

Required unresolved evidence:

- unresolved bucket and occurrence counts;
- fact family counts;
- diagnostic classification/actionability counts if current graph has them;
- top target texts and examples;
- source node labels and path buckets;
- examples where source topology is known but resolution confidence is degraded.

Required provisional App Layer evidence:

- path/package rules used only for sizing;
- classification notes for these seed paths when present: `avmatrix-web/src/**`, `avmatrix-web/test/**`, `avmatrix-web/e2e/**`, `internal/**`, `cmd/**`, `internal/httpapi/**`, `avmatrix-web/src/services/backend-client.ts`, `avmatrix-launcher/**`, `contracts/**`, `internal/contracts/**`, `cmd/generate-web-contracts/**`, `docs/**`, `reports/**`, `*.md`, `*_test.go`, `test/fixtures/**`, config files, package files, and build scripts;
- counts by backend, api, frontend, cli_launcher, shared_contract, api_contract, api_shared_contract, frontend_api_client, backend_test, frontend_test, api_test, generated_contract, docs, config, generated, mixed, and unknown;
- examples proving why ambiguous buckets should remain unknown until better rules exist.

## E5 - Baseline Query Evidence

Status: pending

Record during P0-F.

| Intent | Expected files/symbols | Actual top results | Hit/miss | Noise reason |
| --- | --- | --- | --- | --- |
| unresolved reference diagnostic generation | pending | pending | pending | pending |
| graph health unknown-connectivity separation | pending | pending | pending | pending |
| App Layer and ResolutionGap layout | pending | pending | pending | pending |
| runtime reset hidden-terminal behavior | pending | pending | pending | pending |
| API contract surfaces | pending | pending | pending | pending |
| frontend graph filter surfaces | pending | pending | pending | pending |

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

Status: pending

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

## E8 - ResolutionGap Persistence Evidence

Status: pending

Record during P3.

Required evidence:

- persisted schema or graph representation;
- source-backed gap examples for unresolved call, access, type-reference, and heritage;
- sourceSiteID, source-site status, and proof kind preservation when a gap originates from Phase 2A call/access source-site inventory;
- examples for builtin/predeclared, standard-library, test-framework, external, in-repo analyzer gap, unclassified, and unknown roles;
- fine-grained relation or typed metadata examples;
- proof that fake resolved in-repo nodes/edges were not created;
- proof that aggregation or dedupe preserves exact counts, buckets, source samples, App Layer/Functional Area distributions, and traceability;
- backend test names/results;
- before/after persisted gap counts.

## E9 - Resolution Health And Inventory Evidence

Status: pending

Record during P4.

Required evidence:

- separate Topology Health and Resolution Health payload examples;
- connected node with gap remains connected;
- `no_incoming` or detached node with gap retains topology and shows degraded resolution confidence;
- API summary examples;
- CLI inventory command output;
- counts by App Layer, Functional Area, fact family, target role, classification, actionability, Resolution Health bucket, and topology;
- backend/CLI test names/results.

## E10 - Query Health Command Evidence

Status: pending

Record during P5.

Required evidence:

- command name and usage;
- suite fixture location and format;
- expected file list for the first suite, including `resolve.go`, `emit.go`, `diagnostics.go`, `compute.go`, `policy.go`, `graph-health-filters.ts`, Web graph layout code, layout optimizer code, and launcher runtime code;
- sample table or JSON output;
- hit@5/hit@10 by intent;
- noise reasons;
- tests for parsing, scoring, semantic output, and threshold failure.

## E11 - Semantic Command Surface Evidence

Status: pending

Record during P6.

| Command | Required evidence |
| --- | --- |
| `query` | node type, App Layer, Functional Area, Resolution Health, and related gap summary when available |
| `context` | symbol/node view with topology, resolution-health summary, and source/nearby gaps |
| `impact` | affected App Layers, affected Functional Areas, and resolution-health risks when supported |
| `detect-changes` | changed App Layers, changed Functional Areas, ResolutionGap changes, and resolution-health impact |

If a command cannot fully expose a semantic layer in this implementation, record the exact limitation and follow-up.

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
