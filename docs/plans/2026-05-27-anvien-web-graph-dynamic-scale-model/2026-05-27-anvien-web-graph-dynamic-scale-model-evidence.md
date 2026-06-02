# Anvien Web Graph Dynamic Scale Model And Zoom Semantics Evidence Ledger

Date: 2026-05-27

Status: Completed

Companion files:

- Plan: [2026-05-27-anvien-web-graph-dynamic-scale-model-plan.md](2026-05-27-anvien-web-graph-dynamic-scale-model-plan.md)
- Benchmark ledger: [2026-05-27-anvien-web-graph-dynamic-scale-model-benchmark.md](2026-05-27-anvien-web-graph-dynamic-scale-model-benchmark.md)

## Evidence Rules

Evidence must be recorded as soon as each evidenced task is completed.

This ledger is for traceability, implementation decisions, command outputs, screenshots, blast-radius summaries, and validation results.

Benchmark values belong in the benchmark ledger. Do not mix benchmark tables into this evidence ledger except by linking to their benchmark IDs.

## E0 - Plan Creation Context

Date: 2026-05-27

Status: recorded.

Rule correction:

```text
Only skip Anvien when running git commit for documentation-only staged changes. All docs technical planning, evidence, benchmark, report, and architecture work must use Anvien and read the codebase.
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

Anvien refresh:

```text
.\anvien\bin\anvien.exe analyze --force
files: scanned=774 parsed=572 unsupported=202 failed=0
graph: nodes=89484 relationships=122631 path=E:\Anvien\.anvien\graph.json
```

Anvien commands used:

```text
.\anvien\bin\anvien.exe query "web graph sigma camera zoom node size color overview" --repo Anvien
.\anvien\bin\anvien.exe context "useSigma" --repo Anvien
.\anvien\bin\anvien.exe context "knowledgeGraphToGraphology" --repo Anvien
.\anvien\bin\anvien.exe context "applyReadableGraphCamera" --repo Anvien
```

Key Anvien findings:

- `useSigma` is in `anvien-web/src/hooks/useSigma.ts` with incoming use from `GraphCanvas.tsx`.
- `applyReadableGraphCamera` is in `anvien-web/src/lib/graph-readable-camera.ts` with incoming call from `useSigma.ts`.
- `knowledgeGraphToGraphology` is in `anvien-web/src/lib/graph-adapter.ts`; tests in `graph-adapter.edge-geometry.test.ts` exercise it heavily.
- `buildLayoutNodeSpacingDiagnostics` is in `GraphCanvas.tsx` and participates in Web graph UI diagnostics.

Code findings:

- `anvien-web/src/hooks/useSigma.ts:257` sets `minCameraRatio: MIN_READABLE_CAMERA_RATIO`.
- `anvien-web/src/hooks/useSigma.ts:259` sets `itemSizesReference: 'positions'`.
- `anvien-web/src/hooks/useSigma.ts:518-520` sets camera state to `{ x: 0.5, y: 0.5, ratio: 1, angle: 0 }`, refreshes, then calls `applyReadableGraphCamera(sigma)`.
- `anvien-web/src/lib/graph-readable-camera.ts:53-75` picks the densest visible island.
- `anvien-web/src/lib/graph-readable-camera.ts:149-159` focuses that island and computes a readable ratio.
- `anvien-web/src/lib/graph-adapter.ts:86-107` uses fixed rendered size and fixed minimum center distance.
- `anvien-web/src/lib/graph-adapter.ts:590-593` derives island gap from fixed multiples.
- `anvien-web/src/lib/graph-adapter.ts:811-848` assigns node color from `getNodeColor(displayNodeType)`, so the palette remains type-based.
- `anvien-web/src/lib/graph-screen-spacing.ts:90-93` uses `sigma.scaleSize` for viewport diagnostics.
- `anvien-web/src/lib/runtime-diagnostics.ts` publishes graph conversion, visual scale, layout rings, layout spacing, screen spacing, readable camera, heartbeat, and reconnect diagnostics. It lacks overview visible color count, visible island count, dominant island share, visible node-type inventory, and filter node-type inventory.
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

Status: completed.

Completed pre-edit evidence on 2026-05-27:

- Anvien refresh command: `.\anvien\bin\anvien.exe analyze --force`
- Anvien refresh output: `files: scanned=774 parsed=572 unsupported=202 failed=0`, `graph: nodes=89489 relationships=122636 path=E:\Anvien\.anvien\graph.json`
- Anvien context/query owners: `useSigma`, `applyReadableGraphCamera`, `knowledgeGraphToGraphology`, `buildLayoutNodeSpacingDiagnostics`, `runtime-diagnostics.ts`
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

- Anvien graph refresh command and output;
- Anvien context for graph layout, Sigma rendering, camera, and diagnostics owners;
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

- `anvien-web/src/hooks/useSigma.ts`
- `anvien-web/src/lib/graph-adapter.ts`
- `anvien-web/src/lib/graph-readable-camera.ts`
- `anvien-web/src/lib/graph-screen-spacing.ts`
- `anvien-web/src/lib/runtime-diagnostics.ts`
- `anvien-web/src/components/GraphCanvas.tsx`
- `anvien-web/e2e/graph-orientation-labels.spec.ts`
- related unit tests under `anvien-web/test/unit`

## E1A - Color Overview Parity Gate

Status: completed.

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

- `anvien-web/src/hooks/useSigma.ts` default `setGraph` path now uses `sigma.getCamera().animatedReset({ duration: 500 })`.
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

Anvien pre-commit detection:

- `anvien analyze --force` after Phase 1 edits: scanned `779`, parsed `574`, unsupported `205`, failed `0`.
- Refreshed graph: `89896` nodes, `123159` relationships.
- `anvien detect-changes --repo Anvien --scope all`: changed files `8`, changed symbols `372`.
- Changed app layers: docs `13`, frontend `61`, frontend_test `298`.
- Changed functional areas: documentation `13`, layout `9`, unknown `344`, web_graph_ui `6`.
- Affected processes: `0`.
- Resolution health degraded nodes: `0`.
- Reported risk level: `low`.

This gate was completed before dynamic scale, zoom, and spacing implementation proceeded.

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

Anvien refresh:

- `anvien analyze --force`: scanned `779`, parsed `574`, unsupported `205`, failed `0`.
- Refreshed graph: `89896` nodes, `123159` relationships.

Anvien owner trace:

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

- Added `anvien-web/src/lib/graph-scale-model.ts`.
- Added `anvien-web/test/unit/graph-scale-model.test.ts`.
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

Anvien pre-commit detection for Phase 3:

- `anvien analyze --force`: scanned `786`, parsed `576`, unsupported `210`, failed `0`.
- Refreshed graph: `90067` nodes, `123382` relationships.
- `anvien detect-changes --repo Anvien --scope all`: risk level `low`.
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

- Updated `anvien-web/src/lib/graph-adapter.ts`.
- Updated `anvien-web/test/unit/graph-adapter.edge-geometry.test.ts`.

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

- `anvien-web/test-results/graph-orientation-labels-G-5b603-labels-on-the-desktop-graph-chromium/graph-orientation-labels-desktop.png`
- `anvien-web/test-results/graph-orientation-labels-G-e5a0a-t-and-updates-after-filters-chromium/graph-orientation-labels-small-filtered.png`
- `anvien-web/test-results/graph-orientation-labels-G-537e2-ory-visible-on-default-load-chromium/graph-node-spacing-dense-desktop.png`
- `anvien-web/test-results/graph-orientation-labels-G-537e2-ory-visible-on-default-load-chromium/graph-node-spacing-dense-small.png`

Anvien pre-commit detection for Phase 4:

- `anvien analyze --force`: scanned `786`, parsed `576`, unsupported `210`, failed `0`.
- Refreshed graph: `90038` nodes, `123327` relationships.
- `anvien detect-changes --repo Anvien --scope all`: risk level `low`.
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

Anvien refresh before Phase 5 edits:

- `anvien analyze --force`: scanned `786`, parsed `576`, unsupported `210`, failed `0`.
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
- `anvien-web/test-results/graph-orientation-labels-G-537e2-ory-visible-on-default-load-chromium/graph-node-spacing-dense-desktop.png`
- `anvien-web/test-results/graph-orientation-labels-G-537e2-ory-visible-on-default-load-chromium/graph-node-spacing-dense-small.png`

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

Status: completed.

Phase 6 diagnostics sampling:

- Added `graphInteraction` diagnostics with current mode, current target node, overview samples, zoom samples, detail/focus samples, and dynamic-gap samples.
- Added bounded sample retention at `12` samples per sample list.
- Added camera mode recording in `useSigma` for overview load, zoom-in, zoom-out, reset, and explicit detail focus.
- Added runtime unit coverage proving overview samples, zoom samples, detail/focus samples, visible island count derivation, target node tracking, and bounded sample retention.
- Inventory counts remain runtime-derived. The dense fixture still defines its own exact visible ring, island, node-type, color, and visible-node expectations.

Phase 6 focused validation:

- Full Web build before tests: `npm run build` passed.
- Focused Web unit test: `npx vitest run test/unit/runtime-diagnostics.test.ts` passed, `1` file and `9` tests.
- Full Web unit test: `npx vitest run test/unit` passed, `50` files and `397` tests.
- First Phase 6 e2e run exposed one incorrect assertion: it checked the display text `Function0` against a graph node id. The assertion now derives the expected target id from the dense fixture.
- Phase 6 e2e hardening added settled-state polling for detail/focus target-gap violations before diagnostics capture.
- Full Web build after e2e hardening: `npm run build` passed.
- Final graph e2e/browser run: `npx playwright test e2e/graph-orientation-labels.spec.ts --project=chromium` passed, `3` tests.
- Browser screenshot validation inspected `graph-node-spacing-dense-desktop.png` and `graph-node-spacing-dense-small.png`; nodes, ring labels, island labels, and multi-cluster overview remained visible.
- Browser plugin tool discovery did not expose `node_repl`; local Playwright screenshots and image inspection were used for the browser validation step.

Phase 6 e2e coverage:

- Overview assertions compare visible ring, island, node-type, color, and node count against inventory derived from the active dense fixture.
- Zoom assertions require sampled zoom-in and zoom-out modes plus rendered radius growth and shrink.
- Detail/focus assertions require `overlapCount = 0`, `targetGapViolationCount = 0`, and dynamic required center distance above rendered diameter after settled focus.
- Dynamic-gap samples are recorded from runtime screen-spacing diagnostics and capped to bounded sample history.

Phase 5 validation:

- Full Web build before tests: `npm run build` passed.
- Focused Web unit tests: `npx vitest run test/unit/graph-scale-model.test.ts test/unit/graph-camera-mode.test.ts test/unit/graph-screen-spacing.test.ts test/unit/graph-adapter.edge-geometry.test.ts` passed, `4` files and `38` tests.
- Web e2e/browser tests: `npx playwright test e2e/graph-orientation-labels.spec.ts --project=chromium` passed, `3` tests.
- Browser screenshot validation inspected Phase 5 overview, search focus, and same-selected refocus artifacts.
- First e2e attempt exposed that search selected a node without calling `useSigma.focusNode`; `GraphCanvasHandle.focusNode` now calls the camera focus path.
- Benchmark capture exposed that raw graph coordinates could focus an empty viewport; `useSigma.focusNode` now converts target node coordinates into framed camera coordinates.

## E7 - Pre-Commit Change Detection

Status: completed.

Phase 5 pre-commit detection:

- `git diff --check`: passed.
- `anvien analyze --force`: scanned `789`, parsed `578`, unsupported `211`, failed `0`.
- Refreshed graph: `90253` nodes, `123573` relationships.
- `anvien detect-changes --repo Anvien --scope all`: risk level `low`.
- Changed files: `8`.
- Changed symbols: `155`.
- Changed app layers: docs `11`, frontend `61`, frontend_test `83`.
- Changed functional areas: documentation `11`, layout `34`, unknown `101`, web_graph_ui `9`.
- Affected count: `0`.
- Affected processes: `0`.
- Resolution gap changes: `119` analyzer bookkeeping entries.
- Changed source nodes with gaps: `0`.
- Resolution health impact: degraded nodes `0`, total resolution gap count `0`.

Phase 6 pre-commit detection:

- `git diff --check`: passed.
- Forbidden wording scan across the three plan ledgers: no matches.
- `anvien analyze --force`: scanned `789`, parsed `578`, unsupported `211`, failed `0`.
- Refreshed graph: `90536` nodes, `123974` relationships.
- `detect_changes` through Anvien MCP with `repo=Anvien`, `scope=all`: risk level `low`.
- Changed files: `7`.
- Changed symbols: `358`.
- Changed app layers: docs `8`, frontend `109`, frontend_test `241`.
- Changed functional areas: documentation `8`, layout `16`, unknown `334`.
- Affected count: `0`.
- Affected processes: `0`.
- Resolution gap changes: `255` analyzer bookkeeping entries.
- Changed source nodes with gaps: `0`.
- Resolution health impact: degraded nodes `0`, total resolution gap count `0`.

## E8 - Final Closure Evidence

Status: completed for Phase 7 closure.

Phase 7 validation:

- Full Web build gate: `npm run build` passed.
- Focused Web unit tests: `npx vitest run test/unit/graph-scale-model.test.ts test/unit/graph-camera-mode.test.ts test/unit/graph-screen-spacing.test.ts test/unit/graph-adapter.edge-geometry.test.ts test/unit/runtime-diagnostics.test.ts` passed, `5` files and `47` tests.
- Full Web unit tests: `npx vitest run test/unit` passed, `50` files and `397` tests.
- Web e2e/browser tests: `npx playwright test e2e/graph-orientation-labels.spec.ts --project=chromium` passed, `3` tests.
- Browser screenshots inspected: dense desktop and dense small viewport both showed nodes, labels, and multi-cluster overview.
- Final `git diff --check`: passed.
- Final forbidden wording scan across the three plan ledgers: no matches.
- Final `anvien analyze --force`: scanned `789`, parsed `578`, unsupported `211`, failed `0`; graph `90536` nodes and `123974` relationships.
- Final Anvien MCP `detect_changes` with `repo=Anvien`, `scope=all`: risk level `low`, changed files `1`, changed symbols `2`, changed app layers docs `2`, affected processes `0`, resolution health degraded nodes `0`, total resolution gap count `0`.
- Dev server on `127.0.0.1:5228` was stopped after e2e validation.

## E9 - Phase 8 Reopened Performance Hardening Evidence

Status: in progress.

P8-A Anvien refresh:

- Date: 2026-05-28.
- Command: `anvien analyze --force`.
- Result: passed.
- Files: scanned `789`, parsed `578`, unsupported `211`, failed `0`.
- Refreshed graph: `90541` nodes, `123979` relationships.

P8-B Anvien owner trace:

- `anvien query "GraphCanvas camera updated screen spacing diagnostics graph overview orientation labels wheel zoom runtime diagnostics launcher build" --repo Anvien --limit 10`: traced the active Web graph owner to `anvien-web/src/components/GraphCanvas.tsx`, screen spacing diagnostics, overview diagnostics, orientation labels, `useSigma`, runtime diagnostics, and launcher-related results.
- `anvien context "GraphCanvas" --repo Anvien`: `GraphCanvas` is consumed by `App.tsx` and `GraphCanvas.selection-performance.test.tsx`.
- `anvien context --uid "Function:anvien-web/src/lib/graph-screen-spacing.ts:buildScreenNodeSpacingDiagnostics" --repo Anvien`: incoming calls from `GraphCanvas.tsx`, `graph-readable-camera.ts`, and unit coverage.
- `anvien context --uid "Function:anvien-web/src/lib/graph-overview-diagnostics.ts:buildGraphOverviewDiagnostics" --repo Anvien`: incoming calls from `GraphCanvas.tsx` and unit coverage.
- `anvien context --uid "Function:anvien-web/src/lib/graph-orientation-labels.ts:buildGraphOrientationLabels" --repo Anvien`: incoming calls from `GraphCanvas.tsx` and unit coverage.
- `anvien context --uid "Function:anvien-web/src/lib/graph-orientation-labels.ts:placeGraphOrientationLabels" --repo Anvien`: incoming calls from `GraphCanvas.tsx` and unit coverage.
- `anvien context "recordGraphInteractionMode" --repo Anvien`: owner is `anvien-web/src/lib/runtime-diagnostics.ts`; calls come from `useSigma.ts` and runtime diagnostics tests.
- `anvien context "useSigma" --repo Anvien`: owner is `anvien-web/src/hooks/useSigma.ts`; call site is `GraphCanvas.tsx`.
- `anvien query "launcher build web-dist AnvienLauncher.exe build.ps1 copy web dist" --repo Anvien --limit 10`: traced `anvien-launcher/build.ps1`, `anvien-launcher/src/main.go`, and launcher runtime processes.
- `anvien query "serve launcher web-dist embedded static files AnvienLauncher" --repo Anvien --limit 10`: traced launcher static serving through `anvien-launcher/src/main.go:serveStaticFile` and launcher runtime processes.
- `anvien context "build.ps1" --repo Anvien`: indexed file owner is `anvien-launcher/build.ps1` with app layer `cli_launcher` and functional area `launcher`.

P8-C/P8-D impact and blast radius:

| Target | Risk | Impacted count | Affected processes | Direct owner evidence |
|---|---:|---:|---:|---|
| `GraphCanvas` | LOW | 2 | 0 | `App.tsx`, `main.tsx` |
| `buildScreenNodeSpacingDiagnostics` | CRITICAL | 7 | 11 | `GraphCanvas.tsx`, `graph-readable-camera.ts`, unit tests |
| `buildGraphOverviewDiagnostics` | CRITICAL | 4 | 11 | `GraphCanvas.tsx`, unit tests |
| `buildGraphOrientationLabels` | CRITICAL | 4 | 11 | `GraphCanvas.tsx`, unit tests |
| `placeGraphOrientationLabels` | CRITICAL | 4 | 11 | `GraphCanvas.tsx`, unit tests |
| `useSigma` | CRITICAL | 4 | 11 | `GraphCanvas.tsx`, `App.tsx`, `main.tsx` |
| `recordScreenNodeSpacing` | CRITICAL | 4 | 11 | `GraphCanvas.tsx`, `App.tsx`, `main.tsx` |
| `recordGraphOverview` | CRITICAL | 4 | 11 | `GraphCanvas.tsx`, `App.tsx`, `main.tsx` |
| `recordGraphInteractionMode` | LOW | 4 | 0 | `useSigma.ts`, `GraphCanvas.tsx`, `App.tsx` |

Blast radius note:

- The CRITICAL results are workflow warnings for Web graph render and diagnostics paths. They do not block edits; they require scoped changes, preservation checks for colors/types/labels/edges, and full build plus e2e validation.

P8-E baseline reproduction:

- Current Web build command before baseline: `npm run build` in `anvien-web`; result passed. Bundle artifact includes `assets/index-Bxk3Ac59.js`.
- Current build preview: served `anvien-web/dist` on `http://127.0.0.1:5231`; stopped after capture.
- Packaged launcher baseline: started `anvien-launcher\AnvienLauncher.exe start` with `ANVIEN_LAUNCHER_NO_BROWSER=1`, captured `http://127.0.0.1:5228`, then stopped with `anvien-launcher\AnvienLauncher.exe stop`.
- Baseline JSON artifacts:
  - `reports/problem/2026-05-28-phase8-current-build-baseline.json`.
  - `reports/problem/2026-05-28-phase8-packaged-launcher-baseline.json`.
