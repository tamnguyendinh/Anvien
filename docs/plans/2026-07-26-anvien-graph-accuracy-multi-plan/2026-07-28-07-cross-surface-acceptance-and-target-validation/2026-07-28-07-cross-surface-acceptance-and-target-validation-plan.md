# Anvien Cross-Surface Acceptance and Target Validation Plan

## Metadata

- Date: `2026-07-28`
- Last revised: `2026-08-10`
- Status: `P0 complete / dependency-blocked`
- Plan: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-07-cross-surface-acceptance-and-target-validation/2026-07-28-07-cross-surface-acceptance-and-target-validation-plan.md`
- Evidence: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-07-cross-surface-acceptance-and-target-validation/2026-07-28-07-cross-surface-acceptance-and-target-validation-evidence.md`
- Benchmark: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-07-cross-surface-acceptance-and-target-validation/2026-07-28-07-cross-surface-acceptance-and-target-validation-benchmark.md`
- Actual status: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-07-cross-surface-acceptance-and-target-validation/2026-07-28-07-cross-surface-acceptance-and-target-validation-actual-status.md`
- Contract: `docs/contracts/graph-accuracy-contract.md`
- Roadmap: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-26-anvien-graph-accuracy-roadmap.md`
- Predecessor: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-plan.md`
- Successor: `none — terminal child`

## Goal

Prove that the graph produced by the current Anvien production path is more correct and precise than the recorded baseline, across repeated normal analyzes, the five bounded defect families, every reader actually affected by the campaign, and the real runtime boundaries those readers expose.

## Rules

- Complete P0 actual status before P7 validation begins.
- Treat `reports/problem/2026-07-26-tinh-chinh-cac-van-de-cua-anvien` as the problem-origin record. Its measured defects and bounded targets are evidence; its proposed implementation design is DRAFT and is not execution authority.
- Use the current repository source, build, tests, runtime, command grammar, and normal graph output. Validation does not introduce a second implementation or a substitute execution path.
- Child 07 is validation-only. It may update plan ledgers, reports, reusable QA scripts when an affected UI boundary requires them, and benchmark artifacts. It must not repair production behavior.
- A failed check reopens the exact owning slice in Child 01 through Child 06. The affected child must implement and accept the repair before Child 07 reruns the failed check.
- Open P7-A only after all six predecessor child handoffs identify accepted commits, completed evidence, and the exact source/config/analyzer basis to validate.
- Run `anvien analyze E:\Anvien --force` before graph-based validation. Before any validation-owned implementation artifact is edited, record its file-detail and impact evidence when applicable.
- Run production code changes before test changes in the owning child. Child 07 consumes accepted behavior; it does not make a test pass by changing production or weakening the oracle.
- Run the full build before runtime validation. If an affected surface is browser-visible, validate the real built runtime and store reusable Playwright evidence under `Reports/qa/playwright/` as JSON and Markdown.
- Derive the P7-C reader denominator from accepted child diffs, impact evidence, persisted fields, and real consumers. Do not impose a fixed prior list on unaffected readers, and do not omit an affected reader merely because it was not named in an earlier document.
- Record benchmark values only when measured with the same corpus, configuration, analyzer build, machine, and cache policy as the accepted baseline.
- Analyze `E:\cheapapp.org` in place. Preserve its source and worktree; its normal operational output remains under `E:\cheapapp.org\.anvien`. Keep all Anvien-owned reports, probes, fixtures, and debug material inside `E:\Anvien`.
- Scanner handling for the eight separately reported path omissions remains outside this child. Those omissions cannot be counted as a success or silently added to the graph-accuracy denominator.
- Update the matching checklist, actual-status row, evidence ID, and benchmark row immediately when a result changes.
- Obtain Supervisor PASS before accepting each P7 slice and before terminal closure.
- Run Anvien detect-changes before committing validation work and keep each P7 slice in its own commit.
- Every new or touched file owns one primary responsibility. No validation artifact may reimplement analyzer, resolver, persistence, or reader policy.

## Problem

