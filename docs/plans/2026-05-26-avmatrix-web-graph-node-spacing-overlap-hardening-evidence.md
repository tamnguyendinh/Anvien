# AVmatrix Web Graph Node Spacing And Overlap Hardening Evidence Ledger

Date: 2026-05-26

Status: Complete

Companion files:

- Plan: [2026-05-26-avmatrix-web-graph-node-spacing-overlap-hardening-plan.md](2026-05-26-avmatrix-web-graph-node-spacing-overlap-hardening-plan.md)
- Benchmark ledger: [2026-05-26-avmatrix-web-graph-node-spacing-overlap-hardening-benchmark.md](2026-05-26-avmatrix-web-graph-node-spacing-overlap-hardening-benchmark.md)

## Evidence Rules

Record evidence as each evidenced task is completed. Evidence should include source traces, impacted files, impact output summaries, test commands, build commands, browser screenshots, DOM diagnostic output, and concise observations needed to audit the plan later.

For doc-only commits, do not use AVmatrix.

Do not record inferred runtime behavior as final evidence. Every behavior claim must include source inspection, command output, test output, browser evidence, or exact geometry measurements.

Keep this file separate from the benchmark ledger. This file records what was inspected, what was run, what changed, and what artifacts prove behavior. Quantitative geometry or performance measurements belong in the benchmark ledger.

## E0 - Plan Creation Evidence

Date: 2026-05-26

Status: recorded

Created file set:

- `docs/plans/2026-05-26-avmatrix-web-graph-node-spacing-overlap-hardening-plan.md`
- `docs/plans/2026-05-26-avmatrix-web-graph-node-spacing-overlap-hardening-evidence.md`
- `docs/plans/2026-05-26-avmatrix-web-graph-node-spacing-overlap-hardening-benchmark.md`

Plan creation scope:

- Treat dense Web graph node overlap/crowding as a product readability bug.
- Define the default requirement as one rendered node diameter of empty edge-to-edge gap between rendered circular nodes.
- Plan a hard layout invariant instead of only increasing a spacing constant.
- Keep evidence and benchmark records separate.
- Preserve existing graph orientation labels, filter behavior, island/ring spacing, and deterministic layout semantics.

Doc-only note:

- This plan creation is documentation-only, so AVmatrix was not used for this commit slice.

## E1 - Initial Source Trace From Prior Investigation

Date: 2026-05-26

Status: preliminary; implementation must re-verify with fresh AVmatrix graph before code edits

Relevant source owners observed:

| Area | Path | Observed responsibility |
|---|---|---|
| Web graph conversion and layout | `avmatrix-web/src/lib/graph-adapter.ts` | Computes rendered node size caps, cluster node spacing, island radius, deterministic node offsets, island placement, and ring placement. |
| Sigma rendering integration | `avmatrix-web/src/hooks/useSigma.ts` | Applies rendered node size caps and camera/rendering behavior. |
| Graph canvas | `avmatrix-web/src/components/GraphCanvas.tsx` | Hosts graph UI and may expose diagnostics or validation hooks if needed. |
| Geometry tests | `avmatrix-web/test/unit/graph-adapter.edge-geometry.test.ts` | Contains existing island/ring geometry tests, but must be extended with pairwise same-island spacing checks. |
| Label tests | `avmatrix-web/test/unit/graph-orientation-labels.test.ts` | Protects graph orientation label metadata and overlap guardrails. |
| Browser/e2e tests | `avmatrix-web/e2e/graph-orientation-labels.spec.ts` | Validates label visibility and overlap behavior in browser. |

Important source search command used during plan creation:

```powershell
rg -n "getClusterNodeSpacing|getClusterIslandRadius|getIslandOffset|MAX_RENDERED_NODE_SIZE|capRenderedNodeSize|golden|GOLDEN_ANGLE|GraphCanvas|useSigma|edge-geometry" avmatrix-web -g "*.ts" -g "*.tsx"
```

Observed source symbols:

