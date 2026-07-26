# P2-C Projection, Command, and Derived-Behavior First Divergence

Date: 2026-07-26 16:12 +07
Target: `E:\cheapapp.org`
Target HEAD: `a869876ab6262dacde6cd5d432d099a91852a646`
Target graph: `E:\cheapapp.org\.anvien\graph.json`
Target graph SHA-256: `DD3A3A987E95C91787E7C70EF68BA5D2F953D24A32D35DEFE38B8E8E606FF090`
Target graph inventory: `84,807` nodes / `114,125` relationships
Classification: bounded command fidelity plus bounded representation non-parity; no global process/semantic claim
Revision: corrected after the 16:32 Supervisor rejection; process-control values and interpretation now match the current probe and five reruns.

## Question

For the fixed P1-D/P1-E cases, does a downstream command or derived layer independently lose or invent facts, or does it expose the upstream graph as supplied? Where the raw graph and Cypher/file-detail shapes differ, what is the first Anvien source boundary that changes the representation?

## Boundary and freshness

- The target graph was read in place. It was not copied into `E:\Anvien`, and no target source, graph, index, generated guidance file, test, or configuration was written.
- No target `anvien analyze` was rerun. The target graph retained the hash, counts, and last-write time `2026-07-26T14:49:43.8894011+07:00` recorded above.
- The Anvien self-index at `E:\Anvien\.anvien` was refreshed before owner `file-detail`/`impact` inspection. That refresh did not touch the target graph.
- P1-D and P1-E command captures were treated as inputs and rechecked against the raw target graph and current Anvien source. This report does not accept a global process, community, semantic, or command-completeness claim.
- All new probes and captures are under `E:\Anvien\.tmp\cheapapp-graph-root-cause-restart\p2c`.

## Result summary

| Surface | Bounded result | First relevant boundary |
|---|---|---|
| `context` incoming/outgoing edges | faithful to relationships present in the raw graph; it does not synthesize the two missing barrel-mediated calls | `internal/mcp/context_index.go:12-77` |
| `impact` traversal | faithful to the configured relationship set; it reaches consumer files through existing `IMPORTS`/`USES` but cannot recover absent `CALLS` | `internal/mcp/impact.go:554-631` |
| `context`/`impact` process arrays | empty for the selected symbols because the raw graph has zero selected `STEP_IN_PROCESS` edges; no command-layer process edge is dropped | `internal/mcp/context_index.go:37-52`; `internal/mcp/impact.go:943-999` |
| `file-detail` unresolved rows | one-to-one record/cardinality parity for the seven canonical facts; shape is deliberately reduced to ten fields | `internal/filecontext/context.go:829-890`; `internal/filecontext/compact.go:169-190,489-507` |
| Cypher `step` on non-process edges | confirmed representation non-parity: absent raw `Relationship.Step` is serialized into database value `0` | `internal/lbugload/csv.go:372-387` |
| Cypher scalar JSON types | confirmed native-read normalization: numeric/boolean values are returned as strings | `internal/lbugnative/native_ladybugdb.go:132-165` |

The fixed missing `readAdminCommercialConfig` consumer calls and false ambient gaps therefore remain upstream P2-B failures. P2-C found no additional node/edge/cardinality loss in `context`, `impact`, or `file-detail`. It did find two bounded database/Cypher representation changes—`nil -> 0` for `step`, and scalar values -> strings—that do not exist in the raw JSON graph.

## 1. `context` reads exact graph relationships

`contextToolInternal` selects the symbol and calls `contextNeighborhood` directly (`internal/mcp/context.go:139-152`). `contextNeighborhood` scans `g.Relationships` and emits a row only when the selected symbol is the exact source or target of a configured relationship (`internal/mcp/context_index.go:22-36`). It has no re-export chase, call inference, or fallback synthesis.

For the fixed wrapper:

- raw graph has barrel File -> wrapper Function `USES`;
- raw graph has consumer File -> barrel File `IMPORTS`;
- raw graph has no consumer Function -> wrapper Function `CALLS`;
- P1-E `context` therefore returns the barrel as the sole incoming symbol-level row, while exposing the wrapper's real outgoing `CALLS`/`USES` rows.

For `boundedReadLimit`, raw graph has one `readEmailOperationsReport -> boundedReadLimit` `CALLS` edge, and P1-E `context` returns exactly that incoming call. Its `Math.max`/`Math.min` gaps are added separately by `contextSourceResolutionGaps`, which scans exact `HAS_RESOLUTION_GAP` edges (`internal/mcp/context.go:564-587`).

Existing control `TestContextToolReturnsSemanticFieldsAndSourceResolutionGaps` adds ordinary and process edges to an in-memory graph and verifies they are returned. The targeted MCP test passed in this slice.

Classification: **bounded valid downstream command behavior**. The missing wrapper consumers first diverge in P2-B, not in `context`.

## 2. `impact` traverses only supplied edges

