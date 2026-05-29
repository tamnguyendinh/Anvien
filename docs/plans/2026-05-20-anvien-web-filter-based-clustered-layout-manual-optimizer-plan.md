# Anvien Web Filter-Based Clustered Layout And Manual Optimizer Plan

Date: 2026-05-20

Status: completed - visual island and documentation center validation recorded

Companion files:

- Benchmark ledger: [2026-05-20-anvien-web-filter-based-clustered-layout-manual-optimizer-benchmark.md](2026-05-20-anvien-web-filter-based-clustered-layout-manual-optimizer-benchmark.md)
- Evidence ledger: [2026-05-20-anvien-web-filter-based-clustered-layout-manual-optimizer-evidence.md](2026-05-20-anvien-web-filter-based-clustered-layout-manual-optimizer-evidence.md)

## Rules

1. Use Anvien for codebase analysis and impact checks while working on implementation slices in this plan.
2. As each task is completed, update the corresponding checklist item immediately.
3. Run a full applicable Web validation before closing the plan; include unit coverage for layout policy and a browser/e2e check for graph load behavior.
4. Record benchmark results as each benchmarkable task is completed. Benchmarkable means measured Web graph load behavior, layout start/stop counts, render/conversion latency, optimizer latency, memory, and graph interaction latency; build/test timings are validation evidence unless the slice changes those systems.
5. Record evidence as each evidenced task is completed.
6. For doc-only commits, do not use Anvien.
7. After each completed implementation slice, commit the work, then continue until the full plan is complete.

## Problem

The Web graph currently relies on runtime layout optimization to make the graph visually usable. This makes the graph harder to reason about because nodes move after load, and it makes large graphs prone to visible lag or expensive post-load work.

The desired product behavior is simpler:

- graph load should immediately show a clear clustered layout;
- clustering should use the Node Type filters already present in the Web UI;
- each node type/filter should form its own visual cluster;
- documentation-system files should form one dedicated Documentation node/filter with its own color;
- the Documentation cluster should sit at the center of the large circular graph field;
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

## Visual Reopen Finding

The corrective implementation is not visually acceptable based on the user-provided screenshot:

- Bad current result: `reports/problem/screenshot_1779285599.png`.
- Target visual reference: `reports/problem/aaaa.jpg`.

The bad result shows compressed rail/grid-like node blocks. Nodes are packed into long straight bands, edges become thick strips, and the clusters do not read as usable regions. This fails the intended meaning of "cluster" even if the positions are deterministic and even if the cluster bounding boxes do not overlap.

The target reference shows separate color islands. The precise target model is: place the clusters as colored archipelagos on one large circle.

- each node type/filter color occupies its own area;
- each area has real two-dimensional spread, not a line, rail, or marching row;
- large clusters receive larger island areas;
- smaller clusters sit around the larger islands with visible whitespace;
- the sample is only a reference for node placement style, not a reference for reducing graph connectivity.

The implementation must not copy the sample by reducing, hiding, filtering, pruning, thinning, or reweighting edges. Edge count, edge visibility rules, and graph relationships must remain governed by the existing graph/filter behavior. The lesson from the sample is how to distribute node clusters as readable islands.

The previous row-major local grid acceptance is superseded. It is now explicitly rejected for medium and large clusters because it can create rigid rows, flat bands, and unreadable dense blocks.

## Documentation Center Finding

The documentation system must be represented as its own visible Web filter/node type.

This is a separate display classification requirement, not a new graph-ranking rule:

- documentation files and documentation-system nodes form one `Documentation` filter;
- `Documentation` uses its own dedicated color, separate from code node type colors;
- the `Documentation` island is placed at the center of the large circular graph field;
- all other filter/color islands are placed around the documentation center;
- raw graph labels, relationships, edge counts, edge visibility rules, and analyzer behavior remain unchanged unless a later implementation finding proves a narrow payload change is required.

Documentation classification should use simple existing node facts such as file path, extension, node name, and existing metadata. It must not use connection count, centrality, hub detection, semantic importance, or graph topology to decide whether a node belongs in the center.

## Separate Work Tracks

These are separate problems and must not be merged into one explanation or workaround:

1. **Filter/color clustering.** Use existing Node Type filters and existing node type colors. Put each node type/color into its own clearly separated canvas region.
2. **Manual-only optimizer.** Layout optimization must never run automatically. The only allowed product path is the explicit Web UI optimizer button.
3. **Product timeout ban.** Product/runtime code must not use timeout, timer reset, or delayed reset as an operating mechanism. This rule is independent of repository size. Timeout is acceptable only as a test/runner guard when a test needs a bounded failure mode.
4. **Documentation center filter.** Documentation files must become one dedicated Web filter/node type with its own color and a center island. This is separate from the outer island placement work.

