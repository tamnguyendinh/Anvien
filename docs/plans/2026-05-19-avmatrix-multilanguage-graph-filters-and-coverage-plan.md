# AVmatrix Multi-Language Graph Filters and Coverage Plan

Date: 2026-05-19

Status: complete - zero-trust follow-up closure recorded 2026-05-19

Companion files:

- Benchmark ledger: [2026-05-19-avmatrix-multilanguage-graph-filters-and-coverage-benchmark.md](2026-05-19-avmatrix-multilanguage-graph-filters-and-coverage-benchmark.md)
- Evidence ledger: [2026-05-19-avmatrix-multilanguage-graph-filters-and-coverage-evidence.md](2026-05-19-avmatrix-multilanguage-graph-filters-and-coverage-evidence.md)

## Rules

1. Use AVmatrix for codebase analysis and impact checks while working on implementation slices in this plan.
2. As each task is completed, update the corresponding checklist item immediately.
3. Run a full build before testing; the test suite must include an e2e test.
4. Record benchmark results as each benchmarkable task is completed. Benchmarkable means measured product/runtime performance, capacity, package/startup size, graph/DB throughput, or conversion-inventory counts; build/test/e2e timings are validation evidence unless the slice changes those systems.
5. Record evidence as each evidenced task is completed.
6. For doc-only commits, do not use AVmatrix.
7. After each completed implementation slice, commit the work, then continue until the full plan is complete.

## Zero-Trust Reopen Note

The previous final closure is reopened after zero-trust review on 2026-05-19.

The core trigger bugs are still considered substantially addressed: duplicate same-pair `EXTENDS` / `INHERITS` display is grouped, and the audited `Restaurant_manager` TypeScript heritage sites now emit ScopeIR facts with resolved in-repo graph edges where applicable. The reopened scope is about overclaimed closure criteria:

- `LANGUAGE_GRAPH_COVERAGE` is currently too generic for provider-backed languages and does not explicitly classify every language/fact-family/node/relationship status.
- The `Restaurant_manager` large-graph e2e smoke depends on the first `/api/repos` entry and is not a stable regression guard.
- The `Restaurant_manager` TypeScript heritage trace is optional unless `AVMATRIX_RESTAURANT_MANAGER_ROOT` is set.
- Provider graph parity evidence is partly count-level and representative, not endpoint-level proof for every claimed fact family.
- Benchmark/evidence/plan text needed stale closure and pending-language drift cleanup.

Phase 8 records the zero-trust follow-up closure. Historical baseline `pending` rows remain only as baseline history; active follow-up requirements are recorded as closed in the benchmark/evidence ledgers.

## Problem

The loaded graph can make graph filters, legend rows, edge categories, and language-specific graph facts look incomplete or misleading.

The concrete trigger is `E:\Restaurant_manager`, but this is only a symptom:

- `Class=3` and `Constructor=3` look suspicious at first, but source verification shows they match the real TypeScript class declarations and constructors in the repo.
- `EXTENDS=6` and `INHERITS=6` do not mean `12` independent relationships. They are the same `6` source-target heritage pairs emitted as two relationship types.
- TypeScript source contains at least `14` `interface ... extends ...` sites and `2` `class ... extends ...` sites, but the current graph has no `EXTENDS` / `INHERITS` / `IMPLEMENTS` relationships involving those TypeScript class/interface heritage sites.
- While fixing the Web UI display for these heritage facts, the tool shell must also keep the graph workflow usable: the top bar needs a Back arrow/button beside `AVmatrix` to return to `Start-AVmatrix.html`, and the left dashboard needs drag resizing so dense node/edge/legend controls can be inspected comfortably.

These trigger findings are not the full boundary of the work. AVmatrix is a multi-language graph tool, so every Web UI filter and graph fact category must be checked against every code language declared by the scanner and generated Web contracts, not only TypeScript and Go and not only heritage relationships.

This is not just a Web UI display issue and not just an analyzer issue. It has five related parts:

- graph filter/display completeness: node type filters, edge type filters, focus-depth behavior, color legend rows, graph adapter conversion, and dashboard counts must match the actual graph contract and current graph payload;
- relationship semantics/display: `EXTENDS` and `INHERITS` can represent the same underlying heritage fact, so the UI must not present duplicate compatibility edges as independent codebase facts;
- extraction/resolution coverage: graph facts that exist in source must either be connected to resolved graph targets or represented as unresolved/external audit data, not silently disappear;
- supported-language contract coverage: JavaScript, TypeScript, Python, Java, C, C++, C#, Go, Ruby, Rust, PHP, Kotlin, Swift, Dart, Vue, Svelte, Astro, and Cobol must each be classified for supported node labels, relationship types, extraction support, graph-resolution support, and Web display/filter behavior;
- Web UI shell ergonomics: users need a clear route back to the Start screen and a resizable left dashboard while inspecting corrected graph filters and dense graph data.

