# Anvien Orphan Node Connectivity Lens Plan

Date: 2026-05-20

Status: complete

Companion files:

- Benchmark ledger: [2026-05-20-anvien-orphan-node-connectivity-lens-benchmark.md](2026-05-20-anvien-orphan-node-connectivity-lens-benchmark.md)
- Evidence ledger: [2026-05-20-anvien-orphan-node-connectivity-lens-evidence.md](2026-05-20-anvien-orphan-node-connectivity-lens-evidence.md)

## Rules

1. Use Anvien for codebase analysis and impact checks while working on implementation slices in this plan.
2. As each task is completed, update the corresponding checklist item immediately.
3. Run a full build before testing; the test suite must include an e2e test if Web UI behavior changes.
4. Record benchmark results as each benchmarkable task is completed. Benchmarkable means measured product/runtime performance, capacity, package/startup size, graph/DB throughput, or graph inventory counts; build/test/e2e timings are validation evidence unless the slice changes those systems.
5. Record evidence as each evidenced task is completed.
6. For doc-only commits, do not use Anvien.
7. After each completed implementation slice, commit the work, then continue until the full plan is complete.

## Problem

Anvien graphs can contain nodes that appear disconnected or under-connected. Users need a way to identify code that may be dead, unwired, missing analyzer edges, unresolved, or intentionally isolated.

The term "orphan node" is currently ambiguous. A node with no visible edges can mean many different things:

- real dead code or an unused symbol;
- a feature module not wired into routes, commands, tools, sessions, or process flows;
- an exported API, framework entrypoint, route handler, CLI command, reflection target, migration hook, or public surface that has no static caller;
- a test helper, fixture, generated file, vendor file, documentation section, or example that is expected to be isolated;
- a parser, provider, resolution, graph emission, or Web payload gap that failed to create the expected edges;
- a Web filter/display state that hides relationships and makes a connected node look isolated.

The plan must turn this ambiguous topology symptom into a precise graph-health workflow without changing semantic node labels such as `Function`, `File`, `Class`, `Interface`, `Struct`, `Method`, or `Section`.

This lens identifies graph-health candidates and explanations. It is not a bug detector by itself, and it must not label code as confirmed dead, buggy, or removable without follow-up evidence from source, runtime/route/export context, tests, or user review.

## Scope Boundary

Implementation may touch:

- graph inventory and graph-health analysis code that derives connectivity status from graph payloads;
- relationship type policy for which edges count as connectivity;
- analyzer/resolution metrics if source facts need explicit unresolved or missing-edge evidence;
- graph accuracy audit code if existing property-specific orphan terminology must be namespaced or bridged;
- internal contracts that expose graph-health metadata to Web consumers;
- Web generated contracts if graph-health filters become generated metadata;
- `anvien-web` graph filters, dashboard, node detail panel, legends, and e2e coverage;
- MCP/query/report surfaces if graph-health summaries are exposed outside Web UI;
- tests and fixtures for connectivity taxonomy, expected-isolated policy, and Web filter behavior.

Out of scope unless a later phase explicitly reopens it:

- deleting or rewriting application code based only on orphan status;
- changing primary semantic node labels to `OrphanNode`;
- treating zero outgoing edges as a defect by default;
- synthesizing fake external target nodes only to reduce orphan counts;
- claiming dead-code accuracy without source, route/export, test/generated/vendor, and analyzer evidence;
- calling a connectivity candidate a confirmed bug, confirmed dead code, or safe deletion target without separate supporting evidence.

## Design Decision

Do not map "orphan node" into a primary node type.

Use derived graph-health metadata as the source of truth, then expose it through consumer surfaces such as API/reporting and Web UI filters. The Web `Graph Health` filter group is a consumer of backend-derived metadata, not the place where graph-health truth is invented.

```text
Node semantic label: Function, File, Class, Interface, Struct, Method, Section, ...
Graph Health metadata:
  topologyStatus: connected | true_isolated | no_incoming | no_outgoing | detached_component | unknown_connectivity
  expectedIsolationReasons: []
  diagnostics: []
  confidence: candidate | expected | unknown | confirmed
Primary UI group: Graph Health
```

