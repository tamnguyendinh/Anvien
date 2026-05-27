# AVmatrix Web Graph Dynamic Scale Model And Zoom Semantics Evidence Ledger

Date: 2026-05-27

Status: Planned

Companion files:

- Plan: [2026-05-27-avmatrix-web-graph-dynamic-scale-model-plan.md](2026-05-27-avmatrix-web-graph-dynamic-scale-model-plan.md)
- Benchmark ledger: [2026-05-27-avmatrix-web-graph-dynamic-scale-model-benchmark.md](2026-05-27-avmatrix-web-graph-dynamic-scale-model-benchmark.md)

## Evidence Rules

Evidence must be recorded as soon as each evidenced task is completed.

This ledger is for traceability, implementation decisions, command outputs, screenshots, blast-radius summaries, and validation results.

Benchmark values belong in the benchmark ledger. Do not mix benchmark tables into this evidence ledger except by linking to their benchmark IDs.

## E0 - Plan Creation Context

Date: 2026-05-27

Status: recorded.

Rule correction:

```text
Only skip AVmatrix when running git commit for documentation-only staged changes. All docs technical planning, evidence, benchmark, report, and architecture work must use AVmatrix and read the codebase.
```

Discussion summary:

- The previous spacing hardening work solved a narrow overlap/readability symptom but changed broader graph UX behavior.
- The graph previously preserved a default overview with multiple visible colors/islands.
- After the spacing/camera changes, default load focuses a dense local `frontend:Function` area, causing visible nodes to appear as one color.
- The first repair gate restores the original color overview behavior.
- The first repair gate proves no node type was lost.
- Any node type present in baseline graph/filter inventory and visibly rendered in baseline overview remains present after the fix.
- The deeper issue is not color assignment. The graph algorithm contains fixed-size and fixed-camera assumptions where the product requires dynamic relationships.
- When node size changes, layout spacing, island radius, ring radius, camera behavior, edge styling, labels, diagnostics, and tests adapt together.
- User zoom visibly enlarges nodes; a zoom action that only changes position scale while nodes remain tiny violates expected graph interaction semantics.

Core failure statement:

```text
One fixed-size assumption can invalidate the whole graph geometry. The graph must use a unified dynamic scale model instead of patching node spacing, camera, and diagnostics independently.
```

## E0A - Planning Re-Read And Codebase Audit

Date: 2026-05-27

Status: recorded.

AVmatrix refresh:

```text
.\avmatrix\bin\avmatrix.exe analyze --force
files: scanned=774 parsed=572 unsupported=202 failed=0
graph: nodes=89484 relationships=122631 path=E:\AVmatrix-GO\.avmatrix\graph.json
```

AVmatrix commands used:

```text
.\avmatrix\bin\avmatrix.exe query "web graph sigma camera zoom node size color overview" --repo AVmatrix
.\avmatrix\bin\avmatrix.exe context "useSigma" --repo AVmatrix
.\avmatrix\bin\avmatrix.exe context "knowledgeGraphToGraphology" --repo AVmatrix
.\avmatrix\bin\avmatrix.exe context "applyReadableGraphCamera" --repo AVmatrix
```

Key AVmatrix findings:

- `useSigma` is in `avmatrix-web/src/hooks/useSigma.ts` with incoming use from `GraphCanvas.tsx`.
- `applyReadableGraphCamera` is in `avmatrix-web/src/lib/graph-readable-camera.ts` with incoming call from `useSigma.ts`.
- `knowledgeGraphToGraphology` is in `avmatrix-web/src/lib/graph-adapter.ts`; tests in `graph-adapter.edge-geometry.test.ts` exercise it heavily.
- `buildLayoutNodeSpacingDiagnostics` is in `GraphCanvas.tsx` and participates in Web graph UI diagnostics.

Code findings:

- `avmatrix-web/src/hooks/useSigma.ts:257` sets `minCameraRatio: MIN_READABLE_CAMERA_RATIO`.
- `avmatrix-web/src/hooks/useSigma.ts:259` sets `itemSizesReference: 'positions'`.
- `avmatrix-web/src/hooks/useSigma.ts:518-520` sets camera state to `{ x: 0.5, y: 0.5, ratio: 1, angle: 0 }`, refreshes, then calls `applyReadableGraphCamera(sigma)`.
- `avmatrix-web/src/lib/graph-readable-camera.ts:53-75` picks the densest visible island.
- `avmatrix-web/src/lib/graph-readable-camera.ts:149-159` focuses that island and computes a readable ratio.
- `avmatrix-web/src/lib/graph-adapter.ts:86-107` uses fixed rendered size and fixed minimum center distance.
- `avmatrix-web/src/lib/graph-adapter.ts:590-593` derives island gap from fixed multiples.
- `avmatrix-web/src/lib/graph-adapter.ts:811-848` assigns node color from `getNodeColor(displayNodeType)`, so the palette remains type-based.
- `avmatrix-web/src/lib/graph-screen-spacing.ts:90-93` uses `sigma.scaleSize` for viewport diagnostics.
- `avmatrix-web/src/lib/runtime-diagnostics.ts` publishes graph conversion, visual scale, layout rings, layout spacing, screen spacing, readable camera, heartbeat, and reconnect diagnostics. It lacks overview visible color count, visible island count, dominant island share, visible node-type inventory, and filter node-type inventory.
- Sigma local package defaults: `itemSizesReference: "screen"` and `zoomToSizeRatioFunction: Math.sqrt`.
- Sigma local package `scaleSize` formula multiplies by `cameraRatio * graphToViewportRatio` only when `itemSizesReference` is `"positions"`.

Git baseline findings:

```text
git rev-parse --short 67ba0dd^
80a7972
```

Commit `67ba0dd` changed the Web graph screen-spacing behavior and introduced:

- `applyReadableGraphCamera`;
- `MIN_READABLE_CAMERA_RATIO`;
- `itemSizesReference: 'positions'`;
- default-load `setState({ x: 0.5, y: 0.5, ratio: 1, angle: 0 })`;
- default-load `applyReadableGraphCamera(sigma)`;
- e2e threshold changes that allowed a single focused dense island to pass.

Decision recorded:

The implementation path restores baseline overview first by reverting the default-load camera/render semantic changes from `67ba0dd`, then introduces a dynamic scale model and deterministic hexagonal packing.

## E1 - Required Pre-Implementation Audit

Status: in progress.

Completed pre-edit evidence on 2026-05-27:

- AVmatrix refresh command: `.\avmatrix\bin\avmatrix.exe analyze --force`
- AVmatrix refresh output: `files: scanned=774 parsed=572 unsupported=202 failed=0`, `graph: nodes=89489 relationships=122636 path=E:\AVmatrix-GO\.avmatrix\graph.json`
- AVmatrix context/query owners: `useSigma`, `applyReadableGraphCamera`, `knowledgeGraphToGraphology`, `buildLayoutNodeSpacingDiagnostics`, `runtime-diagnostics.ts`
- Non-destructive baseline worktree command: `git worktree add .tmp\graph-baseline-80a7972 80a7972`
- Baseline worktree result: detached worktree at commit `80a7972`

Pre-edit impact evidence:

- `useSigma`: CRITICAL, affected frontend GraphCanvas flow and 11 detected processes.
- `applyReadableGraphCamera`: LOW, direct caller `useSigma.ts`.
- `knowledgeGraphToGraphology`: CRITICAL with tests included, affected frontend GraphCanvas, unit tests, and 11 detected processes.
- `buildLayoutNodeSpacingDiagnostics`: LOW in upstream impact.
- `createDiagnostics`: LOW in upstream impact.
- `recordScreenNodeSpacing`: CRITICAL, affected GraphCanvas, runtime diagnostics test, and 11 detected processes.
- `recordCurrentScreenNodeSpacing`: LOW in upstream impact.

Blast-radius decision:

HIGH/CRITICAL results are workflow warnings for central Web graph behavior. The Phase 1 slice proceeds with scoped edits to default graph load, overview diagnostics, and parity tests.

Before implementation edits, record:

- AVmatrix graph refresh command and output;
- AVmatrix context for graph layout, Sigma rendering, camera, and diagnostics owners;
- impact analysis for every function/class/method/exported symbol to be edited;
- blast radius for HIGH/CRITICAL symbols as warnings, not edit bans;
- baseline source `80a7972`, captured through `.tmp/graph-baseline-80a7972`;
- baseline browser screenshot paths and visible color/island/label metrics;
- baseline graph/filter node-type inventory and overview-visible node-type inventory;
- current HEAD browser screenshot paths and visible color/island/label metrics;
- current HEAD graph/filter node-type inventory and overview-visible node-type inventory;
- exact regression delta between baseline and current HEAD;
- exact missing-node-type delta between baseline and current HEAD;
- current fixed-size assumptions in layout, render, camera, diagnostics, and tests;
- Sigma node-size semantics verified from local behavior;
- screenshots and browser diagnostics proving the current regression.

Expected code areas to audit:

