# Visible Main Orchestration Rotation Handoff

## Authority transfer record

- Campaign: Anvien Graph Accuracy / Analyze Performance.
- Outgoing Visible Main: `01a031fa-6286-7d11-a4d2-d1b8363ffe17`.
- Prepared incoming successor: `01a0322e-b994-7d03-a4bd-b78fa772b8ff`.
- Raw outgoing official authority UUIDv7: `01a03205-21c0-7d92-907e-7f7dd9f797e3`.
- Authority start: `2026-08-24T11:26:32.256+07:00`.
- Warm target: `2026-08-24T12:11:32.256+07:00`.
- Hard transfer deadline: `2026-08-24T12:26:32.256+07:00`.
- Successor creation UUIDv7: `01a0322e-b994-7d03-a4bd-b78fa772b8ff`.
- Successor creation time: `2026-08-24T12:11:58.100+07:00`.
- Exact warm lateness: `25.844s — HIGH`. Preserve exactly; do not normalize or downgrade.
- Successor returned `UNDERSTOOD / WAITING_FOR_OFFICIAL_TRANSFER` and remains idle. Do not create another successor.
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
4. Prior sealed handoff `E:\Anvien\reports\Investigation\rp_main_260824_112500_orchestration_rotation_handoff.md`.
5. Current P6-E worker checkpoint files named below.

Do not use summary, compacted context, prior memory, or file names instead of raw source. After auto-compaction, perform orchestration section 7 FULL RAW re-anchor and continue from the first uncompleted gate; do not reopen a PASSed gate without proved invalidation.

The independent governance Guard is currently task `01a031fd-366f-79f2-a9e3-04c365403fd3`. Old Guard `01a0307e-fe39-73c2-83df-e2bb234143a5` is historical only. The pre-transfer preparation prompt sent to this successor contained that stale ID; this handoff corrects the topology. Main must not read/use `governance-rule-guard`, message/control the Guard, create a replacement Guard, or self-assume Guard authority.

## Current campaign state

- Sole workspace: `E:\Anvien`, saved project local, no worktree.
- Branch: `master`.
- Current HEAD: `599e401548bd69e5214148a9d2ffd515a7e3305d`.
- Parent: `81163e39718b94a509e41114cada224e8f269e36`.
- Tree: `cf35c431962627167f83619c7512ded78705c202`.
- P6-D is CLOSED at accepted implementation commit `81163e39718b94a509e41114cada224e8f269e36`; do not reopen accepted gates absent proved invalidation.
- P6-E is the sole open Child 06 slice and exactly one commit boundary: `U1 -> U2 -> rebaseline/ownership -> U3 -> conditional U4 -> Pareto -> final equivalence`.
- U1-U4 are internal ordered units, not independent slices, gates, or commits. Pn-A and later gates remain locked.
- No P6-E production/test edit, A/A invocation, comparable baseline, U1, stage, commit, push, or target access exists.
- Approximate whole-slice progress remains about `15%`; this is a status estimate only, never acceptance evidence.

## Four-ledger planner synchronization

The four Child 06 ledgers were synchronized before this handoff. Roadmap, production, and tests were not edited by that synchronization. Binding wording now states: exactly `10` valid A/A pairs; worker stops before U1; one separate visible Architect alone seals numeric materiality/resource rules; Main relays the seal; Owner receives no intermediate request.

Strict identities at the handoff boundary:

- Plan: `124,309` bytes / `677` LF / `0` CR / SHA-256 `93DDF8703BF0ED772FFEB6FD1435A78FCF631596C88D52C3C4A6444294AD1C37`.
- Evidence: `166,480` bytes / `707` LF / `0` CR / SHA-256 `97A026C79189429F16D96A44FB91A5B52CD36A70BE70EBC3B6D72C7C57A39CA6`.
- Benchmark: `35,995` bytes / `161` LF / `0` CR / SHA-256 `8E61F3B5BAE6E438F1517BE0D734E52F6CF3C09FEE1C5B642DD9A8D1D0FE5362`.
- Actual status: `84,112` bytes / `307` LF / `0` CR / SHA-256 `42F93E6EF3B67BA31412DA1502BE4E6329FFC8287330102B74ED3EF258ECA0F0`.

All are strict UTF-8 without BOM. Their document impacts are LOW/zero. `git diff --check` passes.

## Binding Owner correction

