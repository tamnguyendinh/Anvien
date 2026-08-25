# Child 06A Accelerate Analyze Without Sacrificing Accuracy Benchmark Ledger

## Metadata

- Date: `2026-08-24`
- Current state: `P1-A complete / CARRY_TO_FIRST_P2A_REFRESH / B1-P1A-OP001 resolution has a complete 17-row child queue / A001 ARCHITECT_CONSISTENCY_PASS / CODER_ACTIVE_PRE_EDIT_GATE`; no new capture, production/test edit at the transition snapshot, candidate, KEEP, or speedup recorded; measurement/disposition/Supervisor/commit remain unauthorized
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
- For the selected child, `child_total_elapsed` is the sole value for child ranking, before/after comparison, and child-improvement acceptance. Active-child sub-metrics and work counters explain cause/comparability only and cannot substitute for it.
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
| `B1-P1A-TOTAL` | N/A — product boundary | total graph generation of real `anvien analyze` | `618.4358209 s` process wall; `615.8656873 s` analyzer-internal total | `E:\Anvien`; `2,200` scanned / `765` parsed code / `0` failed; `123,277` nodes / `169,976` relationships; `20,357` dependency edges / `479` unresolved | immutable uninstrumented reference `605.732722 s` process / `602.5278811 s` analyzer / `587.2385976 s` phase sum; not a like-for-like P2 baseline | `618.4358209 s` carried instrumented basis | arithmetic versus immutable reference `+12.7030989 s`; non-comparable due binary/HEAD/workload/output mismatch; process CPU `799.4062500 s` | product boundary accepted for carried instrumentation state; no speedup claim | N/A | N/A | none | not measured | not measured | `618.4358209 s` | not measured | not meaningful across mismatched states | `0` | `CARRY_TO_FIRST_P2A_REFRESH`; P2 parent drill-down pending | process `E:\Anvien\anvien\bin\anvien.exe`; analyzer `internal/analyze.Run`; CLI analyze command | `E1-P1A-TOTAL1`, `E1-P1A-INSTR1`, `E1-P1A-OPS1`, `E1-P1A-RANK1`, `E1-P1A-EQUIV1` |
| `B1-P1A-OP001` | 1 | `analyzer_internal` / `resolution` | carried parent `544.8094574000 s`; current child-drilldown parent `545.434182000 s` | `{"runs":1}` | `544.8094574000 s` | child-attribution basis `545.434182000 s` | cross-capture parent delta is not speed evidence; same-run child sum `545.364656400 s`, residual `0.069525600 s` | active / complete child queue; A001 Architect consistency PASS / Coder active pre-edit gate | `0/17` checked/measured | `B1-P1A-OP001` unchecked | `A001` / `B2-P2A-A001-D001`; sole Coder active pre-edit; no production/test edit at transition snapshot | not measured | not measured | same-run drilldown process `653.475797100 s`; analyzer `630.160832200 s` | not measured | not comparable across moving captures | `0` | active parent / child selected; no attempt disposition | `resolution.ResolveBoundInto`; selected child `resolveCall`; complete child call paths in B2 table | `E1-P1A-OPS1`, `E1-P1A-RANK1`, `E2-P2A-A001DRILL1`, `E2-P2A-A001CURRENT1`, `E2-P2A-A001ARCH1`, `E2-P2A-A001PLAN1`, `E2-P2A-A001CONSISTENCY1` |
| `B1-P1A-OP002` | 2 | `analyzer_internal` / `db_load` | `33.9223067000 s` | `{"runs":1}` | `33.9223067000 s` | `33.9223067000 s` | `5.485178%` of instrumented process wall; baseline delta `0` | unchecked / queued | not measured | `B1-P1A-OP002` unchecked | none | not measured | not measured | `618.4358209 s` | not measured | `0` | `0` | pending | not yet proven | `E1-P1A-OPS1`, `E1-P1A-RANK1` |
| `B1-P1A-OP003` | 3 | `analyzer_internal` / `parse` | `14.9652984000 s` | `{"runs":1}` | `14.9652984000 s` | `14.9652984000 s` | `2.419863%` of instrumented process wall; baseline delta `0` | unchecked / queued | not measured | `B1-P1A-OP003` unchecked | none | not measured | not measured | `618.4358209 s` | not measured | `0` | `0` | pending | not yet proven | `E1-P1A-OPS1`, `E1-P1A-RANK1` |
| `B1-P1A-OP004` | 4 | `analyzer_internal` / `graph_snapshot` | `11.3731742000 s` | `{"nodes":123277,"relationships":169976}` | `11.3731742000 s` | `11.3731742000 s` | `1.839023%` of instrumented process wall; baseline delta `0` | unchecked / queued | not measured | `B1-P1A-OP004` unchecked | none | not measured | not measured | `618.4358209 s` | not measured | `0` | `0` | pending | not yet proven | `E1-P1A-OPS1`, `E1-P1A-RANK1` |
| `B1-P1A-OP005` | 5 | `analyzer_internal` / `semantic_enrichment` | `3.5949223000 s` | `{"runs":1}` | `3.5949223000 s` | `3.5949223000 s` | `0.581293%` of instrumented process wall; baseline delta `0` | unchecked / queued | not measured | `B1-P1A-OP005` unchecked | none | not measured | not measured | `618.4358209 s` | not measured | `0` | `0` | pending | not yet proven | `E1-P1A-OPS1`, `E1-P1A-RANK1` |
| `B1-P1A-OP006` | 6 | `analyzer_internal` / `db_runner_resolve` | `2.3800947000 s` | `{"runners":1}` | `2.3800947000 s` | `2.3800947000 s` | `0.384857%` of instrumented process wall; baseline delta `0` | unchecked / queued | not measured | `B1-P1A-OP006` unchecked | none | not measured | not measured | `618.4358209 s` | not measured | `0` | `0` | pending | not yet proven | `E1-P1A-OPS1`, `E1-P1A-RANK1` |
| `B1-P1A-OP007` | 7 | `analyzer_internal` / `cross_file_binding` | `1.7802921000 s` | `{"runs":1}` | `1.7802921000 s` | `1.7802921000 s` | `0.287870%` of instrumented process wall; baseline delta `0` | unchecked / queued | not measured | `B1-P1A-OP007` unchecked | none | not measured | not measured | `618.4358209 s` | not measured | `0` | `0` | pending | not yet proven | `E1-P1A-OPS1`, `E1-P1A-RANK1` |
| `B1-P1A-OP008` | 8 | `cli_outer` / `ai_context` | `1.6165628000 s` | `{"baseSkills":49,"generatedFiles":4}` | `1.6165628000 s` | `1.6165628000 s` | `0.261395%` of instrumented process wall; baseline delta `0` | unchecked / queued | not measured | `B1-P1A-OP008` unchecked | none | not measured | not measured | `618.4358209 s` | not measured | `0` | `0` | pending | not yet proven | `E1-P1A-OPS1`, `E1-P1A-RANK1` |
| `B1-P1A-OP009` | 9 | `analyzer_internal` / `analyzer_orchestration` | `0.9917413000 s` | `{"analyzeRuns":1}` | `0.9917413000 s` | `0.9917413000 s` | `0.160363%` of instrumented process wall; baseline delta `0` | unchecked / queued | not measured | `B1-P1A-OP009` unchecked | none | not measured | not measured | `618.4358209 s` | not measured | `0` | `0` | pending | not yet proven | `E1-P1A-OPS1`, `E1-P1A-RANK1` |
| `B1-P1A-OP010` | 10 | `analyzer_internal` / `processes` | `0.4515488000 s` | `{"runs":1}` | `0.4515488000 s` | `0.4515488000 s` | `0.073015%` of instrumented process wall; baseline delta `0` | unchecked / queued | not measured | `B1-P1A-OP010` unchecked | none | not measured | not measured | `618.4358209 s` | not measured | `0` | `0` | pending | not yet proven | `E1-P1A-OPS1`, `E1-P1A-RANK1` |
| `B1-P1A-OP011` | 11 | `analyzer_internal` / `scan` | `0.4365820000 s` | `{"runs":1}` | `0.4365820000 s` | `0.4365820000 s` | `0.070595%` of instrumented process wall; baseline delta `0` | unchecked / queued | not measured | `B1-P1A-OP011` unchecked | none | not measured | not measured | `618.4358209 s` | not measured | `0` | `0` | pending | not yet proven | `E1-P1A-OPS1`, `E1-P1A-RANK1` |
| `B1-P1A-OP012` | 12 | `cli_outer` / `file_projection` | `0.4041720000 s` | `{"files":2200,"nodes":123277,"relationships":169976}` | `0.4041720000 s` | `0.4041720000 s` | `0.065354%` of instrumented process wall; baseline delta `0` | unchecked / queued | not measured | `B1-P1A-OP012` unchecked | none | not measured | not measured | `618.4358209 s` | not measured | `0` | `0` | pending | not yet proven | `E1-P1A-OPS1`, `E1-P1A-RANK1` |
| `B1-P1A-OP013` | 13 | `analyzer_internal` / `routes` | `0.3293505000 s` | `{"runs":1}` | `0.3293505000 s` | `0.3293505000 s` | `0.053255%` of instrumented process wall; baseline delta `0` | unchecked / queued | not measured | `B1-P1A-OP013` unchecked | none | not measured | not measured | `618.4358209 s` | not measured | `0` | `0` | pending | not yet proven | `E1-P1A-OPS1`, `E1-P1A-RANK1` |
| `B1-P1A-OP014` | 14 | `analyzer_internal` / `documents` | `0.2766779000 s` | `{"runs":1}` | `0.2766779000 s` | `0.2766779000 s` | `0.044738%` of instrumented process wall; baseline delta `0` | unchecked / queued | not measured | `B1-P1A-OP014` unchecked | none | not measured | not measured | `618.4358209 s` | not measured | `0` | `0` | pending | not yet proven | `E1-P1A-OPS1`, `E1-P1A-RANK1` |
| `B1-P1A-OP015` | 15 | `analyzer_internal` / `db_runner_close` | `0.2167895000 s` | `{"runners":1}` | `0.2167895000 s` | `0.2167895000 s` | `0.035054%` of instrumented process wall; baseline delta `0` | unchecked / queued | not measured | `B1-P1A-OP015` unchecked | none | not measured | not measured | `618.4358209 s` | not measured | `0` | `0` | pending | not yet proven | `E1-P1A-OPS1`, `E1-P1A-RANK1` |
| `B1-P1A-OP016` | 16 | `analyzer_outer` / `analyze_setup` | `0.1697595000 s` | `{"analyzeRuns":1}` | `0.1697595000 s` | `0.1697595000 s` | `0.027450%` of instrumented process wall; baseline delta `0` | unchecked / queued | not measured | `B1-P1A-OP016` unchecked | none | not measured | not measured | `618.4358209 s` | not measured | `0` | `0` | pending | not yet proven | `E1-P1A-OPS1`, `E1-P1A-RANK1` |
| `B1-P1A-OP017` | 17 | `analyzer_internal` / `tools` | `0.1611048000 s` | `{"runs":1}` | `0.1611048000 s` | `0.1611048000 s` | `0.026050%` of instrumented process wall; baseline delta `0` | unchecked / queued | not measured | `B1-P1A-OP017` unchecked | none | not measured | not measured | `618.4358209 s` | not measured | `0` | `0` | pending | not yet proven | `E1-P1A-OPS1`, `E1-P1A-RANK1` |
| `B1-P1A-OP018` | 18 | `analyzer_internal` / `communities` | `0.1281742000 s` | `{"runs":1}` | `0.1281742000 s` | `0.1281742000 s` | `0.020726%` of instrumented process wall; baseline delta `0` | unchecked / queued | not measured | `B1-P1A-OP018` unchecked | none | not measured | not measured | `618.4358209 s` | not measured | `0` | `0` | pending | not yet proven | `E1-P1A-OPS1`, `E1-P1A-RANK1` |
| `B1-P1A-OP019` | 19 | `cli_outer` / `registry_meta` | `0.0906005000 s` | `{"repositories":1}` | `0.0906005000 s` | `0.0906005000 s` | `0.014650%` of instrumented process wall; baseline delta `0` | unchecked / queued | not measured | `B1-P1A-OP019` unchecked | none | not measured | not measured | `618.4358209 s` | not measured | `0` | `0` | pending | not yet proven | `E1-P1A-OPS1`, `E1-P1A-RANK1` |
| `B1-P1A-OP020` | 20 | `cli_outer` / `cli_preparation` | `0.0771271000 s` | `{"commands":1}` | `0.0771271000 s` | `0.0771271000 s` | `0.012471%` of instrumented process wall; baseline delta `0` | unchecked / queued | not measured | `B1-P1A-OP020` unchecked | none | not measured | not measured | `618.4358209 s` | not measured | `0` | `0` | pending | not yet proven | `E1-P1A-OPS1`, `E1-P1A-RANK1` |
| `B1-P1A-OP021` | 21 | `cli_outer` / `output_publication` | `0.0245234000 s` | `{"outputs":1}` | `0.0245234000 s` | `0.0245234000 s` | `0.003965%` of instrumented process wall; baseline delta `0` | unchecked / queued | not measured | `B1-P1A-OP021` unchecked | none | not measured | not measured | `618.4358209 s` | not measured | `0` | `0` | pending | not yet proven | `E1-P1A-OPS1`, `E1-P1A-RANK1` |
| `B1-P1A-OP022` | 22 | `analyzer_internal` / `orm` | `0.0198245000 s` | `{"runs":1}` | `0.0198245000 s` | `0.0198245000 s` | `0.003206%` of instrumented process wall; baseline delta `0` | unchecked / queued | not measured | `B1-P1A-OP022` unchecked | none | not measured | not measured | `618.4358209 s` | not measured | `0` | `0` | pending | not yet proven | `E1-P1A-OPS1`, `E1-P1A-RANK1` |
| `B1-P1A-OP023` | 23 | `cli_outer` / `cli_startup` | `0.0144171000 s` | `{"commands":1}` | `0.0144171000 s` | `0.0144171000 s` | `0.002331%` of instrumented process wall; baseline delta `0` | unchecked / queued | not measured | `B1-P1A-OP023` unchecked | none | not measured | not measured | `618.4358209 s` | not measured | `0` | `0` | pending | not yet proven | `E1-P1A-OPS1`, `E1-P1A-RANK1` |
| `B1-P1A-OP024` | 24 | `analyzer_internal` / `structure` | `0.0124608000 s` | `{"runs":1}` | `0.0124608000 s` | `0.0124608000 s` | `0.002015%` of instrumented process wall; baseline delta `0` | unchecked / queued | not measured | `B1-P1A-OP024` unchecked | none | not measured | not measured | `618.4358209 s` | not measured | `0` | `0` | pending | not yet proven | `E1-P1A-OPS1`, `E1-P1A-RANK1` |
| `B1-P1A-OP025` | 25 | `analyzer_internal` / `graph_compact` | `0.0111125000 s` | `{"inputNodes":123277,"inputRelationships":169976}` | `0.0111125000 s` | `0.0111125000 s` | `0.001797%` of instrumented process wall; baseline delta `0` | unchecked / queued | not measured | `B1-P1A-OP025` unchecked | none | not measured | not measured | `618.4358209 s` | not measured | `0` | `0` | pending | not yet proven | `E1-P1A-OPS1`, `E1-P1A-RANK1` |
| `B1-P1A-OP026` | 26 | `analyzer_internal` / `mro` | `0.0033453000 s` | `{"runs":1}` | `0.0033453000 s` | `0.0033453000 s` | `0.000541%` of instrumented process wall; baseline delta `0` | unchecked / queued | not measured | `B1-P1A-OP026` unchecked | none | not measured | not measured | `618.4358209 s` | not measured | `0` | `0` | pending | not yet proven | `E1-P1A-OPS1`, `E1-P1A-RANK1` |
| `B1-P1A-OP027` | 27 | `analyzer_internal` / `benchmark_write` | `0.0010538000 s` | `{"artifacts":1}` | `0.0010538000 s` | `0.0010538000 s` | `0.000170%` of instrumented process wall; baseline delta `0` | unchecked / queued | not measured | `B1-P1A-OP027` unchecked | none | not measured | not measured | `618.4358209 s` | not measured | `0` | `0` | pending | not yet proven | `E1-P1A-OPS1`, `E1-P1A-RANK1` |
| `B1-P1A-OP028` | 28 | `analyzer_internal` / `cobol` | `0.0003756000 s` | `{"runs":1}` | `0.0003756000 s` | `0.0003756000 s` | `0.000061%` of instrumented process wall; baseline delta `0` | unchecked / queued | not measured | `B1-P1A-OP028` unchecked | none | not measured | not measured | `618.4358209 s` | not measured | `0` | `0` | pending | not yet proven | `E1-P1A-OPS1`, `E1-P1A-RANK1` |
| `B1-P1A-OP029` | 29 | `cli_outer` / `memory_profile` | `0.0000000000 s` | `{"profiles":0}` | `0.0000000000 s` | `0.0000000000 s` | `0.000000%` of instrumented process wall; baseline delta `0` | unchecked / queued | not measured | `B1-P1A-OP029` unchecked | none | not measured | not measured | `618.4358209 s` | not measured | `0` | `0` | pending | not yet proven | `E1-P1A-OPS1`, `E1-P1A-RANK1` |
| `B1-P1A-OP030` | 30 | `cli_outer` / `cpu_profile_completion` | `0.0000000000 s` | `{"profiles":0}` | `0.0000000000 s` | `0.0000000000 s` | `0.000000%` of instrumented process wall; baseline delta `0` | unchecked / queued | not measured | `B1-P1A-OP030` unchecked | none | not measured | not measured | `618.4358209 s` | not measured | `0` | `0` | pending | not yet proven | `E1-P1A-OPS1`, `E1-P1A-RANK1` |

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

