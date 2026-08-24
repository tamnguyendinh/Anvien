# Main Orchestration Rotation Handoff — Child 06A P1-A Clean Canonical Build

Status: `DRAFT_UNSEALED`

## Authority Transfer

- Outgoing Visible Main: `01a034d7-07f8-7310-b224-48a90ba8c791`.
- Prepared clean Visible Main successor: `01a03504-5da4-79c2-a132-6aeab8568d36`, host `local`, saved project `E:\Anvien`, no worktree.
- Outgoing authority start: `2026-08-25T00:47:17.349+07:00`.
- Warm successor target: `2026-08-25T01:32:17.349+07:00`.
- Hard transfer deadline: `2026-08-25T01:47:17.349+07:00`.
- The successor was created and returned `UNDERSTOOD / WAITING_FOR_OFFICIAL_TRANSFER` before the warm target. It used no tool and took no authority.
- Authority remains exclusively with the outgoing Main until the successor receives a follow-up headed exactly `OFFICIAL AUTHORITY TRANSFER` citing the final detached seal of this report.
- After that official message, outgoing Main terminates and must not continue orchestration.

## Mandatory Raw Startup

The successor must read FULL RAW sequentially through EOF and must never substitute this handoff, memory, summaries, or compacted context for the sources below:

1. `E:\Anvien\AGENTS.md`.
2. `E:\Anvien\.agents\skills\working-rules\SKILL.md`.
3. `E:\Anvien\.agents\skills\orchestration\SKILL.md`.
4. `E:\Anvien\.agents\skills\governance-rule-guard\SKILL.md`.
5. `E:\Anvien\.agents\skills\planner\SKILL.md` and all four planner templates.
6. All five current Child 06A files under `E:\Anvien\docs\plans\2026-07-26-anvien-graph-accuracy-multi-plan\2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy`.
7. `E:\Anvien\reports\Investigation\rp_main_260824_204615_child06a_optimization_method_discussion.md`.
8. `E:\Anvien\reports\Supervisor\rp_supervisor_260824_223907_by_gpt-5_child06a_plan_rewrite_resubmission.md`.
9. Accepted raw capture artifacts under `E:\Anvien\.tmp\child06a_p1a_initial_20260824_225900`.
10. `E:\Anvien\reports\Investigation\rp_main_260824_234000_orchestration_rotation_handoff.md`.
11. `E:\Anvien\reports\Investigation\rp_main_260825_004000_orchestration_rotation_handoff.md`.
12. This handoff report.

AGENTS.md, working-rules, orchestration, and governance rules must be applied 100% raw, unmodified, and unsummarized. The successor must explicitly preserve this anti-summarization requirement in its next Main handoff. Re-anchoring must not restart the accepted P1-A capture, passed documentation review, doc checkpoint, completed impacts, completed graph refresh, or any later PASS recorded below.

## Exact Campaign Goal And Current Gate

Exact goal:

> Complete Child 06A by reducing the measured elapsed wall-clock time of the real `anvien analyze` graph-generation pipeline as far as current evidence safely permits, while preserving accepted accuracy, semantic completeness, graph correctness, output, persistence, readers, determinism, freshness, failure, transaction, temporary-file, and publication behavior; then close through exactly one final Supervisor, exact cleanup, one post-cleanup detect, exactly one implementation commit, and hand off Child 07.

Current plan:

`E:\Anvien\docs\plans\2026-07-26-anvien-graph-accuracy-multi-plan\2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy\2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy-plan.md`

Current gate:

- `P0-A` is complete.
- `P1-A` remains the first incomplete checklist item.
- Accepted initial capture remains valid and must not be rerun as if absent.
- `E1-P1A-INSTR1` remains open. Instrumentation source/test work exists, but a canonical clean full-build PASS and the like-for-like instrumented capture do not yet exist.
- `P2-A` is locked. There is no complete `B1-P1A-OPnnn` ranking, matching parent checklist, child list, Architect attempt, production Coder attempt, speedup, final equivalence, final Supervisor, detect, implementation commit, or Child 07 handoff.

## Accepted P1-A Capture — Immutable Baseline

- Capture: `child06a-p1a-initial-20260824-225900`.
- Process wall: `605.732722 s`.
- Analyzer internal total: `602.5278811 s`.
- Fifteen phase sum: `587.2385976 s`.
- Analyzer residual: `15.2892835 s`.
- Outer CLI/process residual: `3.2048409 s`.
- Largest provisional phase: resolution `532.4778258 s`, about `88.3740%` of analyzer internal total.
- Work/output denominator: `2,196` scanned, `765` parsed code, `0` failed, `123,075` nodes, `169,739` relationships, `20,363` dependency edges, `479` unresolved.
- The graph-freshness maintenance run remains excluded only because it lacked `--benchmark-json` and was not a like-for-like accepted capture.

## Plan Checkpoint And Worktree

