# AVmatrix Web Graph Dynamic Scale Model And Zoom Semantics Plan

Date: 2026-05-27
Reopened: 2026-05-28

Status: Reopened - Phase 8 pending

Companion files:

- Evidence ledger: [2026-05-27-avmatrix-web-graph-dynamic-scale-model-evidence.md](2026-05-27-avmatrix-web-graph-dynamic-scale-model-evidence.md)
- Benchmark ledger: [2026-05-27-avmatrix-web-graph-dynamic-scale-model-benchmark.md](2026-05-27-avmatrix-web-graph-dynamic-scale-model-benchmark.md)

## Master Rules

1. Use AVmatrix for codebase analysis and impact checks while working on implementation slices in this plan.
2. As each task is completed, update the corresponding checklist item immediately.
3. Run the full build gate before testing. For this plan, full build means `powershell -ExecutionPolicy Bypass -File avmatrix-launcher\build.ps1`, and it must rebuild `avmatrix-launcher\AVmatrixLauncher.exe`.
4. Record benchmark results as each benchmarkable task is completed. Benchmarkable means measured product/runtime performance, capacity, package/startup size, graph/DB throughput, graph inventory counts, graph layout geometry metrics, rendered node radius metrics, zoom growth metrics, visible color/island counts, and browser screenshot artifact inventories. Build/test/e2e timings are validation evidence for this plan.
5. Record evidence as each evidenced task is completed.
6. Only skip AVmatrix when running `git commit` for documentation-only staged changes. All docs technical planning, evidence, benchmark, report, and architecture work must use AVmatrix and read the codebase.
7. After each completed implementation slice, commit the work, then continue until the full plan is complete.
8. HIGH and CRITICAL blast radius means the code area is important and must be changed carefully when required; it is not a ban on edits.

## Problem

The Web graph spacing fix solved a narrow overlap symptom while breaking broader graph behavior.

The current failure is not only dense-node overlap. The deeper failure is that graph geometry is not governed by one dynamic scale model. Node size, layout distance, island radius, ring radius, camera ratio, zoom behavior, edge thickness, labels, diagnostics, and tests can disagree because several modules still use fixed constants while Sigma render behavior uses camera-dependent math.

This produces visible regressions:

- default graph load loses the previous multi-color overview and focuses a dense local island;
- visible nodes appear as one color because the camera is forced into a single node-type island;
- zoom fails the expected user model where zooming in makes nodes visibly larger;
- spacing tests pass while product UX regresses because they protect one metric instead of the full graph interaction model.

The root product rule is:

```text
When one graph dimension changes, all dependent dimensions must move with it. No fixed-size assumption is allowed to silently control layout, render, camera, diagnostics, and tests.
```

## Reopened Failure - 2026-05-28

The previous closure is invalid because the completed implementation still creates product-visible lag and incomplete zoom semantics.

The current graph render path is still doing full-graph work in camera interaction hot paths. This causes the UI to freeze after graph load and during wheel zoom even when total machine CPU and RAM look low. Whole-machine CPU is not the deciding metric here; the browser tab main thread is blocked by long synchronous tasks.

Measured current-build dense fixture evidence:

- Fixture size: `2157` nodes and `2156` relationships.
- Current build preview load: `readyMs=1690`.
- Load long tasks: `3`, total `979ms`, longest `531ms`, max frame delta `567ms`.
- Wheel zoom changes camera ratio from `1` to `0.5882352941176471`, so wheel input reaches Sigma.
- One wheel burst writes `screenNodeSpacing` `6` times and `graphOverview` `6` times.
- Wheel long tasks: `4`, total `907ms`, longest `728ms`, max frame delta `717ms`.
- Button zoom changes camera ratio from `0.5882352941176471` to `0.39215686274509803`, writes diagnostics `7` times, and records `zoom-in` samples.
- Wheel zoom remains recorded as `overview`, so interaction telemetry does not identify wheel zoom.

The current diagnostics also expose a correctness gap:

- `layoutNodeSpacing` reports `overlapCount=0` and `targetGapViolationCount=0`.
- `screenNodeSpacing` on default overview reports `overlapCount=5976` and `targetGapViolationCount=17443`.

This means graph-coordinate layout validation can pass while screen-space product behavior still fails. The plan must protect viewport behavior directly, not only graph-layout geometry.

Packaging freshness is also part of the failure surface:

- The currently served launcher bundle at `avmatrix-launcher\web-dist\assets\index-B3tJBsjs.js` does not contain `graphOverview` or `graphInteraction`.
- The current Web build bundle at `avmatrix-web\dist\assets\index-Bxk3Ac59.js` does contain `graphOverview` and `graphInteraction`.
- User-facing launcher validation must therefore verify that packaged `web-dist` is rebuilt and serving the same behavior as the current Web build.
- The plan's full build gate must rebuild the launcher executable at `avmatrix-launcher\AVmatrixLauncher.exe`, not only rebuild Web assets.

## Codebase Findings

AVmatrix graph refresh was run on 2026-05-27 before this planning update:

```text
files: scanned=774 parsed=572 unsupported=202 failed=0
graph: nodes=89484 relationships=122631 path=E:\AVmatrix-GO\.avmatrix\graph.json
```

AVmatrix graph refresh was run again on 2026-05-28 before reopening this plan:

```text
files: scanned=789 parsed=578 unsupported=211 failed=0
graph: nodes=90537 relationships=123975 path=E:\AVmatrix-GO\.avmatrix\graph.json
```

AVmatrix query/context identified the current owners:

- `avmatrix-web/src/hooks/useSigma.ts`: Sigma setup, default graph load, camera reset.
- `avmatrix-web/src/lib/graph-readable-camera.ts`: readable camera selection and application.
- `avmatrix-web/src/lib/graph-adapter.ts`: graphology conversion, node color, node size, layout spacing, island geometry.
- `avmatrix-web/src/lib/graph-screen-spacing.ts`: viewport spacing diagnostics using `sigma.scaleSize`.
- `avmatrix-web/src/components/GraphCanvas.tsx`: diagnostics publication and Web graph integration.
- `avmatrix-web/src/lib/runtime-diagnostics.ts`: browser diagnostics object for e2e-visible runtime metrics.

Exact current facts:

- `useSigma.ts` sets `minCameraRatio: MIN_READABLE_CAMERA_RATIO`.
- `useSigma.ts` sets `itemSizesReference: 'positions'`.
- `useSigma.ts` loads a new graph, sets camera state to `{ x: 0.5, y: 0.5, ratio: 1, angle: 0 }`, refreshes, then calls `applyReadableGraphCamera(sigma)`.
- `graph-readable-camera.ts` selects `getDensestVisibleIsland` and applies the camera to that island.
- `graph-adapter.ts` assigns node color from `getNodeColor(displayNodeType)`. The color palette is still present; the color loss is caused by viewport collapse to one node-type island.
- `graph-adapter.ts` defines fixed rendered-size assumptions: `MAX_RENDERED_NODE_SIZE = 3`, `MAX_DENSE_RENDERED_NODE_SIZE = 3`, fixed minimum edge gap, and fixed center distance.
- Current runtime diagnostics publish screen spacing and `visibleViewportNodeCount`; they do not publish overview visible color count, visible island count, dominant island share, and visible node-type inventory.
- Sigma package defaults are `itemSizesReference: "screen"` and `zoomToSizeRatioFunction: Math.sqrt`.
- Sigma `scaleSize` uses `itemSizesReference === "positions"` to multiply by `cameraRatio * graphToViewportRatio`. That setting changes node-size zoom semantics.
- Commit `67ba0dd` introduced the default-load readable camera path and changed the previous `animatedReset`.
- The pre-regression baseline commit is `80a7972`, which is `67ba0dd^`.

Additional reopened facts from 2026-05-28:

