# Main Forced Rotation Handoff — Child 06A A003 Residual Attribution

## Metadata

- Written: `2026-08-26 13:43:08 +07:00`
- Workspace: `E:\Anvien`
- Outgoing Main: `01a03cc6-952f-77e1-8c23-9d433bdb1aef`
- Successor Main: `01a03cd1-6dd7-7fc1-b744-ab73e9cc61e8` (`UNDERSTOOD` ACK received; active)
- Independent Governance Rule Guard fact: task `01a03c6d-d11d-7ae2-ab74-2fe91a757681` exists independently; Main must not command, retarget, request ACK from, wait for, duplicate, or functionally use it
- Rotation reason: latest explicit Owner STOP-and-handoff command withdrew this Main's operating authority after it improperly read/applied the Guard-only skill and compromised the Main/Guard independence boundary
- Safe-boundary time: `2026-08-26 13:43:08 +07:00`
- Current slice: `P2-A / A003 / B1-P1A-OP001 resolution / B2-P2A-A001-D001 resolve_calls`
- Current state: `A002 KEEP / A003 CURRENT_BASIS_RECORDED / RESIDUAL_CAUSE_PENDING / ARCHITECT_LOCKED`
- Git HEAD: `cc420e7ad719d90dc4b2d9991be0249e8d648daa`
- Staged path count: `0`

## Owner Stop And Governance Correction

The Owner ordered this Main to stop all campaign work immediately and transfer authority to a new visible Main. This Main must perform only the mandatory handoff, wait for successor ACK, record the successor ID, transfer authority, and terminate.

The exact violation is recorded truthfully:

- `orchestration §12` states that Main reads a `SKILL.md` only when Main directly uses that skill.
- `governance-rule-guard/SKILL.md` belongs exclusively to the independent continuous Guard lane and is not a Main capability.
- This Main incorrectly followed the prior handoff's invalid startup instruction, read that Guard skill, and stated that it was applying it.
- This was a pre-existing Main/Guard boundary violation, not a new Owner-created rule or a supersession of valid authority.
- The successor MUST NOT read or apply `governance-rule-guard/SKILL.md`. Future Main handoffs must omit that requirement and pass only the factual Guard task identity/non-interference boundary stated in Metadata.

## Mandatory Successor ACK And Raw Main-Owned Seal

The successor's first response MUST say `UNDERSTOOD` or `NOT UNDERSTOOD`, then briefly state the goal, current slice, boundary, and first action. If `NOT UNDERSTOOD`, it must stop before tools.

Before campaign action, the successor reads fully through EOF, in rule-before-skill order, only the authorities Main directly owns for the immediate work:

1. `E:\Anvien\AGENTS.md`
2. `E:\Anvien\.agents\skills\working-rules\SKILL.md`
3. `E:\Anvien\.agents\skills\orchestration\SKILL.md`
4. `E:\Anvien\.agents\skills\planner\SKILL.md` because the first campaign action is the bounded actual-status synchronization in the active plan

The successor MUST NOT read `E:\Anvien\.agents\skills\governance-rule-guard\SKILL.md`. It reads any later skill only when Main itself directly uses that skill for the current action. Long functional or acceptance roles remain separate visible lanes under orchestration rules.

You MUST apply 100% of the raw, unmodified `AGENTS.md`, `working-rules/SKILL.md`, and `orchestration/SKILL.md` at all times. You MUST NOT use this prompt, this handoff, any summary, memory, or compacted context as a substitute for those three iron rule files. This exact standard must be passed verbatim to every later Main successor.

After the raw Main-owned seal, read fully through EOF:

- this handoff;
- `E:\Anvien\reports\Investigation\rp_main_260826_133151_by_gpt-5_child06a_a003_rotation_handoff.md`;
- `E:\Anvien\docs\plans\2026-07-26-anvien-graph-accuracy-multi-plan\2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy\plan-rules.md`;
- all four standard Child 06A ledgers in that directory;
- `E:\Anvien\reports\system-architect\rp_system-architect_260825_184751_by_gpt-5_child06a_a002_run_scoped_diagnostic_appender.md`;
- `E:\Anvien\reports\coder\rp_coder_260825_193200_by_gpt-5_child06a_a002_paused_handoff.md`;
- `E:\Anvien\reports\Supervisor\rp_supervisor_260826_130617_by_gpt-5_child06a_a002_run_scoped_diagnostic_appender.md`.

