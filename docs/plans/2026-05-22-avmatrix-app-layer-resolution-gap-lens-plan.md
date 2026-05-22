# AVmatrix App Layer Resolution Gap Lens Plan

Date: 2026-05-22

Status: in progress; Phase 0 closure audit complete; Phase 2 complete; Phase 2A proof-based CALLS/ACCESSES and source-site bridge slices complete; Phase 3 complete; Phase 4 complete; Phase 5 complete; Phase 6 remains next

Source discussion:

- [reports/problem/2026-05-22-avmatrix-app-layer-resolution-gap-discussion.md](../../reports/problem/2026-05-22-avmatrix-app-layer-resolution-gap-discussion.md)

Companion files:

- Benchmark ledger: [2026-05-22-avmatrix-app-layer-resolution-gap-lens-benchmark.md](2026-05-22-avmatrix-app-layer-resolution-gap-lens-benchmark.md)
- Evidence ledger: [2026-05-22-avmatrix-app-layer-resolution-gap-lens-evidence.md](2026-05-22-avmatrix-app-layer-resolution-gap-lens-evidence.md)

## Master rules

1. Use AVmatrix for codebase analysis and impact checks while working on implementation slices in this plan.
2. As each task is completed, update the corresponding checklist item immediately.
3. Run the full build gate before testing; include backend, contract, Web unit, and browser/e2e validation before closing the plan.
4. Record benchmark results as each benchmarkable task is completed. Benchmarkable means measured product/runtime performance, capacity, package/startup size, graph/DB throughput, query hit rate, command output inventory, graph inventory counts, source-site inventory counts, or resolved-edge accuracy; build/test/e2e timings are validation evidence unless the slice changes those systems.
5. Record evidence as each evidenced task is completed.
6. For doc-only commits, do not use AVmatrix.
7. After each completed implementation slice, commit the work, then continue until the full plan is complete.

## Rules of plan

1. Follow active workspace and repository instructions, including `AGENTS.md`; this plan records product work and validation, it does not replace repository rules.
2. Use AVmatrix according to active repository instructions for implementation slices; do not use AVmatrix for doc-only commits.
3. As each task is completed, update the corresponding checklist item immediately.
4. Run a full build before testing. For this plan the full build gate is `powershell -ExecutionPolicy Bypass -File avmatrix-launcher\build.ps1`.
5. Because this changes graph semantics, CLI command behavior, API contracts, and Web UI graph behavior, validation must include backend tests, contract checks, Web unit tests, and Web e2e tests.
6. Record benchmark results as each benchmarkable task is completed. Benchmarkable means measured product/runtime performance, capacity, package/startup size, graph/DB throughput, query hit rate, command output inventory, graph inventory counts, source-site inventory counts, or resolved-edge accuracy; build/test/e2e timings are validation evidence unless the slice changes those systems.
7. Record evidence as each evidenced task is completed.
8. After each completed implementation slice, commit the work, then continue until the full plan is complete.
9. Do not use product/runtime timeouts, elapsed-time budgets, delayed refresh, or automatic layout optimizer runs as a fix mechanism.
10. Do not implement stale graph fallback. Graph-based work starts from `avmatrix analyze --force`, and fresh analyze output is the source of truth.
11. Do not reduce graph evidence only to make the graph smaller. Aggregation or dedupe is allowed only when it preserves meaning, counts, samples, and traceability.

## Problem

The graph currently exposes primary node labels such as `Function`, `Method`, `Struct`, `Interface`, `Variable`, `Property`, `File`, `Folder`, `Package`, `Process`, `Community`, and `Section`. Those labels describe the shape of a symbol, but they do not answer which product surface owns the node.

A `Function` can belong to backend resolution, frontend React code, API handlers, launcher/runtime code, test helpers, generated contracts, shared packages, or documentation tooling. When unresolved references appear, reading the graph data does not reliably answer:

- whether the source is Backend, API, Frontend, Shared, Test, Docs, Config, Generated, or a mixed category;
- which functional area owns the source node;
- whether the unresolved target looks like a callable, member, type, external symbol, builtin, test helper, or unknown target;
- whether topology conclusions such as `no_incoming`, `detached_component`, or `true_isolated` have degraded resolution confidence;
- whether the issue is an in-repo analyzer gap, an external unresolved symbol, a non-actionable builtin/stdlib/test reference, or still unclassified.

The Web UI only renders data provided by the graph/API. Therefore the fix must make the graph/API answer these questions first, then let the Web UI filter, group, color, and explain the result.

An independent CALLS/ACCESSES benchmark found a deeper graph-accuracy requirement: AVmatrix currently has much cleaner ACCESSES output than GitNexus and fewer CALLS duplicates, but a resolved edge is still wrong if it was produced from an unproven name match. For example, a bare Go call such as `stop()` must not be connected to an unrelated TypeScript method named `stop` unless the resolver proves that binding. The correct graph contract is proof-based: source sites must be inventoried, but only proven relationships may become resolved `CALLS` or `ACCESSES` edges.

## Scope Boundary

Implementation may touch:

- graph node and relationship schema;
- analyzer/graph generation metadata;
- reference-resolution proof metadata for `CALLS` and `ACCESSES`;
- call/access source-site inventory records;
- analyze pipeline phase ordering and semantic enrichment flow;
- resolution diagnostic emission and graph-health diagnostic attachment;
- persisted graph snapshot shape under `.avmatrix`;
- LadybugDB schema/export/load surfaces used by graph-backed query and Cypher;
- graph-health and resolution-health summaries;
- HTTP graph APIs and generated Web contracts;
- CLI/query/MCP command output for `query`, `context`, `impact`, `detect-changes`, API-specific MCP tools, and any new query-health or resolution-inventory command;
- Web graph state, filters, graph detail panels, dashboards, generated types, and layout placement;
- backend, contract, Web unit, and Web e2e tests.

Out of scope unless a later phase explicitly reopens it:

- treating unresolved references as confirmed dead code;
- merging `unknown_connectivity` back into `unresolved_reference`;
- synthesizing fake resolved in-repo target nodes or fake semantic edges for unresolved targets;
- emitting resolved `CALLS` or `ACCESSES` edges from unproven global name fallback, cross-language name matching, or coarse file-level evidence;
- using UI-only virtual data as the source of truth for graph semantics;
- allowing overlapping primary App Layer labels on one node;
- using low-confidence Functional Area guessing just to reduce `unknown` counts;
- stale graph compatibility as a product path;
- product/runtime timeout behavior, delayed refresh, or automatic layout optimizer runs.

## Design Decisions

App Layer is a primary, non-overlapping classification. A node must have one primary App Layer category. If a node belongs to a mixed concern, that mixed concern becomes its own category such as `frontend_test`, `api_contract`, `frontend_api_client`, or `api_shared_contract`; it must not be represented by overlapping primary labels.

API is a first-class App Layer. API handlers, API contracts, shared API schemas, generated API contracts, and frontend API clients may need separate rings/categories if that is clearer than forcing them into Backend or Frontend.

Functional Area is a second semantic layer under App Layer. It may include areas such as `resolution`, `graph_health`, `query`, `mcp`, `web_graph_ui`, `layout`, `contracts`, `providers`, `runtime`, `analyzer`, `session`, `launcher`, `cli`, and `reporting`, but only rules with strong evidence should be implemented. Ambiguous nodes should remain unknown rather than guessed.

