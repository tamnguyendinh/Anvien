# Web UI Left Dashboard Graph Display Completeness Plan

Date: 2026-05-19

Status: active

Companion files:

- Benchmark ledger: [2026-05-19-web-ui-left-dashboard-graph-display-completeness-benchmark.md](2026-05-19-web-ui-left-dashboard-graph-display-completeness-benchmark.md)
- Evidence ledger: [2026-05-19-web-ui-left-dashboard-graph-display-completeness-evidence.md](2026-05-19-web-ui-left-dashboard-graph-display-completeness-evidence.md)

## Rules

1. Use AVmatrix for codebase analysis and impact checks while working on implementation slices in this plan.
2. As each task is completed, update the corresponding checklist item immediately.
3. Run a full build before testing; the test suite must include an e2e test.
4. Record benchmark results as each benchmarkable task is completed. Benchmarkable means measured product/runtime performance, capacity, package/startup size, graph/DB throughput, or conversion-inventory counts; build/test/e2e timings are validation evidence unless the slice changes those systems.
5. Record evidence as each evidenced task is completed.
6. For doc-only commits, do not use AVmatrix.
7. After each completed implementation slice, commit the work, then continue until the full plan is complete.

## Problem

The Web UI left dashboard does not yet show or control the graph completely enough for a user to understand what is actually in the loaded graph.

This is not a Go-language problem and not primarily a source-of-truth wording problem. The practical problem is display completeness:

- the Node Types section exposes only a small fixed subset of graph node labels;
- the Edge Types section exposes only a small fixed subset of graph relationship types;
- the Color Legend is hard-coded and can be incomplete or misleading;
- the graph canvas adapter has places where relationship types can be ignored, underused for layout, or collapsed when multiple relationships exist between the same node pair;
- rendered node scale can become visually disproportionate, as shown by `reports/problem/screenshot_1779178877.png`, where many purple structural/folder-like nodes are far too large compared with surrounding nodes;
- after the graph appears to finish loading in the Web UI, the launcher lifecycle can expire and close the session/backend, causing the UI to show `Server connection lost - reconnecting...`;
- after entering the Web UI tool, the top bar has no direct Back arrow/button next to the `AVmatrix` title to return to `Start-AVmatrix.html`, making the runtime flow inconvenient;
- the left dashboard width is fixed, so users cannot drag the panel boundary to expand or shrink it when node/edge lists, legends, repo paths, or graph controls need more or less room.

The UI must remain language-agnostic. AVmatrix supports many languages, and the dashboard must represent graph components produced from all supported language families.

## Scope Boundary

This plan covers Web UI graph display, filtering, legend, and canvas adapter behavior:

- `avmatrix-web/src/lib/constants.ts`
- `avmatrix-web/src/components/Header.tsx`
- `avmatrix-web/src/components/FileTreePanel.tsx`
- `avmatrix-web/src/lib/graph-adapter.ts`
- `avmatrix-web/src/hooks/useSigma.ts`
- `Start-AVmatrix.html`
- graph state and tests that control node/edge visibility
- top bar navigation from the Web UI tool back to the Start AVmatrix entry screen
- left dashboard layout width and pointer-drag resizing behavior
- node size scaling and canvas readability after graph load
- heartbeat/server connection behavior during and after graph load
- e2e coverage for the visible dashboard behavior

This plan does not change extractor semantics, graph fact emission, property ownership rules, or analyze runtime behavior except where test fixtures are needed to validate UI behavior. If the connection-loss investigation proves the issue is in the launcher/backend heartbeat or process lifecycle rather than the Web UI, the fix may touch that runtime layer, but the acceptance gate remains the Web UI staying connected after graph load.

## Completeness Guardrails

The product target is simple: the left dashboard must show the graph components users need to understand the loaded graph. Generated graph contracts and loaded graph payloads are technical guardrails only; they are used to catch missing UI coverage, not to redefine the product problem.

Required behavior:

- every node label present in the loaded graph is displayed as an individual filter option;
- every relationship type present in the loaded graph is displayed as an individual filter option;
- every visualized node type has a color and size;
- every visualized relationship type has a color/style;
- no relationship should disappear only because another relationship already exists between the same source and target;
- visual grouping is allowed only as sectioning around individual options, not as a replacement that hides the real node or relationship type;
- metadata/system graph types may default off, but they must not be invisible by accident;
- structural nodes may be larger than leaf/code nodes, but no node type may render at a scale that visually dominates the graph out of proportion with the rest of the node set;
- loading and laying out a large graph must not make the app report a server connection loss after the graph is already visible;
- the top bar must provide a clear Back arrow/button beside the `AVmatrix` title that returns to `Start-AVmatrix.html`, with accessible naming and keyboard/click support;
- the left dashboard must be resizable by dragging its boundary, with bounded min/max widths, no layout overlap, and usable behavior on dense graph-control content.

