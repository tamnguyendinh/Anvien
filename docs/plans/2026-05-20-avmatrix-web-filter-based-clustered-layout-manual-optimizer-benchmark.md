# AVmatrix Web Filter-Based Clustered Layout And Manual Optimizer Benchmark Ledger

Date: 2026-05-20

Status: active

Companion files:

- Plan: [2026-05-20-avmatrix-web-filter-based-clustered-layout-manual-optimizer-plan.md](2026-05-20-avmatrix-web-filter-based-clustered-layout-manual-optimizer-plan.md)
- Evidence ledger: [2026-05-20-avmatrix-web-filter-based-clustered-layout-manual-optimizer-evidence.md](2026-05-20-avmatrix-web-filter-based-clustered-layout-manual-optimizer-evidence.md)

## Benchmark Rules

Record benchmarkable results when they measure product/runtime performance, capacity, memory, graph load behavior, optimizer behavior, render latency, or interaction latency. Build/test/e2e timings are validation evidence unless the slice changes those systems.

Do not record inferred or estimated values. Every benchmark row must name the command, repo path, graph source, commit, and interpretation.

## B0 - Required Baseline

Date: 2026-05-20

Status: pending measurement

Required baseline before implementation:

| Measurement | Source | Before | After | Status |
|---|---|---:|---:|---|
| Graph node count | runtime diagnostics or graph payload | pending | pending | pending |
| Graph relationship count | runtime diagnostics or graph payload | pending | pending | pending |
| Graph conversion duration | `window.__AVMATRIX_WEB_DIAGNOSTICS__` | pending | pending | pending |
| Layout starts after initial graph load | `window.__AVMATRIX_WEB_DIAGNOSTICS__.layout.starts` | pending | pending | pending |
| Layout stops after initial graph load | `window.__AVMATRIX_WEB_DIAGNOSTICS__.layout.stops` | pending | pending | pending |
| Initial clustered layout rendered without manual optimizer | browser/e2e observation | pending | pending | pending |
| Manual optimizer invocations after user click | runtime diagnostics | pending | pending | pending |
| Manual optimizer stop/apply result | runtime diagnostics | pending | pending | pending |
| Node filter toggle latency | browser/e2e or measured script | pending | pending | pending |
| Selection/focus latency | browser/e2e or measured script | pending | pending | pending |

Expected final interpretation:

- initial graph load layout starts should be `0`;
- manual optimizer invocation count should increase only after explicit user action;
- graph renders in clustered form before any manual optimizer action;
- node filters continue to hide/show their matching clusters.

## B1 - Cluster Layout Inventory

Date: 2026-05-20

Status: pending measurement

Record the final cluster inventory after implementation:

| Cluster source | Count | Notes |
|---|---:|---|
| existing filterable node labels | pending | should match `FILTERABLE_LABELS` |
| node labels present in graph payload | pending | should be subset or fallback-compatible |
| clusters rendered on initial load | pending | only labels present in graph; known labels follow `FILTERABLE_LABELS`, unknown labels append by label string |
| nodes with deterministic position changes across repeated conversion | pending | expected `0` for same input |

Required command or test:

```powershell
npm --prefix avmatrix-web run test -- test/unit/graph-adapter.edge-geometry.test.ts
```

If a more specific layout test file is added, record that command instead.

## B2 - No-Auto-Optimizer Benchmark

Date: 2026-05-20

Status: pending measurement

Record before/after behavior:

| Scenario | Before | After | Expected after |
|---|---:|---:|---:|
| load graph, do not click optimizer, layout starts | pending | pending | 0 |
| load graph, do not click optimizer, layout running indicator appears | pending | pending | 0 |
| click optimizer once, optimizer invocation count | pending | pending | 1 |
| click optimizer once, graph remains interactive | pending | pending | observed pass |

Use runtime diagnostics and browser/e2e artifacts. Do not substitute unit test pass/fail for runtime measurement if browser behavior is the claim.

## B3 - Interaction And Readability Benchmark

Date: 2026-05-20

Status: pending measurement

Record final Web interaction observations:

| Interaction | Before | After | Status |
|---|---:|---:|---|
| initial graph usable before optimizer | pending | pending | pending |
| node type filter toggle remains responsive | pending | pending | pending |
| graph-health filter toggle remains responsive | pending | pending | pending |
| selected node focus still works | pending | pending | pending |
| edge visibility toggle still works | pending | pending | pending |

Readability acceptance is qualitative but must be backed by screenshot, browser observation, or e2e assertion. The key product benchmark is not "optimizer is faster"; it is "graph is readable before optimizer runs".

## B4 - Final Benchmark Summary

Date: 2026-05-20

Status: pending

Fill at closure:

- before commit:
- after commit:
- graph source:
- repo path:
- final node count:
- final relationship count:
- initial layout starts after load:
- manual optimizer invocations after click:
- deterministic layout test result:
- browser/e2e result:
- residual risk:
