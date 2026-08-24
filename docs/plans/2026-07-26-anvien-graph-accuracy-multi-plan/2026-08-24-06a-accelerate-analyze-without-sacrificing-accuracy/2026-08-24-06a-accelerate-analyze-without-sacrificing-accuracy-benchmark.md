# Child 06A Accelerate Analyze Without Sacrificing Accuracy Benchmark Ledger

## Metadata

- Date: `2026-08-24`
- Plan: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy/2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy-plan.md`
- Plan rules: [plan-rules.md](plan-rules.md)
- Evidence: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy/2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy-evidence.md`
- Benchmark: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy/2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy-benchmark.md`
- Actual status: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy/2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy-actual-status.md`
- Method authority: `E:\Anvien\reports\Investigation\rp_main_260824_204615_child06a_optimization_method_discussion.md`
- Required provenance and handoff reference: `E:\Anvien\reports\Investigation\rp_huong_dan_lam_viec_de_slice_p6e_tien_hanh_nhanh_hon.md`

## Benchmark Rules

- This ledger is the central numeric source of truth and optimization control surface. No plan prose, actual-status copy, cost-map file, lane report, historical profile, or solution list can compete with its current accepted rows.
- Elapsed wall-clock time is the primary and controlling metric. Current absolute elapsed time alone orders the bottleneck list.
- P1-A writes the first real total graph-generation time and one parent row per measured top-level operation with its work denominator. It creates the complete top-level list in descending current absolute elapsed-time order; every smaller measured parent remains mandatory queued work.
- Exactly one visible measurement executor produces the initial accepted P1-A capture. Its exact real `anvien analyze` options/workload/cache/runtime identity is recorded in evidence at execution; after acceptance, up to three ACTIVE read-only measurement analysts may share that capture without running competing benchmark processes.
- P2-A selects the largest unchecked parent first. Processing that parent begins by measuring inside only its real boundary and creating the complete child list in descending absolute elapsed-time order, including every smaller measured child.
- Dynamic `B2-P2A-Annn-Dnnn` rows are the child control list for the active parent. The largest unchecked child is processed first, then every smaller child sequentially. No child may be omitted or terminalized because its elapsed time is small.
- Benchmark parent/child row cardinality must match the plan checklist mirror exactly. If drill-down finds `10` children, this ledger contains exactly `10` child rows for that parent and plan contains exactly `10` nested child checkboxes before Architect.
- Record every measurement immediately. Every numeric result references exact evidence that proves command/workload validity, operation boundary, denominator, and output comparability.
- `Initial baseline` is immutable once P1-A accepts the first comparable value. `Current accepted` is the value on the last correct retained state. `Current before` is the accepted value entering the latest attempt. `Candidate after` remains unpromoted until disposition.
- A candidate is promoted only when the selected-child value improves, the active-parent value improves, the end-to-end value retains the benefit, and that attempt's Visible Supervisor returns accuracy/equivalence `PASS`. On promotion, current accepted child/parent/total values become the candidate basis and cumulative deltas are recalculated from immutable initial baselines.
- CPU, RAM, allocation, GC, I/O, wait, call count, byte count, denominator, and share values are secondary only: they explain cause, guide the Architect, or prove comparable work. They cannot rank a row, prove speedup, or compensate for unchanged or worse total elapsed time.
- Supervisor `REJECT`, no retainable child/parent/end-to-end improvement, `REWORK`, or `ROLLBACK` without `KEEP` leaves accepted values unchanged and increments the selected child's consecutive unsuccessful-attempt count. Rejected candidate numbers remain only in attempt history and never drive ranking.
- Every further production edit is a new attempt with a new Architect decision, Planner refresh, Coder result, measurements, and Supervisor check. Benchmark does not reuse a prior attempt's decision or candidate value.
- `KEEP` resets only the selected child's consecutive unsuccessful-attempt count to `0`, promotes the accepted baseline, and keeps that child active. Remeasure child/parent/end-to-end and continue attempts on the same child until terminal.
- On the third consecutive attempt without `KEEP` at one accepted baseline, preserve accepted child/parent/total values, cite all three attempt rows/evidence, and set only the child to terminal `SYSTEM_CHARACTERISTIC`.
- `SYSTEM_CHARACTERISTIC` is not a speedup, accuracy waiver, acceptance of a rejected candidate, or parent completion. It checks off only that child; the largest remaining unchecked child of the same parent is next.
- A parent is processed only after every measured child is terminal and checked. Then remeasure parent/full pipeline, refresh the complete top-level list, check the parent, and select the largest remaining unchecked parent.
- `BLOCKED` is permitted only for concrete unavailable authority, dependency, or evidence. It stays unchecked and blocks parent/P2-A completion; it cannot skip a small row.
- Use existing timing/benchmark/profile capability first. Compare before/after only on a like-for-like workload, operation denominator, and instrumentation state. Do not rank inferred values, overlapping intervals, inclusive profile samples, or mismatched instrumentation states as if they were comparable elapsed times.
- If one required P1-A timing needs source instrumentation, its numeric row is ineligible until `E1-P1A-INSTR1` proves the sole visible sequential writer, exact owner and fresh graph/file-detail/impact, canonical full-build PASS before use, accepted runtime/instrumentation identity, overhead/comparability, and output equivalence. The branch must then record exactly one disposition: carry exact ownership into the first refreshed P2-A attempt, or remove the owned bytes, full-build again, and re-establish and remeasure the accepted timing basis. Benchmark never mixes pre-instrumented, instrumented, or post-removal values in one comparison.
- Supporting CPU, RAM, allocation, GC, I/O, and wait measurements are added only when they explain the selected elapsed-time cost or a required resource boundary. They never replace child, parent, and end-to-end wall-time proof.
- The available pool contains `10` pre-opened visible measurement lanes, with at most `3` ACTIVE in one wave. Each active lane owns exactly one already-measured parent/child bottleneck problem; unused lanes wait and no work is invented to fill the pool. Prefer one accepted capture shared by the active wave and never run duplicate competing benchmark processes merely to use concurrency. A slot is released only after its number, evidence proof, matching plan checklist state, and actual-status pointer are recorded; then the next waiting lane may be dispatched.
- If accepted measurement later exposes a missing top-level parent or child, append its real numeric row and require the matching unchecked plan checkbox immediately. Refresh the affected complete ordered list before selecting another target; no discovered row may live in only one ledger.
- Raw debug/profile material is temporary under repository-local `E:\Anvien\.tmp`. There is no durable per-run benchmark directory or separate cost-map/report tree.
- Build/test/review/detect/commit pass-fail belongs in evidence. This file records only measurements, rankings, attempt counts, dispositions, and evidence IDs.
- Final acceptance requires a lower comparable total graph-generation time on the same workload. No current measurement, baseline, ranking, attempt, or speedup is claimed by this documentation rewrite.