- `avmatrix-web/src/hooks/useSigma.ts`
- `avmatrix-web/src/lib/graph-adapter.ts`
- `avmatrix-web/src/lib/graph-readable-camera.ts`
- `avmatrix-web/src/lib/graph-screen-spacing.ts`
- `avmatrix-web/src/lib/runtime-diagnostics.ts`
- `avmatrix-web/src/components/GraphCanvas.tsx`
- `avmatrix-web/e2e/graph-orientation-labels.spec.ts`
- related unit tests under `avmatrix-web/test/unit`

## E1A - Color Overview Parity Gate

Status: in progress.

Baseline capture:

- Source: commit `80a7972`
- Worktree: `.tmp\graph-baseline-80a7972`
- Screenshot: `reports/problem/2026-05-27-graph-baseline-80a7972-dense-overview.png`
- Metrics: `reports/problem/2026-05-27-graph-baseline-80a7972-dense-overview.json`
- Fixture node count: 1480
- Baseline ring label count: 3
- Baseline ring labels: `backend`, `docs`, `frontend`
- Baseline island label count: 4
- Baseline island labels: `backend:Function`, `docs:Documentation`, `frontend:Function`, `frontend:Method`
- Baseline layout island count: 5
- Baseline spacing violations: overlap `0`, target-gap `0`

Current HEAD pre-fix capture:

- Source: current HEAD before Phase 1 edits
- Screenshot: `reports/problem/2026-05-27-graph-current-before-phase1-dense-overview.png`
- Metrics: `reports/problem/2026-05-27-graph-current-before-phase1-dense-overview.json`
- Fixture node count: 1480
- Current ring label count: 2
- Current ring labels: `backend`, `frontend`
- Current island label count: 0
- Current screen visible viewport node count: 40
- Current screen visible island counts: `frontend:Function` = 40
- Current readable camera applied: `true`
- Current readable camera focused island: `frontend:Function`
- Current camera ratio: `0.010029304328422596`

Regression delta:

- Ring labels dropped from 3 to 2.
- Island labels dropped from 4 to 0.
- Default viewport collapsed to a single visible island, `frontend:Function`.
- Readable camera applied during default graph load and focused the densest island.
- Baseline preserved overview labels for backend, docs, frontend, and multiple island labels.

Count interpretation:

- The recorded `3` ring labels and `4` island labels are measured baseline values for this dense fixture only.
- Product behavior must not hardcode those numbers for every repository.
- Tests must derive expected ring, island, color, and node-type inventories from the active fixture/baseline data.
- A repository with different languages and taxonomy displays its own computed inventory count.
- The current dense fixture has `2157` graph nodes and `2077` default-visible nodes after default filters; its computed overview inventory is exactly `3` rings, `4` islands, `3` visible node types, and `3` visible colors.

Phase 1 camera/render correction:

- `avmatrix-web/src/hooks/useSigma.ts` default `setGraph` path now uses `sigma.getCamera().animatedReset({ duration: 500 })`.
- Default graph load no longer calls `applyReadableGraphCamera(sigma)`.
- Overview `minCameraRatio` restored to `0.002`.
- Sigma `itemSizesReference: 'positions'` override removed so Sigma uses default screen sizing semantics.

Phase 1 overview diagnostics correction:

- Added `graphOverview` diagnostics with visible color count, visible ring count, visible island count, dominant island share, visible color counts, visible ring counts, visible island counts, visible node-type counts, graph ring counts, graph node-type counts, visible ring inventory, visible node-type inventory, graph ring inventory, filter node-type inventory, and camera state.
- Added `buildGraphOverviewDiagnostics` to derive inventories from the active graph and viewport.
- Added unit coverage for graph overview diagnostics.
- Updated dense graph e2e assertions to derive expected ring, island, node-type, and node-count inventory from the fixture graph instead of hardcoding global counts.
- Removed default-overview screen-space overlap expectations from the Phase 1 e2e path. Screen-space no-overlap belongs to explicit detail/focus mode in later phases.
- Plan and benchmark targets now treat measured `3` and `4` values as fixture-specific baseline evidence, not global product constants.

Restored Phase 1 capture:

- Screenshot: `reports/problem/2026-05-27-graph-restored-phase1-dense-overview.png`
- Metrics: `reports/problem/2026-05-27-graph-restored-phase1-dense-overview.json`
- Restored ring label count: 3
- Restored ring labels: `backend`, `docs`, `frontend`
- Restored island label count: 4
- Restored island labels: `backend:Function`, `docs:Documentation`, `frontend:Function`, `frontend:Method`
- Restored visible viewport node count: 1400
- Restored visible color count: 3
- Restored visible ring count: 3
- Restored visible island count: 4
- Restored dominant island share: `0.7142857142857143`
- Restored visible ring inventory: `backend`, `docs`, `frontend`
- Restored visible node-type inventory: `Documentation`, `Function`, `Method`
- Restored graph ring inventory: `backend`, `docs`, `frontend`
- Restored graph island inventory: `backend:Function`, `docs:Documentation`, `frontend:Function`, `frontend:Method`
- Restored filter node-type inventory: `Documentation`, `Function`, `Method`
- Restored readable camera applied: `false`
- Restored camera ratio: `1`

Validation completed before Phase 1 pre-commit detection:

- Full Web build: `npm run build` passed.
- Focused Web unit tests: `npx vitest run test/unit/runtime-diagnostics.test.ts test/unit/graph-overview-diagnostics.test.ts` passed, 2 files, 9 tests.
- Web e2e/browser tests: first run failed because the Vite server on `127.0.0.1:5228` was not running; after starting `npm run dev`, `npx playwright test e2e/graph-orientation-labels.spec.ts --project=chromium` passed, 3 tests. The dev server was stopped after the run.
- Browser screenshots captured for baseline, current pre-fix, and restored Phase 1 overview.

AVmatrix pre-commit detection:

- `avmatrix analyze --force` after Phase 1 edits: scanned `779`, parsed `574`, unsupported `205`, failed `0`.
- Refreshed graph: `89896` nodes, `123159` relationships.
- `avmatrix detect-changes --repo AVmatrix --scope all`: changed files `8`, changed symbols `372`.
- Changed app layers: docs `13`, frontend `61`, frontend_test `298`.
- Changed functional areas: documentation `13`, layout `9`, unknown `344`, web_graph_ui `6`.
- Affected processes: `0`.
- Resolution health degraded nodes: `0`.
- Reported risk level: `low`.

This gate must be completed before dynamic scale, zoom, and spacing implementation proceeds.

Required evidence:

- baseline source: commit `80a7972`;
- command and browser steps used to capture the baseline through `.tmp/graph-baseline-80a7972`;
- baseline screenshot paths;
- current HEAD screenshot paths;
- restored screenshot paths;
- visible color count comparison;
- visible ring count comparison;
- visible island count comparison;
- graph/visible ring inventory comparison;
- graph/filter node-type inventory comparison;
- overview-visible node-type inventory comparison;
- missing node-type list, expected empty;
- ring/island label comparison;
- Phase 1 overview diagnostics shape and test evidence;
- decision record: default-load readable camera removed from overview, Sigma screen sizing restored, baseline reset restored;
- test evidence proving a one-color default viewport fails on a multi-color fixture;
- test evidence proving a missing baseline node type fails before dynamic scale work proceeds.

## E1B - Phase 2 Scale, Zoom, Spacing Audit

Status: completed.

AVmatrix refresh:

- `avmatrix analyze --force`: scanned `779`, parsed `574`, unsupported `205`, failed `0`.
- Refreshed graph: `89896` nodes, `123159` relationships.

AVmatrix owner trace:

- Query `graph layout sigma camera zoom rendered node radius spacing diagnostics fixed constants` ranked `knowledgeGraphToGraphology`, `buildLayoutNodeSpacingDiagnostics`, `useSigma`, `getRenderedNodeRadius`, and `capRenderedNodeSize` as primary owners.
- `useSigma` context maps Sigma initialization, camera config, node reducer, edge reducer, selection, zoom buttons, reset, and focus.
- `knowledgeGraphToGraphology` context maps graph conversion, node size, node color, layout, edge size, and display filter linkage.
- `buildGraphOverviewDiagnostics` context maps Phase 1 graph overview inventory diagnostics.

Impact summary:

| Symbol | Risk | Affected processes | Notes |
|---|---:|---:|---|
| `useSigma` | CRITICAL | 11 | central Sigma render, reducer, camera, selection |
| `knowledgeGraphToGraphology` | CRITICAL | 11 | graph conversion, layout, unit tests, GraphCanvas |
| `buildLayoutNodeSpacingDiagnostics` | LOW | 0 | local GraphCanvas layout diagnostics |
| `buildGraphOverviewDiagnostics` | CRITICAL | 11 | overview diagnostics consumed by GraphCanvas and e2e |
| `buildScreenNodeSpacingDiagnostics` | CRITICAL | 11 | screen spacing diagnostics, readable camera, tests |
| `applyReadableGraphCamera` | LOW | 0 | readable camera apply path |
| `getRenderedNodeRadius` | HIGH | 4 | rendered diameter and layout metrics |
| `capRenderedNodeSize` | CRITICAL | 11 | useSigma reducer and edge-geometry tests |
| `getMinimumNodeCenterDistance` | CRITICAL | 15 | GraphCanvas diagnostics, graph layout, edge-geometry tests |
| `getClusterNodeSpacing` | LOW | 0 | graph-adapter layout helper |
| `getIslandOffsets` | LOW | 0 | graph-adapter layout helper |

