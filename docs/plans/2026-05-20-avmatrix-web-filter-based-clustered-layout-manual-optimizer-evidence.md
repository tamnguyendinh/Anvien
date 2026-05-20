# AVmatrix Web Filter-Based Clustered Layout And Manual Optimizer Evidence Ledger

Date: 2026-05-20

Status: recorded - visual island and documentation center validation passing

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
- Product/runtime timeout is not an accepted operating mechanism. Timeout may be used only by tests or e2e runners as a guard.

Rejected directions:

- Do not add a new taxonomy for layout clusters.
- Do not move "important" or highly connected nodes to the center.
- Do not use runtime optimizer as the primary way to make the graph readable.
- Do not use elapsed-time budget as the layout correctness mechanism.
- Do not use timeout, timer reset, or delayed UI reset in product/runtime code as a graph-load, layout, reconnect, reset, or lag-control mechanism.

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

## E7 - Corrective Review Evidence

Date: 2026-05-20

Status: recorded

Reason for reopening:

- User reported the Web graph still does not read as true clusters.
- User clarified that "cluster" means one Node Type filter equals one clear visual region and one node type color.
- User reported layout optimization appearing after render without a manual click.
- User rejected timeout-based behavior as a product/runtime mechanism. This rejection is not conditional on repository size.

Inspection commands:

```powershell
rg -n "startLayout|stopLayout|setIsLayoutRunning|isLayoutRunning|Layout optimizing|Optimize Layout|Run Layout|recordLayoutStart|recordLayoutStop|ForceAtlas|noverlap|animatedReset|setGraph\(" avmatrix-web\src avmatrix-web\test avmatrix-web\e2e
rg -n "fetchGraph\(|loadGraph|setCurrentView|setGraph\(|error|timeout|reset|onboarding|analyze" avmatrix-web\src\App.tsx avmatrix-web\src\hooks\useAppState.local-runtime.tsx avmatrix-web\src\services\backend-client.ts
Get-Content -Path avmatrix-web\src\lib\graph-adapter.ts -TotalCount 380
Get-Content -Path avmatrix-web\src\hooks\useSigma.ts -TotalCount 620
Get-Content -Path avmatrix-web\src\services\backend-client.ts -TotalCount 540
```

Observed facts:

- `GraphCanvas` calls `startLayout` only from the `Optimize Layout` button.
- `useSigma.startLayout` currently reapplies `applyFilterBasedClusteredLayout`, records manual optimizer diagnostics, and toggles `isLayoutRunning`.
- `useSigma.setGraph` still calls `sigma.getCamera().animatedReset({ duration: 500 })` after graph replacement. This is camera animation, not layout optimization, but it can be confused with post-render runtime movement and must be reviewed against the user's "no auto optimizing after render" requirement.
- `backend-client.fetchGraph` still passes a hard `120_000ms` timeout to `fetchWithTimeout` for graph fetches. This is not layout logic, but it is product/runtime timeout behavior and must be handled as a separate issue from clustering and optimizer behavior.
- `knowledgeGraphToGraphology` still accepted `communityMemberships` and used community color for `COMMUNITY_COLORED_NODE_LABELS` before the corrective review. That mixed multiple colors inside the same node type/filter cluster and violated the user's one-color-per-cluster expectation.

Corrective impact evidence:

```powershell
go run .\cmd\avmatrix analyze --force
go run .\cmd\avmatrix impact knowledgeGraphToGraphology --repo AVmatrix --direction upstream --depth 2 --include-tests
go run .\cmd\avmatrix impact applyFilterBasedClusteredLayout --repo AVmatrix --direction upstream --depth 2 --include-tests
go run .\cmd\avmatrix impact useSigma --repo AVmatrix --direction upstream --depth 2 --include-tests
```

Results:

- `analyze --force`: `files scanned=714 parsed=538 unsupported=176 failed=0`, graph `nodes=21685 relationships=54025`.
- `knowledgeGraphToGraphology`: MEDIUM risk, impacted count `8`.
- `applyFilterBasedClusteredLayout`: HIGH risk, impacted count `5`, affected modules `Cluster` and `Hooks`, affected process count `4`.
- `useSigma`: LOW risk, impacted count `4`.

Warning:

- `applyFilterBasedClusteredLayout` is HIGH risk because it is the primary Web graph layout policy and directly affects `useSigma` and graph render behavior. Corrective implementation must stay narrowly scoped and must be validated visually and with tests.

Separation note:

- Clustering, optimizer invocation, and timeout policy are three separate work tracks.
- Clustering must be solved by existing node type/filter/color placement.
- Optimizer invocation must be solved by call-path audit and manual-only UI trigger.
- Timeout policy must be solved by removing product/runtime timeout/reset mechanisms, not by tuning durations or making repo-size exceptions.

Unreviewed code diff created before this evidence update:

- At the time of E7, `avmatrix-web/src/lib/graph-adapter.ts` had a local change that removed community color as the primary render color and used `getNodeColor(node.label)` instead.
- This diff is not considered complete implementation until the plan is agreed, tests are updated, and browser/e2e behavior is verified.

## E8 - Corrective Implementation Evidence

Date: 2026-05-20

Status: recorded

Files changed in the corrective implementation:

- `avmatrix-web/src/lib/graph-adapter.ts`
- `avmatrix-web/src/lib/runtime-diagnostics.ts`
- `avmatrix-web/src/services/backend-client.ts`
- `avmatrix-web/src/App.tsx`
- `avmatrix-web/src/hooks/useAppState.local-runtime.tsx`
- `avmatrix-web/src/hooks/useBackend.ts`
- `avmatrix-web/src/components/AnalyzeProgress.tsx`
- `avmatrix-web/src/components/CodeReferencesPanel.tsx`
- `avmatrix-web/src/components/DropZone.tsx`
- `avmatrix-web/src/components/EncouragementLine.tsx`
- `avmatrix-web/src/components/MarkdownRenderer.tsx`
- `avmatrix-web/src/components/MermaidDiagram.tsx`
- `avmatrix-web/src/components/OnboardingGuide.tsx`
- `avmatrix-web/src/components/ProcessesPanel.tsx`
- `avmatrix-web/src/components/RepoAnalyzer.tsx`
- `avmatrix-web/src/config/ui-constants.ts`
- `avmatrix-web/test/unit/graph-adapter.edge-geometry.test.ts`
- `avmatrix-web/test/unit/heartbeat.test.ts`
- `avmatrix-web/test/unit/runtime-diagnostics.test.ts`
- `avmatrix-web/test/unit/server-connection.test.ts`
- `avmatrix-web/e2e/onboarding.spec.ts`
- `avmatrix-web/e2e/repo-switching.spec.ts`
- `avmatrix-web/e2e/server-connect.spec.ts`

Corrective implementation summary:

- `knowledgeGraphToGraphology` now uses `getNodeColor(node.label)` as the primary render color for every node, while retaining `community` and `communityColor` as metadata only.
- `applyFilterBasedClusteredLayout` remains the layout source for initial placement and manual optimization. It groups by `nodeType`, orders clusters by the existing filter label order, and places each type into separated deterministic regions.
- Runtime diagnostics no longer expose elapsed-time layout budget fields or layout stop reasons from the old optimizer model.
- Product/runtime code in `avmatrix-web/src` no longer contains timeout/timer/reset mechanisms found by the final inspection command.
- Process list View/lightbulb behavior no longer calls `/api/query` for process steps. It derives `STEP_IN_PROCESS` steps and `CALLS` edges from the graph already loaded in the client, avoiding backend query stalls on large repos without adding timeout logic.
- Heartbeat reconnect behavior relies on browser `EventSource` reconnect semantics instead of application retry timers.

Final inspection commands:

```powershell
rg -n "setTimeout|clearTimeout|setInterval|clearInterval|timeout|Timeout|TIMEOUT|durationBudget|duration-elapsed|noverlap|lastReason" avmatrix-web\src
rg -n "runQuery|focusLoadingProcess|focusRequestIdRef|isSafeId|/api/query|STEP_IN_PROCESS|CALLS" avmatrix-web\src\components\ProcessesPanel.tsx
```

Results:

- First command returned no matches in `avmatrix-web/src`.
- Second command returned only the expected local graph relationship checks for `STEP_IN_PROCESS` and `CALLS`; `ProcessesPanel` no longer calls `runQuery`.

## E9 - Corrective Validation Evidence

Date: 2026-05-20

Status: recorded

Commands and results:

```powershell
npm --prefix avmatrix-web run build
```

- Passed: TypeScript build and Vite production build completed.
- Vite reported existing chunk-size/dynamic-import warnings; no build failure.

```powershell
npm --prefix avmatrix-web run test
```

- Passed: `43` files, `336` tests.

```powershell
npm --prefix avmatrix-web run test:e2e -- server-connect.spec.ts -g "shows process list and View button works|lightbulb highlights nodes in graph" --workers=1
```

- Passed: `2` Playwright tests.
- This was rerun after moving process modal/highlight data extraction to the loaded client graph.

```powershell
npm --prefix avmatrix-web run test:e2e -- --workers=1
```

- Passed: `42` Playwright tests.
- Duration: `20.7m`.
- Key final e2e observations:
  - `keeps connection stable after large graph load without automatic layout optimizer` passed.
  - `invokes manual layout optimizer only after user action` passed.
  - `shows process list and View button works` passed after the corrective change.
  - `lightbulb highlights nodes in graph` passed.

Scope review:

- No backend graph schema changes were made.
- No new node label taxonomy was added.
- No hub, degree, centrality, or importance calculation was added.
- Product/runtime timeout and delayed-reset logic was removed instead of tuned.

## E10 - Final Detect Changes Evidence

Date: 2026-05-20

Status: recorded

Commands:

```powershell
avmatrix analyze --force
mcp__avmatrix__.detect_changes({ repo: "AVmatrix", scope: "all" })
```

Results:

- `analyze --force`: `files scanned=714 parsed=538 unsupported=176 failed=0`, graph `nodes=20760 relationships=49625`.
- `detect_changes`: changed files `23`, changed symbols `125`, affected process count `31`, risk level `critical`.

Affected process themes:

- repo/connect graph fetch flows through `fetchFromBackend`, `fetchGraph`, `fetchRepoInfo`, and `BackendError`;
- auto-connect flow through `useBackend` and `probeBackend`;
- process panel local graph extraction helpers introduced in `ProcessesPanel`;
- graph adapter node render color and layout conversion tests;
- plan/evidence/benchmark documentation sections.

Interpretation:

- The `critical` risk level is expected for this corrective implementation because the product timeout ban intentionally removed timeout/retry/timer behavior from shared backend client and connect flows.
- The affected scope matches the implementation tracks in the plan: filter/color clustering, manual-only optimizer, and product/runtime timeout removal.
- The risk is mitigated by final validation: production build passed, product Go build passed, launcher Go module builds passed, full unit passed (`43` files, `336` tests), and full e2e passed (`42` tests in `20.7m`).

## E11 - Full Product Build Evidence

Date: 2026-05-20

Status: recorded

Commands and results:

```powershell
npm --prefix avmatrix-web run build
```

- Passed: TypeScript build and Vite production build completed.
- Vite reported existing chunk-size/dynamic-import warnings; no build failure.

```powershell
go build ./cmd/... ./internal/...
```

- Passed: root product CLI and internal Go packages built successfully.

```powershell
go build ./...
```