### Carried Instrumented Capture Numeric Basis

| Capture field | Measured value | Evidence |
|---------------|----------------|----------|
| Capture label | `child06a-p1a-instrumented-20260825-040400` | `E1-P1A-INSTR1` |
| Process wall / CPU | `618.4358209 / 799.4062500 s` | `E1-P1A-TOTAL1`, `E1-P1A-INSTR1` |
| Analyzer internal / phase sum | `615.8656873 / 600.8917213 s` | `E1-P1A-OPS1` |
| Analyzer non-phase / analyzer residual | `14.9739660 / 0 s` | `E1-P1A-OPS1` |
| Analyzer outer / CLI outer / all operations | `0.1697595 / 2.2274029 / 618.2628497 s` | `E1-P1A-OPS1` |
| Uncovered process residual | `0.1729712 s` | `E1-P1A-OPS1` |
| Operation count | `30`: `21 analyzer_internal`, `1 analyzer_outer`, `8 cli_outer` | `E1-P1A-OPS1`, `E1-P1A-RANK1` |
| Workload/output | `2,200` scanned / `765` parsed / `0` failed; `123,277` nodes / `169,976` relationships; projection `20,357 / 479` | `E1-P1A-EQUIV1` |
| Instrumentation disposition | `CARRY_TO_FIRST_P2A_REFRESH` | `E1-P1A-INSTR1` |