- `MAX_RENDERED_NODE_SIZE`
- `capRenderedNodeSize`
- `getClusterNodeSpacing`
- `getClusterIslandRadius`
- `getIslandOffset`
- `GOLDEN_ANGLE`
- `GraphCanvas`
- `useSigma`

## E2 - Initial Problem Finding From Prior Investigation

Date: 2026-05-26

Status: preliminary; implementation must reproduce in Phase 0 before code edits

Observed current behavior:

- Current layout uses deterministic spiral placement for nodes inside an island.
- The broad layout has island/ring spacing tests, but no hard same-island pairwise spacing invariant was observed.
- A dense island can pass island/ring separation checks while still placing two rendered nodes too close together.

Preliminary conclusion:

- The user's report is credible and aligns with the current layout structure.
- The proposed UX rule is correct as a product default, but the implementation should express it as a minimum center-distance invariant derived from rendered node size semantics.
- Merely increasing a global spacing constant is not enough proof, because perturbation, future tuning, and camera fit behavior can still reintroduce visual crowding.

## E3 - Plan Creation Source Search

Date: 2026-05-26

Status: recorded

Command:

```powershell
rg -n "node spacing|island radius|edge gap|overlap|cluster island|graph label|ring label" docs\plans avmatrix-web -g "*.md" -g "*.ts" -g "*.tsx"
```

Observed related historical plan context:

- The completed skill-system plan already included adaptive island/ring spacing and graph orientation labeling work.
- That plan's acceptance language included readable island spacing, but the follow-up issue is narrower: enforce pairwise node-node clearance inside dense islands.
- The new plan is a follow-up bug hardening plan and must not rewrite prior closed plan evidence.

## E4 - Pending Implementation Evidence

Date: 2026-05-26

Status: pending

Record implementation evidence here as phases complete:

- fresh AVmatrix graph counts before graph-based implementation work;
- impact analysis blast radius for edited graph layout/Sigma/canvas symbols;
- source diffs and touched files;
- geometry test commands and results;
- Web unit test commands and results;
- e2e/browser test commands and results;
- screenshot artifact paths;
- `detect-changes` output before implementation commits;
- commit hashes for completed implementation slices.

## E5 - P0-A Fresh Graph Refresh

Date: 2026-05-26

Status: recorded

Command:

```powershell
.\avmatrix\bin\avmatrix.exe analyze --force
```

Output summary:

```text
analyzed E:\AVmatrix-GO
files: scanned=777 parsed=578 unsupported=199 failed=0
graph: nodes=88619 relationships=121516 path=E:\AVmatrix-GO\.avmatrix\graph.json
```

Conclusion:

- The AVmatrix graph was refreshed before graph-based implementation work.
- Subsequent AVmatrix query/context/impact/detect-changes commands in this plan use this fresh graph unless another stale warning requires a new refresh.

## E6 - P0-B Web Graph Layout Owner Trace

Date: 2026-05-26

Status: recorded

AVmatrix query:

```powershell
.\avmatrix\bin\avmatrix.exe query "Web graph layout node spacing island radius sigma camera dense node overlap" --repo AVmatrix
```

Top owner results:

| Rank | Symbol | Path | Finding |
|---:|---|---|---|
| 1 | `knowledgeGraphToGraphology` | `avmatrix-web/src/lib/graph-adapter.ts` | Converts backend graph into Sigma graph and applies layout. |
| 2 | `useSigma` | `avmatrix-web/src/hooks/useSigma.ts` | Sigma rendering, node reducers, camera/rendering integration. |
| 4 | `getClusterIslandRadius` | `avmatrix-web/src/lib/graph-adapter.ts` | Computes island footprint from node count and node spacing. |
| 5 | `getClusterNodeSpacing` | `avmatrix-web/src/lib/graph-adapter.ts` | Computes current graph-level spacing bucket. |

Source search:

```powershell
rg -n "getClusterNodeSpacing|getClusterIslandRadius|getIslandOffset|place|ringGap|islandGap|largestAdjacentClusterSpan|largestAdjacentRingSpan|capRenderedNodeSize|MAX_RENDERED_NODE_SIZE|fit|camera|graph.order|Sigma" avmatrix-web\src avmatrix-web\test avmatrix-web\e2e -g "*.ts" -g "*.tsx"
```

Recorded owners:

| Area | Path / symbol | Finding |
|---|---|---|
| Rendered size cap | `avmatrix-web/src/lib/graph-adapter.ts` `MAX_RENDERED_NODE_SIZE`, `MAX_DENSE_RENDERED_NODE_SIZE`, `getMaxRenderedNodeSize`, `capRenderedNodeSize` | Current maximum rendered node size is capped at `3`. |
| Graph-level node spacing | `getClusterNodeSpacing` | Current buckets return `42`, `36`, `30`, `32`, or `34` depending on total graph size. |
| Island radius | `getClusterIslandRadius` | Current formula is `nodeSpacing * sqrt(nodeCount - 1) * 1.22 + nodeSpacing * 5`; it does not use rendered node diameter or hard pairwise spacing. |
| Same-island placement | `getIslandOffset` | Current deterministic golden-angle spiral includes radial `organicWave` and angular jitter. |
| Island placement | `applyFilterBasedClusteredLayout` local `placeCluster` / `placeRingIslands` | Nodes are sorted deterministically, offsets are centered, islands are placed on balanced slots/pinwheel orbit. |
| Island/ring gutters | `islandGap`, `largestAdjacentClusterSpan`, `ringGap`, `largestAdjacentRingSpan` | Broad island/ring spacing is adaptive, but independent from pairwise same-island gap. |
| Sigma rendering | `avmatrix-web/src/hooks/useSigma.ts` `capNodeReducerSize`, `useSigma` | Node reducer caps `size`; hover/label rendering receives scaled `data.size`. |
| Graph canvas labels | `avmatrix-web/src/components/GraphCanvas.tsx` `createLayoutRingBounds`, orientation label refresh | Graph orientation label overlay depends on layout coordinates and camera state. |
| Unit geometry tests | `avmatrix-web/test/unit/graph-adapter.edge-geometry.test.ts` | Existing tests cover rendered size caps, island separation, ring expansion, and non-rail shape, but not pairwise same-island node gap. |
| Browser label tests | `avmatrix-web/e2e/graph-orientation-labels.spec.ts` | Existing e2e validates labels and label overlap, not dense node-node spacing. |

## E7 - P0-C Dense Node Overlap Reproduction

Date: 2026-05-26

Status: recorded

Reproduction command:

```powershell
@'
// Inline Node.js reproduction of current getClusterNodeSpacing/getIslandOffset formula.
'@ | node -
```

Reproduction details:

- Used the current `GOLDEN_ANGLE`, `getStableLabelSeed`, `getClusterNodeSpacing`, and `getIslandOffset` formulas from `avmatrix-web/src/lib/graph-adapter.ts`.
- Used `MAX_RENDERED_NODE_SIZE = 3` as rendered radius for baseline geometry because local hover/label code and Sigma display data treat `size` as radius.
- For total graph size `6100`, current `getClusterNodeSpacing` returns `30`.

Observed failure:

- `large dense island`: `1800` nodes, seed label `frontend:Function`, minimum center distance `0.361`, minimum edge gap `-5.639`, overlap count `22`, target gap violation count `56`.
- This proves the current layout can create real rendered-circle overlap, not only insufficient whitespace.

## E8 - P0-D Rendered Size Semantics

Date: 2026-05-26

Status: recorded

Source evidence:

| Source | Finding |
|---|---|
| `avmatrix-web/src/lib/graph-adapter.ts` | `MAX_RENDERED_NODE_SIZE = 3`; `capRenderedNodeSize` clamps graph node `size` to that maximum. |
| `avmatrix-web/src/hooks/useSigma.ts` | `capNodeReducerSize` applies `capRenderedNodeSize` in Sigma reducers. |
| `avmatrix-web/src/hooks/useSigma.ts` hover renderer | `const nodeSize = capRenderedNodeSize(data.size || 8, graph.order)` then uses `context.arc(data.x, data.y, nodeSize + 4, ...)`, which treats `nodeSize` as radius. |
| `avmatrix-web/node_modules/sigma/dist/index-16136237.cjs.prod.js` | Sigma label/hover paths position labels with `data.x + data.size + 3` and draw hover circles using `data.size + PADDING`, which treats display `size` as a radius. |
| `avmatrix-web/node_modules/sigma/dist/index-16136237.cjs.prod.js` | Edge rendering uses target node `size` as a node radius when offsetting edge/arrow geometry. |

Conclusion:

- For this plan's layout checks, Web/Sigma `size` is treated as rendered radius.
- With `MAX_RENDERED_NODE_SIZE = 3`, rendered diameter is `6`.
- The no-overlap center threshold is `6`; the plan's stronger default one-node-diameter edge gap requires center distance `12`.

## E9 - P0-E Existing Test Inventory

Date: 2026-05-26

Status: recorded

Inventory:

| Test surface | Current coverage | Missing coverage for this plan |
|---|---|---|
| `graph-adapter.edge-geometry.test.ts` rendered size tests | Verifies scaled/rendered node size caps. | Does not assert node-node minimum edge gap inside dense islands. |
| `graph-adapter.edge-geometry.test.ts` island separation tests | Verifies separate visual islands and pinwheel gutters. | Island bounding separation can pass while nodes inside one island overlap. |
| `graph-adapter.edge-geometry.test.ts` non-rail tests | Verifies medium/large clusters are two-dimensional. | Does not enforce pairwise clearance. |
| `graph-orientation-labels.test.ts` | Verifies label metadata and viewport label overlap guardrails. | Does not measure node-node overlap. |
| `graph-orientation-labels.spec.ts` | Browser validates ring/island labels, filter updates, screenshots, and label overlap count. | Does not expose dense node spacing diagnostics or screenshots focused on node-node gap. |

Conclusion:

- Existing tests protect broad layout shape and label orientation.
- This plan needs new unit geometry checks for same-island pairwise spacing and at least one browser/e2e diagnostic for dense node spacing.

## E10 - P0-F Impact Analysis

Date: 2026-05-26

Status: recorded

Commands:

```powershell
.\avmatrix\bin\avmatrix.exe impact "getIslandOffset" --repo AVmatrix --direction upstream
.\avmatrix\bin\avmatrix.exe impact "getClusterIslandRadius" --repo AVmatrix --direction upstream
.\avmatrix\bin\avmatrix.exe impact "applyFilterBasedClusteredLayout" --repo AVmatrix --direction upstream
.\avmatrix\bin\avmatrix.exe impact "capRenderedNodeSize" --repo AVmatrix --direction upstream
```

Results:

| Target | Risk | Summary |
|---|---|---|
| `getIslandOffset` | LOW | 1 direct affected symbol: local `placeCluster`; frontend/layout only. |
| `getClusterIslandRadius` | LOW | 1 direct affected symbol: local `placeCluster`; frontend/layout only. |
| `applyFilterBasedClusteredLayout` | CRITICAL | 6 affected frontend symbols across layout and Web graph UI, 16 affected processes; direct caller includes `useSigma`. |
| `capRenderedNodeSize` | UNKNOWN / ambiguous | AVmatrix found the real function plus ResolutionGap candidates; no edit to this symbol is currently planned. |

Blast radius warning:

- `applyFilterBasedClusteredLayout` is CRITICAL because it is the Web graph layout entry point used by `useSigma` and graph UI flows.
- This is workflow safety information, not a blocker. The implementation must keep changes scoped, deterministic, and test-backed.

## E11 - Phase 1 And Phase 2 Layout Implementation

Date: 2026-05-26

Status: implementation edited; validation pending

Touched files:

| Path | Change |
|---|---|
| `avmatrix-web/src/lib/graph-adapter.ts` | Added rendered node radius/diameter/edge-gap/center-distance helpers; added deterministic same-island offset generation with spatial-grid pairwise clearance; updated island radius to use actual centered offsets plus spacing margin; reused expanded island radii for existing island/ring placement. |
| `avmatrix-web/src/components/GraphCanvas.tsx` | Added layout node-spacing diagnostics derived from the rendered Sigma graph and recorded them into runtime diagnostics for browser validation. |
| `avmatrix-web/src/lib/runtime-diagnostics.ts` | Added `layoutNodeSpacing` diagnostics and `recordLayoutNodeSpacing`. |

Design summary:

- `size` is treated as rendered radius.
- Default rendered edge gap is one rendered node diameter.
- Default minimum center distance is two rendered node diameters.
- Candidate spiral positions still use the existing deterministic golden-angle formula.
- A spatial grid rejects candidate positions that would fall below the minimum center distance inside the same island.
- Organic radial/angular perturbation is preserved only when it does not violate the hard gap invariant.
- If spiral candidates cannot satisfy the invariant within bounded attempts, a deterministic fallback ring placement continues outward until a valid position is found.
- Island radius is expanded from the actual centered offsets plus the minimum center-distance margin.
- Existing island and ring placement already consumes each island's `radius`, so expanded dense islands automatically feed neighboring island gutters and macro-ring radius.

Additional impact before GraphCanvas/runtime diagnostics edit:

```powershell
.\avmatrix\bin\avmatrix.exe analyze --force
.\avmatrix\bin\avmatrix.exe impact "createLayoutRingBounds" --repo AVmatrix --direction upstream
.\avmatrix\bin\avmatrix.exe impact "WebRuntimeDiagnostics" --repo AVmatrix --direction upstream
```

Results:

| Target | Risk | Summary |
|---|---|---|
| `createLayoutRingBounds` | LOW | 0 upstream impacted symbols reported. |
| `WebRuntimeDiagnostics` | MEDIUM | 29 impacted symbols across frontend diagnostics consumers; change is additive and backward-compatible. |

## E12 - Phase 3 And Phase 4 Test Scaffolding

Date: 2026-05-26

Status: tests added; validation pending

Touched files:

| Path | Change |
|---|---|
| `avmatrix-web/test/unit/graph-adapter.edge-geometry.test.ts` | Added spacing contract test, dense same-island pairwise spacing test, and regression coverage for the previous dense spiral close-pair condition. |
| `avmatrix-web/test/unit/runtime-diagnostics.test.ts` | Added runtime diagnostics unit coverage for `layoutNodeSpacing`. |
| `avmatrix-web/e2e/graph-orientation-labels.spec.ts` | Added dense spacing fixture, runtime diagnostics polling, overlap/gap assertions, and desktop/smaller-viewport screenshot capture for graph node spacing. |

Pending validation:

- Full build gate must run before tests.
- Focused Web unit tests must pass.
- Full Web unit tests must pass.
- Web e2e/browser validation must pass and produce screenshot artifacts.

## E13 - P5-A Full Build Gate

Date: 2026-05-26

Status: pass

Command:

```powershell
powershell -ExecutionPolicy Bypass -File avmatrix-launcher\build.ps1
```

Result:

- Exit code: `0`
- Go runtime detected: `go version go1.26.3 windows/amd64`
- Web build command: `tsc -b && vite build`
- Vite modules transformed: `2931`
- Vite build duration: `18.66s`
- Native runtime file `avmatrix\bin\lbug_shared.dll` was already up to date.

Warnings:

- Vite reported existing chunk-size and mixed dynamic/static import warnings. The build still passed.

## E14 - P5-B Focused Web Unit Tests And P3-E Geometry Metrics

Date: 2026-05-26

Status: pass

Focused test command:

```powershell
npm run test -- graph-adapter.edge-geometry.test.ts graph-orientation-labels.test.ts runtime-diagnostics.test.ts
```

