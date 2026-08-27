# Official Main Orchestration Handoff — Child 06A D001 Evidence Exhaustion To D002

## Authority Transfer

- Date/time: `2026-08-27 10:55:21 +07:00`.
- Outgoing visible Main: task `01a03cd1-6dd7-7fc1-b744-ab73e9cc61e8`.
- Successor visible Main: task `01a0415f-7e14-7193-9620-70b09d021b79`.
- Transfer authority: latest direct Owner command, `hand off cho main moi lam viec ngay lap tuc`.
- This command requires the current transfer now. It is not blanket authority for later automatic rotation. After activation, the successor must not rotate/handoff again without a new direct Owner command.
- The outgoing Main must stop immediately after the successor acknowledges and becomes active.

## Raw Rule Seal — Mandatory Before Successor Tools

The successor must reply `UNDERSTOOD` or `NOT UNDERSTOOD` before any tool, then briefly state the goal, current slice, boundary, and first action. If `NOT UNDERSTOOD`, stop.

Before campaign action, read fully through EOF in this order:

1. `E:\Anvien\AGENTS.md`.
2. `E:\Anvien\.agents\skills\working-rules\SKILL.md`.
3. `E:\Anvien\.agents\skills\orchestration\SKILL.md`.
4. `E:\Anvien\.agents\skills\planner\SKILL.md`, because an active visible Planner is materializing the next control transition.
5. Binding `E:\Anvien\docs\plans\2026-07-26-anvien-graph-accuracy-multi-plan\2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy\plan-rules.md`.
6. All four standard Child 06A ledgers in that directory, fully through EOF.
7. This handoff report fully through EOF.

Apply 100% of the raw, unmodified `AGENTS.md`, `working-rules/SKILL.md`, and `orchestration/SKILL.md` at all times. This report, summaries, memory, and compacted context are not substitutes for those sources.

Independent Guard task `01a03c6d-d11d-7ae2-ab74-2fe91a757681` already exists. It is independent. Main must not command, retarget, request ACK from, wait for, duplicate, or functionally use it. Guard warnings are compliance evidence; Guard does not choose Main workflow.

## Campaign Goal And Current Slice

- Goal: continue measurement-driven Child 06A optimization of real `anvien analyze` as far as evidence safely permits, preserving accuracy, graph/output, persistence/readers, ordering, failure/lifecycle, and target-specific workload contracts.
- Current implementation slice remains `P2-A`.
- Current parent: `B1-P1A-OP001 resolution`, unchecked.
- Current child at entry: `B2-P2A-A001-D001 resolve_calls`, unchecked at the committed blocker checkpoint.
- Transition being materialized now: terminalize D001 as narrowly defined `EVIDENCE_EXHAUSTED`, then activate `B2-P2A-A001-D002 resolve_accesses` for read-only residual attribution.
- D002 must remain unchecked. D003-D017 stay queued/unopened. Parent remains unchecked. P3 and Child 07 remain closed.

## Git And Durable Checkpoints

- Current HEAD at handoff start: `77feee35585ce3519867614568845979caa2ff83` — `docs(plan): record A006 evidence blocker`.
- Parent checkpoints:
  - A003 accepted implementation/measurement checkpoint: `b6bf45bce95323aa6b53b182edfea8628bd8b463`.
  - WAL force-fix checkpoint: `0f3a572331dd23d17688886fcbfebeb7d37ee35d`.
  - Prior M2 docs checkpoint: `506490011c190fbeeaf81db8c2a13da09d87420a`.
- Worktree/staged snapshot at `2026-08-27 10:55:21 +07:00`: clean; staged set empty.
- The active Planner may create exactly six allowed modifications after this snapshot. Inspect current reality once when receiving its handoff; do not assume the snapshot remains clean.

## Accepted Optimization Result

A001, A002, and A003 are `KEEP`. A004 and A005 are `SUPERVISOR_PASS / NO_KEEP / ROLLBACK_COMPLETE`. A006 M1/M2 are attribution-only and never created a production attempt or streak event.

Accepted A003 values remain separate and must never be averaged or combined:

| Target | D001 | Parent resolution | Analyzer | Process | Denominator |
|---|---:|---:|---:|---:|---|
| Cheapapp | `3.447846300 s` | `20.472602300 s` | `93.531974900 s` | `95.630648200 s` | `calls=27890; files=887` |
| Restaurant Manager | `9.401585300 s` | `20.850792800 s` | `98.020546700 s` | `101.096911900 s` | `calls=86030; files=1234` |

Initial-to-A003 process reduction:

- Cheapapp: `890.314783200 -> 95.630648200 s`, reduction `794.684135000 s` (`89.3%`, about `9.3x`).
- Restaurant Manager: `1178.391336900 -> 101.096911900 s`, reduction `1077.294425000 s` (`91.4%`, about `11.7x`).

## A006 Technical Closure And Current Governance Fact

Durable evidence:

- Architect report: `E:\Anvien\reports\system-architect\rp_system-architect_260827_by_gpt-5_child06a_a006_residual_direction.md`.
- Main handoff report: `E:\Anvien\reports\Supervisor\rp_supervisor_260827_100511_by_gpt-5_child06a_a006_post_m2_architect_handoff.md`.
- Exact verdict: `ARCHITECT_A006_NO_FURTHER_MEASUREMENT_JUSTIFIED`.
- Main handoff verification: `PASS`.
- Current committed cursor: `A006_NO_FURTHER_MEASUREMENT_JUSTIFIED / D001_BLOCKED_OPEN / EVIDENCE_PROVEN_TWO_TARGET_OWNER_UNAVAILABLE / D001_STREAK_2`.
- M2 Cheapapp receiver recheck/control/net: `10,466,100 / 1,000,600 / 9,465,500 ns`, false/true `3549/0`.
- M2 Restaurant receiver recheck/control/net: `0 / 0 / 0 ns`, false/true `638/0`.
- Both packets pass `30/30`, `17/17`, conservation, parity, and overlap `0`; cross-target gap `213,628,851,200 ns`.
- The duplicate receiver-claim read is therefore falsified as a two-target production owner.
- No remaining evidence-proven, two-target, synchronously removable D001 owner exists outside A001-A005.
- M3, another split, repeated attribution, a manufactured third attempt, direct Planner/Coder, and `SYSTEM_CHARACTERISTIC` are forbidden from this evidence.

## Active Visible Planner — Must Be Monitored Continuously

- Fresh Planner task: `01a04153-b680-7e63-ae60-7108fe00278c`.
- Active turn at handoff: `01a04153-b910-7591-9132-860e68ae92b4`.
- Title: `A006 Planner — Evidence Exhaustion Transition`.
- Shared checkout: saved project `E:\Anvien`, local/direct, no worktree.
- Current observed progress: raw authority read through EOF; Planner stated the technical facts correctly and was about to run `anvien --help` followed by exactly one `anvien analyze --force` for doc-plan authoring.
- Superseded transport-failed Planner task `01a03f86-cd96-7573-97b1-52d9548768fb` ended idle with `items=[]`, no ACK/tool/write. It was explicitly told STOP and owns no work. Do not reuse it.

Planner exact required outcome:

1. Define a narrow binding `EVIDENCE_EXHAUSTED` terminal applicable only when the accepted baseline is preserved, a fresh attempt-local Architect proves no further bounded measurement is justified and releases no production direction, Main handoff verification passes, no attempt/streak event is manufactured, and an exact unavailable-evidence/resume condition is recorded.
2. Explicitly distinguish it from `KEEP`, `NO_KEEP`, `ROLLBACK`, `SYSTEM_CHARACTERISTIC`, a third attempt, a speedup, or an accuracy waiver.
3. Record `E2-P2A-A006EXHAUST1` from `A006ARCH3/A006MAINVERIFY2/A006BLOCK1`.
4. Transition `D001_BLOCKED_OPEN -> D001_EVIDENCE_EXHAUSTED_TERMINAL`.
5. Preserve A003 numeric values and D001 streak exactly `2`.
6. Check only D001; keep the parent unchecked.
7. Activate D002 unchecked with cursor `D002_CURRENT_BASIS_RECORDED / D002_RESIDUAL_ATTRIBUTION_PENDING / ARCHITECT_LOCKED`.
8. Keep D003-D017 queued/unopened and P3/Child 07 closed.
9. Change no measured number, denominator, target separation, Attempt Numeric History, A001-A006 disposition, accepted baseline, or streak.
10. Write exactly these six paths:
   - `plan-rules.md`;
   - four standard ledgers;
   - `E:\Anvien\reports\planner\rp_planner_260827_by_gpt-5_child06a_a006_evidence_exhaustion_transition.md`.