The immutable uninstrumented reference remains `605.732722 / 602.5278811 / 587.2385976 s`. Arithmetic deltas are recorded, but are not performance evidence because binary, HEAD, workload, and output counts differ. P2 uses the carried instrumented capture as its current same-instrumentation basis.

### Accepted Uninstrumented Capture Phase Timings — Historical Visibility

These are every phase timer emitted by the immutable uninstrumented capture, shown for provenance only. They are not the current queue rows; the complete current `B1-P1A-OP001..OP030` ranking comes from the carried instrumented capture above.

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

### Historical Timing Coverage Gap — Closed By Carried Instrumentation

| Residual boundary | Exact measured residual | What source proves is inside/outside | Ranking consequence | Evidence |
|-------------------|-------------------------|-------------------------------------|---------------------|----------|
| Analyzer-internal total minus 15 phase timers | `15.2892835 s` (`2.5375%` of internal total) | unphased `Graph.Compact`, DB runner resolve/close, Graph JSON `writeGraphSnapshot`, and surrounding finalization are inside `totalDuration` but outside every `runPhase` timer | required real operations cannot receive separate elapsed values or rank | `E1-P1A-OPS1`, open `E1-P1A-INSTR1` |
| Process wall minus analyzer-internal total | `3.2048409 s` (`0.5291%` of process wall) | process startup; pre-timer lock/storage force preparation; benchmark write; registry/meta recording; generated AI context; file projection; JSON output are not separately timed by benchmark JSON | CLI/post-run publication cannot receive a real elapsed rank | `E1-P1A-OPS1`, open `E1-P1A-INSTR1` |
| Process wall minus 15 phase timers | `18.4941244 s` (`3.0534%` of process wall) | exact combined unattributed wall; it is an aggregate, not a fabricated operation | complete `B1-P1A-OPnnn` list and parent checklist remain absent | `E1-P1A-RANK1` pending |

