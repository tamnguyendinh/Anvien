# AVmatrix Web Graph Dynamic Scale Model And Zoom Semantics Benchmark Ledger

Date: 2026-05-27

Status: Completed

Companion files:

- Plan: [2026-05-27-avmatrix-web-graph-dynamic-scale-model-plan.md](2026-05-27-avmatrix-web-graph-dynamic-scale-model-plan.md)
- Evidence ledger: [2026-05-27-avmatrix-web-graph-dynamic-scale-model-evidence.md](2026-05-27-avmatrix-web-graph-dynamic-scale-model-evidence.md)

## Benchmark Rules

Benchmark results must be recorded as soon as each benchmarkable task is completed.

Benchmarkable items for this plan include:

- measured rendered node radius at multiple camera zoom ratios;
- rendered node radius growth factor after user zoom-in actions;
- visible viewport color count, ring count, and island count in default overview;
- visible viewport node count, color count, ring count, and island count in detail/focus views;
- same-island screen-space overlap count and target-gap violation count;
- minimum observed screen-space edge gap;
- graph coordinate bounds and viewport dimensions;
- layout conversion time for dense fixtures;
- graph inventory counts for dense fixtures;
- browser screenshot artifact inventory;
- e2e validation metrics for desktop and small viewports.

Build/test/e2e durations are validation evidence for this plan.

## B0 - Required Metric Inventory

Status: completed.

| Metric group | Completion evidence |
|---|---|
| Baseline and restored overview parity | B1, B1A |
| Fixed-assumption audit baseline | B1B |
| Dynamic scale model metrics | B2 |
| Dynamic hex layout metrics | B2A |
| Overview and zoom browser metrics | B3 |
| Detail/focus browser metrics | B4, B4A |
| Final validation inventory | B5 |

## B1 - Baseline Capture Before Implementation

Status: completed.

Baseline worktree:

- Source commit: `80a7972`
- Worktree path: `.tmp/graph-baseline-80a7972`

Recorded metrics:

| Metric | Baseline `80a7972` | Current pre-fix |
|---|---:|---:|
| Dense fixture node count | 1480 | 1480 |
| Overview ring label count | 3 | 2 |
| Overview island label count | 4 | 0 |
| Layout island count | 5 | 5 |
| Layout overlap count | 0 | 0 |
| Layout target-gap violation count | 0 | 0 |
| Current visible viewport node count | not available in baseline | 40 |
| Current visible viewport island count | not available in baseline | 1 |
| Current readable camera applied | not present in baseline | true |

Restored Phase 1 metrics:

| Metric | Restored Phase 1 |
|---|---:|
| Dense fixture node count | 1480 |
| Overview ring label count | 3 |
| Overview island label count | 4 |
| Visible viewport node count | 1400 |
| Visible color count | 3 |
| Visible ring count | 3 |
| Visible island count | 4 |
| Dominant island share | 0.7142857142857143 |
| Readable camera applied | false |
| Camera ratio | 1 |

Completed before code edits:

- create `.tmp/graph-baseline-80a7972` from commit `80a7972`;
- capture baseline browser screenshot and metrics;
- capture current default overview screenshot;
- capture current two-step zoom screenshot sequence;
- record rendered node radius at initial, zoom step 1, and zoom step 2;
- record visible color count and visible island count at initial view;
- record dominant island share;
- record current detail/focus overlap and target-gap metrics.

## B1A - Color Overview Parity Gate

Status: completed for Phase 1.

This benchmark gate must pass before dynamic scale, spacing, and zoom implementation work proceeds.

Record:

- baseline commit `80a7972`;
- baseline visible color count;
- baseline visible ring count;
- baseline visible island count;
- baseline graph/visible ring inventory;
- baseline graph/filter node-type inventory;
- baseline overview visible node-type inventory;
- baseline ring/island label counts;
- current HEAD visible color count;
- current HEAD visible ring count;
- current HEAD visible island count;
- current HEAD graph/visible ring inventory;
- current HEAD graph/filter node-type inventory;
- current HEAD overview visible node-type inventory;
- restored visible color count;
- restored visible ring count;
- restored visible island count;
- restored graph/visible ring inventory;
- restored graph/filter node-type inventory;
- restored overview visible node-type inventory;
- parity result for baseline overview behavior.

## B1B - Phase 2 Audit Browser Baseline

Status: completed.

Artifacts:

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

Dense fixture metrics:

| Stage | Camera ratio | Max rendered radius px | Visible nodes | Visible island count | Overlap count | Target-gap violations |
|---|---:|---:|---:|---:|---:|---:|
| Overview | 1 | 3 | 1400 | 4 | 27425 | 97426 |
| Zoom in 1 | 0.6666666666666666 | 3.6742346141747673 | 40 | 1 | 18411 | 68085 |
| Zoom in 2 | 0.4444444444444444 | 4.5 | 0 | 0 | 12240 | 47002 |
| Zoom out | 0.6666666666666666 | 3.6742346141747673 | 40 | 1 | 18411 | 68085 |