The problem-origin report records five bounded accuracy defects in the graph generated from the real target:

1. Two of four same-name `time`/`now` declarations survive as distinct graph identities.
2. None of six destructured binding leaves is represented through the required definition/binding flow.
3. None of 21 bounded direct exports has the required export semantics.
4. Neither of two bounded calls through a barrel resolves to its terminal function.
5. `Promise`, `Math.max`, and `Math.min` do not have truthful external or external-capability outcomes and are exposed as in-repository gaps.

Children 01 through 06 own the corresponding production corrections. This terminal child exists because local tests alone cannot prove that the final graph is deterministic, structurally intact, correct at the original source sites, visible through affected consumers, and operationally acceptable.

## Scope

In scope:

- repeated normal analyze determinism on identical source/config/analyzer inputs;
- canonical node and relationship record comparison without assuming an unapproved identity shape;
- duplicate-loss, missing-endpoint, orphan-reference, and occurrence-conservation checks supported by the accepted contract;
- exact target validation for the five bounded defect families;
- Graph JSON and Ladybug checks for corrected facts;
- the complete set of CLI, MCP, HTTP, Web, cache, embedding, registry, group, process, community, rename, context, impact, or other readers proven affected by accepted changes;
- full-build, nearest-real-boundary, runtime, visual, and performance evidence required by the affected-surface inventory;
- final target-boundary and worktree-integrity evidence.

Out of scope:

- production fixes inside Child 07;
- a scanner remediation or a claim of global TypeScript compiler conformance;
- reader work that has no demonstrated dependency on a changed field, identity, relationship, outcome, or persistence projection;
- target source edits, target-hosted reports, or target-derived fixtures;
- a new production validation package or duplicated semantic algorithms created only to satisfy acceptance;
- prescribing a graph-writing mechanism before an owning-child evidence gate establishes that it is needed.

## Requirements

1. P7-A must bind every comparison run to the same source commit, configuration, analyzer build, command, and normalization rules.
2. P7-A must use normal supported analyze invocations. A failed invocation is recorded as a failure and contributes no graph result to the determinism denominator.
3. Determinism compares the canonical facts established by the accepted Child 01 and Child 02 contracts. If relationship IDs are not part of that accepted contract, compare normalized relationship records rather than inventing IDs here.
4. Integrity checks must prove that corrected source occurrences are not lost, relationship endpoints exist, and persisted projections do not silently drop conflicting records.
5. P7-B must use the exact original source sites and record numerator, denominator, source evidence, graph evidence, and outcome for each defect family.
6. Ambient acceptance allows a correctly resolved external target or an explicit external-capability outcome supported by the accepted Child 06 contract. None of the three target sites may remain an in-repository analyzer gap.
7. P7-C must create an affected-surface inventory before running parity. Each row states why the surface is affected, the field or behavior checked, its real boundary, and the evidence result.
8. P7-C validates the observable behavior of existing production surfaces. It does not add production oracle files or duplicate domain logic.
9. Runtime and UI evidence is required only for an affected runtime/UI surface, but once such a surface is in the inventory it cannot be replaced with a source-only assertion or a development-server shortcut.
10. Performance acceptance records at least analyze duration, graph size, and peak memory when the campaign changes them; Ladybug load and query latency are included when the accepted affected-surface inventory reaches those paths. Any budget must already be recorded before the final run; Child 07 cannot invent a favorable threshold afterward.
11. Every failed terminal row maps back to one owning child and evidence ID. Child 07 never changes the expected result to accept current output.

## Acceptance Criteria

- All six predecessor handoffs are accepted and identify the exact commits under test.
- At least five normal analyzes of identical inputs produce equal canonical fact sets under the accepted normalization rules.
- Integrity checks report zero unexplained lost occurrences, zero missing relationship endpoints, and zero orphan references in every affected persisted projection.
- The terminal target table passes exactly:

| Defect | Baseline | PASS |
|--------|---------:|-----:|
| Same-name identity | 2/4 | 4/4 |
| Binding patterns | 0/6 | 6/6 |
| Direct exports | 0/21 | 21/21 |
| Barrel calls | 0/2 | 2/2 |
| Ambient sites | 0/3 | 3/3 correct external/capability outcomes |

- Every affected-surface inventory row passes at its nearest real boundary; every unaffected omission has explicit evidence showing why it is outside the denominator.
- Full build passes before runtime checks. Any affected user-visible surface passes on the real built runtime with reusable JSON, Markdown, and visually inspected evidence.
- Final performance measurements use the accepted baseline method and remain within its pre-recorded budget, or the campaign remains blocked with the measured regression.
- The target source/worktree is unchanged except for normal ignored operational output explicitly recorded by the command boundary.
- P7 contains no production repair. Any failed invariant is closed in its owning child and then rerun here.
- Supervisor accepts P7-A, P7-B, P7-C, and the terminal campaign claim.

## Checklist

- [x] P0-A: Complete actual status before validation work.
  - Goal: classify the real terminal-validation state without treating a DRAFT implementation proposal as authority.
  - Work Steps:
    1. Read the working rules, planner skill/templates, problem-origin report, roadmap, contract, all predecessor handoff requirements, and all four Child 07 ledgers in full.
    2. Refresh the Anvien graph, record documentation relationship evidence, classify missing predecessor results, and update P7-A/P7-B/P7-C from the measured report denominators.
  - Implementation Gate: no P7 validation run starts until the actual-status file records a final P0 decision.
  - Acceptance: the status matrix identifies the baseline defects, missing handoffs, validation-only boundary, target boundary, affected-surface discovery rule, and next action through `E0-P0A-RULE1`, `E0-P0A-PLANNER1`, `E0-P0A-REPORT1`, `E0-P0A-REPORT2`, `E0-P0A-REPORT3`, `E0-P0A-BOUNDARY1`, `E0-P0A-GRAPH1`, `E0-P0A-FD1`, `E0-P0A-DOC1`, `E0-P0A-STATUS1`, and `E0-P0A-DIFF1`.

### P7: Terminal graph-accuracy acceptance

- Phase Goal: independently prove determinism, structural integrity, bounded target accuracy, affected-surface behavior, runtime behavior, and measured performance.
- Phase Boundary:
  - In scope: validation evidence and reports for accepted production work.
  - Out of scope: production repair or new semantic infrastructure.
  - Dependencies: six accepted predecessor handoffs and a fresh Child 07 actual-status refresh.
- Phase Implementation Rule: execute `P7-A`, accept and commit its evidence, refresh actual status, then execute `P7-B`, then `P7-C`. A failure returns to its owning child.
- Ordered Slice List:
  - P7-A: Determinism, integrity, and repeated normal analyze.
  - P7-B: Five-family terminal target acceptance.
  - P7-C: Affected-surface, runtime, and performance acceptance.