## Scope

Implementation may touch:

- `internal/scanner/*`
- `internal/parser/*`
- `internal/scopeir/*`
- `internal/providers/*`
- `internal/resolution/*`
- `internal/graph/*`
- `internal/contracts/*`
- `internal/mcp/*`
- `avmatrix-web/src/components/Header.tsx`
- `avmatrix-web/src/lib/constants.ts`
- `avmatrix-web/src/lib/graph-adapter.ts`
- `avmatrix-web/src/lib/graph-links-visibility.ts`
- `avmatrix-web/src/lib/graph-edge-visibility-mode.ts`
- `avmatrix-web/src/hooks/useSigma.ts`
- `avmatrix-web/src/components/FileTreePanel.tsx`
- `avmatrix-web/src/App.tsx`
- `Start-AVmatrix.html`
- `avmatrix-launcher/src/main.go`
- launcher static asset packaging/serving paths that affect `Start-AVmatrix.html`
- Web and Go tests for node/edge filter inventory, language coverage, graph conversion, dashboard controls, and e2e display

This plan does not change the verified `Class=3` and `Constructor=3` count for `E:\Restaurant_manager` unless a later source audit proves new class declarations were missed.

## Acceptance Guardrails

- A user looking at the Web UI must not interpret one underlying heritage source-target pair as two independent relationships only because both `EXTENDS` and `INHERITS` were emitted for compatibility.
- If `INHERITS` remains in the graph payload, the UI must explain or group it as normalized/compatibility heritage rather than an additional source-code edge.
- Every supported code language must have an explicit graph coverage classification for node labels, relationship types, source fact extraction, graph resolution, and Web filter/display behavior: supported and resolved, supported but unresolved/external, supported only at extraction level, not applicable by language semantics, or scanned but not extracted.
- Every filterable node label in `NODE_LABELS` and every filterable relationship type in `GRAPH_RELATIONSHIP_TYPES` must either be displayed and toggled correctly or explicitly classified as absent from the current graph.
- For every language that can express a graph fact category, source sites must have deterministic graph behavior:
  - resolved in-repo targets become graph relationships;
  - unresolved external targets are explicitly represented or explicitly recorded as unresolved evidence;
  - source sites must not disappear without an audit trail.
- Analyzer and UI behavior must remain language-system-wide. Go embedded structs, TS class/interface heritage, Java/C#/Python/Kotlin/Swift/Dart/PHP/C++/Ruby/Rust graph forms, script-container facts from Vue/Svelte/Astro, and future provider facts should follow one documented graph contract.
- Benchmark/evidence must report raw node/edge counts, filter-visible counts, per-language coverage matrix counts, semantic duplicate counts, and missing/unresolved source-site counts.
- The Web UI must provide a clear Back arrow/button beside the `AVmatrix` title that returns to `Start-AVmatrix.html`, with accessible naming and keyboard/click support.
- The left dashboard must be resizable by dragging its boundary, with bounded min/max widths, no layout overlap, and no breakage of node type, edge type, legend, or graph canvas interactions.

## Baseline Findings

- [x] [P0-A] Verify whether `Class=3` and `Constructor=3` in `E:\Restaurant_manager` are real. Result: verified as real source counts for `SSEListener`, `ApiError`, and `ErrorBoundary`.
- [x] [P0-B] Verify whether `EXTENDS=6` and `INHERITS=6` are independent relationships. Result: all `6` source-target pairs are duplicated as both `EXTENDS` and `INHERITS`.
- [x] [P0-C] Trace graph emission source. Result: `emitHeritageCompatibilityEdges` emits `EXTENDS` or `IMPLEMENTS`, then emits `ReferenceInherits` when compatibility is enabled; `ReferenceInherits` maps to `INHERITS`.
- [x] [P0-D] Check TypeScript source heritage sites in `E:\Restaurant_manager`. Result: at least `14` `interface extends` sites and `2` `class extends` sites exist.
- [x] [P0-E] Check whether current graph has TS class/interface heritage edges. Result: no `EXTENDS` / `INHERITS` / `IMPLEMENTS` edges were found involving the audited TS class/interface heritage sites.
- [x] [P0-F] Identify the supported code-language surface from scanner/Web contracts. Result: `javascript`, `typescript`, `python`, `java`, `c`, `cpp`, `csharp`, `go`, `ruby`, `rust`, `php`, `kotlin`, `swift`, `dart`, `vue`, `svelte`, `astro`, and `cobol` must be covered by the graph filter and graph fact contract matrix.
- [x] [P0-G] Identify current Web filter source. Result: filterable node labels come from generated `NODE_LABELS`, filterable edge types come from generated `GRAPH_RELATIONSHIP_TYPES`, while default visibility, colors, sizes, labels, and graph adapter behavior are maintained in Web constants/adapter code and must be audited against all graph labels/types.
- [x] [P0-H] Run AVmatrix analyze and impact/context checks for the implementation surface. Result: `FileTreePanel` impacts `App.tsx` and `FileTreePanel.dashboard-completeness.test.tsx`; `knowledgeGraphToGraphology` impacts `GraphCanvas.tsx` and `graph-adapter.edge-geometry.test.ts`; `getGraphEdgeVisibilityMode` impacts `useSigma.ts` and `graph-edge-visibility-mode.test.ts`; `WebUIContract` feeds generated Web graph/language constants; `parseFiles` routes extractor support and has tests for most provider languages.
- [x] [P0-I] Identify exact current coverage gap shape from code. Result: the dashboard can enumerate graph-present node labels and edge types from generated contracts, but graph adapter behavior is still hand-classified for selected structural/symbol node types, selected hierarchy edges, selected edge size multipliers, and selected provider parity fixtures. The plan must fix or document those concrete gaps, not just say "multi-language".

