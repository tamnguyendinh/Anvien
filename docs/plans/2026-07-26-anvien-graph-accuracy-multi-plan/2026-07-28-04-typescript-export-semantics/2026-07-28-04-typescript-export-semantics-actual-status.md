# Anvien First-Class TypeScript Export Semantics Actual Status

Title: Anvien First-Class TypeScript Export Semantics
Date: 2026-07-28
Status: Draft / P0 incomplete / Child 03 predecessor handoff recorded / opens only after Pn-C commit / implementation remains blocked
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

The bounded investigation identifies candidate source locations, but current related-file counts and blast radius must be refreshed before editing.

| Unit / File / Surface | Current Evidence | Related File Count | Relationship Summary | Impact Note |
|-----------------------|------------------|-------------------:|----------------------|-------------|
| Current TS export extraction owner | `E0-P0A-VERIFY1` bounded source finding; `E0-P0A-FD1` pending | pending fresh P0 | export AST to current provider fact | edit blocked until fresh file-detail/impact |
| Current fact/meaning owner | `E0-P0A-SRC1` pending | pending fresh P0 | module export kind/name/range/meaning contract | owner not inferred from a proposed filename |
| Current graph export projection owner | `E0-P0A-VERIFY1`; current detail pending | pending fresh P0 | provider export fact to graph fields/edges | exact impact pending |
| Current persistence mapping | accepted source observation of representation drift; current detail pending | pending fresh P0 | graph export fields to stored/query records | edit only if changed fields are consumed |
| Existing CLI/MCP/HTTP/Web/cache/derived surfaces | `E0-P0A-SRC1` pending | pending fresh P0 | possible consumers of export fields | inspect-only unless exact consumption is proven |
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
| Child 03 predecessor handoff | Child 03 binding-pattern invariant is accepted and closed at its isolated Pn-C boundary; non-claims explicitly exclude export semantics | consume only accepted predecessor facts and keep Child 04 P0 current-source gates open | `correct` | `1` inbound handoff/roadmap relationship | `E3-PNC-HANDOFF1`, `E3-PNB-COMMIT1`, `E3-PNA-REVIEW1` | open only Child 04 P0-A; do not infer export owners or implementation details |
| Current production owner inventory | Candidate provider/emitter mechanism is known, but current owners and impacts are not yet fully revalidated | exact editable/inspect-only/preserve-only owners with current counts | `blocked` | pending | `E0-P0A-SRC1`, `E0-P0A-FD1`, `E0-P0A-IMPACT1` pending | complete P0 before P4-A |
| Bounded direct-export definitions | 21 selected definitions were present at the accepted baseline | preserve the definitions and attach truthful direct export facts | `partial` | pending | `E0-P0A-VERIFY1`, `E0-P0A-VERIFY2` | preserve definition extraction; repair export facts |
| Bounded export metadata | 0 of 21 selected definitions carried export/visibility metadata at the accepted baseline | 21 of 21 correct direct export facts and projections | `wrong` | pending | `E0-P0A-VERIFY1`, `E0-P0A-VERIFY2` | P4-A, P4-B, P4-C |
| Access visibility versus module export | Historical pipeline used/omitted a visibility-like field instead of a first-class verified module-export fact | independent concepts with any compatibility field derived from one export fact | `wrong` | pending | `E0-P0A-ORIGIN1`, `E0-P0A-VERIFY1` | establish boundary in P4-A |
| General direct/default/alias/type-only syntax | Proposed by the DRAFT report but not measured as complete | one truthful fact per supported exported binding/specifier | `blocked` | pending | `E0-P0A-ORIGIN1`, current source pending | inventory in P0; implement gaps in P4-B |
| Star/namespace/re-export syntax | Proposed by the DRAFT report; current syntax-fact coverage unmeasured | immutable syntax facts without terminal resolution | `blocked` | pending | `E0-P0A-ORIGIN1`, current source pending | inventory in P0; implement gaps in P4-B1 |
| Export graph projection | Historical emitter omitted empty provider metadata, so bounded exports appeared unexported | corrected fact kind/names/meaning/provenance and access separation | `wrong` | pending | `E0-P0A-VERIFY1` | P4-C after extraction |
| Persistence field parity | Source evidence identified differing export property names across graph and Ladybug paths | same affected export facts in every proven persistence owner | `blocked` | pending | current `E0-P0A-SRC1` required | P0 determines whether P4-C edits persistence |
| Terminal module/re-export resolution | Separate bounded barrel defect exists | Child 05 resolves syntax facts to terminal outcomes | `wrong` but out of scope | pending | `E0-P0A-VERIFY1` | preserve boundary; no Child 05 logic here |
| Target boundary | Target analyzed in place; investigation artifacts remained in Anvien | preserve target source and keep normal operational output target-local | `correct` | N/A | `E0-P0A-VERIFY1`, `E0-P0A-VERIFY2` | preserve in P4-C2 |

## Status Refresh Log