- Baseline screenshot artifacts:
  - `reports/problem/2026-05-28-phase8-current-build-baseline-load.png`.
  - `reports/problem/2026-05-28-phase8-current-build-baseline-wheel.png`.
  - `reports/problem/2026-05-28-phase8-packaged-launcher-baseline-load.png`.
- Current build baseline confirmed the regression remains before edits: wheel input changes camera ratio from `1` to `0.5882352941176471`, but `graphInteraction.currentMode` stays `overview` and `zoomSamples` stays `0`.
- Current build load retained graph inventory: `visibleColorCount=3`, `visibleRingCount=3`, `visibleIslandCount=4`, `graphRingInventory=[backend, docs, frontend]`, `graphIslandInventory=[backend:Function, docs:Documentation, frontend:Function, frontend:Method]`.
- Packaged launcher baseline confirmed the serving bug remains before edits: `screenNodeSpacing` exists, but `graphOverview` and `graphInteraction` are absent from the served bundle; wheel input did not change the stale launcher camera ratio (`0.004001487892151231` before and after wheel).

P8-O/P8-P/P8-Q full build gate:

- Command: `powershell -ExecutionPolicy Bypass -File anvien-launcher\build.ps1`.
- Result: passed.
- Web build bundle: `anvien-web\dist\assets\index-cazHqvFt.js`, length `2053493`, SHA256 `CD837CA78F245717A6316A8DEB95D0309D3B9A61FEE2B5B806A715AECCDF7AEA`.
- Launcher web bundle: `anvien-launcher\web-dist\assets\index-cazHqvFt.js`, length `2053493`, SHA256 `CD837CA78F245717A6316A8DEB95D0309D3B9A61FEE2B5B806A715AECCDF7AEA`, copied match `true`.
- Launcher executable: `anvien-launcher\AnvienLauncher.exe`, LastWriteTime `2026-05-28 09:45:29 +07:00`, size `6993408`, SHA256 `1802C7DD2EC4114E2B123895205E08FA51E4D0A06827192DA089D2919ABD0945`.
- Canonical CLI executable: `anvien\bin\anvien.exe`, LastWriteTime `2026-05-28 09:45:28 +07:00`, size `50528768`, SHA256 `AA8AA6F57663AA4793C63A078C079D6A6A330A747F441CCB3F5F9C53BC48DC7C`.
- Launcher server executable: `anvien-launcher\server-bundle\anvien-server.exe`, LastWriteTime `2026-05-28 09:45:30 +07:00`, size `2053632`, SHA256 `1C61675647F748460DFED28BEAC2896F791EFD34F1E21215C1E522F7015A9103`.