- `GraphCanvas.tsx` registers `camera?.on?.('updated', handleCameraUpdated)`.
- `handleCameraUpdated` recomputes orientation labels and schedules screen diagnostics.
- `recordCurrentScreenNodeSpacing` calls both `buildScreenNodeSpacingDiagnostics(sigma)` and `buildGraphOverviewDiagnostics(sigma)`.
- `buildScreenNodeSpacingDiagnostics` walks `graph.nodes()`, calls `sigma.getNodeDisplayData`, `sigma.graphToViewport`, `sigma.scaleSize`, then performs same-island neighbor-pair checks.
- `buildGraphOverviewDiagnostics` walks `graph.nodes()`, calls `sigma.getNodeDisplayData` and `sigma.graphToViewport`, then counts visible colors, rings, islands, and node types.
- `buildGraphOrientationLabels` walks the graph with `graph.forEachNode`.
- `useSigma.zoomIn` and `useSigma.zoomOut` record interaction modes for button zoom only.
- Wheel zoom reaches Sigma and changes camera ratio, but no dedicated wheel interaction mode is recorded.
- AVmatrix impact for `buildScreenNodeSpacingDiagnostics`, `buildGraphOverviewDiagnostics`, and `buildGraphOrientationLabels` is `CRITICAL`.
- AVmatrix impact for `applyFilterBasedClusteredLayout` is `HIGH`.
- `avmatrix-launcher\build.ps1` sets `$LauncherOutPath = Join-Path $LauncherRoot "AVmatrixLauncher.exe"`.
- `avmatrix-launcher\build.ps1` runs `npm run build` in `avmatrix-web`.
- `avmatrix-launcher\build.ps1` builds the canonical CLI at `avmatrix\bin\avmatrix.exe`.
- `avmatrix-launcher\build.ps1` builds the launcher with `go build -ldflags="-s -w -H=windowsgui" -o $LauncherOutPath .`.
- `avmatrix-launcher\build.ps1` builds `avmatrix-launcher\server-bundle\avmatrix-server.exe`.
- `avmatrix-launcher\build.ps1` deletes and recopies `avmatrix-web\dist` to `avmatrix-launcher\web-dist`.
- Therefore the script is the correct full build gate for this requirement, but validation must prove the launcher exe and packaged Web bundle were actually refreshed.

## Non-Negotiable Product Invariants

- First restore the original color/overview behavior. Scale, spacing, zoom, and camera implementation starts after this gate passes.
- The first implementation slice reverts the default-load camera path introduced in `67ba0dd`.
- Fixes preserve every baseline node type: none are removed; none are hidden; none become unreachable; none are visually de-emphasized.
- Default load is an overview of the graph, not an automatic jump into the densest local patch.
- Default overview preserves multi-island and multi-color information on fixtures that contain multiple visible node types and app layers.
- Expected ring, island, color, node-type, and visible-node counts are computed from the active graph inventory. For the current dense fixture, the computed product expectation is exactly `3` rings, `4` islands, `3` visible node types, and `2077` default-visible nodes; repositories with different language/taxonomy inventories must produce their own exact counts.
- Every exact inventory count in this plan is measured fixture evidence, not a product constant.
- User zoom-in visibly increases rendered node radius in viewport pixels.
- User zoom-out visibly decreases rendered node radius in viewport pixels.
- Detail/focus mode makes raw nodes readable without overlap.
- Same-island detail spacing derives from actual rendered node size.
- Island radius derives from node count, node size, required spacing, and density.
- Ring radius and ring gaps derive from expanded island footprints.
- Edge thickness and label behavior stay coherent with current camera scale.
- Tests protect overview, node-type inventory, zoom semantics, and detail spacing together.

## Scope

Implementation touches:

- Web graph layout geometry in `avmatrix-web/src/lib/graph-adapter.ts`;
- a new single-owner scale model module `avmatrix-web/src/lib/graph-scale-model.ts`;
- Sigma/camera/render behavior in `avmatrix-web/src/hooks/useSigma.ts`;
- readable camera behavior in `avmatrix-web/src/lib/graph-readable-camera.ts`;
- graph screen diagnostics in `avmatrix-web/src/lib/graph-screen-spacing.ts`;
- browser runtime diagnostics in `avmatrix-web/src/lib/runtime-diagnostics.ts`;
- graph canvas diagnostics/test hooks in `avmatrix-web/src/components/GraphCanvas.tsx`;
- Web unit tests under `avmatrix-web/test/unit`;
- Web e2e tests under `avmatrix-web/e2e`;
- evidence and benchmark ledgers for this plan.
- graph diagnostics scheduling and throttling in `avmatrix-web/src/components/GraphCanvas.tsx`;
- graph overview and screen-spacing diagnostic split/caching behavior;
- orientation label inventory caching and camera-placement refresh behavior;
- wheel-zoom interaction telemetry and e2e behavior coverage;
- packaged launcher Web dist freshness validation.

Out of scope:

- backend graph analysis semantics;
- graph schema and relationship extraction;
- generated API contract semantics unrelated to Web graph rendering;
- replacing Sigma;
- changing node taxonomy and graph filter taxonomy;
- editing generated `AGENTS.md`, `CLAUDE.md`, and generated skill content directly as source of truth.

