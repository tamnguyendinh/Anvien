# Anvien Orphan Node Connectivity Lens Benchmark Ledger

Date: 2026-05-20

Status: complete

Companion files:

- Plan: [2026-05-20-anvien-orphan-node-connectivity-lens-plan.md](2026-05-20-anvien-orphan-node-connectivity-lens-plan.md)
- Evidence ledger: [2026-05-20-anvien-orphan-node-connectivity-lens-evidence.md](2026-05-20-anvien-orphan-node-connectivity-lens-evidence.md)

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

Status: Phase 1 accepted policy + fresh `E:\Anvien` + representative `Restaurant_manager` baselines recorded; cross-repo re-measure on future index updates noted as follow-up

### Accepted Counted Edge Policy (Phase 1 final for baseline)

**Counted relationships** (contribute to per-node incoming/outgoing for topologyStatus; these represent semantic wiring, usage, call-flow, heritage, registration, and process participation):
- CALLS, ACCESSES, INHERITS, IMPLEMENTS, EXTENDS, METHOD_OVERRIDES, METHOD_IMPLEMENTS, IMPORTS, USES, DECORATES, WRAPS, QUERIES, FETCHES, STEP_IN_PROCESS, HANDLES_ROUTE, HANDLES_TOOL, ENTRY_POINT_OF

**Non-counted relationships** (structural/ownership/containment; excluded because they always attach symbols to files/containers and would mask isolation candidates):
- CONTAINS, DEFINES, HAS_METHOD, HAS_PROPERTY, MEMBER_OF

Rationale: Empirical (raw-all yields 0 code zero-incoming on both repos because every symbol has a DEFINES); layout hierarchy in graph-adapter.ts already separates these; processes.go and ignore policy treat structural as given. Matches design decision that "no_incoming" must surface real unwired candidates after ownership is stripped.

### Representative Repo Selection Criteria (P1-E1)
- Large indexed repos (>5k files or >50k nodes) outside the Anvien codebase itself to test generalizability.
- Diverse label mix (high code + Process/Route/Tool for entry surfaces, high Section for doc coverage).
- Presence of framework entry patterns (routes, tools, processes) so detached_component rules have positive examples.
- Recent or usable .anvien/graph.json snapshot available.
- Restaurant_manager selected: 6198 files, 78k nodes, 505 processes, heavy on Sections + Functions/Methods, representative of fullstack app with UI+backend.

### Measured Baselines (using accepted non-counted structural policy + fresh Anvien index 2026-05-20 + Restaurant snapshot)

**E:\Anvien** (graph: nodes=21091, rels=52445, timestamp from analyze --force 2026-05-20, code_nodes=4934):
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

All counts produced by reproducible python loader over graph.json (no Anvien MCP used for final numbers; commands logged in evidence). Raw counts prove why structural exclusion is mandatory.

Required repos:

| Repo | Required | Status |
|---|---|---|
| `E:\Anvien` | yes | initial measured |
| Representative indexed repo set | yes, when criteria are recorded and repos are available | criteria recorded; `Restaurant_manager` selected and measured |

Required baseline inventory:

| Metric | `E:\Anvien` | Notes |
|---|---:|---|
| Raw node count | `20,967` | measured from graph payload |
| Raw relationship count | `52,302` | measured from graph payload |
| Semantic node labels | recorded below | count by label |
| Relationship types | recorded below | count by type |
| Counted edge policy version | raw-all, provisional non-structural, and accepted non-structural policy | accepted policy recorded above |
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
| Detached components | measured in B1 | requires accepted root/entry policy |
| Unresolved references | measured in B1/P2-E | only source-backed diagnostics become node diagnostics |
| Unknown connectivity | measured in B1/P2-E | source-backed unresolved diagnostics move affected nodes to unknown |

Representative indexed repo baseline:

| Requirement | Status |
|---|---|
| Selection criteria | recorded in Phase 1 |
| Selected repo list | `Restaurant_manager` selected |
| Raw/provisional counts | recorded above |
| Accepted-policy counts | recorded above |

Top node labels:

| Repo | Top labels |
|---|---|
| `E:\Anvien` | `Variable:8627`, `Function:3428`, `Property:3235`, `Section:1071`, `Community:923`, `Method:809`, `File:694`, `Process:640`, `Struct:508`, `Package:413`, `Const:323`, `Folder:112` |

Relationship type inventory:

| Repo | Relationship types |
|---|---|
| `E:\Anvien` | `ACCESSES:5179`, `CALLS:8580`, `CONTAINS:1852`, `DEFINES:17527`, `ENTRY_POINT_OF:640`, `HAS_METHOD:339`, `HAS_PROPERTY:2891`, `IMPORTS:3755`, `MEMBER_OF:3909`, `STEP_IN_PROCESS:2362`, `USES:5268` |

Initial benchmark interpretation:

- Raw all-relationship connectivity is not sufficient for orphan/dead-code triage because structural edges such as `DEFINES`, `CONTAINS`, `HAS_METHOD`, `HAS_PROPERTY`, and `MEMBER_OF` dominate incoming connectivity.
- The large difference between raw-all and provisional non-structural counts is the concrete reason Phase 1 cannot be skipped.
- The provisional policy numbers are evidence for planning only; they are not accepted product behavior.
- Representative `Restaurant_manager` measurements are recorded; re-run on future index updates is recommended.

## B1 - Implementation Benchmark

Status: Phase 2/3 backend/API/contract slices recorded through P2-E on 2026-05-20

Source:

- Repo: `E:\Anvien`
- Graph: `.anvien/graph.json`
- Policy version: `graph-health-non-structural-v1`
- P2-D refresh command: `go run .\cmd\anvien analyze --force [redacted removed argument] --no-stats`
- P2-D refresh result: `nodes=21323 relationships=52940`
- P2-E refresh command: `go run .\cmd\anvien analyze --force [redacted removed argument] --no-stats --benchmark-json .tmp\p2e-unresolved-diagnostics-benchmark.json --benchmark-label p2e-unresolved-diagnostics`
- P2-E refresh result: `files: scanned=707 parsed=534 unsupported=173 failed=0`; `graph: nodes=21388 relationships=53228`
- Measurement code: temporary Go runner inside repo root importing `internal/graph` and `internal/graphhealth`, loading `.anvien/graph.json`, calling `graphhealth.ComputeSummary(&g)`, measuring content-stripped graph payload size before/after public Graph Health metadata, then removing the temporary runner.

### P2-D Detached-Component Output

This historical slice output is kept to show the before-diagnostics component traversal measurement.

```text
runtime_ms=181.027
base_payload_bytes=28415767
annotated_payload_bytes=39167950
delta_bytes=10752183
delta_percent=37.84

topologyStatusCounts:
  connected=2581
  true_isolated=13956
  no_incoming=1671
  no_outgoing=2958
  detached_component=157
  unknown_connectivity=0

expectedIsolationReasonCounts:
  documentation=1258
  framework_entry=798
  test=5780

confidenceCounts:
  candidate=13487
  expected=7836
  unknown=0
  confirmed=0

countedRelationshipCount=26033
componentCount=14006
detachedComponentCount=48
rootNodeCount=798
excludedEdgeCounts.structural=26909
diagnosticCounts={}
```

### P2-E Source-Backed Unresolved Diagnostics Output

```text
runtime_ms=182.586
graph_snapshot_bytes=41745701
base_payload_bytes=37175228
annotated_payload_bytes=46975939
delta_bytes=9800711
delta_percent=26.36

nodeCount=21388
relationshipCount=53228
countedRelationshipCount=26201
componentCount=14031
detachedComponentCount=48
rootNodeCount=797

unresolvedReferenceCount=49576
sourceBackedUnresolvedReferenceCount=49576
unattributedUnresolvedReferenceCount=0

topologyStatusCounts:
  connected=602
  true_isolated=13812
  no_incoming=163
  no_outgoing=2451
  detached_component=59
  unknown_connectivity=4301

expectedIsolationReasonCounts:
  documentation=1258
  framework_entry=797
  test=5793

confidenceCounts:
  candidate=10997
  expected=6090
  unknown=4301
  confirmed=0

diagnosticCounts:
  unresolved_reference=49576

excludedEdgeCounts:
  structural=27027

diagnosticRecords=8756
```

Notes:

- `detached_component` represents nodes inside weak counted-edge components that have internal counted edges but no accepted root reachability.
- P2-E changes many candidate statuses to `unknown_connectivity` only when source-backed unresolved diagnostics exist on the node.
- `diagnosticCounts.unresolved_reference=49576` counts diagnostic occurrences after aggregation, not the number of serialized diagnostic records.
- Public diagnostics are aggregated into 8,756 records; the first non-aggregated attempt produced 49,561 records and an excessive payload, so aggregation is the accepted implementation.
- Payload size compares a content-stripped API-like graph JSON payload before and after adding public per-node Graph Health metadata plus top-level summary. The internal raw `graphHealthDiagnostics` node property is stripped from HTTP JSON/NDJSON responses.
- Component root IDs are not repeated per node; they are reserved for component summaries to avoid excessive payload growth.

