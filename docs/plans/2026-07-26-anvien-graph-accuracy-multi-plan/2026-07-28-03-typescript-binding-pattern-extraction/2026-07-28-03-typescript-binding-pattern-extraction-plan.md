# Anvien TypeScript Binding-Pattern Extraction Plan

## Metadata

- Date: `2026-07-28`
- Status: `draft / P0 incomplete`
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

- [ ] P0-A: Complete current actual status before implementation.
  - Goal: replace historical assumptions with current source, graph, and impact evidence for the binding pipeline.
  - Work Steps:
    1. Refresh the Anvien graph and record current HEAD/source basis.
    2. Read the current TS binding extraction, fact/index, graph projection, and relevant tests in full; inventory every binding context and silent-exit path.
    3. Run `file-detail` and upstream `impact` for every candidate owner and record exact related-file counts and blast radius.
    4. Classify correct, partial, wrong, missing, and blocked rows in actual status; update P3-A work steps if exact ownership differs from this plan.
  - Implementation Gate: no production edit until the actual-status Final P0 Decision permits P3-A.
  - Acceptance: `E0-P0A-RULE1`, `E0-P0A-SKILL1`, `E0-P0A-ORIGIN1`, `E0-P0A-VERIFY1`, `E0-P0A-VERIFY2`, `E0-P0A-SCOPE1`, `E0-P0A-GRAPH1`, `E0-P0A-SRC1`, `E0-P0A-FD1`, `E0-P0A-IMPACT1`, and `E0-P0A-STATUS1` are recorded with no unresolved owner boundary.

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

- [ ] P3-A: Implement the binding-pattern contract and recursive leaf walker.
  - Goal: enumerate legal binding leaves, paths, ranges, modifiers, and unsupported syntax without graph emission.
  - Scope Boundary:
    - Editable: exact fact and TS walker owners selected by P0.
    - Inspect-only: existing collector, type inference, graph projection, and import extraction.
    - Preserve-only: identifier-only declaration behavior and all export behavior.
    - Out of scope: declaration-context wiring and graph emission.
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
    1. Refresh graph evidence, run exact file-detail/impact, confirm one-file responsibility, and implement the production fact/walker plus structured unsupported diagnostic.
       - UI flow check: N/A; no visible flow exists.
       - DB/data flow check: verify deterministic one-leaf-to-one-fact production in memory.
       - Render location check: N/A; inspect ScopeIR output.
       - Mini QA: after the full build in step 2, exercise the real ScopeIR extraction boundary; browser controls are N/A.
       - Evidence target: `E3-P3A-IMPACT1`, `E3-P3A-SRC1`.
    2. After production behavior is correct, add focused fixtures/tests, run the full build, validate ScopeIR output and diagnostic counts, refresh ledgers, obtain Supervisor PASS, run detect-changes, and commit.
       - UI flow check: N/A; no visible flow exists.
       - DB/data flow check: verify exact paths/ranges/modifiers and no duplicate facts.
       - Render location check: evidence ledger and focused test output.
       - Mini QA: exercise the built non-UI provider boundary and record the command and observed result.
       - Evidence target: `E3-P3A-BUILD1`, `E3-P3A-TEST1`, `E3-P3A-BOUNDARY1`, `E3-P3A-REVIEW1`, `E3-P3A-DETECT1`, `E3-P3A-COMMIT1`.
  - Implementation Gate: P0 is complete; accepted identity/range fields needed by binding facts are available; exact owners have fresh impact evidence.
  - Acceptance:
    - Source: every supported leaf has exactly one fact, correct path/range/modifiers, and unsupported syntax is diagnosed.
    - Runtime/UI: N/A; nearest real boundary is ScopeIR extraction after the full build.
    - Data: no graph or persistent-store mutation in this slice.
    - Behavior test: general pattern matrix passes deterministically.
    - Cleanup: no debug or target artifact remains.
    - Evidence IDs: `E3-P3A-IMPACT1`, `E3-P3A-SRC1`, `E3-P3A-BUILD1`, `E3-P3A-TEST1`, `E3-P3A-BOUNDARY1`, `E3-P3A-REVIEW1`, `E3-P3A-DETECT1`, `E3-P3A-COMMIT1`.
    - Actual-status rows refreshed: binding contract/walker.
  - Evidence Targets: source diff, pattern matrix, diagnostic count, build, boundary output, Supervisor, detect-changes, commit.
  - Actual-status Update: transition the binding walker row from its P0 classification to `correct`; leave declaration contexts unchanged.
  - Commit Boundary: one P3-A commit after all acceptance evidence passes.

