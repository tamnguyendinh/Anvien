# Anvien Test File Unresolved Separation Plan

Date: 2026-06-02

Status: Active

Companion files:

- Evidence ledger: [2026-06-02-anvien-test-file-unresolved-separation-evidence.md](2026-06-02-anvien-test-file-unresolved-separation-evidence.md)
- Benchmark ledger: [2026-06-02-anvien-test-file-unresolved-separation-benchmark.md](2026-06-02-anvien-test-file-unresolved-separation-benchmark.md)

## Master Rules

1. Follow active workspace and repository instructions, including generated `AGENTS.md`; this plan records product work and validation, it does not replace repository rules.
2. Do not use Anvien for doc-only commits.
3. Use Anvien for implementation slices that inspect code ownership, graph impact, file projection behavior, graph-health behavior, or Web UI impact.
4. Refresh graph evidence with `anvien analyze --force` before graph-based implementation evidence.
5. Run impact analysis before editing graph builders, file projection logic, resolution metrics, API handlers/contracts, Web graph views, or shared test classification code.
6. Do not delete unresolved/ResolutionGap facts just to reduce counts; preserve raw diagnostic data and change default classification, ranking, and display behavior.
7. Treat generated `AGENTS.md`, `CLAUDE.md`, and `.claude/skills/anvien/**` as generated output only.
8. Run the full build before testing. For this repo the full build gate is `powershell -ExecutionPolicy Bypass -File anvien-launcher\build.ps1`.
9. If Web UI behavior changes, include relevant Web/e2e validation or record why no browser validation is required.
10. Record evidence as each evidenced task completes.
11. Record benchmarkable inventory/count changes as each benchmarkable task completes.
12. Run `anvien detect-changes --repo Anvien --scope all` before implementation commits.
13. Commit each completed implementation slice after evidence and benchmark ledgers are updated.

## Goal

Separate test-file unresolved references from the default production unresolved signal so test files no longer dominate unresolved hotspots or clutter the default graph with low-value ResolutionGap detail.

The target behavior is:

- test files are visible as their own `Test File` region/node type;
- default graph/file hotspot views show that a file is a test file;
- default views show the file/symbol/API/flow relationships that the test covers;
- default views do not expand unresolved details inside test files;
- raw test unresolved data remains available through explicit drill-down or filter when debugging tests.

## Problem

Current analyze output shows unresolved hotspots are dominated by test surfaces instead of production surfaces.

Latest observed hotspots:

| File | Unresolved | Risk |
|---|---:|---|
| `internal/mcp/server_test.go` | 1445 | high |
| `internal/cli/command_test.go` | 1119 | high |
| `internal/analyze/analyze_test.go` | 1052 | high |
| `internal/resolution/resolution_test.go` | 934 | high |
| `anvien-web/e2e/graph-orientation-labels.spec.ts` | 856 | high |

That means the current default unresolved signal is mixing fundamentally different things:

- production unresolved that can indicate graph/analyzer/product issues;
- test unresolved caused by assertions, mocks, fixtures, intentionally invalid inputs, test helpers, and test frameworks;
- non-actionable unresolved such as builtins and standard library calls;
- unknown unresolved that still needs investigation.

For the default graph, test files do not need unresolved children. Their primary value is to show: "I am a test file" and "I test these files/symbols/API routes/flows."

## Scope

In scope:

- file classification used by file projection, graph health, hotspot ranking, and Web file views;
- unresolved and ResolutionGap count aggregation by file kind/app layer;
- default hotspot ranking and default graph visibility;
- Web UI display behavior for test files and unresolved/ResolutionGap nodes;
- optional explicit drill-down/filter for test unresolved;
- CLI/API contracts used by Web file map/detail/graph views;
- tests and benchmarks proving default unresolved signal is no longer dominated by test files.

Out of scope:

- removing ResolutionGap nodes from the canonical graph;
- hiding production unresolved references;
- changing parser/source-site resolution semantics for production files;
- changing test execution behavior;
- changing unrelated graph layout or visual style work;
- forcing all test files to resolve every assertion/mock/helper.

## Requirements

1. Test files must be classified as a stable file category in file projection and Web output.
2. Default unresolved hotspot ranking must not be dominated by test/e2e files.
3. Default graph/file views must collapse or hide unresolved details inside test files.
4. Test file nodes must remain visible and identifiable as `Test File`.
5. Test file relationships to tested targets must remain visible when available.
6. Raw test unresolved data must remain accessible through explicit drill-down, debug view, JSON field, or filter.
7. Metrics must separate at least:
   - production/actionable unresolved;
   - test unresolved;
   - non-actionable unresolved;
   - unknown unresolved.
