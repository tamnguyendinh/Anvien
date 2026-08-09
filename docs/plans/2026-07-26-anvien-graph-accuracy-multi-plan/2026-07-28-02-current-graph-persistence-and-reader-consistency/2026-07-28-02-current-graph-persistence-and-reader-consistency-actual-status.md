# Anvien Current Graph Persistence and Reader Consistency Actual Status

Title: Anvien Current Graph Persistence and Reader Consistency

Date: 2026-07-28

Status: Draft / P0 Incomplete / Dependency Blocked

Companion plan: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-02-current-graph-persistence-and-reader-consistency/2026-07-28-02-current-graph-persistence-and-reader-consistency-plan.md`

Companion evidence: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-02-current-graph-persistence-and-reader-consistency/2026-07-28-02-current-graph-persistence-and-reader-consistency-evidence.md`

Companion benchmark: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-02-current-graph-persistence-and-reader-consistency/2026-07-28-02-current-graph-persistence-and-reader-consistency-benchmark.md`

## Purpose

This file records what is actually known about corrected-field persistence and reader behavior before Child 02 implementation. It deliberately leaves the affected-reader denominator open until P2-A proves it from current source.

Implementation cannot start until Child 01 supplies accepted corrected fields and P0 refreshes every candidate persistence/reader owner at the implementation HEAD.

## Freshness / Refresh Rules

- Refresh the graph before graph-based status work.
- Refresh candidate file counts and exact impacts before changing touch mode from inspect-only to edit.
- After P2-A, replace broad candidate rows with exact affected-field/reader rows.
- After each accepted slice, update only affected classifications and append a refresh-log row.
- Remove obsolete historical matrices rather than keeping them as active status.
- Keep proof in evidence and measurements in benchmark.

## Scope

Target scope:

- accepted Child 01 corrected fields;
- Graph JSON and Ladybug preservation of those fields and records;
- exact readers/query adapters proven affected by P2-A;
- repeated normal analyze on the same repository/artifact path;
- clear failure at affected normal boundaries;
- field parity, dropped-record count, affected-reader count, and handoff.

Out of scope:

- readers unrelated to corrected fields;
- unrelated metadata/negotiation behavior or historical reader denominators;
- artifact-write mechanics not required by corrected-field evidence;
- binding, export, module/re-export, ambient/external, and graph-health implementation;
- scanner repair, target source changes, and unrelated refactors.

## Relationship / Impact Evidence

The rows below are candidate boundaries from the latest production-source audit, not an affected-reader list. P2-A must re-run source/file-detail/impact and may remove every candidate not connected to an accepted corrected field.

| Unit / File / Surface | File Detail Evidence | Related File Count | Relationship Summary | Impact Note |
|-----------------------|----------------------|-------------------:|----------------------|-------------|
| `internal/lbugload/csv.go` | `E0-P0A-FD1` | 19 files | graph-to-Ladybug record projection/load preparation | HIGH candidate persistence boundary |
| `internal/analyze/analyze.go` | `E0-P0A-FD2` | 182 files | normal analyze orchestration and artifact calls | CRITICAL candidate; inspect first, preserve if acceptance already holds |
| `internal/httpapi/graph.go` | `E0-P0A-FD3` | 22 files | graph response candidate | HIGH candidate reader; affected status unproven |
| `anvien-web/src/services/backend-client.ts` | `E0-P0A-FD4` | 24 files | Web graph response candidate | HIGH candidate reader; affected status unproven |

## Status Rules

| Status | Meaning | Allowed next action |
|--------|---------|---------------------|
| `correct` | Current behavior already satisfies the requirement. | Preserve and validate. |
| `partial` | Some required behavior exists but acceptance is incomplete. | Correct only the measured gap. |
| `wrong` | Current behavior conflicts with the required corrected fact. | Change the exact evidence-backed owner. |
| `missing` | Required inventory, behavior, or field is absent. | Add only after evidence identifies the owner. |
| `unbound` | Corrected fact exists but does not reach a real reader. | Bind the proven affected reader. |
| `fake-or-stub` | Placeholder output is used as real data. | Remove or replace truthfully. |
| `blocked` | Predecessor or current evidence is absent. | Do not implement. |

## Current Status Matrix

| Unit | Current State | Required State | Status | Relationship Count | Evidence | Next Plan Decision |
|------|---------------|----------------|--------|-------------------:|----------|--------------------|
| Child 01 handoff | accepted corrected identity/range field set is not yet recorded | exact accepted fields and source evidence | blocked | predecessor boundary | `E0-P0A-HANDOFF1` pending | no P2 production edit |
| Graph JSON preservation | baseline graph serialization exists; corrected Child 01 field round-trip is not yet measured | all accepted corrected records/fields preserved without drop | partial/blocked | P2-A source inventory pending | `E0-P0A-SOURCE1` pending | inventory in P2-A; implement only if P2-B proves gap |
| Ladybug preservation | bounded investigation observed representation changes and current source has a known projection boundary; effect on corrected fields is not yet classified | lossless corrected facts and explicit representation mapping | partial | 19 candidate related files | `E0-P0A-VERIFY1`, `E0-P0A-FD1`; current field impact pending | P2-A then P2-B |
| affected reader inventory | no accepted source-proven affected-reader denominator | exact rows with symbol/path, field, impact, touch mode, evidence | missing | unknown until P2-A | `E2-P2A-INVENTORY1` pending | P2-A only |
| HTTP graph reader candidate | current candidate exists; dependency on corrected fields unproven | edit only if P2-A demonstrates a corrected-field dependency | blocked pending inventory | 22 candidate related files | `E0-P0A-FD3` | inspect in P2-A |
| Web graph reader candidate | current candidate exists; dependency on corrected fields unproven | edit only if P2-A demonstrates a corrected-field dependency | blocked pending inventory | 24 candidate related files | `E0-P0A-FD4` | inspect in P2-A |
| repeated normal analyze | command is repeatedly usable in the product, but Child 02 matched-run/current-fact/failure acceptance is not measured | all declared runs expose matching corrected facts; failures are clear | partial/unknown | 182 candidate related files | `E0-P0A-FD2`, `E0-P0A-SOURCE2` pending | first measure in P2-D; preserve correct mechanics |
| target boundary | target source/worktree remains validation-only | normal target output only; no source contamination | correct | external repository boundary | `E0-P0A-BOUNDARY1` | preserve through validation |

## Status Refresh Log

| Refresh | Date | Repo Basis | Changed Scope | Status Changes | Evidence | Next Phase Update |
|---------|------|------------|---------------|----------------|----------|-------------------|
| R0 | 2026-08-10 | documentation correction worktree; retained production-source candidate counts | Child 02 only | campaign-wide reader assumptions removed; affected inventory missing; predecessor and fresh P0 gates pending | `E0-P0A-REPORT1`, `E0-P0A-VERIFY1`, `E0-P0A-FD1..E0-P0A-FD4`, `E0-P0A-STATUS1` | keep P2-A through P2-E blocked until accepted Child 01 handoff and current P0 evidence |

## Phase Touch Map

| Unit / File / Surface | Plan-Relevant Relationship File | Relationship to Target | Plan Item | Touch Mode | Evidence | Constraint |
|-----------------------|---------------------------------|------------------------|-----------|------------|----------|------------|
| Child 01 accepted fields | Graph JSON/Ladybug source owners | predecessor contract | P0/P2-A | inspect-only | `E0-P0A-HANDOFF1` pending | no schema assumption before handoff |
| candidate Graph JSON/analyze flow | current source discovered in P0 | persistence/orchestration | P2-A/P2-B/P2-D | inspect-only until exact impact | `E0-P0A-SOURCE1`, `E0-P0A-SOURCE2` pending | output behavior remains inspect-only until exact impact |
| `internal/lbugload/csv.go` | current graph records | Ladybug projection candidate | P2-A/P2-B | inspect-only until affected-field proof | `E0-P0A-FD1` | zero silent corrected-record drops |
| `internal/httpapi/graph.go` | Graph JSON/Ladybug query output | public reader candidate | P2-A/P2-C | inspect-only; edit only if affected | `E0-P0A-FD3` | no category-wide reader work |
| `anvien-web/src/services/backend-client.ts` | HTTP graph response | UI reader candidate | P2-A/P2-C | inspect-only; edit only if affected | `E0-P0A-FD4` | real built UI QA only if opened |
| `E:\cheapapp.org` | normal target output | external validation boundary | P2-E if required | preserve source; validate only | `E0-P0A-BOUNDARY1` | no target fixture/report/debug artifacts |

## Detailed Findings

### Persistence and representation

Current state:

The bounded investigation found one-to-one cardinality for selected raw/Cypher/file-detail facts while recording representation changes such as reduced output shape and scalar/null handling. That does not prove corrected Child 01 fields are preserved, nor does it prove a loss in every reader.

Required state:

```text
accepted corrected record
-> Graph JSON record
-> Ladybug record
-> exact source-proven affected reader
```

Each arrow must conserve the corrected semantic fact, and any representation mapping must be explicit.

Evidence:

- `E0-P0A-VERIFY1`: bounded representation observation, not a universal reader defect.
- `E0-P0A-HANDOFF1`: pending accepted corrected-field set.
- `E2-P2A-INVENTORY1`: pending exact field-flow inventory.

Relationship and impact:

- Related file counts: 19 Ladybug candidate files; 22 HTTP candidate files; 24 Web candidate files; 182 analyze candidate files.
- Impact note: candidates remain inspect-only; HIGH/CRITICAL does not authorize broad edits.

Classification: persistence `partial/blocked`; affected-reader inventory `missing`.

Allowed next action: complete predecessor/P0 gates, then execute P2-A source inventory.

Forbidden next action: reuse a historical reader denominator or introduce storage behavior before proving a corrected-field gap.

### Repeated analyze

Current state:

The command is called repeatedly in normal use, but the Child 02 matched-run and failure matrix has not been measured at the implementation HEAD.

Required state:

```text
normal analyze invocation
-> normal artifact path implemented by source
-> affected reader returns matching corrected facts

