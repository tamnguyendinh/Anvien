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

This is not just a Web UI display issue and not just an analyzer issue. It has two related parts:

- relationship semantics/display: `EXTENDS` and `INHERITS` can represent the same underlying heritage fact, so the UI must not present duplicate compatibility edges as independent codebase facts;
- extraction/resolution coverage: TypeScript heritage sites that exist in the repo must either be connected to resolved graph targets or represented as unresolved/external heritage, not silently disappear.

## Scope

Implementation may touch:

- `internal/providers/tsjs/*`
- `internal/resolution/*`
- `internal/graph/*`
- `internal/contracts/*`
- `internal/mcp/*`
- `avmatrix-web/src/lib/constants.ts`
- `avmatrix-web/src/lib/graph-adapter.ts`
- `avmatrix-web/src/components/FileTreePanel.tsx`
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

## Baseline Findings

- [x] [P0-A] Verify whether `Class=3` and `Constructor=3` in `E:\Restaurant_manager` are real. Result: verified as real source counts for `SSEListener`, `ApiError`, and `ErrorBoundary`.
- [x] [P0-B] Verify whether `EXTENDS=6` and `INHERITS=6` are independent relationships. Result: all `6` source-target pairs are duplicated as both `EXTENDS` and `INHERITS`.
- [x] [P0-C] Trace graph emission source. Result: `emitHeritageCompatibilityEdges` emits `EXTENDS` or `IMPLEMENTS`, then emits `ReferenceInherits` when compatibility is enabled; `ReferenceInherits` maps to `INHERITS`.
- [x] [P0-D] Check TypeScript source heritage sites in `E:\Restaurant_manager`. Result: at least `14` `interface extends` sites and `2` `class extends` sites exist.
- [x] [P0-E] Check whether current graph has TS class/interface heritage edges. Result: no `EXTENDS` / `INHERITS` / `IMPLEMENTS` edges were found involving the audited TS class/interface heritage sites.

## Phase 1 - Semantics Decision

- [ ] [P1-A] Define whether `INHERITS` is a compatibility/normalized edge, a first-class display edge, or both.
- [ ] [P1-B] Define canonical Web display behavior for pairs that have both `EXTENDS` and `INHERITS`.
- [ ] [P1-C] Define raw graph export behavior: keep both edge types, collapse at graph payload generation, or preserve raw edges and add semantic grouping metadata.
- [ ] [P1-D] Define unresolved/external heritage policy for targets such as `Error`, `Component`, DOM/React interfaces, and external package types.
- [ ] [P1-E] Record the selected policy in plan, benchmark, evidence, MCP graph schema docs, and Web dashboard wording.

## Phase 2 - Analyzer Coverage

- [ ] [P2-A] Add focused TS fixtures for `class extends`, `class implements`, `interface extends`, same-file targets, cross-file imported targets, and external unresolved targets.
- [ ] [P2-B] Add source-site inventory assertions so every heritage site is counted even if unresolved.
- [ ] [P2-C] Fix TS provider/resolution behavior so resolved in-repo class/interface heritage emits the selected graph relationship contract.
- [ ] [P2-D] Fix or document unresolved/external heritage handling so those source sites are visible in graph audit data.
- [ ] [P2-E] Verify Go embedded struct behavior still emits the expected heritage facts after any contract changes.
- [ ] [P2-F] Run provider parity tests for supported language heritage forms that already have coverage.

## Phase 3 - UI Relationship Display

- [ ] [P3-A] Update dashboard Edge Types counts so compatibility duplicates do not mislead users.
- [ ] [P3-B] Update graph canvas edge conversion so duplicate compatibility heritage pairs do not draw as two unrelated relationships unless explicitly requested.
- [ ] [P3-C] Update color legend labels/tooltips for `EXTENDS`, `IMPLEMENTS`, and `INHERITS` to reflect the selected semantics.
- [ ] [P3-D] Add unit tests for duplicate `EXTENDS` + `INHERITS` source-target pairs.
- [ ] [P3-E] Add e2e coverage proving the Web UI displays heritage counts and relationship toggles according to the selected policy.

## Phase 4 - Benchmarks and Accuracy Measurement

- [ ] [P4-A] Analyze `E:\AVmatrix-GO` and record raw relationship counts, semantic unique heritage counts, and duplicate compatibility pair counts.
- [ ] [P4-B] Analyze `E:\Restaurant_manager` and record the same heritage metrics.
- [ ] [P4-C] Record TypeScript heritage source-site coverage for `E:\Restaurant_manager`: source sites, resolved graph relationships, unresolved/external sites, and missing sites.
- [ ] [P4-D] Record UI display inventory: visible dashboard edge rows, raw edge counts, displayed semantic counts, and graph canvas duplicate-pair rendering behavior.
- [ ] [P4-E] Update benchmark and evidence ledgers immediately after each measured slice.

## Phase 5 - Validation

- [ ] [P5-A] Run full Go build before tests.
- [ ] [P5-B] Run focused Go tests for TS heritage extraction/resolution and relationship contract behavior.
- [ ] [P5-C] Run full applicable Go test suite for `cmd` and `internal`.
- [ ] [P5-D] Run Web build before Web tests.
- [ ] [P5-E] Run focused Web unit tests for edge type dashboard, legend, and graph adapter behavior.
- [ ] [P5-F] Run full Web unit suite.
- [ ] [P5-G] Run e2e test covering heritage edge display and duplicate-pair behavior.
- [ ] [P5-H] Re-run analyze on `E:\Restaurant_manager` and verify final graph metrics match the accepted contract.

## Phase 6 - Closure

- [ ] [P6-A] Update this plan checklist after each completed slice.
- [ ] [P6-B] Update benchmark ledger with initial, intermediate, and final counts.
- [ ] [P6-C] Update evidence ledger with commands, files changed, tests, and conclusions.
- [ ] [P6-D] Commit each completed implementation slice.
- [ ] [P6-E] Final closure: confirm heritage graph semantics, TS coverage, UI display, benchmark, evidence, full build, unit tests, and e2e tests are complete.
