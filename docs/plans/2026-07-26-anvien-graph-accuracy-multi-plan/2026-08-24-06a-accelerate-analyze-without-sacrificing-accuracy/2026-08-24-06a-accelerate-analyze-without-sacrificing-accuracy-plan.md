# Child 06A Accelerate Analyze Without Sacrificing Accuracy Plan

## Metadata

- Date: `2026-08-24`
- Status: `P0-A complete / P1-A accepted capture recorded but timing gap open / no complete bottleneck ranking or parent checklist / P2-A implementation not started / no Child 06A implementation commit`
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
- sequential production optimization of current measured bottlenecks, with tests, full build, direct and end-to-end remeasurement, per-attempt Visible Supervisor accuracy/equivalence review, disposition, baseline promotion or restoration, and reranking;
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
- Coder implements only that refreshed direction, production first and tests second, after fresh required Anvien graph/file-detail/impact on the exact owner.
- Run the canonical full build, remeasure the selected child with the same denominator, remeasure its parent, and remeasure total graph-generation time on the same workload.
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
- Every retained production change has exact parent/child cause/owner/impact evidence, production-first implementation, tests, canonical full build, child before/after, parent before/after, end-to-end before/after, affected equivalence, Supervisor `PASS`, and an immediate `KEEP` disposition.
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

Current checklist entries: none. One accepted P1-A capture exists, but `E1-P1A-INSTR1` is open because existing timers leave required real operations unattributed. No complete `B1-P1A-OPnnn` ranking exists, so no parent or child checklist item can be truthfully created yet.

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
| Attempt state | `not opened; P1-A accepted capture exists, but timing coverage/ranking is incomplete and E1-P1A-INSTR1 is open` |
| Attempt ID | none |
| Attempt goal | when opened: reduce elapsed time of the exact selected benchmark row and total graph-generation elapsed time from the current accepted baseline while preserving all required invariants |
| Active parent benchmark row / checklist item | none; no top-level bottleneck is ranked or listed |
| Complete parent child list | none; no parent has been selected or drilled down |
| Active child benchmark row / checklist item | none |
| Remaining child queue | none; no child rows exist |
| Current accepted baseline | uninstrumented P1-A total is recorded at `E1-P1A-TOTAL1`; rankable like-for-like baseline remains pending instrumentation disposition |
| Consecutive unsuccessful attempts at current child baseline | `0`; no child or attempt exists |
| Selected child exact cause / owner / complete call path | not measured or evidenced |
| Fresh Architect decision | none; no production attempt exists |
| Planner refresh evidence | this initial method rewrite only; no production-attempt refresh exists |
| Allowed production/test surfaces | none |
| Expected observable gain | not defined before measurement and Architect decision |
| Preserved invariants / validation/resource boundary | inherited plan-wide invariants; attempt-specific boundary not yet defined |
| Rollback condition / exact owned bytes | no production bytes owned |
| Coder status | not authorized because no attempt has been selected and refreshed |
| Post-measurement Supervisor | none; no changed candidate exists |
| Next action | Main assigns one separate visible sequential sole instrumentation writer for the exact `E1-P1A-INSTR1` gap; P2-A remains closed |

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

