# Anvien Graph Accuracy Roadmap

Date: 2026-07-26
Last revised: 2026-08-24
Status: active campaign; Child 01 P0/P1-A/P1-B/P1-C/P1-D/P1-E/Pn-A/Pn-B/Pn-C accepted at isolated commit boundaries; Child 02 P0-A/P2-A/P2-B/P2-C/P2-D/P2-E/Pn-A/Pn-B/Pn-C accepted at isolated commit boundaries and closed at `181b8cb800f5fe34fa6fe85ddd359f514ead9fb0`; Child 03 P0-A through P3-B2A retain their accepted isolated boundaries; P3-C closes at `656a0445ff3e25b6225b994cdaf7cf1b35eb665c`; P3-C2 closes at `8784c6c21da842b188f136b95ec97ab8df9f20e8`; aggregate `E3-PNA-REVIEW1` and cleanup `E3-PNB-REVIEW1` are `PASS`; `E3-PNB-COMMIT1` is `0231d7f9cc7cfd05d8d5da4787161a8ec2cd1550`; Child 03 Pn-C closes at `3e25f9ca75c9cb7d59bf228e6cc9aa0b81d738b4`; Child 04 P0-A/P4-A/P4-B/P4-B1/P4-C/P4-C2 retain their accepted isolated commits; P4-C2 closes at `03f09b43f652b9a14b3e49774dc805c0dfd24a27`; aggregate `E4-PNA-REVIEW1` and cleanup `E4-PNB-REVIEW1` are `PASS`; `E4-PNB-DETECT1` is `PASS` and `E4-PNB-COMMIT1` closes at `d1d8eb9002ce9c449c3713de0837ac8216d17a8d`; Child 04 Pn-C closes only the Child 04 plan by the current exact five-document commit boundary, with final commit hash reported externally; Child 05 transition occurs after that commit as the next campaign step

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
| 06 | ambient/external resolution, diagnostics, and end-to-end analyze runtime performance | 7 | evidence-backed declaration authority, structured outcomes, `3/3` bounded ambient sites, and faster end-to-end `anviens analyze` runtime with exact accuracy, semantic, graph, determinism, freshness, failure/publication, and persistence/reader parity |
| 07 | cross-surface acceptance and target validation | 3 | deterministic, parity-valid, runtime-validated closure of all five oracles |
| **Total** |  | **36** | **seven accepted Child closures** |

## Child plan inventory

| Child | Plan | Scope | Status |
|-------|------|-------|--------|
| 01 | [graph identity contract and strict construction](2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-plan.md) | identity, range, lexical owner, occurrence conservation, proven collision owner | P0/P1-A/P1-B/P1-C/P1-D/P1-E/Pn-A/Pn-B/Pn-C accepted at isolated commit boundaries |
| 02 | [current graph persistence and reader consistency](2026-07-28-02-current-graph-persistence-and-reader-consistency/2026-07-28-02-current-graph-persistence-and-reader-consistency-plan.md) | persistence parity, affected readers, repeated normal analyze | P0-A/P2-A/P2-B/P2-C/P2-D/P2-E/Pn-A/Pn-B/Pn-C accepted at isolated commit boundaries; Pn-B commit `9b65f3ef3f3f377d0dcb4a9e0cd5eca444cf51c6`; Pn-C closure commit `181b8cb800f5fe34fa6fe85ddd359f514ead9fb0` |
| 03 | [TypeScript binding-pattern extraction](2026-07-28-03-typescript-binding-pattern-extraction/2026-07-28-03-typescript-binding-pattern-extraction-plan.md) | binding-pattern defect only | all seven slices, aggregate review, cleanup, Pn-B commit, and Pn-C handoff/detect are closed at isolated Pn-C commit `3e25f9ca75c9cb7d59bf228e6cc9aa0b81d738b4` |
| 04 | [TypeScript export semantics](2026-07-28-04-typescript-export-semantics/2026-07-28-04-typescript-export-semantics-plan.md) | export-fact defect only | P0-A/P4-A/P4-B/P4-B1/P4-C retain their recorded isolated commits; P4-C2 closes at `03f09b43f652b9a14b3e49774dc805c0dfd24a27`; aggregate Pn-A and cleanup review are `PASS`; Pn-B detect/isolated commit closes at `d1d8eb9002ce9c449c3713de0837ac8216d17a8d`; Pn-C closes the plan by the current exact five-document boundary with final hash external |
| 05 | [module export and re-export resolution](2026-07-28-05-module-export-and-reexport-resolution/2026-07-28-05-module-export-and-reexport-resolution-plan.md) | module/export resolution defect only | dependency-blocked until the Child 04 plan-closure commit succeeds; transition and P5-A opening are the next campaign step, outside Child 04 Pn-C |
| 06 | [ambient/external resolution and diagnostics](2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-plan.md) | ambient/external correctness plus end-to-end `anviens analyze` runtime speedup after correctness closure | P6-D independent Supervisor `PASS`: canonical runtime/readers pass, target capability outcomes are `3/3`, in-repository analyzer gaps are `0/3`, target boundary passes, and procedure is clean. Main-owned fresh current-byte detect now passes on graph `2,170/765/0`, `122,649/169,300`, with summary LOW, file layer HIGH, `13` changed / `8` affected tracked files, ResolutionGap `245/245`, health `0/0/0`, and complete semantic fields. P6-D remains unchecked only for the isolated commit; P6-E remains exactly one U1-U4 slice and stays locked until that commit |
| 07 | [cross-surface acceptance and target validation](2026-07-28-07-cross-surface-acceptance-and-target-validation/2026-07-28-07-cross-surface-acceptance-and-target-validation-plan.md) | acceptance only; no production repair | dependency-blocked |

