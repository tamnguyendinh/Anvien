# Official Main Orchestration Rotation Handoff — Child 06A D002 Evidence Disposition

## Authority Transfer

- Date/time: `2026-08-27 14:25:08 +07:00`.
- Outgoing visible Main: task `01a0415f-7e14-7193-9620-70b09d021b79`.
- Successor visible Main: task `01a0421e-a38e-7192-8e98-8a09fa72f04d`.
- Transfer reason: the raw orchestration 120-minute Main-only rotation deadline has elapsed. This rotation preserves continuous campaign execution; it is not a pause, stop, workflow hand-back, or request for Owner direction.
- The outgoing Main must stop immediately after the successor acknowledges, this report records the exact successor task ID, and authority-transfer-complete is sent.

## Raw Rule Seal — Mandatory Before Successor Campaign Action

The successor's first response must be `UNDERSTOOD` or `NOT UNDERSTOOD`, then briefly state the goal, current slice, boundary, and first action. If `NOT UNDERSTOOD`, stop before every tool.

Read fully through EOF in this order:

1. `E:\Anvien\AGENTS.md`.
2. `E:\Anvien\.agents\skills\working-rules\SKILL.md`.
3. `E:\Anvien\.agents\skills\orchestration\SKILL.md`.
4. `E:\Anvien\.agents\skills\planner\SKILL.md` because Main owns mechanical current-state synchronization.
5. `E:\Anvien\docs\plans\2026-07-26-anvien-graph-accuracy-multi-plan\2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy\plan-rules.md`.
6. The four standard Child 06A ledgers in that directory, fully through EOF.
7. This handoff report fully through EOF.

Apply 100% of the raw, unmodified `AGENTS.md`, `working-rules/SKILL.md`, and `orchestration/SKILL.md`; no prompt, summary, memory, compacted context, or handoff substitutes for them.

An independent visible governance Guard already exists and reports through task/source `01a04169-e6b0-7780-89f8-5a44718042f8` (the earlier handoff also records independent Guard task `01a03c6d-d11d-7ae2-ab74-2fe91a757681`). Do not command, retarget, wait for, request ACK from, duplicate, or functionally use Guard. Guard warns; Main chooses workflow from raw rules and evidence.

## Mandatory Operating Corrections From Owner Experience

- Main is the executive: understand the plan, choose the next valid workflow, issue complete lane contracts, monitor actual behavior, verify boundaries, commit promptly, and advance without asking Owner to operate the state machine.
- A reply, correction, or status question is not a pause. Respond via commentary and continue. Only explicit `PAUSE` or `STOP` halts campaign work.
- Main handles small mechanical ledger/status changes itself with planner skill. Never reactivate Planner for a few stale pointers or turn documentation sync into an audit loop.
- A lane owns an independent outcome; a skill only supplies capability. Do not assign work to a lane whose function/ownership does not match.
- Main is not the functional worker. Attribution, architecture, implementation, measurement, and independent acceptance stay in visible specialist lanes. Main performs only identity, boundary, state, staging, commit, and transition checks.
- Do not call a Main-owned boundary check a Supervisor review or issue a Supervisor acceptance verdict. Actual acceptance requires a separate visible Supervisor.
- Never rerun a still-valid gate merely because Main rotated, context compacted, documentation changed, or a transport retry occurred.
- Every simple lane receives a hand-held contract: exact goal, slice, ownership, authority, inputs, paths, commands, output, validation, stop conditions, verdict, and next owner.
- Continuously monitor active lanes. Intervene immediately on scope drift or audit loops. Do not final/yield while a gate is active.
- Transport recovery is evidence-based: a task with `items=[]` and no ACK/tool/write can be revoked and replaced once; never leave two functional owners.
- Measurement truth stays target-separated. Elapsed wall-clock controls; CPU/profile data is causal only. Attribution is not a production attempt, streak event, disposition, or speedup.
- Raw 120-minute rotation remains binding. Preserve campaign continuity through compliant handoff; do not treat an earlier Owner sentence as a technical override of raw rules.

## Campaign Goal And Current Slice

