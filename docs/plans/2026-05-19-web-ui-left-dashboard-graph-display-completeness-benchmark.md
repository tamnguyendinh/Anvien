# Web UI Left Dashboard Graph Display Completeness Benchmark Ledger

Date: 2026-05-19

Status: active

Companion files:

- Plan: [2026-05-19-web-ui-left-dashboard-graph-display-completeness-plan.md](2026-05-19-web-ui-left-dashboard-graph-display-completeness-plan.md)
- Evidence ledger: [2026-05-19-web-ui-left-dashboard-graph-display-completeness-evidence.md](2026-05-19-web-ui-left-dashboard-graph-display-completeness-evidence.md)

## Benchmark Rules

Record benchmarkable results when they measure product/runtime performance, capacity, package/startup size, graph/DB throughput, or conversion-inventory counts. Build/test/e2e timings are validation evidence unless the slice changes those systems.

For this Web UI plan, benchmarkable measurements include:

- graph inventory counts used by dashboard controls;
- dashboard coverage counts;
- graph adapter conversion time;
- relationship preservation or aggregation counts;
- node visual scale ratios and size-cap behavior;
- post-load connection stability and reconnect-banner counts;
- large-graph render/load capacity where measured.

Inventory measurements must include both:

- the current real graph loaded from `.avmatrix/graph.json`;
- a representative Web UI fixture that covers known and fallback graph display cases, because the current real graph may not contain every supported label or relationship type.

## B0 - Initial Inventory Baseline

Date: 2026-05-19

Graph snapshot:

- path: `.avmatrix/graph.json`
- repo: `E:\AVmatrix-GO`
- nodes: `20,354`
- relationships: `50,980`
- unique node labels present: `16`
- unique relationship types present: `11`

### Node Label Counts

| Node label | Count |
|---|---:|
| Class | 4 |
| Community | 934 |
| Const | 321 |
| Constructor | 5 |
| File | 682 |
| Folder | 112 |
| Function | 3,339 |
| Interface | 98 |
| Method | 804 |
| Package | 413 |
| Process | 644 |
| Property | 3,096 |
| Section | 889 |
| Struct | 500 |
| TypeAlias | 69 |
| Variable | 8,444 |

### Relationship Type Counts

| Relationship type | Count |
|---|---:|
| ACCESSES | 5,024 |
| CALLS | 8,396 |
| CONTAINS | 1,658 |
| DEFINES | 17,093 |
| ENTRY_POINT_OF | 644 |
| HAS_METHOD | 336 |
| HAS_PROPERTY | 2,769 |
| IMPORTS | 3,713 |
| MEMBER_OF | 3,826 |
| STEP_IN_PROCESS | 2,373 |
| USES | 5,148 |

### Current Dashboard Coverage

| Area | Current UI coverage | Present in graph | Missing present items |
|---|---:|---:|---:|
| Node type controls | 7 present labels controllable | 16 labels | 9 labels |
| Edge type controls | 4 present relationship types controllable | 11 types | 7 types |
| Hard-coded node filters | 11 labels total | 36 contract labels | 25 contract labels |
| Hard-coded edge filters | 6 types total | 22 graph relationship types | 16 graph relationship types |

Notes:

- `22 graph relationship types` means relationship types in the graph payload contract used by `GraphRelationship`.
- LadybugDB-only relationship constants are outside dashboard scope unless they appear in graph payloads.
- Current real graph coverage is not a full schema coverage test because many known labels and relationship types have zero count in this repo snapshot.

Missing current graph node labels:

| Missing node label | Count |
|---|---:|
| Community | 934 |
| Const | 321 |
| Constructor | 5 |
| Package | 413 |
| Process | 644 |
| Property | 3,096 |
| Section | 889 |
| Struct | 500 |
| TypeAlias | 69 |

Missing current graph relationship types:

| Missing relationship type | Count |
|---|---:|
| ACCESSES | 5,024 |
| ENTRY_POINT_OF | 644 |
| HAS_METHOD | 336 |
| HAS_PROPERTY | 2,769 |
| MEMBER_OF | 3,826 |
| STEP_IN_PROCESS | 2,373 |
| USES | 5,148 |

