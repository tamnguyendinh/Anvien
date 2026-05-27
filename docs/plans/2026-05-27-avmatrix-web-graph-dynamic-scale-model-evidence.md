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

Status: pending.

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
- `avmatrix-web/src/components/GraphCanvas.tsx`
- `avmatrix-web/e2e/graph-orientation-labels.spec.ts`
- related unit tests under `avmatrix-web/test/unit`

## E1A - Color Overview Parity Gate

Status: pending.

This gate must be completed before dynamic scale, zoom, and spacing implementation proceeds.

Required evidence:

- baseline source: commit `80a7972`;
- command and browser steps used to capture the baseline through `.tmp/graph-baseline-80a7972`;
- baseline screenshot paths;
- current HEAD screenshot paths;
- restored screenshot paths;
- visible color count comparison;
- visible island count comparison;
- graph/filter node-type inventory comparison;
- overview-visible node-type inventory comparison;
- missing node-type list, expected empty;
- ring/island label comparison;
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
