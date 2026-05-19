# Web UI Left Dashboard Graph Display Completeness Evidence Ledger

Date: 2026-05-19

Status: active

Companion files:

- Plan: [2026-05-19-web-ui-left-dashboard-graph-display-completeness-plan.md](2026-05-19-web-ui-left-dashboard-graph-display-completeness-plan.md)
- Benchmark ledger: [2026-05-19-web-ui-left-dashboard-graph-display-completeness-benchmark.md](2026-05-19-web-ui-left-dashboard-graph-display-completeness-benchmark.md)

## Evidence Rules

Record evidence as each evidenced task is completed. Evidence should include commands, impacted files, test results, e2e artifacts, and concise observations needed to audit the plan later.

For doc-only commits, do not use AVmatrix.

## E0 - Plan Creation Evidence

Date: 2026-05-19

Created file set:

- `docs/plans/2026-05-19-web-ui-left-dashboard-graph-display-completeness-plan.md`
- `docs/plans/2026-05-19-web-ui-left-dashboard-graph-display-completeness-benchmark.md`
- `docs/plans/2026-05-19-web-ui-left-dashboard-graph-display-completeness-evidence.md`

Reason:

The Web UI left dashboard currently does not show or control all graph components that users need to understand the loaded graph. This plan covers Node Types, Edge Types, Color Legend, and canvas relationship rendering/completeness.

Plan review update:

- tightened the requirement so every node label and relationship type present in the loaded graph must appear as an individual dashboard option;
- clarified that visual grouping is allowed only for sectioning, not for hiding real graph types;
- added representative fixture coverage because the current AVmatrix-GO graph does not contain every known graph label or relationship type;
- added fallback requirements for unknown/future node labels and relationship types;
- documented that dashboard edge controls target graph payload relationship types, not storage-only LadybugDB constants unless those appear in graph payloads;
- added the oversized purple structural-node visual-scale issue from `reports/problem/screenshot_1779178877.png`;
- added the post-load launcher lifecycle timeout / `Server connection lost - reconnecting...` runtime stability issue.

## E1 - Initial Source Inspection

Date: 2026-05-19

Observed files:

- `avmatrix-web/src/lib/constants.ts`
- `avmatrix-web/src/components/FileTreePanel.tsx`
- `avmatrix-web/src/lib/graph-adapter.ts`
- `avmatrix-web/src/hooks/useSigma.ts`
- `avmatrix-web/src/generated/avmatrix-contracts.ts`
- `internal/graph/types.go`
- `internal/httpapi/graph.go`
- `internal/httpapi/heartbeat.go`
- `avmatrix-web/src/services/backend-client.ts`
- `avmatrix-web/src/App.tsx`
- `avmatrix-launcher/src/main.go`

Observed implementation facts:

- `FILTERABLE_LABELS` is a hard-coded list of 11 labels in `avmatrix-web/src/lib/constants.ts`.
- `EdgeType` is a hard-coded union of 6 relationship types in `avmatrix-web/src/lib/constants.ts`.
- `Color Legend` in `FileTreePanel.tsx` is a hard-coded list of 10 labels.
- `graph-adapter.ts` only treats `CONTAINS`, `DEFINES`, and `IMPORTS` as hierarchy relationships for positioning.
- `graph-adapter.ts` checks whether an edge already exists between a source and target before adding a relationship, which can lose information when multiple relationship types share the same endpoints.
- The HTTP graph API returns nodes and relationships from the graph payload, including relationship audit metadata.
- The Web client NDJSON graph loader pushes every relationship record into the in-memory graph before the canvas adapter runs.
- The current real graph is a necessary baseline but not sufficient as the only UI completeness fixture because it contains `16 / 36` known node labels and `11 / 22` graph payload relationship types.

Checklist updates from this review:

- `P1-A` completed: source lists, size maps, edge style maps, layout groups, and heartbeat paths were inspected.
- `P1-B` completed: `.avmatrix/graph.json` inventory was recorded in the benchmark ledger.
- `P1-E` completed: missing labels/types and parallel relationship risk were recorded.
- `P1-H` completed: oversized purple-node code path was reviewed and conclusions/hypotheses were recorded.
- `P1-I` completed: post-load reconnect code path was reviewed and conclusions/hypotheses were recorded.

## E1B - User-Reported Visual and Runtime Problems

Date: 2026-05-19

Screenshot artifact:

```text
reports/problem/screenshot_1779178877.png
```

File check:

```text
FullName: E:\AVmatrix-GO\reports\problem\screenshot_1779178877.png
Length: 1,144,772 bytes
LastWriteTime: 2026-05-19 15:21:30 local time
```

Problem 1:

- Many purple circular structural/folder-like nodes are much too large compared with surrounding nodes.
- Different node sizes are expected, but this observed scale is disproportionate and harms graph readability.
- The plan must audit node-size constants, graph-size scaling, community/structural styling, zoom behavior, and final rendered size ratios.
- Code review found no final rendered-size cap after `useSigma` reducer multipliers.
- Root cause: `NODE_SIZES` intentionally gives structural/purple families large base sizes. In the current graph size band, `getScaledNodeSize` leaves `Folder=5`, `Project=10`, `Package=8`, and `Module=6.5`, while many code/member nodes are `1.5-2`. Because node size is rendered as circle radius, `Folder=5` has about `11x` the area of a `Property=1.5` node. Selected/highlight/glow/ripple states can multiply those values further by `1.6x-2.5x`.

