# Anvien TypeScript Binding-Pattern Extraction Plan

## Metadata

- Date: `2026-07-28`
- Status: `active / Child 02 closed at 181b8cb8 / P0-A accepted and committed at 17bd4584 / P3-A accepted and committed at b4dbe5c / P3-B accepted and committed at 17254549 / P3-B1 committed-pushed at 01f160e6 / P3-B2 committed-pushed at 19247b4e / P3-B2A REVIEW2 PASS, detect 348/14/14 health 0/0/0, isolated commit-push closure in progress / later slices locked`
- Plan: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-03-typescript-binding-pattern-extraction/2026-07-28-03-typescript-binding-pattern-extraction-plan.md`
- Evidence: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-03-typescript-binding-pattern-extraction/2026-07-28-03-typescript-binding-pattern-extraction-evidence.md`
- Benchmark: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-03-typescript-binding-pattern-extraction/2026-07-28-03-typescript-binding-pattern-extraction-benchmark.md`
- Actual status: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-03-typescript-binding-pattern-extraction/2026-07-28-03-typescript-binding-pattern-extraction-actual-status.md`
- Roadmap: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-26-anvien-graph-accuracy-roadmap.md`
- Contract: `docs/contracts/graph-accuracy-contract.md`
- Predecessor child: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-02-current-graph-persistence-and-reader-consistency/2026-07-28-02-current-graph-persistence-and-reader-consistency-plan.md`
- Successor child: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-04-typescript-export-semantics/2026-07-28-04-typescript-export-semantics-plan.md`

## Goal

Make TypeScript binding-pattern extraction preserve every legal bound identifier as its own declaration and binding, with the correct lexical scope, source ranges, binding path, and downstream reference resolution.

## Rules

- Work directly on the current production source and command path. Keep one production graph path; a slice may replace a component only when its source and impact evidence require that change.
- The current extraction and analyze pipeline is baseline evidence, not an immutable design. Change only the owners that fresh P0 source and impact evidence prove necessary.
- Treat `reports/problem/2026-07-26-tinh-chinh-cac-van-de-cua-anvien` as the problem origin. Its remediation architecture is explicitly DRAFT and is not implementation authority.
- Treat the causal synthesis and final Supervisor report as accepted bounded verification of the six missing bindings only; they do not approve a remediation design or establish global TypeScript conformance.
- Complete P0 actual status before production edits. Refresh the graph, then record `file-detail` and upstream `impact` for every exact file and symbol opened by a slice.
- Implement production behavior before adding or updating tests. Run the full build before validating the nearest real ScopeIR, graph, CLI, or persisted-graph boundary.
- Update checklist, actual status, evidence, and benchmark immediately when a measured state changes.
- Each production and test file owns one primary semantic responsibility. Do not introduce catch-all helper files or refactor unrelated legacy responsibilities.
- Synthetic fixtures belong under the owning package's `testdata/`; they must be independently authored and must not copy target source.
- Analyze `E:\cheapapp.org` in place only for P3-C2. Its source is preserve-only; normal operational output under its own `.anvien` is the only allowed write there.
- Scanner remediation and scanner-quality acceptance are outside this child.
- After each implementation slice: refresh ledgers, remove superseded slice artifacts, obtain Supervisor PASS, run `anvien detect-changes --repo E:\Anvien --scope all`, and commit that slice before opening the next one.
- Hidden fallback is forbidden. Unsupported binding syntax must produce a structured, countable extraction diagnostic instead of disappearing silently.

## Problem

The problem-origin report identifies missing TypeScript destructuring bindings. The accepted bounded causal synthesis later reproduced the first divergence: in the investigated source revision, the TS/JS `variable_declarator` path returned when the declaration name was not a plain identifier. Six legal array-bound names existed in the TypeScript AST but never produced `DefinitionFact` or local-binding facts, so their downstream `.map()` uses became gaps.

That evidence proves a bounded defect, not the full implementation shape. The report's proposed recursive walker and fact schema are candidate directions. P0 must revalidate the current source, enumerate every declaration context and existing fact contract, and identify the smallest production owners before implementation begins.

## Scope

In scope:

- recursive extraction of identifier leaves from array and object binding patterns, including nesting, aliases, rest, defaults, computed keys, and array holes;
- typed binding paths that distinguish property keys from local identifiers and preserve source indexes across holes;
- variable-declaration, parameter, catch, and `for-of`/`for-in` declaration contexts;
- declaration-versus-assignment behavior, type-inference independence, lexical shadowing, and import-binding non-duplication;
- structured diagnostics and metrics for unsupported binding syntax;
- projection of accepted binding facts into the current graph and only those persistence/resolution owners proven affected by P0/impact evidence;
- focused synthetic fixtures and the bounded `6/6` target validation.

Out of scope:

- export syntax and export metadata, owned by Child 04;
- module and re-export resolution, owned by Child 05;
- ambient/external declaration resolution, owned by Child 06;
- blanket changes to CLI, MCP, HTTP, Web, caches, embeddings, groups, processes, or communities without an observed affected contract;
- unrelated artifact/reader policy, scanner remediation, unrelated provider refactors, and target-source changes.

## Requirements

- Every legal bound identifier leaf emits exactly one current production declaration fact and one binding fact or their accepted current-contract equivalents.
- A binding record preserves local name, lexical scope, declaration kind, full range, selection range, binding path, rest/default modifiers, and provenance needed by graph projection.
- Property keys such as `source` in `{source: local}` are not local declarations; `local` is the declaration and its path records `.source`.
- Array holes preserve indexes without creating bindings. For `[a,,b]`, the paths are index `0` and index `2`.
- Default initializers modify a binding; they do not create an extra declaration or an identity path segment.
- Declaration destructuring creates declarations. Assignment destructuring creates writes/references to existing bindings and creates no declaration.
- Emission of a legal binding is independent of whether type inference succeeds for that leaf.
- Variable declarations, parameters, catch bindings, and loop declarations use the same accepted leaf semantics while retaining their distinct lexical-scope rules.
- Existing identifier-only declarations and import bindings remain correct. If imports share any extraction path, evidence must prove each import binding is counted exactly once.
- Unsupported syntax emits a structured diagnostic containing the source site and reason, and increments a measurable unsupported-pattern count.
- Graph projection preserves one declaration/binding occurrence per accepted source leaf and resolves same-name leaves by lexical scope rather than by global name.
- The exact files and symbols are not predetermined by this plan. P0 and per-slice impact evidence select them; if that evidence changes a slice boundary, update the plan before editing code.

## Acceptance Criteria

- The six bounded target names produce `6/6` declarations and `6/6` binding facts with correct range, selection range, scope, and binding path.
- The six bounded downstream `.map()` sites have `0` binding-caused `ResolutionGap` records and resolve to the expected lexical symbols.
- General synthetic fixtures pass for nested array/object patterns, aliases, rest, defaults, holes, parameters, catch bindings, and loop declarations.
- Assignment destructuring creates `0` new declarations; import-binding count delta is `0`; nested shadowing resolves to the expected lexical symbol.
- Unsupported fixture cases have `100%` structured diagnostic coverage and no silent return.
- Existing identifier-only declaration behavior and sibling language providers pass the regression boundary selected by fresh impact evidence.
- The target source/worktree remains unchanged apart from documented normal analyzer-owned operational output.
- Each implementation slice has complete source, build, behavior, boundary, Supervisor, detect-changes, and commit evidence.

## Checklist

- [x] P0-A: Complete current actual status before implementation.
  - Goal: replace historical assumptions with current source, graph, and impact evidence for the binding pipeline.
  - Work Steps:
    1. Refresh the Anvien graph and record current HEAD/source basis.
    2. Read the current TS binding extraction, fact/index, graph projection, and relevant tests in full; inventory every binding context and silent-exit path.
    3. Run `file-detail` and upstream `impact` for every candidate owner and record exact related-file counts and blast radius.
    4. Classify correct, partial, wrong, missing, and blocked rows in actual status; update P3-A work steps if exact ownership differs from this plan.
  - Implementation Gate: the accepted Child 02 handoff is recorded as `E0-P0A-HANDOFF1` and its Pn-C commit is confirmed; no production edit until the actual-status Final P0 Decision permits P3-A.
  - Acceptance: `E0-P0A-HANDOFF1`, `E0-P0A-RULE1`, `E0-P0A-SKILL1`, `E0-P0A-ORIGIN1`, `E0-P0A-VERIFY1`, `E0-P0A-VERIFY2`, `E0-P0A-SCOPE1`, `E0-P0A-GRAPH1`, `E0-P0A-SRC1`, `E0-P0A-FD1`, `E0-P0A-IMPACT1`, `E0-P0A-STATUS1`, and independent `E0-P0A-REVIEW1` are recorded with no unresolved owner boundary; `E0-P0A-DETECT1` and `E0-P0A-COMMIT1` close the isolated documentation/evidence boundary.
  - Current State: Child 02 is closed at entry HEAD `181b8cb800f5fe34fa6fe85ddd359f514ead9fb0`. The exact Set A/B/C manifest and partition checks are durable in the QA report. Independent Supervisor reports `E0-P0A-REVIEW1` and `E0-P0A-REVIEW2` returned `PASS`, including the byte-clean artifact committed as `17bd45845a218dccfa7cb09bde8195a14693bf50`. P3-A opened afterward; `E3-P3A-REVIEW1` and `E3-P3A-REVIEW2` preserve the two REJECT histories, `E3-P3A-REVIEW3` accepts the repaired helper/contract boundary, and the isolated slice is committed at `b4dbe5ccc2d0a77d0986b647c8054427ecca73c4`. P3-B source/test/build invariants were locked by `E3-P3B-REVIEW1`, its sole ledger-currentness rejection was repaired, focused `E3-P3B-REVIEW2` returned `PASS`, final staged `E3-P3B-DETECT1` is `222/13/13` with graph health `0/0/0`, and the slice is committed at `17254549a13ad81a560c18fbcc6ab8fe3ce5f111`. P3-B1 retains `E3-P3B1-REVIEW1` as `REJECT`, closes `P3B1-PARAMETER-CONTEXT-DISAMBIGUATION-1` with focused `E3-P3B1-REVIEW2` `PASS`, records final `E3-P3B1-DETECT1` as `309/17/13` with graph health `0/0/0`, and is committed/pushed at `01f160e6e28ad74c1f379ce5ea47e643a5a14652`. P3-B2 has independent `E3-P3B2-REVIEW1` `PASS`, final staged `E3-P3B2-DETECT1` `242/11/11` with health `0/0/0`, and commit/push closure at `19247b4eb58a4e01a6256f3d63bbb59839644d64`. P3-B2A retains `E3-P3B2A-REVIEW1` as `REJECT`, closes `P3B2A-ASSIGNMENT-TARGET-WRAPPER-CLOSURE-1` with focused `E3-P3B2A-REVIEW2` `PASS`, and awaits only orchestration-owned detect, isolated commit, and immediate push. P3-C and every later slice stay locked until that push succeeds.
  - Commit Boundary: four Child 03 ledgers, one QA report, immutable prior REJECT provenance, and two independent PASS reports; no production/test/fixture/target/runtime path.

### P3: TypeScript binding-pattern extraction

- Phase Goal: preserve each legal binding leaf through extraction, graph projection, and bounded target resolution.
- Phase Boundary:
  - In scope: seven ordered slices listed below.
  - Out of scope: export, module-resolution, ambient/external, scanner, and unrelated consumer changes.
  - Dependencies: P0 complete and predecessor child accepted/committed.
- Ordered Slice List:
  - P3-A: Implement the binding-pattern contract and recursive leaf walker.
  - P3-B: Integrate variable-declaration contexts.
  - P3-B1: Integrate parameter contexts.
  - P3-B2: Integrate catch contexts.
  - P3-B2A: Integrate `for-of`/`for-in` declaration contexts.
  - P3-C: Project and resolve binding occurrences in the current graph.
  - P3-C2: Validate the complete binding contract against the real target.

- [x] P3-A: Implement the binding-pattern contract and recursive leaf walker.
  - Goal: enumerate legal binding leaves, paths, ranges, modifiers, and unsupported syntax without graph emission.
  - Scope Boundary:
    - Editable: add one focused `internal/providers/tsjs/binding_patterns.go` owner for recursive pattern traversal; extend only the exact ScopeIR binding-pattern contract owners in `internal/scopeir/facts.go`, `internal/scopeir/ir.go`, and `internal/scopeir/sort_keys.go` when the accepted contract requires it.
    - Inspect-only: current collector entry/traversal and declaration/type/scoping/reference owners (`extract.go`, `nodes.go`, `definitions.go`, `scopes.go`, `references.go`, `types.go`) plus downstream projection/resolution and graph-accuracy audit readers.
    - Preserve-only: identifier-only declaration behavior, the separate import pipeline in `imports.go`, and all export behavior.
    - Out of scope: variable/parameter/catch/loop declaration-context wiring, assignment destructuring, graph emission/resolution edits, export/module/ambient/scanner work, target validation, and unrelated readers.
  - Non-Goals: do not make assignment destructuring create declarations and do not add resolver behavior.
  - Pre-flight Questions:
    - Data source: current Tree-sitter binding-pattern nodes and current ScopeIR contracts.
    - Display permission: N/A; this slice changes no visible or permission-controlled surface.
    - DB read flow: N/A; the slice reads AST and in-memory ScopeIR only.
    - DB write flow: N/A; it writes in-memory facts and package-local test output only.
    - Render location: N/A; proof is recorded in ScopeIR output and the evidence ledger.
    - UI behavior flow: N/A; this is a non-UI extraction contract.
    - Docker runtime: N/A; no browser/app runtime is changed, but the repository full build remains mandatory.
    - Playwright target: N/A; exercise the real ScopeIR extraction boundary instead.
    - Behavior test: nested array/object, aliases, rest, defaults, holes, computed keys, deterministic order, and unsupported diagnostic.
    - Cleanup/quarantine: keep reusable fixtures under package `testdata/` and remove debug output.
    - External side effects: none.
    - N/A notes: UI, DB, Docker, and Playwright do not participate in this in-memory provider slice.
  - Work Steps:
    1. After independent P0-A acceptance/commit, refresh graph evidence and re-run exact file-detail/impact for the accepted owners; define deterministic path/rest/default/computed-key/hole/provenance and structured unsupported-diagnostic semantics in the ScopeIR contract, then implement the one-responsibility recursive walker in `internal/providers/tsjs/binding_patterns.go` without declaration-context or graph wiring.
       - UI flow check: N/A; no visible flow exists.
       - DB/data flow check: verify deterministic one-leaf-to-one-fact production in memory.
       - Render location check: N/A; inspect ScopeIR output.
       - Mini QA: after the full build in step 2, exercise the real ScopeIR extraction boundary; browser controls are N/A.
       - Evidence target: `E3-P3A-IMPACT1`, `E3-P3A-SRC1`.
    2. After production behavior is correct, add focused tests and fixture artifacts owned by the exact P3-A test file, run the required clean-holder gate before the full build, validate ScopeIR output and diagnostic counts, preserve identifier/import behavior, refresh ledgers, obtain Supervisor PASS, run detect-changes, and commit.
       - UI flow check: N/A; no visible flow exists.
       - DB/data flow check: verify exact paths/ranges/modifiers and no duplicate facts.
       - Render location check: evidence ledger and focused test output.
       - Mini QA: exercise the built non-UI provider boundary and record the command and observed result.
       - Evidence target: `E3-P3A-BUILD1`, `E3-P3A-TEST1`, `E3-P3A-BOUNDARY1`, `E3-P3A-REVIEW1`, `E3-P3A-REVIEW2`, `E3-P3A-REVIEW3`, `E3-P3A-PLAN1`, `E3-P3A-DETECT1`, `E3-P3A-DETECT2`, `E3-P3A-DETECT3`, `E3-P3A-COMMIT1`.
  - Implementation Gate: independent Supervisor accepts the P0-A candidate, P0-A is isolated in a commit, and the exact fact/walker owners have refreshed impact evidence. P0-A completion alone does not open P3-A.
  - Acceptance:
    - Source: every supported leaf has exactly one fact, correct path/range/modifiers, and unsupported syntax is diagnosed.
    - Runtime/UI: N/A; nearest real boundary is ScopeIR extraction after the full build.
    - Data: no graph or persistent-store mutation in this slice.
    - Behavior test: general pattern matrix passes deterministically.
    - Cleanup: no debug or target artifact remains.
    - Evidence IDs: `E3-P3A-IMPACT1`, `E3-P3A-SRC1`, `E3-P3A-BUILD1`, `E3-P3A-TEST1`, `E3-P3A-BOUNDARY1`, `E3-P3A-REVIEW1`, `E3-P3A-REVIEW2`, `E3-P3A-REVIEW3`, `E3-P3A-PLAN1`, `E3-P3A-DETECT1`, `E3-P3A-DETECT2`, `E3-P3A-DETECT3`, `E3-P3A-COMMIT1`.
    - Actual-status rows refreshed: binding contract/walker.
  - Evidence Targets: source diff, pattern matrix, diagnostic count, build, boundary output, Supervisor, detect-changes, commit.
  - Actual-status Update: transition the binding walker row from its P0 classification to `correct`; leave declaration contexts unchanged.
  - Commit Boundary: one P3-A commit after all acceptance evidence passes.
  - Coder Candidate (2026-08-14): the first candidate was rejected by `E3-P3A-REVIEW1` (`reports/Supervisor/rp_supervisor_260814_085244_by_gpt-5_child03_p3a.md`) for valid-`undefined`, malformed/context-invalid, clone-oracle, stale detect, ledger-currentness, and cleanup invariants. Repair remains in this exact P3-A boundary. Fresh repair evidence must replace the superseded first-candidate BUILD/TEST/BOUNDARY/DETECT claims before re-review; this checkbox and `E3-P3A-COMMIT1` remain open, and P3-B is not opened.
  - Repair Checkpoint (2026-08-14): the source/test repair is complete within the same boundary; fresh analyze, file-detail/impact, clean-gated build, focused tests, raw/structured detect artifacts, and cleanup evidence are being recorded under the repair coder report. `E3-P3A-REVIEW1` remains the prior REJECT until a new independent review; `E3-P3A-COMMIT1` remains pending and P3-B stays locked. Roadmap reconciliation is owned by orchestration/docs-owner.
  - Acceptance Checkpoint (2026-08-14): `E3-P3A-REVIEW2` retained the second REJECT because the `292/9` detect pair predated roadmap reconciliation. `E3-P3A-DETECT2` then supplied a fresh durable `294/10` raw/JSON package plus an exact untracked-walker manifest, `E3-P3A-REVIEW3` returned `PASS`, and final `E3-P3A-DETECT3` records `296/10/10` after the planner refresh. `E3-P3A-COMMIT1` is the isolated commit containing this checklist; Git/handoff records its exact hash. P3-B remains locked pending Owner transfer.
    - [x] Preserve the first independent REJECT as `E3-P3A-REVIEW1`.
    - [x] Preserve the second independent REJECT as `E3-P3A-REVIEW2`.
    - [x] Record the fresh durable `294/10` package and untracked-walker manifest as `E3-P3A-DETECT2` without reusing `DETECT1`.
    - [x] Record the final independent PASS as `E3-P3A-REVIEW3`.
    - [x] Refresh checklist/evidence/benchmark/actual-status with the planner gate recorded as `E3-P3A-PLAN1`.
    - [x] Run and record the final orchestration-owned pre-commit detect as `E3-P3A-DETECT3` (`296` changed symbols / `10` changed files / `10` affected files).
    - [x] Create the isolated P3-A commit boundary as `E3-P3A-COMMIT1`; Git/handoff records the exact final hash.

- [x] P3-B: Integrate variable-declaration contexts.
  - Goal: emit binding leaves for variable declarations even when per-leaf type inference is unavailable.
  - Scope Boundary:
    - Editable: `internal/providers/tsjs/definitions.go` only for the `variable_declarator` adapter and leaf-to-Definition/local-Binding emission; `internal/providers/tsjs/extract.go` only for the minimal collector fields plus `result()` plumbing required to return variable `BindingLeaves` and `ExtractionDiagnostics`; `internal/providers/tsjs/extract_test.go` as the focused real-`Extract` test owner; `internal/providers/tsjs/definition_position_inputs_test.go` only for the stale-oracle transition in `TestP1BPreservesBindingPatternBoundary` from zero destructured definitions to exactly one correctly ranged/scoped `left` Definition/local Binding.
    - Inspect-only: accepted walker/ScopeIR contract owners, type inference, assignment/reference extraction, imports, collector traversal, scopes, nodes, and sibling provider adapters.
    - Preserve-only: identifier declarators and non-variable contexts.
    - Out of scope: parameters, catch, loops, and graph projection.
  - Non-Goals: do not invent types or change assignment semantics.
  - Pre-flight Questions:
    - Data source: current variable-declarator AST and accepted P3-A facts.
    - Display permission: N/A; this slice changes no visible or permission-controlled surface.
    - DB read flow: N/A; it reads AST and in-memory ScopeIR only.
    - DB write flow: N/A; it writes in-memory facts and package-local test output only.
    - Render location: N/A; proof is ScopeIR output and ledger evidence.
    - UI behavior flow: N/A; nearest boundary is ScopeIR variable facts.
    - Docker runtime: N/A; no browser/app runtime is changed, but the full build is mandatory.
    - Playwright target: N/A; exercise the real variable-fact boundary instead.
    - Behavior test: array/object/nested/rest/default variables, identifier regression, type-inference failure, assignment non-declaration, and import count.
    - Cleanup/quarantine: keep fixtures under package `testdata/`; remove debug output.
    - External side effects: none.
    - N/A notes: UI, DB, Docker, and Playwright do not participate in this provider slice.
  - Work Steps:
    1. Reconfirm exact impact/ownership, record why `definitions.go` cannot return accepted leaf/diagnostic collections without collector storage/result plumbing, update this boundary before code, then implement production variable-context wiring independently of type inference. Keep `extract.go` changes limited to fields and `result()`; do not change `Extract`, `walkKind`, traversal, or context dispatch.
       - UI flow check: N/A; no visible flow exists.
       - DB/data flow check: verify each legal variable leaf reaches one in-memory declaration/binding fact.
       - Render location check: N/A; inspect ScopeIR output.
       - Mini QA: after the full build in step 2, exercise the variable ScopeIR boundary; browser controls are N/A.
       - Evidence target: `E3-P3B-IMPACT1`, `E3-P3B-SRC1`.
    2. Add focused tests after production code; update only the impact-proven stale oracle `TestP1BPreservesBindingPatternBoundary`; rerun the holder/lock gate and full build after the final test bytes; validate identifiers/type-failure/assignment/import controls, refresh ledgers, obtain Supervisor PASS, detect changes, and commit.
       - UI flow check: N/A; no visible flow exists.
       - DB/data flow check: assert correct ranges/scopes and zero false declaration/import deltas.
       - Render location check: evidence ledger and focused test output.
       - Mini QA: exercise the built variable-fact boundary and record the observed result.
       - Evidence target: `E3-P3B-BUILD1`, `E3-P3B-TEST1`, `E3-P3B-BOUNDARY1`, `E3-P3B-REVIEW1`, `E3-P3B-REVIEW2`, `E3-P3B-DETECT1`, `E3-P3B-COMMIT1`.
  - Implementation Gate: P3-A accepted and committed at `b4dbe5ccc2d0a77d0986b647c8054427ecca73c4`; Owner explicitly transferred P3-B; fresh `E3-P3B-IMPACT1` proves the exact four-file editable boundary, with the fourth owner limited to one stale test function and no helper/sibling-test edit.
  - Acceptance:
    - Source: variable-pattern leaves emit once with correct scope/ranges and survive type-inference failure.
    - Runtime/UI: N/A; ScopeIR output is the real boundary.
    - Data: assignment destructuring adds `0` declarations and import-binding delta is `0`.
    - Behavior test: variable and identifier regression matrices pass.
    - Cleanup: fixture/debug state is confined to approved repo paths.
    - Evidence IDs: `E3-P3B-IMPACT1`, `E3-P3B-SRC1`, `E3-P3B-BUILD1`, `E3-P3B-TEST1`, `E3-P3B-BOUNDARY1`, `E3-P3B-REVIEW1`, `E3-P3B-REVIEW2`, `E3-P3B-DETECT1`, `E3-P3B-COMMIT1`.
    - Actual-status rows refreshed: variable declaration extraction.
  - Evidence Targets: variable fact oracle, type-inference-failure row, assignment/import deltas, build, Supervisor, detect-changes, commit.
  - Actual-status Update: transition the variable-declaration row to `correct`; leave other contexts pending.
  - Commit Boundary: one P3-B commit.
  - Coder Candidate Opening (2026-08-14): Owner transferred P3-B as the sole open Child 03 implementation lane. Fresh graph/source/file-detail/exact-UID impact proves `definitions.go` owns the variable decision/emission adapter, while `extract.go` is unavoidably required only for collector storage and `ScopeIR` result plumbing. `extract_test.go` is the focused real-boundary test owner. The checklist remains open; independent `REVIEW1`, orchestration `DETECT1`, and `COMMIT1` remain reserved.
  - Historical Coder Candidate Blocker (2026-08-14, superseded): the focused P3-B ScopeIR matrix passed `4/4`, but required targeted regression failed only `TestP1BPreservesBindingPatternBoundary` because its historical P1-B guard required destructured `left` to emit zero variable definitions. That failed command remains non-PASS evidence; work stopped before touching the then-unauthorized fourth owner.
  - Coder Candidate Resume (2026-08-14): Orchestration authorized the stale-oracle transition in that exact test function. Fresh graph `98,521/138,234`, file-detail, file impact, and exact-UID upstream impact classify the test owner/function LOW with zero affected files, symbols, modules, processes, or flows. The editable boundary is now exactly four files; no helper or sibling test is opened. P3-B remains unchecked, and independent `REVIEW1`, orchestration `DETECT1`, and `COMMIT1` remain reserved.
  - Coder Candidate Readiness (2026-08-14): final test bytes passed the holder-clean full build, the exact stale oracle, focused `4/4`, targeted `8/8`, and repo-native product regression `61/61` packages. The ScopeIR boundary records exact ranges/scopes/paths plus `100%` inference-miss emission, assignment false declarations `0`, and import-binding delta `0`. Cleanup and diff checks are clean. Candidate state is `READY_FOR_SUPERVISOR`; the P3-B checkbox and `REVIEW1`/`DETECT1`/`COMMIT1` remain open.
  - Supervisor Closure Checkpoint (2026-08-14): `E3-P3B-REVIEW1` retains the first independent `REJECT`, whose only blocker was `P3B-LEDGER-CURRENTNESS-1`; every source, test, build, impact, plan, evidence, and benchmark invariant cleared by that review stayed hash-locked. The focused ledger-only repair appended R14, and `E3-P3B-REVIEW2` returned `PASS` with no residual same-invariant surface. Orchestration owns only the planner refresh, final `E3-P3B-DETECT1`, and isolated `E3-P3B-COMMIT1`; P3-B1 and every later slice remain locked until that commit exists.
    - [x] Preserve the first independent REJECT as `E3-P3B-REVIEW1`.
    - [x] Record the focused independent PASS as `E3-P3B-REVIEW2`.
    - [x] Refresh roadmap, checklist, evidence, benchmark, and actual status without changing locked source/test evidence.
    - [x] Run and record the final orchestration-owned staged pre-commit detect as `E3-P3B-DETECT1` (`222` changed symbols / `13` changed files / `13` affected files; `0` affected processes/flows; current graph health `0/0/0`).
    - [x] Create this isolated P3-B commit boundary as `E3-P3B-COMMIT1`; Git/handoff records the exact final hash.
  - Commit and Successor Opening (2026-08-14): `E3-P3B-COMMIT1` is `17254549a13ad81a560c18fbcc6ab8fe3ce5f111`. The current Owner instruction transfers P3-B1 as the sole open implementation slice. P3-B2 and every later slice remain locked; fresh parameter-owner impact must select the exact editable production/test boundary before code.

- [x] P3-B1: Integrate parameter contexts.
  - Goal: apply accepted leaf semantics to function, method, constructor, and arrow-function parameters.
  - Scope Boundary:
    - Accepted repair boundary: `internal/providers/tsjs/definitions.go` for parameter syntax-role dispatch/emission and tightly related private helpers; `internal/providers/tsjs/extract_test.go` for the focused real-`Extract` former-failure oracle. Both are hash-locked by `E3-P3B1-REVIEW2` pending isolated commit.
    - Accepted dependent boundary: `internal/providers/tsjs/testdata/typescript_scopeir_signature.golden.json`, `internal/providers/vue/extract_test.go`, `internal/providers/vue/testdata/vue_scopeir_signature.golden.json`, `internal/resolution/graph_parity_test.go`, `internal/resolution/testdata/typescript_graph_baseline_counts.golden.json`, and `internal/resolution/testdata/typescript_graph_signature.golden.json` remain byte-locked by `E3-P3B1-REVIEW2`.
    - Inspect-only: accepted walker and scope construction; `internal/providers/vue/extract.go`; `internal/resolution/resolution_test.go` and current resolution production owners that consume generic Definition facts.
    - Preserve-only: variables, catch, loops, calls, references, type inference, imports, every sibling Vue/resolution test function, and all graph projection/resolution source.
    - Out of scope: any production Vue/resolution edit, new graph projection/resolution behavior, P3-C acceptance, and later declaration contexts.
  - Non-Goals: no broad collector rewrite.
  - Pre-flight Questions:
    - Data source: current parameter AST nodes and accepted P3-A contract.
    - Display permission: N/A; no visible or permission-controlled surface changes.
    - DB read flow: N/A; the slice reads parameter AST and in-memory ScopeIR.
    - DB write flow: N/A; only in-memory facts and package-local test output are written.
    - Render location: N/A; inspect parameter ScopeIR and ledger evidence.
    - UI behavior flow: N/A; nearest boundary is parameter ScopeIR.
    - Docker runtime: N/A; no browser/app runtime changes, but the full build is mandatory.
    - Playwright target: N/A; exercise the real parameter-fact boundary.
    - Behavior test: nested/default/rest parameters, lexical scope, shadowing, and sibling-context preservation.
    - Cleanup/quarantine: keep fixtures under package `testdata/`; remove debug output.
    - External side effects: none.
    - N/A notes: UI, DB, Docker, and Playwright are not applicable to this provider slice.
  - Work Steps:
    1. Record fresh parameter-owner impact and implement production parameter wiring.
       - UI flow check: N/A; no visible flow exists.
       - DB/data flow check: verify one binding fact per legal parameter leaf in the correct scope.
       - Render location check: N/A; inspect parameter ScopeIR.
       - Mini QA: after the full build in step 2, exercise the parameter-fact boundary.
       - Evidence target: `E3-P3B1-IMPACT1`, `E3-P3B1-SRC1`.
    2. Add focused tests after code and run the full build plus ScopeIR validation. If repo-native product regression proves that generic consumers now expose the new parameter Definitions, stop for fresh file/symbol impact and a plan boundary update; then repair only the exact stale Vue/resolution expected-output artifacts listed above, rerun a holder-clean full build on final bytes, rerun focused/sibling/product validation, refresh ledgers, obtain Supervisor PASS, detect changes, and commit.
       - UI flow check: N/A; no visible flow exists.
       - DB/data flow check: verify parameter paths/scopes and sibling preservation.
       - Render location check: evidence ledger and focused test output.
       - Mini QA: exercise the built parameter-fact boundary and record the result.
       - Evidence target: `E3-P3B1-BUILD1`, `E3-P3B1-TEST1`, `E3-P3B1-BOUNDARY1`, `E3-P3B1-REVIEW1`, `E3-P3B1-DETECT1`, `E3-P3B1-COMMIT1`.
    3. After `E3-P3B1-REVIEW1` rejected only `P3B1-PARAMETER-CONTEXT-DISAMBIGUATION-1`, prove each `required_parameter`/`optional_parameter` belongs to its owning callable's own formal-parameter surface, preserve unparenthesized TypeScript arrows and explicit `this`, and add the tuple/function-type former-failure test. On final repair bytes, rerun one holder-clean full build plus the new test, focused `4/4`, former failures `4/4`, P3-A/P3-B preservation `10/10`, four-package, and repo-native product gates; rerun the direct JSON audit only if an oracle byte changes or output drift appears; then request `E3-P3B1-REVIEW2`. Do not run detect or open P3-B2/P3-C before `REVIEW2` PASS.
       - Evidence target: repaired `E3-P3B1-SRC1`, refreshed `E3-P3B1-BUILD1`, `E3-P3B1-TEST1`, `E3-P3B1-BOUNDARY1`, and pending `E3-P3B1-REVIEW2`.
  - Implementation Gate: P3-B is accepted and committed. Fresh `E3-P3B1-IMPACT1` must prove the exact TSJS production/test boundary and every dependent-oracle owner/artifact before it is edited; no production Vue/resolution owner may open.
  - Acceptance:
    - Source: every legal parameter leaf emits once in the correct lexical scope.
    - Runtime/UI: N/A; ScopeIR parameter facts are validated after the full build.
    - Data: variable/catch/loop outputs are unchanged; Vue and resolution deltas contain only the parameter Definition/node/`DEFINES` facts produced by the existing generic pipeline, with no projection-source edit and no P3-C acceptance claim.
    - Behavior test: parameter, shadowing, parity, P3-A/P3-B preservation, TSJS/ScopeIR package, and repo-native product matrices pass on final bytes.
    - Cleanup: no unapproved artifact remains.
    - Evidence IDs: `E3-P3B1-IMPACT1`, `E3-P3B1-SRC1`, `E3-P3B1-BUILD1`, `E3-P3B1-TEST1`, `E3-P3B1-BOUNDARY1`, `E3-P3B1-REVIEW1`, `E3-P3B1-REVIEW2`, `E3-P3B1-DETECT1`, `E3-P3B1-COMMIT1`.
    - Actual-status rows refreshed: parameter extraction and exact dependent-oracle boundary.
  - Evidence Targets: parameter oracle, scope/shadowing results, exact Vue/resolution expected-output delta, build, Supervisor, detect-changes, commit.
  - Actual-status Update: transition the parameter row to `correct`.
  - Commit Boundary: one P3-B1 commit.
  - Opening Checkpoint (2026-08-14): P3-B is accepted and committed at `17254549a13ad81a560c18fbcc6ab8fe3ce5f111`; the current Owner transfer opens this slice only. All P3-B1 evidence IDs remain pending, and no production/test owner is editable until fresh file-detail and exact file/symbol impact prove ownership.
  - Boundary Adjustment Checkpoint (2026-08-14): the first exact three-file candidate passed focused/parity, P3-A/P3-B preservation, TSJS/ScopeIR, and a holder-clean full build, but repo-native product regression exposed four deterministic stale-oracle failures: Vue parity gains parameter Definitions `repo`, `user`, and `value` plus `DEFINES 10 -> 13`; the Go resolution fixture gains six parameter Variable nodes and six `DEFINES` edges. Fresh file-detail, upstream file impact, exact test-function impact, and exact `graphCountBaseline` contract impact are LOW with `0` affected files/modules/processes/flows; all three golden artifacts are LOW/`0`. The historical TypeScript resolution-only baseline remains immutable at `24` nodes / `54` relationships / `18` `DEFINES`; the allowed graph-parity repair must record the current Go node count and classified `+6` parameter-node delta separately, update only the Go-side `DEFINES` delta, and preserve the generated TypeScript values/source. This checkpoint opens only the exact test/oracle edits named in the Scope Boundary. The failed regression remains non-PASS evidence, prior builds are superseded after any final-byte change, and P3-C remains locked.
  - Supervisor Repair Checkpoint (2026-08-14, satisfied by REVIEW2): `E3-P3B1-REVIEW1` is the immutable `REJECT` in `reports/Supervisor/rp_supervisor_260814_222415_by_gpt-5_child03_p3b1_parameter_contexts.md` (`19,388` bytes; SHA-256 `93A9427D84405ADD763140EEC7241FB1A795E796B41DDCCD4F786523F2E40184`). Its sole rejected invariant was `P3B1-PARAMETER-CONTEXT-DISAMBIGUATION-1`: TypeScript labeled tuple members and nested type-signature parameters share `required_parameter`/`optional_parameter` kinds and must not become runtime parameter facts merely because a callable is an ancestor. At that checkpoint, repair was restricted to the exact two files above, all six dependent files and every cleared invariant were locked, final-byte build/behavior gates required refresh, and detect/commit/push plus every later slice were forbidden before PASS.
  - Supervisor Acceptance Checkpoint (2026-08-14): focused `E3-P3B1-REVIEW2` returned `PASS` in `reports/Supervisor/rp_supervisor_260814_233151_by_gpt-5_child03_p3b1_parameter_contexts_review2.md` (`15,253` bytes; SHA-256 `0A0240FB56235EF9BE24F1C02E379DD83C4E7D2B8434F157B1203374B9774299`). Direct `formal_parameters` ownership, runtime-callable classification, and exact callable `parameters` field node-ID identity close the sole blocker; independent final-byte full build, former-failure `1/1`, parameter/parity `5/5`, dependent `4/4`, preservation `10/10`, four-package, and repo-native gates all passed. Final staged `E3-P3B1-DETECT1` then passed at `309` changed symbols / `17` changed files / `13` affected files, zero affected processes/flows, and graph health `0/0/0`. Orchestration owns only isolated `E3-P3B1-COMMIT1` and immediate push. P3-B2 and every later slice remain locked until commit and push succeed.
    - [x] Preserve the first independent REJECT as `E3-P3B1-REVIEW1`.
    - [x] Record focused independent PASS as `E3-P3B1-REVIEW2`.
    - [x] Refresh roadmap, checklist, evidence, benchmark, and actual status without changing locked source/test/report bytes.
    - [x] Run and record final orchestration-owned staged detect as `E3-P3B1-DETECT1` (`309` changed symbols / `17` changed files / `13` affected files; `0` affected processes/flows; current graph health `0/0/0`).
    - [x] Create and immediately push the isolated P3-B1 commit as `E3-P3B1-COMMIT1`; commit `01f160e6e28ad74c1f379ce5ea47e643a5a14652` was pushed to `origin/master`.
  - Commit and Successor Opening (2026-08-15): `E3-P3B1-COMMIT1` is `01f160e6e28ad74c1f379ce5ea47e643a5a14652`, parent `06229bea5735e75c8d8a476c738f627bf93def8d`; the exact `17`-path slice was pushed once to `origin/master`, and local/remote `master` match. P3-B2 is now the sole open coder slice. Its production/test owners remain unselected until fresh catch-context file-detail and exact file/symbol impact evidence are recorded; P3-B2A and every later slice remain locked.

- [x] P3-B2: Integrate catch contexts.
  - Goal: apply accepted leaf semantics to catch-clause declarations with correct catch scope.
  - Scope Boundary:
    - Editable: `internal/providers/tsjs/definitions.go` only for `collector.emitDefinitionKind` catch dispatch plus private catch-only leaf/Definition/OwnedDefID/local-Binding helpers; `internal/providers/tsjs/scopes.go` only for `collector.collectScopeCandidateForKind` to add `catch_clause -> ScopeBlock` over the catch-clause range; `internal/providers/tsjs/extract_test.go` only for focused real-`Extract` catch tests and the exact `caught` expectation transition inside `TestExtractParameterBindingPatternsPreserveShadowingAndSiblingContexts`.
    - Inspect-only: accepted walker/ScopeIR contract, collector traversal/result plumbing, pinned TS/JS catch grammar, references, type inference, imports, and existing scope construction outside the one authorized catch case.
    - Preserve-only: every existing module/class/function scope rule and scope sorting/parent/ID/containment helper; accepted variable/parameter behavior; loop behavior and the assignment-form non-declaration invariant; assignment write/reference behavior remains partial/missing and is deferred to P3-B2A/P3-C.
    - Out of scope: exception-flow modeling, general block-scope modeling, graph projection/resolution, persistence, target validation, and all later contexts.
  - Non-Goals: no changes to exception-flow modeling.
  - Pre-flight Questions:
    - Data source: current catch-clause AST nodes.
    - Display permission: N/A; no visible or permission-controlled surface changes.
    - DB read flow: N/A; the slice reads catch AST and in-memory ScopeIR.
    - DB write flow: N/A; only in-memory facts and package-local test output are written.
    - Render location: N/A; inspect catch ScopeIR and ledger evidence.
    - UI behavior flow: N/A; nearest boundary is catch ScopeIR.
    - Docker runtime: N/A; no browser/app runtime changes, but the full build is mandatory.
    - Playwright target: N/A; exercise the real catch-fact boundary.
    - Behavior test: identifier and supported catch patterns, exact catch scope, shadowing, and no duplicate declaration.
    - Cleanup/quarantine: keep fixtures under package `testdata/`; remove debug output.
    - External side effects: none.
    - N/A notes: UI, DB, Docker, and Playwright are not applicable to this provider slice.
  - Work Steps:
    1. Preserve the accepted `E3-P3B2-IMPACT1` three-file/three-existing-symbol boundary, add only the catch-clause `ScopeBlock` candidate in `scopes.go`, and implement catch dispatch/emission in `definitions.go` using the accepted walker exactly once; `catch {}` returns with zero facts/diagnostics, and no type-inference path is called.
       - UI flow check: N/A; no visible flow exists.
       - DB/data flow check: verify one binding fact per supported catch leaf in the correct scope.
       - Render location check: N/A; inspect catch ScopeIR.
       - Mini QA: after the full build in step 2, exercise the catch-fact boundary.
       - Evidence target: `E3-P3B2-IMPACT1`, `E3-P3B2-SRC1`.
    2. After production behavior is correct, add focused TypeScript/JavaScript identifier/pattern/optional-catch/scope/shadowing/zero-duplicate tests and update only the stale `caught` assertion named above; run the holder-clean full build, ScopeIR/preservation/product validation, refresh ledgers, obtain Supervisor PASS, detect changes, and commit.
       - UI flow check: N/A; no visible flow exists.
       - DB/data flow check: verify catch scope/shadowing and zero duplicate declarations.
       - Render location check: evidence ledger and focused test output.
       - Mini QA: exercise the built catch-fact boundary and record the result.
       - Evidence target: `E3-P3B2-BUILD1`, `E3-P3B2-TEST1`, `E3-P3B2-BOUNDARY1`, `E3-P3B2-REVIEW1`, `E3-P3B2-DETECT1`, `E3-P3B2-COMMIT1`.
  - Implementation Gate: P3-B1 is accepted, committed, and pushed at `01f160e6e28ad74c1f379ce5ea47e643a5a14652`; fresh `E3-P3B2-IMPACT1` and Orchestration's planner update authorize exactly the three files and three existing symbols above. No other production/test owner may open without a new stop-and-plan-update checkpoint.
  - Acceptance:
    - Source: supported catch leaves emit once in the exact catch `ScopeBlock`; optional `catch {}` emits zero leaves/diagnostics/definitions; no type or exception-flow behavior is invented.
    - Runtime/UI: N/A; ScopeIR catch facts are the real boundary.
    - Data: sibling contexts and assignment behavior remain unchanged.
    - Behavior test: TypeScript and JavaScript catch, identifier/pattern/default/rest, optional catch, scope parent/range, shadowing, zero-duplicate, and P3-A/P3-B/P3-B1 preservation matrices pass.
    - Cleanup: no unapproved artifact remains.
    - Evidence IDs: `E3-P3B2-IMPACT1`, `E3-P3B2-SRC1`, `E3-P3B2-BUILD1`, `E3-P3B2-TEST1`, `E3-P3B2-BOUNDARY1`, `E3-P3B2-REVIEW1`, `E3-P3B2-DETECT1`, `E3-P3B2-COMMIT1`.
    - Actual-status rows refreshed: catch extraction.
  - Evidence Targets: catch oracle, scope results, build, Supervisor, detect-changes, commit.
  - Actual-status Update: transition the catch row to `correct`.
  - Commit Boundary: one P3-B2 commit.
  - Opening Checkpoint (2026-08-15): predecessor `E3-P3B1-COMMIT1` is satisfied and pushed at `01f160e6e28ad74c1f379ce5ea47e643a5a14652`. This opens P3-B2 only. The coder must begin from fresh current graph/source/file-detail/impact evidence, select the smallest exact catch production/test boundary, and update this plan before code if evidence changes the generic owner wording. No P3-B2 evidence ID is accepted yet; P3-B2A/P3-C and later slices remain locked.
  - Impact Boundary Checkpoint (2026-08-15): the visible coder task stopped before code because source proves scope collection/build precedes definition emission and the prior inspect-only lexical-scope wording could not produce correct catch ownership. Durable `E3-P3B2-IMPACT1` in `reports/coder/rp_coder_260815_002037_by_gpt-5_child03_p3b2_catch_contexts.md` (`11,579` bytes; SHA-256 `77E2EC40404015338B75573EEB7D0303BC0EE90279B4AFBF77676E47161022D1`) selects the exact three-file boundary above. File risks are definitions HIGH/MEDIUM, scopes HIGH/MEDIUM, and test LOW; exact `emitDefinitionKind`, `collectScopeCandidateForKind`, and stale-test-function impacts are LOW `0/0/0`, with no affected process/flow. This planner update authorizes resume in the same visible coder task only; no implementation/build/test/review/detect/commit evidence is accepted yet.
  - Supervisor Acceptance Checkpoint (2026-08-15): independent `E3-P3B2-REVIEW1` returned `PASS` in `reports/Supervisor/rp_supervisor_260815_012555_by_gpt-5_child03_p3b2_catch_contexts.md` (`14,774` bytes; SHA-256 `E80DA8AEB480678708F4ED02FDBED455D21206D3F0C79221E296286A605116D4`). Source review cleared the exact catch dispatch/emission and one-symbol catch-scope changes. Independent holder-clean full build, focused catch `4/4`, P3-A/P3-B/P3-B1 preservation `14/14`, TSJS scope/parity `6/6`, ScopeIR tree/normalization, four-package, and repo-native command/internal gates all passed. Optional catch emits zero binding facts, every accepted catch leaf has one exact catch-scope Definition/OwnedDefID/BindingLocal, shadowing stays distinct, assignment `written` remains zero declarations, and no type/fallback/exception/loop/graph behavior was added. Acceptance is P3-B2 only for planner refresh, final detect, isolated commit, and immediate push; P3-B2A and later slices remain locked.
    - [x] Record independent PASS as `E3-P3B2-REVIEW1`.
    - [x] Refresh roadmap, checklist, evidence, benchmark, and actual status without changing accepted source/test/report bytes.
    - [x] Run and record final orchestration-owned staged detect as `E3-P3B2-DETECT1` (`242` changed symbols / `11` changed files / `11` affected files; `0` affected processes/flows; current graph health `0/0/0`).
    - [x] Create and immediately push the isolated P3-B2 commit as `E3-P3B2-COMMIT1`; commit `19247b4eb58a4e01a6256f3d63bbb59839644d64` was pushed once to `origin/master`.
  - Commit and Successor Opening (2026-08-15): `E3-P3B2-COMMIT1` is `19247b4eb58a4e01a6256f3d63bbb59839644d64`, parent `01f160e6e28ad74c1f379ce5ea47e643a5a14652`; the exact `11`-path P3-B2 slice was pushed once to `origin/master`, and local/remote `master` match. P3-B2A is now the sole open slice. Its loop production/test owners remain unselected until fresh source, file-detail, and exact file/symbol impact evidence are recorded; P3-C and every later slice remain locked.

- [x] P3-B2A: Integrate `for-of`/`for-in` declaration contexts.
  - Goal: apply accepted leaf semantics to loop declarations while keeping assignment-form loops as writes.
  - Scope Boundary:
    - Editable: `internal/providers/tsjs/definitions.go` only at `collector.emitDefinitionKind` plus private loop-only declaration/context/ownership/emission helpers; `internal/providers/tsjs/scopes.go` only at `collector.collectScopeCandidateForKind` plus an optional private predicate for lexical `let|const` loop scope; `internal/providers/tsjs/references.go` only at `collector.emitReferenceKind` plus private assignment-target `AccessWrite` helpers and suppression of the current loop-member false read; `internal/providers/tsjs/extract_test.go` only for focused real-parser/real-`Extract` loop tests and the exact stale loop assertion in `TestExtractParameterBindingPatternsPreserveShadowingAndSiblingContexts`.
    - Inspect-only: accepted walker/ScopeIR contracts, collector traversal/result plumbing, AST helpers, pinned TypeScript/JavaScript grammar, type inference, imports, and existing scope/reference behavior outside the exact loop cases above.
    - Preserve-only: identifier and destructured variables, parameters including type-only-label exclusion, catch bindings and catch scope, import bindings, module/class/function/catch scope rules, calls and member accesses outside loop assignment targets, accepted golden artifacts, Vue/resolution behavior, and every prior accepted P3 invariant.
    - Out of scope: graph projection/resolution/persistence, new ScopeIR contracts, `await`/`using` semantics, general block-scope repair, control-flow/process modeling, target validation, and every later slice.
  - Non-Goals: no general reference-system rewrite, loop-flow/process-model change, or unsupported `await`/`using` expansion.
  - Pre-flight Questions:
    - Data source: current `for-of`/`for-in` AST nodes.
    - Display permission: N/A; no visible or permission-controlled surface changes.
    - DB read flow: N/A; the slice reads loop AST and in-memory ScopeIR/reference state.
    - DB write flow: N/A; only in-memory facts and package-local test output are written.
    - Render location: N/A; inspect loop ScopeIR and ledger evidence.
    - UI behavior flow: N/A; nearest boundary is loop ScopeIR.
    - Docker runtime: N/A; no browser/app runtime changes, but the full build is mandatory.
    - Playwright target: N/A; exercise the real loop-fact boundary.
    - Behavior test: declaration patterns, assignment forms, scopes, holes/index paths, and sibling preservation.
    - Cleanup/quarantine: keep fixtures under package `testdata/`; remove debug output.
    - External side effects: none.
    - N/A notes: UI, DB, Docker, and Playwright are not applicable to this provider slice.
  - Work Steps:
    1. Preserve accepted `E3-P3B2A-IMPACT1`, then implement production behavior inside the exact three production owners: discriminate declaration form by non-empty `for_in_statement.kind`; create an exact loop `ScopeBlock` only for lexical `let|const`; emit identifier or accepted pattern leaves once with `var` owned by the nearest function/module and `let|const` by the loop scope; route no-`kind` assignment targets to truthful `AccessWrite` facts while creating zero declarations and eliminating the loop-member false read.
       - UI flow check: N/A; no visible flow exists.
       - DB/data flow check: verify one fact per declaration leaf, exact lexical ownership, zero declaration for assignment forms, and one truthful write per representable assignment target.
       - Render location check: N/A; inspect loop ScopeIR.
       - Mini QA: after the full build in step 2, exercise the loop-fact boundary.
       - Evidence target: `E3-P3B2A-IMPACT1`, `E3-P3B2A-SRC1`.
    2. After production behavior is correct, add focused TypeScript/JavaScript declaration, assignment-write, scope/shadowing, path/modifier, and parity tests plus only the authorized stale assertion transition; terminate build holders, prove the build lock free, run the full build on final bytes, then validate the real ScopeIR boundary, P3-A/P3-B/P3-B1/P3-B2 preservation, TSJS/ScopeIR, four-package, and repo-native command/internal gates. Remove debug output, create the durable coder handoff, and stop for independent Supervisor review; detect, stage, commit, and push remain Orchestration-owned after PASS.
       - UI flow check: N/A; no visible flow exists.
       - DB/data flow check: verify loop paths/scopes, assignment declaration delta `0`, exact write kinds with no duplicate/false read, and sibling preservation.
       - Render location check: evidence ledger and focused test output.
       - Mini QA: exercise the built loop-fact boundary and record the result.
       - Evidence target: `E3-P3B2A-BUILD1`, `E3-P3B2A-TEST1`, `E3-P3B2A-BOUNDARY1`, `E3-P3B2A-REVIEW1`, `E3-P3B2A-DETECT1`, `E3-P3B2A-COMMIT1`.
    3. After `E3-P3B2A-REVIEW1` rejected only `P3B2A-ASSIGNMENT-TARGET-WRAPPER-CLOSURE-1`, preserve every locked declaration/scope/direct-target byte and repair only legal representable target wrappers in `references.go`: recurse through TypeScript/JavaScript `parenthesized_expression` and TypeScript `non_null_expression` in both assignment-write extraction and exact-member containment/suppression. After production is correct, add the focused real-parser/real-`Extract` wrapper oracle in `extract_test.go`, including the explicit bracket-target out-of-contract control. Rerun one holder-clean full build and the complete original validation matrix on final repair bytes, create a new coder resubmission report, and request `E3-P3B2A-REVIEW2`; do not run detect or open P3-C before PASS.
       - Evidence target: repaired `E3-P3B2A-SRC1`, refreshed `E3-P3B2A-BUILD1`, `E3-P3B2A-TEST1`, `E3-P3B2A-BOUNDARY1`, and pending `E3-P3B2A-REVIEW2`.
  - Implementation Gate: P3-B2 is accepted, committed, and pushed at `19247b4eb58a4e01a6256f3d63bbb59839644d64`; accepted `E3-P3B2A-IMPACT1` and this Orchestration planner update authorize exactly the four files and four existing symbols above. No other production/test/fixture/contract owner may open without a new stop-and-plan-update checkpoint.
  - Acceptance:
    - Source: declaration-form loop leaves emit once with `var -> nearest function/module` and `let|const -> exact loop ScopeBlock`; assignment-form loops create `0` declarations and emit truthful `AccessWrite` facts without the current member false read.
    - Runtime/UI: N/A; ScopeIR loop facts are the real boundary.
    - Data: all prior context outputs remain correct; no new ScopeIR, graph, resolution, type, import, flow, `await`, or `using` semantics are introduced.
    - Behavior test: TypeScript/JavaScript declaration/assignment, scope/shadowing, path/modifier, write-kind, zero-duplicate, and prior-slice preservation matrices pass.
    - Cleanup: no unapproved artifact remains.
    - Evidence IDs: `E3-P3B2A-IMPACT1`, `E3-P3B2A-SRC1`, `E3-P3B2A-BUILD1`, `E3-P3B2A-TEST1`, `E3-P3B2A-BOUNDARY1`, `E3-P3B2A-REVIEW1`, `E3-P3B2A-REVIEW2`, `E3-P3B2A-DETECT1`, `E3-P3B2A-COMMIT1`.
    - Actual-status rows refreshed: loop declaration extraction.
  - Evidence Targets: loop oracle, declaration/assignment delta, build, Supervisor, detect-changes, commit.
  - Actual-status Update: transition the loop row to `correct`.
  - Commit Boundary: one P3-B2A commit.
  - Opening Checkpoint (2026-08-15): predecessor `E3-P3B2-COMMIT1` is satisfied and pushed at `19247b4eb58a4e01a6256f3d63bbb59839644d64`. This opens P3-B2A only. The visible coder task must begin read-only from fresh current graph/source/file-detail/impact evidence, distinguish declaration-form `for-in`/`for-of` from assignment forms, and stop for Orchestration planner authorization before editing any exact production/test owner. No P3-B2A evidence ID is accepted yet; P3-C and every later slice remain locked.
  - Impact Boundary Checkpoint (2026-08-15): the visible coder task stopped before code and produced durable `E3-P3B2A-IMPACT1` in `reports/coder/rp_coder_260815_021253_by_gpt-5_child03_p3b2a_loop_contexts.md` (`18,519` bytes; SHA-256 `5D29AFBABB96A0549705A445D1E2262154F37A442947B01F102A5E75DC673612`). Pinned TypeScript/JavaScript grammar and real parser output prove both operators use `for_in_statement`, declaration form has a non-empty `kind`, assignment form has none, and `left` is direct rather than a `variable_declarator`. Current source proves declaration emission, lexical loop scope, and assignment-target write truth require the exact `definitions.go`, `scopes.go`, `references.go`, and `extract_test.go` boundary above. File risks are HIGH/HIGH/HIGH/LOW with upstream file impact MEDIUM `9/2/0`, MEDIUM `9/2/0`, LOW `1/1/0`, and LOW `0/0/0`; all four exact existing-symbol impacts are LOW `0/0/0`. This planner update authorizes implementation resume in the same visible coder task only. `E3-P3B2A-SRC1` and every later gate remain pending; P3-C and later slices stay locked.
  - Supervisor Repair Checkpoint (2026-08-15): independent `E3-P3B2A-REVIEW1` returned `REJECT` in `reports/Supervisor/rp_supervisor_260815_082939_by_gpt-5_child03_p3b2a_loop_contexts.md` (`15,125` bytes; SHA-256 `6D9CD9B381D851762A65A42B0DE2170DB5F242A8F1555F54E91C967DF494A628`). Its sole rejected invariant is `P3B2A-ASSIGNMENT-TARGET-WRAPPER-CLOSURE-1`: legal TS/JS parenthesized targets and TypeScript non-null targets wrap already-representable identifier/member leaves but currently lose their write or retain a false member read. Bracket/subscript targets are explicitly not the rejection basis because the current dot-member `AccessFact` contract cannot represent computed access truthfully without out-of-scope contract expansion. Repair is restricted to `references.go` plus `extract_test.go`; `definitions.go`, `scopes.go`, direct target recursion, declaration/ownership/scope behavior, governance, and both prior coder reports remain byte-locked. Final-byte build/test/boundary gates require one refresh after repair; `E3-P3B2A-REVIEW2`, detect, commit, push, and P3-C remain pending and locked.
  - Supervisor Acceptance Checkpoint (2026-08-15): focused `E3-P3B2A-REVIEW2` returned `PASS` in `reports/Supervisor/rp_supervisor_260815_091529_by_gpt-5_child03_p3b2a_loop_contexts_review2.md` (`13,197` bytes; SHA-256 `80BC32BD13A34E23CA1F518FC20B1C5E04CB1882516D82763728E1DF2E552C31`). The exact two-file repair unwraps only parenthesized and TypeScript non-null assignment targets in write extraction and exact-member containment. Independent final-byte full build, loop `5/5` with wrapper controls `9/9`, prior-context preservation `14/14`, catch preservation `4/4`, TSJS parity `6/6`, ScopeIR `3/3`, and four-package gates passed. The repo-native command exited `1` only in two CLI skill-generation tests caused by the explicitly excluded Owner-owned `internal/aicontext/skills/ui-ux-pro-max-skill/` tree; the report records that failure truthfully and does not use it as P3-B2A evidence. Acceptance is P3-B2A only for planner refresh, final detect, isolated commit, and immediate push; P3-C and every later slice remain locked until push succeeds.
    - [x] Preserve the first independent REJECT as `E3-P3B2A-REVIEW1`.
    - [x] Record focused independent PASS as `E3-P3B2A-REVIEW2`.
    - [x] Refresh roadmap, checklist, evidence, benchmark, and actual status without changing accepted source/test/report bytes.
    - [x] Run and record final orchestration-owned staged detect as `E3-P3B2A-DETECT1` using `anvien detect-changes --repo E:\Anvien --scope staged` (`348` changed symbols / `14` changed files / `14` affected files; `0` affected processes/flows; current health `0/0/0`); `--scope all` remains external-state disclosure only.
    - [x] Create and immediately push the isolated P3-B2A commit as `E3-P3B2A-COMMIT1` (the current 14-path orchestration closure action; Git/handoff records the resulting hash).

- [ ] P3-C: Project and resolve binding occurrences in the current graph.
  - Goal: carry accepted binding facts into the graph so each source leaf remains distinct and downstream references select the correct lexical symbol.
  - Scope Boundary:
    - Editable: only graph, persistence, and reference-resolution owners proven affected by fresh impact evidence.
    - Inspect-only: completed extraction contexts and Child 02 persistence contract.
    - Preserve-only: export/module/ambient semantics and unaffected readers.
    - Out of scope: unrelated transport/cache changes and target validation.
  - Non-Goals: do not infer missing bindings by name rescue.
  - Pre-flight Questions:
    - Data source: accepted P3-A through P3-B2A facts.
    - Display permission: preserve existing access; no new permission behavior is planned.
    - DB read flow: read accepted binding facts and the current graph/persistence contracts proven affected.
    - DB write flow: write the current in-memory graph and isolated test stores only where impact proves persistence is affected.
    - Render location: graph/persisted-graph command output and evidence ledger; no UI is planned.
    - UI behavior flow: N/A unless impact proves an existing public surface consumes changed fields; update the plan before expansion.
    - Docker runtime: N/A for the planned non-UI boundary; if P0 proves a served UI is affected, update the plan and require the built Docker runtime.
    - Playwright target: N/A for the planned graph/CLI boundary; required only after an evidence-backed UI scope update.
    - Behavior test: declaration/binding conservation, lexical shadowing, exact reference endpoints, import count, assignment non-declaration, and persistence parity where applicable.
    - Cleanup/quarantine: use isolated repo-local stores and remove debug output; no target run.
    - External side effects: none.
    - N/A notes: no UI/Playwright work is authorized by the current affected-surface evidence.
  - Work Steps:
    1. Refresh graph evidence, map the exact extraction-to-graph-to-reference path, update the plan if scope changes, and implement production projection/resolution in the smallest proven owners.
       - UI flow check: N/A under the current non-UI scope.
       - DB/data flow check: verify every accepted fact reaches one occurrence and the correct lexical target.
       - Render location check: graph/persisted-graph command output.
       - Mini QA: after the full build in step 2, exercise the nearest real graph boundary.
       - Evidence target: `E3-P3C-IMPACT1`, `E3-P3C-SRC1`.
    2. Add focused graph/reference tests after code, run the full build, validate affected persistence and command output, refresh ledgers, obtain Supervisor PASS, detect changes, and commit.
       - UI flow check: N/A unless an approved scope update added a public surface.
       - DB/data flow check: verify occurrence conservation, zero orphan/drift, correct shadowing, and zero assignment/import deltas.
       - Render location check: evidence ledger and real graph/persisted-graph output.
       - Mini QA: exercise the built graph/loader/CLI boundary and record command plus observed result.
       - Evidence target: `E3-P3C-BUILD1`, `E3-P3C-TEST1`, `E3-P3C-BOUNDARY1`, `E3-P3C-REVIEW1`, `E3-P3C-DETECT1`, `E3-P3C-COMMIT1`.
  - Implementation Gate: P3-A through P3-B2A accepted and committed; every editable owner has current file-detail/impact evidence; no unplanned reader is modified.
  - Acceptance:
    - Source: every accepted binding leaf reaches one graph declaration/binding occurrence and the correct lexical symbol.
    - Runtime/UI: nearest real graph command or persisted-graph boundary returns the corrected facts; no unrelated public surface is changed.
    - Data: occurrence conservation holds; assignment declarations and import-count deltas are `0`.
    - Behavior test: graph endpoints, shadowing, and binding-caused gap regression pass.
    - Cleanup: isolated stores/debug output are removed or retained only as valid evidence.
    - Evidence IDs: `E3-P3C-IMPACT1`, `E3-P3C-SRC1`, `E3-P3C-BUILD1`, `E3-P3C-TEST1`, `E3-P3C-BOUNDARY1`, `E3-P3C-REVIEW1`, `E3-P3C-DETECT1`, `E3-P3C-COMMIT1`.
    - Actual-status rows refreshed: graph projection, lexical resolution, and affected persistence rows.
  - Evidence Targets: exact blast radius, occurrence/reference oracle, affected persistence parity, build, boundary output, Supervisor, detect-changes, commit.
  - Actual-status Update: transition only proven affected graph/resolution rows to `correct`.
  - Commit Boundary: one P3-C commit.

- [ ] P3-C2: Validate the complete binding contract against the real target.
  - Goal: prove the bounded defect is fixed on the real target without modifying target source or broadening into scanner work.
  - Scope Boundary:
    - Editable: this child’s Anvien-side ledgers and reusable validation assets only if an existing owner is proven necessary.
    - Inspect-only: target source and current target graph output.
    - Preserve-only: target source/worktree.
    - Out of scope: production repair, target fixtures, and all other child defects.
  - Non-Goals: do not copy target source or claim global TypeScript accuracy.
  - Pre-flight Questions:
    - Data source: independent six-site source oracle plus the target's normal analyze output.
    - Display permission: preserve existing command access; no permission behavior changes.
    - DB read flow: read target source, target-local graph output, and only affected persisted records.
    - DB write flow: normal target-local `.anvien` operational output only.
    - Render location: Anvien-side evidence ledger and human-readable command output; target-side reports are forbidden.
    - UI behavior flow: N/A unless an already approved public surface was affected in P3-C.
    - Docker runtime: N/A for the planned non-UI target boundary; use built Docker only if an approved UI scope exists.
    - Playwright target: N/A for source/graph validation; required only for an approved UI scope.
    - Behavior test: named `6/6`, exact paths/ranges/scopes, six downstream references, zero binding-caused gaps, and target boundary.
    - Cleanup/quarantine: all reports remain in Anvien; remove debug artifacts and never write target fixtures/probes.
    - External side effects: normal analyzer-owned target-local graph output only.
    - N/A notes: validation is non-UI unless P3-C recorded an approved public-surface impact.
  - Work Steps:
    1. Capture target HEAD/status and pre-run boundary, then run the normal built analyzer and compare source oracle, extracted facts, graph occurrences, and reference outcomes.
       - UI flow check: N/A unless an approved public surface exists.
       - DB/data flow check: read only the target's current analyzer output and compare all six exact sites.
       - Render location check: Anvien-side command captures and ledger evidence.
       - Mini QA: exercise the real built target analyze/query boundary and inspect human-readable output.
       - Evidence target: `E3-P3C2-ORACLE1`, `E3-P3C2-TARGET1`.
    2. Record gap/shadowing/assignment/import and target-contamination results, refresh ledgers, obtain Supervisor PASS, run the boundary/change check, and commit valid Anvien-side evidence only.
       - UI flow check: N/A unless an approved public surface exists.
       - DB/data flow check: verify target source preservation, `6/6`, and zero binding-caused gaps.
       - Render location check: final evidence ledger and boundary manifest under Anvien.
       - Mini QA: visually inspect the retained result and confirm no target-side report/fixture exists.
       - Evidence target: `E3-P3C2-BOUNDARY1`, `E3-P3C2-REVIEW1`, `E3-P3C2-DETECT1`, `E3-P3C2-COMMIT1`.
  - Implementation Gate: P3-C accepted and committed; the oracle names all six sites without copied target source; target pre-state is recorded.
  - Acceptance:
    - Source: all six target bindings have correct declaration/binding/path/range/scope records.
    - Runtime/UI: the real target graph boundary shows `6/6` corrected sites and `0` binding-caused gaps.
    - Data: target source/worktree boundary is preserved.
    - Behavior test: assignment/import/shadowing controls remain correct.
    - Cleanup: no target report, fixture, probe, or debug artifact remains.
    - Evidence IDs: `E3-P3C2-ORACLE1`, `E3-P3C2-TARGET1`, `E3-P3C2-BOUNDARY1`, `E3-P3C2-REVIEW1`, `E3-P3C2-DETECT1`, `E3-P3C2-COMMIT1`.
    - Actual-status rows refreshed: all Child 03 rows and Child 04 predecessor state.
  - Evidence Targets: six-site oracle, target graph/reference results, boundary manifest, Supervisor, change check, evidence commit.
  - Actual-status Update: transition the bounded target row to `correct` and append the successor handoff refresh.
  - Commit Boundary: commit only Anvien-side validation/ledger evidence after acceptance.

- [ ] Pn-A: Run the Supervisor acceptance loop.
  - Goal: independently verify all seven slices, source diff, build, real-boundary results, benchmark, target boundary, and evidence integrity.
  - Work Steps: invoke the Supervisor skill; return only rejected invariants to the owning slice; repeat until PASS or a documented blocker.
  - Implementation Gate: every P3 slice is complete or explicitly blocked.
  - Acceptance: `E3-PNA-REVIEW1` records Supervisor PASS or a precise blocker; no self-acceptance is used.
- [ ] Pn-B: Remove dead work created by this child.
  - Goal: leave no superseded fixture, debug output, failed evidence, duplicate report, or rejected approach.
  - Work Steps: inventory child-created artifacts, remove only dead child work, and obtain Supervisor review of cleanup.
  - Implementation Gate: Pn-A has completed its first review.
  - Acceptance: `E3-PNB-CLEAN1` and `E3-PNB-REVIEW1` prove the retained artifact set is current.
- [ ] Pn-C: Close Child 03 and hand off Child 04.
  - Goal: record final validation, commit state, and the exact successor opening condition.
  - Work Steps:
    1. Confirm all slice commits, final build/boundary evidence, benchmark rows, and target boundary.
    2. Run final detect-changes when implementation work is present.
    3. Refresh Child 04 actual status from accepted Child 03 evidence and record the handoff.
  - Implementation Gate: Pn-A and Pn-B PASS.
  - Acceptance: `E3-PNC-DETECT1`, `E3-PNC-COMMIT1`, and `E3-PNC-HANDOFF1` are recorded; worktree state is known.

## Risk Notes

- The accepted defect is bounded to six sites; treating it as proof of every TypeScript pattern form would overstate evidence. General fixtures are required before broad claims.
- Type inference is a likely accidental gate, but P0 must prove the current call path before code changes.
- Property keys, defaults, array holes, assignment patterns, and import bindings can create false declarations or double counts if the walker loses syntax role.
- Lexical-scope mistakes can make counts look correct while references bind to the wrong same-name symbol; acceptance therefore checks exact symbol targets, not totals alone.
- Graph and persistence owners can have HIGH/CRITICAL blast radius. That warning requires narrow scope and regression evidence; it does not prohibit the necessary edit.
- Any newly discovered affected reader or contract changes the slice boundary and requires a plan update before code; it does not expand unrelated consumers automatically.
