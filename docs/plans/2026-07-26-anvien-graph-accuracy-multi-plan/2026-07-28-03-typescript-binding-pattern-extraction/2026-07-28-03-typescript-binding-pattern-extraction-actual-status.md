# Anvien TypeScript Binding-Pattern Extraction Actual Status

Title: Anvien TypeScript Binding-Pattern Extraction
Date: 2026-08-14
Status: Active / P0-A Accepted / Isolated Documentation-Evidence Commit Boundary Prepared / P3-A Closed Until Git Confirms Commit
Companion plan: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-03-typescript-binding-pattern-extraction/2026-07-28-03-typescript-binding-pattern-extraction-plan.md`
Companion evidence: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-03-typescript-binding-pattern-extraction/2026-07-28-03-typescript-binding-pattern-extraction-evidence.md`
Companion benchmark: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-03-typescript-binding-pattern-extraction/2026-07-28-03-typescript-binding-pattern-extraction-benchmark.md`
Roadmap: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-26-anvien-graph-accuracy-roadmap.md`
Predecessor child: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-02-current-graph-persistence-and-reader-consistency/2026-07-28-02-current-graph-persistence-and-reader-consistency-plan.md`
Successor child: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-04-typescript-export-semantics/2026-07-28-04-typescript-export-semantics-plan.md`

## Purpose

This file classifies the current binding-pattern scope before implementation. Child 02 is closed at entry HEAD `181b8cb800f5fe34fa6fe85ddd359f514ead9fb0`. Historical investigation remains bounded evidence; R2 and R2-C1 replace historical ownership assumptions with current-source, graph, file-detail, and impact evidence. Independent Supervisor report `E0-P0A-REVIEW1` accepts the corrected inventory; P3-A remains closed until the isolated P0-A commit is confirmed.

Detailed proof belongs in the evidence ledger. This file stores classifications, touch modes, plan consequences, and status transitions.

## Freshness / Refresh Rules

- Complete a fresh P0 against the current HEAD before P3-A.
- Refresh affected rows after every accepted slice and before the next slice opens.
- Append refresh-log rows; do not erase the bounded baseline or earlier transitions.
- Use exact evidence IDs. A reserved pending ID is not proof.
- If fresh source evidence changes an owner or affected surface, update the next slice before production code is edited.

## Scope

Target scope:

- TypeScript binding-pattern facts and recursive leaf extraction;
- variable, parameter, catch, and `for-of`/`for-in` declaration contexts;
- declaration-versus-assignment behavior, type-inference independence, lexical shadowing, import non-duplication, diagnostics, graph projection, and bounded target resolution;
- only persistence or reader owners proven directly affected by fresh P0/impact evidence.

Out of scope:

- export, module/re-export, and ambient/external semantics;
- unrelated artifact behavior or broad reader policy;
- scanner remediation, blanket transport/cache changes, unrelated refactors, and target-source edits.

## Relationship / Impact Evidence

R2 froze `15` production owners and `12` exact test owners before file-detail/impact. Production relationship/impact metrics are current at entry HEAD; HIGH/CRITICAL remain scope warnings.

| Unit / File / Surface | Current Evidence | Related File Count | Relationship Summary | Impact Note |
|-----------------------|------------------|-------------------:|----------------------|-------------|
| `internal/providers/tsjs/{extract,nodes,definitions,scopes,references,types}.go` | `E0-P0A-SRC1`, `E0-P0A-FD1`, `E0-P0A-IMPACT1` | 23 / 20 / 16 / 17 / 15 / 17 | collector traversal, identifier gate, declaration/type/scope/reference contexts | file impact: CRITICAL / CRITICAL / CRITICAL / MEDIUM / LOW / LOW |
| `internal/scopeir/{facts,ir,sort_keys}.go` | same | 239 / 237 / 234 | current Definition/Binding/Scope contracts, deterministic normalization/sorting | file impact: CRITICAL / CRITICAL / HIGH |
| `internal/providers/tsjs/imports.go` | same | 15 | separate import-binding pipeline | preserve-only; file impact LOW |
| `internal/resolution/{indexes,emit,resolve}.go` | same | 49 / 42 / 53 | index accepted facts, project Definition/`DEFINES`, lexical resolution/gaps | inspect-only; all file impact CRITICAL |
| `internal/graphaccuracy/{property_access,source_site_accuracy}.go` | same | 6 / 6 | consumes existing graph properties/unresolved diagnostics | inspect-only; HIGH / CRITICAL; does not recover omitted leaves |
| `E:\cheapapp.org` source/worktree | `E0-P0A-VERIFY1`, `E0-P0A-VERIFY2` | N/A | real bounded integration target | preserve-only; normal target-local analyzer output only |