11. Return exact verdict `PLANNER_A006_EVIDENCE_EXHAUSTION_READY_FOR_MAIN_VERIFY`, scoped diff-check PASS, and staged-empty.

Successor first action after raw re-anchor: monitor this Planner with bounded `wait_threads`. Do not open D002 early. When it completes, perform one Main-owned boundary verification only:

- exact six-path set;
- scoped `git diff --check` PASS;
- D001 only checked; parent/D002/D003-D017 states correct;
- every measured number and D001 streak `2` unchanged;
- exact `E2-P2A-A006EXHAUST1` and D002 cursor present;
- staged set empty.

Do not open a Supervisor/doc-audit loop for this plan transition. If exact boundary passes, stage the six paths and create a docs/report checkpoint immediately. Suggested subject: `docs(plan): terminalize D001 evidence exhaustion`. Verify manifest and clean worktree, then dispatch D002 attribution without waiting for Owner approval.

## D002 Attribution Packet — Ready For Immediate Dispatch After Commit

The successor must open one fresh separate visible read-only attribution task in the same saved project/local shared checkout. Do not use an internal hidden worker or Main as functional investigator.

Exact D002 accepted current basis from the A003 raw packets:

| Target | D002 `resolve_accesses` | Parent | Analyzer | Process | Denominator |
|---|---:|---:|---:|---:|---|
| Cheapapp | `9.380783200 s` | `20.472602300 s` | `93.531974900 s` | `95.630648200 s` | `accesses=26042; files=887` |
| Restaurant Manager | `2.254679300 s` | `20.850792800 s` | `98.020546700 s` | `101.096911900 s` | `accesses=50554; files=1234` |

Accepted raw benchmark inputs:

- `E:\Anvien\.tmp\child06a_a003_cheapapp_frozen_17child_20260826\candidate\benchmark.json`.
- `E:\Anvien\.tmp\child06a_a003_restaurant_manager_frozen_17child_20260826\candidate\benchmark.json`.

Accepted CPU profiles:

- Cheapapp: `E:\Anvien\.tmp\child06a_a003_cheapapp_frozen_17child_20260826\candidate\resolution.cpu.pprof`, `65,363` bytes, SHA-256 `806FAD79B12C567618778A8232E8521991882CE6C745AC2EBBE7B09ECC267107`.
- Restaurant: `E:\Anvien\.tmp\child06a_a003_restaurant_manager_frozen_17child_20260826\candidate\resolution.cpu.pprof`, `56,397` bytes, SHA-256 `4272FFAB0ABE4C14C470A139A31A601F31323B1373E0FA2986A09935456FA3E2`.

Existing measured child inventory/source-map input:

- `E:\Anvien\.tmp\child06a_p2a_op001_resolution_source_map_segment03\source_map.json`.
- `resolve_accesses` owner inventory: `internal/resolution/resolve.go::resolveAccess`.
- Measured interval: `ResolveBoundInto` per-file `ir.Accesses` loop, inventory lines `95-97`.
- Inventory source-owner range: `resolveAccess`, lines `630-741` at that source-map basis; the attribution lane must obtain current exact lines from fresh Anvien/current source and must not blindly reuse old line numbers.
- Complete product path inventory: `internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> for each ir in w.files -> for each access in ir.Accesses -> resolution.resolveAccess`.

Lane capabilities: `working-rules` plus `debugging`; report-only read-only functional attribution. Require the lane to read full raw files itself.

Exact lane commands/sequence:

1. `anvien --help`.
2. Exactly one fresh `anvien analyze --force` from `E:\Anvien`.
3. Task-matching Anvien evidence after refresh:
   - `anvien file-detail E:\Anvien\internal\resolution\resolve.go --repo E:\Anvien --json`;
   - `anvien context symbol "resolveAccess" --repo E:\Anvien`;
   - `anvien impact symbol "resolveAccess" --repo E:\Anvien --direction upstream` for exact owner/caller/flow scope; this is read-only attribution evidence, not edit permission.
4. Verify the two profile hashes above before consumption.
5. Run bounded profile queries independently per target, focusing on stacks containing `resolution.resolveAccess`; use the installed Go toolchain and record exact commands/output. Do not add/average target samples and do not use CPU samples as elapsed proof.
6. Read only the current source ranges and direct helpers needed to explain the retained D002 stack after Anvien owner discovery. Do not broadly audit resolution.
7. Prove, or explicitly fail to prove, all three required facts:
   - exact current residual cause materially present on both targets;
   - exact current production owner distinct from already accepted/rejected directions where applicable;
   - complete call path from CLI/analyze through D002 and all synchronous downstream graph/outcome/diagnostic/reference carriers affected by that cause.
8. Do not design the fix or select architecture. Do not edit source/test/plan/ledger/script/target. Do not run build/test/target analyze/new profile/benchmark/detect/stage/commit.
9. Write only `E:\Anvien\reports\Investigation\rp_child06a_d002_residual_attribution.md` via `apply_patch`.
10. Minimum validation: exact report-only changed-path boundary; `git diff --check -- reports/Investigation/rp_child06a_d002_residual_attribution.md`; staged set empty.
11. Terminal verdict exactly one of:
   - `D002_RESIDUAL_ATTRIBUTION_COMPLETE_READY_FOR_ARCHITECT` when cause, exact owner, and complete call path are all proven;
   - `D002_RESIDUAL_ATTRIBUTION_BLOCKED` with the one exact missing fact/evidence, without inventing another measurement workflow.
12. Next owner: Main Orchestration. Architect remains locked until the complete verdict exists.

After a complete attribution handoff, Main performs one zero-trust boundary check, records the packet in current ledgers using planner skill, and opens exactly one fresh visible D002 Architect. Do not ask Owner to choose the workflow.

## Learned Operating Habits — Mandatory Successor Behavior

These are not new rules; they are the practical operating form of rules repeatedly corrected by the Owner during this campaign.

### 1. Main is the executive, not an employee waiting for instructions

- Main understands the whole campaign, chooses the next valid workflow, issues complete commands, monitors, verifies boundaries, commits, and advances immediately.
- Never ask Owner to decide a workflow that repo rules/evidence already determine.
- Never push the responsibility to advance the plan back to Owner.
- Ask only when new authority, an external decision, or a truly material product/architecture choice cannot be derived from current authority.

Why this habit is necessary: this campaign stalled whenever Main treated Owner as the person who must design the next workflow. The clearest failure was after A006 was recorded as `D001_BLOCKED_OPEN`: Main reported the blocker and asked Owner to approve a new terminal category instead of using its executive responsibility to reconcile the plan method and keep the measured queue moving. Owner had already supplied the goal, rule hierarchy, and authority boundary; asking Owner to operate the state machine forced Owner to become the orchestrator. The successor must therefore distinguish two questions. “What does the evidence technically prove?” must be answered from source and durable evidence. “What valid workflow consumes that fact?” is Main's responsibility unless the workflow would expand authority beyond the campaign.

### 2. A reply is not a stopping point

- Owner reminder, question, correction, or status request is not PAUSE.
- Reply through commentary and continue seamlessly in the same turn.
- Only explicit PAUSE/STOP halts functional orchestration.
- If waiting for a lane, continuously monitor with bounded waits. Do not final/yield and go idle while a gate is active.

Why this habit is necessary: several Owner messages asked for status, clarification, or correction while the campaign still had an active gate. Main sometimes produced a polished final response and ended the turn, even though no PAUSE/STOP existed. That changed a normal conversation into a silent operational stop, and the Owner had to send “continue” repeatedly. The correct behavior is to answer the question in commentary, preserve the active workflow, and keep calling bounded monitoring or doing safe Main-owned preparation. A final answer is appropriate only when the requested outcome is actually complete, the campaign has a real recorded blocker with no authorized continuation, or the Owner explicitly pauses/stops it.

