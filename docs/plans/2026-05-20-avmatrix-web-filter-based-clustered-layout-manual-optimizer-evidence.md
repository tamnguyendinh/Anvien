# AVmatrix Web Filter-Based Clustered Layout And Manual Optimizer Evidence Ledger

Date: 2026-05-20

Status: complete

Companion files:

- Plan: [2026-05-20-avmatrix-web-filter-based-clustered-layout-manual-optimizer-plan.md](2026-05-20-avmatrix-web-filter-based-clustered-layout-manual-optimizer-plan.md)
- Benchmark ledger: [2026-05-20-avmatrix-web-filter-based-clustered-layout-manual-optimizer-benchmark.md](2026-05-20-avmatrix-web-filter-based-clustered-layout-manual-optimizer-benchmark.md)

## Evidence Rules

Record evidence as each evidenced task is completed. Evidence should include commands, impacted files, test results, e2e artifacts, screenshots when useful, and concise observations needed to audit the plan later.

For doc-only commits, do not use AVmatrix.

Do not record inferred performance or behavior as fact. Every behavior claim must include a command, inspected file, test result, browser/e2e artifact, or runtime diagnostic.

## E0 - Plan Creation Evidence

Date: 2026-05-20

Status: recorded

Created file set:

- `docs/plans/2026-05-20-avmatrix-web-filter-based-clustered-layout-manual-optimizer-plan.md`
- `docs/plans/2026-05-20-avmatrix-web-filter-based-clustered-layout-manual-optimizer-evidence.md`
- `docs/plans/2026-05-20-avmatrix-web-filter-based-clustered-layout-manual-optimizer-benchmark.md`

Plan creation scope:

- Use the repository's established plan, evidence, and benchmark ledger format.
- Capture the user-approved design direction before implementation.
- Keep this as a Web UI layout plan, not a backend graph or analyzer plan.

Convention inspection commands:

```powershell
git status --short
Get-Content -Path docs\plans\2026-05-20-avmatrix-orphan-node-connectivity-lens-plan.md -TotalCount 80
Get-Content -Path docs\plans\2026-05-20-avmatrix-orphan-node-connectivity-lens-evidence.md -TotalCount 80
Get-Content -Path docs\plans\2026-05-20-avmatrix-orphan-node-connectivity-lens-benchmark.md -TotalCount 80
```

Observed planning standard:

- Plan file has date, status, companion files, rules, problem/scope, design decision, checklist, acceptance criteria, and closure definition.
- Evidence ledger records commands, inspected files, decisions, and validation artifacts.
- Benchmark ledger records product/runtime measurements and keeps unmeasured values explicit until they are recorded.

## E1 - Product Decision Evidence

Date: 2026-05-20

Status: recorded

Accepted direction from discussion:

- Node clustering should be the default layout foundation.
- Clustering should be based only on existing Node Type filters.
- Each node type/filter should map to one visual cluster.
- Do not compute node importance, hubs, degree, centrality, or "command" status for placement.
- The purpose is visual clarity, not ranking code by connectivity.
- Layout optimization is a separate manual cleanup feature.
- Layout optimization must not run automatically after graph load.

Rejected directions:

- Do not add a new taxonomy for layout clusters.
- Do not move "important" or highly connected nodes to the center.
- Do not use runtime optimizer as the primary way to make the graph readable.
- Do not use elapsed-time budget as the layout correctness mechanism.

## E2 - Current Codebase Facts From Initial Inspection

Date: 2026-05-20

Status: recorded

Initial code inspection commands:

```powershell
rg --files -g "*graph*" -g "*Graph*" -g "*.tsx" -g "*.ts" avmatrix-web internal cmd docs | Select-Object -First 120
rg -n "layout|optimiz|force|simulation|dagre|elk|cytoscape|sigma|cluster|community|graph" avmatrix-web/src avmatrix-web/e2e avmatrix-web/test internal -S
Get-Content -Path avmatrix-web\src\hooks\useSigma.ts -TotalCount 260
Get-Content -Path avmatrix-web\src\lib\graph-adapter.ts -TotalCount 260
Get-Content -Path avmatrix-web\src\hooks\app-state\graph.tsx -TotalCount 320
Get-Content -Path avmatrix-web\src\components\GraphCanvas.tsx -TotalCount 320
Get-Content -Path avmatrix-web\src\lib\constants.ts -TotalCount 180
```

Observed current implementation:

- `GraphCanvas` converts the loaded `KnowledgeGraph` into a Sigma graph through `knowledgeGraphToGraphology`.
- `GraphCanvas` builds `communityMemberships` from `MEMBER_OF` relationships before conversion.
- `knowledgeGraphToGraphology` already owns initial node position assignment.
- `knowledgeGraphToGraphology` currently mixes structural parent placement, community placement, and random jitter.
- `useSigma.setGraph` currently starts layout after assigning the Sigma graph.
- `useSigma` imports and uses `graphology-layout-forceatlas2/worker`.
- `useSigma` also runs `noverlap.assign` during layout stop/cleanup paths.
- Existing node filter order and known labels are exposed through Web constants such as `FILTERABLE_LABELS`; implementation should not use `DEFAULT_VISIBLE_LABELS` as the cluster order source.

Implementation implication:

- The likely smallest code change is Web-only:
  - change initial placement policy in `graph-adapter.ts`;
  - stop auto-starting layout in `useSigma.ts`;
  - keep optimizer as a manual GraphCanvas action.

## E3 - Required AVmatrix Impact Evidence

Date: 2026-05-20

Status: recorded

Commands:

```powershell
go run .\cmd\avmatrix analyze --force
go run .\cmd\avmatrix impact knowledgeGraphToGraphology --repo AVmatrix --direction upstream --depth 2 --include-tests
go run .\cmd\avmatrix impact useSigma --repo AVmatrix --direction upstream --depth 2 --include-tests
go run .\cmd\avmatrix impact GraphCanvas --repo AVmatrix --direction upstream --depth 2 --include-tests
```

Results:

- `analyze --force`: `files scanned=714 parsed=538 unsupported=176 failed=0`, graph `nodes=21700 relationships=54010`.
- `knowledgeGraphToGraphology`: MEDIUM risk, impacted count `8`; direct surfaces include `GraphCanvas.tsx`, `graph-adapter.edge-geometry.test.ts`, and `graph.test.ts`.
- `useSigma`: LOW risk, impacted count `4`; direct surface `GraphCanvas.tsx`.
- `GraphCanvas`: LOW risk, impacted count `3`; direct surfaces `App.tsx` and `GraphCanvas.selection-performance.test.tsx`.
- No HIGH or CRITICAL warning was reported.

## E4 - Implementation Evidence

Date: 2026-05-20

Status: recorded

Files changed:

- `avmatrix-web/src/lib/graph-adapter.ts`
- `avmatrix-web/src/hooks/useSigma.ts`
- `avmatrix-web/src/components/GraphCanvas.tsx`
- `avmatrix-web/src/lib/runtime-diagnostics.ts`
- `avmatrix-web/test/unit/graph-adapter.edge-geometry.test.ts`
- `avmatrix-web/test/unit/GraphCanvas.selection-performance.test.tsx`
- `avmatrix-web/test/unit/runtime-diagnostics.test.ts`
- `avmatrix-web/e2e/server-connect.spec.ts`

Implementation summary:

- Added `applyFilterBasedClusteredLayout` with cluster key `nodeType`, known order from `FILTERABLE_LABELS`, unknown labels sorted by label string, and deterministic in-cluster row-major local grids.
- Replaced random/hierarchy/community initial positioning in `knowledgeGraphToGraphology` with the filter-based clustered layout.
- Preserved existing colors, node sizes, graph-health metadata, community color metadata, and edge conversion.
- Removed automatic ForceAtlas2 worker startup from `useSigma.setGraph`.
- Reworked `startLayout` as a manual clustered layout cleanup action that reapplies the deterministic clustered layout and records a manual optimizer invocation.
- Changed the graph button label from `Run Layout Again` to `Optimize Layout`.
- Added `manualOptimizerInvocations`, `lastManualOptimizerInvokedAt`, and `lastManualOptimizerRunMs` runtime diagnostics.
- Updated e2e coverage so large graph load expects `layout.starts=0` and manual optimizer invocation increases only after clicking `Optimize Layout`.

## E5 - Validation Evidence

Date: 2026-05-20

Status: recorded

Commands and results:

```powershell
npm --prefix avmatrix-web run test -- test/unit/graph-adapter.edge-geometry.test.ts
```

- Passed: `1` file, `9` tests.

```powershell
npm --prefix avmatrix-web run test -- test/unit/runtime-diagnostics.test.ts test/unit/GraphCanvas.selection-performance.test.tsx
```

- Passed: `2` files, `7` tests.

```powershell
npm --prefix avmatrix-web run test -- test/unit/graph.test.ts test/unit/graph-edge-visibility-mode.test.ts test/unit/graph-edge-render-style.test.ts test/unit/selected-graph-context.test.ts
```

- Passed: `4` files, `17` tests.

```powershell
npm --prefix avmatrix-web run build
```

- Passed: TypeScript build and Vite production build completed.
- Vite reported existing chunk-size/dynamic-import warnings; no build failure.

```powershell
npm --prefix avmatrix-web run test:e2e -- server-connect.spec.ts -g "without automatic layout optimizer|manual layout optimizer" --workers=1 --timeout=120000
```

- Passed: `2` Playwright tests.
- `--timeout=120000` is a Playwright runner guard for validation execution only. It is not an accepted product or layout correctness mechanism.

Runtime diagnostic capture command:

```powershell
@'
const { chromium, expect } = require('@playwright/test');
// loads first indexed repo, captures diagnostics, clicks Optimize Layout, captures diagnostics again
'@ | node -
```

Observed runtime diagnostic capture:

- Repo: `Restaurant_manager`.
- Graph payload: `78358` nodes, `130588` relationships.
- Before manual optimizer: `layout.starts=0`, `layout.stops=0`, `manualOptimizerInvocations=0`.
- After clicking `Optimize Layout`: `layout.starts=0`, `layout.stops=0`, `manualOptimizerInvocations=1`.
- Manual optimizer run time recorded by diagnostics: `1183ms`.
- Heartbeat reconnects: `0`; reconnect banner shows: `0`.

Browser plugin note:

- Browser plugin instructions were read, but the required Node REPL browser-control tool was not exposed by tool discovery in this session. UI verification was completed through Playwright e2e and a direct Playwright diagnostic capture instead.

## E6 - Detect Changes Evidence

Date: 2026-05-20

Status: recorded

Command:

```powershell
mcp__avmatrix__.detect_changes({ repo: "AVmatrix", scope: "all" })
```

Result summary:

- Changed files: `11`
- Changed symbols: `66`
- Affected process count: `4`
- Risk level: `medium`
- Affected processes reported around removed/replaced graph-adapter hierarchy helpers, including `AddHierarchyLink -> CompareKnownOrder`, `AddHierarchyLink -> GetDisplayGraphRelationships`, `AddHierarchyLink -> GetGroupedHeritageCompatibilityKeys`, and `AddHierarchyLink -> GetNodeLabelCounts`.
- No HIGH or CRITICAL warning was reported.

Interpretation:

- Medium risk is expected because the graph adapter's previous hierarchy-placement helper was removed and replaced by filter-based clustered placement.
- Scope matches the intended Web layout, runtime diagnostics, e2e, unit test, and plan ledger changes.
