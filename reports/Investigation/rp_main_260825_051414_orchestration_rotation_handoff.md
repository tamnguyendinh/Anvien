# Main Orchestration Rotation Handoff — 2026-08-25 05:14:14 +07:00

## Authority

- Outgoing Main: `01a03592-cd73-7bb0-af53-443f849343fb`.
- Clean visible successor: `01a035c8-7943-7fb2-9ee1-5d5daf991b0b` on host `local`.
- Outgoing authority start: `2026-08-25T04:14:14.223+07:00`.
- Successor warm creation/ready: created before `2026-08-25T04:58:53.0392760+07:00`; mandatory WAIT ACK completed before warm target.
- Successor warm target: `2026-08-25T04:59:14.223+07:00`.
- Successor authority start / exact transfer point: `2026-08-25T05:14:14.223+07:00`.
- Successor hard deadline: `2026-08-25T06:14:14.223+07:00`; warm target `2026-08-25T05:59:14.223+07:00`.
- Exact follow-up headed `OFFICIAL AUTHORITY TRANSFER` sent after this file is sealed is the sole new authority source of truth.

## Mandatory Raw Startup

Read every source FULL RAW sequentially through EOF. Summaries, compact context, this handoff, and prior memory are not substitutes:

1. `E:\Anvien\AGENTS.md`
2. `E:\Anvien\.agents\skills\working-rules\SKILL.md`
3. `E:\Anvien\.agents\skills\orchestration\SKILL.md`
4. `E:\Anvien\.agents\skills\governance-rule-guard\SKILL.md`
5. `E:\Anvien\.agents\skills\planner\SKILL.md`
6. All four files under `E:\Anvien\.agents\skills\planner\templates\` named `plan.template.md`, `evidence.template.md`, `benchmark.template.md`, and `actual-status.template.md`, in that order.
7. `E:\Anvien\docs\plans\2026-07-26-anvien-graph-accuracy-multi-plan\2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy\2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy-plan.md`
8. `E:\Anvien\docs\plans\2026-07-26-anvien-graph-accuracy-multi-plan\2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy\2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy-evidence.md`
9. `E:\Anvien\docs\plans\2026-07-26-anvien-graph-accuracy-multi-plan\2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy\2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy-benchmark.md`
10. `E:\Anvien\docs\plans\2026-07-26-anvien-graph-accuracy-multi-plan\2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy\2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy-actual-status.md`
11. `E:\Anvien\docs\plans\2026-07-26-anvien-graph-accuracy-multi-plan\2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy\plan-rules.md`
12. This sealed handoff.

Apply the raw rules unmodified. Do not restart accepted capture, instrumentation impact/source/test work, canonical build/full-build validation, previous graph refreshes, documentation review, commits, or old measurement work.

## Startup Compliance And Preserved Violation

- Immediately after the preceding official transfer, this Main initially read `working-rules/SKILL.md` before reading FULL RAW `AGENTS.md`. Guard correctly issued `HIGH — VERIFIED VIOLATION`.
- Main stopped all campaign work, preserved the violation without downgrade, then restarted the entire Mandatory Raw Startup sequence with `AGENTS.md` first and every listed source read through EOF in exact order.
- Large plan/ledger files were reread in contiguous chunks after one whole-file output was visibly truncated; no truncated output was treated as FULL RAW.
- Do not normalize, omit, justify, or downgrade this startup-order violation. The corrective restart restores compliance; it does not erase the event.

## Exact Campaign Goal And Cursor

- Exact goal: complete Child 06A using exhaustive measured-bottleneck optimization of real `anvien analyze`, preserving final equivalence; exactly one final P3-A Supervisor boundary, P3-B cleanup, P3-C detect and exactly one implementation commit, then direct Child 07 handoff.
- P0-A and P1-A are complete.
- `E1-P1A-INSTR1` is terminal `CARRY_TO_FIRST_P2A_REFRESH`.
- The complete top-level queue is `B1-P1A-OP001..OP030`, with exactly `30` matching unchecked plan parent items.
- P2-A is the first incomplete slice and is open only at `PARENT_DRILLDOWN_PENDING` on active parent `B1-P1A-OP001 resolution`.
- No dynamic child row, nested child checklist item, A001, Architect decision, Planner production-attempt refresh, production optimization, Coder result, per-attempt Supervisor, `KEEP`, speedup, final Supervisor, cleanup, detect, or Child 06A implementation commit exists.
- Production Coder remains locked until the active parent has a complete measured child list, exact benchmark/checklist cardinality, the largest unchecked child's exact cause/source owner/complete call path, a fresh visible Architect decision, and Main/planner refresh.

## Accepted P1-A Numeric Basis

- Immutable uninstrumented reference: process wall `605.732722 s`, analyzer internal `602.5278811 s`, 15-phase sum `587.2385976 s`.
- Accepted carried-instrumentation capture: `child06a-p1a-instrumented-20260825-040400` under `E:\Anvien\.tmp\child06a_p1a_instrumented_20260825_040400`.
- Exact accepted command: `E:\Anvien\anvien\bin\anvien.exe analyze E:\Anvien --force --json --progress --benchmark-json E:\Anvien\.tmp\child06a_p1a_instrumented_20260825_040400\benchmark.json --benchmark-label child06a-p1a-instrumented-20260825-040400`.
- PID/start/end/exit: `14956`; `2026-08-25T04:11:18.7266556+07:00`; `2026-08-25T04:21:37.1624765+07:00`; `0`; exactly one launch and no retry.
- Process wall/CPU/analyzer/phase: `618.4358209 / 799.4062500 / 615.8656873 / 600.8917213 s`.
- Workload/output: `2,200` scanned, `765` parsed, `0` failed, `123,277` nodes, `169,976` relationships, projection `20,357 / 479`.
- Operation integrity: `30` globally unique rows; `21 analyzer_internal`, `1 analyzer_outer`, `8 cli_outer`; all durations and denominators non-negative; all 15 phase names/order/durations match; analyzer non-phase `14.9739660 s`; analyzer-internal sum equals total exactly; analyzer residual `0`; analyzer outer `0.1697595 s`; CLI outer `2.2274029 s`; all operations `618.2628497 s`; process wrapper residual `0.1729712 s`.
- Raw artifact identities and cross-capture workload/hash/count mismatches are recorded in the current evidence ledger. Arithmetic differences versus the immutable reference are non-comparable and are not speed/regression or instrumentation-overhead evidence.

## Build And Instrumentation Boundary

- Sole instrumentation writer: `01a0348f-3bb7-7c70-8ec5-61176ce00591`, idle with exact carried ownership.
- Protected writer-owned diff: `internal/analyze/analyze.go +117/-15`, `internal/cli/command.go +157/-38`, `internal/cli/command_test.go +75/-0`, total `+349/-53`.
- Canonical runtime build passed in `25.6760326 s`.
- Canonical full-build validation passed with exit `0` in `700.1712706 s`, from `2026-08-25T03:25:06.0193408+07:00` to `2026-08-25T03:36:46.1906114+07:00`.
- Full-build validation includes maintenance `analyze . --force`; `700.1712706 s` is never build performance and never P1/P2 benchmark data. Do not conflate build, full-build validation, or an explicit benchmark capture.
- Runtime `1.2.8`, `73,764,864` bytes, SHA-256 `62920CBF15921EF8A6D2FAC671776BAA3C312EFEA1BEED53B721C4AFF5E1B6C5`.

## Main-Owned Ledger Sync

- Main corrected the remaining stale current-state text in `actual-status.md` once; historical R4/R5 and state-machine definitions were preserved.
- Validation after sync: exactly `30` benchmark `B1-P1A-OPnnn` parent rows; exactly `30` matching plan parent checkboxes; P1-A checked; OP001 active; zero dynamic child rows; zero production attempt rows; current attempt/evidence/benchmark rows explicitly none; four-ledger `git diff --check` PASS.
- `actual-status.md` R7 records the exact live read-only measurement assignment and keeps Architect/Coder locked.
- Do not turn this documentation sync into another audit/report/review loop.

## Live Visible Lane Hierarchy

- Continuous sole Guard: `01a03339-144d-7383-8cd1-72fcd5e5417c`, active governance-only. Retarget this exact Guard to the successor at official transfer; never duplicate or control it as a functional lane.
- Exactly one preserved read-only measurement delivery is active: `01a033a5-4700-7272-a28d-de3a71f58135` (`P6-E Measure Segment 04`). It started at `2026-08-25T05:04:15+07:00`; as of the final observation `2026-08-25T05:12:34.7703740+07:00`, first ACK/tool/artifact remains absent. Current authority is only the complete `OP001 resolution` child drill-down.
- The measurement lane must consume the accepted capture/source first, enumerate every real measured child with denominators/non-overlap/sum/residual, and prove the largest child's exact cause/source owner/complete call path. It cannot edit source/test/docs/ledgers, build, run analyze, or launch a child benchmark without a separate explicit `MEASUREMENT SLOT` follow-up.
- The lane may terminate with either a complete raw child table or `READY_FOR_CHILD_MEASUREMENT_SLOT` containing the smallest exact proposed invocation/owned bytes. Previous assignment to Segment 03 `01a033a5-414c-7721-864f-a8f4642d46e3` ended two empty turns totaling no user-message materialization, ACK, tool, artifact, or result; it returned idle and owns no measurement. Main then assigned the same exact problem to Segment 04 without overlapping functional work. Preserve the actual observed state at transfer.
- Revoked/idle old capture executor: `01a033a5-3a44-7c43-9aca-c85e3e932a0f`; never resume it.
- Successful capture executor `01a03597-6afb-73e1-bee6-4306836cce01` is idle after one successful capture; do not rerun it.
- Sole instrumentation writer `01a0348f-3bb7-7c70-8ec5-61176ce00591` is idle; do not duplicate it.
- Preserved Planner `01a033b7-9e01-7151-bdce-5d87173e1ac5` remains HOLD until a fresh child-specific Architect decision exists.
- Production Coder `01a03375-f616-7651-ba81-99bad6536746` remains STOP/HOLD.
- Build-only executor `01a034fe-f5d3-76c2-b58e-411bd33f3536` remains idle; do not rerun accepted build/full-build.
- No hidden functional agents; never `fork_thread`.

## Required Next Transition

1. Monitor the exact active measurement lane's transcript and actual commands. If it deviates, stop the deviation immediately.
2. If it returns `READY_FOR_CHILD_MEASUREMENT_SLOT`, verify the proposed boundary is complete/non-overlapping and is the smallest child-specific measurement needed. Grant at most one exact slot; do not permit a duplicate top-level capture, source/test edit, or build unless a new plan-authorized writer boundary is explicitly established.
3. If it returns a complete child table, consume raw artifacts once and independently verify every duration, denominator, child sum/residual, source boundary, largest-child cause, owner, and full call path. Do not accept prose alone.
4. Main immediately writes every real measured child as a dynamic benchmark row and exactly one matching nested unchecked plan child item, updates evidence and actual-status, and checks exact cardinality. No row may be omitted by elapsed-time size.
5. Only after the complete child list and exact largest-child cause/owner/call path are proven may Main open fresh visible Architect A001. Planner then refreshes the exact current attempt before production Coder.
6. P2-A remains one exhaustive implementation slice. Every actual attempt follows child/parent/total basis -> fresh Architect -> Main/planner refresh -> Coder -> canonical full build -> child/parent/E2E remeasurement -> fresh per-attempt Supervisor -> disposition.
7. `KEEP` requires lower child, retained parent, and retained end-to-end elapsed times plus that attempt's Supervisor `PASS`. `KEEP` promotes the baseline, resets that child's no-KEEP streak to zero, leaves the child unchecked, and continues the same child.
8. Three consecutive no-KEEP attempts terminalize only that child as `SYSTEM_CHARACTERISTIC`. Every smaller child and parent remains mandatory; concrete `BLOCKED` remains unchecked.

## Git And Workspace Boundary

- Current HEAD remains `8df19258fbcb18e14841bf0dae036400aa9b22a3`; commits `7c6869ba` and `8df19258` are not P3-C.
- Handoff-boundary snapshot has exactly `20` tracked dirty rows, zero staged rows, and exactly two untracked protected rotation handoffs: `rp_main_260825_041414_orchestration_rotation_handoff.md` and this file. No mutation was performed by the snapshot.
- Sole workspace `E:\Anvien`, saved project local, no worktree.
- Preserve all dirty/protected state, including the instrumentation files, build/package work, the Child 06A four ledgers, roadmap/predecessor/successor changes, and Owner deletion of `CONTRIBUTING.md`.
- This handoff remains uncommitted and protected. No reset, stash, checkout, broad cleanup, broad stage, push, alternate worktree, target repo, network, install/package scripts, C-drive workspace/read/write, shared/global/quarantined cache, or blind process termination/retry.
- Temporary/debug artifacts belong only under repo-local `E:\Anvien\.tmp`; do not broadly content-scan `.tmp`.

## Preserved Violations

Do not normalize, downgrade, omit, or justify any of these:

- predecessor hard miss `4.646s — HIGH — VERIFIED VIOLATION`;
- continuous-monitoring gap at least `255s — HIGH — VERIFIED VIOLATION`;
- Guard warm warning `40.891s — HIGH — VERIFIED VIOLATION`;
- predecessor successor-creation warm miss `146.891s — HIGH — VERIFIED VIOLATION`;
- earlier outgoing predecessor Main hard miss `5216.875s — HIGH — VERIFIED VIOLATION`;
- graph queries used after `status` without mandatory preceding `anvien analyze --force`: `HIGH — VERIFIED VIOLATION`; those graph results are not decision evidence;
- a prior Main transferred Main-owned ledger/progression responsibility to a worker and stalled the campaign roughly `03:38–03:48`: `HIGH — VERIFIED VIOLATION`;
- a prior Main coordinated passively and required repeated Owner intervention: `HIGH — VERIFIED VIOLATION`;
- predecessor successor creation missed its warm target by `54.777s`: `HIGH — VERIFIED VIOLATION`;
- this outgoing Main initially read `working-rules` before FULL RAW `AGENTS.md` after official transfer: `HIGH — VERIFIED VIOLATION`; corrective full raw startup later completed in exact order.

## Handoff Completion

At the exact transfer point, outgoing Main sends the official transfer with this detached seal, retargets the existing Guard, and transfers the live measurement-lane cursor. Successor verifies the detached seal once, reads every Mandatory Raw Startup source through EOF in order, inspects the live lane and handoff-relevant Git boundary once, then continues the first incomplete gate without restarting accepted work. Outgoing Main terminates immediately after transfer.