The proposed initial topology taxonomy is:

| Status | Definition | Default interpretation |
|---|---|---|
| `connected` | one or more counted incoming edges and one or more counted outgoing edges, and not detached from accepted roots under the component policy | normal graph-health baseline; usually not shown as a triage filter |
| `true_isolated` | zero counted incoming edges and zero counted outgoing edges | needs triage; could be real dead artifact, expected fixture/generated/vendor, or analyzer miss |
| `no_incoming` | zero counted incoming edges and one or more counted outgoing edges | strongest dead-code or unwired candidate after entrypoint/export/test/generated exclusions |
| `no_outgoing` | one or more counted incoming edges and zero counted outgoing edges | often normal leaf behavior; inspect only, do not flag as bug by default |
| `detached_component` | a connected component has internal edges but no counted path to accepted entry/process/root surfaces | strong candidate for unwired feature or missing root edge |
| `unknown_connectivity` | graph does not contain enough evidence to classify safely | do not count as bug |

Phase 1 must either accept this taxonomy as final or revise it with recorded rationale, tests, evidence-ledger notes, and benchmark implications before any user-facing status is implemented.

Expected-isolated classification is not a topology status. It is an overlay reason list that can apply to any topology status, for example a test helper can be `no_incoming` with `expectedIsolationReasons: ["test"]`. Public/exported status alone is a prioritization modifier, not sufficient evidence to auto-hide a node.

Unresolved references are diagnostics, not node topology statuses, unless a source fact can be attached to a specific in-repo source node. When source/resolution evidence exists, emit a diagnostic such as `unresolved_reference` with fact family, source node, target text, and resolution source. When no source-node evidence exists, include the count in summaries and classify affected topology as `unknown_connectivity` rather than pretending there is a filterable node status.

Existing property-access accuracy output is a separate diagnostic surface. `cmd/property-access-audit` and `internal/graphaccuracy/property_access.go` already emit a Property-only `orphanStatus` taxonomy such as `owner_linked`, `false_orphan`, `true_orphan`, `unknown`, `external_library_owned`, and `intentionally_unmodeled`. The Graph Health taxonomy must stay namespaced from that audit output unless a phase explicitly defines and tests a compatibility mapping.

## Connectivity Edge Policy

Phase 1 must define the exact relationship types counted for connectivity.

The policy must separate:

- code dependency edges such as calls, definitions, imports, type references, accesses, members, ownership, routes, tools, processes, and inheritance;
- structural/container edges such as file/package/directory containment;
- documentation/report edges;
- display-only or compatibility edges;
- hidden edges suppressed by Web filters.
- root/path traversal rules for `detached_component`, including directed versus undirected traversal, accepted root labels, accepted entry/process/resource surfaces, and how `ENTRY_POINT_OF`, `STEP_IN_PROCESS`, `HANDLES_ROUTE`, and `HANDLES_TOOL` directions are interpreted.

No count is valid until the edge policy is recorded in this plan, the evidence ledger, and tests.

## Expected-Isolated Policy

Phase 1 must define exclusion rules before any status is presented as a bug candidate.

The expected-isolated policy must classify at least:

- test files and test helpers;
- fixtures and sample repositories;
- generated files;
- nodes from vendor/dependency paths if such nodes are present in the graph;
- documentation/report/section nodes;
- migrations and scripts;
- exported APIs and public package surfaces, with an explicit evidence rule that exported-only is a prioritization modifier and not automatic expected-isolated status;
- framework entrypoints and route handlers only when route/tool/process/framework evidence identifies them as entry surfaces;
- CLI commands, MCP tools, session handlers, background jobs, and reflection/config-discovered surfaces.
- scanner-ignored paths as out-of-graph inputs, not Graph Health status counts.

## Acceptance Criteria

