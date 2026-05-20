# AVmatrix Web Filter-Based Clustered Layout And Manual Optimizer Benchmark Ledger

Date: 2026-05-20

Status: reopened - visual island distribution benchmark required

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

Status: recorded, superseded for visual readability by B6

Record final Web interaction observations:

| Interaction | Before | After | Status |
|---|---:|---:|---|
| initial graph usable before optimizer | previous behavior auto-started runtime layout | clustered layout shown without auto optimizer | passed by e2e/load diagnostics |
| node type filter toggle remains responsive | existing behavior | unchanged | covered by graph dashboard e2e retained and focused unit coverage |
| graph-health filter toggle remains responsive | existing behavior | unchanged | covered by `graph.test.ts` and existing dashboard tests retained |
| selected node focus still works | existing behavior | unchanged | covered by `selected-graph-context.test.ts` and `GraphCanvas.selection-performance.test.tsx` |
| edge visibility toggle still works | existing behavior | unchanged | covered by edge visibility/style tests |

Readability acceptance is qualitative but must be backed by screenshot, browser observation, or e2e assertion. The key product benchmark is not "optimizer is faster"; it is "graph is readable before optimizer runs".

Visual reopen note:

- The later user-provided screenshot `reports/problem/screenshot_1779285599.png` proves this B3 readability record was not sufficient. B6 is now the controlling visual benchmark.

## B4 - Initial Implementation Benchmark Summary

Date: 2026-05-20

Status: superseded by corrective review

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

Corrective note:

- This benchmark set proved that `layout.starts` stayed `0` in the targeted diagnostic run, but it did not prove the visual requirement that each node type/filter cluster is a clear separated region with only that node type's color.
- This benchmark set also did not cover the user's later report that optimization appears after render or that product/runtime timeout/reset behavior occurs after the graph is already visible.
- The final benchmark summary must be re-run after corrective implementation.

## B5 - Corrective Benchmark Requirements

Date: 2026-05-20

Status: recorded, superseded for visual shape by B6

The corrective benchmark must record:

| Measurement | Required evidence | Expected result |
|---|---|---|
| Node type color purity | unit test or browser diagnostic grouping nodes by `nodeType` and `color` | each `nodeType` has exactly one primary render color, equal to `getNodeColor(nodeType)` |
| Visual cluster separation | browser/e2e diagnostic or screenshot-backed bounds per node type | each node type/filter occupies a distinct non-overlapping region, allowing intentional spacing only |
| No auto optimizer after render | browser/e2e diagnostics sampled after graph ready and after a post-render wait | `layout.starts=0`, `layout.stops=0`, `manualOptimizerInvocations=0` until user clicks |
| Manual optimizer only | browser/e2e click on `Optimize Layout` | manual invocation count increments only after click |
| Product/runtime timeout ban | code inspection plus browser/e2e evidence where needed | graph load, render, reconnect, and UI state transitions do not rely on product timeout, timer reset, or delayed reset behavior |

The benchmark must not use an elapsed-time budget or timeout as the definition of layout correctness. Test-runner timeouts are allowed only as test guards and must not be cited as product behavior.

Corrective benchmark results, historical before visual reopen:

| Measurement | Evidence | Result |
|---|---|---|
| Node type color purity | `graph-adapter.edge-geometry.test.ts` plus code inspection of `knowledgeGraphToGraphology` | each node uses `getNodeColor(node.label)` as primary render color |
| Visual cluster separation | `graph-adapter.edge-geometry.test.ts` bounding-box coverage | node type visual regions are deterministic and non-overlapping in the tested graph |
| No auto optimizer after render | full e2e `server-connect.spec.ts` | `keeps connection stable after large graph load without automatic layout optimizer` passed |
| Manual optimizer only | full e2e `server-connect.spec.ts` | `invokes manual layout optimizer only after user action` passed |
| Product/runtime timeout ban | `rg -n "setTimeout|clearTimeout|setInterval|clearInterval|timeout|Timeout|TIMEOUT|durationBudget|duration-elapsed|noverlap|lastReason" avmatrix-web\src` | no matches in product/runtime source |
| Process modal latency on large repo | full e2e `server-connect.spec.ts` and `shell-interactions.spec.ts` | process View/modal and lightbulb tests passed after reading process steps from loaded graph |