- [ ] P3-B: Integrate variable-declaration contexts.
  - Goal: emit binding leaves for variable declarations even when per-leaf type inference is unavailable.
  - Scope Boundary:
    - Editable: variable-declaration adapter selected by impact evidence.
    - Inspect-only: walker, type inference, assignment/reference extraction, imports.
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
    1. Reconfirm exact impact/ownership and implement production variable-context wiring independently of type inference.
       - UI flow check: N/A; no visible flow exists.
       - DB/data flow check: verify each legal variable leaf reaches one in-memory declaration/binding fact.
       - Render location check: N/A; inspect ScopeIR output.
       - Mini QA: after the full build in step 2, exercise the variable ScopeIR boundary; browser controls are N/A.
       - Evidence target: `E3-P3B-IMPACT1`, `E3-P3B-SRC1`.
    2. Add focused tests after production code, run the full build, validate identifiers/type-failure/assignment/import controls, refresh ledgers, obtain Supervisor PASS, detect changes, and commit.
       - UI flow check: N/A; no visible flow exists.
       - DB/data flow check: assert correct ranges/scopes and zero false declaration/import deltas.
       - Render location check: evidence ledger and focused test output.
       - Mini QA: exercise the built variable-fact boundary and record the observed result.
       - Evidence target: `E3-P3B-BUILD1`, `E3-P3B-TEST1`, `E3-P3B-BOUNDARY1`, `E3-P3B-REVIEW1`, `E3-P3B-DETECT1`, `E3-P3B-COMMIT1`.
  - Implementation Gate: P3-A accepted and committed.
  - Acceptance:
    - Source: variable-pattern leaves emit once with correct scope/ranges and survive type-inference failure.
    - Runtime/UI: N/A; ScopeIR output is the real boundary.
    - Data: assignment destructuring adds `0` declarations and import-binding delta is `0`.
    - Behavior test: variable and identifier regression matrices pass.
    - Cleanup: fixture/debug state is confined to approved repo paths.
    - Evidence IDs: `E3-P3B-IMPACT1`, `E3-P3B-SRC1`, `E3-P3B-BUILD1`, `E3-P3B-TEST1`, `E3-P3B-BOUNDARY1`, `E3-P3B-REVIEW1`, `E3-P3B-DETECT1`, `E3-P3B-COMMIT1`.
    - Actual-status rows refreshed: variable declaration extraction.
  - Evidence Targets: variable fact oracle, type-inference-failure row, assignment/import deltas, build, Supervisor, detect-changes, commit.
  - Actual-status Update: transition the variable-declaration row to `correct`; leave other contexts pending.
  - Commit Boundary: one P3-B commit.

- [ ] P3-B1: Integrate parameter contexts.
  - Goal: apply accepted leaf semantics to function, method, constructor, and arrow-function parameters.
  - Scope Boundary:
    - Editable: exact parameter-context owner and its focused test owner.
    - Inspect-only: walker and scope construction.
    - Preserve-only: variables, catch, loops, calls, and type inference.
    - Out of scope: graph projection.
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
    2. Add focused tests after code, run the full build and ScopeIR validation, refresh ledgers, obtain Supervisor PASS, detect changes, and commit.
       - UI flow check: N/A; no visible flow exists.
       - DB/data flow check: verify parameter paths/scopes and sibling preservation.
       - Render location check: evidence ledger and focused test output.
       - Mini QA: exercise the built parameter-fact boundary and record the result.
       - Evidence target: `E3-P3B1-BUILD1`, `E3-P3B1-TEST1`, `E3-P3B1-BOUNDARY1`, `E3-P3B1-REVIEW1`, `E3-P3B1-DETECT1`, `E3-P3B1-COMMIT1`.
  - Implementation Gate: P3-B accepted and committed.
  - Acceptance:
    - Source: every legal parameter leaf emits once in the correct lexical scope.
    - Runtime/UI: N/A; ScopeIR parameter facts are validated after the full build.
    - Data: variable/catch/loop outputs are unchanged.
    - Behavior test: parameter and shadowing matrices pass.
    - Cleanup: no unapproved artifact remains.
    - Evidence IDs: `E3-P3B1-IMPACT1`, `E3-P3B1-SRC1`, `E3-P3B1-BUILD1`, `E3-P3B1-TEST1`, `E3-P3B1-BOUNDARY1`, `E3-P3B1-REVIEW1`, `E3-P3B1-DETECT1`, `E3-P3B1-COMMIT1`.
    - Actual-status rows refreshed: parameter extraction.
  - Evidence Targets: parameter oracle, scope/shadowing results, build, Supervisor, detect-changes, commit.
  - Actual-status Update: transition the parameter row to `correct`.
  - Commit Boundary: one P3-B1 commit.

