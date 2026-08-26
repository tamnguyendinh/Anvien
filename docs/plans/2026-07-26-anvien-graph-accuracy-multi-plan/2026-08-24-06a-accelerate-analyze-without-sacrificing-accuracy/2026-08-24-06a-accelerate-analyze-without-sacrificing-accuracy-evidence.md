# Child 06A Accelerate Analyze Without Sacrificing Accuracy Evidence Ledger

## Metadata

- Date: `2026-08-24`
- Current state: `A003_CHECKPOINT_COMPLETE / WAL_FIX_CHECKPOINT_COMPLETE / A004 SUPERVISOR_PASS / NO_KEEP / ROLLBACK_COMPLETE / A005_CURRENT_BASIS_RECORDED / A005_ATTRIBUTION_COMPLETE / A005_ARCHITECT_COMPLETE / A005_PLAN_COMPLETE / A005_MAIN_VERIFIED / CODER_A005_READY_FOR_MEASUREMENT_HANDOFF / MAIN_HANDOFF_PASS / A005_FROZEN_BUILD_FAILED / OVERLAY_BUNDLE_ADAPTATION_RECOVERY_READY / MEASUREMENT_LOCKED / D001_STREAK_1`; accepted A003 remains baseline
- Plan: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy/2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy-plan.md`
- Plan rules: [plan-rules.md](plan-rules.md)
- Evidence: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy/2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy-evidence.md`
- Benchmark: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy/2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy-benchmark.md`
- Actual status: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy/2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy-actual-status.md`
- Roadmap: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-26-anvien-graph-accuracy-roadmap.md`
- Predecessor: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-plan.md`
- Successor: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-07-cross-surface-acceptance-and-target-validation/2026-07-28-07-cross-surface-acceptance-and-target-validation-plan.md`
- Method authority: `E:\Anvien\reports\Investigation\rp_main_260824_204615_child06a_optimization_method_discussion.md`
- Required provenance and handoff reference: `E:\Anvien\reports\Investigation\rp_huong_dan_lam_viec_de_slice_p6e_tien_hanh_nhanh_hon.md`

## Evidence Rules

- Evidence explains why a command, measurement, architecture decision, plan refresh, implementation, validation, review, disposition, or transition is valid. Numeric values and rankings live in benchmark.
- Elapsed wall-clock time is the controlling performance proof. Secondary resource/count evidence may explain cause or comparability but cannot select a bottleneck or justify `KEEP`.
- Record every result immediately. A pending evidence ID is not completed proof and cannot authorize a later state.
- P1-A records only the real `anvien analyze` command/workload identity, total timing validity, operation boundaries/denominators, conditional minimum instrumentation when used, initial equivalence, and ranking derivation. It contains no production-optimization proof and does not predeclare an unmeasured options/workload/cache tuple.
- Exactly one visible measurement executor produces each accepted P1-A runtime capture. After acceptance, up to three ACTIVE read-only measurement analysts may share that capture for independent measured problems; analysts do not edit source or launch duplicate competing benchmark processes.
- Existing timing/benchmark/profile capability is used first. If one required P1-A timing is missing, Main assigns exactly one visible sequential instrumentation-writer lane outside the read-only analysts, and that lane is the sole source writer for the conditional branch.
- `E1-P1A-INSTR1` must bind the writer, exact owner and owned bytes, fresh required graph/file-detail/impact, canonical full-build PASS before instrumented timing use, runtime/instrumentation identity, overhead, denominator/comparability, output equivalence, and like-for-like measurement state. It then records exactly one terminal disposition: carry exact instrumentation ownership into the first refreshed P2-A attempt, or remove the exact owned bytes, full-build again, and re-establish and remeasure the accepted timing basis. A half-open or mixed-instrumentation state cannot authorize ranking or P2-A.
- The conditional instrumentation branch is not a phase, implementation slice, optimization attempt, harness project, Architect/Supervisor gate, progress claim, or separate commit/report system.
- P2-A is one implementation slice with a monotonically numbered attempt sequence `A001`, `A002`, and onward. Attempt count is created by actual production work; it is not predeclared.
- Every production attempt requires an active parent, its complete measured child list, the exact selected child, current child/parent/total timing basis, denominator/cause/owner/complete call path, a new Visible Architect decision, Planner refresh, fresh owner impact, Coder source/tests/build evidence, child/parent/end-to-end remeasurement, and a new Visible Supervisor accuracy/equivalence result before disposition.
- `AnnnDRILL1` is mandatory before the first child Architect attempt of a parent and must remain current for later attempts. It proves the selected parent boundary and the complete measured child list in descending absolute elapsed-time order, including smaller child costs. The benchmark child-row count and nested plan-checklist child-item count must match exactly; if drill-down measures `10` children, both files contain exactly `10` corresponding rows/items before Architect.
- An Architect decision is attempt-local. It must record exact cause, technical direction, allowed owners, expected observable gain, preserved invariants, validation/resource boundary, and rollback condition. It expires at disposition and cannot authorize later code.
- Planner refresh evidence must point to the concrete current-attempt update in the four plan ledgers. Coder may start only after that refresh exists and may not infer, reuse, broaden, or replace architecture direction.
- The per-attempt Supervisor reviews the exact changed candidate and affected accuracy/correctness/equivalence/output/lifecycle evidence. It neither edits code nor audits documentation wording, report hashes, or evidence chains.
- `KEEP` requires selected-child improvement, retained parent improvement, retained end-to-end benefit, and that attempt's Supervisor `PASS`. Only `KEEP` evidence authorizes baseline promotion; it resets the active child's streak but does not terminalize or check that child.
- After Supervisor `PASS` plus Main `KEEP`, an accepted attempt may receive the scoped detected progress checkpoint required by `plan-rules.md`; that checkpoint records accepted bytes but does not close P2-A, check a parent/child, or replace P3-C closure.
- Each `AnnnCURRENT1`/`AnnnDRILL1`/`AnnnPLAN1` set must state the exact attempt goal: reduce elapsed time of the selected child, its parent, and total graph generation from the current accepted baseline while preserving required invariants. Supporting drill-down instrumentation/reporting is never the goal.
- Supervisor `REJECT` requires exact attempt-owned restoration/retention of the last accepted baseline. Rejection evidence and the violated invariant return to a new Architect; a new Planner refresh precedes any further Coder action. Rejected bytes cannot be ranked or used to switch bottlenecks.
- Track consecutive unsuccessful attempts per exact child row and current accepted baseline. Supervisor `REJECT`, no retainable child/parent/end-to-end gain, `REWORK`, or `ROLLBACK` without `KEEP` increments the streak. `KEEP` resets it to `0` on the promoted baseline.
- On the third consecutive attempt without `KEEP`, evidence must cite the exact parent/child rows, denominator, current retained child/parent times, all three attempt families, and why none was retainable. Then only that child may receive terminal `SYSTEM_CHARACTERISTIC`; evidence and plan mark it before selecting the next-largest unchecked child of the same parent.
- `SYSTEM_CHARACTERISTIC` preserves the last accepted correct state. It is not a speedup claim, accuracy waiver, permission to retain a rejected candidate, or parent completion. A parent closes only after every measured child is terminal and checked.
- After `SYSTEM_CHARACTERISTIC`, accepted child/parent/end-to-end remeasurement and complete child/top-level list refresh must be evidenced before the next child is selected. If accepted measurement exposes a missing parent or child, evidence must bind its new benchmark row and matching unchecked plan checkbox immediately.
- A concrete unavailable authority/dependency/evidence `BLOCKED` row remains unchecked and blocks the parent/P2-A. It cannot terminalize a measured row or skip a small elapsed-time cost.
- Fresh required Anvien graph/file-detail/impact runs immediately before editing the exact measured owner. Production implementation precedes authorized test edits; build-holder preflight and canonical full-build `PASS` then precede focused/package test execution; later child -> parent -> end-to-end remeasurement and Supervisor remain separate.
- A full-package test is never relabeled `PASS`. When an isolated candidate rerun and an exact-HEAD overlay rerun reproduce the same out-of-scope golden failure, record it as a pre-existing unrelated stale validation failure: neither A001 failure nor A001 validation proof, and never permission to edit its preserve-only owner.
- Raw debug/profile material remains under repository-local `E:\Anvien\.tmp`. The four ledgers are the durable recording system; do not create a per-run evidence tree, cost-map report, report audit, or evidence-about-evidence chain.
- `E3-P3A-REVIEW1` is the one final whole-candidate Supervisor result. It is additional to, and cannot replace, the required per-attempt `E2-P2A-AnnnREVIEW1` results.
- `E3-P3B-CLEAN1` opens no review and cannot change accepted production/test bytes. `E3-P3C-DETECT1` precedes the sole `E3-P3C-COMMIT1`; `E3-P3C-HANDOFF1` opens Child 07 from that commit.

### Evidence ID Naming

- `E0-P0A-*` records current truth, predecessor/order, method, and provenance.
- `E1-P1A-*` records initial detailed measurement and ranking.
- `E2-P2A-Annn*` records one actual production attempt, where `nnn` is a zero-padded chronological sequence beginning with `001`.
- `E2-P2A-FINAL*` and `E2-P2A-EXHAUST1` record stable final P2-A proof.
- `E3-P3A-*`, `E3-P3B-*`, and `E3-P3C-*` record final whole-candidate review, cleanup, detect/commit, and handoff.
- Never reuse an attempt number or evidence ID after rejection, rollback, compaction, rotation, or another edit to the same bottleneck.

## Legacy Evidence Provenance Mapping

Legacy IDs remain searchable provenance. Current work records only current IDs and, when relevant, cites the legacy source alongside them.

| Legacy Child 06 evidence | Current Child 06A evidence | Mapping rule |
|--------------------------|----------------------------|--------------|
| `E6-P6E-RC1`, `E6-P6E-ARCH1`, `E6-P6E-WORKFLOW1`, `E6-P6E-FEAS1` | `E0-P0A-PROVENANCE1` | historical attribution/design/workflow/feasibility only; no current target or acceptance authority |
| `E6-P6E-ANCHOR1` | `E0-P0A-ANCHOR1` | accepted P6-D closure at `81163e39718b94a509e41114cada224e8f269e36` |
| `E6-P6E-PLAN1` | `E0-P0A-ORDER1`, `E0-P0A-METHOD1` | independent Child 06A topology and superseding empirical method |
| `E6-P6E-AUTH1` | `E1-P1A-IDENTITY1` | current real command/runtime/workload validity, recorded only when measured |
| `E6-P6E-COSTMAP1`, `E6-P6E-BASELINE1` | `E1-P1A-TOTAL1`, `E1-P1A-OPS1`, `E1-P1A-RANK1` | initial total/operation timings and absolute ranking; old protocol is not executable |
| `E6-P6E-IMPACT1` | dynamic `E2-P2A-AnnnIMPACT1` | fresh exact-owner impact immediately before each actual edit |
| legacy U1/U2/U3/U4 source/build/test/benchmark/equivalence families | dynamic `E2-P2A-AnnnARCH1/PLAN1/IMPACT1/SRC1/TEST1/BUILD1/DIRECT1/PARENT1/E2E1/EQUIV1/REVIEW1/DECISION1` | historical idea family maps only if current measurement selects its real operation; no family is pre-opened |
| `E6-P6E-REBASE1`, `E6-P6E-REBASE2`, `E6-P6E-PARETO1` | dynamic `E2-P2A-AnnnRANK1` and living benchmark ranking | accepted-state remeasurement/reranking after dispositions |
| `E6-P6E-DETERMINISM1/FRESHNESS1/PARITY1/FAILURE1/PUBLICATION1/RESOURCE1` | per-attempt `E2-P2A-AnnnEQUIV1/REVIEW1` plus `E2-P2A-FINALEQUIV1` | exact affected checks per attempt and aggregate final equivalence |
| `E6-P6E-REVIEW1`, dependent `E6-PNA-REVIEW1` | `E3-P3A-REVIEW1` | one final whole-candidate Supervisor; per-attempt reviews are new stronger gates |
| `E6-P6E-CLEAN1`, dependent `E6-PNB-CLEAN1` | `E3-P3B-CLEAN1` | exact cleanup without another review |
| `E6-P6E-DETECT1`, dependent `E6-PNC-DETECT1` | `E3-P3C-DETECT1` | sole post-cleanup detect |
| `E6-P6E-COMMIT1`, performance-dependent `E6-PNC-COMMITS1` | `E3-P3C-COMMIT1` | sole Child 06A implementation commit |
| performance-dependent `E6-PNC-HANDOFF1` | `E3-P3C-HANDOFF1` | Child 07 opens from Child 06A closure commit |

## E0 - P0 Evidence

Matching plan item(s): `P0-A`

- `E0-P0A-ANCHOR1` — recorded: Child 06 P6-D is accepted and closed at isolated commit `81163e39718b94a509e41114cada224e8f269e36`. Child 06A inherits that correctness/runtime/reader boundary without reopening it.
- `E0-P0A-ORDER1` — recorded: direct campaign order is Child 06 -> Child 06A -> Child 07; former performance-dependent aggregate responsibilities no longer block closed Child 06, and Child 07 waits for the Child 06A closure commit.
- `E0-P0A-TRUTH1` — recorded: Child 06A has no detailed current graph-generation timing table, ranked bottleneck, production implementation, attempt, accepted speedup, final Supervisor result, detect, cleanup, or implementation commit.
- `E0-P0A-METHOD1` — recorded: FULL RAW `E:\Anvien\reports\Investigation\rp_main_260824_204615_child06a_optimization_method_discussion.md` consolidates the real `anvien analyze` measure/rank/optimize/remeasure method, one visible initial-capture executor, the conditional P1-A sole-writer/build-before-use/comparability/disposition branch, fresh Architect + Planner refresh + Coder + post-measurement Supervisor for every production attempt, elapsed-time authority, the mandatory benchmark list, and the three-unsuccessful-attempt terminal rule. The long rule body is stored losslessly in linked auxiliary `plan-rules.md`; the four standard ledgers keep their original responsibilities and no new phase, implementation slice, Architect/Supervisor gate, ledger, or permission boundary exists.
- `E0-P0A-PROVENANCE1` — recorded: legacy P6-E evidence and benchmark families remain traceable through the mapping tables, but historical attribution and fixed solution ideas cannot select current work.

No functional analyze, build, test, benchmark, implementation, Architect attempt, Supervisor attempt, detect, cleanup, or commit was performed by this documentation rewrite.

## E1 - P1 Evidence

Matching plan item(s): `P1-A`

| Evidence ID | Required proof | Current status |
|-------------|----------------|----------------|
| `E1-P1A-IDENTITY1` | exact visible measurement-executor identity, accepted capture identity, real executable/runtime, `anvien analyze` command/options, repository/input workload, relevant cache regime, timestamps, exit, output identity, and denominator | recorded for capture `child06a-p1a-initial-20260824-225900`; exact proof below |
| `E1-P1A-TOTAL1` | current total graph-generation elapsed time from the accepted real path | complete: immutable uninstrumented process/analyzer/phase reference `605.732722 / 602.5278811 / 587.2385976 s`; carried instrumented P2 basis `618.4358209 / 615.8656873 / 600.8917213 s`; cross-capture arithmetic is explicitly non-comparable performance evidence |
| `E1-P1A-OPS1` | real internal operation boundaries, elapsed times, denominators, and validity for every ranked row | complete: `30` rows (`21 analyzer_internal`, `1 analyzer_outer`, `8 cli_outer`), unique names, non-negative durations/denominators, exact 15-phase identity, exact analyzer-internal sum, and recorded process residual |
| `E1-P1A-INSTR1` | exact missing timing and why existing capability is insufficient; sole visible sequential writer; exact owner/owned bytes; fresh graph/file-detail/impact; canonical full-build PASS before use; accepted runtime/instrumentation identity; overhead, denominator, like-for-like comparability, and output equivalence; exactly one carry-forward ownership or remove/rebuild/remeasure disposition | complete / terminal `CARRY_TO_FIRST_P2A_REFRESH`: exact writer-owned `+349/-53` instrumentation remains carried so P2 measurements keep the same instrumentation state; no overhead or speed claim is made versus the different-HEAD/different-workload reference |
| `E1-P1A-EQUIV1` | initial workload/output/persistence/publication validity sufficient to use the measurement | complete for P1 transition: exit `0`, valid JSON/stdout/stderr, all `15` phases, graph/Ladybug/meta publication, zero parse/DB failures, and exact cross-capture mismatches recorded; hashes/counts are not falsely claimed equal across different HEAD/workload states |
| `E1-P1A-RANK1` | complete benchmark-owned top-level list of every measured real operation in descending elapsed-time order, including smaller rows, with current rank/time, denominator, meaningful share/delta, processing state, proven owner/call path when known, disposition/evidence, and an exact one-to-one unchecked parent checklist mirror in plan | complete: `B1-P1A-OP001..OP030` and exactly `30` matching unchecked plan parent items; `OP001 resolution` is the largest active parent |

### P1-A Result Record

Append the first completed measurement record here and update benchmark/status immediately. Current rows: none.

| Evidence ID | Recorded at | Visible executor / accepted capture | Exact command / exit | Runtime and workload identity | Operation boundary / denominator | Benchmark rows | Output validity | Instrumentation branch / disposition | Next exact action |
|-------------|-------------|-------------------------------------|----------------------|-------------------------------|----------------------------------|----------------|-----------------|--------------------------------------|-------------------|
| `E1-P1A-IDENTITY1/TOTAL1/OPS1/EQUIV1` | `2026-08-24T22:58:05.2715182+07:00` -> `2026-08-24T23:08:11.1166985+07:00` | executor `01a033a5-3a44-7c43-9aca-c85e3e932a0f`; capture `child06a-p1a-initial-20260824-225900` | exact command below; exit `0` | repo-owned runtime `1.2.8`; workload `E:\Anvien`; forced cold analyzer outputs; isolated E-only process environment | process wall/CPU plus all `15` emitted phases; exact workload/output denominators below | `B1-P1A-TOTAL` plus provisional `B1-P1A-PHASE-*`; no `OPnnn` queue rows | graph/Ladybug/meta and raw capture identities below; tracked/index boundary preserved | `E1-P1A-INSTR1` OPEN, no writer/edit/build/disposition; ranking blocked | Main assigns the separate visible sequential sole writer or records a concrete blocker; this executor stops |
| `E1-P1A-INSTR1` | `2026-08-25T03:25:06.0193408+07:00` -> `2026-08-25T03:36:46.1906114+07:00` full-build validation window | sole writer `01a0348f-3bb7-7c70-8ec5-61176ce00591`; no performance capture | `pwsh -NoLogo -NoProfile -NonInteractive -ExecutionPolicy Bypass -File E:\Anvien\scripts\full-build.ps1`; exit `0` | HEAD `8df19258fbcb18e14841bf0dae036400aa9b22a3`; runtime `1.2.8`, `73,764,864` bytes, SHA-256 `62920CBF15921EF8A6D2FAC671776BAA3C312EFEA1BEED53B721C4AFF5E1B6C5` | flat non-overlapping `operations` contract and denominators below; no elapsed operation values recorded | unchanged: provisional rows only; no `OPnnn` queue rows | source/test diff-check clean; maintenance analyze completed with `2,200` scanned, `765` parsed, `0` failed, `123,273` nodes, `169,972` relationships | build-before-use complete; capture/comparability/equivalence/disposition pending | hand off `READY_FOR_INSTRUMENTED_CAPTURE` to Main; Main alone releases preserved measurement executor for the separate accepted tuple |
| `E1-P1A-INSTR1/OPS1/RANK1/EQUIV1` | `2026-08-25T04:11:18.7266556+07:00` -> `2026-08-25T04:21:37.1624765+07:00` | sole executor `01a03597-6afb-73e1-bee6-4306836cce01`; capture `child06a-p1a-instrumented-20260825-040400`; PID `14956`; one launch/no retry | accepted benchmark command; exit `0`; process wall/CPU `618.4358209 / 799.4062500 s` | HEAD/runtime/hash above; E-only isolated environment; `2,200` scanned, `765` parsed, `0` failed; graph `123,277 / 169,976` | `30` exact operation rows; phase `600.8917213 s`; analyzer `615.8656873 s`; all operations `618.2628497 s`; process residual `0.1729712 s` | `B1-P1A-OP001..OP030`; exactly `30` matching plan parent items | graph/Ladybug/meta published; exact artifacts/hashes and cross-capture mismatches below | terminal `CARRY_TO_FIRST_P2A_REFRESH`; cross-capture deltas are non-comparable arithmetic only | P2-A `PARENT_DRILLDOWN_PENDING` on `B1-P1A-OP001 resolution`; no child/attempt/Architect yet |

### E1-P1A Accepted Capture Detail

`E1-P1A-IDENTITY1`:

- Main authority: task `01a03472-893f-7c92-95ac-7221631c002e`; visible measurement executor `01a033a5-3a44-7c43-9aca-c85e3e932a0f`; exactly one capture process was launched.
- Exact executable: `E:\Anvien\anvien\bin\anvien.exe`, version `1.2.8`, `73,750,016` bytes, SHA-256 `D4EC8B58C41B9F0A95359CFC014DB07D11653F25736728549555F74D204ABD19`. This is the actual current P1-A runtime identity; the older P6-D executable identity was not substituted.
- Exact command:

```text
E:\Anvien\anvien\bin\anvien.exe analyze E:\Anvien --force --json --progress --benchmark-json E:\Anvien\.tmp\child06a_p1a_initial_20260824_225900\benchmark.json --benchmark-label child06a-p1a-initial-20260824-225900
```

- Cache/runtime regime: `--force` removed the existing analyzer graph/Ladybug outputs; process `HOME`, `USERPROFILE`, `TEMP`, `TMP`, `APPDATA`, `LOCALAPPDATA`, `XDG_CACHE_HOME`, and `XDG_CONFIG_HOME` were isolated beneath `E:\Anvien\.tmp\child06a_p1a_initial_20260824_225900`; process-local Git safe-directory was `E:/Anvien`; no OS filesystem-cache flush was performed.
- Machine/process observation: initial preflight at `2026-08-24T22:57:14.5940676+07:00` found CPU `50%`, available memory `21,625 MiB`, six installed `anvien mcp` service processes, and zero `anvien analyze`/`benchmark-compare` competitors. The immediate capture-start observation was CPU `15%`, available memory `21,648 MiB`, competitors `0`; capture-end observation was CPU `41%`, available memory `21,165 MiB`.
- Capture start/end/exit: `2026-08-24T22:58:05.2715182+07:00` / `2026-08-24T23:08:11.1166985+07:00` / `0`.

`E1-P1A-TOTAL1`:

- Total process wall: `605.732722 s`.
- Analyzer-internal `totalDuration`: `602.5278811 s`.
- Process CPU: `803.093750 s` = user `789.515625 s` + kernel `13.578125 s`; CPU/wall `132.582%`.
- Go memory counters from the benchmark: start alloc `1,431,736` bytes; end alloc `996,685,256` bytes; maximum observed `Sys` `2,387,860,088` bytes. Peak process working set and allocated object count were not captured/instrumented and are not inferred.

`E1-P1A-OPS1` partial proof:

- `internal/analyze.runPhase` appends one duration after each synchronous phase returns. Accepted stderr records exactly `15` ordered `done` events: scan, structure, documents, cobol, parse, routes, tools, orm, cross-file binding, resolution, MRO, communities, processes, semantic enrichment, and DB load.
- The `15` phase durations sum to `587.2385976 s`. Analyzer-internal total minus phase sum is `15.2892835 s`; process wall minus analyzer-internal total is `3.2048409 s`; process wall minus phase sum is `18.4941244 s`.
- Source boundary proof: `internal/analyze/analyze.go` times the named phases but leaves `Graph.Compact`, DB runner resolve/close, and `writeGraphSnapshot` outside `runPhase`; `internal/cli/command.go` performs benchmark/profile completion, registry/meta recording, generated AI context, file projection, and JSON output after `analyze.Run`. Existing benchmark JSON therefore cannot assign the residual to real non-overlapping operations or provide a Graph JSON persistence/publication elapsed value.
- Every emitted phase duration and available denominator is recorded in benchmark under provisional `B1-P1A-PHASE-*` metrics. Those metrics are not promoted to `B1-P1A-OPnnn` because the missing residual could change the complete absolute ordering.

Workload/output denominator:

- repo `E:\Anvien`; `2,196` scanned / `765` parsed code / `0` failed; parser bytes `7,086,653`;
- graph `123,075` nodes / `169,739` relationships; file projection `20,363` dependency edges / `479` unresolved;
- DB load `123,075` node rows / `169,739` relationship rows, COPY calls `18 / 93`, fallback/failure/skip `0 / 0 / 0`.

Raw capture and output identities:

| Artifact | Bytes | SHA-256 |
|----------|------:|--------|
| `E:\Anvien\.tmp\child06a_p1a_initial_20260824_225900\benchmark.json` | `10,446` | `B235AAFAE36AD4C13ED8FCE340D6428CB235DADC9ACE657003C395072F4D492D` |
| `E:\Anvien\.tmp\child06a_p1a_initial_20260824_225900\stdout.json` | `9,130` | `57718F513AC7F72E03D86D06C17EEBE7BB990899636C68616FAAB5E4BD460F0E` |
| `E:\Anvien\.tmp\child06a_p1a_initial_20260824_225900\stderr.log` | `213` | `0FDCF15222C02C25771C98E17D9220C89FDCD0294137A6AE4D9B3AAFD4761ADC` |
| `E:\Anvien\.anvien\graph.json` | `673,680,603` | `67F558C36EDB6BB01D20DDCDF96BE7490C734D272BB436D24CC5E3FC28EC6091` |
| `E:\Anvien\.anvien\lbug` | `389,046,272` | `49A2E79F687DA49C70778941F9E29B79F45019E1B99400D677EF656DA64C1195` |
| `E:\Anvien\.anvien\meta.json` | `255` | `D93267E62954DDD667BED58CFCD079A2DA7D897032ED62DB1F2294093D6C60E6` |

`E1-P1A-EQUIV1`:

- Exit is `0`; stdout decodes as JSON and names `E:\Anvien\.anvien\graph.json`; stderr contains all `15` phase completions and no error.
- `meta.json` binds repo `E:\Anvien`, commit `1c5de4ef6875a5e7b3329f04dafd1189c7622e4d`, and the same `2,196` files / `123,075` nodes / `169,739` edges / `1,529` communities / `691` processes.
- Git index remained empty. Tracked dirty rows were `10` before and after; the post-capture set remains the preserved Owner deletion plus existing roadmap/Child 06/Child 07 rows. `AGENTS.md` and `CLAUDE.md` have no tracked diff.

`E1-P1A-RANK1` is complete. The immutable uninstrumented reference remains preserved, while the instrumented carry capture supplies the current rankable operation basis. The two captures are not treated as a paired performance comparison because HEAD, binary identity, workload, and output counts differ. Benchmark contains all `30` real operation rows and plan contains exactly `30` matching unchecked parent items.

### Conditional P1-A Instrumentation Contract

This table is populated only if `E1-P1A-INSTR1` opens. It is evidence for one conditional branch, not a new phase, implementation slice, attempt, gate, or report system.

| Required pointer | Proof required before P1-A transition |
|------------------|---------------------------------------|
| Missing timing | CLOSED: the instrumented capture measures all `30` emitted operation rows. Analyzer-internal operation sum equals `615.8656873 s` `totalDuration` exactly; all-operation coverage is `618.2628497 s`, leaving only `0.1729712 s` process-level wrapper/final-rewrite residual, recorded rather than inferred. |
| Writer / owner | SOLE WRITER `01a0348f-3bb7-7c70-8ec5-61176ce00591`. Exact rollback map: `internal/analyze/analyze.go +117/-15` (`Metrics`, `OperationMetric`, `Run`, `runPhase`, operation helpers); `internal/cli/command.go +157/-38` (`newAnalyzeCommand`); `internal/cli/command_test.go +75/-0` (`TestAnalyzeCommandRunsGoPipelineAndWritesBenchmark`). Total `+349/-53`. |
| Pre-edit topology | COMPLETE. Production file-detail was consumed for both exact owners. Accepted impact artifacts under `E:\Anvien\.tmp\child06a_p1a_instrumentation`: `analyze-file-impact.stdout.json` CRITICAL (`52` impacted, `15` files, `1` flow); `command-file-impact.stdout.json` HIGH (`11` impacted, `3` files, `1` flow); exact symbols `Metrics` LOW, `Run` CRITICAL, `runPhase` CRITICAL, and `newAnalyzeCommand` CRITICAL. All accepted command stderr artifacts are empty; mixed JSON-prefix plus guidance suffix is preserved where emitted. Fresh post-checkpoint test evidence: file-detail PASS/current; `fresh-command-test-file-impact.stdout.json` LOW (`0` impacted, `530` symbols, `7` linked tests); exact test-function impact LOW (`0` impacted, `1` affected file). |
| Source/test contract | COMPLETE FOR PRE-CAPTURE USE. `Metrics.Operations []OperationMetric` adds benchmark-only `boundary`, `name`, `duration`, and optional `denominators`. The existing focused test `TestAnalyzeCommandRunsGoPipelineAndWritesBenchmark` asserts unique/non-negative operations, phase-to-operation duration identity, required boundaries, and denominator keys. Owner resume authority records the focused test as already PASS; no standalone duration is inferred where no separate raw test artifact was preserved. Scoped `git diff --check` PASS. |
| Build-before-use | PASS. Exact full-build validation command above ran on HEAD `8df19258fbcb18e14841bf0dae036400aa9b22a3`, exit `0`. Start/end `2026-08-25T03:25:06.0193408+07:00` / `2026-08-25T03:36:46.1906114+07:00`; `700.1712706 s` is FULL-BUILD VALIDATION ELAPSED ONLY because this boundary includes build work plus maintenance `analyze . --force`. It is neither build elapsed nor a P1-A product/analyze benchmark result. Runtime version/hash and maintenance-analyze counts are recorded above. Only Vite chunk/dynamic-import warnings were written to stderr. |
| Build history | Two earlier invocations (`canonical-full-build.*.log`, `canonical-full-build-retry.*.log`) exited `1` because `E:\Anvien\anvien\bin\anvien.exe` was held open. A later post-holder invocation under `E:\Anvien\.tmp\child06a_p1a_post_holder_build_20260825_011700` exited `1` because nested vendor verification could not resolve `Get-FileHash`. Owner corrected the build path; current PASS artifacts are `canonical-full-build-current-source-20260825-032500.{preflight.json,stdout.log,stderr.log,result.json}` under the instrumentation root. These failures are validation history, not benchmark rows. |
| Runtime validity | CAPTURED: version `1.2.8`, `73,764,864` bytes, SHA-256 `62920CBF15921EF8A6D2FAC671776BAA3C312EFEA1BEED53B721C4AFF5E1B6C5`; exact accepted benchmark tuple ran once, exit `0`. The maintenance analyze embedded in full build remains excluded. Cross-capture overhead is not measured because no same-HEAD uninstrumented pair exists. |
| Terminal disposition | `CARRY_TO_FIRST_P2A_REFRESH`. Exact writer-owned `+349/-53` bytes remain instrumentation-owned through the first refreshed P2-A attempt and all P2 timing comparisons must retain the same instrumentation state. |

#### Instrumentation operation identity and denominators

| Boundary | Operation | Exact timed work | Denominators |
|----------|-----------|------------------|--------------|
| `analyzer_outer` | `analyze_setup` | `analyze.Run` entry through resolved-path/result setup immediately before analyzer total/orchestration begins | `analyzeRuns` |
| `analyzer_internal` | every existing phase name | the same synchronous body already measured by `runPhase`; operation duration is exactly the phase duration | `runs` |
| `analyzer_internal` | `graph_compact` | `result.Graph.Compact()` only | `inputNodes`, `inputRelationships` |
| `analyzer_internal` | `db_runner_resolve` | `resolveDBRunner` only | `runners` (`0` or `1`) |
| `analyzer_internal` | `db_runner_close` | resolved close callback only; emitted only when a closer exists | `runners` |
| `analyzer_internal` | `graph_snapshot` | `writeGraphSnapshot` only when enabled | `nodes`, `relationships` |
| `analyzer_internal` | `analyzer_orchestration` | exclusive accumulated `analyze.Run` time outside phases and the explicitly timed residual operations | `analyzeRuns` |
| `analyzer_internal` | `benchmark_write` | analyzer's benchmark write | `artifacts` |
| `cli_outer` | `cli_startup` | command construction/startup until `RunE` begins | `commands` |
| `cli_outer` | `cli_preparation` | `RunE` preparation through immediately before `analyze.Run` | `commands` |
| `cli_outer` | `memory_profile` | post-run memory-profile write | `profiles` (`0` or `1`) |
| `cli_outer` | `registry_meta` | analyze registry/meta recording | `repositories` |
| `cli_outer` | `ai_context` | generated AI-context work | `generatedFiles`, `baseSkills` |
| `cli_outer` | `file_projection` | file-projection construction | `files`, `nodes`, `relationships` |
| `cli_outer` | `output_publication` | existing JSON or text/publication path | `outputs` |
| `cli_outer` | `cpu_profile_completion` | CPU-profile stop/completion | `profiles` (`0` or `1`) |

Timers are top-level and non-overlapping within each boundary. Analyzer orchestration pauses around every phase and explicit analyzer operation. CLI operations execute sequentially around the single `analyze.Run` call. After all CLI timings exist, the CLI rewrites the same benchmark artifact so the final file carries the complete list; that final rewrite is instrumentation overhead and is not mislabeled as product work. Public CLI stdout/JSON and persistence calls remain on their existing paths.

#### Full-build validation artifacts

| Artifact | SHA-256 |
|----------|--------|
| `E:\Anvien\.tmp\child06a_p1a_instrumentation\canonical-full-build-current-source-20260825-032500.preflight.json` | `E73A82F1286B02AC48BCE1AA791BF7C48D8FA112BC36214DE81956490FF90544` |
| `E:\Anvien\.tmp\child06a_p1a_instrumentation\canonical-full-build-current-source-20260825-032500.stdout.log` | `F5AC4849A0E7F49032D0530335C862FBCB2CBE7DB782A0D193B1541436F0E309` |
| `E:\Anvien\.tmp\child06a_p1a_instrumentation\canonical-full-build-current-source-20260825-032500.stderr.log` | `BB87ED8391C6BFA1A011F1556CFCCB01AB339F7F818CB65BD0DF8839123859BB` |
| `E:\Anvien\.tmp\child06a_p1a_instrumentation\canonical-full-build-current-source-20260825-032500.result.json` | `6274812BD8DFFF5D541468EA005BE253665D28C7249A52BC2267A07B1CD7C05E` |

Handoff state: `TOP_LEVEL_QUEUE_READY -> PARENT_DRILLDOWN_PENDING`. Capture executor `01a03597-6afb-73e1-bee6-4306836cce01` is idle after `INSTRUMENTED_CAPTURE_COMPLETE`; old executor remains revoked/idle. P1-A is complete, `OP001 resolution` is active, and no child row, Architect, Coder, Supervisor, detect, commit, or optimization result exists.

### E1-P1A Instrumented Capture And Terminal Disposition

- Exact command: `E:\Anvien\anvien\bin\anvien.exe analyze E:\Anvien --force --json --progress --benchmark-json E:\Anvien\.tmp\child06a_p1a_instrumented_20260825_040400\benchmark.json --benchmark-label child06a-p1a-instrumented-20260825-040400`.
- PID/start/end/exit: `14956`; `2026-08-25T04:11:18.7266556+07:00`; `2026-08-25T04:21:37.1624765+07:00`; `0`. Exactly one launch, no retry, zero postflight matching analyze processes.
- Process CPU: `799.4062500 s` = user `785.2500000 s` + kernel `14.1562500 s`. No filesystem-I/O trace or exact peak working set was measured; maximum sampled working set was `2,617,962,496` bytes.
- Raw artifacts: `benchmark.json` `15,447` bytes / `0B089D679EF92F941AD53BE3A621D39E32F0F1D287BBB9EAB42B21D48D97F1EA`; `stdout.json` `9,130` / `3BB3AA3E6D7371F656B2359BD602FC46CC058DCB16593589B8D77A061F2D2601`; `stderr.log` `213` / `0FDCF15222C02C25771C98E17D9220C89FDCD0294137A6AE4D9B3AAFD4761ADC`.
- Published outputs: `.anvien/graph.json` `674,538,674` / `F6ADF32A4FF9A13A8E009BECA3EAA38DE3A02E5D2449EFDA132B6611D908F853`; `.anvien/lbug` `389,373,952` / `D72FBC8A10F7FB3AFF29630A1A7E28A4C5C391224111A22CC96AEB3357B187E7`; `.anvien/meta.json` `255` / `328ADA4B3BFDCF6EE6D5DE2748C364A2F3C7D852D3103FDC08DF315875EF7F59`.
- Arithmetic integrity: `30` globally unique operation names; all durations/denominators non-negative; only disabled `memory_profile` and `cpu_profile_completion` are zero with denominator `0`; all `15` phase names/order/durations match; phase sum `600.8917213 s`; analyzer non-phase sum `14.9739660 s`; analyzer-internal sum `615.8656873 s` equals `totalDuration`; analyzer outer `0.1697595 s`; CLI outer `2.2274029 s`; all operations `618.2628497 s`; uncovered process residual `0.1729712 s`.
- Comparability: same executable path/version, repository path, command options, forced regime, E-only topology, phase order, stderr, parsed `765`, failed `0`, projection unresolved `479`, COPY `18/93`, and DB fallback/failure/skip `0/0/0`. Exact mismatches versus immutable reference: binary `+14,848` bytes/different hash and different HEAD; scanned/documents `+4/+4`; parser bytes `+4,640`; nodes/relationships `+202/+237`; dependency edges `-6`; communities/memberships `-20/-4`; processes/steps `+4/+13`; graph/Ladybug/meta/stdout hashes differ.
- Arithmetic deltas versus immutable reference are process `+12.7030989 s`, analyzer `+13.3378062 s`, and phase sum `+13.6531237 s`; these are neither instrumentation overhead nor speed/regression evidence because no same-HEAD uninstrumented pair exists and workload/output counts differ.
- Main disposition: `CARRY_TO_FIRST_P2A_REFRESH`. This preserves one instrumentation state for all subsequent child/parent/end-to-end measurements and avoids discarding the only complete operation timing contract. The immutable reference remains preserved for provenance and final reporting without being mislabeled as the current like-for-like P2 baseline.

## E2 - P2 Evidence

Matching plan item(s): `P2-A`

### Attempt Evidence Contract

Every actual production attempt instantiates every mandatory ID below with one immutable attempt number; A001 additionally uses the Owner-specific `E2-P2A-A001CONSISTENCY1`, while `RESTORE1` and `SYSTEM1` are required only when their stated condition occurs.

| Attempt evidence kind | Required proof |
|-----------------------|----------------|
| `E2-P2A-AnnnCURRENT1` | exact child-attempt goal; active parent row; selected child row; accepted child/parent/total elapsed times; denominators; child unsuccessful streak; cause/owner/complete call path; preserved invariants; complete child-list and plan-checklist pointers |
| `E2-P2A-AnnnDRILL1` | measurement inside only the active parent boundary; complete current absolute elapsed-time ordered child list including smaller costs; exact benchmark-row/plan-checkbox cardinality; selected largest unchecked child; each child's denominator and validity |
| `E2-P2A-AnnnARCH1` | new Visible Architect identity/turn; exact parent row, complete child list, selected child elapsed-time owner/cause, current child/parent/total basis, attempt-local direction, allowed owners, expected gain, invariants, validation/resources, and rollback |
| `E2-P2A-AnnnPLAN1` | Planner refreshed the living P2-A attempt and exact implementation/test/build/measure/review/rollback steps before Coder |
| `E2-P2A-A001CONSISTENCY1` | A001-only direct-Owner sequence: Main four-ledger verification, independent Architect consistency result, Main acceptance/transition decision, revoked old Coder identity, sole active Coder identity/current pre-edit state, and explicit non-implementation/non-measurement/non-Supervisor/non-disposition/non-commit boundary |
| `E2-P2A-A001COMMIT1` | A001-only exact implementation commit hash/message, verified 12-path manifest, clean staged set, separation from unrelated dirty work, and explicit non-P2-closure/non-A002-authorization boundary |
| `E2-P2A-AnnnCOMMIT1` | accepted Supervisor-PASS/Main-KEEP progress checkpoint hash/parent/subject, exact scoped manifest, pre-commit detect/diff proof, clean staged set, and explicit non-P2-closure/non-checkbox/non-P3C boundary |
| `E2-P2A-AnnnIMPACT1` | fresh required graph/file-detail/impact for exact owner immediately before edit, including full HIGH/CRITICAL scope |
| `E2-P2A-AnnnSRC1` | exact production-first change and attempt-owned rollback boundary |
| `E2-P2A-AnnnBUILD1` | build-holder/process preflight plus canonical full-build `PASS` after production and authorized test edits, before test execution |
| `E2-P2A-AnnnTEST1` | authorized behavior-test edits made only after production is correct, followed by focused and package test execution only after full-build `PASS` |
| `E2-P2A-AnnnDIRECT1` | selected-child before/after with same denominator and benchmark rows |
| `E2-P2A-AnnnPARENT1` | active-parent before/after on the same accepted work and benchmark rows |
| `E2-P2A-AnnnE2E1` | total graph-generation before/after on same workload and benchmark rows |
| `E2-P2A-AnnnEQUIV1` | exact affected correctness/output/persistence/reader/lifecycle result |
| `E2-P2A-AnnnREVIEW1` | new Visible Supervisor identity/turn and independent accuracy/equivalence/invariant `PASS` or `REJECT` |
| `E2-P2A-AnnnDECISION1` | disposition and precise promotion/restoration/next-owner effect after Supervisor |
| `E2-P2A-AnnnRESTORE1` | required on rejected/non-retained candidate: exact owned rollback and last accepted state restored/retained |
| `E2-P2A-AnnnSYSTEM1` | only on third consecutive unsuccessful child attempt: exact parent/child rows, denominator, retained child/parent times, three attempts/reasons, terminal child `SYSTEM_CHARACTERISTIC`, and matching child checklist check |
| `E2-P2A-AnnnRANK1` | accepted-state child/parent/full-pipeline timing refresh, complete child/top-level list update, and exact next unchecked child/parent pointer |

Current attempt basis: A001/A002/A003 evidence and checkpoints remain complete history. A003 checkpoint is `b6bf45bce95323aa6b53b182edfea8628bd8b463`; accepted Cheapapp and Restaurant Manager values remain separate. WAL fix checkpoint `0f3a572331dd23d17688886fcbfebeb7d37ee35d` is complete. A004 is `SUPERVISOR_PASS / NO_KEEP / ROLLBACK_COMPLETE`; its values remain rejected-candidate evidence. D001 remains active/unchecked with streak `1`; parent and all children remain unchecked; D002-D017 remain queued. A005 source/build/test remains valid; frozen build invocation 1 produced no candidate because the copied 17-child overlay still passed the new private finalized bundle into two slice-denominator calls and `len`. State is `A005_FROZEN_BUILD_FAILED / OVERLAY_BUNDLE_ADAPTATION_RECOVERY_READY / MEASUREMENT_LOCKED` with no benchmark/streak/checklist effect.

### `E2-P2A-A001DRILL1` — Complete OP001 Resolution Child Measurement

- Exact launch identity: one Owner-authorized replacement `anvien analyze E:\Anvien --force --skip-git --json --progress --benchmark-json ...`, launch count `1`, exit `0`; process wall/CPU `653.475797100 / 830.109375000 s` and analyzer `630.160832200 s`.
- Parent boundary: `analyzer_internal/resolution`, `545.434182000 s`, denominator `{"runs":1}`.
- Complete child result: exactly `17` unique source-map IDs; child sum `545.364656400 s`; residual `0.069525600 s`; `2309` exclusive intervals; independently recomputed overlap count `0`.
- Arithmetic and ordering: every child duration equals the sum of its raw interval pairs; child sum and residual recompute exactly; benchmark parent row matches; ranking is non-increasing; every displayed percentage recomputes from the same-run parent.
- Raw numeric authority: `E:\Anvien\.tmp\child06a_p2a_op001_resolution_slot_segment05\process.json`, `benchmark.json`, `child_metrics.json`, `bottleneck_ranking.json`, `measurement_validation.json`, and `resolution.cpu.pprof`.
- Boundary authority: Segment 03 `source_map.json` and `validation.json` PASS; all `17` IDs, denominator expressions, source ranges, and call paths match the measured rows.
- Moving-repo correction: `GraphNodesEmitted 67882` versus `67783` and `GraphRelationshipsEmitted 106991` versus `106892` are cross-capture relative counts. Owner explicitly rejected treating absolute counts from a continuously changing repo as semantic-regression evidence. These deltas are not before/after speed evidence.
- Result: complete child queue is valid for the P2-A largest-first control surface. It does not prove a production speedup, candidate equivalence, `KEEP`, or final equivalence.

### `E2-P2A-A001CURRENT1` — Selected Child Current Basis

- Selected row: `B2-P2A-A001-D001 resolve_calls`, `243.794553800 s`, `44.697337%` of the same-run parent, denominator `calls=72976; files=765`.
- Owner: `E:\Anvien\internal\resolution\resolve.go:resolveCall:92-94`.
- Complete call path: `internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> for each ir in w.files -> for each call in ir.Calls -> resolution.resolveCall`.
- Cause evidence: source repeatedly invokes import-claim lookup during call resolution; `explicitImportNameClaimed` and `explicitImportCallState` scan `w.imports`, while `cleanPath` normalizes import paths inside those scans. The resolution-only CPU profile attributes cumulative samples inside `resolveCall` to `explicitImportNameClaimed` (`134.42 s`), `cleanPath` (`161.32 s`), and `explicitImportCallState` (`49.27 s`). CPU samples explain cause only and do not replace elapsed ranking.
- Current basis: child `243.794553800 s`; parent `545.434182000 s`; process/analyzer `653.475797100 / 630.160832200 s`; no-KEEP streak `0`; all `17` child checkboxes remain unchecked.
- Gate at capture time: fresh Architect A001 was next. `E2-P2A-A001ARCH1` and `E2-P2A-A001PLAN1` were then recorded; Main verification and the independent Owner-required consistency review later passed, and `E2-P2A-A001CONSISTENCY1` now records the Main-authorized Coder transition without changing this captured numeric basis.

### `E2-P2A-A001ARCH1` — A001 Architecture Decision

Status: recorded from the visible Architect result; this is architecture authority for Planner translation only, not Coder authorization, implementation evidence, a benchmark result, disposition, or Supervisor `PASS`.

#### Identity and acceptance boundary

- Visible Architect task: `Child 06A A001 Architect — resolve_…`.
- Thread: `01a036f5-c602-7422-b484-ac8e58a34eb9`.
- Report: `E:\Anvien\reports\system-architect\rp_system-architect_260825_121706_by_gpt-5_child06a_active_child_measurement_and_a001_direction.md`.
- Report status: `DECIDED / READY_FOR_PLANNER_TRANSLATION`.
- Report-only commit: `047f4d8ff5c850cf6fbe280627863537dd54552a`; parent `137c4e3071c2bc26f40e680e83f61dd8648e774e`.
- Main independently verified that HEAD/commit identity and that the commit adds only the one architecture report. That verification accepts the report as Planner input; it is not a per-attempt Supervisor review or implementation acceptance.

#### Exact current selection and cause

- Active parent: unchecked `B1-P1A-OP001 resolution`, same-run elapsed `545.434182000 s`.
- Complete ordered child queue remains exactly `B2-P2A-A001-D001..D017`, all unchecked: `resolve_calls`, `resolve_accesses`, `resolve_type_annotations`, `project_resolution_outcomes`, `emit_definition_nodes`, `finalize_resolution_outcomes`, `build_binding_occurrence_index`, `finalize_typescript_authority_results`, `emit_import_edges`, `emit_typescript_external_symbols`, `emit_method_dispatch_edges`, `assemble_resolution_result`, `binding_accumulator_dispose`, `emit_heritage_compatibility_edges`, `emit_unresolved_heritage_diagnostics`, `finalize_resolution_metadata`, and `runtime_setup`.
- Selected child: unchecked `B2-P2A-A001-D001 resolve_calls`, `243.794553800 s`, denominator `calls=72976; files=765`, no-KEEP streak `0`.
- Same-run process/analyzer basis: `653.475797100 / 630.160832200 s`; child sum/residual `545.364656400 / 0.069525600 s`; `2309` exclusive parent intervals; overlap count `0`.
- Cause: `resolveCall` repeatedly reaches `explicitImportNameClaimed` and `explicitImportCallState`; both perform global `w.imports` candidate scans and normalize visited import source paths with `cleanPath`. `repositoryReceiverClaimed` also reaches the name-claim scan. The current workload has `4887` imports and `72976` calls; one hypothetical full scan per call is `356633712` candidate visits, which is a structural scale calculation rather than an observed counter. Resolution-only CPU cumulative samples support this mechanism but overlap and cannot be added or substituted for elapsed ranking.

#### Measurement and work-count decision

- `child_total_elapsed` is the sole child-ranking, selection, before/after, and child-improvement acceptance value.
- Detailed attribution is active-child-only for D001; coarse `child_total_elapsed` remains for the complete 17-child list. D002-D017 stay queued/unchecked and are neither deep-instrumented nor opened by A001.
- The active-child exclusive contract is:

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

- `prepare_lookup_elapsed` owns path/input normalization, index/import scanning, and candidate preparation or lookup before a semantic decision.
- `resolution_decision_elapsed` owns binding, member, package, authority, target, proof, and resolution-outcome decision work not already owned by lookup.
- `graph_mutation_elapsed` owns node, relationship, reference, property, and metadata creation or merge work.
- `diagnostic_outcome_elapsed` owns resolution outcome and unresolved/diagnostic recording.
- `metrics_bookkeeping_elapsed` owns metrics/counter bookkeeping not already owned by another category.
- `residual_elapsed` owns child elapsed not attributed to the five explicit categories.
- Pending D001 comparability counters are files, calls, import-claim helper invocations, global-import candidates visited, import-path normalizations at claim lookup sites, index entries, unique lookup keys, matching-bucket candidates visited, resolution targets considered, references/relationships emitted, outcomes/diagnostics recorded, and affected graph node/relationship/property inventories.
- Counters prove comparable work but never replace elapsed wall time. Do not create per-call logs, traces, or timer artifacts for `72976` calls. Use the smallest local cumulative accumulators and merge once. Preserve identical instrumentation identity before/after. Invalid or mixed instrumentation cannot rank or prove gain. If recorded, a full-scan fallback counter must be zero for an accepted candidate.

#### Single selected production direction

- Create one separate run-scoped all-import claim index: `{canonical source file path, exact local import name} -> []original w.imports indices in existing order`.
- Build it during import construction after source-path canonicalization. Include resolved/unresolved, semantic/nonsemantic, and duplicate imports; append original indices in existing order.
- Replace only global candidate selection in `(*workspace).explicitImportNameClaimed` and the candidate-selection loop of `(*workspace).explicitImportCallState`.
- Preserve semantic filters, target checks, `allowed`/`importTargets`, target-nil behavior, lexical shadow behavior, and import-export resolution. Do not reuse/change resolved-only `importsByReceiver`, cache final answers, or iterate map keys for resolution order.
- Expected complexity is one `O(imports)` construction plus expected `O(1)` key lookup and `O(matching bucket size)` traversal. No timing threshold or elapsed saving is predeclared.

#### Exact owner and blast-radius boundary

- Allowed future production owners only:
  - `internal/resolution/indexes.go`: one private `workspace` field, its `buildWorkspace` initialization, and original-index population in `(*workspace).resolveImports`.
  - `internal/resolution/resolve.go`: body of `(*workspace).explicitImportNameClaimed`.
  - `internal/resolution/export_resolution.go`: candidate-selection loop of `(*workspace).explicitImportCallState`.
- Allowed future test edits only after production is complete: `internal/resolution/export_resolution_test.go` and `internal/resolution/resolution_test.go`.
- Every other production/helper/contract/test surface is inspect-only, run-only, or preserve-only. In particular preserve `resolveCall`, `repositoryReceiverClaimed`, `cleanPath`, `scopeFilePath`, import target/file resolution, graph emitters, persistence/readers, and every exported/shared contract. Needing another production owner/helper requires STOP and a fresh architecture decision.
- Blast radius is CRITICAL. Architect evidence records `resolveCall`, `workspace`, `buildWorkspace`, and `resolveImports` as CRITICAL owner/symbol boundaries. Helper-level `LOW/0` results do not provide impact assurance because unresolved method-call edges remain. A future Coder must run fresh graph/file-detail/exact-impact evidence immediately before edit.

#### Preserved invariants

- Exact resolution outcomes, confidence, proof kind, unresolved notes, and call-branch order.
- Exact query trimming and case-sensitive import-local-name semantics.
- Original import ordering, duplicates, first-match behavior, and traversal order.
- Current `cleanPath` behavior without lowercase, absolutization, or symlink resolution.
- Relative, absolute, language-specific, and Go/package import handling.
- Unresolved import claim and global-rescue blocking behavior.
- Nonsemantic import name claims and semantic call-state filtering.
- Graph node, relationship, ID, label, property, count, and ordered-output equivalence.
- Resolution outcome and index equivalence.
- Canonical in-memory graph, Graph JSON, Ladybug/native persistence/readback, and affected native/product readers.
- Deterministic replay/output; freshness/invalidation; failure propagation; transaction rollback; temporary-file flush/close/rename; publication visibility.
- Private run-scoped in-memory lifecycle with no new serialization, I/O, goroutine, global cache, public schema, or shared contract.

#### Future validation, resource, and rollback boundary

- Future Coder must freshly run `anvien --help`, `anvien analyze --force`, file-detail, and exact file/symbol impact before editing. The Architect graph refresh does not replace this pre-edit gate; HIGH/CRITICAL is a warning requiring scoped handling, not an edit ban.
- Production implementation precedes authorized test edits. Only after production behavior is correct may the two authorized test files be edited; do not execute tests yet.
- Perform build-lock/process preflight, then run exactly `pwsh -NoLogo -NoProfile -NonInteractive -ExecutionPolicy Bypass -File E:\Anvien\scripts\full-build.ps1` and require `PASS`.
- Only after the canonical full build passes, execute focused correctness tests and the `internal/resolution` package tests, including unchanged affected resolution/equivalence tests.
- Later and separately, remeasure with the same instrumentation and comparable work in order: D001 `resolve_calls`, parent `resolution`, then end-to-end graph generation/analyzer. Validate work counts and affected graph/outcome/persistence/reader equivalence.
- Memory must remain `O(imports + unique keys)` and store original indices rather than duplicate import objects.
- `KEEP` requires lower child elapsed, lower parent elapsed, lower end-to-end elapsed, and a later visible per-attempt Supervisor `PASS`. Build/tests/microbenchmark/CPU/profile/secondary counters cannot substitute.
- A001 rollback triggers on any build/test/outcome/order/graph/persistence/reader/determinism/failure/lifecycle change, unstable or non-linear memory, no valid D001 gain or no corresponding net parent/end-to-end benefit, or crossing the exact owner boundary.
- Attempt-owned bytes are only the private field/init/population hunks, the two claim-helper lookup/traversal hunks, and focused authorized test bytes. Rollback only those hunks; preserve user/protected bytes and carried P1 instrumentation `internal/analyze/analyze.go +117/-15`, `internal/cli/command.go +157/-38`, and `internal/cli/command_test.go +75/-0` as non-A001 work.

#### Residual NO EVIDENCE

- Exact helper-invocation counts, observed global-import visits, matching-bucket distribution, and allocation delta are not measured.
- CPU samples are cumulative/overlapping and cannot predict elapsed savings.
- Cheapapp and Restaurant Manager reports each remain measurement-only sources and do not themselves supply governance acceptance. Exact helper-invocation/deep-breakdown counters remain unmeasured. Independent Supervisor acceptance exists under `E2-P2A-A001REVIEW1`, Main `KEEP`/promotion exists separately under `E2-P2A-A001DECISION1`, and the A001 implementation commit exists under `E2-P2A-A001COMMIT1`; no average/combined baseline, A002 result, detect, or stage exists.

### `E2-P2A-A001PLAN1` — Exact Planner Refresh

Status: recorded / `PLAN_REFRESH_WRITTEN / OWNER_REQUIRED_ARCHITECT_REVIEW_PENDING`; review-ready, not Coder-ready.

- `plan.md` now binds A001 state, `E2-P2A-A001ARCH1`, the exact selected direction and production/test owners, active-child measurement/work-count/observer contract, preserved invariants, future production-first/build/remeasure/Supervisor flow, attempt-owned rollback, and the A001 Owner-direct review gate.
- `evidence.md` records this immutable `ARCH1` decision and this `PLAN1` materialization without claiming implementation or validation evidence.
- `benchmark.md` keeps all existing 17 child values/order unchanged, makes `child_total_elapsed` the controlling numeric field, and adds only pending active-child exclusive/work-count control rows with no invented breakdown measurement.
- `actual-status.md` moves the living pointer to `PLAN_REFRESH_WRITTEN / OWNER_REQUIRED_ARCHITECT_REVIEW_PENDING`, keeps all P2-A/parent/child checkboxes unchecked, points to `ARCH1/PLAN1`, and appends one refresh-log transition.
- Coder remains `LOCKED`. No production/test/source edit, build, test, runtime validation, measurement capture, benchmark result, candidate, `KEEP/REWORK/ROLLBACK/SYSTEM_CHARACTERISTIC`, per-attempt Supervisor result, detect, cleanup, staging, or commit was performed or fabricated by this Planner step.
- Owner-direct A001 sequencing: Main independently verifies the exact four diffs/artifacts, then sends the exact updated artifacts plus Planner handoff to Architect thread `01a036f5-c602-7422-b484-ac8e58a34eb9` for independent architecture-consistency review. This is the latest direct Owner requirement, not Architect command authority over Main, and it is not generalized here as a gate for every attempt.
- Architect reject returns exact correction to this Planner lane with no Coder. Architect pass returns to Main; Main alone decides the next governance transition.

### `E2-P2A-A001CONSISTENCY1` — Owner-Required Consistency PASS And Main Coder Transition

Status: recorded / `ARCHITECT_CONSISTENCY_PASS / CODER_ACTIVE_PRE_EDIT_GATE`.

- Main zero-trust verification of the four refreshed ledgers passed before the independent consistency review.
- Visible Architect thread `01a036f5-c602-7422-b484-ac8e58a34eb9`, turn `01a0377a-34a0-7523-b10c-820f84553fdb`, returned exactly `ARCHITECT_CONSISTENCY_REVIEW_PASS_READY_FOR_MAIN` after read-only commands and zero file changes.
- Main accepted that result as planning/architecture-gate evidence only and made the governance decision authorizing the A001 Coder transition. It is not implementation evidence, a benchmark/measurement result, disposition, Supervisor `PASS`, or commit authorization.
- Old Coder task `01a03375-f616-7651-ba81-99bad6536746` was revoked before ACK, tool use, or write; its active turn had `items=[]`.
- Sole active visible Coder task `01a03786-cf51-7221-83fa-1c3fbf13cfa1`, turn `01a03786-d11b-7493-9894-17d2acc91651`, ACKed `HIỂU`, completed FULL RAW startup, and is in the fresh pre-edit Anvien gate. No production/test edit existed at this state-sync snapshot.
- Current authorization is limited to the exact pre-edit gate followed by the refreshed production-first/test-edit/build/test-validation sequence. Measurement, disposition, per-attempt Supervisor, and commit remain unauthorized; no Architect re-review is required for this truthful state/order sync.

### `E2-P2A-A001IMPACT1` — Fresh Pre-Edit Anvien Impact

Status: recorded / complete for the exact A001 candidate owner boundary.

- Fresh `anvien analyze --force` exited `0`: `2211` scanned, `765` parsed, `0` failed.
- File risk is HIGH for all three production files: `internal/resolution/indexes.go`, `internal/resolution/resolve.go`, and `internal/resolution/export_resolution.go`.
- Exact symbol impact: `workspace` CRITICAL (`145` symbols / `30` files / `71` processes / `7` modules); `buildWorkspace` CRITICAL (`62 / 26 / 23 / 7`); `resolveImports` CRITICAL (`37 / 19 / 17 / 3`).
- `explicitImportNameClaimed` and `explicitImportCallState` returned `LOW/0`; this is explicitly not impact assurance because unresolved method-call edges remain. The candidate stays inside the exact ARCH1 owner boundary.

### `E2-P2A-A001SRC1` — Exact Production-First Candidate

Status: recorded / uncommitted candidate at HEAD `e087f77032ff871b766c6af769eaa7c1aece9c73`; staged set empty; scoped diff-check PASS.

- Exact A001 production/test files and current SHA-256:
  - `internal/resolution/indexes.go` — `AA2118AAF65349778E7EAD288FBBDBE5150C8143ABF4250AE484A786D064D35D`.
  - `internal/resolution/resolve.go` — `8CEEDBA1883314EE8883320D3647C25DEF6F19F043D57881A893FBA73BA210D9`.
  - `internal/resolution/export_resolution.go` — `CA36FD64D1EDEE25DFBCC18AF1A09FE30CE0D748C8714F25057D78C5468B58E7`.
  - `internal/resolution/export_resolution_test.go` — `F7A479D4BE87FBD0B21E77FAC00EBF3D101B7A43FAEE579ACD17D0D072BB75AE`.
  - `internal/resolution/resolution_test.go` — `E229FE6F6E10AC47CB1ACC4BCAE7E8B31A0EDF1596DC33F877C613A99D890CF1`.
- Exact five-file candidate diff is `+174/-10`.
- `indexes.go` adds private `importClaimsByReceiver`, initializes it in `buildWorkspace`, and appends original `w.imports` indices for every nonempty canonical source-path/exact-local-name claim. Resolved/unresolved, semantic/nonsemantic, and duplicate imports remain included; resolved-only `importsByReceiver` stays nested and semantically unchanged.
- `resolve.go` replaces the `explicitImportNameClaimed` whole-import scan with exact bucket presence. `export_resolution.go` replaces only `explicitImportCallState` candidate acquisition with indexed bucket traversal; semantic, `TargetDef`, `resolveImportExport`, `allowed`, `importTargets`, target-nil, and shadow logic remain unchanged.
- The index remains private, run-scoped, and in-memory with no serialization, I/O, goroutine, global cache, or public contract. Memory remains `O(imports + unique keys)` and stores original indices rather than duplicate import objects.
- Tests were edited only after production inspection and only in the two authorized files. Attempt-owned rollback remains exactly these A001 hunks; carried P1 instrumentation is untouched.

### `E2-P2A-A001BUILD1` — Canonical Full-Build Validation

Status: `PASS` / exit `0`; this is validation evidence only, never product-performance evidence.

- Build-holder preflight found the analyze lock free and no repo-local runtime or repo-scoped Go/Node build holder; nothing was terminated.
- Exact command: `pwsh -NoLogo -NoProfile -NonInteractive -ExecutionPolicy Bypass -File E:\Anvien\scripts\full-build.ps1`.
- Start/end: `2026-08-25T13:40:12.1770424+07:00` / `2026-08-25T13:46:27.4898463+07:00`; exit `0`.
- Maintenance analyze completed with `2211` scanned / `765` parsed / `0` failed.
- Runtime: version `1.2.8`, `73,764,864` bytes, SHA-256 `C03380CA8E47C48361C8A88774378E01D191726E5302A52585F3BD5F7024D6E8`.
- The build boundary occurred before all test execution. Its elapsed window is validation only and cannot become child, parent, end-to-end, or speed evidence.

### `E2-P2A-A001TEST1` — Focused PASS And Pre-Existing Package Golden FAIL

Status: focused `PASS`; full package truthfully `FAIL` on one proven pre-existing preserve-only golden. The package is not labeled PASS.

- Focused command: `go test ./internal/resolution -run '^(TestBuildWorkspaceIndexesAllImportClaimsByCanonicalSourceAndExactName|TestExplicitImportCallStatePreservesIndexedCandidateSemantics)$' -count=1 -v`; exit `0`. Main independently reran it and both tests passed.
- Full-package command: `go test ./internal/resolution -count=1`; exit `1`, solely `TestProofBasedCallAccessGoldenCorpus` at `proof_accuracy_golden_test.go:166`.
- Expected callback metadata classification/actionability is `in_repo_unresolved/analyzer_gap`; actual is `unclassified/review`. An isolated current-candidate rerun fails identically.
- A repo-local Go overlay replaced all five authorized A001 files with their exact HEAD blobs; all `5/5` blob IDs matched HEAD, and the same targeted golden failed with the identical Diagnostic payload. This proves the failure predates A001 and is outside the editable/preserve-only surface.
- Under AGENTS rule 8, this is a pre-existing unrelated stale validation failure. It is neither A001 failure nor A001 validation proof; `proof_accuracy_golden_test.go` remains untouched and preserve-only.
- The candidate remains unchanged. The two independent target measurement reports are now recorded below; `TEST1` itself remains validation-only and supplies no performance/equivalence acceptance, Supervisor result, or disposition.

### Cheapapp Measurement Report Evidence — Measurement Only

- Report: `E:\Anvien\reports\Investigation\rp_child06a_a001_cheapapp_benchmark.md`; status `MEASUREMENT_COMPLETE`; expected lane `01a033a5-4700-7272-a28d-de3a71f58135`; target `E:\cheapapp.org`.
- Baseline executable: `E:\Anvien\.tmp\child06a_a001_cheapapp_benchmark\baseline\anvien-pre-a001.exe`, version `1.2.8`, `73814016` bytes, SHA-256 `5C0F4C6BC13204ABFCBBD05C53499F7C141C819875BB29D006EE130A2A04C53F`; exact command uses `analyze E:\cheapapp.org --force --skip-git --json --progress` and the baseline benchmark path/label recorded in the report.
- Optimized executable: `E:\Anvien\.tmp\child06a_a001_cheapapp_benchmark\optimized\anvien-current-a001.exe`, version `1.2.8`, `73814016` bytes, SHA-256 `DA34A01158811833F3DE780E8308407FEC3B97CD903E446FD8BA2ED9D8164F51`; exact command uses the same target/options and the optimized benchmark path/label recorded in the report.
- Complete numeric inventory: product boundary plus all `30` top-level operation rows and all `17` resolution-child rows, with identical denominators per baseline/optimized pair. D001 is `209.823705100 -> 25.045225300 s`, `calls=27890`, `files=887`, delta `-184.778479800 s (-88.063682%)`; parent is `733.384225100 -> 184.481061700 s`; process wall is `890.314783200 -> 279.105934600 s`; analyzer is `851.802490600 -> 274.474620900 s`.
- Conservation: child sum `733.255678600 -> 184.422170400 s`; parent residual `0.128546500 -> 0.058891300 s`; `2675` exclusive intervals and `0` overlaps in both runs.
- Equivalence facts: canonical Graph JSON, stdout, stderr, workload/graph counts, resolution semantic metrics, and all denominators match exactly. Ladybug and meta bytes/hashes differ; the report makes no acceptance inference from those differences.
- Raw roots: `E:\Anvien\.tmp\child06a_a001_cheapapp_benchmark\owner-final-run\baseline` and `E:\Anvien\.tmp\child06a_a001_cheapapp_benchmark\owner-final-run\optimized`; comparison appendix `E:\Anvien\.tmp\child06a_a001_cheapapp_benchmark\owner-final-run\comparison.json`.
- This report evidence is measurement-only and did not itself promote the candidate. Independent acceptance and Main promotion occur only under `E2-P2A-A001REVIEW1/DECISION1`; no Restaurant capture was read or run by the Cheapapp lane.

### Restaurant Manager Measurement Report Evidence — Measurement Only

- Report: `E:\Anvien\reports\Investigation\rp_child06a_a001_restaurant_manager_benchmark.md`; status `MEASUREMENT_COMPLETE`; lane `01a033a5-4ec3-7e42-8946-0ab9172f6088`; target `E:\Restaurant_manager`; target HEAD `605c0bda99491789e7f07628ec4b3d39a3ae1c67`.
- Both runs use the exact identical exclusion `electron/renderer/src/api/userApi.ts`. Baseline executable SHA-256 is `F25C9C17E232834F0ED5FB8E7F3935BE61F5F457E181F105DB76EFA90CA20642`; optimized executable SHA-256 is `DA34A01158811833F3DE780E8308407FEC3B97CD903E446FD8BA2ED9D8164F51`; Go identity is `go version go1.26.3 windows/amd64`.
- Baseline command: `& "E:\Anvien\.tmp\child06a_a001_restaurant_manager_benchmark\fresh-supported-exclude-userapi-20260825\runs\baseline\anvien-restaurant-manager-pre-a001.exe" analyze E:\Restaurant_manager --force --json --progress --exclude electron/renderer/src/api/userApi.ts --benchmark-json E:\Anvien\.tmp\child06a_a001_restaurant_manager_benchmark\fresh-supported-exclude-userapi-20260825\runs\baseline\benchmark.json --benchmark-label child06a-a001-restaurant-manager-pre-a001`.
- Optimized command: `& "E:\Anvien\.tmp\child06a_a001_restaurant_manager_benchmark\fresh-supported-exclude-userapi-20260825\runs\optimized\anvien-restaurant-manager-current-a001.exe" analyze E:\Restaurant_manager --force --json --progress --exclude electron/renderer/src/api/userApi.ts --benchmark-json E:\Anvien\.tmp\child06a_a001_restaurant_manager_benchmark\fresh-supported-exclude-userapi-20260825\runs\optimized\benchmark.json --benchmark-label child06a-a001-restaurant-manager-current-a001`.
- Baseline and optimized `process.json` each record launch count `1` and exit code `0`. Complete numeric inventory is one product boundary, all `30` top-level operation rows, and all `17` resolution-child rows with no denominator mismatch. D001 is `501.638742800 -> 40.769294200 s`, `calls=86030`, `files=1234`, delta `-460.869448600 s (-91.872778%)`, rank `1 -> 2`; parent is `1090.085959900 -> 136.436879300 s`; process wall is `1178.391336900 -> 218.680628900 s`; analyzer is `1174.442838300 -> 215.972455200 s`.
- Workload/output equivalence facts: scanned `1556`, parsed `1234`, failed `0`, graph nodes/relationships `137229 / 190644`, output counts, target-local Graph JSON SHA-256, and stdout SHA-256 match exactly; `denominatorMismatchRows` is empty. The identical target-local graph SHA-256 is `09708B038AD63B659A95BAC34304A752F2C6350952B91CF435EBD5C5F9D94B7C`.
- Raw root: `E:\Anvien\.tmp\child06a_a001_restaurant_manager_benchmark\fresh-supported-exclude-userapi-20260825`; baseline/optimized artifacts are under its `runs` directory; machine-readable comparison is `E:\Anvien\.tmp\child06a_a001_restaurant_manager_benchmark\fresh-supported-exclude-userapi-20260825\comparison.json`.
- This report evidence is measurement-only and did not itself promote the candidate. Independent acceptance and Main promotion occur only under `E2-P2A-A001REVIEW1/DECISION1`; it is never averaged or combined with Cheapapp.

### A001 Owner-Approved Two-Target Benchmark Plan

The two independent reports below are now both recorded as the only repo-specific numeric inputs replacing the superseded `E:\Anvien` control surface. Each lane compares pre-A001 baseline with current optimized A001 on its one assigned target and reports D001 `resolve_calls` elapsed/denominator, parent `resolution` elapsed, end-to-end analyze elapsed, graph/output equivalence, and exact executable/command/target identity. Targets are never averaged or combined, and neither report promotes the candidate.

| Target | Existing visible measurement lane | Exact required report | Current value state |
|--------|-----------------------------------|-----------------------|---------------------|
| `E:\cheapapp.org` | `01a033a5-4700-7272-a28d-de3a71f58135` | `E:\Anvien\reports\Investigation\rp_child06a_a001_cheapapp_benchmark.md` | `MEASUREMENT_COMPLETE`; optimized values later promoted for this repo only by `E2-P2A-A001DECISION1` |
| `E:\Restaurant_manager` | `01a033a5-4ec3-7e42-8946-0ab9172f6088` | `E:\Anvien\reports\Investigation\rp_child06a_a001_restaurant_manager_benchmark.md` | `MEASUREMENT_COMPLETE`; optimized values later promoted separately by `E2-P2A-A001DECISION1`, with identical exclusion in both runs |

The two report blocks were consumed by the independent A001 review and Main disposition recorded below. Planner did not assign, command, or coordinate either lane. Each target remains separate; A001 optimized values are promoted only by `E2-P2A-A001DECISION1`, not by either measurement report alone.

### `E2-P2A-A001REVIEW1` — Independent Per-Attempt Supervisor Acceptance

Status: `SUPERVISOR_A001_PASS`; measurement/equivalence acceptance only. This evidence does not itself make `KEEP`, promotion, staging, commit, or A002 authorization decisions.

- Supervisor report: `E:\Anvien\reports\Supervisor\rp_supervisor_260825_172308_by_gpt-5_a001_measurement_equivalence_acceptance.md`.
- Review time/reviewer: `2026-08-25 17:23:08 +07:00` / `gpt-5`. No separate Supervisor task or turn identity is stated in the report, so none is fabricated here.
- Exact verdict: `SUPERVISOR_A001_PASS`.
- Candidate boundary at review: HEAD `72f7da1dc41d4d0b65e6a5477fa12640b51a7454`, which descends from recorded source-evidence HEAD `e087f77032ff871b766c6af769eaa7c1aece9c73`; the intervening commit changes none of the five A001 paths. The exact five-file candidate remains `+174/-10`, scoped diff-check clear, staged set empty, and all five SHA-256 values match `E2-P2A-A001SRC1` plus measurement preflight identities.
- Cheapapp accepted review facts: D001 `209.823705100 -> 25.045225300 s`, parent `733.384225100 -> 184.481061700 s`, process E2E `890.314783200 -> 279.105934600 s`, denominator `calls=27890; files=887`, `30/30` operations, `17/17` children, zero denominator mismatches, canonical Graph JSON/output/workload equivalence.
- Restaurant Manager accepted review facts: D001 `501.638742800 -> 40.769294200 s`, parent `1090.085959900 -> 136.436879300 s`, process E2E `1178.391336900 -> 218.680628900 s`, denominator `calls=86030; files=1234`, `30/30` operations, `17/17` children, zero denominator mismatches, canonical Graph JSON/output/workload equivalence.
- Same-work/invariant closure: all operation/child denominators, resolution semantic metrics, workload/output counts, canonical Graph JSON, and public stdout/stderr match within each target pair. Cheapapp Ladybug/meta bytes differ but are explicitly non-blocking for this acceptance boundary because canonical graph, DB readback counts, public output, and affected semantics match and A001 does not edit persistence/metadata code.
- The known package golden failure remains proven pre-existing/preserve-only. The Supervisor did not rerun benchmarks, full build, or tests; it proved current candidate-byte identity and consumed recorded evidence without reopening Architect, Planner, Coder, impact, staging, or commit gates.
- Exact authority boundary: the Supervisor accepts the evidence as eligible input only. Main alone owns `KEEP`, disposition, baseline promotion, checklist transition, staging, and commit.

### `E2-P2A-A001DECISION1` — Official Main KEEP And Repo-Specific Promotion

Status: `A001 KEEP / A002 ARCHITECT_PENDING`.

- Main Orchestration accepted `E2-P2A-A001REVIEW1` and officially decided `A001 = KEEP` because D001 `resolve_calls`, parent `resolution`, and process E2E all improved on both independent targets while same-work canonical graph/output equivalence passed and candidate bytes matched recorded source/build/test evidence.
- Promotion remains repo-specific and never averaged or combined:
  - `E:\cheapapp.org` accepted current A001 baseline: D001 `25.045225300 s` with `calls=27890; files=887`; parent `184.481061700 s`; process/analyzer `279.105934600 / 274.474620900 s`.
  - `E:\Restaurant_manager` accepted current A001 baseline: D001 `40.769294200 s` with `calls=86030; files=1234`; parent `136.436879300 s`; process/analyzer `218.680628900 / 215.972455200 s`.
- The pre-A001 values remain immutable comparison history in each target block; the optimized values above are now the accepted current baseline for that same target only.
- Hierarchy/checklist effect: P2-A, unchecked parent `B1-P1A-OP001`, and unchecked active child `B2-P2A-A001-D001` remain open. D001's consecutive no-KEEP streak resets/remains `0`. D002-D017 remain queued, unchecked, and unopened. No parent or child is terminalized.
- Next campaign cursor: `A002 / ARCHITECT_PENDING` on the same active D001. A001 architecture authority expired at disposition; Main must open a fresh visible A002 Architect before a new Planner refresh or Coder. Coder is not reopened by this decision.
- At A001 disposition time the accepted production/test bytes remained unchanged and uncommitted. Their later exact implementation commit is appended separately under `E2-P2A-A001COMMIT1`; this does not fabricate or satisfy final Child 06A `E3-P3C-COMMIT1` closure.

### `E2-P2A-A001COMMIT1` — A001 Implementation Commit Complete

Status: `A001_COMMIT_COMPLETE`; fact sync only. This commit records accepted A001 and does not close P2-A/Child 06A, check any parent/child, authorize A002 Coder, or satisfy final `E3-P3C-COMMIT1` closure.

- Exact commit: `17a1f3af37dcb61f9d389345822b6470a8f772cc`.
- Exact subject: `perf(resolution): index import claims for call resolution`.
- Parent: `72f7da1dc41d4d0b65e6a5477fa12640b51a7454`.
- Verified manifest contains exactly `12` paths: the five A001 source/test files, the four Child 06A standard ledgers, the two target benchmark reports, and the one A001 Supervisor report.
- Exact five source/test paths: `internal/resolution/indexes.go`, `internal/resolution/resolve.go`, `internal/resolution/export_resolution.go`, `internal/resolution/export_resolution_test.go`, and `internal/resolution/resolution_test.go`.
- Exact report paths: `reports/Investigation/rp_child06a_a001_cheapapp_benchmark.md`, `reports/Investigation/rp_child06a_a001_restaurant_manager_benchmark.md`, and `reports/Supervisor/rp_supervisor_260825_172308_by_gpt-5_a001_measurement_equivalence_acceptance.md`.
- Post-implementation Coder completion/commit evidence pointer: `E:\Anvien\reports\coder\rp_coder_260825_175710_by_gpt-5_child06a_a001_import_claim_index.md`. Full report verification confirms the five-file implementation, build-before-tests order, pre-existing golden failure, both separate target measurements, `SUPERVISOR_A001_PASS`, Main `KEEP`, detect-changes, and exact implementation commit. The report was created after `17a1f3af37dcb61f9d389345822b6470a8f772cc`, is not in its 12-path manifest, and is committed separately at `18b5063d236f9f2567fea90e48eca8f1501bd1eb` with subject `docs(report): record Child 06A A001 coder handoff`; that report commit contains exactly one path, `reports/coder/rp_coder_260825_175710_by_gpt-5_child06a_a001_import_claim_index.md`.
- Staged set is empty after the commit. Unrelated pre-existing dirty/untracked files remain outside the manifest, unstaged, and untouched by this fact sync.
- State effect: preserve A001 `KEEP`, separate accepted Cheapapp/Restaurant baselines, active unchecked D001, streak `0`, queued/unopened D002-D017, and exact current cursor `A002 / ARCHITECT_PENDING`. No A002 evidence or Coder authorization is created.

### `E2-P2A-A002ARCH1` — A002 Run-Scoped Diagnostic Appender Architecture

Status: recorded from the fresh Architect report / `ARCHITECT_A002_READY_FOR_PLANNER`. This is attempt-local architecture input for Planner translation only; it is not Coder authorization, implementation/build/test evidence, an elapsed-speed claim, measurement result, Supervisor result, disposition, staging, or implementation commit.

#### Identity and current selection

- Report: `E:\Anvien\reports\system-architect\rp_system-architect_260825_184751_by_gpt-5_child06a_a002_run_scoped_diagnostic_appender.md`.
- Report-only commit: `dafcf7a25e254ef8db09d28eca3575d5225c553e`; parent/current source-and-index authority `56090ac3c396b11b7a7aa0e9240ec5acbe562f01`; subject `docs(architecture): define Child 06A A002 diagnostic appender`; manifest exactly the one architecture report.
- Verdict/handoff: `ARCHITECT_A002_READY_FOR_PLANNER`.
- Architect-recorded fresh graph state: `2216` files scanned, `765` parsed, `0` failures. This report evidence does not replace the future Coder's fresh pre-edit Anvien gate.
- Active parent remains unchecked `B1-P1A-OP001 resolution`; active child remains unchecked `B2-P2A-A001-D001 resolve_calls`; D001 no-KEEP streak remains `0`; D002-D017 remain queued, unchecked, and unopened.
- Accepted A001 baselines remain separate: Cheapapp D001/parent/process/analyzer `25.045225300 / 184.481061700 / 279.105934600 / 274.474620900 s`, denominator `calls=27890; files=887`, diagnostics/outcomes `57683 / 86742`; Restaurant D001/parent/process/analyzer `40.769294200 / 136.436879300 / 218.680628900 / 215.972455200 s`, denominator `calls=86030; files=1234`, diagnostics/outcomes `129009 / 186251`.

#### Residual cause and causal CPU evidence

- `internal/graphhealth/diagnostics.go`: `AppendDiagnosticToNode` normalizes the incoming diagnostic; `appendDiagnosticsFromProperties` obtains all existing diagnostics and normalizes the incoming diagnostic again; stored `[]Diagnostic` enters `normalizeDiagnosticSlice`, which allocates a replacement slice and re-normalizes every prior entry; each structured normalization enters `decodeStructuredResolutionOutcome`, which unmarshals the same JSON `Note` into an envelope and then the outcome.
- Production resolution has exactly two current `AppendDiagnosticToNode` call sites: `internal/resolution/emit.go:218` and `internal/resolution/outcome.go:392`, owned by `(*emitter).emitUnresolvedReference` and `(*emitter).emitTypeScriptOutcomeDiagnostic`.
- Inference, not an operation counter: because each successful append writes `[]Diagnostic` back, repeated prior-entry normalization has the per-node triangular shape `0 + 1 + ... + (k-1)`. Profiles prove the stack is material but do not measure the exact repeated decode count.
- Profiles were filtered to stacks containing `resolution.resolveCall`; values are cumulative CPU samples, overlap, are not additive, are not elapsed-speed claims, and are never averaged: Cheapapp `AppendDiagnosticToNode 20.30 s` / `normalizeDiagnosticMetadata 20.09 s`; Restaurant Manager `30.18 s / 29.56 s`.
- Raw identities: Cheapapp benchmark/profile SHA-256 `BD598F4D5BCC52E7FB248988889431E7DF0CEC6788D583A4BF3637DE00CDC3BA` / `81606B16916225E5FCDAABEA7BE542F680626BF6093D1B026342608026289484`; Restaurant `72B9B441A84ABBA427F844C728B85DF1842C7A25187B785D5F7344CAFE9C1666` / `3A4C1BA5C535124DE3A5E898FA316C612AA7697523CF6B8819275DF468A1B961`. Roots remain `E:\Anvien\.tmp\child06a_a001_cheapapp_benchmark\owner-final-run\optimized` and `E:\Anvien\.tmp\child06a_a001_restaurant_manager_benchmark\fresh-supported-exclude-userapi-20260825\runs\optimized`.

#### One selected direction and exact owners

- Add one private, reversible, write-through diagnostic appender owned by the existing per-run `emitter`. During one sequential `ResolveBoundInto`, normalize each touched node's existing supported diagnostics once on first successful touch, normalize each incoming diagnostic once, merge normalized inputs, and keep writing the same `[]graphhealth.Diagnostic` representation through `Graph.AddNode` after every append.
- Only `internal/graphhealth/diagnostics.go`: run-scoped appender plus one private normalized-input merge helper; cache only `map[nodeID][]Diagnostic` slice headers/state; first successful touch reads the current node and normalizes its existing supported representation once; every successful append re-reads the node, preserves all other properties, writes `DiagnosticPropertyKey` as `[]Diagnostic`, and calls `Graph.AddNode`. `AppendDiagnosticToNode` remains available for all other callers with current normalization/observable semantics; helper refactoring is allowed only to avoid duplicated rules.
- Only `internal/resolution/emit.go`: add appender field to `emitter`, initialize it in `newEmitter` from the same graph, and route `emitUnresolvedReference` through it.
- Only `internal/resolution/outcome.go`: route `emitTypeScriptOutcomeDiagnostic` through the same appender.
- Only after production is correct, add focused test bytes in one new file: `internal/graphhealth/diagnostics_test.go`.
- No second optimization: do not redesign diagnostic policy, remove the two unmarshals inside one structured decode, change bucket lookup/sorting, or open D002-D017.

#### Blast radius and preserved invariants

- Architect file risk: all three production files HIGH. `diagnostics.go`: `154` symbols, inbound/outbound `50/70`, flows/tests `1/19`; `emit.go`: `154`, `81/231`, `13/24`; `outcome.go`: `119`, `87/124`, `11/23`.
- Architect symbol risk: `AppendDiagnosticToNode` CRITICAL (`7/2` impacted/direct, `1/24` modules/processes); `emitter` CRITICAL (`28/25`, `2/49`); `newEmitter` CRITICAL (`6/1`, `4/32`); `emitUnresolvedReference` CRITICAL (`7/4`, `2/36`); `emitTypeScriptOutcomeDiagnostic` LOW/0, which does not override HIGH file risk.
- Preserve false-return behavior for nil graph, blank node ID, blank diagnostic kind, and absent source node; preserve `SourceNodeID` and `Count` defaults.
- Preserve structured fail-closed and legacy classification/actionability policy; bucket identity; merged counts; empty-target fill; earliest source line; stable order; and write-through visibility after every call.
- Preserve `[]Diagnostic` graph property, Graph JSON, persistence/readback, health projections, public output, and deterministic replay.
- Preserve every resolution outcome, encoded `Note`, authority/proof/source-site field, metric, node, relationship, ID, label, property, branch, and order.
- Normalize pre-existing supported diagnostic representations once on first touch exactly as the legacy function would. Cache validity is only under existing sequential `ResolveBoundInto`; do not cache whole nodes or silently overwrite an unproven concurrent/out-of-band writer.
- Let `T` be source nodes touched by resolution diagnostics. Retained state is `O(T)` map entries/slice headers; cache and graph point to the same normalized slice and retain no second diagnostic-object copy. The first-touch replacement allocation may occur once, and the old property can then be collected.
- Map lifecycle: created in `newEmitter`, no serialization/I/O/goroutine/lock/global cache, unreachable when `ResolveBoundInto` returns; diagnostic slices remain owned by the returned graph; every append is write-through, so no flush/finalizer.

#### Required validation, measurement, rollback, and STOP

1. Future Coder repeats fresh `anvien --help`, `anvien analyze --force`, file-detail for all three production owners, and exact symbol impacts; Architect evidence does not replace the gate.
2. Edit the three production owners first. Only after production is correct add `internal/graphhealth/diagnostics_test.go`, comparing legacy and repeated appender writes for first-touch decoded input, multiple unique entries, duplicate-bucket merging, stable order, policy fields, per-append write-through, and absent-node failure.
3. Perform build-holder/lock preflight, then require canonical full-build `PASS` before any test execution: `pwsh -NoLogo -NoProfile -NonInteractive -ExecutionPolicy Bypass -File E:\Anvien\scripts\full-build.ps1`.
4. Post-build focused tests include the new appender-equivalence test, `TestP6DStructuredOutcomeProjectionPreservesGraphJSONAndHealthParity`, `TestResolveAttachesSourceBackedUnresolvedDiagnostics`, `TestP6C3FinalOutcomeStatusPrecedenceAndCarriageMatrix`, and `TestP6BResolverCarriesCatalogValidationFailuresPerSourceSite`; then `go test ./internal/graphhealth -count=1` and `go test ./internal/resolution -count=1`. The known preserve-only `TestProofBasedCallAccessGoldenCorpus` baseline failure needs exact unchanged-baseline proof, is not PASS, and any new failure blocks A002.
5. Later use each accepted A001 artifact as its repo's own `before`; build one identifiable A002 candidate from accepted A001 plus only the authorized diff, preserving identical 17-child instrumentation and runtime/native/build contract. Cheapapp retains `E:\cheapapp.org --force --skip-git --json --progress`; Restaurant retains `E:\Restaurant_manager --force --json --progress` and exactly one `--exclude electron/renderer/src/api/userApi.ts`; only executable identity, benchmark path, and label may differ.
6. Each target remains separate with graph under target-local `.anvien`; record target HEAD/status/workload plus process/executable/artifact hashes, D001, parent, analyzer total, process wall, calls/files, all 30 operation rows, all 17 child rows, scanned/parsed/failed, graph nodes/relationships, DB readback counts, resolution semantic metrics, diagnostics/outcomes, canonical Graph JSON, public stdout, `startAllocBytes`, `endAllocBytes`, `maxObservedSys`, and `O(T)` no-duplicate proof. Ladybug/meta byte equality alone is not the semantic gate; build/test/microbenchmark/profile duration cannot substitute.
7. No numeric speed threshold is predicted. `KEEP` requires valid same-work D001 improvement plus corresponding net parent/process benefit on both repositories independently, all correctness/resource gates, and later per-attempt Supervisor `PASS`.
8. Roll back only A002 appender/helper, `emitter` field/init, two route substitutions, and new focused test file if any build, correctness, equivalence, determinism, lifecycle, memory, D001, parent, or process gate fails. Preserve A001/P1/protected work.
9. STOP and return to Main for fresh architecture if any additional production owner is required, including `internal/graph/types.go`, `internal/graphhealth/policy.go`/`Diagnostic`, `internal/resolution/resolve.go`, outcome/metric schemas, persistence/readers, public contracts, global/concurrent caching, or timing instrumentation. Any new residual belongs to a later attempt.

#### Residual NO EVIDENCE

- No A002 source/test edit, candidate binary, build, test, wall-time measurement, allocation result, graph/output equivalence result, Supervisor result, disposition, detect, stage, or implementation commit exists.
- CPU samples do not predict elapsed savings; exact repeated-decode counts per node are not measured; no candidate `startAllocBytes`, `endAllocBytes`, or `maxObservedSys` exists yet.

### `E2-P2A-A002PLAN1` — Exact A002 Planner Refresh

Status: recorded / `A002 PLAN_READY_FOR_MAIN_VERIFY / CODER_LOCKED`.

- `plan.md` binds the A002 selected direction, exact production/test owner boundary, preserved diagnostic/graph/output/lifecycle/resource invariants, production-first/full-build-before-tests order, exact focused/package validation, independent two-target measurement contract, scoped rollback, and mandatory STOP.
- `evidence.md` records immutable `E2-P2A-A002ARCH1` and this materialization without fabricating implementation or validation proof.
- `benchmark.md` preserves every A001 numeric row/order and both accepted repo-specific baselines, records the two residual CPU samples separately as causal/non-elapsed evidence, and adds only pending A002 candidate/resource fields with no invented value.
- `actual-status.md` moves current state from `A002 ARCHITECT_PENDING` to `A002 PLAN_READY_FOR_MAIN_VERIFY`, keeps all P2-A/parent/child checkboxes unchecked, keeps D001 active/streak `0` and D002-D017 queued, points to `ARCH1/PLAN1`, and appends exactly one current refresh row.
- Coder remains locked. This Planner refresh performs no source/test/report/plan-rules edit, graph/analyze/build/test/measurement, Supervisor/disposition, stage, or commit.
- Next owner is Main Orchestration: independently verify the exact four ledger diffs/artifacts, then alone decide the next governance transition. No extra Architect consistency re-review is created by this translation.

### `E2-P2A-A002IMPACT1` — Fresh Pre-Edit Anvien Impact

Status: recorded / complete for the exact A002 production owner boundary.

- Exact `anvien --help`: PASS.
- Fresh pre-edit `anvien analyze --force`: exit `0`; `2217` scanned, `765` parsed, `0` failed; graph `123605` nodes / `170322` relationships; indexed commit `80eb8c3c9a9cf516dd17f8b3bf15802694bb1277`.
- File risk: `internal/graphhealth/diagnostics.go`, `internal/resolution/emit.go`, and `internal/resolution/outcome.go` are HIGH.
- Symbol risk: `AppendDiagnosticToNode`, `appendDiagnosticsFromProperties`, `emitter`, `newEmitter`, and `emitUnresolvedReference` are CRITICAL; `emitTypeScriptOutcomeDiagnostic` is LOW/0 while its file remains HIGH.
- HIGH/CRITICAL were treated as blast-radius warnings. The implementation stayed inside the exact `E2-P2A-A002ARCH1/PLAN1` boundary.

### `E2-P2A-A002SRC1` — Exact Production-First Candidate

Status: recorded / uncommitted candidate at HEAD `cc420e7ad719d90dc4b2d9991be0249e8d648daa`; staged set empty; scoped diff-check PASS.

- Exact production delta is `+56/-4`: `internal/graphhealth/diagnostics.go +52/-2`, `internal/resolution/emit.go +3/-1`, and `internal/resolution/outcome.go +1/-1`.
- One new post-production test file exists: `internal/graphhealth/diagnostics_test.go`. No existing test file changed.
- Final SHA-256 values: `diagnostics.go AEABA8A541D1C293DFA4FE411253A85A26B6AC0DD18047AFA6F9151B4932BEE8`; `emit.go 73AE20E8A18AC25F8019E2A3223407621AE7ABC358DBAF4F058E1D9AC3136060`; `outcome.go 02092F9FE7DA2A4BDB49E13056FBA3C97DC24F416141E7F27866EE80F60C1F7E`; `diagnostics_test.go 58E8CAE2C4EBBF4672D2D338CC5E1419B7CECA9F7389B0E24F51BF2F975A7CEA`.
- The candidate adds one emitter-owned, write-through diagnostic appender; first touch normalizes supported existing diagnostics once, each incoming diagnostic once, and every successful append writes `[]Diagnostic` through `Graph.AddNode`. Legacy `AppendDiagnosticToNode` remains available for all other callers.
- No forbidden owner, schema, public contract, persistence/reader, timing instrumentation, D002-D017, or unrelated dirty path changed.
- Coder report: `E:\Anvien\reports\coder\rp_coder_260825_193200_by_gpt-5_child06a_a002_paused_handoff.md`; status `CODER_A002_SOURCE_BUILD_READY_FOR_MAIN`.

### `E2-P2A-A002BUILD1` — Canonical Full-Build Validation

Status: `PASS` / exit `0`; validation only, not product-performance evidence.

- Exact command: `pwsh -NoLogo -NoProfile -NonInteractive -ExecutionPolicy Bypass -File E:\Anvien\scripts\full-build.ps1`.
- Build-holder preflight found the analyze lock free and only editor-owned MCP services; nothing was terminated.
- Vendor/runtime source verification passed with `1798` files, `0` missing, `0` extra, and `0` mismatched; canonical runtime and launcher/Vite build passed.
- Runtime: version `1.2.8`, `73,767,936` bytes, SHA-256 `11895442CA4A5708FA439A787BC541102BE7144856315BCD17E9A7B772658AED`.
- Maintenance analyze exited `0` with `2221` scanned, `767` parsed, `0` failed, graph `123811 / 170572`. This maintenance run and build duration are not A002 benchmark observations.

### `E2-P2A-A002TEST1` — Focused/Graphhealth PASS And Pre-Existing Resolution Golden FAIL

Status: all A002-focused checks and `internal/graphhealth` PASS; `internal/resolution` truthfully FAILs only on one proven pre-existing preserve-only golden. The package is not labeled PASS.

- PASS after the definitive full-build PASS, in required order: `TestDiagnosticAppenderMatchesLegacySemantics`; `TestP6DStructuredOutcomeProjectionPreservesGraphJSONAndHealthParity`; `TestResolveAttachesSourceBackedUnresolvedDiagnostics`; `TestP6C3FinalOutcomeStatusPrecedenceAndCarriageMatrix`; `TestP6BResolverCarriesCatalogValidationFailuresPerSourceSite`; `go test ./internal/graphhealth -count=1`.
- `go test ./internal/resolution -count=1` exits `1` solely at `TestProofBasedCallAccessGoldenCorpus` with `unclassified/review` instead of the stale expected classification/actionability.
- A repo-local exact current-HEAD overlay for the three A002 production files reproduces the identical complete diagnostic payload. The failure is pre-existing/preserve-only, is neither A002 regression nor A002 validation proof, and does not authorize editing its owner.
- Main independently verified HEAD, staged set, scoped status, final hashes, exact `+56/-4` production diff, new test source, and diff-check before advancing the ledger cursor.
- Current transition after later measurement history: `A002 KEEP / A003 CURRENT_BASIS_RECORDED / RESIDUAL_CAUSE_PENDING / ARCHITECT_LOCKED`. Invalid history remains preserved; `REVIEW1/DECISION1` close A002 as KEEP; `A003CURRENT1` records the promoted basis. No A003 cause/Architect/Planner/Coder evidence, detect, stage, or commit exists.

### `E2-P2A-A002MEASUREFAIL1` — Cheapapp Capture Ineligible: Missing 17-Child Instrumentation

Status: `MEASUREMENT_A002_CHEAPAPP_INCOMPLETE_MISSING_17_CHILD_INSTRUMENTATION`. This is an invalid measurement invocation, not a production attempt result, Supervisor result, no-KEEP disposition, baseline promotion/restoration, or retry authorization.

- Visible lane/report: `01a033a5-4700-7272-a28d-de3a71f58135`; `E:\Anvien\reports\Investigation\rp_child06a_a002_cheapapp_benchmark.md`; report SHA-256 `17F1D449A2766858EA70AB32D565EBBF078DA8830B36928816CB741815CE5AE8`.
- Raw root: `E:\Anvien\.tmp\child06a_a002_cheapapp_benchmark`; accepted A001 before root remained `E:\Anvien\.tmp\child06a_a001_cheapapp_benchmark\owner-final-run\optimized` and was not rerun.
- Exact one-shot command: `E:\Anvien\.tmp\child06a_a002_cheapapp_benchmark\candidate\anvien-a002-cheapapp-candidate.exe analyze E:\cheapapp.org --force --skip-git --json --progress --benchmark-json E:\Anvien\.tmp\child06a_a002_cheapapp_benchmark\candidate\benchmark.json --benchmark-label child06a-a002-cheapapp-candidate`.
- Identity/result: launch count `1`; PID `12756`; start/end `2026-08-25T21:40:32.8136848+07:00` / `2026-08-25T21:43:09.4112295+07:00`; exit `0`; executable `73,767,936` bytes / SHA-256 `11895442CA4A5708FA439A787BC541102BE7144856315BCD17E9A7B772658AED`; target HEAD `a869876ab6262dacde6cd5d432d099a91852a646`.
- Raw measured boundaries, all ineligible: process wall/CPU `156.946718300 / 110.906250000 s`; analyzer `120.335153100 s`; parent `resolution 22.664132600 s`; memory start/end/max-Sys `1,299,272 / 666,783,664 / 1,972,447,736` bytes.
- Coverage: exactly `30/30` top-level operations with matching denominators, but the benchmark root has no `resolutionChildren`/equivalent collection and its `resolution` object contains semantic counters only. Candidate child coverage is exactly `0/17`; D001 duration/calls/files/rank/share/delta, all D002-D017 values, child sum, parent residual, exclusive intervals, overlaps, and resolution CPU profile are unavailable and were not inferred.
- Same-work facts that did pass: scanned/parsed/failed `1359/887/0`; parser bytes `5,430,581`; graph `95,762 / 129,716`; projection `13,360 / 735`; DB readback `95,762 / 129,716`; all resolution semantic counters equal accepted A001.
- Main independently recomputed raw hashes: benchmark `4CDF0F9813D977EE9C5B303EB94E1376800C4FAED98B7FC61CDB82EBC07747B3`; process `E8D026F49B30A45503C0718EC0A611EB6021D77AC34538056B144DA238A2704E`; graph `8763AA2D089328F3D782A5F01E06178D171922E3A5CE96A4EE6707510976C920`; stdout `F52193AD16DD18C3DBD78CC161B503CF4091BDFAC1B363962FEBC2C50AD7FEF4`; stderr `0FDCF15222C02C25771C98E17D9220C89FDCD0294137A6AE4D9B3AAFD4761ADC`. Graph/stdout/stderr exactly match the accepted A001 authority; Ladybug/meta raw-byte differences remain non-semantic and make no acceptance claim.
- Exact cause: the frozen canonical candidate contained the A002 production appender but did not contain the accepted `resolve.go`/`types.go` child instrumentation. Setting a profile environment variable cannot create child metrics in a binary that was not built with the overlay. Raw Segment 05 `resolve.go` is forbidden for correction because it predates accepted A001 and would restore its old global import traversal.
- Required correction: preserve the unchanged A002 source candidate; independently verify and use the accepted A001 optimized child-instrumentation overlay mapping only for `internal/resolution/resolve.go` and `internal/resolution/types.go`; build a frozen measurement binary with that overlay, then run exactly one corrected Cheapapp capture after zero-competitor preflight. Restaurant Manager remains queued until the corrected Cheapapp packet is valid.
- State effect: no A002 value is promotable; A001 accepted baselines remain authoritative; D001 streak remains `0`; every checkbox remains unchanged; no retry was launched by the failing lane.

### `E2-P2A-A002OVERLAY1` — Corrected 17-Child Measurement Overlay Identity

Status: `CORRECTED_OVERLAY_VERIFIED_STATIC / BUILD_AND_RUNTIME_PENDING`. Two independent read-only checks plus Main verification agree; this authorizes only the corrected measurement build/run, not a production edit, benchmark result, Supervisor, disposition, or commit.

- Accepted optimized manifest: `E:\Anvien\.tmp\child06a_a001_cheapapp_benchmark\optimized-overlay.json`, SHA-256 `7B138FBA06B41CBD4C6709E6DF00C80C03E4015B615D1E42F787DE2A65D378F7`. It has exactly two replacements: current `internal/resolution/resolve.go` -> optimized overlay SHA-256 `92A89227E1B1B1C159DE8BE77A1F060361C9058C4F83C78BFC32E9AB40DADEA9`; current `internal/resolution/types.go` -> optimized overlay SHA-256 `A8DDCB1C61C312D3D3B164B1DD02BFD91C640F1FEB8B6DB511B73FD923C802B8`.
- A001 preservation: Main extracted `(*workspace).explicitImportNameClaimed` and proved current production equals optimized overlay byte-for-byte, while raw Segment 05 differs. The optimized overlay retains `w.importClaimsByReceiver`; raw Segment 05 would restore linear `w.imports` traversal and is forbidden.
- Instrumentation completeness: optimized overlay declares `Metrics.ChildMeasurements`, the 17 exact child IDs, interval/source/call-path/denominator fields, and profile env `ANVIEN_OP001_RESOLUTION_CPU_PROFILE`.
- A002 compatibility: the manifest does not map `internal/graphhealth/diagnostics.go`, `internal/resolution/emit.go`, or `internal/resolution/outcome.go`; build therefore consumes current A002 hashes `AEABA8A...BEE8`, `73AE20E8...36060`, and `02092F9F...C1F7E` directly. Static review found no signature/type conflict; actual compilation remains pending in the measurement lane.
- Pinned native/toolchain contract: Go `go1.26.3 windows/amd64`; Ladybug directory `E:\Anvien\third_party\ladybugdb\v0.19.1\windows-x86_64`; pinned DLL SHA-256 `20CBD87840483A2053CFF3FC2DB23A86DD802B8915D86509D41A4B709624CDB7`.
- Stop boundary: do not use raw Segment 05 overlay; stop on any hash/target/source/native drift, competing analyze/benchmark process, build/analyze failure, absent/incorrect 17-child schema/order/interval conservation/profile, or workload/semantic/graph/output mismatch. No automatic retry.

### `E2-P2A-A002MEASUREFAIL2` — Corrected Cheapapp Measurement Build Stopped Before Runtime

Status: `MEASUREMENT_INCOMPLETE_BUILD_FAILURE_MISSING_LBUG_HEADER / ANALYZE_LAUNCH_COUNT_0`. This is measurement-infrastructure failure evidence, not a production-attempt result, no-KEEP disposition, baseline promotion/restoration, or retry by the failed lane.

- Visible lane/report: `01a033a5-4700-7272-a28d-de3a71f58135`; `E:\Anvien\reports\Investigation\rp_child06a_a002_cheapapp_benchmark_corrected_17child.md`; report SHA-256 `AB0D18A91D5D344C3975B1D19597E088B5CFECF639E02392C2BAA6A040408EC9`.
- Raw root: `E:\Anvien\.tmp\child06a_a002_cheapapp_benchmark_corrected_17child`; machine-readable failure: `build-result.json`.
- One build attempt only: start/end `2026-08-25T22:16:04.2598316+07:00` / `2026-08-25T22:18:55.9468223+07:00`; wall `171.675682500 s`; exit `1`; candidate binary absent; analyze launch count `0`; benchmark/profile absent.
- Exact failure: `internal\lbugnative\native_ladybugdb.go:6:10: fatal error: lbug.h: No such file or directory`.
- All identity gates passed before build: Anvien HEAD `cc420e7ad719d90dc4b2d9991be0249e8d648daa`, staged count `0`, Cheapapp HEAD `a869876ab6262dacde6cd5d432d099a91852a646`, zero competitors, Go `go1.26.3 windows/amd64`, exact manifest/two mapped-file/A002/native-DLL hashes under `OVERLAY1`.
- The historical attempt used a direct Go `-overlay` command with repo-local HOME/TEMP/XDG/Go caches, `GOPROXY/GOSUMDB=off`, `CGO_ENABLED=1`, and only a native PATH prefix. `CGO_CFLAGS` and `CGO_LDFLAGS` were unset for this sole build, so it could not locate `lbug.h`. The build was not interrupted or retried. The Owner subsequently corrected the workflow: direct Go compilation is not a valid recovery path in this workspace; the two repository build scripts are authoritative.
- Target source, target HEAD/status, A002 source/test bytes, existing reports, ledgers, A001 artifacts, and Restaurant artifacts were unchanged. No target analyze or target-local graph mutation occurred.
- State effect: A001 baselines remain authoritative and separate; D001 streak stays `0`; every checkbox remains unchanged; no product elapsed/resource/equivalence value is created from this failure.
- Required recovery was completed by `E2-P2A-A002SCRIPTPLAN1/SCRIPTBUILD1`: preserve this failed root/report, do not repeat the ad-hoc direct build, and consume the validated reusable packet for one corrected Cheapapp capture. Restaurant and Supervisor remain locked until a valid Cheapapp packet is durably recorded.

### `E2-P2A-A002BUILDRECOVERY1` — Existing Script-Compatible Overlay Path Is Missing

Status: `NO_EXISTING_SCRIPT_COMPATIBLE_OVERLAY_PATH / SUPERSEDED_NEXT_STEP_BY_OWNER_SCRIPT_DIRECTION`. This is read-only build-interface evidence, not a production-attempt result, no-KEEP disposition, target measurement, or authorization to edit the existing build infrastructure.

- Owner workflow authority: canonical builds use `E:\Anvien\scripts\full-build.ps1` and `E:\Anvien\anvien-launcher\build.ps1`; a direct `go build` retry is prohibited.
- Fresh Main source inspection plus two independent read-only checks agree. `full-build.ps1` has no parameters, derives the fixed repository root, runs `go run -mod=vendor ..\cmd\anvien package build-runtime` from `E:\Anvien\anvien`, writes the fixed canonical runtime, invokes the fixed launcher build, then runs canonical `analyze . --force`.
- `package build-runtime` uses `cobra.NoArgs`. `resolvePackageSourceRoot` always selects parent source `E:\Anvien` before `E:\Anvien\anvien\go-src` because all required source anchors exist. `buildGoRuntimePackage` hard-codes the canonical output and build flags, has no `-overlay`, and forces `GOFLAGS` empty; ambient overlay injection is therefore unavailable.
- `prepare-go-source` has no arguments, deletes/recreates fixed `anvien\go-src` from current repo source, and has no overlay hook. The canonical workspace build ignores that fallback while parent source exists.
- `anvien-launcher\build.ps1` has no parameters and does not build the CLI runtime; it requires the fixed canonical runtime/metadata, builds Web/launcher/server outputs, copies the native DLL, and registers the protocol.
- Existing environment overrides affect Go executable or native authority only; none selects runtime source, overlay, or output. A temp-package/go-src/direct-Go workaround would be an ad-hoc workflow outside the two Owner-approved scripts.
- Exact historical conclusion: no safe pre-existing command could build the verified two-file instrumentation overlay while preserving the script contract. The latest Owner direction supersedes only its next step: keep both canonical scripts unchanged and implement the reusable `scripts/build-a00x-benchmark.ps1`. Coder owns that exact script; Cheapapp runtime launch, Restaurant Manager, Supervisor, staging, and commit remain locked until script build validation completes.

### `E2-P2A-A002SCRIPTPLAN1` — Owner-Directed Reusable A00x Benchmark Build Path

Status: `SCRIPT_PLAN_RECORDED / CONSUMED_BY_E2-P2A-A002SCRIPTBUILD1`. This is Planner control evidence for measurement-support tooling, not A002 production evidence, benchmark data, Supervisor acceptance, disposition, or commit authorization.

- Latest Owner authority requires the complete reusable contract to be durable, not remembered only from chat: `scripts/build-a00x-benchmark.ps1` is the common build path for A002 and later A00x overlay/instrumented benchmark candidates.
- Required inputs are explicit attempt ID, overlay manifest, output directory, expected overlay hash, expected mapped-source hashes, and expected candidate-source hashes. For A002 the overlay is the exact two-mapping `resolve.go`/`types.go` authority from `E2-P2A-A002OVERLAY1`, while the three A002 production hashes remain those in `E2-P2A-A002SRC1`.
- The script carries the canonical `-mod=vendor`, `ladybugdb`, trim/ldflags, Ladybug native include/link/PATH, CGO, offline Go, and toolchain environment; adds only the verified overlay, `-buildvcs=false`, and explicit repo-local `.tmp` output; copies the pinned DLL; and emits executable/DLL/overlay/source/Go/native/command provenance as JSON.
- It never edits, copies, wraps through, or replaces `scripts/full-build.ps1` or `anvien-launcher/build.ps1`; never writes canonical runtime/launcher/server/web outputs; never registers protocol or runs `analyze`; and never mutates target repositories or shared/global caches.
- After this script/native contract passes once, unchanged later A00x attempts call it with their own explicit input/hash contract and do not repeat general build-interface discovery or audit.
- If a later A00x has a genuinely different feature/input and the script fails, the exact script error is the starting evidence. Planner refreshes that attempt, and Coder adjusts only the narrow reusable-script parameter/guard/environment/tag/generated-input/provenance behavior required by that A00x. Do not audit the canonical scripts again, alter production to evade the build failure, or change canonical build interfaces unless the Owner explicitly changes this rule.
- Current Coder boundary: add only `scripts/build-a00x-benchmark.ps1` and update the existing A002 Coder handoff. After holder preflight, canonical full-build `PASS` must precede script execution; then verify PowerShell parse, one A002 script build, binary/version/DLL/provenance identities, scoped diff-check, and empty staged set. The Coder runs no target analyze.
- Failure/rollback: remove only the new script and exact `.tmp` packet if invalid. A script failure creates no numeric row, no no-KEEP increment, no baseline/checklist change, and no A002 production rollback.
- Fresh plan-authoring evidence: `anvien --help` PASS; `anvien analyze --force` exit `0`, `2223` scanned / `767` parsed / `0` failed, graph `123840 / 170601`; all four standard ledgers are current/non-stale LOW-risk documentation files with `2` inbound / `1` outbound relationships and `0` unresolved references.

### `E2-P2A-A002SCRIPTBUILD1` — Reusable A00x Script And Frozen A002 Packet Validated

Status: `PASS / CODER_A002_A00X_SCRIPT_BUILD_READY_FOR_MAIN / CORRECTED_CHEAPAPP_MEASUREMENT_PENDING`. This is measurement-support build evidence only; it creates no product benchmark value, no no-KEEP increment, no baseline/checklist change, and no A002 disposition.

- Added only `scripts/build-a00x-benchmark.ps1`, SHA-256 `ADA407C7496FCEA988276F03BAD5001ED139A4AEC9A16B9C32947DA440814EC5`; updated only the existing Coder report. Staged set remains empty and all canonical scripts/package-runtime/A002 source-test/ledger/target identities remain unchanged.
- Exact holder preflight first misclassified `MsMpEng` PID `3840`; termination returned Access Denied and no build started. Narrow Restart Manager proof then identified five actual canonical EXE/DLL holders (`5564`, `8320`, `10888`, `12552`, `7484`), all installed `anvien mcp` processes. Only those proven holders were terminated; all five disappeared, both outputs accepted exclusive access, and analyze lock was free before build.
- Canonical `scripts/full-build.ps1` ran after holder clearance, exit `0`, start/end `2026-08-26T02:52:20.7205211Z / 2026-08-26T02:58:30.0364325Z`; elapsed `369.3159114 s` is validation only. Vendor verification passed; maintenance analyze recorded `2224` scanned / `767` parsed / `0` failed and is not target benchmark evidence.
- PowerShell parse PASS (`4924` tokens, `0` errors). Bad-overlay-hash and output-outside-`.tmp` guards each failed closed before any output/work root and do not count as valid builds.
- Exactly one valid A002 script invocation ran, exit `0`, start/end `2026-08-26T03:05:27.6669857Z / 2026-08-26T03:08:38.2045746Z`. Go identity is `go version go1.26.3 windows/amd64`.
- Frozen packet root: `E:\Anvien\.tmp\child06a_a002_a00x_benchmark_build`. Executable `anvien-a002-benchmark.exe`: version `1.2.8`, `73,816,576` bytes, SHA-256 `0F8B8244ABC80339A73E3A29F38D32F55A7A6A65281A64C291785ACF8F4A241E`. DLL SHA-256 `20CBD87840483A2053CFF3FC2DB23A86DD802B8915D86509D41A4B709624CDB7`. Provenance SHA-256 `673FCCFEA17C48640C0AC62243479FE5EFD71E172CB1D4136D046C806E5855F1`.
- Provenance independently verifies schema `1`, attempt `A002`, HEAD `cc420e7ad719d90dc4b2d9991be0249e8d648daa`, exact build argv/flags, `2/2` overlay mappings, `3/3` A002 candidate sources, `3/3` native inputs, output hashes/version/bytes, exit `0`, and all Go cache/temp/home/AppData/XDG paths beneath the packet-local `.a00x-A002-work` directory.
- Validation artifacts are under `E:\Anvien\.tmp\child06a_a002_a00x_script_validation`; packet verification and preserve-boundary verification both record `PASS`. Coder report: `E:\Anvien\reports\coder\rp_coder_260825_193200_by_gpt-5_child06a_a002_paused_handoff.md`.
- Next action: run exactly one corrected Cheapapp capture with the frozen executable and accepted A001 before; Restaurant Manager and per-attempt Supervisor remain locked until the complete Cheapapp packet is durably recorded.

### `E2-P2A-A002CHEAPAPP1` — Valid Frozen Cheapapp Candidate Measurement Recorded

Status: `MEASUREMENT_COMPLETE / CHEAPAPP_MEASUREMENT_RECORDED / RESTAURANT_MEASUREMENT_PENDING`. This is one target's measurement/equivalence/resource evidence only. It does not promote A002, decide `KEEP`, change the accepted A001 baseline, increment/reset the D001 no-KEEP streak, check any parent/child, open Supervisor, or authorize detect/stage/commit.

- Visible lane/turn: task `01a033a5-4700-7272-a28d-de3a71f58135`, turn `01a03c2e-64d2-7f31-ad22-5f4488977209`, completed without error after exactly one target analyze launch.
- Report: `E:\Anvien\reports\Investigation\rp_child06a_a002_cheapapp_benchmark_frozen_17child.md`.
- Raw root: `E:\Anvien\.tmp\child06a_a002_cheapapp_frozen_17child_20260826`; machine-readable authority is `preflight.json`, `validation.json`, `comparison.json`, `output-manifest.json`, and `candidate/{benchmark.json,process.json,stdout.json,stderr.log,resolution.cpu.pprof}`.
- Exact frozen executable: `E:\Anvien\.tmp\child06a_a002_a00x_benchmark_build\anvien-a002-benchmark.exe`, version `1.2.8`, `73,816,576` bytes, SHA-256 `0F8B8244ABC80339A73E3A29F38D32F55A7A6A65281A64C291785ACF8F4A241E`; adjacent DLL SHA-256 `20CBD87840483A2053CFF3FC2DB23A86DD802B8915D86509D41A4B709624CDB7`; provenance SHA-256 `673FCCFEA17C48640C0AC62243479FE5EFD71E172CB1D4136D046C806E5855F1`.
- Exact command: `anvien-a002-benchmark.exe analyze E:\cheapapp.org --force --skip-git --json --progress --benchmark-json E:\Anvien\.tmp\child06a_a002_cheapapp_frozen_17child_20260826\candidate\benchmark.json --benchmark-label child06a-a002-cheapapp-frozen-17child-20260826`.
- Process proof: launch count `1`; PID `13232`; start/end `2026-08-26T10:58:53.0369124+07:00` / `2026-08-26T11:01:09.5419385+07:00`; exit `0`; process wall/CPU `136.729876000 / 105.125000000 s`; zero competing analyze/benchmark processes immediately before launch.
- Complete schema/conservation proof: `30/30` unique top-level operations; `17/17` unique resolution children; raw instrumentation order matches accepted A001 and `comparison.json` normalizes the exact rows to `B2-P2A-A001-D001..D017`; D001 denominator `calls=27890; files=887`; `2675` exclusive intervals; overlap `0`; every child duration equals its interval sum; parent/child-sum/residual `19.040468000 / 19.024756200 / 0.015711800 s`; all durations/denominators nonnegative; profile `61,264` bytes and nonempty.
- Accepted A001 -> unpromoted A002 candidate elapsed values:
  - D001 `resolve_calls`: `25.045225300 -> 3.090914200 s`; delta `-21.954311100 s (-87.658669%)`.
  - parent `resolution`: `184.481061700 -> 19.040468000 s`; delta `-165.440593700 s (-89.678904%)`.
  - analyzer total: `274.474620900 -> 100.843249000 s`; delta `-173.631371900 s (-63.259536%)`.
  - process wall: `279.105934600 -> 136.729876000 s`; delta `-142.376058600 s (-51.011477%)`.
- Same-work/equivalence proof: files `1359/887/0`; parser bytes `5,430,581`; graph and DB rows `95,762 / 129,716`; dependency edges/projection unresolved `13,360 / 735`; all operation/child denominators match accepted A001; resolution semantic counters, diagnostics, and outcomes are equal. Canonical Graph JSON, stdout, and stderr hashes match exactly: `8763AA2D089328F3D782A5F01E06178D171922E3A5CE96A4EE6707510976C920`, `F52193AD16DD18C3DBD78CC161B503CF4091BDFAC1B363962FEBC2C50AD7FEF4`, and `0FDCF15222C02C25771C98E17D9220C89FDCD0294137A6AE4D9B3AAFD4761ADC`.
- Ladybug/meta raw bytes differ and are not used as the semantic gate: candidate Ladybug `597,946,368` bytes / `2FEC9F35F7662B45789865DC9EB40A97555AAA1A68DB243BBC3A8D9B778C5AF4`; candidate meta SHA-256 `2FC9B672D97543DF42BED79F7CF19F113A70E1AA3A0561A871DA263557A9D43E`. Logical DB counts, canonical Graph JSON, public output, and affected semantics match.
- Resource fields: accepted A001 -> A002 start allocation `1,309,264 -> 1,315,328` bytes (`+6,064`); end allocation `863,629,352 -> 1,055,121,264` (`+191,491,912`); maximum observed `Sys` `1,989,634,552 -> 1,972,115,960` (`-17,518,592`). Structural source/Coder proof remains `O(T)` map entries/slice headers with the graph and cache sharing one normalized slice and no second retained diagnostic-object copy.
- Source/target identity: Anvien HEAD stayed `cc420e7ad719d90dc4b2d9991be0249e8d648daa`, staged set stayed empty, the four A002 hashes remained exact, Cheapapp HEAD stayed `a869876ab6262dacde6cd5d432d099a91852a646`, and its complete `13`-row pre/post status arrays are identical.
- Main independently recomputed the process/schema/conservation/denominator/profile/workload/graph/DB/hash/status facts from raw artifacts. A separate bounded read-only check found `FAILED: none` and `NO_EVIDENCE: none` for the assigned invariants. The lane report is therefore accepted as a valid measurement input, not as a campaign verdict.
- The next action after this Cheapapp record was Restaurant recovery. That transition completed later under `E2-P2A-A002RESTAURANTBLOCK1/RESTAURANT1`; the current gate is now the fresh visible A002 per-attempt Supervisor.

### `E2-P2A-A002RESTAURANTBLOCK1` — Handoff Restaurant Task Unavailable

Status: `HISTORICAL_BLOCKER / RESOLVED_BY_REPLACEMENT_LANE`. This was an orchestration/task-availability blocker, not a product measurement failure, no-KEEP result, source defect, baseline change, or disposition.

- Handoff-specified task: `01a033a5-4ec3-7e42-8946-0ab9172f6088`.
- Fresh `read_thread` on host `local` returned `No Codex thread found`.
- Fresh active task listing did not contain the ID; archived task listing returned no tasks.
- Reversible recovery attempt `set_thread_archived(archived=false)` on the exact ID also returned `No Codex thread found`.
- Local Codex project metadata still contains a project-reference entry for the ID, but no rollout/session file exists beneath the local sessions tree; the reference is stale metadata rather than a callable task.
- No Restaurant command was sent, no target process launched, no target/source/artifact changed, and no substitute lane was opened.
- Main cannot perform the external target measurement itself or use an internal/hidden lane because the orchestration boundary requires a separate visible functional lane. Creating a replacement visible task is a materially new task-management action and requires explicit Owner authority.
- Historical effect: Cheapapp remained valid/unpromoted; accepted A001 baselines stayed separate; D001 streak stayed `0`; all parent/child checkboxes stayed unchecked; no detect/stage/commit/P3/Child 07 action occurred.
- Resolution: replacement visible task `01a03c72-5f6d-7fd0-948a-88b06a1ebb65` executed the frozen packet once and produced the valid `E2-P2A-A002RESTAURANT1` packet. This historical blocker no longer controls the current cursor.

### `E2-P2A-A002RESTAURANT1` — Valid Frozen Restaurant Manager Candidate Measurement Recorded

Status: `MEASUREMENT_COMPLETE / BOTH_TARGET_MEASUREMENTS_RECORDED / SUPERVISOR_PENDING`. This is the second target's measurement/equivalence/resource evidence only. It does not promote A002, decide `KEEP`, change either accepted A001 baseline, change the D001 no-KEEP streak, check any parent/child, supply a Supervisor verdict, or authorize detect/stage/commit.

- Visible replacement task: `01a03c72-5f6d-7fd0-948a-88b06a1ebb65`; exactly one target analyze launch.
- Report: `E:\Anvien\reports\Investigation\rp_child06a_a002_restaurant_manager_benchmark_frozen_17child.md`; report SHA-256 `77BA7EA058A8F4F42059ED15E28F896840C31E6E70A5098E50DB26773A9DCCB1`.
- Raw root: `E:\Anvien\.tmp\child06a_a002_restaurant_manager_frozen_17child_20260826`; machine-readable authority includes `candidate/benchmark.json`, `candidate/process.json`, canonical outputs, profile, and the packet comparison/validation artifacts.
- Exact frozen executable: `E:\Anvien\.tmp\child06a_a002_a00x_benchmark_build\anvien-a002-benchmark.exe`, version `1.2.8`, `73,816,576` bytes, SHA-256 `0F8B8244ABC80339A73E3A29F38D32F55A7A6A65281A64C291785ACF8F4A241E`; adjacent DLL SHA-256 `20CBD87840483A2053CFF3FC2DB23A86DD802B8915D86509D41A4B709624CDB7`; provenance SHA-256 `673FCCFEA17C48640C0AC62243479FE5EFD71E172CB1D4136D046C806E5855F1`.
- Exact command: `anvien-a002-benchmark.exe analyze E:\Restaurant_manager --force --json --progress --exclude electron/renderer/src/api/userApi.ts --benchmark-json E:\Anvien\.tmp\child06a_a002_restaurant_manager_frozen_17child_20260826\candidate\benchmark.json --benchmark-label child06a-a002-restaurant-manager-frozen-17child-20260826`.
- Process proof: launch count `1`; PID `15688`; start/end `2026-08-26T05:12:53.9496932+00:00` / `2026-08-26T05:15:19.0185984+00:00`; exit `0`; process wall/CPU `145.066210900 / 114.203125000 s`.
- Complete schema/conservation proof: `30/30` top-level operations; ordered `17/17` resolution children; D001 denominator `calls=86030; files=1234`; `3716` exclusive intervals; overlap `0`; parent/child-sum/residual `21.242055400 / 21.217471500 / 0.024583900 s`; conservation PASS; nonempty profile `59,197` bytes.
- Accepted A001 -> unpromoted A002 candidate elapsed values:
  - D001 `resolve_calls`: `40.769294200 -> 9.909636600 s`; delta `-30.859657600 s (-75.693382%)`.
  - parent `resolution`: `136.436879300 -> 21.242055400 s`; delta `-115.194823900 s (-84.430855%)`.
  - analyzer total: `215.972455200 -> 109.339859600 s`; delta `-106.632595600 s (-49.373239%)`.
  - process wall: `218.680628900 -> 145.066210900 s`; delta `-73.614418000 s (-33.662981%)`.
- Same-work/equivalence proof: files `1556/1234/0`; graph and DB rows `137229 / 190644`; all operation/child denominators, resolution semantic counters, diagnostics, and outcomes PASS. Canonical Graph JSON, stdout, and stderr match accepted A001 exactly: `09708B038AD63B659A95BAC34304A752F2C6350952B91CF435EBD5C5F9D94B7C`, `CD0028F3B3939451FF4D3623DF7DD7482A993E0934D2DFA773682D37D4FB0E94`, and `0FDCF15222C02C25771C98E17D9220C89FDCD0294137A6AE4D9B3AAFD4761ADC`.
- Ladybug/meta raw bytes are recorded but non-governing alone: Ladybug `503,070,720` bytes / SHA-256 `5CDE4541F687E515CF04C2824F89EE9E6C701B38EE7D090AB75438B9572E1E48`; meta `267` bytes / SHA-256 `B45157FC1B230FE1522E0A395226997ED071534F4BC73DDED8C870F5CE03078A`.
- Resource fields: accepted A001 -> A002 start allocation `1,461,312 -> 1,475,192` bytes (`+13,880`); end allocation `876,051,256 -> 1,119,052,736` (`+243,001,480`); maximum observed `Sys` `3,050,756,728 -> 2,911,591,032` (`-139,165,696`). Structural proof remains `O(T)` slice-header map with graph/cache sharing the same normalized slice and no second retained diagnostic-object copy.
- Source/target boundary: Anvien HEAD remained `cc420e7ad719d90dc4b2d9991be0249e8d648daa`; Restaurant HEAD remained `605c0bda99491789e7f07628ec4b3d39a3ae1c67`; staged set and the exact pre-existing Restaurant status rows remained unchanged.
- Raw hashes: benchmark `C2E94CD1939C8A0435A970D00341EDF4F366E9418B53284058D92355195E51FF`; process `2045359544842E3E110C32B1E93CDB19BAFA115F4B64FE4B16C77CF3566B7E16`; profile `97F0BB7E24D5DF9585F0B0221B352E63B2E603E26BE9CB959371885CE646F4DB`.
- Current effect: `CHEAPAPP1/RESTAURANT1` form the complete two-target measurement input. Both remain separate and unpromoted; accepted baselines, D001 streak `0`, and every checkbox remain unchanged. The next owner is one fresh visible A002 per-attempt Supervisor.

### `E2-P2A-A002REVIEW1` — Independent A002 Per-Attempt Supervisor PASS

Status: `SUPERVISOR_A002_PASS`. This is independent acceptance evidence; it does not itself decide `KEEP`, promote baselines, change streak/checklists, detect, stage, commit, or open P3.

- Visible Supervisor task/turn: `01a03c97-02f9-7560-ae31-4a333c31798e` / `01a03c97-05a9-7242-91ce-374722cf204d`.
- Report: `E:\Anvien\reports\Supervisor\rp_supervisor_260826_130617_by_gpt-5_child06a_a002_run_scoped_diagnostic_appender.md`; SHA-256 `11BEA6715CD8626F80746B158FE755347E6924EAC560B2C5171525FBDE58A78B`.
- Exact verdict: `SUPERVISOR_A002_PASS`.
- Current identity recheck: HEAD `cc420e7ad719d90dc4b2d9991be0249e8d648daa`; staged set empty; production numstat exactly `+56/-4`; all four A002 hashes match; explicit forbidden owners are clean.
- Source/invariant clearance: legacy callers remain on `AppendDiagnosticToNode`; the two resolution routes share one closure created by the single per-run `newEmitter`; first touch and incoming normalization, merge policy, per-append write-through, `[]Diagnostic` representation, sequential lifecycle, `O(T)` map/slice-header ownership, and no I/O/goroutine/lock/global cache/flush/finalizer pass.
- Runtime/build/test clearance: frozen executable/DLL/provenance identities match. Existing canonical full-build evidence remains identity-bound. Focused checks and `internal/graphhealth` remain PASS; `internal/resolution` remains truthfully FAIL only on the identical current-HEAD-overlay preserve-only golden and is not called PASS.
- Target clearance: both raw packets independently pass one launch/exit `0`, `30/30`, ordered `17/17`, denominator equality, conservation/zero overlap, nonempty profile, workload/graph/DB/semantic/Graph JSON/stdout/stderr equivalence, resource fields, and unchanged target boundary.
- Elapsed clearance remains separate: Cheapapp D001/parent/process `25.045225300 -> 3.090914200`, `184.481061700 -> 19.040468000`, `279.105934600 -> 136.729876000 s`; Restaurant `40.769294200 -> 9.909636600`, `136.436879300 -> 21.242055400`, `218.680628900 -> 145.066210900 s`.
- Residual same-invariant blockers: none. Next owner: Main Orchestration for disposition only.

### `E2-P2A-A002DECISION1` — Official Main KEEP And Separate Baseline Promotion

Status: `A002 KEEP / A003 RESIDUAL_ATTRIBUTION_COMPLETE / ARCHITECT_ACTIVE`.

- Main accepts `E2-P2A-A002REVIEW1` and decides `A002 = KEEP` because D001, parent, and process elapsed all decrease on Cheapapp independently and Restaurant Manager independently, while correctness/equivalence/resource/lifecycle gates and Supervisor pass.
- Promotion remains repo-specific and never averaged or combined:
  - Cheapapp accepted current A002 baseline: D001 `3.090914200 s`; parent `19.040468000 s`; analyzer/process `100.843249000 / 136.729876000 s`; denominator `calls=27890; files=887`.
  - Restaurant accepted current A002 baseline: D001 `9.909636600 s`; parent `21.242055400 s`; analyzer/process `109.339859600 / 145.066210900 s`; denominator `calls=86030; files=1234`.
- Checklist/streak effect: P2-A, `B1-P1A-OP001`, and D001 remain unchecked; D002-D017 remain queued/unopened; D001 no-KEEP streak resets/remains `0`. No child or parent is terminalized.
- Source effect: the exact A002 four-file source/test boundary becomes accepted current state. Corrected `plan-rules.md` requires its scoped detected implementation progress checkpoint before A003 production editing; that checkpoint does not close P2-A or check D001/parent.
- Commit authority: accepted-attempt progress checkpoint only. P3-C retains one local closure commit and does not prohibit earlier scoped checkpoints.
- Next cursor: A003 on the same active D001. A002 Architect direction expired at disposition. No production edit is permitted until current accepted residual cause/owner/full call path is proven, a fresh visible A003 Architect decides direction, and Planner records A003 `PLAN1`.

### `E2-P2A-A002COMMIT1` — Accepted A002 Progress Checkpoint Complete

Status: `A002_CHECKPOINT_COMPLETE`; accepted-state fact only. This checkpoint does not close P2-A, check the parent/D001, change streak `0`, authorize A003 by itself, or satisfy P3-C closure.

- Before commit, Main completed fresh `anvien analyze --force` with `2231` scanned / `767` parsed / `0` failed, `anvien detect-changes --repo E:\Anvien --scope all` exit `0`, and the exact cached scoped diff-check PASS. These are Main-owned transition facts; Planner did not rerun them.
- Exact commit: `ecf825d709b761390a5df4a2147b6ed6eec04499`; parent `2e20a44d5ef375cac17968d76f170676262db92c`; subject `perf(graphhealth): checkpoint accepted A002 diagnostic appender`.
- Exact nine-path manifest: `internal/graphhealth/diagnostics.go`; new `internal/graphhealth/diagnostics_test.go`; `internal/resolution/emit.go`; `internal/resolution/outcome.go`; `scripts/build-a00x-benchmark.ps1`; `reports/coder/rp_coder_260825_193200_by_gpt-5_child06a_a002_paused_handoff.md`; `reports/Investigation/rp_child06a_a002_cheapapp_benchmark_frozen_17child.md`; `reports/Investigation/rp_child06a_a002_restaurant_manager_benchmark_frozen_17child.md`; and `reports/Supervisor/rp_supervisor_260826_130617_by_gpt-5_child06a_a002_run_scoped_diagnostic_appender.md`.
- Staged set is empty after commit. Unrelated P1/package/launcher/Child 06/Child 07/Owner paths were excluded.
- State effect: accepted A002 bytes are now the exact A003 source checkpoint. All numeric baselines, denominators, checkboxes, and D001 streak remain unchanged.

### `E2-P2A-A003CURRENT1` — Accepted A002 Basis And Residual-Attribution Lock

Status: `CURRENT_BASIS_RECORDED / RESIDUAL_CAUSE_PENDING / ARCHITECT_LOCKED`.

- Active parent/child remain `B1-P1A-OP001 resolution` and `B2-P2A-A001-D001 resolve_calls`; complete child list remains `17/17`; D002-D017 stay queued/unopened.
- Current accepted values are the separate promoted A002 baselines recorded under `DECISION1`; D001 streak is `0`.
- `E2-P2A-A003ATTRIB1` records the existing residual packet without rerunning work. Separate profile facts: Cheapapp `resolveCall 2.76 s`, appender `0.92 s`, normalization/policy/decoder `0.87 s`; Restaurant `resolveCall 9.19 s`, appender `2.35 s`, normalization/policy `2.11 s`, decoder `2.10 s`. These are cumulative CPU samples, overlapping, non-additive, non-averaged, and not elapsed proof.
- Retained cause: per diagnostic, `diagnosticAppender.appendToNode` enters structured normalization and `decodeStructuredResolutionOutcome` decodes the same trimmed `Diagnostic.Note` into an envelope and again into a typed outcome. Exact retained-work owner: `internal/graphhealth/diagnostics.go::decodeStructuredResolutionOutcome:377-400`; decode sites `385/396`; per-diagnostic entry `appendToNode:60-88`, normalization line `77`.
- Complete path: `internal/analyze.Run:365-370 -> runPhase:1134-1155 -> resolution.ResolveBoundInto:57-128 -> w.files/ir.Calls:91-94 -> resolveCall:385-605`. Repository-unresolved branches at lines `388/392/560/564` call `emitUnresolvedReference:182-225 -> e.diagnosticAppender:220`; the TypeScript branch `516-535` calls `recordTypeScriptLookup:926-977 -> emitTypeScriptOutcomeDiagnostic:372-397 -> e.diagnosticAppender:392`. `newEmitter:31-42` binds `NewDiagnosticAppender:52-58` to `appendToNode`.
- No exact invocation count, path-family split, or elapsed-time share inside D001 is claimed; not all `encoding/json.Unmarshal` samples are assigned exclusively to this decoder. No solution or gain is inferred by this evidence.
- Visible Architect task `01a03cea-15cc-7d73-bf07-a984c69517d9` returned the updated canonical direction below. No A003 source/build/test/measurement/Supervisor/disposition exists yet.

### `E2-P2A-A003ARCH1` — Canonical Single-Interpretation Decoder Architecture

Status: `ARCHITECT_A003_READY_FOR_PLANNER`.

- Report: `E:\Anvien\reports\system-architect\rp_system-architect_260826_141200_by_gpt-5_child06a_a003_residual_direction.md`; visible task `01a03cea-15cc-7d73-bf07-a984c69517d9`.
- The updated canonical direction supersedes and withdraws the earlier status-marker fast-path wording: one Diagnostic -> one canonical outer-Note interpretation -> one semantic result (`UNSTRUCTURED`, `STRUCTURED_VALID`, or `STRUCTURED_INVALID`) -> one policy decision -> one graph write. There is no production primary/recovery/legacy/retry/fallback decoder.
- Exact semantic owner: `internal/graphhealth/diagnostics.go::decodeStructuredResolutionOutcome:377-400`. Allowed implementation surface is that decoder plus, only if necessary, one private non-exported decode-result/wire representation in the same file exclusively owned by it. No other existing semantic owner or production file may change.
- Test owner after production: append A003-only differential cases to `internal/graphhealth/diagnostics_test.go`; the test-local pre-A003 oracle is not a production path. `compute_test.go` and `p6d_outcome_projection_test.go` remain run-only.
- Anvien: containing file HIGH; exact symbol LOW; `12` impacted symbols, `1` direct caller, `1` module, `0` processes across three graph-health files. HIGH remains a scope warning.
- Preserve the externally observed `(outcome, structured, valid)` tuple; exact case-sensitive marker presence and `SourceSiteStatus`; invalid structured evidence remains fail-closed; nested authority decode; diagnostic/graph/output/persistence/lifecycle/resource contracts. No retained cache/shared contract/I/O/goroutine/lock/global state.
- Expected gain is qualitative only at D001, parent, and process boundaries on each target independently. No numeric prediction or cross-target aggregate exists.
- Exact rollback owns only the canonical decoder/private-result hunk and A003-specific test additions. STOP on any need for a second decoder route, another semantic owner/file, shared/public contract, cache, emitter/persistence/instrumentation surface, D002-D017, or weakened fail-closed behavior.

### `E2-P2A-A003PLAN1` — Exact Planner Translation

Status: `SUPERSEDED MAIN-AUTHORED DRAFT / NOT EXECUTION AUTHORITY`.

- Provenance correction: Main materialized this draft at commit `2e20a44d5ef375cac17968d76f170676262db92c` before the separate visible Planner lane completed. It remains searchable historical draft evidence and is not overwritten, reused, or treated as a fresh Planner fact.
- `E2-P2A-A003PLAN2` below supersedes this draft for owner, algorithm, tests, build, measurement, rollback, STOP, and current-lane authority.

- Production-first: replace the two independent full-note interpretations with one canonical presence-aware interpretation owned by `decodeStructuredResolutionOutcome`; it produces structure evidence, typed outcome, and validity together. No second production decode route survives or is introduced.
- Test-after-production: append differential tuple/policy cases covering status evidence, exact note markers, both evidence sources, markerless valid structured notes, unstructured/malformed/non-object/null/type-error inputs, missing/exact/case-variant markers, duplicates, unknown fields, conflicting evidence, and invalid authority/proof.
- Validation order: current Coder Anvien file-detail/impact -> production -> tests -> holder/lock preflight -> canonical full build -> focused parity/appender/P6D/resolution checks -> `internal/graphhealth` -> truthful `internal/resolution` execution with the unchanged preserve-only golden classification.
- Measurement: unchanged accepted A00x build script/17-child overlay contract; Cheapapp and Restaurant Manager separately compare accepted A002 with accepted-A002-plus-only-A003, recording D001/parent/analyzer/process, 30/17 rows, denominators, graph/output/semantic/resource parity; no renewed build-interface audit.
- After both packets, a fresh visible Supervisor reviews the exact candidate. Main alone decides disposition. Accepted A003 receives a scoped detected progress commit; rejected A003 rolls back exact owned bytes and increments the D001 streak.
- Historical next transition is superseded. A002 checkpoint is complete; no Coder is active; current transition is PLAN2 handoff to Main followed by a fresh visible Coder.

### `E2-P2A-A003PLAN2` — Visible Planner Canonical Decoder Execution Authority

Status: `PLANNER_READY_FOR_MAIN_HANDOFF / NO CODER ACTIVE / FRESH CODER PENDING AFTER PLANNER HANDOFF`.

- Planner identity: visible task `01a03d1b-662d-7f90-9a60-38271a10c077`. Architecture input: task `01a03cea-15cc-7d73-bf07-a984c69517d9`, report `E:\Anvien\reports\system-architect\rp_system-architect_260826_141200_by_gpt-5_child06a_a003_residual_direction.md`, exact verdict `ARCHITECT_A003_READY_FOR_PLANNER`. No attribution/A001/A002 gate or architecture review was rerun.
- Plan-authoring evidence only: `anvien --help` PASS; `anvien analyze --force` exit `0`, `2231` scanned / `767` parsed / `0` failed, graph `123946 / 170707`; each standard ledger is non-stale LOW-risk documentation with `2` inbound / `1` outbound / `0` unresolved relationships. This proves plan freshness/relationships only and is not product/build/test/benchmark/Supervisor evidence.
- Exact production owner: only `internal/graphhealth/diagnostics.go::decodeStructuredResolutionOutcome`; at most one private non-exported decode-result/wire representation in the same file, exclusively owned by the decoder. No other existing semantic owner or production file may change.
- Canonical algorithm: one parse/traversal of the outer `Diagnostic.Note` must simultaneously produce exact case-sensitive top-level `schemaVersion`/`status` presence, typed structured outcome/wire fields, and syntax/type/validation error state. One decision combines those products with existing `SourceSiteStatus` evidence and yields exactly `UNSTRUCTURED`, `STRUCTURED_VALID`, or `STRUCTURED_INVALID`; the existing policy then leads to one graph write. No production primary/recovery/legacy/retry/fallback decoder or second full-note authority exists. Nested Authority subdocument decode remains unchanged.
- External parity/fail-closed: preserve the complete externally observed `(outcome, structured, valid)` tuple and classification/actionability for every accepted input class. If structured evidence exists from `SourceSiteStatus` or exact marker presence, any syntax/type/required-field/proof/authority/contract error is `STRUCTURED_INVALID`, retains current fail-closed `unclassified/review`, and is never reinterpreted as unstructured.
- Production first, tests second: only after production is correct, append A003-specific bytes to existing `internal/graphhealth/diagnostics_test.go`, preserving all A002 bytes. Planned focused test `TestDecodeStructuredResolutionOutcomeCanonicalSingleInterpretationParity` uses a test-local pre-A003 oracle—not a production fallback—to compare full tuple/policy parity across status evidence, exact marker evidence, both evidence sources, valid markerless structured notes, unstructured/malformed/non-object/null/type-error inputs, missing/exact/case-variant markers, duplicate keys, unknown fields, conflicting evidence, and invalid proof/authority. `compute_test.go` and `p6d_outcome_projection_test.go` are run-only.
- Fresh Coder pre-edit sequence: `anvien --help` -> `anvien analyze --force` -> file-detail `internal/graphhealth/diagnostics.go` -> exact upstream impact for `decodeStructuredResolutionOutcome` -> explicit blast-radius record. Containing-file HIGH is a scope warning; the exact owner boundary remains mandatory.
- Holder/build/test order: prove and terminate all build-output holders, then run canonical `pwsh -NoLogo -NoProfile -NonInteractive -ExecutionPolicy Bypass -File E:\Anvien\scripts\full-build.ps1`. Only after PASS run the new decoder parity test, `TestDiagnosticAppenderMatchesLegacySemantics`, `TestP6DStructuredOutcomeProjectionPreservesGraphJSONAndHealthParity`, `TestResolveAttachesSourceBackedUnresolvedDiagnostics`, `TestP6C3FinalOutcomeStatusPrecedenceAndCarriageMatrix`, `TestP6BResolverCarriesCatalogValidationFailuresPerSourceSite`, `go test ./internal/graphhealth -count=1`, and truthfully `go test ./internal/resolution -count=1`. Only unchanged `TestProofBasedCallAccessGoldenCorpus` may fail; the package is never called PASS and any new/changed failure blocks A003.
- Measurement: reuse unchanged `scripts/build-a00x-benchmark.ps1` and the accepted 17-child overlay/native/runtime/provenance contract. Build no new interface and edit no script. Each target compares accepted A002 checkpoint plus only A003 bytes against its own A002 baseline. Cheapapp and Restaurant Manager independently record D001/parent/analyzer/process, `30/30`, `17/17`, denominators, graph/readback/output/semantic parity, and resources; never average/combine.
- Acceptance history: both packets and the fresh visible A003 Supervisor completed; exact verdict is `SUPERVISOR_A003_PASS`. The later Owner disposition is recorded only under `E2-P2A-A003OWNERKEEP1`. PLAN2 itself changes no benchmark number, checkbox, streak, or speed claim and cannot authorize a new attempt.
- Exact rollback: remove only the A003 canonical-decoder hunk, any decoder-exclusive private representation, and A003-appended test bytes. Preserve the entire A002 checkpoint and every A002 test byte, A001/P1 work, script, reports, ledgers, and unrelated/user/protected changes.
- STOP: second decoder/path/direction; another owner/file/shared/public contract/cache/emitter/policy/Authority decoder/persistence/reader/instrumentation/D002-D017 surface; weakened fail-closed semantics; new validation failure; denominator/workload mismatch; inability to prove parity; or unavailable independent two-target evidence.
- Lane state: task `01a03d17-45f5-7831-8e84-3ca8181d69c6` was archived before graph/source/test/build/edit activity because it carried stale pre-PLAN2 context; it is not inspected, resumed, or reused. Main opens one fresh visible Coder from checkpoint `ecf825d7` plus PLAN2.
- Durable Planner report: `E:\Anvien\reports\planner\rp_planner_260826_by_gpt-5_child06a_a003_canonical_decoder.md`. Next owner is Main Orchestration for handoff acceptance and fresh-Coder creation.

### `E2-P2A-A003IMPACT1/SRC1/BUILD1/TEST1/PACKET1` — Coder And Frozen Candidate Complete

Status: `CODER_A003_SOURCE_BUILD_COMPLETE / CANDIDATE_PACKET_READY / SUPERVISOR_A003_PASS / OWNER_KEEP / RESTORE_COMPLETE / A003_CHECKPOINT_COMPLETE`.

- Fresh visible Coder task `01a03d34-b740-7982-8ff7-c2fec47940a6` used PLAN2 only. Fresh analyze passed at clean HEAD `90edf7fe99cd9600b99c1947c2483f8c5fe2c67c`: `2231` scanned / `766` parsed / `0` failed, graph `123887 / 170645`. `diagnostics.go` is HIGH; exact upstream `decodeStructuredResolutionOutcome` impact is LOW with `7` impacted symbols, `1` direct caller, `1` module, `0` processes.
- Source is exactly `internal/graphhealth/diagnostics.go +109/-7` and `internal/graphhealth/diagnostics_test.go +182/-0`. One outer-note `json.Decoder` traversal replaces the two full-note decodes; no second production decoder/path or other owner changed. Production preceded the appended 25-case test-local oracle parity test. Source SHA-256 values are `6DE54D97A1B95E877B686DC3459238598F6060D28CE6329357A80BD7FC376D30` and `06677DBA4FE9EDA4FE8651E08B93ABC8659129A61355E34BFFA10380AE429371`; scoped diff-check passed; staged set is empty.
- Canonical full build passed before tests. Focused graphhealth, focused resolution, and full `internal/graphhealth` passed. Full `internal/resolution` truthfully failed only at unchanged preserve-only `TestProofBasedCallAccessGoldenCorpus` with the documented `unclassified/review` payload; it is not called package PASS and its owner remains untouched.
- Coder report: `E:\Anvien\reports\coder\rp_coder_260826_by_gpt-5_child06a_a003_canonical_decoder.md`; verdicts `CODER_A003_SOURCE_BUILD_READY_FOR_MAIN` and `CODER_A003_CANDIDATE_PACKET_READY_FOR_MAIN`.
- Frozen packet: `E:\Anvien\.tmp\child06a_a003_a00x_benchmark_build`; executable version/bytes/SHA-256 `1.2.8 / 73,823,232 / 2DBBD7AC70C04FB62C3D8AB4A90F50E7F90C51251E82B67883A58B61D42B426B`; DLL SHA-256 `20CBD87840483A2053CFF3FC2DB23A86DD802B8915D86509D41A4B709624CDB7`; provenance SHA-256 `CB91D5F7FC2EB2810E6A02AED36B5C7E0791F3973487356336B67A1BC8A66F76`. Provenance records schema/attempt `1/A003`, build exit `0`, overlay `2/2`, candidate sources `3/3`, and native inputs `3/3`.
- Main consumed the completed source/build/test handoff once and released both target measurements without rerunning passed Coder gates. `E2-P2A-A003CHEAPAPP1/RESTAURANT1/REVIEW1` record the separate packets and `SUPERVISOR_A003_PASS`. Owner selects exact A003 `KEEP`; restoration is complete at the reviewed hashes and checkpoint `b6bf45bce95323aa6b53b182edfea8628bd8b463` records the accepted state. All checkboxes remain unchanged.

### `E2-P2A-A003CHEAPAPP1` — Valid Frozen Cheapapp Candidate Measurement

Status: `MEASUREMENT_A003_CHEAPAPP_COMPLETE / MEASUREMENT_ONLY`.

- Finalizer task `01a03d66-baca-7353-958d-b2cefd3ca087` consumed the sole valid launch from raw root `E:\Anvien\.tmp\child06a_a003_cheapapp_frozen_17child_20260826`; launch `1`, PID `15988`, exit `0`, process wall/CPU `95.630648200 / 101.687500000 s`. No rerun occurred.
- Report: `E:\Anvien\reports\Investigation\rp_child06a_a003_cheapapp_benchmark_frozen_17child.md`.
- A002 -> A003 D001/parent/analyzer/process: `3.090914200 -> 3.447846300`, `19.040468000 -> 20.472602300`, `100.843249000 -> 93.531974900`, `136.729876000 -> 95.630648200 s`; denominator `calls=27890; files=887`.
- Packet validity: `30/30` operations, `17/17` children, `2675` exclusive intervals, overlap `0`, parent/child-sum/residual `20.472602300 / 20.449866500 / 0.022735800 s`, nonempty profile, equal denominators/workload/graph/DB/semantic/Graph JSON/stdout/stderr, and unchanged candidate-source/target boundary.
- Resource before -> candidate: start alloc `1,315,328 -> 1,315,192`; end alloc `1,055,121,264 -> 886,583,856`; max Sys `1,972,115,960 -> 1,995,295,224` bytes.
- This evidence makes no disposition or baseline promotion claim.

### `E2-P2A-A003RESTAURANT1` — Valid Frozen Restaurant Manager Candidate Measurement

Status: `MEASUREMENT_A003_RESTAURANT_COMPLETE / MEASUREMENT_ONLY`; later review/disposition lives under `REVIEW1/OWNERKEEP1`.

- Visible task `01a03d8a-d0c3-7040-9b9a-69be2d8b419b` ran exactly one launch from raw root `E:\Anvien\.tmp\child06a_a003_restaurant_manager_frozen_17child_20260826`; PID `9732`, exit `0`, process wall/CPU `101.096911900 / 110.656250000 s`.
- Report: `E:\Anvien\reports\Investigation\rp_child06a_a003_restaurant_manager_benchmark_frozen_17child.md`.
- A002 -> A003 D001/parent/analyzer/process: `9.909636600 -> 9.401585300`, `21.242055400 -> 20.850792800`, `109.339859600 -> 98.020546700`, `145.066210900 -> 101.096911900 s`; denominator `calls=86030; files=1234`.
- Packet validity: `30/30` operations, `17/17` children, `3716` exclusive intervals, overlap `0`, parent/child-sum/residual `20.850792800 / 20.834583800 / 0.016209000 s`, nonempty profile, equal denominators/workload/graph/DB/semantic/Graph JSON/stdout/stderr, and unchanged candidate-source/target boundary.
- Resource before -> candidate: start alloc `1,475,192 -> 1,474,592`; end alloc `1,119,052,736 -> 1,306,049,888`; max Sys `2,911,591,032 -> 2,634,607,224` bytes.
- Both packets remain separate and unpromoted. Next owner is one fresh visible A003 per-attempt Supervisor.

### `E2-P2A-A003REVIEW1` — Independent A003 Per-Attempt Supervisor PASS

Status: `SUPERVISOR_A003_PASS`.

- Visible Supervisor task `01a03d9b-db3b-71c2-a637-b6fccf6591ce`; report `E:\Anvien\reports\Supervisor\rp_supervisor_260826_by_gpt-5_child06a_a003_canonical_decoder.md`.
- Fresh graph `2235/766/0`, `123974/170800`; file HIGH, exact decoder impact LOW with `7` upstream symbols, `1` direct caller/file, `2` communities, `0` processes.
- Exact two-file candidate/hash/staged boundary passed. One outer-Note decoder, tuple/policy fail-closed parity, nested Authority preservation, A002 appender/write-through/lifecycle, consumed build/test truth, and both separate target equivalence/resource packets passed.
- Exact verdict: `SUPERVISOR_A003_PASS`. KEEP/no-KEEP and rollback remain Main authority.

### `E2-P2A-A003DECISION1` — Prior Main NO_KEEP And Exact Rollback History

Status: `HISTORICAL NO_KEEP / ROLLBACK_COMPLETE / SUPERSEDED_BY_OWNER_KEEP`.

- Main accepts `SUPERVISOR_A003_PASS` for accuracy/equivalence, but A003 cannot receive `KEEP`: Cheapapp D001 `3.090914200 -> 3.447846300 s` and parent `19.040468000 -> 20.472602300 s` both worsen. Restaurant improvement and both process improvements cannot override the binding all-target child/parent/process requirement.
- Candidate values remain unpromoted. Accepted A002 baselines and checkpoint `ecf825d709b761390a5df4a2147b6ed6eec04499` remain authoritative; D001 unsuccessful streak changes `0 -> 1`; P2-A, parent, D001, and D002-D017 checklist states remain unchanged.
- `E2-P2A-A003ROLLBACK1`: fresh visible rollback task `01a03db2-20ad-7200-9378-716e2f61d49f` surgically removed only A003 decoder/test hunks. Post-hashes are `AEABA8A541D1C293DFA4FE411253A85A26B6AC0DD18047AFA6F9151B4932BEE8` and `58E8CAE2C4EBBF4672D2D338CC5E1419B7CECA9F7389B0E24F51BF2F975A7CEA`; diff against `ecf825d7` is empty, scoped diff-check PASS, staged set empty, and no build/test/benchmark/analyze/detect/commit reran.
- The historical next-owner pointer to A004 is superseded. No A004 Architect was opened; current transition is the exact A003 restoration contract below.

### `E2-P2A-A003OWNERKEEP1/COMMIT1` — Owner KEEP, Exact Restoration, And Accepted Checkpoint

Status: `A003_CHECKPOINT_COMPLETE / WAL_FIX_CHECKPOINT_COMPLETE / A004_ARCHITECT_OWNER_APPROVED / A004_PLAN_READY_FOR_MAIN_VERIFY / CODER_LOCKED / D001_STREAK_0`.

- Owner authority is A003-specific: `lệch tầm 1s thì ko phải là con số đáng kể. keep được rồi`. This does not replace or weaken the measurements. Cheapapp D001 `3.090914200 -> 3.447846300 s` (`+0.356932100 s`) and parent `19.040468000 -> 20.472602300 s` (`+1.432134300 s`) remain objective measured variation, not absence, error, or automatically labeled noise.
- Contextual evidence for this one disposition remains target-separated: Cheapapp analyzer/process improve `100.843249000 -> 93.531974900 / 136.729876000 -> 95.630648200 s`; Restaurant D001/parent/analyzer/process improve `9.909636600 -> 9.401585300 / 21.242055400 -> 20.850792800 / 109.339859600 -> 98.020546700 / 145.066210900 -> 101.096911900 s`; exact independent verdict is `SUPERVISOR_A003_PASS`.
- Owner therefore corrects the prior A003-only mechanical interpretation that any single positive local delta must force `NO_KEEP`. No fixed tolerance or reusable A004+ exception is created. D001 unsuccessful streak resets `1 -> 0`; P2-A, parent, D001, D002-D017, and every checkbox remain unchanged; A004 Architect is pending.
- Materialized-source truth: `E2-P2A-A003ROLLBACK1` remains exact historical fact. Current source/test bytes equal exact final A003 hashes. Restoration, fresh analyze, required detect, and accepted checkpoint are complete.
- Recovery source: `C:\Users\TAM NGUYEN\.codex\sessions\2026\08\26\rollout-2026-08-26T15-34-20-01a03d34-b740-7982-8ff7-c2fec47940a6.jsonl`. The fresh restoration Coder first verifies A002 SHA-256 `diagnostics.go=AEABA8A541D1C293DFA4FE411253A85A26B6AC0DD18047AFA6F9151B4932BEE8` and `diagnostics_test.go=58E8CAE2C4EBBF4672D2D338CC5E1419B7CECA9F7389B0E24F51BF2F975A7CEA`, then decodes original patch payloads from 1-based lines `433`, `461`, `489` and submits each through `apply_patch` in that exact order. Prior successful application records are lines `434`, `462`, `490`. Patch bodies are not duplicated or manually reconstructed.
- Exact final identity is materialized: `diagnostics.go=6DE54D97A1B95E877B686DC3459238598F6060D28CE6329357A80BD7FC376D30`; `diagnostics_test.go=06677DBA4FE9EDA4FE8651E08B93ABC8659129A61355E34BFFA10380AE429371`.
- `E2-P2A-A003COMMIT1`: full build exited `0`, including canonical maintenance analyze `2237/766/0` and graph `123997/170823`; detect exited `0`; exact nine-path commit is `b6bf45bce95323aa6b53b182edfea8628bd8b463` with subject `perf(graphhealth): checkpoint accepted A003 canonical decoder`.
- Gate reuse: exact identity and checkpoint passed; do not rerun build, tests, benchmark, profile, target measurement, or Supervisor.
- Next owner: consumed by the completed WAL-fix Coder and Supervisor. Main creates the accepted checkpoint, then resumes A004.

### Attempt State Machine

### `E2-P2A-WALFORCEPLAN1` — Confirmed `--force` WAL-Sidecar Cleanup Fix Contract

- Reproduced on current code: `anvien analyze --force` failed `lbug_database_init failed with state 1` after lock/process clearance. The storage retained Ladybug checkpoint sidecars even when `.wal/.lock` were absent: `.wal.checkpoint`, `.checkpoint.apply.lock`, and `.checkpoint.intent.lock`.
- Root cause: `prepareStorage` omitted physical database-family cleanup, and the old runtime inventory modeled only `.wal/.lock`, omitting the three checkpoint sidecars actually produced by Ladybug.
- Approved architecture: `lbugruntime` owns exactly the five physical sidecars and full database-generation reset; `analyze` decides when `force` requires that reset and never knows suffixes. WAL recovery continues through only its explicit corruption classifier and never treats generic state `1` as recoverable WAL corruption.
- Exact source/test boundary, validation order, and STOP conditions were consumed by Coder task `01a03e29-7983-7441-8f33-8e8edf4a2493`; no glob, native wrapper/error-plumbing, non-force, A003, A004, target, or benchmark edit was authorized.
- Safe recovery/full-build evidence from Main is preserved: protected root `E:\Anvien\.tmp\lbug-recovery-20260826-a003`; canonical full build PASS including maintenance analyze `2237/766/0`, graph `123997/170823`.
- Historical cursor: consumed by `E2-P2A-WALFORCEFIX1/REVIEW1/COMMIT1`; it released the A004 Architect whose finalized result is now recorded under `E2-P2A-A004ARCH1/PLAN1`.

### `E2-P2A-WALFORCEFIX1/REVIEW1/COMMIT1` — Exact Fix, Supervisor PASS, And Checkpoint

Status: `WAL_FIX_SUPERVISOR_PASS / WAL_FIX_CHECKPOINT_COMPLETE`; accepted historical prerequisite to A004.

- Coder task/report: `01a03e29-7983-7441-8f33-8e8edf4a2493`; `E:\Anvien\reports\coder\rp_coder_260826_by_gpt-5_child06a_wal_force_cleanup.md`; verdict `CODER_CHILD06A_WAL_FORCE_FIX_READY_FOR_SUPERVISOR`.
- Exact candidate: `internal/lbugruntime/wal_recovery.go`, `internal/lbugruntime/extensions_retry_test.go`, `internal/analyze/analyze.go`, and `internal/analyze/analyze_test.go`; diff `81` insertions / `11` deletions; scoped diff-check PASS; staged set empty.
- Production ownership: `DatabaseSidecarPaths` is the sole explicit five-sidecar inventory; `RemoveDatabaseArtifacts` removes the main DB plus that inventory; `prepareStorage(force)` delegates physical cleanup and retains Graph JSON cleanup; `RunWithWALRecovery` remains explicitly classified, sidecar-only, main-DB preserving, and one-retry.
- Production-first/test evidence: production file timestamps precede both test-file changes. Tests lock the exact inventory, full-generation cleanup, sidecar-only WAL recovery, generic `lbug_database_init failed with state 1` non-recovery, and force cleanup.
- Coder validation: final canonical full build exit `0`; focused lbugruntime and analyze tests PASS; rebuilt `anvien.exe` five-sidecar repro exits `0`, reports `scanned=0`, creates main DB/Graph JSON/meta, and leaves all five seeded sidecars absent.
- Independent Supervisor task/report: `01a03e75-c596-7e80-8880-698e3bce145a`; `E:\Anvien\reports\Supervisor\rp_supervisor_260826_by_gpt-5_child06a_wal_force_cleanup.md`; exact verdict `SUPERVISOR_CHILD06A_WAL_FORCE_FIX_PASS`.
- Supervisor fresh graph: `2240` files scanned / `766` parsed / `0` failed, graph `124025 / 170891`. File risk HIGH for both production files. Exact critical impacts: `RemoveDatabaseSidecars` `6` symbols / `3` files / `12` processes; `RemoveDatabaseArtifacts` `4 / 3 / 21`; `prepareStorage` `5 / 4 / 31`. Tight source inspection, canonical build, focused tests, and independent runtime repro close these blast-radius warnings.
- Final hashes: `wal_recovery.go 3EAEAFF73B1EE27C69FBBBCB0D8D52559E210DA187A88CC613671BDD4879F490`; `extensions_retry_test.go B54866C5CF749C5CD7789C15D70B15D20FF3656842F2757CEDBB4243AB5DA5B8`; `analyze.go 0815E6E181B91490B5D47371D6862D8ABD9E68C2AC59B374CF95576E6EE40AB4`; `analyze_test.go CD40D4251CFDED2140E2D5728C5A8C3F338AC4F7964734C9D735403A86B3AA3F`; rebuilt runtime `EAC01B85E54FFA76502E122FDBF99C5221D4AE22D1CA823EAA48415D09A54A58`.
- Preserved: `E:\Anvien\.tmp\lbug-recovery-20260826-a003`, `E:\Anvien\.tmp\wal-fix-preedit-sidecars`, paused A004 report, all target measurements, A003 evidence, metrics, checkboxes, queues, and D001 streak `0`.
- Main detect: fresh graph `2241` scanned / `766` parsed / `0` failed, graph `124046 / 170912`; `anvien detect-changes --repo E:\Anvien --scope all` PASS. It records `9` tracked changed files, `79` changed graph entities, source file risk HIGH, `affected_processes=[]`, and summary risk `low`. Five reported gap occurrences are graph-resolution gaps for builtin/call sites and did not invalidate compiled/runtime behavior.
- `E2-P2A-WALFORCECOMMIT1`: exact commit `0f3a572331dd23d17688886fcbfebeb7d37ee35d`; subject `fix(storage): reset Ladybug artifact family on force`; exact `11`-path manifest is the four source/test files, four standard ledgers, Planner report, Coder report, and Supervisor report; staged diff-check PASS and staged set empty after commit. Paused A004 report remained excluded/untracked.
- A004 task `01a03e09-77e0-7c03-b144-bed526838f37` resumed immediately after the checkpoint, returned the finalized Owner-approved direction recorded below, and opened only the visible Planner translation. No Coder transition is authorized by that history.

### `E2-P2A-A004ARCH1` — Owner-Approved Export-Binding Evidence Ordering Architecture

Status: `ARCHITECT_A004_OWNER_APPROVED_READY_FOR_PLANNER`. This is architecture authority for exact Planner translation only; it is not Coder authorization, source/test/build evidence, measurement, Supervisor review, disposition, staging, or commit authority.

- Report: `E:\Anvien\reports\system-architect\rp_system-architect_260826_by_gpt-5_child06a_a004_residual_direction.md`; current report/source basis HEAD `54b2bd9b8b10b84861abd81f1686874cbd23c16c`; accepted A003 checkpoint `b6bf45bce95323aa6b53b182edfea8628bd8b463`; accepted WAL checkpoint `0f3a572331dd23d17688886fcbfebeb7d37ee35d`.
- Current hierarchy is unchanged: active unchecked parent `B1-P1A-OP001 resolution`, active unchecked child `B2-P2A-A001-D001 resolve_calls`, D001 streak `0`, D002-D017 queued/unopened, and Cheapapp/Restaurant Manager accepted A003 bases remain separate in benchmark.
- Proven residual: `appendExportBindingEvidence` serializes export-binding notes, sorts the projection, then calls `mergeExportBindingEvidence`; the merge exact-tuple deduplicates/partitions and sorts projected evidence again. `sortExportBindingEvidence` uses a stable comparator that repeatedly invokes `exportBindingEvidenceOrderFor`, which decodes final JSON `Evidence.Note` for terminal/hop/failure records. The same stack appears beneath D001 in both retained target profiles; CPU samples are overlapping causal evidence only and predict no numeric gain.
- Complete primary path: `internal/analyze.Run -> runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> w.files -> ir.Calls -> resolveCall -> appendExportBindingEvidence -> mergeExportBindingEvidence -> sortExportBindingEvidence -> comparator -> exportBindingEvidenceOrderFor -> json.Unmarshal`. Shared preserve/run-only consumers include resolved relationship coalescing through `emitReference/emitRelationship/mergeRelationship`, final resolution-outcome/reference-index projection, and sibling `resolveAccess`.
- Locked direction: `exact-tuple dedupe -> derive the canonical order key from final serialized Evidence.Note exactly once per unique projected export-binding evidence record -> exactly one final stable sort using cached keys -> discard transient keys -> return byte/order-identical []graph.Evidence`. Generic evidence retains first-seen order; projected comparison retains exact kind rank, proof ordinal, hop ordinal, and lexical Note tuple.
- Allowed production surface is only `internal/resolution/export_binding_proof.go`: `appendExportBindingEvidence`, `mergeExportBindingEvidence`, `sortExportBindingEvidence`, `exportBindingEvidenceOrderFor`, existing private order/key types, and at most one private transient decorated representation/helper. Existing signatures remain unchanged. Only after production is correct may A004-specific tests be appended to `internal/resolution/export_binding_proof_test.go`.
- Forbidden architecture: producer-carried typed alternate keys/private overloads, JSON decode/marshal/validation/mutation in the comparator, more than one final projected-evidence sort, public/persisted order fields, run-wide/cross-call/global cache, concurrency/I/O/lifecycle ownership, another production file, A003 decoder/appender changes, or D002-D017 implementation scope.
- Blast radius: containing file HIGH. Exact upstream impacts are CRITICAL: `appendExportBindingEvidence` `36` symbols / `5` direct callers / `14` files / `7` modules / `39` processes; `mergeExportBindingEvidence` `51 / 5 / 19 / 7 / 49`; `sortExportBindingEvidence` `23 / 2 / 10 / 2 / 34`; `exportBindingEvidenceOrderFor` `14 / 1 / 6 / 1 / 19`. These are scope warnings requiring the one-file boundary and broad run-only regressions, not edit bans.
- Exact parity preserves byte-for-byte `Kind/Weight/Note`, full slice length/content/order, generic first-seen order, exact-tuple dedupe, terminal/hop/failure/default rank, ordinal and lexical ties, malformed-note max-ordinal fallback, stable equal keys, SourceSiteIDs, repeated-coalescing idempotency, marshal/parse failure behavior, input non-mutation/aliasing, relationship/reference/outcome/diagnostic carriage, graph/output/persistence/readers, determinism/freshness/failure/transaction/temp/publication behavior.
- Transient resource boundary is private `O(P)` per merge for `P` unique projected records; it may copy evidence values/string headers but not duplicate retained Note byte storage, and no state survives the call.
- Tests require a test-local pre-A004 oracle, never a production fallback, comparing full ordered `[]graph.Evidence` equality across terminal/hop/failure proofs/hops, mixed generic/projected inputs, permutations/repeated coalescing, duplicates, ordinal ties/different Notes, multiple SourceSiteIDs, malformed JSON for each export kind, unknown/non-export kinds, equal-key stability, empty/no-proof paths, input non-mutation, and idempotency.
- Validation/measurement: fresh future Coder graph/file-detail/exact impacts for every edited helper; production first; tests second; holder clearance; canonical full build PASS before tests; differential/export-binding/call/access/outcome/Graph JSON regressions and full `internal/resolution` with the known golden truthful; unchanged A00x contract; separate Cheapapp and Restaurant Manager A003-to-A004 packets covering D001/parent/analyzer/process, `30/30`, `17/17`, denominators, conservation, evidence order, Graph JSON, stdout/stderr, graph/DB/semantics, and resources; one fresh A004 Supervisor only after both packets.
- Exact rollback owns only the A004 production hunk/private transient helper/type and A004-appended test bytes. Mandatory STOP applies on expanded production ownership, alternate key authority, comparator JSON decoding, second final sort, retained cache/concurrency, malformed evidence rewrite/drop, changed contract/output, D002-D017 opening, unavailable independent packets, or unprovable exact parity.

### `E2-P2A-A004PLAN1` — Exact Visible Planner Translation

Status: `A004_PLAN_READY_FOR_MAIN_VERIFY / CODER_LOCKED`.

- The living Current P2-A Attempt card and P2-A work steps now bind `E2-P2A-A004ARCH1` exactly: one-file production owner, dedupe-before-decoration, one canonical final-Note key extraction per unique projected record, one cached-key stable sort, unchanged ordered output, production-first tests, full validation/measurement/Supervisor order, exact rollback, and mandatory STOP.
- Plan authoring gate: `anvien --help` PASS; one fresh `anvien analyze --force` exited `0` with `2241` scanned / `766` parsed / `0` failed and graph `124071 / 170937`. Indexed/current commit is `54b2bd9b8b10b84861abd81f1686874cbd23c16c`, stale `false`.
- File-detail after that refresh passed for all four standard ledgers. Each is parsed/current/non-stale LOW-risk documentation with `2` inbound references, `1` outbound reference, `0` unresolved sites, and exactly `2` related files (`plan-rules.md` and the roadmap). This proves plan freshness/traceability only.
- Exact Planner report: `E:\Anvien\reports\planner\rp_planner_260826_by_gpt-5_child06a_a004_export_evidence_ordering.md`.
- `benchmark.md` remains the sole numeric source. Existing Cheapapp and Restaurant Manager A003 bases, denominators, rows, and final-current values are unchanged; this translation creates no A004 value or performance claim.
- P2-A, parent, D001, and D002-D017 checkboxes remain unchanged; D001 streak remains `0`; queue/order and accepted baselines remain unchanged. No A001/A002/A003/WAL gate was rerun or reinterpreted.
- Planner performed no source/test/script/target/Architect-report edit, source impact, detect, build, test, benchmark capture, profile, target analyze, Supervisor, stage, or commit. The architecture report remains outside Planner write ownership.
- Next owner is Main Orchestration for independent verification of the exact five owned artifacts and a separate future Coder-release decision. Planner neither opens nor commands Coder.

### `E2-P2A-A004MAINVERIFY1` — Main Verification And Sole Coder Release

Status: `A004_PLAN_COMPLETE / A004_MAIN_VERIFIED / A004_CODER_ACTIVE`.

- Main independently checked the approved Architect report, all four living ledgers, and the visible Planner report. Required A004 algorithm, one-file production boundary, tests-after-production rule, separate target contract, rollback/STOP boundary, unchanged checklist/streak/queue, and `E2-P2A-A004ARCH1/PLAN1` pointers are consistent.
- Verification report: `E:\Anvien\reports\Supervisor\rp_supervisor_260826_232341_by_gpt-5_child06a_a004_planner_translation.md`; verdict `PASS`; scoped diff-check passed, staged set was empty, and checklist-line delta count was `0`.
- Main released sole visible Coder task `01a03ee4-3c2a-72d0-afe7-08be98fe982e`. Its ownership is limited to `internal/resolution/export_binding_proof.go`, post-production A004 bytes in `internal/resolution/export_binding_proof_test.go`, and its Coder report.
- Target measurement, post-measurement Supervisor, disposition, detect, staging, and implementation commit remain locked until Coder returns valid source/build/test evidence.

### `E2-P2A-A004IMPACT1/SRC1/BUILD1/TEST1` — Coder Candidate Ready For Measurement

Status: `CODER_A004_READY_FOR_MEASUREMENT_HANDOFF / MAIN_HANDOFF_PASS / A004_MEASUREMENT_READY`.

- Coder task/report: `01a03ee4-3c2a-72d0-afe7-08be98fe982e`; `E:\Anvien\reports\coder\rp_coder_260826_by_gpt-5_child06a_a004_export_evidence_ordering.md`; report SHA-256 `52C62354D4F3D90BB4C6C922C77E1CA51E3E92169B86F580938625E00301A620`.
- `A004IMPACT1`: fresh graph `2243/766/0`, `124096/170962`; file HIGH/current; helper impacts CRITICAL with exact fresh counts recorded in the Coder report.
- `A004SRC1`: production only `internal/resolution/export_binding_proof.go`, `+18/-4`, SHA-256 `36D41BC7336A5A04B4C77D1EE9DEB45DE4E5E2BA120EEFBF98419BD2D846C4D2`; tests only `internal/resolution/export_binding_proof_test.go`, `+477/-0`, SHA-256 `99C6F6FC5FD0AE4D1BFDFE547D5F67C958544AA61FFD89895DC0BAC79C839BBC`.
- `A004BUILD1`: holder preflight found no proven build-output holder; canonical `scripts/full-build.ps1` PASS before tests; runtime SHA-256 `7C077F8555ABDD9D1A76ED19931B448B208622A3B13F170BB71131A463E826F6`; maintenance analyze `2243/766/0`, graph `124258/171186`.
- `A004TEST1`: A004 differential/export-binding, focused call/access/outcome, graphhealth/Graph JSON, and analyzer carriage suites PASS. Full `internal/resolution` truthfully exits `1` only on the identical recorded `TestProofBasedCallAccessGoldenCorpus`; no new failure and no golden edit.
- Main verification report `E:\Anvien\reports\Supervisor\rp_supervisor_260826_235532_by_gpt-5_child06a_a004_coder_handoff.md` returns PASS for measurement handoff only. Candidate remains uncommitted/unstaged; target measurement, Supervisor disposition, detect, and implementation commit remain pending.

### `E2-P2A-A004PACKET1` — Frozen Candidate And Independent Measurement Launches

Status: `A004_FROZEN_PACKET_READY / CHEAPAPP_MEASUREMENT_ACTIVE / RESTAURANT_MEASUREMENT_ACTIVE`.

- Build report: `E:\Anvien\reports\Investigation\rp_child06a_a004_frozen_candidate_build.md`; reusable script ran exactly once, exit `0`, marker `A00X_BENCHMARK_BUILD_COMPLETE`; no target analyze occurred in the build lane.
- Frozen executable `E:\Anvien\.tmp\child06a_a004_a00x_benchmark_build\anvien-a004-benchmark.exe`: `73,825,280` bytes, SHA-256 `6D319467D198B8BBA2375339CDF6BD7634FA97E7C503D3DFC0C8C315D965352C`, version `1.2.8`.
- Frozen DLL SHA-256 `20CBD87840483A2053CFF3FC2DB23A86DD802B8915D86509D41A4B709624CDB7`; provenance SHA-256 `4EBF871D0008B6BB1A7FDB433DE07CE437F859DFB34221C94D360E866711A20D`, schema/attempt/mappings/candidates/native/exit `1/A004/2/4/3/0`; all expected/actual hashes match.
- Cheapapp one-launch task `01a03f0a-8a2c-71a3-a889-428daf219ba7` owns only `E:\cheapapp.org`, raw root `E:\Anvien\.tmp\child06a_a004_cheapapp_frozen_17child_20260827`, and report `reports/Investigation/rp_child06a_a004_cheapapp_benchmark_frozen_17child.md`.
- Restaurant one-launch task `01a03f0b-1eba-78b1-8845-9ca92bb58d29` owns only `E:\Restaurant_manager`, raw root `E:\Anvien\.tmp\child06a_a004_restaurant_manager_frozen_17child_20260827`, and report `reports/Investigation/rp_child06a_a004_restaurant_manager_benchmark_frozen_17child.md`.
- Each lane preserves its own accepted A003 before packet and cannot retry, average targets, decide disposition, edit ledgers/source/test, run Supervisor/detect, stage, or commit. Post-measurement Supervisor remains locked until both valid packets exist.

### `E2-P2A-A004CHEAPAPP1/RESTAURANT1` — Complete Separate Target Measurements

Status: `BOTH_TARGET_MEASUREMENTS_RECORDED / SUPERVISOR_PENDING`; numbers are unpromoted measurement input only.

- Cheapapp task/report: `01a03f0a-8a2c-71a3-a889-428daf219ba7`; `E:\Anvien\reports\Investigation\rp_child06a_a004_cheapapp_benchmark_frozen_17child.md`; one launch, exit `0`, `30/30`, `17/17`, `calls=27890; files=887`, exact conservation and output/order equivalence. A003 -> A004: D001 `3.447846300 -> 2.074182500 s` (`-1.373663800 s`, `-39.841213%`); parent `20.472602300 -> 13.265999200 s` (`-7.206603100 s`, `-35.201207%`); analyzer `93.531974900 -> 107.287054400 s` (`+13.755079500 s`, `+14.706286%`); process `95.630648200 -> 144.975972400 s` (`+49.345324200 s`, `+51.599906%`).
- Restaurant task/report: `01a03f0b-1eba-78b1-8845-9ca92bb58d29`; `E:\Anvien\reports\Investigation\rp_child06a_a004_restaurant_manager_benchmark_frozen_17child.md`; one launch, exit `0`, `30/30`, `17/17`, `calls=86030; files=1234`, exact conservation and output/order equivalence. A003 -> A004: D001 `9.401585300 -> 8.975767700 s` (`-0.425817600 s`, `-4.529211%`); parent `20.850792800 -> 19.416099500 s` (`-1.434693300 s`, `-6.880761%`); analyzer `98.020546700 -> 101.406172300 s` (`+3.385625600 s`, `+3.453996%`); process `101.096911900 -> 135.569489100 s` (`+34.472577200 s`, `+34.098546%`).
- Both targets preserve exact files/parser/graph/DB/dependency/projection/semantic/diagnostic/outcome identities; canonical graph.json, stdout, and stderr hashes equal each target's accepted A003 packet. Canonical graph byte identity proves full ordered Evidence arrays are unchanged.
- Resource observations remain separate: Cheapapp end allocation and max observed system memory fall; Restaurant end allocation falls while max observed system memory rises. Secondary resources cannot override elapsed disposition.
- No target is averaged or used to subsidize the other. No Supervisor verdict or disposition exists yet; accepted A003 baselines and D001 streak `0` remain unchanged.
- Sole fresh Supervisor task `01a03f21-63a6-7c53-a793-bb7cd459b858` is active for exact A004 candidate correctness/equivalence/output/lifecycle only. It cannot decide performance disposition.

### `E2-P2A-A004REVIEW1/DECISION1` — Supervisor PASS And Main NO_KEEP

Status: `SUPERVISOR_A004_PASS / NO_KEEP / ROLLBACK_ACTIVE / D001_STREAK_1`.

- Supervisor task/report: `01a03f21-63a6-7c53-a793-bb7cd459b858`; `E:\Anvien\reports\Supervisor\rp_supervisor_260827_by_gpt-5_child06a_a004_export_evidence_ordering.md`; exact verdict `SUPERVISOR_A004_PASS` for correctness/equivalence/output/lifecycle only.
- Fresh review graph `2248/766/0`, `124309/171237`; file HIGH; current impacts append/merge/sort CRITICAL and extractor LOW. Source/test/frozen/packet identities and staged-empty boundary match exactly.
- Main disposition is `NO_KEEP`: the plan requires lower selected child, parent, and process elapsed on each target plus Supervisor PASS. A004 lowers D001 and parent on both targets but process increases Cheapapp `+49.345324200 s` and Restaurant `+34.472577200 s`; analyzer is also higher on both. No A004 exception or tolerance exists.
- Accepted A003 target-separated values remain current; A004 values are not promoted. D001 consecutive unsuccessful-attempt streak changes `0 -> 1`; P2-A, parent, D001, D002-D017, and every checkbox remain unchanged.
- Exact rollback owns only the `+18/-4` A004 production hunk and `+477/-0` A004 test bytes. Passed build/test/measurement/Supervisor gates remain valid evidence and must not be rerun solely for rollback identity.

### `E2-P2A-A004ROLLBACK1/A005CURRENT1` — Accepted Baseline Restored And Next Basis Locked

Status: `A004_ROLLBACK_COMPLETE / A005_CURRENT_BASIS_RECORDED / RESIDUAL_ATTRIBUTION_PENDING / ARCHITECT_LOCKED / D001_STREAK_1`.

- Same Coder task `01a03ee4-3c2a-72d0-afe7-08be98fe982e` mechanically reversed only A004 production/test bytes with `apply_patch`; no checkout/reset/restore/stash/whole-file overwrite.
- Final `internal/resolution/export_binding_proof.go` SHA-256 `4BFE1976766FBF2E2257102070FF43CAC6D757E32E71DD583F8824781EAB6A8E`; final `internal/resolution/export_binding_proof_test.go` SHA-256 `AD3DCA9E82EACFB31137560636B59D426A063AEB967613E724D5EE3017AD5812`; both have zero diff against HEAD. Coder report appendix SHA-256 `2C93A4A5A2CC0240DE2EFF7F3F413C03EB15DCCF424DB59174990595019A7D59`.
- Frozen and raw A004 artifacts remain preserved as rejected-candidate evidence. No Anvien/build/test/measurement/Supervisor/detect/stage/commit gate was rerun for rollback.
- A005 current numeric basis is the accepted target-separated A003 row; parent/D001 remain active/unchecked, D002-D017 queued, and streak is `1`.
- A005 Architect is locked until a separate visible read-only attribution lane proves current residual cause, exact owner, and complete call path from both retained A004 CPU profiles plus current restored graph/source.
- That sole attribution lane is active as task `01a03f3f-34e5-7032-9944-3509d3cf39cf`; it may write only `reports/Investigation/rp_child06a_a005_residual_attribution.md` and cannot select architecture or edit source/ledgers.

### `E2-P2A-A005ATTRIB1` — Record-Time Outcome Serialization Residual

Status: `A005_ATTRIBUTION_COMPLETE / A005_ARCHITECT_COMPLETE / A005_PLAN_READY_FOR_MAIN_VERIFY / CODER_LOCKED`.

- Attribution task/report: `01a03f3f-34e5-7032-9944-3509d3cf39cf`; `E:\Anvien\reports\Investigation\rp_child06a_a005_residual_attribution.md`; clean authoritative HEAD `6ee70e29`; one fresh graph `2249/766/0`, `124165/171031`.
- Selected cause: `(*resolutionOutcomeCollector).record` eagerly calls `marshalResolutionOutcome/json.Marshal` for every accepted outcome. Repository-resolved callers discard the returned encoded bytes; resolved-external TypeScript may also not consume them; later `projectResolutionOutcomes` serializes every finalized retained outcome again for canonical graph/reference/diagnostic carriers.
- Exact owner: `internal/resolution/outcome.go::(*resolutionOutcomeCollector).record` lines `83-106`, sampled unique path line `100`; helper `marshalResolutionOutcome` lines `220-226`, `json.Marshal` line `221`.
- Four retained profiles are separately pinned. Under D001, `record/marshal` remains sampled in Cheapapp A003/A004 and Restaurant A003/A004; CPU samples are overlapping causal evidence only. A004 profiles expose the family after removing the rejected export-ordering work, but A004 values remain rejected.
- Complete call path: `cli analyze -> analyze.Run -> PhaseResolution -> ResolveBoundInto -> resolveCall -> repository-resolved/repository-unresolved/TypeScript recorders -> collector.record -> marshalResolutionOutcome -> finalize -> projectResolutionOutcomes -> relationship/reference/diagnostic carriers -> analyze.Result -> graph consumers -> Ladybug DB/Graph JSON -> CLI consumers`.
- Current graph: containing file HIGH (`119` symbols, `87` inbound, `122` outbound, `11` linked flows, `23` linked tests); file impact CRITICAL (`177` symbols, `55` files). Method `record` CRITICAL (`5` impacted, `5` direct, `10` processes, `2` files). Helper `marshalResolutionOutcome` CRITICAL (`127` impacted, `3` direct, `48` files, `25` modules, `52` processes). These are warnings, not edit permission.
- Architect boundary: decide ownership/lifecycle of record-time encoded bytes versus immediate unresolved/non-resolved diagnostic needs and final projection serialization; preserve clone/validation/conflict/error timing, SourceSiteID/order, byte-identical Diagnostic/Evidence Notes, exact graph/output/persistence/readers, and bounded private resources. Attribution chooses no implementation.
- A004 direction is excluded: it passed correctness but failed process elapsed on both targets, was `NO_KEEP`, and was rolled back. A001-A003 remain accepted preserve-only directions.
- Fresh visible Architect task `01a03f5e-b844-7ac3-b23b-b9c9cdc0374d` completed its A005-only direction; the visible Planner translation is now recorded under `E2-P2A-A005PLAN1`, and Coder remains locked.

### `E2-P2A-A005ARCH1` — Canonical Outcome-Byte Ownership Architecture

Status: `A005_ARCHITECT_COMPLETE / A005_PLAN_READY_FOR_MAIN_VERIFY / CODER_LOCKED`.

- Architect task/report/verdict: `01a03f5e-b844-7ac3-b23b-b9c9cdc0374d`; `E:\Anvien\reports\system-architect\rp_system-architect_260827_by_gpt-5_child06a_a005_outcome_serialization.md`; `ARCHITECT_A005_READY_FOR_PLANNER`.
- Fresh Architect graph passed at HEAD `0030c97bcd0fc6f94d4146842e6b0d21d8ac3f1d`: `2250/766/0`, graph `124189/171055`; report-only diff check passed and staged set was empty.
- Direction: encode each retained outcome exactly once at record time, retain the canonical JSON string in one private run-scoped `SourceSiteID -> canonical JSON` sidecar, and share it with immediate diagnostics and final projection. `projectResolutionOutcomes` is a strict consumer/validator with no re-encoding fallback.
- Exact architecture ownership is private outcome state/functions in `internal/resolution/outcome.go` plus only the existing finalize/project/result wiring block in `internal/resolution/resolve.go`; tests follow production in one new focused file plus minimal direct-private-call adaptation in the existing P6C3 test. Public/persisted shapes, A001-A004, graph/output/persistence/readers, workloads, scripts, and target separation remain preserve-only.
- Main handoff verification report `E:\Anvien\reports\Supervisor\rp_supervisor_260827_021244_by_gpt-5_child06a_a005_architect_handoff.md` records `PASS`. It released exactly one visible Planner to translate the report without redesign; `E2-P2A-A005PLAN1` records that completed translation and still does not release Coder.

### `E2-P2A-A005PLAN1` — Exact Visible Planner Canonical Outcome-Byte Translation

Status: `A005_PLAN_COMPLETE / A005_MAIN_VERIFIED / CODER_RELEASE_PENDING`.

- Planner role/basis: one visible A005 Planner in shared checkout `E:\Anvien`, saved project directly, at checkpoint HEAD `e6ab44705b9abb4d1a55094cfb1a2bce1f06ea8d`. It read `AGENTS.md`, `working-rules`, `planner`, binding `plan-rules.md`, all four standard ledgers, the full A005 Architect report, and Main handoff verification through EOF in the required order; `anvien --help` passed. It consumed the fresh Architect graph/impact packet and did not rerun graph attribution, architecture, target analysis, build, tests, measurement, Supervisor, detect, stage, or commit.
- Planner report: `E:\Anvien\reports\planner\rp_planner_260827_by_gpt-5_child06a_a005_outcome_serialization.md`; verdict `PLANNER_A005_READY_FOR_MAIN_VERIFY`.
- Locked byte lifecycle: retain the existing cloned semantic map and add exactly one private run-scoped `SourceSiteID -> canonical JSON string` sidecar. A unique valid `record` retains the clone, invokes unchanged `marshalResolutionOutcome` once at the existing record boundary, stores bytes only on success, and preserves its signature/returns. Equal duplicates reuse the stored immutable clone/string with `added=false`; conflicts and record-time marshal failures preserve current first-error behavior and never retry. One private `finalizedResolutionOutcomes` bundle carries SourceSiteID-sorted cloned values plus the sealed sidecar. `projectResolutionOutcomes` strictly consumes/validates it with no marshal, fallback, second encoder, or missing-byte tolerance; `ResolveBoundInto` passes the bundle to projection and unwraps only its values for public `Result.ResolutionOutcomes`.
- Exact production owners: only private collector state/init, `record`, the private finalized bundle, `finalize`, and `projectResolutionOutcomes` in `internal/resolution/outcome.go`; plus only the existing `finalize -> project -> error/result ResolutionOutcomes` wiring block in `internal/resolution/resolve.go`. `marshalResolutionOutcome` body/signature, validation/clone, outcome constructors, immediate diagnostic callers/emitters, graph/public/persistence/readers, instrumentation, scripts, A001-A004, and D002-D017 are preserve-only.
- Production first, tests second: after production is correct, create only `internal/resolution/outcome_serialization_test.go` and mechanically adapt only three direct private `finalize` calls plus one direct `projectResolutionOutcomes` call in `internal/resolution/p6c3_structured_outcome_test.go`. The matrix proves one canonical string across repository/intrinsic/TypeScript resolved and non-resolved families; exact immediate diagnostic/final relationship/reference byte parity; equal duplicate/conflict/first-error/`added`; deep clone immutability; record-time non-finite-weight marshal failure; missing/extra sidecar coverage failure without re-encoding; duplicate finalized sites, resolved/unresolved overlap, diagnostic/outcome and reference-carrier drift rejection; SourceSiteID order, complete ordered evidence, graph/reference traversal, and repeated-run determinism. Test-local expected bytes are never a production fallback.
- Fresh future Coder gate: after Main independently verifies the four ledgers plus Planner report and separately releases one Coder, run `anvien --help`, exactly one fresh `anvien analyze --force`, file-detail for both production files, and exact upstream impact for every edited symbol. Record the Architect HIGH/CRITICAL blast radius as scope warnings. Any additional owner is a STOP before tests.
- Holder/build/test order: after production and authorized tests, run `anvien doctor locks --repo E:\Anvien --json` and `anvien doctor processes --json`, terminate only proven build-output holders, then require canonical `scripts/full-build.ps1` exit `0`. Only after PASS run A005 focused tests plus `TestP6C3P5ProofNestingConflictAndImmutableReplay`, `TestP6C3FinalOutcomeStatusPrecedenceAndCarriageMatrix`, `TestResolveAttachesSourceBackedUnresolvedDiagnostics`, `TestP6BResolverCarriesCatalogValidationFailuresPerSourceSite`, `TestP6DStructuredOutcomeProjectionPreservesGraphJSONAndHealthParity`, `TestP6C3AnalyzeResultPreservesFinalOutcomesAndGraphCarriage`, `TestP6C3AnalyzeCapabilityOutcomesRetainAcceptedAuthorityStatus`, `TestP6C3ResolutionOutcomesPreserveGraphJSONAndLadybugParity`, and `TestP6C3NativeLadybugResolutionOutcomeReadback`; then run package regressions for `internal/resolution`, `internal/graphhealth`, `internal/analyze`, `internal/lbugload`, and `internal/lbugnative`. The unchanged preserve-only golden is never called PASS; any new or changed failure blocks A005. Scoped diff-check follows; detect/stage/commit remain locked.
- Candidate and separate measurements: reuse unchanged `scripts/build-a00x-benchmark.ps1`; mechanically regenerate/verify the overlay-mapped A005 `resolve.go` from exact production wiring plus accepted 17-child instrumentation, keep `types.go` unchanged absent exact compilation evidence, and exclude rejected A004 bytes. Cheapapp and Restaurant each run once against its own accepted A003 packet with unchanged options/exclusion. Each packet records D001/parent/analyzer/process, `30/30`, `17/17`, calls/files, conservation/zero overlap, workload, diagnostics/outcomes, full ordered evidence, Graph JSON, graph/DB readback, stdout/stderr, semantic counters, allocation/System fields, and private `O(U+B)` one-payload lifecycle proof. Accepted bases are never rerun, averaged, or combined.
- Later boundary: only after both packets exist may one fresh visible A005 Supervisor review exact correctness/equivalence/output/lifecycle/resource behavior. Main alone decides `KEEP`, `REWORK`, `ROLLBACK`, or streak effect. Current streak remains `1`; one unsuccessful A005 can move it only to `2`, never terminal `3`.
- Exact rollback/STOP: rollback only the sidecar/init, `record` reuse hunk, private finalized bundle/`finalize`, projection consumer change, narrow `ResolveBoundInto` wiring, new focused test, mechanical P6C3 adaptations, and rejected A005 frozen/overlay packet. STOP for any other owner/symbol; change to the marshaler/callers/schema/public/persisted shape; moved validation or initial marshal timing; changed conflict/first-error/`added`/clone/SourceSiteID behavior; immediate byte drift; fallback/second encoder; A004 ordering drift; resource growth beyond `O(U+B)` or payload duplication/cross-run state; build/new-test/golden change; stale/mixed overlay; missing/incomparable target packet; unexplained sibling/resource regression; or D002-D017/P3/Child 07 opening.
- State effect: no benchmark number, baseline, checkbox, queue, or streak changed; accepted A003 target separation remains authoritative; A004 remains rejected/rolled back. Main verification below accepts this translation; one visible Coder may be released next.
- Authoring verification: HEAD remains `e6ab44705b9abb4d1a55094cfb1a2bce1f06ea8d`; changed-path set equals exactly the four standard ledgers plus the Planner report; scoped `git diff --check` exit `0`; report trailing-whitespace count `0`; all `53` plan checkbox lines match HEAD; all `1,470` benchmark elapsed-value tokens match HEAD; Attempt Numeric History matches HEAD exactly; staged set is empty. No detect, stage, or commit ran.

### `E2-P2A-A005MAINVERIFY1` — Main Verification And Coder Release Authority

Status: `A005_PLAN_COMPLETE / A005_MAIN_VERIFIED / CODER_RELEASE_PENDING`.

- Main reviewed the real Planner claim against `AGENTS.md`, working-rules, orchestration, planner, supervisor, binding `plan-rules.md`, `E2-P2A-A005ATTRIB1/ARCH1`, the Architect report, and Main Architect-handoff report.
- Exact reviewed set: four standard Child 06A ledgers plus `E:\Anvien\reports\Planner\rp_planner_260827_by_gpt-5_child06a_a005_outcome_serialization.md` at HEAD `e6ab44705b9abb4d1a55094cfb1a2bce1f06ea8d`; Planner task `01a03f86-cd96-7573-97b1-52d9548768fb` returned `PLANNER_A005_READY_FOR_MAIN_VERIFY`.
- Independent result: exact five-path boundary; scoped diff-check exit `0`; staged set empty; `53/53` plan checkbox lines exact; `1,470/1,470` benchmark decimal elapsed tokens exact and ordered; complete Attempt Numeric History exact; accepted/rejected bases, D001 streak `1`, target separation, queue, architecture, owner/test/build/measurement/Supervisor/rollback/STOP contracts all preserved.
- Verification report `E:\Anvien\reports\Supervisor\rp_supervisor_260827_025255_by_gpt-5_child06a_a005_planner_handoff.md` returns `PASS`. This is plan-handoff acceptance only, not source/build/test/performance acceptance.
- State effect: `A005_PLAN_READY_FOR_MAIN_VERIFY / CODER_LOCKED -> A005_PLAN_COMPLETE / A005_MAIN_VERIFIED / CODER_RELEASE_PENDING`; exactly one visible Coder may open from this authority and must begin with its own fresh Anvien/file-detail/impact gate. Measurement, later Supervisor, disposition, detect, and implementation commit remain locked.

| From | Required evidence | To / next owner |
|------|-------------------|-----------------|
| complete top-level list | `E1-P1A-RANK1` plus exact parent checklist mirror | select largest unchecked parent and measure its complete child list |
| complete parent drill-down | `AnnnDRILL1`, all dynamic child rows, and equal nested checklist cardinality | select largest unchecked child and record `AnnnCURRENT1` |
| selected child current basis | `AnnnCURRENT1` | new Visible Architect receives exact parent, complete child list, selected child/cause/owner/call path |
| Architect decision complete | `AnnnARCH1` | Planner refreshes current P2-A attempt |
| A001 Planner refresh written | `E2-P2A-A001PLAN1` | Owner-direct exception required Main verification and independent Architect consistency review before any Coder transition |
| A001 independent consistency review passes and Main decides transition | `E2-P2A-A001CONSISTENCY1` | sole Coder completes the fresh pre-edit gate, production implementation, authorized test edits, build-holder preflight, canonical full build, and only then focused/package test execution |
| Planner refresh complete for an attempt without an additional active Owner-specific gate | `AnnnPLAN1` | Coder may edit only refreshed surfaces |
| Coder production/build/test validation complete | `E2-P2A-A001IMPACT1/SRC1/BUILD1/TEST1` plus both recorded target measurement reports | both target blocks recorded separately; A001 then proceeded through `E2-P2A-A001REVIEW1/DECISION1` without an Architect/Coder reopen |
| measurements/equivalence complete | `AnnnDIRECT1/PARENT1/E2E1/EQUIV1` | new Visible Supervisor |
| Supervisor `PASS`, child/parent/end-to-end elapsed times lower | `AnnnREVIEW1`, `AnnnDECISION1` | `KEEP`; promote baseline, reset active-child streak to `0`, remeasure, and open another attempt on the same unchecked child |
| A001 `KEEP` recorded | `E2-P2A-A001REVIEW1`, `E2-P2A-A001DECISION1`; two repo-specific accepted baselines remain separate | cursor `A002 / ARCHITECT_PENDING` on the same unchecked D001; Main opens a fresh visible Architect, never Coder directly |
| A001 implementation commit recorded | `E2-P2A-A001COMMIT1`; exact commit/subject, 12-path manifest, staged set empty | `A001_COMMIT_COMPLETE`; cursor remains `A002 / ARCHITECT_PENDING`, with no A002 evidence or Coder authorization |
| A002 architecture and Planner refresh recorded | `E2-P2A-A002ARCH1`, `E2-P2A-A002PLAN1`; exact report/commit, cause, three production owners, one post-production test owner, invariants, validation/resource/measurement/rollback/STOP boundary | `A002 PLAN_READY_FOR_MAIN_VERIFY`; Main verifies the four ledgers and alone decides the next governance transition; Coder remains locked |
| A002 source/build/test validation complete | `E2-P2A-A002IMPACT1/SRC1/BUILD1/TEST1`; exact four-file candidate, definitive full-build PASS before tests, focused/graphhealth PASS, identical baseline-overlay proof for the preserve-only resolution golden, staged set empty | `A002 CODER_COMPLETE / TWO_TARGET_MEASUREMENT_PENDING`; Main dispatches independent Cheapapp then Restaurant candidate measurements without overlap or aggregation |
| Cheapapp invocation omits the complete child collection | `E2-P2A-A002MEASUREFAIL1`; launch `1`, exit `0`, `30/30` operations, `0/17` children, D001 unavailable | keep A001 baseline/streak/checklists unchanged; verify the accepted A001 optimized `resolve.go`/`types.go` instrumentation overlay and run exactly one corrected Cheapapp capture; Restaurant remains queued |
| Corrected Cheapapp measurement-binary build cannot find `lbug.h`, and the supported scripts expose no overlay input | `E2-P2A-A002MEASUREFAIL2/BUILDRECOVERY1/SCRIPTPLAN1`; build attempt `1`, exit `1`, no binary, analyze launch `0`; reusable script plan bound | keep A001 baseline/streak/checklists unchanged; preserve failed root; consume `SCRIPTBUILD1`; Restaurant remains queued |
| Owner-directed reusable A00x script validated for A002 | `E2-P2A-A002SCRIPTPLAN1/SCRIPTBUILD1`; explicit overlay/output/hash inputs, canonical vendor/native/CGO/offline flags, `.tmp` binary/DLL/provenance, canonical scripts preserve-only, unchanged-contract reuse rule | `A00X_SCRIPT_BUILD_COMPLETE / CORRECTED_CHEAPAPP_MEASUREMENT_PENDING`; run one corrected Cheapapp capture; later A00x reuses the unchanged script unless an exact new feature/input failure proves a narrow script-contract change is needed |
| Valid Cheapapp frozen candidate packet independently recorded | `E2-P2A-A002CHEAPAPP1`; one launch/exit `0`, frozen executable identity, `30/30` operations, `17/17` children, denominator `27890/887`, `2675` intervals/zero overlap, conservation, workload/graph/DB/output/semantic/resource/source-status proof | keep values unpromoted and accepted A001 baseline/streak/checklists unchanged; Restaurant recovery/measurement follows |
| Handoff Restaurant task unavailable | `E2-P2A-A002RESTAURANTBLOCK1`; exact read/unarchive/list/session evidence proves ID `01a033a5-4ec3-7e42-8946-0ab9172f6088` was not callable | historical blocker resolved by replacement task `01a03c72-5f6d-7fd0-948a-88b06a1ebb65`; no product-state effect |
| Valid Restaurant frozen candidate packet independently recorded | `E2-P2A-A002RESTAURANT1`; one launch/exit `0`, identical frozen executable identity, exact one userApi exclusion, `30/30` operations, ordered `17/17` children, denominator `86030/1234`, `3716` intervals/zero overlap, conservation, workload/graph/DB/output/semantic/resource/source-status proof | `BOTH_TARGET_MEASUREMENTS_RECORDED / SUPERVISOR_PENDING`; keep both packets unpromoted and accepted baselines/streak/checklists unchanged; open one fresh visible A002 per-attempt Supervisor |
| A002 Supervisor PASS and Main KEEP | `E2-P2A-A002REVIEW1/DECISION1`; exact Supervisor report/verdict, current identity, both target elapsed/equivalence/resource gates | promote each repo's A002 values separately, reset/retain D001 streak `0`, keep parent/D001 unchecked, and record `E2-P2A-A003CURRENT1`; no per-attempt detect/stage/commit |
| A003 Architect and visible Planner complete | `E2-P2A-A003CURRENT1/ATTRIB1/ARCH1/PLAN2`; accepted basis, residual packet, canonical direction, exact owner/test/validation/measurement/rollback/STOP contract; `PLAN1` superseded; A002 checkpoint `ecf825d7` complete | historical transition completed through Coder, measurement, and Supervisor; `ARCH1/PLAN2` now prove exact prior scope/identity only and cannot authorize a new attempt |
| A003 Owner KEEP after exact rollback | `E2-P2A-A003REVIEW1/DECISION1/ROLLBACK1/OWNERKEEP1/COMMIT1`; exact target-separated metrics, `SUPERVISOR_A003_PASS`, restored hashes, full-build PASS, detect, and checkpoint `b6bf45bce95323aa6b53b182edfea8628bd8b463` | `OWNER_KEEP / RESTORE_COMPLETE / A003_CHECKPOINT_COMPLETE / A004 ARCHITECT_PENDING / D001_STREAK_0`; Main opens A004 Architect; no A003 gate rerun |
| WAL correctness fix independently passes and checkpoints | `E2-P2A-WALFORCEPLAN1/FIX1/REVIEW1/COMMIT1`; exact four-file candidate, final full-build/focused-test/runtime-repro PASS, exact verdict `SUPERVISOR_CHILD06A_WAL_FORCE_FIX_PASS`, detect PASS, checkpoint `0f3a572331dd23d17688886fcbfebeb7d37ee35d` | A004 Architect active; stop after its direction for Owner discussion, with Planner/Coder locked |
| A004 architecture receives Owner approval | `E2-P2A-A004ARCH1`; finalized report/verdict, exact residual/call paths, one-file owner, HIGH/CRITICAL blast radius, single-key/single-final-sort direction, parity/resource/validation/measurement/rollback/STOP contract | visible Planner translates exactly; Coder remains locked |
| A004 Planner translation complete | `E2-P2A-A004PLAN1`; living attempt card, ordered P2-A steps, four-ledger pointers, and Planner report are current | `A004_PLAN_READY_FOR_MAIN_VERIFY / CODER_LOCKED`; Main independently verifies the five owned artifacts and alone decides a separate future Coder release |
| A004 Supervisor PASS, Main NO_KEEP, and exact rollback complete | `E2-P2A-A004REVIEW1/DECISION1/ROLLBACK1/A005CURRENT1`; separate rejected measurements and restored source/test identities | retain accepted A003 bases; set D001 streak `1`; keep all checkboxes/queues unchanged; open A005 attribution then Architect, never Coder |
| A005 attribution and architecture complete | `E2-P2A-A005ATTRIB1/ARCH1`; exact record/projection cause, owner/call path, canonical-byte architecture, Main handoff verification `PASS` | one visible Planner translates exactly; Coder remains locked |
| A005 Planner translation complete | `E2-P2A-A005PLAN1`; living attempt card, concrete P2-A steps, four-ledger pointers, and Planner report are current | `A005_PLAN_READY_FOR_MAIN_VERIFY / CODER_LOCKED`; Main independently verifies the five Planner-owned artifacts and alone decides a separate future Coder release |
| A005 Main verification PASS | `E2-P2A-A005MAINVERIFY1`; exact five-artifact diff/scope/state checks and Supervisor-protocol report | `A005_PLAN_COMPLETE / A005_MAIN_VERIFIED / CODER_RELEASE_PENDING`; open exactly one visible Coder, then require its fresh pre-edit gate |
| A005 Coder and Main handoff PASS | `E2-P2A-A005IMPACT1/SRC1/BUILD1/TEST1`; exact source/test/build/test/report identities | `A005_MEASUREMENT_READY`; build one frozen candidate, then launch separate target measurement lanes only from its accepted packet |
| Supervisor `REJECT` or no retainable gain | `AnnnREVIEW1`, `AnnnDECISION1`, `AnnnRESTORE1` | increment unsuccessful streak; rejected candidate is not ranked |
| unsuccessful child streak `1` or `2` | restored accepted state plus rejection/no-gain packet | new attempt records current parent/child basis before a new Visible Architect for the same child |
| unsuccessful child streak `3` | three full attempt families plus `AnnnSYSTEM1` and accepted `AnnnRANK1` refresh | child `SYSTEM_CHARACTERISTIC`; remeasure accepted child/parent/E2E, refresh both ordered lists, check only that child, and select the largest remaining unchecked child of the same parent |
| all active-parent children checked | complete child ledger/checklist plus accepted parent/full-pipeline remeasurement | check parent, refresh top-level list, select largest remaining unchecked parent |

### `E2-P2A-A005IMPACT1/SRC1/BUILD1/TEST1` — Coder Candidate Ready For Measurement

Status: `CODER_A005_READY_FOR_MEASUREMENT_HANDOFF / MAIN_HANDOFF_PASS / A005_MEASUREMENT_READY`.

- Coder task/report: `01a03fa8-680f-7400-9ba6-f0d1db5c59dd`; `E:\Anvien\reports\coder\rp_coder_260827_by_gpt-5_child06a_a005_outcome_serialization.md`; report SHA-256 `26360A4BEAE5304ADF261CE5FFC444DBA09FA2E61F7262AE81F7D10BBF841EC8`.
- `A005IMPACT1`: clean base HEAD `cf230eeee87e2cc212c42639c8c7e99eba084bc7`; one fresh graph `2254/766/0`, `124251/171117`; both production files HIGH/current/non-stale; exact collector/init/record/finalize/projection/ResolveBoundInto impacts CRITICAL with complete counts in the Coder report.
- `A005SRC1`: exact production/test boundary only. `outcome.go +43/-24`, SHA-256 `18203DFAB9A227B526F8F7478B516AE6673F635BABC02D9463975E428A3983AF`; `resolve.go +2/-2`, SHA-256 `76B7B62A060B36EE2438E76689E858544358AD681DB42F6A4FC47D271F1749A1`; new `outcome_serialization_test.go +652/-0`, SHA-256 `89B168B2764B9A9B8EACDAB37A071BE0E1C51A8F791FE158B111651E9640C957`; P6C3 adaptation `+8/-5`, SHA-256 `E561280B3F8420D2288001179431A37FC39E6E2B8564863C9AD644EEE005A2E6`; total `+705/-31`; diff-check/gofmt clean; staged empty.
- `A005BUILD1`: holder/process preflight found no proven build-output holder; canonical `scripts/full-build.ps1` exited `0` before tests; vendor `1798/0/0/0`; runtime SHA-256 `CFDDA72DBB448AB9BE20FC58A7C20A883353E7F0B3EE7DF60A6DDD0968163AC7`; maintenance analyze `2255/767/0`, graph `124463/171506`.
- `A005TEST1`: exact focused A005 plus nine named P6B/P6C3/P6D/Graph JSON/Ladybug regressions exited `0`; `internal/graphhealth`, `internal/analyze`, `internal/lbugload`, and `internal/lbugnative` exit `0`. Full `internal/resolution` truthfully exits `1` only on unchanged `TestProofBasedCallAccessGoldenCorpus`; the package is not called PASS and golden blob remains exact HEAD `00d0ae042f109cd7d3c8aeadf855d192a5aa8172`.
- Main source/diff/report verification `E:\Anvien\reports\Supervisor\rp_supervisor_260827_034025_by_gpt-5_child06a_a005_coder_handoff.md` returns `PASS` for frozen-candidate/measurement handoff only. Candidate remains uncommitted/unstaged; no target measurement, per-attempt Supervisor disposition, detect, stage, or implementation commit exists.
- Next owner: one exact frozen-candidate builder using the unchanged A00x script and A005 overlay/hash inputs. It must not run target analyze.

### `E2-P2A-A005BUILDFAIL1` — Frozen Overlay Bundle Compile Mismatch

Status: `A005_FROZEN_BUILD_FAILED / OVERLAY_BUNDLE_ADAPTATION_RECOVERY_READY / MEASUREMENT_LOCKED`.

- Initial builder task `01a03fd6-6a2d-7573-85b8-a503d67cd299` copied the accepted overlay then stopped before build because Main’s first patch pattern omitted the existing line context. Resume task `01a03fd9-7029-7ee3-bd2e-4a94b5ef0296` completed the intended two public-result `outcomes.values` substitutions, verified overlay/manifest hashes `90FAA949...BA13E` / `FD555BCE...EA53`, and ran the first actual script invocation exactly once.
- Build invocation 1 exited `1` and emitted no `A00X_BENCHMARK_BUILD_COMPLETE`. Exact compile errors are copied overlay `resolve.go` lines `226/232`, where `resolutionOutcomeProjectionDenominators` still receives private `finalizedResolutionOutcomes`, and line `260`, where `len(outcomes)` targets that bundle rather than its slice. No binary, valid provenance packet, target analyze, measurement, benchmark value, disposition, streak, checklist, detect, stage, or commit was created.
- Exact narrow recovery: preserve the failed root; change only both `resolutionOutcomeProjectionDenominators(..., outcomes)` calls to `outcomes.values` and only the `resolutionChildAssembleResolutionResult` count from `len(outcomes)` to `len(outcomes.values)`. Do not change the separate helper-local `len(outcomes)` at line `561`, whose parameter remains a slice. Recovered overlay SHA-256 must equal `304F65E40629AE9B32803BA3D61ECD093022481381C948D8CF0B09AD75BFF788`; manifest/types remain unchanged.
- Recovery uses a new output root and one invocation of the unchanged accepted script. This is a generated-input compatibility correction under binding plan rules, not a production attempt, no-KEEP result, source defect, script audit, or permission to edit production/tests/scripts.

`REWORK` and `ROLLBACK` never authorize Coder directly. They describe the failed attempt's disposition; any next production edit starts a new attempt at Visible Architect.

### Attempt Result Ledger

Append one row after every attempt result. Current rows: `3`.

| Attempt | Parent row / selected child / accepted baseline | Complete child list / checklist | Architect evidence | Planner refresh | Coder / build | Child / parent / E2E rows | Supervisor result | Disposition | Child unsuccessful streak after | Promotion / restoration / checklist effect | Next owner / action |
|---------|-------------------------------------------------|---------------------------------|--------------------|-----------------|---------------|---------------------------|-------------------|-------------|-----------------------------------|------------------------------------------|---------------------|
| `A001` | `B1-P1A-OP001` / `B2-P2A-A001-D001`; pre-A001 baselines remain separate: Cheapapp D001/parent/process `209.823705100 / 733.384225100 / 890.314783200 s`; Restaurant D001/parent/process `501.638742800 / 1090.085959900 / 1178.391336900 s` | complete `17/17` rows per target; parent and all `17` child checklist items remain unchecked | `E2-P2A-A001ARCH1` | `E2-P2A-A001PLAN1`, `E2-P2A-A001CONSISTENCY1` | `E2-P2A-A001IMPACT1/SRC1/BUILD1/TEST1/COMMIT1`; exact five-file `+174/-10` candidate, build PASS, focused PASS, pre-existing package golden FAIL recorded truthfully, implementation commit `17a1f3af37dcb61f9d389345822b6470a8f772cc` complete; Coder report `E:\Anvien\reports\coder\rp_coder_260825_175710_by_gpt-5_child06a_a001_import_claim_index.md` remains outside that implementation manifest and is committed separately at `18b5063d236f9f2567fea90e48eca8f1501bd1eb` with a one-file report manifest | separate accepted after values: Cheapapp D001/parent/process `25.045225300 / 184.481061700 / 279.105934600 s`; Restaurant `40.769294200 / 136.436879300 / 218.680628900 s`; denominators unchanged | `SUPERVISOR_A001_PASS`; `E2-P2A-A001REVIEW1` | `KEEP`; `E2-P2A-A001DECISION1` | `0` | promote each repo's optimized values as that repo's accepted current baseline; never average/combine; leave P2-A/parent/D001 unchecked and D002-D017 queued/unopened | `A001_COMMIT_COMPLETE`; Main Orchestration opens fresh visible Architect A002 on the same active D001; Coder remains closed |
| `A002` | `B1-P1A-OP001` / `B2-P2A-A001-D001`; accepted A001 baselines separate: Cheapapp D001/parent/process `25.045225300 / 184.481061700 / 279.105934600 s`; Restaurant `40.769294200 / 136.436879300 / 218.680628900 s` | complete `17/17` rows per target; parent and all `17` child checklist items remain unchecked | `E2-P2A-A002ARCH1` | `E2-P2A-A002PLAN1` | `E2-P2A-A002IMPACT1/SRC1/BUILD1/TEST1/SCRIPTBUILD1/COMMIT1`; exact three-production-plus-one-test boundary, build PASS, focused/graphhealth PASS, pre-existing package golden FAIL recorded truthfully; accepted checkpoint `ecf825d7` complete | separate promoted after values: Cheapapp D001/parent/process `3.090914200 / 19.040468000 / 136.729876000 s`; Restaurant `9.909636600 / 21.242055400 / 145.066210900 s`; denominators unchanged | `SUPERVISOR_A002_PASS`; `E2-P2A-A002REVIEW1` | `KEEP`; `E2-P2A-A002DECISION1` | `0` | promote each repo's A002 values separately; checkpoint accepted bytes; leave P2-A/parent/D001 unchecked and D002-D017 queued | `E2-P2A-A003CURRENT1/ATTRIB1/ARCH1/PLAN2`; Main opens a fresh visible Coder only after PLAN2 handoff |
| `A003` | `B1-P1A-OP001` / `B2-P2A-A001-D001`; A002 measurements remain the separate before basis | complete `17/17` rows per target; parent and all `17` child checklist items remain unchecked | `E2-P2A-A003ARCH1` historical exact-candidate scope only | `E2-P2A-A003PLAN2/OWNERKEEP1` | prior `IMPACT1/SRC1/BUILD1/TEST1/PACKET1` complete; exact two-file A003 hashes are materialized; checkpoint `b6bf45bce95323aa6b53b182edfea8628bd8b463` complete; no gate rerun | Cheapapp D001/parent/analyzer/process `3.090914200 -> 3.447846300 / 19.040468000 -> 20.472602300 / 100.843249000 -> 93.531974900 / 136.729876000 -> 95.630648200 s`; Restaurant `9.909636600 -> 9.401585300 / 21.242055400 -> 20.850792800 / 109.339859600 -> 98.020546700 / 145.066210900 -> 101.096911900 s`; denominators unchanged | `SUPERVISOR_A003_PASS`; `E2-P2A-A003REVIEW1` | Owner `KEEP / RESTORE_COMPLETE / A003_CHECKPOINT_COMPLETE` | `0` | exact A003 bytes checkpointed; all checkboxes unchanged; WAL fix later checkpointed separately | `E2-P2A-A004ARCH1/PLAN1` complete; Main verifies A004 plan while Coder remains locked |

### Final P2-A Evidence

| Evidence ID | Required proof | Current status |
|-------------|----------------|----------------|
| `E2-P2A-FINALBUILD1` | canonical full build of stable accepted complete state | pending |
| `E2-P2A-FINALTIME1` | initial-versus-final total graph-generation measurement on same workload | pending |
| `E2-P2A-FINALEQUIV1` | aggregate correctness/output/persistence/reader/determinism/freshness/failure/transaction/temp/publication equivalence | pending |
| `E2-P2A-EXHAUST1` | every measured top-level parent and every measured child has retained terminal evidence and a checked plan item; no smaller row is omitted; no `BLOCKED` row remains unchecked | pending |

## E3 - P3 Evidence

Matching plan item(s): `P3-A`, `P3-B`, `P3-C`

| Evidence ID | Required proof | Current status |
|-------------|----------------|----------------|
| `E3-P3A-REVIEW1` | one final whole-candidate Supervisor `PASS` over initial/final measurements, all per-attempt reviews, stable source/build, and preserved invariants | pending/locked |
| `E3-P3B-CLEAN1` | exact Child 06A dead/debug work disposition and accepted production/test bytes unchanged from P3-A | pending/locked |
| `E3-P3C-DETECT1` | fresh post-cleanup detect-changes on accepted implementation boundary | pending/locked |
| `E3-P3C-COMMIT1` | exactly one isolated Child 06A implementation commit and verified manifest | pending/locked |
| `E3-P3C-HANDOFF1` | Child 07 predecessor/opening points to exact Child 06A closure commit and accepted performance/equivalence basis | pending/locked |

## Handoff Rules

Every Child 06A execution handoff must include:

- FULL RAW method authority `E:\Anvien\reports\Investigation\rp_main_260824_204615_child06a_optimization_method_discussion.md` and exact required provenance path `E:\Anvien\reports\Investigation\rp_huong_dan_lam_viec_de_slice_p6e_tien_hanh_nhanh_hon.md`, with an explicit note that current per-attempt Architect/Planner/Supervisor and three-attempt rules supersede conflicting older workflow text;
- exact active parent row, complete child-list snapshot, active child row, remaining parent/child checklist queues, accepted baseline, denominators, child unsuccessful-attempt streak, and disposition;
- active attempt ID and exact completed `CURRENT/DRILL/ARCH/PLAN/IMPACT/SRC/TEST/BUILD/DIRECT/PARENT/E2E/EQUIV/REVIEW/DECISION/RESTORE/SYSTEM/RANK` evidence as applicable;
- exact current owner and next owner/action; after Supervisor `REJECT`, next owner is Visible Architect after accepted-state restoration, never Coder;
- during P1-A, exact visible capture-executor/accepted-capture identity and zero-to-three read-only analyst assignments; if `E1-P1A-INSTR1` opened, the sole writer, exact owner/impact, build-before-use, comparability/equivalence, and completed carry-forward or remove/rebuild/remeasure disposition;
- measurement-pool state: `10` pre-opened visible lanes total, exact zero-to-three ACTIVE assignments (one measured parent/child problem per active lane), remaining waiting count, shared accepted capture identity when used, and one shared-worktree production/test writer at most; no active slot is freed until that lane's benchmark number, evidence proof, plan checklist state, and actual-status pointer are recorded;
- no rerun of still-valid accepted evidence when runtime/workload/boundary is unchanged;
- lane report limited to lane/work/result/checkpoint/next owner, with no report audit, hash chain, evidence-about-evidence, or documentation review.

`E3-P3C-HANDOFF1` is reserved for opening Child 07 from the one Child 06A closure commit.

## Closure Evidence

Closure requires any opened `E1-P1A-INSTR1` branch to have exactly one completed carry-forward or remove/rebuild/remeasure disposition; every measured parent and child benchmark row to have retained terminal evidence and a checked plan item; no unresolved unchecked `BLOCKED` row; lower comparable final total; complete per-attempt parent/child/Architect/Planner/Coder/Supervisor evidence; final equivalence; `E3-P3A-REVIEW1` `PASS`; exact cleanup; one detect; one implementation commit; and `E3-P3C-HANDOFF1`. No current closure evidence exists.