Problem 2:

- After the Web UI finishes loading/rendering the graph and runs for a short while, the UI can show `Server connection lost - reconnecting...`.
- This is a real stability problem, not just a transient banner.
- Confirmed root cause chain from log/code: `avmatrix-launcher/src/main.go` injects a script that pings every `5s`, while launcher expiry is `15s`. The launcher log shows `web ui session closed` at `2026-05-19 15:21:45`, 15 seconds after the screenshot timestamp window, while the graph was visible. That line is emitted only when the lifecycle monitor fires. Because this launcher owned backend pid `13752`, lifecycle expiry can stop the backend, which makes the Web heartbeat SSE fail and show the reconnect banner.
- The remaining subcause to measure is why launcher heartbeat was not delivered in time. The leading candidate is browser main-thread starvation from synchronous graph conversion/layout/noverlap during heavy graph load.

## E1C - Root-Cause Code Review Details

Date: 2026-05-19

Oversized node code paths reviewed:

- `avmatrix-web/src/lib/constants.ts`
  - `NODE_SIZES` sets `Project=20`, `Package=16`, `Module=13`, `Folder=10`, and leaf/member nodes often `1.5-2`.
- `avmatrix-web/src/lib/graph-adapter.ts`
  - `getScaledNodeSize` uses `baseSize * 0.5` when node count is greater than `20,000`.
  - current graph size is `20,354`, so Project scales to `10`, Package to `8`, Module to `6.5`, while member/leaf nodes floor at `1.5`.
- `avmatrix-web/src/hooks/useSigma.ts`
  - selected node multiplier is `1.8x`;
  - query highlight multiplier is `1.6x`;
  - glow animation multiplier reaches `2.0x`;
  - ripple animation multiplier reaches `2.5x`;
  - there is no cap after these multipliers.

Post-load reconnect code paths reviewed:

- `avmatrix-web/src/App.tsx`
  - heartbeat is created only when `viewMode === 'exploring'`;
  - the banner text is shown when `serverDisconnected` is true.
- `avmatrix-web/src/services/backend-client.ts`
  - `connectHeartbeat` opens `EventSource('/api/heartbeat')`;
  - on any EventSource error it immediately calls `onReconnecting`, which shows the banner;
  - it retries indefinitely.
- `internal/httpapi/heartbeat.go`
  - `/api/heartbeat` writes an SSE comment immediately, then every `15s`.
- `avmatrix-launcher/src/main.go`
  - lifecycle script sends launcher heartbeat every `5s`;
  - launcher UI timeout is `15s`;
  - when lifecycle expires, `waitForExit` treats the UI session as closed and removes state.

Root-cause conclusions and remaining hypotheses:

- visual scale confirmed root: structural/purple node base sizes and radius-based rendering create excessive area ratios; highlight/animation multipliers can make it worse;
- reconnect confirmed root chain: launcher lifecycle timeout closes the UI session/backend while graph is visible;
- reconnect remaining hypothesis: heavy graph conversion/layout/noverlap blocks or delays launcher heartbeat delivery long enough to trigger the `15s` timeout.

Launcher log excerpt:

```text
2026/05/19 15:20:03.638886 start root=E:\AVmatrix-GO
2026/05/19 15:20:03.646305 backend pid=13752
2026/05/19 15:21:45.648312 web ui session closed
```

## E2 - Initial Graph Inventory Evidence

Date: 2026-05-19

Command:

```powershell
.\avmatrix-launcher\server-bundle\avmatrix.exe analyze E:\AVmatrix-GO --force --skip-agents-md --no-stats
```

Result:

```text
analyzed E:\AVmatrix-GO
files: scanned=682 parsed=527 unsupported=155 failed=0
graph: nodes=20354 relationships=50980 path=E:\AVmatrix-GO\.avmatrix\graph.json
```

Node label inventory command:

```powershell
$g = Get-Content -Raw -LiteralPath '.avmatrix\graph.json' | ConvertFrom-Json
$g.nodes | Group-Object label | Sort-Object Name
```

Result summary:

```text
Class 4
Community 934
Const 321
Constructor 5
File 682
Folder 112
Function 3339
Interface 98
Method 804
Package 413
Process 644
Property 3096
Section 889
Struct 500
TypeAlias 69
Variable 8444
```

Relationship type inventory command:

```powershell
$g.relationships | Group-Object type | Sort-Object Name
```

Result summary:

```text
ACCESSES 5024
CALLS 8396
CONTAINS 1658
DEFINES 17093
ENTRY_POINT_OF 644
HAS_METHOD 336
HAS_PROPERTY 2769
IMPORTS 3713
MEMBER_OF 3826
STEP_IN_PROCESS 2373
USES 5148
```

Parallel relationship risk command:

```powershell
$pairs = $g.relationships | Group-Object { "$($_.sourceId)->$($_.targetId)" }
$parallel = $pairs | Where-Object { ($_.Group | Group-Object type).Count -gt 1 }
$parallel.Count
```

Result:

```text
1421
```

## E3 - Implementation Evidence

Status: active

Record each implementation slice here:

- files changed;
- AVmatrix impact/check results where applicable;
- build/test/e2e commands;
- important screenshots or textual e2e assertions;
- benchmark ledger entries updated.

