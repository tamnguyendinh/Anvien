# Visible Main Orchestration Rotation Handoff

## Authority transfer record

- Campaign: Anvien Graph Accuracy / Analyze Performance.
- Outgoing Visible Main: `01a03290-28f5-7870-a643-871af4aa24d4`.
- Clean prepared successor: `01a032e4-e726-7b12-919a-ef6e617414f6`.
- Raw outgoing official authority UUIDv7: `01a03297-41ff-7260-9e05-c4ea9accee6d`.
- Outgoing authority start: `2026-08-24T14:06:08.767+07:00`.
- Outgoing warm target: `2026-08-24T14:51:08.767+07:00`.
- Outgoing hard deadline: `2026-08-24T15:06:08.767+07:00`.
- App reset interrupted the active Main/Worker/Guard turns before scheduled warm creation. The overdue state is historical fact, not justification or normalization.
- Clean successor was created by `create_thread` against exact saved project ID `local-dc76d679b5f0e3c8ddbb8f9218f46f6`, path `E:\Anvien`, environment `local`, no worktree.
- Clean successor creation UUIDv7: `01a032e4-e726-7b12-919a-ef6e617414f6`.
- Clean successor creation time: `2026-08-24T15:30:57.318+07:00`.
- Exact clean-successor warm lateness: `2388.551s — HIGH`. Preserve exactly.
- Exact clean-successor hard-deadline lateness at creation: `1488.551s — HIGH`. Preserve exactly.
- Clean successor returned `UNDERSTOOD` and `WAITING_FOR_OFFICIAL_TRANSFER`, stated the campaign goal, P6-E sole-open-slice, Main-only boundary, bans, Worker/Guard preservation, runtime-identity blocker, and invalid-fork exclusion. It used no tool and remains idle.
- The raw UUIDv7 of the follow-up headed exactly `OFFICIAL AUTHORITY TRANSFER` is transport source of truth for incoming authority time. Incoming Main computes warm target `+45 minutes` and hard deadline `+60 minutes` from that UUIDv7.
- Outgoing Main terminates absolutely immediately after official dispatch.

## Invalid fork exclusion

- Three attempts to fork the exact current Main failed with thread-store error `expected ordinal 1669, got 1668`.
- Outgoing Main then incorrectly used a same-directory fork from completed prior Main as fallback.
- Invalid fork task: `01a032d9-44a3-7fe2-b97b-8a8fe89fe8f3`.
- Invalid fork creation time: `2026-08-24T15:18:14.819+07:00`.
- Its warm lateness was `1626.052s — HIGH`; hard-deadline lateness at creation was `726.052s — HIGH`.
- Governance correctly classified the fork as `HIGH — VERIFIED VIOLATION`: a fork inheriting completed history is not the clean new visible successor required by current Owner authority.
- The invalid fork received only a PRE-TRANSFER waiting prompt. It never received official transfer, has no authority, must never receive authority, and is excluded from all future rotation decisions.
- Never use `fork_thread` as successor fallback again. A clean task must be created by `create_thread`; if blocked, report the exact blocker.

## Mandatory iron startup and anti-summary rule

Before every orchestration action, incoming Main must read FULL RAW sequentially through EOF:

1. `E:\Anvien\AGENTS.md`.
2. `E:\Anvien\.agents\skills\working-rules\SKILL.md`.
3. `E:\Anvien\.agents\skills\orchestration\SKILL.md`.
4. `E:\Anvien\.agents\skills\governance-rule-guard\SKILL.md`.
5. This entire sealed handoff and verify the detached seal supplied in the official transfer message.

Before plan, ledger, lane-control, corpus, baseline, implementation authorization, acceptance, or Worker-resume action, continue FULL RAW through EOF:

1. `E:\Anvien\.agents\skills\planner\SKILL.md`.
2. `E:\Anvien\docs\plans\2026-07-26-anvien-graph-accuracy-multi-plan\2026-07-26-anvien-graph-accuracy-roadmap.md`.
3. All four Child 06 ledgers under `E:\Anvien\docs\plans\2026-07-26-anvien-graph-accuracy-multi-plan\2026-07-28-06-ambient-external-resolution-and-diagnostics\`.
4. Prior sealed handoff `E:\Anvien\reports\Investigation\rp_main_260824_140322_orchestration_rotation_handoff.md`.
5. Current P6-E Worker checkpoint files named below.
6. The final independent runtime-identity Supervisor report if it has materialized.

Do not use summary, compacted context, prior memory, filenames, task titles, or the invalid self-PASS instead of raw source. After auto-compaction, perform orchestration section 7 FULL RAW re-anchor and continue from the first uncompleted gate; do not reopen a passed gate without proved invalidation.

## Current campaign state

- Sole workspace: `E:\Anvien`, saved project local, no worktree.
- Branch: `master`.
- Current HEAD at handoff snapshot: `1a9b5310bc3204c610e9e6c673afc214cd049991`.
- Parent: `d02e4eb76ee6e42848ea0ffb627449ac7d9df092`.
- Tree: `f124bc5811ae26b2751c1b09fc8da673b9f4ab1c`.
- P6-D remains CLOSED at accepted implementation commit `81163e39718b94a509e41114cada224e8f269e36`; do not reopen accepted gates absent proved invalidation.
- P6-E is the sole open Child 06 slice and exactly one commit boundary: `U1 -> U2 -> rebaseline/ownership -> U3 -> conditional U4 -> Pareto -> final equivalence`.
- U1-U4 are internal ordered units, not independent slices, gates, or commits. Pn-A and later gates remain locked.
- No valid current full-row-equal corpus seal, graph/storage authority, valid A/A invocation, comparable baseline, U1, stage, commit, push, or target access exists. A/A is exactly `0/20`.
- Approximate whole-slice progress remains about `15%`; this is status estimate only and never acceptance evidence.

## Runtime-identity role-separation blocker

- Prior Main directly read the Supervisor skill, directly authored `E:\Anvien\reports\Supervisor\rp_supervisor_260824_144039_by_gpt-5_p6e_runtime_identity_supersession.md`, issued `Verdict identity-only: PASS`, and used that self-PASS to supersede runtime identity and resume Worker.
- Governance classified this `HIGH — VERIFIED VIOLATION`. Main is not Supervisor; only an independent visible Supervisor may issue an acceptance verdict.
- The self-authored report and self-PASS are invalid as Supervisor authority or acceptance evidence. Preserve their bytes as historical provenance/allegation only; do not delete, rewrite, normalize, or cite them as acceptance.
- Worker is already at a safe boundary because the app reset interrupted its raw-ledger re-anchor. No downstream harness/corpus/graph/A-A/U1 action followed the reset.
- A clean independent visible Supervisor task is active: `01a032e6-d308-7751-b82c-92a4652dc93b`, saved project `E:\Anvien`, local, no worktree.
- Supervisor scope is identity-only: independently verify current Git/runtime/vendor/native/sidecar/process/exclusive-open evidence and return PASS/REJECT/BLOCKED. It must not accept harness, corpus, graph, A/A, baseline, U1, or implementation.
- Supervisor was instructed to create exactly one new durable report under `E:\Anvien\reports\Supervisor\` and seal bytes/LF/CR/UTF-8/BOM/SHA-256/timestamps. At this handoff snapshot it is ACTIVE and has not returned a verdict.
- Incoming Main must monitor that exact Supervisor task. Do not resume Worker based on the old self-PASS. After the independent final arrives, read FULL RAW Supervisor skill and the new durable report, verify its seal/source/Git boundary, then relay only the accepted exact correction to the existing Worker. On REJECT/BLOCKED, preserve Worker stop and act only on the exact rejected invariant.

### Alleged runtime identity pending independent verdict

- Runtime path: `E:\Anvien\anvien\bin\anvien.exe`.
- Claimed public version: `1.2.8`.
- Claimed bytes: `73,750,016`.
- Claimed SHA-256: `D4EC8B58C41B9F0A95359CFC014DB07D11653F25736728549555F74D204ABD19`.
- Claimed last-write: `2026-08-24T14:22:30.7298427+07:00`.
- Claimed Go build identity: Go `1.26.3`, `ladybugdb`, `trimpath`, `vcs.revision=1a9b5310bc3204c610e9e6c673afc214cd049991`.
- Claimed vendor manifest SHA-256: `678A598E0D84B4903DB9684B64F20CB1F273A4785FC65CE9B1718DBC67C7ECDC`.
- Claimed native DLL SHA-256: `20CBD87840483A2053CFF3FC2DB23A86DD802B8915D86509D41A4B709624CDB7`.
- Worker blocker disposition: `E:\Anvien\reports\coder\p6e_analyze_performance_260824_093855_raw\preflight\runtime-identity-drift-260824_143600.json`.
- Every value above is allegation/pending evidence until the independent Supervisor verdict closes the identity gate.

## Binding Owner and Architect authority

- Never ask Owner to approve, choose, review, or comment on numeric materiality/resource rules or any other intermediate P6-E decision.
- After identity and harness/corpus gates are valid, Worker executes exactly `10` valid A/A pairs / `20` unprofiled invocations, no early stop, then stops before U1 and hands the raw summary plus numeric proposal to Main.
- Main opens one separate clean visible Architect lane. Architect alone owns and seals numeric materiality/resource rules.
- Main relays the sealed Architect decision to the existing Worker; only that seal opens U1.
- Owner receives only the final completed campaign result unless Owner independently issues `PAUSE`, `STOP`, or a scope override.

## Existing P6-E Worker

- Task: `01a0319a-62c9-7650-ba6a-ef59dfcf1973`.
- Saved project: `E:\Anvien`; local; no worktree.
- Durable root: `E:\Anvien\reports\coder\p6e_analyze_performance_260824_093855_raw`.
- Scratch root: `E:\Anvien\.tmp\p6e-worker-01a0319a`; disposable scratch only, never evidence/authority/blocker.
- App reset interrupted the active Worker turn. Current thread state is idle/interrupted at a safe boundary.
- Latest visible Worker message: roadmap and Child 06 plan were raw-read; three remaining ledgers were being read; no harness edit followed that compaction.
- Preserve the exact Worker. Do not duplicate, replace, pressure, archive, transfer, or open a second Worker.
- Worker owns harness/schema, fresh corpus/graph/impact, exactly 10 valid A/A pairs, raw numeric proposal, and ordered implementation only after independent runtime identity and Architect gates.
- Main is not Worker.

Read these Worker files FULL RAW before any Worker direction:

- `README.md`.
- `preflight/lifecycle.md`.
- `preflight/frozen_boundary.json`.
- `preflight/post-handoff-corpus-freeze.seal.json`.
- `preflight/post-handoff-graph-authority-corpus.seal.json`.
- `preflight/aa-baseline-protocol.md`.
- `costmap/living-cost-map-v1.md`.
- `graph/impact-summary.md`.

Historical `2,173`-row corpus, manifest `c81d33e3acdf21ca7689b351ad53230cf693b1203d9b72f87e9fcd20c821d853`, graph `2,173/765/0`, `122,696/169,347`, and runtime `73,749,504 / 64EECEBA...DEBB` are lineage only for A/A. Owner commits, Main handoffs, ledger changes, runtime drift, Supervisor outputs, and Worker artifacts require one new exact current corpus/identity chain.

## Required Worker continuation after independent identity PASS

1. Relay the exact independent Supervisor seal/correction to the existing Worker.
2. Worker resumes raw re-anchor at the first incomplete source; do not restart passed product gates.
3. Finish and seal one compiled fail-closed producer contract, including independent review of bounded hidden-subagent input and exact lifecycle/schema/cleanup tests.
4. Produce two consecutive full-row-equal corpus inventories excluding only the exact Worker durable root and including every current protected/Owner/Main/Supervisor row.
5. Run one fresh excluded graph/storage authority and current file-detail/impact rerun. Graph authority is not an A/A invocation.
6. Run exactly `10` valid A/A pairs / `20` unprofiled invocations, no early stop. Invalid runs remain durable and execution continues until exactly 10 valid pairs exist.
7. Worker stops before U1 and returns raw summary/numeric proposal to Main.
8. Main opens one separate clean visible Architect, relays its sealed decision to the same Worker, and only then opens U1.

## Existing visible lanes and governance Guard

- P6-D Coder `01a02f10-c175-7162-ac32-6df30a4d604a`: idle after accepted handoff; preserve.
- P6-D Supervisor `01a030d2-3ae2-7b33-a57a-c23503703b1f`: idle after PASS; preserve.
- P6-E Worker `01a0319a-62c9-7650-ba6a-ef59dfcf1973`: idle/interrupted after app reset; preserve and hold pending independent identity verdict.
- P6-E runtime-identity Supervisor `01a032e6-d308-7751-b82c-92a4652dc93b`: active independent visible review; monitor to verdict.
- Historical Guard `01a031fd-366f-79f2-a9e3-04c365403fd3` rotated to `01a032b4-00b2-7282-91f0-6d8161299940`; current governance messages arrive from `01a032b6-7eaf-7b32-94f1-c693be834928`. Preserve/reuse current Guard lineage, do not duplicate it, and do not direct it as a functional lane.
- Invalid fork `01a032d9-44a3-7fe2-b97b-8a8fe89fe8f3`: no authority, never transfer to it.

## Git/worktree snapshot

- Snapshot: `2026-08-24T15:35:50.1543017+07:00`.
- Branch/HEAD/parent/tree: `master` / `1a9b5310bc3204c610e9e6c673afc214cd049991` / `d02e4eb76ee6e42848ea0ffb627449ac7d9df092` / `f124bc5811ae26b2751c1b09fc8da673b9f4ab1c`.
- Index: empty, `0` staged paths.
- Status before this handoff: `221` rows = `184` exact Worker-root rows + `37` outside-Worker rows; `5` tracked dirty rows + `216` untracked rows.
- This handoff adds one protected outside-Worker row. Active independent Supervisor may later add exactly one additional protected report row.
- `git diff --check`: PASS.
- Preserve Owner deletion `CONTRIBUTING.md`, all four Child 06 ledger edits, Owner aicontext/governance commits/bytes, Microsoft PowerShell artifact, every historical Main handoff, every P6-D/P6-E report, every Worker artifact, vendor, third_party, and all unrelated/protected rows.
- Never normalize, stash, reset, checkout, broad-stage, broad-clean, or broad-commit this worktree.

## Verified governance findings

- Historical first post-transfer ACK omission remains `MEDIUM — VERIFIED VIOLATION`; corrective ACK did not reset passed gates.
- Historical unscoped `list_projects {}` call remains `HIGH — VERIFIED VIOLATION`; its output was discarded and never reused.
- Historical misapplication of orchestration section 2 to the Worker's bounded hidden subagent remains `HIGH — VERIFIED VIOLATION`; the Worker finding was fully withdrawn and must never be resurrected.
- Historical warm lateness `0.205s — HIGH — VERIFIED VIOLATION`, `44.162s — HIGH`, `25.844s — HIGH`, `0.240s — HIGH`, and `0.154s — HIGH` remain immutable history.
- Current clean-successor warm lateness is exactly `2388.551s — HIGH`; hard-deadline lateness at clean creation is `1488.551s — HIGH`.
- Invalid successor fork is `HIGH — VERIFIED VIOLATION`; it has no authority.
- Main self-authored Supervisor report/verdict is `HIGH — VERIFIED VIOLATION`; self-PASS is invalid and the independent Supervisor gate is open.
- A later old-Guard `RISK — HIGH` inference from null transcript markers was fully retracted as invalid and remains `NO EVIDENCE`; do not resurrect it.
- Owner later explicitly granted full authority to run the required project lookup. The one current `list_projects` call was therefore authorized, filtered only exact `E:\Anvien`, and yielded saved project ID `local-dc76d679b5f0e3c8ddbb8f9218f46f6d` for clean task creation.

## Immediate continuation for incoming Main

1. Complete mandatory FULL RAW startup and detached-seal verification.
2. Monitor exact independent Supervisor `01a032e6-d308-7751-b82c-92a4652dc93b`; do not resume Worker while verdict/report is absent.
3. On final verdict, read FULL RAW Supervisor skill and new durable report, verify seal/current Git/runtime boundary, and relay only the valid exact decision to existing Worker.
4. On PASS, resume existing Worker from interrupted raw re-anchor, then harness seal -> fresh current corpus -> fresh excluded graph/storage/current impacts -> exactly 10 A/A pairs.
5. On REJECT/BLOCKED, preserve Worker hold and send only the exact corrective invariant to the appropriate existing lane; do not self-accept.
6. After A/A, open one separate clean visible Architect; never ask Owner for intermediate decision.
7. Preserve one P6-E commit boundary and ordered U1-U4. Open independent P6-E Supervisor only after a complete no-commit candidate and durable Coder handoff.
8. After final Supervisor PASS: planner refresh, fresh current E-only detect, exact cleanup disposition, exact isolated staging, one P6-E commit under standing AGENTS Rule 12 authority, no push.
9. Initialize the next clean visible Main successor at the new exact warm target and transfer no later than the new hard deadline.

## Workspace bans and standing authority

- Sole workspace is `E:\Anvien`, saved project local, no worktree.
- No C-drive workspace/read/write, alternate worktree, target access, implicit network, install/package-script side effects, shared/global/protected/quarantined cache authority, reset/stash/checkout/broad cleanup, or push.
- Every `E:\Anvien\.tmp` path is disposable reproducible scratch and never evidence, review/handoff/recovery authority, source of truth, acceptance dependency, or blocker.
- Preserve vendor, third_party, durable reports, Owner work/deletion, current Owner commits, unrelated/protected rows, and all historical evidence.
- Owner `PAUSE` or `STOP` halts commands, edits, QA, commits, and lane control immediately.
- Commit after completed accepted P6-E is standing AGENTS Rule 12 authority; never ask Owner to grant it again.
- No push unless Owner explicitly requests push.

Outgoing Main terminates absolutely after official transfer dispatch.
