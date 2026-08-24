# Main Orchestration Rotation Handoff — Child 06A Implementation Opening

Status: `SEALED_HANDOFF_READY_FOR_OFFICIAL_TRANSFER`

## Authority Transfer

- Outgoing Visible Main: `01a033bf-1cec-7bf3-ab52-545915b3c3fe`.
- Incoming clean Visible Main successor: `01a03472-893f-7c92-95ac-7221631c002e`, host `local`, saved project `E:\Anvien`, no worktree.
- Latest direct User command: hand off to a new Main lane to implement the campaign. This command ends the prior User-directed rotation pause and authorizes the successor to begin implementation orchestration without asking again.
- Authority remains with the outgoing Main until the successor receives the follow-up headed exactly `OFFICIAL AUTHORITY TRANSFER` citing this sealed report.
- Incoming authority start, warm target `+45 minutes`, and hard rotation deadline `+60 minutes` are recorded in that official follow-up.
- After official transfer, outgoing Main terminates and must not continue campaign orchestration.

## Mandatory Raw Startup

The successor must read FULL RAW sequentially through EOF and must not substitute this handoff, memory, summary, or compacted context for any source:

1. `E:\Anvien\AGENTS.md`.
2. `E:\Anvien\.agents\skills\working-rules\SKILL.md`.
3. `E:\Anvien\.agents\skills\orchestration\SKILL.md`.
4. `E:\Anvien\.agents\skills\governance-rule-guard\SKILL.md`.
5. `E:\Anvien\.agents\skills\planner\SKILL.md` and its four templates.
6. All five current Child 06A files under `E:\Anvien\docs\plans\2026-07-26-anvien-graph-accuracy-multi-plan\2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy`.
7. `E:\Anvien\reports\Investigation\rp_main_260824_204615_child06a_optimization_method_discussion.md`.
8. `E:\Anvien\reports\Supervisor\rp_supervisor_260824_223907_by_gpt-5_child06a_plan_rewrite_resubmission.md`.
9. This handoff report.

AGENTS.md, working-rules, and orchestration must be applied 100% raw and unmodified. After auto-compact or context loss, reread the raw sources as required; do not restart passed gates merely to re-anchor.

## Exact Campaign Goal And Current Slice

Exact goal:

> Complete Child 06A by reducing the measured elapsed wall-clock time of the real `anvien analyze` graph-generation pipeline as far as current evidence safely permits, while preserving accepted accuracy, semantic completeness, graph correctness, output, persistence, readers, determinism, freshness, failure, transaction, temporary-file, and publication behavior; then close through one final Supervisor, cleanup/detect, exactly one implementation commit, and hand off Child 07.

Current plan:

`E:\Anvien\docs\plans\2026-07-26-anvien-graph-accuracy-multi-plan\2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy\2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy-plan.md`

Current state:

- `P0-A` is complete.
- `P1-A` is the first incomplete and immediately executable item.
- No Child 06A functional measurement exists.
- `B1-P1A-TOTAL` is not measured.
- No accepted capture, detailed timing map, top-level bottleneck rows, parent checklist, child rows, active bottleneck, Architect attempt, Coder attempt, speedup, final equivalence, final Supervisor result, cleanup, detect, or implementation commit exists.
- The plan rewrite has independent documentation acceptance `PASS` at `reports/Supervisor/rp_supervisor_260824_223907_by_gpt-5_child06a_plan_rewrite_resubmission.md`.
- No new documentation audit or plan-authoring loop is needed before P1-A. Start the work and update the living ledgers from actual results.

## Child Order And Pn Clarification

- P6-D is accepted at commit `81163e39718b94a509e41114cada224e8f269e36`; it remains an ancestor of current HEAD.
- Direct order is `Child 06 -> Child 06A -> Child 07`.
- Current Child 06 checklist ends at `P6-D`. Current Child 06 does not contain phases `Pn-A`, `Pn-B`, or `Pn-C`.
- The pre-split committed Child 06 plan did contain `Pn-A/B/C` and evidence IDs `E6-PNA/PNB/PNC-*`. Current Child 06A references those only as legacy provenance. Do not treat them as current Child 06 gates and do not add them back.
- Child 06A `P3-A/P3-B/P3-C` are the new closure sequence: final whole-candidate Supervisor, exact cleanup, detect/one implementation commit/Child 07 handoff.

## Binding Optimization Method