## Required Technical Direction

The implementation introduces one scale model module: `avmatrix-web/src/lib/graph-scale-model.ts`.

That module owns these calculations:

- rendered node radius at a camera ratio;
- graph-unit-to-viewport-pixel conversion at a camera ratio;
- required screen-space edge gap;
- required screen-space center distance;
- required graph-coordinate center distance;
- overview camera policy;
- detail/focus camera policy;
- edge-width and label-threshold policy bound to the same camera scale.

The model rejects a global `3px` node radius, rejects a globally valid dense spacing constant, and rejects camera zoom changes that ignore node size.

### Baseline Overview Restoration

The first implementation slice restores default overview behavior exactly:

1. Create a non-destructive baseline worktree from commit `80a7972` at `.tmp/graph-baseline-80a7972`.
2. Capture baseline screenshot and metrics from that worktree.
3. In current HEAD, replace the current default-load path in `useSigma.setGraph` with `sigma.getCamera().animatedReset({ duration: 500 })`.
4. Remove the default-load call to `applyReadableGraphCamera(sigma)`.
5. Restore `minCameraRatio: 0.002` for overview behavior.
6. Remove `itemSizesReference: 'positions'` so Sigma uses its default `"screen"` sizing semantics.
7. Keep readable camera code out of default load.
8. Add Phase 1 overview diagnostics for visible color count, visible island count, dominant island share, visible node-type inventory, and filter node-type inventory.
9. Add tests that fail when default overview collapses to one visible color. Add tests that fail when default overview collapses to one visible island. Add tests that fail when baseline node-type inventory is reduced.

This slice fixes the color regression before spacing work starts because the palette is still correct and the regression is a camera/render semantic regression.

### Overview Vs Detail

The implementation separates two modes:

| Mode | Purpose | Behavior |
|---|---|---|
| Overview | Preserve whole-codebase structure, color, ring/island orientation, and navigation | Uses baseline reset behavior, smaller nodes, multi-color/multi-island viewport |
| Detail/focus | Inspect raw nodes in a local area | Uses explicit focus/search/selection trigger, readable node size, no-overlap spacing |

Default load always uses overview.

Readable/detail behavior is triggered only by explicit selection, search result focus, island focus, node focus, and a named detail mode command.

### Zoom Semantics

Zoom is tested as product behavior:

1. record rendered radius at initial camera ratio;
2. trigger zoom in;
3. record rendered radius again;
4. assert radius increased;
5. trigger another zoom in;
6. assert radius increased again;
7. trigger zoom out;
8. assert radius decreased.

The implementation uses Sigma `"screen"` sizing semantics so user zoom visibly changes node radius according to Sigma's `zoomToSizeRatioFunction`.

### Dynamic Scale Model Algorithm

The scale model uses viewport-pixel measurement as the authority:

```text
pxPerGraphUnit = distance(
  sigma.graphToViewport({ x: probe.x + 1, y: probe.y }),
  sigma.graphToViewport({ x: probe.x, y: probe.y })
)

renderedRadiusPx(node, cameraRatio) = sigma.scaleSize(node.size, cameraRatio)

requiredEdgeGapPx = max(renderedDiameterPx among compared nodes)

requiredCenterDistancePx(nodeA, nodeB) =
  renderedRadiusPx(nodeA) + renderedRadiusPx(nodeB) + requiredEdgeGapPx

requiredCenterDistanceGraph(nodeA, nodeB) =
  requiredCenterDistancePx(nodeA, nodeB) / pxPerGraphUnit
```

For equal-size nodes, this enforces the requested rule:

```text
edge-to-edge gap = one node diameter
center-to-center distance = radius + diameter + radius
```

### Layout Algorithm

Same-island dense placement uses deterministic hexagonal packing.

Algorithm:

1. Sort nodes by stable node id.
2. Compute `cellSpacingGraph` from the dynamic scale model's required center distance at detail/focus camera ratio.
3. Place nodes on axial hex coordinates expanding ring by ring around island center:

```text
x = centerX + cellSpacingGraph * (q + r / 2)
y = centerY + cellSpacingGraph * sqrt(3) / 2 * r
```

4. Validate every same-island neighbor pair in viewport pixels.
5. Expand `cellSpacingGraph` by the measured violation ratio and recompute the island until overlap count is zero and target-gap violation count is zero.
6. Compute island radius from final placed bounds plus dynamic required graph margin.
7. Compute app-layer ring radii from island footprints.
8. Compute inter-ring gaps from the largest adjacent expanded footprint.
9. Remove post-placement jitter from paths protected by the no-overlap invariant.