8. Existing unresolved raw counts must remain traceable for audit and benchmark comparison.
9. API/Web contract changes must be additive where possible and covered by tests.
10. Generated output or docs must not claim unresolved was "fixed" by deletion; the behavior change is classification and default visibility.

## Invariants

1. Canonical symbol/source-site graph facts remain the source of truth.
2. File projection can derive display/ranking groups from canonical facts, but must not rewrite canonical truth.
3. Test unresolved is diagnostic data, not the default production graph signal.
4. `Test File` visibility is more important in default views than every unresolved child detail.
5. A user debugging tests must still be able to find raw unresolved evidence.
6. Default hotspot ranking must optimize for actionable production investigation.
7. Benchmark totals must distinguish raw unresolved from default-visible unresolved.

## Technical Direction

Implementation should first find the actual owners before editing. Likely ownership areas include file projection, graph-health/file-hotspot aggregation, HTTP/API response shape for file detail, and Web graph/file map rendering. Do not assume the exact owner names before source inspection.

Preferred model:

```text
rawUnresolvedCount
productionUnresolvedCount
testUnresolvedCount
nonActionableUnresolvedCount
unknownUnresolvedCount
defaultVisibleUnresolvedCount
```

Default hotspot ranking should use `defaultVisibleUnresolvedCount` or `productionUnresolvedCount`, not raw unresolved count. Test files should still be findable by a file kind/app layer filter.

For Web graph display:

```text
Test file node/region
  -> tested file/symbol/API/flow edges
  -> no default unresolved child expansion
  -> explicit "show test unresolved" drill-down/filter when needed
```

For CLI/API output, prefer additive fields over breaking existing fields:

- keep existing `unresolved` raw count if clients depend on it;
- add separated counts for default ranking/display;
- add `isTestFile` or reuse existing `kind=test`/`appLayer=backend_test|api_test|frontend_test`;
- add a clear visibility/ranking field if Web UI needs it.

## Definition Of Done

The plan is complete when:

1. analyze/file projection records separated unresolved counts for test and production files;
2. default hotspot ranking no longer lists test/e2e files as top unresolved hotspots solely because of test unresolved detail;
3. Web UI default graph/file views render test files as `Test File` and do not expand test unresolved detail by default;
4. raw test unresolved remains available through explicit drill-down or filter;
5. tests cover classification, metric separation, hotspot ranking, and Web/API behavior affected by the change;
6. before/after benchmarks record raw unresolved count, default-visible unresolved count, test unresolved count, production unresolved count, and top hotspot composition;
7. full build, focused tests, any required Web/e2e validation, analyze regeneration, and detect-changes evidence are recorded;
8. implementation work is committed after evidence and benchmark ledgers are updated.

## Phase Checklist

- [ ] [P0-A] Establish baseline and owner evidence.
  - Goal: record the current unresolved hotspot problem and identify source owners before edits.
  - Work Steps: run `anvien analyze --force`; record file counts, graph counts, unresolvedFiles, top hotspot list, and whether each hotspot is test/e2e; use direct Anvien command selection to inspect likely owners for file projection, graph-health/file-hotspots, HTTP file detail, and Web graph/file map rendering; inspect source manually after owner discovery.
  - Implementation Gate: no code edits in this phase.
  - Acceptance: evidence records current hotspot composition and owner files; benchmark records baseline raw unresolved and hotspot composition.

- [ ] [P1-A] Define test-file classification source of truth.
  - Goal: make test-file detection stable and shared by projection/ranking/UI.
  - Work Steps: inspect existing file kind, appLayer, path-pattern, and provider classification fields; decide whether to reuse `kind=test`/test app layers or add a small helper; cover Go `_test.go`, TS/JS `.spec`/`.test`, e2e paths, and current backend/api/frontend test app layers; add unit tests for classification boundaries.
  - Implementation Gate: do not change unresolved counts yet; only establish reliable classification.
  - Acceptance: source has one clear classification path that can tell production files from test/e2e files without ad hoc UI-only checks.

- [ ] [P1-B] Separate unresolved metric buckets.
  - Goal: preserve raw unresolved data while creating metrics that distinguish production/test/non-actionable/unknown unresolved.
  - Work Steps: update file projection aggregation to compute raw unresolved plus separated buckets; classify ResolutionGap rows by file classification and existing actionability/classification metadata; keep old fields if compatibility requires; add tests proving raw totals stay traceable while production/test buckets split correctly.
  - Implementation Gate: do not hide data in UI until the API/projection contract can represent both raw and separated counts.
  - Acceptance: JSON/API/file summaries expose separated unresolved metrics, and tests prove test unresolved is not counted as production unresolved.