Re-anchoring restores context only. Do not rerun or reopen A001, A002 architecture/Planner/source/build/test, overlay verification, reusable-script validation, two target measurements, or A002 Supervisor while their source/runtime/boundary remains valid.

## Exact Read Boundary At Forced Stop

This outgoing Main completed the following reads before the Owner STOP:

- full `AGENTS.md`;
- full `working-rules/SKILL.md`;
- full `orchestration/SKILL.md`;
- full `planner/SKILL.md`;
- full `supervisor/SKILL.md`;
- full prior rotation handoff `rp_main_260826_133151_by_gpt-5_child06a_a003_rotation_handoff.md`;
- full `plan-rules.md`;
- full standard `plan.md`;
- full standard `evidence.md`;
- full standard `benchmark.md`;
- only lines `1..225` of standard `actual-status.md`.

The outgoing Main had not yet completed `actual-status.md` or read the three required A002 reports when the Owner STOP arrived. The successor must complete those required reads from source rather than relying on this handoff. Do not turn that completion into a new audit or rerun a passed gate.

This outgoing Main performed no campaign mutation, ledger synchronization, `git diff --check`, residual attribution, graph command, Architect opening, source edit, detect, stage, commit, P3, or Child 07 action. The only new file is this Owner-required handoff report.

## Verified Completed Campaign State

- P0-A and P1-A are complete. P2-A is the sole open implementation slice.
- A001 is `KEEP`, committed at `17a1f3af37dcb61f9d389345822b6470a8f772cc`; its later one-file Coder-report commit is `18b5063d236f9f2567fea90e48eca8f1501bd1eb`.
- A002 exact accepted production delta is `+56/-4` across `internal/graphhealth/diagnostics.go`, `internal/resolution/emit.go`, and `internal/resolution/outcome.go`, plus new `internal/graphhealth/diagnostics_test.go`.
- A002 canonical full build passed before tests; five focused checks and `internal/graphhealth` passed. `internal/resolution` truthfully fails only on the identical current-HEAD-overlay preserve-only `TestProofBasedCallAccessGoldenCorpus`; never call that package PASS and do not edit its owner.
- Reusable frozen A00x script/packet evidence is complete under `E2-P2A-A002SCRIPTBUILD1`.
- Valid Cheapapp measurement is `E2-P2A-A002CHEAPAPP1`.
- Valid Restaurant Manager measurement is `E2-P2A-A002RESTAURANT1`; replacement task was `01a03c72-5f6d-7fd0-948a-88b06a1ebb65`.
- A002 Supervisor task `01a03c97-02f9-7560-ae31-4a333c31798e` returned exact verdict `SUPERVISOR_A002_PASS`.
- Supervisor report SHA-256 is `11BEA6715CD8626F80746B158FE755347E6924EAC560B2C5171525FBDE58A78B`.
- Main disposition is `A002 KEEP` under `E2-P2A-A002DECISION1`.

Separate accepted A002 baselines:

| Target | D001 `resolve_calls` | Parent `resolution` | Analyzer | Process | Denominator |
|---|---:|---:|---:|---:|---|
| `E:\cheapapp.org` | `3.090914200 s` | `19.040468000 s` | `100.843249000 s` | `136.729876000 s` | `calls=27890; files=887` |
| `E:\Restaurant_manager` | `9.909636600 s` | `21.242055400 s` | `109.339859600 s` | `145.066210900 s` | `calls=86030; files=1234` |

Never average or combine the targets. D001 remains active and unchecked with streak `0`; parent remains unchecked; D002-D017 remain queued and unopened.

## Commit Boundary Authority

Do not run detect, stage, or commit after A002.

Binding `plan-rules.md` states:

- P2-A is the only implementation slice and remains open across attempts;
- per-attempt commits are prohibited;
- P3-C owns the sole final detect and implementation commit after P2-A exhaustion and final review.

A002 `KEEP` therefore continues inside open P2-A. Accepted A002 source/test bytes remain uncommitted and staged set remains empty.

## First Unfinished Gate — One-Time Bounded Actual-Status Sync

