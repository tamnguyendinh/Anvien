# Anvien Current Graph Persistence and Reader Consistency Actual Status

Title: Anvien Current Graph Persistence and Reader Consistency

Date: 2026-07-28

Status: Active / P0-A Accepted / P2-A Source Inventory Open

Companion plan: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-02-current-graph-persistence-and-reader-consistency/2026-07-28-02-current-graph-persistence-and-reader-consistency-plan.md`

Companion evidence: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-02-current-graph-persistence-and-reader-consistency/2026-07-28-02-current-graph-persistence-and-reader-consistency-evidence.md`

Companion benchmark: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-02-current-graph-persistence-and-reader-consistency/2026-07-28-02-current-graph-persistence-and-reader-consistency-benchmark.md`

## Purpose

This file records what is actually known about corrected-field persistence and reader behavior before Child 02 implementation. P0-A established the current-codebase baseline, accepted Child 01 dependency, candidate owners, blast radii, preserve-only boundaries, and next-slice decision. It did not establish the exact affected-reader denominator.

Production implementation cannot start until P2-A independently verifies the exact affected persistence/reader inventory and routes every finding to its owning later slice.

## Freshness / Refresh Rules

- Refresh the graph before graph-based status work.
- Refresh candidate file counts and exact impacts before changing touch mode from inspect-only to edit.
- Carry P0-A rows into P2-A as candidate leads only; refresh evidence invalidated by repo drift and independently verify or exclude every row before it becomes an accepted denominator.
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

The rows below are the P0-A current-source/impact baseline at HEAD `22b8ed4cb0d7393cc5c982586b2e9a5c4c049760`. Every P2 touch mode remains a candidate until P2-A verifies ownership and the owning implementation slice opens.

| Unit / File / Surface | File Detail Evidence | Related File Count | Relationship Summary | Impact Note |
|-----------------------|----------------------|-------------------:|----------------------|-------------|
| `internal/lbugload/csv.go` | `E0-P0A-FD1` | 19 files | Ladybug symbol projection candidate | CRITICAL; verify in P2-A before any P2-B ownership |
| `internal/analyze/analyze.go` | `E0-P0A-FD2` | 182 files | generic Graph JSON writer and persistence orchestration | CRITICAL; validate-only |
| `internal/httpapi/graph.go` | `E0-P0A-FD3` | 22 files | transparent graph record transport | MEDIUM; preserve-only |
| `anvien-web/src/services/backend-client.ts` | `E0-P0A-FD4` | 24 files | transparent Web graph record transport | CRITICAL; preserve-only |
| `useAppState.local-runtime.tsx::resolveNodeIds` | `E0-P0A-FD5` | 30 files | opaque-ID interpretation lead | exact symbol LOW; verify/classify in P2-A |
| `internal/resolution/emit.go::emitDefinitionNodes` | `E0-P0A-FD6` | 42 files | corrected-field projection candidate | CRITICAL; verify in P2-A before any P2-B ownership |
| `internal/lbugschema/schema.go::NodeSchema` | `E0-P0A-FD7` | 18 files | Ladybug symbol schema candidate | LOW; verify in P2-A before any P2-B ownership |
| `internal/mcp/tools.go::runCypherRead` | `E0-P0A-FD8` | 55 files | Ladybug-to-Graph-JSON fallback boundary | LOW; future P2-D validate-only |
| `internal/filecontext/context.go::nodeRange` | `E0-P0A-FD9` | 44 files | corrected-range reader lead | CRITICAL; verify/classify in P2-A |

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
| Child 01 handoff | accepted corrected identity/range/occurrence/collision/endpoint field set is recorded by orchestration/main from Child 01's accepted evidence and commit chain | exact accepted fields, source evidence, explicit non-claims, and predecessor HEAD/commit chain | correct | predecessor boundary | `E0-P0A-HANDOFF1`, `E2-PNC-HANDOFF1`, `E2-PNA-SUPERVISOR1`, `E2-PNB-COMMIT1` | accepted input to P2-A; no production edit follows from the handoff alone |
| Graph JSON preservation | generic writer and projection candidates are recorded, but corrected-field round-trip is not yet measured | all accepted corrected records/fields preserved without drop | partial | candidate emission/persistence flow recorded | `E0-P0A-QA1`, `E0-P0A-SOURCE1`, `E0-P0A-FD6` | P2-A verifies ownership; P2-B owns any proven correction |
| Ladybug preservation | projection/schema candidates are recorded; exact corrected-field lossless mapping is not yet accepted | lossless corrected facts and explicit representation mapping | partial | candidate projection/schema owners recorded | `E0-P0A-QA1`, `E0-P0A-SOURCE1`, `E0-P0A-FD1`, `E0-P0A-FD7` | P2-A verifies ownership; P2-B owns any proven correction |
| affected reader inventory | P0-A recorded candidate reader leads; review identified additional direct-consumer leads outside the P0-A acceptance boundary | exact rows with symbol/path, field, impact, touch mode, evidence and zero unassigned consumers | missing / P2-A-owned | unknown until P2-A | `E0-P0A-QA1`, `E2-P2A-INPUT1` | P2-A verifies or excludes every lead and establishes the denominator; this did not block P0-A |
| HTTP graph transport | current source transports graph records without interpreting corrected fields | preserve existing transparent transport | correct | 22 related files | `E0-P0A-QA1`, `E0-P0A-FD3`, `E0-P0A-IMPACT3` | preserve-only |
| Web graph transport | backend client transports graph records without interpreting corrected fields | preserve existing transparent transport | correct | 24 related files | `E0-P0A-QA1`, `E0-P0A-FD4`, `E0-P0A-IMPACT4` | preserve-only |
| repeated normal analyze | command is repeatedly usable in the product, but Child 02 matched-run/current-fact/failure acceptance is not measured | all declared runs expose matching corrected facts; failures are clear | partial/unknown | 182 related files | `E0-P0A-FD2`; P2-D evidence pending | first measure in P2-D; preserve correct mechanics |
| target boundary | target source/worktree remains validation-only | normal target output only; no source contamination | correct | external repository boundary | `E0-P0A-BOUNDARY1` | preserve through validation |

## Status Refresh Log

| Refresh | Date | Repo Basis | Changed Scope | Status Changes | Evidence | Next Phase Update |
|---------|------|------------|---------------|----------------|----------|-------------------|
| R0 | 2026-08-10 | documentation correction worktree; retained production-source candidate counts | Child 02 only | campaign-wide reader assumptions removed; affected inventory missing; predecessor and fresh P0 gates pending | `E0-P0A-REPORT1`, `E0-P0A-VERIFY1`, `E0-P0A-FD1..E0-P0A-FD4`, `E0-P0A-STATUS1` | keep P2-A through P2-E blocked until accepted Child 01 handoff and current P0 evidence |
| R1 | 2026-08-11 | Child 01 accepted HEAD chain through Pn-B `da49506a71e006b9ab48137b780e185bf14582fb`; fresh Anvien graph `1,580/680/0`, `95,819/134,750` | Child 01 handoff authority and P0 predecessor row | `Child 01 handoff missing/blocked -> recorded/correct`; P0 remains incomplete because current Child 02 source/file-detail/impact and Supervisor evidence are still pending | `E0-P0A-HANDOFF1`, `E2-PNC-HANDOFF1`, `E2-PNA-SUPERVISOR1`, `E2-PNB-COMMIT1` | run Child 02 P0-A source inventory and current graph/file-detail/impact; do not open P2-B/P2-C/P2-D/P2-E or edit production from the handoff alone |
| R2 | 2026-08-11 | Child 01 Pn-C isolated closure commit `22b8ed4cb0d7393cc5c982586b2e9a5c4c049760`; main opened the next slice from the accepted handoff | Child 02 P0-A execution boundary | P0-A `blocked-before-open -> open`; predecessor facts remain correct; no production touch mode changes before source/file-detail/impact inventory | `E0-P0A-HANDOFF1`, `E2-PNC-COMMIT1` | run only P0-A source inventory; record exact affected/unaffected rows and obtain P0-A Supervisor decision before opening P2-A |
| R3 | 2026-08-13 | same accepted HEAD; fresh graph `1,581/680/0`, `95,830/134,761`; QA no-fix report imported into main workspace | bounded persistence/reader inventory | inventory `missing -> recorded/pending acceptance`; affected readers `unknown -> 2`, editable subset `1`; HTTP/Web transports `candidate -> preserve-only`; exact P2-B/P2-C/P2-D owners recorded | `E0-P0A-QA1`, `E0-P0A-GRAPH1`, `E0-P0A-SOURCE1`, `E0-P0A-SOURCE2`, `E0-P0A-FD1..E0-P0A-FD9`, `E0-P0A-IMPACT1..E0-P0A-IMPACT9` | open only the P0-A Supervisor gate; P0/P2 remain closed until `E0-P0A-REVIEW1` PASS |
| R4 | 2026-08-13 | P0-A QA and zero-trust review completed at the same production HEAD; fresh review graph `1,583/680/0`, `95,860/134,791`; long impact batch completed in `162s` | P0 scope/dependency decision and out-of-slice finding routing | P0-A current-state/dependency/source/impact/no-code-diff gates accepted; the review's direct-consumer completeness finding is transferred without loss to `E2-P2A-INPUT1` because denominator closure belongs to P2-A | `E0-P0A-REVIEW1`, `E2-P2A-INPUT1` | close P0-A, open only P2-A, do not rerun P0 QA/review |

## Phase Touch Map

| Unit / File / Surface | Plan-Relevant Relationship File | Relationship to Target | Plan Item | Touch Mode | Evidence | Constraint |
|-----------------------|---------------------------------|------------------------|-----------|------------|----------|------------|
| Child 01 accepted fields | Graph JSON/Ladybug source owners | predecessor contract | P0/P2-A | inspect-only | `E0-P0A-HANDOFF1`, `E2-PNC-HANDOFF1` | accepted facts are inputs to P2-A inventory; no persistence schema or affected-reader assumption |
| `internal/resolution/emit.go::emitDefinitionNodes` | accepted Definition facts | graph property projection | P2-A/P2-B | P2-A candidate; production locked | `E0-P0A-FD6`, `E0-P0A-IMPACT6` | verify ownership in P2-A; if affected, route only accepted range/selection projection to P2-B |
| `internal/lbugload/csv.go` and `internal/lbugschema/schema.go::NodeSchema` | current graph records | Ladybug symbol projection/schema | P2-A/P2-B | P2-A candidates; production locked | `E0-P0A-FD1`, `E0-P0A-FD7` | verify ownership in P2-A; if affected, route zero-loss correction to P2-B |
| `internal/analyze/analyze.go` | graph records | generic Graph JSON writer/orchestration | P2-A/P2-B/P2-D | validate-only | `E0-P0A-FD2`, `E0-P0A-IMPACT2` | preserve generic writer mechanics |
| `internal/httpapi/graph.go` and `backend-client.ts` | Graph JSON records | transparent transport | P2-A/P2-C | preserve-only | `E0-P0A-FD3`, `E0-P0A-FD4` | excluded from affected-reader denominator |
| `useAppState.local-runtime.tsx::resolveNodeIds` | opaque node ID | semantic-reader lead | P2-A/P2-C | P2-A candidate; production locked | `E0-P0A-FD5`, `E0-P0A-IMPACT5`, `E2-P2A-INPUT1` | verify together with sibling Web grounding; route any proven reader correction to P2-C |
| `internal/filecontext/context.go::nodeRange` | construct range fields | semantic-reader lead | P2-A/P2-C | P2-A candidate; production locked | `E0-P0A-FD9`, `E2-P2A-INPUT1` | verify together with other direct filecontext consumers; route any proven reader correction to P2-C |
| `internal/mcp/tools.go::runCypherRead` | Ladybug/Graph JSON query boundary | repeated-read fallback | P2-A/P2-D | validate-only | `E0-P0A-FD8`, `E0-P0A-IMPACT8` | excluded from P2-C denominator |
| `E:\cheapapp.org` | normal target output | external validation boundary | P2-E if required | preserve source; validate only | `E0-P0A-BOUNDARY1` | no target fixture/report/debug artifacts |

## Detailed Findings

### Accepted Child 01 handoff

Current state:

The predecessor handoff is durable and source-backed. Child 01 accepted the identity tuple, construct/selection position facts, production `defsByFile` occurrence denominator, canonical Definition conflict behavior, and endpoint integrity described in `E0-P0A-HANDOFF1`. The full predecessor chain through Pn-C is accepted at `22b8ed4cb0d7393cc5c982586b2e9a5c4c049760`.

Required state:

```text
accepted Child 01 facts
-> current Child 02 source inventory
-> exact Graph JSON/Ladybug/reader parity decision
```

Evidence:

- `E0-P0A-HANDOFF1`: exact accepted facts, source evidence, non-claims, and predecessor chain.
- `E2-PNC-HANDOFF1`: parent Pn-C handoff record.

Relationship and impact:

- Related file count: predecessor boundary; Child 02's exact current rows are recorded in `E0-P0A-QA1`.
- Impact note: the handoff and QA inventory are evidence inputs only; neither authorizes production edits before the owning P2 slice.

Classification:

`correct` for the predecessor contract; Child 02 P0-A is accepted.

Allowed next action:

Execute only the exact P2-A documentation/source-audit slice. Treat P0 candidate rows and review findings as leads that must be verified or excluded; do not edit production.

Forbidden next action:

Do not infer Graph JSON/Ladybug parity, affected-reader membership, repeated-analyze reader behavior, or column/SelectionRange projection from the Child 01 handoff; do not edit production before P2-A proves an owner.

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
- `E0-P0A-HANDOFF1`: recorded accepted corrected-field set; Graph JSON/Ladybug/reader parity remains unmeasured and is still owned by P2-A/P2-B/P2-C.
- `E0-P0A-QA1`: current-codebase field-flow/candidate inventory recorded and accepted for the bounded P0-A purpose; later P2-A evidence mapping remains independent.

Relationship and impact:

- Exact related-file counts and impacts are recorded in `E0-P0A-FD1..E0-P0A-FD9` and `E0-P0A-IMPACT1..E0-P0A-IMPACT9`.
- Impact note: HIGH/CRITICAL warnings keep future edits bounded and do not authorize broad changes.

Classification: P0 baseline `recorded/accepted`; persistence and reader rows remain unverified P2-A candidates, with final denominator unresolved.

Allowed next action: execute P2-A only. Verify every candidate and additional direct-consumer lead, establish the exact denominator, and route any discovered gap to P2-B/P2-C/P2-D without implementing it in P2-A.

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
| P2-A | P0 baseline accepted; candidate denominator and direct-consumer leads are unverified | open source-only inventory; verify/exclude every lead and establish exact affected rows |
| P2-B | candidate persistence owners recorded | remain closed; edit only owners proven and assigned by accepted P2-A plus fresh pre-edit impact |
| P2-C | candidate reader leads recorded; final denominator/editable subset unresolved | remain closed; consume only accepted P2-A reader rows after P2-B and slice gates |
| P2-D | repeated behavior unmeasured | test existing built behavior before implementation change |
| P2-E | no accepted parity denominator | validation-only after P2-A through P2-D pass |

## Implementation Gate

- [x] Accepted Child 01 handoff is recorded.
- [x] Target scope is current in the status matrix.
- [x] Every P0 target unit has current baseline evidence and status; P2-A owns completeness.
- [x] Candidate file counts are refreshed at the implementation HEAD.
- [x] Exact persistence/reader symbols have current impact evidence.
- [ ] P2-A affected-field/reader inventory is complete before P2-B/P2-C (P2-A acceptance gate, not P0-A).
- [x] Future edit candidates remain locked rather than authorized.
- [x] Target boundary is preserve-only.
- [x] Next-slice assumptions and work steps are refreshed from current evidence.
- [x] Status Refresh Log has an R0 correction row.
- [x] P0 QA and zero-trust review gates completed; Owner corrected the out-of-slice denominator interpretation.
- [x] P0 scope/dependency disposition passes without rerunning accepted gates.

## Final P0 Decision

- [ ] P0 actual-status incomplete. Implementation is blocked.
- [ ] P0 complete. Next phase can proceed unchanged.
- [x] P0 complete. Next phase status, next action, or work steps must be updated before implementation.
- [ ] P0 complete. Target scope is preserve-only.
- [ ] P0 complete. Implementation is blocked by missing authority or evidence.

Decision note:

The accepted Child 01 handoff, current codebase baseline, candidate files/owners, impact warnings, preserve-only boundaries, and no-production-diff gate are recorded. P0-A QA and zero-trust review completed. Owner corrected the review's phase-boundary error: exact direct-consumer completeness belongs to P2-A, so it is preserved as `E2-P2A-INPUT1` rather than treated as a P0-A rejection. P0-A is complete; P2-A is the only open slice, and all production slices remain closed.
