# AVmatrix Web Graph Node Spacing And Overlap Hardening Benchmark Ledger

Date: 2026-05-26

Status: Active

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
| Existing hard pairwise same-island node gap tests | tests | pending | pending | pending | >= 2 |
| Existing dense-overlap regression fixtures | fixtures | pending | pending | pending | >= 1 |

## B1 - Node Spacing Contract Metrics

Status: pending implementation measurement.

| Metric | Unit | Baseline | Latest | Delta | Final target |
|---|---:|---:|---:|---:|---:|
| Maximum rendered node size used for spacing | layout units | pending | pending | pending | record |
| Verified rendered node diameter | layout units | pending | pending | pending | record |
| Required minimum edge gap | layout units | pending | pending | pending | >= 1 rendered node diameter |
| Required minimum center distance | layout units | pending | pending | pending | >= 2 rendered node diameters |
| Organic/jitter violations allowed | violations | pending | pending | pending | 0 |

## B2 - Dense Island Geometry Metrics

Status: pending baseline and final measurement.

| Fixture | Node count | Baseline min center distance | Baseline min edge gap | Baseline overlap count | Final min center distance | Final min edge gap | Final overlap count | Target |
|---|---:|---:|---:|---:|---:|---:|---:|---|
| Small island | pending | pending | pending | pending | pending | pending | pending | no overlap and edge gap >= 1 diameter |
| Medium island | pending | pending | pending | pending | pending | pending | pending | no overlap and edge gap >= 1 diameter |
| Large dense island | pending | pending | pending | pending | pending | pending | pending | no overlap and edge gap >= 1 diameter |
| Regression fixture from current bug | pending | pending | pending | pending | pending | pending | pending | no overlap and edge gap >= 1 diameter |

## B3 - Layout Footprint Metrics

Status: pending implementation measurement.

| Metric | Unit | Baseline | Latest | Delta | Final target |
|---|---:|---:|---:|---:|---:|
| Largest dense island radius | layout units | pending | pending | pending | derived from node count and minimum gap |
| Minimum island-to-island gutter | layout units | pending | pending | pending | >= configured island gutter |
| Minimum macro-ring gutter | layout units | pending | pending | pending | >= configured ring gutter |
| Graph layout width for dense fixture | layout units | pending | pending | pending | record |
| Graph layout height for dense fixture | layout units | pending | pending | pending | record |
| Rail-like dense island shape violations | violations | pending | pending | pending | 0 |

## B4 - Browser UX Metrics

Status: pending browser validation.

| Metric | Unit | Baseline | Latest | Delta | Final target |
|---|---:|---:|---:|---:|---:|
| Desktop visible node overlap violations | violations | pending | pending | pending | 0 |
| Smaller-viewport visible node overlap violations | violations | pending | pending | pending | 0 |
| Desktop visible label overlap violations | violations | pending | pending | pending | 0 |
| Smaller-viewport visible label overlap violations | violations | pending | pending | pending | 0 |
| Dense graph desktop screenshots captured | files | pending | pending | pending | >= 1 |
| Dense graph smaller-viewport screenshots captured | files | pending | pending | pending | >= 1 |
| Filter-change stale layout/label violations | violations | pending | pending | pending | 0 |

## B5 - Validation Metrics

Status: pending implementation validation.

| Metric | Unit | Baseline | Latest | Delta | Final target |
|---|---:|---:|---:|---:|---:|
| Full build gate result | pass/fail | pending | pending | pending | pass |
| Focused graph geometry tests result | pass/fail | pending | pending | pending | pass |
| Full Web unit test result | pass/fail | pending | pending | pending | pass |
| Web e2e/browser graph spacing result | pass/fail | pending | pending | pending | pass |
| `detect-changes` pre-commit result | pass/fail | pending | pending | pending | pass |
| Implementation commits recorded | commits | 0 | 0 | 0 | >= 1 after implementation slice |