P8-F through P8-M implementation evidence:

- Added reusable browser performance probe helpers in `anvien-web\e2e\graph-orientation-labels.spec.ts` for long tasks, max frame delta, frame drops, and diagnostics-write counts across dense load, wheel zoom, button zoom, and camera settle.
- Added `wheel-zoom` to `GraphInteractionMode` and routed `wheel-zoom` samples into `graphInteraction.zoomSamples`.
- Added `graphCamera` diagnostics for cheap camera-only samples without running full screen spacing or overview audits.
- Coalesced `graphCamera` writes through one pending RAF so repeated camera updates in the same frame keep only the latest camera sample.
- Added capture-phase wheel listener in `useSigma` so mouse wheel zoom records `wheel-zoom` before Sigma consumes the event path.
- Reworked `GraphCanvas` orientation labels so source inventory is built on graph/filter changes and camera/afterRender only places cached labels.
- Reworked `GraphCanvas` camera hot path so full `buildScreenNodeSpacingDiagnostics` and `buildGraphOverviewDiagnostics` run after camera settle through idle work instead of on every camera update.
- Preserved graph/filter-triggered full diagnostics and initial load diagnostics so existing screen spacing and overview inventory evidence remains available.

P8-R launcher-served bundle verification:

- Launcher server: `http://127.0.0.1:5228`.
- Served bundle: `assets/index-cazHqvFt.js`.
- Served bundle contains `graphOverview`: `true`.
- Served bundle contains `graphInteraction`: `true`.
- Served bundle contains `wheel-zoom`: `true`.