The historical gap is closed for current queue construction. The carried instrumented capture supplies every emitted operation value, exact denominators, a zero analyzer-internal residual, and the complete `OP001..OP030` list. `OP001 resolution` is active; `OP002..OP030` remain queued. No child inventory or production attempt exists yet.

### Captured Instrumentation Contract And Terminal Disposition

This table records the operation identity/denominator coverage that produced the current queue. Elapsed values and ranks are in `B1-P1A-OP001..OP030`; this section does not duplicate them.

| Boundary | Captured operation names | Denominators emitted | Numeric state |
|----------|-----------------------------------------------|----------------------|-------------------------|
| `analyzer_outer` | `analyze_setup` | `analyzeRuns` | captured / ranked |
| `analyzer_internal` | every existing phase name | `runs` | captured / ranked; all durations exactly match phases |
| `analyzer_internal` | `graph_compact` | `inputNodes`, `inputRelationships` | captured / ranked |
| `analyzer_internal` | `db_runner_resolve`; `db_runner_close` | `runners` | captured / ranked |
| `analyzer_internal` | `graph_snapshot` | `nodes`, `relationships` | captured / ranked |
| `analyzer_internal` | `analyzer_orchestration` | `analyzeRuns` | captured / ranked; exclusive residual outside phases/explicit analyzer operations |
| `analyzer_internal` | `benchmark_write` | `artifacts` | captured / ranked |
| `cli_outer` | `cli_startup`, `cli_preparation` | `commands` | captured / ranked |
| `cli_outer` | `memory_profile`, `cpu_profile_completion` | `profiles` | captured / ranked at zero with denominator zero |
| `cli_outer` | `registry_meta` | `repositories` | captured / ranked |
| `cli_outer` | `ai_context` | `generatedFiles`, `baseSkills` | captured / ranked |
| `cli_outer` | `file_projection` | `files`, `nodes`, `relationships` | captured / ranked |
| `cli_outer` | `output_publication` | `outputs` | captured / ranked |

