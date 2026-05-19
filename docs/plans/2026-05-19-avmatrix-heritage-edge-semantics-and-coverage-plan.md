# AVmatrix Heritage Edge Semantics and Coverage Plan

Date: 2026-05-19

Status: active

Companion files:

- Benchmark ledger: [2026-05-19-avmatrix-heritage-edge-semantics-and-coverage-benchmark.md](2026-05-19-avmatrix-heritage-edge-semantics-and-coverage-benchmark.md)
- Evidence ledger: [2026-05-19-avmatrix-heritage-edge-semantics-and-coverage-evidence.md](2026-05-19-avmatrix-heritage-edge-semantics-and-coverage-evidence.md)

## Rules

1. Use AVmatrix for codebase analysis and impact checks while working on implementation slices in this plan.
2. As each task is completed, update the corresponding checklist item immediately.
3. Run a full build before testing; the test suite must include an e2e test.
4. Record benchmark results as each benchmarkable task is completed. Benchmarkable means measured product/runtime performance, capacity, package/startup size, graph/DB throughput, or conversion-inventory counts; build/test/e2e timings are validation evidence unless the slice changes those systems.
5. Record evidence as each evidenced task is completed.
6. For doc-only commits, do not use AVmatrix.
7. After each completed implementation slice, commit the work, then continue until the full plan is complete.

## Problem

The loaded graph can make heritage relationships look wrong or incomplete.

The concrete trigger is `E:\Restaurant_manager`:

- `Class=3` and `Constructor=3` look suspicious at first, but source verification shows they match the real TypeScript class declarations and constructors in the repo.
- `EXTENDS=6` and `INHERITS=6` do not mean `12` independent relationships. They are the same `6` source-target heritage pairs emitted as two relationship types.
- TypeScript source contains at least `14` `interface ... extends ...` sites and `2` `class ... extends ...` sites, but the current graph has no `EXTENDS` / `INHERITS` / `IMPLEMENTS` relationships involving those TypeScript class/interface heritage sites.
- While fixing the Web UI display for these heritage facts, the tool shell must also keep the graph workflow usable: the top bar needs a Back arrow/button beside `AVmatrix` to return to `Start-AVmatrix.html`, and the left dashboard needs drag resizing so dense node/edge/legend controls can be inspected comfortably.

This is not just a Web UI display issue and not just an analyzer issue. It has two related parts:

- relationship semantics/display: `EXTENDS` and `INHERITS` can represent the same underlying heritage fact, so the UI must not present duplicate compatibility edges as independent codebase facts;
- extraction/resolution coverage: TypeScript heritage sites that exist in the repo must either be connected to resolved graph targets or represented as unresolved/external heritage, not silently disappear.
- Web UI shell ergonomics: users need a clear route back to the Start screen and a resizable left dashboard while inspecting the corrected heritage graph display.

## Scope

Implementation may touch:

- `internal/providers/tsjs/*`
- `internal/resolution/*`
- `internal/graph/*`
- `internal/contracts/*`
- `internal/mcp/*`
- `avmatrix-web/src/components/Header.tsx`
- `avmatrix-web/src/lib/constants.ts`
- `avmatrix-web/src/lib/graph-adapter.ts`
- `avmatrix-web/src/components/FileTreePanel.tsx`
- `avmatrix-web/src/App.tsx`
- `Start-AVmatrix.html`
- Web and Go tests for relationship inventory, graph conversion, dashboard controls, and e2e display

This plan does not change the verified `Class=3` and `Constructor=3` count for `E:\Restaurant_manager` unless a later source audit proves new class declarations were missed.

## Acceptance Guardrails

- A user looking at the Web UI must not interpret one underlying heritage source-target pair as two independent relationships only because both `EXTENDS` and `INHERITS` were emitted for compatibility.
- If `INHERITS` remains in the graph payload, the UI must explain or group it as normalized/compatibility heritage rather than an additional source-code edge.
- TypeScript `class extends`, `interface extends`, and `implements` sites must have deterministic graph behavior:
  - resolved in-repo targets become graph relationships;
  - unresolved external targets are explicitly represented or explicitly recorded as unresolved heritage evidence;
  - source sites must not disappear without an audit trail.
- Analyzer and UI behavior must remain language-agnostic. Go embedded structs, TS class/interface heritage, Java/C#/Python heritage, and future provider heritage should follow one documented graph contract.
- Benchmark/evidence must report raw edge counts, unique semantic heritage pair counts, duplicate compatibility pair counts, and missing/unresolved source-site counts.
- The Web UI must provide a clear Back arrow/button beside the `AVmatrix` title that returns to `Start-AVmatrix.html`, with accessible naming and keyboard/click support.
- The left dashboard must be resizable by dragging its boundary, with bounded min/max widths, no layout overlap, and no breakage of node type, edge type, legend, or graph canvas interactions.

