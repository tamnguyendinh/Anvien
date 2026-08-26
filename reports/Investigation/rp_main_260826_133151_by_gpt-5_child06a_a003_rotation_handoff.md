# Main Rotation Handoff — Child 06A A003 Residual Attribution

## Metadata

- Written: `2026-08-26 13:31:51 +07:00`
- Workspace: `E:\Anvien`
- Outgoing Main: `01a03c31-2456-79f2-bab6-e71d42337052`
- Successor Main: `01a03cc6-952f-77e1-8c23-9d433bdb1aef`
- Independent Governance Rule Guard: `01a03c6d-d11d-7ae2-ab74-2fe91a757681`
- Rotation reason: the outgoing Main's mandatory 120-minute Main-only deadline (`2026-08-26 12:50:49 +07:00`) is overdue; this handoff occurs immediately after the mandatory raw startup seal.
- Current slice: `P2-A / A003 / B1-P1A-OP001 resolution / B2-P2A-A001-D001 resolve_calls`
- Current state: `A002 KEEP / A003 CURRENT_BASIS_RECORDED / RESIDUAL_CAUSE_PENDING / ARCHITECT_LOCKED`
- Git HEAD: `cc420e7ad719d90dc4b2d9991be0249e8d648daa`
- Staged path count: `0`

## Mandatory Full-Raw Successor Seal

The successor's first response MUST say `UNDERSTOOD` or `NOT UNDERSTOOD`, then briefly state goal, current slice, boundary, and first action. If `NOT UNDERSTOOD`, it must stop before tools.

Before any campaign action, read fully through EOF in this exact rule-before-skill order:

1. `E:\Anvien\AGENTS.md`
2. `E:\Anvien\.agents\skills\working-rules\SKILL.md`
3. `E:\Anvien\.agents\skills\orchestration\SKILL.md`
4. `E:\Anvien\.agents\skills\governance-rule-guard\SKILL.md`
5. `E:\Anvien\.agents\skills\planner\SKILL.md`
6. `E:\Anvien\.agents\skills\supervisor\SKILL.md`

You MUST apply 100% of the raw, unmodified `AGENTS.md`, `working-rules/SKILL.md`, and `orchestration/SKILL.md` at all times. You MUST NOT use this prompt, the handoff, any summary, memory, or compacted context as a substitute for those three iron rule files. This exact standard must be passed verbatim to every later Main successor.

After that seal, read fully through EOF:

- this handoff;
- `E:\Anvien\docs\plans\2026-07-26-anvien-graph-accuracy-multi-plan\2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy\plan-rules.md`;
- all four standard Child 06A ledgers in that directory;
- `E:\Anvien\reports\system-architect\rp_system-architect_260825_184751_by_gpt-5_child06a_a002_run_scoped_diagnostic_appender.md`;
- `E:\Anvien\reports\coder\rp_coder_260825_193200_by_gpt-5_child06a_a002_paused_handoff.md`;
- `E:\Anvien\reports\Supervisor\rp_supervisor_260826_130617_by_gpt-5_child06a_a002_run_scoped_diagnostic_appender.md`.

Re-anchoring restores context only. Do not rerun or reopen A001, A002 architecture/Planner/source/build/test, overlay verification, reusable-script validation, two target measurements, or A002 Supervisor while their source/runtime/boundary remains valid.

## Latest Owner Authority And Lane-Design Correction

- Communicate with the Owner in Vietnamese.
- Owner explicitly ordered continuation; questions, reminders, status requests, and corrections are not a pause. Only explicit `PAUSE` or `STOP` halts work.
- Classify work before delegation. A simple execution lane must not infer or search for tools. Main supplies the exact script, invocation, inputs, output path, report path/template, validation minimum, stop condition, and next owner. The lane then only runs the exact command, captures the result, and writes the bounded report.
- A complex lane may reason only inside a complete explicit boundary.
- Supervisor accepts the exact candidate code and related measurement/correctness/equivalence boundary; it does not audit the whole codebase or documentation.
- Do not command, retarget, request ACK from, wait for, or functionally use the independent Guard. Do not create a duplicate Guard.

## Verified Completed State

- P0-A and P1-A are complete. P2-A is the sole open implementation slice.
- A001 is `KEEP`, committed at `17a1f3af37dcb61f9d389345822b6470a8f772cc`; its later one-file Coder-report commit is `18b5063d236f9f2567fea90e48eca8f1501bd1eb`.
- A002 exact production delta is `+56/-4` across `internal/graphhealth/diagnostics.go`, `internal/resolution/emit.go`, and `internal/resolution/outcome.go`, plus new `internal/graphhealth/diagnostics_test.go`.
- A002 canonical full build passed before tests; five focused checks and `internal/graphhealth` passed. `internal/resolution` remains truthfully FAIL only on the identical current-HEAD-overlay preserve-only `TestProofBasedCallAccessGoldenCorpus`; do not call that package PASS and do not edit its owner.
- Reusable frozen A00x script/packet evidence is complete under `E2-P2A-A002SCRIPTBUILD1`.
- Valid Cheapapp measurement is `E2-P2A-A002CHEAPAPP1`.
- Valid Restaurant Manager measurement is `E2-P2A-A002RESTAURANT1`; replacement task was `01a03c72-5f6d-7fd0-948a-88b06a1ebb65`.
- A002 Supervisor task `01a03c97-02f9-7560-ae31-4a333c31798e` returned exact verdict `SUPERVISOR_A002_PASS`.
- Supervisor report SHA-256: `11BEA6715CD8626F80746B158FE755347E6924EAC560B2C5171525FBDE58A78B`.
- Main disposition is `A002 KEEP` under `E2-P2A-A002DECISION1`.