## Status Rules

| Status | Meaning | Allowed next action |
|--------|---------|---------------------|
| `correct` | Current evidence proves the required behavior. | Preserve; add regression evidence only when needed. |
| `partial` | Some required behavior exists, but the contract is incomplete. | Change only the missing behavior after impact evidence. |
| `wrong` | Accepted evidence proves behavior conflicts with the requirement. | Repair the exact responsible boundary. |
| `missing` | A required behavior or contract does not exist. | Implement the missing responsibility only. |
| `unbound` | A fact exists but is not connected to the real downstream flow. | Bind it at the proven owner. |
| `fake-or-stub` | Placeholder behavior is being presented as real. | Remove or replace it with truthful behavior. |
| `blocked` | Current authority, ownership, or evidence is incomplete. | Do not implement until P0 resolves it. |

## Current Status Matrix

| Unit | Current State | Required State | Status | Relationship Count | Evidence | Next Plan Decision |
|------|---------------|----------------|--------|-------------------:|----------|--------------------|
| Problem authority | Original report is DRAFT; causal synthesis and Supervisor PASS verify only the bounded defect | problem findings retained; proposed architecture treated as non-authoritative | `correct` | N/A | `E0-P0A-ORIGIN1`, `E0-P0A-VERIFY1`, `E0-P0A-VERIFY2` | preserve this evidence hierarchy |
| Child 02 predecessor handoff | Pn-C closure is entry HEAD `181b8cb8`; exact persistence/reader guarantees and non-claims recorded | predecessor closed without pre-accepting Child 03 | `correct` | predecessor boundary | `E0-P0A-HANDOFF1` | P0-A candidate eligible; P3-A still closed |
| Current production owner inventory | frozen `15` production + `12` exact test owners; `35` explicit exclusions; duplicates/unassigned `0` | exact editable/inspect-only/preserve-only ownership | `correct` | 27 owners | `E0-P0A-SRC1`, `E0-P0A-FD1`, `E0-P0A-IMPACT1` | independent review of P0-A candidate |
| Identifier-only declarators | current `variable_declarator` identifier path emits DefinitionFact, OwnedDefID and local BindingFact | preserve current identifier behavior | `correct` | definitions FD related 16 | `E0-P0A-SRC1` | preserve in P3-A/P3-B |
| Object/array/alias/rest/default/computed/hole patterns | non-identifier declarator names silently return; recursive walker absent | deterministic legal leaves plus explicit unsupported diagnostics | `wrong / missing` | no emitted leaf facts | `E0-P0A-SRC1` | P3-A contract/walker, then P3-B wiring |
| Binding contract | `BindingFact` exists, but current ScopeIR has no pattern path/rest/default/provenance or extraction-diagnostic contract | typed deterministic pattern contract | `partial` | facts FD related 239 | `E0-P0A-SRC1`, `E0-P0A-FD1` | narrow P3-A ScopeIR extension |
| Parameter definitions/local bindings | identifier parameters can receive a separate TypeBindingFact, but no parameter-specific DefinitionFact/local-binding record with binding-pattern metadata is emitted | one declaration/binding per legal leaf in callable scope | `missing` | type path independent | `E0-P0A-SRC1` | P3-B1 after P3-A/P3-B |
| Catch bindings/scope | no catch-clause binding handler or catch scope | catch-local legal leaves | `missing` | zero owner behavior | `E0-P0A-SRC1` | P3-B2 |
| Loop forms | declaration-form identifier may emit in enclosing scope; patterns reject; assignment forms emit no declaration; plain-identifier writes are absent | loop lexical ownership and declaration/assignment distinction, plus write/reference contract | `partial / wrong` | current declaration path only | `E0-P0A-SRC1` | P3-B2A; retain non-declaration invariant, implement/prove writes/references in owning later slice |
| Type inference independence | Definition emission and type inference are separate, but both independently early-return on non-identifiers | declarations survive inference failure | `partial` | separate owners | `E0-P0A-SRC1` | preserve independence; add regression control |
| Imports | separate import pipeline emits import bindings | no double count from pattern work | `correct` | imports FD related 15 | `E0-P0A-SRC1` | preserve-only |
| Assignment writes/references | assignment-form loops create no declarations, but plain-identifier writes and assignment-destructuring write/reference behavior are not implemented or fully proven | non-declaration invariant plus correct write/reference facts | `partial / missing` | no accepted write contract | `E0-P0A-SRC1` | route to P3-B2A/P3-C; do not claim preserve-only |
| Projection and lexical resolution | every accepted Definition projects to a graph node/`DEFINES`; resolver searches inner-to-parent and fails ambiguity; omitted leaves never arrive | consume valid new leaf facts without name rescue | `correct downstream / missing input` | indexes/emit/resolve related 49/42/53 | `E0-P0A-SRC1` | inspect-only until P3-C evidence proves an edit |
| Unsupported-pattern diagnostics | unsupported/silent paths have no structured extraction diagnostic; graph audit readers only see existing generic diagnostics | structured countable extraction diagnostic | `missing` | no current contract | `E0-P0A-SRC1` | P3-A |
| Target boundary | Target analyzed in place; investigation artifacts remained in Anvien | preserve target source and keep normal operational output target-local | `correct` | N/A | `E0-P0A-VERIFY1`, `E0-P0A-VERIFY2` | preserve in P3-C2 |