### E3A - Launcher Lifecycle Heartbeat Budget Removal

Date: 2026-05-19

Scope:

- `avmatrix-launcher/src/main.go`
- `avmatrix-launcher/src/main_test.go`
- `docs/plans/2026-05-19-web-ui-left-dashboard-graph-display-completeness-plan.md`
- `docs/plans/2026-05-19-web-ui-left-dashboard-graph-display-completeness-benchmark.md`
- `docs/plans/2026-05-19-web-ui-left-dashboard-graph-display-completeness-evidence.md`

AVmatrix context/impact used before implementation:

```powershell
.\avmatrix-launcher\server-bundle\avmatrix.exe context newWebLifecycleMonitor --repo AVmatrix
.\avmatrix-launcher\server-bundle\avmatrix.exe impact newWebLifecycleMonitor --repo AVmatrix --direction upstream --depth 2 --include-tests
.\avmatrix-launcher\server-bundle\avmatrix.exe context connectHeartbeat --repo AVmatrix
.\avmatrix-launcher\server-bundle\avmatrix.exe impact connectHeartbeat --repo AVmatrix --direction upstream --depth 2 --include-tests
```

Result summary:

- `newWebLifecycleMonitor` direct callers: `startRuntime` and launcher lifecycle tests.
- `connectHeartbeat` direct use remains `AppContent` and heartbeat unit tests; Web heartbeat displays the banner after backend/session loss but is not the root shutdown decision.
- Implementation risk was centered in launcher lifecycle and covered by launcher tests plus Web graph-load e2e.

Implementation:

- Removed launcher heartbeat-missing auto-shutdown. The injected heartbeat still pings every `5s`, but stale heartbeat age is diagnostic only.
- Kept explicit `pagehide` close handling with short close grace so real browser close/reload can still stop the launcher-owned runtime.
- Added `webLifecycleSnapshot` with heartbeat age, close age, close-grace, close reason, and backend ownership logging.
- Added launcher tests proving stale heartbeat and large heartbeat gaps do not expire lifecycle, while explicit close still expires after grace.

Build before tests:

```powershell
go build ./...
```

Result:

```text
failed in intentionally non-standalone fixture packages under avmatrix/test/fixtures:
- package models is not in std
- package animal is not in std; mixed packages animal/main
- C source files not allowed when not using cgo or SWIG
```

Repo-specific build commands used after confirming the known fixture boundary:

```powershell
npm --prefix avmatrix-web run build
go build -trimpath -o .tmp\avmatrix.exe .\cmd\avmatrix
go build .
```

Results:

```text
npm --prefix avmatrix-web run build: passed, built in 21.74s
go build -trimpath -o .tmp\avmatrix.exe .\cmd\avmatrix: passed
go build . in avmatrix-launcher/src: passed
```

Focused tests:

```powershell
go test .
```

Run from:

```text
E:\AVmatrix-GO\avmatrix-launcher\src
```

Result:

```text
ok avmatrix-launcher 3.871s
```

Broader Go validation:

```powershell
go test ./cmd/... ./internal/... -count=1
```

Result:

```text
all reported packages passed except internal/session TestControllerCancelsPreviousSessionForSameRepo.
The failing assertion was runs = second,first.
```

Follow-up classification:

```powershell
go test ./internal/session -run TestControllerCancelsPreviousSessionForSameRepo -count=1 -v
go test ./internal/session -run TestControllerCancelsPreviousSessionForSameRepo -count=3 -v
```

Result:

```text
-count=1 passed.
-count=3 passed twice and failed once with runs = first.
```

Interpretation: the `internal/session` failure is pre-existing/flaky ordering behavior outside this launcher slice; it is not caused by the launcher lifecycle change.

Web unit tests:

```powershell
npm --prefix avmatrix-web run test
```

Result:

```text
39 test files passed
296 tests passed
```

E2E:

```powershell
npm --prefix avmatrix-web run test:e2e -- server-connect.spec.ts -g "selects a repo from landing and loads graph" --workers=1 --timeout=120000
```

Result:

```text
1 passed
test duration: 25.8s
total duration: 31.4s
```

Known validation note:

- Full `server-connect.spec.ts` run was attempted with a `180s` command timeout and did not complete before the shell timeout. The focused graph-load e2e above completed successfully.
- An initial `npm --prefix avmatrix-web run test -- --runInBand` attempt failed because Vitest does not support the Jest `--runInBand` option; the normal Vitest command passed.

Benchmark ledger updated:

- `B4A - Launcher Lifecycle Budget Removal`

### E3B - Dashboard Node/Edge/Legend Completeness

Date: 2026-05-19

Scope:

- `avmatrix-web/src/lib/constants.ts`
- `avmatrix-web/src/components/FileTreePanel.tsx`
- `avmatrix-web/src/hooks/app-state/graph.tsx`
- `avmatrix-web/src/hooks/useAppState.local-runtime.tsx`
- `avmatrix-web/src/lib/graph-adapter.ts`
- `avmatrix-web/test/unit/constants.test.ts`
- `avmatrix-web/test/unit/filter-panel.test.ts`
- `avmatrix-web/test/unit/FileTreePanel.dashboard-completeness.test.tsx`
- `docs/plans/2026-05-19-web-ui-left-dashboard-graph-display-completeness-plan.md`
- `docs/plans/2026-05-19-web-ui-left-dashboard-graph-display-completeness-benchmark.md`
- `docs/plans/2026-05-19-web-ui-left-dashboard-graph-display-completeness-evidence.md`

