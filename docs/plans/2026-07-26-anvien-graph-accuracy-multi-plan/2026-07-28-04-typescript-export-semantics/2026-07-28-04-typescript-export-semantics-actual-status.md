# Anvien First-Class TypeScript Export Semantics Actual Status

Title: Anvien First-Class TypeScript Export Semantics
Date: 2026-07-28
Status: P4-A source/build/boundary/Supervisor/detect PASS / isolated commit pending / P4-B locked until commit success
Companion plan: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-04-typescript-export-semantics/2026-07-28-04-typescript-export-semantics-plan.md`
Companion evidence: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-04-typescript-export-semantics/2026-07-28-04-typescript-export-semantics-evidence.md`
Companion benchmark: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-04-typescript-export-semantics/2026-07-28-04-typescript-export-semantics-benchmark.md`
Roadmap: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-26-anvien-graph-accuracy-roadmap.md`
Predecessor child: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-03-typescript-binding-pattern-extraction/2026-07-28-03-typescript-binding-pattern-extraction-plan.md`
Successor child: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-05-module-export-and-reexport-resolution/2026-07-28-05-module-export-and-reexport-resolution-plan.md`

## Purpose

This file classifies the current TypeScript export-syntax and direct-projection scope before implementation. Historical investigation proves a bounded 21-export defect but does not substitute for fresh current-source, graph, file-detail, impact, and affected-consumer evidence. `E3-PNC-HANDOFF1` now records the accepted Child 03 predecessor closure; it supplies no export implementation evidence and does not waive Child 04 P0.

Detailed proof belongs in the evidence ledger. This file stores classifications, touch modes, plan consequences, and status transitions.

## Freshness / Refresh Rules

- Complete a fresh P0 against the current HEAD before P4-A.
- Refresh affected rows after every accepted slice and before the next slice opens.
- Append refresh-log rows; do not erase the bounded baseline or earlier transitions.
- Use exact evidence IDs. A reserved pending ID is not proof.
- If fresh source evidence changes an owner or affected surface, update the next slice before production code is edited.

## Scope

Target scope:

- TypeScript direct/default/alias/type-only export facts and star/namespace/re-export syntax facts;
- value/type/namespace meaning and separation of module export from access visibility;
- direct-export compatibility derivation, graph projection, and only persistence/read owners proven affected;
- unsupported-syntax diagnostics, negative controls, and bounded `21/21` target validation.

Out of scope:

- terminal export resolution, alias traversal, cycles, ambiguity, and package public API;
- binding and ambient/external semantics;
- unrelated artifact behavior or broad reader policy;
- scanner remediation, blanket transport/cache changes, unrelated refactors, and target-source edits.

## Relationship / Impact Evidence

P0-A resolves the current owners, related-file counts, and blast radii. CRITICAL/HIGH results are scope warnings; each implementation slice must refresh its own exact edit boundary before production changes.