### Parallel Relationship Risk

Current graph has `1,421` source-target pairs with more than one relationship type. The current graph adapter check that prevents adding an edge when an edge already exists between the same source and target is therefore not safe unless it intentionally aggregates and exposes all relationship types.

Sample pairs:

| Source-target pair | Relationship types |
|---|---|
| `Class:avmatrix-web/test/unit/use-auto-scroll.test.tsx:ResizeObserverMock -> Property:avmatrix-web/test/unit/use-auto-scroll.test.tsx:ResizeObserverMock.observedElements` | `ACCESSES`, `HAS_PROPERTY` |
| `Const:internal/aicontext/aicontext.go:startMarker -> Struct:internal/aicontext/aicontext.go:baseSkill` | `CALLS`, `USES` |
| `File:avmatrix-web/test/unit/graph-edge-render-style.test.ts -> Function:avmatrix-web/src/lib/graph-edge-render-style.ts:getSelectedContextEdgeSize` | `CALLS`, `USES` |

### Visual Scale Problem Baseline

Status: initial code-derived baseline recorded; rendered screenshot measurement pending

Known artifact:

- screenshot: `reports/problem/screenshot_1779178877.png`
- issue: many purple circular structural/folder-like nodes render far too large relative to surrounding graph nodes.

Record during implementation:

- suspected node label and node id if recoverable;
- rendered size/radius of the oversized node;
- median rendered node size for visible nodes;
- max/median rendered node size ratio;
- zoom/camera ratio when observed;
- node count and graph size at observation time;
- size constants and scaled-size output before/after fix.

Initial code-derived size baseline for the current `20,354` node graph:

| Label | Base size | Scaled size | Selected size | Glow max | Ripple max |
|---|---:|---:|---:|---:|---:|
| Project | 20 | 10.00 | 18.00 | 20.00 | 25.00 |
| Package | 16 | 8.00 | 14.40 | 16.00 | 20.00 |
| Module | 13 | 6.50 | 11.70 | 13.00 | 16.25 |
| Folder | 10 | 5.00 | 9.00 | 10.00 | 12.50 |
| Class | 8 | 4.00 | 7.20 | 8.00 | 10.00 |
| File | 6 | 3.00 | 5.40 | 6.00 | 7.50 |
| Function | 4 | 2.00 | 3.60 | 4.00 | 5.00 |
| Method | 3 | 1.50 | 2.70 | 3.00 | 3.75 |
| Property | 2 | 1.50 | 2.70 | 3.00 | 3.75 |
| Variable | 2 | 1.50 | 2.70 | 3.00 | 3.75 |

Area-ratio note:

- scaled `Folder=5` versus scaled `Property=1.5` is `3.3x` radius and `11.1x` circle area;
- scaled `Project=10` versus scaled `Property=1.5` is `6.7x` radius and `44.4x` circle area;
- selected/glow/ripple multipliers can increase the rendered ratio again.

This is enough to explain the screenshot class of problem: structural/purple nodes can dominate the canvas even before any semantic evidence says they deserve that much visual weight.

### Post-Load Connection Stability Baseline

Status: launcher lifecycle expiry confirmed from logs; missing-heartbeat subcause requires instrumentation

Known symptom:

- after the graph appears to load fully in the Web UI, the banner `Server connection lost - reconnecting...` can appear after a short delay.

Record during implementation:

- graph load time;
- layout start/stop time;
- heartbeat/EventSource reconnect count;
- backend/launcher process continuity;
- browser console errors;
- main-thread long-task observations if measured;
- whether the reconnect banner appears during a defined post-load stability window.

Initial code-path baseline:

| Component | Current behavior | Risk |
|---|---|---|
| Web app heartbeat | `connectHeartbeat` starts only in `exploring` mode | heartbeat starts after graph data handoff, while canvas work begins |
| Backend heartbeat endpoint | `/api/heartbeat` sends SSE `:ok` immediately and every `15s` | EventSource errors if backend/connection drops |
| Launcher UI heartbeat | injected script pings `/__avmatrix_launcher/heartbeat` every `5s` | browser main-thread stall can delay interval/fetch |
| Launcher lifecycle timeout | `15s` without UI heartbeat closes session | heavy graph conversion/layout may exceed timeout budget |
| Graph conversion | graphology conversion runs synchronously on main thread | can starve timers/fetches on large graphs |
| Layout cleanup | `noverlap.assign` runs on main thread after layout stop | can add another stall after graph is already visible |