- [ ] P1-A: Measure Detailed Graph-Generation Time And List Every Top-Level Bottleneck.
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
  - Current Execution Cursor (`2026-08-24`):
    - Step 1 complete: one visible executor produced accepted capture `child06a-p1a-initial-20260824-225900`; `E1-P1A-IDENTITY1`, `E1-P1A-TOTAL1`, and `E1-P1A-EQUIV1` are recorded.
    - Step 2 partial: every emitted phase timer and available denominator is recorded under provisional benchmark phase metrics, but the existing timer boundary leaves required graph-finalization/persistence and CLI publication work unattributed; `E1-P1A-OPS1` remains partial.
    - Step 3 open/pending Main assignment: `E1-P1A-INSTR1` records the exact timing gap, but no sole writer, exact owned bytes, fresh owner impact, build, instrumented capture, comparability proof, or terminal disposition exists.
    - Step 4 blocked: do not create `B1-P1A-OPnnn`, parent checklist entries, a largest parent pointer, or `E1-P1A-RANK1` until the conditional branch closes on a rankable like-for-like timing basis.
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
    2. Select the largest unchecked child, bind its exact goal/current child-parent-total basis/denominator/cause/owner/complete call path, then open a new Visible Architect and obtain the exact attempt-scoped decision.
       - UI flow check: N/A — architecture decision for CLI production work.
       - DB/data flow check: decision names affected data/read/lifecycle invariants and resource boundary.
       - Render location check: record `E2-P2A-AnnnCURRENT1` and `ARCH1` linked to the exact parent row, complete child list, and selected child row; do not create a separate architecture report gate.
       - Mini QA: Architect receives the full current parent-child list and selected child's measured owner, then names exact cause, direction, allowed owners, expected observable gain, validation/resources, and rollback; a coarse parent name is insufficient and the decision cannot be reused.
       - Evidence target: `E2-P2A-AnnnCURRENT1`, `E2-P2A-AnnnARCH1`.
    3. Planner refreshes the Current P2-A Attempt card and these execution steps to the Architect decision before Coder begins.
       - UI flow check: N/A — plan control update.
       - DB/data flow check: concrete validation covers every data/output/lifecycle invariant named by the decision.
       - Render location check: exact owner/files/symbols/change/tests/build/measure/equivalence/rollback replace the unbound attempt fields in this plan set.
       - Mini QA: Coder remains blocked until `E2-P2A-AnnnPLAN1` points to a complete, internally consistent refresh across plan/evidence/benchmark/status.
       - Evidence target: `E2-P2A-AnnnPLAN1`.
    4. Coder runs fresh required Anvien graph/file-detail/impact, implements only the refreshed production direction, then updates tests and runs the canonical full build.
       - UI flow check: N/A unless activated by impact.
       - DB/data flow check: preserve exact affected data, ordering, failure, transaction/temp, publication, and reader semantics.
       - Render location check: cause/impact/source/build/test proof goes to evidence; product output remains observable boundary.
       - Mini QA: production is correct before tests change; no owner beyond the refreshed allowed surface is edited.
       - Evidence target: `E2-P2A-AnnnIMPACT1`, `E2-P2A-AnnnSRC1`, `E2-P2A-AnnnTEST1`, `E2-P2A-AnnnBUILD1`.
    5. Remeasure the selected child, its parent, and total graph-generation time on the same work, then open a new Visible Supervisor accuracy/equivalence check for this attempt.
       - UI flow check: N/A unless activated by impact.
       - DB/data flow check: compare same child/parent/workload denominators and exact affected outputs/readers/lifecycle; review real evidence and candidate only.
       - Render location check: candidate numbers enter benchmark as unpromoted; Supervisor proof enters `E2-P2A-AnnnREVIEW1`.
       - Mini QA: no disposition or baseline promotion occurs before the Supervisor verdict; Supervisor does not edit code or audit report wording/hashes; improved secondary metrics cannot substitute for lower child, parent, and total elapsed time.
       - Evidence target: `E2-P2A-AnnnDIRECT1`, `E2-P2A-AnnnPARENT1`, `E2-P2A-AnnnE2E1`, `E2-P2A-AnnnEQUIV1`, `E2-P2A-AnnnREVIEW1`.
    6. Record disposition, promote only an eligible child+parent+end-to-end `KEEP`, and keep the same child active until terminal.
       - UI flow check: N/A — ledger-controlled transition.
       - DB/data flow check: rejected bytes never become accepted state or ranking input.
       - Render location check: update child/parent/E2E benchmark rows, promotion/history/streak, evidence disposition, plan checklist, actual-status queues, and this attempt card immediately.
       - Mini QA: `KEEP` resets the active child's streak but leaves its checkbox unchecked and starts another attempt on that child; reject restores accepted state and returns through fresh Architect/Planner, never Coder; only the third consecutive no-KEEP records `SYSTEM_CHARACTERISTIC`, refreshes accepted child/parent/E2E timings and both ordered lists, and checks that child.
       - Evidence target: `E2-P2A-AnnnDECISION1`, conditional `E2-P2A-AnnnRESTORE1`, terminal `E2-P2A-AnnnSYSTEM1`, and `E2-P2A-AnnnRANK1`.
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
    - Every refreshed attempt states the child + parent + total elapsed-time reduction goal against the exact benchmark rows/current accepted baseline; supporting tooling cannot replace that goal.
    - Only one production/test writer is active. The measurement pool has `10` pre-opened visible lanes, at most `3` ACTIVE in one wave, exactly one measured parent/child problem per active lane, and no slot reuse until that lane's benchmark/evidence/checklist/status updates are recorded.
    - The previous attempt has a recorded Supervisor result, disposition, accepted-state promotion/restoration, unsuccessful-streak update, and current ledgers before another attempt opens.
  - Acceptance:
    - Source: only fresh-Architect-allowed measured owners changed; Coder followed the refreshed attempt exactly.
    - Runtime/UI: selected child, active parent, and total graph-generation elapsed wall time all improve for every `KEEP`; final total is lower; UI is N/A unless activated.
    - DB/data: each attempt's Supervisor and final evidence pass affected correctness, persistence, reader, failure, transaction/temp, and publication invariants.
    - Behavior test: production-first tests, full build, same-work remeasurement, per-attempt review, and final equivalence pass.
    - Cleanup/quarantine: failed/superseded/debug work is identified for P3-B; protected work is untouched.
    - Checklist: every measured child and parent item is checked from retained terminal evidence; no `BLOCKED` or smaller measured row is hidden.
    - Evidence IDs: dynamic attempt families including `E2-P2A-AnnnDRILL1/PARENT1`, plus `E2-P2A-FINALBUILD1`, `E2-P2A-FINALTIME1`, `E2-P2A-FINALEQUIV1`, `E2-P2A-EXHAUST1`.
    - Actual-status rows refreshed: active parent/child, remaining parent/child queues, attempt state, unsuccessful streak, Architect/Planner/Coder/Supervisor pointers, last disposition, accepted baseline, and next action.
  - Evidence Targets: complete top-level list, complete child lists, exact plan checklist mirror, parent/child selection, cause/owner/call path, Architect, Planner refresh, impact/source/tests/build, child/parent/E2E measurements, Supervisor, disposition, restoration/promotion, three-attempt terminal proof, parent completion, reranking, exhaustion, final proof.
  - Actual-status Update: refresh after every result; open P3-A only after final speedup/equivalence and terminal-disposition criteria pass.
  - Commit Boundary: no attempt commit. The complete accepted Child 06A implementation is committed exactly once in P3-C.

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