Graph relationship scope:

- the dashboard controls graph payload relationship types used by the canvas;
- LadybugDB-only relation constants are not required in the dashboard unless they appear in graph payloads;
- if a future graph payload contains a relationship type unknown to the current display table, the UI must still show it with a safe fallback label/style instead of dropping it.

Dashboard relationship boundary:

- `GraphRelationship.type` / generated `GRAPH_RELATIONSHIP_TYPES` is the Web UI source for Node/Edge controls, canvas edge rendering, and legend entries.
- LadybugDB storage/query constants are backend persistence vocabulary; they only enter the Web dashboard when serialized into the graph payload.
- Tests must cover all generated graph payload relationship types plus unknown/future fallback types. They should not require dashboard controls for LadybugDB-only relations that are never present in Web graph payloads.

## Baseline

Current graph snapshot after analyze:

- graph path: `.avmatrix/graph.json`
- nodes: `20,354`
- relationships: `50,980`
- unique node labels in current graph: `16`
- unique relationship types in current graph: `11`

Current dashboard coverage:

- dashboard node filters: `11` hard-coded labels;
- current graph labels directly controllable by those filters: `7 / 16`;
- current graph labels missing from the dashboard: `Community`, `Const`, `Constructor`, `Package`, `Process`, `Property`, `Section`, `Struct`, `TypeAlias`;
- dashboard edge filters: `6` hard-coded relationship types;
- current graph relationship types directly controllable by those filters: `4 / 11`;
- current graph relationship types missing from the dashboard: `ACCESSES`, `ENTRY_POINT_OF`, `HAS_METHOD`, `HAS_PROPERTY`, `MEMBER_OF`, `STEP_IN_PROCESS`, `USES`;
- current graph has `1,421` source-target pairs with multiple relationship types, so the canvas cannot safely key edges only by source and target.

The current graph is useful evidence, but it is not enough by itself because it only contains `16 / 36` known node labels and `11 / 22` known graph relationship types. This plan therefore also needs a representative Web UI fixture that includes uncommon and currently-zero labels/relationships so dashboard completeness is tested beyond today's graph.

Known visual/runtime problems to include in validation:

- `reports/problem/screenshot_1779178877.png`: many purple circular structural/folder-like nodes are too large relative to other nodes and the overall rendered graph scale;
- after graph rendering finishes, `Server connection lost - reconnecting...` can appear, which must be diagnosed instead of dismissed as a transient UI message;
- the top bar lacks a Start-screen return affordance next to `AVmatrix`, so users cannot conveniently navigate back to `Start-AVmatrix.html` from inside the tool;
- the left dashboard cannot be expanded or narrowed by dragging, which makes long graph-control lists and labels harder to inspect on different screen sizes.

## Codebase Findings Before Implementation

These findings come from the initial codebase review and should guide implementation.

Oversized purple structural-node root cause:

- The screenshot shows many large purple nodes around the graph perimeter with folder-like labels such as `service`, `config`, `app`, `audit`, and `db_test`, not only one selected node.
- `avmatrix-web/src/lib/constants.ts` intentionally uses "dramatic size differences": `Project=20`, `Package=16`, `Module=13`, `Folder=10`, while leaf/member nodes are often `1.5-2`.
- `avmatrix-web/src/lib/graph-adapter.ts` scales the current `20,354` node graph with `baseSize * 0.5`, leaving `Project=10`, `Package=8`, `Module=6.5`, and most variable/property/member nodes at `1.5`.
- Even without selection/highlight, a scaled `Folder=5` has more than `11x` the circle area of a scaled `Function=2` or `Property=1.5` node. `Project=10` has more than `44x` the area of a `Property=1.5` node.
- `avmatrix-web/src/hooks/useSigma.ts` can multiply node sizes again for selected/highlighted/animated states: selected `1.8x`, query highlight `1.6x`, glow animation up to `2.0x`, ripple animation up to `2.5x`.
- There is no final rendered-size cap after base scaling or reducer multipliers. Structural/folder/purple nodes can therefore dominate the canvas out of proportion with the rest of the graph.
- The fix should not make all nodes equal; it should introduce bounded, proportional scale so larger structural nodes remain readable without dominating the canvas.

