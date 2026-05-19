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
- after the graph appears to finish loading in the Web UI, the launcher lifecycle can expire and close the session/backend, causing the UI to show `Server connection lost - reconnecting...`.

The UI must remain language-agnostic. AVmatrix supports many languages, and the dashboard must represent graph components produced from all supported language families.

## Scope Boundary

This plan covers Web UI graph display, filtering, legend, and canvas adapter behavior:

- `avmatrix-web/src/lib/constants.ts`
- `avmatrix-web/src/components/FileTreePanel.tsx`
- `avmatrix-web/src/lib/graph-adapter.ts`
- `avmatrix-web/src/hooks/useSigma.ts`
- graph state and tests that control node/edge visibility
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
- loading and laying out a large graph must not make the app report a server connection loss after the graph is already visible.

Graph relationship scope:

- the dashboard controls graph payload relationship types used by the canvas;
- LadybugDB-only relation constants are not required in the dashboard unless they appear in graph payloads;
- if a future graph payload contains a relationship type unknown to the current display table, the UI must still show it with a safe fallback label/style instead of dropping it.

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
- after graph rendering finishes, `Server connection lost - reconnecting...` can appear, which must be diagnosed instead of dismissed as a transient UI message.

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
- The exact subcause for the missing launcher heartbeat still needs instrumentation: likely browser main-thread starvation from synchronous graph conversion/layout/noverlap, but the primary runtime design defect is already clear: a 5s ping / 15s timeout lifecycle is not robust enough for heavy graph loads.
- The fix must treat heavy repo load as a supported workload, not an exceptional case.

## Acceptance Criteria

- [ ] The left dashboard shows node type controls for all node labels present in the loaded graph, with counts.
- [ ] The left dashboard shows relationship type controls for all relationship types present in the loaded graph, with counts.
- [ ] Node/edge controls are driven by the loaded graph plus a maintained display policy, not by an incomplete fixed list.
- [ ] Color legend reflects the labels and edge types currently visible or available in the dashboard.
- [ ] If community coloring is active, the legend states or shows that symbol colors can be community colors instead of static node-type colors.
- [ ] The graph adapter preserves or explicitly aggregates parallel relationships between the same source and target without losing relationship type information.
- [ ] The edge visibility filter works for every relationship type in the loaded graph.
- [ ] The node visibility filter works for every node label in the loaded graph.
- [ ] Layout hierarchy uses owner/process/route/member relationships where they materially improve graph readability.
- [ ] Tests fail if a graph label or relationship type appears in fixtures but is not representable in the dashboard.
- [ ] Unknown future node labels and relationship types render with safe fallback labels, colors, icons, sizes, and edge styles.
- [ ] Node size scaling is bounded and proportional across structural, metadata, and code nodes; the purple oversized-node screenshot is reproduced or explained, then fixed.
- [ ] The Web UI remains connected after graph load and layout; `Server connection lost - reconnecting...` must not appear during the post-load stability window.
- [ ] Full build passes before the test suite.
- [ ] Test suite includes e2e coverage for node type toggles, edge type toggles, legend behavior, visual-scale bounds, and post-load connection stability.

## Phase 1 - Inventory, Root-Cause Review, and Display Policy

- [x] [P1-A] Inventory current Web UI graph display lists: node filters, edge filters, color legend, icon map, node sizes, edge styles, and layout relationship groups.
- [x] [P1-B] Inventory loaded graph labels and relationship types from `.avmatrix/graph.json` and record counts in the benchmark ledger.
- [ ] [P1-C] Define a display policy for graph categories: structure, code symbols, member/property facts, references/calls/imports, process/community metadata, route/tool/API facts, and unknown/future graph types.
- [ ] [P1-D] Decide which zero-count labels/relationship types should stay hidden by default and how users can reveal them if needed.
- [x] [P1-E] Record evidence showing current missing node labels, missing relationship types, and parallel relationship risk.
- [ ] [P1-F] Build or update a representative Web UI graph fixture that includes every known graph node label and graph payload relationship type, plus one unknown/future label and relationship type for fallback behavior.
- [ ] [P1-G] Document the boundary between graph payload relationship types used by the canvas and LadybugDB-only relationship constants used by query/storage layers.
- [x] [P1-H] Review code paths for the oversized purple-node symptom and record concrete root-cause conclusions/hypotheses.
- [x] [P1-I] Review code paths for the post-load reconnect symptom and record concrete root-cause conclusions/hypotheses.

## Phase 2 - Heavy Graph Runtime Stability

- [x] [P2-A] Add launcher lifecycle instrumentation that records heartbeat age, close-grace expiry reason, and whether the backend was stopped because UI lifecycle exit occurred. Result: `webLifecycleSnapshot` now records heartbeat age, close age, close-grace, exit reason, and backend ownership in launcher logs.
- [ ] [P2-B] Add Web-side instrumentation or test hooks for graph conversion time, layout start/stop time, noverlap duration, heartbeat reconnects, and reconnect-banner state.
- [ ] [P2-C] Reproduce the post-load `Server connection lost - reconnecting...` behavior under controlled conditions with the current large graph and recorded launcher/Web timings.
- [x] [P2-D] Fix the lifecycle design so heavy graph load cannot close the runtime while the page is alive. Result: removed heartbeat-budget auto-shutdown entirely; heartbeat is now instrumentation/liveness evidence only, while explicit page-close signals still close after a short reload grace.
- [x] [P2-E] Ensure a lifecycle timeout does not stop the backend during a known active graph load/layout window. Result: there is no heartbeat timeout path left to stop the backend during graph load/layout.
- [x] [P2-F] Add launcher/backend/Web tests covering heartbeat gaps longer than the old `15s` threshold during active graph load and confirming no forced shutdown. Result: launcher tests cover stale heartbeat, `3h` and `24h` heartbeat gaps, close-grace expiry, and close without prior heartbeat.
- [ ] [P2-G] Add an e2e stability test that loads a large graph, waits through the post-load/layout window, and asserts that the reconnect banner does not appear.
- [ ] [P2-H] Record post-load stability timing, launcher heartbeat age, backend process continuity, and connection-loss count in the benchmark and evidence ledgers.

