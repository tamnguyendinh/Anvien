# Anvien Current Graph Persistence and Reader Consistency Actual Status

Title: Anvien Current Graph Persistence and Reader Consistency

Date: 2026-07-28

Status: Active / P0-A, P2-A, P2-B, and P2-C Accepted at Isolated Boundaries / P2-D Unconditional Supervisor PASS / Final Detect and Isolated Commit Immediately Follow

Companion plan: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-02-current-graph-persistence-and-reader-consistency/2026-07-28-02-current-graph-persistence-and-reader-consistency-plan.md`

Companion evidence: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-02-current-graph-persistence-and-reader-consistency/2026-07-28-02-current-graph-persistence-and-reader-consistency-evidence.md`

Companion benchmark: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-02-current-graph-persistence-and-reader-consistency/2026-07-28-02-current-graph-persistence-and-reader-consistency-benchmark.md`

## Purpose

This file records the living current state of corrected-field persistence and reader behavior. P0-A established the current-codebase baseline and accepted Child 01 dependency; P2-A froze the corrected 19-row inventory; P2-B preserved the corrected facts; P2-C corrected all eight source-proven readers; and P2-D now has unconditional Supervisor PASS for repeated normal analyze and clear failure behavior.

P2-D requires only the main-owned final detect and immediate isolated commit boundary. P2-E remains closed until Git confirms that commit; no accepted P2-D production or existing product-test edit exists.

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
| Graph JSON preservation | C01 now emits construct columns and optional all-or-none SelectionRange; unchanged C04 generic writer preserves them | all accepted corrected records/fields preserved without drop | correct / accepted at P2-B | `36,330/36,330` Definition records; `0` partial selection; `0` missing endpoints | `E2-P2B-SOURCE1`, `E2-P2B-PARITY1`, `E2-P2B-REVIEW1` | preserve in P2-C; final independent parity remains P2-E-owned |
| Ladybug preservation | C02/C03 now align all Definition CSV/schema branches; C05/C06 persist explicit embedding label; C07/C08 remain endpoint-aligned | lossless corrected facts and explicit representation mapping | correct / accepted at P2-B | `36,330/36,330` matched records and endpoint pairs; `0` missing/extra/mismatch/drop | `E2-P2B-SOURCE1`, `E2-P2B-TEST1`, `E2-P2B-PARITY1`, `E2-P2B-REVIEW1` | preserve corrected persistence; after detect/commit, P2-C consumes it only in accepted reader rows |
| affected reader inventory | corrected matrix retains eight source-proven first interpreters and reclassifies semantic-search hydration as future edit; two standalone audit/probe rows are explicitly out of campaign | exact rows with symbol/path, field, touch mode, evidence and zero duplicate/unassigned/unclassified consumers | correct / accepted | `8` affected readers: `4` future edit + `4` validate-only | `E2-P2A-INVENTORY2`, `E2-P2A-MATRIX2`, `E2-P2A-REVIEW1` | P2-C stays closed; P2-B first implements the accepted persistence owners |
| affected reader behavior | C09 exact-ID and C10 unique grounding fail closed; C11 coordinates are consistent; C16 consumes explicit persisted label; C12-C15 remain aligned and unchanged | all eight readers expose the same corrected semantic facts as P2-B without opaque-ID reconstruction or silent ambiguity | correct / accepted at P2-C | `8/8` readers; `4` corrected + `4` validate-only | `E2-P2C-SOURCE1`, `E2-P2C-TEST1`, `E2-P2C-RUNTIME1`, `E2-P2C-REVIEW1` | preserve in P2-D; open repeated-analyze measurement only after P2-C detect/commit succeeds |
| HTTP graph transport | current source transports graph records without interpreting corrected fields | preserve existing transparent transport | correct | 22 related files | `E0-P0A-QA1`, `E0-P0A-FD3`, `E0-P0A-IMPACT3` | preserve-only |
| Web graph transport | backend client transports graph records without interpreting corrected fields | preserve existing transparent transport | correct | 24 related files | `E0-P0A-QA1`, `E0-P0A-FD4`, `E0-P0A-IMPACT4` | preserve-only |
| repeated normal analyze | normal built analyze and `Server.runCypherRead` preserve current accepted facts, expose changed input, fail clearly at owned faults, and recover without substitute success | all declared runs expose matching corrected facts; failures are clear | correct / accepted at P2-D | `2` P2-D boundaries; `7/7` matrix; `4/4` successful analyze runs; `0/2` owned failure attempts returned substitute success | `E2-P2D-IMPACT1`, `E2-P2D-SOURCE1`, `E2-P2D-BUILD1`, `E2-P2D-TEST1`, `E2-P2D-REPEAT1`, `E2-P2D-REVIEW1` | preserve accepted mechanics; complete final detect/isolated commit before opening validation-only P2-E |
| target boundary | target source/worktree remains validation-only | normal target output only; no source contamination | correct | external repository boundary | `E0-P0A-BOUNDARY1` | preserve through validation |

## Status Refresh Log

| Refresh | Date | Repo Basis | Changed Scope | Status Changes | Evidence | Next Phase Update |
|---------|------|------------|---------------|----------------|----------|-------------------|
| R0 | 2026-08-10 | documentation correction worktree; retained production-source candidate counts | Child 02 only | campaign-wide reader assumptions removed; affected inventory missing; predecessor and fresh P0 gates pending | `E0-P0A-REPORT1`, `E0-P0A-VERIFY1`, `E0-P0A-FD1..E0-P0A-FD4`, `E0-P0A-STATUS1` | keep P2-A through P2-E blocked until accepted Child 01 handoff and current P0 evidence |
| R1 | 2026-08-11 | Child 01 accepted HEAD chain through Pn-B `da49506a71e006b9ab48137b780e185bf14582fb`; fresh Anvien graph `1,580/680/0`, `95,819/134,750` | Child 01 handoff authority and P0 predecessor row | `Child 01 handoff missing/blocked -> recorded/correct`; P0 remains incomplete because current Child 02 source/file-detail/impact and Supervisor evidence are still pending | `E0-P0A-HANDOFF1`, `E2-PNC-HANDOFF1`, `E2-PNA-SUPERVISOR1`, `E2-PNB-COMMIT1` | run Child 02 P0-A source inventory and current graph/file-detail/impact; do not open P2-B/P2-C/P2-D/P2-E or edit production from the handoff alone |
| R2 | 2026-08-11 | Child 01 Pn-C isolated closure commit `22b8ed4cb0d7393cc5c982586b2e9a5c4c049760`; main opened the next slice from the accepted handoff | Child 02 P0-A execution boundary | P0-A `blocked-before-open -> open`; predecessor facts remain correct; no production touch mode changes before source/file-detail/impact inventory | `E0-P0A-HANDOFF1`, `E2-PNC-COMMIT1` | run only P0-A source inventory; record exact affected/unaffected rows and obtain P0-A Supervisor decision before opening P2-A |
| R3 | 2026-08-13 | same accepted HEAD; fresh graph `1,581/680/0`, `95,830/134,761`; QA no-fix report imported into main workspace | bounded persistence/reader inventory | inventory `missing -> recorded/pending acceptance`; affected readers `unknown -> 2`, editable subset `1`; HTTP/Web transports `candidate -> preserve-only`; exact P2-B/P2-C/P2-D owners recorded | `E0-P0A-QA1`, `E0-P0A-GRAPH1`, `E0-P0A-SOURCE1`, `E0-P0A-SOURCE2`, `E0-P0A-FD1..E0-P0A-FD9`, `E0-P0A-IMPACT1..E0-P0A-IMPACT9` | open only the P0-A Supervisor gate; P0/P2 remain closed until `E0-P0A-REVIEW1` PASS |
| R4 | 2026-08-13 | P0-A QA and zero-trust review completed at the same production HEAD; fresh review graph `1,583/680/0`, `95,860/134,791`; long impact batch completed in `162s` | P0 scope/dependency decision and out-of-slice finding routing | P0-A current-state/dependency/source/impact/no-code-diff gates accepted; the review's direct-consumer completeness finding is transferred without loss to `E2-P2A-INPUT1` because denominator closure belongs to P2-A | `E0-P0A-REVIEW1`, `E2-P2A-INPUT1` | close P0-A, open only P2-A, do not rerun P0 QA/review |
| R5 | 2026-08-13 | HEAD `f8b0717752c3d98e55556219567e21685c648207`; QA refresh `1,582/680/0`, `95,848/134,779`; main planner refresh `1,583/680/0`, `95,866/134,797`; production diff remains zero | P2-A frozen persistence/reader inventory and finding ownership | inventory `missing -> recorded/pending Supervisor`; persistence denominator `unknown -> 4` (`3` future edit); affected readers `unknown -> 8` (`3` future edit, `5` validate-only); P2-D boundaries `unknown -> 2`; duplicate/unassigned `unknown -> 0/0`; A14/A15 corrected to completed standalone audit/probe owners outside this campaign | `E2-P2A-IMPACT1`, `E2-P2A-SOURCE1`, `E2-P2A-INVENTORY1`, `E2-P2A-MATRIX1` | open only the P2-A Supervisor gate; do not open or edit P2-B/P2-C/P2-D until P2-A PASS and isolated commit |
| R6 | 2026-08-13 | same production HEAD; bounded Supervisor report SHA-256 `F02B0B894C6BFEECD379E7BF265721C5753FB5F74259FCE600B393418E026636`; production/test/runtime diff remains zero | only the three rejected P2-A inventory invariants | first inventory `recorded -> rejected for exactness`; corrected candidate refrozen to `8` persistence (`5/3`), `8` readers (`4/4`), `2` P2-D boundaries, `2` out-of-campaign exclusions, `19` total rows, and `0/0/0` duplicate/unassigned/unclassified; A02 exact branches, A12 linked persistence/reader roles, and `DEFINES` endpoints are now explicit | `E2-P2A-REJECT1`, `E2-P2A-SOURCE2`, `E2-P2A-INVENTORY2`, `E2-P2A-MATRIX2` | run only bounded P2-A re-review; keep P2-B/P2-C/P2-D closed and do not rerun accepted QA/impact gates without invalidation |
| R7 | 2026-08-13 | same production HEAD; bounded PASS report SHA-256 `988C137D5C864192C38180A47812BAB53FE315204DA4907F6FFCC15EEE659369`; exact five-doc/three-report boundary; staged paths were `0` at review | corrected P2-A acceptance and P2-B handoff | corrected inventory `pending re-review -> correct/accepted`; Supervisor independently cleared C02, C05/C06/C16, C07/C08, and `19`-row arithmetic with no residual same-invariant blocker; no production/runtime acceptance claimed | `E2-P2A-REVIEW1`, `E2-P2A-COMMIT1` | commit P2-A immediately, then open P2-B only; P2-C/P2-D remain closed |
| R8 | 2026-08-13 | HEAD before isolated commit `0dd955d40dbb422b90273b23d96e03105bd8fbd7`; independent Supervisor PASS report SHA-256 `24F403FC7A058AD36943A4664B0EFF8ABE2784644BA08BBC79D94B16C8CC750E`; fresh full build, normal artifacts, and final detect | P2-B corrected persistence and endpoint conservation | Graph JSON/Ladybug `partial -> correct/accepted`; `36,330/36,330` records and endpoint pairs, zero field mismatch/drop/missing endpoints; optional SelectionRange and zero/NULL semantics accepted; explicit CodeEmbedding label persisted; detect HIGH with `7` affected flows and zero degraded nodes | `E2-P2B-IMPACT1`, `E2-P2B-SOURCE1`, `E2-P2B-BUILD1`, `E2-P2B-TEST1`, `E2-P2B-PARITY1`, `E2-P2B-REVIEW1`, `E2-P2B-DETECT1`, `E2-P2B-COMMIT1` | exact isolated commit follows this row; after Git confirms success, open P2-C with C09/C10/C11/C16 edit and C12-C15 validate-only |
| R9 | 2026-08-13 | P2-C implementation baseline `4d456446fcc49aed0c6d489aa9c63e00d030b53c`; current HEAD `43aff01e882b787c91da84676e23f3ae28c05720` differs only by an unrelated one-file orchestration-skill commit; unconditional re-review SHA-256 `F54C4225520FFC57A6438CE3F4C3B0816FE8095B31AFB55FD004E91A89808698`; final detect HIGH file-layer scope with zero degraded/active gaps | C09-C16 affected readers and exact HTTP label-fixture sibling contract | affected readers `partial/wrong/unbound -> correct/accepted`; C09-C16 `8/8`; first HTTP sibling review blocker corrected without production change; VECTOR outage classified external/non-blocking; final detect PASS | `E2-P2C-IMPACT1`, `E2-P2C-SOURCE1`, `E2-P2C-BUILD1`, `E2-P2C-TEST1`, `E2-P2C-RUNTIME1`, `E2-P2C-REVIEW1`, `E2-P2C-DETECT1`, `E2-P2C-COMMIT1` | exact isolated P2-C commit follows; keep P2-D closed until Git confirms success, then measure existing behavior before any edit |
| R10 | 2026-08-13 | HEAD `927a676653963e8001d7789291010d5b819bac83`; fresh binary SHA-256 `C77D3E1F2C13BDC0BB3F089D476E9B340385E0385CA4F7DF7689D7ECFDD40BFB`; independent Supervisor PASS SHA-256 `C5582385945607FDB730B918EAA758AB2C7A3D37D3772B884D667B0E6133204D`; staged paths `0` at review | C04/C17 repeated analyze, current read, owned faults, recovery, and semantic determinism | repeated analyze `partial/unknown -> correct/accepted`; matrix `7/7`; successful analyze runs `4/4`; exact Definition/endpoint facts stable; changed input current; owned failures clear; exact build-capability sentinel classified non-blocking | `E2-P2D-IMPACT1`, `E2-P2D-SOURCE1`, `E2-P2D-BUILD1`, `E2-P2D-TEST1`, `E2-P2D-REPEAT1`, `E2-P2D-REVIEW1` | run main-owned final detect and isolated P2-D commit; keep P2-E closed until Git confirms success |

## Phase Touch Map

| Unit / File / Surface | Plan-Relevant Relationship File | Relationship to Target | Plan Item | Touch Mode | Evidence | Constraint |
|-----------------------|---------------------------------|------------------------|-----------|------------|----------|------------|
| Child 01 accepted fields | Graph JSON/Ladybug source owners | predecessor contract | P0/P2-A | inspect-only | `E0-P0A-HANDOFF1`, `E2-PNC-HANDOFF1` | accepted facts are inputs to P2-A inventory; no persistence schema or affected-reader assumption |
| `internal/resolution/emit.go::emitDefinitionNodes` | accepted Definition facts | graph property projection (C01) | P2-A/P2-B | future edit; production locked | `E2-P2A-IMPACT1`, `E2-P2A-SOURCE2`, `E2-P2A-INVENTORY2` | after P2-A PASS/commit, route only accepted construct-column/SelectionRange projection to P2-B |
| `internal/lbugload/csv.go::{ExportGraphCSVs,symbolNodeColumns,methodNodeColumns,defaultNodeColumns,nodeColumnLookup,nodeColumns,nodeCSVRow}` and `internal/lbugschema/schema.go::NodeSchema` | current Definition graph records | grouped Ladybug Definition node projection (C02) and schema (C03) | P2-A/P2-B | future edit; production locked | `E2-P2A-REJECT1`, `E2-P2A-SOURCE2`, `E2-P2A-INVENTORY2` | P2-B must cover Function/Class/Interface/CodeElement, Method, and default valid Definition-table branches; no narrower header subset |
| `internal/embeddings/pipeline.go::{EmbeddingUpdate,prepareBatch,CreateEmbeddingQuery}` and `internal/lbugschema/schema.go::EmbeddingSchema` | explicit label from `EmbeddableNode` | embedding label persistence/schema (C05/C06) | P2-A/P2-B | future edit; production locked | `E2-P2A-REJECT1`, `E2-P2A-SOURCE2`, `E2-P2A-INVENTORY2` | P2-B persists explicit label for the linked semantic-search reader; do not derive another identity contract |
| `internal/analyze/analyze.go::{writeGraphSnapshotJSON,Run}` | graph records | generic Graph JSON writer/normal analyze (C04) | P2-A/P2-B/P2-D | correct / accepted at P2-D; validate-only | `E2-P2D-SOURCE1`, `E2-P2D-TEST1`, `E2-P2D-REPEAT1`, `E2-P2D-REVIEW1` | preserve atomic Graph JSON selection and truthful normal analyze success/failure mechanics |
| `internal/lbugload/csv.go::{ExportGraphCSVs,relationshipColumns,relationshipCSVRow}` and `internal/lbugschema/schema.go::{RelationPairs,RelationSchema}` | accepted `DEFINES` relationships | relationship endpoint persistence (C07/C08) | P2-A/P2-B | validate-only; production locked | `E2-P2A-REJECT1`, `E2-P2A-SOURCE2`, `E2-P2A-INVENTORY2` | prove zero endpoint drop in P2-B; downstream COPY/load remains transparent unless new evidence invalidates it |
| `internal/httpapi/graph.go` and `backend-client.ts` | Graph JSON records | transparent transport | P2-A/P2-C | preserve-only | `E0-P0A-FD3`, `E0-P0A-FD4` | excluded from affected-reader denominator |
| `useAppState.local-runtime.tsx::{resolveNodeIds,handleNodeGroundingReference}` and `CodeReferencesPanel` | opaque identity and accepted one-based graph lines | semantic readers C09-C11 | P2-A/P2-C | correct / preserve after P2-C | `E2-P2C-SOURCE1`, `E2-P2C-TEST1`, `E2-P2C-RUNTIME1`, `E2-P2C-REVIEW1` | preserve exact-ID, unique-grounding, and coordinate behavior during P2-D; no reader redesign |
| `internal/embeddings/search.go::{SemanticSearch,hydrateSearchResults,metadataQuery}` plus label propagation in vector-result structs/helpers | explicit persisted label | semantic-search reader C16 | P2-A/P2-C | correct / preserve after P2-C | `E2-P2C-SOURCE1`, `E2-P2C-TEST1`, `E2-P2C-REVIEW1` | preserve explicit-label propagation, opaque identity, dedup, and fail-closed behavior; do not restore NodeID parsing |
| `internal/httpapi/search_test.go` semantic/hybrid vector rows | persisted `CodeEmbedding.label` test contract | C16 public HTTP sibling regression | P2-C | correct / test-only | `E2-P2C-TEST1`, `E2-P2C-REVIEW1` | positive persisted-vector fixtures carry explicit `Function` label; intentional missing-label negative remains in embeddings tests |
| `nodeRange`, Definition MCP context payload, `detectChangedSymbols`, `collectRenameChanges` | accepted identity/path/range facts | semantic readers C12-C15 | P2-A/P2-C | correct / validate-only | `E2-P2C-TEST1`, `E2-P2C-REVIEW1` | preserve aligned source; no rewrite without new slice evidence |
| `internal/mcp/tools.go::Server.runCypherRead` | Ladybug query boundary | repeated-read/failure boundary C17 | P2-A/P2-D | correct / accepted at P2-D; validate-only | `E2-P2D-SOURCE1`, `E2-P2D-TEST1`, `E2-P2D-REPEAT1`, `E2-P2D-REVIEW1` | preserve Ladybug-primary current read, non-sentinel fail-closed behavior, and same-repository source-classified capability fallback |
| `buildPropertyAccessAudit`; `nodeCanonicalKey`/`idName` | standalone audit/probe inputs | completed plan-owned validation tools C18/C19 | none in this campaign | out-of-campaign / preserve-only | `E2-P2A-SOURCE2`, `E2-P2A-MATRIX2` | owners are the completed 2026-05-19 property/access plan and completed 2026-05-16 Node-vs-Go accuracy plan; do not assign Child 03 or Child 07 |
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

Classification: P0 baseline `recorded/accepted`; the first P2-A candidate was rejected only for inventory exactness, and the corrected C01-C19 inventory is now `correct/accepted` under `E2-P2A-REVIEW1` and the immediate `E2-P2A-COMMIT1` boundary.

Allowed next action: after the isolated P2-A commit succeeds, open P2-B only. Refresh the graph, run file-detail and upstream impact on C01/C02/C03/C05/C06 before editing, report every blast radius, and validate C04/C07/C08 without expanding into P2-C/P2-D.

Forbidden next action: reuse a historical reader denominator or introduce storage behavior before proving a corrected-field gap.

### Repeated analyze

Current state:

The current normal packaged command passed the complete P2-D matrix at HEAD `927a676653963e8001d7789291010d5b819bac83` without production or existing product-test repair. Two unchanged runs preserve exact accepted Definition identities/ranges and exact `DEFINES` pairs; changed input removes the stale second-`now` identity and exposes `later`; analyze storage failure and no-readable-backend failure return clear non-success; and a subsequent analyze/read restores the baseline facts. Ladybug whole-file byte hashes may differ while the accepted semantic facts remain deterministic. The exact `ErrUnavailable` capability sentinel is unreachable in the tagged production binary without manufacturing a different runtime, so source classification plus all production-reachable dynamic outcomes is accepted as sufficient.

Required state:

```text
normal analyze invocation
-> normal artifact path implemented by source
-> affected reader returns matching corrected facts

