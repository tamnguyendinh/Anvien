# Child 06A Accelerate Analyze Without Sacrificing Accuracy Plan

## Metadata

- Date: `2026-08-24`
- Status: `P0-A complete / P1-A complete / A001 KEEP / A001_COMMIT_COMPLETE / A002 ARCHITECT_PENDING / PREEXISTING_PRESERVE_ONLY_GOLDEN_FAILURE_RECORDED / repo-specific optimized baselines accepted separately`
- Plan: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy/2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy-plan.md`
- Plan rules: [plan-rules.md](plan-rules.md)
- Evidence: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy/2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy-evidence.md`
- Benchmark: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy/2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy-benchmark.md`
- Actual status: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy/2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy-actual-status.md`
- Roadmap: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-26-anvien-graph-accuracy-roadmap.md`
- Predecessor: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-plan.md`
- Predecessor closure commit: `81163e39718b94a509e41114cada224e8f269e36`
- Successor: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-07-cross-surface-acceptance-and-target-validation/2026-07-28-07-cross-surface-acceptance-and-target-validation-plan.md`
- Method authority: `E:\Anvien\reports\Investigation\rp_main_260824_204615_child06a_optimization_method_discussion.md`
- Required provenance and handoff reference: `E:\Anvien\reports\Investigation\rp_huong_dan_lam_viec_de_slice_p6e_tien_hanh_nhanh_hon.md`

## Goal

Reduce as much as current evidence safely permits the measured graph-generation elapsed time of the real `anvien analyze` pipeline on the same accepted workload, with direct before/after and end-to-end before/after measurements, while preserving every accepted correctness, output, persistence, reader, determinism, freshness, failure, transaction, temporary-file, and publication invariant.

## Rules

- Binding detailed rule authority: [plan-rules.md](plan-rules.md).
- The linked auxiliary file contains the complete Child 06A plan-wide rules. It is not a fifth standard ledger, phase, gate, report, evidence source, numeric control surface, or permission boundary; this plan continues to control goals, checklists, the Current P2-A Attempt, work steps, gates, and acceptance.

## Problem

Child 06 completed the accepted correctness/runtime/reader boundary at P6-D. P1-A now has one accepted real capture and all timing values exposed by the current benchmark output, but untimed graph-finalization/persistence and CLI publication work prevents a complete real-operation timing map and evidence-ranked bottleneck list. Child 06A still has no production implementation, comparable initial-versus-final result, or implementation commit.

The superseded method made a measurement framework and a predetermined solution sequence control the work. That made execution depend on protocol completion rather than on real runtime costs. Child 06A must instead let current absolute measurements select the next operation and require a fresh architecture decision, concrete Planner refresh, implementation, remeasurement, and independent accuracy check for every production attempt.

## Scope

In scope:

- the accepted real `anvien analyze` command and same input workload inherited from the P6-D runtime boundary, with exact options/cache/runtime identity recorded only from P1-A execution;
- initial total graph-generation elapsed time and detailed timings for real internal operations present in the current pipeline;
- a living absolute bottleneck ranking with a work denominator for every measured operation;
- parent-local drill-down measurement inside only the selected top-level bottleneck, producing a complete absolute elapsed-time ordered child list, including smaller measured child costs, before any child Architect attempt;
- a single conditional P1-A instrumentation-writer branch only when existing timing/benchmark/profile output lacks one required timing, with exact edit ownership, pre-edit graph/file-detail/impact, build-before-use, like-for-like comparability/output equivalence, and one carry-forward or remove/rebuild/remeasure disposition;
- per-attempt fresh Visible Architect decision and concrete Planner refresh before production editing;
- sequential production optimization of current measured bottlenecks, with production-first implementation, authorized test edits, full build, post-build test execution, direct and end-to-end remeasurement, per-attempt Visible Supervisor accuracy/equivalence review, disposition, baseline promotion or restoration, and reranking;
- final initial-versus-final graph-generation proof on the same workload;
- one final whole-candidate Supervisor boundary, exact dead-work cleanup, one detect, one implementation commit, and Child 07 handoff.

Preserve/validate boundary:

- accepted accuracy, semantic completeness, graph correctness, ordering of outcomes/diagnostics/evidence, and fail-closed behavior;
- canonical in-memory graph, Graph JSON, Ladybug/native persistence, and affected native/product readers;
- deterministic replay, source freshness/invalidation, failure propagation, transaction rollback, temporary-file flush/close/rename, and publication visibility;
- exact workload identity and operation-specific work denominators needed for comparable before/after results.

## Non-Goals

- No implementation of a historical idea merely because it appeared in a prior report or legacy plan item.
- No predefined bottleneck, source owner, file list, optimization list, or reusable architecture decision.
- No general measurement framework, permanent observability project, or durable run-by-run report hierarchy.
- No fixed repetition count or numeric threshold before current evidence shows what comparability support is needed.
- No parallel production editing, per-bottleneck implementation slice, per-attempt commit, or documentation acceptance loop.
- No promotion based only on timing without the exact attempt's Supervisor accuracy/equivalence `PASS`.
- No Child 06 correctness re-review and no direct Child 06 -> Child 07 opening.
- No target-repository access or terminal Child 07 validation in this child.
- No performance claim from profiled cumulative samples, historical wall times, build/test duration, or an unmeasured resource hypothesis.
- No elapsed-time size cutoff or discretionary queue filter that can omit a measured parent or child. `BLOCKED` is valid only for concrete unavailable authority, dependency, or evidence and remains an unchecked plan blocker.

## Requirements

### Initial detailed measurement

- Start with exactly one visible measurement executor running the accepted real `anvien analyze` command and workload, recording only identity needed to prove the same work ran: executable/runtime, exact command/options, repository/input, relevant cache regime, start/end/exit, output identity, and workload denominator. Do not invent the exact options/workload/cache tuple in advance.
- Record total graph-generation elapsed time and real internal operation boundaries present in that run. Categories such as scan, parse, resolution, graph construction/insertion, persistence, diagnostics, and publication are examples only; use only boundaries proven to exist.
- Record every operation's absolute elapsed time and denominator in the living benchmark table. Record command validity and operation-boundary proof in evidence.
- Add CPU, RAM, allocation, GC, I/O, or wait evidence only when it explains the selected elapsed-time cost or proves a required resource invariant.
- Use existing timing, benchmark output, and profiles first. Only if one required P1-A timing is missing may Main assign exactly one visible sequential instrumentation-writer lane outside the read-only measurement analysts; that lane is the sole source writer for the branch.
- Before the conditional writer edits the exact timing owner, require the fresh repository graph plus file-detail and impact. Require the canonical full build to PASS before any instrumented timing is used or ranked; record exact writer/owner/owned bytes, runtime and instrumentation identity, overhead/comparability, denominator, and output equivalence, and compare only like-for-like instrumentation states.
- Close that conditional branch with exactly one disposition: carry the instrumentation with exact ownership into the first refreshed P2-A attempt, or remove the exact owned bytes, run the canonical full build again, and re-establish and remeasure the accepted timing basis. The branch is not a phase, implementation slice, optimization attempt, harness project, Architect/Supervisor gate, progress claim, or separate commit/report system.
- After the initial accepted capture exists, at most three ACTIVE read-only measurement analysts from the ten-lane waiting pool may share it for independent measured problems. They do not edit source or run duplicate competing benchmark processes.
- End P1-A with an ordered list derived entirely from current absolute measurements. Record owner/call path only after evidence identifies it; unknown ownership cannot be guessed.

### Per-attempt architecture, planning, implementation, and accuracy gate

- Select the largest current unprocessed top-level parent from the complete benchmark list and establish its current accepted timing/denominator. Later smaller parents remain queued.
- As the first processing step for that parent, measure more deeply inside only its proven runtime boundary, record elapsed time and denominator for every real measured child, and create the complete child list in descending absolute elapsed-time order. Do not drill down every parent upfront or create a separate phase.
- Mirror the exact parent and complete child row IDs/names into the plan checklist. Select the largest unprocessed child, prove its exact cause, source owner, complete call path, and work denominator, and keep every smaller child queued under the same parent.
- Open a new Visible Architect turn for exactly the selected child attempt only after the complete child list exists. The packet contains the parent row, complete current child list, selected child row, current accepted child/parent/total basis, denominator, exact cause/owner/call path, preserved invariants, and any prior rejection evidence.
- Record the Architect's exact direction, allowed owners, expected observable gain, validation/resource boundaries, and rollback condition. The decision expires when the attempt reaches disposition.
- Planner then refreshes the Current P2-A Attempt card and replaces generic steps with concrete files/symbols, production change, tests, build, measurements, equivalence checks, and rollback for that attempt. No Coder begins before this refresh is recorded.
- Coder implements only that refreshed direction after fresh required Anvien graph/file-detail/impact on the exact owner. Production implementation comes first; only after production behavior is correct may Coder edit the authorized test files.
- Run the exact build-holder/process preflight and canonical full build. Only after full-build `PASS` run the focused tests and `internal/resolution` package tests as validation evidence.
- Later and separately, remeasure the selected child with the same denominator, remeasure its parent, and remeasure total graph-generation time on the same workload.
- Open a new Visible Supervisor check for that attempt's changed candidate and affected accuracy/correctness/equivalence/output/lifecycle boundary.
- Record one disposition only after the Supervisor result: `KEEP`, `REWORK`, `ROLLBACK`, concrete `BLOCKED`, or terminal `SYSTEM_CHARACTERISTIC` after the required three-attempt history. A row's small time is never a disposition.
- `KEEP` requires lower selected-child elapsed wall-clock time, lower retained parent elapsed wall-clock time, lower retained end-to-end graph-generation elapsed wall-clock time, and Supervisor `PASS`; only then does the candidate become the current accepted baseline.
- Supervisor `REJECT` forces scoped restoration of the last accepted baseline, a new attempt-local internal drill-down on that accepted state, and a correction packet with the current drill-down back to a new Architect. Rejected bytes cannot become baseline, cannot be ranked, and cannot authorize a different bottleneck. If the unsuccessful streak is below `3`, a new attempt on the same bottleneck is mandatory unless a concrete blocker exists.
- Count an attempt as unsuccessful when it produces no retainable `KEEP`: Supervisor `REJECT`, no retainable direct/end-to-end improvement, or a `REWORK`/`ROLLBACK` outcome. Store the streak against the exact bottleneck row and current accepted baseline.
- On `KEEP`, promote the candidate, reset the selected child's consecutive unsuccessful-attempt count to `0`, remeasure that child/parent/full pipeline, and continue the same child until its remaining retained cost reaches terminal `SYSTEM_CHARACTERISTIC` or a concrete blocker is resolved.
- On the third consecutive unsuccessful attempt, record the exact three attempt/evidence references and reasons, preserve the last accepted correct timing, and terminalize only the selected child as `SYSTEM_CHARACTERISTIC`.
- After a child becomes terminal, mark its plan checkbox and select the largest remaining unprocessed child under the active parent. A concrete `BLOCKED` child stays unchecked and blocks parent completion; it cannot be used as a small-row shortcut.
- Mark a parent complete only after every measured child is checked with retained terminal evidence. Then remeasure the parent/full pipeline, refresh the complete top-level list, and process the largest remaining unchecked parent. Continue until every measured parent and child is checked.

### Final equivalence and closure

- Preserve initial total and operation baselines for final comparison even though every retained attempt uses the immediately preceding accepted state as its incremental baseline.
- Final acceptance requires lower measured total graph-generation time on the same workload, no unexplained transfer of cost to another operation, and exact required equivalence across all affected accepted boundaries.
- One stable final state receives the final whole-candidate P3-A Supervisor review; this does not replace or retroactively supply any per-attempt review.
- Cleanup removes only Child 06A-created failed, superseded, or debug work and must not mutate accepted production/test bytes.
- One post-cleanup detect and one isolated implementation commit close Child 06A; Child 07 consumes that exact closure commit.

## Acceptance Criteria

- P6-D commit `81163e39718b94a509e41114cada224e8f269e36` remains the immutable direct predecessor.
- P1-A records a valid current total graph-generation time, real internal operation timings, denominators, and an ordered absolute bottleneck ranking without production optimization.
- If P1-A uses timing instrumentation, `E1-P1A-INSTR1` proves the sole visible writer, exact owner/impact, full-build PASS before timing use, runtime/instrumentation identity, like-for-like overhead/comparability, output equivalence, and exactly one completed carry-forward or remove/rebuild/remeasure disposition; no instrumentation branch remains half-open.
- The benchmark living table contains current rank, real operation/bottleneck, current elapsed time, denominator, initial baseline, current before/after, meaningful share/delta, actionability, current attempt, direct delta, end-to-end before/after, cumulative initial delta, consecutive unsuccessful-attempt count, disposition, proven owner/call path when available, and exact evidence.
- P1-A cannot close and P2-A cannot open until benchmark contains the complete ordered top-level bottleneck list for every currently measured real operation and plan contains one unchecked parent checklist row for each benchmark parent. The largest unchecked parent is first; every smaller parent remains mandatory.
- P2-A is the only implementation slice and exhaustively processes every measured parent and child in largest-first order at each level.
- Before the first child attempt of a parent, P2-A creates the complete measured child list, mirrors every child into the nested plan checklist, selects the largest unchecked child, and proves its cause/owner/complete call path; no coarse parent name or incomplete child inventory can authorize architecture or code.
- Every production attempt has a distinct fresh Visible Architect decision, a recorded Planner refresh of the current P2-A attempt before Coder, and a distinct post-remeasurement Visible Supervisor accuracy/equivalence result.
- Every retained production change has exact parent/child cause/owner/impact evidence, production-first implementation, authorized test edits only after production is correct, canonical full build, post-build focused/package test execution, child before/after, parent before/after, end-to-end before/after, affected equivalence, Supervisor `PASS`, and an immediate `KEEP` disposition.
- Only `KEEP` attempts are promoted; each promoted state becomes the next current baseline and resets the row's unsuccessful streak. Rejected candidates are restored without ranking or transition on rejected bytes.
- The same child continues through new attempts after each `KEEP`. Three consecutive attempts without `KEEP` at one accepted baseline terminalize only that child as `SYSTEM_CHARACTERISTIC`; the next-largest unchecked child of the same parent follows.
- Every measured child of every measured parent has a retained terminal disposition and a checked plan item; every parent is checked only after all its children are checked. No small measured row is omitted.
- Final total graph-generation time is lower than the initial comparable value on the same workload.
- Accuracy, semantic completeness, graph correctness, ordered output, determinism, freshness, failure, transaction, temp, publication, Graph JSON, Ladybug, and native-reader invariants pass at every affected boundary.
- Exactly one final whole-candidate `P3-A` Supervisor boundary passes, `P3-B` cleanup leaves accepted production/test bytes unchanged, and `P3-C` records one detect, one implementation commit, and the direct Child 07 handoff.
- No removed-method control structure, template placeholder, parallel report tree, reusable stale architecture decision, or fabricated measurement/result remains in the four ledgers.
- The plan directory contains exactly four standard planner files plus the auxiliary [plan-rules.md](plan-rules.md); standard companion/suffix checks still count exactly the four standard files.

## Legacy Provenance Mapping

| Legacy authority | Current Child 06A authority | Traceability rule |
|------------------|-----------------------------|-------------------|
| Child 06 `P6-E` | P1-A measurement + sole P2-A implementation + P3-A/P3-B/P3-C closure | intent transferred without retaining a fixed solution order |
| legacy `E6-P6E-*` | `E0-P0A-*`, `E1-P1A-*`, dynamic `E2-P2A-*`, and `E3-P3A/P3B/P3C-*` | exact family mapping is maintained in evidence |
| dependent legacy `E6-PNA/PNB/PNC-*` | `E3-P3A/P3B/P3C-*` | aggregate performance review/cleanup/commit/handoff moved out of closed Child 06 |
| legacy `B6-P6E-*` | historical `B0-P0A-*`, initial `B1-P1A-*`, and dynamic `B2-P2A-*` | exact mapping is maintained in benchmark; historical values never select work |

## Measured Bottleneck Checklist

This is the living execution checklist requested for the bottlenecks discovered by measurement. It mirrors exact benchmark row IDs and real operation names so an executor can see and mark what has been processed; it never copies elapsed-time values or replaces benchmark ordering.

Current checklist entries: `30` top-level parent items plus `17` nested child items under active parent `B1-P1A-OP001`. The child items exactly match `B2-P2A-A001-D001..D017`, remain unchecked, and preserve their complete recorded identities. Main recorded A001 `KEEP` after `SUPERVISOR_A001_PASS`; `B2-P2A-A001-D001` (`resolve_calls`) remains the active unchecked child with streak `0`, and the next campaign cursor is A002 at fresh Architect pending. The parent remains unchecked and D002-D017 remain queued/unopened.

- [ ] `B1-P1A-OP001` — `resolution`
  - [ ] `B2-P2A-A001-D001` — `resolve_calls`
  - [ ] `B2-P2A-A001-D002` — `resolve_accesses`
  - [ ] `B2-P2A-A001-D003` — `resolve_type_annotations`
  - [ ] `B2-P2A-A001-D004` — `project_resolution_outcomes`
  - [ ] `B2-P2A-A001-D005` — `emit_definition_nodes`
  - [ ] `B2-P2A-A001-D006` — `finalize_resolution_outcomes`
  - [ ] `B2-P2A-A001-D007` — `build_binding_occurrence_index`
  - [ ] `B2-P2A-A001-D008` — `finalize_typescript_authority_results`
  - [ ] `B2-P2A-A001-D009` — `emit_import_edges`
  - [ ] `B2-P2A-A001-D010` — `emit_typescript_external_symbols`
  - [ ] `B2-P2A-A001-D011` — `emit_method_dispatch_edges`
  - [ ] `B2-P2A-A001-D012` — `assemble_resolution_result`
  - [ ] `B2-P2A-A001-D013` — `binding_accumulator_dispose`
  - [ ] `B2-P2A-A001-D014` — `emit_heritage_compatibility_edges`
  - [ ] `B2-P2A-A001-D015` — `emit_unresolved_heritage_diagnostics`
  - [ ] `B2-P2A-A001-D016` — `finalize_resolution_metadata`
  - [ ] `B2-P2A-A001-D017` — `runtime_setup`
- [ ] `B1-P1A-OP002` — `db_load`
- [ ] `B1-P1A-OP003` — `parse`
- [ ] `B1-P1A-OP004` — `graph_snapshot`
- [ ] `B1-P1A-OP005` — `semantic_enrichment`
- [ ] `B1-P1A-OP006` — `db_runner_resolve`
- [ ] `B1-P1A-OP007` — `cross_file_binding`
- [ ] `B1-P1A-OP008` — `ai_context`
- [ ] `B1-P1A-OP009` — `analyzer_orchestration`
- [ ] `B1-P1A-OP010` — `processes`
- [ ] `B1-P1A-OP011` — `scan`
- [ ] `B1-P1A-OP012` — `file_projection`
- [ ] `B1-P1A-OP013` — `routes`
- [ ] `B1-P1A-OP014` — `documents`
- [ ] `B1-P1A-OP015` — `db_runner_close`
- [ ] `B1-P1A-OP016` — `analyze_setup`
- [ ] `B1-P1A-OP017` — `tools`
- [ ] `B1-P1A-OP018` — `communities`
- [ ] `B1-P1A-OP019` — `registry_meta`
- [ ] `B1-P1A-OP020` — `cli_preparation`
- [ ] `B1-P1A-OP021` — `output_publication`
- [ ] `B1-P1A-OP022` — `orm`
- [ ] `B1-P1A-OP023` — `cli_startup`
- [ ] `B1-P1A-OP024` — `structure`
- [ ] `B1-P1A-OP025` — `graph_compact`
- [ ] `B1-P1A-OP026` — `mro`
- [ ] `B1-P1A-OP027` — `benchmark_write`
- [ ] `B1-P1A-OP028` — `cobol`
- [ ] `B1-P1A-OP029` — `memory_profile`
- [ ] `B1-P1A-OP030` — `cpu_profile_completion`

Population and marking rules:

- Before P1-A closes, insert one unchecked top-level parent checklist item for every measured `B1-P1A-OPnnn` row, in the same descending elapsed-time order as benchmark. Each item names its exact benchmark row ID and real operation.
- When P2-A selects a parent and completes drill-down, insert one nested unchecked child checklist item for every measured child row under that parent, in the same descending elapsed-time order as the complete benchmark child list. Do this before opening the first child Architect attempt.
- Cardinality must match exactly: if one parent drill-down measures `10` child bottlenecks, benchmark must contain exactly `10` child rows for that parent and this plan must contain exactly `10` nested child checkbox items before Architect. A missing or extra row/item blocks the attempt.
- Do not insert generic candidate, filename, solution, or empty placeholder items. Every checkbox must come from a real measured benchmark row.
- Keep the active child unchecked after `KEEP`; the retained baseline improves, its no-KEEP counter resets, and optimization continues on that same child.
- Check a child only after retained evidence records terminal `SYSTEM_CHARACTERISTIC`. A concrete unavailable authority/dependency/evidence `BLOCKED` row remains unchecked and blocks parent completion.
- Check a parent only after every measured child beneath it is checked and parent/full-pipeline remeasurement is recorded. If later accepted measurement exposes a previously missing child, reopen that parent, append the real child row, and continue P2-A.
- If later accepted measurement exposes a previously missing top-level parent or child, append its real benchmark row and matching unchecked plan checkbox immediately. A new parent joins the remaining top-level queue; a new child reopens/remains under its parent. No discovered row may exist in only one ledger.
- P2-A and `E2-P2A-EXHAUST1` cannot close until every measured parent and every nested measured child in this checklist is checked.

## Current P2-A Attempt

This is the living Planner refresh surface for the one implementation slice. Planner replaces this row and the matching concrete attempt steps after each fresh Architect decision and before any Coder edit. Metrics remain in benchmark.

| Field | Current value |
|-------|---------------|
| Attempt state | `A001 KEEP / A002 ARCHITECT_PENDING`; `SUPERVISOR_A001_PASS` is recorded under `E2-P2A-A001REVIEW1`, Main's separate disposition/promotion is recorded under `E2-P2A-A001DECISION1`, and Coder remains closed for A002 |
| Attempt ID | `A002`; A001 is closed as `KEEP` history |
| Attempt goal | obtain a fresh A002 Architect decision for the same active unchecked `B2-P2A-A001-D001 resolve_calls`, then reduce that child, retained parent `B1-P1A-OP001 resolution`, and retained end-to-end elapsed time from each repo's own accepted A001 baseline while preserving graph/output/lifecycle invariants |
| Active parent benchmark row / checklist item | `B1-P1A-OP001` / unchecked `resolution` parent item |
| Complete parent child list | `17/17` rows `B2-P2A-A001-D001..D017` are recorded independently for both targets and exactly `17` nested checklist items remain unchecked; D002-D017 remain queued/unopened and target-specific numbers stay separate in benchmark |
| Active child benchmark row / checklist item | `B2-P2A-A001-D001` / unchecked `resolve_calls` |
| Remaining child queue | `B2-P2A-A001-D002..D017`, all unchecked and ordered by descending same-run elapsed time |
| Current benchmark authority | A001 optimized values are the accepted current baselines separately for `E:\cheapapp.org` and `E:\Restaurant_manager` under `E2-P2A-A001DECISION1`; existing `E:\Anvien` rows stay superseded and the two repos are never averaged or combined |
| Consecutive unsuccessful attempts at current child baseline | `0`; A001 `KEEP` reset/retained the D001 streak at `0` on the promoted repo-specific baselines |
| Selected child exact cause / owner / complete call path | A001 historical owner/cause remains `E:\Anvien\internal\resolution\resolve.go:resolveCall:92-94`, where repeated import-claim scans/path normalization were removed by the accepted candidate; path remains `internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> for each ir in w.files -> for each call in ir.Calls -> resolution.resolveCall`; A002 cause/direction must be freshly decided and is not inferred |
| Fresh Architect decision | A002 `ARCH1` pending; A001 `E2-P2A-A001ARCH1` expired at A001 disposition and remains history only. Main Orchestration must open a fresh visible Architect for A002 before any new Coder |
| Planner refresh evidence | A002 `PLAN1` pending until its fresh Architect decision; A001 `E2-P2A-A001PLAN1` remains completed history |
| Owner-required consistency / transition evidence | `E2-P2A-A001CONSISTENCY1`; Architect thread `01a036f5-c602-7422-b484-ac8e58a34eb9`, turn `01a0377a-34a0-7523-b10c-820f84553fdb`, returned `ARCHITECT_CONSISTENCY_REVIEW_PASS_READY_FOR_MAIN`; Main accepted that planning/architecture evidence and authorized the Coder transition |
| Fresh pre-edit impact | A002 pending; A001 `E2-P2A-A001IMPACT1` remains accepted history and does not authorize a new edit |
| Allowed production/test surfaces | unbound for A002 until its fresh Architect decision and Planner refresh; A001's exact accepted five-file candidate remains preserve-only current baseline |
| Expected observable gain | unbound for A002; no cause, direction, owner, elapsed threshold, or saving is fabricated before its fresh Architect decision |
| Preserved invariants / validation/resource boundary | exact A001 resolution, order, graph/output, persistence/reader, determinism/freshness/failure/lifecycle, and private run-scoped boundaries are listed below; memory must stay `O(imports + unique keys)` and store indices, not duplicate import objects |
| Rollback condition / exact owned bytes | A001 rollback is closed by Main `KEEP`; its exact accepted hunks remain current baseline. A002 rollback is unbound until fresh architecture and must never disturb A001 accepted bytes or carried P1 instrumentation |
| Coder source/build/test evidence | A001 `E2-P2A-A001SRC1/BUILD1/TEST1` remains accepted input to `REVIEW1/DECISION1`: exact five-file candidate `+174/-10`, canonical full build PASS before tests, focused PASS, and full package truthfully FAIL only on the proven pre-existing preserve-only golden; the accepted bytes are committed under `E2-P2A-A001COMMIT1` |
| Coder status | A001 complete/closed and `A001_COMMIT_COMPLETE`; A002 Coder not opened or authorized |
| A001 implementation commit | `17a1f3af37dcb61f9d389345822b6470a8f772cc` — `perf(resolution): index import claims for call resolution`; verified manifest exactly `12` paths and staged set empty; this records A001 only and does not check P2-A/parent/D001 or satisfy final P3-C closure |
| Coder completion/commit report pointer | `E:\Anvien\reports\coder\rp_coder_260825_175710_by_gpt-5_child06a_a001_import_claim_index.md`; created after implementation commit `17a1f3af37dcb61f9d389345822b6470a8f772cc` and not part of its 12-path manifest; committed separately at `18b5063d236f9f2567fea90e48eca8f1501bd1eb` with subject `docs(report): record Child 06A A001 coder handoff` and a one-file manifest exactly `reports/coder/rp_coder_260825_175710_by_gpt-5_child06a_a001_import_claim_index.md` |
| Post-measurement Supervisor | `E2-P2A-A001REVIEW1`; report `E:\Anvien\reports\Supervisor\rp_supervisor_260825_172308_by_gpt-5_a001_measurement_equivalence_acceptance.md`; exact verdict `SUPERVISOR_A001_PASS`; acceptance evidence only, while Main separately owns and issued `KEEP` |
| Target 1 / lane / report | `E:\cheapapp.org` / `01a033a5-4700-7272-a28d-de3a71f58135` / `E:\Anvien\reports\Investigation\rp_child06a_a001_cheapapp_benchmark.md`; optimized D001/parent/process values accepted as this repo's current baseline under `E2-P2A-A001DECISION1` |
| Target 2 / lane / report | `E:\Restaurant_manager` / `01a033a5-4ec3-7e42-8946-0ab9172f6088` / `E:\Anvien\reports\Investigation\rp_child06a_a001_restaurant_manager_benchmark.md`; optimized D001/parent/process values accepted separately as this repo's current baseline under `E2-P2A-A001DECISION1`; identical exclusion `electron/renderer/src/api/userApi.ts` remains recorded |
| Two-target measurement contract | each lane independently measures pre-A001 baseline versus current optimized A001 for D001 elapsed/denominator, parent `resolution`, end-to-end analyze, graph/output equivalence, and exact executable/command/target identity; the two reports remain separate and are never averaged or combined |
| Next action | next owner Main Orchestration opens a fresh visible Architect for A002 on the same active unchecked D001 using both separate accepted A001 baselines; no Coder opens until A002 `ARCH1` and Planner refresh exist |

### A001 Measurement And Work-Count Contract

- `child_total_elapsed` is the sole value for child ranking, active-child selection, D001 before/after comparison, and child-improvement acceptance. CPU samples, internal sub-metrics, build/full-build duration, microbenchmarks, and counters cannot replace it.
- Deep attribution is active-child-only for `B2-P2A-A001-D001`. Coarse `child_total_elapsed` remains enabled for the full `B2-P2A-A001-D001..D017` list; D002-D017 stay queued/unchecked and receive no deep instrumentation or implementation scope now.
- The future D001 breakdown is mutually exclusive and must conserve the full child wall time without invented values:

```text
prepare_lookup_elapsed
+ resolution_decision_elapsed
+ graph_mutation_elapsed
+ diagnostic_outcome_elapsed
+ metrics_bookkeeping_elapsed
+ residual_elapsed
= child_total_elapsed

