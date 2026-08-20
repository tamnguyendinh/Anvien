# Anvien Graph Accuracy Roadmap

Date: 2026-07-26
Last revised: 2026-08-20
Status: active campaign; Child 01 P0/P1-A/P1-B/P1-C/P1-D/P1-E/Pn-A/Pn-B/Pn-C accepted at isolated commit boundaries; Child 02 P0-A/P2-A/P2-B/P2-C/P2-D/P2-E/Pn-A/Pn-B/Pn-C accepted at isolated commit boundaries and closed at `181b8cb800f5fe34fa6fe85ddd359f514ead9fb0`; Child 03 P0-A through P3-B2A retain their accepted isolated boundaries; P3-C closes at `656a0445ff3e25b6225b994cdaf7cf1b35eb665c`; P3-C2 closes at `8784c6c21da842b188f136b95ec97ab8df9f20e8`; aggregate `E3-PNA-REVIEW1` and cleanup `E3-PNB-REVIEW1` are `PASS`; `E3-PNB-COMMIT1` is `0231d7f9cc7cfd05d8d5da4787161a8ec2cd1550`; Child 03 Pn-C closes at `3e25f9ca75c9cb7d59bf228e6cc9aa0b81d738b4`; Child 04 P0-A current-state discovery is complete by the current isolated five-document boundary, and P4-A opens automatically after Git reports that boundary committed