## Standard file inventory

Each Child owns one plan, one evidence ledger, one benchmark ledger, and one actual-status ledger. Together with this roadmap, the contract, and the four-file correction plan, the active authority contains 34 documents. P6-E is added inside the existing Child 06 four-file authority, so the active-document count remains 34.

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
6. Child 06 supplies evidence-backed ambient/external authority and structured resolution outcomes, projects them truthfully to graph health, recovers the canonical current-runtime path from the durable LadybugDB bundle plus the checked-in standard-vendor/CGO-overlay/one-patch authority under unique lane-owned derived roots, then executes exactly one P6-E slice containing ordered units U1-U4 while preserving accuracy, semantic completeness, graph correctness, deterministic output, freshness, failure/publication behavior, and persistence/reader parity.
7. Child 07 runs final acceptance; each failure is returned to its owning Child.

Each non-terminal Child closes only after its ledgers identify the accepted output contract and refresh the successor's actual status. A handoff communicates facts and evidence; it does not force the successor to preserve an implementation detail that its own source audit disproves.

## Shared artifact ownership

| Artifact or authority | Owner | Consumer rule |
|-----------------------|-------|---------------|
| `docs/contracts/graph-accuracy-contract.md` | Child 01 P1-A | all Children use the accepted invariants |
| original problem report and target oracle | immutable evidence source | findings and target counts only; DRAFT design remains non-authoritative |
| current-graph reader impact inventory | Child 02 P2-A | contains only source-proven affected readers and exact evidence |
| `reports/QA/child04-p4c2/oracle/<oracle_id>/` source-only oracle bundle | clean-context P4-C2 Oracle Authoring lane; Main routes seal identities | born durable; exactly 21 positive + 11 negative rows; QA consumes read-only after seal verification; analyzer/QA output cannot author expected values |
| module export table/result contract | Child 05 | Child 06 inspects; Child 07 validates |
| structured external-resolution outcome | Child 06 | Child 07 validates |
| `third_party/ladybugdb/v0.19.1/windows-x86_64` native input bundle | Child 06 P6-D | repository-owned durable `{lbug.h, lbug_shared.lib, lbug_shared.dll}` authority; `.tmp\ladybug-native` is not an accepted input, cache, recovery source, or reproducibility dependency |
| root `vendor/` plus `third_party/go-vendor` closure authority | Child 06 P6-D | deterministic standard baseline, checksum-derived complete-root CGO overlay, exactly one guarded reviewed patch, manifest contract v2, licenses `45/45`, and dual-shell verifier; cache/network fallback is forbidden |
| campaign order and closure state | this roadmap | each Child updates only its own status/handoff row |