## Scope Boundary

Implementation may touch:

- `anvien-web/src/lib/graph-adapter.ts` initial node placement policy;
- `anvien-web/src/hooks/useSigma.ts` layout lifecycle and manual optimizer behavior;
- `anvien-web/src/components/GraphCanvas.tsx` optimizer button labeling and behavior;
- Web runtime/load code only where needed to remove product timeout/reset behavior;
- Web graph constants and display classification only where needed to add the Documentation filter/color;
- unit tests for deterministic clustered placement and no auto optimizer;
- e2e/browser checks for initial clustered rendering and manual optimizer triggering;
- runtime diagnostics if needed to verify no automatic optimizer starts on graph load.

Out of scope:

- backend graph schema changes;
- new backend graph labels or relationship types unrelated to the explicit Documentation display filter;
- broad new cluster taxonomy independent of existing Node Type filters and the explicit Documentation display filter;
- degree, centrality, PageRank, hub scoring, or semantic importance scoring;
- automatic layout optimization after graph load;
- using elapsed-time budget as the layout correctness mechanism;
- using product/runtime timeout as the graph-load, layout, reconnect, reset, or lag-control mechanism;
- changing graph-health filters, edge filters, or analyzer behavior;
- changing graph payload shape unless a later implementation finding proves it is required.

## Design Decision

Use existing Web Node Type filters as the layout clustering source of truth, with one explicit Documentation display-filter exception.

```text
Raw graph label = node.label
Display cluster key = Documentation for documentation-classified nodes, otherwise node.label
Documentation cluster = fixed center island
Outer cluster order = FILTERABLE_LABELS order
Unknown label order = sort unknown labels by label string, appended after FILTERABLE_LABELS
Node order inside cluster = filePath -> name -> id
Initial placement = Documentation at center; other clusters as deterministic separated islands on one large circle
In-cluster placement = deterministic two-dimensional island/cloud per cluster
Optimizer = manual-only cleanup
Render color = getNodeColor(display cluster key), not community color
```

No node is moved toward a center because it is "important". No graph connectivity metric decides placement. The goal is a stable, readable grouping that matches controls the user already understands.

The in-cluster island placement must use only already-available facts: display filter, raw node label, stable node order, visible node count, node radius, and configured spacing. Do not add centrality, degree, hub, or connectivity scoring. A deterministic spiral/ring/low-discrepancy placement is acceptable because it creates a two-dimensional island without inventing graph importance.

Cluster macro-placement must give each filter/color its own region with visible gutters. The Documentation island occupies the center. Large non-documentation clusters get proportionally larger outer island regions; smaller non-documentation clusters are distributed into separate peripheral regions. The result should resemble colored archipelagos distributed across one large circular graph field, not a row of packed containers.

Community membership may remain as metadata, but it must not override the primary node type color in the main graph canvas while this filter-based clustering mode is active.

## Target Behavior

On graph load:

1. Build one logical visual cluster per display filter present in the graph.
2. Classify documentation-system files into the dedicated `Documentation` display filter.
3. For non-documentation nodes, use the existing node label as the display filter.
4. Order known outer labels with `FILTERABLE_LABELS`; append graph labels unknown to `FILTERABLE_LABELS` sorted by label string.
5. Place the Documentation cluster at the center of the large circular graph field.
6. Place nodes inside each cluster using deterministic ordering and a two-dimensional island/cloud distribution.
7. Place non-documentation clusters into clear separated outer regions so different node type colors do not appear interleaved or compressed into adjacent rails.
8. Render each node with its display filter color.
9. Render immediately without starting layout optimization.
10. Keep filters compatible: toggling a display filter hides/shows that filter's existing cluster without causing unrelated clusters to jump.

Cluster geometry requirements:

- the whole graph should read as one large circular field containing multiple colored archipelagos;
- the Documentation island should be visibly centered inside that circular field;
- medium and large clusters must have meaningful width and height;
- no medium or large cluster may collapse into a line, rail, single row, or extreme rectangle;
- cluster area must scale with node count and capped node diameter so nodes are not visually stacked into dense blocks;
- visible gutters must exist between different display-filter/color islands;
- cluster layout correctness must be verified by shape/density metrics or screenshot-backed diagnostics, not only by non-overlapping bounding boxes.

When the user clicks the optimizer button:

1. Run layout cleanup manually.
2. Preserve the display-filter cluster boundary as the primary visual rule.
3. Improve overlap and spacing within clusters and between clusters where practical.
4. Apply the result without turning graph load into a long live animation.
5. Do not run a global layout that mixes different display-filter clusters together.
6. Do not reuse the existing global ForceAtlas2 optimizer unchanged unless it is constrained to preserve display-filter clusters.

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
- Do not use a row-major grid as the final in-cluster placement for medium or large clusters.
- Do not arrange clusters as long thin rows, rails, marching lines, or packed rectangular blocks.
- Do not rewrite backend analyzer schema just to create the Documentation display filter unless implementation evidence proves it is required.
- Do not classify documentation by graph topology, connection count, or importance.

## Implementation Phases

P0-P5 record the already-completed historical implementation. After the visual reopen, they are not sufficient for closure. P6-P10 are the required corrective work to close this plan again.

### P0 - Discovery And Guardrails

- [x] [P0-A] Run Anvien refresh and impact checks for the implementation symbols before code edits.
- [x] [P0-B] Confirm direct callers and tests for `knowledgeGraphToGraphology`, `useSigma`, and `GraphCanvas`.
- [x] [P0-C] Record current graph load behavior: whether layout starts automatically after `setGraph`.
- [x] [P0-D] Record current initial placement behavior, including any nondeterministic `Math.random()` use.
- [x] [P0-E] Define the final cluster order source as `FILTERABLE_LABELS`, not `DEFAULT_VISIBLE_LABELS`.

### P1 - Filter-Based Clustered Initial Layout

- [x] [P1-A] Replace initial placement with deterministic grouping by `node.label`.
- [x] [P1-B] Use `FILTERABLE_LABELS` for known cluster ordering, with unknown labels appended by label string.
- [x] [P1-C] Sort nodes inside each cluster by `filePath`, then `name`, then `id`.
- [x] [P1-D] Place clusters using deterministic separated regions that read as one clear region per node type/filter, not as scattered mixed-color blocks.
- [x] [P1-E] Historical completed behavior: place nodes inside each cluster using a deterministic row-major local grid with `columns = ceil(sqrt(clusterNodeCount))`. This is superseded by P6/P7 and must not be the final medium/large cluster layout.
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
- [x] [P3-C] Keep optimizer output subordinate to cluster readability; it may clean spacing but must not become the primary layout source or mix display-filter clusters together.
- [x] [P3-D] Avoid introducing elapsed-time budget as the layout correctness mechanism.
- [x] [P3-E] Do not reuse the current global ForceAtlas2 optimizer unchanged unless the implementation constrains it to preserve display-filter clusters.

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
- [x] [P5-G] Run full product build validation and record the exact product build scope: Web build, root Go `cmd`/`internal`, and launcher Go modules.

### P6 - Visual Reopen From Screenshot Evidence

- [x] [P6-A] Record `reports/problem/screenshot_1779285599.png` as failing visual evidence for the current clustered layout.
- [x] [P6-B] Record `reports/problem/aaaa.jpg` as the target reference for separated 2D color islands.
- [x] [P6-C] Supersede row-major local grid acceptance for medium and large clusters.
- [x] [P6-D] Update benchmark and evidence ledgers so previous bounding-box-only validation is treated as insufficient.
- [x] [P6-E] Add or update diagnostics/tests so rail-like or line-like cluster compression fails validation.

### P7 - Organic Filter-Island Layout Correction

- [x] [P7-A] Replace medium/large in-cluster row-major placement with deterministic two-dimensional island/cloud placement.
- [x] [P7-B] Keep the outer cluster source as existing Node Type filters and existing node type colors, with the P8 Documentation display-filter exception handled separately.
- [x] [P7-C] Size each cluster island from visible node count, capped node diameter, and spacing so density is bounded.
- [x] [P7-D] Pack cluster islands into separated macro-regions with visible gutters; large clusters receive larger regions and small clusters occupy secondary/peripheral regions.
- [x] [P7-E] Prevent medium and large clusters from producing extreme aspect ratios or thin rail-like shapes.
- [x] [P7-F] Keep placement free of centrality, degree, hub, semantic importance, or connectivity ranking.
- [x] [P7-G] Preserve existing edge data and edge visibility behavior; do not reduce edge count or hide cross-cluster links to imitate the sample image.

### P8 - Documentation Filter And Center Island