- Elapsed wall-clock time is the primary and controlling product metric.
- CPU, RAM, allocation, GC, I/O, waits, calls, bytes, counts, denominators, and shares are secondary explanation/comparability evidence. They cannot choose priority or prove a speedup.
- P1-A creates the complete top-level list from current absolute elapsed times and inserts one exact unchecked parent checklist item per benchmark row.
- P2-A is exactly one implementation slice. It selects the largest unchecked parent, measures deeply inside that parent, creates the complete child list and exact nested checklist, processes the largest child first, and then every smaller measured child. No measured row is skipped because it is small.
- Every production attempt follows: current child/parent/end-to-end basis -> fresh Visible Architect -> Planner refresh -> Coder -> canonical full build -> child/parent/end-to-end remeasurement -> fresh Visible Supervisor accuracy/equivalence review -> disposition.
- `KEEP` requires all four: child elapsed time lower, parent elapsed time lower, end-to-end elapsed time lower, and that attempt's Supervisor `PASS`.
- On `REJECT` or no retainable gain, restore/retain the last accepted baseline; rejected bytes never drive ranking. Another production edit requires a new Architect decision, Planner refresh, Coder, measurement, and Supervisor.
- Three consecutive attempts without `KEEP` terminalize only the exact child as `SYSTEM_CHARACTERISTIC`. That child is checked; its parent remains open until every child is terminal and checked.
- A concrete `BLOCKED` row remains unchecked and blocks completion.
- P3-A is one final whole-candidate Supervisor in addition to per-attempt reviews. P3-B changes no accepted production/test bytes. P3-C performs one detect, one implementation commit, and the Child 07 handoff.

## P1-A Opening Command Boundary

The successor must immediately orchestrate P1-A, not ask for another User START:

1. Assign exactly one visible measurement executor to produce the initial accepted real `anvien analyze` capture on the accepted workload. The executor records the actual executable/options/workload/cache/runtime/output identity at execution; do not invent a command tuple or numeric result in advance.
2. Use existing timing, benchmark output, and profiles first.
3. Record total graph-generation time, real top-level operation times, boundaries, and denominators immediately in benchmark/evidence/status.
4. After an accepted capture exists, at most three ACTIVE read-only measurement analysts from the ten-lane pool may share it for independent measured problems. They do not edit source and do not launch duplicate competing benchmark processes.
5. Only if one required P1 timing is genuinely missing may Main assign one separate visible sequential sole instrumentation writer. Before edit: fresh required graph, file-detail, and impact on the exact owner. After edit: canonical full-build `PASS` before timing use. Record like-for-like instrumentation identity, overhead/comparability, denominator, and output equivalence. Close with exactly one disposition: carry exact ownership into the first refreshed P2-A attempt, or remove exact owned bytes, rebuild, and re-establish/remeasure the accepted timing basis.
6. P1-A does not open Architect, Coder optimization, or Supervisor. Those begin only for a production optimization attempt after current measurement creates a complete parent/child inventory and selected child owner/cause.

## Visible Lane State

- Planner `01a033b7-9e01-7151-bdce-5d87173e1ac5`: idle/preserved. Use it for each P2-A attempt refresh after a fresh Architect decision; do not let Planner choose workflow.
- Replacement Coder `01a03375-f616-7651-ba81-99bad6536746`: preserved/HOLD. Do not activate for P1 pure measurement. Activate only after the first production attempt has complete parent/child evidence, fresh Architect direction, and Planner refresh. If conditional P1 instrumentation is needed, Main must explicitly scope the single writer branch rather than treating it as a production optimization attempt.
- Continuous Guard lineage: `01a03339-144d-7383-8cd1-72fcd5e5417c`; prior related lineage includes `01a03373-96ae-70d2-94c4-ac5e631a2a72`. Do not duplicate or control Guard as a functional lane. The official Main transfer record supplies the successor target.
- Measurement pool: User and current plan establish ten pre-opened visible lanes, maximum three ACTIVE, one measured problem per lane. The current fifty-thread snapshot exposed the exact waiting Segment 02-10 tasks:
  - Segment 02 `01a033a5-3a44-7c43-9aca-c85e3e932a0f`
  - Segment 03 `01a033a5-414c-7721-864f-a8f4642d46e3`
  - Segment 04 `01a033a5-4700-7272-a28d-de3a71f58135`
  - Segment 05 `01a033a5-4ec3-7e42-8946-0ab9172f6088`
  - Segment 06 `01a033a5-552d-7703-aa32-083c165cb55a`
  - Segment 07 `01a033a5-5dbb-7ce0-b1be-2ef3d8c98623`
  - Segment 08 `01a033a5-64ec-7873-ac33-2065b9ef3812`
  - Segment 09 `01a033a5-6cb1-7d62-8426-f20ab3128ebf`
  - Segment 10 `01a033a5-7479-7c21-b297-2cd0fa4783fc`
- Segment 01 exact ID was not exposed by that bounded current snapshot. Do not create a duplicate merely from that absence. Recover its exact visible identity only if needed; P1-A needs exactly one executor, not all ten ACTIVE.
- No functional lane is active at transfer.

## Main Authority And Role Boundary