Blast radius note:

- HIGH and CRITICAL mark central graph UI paths. They are warnings that require scoped edits, not edit bans.

Fixed assumption inventory:

- `graph-adapter.ts`: `MAX_RENDERED_NODE_SIZE = 3`, `MAX_DENSE_RENDERED_NODE_SIZE = 3`, and `getMaxRenderedNodeSize` always returns the dense cap.
- `graph-adapter.ts`: `getClusterNodeSpacing` returns `42` below `1000` nodes and `getMinimumNodeCenterDistance` above that threshold.
- `graph-adapter.ts`: island gap uses `nodeSpacing * 34`, `minimumNodeCenterDistance * 75`, and `largestClusterRadius * 0.55`.
- `graph-adapter.ts`: island orbit uses `largestClusterRadius * 1.85`, adjacent span math, and `nodeSpacing * 28`.
- `graph-adapter.ts`: ring gap uses `nodeSpacing * 70` and `largestOuterRingRadius * 0.75`; outer orbit also uses `largestOuterRingRadius * 2.1` and `nodeSpacing * 72`.
- `graph-adapter.ts`: node size scaling is step-based at `1000`, `5000`, `20000`, and `50000` graph nodes.
- `graph-adapter.ts`: edge base size is step-based at `5000` and `20000` graph nodes.
- `useSigma.ts`: camera config uses `minCameraRatio: 0.002` and `maxCameraRatio: 50`.
- `useSigma.ts`: default load reset duration is `500`; focus camera uses hardcoded `ratio: 0.15` and duration `400`.
- `useSigma.ts`: zoom button durations are `200`; reset duration is `300`.
- `useSigma.ts`: node reducer uses fixed highlight multipliers including `1.8`, `1.6`, `1.4`, `1.3`, `0.6`, `0.5`, and `0.4`.
- `useSigma.ts`: edge reducer uses fixed dense ambient edge size and opacity policy.
- `graph-readable-camera.ts`: readable threshold is `1000` nodes, readable radius target is `2px`, and minimum readable camera ratio is `0.00004`.
- `graph-readable-camera.ts`: readable camera selects the densest visible island.
- `graph-screen-spacing.ts`: screen diagnostics use Sigma `scaleSize`, pair edge gap as largest rendered diameter, and center distance as two radii plus required edge gap.
- `GraphCanvas.tsx`: diagnostics refresh on resize and setGraph. Camera `updated` currently refreshes orientation labels but does not refresh screen diagnostics.

Phase 2 browser artifacts:

- `reports/problem/2026-05-27-graph-phase2-overview.png`
- `reports/problem/2026-05-27-graph-phase2-overview.json`
- `reports/problem/2026-05-27-graph-phase2-zoom-in-1.png`
- `reports/problem/2026-05-27-graph-phase2-zoom-in-1.json`
- `reports/problem/2026-05-27-graph-phase2-zoom-in-2.png`
- `reports/problem/2026-05-27-graph-phase2-zoom-in-2.json`
- `reports/problem/2026-05-27-graph-phase2-zoom-out.png`
- `reports/problem/2026-05-27-graph-phase2-zoom-out.json`
- `reports/problem/2026-05-27-graph-phase2-detail-focus.png`
- `reports/problem/2026-05-27-graph-phase2-detail-focus.json`

Dense browser baseline:

| Stage | Camera ratio | Max radius px | Visible nodes | Visible island inventory | Overlap count | Target-gap violations |
|---|---:|---:|---:|---|---:|---:|
| Overview | 1 | 3 | 1400 | `backend:Function`, `docs:Documentation`, `frontend:Function`, `frontend:Method` | 27425 | 97426 |
| Zoom in 1 | 0.6666666666666666 | 3.6742346141747673 | 40 | `docs:Documentation` | 18411 | 68085 |
| Zoom in 2 | 0.4444444444444444 | 4.5 | 0 | empty | 12240 | 47002 |
| Zoom out | 0.6666666666666666 | 3.6742346141747673 | 40 | `docs:Documentation` | 18411 | 68085 |