`plan.md`, `evidence.md`, and `benchmark.md` already record `A002 KEEP / A003 CURRENT_BASIS_RECORDED / RESIDUAL_CAUSE_PENDING / ARCHITECT_LOCKED` substantially correctly.

Use Planner skill once to correct only these stale current-state statements in `actual-status.md`; preserve historical R0-R50 rows exactly and do not create a documentation review gate:

1. Relationship/Impact row still says `BOTH_TARGET_MEASUREMENTS_RECORDED / SUPERVISOR_PENDING`.
2. Current Control Snapshot follow-up paragraph still says Main owns corrected Cheapapp and no disposition exists.
3. P1 pointer follow-up paragraph still says Restaurant remains pending.
4. Current Status Matrix rows for current child/owner, P2-A attempts, A002 gate, and final speedup still describe pre-Supervisor/pre-KEEP state.
5. Detailed Findings' required-state sequence still ends at `fresh per-attempt Supervisor pending`, and the following rejection-loop section says A002 measurement/Supervisor remain pending.
6. The final `Decision note` has two paragraphs; delete the obsolete first paragraph and keep the A002-KEEP/A003 paragraph.
7. Add/check exact implementation-gate statements for `E2-P2A-A002REVIEW1`, `E2-P2A-A002DECISION1`, and `E2-P2A-A003CURRENT1` if absent.

After this single bounded synchronization, run `git diff --check` on the exact four ledgers plus the A002 Restaurant and Supervisor reports. Do not audit history and do not create a report, Supervisor loop, or documentation-review gate for the sync.

## Next Functional Gate After Sync

Perform bounded read-only attribution from the two accepted A002 CPU profiles and current source to prove all three before Architect:

1. current retained D001 residual cause;
2. exact source owner;
3. complete call path from analyze/resolution entry through `resolveCall` to the residual owner.

Profiles remain separate cumulative CPU samples. They may support cause but cannot be added, averaged, ranked as elapsed time, or used to predict speedup. Do not reuse A002 architecture.

Only when cause, owner, and call path are complete may Main open one fresh visible A003 Architect with the exact active parent, full 17-child list, separate accepted A002 baselines, denominators, source/profile evidence, preserved invariants, explicit non-goals, deliverable, stop/acceptance conditions, and next owner. No A003 Coder, source edit, detect, stage, commit, P3, or Child 07 action is permitted early.

## Workspace And Lane Boundary

- Worktree is intentionally dirty with Owner/campaign/protected work. Preserve it; no broad reset, checkout, stash, cleanup, or casual staging.
- A002 accepted source/test bytes and reusable script remain uncommitted inside P2-A.
- `CONTRIBUTING.md`, launcher/package-runtime changes, Child 06/07 ledgers, Vite timestamp, and unrelated dirty paths remain outside the A003 attribution boundary.
- Functional measurement lanes: none active.
- A002 Supervisor: complete/idle after PASS.
- A003 Architect: not opened and locked pending residual attribution.
- A003 Coder: not opened.
- Independent Guard fact: task `01a03c6d-d11d-7ae2-ab74-2fe91a757681` exists; do not command, retarget, request ACK from, wait for, duplicate, or functionally use it.

## Owner Communication And Lane Design

- Communicate with the Owner in Vietnamese.
- Questions, reminders, status requests, and corrections are not a pause. Only explicit `PAUSE` or `STOP` halts work.
- Classify work before delegation. For a simple execution lane, Main supplies the exact script, invocation, inputs, output path, report path/template, minimum validation, stop condition, and next owner; the lane must not discover tools or design a workflow.
- A complex lane may reason only inside a complete explicit boundary.
- Supervisor accepts the exact candidate code and related measurement/correctness/equivalence boundary; it does not audit the whole codebase or documentation.

## Official Transfer Procedure

1. Open one new visible Main in the existing saved project `E:\Anvien`, local shared checkout directly, no worktree and no fork.
2. Wait for the successor's mandatory `UNDERSTOOD` ACK.
3. Replace `PENDING_VISIBLE_SUCCESSOR_ACK` in this report with the exact successor task ID.
4. Send the successor the final report identity if needed, announce authority transfer, and terminate this outgoing Main immediately.
5. The successor starts its own 120-minute Main-only clock at activation, initializes another visible successor 15 minutes before its deadline, and hands off exactly at the deadline.