This algorithm is deterministic, scales with node count, and keeps geometry derived from rendered size instead of fixed constants.

### Fixed-Size Assumption Removal

Implementation replaces fixed geometry with scale-model policy inputs:

- max rendered node size;
- dense node size;
- minimum camera ratio;
- cluster node spacing;
- island gap;
- ring gap;
- edge size in dense mode;
- label thresholds;
- e2e thresholds that lock success to one island.

Only named product policy constants remain, such as minimum readable pixel target. Overview color, ring, island, and node-type expectations come from the active graph inventory.

### Reopened Performance And Interaction Direction

The Phase 8 implementation keeps every feature and changes scheduling, caching, and instrumentation only.

The Web graph must separate three work classes:

| Work class | Trigger | Main-thread budget |
|---|---|---|
| Static graph inventory | graph/filter/semantic-filter changes | full graph walk allowed |
| Cheap camera sample | camera update, wheel, button zoom, pan, resize | one coalesced RAF, no pairwise full-graph audit |
| Full viewport spacing audit | camera settled, explicit validation, e2e/debug sample | idle/settled task, never every camera tick |

Required scheduler algorithm:

```text
onGraphOrFilterChanged:
  rebuild cached graph inventories
  rebuild cached orientation label sources
  schedule full viewport spacing audit

onCameraUpdated:
  mark camera moving
  schedule one RAF if none is pending
  reset camera-settled timer

RAF camera sample:
  place orientation labels from cached label sources
  record cheap camera and rendered-radius sample
  do not run full graph pairwise spacing audit
  do not rebuild static inventory

camera-settled timer:
  after 160ms without camera update, schedule full audit with requestIdleCallback
  fallback to setTimeout when requestIdleCallback is unavailable

idle full audit:
  run buildScreenNodeSpacingDiagnostics once
  run visible overview diagnostics once
  publish graphInteraction sample with the current interaction mode
```

Wheel zoom must get first-class semantics:

```text
on wheel zoom:
  recordGraphInteractionMode({ mode: 'wheel-zoom' })
  assert cameraRatio changes
  assert rendered node radius changes
```

The implementation must not remove node types, edges, colors, filters, labels, diagnostics, or screenshots. Any reduction in work must come from scheduling, caching, coalescing, and moving full audits out of interactive frames.

### Full Build Gate Definition

For every reopened implementation slice, the full build gate is:

```powershell
powershell -ExecutionPolicy Bypass -File avmatrix-launcher\build.ps1
```

This gate fails unless all required product artifacts are freshly produced and verified:

- `avmatrix-web\dist`;
- `avmatrix-launcher\web-dist`;
- `avmatrix\bin\avmatrix.exe`;
- `avmatrix-launcher\server-bundle\avmatrix-server.exe`;
- `avmatrix-launcher\AVmatrixLauncher.exe`.

The evidence ledger must record the launcher exe `LastWriteTime`, file size, and content hash after the build. Browser validation must use the launcher-served app on `127.0.0.1:5228` after this build, not only Vite preview or `avmatrix-web\dist`.

## Acceptance Criteria