Post-load reconnect root cause chain:

- The launcher log records `start root=E:\AVmatrix-GO` at `2026-05-19 15:20:03`, screenshot evidence at `15:21:30`, then `web ui session closed` at `15:21:45`.
- In `avmatrix-launcher/src/main.go`, `web ui session closed` is emitted only when `lifecycleDone(...)` fires. That means the launcher lifecycle monitor expired; this is not merely a cosmetic Web banner.
- The lifecycle script pings `/__avmatrix_launcher/heartbeat` every `5s`, but `launcherUITimeout` is only `15s`. A heavy graph load/render/layout path can exceed that budget and make the launcher close the session while the user is still looking at the graph.
- `waitForExit` then returns, and `startRuntime` has `defer stopPID(backend.pid)`, so when the launcher owns the backend process, lifecycle expiry can stop the backend. The Web heartbeat SSE `/api/heartbeat` then errors and `App.tsx` shows `Server connection lost - reconnecting...`.
- The exact subcause for the missing launcher heartbeat still needs instrumentation: likely browser main-thread starvation from synchronous graph conversion/layout/noverlap, but the primary runtime design defect is already clear: using heartbeat age as an auto-shutdown budget is invalid for heavy graph loads.
- The fix must treat heavy repo load as a supported workload, not an exceptional case. It must not merely raise the heartbeat budget; it must remove heartbeat age as a runtime shutdown condition and reserve shutdown for explicit close/user lifecycle signals.

## Acceptance Criteria

- [x] The left dashboard shows node type controls for all node labels present in the loaded graph, with counts.
- [x] The left dashboard shows relationship type controls for all relationship types present in the loaded graph, with counts.
- [x] Node/edge controls are driven by the loaded graph plus a maintained display policy, not by an incomplete fixed list.
- [x] Color legend reflects the labels and edge types currently visible or available in the dashboard.
- [x] If community coloring is active, the legend states or shows that symbol colors can be community colors instead of static node-type colors.
- [x] The graph adapter preserves or explicitly aggregates parallel relationships between the same source and target without losing relationship type information.
- [x] The edge visibility filter works for every relationship type in the loaded graph.
- [x] The node visibility filter works for every node label in the loaded graph.
- [x] Layout hierarchy uses owner/process/route/member relationships where they materially improve graph readability.
- [x] Tests fail if a graph label or relationship type appears in fixtures but is not representable in the dashboard.
- [x] Unknown future node labels and relationship types render with safe fallback labels, colors, icons, sizes, and edge styles.
- [x] Node size scaling is bounded and proportional across structural, metadata, and code nodes; the purple oversized-node screenshot is reproduced or explained, then fixed.
- [x] The Web UI remains connected after graph load and layout; `Server connection lost - reconnecting...` must not appear during the post-load stability window.
- [ ] The top bar shows a Back arrow/button next to the `AVmatrix` title and returns to `Start-AVmatrix.html` from the running Web UI.
- [ ] The Start-screen return behavior has unit or e2e coverage and does not introduce a false reconnect-banner failure during navigation.
- [ ] The left dashboard can be resized by mouse/pointer drag within bounded min/max widths.
- [ ] Resizing the left dashboard does not hide controls, overlap the graph canvas, or break node/edge/legend interactions.
- [x] Full build passes before the test suite.
- [x] Test suite includes e2e coverage for node type toggles, edge type toggles, legend behavior, visual-scale bounds, and post-load connection stability.

## Phase 1 - Inventory, Root-Cause Review, and Display Policy