- "Orphan" is not introduced as a primary node label.
- Backend graph-health derivation is the source of truth for topology status, expected-isolated overlays, diagnostics, and confidence.
- At least one non-Web consumer surface exposes graph-health/connectivity summaries and explanations before Web UI filters are treated as complete.
- Web UI exposes graph-health/connectivity filters separately from node-type filters when the Web phase is implemented.
- Every flagged node has an explanation: topology status, counted incoming edges, counted outgoing edges, excluded edge categories, source path policy, expected-isolated reasons, diagnostics, and confidence.
- `no_incoming` and `detached_component` are prioritized as actionable candidates; `no_outgoing` is not presented as a bug by default.
- Graph Health statuses are candidate/explanation signals, not confirmed bug or deletion verdicts.
- Expected-isolated nodes can be hidden or de-emphasized without deleting their semantic labels or overwriting their topology status.
- A user can distinguish a real candidate from analyzer/resolution uncertainty.
- A node can simultaneously retain topology status, expected-isolated reasons, and diagnostics without those fields overwriting each other.
- Benchmark ledger records measured counts for all topology statuses, expected-isolated reasons, diagnostics, and confidence levels on `E:\Anvien` and on representative indexed repos selected by documented Phase 1 criteria when available.
- Evidence ledger records commands, impacted files, tests, e2e artifacts, and conclusions for each implementation slice.
- Full build, unit tests, and relevant e2e tests pass before closure.

## Baseline Requirements

The following baseline must be measured before implementation claims:

- current graph node count by semantic label;
- current relationship count by type;
- count of nodes with zero incoming counted edges;
- count of nodes with zero outgoing counted edges;
- count of nodes with zero counted edges both ways;
- connected component count and largest detached component candidates;
- expected-isolated count by reason;
- unresolved reference count by source fact family if available;
- comparison between raw graph connectivity and Web-visible connectivity after filters.
- compatibility impact on existing property-access `orphanStatus` output.
- graph source timestamp or hash and exact count commands/scripts sufficient to reproduce every baseline row.

Initial `E:\Anvien` codebase-reviewed baseline measurements are recorded in the benchmark ledger. Phase 1 later locked the counted-edge and expected-isolated policies, selected `Restaurant_manager` as the representative cross-repo baseline, and recorded accepted-policy counts before implementation.

## Codebase Findings Before Implementation

Anvien index and source inspection on 2026-05-20 identified these existing implementation facts:

- Graph storage currently has no graph-health metadata. `internal/graph/types.go` defines `Node`, `Relationship`, and `Graph`; relationships have `sourceId`, `targetId`, `type`, confidence, reason, step, resolution source, file hash, and evidence, but nodes have no `topologyStatus`, expected-isolated reasons, diagnostics, or per-node incoming/outgoing summary.
- The HTTP graph endpoint currently streams only raw graph nodes and relationships. `internal/httpapi/graph.go` exposes `graphPayload` and `streamGraphNDJSON`; `graphNodeForResponse` only strips `content` unless requested.
- Generated Web contracts currently expose semantic node labels and relationship types, not Graph Health metadata. `internal/contracts/web_ui.go` defines `nodeLabels`, `graphRelationshipTypes`, relationship display policy, and language coverage metadata.
- The Web graph model currently mirrors generated `GraphNode` and `GraphRelationship` arrays. `anvien-web/src/core/graph/types.ts` defines `KnowledgeGraph` as nodes plus relationships with no derived graph-health field.
- Web filter UI currently has `Node Types`, `Edge Types`, `Focus Depth`, and `Color Legend` inside `FileTreePanel`. There is no `Graph Health` filter group.
- Web graph conversion already has an "orphan nodes" placement step in `knowledgeGraphToGraphology`, but that phrase only means nodes not reached by the layout hierarchy BFS. It is not a product taxonomy and does not mean dead or unwired code.
- Existing layout hierarchy logic in `knowledgeGraphToGraphology` treats `CONTAINS`, `HAS_METHOD`, `HAS_PROPERTY`, `DEFINES`, `IMPORTS`, `WRAPS`, `STEP_IN_PROCESS`, `ENTRY_POINT_OF`, `HANDLES_ROUTE`, `HANDLES_TOOL`, and `MEMBER_OF` as parent/child layout signals with priorities.
- Existing filter counts in `FileTreePanel` use graph-present semantic labels and relationship types; relationship display counts collapse grouped `INHERITS` compatibility edges through `getDisplayRelationshipTypeCounts`.
- Existing Web graph filter state is split across `FileTreePanel`, `GraphStateProvider`, `GraphCanvas`, `knowledgeGraphToGraphology`, and Sigma node/edge attributes. Graph Health filters must compose with node-type filters, edge-type visibility, focus-depth filtering, and graph canvas refresh behavior.
- Existing process detection already uses a narrower connectivity idea than raw graph connectivity. `internal/processes/processes.go` builds process traces from `CALLS`, ignores low-confidence calls, excludes test files for process entrypoints, links `Route`/`Tool` nodes to processes with `ENTRY_POINT_OF`, and scores exported/framework-like functions.
- Existing ignore policy already recognizes many expected-isolated path classes. `internal/ignore/constants.go` ignores directories such as `node_modules`, `vendor`, `dist`, `build`, generated directories, fixtures, snapshots, caches, and test fixture directories. `internal/processes/processes.go` also has `isTestFile` logic for `.test.`, `.spec.`, `/test/`, `/tests/`, `__tests__`, `_test.go`, and `_test.py`.
- Existing graph accuracy code already uses orphan terminology for one specialized audit. `internal/graphaccuracy/property_access.go` classifies Property nodes into `owner_linked`, `false_orphan`, `true_orphan`, `unknown`, `external_library_owned`, and `intentionally_unmodeled`; `cmd/property-access-audit` prints those counts as `orphan.*` summary lines. This is not a global graph-health/node-connectivity taxonomy.
- Existing generated Web contracts already document unresolved/external policy for provider fact coverage, but unresolved/external targets are retained in metrics/evidence rather than emitted as resolved graph edges.

