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

## E2 - Dynamic Scale Model Evidence

Status: pending.

Record once implemented:

- `graph-scale-model.ts` as the single source of truth for graph scale;
- how node screen radius is measured with Sigma `scaleSize`;
- how graph-coordinate distance converts to viewport pixels;
- how required node center distance is derived;
- how zoom changes node radius;
- how overview differs from detail/focus mode;
- which old fixed constants were removed, replaced, and constrained.

## E3 - Overview Preservation Evidence

Status: pending.

Record once implemented:

- default load behavior;
- screenshot paths for overview;
- visible color count;
- visible island count;
- dominant island share;
- ring/island label counts;
- proof that default load does not auto-focus only the densest local island.

## E4 - Zoom Semantics Evidence

Status: pending.

Record once implemented:

- node radius at initial camera ratio;
- node radius after first zoom in;
- node radius after second zoom in;
- node radius after zoom out;
- proof that node radius changes monotonically with zoom direction;
- screenshot paths for zoom sequence.

## E5 - Detail Spacing Evidence

Status: pending.

Record once implemented:

- detail/focus behavior trigger;
- same-island node spacing diagnostics;
- overlap count;
- target-gap violation count;
- minimum edge gap;
- screenshot paths for detail/focus views.

## E6 - Validation Evidence

Status: pending.

Record:

- full build gate command and result;
- focused unit test command and result;
- full Web unit test command and result;
- Web e2e/browser command and result;
- screenshot artifact paths;
- every failure and fix loop before final pass.

## E7 - Pre-Commit Change Detection

Status: pending.

Before committing implementation work, record:

- final `git diff --check`;
- final `avmatrix analyze --force`;
- final `avmatrix detect-changes --repo AVmatrix --scope all`;
- summary: risk level, changed files, changed symbols, affected count, changed app layers, changed functional areas, resolution health impact.
