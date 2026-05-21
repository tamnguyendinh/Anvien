# AVmatrix Graph Health Unknown Connectivity Separation Evidence Ledger

Date: 2026-05-21

Status: draft

Plan: [2026-05-21-avmatrix-graph-health-unknown-connectivity-separation-plan.md](2026-05-21-avmatrix-graph-health-unknown-connectivity-separation-plan.md)

Benchmark: [2026-05-21-avmatrix-graph-health-unknown-connectivity-separation-benchmark.md](2026-05-21-avmatrix-graph-health-unknown-connectivity-separation-benchmark.md)

## E0 - User Problem Statement

Status: recorded

User observed that the `Graph Health` filter shows a very large `Unknown` node count. Turning off `Unknown` hides a large portion of the graph, including major connected graph islands.

User clarification:

- The large `Unknown` group may contain both misclassified analyzer uncertainty and real dead-code candidates.
- The issue must be analyzed, not hidden.
- The graph visibly shows many wires connected to nodes in the `Unknown` group, so `Unknown` cannot simply mean no connectivity.
- The desired outcome is to find out whether these are dead code or classification/resolution problems.

Interpretation:

- `Unknown` is currently too broad and misleading as a topology filter.
- The implementation must split topology from diagnostic uncertainty.

## E1 - Source Inspection Evidence

Status: recorded

Files inspected:

- `internal/graphhealth/policy.go`
- `internal/graphhealth/compute.go`
- `internal/graphhealth/diagnostics.go`
- `internal/resolution/emit.go`
- `internal/resolution/resolve.go`
- `internal/httpapi/graph.go`
- `avmatrix-web/src/lib/graph-health-filters.ts`
- `avmatrix-web/src/components/FileTreePanel.tsx`

Relevant findings:

- `internal/graphhealth/policy.go` defines `unknown_connectivity`, `unknown` confidence, and `unresolved_reference`.
- `internal/graphhealth/compute.go` checks `hasUnresolvedDiagnostic` before topology checks.
- Because of that order, any node with unresolved diagnostics becomes `TopologyUnknown`.
- `internal/resolution/emit.go` emits unresolved diagnostics for unresolved call, access, type-reference, and heritage facts.
- `internal/graphhealth/diagnostics.go` attaches those diagnostics to real graph nodes under `graphHealthDiagnostics`.
- `internal/httpapi/graph.go` computes graph health before returning `/api/graph`, report, and explain payloads.
- Web filters are consuming backend-derived topology data. The Web UI is not inventing the large `Unknown` count.

Conclusion:

- The primary issue is backend Graph Health semantics, not the left dashboard filter implementation.

## E2 - Baseline Inventory Evidence

Status: recorded

Command context:

- Repository: `E:\AVmatrix-GO`
- Graph refresh before inspection: `avmatrix analyze --force`
- Result: `files: scanned=721 parsed=539 unsupported=182 failed=0`
- Graph: `nodes=21941`, `relationships=54489`

Raw graph inspection showed:

- `.avmatrix\graph.json` itself does not contain computed `graphHealth` fields until the API/Graph Health compute path runs.
- It does contain source-backed `graphHealthDiagnostics` fields from resolution.
- Nodes with `graphHealthDiagnostics`: `4345`

This matches the user-observed `Unknown` scale because current backend compute turns unresolved diagnostics into `unknown_connectivity`.

## E3 - Diagnostic Node Degree Evidence

Status: recorded

Using the accepted Graph Health counted-edge policy from `internal/graphhealth/policy.go`, the diagnostic nodes break down as:

| Counted-edge bucket | Node count |
| --- | ---: |
| Connected both incoming and outgoing | 2036 |
| No incoming, has outgoing | 1581 |
| Has incoming, no outgoing | 559 |
| No incoming and no outgoing | 169 |

Conclusion:

- At least `2036` nodes currently classified as `Unknown` are actually connected by counted graph edges.
- Therefore `Unknown` is not a reliable dead-code or no-connectivity signal.

## E4 - Diagnostic Label and Source Evidence

Status: recorded

Diagnostic nodes by semantic label:

| Label | Count |
| --- | ---: |
| Function | 3057 |
| Method | 738 |
| Struct | 292 |
| Variable | 166 |
| File | 22 |
| Property | 21 |
| Package | 18 |
| Const | 10 |
| Interface | 7 |
| Constructor | 6 |
| Class | 4 |
| TypeAlias | 4 |