failure
-> clear non-success at the normal boundary
```

Classification: `correct / accepted at P2-D`; preserve through the final detect/commit and P2-E independent validation.

## Next Phase Status Decisions

| Plan Item | Actual Status Finding | Required Status / Next-Action Update |
|-----------|-----------------------|--------------------------------------|
| P2-A | corrected C01-C19 inventory received bounded Supervisor PASS with `8` persistence, `8` readers, `2` P2-D boundaries, `2` out-of-campaign rows, and `0/0/0` duplicate/unassigned/unclassified | complete the immediate isolated documentation commit; do not rerun completed QA/impact gates without invalidation |
| P2-B | C01/C02/C03/C05/C06 correction and C04/C07/C08 validation received independent PASS with zero parity/drop/endpoint failures; detect reports HIGH scoped impact, seven affected flows, and zero degraded nodes | accepted; preserve the exact isolated source/test/ledger/report commit boundary |
| P2-C | C09-C16 received unconditional `8/8 PASS`; exact HTTP explicit-label fixture blocker is closed; final detect passed with scoped HIGH file-layer warning and zero degraded/active gaps | commit the exact isolated boundary immediately; do not rerun retained QA/review gates without invalidation |
| P2-D | C04/C17 existing behavior received unconditional PASS for `7/7` matrix rows, semantic determinism, clear owned failures, and recovery; exact production-unreachable capability fallback is explicitly non-blocking | complete main-owned final detect and immediate isolated commit without production repair |
| P2-E | P2-A denominator and P2-B/P2-C/P2-D accepted evidence are available, but the predecessor commit gate is not yet complete | remain closed until Git confirms the isolated P2-D commit; then open validation-only P2-E without production repair |

## Implementation Gate

- [x] Accepted Child 01 handoff is recorded.
- [x] Target scope is current in the status matrix.
- [x] Every P0 target unit has current baseline evidence and status; P2-A owns completeness.
- [x] Candidate file counts are refreshed at the implementation HEAD.
- [x] Exact persistence/reader symbols have current impact evidence.
- [x] First P2-A inventory received a bounded zero-trust verdict; only its three rejected inventory invariants were returned for correction.
- [x] Corrected P2-A candidate matrix C01-C19 and finding routes are recorded with zero duplicate/unassigned/unclassified rows.
- [x] P2-A bounded Supervisor PASS and immediate isolated documentation commit boundary are recorded before P2-B/P2-C.
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

The accepted Child 01 handoff and P0-A gates remain complete and are not rerun. P2-A's corrected inventory remains `8` persistence rows (`5/3`), `8` reader rows (`4/4`), `2` P2-D boundaries, `2` out-of-campaign rows, `19` total rows, and `0/0/0` duplicate/unassigned/unclassified. P2-B persistence and P2-C readers are accepted at isolated commits. P2-D now has unconditional `7/7 PASS` with no production or existing product-test repair; semantic determinism, current changed facts, clear owned failures, and recovery are accepted, and the production-unreachable `ErrUnavailable` capability sentinel is explicitly non-blocking. Final detect and the isolated P2-D commit immediately follow; P2-E stays closed until Git confirms that commit.