| Refresh | Date | Repo Basis | Changed Scope | Status Changes | Evidence | Next Phase Update |
|---------|------|------------|---------------|----------------|----------|-------------------|
| R0 | 2026-08-10 | documentation correction against full problem-origin and bounded-verification reports | Child 04 plan authority and scope | removed unrelated campaign assumptions; separated syntax/direct projection from Child 05 resolution; P0 reset to incomplete pending current owner/impact proof | `E0-P0A-RULE1`, `E0-P0A-SKILL1`, `E0-P0A-ORIGIN1`, `E0-P0A-VERIFY1`, `E0-P0A-VERIFY2` | run fresh P0; do not open P4-A yet |
| R1 | 2026-08-20 | Child 03 Pn-B commit `0231d7f9cc7cfd05d8d5da4787161a8ec2cd1550`; excluded graph `1,126/626/0`, `80,908/120,167`; roadmap/Child 03/Child 04 actual/handoff governance rows LOW/non-stale with zero upstream affected files/processes/flows/tests | predecessor handoff and successor opening condition | Child 03 predecessor `dependency-blocked -> accepted handoff`; Child 04 P0 remains incomplete, current owner/file-detail/impact/syntax/consumer evidence remains pending | `E3-PNC-HANDOFF1`, `E3-PNB-COMMIT1` | after Pn-C docs-only commit, open only Child 04 P0-A; no P4 implementation or redundant Supervisor loop |

## Phase Touch Map

| Unit / File / Surface | Plan-Relevant Relationship | Plan Item | Touch Mode | Evidence | Constraint |
|-----------------------|----------------------------|-----------|------------|----------|------------|
| current export fact/meaning owner | source of module-export syntax truth | P4-A | block until P0, then edit if proven | `E0-P0A-SRC1`, `E0-P0A-FD1`, `E0-P0A-IMPACT1` pending | access visibility remains independent |
| current direct export extractor | produces direct/default/alias/type-only facts | P4-B | inspect, then narrow edit | pending P0 | preserve definitions and non-export behavior |
| current re-export syntax extractor | produces named/star/namespace syntax facts | P4-B1 | inspect, then narrow edit | pending P0 | no terminal resolution |
| current graph/persistence owners | project exact export facts | P4-C | inspect, then edit only proven consumers | pending P0/P4 impact | one source fact; only named affected consumers |
| unaffected readers/transports/caches | possible downstream consumers | all P4 | inspect-only unless changed-field consumption is proven | pending impact | plan update required before expansion |
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

Classification: `wrong` for the bounded baseline; current ownership remains `blocked` until fresh P0.

Allowed next action: revalidate the current provider/fact/projection path and actual consumers, then implement P4-A in the exact owner.

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
| Predecessor handoff | Child 03 final evidence/benchmark/target boundary is accepted; no export claim is transferred | `correct`; consume only as dependency input and keep Child 04 P0 gates pending |
| P4-A | bounded omission verified; current fact/meaning owner and compatibility consumers pending | keep blocked until P0 current source/file-detail/impact complete |
| P4-B | 21 direct definitions exist but metadata is wrong | preserve definition extraction; implement exact missing syntax facts after P4-A |
| P4-B1 | current re-export syntax coverage unmeasured | inventory in P0; emit syntax only, no terminal state |
| P4-C | bounded graph projection wrong; affected persistence surface unknown | open after extraction; touch only proven graph/persistence owners |
| P4-C2 | target boundary known and preserve-only | run after P4-C commit with independent 21-entry oracle and negative controls |

## Implementation Gate

- [ ] Current graph and HEAD basis are recorded.
- [ ] Current production source and existing tests for export facts/extraction/projection/persistence are read in full.
- [ ] Every candidate editable file has fresh `file-detail` related-file count evidence.
- [ ] Every candidate function/method/exported symbol has fresh upstream impact evidence and reported blast radius.
- [ ] Current syntax coverage and unsupported paths are inventoried without inferring from report proposals.
- [ ] Every actual consumer of a changed export field is classified; unaffected consumers are preserve-only.
- [ ] Child 04 syntax/direct-state ownership is separated from Child 05 terminal-resolution state.
- [ ] Exact production, test, fixture, generated-output, and validation ownership is recorded for P4-A.
- [ ] P4-A work steps match the discovered current owners.
- [ ] Status Refresh Log contains the completed P0 refresh.
- [x] Child 03 predecessor handoff `E3-PNC-HANDOFF1` is recorded without selecting a Child 04 implementation owner.

## Final P0 Decision

- [x] P0 actual-status incomplete. Implementation is blocked.
- [ ] P0 complete. Next phase can proceed unchanged.
- [ ] P0 complete. Next phase status, next action, or work steps must be updated before implementation.
- [ ] P0 complete. Target scope is preserve-only.
- [ ] P0 complete. Implementation is blocked by missing authority or evidence.

Decision note:

The bounded 21-export defect and target acceptance number are verified. Child 03 is closed and its accepted non-export handoff is recorded, but current-source ownership, relationship counts, impact, syntax coverage, and actual affected persistence consumers are not yet refreshed, so Child 04 P0 remains incomplete and P4-A cannot open.
