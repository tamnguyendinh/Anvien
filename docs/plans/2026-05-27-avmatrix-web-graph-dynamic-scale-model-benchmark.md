# AVmatrix Web Graph Dynamic Scale Model And Zoom Semantics Benchmark Ledger

Date: 2026-05-27

Status: Planned

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

Status: pending.

| Metric | Unit | Baseline | Final | Target |
|---|---:|---:|---:|---:|
| Baseline overview visible color count | colors | pending | pending | reference value from `80a7972` |
| Current HEAD overview visible color count | colors | pending | pending | restored to baseline parity before scale work |
| Baseline overview visible ring count | rings | pending | pending | reference value from `80a7972` |
| Current HEAD overview visible ring count | rings | pending | pending | restored to baseline parity before scale work |
| Baseline overview visible island count | islands | pending | pending | reference value from `80a7972` |
| Current HEAD overview visible island count | islands | pending | pending | restored to baseline parity before scale work |
| Baseline graph/filter node-type inventory | node types | pending | pending | reference inventory from `80a7972` |
| Current/restored graph/filter node-type inventory | node types | pending | pending | zero missing baseline node type |
| Baseline overview visible node-type count | node types | pending | pending | reference visible inventory from `80a7972` |
| Current/restored overview visible node-type count | node types | pending | pending | zero missing baseline-visible node type |
| Phase 1 overview diagnostics inventory | fields | pending | pending | visible color count, visible ring count, visible island count, dominant island share, visible ring inventory, visible node-type inventory, graph ring inventory, filter node-type inventory |
| Baseline overview screenshot inventory | files | pending | pending | captured from `.tmp/graph-baseline-80a7972` |
| Current/restored overview screenshot inventory | files | pending | pending | restored screenshots match baseline behavior |
| Initial overview visible color count | colors | pending | pending | equals computed visible color inventory for fixture/baseline |
| Initial overview visible ring count | rings | pending | pending | equals computed visible ring inventory for fixture/baseline |
| Initial overview visible island count | islands | pending | pending | equals computed visible island inventory for fixture/baseline |
| Initial overview ring label count | labels | pending | pending | equals computed visible ring label inventory for fixture/baseline |
| Initial overview dominant island share | ratio | pending | pending | `< 0.85` on multi-island dense fixture |
| Initial overview max rendered node radius | px | pending | pending | record, not forced to detail size |
| First zoom-in rendered node radius growth | ratio | pending | pending | `> 1.0` |
| Second zoom-in rendered node radius growth | ratio | pending | pending | greater than first zoom-in |
| Zoom-out rendered node radius shrink | ratio | pending | pending | `< 1.0` against previous zoom step |
| Detail/focus max rendered node radius | px | pending | pending | readable target from scale model |
| Detail/focus overlap count | violations | pending | pending | `0` |
| Detail/focus target-gap violation count | violations | pending | pending | `0` |
| Detail/focus minimum edge gap | px | pending | pending | `>= dynamic required gap` |
| Dense layout conversion time | ms | pending | pending | no material regression without evidence |
| Dense graph node count | nodes | pending | pending | record |
| Dense graph edge count | edges | pending | pending | record |

## B1 - Baseline Capture Before Implementation

Status: in progress.

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

Required before code edits:

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

Status: pending.

Record after camera/zoom behavior is implemented:

- initial overview visible color count;
- initial overview visible island count;
- initial overview dominant island share;
- zoom-in rendered radius growth;
- zoom-out rendered radius shrink;
- reset returns to overview, not a dense local patch.

## B4 - Detail Spacing Browser Metrics

Status: pending.

Record after detail/focus behavior is implemented:

- selected detail viewport node count;
- selected detail viewport same-island node count;
- overlap count;
- target-gap violation count;
- minimum observed edge gap;
- maximum rendered node diameter;
- screenshot artifacts.

## B5 - Final Validation Inventory

Status: pending.

Record final:

- full build result;
- focused unit test inventory;
- full Web unit test inventory;
- Web e2e/browser test inventory;
- screenshot artifact paths;
- AVmatrix change detection summary before implementation commit.