Observed launcher log evidence:

| Time | Log / artifact | Meaning |
|---|---|---|
| 2026-05-19 15:20:03.638886 | `start root=E:\AVmatrix-GO` | launcher session started |
| 2026-05-19 15:20:03.646305 | `backend pid=13752` | launcher owned backend process |
| 2026-05-19 15:21:30 | `reports/problem/screenshot_1779178877.png` timestamp | graph UI visible before session closed |
| 2026-05-19 15:21:45.648312 | `web ui session closed` | lifecycle monitor fired while UI should still have been usable |

Confirmed chain to validate/fix:

1. Launcher lifecycle monitor expired.
2. `waitForExit` treated the UI session as closed.
3. Because launcher owned backend pid `13752`, `defer stopPID(backend.pid)` could stop the backend.
4. Web `/api/heartbeat` EventSource errors after backend/session loss.
5. App shows `Server connection lost - reconnecting...`.

## B1 - Dashboard Completeness After Node Type Slice

Status: pending

Record:

- labels present in loaded graph;
- labels present in representative fixture;
- labels visible in dashboard controls;
- labels with colors/icons/sizes;
- labels with unknown/future fallback coverage;
- visual groups used for sectioning;
- missing labels count.

## B2 - Dashboard Completeness After Edge Type Slice

Status: pending

Record:

- relationship types present in loaded graph;
- relationship types present in representative fixture;
- relationship types visible in dashboard controls;
- relationship types with labels/colors/styles;
- relationship types with unknown/future fallback coverage;
- visual groups used for sectioning;
- missing relationship types count.

## B3 - Graph Adapter Performance and Preservation

Status: pending

Record:

- graphology conversion time before/after;
- input relationship count;
- rendered or aggregated relationship count;
- parallel relationship pairs preserved;
- relationship type loss count.

## B4 - Visual Scale and Connection Stability

Status: partial; launcher lifecycle auto-shutdown slice recorded, visual scale and post-load reconnect-count validation still pending

Record:

- node max/median size ratio before/after;
- screenshot or deterministic fixture used;
- post-load wait duration;
- reconnect banner count;
- heartbeat continuity result;
- backend/launcher process continuity result.

### B4A - Launcher Lifecycle Budget Removal

Date: 2026-05-19

Measured/validated behavior:

| Metric | Before | After |
|---|---:|---:|
| Launcher heartbeat interval | `5s` | `5s` |
| Heartbeat-missing auto-shutdown budget | `15s` | none |
| Tested stale-heartbeat gap without lifecycle expiry | not supported | `3h` and `24h` |
| Explicit close signal grace | `2s` | `2s` |
| Backend stop caused by heartbeat age alone | possible | removed |

Focused e2e graph-load evidence:

| Command | Result | Duration |
|---|---|---:|
| `npm --prefix avmatrix-web run test:e2e -- server-connect.spec.ts -g "selects a repo from landing and loads graph" --workers=1 --timeout=120000` | passed, `1 / 1` | `25.8s` test time, `31.4s` total |

Notes:

- This slice intentionally does not set a larger heartbeat budget. It removes the heartbeat budget as a shutdown mechanism because any finite budget can fail on a sufficiently large repo/load.
- Heartbeat age is still logged for diagnosis, but it is no longer allowed to close the Web UI runtime.
- Full post-load reconnect-banner count remains pending for `P2-G/P2-H`; this slice only proves the launcher can no longer stop the backend because the page missed heartbeat pings during heavy graph work.

## B5 - Final Benchmark

Status: pending

Record final measurements after all implementation slices:

- dashboard node type completeness;
- dashboard edge type completeness;
- legend completeness;
- graph adapter conversion performance;
- parallel relationship preservation;
- node visual-scale proportionality;
- post-load connection stability;
- large-graph load/render capacity if measured.
