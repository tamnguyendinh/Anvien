# AVmatrix Web Graph Node Spacing And Overlap Hardening Benchmark Ledger

Date: 2026-05-26

Status: Complete

Companion files:

- Plan: [2026-05-26-avmatrix-web-graph-node-spacing-overlap-hardening-plan.md](2026-05-26-avmatrix-web-graph-node-spacing-overlap-hardening-plan.md)
- Evidence ledger: [2026-05-26-avmatrix-web-graph-node-spacing-overlap-hardening-evidence.md](2026-05-26-avmatrix-web-graph-node-spacing-overlap-hardening-evidence.md)

## Benchmark Rules

This file records quantitative data only: graph layout geometry metrics, overlap counts, minimum distance measurements, graph footprint size, fixture sizes, label counts, visible overlap counts, test pass/fail counts, and measured runtime behavior.

Narrative evidence, commands, logs, source observations, screenshots, and source traces belong in the evidence ledger.

Use `pending` only when a future phase has not measured that value yet.

Build/test/e2e timings are validation evidence unless the implementation changes build/test/e2e performance behavior directly.

## B0 - Source And Test Inventory Counts

Status: plan creation measured.

| Metric | Unit | Baseline | Latest | Delta | Final target |
|---|---:|---:|---:|---:|---:|
| Known Web layout owner files | files | 1 | 1 | 0 | >= 1 |
| Known Web render/camera owner files | files | 1 | 1 | 0 | >= 1 if camera behavior changes |
| Known graph canvas owner files | files | 1 | 1 | 0 | >= 1 if diagnostics change |
| Existing unit geometry test files identified | files | 1 | 1 | 0 | >= 1 |
| Existing browser graph label e2e files identified | files | 1 | 1 | 0 | >= 1 |
| Existing hard pairwise same-island node gap tests | tests | 0 | 2 | +2 | >= 2 |
| Existing dense-overlap regression fixtures | fixtures | 0 | 1 | +1 | >= 1 |

## B1 - Node Spacing Contract Metrics

Status: baseline measured; final pending.

| Metric | Unit | Baseline | Latest | Delta | Final target |
|---|---:|---:|---:|---:|---:|
| Maximum rendered node size used for spacing | layout units | 3 | 3 | 0 | record |
| Verified rendered node diameter | layout units | 6 | 6 | 0 | record |
| Required minimum edge gap | layout units | 6 | 6 | 0 | >= 1 rendered node diameter |
| Required minimum center distance | layout units | 12 | 12 | 0 | >= 2 rendered node diameters |
| Organic/jitter violations allowed | violations | pending | pending | pending | 0 |

## B2 - Dense Island Geometry Metrics

Status: baseline measured; final pending.

| Fixture | Node count | Baseline min center distance | Baseline min edge gap | Baseline overlap count | Final min center distance | Final min edge gap | Final overlap count | Target |
|---|---:|---:|---:|---:|---:|---:|---:|---|
| Small island | 100 | 40.535 | 34.535 | 0 | 40.535 | 34.535 | 0 | no overlap and edge gap >= 1 diameter |
| Medium island | 1000 | 3.144 | -2.856 | 4 | 13.417 | 7.417 | 0 | no overlap and edge gap >= 1 diameter |
| Large dense island | 1800 | 0.361 | -5.639 | 22 | 12.048 | 6.048 | 0 | no overlap and edge gap >= 1 diameter |
| Regression fixture from current bug | 1800 | 2.911 | -3.089 | 4 | 12.107 | 6.107 | 0 | no overlap and edge gap >= 1 diameter |

## B3 - Layout Footprint Metrics

Status: implementation measured; browser footprint pending.

| Metric | Unit | Baseline | Latest | Delta | Final target |
|---|---:|---:|---:|---:|---:|
| Largest dense island radius | layout units | pending | pending | pending | derived from node count and minimum gap |
| Minimum island-to-island gutter | layout units | pending | pending | pending | >= configured island gutter |
| Minimum macro-ring gutter | layout units | pending | pending | pending | >= configured ring gutter |
| Graph layout width for dense fixture | layout units | pending | 2633.931 | pending | record |
| Graph layout height for dense fixture | layout units | pending | 2642.245 | pending | record |
| Rail-like dense island shape violations | violations | pending | 0 | pending | 0 |

## B4 - Browser UX Metrics

Status: browser validation measured.