`runImpactBFSProfiled` creates a frontier at the selected graph node, scans `g.Relationships`, filters to configured relationship types/confidence, and traverses the exact source/target direction (`internal/mcp/impact.go:554-619`). It neither resolves imports nor manufactures missing calls.

This explains the P1-E wrapper result:

- depth 1 reaches the barrel File through existing `USES`;
- depth 2 reaches consumer Files through existing file-level `IMPORTS`;
- no depth contains the missing consumer Function -> wrapper Function `CALLS`, because those relationships are absent upstream.

The `LOW`, count `7`, direct `1` result is therefore a risk computation over the incomplete input graph. It is not an independent resolver result and cannot be interpreted as the source truth blast radius.

Classification: **bounded valid traversal of an upstream-incomplete graph**.

## 3. Empty process arrays do not prove no product process

The target graph contains `662` Process nodes and `2,761` `STEP_IN_PROCESS` relationships overall. The raw probe found zero `STEP_IN_PROCESS` relationships for each selected file and each selected symbol:

- `readAdminCommercialConfig`;
- `boundedReadLimit`;
- `readEmailOperationsReport`;
- `validateCommercialFreshness` (the enclosing source function for the line-142 call);
- `readAdminCommercialConfigRouteView`.

`contextNeighborhood` adds a process row only for a direct `symbol -> Process` `STEP_IN_PROCESS` relationship (`internal/mcp/context_index.go:37-52`). `impactAffectedProcesses` likewise aggregates only `STEP_IN_PROCESS` relationships whose source is already in the impacted-node set (`internal/mcp/impact.go:943-999`). Thus P1-E's empty arrays match the raw graph; the command layer does not drop an existing selected process membership.

The derived process producer is intentionally heuristic. `processes.Apply`:

- builds only from `CALLS` with confidence at least `0.5` (`internal/processes/processes.go:242-255`);
- scores and caps entry candidates at `200` (`internal/processes/processes.go:271-310`);
- requires at least three steps by default (`internal/processes/processes.go:151-166,313-363`);
- deduplicates/subsets and caps emitted traces (`internal/processes/processes.go:69-89`);
- emits `STEP_IN_PROCESS` only for traces that survive those gates (`internal/processes/processes.go:95-133`).

An in-memory differential control removed existing derived process nodes, reran the producer, and then added only the two missing consumer -> wrapper `CALLS` edges:

| Run | `CALLS` considered | Process nodes | Step edges | selected consumer process memberships | wrapper memberships |
|---|---:|---:|---:|---:|---:|
| actual target graph | 3,771 | 662 | 2,761 | 0 | 0 |
| + two synthetic missing calls | 3,773 | 662 | 2,761 | 0 | 0 |

The exact two-edge control changes only the calls-considered count (`3,771` to `3,773`); process/step counts and selected memberships remain identical. This single control establishes only that no output change was observed under this exact configuration. It does not support any broader process property, and the wrapper remaining absent prevents a stronger bounded flow claim.

Classification: **command projection valid for present edges; derived-process completeness unresolved**. No global process/semantic conclusion is authorized.

## 4. `file-detail` preserves facts but reduces their shape

P1-D established one raw gap node, one `HAS_RESOLUTION_GAP` edge, one Cypher gap row, and one file-detail sample for every selected canonical source-site ID. The current source trace explains why file-detail omits raw fields without losing those seven rows:

1. `Builder.buildUnresolved` indexes all gap nodes for the requested file and constructs `UnresolvedSample` with exactly ten fields: line, column, target text, source symbol, gap kind, classification, actionability, proof kind, source-site ID, and source-site status (`internal/filecontext/context.go:829-890`).
2. The compact schema declares exactly those ten columns (`internal/filecontext/compact.go:169-190`).
3. `unresolvedRows` interns only those declared values (`internal/filecontext/compact.go:489-507`).

Consequently `endLine`, `endCol`, `fileHash`, `targetRole`, `note`, `resolutionSource`, raw node ID, and edge fields are outside the file-detail DTO. This is the first shape reduction. It is not a missing raw record, and the full-row compact default plus P1-D's explicit high expanded cap rule out a sample-cap loss for the seven facts.

Classification: **bounded valid cardinality; explicit reduced projection shape**.

## 5. Database/Cypher changes nullability and scalar type

### 5.1 Absent `step` becomes `0`

Every selected raw `HAS_RESOLUTION_GAP` relationship omits the JSON `step` key, which corresponds to `Relationship.Step == nil`. `relationshipCSVRow` initializes `step := 0`, overwrites it only when the pointer is non-nil, and always writes the decimal string to the relationship CSV (`internal/lbugload/csv.go:372-387`). The single `CodeRelation.step INT32` column then contains `0` for those non-process edges.

The actual target Cypher capture returns `step: "0"` for all seven rows. The first divergence from raw nullability is therefore the CSV/database projection, before the query is executed.

Classification: **confirmed bounded representation conflation (`nil` and explicit zero become indistinguishable)**. It does not create an extra edge or source site.

### 5.2 Native query scalars become strings