ResolutionGap or UnresolvedSymbol must be persisted into graph output. It must not be only a virtual API/UI lens. Persisted gaps must retain enough source data for CLI, query, context, impact, detect-changes, Web UI, and future MCP consumers to inspect the same truth.

Resolution Health is separate from Topology Health. A node can remain `connected`, `no_incoming`, `no_outgoing`, `detached_component`, or `true_isolated` while also carrying resolution gaps and degraded resolution confidence.

Web layout should use App Layer as the macro placement ring and existing node type/filter kind as the micro island inside that ring. The goal is readable separation: Backend/API/Frontend/Shared/Test/Docs/Config/Generated/Mixed rings when present, with same-color/type islands kept together instead of mixed across a ring. API must be treated as the bridge between Backend and Frontend; contract rings should be placed near API when present.

Query-related commands should understand the new semantic layers. `analyze` remains the base graph-producing command; child commands such as `query`, `context`, `impact`, `detect-changes`, and a new query-health/inventory command should surface App Layer, Functional Area, and ResolutionGap meaning when the graph provides it.

Existing graph-health diagnostic classification and actionability are product code that should be reused and extended where accurate. The plan must not create a second incompatible classification path for builtin, standard-library, test-framework, external-library, in-repo analyzer-gap, or unclassified cases.

Resolved graph edges must be proof-based. `CALLS` is only for a call site whose target has been proven by local scope binding, import/module binding, receiver/type binding, or another explicitly recorded proof. `ACCESSES` is only for field/property read/write semantics where the target is proven to be a property/field-like node. A selector such as `config.NewConfig()` is not an `ACCESSES` edge when `NewConfig` is a function, and a method call must be represented as a call rather than as a property access unless a separate property read is proven.

Every syntactic call/access source site must be inventoried even when it cannot emit a resolved edge. Source-site records need a stable ID, source node, file/range, source text or target text, fact family, status, proof kind, target role, and diagnostic/actionability fields when available. Valid statuses include at least `resolved`, `unresolved_local_binding`, `unresolved_external`, `ambiguous`, `dynamic`, `unsupported_syntax`, and `unknown`. A source site with no proof must become a persisted unresolved/ambiguous/external fact instead of a guessed edge.

The graph-accuracy target is: resolved edge false positives are zero in the golden corpus, source-site inventory misses are zero in the golden corpus, unresolved/ambiguous classification is explicit, and silent missing facts are not allowed. More edges is not a success metric unless those edges are proven.

The deterministic initial graph placement is separate from the manual layout optimizer. The Web graph already applies an initial filter-based clustered placement during graph conversion, while the optimizer button invokes layout work manually. This plan changes the deterministic initial placement into App Layer rings and type islands; it must not add automatic optimizer execution after render, load, filter changes, or refresh.

Semantic enrichment is a graph/analyze pipeline concern. The implementation must choose and test one flow that produces the most accurate graph facts while preserving analyzer speed. Current source inspection shows analyze resolves references, applies MRO, communities, and processes, compacts the graph, loads LadybugDB, then writes the graph snapshot. App Layer, Functional Area, and ResolutionGap enrichment must run before LadybugDB load and graph snapshot, and after every upstream signal it depends on is available. If Functional Area depends on process/community membership, the enrichment phase belongs after those phases and before compact/load/snapshot. If ResolutionGap facts are produced during resolution, their raw target identity must be captured then and finalized later during enrichment.

The target analyze flow for this plan is:

```text
scan/parse
-> build graph
-> cross-file binding
-> resolution
   -> capture call/access source sites with sourceSiteID, sourceNodeID, factFamily, targetText, filePath, range, status, proofKind, and evidence
   -> capture raw unresolved facts with sourceNodeID, factFamily, targetText, filePath, range, resolutionSource, and note
-> MRO
-> communities
-> processes
-> semantic enrichment
   -> App Layer
   -> Functional Area
   -> source-site status/proof summaries
   -> ResolutionGap / UnresolvedSymbol
   -> Resolution Health inventory
-> graph.Compact()
-> LadybugDB load
-> graph.json snapshot
```

The enrichment phase must not rescan files or reparse ASTs. It should build reusable indexes once, such as `nodeID -> node`, `filePath -> App Layer`, `nodeID -> process/community`, `sourceSiteID -> call/access site`, and `sourceNodeID -> raw gaps`, then run near `O(nodes + relationships + rawSites + rawGaps)` work. App Layer should primarily use cached path/package rules, Functional Area should use only accepted high-confidence signals, source-site status should preserve resolver proof, and ResolutionGap bucketing must include `targetText` so repeated unresolved facts do not lose identity.

## Acceptance Criteria

- Graph nodes expose a persisted primary App Layer category with no overlapping primary labels.
- API and API-related mixed categories are first-class, not hidden under Backend or Frontend.
- Functional Area metadata is persisted only where evidence is accurate enough; ambiguous nodes remain unknown.
- Source-backed unresolved references are represented as persisted ResolutionGap/UnresolvedSymbol graph entities or equivalent persisted graph records, not only diagnostic text.
- `CALLS` and `ACCESSES` edges are emitted only when the resolver records an explicit proof; unproven source sites are persisted as unresolved, ambiguous, external, dynamic, unsupported, or unknown source-site facts.
- Every call/access source site has an inventory record with source node, file/range, target text, fact family, status, proof kind, target role when known, and evidence.
- `ACCESSES` resolved targets are property/field-like targets only; functions, methods, consts, structs, variables, imports, and coarse file-level matches are not emitted as `ACCESSES` unless a field/property access proof exists.
- Bare calls and cross-language same-name matches do not use global fallback to produce resolved `CALLS` edges. If local binding, import binding, receiver binding, or other accepted proof is absent, the site remains unresolved or ambiguous.
- Golden accuracy tests prove expected edges exist, known false edges do not exist, unresolved source sites are not lost, duplicate resolved edges are controlled, and source-site inventory has no silent missing facts.
- Repeated unresolved references with different target text are not collapsed into a bucket that loses target identity.
- Fine-grained gap relationships or typed gap metadata preserve call, access, type-reference, heritage, external, builtin, test, analyzer-gap, and unknown distinctions where evidence supports them.
- Topology Health and Resolution Health remain separate in graph/API/CLI/Web output.
- CLI/query command surfaces can report App Layer, Functional Area, ResolutionGap, and resolution-health information from persisted graph data.
- A query-health benchmark command exists and reports hit@5/hit@10, expected files/symbols, actual results, noise reason, and pass/fail.
- Query-health captures the current noisy baseline and then validates the improved `query` implementation against expected source files and symbols.
- Web UI exposes App Layer filters, Resolution Health filters, and a multi-ring layout where App Layer controls macro placement and node type/gap kind controls islands inside rings.
- The Web UI does not invent App Layer, Functional Area, or ResolutionGap truth on the client.
- Layout optimizer remains manual-only and is not auto-run after render.
- Full build, backend tests, contract checks, Web unit tests, and Web e2e tests pass before closure.
- Benchmark and evidence ledgers contain baseline and final inventories for App Layer, Functional Area, ResolutionGap, Resolution Health, query benchmark, CLI semantic output, and Web rings/filters.
- Missing App Layer, Functional Area, source-site proof/status, ResolutionGap, or Resolution Health metadata in loaded graph data is treated as stale/incomplete graph evidence when that metadata is required by the active schema, not as a reason to guess at API/UI load time.
- If ResolutionGap data is aggregated or deduped, the aggregate keeps exact counts and representative source evidence without capping away meaning.
- Resolution inventory and Resolution Health APIs/commands expose full counts and must not rely on capped graph-health triage report candidates.
- Analyze produces one consistent semantic graph across in-memory graph, LadybugDB export, graph JSON, HTTP API, MCP tools, and Web contracts.
- Semantic enrichment has recorded runtime/memory/graph-size benchmarks and does not add avoidable rescans or duplicate graph traversals.