AVmatrix refresh:

```powershell
.\avmatrix-launcher\server-bundle\avmatrix.exe analyze E:\AVmatrix-GO --force --skip-agents-md --no-stats
```

Result:

```text
analyzed E:\AVmatrix-GO
files: scanned=685 parsed=527 unsupported=158 failed=0
graph: nodes=20421 relationships=51111 path=E:\AVmatrix-GO\.avmatrix\graph.json
```

Implementation:

- Node Types controls are now derived from labels present in the loaded graph and show counts.
- Edge Types controls are now derived from relationship types present in the loaded graph and show counts.
- `EdgeType` visibility state now supports generated graph relationship types and unknown/future strings.
- `EDGE_INFO` covers all `22` generated graph payload relationship types and has fallback display info for future types.
- `FILTERABLE_LABELS` now follows the generated `NODE_LABELS` contract order.
- `Community` and `Process` now have non-zero sizes so metadata nodes are inspectable when toggled on.
- Color Legend is derived from loaded node and relationship inventory instead of a hard-coded 10-label list.
- Graph adapter node color/size and edge color now use the same fallback policy.

Build before tests:

```powershell
go build -trimpath -o .tmp\avmatrix.exe .\cmd\avmatrix
npm --prefix avmatrix-web run build
```

Results:

```text
go build -trimpath -o .tmp\avmatrix.exe .\cmd\avmatrix: passed
npm --prefix avmatrix-web run build: passed, built in 25.01s
```

Tests:

```powershell
npm --prefix avmatrix-web run test -- test/unit/constants.test.ts test/unit/filter-panel.test.ts test/unit/graph-links-visibility.test.ts test/unit/graph-edge-visibility-mode.test.ts
npm --prefix avmatrix-web run test -- test/unit/FileTreePanel.dashboard-completeness.test.tsx
npm --prefix avmatrix-web run test
go test ./cmd/... ./internal/... -count=1
```

Results:

```text
focused constants/filter/edge tests: 4 files passed, 39 tests passed
FileTreePanel dashboard completeness: 1 file passed, 2 tests passed
full Web unit suite: 40 files passed, 305 tests passed
go test ./cmd/... ./internal/... -count=1: passed
```

E2E:

```powershell
npm --prefix avmatrix-web run test:e2e -- server-connect.spec.ts -g "selects a repo from landing and loads graph" --workers=1 --timeout=120000
```

Result:

```text
1 passed
test duration: 29.7s
total duration: 35.3s
```

Runtime UI inspection with Playwright:

```text
repo: AVmatrix
Community (930): 1
Property (3106): 1
TypeAlias (70): 1
Accesses (5031): 1
Has Property (2779): 1
Step In Process (2376): 1
```

Benchmark ledger updated:

- `B1 - Dashboard Completeness After Node Type Slice`
- `B2 - Dashboard Completeness After Edge Type Slice`

### E3C - Canvas Relationship Preservation and Visual Scale Bound

Date: 2026-05-19

Scope:

- `avmatrix-web/src/lib/graph-adapter.ts`
- `avmatrix-web/src/hooks/useSigma.ts`
- `avmatrix-web/test/unit/graph-adapter.edge-geometry.test.ts`
- `docs/plans/2026-05-19-web-ui-left-dashboard-graph-display-completeness-plan.md`
- `docs/plans/2026-05-19-web-ui-left-dashboard-graph-display-completeness-benchmark.md`
- `docs/plans/2026-05-19-web-ui-left-dashboard-graph-display-completeness-evidence.md`

AVmatrix context and impact:

```powershell
avmatrix status
avmatrix context knowledgeGraphToGraphology --repo AVmatrix
avmatrix impact knowledgeGraphToGraphology --repo AVmatrix --direction upstream --depth 2 --include-tests
```

Result:

```text
status: stale before refresh, index at 9728cd0 while worktree was at adbf58d
direct callers: GraphCanvas.tsx, graph-adapter.edge-geometry.test.ts
depth-2 affected: App.tsx, GraphCanvas.selection-performance.test.tsx
risk: low, adapter contract and UI canvas behavior require focused Web tests plus e2e graph-load validation
```

AVmatrix refresh:

```powershell
.\avmatrix-launcher\server-bundle\avmatrix.exe analyze E:\AVmatrix-GO --force --skip-agents-md --no-stats
```

Result:

```text
analyzed E:\AVmatrix-GO
files: scanned=686 parsed=528 unsupported=158 failed=0
graph: nodes=20436 relationships=51176 path=E:\AVmatrix-GO\.avmatrix\graph.json
```

Implementation:

- `knowledgeGraphToGraphology` now creates a graphology `MultiDirectedGraph` so multiple relationships between the same source and target are preserved.
- Every renderable relationship is added with a stable edge key using the relationship id when available, with a deterministic fallback key and duplicate suffix.
- Each edge keeps its own `relationType`, so existing edge visibility logic filters parallel edges by relationship type instead of losing them.
- Large-graph node scaling now caps the base scaled size by graph density.
- `useSigma` caps final rendered node size after selection, highlight, ripple, and animation reducer multipliers.
- Graph-adapter tests now cover `CALLS + USES` and `HAS_PROPERTY + ACCESSES` parallel relationship pairs.
- Graph-adapter tests also cover bounded node-size behavior across small, medium, large, and very large graph sizes.