Instrumentation identity: sole writer `01a0348f-3bb7-7c70-8ec5-61176ce00591`; scoped diff `internal/analyze/analyze.go +117/-15`, `internal/cli/command.go +157/-38`, `internal/cli/command_test.go +75/-0`; runtime `1.2.8`, SHA-256 `62920CBF15921EF8A6D2FAC671776BAA3C312EFEA1BEED53B721C4AFF5E1B6C5`. Canonical full-build validation passed with exit `0`. Its `700.1712706 s` validation elapsed includes build work plus maintenance `analyze . --force`; it is not build performance and remains excluded from every benchmark row. Terminal disposition is `CARRY_TO_FIRST_P2A_REFRESH`.

The completed numeric authority used exactly:

```text
E:\Anvien\anvien\bin\anvien.exe analyze E:\Anvien --force --json --progress --benchmark-json E:\Anvien\.tmp\child06a_p1a_instrumented_20260825_040400\benchmark.json --benchmark-label child06a-p1a-instrumented-20260825-040400
```

Executor `01a03597-6afb-73e1-bee6-4306836cce01` completed one launch and is idle. `E1-P1A-RANK1` is complete, P1-A is checked, and P2-A is open only at active-parent drill-down; no child or production attempt exists.

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

After B1 selects the largest unchecked parent, measure only inside that parent and append every measured child here in descending current absolute elapsed-time order. The largest unchecked child is processed first, then every smaller child. Current rows: `17`.

Same-run numeric basis: process wall `653.475797100 s`; process CPU `830.109375000 s`; analyzer `630.160832200 s`; parent `545.434182000 s` with `{"runs":1}`; child sum `545.364656400 s`; residual `0.069525600 s`; `2309` exclusive intervals; `0` overlaps. Graph-count differences against the earlier moving-repo capture are cross-capture relative counts, not a speed or semantic-regression verdict.