## Current Code Facts To Account For

The following facts came from planning and source inspection. Facts not already superseded by recorded Phase 1/Phase 2 evidence must be re-verified before the implementation slice that depends on them:

- unresolved references are currently emitted from `internal/resolution/emit.go` and attached to source nodes with `graphhealth.AppendDiagnosticToNode`;
- `internal/graphhealth/diagnostics.go` currently aggregates diagnostic buckets without `TargetText` in the bucket key, so different unresolved targets can collapse into one bucket and keep only the first target text;
- graph-health diagnostic classification/actionability already exists in `internal/graphhealth/policy.go` and `internal/graphhealth/diagnostics.go`;
- graph-health summaries are computed by HTTP graph paths such as `internal/httpapi/graph.go`, so persisted ResolutionGap/App Layer semantics must be produced earlier than API response shaping;
- analyze currently runs resolution, MRO, communities, processes, graph compaction, LadybugDB load, and graph snapshot writing in `internal/analyze/analyze.go`;
- LadybugDB data is exported from the in-memory graph through `internal/lbugload`, so new persisted semantic fields or entities must be exported there when query/Cypher consumers need them;
- API-specific MCP surfaces such as `route_map`, `shape_check`, and `api_impact` already exist and must be considered because API is a first-class App Layer;
- `/api/graph/report` has a capped candidate limit and is not sufficient as the full ResolutionGap inventory source;
- `query` currently ranks process matches with simple contains scoring in `internal/mcp/tools.go`, and definition matching is limited enough that function/method-centric intents can miss the expected files;
- Web graph filters currently consume graph/API metadata, generated contracts, and client-side filter state, but App Layer and Resolution Health filters do not exist yet;
- Web graph filter state is managed through `avmatrix-web/src/hooks/app-state/graph.tsx` in addition to panel and adapter files;
- Web graph conversion currently applies deterministic filter-based clustered layout in `avmatrix-web/src/lib/graph-adapter.ts`, and the manual optimizer button calls layout work through `avmatrix-web/src/hooks/useSigma.ts`;
- generated Web contracts in `internal/contracts/web_ui.go` are the boundary that should prevent the UI from relying on ad hoc shape guesses.
- `internal/resolution/resolve.go` currently reaches `resolveGlobalCallName` as a fallback for constructor, member, and free-call resolution, then emits `CALLS` through `emitReference` when it returns a target, including confidence `0.5` paths. Phase 2A must treat low-confidence simple-name/global matches as source-site status unless an accepted proof kind exists.
- `internal/resolution/indexes.go` defines the resolver proof surfaces that Phase 2A must audit: `resolveGlobalCallName`, `resolveSameFileName`, `resolveGoSamePackageFunction`, `resolveMember`, `resolveImportedMember`, `callableLabels`, `propertyLabels`, and dispatch-owner indexes.
- `internal/resolution/emit.go` currently dedupes semantic edges with `semanticEdgeKey` using source, target, relationship type, and limited call/access details; it does not preserve a `sourceSiteID` on the resolved relationship, so Phase 2A must keep exact source-site occurrences or exact occurrence counts before dedupe hides them.
- `internal/graph/types.go` `Relationship` currently has evidence, confidence, and resolution source fields, but no source-site status, proof kind, sourceSiteID, or source range fields. Phase 2A must choose an explicit persisted schema: extended relationship metadata, a persisted SourceSite entity/record, or both where each serves a clear consumer.
- `internal/scopeir/facts.go` defines `CallSiteFact` and `AccessFact`; provider reference collectors, provider parity tests, and generated contract tests are implementation surfaces for making source-site inventory complete across languages.
- `internal/resolution/indexes.go` `propertyLabels()` and `internal/contracts/web_ui.go` provider fact coverage currently allow `Property`, `Variable`, `Const`, and `Static` as ACCESSES-like targets. Phase 2A must align this with the strict `ACCESSES = proven property/field read/write` contract or split non-property uses into a separate relation/fact role before emitting.
- `internal/graphaccuracy/graphaccuracy.go` currently validates Go definitions/imports and a direct CALLS subset; it does not yet validate ACCESSES precision, source-site inventory completeness, false resolved edges, or the cross-language `stop()` false-positive class.
- independent benchmark evidence shows AVmatrix ACCESSES is much cleaner than GitNexus because all sampled AVmatrix ACCESSES targets are Property-like while GitNexus over-emits function/method/const targets and duplicates heavily; this should be preserved with a strict ACCESSES contract rather than relaxed.
- independent benchmark evidence shows AVmatrix CALLS has zero duplicate pairs in the sampled comparison but still has a real false resolved edge where Go `main.stop()` is connected to TypeScript `SSEListener.stop`; this requires removing or gating dangerous simple-name/global fallback.
- GitNexus emits more CALLS but includes coarse `File -> Function` edges; AVmatrix should keep symbol-level source precision and should not use file-level source edges as proof of a resolved call.

## Checklist Item Standard

Each checkbox below is a concrete unit of work with a visible output in code, generated contracts, tests, benchmark data, or evidence ledgers. Constraints and cautions may appear inside an item, but only as part of doing that concrete work correctly.

## Phase 0 - Baseline, Discussion Coverage, And Source Trace

- [x] [P0-A] Read the full discussion report and write a coverage table in the evidence ledger that maps every major decision to this plan: node type insufficiency, BE/API/FE/App Layer rings, non-overlapping mixed categories, Functional Area accuracy, proof-based CALLS/ACCESSES, source-site inventory, persisted ResolutionGap, fine-grained gap relations, Resolution Health, query-health command, semantic command output, multi-ring Web layout, no stale fallback, and no timeout/auto optimizer behavior.

- [x] [P0-B] Run `avmatrix analyze --force` before the next graph-based implementation slice that depends on baseline counts, then record scanned/parsed/unsupported/failed file counts, graph node count, graph relationship count, counted semantic relationship count, execution-flow count, `unknown_connectivity` count, and graph timestamp/hash in the benchmark ledger. Compare the fresh baseline against the discussion observations of about `22010` nodes, `26906` counted semantic relationships, `0` `unknown_connectivity`, `51232` unresolved occurrences, and `8880` unresolved buckets without assuming those old numbers are still exact.

- [x] [P0-C] Trace the current unresolved-reference pipeline from resolution source facts to graph/API output. Record the exact files/symbols that create unresolved call, access, type-reference, and heritage diagnostics; record where target text, source node, fact family, classification, actionability, and source location are currently kept or lost, including whether `sameDiagnosticBucket` collapses distinct target text in the same source/fact/file/note bucket.

- [x] [P0-D] Measure the current unresolved-reference inventory before changing semantics using full graph data or a temporary audit script that preserves full counts when existing APIs/commands are capped. Record bucket count, occurrence count, fact family counts, diagnostic classification/actionability counts if present, top target texts, source node labels, source path buckets, and whether each source already has a topology status.