- Never ask Owner to approve, choose, review, or comment on numeric materiality/resource rules or any other intermediate P6-E decision.
- Worker executes exactly `10` valid A/A pairs / `20` unprofiled invocations, no early stop, then stops before U1 and hands the raw summary plus numeric proposal to Main.
- Main opens one separate visible Architect lane. Architect alone owns and seals numeric materiality/resource rules.
- Main relays the sealed Architect decision to the existing worker; only that seal opens U1.
- Owner receives only the final completed campaign result unless Owner independently issues `PAUSE`, `STOP`, or a scope override.

## Active P6-E worker

- Task: `01a0319a-62c9-7650-ba6a-ef59dfcf1973`.
- Saved project: `E:\Anvien`; local; no worktree.
- Durable root: `E:\Anvien\reports\coder\p6e_analyze_performance_260824_093855_raw`.
- Scratch root: `E:\Anvien\.tmp\p6e-worker-01a0319a`; disposable reproducible scratch only, never evidence/authority/blocker.
- Latest direct snapshot: `ACTIVE`. Worker has read the roadmap and all four current ledgers and is raw re-anchoring P6-E/P6-D reports. It is holding corpus freeze, graph refresh, A/A, and U1.
- Preserve and monitor. Do not duplicate, replace, stop, pressure, archive, or transfer this worker.
- Worker owns harness/schema, fresh corpus/graph/impact, exactly 10 valid A/A pairs, raw proposal, then ordered implementation only after Main relays the sealed Architect decision.
- Main is not the worker.

Read these worker files FULL RAW before directing the lane:

- `README.md`.
- `preflight/lifecycle.md`.
- `preflight/frozen_boundary.json`.
- `preflight/post-handoff-corpus-freeze.seal.json`.
- `preflight/post-handoff-graph-authority-corpus.seal.json`.
- `preflight/aa-baseline-protocol.md`.
- `costmap/living-cost-map-v1.md`.
- `graph/impact-summary.md`.

All are below the durable root.

## Corpus and graph authority correction

- Historical seal was `2,173` rows with manifest `c81d33e3acdf21ca7689b351ad53230cf693b1203d9b72f87e9fcd20c821d853`; historical graph authority was `2,173/765/0`, graph `122,696/169,347`, file projection `20,351/479`, DB fallback/failure/skip `0`.
- These values remain lineage only for A/A because protected handoff `rp_main_260824_112500_orchestration_rotation_handoff.md`, the four ledger changes, and this new handoff exist outside the exact worker-root exclusion.
- Worker correctly stopped A/A and U1. Do not hide/delete protected rows or reuse the old seal as current authority.
- After this handoff is sealed, worker may perform one fresh post-handoff full-row-equal corpus reseal including every protected row, then one fresh excluded graph/storage authority and current file-detail/impact rerun.
- Only after those current gates pass may the worker run exactly `10` valid A/A pairs. The graph-authority run is not one of the 20 A/A invocations.

Frozen A/A protocol remains binding:

- exactly `10` valid pairs / `20` unprofiled invocations;
- alternating labels `L/R`, no early stop, no statistical outlier exclusion;
- invalid runs remain durable and execution continues until exactly 10 valid pairs exist;
- identical runtime/source/corpus/storage/cache/instrumentation and exact output/reader identity;
- stop before U1 and route numeric proposal to separate visible Architect through Main.

## Current cost and impact checkpoint

- The first profiled run remains `ATTRIBUTION_ATTEMPT_ONLY`, exit `0`, wall `571,249.2723 ms`; it observed worker-owned rows and is forbidden as graph authority, A/A, comparable baseline, numeric threshold, or speedup proof.
- Historical disjoint profiled phase wall: resolution `499,895.9298 ms`, DB load `30,635.1092 ms`, parse `14,599.2618 ms`, semantic enrichment `3,928.1245 ms`, cross-file binding `1,627.2942 ms`, other named phases `2,021.7028 ms`, explicit residual `14,818.9314 ms`, benchmark `567,526.3537 ms`, wrapper residual `3,722.9186 ms`, wrapper wall `571,249.2723 ms`.
- Inclusive CPU/allocation rows overlap and must never be summed into wall savings.
- Historical file impacts remain CRITICAL lineage and require current rerun before first edit: `indexes.go 253/38`, `resolve.go 155/48`, `export_resolution.go 163/23`, `diagnostics.go 107/31` (`impacted symbols / affected files`).
- Important symbol lineage includes `workspace 145/30/7/71/92`, `buildWorkspace 62/26/7/23/18`, `cleanPath 129/30/5/66/49`, `ResolveBoundInto 120/47/25/31/4`, `AppendDiagnosticToNode 11/7/2/24/6`, and `appendDiagnosticsFromProperties 11/8/2/13/1`.
- HIGH/CRITICAL is a scope warning requiring surgical edits and affected-boundary regression, not an edit prohibition.

