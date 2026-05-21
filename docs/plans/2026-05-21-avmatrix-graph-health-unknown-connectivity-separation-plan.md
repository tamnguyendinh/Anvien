# AVmatrix Graph Health Unknown Connectivity Separation Plan

Date: 2026-05-21

Status: draft

Companion files:

- Benchmark ledger: [2026-05-21-avmatrix-graph-health-unknown-connectivity-separation-benchmark.md](2026-05-21-avmatrix-graph-health-unknown-connectivity-separation-benchmark.md)
- Evidence ledger: [2026-05-21-avmatrix-graph-health-unknown-connectivity-separation-evidence.md](2026-05-21-avmatrix-graph-health-unknown-connectivity-separation-evidence.md)

## Rules

1. Follow active workspace and repository instructions, including `AGENTS.md`, for implementation workflow. This plan records product work and validation; it does not replace those rules.
2. Use AVmatrix according to the active repo instructions for implementation slices; do not use AVmatrix for doc-only commits.
3. As each implementation task is completed, update the corresponding checklist item immediately.
4. Run the full build before testing. For this plan the full build gate is `powershell -ExecutionPolicy Bypass -File avmatrix-launcher\build.ps1`.
5. Because this changes Web UI behavior and graph-health semantics, validation must include backend/unit coverage plus Web e2e coverage after the full build passes.
6. Record benchmark results as each benchmarkable task is completed. Benchmarkable means graph inventory counts, diagnostic counts, graph/API throughput, runtime/package size, or user-visible graph counts.
7. Record evidence as each evidenced task is completed.
8. After each completed implementation slice, commit the work, then continue until the full plan is complete.

## Problem

The Web UI Graph Health filter currently shows a very large `Unknown` bucket. On the current `E:\AVmatrix-GO` graph, investigation found `4345` nodes in that bucket. Turning off this bucket visually removes a large part of the graph, including many connected "islands".

This is not a small count and cannot be treated as cosmetic.

The current behavior makes `Unknown` ambiguous:

- it may contain real dead or unwired code candidates;
- it may contain connected code with analyzer uncertainty;
- it may contain source-backed unresolved references to Go builtins, standard library APIs, external library types, or test framework helpers;
- it may contain real in-repo resolution misses that should be fixed by the analyzer.

The user-visible problem is that `Unknown` looks like a topology class, but the implementation currently uses it as a diagnostic override. That hides the real topology and makes the filter misleading.

## Root Finding

Current code in `internal/graphhealth/compute.go` assigns topology in this order:

```go
case hasUnresolvedDiagnostic:
    status = TopologyUnknown
case component.Detached:
    status = TopologyDetached
case in > 0 && out > 0:
    status = TopologyConnected
```

This means any node with `unresolved_reference` diagnostics becomes `unknown_connectivity` before the code checks whether it is connected, no-incoming, no-outgoing, isolated, or detached.

Current baseline confirms the problem:

- `4345` nodes have `graphHealthDiagnostics`.
- `4345` nodes are therefore effectively shown as `Unknown`.
- `2036` of those diagnostic nodes still have both counted incoming and counted outgoing edges.
- Top unresolved targets include `testing.T`, `make`, `string`, `len`, `append`, `fmt.Errorf`, `time.Second`, `t.Helper`, and `t.TempDir`.

Therefore current `Unknown` is mostly "has resolution diagnostic", not "cannot classify topology".

## Design Decision

Separate graph topology from diagnostic uncertainty.

Graph Health must preserve independent fields:

```text
topologyStatus:
  connected | true_isolated | no_incoming | no_outgoing | detached_component | unknown_connectivity

diagnostics:
  unresolved_reference and future diagnostic evidence

confidence:
  candidate | expected | unknown | confirmed
```

Rules:

- `topologyStatus` must be derived from counted-edge topology and component/root reachability.
- `diagnostics` must not overwrite `topologyStatus`.
- `confidence: unknown` may mark analyzer/resolution uncertainty while preserving topology.
- `unknown_connectivity` must mean topology truly cannot be classified, not merely "this node has an unresolved reference".
- For a valid in-memory graph node with a known ID and computed counted-degree/component metadata, topology is classifiable. In that case `unknown_connectivity` must not be emitted.
- `unknown_connectivity` is reserved for malformed/incomplete graph payloads or future graph-health inputs where counted-degree/component computation cannot be completed.
- Dead-code triage must come from topology candidates such as `no_incoming`, `true_isolated`, and `detached_component`, with expected-isolated and diagnostic overlays considered separately.
- Builtin, standard-library, external-library, generated, and test-framework unresolved references must be classified or de-emphasized before being presented as actionable graph-health uncertainty.

## Non-Goals

- Do not delete application code based on this plan.
- Do not call `Unknown` a dead-code verdict.
- Do not hide analyzer misses by removing diagnostics from summaries.
- Do not synthesize fake in-repo nodes or fake edges only to reduce the `Unknown` count.
- Do not collapse `diagnostics`, `confidence`, and `topologyStatus` into one UI filter.
- Do not change primary semantic node labels such as `Function`, `File`, `Method`, `Struct`, or `Property`.
- Do not introduce product/runtime timeout or elapsed-time budget behavior.

## Target Design

### Backend Graph Health

- Compute counted incoming/outgoing degrees and component reachability first.
- Assign `topologyStatus` from topology only.
- Attach `diagnostics` independently.
- Set `confidence: unknown` when diagnostics indicate analyzer uncertainty.
- Keep `expectedIsolationReasons` as overlays, not topology replacements.
- Keep graph-level diagnostic summary counts.

Example target behavior:

| Node evidence | Target topologyStatus | Target confidence | Target diagnostics |
| --- | --- | --- | --- |
| in > 0 and out > 0, no diagnostics | `connected` | `candidate` or `expected` | none |
| in > 0 and out > 0, unresolved diagnostics | `connected` | `unknown` | present |
| in == 0 and out > 0, unresolved diagnostics | `no_incoming` | `unknown` | present |
| detached component with unresolved diagnostics | `detached_component` | `unknown` | present |
| topology cannot be computed from malformed graph data | `unknown_connectivity` | `unknown` | present if available |

### Diagnostic Classification

The current largest unresolved targets strongly suggest unresolved references are over-counting known non-actionable references. The implementation must classify at least:

- Go builtins and predeclared types/functions such as `make`, `len`, `append`, `string`, `int`, `uint`, `byte`, `float64`, `any`, and composite type text such as `map[string]any`.
- Go standard library package/type/member references such as `testing.T`, `testing.B`, `fmt.Errorf`, `time.Second`, `filepath.Join`, `strings.TrimSpace`, and test helper methods reached through `*testing.T`.
- external package/library references that are outside the analyzed repository.
- unresolved in-repo references that are still actionable analyzer or graph-accuracy candidates.

Each unresolved diagnostic must carry classification metadata when classification is implemented:

```text
classification:
  builtin | standard_library | test_framework | external_library | in_repo_unresolved | unclassified

actionability:
  non_actionable | review | analyzer_gap
```

Classification rules:

- `builtin`, `standard_library`, and `test_framework` diagnostics default to `actionability: non_actionable` unless an implementation-specific reason says otherwise.
- `external_library` diagnostics default to `actionability: review`, because the external reference may be expected but may still reveal missing import/package modeling.
- `in_repo_unresolved` diagnostics default to `actionability: analyzer_gap`.
- `unclassified` diagnostics default to `actionability: review` until a later classifier proves they are non-actionable or analyzer gaps.
- Classification must not create fake graph edges or fake in-repo target nodes.

The plan does not require all external references to become graph edges. It requires the diagnostic surface to stop treating recognized builtin/stdlib/external cases as topology uncertainty.

### Web UI

- The `Graph Health` section must make topology filters and diagnostic filters visually and semantically separate.
- `Unknown` under topology must not represent all unresolved diagnostics.
- If a node is `connected` with diagnostics, turning off `Unknown` topology must not hide that node.
- The UI may expose a separate diagnostic filter such as `Unresolved reference` or `Resolution uncertainty`, but that filter must be clearly diagnostic, not dead-code topology.
- Node detail must show both topology and diagnostic evidence when both exist.
- Hiding `Unknown` topology must only hide nodes whose actual `topologyStatus` is `unknown_connectivity`; it must not hide all nodes with diagnostic uncertainty.
- If diagnostic classification is available, the UI must expose or display enough detail for users to distinguish non-actionable builtin/stdlib/test diagnostics from in-repo analyzer gaps.