| Unit / File / Surface | Current Evidence | Related File Count | Relationship Summary | Impact Note |
|-----------------------|------------------|-------------------:|----------------------|-------------|
| Current TS export extraction owner | `E0-P0A-SRC1`, `E0-P0A-FD1`, `E0-P0A-IMPACT1` | `imports.go=17`; `definitions.go=24`; `extract.go=25` | `emitExportStatement` is the syntax owner; `extract.go` wires collection; `definitions.go` preserves named declarations | P4-B/B1 may edit only `imports.go`/`extract.go` after slice-local refresh; `definitions.go` is inspect/preserve-only |
| Current fact/meaning owner | `E0-P0A-SRC1`, `E0-P0A-FD1`, `E0-P0A-IMPACT1` | `facts.go=245`; `kinds.go=239`; `ir.go=243`; `sort_keys.go=239` | four-file deterministic module-export fact/meaning/collection/order contract | exact P4-A production boundary; CRITICAL/HIGH shared-contract scope |
| Current graph export projection owner | `E0-P0A-SRC1`, `E0-P0A-FD1`, `E0-P0A-IMPACT1` | `resolution/emit.go=43` | `emitDefinitionNodes` currently maps only `DefinitionFact.Visibility` to `visibility` | inspect-only until P4-C; current impact CRITICAL `26/20/5/1` |
| Current persistence mapping | `E0-P0A-SRC1`, `E0-P0A-FD1`, `E0-P0A-IMPACT1` | `lbugload/csv.go=21`; `lbugschema/schema.go=21` | Ladybug persists field-specific `isExported` and drops `visibility`; Graph JSON preserves arbitrary properties | inspect-only until P4-C proves the final changed fields |
| Existing CLI/MCP/HTTP/Web/cache/derived surfaces | `E0-P0A-SRC1`, `E0-P0A-FD1`, `E0-P0A-IMPACT1` | `filecontext/context.go=44`; `mcp/context.go=28` | file-context is a semantic consumer; CLI/MCP/HTTP expose derived results; Web/generated contracts are carriers | only proven semantic consumers may change in P4-C; transports/carriers remain preserve-only |
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
| Problem authority | Original report is DRAFT; causal synthesis and Supervisor PASS verify only the bounded defect | findings retained; proposed architecture treated as non-authoritative | `correct` | N/A | `E0-P0A-ORIGIN1`, `E0-P0A-VERIFY1`, `E0-P0A-VERIFY2` | preserve evidence hierarchy |
| Child 03 predecessor handoff | Child 03 binding-pattern invariant is accepted and closed at isolated Pn-C commit `3e25f9ca75c9cb7d59bf228e6cc9aa0b81d738b4`; non-claims explicitly exclude export semantics | consume only accepted predecessor facts and independently resolve Child 04 owners | `correct` | `1` inbound handoff/roadmap relationship | `E3-PNC-HANDOFF1`, `E3-PNC-COMMIT1`, `E3-PNA-REVIEW1` | predecessor gate closed; no export behavior inferred from Child 03 |
| Current production owner inventory | Four ScopeIR contract owners, TSJS extraction owners, graph/persistence owners, semantic consumers, and preserve-only carriers are current and impact-classified | exact editable/inspect-only/preserve-only owners with current counts | `correct` | `12` P0 file-detail rows; `4` P4-A final-byte rows | `E0-P0A-SRC1`, `E0-P0A-FD1`, `E0-P0A-IMPACT1`, `E4-P4A-IMPACT1` | preserve accepted P4-A owner boundary; P4-B requires its own extraction-owner refresh after commit |
| First-class export fact contract | Exact six-file candidate adds one source-site `ExportFact`, seven source-form kinds, three meaning lanes, explicit type-only/provenance/diagnostic state, owning deep copies, and deterministic JSON ordering | one canonical deterministic in-memory export fact/diagnostic contract with no later-slice state | `correct` | `4` production owners + `2` test-after-code files | `E4-P4A-SRC1`, `E4-P4A-TEST1`, `E4-P4A-BOUNDARY1`, `E4-P4A-REVIEW1` | preserve after isolated commit; P4-B may consume but not redesign the contract |
| Bounded direct-export definitions | 21 selected definitions were present at the accepted baseline | preserve the definitions and attach truthful direct export facts | `partial` | pending | `E0-P0A-VERIFY1`, `E0-P0A-VERIFY2` | preserve definition extraction; repair export facts |
| Bounded export metadata | 0 of 21 selected definitions carried export/visibility metadata at the accepted baseline | 21 of 21 correct direct export facts and projections | `wrong` | pending | `E0-P0A-VERIFY1`, `E0-P0A-VERIFY2` | P4-A, P4-B, P4-C |
| Access visibility versus module export | P4-A now provides an independent export fact contract; `DefinitionFact.Visibility` and preserve-only projection/compatibility writers are unchanged | independent concepts at the contract boundary; any later compatibility field derives from the accepted fact | `correct` for P4-A contract boundary | `4` production owners + preserve-only siblings | `E4-P4A-SRC1`, `E4-P4A-TEST1`, `E4-P4A-REVIEW1` | preserve separation in P4-B/B1; implement derivation only in P4-C |
| General direct/default/alias/type-only syntax | source-less export statements return without a fact or diagnostic; direct/local/default/type-only forms have no first-class representation | one truthful fact per supported exported binding/specifier | `missing` | `0/4` first-class syntax paths | `E0-P0A-SRC1` | implement fact production in P4-B after P4-A |
| Star/namespace/re-export syntax | named re-export becomes `ImportReexport`; star becomes `ImportWildcard`; namespace alias and type-only marker are lost | immutable export syntax facts without terminal resolution | `partial` | `2/4` compatibility paths; `0/4` export-fact paths | `E0-P0A-SRC1` | replace/derive compatibility from first-class syntax facts in P4-B1 |
| Export graph projection | Historical emitter omitted empty provider metadata, so bounded exports appeared unexported | corrected fact kind/names/meaning/provenance and access separation | `wrong` | pending | `E0-P0A-VERIFY1` | P4-C after extraction |
| Persistence field parity | Graph JSON preserves arbitrary properties, while Ladybug uses label-specific `isExported` columns and drops `visibility` | same affected export facts in every proven persistence owner | `wrong` | `2` persistence owners plus generic Graph JSON | `E0-P0A-SRC1`, `E0-P0A-FD1`, `E0-P0A-IMPACT1` | P4-C reconciles only fields actually emitted by the accepted fact projection |
| Terminal module/re-export resolution | Separate bounded barrel defect exists | Child 05 resolves syntax facts to terminal outcomes | `wrong` but out of scope | pending | `E0-P0A-VERIFY1` | preserve boundary; no Child 05 logic here |
| Target boundary | Target analyzed in place; investigation artifacts remained in Anvien | preserve target source and keep normal operational output target-local | `correct` | N/A | `E0-P0A-VERIFY1`, `E0-P0A-VERIFY2` | preserve in P4-C2 |