## Target and repository boundary

- Production and plan work occur in `E:\Anvien`.
- `E:\cheapapp.org` is source-inspected read-only by the clean P4-C2 Oracle Authoring lane, then analyzed in place only by the validation slices that name it after the durable oracle seal verifies.
- Normal operational output for the target remains under its own `.anvien` directory.
- Plan, report, fixture, QA, and Supervisor artifacts remain in `E:\Anvien`.
- Every evidence-bearing P4-C2 artifact must originate directly under `reports/QA/child04-p4c2/...`; `.tmp` is disposable debug-only and no artifact born there may close an evidence ID or be restored/promoted/copied/renamed as evidence.
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
- the seven Child plans declare exactly 36 implementation slices in the phase-map counts above;
- every checklist slice has exact evidence IDs declared in its own evidence ledger;
- no evidence row remains for a removed slice;
- actual-status P0 checkboxes match their pending gates and `Related File Count` values count files;
- each Child contains only its assigned semantic responsibility;
- no document treats the initial analyze implementation as fixed or invents unproven write/reader behavior;
- no active document treats the DRAFT architecture recommendations as approved source authority;
- no active document treats `.tmp` as an oracle/evidence path, recovery source, promotion route, or reproducibility dependency;
- P6-D's canonical `scripts/full-build.ps1` path and every affected helper/package/launcher call site consume the repository-owned LadybugDB bundle and explicit unique lane-owned cache/runtime roots; no `.tmp\ladybug-native` shared authority remains;
- P6-D's Go source authority proves standard baseline `2/2`, machine-derived supplemental fixed-point `2/2`, reviewed patch cardinality exactly `1`, final tree `2/2`, unchanged `go.mod`/`go.sum`, and offline compile closure before target access;
- P6-E remains one checklist slice and one commit boundary; U1-U4 are ordered internal units, never separate slices or commits;
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

## Child 04 P4-A Acceptance Checkpoint

- `E4-P4A-IMPACT1` retains the exact four production-owner blast radius: `facts.go`, `kinds.go`, and `ir.go` are CRITICAL; `sort_keys.go` is HIGH. These are scope warnings; the accepted candidate remains limited to those four owners plus `scopeir_test.go` and `scopeir.golden.json` after production code.
- `E4-P4A-SRC1` is the exact six-file candidate, `570` insertions / `0` deletions: one source-site `ExportFact`, seven source-form kinds, three meaning lanes, explicit type-only/provenance/diagnostic state, deep-copy normalization, deterministic JSON ordering, unchanged `DefinitionFact.Visibility`, and no AST extraction, projection, compatibility writer, terminal resolution, barrel, ambiguity, cycle, or public-API state.
- `E4-P4A-BUILD1`, `E4-P4A-TEST1`, and `E4-P4A-BOUNDARY1` are PASS: canonical `npm run full-build` exit `0`; exact focused ScopeIR matrix `7/7`; `go test ./internal/scopeir -count=1` exit `0`; nearest six-package product boundary exit `0`. Repo-wide `go test ./... -count=1` remains exit `1` from the recorded out-of-slice baseline and is not PASS evidence.
- `E4-P4A-REVIEW1` is independent `PASS`: `reports/Supervisor/rp_supervisor_260820_221058_by_gpt-5_child04_p4a_export_fact_boundary.md`, `16,898` bytes / `156` LF / SHA-256 `1B8DEB2F8D5F49F285BE5AA4DF817304F8A9D8DE61E112BD2FCFEECC175573B2`.
- Closure graph for `E4-P4A-DETECT1` is fresh and excluded: `1,130/626/0` scanned/parsed/failed and `81,132/120,514` nodes/relationships. `anvien detect-changes --repo E:\Anvien --scope all` exits `0`: `196` changed semantic units, `11` tracked changed files, `10` affected files, `3` affected processes, overall MEDIUM risk, complete semantic fields, `0` persisted resolution gaps, `0` nodes with gaps, and `0` degraded nodes. Changed paths equal the six candidate plus five living documents; report provenance is verified separately and included only at exact staging.
- `E4-P4A-DETECT1` is PASS; `E4-P4A-COMMIT1` remains pending. P4-B remains locked until the exact accepted implementation/tests/golden, valid reports, and five refreshed living documents are committed together; no push is authorized.