Separate accepted A002 baselines:

| Target | D001 `resolve_calls` | Parent `resolution` | Analyzer | Process | Denominator |
|---|---:|---:|---:|---:|---|
| `E:\cheapapp.org` | `3.090914200 s` | `19.040468000 s` | `100.843249000 s` | `136.729876000 s` | `calls=27890; files=887` |
| `E:\Restaurant_manager` | `9.909636600 s` | `21.242055400 s` | `109.339859600 s` | `145.066210900 s` | `calls=86030; files=1234` |

Never average or combine the targets. D001 remains active and unchecked with streak `0`; parent remains unchecked; D002-D017 remain queued and unopened.

## Commit-Boundary Authority

Do not run detect, stage, or commit after A002.

Binding `plan-rules.md` states that:

- `P2-A` is the only implementation slice and stays open across attempts;
- per-attempt commits are prohibited;
- P3-C owns the sole final detect and implementation commit after P2-A exhaustion and final review.

The general orchestration/working-rule commit language therefore does not make A002 a new slice boundary. A002 `KEEP` continues inside open P2-A. The Guard warning that requested an A002 commit conflicts with the exact active-plan authority and must not be executed.

## Current Ledger State And One-Time Sync Still Needed

`plan.md`, `evidence.md`, and `benchmark.md` already record `A002 KEEP / A003 CURRENT_BASIS_RECORDED / RESIDUAL_CAUSE_PENDING / ARCHITECT_LOCKED` substantially correctly.

`actual-status.md` still contains a small number of stale current-state statements mixed with correct history. Use Planner skill to correct them once, without auditing historical R0-R50 rows or generic state-model text:

1. Relationship/Impact row still says `BOTH_TARGET_MEASUREMENTS_RECORDED / SUPERVISOR_PENDING`.
2. Current Control Snapshot follow-up paragraph still says Main owns corrected Cheapapp and no disposition exists.
3. P1 pointer follow-up paragraph still says Restaurant remains pending.
4. Current Status Matrix rows for current child/owner, P2-A attempts, A002 gate, and final speedup still describe pre-Supervisor/pre-KEEP state.
5. Detailed Findings' required-state sequence still ends at `fresh per-attempt Supervisor pending`, and the following rejection-loop section says A002 measurement/Supervisor remain pending.
6. The final `Decision note` has two paragraphs; delete the obsolete first paragraph and keep the A002-KEEP/A003 paragraph.
7. Add/check exact implementation-gate statements for `E2-P2A-A002REVIEW1`, `E2-P2A-A002DECISION1`, and `E2-P2A-A003CURRENT1` if absent.

Preserve historical R0-R50 rows exactly. After this bounded one-time sync, run `git diff --check` on the four ledgers plus the A002 Restaurant and Supervisor reports. Do not create a report or extra review gate for this documentation sync.

## Exact Next Functional Gate

After the one-time ledger sync, perform bounded read-only attribution from the two accepted A002 CPU profiles and current source to prove all three before Architect:

1. current retained D001 residual cause;
2. exact source owner;
3. complete call path from analyze/resolution entry through `resolveCall` to the residual owner.

Profiles remain separate cumulative CPU samples: they may support cause but cannot be added, averaged, ranked as elapsed time, or used to predict speedup. Do not reuse A002 architecture.

Only if cause/owner/call path are complete, open one fresh visible A003 Architect with the exact active parent, full 17-child list, separate accepted A002 baselines, denominator, source/profile evidence, preserved invariants, explicit non-goals, deliverable, stop/acceptance conditions, and next owner. No A003 Coder, source edit, detect, stage, commit, P3, or Child 07 action exists or is allowed before the fresh Architect decision and Planner refresh.

## Workspace Boundary

- HEAD is `cc420e7ad719d90dc4b2d9991be0249e8d648daa`.
- Staged set is empty.
- The worktree is intentionally dirty with Owner/campaign/protected work. Preserve it; no broad reset, checkout, stash, cleanup, or casual staging.
- A002 accepted source/test bytes and the reusable script remain uncommitted inside P2-A.
- `CONTRIBUTING.md`, launcher/package-runtime changes, Child 06/07 ledgers, Vite timestamp, and unrelated dirty paths are outside the current A003 attribution boundary.

## Active Lanes At Transfer

- Functional measurement lanes: none active; both target packets are complete.
- A002 Supervisor: complete/idle after PASS.
- A003 Architect: not opened and locked pending residual attribution.
- A003 Coder: not opened.
- Independent Guard: continuous task `01a03c6d-d11d-7ae2-ab74-2fe91a757681`; independent and not under Main control.

## Official Transfer

When the successor task ID is written into Metadata and the successor becomes active, it is the sole visible Main Orchestration authority. Its 120-minute Main-only rotation clock starts at activation. The outgoing Main must stop immediately and perform no further campaign action.