- [ ] [P1-C] Change hotspot ranking to actionable default signal.
  - Goal: stop test files from dominating default unresolved hotspot lists.
  - Work Steps: find current hotspot ranking logic; switch default unresolved ranking to production/actionable/default-visible unresolved; keep optional raw/test ranking mode if current command/UI supports sorting by raw unresolved; update CLI/API tests and benchmark outputs.
  - Implementation Gate: raw unresolved count must remain available somewhere for audit.
  - Acceptance: default top hotspots are ranked by production/actionable unresolved; test files can still appear in a test-specific or raw-unresolved view.

- [ ] [P2-A] Update Web file map and graph default display.
  - Goal: make test files visible as test files without expanding unresolved child detail by default.
  - Work Steps: inspect Web graph/file map components; add UI treatment for test file nodes/rows using existing visual language; hide/collapse test unresolved child nodes in default graph view; ensure test file -> tested target relationships remain visible; keep text short and avoid explanatory UI copy.
  - Implementation Gate: if API fields needed by UI are missing, return to P1-B/P1-C instead of hard-coding path checks in UI.
  - Acceptance: default Web view shows `Test File` identity and linked tested targets, while test unresolved detail is not rendered as default graph clutter.

- [ ] [P2-B] Add explicit test unresolved drill-down/filter.
  - Goal: preserve access to raw test unresolved for debugging without polluting the default view.
  - Work Steps: add or reuse a filter/toggle/detail section for test unresolved; make it off by default; display bucket counts and samples only after explicit user action; keep raw evidence traceable to source-site/ResolutionGap IDs.
  - Implementation Gate: do not make the toggle affect production unresolved visibility.
  - Acceptance: a user can intentionally inspect test unresolved details, but default graph/hotspot views remain production-focused.

- [ ] [P3-A] Update API/contract and unit tests.
  - Goal: prevent drift across backend projection, API output, and Web consumers.
  - Work Steps: update generated Web contracts if necessary; add backend tests for separated counts and ranking; add Web/client tests for display defaults; update existing snapshots or fixtures only when behavior changed intentionally.
  - Implementation Gate: run full build before tests.
  - Acceptance: affected package tests pass and prove old behavior would fail.

- [ ] [P3-B] Run Web/e2e validation for UI behavior.
  - Goal: verify the visible graph/file behavior if Web UI changes.
  - Work Steps: run the relevant Web build/test/e2e path; use browser or Playwright validation if a local UI change is present; capture screenshot or trace evidence when layout/visibility changes.
  - Implementation Gate: if no Web UI behavior changes are made, record that this phase is not applicable and why.
  - Acceptance: evidence proves test unresolved is hidden/collapsed by default and visible only through explicit action.

- [ ] [P4-A] Analyze, benchmark, and compare before/after.
  - Goal: prove the metric/ranking behavior improved without deleting data.
  - Work Steps: run `anvien analyze --force`; record raw graph counts, raw unresolved counts, separated unresolved buckets, default-visible unresolved count, top hotspot composition, and generated file projection stats; compare against B0 baseline.
  - Implementation Gate: do not mark complete if raw unresolved disappears without a traceable replacement.
  - Acceptance: benchmark shows raw unresolved remains measurable, test unresolved is separated, and default hotspots are no longer dominated by test files.

- [ ] [P4-B] Detect changes, record closure evidence, and commit.
  - Goal: close the implementation slice with synchronized plan/evidence/benchmark state.
  - Work Steps: update evidence and benchmark ledgers; run `anvien detect-changes --repo Anvien --scope all`; review affected processes and high-risk files; commit the completed implementation slice.
  - Implementation Gate: do not commit until detect-changes and ledgers are updated.
  - Acceptance: commit hash is recorded in evidence and the plan checklist reflects completed tasks.

## Risk Notes

- Hiding test unresolved by deleting ResolutionGap facts would make graph diagnostics less honest; this plan requires classification and default visibility changes instead.
- Ranking changes can break users who rely on raw unresolved sorting; keep raw/test-specific access where practical.
- UI-only path checks would drift from backend truth; prefer backend/file-projection classification fields.
- Test files sometimes expose real production graph issues through imports or tested-target edges; keep test-to-target relationships visible.
- Existing counts may appear to drop in default views; benchmark must explain raw versus default-visible counts clearly.