Detail/focus fixture metrics:

| Metric | Value |
|---|---:|
| Selection succeeded | true |
| Focus button clicked | true |
| Camera ratio after focus | 1 |
| Max rendered radius after focus | 3 |
| Visible viewport node count after focus | 1 |

Audit benchmark result:

- Zoom radius grows from `3` to `3.6742346141747673` to `4.5`.
- Visible dense overview inventory regresses during zoom: visible island count drops from `4` to `1` to `0`.
- Detail/focus click path keeps camera ratio at `1` after focus button click.
- Diagnostics require a resize event to record zoom camera changes in the current code.

## B2 - Dynamic Scale Model Metrics

Status: completed for Phase 3 scale-model slice.

Focused unit metrics:

| Metric | Value |
|---|---:|
| Radius for size `3` at camera ratio `1` | 3 |
| Radius for size `3` at camera ratio `0.25` | 6 |
| Required edge gap for two radius-`3` nodes | 6 |
| Required center distance for two radius-`3` nodes | 12 |
| Projection scale from `x * 5` fixture | 5 px/graph-unit |
| Model camera ratio fixture | 0.25 |
| Model min radius from sizes `2`, `3` | 4 |
| Model max radius from sizes `2`, `3` | 6 |
| Model required edge gap | 12 |
| Model required center distance px | 24 |
| Model required center distance graph units | 12 |

Validation inventory:

- Full Web build: passed.
- Focused Web unit tests: `3` files, `34` tests passed.
- Scale model test inventory: `6` tests passed.

## B2A - Dynamic Hex Layout Metrics

Status: completed for Phase 4 layout slice.

Dense unit fixture:

| Metric | Value |
|---|---:|
| Dense fixture node count | 6100 |
| Function island node count | 1800 |
| Method island node count | 1400 |
| Dynamic center-distance floor | 12 |
| Dynamic edge-gap floor | 6 |
| Function overlap count | 0 |
| Function target-gap violation count | 0 |
| Method overlap count | 0 |
| Method target-gap violation count | 0 |

Validation inventory:

- Full Web build: passed.
- Focused Web unit tests: `3` files, `34` tests passed.
- Web e2e/browser tests: `3` tests passed.
- Browser screenshot artifacts: `4` PNG files.

E2E dense fixture:

| Metric | Value |
|---|---:|
| Dense fixture graph nodes | 2157 |
| Dense fixture graph edges | 2156 |
| Default-visible node count | 2077 |
| Expected ring count from fixture inventory | 3 |
| Expected island count from fixture inventory | 4 |
| Expected visible node-type count from fixture inventory | 3 |

## B3 - Overview And Zoom Browser Metrics

Status: completed for Phase 5 camera and zoom slice.

Artifact:

- `reports/problem/2026-05-27-graph-phase5-camera-zoom-metrics.json`

Dense fixture inventory:

| Metric | Value |
|---|---:|
| Dense fixture graph nodes | 2157 |
| Dense fixture graph edges | 2156 |
| Default-visible node count | 2077 |
| Visible color count from active inventory | 3 |
| Visible ring count from active inventory | 3 |
| Visible island count from active inventory | 4 |
| Visible node-type count from active inventory | 3 |

Zoom metrics:

| Stage | Camera ratio | Max rendered radius px | Visible nodes | Visible islands | Dominant island share |
|---|---:|---:|---:|---:|---:|
| Initial overview | 1 | 3 | 2077 | 4 | 0.8074145402022147 |
| Zoom in 1 | 0.6666666666666666 | 3.6742346141747673 | 773 | 2 | 0.9482535575679172 |
| Zoom in 2 | 0.4444444444444444 | 4.5 | 40 | 1 | 1 |
| Zoom out | 0.6666666666666666 | 3.6742346141747673 | 773 | 2 | 0.9482535575679172 |

Growth metrics:

| Metric | Value |
|---|---:|
| Zoom-in 1 radius growth | 1.2247448713915892 |
| Zoom-in 2 radius growth | 1.224744871391589 |
| Zoom-out radius shrink | 0.816496580927726 |

## B4 - Detail Spacing Browser Metrics

Status: completed.

Phase 5 explicit focus metrics:

| Stage | Camera ratio | Max rendered radius px | Visible nodes | Visible islands |
|---|---:|---:|---:|---:|
| Search focus | 0.140625 | 8 | 787 | 1 |
| Same selection shifted by zoom-out | 0.2109375 | 6.531972647421808 | 1397 | 1 |
| Same selection refocus | 0.140625 | 8 | 787 | 1 |

Focus growth metrics:

| Metric | Value |
|---|---:|
| Search focus radius growth against prior zoom-out | 2.1773242158072694 |
| Same-selected refocus radius growth after zoom-out | 1.224744871391589 |

Screenshot artifacts:

- `reports/problem/2026-05-27-graph-phase5-search-focus.png`
- `reports/problem/2026-05-27-graph-phase5-same-selected-refocus.png`

Detail spacing diagnostics were completed in B4A.

## B4A - Phase 6 Diagnostics And E2E Metrics

Status: completed for Phase 6 diagnostics and e2e slice.

Dense fixture inventory from active graph:

| Metric | Value |
|---|---:|
| Dense fixture graph nodes | 2157 |
| Dense fixture graph edges | 2156 |
| Default-visible node count | 2077 |
| Visible ring count from active inventory | 3 |
| Visible island count from active inventory | 4 |
| Visible node-type count from active inventory | 3 |
| Visible color count from active inventory | 3 |

Runtime interaction sample inventory:

| Sample list | Count |
|---|---:|
| Overview samples | 6 |
| Zoom samples | 12 |
| Detail/focus samples | 7 |
| Dynamic-gap samples | 12 |

Overview sample metrics:

| Metric | Value |
|---|---:|
| Visible viewport nodes | 2077 |
| Visible viewport islands | 4 |
| Camera ratio | 1 |
| Max rendered radius px | 3 |
| Dominant island share | 0.8074145402022147 |

Zoom sample metrics:

| Stage | Camera ratio | Max rendered radius px | Visible nodes | Visible islands |
|---|---:|---:|---:|---:|
| Initial overview | 1 | 3 | 2077 | 4 |
| Zoom in 1 | 0.6666666666666666 | 3.6742346141747673 | 773 | 2 |
| Zoom in 2 | 0.4444444444444444 | 4.5 | 40 | 1 |
| Zoom out | 0.6666666666666666 | 3.6742346141747673 | 773 | 2 |

Detail/focus diagnostic metrics:

| Metric | Value |
|---|---:|
| Search focus visible viewport nodes | 800 |
| Search focus visible viewport islands | 1 |
| Search focus max rendered radius px | 7.8670397179696385 |
| Search focus max rendered diameter px | 15.734079435939277 |
| Search focus minimum observed edge gap px | 16.412827972783575 |
| Search focus max required center distance px | 31.468158871878554 |
| Search focus overlap count | 0 |
| Search focus target-gap violation count | 0 |

Validation inventory:

- Full Web build before tests: passed.
- Focused runtime diagnostics unit test: `1` file, `9` tests passed.
- Full Web unit tests: `50` files, `397` tests passed.
- Graph e2e/browser tests: `3` tests passed.
- Screenshot artifacts: `4` PNG files.

## B5 - Final Validation Inventory

Status: completed.

Final validation inventory:

| Gate | Result |
|---|---:|
| Full Web build | passed |
| Focused Web unit files | 5 |
| Focused Web unit tests | 47 |
| Full Web unit files | 50 |
| Full Web unit tests | 397 |
| Web e2e/browser files | 1 |
| Web e2e/browser tests | 3 |
| Browser screenshot PNG artifacts inspected | 2 |
| Final AVmatrix changed files at closure detect | 1 |
| Final AVmatrix changed symbols at closure detect | 2 |
| Final affected processes | 0 |
| Final degraded resolution-health nodes | 0 |

Final screenshot artifacts inspected:

- `avmatrix-web/test-results/graph-orientation-labels-G-537e2-ory-visible-on-default-load-chromium/graph-node-spacing-dense-desktop.png`
- `avmatrix-web/test-results/graph-orientation-labels-G-537e2-ory-visible-on-default-load-chromium/graph-node-spacing-dense-small.png`

## B6 - Phase 8 Baseline Performance

Status: completed for P8-E baseline.

Current Web build baseline:

| Metric | Value |
|---|---:|
| Dense fixture graph nodes | 2157 |
| Dense fixture graph relationships | 2156 |
| Ready time ms | 1591 |
| Load long tasks | 3 |
| Load total long-task ms | 1085 |
| Load longest long-task ms | 461 |
| Load max frame delta ms | 616.6 |
| Load screen diagnostics writes | 2 |
| Load overview diagnostics writes | 2 |
| Graph conversion last ms | 43.80000001192093 |
| Load camera ratio | 1 |
| Load max rendered radius px | 3 |
| Load visible color count | 3 |
| Load visible ring count | 3 |
| Load visible island count | 4 |
| Wheel camera ratio | 0.5882352941176471 |
| Wheel max rendered radius px | 3.9115214431215892 |
| Wheel long tasks | 4 |
| Wheel total long-task ms | 446 |
| Wheel longest long-task ms | 119 |
| Wheel max frame delta ms | 133.30000000000018 |
| Wheel screen diagnostics writes | 4 |
| Wheel overview diagnostics writes | 4 |
| Wheel zoom samples | 0 |
| Button camera ratio | 0.39215686274509803 |
| Button max rendered radius px | 4.790615826801394 |
| Button long tasks | 4 |
| Button total long-task ms | 329 |
| Button max frame delta ms | 83.40000000000009 |
| Button screen diagnostics writes | 4 |
| Button overview diagnostics writes | 4 |
| Button zoom samples | 4 |