P8-S/P8-T/P8-U validation:

- First launcher e2e attempt exposed the exact remaining wheel bug: camera/radius changed, but `wheel-zoom` was not present in `zoomSamples`. The listener was then changed from bubble-phase to capture-phase.
- Focused Web unit tests after the capture-phase fix: `npx vitest run test/unit/runtime-diagnostics.test.ts test/unit/graph-screen-spacing.test.ts test/unit/graph-orientation-labels.test.ts test/unit/graph-adapter.edge-geometry.test.ts` passed, `4` files and `45` tests.
- Full Web unit tests after the capture-phase fix: `npx vitest run test/unit` passed, `50` files and `398` tests.
- Web e2e/browser tests against launcher-served `127.0.0.1:5228`: `npx playwright test e2e/graph-orientation-labels.spec.ts --project=chromium` passed, `3` tests.
- Browser screenshot validation inspected:
  - `anvien-web\test-results\graph-orientation-labels-G-537e2-ory-visible-on-default-load-chromium\graph-node-spacing-dense-desktop.png`.
  - `anvien-web\test-results\graph-orientation-labels-G-537e2-ory-visible-on-default-load-chromium\graph-node-spacing-dense-small.png`.
  - `reports\problem\2026-05-28-phase8-launcher-final-load.png`.
  - `reports\problem\2026-05-28-phase8-launcher-final-wheel.png`.