## Baseline Findings

- [x] [P0-A] Verify whether `Class=3` and `Constructor=3` in `E:\Restaurant_manager` are real. Result: verified as real source counts for `SSEListener`, `ApiError`, and `ErrorBoundary`.
- [x] [P0-B] Verify whether `EXTENDS=6` and `INHERITS=6` are independent relationships. Result: all `6` source-target pairs are duplicated as both `EXTENDS` and `INHERITS`.
- [x] [P0-C] Trace graph emission source. Result: `emitHeritageCompatibilityEdges` emits `EXTENDS` or `IMPLEMENTS`, then emits `ReferenceInherits` when compatibility is enabled; `ReferenceInherits` maps to `INHERITS`.
- [x] [P0-D] Check TypeScript source heritage sites in `E:\Restaurant_manager`. Result: at least `14` `interface extends` sites and `2` `class extends` sites exist.
- [x] [P0-E] Check whether current graph has TS class/interface heritage edges. Result: no `EXTENDS` / `INHERITS` / `IMPLEMENTS` edges were found involving the audited TS class/interface heritage sites.

## Review Direction

This plan should solve the real problem only if it keeps these conclusions separate:

- `Class=3` and `Constructor=3` are not a graph bug unless later parser-based inventory proves additional class declarations were missed.
- `EXTENDS` + `INHERITS` duplication is a semantics/display contract problem.
- Missing TypeScript heritage is an analyzer coverage/resolution problem.
- Go embedded structs currently appear as `EXTENDS`/`INHERITS`; the plan must decide whether that user-facing label is correct or whether the graph needs a more precise heritage kind/display label.

The implementation must trace each missing TS source site through the pipeline before changing behavior:

```text
source AST -> ScopeIR HeritageFact -> workspace heritage resolution -> graph relationship -> Web payload -> dashboard/canvas display
```

Regex source-site counts are useful baseline evidence only. Final accuracy checks must use parser/ScopeIR or graph-internal inventory where practical.

## Phase 1 - Semantics Decision

- [ ] [P1-A] Define whether `INHERITS` is a compatibility/normalized edge, a first-class display edge, or both.
- [ ] [P1-B] Define canonical Web display behavior for pairs that have both `EXTENDS` and `INHERITS`.
- [ ] [P1-C] Define raw graph export behavior: keep both edge types, collapse at graph payload generation, or preserve raw edges and add semantic grouping metadata.
- [ ] [P1-D] Define unresolved/external heritage policy for targets such as `Error`, `Component`, DOM/React interfaces, and external package types.
- [ ] [P1-E] Define user-facing terminology for language-specific heritage forms: class extends, interface extends, implements, Go embedding, trait/include/extend/prepend, and normalized inheritance.
- [ ] [P1-F] Define compatibility behavior for MCP/context/impact/MRO consumers that may currently depend on `INHERITS`.
- [ ] [P1-G] Record the selected policy in plan, benchmark, evidence, MCP graph schema docs, generated Web contracts, and Web dashboard wording.

## Phase 2 - Root-Cause Trace Before Fixing

- [ ] [P2-A] For representative `E:\Restaurant_manager` TS files, capture parser/ScopeIR output and verify whether `HeritageFact` is emitted for every audited `extends` / `implements` site.
- [ ] [P2-B] Trace same-file TS interface inheritance such as `AreaWithTableCount extends Area` and `ShiftWithCounts extends Shift` through workspace binding and relationship emission.
- [ ] [P2-C] Trace cross-file/imported TS heritage targets and record whether import binding or name resolution is the failure point.
- [ ] [P2-D] Trace external TS heritage targets such as `Error`, `Component`, and React/DOM interfaces and classify them as external/unresolved by policy.
- [ ] [P2-E] Trace Go embedded struct relationships and classify whether their user-facing label should be `EXTENDS`, `EMBEDS`, `INHERITS`, or a grouped heritage display.
- [ ] [P2-F] Record the exact loss point for each missing heritage source-site class: extraction loss, owner-scope loss, import/name-resolution loss, graph emission loss, payload loss, or UI display loss.

## Phase 3 - Analyzer Coverage

- [ ] [P3-A] Add focused TS fixtures for `class extends`, `class implements`, `interface extends`, same-file targets, cross-file imported targets, and external unresolved targets.
- [ ] [P3-B] Add parser/ScopeIR source-site inventory assertions so every heritage site is counted even if unresolved.
- [ ] [P3-C] Fix TS provider/resolution behavior so resolved in-repo class/interface heritage emits the selected graph relationship contract.
- [ ] [P3-D] Fix or document unresolved/external heritage handling so those source sites are visible in graph audit data.
- [ ] [P3-E] Verify Go embedded struct behavior still emits the expected heritage facts after any contract changes.
- [ ] [P3-F] Run provider parity tests for supported language heritage forms that already have coverage.
- [ ] [P3-G] Update generated contracts or schema docs if the selected policy adds relationship metadata, external target facts, or display-group fields.

