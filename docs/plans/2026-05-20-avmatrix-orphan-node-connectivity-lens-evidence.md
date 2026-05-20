# AVmatrix Orphan Node Connectivity Lens Evidence Ledger

Date: 2026-05-20

Status: active

Companion files:

- Plan: [2026-05-20-avmatrix-orphan-node-connectivity-lens-plan.md](2026-05-20-avmatrix-orphan-node-connectivity-lens-plan.md)
- Benchmark ledger: [2026-05-20-avmatrix-orphan-node-connectivity-lens-benchmark.md](2026-05-20-avmatrix-orphan-node-connectivity-lens-benchmark.md)

## Evidence Rules

Record evidence as each evidenced task is completed. Evidence should include commands, impacted files, test results, e2e artifacts, and concise observations needed to audit the plan later.

For doc-only commits, do not use AVmatrix.

Do not record inferred graph counts. Every count must include the command, source graph, repo path, commit or graph timestamp when available, and interpretation.

## E0 - Plan Creation Evidence

Date: 2026-05-20

Status: recorded

Created file set:

- `docs/plans/2026-05-20-avmatrix-orphan-node-connectivity-lens-plan.md`
- `docs/plans/2026-05-20-avmatrix-orphan-node-connectivity-lens-evidence.md`
- `docs/plans/2026-05-20-avmatrix-orphan-node-connectivity-lens-benchmark.md`

Plan creation scope:

- Use the repository's established plan, evidence, and benchmark ledger format.
- Ground implementation work in codebase-reviewed facts.
- Leave unknown graph-health counts pending until measured.

Convention inspection commands:

```powershell
Get-ChildItem docs\plans | Sort-Object Name | Select-Object -Last 20 | Format-Table -AutoSize
rg -n "^# |^Status:|^## |Acceptance|Validation|Closure|Evidence|Benchmark|Zero-Trust|Phase" docs\plans
Get-Content docs\plans\2026-05-19-avmatrix-multilanguage-graph-filters-and-coverage-plan.md -TotalCount 280
Get-Content docs\plans\2026-05-19-avmatrix-multilanguage-graph-filters-and-coverage-evidence.md -TotalCount 160
Get-Content docs\plans\2026-05-19-avmatrix-multilanguage-graph-filters-and-coverage-benchmark.md -TotalCount 160
```

Observed planning standard:

- Plan file has date, status, companion files, rules, problem/scope, acceptance guardrails, phase checklist, ledger, and closure definition.
- Evidence ledger records commands, files, status, observations, and validation artifacts.
- Benchmark ledger records measured counts and benchmarkable inventory; unmeasured values remain pending.

Plan creation decisions:

- Status is `active` because this is a formal plan for future implementation.
- No graph counts were invented.
- At initial creation, baseline graph-health counts were left `pending measurement`; E3/E4 supersede that gap with codebase review and initial measured baselines.
- "Orphan node" is defined as derived connectivity status, not a primary semantic node label.

## E1 - Initial Product Reasoning Evidence

Date: 2026-05-20

Status: recorded

Discussion summary:

- The product question is whether lonely/orphan nodes should be classified or mapped into a separate node/filter type to manage code, buggy functions, and unwired code.
- The accepted planning direction is to classify this as a graph-health/connectivity lens, not a semantic node label.
- The plan requires taxonomy and evidence before presenting any orphan status as a bug.

Current claim boundary:

- No current orphan counts are claimed.
- No current analyzer defect is claimed.
- No dead-code count is claimed.
- No UI implementation detail is considered accepted until inspected during implementation phases.

## E2 - Pending Baseline Evidence

Date: 2026-05-20

Status: partially superseded by E3/E4 initial `E:\AVmatrix-GO` measured baseline; representative cross-repo baseline pending

Required before implementation claims:

- measured connectivity inventory for `E:\AVmatrix-GO`;
- measured connectivity inventory for representative indexed repos selected by documented criteria when available;
- recorded edge policy used for the measurements;
- expected-isolated policy and count by reason;
- comparison of raw graph connectivity versus Web-visible connectivity.