Packaged launcher baseline:

| Metric | Value |
|---|---:|
| Dense fixture graph nodes | 2157 |
| Dense fixture graph relationships | 2156 |
| Ready time ms | 2427 |
| Load long tasks | 4 |
| Load total long-task ms | 1766 |
| Load longest long-task ms | 964 |
| Load max frame delta ms | 1166.7 |
| Load screen diagnostics writes | 2 |
| Load overview diagnostics writes | 0 |
| Graph conversion last ms | 329.69999998807907 |
| Load camera ratio | 0.004001487892151231 |
| Load max rendered radius px | 2.0000000473025636 |
| Load graph overview present | 0 |
| Load graph interaction present | 0 |
| Wheel camera ratio | 0.004001487892151231 |
| Wheel long tasks | 6 |
| Wheel total long-task ms | 374 |
| Wheel max frame delta ms | 66.79999999999927 |
| Wheel screen diagnostics writes | 0 |
| Wheel overview diagnostics writes | 0 |

Artifacts:

- `reports/problem/2026-05-28-phase8-current-build-baseline.json`
- `reports/problem/2026-05-28-phase8-current-build-baseline-load.png`
- `reports/problem/2026-05-28-phase8-current-build-baseline-wheel.png`
- `reports/problem/2026-05-28-phase8-packaged-launcher-baseline.json`
- `reports/problem/2026-05-28-phase8-packaged-launcher-baseline-load.png`

## B7 - Phase 8 Final Launcher Performance

Status: completed for P8-V.

Final launcher benchmark:

| Metric | Value |
|---|---:|
| Dense fixture graph nodes | 2157 |
| Dense fixture graph relationships | 2156 |
| Ready time ms | 1431 |
| Load long tasks | 4 |
| Load total long-task ms | 1122 |
| Load longest long-task ms | 446 |
| Load max frame delta ms | 633.3000000000001 |
| Load screen diagnostics writes | 2 |
| Load overview diagnostics writes | 2 |
| Graph conversion last ms | 51.5 |
| Load camera ratio | 1 |
| Load max rendered radius px | 3 |
| Load visible color count | 3 |
| Load visible ring count | 3 |
| Load visible island count | 4 |
| Wheel camera ratio | 0.5882352941176471 |
| Wheel max rendered radius px | 3.9115214431215892 |
| Wheel long tasks | 6 |
| Wheel total long-task ms | 420 |
| Wheel longest long-task ms | 83 |
| Wheel max frame delta ms | 83.30000000000018 |
| Wheel screen diagnostics writes | 1 |
| Wheel overview diagnostics writes | 1 |
| Wheel zoom samples | 1 |
| Button camera ratio | 0.39215686274509803 |
| Button max rendered radius px | 4.790615826801394 |
| Button long tasks | 5 |
| Button total long-task ms | 308 |
| Button longest long-task ms | 75 |
| Button max frame delta ms | 66.69999999999982 |
| Button screen diagnostics writes | 1 |
| Button overview diagnostics writes | 1 |
| Button zoom samples | 2 |

Baseline to final comparison:

| Metric | Baseline current build | Baseline packaged launcher | Final launcher |
|---|---:|---:|---:|
| Ready time ms | 1591 | 2427 | 1431 |
| Load longest long-task ms | 461 | 964 | 446 |
| Load max frame delta ms | 616.6 | 1166.7 | 633.3000000000001 |
| Load visible color count | 3 | unavailable | 3 |
| Load visible ring count | 3 | unavailable | 3 |
| Load visible island count | 4 | unavailable | 4 |
| Wheel screen diagnostics writes | 4 | 0 | 1 |
| Wheel overview diagnostics writes | 4 | 0 | 1 |
| Wheel zoom samples | 0 | unavailable | 1 |
| Wheel max frame delta ms | 133.30000000000018 | 66.79999999999927 | 83.30000000000018 |
| Button screen diagnostics writes | 4 | unavailable | 1 |
| Button overview diagnostics writes | 4 | unavailable | 1 |
| Button zoom samples | 4 | unavailable | 2 |

Final artifacts:

- `reports/problem/2026-05-28-phase8-launcher-final.json`
- `reports/problem/2026-05-28-phase8-launcher-final-load.png`
- `reports/problem/2026-05-28-phase8-launcher-final-wheel.png`