| Metric | Baseline | After P2-E | Delta | Interpretation |
|---|---:|---:|---:|---|
| Graph-health summary runtime | n/a | `182.586ms` | n/a | measured server-side derivation on current `E:\Anvien` graph with component traversal and diagnostic aggregation |
| Graph payload size | `37,175,228 bytes` | `46,975,939 bytes` | `+9,800,711 bytes` (`+26.36%`) | content-stripped JSON payload with public metadata embedded and internal raw diagnostics stripped |
| Web graph filter render latency | pending | pending | pending | only if Web filter implementation changes rendering |
| `true_isolated` count | pending | `13,812` | pending | measured by `ComputeSummary`; includes expected/candidate overlay separation |
| `no_incoming` count | pending | `163` | pending | measured by `ComputeSummary` after source-backed unresolved nodes move to unknown |
| `no_outgoing` count | pending | `2,451` | pending | measured by `ComputeSummary`; low-priority triage by policy |
| `connected` count | pending | `602` | pending | measured by `ComputeSummary` |
| `detached_component` node count | pending | `59` | pending | measured by component traversal after source-backed unresolved nodes move to unknown |
| Detached component count | pending | `48` | pending | measured by component traversal |
| Component count | pending | `14,031` | pending | measured by component traversal |
| Root node count | pending | `797` | pending | accepted roots used for directed reachability |
| `unknown_connectivity` count | pending | `4,301` | pending | nodes with source-backed unresolved diagnostics |
| Expected-isolated reason count: `test` | pending | `5,793` | pending | measured by `ComputeSummary` |
| Expected-isolated reason count: `documentation` | pending | `1,258` | pending | measured by `ComputeSummary` |
| Expected-isolated reason count: `framework_entry` | pending | `797` | pending | measured by accepted root policy |
| Confidence count: `candidate` | pending | `10,997` | pending | measured by `ComputeSummary`; not a bug/deletion verdict |
| Confidence count: `expected` | pending | `6,090` | pending | measured by `ComputeSummary` |
| Confidence count: `unknown` | pending | `4,301` | pending | source-backed unresolved diagnostics |
| Diagnostic occurrence count: `unresolved_reference` | pending | `49,576` | pending | source-backed unresolved diagnostics emitted by resolution and counted via aggregated `count` |
| Diagnostic record count | pending | `8,756` | pending | serialized aggregated records, not occurrences |
| Source-backed unresolved references | pending | `49,576` | pending | all unresolved references in this snapshot had source-node or file-node evidence |
| Unattributed unresolved references | pending | `0` | pending | no unresolved references lacked attachable source evidence in this snapshot |

### P2-F Coverage Slice

P2-F is a test-only coverage slice. It changes no production graph-health derivation logic, graph inventory counts, payload shape, runtime path, or Web rendering behavior, so it has no new benchmarkable product/runtime metric. Validation evidence is recorded in E9.

### P3-C Explain Endpoint Slice

P3-C adds a bounded explain API surface and generated response contracts. It does not change existing `/api/graph` inventory counts, graph summary counts, or Web render behavior. Component explain responses are bounded by `sampleLimit=20`; endpoint validation evidence is recorded in E10. Dedicated explain-endpoint runtime benchmarking remains a follow-up if this endpoint becomes part of a high-frequency Web interaction.

### P3-D Report Export Slice

P3-D adds JSON report/export contracts and `GET /api/graph/report`. It does not change existing graph inventory counts, `/api/graph` payload size, or Web render behavior. Report responses default to `limit=100`, cap at `1000`, and preserve `candidate_not_confirmed` wording. Dedicated report runtime benchmarking remains a follow-up if the endpoint is used for large automated exports.

### P4-A/B/C/G/H/I Web Graph Health Filter Composition Slice

P4 filter-composition changes only Web state, dashboard controls, graph conversion attributes, and client-side filtering. It does not change backend graph inventory counts, graph-health derivation counts, `/api/graph` payload size, explain/report output, or graph snapshot shape.

Benchmarkable package-size observation from the accepted Web build:

```text
Command: npm --prefix anvien-web run build
CSS: assets/index-BMziPvPs.css 55.17 kB, gzip 10.55 kB
Main JS: assets/index-BmCY7_Nr.js 2,018.26 kB, gzip 601.51 kB
Build warnings: existing ProcessFlowModal dynamic/static import warning and >500 kB chunk-size warning
```