## E3 - AVmatrix And Source Code Review Evidence

Date: 2026-05-20

Status: recorded

Reason:

The plan is grounded in AVmatrix and source inspection before it drives implementation.

AVmatrix index command:

```powershell
go run .\cmd\avmatrix analyze --force --skip-agents-md --no-stats
```

Result:

```text
analyzed E:\AVmatrix-GO
files: scanned=694 parsed=530 unsupported=164 failed=0
graph: nodes=20967 relationships=52302 path=E:\AVmatrix-GO\.avmatrix\graph.json
```

AVmatrix context commands:

```powershell
go run .\cmd\avmatrix context Graph --repo AVmatrix
go run .\cmd\avmatrix context knowledgeGraphToGraphology --repo AVmatrix
go run .\cmd\avmatrix context FileTreePanel --repo AVmatrix
go run .\cmd\avmatrix context GraphRelationshipTypes --repo AVmatrix
```

Observed AVmatrix/codebase targets:

- `internal/graph/types.go`
- `internal/httpapi/graph.go`
- `internal/contracts/web_ui.go`
- `internal/ignore/constants.go`
- `internal/processes/processes.go`
- `internal/graphaccuracy/property_access.go`
- `cmd/property-access-audit/main.go`
- `avmatrix-web/src/core/graph/types.ts`
- `avmatrix-web/src/services/backend-client.ts`
- `avmatrix-web/src/lib/constants.ts`
- `avmatrix-web/src/lib/graph-adapter.ts`
- `avmatrix-web/src/components/FileTreePanel.tsx`

Source inspection commands:

```powershell
Get-Content internal\graph\types.go
Get-Content internal\httpapi\graph.go | Select-Object -First 260
Get-Content internal\contracts\web_ui.go | Select-Object -First 220
Get-Content internal\ignore\constants.go | Select-Object -First 140
Get-Content internal\processes\processes.go | Select-Object -Skip 180 -First 320
Get-Content internal\graphaccuracy\property_access.go | Select-Object -First 280
Get-Content internal\graphaccuracy\property_access.go | Select-Object -Skip 380 -First 100
rg -n "property-access-audit|OrphanStatus|true_orphan|false_orphan|orphan\\.\\*" cmd internal
Get-Content avmatrix-web\src\core\graph\types.ts
Get-Content avmatrix-web\src\services\backend-client.ts | Select-Object -Skip 480 -First 100
Get-Content avmatrix-web\src\lib\constants.ts | Select-Object -First 480
Get-Content avmatrix-web\src\lib\graph-adapter.ts | Select-Object -First 460
Get-Content avmatrix-web\src\components\FileTreePanel.tsx | Select-Object -First 900
```

Codebase findings:

- `internal/graph/types.go` has no `connectivityStatus` metadata today.
- `internal/httpapi/graph.go` streams raw graph nodes/relationships; it strips node `content` by default but does not derive connectivity.
- `internal/contracts/web_ui.go` exposes semantic node labels and relationship types but no Graph Health contract.
- `avmatrix-web/src/core/graph/types.ts` models `KnowledgeGraph` as raw node and relationship arrays.
- `avmatrix-web/src/lib/constants.ts` owns node/edge labels, colors, sizes, filterable labels, edge display metadata, and grouped heritage compatibility counts.
- `avmatrix-web/src/components/FileTreePanel.tsx` currently renders Node Types, Edge Types, Focus Depth, and Color Legend. There is no Graph Health group.
- `avmatrix-web/src/lib/graph-adapter.ts` has a layout-only comment about "orphan nodes that weren't reached" after hierarchy BFS. This is not an orphan/dead-code taxonomy.
- `internal/processes/processes.go` already encodes relevant graph-health heuristics: process entrypoints are Function/Method, tests are excluded by `isTestFile`, calls under confidence `0.5` are ignored, and Route/Tool resources are linked to processes by `ENTRY_POINT_OF`.
- `internal/ignore/constants.go` contains existing path policy that should inform expected-isolated classification for vendor, generated, fixture, test, build, cache, and dependency directories.
- `internal/graphaccuracy/property_access.go` already emits a Property-specific `orphanStatus` taxonomy: `owner_linked`, `false_orphan`, `true_orphan`, `unknown`, `external_library_owned`, and `intentionally_unmodeled`.
- `cmd/property-access-audit` prints property audit counts as `orphan.*` summary lines. The Graph Health plan must keep the new node-connectivity taxonomy separate unless compatibility mapping is explicitly designed and tested.
- Scanner-ignored paths such as vendor, dependency, generated, fixture, build, and cache directories may never enter the graph. Expected-isolated policy must distinguish graph-present nodes from out-of-graph ignored inputs.

