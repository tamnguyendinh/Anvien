# Child 06A Accelerate Analyze Without Sacrificing Accuracy Plan

## Metadata

- Date: `2026-08-24`
- Status: `P0-A complete / P1-A complete / P2-A open / OWNER_COST_FLOOR_APPLIED / TOP_LEVEL_25_OF_30_TERMINAL / OP001_CHILDREN_16_OF_17_TERMINAL / D001_EVIDENCE_EXHAUSTED_TERMINAL / D001_STREAK_2 / D002_EVIDENCE_EXHAUSTED_TERMINAL / D003_PASS_BY_OWNER_COST_FLOOR / D004_A007_CODER_COMPLETE / D005-D017_PASS_BY_OWNER_COST_FLOOR / OWNER_FINAL_SCOPE_BOUND / OP002-OP005_DEFER_AFTER_D004 / SAME_CODER_MEASUREMENT_PENDING / CAMPAIGN_CLOSE_AFTER_A007 / CHILD07_NOT_OPENED_BY_OWNER / A003_CHECKPOINT_COMPLETE / TARGET_SEPARATION_PRESERVED`
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
- The latest Owner cost-floor rule in [plan-rules.md](plan-rules.md) supersedes every conflicting statement in this plan that forbids a size/cost cutoff or requires functional processing of every smaller measured row. The complete row remains recorded; strict floor qualification terminalizes it in place rather than omitting it.

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
- one final whole-candidate Supervisor boundary, exact dead-work cleanup, one detect, one Child 06A campaign-closure commit, and terminal campaign closure without Child 07 opening.

Preserve/validate boundary:

- accepted accuracy, semantic completeness, graph correctness, ordering of outcomes/diagnostics/evidence, and fail-closed behavior;
- canonical in-memory graph, Graph JSON, Ladybug/native persistence, and affected native/product readers;
- deterministic replay, source freshness/invalidation, failure propagation, transaction rollback, temporary-file flush/close/rename, and publication visibility;
- exact workload identity and operation-specific work denominators needed for comparable before/after results.

## Non-Goals

- No implementation of a historical idea merely because it appeared in a prior report or legacy plan item.
- No predefined bottleneck, source owner, file list, optimization list, or reusable architecture decision.
- No general measurement framework, permanent observability project, or durable run-by-run report hierarchy.
- No numeric threshold other than the binding Owner cost floor in [plan-rules.md](plan-rules.md); the strict `3.000000000 s` all-target predicate controls current queue terminalization.
- No parallel production editing, per-bottleneck implementation slice, unaccepted/rejected-attempt commit, or documentation acceptance loop. Accepted Supervisor-`PASS` + Main-`KEEP` progress checkpoints follow [plan-rules.md](plan-rules.md) without closing P2-A.
- No promotion based only on timing without the exact attempt's Supervisor accuracy/equivalence `PASS`.
- No Child 06 correctness re-review and no Child 07 opening.
- No terminal Child 07 validation; the latest Owner ends the campaign at the Child 06A closure commit.
- No performance claim from profiled cumulative samples, historical wall times, build/test duration, or an unmeasured resource hypothesis.
- The former prohibition on an elapsed-time size cutoff is superseded by the latest Owner authority. `PASS_BY_OWNER_COST_FLOOR` checks a measured row without omitting it only when every accepted target has a strictly `< 3.000000000 s` controlling wall value; any `>= 3.000000000 s`, missing value, or comparability loss remains open. Ordinary `BLOCKED_OPEN` and binding `EVIDENCE_EXHAUSTED` retain their separate meanings.

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

- Classify every current top-level row from the accepted target-separated packet first. Check each strict all-target floor qualifier directly; retain every other row open in its existing relative order. Continue the already-active OP001 parent until its remaining child is terminal.
- As the first processing step for that parent, measure more deeply inside only its proven runtime boundary, record elapsed time and denominator for every real measured child, and create the complete child list in descending absolute elapsed-time order. Do not drill down every parent upfront or create a separate phase.
- Mirror the exact parent and complete child row IDs/names into the plan checklist. Check each strict all-target floor qualifier directly, then select the largest remaining open child and keep only other nonqualifying children queued.
- Open a new Visible Architect turn for exactly the selected child attempt only after the complete child list exists. The packet contains the parent row, complete current child list, selected child row, current accepted child/parent/total basis, denominator, exact cause/owner/call path, preserved invariants, and any prior rejection evidence.
- Record the Architect's exact direction, allowed owners, expected observable gain, validation/resource boundaries, and rollback condition. The decision expires when the attempt reaches disposition.
- Planner then refreshes the Current P2-A Attempt card and replaces generic steps with concrete files/symbols, production change, tests, build, measurements, equivalence checks, and rollback for that attempt. No Coder begins before this refresh is recorded.
- Coder implements only that refreshed direction after fresh required Anvien graph/file-detail/impact on the exact owner. Production implementation comes first; only after production behavior is correct may Coder edit the authorized test files.
- Run the exact build-holder/process preflight and canonical full build. Only after full-build `PASS` run the focused tests and `internal/resolution` package tests as validation evidence.
- Later and separately, remeasure the selected child with the same denominator, remeasure its parent, and remeasure total graph-generation time on the same workload.
- Open a new Visible Supervisor check for that attempt's changed candidate and affected accuracy/correctness/equivalence/output/lifecycle boundary.
- Record one production-attempt disposition only after the Supervisor result: `KEEP`, `REWORK`, `ROLLBACK`, concrete `BLOCKED`, or terminal `SYSTEM_CHARACTERISTIC` after the required three-attempt history. Separately, `EVIDENCE_EXHAUSTED` follows its exact governance proof, while `PASS_BY_OWNER_COST_FLOOR` is a distinct Owner-authorized non-production terminal control state and never a production disposition.
- `KEEP` requires lower selected-child elapsed wall-clock time, lower retained parent elapsed wall-clock time, lower retained end-to-end graph-generation elapsed wall-clock time, and Supervisor `PASS`; only then does the candidate become the current accepted baseline.
- Supervisor `REJECT` forces scoped restoration of the last accepted baseline, a new attempt-local internal drill-down on that accepted state, and a correction packet with the current drill-down back to a new Architect. Rejected bytes cannot become baseline, cannot be ranked, and cannot authorize a different bottleneck. If the unsuccessful streak is below `3`, a new attempt on the same bottleneck is mandatory unless a concrete blocker exists.
- Count an attempt as unsuccessful when it produces no retainable `KEEP`: Supervisor `REJECT`, no retainable direct/end-to-end improvement, or a `REWORK`/`ROLLBACK` outcome. Store the streak against the exact bottleneck row and current accepted baseline.
- On `KEEP`, promote the candidate, reset the selected child's consecutive unsuccessful-attempt count to `0`, remeasure that child/parent/full pipeline, and continue the same child until its remaining retained cost reaches terminal `SYSTEM_CHARACTERISTIC` or a concrete blocker is resolved.
- On the third consecutive unsuccessful attempt, record the exact three attempt/evidence references and reasons, preserve the last accepted correct timing, and terminalize only the selected child as `SYSTEM_CHARACTERISTIC`.
- After a child becomes terminal by `SYSTEM_CHARACTERISTIC`, exact `EVIDENCE_EXHAUSTED`, or strict `PASS_BY_OWNER_COST_FLOOR` evidence, mark only its plan checkbox and select the largest remaining open child under the active parent. An ordinary concrete `BLOCKED_OPEN` child stays unchecked and blocks parent completion.
- Mark OP001 complete only after every measured child is checked with retained terminal evidence. For the final A007 transition, use its retained packet after `KEEP` or the already accepted A003 packet after restoration; do not rerun the accepted baseline solely for closure. Then apply exact Owner-scope deferral to OP002-OP005 rather than opening further optimization work.

### Final equivalence and closure

- Preserve initial total and operation baselines for final comparison even though every retained attempt uses the immediately preceding accepted state as its incremental baseline.
- Final acceptance requires lower measured total graph-generation time on the same workload, no unexplained transfer of cost to another operation, and exact required equivalence across all affected accepted boundaries.
- One stable final state receives the final whole-candidate P3-A Supervisor review; this does not replace or retroactively supply any per-attempt review.
- Cleanup removes only Child 06A-created failed, superseded, or debug work and must not mutate accepted production/test bytes.
- One post-cleanup detect and one isolated implementation commit close Child 06A and the campaign; Child 07 remains unopened by Owner authority.

## Acceptance Criteria