failure
-> clear non-success at the normal boundary
```

Classification: `partial/unknown`; measure before changing code.

## Next Phase Status Decisions

| Plan Item | Actual Status Finding | Required Status / Next-Action Update |
|-----------|-----------------------|--------------------------------------|
| P2-A | affected denominator missing; predecessor fields pending | remain blocked, then perform source-only inventory first |
| P2-B | persistence gap not yet measured | edit only exact affected owners after P2-A |
| P2-C | reader candidates are not accepted affected readers | exclude every unproven candidate; edit exact P2-A rows only |
| P2-D | repeated behavior unmeasured | test existing built behavior before implementation change |
| P2-E | no accepted parity denominator | validation-only after P2-A through P2-D pass |

## Implementation Gate

- [ ] Accepted Child 01 handoff is recorded.
- [ ] Target scope is current in the status matrix.
- [ ] Every target unit has current evidence and status.
- [ ] Candidate file counts are refreshed at the implementation HEAD.
- [ ] Exact persistence/reader symbols have current impact evidence.
- [ ] P2-A affected-field/reader inventory is complete before P2-B/P2-C.
- [x] Candidate readers remain inspect-only rather than assumed affected.
- [x] Target boundary is preserve-only.
- [ ] Next-slice assumptions and work steps are refreshed from current evidence.
- [x] Status Refresh Log has an R0 correction row.
- [ ] P0 Supervisor review passes.

## Final P0 Decision

- [x] P0 actual-status incomplete. Implementation is blocked.
- [ ] P0 complete. Next phase can proceed unchanged.
- [ ] P0 complete. Next phase status, next action, or work steps must be updated before implementation.
- [ ] P0 complete. Target scope is preserve-only.
- [ ] P0 complete. Implementation is blocked by missing authority or evidence.

Decision note:

Child 01 handoff, current source/file-detail/impact refresh, exact affected-reader inventory, and P0 Supervisor acceptance are pending. No production slice is open.