- AVmatrix refresh and impact checks are recorded before implementation edits.
- The known-good color overview is captured from baseline commit `80a7972` using a non-destructive worktree.
- Current UI restores color/overview parity with baseline before dynamic scale work starts.
- Current UI restores node-type parity with baseline before dynamic scale work starts.
- Every baseline graph/filter node type remains present after the fix.
- Every baseline-visible overview node type remains visibly rendered after the fix.
- Fixed-size assumptions are inventoried before edits.
- Sigma node-size behavior under zoom is empirically verified.
- Phase 1 overview diagnostics exist before e2e parity assertions are written.
- `graph-scale-model.ts` exists and owns graph scale derivation.
- Default graph load restores overview behavior.
- Default overview on a multi-type dense fixture reports the computed visible color, island, ring, node-type, and default-visible node inventories for that fixture.
- The current dense fixture reports exactly `3` rings, `4` islands, `3` visible node types, and `2077` default-visible nodes, with those expectations derived from the fixture graph instead of hardcoded into assertions.
- Default overview dominant island share stays below `0.85` on the multi-island dense fixture.
- Zoom-in increases rendered node radius.
- Zoom-out decreases rendered node radius.
- Detail/focus mode produces readable raw nodes without same-island overlap.
- Detail/focus target gap derives from current rendered node size.
- Tests no longer encode `frontend:Function` as the successful dense viewport.
- Browser screenshots prove overview, zoom sequence, and detail/focus behavior.
- Full build, focused Web unit tests, full Web unit tests, and Web e2e/browser tests pass.
- AVmatrix detect-changes is run before implementation commit.
- Each completed implementation slice is committed before continuing.
- Reopened Phase 8 records the current long-task, frame-delta, diagnostics-write, and wheel-ratio baseline before edits.
- Reopened Phase 8 uses `powershell -ExecutionPolicy Bypass -File avmatrix-launcher\build.ps1` as the full build gate.
- The full build evidence proves `avmatrix-launcher\AVmatrixLauncher.exe` was rebuilt.
- The full build evidence proves `avmatrix-launcher\web-dist` was recopied from the current `avmatrix-web\dist`.
- Camera update handling no longer runs full `buildScreenNodeSpacingDiagnostics` and `buildGraphOverviewDiagnostics` on every camera tick.
- Orientation labels preserve the same visible label behavior without walking the full graph on every camera tick.
- Wheel zoom has a dedicated runtime interaction mode and e2e coverage.
- Wheel zoom changes camera ratio and rendered node radius in the browser test.
- Dense fixture wheel burst no longer produces a single long task above `100ms`.
- Dense fixture wheel burst max frame delta stays below `150ms`.
- Full pairwise screen-spacing audit runs after camera settle and remains available for e2e/debug validation.
- Static graph inventories remain exact and are recomputed on graph/filter changes.
- No node type, edge type, color, label, filter, or diagnostic field is removed to meet performance targets.
- `avmatrix-launcher\web-dist` is rebuilt from the current Web build and verified to contain the current diagnostics/interaction code.
- Packaged launcher validation proves the user-facing `5228` bundle matches current Web behavior.

## Phase 0 - Plan Creation

- [x] P0-A Create dynamic scale model plan file.
- [x] P0-B Create separate evidence ledger.
- [x] P0-C Create separate benchmark ledger.
- [x] P0-D Re-read plan, codebase, AVmatrix graph evidence, and revise this plan to a single concrete implementation direction.

## Phase 1 - Restore Original Color Overview Gate

- [x] P1-A Run `avmatrix analyze --force`.
- [x] P1-B Use AVmatrix context/query to trace current graph color assignment, initial camera, Sigma render, readable camera, layout, and overview diagnostics owners.
- [x] P1-C Run impact analysis before editing each planned function/class/method/exported symbol involved in color, camera, render, layout, and diagnostics.
- [x] P1-D Create non-destructive baseline worktree `.tmp/graph-baseline-80a7972` from commit `80a7972`.
- [x] P1-E Capture baseline browser screenshots and metrics for visible colors, visible islands, visible node types, filter node-type inventory, labels, and camera state.
- [x] P1-F Capture current HEAD browser screenshots and metrics for the same viewport and fixture.
- [x] P1-G Compare current HEAD against the baseline and record the exact color/island/node-type/label regression.
- [x] P1-H Restore default-load overview by reverting the `useSigma.setGraph` camera path to `animatedReset({ duration: 500 })`.
- [x] P1-I Remove default-load `applyReadableGraphCamera(sigma)` and keep readable camera out of graph load.
- [x] P1-J Restore overview camera sizing by using `minCameraRatio: 0.002`.
- [x] P1-K Restore Sigma default screen-size semantics by removing `itemSizesReference: 'positions'`.
- [x] P1-L Add Phase 1 overview diagnostics in `runtime-diagnostics.ts` and `GraphCanvas.tsx` for visible color count, visible ring count, visible island count, dominant island share, visible ring inventory, visible node-type inventory, graph ring inventory, and filter node-type inventory.
- [x] P1-M Add tests that fail when the default viewport collapses to one color on a multi-color fixture.
- [x] P1-N Add tests that fail when any baseline node type disappears from graph/filter inventory and baseline-visible overview rendering.
- [x] P1-O Record color and node-type parity evidence and benchmark metrics.
- [x] P1-P Run full build before tests, then focused Web unit tests, e2e tests, and browser screenshot validation for this slice.
- [x] P1-Q Run AVmatrix detect-changes for the slice.
- [x] P1-R Commit the completed color/overview/node-type restoration slice.

