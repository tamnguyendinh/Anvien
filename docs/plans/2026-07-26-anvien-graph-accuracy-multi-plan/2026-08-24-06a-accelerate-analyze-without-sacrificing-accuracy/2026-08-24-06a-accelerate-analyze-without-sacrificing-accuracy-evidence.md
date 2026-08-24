# Child 06A Accelerate Analyze Without Sacrificing Accuracy Evidence Ledger

## Metadata

- Date: `2026-08-24`
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
- Production implementation precedes tests; the canonical full build precedes remeasurement/validation. Fresh required Anvien graph/file-detail/impact runs immediately before editing the exact measured owner.
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
| `E1-P1A-TOTAL1` | current total graph-generation elapsed time from the accepted real path | recorded: process wall `605.732722 s`; analyzer-internal total `602.5278811 s`; process CPU `803.093750 s` |
| `E1-P1A-OPS1` | real internal operation boundaries, elapsed times, denominators, and validity for every ranked row | partial: all `15` emitted phase timers and available denominators recorded; `15.2892835 s` internal plus `3.2048409 s` process/CLI residual lacks complete operation attribution |
| `E1-P1A-INSTR1` | exact missing timing and why existing capability is insufficient; sole visible sequential writer; exact owner/owned bytes; fresh graph/file-detail/impact; canonical full-build PASS before use; accepted runtime/instrumentation identity; overhead, denominator, like-for-like comparability, and output equivalence; exactly one carry-forward ownership or remove/rebuild/remeasure disposition | OPEN / NOT COMPLETED: timing gap proven; no writer assigned, no source edited, no build or instrumented capture run, and no terminal disposition recorded |
| `E1-P1A-EQUIV1` | initial workload/output/persistence/publication validity sufficient to use the measurement | recorded for the uninstrumented capture: exit `0`, all `15` phases completed, output counts/identities recorded, index remained empty, and tracked dirty count stayed `10` with no AGENTS/CLAUDE diff |
| `E1-P1A-RANK1` | complete benchmark-owned top-level list of every measured real operation in descending elapsed-time order, including smaller rows, with current rank/time, denominator, meaningful share/delta, processing state, proven owner/call path when known, disposition/evidence, and an exact one-to-one unchecked parent checklist mirror in plan | pending/blocked by `E1-P1A-INSTR1`; no `B1-P1A-OPnnn` rows or plan parent items created |

### P1-A Result Record

Append the first completed measurement record here and update benchmark/status immediately. Current rows: none.

| Evidence ID | Recorded at | Visible executor / accepted capture | Exact command / exit | Runtime and workload identity | Operation boundary / denominator | Benchmark rows | Output validity | Instrumentation branch / disposition | Next exact action |
|-------------|-------------|-------------------------------------|----------------------|-------------------------------|----------------------------------|----------------|-----------------|--------------------------------------|-------------------|
| `E1-P1A-IDENTITY1/TOTAL1/OPS1/EQUIV1` | `2026-08-24T22:58:05.2715182+07:00` -> `2026-08-24T23:08:11.1166985+07:00` | executor `01a033a5-3a44-7c43-9aca-c85e3e932a0f`; capture `child06a-p1a-initial-20260824-225900` | exact command below; exit `0` | repo-owned runtime `1.2.8`; workload `E:\Anvien`; forced cold analyzer outputs; isolated E-only process environment | process wall/CPU plus all `15` emitted phases; exact workload/output denominators below | `B1-P1A-TOTAL` plus provisional `B1-P1A-PHASE-*`; no `OPnnn` queue rows | graph/Ladybug/meta and raw capture identities below; tracked/index boundary preserved | `E1-P1A-INSTR1` OPEN, no writer/edit/build/disposition; ranking blocked | Main assigns the separate visible sequential sole writer or records a concrete blocker; this executor stops |

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

`E1-P1A-RANK1` remains pending. The capture proves useful provisional phase order, but complete current top-level operation ranking and its one-to-one plan parent checklist do not exist while `E1-P1A-INSTR1` is open.