### Reports and API

- `/api/graph` must return per-node graph health where topology and diagnostics are independent.
- `/api/graph/report` must not rank `unresolved_reference` as a replacement topology.
- Connected nodes with diagnostics must not be ranked as topology defects. They may appear in a separate diagnostic section/report, or as diagnostic evidence attached to their real topology.
- Summary counts must include:
  - topology status counts;
  - confidence counts;
  - diagnostic counts;
  - diagnostic classification counts when implemented.

## Acceptance Criteria

- A node with counted incoming and counted outgoing edges plus unresolved diagnostics remains `topologyStatus: connected`.
- A node with `no_incoming` topology plus unresolved diagnostics remains `topologyStatus: no_incoming`.
- `confidence: unknown` and `diagnostics` remain visible without overwriting topology.
- Current `E:\AVmatrix-GO` `unknown_connectivity` count is no longer equal to the `graphHealthDiagnostics` node count.
- Current `E:\AVmatrix-GO` connected-with-diagnostics nodes are still visible when only `Unknown` topology is hidden.
- Builtin and standard-library unresolved references are classified or excluded from actionable topology uncertainty.
- Diagnostic classification includes `classification` and `actionability` fields, or an equivalent typed representation with the same values and semantics.
- Graph Health UI lets the user distinguish:
  - actual topology candidates;
  - expected-isolated nodes;
  - analyzer/resolution uncertainty;
  - connected nodes that merely carry diagnostics.
- `/api/graph/report` does not treat connected diagnostic nodes as dead-code/unwired topology candidates.
- Backend tests cover topology-plus-diagnostic combinations.
- Web e2e covers hiding `Unknown` topology without hiding connected diagnostic nodes.
- Full build passes before tests.
- Full relevant unit tests and Web e2e tests pass.
- Benchmark ledger records before/after counts for topology, diagnostics, confidence, and major unresolved target classes.

## Current Code Facts

- `internal/graphhealth/policy.go` defines:
  - `TopologyUnknown = "unknown_connectivity"`;
  - `ConfidenceUnknown = "unknown"`;
  - `DiagnosticUnresolvedReference = "unresolved_reference"`.
- `internal/graphhealth/compute.go` currently checks `hasUnresolvedDiagnostic` before topology cases, causing diagnostic evidence to overwrite topology.
- `internal/graphhealth/diagnostics.go` appends source-backed unresolved diagnostics to real graph nodes under `graphHealthDiagnostics`.
- `internal/resolution/emit.go` emits `unresolved_reference` diagnostics from unresolved call, access, type-reference, and heritage facts.
- `internal/httpapi/graph.go` uses `graphhealth.ComputeSummary(g)` for `/api/graph`, `/api/graph/report`, and graph-health explain/report surfaces.
- `avmatrix-web/src/lib/graph-health-filters.ts` labels `unknown_connectivity` as `Unknown` and uses topology filters independently from diagnostic kind filters, but backend data currently conflates the two.
- `avmatrix-web/src/components/FileTreePanel.tsx` exposes Graph Health filters and counts from graph payload node metadata.
- Existing evidence shows `4345` diagnostic nodes on `E:\AVmatrix-GO`, with `2036` still connected by counted edges.

## Implementation Slices

### P0 - Baseline and Reproduction

- [ ] [P0-A] Reproduce current `/api/graph` Graph Health summary on a freshly analyzed `E:\AVmatrix-GO` graph.
- [ ] [P0-B] Record raw `.avmatrix\graph.json` diagnostic-node inventory and counted-edge degree buckets.
- [ ] [P0-C] Record top unresolved target texts, fact families, files, node labels, and path buckets.
- [ ] [P0-D] Record the old connected-plus-diagnostic bug in evidence, then add target-behavior coverage that a connected node with unresolved diagnostics remains `connected` with `confidence: unknown`.
- [ ] [P0-E] Record baseline counts in the benchmark ledger before any semantic change.