| Metric | Unit | Baseline | Latest | Delta | Final target |
|---|---:|---:|---:|---:|---:|
| Dense browser fixture node count | nodes | pending | 1480 | pending | record |
| Dense browser fixture relationship count | relationships | pending | 1479 | pending | record |
| Desktop viewport width | px | pending | 1280 | pending | record |
| Desktop viewport height | px | pending | 800 | pending | record |
| Smaller viewport width | px | pending | 520 | pending | record |
| Smaller viewport height | px | pending | 720 | pending | record |
| Desktop visible node overlap violations | violations | pending | 0 | pending | 0 |
| Smaller-viewport visible node overlap violations | violations | pending | 0 | pending | 0 |
| Desktop target gap violations | violations | pending | 0 | pending | 0 |
| Smaller-viewport target gap violations | violations | pending | 0 | pending | 0 |
| Desktop browser diagnostic minimum edge gap | layout units | pending | 6 | pending | >= 6 |
| Smaller-viewport browser diagnostic minimum edge gap | layout units | pending | 6 | pending | >= 6 |
| Desktop visible label overlap violations | violations | pending | 0 | pending | 0 |
| Smaller-viewport visible label overlap violations | violations | pending | 0 | pending | 0 |
| Desktop ring labels visible | labels | pending | 3 | pending | >= 3 |
| Desktop island labels visible | labels | pending | 4 | pending | >= 1 |
| Smaller-viewport ring labels visible | labels | pending | 3 | pending | >= 3 |
| Smaller-viewport island labels visible | labels | pending | 3 | pending | >= 1 |
| Dense graph desktop screenshots captured | files | pending | 1 | pending | >= 1 |
| Dense graph smaller-viewport screenshots captured | files | pending | 1 | pending | >= 1 |
| Filter-change stale layout/label violations | violations | pending | 0 | pending | 0 |

## B5 - Validation Metrics

Status: pending implementation validation.

| Metric | Unit | Baseline | Latest | Delta | Final target |
|---|---:|---:|---:|---:|---:|
| Full build gate result | pass/fail | pending | pass | pass | pass |
| Focused graph geometry tests result | pass/fail | pending | pass | pass | pass |
| Full Web unit test result | pass/fail | pending | pass | pass | pass |
| Web e2e/browser graph spacing result | pass/fail | pending | pass | pass | pass |
| `detect-changes` pre-commit result | pass/fail | pending | pass | pass | pass |
| Implementation commits recorded | commits | 0 | 1 | +1 | >= 1 after implementation slice |

## B6 - Reopened Screen-Space Failure Metrics

Status: final measured.

| Metric | Unit | Previous latest | Reopen observation | Final measured | Target |
|---|---:|---:|---:|---:|---:|
| Dense browser fixture node count | nodes | 1480 | pending | 2157 | record |
| Dense browser fixture relationship count | relationships | 1479 | pending | 2156 | record |
| Dense fixture `Function` island size | nodes | pending | 1677 | 1677 | >= 1677 |
| Browser diagnostic coordinate space | category | layout coordinates | insufficient for screenshot failure | viewport_px | screen-projected viewport pixels |
| Layout diagnostic node count | nodes | 1480 | pending | 2157 | record |
| Layout diagnostic island count | islands | pending | pending | 5 | record |
| Layout minimum center distance | layout units | 12 | pending | 12 | >= 12 |
| Layout minimum edge gap | layout units | 6 | pending | 6 | >= 6 |
| Layout overlap count | violations | 0 | pending | 0 | 0 |
| Layout target gap violations | violations | 0 | pending | 0 | 0 |
| Desktop screen diagnostic node count | nodes | pending | pending | 2077 | record |
| Desktop screen diagnostic island count | islands | pending | pending | 4 | record |
| Desktop viewport width | px | 1280 page | pending | 1032 graph canvas | record |
| Desktop viewport height | px | 800 page | pending | 702 graph canvas | record |
| Desktop camera ratio | ratio | pending | pending | 0.008524702358442132 | record |
| Desktop min rendered radius | px | pending | pending | 0.6249999779166772 | record |
| Desktop max rendered radius | px | pending | pending | 0.7499999735000125 | >= 0.74 |
| Desktop max rendered diameter | px | pending | pending | 1.499999947000025 | record |
| Desktop screen minimum center distance | px | pending | pending | 2.99999989400005 | >= max required center |
| Desktop screen minimum edge gap | px | pending | pending | 1.499999947000025 | >= one rendered node diameter |
| Desktop screen max required center distance | px | pending | pending | 2.99999989400005 | record |
| Desktop screen overlap count | violations | pending | visible blob | 0 | 0 |
| Desktop screen target gap violations | violations | pending | pending | 0 | 0 |
| Desktop ring labels visible | labels | 3 | pending | 2 | >= 1 |
| Desktop island labels visible | labels | 4 | pending | 1 | record |
| Desktop label overlap count | violations | 0 | pending | 0 | 0 |
| Small screen diagnostic node count | nodes | pending | pending | 2077 | record |
| Small screen diagnostic island count | islands | pending | pending | 4 | record |
| Small viewport width | px | 520 page | pending | 272 graph canvas | record |
| Small viewport height | px | 720 page | pending | 606 graph canvas | record |
| Small camera ratio | ratio | pending | pending | 0.00043012603675720635 | record |
| Small min rendered radius | px | pending | pending | 0.624999992016073 | record |
| Small max rendered radius | px | pending | pending | 0.7499999904192876 | >= 0.74 |
| Small max rendered diameter | px | pending | pending | 1.499999980838575 | record |
| Small screen minimum center distance | px | pending | pending | 2.99999996167715 | >= max required center |
| Small screen minimum edge gap | px | pending | pending | 1.499999980838575 | >= one rendered node diameter |
| Small screen max required center distance | px | pending | pending | 2.99999996167715 | record |
| Small screen overlap count | violations | pending | pending | 0 | 0 |
| Small screen target gap violations | violations | pending | pending | 0 | 0 |
| Small ring labels visible | labels | 3 | pending | 2 | >= 1 |
| Small island labels visible | labels | 3 | pending | 0 | record |
| Small label overlap count | violations | 0 | pending | 0 | 0 |
| Dense graph desktop screenshots captured | files | 1 | pending | 1 | >= 1 |
| Dense graph small screenshots captured | files | 1 | pending | 1 | >= 1 |
| Closure status | status | complete | reopened | complete | complete only after screen-space metrics pass |