- [ ] P7-A: Determinism, integrity, and repeated normal analyze.
  - Goal: prove that accepted graph facts are stable and structurally complete across repeated supported analyzes.
  - Scope Boundary:
    - Editable: Child 07 ledgers and Anvien-owned validation reports.
    - Inspect-only: production source, tests, build output, graph output, and accepted predecessor evidence.
    - Preserve-only: production behavior and target source.
    - Out of scope: repairs and new production validation packages.
  - Non-Goals: fault-injection architecture, broad reader changes, or a new graph-writing contract.
  - Pre-flight Questions:
    - Data source: exact commits, source/config/analyzer inputs, and canonical fact definitions from all six accepted handoffs.
    - Display permission: N/A — this slice has no visual behavior.
    - DB read flow: read the normal Graph JSON and Ladybug results produced by each successful analyze.
    - DB write flow: normal supported analyze output only; validation reports stay in Anvien.
    - Render location: evidence and benchmark ledgers.
    - UI behavior flow: N/A — nearest boundaries are analyze, graph loading, and graph queries.
    - Docker runtime: N/A unless the accepted build contract requires it for the normal packaged analyzer.
    - Playwright target: N/A — no browser-visible behavior is accepted in this slice.
    - Behavior test: five or more identical-input analyzes, canonical fact equality, occurrence conservation, duplicate detection, endpoints, and orphan references.
    - Cleanup/quarantine: remove failed or superseded report attempts; retain the final accepted result only.
    - External side effects: repository-local operational graph output.
    - N/A notes: no DB-backed product UI is changed.
  - Work Steps:
    1. Verify six handoffs and fix the comparison tuple for the run set: source commit, config, analyzer build, command, machine policy, and normalization rules.
       - UI flow check: N/A — no UI surface.
       - DB/data flow check: all runs use identical graph inputs and normal output paths.
       - Render location check: Child 07 evidence ledger.
       - Mini QA: run the nearest real analyze and graph-reader boundaries after the full build.
       - Evidence target: `E7-P7A-INPUT1`, `E7-P7A-BUILD1`.
    2. Execute at least five successful normal analyzes, compare canonical facts, and run integrity checks against Graph JSON and Ladybug.
       - UI flow check: N/A — no UI surface.
       - DB/data flow check: compare successful results only; a command failure is a failed run, never an accepted empty result.
       - Render location check: evidence and benchmark ledgers.
       - Mini QA: inspect actual command results and representative corrected records.
       - Evidence target: `E7-P7A-ANALYZE1`, `E7-P7A-DETERMINISM1`, `E7-P7A-INTEGRITY1`.
  - Implementation Gate:
    - All six predecessor handoffs and the accepted canonical comparison contract exist.
    - Full build passes before analyze validation.
  - Acceptance:
    - Source: comparison inputs and normalization rules are exact and fixed for the run set.
    - Runtime/UI: normal analyze and graph-reader boundaries succeed; UI is N/A.
    - DB/data: canonical facts are equal across at least five successful runs; no unexplained lost occurrences, missing endpoints, or orphan references exist.
    - Behavior test: repeated analyze and integrity checks pass without weakening the accepted contract.
    - Cleanup/quarantine: only the final accepted evidence remains.
    - Evidence IDs: `E7-P7A-INPUT1`, `E7-P7A-BUILD1`, `E7-P7A-ANALYZE1`, `E7-P7A-DETERMINISM1`, `E7-P7A-INTEGRITY1`, `E7-P7A-REVIEW1`, `E7-P7A-DETECT1`, `E7-P7A-COMMIT1`.
    - Actual-status rows refreshed: determinism, integrity, and P7-B opening gate.
  - Evidence Targets: comparison tuple, run results, canonical hashes/records, integrity counts, Supervisor verdict, detect-changes, and commit.
  - Actual-status Update: mark P7-A correct or record the exact owning child reopened by the failed invariant.
  - Commit Boundary: commit accepted P7-A evidence and ledger updates only.

