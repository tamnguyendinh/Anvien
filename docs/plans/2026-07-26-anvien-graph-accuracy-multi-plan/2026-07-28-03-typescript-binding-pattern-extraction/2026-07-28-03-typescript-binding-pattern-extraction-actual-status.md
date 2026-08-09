# Anvien TypeScript Binding-Pattern Extraction Actual Status

Title: Anvien TypeScript Binding-Pattern Extraction
Date: 2026-07-28
Status: Draft / P0 incomplete
Companion plan: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-03-typescript-binding-pattern-extraction/2026-07-28-03-typescript-binding-pattern-extraction-plan.md`
Companion evidence: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-03-typescript-binding-pattern-extraction/2026-07-28-03-typescript-binding-pattern-extraction-evidence.md`
Companion benchmark: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-03-typescript-binding-pattern-extraction/2026-07-28-03-typescript-binding-pattern-extraction-benchmark.md`
Roadmap: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-26-anvien-graph-accuracy-roadmap.md`
Predecessor child: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-02-current-graph-persistence-and-reader-consistency/2026-07-28-02-current-graph-persistence-and-reader-consistency-plan.md`
Successor child: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-04-typescript-export-semantics/2026-07-28-04-typescript-export-semantics-plan.md`

## Purpose

This file classifies the current binding-pattern scope before implementation. Historical investigation proves a bounded defect but does not substitute for fresh current-source, graph, file-detail, and impact evidence.

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

The bounded investigation identifies candidate source locations, but current related-file counts and blast radius must be refreshed before editing.

| Unit / File / Surface | Current Evidence | Related File Count | Relationship Summary | Impact Note |
|-----------------------|------------------|-------------------:|----------------------|-------------|
| Current TS variable-declaration extraction owner | `E0-P0A-VERIFY1` bounded source finding; `E0-P0A-FD1` pending | pending fresh P0 | AST declaration name to current fact emission | edit blocked until fresh file-detail/impact |
| Current binding fact/index owner | `E0-P0A-SRC1` pending | pending fresh P0 | binding leaf, scope, range, path, and index contract | owner not inferred from a filename |
| Current graph/reference projection owner | `E0-P0A-SRC1` pending | pending fresh P0 | extracted binding to graph occurrence and reference target | likely broad blast radius; exact impact pending |
| Existing persistence/reader surfaces | `E0-P0A-SRC1` pending | pending fresh P0 | include only surfaces that consume a changed binding contract | inspect-only until proven affected |
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
| Current production owner inventory | Candidate source mechanism is known, but current HEAD owners and impacts are not yet fully revalidated | exact editable/inspect-only/preserve-only owners with current counts | `blocked` | pending | `E0-P0A-SRC1`, `E0-P0A-FD1`, `E0-P0A-IMPACT1` pending | complete P0 before P3-A |
| Identifier-only declarators | Accepted investigation observed existing plain-identifier emission | preserve current correct identifier behavior | `partial` | pending | `E0-P0A-VERIFY1` | revalidate and mark preserve-only in P0 |
| Bounded array-declaration bindings | Six legal names existed in AST and produced no definition/local binding at the accepted baseline | six independent declarations and bindings | `wrong` | pending | `E0-P0A-VERIFY1`, `E0-P0A-VERIFY2` | P3-A then P3-B |
| General object/array/rest/default/hole semantics | Proposed in the DRAFT report but not accepted as current-source coverage | general legal binding matrix with typed paths | `blocked` | pending | `E0-P0A-ORIGIN1` | P0 inventories current support; P3-A fills proven gaps |
| Parameter/catch/loop contexts | Required by the plan but not measured by the bounded investigation | accepted leaf semantics in each lexical context | `blocked` | pending | `E0-P0A-SRC1` pending | P3-B1, P3-B2, P3-B2A after P0 |
| Unsupported-pattern diagnostics | Historical failure was a silent early return; current diagnostic coverage is unmeasured | structured, countable diagnostic for every unsupported fixture | `blocked` | pending | `E0-P0A-VERIFY1`, current evidence pending | P3-A after source confirmation |
| Binding graph/reference outcome | Six downstream uses became gaps because no binding fact existed | one graph occurrence per leaf and correct lexical reference target | `wrong` | pending | `E0-P0A-VERIFY1` | P3-C after all extraction contexts |
| Assignment/import/shadowing controls | Not established by the bounded investigation | zero false declarations, zero import double count, correct lexical symbol | `blocked` | pending | reserved P3 evidence | measure in P3-A/P3-B/P3-C |
| Target boundary | Target analyzed in place; investigation artifacts remained in Anvien | preserve target source and keep normal operational output target-local | `correct` | N/A | `E0-P0A-VERIFY1`, `E0-P0A-VERIFY2` | preserve in P3-C2 |

