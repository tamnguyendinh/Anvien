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
- visible viewport color count and island count in default overview;
- visible viewport node count, color count, and island count in detail/focus views;
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
| Baseline overview visible island count | islands | pending | pending | reference value from `80a7972` |
| Current HEAD overview visible island count | islands | pending | pending | restored to baseline parity before scale work |
| Baseline graph/filter node-type inventory | node types | pending | pending | reference inventory from `80a7972` |
| Current/restored graph/filter node-type inventory | node types | pending | pending | zero missing baseline node type |
| Baseline overview visible node-type count | node types | pending | pending | reference visible inventory from `80a7972` |
| Current/restored overview visible node-type count | node types | pending | pending | zero missing baseline-visible node type |
| Phase 1 overview diagnostics inventory | fields | pending | pending | visible color count, visible island count, dominant island share, visible node-type inventory, filter node-type inventory |
| Baseline overview screenshot inventory | files | pending | pending | captured from `.tmp/graph-baseline-80a7972` |
| Current/restored overview screenshot inventory | files | pending | pending | restored screenshots match baseline behavior |
| Initial overview visible color count | colors | pending | pending | `>= 3` on multi-type dense fixture |
| Initial overview visible island count | islands | pending | pending | `>= 3` on multi-island dense fixture |
| Initial overview ring label count | labels | pending | pending | `>= 3` on multi-ring fixture |
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

Status: pending.

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

Status: pending.

This benchmark gate must pass before dynamic scale, spacing, and zoom implementation work proceeds.

Record:

- baseline commit `80a7972`;
- baseline visible color count;
- baseline visible island count;
- baseline graph/filter node-type inventory;
- baseline overview visible node-type inventory;
- baseline ring/island label counts;
- current HEAD visible color count;
- current HEAD visible island count;
- current HEAD graph/filter node-type inventory;
- current HEAD overview visible node-type inventory;
- restored visible color count;
- restored visible island count;
- restored graph/filter node-type inventory;
- restored overview visible node-type inventory;
- parity result for baseline overview behavior.

## B2 - Dynamic Scale Model Metrics

Status: pending.

Record after the scale model is implemented:

- graph-unit-to-viewport-pixel scale at each measured camera ratio;
- measured Sigma node-size behavior at each measured camera ratio;
- expected vs observed rendered node radius;
- expected vs observed required center distance;
- layout bounds before and after dynamic spacing.

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