- [x] [P0-E] Measure a provisional path/package-derived App Layer inventory without changing product behavior. Seed the audit with the discussion examples `avmatrix-web/src/**`, `avmatrix-web/test/**`, `avmatrix-web/e2e/**`, `internal/**`, `cmd/**`, `internal/httpapi/**`, `avmatrix-web/src/services/backend-client.ts`, `avmatrix-launcher/**`, `contracts/**`, `internal/contracts/**`, `cmd/generate-web-contracts/**`, `docs/**`, `reports/**`, `*.md`, `*_test.go`, `test/fixtures/**`, config files, package files, and build scripts. The evidence must show candidate counts for backend, api, frontend, cli_launcher, shared_contract, api_contract, api_shared_contract, frontend_api_client, backend_test, frontend_test, api_test, generated_contract, docs, config, generated, mixed, and unknown, plus examples explaining every uncertain bucket.

- [x] [P0-F] Audit current query behavior with fixed benchmark intents before adding the new command. At minimum test intents for unresolved reference diagnostic generation, graph health unknown-connectivity separation, App Layer/resolution-gap layout, API contract surfaces, frontend graph filter surfaces, and runtime reset hidden-terminal behavior; record expected files/symbols, actual top results, hit/miss, and noise reason. The audit must include the current `internal/mcp/tools.go` process scoring and definition matching behavior so later query work fixes the actual retrieval path rather than only adding a benchmark wrapper.

- [x] [P0-G] Locate the implementation surfaces for graph schema, graph generation, resolution diagnostics, graph health, HTTP graph APIs, generated Web contracts, CLI query/context/impact/detect-changes, Web graph filters, Web graph detail panels, and Web layout. Record exact file paths in evidence, including `internal/resolution/emit.go`, `internal/resolution/resolve.go`, `internal/resolution/indexes.go`, `internal/scopeir/facts.go`, `internal/graph/types.go`, `internal/graphaccuracy/graphaccuracy.go`, `internal/graphhealth/diagnostics.go`, `internal/graphhealth/policy.go`, `internal/httpapi/graph.go`, `internal/contracts/web_ui.go`, `internal/mcp/tools.go`, `internal/cli/tool_command.go`, `internal/providers/*/references.go`, `internal/providers/*/extract_test.go`, `internal/providers/provider_parity_test.go`, `avmatrix-web/src/lib/graph-adapter.ts`, `avmatrix-web/src/lib/graph-health-filters.ts`, `avmatrix-web/src/components/FileTreePanel.tsx`, `avmatrix-web/src/components/GraphCanvas.tsx`, and `avmatrix-web/src/hooks/useSigma.ts`.

- [x] [P0-H] Trace the current analyze persistence flow and choose the semantic enrichment insertion point. Record the exact order around resolution, MRO, communities, processes, graph compact, LadybugDB load, and graph snapshot; define which App Layer, Functional Area, and ResolutionGap inputs are available at that point; record the performance constraints for keeping enrichment accurate and fast.

- [x] [P0-I] Specify the semantic enrichment input indexes and complexity budget before implementation. The design must use already-produced graph/resolution/process/community facts, build reusable maps once, avoid file rescans and AST reparses, and target near `O(nodes + relationships + rawSites + rawGaps)` behavior when source-site inventory is present.

## Phase 1 - App Layer Taxonomy And Persistence

- [x] [P1-A] Define the persisted App Layer category registry as a primary one-of field, not overlapping tags. The category list must include normal categories and mixed categories where needed: backend, api, frontend, cli_launcher, shared_contract, api_contract, api_shared_contract, frontend_api_client, backend_test, frontend_test, api_test, generated_contract, docs, config, generated, mixed, unknown, and any additional categories proven necessary by P0 evidence.

- [x] [P1-B] Define the source evidence required for every App Layer category and the explicit `unknown` assignment rule for insufficient evidence. Use high-confidence signals such as path, package/module, generated-contract location, API route/handler location, frontend source roots, test naming, docs/config paths, and launcher/runtime ownership; start from the P0-E path seed list and refine it with source inspection.

- [x] [P1-C] Implement the backend App Layer classifier as a dedicated semantic classification surface with unit coverage for category evidence rules, mixed categories, and `unknown` assignment. Keep this classifier reusable by the analyze enrichment phase, inventory code, and tests.

- [x] [P1-C2] Wire the App Layer classifier into the analyze semantic enrichment phase and persist the result in graph output. Place it according to P0-H so App Layer metadata is present before graph compact, LadybugDB load, graph snapshot writing, HTTP API consumption, and MCP graph reads. Record enrichment latency, memory, and graph-size impact in the benchmark ledger.

- [x] [P1-D] Make API first-class by classifying server handlers, API graph endpoints, API contract/schema code, generated API contract files, and frontend API clients into separate categories when evidence supports that separation. If a surface is both API and shared contract, use a mixed category such as `api_shared_contract` rather than overlapping labels.

- [x] [P1-E] Update graph schema, snapshot serialization, LadybugDB export/load schema, HTTP graph payloads, generated Web contracts, and any contract tests so App Layer values are stable public fields. A freshly analyzed but ambiguous node may use `unknown`; graph data that lacks the new metadata entirely must be treated as stale/incomplete schema evidence and must not trigger load-time classification heuristics.

- [x] [P1-F] Add tests for simple and mixed App Layer examples. Coverage must prove one primary App Layer per node, correct API classification, correct mixed category classification, correct unknown handling for ambiguous input, and no accidental multi-label primary classification.

- [x] [P1-G] Record before/after App Layer counts, examples by category, changed schema fields, generated contract output, and test evidence in the benchmark and evidence ledgers.

- [x] [P1-H] Define user-facing and machine-facing names for App Layer, API Layer, API Contract, Frontend API Client, Resolution Gap, Unresolved Symbol, Analyzer Gap, External Reference, and Non-actionable Reference. Record enum keys, display labels, CLI labels, and Web labels so API, CLI, and UI do not drift.

## Phase 2 - Functional Area Accuracy Gate

- [x] [P2-A] Evaluate candidate Functional Area signals from P0 evidence: path prefix, package/module name, process membership, community detection, import/call neighborhood, explicit config, and any AI-assisted labeling only if it is reproducible and verifiable. Record accepted and rejected signals with examples instead of choosing a weak signal because it is easy to code.

- [x] [P2-B] Define a Functional Area registry for high-confidence areas only. Initial candidates may include resolution, graph_health, query, mcp, web_graph_ui, layout, contracts, providers, runtime, analyzer, session, launcher, cli, reporting, and unknown; each accepted area must have exact evidence rules.

- [x] [P2-C] Persist Functional Area metadata in graph output only for nodes that meet accepted rules. Ambiguous nodes must stay `unknown` or equivalent rather than being forced into a functional group.

- [x] [P2-C2] Extend the analyze semantic enrichment phase with Functional Area assignment after the required signals from P0-H are available. If accepted rules depend on process or community membership, run this assignment after those phases and before graph compact, LadybugDB load, and graph snapshot writing.

- [x] [P2-D] Expose Functional Area through API payloads, generated contracts, CLI surfaces, and Web detail/filter data where available. Consumers must distinguish "unknown because not enough evidence" from "missing field because graph was not freshly analyzed".

- [x] [P2-E] Add tests for accepted Functional Area rules, rejected low-confidence rules, ambiguous nodes, and command/API/Web contract visibility.

- [x] [P2-F] Record selected rules, rejected rules, counts, unknown counts, example nodes, and test evidence in the evidence and benchmark ledgers.

## Phase 2A - Proof-Based CALLS/ACCESSES And Source-Site Inventory