### Conditional P1-A Instrumentation Contract

This table is populated only if `E1-P1A-INSTR1` opens. It is evidence for one conditional branch, not a new phase, implementation slice, attempt, gate, or report system.

| Required pointer | Proof required before P1-A transition |
|------------------|---------------------------------------|
| Missing timing | OPEN: `15.2892835 s` analyzer-internal residual contains untimed graph compaction, DB runner resolve/close, and Graph JSON snapshot/finalization; `3.2048409 s` process/CLI residual contains untimed startup/storage preparation, benchmark write, registry/meta, generated context, file projection, and JSON output. Existing benchmark/profile capability supplies no separate wall boundary for these real operations, so complete absolute ranking is unavailable. |
| Writer / owner | NOT ASSIGNED. One separate visible sequential instrumentation writer outside read-only analysts is required; no owner or byte range is authorized by this executor. Source inspection identifies the timing surfaces in `internal/analyze/analyze.go` and `internal/cli/command.go`, but fresh owner proof remains mandatory before edit. |
| Pre-edit topology | PENDING. The accepted capture refreshed the repository graph, but the future sole writer must still record the exact required fresh graph/file-detail/impact immediately before its edit. |
| Build-before-use | PENDING; no source edit or build occurred in this executor. |
| Runtime validity | Current accepted uninstrumented capture is `child06a-p1a-initial-20260824-225900`; an instrumented capture does not exist. Like-for-like overhead/comparability and output equivalence remain pending. |
| Terminal disposition | PENDING; neither `CARRY_TO_FIRST_P2A_REFRESH` nor `REMOVE_REBUILD_REMEASURE` has been selected or completed. |

## E2 - P2 Evidence

Matching plan item(s): `P2-A`

### Attempt Evidence Contract

Every actual production attempt instantiates every mandatory ID below with one immutable attempt number; `RESTORE1` and `SYSTEM1` are additionally required when their stated condition occurs.

| Attempt evidence kind | Required proof |
|-----------------------|----------------|
| `E2-P2A-AnnnCURRENT1` | exact child-attempt goal; active parent row; selected child row; accepted child/parent/total elapsed times; denominators; child unsuccessful streak; cause/owner/complete call path; preserved invariants; complete child-list and plan-checklist pointers |
| `E2-P2A-AnnnDRILL1` | measurement inside only the active parent boundary; complete current absolute elapsed-time ordered child list including smaller costs; exact benchmark-row/plan-checkbox cardinality; selected largest unchecked child; each child's denominator and validity |
| `E2-P2A-AnnnARCH1` | new Visible Architect identity/turn; exact parent row, complete child list, selected child elapsed-time owner/cause, current child/parent/total basis, attempt-local direction, allowed owners, expected gain, invariants, validation/resources, and rollback |
| `E2-P2A-AnnnPLAN1` | Planner refreshed the living P2-A attempt and exact implementation/test/build/measure/review/rollback steps before Coder |
| `E2-P2A-AnnnIMPACT1` | fresh required graph/file-detail/impact for exact owner immediately before edit, including full HIGH/CRITICAL scope |
| `E2-P2A-AnnnSRC1` | exact production-first change and attempt-owned rollback boundary |
| `E2-P2A-AnnnTEST1` | behavior tests updated after production and exact affected boundary proved |
| `E2-P2A-AnnnBUILD1` | canonical full build result before runtime validation |
| `E2-P2A-AnnnDIRECT1` | selected-child before/after with same denominator and benchmark rows |
| `E2-P2A-AnnnPARENT1` | active-parent before/after on the same accepted work and benchmark rows |
| `E2-P2A-AnnnE2E1` | total graph-generation before/after on same workload and benchmark rows |
| `E2-P2A-AnnnEQUIV1` | exact affected correctness/output/persistence/reader/lifecycle result |
| `E2-P2A-AnnnREVIEW1` | new Visible Supervisor identity/turn and independent accuracy/equivalence/invariant `PASS` or `REJECT` |
| `E2-P2A-AnnnDECISION1` | disposition and precise promotion/restoration/next-owner effect after Supervisor |
| `E2-P2A-AnnnRESTORE1` | required on rejected/non-retained candidate: exact owned rollback and last accepted state restored/retained |
| `E2-P2A-AnnnSYSTEM1` | only on third consecutive unsuccessful child attempt: exact parent/child rows, denominator, retained child/parent times, three attempts/reasons, terminal child `SYSTEM_CHARACTERISTIC`, and matching child checklist check |
| `E2-P2A-AnnnRANK1` | accepted-state child/parent/full-pipeline timing refresh, complete child/top-level list update, and exact next unchecked child/parent pointer |