- HEAD: `7b3bccb3b1abd09d1d2d3bdc9d864f55d9bd4965`, subject `docs(plan): checkpoint Child 06A optimization plan`.
- The checkpoint is doc-only and does not count as the one P3-C implementation commit.
- Worktree is intentionally dirty and protected. Preserve Owner deletion `CONTRIBUTING.md`, modified roadmap/Child 06/Child 07 ledgers, reports, raw coder artifacts, and all unrelated/protected rows.
- Current scoped instrumentation diff remains:
  - `internal/analyze/analyze.go`: `117` insertions / `15` deletions.
  - `internal/cli/command.go`: `157` insertions / `38` deletions.
  - `internal/cli/command_test.go`: `75` insertions / `0` deletions.
- No implementation commit exists.

## Global Anvien Architecture And Owner Process Direction

- Anvien is a repo-agnostic global CLI/MCP runtime. `anvien analyze [path]` targets a repository; the global MCP runtime can hold a packaged Anvien executable while the Anvien source repo itself is being rebuilt.
- Owner confirms no other repository/user currently needs Anvien. This campaign may terminate verified Anvien MCP/CLI holders to create a clean build/analyze window.
- For full build, the decisive sequence is exactly: `anvien clean --force` -> terminate remaining Anvien MCP/CLI processes and `anvien.cmd` wrappers -> canonical full build.
- `anvien clean` deletes the repository index; it does not release a held executable. Process holders must still be terminated separately.
- A legitimate full build takes about `12–15 minutes` because it includes the expensive real `anvien analyze` graph-generation step. Never kill a process that is a descendant of the active canonical build tree. Kill only unrelated Anvien holders outside that ancestry.
- Before an unfamiliar Anvien command, use `anvien --help` and the exact subcommand help. Do not replace command-surface use with broad source speculation.
- Owner forbids audit/documentation loops that block progress: a check must unlock the next action; a passed gate is not rerun unless invalidated; no report-about-report, hash/wording audit, or documentation acceptance loop.

## Canonical Build History In This Authority Window

### Build after holder cleanup — failed before analyze

- Command: `pwsh -NoProfile -File E:\Anvien\scripts\full-build.ps1`.
- Start: `2026-08-25T01:18:25.4775059+07:00`.
- End: `2026-08-25T01:23:32.6885829+07:00`.
- Exit: `1`.
- It passed the outer Go-vendor verification, then failed inside `go run ... package build-runtime` because nested `verify-go-vendor.ps1` could not resolve `Get-FileHash`. It never reached graph-generation analyze and did not update the canonical runtime.
- Artifacts: `E:\Anvien\.tmp\child06a_p1a_post_holder_build_20260825_011700`.
- Two unrelated MCP holders, PIDs `14004` and `14068`, were terminated; no legitimate build descendant was terminated.
- A bounded environment diagnosis subsequently proved the exact nested verifier can resolve `Get-FileHash` and passed at `2026-08-25T01:31:41.713+07:00`. A complex retry wrapper then failed before starting a build; it is not a build attempt, PASS, or product failure and must not become an audit loop.

### Owner-mandated clean/kill/canonical build — currently running

- `anvien clean --force` ran from `E:\Anvien`, exited `0`, and deleted `E:\Anvien\.anvien`.
- Process cleanup completed at `2026-08-25T01:34:35.1291488+07:00`; `remaining_count=0`, and there were no live PIDs to terminate in that exact check.
- Exact command: `pwsh -NoProfile -File E:\Anvien\scripts\full-build.ps1`.
- Start: `2026-08-25T01:34:53.276287+07:00`.
- Wrapper PID: `7044`.
- Full-build PID: `12520`.
- Current observed stage: outer Go-vendor verification `PASS`; packaged Go runtime build is running.
- This invocation is the sole valid build. Let it run for the legitimate `12–15 minute` duration. Do not start a competing build/analyze and do not kill its descendants.
- If PASS: verify the canonical build artifact once, direct sole writer to update the four P1-A ledgers and return `READY_FOR_INSTRUMENTED_CAPTURE`, then reactivate exactly the preserved measurement executor for one like-for-like instrumented capture.
- If FAIL: capture the exact new failure and debug only that bug. Do not restart a broad holder/architecture/audit loop.

## Exact Next Cursor After Build

1. Monitor only the current canonical build until it exits.
2. On PASS, self-verify the exact command/exit/provenance and current scoped diff once; then direct sole writer to record build PASS and handoff `READY_FOR_INSTRUMENTED_CAPTURE` in the four standard ledgers.
3. Reactivate exactly measurement executor `01a033a5-3a44-7c43-9aca-c85e3e932a0f` for one like-for-like instrumented capture. Before capture, use the Anvien command surface and preserve the accepted cache/workload regime. `anvien clean --force` is used only if the accepted comparison requires that same clean-index regime; never silently mix cache/instrumentation states.
4. Complete `E1-P1A-INSTR1` with exactly one disposition: carry exact instrumentation ownership into the first refreshed P2-A attempt, or remove exact owned bytes, canonical full-build, and re-establish/re-measure the accepted basis.
5. Populate the complete `B1-P1A-OPnnn` list for every measured top-level operation and the exact matching unchecked parent checklist before opening P2-A.
6. Then execute the required per-attempt chain: complete selected-parent child list -> fresh Visible Architect -> Planner refresh -> Coder -> canonical build -> child/parent/E2E remeasurement -> fresh Visible Supervisor -> disposition.

