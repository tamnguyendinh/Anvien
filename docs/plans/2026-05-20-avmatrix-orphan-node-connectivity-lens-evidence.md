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

User requirement:

- Create a proper repo-standard plan, not a chat-only outline.
- Do not create a speculative plan.
- Follow the repository's established planning rules and required companion ledgers.

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
- Doc-only commit mechanics do not require AVmatrix, but plan creation and material plan updates require AVmatrix/codebase inspection when the plan defines implementation work.

Plan creation decisions:

- Status is `active` because the user requested a formal plan for future implementation.
- No graph counts were invented.
- At initial creation, baseline graph-health counts were left `pending measurement`; E3/E4 supersede that gap with codebase review and initial measured baselines.
- "Orphan node" is defined as derived connectivity status, not a primary semantic node label.

Correction note:

- The first committed plan was structurally correct but insufficient because it did not use AVmatrix/codebase inspection before defining work. This evidence ledger now records the required codebase review and measured baseline used to correct the plan.

## E1 - Initial Product Reasoning Evidence

Date: 2026-05-20

Status: recorded

Discussion summary:

- A user asked whether lonely/orphan nodes should be classified or mapped into a separate node/filter type to manage code, buggy functions, and unwired code.
- The accepted planning direction is to classify this as a graph-health/connectivity lens, not a semantic node label.
- The plan requires taxonomy and evidence before presenting any orphan status as a bug.

Non-speculative boundary:

- No current orphan counts are claimed.
- No current analyzer defect is claimed.
- No dead-code count is claimed.
- No UI implementation detail is considered accepted until inspected during implementation phases.

## E2 - Pending Baseline Evidence

Date: 2026-05-20

Status: superseded by E3 initial measured baseline

Required before implementation claims:

- measured connectivity inventory for `E:\AVmatrix-GO`;
- measured connectivity inventory for one large indexed repo when available;
- recorded edge policy used for the measurements;
- expected-isolated policy and count by reason;
- comparison of raw graph connectivity versus Web-visible connectivity.

## E3 - AVmatrix And Source Code Review Correction

Date: 2026-05-20

Status: recorded

Reason:

The plan must be grounded in codebase facts. This correction uses AVmatrix and source inspection before allowing the plan to drive implementation.

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

Conclusion:

The plan must not derive orphan status from raw incoming/outgoing edges. Structural edges such as `DEFINES`, `CONTAINS`, `HAS_METHOD`, `HAS_PROPERTY`, and `MEMBER_OF` materially change zero-incoming/zero-outgoing counts and can hide dead/unwired candidates if counted blindly.

## E4 - Initial Connectivity Baseline Commands

Date: 2026-05-20

Status: recorded

Raw all-relationship baseline command:

```powershell
$repos = @(
  @{Name='AVmatrix-GO'; Path='E:\AVmatrix-GO\.avmatrix\graph.json'},
  @{Name='Restaurant_manager'; Path='E:\Restaurant_manager\.avmatrix\graph.json'}
)
$codeLabels = @('Class','Function','Method','Interface','Struct','Trait','Impl','TypeAlias','Enum','Record','Delegate','Constructor','Route','Tool')
# For each repo: count incoming/outgoing by all relationships, then summarize raw zero-incoming,
# zero-outgoing, zero-both, code-label subsets, path-policy candidates, node labels, and relationship types.
```

Raw all-relationship result summary:

```text
AVmatrix-GO:
nodes=20967 relationships=52302
rawZeroIncoming=20 rawZeroOutgoing=15216 rawZeroBoth=8
codeNodeCount=4929 codeZeroIncoming=0 codeZeroOutgoing=200 codeZeroBoth=0
pathExpectedCandidateNodes=5743 pathExpectedCandidateZeroBoth=0

Restaurant_manager:
nodes=78358 relationships=130588
rawZeroIncoming=14 rawZeroOutgoing=56470 rawZeroBoth=2
codeNodeCount=10258 codeZeroIncoming=0 codeZeroOutgoing=1097 codeZeroBoth=0
pathExpectedCandidateNodes=12934 pathExpectedCandidateZeroBoth=0
```

Provisional non-structural policy command:

```powershell
$nonStructuralTypes = @('CALLS','INHERITS','METHOD_OVERRIDES','METHOD_IMPLEMENTS','IMPORTS','USES','DECORATES','IMPLEMENTS','EXTENDS','ACCESSES','STEP_IN_PROCESS','HANDLES_ROUTE','FETCHES','HANDLES_TOOL','ENTRY_POINT_OF','WRAPS','QUERIES')
$callGraphTypes = @('CALLS','HANDLES_ROUTE','HANDLES_TOOL','ENTRY_POINT_OF','STEP_IN_PROCESS')
# For each repo: count code/callable node incoming/outgoing using those relationship type sets only.
```

Provisional policy result summary:

```text
AVmatrix-GO:
nonStructural code nodes=4929 zeroIn=1616 zeroOut=1100 zeroBoth=133
callable flow nodes=4242 zeroIn=1587 zeroOut=1323 zeroBoth=264

Restaurant_manager:
nonStructural code nodes=10258 zeroIn=4190 zeroOut=3488 zeroBoth=909
callable flow nodes=8349 zeroIn=4172 zeroOut=3456 zeroBoth=1472
```

Interpretation:

- Raw all-relationship counts are not suitable as the orphan/dead-code denominator because `DEFINES` and container/ownership edges give code nodes incoming edges.
- The provisional policy is not final product behavior. It is recorded only to prove why Phase 1 must define a counted-edge policy before implementation.