## Legacy Benchmark Provenance Mapping

Legacy numeric rows remain at their original Child 06 locations and are linked here without becoming current control rows.

| Legacy Child 06 benchmark | Current Child 06A benchmark | Mapping rule |
|---------------------------|-----------------------------|--------------|
| `B6-P6E-ATTRIBUTION1`, `B6-P6E-HISTORICAL1` | `B0-P0A-PROVENANCE1` | historical/profile attribution pointer only; non-comparable and excluded from ranking |
| `B6-P6E-AA1`, `B6-P6E-BASELINE1` | `B1-P1A-TOTAL` plus actual `B1-P1A-OPnnn` rows | no old numeric result carries forward; current real detailed measurement creates the baseline/ranking |
| `B6-P6E-U1AB1`, `B6-P6E-U2AB1`, `B6-P6E-U3AB1`, `B6-P6E-U4AB1` | dynamic `B2-P2A-Annn` attempt rows only if current ranking selects the corresponding real operation | legacy names are traceability, not a candidate list or execution order |
| `B6-P6E-REBASE1`, `B6-P6E-REBASE2`, `B6-P6E-PARETO1` | living control-surface refresh plus dynamic accepted-state `B2-P2A-Annn` history | every rerank is derived from current accepted absolute measurements |
| `B6-P6E-FINAL1` | `B2-P2A-FINAL1` | initial-versus-final total graph-generation result on the same workload |