- [x] [P1-A] Inventory current Web UI graph display lists: node filters, edge filters, color legend, icon map, node sizes, edge styles, and layout relationship groups.
- [x] [P1-B] Inventory loaded graph labels and relationship types from `.avmatrix/graph.json` and record counts in the benchmark ledger.
- [x] [P1-C] Define a display policy for graph categories: structure, code symbols, member/property facts, references/calls/imports, process/community metadata, route/tool/API facts, and unknown/future graph types. Result: generated contract order is the stable known-type order, loaded-graph inventory drives controls, and unknown/future labels/types use safe fallback color/size/label/icon.
- [x] [P1-D] Decide which zero-count labels/relationship types should stay hidden by default and how users can reveal them if needed. Result: loaded graph controls show present labels/types with counts; zero-count known types are covered by fixture/fallback tests and appear when present in a graph payload.
- [x] [P1-E] Record evidence showing current missing node labels, missing relationship types, and parallel relationship risk.
- [x] [P1-F] Build or update a representative Web UI graph fixture that includes every known graph node label and graph payload relationship type, plus one unknown/future label and relationship type for fallback behavior. Result: `FileTreePanel.dashboard-completeness.test.tsx` covers every generated label/type plus `FutureNode` and `FUTURE_RELATIONSHIP`.
- [x] [P1-G] Document the boundary between graph payload relationship types used by the canvas and LadybugDB-only relationship constants used by query/storage layers. Result: plan now defines `GraphRelationship.type` / generated `GRAPH_RELATIONSHIP_TYPES` as the dashboard source, while LadybugDB-only relation constants remain backend storage/query vocabulary unless serialized into graph payloads.
- [x] [P1-H] Review code paths for the oversized purple-node symptom and record concrete root-cause conclusions/hypotheses.
- [x] [P1-I] Review code paths for the post-load reconnect symptom and record concrete root-cause conclusions/hypotheses.

## Phase 2 - Heavy Graph Runtime Stability

- [x] [P2-A] Add launcher lifecycle instrumentation that records heartbeat age, close-grace expiry reason, and whether the backend was stopped because UI lifecycle exit occurred. Result: `webLifecycleSnapshot` now records heartbeat age, close age, close-grace, exit reason, and backend ownership in launcher logs.
- [x] [P2-B] Add Web-side instrumentation or test hooks for graph conversion time, layout start/stop time, noverlap duration, heartbeat reconnects, and reconnect-banner state. Result: `window.__AVMATRIX_WEB_DIAGNOSTICS__` records graph conversion, layout/noverlap, heartbeat, and reconnect-banner counters.
- [x] [P2-C] Reproduce the post-load `Server connection lost - reconnecting...` behavior under controlled conditions with the current large graph and recorded launcher/Web timings. Result: controlled large-graph e2e no longer reproduces the banner after lifecycle budget removal; diagnostics recorded active layout with `0` heartbeat reconnects and `0` banner shows.
- [x] [P2-D] Fix the lifecycle design so heavy graph load cannot close the runtime while the page is alive. Result: removed heartbeat-budget auto-shutdown entirely; heartbeat is now instrumentation/liveness evidence only, while explicit page-close signals still close after a short reload grace.
- [x] [P2-E] Ensure a lifecycle timeout does not stop the backend during a known active graph load/layout window. Result: there is no heartbeat timeout path left to stop the backend during graph load/layout.
- [x] [P2-F] Add launcher/backend/Web tests covering heartbeat gaps longer than the old `15s` threshold during active graph load and confirming no forced shutdown. Result: launcher tests cover stale heartbeat, `3h` and `24h` heartbeat gaps, close-grace expiry, and close without prior heartbeat.
- [x] [P2-G] Add an e2e stability test that loads a large graph, waits through the post-load/layout window, and asserts that the reconnect banner does not appear. Result: `server-connect.spec.ts` now waits for active layout, holds a `30s` post-load stability window, and asserts no reconnect banner or heartbeat reconnect.
- [x] [P2-H] Record post-load stability timing, launcher heartbeat age, backend process continuity, and connection-loss count in the benchmark and evidence ledgers. Result: benchmark/evidence now record Web diagnostics, stability duration, `0` reconnects, `0` banner shows, backend PID continuity, and launcher-heartbeat coverage from the budget-removal slice.

## Phase 3 - Node Type Dashboard Completeness

- [x] [P3-A] Replace the hard-coded `FILTERABLE_LABELS` behavior with a graph-aware node type list that includes every node label present in the loaded graph.
- [x] [P3-B] Add node counts to the Node Types controls so the user can see whether a type exists in the current graph.
- [x] [P3-C] Add or centralize icon fallback policy for all node labels, including `Property`, `Struct`, `Const`, `Constructor`, `Package`, `Section`, `Process`, `Community`, `Route`, and `Tool`.
- [x] [P3-D] Keep default visibility useful without hiding types permanently; noisy types may default off but must remain toggleable.
- [x] [P3-E] Add unit tests proving every node label in a representative graph fixture appears in the dashboard controls.
- [x] [P3-F] Add fallback tests for unknown/future node labels.