### 3. Operate continuously, quickly, and accurately

- Work rolls forward in overlapping safe boundaries: while a functional lane works, prepare the exact next packet or complete unrelated Main-owned coordination.
- As soon as one gate returns the required durable result, update state once, commit at the applicable boundary, and open the next valid lane immediately.
- Do not wait for ceremonial reviews, hashes, or wording checks that are not real gates.
- Keep Owner updates concise, concrete, and evidence-labeled: verified, checking, no evidence, blocked.

Why this habit is necessary: Child 06A contains expensive operations—full builds, target analyzes, profiles, and visible specialist lanes. Serializing every preparatory action after the previous lane finishes creates large avoidable gaps. At the same time, unsafe overlap in the same files destroys attribution. “Continuous” therefore means overlap only non-conflicting work: while Planner reads/writes its six documents, Main may prepare the future D002 attribution packet from accepted artifacts, but Main may not edit those documents or open D002 before the transition commit. This gives speed without mixed ownership. “Accurate” means that every early preparation is explicitly provisional and no state transition is claimed before its gate exists.

### 4. Never rerun what already passed

- Re-anchor after compaction restores context only.
- Passed A001/A002/A003, WAL, A004/A005 review/rollback, and A006 M1/M2 remain valid unless their exact source/evidence/boundary is invalidated.
- Existing numbers, profiles, reports, hashes, and source maps are inputs. Consume them; do not recreate them merely to confirm them.
- Rerunning a valid gate, auditing a passed conclusion, or creating a confirmation loop is prohibited wasted work.

Why this habit is necessary: accepted target analyzes take roughly one to twenty minutes and can vary with real machine conditions. Rerunning them merely to reconfirm a durable packet wastes time and may create a different but still valid timing, which then tempts the campaign to audit natural variation. The same applies to builds and Supervisor reviews. A gate is rerun only when its exact source bytes, workload, executable, boundary, or evidence validity has been invalidated. Compaction, a new Main, a documentation edit, or discomfort with a result does not invalidate it. This is why the successor must consume the A003 profiles and D002 child row directly instead of launching new target analyzes.

### 5. Distinguish a wrong process from a valid result

- A rule violation in how evidence was obtained does not automatically make the result false.
- Stop the invalid behavior, preserve the existing result as an input, and continue from it when its technical evidence remains valid.
- Do not rerun correct data simply to repair ownership/procedure. Correct the procedure at the next transition.
- Conversely, Owner direction controls authority/workflow but is not automatically technical truth; verify technical claims against source/evidence and never say Owner “validated” a technical fact unless Owner actually did.

Why this habit is necessary: during A003 residual attribution, Main crossed its role boundary by manually duplicating functional pprof/source work. That process was wrong, but the resulting data did not become false merely because the wrong role obtained it. Repeating the profile analysis only to repair ownership would create a loop and potentially produce needless variance. The correction was to stop the improper behavior, treat the existing facts as inputs, restore the correct lane chain, and continue. The opposite error also occurred when Main said Owner had technically confirmed evidence, although Owner had only ordered that valid existing results not be rerun. Authority statements and technical proof must remain separate.

### 6. Documentation follows work; it does not become work

- Plan/ledger documentation is a continuously updated control surface, not an audit phase.
- If one word/row/cursor is stale, fix exactly that place once and continue.
- Do not open documentation audit lanes, report-about-report, hash-audit loops, or repeated cross-checks.
- Update plan/evidence/benchmark/actual status as state changes, often in parallel with the active lane, so every worker receives synchronized authority.

Why this habit is necessary: stale plan cursors can cause a real ordering error—for example, opening Architect or Coder against `ARCHITECT_LOCKED`—so documentation cannot be ignored. But repeatedly reviewing all four ledgers after changing one status word is equally harmful and was a recurring source of delay. The correct pattern is one precise update at the transition: change the exact status, evidence pointer, benchmark control wording, and actual-status row affected; run the smallest syntax check; then continue. Documentation proves coordination state, not the technical result itself. A documentation correction never justifies rerunning the underlying Architect, measurement, or Supervisor gate.