## Review Direction

This plan should solve the real problem only if it keeps these conclusions separate:

- `Class=3` and `Constructor=3` are not a graph bug unless later parser-based inventory proves additional class declarations were missed.
- `EXTENDS` + `INHERITS` duplication is a semantics/display contract problem.
- Missing TypeScript heritage is the first observed analyzer coverage/resolution problem, not the full language boundary of the fix.
- Go embedded structs currently appear as `EXTENDS`/`INHERITS`; the plan must decide whether that user-facing label is correct or whether the graph needs a more precise heritage kind/display label.
- Every supported language must be audited through a matrix so unsupported or non-applicable graph facts are explicit rather than silently omitted.
- All Web UI graph filters must be generated, displayed, counted, and toggled from the full graph contract, not from a TS/Go-shaped subset.
- Code paths already identified by AVmatrix must drive implementation order: `internal/contracts/web_ui.go`, `internal/analyze/analyze.go`, `internal/providers/*`, `internal/resolution/*`, `avmatrix-web/src/lib/constants.ts`, `avmatrix-web/src/components/FileTreePanel.tsx`, `avmatrix-web/src/lib/graph-adapter.ts`, `avmatrix-web/src/lib/graph-edge-visibility-mode.ts`, `avmatrix-web/src/hooks/useSigma.ts`, and their tests.

The implementation must trace each missing or contract-relevant source site through the pipeline before changing behavior:

```text
source AST -> ScopeIR facts -> workspace resolution -> graph node/relationship -> Web payload -> filter/dashboard/canvas display
```

Regex source-site counts are useful baseline evidence only. Final accuracy checks must use parser/ScopeIR or graph-internal inventory where practical.

## Phase 1 - Graph Filter And Language Contract Decision

