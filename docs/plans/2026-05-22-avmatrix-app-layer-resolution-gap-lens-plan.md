# AVmatrix App Layer Resolution Gap Lens Plan

Date: 2026-05-22

Status: planned

Source discussion:

- [reports/problem/2026-05-22-avmatrix-app-layer-resolution-gap-discussion.md](../../reports/problem/2026-05-22-avmatrix-app-layer-resolution-gap-discussion.md)

Companion files:

- Benchmark ledger: [2026-05-22-avmatrix-app-layer-resolution-gap-lens-benchmark.md](2026-05-22-avmatrix-app-layer-resolution-gap-lens-benchmark.md)
- Evidence ledger: [2026-05-22-avmatrix-app-layer-resolution-gap-lens-evidence.md](2026-05-22-avmatrix-app-layer-resolution-gap-lens-evidence.md)

## Rules

1. Follow active workspace and repository instructions, including `AGENTS.md`; this plan records product work and validation, it does not replace repository rules.
2. Use AVmatrix according to active repository instructions for implementation slices; do not use AVmatrix for doc-only commits.
3. As each task is completed, update the corresponding checklist item immediately.
4. Run a full build before testing. For this plan the full build gate is `powershell -ExecutionPolicy Bypass -File avmatrix-launcher\build.ps1`.
5. Because this changes graph semantics, CLI command behavior, API contracts, and Web UI graph behavior, validation must include backend tests, contract checks, Web unit tests, and Web e2e tests.
6. Record benchmark results as each benchmarkable task is completed. Benchmarkable means measured product/runtime performance, capacity, package/startup size, graph/DB throughput, query hit rate, command output inventory, or graph inventory counts; build/test/e2e timings are validation evidence unless the slice changes those systems.
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

## Scope Boundary

Implementation may touch:

- graph node and relationship schema;
- analyzer/graph generation metadata;
- resolution diagnostic emission and graph-health diagnostic attachment;
- persisted graph snapshot shape under `.avmatrix`;
- graph-health and resolution-health summaries;
- HTTP graph APIs and generated Web contracts;
- CLI/query/MCP command output for `query`, `context`, `impact`, `detect-changes`, and any new query-health or resolution-inventory command;
- Web graph filters, graph detail panels, dashboards, generated types, and layout placement;
- backend, contract, Web unit, and Web e2e tests.

Out of scope unless a later phase explicitly reopens it:

- treating unresolved references as confirmed dead code;
- merging `unknown_connectivity` back into `unresolved_reference`;
- synthesizing fake resolved in-repo target nodes or fake semantic edges for unresolved targets;
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

## Acceptance Criteria

- Graph nodes expose a persisted primary App Layer category with no overlapping primary labels.
- API and API-related mixed categories are first-class, not hidden under Backend or Frontend.
- Functional Area metadata is persisted only where evidence is accurate enough; ambiguous nodes remain unknown.
- Source-backed unresolved references are represented as persisted ResolutionGap/UnresolvedSymbol graph entities or equivalent persisted graph records, not only diagnostic text.
- Fine-grained gap relationships or typed gap metadata preserve call, access, type-reference, heritage, external, builtin, test, analyzer-gap, and unknown distinctions where evidence supports them.
- Topology Health and Resolution Health remain separate in graph/API/CLI/Web output.
- CLI/query command surfaces can report App Layer, Functional Area, ResolutionGap, and resolution-health information from persisted graph data.
- A query-health benchmark command exists and reports hit@5/hit@10, expected files/symbols, actual results, noise reason, and pass/fail.
- Web UI exposes App Layer filters, Resolution Health filters, and a multi-ring layout where App Layer controls macro placement and node type/gap kind controls islands inside rings.
- The Web UI does not invent App Layer, Functional Area, or ResolutionGap truth on the client.
- Layout optimizer remains manual-only and is not auto-run after render.
- Full build, backend tests, contract checks, Web unit tests, and Web e2e tests pass before closure.
- Benchmark and evidence ledgers contain baseline and final inventories for App Layer, Functional Area, ResolutionGap, Resolution Health, query benchmark, CLI semantic output, and Web rings/filters.
- Missing App Layer or ResolutionGap metadata in loaded graph data is treated as stale/incomplete graph evidence, not as a reason to guess at API/UI load time.
- If ResolutionGap data is aggregated or deduped, the aggregate keeps exact counts and representative source evidence without capping away meaning.