Conclusion:

The plan must not derive orphan status from raw incoming/outgoing edges. Structural edges such as `DEFINES`, `CONTAINS`, `HAS_METHOD`, `HAS_PROPERTY`, and `MEMBER_OF` materially change zero-incoming/zero-outgoing counts and can hide dead/unwired candidates if counted blindly.

## E4 - Initial Connectivity Baseline Commands

Date: 2026-05-20

Status: recorded

## E5 - Phase 1 Policy Decisions: Counted Edge + Expected-Isolated + Taxonomy + Ownership + Root Rules + Confidence

Date: 2026-05-20

Status: recorded (completes P1-A1..P1-I)

### Investigation Commands Used (all AVmatrix + source per AGENTS.md + plan rules)
- `go run ./cmd/avmatrix analyze --force --skip-agents-md --no-stats` (refreshed to nodes=21091, rels=52445)
- `avmatrix__list_repos`, `avmatrix__query` (multiple for "orphan nodes...", "knowledgeGraphToGraphology...", "graph connectivity")
- `avmatrix__context` on key symbols (Graph, FileTreePanel, etc. in prior E3)
- Source reads: internal/graph/types.go (all 22 Rel* consts), internal/processes/processes.go (isTestFile, ENTRY_POINT_OF emission, findEntryPoints, buildCallsGraph), internal/ignore/constants.go (full ignore sets), internal/contracts/web_ui.go (relationshipDisplayPolicies, graphRelationshipTypes list), avmatrix-web/src/lib/graph-adapter.ts (forward/reverseHierarchyRelations exactly matching structural set), internal/graphaccuracy/property_access.go (separate orphanStatus)
- Python graph.json loaders for exact degree counts on both E:\AVmatrix-GO and E:\Restaurant_manager (reproducible, no inference)
- `avmatrix__cypher` attempted (limited on Go adapter)

### Counted Edge Policy (P1-A1/A2/A3 - finalized)
See benchmark B0 for the exact table and rationale. In short: count only "wiring/usage/flow" edges; exclude the 5 structural ownership edges that the layout already treats separately and that empirically collapse all zero-incoming to zero. This was the key blocker identified in E3/E4.

Recorded in plan design section, benchmark, this evidence. No change to graph.Rel* consts or emission — pure derivation policy.