- [x] [P1-A] Define the full filter contract for node labels: dashboard row, count source, default visibility, color legend, size policy, graph adapter behavior, and unknown-label fallback. Result: `NODE_LABELS` remains generated, graph-present labels are dashboard/legend rows, structural/community/mass classification is explicit in Web constants, and unknown labels retain fallback color/size behavior.
- [x] [P1-B] Define the full filter contract for relationship types: dashboard row, count source, default visibility, color/label legend, duplicate/semantic grouping policy, graph adapter behavior, and unknown-edge fallback. Result: `GRAPH_RELATIONSHIP_TYPES` remains generated, relationship display policy is generated, dashboard/legend counts use semantic display counts, and unknown edge types retain fallback style behavior.
- [x] [P1-C] Define whether `INHERITS` is a compatibility/normalized edge, a first-class display edge, or both. Result: raw graph keeps `INHERITS` as normalized/compatibility heritage; Web display labels it `Normalized Heritage` and groups it with same source-target `EXTENDS`/`IMPLEMENTS`.
- [x] [P1-D] Define canonical Web display behavior for pairs that have both `EXTENDS` and `INHERITS`. Result: dashboard counts and graph adapter rendering collapse duplicate compatibility `INHERITS` for the same source-target pair when `EXTENDS` or `IMPLEMENTS` is also present, while preserving raw counts in titles/tooltips.
- [x] [P1-E] Define raw graph export behavior: keep both edge types, collapse at graph payload generation, or preserve raw edges and add semantic grouping metadata. Result: raw graph payload preserves both edge types for compatibility; Web contracts now expose relationship display policy metadata.
- [x] [P1-F] Define unresolved/external source fact policy for targets such as `Error`, `Component`, DOM/React interfaces, external package types, missing imports, framework symbols, and scan-only language surfaces. Result: MCP schema now documents that resolved in-repo targets emit graph relationships while unresolved/external targets are not synthesized as graph nodes and remain audit/metrics evidence.
- [x] [P1-G] Define user-facing terminology for language-specific graph forms: class extends, interface extends, implements, Go embedding, trait/include/extend/prepend, namespace/module/package/struct/record/delegate/annotation, and normalized inheritance. Result: MCP schema now records language heritage terminology for TS, Go, Java/C#/Kotlin/Dart/PHP, and Python/C++/Swift/Ruby/Rust provider-specific forms.
- [x] [P1-H] Define compatibility behavior for MCP/context/impact/MRO consumers that may currently depend on raw relationship types such as `INHERITS`. Result: raw graph keeps `INHERITS` beside resolved `EXTENDS`/`IMPLEMENTS`; user display groups duplicate same-pair compatibility edges, while MCP/context/impact/MRO continue to see raw compatibility relationships.
- [x] [P1-I] Define the supported-language graph coverage matrix for `javascript`, `typescript`, `python`, `java`, `c`, `cpp`, `csharp`, `go`, `ruby`, `rust`, `php`, `kotlin`, `swift`, `dart`, `vue`, `svelte`, `astro`, and `cobol`. Result: generated Web contract includes `LANGUAGE_GRAPH_COVERAGE` for all 18 code languages.
- [x] [P1-J] For each supported language, classify every graph fact family as one of: resolved graph node/relationship, extracted but unresolved/external, extraction-only inventory, not applicable by language semantics, or scanned but not extracted. Result: generated `LANGUAGE_GRAPH_COVERAGE` now includes explicit per-language `factFamilies`, `supportedNodeLabels`, `supportedRelationshipTypes`, proof level, default regression gate, and optional audit metadata; generic `heritage-where-language-supports-it` was removed.
- [x] [P1-K] Define the COBOL policy explicitly: it is scanned and has a dedicated analyzer phase/metrics, but it is not currently routed through ScopeIR provider extraction by `hasExtractor`; graph filter behavior must reflect that instead of implying full ScopeIR parity. Result: generated coverage matrix marks COBOL as `dedicated-analyzer-phase`, not provider-backed ScopeIR parity.
- [x] [P1-L] Record the selected policy in plan, benchmark, evidence, MCP graph schema docs, generated Web contracts, and Web dashboard wording. Result: policy is recorded in this plan, benchmark/evidence ledgers, MCP schema resource, generated Web contracts, and Web dashboard relationship display wording.

## Phase 2 - Root-Cause Trace Before Fixing

