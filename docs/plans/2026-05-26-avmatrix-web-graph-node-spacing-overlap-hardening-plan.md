# AVmatrix Web Graph Node Spacing And Overlap Hardening Plan

Date: 2026-05-26

Status: Complete

Companion files:

- Evidence ledger: [2026-05-26-avmatrix-web-graph-node-spacing-overlap-hardening-evidence.md](2026-05-26-avmatrix-web-graph-node-spacing-overlap-hardening-evidence.md)
- Benchmark ledger: [2026-05-26-avmatrix-web-graph-node-spacing-overlap-hardening-benchmark.md](2026-05-26-avmatrix-web-graph-node-spacing-overlap-hardening-benchmark.md)

## Master rules

1. Use AVmatrix for codebase analysis and impact checks while working on implementation slices in this plan.
2. As each task is completed, update the corresponding checklist item immediately.
3. Run the full build gate before testing; include focused Web unit/e2e/browser screenshot validation for graph spacing behavior.
4. Record benchmark results as each benchmarkable task is completed. Benchmarkable means measured product/runtime performance, capacity, package/startup size, graph/DB throughput, query hit rate, command output inventory, graph inventory counts, source-site inventory counts, generated skill inventory counts, setup/package file inventories, resolved-edge accuracy, or graph layout geometry metrics. Build/test/e2e timings are validation evidence unless the slice changes those systems.
5. Record evidence as each evidenced task is completed.
6. For doc-only commits, do not use AVmatrix.
7. After each completed implementation slice, commit the work, then continue until the full plan is complete.

## Problem

The Web graph layout can place nodes too close together inside large node islands. On large repositories, high node counts make this visible as stacked, crowded, or overlapping nodes even when island-to-island and ring-to-ring spacing appears acceptable.

The current layout path in `avmatrix-web/src/lib/graph-adapter.ts` computes density-aware cluster spacing and island radius, then places nodes with a deterministic golden-angle spiral. That layout protects broad island footprint, but it does not define or test a hard pairwise minimum distance between the rendered circles of two nodes in the same island.

This is a product readability bug. A graph may pass broad bounding-box tests while still failing the user's actual visual task: seeing individual nodes clearly enough to inspect a large repository.

The required default behavior is that two rendered circular nodes have at least one rendered node diameter of empty space between their edges. In geometry terms:

- if rendered node diameter is `d`, edge-to-edge gap must be `>= d`;
- therefore center-to-center distance must be `>= 2d`;
- if the renderer treats the Web `size` attribute as radius `r`, then `d = 2r` and center-to-center distance must be `>= 4r`.

The implementation must verify the renderer size semantics rather than assuming whether Sigma `size` is a radius or diameter.

## Scope

Implementation may touch:

- Web graph layout geometry in `avmatrix-web/src/lib/graph-adapter.ts`;
- node-size capping and layout-size helpers if the minimum gap needs a shared rendered-size contract;
- Sigma/camera behavior in `avmatrix-web/src/hooks/useSigma.ts` only if default fit-to-view compresses the graph into unreadable visual density;
- graph canvas diagnostics or test hooks in `avmatrix-web/src/components/GraphCanvas.tsx` only if e2e validation needs measurable layout output;
- Web unit tests under `avmatrix-web/test/unit`, especially graph adapter geometry tests;
- Web e2e tests under `avmatrix-web/e2e` for browser-level graph spacing screenshots and assertions;
- benchmark/evidence ledgers for minimum node gap, overlap counts, graph footprint, visible labels, screenshot artifacts, and validation output.

Out of scope unless source inspection proves it is required:

- backend graph analysis semantics;
- graph schema or relationship extraction;
- generated API contract semantics unrelated to Web graph layout payloads;
- replacing Sigma or the full graph rendering stack;
- changing graph labels, node classification, or filter taxonomy;
- editing generated `AGENTS.md`, `CLAUDE.md`, or `.claude/skills/avmatrix/**` directly as source of truth.

## Design Decisions

The fix must be an explicit layout invariant, not only a larger constant. Increasing `nodeSpacing` may reduce overlap in common cases, but it does not prove that the golden-angle spiral, jitter, organic wave, camera fitting, or future density changes cannot reintroduce local collisions.