## Status Refresh Log

| Refresh | Date | Repo Basis | Changed Scope | Status Changes | Evidence | Next Phase Update |
|---------|------|------------|---------------|----------------|----------|-------------------|
| R0 | 2026-08-10 | documentation correction against full problem-origin and bounded-verification reports | Child 03 plan authority and scope | removed unrelated campaign assumptions; P0 reset to incomplete because current owner/impact proof is pending | `E0-P0A-RULE1`, `E0-P0A-SKILL1`, `E0-P0A-ORIGIN1`, `E0-P0A-VERIFY1`, `E0-P0A-VERIFY2` | run fresh P0; do not open P3-A yet |

## Phase Touch Map

| Unit / File / Surface | Plan-Relevant Relationship | Plan Item | Touch Mode | Evidence | Constraint |
|-----------------------|----------------------------|-----------|------------|----------|------------|
| current TS binding walker/fact owner | source of legal binding leaves | P3-A | block until P0, then edit if proven | `E0-P0A-SRC1`, `E0-P0A-FD1`, `E0-P0A-IMPACT1` pending | one semantic responsibility |
| current variable context owner | invokes accepted leaf extraction | P3-B | inspect, then edit if proven | pending P0 | preserve identifiers, imports, assignments |
| current parameter owner | parameter lexical context | P3-B1 | inspect, then edit if proven | pending P0 | no variable/catch/loop logic |
| current catch owner | catch lexical context | P3-B2 | inspect, then edit if proven | pending P0 | no exception-flow redesign |
| current loop owner | loop declaration context | P3-B2A | inspect, then edit if proven | pending P0 | assignment-form loops remain writes |
| current graph/reference owners | occurrence projection and lexical target | P3-C | inspect, then narrow edit | pending P0/P3 impact | no name-based rescue; only named affected consumers |
| unaffected readers/transports/caches | possible downstream consumers | all P3 | inspect-only unless exact changed-field consumption is proven | pending impact | plan update required before expansion |
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
| P3-A | defect is verified, but current fact/walker ownership and general support matrix are pending | keep blocked until P0 current source/file-detail/impact are complete |
| P3-B | bounded variable omission is wrong | open only after P3-A; require type-inference-failure and assignment/import controls |
| P3-B1 | current parameter support unmeasured | inventory in P0; implement only missing behavior |
| P3-B2 | current catch support unmeasured | inventory in P0; implement only missing behavior |
| P3-B2A | current loop support unmeasured | inventory in P0; distinguish declaration and assignment forms |
| P3-C | downstream graph/reference outcome is wrong for the six sites | open only after all context slices; touch only proven graph/persistence owners |
| P3-C2 | target boundary is known and preserve-only | run only after P3-C commit with independent six-site oracle |

## Implementation Gate

- [ ] Current graph and HEAD basis are recorded.
- [ ] Current production source and existing tests for every binding context are read in full.
- [ ] Every candidate editable file has fresh `file-detail` related-file count evidence.
- [ ] Every candidate function/method/exported symbol has fresh upstream impact evidence and reported blast radius.
- [ ] Current Status Matrix rows are refreshed from current source, not historical filenames alone.
- [ ] Correct identifier/import/assignment behavior is marked preserve-only.
- [ ] Exact production, test, fixture, generated-output, and validation ownership is recorded for P3-A.
- [ ] P3-A work steps match the discovered current owners.
- [ ] Status Refresh Log contains the completed P0 refresh.

## Final P0 Decision

- [x] P0 actual-status incomplete. Implementation is blocked.
- [ ] P0 complete. Next phase can proceed unchanged.
- [ ] P0 complete. Next phase status, next action, or work steps must be updated before implementation.
- [ ] P0 complete. Target scope is preserve-only.
- [ ] P0 complete. Implementation is blocked by missing authority or evidence.

Decision note:

The bounded defect and target acceptance numbers are verified. Current-source ownership, relationship counts, impact, and general-context support are not yet refreshed, so P3-A cannot open.