## Campaign completion definition

The campaign is complete only when all 36 implementation slices have their required source/build/runtime/behavior evidence and isolated commits; all five target oracles pass; corrected-field Graph JSON/Ladybug and affected-reader parity pass; repeated analyze is deterministic for the accepted facts; P6-E proves faster end-to-end `anviens analyze` runtime against a comparable baseline while preserving accuracy, semantic completeness, graph correctness, deterministic output, freshness, failure/publication behavior, and persistence/native-reader parity; every Child cleanup and Supervisor gate passes; and this roadmap records the final accepted handoffs with no blocker.

## Child 04 P4-B Acceptance Checkpoint

- `E4-P4B-REVIEW1` is `PASS` on the resubmission report `reports/Supervisor/rp_supervisor_260821_001554_by_gpt-5_child04_p4b_export_facts_resubmission.md`; the prior `REJECT` is retained as historical evidence and its sole artifact-lifecycle finding is closed.
- Exact candidate remains `internal/providers/tsjs/imports.go`, `internal/providers/tsjs/extract.go`, and `internal/providers/tsjs/extract_test.go`; source/test diff is `1,186` insertions / `10` deletions, with no ScopeIR, graph, persistence, target, or forbidden-tree bytes.
- Fresh cleanup evidence proves `.tmp/p4b_ast_probe/` and all matching probe directories are absent; focused export tests and the nearest three-package boundary pass; canonical `npm run full-build` exit `0` remains valid from independent REVIEW1 evidence.
- `E4-P4B-DETECT1` and `E4-P4B-COMMIT1` close the documentation lag: final detect passed before isolated commit `11a37aa8ec0320dd93258c058b088d1070aa778d`; no P4-B source was reopened. P4-B1 is now closed at `42d167aaf28446ac0b3de479a8afefabb8d06736`; P4-C, P4-C2, Child 05, and target access remain locked.

## Child 04 P4-B1 REVIEW3 Acceptance Checkpoint

- `E4-P4B1-REVIEW1` is independent `PASS` on `reports/Supervisor/rp_supervisor_260821_031004_by_gpt-5_child04_p4b1_comment_recovery_review3.md`; identity `11,830` bytes / `136` LF / SHA-256 `07DD5BB92F169C5923C0DBCB597F914A28E594496ACDADB55D41F13DE364421C`.
- Exact candidate remains only `internal/providers/tsjs/imports.go` and `internal/providers/tsjs/extract_test.go`; current hashes are `52AFBA2FC9A7ACEA314D5B39043A054F2671B2EF987161D8F1C3D641B2382749` and `07CF7D49715CA0398DA9485086A37E395FCB8E2E695AF2B649B6BC05074C604D`; diff is `736` insertions / `24` deletions; index is empty.
- Fresh canonical build, focused `6/6`, full `tsjs`, nearest `3/3`, and `resolution/analyze 2/2` pass. Independent TS/JS parser-to-ScopeIR probes close both comment-bearing forms at `2/1/2`, with no `Broken` fact/import and zero terminal state. Fresh excluded graph is `1,144/626/0`, `82,059/121,760` nodes/relationships.
- Main-owned `E4-P4B1-DETECT1` is recorded as PASS from the fresh excluded graph/change detection (`305` changed semantic units, `7` changed/affected files, `1` affected process, MEDIUM risk, zero persisted resolution-health degradation), and `E4-P4B1-COMMIT1` closes at isolated commit `42d167aaf28446ac0b3de479a8afefabb8d06736`; current successor basis is clean HEAD `871189b8c6a4e4bb9ff538407232c913b8cf4db6`, with the two external docs-only commits preserved and no push. Explicit successor authority opens only P4-C under `E4-P4C-AUTH1`; P4-C2/Child 05 and `E:\cheapapp.org` remain locked.

## Child 04 P4-C Supervisor REVIEW1 Acceptance Checkpoint

