# AVmatrix Web Filter-Based Clustered Layout And Manual Optimizer Plan

Date: 2026-05-20

Status: active

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

## Scope Boundary

Implementation may touch:

- `avmatrix-web/src/lib/graph-adapter.ts` initial node placement policy;
- `avmatrix-web/src/hooks/useSigma.ts` layout lifecycle and manual optimizer behavior;
- `avmatrix-web/src/components/GraphCanvas.tsx` optimizer button labeling and behavior;
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
- changing graph-health filters, edge filters, or analyzer behavior;
- changing graph payload shape unless a later implementation finding proves it is required.

## Design Decision

Use existing Web Node Type filters as the layout clustering source of truth.

```text
Cluster key = node.label
Cluster order = FILTERABLE_LABELS order
Unknown label order = sort by label string, then node id, appended after FILTERABLE_LABELS
Node order inside cluster = filePath -> name -> id
Initial placement = deterministic row-major cluster grid
In-cluster placement = deterministic local grid; columns = ceil(sqrt(clusterNodeCount))
Optimizer = manual-only cleanup
```

No node is moved toward a center because it is "important". No graph connectivity metric decides placement. The goal is a stable, readable grouping that matches controls the user already understands.

## Target Behavior

On graph load:

1. Build one logical visual cluster per filterable node label, using `FILTERABLE_LABELS`.
2. Append graph labels unknown to `FILTERABLE_LABELS` after known labels in stable order.
3. Place nodes inside each cluster using deterministic ordering and a row-major local grid.
4. Render immediately without starting layout optimization.
5. Keep filters compatible: toggling a node type hides/shows that type's existing cluster without causing unrelated clusters to jump.

When the user clicks the optimizer button:

1. Run layout cleanup manually.
2. Preserve the node-label cluster boundary as the primary visual rule.
3. Improve overlap and spacing within clusters and between clusters where practical.
4. Apply the result without turning graph load into a long live animation.
5. Do not run a global layout that mixes different node-label clusters together.
6. Do not reuse the existing global ForceAtlas2 optimizer unchanged unless it is constrained to preserve node-label clusters.

## Non-Goals

- Do not compute "important" nodes.
- Do not rank nodes by connection count.
- Do not infer hubs.
- Do not add center-vs-edge placement logic based on graph topology.
- Do not make the optimizer mandatory for readability.
- Do not use elapsed-time budget as the layout correctness mechanism.

## Implementation Phases

### P0 - Discovery And Guardrails

- [ ] [P0-A] Run AVmatrix refresh and impact checks for the implementation symbols before code edits.
- [ ] [P0-B] Confirm direct callers and tests for `knowledgeGraphToGraphology`, `useSigma`, and `GraphCanvas`.
- [ ] [P0-C] Record current graph load behavior: whether layout starts automatically after `setGraph`.
- [ ] [P0-D] Record current initial placement behavior, including any nondeterministic `Math.random()` use.
- [ ] [P0-E] Define the final cluster order source as `FILTERABLE_LABELS`, not `DEFAULT_VISIBLE_LABELS`.

### P1 - Filter-Based Clustered Initial Layout

- [ ] [P1-A] Replace initial placement with deterministic grouping by `node.label`.
- [ ] [P1-B] Use `FILTERABLE_LABELS` for cluster ordering, with unknown labels appended in stable order.
- [ ] [P1-C] Sort nodes inside each cluster by `filePath`, then `name`, then `id`.
- [ ] [P1-D] Place clusters using a deterministic row-major grid that does not require graph connectivity scoring.
- [ ] [P1-E] Place nodes inside each cluster using a deterministic row-major local grid with `columns = ceil(sqrt(clusterNodeCount))`.
- [ ] [P1-F] Preserve existing node color, size, graph-health metadata, edge conversion, and filter composition.

### P2 - Disable Automatic Optimizer On Load

- [ ] [P2-A] Remove automatic optimizer start from graph load.
- [ ] [P2-B] Ensure loading or replacing a graph renders clustered positions immediately.
- [ ] [P2-C] Ensure no layout start diagnostic is recorded during graph load unless the user explicitly starts optimization.
- [ ] [P2-D] Keep manual stop/cancel behavior correct if an optimizer is already running and the graph is replaced.

### P3 - Manual Optimizer Semantics

- [ ] [P3-A] Rename or clarify the current "Run Layout Again" control as manual layout optimization.
- [ ] [P3-B] Ensure clicking the control is the only normal path that starts optimization.
- [ ] [P3-C] Keep optimizer output subordinate to cluster readability; it may clean spacing but must not become the primary layout source or mix node-label clusters together.
- [ ] [P3-D] Avoid introducing elapsed-time budget as the layout correctness mechanism.
- [ ] [P3-E] Do not reuse the current global ForceAtlas2 optimizer unchanged unless the implementation constrains it to preserve node-label clusters.

### P4 - Tests And Diagnostics

- [ ] [P4-A] Add unit tests proving same input graph produces stable initial positions.
- [ ] [P4-B] Add unit tests proving each node label/filter forms a separate cluster.
- [ ] [P4-C] Add unit tests proving nodes inside a cluster use `filePath -> name -> id` ordering.
- [ ] [P4-D] Add unit or integration coverage proving graph load does not auto-start the optimizer.
- [ ] [P4-E] Add browser/e2e coverage proving the manual optimizer button starts layout only after user action.
- [ ] [P4-F] Validate existing graph filters, graph-health filters, depth filter, selection, and focus behavior still compose with clustered layout.

### P5 - Benchmark And Closure

- [ ] [P5-A] Measure before/after graph load diagnostics on a representative graph.
- [ ] [P5-B] Measure layout start count after initial graph load; expected final value is `0`.
- [ ] [P5-C] Measure manual optimizer start count after clicking the optimizer button; expected final value is `1`.
- [ ] [P5-D] Record render/conversion latency and interaction observations in the benchmark ledger.
- [ ] [P5-E] Run focused Web validation and record results in the evidence ledger.
- [ ] [P5-F] Review the final diff for scope creep against this plan's non-goals.

## Acceptance Criteria

- [ ] Initial graph layout is grouped by existing Node Type filters.
- [ ] There is no new centrality, hub, degree, or importance calculation for placement.
- [ ] Layout is deterministic for the same graph input.
- [ ] Cluster order follows `FILTERABLE_LABELS`, with unknown labels appended in stable order.
- [ ] Unknown labels are ordered by label string, then node id.
- [ ] In-cluster node placement uses row-major local grid with `columns = ceil(sqrt(clusterNodeCount))`.
- [ ] Filter state hides/shows existing clusters and does not create a separate layout policy.
- [ ] Graph load does not automatically start layout optimization.
- [ ] Manual optimizer starts only from explicit user action.
- [ ] Manual optimizer does not mix different node-label clusters together.
- [ ] Existing global ForceAtlas2 optimizer is not reused unchanged unless constrained to preserve node-label clusters.
- [ ] Existing node filters still work and visually map to the same clusters.
- [ ] Existing edge visibility, graph-health filters, depth filters, selection, and focus behavior still work.
- [ ] Unit tests cover deterministic clustered placement.
- [ ] Web/browser or e2e validation covers no-auto-optimizer and manual optimizer trigger.
- [ ] Benchmark ledger records before/after load and optimizer behavior.

## Closure Definition

The plan can be marked complete when the Web graph renders a deterministic node-label clustered layout immediately on load, no automatic optimizer starts on graph load, the user can manually start layout optimization, validation passes, benchmark evidence is recorded, and the final diff has no backend/schema/ranking-system creep.
