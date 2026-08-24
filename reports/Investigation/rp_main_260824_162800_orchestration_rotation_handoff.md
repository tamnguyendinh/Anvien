# Visible Main Orchestration Rotation Handoff

## Authority transfer record

- Campaign: Anvien Graph Accuracy / Analyze Performance.
- Outgoing Visible Main: `01a032e4-e726-7b12-919a-ef6e617414f6`.
- Clean prepared successor: `01a03316-a186-7fe0-aa6e-cb4406fea297`.
- Raw incoming official-authority UUIDv7 for this outgoing Main: `01a032ec-a59b-7491-9348-dfb7028af46b`.
- Outgoing authority start: `2026-08-24T15:39:24.827+07:00`.
- Outgoing warm target: `2026-08-24T16:24:24.827+07:00`.
- Outgoing hard transfer deadline: `2026-08-24T16:39:24.827+07:00`.
- The raw UUIDv7 of the follow-up headed exactly `OFFICIAL AUTHORITY TRANSFER` is the incoming Main's authority-time source of truth. Incoming Main computes its next warm target at `+45 minutes` and hard deadline at `+60 minutes` from that UUIDv7.
- Outgoing Main terminates absolutely immediately after official dispatch.

## Current rotation transport and verified violation

- The first scheduled `create_thread` transport used saved-project ID `local-dc76d679b5f0e3c8ddbb8f9218f46f6`, exactly as written in the outgoing official-transfer message. The call failed `Unknown projectId` at observed dispatch timestamp `2026-08-24T16:24:43.296+07:00`.
- The prior sealed handoff contained the valid exact saved-project ID with a trailing `d`: `local-dc76d679b5f0e3c8ddbb8f9218f46f6d`.
- Main immediately retried with that exact sealed ID, saved project local `E:\Anvien`, no worktree, and did not use project discovery or `fork_thread`.
- Clean successor task: `01a03316-a186-7fe0-aa6e-cb4406fea297`.
- Raw successor UUIDv7 creation time: `2026-08-24T16:25:16.294+07:00`; raw-UUID lateness from warm target is `51.467s`.
- Tool-observed successful dispatch timestamp: `2026-08-24T16:25:17.284+07:00`.
- Governance's binding observation is warm lateness exactly `52.457s — HIGH — VERIFIED VIOLATION`. Preserve exactly; do not normalize, downgrade, or justify it from the project-ID discrepancy.
- Successor returned `UNDERSTOOD` and `WAITING_FOR_OFFICIAL_TRANSFER`, stated the exact campaign goal, P6-E sole-open-slice, Main-only boundary, workspace/bans, Worker/Guard preservation, identity-only PASS boundary, invalid-fork exclusion, and absolute no-action rule before official transfer. It used no tool and is idle.
- Successor ACK turn: `01a03316-a395-7690-8ec4-9510ae8a5e84`.
- The latest Guard warning requires this handoff and official transfer before `2026-08-24T16:39:24.827+07:00`. No campaign lane action is authorized during that priority transition.

## Invalid transport exclusion

- Historical invalid fork `01a032d9-44a3-7fe2-b97b-8a8fe89fe8f3` never received official authority and remains permanently excluded.
- Never use `fork_thread` for Main rotation. Use a clean `create_thread` on the exact saved project local `E:\Anvien`, no worktree.
- Do not reuse the failed saved-project ID without the trailing `d`; current valid ID is `local-dc76d679b5f0e3c8ddbb8f9218f46f6d`.

## Mandatory iron startup and anti-summary rule

Before every orchestration action, incoming Main must read FULL RAW sequentially through EOF:

1. `E:\Anvien\AGENTS.md`.
2. `E:\Anvien\.agents\skills\working-rules\SKILL.md`.
3. `E:\Anvien\.agents\skills\orchestration\SKILL.md`.
4. `E:\Anvien\.agents\skills\governance-rule-guard\SKILL.md`.
5. This entire sealed handoff and verify the detached seal supplied in the official transfer message.

Before plan, ledger, lane-control, corpus, baseline, implementation authorization, acceptance, or Worker direction, continue FULL RAW through EOF:

1. `E:\Anvien\.agents\skills\planner\SKILL.md`.
2. `E:\Anvien\docs\plans\2026-07-26-anvien-graph-accuracy-multi-plan\2026-07-26-anvien-graph-accuracy-roadmap.md`.
3. All four Child 06 ledgers under `E:\Anvien\docs\plans\2026-07-26-anvien-graph-accuracy-multi-plan\2026-07-28-06-ambient-external-resolution-and-diagnostics\`.
4. Prior sealed handoffs `E:\Anvien\reports\Investigation\rp_main_260824_153550_orchestration_rotation_handoff.md` and `E:\Anvien\reports\Investigation\rp_main_260824_140322_orchestration_rotation_handoff.md`.
5. The eight Worker checkpoint files named below and the latest Worker provenance/checkpoint.
6. Independent runtime-identity Supervisor report `E:\Anvien\reports\Supervisor\rp_supervisor_260824_154625_by_gpt-5_p6e_runtime_identity_independent_review.md`.

Do not use summary, compacted context, prior memory, filenames, task titles, or the invalid historical self-PASS instead of raw source. After auto-compaction, perform orchestration section 7 FULL RAW re-anchor and continue from the first uncompleted gate; do not reopen a passed gate without proved invalidation.

## Current campaign state

- Sole workspace: `E:\Anvien`, saved project local, no worktree.
- Branch: `master`.
- Current HEAD: `1a9b5310bc3204c610e9e6c673afc214cd049991`.
- Parent: `d02e4eb76ee6e42848ea0ffb627449ac7d9df092`.
- Tree: `f124bc5811ae26b2751c1b09fc8da673b9f4ab1c`.
- P6-D remains CLOSED at accepted implementation commit `81163e39718b94a509e41114cada224e8f269e36`; do not reopen accepted gates absent proved invalidation.
- P6-E is the sole open Child 06 slice and exactly one commit boundary: `U1 -> U2 -> rebaseline/ownership -> U3 -> conditional U4 -> Pareto -> final equivalence`.
- U1-U4 are ordered internal units, not independent slices, gates, or commits. Pn-A and later gates remain locked.
- Approximate whole-slice progress is about `15%`; this is an Owner-facing estimate only, never acceptance evidence.
- A/A is exactly `0/20`. No accepted current harness seal, full-row-equal corpus seal, current graph/storage authority, comparable baseline, U1, stage, P6-E commit, push, or target access exists.
- The runtime-identity blocker is closed only at the identity continuation boundary. The first uncompleted technical gate is the compiled fail-closed producer/harness contract.

## Independent runtime-identity closure

- Independent visible Supervisor: `01a032e6-d308-7751-b82c-92a4652dc93b`.
- Durable report: `E:\Anvien\reports\Supervisor\rp_supervisor_260824_154625_by_gpt-5_p6e_runtime_identity_independent_review.md`.
- Verdict: `PASS`, identity-only.
- Identity: `13,586` bytes / `174` LF / `0` CR / strict UTF-8 without BOM.
- SHA-256: `467D9C75E2BA73FD70D8A1CC42A140400859AE751384790C464F8FE6B9A655A1`.
- Created/last-write: `2026-08-24T15:48:19.1841411+07:00`.
- Accepted current direct runtime: `E:\Anvien\anvien\bin\anvien.exe`, version `1.2.8`, `73,750,016` bytes, SHA-256 `D4EC8B58C41B9F0A95359CFC014DB07D11653F25736728549555F74D204ABD19`.
- Sidecar SHA-256: `C454EBDF4CD28AD61D0D9A903C0C761F2E135316BE2325E3F59C6A3A45B2C78D`.
- Vendor manifest SHA-256: `678A598E0D84B4903DB9684B64F20CB1F273A4785FC65CE9B1718DBC67C7ECDC`.
- Native DLL SHA-256: `20CBD87840483A2053CFF3FC2DB23A86DD802B8915D86509D41A4B709624CDB7`.
- Worker consumed this authority in `preflight\runtime-identity-independent-supersession-260824_155835.json`, written `2026-08-24T15:59:29.5413473+07:00`.
- The prior Main-authored `rp_supervisor_260824_144039_by_gpt-5_p6e_runtime_identity_supersession.md` and its self-PASS remain `HIGH — VERIFIED VIOLATION`, historical allegation/provenance only, and invalid as Supervisor authority.
- Identity PASS does not accept harness, corpus, graph/storage authority, A/A, comparable baseline, numeric rules, U1, implementation, target, stage, commit, or push.

## Existing P6-E Worker and current progress

- Task: `01a0319a-62c9-7650-ba6a-ef59dfcf1973`.
- Saved project: `E:\Anvien`; local; no worktree.
- Durable root: `E:\Anvien\reports\coder\p6e_analyze_performance_260824_093855_raw`.
- Scratch root: `E:\Anvien\.tmp\p6e-worker-01a0319a`; disposable scratch only, never evidence/authority/blocker.
- Preserve this exact Worker. Do not duplicate, replace, pressure, archive, transfer, or open a second Worker.
- Worker is actively repairing the fail-closed harness under the durable root. At snapshot `2026-08-24T16:27:31.3300762+07:00`, the root contained `185` files and current edits reached:
  - `scripts\Compare-P6EAAPair.ps1`, last-write `2026-08-24T16:26:39.9594446+07:00`;
  - `scripts\P6EHarness.Common.ps1`, last-write `2026-08-24T16:26:38.3571582+07:00`;
  - `scripts\Compare-P6ECorpusInventories.ps1`, last-write `2026-08-24T16:26:23.8563980+07:00`;
  - `scripts\Invoke-P6EScannerInventory.ps1`, last-write `2026-08-24T16:26:15.2802225+07:00`;
  - `scripts\Invoke-P6EAnvien.ps1`, last-write `2026-08-24T16:25:42.8048068+07:00`;
  - `scripts\Invoke-P6EAARun.ps1`, last-write `2026-08-24T16:24:28.1846921+07:00`.
- No `benchmarks\aa-baseline` directory exists at that snapshot; therefore no A/A invocation exists and A/A remains exactly `0/20`.
- No newer harness seal, corpus seal, graph authority seal, or durable Worker handoff exists at outgoing transfer time.
- Thread-status snapshots can remain stale after app reset; current durable filesystem writes prove the Worker is active. Do not interrupt it solely because visible commentary is old.

Read these Worker files FULL RAW before any Worker direction:

- `README.md`.
- `preflight/lifecycle.md`.
- `preflight/frozen_boundary.json`.
- `preflight/post-handoff-corpus-freeze.seal.json`.
- `preflight/post-handoff-graph-authority-corpus.seal.json`.
- `preflight/aa-baseline-protocol.md`.
- `costmap/living-cost-map-v1.md`.
- `graph/impact-summary.md`.
- Latest provenance `preflight/runtime-identity-independent-supersession-260824_155835.json`.

Historical `2,173`-row corpus, manifest `c81d33e3acdf21ca7689b351ad53230cf693b1203d9b72f87e9fcd20c821d853`, graph `2,173/765/0`, `122,696/169,347`, and old runtime `73,749,504 / 64EECEBA...DEBB` are lineage only. They are not current P6-E corpus, graph, A/A, baseline, or runtime authority.

## Required Worker continuation

1. Let the same Worker finish and seal one compiled fail-closed producer/harness contract, including focused lifecycle/schema/cleanup tests.
2. Produce two consecutive full-row-equal current corpus inventories excluding only the exact Worker durable root and including this Main handoff, the independent Supervisor report, every current protected report, and all Owner rows.
3. Run one fresh excluded graph/storage authority and current file-detail/impact rerun. This graph-authority invocation is not A/A.
4. Run exactly `10` valid A/A pairs / `20` unprofiled invocations, no early stop. Invalid runs remain durable and execution continues until exactly ten valid pairs exist.
5. Worker stops before U1 and returns the raw A/A summary plus numeric proposal to Main.
6. Main opens one separate clean visible Architect. Architect alone owns and seals numeric materiality/resource rules; Owner receives no intermediate decision request.
7. Main relays the sealed Architect decision to the same Worker; only that seal opens U1.
8. Preserve ordered U1 -> U2 -> rebaseline/ownership -> U3 -> conditional U4 -> Pareto -> final equivalence and one final P6-E commit boundary.

## Existing lanes and governance Guard

- P6-D Coder `01a02f10-c175-7162-ac32-6df30a4d604a`: idle after accepted handoff; preserve.
- P6-D Supervisor `01a030d2-3ae2-7b33-a57a-c23503703b1f`: idle after PASS; preserve.
- P6-E Worker `01a0319a-62c9-7650-ba6a-ef59dfcf1973`: active harness repair; preserve.
- Runtime-identity Supervisor `01a032e6-d308-7751-b82c-92a4652dc93b`: completed identity-only PASS; preserve.
- Governance messages previously arrived through Guard lineage `01a032b6-7eaf-7b32-94f1-c693be834928`; the latest visible warning transport is `01a032ff-de24-7352-9c5d-39ce5ee0a106`. Preserve/reuse the existing continuous Guard lineage, do not duplicate it, and do not control it as a functional lane.
- Invalid fork `01a032d9-44a3-7fe2-b97b-8a8fe89fe8f3`: no authority, never transfer to it.

## Git/worktree snapshot

- Snapshot: `2026-08-24T16:27:31.3300762+07:00`.
- Branch/HEAD/parent/tree: `master` / `1a9b5310bc3204c610e9e6c673afc214cd049991` / `d02e4eb76ee6e42848ea0ffb627449ac7d9df092` / `f124bc5811ae26b2751c1b09fc8da673b9f4ab1c`.
- Index: empty, `0` staged paths.
- Status before this handoff: `224` rows = `185` exact Worker-root rows + `39` outside-Worker rows.
- This handoff adds exactly one protected outside-Worker row. Worker may independently add, update, or remove only its own in-progress rows under the exact durable root while continuing harness work; therefore the live Worker-row count may change without changing the protected outside boundary.
- `git diff --check`: PASS.
- Preserve Owner deletion `CONTRIBUTING.md`, all four Child 06 ledger edits, Owner aicontext/governance history and bytes, Microsoft PowerShell artifact, every Main handoff, every Supervisor/Coder report, all Worker artifacts, vendor, third_party, and all unrelated/protected rows.
- Never normalize, stash, reset, checkout, broad-stage, broad-clean, or broad-commit this worktree.

## Governance findings to preserve

- Current successor warm lateness is exactly `52.457s — HIGH — VERIFIED VIOLATION`; failed project-ID transport and successful retry are immutable history.
- The outgoing predecessor's clean-successor warm lateness `2388.551s — HIGH`, hard-deadline lateness `1488.551s — HIGH`, invalid successor fork, and Main self-authored Supervisor PASS remain immutable verified history.
- Historical first post-transfer ACK omission, unscoped discovery, and bounded-subagent misclassification remain exactly as recorded in prior sealed handoffs.
- Do not normalize any historical violation, lateness, retraction, or `NO EVIDENCE` disposition.

## Immediate continuation for incoming Main

1. Complete mandatory FULL RAW startup and detached-seal verification.
2. Re-anchor to the independent identity-only PASS and latest Worker provenance without reopening that passed identity gate.
3. Monitor the existing Worker actual durable writes. Do not duplicate, replace, pressure, or interrupt it.
4. Wait for the compiled fail-closed producer/harness seal, then allow current corpus -> graph/storage/current impact -> exactly 10 A/A pairs in the required order.
5. After A/A, open one clean visible Architect; never ask Owner for an intermediate numeric decision.
6. Open an independent P6-E Supervisor only after a complete no-commit candidate and durable Coder handoff.
7. After final Supervisor PASS: planner refresh, fresh current E-only detect, exact cleanup disposition, exact isolated staging, one P6-E commit under standing AGENTS Rule 12 authority, no push.
8. Initialize the next clean visible Main successor exactly at the incoming warm target and transfer no later than the incoming hard deadline.

## Workspace bans and standing authority

- Sole workspace is `E:\Anvien`, saved project local, no worktree.
- No C-drive workspace/read/write, alternate worktree, target access, implicit network, install/package-script side effects, shared/global/protected/quarantined cache authority, reset/stash/checkout/broad cleanup, unscoped discovery, or push.
- Every `E:\Anvien\.tmp` path is disposable reproducible scratch and never evidence, review/handoff/recovery authority, source of truth, acceptance dependency, or blocker.
- Preserve vendor, third_party, durable reports, Owner work/deletion, current Owner commits, unrelated/protected rows, and all historical evidence.
- Owner `PAUSE` or `STOP` halts commands, edits, QA, commits, and lane control immediately.
- Commit after completed accepted P6-E is standing AGENTS Rule 12 authority; never ask Owner to grant it again.
- No push unless Owner explicitly requests push.

Outgoing Main terminates absolutely after official transfer dispatch.