overlap_count = 0
```

- `prepare_lookup_elapsed`: path/input normalization, index/import scanning, and candidate preparation or lookup before a semantic decision.
- `resolution_decision_elapsed`: binding, member, package, authority, target, proof, and resolution-outcome decision work not already owned by lookup.
- `graph_mutation_elapsed`: node, relationship, reference, property, and metadata creation or merge work.
- `diagnostic_outcome_elapsed`: resolution outcome and unresolved/diagnostic recording.
- `metrics_bookkeeping_elapsed`: metrics/counter bookkeeping not already owned by another category.
- `residual_elapsed`: child elapsed not attributed to the five explicit categories.
- D001 before/after work comparability requires pending counters for files, calls, import-claim helper invocations, global-import candidates visited, import-path normalizations at claim lookup sites, index entries, unique lookup keys, matching-bucket candidates visited, resolution targets considered, references/relationships emitted, outcomes/diagnostics recorded, and affected graph node/relationship/property inventories.
- Counters prove comparable work and never replace elapsed wall time. Do not emit a per-call log, trace, or timer artifact for `72976` calls; use the smallest cumulative local accumulators and merge each once. Preserve the same instrumentation identity before and after. Invalid or mixed instrumentation cannot rank or prove gain. If a full-scan fallback counter is recorded, its accepted candidate value must be zero.

### A001 Production Direction And Exact Owner Boundary

- Create a separate run-scoped all-import claim index: `{canonical source file path, exact local import name} -> []original w.imports indices in existing order`.
- Build it during import construction after source-path canonicalization. Include resolved/unresolved, semantic/nonsemantic, and duplicate imports. Append original indices in existing `w.imports` order; never derive resolution order by iterating map keys.
- Replace only global candidate selection in `(*workspace).explicitImportNameClaimed` and the candidate-selection loop of `(*workspace).explicitImportCallState`. Preserve their semantic filters, target checks, `allowed`/`importTargets`, target-nil behavior, lexical shadow behavior, and import-export resolution.
- Do not reuse or change resolved-only `importsByReceiver` semantics. Do not cache final resolution answers. Do not change `resolveCall` branch ordering or any preserved helper/contract.
- Exact future production owners:
  - `internal/resolution/indexes.go`: one private `workspace` field; its `buildWorkspace` initialization; original-index population in `(*workspace).resolveImports`.
  - `internal/resolution/resolve.go`: body of `(*workspace).explicitImportNameClaimed` only.
  - `internal/resolution/export_resolution.go`: candidate-selection loop of `(*workspace).explicitImportCallState` only.
- Exact future test owners, after production is correct: `internal/resolution/export_resolution_test.go` and `internal/resolution/resolution_test.go`.
- Preserve/inspect only: `resolveCall`, `repositoryReceiverClaimed`, `cleanPath`, `scopeFilePath`, import target/file resolution, graph emitters, persistence/readers, every exported/shared contract, and every other production/helper/test surface. If Coder needs any other production owner or helper, STOP and return for a fresh architecture decision; do not broaden A001.

### A001 Preserved Invariants

- Preserve exact resolution outcomes, confidence, proof kind, unresolved notes, and call-branch order.
- Preserve exact query trimming and case-sensitive import-local-name semantics.
- Preserve original import ordering, duplicates, first-match behavior, and traversal order.
- Preserve current `cleanPath` canonical-path behavior without lowercase, absolutization, or symlink resolution.
- Preserve relative, absolute, language-specific, and Go/package import handling.
- Preserve unresolved import claims and global-rescue blocking behavior.
- Preserve nonsemantic import-name claims and semantic call-state filtering.
- Preserve graph nodes, relationships, IDs, labels, properties, counts, ordered output, resolution outcomes, and resolution indexes.
- Preserve canonical in-memory graph, Graph JSON, Ladybug/native persistence and readback, and affected native/product readers.
- Preserve deterministic replay/output, source freshness/invalidation, failure propagation, transaction rollback, temporary-file flush/close/rename, and publication visibility.
- Keep the index private, run-scoped, and in-memory with no new serialization, I/O, goroutine, global cache, public schema, or shared contract.

### A001 Future Validation, Resource, And Rollback Sequence

1. COMPLETE under `E2-P2A-A001IMPACT1`: fresh `anvien analyze --force`, file-detail, and exact impact ran immediately before edit; CRITICAL/HIGH scope remains recorded without treating helper `LOW/0` as assurance.
2. COMPLETE under `E2-P2A-A001SRC1`: only the exact production direction was implemented first.
3. COMPLETE under `E2-P2A-A001SRC1/TEST1`: only the two authorized test files were edited after production inspection.
4. COMPLETE under `E2-P2A-A001BUILD1`: build-holder/process preflight found no holder to terminate; exact canonical full build returned `PASS` before test execution.
5. RECORDED under `E2-P2A-A001TEST1`: focused tests PASS; the full package is truthfully `FAIL` only at a golden proven identical with all five A001 files overlaid from HEAD, so it is pre-existing preserve-only and neither package PASS nor A001 failure/proof.
6. COMPLETE FOR MEASUREMENT RECORDING: Cheapapp and Restaurant Manager reports are each recorded independently with exact D001/parent/end-to-end/identity/equivalence plus all 30/17 rows; targets remain separate and are never averaged or combined.
7. COMPLETE FOR A001 REVIEW/DISPOSITION: `E2-P2A-A001REVIEW1` records exact verdict `SUPERVISOR_A001_PASS`; Main then issued A001 `KEEP` under `E2-P2A-A001DECISION1` because D001, parent, and process E2E all improved on both targets with same-work canonical graph/output equivalence. Each repo's optimized values are now that repo's accepted current baseline; D001 remains unchecked/active, streak is `0`, and A002 is fresh Architect pending.
8. Keep index memory `O(imports + unique keys)` and store original indices rather than duplicate import objects. Roll back exact A001-owned hunks if any build/test/outcome/order/graph/persistence/reader/determinism/failure/lifecycle invariant changes, memory becomes unstable or exceeds the linear model, valid same-work measurements do not improve D001 with corresponding parent/end-to-end net benefit, or the implementation crosses the exact owner boundary.
9. Scoped rollback reverses only the private field/init/population, two claim-helper traversal hunks, and authorized focused-test bytes. It must not reset/checkout/stash broadly or disturb user/protected work or the carried P1 instrumentation in `internal/analyze/analyze.go`, `internal/cli/command.go`, and `internal/cli/command_test.go`.

### A001 Owner-Direct Review Gate — Completed

This gate applies to A001 because it is the latest direct Owner sequencing requirement; it is not an Architect command to Main and is not generalized to every future attempt.

1. Main independently verified the exact four refreshed ledgers.
2. Architect thread `01a036f5-c602-7422-b484-ac8e58a34eb9`, turn `01a0377a-34a0-7523-b10c-820f84553fdb`, independently reviewed the refreshed artifacts and returned `ARCHITECT_CONSISTENCY_REVIEW_PASS_READY_FOR_MAIN` with zero file changes and read-only commands only.
3. Main accepted the result as planning/architecture-gate evidence only and authorized the A001 Coder transition. `E2-P2A-A001CONSISTENCY1` records this completed sequence; it is not Supervisor `PASS`, implementation evidence, measurement, disposition, or commit authorization.
4. No Architect re-review was required for that truthful state/order sync. Main retained governance authority; the sole Coder authorization was limited to A001 pre-edit/production/test/build validation. A001 later completed measurement, `SUPERVISOR_A001_PASS`, Main `KEEP`, and exact implementation commit `17a1f3af37dcb61f9d389345822b6470a8f772cc`; A002 still starts at fresh Architect rather than reopening Coder.

## Checklist

- [x] P0-A: Establish current truth and direct campaign order.
  - Goal: record transferred state without inventing a baseline, bottleneck, result, or implementation completion.
  - Work Steps:
    1. Preserve the accepted P6-D closure anchor and direct Child 06 -> Child 06A -> Child 07 order.
    2. Record that no current detailed timing map, ranked bottleneck, Child 06A implementation, final result, Supervisor verdict, or implementation commit exists.
    3. Bind the empirical method, per-attempt architecture/planning/accuracy chain, and lossless legacy mapping into the four ledgers.
  - Implementation Gate: documentation-only P0 creates no functional evidence and cannot claim a measured baseline.
  - Acceptance: `E0-P0A-ANCHOR1`, `E0-P0A-ORDER1`, `E0-P0A-TRUTH1`, `E0-P0A-METHOD1`, and `E0-P0A-PROVENANCE1` record current state and open P1-A directly.

### P1: Measure The Real Pipeline And List Every Top-Level Bottleneck

- Phase Goal: create the first trustworthy numeric map of current graph-generation time, the complete largest-first top-level queue, and its exact plan checklist mirror.
- Phase Boundary:
  - In scope: one measurement-only slice, P1-A.
  - Out of scope: production optimization, a fixed solution list, Architect/Coder/Supervisor attempt work, or a speedup claim.
  - Dependencies: P0-A and accepted P6-D runtime/workload authority.
- Phase Implementation Rule: P1-A writes measurements directly into benchmark and proof into evidence; a conditional sole-writer timing edit follows the linked rules, creates no optimization attempt or implementation commit, and cannot supply rankable timing before its canonical full-build PASS.
- Ordered Slice List:
  - P1-A: Measure Detailed Graph-Generation Time And List Every Top-Level Bottleneck.

- [x] P1-A: Measure Detailed Graph-Generation Time And List Every Top-Level Bottleneck.
  - Goal: measure the current real pipeline deeply enough to produce a complete evidence-backed top-level ranking, preserve every smaller measured row as mandatory work, and materialize the matching plan checklist.
  - Scope Boundary:
    - Editable: four Child 06A ledgers; only the exact timing owner may additionally be edited by the one visible sequential instrumentation writer when existing timing/benchmark/profile capability lacks one required P1-A timing.
    - Inspect-only: current real pipeline boundaries, runtime command, workload, operation owners, and output/persistence/read surfaces.
    - Preserve-only: production behavior, accepted P6-D outputs and lifecycle, historical/protected artifacts.
    - Out of scope: production optimization, broad instrumentation, permanent report systems, target access, commit, or acceptance review.
  - Non-Goals: no bottleneck or owner is selected before current absolute timings exist; pure measurement requires no Architect.
  - Pre-flight Questions:
    - Data source: accepted real `anvien analyze` runtime and same workload inherited from P6-D, with exact options/cache/runtime identity recorded at execution.
    - Display permission: N/A — command-line analyzer work has no UI display scope.
    - DB read flow: inspect only the real persistence/read path needed to identify operation boundaries and output identity.
    - DB write flow: run the normal accepted analyze path; do not alter transaction or publication behavior.
    - Render location: numeric results in benchmark, validity proof in evidence, current pointer in actual status.
    - UI behavior flow: N/A — no browser-visible behavior is in scope.
    - Docker runtime: N/A — accepted repository-local CLI/runtime governs this non-UI measurement.
    - Playwright target: N/A — no browser target exists.
    - Behavior test: confirm accepted workload/output denominator and normal completion before ranking a run.
    - Cleanup/quarantine: temporary debug/profile material is confined to `E:\Anvien\.tmp` and is not durable evidence.
    - External side effects: normal repo-local analyze outputs only; no network, install, target, shared/global cache, or alternate worktree.
    - N/A notes: one visible measurement executor alone produces the initial accepted capture. The `10` pre-opened visible measurement lanes are waiting capacity, not a fixed work list; after capture acceptance, activate at most three read-only analysts for already-evidenced independent parent/child problems, one problem per lane, and release a slot only after benchmark/evidence/checklist/status recording is complete. The conditional instrumentation writer is a separate sequential sole-writer role, not a read-only analyst.
  - Work Steps:
    1. Assign exactly one visible measurement executor to run the accepted real command on the accepted workload and record current total graph-generation elapsed time; do not dispatch a competing capture.
       - UI flow check: N/A — CLI only.
       - DB/data flow check: record exact workload/output denominator, relevant cache regime, exit, and normal persistence completion.
       - Render location check: write the number to `B1-P1A-TOTAL`; write validity to `E1-P1A-IDENTITY1` and `E1-P1A-TOTAL1` immediately.
       - Mini QA: reject a run whose command, workload, output, exit, or relevant cache identity is not comparable.
       - Evidence target: `E1-P1A-IDENTITY1`, `E1-P1A-TOTAL1`.
    2. Measure elapsed time for real internal operations present in that accepted capture, using existing timing/benchmark/profile output first.
       - UI flow check: N/A — CLI only.
       - DB/data flow check: bind every operation to its real boundary and denominator; preserve output and lifecycle behavior.
       - Render location check: append/update living `B1-P1A-OPnnn` rows and boundary proof `E1-P1A-OPS1`; no separate cost map.
       - Mini QA: do not sum overlapping intervals or infer an unmeasured value; record any residual honestly.
       - Evidence target: `E1-P1A-OPS1`.
    3. If and only if one required operation timing is still missing, Main assigns exactly one visible sequential instrumentation-writer lane outside the read-only analysts and records the conditional branch immediately.
       - UI flow check: N/A — CLI timing source only.
       - DB/data flow check: before edit, refresh the required graph and record file-detail plus impact for the exact timing owner; preserve output, persistence, reader, and lifecycle behavior.
       - Render location check: `E1-P1A-INSTR1` records writer, exact owner/owned bytes, graph/file-detail/impact, canonical full-build PASS before timing use, accepted runtime/instrumentation identity, overhead/comparability, denominator, and output equivalence; benchmark receives numbers only after this proof exists.
       - Mini QA: compare like-for-like instrumentation states and finish with exactly one disposition: carry exact instrumentation ownership into the first refreshed P2-A attempt, or remove exact owned bytes, full-build again, and re-establish and remeasure the accepted timing basis. A half-open, unbuilt, mixed-instrumentation branch blocks ranking/P1 closure.
       - Evidence target: conditional `E1-P1A-INSTR1`, plus refreshed `E1-P1A-IDENTITY1`, `E1-P1A-OPS1`, and `E1-P1A-EQUIV1` as applicable.
    4. Sort every currently measured operation row by absolute elapsed time, publish the complete ordered top-level bottleneck list in benchmark, and insert one unchecked parent item for every exact row into the Measured Bottleneck Checklist in the same order.
       - UI flow check: N/A — ledger state only.
       - DB/data flow check: ranking uses the same accepted workload and denominators; owner/call path is recorded only when proven.
       - Render location check: benchmark owns the ordered numeric table; plan mirrors row IDs/names as checkboxes; actual status points to the largest unchecked parent without copying numbers.
       - Mini QA: every measured row appears both in benchmark and as one unchecked plan parent item, including smaller rows; each benchmark row records current elapsed time, denominator, meaningful share/delta when available, processing/terminal state, proven owner/call path when known, no-KEEP/disposition state, and exact evidence; no historical value or size cutoff controls inclusion.
       - Evidence target: `E1-P1A-RANK1`, `E1-P1A-EQUIV1`.
  - Current Execution Cursor (`2026-08-25`):
    - Step 1 complete: one visible executor produced accepted capture `child06a-p1a-initial-20260824-225900`; `E1-P1A-IDENTITY1`, `E1-P1A-TOTAL1`, and `E1-P1A-EQUIV1` are recorded.
    - Step 2 complete: the sole instrumented capture emits `30` unique, non-negative operation rows across `21 analyzer_internal`, `1 analyzer_outer`, and `8 cli_outer` boundaries; all `15` phase-operation names/order/durations match exactly and analyzer-internal sum equals `totalDuration` with residual `0`.
    - Step 3 complete: sole writer `01a0348f-3bb7-7c70-8ec5-61176ce00591` retains exact ownership of `internal/analyze/analyze.go +117/-15`, `internal/cli/command.go +157/-38`, and `internal/cli/command_test.go +75/-0` (`+349/-53` total). The successful capture ran once at HEAD `8df19258fbcb18e14841bf0dae036400aa9b22a3`, exit `0`, and the branch is terminal `CARRY_TO_FIRST_P2A_REFRESH`. Its timing delta versus the immutable uninstrumented reference is explicitly non-comparable because HEAD/workload/output counts differ.
    - Step 4 complete: benchmark contains every measured real operation as `B1-P1A-OP001..OP030` in descending absolute elapsed order and this plan contains exactly `30` matching unchecked parent items. `B1-P1A-OP001` (`resolution`) is active; P2-A opens only at `PARENT_DRILLDOWN_PENDING` with no child or production attempt yet.
  - Instrumentation Writer Checkpoint:
    - Identity/boundary: `Metrics.Operations []OperationMetric` serializes flat `boundary`, `name`, `duration`, and `denominators`; production behavior, CLI stdout/JSON, graph persistence, readers, and lifecycle remain unchanged apart from the current benchmark contract gaining timing fields.
    - Analyzer operations: `analyze_setup`; every existing `runPhase` phase; `graph_compact`; `db_runner_resolve`; conditional `db_runner_close`; `graph_snapshot`; exclusive residual `analyzer_orchestration`; and `benchmark_write`.
    - CLI operations: `cli_startup`, `cli_preparation`, `memory_profile`, `registry_meta`, `ai_context`, `file_projection`, `output_publication`, and `cpu_profile_completion`. The final benchmark rewrite only publishes the completed CLI operation list and is instrumentation overhead, not a product-operation timing row.
    - Denominators: analyze/command/run/artifact counts; graph input/output node and relationship counts; runner/profile/repository/output counts; generated AI-context file/base-skill counts; and file-projection file/node/relationship counts as recorded in `E1-P1A-INSTR1`.
    - Validation boundary: `pwsh -NoLogo -NoProfile -NonInteractive -ExecutionPolicy Bypass -File E:\Anvien\scripts\full-build.ps1`, start `2026-08-25T03:25:06.0193408+07:00`, end `2026-08-25T03:36:46.1906114+07:00`, exit `0`. Its `700.1712706 s` is full-build validation elapsed only, including maintenance `analyze . --force`; it is not build elapsed and not a P1-A benchmark observation.
    - Runtime identity ready for capture: version `1.2.8`, `73,764,864` bytes, SHA-256 `62920CBF15921EF8A6D2FAC671776BAA3C312EFEA1BEED53B721C4AFF5E1B6C5`.
    - Terminal disposition/state: Main selected `CARRY_TO_FIRST_P2A_REFRESH`; the accepted separate capture tuple completed exactly once and the maintenance analyze embedded in full build remains excluded from benchmark evidence. Next state is P2-A `PARENT_DRILLDOWN_PENDING` on `B1-P1A-OP001`.
  - Implementation Gate:
    - No production optimization is allowed in P1-A and no Architect is required for pure measurement/inventory.
    - Read-only measurement analysts cannot edit source. A missing timing may open only the one visible sequential instrumentation-writer branch; its fresh graph/file-detail/impact and full-build PASS must exist before instrumented timing can enter benchmark.
    - Minimum timing instrumentation must preserve accepted output, be measured like-for-like, and end in exactly one recorded state: explicitly carried with ownership into the first refreshed P2-A attempt, or removed with exact owned-byte cleanup followed by full build and accepted-basis remeasurement.
    - P2-A remains closed until the complete ordered top-level list exists in benchmark and every row has an exact unchecked plan parent item; a prose summary, hand-picked target, or largest-only item is insufficient.
  - Acceptance:
    - Source: no production optimization or preselected owner was introduced.
    - Runtime/UI: one visible executor produced the valid accepted capture; one valid current total, detailed real operation timings, the complete benchmark-owned ordered top-level list, and the matching plan parent checklist exist; UI is N/A.
    - DB/data: workload/output denominator and normal persistence/publication boundary are identified.
    - Behavior test: measurement validity and affected initial equivalence pass without a speedup claim.
    - Cleanup/quarantine: no durable per-run report tree exists; temporary debug material is repo-local; any conditional instrumentation has one completed carry-forward or remove/rebuild/remeasure disposition.
    - Evidence IDs: `E1-P1A-IDENTITY1`, `E1-P1A-TOTAL1`, `E1-P1A-OPS1`, conditional `E1-P1A-INSTR1`, `E1-P1A-RANK1`, `E1-P1A-EQUIV1`.
    - Actual-status rows refreshed: accepted-capture executor/identity, read-only analyst assignments, conditional writer/owner/impact/build/comparability/disposition pointers when opened, current total pointer, active largest unchecked parent, remaining parent queue, active owner if known, last disposition, and exact next action.
  - Evidence Targets: command/identity validity, operation boundaries/denominators, initial output validity, and ranking derivation.
  - Actual-status Update: change `no current bottleneck` to the exact largest unchecked parent `B1-P1A-OPnnn`, point to the complete remaining parent checklist/queue, and make P2-A complete-child-list drill-down the next action. Only that complete child inventory can open a fresh Architect decision packet for attempt A001.
  - Commit Boundary: no implementation commit; measurements and ledgers remain inside the eventual single Child 06A commit.

### P2: Exhaust Every Measured Parent And Child Bottleneck

- Phase Goal: retain safe cumulative end-to-end improvement by processing every measured parent and every measured child largest first at each level, including all smaller rows.
- Phase Boundary:
  - In scope: exactly one implementation slice, P2-A, repeated internally for dynamically selected attempts.
  - Out of scope: predeclared solution families, overlapping production owners, reused architecture decisions, or per-attempt commits.
  - Dependencies: accepted P1-A complete top-level table, matching parent checklist, and exact largest unchecked parent row.
- Phase Implementation Rule: P2-A stays open across all attempts. Each production attempt receives a new Architect, Planner refresh, Coder execution, and post-measurement Supervisor check. No commit occurs before P3-C.
- Ordered Slice List:
  - P2-A: Optimize Every Measured Parent And Child In Largest-First Order.

- [ ] P2-A: Optimize Every Measured Parent And Child In Largest-First Order.
  - Goal: for every dynamically opened child attempt, reduce the selected child's measured elapsed time, its active parent's elapsed time, and total graph-generation elapsed time from the current accepted baseline, while preserving accuracy/equivalence and all required lifecycle invariants; exhaust every parent/child checklist row rather than stopping at the largest.
  - Scope Boundary:
    - Editable: only exact measured owner/tests allowed by the current attempt's fresh Architect decision and Planner refresh.
    - Inspect-only: complete active-parent child list, current call path, sibling operations, output/persistence/read consumers, and resource behavior needed for that attempt.
    - Preserve-only: unselected owners, workload, output format, accepted Child 06 behavior, historical/protected/shared work.
    - Out of scope: another implementation slice, speculative refactor, unmeasured concurrency, workload shortcut, or intermediate commit.
  - Non-Goals: completing a historical idea is never an attempt objective. Harness, instrumentation, profiling, report, audit, or architecture documentation cannot be an attempt goal; each may be only the smallest supporting step tied to reducing the selected child, parent, and total elapsed times.
  - Pre-flight Questions:
    - Data source: exact active parent `B1-P1A-OPnnn`, its complete child list, the largest unchecked child row, and same-work denominators.
    - Display permission: N/A — CLI analyzer work.
    - DB read flow: identify/validate only readers affected by the selected owner.
    - DB write flow: preserve transaction, snapshot/temp, persistence, rollback, and publication semantics.
    - Render location: numbers in benchmark, proof in evidence, active attempt in plan/status, normal product outputs unchanged.
    - UI behavior flow: N/A unless fresh impact proves a public UI consumer, which requires a plan/status refresh before edit.
    - Docker runtime: N/A for this repository-local CLI; canonical full build remains mandatory.
    - Playwright target: N/A unless impact activates a browser-visible consumer.
    - Behavior test: exact selected-child, parent, and end-to-end before/after plus affected correctness/output/lifecycle equivalence.
    - Cleanup/quarantine: identify attempt-owned instrumentation, failed code, tests, and debug artifacts; rollback only those bytes.
    - External side effects: normal repo-local build/analyze outputs only; no network, install, target, shared/global mutation, push, or alternate worktree.
    - N/A notes: CPU/RAM/allocation/GC/I/O/wait metrics are collected only when they explain selected elapsed-time cost.
  - Work Steps:
    1. Select the largest unchecked top-level parent, measure more deeply inside only its real boundary, create the complete child list in descending current absolute elapsed-time order, and mirror every exact child row/name as a nested unchecked plan item before any Architect opens.
       - UI flow check: N/A — CLI only.
       - DB/data flow check: prove the parent and every measured child belong to the accepted real path and bind each child to its work denominator.
       - Render location check: benchmark owns the complete child numeric list; plan mirrors every exact child as a nested checkbox; actual status points to active parent, active child, and remaining child queue.
       - Mini QA: reject an incomplete child list, largest-only list, size-based omission, historical hypothesis, secondary-metric ordering, or generic checklist item.
       - Evidence target: `E2-P2A-AnnnDRILL1`, exact parent `B1-P1A-OPnnn`, and complete dynamic child rows `B2-P2A-Annn-Dnnn`.
       - Current A001 cursor: complete for `B1-P1A-OP001`; `E2-P2A-A001DRILL1` binds all `17` rows/items, and all parent/child checkboxes remain unchecked.
    2. Select the largest unchecked child, bind its exact goal/current child-parent-total basis/denominator/cause/owner/complete call path, then open a new Visible Architect and obtain the exact attempt-scoped decision.
       - UI flow check: N/A — architecture decision for CLI production work.
       - DB/data flow check: decision names affected data/read/lifecycle invariants and resource boundary.
       - Render location check: record `E2-P2A-AnnnCURRENT1` and `ARCH1` linked to the exact parent row, complete child list, and selected child row; do not create a separate architecture report gate.
       - Mini QA: Architect receives the full current parent-child list and selected child's measured owner, then names exact cause, direction, allowed owners, expected observable gain, validation/resources, and rollback; a coarse parent name is insufficient and the decision cannot be reused.
       - Evidence target: `E2-P2A-AnnnCURRENT1`, `E2-P2A-AnnnARCH1`.
       - Current A001 cursor: complete through `E2-P2A-A001ARCH1`; the selected direction is the separate all-import claim index and the decision remains attempt-local.
    3. Planner refreshes the Current P2-A Attempt card and these execution steps to the Architect decision before Coder begins.
       - UI flow check: N/A — plan control update.
       - DB/data flow check: concrete validation covers every data/output/lifecycle invariant named by the decision.
       - Render location check: exact owner/files/symbols/change/tests/build/measure/equivalence/rollback replace the unbound attempt fields in this plan set.
       - Mini QA: `E2-P2A-A001PLAN1` is materialized; Main verification and the Owner-required Architect consistency review passed, and Main authorized the sole Coder transition under `E2-P2A-A001CONSISTENCY1`.
       - Evidence target: `E2-P2A-AnnnPLAN1`.
       - Current A001 cursor: consistency/Main authorization remains complete; Coder source/build/test work is now recorded under `E2-P2A-A001IMPACT1/SRC1/BUILD1/TEST1`.
    4. The sole Coder completes fresh required Anvien graph/file-detail/impact, implements only the refreshed production direction, edits authorized tests only after production is correct, performs the exact build-holder preflight, obtains canonical full-build `PASS`, and only then executes focused and `internal/resolution` package tests.
       - UI flow check: N/A unless activated by impact.
       - DB/data flow check: preserve exact affected data, ordering, failure, transaction/temp, publication, and reader semantics.
       - Render location check: cause/impact/source/build/test proof goes to evidence; product output remains observable boundary.
       - Mini QA: production is correct before tests change; test execution occurs only after full-build `PASS`; no owner beyond the exact A001 `indexes.go`, `resolve.go`, and `export_resolution.go` symbols is edited; needing another production helper forces STOP and fresh architecture.
       - Evidence target in execution order: `E2-P2A-AnnnIMPACT1`, `E2-P2A-AnnnSRC1`, `E2-P2A-AnnnBUILD1`, `E2-P2A-AnnnTEST1`; `TEST1` records both authorized post-production test edits and test execution only after full-build `PASS`.
       - Current A001 cursor: COMPLETE; exact five-file candidate `+174/-10`, build PASS, focused PASS, and pre-existing preserve-only package golden failure recorded. P2-A/parent/children remain unchecked.
    5. Record the two independent target measurements without combining them. Cheapapp lane `01a033a5-4700-7272-a28d-de3a71f58135` and Restaurant lane `01a033a5-4ec3-7e42-8946-0ab9172f6088` have each completed their assigned report; Planner records each repo-specific block separately and owns neither lane.
       - UI flow check: N/A unless activated by impact.
       - DB/data flow check: each target records its own D001 denominator, parent, end-to-end, graph/output equivalence, and exact executable/command/target identity; never average or combine targets.
       - Render location check: lane `01a033a5-4700-7272-a28d-de3a71f58135` writes only `rp_child06a_a001_cheapapp_benchmark.md`; lane `01a033a5-4ec3-7e42-8946-0ab9172f6088` writes only `rp_child06a_a001_restaurant_manager_benchmark.md` at the exact report paths above.
       - Mini QA: both target blocks remain complete and separate; old `E:\Anvien` rows are not the current control surface; A001 promotion follows only `REVIEW1/DECISION1`, never averaging, combination, or a fabricated number.
       - Evidence target: `E2-P2A-AnnnDIRECT1`, `E2-P2A-AnnnPARENT1`, `E2-P2A-AnnnE2E1`, `E2-P2A-AnnnEQUIV1`, `E2-P2A-AnnnREVIEW1`.
       - Current A001 cursor: COMPLETE; both measurements fed `E2-P2A-A001REVIEW1`, which returned `SUPERVISOR_A001_PASS`; Main issued `KEEP` under `E2-P2A-A001DECISION1` and promoted each repo's optimized values separately.
    6. Record disposition, promote only an eligible child+parent+end-to-end `KEEP`, and keep the same child active until terminal.
       - UI flow check: N/A — ledger-controlled transition.
       - DB/data flow check: rejected bytes never become accepted state or ranking input.
       - Render location check: update child/parent/E2E benchmark rows, promotion/history/streak, evidence disposition, plan checklist, actual-status queues, and this attempt card immediately.
       - Mini QA: `KEEP` resets the active child's streak but leaves its checkbox unchecked and starts another attempt on that child; reject restores accepted state and returns through fresh Architect/Planner, never Coder; only the third consecutive no-KEEP records `SYSTEM_CHARACTERISTIC`, refreshes accepted child/parent/E2E timings and both ordered lists, and checks that child.
       - Evidence target: `E2-P2A-AnnnDECISION1`, conditional `E2-P2A-AnnnRESTORE1`, terminal `E2-P2A-AnnnSYSTEM1`, and `E2-P2A-AnnnRANK1`.
       - Current cursor: A001 `KEEP` is recorded under `E2-P2A-A001REVIEW1/DECISION1` and `A001_COMMIT_COMPLETE` under `E2-P2A-A001COMMIT1`; `B1-P1A-OP001` and `B2-P2A-A001-D001` remain unchecked, D001 streak is `0`, D002-D017 remain queued, and A002 is `ARCHITECT_PENDING` on the same child.
    7. After a child becomes terminal, select the largest remaining unchecked child of the same parent. After every child is checked, remeasure parent/full pipeline, check the parent, refresh the complete top-level list, and select the largest remaining unchecked parent.
       - UI flow check: N/A — ledger-controlled hierarchy.
       - DB/data flow check: parent completion requires every measured child terminal on retained correct state; a new measured child reopens the parent.
       - Render location check: benchmark owns current hierarchy/order; plan marks exact rows; actual status points to remaining parent/child queues.
       - Mini QA: never move to another parent while a child is unchecked; never omit or check a small row by size; concrete `BLOCKED` stays unchecked; any later-discovered parent/child gets an immediate benchmark row plus matching unchecked plan checkbox before queue selection continues.
       - Evidence target: `E2-P2A-AnnnRANK1`, plan checklist refresh, parent-completion evidence in the matching result record.
    8. After every measured parent and child checklist item is checked, record final initial-versus-current total and complete final equivalence on the stable accepted implementation state.
       - UI flow check: N/A unless activated by impact.
       - DB/data flow check: verify canonical in-memory, Graph JSON, Ladybug/native, affected readers, determinism, freshness, failure, transaction/temp, and publication.
       - Render location check: final numbers in `B2-P2A-FINAL1`; proof in evidence; closure pointer in actual status.
       - Mini QA: final total is lower on same workload and every unexplained sibling/resource regression is resolved.
       - Evidence target: `E2-P2A-FINALBUILD1`, `E2-P2A-FINALTIME1`, `E2-P2A-FINALEQUIV1`, `E2-P2A-EXHAUST1`.
  - Implementation Gate:
    - Every production edit has a unique attempt ID, active parent, complete child list, selected largest unchecked child, current child/parent/total benchmark basis, exact cause/owner/call path, fresh attempt-specific Architect decision, and completed Planner refresh.
    - For A001 only, the latest direct Owner sequencing required Main verification of the four refreshed artifacts followed by independent consistency review in Architect thread `01a036f5-c602-7422-b484-ac8e58a34eb9`. That sequence passed and Main authorized the sole Coder transition under `E2-P2A-A001CONSISTENCY1`. This was Owner authority, not Architect command authority, and is not a generic gate for every attempt.
    - Every refreshed attempt states the child + parent + total elapsed-time reduction goal against the exact benchmark rows/current accepted baseline; supporting tooling cannot replace that goal.
    - Only one production/test writer is active. The measurement pool has `10` pre-opened visible lanes, at most `3` ACTIVE in one wave, exactly one measured parent/child problem per active lane, and no slot reuse until that lane's benchmark/evidence/checklist/status updates are recorded.
    - The previous attempt has a recorded Supervisor result, disposition, accepted-state promotion/restoration, unsuccessful-streak update, and current ledgers before another attempt opens.
  - Acceptance:
    - Source: only fresh-Architect-allowed measured owners changed; Coder followed the refreshed attempt exactly.
    - Runtime/UI: selected child, active parent, and total graph-generation elapsed wall time all improve for every `KEEP`; final total is lower; UI is N/A unless activated.
    - DB/data: each attempt's Supervisor and final evidence pass affected correctness, persistence, reader, failure, transaction/temp, and publication invariants.
    - Behavior test: production implementation, authorized test edits after production is correct, canonical full build, post-build focused/package test execution, later same-work remeasurement, per-attempt review, and final equivalence pass in that order.
    - Cleanup/quarantine: failed/superseded/debug work is identified for P3-B; protected work is untouched.
    - Checklist: every measured child and parent item is checked from retained terminal evidence; no `BLOCKED` or smaller measured row is hidden.
    - Evidence IDs: dynamic attempt families including `E2-P2A-AnnnDRILL1/PARENT1`, plus `E2-P2A-FINALBUILD1`, `E2-P2A-FINALTIME1`, `E2-P2A-FINALEQUIV1`, `E2-P2A-EXHAUST1`.
    - Actual-status rows refreshed: active parent/child, remaining parent/child queues, attempt state, unsuccessful streak, Architect/Planner/Coder/Supervisor pointers, last disposition, accepted baseline, and next action.
  - Evidence Targets: complete top-level list, complete child lists, exact plan checklist mirror, parent/child selection, cause/owner/call path, Architect, Planner refresh, impact/source/tests/build, child/parent/E2E measurements, Supervisor, disposition, restoration/promotion, three-attempt terminal proof, parent completion, reranking, exhaustion, final proof.
  - Actual-status Update: refresh after every result; open P3-A only after final speedup/equivalence and terminal-disposition criteria pass.
  - Commit Boundary: verified current fact `E2-P2A-A001COMMIT1` records A001 implementation commit `17a1f3af37dcb61f9d389345822b6470a8f772cc`. It does not close P2-A/Child 06A, check any parent/child, satisfy `E3-P3C-COMMIT1`, or authorize A002 Coder; final closure remains governed separately by P3-C.

### P3: Final Acceptance And Closure

- Phase Goal: review the stable whole candidate, clean it without changing accepted behavior, detect it, commit it once, and hand it to Child 07.
- Phase Boundary:
  - In scope: P3-A, P3-B, and P3-C only.
  - Out of scope: new optimization work, another implementation slice, another final Supervisor boundary, or another commit.
  - Dependencies: P2-A final speedup, per-attempt reviews, final equivalence, and exhaustion evidence.
- Phase Implementation Rule: P3-A is the one final whole-candidate Supervisor; P3-B performs no review; P3-C owns the sole detect/commit/handoff.

- [ ] P3-A: Run The Final Whole-Candidate Supervisor Review.
  - Goal: independently verify stable complete P2-A against current ledgers, source, build, every attempt review, measured initial/final result, and preserved invariants.
  - Work Steps:
    1. Review the exact stable implementation boundary and issue `PASS`, `REJECT`, or a concrete blocker.
    2. On reject, return only the failed measured owner/invariant through a new P2-A current drill-down -> Architect -> Planner -> Coder -> remeasure -> per-attempt Supervisor chain; resubmission stays inside this same final whole-candidate Supervisor boundary.
  - Implementation Gate: every measured parent and child benchmark row has retained terminal evidence, every matching plan checklist item is checked, and `B2-P2A-FINAL1` has lower comparable final total with final equivalence evidence.
  - Acceptance: `E3-P3A-REVIEW1` records the one effective final whole-candidate `PASS` with no unresolved same-scope finding.

- [ ] P3-B: Remove Exact Child 06A Dead Work Without A New Review.
  - Goal: remove only failed, superseded, temporary, or unused work created by Child 06A.
  - Work Steps:
    1. Inventory/remove exact Child 06A-owned dead/debug material; preserve historical/protected/shared/quarantined work.
    2. Prove accepted production/test bytes are unchanged from P3-A. If changed, invalidate P3-A and return to a new exact-owner P2-A attempt before closure.
  - Implementation Gate: `E3-P3A-REVIEW1` is `PASS`; cleanup cannot broaden scope or create documentation review.
  - Acceptance: `E3-P3B-CLEAN1` records removed/preserved/blocked items and accepted production/test identity remains unchanged.

- [ ] P3-C: Detect, Create One Implementation Commit, And Hand Off Child 07.
  - Goal: close Child 06A at one isolated implementation commit and make it Child 07's direct predecessor.
  - Work Steps:
    1. Refresh final ledgers and run post-cleanup `anvien detect-changes --repo E:\Anvien --scope all` on the accepted implementation boundary.
    2. Stage only accepted Child 06A implementation/tests/necessary instrumentation/evidence/benchmark/plan ledgers, create exactly one implementation commit, and verify its manifest.
    3. Update Child 07 opening evidence to the exact Child 06A closure commit and hand off current benchmark/equivalence basis.
  - Implementation Gate: P3-A remains valid, P3-B changed no accepted production/test bytes, detect passes, and staging excludes unrelated/protected work.
  - Acceptance: `E3-P3C-DETECT1`, `E3-P3C-COMMIT1`, and `E3-P3C-HANDOFF1` record one detect, one implementation commit, and direct Child 07 handoff.

## Risk Notes

- An inaccurate operation boundary can optimize the wrong work. Rank only current absolute elapsed timings with proven denominators and boundary evidence.
- Instrumentation overhead can distort rankings. Use existing timing first; if the sole-writer branch opens, require fresh owner impact, full-build PASS before use, like-for-like instrumentation states, output equivalence, and a completed carry-forward or remove/rebuild/remeasure disposition.
- A stale internal ranking or architecture decision can make Coder optimize the wrong current cause. Every production edit requires current attempt-local drill-down evidence, a new attempt-specific Architect decision, and Planner refresh.
- A local child saving can move cost into its parent or elsewhere. Every retained attempt requires child, parent, and end-to-end proof.
- Timing improvement can hide an accuracy regression. Candidate values remain unpromoted until that attempt's independent Supervisor passes the affected equivalence boundary.
- A faster path can silently weaken output, ordering, persistence, readers, determinism, freshness, or failure/publication behavior. Validate the exact affected inherited boundary before `KEEP`.
- Parallel shared-worktree implementation destroys attribution. Measurement parallelism is optional/bounded; production/test changes remain sequential.
- Cleanup after final acceptance can invalidate stable state. Any accepted-byte mutation returns to a new P2-A attempt and the same final P3-A boundary.
- Historical profile values may suggest investigation but cannot outrank current absolute measurements or become a baseline.