The layout must compute a minimum rendered node clearance from the actual maximum rendered node size used by the graph. The default minimum must match the product rule: one empty node diameter between two node edges.

The implementation should introduce named geometry constants or helpers such as:

| Concept | Purpose |
|---|---|
| rendered node radius/diameter | Converts Sigma node size into geometry units. |
| minimum edge gap | Required empty space between rendered node circles. |
| minimum center distance | Distance threshold used by deterministic placement and tests. |
| spiral band gap | Prevents adjacent spiral bands from becoming visually stacked. |
| island gutter | Keeps neighboring islands separate after island radius expands. |
| ring gutter | Keeps macro rings separate after island footprints expand. |

The node placement algorithm must preserve deterministic output. Acceptable implementation strategies include:

- deriving spiral radial increments from the minimum center distance and verifying pairwise clearance for nearby candidates;
- using a deterministic low-discrepancy or ring-packing strategy with a hard minimum center distance;
- adding a deterministic collision-relaxation pass that never moves nodes below the minimum gap and does not require manual optimizer interaction;
- removing or clamping organic perturbation when it would violate the spacing invariant.

The implementation must avoid rail-like or line-like dense shapes. A layout that satisfies minimum pairwise distance but turns islands into long straight bands is not acceptable.

The graph/canvas footprint may expand for large repositories. The product should prefer a larger deterministic layout with usable camera behavior over compressing nodes into overlap. If default fit-to-view makes the first viewport too zoomed out for large graphs, the implementation must record that as evidence and adjust camera/readability behavior or define an explicit large-graph UX boundary.

The pairwise spacing check should be efficient enough for large islands. Unit tests can use exact all-pairs checks on bounded fixtures, while production code should avoid quadratic work on full large graphs unless it is proven acceptable by benchmark.

Existing label and island/ring spacing behavior must remain intact. The fix must preserve macro ring labels, island labels, filter visibility behavior, and current graph orientation semantics.

## Acceptance Criteria

- The current node placement owner, rendered node size contract, island radius computation, island placement, ring placement, and camera fit behavior are traced and recorded in the evidence ledger.
- Impact analysis is run before editing graph layout, Sigma/camera, or graph canvas symbols. HIGH or CRITICAL blast radius is reported as a warning and handled carefully, not treated as a ban on required edits.
- The implementation defines a default minimum edge-to-edge node gap of one rendered node diameter.
- The implementation computes minimum center distance from actual rendered node size semantics and records the interpretation in evidence.
- Dense island placement has a hard pairwise no-overlap invariant in layout coordinates.
- Dense island placement satisfies the stronger default gap target: edge gap `>= one rendered node diameter`.
- Island radius grows from node count, rendered node size, minimum center distance, and any spiral band gap.
- Neighboring islands are spaced from expanded island radii plus gutter.
- Macro ring radius grows from expanded island footprints plus gutter.
- Organic/jitter perturbation cannot violate minimum pairwise node spacing.
- Unit tests include dense same-island fixtures that assert no overlap and the one-node-diameter edge gap.
- Unit tests include regression coverage for large islands where the previous spiral formula could place nodes too close.
- Unit tests preserve existing island/ring separation and label visibility behavior.
- Browser/e2e validation captures desktop and smaller-viewport screenshots for a dense graph and records node overlap/gap diagnostics.
- Web build and full repository build gate pass before closure.
- `avmatrix detect-changes --repo AVmatrix --scope all` runs before each implementation commit and the changed scope is recorded.
- Evidence and benchmark ledgers are updated after each completed task or measured benchmark, with evidence kept separate from quantitative metrics.

## Phase 0 - Baseline Trace And Reproduction

- [x] [P0-A] Refresh the AVmatrix graph before graph-based implementation work and record graph counts in the evidence ledger.
- [x] [P0-B] Trace Web graph layout ownership and record the files/symbols for node spacing, node size capping, island radius, island offset, island placement, ring placement, Sigma size handling, camera fit, and relevant tests.
- [x] [P0-C] Reproduce the dense-node overlap/crowding issue with a deterministic fixture or script using the current layout formula. Record the minimum center distance, minimum edge gap, overlap count, and fixture size in the benchmark ledger.
- [x] [P0-D] Verify whether Sigma/Web `size` behaves as rendered radius or diameter for the purpose of graph layout checks. Record the source/test/browser evidence.
- [x] [P0-E] Inventory existing geometry tests and e2e tests to identify which checks already cover island/ring separation and which checks are missing pairwise node spacing.
- [x] [P0-F] Run impact analysis before editing any graph layout, Sigma/camera, or graph canvas symbols. Record blast radius and affected flows.