## Current Code Facts To Verify

The following facts are expected from the recent graph-health work and must be verified in Phase 0 before implementation edits:

- unresolved references are currently emitted as diagnostics from resolution paths and attached to source nodes;
- `unknown_connectivity` is already separated from ordinary unresolved diagnostics;
- Web graph filters currently consume graph/API metadata, generated contracts, and client-side filter state;
- graph layout already has filter-based clustering and a manual optimizer path;
- command surfaces exist for query/context/impact/detect-changes but do not yet expose the proposed App Layer and ResolutionGap semantics consistently;
- generated Web contracts are the boundary that should prevent the UI from relying on ad hoc shape guesses.

## Phase 0 - Baseline, Discussion Coverage, And Source Trace

- [ ] [P0-A] Read the full discussion report and write a coverage table in the evidence ledger that maps every major decision to this plan: node type insufficiency, BE/API/FE/App Layer rings, non-overlapping mixed categories, Functional Area accuracy, persisted ResolutionGap, fine-grained gap relations, Resolution Health, query-health command, semantic command output, multi-ring Web layout, no stale fallback, and no timeout/auto optimizer behavior.

- [ ] [P0-B] Run `avmatrix analyze --force` at implementation start, then record scanned/parsed/unsupported/failed file counts, graph node count, graph relationship count, counted semantic relationship count, execution-flow count, `unknown_connectivity` count, and graph timestamp/hash in the benchmark ledger. Compare the fresh baseline against the discussion observations of about `22010` nodes, `26906` counted semantic relationships, `0` `unknown_connectivity`, `51232` unresolved occurrences, and `8880` unresolved buckets without assuming those old numbers are still exact.

- [ ] [P0-C] Trace the current unresolved-reference pipeline from resolution source facts to graph/API output. Record the exact files/symbols that create unresolved call, access, type-reference, and heritage diagnostics; record where target text, source node, fact family, classification, actionability, and source location are currently kept or lost.

- [ ] [P0-D] Measure the current unresolved-reference inventory before changing semantics. Record bucket count, occurrence count, fact family counts, diagnostic classification/actionability counts if present, top target texts, source node labels, source path buckets, and whether each source already has a topology status.

- [ ] [P0-E] Measure a provisional path/package-derived App Layer inventory without changing product behavior. Seed the audit with the discussion examples `avmatrix-web/src/**`, `avmatrix-web/test/**`, `avmatrix-web/e2e/**`, `internal/**`, `cmd/**`, `internal/httpapi/**`, `avmatrix-web/src/services/backend-client.ts`, `avmatrix-launcher/**`, `contracts/**`, `internal/contracts/**`, `cmd/generate-web-contracts/**`, `docs/**`, `reports/**`, `*.md`, `*_test.go`, `test/fixtures/**`, config files, package files, and build scripts. The evidence must show candidate counts for backend, api, frontend, cli_launcher, shared_contract, api_contract, api_shared_contract, frontend_api_client, backend_test, frontend_test, api_test, generated_contract, docs, config, generated, mixed, and unknown, plus examples explaining every uncertain bucket.

- [ ] [P0-F] Audit current query behavior with fixed benchmark intents before adding the new command. At minimum test intents for unresolved reference diagnostic generation, graph health unknown-connectivity separation, App Layer/resolution-gap layout, and runtime reset hidden-terminal behavior; record expected files/symbols, actual top results, hit/miss, and noise reason.