## Phase 2 - Audit Scale, Zoom, Spacing Failure And Blast Radius

- [x] P2-A Use AVmatrix context/query to trace graph layout, Sigma render, camera, diagnostics, and e2e owners after Phase 1.
- [x] P2-B Run impact analysis before editing each planned function/class/method/exported symbol.
- [x] P2-C Record HIGH/CRITICAL blast radius as warnings where present.
- [x] P2-D Inventory all fixed-size and fixed-camera assumptions.
- [x] P2-E Capture browser baseline screenshots for overview, zoom sequence, and detail/focus after color parity is restored.
- [x] P2-F Record baseline benchmark metrics in `B1`.
- [x] P2-G Update evidence ledger with audit findings.

## Phase 3 - Implement Dynamic Graph Scale Model

- [x] P3-A Add `avmatrix-web/src/lib/graph-scale-model.ts`.
- [x] P3-B Implement rendered radius derivation through Sigma `scaleSize`.
- [x] P3-C Implement graph-unit-to-viewport-pixel conversion.
- [x] P3-D Implement required edge gap, center distance in pixels, and center distance in graph units.
- [x] P3-E Replace fixed rendered-size helpers with scale-model policy inputs.
- [x] P3-F Add focused unit tests for scale derivation and zoom radius behavior.
- [x] P3-G Record scale-model evidence and metrics.
- [x] P3-H Run full build before tests, then focused Web unit tests for this slice.
- [x] P3-I Run AVmatrix detect-changes for the slice.
- [x] P3-J Commit the completed scale-model slice.

## Phase 4 - Refactor Layout To Dynamic Hex Packing

- [x] P4-A Refactor same-island placement to deterministic hexagonal packing using scale-model center distance.
- [x] P4-B Refactor island radius to derive from final placed bounds and dynamic margin.
- [x] P4-C Refactor island gap to derive from expanded island footprints.
- [x] P4-D Refactor ring radius/gap to derive from adjacent expanded island footprints.
- [x] P4-E Remove post-placement jitter from no-overlap layout paths.
- [x] P4-F Add dense layout unit tests for multi-size and high-node-count fixtures.
- [x] P4-G Record layout benchmark metrics.
- [x] P4-H Run full build before tests, then focused Web unit tests for this slice.
- [x] P4-I Run AVmatrix detect-changes for the slice.
- [x] P4-J Commit the completed layout slice.

## Phase 5 - Correct Zoom And Detail Behavior

- [x] P5-A Keep default reset/load behavior in overview mode.
- [x] P5-B Ensure zoom-in visibly increases node radius.
- [x] P5-C Ensure zoom-out visibly decreases node radius.
- [x] P5-D Move readable camera behavior behind explicit selection, search result focus, island focus, node focus, and named detail mode.
- [x] P5-E Keep selection/search/blast-radius highlight behavior intact.
- [x] P5-F Add unit tests for camera mode selection.
- [x] P5-G Record overview and zoom benchmark metrics.
- [x] P5-H Run full build before tests, then focused Web unit tests and e2e zoom tests for this slice.
- [x] P5-I Run AVmatrix detect-changes for the slice.
- [x] P5-J Commit the completed camera/zoom slice.

## Phase 6 - Diagnostics And E2E Coverage

- [x] P6-A Extend Phase 1 overview diagnostics with zoom radius samples, detail/focus samples, and dynamic gap samples.
- [x] P6-B Replace e2e assertions that accept a single `frontend:Function` viewport as success.
- [x] P6-C Add overview e2e assertions for baseline color parity and multi-color/multi-island default load.
- [x] P6-D Add zoom e2e assertions for rendered radius growth/shrink.
- [x] P6-E Add detail/focus e2e assertions for no overlap and dynamic target gap.
- [x] P6-F Capture desktop and small-viewport screenshots.
- [x] P6-G Record diagnostics evidence and benchmark metrics.
- [x] P6-H Run full build before tests, then full Web unit tests and Web e2e/browser tests for this slice.
- [x] P6-I Run AVmatrix detect-changes for the slice.
- [x] P6-J Commit the completed diagnostics/e2e slice.

## Phase 7 - Full Validation And Closure

