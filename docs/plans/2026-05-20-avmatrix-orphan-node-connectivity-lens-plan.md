# AVmatrix Orphan Node Connectivity Lens Plan

Date: 2026-05-20

Status: active

Companion files:

- Benchmark ledger: [2026-05-20-avmatrix-orphan-node-connectivity-lens-benchmark.md](2026-05-20-avmatrix-orphan-node-connectivity-lens-benchmark.md)
- Evidence ledger: [2026-05-20-avmatrix-orphan-node-connectivity-lens-evidence.md](2026-05-20-avmatrix-orphan-node-connectivity-lens-evidence.md)

## Rules

1. Use AVmatrix for codebase analysis and impact checks while working on implementation slices in this plan.
2. As each task is completed, update the corresponding checklist item immediately.
3. Run a full build before testing; the test suite must include an e2e test if Web UI behavior changes.
4. Record benchmark results as each benchmarkable task is completed. Benchmarkable means measured product/runtime performance, capacity, package/startup size, graph/DB throughput, or graph inventory counts; build/test/e2e timings are validation evidence unless the slice changes those systems.
5. Record evidence as each evidenced task is completed.
6. For doc-only commits, do not use AVmatrix.
7. After each completed implementation slice, commit the work, then continue until the full plan is complete.
8. Do not mark orphan/connectivity findings as bugs without evidence. Connectivity status is a graph-derived candidate classification until source, analyzer, route, export, generated, test, or fixture evidence confirms the interpretation.
9. Do not create fake baseline counts. Any graph count, orphan count, detached-component count, or unresolved-reference count must come from a recorded command or test artifact.

## Problem

AVmatrix graphs can contain nodes that appear disconnected or under-connected. Users need a way to identify code that may be dead, unwired, missing analyzer edges, unresolved, or intentionally isolated.

The term "orphan node" is currently ambiguous. A node with no visible edges can mean many different things:

- real dead code or an unused symbol;
- a feature module not wired into routes, commands, tools, sessions, or process flows;
- an exported API, framework entrypoint, route handler, CLI command, reflection target, migration hook, or public surface that has no static caller;
- a test helper, fixture, generated file, vendor file, documentation section, or example that is expected to be isolated;
- a parser, provider, resolution, graph emission, or Web payload gap that failed to create the expected edges;
- a Web filter/display state that hides relationships and makes a connected node look isolated.

The plan must turn this ambiguous topology symptom into a precise graph-health workflow without changing semantic node labels such as `Function`, `File`, `Class`, `Interface`, `Struct`, `Method`, or `Section`.

## Scope Boundary

Implementation may touch:

- graph inventory and graph-health analysis code that derives connectivity status from graph payloads;
- relationship type policy for which edges count as connectivity;
- analyzer/resolution metrics if source facts need explicit unresolved or missing-edge evidence;
- internal contracts that expose graph-health metadata to Web consumers;
- Web generated contracts if graph-health filters become generated metadata;
- `avmatrix-web` graph filters, dashboard, node detail panel, legends, and e2e coverage;
- MCP/query/report surfaces if graph-health summaries are exposed outside Web UI;
- tests and fixtures for connectivity taxonomy, expected-isolated policy, and Web filter behavior.

Out of scope unless a later phase explicitly reopens it:

- deleting or rewriting application code based only on orphan status;
- changing primary semantic node labels to `OrphanNode`;
- treating zero outgoing edges as a defect by default;
- synthesizing fake external target nodes only to reduce orphan counts;
- claiming dead-code accuracy without source, route/export, test/generated/vendor, and analyzer evidence.

## Design Decision

Do not map "orphan node" into a primary node type.

Use derived connectivity metadata and a separate Web filter group:

```text
Node semantic label: Function, File, Class, Interface, Struct, Method, Section, ...
Derived status: connectivityStatus
UI group: Graph Health
```

The first accepted taxonomy is:

| Status | Definition | Default interpretation |
|---|---|---|
| `true_isolated` | zero counted incoming edges and zero counted outgoing edges | needs triage; could be real dead artifact, expected fixture/generated/vendor, or analyzer miss |
| `no_incoming` | zero counted incoming edges and one or more counted outgoing edges | strongest dead-code or unwired candidate after entrypoint/export/test/generated exclusions |
| `no_outgoing` | one or more counted incoming edges and zero counted outgoing edges | often normal leaf behavior; inspect only, do not flag as bug by default |
| `detached_component` | a connected component has internal edges but no counted path to accepted entry/process/root surfaces | strong candidate for unwired feature or missing root edge |
| `unresolved_reference` | source fact or relationship evidence names a target that cannot be resolved to an in-repo graph node | analyzer/resolution/import/external policy issue until classified |
| `expected_isolated` | node matches explicit expected-isolated policy such as fixture, generated, vendor, public API, test-only, docs, example, migration, or framework entrypoint | hidden or de-emphasized by default |
| `unknown_connectivity` | graph does not contain enough evidence to classify safely | do not count as bug |

## Connectivity Edge Policy

Phase 1 must define the exact relationship types counted for connectivity.

The policy must separate:

- code dependency edges such as calls, definitions, imports, type references, accesses, members, ownership, routes, tools, processes, and inheritance;
- structural/container edges such as file/package/directory containment;
- documentation/report edges;
- display-only or compatibility edges;
- hidden edges suppressed by Web filters.

No count is valid until the edge policy is recorded in this plan, the evidence ledger, and tests.

## Expected-Isolated Policy

Phase 1 must define exclusion rules before any status is presented as a bug candidate.

The expected-isolated policy must classify at least:

- test files and test helpers;
- fixtures and sample repositories;
- generated files;
- vendor/dependency directories;
- documentation/report/section nodes;
- migrations and scripts;
- exported APIs and public package surfaces;
- framework entrypoints and route handlers;
- CLI commands, MCP tools, session handlers, background jobs, and reflection/config-discovered surfaces.

## Acceptance Criteria

- "Orphan" is not introduced as a primary node label.
- Web UI exposes graph-health/connectivity filters separately from node-type filters.
- Every flagged node has an explanation: counted incoming edges, counted outgoing edges, excluded edge categories, source path policy, expected-isolated reason, and confidence.
- `no_incoming` and `detached_component` are prioritized as actionable candidates; `no_outgoing` is not presented as a bug by default.
- Expected-isolated nodes can be hidden or de-emphasized without deleting their semantic labels.
- A user can distinguish a real candidate from analyzer/resolution uncertainty.
- Benchmark ledger records measured counts for all statuses on at least `E:\AVmatrix-GO` and one large indexed repo when available.
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

Baseline counts are currently `pending measurement`.

## Phase 1 - Taxonomy, Policy, and Baseline

- [ ] [P1-A] Define counted edge types, excluded edge types, and compatibility/display edge handling for connectivity status.
- [ ] [P1-B] Define expected-isolated policy for test, fixture, generated, vendor, docs, migrations, scripts, exported API, route, tool, command, session, framework, and reflection/config surfaces.
- [ ] [P1-C] Define confidence levels: `candidate`, `expected`, `unknown`, and `confirmed`, with evidence requirements for each.
- [ ] [P1-D] Measure baseline graph connectivity for `E:\AVmatrix-GO` and record commands/results in benchmark and evidence ledgers.
- [ ] [P1-E] Measure baseline graph connectivity for one large indexed repo when available and record commands/results in benchmark and evidence ledgers.
- [ ] [P1-F] Record the selected taxonomy in this plan, generated contracts if needed, and user-facing Web wording.

## Phase 2 - Backend Graph-Health Derivation

- [ ] [P2-A] Identify the graph data boundary that should own derived connectivity status: graph package, analyzer output, HTTP graph payload, contract layer, or Web-only derived state.
- [ ] [P2-B] Implement deterministic connectivity summary generation using the Phase 1 edge policy.
- [ ] [P2-C] Add per-node derived metadata for connectivity status, reasons, counted incoming/outgoing counts, expected-isolated reason, and confidence.
- [ ] [P2-D] Add detached-component grouping and component-level explanations.
- [ ] [P2-E] Add unresolved-reference classification only where source/resolution evidence exists; otherwise classify as `unknown_connectivity`.
- [ ] [P2-F] Add unit tests for every taxonomy status and exclusion rule.

## Phase 3 - Contract, API, and Reporting Surface

- [ ] [P3-A] Decide whether graph-health metadata is emitted in Web graph payloads, generated Web contracts, MCP resources, or a dedicated endpoint.
- [ ] [P3-B] Add graph-health summary output with counts by status and expected-isolated reason.
- [ ] [P3-C] Add explain output for a single node or component.
- [ ] [P3-D] Add report/export path if needed for dead-code or unwired-candidate review.
- [ ] [P3-E] Add contract tests proving status fields are stable, explicit, and not confused with semantic labels.