Web render/filter latency was not measured in this slice. Functional coverage is recorded in E12 through focused unit tests, full Web unit suite, and targeted Playwright dashboard e2e. Dedicated latency measurement remains a follow-up if Graph Health filtering becomes a performance-sensitive interaction on large graphs.

### P4-D/E/F Web Graph Health Detail And Detached Focus Slice

P4 detail/focus changes Web presentation, dashboard tooltips/counts, selected-node detail rendering, and user-driven component highlighting. It does not change backend graph inventory counts, graph-health derivation counts, `/api/graph` payload size, explain/report output, or graph snapshot shape.

Benchmarkable package-size observation from the accepted Web build:

```text
Command: npm --prefix anvien-web run build
CSS: assets/index-CernHlw7.css 55.27 kB, gzip 10.57 kB
Main JS: assets/index-DQRhUrAU.js 2,026.11 kB, gzip 604.14 kB
Build warnings: existing ProcessFlowModal dynamic/static import warning and >500 kB chunk-size warning
```

Dedicated Web filter/detail latency measurement remains pending. Functional coverage is recorded in E13 through focused unit tests, full Web unit suite, targeted large-graph Playwright dashboard e2e, and deterministic mocked Graph Health e2e.

### P5 Triage Workflow Reconciliation

P5 is a doc-only reconciliation of behavior already implemented in P3-D and P4. It changes no runtime code, graph inventory, graph-health derivation counts, payload size, package size, or Web render behavior. No new benchmarkable metric is recorded for this slice; evidence is recorded in E14.

## B2 - Final Benchmark

Status: recorded 2026-05-20

Source:

- Repo: `E:\Anvien`
- HTTP graph summary command: `$payload = Invoke-RestMethod -Uri 'http://127.0.0.1:4848/api/graph?repo=Anvien&stream=false' -TimeoutSec 120; $payload.graphHealth | ConvertTo-Json -Depth 8`
- Repo registry observation: `Anvien` indexed at `2026-05-20T08:23:34Z`, graph `21,658` nodes / `53,962` relationships.
- Representative Web latency repo: `Restaurant_manager`, graph `78,358` nodes / `130,588` relationships.

Final `Anvien` graph-health counts:

| Metric | Final count |
|---|---:|
| Node count | `21,658` |
| Relationship count | `53,962` |
| Counted relationship count | `26,578` |
| Component count | `14,186` |
| Detached component count | `48` |
| Root node count | `805` |
| Unresolved references | `50,543` |
| Source-backed unresolved references | `50,543` |
| Unattributed unresolved references | `0` |
| Excluded structural edges | `27,384` |

Final topology status counts:

| Topology status | Count |
|---|---:|
| `connected` | `605` |
| `true_isolated` | `13,964` |
| `no_incoming` | `166` |
| `no_outgoing` | `2,528` |
| `detached_component` | `59` |
| `unknown_connectivity` | `4,336` |

Final expected-isolated reason counts:

| Reason | Count |
|---|---:|
| `test` | `5,857` |
| `documentation` | `1,271` |
| `framework_entry` | `805` |

Final confidence counts:

| Confidence | Count |
|---|---:|
| `candidate` | `11,164` |
| `expected` | `6,158` |
| `unknown` | `4,336` |
| `confirmed` | `0` |

Final diagnostic counts:

| Diagnostic | Count |
|---|---:|
| `unresolved_reference` | `50,543` |

Web Graph Health filter/detail latency on `Restaurant_manager`:

| Interaction | Measured latency |
|---|---:|
| `No incoming` off | `4,236.831ms` |
| `No incoming` on | `3,541.886ms` |
| `Test` off | `4,198.017ms` |
| `Test` on | `4,256.324ms` |
| `Unresolved reference` off | `3,856.182ms` |
| `Unresolved reference` on | `6,402.975ms` |
| Average | `4,415.369ms` |
| p95 | `6,402.975ms` |

Web runtime diagnostics captured during the latency benchmark:

| Metric | Value |
|---|---:|
| Graph conversion time | `4,980.800ms` |
| Visual graph node count | `78,358` |
| Max node size | `3` |
| Max rendered node size cap | `3` |
| Structural-to-leaf ratio | `3` |

Interpretation:

- `candidate` remains a triage signal, not a bug or deletion verdict.
- `unknown_connectivity` is driven by source-backed unresolved diagnostics and should be inspected before topology is treated as meaningful.
- `confirmed` remains `0` because Graph Health derivation does not auto-confirm dead code or defects.
- The measured large-graph Web toggle latency is noticeable on `Restaurant_manager`; future optimization can target incremental Sigma visibility updates if this workflow becomes high-frequency.