- [x] P7-A Run the full build gate before tests.
- [x] P7-B Run focused Web unit tests.
- [x] P7-C Run full Web unit tests.
- [x] P7-D Run Web e2e/browser tests and inspect screenshots.
- [x] P7-E Run final `git diff --check`.
- [x] P7-F Run final `avmatrix analyze --force`.
- [x] P7-G Run final `avmatrix detect-changes --repo AVmatrix --scope all`.
- [x] P7-H Record final evidence and benchmark summaries.
- [x] P7-I Commit the final implementation slice.
- [x] P7-J Confirm working tree state and remaining plan items.

## Phase 8 - Reopened Render Hot Path And Wheel Zoom Hardening

- [x] P8-A Run `avmatrix analyze --force`.
- [x] P8-B Use AVmatrix query/context to trace `GraphCanvas` camera updates, screen spacing diagnostics, graph overview diagnostics, orientation labels, wheel zoom, and package/launcher serving paths.
- [x] P8-C Run impact analysis before editing each planned function/class/method/exported symbol in graph render, diagnostics, orientation labels, zoom interaction, and package build flow.
- [x] P8-D Record HIGH/CRITICAL blast radius warnings in evidence and proceed carefully.
- [x] P8-E Reproduce and record baseline browser performance metrics for current build and packaged launcher: load long tasks, wheel long tasks, max frame delta, diagnostics write counts, camera ratio, rendered radius, visible color count, visible ring count, visible island count, and screenshot artifacts.
- [x] P8-F Add reusable browser performance probe coverage for dense graph load, wheel zoom, button zoom, and camera settle.
- [x] P8-G Add `wheel-zoom` interaction mode in runtime diagnostics.
- [x] P8-H Add e2e coverage that sends mouse wheel input and asserts camera ratio and rendered node radius change.
- [x] P8-I Split graph diagnostics into cached static inventory, cheap camera sample, and settled full viewport spacing audit.
- [x] P8-J Coalesce camera-update diagnostics to one pending RAF and prevent duplicate diagnostics writes during the same frame.
- [x] P8-K Move full pairwise `buildScreenNodeSpacingDiagnostics` execution to camera-settled idle work.
- [x] P8-L Move full visible overview diagnostics execution to graph/filter changes and camera-settled idle work.
- [x] P8-M Cache orientation label source inventory on graph/filter changes and use camera updates only for placement from cached sources.
- [x] P8-N Preserve all existing node types, edges, colors, filters, labels, graph inventories, screen spacing diagnostics, graph overview diagnostics, and screenshots.
- [x] P8-O Run the full build gate: `powershell -ExecutionPolicy Bypass -File avmatrix-launcher\build.ps1`.
- [x] P8-P Verify the full build refreshed `avmatrix-launcher\AVmatrixLauncher.exe` by recording `LastWriteTime`, file size, and hash.
- [x] P8-Q Verify `avmatrix-launcher\web-dist` was recopied from the current `avmatrix-web\dist`.
- [x] P8-R Verify the launcher-served bundle on `127.0.0.1:5228` contains current graph overview, graph interaction, and wheel zoom diagnostics.
- [x] P8-S Run focused Web unit tests for diagnostics scheduling, wheel interaction mode, orientation label cache behavior, and graph inventory preservation.
- [x] P8-T Run full Web unit tests.
- [x] P8-U Run Web e2e/browser tests with screenshots for dense overview, wheel zoom, button zoom, and detail/focus.
- [x] P8-V Record performance benchmark results after each benchmarkable task.
- [x] P8-W Record evidence after each evidenced task.
- [x] P8-X Run `avmatrix detect-changes --repo AVmatrix --scope all`.
- [x] P8-Y Commit the completed reopened render hot path and wheel zoom hardening slice.
- [x] P8-Z Confirm working tree state and close the reopened plan only after launcher and current build both pass.

## Risk Notes

Blast radius is HIGH/CRITICAL because graph layout, Sigma rendering, camera behavior, and e2e fixtures are central Web UI behavior. The implementation proceeds carefully with evidence and tests; HIGH/CRITICAL is not a reason to avoid required code changes.

The main risk is solving only one symptom again. This plan closes only after overview, node-type preservation, zoom semantics, and detail spacing pass together.

The reopened risk is treating performance as a reason to remove graph information. That is forbidden for this plan. Performance work must preserve the feature set and reduce repeated work by using scheduling, caching, coalescing, and settled/idle audits.