## Phase 4 - Web UI Graph Health Filters

- [ ] [P4-A] Add a `Graph Health` filter group separate from `Node Types` and `Edge Types`.
- [ ] [P4-B] Add toggles for `true_isolated`, `no_incoming`, `no_outgoing`, `detached_component`, `unresolved_reference`, `expected_isolated`, and `unknown_connectivity`.
- [ ] [P4-C] Add summary counts and tooltips that explain status meaning without calling candidates bugs.
- [ ] [P4-D] Add node detail panel explanations: counted incoming/outgoing edges, expected-isolated reason, confidence, and next triage action.
- [ ] [P4-E] Add detached-component interaction that focuses a component and shows why it is detached.
- [ ] [P4-F] Ensure existing node-type, edge-type, legend, focus-depth, and graph canvas behavior still works with Graph Health filters.

## Phase 5 - Triage Workflow

- [ ] [P5-A] Define default triage order: `no_incoming` production symbols, `detached_component`, `unresolved_reference`, `true_isolated`, then optional `no_outgoing`.
- [ ] [P5-B] Add report wording for "candidate" versus "confirmed" findings.
- [ ] [P5-C] Add documentation or in-product text for why a status was assigned and what the next action should be.
- [ ] [P5-D] Add a way to hide or de-emphasize expected-isolated nodes without changing raw graph data.

## Phase 6 - Validation

- [ ] [P6-A] Run full Go build before tests.
- [ ] [P6-B] Run focused Go tests for graph-health derivation, taxonomy, expected-isolated policy, and contract/API behavior.
- [ ] [P6-C] Run full applicable Go test suite for `cmd` and `internal`.
- [ ] [P6-D] Run Web build before Web tests if Web UI changes.
- [ ] [P6-E] Run focused Web unit tests for Graph Health filters, node detail explanations, counts, legends, and filter interactions.
- [ ] [P6-F] Run full Web unit suite if Web UI changes.
- [ ] [P6-G] Run e2e covering Graph Health filter visibility, node explanation, detached-component focus, and interaction with existing node/edge filters if Web UI changes.
- [ ] [P6-H] Re-run baseline graph-health inventory after implementation and record before/after counts.

## Phase 7 - Closure

- [ ] [P7-A] Update this plan checklist after each completed slice.
- [ ] [P7-B] Update benchmark ledger with initial, intermediate, and final counts.
- [ ] [P7-C] Update evidence ledger with commands, files changed, tests, and conclusions.
- [ ] [P7-D] Commit each completed implementation slice.
- [ ] [P7-E] Final closure: confirm taxonomy, backend derivation, contract/API surface, Web UI filters, triage workflow, benchmark, evidence, full build, unit tests, and e2e tests are complete.

## Ledger

| ID | Area | Scope | Target | Benchmark | Evidence | Commit | Status |
|---|---|---|---|---|---|---|---|
| P1-A..P1-F | Policy | taxonomy and baseline | no ambiguous orphan claims | pending | pending | pending | open |
| P2-A..P2-F | Backend | derived graph-health metadata | deterministic status and reasons | pending | pending | pending | open |
| P3-A..P3-E | Contract/API | consumer surface | stable explicit status fields | pending | pending | pending | open |
| P4-A..P4-F | Web UI | graph-health filters | separate filters and explanations | pending | pending | pending | open |
| P5-A..P5-D | Workflow | triage/reporting | candidate-vs-confirmed workflow | pending | pending | pending | open |
| P6-A..P6-H | Validation | build/tests/e2e | full validation recorded | pending | pending | pending | open |
| P7-A..P7-E | Closure | ledgers and commits | complete closure package | pending | pending | pending | open |

## Definition Of Done

- The plan, evidence, and benchmark ledgers are all updated.
- Connectivity taxonomy has explicit statuses, edge policy, expected-isolated policy, and confidence rules.
- No UI or API labels "orphan" as a bug without evidence.
- Graph Health filters are separate from semantic node labels.
- Node explanations are auditable from recorded graph data.
- Benchmarks record measured counts, not assumptions.
- Full build and applicable Go/Web/e2e validation pass.
- AVmatrix status is up to date after final implementation commit.
