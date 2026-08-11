# Anvien Graph Identity Contract and Strict Construction Actual Status

Title: Anvien Graph Identity Contract and Strict Construction
Date: 2026-08-11
Status: P1-A Complete / P1-B Accepted and Committed / P1-C Accepted at Isolated Commit Boundary / P1-D Not Yet Opened
Companion plan: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-plan.md`
Companion evidence: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-evidence.md`
Companion benchmark: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-benchmark.md`

## Purpose

This file records the current source-backed identity/range/occurrence/collision boundary for Child 01. It separates the measured target finding from the DRAFT architecture proposal. P1-B is accepted in isolated commit `6b3a3fc4480c7f4c2291d8004207f50b52d91d21`. P1-C's scoped production mapping, package QA, full build, normal built-runtime fixture oracle, 61-package product regression, zero-trust history-closure review, exact artifact cleanup, and detect-changes pass at its isolated commit boundary. P1-D is not yet opened and later slices remain closed.

## Freshness / Refresh Rules

- The Anvien graph was refreshed with `anvien analyze --force --json` at P1-C implementation HEAD `6b3a3fc4480c7f4c2291d8004207f50b52d91d21` before the current file-detail and impact queries; `anvien status` and `.anvien/meta.json` report the same indexed/current commit and `stale=false`.
- After the production/test edits, the required QA-impact refresh on 2026-08-11 reports `1,566` scanned, `679` parsed-code, `0` failed files, `85,295` nodes, and `124,278` relationships; current test file-detail rows are `stale=false/changedSinceAnalyze=false`.
- The required full build then completed at `214.1s` with built-CLI self-analyze `1,566/679/0` and `95,539` nodes / `134,405` relationships; this inventory is validation evidence, not an occurrence-conservation denominator.
- All file-detail rows below report `stale=false` and `changedSinceAnalyze=false`.
- Target validation is not run in P0; target source, graph, and pre-existing worktree state are preserve-only.
- Future slices must refresh only rows affected by accepted evidence or a completed slice and append a status-refresh row.

## Scope

Target scope:

- declaration/symbol identity and lexical-owner/meaning inputs;
- full range and selection-range facts;
- declaration occurrence conservation through the production ScopeIR-to-graph path;
- the proven graph node collision/replacement boundary;
- deterministic projection and endpoint integrity;
- bounded `time`/`now` target oracle at P1-E.

Out of scope:

- Graph JSON/Ladybug persistence and readers;
- TypeScript binding-pattern extraction, export semantics, module/re-export, ambient/external, and graph-health semantics;
- relationship identity/merge policy unless a later Child 01 impact proves it is required;
- scanner repair, target-source edits, and target analyze in P0.

## Relationship / Impact Evidence

P0 rows retain their recorded baseline evidence at `7042dda8bfc02ee42b40c3a1c0ede89138439481`; the P1-C candidate rows for `internal/resolution/indexes.go` and inspect-only `internal/resolution/emit.go` were refreshed at `6b3a3fc4480c7f4c2291d8004207f50b52d91d21`. Related-file counts are file counts, not symbol counts.

| Unit / File / Surface | File Detail Evidence | Related File Count | Relationship Summary | Impact Note |
|-----------------------|----------------------|--------------------:|----------------------|-------------|
| `internal/resolution/indexes.go` | `E0-P0A-FD1`, `E1-P1C-IMPACT1` | 46 | production workspace/index and graph-ID construction | high file risk; current `graphIDForDef`/`buildWorkspace` impacts CRITICAL |
| `internal/scopeir/range.go` | `E0-P0A-FD2` | 227 | shared four-coordinate range shape | medium file risk; `Range` CRITICAL |
| `internal/scopeir/definition_index.go` | `E0-P0A-FD3` | 225 | standalone definition-index utility | medium file risk; `BuildDefinitionIndex` LOW, test-only |
| `internal/graph/types.go` | `E0-P0A-FD4` | 238 | graph storage, duplicate-ID mutation, index rebuild | high file risk; `Graph.AddNode` CRITICAL |
| `internal/resolution/emit.go` | `E0-P0A-FD5`, `E1-P1C-IMPACT1` | 42 | definition projection, endpoints, separate relationship merge | high file risk; current `emitNode`/`emitDefinitionNodes` impacts CRITICAL; inspect-only for P1-C |
| `internal/scopeir/facts.go` | `E0-P0A-FD6` | 231 | `DefinitionFact`/`ScopeFact` shared contracts | medium file risk; `DefinitionFact` CRITICAL |
| `internal/scopeir/ir.go` | `E0-P0A-FD7` | 229 | owned copies and deterministic normalization | high file risk; inspect-only |
| `internal/scopeir/sort_keys.go` | `E0-P0A-FD8` | 226 | definition/range sort keys | high file risk; inspect-only |
| `internal/providers/tsjs/definitions.go` | `E0-P0A-FD9` | 16 | TS definition construct range, scope binding, visibility input | high file risk; exact functions conditional P1-B/P1-C |
| `internal/providers/tsjs/nodes.go` | `E0-P0A-FD10` | 20 | tree-sitter position conversion | high file risk; `nodeRange` MEDIUM |
| `internal/providers/tsjs/scopes.go` | `E0-P0A-FD11` | 17 | lexical scope candidates and containment | high file risk; inspect-only unless P1-B proves need |
| `internal/providers/tsjs/extract.go` | `E0-P0A-FD12` | 23 | provider orchestration into ScopeIR | high file risk; preserve orchestration |
| `internal/resolution/resolve.go` | `E0-P0A-FD13` | 50 | binding/result orchestration and emission order | high file risk; preserve unless a slice proves plumbing need |
| `internal/resolution/types.go` | `E0-P0A-FD14` | 31 | resolution result/metrics contracts | high file risk; inspect-only |
| `internal/analyze/analyze.go` | `E0-P0A-FD15` | 182 | normal analyze phase orchestration | high file risk; preserve; no P0 edit |

Exact symbol impacts (`--direction upstream --include-tests`) are recorded in `E0-P0A-IMPACT1..E0-P0A-IMPACT2`. CRITICAL/HIGH is a blast-radius warning, not a prohibition.

## Status Rules

| Status | Meaning | Allowed next action |
|--------|---------|---------------------|
| `correct` | Current behavior satisfies the scoped requirement or a preserve-only sibling contract. | Preserve and validate. |
| `partial` | A required fact exists but is incomplete or not projected. | Change only the proven gap after the slice gate. |
| `wrong` | Current behavior contradicts a measured requirement. | Correct only at the evidence-backed owner. |
| `missing` | Required behavior or fact does not exist. | Implement only after owner evidence. |
| `unbound` | A fact exists but is not connected to its real flow or contract. | Bind only at the proven boundary. |
| `fake-or-stub` | Placeholder or fallback behavior is presented as real. | Remove or replace truthfully. |
| `blocked` | Required authority or acceptance evidence is not yet closed. | Do not implement. |

## Current Status Matrix

| Unit | Current State | Required State | Status | Relationship Count | Evidence | Next Plan Decision |
|------|---------------|----------------|--------|--------------------|----------|--------------------|
| problem/acceptance authority | bounded reports prove four ScopeIR occurrences become two graph nodes; report repair design remains DRAFT | use only measured `2/4 -> 4/4` finding and source-backed invariants | correct | N/A | `E0-P0A-REPORT1`, `E0-P0A-VERIFY1` | preserve; ratify mechanics only in P1-A |
| graph-accuracy contract authority | active contract records the production occurrence denominator, current identity inputs/collision owner, endpoint traceability, and separate relationship boundary without prescribing a repair topology | source-backed Child 01 invariants are authoritative while implementation shape remains slice-evidence-driven | correct | 0 related files | `E1-P1A-SOURCE1`, `E1-P1A-CONTRACT1`, `E1-P1A-REVIEW1`, `E1-P1A-COMMIT1` | preserve contract wording; P1-C uses only its recorded source/impact owner boundary |
| Anvien graph baseline/current candidate | P0 baseline at `7042dda8` was `1,557/676/0`, `85,114` nodes, `123,982` relationships after its ledger refresh; the final P1-B pre-commit refresh was `1,564/677/0`, `85,247` nodes, `124,149` relationships; the P1-C pre-flight was `1,564` files, `85,248` nodes, `124,150` relationships; the post-edit QA-impact refresh is `1,566/679/0`, `85,295/124,278`; the full-build self-analyze reports `95,539/134,405` | graph commands use a fresh index and inventory deltas are not treated as behavior proof | correct | N/A | `E0-P0A-GRAPH1`, `E1-P1B-DETECT1`, `E1-P1C-IMPACT1`, `E1-P1C-TEST1`, `E1-P1C-BUILD1` | current build evidence; refresh again before future graph commands after repository changes |
| declaration contract (`DefinitionFact`) | full construct range plus optional declaring-token selection range, label, owner ID, and visibility inputs exist; TSJS supplies selection range from `nameNode`, while unrelated providers omit the optional field | source-backed full/selection position plus identity inputs without inventing a topology | correct | 231 | `E1-P1B-IMPACT1`, `E1-P1B-SOURCE1`, `E1-P1B-TEST1`, `E1-P1B-RUNTIME1` | preserve the P1-B fact shape; graph consumption remains P1-C/P1-E |
| range semantics | TSJS explicitly retains tree-sitter's one-based lines, zero-based UTF-8 byte columns, and exclusive end point; construct/selection ranges survive LF/CRLF, Unicode, tab, normalization, and JSON round-trip; shared `Range`/legacy `PositionIndex` semantics are unchanged | explicit source-backed provider coordinate/encoding/interval semantics and a provider-supported selection range | correct | 227 | `E1-P1B-IMPACT1`, `E1-P1B-SOURCE1`, `E1-P1B-TEST1`, `E1-P1B-RUNTIME1` | preserve TSJS contract; no guessed cross-provider conversion |
| lexical owner / meaning | `buildWorkspace` supplies `scopeByDef` lexical membership to graph identity; provider `Label` remains the meaning prefix and `OwnerID` is part of the occurrence tuple; package, built-runtime, product-regression, review, and detect evidence distinguish scope and `Shared` meaning lanes | distinct lexical owners and provider-distinguished meanings remain distinguishable through the graph identity boundary | correct | 46 identity / 17 TS scope | `E1-P1C-SOURCE1`, `E1-P1C-TEST1`, `E1-P1C-ORACLE1`, `E1-P1C-REVIEW1`, `E1-P1C-DETECT1` | preserve at the isolated P1-C commit boundary; P1-D must run its own fresh pre-flight |
| graph identity | `graphIDForDef` maps the accepted repo-relative source inputs through an unambiguous deterministic tuple; focused/full package tests, the normal built CLI fixture, the 61-package product regression, history-closure review, and detect-changes pass | injective deterministic identity for distinct source occurrences and supported semantic lanes | correct | 46 | `E1-P1C-IMPACT1`, `E1-P1C-SOURCE1`, `E1-P1C-TEST1`, `E1-P1C-BUILD1`, `E1-P1C-ORACLE1`, `E1-P1C-REVIEW1`, `E1-P1C-DETECT1`, `E1-P1C-COMMIT1` | preserve the accepted P1-C boundary; do not reopen it without a failed later gate |
| production occurrence path | all 13 production providers retain occurrence IDs; `defsByFile` remains the denominator; package fixture conservation is `100%`; built runtime emits `10/10` unique definitions and `10/10` `DEFINES` endpoints; product regression, review, and detect-changes pass | every required input occurrence reaches a distinct accepted graph fact | correct | 46 / 42 | `E1-P1C-SOURCE1`, `E1-P1C-TEST1`, `E1-P1C-ORACLE1`, `E1-P1C-REVIEW1`, `E1-P1C-DETECT1`, `E1-P1C-COMMIT1` | preserve the accepted occurrence denominator and identity output |
| graph collision/mutation | duplicate `Graph.AddNode` IDs replace payloads; `Graph.init` does not reject pre-existing duplicate IDs | conflicting canonical facts fail clearly or follow an explicit proven enrichment rule | wrong | 238 | `E0-P0A-SOURCE4`, `E0-P0A-FD4`, `E0-P0A-IMPACT2` | P1-D is not yet opened; refresh its plan status and pre-flight evidence only after Git confirms the P1-C commit boundary |
| definition projection/endpoints | `emitDefinitionNodes` remains unchanged and emits one node/`DEFINES` edge per occurrence; package and built-runtime evidence prove every P1-C definition has its node/`DEFINES` target and all runtime relationships have endpoints, while later range projection remains outside this slice | corrected identity facts reach nodes and every affected endpoint exists | partial | 42 | `E0-P0A-SOURCE8`, `E1-P1C-TEST1`, `E1-P1C-ORACLE1`, `E1-P1C-REVIEW1`, `E1-P1C-DETECT1` | P1-C endpoint behavior is accepted; preserve source, while P1-E range projection remains closed |
| standalone definition index | first-writer duplicate behavior is covered by legacy tests but has zero production callers | retain its existing test contract; do not mislabel it as runtime loss owner | correct | 225 | `E0-P0A-SOURCE3`, `E0-P0A-FD3`, `E0-P0A-IMPACT2` | preserve-only; do not touch in Child 01 |
| analyze/resolution orchestration | direct provider → ScopeIR → workspace → bound resolution → definition emission path is source-proven; no hidden alternate runtime found | use normal production path; change orchestration only with new evidence | correct | 182 / 50 | `E0-P0A-SOURCE5`, `E0-P0A-FD13`, `E0-P0A-FD15` | preserve-only; no P0 edit; inspect each slice |
| target boundary | target source and worktree contain pre-existing user state; oracle file hash and graph snapshot are known | preserve all 13 entries and source; validate `4/4` only in P1-E | correct | external repository | `E0-P0A-BOUNDARY1`, `E0-P0A-BOUNDARY2` | preserve-only; no P0 target analyze or writes |

## Status Refresh Log

| Refresh | Date | Repo Basis | Changed Scope | Status Changes | Evidence | Next Phase Update |
|---------|------|-----------|---------------|----------------|----------|-------------------|
| R0 | 2026-08-10 | correction-closed HEAD `7042dda8`; prior bounded reports | Child 01 identity/range candidates | broad retained candidates pending refresh | `E0-P0A-REPORT1`, `E0-P0A-VERIFY1` | refresh all candidate owners before implementation |
| R1 | 2026-08-10 | fresh Anvien analyze + source/impact/target audits at `7042dda8` | Child 01 P0 source boundary | `BuildDefinitionIndex` causal candidate -> preserve-only; graph identity/collision -> confirmed wrong; range/IR inputs -> partial; orchestration -> preserve-only | `E0-P0A-GRAPH1`, `E0-P0A-SOURCE1..E0-P0A-SOURCE9`, `E0-P0A-FD1..E0-P0A-FD15`, `E0-P0A-IMPACT1..E0-P0A-IMPACT2`, `E0-P0A-BOUNDARY2` | P1-A remains gated by Supervisor; P1-B/P1-C/P1-D work steps narrowed to exact owners |
| R2 | 2026-08-10 | independent Supervisor report at `7042dda8` | P0-A acceptance and next-slice opening | P0 source/impact/status boundary -> accepted; P1-A documentation contract slice -> open; production implementation remains closed until P1-A and later implementation gates | `E0-P0A-REVIEW1` | P1-A may reconcile the source-backed contract; P1-B/P1-C/P1-D/P1-E remain ordered and closed |
| R3 | 2026-08-10 | P1-A contract reconciliation at `5dc9855b` with fresh contract file-detail/impact | graph-accuracy contract authority and P1-B assumptions | contract `partial -> correct`; occurrence denominator/collision/relationship boundaries ratified; implementation topology remains unprescribed | `E1-P1A-SOURCE1`, `E1-P1A-CONTRACT1` | P1-A awaits Supervisor/commit; P1-B remains closed and must use its own exact owner/impact gate |
| R4 | 2026-08-10 | P1-A Supervisor PASS and isolated documentation commit boundary | contract authority and P1-B opening | P1-A `review pending -> accepted`; production remains unchanged; P1-B becomes the only open implementation slice | `E1-P1A-REVIEW1`, `E1-P1A-COMMIT1` | open P1-B pre-flight only; P1-C/P1-D/P1-E remain closed |
| R5 | 2026-08-10 | P1-B production diff at `8cbf6006` plus full-build, compiled provider boundary, focused matrix, and 61-package product regression | range/selection and identity-input boundary | `DefinitionFact` selection input `partial -> correct`; TSJS coordinate/interval contract `partial -> correct`; lexical/meaning inputs verified correct at ScopeIR while graph consumption remains `partial` | `E1-P1B-IMPACT1`, `E1-P1B-SOURCE1`, `E1-P1B-BUILD1`, `E1-P1B-TEST1`, `E1-P1B-RUNTIME1` | P1-B awaits Supervisor, detect-changes, and isolated commit; P1-C/P1-D/P1-E stay closed |
| R6 | 2026-08-10 | first P1-B Supervisor REJECT, ledger-only correction, and history-closure resubmission PASS | actual-status consistency and review gate | stale Detailed Findings reconciled with matrix/R5; P1-B `review pending -> Supervisor accepted`; production/test behavior unchanged | `E1-P1B-REVIEW1` | run fresh detect-changes and isolated P1-B commit only; P1-C/P1-D/P1-E stay closed |
| R7 | 2026-08-10 | final P1-B candidate after fresh `1,564/677/0` graph refresh and detect-changes | P1-B acceptance and Git boundary | detect-changes PASSes at overall risk `low` with `50` changed symbols / `7` changed files / `0` affected symbols / no affected process; P1-B `Supervisor accepted -> accepted at isolated commit boundary` | `E1-P1B-DETECT1`, `E1-P1B-COMMIT1` | create the isolated P1-B commit; P1-C opens only after that Git boundary and its own pre-flight refresh |
| R8 | 2026-08-10 | clean isolated P1-B commit `6b3a3fc4`; fresh P1-C analyze/file-detail/impact and source-owner audit | P1-C identity/occurrence pre-flight | P1-B Git gate confirmed; `graphIDForDef`/`buildWorkspace` remain the only editable production owner; projection and mutation remain inspect-only; absent historical `graphidentity`/declaration-occurrence modules are not production candidates | `E1-P1B-COMMIT1`, `E1-P1C-IMPACT1` | open P1-C only; production implementation may begin in `internal/resolution/indexes.go`; P1-D/P1-E and target remain closed |
| R9 | 2026-08-11 | P1-C scoped production diff, production-first QA, fresh QA-impact refresh, focused matrix, and resolution-package regression | identity, lexical/meaning consumption, occurrence conservation, endpoint QA | graph identity/occurrence `wrong -> partial` with package boundary PASS; full build/runtime/Supervisor/detect/commit remain pending | `E1-P1C-SOURCE1`, `E1-P1C-TEST1` | continue P1-C validation only; P1-D/P1-E and target remain closed |
| R10 | 2026-08-11 | repo-native full build and built-CLI self-analyze | P1-C build gate | build `pending -> PASS`; production/runtime behavior still requires the fixture boundary | `E1-P1C-BUILD1` | run built CLI fixture oracle; do not open P1-D/P1-E |
| R11 | 2026-08-11 | packaged CLI against repo-local copy of the stable P1-C fixture plus direct graph inspection | P1-C nearest normal runtime/oracle boundary | runtime oracle `pending -> PASS`: `10/10` unique definitions, `time 2/2`, `now 2/2`, separate `Shared` lanes, `0` missing endpoints | `E1-P1C-ORACLE1` | run wider regression and prepare zero-trust review; target/P1-D/P1-E remain closed |
| R12 | 2026-08-11 | first post-build 61-package product regression | P1-C sibling QA | all reported product packages pass except `internal/providers`; its representative heritage relationships exist but the test still expects legacy literal IDs | `E1-P1C-TEST1` | refresh graph and impact the exact provider parity test before QA-only correction; no production scope expansion |
| R13 | 2026-08-11 | fresh provider-parity QA impact, semantic endpoint assertion, 12-language focused matrix, and provider package | P1-C sibling QA repair | stale heritage assertion `FAIL -> PASS`; production/provider behavior unchanged | `E1-P1C-TEST1` | rerun the exact 61-package product regression |
| R14 | 2026-08-11 | exact post-fix 61-package product regression | P1-C full product regression | regression `FAIL -> 61/61 PASS`; only 11 measured non-standalone fixture corpus packages excluded | `E1-P1C-TEST1` | prepare coder report and zero-trust Supervisor review; P1-D/P1-E remain closed |
| R15 | 2026-08-11 | first zero-trust P1-C review plus ledger-only history correction; production/QA hashes unchanged | P1-C actual-status consistency | technical P1-C boundary cleared; review `REJECT` because Detailed Findings still described completed build/runtime/regression gates as pending; stale narrative reconciled without changing implementation or acceptance state | `reports/Supervisor/rp_supervisor_260811_090631_by_gpt-5-codex_child01_p1c.md` | resubmit the ledger-only correction; P1-D/P1-E/target remain closed and `E1-P1C-REVIEW1` remains pending |
| R16 | 2026-08-11 | history-closure Supervisor PASS plus exact current-slice artifact inventory | P1-C review and artifact lifecycle | review `REJECT -> PASS`; production/direct-test/provider-test hashes unchanged; `50/50` consistency assertions PASS; six pre-flight JSON files, one runtime debug tree, and two lane-input Markdown files removed from repo-local `.tmp`, while historical `.tmp/p1c0*` work was preserved | `E1-P1C-REVIEW1` | refresh graph, run detect-changes, record exact evidence, then create the isolated P1-C commit; P1-D/P1-E/target remain closed |
| R17 | 2026-08-11 | mandatory fresh `1,570/679/0` graph and staged all-scope detect-changes across the accepted 19-path boundary | P1-C detect and Git closure | detect-changes exits `0`, risk `medium`, `223` changed symbols / `19` changed files, one affected process (`Main -> CleanPath`, changed step `buildWorkspace` 4/5), zero persisted gaps/degraded nodes/nodes-with-gaps; P1-C `partial -> correct` at its isolated commit boundary | `E1-P1C-DETECT1`, `E1-P1C-COMMIT1` | create the isolated commit immediately; P1-D remains unopened until Git confirms the boundary |

## Phase Touch Map

Related files are not automatically editable. P0 touch mode is inspect-only or preserve-only for every production surface.

| Unit / File / Surface | Plan-Relevant Relationship File | Relationship to Target | Plan Item | Touch Mode | Evidence | Constraint |
|-----------------------|---------------------------------|------------------------|-----------|------------|----------|------------|
| graph-accuracy contract | Child 01 plan and ledgers | authority | P1-A | edit documentation | `E0-P0A-REPORT1`, `E0-P0A-VERIFY1` | exclude DRAFT implementation mechanics |
| `DefinitionFact`/`ScopeFact` | `internal/scopeir/facts.go` | shared IR contract | P1-B/P1-C | P1-B edited only optional `SelectionRange`; `ScopeFact` preserved | `E1-P1B-IMPACT1`, `E1-P1B-SOURCE1` | no redundant lexical/meaning topology; later consumption belongs to P1-C |
| range shape/semantics | `internal/scopeir/range.go`, `internal/providers/tsjs/nodes.go` | shared/provider position input | P1-B | `nodes.go` documents existing TSJS contract; `range.go` preserved | `E1-P1B-IMPACT1`, `E1-P1B-SOURCE1` | no guessed conversion or shared interval rewrite |
| TS definition constructor | `internal/providers/tsjs/definitions.go` | construct range/scope/visibility input | P1-B/P1-C | P1-B edited exact `addDefinition` path to retain `nameNode` range | `E1-P1B-IMPACT1`, `E1-P1B-SOURCE1` | binding-pattern branch remains unchanged for Child 03 |
| TS scope builder | `internal/providers/tsjs/scopes.go` | lexical owner source | P1-B/P1-C | inspect-only | `E0-P0A-FD11`, `E0-P0A-SOURCE7` | preserve containment semantics unless P1 proves gap |
| normalization/sort | `internal/scopeir/ir.go`, `internal/scopeir/sort_keys.go` | deterministic occurrence ordering | P1-C/P1-E | inspect-only | `E0-P0A-FD7`, `E0-P0A-FD8` | no deduplication by name or ID without contract |
| graph identity/index | `internal/resolution/indexes.go` | `DefinitionFact` + existing lexical membership -> `defRef.GraphID` | P1-C | edit only `graphIDForDef`/`buildWorkspace` | `E0-P0A-FD1`, `E0-P0A-IMPACT1`, `E1-P1C-IMPACT1` | CRITICAL warning; no broad resolver rewrite, provider edit, or new identity package |
| identity behavior QA | `internal/resolution/identity_occurrence_test.go`, `internal/resolution/testdata/p1c_identity_repo/src/identity.ts`, affected legacy resolution tests | identity/relationship behavior and stale literal-ID consumers | P1-C | edit tests/fixture only after production behavior | `E1-P1C-TEST1` | semantic relationship tests resolve actual definition nodes; no compatibility fallback or target-derived fixture |
| graph mutation | `internal/graph/types.go` | duplicate canonical node handling | P1-D | inspect-only; exact `Graph.AddNode` candidate | `E0-P0A-FD4`, `E0-P0A-IMPACT2` | classify 11 production caller families and enrichment first |
| node projection | `internal/resolution/emit.go` | identity/range -> graph node/endpoints | P1-C/P1-E | inspect/validate; edit only if impact proves need | `E0-P0A-FD5`, `E0-P0A-SOURCE8` | `emitRelationship` is separate and preserve-only |
| standalone `DefinitionIndex` | `internal/scopeir/definition_index.go` | legacy test utility | P1-C/P1-D | preserve-only / validate-only | `E0-P0A-FD3`, `E0-P0A-IMPACT2` | do not claim it is production loss owner |
| resolution/analyze orchestration | `internal/resolution/resolve.go`, `internal/resolution/types.go`, `internal/analyze/analyze.go` | normal command flow | all P1 slices | preserve-only / inspect-only | `E0-P0A-FD13..E0-P0A-FD15`, `E0-P0A-SOURCE5` | no orchestration change absent new proof |
| target source/graph/worktree | `E:\cheapapp.org` | bounded oracle and external boundary | P1-E | preserve-only / validate-only | `E0-P0A-BOUNDARY1`, `E0-P0A-BOUNDARY2` | no source copy, manual graph edit, or P0 analyze |

## Detailed Findings

### Identity, lexical owner, and meaning

Current state:

The source-backed ScopeIR path retains distinct `DefinitionFact.ID` values, exact lexical scope ownership/bindings, provider `Label`, and member `OwnerID`. P1-C consumes provider occurrence ID, lexical scope, owner, label/meaning, semantic name, arity, and cleaned repository-relative file at the existing graph-ID owner. Package tests prove the accepted inputs are distinguishable and invariant to fact ordering; the normal built-runtime fixture and 61-package product regression pass at the owned boundary.

Required state:

```text
source occurrence + lexical owner + provider-supported meaning
-> deterministic injective graph identity
-> no name-only/range-only collapse
```

Evidence: `E0-P0A-SOURCE1`, `E0-P0A-SOURCE7`, `E1-P1B-SOURCE1`, `E1-P1C-IMPACT1`, `E1-P1C-SOURCE1`, `E1-P1C-BUILD1`, `E1-P1C-TEST1`, `E1-P1C-ORACLE1`.

Classification: `correct` at the provider/ScopeIR, package, built-runtime, product-regression, Supervisor, and detect/commit boundaries.

Allowed next action: preserve the accepted P1-C identity behavior; after Git confirms this isolated commit, refresh P1-D status/pre-flight before any collision-policy work.

Forbidden next action: import the problem report's DRAFT two-tier topology, add only line/range to the old ID, use random/absolute-path identity, or fall back to a global name.

### Ranges and selection range

Current state:

`DefinitionFact` now carries the existing construct `Range` plus an optional declaring-token `SelectionRange`, and TSJS `addDefinition` supplies both from the construct and `nameNode`. TSJS `nodeRange` explicitly records its source-backed one-based lines, zero-based UTF-8 byte columns, and exclusive end point. Shared `Range`, legacy `PositionIndex`, and TSJS scope containment remain unchanged; definition graph nodes still drop columns/selection data at the later P1-E projection boundary.

Required state:

```text
provider position facts -> explicit coordinate/interval semantics
                       -> full declaration range + identifier selection range
                       -> stable identity/projection inputs