- [ ] P7-B: Five-family terminal target acceptance.
  - Goal: prove the five report-origin defects at their original real-source sites.
  - Scope Boundary:
    - Editable: Child 07 ledgers and Anvien-owned target-validation reports.
    - Inspect-only: accepted implementation, target source, and target graph output.
    - Preserve-only: target source/worktree and Anvien production behavior.
    - Out of scope: scanner remediation, target fixtures, and production fixes.
  - Non-Goals: a broad target-quality claim beyond the five bounded denominators.
  - Pre-flight Questions:
    - Data source: recorded target commit/worktree state and exact report-origin source sites.
    - Display permission: N/A unless an affected Web reader is deferred to P7-C.
    - DB read flow: read the normal target Graph JSON and Ladybug outputs plus existing real query boundaries.
    - DB write flow: target's normal operational analyze output only.
    - Render location: all reports stay under Anvien.
    - UI behavior flow: N/A for semantic counting; affected UI projection belongs to P7-C.
    - Docker runtime: N/A for the bounded source/graph oracle.
    - Playwright target: N/A for the bounded source/graph oracle.
    - Behavior test: exact `4/4`, `6/6`, `21/21`, `2/2`, and `3/3` rows.
    - Cleanup/quarantine: remove superseded Anvien-owned attempts; do not write reports or probes into the target.
    - External side effects: target-local normal operational graph output only.
    - N/A notes: scanner omissions remain separately owned.
  - Work Steps:
    1. Record target pre-state, run the supported analyze in place, and bind each expected source site to its graph fact or structured outcome.
       - UI flow check: N/A — semantic source/graph boundary.
       - DB/data flow check: use the target's normal current graph outputs.
       - Render location check: Anvien evidence ledger and report directory.
       - Mini QA: inspect each named source site and its actual graph/query result.
       - Evidence target: `E7-P7B-TARGET1`, `E7-P7B-SITES1`.
    2. Calculate the five terminal rows, inspect failure classifications, verify target post-state, and obtain independent Supervisor review.
       - UI flow check: N/A — P7-C owns affected visual surfaces.
       - DB/data flow check: numerators and denominators are backed by exact source and graph records.
       - Render location check: terminal target table in evidence and benchmark ledgers.
       - Mini QA: verify no passing row depends on a missing, empty, or unrelated result.
       - Evidence target: `E7-P7B-ORACLE1`, `E7-P7B-BOUNDARY1`, `E7-P7B-REVIEW1`.
  - Implementation Gate: P7-A is accepted and the target pre-state is recorded.
  - Acceptance:
    - Source: all exact report-origin sites are accounted for.
    - Runtime/UI: normal target analyze/query boundaries pass; UI is deferred to P7-C only when affected.
    - DB/data: terminal rows equal `4/4`, `6/6`, `21/21`, `2/2`, and `3/3` with truthful external/capability outcomes.
    - Behavior test: no target row is accepted from unrelated, stale, or pass-by-default evidence.
    - Cleanup/quarantine: target source/worktree boundary is preserved and superseded Anvien reports are removed.
    - Evidence IDs: `E7-P7B-TARGET1`, `E7-P7B-SITES1`, `E7-P7B-ORACLE1`, `E7-P7B-BOUNDARY1`, `E7-P7B-REVIEW1`, `E7-P7B-DETECT1`, `E7-P7B-COMMIT1`.
    - Actual-status rows refreshed: all five terminal target rows and P7-C opening gate.
  - Evidence Targets: target pre/post state, per-site source facts, per-site graph facts/outcomes, terminal table, Supervisor verdict, detect-changes, and commit.
  - Actual-status Update: mark each defect row correct or reopen its owning child.
  - Commit Boundary: commit accepted P7-B evidence and ledger updates only.