## Status Refresh Log

| Refresh | Date | Repo Basis | Changed Scope | Status Changes | Evidence | Next Phase Update |
|---------|------|------------|---------------|----------------|----------|-------------------|
| R0 | 2026-08-10 | documentation correction against full problem-origin and bounded-verification reports | Child 03 plan authority and scope | removed unrelated campaign assumptions; P0 reset to incomplete because current owner/impact proof is pending | `E0-P0A-RULE1`, `E0-P0A-SKILL1`, `E0-P0A-ORIGIN1`, `E0-P0A-VERIFY1`, `E0-P0A-VERIFY2` | run fresh P0; do not open P3-A yet |
| R1 | 2026-08-14 | accepted Child 02 chain through Pn-B `9b65f3ef3f3f377d0dcb4a9e0cd5eca444cf51c6`; main Pn-C handoff candidate; Child 02 pre-candidate graph `1,628/688/0`, `97,329/136,603` | predecessor dependency and Child 03 P0 entry boundary | predecessor `missing/dependency-blocked -> recorded/pending closure commit`; accepted persistence/read guarantees and explicit binding non-claims added; all current binding owner/impact rows remain blocked pending P0 | `E0-P0A-HANDOFF1` | after Child 02 Pn-C commit, open only P0-A; P3-A and all implementation remain closed |
| R2 | 2026-08-14 | `master` at `181b8cb800f5fe34fa6fe85ddd359f514ead9fb0`; fresh graph `1,629/688/0`, `97,340/136,614`; status current | complete P0-A source/owner/impact inventory | five current-state IDs recorded; owner boundary resolved; P3-A generic owner/work-step text narrowed | `E0-P0A-GRAPH1`, `E0-P0A-SRC1`, `E0-P0A-FD1`, `E0-P0A-IMPACT1`, `E0-P0A-STATUS1` | keep P3-A closed pending independent Supervisor PASS and isolated P0-A commit |
| R2-C1 | 2026-08-14 | authorized evidence-repair pass at `181b8cb8`; fresh graph `1,631/688/0`, `97,369/136,643`; prior file-detail/impact gates preserved at original capture boundary | exact Set A/B/C path manifest and reproducible partition checks | evidence-integrity blocker repaired: `A=62`, `B=27`, `C=35`, duplicates `0`, overlap `0`, union gaps `0`, assigned `27/27`, unassigned `0`; QA candidate remained unchecked pending review | `E0-P0A-SRC1`, `E0-P0A-STATUS1`, QA report durable manifest section | obtain fresh independent Supervisor review; keep P3-A closed |
| R2-C2 | 2026-08-14 | fresh independent Supervisor re-review at `181b8cb8`; candidate hash `AE4EA64DBD8BE18DB709BD49A21E4E7DAB38380B37F4ED1633532148530A64C2` | P0-A manifest completeness and temporal evidence separation | `PASS`; exact `62` rows, `B=27`, `C=35`, zero overlap/union gaps, all 13 provisional removals reconciled; two presentation notes are non-blocking | `E0-P0A-REVIEW1` | run detect-changes and create isolated P0-A documentation/evidence commit; do not open P3-A |
| R2-C3 | 2026-08-14 | fresh graph refresh after Supervisor PASS: `1,632/688/0`, graph `97,380/136,654`; `anvien detect-changes --repo E:\Anvien --scope all` exit `0` | isolated P0-A docs/evidence impact | `4/4` changed/affected docs, `21` changed sections, LOW risk, zero affected processes/gaps/degraded nodes | `E0-P0A-DETECT1` | stage exact P0-A artifacts, run staged detect confirmation, then commit; P3-A remains closed |
| R2-C4 | 2026-08-14 | byte-clean QA report hash `5B2A567E965521C680B9254873C8049DA69EB3A3F8C21767AF5C13B215DF4857`; fresh Supervisor report hash `023AD68F321B9C7755E52C56FCDDD41F761BD1CAF0DE38583443207AE0D96C43` | final report integrity before isolated commit | `PASS`; exact manifest remains `62/27/35`; no content/scope change beyond whitespace cleanup | `E0-P0A-REVIEW2` | include clean report in staged boundary; rerun fresh detect and commit; P3-A remains closed |
| R2-C5 | 2026-08-14 | final staged-boundary graph `1,633/688/0`, `97,388/136,662`; final detect exit `0` | complete P0-A docs/evidence change inventory | `8/8` changed/affected docs/reports, `62` sections, LOW risk, docs/reporting only, zero affected processes/flows/gaps/degraded nodes | `E0-P0A-DETECT1` | record staged detect confirmation and create isolated P0-A commit; P3-A remains closed |