Benchmark measurement:

```json
{
  "inputNodes": 20436,
  "inputRelationships": 51176,
  "renderableRelationships": 51176,
  "graphologyNodes": 20436,
  "graphologyEdges": 51176,
  "relationshipLoss": 0,
  "parallelSourceTargetPairs": 1412,
  "conversionMs": 478.37,
  "currentLargeGraphProjectSize": 3,
  "currentLargeGraphPropertySize": 1,
  "currentLargeGraphRadiusRatio": 3,
  "renderedSizeCap": 9
}
```

Build before tests:

```powershell
go build -trimpath -o .tmp\avmatrix.exe .\cmd\avmatrix
npm --prefix avmatrix-web run build
```

Results:

```text
go build -trimpath -o .tmp\avmatrix.exe .\cmd\avmatrix: passed
npm --prefix avmatrix-web run build: passed, built in 24.55s
```

Tests:

```powershell
npm --prefix avmatrix-web run test -- test/unit/graph-adapter.edge-geometry.test.ts test/unit/selected-graph-context.test.ts test/unit/graph-edge-visibility-mode.test.ts
npm --prefix avmatrix-web run test -- test/unit/graph-adapter.edge-geometry.test.ts
npm --prefix avmatrix-web run test
go test ./cmd/... ./internal/... -count=1
go test ./internal/embeddings -run TestHTTPEmbedderDoesNotRetryTimeout -count=3 -v
```

Results:

```text
focused graph adapter/context/edge tests: 3 files passed, 10 tests passed
focused graph-adapter edge geometry rerun after HAS_PROPERTY/ACCESSES fixture: 1 file passed, 4 tests passed
full Web unit suite: 40 files passed, 309 tests passed
go test ./cmd/... ./internal/... -count=1: failed once in unrelated internal/embeddings TestHTTPEmbedderDoesNotRetryTimeout with calls = 0, want 1
targeted embeddings rerun: passed 3 / 3
```

E2E:

```powershell
npm --prefix avmatrix-web run test:e2e -- server-connect.spec.ts -g "selects a repo from landing and loads graph" --workers=1 --timeout=120000
```

Result:

```text
1 passed
test duration: 29.2s
total duration: 34.5s
```

Benchmark ledger updated:

- `B3 - Graph Adapter Performance and Preservation`
- `B4B - Node Visual Scale Bound`

### E3D - Web Runtime Diagnostics and Post-Load Connection Stability

Date: 2026-05-19

Scope:

- `avmatrix-web/src/lib/runtime-diagnostics.ts`
- `avmatrix-web/src/services/backend-client.ts`
- `avmatrix-web/src/App.tsx`
- `avmatrix-web/src/components/GraphCanvas.tsx`
- `avmatrix-web/src/hooks/useSigma.ts`
- `avmatrix-web/test/unit/heartbeat.test.ts`
- `avmatrix-web/test/unit/runtime-diagnostics.test.ts`
- `avmatrix-web/e2e/server-connect.spec.ts`
- `docs/plans/2026-05-19-web-ui-left-dashboard-graph-display-completeness-plan.md`
- `docs/plans/2026-05-19-web-ui-left-dashboard-graph-display-completeness-benchmark.md`
- `docs/plans/2026-05-19-web-ui-left-dashboard-graph-display-completeness-evidence.md`

AVmatrix refresh and impact:

```powershell
.\avmatrix-launcher\server-bundle\avmatrix.exe analyze E:\AVmatrix-GO --force --skip-agents-md --no-stats
.\avmatrix-launcher\server-bundle\avmatrix.exe context connectHeartbeat --repo AVmatrix
.\avmatrix-launcher\server-bundle\avmatrix.exe impact connectHeartbeat --repo AVmatrix --direction upstream --depth 2 --include-tests
.\avmatrix-launcher\server-bundle\avmatrix.exe context useSigma --repo AVmatrix
.\avmatrix-launcher\server-bundle\avmatrix.exe impact useSigma --repo AVmatrix --direction upstream --depth 2 --include-tests
```

Results:

```text
analyze: files scanned=686 parsed=528 unsupported=158 failed=0; graph nodes=20476 relationships=51211
connectHeartbeat direct callers: App.tsx and heartbeat.test.ts
connectHeartbeat upstream risk: LOW, impacted count 5
useSigma direct caller: GraphCanvas.tsx
useSigma upstream risk: LOW, impacted count 4
```

Implementation:

- Added `runtime-diagnostics.ts` with `window.__AVMATRIX_WEB_DIAGNOSTICS__` and `window.__AVMATRIX_RESET_WEB_DIAGNOSTICS__` test hooks.
- Recorded graph conversion count/timing and graph node/relationship counts in `GraphCanvas`.
- Recorded layout start/stop, duration budget, active state, and noverlap duration in `useSigma`.
- Recorded heartbeat connect/reconnect counters in `connectHeartbeat`.
- Recorded reconnect-banner transitions in `App` and added `data-testid="server-reconnect-banner"`.
- Added unit tests for heartbeat diagnostics and runtime diagnostic counters.
- Added e2e coverage that loads a large graph, waits for active layout, holds a `30s` post-load stability window, and asserts no heartbeat reconnect or reconnect banner.

