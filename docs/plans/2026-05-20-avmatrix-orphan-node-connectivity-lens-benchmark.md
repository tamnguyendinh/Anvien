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

Status: initial `E:\AVmatrix-GO` measurements recorded; accepted-policy and representative baselines pending

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
