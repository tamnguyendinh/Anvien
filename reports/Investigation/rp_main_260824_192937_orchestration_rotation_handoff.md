# Main Orchestration Rotation Handoff — Child 06A Plan Split

## Authority

- Outgoing Main: `01a03375-ef3c-73c1-972a-5532b262558d`.
- Incoming prepared Main: `01a033bf-1cec-7bf3-ab52-545915b3c3fe`.
- Outgoing official authority UUID: `01a0337a-a734-73c3-b9b8-2e3dea0b29d6`.
- Outgoing authority start: `2026-08-24T18:14:31.348+07:00`.
- Warm target: `2026-08-24T18:59:31.348+07:00`.
- Hard deadline: `2026-08-24T19:14:31.348+07:00`.
- First detected post-deadline clock: `2026-08-24T19:28:26.992+07:00`.
- Current rotation violation: warm miss `1,735.644s`; hard miss `835.644s` — `HIGH — VERIFIED VIOLATION`. Preserve without normalization or justification.
- Preserve earlier verified violations exactly: `52.457s`; `3.488s`; interrupted Main monitoring turn; passive wait-only orchestration; prior warm miss `36.626s`.
- Incoming authority time must be derived from the raw UUIDv7 of the exact `OFFICIAL AUTHORITY TRANSFER` follow-up, then warm `+45m` and hard `+60m`.

## Latest Owner Authority

The Owner has assigned complete execution authority below the Owner to Main. After discussion, Main must execute; Main must not ask for permission again or push continuation/workflow decisions back to the Owner. Main commands workflow and lanes. Lanes execute and report; their preferred workflow is not authority. Only a later direct Owner command on a concrete issue overrides Main.

Current exact task: split all executable P6-E work out of Child 06 into a standard four-file `Child 06A` plan positioned after Child 06 and before Child 07.

## Current Functional State

- Workspace: saved project `local-dc76d679b5f0e3c8ddbb8f9218f46f6d`, exact `E:\Anvien`, local, no worktree.
- P6-D is closed at `81163e39718b94a509e41114cada224e8f269e36`.
- Child 06 must end/close at that P6-D boundary; old P6-E executable content and performance-dependent closure move to Child 06A.
- Child 06A is the direct successor of Child 06; Child 07 follows only Child 06A.
- Child 06A contains exactly one implementation slice and one eventual implementation commit boundary. Logical order is `U1 -> U2 -> rebaseline/ownership -> U3 -> rebaseline -> conditional U4 -> dynamic Pareto -> final equivalence -> one final Supervisor -> detect/cleanup -> one commit`.
- If Pn-A/Pn-B/Pn-C are retained for template compliance: Pn-A is the sole final Supervisor; Pn-B performs exact cleanup without a second review; Pn-C runs detect, creates the one implementation commit, and hands to Child 07. Cleanup changing production/test invalidates Pn-A and returns to the owning unit.
- Do not preselect lane count from examples. Actual cost centers determine `N`; at most three measurement lanes run concurrently later.
- After every measurement/result: result -> JSON/Markdown report -> immediate evidence/benchmark/status record -> next command. No end-of-batch documentation.
- Current performance truth remains: A/A `0/20`; harness unsealed; no comparable baseline; U1 locked; no Child 06A implementation or speedup claim.
- Mandatory operational report: `E:\Anvien\reports\Investigation\rp_huong_dan_lam_viec_de_slice_p6e_tien_hanh_nhanh_hon.md`.

## Active Visible Lane

- Visible Planner: `01a033b7-9e01-7151-bdce-5d87173e1ac5`, host `local`.
- It is actively reading FULL RAW rules/templates/roadmap/Child 06 ledgers/Child 07 ledgers/mandatory report before patching.
- It has received the binding Main clarification above, including exact Child 06 closure and one-Supervisor/one-commit structure.
- At `2026-08-24T19:28:26.992+07:00`, no Child 06A file existed yet.
- Preserve this Planner and monitor its actual commands/writes. Do not merely await prose.
- Internal hidden planner `/root/p6e_child6a_planner` was a routing error, was interrupted, and produced no Child 06A file. Never reuse it.

## Held Lanes

- All Coder and measurement lanes remain HOLD. Do not resume them during this plan-only task.
- Old Worker `01a0319a-62c9-7650-ba6a-ef59dfcf1973` remains permanently STOPPED.
- Replacement Coder `01a03375-f616-7651-ba81-99bad6536746` remains preserved/HOLD.
- Preserve continuous Guard lineage: `01a03339-144d-7383-8cd1-72fcd5e5417c`; latest warning lineage includes `01a03373-96ae-70d2-94c4-ac5e631a2a72`. Do not duplicate or control Guard as a functional lane.
- Invalid fork `01a032d9-44a3-7fe2-b97b-8a8fe89fe8f3` is excluded permanently. Never use `fork_thread`.

## Current Files and Boundary

- No `06a` plan file existed at the last scoped check.
- Existing Child 06 four ledgers are already modified in the Owner worktree and must not be reverted or overwritten broadly:
  - `2026-07-28-06-ambient-external-resolution-and-diagnostics-plan.md`
  - `2026-07-28-06-ambient-external-resolution-and-diagnostics-evidence.md`
  - `2026-07-28-06-ambient-external-resolution-and-diagnostics-benchmark.md`
  - `2026-07-28-06-ambient-external-resolution-and-diagnostics-actual-status.md`
- Planner is authorized only within `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/` for this assignment.
- Expected plan edits: roadmap; Child 06 four ledgers; new Child 06A standard four-file set; Child 07 predecessor/opening references wherever present.
- Do not change code, reports, runtime, graph, target, benchmarks, or functional artifacts for this plan split.
- Do not commit or push.

## Mandatory Incoming Startup and Immediate Work

1. Read FULL RAW sequentially through EOF: `E:\Anvien\AGENTS.md`, working-rules, orchestration, governance-rule-guard, planner, all four planner templates, this handoff, roadmap, all four Child 06 ledgers, relevant Child 07 ledgers, and the mandatory P6-E work report. No summaries substitute.
2. Verify this handoff once at the genuine transfer boundary; do not create a report/hash audit loop.
3. Record exact incoming authority/warm/hard from the official transfer UUID.
4. Immediately monitor visible Planner `01a033b7-9e01-7151-bdce-5d87173e1ac5` actual commands/writes and enforce scope.
5. On handoff, inspect the scoped diff and verify: four new standard files; no template placeholders; Child 06 contains no executable P6-E section and closes at P6-D; roadmap order `06 -> 06A -> 07`; Child 07 predecessor points to 06A; legacy ID mapping is lossless; one Supervisor and one commit only; all links resolve; `git diff --check` passes for exact edited paths.
6. Route proportionate visible acceptance only if required by raw rules; do not open documentation audit loops or rerun unchanged functional gates.
7. Deliver the completed plan split without asking Owner to authorize continuation.

## Bans

No C-drive workspace/read/write, alternate worktree, target access, network, install/package scripts, protected/shared/global/quarantined cache, reset/stash/checkout, broad cleanup, broad stage, push, unscoped discovery, hidden functional agents, or `fork_thread`. Repo-local `.tmp` is disposable scratch only.