## Phase 4 - Edge Type Dashboard Completeness

- [x] [P4-A] Replace the narrow `EdgeType` union with relationship visibility support for every relationship type that can appear in the graph payload.
- [x] [P4-B] Add relationship counts to the Edge Types controls.
- [x] [P4-C] Add labels, colors, and styles for all existing relationship types, including `USES`, `ACCESSES`, `HAS_PROPERTY`, `HAS_METHOD`, `MEMBER_OF`, `STEP_IN_PROCESS`, `HANDLES_ROUTE`, `FETCHES`, `HANDLES_TOOL`, `ENTRY_POINT_OF`, `WRAPS`, and `QUERIES`.
- [x] [P4-D] Preserve a clear grouping/order so users can scan relationship types instead of seeing an arbitrary alphabetized list.
- [x] [P4-E] Add unit tests proving every relationship type in a representative graph fixture appears in the dashboard controls and can be toggled.
- [x] [P4-F] Add fallback tests for unknown/future relationship types.

## Phase 5 - Canvas Relationship Preservation, Layout, and Visual Scale

- [x] [P5-A] Fix graphology edge creation so multiple relationship types between the same source and target are preserved or explicitly aggregated with all relationship types retained. Result: graph adapter now emits a `MultiDirectedGraph` and adds every renderable relationship with a stable edge key.
- [x] [P5-B] Ensure edge filtering works correctly after the parallel-edge fix or aggregation model. Result: every parallel edge retains its own `relationType`, so existing edge visibility logic can filter per relationship type.
- [x] [P5-C] Expand layout hierarchy/grouping logic beyond `CONTAINS`, `DEFINES`, and `IMPORTS` where appropriate, including owner/member and process/route relationships. Result: graph adapter now uses priority-based layout parents for owner, process, route/tool, wrapper, and member relations without letting low-priority community membership override stronger containment/ownership.
- [x] [P5-D] Add graph-adapter tests for parallel relationships such as `CALLS + USES` and `HAS_PROPERTY + ACCESSES` on the same source-target pair. Result: unit coverage proves both pairs are preserved as separate graphology edges.
- [x] [P5-E] Audit current node size rules, scaled size output, zoom behavior, and community/structural node styling against `reports/problem/screenshot_1779178877.png`. Result: current graph baseline had `Project=10` versus `Property=1.5`, a `6.7x` radius and `44.4x` area ratio before reducer multipliers.
- [x] [P5-F] Define and implement a proportional node-size cap after base scaling and reducer multipliers so important structural/highlighted nodes remain larger without producing oversized circles that distort graph readability. Result: dense-graph base scaling caps generic structural nodes at `3.0`, caps `Package` at `1.5`, caps `Section` at `1.0`, and caps dense post-reducer display size at `3.0`; smaller graphs retain the broader reducer cap.
- [x] [P5-G] Add unit tests for node-size scaling boundaries on small, medium, large, and very large graph sizes. Result: tests cover `100`, `1,500`, `6,000`, `20,421`, and `60,000` node graph sizes.
- [x] [P5-H] Measure graph adapter conversion time, edge preservation counts, and node-size ratio bounds on the current AVmatrix-GO graph. Result: `51,176 / 51,176` relationships preserved, `1,412` parallel source-target pairs preserved, conversion `478.37ms`, current graph structural/property radius ratio `3x`.
- [x] [P5-I] Re-check dense-graph size output for every known node label after reducing `Package` and `Section`. Result: all known node labels stay within `1.0-3.0` at `78,350` nodes; `Package=1.5`, `Section=1.0`, and the largest labels are capped at `3.0`.

## Phase 6 - Legend Accuracy

- [x] [P6-A] Replace the hard-coded Color Legend list with a legend derived from the labels and relationship types available in the loaded graph.
- [x] [P6-B] Show counts or availability state in the legend where useful. Result: Node Type and Edge Type legend entries now include loaded-graph counts.
- [x] [P6-C] Distinguish node-type coloring from community coloring so the legend does not mislead users. Result: community coloring appears as a separate community color set entry when `MEMBER_OF` relationships are present.
- [x] [P6-D] Add tests covering legend completeness for uncommon labels and relationship types.

## Phase 7 - Validation