## B0 - P0 Benchmarks

| Benchmark ID | Numeric authority | Current status | Ranking eligibility | Evidence |
|--------------|-------------------|----------------|---------------------|----------|
| `B0-P0A-PROVENANCE1` | exact historical values remain in legacy `B6-P6E-ATTRIBUTION1` and `B6-P6E-HISTORICAL1` | provenance only; not copied into current rows | excluded | `E0-P0A-PROVENANCE1` |

P0 creates no product-performance measurement. Child 06A implementation commit count `0` is structural current-state evidence, not a performance benchmark.

## B1 - P1 Initial Measurement And Living Optimization Control Surface

### Selection Rule

1. `B1-P1A-TOTAL` is the end-to-end product boundary and is not ranked against its internal operations.
2. P1-A appends `B1-P1A-OP001`, `B1-P1A-OP002`, and onward only from actual measured operations and orders them by `Current accepted` absolute elapsed time.
3. The largest unchecked parent is processed first; every smaller measured parent stays queued and must eventually be processed.
4. After child `KEEP` or `SYSTEM_CHARACTERISTIC`, accepted remeasurement updates parent/full-pipeline values without switching away from an active parent that still has unchecked children. After all children are checked, refresh this complete list and select the largest remaining unchecked parent.
5. Plan mirrors every parent row ID/name as one checkbox, while actual status references the active parent and remaining queue without copying numbers.

### Living Numeric Control Surface And Mandatory Ordered Bottleneck List

P1-A cannot close and P2-A cannot open until this table contains every currently measured top-level operation in descending accepted absolute elapsed-time order and plan contains one matching unchecked parent checkbox per row. The largest unchecked parent is first; all smaller parent rows remain mandatory.

| Benchmark row | Rank | Real parent operation / bottleneck | Current accepted elapsed time | Work denominator | Initial baseline | Current before | Meaningful share / delta | Processing state | Child list checked / measured | Plan parent checkbox | Current child attempt | Current after (candidate until promoted) | Parent direct delta | End-to-end before | End-to-end after | Cumulative parent delta from initial | Consecutive no-KEEP | Disposition | Proven owner / call path | Exact evidence |
|---------------|------|------------------------------------|-------------------------------|------------------|------------------|----------------|--------------------------|------------------|-------------------------------|----------------------|-----------------------|------------------------------------------|---------------------|-------------------|------------------|--------------------------------------|---------------------|-------------|--------------------------|----------------|
| `B1-P1A-TOTAL` | N/A — product boundary | total graph generation of real `anvien analyze` | `605.732722 s` process wall; `602.5278811 s` analyzer-internal timed total | `E:\Anvien`; `2,196` scanned / `765` parsed code / `0` failed; `123,075` nodes / `169,739` relationships; `20,363` dependency edges / `479` unresolved | `605.732722 s` accepted uninstrumented capture; final P1-A like-for-like baseline pending timing-gap disposition | `605.732722 s` | process CPU `803.093750 s`; CPU/wall `132.582%` | product boundary measured; P1-A ranking blocked by untimed residual | N/A | N/A | none | not measured | not measured | not measured | not measured | not measured | `0` | accepted capture; `P1_TIMING_GAP`; no speedup claim | process `E:\Anvien\anvien\bin\anvien.exe`; analyzer `internal/analyze.Run`; CLI `internal/cli` analyze command | `E1-P1A-IDENTITY1`, `E1-P1A-TOTAL1`, `E1-P1A-EQUIV1`; `E1-P1A-OPS1/INSTR1/RANK1` open |