- Failed as an acceptance command because it includes analysis fixtures under `avmatrix/test/fixtures`.
- The failing fixture errors were expected fixture-shape errors, including unresolved sample imports, mixed fixture packages, and C source fixture input.
- This command is not a valid full-product build command for this repository because the fixture directories are source samples for AVmatrix analysis, not buildable Go packages.

```powershell
go build ./...
```

- Working directory: `avmatrix-launcher/server-wrapper`.
- Passed.
- The produced local `.exe` artifact was removed after validation.

```powershell
go build ./...
```

- Working directory: `avmatrix-launcher/src`.
- Passed.
- The produced local `.exe` artifact was removed after validation.

Final status:

- Web production build: passed.
- Root product Go build: passed for `cmd` and `internal`.
- Launcher Go builds: passed.
- Worktree returned clean after removing build artifacts.

## E12 - Visual Island Distribution Reopen Evidence

Date: 2026-05-20

Status: recorded for plan update and implementation

User-provided artifacts:

- Current failing output: `reports/problem/screenshot_1779285599.png`.
- Target placement reference: `reports/problem/aaaa.jpg`.

Observed problem:

- The current clustered layout still reads as compressed rails or packed blocks, not as readable node type/color islands.
- Nodes are too close together for the visible graph scale, and medium/large clusters do not have enough two-dimensional spread.
- Previous validation accepted deterministic non-overlapping cluster bounds, but that was insufficient. It did not detect rail-like cluster shape, excessive density, poor whitespace, or screenshot-level readability failure.

Clarification recorded:

- The intended visual model is colored archipelagos on one large circular graph field.
- The sample image is only a reference for how to distribute node clusters visually.
- The implementation must not copy the sample by reducing, hiding, filtering, pruning, thinning, or reweighting graph edges.
- Relationship data, edge count, cross-cluster links, and existing edge visibility behavior must be preserved unless the user creates a separate edge-display plan.

Evidence recorded in E14:

- Code diff review proving only node placement/cluster geometry changed, not graph relationship data.
- Browser screenshot after graph ready on the representative large repo showing separated two-dimensional color islands.
- Per-cluster diagnostics for node count, color, width, height, aspect ratio, density, and inter-cluster gutter.
- Edge preservation evidence showing relationship count and existing edge visibility behavior remain intact.
- Re-run focused layout tests, full Web build, full Web unit tests, full Web e2e tests, root product Go build, and launcher Go builds.

## E13 - Documentation Center Reopen Evidence

Date: 2026-05-20

Status: recorded for plan update and implementation

User clarification:

- Documentation-system files must be separated into one dedicated node/filter type.
- The Documentation filter must have its own color.
- The Documentation work is its own implementation phase, not an implicit part of the generic outer island layout.
- The Documentation island must be placed at the center of the large circular graph field.

Plan impact:

- The plan now treats `Documentation` as one explicit display-filter exception to the existing Node Type filter source of truth.
- Raw graph labels, relationship data, edge counts, edge visibility rules, and analyzer behavior remain preserved unless implementation evidence proves a narrow payload change is required.
- Outer filter/color islands remain arranged as colored archipelagos around the Documentation center.

Evidence recorded in E14:

- Code diff showing the Documentation filter/color/classification path.
- Unit or browser diagnostic proving documentation nodes are grouped into one `Documentation` display filter.
- Unit or browser diagnostic proving the Documentation island is centered before manual optimizer action.
- Screenshot evidence showing the Documentation center island and surrounding outer color islands.
- Edge preservation evidence showing relationship count and existing edge visibility behavior remain intact.
- Re-run focused layout tests, full Web build, full Web unit tests, full Web e2e tests, root product Go build, and launcher Go builds.

## E14 - Documentation Center Implementation And Validation Evidence

Date: 2026-05-20

Status: recorded

AVmatrix impact and risk evidence:

- `avmatrix analyze --force` refreshed the index before implementation.
- `avmatrix impact applyFilterBasedClusteredLayout --repo "E:\AVmatrix-GO" --direction upstream --depth 2 --include-tests` returned HIGH risk; blast radius included `useSigma` and `GraphCanvas`.
- `avmatrix impact knowledgeGraphToGraphology --repo "E:\AVmatrix-GO" --direction upstream --depth 2 --include-tests` returned MEDIUM risk; blast radius included `GraphCanvas` and unit tests.
- `avmatrix impact getNodeColor`, `getNodeLabelCounts`, `getFilterableNodeLabelsForGraph`, and `getNodeTypeIcon` returned CRITICAL risk because they affect filter UI and graph display inventory.

Implementation files:

- `avmatrix-web/src/lib/constants.ts`
- `avmatrix-web/src/lib/graph-adapter.ts`
- `avmatrix-web/src/components/FileTreePanel.tsx`
- `avmatrix-web/test/unit/constants.test.ts`
- `avmatrix-web/test/unit/filter-panel.test.ts`
- `avmatrix-web/test/unit/graph-adapter.edge-geometry.test.ts`

Implementation summary:

- Added display-only `Documentation` filter label with dedicated color `#84cc16`.
- Added documentation classification from existing facts: file path, directory segment, extension, file name, and node metadata.
- Added `getNodeDisplayLabel` so documentation nodes become `Documentation` for Web filtering and layout while preserving raw graph label in `rawNodeType`.
- Updated clustered layout so `Documentation` is placed at the graph layout origin and all other display-filter islands are distributed around it.
- Re-centered per-cluster organic offsets so the intended island center is the actual island bounding center.
- Kept relationship conversion unchanged; no edge pruning, hiding, thinning, or backend schema change was introduced.

Validation commands:

- `npm --prefix avmatrix-web run test -- test/unit/constants.test.ts test/unit/filter-panel.test.ts test/unit/graph-adapter.edge-geometry.test.ts`
  - Passed: 3 files, 53 tests.
- `npm --prefix avmatrix-web run build`
  - Passed; Vite reported existing chunk-size/dynamic-import warnings only.
- `npm --prefix avmatrix-web run test`
  - Passed: 43 files, 344 tests.
- `npm --prefix avmatrix-web run test:e2e -- --workers=1`
  - Passed: 42 tests in 25.6 minutes.
- `go build ./cmd/... ./internal/...`
  - Passed.
- `go build ./...` in `avmatrix-launcher/src`
  - Passed.
- `go build ./...` in `avmatrix-launcher/server-wrapper`
  - Passed.

Pre-commit AVmatrix change detection:

- `avmatrix analyze --force`
  - Passed before the final change-scope check.
- `avmatrix detect-changes --repo "E:\AVmatrix-GO" --scope all`
  - Passed with expected Web-only/doc scope: changed files `9`, changed symbols `156`, affected process count `6`, risk level `high`.
  - Affected processes: `PlaceCluster -> CompareKnownOrder`, `PlaceCluster -> GetDisplayGraphRelationships`, `PlaceCluster -> GetFileExtension`, `PlaceCluster -> GetFileStem`, `PlaceCluster -> GetCommunityColor`, and `PlaceCluster -> GetEdgeInfo`.
  - Scope review: expected Web constants/filter classification, graph adapter layout, FileTreePanel icon, unit tests, and plan evidence files. No backend graph schema, analyzer, relationship extraction, or edge data contract change was included.

Browser evidence:

- Screenshot: `reports/problem/screenshot_20260520_documentation_center_after.png`.
- Diagnostics: `reports/problem/screenshot_20260520_documentation_center_after-diagnostics.json`.
- Project: `AVmatrix`.
- UI filter evidence: `Documentation (1335)`.
- Runtime diagnostics: `nodeCount=21761`, `lastRelationshipCount=54298`, `layout.starts=0`, `layout.stops=0`, `manualOptimizerInvocations=0`, `heartbeat.reconnects=0`, `reconnectBanner.shows=0`.