Multi-plan root: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/`

Correction plan: `docs/plans/2026-08-09-anvien-graph-accuracy-multi-plan-correction/2026-08-09-anvien-graph-accuracy-multi-plan-correction-plan.md`

Contract: `docs/contracts/graph-accuracy-contract.md`

## Goal

Generate a graph that is more correct and precise than the measured baseline graph by closing the five bounded defect families established by the original problem report:

1. distinct same-name declarations collapse;
2. TypeScript binding-pattern leaves are omitted;
3. TypeScript export facts are missing or drift across representations;
4. barrel imports do not resolve through re-export bindings to terminal symbols;
5. TypeScript standard-library sites are reported as in-repository gaps instead of correct external or explicit external-capability outcomes.

The report at `reports/problem/2026-07-26-tinh-chinh-cac-van-de-cua-anvien` is the problem origin and oracle source. Its DRAFT architecture proposals are not automatically accepted. The causal synthesis and bounded Supervisor PASS verify the findings only.

## Execution authority

- All implementation work occurs directly in the production codebase at the repository root.
- Commands retain their normal grammar and run against the normally built runtime.
- The analyze pipeline present when a slice opens is baseline evidence. It may be preserved or changed according to current source, file-detail, impact, and behavior evidence; it is not frozen by this roadmap.
- Every Child works on the single production implementation and normal command path; comparison uses independent evidence from that same path.
- Git provides one commit and rollback boundary per accepted implementation slice; it is not a second runtime or semantic oracle.
- Production code is corrected before tests are updated. Full build precedes real-boundary validation.
- Every Child keeps its actual-status, evidence, and benchmark ledgers current. Every implementation slice requires the relevant impact evidence, behavior proof, Supervisor PASS, detect-changes, and its own commit.
- A later acceptance failure returns to the exact owning Child and slice. Validation work does not absorb production repair.

## Campaign phase map

| Child | Responsibility | Implementation slices | Required result |
|-------|----------------|----------------------:|-----------------|
| 01 | graph identity contract and strict construction | 5 | distinct declaration/symbol identity, correct ranges and lexical owner, no silent collision loss, `4/4` target oracle |
| 02 | current graph persistence and reader consistency | 5 | corrected facts preserved through Graph JSON/Ladybug and source-proven affected readers; repeated normal analyze verified |
| 03 | TypeScript binding-pattern extraction | 7 | recursive declaration bindings and `6/6` bounded downstream correctness |
| 04 | TypeScript export semantics | 5 | direct/default/alias/type-only/star/namespace/re-export facts and `21/21` direct exports |
| 05 | module export and re-export resolution | 4 | export table and terminal resolution with `2/2` bounded calls |
| 06 | ambient/external resolution and diagnostics | 6 | evidence-backed declaration authority, structured outcomes, and `3/3` bounded ambient sites |
| 07 | cross-surface acceptance and target validation | 3 | deterministic, parity-valid, runtime-validated closure of all five oracles |
| **Total** |  | **35** | **seven accepted Child closures** |

## Child plan inventory

| Child | Plan | Scope | Status |
|-------|------|-------|--------|
| 01 | [graph identity contract and strict construction](2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-plan.md) | identity, range, lexical owner, occurrence conservation, proven collision owner | P0/P1-A/P1-B/P1-C/P1-D/P1-E/Pn-A/Pn-B/Pn-C accepted at isolated commit boundaries |
| 02 | [current graph persistence and reader consistency](2026-07-28-02-current-graph-persistence-and-reader-consistency/2026-07-28-02-current-graph-persistence-and-reader-consistency-plan.md) | persistence parity, affected readers, repeated normal analyze | P0-A/P2-A/P2-B/P2-C/P2-D/P2-E/Pn-A/Pn-B/Pn-C accepted at isolated commit boundaries; Pn-B commit `9b65f3ef3f3f377d0dcb4a9e0cd5eca444cf51c6`; Pn-C closure commit `181b8cb800f5fe34fa6fe85ddd359f514ead9fb0` |
| 03 | [TypeScript binding-pattern extraction](2026-07-28-03-typescript-binding-pattern-extraction/2026-07-28-03-typescript-binding-pattern-extraction-plan.md) | binding-pattern defect only | all seven slices, aggregate review, cleanup, Pn-B commit, and Pn-C handoff/detect are closed at isolated Pn-C commit `3e25f9ca75c9cb7d59bf228e6cc9aa0b81d738b4` |
| 04 | [TypeScript export semantics](2026-07-28-04-typescript-export-semantics/2026-07-28-04-typescript-export-semantics-plan.md) | export-fact defect only | predecessor handoff accepted; P0-A current-source/file-detail/impact discovery complete by the current docs-only boundary; P4-A opens only after that boundary commits |
| 05 | [module export and re-export resolution](2026-07-28-05-module-export-and-reexport-resolution/2026-07-28-05-module-export-and-reexport-resolution-plan.md) | module/export resolution defect only | dependency-blocked |
| 06 | [ambient/external resolution and diagnostics](2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-plan.md) | ambient/external defect only | dependency-blocked |
| 07 | [cross-surface acceptance and target validation](2026-07-28-07-cross-surface-acceptance-and-target-validation/2026-07-28-07-cross-surface-acceptance-and-target-validation-plan.md) | acceptance only; no production repair | dependency-blocked |

## Standard file inventory

Each Child owns one plan, one evidence ledger, one benchmark ledger, and one actual-status ledger. Together with this roadmap, the contract, and the four-file correction plan, the active authority contains 34 documents.

| Child | Plan | Evidence | Benchmark | Actual status |
|-------|------|----------|-----------|---------------|
| 01 | [plan](2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-plan.md) | [evidence](2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-evidence.md) | [benchmark](2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-benchmark.md) | [actual status](2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-actual-status.md) |
| 02 | [plan](2026-07-28-02-current-graph-persistence-and-reader-consistency/2026-07-28-02-current-graph-persistence-and-reader-consistency-plan.md) | [evidence](2026-07-28-02-current-graph-persistence-and-reader-consistency/2026-07-28-02-current-graph-persistence-and-reader-consistency-evidence.md) | [benchmark](2026-07-28-02-current-graph-persistence-and-reader-consistency/2026-07-28-02-current-graph-persistence-and-reader-consistency-benchmark.md) | [actual status](2026-07-28-02-current-graph-persistence-and-reader-consistency/2026-07-28-02-current-graph-persistence-and-reader-consistency-actual-status.md) |
| 03 | [plan](2026-07-28-03-typescript-binding-pattern-extraction/2026-07-28-03-typescript-binding-pattern-extraction-plan.md) | [evidence](2026-07-28-03-typescript-binding-pattern-extraction/2026-07-28-03-typescript-binding-pattern-extraction-evidence.md) | [benchmark](2026-07-28-03-typescript-binding-pattern-extraction/2026-07-28-03-typescript-binding-pattern-extraction-benchmark.md) | [actual status](2026-07-28-03-typescript-binding-pattern-extraction/2026-07-28-03-typescript-binding-pattern-extraction-actual-status.md) |
| 04 | [plan](2026-07-28-04-typescript-export-semantics/2026-07-28-04-typescript-export-semantics-plan.md) | [evidence](2026-07-28-04-typescript-export-semantics/2026-07-28-04-typescript-export-semantics-evidence.md) | [benchmark](2026-07-28-04-typescript-export-semantics/2026-07-28-04-typescript-export-semantics-benchmark.md) | [actual status](2026-07-28-04-typescript-export-semantics/2026-07-28-04-typescript-export-semantics-actual-status.md) |
| 05 | [plan](2026-07-28-05-module-export-and-reexport-resolution/2026-07-28-05-module-export-and-reexport-resolution-plan.md) | [evidence](2026-07-28-05-module-export-and-reexport-resolution/2026-07-28-05-module-export-and-reexport-resolution-evidence.md) | [benchmark](2026-07-28-05-module-export-and-reexport-resolution/2026-07-28-05-module-export-and-reexport-resolution-benchmark.md) | [actual status](2026-07-28-05-module-export-and-reexport-resolution/2026-07-28-05-module-export-and-reexport-resolution-actual-status.md) |
| 06 | [plan](2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-plan.md) | [evidence](2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-evidence.md) | [benchmark](2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-benchmark.md) | [actual status](2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-actual-status.md) |
| 07 | [plan](2026-07-28-07-cross-surface-acceptance-and-target-validation/2026-07-28-07-cross-surface-acceptance-and-target-validation-plan.md) | [evidence](2026-07-28-07-cross-surface-acceptance-and-target-validation/2026-07-28-07-cross-surface-acceptance-and-target-validation-evidence.md) | [benchmark](2026-07-28-07-cross-surface-acceptance-and-target-validation/2026-07-28-07-cross-surface-acceptance-and-target-validation-benchmark.md) | [actual status](2026-07-28-07-cross-surface-acceptance-and-target-validation/2026-07-28-07-cross-surface-acceptance-and-target-validation-actual-status.md) |

## Ordered implementation and handoffs

1. Child 01 establishes source-backed identity/range rules, fixes the proven collision/loss owners, and proves the same-name oracle.
2. Child 02 discovers which persistence/read paths consume the corrected fields, preserves them across Graph JSON/Ladybug, corrects only affected readers, and validates repeated normal analyze.
3. Child 03 repairs recursive TypeScript binding-pattern extraction.
4. Child 04 records TypeScript module export semantics independently from access visibility.
5. Child 05 resolves module export surfaces and re-export chains to terminal symbols.
6. Child 06 supplies evidence-backed ambient/external authority and structured resolution outcomes, then projects them truthfully to graph health.
7. Child 07 runs final acceptance; each failure is returned to its owning Child.

Each non-terminal Child closes only after its ledgers identify the accepted output contract and refresh the successor's actual status. A handoff communicates facts and evidence; it does not force the successor to preserve an implementation detail that its own source audit disproves.

## Shared artifact ownership

| Artifact or authority | Owner | Consumer rule |
|-----------------------|-------|---------------|
| `docs/contracts/graph-accuracy-contract.md` | Child 01 P1-A | all Children use the accepted invariants |
| original problem report and target oracle | immutable evidence source | findings and target counts only; DRAFT design remains non-authoritative |
| current-graph reader impact inventory | Child 02 P2-A | contains only source-proven affected readers and exact evidence |
| module export table/result contract | Child 05 | Child 06 inspects; Child 07 validates |
| structured external-resolution outcome | Child 06 | Child 07 validates |
| campaign order and closure state | this roadmap | each Child updates only its own status/handoff row |

## Target and repository boundary

- Production and plan work occur in `E:\Anvien`.
- `E:\cheapapp.org` is analyzed in place only by the validation slices that name it.
- Normal operational output for the target remains under its own `.anvien` directory.
- Plan, report, fixture, QA, and Supervisor artifacts remain in `E:\Anvien`.
- The target's pre-existing worktree is preserved.
- The scanner defect involving eight omitted paths is not part of this campaign; the campaign must introduce no additional omission.

## Campaign acceptance matrix

| Defect | Baseline | PASS | Owning Child |
|--------|---------:|-----:|--------------|
| Same-name identity | `2/4` | `4/4` | 01 |
| Binding patterns | `0/6` | `6/6` | 03 |
| Direct exports | `0/21` | `21/21` | 04 |
| Barrel calls | `0/2` | `2/2` | 05 |
| Ambient sites | `0/3` | `3/3` correct external/capability outcomes | 06 |

## Active plan integrity gate

- all 34 active documents resolve to the paths declared by this roadmap;
- the seven Child plans declare exactly 35 implementation slices in the phase-map counts above;
- every checklist slice has exact evidence IDs declared in its own evidence ledger;
- no evidence row remains for a removed slice;
- actual-status P0 checkboxes match their pending gates and `Related File Count` values count files;
- each Child contains only its assigned semantic responsibility;
- no document treats the initial analyze implementation as fixed or invents unproven write/reader behavior;
- no active document treats the DRAFT architecture recommendations as approved source authority;
- link, forbidden-concept, slice, evidence, diff, fresh-analyze, and Supervisor gates all pass before production implementation opens.

## Child 03 Pn-B/Pn-C Closure Checkpoint

- `E3-PNB-CLEAN1` is the cleanup executor's current candidate: `28,632` bytes / `412` LF lines / SHA-256 `D63362A7B382F8382875E71718DDC580B34A21DEBCA71BC327509E56DAC1E8D4`.
- Independent `E3-PNB-REVIEW1` is `PASS`: `34,757` bytes / `455` LF lines / SHA-256 `533A957569BB929FFFD8C269BACAB781C14CB0568B4F465DE67E5D8C81A6943D`; the exact eleven-path absence, `107` denominator, `.tmp=738` census, shared/protected hashes, and no-missed-artifact sweep are closed.
- Planner refresh basis is the excluded graph `1,124/626/0` with `80,879/120,138` nodes/relationships; the roadmap is LOW with `28` outbound links, each Child 03 ledger is LOW with one inbound link, and all five upstream impacts are `0` affected files/processes/flows/tests.
- The Main-owned isolated Pn-B boundary is the five living documents, the cleanup Coder report, the Supervisor report, and the explicitly authorized concurrent Main handoff provenance. No production/test/probe/target/forbidden-tree path is in scope; `E3-PNB-COMMIT1` closes this boundary before Pn-C opens.
- Staged `E3-PNB-DETECT1` is `PASS`: `75` changed documentation/reporting sections, `8` changed files, `8` affected files, LOW risk, `0` affected processes/flows, gap delta `0/0`, and current health `0/0/0`. The exact eight-path manifest remains isolated, and `E3-PNB-COMMIT1` closes the former sole remaining Pn-B gate.

## Child 03 Pn-C Closure and Child 04 Handoff

- `E3-PNB-COMMIT1` is verified at `0231d7f9cc7cfd05d8d5da4787161a8ec2cd1550`, parent `0dd710bb4b0f37072854071058af58bcf9b9e73d`, with the exact eight-path manifest, clean post-commit worktree/index, and no push.
- `E3-PNC-HANDOFF1` records the accepted final Child 03 evidence boundary: all seven slice commits are ancestors, final build/real-boundary evidence and benchmark rows remain accepted, the target six-site boundary is `6/6` with binding-caused gaps `0`, and no source/test/probe/target/forbidden-tree bytes are reopened.
- Child 04 actual status at handoff retained the bounded `21`-export finding as historical evidence and opened only P0-A; the subsequent P0-A checkpoint below independently resolves current source/file-detail/impact/syntax/consumer evidence without reopening Child 03.
- `E3-PNC-DETECT1` is `PASS`: `50` changed documentation/reporting sections, `8` changed files, `8` affected files, LOW/LOW risk, docs `50`, documentation/reporting `30/20`, affected processes/flows `0/0`, gap delta `0/0`, health `0/0/0`, and complete semantic fields; changed and affected path sets equal the exact manifest.
- `E3-PNC-COMMIT1` is the exact eight-path docs/handoff commit `3e25f9ca75c9cb7d59bf228e6cc9aa0b81d738b4`, parent `0231d7f9cc7cfd05d8d5da4787161a8ec2cd1550`; no push occurred.

## Child 04 P0-A Current-State Checkpoint

- `E0-P0A-GRAPH1` records the fresh excluded graph at Child 03 closure HEAD `3e25f9ca75c9cb7d59bf228e6cc9aa0b81d738b4`: `1,126/626/0` scanned/parsed/failed and `80,908/120,167` nodes/relationships.
- `E0-P0A-SRC1` proves there is no first-class export fact, meaning contract, collection, ordering, or export-specific diagnostic. Direct/local/default/type-only syntax emits no export fact; named/star re-exports survive only as partial import compatibility facts; namespace alias and type-only detail are lost.
- `E0-P0A-FD1` and `E0-P0A-IMPACT1` classify the exact four-file P4-A ScopeIR contract boundary, later TSJS extraction owners, graph/Ladybug projection owners, semantic consumers, and preserve-only carriers. CRITICAL/HIGH results are recorded scope warnings, not edit prohibitions.
- `E0-P0A-DETECT1` is staged `PASS`: fresh excluded closure graph `1,126/626/0`, `80,913/120,172`; LOW `29/5/5` changed sections/files/affected files; affected processes/flows, gap delta, and health are all zero; semantic fields are complete.
- P0-A changes only this roadmap and the four Child 04 living ledgers. P4-A opens only after the exact five-document docs-only boundary is committed; no production/test/generated/target path and no push are authorized in this closure.

## Campaign completion definition

The campaign is complete only when all 35 implementation slices have their required source/build/runtime/behavior evidence and isolated commits; all five target oracles pass; corrected-field Graph JSON/Ladybug and affected-reader parity pass; repeated analyze is deterministic for the accepted facts; performance and capacity measurements are recorded where implementation changes them; every Child cleanup and Supervisor gate passes; and this roadmap records the final accepted handoffs with no blocker.