## Binding Optimization Method

- Elapsed wall-clock time controls ranking and speedup. CPU/RAM/allocation/GC/I/O/wait/count evidence is secondary only.
- P1-A cannot close until every real measured top-level operation has one `B1-P1A-OPnnn` row and one matching unchecked parent checkbox.
- P2-A is exactly one implementation slice and exhausts every measured parent and every measured child largest-first, including small rows.
- Every production attempt requires a distinct fresh Architect decision, Planner refresh, Coder, canonical full build, selected-child/parent/end-to-end remeasurement, fresh Visible Supervisor, and disposition.
- `KEEP` requires lower selected-child elapsed time, lower retained parent elapsed time, lower retained end-to-end elapsed time, and Supervisor `PASS`.
- Three consecutive attempts without `KEEP` terminalize only that child as `SYSTEM_CHARACTERISTIC`; a concrete blocker remains unchecked.
- P3-A is the one final whole-candidate Supervisor. P3-B changes no accepted production/test bytes. P3-C performs the only post-cleanup detect, the only implementation commit, and Child 07 handoff.

## Visible Lane State

- Prepared successor `01a03504-5da4-79c2-a132-6aeab8568d36`: idle, `WAITING_FOR_OFFICIAL_TRANSFER`, no tool use.
- Sole instrumentation writer `01a0348f-3bb7-7c70-8ec5-61176ce00591`: preserved sole source/test writer; do not duplicate. It must not start a competing build while the current build executor owns the invocation.
- Canonical build executor `01a034fe-f5d3-76c2-b58e-411bd33f3536`: ACTIVE on the sole current build; no source/doc edit authority.
- Initial measurement executor `01a033a5-3a44-7c43-9aca-c85e3e932a0f`: HOLD, preserved for the one like-for-like instrumented capture after exact build/ledger gate.
- Planner `01a033b7-9e01-7151-bdce-5d87173e1ac5`: HOLD, preserved for P2-A refresh after a fresh Architect decision.
- Production Coder `01a03375-f616-7651-ba81-99bad6536746`: HOLD, never use for P1-A instrumentation.
- Continuous Guard lineage `01a03339-144d-7383-8cd1-72fcd5e5417c`: ACTIVE, governance-only; preserve/retarget through official transfer and never duplicate/control it as a functional lane.
- Other measurement lanes remain HOLD. At most three read-only analysts may be ACTIVE only for current independent measured problems. No duplicate competing benchmark process.

## Main Authority And Role Boundary

- User commands directly. Main has complete command authority below User and must advance without asking for repeat authorization.
- Functional lanes execute and report; they do not choose workflow or command Main.
- Main performs orchestration, monitoring, boundary verification, planner-ledger transitions, acceptance routing, and commit transition; it does not duplicate production or measurement work.
- No hidden functional agents. Never use `fork_thread`.
- Current Child 06 has no current Pn-A/B/C. Legacy `E6-PNA/PNB/PNC-*` is provenance only. Child 06A P3-A/P3-B/P3-C is the current closure sequence.

## Workspace, Git, And Bans

- Sole workspace: `E:\Anvien`, saved project local, branch `master`, no worktree.
- Repository-local `E:\Anvien\.tmp` is the only scratch location.
- No reset, stash, checkout, broad cleanup, broad stage, push, C-drive workspace/read/write, alternate worktree/target, network, install/package scripts, or protected/shared/global/quarantined cache.
- Do not run detect-changes before P3-C; the plan reserves one post-cleanup detect there.
- Preserve dirty/protected state and do not normalize unrelated work.

## Preserved Verified Violations

Preserve each item exactly without justification, normalization, downgrade, or omission:

- `52.457s`.
- `3.488s`.
- interrupted Main monitoring turn.
- passive wait-only orchestration.
- prior warm miss `36.626s`.
- prior current warm miss `1,735.644s`.
- prior current hard miss `835.644s`.
- predecessor warm miss `30.858s — HIGH — VERIFIED VIOLATION`.
- predecessor hard miss `4.646s — HIGH — VERIFIED VIOLATION`.
- continuous-monitoring gap at least `255s — HIGH — VERIFIED VIOLATION`.
- Guard warm warning `40.891s — HIGH — VERIFIED VIOLATION`.
- successor-creation warm miss `146.891s — HIGH — VERIFIED VIOLATION`.
- graph-query-without-refresh `HIGH — VERIFIED VIOLATION`.

The four graph-query results implicated by the last item are excluded from evidence and decision-making. No further graph-based command may run without the required fresh `anvien analyze --force`; the current clean canonical full build may itself provide the next legitimate graph refresh only if it reaches and completes that analyze step.

## Detached Seal

`UNSEALED_DRAFT` — replace this section only after the final write, then compute exact byte length and SHA-256 once for the official transfer message.