- [ ] [P0-G] Locate the implementation surfaces for graph schema, graph generation, resolution diagnostics, graph health, HTTP graph APIs, generated Web contracts, CLI query/context/impact/detect-changes, Web graph filters, Web graph detail panels, and Web layout. Record exact file paths in evidence so later phases edit the identified surfaces rather than guessing.

## Phase 1 - App Layer Taxonomy And Persistence

- [ ] [P1-A] Define the persisted App Layer category registry as a primary one-of field, not overlapping tags. The category list must include normal categories and mixed categories where needed: backend, api, frontend, cli_launcher, shared_contract, api_contract, api_shared_contract, frontend_api_client, backend_test, frontend_test, api_test, generated_contract, docs, config, generated, mixed, unknown, and any additional categories proven necessary by P0 evidence.

- [ ] [P1-B] Define the source evidence required for every App Layer category. Use high-confidence signals such as path, package/module, generated-contract location, API route/handler location, frontend source roots, test naming, docs/config paths, and launcher/runtime ownership; start from the P0-E path seed list and refine it with source inspection. Do not classify a node into a precise category when evidence is insufficient.

- [ ] [P1-C] Implement App Layer classification during analyze/graph generation and persist the result in graph output. The Web UI and API consumers must read the persisted field; they must not classify nodes at load time except for defensive display of missing data.

- [ ] [P1-D] Make API first-class by classifying server handlers, API graph endpoints, API contract/schema code, generated API contract files, and frontend API clients into separate categories when evidence supports that separation. If a surface is both API and shared contract, use a mixed category such as `api_shared_contract` rather than overlapping labels.

- [ ] [P1-E] Update graph schema, snapshot serialization, HTTP graph payloads, generated Web contracts, and any contract tests so App Layer values are stable public fields. A freshly analyzed but ambiguous node may use `unknown`; graph data that lacks the new metadata entirely must be treated as stale/incomplete schema evidence and must not trigger load-time classification heuristics.

- [ ] [P1-F] Add tests for simple and mixed App Layer examples. Coverage must prove one primary App Layer per node, correct API classification, correct mixed category classification, correct unknown handling for ambiguous input, and no accidental multi-label primary classification.

- [ ] [P1-G] Record before/after App Layer counts, examples by category, changed schema fields, generated contract output, and test evidence in the benchmark and evidence ledgers.

- [ ] [P1-H] Define user-facing and machine-facing names for App Layer, API Layer, API Contract, Frontend API Client, Resolution Gap, Unresolved Symbol, Analyzer Gap, External Reference, and Non-actionable Reference. Record enum keys, display labels, CLI labels, and Web labels so API, CLI, and UI do not drift.

## Phase 2 - Functional Area Accuracy Gate

- [ ] [P2-A] Evaluate candidate Functional Area signals from P0 evidence: path prefix, package/module name, process membership, community detection, import/call neighborhood, explicit config, and any AI-assisted labeling only if it is reproducible and verifiable. Record accepted and rejected signals with examples instead of choosing a weak signal because it is easy to code.

- [ ] [P2-B] Define a Functional Area registry for high-confidence areas only. Initial candidates may include resolution, graph_health, query, mcp, web_graph_ui, layout, contracts, providers, runtime, analyzer, session, launcher, cli, reporting, and unknown; each accepted area must have exact evidence rules.

- [ ] [P2-C] Persist Functional Area metadata in graph output only for nodes that meet accepted rules. Ambiguous nodes must stay `unknown` or equivalent rather than being forced into a functional group.

- [ ] [P2-D] Expose Functional Area through API payloads, generated contracts, CLI surfaces, and Web detail/filter data where available. Consumers must distinguish "unknown because not enough evidence" from "missing field because graph was not freshly analyzed".

- [ ] [P2-E] Add tests for accepted Functional Area rules, rejected low-confidence rules, ambiguous nodes, and command/API/Web contract visibility.

- [ ] [P2-F] Record selected rules, rejected rules, counts, unknown counts, example nodes, and test evidence in the evidence and benchmark ledgers.