| Parent row | Parent plan checkbox | Child row | Child rank | Real child operation / bottleneck | Current accepted child elapsed time | Work denominator | Initial child baseline | Current child before | Processing state | Plan child checkbox | Current attempt | Current child after | Direct child delta | Parent before | Parent after | Parent delta | End-to-end before | End-to-end after | End-to-end delta | Cumulative parent delta | Cumulative end-to-end delta | Consecutive child no-KEEP | Child disposition | Proven owner / complete call path | Exact evidence |
|------------|----------------------|-----------|------------|-----------------------------------|-------------------------------------|------------------|------------------------|----------------------|------------------|---------------------|-----------------|---------------------|--------------------|---------------|--------------|--------------|-------------------|------------------|------------------|-------------------------|-----------------------------|---------------------------|-------------------|-----------------------------------|----------------|
| `B1-P1A-OP001` | unchecked | `B2-P2A-A001-D001` | 1 | `resolve_calls` | `243.794553800 s` (`44.697337%`) | `calls=72976; files=765` | `243.794553800 s` | `243.794553800 s` | active / `ARCHITECT_CONSISTENCY_PASS` / `CODER_ACTIVE_PRE_EDIT_GATE` | unchecked | `A001`; sole Coder active pre-edit; no production/test edit at transition snapshot | not measured | not measured | `545.434182000 s` | not measured | not measured | process `653.475797100 s`; analyzer `630.160832200 s` | not measured | not measured | `0` | `0` | `0` | active; no disposition | `E:\Anvien\internal\resolution\resolve.go:resolveCall:92-94`; `internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> for each ir in w.files -> for each call in ir.Calls -> resolution.resolveCall` | `E2-P2A-A001DRILL1`, `E2-P2A-A001CURRENT1`, `E2-P2A-A001ARCH1`, `E2-P2A-A001PLAN1`, `E2-P2A-A001CONSISTENCY1` |
| `B1-P1A-OP001` | unchecked | `B2-P2A-A001-D002` | 2 | `resolve_accesses` | `204.093359700 s` (`37.418513%`) | `accesses=43826; files=765` | `204.093359700 s` | `204.093359700 s` | unchecked / queued | unchecked | none | not measured | not measured | `545.434182000 s` | not measured | not measured | process `653.475797100 s`; analyzer `630.160832200 s` | not measured | not measured | `0` | `0` | `0` | queued | `E:\Anvien\internal\resolution\resolve.go:resolveAccess:95-97`; `internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> for each ir in w.files -> for each access in ir.Accesses -> resolution.resolveAccess` | `E2-P2A-A001DRILL1` |
| `B1-P1A-OP001` | unchecked | `B2-P2A-A001-D003` | 3 | `resolve_type_annotations` | `92.316121800 s` (`16.925254%`) | `files=765; typeAnnotations=37389` | `92.316121800 s` | `92.316121800 s` | unchecked / queued | unchecked | none | not measured | not measured | `545.434182000 s` | not measured | not measured | process `653.475797100 s`; analyzer `630.160832200 s` | not measured | not measured | `0` | `0` | `0` | queued | `E:\Anvien\internal\resolution\resolve.go:resolveTypeAnnotation:98-100`; `internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> for each ir in w.files -> for each annotation in ir.TypeAnnotations -> resolution.resolveTypeAnnotation` | `E2-P2A-A001DRILL1` |
| `B1-P1A-OP001` | unchecked | `B2-P2A-A001-D004` | 4 | `project_resolution_outcomes` | `4.067254700 s` (`0.745691%`) | `graphNodes=67882; graphRelationships=106991; outcomes=150404; referencesBySourceScope=46600; referencesByTargetDef=46600; workUnits=418477` | `4.067254700 s` | `4.067254700 s` | unchecked / queued | unchecked | none | not measured | not measured | `545.434182000 s` | not measured | not measured | process `653.475797100 s`; analyzer `630.160832200 s` | not measured | not measured | `0` | `0` | `0` | queued | `E:\Anvien\internal\resolution\outcome.go:projectResolutionOutcomes:114-116`; `internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> resolution.projectResolutionOutcomes` | `E2-P2A-A001DRILL1` |
| `B1-P1A-OP001` | unchecked | `B2-P2A-A001-D005` | 5 | `emit_definition_nodes` | `0.387684000 s` (`0.071078%`) | `definitions=46608; exports=417; files=765; runs=1; workUnits=47790` | `0.387684000 s` | `0.387684000 s` | unchecked / queued | unchecked | none | not measured | not measured | `545.434182000 s` | not measured | not measured | process `653.475797100 s`; analyzer `630.160832200 s` | not measured | not measured | `0` | `0` | `0` | queued | `E:\Anvien\internal\resolution\emit.go:emitDefinitionNodes:77-81`; `internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> resolution.emitDefinitionNodes` | `E2-P2A-A001DRILL1` |
| `B1-P1A-OP001` | unchecked | `B2-P2A-A001-D006` | 6 | `finalize_resolution_outcomes` | `0.328439300 s` (`0.060216%`) | `outcomeMapEntries=150404` | `0.328439300 s` | `0.328439300 s` | unchecked / queued | unchecked | none | not measured | not measured | `545.434182000 s` | not measured | not measured | process `653.475797100 s`; analyzer `630.160832200 s` | not measured | not measured | `0` | `0` | `0` | queued | `E:\Anvien\internal\resolution\outcome.go:(*resolutionOutcomeCollector).finalize:110-113`; `internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> (*resolutionOutcomeCollector).finalize` | `E2-P2A-A001DRILL1` |
| `B1-P1A-OP001` | unchecked | `B2-P2A-A001-D007` | 7 | `build_binding_occurrence_index` | `0.172202300 s` (`0.031572%`) | `bindingLeaves=2674; definitionsVisited=46608; filePasses=1530; files=765; ownedDefIDsVisited=46608; ownerBindingsInspected=30782; runs=1; scopesVisited=11368; workUnits=139570` | `0.172202300 s` | `0.172202300 s` | unchecked / queued | unchecked | none | not measured | not measured | `545.434182000 s` | not measured | not measured | process `653.475797100 s`; analyzer `630.160832200 s` | not measured | not measured | `0` | `0` | `0` | queued | `E:\Anvien\internal\resolution\resolve.go:buildBindingOccurrenceIndex:62-65`; `internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> resolution.buildBindingOccurrenceIndex` | `E2-P2A-A001DRILL1` |
| `B1-P1A-OP001` | unchecked | `B2-P2A-A001-D008` | 8 | `finalize_typescript_authority_results` | `0.096993100 s` (`0.017783%`) | `authorityResults=5704` | `0.096993100 s` | `0.096993100 s` | unchecked / queued | unchecked | none | not measured | not measured | `545.434182000 s` | not measured | not measured | process `653.475797100 s`; analyzer `630.160832200 s` | not measured | not measured | `0` | `0` | `0` | queued | `E:\Anvien\internal\resolution\resolve.go:finalizeTypeScriptAuthorityResults:103-106`; `internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> resolution.finalizeTypeScriptAuthorityResults` | `E2-P2A-A001DRILL1` |
| `B1-P1A-OP001` | unchecked | `B2-P2A-A001-D009` | 9 | `emit_import_edges` | `0.088903700 s` (`0.016300%`) | `imports=4887; targetDefinitions=1381; targetFiles=5562; workUnits=11830` | `0.088903700 s` | `0.088903700 s` | unchecked / queued | unchecked | none | not measured | not measured | `545.434182000 s` | not measured | not measured | process `653.475797100 s`; analyzer `630.160832200 s` | not measured | not measured | `0` | `0` | `0` | queued | `E:\Anvien\internal\resolution\emit.go:emitImportEdges:83`; `internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> resolution.emitImportEdges` | `E2-P2A-A001DRILL1` |
| `B1-P1A-OP001` | unchecked | `B2-P2A-A001-D010` | 10 | `emit_typescript_external_symbols` | `0.011376900 s` (`0.002086%`) | `authorityResults=5495; resolvedRecords=0; uniqueResolvedSymbols=0; workUnits=5495` | `0.011376900 s` | `0.011376900 s` | unchecked / queued | unchecked | none | not measured | not measured | `545.434182000 s` | not measured | not measured | process `653.475797100 s`; analyzer `630.160832200 s` | not measured | not measured | `0` | `0` | `0` | queued | `E:\Anvien\internal\resolution\external_symbol.go:emitTypeScriptExternalSymbols:107-109`; `internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> resolution.emitTypeScriptExternalSymbols` | `E2-P2A-A001DRILL1` |
| `B1-P1A-OP001` | unchecked | `B2-P2A-A001-D011` | 11 | `emit_method_dispatch_edges` | `0.007767100 s` (`0.001424%`) | `heritageFacts=16; memberEntries=6510; ownerMemberOwners=1075` | `0.007767100 s` | `0.007767100 s` | unchecked / queued | unchecked | none | not measured | not measured | `545.434182000 s` | not measured | not measured | process `653.475797100 s`; analyzer `630.160832200 s` | not measured | not measured | `0` | `0` | `0` | queued | `E:\Anvien\internal\resolution\emit.go:emitMethodDispatchEdges:102`; `internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> resolution.emitMethodDispatchEdges` | `E2-P2A-A001DRILL1` |
| `B1-P1A-OP001` | unchecked | `B2-P2A-A001-D012` | 12 | `assemble_resolution_result` | `0.000000000 s` (`0.000000%`) | `authorityResults=5495; graphPointers=1; metricsValues=1; outcomes=150404; referenceIndexes=1; resultAssemblies=1` | `0.000000000 s` | `0.000000000 s` | unchecked / queued | unchecked | none | not measured | not measured | `545.434182000 s` | not measured | not measured | process `653.475797100 s`; analyzer `630.160832200 s` | not measured | not measured | `0` | `0` | `0` | queued | `E:\Anvien\internal\resolution\resolve.go:ResolveBoundInto:121-127`; `internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> construct resolution.Result return value` | `E2-P2A-A001DRILL1` |
| `B1-P1A-OP001` | unchecked | `B2-P2A-A001-D013` | 13 | `binding_accumulator_dispose` | `0.000000000 s` (`0.000000%`) | `accumulatedEntries=904; deferredExecutions=1; fileEntryBuckets=139; fileScopeBuckets=139; workUnits=1182` | `0.000000000 s` | `0.000000000 s` | unchecked / queued | unchecked | none | not measured | not measured | `545.434182000 s` | not measured | not measured | process `653.475797100 s`; analyzer `630.160832200 s` | not measured | not measured | `0` | `0` | `0` | queued | `E:\Anvien\internal\resolution\binding_accumulator.go:(*bindingAccumulator).dispose:71-74`; `internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto deferred closure -> (*bindingAccumulator).dispose` | `E2-P2A-A001DRILL1` |
| `B1-P1A-OP001` | unchecked | `B2-P2A-A001-D014` | 14 | `emit_heritage_compatibility_edges` | `0.000000000 s` (`0.000000%`) | `heritageFacts=16; runs=1` | `0.000000000 s` | `0.000000000 s` | unchecked / queued | unchecked | none | not measured | not measured | `545.434182000 s` | not measured | not measured | process `653.475797100 s`; analyzer `630.160832200 s` | not measured | not measured | `0` | `0` | `0` | queued | `E:\Anvien\internal\resolution\emit.go:emitHeritageCompatibilityEdges:85-89`; `internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> for each w.heritage -> resolution.emitHeritageCompatibilityEdges` | `E2-P2A-A001DRILL1` |
| `B1-P1A-OP001` | unchecked | `B2-P2A-A001-D015` | 15 | `emit_unresolved_heritage_diagnostics` | `0.000000000 s` (`0.000000%`) | `runs=1; unresolvedHeritageFacts=45` | `0.000000000 s` | `0.000000000 s` | unchecked / queued | unchecked | none | not measured | not measured | `545.434182000 s` | not measured | not measured | process `653.475797100 s`; analyzer `630.160832200 s` | not measured | not measured | `0` | `0` | `0` | queued | `E:\Anvien\internal\resolution\resolve.go:emitUnresolvedHeritageDiagnostics:82`; `internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> resolution.emitUnresolvedHeritageDiagnostics` | `E2-P2A-A001DRILL1` |
| `B1-P1A-OP001` | unchecked | `B2-P2A-A001-D016` | 16 | `finalize_resolution_metadata` | `0.000000000 s` (`0.000000%`) | `graphNodes=67882; graphRelationships=106991; metadataUpdates=1; workUnits=174874` | `0.000000000 s` | `0.000000000 s` | unchecked / queued | unchecked | none | not measured | not measured | `545.434182000 s` | not measured | not measured | process `653.475797100 s`; analyzer `630.160832200 s` | not measured | not measured | `0` | `0` | `0` | queued | `E:\Anvien\internal\resolution\resolve.go:ResolveBoundInto:118-120`; `internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> graphhealth.SetResolutionMetadata` | `E2-P2A-A001DRILL1` |
| `B1-P1A-OP001` | unchecked | `B2-P2A-A001-D017` | 17 | `runtime_setup` | `0.000000000 s` (`0.000000%`) | `graphAllocations=0; newEmitterInvocations=1; resolveBoundIntoInvocations=1` | `0.000000000 s` | `0.000000000 s` | unchecked / queued | unchecked | none | not measured | not measured | `545.434182000 s` | not measured | not measured | process `653.475797100 s`; analyzer `630.160832200 s` | not measured | not measured | `0` | `0` | `0` | queued | `E:\Anvien\internal\resolution\resolve.go:ResolveBoundInto:66-75`; `internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> conditional graph.New -> resolution.newEmitter` | `E2-P2A-A001DRILL1` |