### 7. Lane and skill are different layers

- A lane owns an outcome. A skill gives that lane capability.
- Planner lane creates/translates implementation plans. Main using planner skill only maintains progress/current status; it does not replace the visible Planner for plan creation/technical translation.
- Architect owns architecture decisions. Main verifies authority/boundary and passes the architecture to Planner; Main does not audit whether architecture is “good enough” as a substitute architecture lane.
- Coder implements the exact Planner authority. Supervisor alone accepts functional completion. Main decides campaign disposition/transition from accepted evidence.
- Do not open Coder before the visible Planner completes. Do not give any lane stale, Main-authored, or mixed authority.

Why this habit is necessary: Main previously used the planner skill and then treated its own A003 plan text as if a visible Planner lane had created the implementation translation. It opened a Coder before the actual Planner completed, so the Coder received stale authority even though a pre-edit lock prevented source mutation. The problem was not lack of planner knowledge; it was wrong outcome ownership. A skill answers “what capability is available?” A lane answers “who owns the independently observable result?” The successor must check both before every transition. Loading a skill never grants Main the lane's functional authority.

### 8. Give every simple execution lane a complete hand-held contract

Every lane prompt must state:

- exact goal and open slice;
- ownership, capability, authority, scope, and non-goals;
- exact inputs and hashes when relevant;
- exact source/artifact paths;
- exact command/script and invocation;
- exact output root and report path/template;
- minimum validation;
- explicit stop conditions;
- completion marker and next owner.

Do not make a simple lane discover tools, invent scripts, design a workflow, or wait passively for Main clarification.

Why this habit is necessary: several early measurement-segment turns either materialized empty or waited without producing usable artifacts because the assignment described an outcome but did not give the exact script, files, schema, output root, and stop rule. Once Main supplied concrete source ranges, ordered child IDs, JSON schema, command, hashes, validation, and handoff marker, the lanes produced deterministic artifacts. A specialist should spend its time executing its specialty, not reverse-engineering Main's workflow. Complete lane contracts also make scope drift visible immediately: any command or file outside the packet is objectively wrong rather than subject to interpretation.

### 9. Main is not the functional worker

- Main performs Main-owned identity, boundary, state, staging, commit, and transition checks.
- Functional attribution, architecture, implementation, measurement, and independent acceptance belong to separate visible lanes when they create an independently observable outcome.
- Internal agents are only for bounded read-only inventory/discovery; their outputs are inputs, never acceptance verdicts.
- Never duplicate a functional lane's task manually while that lane is active.

Why this habit is necessary: when Main manually performed residual pprof/source synthesis while hidden discovery workers were doing the same job, it erased independent ownership and made the gate invisible to Owner. It also created self-review: Main both produced the technical conclusion and decided the transition based on it. Main must understand enough source/evidence to design and boundary-check a lane, but understanding is not permission to close the lane's functional outcome. Keeping Main at the executive layer preserves traceability, Owner control, and independent acceptance while allowing Main to act quickly on the returned result.

### 10. Commit boundaries preserve rollback and progress

- After an accepted implementation slice/checkpoint, stage only owned paths and commit promptly so rollback is possible.
- Docs/report checkpoint commits are allowed whenever their current checkpoint is complete and scoped.
- `P3-C` has a local final detect/closure-commit rule; it is not a blanket ban on earlier P2 progress checkpoints.
- Implementation commit requires fresh Anvien detect. Docs-only commit does not require implementation detect.
- Always verify exact staged manifest and finish with staged/worktree cleanliness before advancing.
- Do not broadly reset, checkout, stash, or clean shared work.

Why this habit is necessary: the Owner correctly emphasized that work without checkpoints is difficult to roll back safely. Earlier wording in `plan-rules.md` was misread as banning all progress commits until P3-C, which would have left multiple accepted optimizations entangled. The corrected rule distinguishes accepted P2 progress checkpoints from P3-C's one local closure commit. A commit is not ceremonial; it freezes a known accepted baseline so a later rejected attempt can be reversed surgically. Conversely, unaccepted A004/A005 candidates were rolled back without promotion commits. The successor must commit by actual boundary, not by phase-name keyword.