- [x] [P2A-A] Define the resolved-edge contract for `CALLS` and `ACCESSES` as a code-facing and test-facing specification. The contract must state accepted proof kinds for calls, accepted proof kinds for field/property accesses, rejected proof sources such as global simple-name fallback, low-confidence simple-name matches, and cross-language same-name fallback, and the exact status values used when a source site cannot produce a proven edge. Confidence may be recorded as evidence, but confidence alone is not proof.

- [x] [P2A-B] Trace the current CALLS/ACCESSES resolver paths and record where resolved edges are emitted, where unresolved facts are emitted, where target role is known, where duplicate suppression occurs, and where false edges can be introduced. The trace must cover `internal/resolution/resolve.go`, `internal/resolution/emit.go`, `internal/resolution/indexes.go`, `internal/scopeir/facts.go`, `internal/graph/types.go`, `internal/providers/*/references.go`, `internal/providers/*/extract_test.go`, `internal/providers/provider_parity_test.go`, `internal/contracts/web_ui.go`, and `internal/graphaccuracy/graphaccuracy.go`; it must specifically cover the `stop()` false-positive class, selector/import cases such as `config.NewConfig()`, receiver method calls, property reads/writes, local function variables, closures, imports, builtins, external packages, and TypeScript/React owner attribution.

- [x] [P2A-C] Add persisted call/access source-site inventory before resolved-edge emission is finalized. Each record must keep sourceSiteID, source node ID, source App Layer, source Functional Area when known, file path, range or line/column, source text or target text, fact family, target role when known, status, proof kind, resolution source, evidence note, and any linked resolved target ID when proof exists. The implementation must choose and document whether this is persisted as extended `graph.Relationship` metadata, SourceSite graph entities/records, or both, then propagate the chosen schema through graph JSON, LadybugDB, HTTP/API contracts, generated Web contracts, and command/query consumers that need it.

- [x] [P2A-C1] Persist source-site and proof metadata on resolved relationships and source-backed unresolved diagnostics as the first source-site inventory slice. Resolved `CALLS`, `ACCESSES`, `USES` type references, and `INHERITS` compatibility edges now carry `sourceSiteId`, `sourceSiteIds`, `sourceSiteCount`, `sourceSiteStatus`, `proofKind`, `targetRole`, `targetText`, file path, and range metadata. Unresolved source-backed diagnostics now carry the same identity/status/proof fields so Phase 3 can consume precise raw facts instead of reconstructing gaps from summary text. This slice intentionally does not create fake in-repo target nodes for unresolved targets.

- [x] [P2A-D] Gate resolved edge emission on proof. `CALLS` must be emitted only for proven call bindings; `ACCESSES` must be emitted only for proven property/field access targets. Bare calls without local/import/receiver proof, function variables without assignment/capture proof, low-confidence `resolveGlobalCallName` results, cross-language same-name matches, import selector references that resolve to functions, and coarse file-level matches must remain source-site facts with unresolved/ambiguous/external/dynamic/unsupported status instead of becoming resolved edges. Accepted same-file, same-package, import, receiver, and member matches must carry named proof kinds so tests can assert why they became resolved edges.

- [x] [P2A-D1] Remove low-confidence global/simple-name fallback as proof for resolved `CALLS`. `resolveCall` now uses scoped local/import binding before same-file and same-package proof, treats `resolveGlobalCallName` matches as unresolved source-backed diagnostics, and includes golden resolver tests proving bare Go `stop()` does not call TypeScript `SSEListener.stop`. The slice also adds PHP `use function`/`use const` import facts so PHP function calls that have import evidence stay proof-backed instead of relying on global fallback.

- [x] [P2A-D2] Gate file-level call sources out of resolved `CALLS`. A call site whose source owner is only the `File` node is now preserved as a source-site diagnostic with `sourceSiteStatus=unsupported_syntax` and note `call source is file-level; resolved edge not emitted`, instead of becoming a coarse `File -> Function` topology edge. This keeps source-site inventory without claiming a symbol-level caller that the analyzer did not prove.

- [x] [P2A-E] Connect source-site inventory to ResolutionGap/UnresolvedSymbol inputs so Phase 3 consumes source-backed unresolved/ambiguous/external call/access sites rather than reconstructing gaps from already-aggregated diagnostics. The implementation must preserve existing diagnostic summaries while making source-site records the more precise source of truth for call/access resolution health.

- [x] [P2A-E1] Add the reusable source-backed ResolutionGap input model in `internal/graphhealth`. `SourceBackedResolutionGapInputs` reads persisted unresolved-reference diagnostics with `sourceSiteId` directly from node diagnostic properties and preserves source node, App Layer, Functional Area, fact family, target text, target role, source-site status, proof kind, classification, actionability, resolution source, file/range, count, and note. `SourceBackedCallAccessResolutionGapInputs` provides the Phase 3 call/access-specific input without reading graph-health summaries or creating fake target nodes/edges. Pre-commit detect-changes after staging the new files reported changed_count `56`, changed_files `5`, affected_count `0`, risk_level `low`.

- [x] [P2A-F] Add a golden accuracy corpus for CALLS and ACCESSES with positive and negative expectations. Required cases include a proven property access, a proven method/function call, a selector call that must not become ACCESSES, the Go `stop()` local-binding false-positive class that must not call `SSEListener.stop`, local function variables, closures, imports, external/builtin calls, TypeScript/React method ownership, provider-level `CallSiteFact` and `AccessFact` preservation, duplicate-edge prevention with exact source-site occurrence counts, and at least one coarse File-source edge case that must not be accepted as symbol-level proof.

- [x] [P2A-G] Add accuracy metrics and benchmark output for source-site and edge correctness. Record raw call sites, raw access sites, resolved `CALLS`, resolved `ACCESSES`, low-confidence/global-fallback source sites, unresolved/ambiguous/external/dynamic/unsupported source sites, false resolved edges found by golden tests, silent missing source sites found by golden tests, duplicate resolved edges, source sites merged by edge dedupe without occurrence evidence, max duplicate, ACCESSES target label distribution, and non-property ACCESSES targets. The target for golden tests is false resolved edges `0`, silent missing source sites `0`, source sites hidden by dedupe `0`, and non-property ACCESSES targets `0` unless they have been split into a separate non-ACCESSES relation/fact role.

- [x] [P2A-H] Update graph snapshot, LadybugDB export/load, HTTP/API contracts, generated Web contracts, CLI/query inventory surfaces, and `internal/contracts/web_ui.go` provider fact coverage where needed so source-site status and proof metadata are available to later ResolutionGap, Resolution Health, query/context/impact, and Web UI work. This item must also update or split the current ACCESSES target-label contract so `Variable`, `Const`, and `Static` are not treated as proven property/field ACCESSES unless a separate relation/fact role explicitly explains them. If a surface cannot expose the full inventory in this slice, record the exact limitation and the persisted source of truth it can read later.

- [x] [P2A-H1] Propagate the source-site/proof relationship and diagnostic fields through graph JSON, LadybugDB relation CSV/export/load schema, fallback relationship inserts, Go Web contract source, generated Web TypeScript contracts, and backend tests. `propertyLabels()` now only accepts `Property` for resolved `ACCESSES`; the access-candidate audit separates rejected `Variable`/`Const`/`Static` selector targets into `non_property_target` instead of treating them as proven field/property access.

- [x] [P2A-I] Extend `internal/graphaccuracy` or add a dedicated CLI/report command for proof-based source-site accuracy. The command/report must read the current graph or accuracy fixtures and emit the B5A metrics: source-site inventory counts, resolved edge counts, unresolved/ambiguous/external/dynamic/unsupported status counts, low-confidence/global-fallback counts, ACCESSES target label distribution, duplicate/merged source-site counts, golden false-positive counts, and silent missing source-site counts.