`E2-P2A-AnnnDRILL1` proves boundary validity, complete child coverage, denominators, ordering, and benchmark/plan checklist cardinality. If the parent has `10` measured children, this table has exactly `10` child rows and plan has exactly `10` nested child checkboxes before `E2-P2A-AnnnARCH1`. Missing/extra rows or checklist items block Architect. A concrete `BLOCKED` child stays unchecked.

### A001 Active-Child Numeric Contract

No benchmark/analyze capture was run during the Planner refresh or this Coder-transition state sync. The complete 17-row values and order above are unchanged. Only D001 is eligible for later deep attribution after the authorized implementation/build/test-validation sequence; D002-D017 retain their coarse `child_total_elapsed` values and remain queued/unchecked. `E2-P2A-A001CONSISTENCY1` is control-state evidence only and supplies no numeric value, candidate, speed result, disposition, Supervisor result, or commit authority.

| A001 metric | Current D001 value | Candidate after | Numeric role / constraint | Evidence |
|-------------|-------------------|-----------------|---------------------------|----------|
| `child_total_elapsed` | `243.794553800 s` | not measured | sole child ranking, before/after, and child-improvement acceptance value | `E2-P2A-A001DRILL1`, `E2-P2A-A001CURRENT1`, `E2-P2A-A001ARCH1` |
| `prepare_lookup_elapsed` | not measured | not measured | path/input normalization, index/import scanning, and candidate preparation/lookup before semantic decision | `E2-P2A-A001ARCH1` |
| `resolution_decision_elapsed` | not measured | not measured | binding/member/package/authority/target/proof/outcome decision work not already owned by lookup | `E2-P2A-A001ARCH1` |
| `graph_mutation_elapsed` | not measured | not measured | node/relationship/reference/property/metadata creation or merge | `E2-P2A-A001ARCH1` |
| `diagnostic_outcome_elapsed` | not measured | not measured | resolution outcome and unresolved/diagnostic recording | `E2-P2A-A001ARCH1` |
| `metrics_bookkeeping_elapsed` | not measured | not measured | metrics/counter bookkeeping not already owned elsewhere | `E2-P2A-A001ARCH1` |
| `residual_elapsed` | not measured | not measured | remaining D001 elapsed after the five explicit exclusive groups | `E2-P2A-A001ARCH1` |
| `overlap_count` for the future D001 exclusive breakdown | not measured | not measured | must equal `0`; the already-recorded `0` above belongs to the coarse parent interval set and is not fabricated as a D001 deep value | `E2-P2A-A001ARCH1` |