- [x] [P2-A] For representative `E:\Restaurant_manager` TS files, capture parser/ScopeIR output and verify whether `HeritageFact` is emitted for every audited `extends` / `implements` site. Result: optional TS provider trace test against `AVMATRIX_RESTAURANT_MANAGER_ROOT=E:\Restaurant_manager` verifies `17` HeritageFact target facts across `13` audited TS/TSX files covering all `16` audited source sites.
- [x] [P2-B] Trace same-file TS interface inheritance such as `AreaWithTableCount extends Area` and `ShiftWithCounts extends Shift` through workspace binding and relationship emission. Result: loss point was workspace name resolution ambiguity for same-file interface targets; fixed with same-file heritage fallback.
- [x] [P2-C] Trace cross-file/imported TS heritage targets and record whether import binding or name resolution is the failure point. Result: the audited resolved TS heritage pairs are same-file targets; imported/built-in targets are React/DOM/Error/Performance external symbols and are classified as unresolved/external by policy, not import-binding failures to in-repo files.
- [x] [P2-D] Trace external TS heritage targets such as `Error`, `Component`, and React/DOM interfaces and classify them as external/unresolved by policy. Result: external TS heritage targets are classified as unresolved/external and are not synthesized as graph nodes in this slice.
- [x] [P2-E] Trace Go embedded struct relationships and classify whether their user-facing label should be `EXTENDS`, `EMBEDS`, `INHERITS`, or a grouped heritage display. Result: Go embedded struct/interface forms remain raw `EXTENDS` plus compatibility `INHERITS`, and Web display groups same-pair compatibility `INHERITS` as normalized heritage rather than an additional independent edge.
- [x] [P2-F] Record the exact loss point for each missing heritage source-site class: extraction loss, owner-scope loss, import/name-resolution loss, graph emission loss, payload loss, or UI display loss. Result: resolved in-repo TS sites were name-resolution losses; external React/DOM/Error/Performance sites remain unresolved by policy.
- [x] [P2-G] Trace representative heritage facts for every heritage-capable provider from parser/ScopeIR through resolution and graph emission, not only TS and Go. Result: provider parity tests now cover heritage extraction for TS, Go, Python, C++, Ruby, Java, C#, Kotlin, Rust, PHP, Dart, and Swift, and graph-resolution parity for TS, Go, Python, Java, C#, Kotlin, C++, PHP, Ruby, Rust, Dart, and Swift.
- [x] [P2-H] Trace representative non-heritage facts for every provider that supports them: definitions, imports, calls, uses/type refs, member/property/method relationships, accesses, routes/tools/process edges where applicable. Result: provider-specific golden ScopeIR fixtures and graph parity count tests cover definitions, imports, calls, accesses, type refs/bindings, members/properties/methods, and graph resolution for provider-backed languages plus Vue/Svelte/Astro script containers; route/tool/process edges remain covered by route/process/tool-specific suites rather than provider parity.
- [x] [P2-I] Audit `internal/analyze/analyze.go` `hasExtractor`, `extractScopeIR`, and `extractScriptContainerScopeIR` against `internal/contracts/web_ui.go` code languages so every declared language is categorized as provider-backed, script-container-backed, dedicated-phase-backed, scan-only, or unsupported. Result: 14 provider-backed languages, 3 script-container-backed languages, and COBOL dedicated analyzer phase are recorded in generated contract metadata.
- [x] [P2-J] Audit `internal/providers/provider_parity_test.go` and provider-specific tests to identify which graph fact families have real parity coverage and which are only covered for TS/Go or a small subset. Result: heritage parity gaps were expanded in `provider_parity_test.go`; broader non-heritage proof and wording are reopened under P8-E.
- [x] [P2-K] Compare actual graph node labels and relationship types from `E:\Restaurant_manager` and `E:\AVmatrix-GO` against Web filter rows, legend rows, color/size mappings, and adapter support. Result: benchmark ledger records final graph-present label/type inventories for both repos and Web unit tests assert graph-present dashboard completeness.
- [x] [P2-L] Audit `graph-adapter.ts` hard-coded classifications: `structuralTypes`, `symbolTypes`, `forwardHierarchyRelations`, `reverseHierarchyRelations`, `getNodeMass`, `EDGE_SIZE_MULTIPLIERS`, node size caps, and community coloring. Result: constants now expose explicit structural/community/edge-size classifications and unit tests assert every generated relationship type has a size policy.

## Phase 3 - Analyzer Coverage

- [x] [P3-A] Build the provider/fact-family matrix from `hasExtractor`, `extractScopeIR`, script-container extraction, Cobol analyzer-phase metrics, provider tests, generated Web contracts, and current graph payloads. Result: generated `LANGUAGE_GRAPH_COVERAGE` records language status and graph fact families for all 18 code languages.
- [x] [P3-B] For every supported language, record provider/dedicated analyzer status, supported node labels, supported relationship types, source fact families, unresolved/external behavior, and current fixture coverage. Result: all 18 code languages now have explicit generated coverage entries with per-fact status, node/relationship inventory, fixture coverage, and proof-level metadata.
- [x] [P3-C] Add or update provider parity fixtures for every claimed graph fact family before marking any language, node label, edge type, or filter complete. Result: existing provider fixtures are now documented as representative endpoint/count-level parity, and endpoint assertions were added for representative non-TS/Go graph facts instead of overclaiming every endpoint.
- [x] [P3-D] Add graph/resolution contract tests proving each supported source fact kind maps to the selected graph node, relationship, display metadata, unresolved audit record, or explicit not-applicable classification. Result: contract tests require explicit status for every generated language fact family and prevent provider-backed languages from using the old generic heritage marker.
- [x] [P3-E] Add focused TS trigger fixtures for `class extends`, `class implements`, `interface extends`, same-file targets, cross-file imported targets, and external unresolved targets. Result: TS provider tests cover interface heritage extraction; resolution tests cover resolved interface inheritance, same-file ambiguity, and unresolved external generic heritage.
- [x] [P3-F] Add parser/ScopeIR source-site inventory assertions so every TS trigger heritage site is counted even if unresolved. Result: `HeritageFactsIndexed` and `UnresolvedInheritance` metrics now account for resolved and unresolved heritage facts.
- [x] [P3-G] Fix TS provider/resolution behavior so resolved in-repo class/interface heritage emits the selected graph relationship contract. Result: `extends_type_clause` is extracted and same-file ambiguous targets resolve; Restaurant_manager now emits 8 resolved TS interface heritage pairs.
- [x] [P3-H] Fix or document unresolved/external heritage handling so those source sites are visible in graph audit data. Result: unresolved heritage is counted in binding metrics; external target graph nodes are not synthesized in this slice.
- [x] [P3-I] Verify Go embedded struct behavior still emits the expected heritage facts after any contract changes. Result: provider graph-resolution parity asserts Go embedded struct `Dog -> Animal` emits both `EXTENDS` and compatibility `INHERITS`.
- [x] [P3-J] Add coverage for missing or weak non-TS/Go provider facts before claiming a filter is multi-language complete, including script-container providers where facts come from embedded JS/TS. Result: non-TS/Go provider and script-container evidence is tied to explicit fact-family classifications; representative endpoint proof was added for C and Java definitions, members, calls, accesses, and type-use edges.
- [x] [P3-K] Update generated contracts or schema docs if the selected policy adds relationship metadata, external target facts, display-group fields, or per-language graph coverage metadata. Result: Web contract schema/generated TS now include relationship display policy and language graph coverage metadata.