- [ ] P7-C: Affected-surface, runtime, and performance acceptance.
  - Goal: prove corrected facts through every genuinely affected consumer and measure the final production runtime.
  - Scope Boundary:
    - Editable: Child 07 ledgers, benchmarks, official QA reports, and a reusable Playwright script only when an affected browser surface requires it.
    - Inspect-only: production source, tests, built artifacts, affected readers, and graph output.
    - Preserve-only: production semantics and target source/worktree.
    - Out of scope: production fixes, unrelated reader changes, and production oracle packages.
  - Non-Goals: validating unrelated consumers or declaring parity from a category label without a concrete reader row.
  - Pre-flight Questions:
    - Data source: accepted child diffs, impact evidence, changed persisted fields, and the real readers of those fields.
    - Display permission: preserve existing permissions on every affected user-visible surface.
    - DB read flow: Graph JSON, Ladybug, and each concrete affected reader identified by inventory.
    - DB write flow: normal built-runtime operational state and Anvien-owned evidence only.
    - Render location: native command output, real exposed runtime, benchmark ledger, and official QA directory when applicable.
    - UI behavior flow: required for each affected browser-visible row; otherwise N/A with source evidence.
    - Docker runtime: required when an affected app/Web row is validated.
    - Playwright target: the real built runtime for each affected browser-visible row.
    - Behavior test: per-reader corrected fields/outcomes, runtime error behavior, and same-method performance comparison.
    - Cleanup/quarantine: remove debug and superseded QA; retain final JSON/Markdown/visual evidence.
    - External side effects: local built runtime and repository-local operational outputs.
    - N/A notes: only source-proven unaffected surfaces may be excluded.
  - Work Steps:
    1. Build the final affected-surface inventory from accepted diffs, impact evidence, persistence fields, and concrete reader call paths; record inclusion or evidence-backed exclusion for each candidate.
       - UI flow check: classify every browser-visible affected row.
       - DB/data flow check: map each row to its actual field/outcome source.
       - Render location check: affected-surface inventory in the evidence ledger.
       - Mini QA: exercise representative real reader boundaries after the full build.
       - Evidence target: `E7-P7C-INVENTORY1`, `E7-P7C-BUILD1`.
    2. Validate every included reader at its nearest real boundary; for included UI rows, run the real built runtime, reusable Playwright, and visual inspection.
       - UI flow check: required for included UI rows; N/A only with exclusion evidence.
       - DB/data flow check: each result must match the accepted graph fact/outcome it consumes.
       - Render location check: CLI/MCP/HTTP/runtime output and official QA artifacts.
       - Mini QA: inspect actual records, errors, screenshots, and traces as applicable.
       - Evidence target: `E7-P7C-RUNTIME1`, `E7-P7C-PARITY1`, `E7-P7C-PLAY1` when applicable.
    3. Measure final analyze duration, graph size, peak memory, and every additional affected performance path using the accepted baseline method.
       - UI flow check: N/A unless a measured UI boundary is in the accepted performance scope.
       - DB/data flow check: baseline and final use the same corpus/config/build/machine/cache policy.
       - Render location check: benchmark ledger.
       - Mini QA: inspect raw measurements and aggregation calculations.
       - Evidence target: `E7-P7C-BENCH1`.
  - Implementation Gate:
    - P7-B is accepted.
    - The affected-surface inventory has concrete reader owners and no unresolved candidate.
    - A comparable pre-recorded performance baseline and budget exist for each metric used as a release gate.
  - Acceptance:
    - Source: every included or excluded reader has evidence-backed classification.
    - Runtime/UI: every included real boundary passes; affected UI rows pass on the real built runtime with visual evidence.
    - DB/data: each affected projection matches its accepted graph fact/outcome with zero unexplained loss or drift.
    - Behavior test: no umbrella command substitutes for a concrete affected reader.
    - Cleanup/quarantine: only final official evidence remains.
    - Evidence IDs: `E7-P7C-INVENTORY1`, `E7-P7C-BUILD1`, `E7-P7C-RUNTIME1`, `E7-P7C-PARITY1`, `E7-P7C-PLAY1` when applicable, `E7-P7C-BENCH1`, `E7-P7C-REVIEW1`, `E7-P7C-DETECT1`, `E7-P7C-COMMIT1`.
    - Actual-status rows refreshed: affected-surface inventory, runtime, UI, performance, and terminal campaign decision.
  - Evidence Targets: complete affected-reader inventory, full build, real-boundary outputs, official UI evidence when applicable, performance measurements, Supervisor verdict, detect-changes, and commit.
  - Actual-status Update: mark final affected surfaces and performance correct, or reopen the exact owning child/metric blocker.
  - Commit Boundary: commit accepted P7-C evidence, QA, benchmark, and terminal ledger updates only.

## Risk Notes

- A stable graph hash can still hide a consistently wrong graph. P7-A integrity and P7-B semantic rows are independent gates.
- A correct source fact can be lost in persistence or misread by a consumer. P7-C inventory must follow actual changed fields and call paths.
- A historical reader list can both over-test unrelated surfaces and omit a newly affected one. The denominator must be rebuilt from accepted diffs and impact evidence.
- Target output can look improved when a query returns an empty or unrelated result. Every numerator requires exact source-site and graph-record evidence.
- Performance comparisons are invalid when corpus, configuration, analyzer build, machine, or cache policy changes.
- Child 07 cannot become a repair phase. Any production edit during P7 invalidates the current terminal run and returns work to the owning child.