Current attempt rows: none. P1-A must first create the complete top-level list and matching parent checklist. P2-A then selects the largest unchecked parent, records its complete child list and exact nested checklist in `DRILL1`, selects the largest unchecked child, records `A001 CURRENT1`, and only afterward obtains the fresh Architect decision plus Planner refresh required for production editing.

### Attempt State Machine

| From | Required evidence | To / next owner |
|------|-------------------|-----------------|
| complete top-level list | `E1-P1A-RANK1` plus exact parent checklist mirror | select largest unchecked parent and measure its complete child list |
| complete parent drill-down | `AnnnDRILL1`, all dynamic child rows, and equal nested checklist cardinality | select largest unchecked child and record `AnnnCURRENT1` |
| selected child current basis | `AnnnCURRENT1` | new Visible Architect receives exact parent, complete child list, selected child/cause/owner/call path |
| Architect decision complete | `AnnnARCH1` | Planner refreshes current P2-A attempt |
| Planner refresh complete | `AnnnPLAN1` | Coder may edit only refreshed surfaces |
| Coder/build complete | `AnnnIMPACT1/SRC1/TEST1/BUILD1` | selected-child, parent, and end-to-end remeasurement |
| measurements/equivalence complete | `AnnnDIRECT1/PARENT1/E2E1/EQUIV1` | new Visible Supervisor |
| Supervisor `PASS`, child/parent/end-to-end elapsed times lower | `AnnnREVIEW1`, `AnnnDECISION1` | `KEEP`; promote baseline, reset active-child streak to `0`, remeasure, and open another attempt on the same unchecked child |
| Supervisor `REJECT` or no retainable gain | `AnnnREVIEW1`, `AnnnDECISION1`, `AnnnRESTORE1` | increment unsuccessful streak; rejected candidate is not ranked |
| unsuccessful child streak `1` or `2` | restored accepted state plus rejection/no-gain packet | new attempt records current parent/child basis before a new Visible Architect for the same child |
| unsuccessful child streak `3` | three full attempt families plus `AnnnSYSTEM1` and accepted `AnnnRANK1` refresh | child `SYSTEM_CHARACTERISTIC`; remeasure accepted child/parent/E2E, refresh both ordered lists, check only that child, and select the largest remaining unchecked child of the same parent |
| all active-parent children checked | complete child ledger/checklist plus accepted parent/full-pipeline remeasurement | check parent, refresh top-level list, select largest remaining unchecked parent |

`REWORK` and `ROLLBACK` never authorize Coder directly. They describe the failed attempt's disposition; any next production edit starts a new attempt at Visible Architect.

### Attempt Result Ledger

Append one row after every attempt result. Current rows: none.

| Attempt | Parent row / selected child / accepted baseline | Complete child list / checklist | Architect evidence | Planner refresh | Coder / build | Child / parent / E2E rows | Supervisor result | Disposition | Child unsuccessful streak after | Promotion / restoration / checklist effect | Next owner / action |
|---------|-------------------------------------------------|---------------------------------|--------------------|-----------------|---------------|---------------------------|-------------------|-------------|-----------------------------------|------------------------------------------|---------------------|

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