Working directory:

```text
E:\AVmatrix-GO\avmatrix-web
```

Result:

- Exit code: `0`
- Test files: `3 passed`
- Tests: `38 passed`
- `graph-adapter.edge-geometry.test.ts`: `25 tests passed`
- `graph-orientation-labels.test.ts`: `7 tests passed`
- `runtime-diagnostics.test.ts`: `6 tests passed`

Focused geometry checks added:

- one rendered node diameter is the default edge gap;
- dense same-island nodes have no overlap and no target gap violations;
- regression coverage for the previous dense spiral close-pair condition.

Final geometry measurement command:

```powershell
@'
// Inline Node.js measurement mirroring the implemented deterministic spacing algorithm.
'@ | node -
```

Result summary:

- Small island, `100` nodes: min center `40.535`, min edge gap `34.535`, overlap `0`, target violations `0`.
- Medium island, `1000` nodes: min center `13.417`, min edge gap `7.417`, overlap `0`, target violations `0`.
- Large dense island, `1800` nodes: min center `12.048`, min edge gap `6.048`, overlap `0`, target violations `0`.
- Regression fixture, `1800` nodes: min center `12.107`, min edge gap `6.107`, overlap `0`, target violations `0`.

## E15 - P5-C Full Web Unit Test First Attempt

Date: 2026-05-26

Status: failed; fixed by E16

Command:

```powershell
npm run test
```

Working directory:

```text
E:\AVmatrix-GO\avmatrix-web
```

Result:

- Exit code: `1`
- Test files: `1 failed`, `45 passed`
- Tests: `4 failed`, `377 passed`

Failure:

- `test/unit/GraphCanvas.selection-performance.test.tsx` failed because its mocked `../../src/lib/graph-adapter` did not provide the newly imported `getMinimumNodeEdgeGap` export.
- Same mock also needs `getMinimumNodeCenterDistance` because `GraphCanvas` now imports it for layout spacing diagnostics.

Next action:

- Update the focused mock to include the new spacing helper exports, then rerun the full Web unit suite.

## E16 - P5-C Full Web Unit Test Rerun

Date: 2026-05-26

Status: pass

Fix applied:

- Updated `avmatrix-web/test/unit/GraphCanvas.selection-performance.test.tsx` mock for `../../src/lib/graph-adapter` to include `getMinimumNodeEdgeGap` and `getMinimumNodeCenterDistance`.

Command:

```powershell
npm run test
```

Working directory:

```text
E:\AVmatrix-GO\avmatrix-web
```

Result:

- Exit code: `0`
- Test files: `46 passed`
- Tests: `381 passed`
- Duration: `31.49s`

## E17 - P5-D Web E2E First Attempt

Date: 2026-05-26

Status: failed; fixture adjusted

Command:

```powershell
npm run test:e2e -- graph-orientation-labels.spec.ts
```

Working directory:

```text
E:\AVmatrix-GO\avmatrix-web
```

Server handling:

- Started `npm run dev` on `127.0.0.1:5228` for the test run.
- Stopped the Vite process tree in `finally` after Playwright completed.

Result:

- Exit code: `1`
- Tests: `2 passed`, `1 failed`

Failure:

- The new dense spacing test passed the node spacing diagnostics, but failed a label-count guard because the dense fixture only produced `2` visible ring labels while the assertion expected at least `3`.
- The failure was in the test fixture/assertion, not the dense node spacing invariant.

Fix:

- Added a backend dense cluster to the fixture so the browser test exercises backend, API, and frontend rings while preserving the dense frontend island.
- Updated the dense fixture node-count expectation from `>= 1400` to `>= 1480`.

## E18 - P5-D Browser E2E And Screenshot Validation

Date: 2026-05-26

Status: pass

Command:

```powershell
npm run test:e2e -- graph-orientation-labels.spec.ts
```

Working directory:

```text
E:\AVmatrix-GO\avmatrix-web
```