## B7 - Node Visibility Regression Metrics

Status: final measured.

| Metric | Unit | P6 final | P7 final | Target |
|---|---:|---:|---:|---:|
| Dense browser fixture node count | nodes | 2157 | 2157 | record |
| Dense browser fixture relationship count | relationships | 2156 | 2156 | record |
| Dense fixture `Function` island size | nodes | 1677 | 1677 | >= 1677 |
| Dense graph readable camera target max radius | px | 0.75 | 2.0 | visible node marker |
| Dense layout spacing for `>1000` nodes | layout units | 36 | 12 | minimum center-distance contract |
| Dense layout minimum center distance | layout units | 12 | 12 | >= 12 |
| Dense layout minimum edge gap | layout units | 6 | 6 | >= 6 |
| Dense layout overlap count | violations | 0 | 0 | 0 |
| Dense layout target gap violations | violations | 0 | 0 | 0 |
| Dense island-to-island gutter floor | layout units | derived from spacing | `minimumNodeCenterDistance * 75` | preserve >= 900 test threshold |
| Desktop graph canvas width | px | 1032 | 1032 | record |
| Desktop graph canvas height | px | 702 | 702 | record |
| Desktop visible viewport node count | nodes | not measured | 20 | >= 16 |
| Desktop visible `frontend:Function` nodes | nodes | not measured | 20 | >= 16 |
| Desktop max rendered node radius | px | 0.75 | 2.0 | >= 1.99 |
| Desktop max rendered node diameter | px | 1.5 | 4.0 | record |
| Desktop screen minimum center distance | px | 3.0 | 8.0 | >= max required center |
| Desktop screen minimum edge gap | px | 1.5 | 4.0 | >= one rendered node diameter |
| Desktop screen overlap count | violations | 0 | 0 | 0 |
| Desktop screen target gap violations | violations | 0 | 0 | 0 |
| Small graph canvas width | px | 272 | 272 | record |
| Small graph canvas height | px | 606 | 606 | record |
| Small visible viewport node count | nodes | not measured | 1 | >= 1 |
| Small visible `frontend:Function` nodes | nodes | not measured | 1 | >= 1 |
| Small max rendered node radius | px | 0.75 | >= 1.99 | >= 1.99 |
| Small screen overlap count | violations | 0 | 0 | 0 |
| Small screen target gap violations | violations | 0 | 0 | 0 |
| Focused Web unit tests | tests | pass | 42 passed | pass |
| Full Web unit tests | tests | 384 passed | 384 passed | pass |
| Web e2e graph tests | tests | 3 passed | 3 passed | pass |
| Dense graph desktop screenshot | files | 1 | 1 | >= 1 |
| Dense graph small screenshot | files | 1 | 1 | >= 1 |