- [x] [P2A-I1] Add the graph-inventory implementation of `avmatrix source-site-accuracy` to the packaged CLI. The command reads `.avmatrix/graph.json` or a supplied `--graph`, writes JSON with `--out` or stdout with `--json`, and reports source-site occurrence counts, resolved `CALLS`/`ACCESSES` counts, unresolved diagnostic status/proof/fact-family counts, low-confidence fallback counts, ACCESSES target labels, duplicate and merged source-site counts, graph-policy violation candidates, and explicit golden-validation availability. Graph inventory mode records golden validation as disabled unless a fixture is supplied.

- [x] [P2A-I2] Add fixture-backed golden validation to `avmatrix source-site-accuracy` with `--golden`. The fixture mode reads expected source-site IDs and known-false resolved edges, then reports expected and matched source-site counts, silent missing source-site counts, expected false-edge counts, false resolved edges found in the graph, and capped examples for both missing sites and false edges. This keeps source-site inventory metrics and golden proof checks in the same packaged command without turning fixture expectations into product graph facts.

- [x] [P2A-J] Add backend tests, contract tests, command/API visibility tests, and benchmark/evidence ledger entries for the proof-based CALLS/ACCESSES slice. Run the full build before tests, include focused resolver/source-site tests and wider graph tests, then record validation and AVmatrix detect-changes before committing the slice.

- [x] [P2A-J1] Add backend and CLI command tests for the graph-inventory source-site accuracy command. The tests cover JSON command output, source-site relationship and diagnostic counts, merged source-site occurrence evidence, low-confidence fallback diagnostics, strict ACCESSES target labeling, duplicate source-target pairs, and graph-policy violation candidate reporting. Validation for this slice uses full build first, then focused `internal/graphaccuracy` and `internal/cli` tests plus wider `internal/...` and `cmd/...` tests.

- [x] [P2A-J2] Add resolver test coverage for the file-source CALLS gate and re-run the source-site accuracy report on a fresh graph. The report must show `coarseFileCallEdges=0`, `falseResolvedEdgeCandidates=0`, no resolved edges without proof, no resolved edges without sourceSiteID, no low-confidence fallback resolved edges, and no non-property ACCESSES targets.

- [x] [P2A-J3] Add the resolver golden corpus validation slice. `TestProofBasedCallAccessGoldenCorpus` covers twelve controlled call/access source sites: five proven `CALLS` occurrences, one proven `ACCESSES` occurrence, six unresolved diagnostics, duplicate helper calls merged with exact `sourceSiteCount=2`, builtin and external diagnostics, a local function variable call that must not resolve, the `stop()` low-confidence false-positive class, a function selector that must not become `ACCESSES`, and a file-level call that must remain `unsupported_syntax`. Full build ran before tests; focused and wider backend suites passed.

- [x] [P2A-J4] Add backend and CLI tests for source-site accuracy golden fixture mode. `TestRunSourceSiteAccuracyValidatesGoldenFixture` covers expected source-site matching, silent missing source-site reporting, known-false edge detection, and summary output. `TestSourceSiteAccuracyCommandOutputsJSON` covers packaged CLI JSON visibility for `--golden`. Full build ran before tests; focused `internal/graphaccuracy` and `internal/cli` tests plus wider `internal/...` and `cmd/...` tests passed. Fresh analyze plus `detect-changes` ran before commit and reported changed_count `74`, changed_files `7`, affected_count `12`, risk_level `high`; this is expected blast-radius context for extending a packaged CLI command and graphaccuracy report flow.

## Phase 3 - Persisted ResolutionGap And UnresolvedSymbol Model

- [x] [P3-A] Define the persisted ResolutionGap/UnresolvedSymbol data model using Phase 2A source-site inventory where available. `internal/graphhealth.ResolutionGapInput` now produces persisted `ResolutionGap` graph nodes keyed by `ResolutionGap:<sourceSiteID>` when source-site identity exists, with fallback identity including source node, fact family, target text, target role, status, proof kind, classification, and actionability. The node preserves sourceSiteID, source node ID/label, source App Layer, source Functional Area when known, fact family, target text, target role, source-site status, proof kind, classification, actionability, resolution source, source, file path/hash/range, occurrence count, and note. It does not claim the unresolved target is resolved.

- [x] [P3-B] Add or extend raw unresolved fact storage from resolution before diagnostic summary aggregation. Phase 3 now consumes the Phase 2A source-backed unresolved diagnostic/source-site records directly through `SourceBackedResolutionGapInputs`; it does not invent a second call/access identity and keeps `sourceSiteID`, `sourceNodeID`, `factFamily`, `targetText`, `filePath`, range, `sourceSiteStatus`, `proofKind`, `resolutionSource`, and note.

- [x] [P3-B2] Finalize source-backed unresolved call, access, type-reference, and heritage raw facts into persisted graph entities or persisted graph records during the analyze semantic enrichment phase. `semantic.Apply` now persists `ResolutionGap` nodes and `HAS_RESOLUTION_GAP` diagnostic relationships after App Layer/Functional Area classification and before graph compact, graph snapshot, and LadybugDB load. Existing diagnostic summaries remain compatible, and persistence consumes source-site/raw unresolved facts rather than capped report candidates.

- [x] [P3-C] Implement fine-grained gap relationships or typed gap metadata for every distinction supported by source evidence. Persisted gaps now carry `gapKind` values for unresolved call, unresolved access, unresolved type-reference, and unresolved heritage; `classification` distinguishes builtin, standard_library, test_framework, external_library, in_repo_unresolved, and unclassified when evidence supplies it; `actionability` distinguishes non_actionable, review, and analyzer_gap; `targetRole` preserves callable/member/type roles from source evidence.

- [x] [P3-C2] Update LadybugDB export/load schema and graph snapshot serialization for persisted ResolutionGap entities, relations, or records so graph JSON, DB-backed Cypher, HTTP API, MCP resources, and Web consumers read the same semantic facts. `ResolutionGap` is now a graph node label, LadybugDB node table, generated Web contract label/table, and Web filter/legend label. `HAS_RESOLUTION_GAP` is now a graph relationship type, LadybugDB relationship type, generated Web contract relationship, and Web edge metadata entry. LadybugDB CSV export/load has zero skipped relationships for the fresh analyzed graph.

- [x] [P3-D] Implement target-role inference for ResolutionGap records from fact family and source evidence. `ResolutionGapInput.InferredTargetRole` now preserves explicit source evidence, maps call/access/type-reference/heritage facts to callable/member/type roles, maps builtin/standard-library/test-framework/external classifications when fact family does not decide, and uses the inferred role in gap node properties, gap relationship metadata, and fallback gap identity without marking any unresolved target as resolved.

- [x] [P3-E] Wire existing graph-health diagnostic classification/actionability into persisted ResolutionGap records and summaries. Persisted gaps keep the graph-health diagnostic `classification` and `actionability` fields; fresh analyze produced builtin, standard_library, test_framework, external_library, and in_repo_unresolved classifications, with non_actionable/review/analyzer_gap actionability counts recorded in the benchmark ledger.

- [x] [P3-F] Add graph validation and backend tests that inspect persisted ResolutionGap output and prove unresolved targets do not create fake in-repo target nodes, fake resolved semantic edges, fake topology edges, or proofless `CALLS`/`ACCESSES`. `ValidateResolutionGapPersistence` checks dangling `HAS_RESOLUTION_GAP` relationships, non-gap relationship targets, counted topology participation, resolved-target claims on gap nodes, and semantic resolved edges that reuse unresolved gap source-site IDs. Backend tests cover accepted source-backed gaps and each violation class.