`nativeResult.Rows` calls `tupleRow`; `tupleRow` calls `tupleString` for every selected column and stores that string in `lbugruntime.Row` (`internal/lbugnative/native_ladybugdb.go:95-110,132-165`). Non-string Ladybug values are rendered with `lbug_value_to_string`, so `DOUBLE 1` becomes `"1.000000"`, `INT32 0` becomes `"0"`, and numeric node fields are also strings. `mcpRowsFromLbugRows` copies those values without retyping (`internal/mcp/tools.go:570-601`).

This exactly explains P1-D's actual CLI JSON output. The raw graph retains numeric `confidence: 1`; Cypher returns `confidence: "1.000000"`.

Classification: **confirmed bounded scalar-type normalization at the native read adapter**. No row/cardinality loss was observed. The current Cypher tool contract does not state whether JSON scalar types or nullable `step` must be preserved, so this report does not escalate the representation differences into a global product-contract verdict.

## Owner file-detail and impact evidence

The refreshed Anvien self-graph was used only for owner inspection. Important blast-radius warnings are:

| Owner symbol | Risk | Impacted | Modules | Processes | Interpretation |
|---|---:|---:|---:|---:|---|
| `contextNeighborhood` | CRITICAL | 1 | 1 | 6 | central context command path |
| `runImpactBFSProfiled` | CRITICAL | 4 | 1 | 14 | central impact traversal |
| `impactAffectedProcesses` | CRITICAL | 4 | 1 | 13 | process summary path |
| `Builder.buildUnresolved` | CRITICAL | 6 | 2 | 23 | file-detail unresolved projection |
| `compactContextBuilder.unresolvedRows` | CRITICAL | 4 | 1 | 6 | compact DTO encoding |
| `relationshipCSVRow` | CRITICAL | 3 | 2 | 8 | all database relationship loads |
| `tupleRow` | LOW | 1 | 1 | 0 | native query row conversion |
| `tupleString` | LOW | 2 | 1 | 0 | native value conversion |
| `processes.Apply` | CRITICAL | 5 | 4 | 39 | global derived process construction |

These are blast-radius warnings, not edit prohibitions. No edit is authorized or performed in this investigation.

## Verification

Targeted existing controls passed:

```text
ok  github.com/tamnguyendinh/anvien/internal/mcp
ok  github.com/tamnguyendinh/anvien/internal/filecontext
ok  github.com/tamnguyendinh/anvien/internal/lbugload
```

The tests cover command exposure of present symbol/process/gap edges, compact row/cardinality behavior, and relationship CSV projection. They validate the traced mechanisms; they do not validate the target's missing upstream calls or establish process completeness.

## Evidence

- `E2-P2C-BOUNDARY1`: target HEAD/status, graph hash/count/last-write recheck; no target analyze or target write.
- `E2-P2C-TARGET1`: `E:\Anvien\.tmp\cheapapp-graph-root-cause-restart\p2c\target_command_boundary_probe.json` — selected raw relationships and zero selected process edges.
- `E2-P2C-CMD1`: P1-E command capture `E:\Anvien\.tmp\cheapapp-graph-root-cause-restart\p1e-command-captures.json`, independently reconciled to raw relationships and current command source.
- `E2-P2C-PROJ1`: P1-D raw/projection captures, especially `p1d-raw-graph-probe-output.json`, `p1d-parity-probe-output.json`, and `p1d-cypher-gap-rels.txt`.
- `E2-P2C-SRC1`: `contextNeighborhood`, `runImpactBFSProfiled`, `impactAffectedProcesses`, `buildUnresolved`, and compact-row source trace.
- `E2-P2C-SRC2`: `relationshipCSVRow` nil-to-zero and native `tupleRow`/`tupleString` scalar-to-string source trace.
- `E2-P2C-DERIVED1`: `E:\Anvien\.tmp\cheapapp-graph-root-cause-restart\p2c\process_control_probe.go` and `process_control_probe_output.json`.
- `E2-P2C-DERIVED2`: five independent current reruns under `p2c\supervisor\process-run-1.json` through `process-run-5.json` reproduce the same `3771/662/2761/0/0` versus `3773/662/2761/0/0` values and retain the same target graph hash.
- `E2-P2C-IMPACT1`: `E:\Anvien\.tmp\cheapapp-graph-root-cause-restart\p2c\owner-evidence\impact2_*.json` plus owner file-detail captures.
- `E2-P2C-TEST1`: `test_mcp_projection.txt`, `test_filecontext_projection.txt`, and `test_lbugload_projection.txt`.
- `E2-P2C-REPORT1`: this report.

## Residual boundary

- The report does not decide whether a particular product flow should exist as an Anvien Process; there is no accepted product authority for that claim, and the producer is heuristic/capped.
- It does not claim all `context`, `impact`, Cypher, file-detail, process, community, or semantic output is complete.
- The `step` and scalar-type findings are representation-level differences only; remediation priority and public-contract requirements remain for a separate task.
- No production code, tests, target source, target graph, or target index was changed. No build, detect-changes, commit, or remediation was performed.