Zoom audit result:

- Rendered radius grows when camera ratio decreases, so Sigma screen sizing now provides visible zoom growth.
- Default zoom from overview collapses visible inventory from full dense overview to `docs:Documentation`, then to zero visible nodes.
- Runtime diagnostics did not update after zoom until a resize event was dispatched; Phase 6 must attach screen diagnostics refresh to camera updates.

Detail/focus baseline:

- Fixture: single-node focus fixture.
- Selection succeeded: `true`.
- Focus button clicked: `true`.
- Camera ratio after focus: `1`.
- Max rendered radius after focus: `3`.
- Visible viewport node count after focus: `1`.
- Code audit explains this: `useSigma.focusNode` skips camera animation when `selectedNodeRef.current === nodeId`, so focusing a node that was just selected by click does not change camera ratio.

## E2 - Dynamic Scale Model Evidence

Status: completed for Phase 3 scale-model slice.

Implemented files:

- Added `avmatrix-web/src/lib/graph-scale-model.ts`.
- Added `avmatrix-web/test/unit/graph-scale-model.test.ts`.
- Updated `graph-screen-spacing.ts` to use `getRequiredCenterDistancePx`.
- Updated `graph-adapter.ts` rendered-size exports to read from `GRAPH_RENDER_SCALE_POLICY`.

Scale model behavior:

- `getRenderedNodeRadiusPx` delegates node size and camera ratio to Sigma `scaleSize`.
- `measureGraphUnitToViewportPx` measures graph-coordinate distance through `graphToViewport`.
- `getRequiredEdgeGapPx` uses the largest rendered diameter of the compared nodes.
- `getRequiredCenterDistancePx` uses left radius plus right radius plus required edge gap.
- `getRequiredCenterDistanceGraph` converts required viewport pixels back to graph units.
- `buildGraphScaleModel` records camera ratio, graph-unit scale, min/max rendered radius, max diameter, required edge gap, required center distance in pixels, and required center distance in graph units.

Focused validation:

- Full Web build: `npm run build` passed.
- Focused tests: `npx vitest run test/unit/graph-scale-model.test.ts test/unit/graph-screen-spacing.test.ts test/unit/graph-adapter.edge-geometry.test.ts` passed, 3 files, 34 tests.

Measured unit facts:

- Radius with size `3` at camera ratio `1`: `3`.
- Radius with size `3` at camera ratio `0.25`: `6`.
- Equal-size radius `3` required edge gap: `6`.
- Equal-size radius `3` required center distance: `12`.
- Graph-unit scale from projection `x * 5`: `5px` per graph unit.
- Built model at camera ratio `0.25` with node sizes `2` and `3`: min radius `4`, max radius `6`, required edge gap `12`, required center distance `24px`, required center distance `12` graph units.

AVmatrix pre-commit detection for Phase 3:

- `avmatrix analyze --force`: scanned `786`, parsed `576`, unsupported `210`, failed `0`.
- Refreshed graph: `90067` nodes, `123382` relationships.
- `avmatrix detect-changes --repo AVmatrix --scope all`: risk level `low`.
- Changed files: `5`.
- Changed symbols: `12`.
- Changed app layers: docs `9`, frontend `3`.
- Changed functional areas: documentation `9`, layout `2`, unknown `1`.
- Affected count: `0`.
- Affected processes: `0`.
- Resolution gap changes: `0`.
- Resolution health impact: degraded nodes `0`, total resolution gap count `0`.
- `git diff --check`: passed.

## E2A - Dynamic Hex Layout Evidence

Status: completed for Phase 4 layout slice.

Blast radius:

- `applyFilterBasedClusteredLayout`: CRITICAL, affected processes `15`.
- `getIslandOffset`, `getFallbackIslandOffset`, `isIslandOffsetFarEnough`, `addIslandOffsetToGrid`: HIGH, affected processes `4` each.
- `getClusterNodeSpacing`, `getClusterIslandRadius`, `getIslandOffsets`, `getPinwheelRadiusMultiplier`, `getPinwheelAngleOffset`: LOW.
- HIGH and CRITICAL were treated as required-care warnings, not edit bans.

Implemented files:

- Updated `avmatrix-web/src/lib/graph-adapter.ts`.
- Updated `avmatrix-web/test/unit/graph-adapter.edge-geometry.test.ts`.

Layout behavior:

- Same-island placement now uses deterministic axial hex coordinates.
- Cell spacing derives from `getMinimumNodeCenterDistance`, which is tied to rendered node diameter plus required edge gap.
- The old spiral attempt search, distance grid, secondary spiral, organic wave, and pinwheel radius/angle jitter were removed from protected no-overlap paths.
- Island radius derives from final offset bounds plus dynamic footprint margin.
- Island and ring gaps derive from adjacent footprint radii plus dynamic center-distance floor.
- Unit tests now compare island/ring whitespace to dynamic graph gap floors instead of fixed `100`, `200`, `500`, and `900` graph-unit thresholds.

Validation:

- Full Web build: `npm run build` passed.
- Focused Web unit tests: `npx vitest run test/unit/graph-scale-model.test.ts test/unit/graph-screen-spacing.test.ts test/unit/graph-adapter.edge-geometry.test.ts` passed, 3 files, 34 tests.
- Web e2e/browser tests: `npx playwright test e2e/graph-orientation-labels.spec.ts --project=chromium` passed, 3 tests.
- Vite dev server was started on `127.0.0.1:5228` for e2e and stopped after the run.

Screenshot validation:

- `avmatrix-web/test-results/graph-orientation-labels-G-5b603-labels-on-the-desktop-graph-chromium/graph-orientation-labels-desktop.png`
- `avmatrix-web/test-results/graph-orientation-labels-G-e5a0a-t-and-updates-after-filters-chromium/graph-orientation-labels-small-filtered.png`
- `avmatrix-web/test-results/graph-orientation-labels-G-537e2-ory-visible-on-default-load-chromium/graph-node-spacing-dense-desktop.png`
- `avmatrix-web/test-results/graph-orientation-labels-G-537e2-ory-visible-on-default-load-chromium/graph-node-spacing-dense-small.png`

AVmatrix pre-commit detection for Phase 4:

- `avmatrix analyze --force`: scanned `786`, parsed `576`, unsupported `210`, failed `0`.
- Refreshed graph: `90038` nodes, `123327` relationships.
- `avmatrix detect-changes --repo AVmatrix --scope all`: risk level `low`.
- Changed files: `5`.
- Changed symbols: `145`.
- Changed app layers: docs `6`, frontend `69`, frontend_test `70`.
- Changed functional areas: documentation `6`, layout `69`, unknown `70`.
- Affected count: `0`.
- Affected processes: `0`.
- Resolution gap changes: `109` bookkeeping entries, actionability analyzer_gap `106`, non_actionable `3`.
- Changed source nodes with gaps: `0`.
- Resolution health impact: degraded nodes `0`, total resolution gap count `0`.
- `git diff --check`: passed.

## E3 - Overview Preservation Evidence

Status: completed for Phase 5 camera and zoom slice.

AVmatrix refresh before Phase 5 edits:

- `avmatrix analyze --force`: scanned `786`, parsed `576`, unsupported `210`, failed `0`.
- Refreshed graph: `90038` nodes, `123327` relationships.

Phase 5 impact summary:

| Symbol | Risk | Affected processes | Notes |
|---|---:|---:|---|
| `useSigma` | CRITICAL | 11 | central Sigma camera, selection, zoom, reducer path |
| `focusNode` variable in `useSigma.ts` | LOW | 0 | direct focus behavior implementation |
| `GraphCanvas` | LOW | 0 | imperative focus handle and diagnostics event wiring |
| `recordCurrentScreenNodeSpacing` | LOW | 0 | browser diagnostics refresh path |
| `applyReadableGraphCamera` | LOW | 0 | confirmed default load has no incoming use |

Blast-radius decision:

- CRITICAL on `useSigma` is a required-care warning because graph camera behavior is central Web UI behavior.
- The edit stayed scoped to graph camera mode, focus coordinates, diagnostics refresh, and tests.

Implemented overview behavior:

- Default graph load still uses overview reset through `buildOverviewCameraAction`.
- Default graph load still keeps readable/detail camera out of initial load.
- Graph overview diagnostics remain inventory-driven.
- For the current dense fixture, runtime inventory recorded `2157` graph nodes, `2077` default-visible nodes, `3` visible colors, `3` visible rings, `4` visible islands, and `3` visible node types.
- These `3`, `4`, and `2077` values are measured dense-fixture inventory values. Repositories with different language/taxonomy content compute their own exact values from active graph inventory.

Screenshot artifacts:

- `reports/problem/2026-05-27-graph-phase5-overview.png`
- `avmatrix-web/test-results/graph-orientation-labels-G-537e2-ory-visible-on-default-load-chromium/graph-node-spacing-dense-desktop.png`
- `avmatrix-web/test-results/graph-orientation-labels-G-537e2-ory-visible-on-default-load-chromium/graph-node-spacing-dense-small.png`

