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

Status: completed for loaded-graph controls and representative fixture coverage

Record:

- labels present in loaded graph;
- labels present in representative fixture;
- labels visible in dashboard controls;
- labels with colors/icons/sizes;
- labels with unknown/future fallback coverage;
- visual groups used for sectioning;
- missing labels count.

Date: 2026-05-19

Current graph snapshot after re-analyze:

- path: `.avmatrix/graph.json`
- repo: `E:\AVmatrix-GO`
- nodes: `20,421`
- relationships: `51,111`
- unique node labels present: `16`
- unique relationship types present: `11`

Node dashboard coverage:

| Coverage target | Result |
|---|---:|
| Current graph labels present | `16` |
| Current graph labels displayed as controls | `16` |
| Current graph missing node controls | `0` |
| Generated contract labels in representative fixture | `36` |
| Unknown/future labels in representative fixture | `1` |
| Representative fixture node controls rendered | `37 / 37` |
| Node labels with colors | `36 / 36` generated labels plus fallback |
| Node labels with non-zero sizes | `36 / 36` generated labels plus fallback |

Previously missing current graph labels now displayed with counts:

| Node label | Count |
|---|---:|
| Community | 930 |
| Const | 323 |
| Constructor | 5 |
| Package | 413 |
| Process | 645 |
| Property | 3,106 |
| Section | 930 |
| Struct | 501 |
| TypeAlias | 70 |

Representative fixture proof:

- `FileTreePanel.dashboard-completeness.test.tsx` renders every generated node label plus `FutureNode`;
- the same test clicks every graph-present node type control and verifies `toggleLabelVisibility(label)` receives the exact label string;
- unknown/future label fallback uses safe icon, color, and size.

## B2 - Dashboard Completeness After Edge Type Slice

Status: completed for loaded-graph controls and representative fixture coverage

Record:

- relationship types present in loaded graph;
- relationship types present in representative fixture;
- relationship types visible in dashboard controls;
- relationship types with labels/colors/styles;
- relationship types with unknown/future fallback coverage;
- visual groups used for sectioning;
- missing relationship types count.

Date: 2026-05-19

Relationship dashboard coverage:

| Coverage target | Result |
|---|---:|
| Current graph relationship types present | `11` |
| Current graph relationship types displayed as controls | `11` |
| Current graph missing relationship controls | `0` |
| Generated graph payload relationship types in representative fixture | `22` |
| Unknown/future relationship types in representative fixture | `1` |
| Representative fixture relationship controls rendered | `23 / 23` |
| Relationship types with labels/colors/styles | `22 / 22` generated types plus fallback |

Previously missing current graph relationship types now displayed with counts:

| Relationship type | Count |
|---|---:|
| ACCESSES | 5,031 |
| ENTRY_POINT_OF | 645 |
| HAS_METHOD | 337 |
| HAS_PROPERTY | 2,779 |
| MEMBER_OF | 3,833 |
| STEP_IN_PROCESS | 2,376 |
| USES | 5,156 |

Representative fixture proof:

- `FileTreePanel.dashboard-completeness.test.tsx` renders every generated graph relationship type plus `FUTURE_RELATIONSHIP`;
- the same test clicks every graph-present relationship control and verifies `toggleEdgeVisibility(type)` receives the exact relationship type string;
- unknown/future relationship fallback uses a safe display label and color.

### B2B - Legend Count and Community-Color Coverage

Date: 2026-05-19

Current graph snapshot after re-analyze:

```text
nodes: 20,548
relationships: 51,416
node labels present: 16
relationship types present: 11
community membership relationships: 3,860
community nodes: 912
```

Legend coverage:

| Coverage target | Result |
|---|---:|
| Node type legend entries with counts | `16 / 16` current labels |
| Edge type legend entries with counts | `11 / 11` current relationship types |
| Representative fixture node legend entries with counts | `37 / 37` |
| Representative fixture edge legend entries with counts | `23 / 23` |
| Community color-set legend shown when `MEMBER_OF` exists | yes |
| Community color swatches shown | `6` preview swatches |

## B3 - Graph Adapter Performance and Preservation

Status: completed for parallel-edge preservation and layout hierarchy expansion slices

Date: 2026-05-19

Current graph input:

```text
repo: E:\AVmatrix-GO
nodes: 20,436
relationships: 51,176
```

Measured graph-adapter result:

| Metric | Result |
|---|---:|
| Input nodes | `20,436` |
| Input relationships | `51,176` |
| Renderable relationships | `51,176` |
| Graphology nodes | `20,436` |
| Graphology edges | `51,176` |
| Relationship loss | `0` |
| Parallel source-target pairs preserved | `1,412` |
| Conversion time | `478.37ms` |

Notes:

- The graph adapter now uses graphology `MultiDirectedGraph`, so relationship identity is no longer collapsed by source-target pair.
- Each edge keeps its own `relationType`, preserving edge visibility/filter behavior for parallel relationships.
- Unit coverage includes both `CALLS + USES` and `HAS_PROPERTY + ACCESSES` parallel relationship pairs.

### B3B - Layout Hierarchy Expansion

Date: 2026-05-19

Current graph relationship inventory relevant to layout hierarchy:

| Relationship type | Count | Layout direction |
|---|---:|---|
| CONTAINS | `1,710` | source parent |
| DEFINES | `17,261` | source parent |
| IMPORTS | `3,727` | source parent |
| HAS_METHOD | `337` | source parent |
| HAS_PROPERTY | `2,829` | source parent |
| MEMBER_OF | `3,860` | target parent, low priority |
| STEP_IN_PROCESS | `2,356` | target parent |
| ENTRY_POINT_OF | `640` | target parent |

Layout parent candidates:

| Candidate set | Count |
|---|---:|
| Previously considered hierarchy relationships | `22,698` |
| Added owner/process/member hierarchy relationships present in current graph | `10,022` |
| Total current hierarchy candidates after expansion | `32,720` |

Notes:

- `HANDLES_ROUTE`, `HANDLES_TOOL`, and `WRAPS` are supported by the layout policy but have zero count in the current AVmatrix-GO graph snapshot.
- Process/route/tool grouping uses higher priority than broad file-level `DEFINES`; community membership stays lower priority so it does not override stronger containment/ownership.

## B4 - Visual Scale and Connection Stability

Status: partial; launcher lifecycle auto-shutdown, visual-scale, and post-load connection-stability slices recorded

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

### B4B - Node Visual Scale Bound

Date: 2026-05-19

Baseline issue:

| Metric | Before |
|---|---:|
| Current graph Project size | `10` |
| Current graph Property size | `1.5` |
| Project/Property radius ratio | `6.7x` |
| Project/Property area ratio | `44.4x` |
| Reducer highlight/selection cap | none |

After visual-scale fix:

| Metric | After |
|---|---:|
| Current graph Project size | `3` |
| Current graph Property size | `1` |
| Project/Property radius ratio | `3x` |
| Dense rendered node size cap after reducer multipliers | `3` |

Measured with the current AVmatrix-GO graph:

```json
{
  "inputNodes": 20436,
  "currentLargeGraphProjectSize": 3,
  "currentLargeGraphPropertySize": 1,
  "currentLargeGraphRadiusRatio": 3,
  "renderedSizeCap": 3
}
```

Notes:

- Structural nodes remain larger than leaf/property nodes, but they no longer dominate the graph at the disproportionate scale shown in `reports/problem/screenshot_1779178877.png`.
- Size-boundary unit tests cover graph sizes `100`, `1,500`, `6,000`, `20,421`, and `60,000`.

### B4C - Post-Load Connection Stability

Date: 2026-05-19

Scenario:

```text
frontend: http://127.0.0.1:5228
backend: http://127.0.0.1:4848
backend PID before/after: 524 / 524
repo selected by e2e backend repo list: Restaurant_manager
stability window: 30,000ms after graph conversion and active layout start
old launcher heartbeat shutdown threshold: 15,000ms
```

Focused e2e result:

| Command | Result | Duration |
|---|---|---:|
| `npm --prefix avmatrix-web run test:e2e -- server-connect.spec.ts -g "keeps connection stable after large graph load and layout window" --workers=1 --timeout=120000` | passed, `1 / 1` | `1.1m` test time, `1.3m` total |

Runtime diagnostics captured from the same server/frontend setup:

```json
{
  "repo": "Restaurant_manager",
  "elapsedMs": 62104,
  "graphConversion": {
    "count": 2,
    "lastMs": 1231.4,
    "maxMs": 2446.1,
    "lastNodeCount": 78350,
    "lastRelationshipCount": 130497
  },
  "layout": {
    "starts": 2,
    "stops": 1,
    "isRunning": true,
    "lastDurationBudgetMs": 45000,
    "lastNoverlapMs": 0
  },
  "heartbeat": {
    "connects": 1,
    "reconnects": 0
  },
  "reconnectBanner": {
    "shows": 0,
    "visible": false
  }
}
```