### Expected-Isolated Overlay Policy (P1-B1/B2/B3 - finalized)
Automatic reasons (path/label/evidence based, hide or de-emphasize by default in triage):
- `test`: isTestFile() patterns (exact copy of processes.go:465) OR /test/ /__tests__/ .test. .spec. _test.* ; also test helpers inside those files.
- `fixture`: /fixtures/ /__snapshots__/ /snapshots__ /testdata/
- `generated`: path contains /generated/ .generated. or parser properties indicating generated code.
- `vendor_dependency`: /vendor/ /node_modules/ /dist/ /build/ (even if scanner let a few through).
- `documentation`: label=="Section" OR *.md files OR /docs/ /README* that carry no CALLS/ACCESSES etc.
- `migration_script`: /migrations/ *.sql change files, db/migrate scripts.
- `cli_mcp`: symbols under internal/cli, internal/mcp, or registered via cmd/* or server tools.

Prioritization modifier only (never auto-expected-isolated, still shown in candidate lists but lower priority):
- `exported_api`: boolProperty "isExported"==true (from processes scoring). These may legitimately have zero internal incoming because they are the public surface for external callers or reflection. Triage must inspect callers outside repo + tests + docs.

Framework entry surfaces (roots, not candidates even if low degree):
- `framework_entry`: label in {Route, Tool} OR has outgoing ENTRY_POINT_OF / HANDLES_ROUTE / HANDLES_TOOL OR isProcessSymbol(main-like) that findEntryPoints would consider. These are the accepted roots for detached traversal.

Evidence rule: automatic reasons take precedence for hiding; exported + framework only affect sort order in reports ("triage these last"). A node can carry multiple reasons (e.g. exported test helper).

This policy bridges the existing ignore + processes.isTestFile + exported scoring without inventing new scanners.

### Topology Taxonomy (P1-F)
Accepted exactly as proposed in plan §Design Decision (connected / true_isolated / no_incoming / no_outgoing / detached_component / unknown_connectivity). No changes. `no_outgoing` remains low-priority (normal leaves).

### Confidence Levels (P1-C)
- `candidate`: isolated topology status + zero expected reasons + zero source-backed diagnostics → actionable triage item.
- `expected`: any expected-isolated reason present (even if topologically isolated).
- `unknown`: unresolved_reference with source node evidence, or external-only targets dominate, or policy edge case.
- `confirmed`: human + additional runtime/test evidence only; never auto-set by derivation.

TopologyStatus and expectedIsolatedReasons and diagnostics are orthogonal fields on the derived metadata; they coexist.

### Root Surfaces + Traversal for detached_component (P1-G)
Accepted roots (start traversal from these):
- All Process nodes
- All nodes that are source of at least one ENTRY_POINT_OF, HANDLES_ROUTE, HANDLES_TOOL relationship
- All nodes with label Route or Tool
- Functions/Methods named "main", "init", "run", "start", "bootstrap" (case-insensitive) that have isExported or high frameworkMultiplier
- MCP tool registrations and CLI command entry symbols

Traversal: directed outgoing along counted edges (CALLS + process links + heritage for type reachability) from the roots. A component (weakly connected subgraph under non-structural edges) that has internal edges but zero path from any accepted root → `detached_component`.

Directed chosen over undirected because call graph and process traces are directional; a "downstream only" module that nothing calls into is still detached if no root reaches it.

Undirected fallback only for pure structural cycles that survived filtering (rare).

### Metadata Ownership (P1-H)
Core graph layer (recommended path):
- New package `internal/graphhealth` (or extension inside `internal/graph`) will expose `AnnotateGraphHealth(g *graph.Graph, policy EdgePolicy, expectedPolicy)` that mutates/adds to each Node.Properties:
  - "topologyStatus": string
  - "countedIncoming": int
  - "countedOutgoing": int
  - "excludedEdgeCategories": map or array
  - "expectedIsolationReasons": []string
  - "diagnostics": [] {kind, evidence...}
  - "confidence": string
- Derivation runs once after full graph assembly (in analyze or http graph handler or dedicated MCP resource), before any consumer (Web, query, reports, cypher views).
- Benefits: MCP `context`/`query`/`impact` can surface health without Web-only derivation; consistent truth for all surfaces (P3 requirement); no duplication.
- Alternative (Web-only) rejected for this reason.

Implementation in Phase 2 will keep raw graph immutable; health is always derived view (or cached annotation).

### Compatibility with Existing orphanStatus (P1-F)
`cmd/property-access-audit` + `internal/graphaccuracy/property_access.go` "orphanStatus" (owner_linked / true_orphan / false_orphan / ...) remains Property-only, audit-only surface. New Graph Health taxonomy is node-global, topology+overlay, and explicitly namespaced. No overwrite, no shared enum. If future phase wants mapping, it will be a separate documented transform with tests.

### All Phase 1 Decisions Documented
- Plan updated with decisions.
- Benchmark B0 now authoritative with policy + numbers for both repos.
- This E5 + E3/E4 provide full audit trail.
- No implementation code changed in this slice (doc-only per plan rule 6).

Conclusion: Phase 1 complete. Ready for Phase 2 backend derivation without risk of ambiguous "orphan" meaning. All guardrails (no primary label change, candidate not verdict, explanations required) satisfied.

Raw all-relationship baseline command:

```powershell
$repos = @(@{Name='AVmatrix-GO'; Path='E:\AVmatrix-GO\.avmatrix\graph.json'})
$codeLabels = @('Class','Function','Method','Interface','Struct','Trait','Impl','TypeAlias','Enum','Record','Delegate','Constructor','Route','Tool')
# Count incoming/outgoing by all relationships, then summarize raw zero-incoming,
# zero-outgoing, zero-both, code-label subsets, path-policy candidates, node labels, and relationship types.
```

Raw all-relationship result summary:

```text
AVmatrix-GO:
nodes=20967 relationships=52302
rawZeroIncoming=20 rawZeroOutgoing=15216 rawZeroBoth=8
codeNodeCount=4929 codeZeroIncoming=0 codeZeroOutgoing=200 codeZeroBoth=0
pathExpectedCandidateNodes=5743 pathExpectedCandidateZeroBoth=0
```

Provisional non-structural policy command:

```powershell
$nonStructuralTypes = @('CALLS','INHERITS','METHOD_OVERRIDES','METHOD_IMPLEMENTS','IMPORTS','USES','DECORATES','IMPLEMENTS','EXTENDS','ACCESSES','STEP_IN_PROCESS','HANDLES_ROUTE','FETCHES','HANDLES_TOOL','ENTRY_POINT_OF','WRAPS','QUERIES')
$callGraphTypes = @('CALLS','HANDLES_ROUTE','HANDLES_TOOL','ENTRY_POINT_OF','STEP_IN_PROCESS')
# Count code/callable node incoming/outgoing using those relationship type sets only.
```

Provisional policy result summary:

```text
AVmatrix-GO:
nonStructural code nodes=4929 zeroIn=1616 zeroOut=1100 zeroBoth=133
callable flow nodes=4242 zeroIn=1587 zeroOut=1323 zeroBoth=264
```

Interpretation:

- Raw all-relationship counts are not suitable as the orphan/dead-code denominator because `DEFINES` and container/ownership edges give code nodes incoming edges.
- The provisional policy is not final product behavior. It is recorded only to prove why Phase 1 must define a counted-edge policy before implementation.
- Representative cross-repo baseline remains pending until Phase 1 records selection criteria.

## E6 - Phase 2/3 Backend Graph-Health Derivation Slice

Date: 2026-05-20

Status: recorded

Scope:

- Implement backend-owned graph-health derivation in `internal/graphhealth`.
- Emit per-node graph-health metadata in HTTP graph JSON and NDJSON node records.
- Emit graph-health summary in non-stream `/api/graph` JSON response.
- Generate explicit Web contract types for Graph Health metadata without adding Web UI filters yet.
- Leave detached-component traversal, source-backed unresolved diagnostics, node/component explain endpoint, and Web filter UI for later phases.

AVmatrix refresh and impact commands:

```powershell
go run .\cmd\avmatrix analyze --force --skip-agents-md --no-stats
go run .\cmd\avmatrix impact Compute --repo AVmatrix --direction upstream --depth 2 --include-tests
go run .\cmd\avmatrix impact graphPayload --repo AVmatrix --direction upstream --depth 2 --include-tests
go run .\cmd\avmatrix impact streamGraphNDJSON --repo AVmatrix --direction upstream --depth 2 --include-tests
go run .\cmd\avmatrix impact graphNodeForResponse --repo AVmatrix --direction upstream --depth 2 --include-tests
go run .\cmd\avmatrix impact WebUIContractTypeScript --repo AVmatrix --direction upstream --depth 2 --include-tests
go run .\cmd\avmatrix impact WebUIContract --repo AVmatrix --direction upstream --depth 2 --include-tests
```

Impact observations:

- `Compute`: LOW risk; direct test impact in `internal/graphhealth`.
- HTTP graph functions: CRITICAL risk because they feed `Server.handleGraph` and graph API consumers. Mitigation: no stream record shape change; focused JSON/NDJSON tests added.
- `WebUIContract` / `WebUIContractTypeScript`: CRITICAL risk because generated contract glue and contract tests depend on it. Mitigation: generated contracts refreshed and `--check` passed.

Changed files:

- `internal/graphhealth/policy.go`
- `internal/graphhealth/compute.go`
- `internal/graphhealth/compute_test.go`
- `internal/httpapi/graph.go`
- `internal/httpapi/handlers_test.go`
- `internal/contracts/web_ui.go`
- `internal/contracts/web_ui_test.go`
- `avmatrix-web/src/generated/avmatrix-contracts.ts`
- `docs/plans/2026-05-20-avmatrix-orphan-node-connectivity-lens-plan.md`
- `docs/plans/2026-05-20-avmatrix-orphan-node-connectivity-lens-evidence.md`
- `docs/plans/2026-05-20-avmatrix-orphan-node-connectivity-lens-benchmark.md`

Implementation notes:

- `graphhealth.ComputeSummary` annotates graph nodes in-place and returns summary counts.
- Per-node `properties.graphHealth` includes topology status, counted incoming/outgoing counts, excluded structural counts, expected-isolated reasons, diagnostics field, and confidence.
- Flat properties (`topologyStatus`, `countedIncoming`, `countedOutgoing`, `expectedIsolationReasons`, `confidence`) remain for simple consumers.
- `exported_api` alone remains `candidate`; automatic expected reasons and `framework_entry` produce `expected`.
- HTTP NDJSON keeps the existing record types (`node`, `relationship`, `error`) and only adds metadata inside node properties.
- Non-stream HTTP JSON adds top-level `graphHealth` summary.
- Generated Web contracts now include `GRAPH_HEALTH_TOPOLOGY_STATUSES`, `GRAPH_HEALTH_CONFIDENCE_LEVELS`, `GRAPH_HEALTH_EXPECTED_ISOLATION_REASONS`, `GraphHealthNodeMetadata`, `GraphHealthSummary`, and `GraphResponse`.

Validation commands and results:

```powershell
go test ./internal/graphhealth
go test ./internal/httpapi -run "TestGraph"
go test ./internal/contracts
go run .\cmd\generate-web-contracts
go build ./...
go build ./cmd/... ./internal/...
go test ./internal/graphhealth ./internal/httpapi ./internal/contracts
go run .\cmd\generate-web-contracts --check
cd avmatrix-web; npm run build
go test ./cmd/... ./internal/...
go run .\cmd\avmatrix analyze --force --skip-agents-md --no-stats
go run .\cmd\avmatrix detect-changes --repo AVmatrix --scope all
```

Results:

- Focused graphhealth/httpapi/contracts tests passed.
- `go run .\cmd\generate-web-contracts` refreshed `avmatrix-web/src/generated/avmatrix-contracts.ts`.
- `go run .\cmd\generate-web-contracts --check` passed.
- `go build ./...` failed on existing fixture packages under `avmatrix/test/fixtures/...` (`models`, `animal`, and C fixture source). This is outside the implementation slice and is why P6-A remains not fully complete.
- `go build ./cmd/... ./internal/...` passed.
- `go test ./cmd/... ./internal/...` passed.
- `npm run build` in `avmatrix-web` passed; Vite reported existing chunk-size/dynamic-import warnings only.
- Post-change AVmatrix refresh passed with `nodes=21233 relationships=52777`.
- `detect-changes` passed and reported `risk_level=high`, `changed_files=12`, `changed_count=79`, `affected_count=14`. The affected processes were expected graph API and contract generation flows, including `HandleGraph -> GraphResponse`, `HandleGraph -> AddNodeHealthToSummary`, and `WebUIContractTypeScript -> ...`.

Current unrelated worktree note:

- `.gitignore` has an existing local change adding `.grok/`. This slice did not depend on it and must not stage it unless separately requested.

Conclusion:

Backend/API/contract Graph Health derivation is implemented for counted-edge topology and expected-isolated overlays. The slice does not complete detached-component traversal, unresolved-reference diagnostics, report/explain endpoints, or Web UI filters.
