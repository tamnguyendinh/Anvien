# Anvien Current Graph Persistence and Reader Consistency Plan

## Metadata

- Date: `2026-07-28`
- Status: `active / P0-A, P2-A, P2-B, P2-C, P2-D, P2-E, and Pn-A accepted at isolated commit boundaries / Pn-B cleanup independently accepted pending isolated commit / Pn-C closed`
- Plan: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-02-current-graph-persistence-and-reader-consistency/2026-07-28-02-current-graph-persistence-and-reader-consistency-plan.md`
- Evidence: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-02-current-graph-persistence-and-reader-consistency/2026-07-28-02-current-graph-persistence-and-reader-consistency-evidence.md`
- Benchmark: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-02-current-graph-persistence-and-reader-consistency/2026-07-28-02-current-graph-persistence-and-reader-consistency-benchmark.md`
- Actual status: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-02-current-graph-persistence-and-reader-consistency/2026-07-28-02-current-graph-persistence-and-reader-consistency-actual-status.md`
- Roadmap: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-26-anvien-graph-accuracy-roadmap.md`
- Contract: `docs/contracts/graph-accuracy-contract.md`
- Predecessor: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-plan.md`
- Successor: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-03-typescript-binding-pattern-extraction/2026-07-28-03-typescript-binding-pattern-extraction-plan.md`

## Goal

Preserve the corrected graph facts through Graph JSON and Ladybug, make only source-proven affected readers consume those facts consistently, and verify repeated normal analyze behavior on the artifact paths and runtime mechanisms actually implemented by the codebase.

## Rules

- Complete P0 actual status and receive the accepted Child 01 handoff before production work.
- Work directly in the repository-root production codebase with normal commands.
- The persistence and analyze implementation observed at P0 is baseline evidence; any change is decided by the owning slice's source and impact proof.
- P2-A must inspect source and produce the exact affected-field/reader inventory. A reader is in scope only when its command or code truly consumes a corrected field.
- Do not use a historical reader list, transport category, or broad command family as the denominator.
- Keep the existing graph output path under observation; add or change a contract only when current source/impact evidence proves the corrected fields require it and this plan is updated.
- Analyze remains a normal repeatedly invoked command. P2-D tests what the built runtime actually does on the same repository artifact path.
- A failed analyze or graph-bound read reports a clear failure at its normal boundary. Data from another artifact is not substituted as a successful result.
- Graph JSON and Ladybug must conserve corrected identifiers, ranges, meanings, facts, and endpoints without silent record loss.
- Implement production behavior before changing tests. Run the full build before nearest-real-boundary validation.
- Refresh the graph and record file-detail plus upstream impact before editing any writer, loader, query adapter, command handler, or shared contract.
- Every touched production/test file owns one primary semantic responsibility. No catch-all adapter or unrelated refactor.
- Update checklist, actual status, evidence, and benchmark as each state changes.
- Every implementation slice requires Supervisor PASS, detect-changes, and an isolated commit before its successor opens.
- Validation-only P2-E does not repair production code; failures return to P2-B, P2-C, or P2-D.
- Target source and pre-existing worktree are preserve-only. Scanner repair remains outside the campaign.
- A finding does not widen the slice that discovered it. Record its exact evidence and route it to the owning slice; only an invariant inside the current slice's stated goal and boundary can PASS or REJECT that slice.

## Problem

Correct source facts are not useful if persistence drops them or readers reconstruct a different meaning. The bounded investigation observed representation changes across raw graph, Ladybug/native projection, and file-detail for selected facts, while preserving cardinality for those selected rows. Separately, the identity correction from Child 01 introduces fields whose real persistence and reader consumers must be discovered from current source.

No reader enters scope from an interface or subsystem label. The plan must first prove the corrected-field flow and then change only the concrete persistence or reader owners that participate in it.

## Scope

In scope:

- inventory of corrected Child 01 fields and every persistence/read path that actually consumes them;
- Graph JSON preservation and decode integrity for those fields;
- Ladybug schema/load/query preservation for those fields;
- exact query/reader adapters proven affected by P2-A;
- repeated normal analyze on the same repository/artifact path;
- clear error behavior at affected normal command/runtime boundaries;
- field-level Graph JSON/Ladybug/affected-reader parity and zero dropped corrected records;
- accepted handoff of persistence facts to Child 03.

## Non-Goals

- a universal reader inventory unrelated to corrected fields;
- a fixed surface taxonomy or historical reader-row denominator;
- unrelated consumer, coordination, or identifier behavior without corrected-field source evidence;
- any storage behavior not required by the corrected-field evidence;
- binding, export, module/re-export, ambient/external, graph-health, or UI redesign;
- changing unrelated analyze stages merely to satisfy the plan;
- target-source edits or scanner repair.

## Requirements

- P2-A records, for every candidate path: exact source symbol/path, corrected field read or written, actual impact, touch mode, and evidence ID.
- The affected-reader denominator is the number of source-proven affected readers, not a historical or transport-wide total.
- A candidate that does not consume a corrected field is preserve-only and excluded from P2-C acceptance counts.
- Graph JSON encode/decode preserves every accepted Child 01 identity/range/meaning/occurrence fact and endpoint.
- Ladybug projection/load/query preserves the same corrected semantic fields and records; any intentional representation mapping is explicit and proven lossless.
- No serializer, CSV/load path, or query adapter silently skips a corrected record.
- P2-C changes only the exact readers identified by P2-A and uses explicit corrected fields rather than deriving meaning from opaque identifiers.
- P2-D invokes the normal built command repeatedly against the same repository path and evaluates the artifact path/mechanism implemented by current source.
- P2-D does not require an implementation change when existing repeated behavior already satisfies the acceptance contract.
- When an invocation or affected read fails, the normal boundary returns a clear failure and no successful result sourced from a different artifact.
- If P2-A/P2-B impact proves an additional persistence or reader owner is required, update the plan and actual-status touch map before editing it.

## Acceptance Criteria

- P2-A has a complete source-proven affected-path inventory with zero unassigned affected readers.
- Graph JSON and Ladybug have zero differing corrected records/fields under the accepted comparison and zero dropped corrected records.
- Every affected reader returns the same corrected semantic fact as its source persistence record.
- No unaffected reader or unrelated transport is edited merely to satisfy a denominator.
- Repeated normal analyze succeeds for every declared matched run and the latest result reflects the analyzed source/configuration.
- Injected or naturally reproduced failure at an owned boundary is visible and does not return a successful graph result from another artifact.
- Full build, nearest real runtime, affected-reader regression, evidence, benchmark, Supervisor, detect-changes, and isolated commit gates pass for every implementation slice.

## Checklist

- [x] P0-A: Complete actual status and predecessor dependency before implementation work.
  - Goal: classify current persistence/read behavior and prevent any reader or writer from entering scope by assumption.
  - Work Steps:
    1. Consume the accepted Child 01 handoff and list its corrected fields without adding new schema requirements.
    2. Refresh the graph; inspect Graph JSON/Ladybug/analyze/read source; run file-detail and impact on every candidate owner.
    3. Record file counts, current classifications, touch modes, and the pending P2-A discovery boundary.
    4. Obtain Supervisor review of the P0 scope/dependency decision.
  - Implementation Gate: no production edit until the accepted Child 01 handoff exists and this actual-status ledger has a complete, current P0 decision.
    - Acceptance: `E0-P0A-REPORT1`, `E0-P0A-VERIFY1`, `E0-P0A-HANDOFF1`, `E0-P0A-QA1`, `E0-P0A-GRAPH1`, `E0-P0A-SOURCE1`, `E0-P0A-SOURCE2`, `E0-P0A-FD1`, `E0-P0A-FD2`, `E0-P0A-FD3`, `E0-P0A-FD4`, `E0-P0A-IMPACT1`, `E0-P0A-BOUNDARY1`, `E0-P0A-STATUS1`, and the scope-corrected review disposition `E0-P0A-REVIEW1` are exact and non-contradictory.
    - Current State: accepted at HEAD `22b8ed4cb0d7393cc5c982586b2e9a5c4c049760`. P0-A recorded the current-codebase baseline, accepted Child 01 dependency, candidate owner counts/impacts, preserve-only boundaries, and the next-slice decision. The review's additional direct-consumer findings are retained as `E2-P2A-INPUT1`; exact denominator and zero-unassigned proof were not P0-A acceptance criteria and remain P2-A-owned. No P0-A QA or review gate is rerun.

### P2: Current graph persistence and reader consistency

- Phase Goal: discover affected paths, preserve corrected fields through persistence, correct affected readers, verify repeated normal analyze, and close parity.
- Phase Boundary:
  - In scope: P2-A through P2-E below.
  - Out of scope: later semantic Children and any unproven storage/reader behavior.
  - Dependencies: P0 and Child 01 handoff complete; each slice accepted/committed before its successor.
- Phase Implementation Rule: execute five independent slices; do not implement P2 as a broad rewrite.
- Ordered Slice List:
  - P2-A: Inventory affected persistence and reader paths.
  - P2-B: Preserve corrected facts through Graph JSON and Ladybug.
  - P2-C: Correct source-proven affected readers.
  - P2-D: Verify repeated normal analyze and clear failure behavior.
  - P2-E: Close persistence/reader parity and handoff.

- [x] P2-A: Inventory affected persistence and reader paths.
  - Goal: establish the exact field-flow denominator before production edits.
  - Scope Boundary:
    - Editable: this Child's four ledgers only.
    - Inspect-only: Graph JSON/Ladybug/analyze/query/reader source and accepted Child 01 output.
    - Preserve-only: all production source, tests, runtime, and target source.
    - Out of scope: reader/writer implementation.
  - Non-Goals: no reused denominator, broad transport row, or guessed owner path.
  - Finding Routing:
    - persistence/projection gaps -> record for P2-B;
    - source-proven semantic reader gaps -> record for P2-C;
    - repeated-analyze or failure-boundary gaps -> record for P2-D;
    - later-Child or unrelated findings -> record the exact deferred owner and keep them out of P2-A acceptance.
  - Pre-flight Questions:
    - Data source: accepted Child 01 fields and current persistence/reader source.
    - Display permission: N/A — source inventory only.
    - DB read flow: inspect actual Graph JSON/Ladybug/query open/read flows.
    - DB write flow: N/A — no runtime state change.
    - Render location: evidence ledger affected-path table.
    - UI behavior flow: N/A unless source proves a Web reader is affected; inventory still does not edit it.
    - Docker runtime: N/A — source/documentation slice.
    - Playwright target: N/A — no UI behavior changed.
    - Behavior test: source anchor existence, corrected-field flow, unique owner, and affected/unaffected classification.
    - Cleanup/quarantine: remove unsupported denominator assumptions from active ledgers.
    - External side effects: none.
    - N/A notes: implementation/runtime gates begin in P2-B.
  - Work Steps:
    1. Consume `E2-P2A-INPUT1` as candidate leads only, refresh graph, and trace each accepted corrected field through Graph JSON, Ladybug, and every reader reached by real source flow; use P0 evidence only as current-codebase baseline, not as a completed denominator.
       - UI flow check: record an affected UI reader only if source proves the field dependency.
       - DB/data flow check: record exact write/read field and backend.
       - Render location check: evidence ledger inventory.
       - Mini QA: validate source anchors and field-flow traces.
       - Evidence target: `E2-P2A-IMPACT1`, `E2-P2A-SOURCE1`, `E2-P2A-INVENTORY1`.
    2. Classify each row as edit, validate-only, preserve-only, or out-of-scope; obtain Supervisor PASS and commit the inventory slice.
       - UI flow check: N/A — no UI changed.
       - DB/data flow check: affected reader count equals exact unique affected rows.
       - Render location check: evidence and actual-status ledgers.
       - Mini QA: zero duplicate/unassigned affected rows and zero broken anchors.
       - Evidence target: `E2-P2A-MATRIX1`, `E2-P2A-REVIEW1`, `E2-P2A-COMMIT1`.
  - Implementation Gate: P0 and Child 01 handoff accepted; graph/source evidence current.
  - Current State: accepted by bounded zero-trust re-review report `reports/Supervisor/rp_supervisor_260813_131254_by_gpt-5-codex_child02_p2a_inventory_rereview.md` at SHA-256 `988C137D5C864192C38180A47812BAB53FE315204DA4907F6FFCC15EEE659369`. The accepted matrix contains `8` persistence rows (`5` future edit, `3` validate-only), `8` affected Child 02 reader rows (`4` future edit, `4` validate-only), `2` P2-D boundaries with `analyze` shared with persistence, `2` out-of-campaign audit/probe exclusions, and `19` total unique classified rows with zero duplicate, unassigned, or unclassified mandatory leads. The review independently closes A02's exact Definition CSV branches, the linked A12 embedding-label persistence/semantic-search reader split, `DEFINES` endpoint CSV/schema ownership, and the refrozen arithmetic. No production/test/runtime behavior is accepted by this docs-only slice. The exact five-doc/three-report boundary is committed immediately after this row as `E2-P2A-COMMIT1`; Git reports the final hash because a commit cannot contain its own hash. P2-B may open only after that isolated commit succeeds.
  - Acceptance:
    - Source: every affected persistence/read path has exact symbol/path and field evidence.
    - Runtime/UI: N/A — no runtime/UI behavior changed.
    - DB/data: affected denominator is exact; unaffected paths are excluded and preserve-only.
    - Behavior test: source-anchor/ownership audit passes with zero unassigned affected readers.
    - Cleanup/quarantine: unsupported reader assumptions absent from active scope.
    - Evidence IDs: `E2-P2A-IMPACT1`, `E2-P2A-SOURCE1`, `E2-P2A-INVENTORY1`, `E2-P2A-MATRIX1`, `E2-P2A-REJECT1`, `E2-P2A-SOURCE2`, `E2-P2A-INVENTORY2`, `E2-P2A-MATRIX2`, `E2-P2A-REVIEW1`, `E2-P2A-COMMIT1`.
    - Actual-status rows refreshed: persistence candidates, affected readers, touch map, P2-B/P2-C steps.
  - Evidence Targets: source trace, affected inventory, review, documentation commit.
  - Actual-status Update: affected-path inventory `missing -> correct`; implementation rows remain partial/wrong as measured.
  - Commit Boundary: documentation/source-audit commit only.

- [x] P2-B: Preserve corrected facts through Graph JSON and Ladybug.
  - Goal: make the exact P2-A persistence owners preserve accepted corrected records and fields without loss.
  - Scope Boundary:
    - Editable: only Graph JSON/Ladybug encode, load, or query owners proven affected by P2-A.
    - Inspect-only: accepted Child 01 source records and P2-C readers.
    - Preserve-only: unaffected fields/backends/readers and analyze orchestration not implicated by impact.
    - Out of scope: reader UI/transport changes and unrelated output behavior.
  - Non-Goals: no unsupported output contract, unrelated scalar rewrite, or silent normalization.
  - Pre-flight Questions:
    - Data source: accepted corrected graph records and P2-A field-flow inventory.
    - Display permission: N/A — persistence layer.
    - DB read flow: decode/query Graph JSON and Ladybug through their actual built paths.
    - DB write flow: use the existing production persistence entrypoints unless impact proves a required change.
    - Render location: parity manifest and command evidence.
    - UI behavior flow: N/A — no UI owner.
    - Docker runtime: N/A unless an affected real boundary requires the packaged service; otherwise full build plus real store/query boundary.
    - Playwright target: N/A — no browser behavior owned.
    - Behavior test: corrected field round-trip, duplicate/drop detection, endpoint integrity, empty/subset, and intentional representation mapping.
    - Cleanup/quarantine: isolated repository-local fixtures; remove superseded debug data.
    - External side effects: none beyond isolated and normal repository artifacts.
    - N/A notes: persistence mechanism is discovered, not prescribed.
  - Work Steps:
    1. Refresh graph; record exact file/symbol impact; implement the smallest production changes in the affected persistence owners.
       - UI flow check: N/A — non-UI.
       - DB/data flow check: trace accepted source record to Graph JSON and Ladybug records.
       - Render location check: evidence ledger.
       - Mini QA: real Graph JSON/Ladybug encode/load/query boundary.
       - Evidence target: `E2-P2B-IMPACT1`, `E2-P2B-SOURCE1`.
    2. Add focused tests after production code, run full build, and compare corrected records field-for-field with zero drops.
       - UI flow check: N/A — non-UI.
       - DB/data flow check: Graph JSON/Ladybug parity and endpoint closure.
       - Render location check: parity manifest.
       - Mini QA: built runtime store/query probe.
       - Evidence target: `E2-P2B-BUILD1`, `E2-P2B-TEST1`, `E2-P2B-PARITY1`.
  - Implementation Gate: P2-A accepted; every editable persistence owner and field is named by an inventory row with fresh impact.
  - Current State: production correction is complete in C01/C02/C03/C05/C06 and C04/C07/C08 remain validate-only. Independent current-HEAD review `reports/Supervisor/rp_supervisor_260813_173541_by_gpt-5_child02_p2b_persistence.md` (SHA-256 `24F403FC7A058AD36943A4664B0EFF8ABE2784644BA08BBC79D94B16C8CC750E`) returned unconditional `PASS` / `ACCEPT_P2B_AND_HAND_BACK`. Fresh root full build, affected-package and CLI/HTTP sibling regressions, native Ladybug readback, and normal Graph JSON/Ladybug record parity all passed. The real artifact contains `36,330/36,330` matching Definition records, zero corrected-field mismatch/drop, `36,330/36,330` exact `DEFINES` pairs, and zero missing endpoints. `E2-P2B-DETECT1` passes with HIGH scoped impact and no degraded nodes; the exact source/test/ledger/report boundary is committed immediately after this row as `E2-P2B-COMMIT1`, with Git reporting the final hash because a commit cannot contain its own hash. P2-C opens only after that commit succeeds.
  - Acceptance:
    - Source: only proven persistence owners changed.
    - Runtime/UI: built persistence/query boundary passes; no UI change.
    - DB/data: zero corrected field differences, zero dropped corrected records, zero affected missing endpoints.
    - Behavior test: positive and fault/round-trip matrices pass.
    - Cleanup/quarantine: only the current output path and accepted evidence remain.
    - Evidence IDs: `E2-P2B-IMPACT1`, `E2-P2B-SOURCE1`, `E2-P2B-BUILD1`, `E2-P2B-TEST1`, `E2-P2B-PARITY1`, `E2-P2B-REVIEW1`, `E2-P2B-DETECT1`, `E2-P2B-COMMIT1`.
    - Actual-status rows refreshed: Graph JSON/Ladybug field and drop rows.
  - Evidence Targets: impact, source, full build, round-trip/parity, Supervisor, detect, commit.
  - Actual-status Update: affected persistence `partial/wrong -> correct`.
  - Commit Boundary: isolated implementation-slice commit.

- [x] P2-C: Correct source-proven affected readers.
  - Goal: make only P2-A affected readers consume the corrected persisted fields consistently.
  - Scope Boundary:
    - Editable: exact reader/query adapter rows marked edit by P2-A.
    - Inspect-only: P2-B persistence records and affected command dispatchers.
    - Preserve-only: all readers not proven affected.
    - Out of scope: any reader selected only by an interface or subsystem category.
  - Non-Goals: no invented reader guard, identifier parsing workaround, or unrelated response redesign.
  - Pre-flight Questions:
    - Data source: P2-B corrected records and P2-A exact reader rows.
    - Display permission: preserve existing permission/visibility; no new access rule.
    - DB read flow: actual backend used by each affected reader.
    - DB write flow: N/A unless an affected reader owns a source-proven cache write; update plan first if so.
    - Render location: each affected reader's normal output surface.
    - UI behavior flow: required only for an affected UI reader proven by P2-A.
    - Docker runtime: required only for affected app/runtime surfaces; otherwise use the nearest built command or query boundary.
    - Playwright target: required only for affected browser-visible readers, using reusable scripts and real built runtime.
    - Behavior test: one positive/negative corrected-field vector per affected reader.
    - Cleanup/quarantine: retain only reusable fixtures and official evidence.
    - External side effects: none.
    - N/A notes: unaffected readers are regression-only and excluded from denominator.
  - Work Steps:
    1. Refresh graph; record exact handler/adapter impact; implement explicit corrected-field consumption only in affected rows.
       - UI flow check: execute only source-proven affected UI rows.
       - DB/data flow check: each reader traces to the P2-B record without reconstructing lost meaning.
       - Render location check: actual command/API/UI output for each affected row.
       - Mini QA: nearest real boundary per affected row.
       - Evidence target: `E2-P2C-IMPACT1`, `E2-P2C-SOURCE1`.
    2. Add tests after production code; run full build and all affected-reader acceptance plus sibling regression required by impact.
       - UI flow check: real built runtime and visual inspection only when affected UI exists.
       - DB/data flow check: affected reader outputs equal corrected persistence facts.
       - Render location check: evidence captures for exact affected rows.
       - Mini QA: commands/runtime listed by the inventory.
       - Evidence target: `E2-P2C-BUILD1`, `E2-P2C-TEST1`, `E2-P2C-RUNTIME1`.
  - Implementation Gate: P2-B accepted; P2-A affected rows and exact edit owners current.
  - Current State: C09-C16 received unconditional `8/8 PASS` from independent re-review `reports/Supervisor/rp_supervisor_260813_203838_by_gpt-5_child02_p2c_affected_readers_rereview.md` (SHA-256 `F54C4225520FFC57A6438CE3F4C3B0816FE8095B31AFB55FD004E91A89808698`). C09 now resolves only exact opaque IDs; C10 fails closed unless grounding has exactly one persisted identity; C11 preserves one-based lines, zero-based UTF-8 byte columns, and exclusive ends through display/highlight/API conversion; C16 consumes persisted `CodeEmbedding.label` through vector rows, dedup, and metadata hydration without NodeID label reconstruction. C12-C15 remain unchanged and pass their nearest regressions. A clean-holder full build passed in `121.4s`, focused frontend passed `11/11`, HTTP/embeddings/filecontext/MCP regressions passed, tagged local-native label persistence/readback/hydration passed, and mounted Playwright evidence passed C09-C11 with six visually accepted screenshots. The first review correctly rejected stale HTTP semantic/hybrid vector fixtures that lacked the explicit label; the bounded follow-up added only three `Function` labels and the re-review cleared the exact blocker. The official VECTOR extension outage is an external, non-blocking evidence limitation rather than a P2-C production defect. Final detect passed with HIGH file-layer scope, `13` tracked changed files, `12` affected files, `154` changed graph entries, no affected processes/flows, and zero degraded/active resolution gaps. The exact implementation/test/plan/report/QA boundary was committed as `927a676653963e8001d7789291010d5b819bac83`, opening P2-D.
  - Acceptance:
    - Source: only affected readers changed; unaffected readers remain preserve-only.
    - Runtime/UI: every affected normal boundary exposes the corrected fact; affected UI, if any, passes real-runtime visual QA.
    - DB/data: reader output agrees with P2-B source record; no silent missing/ambiguous field.
    - Behavior test: affected readers passed/total equals the exact P2-A denominator.
    - Cleanup/quarantine: no category-wide adapter or dead QA artifact remains.
    - Evidence IDs: `E2-P2C-IMPACT1`, `E2-P2C-SOURCE1`, `E2-P2C-BUILD1`, `E2-P2C-TEST1`, `E2-P2C-RUNTIME1`, `E2-P2C-REVIEW1`, `E2-P2C-DETECT1`, `E2-P2C-COMMIT1`.
    - Actual-status rows refreshed: every affected reader and preserved sibling regression.
  - Evidence Targets: impact, source, build, affected-reader runtime/QA, Supervisor, detect, commit.
  - Actual-status Update: affected reader rows `partial/wrong/unbound -> correct`.
  - Commit Boundary: isolated implementation-slice commit.

- [x] P2-D: Verify repeated normal analyze and clear failure behavior.
  - Goal: prove that repeatedly invoking the normal built analyze command on the same repository path yields current, readable corrected facts or a clear failure.
  - Scope Boundary:
    - Editable: only exact analyze/read owner proven wrong by P2-D pre-flight; validation harnesses and ledgers.
    - Inspect-only: current analyze implementation, artifact locations, P2-B persistence, and P2-C affected readers.
    - Preserve-only: output mechanics that already satisfy acceptance and unrelated runtime paths.
    - Out of scope: changing output behavior by assumption.
  - Non-Goals: no additional graph output, substitute result, or unrelated output change by assumption.
  - Pre-flight Questions:
    - Data source: normal source/config inputs and accepted P2-B/P2-C corrected facts.
    - Display permission: preserve current access behavior.
    - DB read flow: read the artifact produced by each normal invocation through affected real readers.
    - DB write flow: invoke the existing normal analyze command; change implementation only if evidence proves acceptance fails.
    - Render location: normal CLI and exact affected reader outputs.
    - UI behavior flow: only if P2-A identified an affected UI reader.
    - Docker runtime: required only for an affected app/runtime boundary; otherwise full build plus packaged command.
    - Playwright target: required only for affected browser-visible behavior.
    - Behavior test: repeated unchanged input, changed source/config input, command failure, read failure, and subsequent successful analyze.
    - Cleanup/quarantine: isolated repo-local fault fixtures; no target contamination.
    - External side effects: normal analyze output only.
    - N/A notes: the output method is observed at the normal command boundary, not selected by this slice.
  - Work Steps:
    1. Refresh graph; inspect current repeated-analyze source and impact; first test existing behavior through the built command.
       - UI flow check: only affected UI rows if present.
       - DB/data flow check: associate each invocation with the facts read from its normal artifact path.
       - Render location check: command and reader evidence.
       - Mini QA: normal built analyze plus exact affected reads.
       - Evidence target: `E2-P2D-IMPACT1`, `E2-P2D-SOURCE1`.
    2. If existing behavior fails, update the plan/actual status, implement only the proven owner, then add tests; run full build and complete the repeat/failure matrix.
       - UI flow check: real built runtime only when affected UI exists.
       - DB/data flow check: all declared repeated runs expose their matching corrected facts; failures are visible and no other artifact supplies success.
        - Render location check: runtime/fault evidence.
       - Mini QA: built normal command and affected reader boundary.
       - Evidence target: `E2-P2D-BUILD1`, `E2-P2D-TEST1`, `E2-P2D-REPEAT1`.
  - Implementation Gate: P2-C accepted; current repeated behavior measured before any production edit.
  - Current State: existing production mechanics required no correction. Independent Supervisor report `reports/Supervisor/rp_supervisor_260813_220611_by_gpt-5_child02_p2d_repeated_analyze.md` (SHA-256 `C5582385945607FDB730B918EAA758AB2C7A3D37D3772B884D667B0E6133204D`) returned unconditional `PASS` / `ACCEPT_P2D_AND_HAND_BACK`. The reusable normal-built matrix passed `7/7`: two unchanged analyzes retained identical accepted Definition facts and exact `DEFINES` pairs; changed input exposed `later` and removed the stale second-`now` identity; owned analyze storage failure returned exit `1`; Ladybug remained readable with Graph JSON absent; both backends absent returned JSON-RPC non-success with no rows; and recovery restored the baseline facts. The clean-holder full build passed in `130.822s`. Whole Ladybug file hashes are informational rather than the determinism contract; exact identities, construct/selection ranges, and endpoints are stable. Exact `ErrUnavailable -> same-repository graph.json` was source-classified but not fabricated because the normal packaged `ladybugdb` binary cannot truthfully emit that compile-time capability sentinel; Supervisor explicitly accepted this as non-blocking after all production-reachable success/failure outcomes passed. Final all/staged detect passed at LOW risk and the exact ten-path boundary was committed as `35939e7e6a621593d3d3065b9493a97c2c9a4f25`, opening P2-E.
  - Acceptance:
    - Source: existing correct mechanics preserved; only a source-proven failing owner changed.
    - Runtime/UI: all declared normal invocations behave truthfully; clear non-success on owned failure.
    - DB/data: each successful run exposes matching corrected facts; failure returns no successful result from another artifact.
    - Behavior test: repeat/change/failure/subsequent-success matrix passes at the normal boundary.
    - Cleanup/quarantine: no additional graph output or debug residue.
    - Evidence IDs: `E2-P2D-IMPACT1`, `E2-P2D-SOURCE1`, `E2-P2D-BUILD1`, `E2-P2D-TEST1`, `E2-P2D-REPEAT1`, `E2-P2D-REVIEW1`, `E2-P2D-DETECT1`, `E2-P2D-COMMIT1`.
    - Actual-status rows refreshed: repeated analyze, failure, and affected read results.
  - Evidence Targets: current source/impact, build, repeat/fault runtime, Supervisor, detect, commit.
  - Actual-status Update: repeated normal analyze `partial/unknown -> correct`.
  - Commit Boundary: isolated implementation/validation-slice commit.

- [x] P2-E: Close persistence/reader parity and handoff.
  - Goal: independently validate Child 02 without repairing production code and hand accepted persistence facts to Child 03.
  - Scope Boundary:
    - Editable: validation manifests and four Child ledgers only.
    - Inspect-only: accepted P2-A through P2-D source/runtime and affected readers.
    - Preserve-only: production source, target source, and unaffected readers.
    - Out of scope: code repair and later semantic implementation.
  - Non-Goals: no expansion to unproven readers or campaign-wide acceptance.
  - Pre-flight Questions:
    - Data source: accepted corrected records, affected reader inventory, and repeated normal analyze results.
    - Display permission: preserve current access behavior.
    - DB read flow: Graph JSON, Ladybug, and exact affected readers.
    - DB write flow: normal analyze and isolated validation output only.
    - Render location: evidence/benchmark ledgers and exact affected surfaces.
    - UI behavior flow: only if an affected UI row exists.
    - Docker runtime: required only for affected app/runtime surfaces.
    - Playwright target: required only for affected browser-visible rows.
    - Behavior test: field parity, dropped-record count, affected-reader count, repeated analyze, and clear failure.
    - Cleanup/quarantine: final evidence only; remove failed/duplicate validation artifacts.
    - External side effects: normal repository output only.
    - N/A notes: failures reopen their exact owning P2 slice.
  - Work Steps:
    1. Run full build and the complete corrected-field parity, affected-reader, and repeated-analyze acceptance matrix.
       - UI flow check: exact affected UI rows only.
       - DB/data flow check: zero differing/dropped corrected records and full affected-reader coverage.
       - Render location check: final manifests and real affected outputs.
       - Mini QA: nearest real boundary for every affected row.
       - Evidence target: `E2-P2E-BUILD1`, `E2-P2E-PARITY1`, `E2-P2E-READERS1`, `E2-P2E-REPEAT1`.
    2. Obtain Supervisor PASS, run detect-changes, commit validation/ledger closure, and refresh Child 03 from exact accepted facts.
       - UI flow check: N/A beyond any affected row already validated.
       - DB/data flow check: handoff lists only accepted persisted fields and affected-reader guarantees.
       - Render location check: Supervisor/evidence/handoff records.
       - Mini QA: link/evidence/status audit.
       - Evidence target: `E2-P2E-REVIEW1`, `E2-P2E-DETECT1`, `E2-P2E-COMMIT1`.
  - Implementation Gate: P2-A through P2-D accepted and committed; exact affected denominator recorded.
  - Current State: independent validation and zero-trust review are complete at production HEAD `35939e7e6a621593d3d3065b9493a97c2c9a4f25`. QA report `reports/QA/rp_qa_260813_233238_by_gpt-5_child02_p2e_persistence_reader_parity.md` (SHA-256 `41509F4379B3B5A4976BC7469831EF8E57EA902FABBAE4BD02184EEFB14D8D16`) records a one-attempt clean-holder full build in `116.702s`, C01-C17 `17/17`, persistence `8/8`, affected readers `8/8`, and repeated-analyze/failure/recovery `7/7`, with no production or existing product-test edit. Field parity is `36,611/36,611` Definitions with zero missing/extra/mismatch/drop, optional SelectionRange `4,941` present / `31,670` absent / `0` partial, `5,650` real zero-valued `startCol`, exact `36,611/36,611` `DEFINES` pairs, and zero missing endpoints. Supervisor report `reports/Supervisor/rp_supervisor_260813_235237_by_gpt-5_child02_p2e_validation.md` (SHA-256 `DC2A7FC600E49529EA675640D03F5251A864B563E23B4F5B6190A7BAF32837F7`) returned unconditional `PASS` / `ACCEPT_P2E_AND_HAND_BACK`, no blocking or non-blocking product findings, and no residual same-invariant surface. The citation-card text `Code not available in memory for src/unique.ts` is an intentional no-snippet card state outside C10/C11; the selected Code Inspector independently read the exact file, highlighted lines 10-11, and excluded line 12. Final graph refresh passed at `1,625/688/0` files and `97,267/136,541` graph nodes/relationships; all-scope detect is LOW with five changed/affected documentation files, 19 changed sections, zero affected flows/processes, zero changed or active ResolutionGap entities, and zero degraded nodes. Exact 28-file staging has no unstaged residue; staged detect reports overall MEDIUM/file-layer HIGH because it now parses the reusable parity harness and report evidence, with 22 parseable changed files, 15 affected files, two harness-local processes, 282 touched gap-evidence entities, zero active gaps, and zero degraded nodes. The isolated validation/ledger commit follows this row; Pn-A and later work remain closed until Git confirms that commit.
  - Acceptance:
    - Source: no production repair in P2-E.
    - Runtime/UI: all affected real boundaries pass; affected UI, if any, visually verified.
    - DB/data: Graph JSON/Ladybug corrected-field parity, zero drops, and complete affected-reader coverage.
    - Behavior test: repeated analyze and clear-failure matrix passes.
    - Cleanup/quarantine: only current final evidence remains.
    - Evidence IDs: `E2-P2E-BUILD1`, `E2-P2E-PARITY1`, `E2-P2E-READERS1`, `E2-P2E-REPEAT1`, `E2-P2E-REVIEW1`, `E2-P2E-DETECT1`, `E2-P2E-COMMIT1`.
    - Actual-status rows refreshed: all Child 02 rows and Child 03 handoff assumptions.
  - Evidence Targets: full build, parity/readers/repeat, Supervisor, detect, commit.
  - Actual-status Update: Child 02 scope `partial/missing -> correct`; refresh successor.
  - Commit Boundary: isolated validation/ledger commit.
  - Commit Result: exact 28-file boundary committed as `593e77a3f36c78447864a906a75c05e0d89530cc` with message `test(graph): close child 02 persistence parity`; post-commit worktree and index are clean.

- [x] Pn-A: Call Supervisor for the implemented-plan acceptance loop.
  - Goal: independently verify source, diff, persistence, readers, runtime, reports, and ledgers.
  - Work Steps:
    1. Run the Supervisor skill on complete Child 02 work.
    2. Return each rejection only to P2-B, P2-C, P2-D, or evidence ownership and repeat review.
  - Implementation Gate: P2-A through P2-E complete or explicitly blocked.
  - Current State: independent report `reports/Supervisor/rp_supervisor_260814_005128_by_gpt-5_child02_pna_full_plan.md` (SHA-256 `60314C3BAFDAB09E4A60539391C6E7A847B5FA411E08129E1CE12555BFECC9E0`) returned unconditional `PASS` / `ACCEPT_CHILD02_PNA_AND_HAND_BACK` at HEAD `593e77a3f36c78447864a906a75c05e0d89530cc`. It independently verified source, all six slice commits, Graph JSON/Ladybug parity, eight readers, repeat/failure/recovery, raw evidence/harnesses/screenshots, full-plan ledger consistency, and zero residual same-invariant surface. Earlier P2-A/P2-C REJECT artifacts are retained immutable provenance whose findings are closed, not live blockers. The exact five-document/one-report boundary commits immediately after this row; Pn-B/Pn-C and Child 03 remain closed until Git confirms success.
  - Commit Result: exact five-document/one-report boundary committed as `e47acfad927425621c3f9048d0a23eed513444a5` with message `docs(plan): accept child 02 full-plan review`; post-commit worktree and index are clean.
  - Acceptance: unconditional PASS recorded as `E3-PNA-SUPERVISOR1`, or exact blocker prevents closure.

- [x] Pn-B: Remove dead work created during this plan.
  - Goal: remove superseded inventories, failed fixtures, duplicate reports, and rejected artifacts.
  - Work Steps:
    1. Identify only Child 02 dead work.
    2. Remove it and obtain Supervisor cleanup review.
  - Implementation Gate: Pn-A reviewed the result.
  - Current State: cleanup executor report `reports/coder/rp_coder_260814_011727_by_gpt-5_child02_pnb_cleanup.md` (SHA-256 `C77A7C3F59EE26CB1841FBFB650CF3C3081B2748CC254025EF22859A14956B18`) inventoried the exact lifecycle boundary and deleted only a nested P2-E execution skeleton: one `120`-byte owner marker plus four exact empty directories under `.tmp/.tmp/qa-child02-p2e`. Independent report `reports/Supervisor/rp_supervisor_260814_015043_by_gpt-5_child02_pnb_cleanup.md` (SHA-256 `04475021762338800B6D8100AB5ABC31388930EB09E4082C2D221FF0818ED357`) returned unconditional `PASS` / `ACCEPT_CHILD02_PNB_CLEANUP_AND_HAND_BACK`, findings `0/0`, and no residual same-invariant surface. It independently reconciled `77` current-chain paths and `92` lifecycle entries; retained `49/49` accepted artifacts, `5/5` created product tests, all rejected-attempt/history provenance, and all three shared roots; verified JSON `8/8`, PNG `12/12`, five protected hashes, exact five-path absence, the retained empty generic parent, branch/HEAD, and staged `0`. No build/test/runtime/QA/graph gate was rerun because Pn-B changed no tracked product surface. The exact five-document/two-report boundary commits immediately after this row; Pn-C and Child 03 remain closed until Git confirms success.
  - Acceptance: cleanup proof recorded as `E3-PNB-CLEANUP1`.
  - Commit Boundary: isolated documentation/report cleanup commit; Git reports the final hash because a commit cannot contain its own hash.

- [ ] Pn-C: Close the plan and hand off accepted persistence facts.
  - Goal: finalize ledgers, detect-changes, commits, and Child 03 actual-status refresh.
  - Work Steps:
    1. Record final measured values and validation state.
    2. Run final detect-changes, record commit/worktree, and refresh Child 03 from exact Child 02 evidence.
  - Implementation Gate: Pn-A/Pn-B pass and every P2 evidence row is complete.
  - Acceptance: `E3-PNC-DETECT1`, `E3-PNC-HANDOFF1`, and `E3-PNC-COMMIT1` recorded.

## Risk Notes

- Persistence and query owners can have HIGH/CRITICAL blast radius; source-proven field flow keeps the edit set bounded.
- Matching record counts can hide field drift; compare corrected fields and endpoints, not totals alone.
- A broad reader matrix can manufacture work unrelated to graph accuracy; P2-A's affected denominator is authoritative.
- Changing analyze behavior before measuring the existing repeated command can introduce a new defect unrelated to the five campaign findings.
- Silent deduplication in serialization/load can recreate the identity loss after Child 01 has fixed construction.
- Opaque identifiers must not be parsed to reconstruct fields, but a consumer is changed only when P2-A proves that it performs such affected parsing.
- Target state and the excluded scanner defect remain separate boundaries.