## Status Refresh Log

| Refresh | Date | Repo Basis | Changed Scope | Status Changes | Evidence | Next Phase Update |
|---------|------|------------|---------------|----------------|----------|-------------------|
| R0 | 2026-08-10 | documentation correction against full problem-origin and bounded-verification reports | Child 04 plan authority and scope | removed unrelated campaign assumptions; separated syntax/direct projection from Child 05 resolution; P0 reset to incomplete pending current owner/impact proof | `E0-P0A-RULE1`, `E0-P0A-SKILL1`, `E0-P0A-ORIGIN1`, `E0-P0A-VERIFY1`, `E0-P0A-VERIFY2` | run fresh P0; do not open P4-A yet |
| R1 | 2026-08-20 | Child 03 Pn-B commit `0231d7f9cc7cfd05d8d5da4787161a8ec2cd1550`; excluded graph `1,126/626/0`, `80,908/120,167`; roadmap/Child 03/Child 04 actual/handoff governance rows LOW/non-stale with zero upstream affected files/processes/flows/tests | predecessor handoff and successor opening condition | Child 03 predecessor `dependency-blocked -> accepted handoff`; Child 04 P0 remains incomplete, current owner/file-detail/impact/syntax/consumer evidence remains pending | `E3-PNC-HANDOFF1`, `E3-PNB-COMMIT1` | after Pn-C docs-only commit, open only Child 04 P0-A; no P4 implementation or redundant Supervisor loop |
| R2 | 2026-08-20 | Child 03 Pn-C commit `3e25f9ca75c9cb7d59bf228e6cc9aa0b81d738b4`; fresh excluded graph `1,126/626/0`, `80,908/120,167`; full current source plus 12 file-detail and impact rows | Child 04 current owner, syntax, projection, persistence, and consumer boundary | owner inventory `blocked -> correct`; direct syntax `blocked -> missing`; re-export syntax `blocked -> partial`; persistence `blocked -> wrong`; P0-A complete | `E0-P0A-GRAPH1`, `E0-P0A-SRC1`, `E0-P0A-FD1`, `E0-P0A-IMPACT1`, `E0-P0A-STATUS1` | commit exact five-document boundary, then open only P4-A with its updated four-file contract scope |
| R3 | 2026-08-20 | P0-A commit `ff2467bb92f94a9c53c4de030685686700051a98`; exact P4-A candidate; canonical full build and nearest boundary PASS; independent Supervisor report `1B8DEB2F...175573B2`; fresh excluded closure graph `1,130/626/0`, `81,132/120,514`; detect exit `0` | first-class export fact/meaning/diagnostic/deterministic ScopeIR boundary | export fact contract `missing -> correct`; access/export contract separation `wrong -> correct`; direct/re-export extraction and projection classifications unchanged | `E4-P4A-IMPACT1`, `E4-P4A-SRC1`, `E4-P4A-BUILD1`, `E4-P4A-TEST1`, `E4-P4A-BOUNDARY1`, `E4-P4A-REVIEW1`, `E4-P4A-DETECT1` | create isolated P4-A commit; keep P4-B locked until commit success |