- `E4-P4C-REVIEW1` is `PASS` at `reports/Supervisor/rp_supervisor_260821_083556_by_gpt-5_child04_p4c_graph_persistence_review1.md`; final identity is `15,261` bytes / `101` LF / SHA-256 `AA94C417A02BE371168D8ADCD5B3F4FECF1941F382518EFF9214EC0FABB93CFF`.
- The review independently confirms `414` Graph Export nodes and `414` File→Export relations, `0` duplicate IDs, `0` orphan local-definition references, `0` forbidden terminal/resolved-target/public-API keys, direct runtime records `240`, runtime local records `239`, compatibility drift `0`, and `28 × 414 = 11,592` normalized Graph JSON↔Ladybug field comparisons with `0` differences and `0` missing IDs in either direction.
- Source, focused owner tests, canonical full build (`1,855/735/0`; `113,496/156,003` graph), CSV/schema/loader, semantic reader, negative controls, and HIGH/CRITICAL impact warnings are closed within P4-C. `E4-P4C-DETECT1` is now `PASS` on fresh graph `1,857/735/0`, `113,523/156,030`, with `180` changed units, `16` changed files, `14` affected files, HIGH risk and resolution health `0/0/0`; `E4-P4C-COMMIT1` remains pending and is Main-owned.
- Review provenance records the empty review-induced parent `.tmp\\p4c-tests` and the prior Supervisor protocol deviation that briefly opened `/root/authority_scan`; it was read-only, created no state, and did not invalidate technical review independence. No target access occurred.

## Child 04 P4-C Commit Checkpoint

- `E4-P4C-COMMIT1` closes at isolated commit `c99c4070b66e7a96be8c9fa2721a0335a1f94877` (`feat(graph): project TypeScript export facts`), with the accepted 23-path boundary, post-commit HEAD match, `git diff --check` PASS, no push/reset/checkout, and no target access.
- The commit contains the five refreshed Child 04 living ledgers, four production owners, four focused tests, seven owner test/golden updates, Coder/Supervisor reports, and the current `0819` Main handoff. Older `0631`/`0721` handoffs remain preserved untracked because their historical blank EOF lines would violate staged `git diff --check`; their bytes were not rewritten.
- At the P4-C commit checkpoint, P4-C2 became the sole open slice and source-only Oracle Authoring was authorized for the three hash-pinned files before analyzer observation. The later Oracle/QA checkpoint below records that lifecycle as completed; Child 05 and all later slices remain locked.

## Child 04 P4-C2 Oracle Lifecycle Checkpoint

- P4-C2 remains the sole open slice. Source-only Oracle Authoring is authorized to inspect the three hash-pinned target source files while the target worktree remains preserve-only; this is ground-truth derivation, not analyzer validation.
- The clean-context Oracle Authoring lane must not observe target `.anvien`, current Anvien implementation/tests/goldens, analyzer or QA output/reports, Child 05 state, or historical `.tmp` oracle material. It seals exactly 21 positive rows plus 11 owner-qualified negative controls before existing QA resumes.
- Every oracle, raw capture, command stream, manifest, provenance record, expected-value input/output, benchmark, and reproducibility artifact is created directly under `reports/QA/child04-p4c2/...`. `.tmp` is debug-only and can never close an evidence ID or be promoted/restored as evidence.
- Historical `.tmp\cheapapp-graph-root-cause-restart\p1b-identity-oracle-output.json` is an invalid-lifecycle debug capture, not an accepted or lost oracle. The QA oracle-gate and Recovery reports remain historical; their `.tmp` acceptance/recovery route is non-authoritative.
- `E4-P4C2-ORACLE1` is `SEALED`: oracle ID `p4c2-oracle-v1-a869876ab626-260821_110849+0700`, bundle digest `7749AB14E02FBF61CF7F81B3A7638888F8AD414064F9F969CFED06EB11355439`, exactly `21+11` rows, and zero target writes/forbidden observations/evidence-bearing `.tmp` artifacts. The gate remains closed.

## Child 04 P4-C2 QA Retry Checkpoint