### Accepted Capture Numeric Basis

| Capture field | Measured value | Evidence |
|---------------|----------------|----------|
| Capture label | `child06a-p1a-initial-20260824-225900` | `E1-P1A-IDENTITY1` |
| Process wall | `605.732722 s` | `E1-P1A-TOTAL1` |
| Analyzer-internal `totalDuration` | `602.5278811 s` | `E1-P1A-TOTAL1` |
| Process CPU | `803.093750 s` (`789.515625 s` user / `13.578125 s` kernel) | `E1-P1A-TOTAL1` |
| Phase count / phase sum | `15` / `587.2385976 s` | `E1-P1A-OPS1` partial |
| Go heap allocation at analyzer start / end | `1,431,736` / `996,685,256` bytes | `E1-P1A-OPS1` partial |
| Maximum observed Go `Sys` | `2,387,860,088` bytes | `E1-P1A-OPS1` partial |
| Process peak working set / allocation objects | `not captured` / `not instrumented` | `E1-P1A-OPS1` partial |

### Accepted Capture Phase Timings — Provisional, Not A Complete Parent Ranking

These are every phase timer emitted by the accepted capture, shown in descending elapsed-time order for numeric visibility only. They are not `B1-P1A-OPnnn` rows and do not populate the plan parent checklist because the untimed residual contains required real graph-finalization/persistence work whose elapsed-time rank is unknown.