- [x] [P3-G] Add backend tests for unresolved call, access, type-reference, heritage, builtin/predeclared, standard-library, test-framework, external, in-repo analyzer-gap, unknown target-role, repeated occurrence aggregation, and multiple different target texts from the same source/fact/file/note bucket. `TestResolutionGapInputInfersTargetRole` covers inference branches, and `TestApplyPersistsResolutionGapRolesClassificationsAndOccurrences` proves target identity, sourceSiteID, source-site status, proof kind, classification, actionability, App Layer, Functional Area, and occurrence counts survive semantic persistence without synthetic target nodes or semantic edges.

- [x] [P3-H] Record persisted gap schema examples, before/after gap counts, top targets, target-role/actionability counts, and test evidence in the evidence and benchmark ledgers. Evidence E8 and benchmark B6/B11/B12 now record the persisted schema, latest graph count `58350` ResolutionGap nodes, `58350` HAS_RESOLUTION_GAP relationships, fact-family/gap-kind counts, target-role counts, classification/actionability counts, validation commands, source-site accuracy output, and the fresh analyze benchmark artifacts through `.tmp\2026-05-22-p3-role-validation-postedit-analyze.json`.

- [x] [P3-I] Implement the aggregation/dedupe policy for unresolved occurrences with exact occurrence counts, target text identity, bucket identity, representative source samples, source App Layer/Functional Area distribution, and traceability back to source diagnostics. `ResolutionGapAggregates` and `SourceBackedResolutionGapAggregates` now compute evidence-preserving aggregate buckets from source-backed inputs without reducing persisted gap entities; buckets include exact input and occurrence counts, full sourceSiteID traceability, capped representative samples, App Layer/Functional Area/file distributions, and target text in bucket identity. Tests prove repeated occurrences, sample caps, distributions, and multiple different target texts from the same source/fact/file/note fixture keep their meaning.

## Phase 4 - Resolution Health Inventory And Topology Separation

- [x] [P4-A] Define Resolution Health buckets that are separate from Topology Health. `graphhealth` now defines `resolved_references`, `unresolved_non_actionable`, `external_unresolved`, `in_repo_analyzer_gap`, `unresolved_call_target`, `unresolved_access_target`, `unresolved_type_target`, `unresolved_heritage_target`, and `unclassified_unknown`, plus `clear/degraded/unknown` Resolution Confidence. These fields are exposed in Go summary structs and generated Web contracts.

- [x] [P4-B] Update graph-health/report summary builders and tests so topology status stays topology-only while resolution status and resolution confidence are overlays. `ComputeSummary` still derives topology only from counted relationships; `HAS_RESOLUTION_GAP` remains excluded from topology. Node health now carries `resolutionHealthBuckets`, `resolutionGapCount`, and `resolutionConfidence`. Backend tests prove a connected node with gaps remains `connected` while showing degraded resolution confidence.

- [x] [P4-C] Add graph/API inventory counts by App Layer, Functional Area, fact family, target role, classification, actionability, Resolution Health bucket, and topology status from the persisted graph/inventory source of truth. `graphhealth.Summary`, HTTP graph/report/explain payloads, generated Web contracts, and `resolution-inventory` all read full persisted graph data, not capped report candidates. Cypher verification confirms LadybugDB can see `58879` `ResolutionGap` nodes, `58879` `HAS_RESOLUTION_GAP` relationships, and gap node fields such as `sourceAppLayer`, `functionalArea`, and `gapKind`.

- [x] [P4-D] Add or extend a CLI inventory command for resolution gaps and semantic graph health. `avmatrix resolution-inventory --graph .avmatrix\graph.json --out .tmp\2026-05-22-p4-resolution-inventory.json` reads persisted analyze output and prints full counts for ResolutionGap nodes/relationships, occurrence counts, resolved references, App Layer, Functional Area, fact family, target role, classification, actionability, Resolution Health bucket, Resolution Confidence, and topology overlay without applying UI/report candidate caps.

- [x] [P4-E] Add backend and CLI tests proving Resolution Health is not a replacement for topology, inventory uses persisted graph data, and connected diagnostic nodes are not ranked as topology defects. Tests added/updated in `internal/graphhealth`, `internal/cli`, `internal/httpapi`, generated contract checks, Web unit tests, and e2e validation cover the new overlay and command output.

- [x] [P4-F] Record separate Resolution Health and Topology Health examples, command output, API payload samples, count tables, and tests in evidence/benchmark. Evidence E9 and benchmark B7/B11/B12 now record the Phase 4 analyze artifact, CLI inventory artifact, source-site accuracy artifact, Cypher checks, build/test/e2e results, and count tables.

## Phase 5 - Query Health Benchmark Command

- [x] [P5-A] Define a query benchmark suite format with intent text, expected files, expected symbols, optional expected App Layer/Functional Area, hit@5 threshold, hit@10 threshold, actual top results, noise reason, and pass/fail status. `query-health` now reads JSON suites with `schemaVersion`, `suite`, and `cases[]` fields; each case records expected files/symbols, optional semantic expectations, hit thresholds, top results, matched/missed targets, noise reason, and pass/fail status.

- [x] [P5-B] Add a CLI command for query health so retrieval accuracy can be checked by running one command. `avmatrix query-health --suite <suite.json> --repo <repo> --out <report.json> --limit 10` verifies the indexed repo is not stale by commit, calls the same local MCP `query` implementation path as `avmatrix query`, scores results, writes JSON reports, and prints readable summary lines; `--json` and `--fail-on-threshold` are available for automation.

- [x] [P5-C] Add initial suite entries for unresolved reference diagnostic generation, graph health unknown-connectivity separation, App Layer/resolution-gap layout, runtime reset hidden-terminal behavior, API contract surfaces, query implementation surfaces, and frontend graph filter surfaces. The suite at `docs/query-health/2026-05-22-avmatrix-app-layer-resolution-gap-suite.json` includes the required resolver, graph-health, query, CLI, API, contract, Web graph filter/layout/optimizer, and launcher runtime files.

- [x] [P5-D] Make command output report expected targets, actual top results, matched files/symbols, hit@5, hit@10, noise reason, pass/fail, any semantic layer fields returned by query results, and explicit miss reasons for function/method targets that current definition matching fails to return. JSON output includes expected targets, matched/missed targets, top results, rank/source/file/symbol fields, optional App Layer/Functional Area/Resolution Health fields when present, and specific miss reasons for missing function/method targets.

- [x] [P5-E] Add tests for suite parsing, scoring, missing expected targets, noisy results, semantic field output, JSON/table output if both exist, and failed threshold behavior. `internal/cli/query_health_command_test.go` covers suite parsing, scoring, noisy/missing targets, semantic field preservation, JSON output, summary output, report writing, and `--fail-on-threshold`.

- [x] [P5-F] Run the command on the current repository after implementation and record baseline/final hit rates, noisy intents, and examples in the evidence and benchmark ledgers. After fresh analyze, `query-health` wrote `.tmp\2026-05-22-p5-query-health-baseline.json`; the current query implementation passed `1/7` cases and failed `6/7`, establishing the baseline that Phase 6 must improve.

## Phase 6 - Semantic Command Surfaces