- Screenshots show graph nodes, multiple color groups, ring labels, island labels, zoom controls, and dense overview/focused wheel views still visible.

P8-N preservation evidence:

- Final dense load inventory retained `visibleColorCount=3`, `visibleRingCount=3`, `visibleIslandCount=4`.
- Final graph ring inventory retained `[backend, docs, frontend]`.
- Final graph island inventory retained `[backend:Function, docs:Documentation, frontend:Function, frontend:Method]`.
- Final filter node type inventory retained `[Documentation, Function, Method]`.
- Final wheel zoom retained graph island inventory and recorded `graphInteraction.currentMode=wheel-zoom`, `zoomSamples=1`, and `graphCamera.mode=wheel-zoom`.
- Final button zoom retained graph island inventory and recorded `graphInteraction.currentMode=zoom-in`, `zoomSamples=2`, and `graphCamera.mode=zoom-in`.

P8-X pre-commit change detection:

- `git diff --check`: passed.
- `anvien analyze --force`: scanned `792`, parsed `578`, unsupported `214`, failed `0`.
- Refreshed graph: `90916` nodes, `124370` relationships.
- `anvien detect-changes --repo Anvien --scope all`: risk level `medium`.
- Changed files: `8`.
- Changed symbols: `412`.
- Changed app layers: docs `9`, frontend `187`, frontend_test `216`.
- Changed functional areas: documentation `9`, layout `12`, unknown `242`, web_graph_ui `149`.
- Affected count: `4`.
- Affected app layers: frontend `4`.
- Affected functional areas: layout `4`.
- Affected processes: `4`, all from the new `handleWheel` flow (`HandleWheel -> ApplyFilterBasedClusteredLayout`, `HandleWheel -> CapRenderedNodeSize`, `HandleWheel -> HexToRgb`, `HandleWheel -> RgbToHex`).
- Resolution gap changes: `319` changed gap entities and `328` changed gap occurrences; changed source nodes with gaps `0`.
- Resolution health impact: degraded nodes `0`, total resolution gap count `0`.

P8-Y/P8-Z closure:

- Implementation commit: `deb5a71 fix(web): harden graph zoom diagnostics`.
- Post-implementation `git status --short`: clean.
- Reopened Phase 8 closure condition met: current build passed through the full launcher build gate, launcher-served bundle passed e2e/browser validation, and benchmark/evidence artifacts were recorded.

## E9 - Phase 9 Plan Reopen Evidence

Status: active for P9-A through P9-E.

Correction:

- A prior response misread "split files" as documentation splitting. That was incorrect.
- The required split is codebase splitting: reduce mixed responsibilities in overloaded Web graph code before restoring ring/spiral topology.
- The temporary doc split files were removed and are not part of the Phase 9 plan.

Anvien refresh:

- Command: `anvien analyze --force`.
- Result: passed.
- Files: scanned `792`, parsed `578`, unsupported `214`, failed `0`.
- Refreshed graph: `90916` nodes, `124370` relationships.

Anvien owner discovery:

- Query: `graph-adapter overloaded layout conversion node color edge filter spiral hex spacing inventory zoom`.
- Primary owners returned:
  - `anvien-web/src/lib/graph-adapter.ts:knowledgeGraphToGraphology`.
  - `anvien-web/src/lib/graph-adapter.ts:getClusterNodeSpacing`.
  - `anvien-web/src/lib/graph-adapter.ts:applyFilterBasedClusteredLayout`.
  - `anvien-web/src/components/GraphCanvas.tsx:buildLayoutNodeSpacingDiagnostics`.
  - `anvien-web/src/hooks/useSigma.ts:useSigma`.

Codebase read:

- `anvien-web/src/lib/graph-adapter.ts` is `764` lines.
- The file currently owns exported Sigma node/edge attributes, rendered-size helpers, app-layer ring taxonomy, island classification, layout geometry, node mass, graph conversion, edge conversion, label filtering, and depth filtering.
- This mixed ownership increases the chance that a layout fix changes unrelated conversion, filtering, color, edge, or interaction behavior.

Topology regression evidence:

- Current layout owner: `applyFilterBasedClusteredLayout`, lines `426-628`.
- Current conversion owner calls layout through `knowledgeGraphToGraphology`, line `737`.
- Current same-island placement uses `getHexAxialRingCoordinates`, `getHexIslandOffset`, and `getIslandOffsets`.
- Commit `0963d96 feat(web): use dynamic hex graph layout` removed `GOLDEN_ANGLE`, `getIslandOffset`, `getFallbackIslandOffset`, `getPinwheelRadiusMultiplier`, and `getPinwheelAngleOffset`.
- The same commit added axial hex coordinate placement and changed island placement to one equal `clusterOrbitRadius`.
- Therefore the missing ring/spiral/galaxy shape is a code regression, not a build-only failure.

Impact evidence:

- `anvien context "applyFilterBasedClusteredLayout" --repo Anvien` found incoming calls from `useSigma.ts` and `graph-adapter.ts`.
- `anvien impact "applyFilterBasedClusteredLayout" --repo Anvien --direction upstream` returned risk `CRITICAL`.
- Affected app layers: frontend `6`.
- Affected functional areas: layout `4`, web_graph_ui `2`.
- Affected processes: `8`, including `HandleWheel -> ApplyFilterBasedClusteredLayout`, `HandleWheel -> CapRenderedNodeSize`, `HandleWheel -> HexToRgb`, `HandleWheel -> RgbToHex`, and `GetRenderedNodeRadius` flows.
- HIGH/CRITICAL is a warning to work carefully, not a ban on changing the layout owner.

Phase 9 planning decision:

- Phase 9 has one implementation direction with two ordered slices.
- Slice 9A is behavior-preserving code file splitting from `graph-adapter.ts` into responsibility-owned modules while keeping the existing public import surface.
- Slice 9B restores collision-safe spiral/ring topology.
- Phase 9 must preserve spacing, dynamic inventory, node types, filters, colors, community colors, zoom, performance strategy, interactions, edges, labels, full launcher build, validation, evidence, and benchmark separation.

## E10 - Phase 9A Graph Adapter Split Evidence

Status: completed for P9-F through P9-J.

P9-F code split:

- `anvien-web/src/lib/graph-adapter.ts` is now a public facade that re-exports the existing public graph adapter surface.
- Added `anvien-web/src/lib/graph-adapter-types.ts` for `SigmaNodeAttributes` and `SigmaEdgeAttributes`.
- Added `anvien-web/src/lib/graph-node-sizing.ts` for rendered node size, scaled node size, center-distance helpers, and node mass.
- Added `anvien-web/src/lib/graph-layout.ts` for app-layer ring taxonomy and `applyFilterBasedClusteredLayout`.
- Added `anvien-web/src/lib/graph-conversion.ts` for `knowledgeGraphToGraphology`.
- Added `anvien-web/src/lib/graph-filtering.ts` for label and depth filtering.
- External imports from `graph-adapter.ts` are intentionally preserved for the split slice.
- The split slice has not changed the hex layout algorithm; topology restoration remains pending P9-K through P9-N.

P9-F pre-validation status:

- `git diff --check` first found one blank line at EOF in `graph-adapter.ts`; it was removed before validation.

P9-H full build first attempt:

- Command: `powershell -ExecutionPolicy Bypass -File anvien-launcher\build.ps1`.
- Web build portion passed: TypeScript project build and Vite build completed.
- Build gate did not pass because launcher packaging could not remove `anvien-launcher\server-bundle\anvien-server.exe`.
- Error: `Remove-Item : Cannot remove item E:\Anvien\anvien-launcher\server-bundle\anvien-server.exe: Access to the path 'anvien-server.exe' is denied.`
- Runtime process trace found launcher-owned processes still holding repo build artifacts:
  - `AnvienLauncher.exe` PID `11592`, path `E:\Anvien\anvien-launcher\AnvienLauncher.exe`, parent `explorer.exe`.
  - `anvien-server.exe` PID `10352`, path `E:\Anvien\anvien-launcher\server-bundle\anvien-server.exe`, parent `AnvienLauncher.exe`.
  - `anvien.exe` PID `12388`, path `E:\Anvien\anvien\bin\anvien.exe`, command `serve --host 127.0.0.1 --port 4848`, parent `anvien-server.exe`.
- Editor-owned Anvien MCP processes were also present and were not part of the build lock.
- Decision: stop only the launcher-owned runtime group, then rerun the same full build gate. P9-H remains incomplete until the retry passes.

P9-H full build retry:

- Launcher-owned runtime group was stopped with `anvien-launcher\AnvienLauncher.exe stop`.
- `anvien doctor processes --json` then showed only editor-owned Anvien MCP processes and a VS Code Playwright test-server; the repo launcher process group was gone.
- Command: `powershell -ExecutionPolicy Bypass -File anvien-launcher\build.ps1`.
- Result: passed.
- Web build completed TypeScript project build and Vite production build.
- `anvien-web\dist` was rebuilt.
- `anvien-launcher\web-dist` was recopied from the current Web dist.
- `anvien\bin\anvien.exe` was rebuilt.
- `anvien-launcher\server-bundle\anvien-server.exe` was rebuilt.
- `anvien-launcher\AnvienLauncher.exe` was rebuilt.
- Build warnings were limited to existing Vite chunk/dynamic-import warnings; they did not fail the gate.
- Artifact hashes and sizes are recorded in benchmark row `B9`.

P9-G/P9-I behavior-preserving validation:

- Focused Web unit command: `npx vitest run test/unit/graph-adapter.edge-geometry.test.ts test/unit/graph-scale-model.test.ts test/unit/graph-screen-spacing.test.ts test/unit/graph-overview-diagnostics.test.ts test/unit/graph-orientation-labels.test.ts test/unit/graph-camera-mode.test.ts test/unit/runtime-diagnostics.test.ts`.
- Focused Web unit result: passed, `7` files and `56` tests.
- Launcher-served app was started with `ANVIEN_LAUNCHER_NO_BROWSER=1` and `anvien-launcher\AnvienLauncher.exe start`.
- Launcher process trace after start:
  - `AnvienLauncher.exe` PID `4284`, path `E:\Anvien\anvien-launcher\AnvienLauncher.exe`.
  - `anvien-server.exe` PID `10368`, path `E:\Anvien\anvien-launcher\server-bundle\anvien-server.exe`.
  - `anvien.exe` PID `6892`, path `E:\Anvien\anvien\bin\anvien.exe`, command `serve --host 127.0.0.1 --port 4848`.
- Port check: `127.0.0.1:5228` returned `TcpTestSucceeded=True`.
- Existing Web e2e/browser command: `npx playwright test e2e/graph-orientation-labels.spec.ts --project=chromium`.
- Existing Web e2e/browser result: passed, `3` tests.
- Screenshot artifacts were produced under `anvien-web\test-results` for desktop graph labels, small filtered labels, dense desktop overview, and dense small overview.
- Behavior-preserving decision: the split slice changed ownership boundaries only. The public `graph-adapter.ts` import surface remained stable, the current hex topology was intentionally left unchanged, and the focused unit plus launcher-served e2e coverage passed before topology edits began.

P9-J pre-commit Anvien detection:

- `git diff --check`: passed.
- `anvien analyze --force`: scanned `797`, parsed `583`, unsupported `214`, failed `0`.
- Refreshed graph: `91018` nodes and `124419` relationships.
- First `detect-changes --scope all` before staging saw only documentation because new split modules were still untracked; the slice was staged and detection was rerun on the staged scope.
- Staged `anvien detect-changes --repo Anvien --scope staged`: risk level `medium`.
- Changed files: `9`.
- Changed symbols: `539`.
- Changed app layers: docs `17`, frontend `522`.
- Changed functional areas: documentation `17`, unknown `522`.
- Affected count: `4`.
- Affected app layers: frontend `4`.
- Affected functional areas: unknown `4`.
- Affected processes:
  - `KnowledgeGraphToGraphology -> CompareKnownOrder`.
  - `KnowledgeGraphToGraphology -> GetDisplayGraphRelationships`.
  - `KnowledgeGraphToGraphology -> GetFileExtension`.
  - `KnowledgeGraphToGraphology -> GetFileStem`.
- Resolution gap changes: `332` changed gap entities and `332` changed gap occurrences; changed source nodes with gaps `0`.
- Resolution health impact: degraded nodes `0`, total resolution gap count `0`.
- `knowledgeGraphToGraphology` context/impact was ambiguous by name after the split, so the exact UID was used for impact.
- Exact UID impact: `Function:anvien-web/src/lib/graph-conversion.ts:knowledgeGraphToGraphology`.
- Exact UID impact result: risk `LOW`, impacted count `9`, processes affected `0`; direct impact is the `graph-adapter.ts` re-export facade and upstream imports through `GraphCanvas.tsx`, `useSigma.ts`, graph diagnostics, readable camera, selected graph context, and `App.tsx`.
- Decision: medium changed-scope risk is expected for moving existing graph conversion/layout code into new modules. The runtime behavior was already validated through the full launcher build, focused unit tests, and launcher-served e2e/browser screenshots.

## E11 - Phase 9B Spiral Topology Restoration Evidence

Status: in progress for P9-K through P9-Z.

P9-K failing topology tests:

- Anvien refresh before topology work: `anvien analyze --force`.
- Result: scanned `797`, parsed `583`, unsupported `214`, failed `0`; graph `91018` nodes and `124419` relationships.
- Exact impact target: `Function:anvien-web/src/lib/graph-layout.ts:applyFilterBasedClusteredLayout`.
- Exact impact result: risk `HIGH`, impacted count `11`, affected module `Cluster`, affected processes `4`.
- Affected processes:
  - `KnowledgeGraphToGraphology -> CompareKnownOrder`.
  - `KnowledgeGraphToGraphology -> GetDisplayGraphRelationships`.
  - `KnowledgeGraphToGraphology -> GetFileExtension`.
  - `KnowledgeGraphToGraphology -> GetFileStem`.
- HIGH is a blast-radius warning for central Web graph layout, not an edit ban.
- Added unit coverage in `anvien-web\test\unit\graph-adapter.edge-geometry.test.ts`:
  - island centers inside one app-layer ring must use radial pinwheel variation, not one equal-radius orbit;
  - same-island nodes must use deterministic collision-safe spiral positions, not hex/grid rows;
  - existing same-island spacing, edge-gap, color, ring, and island tests remain in the same file.
- Failing command before implementation: `npx vitest run test/unit/graph-adapter.edge-geometry.test.ts`.
- Expected failure result before implementation: `2` failed tests and `23` passed tests.
- Failing assertions:
  - pinwheel variation: expected center-radius ratio `1.0000000000000002` to be at least `1.08`;
  - spiral/non-grid placement: expected rounded unique x-count `27` to be greater than `122.39999999999999`.

P9-L through P9-S implementation evidence:

- Replaced the hex-only same-island placement path in `anvien-web\src\lib\graph-layout.ts` with deterministic golden-angle/radial candidate placement.
- Added a collision grid and fallback radial candidate path so same-island placement remains deterministic and collision-safe.
- Restored app-layer island radial/pinwheel variation through per-slot radius and angle offsets while keeping dynamic footprint gaps.
- Tightened Web e2e layout diagnostics in `anvien-web\e2e\graph-orientation-labels.spec.ts`:
  - layout ring inventory is derived from the active dense graph instead of a fixed count;
  - ring center/order checks preserve API between backend/frontend and docs centered behavior;
  - same-color island violations must remain zero;
  - frontend ring must keep multiple islands.
- Added focused camera spacing protection after the first topology e2e run exposed a remaining failure:
  - failing symptom: detail/search focus reached `targetGapViolationCount=1188`;
  - impacted symbols were checked with Anvien before edits;
  - detail focus ratio now uses scale-model spacing constraints so explicit search/focus can zoom deeper when screen-space spacing requires it.