```

Evidence: `E0-P0A-SOURCE2`, `E0-P0A-FD2`, `E0-P0A-FD10`, `E1-P1B-SOURCE1`, `E1-P1B-TEST1`, `E1-P1B-RUNTIME1`.

Classification: `correct` at the TSJS provider/ScopeIR P1-B boundary; `partial` only at the later graph projection boundary.

Allowed next action: preserve the accepted P1-B range/position inputs and the completed P1-C identity consumption; graph range projection remains closed until P1-E opens after the required earlier slice boundaries.

Forbidden next action: silently convert byte columns to UTF-16, assume half-open/inclusive semantics, or make range alone the identity key.

### Occurrence conservation and collision

Current state:

The production `defsByFile` slice preserves each incoming definition occurrence. The identity mapping gives each independently evidenced occurrence a unique GraphID at the P1-C package boundary; package QA proves `len(defsByFile[file]) == len(ScopeIR.Definitions)`, and the normal built-runtime fixture proves `10/10` unique definition IDs, `10/10` `DEFINES` endpoints, and zero missing relationship endpoints. The 61-package product regression passes. `Graph.AddNode` conflict policy remains unchanged and belongs to the still-closed P1-D slice. `BuildDefinitionIndex` remains preserve-only.

Required state:

```text
input occurrence denominator -> accepted graph occurrences
                         100% with zero unexplained drops