- [ ] P3-B2: Integrate catch contexts.
  - Goal: apply accepted leaf semantics to catch-clause declarations with correct catch scope.
  - Scope Boundary:
    - Editable: exact catch-context owner and focused test owner.
    - Inspect-only: walker and lexical scope construction.
    - Preserve-only: variables, parameters, loops, and assignment writes.
    - Out of scope: graph projection.
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
    1. Record fresh catch-owner impact and implement production catch wiring.
       - UI flow check: N/A; no visible flow exists.
       - DB/data flow check: verify one binding fact per supported catch leaf in the correct scope.
       - Render location check: N/A; inspect catch ScopeIR.
       - Mini QA: after the full build in step 2, exercise the catch-fact boundary.
       - Evidence target: `E3-P3B2-IMPACT1`, `E3-P3B2-SRC1`.
    2. Add focused tests after code, run the full build and ScopeIR validation, refresh ledgers, obtain Supervisor PASS, detect changes, and commit.
       - UI flow check: N/A; no visible flow exists.
       - DB/data flow check: verify catch scope/shadowing and zero duplicate declarations.
       - Render location check: evidence ledger and focused test output.
       - Mini QA: exercise the built catch-fact boundary and record the result.
       - Evidence target: `E3-P3B2-BUILD1`, `E3-P3B2-TEST1`, `E3-P3B2-BOUNDARY1`, `E3-P3B2-REVIEW1`, `E3-P3B2-DETECT1`, `E3-P3B2-COMMIT1`.
  - Implementation Gate: P3-B1 accepted and committed.
  - Acceptance:
    - Source: supported catch leaves emit once in the correct catch scope.
    - Runtime/UI: N/A; ScopeIR catch facts are the real boundary.
    - Data: sibling contexts and assignment behavior remain unchanged.
    - Behavior test: catch and shadowing matrices pass.
    - Cleanup: no unapproved artifact remains.
    - Evidence IDs: `E3-P3B2-IMPACT1`, `E3-P3B2-SRC1`, `E3-P3B2-BUILD1`, `E3-P3B2-TEST1`, `E3-P3B2-BOUNDARY1`, `E3-P3B2-REVIEW1`, `E3-P3B2-DETECT1`, `E3-P3B2-COMMIT1`.
    - Actual-status rows refreshed: catch extraction.
  - Evidence Targets: catch oracle, scope results, build, Supervisor, detect-changes, commit.
  - Actual-status Update: transition the catch row to `correct`.
  - Commit Boundary: one P3-B2 commit.

- [ ] P3-B2A: Integrate `for-of`/`for-in` declaration contexts.
  - Goal: apply accepted leaf semantics to loop declarations while keeping assignment-form loops as writes.
  - Scope Boundary:
    - Editable: exact loop-declaration owner and focused test owner.
    - Inspect-only: walker, scope construction, and assignment/reference extraction.
    - Preserve-only: variables, parameters, catch, and import bindings.
    - Out of scope: graph projection.
  - Non-Goals: no loop-flow or process-model change.
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
    1. Record fresh loop-owner impact and implement production declaration-context wiring.
       - UI flow check: N/A; no visible flow exists.
       - DB/data flow check: verify one fact per declaration leaf and no declaration for assignment forms.
       - Render location check: N/A; inspect loop ScopeIR.
       - Mini QA: after the full build in step 2, exercise the loop-fact boundary.
       - Evidence target: `E3-P3B2A-IMPACT1`, `E3-P3B2A-SRC1`.
    2. Add focused tests after code, run the full build and ScopeIR validation, refresh ledgers, obtain Supervisor PASS, detect changes, and commit.
       - UI flow check: N/A; no visible flow exists.
       - DB/data flow check: verify loop paths/scopes, assignment delta, and sibling preservation.
       - Render location check: evidence ledger and focused test output.
       - Mini QA: exercise the built loop-fact boundary and record the result.
       - Evidence target: `E3-P3B2A-BUILD1`, `E3-P3B2A-TEST1`, `E3-P3B2A-BOUNDARY1`, `E3-P3B2A-REVIEW1`, `E3-P3B2A-DETECT1`, `E3-P3B2A-COMMIT1`.
  - Implementation Gate: P3-B2 accepted and committed.
  - Acceptance:
    - Source: declaration-form loop leaves emit once; assignment-form loops create `0` declarations.
    - Runtime/UI: N/A; ScopeIR loop facts are the real boundary.
    - Data: all prior context outputs remain correct.
    - Behavior test: loop declaration/assignment and scope matrices pass.
    - Cleanup: no unapproved artifact remains.
    - Evidence IDs: `E3-P3B2A-IMPACT1`, `E3-P3B2A-SRC1`, `E3-P3B2A-BUILD1`, `E3-P3B2A-TEST1`, `E3-P3B2A-BOUNDARY1`, `E3-P3B2A-REVIEW1`, `E3-P3B2A-DETECT1`, `E3-P3B2A-COMMIT1`.
    - Actual-status rows refreshed: loop declaration extraction.
  - Evidence Targets: loop oracle, declaration/assignment delta, build, Supervisor, detect-changes, commit.
  - Actual-status Update: transition the loop row to `correct`.
  - Commit Boundary: one P3-B2A commit.

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