- [x] [P8-A] Define documentation node classification using simple existing node facts such as file path, extension, name, and existing metadata.
- [x] [P8-B] Add a `Documentation` Web filter/node type entry with its own dedicated color.
- [x] [P8-C] Route documentation-classified nodes into the `Documentation` display cluster while preserving raw graph label and relationship metadata.
- [x] [P8-D] Place the `Documentation` island at the center of the large circular graph field.
- [x] [P8-E] Place all other filter/color islands around the documentation center with visible gutters.
- [x] [P8-F] Preserve relationship data, edge count, and existing edge visibility behavior.
- [x] [P8-G] Add tests or diagnostics proving the Documentation cluster has one color, is centered, and does not mix with outer clusters.

### P9 - Manual Optimizer Recheck

- [x] [P9-A] Reconfirm graph load renders the island layout immediately without starting any optimizer.
- [x] [P9-B] Reconfirm the optimizer starts only from the explicit Web UI button.
- [x] [P9-C] Ensure manual optimization preserves display-filter/color island boundaries and does not mix labels into one global layout.

### P10 - Visual And Full Validation

- [x] [P10-A] Capture browser evidence after graph ready on the representative large repo and compare it against the new island-distribution criteria.
- [x] [P10-B] Record per-cluster diagnostics for color, node count, width, height, aspect ratio, density, inter-cluster gutters, and Documentation center distance.
- [x] [P10-C] Run focused unit/e2e tests for island layout, Documentation center behavior, and manual-only optimizer behavior.
- [x] [P10-D] Run full Web build, full Web unit tests, full Web e2e tests, root product Go build, and launcher Go builds before closing again.

## Acceptance Criteria

- [x] Initial graph layout is grouped by existing display filters as clear separated two-dimensional color islands.
- [x] Documentation files are grouped into one dedicated `Documentation` Web filter/node type with a dedicated color.
- [x] The Documentation island is centered in the large circular graph field, with other filter/color islands distributed around it.
- [x] There is no new centrality, hub, degree, or importance calculation for placement.
- [x] Layout is deterministic for the same graph input.
- [x] Outer cluster order follows `FILTERABLE_LABELS` for known labels, with unknown labels appended by label string.
- [x] Only display filters present in the graph create rendered clusters; `FILTERABLE_LABELS` is the outer ordering source, not a requirement to render empty clusters.
- [x] In-cluster node placement uses deterministic two-dimensional island/cloud placement, not row-major grid placement for medium and large clusters.
- [x] Medium and large clusters have bounded aspect ratio and cannot collapse into a long line, rail, or single packed row.
- [x] Cluster area scales with visible node count and capped node diameter so nodes are not stacked into dense blocks.
- [x] Different display-filter/color islands have visible gutters and do not visually merge into adjacent colored blocks.
- [x] Each display-filter cluster uses only that filter's own render color.
- [x] Community color does not override node type/filter color in the main graph canvas.
- [x] Filter state hides/shows existing clusters and does not create a separate layout policy.
- [x] Graph load does not automatically start layout optimization, including after render/camera fit/progress completion.
- [x] Manual optimizer starts only from explicit user action.
- [x] Manual optimizer does not mix different display-filter clusters together.
- [x] Existing global ForceAtlas2 optimizer is not reused unchanged unless constrained to preserve display-filter clusters.
- [x] Existing node filters still work and visually map to the same clusters.
- [x] Existing edge visibility, graph-health filters, depth filters, selection, and focus behavior still work.
- [x] Existing relationship data and cross-cluster edge visibility are preserved; the sample image affects node placement only.
- [x] Unit tests cover deterministic island placement, shape bounds, Documentation center placement, and spacing for representative cluster sizes.
- [x] Unit/e2e tests cover one-color-per-filter-cluster behavior.
- [x] Web/browser or e2e validation covers no-auto-optimizer after render and manual optimizer trigger after island correction.
- [x] Web/browser or e2e validation includes screenshot-backed or diagnostic-backed evidence that `reports/problem/screenshot_1779285599.png` no longer represents the output shape.
- [x] Benchmark ledger records before/after load, optimizer behavior, and island geometry after the visual correction.
- [x] Product build evidence is recorded without treating intentionally non-buildable analysis fixtures as product packages.

## Closure Definition

The plan can be marked complete when the Web graph renders deterministic display-filter color islands immediately on load, Documentation is its own colored center island, medium and large outer clusters have two-dimensional readable spread with visible gutters, no automatic optimizer starts on graph load, the user can manually start layout optimization, validation passes, product build evidence is recorded, benchmark evidence is recorded, and the final diff has no backend/schema/ranking-system/ranking-algorithm creep.