- P6-D commit `81163e39718b94a509e41114cada224e8f269e36` remains the immutable direct predecessor.
- P1-A records a valid current total graph-generation time, real internal operation timings, denominators, and an ordered absolute bottleneck ranking without production optimization.
- If P1-A uses timing instrumentation, `E1-P1A-INSTR1` proves the sole visible writer, exact owner/impact, full-build PASS before timing use, runtime/instrumentation identity, like-for-like overhead/comparability, output equivalence, and exactly one completed carry-forward or remove/rebuild/remeasure disposition; no instrumentation branch remains half-open.
- The benchmark living table contains current rank, real operation/bottleneck, current elapsed time, denominator, initial baseline, current before/after, meaningful share/delta, actionability, current attempt, direct delta, end-to-end before/after, cumulative initial delta, consecutive unsuccessful-attempt count, disposition, proven owner/call path when available, and exact evidence.
- P1-A cannot close and P2-A cannot open until benchmark contains the complete ordered top-level bottleneck list for every currently measured real operation and plan contains one unchecked parent checklist row for each benchmark parent. All rows remain mandatory control records; latest Owner scope changes OP002-OP005 from future Child 06A optimization work into post-A007 deferrals without removing them.
- P2-A is the only implementation slice. It runs A007 as the one final D004 production chain; after its safe Supervisor/Main disposition, D004 and OP001 close and OP002-OP005 are terminalized only as `DEFERRED_BY_OWNER_SCOPE`. No OP002-OP005 optimization lane opens in Child 06A.
- Before the first child attempt of a parent, P2-A creates the complete measured child list, mirrors every child into the nested plan checklist, selects the largest unchecked child, and proves its cause/owner/complete call path; no coarse parent name or incomplete child inventory can authorize architecture or code.
- Every production attempt has a distinct fresh Visible Architect decision, a recorded Planner refresh of the current P2-A attempt before Coder, and a distinct post-remeasurement Visible Supervisor accuracy/equivalence result.
- Every retained production change has exact parent/child cause/owner/impact evidence, production-first implementation, authorized test edits only after production is correct, canonical full build, post-build focused/package test execution, child before/after, parent before/after, end-to-end before/after, affected equivalence, Supervisor `PASS`, and an immediate `KEEP` disposition.
- Only `KEEP` attempts are promoted; each promoted state becomes the next current baseline and resets the row's unsuccessful streak. Rejected candidates are restored without ranking or transition on rejected bytes.
- The same child continues through new attempts after each `KEEP`. Three consecutive attempts without `KEEP` at one accepted baseline terminalize only that child as `SYSTEM_CHARACTERISTIC`; independently, exact binding `EVIDENCE_EXHAUSTED` proof may terminalize only that measured child without manufacturing a third attempt or changing its streak. The next-largest unchecked child of the same parent follows.
- Every measured child of every measured parent has either retained production/governance terminal evidence or `PASS_BY_OWNER_COST_FLOOR` packet proof and a checked plan item; every open parent is checked only after all its children are checked. No row is omitted.
- Final total graph-generation time is lower than the initial comparable value on the same workload.
- Accuracy, semantic completeness, graph correctness, ordered output, determinism, freshness, failure, transaction, temp, publication, Graph JSON, Ladybug, and native-reader invariants pass at every affected boundary.
- Exactly one final whole-candidate `P3-A` Supervisor boundary passes, `P3-B` cleanup leaves accepted production/test bytes unchanged, and `P3-C` records one detect, one campaign-closure commit, and terminal `CAMPAIGN_CLOSED / CHILD07_NOT_OPENED_BY_OWNER` status.
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

Current checklist entries: `30` top-level rows plus `17` nested child rows under active parent `B1-P1A-OP001`. `E2-P2A-OWNERCOSTFLOOR1` and `B2-P2A-OWNERCOSTFLOOR1` classify the exact accepted A003 packets: top-level checked count is `25/30` and OP001 checked-child count is `16/17`. D001 remains terminal `EVIDENCE_EXHAUSTED` with streak exactly `2`; D002 remains terminal `EVIDENCE_EXHAUSTED` with no streak. D003 and D005-D017 are terminal `PASS_BY_OWNER_COST_FLOOR`. D004 is the only open OP001 child; A007 is its one final production chain and is Planner-ready for exactly one Coder. OP001 remains unchecked. OP002-OP005 remain numerically present and unchecked until A007 has a safe Supervisor/Main disposition, at which point they become `DEFERRED_BY_OWNER_SCOPE` rather than optimization results. Accepted A003 values and target separation are unchanged.

- [ ] `B1-P1A-OP001` — `resolution`
  - [x] `B2-P2A-A001-D001` — `resolve_calls` — `EVIDENCE_EXHAUSTED`
  - [x] `B2-P2A-A001-D002` — `resolve_accesses` — `EVIDENCE_EXHAUSTED`
  - [x] `B2-P2A-A001-D003` — `resolve_type_annotations` — `PASS_BY_OWNER_COST_FLOOR`
  - [ ] `B2-P2A-A001-D004` — `project_resolution_outcomes`
  - [x] `B2-P2A-A001-D005` — `emit_definition_nodes` — `PASS_BY_OWNER_COST_FLOOR`
  - [x] `B2-P2A-A001-D006` — `finalize_resolution_outcomes` — `PASS_BY_OWNER_COST_FLOOR`
  - [x] `B2-P2A-A001-D007` — `build_binding_occurrence_index` — `PASS_BY_OWNER_COST_FLOOR`
  - [x] `B2-P2A-A001-D008` — `finalize_typescript_authority_results` — `PASS_BY_OWNER_COST_FLOOR`
  - [x] `B2-P2A-A001-D009` — `emit_import_edges` — `PASS_BY_OWNER_COST_FLOOR`
  - [x] `B2-P2A-A001-D010` — `emit_typescript_external_symbols` — `PASS_BY_OWNER_COST_FLOOR`
  - [x] `B2-P2A-A001-D011` — `emit_method_dispatch_edges` — `PASS_BY_OWNER_COST_FLOOR`
  - [x] `B2-P2A-A001-D012` — `assemble_resolution_result` — `PASS_BY_OWNER_COST_FLOOR`
  - [x] `B2-P2A-A001-D013` — `binding_accumulator_dispose` — `PASS_BY_OWNER_COST_FLOOR`
  - [x] `B2-P2A-A001-D014` — `emit_heritage_compatibility_edges` — `PASS_BY_OWNER_COST_FLOOR`
  - [x] `B2-P2A-A001-D015` — `emit_unresolved_heritage_diagnostics` — `PASS_BY_OWNER_COST_FLOOR`
  - [x] `B2-P2A-A001-D016` — `finalize_resolution_metadata` — `PASS_BY_OWNER_COST_FLOOR`
  - [x] `B2-P2A-A001-D017` — `runtime_setup` — `PASS_BY_OWNER_COST_FLOOR`
- [ ] `B1-P1A-OP002` — `db_load` — pending `DEFERRED_BY_OWNER_SCOPE` after A007 disposition
- [ ] `B1-P1A-OP003` — `parse` — pending `DEFERRED_BY_OWNER_SCOPE` after A007 disposition
- [ ] `B1-P1A-OP004` — `graph_snapshot` — pending `DEFERRED_BY_OWNER_SCOPE` after A007 disposition
- [ ] `B1-P1A-OP005` — `semantic_enrichment` — pending `DEFERRED_BY_OWNER_SCOPE` after A007 disposition
- [x] `B1-P1A-OP006` — `db_runner_resolve` — `PASS_BY_OWNER_COST_FLOOR`
- [x] `B1-P1A-OP007` — `cross_file_binding` — `PASS_BY_OWNER_COST_FLOOR`
- [x] `B1-P1A-OP008` — `ai_context` — `PASS_BY_OWNER_COST_FLOOR`
- [x] `B1-P1A-OP009` — `analyzer_orchestration` — `PASS_BY_OWNER_COST_FLOOR`
- [x] `B1-P1A-OP010` — `processes` — `PASS_BY_OWNER_COST_FLOOR`
- [x] `B1-P1A-OP011` — `scan` — `PASS_BY_OWNER_COST_FLOOR`
- [x] `B1-P1A-OP012` — `file_projection` — `PASS_BY_OWNER_COST_FLOOR`
- [x] `B1-P1A-OP013` — `routes` — `PASS_BY_OWNER_COST_FLOOR`
- [x] `B1-P1A-OP014` — `documents` — `PASS_BY_OWNER_COST_FLOOR`
- [x] `B1-P1A-OP015` — `db_runner_close` — `PASS_BY_OWNER_COST_FLOOR`
- [x] `B1-P1A-OP016` — `analyze_setup` — `PASS_BY_OWNER_COST_FLOOR`
- [x] `B1-P1A-OP017` — `tools` — `PASS_BY_OWNER_COST_FLOOR`
- [x] `B1-P1A-OP018` — `communities` — `PASS_BY_OWNER_COST_FLOOR`
- [x] `B1-P1A-OP019` — `registry_meta` — `PASS_BY_OWNER_COST_FLOOR`
- [x] `B1-P1A-OP020` — `cli_preparation` — `PASS_BY_OWNER_COST_FLOOR`
- [x] `B1-P1A-OP021` — `output_publication` — `PASS_BY_OWNER_COST_FLOOR`
- [x] `B1-P1A-OP022` — `orm` — `PASS_BY_OWNER_COST_FLOOR`
- [x] `B1-P1A-OP023` — `cli_startup` — `PASS_BY_OWNER_COST_FLOOR`
- [x] `B1-P1A-OP024` — `structure` — `PASS_BY_OWNER_COST_FLOOR`
- [x] `B1-P1A-OP025` — `graph_compact` — `PASS_BY_OWNER_COST_FLOOR`
- [x] `B1-P1A-OP026` — `mro` — `PASS_BY_OWNER_COST_FLOOR`
- [x] `B1-P1A-OP027` — `benchmark_write` — `PASS_BY_OWNER_COST_FLOOR`
- [x] `B1-P1A-OP028` — `cobol` — `PASS_BY_OWNER_COST_FLOOR`
- [x] `B1-P1A-OP029` — `memory_profile` — `PASS_BY_OWNER_COST_FLOOR`
- [x] `B1-P1A-OP030` — `cpu_profile_completion` — `PASS_BY_OWNER_COST_FLOOR`

Population and marking rules:

