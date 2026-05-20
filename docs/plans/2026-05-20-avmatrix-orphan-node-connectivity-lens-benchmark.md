# AVmatrix Orphan Node Connectivity Lens Benchmark Ledger

Date: 2026-05-20

Status: active

Companion files:

- Plan: [2026-05-20-avmatrix-orphan-node-connectivity-lens-plan.md](2026-05-20-avmatrix-orphan-node-connectivity-lens-plan.md)
- Evidence ledger: [2026-05-20-avmatrix-orphan-node-connectivity-lens-evidence.md](2026-05-20-avmatrix-orphan-node-connectivity-lens-evidence.md)

## Benchmark Rules

Record benchmarkable results when they measure product/runtime performance, capacity, package/startup size, graph/DB throughput, or graph inventory counts. Build/test/e2e timings are validation evidence unless the slice changes those systems.

For this plan, benchmarkable measurements include:

- raw graph node count by semantic label;
- raw graph relationship count by type;
- counted incoming edge distribution under the accepted edge policy;
- counted outgoing edge distribution under the accepted edge policy;
- `true_isolated` node count;
- `no_incoming` node count;
- `no_outgoing` node count;
- `detached_component` count and largest component sizes;
- `unresolved_reference` count by fact family if source/resolution evidence exists;
- `expected_isolated` count by reason;
- `unknown_connectivity` count;
- before/after graph-health filter counts in Web UI;
- graph-health summary generation runtime if implemented server-side;
- graph payload size delta if connectivity metadata is added to payloads;
- Web render/filter latency if Graph Health filters change graph interaction performance.

Do not record inferred or estimated counts. Every benchmark row must name the command, repo path, graph source, and commit or graph timestamp when available.

## B0 - Required Initial Connectivity Baseline

Date: 2026-05-20

Status: Phase 1 accepted policy + fresh `E:\AVmatrix-GO` + representative `Restaurant_manager` baselines recorded; cross-repo re-measure on future index updates noted as follow-up

### Accepted Counted Edge Policy (Phase 1 final for baseline)

**Counted relationships** (contribute to per-node incoming/outgoing for topologyStatus; these represent semantic wiring, usage, call-flow, heritage, registration, and process participation):
- CALLS, ACCESSES, INHERITS, IMPLEMENTS, EXTENDS, METHOD_OVERRIDES, METHOD_IMPLEMENTS, IMPORTS, USES, DECORATES, WRAPS, QUERIES, FETCHES, STEP_IN_PROCESS, HANDLES_ROUTE, HANDLES_TOOL, ENTRY_POINT_OF

**Non-counted relationships** (structural/ownership/containment; excluded because they always attach symbols to files/containers and would mask isolation candidates):
- CONTAINS, DEFINES, HAS_METHOD, HAS_PROPERTY, MEMBER_OF

Rationale: Empirical (raw-all yields 0 code zero-incoming on both repos because every symbol has a DEFINES); layout hierarchy in graph-adapter.ts already separates these; processes.go and ignore policy treat structural as given. Matches design decision that "no_incoming" must surface real unwired candidates after ownership is stripped.

### Representative Repo Selection Criteria (P1-E1)
- Large indexed repos (>5k files or >50k nodes) outside the AVmatrix codebase itself to test generalizability.
- Diverse label mix (high code + Process/Route/Tool for entry surfaces, high Section for doc coverage).
- Presence of framework entry patterns (routes, tools, processes) so detached_component rules have positive examples.
- Recent or usable .avmatrix/graph.json snapshot available.
- Restaurant_manager selected: 6198 files, 78k nodes, 505 processes, heavy on Sections + Functions/Methods, representative of fullstack app with UI+backend.

### Measured Baselines (using accepted non-counted structural policy + fresh AVmatrix index 2026-05-20 + Restaurant snapshot)

**E:\AVmatrix-GO** (graph: nodes=21091, rels=52445, timestamp from analyze --force 2026-05-20, code_nodes=4934):
- raw-all: zero_in=20, zero_out=15314, zero_both=8; code zero_in=0, zero_out=200, zero_both=0
- accepted-policy non-structural: code zero_in=1620, zero_out=1098, zero_both=133
- path-expected-candidate rough filter (test/fixture/generated/vendor/doc): ~6998 nodes
- top rels: DEFINES:17545, CALLS:8588, USES:5270, ACCESSES:5180, MEMBER_OF:3914, IMPORTS:3761, HAS_PROPERTY:2891, STEP_IN_PROCESS:2361, CONTAINS:1956, ENTRY_POINT_OF:640, HAS_METHOD:339
- top labels: Variable:8640, Function:3433, Property:3235, Section:1165, ...

**Restaurant_manager** (graph: nodes=78358, rels=130588, snapshot 2026-05-19, code_nodes=10258):
- raw-all: code zero_in=0, zero_out=1097, zero_both=0
- accepted-policy non-structural: code zero_in=4191, zero_out=3488, zero_both=909
- Note: higher structural ratio (CONTAINS 42k, DEFINES 34k) and many Sections (35k); provisional only — re-run on fresh index recommended after policy-locked implementation.
- Demonstrates policy scales: ~41% of code nodes appear no_incoming under accepted policy, as expected for a large app (many UI components, helpers, exported surfaces).

