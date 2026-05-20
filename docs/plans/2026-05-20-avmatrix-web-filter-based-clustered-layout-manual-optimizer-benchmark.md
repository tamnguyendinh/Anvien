# AVmatrix Web Filter-Based Clustered Layout And Manual Optimizer Benchmark Ledger

Date: 2026-05-20

Status: complete

Companion files:

- Plan: [2026-05-20-avmatrix-web-filter-based-clustered-layout-manual-optimizer-plan.md](2026-05-20-avmatrix-web-filter-based-clustered-layout-manual-optimizer-plan.md)
- Evidence ledger: [2026-05-20-avmatrix-web-filter-based-clustered-layout-manual-optimizer-evidence.md](2026-05-20-avmatrix-web-filter-based-clustered-layout-manual-optimizer-evidence.md)

## Benchmark Rules

Record benchmarkable results when they measure product/runtime performance, capacity, memory, graph load behavior, optimizer behavior, render latency, or interaction latency. Build/test/e2e timings are validation evidence unless the slice changes those systems.

Do not record inferred or estimated values. Every benchmark row must name the command, repo path, graph source, commit, and interpretation.

## B0 - Required Baseline

Date: 2026-05-20

Status: recorded

Representative runtime measurement:

- Repo: `Restaurant_manager`
- Graph source: local indexed graph served through `go run .\cmd\avmatrix serve`
- Frontend: local Vite dev server on `127.0.0.1:5228`
- Capture method: Playwright opened the Web UI, waited for `Ready`, captured `window.__AVMATRIX_WEB_DIAGNOSTICS__`, clicked `Optimize Layout`, then captured diagnostics again.

| Measurement | Source | Before | After | Status |
|---|---|---:|---:|---|
| Graph node count | runtime diagnostics | `78358` | `78358` | recorded |
| Graph relationship count | runtime diagnostics | `130588` | `130588` | recorded |
| Graph conversion duration | diagnostics `graphConversion.lastMs` | `1029.5ms` | `1029.5ms` | recorded |
| Layout starts after initial graph load | diagnostics `layout.starts` | `0` | `0` | passed |
| Layout stops after initial graph load | diagnostics `layout.stops` | `0` | `0` | passed |
| Initial clustered layout rendered without manual optimizer | e2e + diagnostics | yes | yes | passed |
| Manual optimizer invocations after user click | diagnostics `layout.manualOptimizerInvocations` | `0` | `1` | passed |
| Manual optimizer apply result | diagnostics `layout.lastManualOptimizerRunMs` | `0ms` | `1183ms` | recorded |
| Heartbeat reconnects | diagnostics `heartbeat.reconnects` | `0` | `0` | passed |
| Reconnect banner shows | diagnostics `reconnectBanner.shows` | `0` | `0` | passed |

Expected final interpretation:

- initial graph load layout starts should be `0`;
- manual optimizer invocation count should increase only after explicit user action;
- graph renders in clustered form before any manual optimizer action;
- node filters continue to hide/show their matching clusters.

## B1 - Cluster Layout Inventory

Date: 2026-05-20

Status: recorded

Record the final cluster inventory after implementation:

| Cluster source | Count | Notes |
|---|---:|---|
| existing filterable node labels | `36` | `FILTERABLE_LABELS = NODE_LABELS` |
| node labels present in measured graph payload | `17` | from diagnostics `visualScale.maxSizeByLabel` keys |
| clusters rendered on initial load | `17` | only labels present in graph; known labels follow `FILTERABLE_LABELS`, unknown labels append by label string |
| nodes with deterministic position changes across repeated conversion | `0` | covered by unit test for repeated conversion |

Required command or test:

```powershell
npm --prefix avmatrix-web run test -- test/unit/graph-adapter.edge-geometry.test.ts
```

If a more specific layout test file is added, record that command instead.

## B2 - No-Auto-Optimizer Benchmark

Date: 2026-05-20

Status: recorded

Record before/after behavior:

| Scenario | Before | After | Expected after |
|---|---:|---:|---:|
| load graph, do not click optimizer, layout starts | `0` | `0` | `0` |
| load graph, do not click optimizer, layout running indicator appears | `0` | `0` | `0` |
| click optimizer once, optimizer invocation count | `0` | `1` | `1` |
| click optimizer once, graph remains interactive | observed pass | observed pass | observed pass |

Use runtime diagnostics and browser/e2e artifacts. Do not substitute unit test pass/fail for runtime measurement if browser behavior is the claim.

## B3 - Interaction And Readability Benchmark

Date: 2026-05-20

Status: recorded

Record final Web interaction observations:

| Interaction | Before | After | Status |
|---|---:|---:|---|
| initial graph usable before optimizer | previous behavior auto-started runtime layout | clustered layout shown without auto optimizer | passed by e2e/load diagnostics |
| node type filter toggle remains responsive | existing behavior | unchanged | covered by graph dashboard e2e retained and focused unit coverage |
| graph-health filter toggle remains responsive | existing behavior | unchanged | covered by `graph.test.ts` and existing dashboard tests retained |
| selected node focus still works | existing behavior | unchanged | covered by `selected-graph-context.test.ts` and `GraphCanvas.selection-performance.test.tsx` |
| edge visibility toggle still works | existing behavior | unchanged | covered by edge visibility/style tests |

Readability acceptance is qualitative but must be backed by screenshot, browser observation, or e2e assertion. The key product benchmark is not "optimizer is faster"; it is "graph is readable before optimizer runs".

## B4 - Final Benchmark Summary

Date: 2026-05-20

Status: recorded

Fill at closure:

- before commit: `441776d docs: tighten clustered layout plan`
- after commit: implementation closure commit created after this ledger update
- graph source: local `Restaurant_manager` index served by `go run .\cmd\avmatrix serve`
- repo path: selected by Web `/api/repos`, project `Restaurant_manager`
- final node count: `78358`
- final relationship count: `130588`
- initial layout starts after load: `0`
- manual optimizer invocations after click: `1`
- deterministic layout test result: `graph-adapter.edge-geometry.test.ts`, `9` tests passed
- browser/e2e result: `server-connect.spec.ts` targeted no-auto/manual optimizer tests, `2` passed
- residual risk: Browser plugin direct in-app verification was unavailable because the required Node REPL control tool was not exposed; Playwright e2e and diagnostic capture covered the behavior.