- Before P1-A closes, insert one unchecked top-level parent checklist item for every measured `B1-P1A-OPnnn` row, in the same descending elapsed-time order as benchmark. Each item names its exact benchmark row ID and real operation.
- When P2-A selects a parent and completes drill-down, insert one nested unchecked child checklist item for every measured child row under that parent, in the same descending elapsed-time order as the complete benchmark child list. Do this before opening the first child Architect attempt.
- Cardinality must match exactly: if one parent drill-down measures `10` child bottlenecks, benchmark must contain exactly `10` child rows for that parent and this plan must contain exactly `10` nested child checkbox items before Architect. A missing or extra row/item blocks the attempt.
- Do not insert generic candidate, filename, solution, or empty placeholder items. Every checkbox must come from a real measured benchmark row.
- Keep the active child unchecked after `KEEP`; the retained baseline improves, its no-KEEP counter resets, and optimization continues on that same child.
- Check a child after retained `SYSTEM_CHARACTERISTIC`, exact five-condition `EVIDENCE_EXHAUSTED`, or strict all-target `PASS_BY_OWNER_COST_FLOOR` proof. An ordinary unavailable authority/dependency/evidence `BLOCKED_OPEN` row remains unchecked and blocks parent completion.
- The latest Owner cost-floor rule supersedes all earlier checklist wording that forbids size/cost terminalization. It checks qualifying rows in place, preserves their row identity and accepted values, and does not count as production processing or an attempt.
- The latest Owner final-scope rule makes A007 the last Child 06A optimization chain. After its safe disposition, check D004 and OP001 as `CLOSED_BY_OWNER_SCOPE`, then check OP002-OP005 as `DEFERRED_BY_OWNER_SCOPE`; preserve their accepted numbers and never describe the deferral as `PASS`, `KEEP`, cost-floor, speedup, or correctness acceptance.
- Check a parent only after every measured child beneath it is checked and parent/full-pipeline remeasurement is recorded. If later accepted measurement exposes a previously missing child, reopen that parent, append the real child row, and continue P2-A.
- If later accepted measurement exposes a previously missing top-level parent or child, append its real benchmark row and matching unchecked plan checkbox immediately. A new parent joins the remaining top-level queue; a new child reopens/remains under its parent. No discovered row may exist in only one ledger.
- P2-A and `E2-P2A-EXHAUST1` cannot close until every measured row is checked either from retained technical terminal evidence or the exact Owner-scope terminal controls above; no row may be omitted.

## Current P2-A Attempt

This is the living Planner refresh surface for the one implementation slice. Planner replaces this row and the matching concrete attempt steps after each fresh Architect decision and before any Coder edit. Metrics remain in benchmark.

| Field | Current value |
|-------|---------------|
| Attempt state | `D004_A007_PLAN_READY / CODER_RELEASE_PENDING / OWNER_FINAL_SCOPE_BOUND`; accepted A003 remains baseline; parent and D004 remain unchecked |
| Attempt ID | `A007` — the one final D004 production chain; it has architecture/Planner authority but no candidate measurement, numeric history row, streak event, or disposition yet |
| Attempt goal | remove the redundant diagnostic-property JSON round-trip only on the exact canonical in-memory representation, retain exact compatibility fallback and fail-closed behavior, and prove lower D004/parent/process wall on both targets without accuracy, analyzer, output, persistence, reader, order, failure, or lifecycle drift |
| Active parent benchmark row / checklist item | `B1-P1A-OP001` / unchecked `resolution` parent item |
| Complete parent child list | `17/17` rows `B2-P2A-A001-D001..D017` remain target-separated; checked-child count is `16/17`. D001/D002 retain their prior terminal history; D003 and D005-D017 are floor-PASS; only D004 is open |
| Active child benchmark row / checklist item | `B2-P2A-A001-D004` / unchecked `project_resolution_outcomes` |
| Remaining child queue | none beyond active D004; every other OP001 child is terminal |
| Current benchmark authority | `B2-P2A-OWNERCOSTFLOOR1` plus the target-separated accepted A003 D004 basis; no accepted before is rerun, averaged, combined, or replaced before A007 measurement |
| Consecutive unsuccessful attempts | D001 terminal record retains exactly `2`. D002, D003, and D004 have no production-attempt or streak event; no streak transfers between children |
| Selected child exact cause / owner / complete call path | `E2-P2A-D004ATTRIB1`: `internal/resolution/outcome.go::resolutionOutcomeDiagnosticSites` owns a graph-wide diagnostic-property JSON marshal/unmarshal round-trip before exact status/byte parity checks; full path is `internal/cli.newAnalyzeCommand.func1 -> analyze.Run -> runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> projectResolutionOutcomes -> resolutionOutcomeDiagnosticSites -> encoding/json -> parity checks -> return` |
| Fresh Architect decision | `E2-P2A-A007ARCH1`: report `E:\Anvien\reports\system-architect\rp_system-architect_260827_by_gpt-5_child06a_d004_diagnostic_roundtrip.md`, SHA-256 `834D4A0C9297602C12B9EF9C55C4149382743B0103A0E592CA830A58C5AA8D21`, verdict `ARCHITECT_D004_READY_FOR_PLANNER` |
| Planner refresh evidence | `E2-P2A-A007PLAN1`; exact translation of the Architect direction plus `E2-P2A-OWNERSCOPE1`, with no numeric/checklist/streak mutation or Coder result |
| Transition authority | `E2-P2A-A007ARCH1/A007BOUNDARY1/A007PLAN1` and latest Owner final-scope authority; Main boundary check is identity/scope/handoff only, not a Supervisor acceptance verdict |
| Future pre-edit Anvien gate | sole Coder must run `anvien --help`, one fresh `anvien analyze --force`, file-detail `internal/resolution/outcome.go`, and upstream impact for every edited symbol immediately before edit; retain full HIGH/CRITICAL counts as scope warnings |
| Allowed production/test surfaces | production: only `internal/resolution/outcome.go::resolutionOutcomeDiagnosticSites`, at most one private non-exported read-only helper and minimal import in that file; test after production: create only `internal/resolution/outcome_diagnostic_sites_test.go` |
| Production algorithm | guarded fast path for exact `[]graphhealth.Diagnostic`: directly iterate only when consumed `SourceSiteID`/`Note` strings are JSON-equivalent; every other or JSON-unstable representation executes the exact existing marshal/unmarshal fallback and error path; actual graph property remains the source-of-truth and every current parity check/order stays unchanged |
| Test contract | new focused file proves canonical result/non-mutation, ignored untracked diagnostics, resolved+diagnostic and payload-drift failures, traversal/error-first order, legacy `[]any`/`map[string]any` fallback, unsupported/unmarshalable errors, and JSON-unstable typed-string fallback equivalence; existing tests are run-only |
| Build/test contract | clear only proven build holders, canonical full build must exit `0` before tests, then run the new focused test, the eight exact named diagnostic/P6C3/P6D/Graph JSON/Ladybug regressions in `E2-P2A-A007ARCH1`, and packages `internal/resolution`, `internal/graphhealth`, `internal/analyze`, `internal/lbugload`, `internal/lbugnative`; known preserve-only golden remains truthful and any new/changed failure blocks A007 |
| Attribution contract | complete under `E2-P2A-D004ATTRIB1`: same residual owner is material on both accepted A003 profiles, A005-controlled projection marshal is excluded as a rejected sibling, target separation is preserved, and no new measurement was run |
| Measurement contract | freeze one candidate with accepted A003 17-child/native/runtime/build contract; launch exactly one Cheapapp and one Restaurant candidate run against each target's own accepted A003 before; require D004/parent/analyzer/process wall, complete `30/30` + ordered `17/17`, exact denominators/conservation/zero overlap, Graph/output/DB/reader parity, resources, hashes, target identity, and one-launch/exit proof |
| Post-measurement Supervisor | exactly one fresh visible A007 Supervisor reviews the candidate and affected correctness/equivalence/output/resource/lifecycle boundary; Main alone owns `KEEP` versus restoration/non-KEEP and Owner-scope closure |
| Resource/invariant boundary | canonical fast path retains only `O(1)` borrowed slice/locals, does not copy/mutate/reorder diagnostics, create sidecars/caches/state/concurrency/I/O, or change public/persisted shapes; noncanonical fallback retains current allocation and fail-closed semantics |
| Exact rollback | only the A007 hunk/helper/import in `outcome.go`, the new D004 test, and its frozen overlay/candidate packet; accepted A003/A001-A006/protected bytes remain untouched |
| Mandatory STOP | any other production file/symbol, existing-test edit, signature/coordinator/graphhealth/collector/resolve change, fallback weakening, unsupported coercion, A005 lifecycle reuse, graph/output/order/error/resource drift, build/new test failure, or mixed/incomplete target packet returns to Main without broadening |
| Owner closure control | after the A007 Supervisor verdict and safe Main promotion/restoration, do not open another D004 attempt. Mark D004/OP001 `CLOSED_BY_OWNER_SCOPE`, mark OP002-OP005 `DEFERRED_BY_OWNER_SCOPE`, preserve all numbers, then proceed directly to lean P3 closure |
| Coder status | `CODER_A007_READY_FOR_MEASUREMENT_HANDOFF`; exact source/test/report hashes and validation truth are recorded under `E2-P2A-A007IMPACT1/SRC1/BUILD1/TEST1/CODERBOUNDARY1` |
| Measurement owner | latest Owner authority resumes the same visible Coder task `01a042fd-4efa-7880-ada8-66ea8ec5871f`; no separate builder or measurement lane opens |
| Next action | checkpoint the Coder report plus five plan files without staging source/test, then resume that same Coder to freeze one A007 candidate and run Cheapapp followed by Restaurant Manager; Supervisor remains locked until both packets exist |