Build before tests:

```powershell
go build -trimpath -o .tmp\avmatrix.exe .\cmd\avmatrix
npm --prefix avmatrix-web run build
```

Results:

```text
go build -trimpath -o .tmp\avmatrix.exe .\cmd\avmatrix: passed
npm --prefix avmatrix-web run build: passed, built in 29.09s
```

Tests:

```powershell
npm --prefix avmatrix-web run test -- test/unit/heartbeat.test.ts test/unit/runtime-diagnostics.test.ts test/unit/GraphCanvas.selection-performance.test.tsx
npm --prefix avmatrix-web run test
```

Results:

```text
focused runtime diagnostics tests: 3 files passed, 12 tests passed
full Web unit suite: 41 files passed, 312 tests passed, duration 33.69s
```

E2E:

```powershell
npm --prefix avmatrix-web run test:e2e -- server-connect.spec.ts -g "keeps connection stable after large graph load and layout window" --workers=1 --timeout=120000
```

Result:

```text
1 passed
test duration: 1.1m
total duration: 1.3m
```

Measured post-load stability diagnostics:

```json
{
  "repo": "Restaurant_manager",
  "elapsedMs": 62104,
  "graphConversion": {
    "count": 2,
    "lastMs": 1231.4,
    "maxMs": 2446.1,
    "lastNodeCount": 78350,
    "lastRelationshipCount": 130497
  },
  "layout": {
    "starts": 2,
    "stops": 1,
    "isRunning": true,
    "lastDurationBudgetMs": 45000,
    "lastNoverlapMs": 0
  },
  "heartbeat": {
    "connects": 1,
    "reconnects": 0
  },
  "reconnectBanner": {
    "shows": 0,
    "visible": false
  },
  "backendPidBeforeAfter": "524 / 524"
}
```

Notes:

- An initial e2e variant that required complete layout stop was rejected as the wrong assertion for very large graphs because ForceAtlas2 can legitimately keep running beyond an `80s` poll window in dev mode.
- The final e2e assertion targets the actual regression: while heavy graph layout is active and the page is alive, the backend connection must remain stable and the reconnect banner must not appear.
- Launcher heartbeat age is not an active shutdown budget anymore. Launcher-side stale heartbeat gaps are covered by the `B4A` tests with `3h` and `24h` gaps.

Benchmark ledger updated:

- `B4C - Post-Load Connection Stability`

### E3E - Layout Hierarchy Expansion and Legend Accuracy

Date: 2026-05-19

Scope:

- `avmatrix-web/src/lib/graph-adapter.ts`
- `avmatrix-web/src/components/FileTreePanel.tsx`
- `avmatrix-web/test/unit/graph-adapter.edge-geometry.test.ts`
- `avmatrix-web/test/unit/FileTreePanel.dashboard-completeness.test.tsx`
- `docs/plans/2026-05-19-web-ui-left-dashboard-graph-display-completeness-plan.md`
- `docs/plans/2026-05-19-web-ui-left-dashboard-graph-display-completeness-benchmark.md`
- `docs/plans/2026-05-19-web-ui-left-dashboard-graph-display-completeness-evidence.md`

AVmatrix refresh and impact:

```powershell
.\avmatrix-launcher\server-bundle\avmatrix.exe analyze E:\AVmatrix-GO --force --skip-agents-md --no-stats
.\avmatrix-launcher\server-bundle\avmatrix.exe impact knowledgeGraphToGraphology --repo AVmatrix --direction upstream --depth 2 --include-tests
.\avmatrix-launcher\server-bundle\avmatrix.exe impact FileTreePanel --repo AVmatrix --direction upstream --depth 2 --include-tests
```

Results:

```text
analyze: files scanned=688 parsed=530 unsupported=158 failed=0; graph nodes=20548 relationships=51416
knowledgeGraphToGraphology upstream risk: LOW, impacted count 6
FileTreePanel upstream risk: LOW, impacted count 3
```

Implementation:

- Added a priority-based layout parent policy in `graph-adapter.ts`.
- Added source-parent hierarchy for `HAS_METHOD`, `HAS_PROPERTY`, and `WRAPS`.
- Added target-parent grouping for `STEP_IN_PROCESS`, `ENTRY_POINT_OF`, `HANDLES_ROUTE`, `HANDLES_TOOL`, and lower-priority `MEMBER_OF`.
- Added `Process`, `Community`, `Route`, and `Tool` as root/grouping node types for initial layout placement.
- Kept community membership lower priority than containment/ownership so community grouping does not override stronger structural ownership.
- Added graph-adapter tests proving `STEP_IN_PROCESS` can override broad file `DEFINES` for process layout and `HAS_PROPERTY` keeps properties near the owning type.
- Added loaded-graph counts to node and edge legend entries.
- Added a separate community color-set legend entry when `MEMBER_OF` relationships are present.
- Documented the dashboard relationship boundary: Web dashboard controls graph payload `GraphRelationship.type` / generated `GRAPH_RELATIONSHIP_TYPES`; LadybugDB-only constants remain backend storage/query vocabulary unless serialized into graph payloads.