## Phase Touch Map

| Unit / File / Surface | Plan-Relevant Relationship | Plan Item | Touch Mode | Evidence | Constraint |
|-----------------------|----------------------------|-----------|------------|----------|------------|
| `internal/scopeir/{facts.go,kinds.go,ir.go,sort_keys.go}` | source of deterministic module-export syntax truth | P4-A | accepted production boundary; preserve after isolated commit | `E4-P4A-IMPACT1`, `E4-P4A-SRC1`, `E4-P4A-REVIEW1` | access visibility remains independent; later slices consume without redesign |
| `internal/providers/tsjs/{imports.go,extract.go}` | produces direct/default/alias/type-only facts | P4-B | inspect now; narrow edit after slice-local refresh | `E0-P0A-SRC1`, `E0-P0A-FD1`, `E0-P0A-IMPACT1` | preserve definitions and non-export behavior |
| `internal/providers/tsjs/{imports.go,extract.go}` | produces named/star/namespace/type-only re-export syntax facts | P4-B1 | narrow edit after P4-B commit and slice-local refresh | `E0-P0A-SRC1`, `E0-P0A-FD1`, `E0-P0A-IMPACT1` | no terminal resolution |
| `internal/resolution/emit.go`, `internal/lbugload/csv.go`, `internal/lbugschema/schema.go`, proven semantic consumers | project/persist/consume exact export facts | P4-C | inspect now; edit only accepted changed-field consumers after fresh impact | `E0-P0A-SRC1`, `E0-P0A-FD1`, `E0-P0A-IMPACT1` | one source fact; only named affected consumers |
| CLI/MCP/HTTP/Web/generated carriers and unaffected caches | transport or non-consuming surfaces | all P4 | preserve-only unless slice evidence proves direct semantic consumption | `E0-P0A-SRC1` | plan update required before expansion |
| Child 05 resolver owners | consume immutable syntax facts | all P4 | preserve-only | successor boundary | no traversal/cycle/ambiguity work here |
| target source/worktree | real integration input | P4-C2 | preserve-only | `E0-P0A-VERIFY1`, `E0-P0A-VERIFY2` | no copied source or target-side report |

## Detailed Findings

### Bounded export-metadata omission

Current evidence:

- The causal synthesis records 21 selected top-level exported TypeScript declarations that survived as definitions.
- At the accepted baseline, none carried the export/visibility metadata expected by graph consumers.
- The first confirmed divergence was provider fact creation; downstream projection serialized only what it received.

Required state:

```text
export syntax site
  -> one immutable export fact with meaning and provenance
  -> direct export projection derived from that fact
  -> affected persistence readers observe the same value

access visibility remains independent
```

Classification: `wrong` for the bounded baseline and current source; P0-A resolves the ownership boundary.

Allowed next action: commit the exact accepted P4-A boundary; only after commit success may P4-B open.

Forbidden next action: fill one field solely to improve counts, treat re-export reachability as direct export, or add transport adapters without affected-field evidence.

### Child 04 versus Child 05 boundary

Current evidence:

- The same causal synthesis identifies barrel lookup as a separate first divergence from export metadata.
- Syntax collection and terminal export resolution have different commands/state/ownership: Child 04 records what source exports; Child 05 determines what terminal symbol a consumer import reaches.

Classification: boundary `correct` after this documentation correction.

Allowed next action: make Child 04 facts complete and immutable enough for Child 05.

Forbidden next action: recurse through barrels, select candidates, or claim public API during Child 04.

## Next Phase Status Decisions

| Plan Item | Actual Status Finding | Required Next-Action Update |
|-----------|-----------------------|-----------------------------|
| Predecessor handoff | Child 03 final evidence/benchmark/target boundary is accepted at `3e25f9ca75c9cb7d59bf228e6cc9aa0b81d738b4`; no export claim is transferred | `correct`; predecessor gate closed |
| P4-A | first-class contract, tests, build, boundary, independent Supervisor acceptance, and closure detect are PASS on the exact candidate | create isolated commit; keep the item open until commit success |
| P4-B | 21 direct definitions exist but metadata is wrong; accepted P4-A contract is available only in the uncommitted candidate | preserve definition extraction; remain locked until P4-A commit success, then refresh extraction-owner impact and implement exact missing syntax facts |
| P4-B1 | named/star compatibility facts exist, while namespace/type-only detail and all first-class export facts are absent | emit syntax only after P4-B; no terminal state |
| P4-C | graph projection, Ladybug persistence, and semantic consumer dialects are classified | open after extraction; touch only consumers of the accepted changed fields |
| P4-C2 | target boundary known and preserve-only | run after P4-C commit with independent 21-entry oracle and negative controls |

## Implementation Gate

- [x] Current graph and HEAD basis are recorded.
- [x] Current production source and existing tests for export facts/extraction/projection/persistence are read in full.
- [x] Every candidate editable file has fresh `file-detail` related-file count evidence.
- [x] Every candidate function/method/exported symbol has fresh upstream impact evidence and reported blast radius.
- [x] Current syntax coverage and unsupported paths are inventoried without inferring from report proposals.
- [x] Every actual consumer of a changed export field is classified; unaffected consumers are preserve-only.
- [x] Child 04 syntax/direct-state ownership is separated from Child 05 terminal-resolution state.
- [x] Exact production, test, fixture, generated-output, and validation ownership is recorded for P4-A.
- [x] P4-A work steps match the discovered current owners.
- [x] Status Refresh Log contains the completed P0 refresh.
- [x] Child 03 predecessor handoff `E3-PNC-HANDOFF1` is recorded without selecting a Child 04 implementation owner.

## Final P0 Decision

- [ ] P0 actual-status incomplete. Implementation is blocked.
- [ ] P0 complete. Next phase can proceed unchanged.
- [x] P0 complete. Next phase status, next action, or work steps must be updated before implementation.
- [ ] P0 complete. Target scope is preserve-only.
- [ ] P0 complete. Implementation is blocked by missing authority or evidence.

Decision note:

The bounded `0/21` baseline remains accepted, Child 03 is closed at `3e25f9ca75c9cb7d59bf228e6cc9aa0b81d738b4`, and current ownership, relationship counts, impacts, syntax gaps, projection/persistence drift, and semantic consumers are recorded. P0-A is complete. Its required plan update narrows P4-A to the four deterministic ScopeIR contract owners; P4-A opens immediately after Git commits the exact five-document boundary.