## Phase 4 - UI Graph Filters, Display, And Shell Ergonomics

- [x] [P4-A] Verify and fix `FileTreePanel` node filter inventory so graph-present labels from `getFilterableNodeLabelsForGraph` are listed, counted, toggleable, colored, and mirrored in the legend. Result: existing graph-present node completeness coverage remains passing in full Web unit suite.
- [x] [P4-B] Verify and fix `FileTreePanel` edge filter inventory so graph-present relationship types from `getFilterableEdgeTypesForGraph` are listed, counted, toggleable, colored, and mirrored in the legend. Result: edge rows now use semantic display counts and full graph-present relationship dashboard coverage remains passing.
- [x] [P4-C] Decide whether zero-count contract labels/types should be shown in loaded-graph mode. If hidden, evidence must prove absence is from graph payload; if shown, counts must be `0` and toggles must be harmless. Result: loaded-graph mode hides zero-count contract labels/types and shows only graph-present labels/types; no-graph fallback still exposes generated contract rows with zero counts.
- [x] [P4-D] Update dashboard Edge Types counts so compatibility duplicates do not mislead users. Result: `FileTreePanel` uses `getDisplayRelationshipTypeCounts` and surfaces grouped/raw counts in titles.
- [x] [P4-E] Update graph canvas edge conversion so duplicate compatibility heritage pairs do not draw as two unrelated relationships unless explicitly requested. Result: `knowledgeGraphToGraphology` converts display relationships and collapses duplicate compatibility `INHERITS`.
- [x] [P4-F] Replace or justify graph-adapter TS/Go-biased hard-coded type sets: structural types, symbol types, mass, hierarchy parent-child relations, edge size multipliers, node size caps, and community coloring must cover all contract labels/types or have documented fallback behavior. Result: generalized constants cover generated relationship types and use explicit structural/community node classifications with fallback behavior.
- [x] [P4-G] Update color legend labels/tooltips for all node labels and relationship types, including `EXTENDS`, `IMPLEMENTS`, `INHERITS`, and any Go embedding display label, to reflect the selected semantics. Result: `INHERITS` display label is `Normalized Heritage`; grouped/raw title text explains compatibility grouping.
- [x] [P4-H] Add unit tests for duplicate `EXTENDS` + `INHERITS` source-target pairs. Result: constants, FileTreePanel, and graph-adapter unit tests cover duplicate heritage grouping/collapse.
- [x] [P4-I] Add unit tests for resolved and unresolved/external graph fact display across representative language-specific forms, including TS class/interface heritage and non-TS/Go provider facts. Result: constants/FileTreePanel tests assert generated language graph coverage unresolved/external policy, normalized heritage display policy, graph-present loaded-mode rows, zero-count hiding, and focus-depth warning behavior.
- [x] [P4-J] Add e2e coverage proving the Web UI displays node filters, edge filters, counts, legends, focus-depth behavior, and relationship toggles according to the selected policy. Result: Playwright shell e2e now opens Filters, verifies node/edge sections, File and Calls counts, legend rows, Calls toggle state, and focus-depth warning/clear behavior.
- [x] [P4-K] Add deterministic fixture coverage and, where feasible, a `Restaurant_manager` large-graph smoke check so the real regression is covered. Result: Playwright now selects `E2E_REPO_NAME` or the default `Restaurant_manager` by stable repo name/path and skips with an explicit reason if unavailable; focused e2e passed `3/3` against `Restaurant_manager`.
- [x] [P4-L] Inspect `avmatrix-web/src/components/Header.tsx`, `avmatrix-web/src/App.tsx`, launcher path handling, and `Start-AVmatrix.html` to confirm the correct Start-screen return target in both dev and packaged launcher modes. Result: Vite dev/build now serves/emits `Start-AVmatrix.html`; Header computes an absolute same-origin start-screen URL.
- [x] [P4-M] Add an icon-first Back arrow/button beside the `AVmatrix` top bar title with an accessible label and keyboard/click activation. Result: Header renders an `ArrowLeft` icon button labelled `Back to Start screen`.
- [x] [P4-N] Implement the return flow to `Start-AVmatrix.html` without showing a stale connection-loss banner during the intentional navigation transition. Result: App tracks intentional start navigation and suppresses the reconnect banner during that transition.
- [x] [P4-O] Inspect the app shell layout, `FileTreePanel` container, graph canvas sizing, and responsive CSS to identify where the left dashboard width is currently fixed. Result: `FileTreePanel` owns the left dashboard width and now exposes bounded runtime resizing.
- [x] [P4-P] Add a visible drag handle on the right edge of the left dashboard and support mouse/pointer dragging to resize the panel within bounded min/max widths. Result: drag handle persists width in local storage with `192px` min and `480px` max.
- [x] [P4-Q] Add unit and/or e2e coverage proving Back navigation, left dashboard resize bounds, and continued dashboard/canvas interaction after resize. Result: Vitest covers the Header handler and resize bounds; Playwright covers Back navigation and resize/canvas usability.