Initial measurements also prove why a raw "zero incoming" filter would be wrong:

- With all relationship types counted, `Anvien` has `0` code nodes with zero raw incoming edges, because structural/ownership relationships such as `DEFINES` connect symbols from files.
- With a provisional non-structural policy that excludes structural/ownership/display grouping edges, `Anvien` has `1,616` code nodes with zero counted incoming edges.
- Therefore Phase 1 must close the counted-edge policy before any Graph Health status is user-facing.

## Key Decisions Required in Phase 1 (Checklist Form)

**All decisions finalized and recorded 2026-05-20 (see E5 in evidence ledger + B0 in benchmark for full tables/rationale).**

- Counted Edge Policy: 17 "wiring" rel types count; 5 structural (CONTAINS/DEFINES/HAS_*/MEMBER_OF) excluded. Rationale + exact list in benchmark B0.
- Expected-Isolated overlay: 8 automatic reasons (test/fixture/generated/vendor/doc/migration/cli-mcp) + 1 modifier (exported_api) + framework_entry as root. Evidence rules defined; bridges existing isTestFile + ignore + exported scoring.
- Root surfaces + traversal: Process + ENTRY_POINT_OF/HANDLES_* sources + Route/Tool + main-like; directed outgoing on counted edges. Details in E5.
- Ownership: core graph layer via internal/graphhealth annotation (for MCP consistency). Web consumes derived payload.
- Representative criteria + selection: large (>5k files), diverse label mix, entry-surface rich; Restaurant_manager chosen + measured (provisional on its snapshot).

All recorded with commands, rationale, cross-repo numbers. No ambiguity remains for Phase 2.

## Phase 1 - Taxonomy, Policy, and Baseline

**Status: COMPLETE 2026-05-20** (all items closed in this doc-only slice; see E5 evidence + B0 benchmark for full artifacts and rationale. No code changes.)