## Phase Touch Map

| Unit / File / Surface | Plan-Relevant Relationship | Plan Item | Touch Mode | Evidence | Constraint |
|-----------------------|----------------------------|-----------|------------|----------|------------|
| future `internal/providers/tsjs/binding_patterns.go` | new one-responsibility recursive leaf walker | P3-A | future edit | `E0-P0A-SRC1`, `E0-P0A-STATUS1` | create only after independent gate |
| `internal/scopeir/facts.go`, `ir.go`, `sort_keys.go` | pattern contract, storage, normalization and ordering | P3-A | future edit | FD related 239/237/234; impact CRITICAL/CRITICAL/HIGH | smallest compatible contract extension |
| `internal/providers/tsjs/extract.go`, `nodes.go` | collector entry/traversal and AST helpers | P3-A | inspect-only | FD related 23/20; impact CRITICAL/CRITICAL | edit only if P3-A impact re-proves need |
| `internal/providers/tsjs/definitions.go` | variable declaration gate and accepted Definition/local binding emission | P3-B | future edit | FD 16; impact CRITICAL | preserve identifier branch |
| `internal/providers/tsjs/scopes.go`, `references.go`, `types.go` | lexical scopes/reference walk/type inference | P3-B/B1/B2/B2A | future edit or inspect-only per slice | FD 17/15/17; impact MEDIUM/LOW/LOW | type inference must not gate declarations |
| `internal/providers/tsjs/imports.go` | import binding pipeline | P3-A/P3-B | preserve-only | FD 15; impact LOW | exactly-once import behavior |
| assignment write/reference surfaces | assignment-form non-declaration invariant is established; write/reference behavior remains partial/missing | P3-B2A/P3-C | inspect-only/deferred | `E0-P0A-SRC1`, later slice evidence required | do not classify all assignment behavior as correct or preserve-only |
| `internal/resolution/indexes.go`, `emit.go`, `resolve.go` | indexes, graph projection/DEFINES, lexical resolution/gaps | P3-C | inspect-only until changed-fact proof | FD 49/42/53; all impact CRITICAL | no name-based rescue |
| `internal/graphaccuracy/property_access.go`, `source_site_accuracy.go` | graph/read audit of already-emitted properties/diagnostics | P3-C | inspect-only | FD 6/6; impact HIGH/CRITICAL | no extraction recovery logic |
| 12 exact test-owner files listed in QA report | contract-specific validation boundaries | owning P3 slices | future test edit/validation | `E0-P0A-SRC1`, `E0-P0A-FD1` | fixtures are artifacts, not separate owners |
| generic graphhealth/persistence/ResolutionGap/readers, SFC wrappers, fixture data | transparent bridge or generic consumption only | Child 02 / deferred | excluded/deferred | `E0-P0A-SRC1` | reopen only with direct changed-field proof |
| target source/worktree | real integration input | P3-C2 | preserve-only | `E0-P0A-VERIFY1`, `E0-P0A-VERIFY2` | no copied source or target-side report |

## Detailed Findings

### Bounded binding omission

Current evidence:

- The causal synthesis identifies the earliest confirmed divergence in the investigated TS/JS variable-declarator path: a non-identifier declaration name returned before a definition/local binding fact was created.
- Six named array-bound identifiers were present in the AST and absent from the graph.
- Their downstream uses became gaps. This is an upstream extraction failure; downstream commands cannot reconstruct missing facts.

Required state:

```text
legal binding leaf
  -> one declaration fact
  -> one binding fact with path/scope/ranges
  -> one graph occurrence
  -> references resolve to the correct lexical symbol
```

Classification: `wrong` for the bounded baseline; current ownership remains `blocked` until fresh P0.

Allowed next action: revalidate the current source path and impact, then implement P3-A in the exact owner.