- [ ] [P6-A] Update `query` retrieval and result output so matching nodes or flows can expose node type, App Layer, Functional Area, Resolution Health, and related ResolutionGap summaries when those fields are available. Add missing/stale semantic-data handling and improve the current simple contains-scoring behavior enough for the P5 benchmark intents to hit expected files/symbols.

- [ ] [P6-B] Update `context` output so a symbol/node view includes node type, App Layer, Functional Area when known, topology status, resolution-health summary, and nearby/source ResolutionGaps. The output must distinguish source-node gaps from unresolved target entities.

- [ ] [P6-C] Update `impact` output so blast-radius summaries include affected App Layers, affected Functional Areas, and resolution-health risks when graph evidence supports those summaries. Add command-output coverage proving high or critical risk warnings are reported as workflow safety information while inspection output remains available.

- [ ] [P6-D] Update `detect-changes` output so changed symbols and affected flows summarize App Layers, Functional Areas, ResolutionGap changes, and resolution-health impact. This command remains the pre-commit graph-diff check required by repository rules.

- [ ] [P6-E] Update API-specific MCP tools such as `route_map`, `shape_check`, and `api_impact` to surface App Layer, Functional Area, and Resolution Health where the persisted graph provides those fields. If a specific API tool cannot use the new semantic layer in this plan, record the exact limitation and follow-up in evidence.

- [ ] [P6-F] Add focused CLI/MCP tests or command-output tests for query/context/impact/detect-changes and API-specific MCP semantic fields, including cases where fields are unknown, missing because the graph is stale, or unavailable because the node is outside classified surfaces.

- [ ] [P6-G] Record command examples, limitations, changed output fields, and test evidence in the evidence and benchmark ledgers.

## Phase 7 - Web UI Filters, Detail Lens, And Multi-Ring Layout

- [ ] [P7-A] Add App Layer filters/lens to the Web UI using backend/API fields. Source filter values such as Backend, API, Frontend, Shared Contract, API Contract, Frontend API Client, Tests, Docs, Config, Generated, Mixed, and Unknown from graph data and render a missing-data state when those fields are absent.

- [ ] [P7-A2] Extend Web graph state management in `avmatrix-web/src/hooks/app-state/graph.tsx` and related app-state types for App Layer and Resolution Health filters. The state must compose with existing node type, edge type, focus-depth, selected-node, and Graph Health filters without resetting unrelated controls.

- [ ] [P7-B] Add Resolution Health filters/lens for fact family, target role, classification, actionability, analyzer-gap concentration, top unresolved target text, and source App Layer. Minimum user-facing lens rows must include Backend unresolved calls, API unresolved handlers/contracts, Frontend unresolved type refs, Shared contract analyzer gaps, External unresolved symbols, Builtin/Test/Stdlib non-actionable references, In-repo analyzer gaps, Resolution gaps by functional area, Top app layers by analyzer gap count, Top functional areas by unresolved count, and Top unresolved target text. These filters must compose with existing node type, edge type, graph health, focus-depth, and selected-node filters.

- [ ] [P7-C] Add node/detail-panel explanations for App Layer, Functional Area, Topology Health, Resolution Health, and related ResolutionGaps. Include user-facing copy and tests that label degraded resolution confidence separately from dead-code conclusions.

- [ ] [P7-D] Replace or extend the current deterministic filter-based clustered placement with deterministic multi-ring placement where App Layer is the macro position and node type or ResolutionGap kind is the micro island inside that ring. Backend, API, and Frontend must be distinct rings when present, with API placed between Backend and Frontend as the bridge; Shared/API Contract rings should sit near API, and Frontend API Client should sit near Frontend/API. Shared/API Contract/Frontend API Client/Test/Docs/Config/Generated/Mixed/Unknown may become additional rings when the graph contains those categories.

- [ ] [P7-E] Update Web graph adapter color and island assignment so existing node type/filter colors stay grouped inside each App Layer ring. Add assertions or screenshot-backed checks that a yellow node type stays in the yellow island, a blue node type stays in the blue island, and unrelated colors are not interleaved in the same island.

- [ ] [P7-F] Add layout optimizer invocation guards and tests around render, data load, filter changes, and ring placement so optimizer execution is only triggered by the manual Web UI control. Verify the implementation has no timeout, delayed refresh, or elapsed-time budget path for layout behavior.

- [ ] [P7-G] Define ring size, spacing, ordering, and default visibility rules before coding layout behavior. Rings may be large or small, and there is no fixed maximum number of rings, but the UI must avoid overlap, must keep node type islands readable, and must explicitly decide which rings/lenses are visible by default versus collapsed or hidden by default. Because rendered node size is already capped at `3`, this task must focus on island spacing, island radius, ring radius, label density, edge density, and screenshot evidence rather than assuming node-size caps solve the unreadable "metal block" layout.

- [ ] [P7-H] Add Web unit tests for App Layer filters, Resolution Health filters, detail-panel fields, ring grouping, color/type island behavior, ring size/default visibility policy, graph-health composition, and no auto optimizer invocation.

- [ ] [P7-I] Add Web e2e tests in a real browser proving Backend/API/Frontend rings are visible with API between Backend and Frontend, additional rings appear when data exists, ResolutionGaps are visible/filterable, filters do not collapse topology and resolution health, node colors remain grouped by type, default visibility behaves as specified, and the optimizer only runs from manual user action.

- [ ] [P7-J] Record screenshots, Playwright artifacts, visible ring/filter counts, ring placement evidence, default visibility evidence, and UI behavior evidence in the evidence and benchmark ledgers.

## Phase 8 - Full Validation And Closure

- [ ] [P8-A] Run the full build gate before tests: `powershell -ExecutionPolicy Bypass -File avmatrix-launcher\build.ps1`. Record command output and generated artifact location in evidence.

- [ ] [P8-B] Run backend tests for analyze semantic enrichment flow, App Layer classification, Functional Area classification, proof-based CALLS/ACCESSES source-site inventory, ResolutionGap persistence, LadybugDB export/load, fine-grained gap relations, Resolution Health inventory, graph/API summaries, CLI inventory, query-health command, API-specific MCP tools, and semantic command output.

- [ ] [P8-C] Run contract generation/checks from the Go contract source and verify generated Web types expose App Layer, Functional Area, source-site status/proof metadata, ResolutionGap, Resolution Health, relation metadata, and query/command-facing enum values. Confirm the generated TypeScript diff comes from the Go contract source.

- [ ] [P8-D] Run Web unit tests for filters, detail panels, graph layout, manual optimizer behavior, and generated contract usage.

- [ ] [P8-E] Run Web e2e tests covering multi-ring layout, App Layer filters, Resolution Health filters, persisted ResolutionGap visibility, node type island grouping, and manual-only optimizer behavior.

- [ ] [P8-F] Run the query-health benchmark command and record final hit@5/hit@10, pass/fail, noisy intents, and regression notes.

- [ ] [P8-G] Run the resolution inventory command and record final counts by App Layer, Functional Area, fact family, target role, actionability, Resolution Health bucket, and topology status.

- [ ] [P8-H] Run `query`, `context`, `impact`, and `detect-changes` examples and record semantic output or explicit limitations in the evidence ledger.

- [ ] [P8-I] Run AVmatrix detect-changes according to repository rules before implementation commits and record the affected symbols/flows in evidence, with the doc-only commit exception handled by the repository rules.

- [ ] [P8-J] Update this plan, the evidence ledger, and the benchmark ledger to implemented status only after all required validation passes or after any failed validation is recorded with a clear follow-up plan.