## Phase 3 - Node Type Dashboard Completeness

- [ ] [P3-A] Replace the hard-coded `FILTERABLE_LABELS` behavior with a graph-aware node type list that includes every node label present in the loaded graph.
- [ ] [P3-B] Add node counts to the Node Types controls so the user can see whether a type exists in the current graph.
- [ ] [P3-C] Add or centralize icon fallback policy for all node labels, including `Property`, `Struct`, `Const`, `Constructor`, `Package`, `Section`, `Process`, `Community`, `Route`, and `Tool`.
- [ ] [P3-D] Keep default visibility useful without hiding types permanently; noisy types may default off but must remain toggleable.
- [ ] [P3-E] Add unit tests proving every node label in a representative graph fixture appears in the dashboard controls.
- [ ] [P3-F] Add fallback tests for unknown/future node labels.

## Phase 4 - Edge Type Dashboard Completeness

- [ ] [P4-A] Replace the narrow `EdgeType` union with relationship visibility support for every relationship type that can appear in the graph payload.
- [ ] [P4-B] Add relationship counts to the Edge Types controls.
- [ ] [P4-C] Add labels, colors, and styles for all existing relationship types, including `USES`, `ACCESSES`, `HAS_PROPERTY`, `HAS_METHOD`, `MEMBER_OF`, `STEP_IN_PROCESS`, `HANDLES_ROUTE`, `FETCHES`, `HANDLES_TOOL`, `ENTRY_POINT_OF`, `WRAPS`, and `QUERIES`.
- [ ] [P4-D] Preserve a clear grouping/order so users can scan relationship types instead of seeing an arbitrary alphabetized list.
- [ ] [P4-E] Add unit tests proving every relationship type in a representative graph fixture appears in the dashboard controls and can be toggled.
- [ ] [P4-F] Add fallback tests for unknown/future relationship types.

## Phase 5 - Canvas Relationship Preservation, Layout, and Visual Scale

- [ ] [P5-A] Fix graphology edge creation so multiple relationship types between the same source and target are preserved or explicitly aggregated with all relationship types retained.
- [ ] [P5-B] Ensure edge filtering works correctly after the parallel-edge fix or aggregation model.
- [ ] [P5-C] Expand layout hierarchy/grouping logic beyond `CONTAINS`, `DEFINES`, and `IMPORTS` where appropriate, including owner/member and process/route relationships.
- [ ] [P5-D] Add graph-adapter tests for parallel relationships such as `CALLS + USES` and `HAS_PROPERTY + ACCESSES` on the same source-target pair.
- [ ] [P5-E] Audit current node size rules, scaled size output, zoom behavior, and community/structural node styling against `reports/problem/screenshot_1779178877.png`.
- [ ] [P5-F] Define and implement a proportional node-size cap after base scaling and reducer multipliers so important structural/highlighted nodes remain larger without producing oversized circles that distort graph readability.
- [ ] [P5-G] Add unit tests for node-size scaling boundaries on small, medium, large, and very large graph sizes.
- [ ] [P5-H] Measure graph adapter conversion time, edge preservation counts, and node-size ratio bounds on the current AVmatrix-GO graph.

## Phase 6 - Legend Accuracy

- [ ] [P6-A] Replace the hard-coded Color Legend list with a legend derived from the labels and relationship types available in the loaded graph.
- [ ] [P6-B] Show counts or availability state in the legend where useful.
- [ ] [P6-C] Distinguish node-type coloring from community coloring so the legend does not mislead users.
- [ ] [P6-D] Add tests covering legend completeness for uncommon labels and relationship types.

## Phase 7 - Validation

- [ ] [P7-A] Run full build before tests.
- [ ] [P7-B] Run focused Web UI unit tests for constants, filter panel, graph adapter, edge visibility, legend behavior, node-size scaling, and connection stability helpers.
- [ ] [P7-C] Add and run an e2e test that loads a graph fixture with uncommon node/edge types, toggles them in the left dashboard, and verifies visible graph behavior.
- [ ] [P7-D] Add and run an e2e visual-scale check using the oversized-node screenshot scenario or an equivalent deterministic graph fixture.
- [ ] [P7-E] Add and run an e2e post-load connection stability check.
- [ ] [P7-F] Run analyze or graph-loading validation as needed to verify the real current graph still loads.
- [ ] [P7-G] Record validation evidence and any performance benchmark results.

Minimum validation commands, unless the implementation discovers a repo-specific replacement:

- `go test ./...`
- `npm --prefix avmatrix-web run build`
- `npm --prefix avmatrix-web run test`
- `npm --prefix avmatrix-web run test:e2e`

## Phase 8 - Closure

- [ ] [P8-A] Update this plan checklist after each completed slice.
- [ ] [P8-B] Update the benchmark ledger for inventory counts, graph-adapter performance, node-size ratio, and post-load connection stability measurements.
- [ ] [P8-C] Update the evidence ledger for commands, tests, screenshots or e2e artifacts, and implementation notes.
- [ ] [P8-D] Commit each completed implementation slice.
- [ ] [P8-E] Final closure: confirm dashboard node/edge/legend completeness, edge preservation, node visual-scale proportionality, connection stability, full build, unit tests, and e2e tests.