| Capture metric | Observed order | Real timed boundary | Elapsed | Share of analyzer-internal total | Work denominator from execution | Queue eligibility | Evidence |
|----------------|----------------|---------------------|---------|----------------------------------|---------------------------------|-------------------|----------|
| `B1-P1A-PHASE-RESOLUTION` | 1 | `resolution.ResolveBoundInto` | `532.4778258 s` | `88.3740%` | definitions `46,593`; imports `5,562`; resolved/unresolved references `46,614 / 94,010`; calls/access/type `19,391 / 12,720 / 14,503`; emitted nodes/relationships `67,699 / 106,802` | measured; parent row withheld pending complete timing coverage | `E1-P1A-OPS1` partial |
| `B1-P1A-PHASE-DB-LOAD` | 2 | `loadGraph` -> CSV export + Ladybug load | `32.1148262 s` | `5.3300%` | node/relationship rows `123,075 / 169,739`; COPY calls `18 / 93`; fallback/fail/skip `0 / 0 / 0` | measured; parent row withheld pending complete timing coverage | `E1-P1A-OPS1` partial |
| `B1-P1A-PHASE-PARSE` | 3 | `parseFiles` | `14.8203538 s` | `2.4597%` | `765` files / `7,086,653` source bytes; `0` failed / `0` unsupported / `0` timed out | measured; parent row withheld pending complete timing coverage | `E1-P1A-OPS1` partial |
| `B1-P1A-PHASE-SEMANTIC` | 4 | `semantic.Apply` | `3.5868183 s` | `0.5953%` | `123,075` nodes visited / `169,739` relationships scanned; `98,730` gap inputs | measured; parent row withheld pending complete timing coverage | `E1-P1A-OPS1` partial |
| `B1-P1A-PHASE-CROSS-FILE` | 5 | `resolution.BuildCrossFileBinding` | `1.8734363 s` | `0.3109%` | definitions `46,593`; imports `5,562`; binding files/entries `139 / 904`; cross-file reprocess `0` | measured; parent row withheld pending complete timing coverage | `E1-P1A-OPS1` partial |
| `B1-P1A-PHASE-SCAN` | 6 | `scanner.WalkRepositoryPaths` | `0.9844880 s` | `0.1634%` | visited/included `2,411 / 2,196`; ignored/large/error `237 / 8 / 0` | measured; parent row withheld pending complete timing coverage | `E1-P1A-OPS1` partial |
| `B1-P1A-PHASE-ROUTES` | 7 | `routes.Apply` | `0.3750270 s` | `0.0622%` | `164` files scanned | measured; parent row withheld pending complete timing coverage | `E1-P1A-OPS1` partial |
| `B1-P1A-PHASE-PROCESSES` | 8 | `processes.Apply` | `0.3323185 s` | `0.0552%` | processes/steps `691 / 2,603`; CALLS edges considered `12,114` | measured; parent row withheld pending complete timing coverage | `E1-P1A-OPS1` partial |
| `B1-P1A-PHASE-DOCUMENTS` | 9 | `documents.Apply` | `0.3168023 s` | `0.0526%` | Markdown/spreadsheet `997 / 59`; sections/links `18,047 / 330`; metadata nodes `1,198` | measured; parent row withheld pending complete timing coverage | `E1-P1A-OPS1` partial |
| `B1-P1A-PHASE-TOOLS` | 10 | `tools.Apply` | `0.1799308 s` | `0.0299%` | `195` files scanned | measured; parent row withheld pending complete timing coverage | `E1-P1A-OPS1` partial |
| `B1-P1A-PHASE-COMMUNITIES` | 11 | `communities.Apply` | `0.1275115 s` | `0.0212%` | nodes/edges considered `6,487 / 10,972`; communities/memberships `1,529 / 6,487` | measured; parent row withheld pending complete timing coverage | `E1-P1A-OPS1` partial |
| `B1-P1A-PHASE-STRUCTURE` | 12 | `structure.Apply` | `0.0286679 s` | `0.0048%` | file/folder nodes `2,196 / 446`; CONTAINS emitted `7,714` | measured; parent row withheld pending complete timing coverage | `E1-P1A-OPS1` partial |
| `B1-P1A-PHASE-ORM` | 13 | `orm.Apply` | `0.0133807 s` | `0.0022%` | `111` files scanned | measured; parent row withheld pending complete timing coverage | `E1-P1A-OPS1` partial |
| `B1-P1A-PHASE-COBOL` | 14 | `cobol.Apply` | `0.0039133 s` | `0.0006%` | unavailable; capture emitted an empty COBOL metrics object | measured; parent row withheld pending complete timing coverage | `E1-P1A-OPS1` partial |
| `B1-P1A-PHASE-MRO` | 15 | `mro.Apply` | `0.0032972 s` | `0.0005%` | `15` classes analyzed | measured; parent row withheld pending complete timing coverage | `E1-P1A-OPS1` partial |

### Timing Coverage Gap

| Residual boundary | Exact measured residual | What source proves is inside/outside | Ranking consequence | Evidence |
|-------------------|-------------------------|-------------------------------------|---------------------|----------|
| Analyzer-internal total minus 15 phase timers | `15.2892835 s` (`2.5375%` of internal total) | unphased `Graph.Compact`, DB runner resolve/close, Graph JSON `writeGraphSnapshot`, and surrounding finalization are inside `totalDuration` but outside every `runPhase` timer | required real operations cannot receive separate elapsed values or rank | `E1-P1A-OPS1`, open `E1-P1A-INSTR1` |
| Process wall minus analyzer-internal total | `3.2048409 s` (`0.5291%` of process wall) | process startup; pre-timer lock/storage force preparation; benchmark write; registry/meta recording; generated AI context; file projection; JSON output are not separately timed by benchmark JSON | CLI/post-run publication cannot receive a real elapsed rank | `E1-P1A-OPS1`, open `E1-P1A-INSTR1` |
| Process wall minus 15 phase timers | `18.4941244 s` (`3.0534%` of process wall) | exact combined unattributed wall; it is an aggregate, not a fabricated operation | complete `B1-P1A-OPnnn` list and parent checklist remain absent | `E1-P1A-RANK1` pending |

Current top-level parent rows: none. The accepted capture and phase metrics are valid, but there is no complete rank-eligible parent queue until the unphased graph-finalization/persistence and CLI publication boundaries receive real non-overlapping elapsed timings under the conditional sole-writer branch. Therefore there is no active parent, remaining parent queue, child inventory, or production optimization target.

### Column Semantics

| Column | Numeric/control meaning |
|--------|-------------------------|
| Rank | current descending order by accepted absolute elapsed time; order chooses which unchecked parent goes first but never removes smaller rows |
| Current accepted elapsed time | latest correct retained absolute value; the only elapsed value eligible for ranking |
| Work denominator | actual work completed by this operation, used to prove same work before/after |
| Initial baseline | first accepted P1-A measurement; immutable for final cumulative comparison |
| Current before | accepted value immediately before the latest production attempt |
| Meaningful share / delta | current share of total or another meaningful comparison when mathematically valid; otherwise literal `not meaningful` |
| Processing state | unchecked / active / checked; `BLOCKED` remains unchecked and blocks completion |
| Child list checked / measured | exact progress cardinality for the complete child list; parent can be checked only when these counts are equal and accepted remeasurement exists |
| Plan parent checkbox | exact parent checklist item identity mirrored in plan; no elapsed values are copied there |
| Current child attempt | chronological attempt ID currently associated with a child of this parent, or none |
| Current after | measured value after latest attempt; candidate-only until all KEEP gates pass |
| Direct delta | candidate-after minus current-before, with percentage only when denominator permits |
| End-to-end before / after | total graph-generation values around the same attempt |
| Cumulative delta from initial | current-accepted minus initial-baseline after promotion |
| Consecutive unsuccessful attempts | child-owned count when one child is active; reset by child KEEP, terminal for that child at `3` |
| Disposition | parent processing state or child-attempt result; rejected values never become current accepted and small time is never a disposition |
| Proven owner / call path | exact current owner/call path only when evidence proves it; unknown is recorded without guessing |
| Evidence ID | exact measurement and validity/disposition proof |

## B2 - P2 Dynamic Optimization Attempts

### Active Parent Complete Child-Bottleneck List

After B1 selects the largest unchecked parent, measure only inside that parent and append every measured child here in descending current absolute elapsed-time order. The largest unchecked child is processed first, then every smaller child. Current rows: none.

| Parent row | Parent plan checkbox | Child row | Child rank | Real child operation / bottleneck | Current accepted child elapsed time | Work denominator | Initial child baseline | Current child before | Processing state | Plan child checkbox | Current attempt | Current child after | Direct child delta | Parent before | Parent after | Parent delta | End-to-end before | End-to-end after | End-to-end delta | Cumulative parent delta | Cumulative end-to-end delta | Consecutive child no-KEEP | Child disposition | Proven owner / complete call path | Exact evidence |
|------------|----------------------|-----------|------------|-----------------------------------|-------------------------------------|------------------|------------------------|----------------------|------------------|---------------------|-----------------|---------------------|--------------------|---------------|--------------|--------------|-------------------|------------------|------------------|-------------------------|-----------------------------|---------------------------|-------------------|-----------------------------------|----------------|

`E2-P2A-AnnnDRILL1` proves boundary validity, complete child coverage, denominators, ordering, and benchmark/plan checklist cardinality. If the parent has `10` measured children, this table has exactly `10` child rows and plan has exactly `10` nested child checkboxes before `E2-P2A-AnnnARCH1`. Missing/extra rows or checklist items block Architect. A concrete `BLOCKED` child stays unchecked.

### Attempt Numeric History

Append one row after every production attempt. The attempt ID is chronological and never reused. Current rows: none.

