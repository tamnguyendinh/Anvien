# AVmatrix Web Filter-Based Clustered Layout And Manual Optimizer Plan

Date: 2026-05-20

Status: complete - corrective implementation validated

Companion files:

- Benchmark ledger: [2026-05-20-avmatrix-web-filter-based-clustered-layout-manual-optimizer-benchmark.md](2026-05-20-avmatrix-web-filter-based-clustered-layout-manual-optimizer-benchmark.md)
- Evidence ledger: [2026-05-20-avmatrix-web-filter-based-clustered-layout-manual-optimizer-evidence.md](2026-05-20-avmatrix-web-filter-based-clustered-layout-manual-optimizer-evidence.md)

## Rules

1. Use AVmatrix for codebase analysis and impact checks while working on implementation slices in this plan.
2. As each task is completed, update the corresponding checklist item immediately.
3. Run a full applicable Web validation before closing the plan; include unit coverage for layout policy and a browser/e2e check for graph load behavior.
4. Record benchmark results as each benchmarkable task is completed. Benchmarkable means measured Web graph load behavior, layout start/stop counts, render/conversion latency, optimizer latency, memory, and graph interaction latency; build/test timings are validation evidence unless the slice changes those systems.
5. Record evidence as each evidenced task is completed.
6. For doc-only commits, do not use AVmatrix.
7. After each completed implementation slice, commit the work, then continue until the full plan is complete.

## Problem

The Web graph currently relies on runtime layout optimization to make the graph visually usable. This makes the graph harder to reason about because nodes move after load, and it makes large graphs prone to visible lag or expensive post-load work.

The desired product behavior is simpler:

- graph load should immediately show a clear clustered layout;
- clustering should use the Node Type filters already present in the Web UI;
- each node type/filter should form its own visual cluster;
- layout optimization should not run automatically after graph load;
- optimization should be an optional manual cleanup action triggered by the user.

The purpose is readability, not graph-theory ranking. This plan must not introduce hub detection, degree scoring, centrality, semantic importance, or any other new ranking system just to decide where nodes belong.

## Corrective Finding

The first implementation did not fully satisfy the user's intended meaning of "cluster".

The intended meaning is visual and filter-based:

- one Node Type filter equals one visual cluster;
- one visual cluster must occupy one clearly separated region of the canvas;
- one visual cluster must use that node type's own color;
- different node type colors must not be mixed inside the same cluster;
- a cluster is not just "roughly grouped coordinates" if the UI still looks like mixed colors scattered across the canvas.

The earlier implementation grouped positions by node label, but it still allowed community-based render colors for some node labels. That means a `Function`, `Class`, `Method`, or similar filter cluster could show multiple colors. This violates the product intent because "yellow is yellow, blue is blue"; color is part of how the user recognizes each filter cluster.

The earlier implementation also used a dynamic cluster grid sized from cluster dimensions. For large graphs this can still read as spread-out blocks rather than clear named/type-colored regions. The corrected implementation must make cluster separation explicit and deterministic enough that each type appears as its own zone.

The user also reported layout optimization appearing after render and runtime reset behavior after a while. The implementation must verify that no normal graph-load path calls the optimizer.

## Separate Work Tracks

These are separate problems and must not be merged into one explanation or workaround:

1. **Filter/color clustering.** Use existing Node Type filters and existing node type colors. Put each node type/color into its own clearly separated canvas region.
2. **Manual-only optimizer.** Layout optimization must never run automatically. The only allowed product path is the explicit Web UI optimizer button.
3. **Product timeout ban.** Product/runtime code must not use timeout, timer reset, or delayed reset as an operating mechanism. This rule is independent of repository size. Timeout is acceptable only as a test/runner guard when a test needs a bounded failure mode.

## Scope Boundary

Implementation may touch:

- `avmatrix-web/src/lib/graph-adapter.ts` initial node placement policy;
- `avmatrix-web/src/hooks/useSigma.ts` layout lifecycle and manual optimizer behavior;
- `avmatrix-web/src/components/GraphCanvas.tsx` optimizer button labeling and behavior;
- Web runtime/load code only where needed to remove product timeout/reset behavior;
- Web graph constants only if existing filter order needs a direct exported helper;
- unit tests for deterministic clustered placement and no auto optimizer;
- e2e/browser checks for initial clustered rendering and manual optimizer triggering;
- runtime diagnostics if needed to verify no automatic optimizer starts on graph load.

Out of scope:

- backend graph schema changes;
- new node labels or relationship types;
- new cluster taxonomy independent of existing Node Type filters;
- degree, centrality, PageRank, hub scoring, or semantic importance scoring;
- automatic layout optimization after graph load;
- using elapsed-time budget as the layout correctness mechanism;
- using product/runtime timeout as the graph-load, layout, reconnect, reset, or lag-control mechanism;
- changing graph-health filters, edge filters, or analyzer behavior;
- changing graph payload shape unless a later implementation finding proves it is required.

## Design Decision

Use existing Web Node Type filters as the layout clustering source of truth.

```text
Cluster key = node.label
Cluster order = FILTERABLE_LABELS order
Unknown label order = sort unknown labels by label string, appended after FILTERABLE_LABELS
Node order inside cluster = filePath -> name -> id
Initial placement = deterministic separated regions by cluster
In-cluster placement = deterministic local grid; columns = ceil(sqrt(clusterNodeCount))
Optimizer = manual-only cleanup
Render color = getNodeColor(node.label), not community color
```

No node is moved toward a center because it is "important". No graph connectivity metric decides placement. The goal is a stable, readable grouping that matches controls the user already understands.

Community membership may remain as metadata, but it must not override the primary node type color in the main graph canvas while this filter-based clustering mode is active.

## Target Behavior

On graph load:

1. Build one logical visual cluster per node label present in the graph.
2. Order known labels with `FILTERABLE_LABELS`; append graph labels unknown to `FILTERABLE_LABELS` sorted by label string.
3. Place nodes inside each cluster using deterministic ordering and a row-major local grid.
4. Place each cluster into a clear separated canvas region so different node type colors do not appear interleaved.
5. Render each node with its node type/filter color.
6. Render immediately without starting layout optimization.
7. Keep filters compatible: toggling a node type hides/shows that type's existing cluster without causing unrelated clusters to jump.

When the user clicks the optimizer button:

1. Run layout cleanup manually.
2. Preserve the node-label cluster boundary as the primary visual rule.
3. Improve overlap and spacing within clusters and between clusters where practical.
4. Apply the result without turning graph load into a long live animation.
5. Do not run a global layout that mixes different node-label clusters together.
6. Do not reuse the existing global ForceAtlas2 optimizer unchanged unless it is constrained to preserve node-label clusters.

No other normal path may run optimization after graph render. Camera fit, graph replacement, progress updates, heartbeat reconnects, and repository loading must not trigger the layout optimizer.

## Non-Goals

- Do not compute "important" nodes.
- Do not rank nodes by connection count.
- Do not infer hubs.
- Do not add center-vs-edge placement logic based on graph topology.
- Do not make the optimizer mandatory for readability.
- Do not use elapsed-time budget as the layout correctness mechanism.
- Do not use request timeout, timer reset, or delayed UI reset as a product/runtime mechanism. This ban is not conditional on repository size.
- Do not remove test-runner timeouts that exist only to bound tests or e2e execution.
- Do not override node type color with community color in the main graph canvas.

## Implementation Phases

### P0 - Discovery And Guardrails

- [x] [P0-A] Run AVmatrix refresh and impact checks for the implementation symbols before code edits.
- [x] [P0-B] Confirm direct callers and tests for `knowledgeGraphToGraphology`, `useSigma`, and `GraphCanvas`.
- [x] [P0-C] Record current graph load behavior: whether layout starts automatically after `setGraph`.
- [x] [P0-D] Record current initial placement behavior, including any nondeterministic `Math.random()` use.
- [x] [P0-E] Define the final cluster order source as `FILTERABLE_LABELS`, not `DEFAULT_VISIBLE_LABELS`.

### P1 - Filter-Based Clustered Initial Layout

- [x] [P1-A] Replace initial placement with deterministic grouping by `node.label`.
- [x] [P1-B] Use `FILTERABLE_LABELS` for known cluster ordering, with unknown labels appended by label string.
- [x] [P1-C] Sort nodes inside each cluster by `filePath`, then `name`, then `id`.
- [x] [P1-D] Place clusters using deterministic separated regions that read as one clear region per node type/filter, not as scattered mixed-color blocks.
- [x] [P1-E] Place nodes inside each cluster using a deterministic row-major local grid with `columns = ceil(sqrt(clusterNodeCount))`.
- [x] [P1-F] Preserve size, graph-health metadata, edge conversion, and filter composition while forcing primary render color to node type/filter color.

### P2 - Disable Automatic Optimizer On Load