Server handling:

- Started `npm run dev` on `127.0.0.1:5228` for the test run.
- Stopped the Vite process tree in `finally` after Playwright completed.

Result:

- Exit code: `0`
- Tests: `3 passed`
- Duration: `14.2s`

Screenshot artifacts:

- `avmatrix-web/test-results/graph-orientation-labels-G-1d271-ted-by-the-default-node-gap-chromium/graph-node-spacing-dense-desktop.png`
- `avmatrix-web/test-results/graph-orientation-labels-G-1d271-ted-by-the-default-node-gap-chromium/graph-node-spacing-dense-small.png`
- `avmatrix-web/test-results/graph-orientation-labels-G-e5a0a-t-and-updates-after-filters-chromium/graph-orientation-labels-small-filtered.png`

Browser diagnostic observation:

- Runtime diagnostics exposed `layoutNodeSpacing` for the dense fixture and the e2e assertion consumed it for node count, island count, required gap, minimum observed gap, overlap count, and target-gap violation count.
- The dense desktop and smaller-viewport screenshot assertions passed with ring/island labels present and no label overlap.
- The existing smaller-viewport filter e2e still passed after the layout changes, proving labels update after filter visibility changes.
- No Sigma camera/default fit change was required for this slice because the dense browser fixture remained inspectable at desktop and smaller viewport sizes.

Quantitative browser metrics are recorded in the benchmark ledger `B4`.

## E19 - P5-E Backend Or Generated-Contract Validation Scope

Date: 2026-05-26

Status: not required

Scope check:

- Implementation touched Web graph layout, GraphCanvas diagnostics, runtime diagnostics, Web unit tests, and Web e2e tests.
- No backend API handler, generated payload, generated contract file, or shared backend/frontend payload schema was changed.

Conclusion:

- Focused backend or generated-contract validation was not required for this implementation slice.

## E20 - P5-F Pre-Commit Change Detection

Date: 2026-05-26

Status: pass

Graph refresh command before change detection:

```powershell
.\avmatrix\bin\avmatrix.exe analyze --force
```

Graph refresh output:

```text
analyzed E:\AVmatrix-GO
files: scanned=777 parsed=578 unsupported=199 failed=0
graph: nodes=89032 relationships=122028 path=E:\AVmatrix-GO\.avmatrix\graph.json
```

Change detection command:

```powershell
.\avmatrix\bin\avmatrix.exe detect-changes --repo AVmatrix --scope all
```

Output summary:

| Field | Value |
|---|---:|
| Risk level | `high` |
| Changed files | `10` |
| Changed symbols | `485` |
| Affected processes | `15` |
| Changed app layers | `docs: 31`, `frontend: 246`, `frontend_test: 208` |
| Changed functional areas | `documentation: 31`, `layout: 130`, `unknown: 255`, `web_graph_ui: 69` |
| Affected app layers | `frontend: 9`, `mixed: 6` |
| Affected functional areas | `layout: 4`, `mixed: 8`, `web_graph_ui: 3` |
| Resolution health degraded nodes | `0` |
| Total resolution gap count | `0` |

Blast radius note:

- `high` is a warning that the frontend layout/Web graph UI area is important and must be handled carefully.
- It is not treated as a ban on the required code changes.

Additional pre-commit check:

```powershell
git diff --check
```

Result:

- Exit code: `0`
- No whitespace errors reported.

## E21 - P5-G And P5-H Commit And Closure

Date: 2026-05-26

Status: complete

Implementation commit:

```text
05d5ae4 Harden graph node spacing layout
```

Closure state:

- Code implementation, focused geometry tests, full Web unit tests, Web e2e/browser screenshots, browser diagnostics, benchmark metrics, pre-commit change detection, and implementation commit are recorded.
- Dense graph same-island nodes no longer overlap in the measured fixtures.
- The default browser/runtime spacing contract records one rendered node diameter of edge-to-edge gap.
- The implementation commit was completed before marking `P5-G` and `P5-H`.