### A005 PLAN1 Canonical Outcome-Byte Ownership Execution Authority

- Authority: `E2-P2A-A005ATTRIB1/ARCH1` and report `E:\Anvien\reports\system-architect\rp_system-architect_260827_by_gpt-5_child06a_a005_outcome_serialization.md`, verdict `ARCHITECT_A005_READY_FOR_PLANNER`; Main Architect-handoff verification `E:\Anvien\reports\Supervisor\rp_supervisor_260827_021244_by_gpt-5_child06a_a005_architect_handoff.md` is `PASS`. `E2-P2A-A005PLAN1` is the exact execution translation; `E2-P2A-A005MAINVERIFY1` and `E:\Anvien\reports\Supervisor\rp_supervisor_260827_025255_by_gpt-5_child06a_a005_planner_handoff.md` accept it and release exactly one future visible Coder.
- Locked byte lifecycle: `clone -> validate -> conflict/equal-duplicate decision -> unique semantic retention -> unchanged record-time marshal once -> private SourceSiteID canonical-string sidecar -> private finalized bundle -> immediate diagnostic and final relationship/reference/diagnostic parity consumers`. Projection validates/consumes and never encodes or falls back.
- Production-first boundary: only `internal/resolution/outcome.go` private collector sidecar/init, `record` reuse semantics with unchanged signature, one private finalized bundle, `finalize` coverage/return, and strict `projectResolutionOutcomes` consumption; plus only the existing `ResolveBoundInto` finalize/project/result wiring block in `internal/resolution/resolve.go`. `marshalResolutionOutcome` body/signature, validation/clone, outcome constructors/diagnostic emitters, graph/public/persistence/readers, instrumentation, scripts, A001-A004, and D002-D017 are preserve-only.
- Tests-after-production boundary: create only `internal/resolution/outcome_serialization_test.go`; mechanically adapt only three direct private `finalize` calls and one direct `projectResolutionOutcomes` call in `internal/resolution/p6c3_structured_outcome_test.go`. Cover every status family, exact diagnostic/evidence byte parity, duplicate/conflict/first-error/`added`, deep clone immutability, record-time marshal failure, missing/extra sidecar coverage, rejected overlap/drift/duplicate sites, SourceSiteID order, complete ordered evidence, and determinism.
- Validation order: fresh future Coder Anvien/file-detail/exact impacts -> production -> production diff inspection/owner STOP -> tests -> holder/lock/process clearance including both `anvien doctor` commands -> canonical full build PASS -> A005 focused plus named P6B/P6C3/P6D/Graph JSON/Ladybug tests -> five package regressions -> scoped diff-check. The known golden is never called PASS; any new or changed failure blocks A005. No detect/stage/commit occurs before later measurements, Supervisor, and Main disposition.
- Candidate/measurement order: reuse the unchanged A00x script, regenerate/verify only the overlay-mapped A005 `resolve.go` wiring with the accepted 17-child instrumentation, preserve `types.go` unless exact compile evidence says otherwise, and exclude rejected A004 bytes. Cheapapp and Restaurant each run once against their own accepted A003 packet and remain separate; each packet proves the complete elapsed, `30/30`, `17/17`, denominator, interval, evidence-order, graph/DB/output/semantic/resource, and private `O(U+B)` lifecycle boundary.
- Acceptance order: only after both packets exist, one fresh visible A005 Supervisor reviews the exact candidate and affected correctness/equivalence/output/lifecycle/resource surface. Main alone owns `KEEP`, `REWORK`, `ROLLBACK`, and streak effect; an unsuccessful A005 can produce only streak `2` from the current streak `1`.
- Rollback/STOP: rollback only exact A005 bytes and its rejected frozen/overlay packet. Any added owner, alternate/fallback encoder, moved record-time semantics/error timing, byte/order/carrier drift, public/persisted shape change, payload duplication/cross-run state, changed build/test/golden result, stale/mixed overlay, unavailable independent packet, or D002-D017/P3/Child 07 opening returns to Main.

### A004 PLAN1 Export-Binding Evidence Ordering Execution Authority

- Authority: `E2-P2A-A004ARCH1` and Owner-approved report `E:\Anvien\reports\system-architect\rp_system-architect_260826_by_gpt-5_child06a_a004_residual_direction.md`, exact verdict `ARCHITECT_A004_OWNER_APPROVED_READY_FOR_PLANNER`. `E2-P2A-A004PLAN1` is the only A004 execution translation; it does not release Coder.
- Locked algorithm: `exact-tuple dedupe -> derive the canonical order key from final serialized Evidence.Note exactly once per unique projected export-binding evidence record -> exactly one final stable sort using cached keys -> discard transient keys -> return byte/order-identical []graph.Evidence`.
- Production-first boundary: change only `internal/resolution/export_binding_proof.go` and only the four named helpers/private key machinery plus at most one private transient decorated representation/helper. Remove the projection pre-sort before changing tests. Keep every signature and public/persisted shape unchanged.
- Tests-after-production boundary: append only A004-specific bytes to `internal/resolution/export_binding_proof_test.go`. A test-local pre-A004 oracle compares the complete ordered slice and includes the full Architect adversarial matrix, input non-mutation, repeated-coalescing idempotency, malformed-note fallback, and equal-key stability.
- Validation order: fresh Coder Anvien graph -> file-detail and exact impact for every edited helper -> production -> tests -> build-holder/process clearance -> canonical full build PASS -> focused differential/export-binding/call/access/outcome/Graph JSON regressions -> full `internal/resolution` with the known golden recorded truthfully. Any new or changed failure blocks A004.
- Candidate/measurement order: invoke the unchanged A00x build contract with explicit A004 inputs and provenance; independently remeasure Cheapapp and Restaurant Manager against their own accepted A003 bases; record the complete child/parent/analyzer/process, `30/30`, `17/17`, denominator, conservation, exact evidence ordering, Graph JSON, stdout/stderr, semantic, graph/DB, and resource packets without aggregation.
- Acceptance order: only after both packets exist, open one fresh visible A004 Supervisor. Main/Owner owns final disposition. Planning changes no benchmark value, streak, checkbox, queue, or accepted baseline.
- Rollback/STOP: rollback only exact A004-owned production/test bytes. Any expanded production owner, alternate/persisted key authority, comparator JSON decode, second final projected-evidence sort, retained cache/concurrency, malformed-evidence behavior change, D002-D017 opening, parity gap, or missing independent target packet returns to Owner.

### Child 06A WAL Force Cleanup Fix Result — Checkpoint Complete

- Confirmed root cause: `--force` removed only the main database and Graph JSON while Ladybug's physical generation also owned exactly `.wal`, `.lock`, `.wal.checkpoint`, `.checkpoint.apply.lock`, and `.checkpoint.intent.lock`; retained checkpoint sidecars could make a new main database fail with generic native state `1`.
- Accepted production ownership: `internal/lbugruntime/wal_recovery.go` owns the explicit five-sidecar inventory and full `RemoveDatabaseArtifacts` reset; `internal/analyze/analyze.go::prepareStorage(force)` decides reset timing and delegates the physical database family without knowing suffixes. WAL recovery reuses sidecar-only cleanup after its existing explicit classifier, preserves the main DB, and retries once.
- Accepted tests: `internal/lbugruntime/extensions_retry_test.go` proves exact inventory, full-generation reset, sidecar-only WAL recovery, and generic state `1` non-recovery; `internal/analyze/analyze_test.go::TestRunForceRemovesPreviousLbugOutput` proves force removes the main DB generation and all five sidecars while preserving existing Graph JSON assertions.
- Validation result: exact four-file diff `+81/-11`; scoped diff-check PASS; canonical full build exit `0`; focused lbugruntime/analyze tests PASS; rebuilt-runtime five-sidecar repro exits `0`, creates new DB/graph/meta, and leaves all five stale sidecars absent.
- Independent review: task `01a03e75-c596-7e80-8880-698e3bce145a` returned `SUPERVISOR_CHILD06A_WAL_FORCE_FIX_PASS`. No target benchmark/profile/A003 rerun occurred; generic native state-1 diagnostic plumbing remains outside this fix.
- Checkpoint: `0f3a572331dd23d17688886fcbfebeb7d37ee35d`, subject `fix(storage): reset Ladybug artifact family on force`, exact `11`-path manifest, staged set empty after commit. A004 resumed immediately; no additional review or audit gate opened.
- Preserve `E:\Anvien\.tmp\lbug-recovery-20260826-a003`. Repro/control roots remain debug-only and are cleaned only after acceptance.

### A003 PLAN2 Canonical Single-Interpretation Execution Authority

