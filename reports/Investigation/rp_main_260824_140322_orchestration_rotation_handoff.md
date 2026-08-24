# Visible Main Orchestration Rotation Handoff

## Authority transfer record

- Campaign: Anvien Graph Accuracy / Analyze Performance.
- Outgoing Visible Main: `01a03263-2bc6-71f0-bb18-00b4eeaedb1e`.
- Prepared incoming successor: `01a03290-28f5-7870-a643-871af4aa24d4`.
- Raw outgoing official authority UUIDv7: `01a03266-f548-7920-83ff-690d3bb97da0`.
- Authority start: `2026-08-24T13:13:23.400+07:00`.
- Warm target: `2026-08-24T13:58:23.400+07:00`.
- Hard transfer deadline: `2026-08-24T14:13:23.400+07:00`.
- Successor creation UUIDv7: `01a03290-28f5-7870-a643-871af4aa24d4`.
- Successor creation time: `2026-08-24T13:58:23.605+07:00`.
- Exact warm lateness: `0.205s — HIGH`. Preserve exactly; do not normalize or downgrade.
- The same-directory fork was created from this exact current task without project-registry discovery. Two ACK-only follow-ups completed with no visible assistant payload; no successor action, file/skill/tool read, or lane control is evidenced. Do not create another successor. The official transfer follow-up must require the complete corrective ACK before the successor acts.
- The raw UUIDv7 of the follow-up headed exactly `OFFICIAL AUTHORITY TRANSFER` is transport source of truth for the incoming authority time. Incoming Main computes its warm target at `+45 minutes` and hard deadline at `+60 minutes` from that UUIDv7.
- Outgoing Main terminates absolutely immediately after official dispatch.

## Mandatory iron startup and anti-summary rule

Before every orchestration action, incoming Main must read FULL RAW sequentially through EOF:

1. `E:\Anvien\AGENTS.md`.
2. `E:\Anvien\.agents\skills\working-rules\SKILL.md`.
3. `E:\Anvien\.agents\skills\orchestration\SKILL.md`.
4. This entire sealed handoff and verify the detached seal supplied in the official transfer message.

Before plan, ledger, lane-control, corpus, baseline, implementation authorization, or acceptance action, continue FULL RAW through EOF:

1. `E:\Anvien\.agents\skills\planner\SKILL.md`.
2. `E:\Anvien\docs\plans\2026-07-26-anvien-graph-accuracy-multi-plan\2026-07-26-anvien-graph-accuracy-roadmap.md`.
3. All four Child 06 ledgers under `E:\Anvien\docs\plans\2026-07-26-anvien-graph-accuracy-multi-plan\2026-07-28-06-ambient-external-resolution-and-diagnostics\`.
4. Prior sealed handoff `E:\Anvien\reports\Investigation\rp_main_260824_131000_orchestration_rotation_handoff.md`.
5. Current P6-E worker checkpoint files named below.

Do not use summary, compacted context, prior memory, or file names instead of raw source. After auto-compaction, perform orchestration section 7 FULL RAW re-anchor and continue from the first uncompleted gate; do not reopen a PASSed gate without proved invalidation.

The independent governance Guard is task `01a031fd-366f-79f2-a9e3-04c365403fd3`. Main must not read/use `governance-rule-guard`, message/control the Guard, create a replacement Guard, or self-assume Guard authority.

No Main/successor operation may call default/shared/unscoped project or registry discovery. The exact current task and saved project `E:\Anvien` are already authoritative. Successor creation/rotation must remain attached to the exact current task and same directory.

## Current campaign state

- Sole workspace: `E:\Anvien`, saved project local, no worktree.
- Branch: `master`.
- Current HEAD at the pre-handoff snapshot: `1a9b5310bc3204c610e9e6c673afc214cd049991`.
- Parent: `d02e4eb76ee6e42848ea0ffb627449ac7d9df092`.
- Tree: `f124bc5811ae26b2751c1b09fc8da673b9f4ab1c`.
- P6-D is CLOSED at accepted implementation commit `81163e39718b94a509e41114cada224e8f269e36`; do not reopen accepted gates absent proved invalidation.
- P6-E is the sole open Child 06 slice and exactly one commit boundary: `U1 -> U2 -> rebaseline/ownership -> U3 -> conditional U4 -> Pareto -> final equivalence`.
- U1-U4 are internal ordered units, not independent slices, gates, or commits. Pn-A and later gates remain locked.
- No valid current P6-E corpus seal, current graph/storage authority, valid A/A invocation, comparable baseline, U1, stage, commit, push, or target access exists.
- Approximate whole-slice progress remains about `15%`. This is status estimation only, never acceptance evidence.

## Owner-owned concurrent repository updates

- Owner explicitly confirmed that `internal/aicontext/skills/governance-rule-guard/SKILL.md` and `internal/aicontext/aicontext.go` are Owner updates and instructed Main/Worker to leave them alone.
- Main did not read/use the governance skill content. It inspected only Git metadata/blob identities required to classify corpus drift.
- Owner governance history is now two clean one-path commits: `943893e70e6a947ed964103711dd96703896eee1` over former campaign HEAD `599e401548bd69e5214148a9d2ffd515a7e3305d`, then `d02e4eb76ee6e42848ea0ffb627449ac7d9df092` over `943893e...`. Current `aicontext.go` commit `1a9b5310bc3204c610e9e6c673afc214cd049991` is one path over `d02e4eb...`.
- These are external/user-owned current corpus bytes, not Worker deviation or P6-E blockers. Preserve exactly; do not read/use/edit/revert/normalize/stage/commit them as P6-E work.
- They must still be bound as ordinary current repository/corpus identity. Only identity change between consecutive corpus observations or within/between benchmark pair runs invalidates that seal/run under the same-condition protocol.

## Binding Owner and Architect authority

- Never ask Owner to approve, choose, review, or comment on numeric materiality/resource rules or any other intermediate P6-E decision.
- Worker executes exactly `10` valid A/A pairs / `20` unprofiled invocations, no early stop, then stops before U1 and hands the raw summary plus numeric proposal to Main.
- Main opens one separate visible Architect lane. Architect alone owns and seals numeric materiality/resource rules.
- Main relays the sealed Architect decision to the existing Worker; only that seal opens U1.
- Owner receives only the final completed campaign result unless Owner independently issues `PAUSE`, `STOP`, or a scope override.

## Active P6-E Worker

- Task: `01a0319a-62c9-7650-ba6a-ef59dfcf1973`.
- Saved project: `E:\Anvien`; local; no worktree.
- Durable root: `E:\Anvien\reports\coder\p6e_analyze_performance_260824_093855_raw`.
- Scratch root: `E:\Anvien\.tmp\p6e-worker-01a0319a`; disposable reproducible scratch only, never evidence/authority/blocker.
- Latest direct snapshot: `ACTIVE / HARNESS_REPAIR_IN_PROGRESS / RAW_REANCHOR_AFTER_COMPACTION`.
- Preserve and monitor. Do not duplicate, replace, stop, pressure, archive, or transfer this Worker.
- Worker owns the harness/schema, fresh corpus/graph/impact, exactly 10 valid A/A pairs, raw numeric proposal, and ordered implementation only after Main relays the sealed Architect decision. Main is not the Worker.
- Current harness progress: digest helper focused compile/tests and real Ladybug `STRING[]` parity previously PASS; bounded `aa_stats_harness` output remains valid input under working-rules section 7; Worker independently found and repaired producer/comparator/lifecycle issues.
- Corrections 1-4 recorded successive external Owner HEAD/dirty-byte transitions. Current correction 4 authorizes only exact clean HEAD `1a9b5310...`; prior identities are provenance history.
- Common identity smoke passed before the last Owner commit; retained-process/resource helper passed after fixing empty-index hashing, status-row emission, ordered-dictionary numeric aggregation, and writer flush/dispose-before-hash ordering.
- Worker compacted and is raw re-anchoring the roadmap and all four Child 06 ledgers before resuming technical action. Latest snapshot had roadmap and plan read through EOF and three ledgers remaining.
- A focused scanner/digest harness build was waiting on raw re-anchor and Rule 13 holder clearance. Fresh corpus/graph/A/A/U1 have not started; A/A remains exactly `0/20`.

Read these Worker files FULL RAW before directing the lane:

- `README.md`.
- `preflight/lifecycle.md`.
- `preflight/frozen_boundary.json`.
- `preflight/post-handoff-corpus-freeze.seal.json`.
- `preflight/post-handoff-graph-authority-corpus.seal.json`.
- `preflight/aa-baseline-protocol.md`.
- `costmap/living-cost-map-v1.md`.
- `graph/impact-summary.md`.

All are below the durable root. Their older `2,173`-row/graph states are lineage only. Incoming Main must also snapshot current Worker progress because harness files are actively changing under the durable root.

## Harness and corpus continuation

- Historical corpus `2,173` rows, manifest `c81d33e3acdf21ca7689b351ad53230cf693b1203d9b72f87e9fcd20c821d853`, and graph `2,173/765/0`, `122,696/169,347`, file projection `20,351/479` remain lineage only.
- Multiple protected Main handoffs, synchronized ledger bytes, Owner commits, and current Worker artifacts invalidate that historical authority for A/A.
- Worker must finish and seal one compiled/fail-closed producer contract. Synthetic statistics PASS is not an A/A invocation.
- Then produce two consecutive full-row-equal corpus inventories excluding only the exact Worker durable root while including every current protected and Owner-owned repository row.
- Then run one fresh excluded graph/storage authority and current file-detail/impact rerun. The graph-authority invocation is not an A/A invocation.
- Only after those gates pass may the Worker run exactly `10` valid A/A pairs.

Frozen A/A protocol remains binding:

- exactly `10` valid pairs / `20` unprofiled invocations;
- alternating labels `L/R`, no early stop, no statistical outlier exclusion;
- invalid runs remain durable and execution continues until exactly 10 valid pairs exist;
- identical runtime/source/corpus/storage/cache/instrumentation and exact output/reader identity;
- stop before U1 and route the numeric proposal to a separate visible Architect through Main.

## Current cost and impact lineage

- First profiled run remains `ATTRIBUTION_ATTEMPT_ONLY`, exit `0`, wall `571,249.2723 ms`; it observed Worker-owned rows and is forbidden as graph authority, A/A, comparable baseline, numeric threshold, or speedup proof.
- Historical disjoint profiled phase wall: resolution `499,895.9298 ms`, DB load `30,635.1092 ms`, parse `14,599.2618 ms`, semantic enrichment `3,928.1245 ms`, cross-file binding `1,627.2942 ms`, other named phases `2,021.7028 ms`, explicit residual `14,818.9314 ms`, benchmark `567,526.3537 ms`, wrapper residual `3,722.9186 ms`, wrapper wall `571,249.2723 ms`.
- Inclusive CPU/allocation rows overlap and must never be summed into wall savings.
- Historical file impacts remain CRITICAL lineage and require current rerun before first product edit: `indexes.go 253/38`, `resolve.go 155/48`, `export_resolution.go 163/23`, `diagnostics.go 107/31` (`impacted symbols / affected files`).
- HIGH/CRITICAL is a scope warning requiring surgical edits and affected-boundary regression, not an edit prohibition.

## Existing visible lanes

- P6-D Coder `01a02f10-c175-7162-ac32-6df30a4d604a`: idle after accepted handoff; preserve.
- P6-D independent Supervisor `01a030d2-3ae2-7b33-a57a-c23503703b1f`: idle after PASS; preserve.
- Current independent Guard `01a031fd-366f-79f2-a9e3-04c365403fd3`: active independently; preserve without messaging or control.
- P6-E Worker `01a0319a-62c9-7650-ba6a-ef59dfcf1973`: active; preserve and monitor.

## Git/worktree snapshot

- Snapshot: `2026-08-24T14:02:13.6925281+07:00`.
- Branch/HEAD/parent/tree: `master` / `1a9b5310bc3204c610e9e6c673afc214cd049991` / `d02e4eb76ee6e42848ea0ffb627449ac7d9df092` / `f124bc5811ae26b2751c1b09fc8da673b9f4ab1c`.
- Index: empty, `0` staged rows.
- Status before this handoff: `217` rows = `182` exact Worker-root rows + `35` outside-Worker rows. This handoff adds one protected outside-Worker row; active Worker artifacts may increase independently.
- `git diff --check`: PASS.
- Preserve Owner deletion `CONTRIBUTING.md`, Microsoft PowerShell artifact, every historical Main handoff, P6-E plan-acceptance report, all protected P6-D outputs, all four ledger edits, every Worker artifact, and every Owner aicontext/governance byte.
- Never normalize, stash, reset, checkout, broad-stage, broad-clean, or broad-commit this worktree.

## Verified governance findings

- This outgoing Main's first post-transfer commentary omitted exact `UNDERSTOOD`, campaign goal, P6-E sole-open-slice, and Main-not-Worker boundary: `MEDIUM — VERIFIED VIOLATION`. Corrective ACK was issued without restarting passed raw gates.
- This outgoing Main called unscoped `list_projects {}` despite exact E-only saved-project authority: `HIGH — VERIFIED VIOLATION`. The output was discarded and must never be used; no repeat/default/shared discovery is allowed.
- Successor warm creation was late exactly `0.205s — HIGH — VERIFIED VIOLATION`. Preserve exactly.
- Prior successor warm lateness `44.162s — HIGH`, `25.844s — HIGH`, `0.240s — HIGH`, and `0.154s — HIGH` remain immutable history.
- Prior outgoing Main's misapplication of orchestration section 2 to a valid bounded Worker subagent remains `HIGH — VERIFIED VIOLATION`; Owner corrected it, the finding against Worker was fully withdrawn, and it must never be resurrected.
- Successor pre-transfer turns returned no visible assistant payload. This is recorded as an observed transport/task state, not normalized into an ACK. Official transfer must require corrective ACK before any action and must not create a duplicate successor.

## Immediate continuation for incoming Main

1. Return exact corrective `UNDERSTOOD` with campaign goal, P6-E sole open slice, Main-only orchestration/monitoring boundary, first action, and bans; then complete mandatory FULL RAW startup and detached-seal verification.
2. Snapshot the existing Worker without interrupting or duplicating it.
3. Relay the exact sealed identity of this handoff to the Worker as `ROTATION_HANDOFF_SEALED`; this handoff row must be included in the fresh corpus.
4. Allow the Worker to finish raw re-anchor, current correction-4 harness guard, focused harness build/tests, and the one fail-closed producer-contract seal.
5. Allow one fresh full-row-equal corpus reseal, one fresh excluded graph/storage authority and current impact rerun, then exactly 10 valid A/A pairs.
6. After A/A, receive raw summary/proposal and open one separate visible Architect lane. Do not ask Owner.
7. Relay the sealed Architect decision to the existing Worker; preserve ordered U1-U4 and one final P6-E commit only.
8. Open one independent P6-E Supervisor only after complete no-commit candidate and durable Coder handoff.
9. After Supervisor PASS: planner refresh, fresh current E-only detect, exact cleanup disposition, exact isolated staging, one P6-E commit under standing AGENTS Rule 12 authority, and no push.
10. Initialize the next visible Main successor at the exact new warm target and transfer no later than the new hard deadline.

## Workspace bans and standing authority

- Sole workspace is `E:\Anvien`, local saved project, no worktree.
- No C-drive workspace/read/write, alternate worktree, target access, implicit network, install/package-script side effects, shared/global/protected/quarantined cache authority, or unscoped project/registry discovery.
- Every `E:\Anvien\.tmp` path is disposable reproducible scratch and never evidence, review/handoff/recovery authority, source of truth, acceptance dependency, or procedure blocker.
- Preserve vendor, third_party, durable reports, Owner work/deletion, unrelated/protected rows, all historical evidence, and current Owner aicontext/governance work.
- Owner `PAUSE` or `STOP` halts commands, edits, QA, commits, and lane control immediately.
- Commit after completed accepted P6-E is standing AGENTS Rule 12 authority; never ask Owner to grant it again.
- No push unless Owner explicitly requests push.

Outgoing Main terminates absolutely after official transfer dispatch.