- Goal: continue measurement-driven Child 06A optimization of real `anvien analyze` as far as evidence safely permits while preserving accepted accuracy, graph/output, persistence/readers, ordering, failure/lifecycle, and target-specific workload contracts.
- Current slice: `P2-A`.
- Active parent: unchecked `B1-P1A-OP001 resolution`.
- D001 `B2-P2A-A001-D001 resolve_calls`: checked terminal `EVIDENCE_EXHAUSTED`; unsuccessful streak exactly `2`; accepted A003 baseline preserved.
- D002 `B2-P2A-A001-D002 resolve_accesses`: active/unchecked.
- D003-D017: queued/unopened. P3 and Child 07: closed.

## Git And Durable Checkpoints

- Current HEAD: `b779651bbbc5f09248d8a3b22ac5b451a0a629ef` — `docs(plan): record D002 attribution blocker`.
- D001 transition commit: `9f804359992590cb914eb3e89c88145aca8c4a7a` — `docs(plan): terminalize D001 evidence exhaustion`.
- Prior succession report commit: `cc27fef6d2298ec5dd42cff0c7e57e0570b3ce38`.
- Accepted A003 implementation/measurement checkpoint: `b6bf45bce95323aa6b53b182edfea8628bd8b463`.
- WAL fix checkpoint: `0f3a572331dd23d17688886fcbfebeb7d37ee35d`.
- `b779651b` exact manifest: the four Child 06A ledgers plus `reports/Investigation/rp_child06a_d002_residual_attribution.md`.
- At `14:25`, one unowned/protected worktree modification exists: `internal/aicontext/skills/orchestration/SKILL.md`, `+3/-0`. It was not created by outgoing Main and was absent immediately after `b779651b`; preserve it, do not stage/revert/edit it, and do not let it block report-only lane work.
- Staged set was empty at the snapshot. The active Architect may create exactly one allowed untracked report after this snapshot.

## Accepted Target-Separated D002 Basis

| Target | D002 `resolve_accesses` | Parent | Analyzer | Process | Denominator |
|---|---:|---:|---:|---:|---|
| Cheapapp | `9.380783200 s` | `20.472602300 s` | `93.531974900 s` | `95.630648200 s` | `accesses=26042; files=887` |
| Restaurant Manager | `2.254679300 s` | `20.850792800 s` | `98.020546700 s` | `101.096911900 s` | `accesses=50554; files=1234` |

Never combine or average targets. CPU/profile samples explain cause only; they do not prove elapsed improvement.

## Completed D002 Attribution

- Superseded D002 task `01a04174-06b5-7ff3-92fe-7e65ed0803c1`: first turn completed substantial valid Anvien/profile/pprof work but failed with transport `429`; its continuation stayed empty. It was explicitly stopped and owns no result beyond the consumed checkpoint.
- Sole replacement attribution task: `01a04209-fd65-7142-b17a-ac081fd2bfee`.
- Durable report: `E:\Anvien\reports\Investigation\rp_child06a_d002_residual_attribution.md`.
- Evidence ID: `E2-P2A-D002ATTRIB1`.
- Exact verdict: `D002_RESIDUAL_ATTRIBUTION_BLOCKED`.
- Broad owner: current `internal/resolution/resolve.go::resolveAccess`; retained graph scope HIGH file / CRITICAL symbol (`6` symbols, `4` files, `4` modules, `32` processes).
- Cheapapp dominant child route: export-binding append/sort/order/JSON decode, already rejected A004 ownership.
- Restaurant dominant child route: unresolved outcome/diagnostic append/structured decode, already accepted A002/A003 ownership.
- Exact missing fact: no same materially attributed synchronous child symbol exists on both targets outside accepted A002/A003 and rejected A004 ownership. The common `resolveAccess` envelope and access/file counts are not an exact residual cause.
- The report records complete synchronous resolved, unresolved/diagnostic, TypeScript, export-binding, graph, outcome, relationship, reference, and post-loop projection paths.
- No D002 production cause/owner/direction, candidate, attempt, streak event, numeric change, disposition, checkbox, or queue transition was created.
- Main boundary check passed: sole report path before ledger sync, scoped diff-check PASS, staged empty. This was Main boundary verification, not Supervisor acceptance.
- Report/four-ledger checkpoint commit `b779651b` is complete and clean except the protected unrelated file above.