- User commands directly. Main has complete command authority below User and must translate settled decisions into lane commands immediately.
- Lanes execute Main commands and report evidence. They do not choose workflow, command Main, or require Main to seek repeat permission.
- Main is an orchestrator, not a functional worker: understand the whole plan, design lane scopes, monitor actual commands/files, block deviations, receive handoffs, verify evidence, and transition the campaign.
- No hidden functional agents. Never use `fork_thread`.
- Every Architect, Coder, Planner, and Supervisor production-attempt role requiring User visibility must use a separate visible task.
- Only Supervisor issues acceptance verdicts. Main independently verifies handoffs before transition.
- Documentation follows actual progress. Do not turn ledgers, reports, hashes, or wording into a standalone audit loop.

## Workspace, Git, And Protected State

- Sole workspace: `E:\Anvien`, saved project local, no worktree.
- Current branch: `master`.
- Current HEAD at handoff: `1c5de4ef6875a5e7b3329f04dafd1189c7622e4d` (`docs: strengthen plan control-plane guidance`).
- P6-D commit `81163e39718b94a509e41114cada224e8f269e36` is an ancestor of current HEAD.
- Git index is empty.
- Worktree is intentionally dirty. Preserve all user/historical/protected rows, including Owner deletion `CONTRIBUTING.md`, modified roadmap/Child 06/Child 07 ledgers, untracked Child 06A directory, reports, and raw coder artifacts. Do not reset, stash, checkout, broadly clean, broadly stage, or rewrite unrelated rows.
- The Child 06A method report and five plan files are uncommitted and protected current authority.
- No C-drive workspace/read/write, alternate worktree, target access, network, install/package scripts, shared/global/quarantined cache, broad cleanup, broad stage, or push.
- Repository-local `E:\Anvien\.tmp` is disposable scratch only and never durable evidence.

## Current Accepted Documentation Identities

- Method report: SHA-256 `65FD5A7FDD102FDF2CA90AD9728822C7BDA3900ABBE9EE192A72AF990B475333`, `53,667` bytes.
- Child 06A plan: SHA-256 `0F2933F6D900CEA5A16F12A7F9D0060CF01E732F71271ABC3671258D5932BA6B`, `53,345` bytes.
- `plan-rules.md`: SHA-256 `C8789E80D08FDA89CDECAD2B99666078C7DFFAC3ECE5139DBFD1E070DF5FFE33`, `13,559` bytes.
- Evidence: SHA-256 `6519DBAC2FED2B6EC597A0A672F5CD82989B1E9E19B9B55BB28A95FA8B0B21B9`, `27,958` bytes.
- Benchmark: SHA-256 `40A7D83D9C22AEB1160048611CC863F7E47F42DD4BD357B414F3A2240DD0DA77`, `21,911` bytes.
- Actual status: SHA-256 `F58F95396A3E496D609B414D78854C6A457DA02D26E0D41AF77844289B93EF1F`, `38,709` bytes.
- Supervisor resubmission PASS: SHA-256 `9FD52EC78856F035BEA941A2CA37A08DBB8D814B5B86B742DE131F66F9478F09`, `6,908` bytes.

These identities prove the transfer boundary only. Do not repeatedly hash/audit unchanged documentation before starting P1-A.

## Preserved Verified Violations

Preserve these prior verified orchestration violations exactly as historical governance facts; do not justify, normalize, omit, or downgrade them:

- `52.457s`
- `3.488s`
- interrupted Main monitoring turn
- passive wait-only orchestration
- prior warm miss `36.626s`
- current warm miss `1,735.644s`
- current hard miss `835.644s`

The User later explicitly paused rotation during plan discussion and now explicitly resumes it through this handoff command. No new lateness finding is created by this report.

## Immediate Successor Actions

After official authority transfer:

1. Perform the mandatory FULL RAW startup and record authority start/warm/hard times.
2. Preserve continuous Guard lineage without functional control or duplication.
3. Confirm current plan cursor is P1-A and current functional lane state is all HOLD/idle; do not restart the completed plan rewrite review.
4. Select and command exactly one visible measurement executor for the initial accepted capture. State exact goal, scope, allowed outputs, evidence, stop conditions, and first action. Monitor actual commands and machine load.
5. Keep replacement Coder and remaining measurement lanes HOLD until evidence creates valid independent work and current plan gates open them.
6. Drive P1-A to the complete top-level benchmark list and matching parent checklist; then begin P2-A through the required visible Architect -> Planner -> Coder -> measurement -> Supervisor loop.
7. Rotate this successor Main at warm/hard targets according to orchestration rules unless a later direct User command changes rotation.

## Detached Seal

The detached file identity is recorded by the outgoing Main after this file is written. The official transfer message must cite the exact path, byte length, and SHA-256. Any later byte change invalidates that seal and requires a new transfer record.