## Phase 5 - Benchmarks and Accuracy Measurement

- [x] [P5-A] Analyze `E:\AVmatrix-GO` and record full graph filter inventory: node label counts, relationship type counts, graph-present filter rows, adapter classification coverage, semantic unique heritage counts, and duplicate compatibility pair counts.
- [x] [P5-B] Analyze `E:\Restaurant_manager` and record the same full graph filter inventory plus the Restaurant_manager trigger heritage metrics.
- [x] [P5-C] Record TypeScript heritage source-site coverage for `E:\Restaurant_manager`: parser/ScopeIR source sites, resolved graph relationships, unresolved/external sites, and missing sites.
- [x] [P5-D] Record Go embedded struct heritage coverage and final user-facing display label/counts.
- [x] [P5-E] Record supported-language graph coverage matrix: language, provider/extractor status, supported node labels, supported relationship types, extraction status, resolution status, graph relationship/display policy, and fixture/e2e evidence. Result: benchmark ledger records 18 language coverage entries and 141 explicit fact-family rows across provider-backed, script-container-backed, and COBOL dedicated analyzer coverage.
- [x] [P5-F] Record UI filter inventory from code and runtime: visible node rows, visible edge rows, raw graph counts, displayed semantic counts, legend rows, focus-depth behavior, edge visibility behavior, graph adapter classification coverage, and graph canvas duplicate-pair rendering behavior.
- [x] [P5-G] Record Web UI shell interaction inventory for Back navigation and left dashboard resize: navigation target, stale connection-loss banner behavior, min/max width bounds, and canvas usable width after resize.
- [x] [P5-H] Update benchmark and evidence ledgers immediately after each measured slice.

## Phase 6 - Validation

- [x] [P6-A] Run full Go build before tests. Result: `go build ./cmd/... ./internal/...` passed.
- [x] [P6-B] Run focused Go tests for extraction/resolution and graph contract behavior across every covered language/provider. Result: focused provider/resolution/contracts tests passed, including optional Restaurant_manager TS heritage trace when `AVMATRIX_RESTAURANT_MANAGER_ROOT` is set.
- [x] [P6-C] Run focused tests for MCP/context/impact/MRO behavior if `INHERITS` compatibility or graph payload semantics change. Result: `go test ./internal/providers ./internal/mcp ./internal/mro ./internal/resolution` passed after documenting MCP raw/display policy and expanding provider heritage parity.
- [x] [P6-D] Run full applicable Go test suite for `cmd` and `internal`. Result: `go test ./cmd/... ./internal/...` passed. Repository-wide `go test ./...` still includes intentionally non-buildable fixture folders under `avmatrix/test/fixtures`.
- [x] [P6-E] Run Web build before Web tests. Result: `npm run build` passed.
- [x] [P6-F] Run focused Web unit tests for node type dashboard, edge type dashboard, legend, filter visibility, and graph adapter behavior. Result: focused constants/FileTreePanel/graph-adapter/Header tests passed.
- [x] [P6-G] Run full Web unit suite. Result: `npm test -- --run` passed with `41` files and `325` tests.
- [x] [P6-H] Run focused Web unit/e2e tests covering top bar Back navigation and left dashboard resize behavior. Result: Playwright `shell-interactions.spec.ts` focused grep passed `3/3` for Back, resize, and graph filter/legend/focus-depth coverage.
- [x] [P6-I] Run e2e test covering node filters, edge filters, legend, focus depth, duplicate-pair behavior, Back navigation, and left dashboard resize behavior. Result: duplicate-pair unit behavior is covered, and focused Playwright e2e covers node/edge filter sections, legend rows, focus-depth warning/clear behavior, Calls relationship toggling, Back navigation, and left dashboard resize.
- [x] [P6-J] Re-run analyze on `E:\AVmatrix-GO` and `E:\Restaurant_manager` and verify final graph/filter metrics match the accepted contract. Result: both repos were analyzed through `go run ./cmd/avmatrix analyze ... --force --skip-agents-md --no-stats`; final counts are in the benchmark ledger.
- [x] [P6-K] Verify the final supported-language graph coverage matrix has no unclassified language entries and no UI filter row disconnected from graph payload reality. Result: contract tests validate explicit fact-family status for all language entries; full Go/Web build/tests and focused deterministic e2e passed.