Diagnostic nodes by path bucket:

| Path bucket | Count |
| --- | ---: |
| Go source | 2577 |
| Test | 1627 |
| Web source | 141 |

Diagnostic buckets:

| Diagnostic bucket | Count |
| --- | ---: |
| `unresolved_reference|call|scope-resolution|scope-resolution` | 3786 |
| `unresolved_reference|type-reference|scope-resolution|scope-resolution` | 3129 |
| `unresolved_reference|access|scope-resolution|scope-resolution` | 1943 |
| `unresolved_reference|heritage|scope-resolution|scope-resolution` | 7 |

Conclusion:

- The largest source is scope-resolution unresolved references, not a graph layout or UI-only problem.
- Tests and Go source account for most of the affected nodes.

## E5 - Top Unresolved Target Evidence

Status: recorded

Top unresolved target texts:

| Target | Count |
| --- | ---: |
| `type-reference:testing.T` | 1123 |
| `type-reference:collector` | 421 |
| `call:t.Helper` | 342 |
| `type-reference:int` | 336 |
| `call:make` | 293 |
| `call:string` | 196 |
| `call:t.TempDir` | 191 |
| `call:len` | 180 |
| `call:node.Kind` | 132 |
| `call:c.text` | 111 |
| `call:append` | 108 |
| `access:time.Second` | 104 |
| `call:t.Fatalf` | 100 |
| `call:strings.TrimSpace` | 89 |
| `call:uint` | 88 |
| `type-reference:Server` | 78 |
| `type-reference:context.Context` | 72 |
| `type-reference:map[string]any` | 51 |
| `call:int` | 49 |
| `type-reference:byte` | 48 |
| `type-reference:testing.B` | 48 |
| `call:filepath.Join` | 47 |
| `type-reference:float64` | 45 |
| `access:result.Graph` | 44 |
| `access:result.Metrics` | 42 |

Interpretation:

- Many top unresolved references are Go builtins, predeclared types, standard library symbols, or test helpers.
- Those references are not evidence of dead code.
- Some targets such as local receiver/member expressions may still indicate analyzer gaps and must remain visible as diagnostics.

## E6 - Current Conclusion

Status: recorded

The current `Unknown` bucket is a mixed bucket caused by semantic conflation:

- topology status is overwritten by diagnostic presence;
- `Unknown` currently means "has unresolved diagnostics" for many nodes;
- connected nodes are being hidden by the `Unknown` topology filter;
- unresolved diagnostics include non-actionable builtin/stdlib/test references and potentially actionable analyzer misses.

Required correction:

- preserve topology first;
- retain diagnostics separately;
- use `confidence: unknown` to communicate uncertainty;
- classify non-actionable unresolved references;
- keep dead-code triage on actual topology candidates.

## E7 - Plan Review Corrections

Status: recorded

Plan review found that the first draft was directionally correct but needed tighter execution rules.

Corrections made:

- Removed ambiguity around `unknown_connectivity`: valid graph nodes with computed counted-degree/component metadata must not emit `unknown_connectivity`; that status is reserved for malformed or incomplete graph-health inputs.
- Reworded P0-D so implementation does not create a regression test that locks the old bug. Evidence should record the old bug; tests should assert target behavior.
- Added explicit diagnostic classification metadata:
  - `classification`: `builtin`, `standard_library`, `test_framework`, `external_library`, `in_repo_unresolved`, `unclassified`;
  - `actionability`: `non_actionable`, `review`, `analyzer_gap`.
- Clarified that P1 delivers the primary user value by preserving topology even before diagnostic classification is complete.
- Clarified that P2 handles classification of builtin/stdlib/test/external diagnostics.
- Clarified that hiding `Unknown` topology must not hide all diagnostic nodes.
- Clarified `/api/graph/report` semantics: connected diagnostic nodes must not be ranked as dead-code or unwired topology candidates.

## E8 - Implementation Evidence

Status: pending

Record code changes by slice.

## E9 - Validation Evidence

Status: pending

Record full build, backend tests, Web unit tests, Web e2e tests, graph refresh, final counts, and required change detection.

## E10 - Commit Evidence

Status: pending

Record implementation commits after each completed slice.