- [x] [P7-A] Run full build before tests. Result: `go build -trimpath -o .tmp\avmatrix.exe .\cmd\avmatrix` and `npm --prefix avmatrix-web run build` passed before validation tests.
- [x] [P7-B] Run focused Web UI unit tests for constants, filter panel, graph adapter, edge visibility, legend behavior, node-size scaling, and connection stability helpers. Result: focused tests passed, followed by full Web unit suite `41` files / `316` tests passed.
- [x] [P7-C] Add and run an e2e test that loads a graph fixture with uncommon node/edge types, toggles them in the left dashboard, and verifies visible graph behavior. Result: e2e toggles uncommon `Property` node and `Accesses` edge controls and verifies `aria-pressed` state plus legend entries.
- [x] [P7-D] Add and run an e2e visual-scale check using the oversized-node screenshot scenario or an equivalent deterministic graph fixture. Result: e2e asserts loaded graph visual-scale diagnostics stay within max-size and ratio bounds.
- [x] [P7-E] Add and run an e2e post-load connection stability check. Result: e2e waits through a `30s` active-layout stability window and confirms `0` heartbeat reconnects and `0` reconnect-banner shows.
- [x] [P7-F] Run analyze or graph-loading validation as needed to verify the real current graph still loads. Result: final analyze produced `20,611` nodes and `51,507` relationships; e2e graph load passed against the running backend/frontend.
- [x] [P7-G] Record validation evidence and any performance benchmark results. Result: final benchmark/evidence ledgers record build, unit, Go, e2e, analyze, and visual-scale diagnostics.

Minimum validation commands, unless the implementation discovers a repo-specific replacement:

- `go test ./...`
- `npm --prefix avmatrix-web run build`
- `npm --prefix avmatrix-web run test`
- `npm --prefix avmatrix-web run test:e2e`

## Phase 8 - Closure

- [x] [P8-A] Update this plan checklist after each completed slice.
- [x] [P8-B] Update the benchmark ledger for inventory counts, graph-adapter performance, node-size ratio, and post-load connection stability measurements.
- [x] [P8-C] Update the evidence ledger for commands, tests, screenshots or e2e artifacts, and implementation notes.
- [x] [P8-D] Commit each completed implementation slice.
- [x] [P8-E] Final closure: confirm dashboard node/edge/legend completeness, edge preservation, node visual-scale proportionality, connection stability, full build, unit tests, and e2e tests.

## Phase 9 - Top Bar Start-Screen Navigation

- [ ] [P9-A] Inspect `avmatrix-web/src/components/Header.tsx`, `avmatrix-web/src/App.tsx`, launcher path handling, and `Start-AVmatrix.html` to confirm the correct navigation target in both dev and packaged launcher modes.
- [ ] [P9-B] Add an icon-first Back arrow/button beside the `AVmatrix` top bar title with an accessible label and keyboard/click activation.
- [ ] [P9-C] Implement the return flow to `Start-AVmatrix.html` without showing an internal reconnect-banner error during the intentional navigation transition.
- [ ] [P9-D] Add unit and/or e2e coverage proving the Back control is visible, activatable, and reaches the Start AVmatrix screen.
- [ ] [P9-E] Run a full build before tests, including the required e2e test, and record evidence. Record benchmark data only if this slice changes startup/runtime timing, package size, or lifecycle behavior.
- [ ] [P9-F] Commit the implementation slice after the checklist, benchmark ledger, and evidence ledger are updated.

## Phase 10 - Left Dashboard Resizable Width

- [ ] [P10-A] Inspect the app shell layout, `FileTreePanel` container, graph canvas sizing, and responsive CSS to identify where the left dashboard width is currently fixed.
- [ ] [P10-B] Add a visible drag handle on the right edge of the left dashboard and support mouse/pointer dragging to resize the panel.
- [ ] [P10-C] Define and enforce min/max dashboard widths so the panel can expand for dense controls and shrink without breaking the graph canvas or hiding essential controls.
- [ ] [P10-D] Ensure resize behavior does not interfere with node type toggles, edge type toggles, legend scrolling, graph canvas pointer interactions, or post-load layout stability.
- [ ] [P10-E] Add unit and/or e2e coverage for drag resizing, width bounds, and continued dashboard interaction after resize.
- [ ] [P10-F] Run a full build before tests, including the required e2e test, and record evidence. Record benchmark data only if resizing changes render/load performance or layout stability metrics.
- [ ] [P10-G] Commit the implementation slice after the checklist, benchmark ledger, and evidence ledger are updated.