## Phase 7 - Closure

- [x] [P7-A] Update this plan checklist after each completed slice.
- [x] [P7-B] Update benchmark ledger with initial, intermediate, and final counts.
- [x] [P7-C] Update evidence ledger with commands, files changed, tests, and conclusions.
- [x] [P7-D] Commit each completed implementation slice. Result: completed implementation slices are committed as they close, including the graph heritage display slice and the current MCP policy/provider parity slice.
- [x] [P7-E] Final closure: confirm graph filter completeness, graph semantics, full supported-language coverage matrix, UI relationship display, Back navigation, left dashboard resize, benchmark, evidence, full build, unit tests, and e2e tests are complete. Result: zero-trust follow-up closure recorded after explicit language coverage, deterministic e2e, default TS heritage fixture, endpoint-level representative provider proof, updated benchmark/evidence, and full validation.

## Phase 8 - Zero-Trust Reopen Follow-Up

- [x] [P8-A] Replace or supplement generic `providerCoverage()` metadata with an explicit per-language graph fact matrix covering node labels, relationship types, extraction, resolution, unresolved/external, extraction-only, not-applicable, scanned-not-extracted, and fixture/e2e status. Result: `LanguageGraphCoverage` now exposes explicit fact-family rows and language-specific node/relationship inventories.
- [x] [P8-B] Add contract tests verifying no provider-backed language inherits a generic fact-family list without explicit supported, not-applicable, unresolved/external, extraction-only, or scanned-not-extracted statuses. Result: `TestWebUIContractManifestUsesGoRuntimeConstants` now enforces explicit fact coverage and rejects the generic heritage marker.
- [x] [P8-C] Make the `Restaurant_manager` large-graph e2e deterministic by selecting a repo by name/path or provisioning a deterministic large fixture; if unavailable, skip with an explicit reason instead of relying on `repos[0]`. Result: `shell-interactions.spec.ts` selects `E2E_REPO_NAME` or `Restaurant_manager` by stable identity and focused e2e passed against that repo.
- [x] [P8-D] Convert the `Restaurant_manager` TypeScript heritage trace from optional audit-only coverage to a deterministic default regression guard, or add a committed fixture reproducing all `17` `HeritageFact` target facts and document the external trace separately. Result: `TestExtractRestaurantManagerTypeScriptHeritageFixture` covers all `17` target facts by default; the external source trace remains optional and passed in this environment.
- [x] [P8-E] Strengthen provider graph parity proof with endpoint-level assertions for representative non-TS/Go fact families, or downgrade plan/benchmark/evidence wording from "every source fact kind" to representative/count-level coverage where appropriate. Result: `TestProviderGraphParityEndpointProofCoversRepresentativeNonTSGoFacts` asserts C/Java endpoints and generated coverage wording records representative endpoint/count-level proof.
- [x] [P8-F] Clean stale status/drift in benchmark and evidence ledgers: top-level statuses, superseded final closure text, historical pending rows, and final metrics must clearly distinguish baseline history from active requirements. Result: benchmark/evidence statuses and superseded sections were updated; historical baseline `pending` rows are treated as baseline history only.
- [x] [P8-G] Re-run required validation after follow-up: Go build/tests, Web build/tests, deterministic e2e, optional external trace if still used, and update benchmark/evidence ledgers with the results. Result: Go build/tests, Web build/tests, focused deterministic Playwright e2e, and optional external trace passed.
- [x] [P8-H] Final zero-trust closure: confirm no unchecked active tasks, no stale "pending" closure language except historical baseline tables explicitly labeled, deterministic regression gates are in place, and the follow-up implementation slice is committed. Result: active checklist is complete and the follow-up slice is included in the closing commit.