Current graph hierarchy inventory:

```text
CONTAINS=1,710
DEFINES=17,261
IMPORTS=3,727
HAS_METHOD=337
HAS_PROPERTY=2,829
MEMBER_OF=3,860
STEP_IN_PROCESS=2,356
ENTRY_POINT_OF=640
previous hierarchy candidates=22,698
added owner/process/member candidates=10,022
total hierarchy candidates=32,720
```

Build before tests:

```powershell
go build -trimpath -o .tmp\avmatrix.exe .\cmd\avmatrix
npm --prefix avmatrix-web run build
```

Result:

```text
go build -trimpath -o .tmp\avmatrix.exe .\cmd\avmatrix: passed
npm --prefix avmatrix-web run build: passed, built in 23.40s
```

Tests:

```powershell
npm --prefix avmatrix-web run test -- test/unit/graph-adapter.edge-geometry.test.ts test/unit/FileTreePanel.dashboard-completeness.test.tsx
```

Result:

```text
2 files passed, 8 tests passed
```

Benchmark ledger updated:

- `B2B - Legend Count and Community-Color Coverage`
- `B3B - Layout Hierarchy Expansion`

### E3F - Dense Graph Node-Size Ratio Correction

Date: 2026-05-19

Scope:

- `avmatrix-web/src/lib/graph-adapter.ts`
- `avmatrix-web/src/lib/runtime-diagnostics.ts`
- `avmatrix-web/src/components/GraphCanvas.tsx`
- `avmatrix-web/e2e/server-connect.spec.ts`
- `avmatrix-web/test/unit/graph-adapter.edge-geometry.test.ts`
- `avmatrix-web/test/unit/runtime-diagnostics.test.ts`
- `avmatrix-web/test/unit/GraphCanvas.selection-performance.test.tsx`
- plan, benchmark, and evidence ledgers

AVmatrix context and impact:

```powershell
.\avmatrix-launcher\server-bundle\avmatrix.exe context getScaledNodeSize --repo AVmatrix
.\avmatrix-launcher\server-bundle\avmatrix.exe impact getScaledNodeSize --repo AVmatrix --direction upstream --depth 2 --include-tests
.\avmatrix-launcher\server-bundle\avmatrix.exe context getNodeSize --repo AVmatrix
```

Result summary:

- `getScaledNodeSize` directly affects `graph-adapter.ts` graph conversion and `graph-adapter.edge-geometry.test.ts`.
- `getNodeSize` remains the source of base label sizes; the dense-graph cap now corrects disproportionate labels during graph conversion.
- High-risk surface is Web graph rendering, so this slice used focused Web unit tests, Web build, and an e2e visual-scale check.

Implementation:

- Dense graphs over `20,000` nodes now use a generic base-size cap of `3.0`.
- `Package` has an explicit dense cap of `1.5`.
- `Section` has an explicit dense cap of `1.0`.
- Runtime diagnostics now record `visualScale.maxSizeByLabel` so e2e and later audits can inspect every loaded label's rendered base size.
- E2E visual-scale checks now assert `maxNodeSize <= 3`, `structuralToLeafRatio <= 3`, `Package <= 1.5`, and `Section <= 1.0`.

Focused unit tests:

```powershell
npm --prefix avmatrix-web run test -- test/unit/graph-adapter.edge-geometry.test.ts test/unit/runtime-diagnostics.test.ts test/unit/GraphCanvas.selection-performance.test.tsx
```

Result:

```text
passed, 3 files / 12 tests
```

Full Web unit suite:

```powershell
npm --prefix avmatrix-web run test
```

Result:

```text
passed, 41 files / 316 tests, duration 33.34s
```

Build before e2e:

```powershell
npm --prefix avmatrix-web run build
```

Result:

```text
passed, built in 45.62s
```

E2E:

```powershell
npm --prefix avmatrix-web run test:e2e -- server-connect.spec.ts -g "keeps loaded graph node visual scale bounded" --workers=1 --timeout=120000
```

Result:

```text
passed, 1 / 1
test duration: 38.0s
total duration: 46.0s
```

Large-graph browser diagnostics:

```json
{
  "repo": "Restaurant_manager",
  "visualScale": {
    "maxSizeByLabel": {
      "Folder": 3,
      "Route": 2,
      "Package": 1.5,
      "Community": 1,
      "Process": 1.6,
      "File": 2.4,
      "Function": 1.6,
      "Struct": 3,
      "Class": 3,
      "Method": 1.2,
      "Section": 1,
      "Const": 1,
      "Variable": 1,
      "Interface": 2.8,
      "Property": 1,
      "TypeAlias": 1.2,
      "Constructor": 1.6
    },
    "nodeCount": 78350,
    "minNodeSize": 1,
    "maxNodeSize": 3,
    "maxRenderedNodeSizeCap": 9,
    "structuralToLeafRatio": 3
  },
  "graphConversion": {
    "count": 2,
    "lastMs": 2692.6,
    "maxMs": 2692.6,
    "lastNodeCount": 78350,
    "lastRelationshipCount": 130497
  }
}
```

All known node-label dense-size check at `78,350` nodes:

```text
max known-label size: 3.0
min known-label size: 1.0
max/min ratio: 3x
Package: 1.5
Section: 1.0
```

Benchmark ledger updated:

- `B4D - Dense Graph 3x Node-Size Ratio Correction`
- `B5 - Final Benchmark`

## E4 - Final Closure Evidence

Status: completed

Date: 2026-05-19

Final implementation scope added after `E3E`:

- `avmatrix-web/src/lib/runtime-diagnostics.ts`
- `avmatrix-web/src/components/GraphCanvas.tsx`
- `avmatrix-web/src/components/FileTreePanel.tsx`
- `avmatrix-web/e2e/server-connect.spec.ts`
- `avmatrix-web/test/unit/GraphCanvas.selection-performance.test.tsx`
- `avmatrix-web/test/unit/runtime-diagnostics.test.ts`
- `internal/session/controller_test.go`
- `internal/embeddings/http_client_test.go`
- plan, benchmark, and evidence ledgers

Implementation notes:

- Added visual-scale diagnostics to `window.__AVMATRIX_WEB_DIAGNOSTICS__` so e2e can assert loaded graph node scale bounds.
- Added `aria-pressed` to Node Type and Edge Type controls so e2e can verify toggle state.
- Added e2e coverage for uncommon `Property` node and `Accesses` edge controls.
- Added e2e visual-scale coverage using loaded graph diagnostics.
- Fixed two validation-blocking flaky tests:
  - `internal/session/controller_test.go` now waits for the first chat to actually start before starting the second chat and protects fake adapter run records with a mutex.
  - `internal/embeddings/http_client_test.go` now uses a deterministic blocking `RoundTripper` for timeout/no-retry behavior instead of racing a `1ms` timeout against an `httptest.Server` handler.

Final AVmatrix analyze:

```powershell
.\avmatrix-launcher\server-bundle\avmatrix.exe analyze E:\AVmatrix-GO --force --skip-agents-md --no-stats
```

Result:

```text
analyzed E:\AVmatrix-GO
files: scanned=688 parsed=530 unsupported=158 failed=0
graph: nodes=20611 relationships=51507 path=E:\AVmatrix-GO\.avmatrix\graph.json
```

Build before final tests:

```powershell
go build -trimpath -o .tmp\avmatrix.exe .\cmd\avmatrix
npm --prefix avmatrix-web run build
```

Results:

```text
go build -trimpath -o .tmp\avmatrix.exe .\cmd\avmatrix: passed
npm --prefix avmatrix-web run build: passed, built in 23.97s
```

Focused validation:

```powershell
npm --prefix avmatrix-web run test -- test/unit/runtime-diagnostics.test.ts test/unit/FileTreePanel.dashboard-completeness.test.tsx
npm --prefix avmatrix-web run test -- test/unit/GraphCanvas.selection-performance.test.tsx test/unit/runtime-diagnostics.test.ts
go test ./internal/session -run TestControllerCancelsPreviousSessionForSameRepo -count=10 -v
go test ./internal/embeddings -run TestHTTPEmbedderDoesNotRetryTimeout -count=10 -v
```

Results:

```text
runtime/FileTreePanel focused Web tests: 2 files passed, 5 tests passed
GraphCanvas/runtime focused Web tests: 2 files passed, 5 tests passed
session cancellation targeted Go test: passed 10 / 10
embedding timeout targeted Go test: passed 10 / 10
```

Full validation:

```powershell
npm --prefix avmatrix-web run test
go test ./cmd/... ./internal/... -count=1
npm --prefix avmatrix-web run test:e2e -- server-connect.spec.ts -g "Graph Dashboard Controls" --workers=1 --timeout=120000
npm --prefix avmatrix-web run test:e2e -- server-connect.spec.ts -g "keeps connection stable after large graph load and layout window" --workers=1 --timeout=120000
```

Results:

```text
npm --prefix avmatrix-web run test: passed, 41 files / 316 tests, duration 33.34s
go test ./cmd/... ./internal/... -count=1: passed
Graph Dashboard Controls e2e: passed, 2 / 2, duration 1.6m
post-load connection stability e2e: passed, 1 / 1, duration 1.2m
```

Final browser visual-scale diagnostics:

```json
{
  "repo": "Restaurant_manager",
  "visualScale": {
    "maxSizeByLabel": {
      "Folder": 3,
      "Route": 2,
      "Package": 1.5,
      "Community": 1,
      "Process": 1.6,
      "File": 2.4,
      "Function": 1.6,
      "Struct": 3,
      "Class": 3,
      "Method": 1.2,
      "Section": 1,
      "Const": 1,
      "Variable": 1,
      "Interface": 2.8,
      "Property": 1,
      "TypeAlias": 1.2,
      "Constructor": 1.6
    },
    "nodeCount": 78350,
    "minNodeSize": 1,
    "maxNodeSize": 3,
    "maxRenderedNodeSizeCap": 9,
    "structuralToLeafRatio": 3
  },
  "graphConversion": {
    "count": 2,
    "lastMs": 2692.6,
    "maxMs": 2692.6,
    "lastNodeCount": 78350,
    "lastRelationshipCount": 130497
  }
}
```

Final closure result:

- dashboard node controls: complete;
- dashboard edge controls: complete;
- legend counts and community-color distinction: complete;
- parallel relationship preservation: complete;
- visual-scale proportionality: complete;
- post-load connection stability: complete;
- full build, full Web unit, scoped Go suite, required e2e tests, and final analyze: passed.