Final validation benchmark, historical before visual reopen:

| Command | Result |
|---|---|
| `npm --prefix avmatrix-web run build` | passed |
| `go build ./cmd/... ./internal/...` | passed |
| `go build ./...` in `avmatrix-launcher/server-wrapper` | passed |
| `go build ./...` in `avmatrix-launcher/src` | passed |
| `npm --prefix avmatrix-web run test` | `43` files, `336` tests passed |
| `npm --prefix avmatrix-web run test:e2e -- --workers=1` | `42` tests passed in `20.7m` |

Final benchmark interpretation, historical before visual reopen:

- The runtime layout optimizer is no longer a graph-load mechanism.
- The manual optimizer remains a user action and reuses the deterministic clustered layout policy.
- The graph is grouped by existing node type/filter color, not by community color.
- Product/runtime timeout and delayed-reset mechanisms were removed from `avmatrix-web/src`; timeout remains only in tests/e2e runner guards.
- Root `go build ./...` is not an acceptance build command for this repository because it includes intentionally non-buildable analysis fixtures under `avmatrix/test/fixtures`; product Go build coverage is `cmd`, `internal`, and the launcher Go modules.

## B6 - Visual Island Distribution Reopen Benchmark

Date: 2026-05-20

Status: pending implementation and validation

User-provided benchmark evidence:

- Failing current output: `reports/problem/screenshot_1779285599.png`.
- Target visual reference: `reports/problem/aaaa.jpg`.

Benchmark interpretation:

- The B5 visual cluster separation benchmark was too weak. It accepted deterministic non-overlapping bounds, but that still allowed compressed rail/grid clusters that are unreadable.
- "Readable before optimizer" now means separate two-dimensional color islands with visible whitespace, not merely `layout.starts=0` and not merely non-overlapping cluster bounds.
- The target visual model is colored archipelagos on one large circular graph field.
- The sample image is a placement reference only. It must not be interpreted as permission to reduce, hide, filter, prune, or thin graph edges.
- The benchmark must not use an elapsed-time budget or product/runtime timeout as a layout correctness mechanism.

Required new measurements:

| Measurement | Required evidence | Expected result |
|---|---|---|
| Cluster color purity | unit or browser diagnostic grouped by `nodeType` and render color | each visible node type/filter has one primary color from `getNodeColor(nodeType)` |
| Cluster island aspect ratio | browser diagnostic or unit geometry metrics per visible node type | medium and large clusters have bounded aspect ratio and do not collapse into long thin rails |
| Cluster island density | browser diagnostic or unit geometry metrics per visible node type | cluster area scales with node count and capped node diameter; nodes are not stacked into dense blocks |
| Inter-cluster gutters | screenshot-backed bounds or browser diagnostics with node-radius padding | different node type/color islands remain visibly separated |
| Rail/grid regression | screenshot-backed review and automated geometry assertion | output must not resemble `reports/problem/screenshot_1779285599.png` |
| Target-shape comparison | browser screenshot after graph ready on `Restaurant_manager` | layout should resemble colored archipelagos on one large circle, using `reports/problem/aaaa.jpg` as the placement reference |
| Edge preservation | edge count diagnostics plus existing edge visibility tests | node placement changes do not reduce relationship count or hide cross-cluster edges |
| No auto optimizer after visual correction | e2e diagnostics sampled after graph ready | `layout.starts=0`, `layout.stops=0`, `manualOptimizerInvocations=0` until user clicks |
| Manual optimizer only after visual correction | e2e click on `Optimize Layout` | manual invocation count increments only after explicit user action |

Required final benchmark table after implementation:

| Graph source | Node count | Relationship count | Cluster count | Visual result | Optimizer result |
|---|---:|---:|---:|---|---|
| `Restaurant_manager` local index | pending | pending | pending | pending screenshot and metrics | pending diagnostics |

Closure requirement:

- B6 must be filled with actual post-implementation measurements before this plan can be closed again.