| Attempt | Parent row | Selected child row | Accepted baseline identity | Work denominator | Child before | Child after | Child delta | Parent before | Parent after | Parent delta | End-to-end before | End-to-end after | End-to-end delta | Cumulative parent delta | Cumulative end-to-end delta | Supervisor accuracy result | Disposition | Child unsuccessful streak after | Baseline / checklist effect | Evidence ID |
|---------|------------|--------------------|----------------------------|------------------|--------------|-------------|-------------|---------------|--------------|--------------|-------------------|------------------|------------------|-------------------------|-----------------------------|----------------------------|-------------|---------------------------------|-----------------------------|-------------|

Numeric promotion rules:

- `PASS` + lower child elapsed time + lower retained parent elapsed time + lower retained end-to-end graph-generation elapsed time -> `KEEP`; promote candidate, reset that child's streak to `0`, leave its checkbox unchecked, and continue the same child.
- Favorable secondary metrics with unchanged or worse end-to-end elapsed time -> no `KEEP`, no optimization claim, and no baseline promotion.
- Supervisor `REJECT` -> do not promote; restore/retain accepted baseline; increment the selected-child streak; next attempt on that child begins at a new Architect after Planner refresh.
- Supervisor `PASS` but no retainable child/parent/end-to-end gain -> do not promote; record `REWORK` or `ROLLBACK` as evidenced; increment the selected-child streak.
- Until streak reaches `3`, the same child remains active. `BLOCKED` is not terminal and leaves it unchecked.
- At streak `3`, write the terminal child record below, set only that child to `SYSTEM_CHARACTERISTIC`, remeasure accepted child/parent/end-to-end values, refresh both complete ordered lists, check its plan item, and select the largest remaining unchecked child of the same parent. Check the parent only after all child rows are checked and accepted parent/full-pipeline remeasurement exists.

### Three-Attempt Terminal Records

Current rows: none.

| Parent row | Child row | Current accepted baseline | Work denominator | Retained child elapsed time | Retained parent elapsed time | Attempt 1 / no-KEEP reason | Attempt 2 / no-KEEP reason | Attempt 3 / no-KEEP reason | Terminal child disposition | Plan child checkbox | Evidence ID |
|------------|-----------|---------------------------|------------------|-----------------------------|------------------------------|----------------------------|----------------------------|----------------------------|----------------------------|---------------------|-------------|

### Supporting Resource Measurements

Append only metrics needed to explain an active elapsed-time cost or validate its resource boundary. Current rows: none.

| Measurement ID | Attempt / control row | Metric | Unit | Current before | Candidate after | Delta | Accepted-state effect | Evidence ID |
|----------------|-----------------------|--------|------|----------------|-----------------|-------|-----------------------|-------------|

### Final Comparable Result

| Benchmark ID | Workload denominator | Initial total | Final accepted total | Absolute delta | Percent delta | Parent checklist checked / measured | Child checklist checked / measured | Unchecked blocked rows | Final disposition | Evidence ID |
|--------------|----------------------|---------------|----------------------|----------------|---------------|-------------------------------------|------------------------------------|------------------------|-------------------|-------------|
| `B2-P2A-FINAL1` | not recorded; P1-A has not run | not measured | not measured | not measured | not measured | `0/0` | `0/0` | none recorded | pending | `E2-P2A-FINALTIME1`, `E2-P2A-EXHAUST1` pending |

## B3 - P3 Benchmarks

P3-A/P3-B/P3-C are final review, cleanup, detect, commit, and handoff operations. They add no product-performance metric. Their pass/fail and one-commit boundary belong in evidence.

## Non-Benchmarkable Notes

- P6-D commit identity, Architect decisions, Planner refreshes, source/test/build results, Supervisor verdicts, cleanup, detect, commit, and handoff are evidence rather than measurements.
- Historical/protected values are provenance only and never enter `Current accepted`, ranking, attempt count, or final speedup.
- `SYSTEM_CHARACTERISTIC` records one child's terminal retained remaining cost after three complete unsuccessful attempts at one accepted baseline. It does not close the parent, omit smaller rows, or reduce the final requirement that total graph generation must be lower on the same workload.