## Phase 3 - Persisted ResolutionGap And UnresolvedSymbol Model

- [ ] [P3-A] Define the persisted ResolutionGap/UnresolvedSymbol data model. It must preserve source node ID, source App Layer, source Functional Area when known, fact family, target text, inferred target role, classification, actionability, resolution source, source file path, line/column when available, occurrence count, sample evidence, and notes for unclassified cases.

- [ ] [P3-B] Promote source-backed unresolved call, access, type-reference, and heritage diagnostics into persisted graph entities or persisted graph records. Existing diagnostic summaries must remain compatible, but diagnostic text alone must no longer be the only machine-readable representation.

- [ ] [P3-C] Implement fine-grained gap relationships or typed gap metadata for every distinction supported by source evidence. Start with unresolved call, unresolved access, unresolved type-reference, unresolved heritage, external symbol, builtin/stdlib reference, test-framework reference, in-repo analyzer gap, and unknown/unclassified; add narrower types if P0 evidence shows they are useful.

- [ ] [P3-D] Infer target role from fact family and source evidence without pretending the target is resolved. Examples: calls map to callable-like gaps, member accesses map to member-like gaps, type annotations and heritage map to type-like gaps, and external/builtin/test evidence maps to those roles only when classification supports it.

- [ ] [P3-E] Preserve actionability as graph-readable data. Builtin/stdlib/test references should be non-actionable when confidently recognized, external references should be reviewable unless rules say otherwise, in-repo unresolved references should be analyzer gaps, and unclassified references should remain review/unknown until better evidence exists.

- [ ] [P3-F] Do not synthesize fake in-repo target nodes, fake resolved semantic edges, or fake topology edges for unresolved targets. The persisted gap may be a node/entity, but it must not claim resolution that the analyzer did not prove.

- [ ] [P3-G] Add backend tests for unresolved call, access, type-reference, heritage, builtin/predeclared, standard-library, test-framework, external, in-repo analyzer-gap, unknown target-role, and repeated occurrence aggregation cases.

- [ ] [P3-H] Record persisted gap schema examples, before/after gap counts, top targets, target-role/actionability counts, and test evidence in the evidence and benchmark ledgers.

- [ ] [P3-I] If unresolved occurrences are aggregated or deduped, preserve exact occurrence counts, bucket identity, representative source samples, source App Layer/Functional Area distribution, and traceability back to source diagnostics. Do not cap, sample, or collapse evidence in a way that changes the meaning of `51232`-scale unresolved inventories.

## Phase 4 - Resolution Health Inventory And Topology Separation

- [ ] [P4-A] Define Resolution Health buckets that are separate from Topology Health. Required buckets include resolved references when measurable, unresolved non-actionable, external unresolved, in-repo analyzer gap, unresolved call target, unresolved access target, unresolved type target, unresolved heritage target, and unclassified/unknown.

- [ ] [P4-B] Update graph-health and report summaries so topology status stays topology-only while resolution status and resolution confidence are overlays. A node with `connected` topology and unresolved gaps must remain connected while showing degraded resolution confidence.

- [ ] [P4-C] Add graph/API inventory counts by App Layer, Functional Area, fact family, target role, classification, actionability, Resolution Health bucket, and topology status. The same source of truth must be usable by CLI and Web consumers.

- [ ] [P4-D] Add or extend a CLI inventory command for resolution gaps and semantic graph health. It must read persisted analyze output and print the same counts available to API/Web consumers, including App Layer and Functional Area grouping.

- [ ] [P4-E] Add backend and CLI tests proving Resolution Health is not a replacement for topology, inventory uses persisted graph data, and connected diagnostic nodes are not ranked as topology defects.

- [ ] [P4-F] Record separate Resolution Health and Topology Health examples, command output, API payload samples, count tables, and tests in evidence/benchmark.

## Phase 5 - Query Health Benchmark Command

- [ ] [P5-A] Define a query benchmark suite format with intent text, expected files, expected symbols, optional expected App Layer/Functional Area, hit@5 threshold, hit@10 threshold, actual top results, noise reason, and pass/fail status.