Notes:

- The e2e assertion deliberately waits beyond the old `15s` heartbeat shutdown threshold while the graph layout is active.
- The test does not require ForceAtlas2 to finish on very large graphs; it verifies the actual regression boundary: no backend disconnect and no reconnect banner during heavy post-load/layout work.
- Launcher heartbeat age as a shutdown condition is absent after `B4A`; launcher unit coverage proves stale heartbeat gaps of `3h` and `24h` do not expire the session.

### B4D - Dense Graph 3x Node-Size Ratio Correction

Date: 2026-05-19

Requirement:

- dense graph base node size must satisfy `maxNodeSize / minNodeSize <= 3`;
- `Package` and `Section` must not share the generic structural cap because they visually dominated the canvas in dense graphs;
- all other known node labels must be checked, not only the labels visible in the screenshot.

Measured large-graph browser diagnostics after correction:

```json
{
  "repo": "Restaurant_manager",
  "visualScale": {
    "maxSizeByLabel": {
      "Folder": 3,
      "Route": 2,
      "Package": 1.5,
      "Community": 1,
      "Process": 1.6,
      "File": 2.4,
      "Function": 1.6,
      "Struct": 3,
      "Class": 3,
      "Method": 1.2,
      "Section": 1,
      "Const": 1,
      "Variable": 1,
      "Interface": 2.8,
      "Property": 1,
      "TypeAlias": 1.2,
      "Constructor": 1.6
    },
    "nodeCount": 78350,
    "minNodeSize": 1,
    "maxNodeSize": 3,
    "maxRenderedNodeSizeCap": 3,
    "structuralToLeafRatio": 3
  },
  "graphConversion": {
    "count": 2,
    "lastMs": 2692.6,
    "maxMs": 2692.6,
    "lastNodeCount": 78350,
    "lastRelationshipCount": 130497
  }
}
```

All known node-label dense-size check at `78,350` nodes:

| Label | Base size | Dense scaled size |
|---|---:|---:|
| Class | `8.0` | `3.0` |
| Folder | `10.0` | `3.0` |
| Module | `13.0` | `3.0` |
| Namespace | `13.0` | `3.0` |
| Project | `20.0` | `3.0` |
| Record | `8.0` | `3.0` |
| Struct | `8.0` | `3.0` |
| Interface | `7.0` | `2.8` |
| Trait | `7.0` | `2.8` |
| File | `6.0` | `2.4` |
| Enum | `5.0` | `2.0` |
| Route | `5.0` | `2.0` |
| Tool | `5.0` | `2.0` |
| Union | `5.0` | `2.0` |
| Constructor | `4.0` | `1.6` |
| Function | `4.0` | `1.6` |
| Process | `4.0` | `1.6` |
| Package | `16.0` | `1.5` |
| Delegate | `3.0` | `1.2` |
| Impl | `3.0` | `1.2` |
| Method | `3.0` | `1.2` |
| Template | `3.0` | `1.2` |
| Type | `3.0` | `1.2` |
| TypeAlias | `3.0` | `1.2` |
| Typedef | `3.0` | `1.2` |
| Annotation | `2.0` | `1.0` |
| CodeElement | `2.0` | `1.0` |
| Community | `2.0` | `1.0` |
| Const | `2.0` | `1.0` |
| Decorator | `2.0` | `1.0` |
| Import | `1.5` | `1.0` |
| Macro | `2.0` | `1.0` |
| Property | `2.0` | `1.0` |
| Section | `8.0` | `1.0` |
| Static | `2.0` | `1.0` |
| Variable | `2.0` | `1.0` |

Result:

- measured loaded graph size range: `1.0-3.0`;
- measured dense post-reducer display cap: `3.0`;
- measured loaded graph structural-to-leaf ratio: `3x`;
- all known node labels stay within the dense graph `3x` cap;
- `Package` is deliberately smaller than generic structural nodes at `1.5`;
- `Section` is deliberately leaf-sized at `1.0` because large section clouds were visually disproportionate.

## B5 - Final Benchmark

Status: completed

Date: 2026-05-19

Final AVmatrix-GO analyze:

```text
files: scanned=688 parsed=530 unsupported=158 failed=0
graph: nodes=20,611 relationships=51,507
```

Final dashboard and graph-rendering coverage:

| Metric | Result |
|---|---:|
| Current graph node labels displayed as controls | `16 / 16` |
| Current graph relationship types displayed as controls | `11 / 11` |
| Representative fixture node controls | `37 / 37` |
| Representative fixture relationship controls | `23 / 23` |
| Current graph node legend entries with counts | `16 / 16` |
| Current graph edge legend entries with counts | `11 / 11` |
| Parallel relationships preserved | `51,176 / 51,176` in measured adapter slice |
| Relationship loss in adapter slice | `0` |
| Parallel source-target pairs preserved | `1,412` |
| Adapter conversion time in measured slice | `478.37ms` |

Final visual-scale and large-graph browser diagnostics:

```json
{
  "repo": "Restaurant_manager",
  "graphConversion": {
    "count": 2,
    "lastMs": 2692.6,
    "maxMs": 2692.6,
    "lastNodeCount": 78350,
    "lastRelationshipCount": 130497
  },
  "visualScale": {
    "maxSizeByLabel": {
      "Folder": 3,
      "Route": 2,
      "Package": 1.5,
      "Community": 1,
      "Process": 1.6,
      "File": 2.4,
      "Function": 1.6,
      "Struct": 3,
      "Class": 3,
      "Method": 1.2,
      "Section": 1,
      "Const": 1,
      "Variable": 1,
      "Interface": 2.8,
      "Property": 1,
      "TypeAlias": 1.2,
      "Constructor": 1.6
    },
    "nodeCount": 78350,
    "minNodeSize": 1,
    "maxNodeSize": 3,
    "maxRenderedNodeSizeCap": 3,
    "structuralToLeafRatio": 3
  }
}
```

Final connection-stability result:

| Metric | Result |
|---|---:|
| Stability window after active layout start | `30,000ms` |
| Old launcher heartbeat shutdown threshold | `15,000ms` |
| Heartbeat reconnects during window | `0` |
| Reconnect banner shows during window | `0` |
| Backend PID before/after measured stability run | `524 / 524` |

Final validation summary:

| Command | Result |
|---|---|
| `go build -trimpath -o .tmp\avmatrix.exe .\cmd\avmatrix` | passed |
| `npm --prefix avmatrix-web run build` | passed, built in `24.02s` |
| `npm --prefix avmatrix-web run test` | passed, `41` files / `316` tests, `33.29s` |
| `go test ./cmd/... ./internal/... -count=1` | passed |
| `npm --prefix avmatrix-web run test:e2e -- server-connect.spec.ts -g "Graph Dashboard Controls" --workers=1 --timeout=120000` | passed, `2 / 2`, `1.6m` |
| `npm --prefix avmatrix-web run test:e2e -- server-connect.spec.ts -g "keeps connection stable after large graph load and layout window" --workers=1 --timeout=120000` | passed, `1 / 1`, `1.2m` |

Test stability fixes completed during final validation:

- `internal/session/controller_test.go`: waits for the first chat goroutine to start before asserting cancellation order and protects fake adapter run records with a mutex.
- `internal/embeddings/http_client_test.go`: replaces a 1ms `httptest.Server` timeout race with a deterministic blocking `RoundTripper`, proving timeout errors are not retried.

## B6 - Top Bar Start-Screen Navigation Addendum

Status: pending

Date: 2026-05-19

Benchmark policy for this addendum:

- adding a Back arrow/button beside the `AVmatrix` title is not benchmarkable by itself;
- record validation evidence for visibility, click/keyboard activation, and navigation to `Start-AVmatrix.html`;
- record benchmark data only if the implementation changes startup timing, runtime lifecycle behavior, package size, page-load time, or reconnect/heartbeat behavior.

Planned measurements if lifecycle or startup behavior changes:

| Metric | Result |
|---|---:|
| Back navigation latency to `Start-AVmatrix.html` | pending |
| Reconnect banner shown during intentional Back navigation | pending |
| Backend/runtime lifecycle behavior after Back navigation | pending |

## B7 - Left Dashboard Resizable Width Addendum

Status: pending

Date: 2026-05-19

Benchmark policy for this addendum:

- adding drag-resize behavior for the left dashboard is primarily interaction validation, not a runtime benchmark;
- record validation evidence for drag behavior, min/max width bounds, no layout overlap, and continued dashboard/canvas interactions after resize;
- record benchmark data only if the implementation changes graph render/layout timing, dashboard render cost, page-load time, or post-load stability metrics.

Planned measurements if layout performance changes:

| Metric | Result |
|---|---:|
| Minimum dashboard width after drag | pending |
| Maximum dashboard width after drag | pending |
| Graph/canvas usable width after dashboard expansion | pending |
| Reconnect banner shown during/after dashboard resize | pending |