### 11. Anvien discipline is task-driven

- Run `anvien analyze --force` before graph-based work.
- Use task-matching Anvien query/context/file-detail/impact rather than raw grep as the discovery authority for unfamiliar code.
- Before editing governed symbols, require file-detail plus impact and report HIGH/CRITICAL as scope warnings, not bans.
- Before implementation commit, require detect-changes.
- Do not use `anvien --help` as task evidence; it only establishes command surface.

Why this habit is necessary: exact owner and complete call-path work can be wrong when based only on grep snippets or a stale graph. The campaign was explicitly corrected when Main told attribution workers not to use graph evidence and then substituted raw source reads. Fresh Anvien evidence connects symbols to callers, flows, modules, and blast radius, which is essential in the high-impact resolution pipeline. However, Anvien itself must not become a fixed ritual: choose the command matching the task, refresh once before graph work, and do not rerun graph analysis for a report-only conclusion that releases no source direction unless plan authoring rules specifically require the refresh.

### 12. Preserve measurement truth

- Elapsed wall-clock remains controlling. CPU/profile data is causal/supporting evidence only.
- Keep Cheapapp and Restaurant Manager independent; never average/combine.
- Small timing variation is an objective runtime fact, not automatically an error or an intolerable regression.
- A003 is the accepted Owner-specific KEEP and current baseline. A004/A005 remain rejected candidate evidence and rolled-back code.
- Do not promote an attribution measurement into a candidate, attempt, streak event, or disposition.

Why this habit is necessary: local metrics can improve while the product becomes slower. A004 and A005 both reduced D001 and parent time on both targets, yet process wall increased materially, so both were correctly `NO_KEEP`. A003 showed the inverse kind of real-world variation: Cheapapp D001/parent rose by about one second while process wall improved strongly on both targets, and Owner made an A003-specific KEEP after Supervisor correctness PASS. These examples show why timing is not a perfectly fixed speedometer reading and why every measured fact must remain objective. Main must neither erase small variation nor invent a universal tolerance from one Owner disposition.

### 13. Handle transport failures without creating functional duplication

- A task with `items=[]`, no ACK/tool/write is not a functional owner.
- Stop/revoke the transport-failed assignment and create one fresh visible replacement with the full packet.
- Never let two functional owners edit the same scope. Do not keep retrying dead tasks or create an audit of task delivery.

Why this habit is necessary: Codex task delivery occasionally accepted a turn while its transcript remained `items=[]`; those tasks showed no ACK, tool, artifact, or write for minutes. Treating such a shell as an active functional owner caused passive waiting, but repeatedly resending the same assignment caused more empty turns. The efficient recovery is evidence-based: confirm `items=[]` and zero work, send one explicit STOP/revocation, then create one fresh visible replacement. The replaced task must be named in the new prompt so late delivery cannot create two legitimate owners. This exact pattern was used for the active A006 Planner: old task `01a03f86...` owns nothing; fresh task `01a04153...` is the sole Planner.

## Absolute Boundaries At Transfer

- Do not reopen or rerun A001-A006 passed gates.
- Do not alter accepted A003 bytes or numbers.
- Do not call D001 `SYSTEM_CHARACTERISTIC` or a third no-KEEP attempt.
- Do not open D002 before the active Planner transition is committed.
- Do not infer D002 cause before the separate attribution report.
- Do not open D002 Architect before all three attribution facts are proven.
- Do not open Planner/Coder/measurement/P3/Child 07 early.
- Do not ask Owner to select the workflow already specified here.
- Only explicit PAUSE/STOP halts work.

## Successor First Unfinished Gate

1. ACK and complete raw re-anchor.
2. Bounded-monitor Planner `01a04153-b680-7e63-ae60-7108fe00278c` to its exact verdict.
3. Verify once, commit the exact six-path docs/report transition, and restore clean repo state.
4. Immediately dispatch the prepared visible D002 attribution packet.
5. Monitor it continuously; on complete cause/owner/call-path proof, sync ledgers once and open fresh D002 Architect.

No additional Owner confirmation is required for these steps. Only a new explicit Owner PAUSE/STOP/scope command changes this direction.