## Existing visible lanes

- P6-D Coder `01a02f10-c175-7162-ac32-6df30a4d604a`: idle after accepted handoff; preserve.
- P6-D independent Supervisor `01a030d2-3ae2-7b33-a57a-c23503703b1f`: idle after PASS; preserve.
- Current independent Guard `01a031fd-366f-79f2-a9e3-04c365403fd3`: active independently; preserve without messaging or control.
- P6-E Worker `01a0319a-62c9-7650-ba6a-ef59dfcf1973`: active; preserve and monitor.

## Git/worktree snapshot

- Snapshot: `2026-08-24T12:20:48.8205100+07:00`.
- Branch/HEAD/parent/tree: `master` / `599e401548bd69e5214148a9d2ffd515a7e3305d` / `81163e39718b94a509e41114cada224e8f269e36` / `cf35c431962627167f83619c7512ded78705c202`.
- Index: empty, `0` staged rows.
- Status immediately before this handoff: `209` rows. This handoff adds one protected outside-worker row; active worker artifacts may increase independently.
- `git diff --check`: PASS.
- Preserve Owner deletion `CONTRIBUTING.md`, Microsoft PowerShell artifact, every historical Main handoff, P6-E plan-acceptance report, all protected P6-D outputs, all four ledger edits, and every worker artifact.
- Never normalize, stash, reset, checkout, broad-stage, broad-clean, or broad-commit this worktree.

## Verified governance findings

- Mandatory post-transfer ACK was initially omitted: `MEDIUM — VERIFIED VIOLATION`. Corrective `UNDERSTOOD` response was then issued; startup was continued without restarting passed gates.
- Current successor warm creation was late exactly `25.844s — HIGH`. Preserve exactly.
- Prior warm lateness `0.240s — HIGH` and historical predecessor warm lateness `0.154s — HIGH` remain immutable history.
- The preparation prompt's stale Guard ID is corrected by this handoff; old ID is historical only.
- Preserve prior documented orchestration violations/corrections from earlier sealed handoffs; do not normalize or resurrect withdrawn claims.

## Immediate continuation for incoming Main

1. Complete mandatory FULL RAW startup and detached-seal verification.
2. Snapshot the existing worker without interrupting or duplicating it.
3. Recognize that outgoing Main has sent `ROTATION_HANDOFF_SEALED` with this handoff identity. Allow the existing worker to finish its raw re-anchor, reseal the full corpus once after all current protected rows, refresh excluded graph/storage authority and current impacts, then run exactly 10 valid A/A pairs.
4. After A/A, receive the raw summary/proposal and open one separate visible Architect lane. Do not ask Owner.
5. Relay the sealed Architect decision to the existing worker; preserve ordered U1-U4 and one final P6-E commit only.
6. Open one independent P6-E Supervisor only after complete no-commit candidate and durable Coder handoff.
7. After Supervisor PASS: planner refresh, fresh current E-only detect, exact cleanup disposition, exact isolated staging, one P6-E commit under standing AGENTS Rule 12 authority, and no push.
8. Initialize the next visible Main successor at the exact new warm target and transfer no later than the new hard deadline.

## Workspace bans and standing authority

- Sole workspace is `E:\Anvien`, local saved project, no worktree.
- No C-drive workspace/read/write, alternate worktree, target access, implicit network, install/package-script side effects, or shared/global/protected/quarantined cache authority.
- Every `E:\Anvien\.tmp` path is disposable reproducible scratch and never evidence, review/handoff/recovery authority, source of truth, acceptance dependency, or procedure blocker.
- Preserve vendor, third_party, durable reports, Owner work/deletion, unrelated/protected rows, and all historical evidence.
- Owner `PAUSE` or `STOP` halts commands, edits, QA, commits, and lane control immediately.
- Commit after completed accepted P6-E is standing AGENTS Rule 12 authority; never ask Owner to grant it again.
- No push unless Owner explicitly requests push.

Outgoing Main terminates absolutely after official transfer dispatch.