- Historical first QA run remains durably `BLOCKED` only on normal-process Git trust; its canonical full build PASS (`1.2.8`, `180.018s`) and preservation evidence remain closed and were not rerun.
- Process-local trust retry preflight and exactly one `anvien analyze E:\cheapapp.org --force` PASS: exit `0`, `77.37s`, `1,359/887/0`, graph `94,422/125,299`; Graph JSON/Ladybug sizes are `432,028,037` / `150,351,872` bytes.
- QA report `reports/QA/child04-p4c2/runs/p4c2-target-validation-retry-260821_115050+0700/p4c2-qa-retry-validation-report.md` is `11,342` bytes / `200` LF / SHA-256 `C831004F049A563A2387B599BE01C943F5B9416C72B1C45E50A8C1F9D2CEFDB4`; manifest run digest is `9F414A2C54C42F4E39AD8ED03DC042CCC3E1FB5993DA842B22F64851D16AABC4`.
- QA is `READY_FOR_SUPERVISOR`, not an acceptance verdict: all `21/21` Export facts exist; `6/21` positive rows pass overall; `P001`–`P014` and `P018` fail only because exported TypeAlias definitions have `isExported=false` and FileContext `exported=false`; six Function rows and `11/11` negative controls pass.
- Graph JSON↔Ladybug parity passes at `588/588` field comparisons with `0` differences, while semantic correctness fails for the same 15 compatibility values. Duplicate IDs, orphan endpoints/local definitions, export diagnostics, and forbidden Child 05 state are all `0`.
- Target HEAD/source/status and four Git-config identities are preserved; only normal analyzer-owned `.anvien` output changed. Independent `E4-P4C2-REVIEW1` is `REJECT`: `reports/Supervisor/rp_supervisor_260821_123030_by_gpt-5_child04_p4c2_durable_retry_review1.md`, `15,671` bytes / `126` LF / canonical SHA-256 `8E37F4B126ABB38A5DCE071C35937F59D9152EFA135B9245408CA138BEC781F7`.
- The rejected invariant is limited to definition-level direct-export compatibility and the downstream FileContext result for `P001`–`P014`,`P018`. Repair must derive direct source-export membership independently from runtime-value eligibility, preserve `typeOnly`/meaning/access separation, and keep negatives/parity/Child 05 exclusions intact. Child 05 remains locked.

## Child 04 P4-C2 Rejection Repair and REVIEW2 Checkpoint

- The exact two-file production-first repair passes canonical full build, focused rejection regression, resolution→FileContext→Ladybug boundary and all four affected packages. Fresh target QA runs exactly one normal analyze (`1,359/887/0`, graph `94,422/125,299`) and passes `21/21` positives, `11/11` negatives, FileContext `17/3/1`, Graph JSON↔Ladybug `588/0`, all integrity/Child 05 zeros and target/config preservation.
- Independent `E4-P4C2-REVIEW2` is `PASS`: `reports/Supervisor/rp_supervisor_260821_133927_by_gpt-5_child04_p4c2_typealias_compatibility_review2.md`, `13,650` bytes / `120` LF / canonical SHA-256 `5B99A74B1A8D91D48F5E62F0BA1FFCB26317BF818AC6AE044E6CD650B208DC0B`. P4-C2 remains open only for Main-owned fresh graph, detect-changes and isolated commit; Child 05 stays locked.
- `E4-P4C2-DETECT1` is `PASS` on fresh self graph `1,939/736/0`, `114,628/157,443`: `50` changed semantic units, exact `7/7` changed/affected tracked files, changed-file risk `HIGH`, overall risk `LOW`, `0` affected processes/flows, persisted resolution health `0/0/0`, and complete semantic fields. Only exact staging and `E4-P4C2-COMMIT1` remain open; Child 05 stays locked.
- `E4-P4C2-COMMIT1` closes at isolated commit `03f09b43f652b9a14b3e49774dc805c0dfd24a27`, parent `310502a88849fe75f86a45a987ba21490d19dbe2`, with exact `89` paths. Post-commit index/tracked worktree are clean, the two P4-C handoffs `0631`/`0721` remain protected untracked, repair identities are unchanged, and no push occurred. Aggregate Child 04 Pn-A opens; Pn-B/Pn-C and Child 05 stay locked.