## Phase 1 - Spacing Contract And Geometry Design

- [x] [P1-A] Define the rendered node spacing contract: rendered size, diameter, minimum edge gap, and minimum center distance.
- [x] [P1-B] Add or update geometry helpers so spacing thresholds are named and testable instead of hidden in numeric constants.
- [x] [P1-C] Decide the deterministic placement strategy that can guarantee pairwise clearance without producing rail-like dense islands.
- [x] [P1-D] Define how organic/jitter perturbation is preserved, reduced, or removed when it conflicts with the minimum gap invariant.
- [x] [P1-E] Define expected large-graph footprint and camera behavior so the layout does not solve density by visual compression.

## Phase 2 - Layout Implementation

- [x] [P2-A] Update internal island node placement so candidate positions satisfy minimum center distance before assignment.
- [x] [P2-B] Update island radius computation so dense islands reserve enough space for the required node gap.
- [x] [P2-C] Update neighboring island spacing so expanded island footprints do not collide.
- [x] [P2-D] Update macro ring radius computation so expanded islands fit inside their rings without compressing into other rings.
- [x] [P2-E] Preserve deterministic ordering, stable anchors, ring labels, island labels, and filter/depth visibility semantics.
- [x] [P2-F] If needed, adjust Sigma camera/default fit behavior so the initial graph remains inspectable after footprint expansion.
- [x] [P2-G] Update the checklist immediately after each completed implementation subtask and record matching evidence/benchmark entries.

## Phase 3 - Unit And Geometry Validation

- [x] [P3-A] Add unit tests for pairwise same-island minimum center distance on dense fixtures.
- [x] [P3-B] Add unit tests for one-node-diameter minimum edge gap using the verified rendered-size semantics.
- [x] [P3-C] Add a regression fixture that fails against the old spiral formula or reproduces the previously observed close-pair condition.
- [x] [P3-D] Keep or extend tests for island-to-island gutter, macro-ring expansion, non-rail dense shapes, and label metadata.
- [x] [P3-E] Record geometry metrics in the benchmark ledger: fixture node count, minimum center distance, minimum edge gap, overlap count, island radius, graph footprint, and changed ratios versus baseline.

## Phase 4 - Browser And UX Validation

- [x] [P4-A] Add or update e2e diagnostics so browser tests can measure visible dense-node overlap/gap behavior without relying only on human screenshot review.
- [x] [P4-B] Capture desktop screenshot evidence for a dense graph showing readable node separation and preserved ring/island labels.
- [x] [P4-C] Capture smaller-viewport screenshot evidence for the same dense graph or representative fixture.
- [x] [P4-D] Validate filter changes do not leave stale spacing labels or stale layout measurements.
- [x] [P4-E] Record screenshot paths, DOM diagnostic output, and observations in the evidence ledger.
- [x] [P4-F] Record quantitative browser metrics in the benchmark ledger: visible overlap count, minimum measured screen gap where feasible, label overlap count, viewport size, and graph fixture size.

## Phase 5 - Build, Test, And Closure

- [x] [P5-A] Run the full build gate before testing: `powershell -ExecutionPolicy Bypass -File avmatrix-launcher\build.ps1`.
- [x] [P5-B] Run focused Web unit tests for graph adapter geometry and graph orientation labels.
- [x] [P5-C] Run full Web unit tests: `cd avmatrix-web; npm run test`.
- [x] [P5-D] Run Web e2e/browser validation for graph spacing and labels.
- [x] [P5-E] Run any focused backend or generated-contract validation only if implementation touches shared contracts or generated payloads.
- [x] [P5-F] Run `avmatrix detect-changes --repo AVmatrix --scope all` and record changed scope before committing implementation work.
- [x] [P5-G] Commit the completed implementation slice, then continue if any checklist items remain.
- [x] [P5-H] Mark the plan complete only after code, tests, browser evidence, benchmark metrics, evidence ledger, and commit state agree that dense graph nodes no longer overlap and default spacing provides one rendered node diameter of edge gap.