## Phase 4 - UI Relationship Display

- [ ] [P4-A] Update dashboard Edge Types counts so compatibility duplicates do not mislead users.
- [ ] [P4-B] Update graph canvas edge conversion so duplicate compatibility heritage pairs do not draw as two unrelated relationships unless explicitly requested.
- [ ] [P4-C] Update color legend labels/tooltips for `EXTENDS`, `IMPLEMENTS`, `INHERITS`, and any Go embedding display label to reflect the selected semantics.
- [ ] [P4-D] Add unit tests for duplicate `EXTENDS` + `INHERITS` source-target pairs.
- [ ] [P4-E] Add unit tests for resolved TS class/interface heritage and unresolved/external TS heritage display.
- [ ] [P4-F] Add e2e coverage proving the Web UI displays heritage counts and relationship toggles according to the selected policy.
- [ ] [P4-G] Add deterministic fixture coverage and, where feasible, a `Restaurant_manager` large-graph smoke check so the real regression is covered.
- [ ] [P4-H] Inspect `avmatrix-web/src/components/Header.tsx`, `avmatrix-web/src/App.tsx`, launcher path handling, and `Start-AVmatrix.html` to confirm the correct Start-screen return target in both dev and packaged launcher modes.
- [ ] [P4-I] Add an icon-first Back arrow/button beside the `AVmatrix` top bar title with an accessible label and keyboard/click activation.
- [ ] [P4-J] Implement the return flow to `Start-AVmatrix.html` without showing an internal reconnect-banner error during the intentional navigation transition.
- [ ] [P4-K] Inspect the app shell layout, `FileTreePanel` container, graph canvas sizing, and responsive CSS to identify where the left dashboard width is currently fixed.
- [ ] [P4-L] Add a visible drag handle on the right edge of the left dashboard and support mouse/pointer dragging to resize the panel within bounded min/max widths.
- [ ] [P4-M] Add unit and/or e2e coverage proving Back navigation, left dashboard resize bounds, and continued dashboard/canvas interaction after resize.

## Phase 5 - Benchmarks and Accuracy Measurement

- [ ] [P5-A] Analyze `E:\AVmatrix-GO` and record raw relationship counts, semantic unique heritage counts, and duplicate compatibility pair counts.
- [ ] [P5-B] Analyze `E:\Restaurant_manager` and record the same heritage metrics.
- [ ] [P5-C] Record TypeScript heritage source-site coverage for `E:\Restaurant_manager`: parser/ScopeIR source sites, resolved graph relationships, unresolved/external sites, and missing sites.
- [ ] [P5-D] Record Go embedded struct heritage coverage and final user-facing display label/counts.
- [ ] [P5-E] Record UI display inventory: visible dashboard edge rows, raw edge counts, displayed semantic counts, and graph canvas duplicate-pair rendering behavior.
- [ ] [P5-F] Record Web UI shell interaction inventory for Back navigation and left dashboard resize: navigation target, reconnect-banner behavior, min/max width bounds, and canvas usable width after resize.
- [ ] [P5-G] Update benchmark and evidence ledgers immediately after each measured slice.

## Phase 6 - Validation

- [ ] [P6-A] Run full Go build before tests.
- [ ] [P6-B] Run focused Go tests for TS heritage extraction/resolution and relationship contract behavior.
- [ ] [P6-C] Run focused tests for MCP/context/impact/MRO behavior if `INHERITS` compatibility or graph payload semantics change.
- [ ] [P6-D] Run full applicable Go test suite for `cmd` and `internal`.
- [ ] [P6-E] Run Web build before Web tests.
- [ ] [P6-F] Run focused Web unit tests for edge type dashboard, legend, and graph adapter behavior.
- [ ] [P6-G] Run full Web unit suite.
- [ ] [P6-H] Run focused Web unit/e2e tests covering top bar Back navigation and left dashboard resize behavior.
- [ ] [P6-I] Run e2e test covering heritage edge display, duplicate-pair behavior, Back navigation, and left dashboard resize behavior.
- [ ] [P6-J] Re-run analyze on `E:\Restaurant_manager` and verify final graph metrics match the accepted contract.

## Phase 7 - Closure

- [ ] [P7-A] Update this plan checklist after each completed slice.
- [ ] [P7-B] Update benchmark ledger with initial, intermediate, and final counts.
- [ ] [P7-C] Update evidence ledger with commands, files changed, tests, and conclusions.
- [ ] [P7-D] Commit each completed implementation slice.
- [ ] [P7-E] Final closure: confirm heritage graph semantics, TS coverage, UI relationship display, Back navigation, left dashboard resize, benchmark, evidence, full build, unit tests, and e2e tests are complete.