- [x] [P1-A1] Draft and document the initial Counted Edge Policy: 17 wiring types vs 5 structural. Table + rationale in benchmark B0.
- [x] [P1-A2] Define explicit rules for structural and ownership edges: they do NOT contribute to counted incoming/outgoing for topology (always-present ownership masks candidates). Matches graph-adapter hierarchy and empirical data.
- [x] [P1-A3] Finalize and record the accepted Counted Edge Policy with rationale and examples; updated benchmark B0 + evidence E5.
- [x] [P1-B1] Draft the Expected-Isolated overlay policy covering test, fixture, generated, vendor, documentation, migration, exported, route/tool/entry, cli/mcp, reflection. Full list + rules in E5.
- [x] [P1-B2] Define for each: automatic (test/fixture/etc.), prioritization modifier only (`exported_api`), framework_entry as root (never candidate). Evidence rules explicit.
- [x] [P1-B3] Finalize and record the accepted Expected-Isolated Policy with evidence rules. Bridges existing isTestFile + ignore + exported logic.
- [x] [P1-C] Define confidence levels: `candidate` / `expected` / `unknown` / `confirmed` with evidence requirements; topology independent of confidence and reasons. Recorded E5.
- [x] [P1-D] Measure baseline for `E:\Anvien` (fresh post-analyze 21091/52445) + reproducible python commands + interpretation. Recorded B0 + E5. Cross-repo also done.
- [x] [P1-E1] Define representative indexed-repo selection criteria (size, label diversity, entry-surface density, non-self). Recorded in B0.
- [x] [P1-E2] Select `Restaurant_manager` (78k nodes, 10k code, 505 processes) and measure provisional accepted-policy baseline (zero_in 4191 etc.). Recorded B0.
- [x] [P1-F] Record the selected topology taxonomy (unchanged from design), expected-isolated, diagnostics, property-access compatibility (namespaced, no conflict), Web wording (Graph Health filter group). All in E5 + plan.
- [x] [P1-G] Define root surfaces (Process, ENTRY_POINT_OF sources, Route/Tool, main-like) + directed outgoing traversal on counted edges for detached_component. Full rules E5.
- [x] [P1-H] Decide ownership: core graph layer (`internal/graphhealth` annotation for MCP consistency). Web/MCP are consumers. Recorded E5.
- [x] [P1-I] Document all Phase 1 decisions in the plan (this section + design), evidence E5, benchmark B0 with clear rationale, commands, and numbers from both repos.

## Phase 2 - Backend Graph-Health Derivation

**Slice 2026-05-20:** backend/API/contract derivation implemented for counted-edge topology, expected-isolated overlays, confidence, excluded structural counts, JSON summary, NDJSON node metadata, and generated Web contract types.

**Detached-component slice 2026-05-20:** backend derivation now computes weak counted-edge components, accepted root surfaces, directed root reachability, per-node component metadata, and largest detached component summaries.

**Unresolved-diagnostics slice 2026-05-20:** resolution now emits source-backed `unresolved_reference` diagnostics only when a source node can be identified, preserves unresolved counts in graph metadata, classifies diagnostic-backed nodes as `unknown_connectivity`, aggregates repeated diagnostics for payload control, and strips the internal raw diagnostic property from HTTP graph payloads.

**Coverage slice 2026-05-20:** P2-F adds graph-health test coverage across all topology statuses, all expected-isolated reasons, diagnostics aggregation/source attribution behavior, confidence transitions, and counted/excluded edge policy.

- [x] [P2-A] Identify the graph data boundary that should own derived graph-health metadata: graph package, analyzer output, HTTP graph payload, contract layer, or Web-only derived state. Decision implemented as `internal/graphhealth` core derivation consumed by HTTP/API and contracts.
- [x] [P2-B] Implement deterministic connectivity summary generation using the Phase 1 edge policy.
- [x] [P2-C] Add per-node derived metadata for topology status, counted incoming/outgoing counts, excluded edge counts by category, expected-isolated reasons, diagnostics, and confidence.
- [x] [P2-D] Add detached-component grouping and component-level explanations using explicit root/path traversal rules from Phase 1.
- [x] [P2-E] Add unresolved-reference diagnostics only where source/resolution evidence exists; otherwise preserve unresolved counts in summaries and classify affected topology as `unknown_connectivity`.
- [x] [P2-F] Add unit tests for every topology status, expected-isolated overlay reason, diagnostics rule, confidence rule, and exclusion rule.

## Phase 3 - Contract, API, and Reporting Surface