- Provenance: visible Planner task `01a03d1b-662d-7f90-9a60-38271a10c077` consumes Architect task `01a03cea-15cc-7d73-bf07-a984c69517d9`, report `E:\Anvien\reports\system-architect\rp_system-architect_260826_141200_by_gpt-5_child06a_a003_residual_direction.md`, and verdict `ARCHITECT_A003_READY_FOR_PLANNER` without architecture re-audit. `E2-P2A-A003PLAN1` is preserved as a superseded Main-authored draft; only `E2-P2A-A003PLAN2` is execution authority.
- Exact production owner: only `internal/graphhealth/diagnostics.go::decodeStructuredResolutionOutcome`. If required, add at most one private non-exported decode-result/wire representation in the same file, exclusively owned by that decoder. Preserve its external signature and every other existing semantic owner/file.
- Canonical algorithm: for each Diagnostic, perform one parse/traversal of the outer `Note`. That single interpretation must simultaneously capture exact case-sensitive top-level `schemaVersion`/`status` presence, populate the typed structured outcome/wire fields, and retain syntax/type/validation error state. Combine those products with the existing `SourceSiteStatus` evidence in one semantic decision that yields exactly `UNSTRUCTURED`, `STRUCTURED_VALID`, or `STRUCTURED_INVALID`; then one existing policy decision leads to one graph write. There is no production primary/recovery/legacy/retry/fallback decoder or second full-note authority. Nested `Authority` subdocument decoding remains unchanged.
- External parity/fail-closed contract: preserve the exact externally observed `(outcome, structured, valid)` tuple and diagnostic classification/actionability for every existing input class. Any structured evidence from `SourceSiteStatus` or exact top-level marker presence plus any syntax, type, required-field, proof, authority, or contract error must remain `STRUCTURED_INVALID` and fail closed through the existing `unclassified/review` behavior; it must never be reinterpreted as unstructured.
- Production-first test owner: after production is correct, append only A003-owned cases to existing `internal/graphhealth/diagnostics_test.go`, preserving every accepted A002 byte. Use a test-local pre-A003 oracle, never a production fallback, to compare full tuple and policy parity for structured status evidence, exact top-level marker evidence, both evidence sources, valid markerless structured notes, unstructured notes, malformed JSON, non-object JSON, `null`, typed-field errors, missing/exact-case/case-variant markers, duplicate keys, unknown fields, conflicting evidence, and invalid proof/authority data. `compute_test.go` and `p6d_outcome_projection_test.go` are run-only.
- Fresh Coder Anvien gate: `anvien --help` -> `anvien analyze --force` -> `anvien file-detail E:\Anvien\internal\graphhealth\diagnostics.go --repo E:\Anvien --json` -> exact upstream impact for `decodeStructuredResolutionOutcome` -> record the containing-file HIGH and exact-symbol blast radius. STOP before edit if graph/source identity or the owner boundary cannot be proven.
- Build/test order: perform holder/lock preflight and terminate every proven build-output holder completely; then run `pwsh -NoLogo -NoProfile -NonInteractive -ExecutionPolicy Bypass -File E:\Anvien\scripts\full-build.ps1`. Only after full-build PASS run the new planned `TestDecodeStructuredResolutionOutcomeCanonicalSingleInterpretationParity`, `TestDiagnosticAppenderMatchesLegacySemantics`, `TestP6DStructuredOutcomeProjectionPreservesGraphJSONAndHealthParity`, `TestResolveAttachesSourceBackedUnresolvedDiagnostics`, `TestP6C3FinalOutcomeStatusPrecedenceAndCarriageMatrix`, and `TestP6BResolverCarriesCatalogValidationFailuresPerSourceSite`; then `go test ./internal/graphhealth -count=1` and truthfully `go test ./internal/resolution -count=1`. Only the unchanged preserve-only `TestProofBasedCallAccessGoldenCorpus` failure is allowable, and the package is never called PASS; any new or changed failure blocks A003.
- Measurement: reuse unchanged `scripts/build-a00x-benchmark.ps1` and the accepted 17-child overlay/native/runtime/provenance contract; no build-interface audit or script edit. Candidate identity is accepted A002 checkpoint plus only A003-owned production/test bytes. Measure Cheapapp and Restaurant Manager independently against their own accepted A002 baselines; for each record D001, parent, analyzer, process, `30/30` operations, `17/17` children, exact denominators, graph/readback/output/semantic parity, and resource fields. Never average/combine targets or substitute build/test/profile time for elapsed boundaries.
- Acceptance owner: only after both packets exist, open one fresh visible A003 Supervisor for the exact candidate and affected tuple/policy/graph/output/lifecycle/resource boundary. Main alone decides disposition; Planner claims no speedup and changes no checkbox.
- Exact rollback: remove only the A003 canonical-decoder hunk, any exclusively owned private representation, and A003-appended test bytes. Preserve the entire accepted A002 checkpoint, all accepted A002 test bytes, A001/P1 work, script, reports, ledgers, and unrelated/user/protected changes.
- Mandatory STOP: second decoder/path or production direction; any other owner/file/shared or public contract/cache/emitter/policy/authority decoder/persistence/reader/instrumentation/D002-D017 surface; weakened fail-closed semantics; new validation failure; denominator/workload mismatch; inability to prove tuple/policy/graph/output/resource parity; or unavailable independent two-target evidence.

### A003 Owner KEEP Exact Candidate Restoration Result

- Disposition: Owner explicitly selects the exact already-reviewed A003 candidate for `KEEP`. Cheapapp D001 `3.090914200 -> 3.447846300 s` and parent `19.040468000 -> 20.472602300 s` remain objective measured variation. The A003-only decision accepts that local up/down variation in context because Cheapapp process improves `136.729876000 -> 95.630648200 s`, Restaurant D001/parent/process improve `9.909636600 -> 9.401585300 / 21.242055400 -> 20.850792800 / 145.066210900 -> 101.096911900 s`, and the exact verdict is `SUPERVISOR_A003_PASS`. No numeric tolerance or A004+ precedent is created.
- Current materialization: exact A003 source/test restoration is complete at SHA-256 `6DE54D97A1B95E877B686DC3459238598F6060D28CE6329357A80BD7FC376D30` and `06677DBA4FE9EDA4FE8651E08B93ABC8659129A61355E34BFFA10380AE429371`; checkpoint `b6bf45bce95323aa6b53b182edfea8628bd8b463` is complete. State is `OWNER_KEEP / RESTORE_COMPLETE / A003_CHECKPOINT_COMPLETE / A004 ARCHITECT_PENDING / D001_STREAK_0`; every checkbox remains unchanged.
- Exact recovery source: `C:\Users\TAM NGUYEN\.codex\sessions\2026\08\26\rollout-2026-08-26T15-34-20-01a03d34-b740-7982-8ff7-c2fec47940a6.jsonl`. The restoration Coder must verify A002 start hashes, decode original patch payloads at 1-based lines `433`, `461`, and `489`, and submit them through `apply_patch` in that exact order. Prior successful application records are at lines `434`, `462`, and `490`. Patch bodies are not copied into this plan and must not be manually recreated.
- Start-hash gate: `internal/graphhealth/diagnostics.go` must be `AEABA8A541D1C293DFA4FE411253A85A26B6AC0DD18047AFA6F9151B4932BEE8`; `internal/graphhealth/diagnostics_test.go` must be `58E8CAE2C4EBBF4672D2D338CC5E1419B7CECA9F7389B0E24F51BF2F975A7CEA`.
- Final identity gate: only the two graphhealth paths may change; exact diff is `diagnostics.go +109/-7` and `diagnostics_test.go +182/-0`; final SHA-256 must be `6DE54D97A1B95E877B686DC3459238598F6060D28CE6329357A80BD7FC376D30` and `06677DBA4FE9EDA4FE8651E08B93ABC8659129A61355E34BFFA10380AE429371`. Prove identity against the existing Coder report, frozen packet/provenance, and Supervisor report; no new design or behavior inference is permitted.
- Validation/commit boundary: exact candidate identity is proven. Do not rerun build, tests, target measurements, benchmark, profile, or Supervisor. Main alone completes required implementation detect/staging/commit and then opens fresh A004 Architect.
- Restoration stop conditions are closed by exact final-hash identity. A004 Architect remains pending until the A003 checkpoint; D002-D017, P3, and Child 07 remain unopened.

### A002 Production Direction And Exact Owner Boundary

- Architecture authority: `E2-P2A-A002ARCH1`; report `E:\Anvien\reports\system-architect\rp_system-architect_260825_184751_by_gpt-5_child06a_a002_run_scoped_diagnostic_appender.md`; report-only commit `dafcf7a25e254ef8db09d28eca3575d5225c553e`; verdict `ARCHITECT_A002_READY_FOR_PLANNER`.
- Selected direction: add one private, reversible, write-through diagnostic appender owned by the existing per-run `emitter`. For one sequential `ResolveBoundInto`, it normalizes each touched node's existing supported diagnostic property once on first successful touch, normalizes each incoming diagnostic once, merges normalized inputs, and writes the same `[]graphhealth.Diagnostic` representation back through `Graph.AddNode` after every append.
- Cause: current `AppendDiagnosticToNode` normalizes the incoming diagnostic, then its property path normalizes it again and re-normalizes every prior `[]Diagnostic` entry; structured normalization decodes the same JSON `Note` through envelope and outcome unmarshals. The observed CPU profiles prove this stack is material under `resolveCall`; the triangular `0 + 1 + ... + (k-1)` repeated-prior-entry shape is an inference, not an operation counter.
- No second optimization is authorized: do not redesign diagnostic policy, remove the two unmarshals inside one structured-outcome decode, change bucket lookup/sorting, open D002-D017, or fold a newly exposed residual into A002.
- Exact production owners only:
  - `internal/graphhealth/diagnostics.go`: add the run-scoped appender and one private normalized-input merge helper; cache only `map[nodeID][]Diagnostic` slice headers/state; first successful touch reads the current node and normalizes its existing supported representation once; every append re-reads the current node, preserves all other properties, writes `DiagnosticPropertyKey` as `[]Diagnostic`, and calls `Graph.AddNode`. Keep `AppendDiagnosticToNode` available for every other caller with its current normalization and observable semantics; helper refactoring is allowed only to prevent duplicated rules.
  - `internal/resolution/emit.go`: add the appender field to `emitter`, initialize it in `newEmitter` from the same graph, and route `(*emitter).emitUnresolvedReference` through it.
  - `internal/resolution/outcome.go`: route `(*emitter).emitTypeScriptOutcomeDiagnostic` through the same appender.