### P1 - Preserve Topology When Diagnostics Exist

- [ ] [P1-A] Update `internal/graphhealth.ComputeSummary` so unresolved diagnostics do not override topology.
- [ ] [P1-B] Keep `confidence: unknown` for nodes with unresolved diagnostics.
- [ ] [P1-C] Implement the rule that valid graph nodes do not emit `unknown_connectivity`; reserve it for malformed/incomplete graph-health inputs only.
- [ ] [P1-D] Update graph-health unit tests for all topology statuses with and without diagnostics.
- [ ] [P1-E] Update HTTP/API tests that assert graph payload, report, and explain output.
- [ ] [P1-F] Record after-counts for topology statuses, confidence levels, and diagnostics.

### P2 - Classify Non-Actionable Unresolved References

- [ ] [P2-A] Identify where Go builtin, predeclared type, stdlib, and external-library references can be classified without inventing fake graph edges.
- [ ] [P2-B] Add diagnostic classification for recognized builtin/predeclared references.
- [ ] [P2-C] Add diagnostic classification for recognized Go standard-library and test-framework references.
- [ ] [P2-D] Add diagnostic classification fields for `classification` and `actionability`, or an equivalent typed representation with the same values.
- [ ] [P2-E] Preserve unresolved in-repo references as actionable analyzer/graph-accuracy diagnostics.
- [ ] [P2-F] Add tests for the top observed cases: `testing.T`, `make`, `len`, `append`, `string`, `int`, `fmt.Errorf`, `time.Second`, `t.Helper`, and `t.TempDir`.
- [ ] [P2-G] Record before/after top unresolved target counts and diagnostic classification counts.

### P3 - Report and API Semantics

- [ ] [P3-A] Update `/api/graph/report` ranking so `unresolved_reference` is not a topology replacement.
- [ ] [P3-B] Keep diagnostic candidates visible as diagnostic evidence.
- [ ] [P3-C] Ensure graph-health explain output for a connected diagnostic node shows `connected` plus diagnostic details.
- [ ] [P3-D] Ensure connected diagnostic nodes are not ranked as dead-code/unwired topology candidates.
- [ ] [P3-E] Add or update report/explain tests for topology plus diagnostic overlay and diagnostic-only triage.
- [ ] [P3-F] Record API payload evidence.

### P4 - Web UI Clarity

- [ ] [P4-A] Review `Graph Health` UI wording so `Unknown` clearly means topology unknown, not "has unresolved reference".
- [ ] [P4-B] Ensure diagnostic filters remain separate from topology filters.
- [ ] [P4-C] Add e2e coverage where hiding `Unknown` topology does not hide a connected node with diagnostics.
- [ ] [P4-D] Add e2e or unit coverage for node detail showing topology and diagnostics together.
- [ ] [P4-E] Record UI evidence and screenshots/artifacts if generated by Playwright.

### P5 - Full Validation and Closure

- [ ] [P5-A] Run full build first: `powershell -ExecutionPolicy Bypass -File avmatrix-launcher\build.ps1`.
- [ ] [P5-B] Run focused backend tests for graph health, resolution diagnostics, HTTP graph/report/explain, and contracts as touched.
- [ ] [P5-C] Run full relevant Go tests for touched packages.
- [ ] [P5-D] Run focused Web unit tests for graph-health filters/details.
- [ ] [P5-E] Run full Web unit tests.
- [ ] [P5-F] Run Web e2e covering Graph Health filters and diagnostic/topology separation.
- [ ] [P5-G] Re-run `avmatrix analyze --force` and record final inventory counts.
- [ ] [P5-H] Run required change detection before commit according to active repo instructions.
- [ ] [P5-I] Commit the completed implementation slice.

## Open Questions

- Should recognized builtin/stdlib diagnostics be retained as low-priority per-node diagnostics, or moved only into aggregate summary counts after classification?
- Should Web UI default-hide diagnostic uncertainty, or keep it visible but separate from topology?

The `unknown_connectivity` meaning is no longer open: it is reserved for malformed/incomplete graph-health inputs, not valid nodes with unresolved diagnostics.

The remaining questions must be resolved in the implementation evidence before closing P2/P4.
