# Child 06A Accelerate Analyze Without Sacrificing Accuracy Evidence Ledger

## Metadata

- Date: `2026-08-24`
- Current state: `A001 KEEP / A001_COMMIT_COMPLETE / A002 ARCHITECT_PENDING`; `E2-P2A-A001REVIEW1` records `SUPERVISOR_A001_PASS`, `E2-P2A-A001DECISION1` records Main's separate `KEEP` and repo-specific baseline promotion, and `E2-P2A-A001COMMIT1` records exact implementation commit `17a1f3af37dcb61f9d389345822b6470a8f772cc`; no A002 architecture/Coder evidence exists
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

Current attempt basis: A001 source/build evidence remains complete and optimized bytes unchanged. `E2-P2A-A001REVIEW1` records `SUPERVISOR_A001_PASS`; Main's `E2-P2A-A001DECISION1` records A001 `KEEP`; `E2-P2A-A001COMMIT1` records `A001_COMMIT_COMPLETE` at exact commit `17a1f3af37dcb61f9d389345822b6470a8f772cc`. P2-A, parent, and all children remain unchecked; D001 remains active and its no-KEEP streak is `0`; D002-D017 remain queued/unopened. Existing `E:\Anvien` numeric rows stay superseded. Cheapapp and Restaurant Manager independently retain complete before/after A001 data, and each repo's optimized values remain its own accepted current baseline without averaging/combining. Current cursor is `A002 / ARCHITECT_PENDING`; A002 Architect/Planner/Coder/measurement/Supervisor evidence, detect, and stage remain absent.

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

### Attempt State Machine

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
| Supervisor `REJECT` or no retainable gain | `AnnnREVIEW1`, `AnnnDECISION1`, `AnnnRESTORE1` | increment unsuccessful streak; rejected candidate is not ranked |
| unsuccessful child streak `1` or `2` | restored accepted state plus rejection/no-gain packet | new attempt records current parent/child basis before a new Visible Architect for the same child |
| unsuccessful child streak `3` | three full attempt families plus `AnnnSYSTEM1` and accepted `AnnnRANK1` refresh | child `SYSTEM_CHARACTERISTIC`; remeasure accepted child/parent/E2E, refresh both ordered lists, check only that child, and select the largest remaining unchecked child of the same parent |
| all active-parent children checked | complete child ledger/checklist plus accepted parent/full-pipeline remeasurement | check parent, refresh top-level list, select largest remaining unchecked parent |

`REWORK` and `ROLLBACK` never authorize Coder directly. They describe the failed attempt's disposition; any next production edit starts a new attempt at Visible Architect.

### Attempt Result Ledger

Append one row after every attempt result. Current rows: `1`.

| Attempt | Parent row / selected child / accepted baseline | Complete child list / checklist | Architect evidence | Planner refresh | Coder / build | Child / parent / E2E rows | Supervisor result | Disposition | Child unsuccessful streak after | Promotion / restoration / checklist effect | Next owner / action |
|---------|-------------------------------------------------|---------------------------------|--------------------|-----------------|---------------|---------------------------|-------------------|-------------|-----------------------------------|------------------------------------------|---------------------|
| `A001` | `B1-P1A-OP001` / `B2-P2A-A001-D001`; pre-A001 baselines remain separate: Cheapapp D001/parent/process `209.823705100 / 733.384225100 / 890.314783200 s`; Restaurant D001/parent/process `501.638742800 / 1090.085959900 / 1178.391336900 s` | complete `17/17` rows per target; parent and all `17` child checklist items remain unchecked | `E2-P2A-A001ARCH1` | `E2-P2A-A001PLAN1`, `E2-P2A-A001CONSISTENCY1` | `E2-P2A-A001IMPACT1/SRC1/BUILD1/TEST1/COMMIT1`; exact five-file `+174/-10` candidate, build PASS, focused PASS, pre-existing package golden FAIL recorded truthfully, implementation commit `17a1f3af37dcb61f9d389345822b6470a8f772cc` complete; Coder report `E:\Anvien\reports\coder\rp_coder_260825_175710_by_gpt-5_child06a_a001_import_claim_index.md` remains outside that implementation manifest and is committed separately at `18b5063d236f9f2567fea90e48eca8f1501bd1eb` with a one-file report manifest | separate accepted after values: Cheapapp D001/parent/process `25.045225300 / 184.481061700 / 279.105934600 s`; Restaurant `40.769294200 / 136.436879300 / 218.680628900 s`; denominators unchanged | `SUPERVISOR_A001_PASS`; `E2-P2A-A001REVIEW1` | `KEEP`; `E2-P2A-A001DECISION1` | `0` | promote each repo's optimized values as that repo's accepted current baseline; never average/combine; leave P2-A/parent/D001 unchecked and D002-D017 queued/unopened | `A001_COMMIT_COMPLETE`; Main Orchestration opens fresh visible Architect A002 on the same active D001; Coder remains closed |

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