- Exact test owner only after production is correct: one new `internal/graphhealth/diagnostics_test.go`. No existing test file is authorized for edit by A002.
- Preserve-only/forbidden: `AppendDiagnosticToNode` behavior for other callers, `internal/resolution/resolve.go`, `internal/graph/types.go`, `internal/graphhealth/policy.go` and `Diagnostic`, graph types, diagnostic policy/contract, resolution outcome/metric schemas, persistence/readers, public APIs/contracts, timing instrumentation, and every other production/test/helper surface.

### A002 Preserved Contracts And Resource Boundary

- Preserve false-return behavior for nil graph, blank node ID, blank diagnostic kind, and absent source node; preserve `SourceNodeID`/`Count` defaults.
- Preserve structured fail-closed and legacy classification/actionability policy; diagnostic bucket identity; count merge; empty-target fill; earliest-source-line rule; stable order; and write-through visibility after every call.
- Preserve `[]Diagnostic` graph-property representation, Graph JSON, persistence/readback, health projections, public output, and deterministic replay.
- Preserve every resolution outcome, encoded `Note`, authority/proof/source-site field, metric, node, relationship, ID, label, property, branch, and execution order.
- Normalize every pre-existing supported diagnostic representation once on first touch exactly as the current function would. Cache validity is restricted to existing sequential `ResolveBoundInto` ownership; do not cache whole nodes or silently tolerate an unproven concurrent/out-of-band diagnostic writer.
- Let `T` be source nodes touched by resolution diagnostics. Retained appender state must be `O(T)` map entries and slice headers. Cache and graph must point to the same normalized slice, with no second retained copy of diagnostic objects. The first-touch replacement slice may allocate once; the previous property may then be collected.
- The appender map is created in `newEmitter`, becomes unreachable when `ResolveBoundInto` returns, and performs no serialization, I/O, goroutine, lock, global cache, explicit flush, or finalizer. Diagnostic slices remain owned only by the returned graph.

### A002 Validation, Two-Target Measurement, Rollback, And STOP

1. Before edit, future Coder repeats fresh `anvien --help`, `anvien analyze --force`, file-detail for all three production owners, and exact impacts for `AppendDiagnosticToNode`, `emitter`, `newEmitter`, `emitUnresolvedReference`, and `emitTypeScriptOutcomeDiagnostic`. Architect evidence does not replace this gate.
2. Edit only the three production owners first. Only after production behavior is correct may Coder add `internal/graphhealth/diagnostics_test.go`; compare repeated appender writes with legacy append semantics for first-touch decoded input, multiple unique entries, duplicate-bucket merging, stable order, policy fields, write-through after every append, and absent-node failure.
3. Perform exact build-holder/lock preflight, then require canonical full-build `PASS` before any test execution: `pwsh -NoLogo -NoProfile -NonInteractive -ExecutionPolicy Bypass -File E:\Anvien\scripts\full-build.ps1`.
4. Post-build focused validation includes the new appender-equivalence test, `TestP6DStructuredOutcomeProjectionPreservesGraphJSONAndHealthParity`, `TestResolveAttachesSourceBackedUnresolvedDiagnostics`, `TestP6C3FinalOutcomeStatusPrecedenceAndCarriageMatrix`, and `TestP6BResolverCarriesCatalogValidationFailuresPerSourceSite`; then run `go test ./internal/graphhealth -count=1` and `go test ./internal/resolution -count=1`. The known `TestProofBasedCallAccessGoldenCorpus` failure may be recorded only with exact unchanged-baseline proof, is never a PASS, and any new failure blocks A002.
5. Later, build one identifiable A002 candidate from accepted A001 plus only the authorized A002 diff, retaining identical 17-child instrumentation and runtime/native/build contract. Each target keeps its graph under its own `.anvien`, and target HEAD/status/workload plus process/executable/artifact hashes are recorded.
6. Measure targets independently. Cheapapp retains target `E:\cheapapp.org` and flags `--force --skip-git --json --progress`; Restaurant Manager retains target `E:\Restaurant_manager`, flags `--force --json --progress`, and exactly one `--exclude electron/renderer/src/api/userApi.ts`. Only executable identity, benchmark path, and label may differ from each target's accepted A001 command.
7. For each repo separately record D001 `resolve_calls`, parent `resolution`, analyzer total, process wall, calls/files, all 30 operation rows, all 17 child rows, scanned/parsed/failed files, graph nodes/relationships, DB readback counts, resolution semantic metrics, diagnostics/outcomes, canonical Graph JSON, public stdout, `startAllocBytes`, `endAllocBytes`, `maxObservedSys`, and proof that retained appender state is `O(T)` without duplicated diagnostics. Ladybug/meta raw-byte identity alone is not the semantic gate; build/test/microbenchmark/profile duration cannot substitute for wall-time boundaries.
8. No speed percentage or threshold is predicted. `KEEP` requires valid same-work D001 improvement plus corresponding net parent/process benefit on both repositories independently, all correctness/resource gates, and a later per-attempt Supervisor `PASS`.
9. Roll back only A002-owned appender/helper, `emitter` field/init, two route substitutions, and the new focused test file if any build/correctness/equivalence/determinism/lifecycle/memory/D001/parent/process gate fails. Do not reset/checkout/stash broadly or disturb A001/P1/protected work.
10. STOP and return to Main for fresh architecture if any additional production owner is required, including graph types, diagnostic policy/contract, `resolve.go`, outcome/metric schemas, persistence/readers, public contracts, global/concurrent caching, or timing instrumentation.

### Owner-Directed Reusable A00x Benchmark Build Script

- Latest Owner authority resolves the measurement-build blocker directly: stop further recovery discussion and create exactly one reusable `scripts/build-a00x-benchmark.ps1`. This is campaign measurement-support tooling first consumed by A002, not a new optimization attempt, production direction, benchmark result, or permission to edit the two canonical build scripts.
- The sole Coder may add only that script and update the existing A002 Coder handoff with its actual validation result. `scripts/full-build.ps1`, `anvien-launcher/build.ps1`, `internal/cli/package_command.go`, `internal/cli/package_runtime.go`, all A002 production/test files, both target repos, and every other path are preserve-only.
- The script must fail closed unless it is run from this repository, the caller supplies an explicit attempt ID, overlay manifest, expected overlay/mapped-source hashes, expected candidate-source hashes, and output directory beneath `E:\Anvien\.tmp`, and Go/native authority is present. For A002, those inputs are exactly the two `resolve.go`/`types.go` mappings from `E2-P2A-A002OVERLAY1` plus the current three A002 production hashes.
- The script must mirror the canonical runtime builder's offline/vendor/Ladybug CGO environment and build flags, adding only the verified Go overlay, `-buildvcs=false`, and an explicit `.tmp` output. It copies the pinned `lbug_shared.dll` beside the executable and writes a machine-readable provenance JSON containing attempt, executable, overlay, mapped-source, candidate-source, Go, native, and command identity.
- The script must not write canonical `anvien\bin`, launcher/server/web outputs, protocol registration, `.anvien`, target repositories, or global/shared caches; it builds only and never launches `analyze`.
- Execution order: write the script; run the canonical full build after exact holder preflight; only after full-build `PASS`, parse/validate the PowerShell file, run it once for A002, verify exit `0`, binary/version/DLL/provenance identities and staged set, then stop for Main verification. The later Cheapapp lane alone consumes the frozen binary for one corrected benchmark capture. Later A00x attempts call the same unchanged script with a new explicit contract and do not repeat interface discovery/audit unless the script, canonical build contract, Go/native authority, or required output semantics changed.
- When a future A00x has a different measured feature or build input and this script fails, use that exact failure to identify the missing attempt-specific contract, refresh the current attempt plan, and change only the relevant script parameter/guard/environment/tag/generated-input/provenance behavior. Do not reopen a general audit of the canonical scripts, and do not alter candidate production merely to make the benchmark binary compile.
- Rollback is limited to the new script and its exact `.tmp` output packet. A script failure changes no A001/A002 accepted baseline, no D001 streak or checkbox, and does not authorize rollback or edits to A002 production.
- Planner evidence: `E2-P2A-A002SCRIPTPLAN1`. Completed implementation/build evidence is recorded under `E2-P2A-A002SCRIPTBUILD1`; no numeric benchmark row, no no-KEEP result, and no accepted-state change is created by building the executable.

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

### P2: Complete Final D004 Attempt And Close Owner-Deferred Rows

- Phase Goal: retain safe cumulative end-to-end improvement through the one final A007 D004 chain, then close Child 06A optimization scope without falsely claiming improvement on deferred OP002-OP005 rows.
- Phase Boundary:
  - In scope: exactly one implementation slice, P2-A, repeated internally for dynamically selected attempts.
  - Out of scope: predeclared solution families, overlapping production owners, reused architecture decisions, or per-attempt commits.
  - Dependencies: accepted P1-A complete top-level table, matching parent checklist, and exact largest unchecked parent row.
- Phase Implementation Rule: P2-A stays open across all attempts. Each production attempt receives a new Architect, Planner refresh, Coder execution, and post-measurement Supervisor check. Only accepted Supervisor-`PASS` + Main-`KEEP` progress checkpoints may commit before P3-C; P3-C still owns its one local closure commit.
- Ordered Slice List:
  - P2-A: Execute Final D004 Attempt And Materialize Owner Scope Closure.

- [ ] P2-A: Execute Final D004 Attempt And Materialize Owner Scope Closure.
  - Goal: complete A007 against D004 with exact technical disposition, then check every remaining row through explicit retained terminal or Owner-deferred control so P3 can close Child 06A without opening OP002-OP005 optimization work.
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
       - Mini QA: reject an incomplete child list, largest-only list, unrecorded omission, historical hypothesis, secondary-metric ordering, or generic checklist item. Strict `PASS_BY_OWNER_COST_FLOOR` terminalization is allowed only with complete all-target packet proof.
       - Evidence target: `E2-P2A-AnnnDRILL1`, exact parent `B1-P1A-OPnnn`, and complete dynamic child rows `B2-P2A-Annn-Dnnn`.
       - Current A001 cursor: complete for `B1-P1A-OP001`; `E2-P2A-A001DRILL1` binds all `17` rows/items, and all parent/child checkboxes remain unchecked.
    2. Select the largest unchecked child, bind its exact goal/current child-parent-total basis/denominator/cause/owner/complete call path, then open a new Visible Architect and obtain the exact attempt-scoped decision.
       - UI flow check: N/A — architecture decision for CLI production work.
       - DB/data flow check: decision names affected data/read/lifecycle invariants and resource boundary.
       - Render location check: record `E2-P2A-AnnnCURRENT1` and `ARCH1` linked to the exact parent row, complete child list, and selected child row; do not create a separate architecture report gate.
       - Mini QA: Architect receives the full current parent-child list and selected child's measured owner, then names exact cause, direction, allowed owners, expected observable gain, validation/resources, and rollback; a coarse parent name is insufficient and the decision cannot be reused.
       - Evidence target: `E2-P2A-AnnnCURRENT1`, `E2-P2A-AnnnARCH1`.
       - Current transition: `E2-P2A-D004ATTRIB1/D004BOUNDARY1` consume that open D004 row and provide one target-separated causal owner/path packet to exactly one fresh visible Architect. They manufacture no attempt, streak event, speedup, production direction, or disposition.
    3. Planner refreshes the Current P2-A Attempt card and these execution steps to the Architect decision before Coder begins.
       - UI flow check: N/A — plan control update.
       - DB/data flow check: concrete validation covers every data/output/lifecycle invariant named by the decision.
       - Render location check: exact owner/files/symbols/change/tests/build/measure/equivalence/rollback replace the unbound attempt fields in this plan set.
       - Mini QA: `E2-P2A-A007PLAN1` translates `E2-P2A-A007ARCH1` without redesign: one-file guarded typed diagnostic read, exact JSON fallback, one new test file after production, fresh Coder Anvien/full-build/post-build regressions, two target-separated measurements, one independent Supervisor, exact rollback/STOP, and Owner final-scope transition. It changes no benchmark number, checkbox, streak, or queue before execution.
       - Evidence target: `E2-P2A-AnnnPLAN1`.
       - Current cursor: accepted A003 checkpoint remains unchanged; D001/D002 retain terminal history and D001 streak `2`; D003 and D005-D017 are `PASS_BY_OWNER_COST_FLOOR`; parent/D004 remain unchecked; A007 is Planner-ready and exactly one Coder release is pending.
    4. The sole Coder completes fresh required Anvien graph/file-detail/impact, implements only the refreshed production direction, edits authorized tests only after production is correct, performs the exact build-holder preflight, obtains canonical full-build `PASS`, and only then executes focused and `internal/resolution` package tests.
       - UI flow check: N/A unless activated by impact.
       - DB/data flow check: preserve exact affected data, ordering, failure, transaction/temp, publication, and reader semantics.
       - Render location check: cause/impact/source/build/test proof goes to evidence; product output remains observable boundary.
       - Mini QA: A007 Coder runs `anvien --help` -> one fresh `anvien analyze --force` -> file-detail `internal/resolution/outcome.go` -> upstream impact for every edited symbol. Production first adds only the guarded exact-typed read/fallback inside `resolutionOutcomeDiagnosticSites` (plus at most one private helper/import). Only after production is correct may it create `outcome_diagnostic_sites_test.go`; every existing test remains run-only. Any other owner, existing-test edit, fallback weakening, graph mutation, signature/shape/order/error/resource drift, or A005 lifecycle reuse forces STOP.
       - Evidence target in execution order: `E2-P2A-AnnnIMPACT1`, `E2-P2A-AnnnSRC1`, `E2-P2A-AnnnBUILD1`, `E2-P2A-AnnnTEST1`; `TEST1` records only the attempt-authorized post-production test bytes and test execution after full-build `PASS`.
       - Holder/build/test sequence: run `anvien doctor locks --repo E:\Anvien --json` and `anvien doctor processes --json`, terminate only proven build-output holders, require canonical full build exit `0`, then run the new D004 focused test plus `TestDiagnosticAppenderMatchesLegacySemantics`, `TestP6C3FinalOutcomeStatusPrecedenceAndCarriageMatrix`, `TestResolveAttachesSourceBackedUnresolvedDiagnostics`, `TestP6DStructuredOutcomeProjectionPreservesGraphJSONAndHealthParity`, `TestP6C3AnalyzeResultPreservesFinalOutcomesAndGraphCarriage`, `TestP6C3AnalyzeCapabilityOutcomesRetainAcceptedAuthorityStatus`, `TestP6C3ResolutionOutcomesPreserveGraphJSONAndLadybugParity`, and `TestP6C3NativeLadybugResolutionOutcomeReadback`; then run the five required packages. The unchanged known golden remains truthful and any new/changed failure blocks A007.
       - Current D004 cursor: `E2-P2A-A007ARCH1/A007PLAN1` authorize only the exact guarded typed-read/fallback implementation; parent and D004 remain unchecked and no candidate result exists.
    5. After A007 production/build/test validation, resume the same visible Coder task to record two sequential target-independent candidate measurements without combining them, using each target's accepted A003 artifact as that target's `before` and one identifiable A007 candidate containing only the authorized PLAN1 production bytes plus the accepted measurement overlay as `after`. Do not open a separate builder or measurement lane.
       - UI flow check: N/A unless activated by impact.
       - DB/data flow check: reuse unchanged `scripts/build-a00x-benchmark.ps1` and accepted 17-child/native/runtime/provenance contract; map only the A007 `outcome.go` candidate plus unchanged accepted measurement overlay inputs. No script/build-interface audit, target edit, accepted-before rerun, or rejected A004/A005 byte is allowed. Cheapapp retains `--force --skip-git --json --progress`; Restaurant retains `--force --json --progress` and exactly one `--exclude electron/renderer/src/api/userApi.ts`.
       - Render location check: benchmark records separate A007 pending/returned numeric blocks against each target's own accepted A003 baseline; evidence records exact executable/command/target/artifact/overlay identity and equivalence; actual status records which packet is pending/received. The same Coder task owns execution/reporting only and cannot issue acceptance or disposition.
       - Mini QA: each repo separately records D004, parent, analyzer, process wall, all `30/30` operations and ordered `17/17` children, exact D004 denominators, conservation/zero overlap, diagnostics/outcomes/SourceSiteIDs/ordered Evidence/reference indexes, canonical Graph JSON/stdout/stderr, logical Ladybug/native reader parity, `startAllocBytes`, `endAllocBytes`, `maxObservedSys`, and exact candidate/overlay/target/one-launch identity. Never rerun accepted A003 bases, average/combine targets, or use CPU/build/test duration as elapsed proof.
       - Evidence target: `E2-P2A-AnnnDIRECT1`, `E2-P2A-AnnnPARENT1`, `E2-P2A-AnnnE2E1`, `E2-P2A-AnnnEQUIV1`, `E2-P2A-AnnnREVIEW1`.
       - Current cursor: A007 candidate measurement stays locked until the sole Coder returns valid source/build/test evidence; no numeric row exists now.
    6. Record disposition, promote only an eligible child+parent+end-to-end `KEEP`, and keep the same child active until terminal.
       - UI flow check: N/A — ledger-controlled transition.
       - DB/data flow check: rejected bytes never become accepted state or ranking input.
       - Render location check: update child/parent/E2E benchmark rows, promotion/history/streak, evidence disposition, plan checklist, actual-status queues, and this attempt card immediately.
       - Mini QA: normal numeric eligibility still requires lower D004/parent/process wall on each target, no unexplained analyzer regression, and Supervisor `PASS`. Main promotes only an eligible `KEEP`; otherwise it exactly restores/retains accepted A003. Latest Owner scope then closes D004 after this one safe disposition, so no second D004 attempt or third-attempt manufacture occurs.
       - Evidence target: `E2-P2A-AnnnDECISION1`, conditional `E2-P2A-AnnnRESTORE1`, terminal `E2-P2A-AnnnSYSTEM1`, and `E2-P2A-AnnnRANK1`.
       - Current cursor: A001/A002/A003 remain accepted history; A004/A005 remain `SUPERVISOR_PASS / NO_KEEP / ROLLBACK_COMPLETE`; D001/D002 retain terminal history; D003 and D005-D017 are floor-PASS; A007 is the final D004 chain.
    7. After A007 is safely dispositioned, apply the Owner final-scope transition once: mark D004 and OP001 `CLOSED_BY_OWNER_SCOPE`, using the retained A007 packet after `KEEP` or the accepted A003 packet after restoration/retention; then mark OP002-OP005 `DEFERRED_BY_OWNER_SCOPE` without opening their drill-down or optimization lanes.
       - UI flow check: N/A — ledger-controlled hierarchy.
       - DB/data flow check: parent completion requires every measured child terminal on retained correct state; a new measured child reopens the parent.
       - Render location check: benchmark owns current hierarchy/order; plan marks exact rows; actual status points to remaining parent/child queues.
       - Mini QA: no row is omitted and every accepted number remains visible. `DEFERRED_BY_OWNER_SCOPE` is not `PASS`, `KEEP`, cost-floor, attempt, streak, speedup, or correctness acceptance. It closes only OP002-OP005 after D004 has a safe technical disposition; a D004 blocker keeps P2-A open.
       - Evidence target: `E2-P2A-AnnnRANK1`, plan checklist refresh, parent-completion evidence in the matching result record.
    8. After all `30/30` top-level and `17/17` OP001 rows are terminal under exact technical or Owner-scope evidence, record the final initial-versus-current total and complete final equivalence on the stable accepted implementation state, then open P3-A immediately.
       - UI flow check: N/A unless activated by impact.
       - DB/data flow check: verify canonical in-memory, Graph JSON, Ladybug/native, affected readers, determinism, freshness, failure, transaction/temp, and publication.
       - Render location check: final numbers in `B2-P2A-FINAL1`; proof in evidence; closure pointer in actual status.
       - Mini QA: final total is lower on same workload and every unexplained sibling/resource regression is resolved.
       - Evidence target: `E2-P2A-FINALBUILD1`, `E2-P2A-FINALTIME1`, `E2-P2A-FINALEQUIV1`, `E2-P2A-EXHAUST1`.
  - Implementation Gate:
    - Every production edit has a unique attempt ID, active parent, complete child list, selected largest unchecked child, current child/parent/total benchmark basis, exact cause/owner/call path, fresh attempt-specific Architect decision, and completed Planner refresh.
    - A003 identity/checkpoint and WAL checkpoint remain complete; A004/A005 remain `SUPERVISOR_PASS / NO_KEEP / ROLLBACK_COMPLETE`. `E2-P2A-A007ARCH1/A007PLAN1` bind D004's one final production chain; exactly one Coder transition is pending and no candidate result exists.
    - Accepted A002 progress checkpoint `E2-P2A-A002COMMIT1` is complete at `ecf825d709b761390a5df4a2147b6ed6eec04499` with its exact nine-path manifest and empty staged set. It satisfies the checkpoint dependency without closing P2-A, checking D001/parent, or changing streak `0`.
    - For A001 only, the latest direct Owner sequencing required Main verification of the four refreshed artifacts followed by independent consistency review in Architect thread `01a036f5-c602-7422-b484-ac8e58a34eb9`. That sequence passed and Main authorized the sole Coder transition under `E2-P2A-A001CONSISTENCY1`. This was Owner authority, not Architect command authority, and is not a generic gate for every attempt.
    - Every refreshed attempt states the child + parent + total elapsed-time reduction goal against the exact benchmark rows/current accepted baseline; supporting tooling cannot replace that goal.
    - Only one production/test writer is active. The measurement pool has `10` pre-opened visible lanes, at most `3` ACTIVE in one wave, exactly one measured parent/child problem per active lane, and no slot reuse until that lane's benchmark/evidence/checklist/status updates are recorded.
    - The previous attempt has a recorded Supervisor result, disposition, accepted-state promotion/restoration, unsuccessful-streak update, and current ledgers before another attempt opens.
  - Acceptance:
    - Source: only fresh-Architect-allowed measured owners changed; Coder followed the refreshed attempt exactly.
    - Runtime/UI: selected child, active parent, and total graph-generation elapsed wall time all improve for every `KEEP`; final total is lower; UI is N/A unless activated.
    - DB/data: each attempt's Supervisor and final evidence pass affected correctness, persistence, reader, failure, transaction/temp, and publication invariants.
    - Behavior test: A007 guarded canonical typed-read/fallback production implementation, one new D004 test file only after production is correct, holder preflight, canonical full build, post-build focused diagnostic/P6C3/P6D/Graph JSON/Ladybug regressions plus truthful five-package execution, later same-work target-separated remeasurement with verified A007 overlay, per-attempt review/disposition, Owner-scope transition, and final equivalence pass in that order.
    - Cleanup/quarantine: failed/superseded/debug work is identified for P3-B; protected work is untouched.
    - Checklist: every measured row is explicit; terminal rows are checked from retained technical evidence, strict Owner cost-floor evidence, or exact post-A007 Owner-scope control, and no open/blocked row is hidden.
    - Evidence IDs: dynamic attempt families including `E2-P2A-AnnnDRILL1/PARENT1`, plus `E2-P2A-FINALBUILD1`, `E2-P2A-FINALTIME1`, `E2-P2A-FINALEQUIV1`, `E2-P2A-EXHAUST1`.
    - Actual-status rows refreshed: active parent/child, remaining parent/child queues, attempt state, unsuccessful streak, Architect/Planner/Coder/Supervisor pointers, last disposition, accepted baseline, and next action.
  - Evidence Targets: complete top-level list, complete child lists, exact plan checklist mirror, parent/child selection, cause/owner/call path, Architect, Planner refresh, impact/source/tests/build, child/parent/E2E measurements, Supervisor, disposition, restoration/promotion, three-attempt terminal proof, parent completion, reranking, exhaustion, final proof.
  - Actual-status Update: refresh after every result; open P3-A only after final speedup/equivalence and terminal-disposition criteria pass.
  - Commit Boundary: `E2-P2A-A001COMMIT1` records A001 commit `17a1f3af37dcb61f9d389345822b6470a8f772cc`; `E2-P2A-A002COMMIT1` records accepted A002 checkpoint `ecf825d709b761390a5df4a2147b6ed6eec04499`; WAL fix checkpoint `0f3a572331dd23d17688886fcbfebeb7d37ee35d` is complete. P3-C retains its local one-closure-commit rule.