- [x] [P2-A] Remove automatic optimizer start from graph load.
- [x] [P2-B] Ensure loading or replacing a graph renders clustered positions immediately.
- [x] [P2-C] Ensure no layout start diagnostic is recorded during graph load unless the user explicitly starts optimization.
- [x] [P2-D] Keep manual stop/cancel behavior correct if an optimizer is already running and the graph is replaced.
- [x] [P2-E] Audit every graph-load/render path to prove no layout optimizer can start after render unless the user clicked the optimizer control.
- [x] [P2-F] Audit product/runtime code for timeout, timer reset, and delayed UI reset behavior; remove product-path usage and keep timeout only in tests/runner guards.

### P3 - Manual Optimizer Semantics

- [x] [P3-A] Rename or clarify the current "Run Layout Again" control as manual layout optimization.
- [x] [P3-B] Ensure clicking the control is the only normal path that starts optimization.
- [x] [P3-C] Keep optimizer output subordinate to cluster readability; it may clean spacing but must not become the primary layout source or mix node-label clusters together.
- [x] [P3-D] Avoid introducing elapsed-time budget as the layout correctness mechanism.
- [x] [P3-E] Do not reuse the current global ForceAtlas2 optimizer unchanged unless the implementation constrains it to preserve node-label clusters.

### P4 - Tests And Diagnostics

- [x] [P4-A] Add unit tests proving same input graph produces stable initial positions.
- [x] [P4-B] Add unit tests proving each node label/filter forms a separate visual region.
- [x] [P4-C] Add unit tests proving nodes inside a cluster use `filePath -> name -> id` ordering.
- [x] [P4-D] Add unit or integration coverage proving graph load does not auto-start the optimizer.
- [x] [P4-E] Add browser/e2e coverage proving the manual optimizer is invoked only after user action.
- [x] [P4-F] Validate existing graph filters, graph-health filters, depth filter, selection, and focus behavior still compose with clustered layout.
- [x] [P4-G] Add unit coverage proving nodes in a single node type/filter cluster all use that node type's render color.
- [x] [P4-H] Add e2e or diagnostic coverage proving no optimizer starts after graph render without a user click.
- [x] [P4-I] Add regression evidence proving product/runtime graph load, render, reconnect, and UI state transitions do not rely on timeout or delayed reset behavior.

### P5 - Benchmark And Closure

- [x] [P5-A] Measure before/after graph load diagnostics on a representative graph.
- [x] [P5-B] Measure layout start count after initial graph load; expected final value is `0`.
- [x] [P5-C] Measure manual optimizer invocation count after clicking the optimizer button; expected final value is `1`.
- [x] [P5-D] Record render/conversion latency and interaction observations in the benchmark ledger.
- [x] [P5-E] Run focused Web validation and record results in the evidence ledger.
- [x] [P5-F] Review the final diff for scope creep against this plan's non-goals after corrective implementation.

## Acceptance Criteria

- [x] Initial graph layout is grouped by existing Node Type filters as clear separated visual regions.
- [x] There is no new centrality, hub, degree, or importance calculation for placement.
- [x] Layout is deterministic for the same graph input.
- [x] Cluster order follows `FILTERABLE_LABELS` for known labels, with unknown labels appended by label string.
- [x] Only labels present in the graph create rendered clusters; `FILTERABLE_LABELS` is the ordering source, not a requirement to render empty clusters.
- [x] In-cluster node placement uses row-major local grid with `columns = ceil(sqrt(clusterNodeCount))`.
- [x] Each node type/filter cluster uses only that node type's own render color.
- [x] Community color does not override node type/filter color in the main graph canvas.
- [x] Filter state hides/shows existing clusters and does not create a separate layout policy.
- [x] Graph load does not automatically start layout optimization, including after render/camera fit/progress completion.
- [x] Manual optimizer starts only from explicit user action.
- [x] Manual optimizer does not mix different node-label clusters together.
- [x] Existing global ForceAtlas2 optimizer is not reused unchanged unless constrained to preserve node-label clusters.
- [x] Existing node filters still work and visually map to the same clusters.
- [x] Existing edge visibility, graph-health filters, depth filters, selection, and focus behavior still work.
- [x] Unit tests cover deterministic clustered placement.
- [x] Unit/e2e tests cover one-color-per-filter-cluster behavior.
- [x] Web/browser or e2e validation covers no-auto-optimizer after render and manual optimizer trigger.
- [x] Benchmark ledger records before/after load and optimizer behavior after corrective implementation.

## Closure Definition

The plan can be marked complete when the Web graph renders a deterministic node-label clustered layout immediately on load, no automatic optimizer starts on graph load, the user can manually start layout optimization, validation passes, benchmark evidence is recorded, and the final diff has no backend/schema/ranking-system creep.