All counts produced by reproducible python loader over graph.json (no AVmatrix MCP used for final numbers; commands logged in evidence). Raw counts prove why structural exclusion is mandatory.

Required repos:

| Repo | Required | Status |
|---|---|---|
| `E:\AVmatrix-GO` | yes | initial measured |
| Representative indexed repo set | yes, when criteria are recorded and repos are available | pending |

Required baseline inventory:

| Metric | `E:\AVmatrix-GO` | Notes |
|---|---:|---|
| Raw node count | `20,967` | measured from graph payload |
| Raw relationship count | `52,302` | measured from graph payload |
| Semantic node labels | recorded below | count by label |
| Relationship types | recorded below | count by type |
| Counted edge policy version | raw-all and provisional non-structural | final accepted policy remains Phase 1 |
| Nodes with zero raw incoming edges | `20` | all relationship types counted |
| Nodes with zero raw outgoing edges | `15,216` | all relationship types counted |
| Nodes with zero raw incoming and outgoing edges | `8` | all relationship types counted |
| Code nodes with zero raw incoming edges | `0 / 4,929` | proves raw-all is not a useful dead-code denominator |
| Code nodes with zero raw outgoing edges | `200 / 4,929` | all relationship types counted |
| Path expected-isolated candidate nodes | `5,743` | existing path/test/generated/vendor/fixture-like policy candidate, not final |
| Path expected-isolated candidate zero-both nodes | `0` | raw-all relationship policy |
| Provisional non-structural code zero incoming | `1,616 / 4,929` | excludes structural/ownership edge types; not final product policy |
| Provisional non-structural code zero outgoing | `1,100 / 4,929` | excludes structural/ownership edge types; not final product policy |
| Provisional non-structural code zero-both | `133 / 4,929` | excludes structural/ownership edge types; not final product policy |
| Callable-flow zero incoming | `1,587 / 4,242` | counted types: `CALLS`, `HANDLES_ROUTE`, `HANDLES_TOOL`, `ENTRY_POINT_OF`, `STEP_IN_PROCESS` |
| Callable-flow zero outgoing | `1,323 / 4,242` | counted types: `CALLS`, `HANDLES_ROUTE`, `HANDLES_TOOL`, `ENTRY_POINT_OF`, `STEP_IN_PROCESS` |
| Callable-flow zero-both | `264 / 4,242` | counted types: `CALLS`, `HANDLES_ROUTE`, `HANDLES_TOOL`, `ENTRY_POINT_OF`, `STEP_IN_PROCESS` |
| Detached components | pending | requires accepted root/entry policy |
| Unresolved references | pending | only if source/resolution evidence exists |
| Unknown connectivity | pending | insufficient evidence |

Representative indexed repo baseline:

| Requirement | Status |
|---|---|
| Selection criteria | pending in Phase 1 |
| Selected repo list | pending |
| Raw/provisional counts | pending |
| Accepted-policy counts | pending |

Top node labels:

| Repo | Top labels |
|---|---|
| `E:\AVmatrix-GO` | `Variable:8627`, `Function:3428`, `Property:3235`, `Section:1071`, `Community:923`, `Method:809`, `File:694`, `Process:640`, `Struct:508`, `Package:413`, `Const:323`, `Folder:112` |

Relationship type inventory:

| Repo | Relationship types |
|---|---|
| `E:\AVmatrix-GO` | `ACCESSES:5179`, `CALLS:8580`, `CONTAINS:1852`, `DEFINES:17527`, `ENTRY_POINT_OF:640`, `HAS_METHOD:339`, `HAS_PROPERTY:2891`, `IMPORTS:3755`, `MEMBER_OF:3909`, `STEP_IN_PROCESS:2362`, `USES:5268` |

Initial benchmark interpretation:

- Raw all-relationship connectivity is not sufficient for orphan/dead-code triage because structural edges such as `DEFINES`, `CONTAINS`, `HAS_METHOD`, `HAS_PROPERTY`, and `MEMBER_OF` dominate incoming connectivity.
- The large difference between raw-all and provisional non-structural counts is the concrete reason Phase 1 cannot be skipped.
- The provisional policy numbers are evidence for planning only; they are not accepted product behavior.
- Cross-repo measurements remain pending until Phase 1 defines representative selection criteria.

## B1 - Implementation Benchmark

Status: pending

Record after backend/contract implementation:

| Metric | Baseline | After implementation | Delta | Interpretation |
|---|---:|---:|---:|---|
| Graph-health summary runtime | pending | pending | pending | only if server-side generation exists |
| Graph payload size | pending | pending | pending | only if metadata is embedded in payload |
| Web graph filter render latency | pending | pending | pending | only if Web filter implementation changes rendering |
| `true_isolated` count | pending | pending | pending | measured, not inferred |
| `no_incoming` count | pending | pending | pending | measured, not inferred |
| `detached_component` count | pending | pending | pending | measured, not inferred |

## B2 - Final Benchmark

Status: pending

Final closure must record:

- final graph-health counts for `E:\AVmatrix-GO` and any representative indexed repos selected during Phase 1;
- final expected-isolated exclusion counts by reason;
- final actionable candidate counts;
- final validation inventory;
- final interpretation explaining what remains candidate/unknown versus confirmed.