- The focus-ratio fix preserves node data, filters, color, labels, diagnostics, interactions, and topology. It changes camera ratio only for explicit detail focus.

P9-L through P9-S Anvien impact evidence:

- Anvien refresh command: `anvien analyze --force`.
- Result: scanned `797`, parsed `583`, unsupported `214`, failed `0`; graph `91147` nodes and `124601` relationships.
- Exact impact `Function:anvien-web/src/lib/graph-camera-mode.ts:buildDetailFocusCameraAction`: risk `CRITICAL`, impacted count `5`, affected processes `5`.
- Exact impact `Function:anvien-web/src/hooks/useSigma.ts:useSigma`: risk `CRITICAL`, impacted count `4`, affected processes `11`.
- Exact impact `Function:anvien-web/src/lib/graph-scale-model.ts:getDetailFocusCameraRatio`: risk `CRITICAL`, impacted count `6`, affected processes `5`.
- Exact impact `Function:anvien-web/src/lib/graph-scale-model.ts:buildGraphScaleModel`: risk `LOW`, impacted count `0`.
- CRITICAL is a blast-radius warning for central graph UI camera/interaction paths, not an edit ban.

P9-T/P9-U full build gate:

- Stopped the launcher-owned runtime group with `anvien-launcher\AnvienLauncher.exe stop` before rebuilding artifacts.
- Command: `powershell -ExecutionPolicy Bypass -File anvien-launcher\build.ps1`.
- Result: passed.
- The build rebuilt `anvien-web\dist`, recopied `anvien-launcher\web-dist`, rebuilt `anvien\bin\anvien.exe`, rebuilt `anvien-launcher\server-bundle\anvien-server.exe`, and rebuilt `anvien-launcher\AnvienLauncher.exe`.
- Build warnings were limited to existing Vite chunk/dynamic-import warnings; they did not fail the gate.
- Launcher executable and packaged bundle metrics are recorded in benchmark row `B10`.

P9-V validation:

- Preliminary focused unit command after the camera-ratio fix: `npx vitest run test/unit/graph-scale-model.test.ts test/unit/graph-camera-mode.test.ts test/unit/graph-adapter.edge-geometry.test.ts`.
- Preliminary focused unit result: passed, `3` files and `38` tests.
- Official focused Web unit command after full build: `npx vitest run test/unit/graph-adapter.edge-geometry.test.ts test/unit/graph-scale-model.test.ts test/unit/graph-screen-spacing.test.ts test/unit/graph-overview-diagnostics.test.ts test/unit/graph-orientation-labels.test.ts test/unit/graph-camera-mode.test.ts test/unit/runtime-diagnostics.test.ts`.
- Official focused Web unit result: passed, `7` files and `59` tests.
- Full Web unit command: `npx vitest run test/unit`.
- Full Web unit result: passed, `50` files and `401` tests.
- Launcher-served app start: `ANVIEN_LAUNCHER_NO_BROWSER=1` and `anvien-launcher\AnvienLauncher.exe start`.
- Port check: `127.0.0.1:5228` returned `TcpTestSucceeded=True`.
- Web e2e/browser command: `npx playwright test e2e/graph-orientation-labels.spec.ts --project=chromium`.
- Web e2e/browser result: passed, `3` tests.
- Screenshot artifacts:
  - `anvien-web\test-results\graph-orientation-labels-G-5b603-labels-on-the-desktop-graph-chromium\graph-orientation-labels-desktop.png`.
  - `anvien-web\test-results\graph-orientation-labels-G-e5a0a-t-and-updates-after-filters-chromium\graph-orientation-labels-small-filtered.png`.
  - `anvien-web\test-results\graph-orientation-labels-G-537e2-ory-visible-on-default-load-chromium\graph-node-spacing-dense-desktop.png`.
  - `anvien-web\test-results\graph-orientation-labels-G-537e2-ory-visible-on-default-load-chromium\graph-node-spacing-dense-small.png`.

P9-W evidence/benchmark update:

- Product/runtime metrics are recorded in benchmark row `B10`.
- This evidence row records traceability, command results, implementation decisions, impact checks, and validation outcomes only.

P9-X pre-commit change detection:

- `git diff --check`: passed.
- `anvien analyze --force`: scanned `797`, parsed `583`, unsupported `214`, failed `0`.
- Refreshed graph: `91187` nodes and `124660` relationships.
- `anvien detect-changes --repo Anvien --scope all`: risk level `medium`.
- Changed files: `11`.
- Changed symbols: `332`.
- Changed app layers: docs `8`, frontend `149`, frontend_test `175`.
- Changed functional areas: documentation `8`, layout `27`, unknown `297`.
- Affected count: `2`.
- Affected app layers: frontend `2`.
- Affected functional areas: layout `2`.
- Affected processes:
  - `HandleWheel -> GetFocusIslandNodeSizes`.
  - `HandleWheel -> GetNodeIslandKey`.
- Resolution gap changes: `236` changed gap entities and `254` changed gap occurrences; changed source nodes with gaps `0`.
- Resolution health impact: degraded nodes `0`, total resolution gap count `0`.
- Decision: medium changed-scope risk is expected because the slice changes graph layout topology, focus-camera spacing, and topology/e2e assertions. Full build, focused unit, full unit, launcher-served e2e, and screenshot validation passed before this detection.

P9-Y/P9-Z closure:

- Implementation commit: `955be36 fix(web): restore spiral graph topology`.
- Post-commit `git status --short`: clean.
- Phase 9 closure condition met: graph adapter split is committed, spiral/pinwheel topology is restored, rendered-size spacing is preserved, color/inventory and node-type preservation e2e checks pass, zoom/focus interactions pass, full launcher build artifacts were rebuilt, evidence and benchmark ledgers are separated and updated.