## E4 - Zoom Semantics Evidence

Status: completed for Phase 5 camera and zoom slice.

Implemented zoom behavior:

- Camera `updated` events now refresh screen spacing and overview diagnostics.
- E2E waits for visible camera-ratio movement instead of accepting the first animation tick.
- Zoom assertions require meaningful radius growth and shrink, not a tiny early-frame delta.

Measured dense-fixture zoom sequence:

| Stage | Camera ratio | Max rendered radius px | Visible viewport nodes | Visible islands |
|---|---:|---:|---:|---:|
| Initial overview | 1 | 3 | 2077 | 4 |
| Zoom in 1 | 0.6666666666666666 | 3.6742346141747673 | 773 | 2 |
| Zoom in 2 | 0.4444444444444444 | 4.5 | 40 | 1 |
| Zoom out | 0.6666666666666666 | 3.6742346141747673 | 773 | 2 |

Radius result:

- Zoom-in 1 radius growth: `1.2247448713915892`.
- Zoom-in 2 radius growth: `1.224744871391589`.
- Zoom-out radius shrink against zoom-in 2: `0.816496580927726`.

## E5 - Detail Spacing Evidence

Status: completed for Phase 5 explicit focus camera behavior.

Implemented detail/focus behavior:

- Search result focus and parent `GraphCanvasHandle.focusNode` now call the real `useSigma.focusNode` camera path.
- Focusing an already selected node still animates camera back to detail mode.
- Detail focus ratio is derived from current rendered node radius through the scale model, targeting `8px` visible node radius.
- Focus camera target uses framed camera coordinates derived from `sigma.graphToViewport` and `sigma.viewportToFramedGraph`; raw graph coordinates are not used as camera coordinates.
- E2E asserts that explicit focus leaves visible viewport node count above `0`.

Measured focus sequence:

| Stage | Camera ratio | Max rendered radius px | Visible viewport nodes | Visible islands |
|---|---:|---:|---:|---:|
| Search focus | 0.140625 | 8 | 787 | 1 |
| Same selection shifted by zoom-out | 0.2109375 | 6.531972647421808 | 1397 | 1 |
| Same selection refocus | 0.140625 | 8 | 787 | 1 |

Focus result:

- Search focus radius growth against prior zoom-out: `2.1773242158072694`.
- Same-selected refocus radius growth after zoom-out: `1.224744871391589`.

Screenshot artifacts:

- `reports/problem/2026-05-27-graph-phase5-search-focus.png`
- `reports/problem/2026-05-27-graph-phase5-same-selected-refocus.png`
- `reports/problem/2026-05-27-graph-phase5-camera-zoom-metrics.json`

## E6 - Validation Evidence

Status: in progress.

Phase 5 validation:

- Full Web build before tests: `npm run build` passed.
- Focused Web unit tests: `npx vitest run test/unit/graph-scale-model.test.ts test/unit/graph-camera-mode.test.ts test/unit/graph-screen-spacing.test.ts test/unit/graph-adapter.edge-geometry.test.ts` passed, `4` files and `38` tests.
- Web e2e/browser tests: `npx playwright test e2e/graph-orientation-labels.spec.ts --project=chromium` passed, `3` tests.
- Browser screenshot validation inspected Phase 5 overview, search focus, and same-selected refocus artifacts.
- First e2e attempt exposed that search selected a node without calling `useSigma.focusNode`; `GraphCanvasHandle.focusNode` now calls the camera focus path.
- Benchmark capture exposed that raw graph coordinates could focus an empty viewport; `useSigma.focusNode` now converts target node coordinates into framed camera coordinates.

## E7 - Pre-Commit Change Detection

Status: in progress.

Phase 5 pre-commit detection:

- `git diff --check`: passed.
- `avmatrix analyze --force`: scanned `789`, parsed `578`, unsupported `211`, failed `0`.
- Refreshed graph: `90253` nodes, `123573` relationships.
- `avmatrix detect-changes --repo AVmatrix --scope all`: risk level `low`.
- Changed files: `8`.
- Changed symbols: `155`.
- Changed app layers: docs `11`, frontend `61`, frontend_test `83`.
- Changed functional areas: documentation `11`, layout `34`, unknown `101`, web_graph_ui `9`.
- Affected count: `0`.
- Affected processes: `0`.
- Resolution gap changes: `119` analyzer bookkeeping entries.
- Changed source nodes with gaps: `0`.
- Resolution health impact: degraded nodes `0`, total resolution gap count `0`.