Mandatory future conservation:

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

### A001 Pending Work-Count And Observer-Overhead Control

| A001 work counter | Current before | Candidate after | Comparability constraint | Evidence |
|-------------------|----------------|-----------------|--------------------------|----------|
| files | `765` | not measured | same accepted work | `E2-P2A-A001CURRENT1` |
| calls | `72976` | not measured | same accepted work; no per-call log/timer artifact | `E2-P2A-A001CURRENT1` |
| import-claim helper invocations | not measured | not measured | cumulative local counter, merged once | `E2-P2A-A001ARCH1` |
| global-import candidates visited | not measured | not measured | proves whole-import scans are removed | `E2-P2A-A001ARCH1` |
| import-path normalizations at claim lookup sites | not measured | not measured | count only the two claim lookup sites | `E2-P2A-A001ARCH1` |
| index entries | not measured | not measured | bounded by imports; store indices, not duplicate import objects | `E2-P2A-A001ARCH1` |
| unique lookup keys | not measured | not measured | memory remains `O(imports + unique keys)` | `E2-P2A-A001ARCH1` |
| matching-bucket candidates visited | not measured | not measured | accepted traversal follows original index order | `E2-P2A-A001ARCH1` |
| resolution targets considered | not measured | not measured | same semantic/target work | `E2-P2A-A001ARCH1` |
| references / relationships emitted | not measured | not measured | exact affected output equivalence | `E2-P2A-A001ARCH1` |
| outcomes / diagnostics recorded | not measured | not measured | exact outcome/diagnostic equivalence | `E2-P2A-A001ARCH1` |
| affected graph node / relationship / property inventories | not measured | not measured | exact affected graph/output equivalence; moving cross-capture counts cannot substitute | `E2-P2A-A001ARCH1` |
| full-scan fallback | not measured | not measured | if recorded for an accepted candidate, must be `0` | `E2-P2A-A001ARCH1` |

Counters prove comparable work and never replace wall time. The future capture must use the same instrumentation identity before/after, the smallest cumulative local accumulators, and one merge per accumulator. Invalid or mixed instrumentation cannot rank or prove gain. No standalone timer/log/trace may be emitted for each of the `72976` calls.

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
| `B2-P2A-FINAL1` | current child-drilldown workload: `2,208` scanned / `765` parsed / `0` failed; graph `123,376 / 170,075`; moving-capture counts remain relative | immutable uninstrumented reference `605.732722 s`; carried P2 basis `618.4358209 s`; current drilldown process wall `653.475797100 s` is attribution-only | not measured | not measured | not measured | `0/30` | `0/17`; active child `B2-P2A-A001-D001` | none recorded | pending; no speedup claim | `E2-P2A-A001DRILL1`; `E2-P2A-FINALTIME1`, `E2-P2A-EXHAUST1` pending |

## B3 - P3 Benchmarks

P3-A/P3-B/P3-C are final review, cleanup, detect, commit, and handoff operations. They add no product-performance metric. Their pass/fail and one-commit boundary belong in evidence.

## Non-Benchmarkable Notes

- P6-D commit identity, Architect decisions, Planner refreshes, source/test/build results, Supervisor verdicts, cleanup, detect, commit, and handoff are evidence rather than measurements.
- `E2-P2A-A001CONSISTENCY1` records the completed Owner-required consistency/Main transition and sole Coder pre-edit identity; it is not a benchmark observation and changes none of the 17 child values or their order.
- Historical/protected values are provenance only and never enter `Current accepted`, ranking, attempt count, or final speedup.
- `SYSTEM_CHARACTERISTIC` records one child's terminal retained remaining cost after three complete unsuccessful attempts at one accepted baseline. It does not close the parent, omit smaller rows, or reduce the final requirement that total graph generation must be lower on the same workload.