### P3: Final Acceptance And Campaign Closure

- Phase Goal: review the stable whole candidate, clean it without changing accepted behavior, detect it, create one Child 06A closure commit, and terminate the campaign without opening Child 07.
- Phase Boundary:
  - In scope: P3-A, P3-B, and P3-C only.
  - Out of scope: new optimization work, another implementation slice, another final Supervisor boundary, or another commit.
  - Dependencies: P2-A final speedup, per-attempt reviews, final equivalence, and exhaustion evidence.
- Phase Implementation Rule: P3-A is the one final whole-candidate Supervisor; P3-B performs no review; P3-C owns one local final detect/closure-commit/handoff sequence without invalidating earlier accepted progress checkpoints.

- [ ] P3-A: Run The Final Whole-Candidate Supervisor Review.
  - Goal: independently verify stable complete P2-A against current ledgers, source, build, every attempt review, measured initial/final result, and preserved invariants.
  - Work Steps:
    1. Review the exact stable implementation boundary and issue `PASS`, `REJECT`, or a concrete blocker.
    2. On reject, return only the failed measured owner/invariant through a new P2-A current drill-down -> Architect -> Planner -> Coder -> remeasure -> per-attempt Supervisor chain; resubmission stays inside this same final whole-candidate Supervisor boundary.
  - Implementation Gate: every measured row has retained technical terminal or exact Owner-scope control evidence, every matching plan checklist item is checked, and `B2-P2A-FINAL1` has lower comparable final total with final equivalence evidence.
  - Acceptance: `E3-P3A-REVIEW1` records the one effective final whole-candidate `PASS` with no unresolved same-scope finding.

- [ ] P3-B: Remove Exact Child 06A Dead Work Without A New Review.
  - Goal: remove only failed, superseded, temporary, or unused work created by Child 06A.
  - Work Steps:
    1. Inventory/remove exact Child 06A-owned dead/debug material; preserve historical/protected/shared/quarantined work.
    2. Prove accepted production/test bytes are unchanged from P3-A. If changed, invalidate P3-A and return to a new exact-owner P2-A attempt before closure.
  - Implementation Gate: `E3-P3A-REVIEW1` is `PASS`; cleanup cannot broaden scope or create documentation review.
  - Acceptance: `E3-P3B-CLEAN1` records removed/preserved/blocked items and accepted production/test identity remains unchanged.

- [ ] P3-C: Detect, Create One Campaign-Closure Commit, And Terminate.
  - Goal: close Child 06A and the full campaign at one isolated implementation commit; do not create or open Child 07.
  - Work Steps:
    1. Refresh final ledgers and run post-cleanup `anvien detect-changes --repo E:\Anvien --scope all` on the accepted implementation boundary.
    2. Stage only accepted Child 06A implementation/tests/necessary instrumentation/evidence/benchmark/plan ledgers, create exactly one implementation commit, and verify its manifest.
    3. Record terminal `CAMPAIGN_CLOSED / CHILD07_NOT_OPENED_BY_OWNER` status against the exact Child 06A closure commit and current benchmark/equivalence basis.
  - Implementation Gate: P3-A remains valid, P3-B changed no accepted production/test bytes, detect passes, and staging excludes unrelated/protected work.
  - Acceptance: `E3-P3C-DETECT1`, `E3-P3C-COMMIT1`, and historical-name `E3-P3C-HANDOFF1` record one detect, one campaign-closure commit, and `CAMPAIGN_CLOSED / CHILD07_NOT_OPENED_BY_OWNER` with no successor task.

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