```

Evidence: `E0-P0A-SOURCE3`, `E0-P0A-SOURCE4`, `E0-P0A-SOURCE5`, `E0-P0A-SOURCE8`, `E1-P1C-IMPACT1`, `E1-P1C-SOURCE1`, `E1-P1C-TEST1`, `E1-P1C-ORACLE1`.

Classification: `correct` at the P1-C package/built-runtime/product-regression/Supervisor/detect/commit boundary. Graph mutation policy remains separately `wrong` for the not-yet-opened P1-D slice; the standalone definition index remains `correct/preserve-only` relative to its tests.

Allowed next action: preserve P1-C occurrence conservation. Do not classify or edit P1-D until Git confirms this isolated commit and its own fresh pre-flight opens.

Forbidden next action: blanket-change every `AddNode` caller, merge relationship deduplication into node identity, or use output node counts alone as conservation proof.

### Projection and endpoint integrity

Current state:

`emitDefinitionNodes` emits a definition node and `DEFINES` edge per `defsByFile` entry. Package and normal built-runtime evidence prove every P1-C definition has a node/`DEFINES` target and the fixture has zero missing relationship endpoints. Projection still does not carry columns, selection range, provider ID, or lexical scope, and endpoint existence alone does not settle the separate P1-D conflict policy or later P1-E range projection.

Required state:

```text
corrected identity/range facts -> graph node properties
                             -> existing endpoints with no silent loss
