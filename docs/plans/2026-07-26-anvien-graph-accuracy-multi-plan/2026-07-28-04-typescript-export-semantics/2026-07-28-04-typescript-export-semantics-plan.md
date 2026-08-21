# Anvien First-Class TypeScript Export Semantics Plan

## Metadata

- Date: `2026-07-28`
- Status: `P4-A/P4-B/P4-B1/P4-C/P4-C2 committed at their recorded isolated boundaries; aggregate E4-PNA-REVIEW1 and cleanup E4-PNB-REVIEW1 are PASS; Pn-B detect/isolated commit is the sole open boundary; Pn-C, Child 05 and later slices remain locked`
- Plan: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-04-typescript-export-semantics/2026-07-28-04-typescript-export-semantics-plan.md`
- Evidence: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-04-typescript-export-semantics/2026-07-28-04-typescript-export-semantics-evidence.md`
- Benchmark: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-04-typescript-export-semantics/2026-07-28-04-typescript-export-semantics-benchmark.md`
- Actual status: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-04-typescript-export-semantics/2026-07-28-04-typescript-export-semantics-actual-status.md`
- Roadmap: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-26-anvien-graph-accuracy-roadmap.md`
- Contract: `docs/contracts/graph-accuracy-contract.md`
- Predecessor child: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-03-typescript-binding-pattern-extraction/2026-07-28-03-typescript-binding-pattern-extraction-plan.md`
- Successor child: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-05-module-export-and-reexport-resolution/2026-07-28-05-module-export-and-reexport-resolution-plan.md`

## Goal

Represent TypeScript module-export syntax accurately and consistently, separate module export from access visibility, and project the same corrected export facts through the current graph and affected persistence boundary.

## Rules

- Work directly on the current production source and command path. Keep one production graph path; a slice may replace a component only when its source and impact evidence require that change.
- The current provider, graph, and persistence pipeline is baseline evidence, not an immutable design. Change only the owners that fresh P0 source and impact evidence prove necessary.
- Treat `reports/problem/2026-07-26-tinh-chinh-cac-van-de-cua-anvien` as the problem origin. Its remediation architecture is explicitly DRAFT and is not implementation authority.
- Treat the causal synthesis and final Supervisor report as accepted bounded verification of 21 missing export-metadata facts only; they do not approve a remediation design or establish complete module semantics.
- Complete P0 actual status before production edits. Refresh the graph, then record `file-detail` and upstream `impact` for every exact file and symbol opened by a slice.
- Implement production behavior before adding or updating tests. Run the full build before validating the nearest real provider, graph, persistence, or command boundary.
- Update checklist, actual status, evidence, and benchmark immediately when a measured state changes.
- Each production and test file owns one primary semantic responsibility. Do not introduce catch-all helpers or refactor unrelated legacy responsibilities.
- Synthetic fixtures belong under the owning package's `testdata/`; they must be independently authored and must not copy target source.
- Analyze `E:\cheapapp.org` in place only for P4-C2. Its source is preserve-only; normal operational output under its own `.anvien` is the only allowed write there.
- Evidence-bearing artifacts are born durable: oracle rows, raw captures, command streams, manifests, provenance, expected-values inputs/outputs, benchmark material, and reproducibility files must be written directly under an approved `reports/QA/child04-p4c2/...` path. `.tmp` is debug-only; anything evidence-bearing created there is invalid and cannot be promoted, restored, copied, renamed, or used to close an evidence ID.
- P4-C2 separates roles and observation order. The source-only Oracle Author may inspect the authorized target source and target Git/source identity, but may not run or read Anvien analyzer output, target `.anvien`, current implementation behavior, tests, goldens, or QA output as expected-value inputs. QA may observe analyzer output only after Oracle Authoring seals the durable source oracle bundle.
- Scanner remediation and scanner-quality acceptance are outside this child.
- Child 04 owns export syntax and direct export projection. Child 05 owns terminal module/re-export resolution, ambiguity, cycles, and package public-API reachability.
- After each implementation slice: refresh ledgers, remove superseded slice artifacts, obtain Supervisor PASS, run `anvien detect-changes --repo E:\Anvien --scope all`, and commit that slice before opening the next one.
- Hidden fallback is forbidden. Unsupported export syntax must produce a structured, countable extraction diagnostic rather than silently losing a fact.

## Problem

The problem-origin report identifies a module-export metadata defect. The accepted bounded causal synthesis later reproduced the first divergence: 21 selected top-level exported TypeScript declarations existed as definitions, but the investigated provider did not populate export/visibility metadata. The graph emitter serialized the field only when populated, so all 21 appeared unexported to downstream consumers.

The existing active plan had expanded this bounded defect into module resolution and unrelated consumer work. Those are separate boundaries. Child 04 must establish truthful syntax facts and their direct projection only. P0 must revalidate the current source and determine whether the smallest correct change extends an existing fact contract or introduces a dedicated export fact; the report's proposed schema is a candidate, not automatic authority.

## Scope

In scope:

- a current production contract for module-export syntax that is distinct from access visibility;
- direct named/default exports, aliases, anonymous default expressions, type-only exports, and multi-declarator/specifier cases;
- named re-export, star, namespace, default re-export, and type-only re-export syntax facts without terminal resolution;
- value/type/namespace meaning needed to prevent type-only facts from being treated as runtime exports;
- direct-export compatibility fields derived from the same source fact when existing consumers still require them;
- graph projection and only persistence/read owners proven directly affected by P0/impact evidence;
- structured unsupported-syntax diagnostics, focused synthetic fixtures, negative unexported controls, and bounded `21/21` target validation.

Out of scope:

- terminal module resolution, export-table traversal, alias-chain resolution, cycle/ambiguity handling, and package public-API reachability, owned by Child 05;
- binding-pattern extraction and ambient/external resolution;
- unrelated artifact or broad reader policy;
- unrelated changes to CLI, MCP, HTTP, Web, caches, embeddings, groups, processes, or communities without an observed affected contract;
- scanner remediation, unrelated provider refactors, and target-source changes.

## Requirements

- Access visibility and module export are independent concepts. Private/protected/public/internal/package data must not be overwritten to express module export.
- The production source of truth records one export syntax fact per exported binding or export specifier, not one ambiguous flag per statement.
- The exact Go type/file is selected by current source evidence. A dedicated export fact is allowed when existing facts cannot preserve the required semantics; extending a suitable existing owner is allowed when it retains one responsibility.
- Direct named/default declarations, aliases, anonymous default expressions, type-only forms, star exports, namespace exports, and named/default re-exports retain exact source range, selection range, exported name, local name when present, target module/name when syntactically present, and meaning lane.
- Multi-declarator and multi-specifier statements emit one fact for each exported binding/specifier.
- Re-export syntax is recorded without choosing a terminal symbol. Child 04 does not traverse barrels, choose a candidate, or claim package public API.
- A direct-export count means definitions exported directly by their own module. It is distinct from resolved export entries and public API reachability.
- If `visibility`, `isExported`, or another compatibility field remains necessary, it is derived from the same accepted direct-export fact. No second independently maintained export truth is allowed.
- A re-export entry does not make every physical definition in the source module directly exported, and a star entry does not synthesize local declarations.
- Type-only export facts cannot be consumed as runtime value exports. Classes/enums or other dual-meaning declarations use provider evidence rather than name-based guessing.
- Unsupported syntax emits a structured diagnostic containing the source site and reason, and increments a measurable unsupported-export count.
- Graph and affected persistence projections preserve export kind, names, meaning/type-only state, source provenance, and separation from access visibility.
- The exact files and symbols are not predetermined by this plan. P0 and per-slice impact evidence select them; if that evidence changes a slice boundary, update the plan before editing code.

## Acceptance Criteria

- The 21 bounded target definitions have `21/21` correct direct-export facts and graph projection values.
- Direct, default, alias, anonymous-default, type-only, named re-export, star, namespace, and default-re-export synthetic matrices pass with one fact per binding/specifier.
- Access visibility remains unchanged by module-export extraction, and all retained compatibility fields derive from the same export fact with `0` drift.
- Child 04 emits no terminal re-export target, barrel reachability, resolved export-entry count, or package public-API claim.
- Graph JSON and Ladybug or any other persistence owner proven affected by P0 preserve the same export facts with `0` differing affected fields and `0` orphan references.
- Negative unexported controls remain unexported; unsupported fixture cases have `100%` structured diagnostic coverage.
- The target source/worktree remains unchanged apart from documented normal analyzer-owned operational output.
- Each implementation slice has complete source, build, behavior, boundary, Supervisor, detect-changes, and commit evidence.

## Checklist

- [x] P0-A: Complete current actual status before implementation.
  - Goal: replace historical assumptions with current source, graph, impact, and affected-surface evidence for export semantics.
  - Work Steps:
    1. Refresh the Anvien graph and record current HEAD/source basis.
    2. Read the current TS export/import/definition extraction, fact contracts, graph projection, persistence mapping, and relevant tests in full.
    3. Inventory every supported export syntax form, current compatibility field, silent exit, and actual downstream consumer of changed fields.
    4. Run `file-detail` and upstream `impact` for every candidate owner; classify exact edit/inspect/preserve modes and update P4-A if ownership differs.
  - Implementation Gate: no production edit until the actual-status Final P0 Decision permits P4-A.
  - Acceptance: `E0-P0A-RULE1`, `E0-P0A-SKILL1`, `E0-P0A-ORIGIN1`, `E0-P0A-VERIFY1`, `E0-P0A-VERIFY2`, `E0-P0A-SCOPE1`, `E0-P0A-GRAPH1`, `E0-P0A-SRC1`, `E0-P0A-FD1`, `E0-P0A-IMPACT1`, and `E0-P0A-STATUS1` are recorded with no unresolved export-owner boundary.
  - Current State (2026-08-20): fresh excluded graph evidence is current at HEAD `3e25f9ca75c9cb7d59bf228e6cc9aa0b81d738b4`; current source, syntax gaps, compatibility consumers, persistence drift, file-detail counts, and upstream impacts are recorded. The P0 decision requires a P4-A work-step update because the deterministic contract spans four existing owners rather than one file.
  - Commit Boundary: roadmap plus the four Child 04 living ledgers only; documentation-only, no production/test/fixture/generated/target path, no push. P4-A opens only after Git commits this exact boundary.

### P4: First-class TypeScript export semantics

- Phase Goal: preserve export syntax facts and direct-export meaning through extraction and the current graph.
- Phase Boundary:
  - In scope: five ordered slices listed below.
  - Out of scope: binding, terminal module resolution, ambient/external, scanner, and unrelated consumer changes.
  - Dependencies: P0 complete and Child 03 accepted/committed.
- Ordered Slice List:
  - P4-A: Establish the export fact, meaning, and visibility boundary.
  - P4-B: Extract direct/default/alias/type-only export facts.
  - P4-B1: Extract star/namespace/re-export syntax facts.
  - P4-C: Project export facts through the current graph and affected persistence owners.
  - P4-C2: Validate the complete export contract against the real target.

- [x] P4-A: Establish the export fact, meaning, and visibility boundary.
  - Goal: implement the smallest production contract that can represent module-export syntax independently from access visibility.
  - Scope Boundary:
    - Editable: `internal/scopeir/facts.go` for fact/diagnostic shapes, `internal/scopeir/kinds.go` for export kind/meaning enums, `internal/scopeir/ir.go` for the `ScopeIR` export collection and clone/normalize/JSON path, and `internal/scopeir/sort_keys.go` for deterministic export ordering.
    - Test-after-code: `internal/scopeir/scopeir_test.go` and `internal/scopeir/testdata/scopeir.golden.json` only.
    - Inspect-only: `internal/providers/tsjs/{definitions.go,imports.go,extract.go}`, current graph projection/persistence consumers, and Child 05 resolver inputs.
    - Preserve-only: access visibility and existing non-export fact contracts.
    - Out of scope: AST extraction and terminal resolution.
  - Non-Goals: do not make a compatibility boolean the source of truth and do not precompute barrel/public-API state.
  - Pre-flight Questions:
    - Data source: current ScopeIR/provider contracts plus the accepted bounded export finding.
    - Display permission: N/A; no visible or permission-controlled surface changes.
    - DB read flow: N/A; the slice reads in-memory fact contracts only.
    - DB write flow: N/A; it writes in-memory serialization output and package-local test artifacts only.
    - Render location: N/A; inspect fact round trips and ledger evidence.
    - UI behavior flow: N/A; this is a non-UI semantic contract.
    - Docker runtime: N/A; no browser/app runtime changes, but the full build is mandatory.
    - Playwright target: N/A; exercise the real fact serialization boundary.
    - Behavior test: field round trip, meaning lanes, access/export independence, one fact per binding/specifier, and unsupported diagnostic contract.
    - Cleanup/quarantine: keep fixtures under package `testdata/` and remove debug output.
    - External side effects: none.
    - N/A notes: UI, DB, Docker, and Playwright do not participate in this in-memory contract slice.
  - Work Steps:
    1. Refresh graph evidence and exact impacts, then implement the four-owner deterministic contract: fact/diagnostic shapes in `facts.go`, kind/meaning enums in `kinds.go`, `ScopeIR` collection/copy/normalize/JSON wiring in `ir.go`, and ordering in `sort_keys.go`. Do not add AST extraction, graph projection, compatibility-field writes, or terminal resolution in P4-A.
       - UI flow check: N/A; no visible flow exists.
       - DB/data flow check: verify deterministic one-site-to-one-fact serialization and access separation.
       - Render location check: N/A; inspect fact round trips.
       - Mini QA: after the full build in step 2, exercise the real fact boundary; browser controls are N/A.
       - Evidence target: `E4-P4A-IMPACT1`, `E4-P4A-SRC1`.
    2. After production code is correct, add focused tests, run the full build, validate fact round trips and diagnostic counts, refresh ledgers, obtain Supervisor PASS, detect changes, and commit.
       - UI flow check: N/A; no visible flow exists.
       - DB/data flow check: verify exact fields, no second export truth, and no terminal state.
       - Render location check: evidence ledger and focused test output.
       - Mini QA: exercise the built non-UI fact boundary and record command plus result.
       - Evidence target: `E4-P4A-BUILD1`, `E4-P4A-TEST1`, `E4-P4A-BOUNDARY1`, `E4-P4A-REVIEW1`, `E4-P4A-DETECT1`, `E4-P4A-COMMIT1`.
  - Implementation Gate: P0 complete and committed; Child 03 handoff commit `3e25f9ca75c9cb7d59bf228e6cc9aa0b81d738b4` accepted; exact owners and the direct-versus-resolved boundary are recorded. `CRITICAL/HIGH` shared-contract impacts are scope warnings requiring the exact four-file boundary, not edit prohibitions.
  - Acceptance:
    - Source: export kind/names/ranges/meaning/type-only/provenance are representable without changing access visibility.
    - Runtime/UI: N/A; nearest real boundary is production fact serialization after the full build.
    - Data: one fact per binding/specifier; no terminal or public-API state exists in Child 04 facts.
    - Behavior test: contract and separation matrices pass deterministically.
    - Cleanup: no debug or target artifact remains.
    - Evidence IDs: `E4-P4A-IMPACT1`, `E4-P4A-SRC1`, `E4-P4A-BUILD1`, `E4-P4A-TEST1`, `E4-P4A-BOUNDARY1`, `E4-P4A-REVIEW1`, `E4-P4A-DETECT1`, `E4-P4A-COMMIT1`.
    - Actual-status rows refreshed: export contract and access/export separation.
  - Evidence Targets: source diff, serialization matrix, unsupported diagnostic, build, boundary output, Supervisor, detect-changes, commit.
  - Actual-status Update: transition the export-contract row from its P0 classification to `correct`.
  - Current State (2026-08-21): `E4-P4A-IMPACT1`, `E4-P4A-SRC1`, `E4-P4A-BUILD1`, `E4-P4A-TEST1`, `E4-P4A-BOUNDARY1`, independent `E4-P4A-REVIEW1`, and `E4-P4A-DETECT1` are recorded and PASS; `E4-P4A-COMMIT1` is closed at isolated commit `479e8ac229a17f2f6f94be9a4d04e07d74ac4d43`. P4-B is the sole open slice.
  - Commit Boundary: one P4-A commit.

- [x] P4-B: Extract direct/default/alias/type-only export facts.
  - Goal: emit truthful facts for direct declarations and local/default/alias/type-only export syntax.
  - Scope Boundary:
    - Editable: exact direct export extraction owner and thin dispatch proven by impact.
    - Inspect-only: definitions, imports, access modifiers, and P4-A contract.
    - Preserve-only: re-export traversal and non-export definition behavior.
    - Out of scope: star/namespace/re-export syntax and graph projection.
  - Non-Goals: do not resolve target modules or terminal symbols.
  - Pre-flight Questions:
    - Data source: current TypeScript direct/default/local export AST nodes.
    - Display permission: N/A; no visible or permission-controlled surface changes.
    - DB read flow: N/A; the slice reads AST and in-memory definition/fact state.
    - DB write flow: N/A; only in-memory facts and package-local test output are written.
    - Render location: N/A; inspect provider/ScopeIR export facts and ledger evidence.
    - UI behavior flow: N/A; nearest boundary is provider/ScopeIR output.
    - Docker runtime: N/A; no browser/app runtime changes, but the full build is mandatory.
    - Playwright target: N/A; exercise the real direct-export fact boundary.
    - Behavior test: named/default declarations, anonymous default expressions, local aliases, `export {x as default}`, type-only forms, multi-declarator/specifier cases, negative unexported controls.
    - Cleanup/quarantine: keep fixtures under package `testdata/`; remove debug output.
    - External side effects: none.
    - N/A notes: UI, DB, Docker, and Playwright do not participate in this provider slice.
  - Work Steps:
    1. Reconfirm exact extraction-owner impact/current coverage and implement production direct/default/alias/type-only extraction.
       - UI flow check: N/A; no visible flow exists.
       - DB/data flow check: verify one in-memory fact per eligible binding/specifier.
       - Render location check: N/A; inspect provider/ScopeIR output.
       - Mini QA: after the full build in step 2, exercise the direct-export fact boundary.
       - Evidence target: `E4-P4B-IMPACT1`, `E4-P4B-SRC1`.
    2. Add focused tests after code, run the full build, validate syntax facts/negative controls, refresh ledgers, obtain Supervisor PASS, detect changes, and commit.
       - UI flow check: N/A; no visible flow exists.
       - DB/data flow check: verify exact fields/counts and unchanged access visibility.
       - Render location check: evidence ledger and focused test output.
       - Mini QA: exercise the built provider/ScopeIR boundary and record the result.
       - Evidence target: `E4-P4B-BUILD1`, `E4-P4B-TEST1`, `E4-P4B-BOUNDARY1`, `E4-P4B-REVIEW1`, `E4-P4B-DETECT1`, `E4-P4B-COMMIT1`.
  - Implementation Gate: P4-A accepted and committed at `479e8ac229a17f2f6f94be9a4d04e07d74ac4d43`; exact P4-B candidate has independent Supervisor `PASS`.
  - Acceptance:
    - Source: every supported direct/local export binding/specifier emits exactly one fact with correct fields.
    - Runtime/UI: N/A; provider/ScopeIR facts are the real boundary.
    - Data: access visibility unchanged; negative controls remain unexported.
    - Behavior test: direct/default/alias/type-only matrices pass.
    - Cleanup: fixtures/debug output remain only in approved repo paths.
    - Evidence IDs: `E4-P4B-IMPACT1`, `E4-P4B-SRC1`, `E4-P4B-BUILD1`, `E4-P4B-TEST1`, `E4-P4B-BOUNDARY1`, `E4-P4B-REVIEW1`, `E4-P4B-DETECT1`, `E4-P4B-COMMIT1`.
    - Actual-status rows refreshed: direct export extraction and negative controls.
  - Evidence Targets: syntax oracle, fact counts/fields, visibility preservation, build, Supervisor, detect-changes, commit.
  - Actual-status Update: transition direct-export extraction and cleanup to `correct`; leave re-export syntax and all graph/persistence projection pending.
  - Current State (2026-08-21): `E4-P4B-IMPACT1`, `E4-P4B-SRC1`, `E4-P4B-BUILD1`, `E4-P4B-TEST1`, `E4-P4B-BOUNDARY1`, `E4-P4B-REVIEW1`, `E4-P4B-DETECT1`, and `E4-P4B-COMMIT1` are closed at isolated commit `11a37aa8ec0320dd93258c058b088d1070aa778d`; the documentation lag is reconciled and P4-B is not reopened.
  - Commit Boundary: one P4-B commit.

- [x] P4-B1: Extract star/namespace/re-export syntax facts.
  - Goal: record re-export syntax completely without choosing terminal declarations.
  - Scope Boundary:
    - Editable: exact re-export syntax owner and focused test owner.
    - Inspect-only: P4-A contract, direct export extraction, and existing module-request syntax.
    - Preserve-only: access visibility and direct facts.
    - Out of scope: export-table construction, alias traversal, cycles, ambiguity, and public API.
  - Non-Goals: do not treat physical definitions in the referenced file as resolved exports.
  - Pre-flight Questions:
    - Data source: current named/default re-export, star, namespace, and type-only re-export AST nodes.
    - Display permission: N/A; no visible or permission-controlled surface changes.
    - DB read flow: N/A; the slice reads AST, accepted export facts, and existing module-request syntax in memory.
    - DB write flow: N/A; only in-memory syntax facts and package-local test output are written.
    - Render location: N/A; inspect provider/ScopeIR syntax facts and ledger evidence.
    - UI behavior flow: N/A; nearest boundary is provider/ScopeIR output.
    - Docker runtime: N/A; no browser/app runtime changes, but the full build is mandatory.
    - Playwright target: N/A; exercise the real re-export syntax boundary.
    - Behavior test: named/default re-export, `export *`, `export * as ns`, `export type *`, ranges, names, module text, meaning, and one fact per specifier.
    - Cleanup/quarantine: keep fixtures under package `testdata/`; remove debug output.
    - External side effects: none.
    - N/A notes: terminal resolution, UI, DB, Docker, and Playwright are outside this syntax slice.
  - Work Steps:
    1. Record fresh re-export-owner impact/module-request reuse and implement production syntax extraction only.
       - UI flow check: N/A; no visible flow exists.
       - DB/data flow check: verify one immutable in-memory syntax fact per eligible specifier.
       - Render location check: N/A; inspect provider/ScopeIR output.
       - Mini QA: after the full build in step 2, exercise the re-export syntax boundary.
       - Evidence target: `E4-P4B1-IMPACT1`, `E4-P4B1-SRC1`.
    2. Add focused tests after code, run the full build, validate syntax facts and zero terminal state, refresh ledgers, obtain Supervisor PASS, detect changes, and commit.
       - UI flow check: N/A; no visible flow exists.
       - DB/data flow check: verify exact syntax fields/counts and absence of resolution outcomes.
       - Render location check: evidence ledger and focused test output.
       - Mini QA: exercise the built provider/ScopeIR boundary and record the result.
       - Evidence target: `E4-P4B1-BUILD1`, `E4-P4B1-TEST1`, `E4-P4B1-BOUNDARY1`, `E4-P4B1-REVIEW1`, `E4-P4B1-DETECT1`, `E4-P4B1-COMMIT1`.
  - Implementation Gate: P4-B is accepted and committed at `11a37aa8ec0320dd93258c058b088d1070aa778d`; the Child 04/05 ownership boundary is explicit.
  - Acceptance:
    - Source: each supported re-export specifier emits one immutable syntax fact with exact fields.
    - Runtime/UI: N/A; provider/ScopeIR syntax output is the real boundary.
    - Data: terminal target, reachability, ambiguity, cycle, and public-API fields are absent.
    - Behavior test: re-export/star/namespace/type-only syntax matrices pass.
    - Cleanup: no unapproved artifact remains.
    - Evidence IDs: `E4-P4B1-IMPACT1`, `E4-P4B1-SRC1`, `E4-P4B1-BUILD1`, `E4-P4B1-TEST1`, `E4-P4B1-BOUNDARY1`, `E4-P4B1-REVIEW1`, `E4-P4B1-DETECT1`, `E4-P4B1-COMMIT1`.
    - Actual-status rows refreshed: re-export syntax and Child 05 input boundary.
  - Evidence Targets: re-export syntax oracle, zero-derived-state assertion, build, Supervisor, detect-changes, commit.
  - Actual-status Update: transition re-export syntax to `correct`; keep terminal resolution out of scope.
  - Current State (2026-08-21): `E4-P4B1-IMPACT1`, `E4-P4B1-SRC1`, `E4-P4B1-BUILD1`, `E4-P4B1-TEST1`, `E4-P4B1-BOUNDARY1`, independent `E4-P4B1-REVIEW1`, Main-owned `E4-P4B1-DETECT1`, and `E4-P4B1-COMMIT1` are PASS/closed on the exact two-file candidate at isolated commit `42d167aaf28446ac0b3de479a8afefabb8d06736`. Successor verification at clean HEAD `871189b8c6a4e4bb9ff538407232c913b8cf4db6` found the accepted source/test bytes unchanged and preserved both external docs-only commits; explicit continuation authority opens only P4-C under `E4-P4C-AUTH1`. P4-C2/Child 05 stay locked and no push occurred.
  - Commit Boundary: one P4-B1 commit.

- [x] P4-C: Project export facts through the current graph and affected persistence owners.
  - Goal: expose corrected direct export facts and syntax provenance consistently without altering access visibility or implementing resolution.
  - Scope Boundary:
    - Editable: only graph and persistence/read owners proven to consume changed export fields.
    - Inspect-only: completed extraction, current persistence contract, and Child 05 inputs.
    - Preserve-only: access visibility, module dependency edges, and unaffected consumers.
    - Out of scope: unrelated transport/cache changes and terminal export resolution.
  - Non-Goals: do not create a second export truth or populate barrel/public-API outcomes.
  - Pre-flight Questions:
    - Data source: accepted P4-A/P4-B/P4-B1 syntax facts.
    - Display permission: preserve existing access; no new permission behavior is planned.
    - DB read flow: read accepted syntax facts and current graph/persistence contracts proven affected.
    - DB write flow: write the current in-memory graph and isolated test stores only where impact proves persistence is affected.
    - Render location: graph/persisted-graph command output and evidence ledger; no UI is planned.
    - UI behavior flow: N/A unless impact proves an existing public surface directly consumes changed fields; update the plan before expanding.
    - Docker runtime: N/A for the planned non-UI boundary; if P0 proves a served UI is affected, update the plan and require the built Docker runtime.
    - Playwright target: N/A for the planned graph/CLI boundary; required only after an evidence-backed UI scope update.
    - Behavior test: graph fact conservation, direct-export count, compatibility derivation, access separation, affected persistence parity, negative controls, and no terminal fields.
    - Cleanup/quarantine: use isolated repo-local stores and remove debug output; no target run.
    - External side effects: none.
    - N/A notes: no UI/Playwright work is authorized by the current affected-surface evidence.
  - Work Steps:
    1. Refresh graph evidence, map exact fact-to-graph-to-persistence consumers, update the plan if scope changes, and implement projection/compatibility derivation in the smallest proven owners.
       - UI flow check: N/A under the current non-UI scope.
       - DB/data flow check: verify every accepted syntax fact reaches the intended graph record and single-source compatibility value.
       - Render location check: graph/persisted-graph command output.
       - Mini QA: after the full build in step 2, exercise the nearest real graph boundary.
       - Evidence target: `E4-P4C-IMPACT1`, `E4-P4C-SRC1`.
    2. Add focused graph/persistence tests after code, run the full build, validate affected persistence and command output, refresh ledgers, obtain Supervisor PASS, detect changes, and commit.
       - UI flow check: N/A unless an approved scope update added a public surface.
       - DB/data flow check: verify fact conservation, zero drift/orphans, unchanged access fields, negative controls, and zero terminal state.
       - Render location check: evidence ledger and real graph/persisted-graph output.
       - Mini QA: exercise the built graph/loader/CLI boundary and record command plus observed result.
       - Evidence target: `E4-P4C-BUILD1`, `E4-P4C-TEST1`, `E4-P4C-BOUNDARY1`, `E4-P4C-REVIEW1`, `E4-P4C-DETECT1`, `E4-P4C-COMMIT1`.
  - Implementation Gate: P4-A/P4-B/P4-B1 accepted and committed; every editable owner has current impact evidence; no unplanned reader is modified.
  - Acceptance:
    - Source: graph export records preserve syntax facts and direct counts; access visibility remains independent.
    - Runtime/UI: nearest real graph/persisted-graph boundary returns the corrected export facts; no unrelated public surface changes.
    - Data: affected persistence fields differ by `0`, orphan references are `0`, and retained compatibility fields have `0` drift from the source fact.
    - Behavior test: negative controls and zero-terminal-state assertions pass.
    - Cleanup: isolated stores/debug output are removed or retained only as valid evidence.
    - Evidence IDs: `E4-P4C-IMPACT1`, `E4-P4C-SRC1`, `E4-P4C-BUILD1`, `E4-P4C-TEST1`, `E4-P4C-BOUNDARY1`, `E4-P4C-REVIEW1`, `E4-P4C-DETECT1`, `E4-P4C-COMMIT1`.
    - Actual-status rows refreshed: graph projection, compatibility derivation, and affected persistence rows.
  - Evidence Targets: exact blast radius, graph/persistence field oracle, direct counts, access separation, build, Supervisor, detect-changes, commit.
  - Actual-status Update: transition only proven affected projection/persistence rows to `correct`.
  - Current State (2026-08-21): Supervisor REVIEW1 and Main-owned `E4-P4C-DETECT1` are `PASS`; the exact isolated P4-C boundary is committed at `c99c4070b66e7a96be8c9fa2721a0335a1f94877`. Post-commit HEAD matches, `git diff --check` passes, and only two preserved older Main handoff reports remain untracked because their historical blank EOF lines were not rewritten. P4-C2 is now the sole open slice; Child 05 remains locked and target access is authorized only inside the P4-C2 validation lane.
  - Commit Boundary: isolated commit `c99c4070b66e7a96be8c9fa2721a0335a1f94877` contains the accepted P4-C production/test/golden boundary, five living ledgers, Coder report, Supervisor report, and current rotation handoff.

- [x] P4-C2: Validate the complete export contract against the real target.
  - Goal: prove the 21 bounded direct exports and negative controls on the real target without modifying target source or entering Child 05 scope.
  - Current State (2026-08-21): `E4-P4C2-ORACLE1` remains `SEALED`; the bounded repair and fresh post-repair QA pass `21/21` positives, `11/11` negatives, FileContext `17/3/1`, persistence parity `588/0`, integrity/Child 05 zeros and target/config preservation. Independent `E4-P4C2-REVIEW2`, Main-owned `E4-P4C2-DETECT1`, and isolated `E4-P4C2-COMMIT1` are closed; commit `03f09b43f652b9a14b3e49774dc805c0dfd24a27` contains the exact 89-path boundary. Aggregate Pn-A and Pn-B cleanup review are now `PASS`; only the Main-owned Pn-B detect/isolated commit remains open. Pn-C and Child 05 remain locked.
  - Scope Boundary:
    - Oracle Authoring inspect-only: exactly the three accepted hash-pinned target source files plus read-only target HEAD/branch/tracked-status metadata; no target `.anvien` observation.
    - QA inspect-only after seal: target source hashes, target-local analyzer output, and only affected persisted records.
    - Preserve-only: all target source and the complete pre-existing target worktree.
    - Anvien-side writable evidence: new durable oracle and QA bundles under `reports/QA/child04-p4c2/...`, plus this child’s ledgers after valid evidence.
    - Rejection-repair editable boundary after `E4-P4C2-REVIEW1`: `internal/resolution/emit.go` only, subject to fresh graph, file-detail, and upstream file/symbol impact before edit.
    - Rejection-repair test-after-code boundary: `internal/resolution/p4c_export_projection_test.go` and only additional existing P4-C owner tests/goldens proven necessary by fresh impact; production behavior must change first.
    - Rejection-repair inspect/preserve boundary: `internal/filecontext/context.go`, Ladybug CSV/schema/loader, accepted Export facts, QA/oracle artifacts, and unaffected P4-C tests. Edit an inspect/preserve owner only if new evidence changes the plan first.
    - Out of scope: any other production/test/golden repair, target fixtures/reports/probes, copied target source, terminal resolution, other Child defects, and every `.tmp` evidence path.
  - Non-Goals: do not recover/promote `.tmp` captures, do not let analyzer/QA output author or revise expected values, and do not claim global module correctness.
  - Artifact Rule: Oracle Authoring writes the 21 positive rows, 11 negative controls, schema, source basis, provenance, human report, and seal directly under `reports/QA/child04-p4c2/oracle/<oracle_id>/`; QA writes its later raw/normalized run bundle directly under `reports/QA/child04-p4c2/runs/<run_id>/`. Exact schema, file list, and non-circular digest are fixed by `reports/Architect/rp_architect_260821_103812_by_gpt-5_p4c2_evidence_lifecycle.md`; `.tmp` is never an input, staging path, or recovery route.
  - Pre-flight Questions:
    - Expected-value source: direct TypeScript semantics read from the exact hash-pinned target source by a clean-context Oracle Authoring lane; current implementation/analyzer/tests/goldens/QA output are forbidden inputs.
    - Actual-value source: one fresh normal built-analyzer run on the same sealed source basis, performed only by the existing QA continuation after seal verification.
    - Display permission: preserve existing command access; no permission behavior changes.
    - DB read flow: Oracle Authoring reads no target graph/persistence; QA may read target-local graph output and only affected persisted records after seal.
    - DB write flow: Oracle Authoring writes only the durable Anvien-side bundle; QA may write normal target-local `.anvien` operational output only after seal.
    - Render location: durable Anvien-side oracle bundle, later QA report, and evidence ledger; target-side reports are forbidden.
    - UI behavior flow: N/A unless an already approved public surface was affected in P4-C.
    - Docker runtime: N/A for the planned non-UI target boundary; use built Docker only if an approved UI scope exists.
    - Playwright target: N/A for source/graph validation; required only for an approved UI scope.
    - Behavior test: `21/21`, exact kind/name/meaning/type-only/access fields, compatibility parity, negative controls, zero terminal-resolution claims, and target boundary.
    - Cleanup/quarantine: all evidence-bearing artifacts originate and remain in durable Anvien paths; `.tmp` may contain only disposable non-evidence debug trash and never participates in a gate; never write target fixtures/probes.
    - External side effects: normal analyzer-owned target-local graph output only.
    - N/A notes: validation is non-UI unless P4-C recorded an approved public-surface impact.
  - Work Steps:
    1. Main opens a clean/no-history Oracle Authoring lane and keeps QA paused. The lane verifies target HEAD/status plus the three accepted source hashes, reads only those source files, authors exactly 21 positive and 11 negative rows directly in `reports/QA/child04-p4c2/oracle/<oracle_id>/`, writes complete provenance, and writes `seal.json` last. Any source mismatch, forbidden input observation, target write, `.tmp` evidence creation, incomplete row, or seal mismatch stops before analyzer validation.
       - UI flow check: N/A; no visible flow exists.
       - DB/data flow check: source-only; target `.anvien` and analyzer/persistence output are forbidden.
       - Render location check: durable oracle JSON/Markdown bundle only.
       - Mini QA: N/A; Oracle Authoring is not QA.
       - Evidence target: `E4-P4C2-ORACLE1` is recorded as `SEALED`; the immutable bundle and identities returned to Main before QA observation.
    2. Main routes the sealed path/digest to the existing QA continuation. QA verifies the seal/source basis, captures target pre-state, performs the canonical full build and one normal built `anvien analyze E:\cheapapp.org --force`, compares all `21+11` rows without editing expected values, records the target post-state, then proceeds through the existing Supervisor/detect/commit closure gates.
       - UI flow check: N/A unless an approved public surface exists.
       - DB/data flow check: actual output only; expected rows remain immutable; verify field parity, access separation, compatibility, zero diagnostics/orphans/Child 05 state, and unchanged non-`.anvien` target state.
       - Render location check: durable QA run bundle, Child 04 evidence ledger, and later Supervisor report.
       - Mini QA: exercise the real built target analyze/query boundary and inspect human-readable output.
       - Evidence target: `E4-P4C2-TARGET1`, `E4-P4C2-BOUNDARY1`, `E4-P4C2-REVIEW1`, `E4-P4C2-DETECT1`, `E4-P4C2-COMMIT1`.
  - QA Handoff: report `reports/QA/child04-p4c2/runs/p4c2-target-validation-retry-260821_115050+0700/p4c2-qa-retry-validation-report.md`, `11,342` bytes / `200` LF / SHA-256 `C831004F049A563A2387B599BE01C943F5B9416C72B1C45E50A8C1F9D2CEFDB4`; comparison SHA-256 `2C78AB3BDF67D857E5C2A1B75B0F1FDFBFEBE2B70D92E7EBC8EB45A0AC5A3F27`; run digest `9F414A2C54C42F4E39AD8ED03DC042CCC3E1FB5993DA842B22F64851D16AABC4`.
  - Supervisor Handoff: `E4-P4C2-REVIEW1` is `REJECT`; report `reports/Supervisor/rp_supervisor_260821_123030_by_gpt-5_child04_p4c2_durable_retry_review1.md`, `15,671` bytes / `126` LF / canonical SHA-256 `8E37F4B126ABB38A5DCE071C35937F59D9152EFA135B9245408CA138BEC781F7`.
  - Supervisor Re-review: `E4-P4C2-REVIEW2` is `PASS`; report `reports/Supervisor/rp_supervisor_260821_133927_by_gpt-5_child04_p4c2_typealias_compatibility_review2.md`, `13,650` bytes / `120` LF / canonical SHA-256 `5B99A74B1A8D91D48F5E62F0BA1FFCB26317BF818AC6AE044E6CD650B208DC0B`.
  - Rejection Repair Work Steps:
    1. Refresh the self graph, run exact `file-detail` and upstream file/symbol impact for `internal/resolution/emit.go` and the rejected compatibility functions before editing; report HIGH/CRITICAL blast radius and keep the candidate narrow.
    2. Correct production behavior so direct source-export membership drives definition compatibility independently from runtime-value eligibility, while `typeOnly`, meaning lanes, access separation, negatives, and Child 05 exclusions remain unchanged.
    3. Only after production behavior is correct, update the rejected focused expectation and any source-proven affected owner test; run canonical full build, nearest owner/reader boundary, fresh real-target comparison on the sealed basis through the authorized QA workflow, independent Supervisor re-review, detect-changes, and isolated commit.
  - Implementation Gate: closed at isolated commit `03f09b43f652b9a14b3e49774dc805c0dfd24a27`; no P4-C2 gate may be reopened absent evidence invalidation.
  - Stop Conditions: Oracle Authoring stops before seal on source/hash/count/schema/provenance drift, forbidden input/context, target write, copied source, or any evidence-bearing `.tmp`/unapproved-path artifact. QA stops before comparison/acceptance on seal/source mismatch, failed/non-normal build or analyze, old/substituted output, target contamination, or any attempt to edit expected values.
  - Acceptance:
    - Source: `21/21` target definitions have correct direct export facts and access visibility remains independent.
    - Runtime/UI: the real target graph/affected persistence boundary exposes the same corrected values.
    - Data: the sealed oracle has 21 complete unique positive rows and 11 exact negative controls; actual negative controls stay unexported, Child 05-derived fields remain absent, and the target boundary is preserved.
    - Behavior test: syntax-kind, meaning, type-only, compatibility, and unsupported controls pass.
    - Cleanup: no evidence-bearing `.tmp` artifact and no target report, fixture, or probe exists; debug trash is irrelevant to every gate and may be deleted without loss.
    - Evidence IDs: `E4-P4C2-ORACLE1`, `E4-P4C2-TARGET1`, `E4-P4C2-BOUNDARY1`, `E4-P4C2-REVIEW1`, `E4-P4C2-REVIEW2`, `E4-P4C2-DETECT1`, `E4-P4C2-COMMIT1`.
    - Actual-status rows refreshed: all Child 04 rows and Child 05 predecessor state.
  - Evidence Targets: durable source-only authoring bundle, seal identity/chronology, 21-entry target fact/projection comparison, 11 negative controls, durable QA run/boundary manifest, later Supervisor, change check, evidence commit.
  - Actual-status Update: transition the bounded target row to `correct` and append the successor handoff refresh.
  - Commit Boundary: isolated commit `03f09b43f652b9a14b3e49774dc805c0dfd24a27` contains the accepted two-file repair, five living documents, durable Oracle/QA lifecycle, Coder/Supervisor evidence and required P4-C2 provenance; no target or `.tmp` artifact is committed or required.

- [x] Pn-A: Run the Supervisor acceptance loop.
  - Goal: independently verify all five slices, source diff, build, real-boundary results, benchmark, target boundary, and evidence integrity.
  - Work Steps: invoke the Supervisor skill; return only rejected invariants to the owning slice; repeat until PASS or a documented blocker.
  - Implementation Gate: every P4 slice is complete or explicitly blocked.
  - Acceptance: `E4-PNA-REVIEW1` records Supervisor PASS or a precise blocker; no self-acceptance is used.
  - Current State (2026-08-21): independent aggregate `E4-PNA-REVIEW1` is `PASS`; report `reports/Supervisor/rp_supervisor_260821_142429_by_gpt-5_child04_pna_aggregate_review1.md`, `23,563` bytes / `178` LF / canonical SHA-256 `7EBFD5087F8593660A94E70B0816A7FC98944FDE7B9D1F1BC9388CEB9F6DC5A8`; residual same-invariant surfaces none. Pn-B cleanup and its distinct review are also `PASS`; Main-owned Pn-B detect/commit remains open. Pn-C and Child 05 remain locked.
- [ ] Pn-B: Remove dead work created by this child.
  - Goal: leave no superseded fixture, debug output, failed evidence, duplicate report, or rejected approach.
  - Work Steps: inventory child-created artifacts, remove only dead child work, obtain distinct Supervisor review of cleanup, then have Main synchronize the five living documents once, run proportionate change/boundary detection, and create the isolated Pn-B commit.
  - Implementation Gate: Pn-A has completed its first review.
  - Acceptance: `E4-PNB-CLEAN1`, `E4-PNB-REVIEW1`, `E4-PNB-DETECT1`, and `E4-PNB-COMMIT1` prove the retained artifact set is current and close an isolated Git boundary.
  - Current State (2026-08-21): `E4-PNB-CLEAN1` removed only the empty review-induced `.tmp/p4c-tests` parent (`0` files / `0` bytes), and independent `E4-PNB-REVIEW1` is `PASS` with no residual same-invariant surface. Main-owned `E4-PNB-DETECT1` and `E4-PNB-COMMIT1` remain; Pn-C and Child 05 stay locked until the isolated commit succeeds.
  - Commit Boundary: five living documents, cleanup Coder report, cleanup Supervisor report, and explicitly authorized current Main handoff provenance only; preserve the older `0631`/`0721` handoffs untracked unless separately proven admissible.
- [ ] Pn-C: Close Child 04 and hand off Child 05.
  - Goal: record final validation, commit state, immutable syntax-fact boundary, and the exact successor opening condition.
  - Work Steps:
    1. Confirm all slice commits, final build/boundary evidence, benchmark rows, and target boundary.
    2. Run final detect-changes when implementation work is present.
    3. Refresh Child 05 actual status from accepted Child 04 evidence and record the handoff.
  - Implementation Gate: Pn-A and Pn-B PASS, with the isolated Pn-B commit recorded.
  - Acceptance: `E4-PNC-DETECT1`, `E4-PNC-COMMIT1`, and `E4-PNC-HANDOFF1` are recorded; worktree state is known.

## Risk Notes

- The accepted defect proves missing metadata for 21 selected exports, not complete support for every TypeScript export form. General fixtures are required before broader claims.
- Filling only `Visibility="exported"` can improve a count while continuing to conflate access and module semantics; acceptance checks the source fact and field separation.
- Re-export syntax and terminal resolution are different pipelines. Resolving physical definitions during Child 04 would cross into Child 05 and can create false positives.
- Type-only and dual-meaning declarations can look correct by name while exposing invalid runtime exports; tests must verify meaning, not only counts.
- Graph and persistence owners can have HIGH/CRITICAL blast radius. That warning requires narrow scope and regression evidence; it does not prohibit the necessary edit.
- Any newly discovered affected reader or contract changes the slice boundary and requires a plan update before code; it does not expand unrelated consumers automatically.