Forbidden next action: special-case the six target names, infer all TypeScript pattern support from the bounded finding, or create broad adapter work without affected-field evidence.

### Report architecture versus accepted evidence

Current evidence:

- The problem-origin report explicitly labels its architecture as DRAFT.
- The causal synthesis and final Supervisor report accept the fixed discrepancy record only and explicitly withhold remediation authorization and global accuracy claims.

Classification: evidence hierarchy `correct` after this documentation correction.

Allowed next action: use findings and target counts as the baseline; let current source/impact evidence choose implementation details.

Forbidden next action: copy a proposed architecture into production requirements merely because it appears beside a confirmed defect.

## Next Phase Status Decisions

| Plan Item | Actual Status Finding | Required Next-Action Update |
|-----------|-----------------------|-----------------------------|
| P3-A | no recursive walker or complete pattern/diagnostic contract exists; generic owner boundary was inexact | use focused future walker plus exact ScopeIR owners; remain closed pending independent P0-A Supervisor PASS and isolated commit |
| P3-B | bounded variable omission is wrong | open only after P3-A; require type-inference-failure and assignment/import controls |
| P3-B1 | parameter TypeBinding can exist but parameter Definition/local Binding is missing | implement only declaration/binding context after accepted walker |
| P3-B2 | catch scope/binding owner behavior is missing | implement catch-local declaration context only |
| P3-B2A | loop declaration handling is partial/wrong; assignment forms remain non-declarations | implement declaration forms and lexical scope; preserve assignment distinction |
| P3-C | downstream graph/reference outcome is wrong for the six sites | open only after all context slices; touch only proven graph/persistence owners |
| P3-C2 | target boundary is known and preserve-only | run only after P3-C commit with independent six-site oracle |

## Implementation Gate

- [x] Accepted Child 02 persistence/reader handoff and explicit non-claims are recorded.
- [x] Child 02 Pn-C isolated closure commit is confirmed at `181b8cb8`.
- [x] Current graph and HEAD basis are recorded.
- [x] Current production source and existing tests for every binding context are read in full.
- [x] Every frozen owner file has fresh `file-detail` related-file count evidence (`27/27`).
- [x] Every canonical candidate function/method/exported contract symbol has exact-UID upstream impact evidence (`39/39`).
- [x] Current Status Matrix rows are refreshed from current source, not historical filenames alone.
- [x] Identifier and import preservation are established; assignment-form non-declaration invariant is established.
- [ ] Assignment write/reference behavior is complete and preserve-only (current evidence says partial/missing; route to P3-B2A/P3-C).
- [x] Exact production, test, fixture-artifact, generated-output, and validation ownership is recorded for P3-A.
- [x] P3-A work steps match the discovered current owners.
- [x] Status Refresh Log contains R2.
- [x] Independent Supervisor PASS and fresh docs-only detect evidence are recorded (`E0-P0A-REVIEW1`, `E0-P0A-REVIEW2`, `E0-P0A-DETECT1`).
- [x] Independent Supervisor PASS for the P0-A candidate is recorded (`E0-P0A-REVIEW1`).
- [x] Isolated P0-A commit boundary is recorded (`E0-P0A-COMMIT1`; final hash reported externally by Git).

## Final P0 Decision

- [ ] P0 actual-status incomplete. Implementation is blocked.
- [ ] P0 complete. Next phase can proceed unchanged.
- [x] P0 complete. Next phase status, next action, or work steps must be updated before implementation; Supervisor PASS is recorded, but the isolated P0-A commit is still pending and P3-A stays closed.
- [ ] P0 complete. Target scope is preserve-only.
- [ ] P0 complete. Implementation is blocked by missing authority or evidence.

Decision note:

All five P0-A current-state evidence IDs are recorded with no unresolved owner boundary, and independent Supervisor acceptance is recorded as `E0-P0A-REVIEW1` plus the byte-clean artifact re-review `E0-P0A-REVIEW2`. The accepted non-implementation boundary remains: introduce a focused future `binding_patterns.go` recursive walker; extend only the exact ScopeIR pattern/diagnostic contract owners; preserve established identifier/import behavior and the assignment-form non-declaration invariant; leave assignment write/reference behavior explicitly partial/missing for the owning P3-B2A/P3-C slices; wire variable, parameter, catch and loop declaration contexts only in their ordered slices; keep projection/resolution/audit consumers inspect-only unless later changed-fact evidence proves an edit. P0-A is complete and its exact eight-path documentation/evidence commit boundary is prepared. P3-A cannot open until Git confirms `E0-P0A-COMMIT1`.