```

Evidence: `E0-P0A-SOURCE8`, `E0-P0A-FD5`, `E0-P0A-IMPACT2`, `E1-P1C-TEST1`, `E1-P1C-ORACLE1`.

Classification: P1-C identity endpoint behavior is `correct` at its owned package/runtime boundary; range/selection projection remains `partial` for the closed P1-E slice, and P1-C overall remains pending only its process gates.

Allowed next action: preserve `emit.go`; do not open P1-E or edit projection from P1-C closure.

Forbidden next action: treat `emitRelationship`'s semantic merge as declaration conservation or broaden Child 01 into persistence/readers.

### Orchestration and target boundary

Current state:

The normal production flow is source-proven and has no hidden definition-index stage. The target has a dirty pre-existing worktree and a pinned source/graph baseline; no target analyze was run in P0.

Classification: orchestration `correct/preserve-only`; target boundary `correct/preserve-only`.

Allowed next action: the normal-runtime P1-C fixture validation, review, cleanup, and detect/commit boundary are complete. P1-D/P1-E and the target oracle remain closed until their ordered status/predecessor gates open.

Forbidden next action: copy target source into Anvien, manually edit target graph, or treat pre-existing target files as this campaign's artifacts.

## Next Phase Status Decisions

| Plan Item | Actual Status Finding | Required Status / Next-Action Update |
|-----------|-----------------------|--------------------------------------|
| P1-A | source-backed contract reconciliation is Supervisor-accepted and committed; DRAFT architecture remains non-authoritative | complete; preserve accepted contract wording |
| P1-B | range/selection and owner/meaning input behavior is implemented, verified, Supervisor-accepted, detect-clean, and committed as `6b3a3fc4480c7f4c2291d8004207f50b52d91d21`; shared `Range` remains unchanged despite CRITICAL blast radius | complete and preserve |
| P1-C | scoped identity mapping, production-first package QA, full build, built-runtime fixture oracle, 61-package product regression, history-closure Supervisor PASS, exact artifact cleanup, and detect-changes are complete; `defsByFile`/runtime conservation is `100%` | complete at the isolated P1-C commit boundary; preserve until a later owning gate proves a regression |
| P1-D | `Graph.AddNode` is the proven silent replacement owner with 11 production caller families; P1-D is not yet opened | after Git confirms the P1-C boundary, refresh plan/actual status and P1-D's own graph/file-detail/impact/caller pre-flight before editing |
| P1-E | target boundary is captured but target analyze is intentionally not run in P0; P1-E is closed | preserve the exact 13-entry pre-state and hashes; do not run the target oracle until P1-B through P1-D are accepted and committed |

## Implementation Gate

- [x] Target scope is listed in the Current Status Matrix.
- [x] Each target unit has a current implementation-HEAD status.
- [x] Each status has exact evidence IDs.
- [x] Each candidate target file has fresh file-detail evidence and a file count.
- [x] Phase Touch Map lists plan-relevant relationship files and touch modes.
- [x] Correct and preserve-only surfaces are explicit.
- [x] Partial/wrong surfaces have exact next actions and forbidden fallbacks.
- [x] Target/pre-existing boundary is recorded.
- [x] Next-phase assumptions, next actions, and work steps are refreshed from P0 evidence.
- [x] Status Refresh Log has R0 and fresh R1 rows.
- [x] P0 Supervisor review passes.

## Final P0 Decision

- [ ] P0 actual-status incomplete. Implementation is blocked.
- [ ] P0 complete. Next phase can proceed unchanged.
- [x] P0 complete. Next phase status, next action, or work steps must be updated before implementation.
- [ ] P0 complete. Target scope is preserve-only.
- [ ] P0 complete. Implementation is blocked by missing authority or evidence.

Decision note:

The source, fresh graph, file-detail, impact, source-flow, and target-boundary evidence closes the P0 classification. P1-A and P1-B are accepted and committed. P1-C production, package QA, full build, built-runtime fixture oracle, 61-package regression, history-closure Supervisor PASS, exact artifact cleanup, and detect-changes are accepted at its isolated commit boundary. P1-D is not yet opened; P1-E and target execution remain closed.