## Active Visible Architect — Monitor Continuously

- Task: `01a0421b-7800-7ee3-9d1b-9321d319ac76`.
- Turn: `01a0421b-7acc-7b51-82f8-e40ccab4c158`.
- Title: `D002 Evidence-Disposition Architect`.
- Latest wait cursor at handoff: `d570b3d2-b515-44b7-9a4c-17f7c5b9c4ff:1`.
- Status at snapshot: active/in progress; it has begun mandatory raw rule reading and has not yet produced a durable report or verdict.
- Allowed report: exactly `E:\Anvien\reports\system-architect\rp_system-architect_260827_by_gpt-5_child06a_d002_evidence_disposition.md`.
- It is decision-only: no Anvien/profile/source discovery rerun, build/test/target analyze/measurement execution, production direction, plan edit, staging, or commit.

Allowed terminal verdicts:

1. `ARCHITECT_D002_NO_FURTHER_MEASUREMENT_JUSTIFIED` — no bounded non-duplicative two-target evidence action remains; no production direction; exact unavailable evidence and strict reopen condition recorded.
2. `ARCHITECT_D002_ONE_BOUNDED_MEASUREMENT_READY` — exactly one finite hypothesis/symbol, both-target predicate, inputs, artifacts, output/conservation/parity/stop contract; not another broad profile/source/family-splitting loop.
3. `ARCHITECT_D002_EVIDENCE_DISPOSITION_BLOCKED` — one exact authority fact is missing even to choose between the first two; no generic audit proposal.

## Immediate Successor Workflow

1. After raw re-anchor, bounded-monitor Architect `01a0421b-7800-7ee3-9d1b-9321d319ac76` at cursor `d570b3d2-b515-44b7-9a4c-17f7c5b9c4ff:1`; do not duplicate or manually perform its architecture task.
2. Intervene if it broadens into source/profile/ledger audit or production design.
3. On verdict, perform one Main-owned boundary check only: exact report-only path, scoped diff-check PASS, staged empty, verdict matches report, no production direction, target separation preserved, no attempt/streak/number/checklist mutation.
4. Use planner skill directly for the smallest mechanical four-ledger sync. Do not open a Planner lane for this state update.
5. Create a scoped report/four-ledger docs checkpoint promptly; preserve the unrelated `internal/aicontext/skills/orchestration/SKILL.md` modification.
6. If `NO_FURTHER_MEASUREMENT_JUSTIFIED`: verify the five binding `EVIDENCE_EXHAUSTED` conditions, terminalize only D002 without manufacturing an attempt/streak event, activate unchecked D003, commit, and dispatch D003's exact next measurement/attribution lane from current benchmark authority.
7. If `ONE_BOUNDED_MEASUREMENT_READY`: materialize only that exact measurement contract in current ledgers, checkpoint it, open one visible measurement executor, and continuously monitor. No production Planner/Coder opens.
8. If `EVIDENCE_DISPOSITION_BLOCKED`: record the exact blocker once and determine the next valid workflow from raw rules/evidence; do not ask Owner to operate the plan and do not create another generic attribution loop.
9. Only explicit Owner `PAUSE`/`STOP` halts work. A question, correction, or status message is commentary-only and campaign work continues.

## Absolute Boundaries

- Do not rerun A001-A006, accepted D002 Anvien/profile/pprof, or the completed attribution.
- Do not alter accepted A003 production bytes or target-separated numbers.
- Do not infer a D002 production direction or open production Planner/Coder from the blocked attribution.
- Do not check D002 except through the exact binding `EVIDENCE_EXHAUSTED` proof; ordinary blocker remains unchecked and parent-blocking.
- Do not open D003-D017, P3, or Child 07 early.
- Do not touch/stage/revert the unrelated `internal/aicontext/skills/orchestration/SKILL.md` change.
- Do not command or duplicate Guard.
- Continue the raw 120-minute Main rotation cycle; initialize the next successor 15 minutes before deadline and transfer at deadline without campaign interruption.

## Successor First Unfinished Gate

ACK, complete raw re-anchor, then continuously monitor the active D002 evidence-disposition Architect to its exact durable verdict. No further Owner permission is required.