- [x] [P3-A] Decide whether graph-health metadata is emitted in Web graph payloads, generated Web contracts, MCP resources, or a dedicated endpoint. Decision for this slice: HTTP graph JSON response includes `graphHealth` summary; JSON and NDJSON node records include per-node `properties.graphHealth`; generated Web contracts define the stable types. MCP/report-specific explain surfaces remain later work.
- [x] [P3-B] Add graph-health summary output with counts by topology status, expected-isolated reason, diagnostics type, and confidence.
- [x] [P3-C] Add explain output for a single node or component. Implemented as GET `/api/graph/explain` with exactly one of `nodeId` or `componentId`; node explain returns health plus counted/excluded relationship evidence, and component explain returns aggregate counts plus bounded samples.
- [x] [P3-D] Add report/export path if needed for dead-code or unwired-candidate review. Implemented as GET `/api/graph/report` JSON export with `candidate_not_confirmed` verdict policy, triage priority ordering, `includeExpected=true`, and limit bounds.
- [x] [P3-E] Add contract tests proving status fields are stable, explicit, and not confused with semantic labels.

## Phase 4 - Web UI Graph Health Filters

**Filter-composition slice 2026-05-20:** Web state, dashboard controls, graph conversion, Sigma attributes, node-type filtering, focus-depth filtering, focused unit tests, full Web unit suite, Web build, and targeted Playwright dashboard e2e are implemented for Graph Health filters.

**Detail/focus slice 2026-05-20:** Dashboard explanatory tooltips, confidence counts, selected-node Graph Health explanations, next triage action, and detached-component focus/highlight interaction are implemented. Phase 4 Web UI scope is closed.

- [x] [P4-A] Add a `Graph Health` filter group separate from `Node Types` and `Edge Types`.
- [x] [P4-B] Add topology toggles for `true_isolated`, `no_incoming`, `no_outgoing`, `detached_component`, and `unknown_connectivity`; optionally show `connected` as a count-only baseline.
- [x] [P4-C] Add separate controls to hide/de-emphasize expected-isolated overlay reasons and diagnostics such as `unresolved_reference` when source-node evidence exists.
- [x] [P4-D] Add summary counts and tooltips that explain topology status, expected-isolated overlays, diagnostics, and confidence without calling candidates bugs.
- [x] [P4-E] Add node detail panel explanations: topology status, counted incoming/outgoing edges, expected-isolated reasons, diagnostics, confidence, and next triage action.
- [x] [P4-F] Add detached-component interaction that focuses a component and shows why it is detached.
- [x] [P4-G] Compose Graph Health filtering through Web state, `GraphCanvas`, `knowledgeGraphToGraphology`, Sigma node attributes, node-type filters, edge-type visibility, and focus-depth filtering.
- [x] [P4-H] Ensure existing node-type, edge-type, legend, focus-depth, graph links visibility, and graph canvas behavior still works with Graph Health filters.
- [x] [P4-I] Explicitly test and validate Graph Health filter composition with all existing filters (Node Types, Edge Types, Focus Depth) and layout hierarchy; record failure modes and guardrails.

## Phase 5 - Triage Workflow

**Status: COMPLETE 2026-05-20.** This phase was closed by prior implementation slices: P3-D report/export added the candidate review order and `candidate_not_confirmed` verdict policy; P4 filters added expected-isolated hiding; P4 detail/focus added in-product status explanations and next actions.

- [x] [P5-A] Define default triage order: `no_incoming` production symbols without expected-isolated reasons, `detached_component`, diagnostics such as source-backed `unresolved_reference`, `true_isolated`, then optional `no_outgoing`.
- [x] [P5-B] Add report wording for "candidate" versus "confirmed" findings.
- [x] [P5-C] Add documentation or in-product text for why a status was assigned and what the next action should be.
- [x] [P5-D] Add a way to hide or de-emphasize nodes by expected-isolated overlay reason without changing raw graph data or topology status.

## Phase 6 - Validation