- [ ] [P5-B] Add a CLI command for query health so retrieval accuracy can be checked by running one command. The command must use fresh analyze output, read the suite, run each query intent, score results, and write readable table or JSON output suitable for evidence and future automation.

- [ ] [P5-C] Add initial suite entries for unresolved reference diagnostic generation, graph health unknown-connectivity separation, App Layer/resolution-gap layout, runtime reset hidden-terminal behavior, API contract surfaces, and frontend graph filter surfaces. The first suite must include expected files from the discussion: `internal/resolution/resolve.go`, `internal/resolution/emit.go`, `internal/graphhealth/diagnostics.go`, `internal/graphhealth/compute.go`, `internal/graphhealth/policy.go`, `avmatrix-web/src/lib/graph-health-filters.ts`, Web graph layout code, layout optimizer code, and `avmatrix-launcher/src/main.go`.

- [ ] [P5-D] Make command output report expected targets, actual top results, matched files/symbols, hit@5, hit@10, noise reason, pass/fail, and any semantic layer fields returned by query results. This output must make noisy retrieval visible rather than hiding it behind a generic score.

- [ ] [P5-E] Add tests for suite parsing, scoring, missing expected targets, noisy results, semantic field output, JSON/table output if both exist, and failed threshold behavior.

- [ ] [P5-F] Run the command on the current repository after implementation and record baseline/final hit rates, noisy intents, and examples in the evidence and benchmark ledgers.

## Phase 6 - Semantic Command Surfaces

- [ ] [P6-A] Update `query` result output so matching nodes or flows can expose node type, App Layer, Functional Area, Resolution Health, and related ResolutionGap summaries when those fields are available. The command must not make up fields when the graph lacks fresh semantic data.

- [ ] [P6-B] Update `context` output so a symbol/node view includes node type, App Layer, Functional Area when known, topology status, resolution-health summary, and nearby/source ResolutionGaps. The output must distinguish source-node gaps from unresolved target entities.

- [ ] [P6-C] Update `impact` output so blast-radius summaries include affected App Layers, affected Functional Areas, and resolution-health risks when graph evidence supports those summaries. High or critical risk warnings remain informational for workflow safety; they must not mean the tool refuses to inspect or change code.

- [ ] [P6-D] Update `detect-changes` output so changed symbols and affected flows summarize App Layers, Functional Areas, ResolutionGap changes, and resolution-health impact. This command remains the pre-commit graph-diff check required by repository rules.

- [ ] [P6-E] Add focused CLI/MCP tests or command-output tests for query/context/impact/detect-changes semantic fields, including cases where fields are unknown, missing because the graph is stale, or unavailable because the node is outside classified surfaces.

- [ ] [P6-F] Record command examples, limitations, changed output fields, and test evidence in the evidence and benchmark ledgers.

## Phase 7 - Web UI Filters, Detail Lens, And Multi-Ring Layout

- [ ] [P7-A] Add App Layer filters/lens to the Web UI using backend/API fields. The UI must show categories such as Backend, API, Frontend, Shared Contract, API Contract, Frontend API Client, Tests, Docs, Config, Generated, Mixed, and Unknown only from graph data, not client-side invented classification.

- [ ] [P7-B] Add Resolution Health filters/lens for fact family, target role, classification, actionability, analyzer-gap concentration, top unresolved target text, and source App Layer. Minimum user-facing lens rows must include Backend unresolved calls, API unresolved handlers/contracts, Frontend unresolved type refs, Shared contract analyzer gaps, External unresolved symbols, Builtin/Test/Stdlib non-actionable references, In-repo analyzer gaps, Resolution gaps by functional area, Top app layers by analyzer gap count, Top functional areas by unresolved count, and Top unresolved target text. These filters must compose with existing node type, edge type, graph health, focus-depth, and selected-node filters.