## Child 04 Aggregate Pn-A Acceptance Checkpoint

- `E4-PNA-REVIEW1` is independent `PASS`: `reports/Supervisor/rp_supervisor_260821_142429_by_gpt-5_child04_pna_aggregate_review1.md`, `23,563` bytes / `178` LF / raw SHA-256 `634555CB9B6917A37EA2826302BC4C8949920B754307EF661C1D4DA3EBC277BC` / canonical SHA-256 `7EBFD5087F8593660A94E70B0816A7FC98944FDE7B9D1F1BC9388CEB9F6DC5A8`.
- Aggregate review independently clears all production groups, commit ancestry, sealed Oracle/QA manifests, `21/21 + 11/11`, TypeAlias `15/15`, parity `588/0`, integrity/Child 05 zeros and target preservation without rerunning closed gates. Residual same-invariant surfaces are none.
- At the aggregate checkpoint, Pn-B opened while Pn-C and Child 05 remained locked; the cleanup checkpoint below records the subsequent acceptance.

## Child 04 Pn-B Cleanup Acceptance Checkpoint

- `E4-PNB-CLEAN1` is `READY_FOR_SUPERVISOR` at `reports/coder/rp_coder_260821_144325_by_gpt-5_child04_pnb_cleanup_ready_for_supervisor.md`: `24,399` bytes / `472` LF / raw SHA-256 `0209C39BE833312100DFA3948B9676A8AC091A40286F94F9E2E1220B3278839C` / canonical SHA-256 `5BD0338A8949B58933988FAA6DF448EFBF5B4F4D506C91D6DD7B5B44B8F7B260`.
- The executor removed only `.tmp/p4c-tests`, an empty review-induced directory with `0` files / `0` bytes. It preserved the `136` tracked Child 04 paths, three Main handoffs, sealed Oracle, all three QA generations, every accepted/historical report, and all unrelated `.tmp` families.
- Independent `E4-PNB-REVIEW1` is `PASS`: `reports/Supervisor/rp_supervisor_260821_150355_by_gpt-5_child04_pnb_cleanup_review1.md`, `14,130` bytes / `160` LF / raw SHA-256 `F114AF17513B56952B81B71110BA1C4D838AD40CE3D4E6049D4AD7C2D0ABD18F` / canonical SHA-256 `EDFCF6CACA23DE0F8F38BA4376A25009B33B525BDAF70D267D062155AEC91E1F`; residual same-invariant surfaces are none.
- Planner refresh basis is the fresh self graph `1,943/736/0`, `114,721/157,536`. The roadmap is LOW with `28` outbound links; each Child 04 ledger is LOW with one inbound link; all five upstream impacts are `0` affected files/flows/tests.
- `E4-PNB-DETECT1` is `PASS`: `19` changed documentation sections, exactly `5/5` changed/affected files, LOW changed-file and overall risk, `0` affected processes/flows, ResolutionGap delta `0/0`, persisted health `0/0/0`, and complete semantic fields.
- `E4-PNB-COMMIT1` closes at `d1d8eb9002ce9c449c3713de0837ac8216d17a8d`, parent `c7997886a0faeb32b7cfe05b4f7d08e38fc57228`, with the exact eight-path cleanup/report/ledger/provenance manifest. Post-commit index and tracked worktree are clean; no push/reset/checkout occurred. Pn-C is now the sole open gate.

## Child 04 Pn-C Plan Closure

- `E4-PNC-CLOSE1` declares the Child 04 plan `CLOSED`: all five P4 slices retain accepted isolated commits, aggregate and cleanup reviews are `PASS`, and Pn-B is committed at `d1d8eb9002ce9c449c3713de0837ac8216d17a8d`. No Child 04 gate remains open.
- Pn-C changes only this roadmap and the four Child 04 ledgers. It does not refresh Child 05, author a handoff, open P5-A, rerun any accepted gate, or add a separate closure report.
- `E4-PNC-COMMIT1` is the exact isolated five-document plan-closure boundary. Its final hash and clean post-state are reported externally after Git succeeds; no push occurs.