- [x] [P6-A] Run full Go build before tests. Attempted `go build ./...`; blocked by existing non-buildable fixture packages under `anvien/test/fixtures/...`. `go build ./cmd/... ./internal/...` passed for applicable Go packages.
- [x] [P6-B] Run focused Go tests for graph-health derivation, taxonomy, expected-isolated policy, and contract/API behavior.
- [x] [P6-C] Run full applicable Go test suite for `cmd` and `internal`.
- [x] [P6-D] Run Web build before Web tests if Web UI changes. Current Web filter slice: `npm --prefix anvien-web run build` passed.
- [x] [P6-E] Run focused Web unit tests for Graph Health filters, node detail explanations, counts, legends, and filter interactions.
- [x] [P6-F] Run full Web unit suite if Web UI changes. Current Web filter slice: `npm --prefix anvien-web run test` passed.
- [x] [P6-G] Run e2e covering Graph Health filter visibility, node explanation, detached-component focus, and interaction with existing node/edge filters if Web UI changes. Targeted large-graph dashboard e2e and deterministic mocked Graph Health e2e passed; the earlier monolithic full-suite timeout remains recorded in E12 as suite-budget risk, not a feature failure.
- [x] [P6-H] Re-run baseline graph-health inventory after implementation and record before/after counts.
- [x] [P6-I] Validate performance impact of graph-health derivation (if done server-side) and Web filter rendering latency under realistic graph sizes. Server-side derivation and payload size measured in B1; Web filter/detail latency measured in E15/B2 on `Restaurant_manager`.
- [x] [P6-J] Validate that Graph Health status, expected-isolated reasons, diagnostics, and confidence can coexist on the same node without one overwriting the others.

## Phase 7 - Closure

- [x] [P7-A] Update this plan checklist after each completed slice.
- [x] [P7-B] Update benchmark ledger with initial, implementation, Web package-size, Web latency, and final graph-health inventory measurements.
- [x] [P7-C] Update evidence ledger with commands, impacted files, tests, e2e artifacts, benchmark commands, and conclusions for all slices.
- [x] [P7-D] Commit each completed implementation slice.
- [x] [P7-E] Final closure (after all phases).

## Ledger

| ID | Area | Scope | Target | Benchmark | Evidence | Commit | Status |
|---|---|---|---|---|---|---|---|
| P1-A1..P1-I | Policy | edge policy, expected-isolated policy, root rules, metadata ownership, taxonomy, baseline, cross-repo criteria | no ambiguous orphan claims; all major decisions recorded | 2026-05-20 (E5 + B0) | 2026-05-20 python + source + MCP queries | 2026-05-20 (doc-only) | closed |
| P2-A..P2-F | Backend | derived graph-health metadata | deterministic status and reasons | B1 implementation counts | E6/E7/E8/E9 implementation validation | current implementation slices | closed |
| P3-A..P3-E | Contract/API | consumer surface | stable explicit status fields | B1 payload size | E6/E10/E11 implementation validation | current implementation slices | closed |
| P4-A..P4-I | Web UI | graph-health filters + composition | separate filters, explanations, and safe composition with existing filters | B1/P4 package size notes | E12/E13 implementation validation | current implementation slices | closed |
| P5-A..P5-D | Workflow | triage/reporting | candidate-vs-confirmed workflow | no new benchmark; existing B1/P4 observations unchanged | E14 doc-only reconciliation | current doc-only slice | closed |
| P6-A..P6-H | Validation | build/tests/e2e | full validation recorded | B1/B2 | E6/E7/E8/E12/E13/E15 | current implementation slices | closed; applicable backend/Web-contract/Web UI validation passed, full `go build ./...` fixture blocker recorded |
| P7-A..P7-E | Closure | ledgers and commits | complete closure package | B2 final benchmark | E16 final closure | current doc-only slice | closed |

## Definition Of Done

- The plan, evidence, and benchmark ledgers are all updated.
- Graph Health model has explicit topology statuses, edge policy, expected-isolated overlays, diagnostics model, and confidence rules.
- No UI or API labels "orphan" as a bug without evidence.
- Graph Health filters are separate from semantic node labels.
- Node explanations are auditable from recorded graph data.
- Benchmarks record measured counts, not assumptions.
- Full build and applicable Go/Web/e2e validation pass.
- Anvien status is up to date after final implementation commit.