- [ ] [P7-C] Add node/detail-panel explanations for App Layer, Functional Area, Topology Health, Resolution Health, and related ResolutionGaps. The detail panel must show why a node has degraded resolution confidence without calling it dead code by default.

- [ ] [P7-D] Implement deterministic multi-ring placement where App Layer is the macro position and node type or ResolutionGap kind is the micro island inside that ring. Backend, API, and Frontend must be distinct rings when present, with API placed between Backend and Frontend as the bridge; Shared/API Contract rings should sit near API, and Frontend API Client should sit near Frontend/API. Shared/API Contract/Frontend API Client/Test/Docs/Config/Generated/Mixed/Unknown may become additional rings when the graph contains those categories.

- [ ] [P7-E] Preserve existing node type/filter colors and keep same-type islands visually together inside their App Layer ring. A yellow node type should stay in the yellow island, a blue node type in the blue island, and unrelated colors must not be interleaved in the same island.

- [ ] [P7-F] Keep layout optimizer manual-only. Do not auto-run optimizer after render, after data load, after filter changes, or after ring placement; do not add timeout, delayed refresh, or elapsed-time budget behavior to hide layout work.

- [ ] [P7-G] Define ring size, spacing, ordering, and default visibility rules before coding layout behavior. Rings may be large or small, and there is no fixed maximum number of rings, but the UI must avoid overlap, must keep node type islands readable, and must explicitly decide which rings/lenses are visible by default versus collapsed or hidden by default.

- [ ] [P7-H] Add Web unit tests for App Layer filters, Resolution Health filters, detail-panel fields, ring grouping, color/type island behavior, ring size/default visibility policy, graph-health composition, and no auto optimizer invocation.

- [ ] [P7-I] Add Web e2e tests in a real browser proving Backend/API/Frontend rings are visible with API between Backend and Frontend, additional rings appear when data exists, ResolutionGaps are visible/filterable, filters do not collapse topology and resolution health, node colors remain grouped by type, default visibility behaves as specified, and the optimizer only runs from manual user action.

- [ ] [P7-J] Record screenshots, Playwright artifacts, visible ring/filter counts, ring placement evidence, default visibility evidence, and UI behavior evidence in the evidence and benchmark ledgers.

## Phase 8 - Full Validation And Closure

- [ ] [P8-A] Run the full build gate before tests: `powershell -ExecutionPolicy Bypass -File avmatrix-launcher\build.ps1`. Record command output and generated artifact location in evidence.

- [ ] [P8-B] Run backend tests for App Layer classification, Functional Area classification, ResolutionGap persistence, fine-grained gap relations, Resolution Health inventory, graph/API summaries, CLI inventory, query-health command, and semantic command output.

- [ ] [P8-C] Run contract generation/checks and verify generated Web types expose App Layer, Functional Area, ResolutionGap, Resolution Health, relation metadata, and query/command-facing enum values.

- [ ] [P8-D] Run Web unit tests for filters, detail panels, graph layout, manual optimizer behavior, and generated contract usage.

- [ ] [P8-E] Run Web e2e tests covering multi-ring layout, App Layer filters, Resolution Health filters, persisted ResolutionGap visibility, node type island grouping, and manual-only optimizer behavior.

- [ ] [P8-F] Run the query-health benchmark command and record final hit@5/hit@10, pass/fail, noisy intents, and regression notes.

- [ ] [P8-G] Run the resolution inventory command and record final counts by App Layer, Functional Area, fact family, target role, actionability, Resolution Health bucket, and topology status.

- [ ] [P8-H] Run `query`, `context`, `impact`, and `detect-changes` examples and record semantic output or explicit limitations in the evidence ledger.

- [ ] [P8-I] Run AVmatrix detect-changes according to repository rules before implementation commits and record the affected symbols/flows in evidence. Do not run this for doc-only commits.

- [ ] [P8-J] Update this plan, the evidence ledger, and the benchmark ledger to implemented status only after all required validation passes or after any failed validation is recorded with a clear follow-up plan.
