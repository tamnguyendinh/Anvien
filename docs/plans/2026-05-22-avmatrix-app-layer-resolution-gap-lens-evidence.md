# AVmatrix App Layer Resolution Gap Lens Evidence Ledger

Date: 2026-05-22

Status: planned

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
- graph-based work starts from fresh analyze output; no stale graph fallback is a product path;
- accuracy is more important than minimizing graph size.
- analyze semantic enrichment must preserve both correctness and speed; neither may be traded away silently.
- aggregation or dedupe must not discard evidence only to make the graph smaller.

## E0A - Planning Codebase Audit Findings

Status: recorded

This audit was performed after the initial plan commit to make the plan match the current source tree. It is planning evidence only; Phase 0 must still rerun the required fresh baseline at implementation start.

Source facts that must shape implementation:

- `avmatrix analyze --force` on 2026-05-22 scanned `728` files, parsed `539`, reported `189` unsupported, `0` failed, and produced `22095` nodes and `54772` relationships at `.avmatrix/graph.json`.
- `internal/resolution/emit.go` emits unresolved references through `emitUnresolvedReference`, then attaches them to source nodes with `graphhealth.AppendDiagnosticToNode`.
- `internal/resolution/resolve.go` emits unresolved heritage, call, access, and type-reference diagnostics.
- `internal/graphhealth/diagnostics.go` stores diagnostics under `graphHealthDiagnostics`; `sameDiagnosticBucket` currently does not include `TargetText`, so different unresolved target texts can collapse into one bucket.
- `internal/graphhealth/policy.go` and `internal/graphhealth/diagnostics.go` already define diagnostic classification/actionability for builtin, standard-library, test-framework, external-library, in-repo unresolved, unclassified, non-actionable, review, and analyzer-gap.
- `internal/httpapi/graph.go` calls graph-health summary computation when building graph responses, and its graph-health report candidate limit is capped. Full ResolutionGap inventory must not rely on that capped report path.
- `internal/analyze/analyze.go` currently runs resolution, MRO, communities, processes, graph compaction, LadybugDB load, and graph snapshot writing in that order. Semantic enrichment must run before LadybugDB load and graph snapshot writing, and after any upstream signal it depends on.
- The target flow is resolution with raw unresolved fact capture, then MRO, communities, processes, semantic enrichment, graph compaction, LadybugDB load, and graph snapshot. Raw unresolved facts should keep `sourceNodeID`, `factFamily`, `targetText`, `filePath`, range or line, `resolutionSource`, and note.
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
| graph schema/snapshot | pending | pending |
| analyze semantic enrichment flow | pending | pending |
| semantic enrichment input indexes and complexity | pending | pending |
| LadybugDB export/load | pending | pending |
| unresolved call emission | pending | pending |
| unresolved access emission | pending | pending |
| unresolved type-reference emission | pending | pending |
| unresolved heritage emission | pending | pending |
| diagnostic attachment | pending | pending |
| graph-health summary/report | pending | pending |
| HTTP graph payload | pending | pending |
| generated Web contracts | pending | pending |
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

Status: pending

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

## E7 - Functional Area Evidence

Status: pending

Record during P2.

Required evidence:

- evaluated candidate signals;
- selected Functional Area rules and examples;
- rejected low-confidence signals and reasons;
- unknown counts and examples;
- schema/API/contract/command visibility;
- tests proving accepted rules and ambiguous unknown behavior.

## E8 - ResolutionGap Persistence Evidence

Status: pending

Record during P3.

Required evidence:

- persisted schema or graph representation;
- source-backed gap examples for unresolved call, access, type-reference, and heritage;
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
- Web unit command/output;
- Web e2e command/output;
- query-health benchmark output;
- resolution inventory output;
- `query`, `context`, `impact`, and `detect-changes` examples;
- AVmatrix detect-changes output for implementation commits;
- final commit hashes;
- residual risks and follow-up plan if any validation fails.
