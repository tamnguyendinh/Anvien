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

Status: pending measurement

Required repos:

| Repo | Required | Status |
|---|---|---|
| `E:\AVmatrix-GO` | yes | pending |
| one large indexed repo when available | yes | pending |

Required baseline inventory:

| Metric | `E:\AVmatrix-GO` | Large repo | Notes |
|---|---:|---:|---|
| Raw node count | pending | pending | measured from graph payload |
| Raw relationship count | pending | pending | measured from graph payload |
| Semantic node labels | pending | pending | count by label |
| Relationship types | pending | pending | count by type |
| Counted edge policy version | pending | pending | must be defined in plan/evidence |
| Nodes with zero counted incoming edges | pending | pending | not equivalent to bug count |
| Nodes with zero counted outgoing edges | pending | pending | not equivalent to bug count |
| Nodes with zero counted incoming and outgoing edges | pending | pending | candidate `true_isolated` before exclusions |
| Expected-isolated nodes | pending | pending | count by reason |
| `true_isolated` after exclusions | pending | pending | measured after expected-isolated policy |
| `no_incoming` after exclusions | pending | pending | actionable candidate pool |
| `no_outgoing` after exclusions | pending | pending | inspect-only by default |
| Detached components | pending | pending | count and top component sizes |
| Unresolved references | pending | pending | only if source/resolution evidence exists |
| Unknown connectivity | pending | pending | insufficient evidence |

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

- final graph-health counts for required repos;
- final expected-isolated exclusion counts by reason;
- final actionable candidate counts;
- final validation inventory;
- final interpretation explaining what remains candidate/unknown versus confirmed.
