# Visible Main Orchestration Rotation Handoff

## Authority transfer record

- Campaign: Anvien Graph Accuracy / Analyze Performance.
- Outgoing Visible Main: `01a0322e-b994-7d03-a4bd-b78fa772b8ff`.
- Prepared incoming successor: `01a03263-2bc6-71f0-bb18-00b4eeaedb1e`.
- Raw outgoing official authority UUIDv7: `01a03239-4c64-7422-8ff7-e3cbdccb86ea`.
- Authority start: `2026-08-24T12:23:31.044+07:00`.
- Warm target: `2026-08-24T13:08:31.044+07:00`.
- Hard transfer deadline: `2026-08-24T13:23:31.044+07:00`.
- Successor creation UUIDv7: `01a03263-2bc6-71f0-bb18-00b4eeaedb1e`.
- Successor creation time: `2026-08-24T13:09:15.206+07:00`.
- Exact warm lateness: `44.162s — HIGH`. Preserve exactly; do not normalize or downgrade.
- Successor returned `UNDERSTOOD` and `WAITING_FOR_OFFICIAL_TRANSFER`; it remains idle with no file, skill, tool, workspace, or lane action. Do not create another successor.
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
4. Prior sealed handoff `E:\Anvien\reports\Investigation\rp_main_260824_122100_orchestration_rotation_handoff.md`.
5. Current P6-E worker checkpoint files named below.

Do not use summary, compacted context, prior memory, or file names instead of raw source. After auto-compaction, perform orchestration section 7 FULL RAW re-anchor and continue from the first uncompleted gate; do not reopen a PASSed gate without proved invalidation.

The independent governance Guard is task `01a031fd-366f-79f2-a9e3-04c365403fd3`. Main must not read/use `governance-rule-guard`, message/control the Guard, create a replacement Guard, or self-assume Guard authority.

## Current campaign state

- Sole workspace: `E:\Anvien`, saved project local, no worktree.
- Branch: `master`.
- Current HEAD: `599e401548bd69e5214148a9d2ffd515a7e3305d`.
- Parent: `81163e39718b94a509e41114cada224e8f269e36`.
- Tree: `cf35c431962627167f83619c7512ded78705c202`.
- P6-D is CLOSED at accepted implementation commit `81163e39718b94a509e41114cada224e8f269e36`; do not reopen accepted gates absent proved invalidation.
- P6-E is the sole open Child 06 slice and exactly one commit boundary: `U1 -> U2 -> rebaseline/ownership -> U3 -> conditional U4 -> Pareto -> final equivalence`.
- U1-U4 are internal ordered units, not independent slices, gates, or commits. Pn-A and later gates remain locked.
- No P6-E production/test edit, valid A/A invocation, comparable baseline, U1, stage, commit, push, or target access exists.
- Approximate whole-slice progress remains about `15%`. This is a status estimate only, never acceptance evidence.

## Binding Owner and Architect authority

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
- Latest direct snapshot: `ACTIVE / HARNESS_REPAIR_IN_PROGRESS`.
- Preserve and monitor. Do not duplicate, replace, stop, pressure, archive, or transfer this worker.
- Worker owns harness/schema, fresh corpus/graph/impact, exactly 10 valid A/A pairs, raw proposal, and ordered implementation after Main relays the sealed Architect decision. Main is not the worker.
- Independent P6-E Supervisor opens only after the complete no-commit candidate and durable Coder handoff.

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

## Current worker checkpoint and harness state

- Worker completed raw authority re-anchor through EOF after compaction, including roadmap, all four current Child 06 ledgers, causal/architecture/workflow/feasibility/plan-acceptance reports, and P6-D PASS.
- Digest helper is worker-owned. Focused `go build -tags ladybugdb` passed with retained exit `0`; scratch executable was `7,365,148` bytes / SHA-256 prefix `0C729694...A9A`.
- Focused digest tests passed, including real Ladybug `STRING[]` CSV-to-COPY-to-native roundtrip and preservation of pre-existing scratch on rejection.
- The bounded hidden subagent `aa_stats_harness` is allowed under working-rules section 7. It touched exactly `Compare-P6EAAPair.ps1`, `Summarize-P6EAA.ps1`, and `Test-P6EAAStatistics.ps1`, then stopped at its boundary with synthetic `10`-pair / `20`-run and negative-case PASS. Worker must still independently review and accept its output.
- Worker review correctly found that comparator live re-hashing would falsely invalidate a prior run after the next valid run replaces normal graph/Ladybug output; run-scoped durable digest/command seals must be authoritative instead.
- Worker also found cleanup pre-existence/reparse gaps and a producer schema mismatch: wrapper still consumed old `graph.semanticSha256` / `persistence.*`, while digest helper emits `graphJson` / `projection`; scanner freeze still used `go run` and lacked protected-byte manifest/process-handle seal.
- Worker is converging one producer contract using compiled scanner/digest executables, exact run identity, fail-closed lifecycle, pre/post full corpus rows, protected bytes, storage/output semantic digests, and sampled resource summary.
- Fresh post-current-handoff corpus freeze, graph authority, A/A, and U1 have not started. Do not treat harness compile/tests or synthetic statistics as an A/A invocation.

## Corpus and graph authority correction

- Historical corpus `2,173` rows, manifest `c81d33e3acdf21ca7689b351ad53230cf693b1203d9b72f87e9fcd20c821d853`, and graph `2,173/765/0`, `122,696/169,347`, file projection `20,351/479` remain lineage only for A/A.
- Protected handoffs `rp_main_260824_112500...`, `rp_main_260824_122100...`, this handoff, and synchronized ledger bytes exist outside the exact worker-root exclusion. Old corpus/graph seals are not current A/A authority.
- After this handoff is sealed, worker must finish and seal the harness, perform one fresh full-row-equal corpus reseal including every protected row, then one fresh excluded graph/storage authority and current file-detail/impact rerun.
- Only after those gates pass may the worker run exactly `10` valid A/A pairs. The graph-authority run is not one of the 20 A/A invocations.

Frozen A/A protocol remains binding:

- exactly `10` valid pairs / `20` unprofiled invocations;
- alternating labels `L/R`, no early stop, no statistical outlier exclusion;
- invalid runs remain durable and execution continues until exactly 10 valid pairs exist;
- identical runtime/source/corpus/storage/cache/instrumentation and exact output/reader identity;
- stop before U1 and route numeric proposal to separate visible Architect through Main.

## Current cost and impact lineage

- First profiled run remains `ATTRIBUTION_ATTEMPT_ONLY`, exit `0`, wall `571,249.2723 ms`; it observed worker-owned rows and is forbidden as graph authority, A/A, comparable baseline, numeric threshold, or speedup proof.
- Historical disjoint profiled phase wall: resolution `499,895.9298 ms`, DB load `30,635.1092 ms`, parse `14,599.2618 ms`, semantic enrichment `3,928.1245 ms`, cross-file binding `1,627.2942 ms`, other named phases `2,021.7028 ms`, explicit residual `14,818.9314 ms`, benchmark `567,526.3537 ms`, wrapper residual `3,722.9186 ms`, wrapper wall `571,249.2723 ms`.
- Inclusive CPU/allocation rows overlap and must never be summed into wall savings.
- Historical file impacts remain CRITICAL lineage and require current rerun before first product edit: `indexes.go 253/38`, `resolve.go 155/48`, `export_resolution.go 163/23`, `diagnostics.go 107/31` (`impacted symbols / affected files`).
- HIGH/CRITICAL is a scope warning requiring surgical edits and affected-boundary regression, not an edit prohibition.

## Worker hidden-subagent rule correction

- Outgoing Main incorrectly applied orchestration section 2 to the worker and interrupted bounded `aa_stats_harness` work.
- Owner clarified that orchestration hidden-lane restrictions bind Main Orchestration only. Worker may use bounded hidden subagents under working-rules section 7 when the subtask is independent, concrete, and has a clear boundary; output remains input that the worker must verify.
- The event is `HIGH — VERIFIED VIOLATION` by outgoing Main, not a worker finding. Main fully withdrew the worker classification, authorized resume/reuse, preserved valid bytes/provenance, and did not reset any gate.
- Incoming Main must not resurrect this as evidence against the worker or P6-E. Evaluate worker subagents by working-rules section 7 and exact assigned boundary.

## Existing visible lanes

- P6-D Coder `01a02f10-c175-7162-ac32-6df30a4d604a`: idle after accepted handoff; preserve.
- P6-D independent Supervisor `01a030d2-3ae2-7b33-a57a-c23503703b1f`: idle after PASS; preserve.
- Current independent Guard `01a031fd-366f-79f2-a9e3-04c365403fd3`: active independently; preserve without messaging or control.
- P6-E Worker `01a0319a-62c9-7650-ba6a-ef59dfcf1973`: active; preserve and monitor.

## Git/worktree snapshot

- Snapshot: `2026-08-24T13:10:23.2000676+07:00`.
- Branch/HEAD/parent/tree: `master` / `599e401548bd69e5214148a9d2ffd515a7e3305d` / `81163e39718b94a509e41114cada224e8f269e36` / `cf35c431962627167f83619c7512ded78705c202`.
- Index: empty, `0` staged rows.
- Status before this handoff: `213` rows = `179` exact worker-root rows + `34` protected outside-worker rows. This handoff adds one protected outside-worker row; active worker artifacts may increase independently.
- `git diff --check`: PASS.
- Preserve Owner deletion `CONTRIBUTING.md`, Microsoft PowerShell artifact, every historical Main handoff, P6-E plan-acceptance report, all protected P6-D outputs, all four ledger edits, and every worker artifact.
- Never normalize, stash, reset, checkout, broad-stage, broad-clean, or broad-commit this worktree.

## Verified governance findings

- Mandatory first post-transfer ACK lacked exact goal, sole-open-slice, and Main-not-worker boundary fields: `MEDIUM — VERIFIED VIOLATION`. Corrective ACK was issued without restarting passed raw gates.
- Outgoing Main misapplied orchestration section 2 to the worker and interrupted a valid bounded hidden subagent: `HIGH — VERIFIED VIOLATION`. Owner corrected it; Main withdrew the worker finding and restored the permitted boundary without gate reset.
- Current successor warm creation was late exactly `44.162s — HIGH`. Preserve exactly.
- Historical warm lateness `25.844s — HIGH`, `0.240s — HIGH`, and `0.154s — HIGH` remain immutable history.
- Preserve prior documented orchestration violations/corrections from earlier sealed handoffs; do not normalize or resurrect withdrawn claims.

## Immediate continuation for incoming Main

1. Complete mandatory FULL RAW startup and detached-seal verification.
2. Snapshot the existing worker without interrupting or duplicating it.
3. Permit worker to complete and seal the one producer-contract harness, including independent review of bounded subagent output and exact fail-closed lifecycle tests.
4. After this handoff row is sealed, allow one fresh full-row-equal corpus reseal, one fresh excluded graph/storage authority, and current impact rerun; then exactly 10 valid A/A pairs.
5. After A/A, receive raw summary/proposal and open one separate visible Architect lane. Do not ask Owner.
6. Relay the sealed Architect decision to the existing worker; preserve ordered U1-U4 and one final P6-E commit only.
7. Open one independent P6-E Supervisor only after complete no-commit candidate and durable Coder handoff.
8. After Supervisor PASS: planner refresh, fresh current E-only detect, exact cleanup disposition, exact isolated staging, one P6-E commit under standing AGENTS Rule 12 authority, and no push.
9. Initialize the next visible Main successor at the exact new warm target and transfer no later than the new hard deadline.

## Workspace bans and standing authority

- Sole workspace is `E:\Anvien`, local saved project, no worktree.
- No C-drive workspace/read/write, alternate worktree, target access, implicit network, install/package-script side effects, or shared/global/protected/quarantined cache authority.
- Every `E:\Anvien\.tmp` path is disposable reproducible scratch and never evidence, review/handoff/recovery authority, source of truth, acceptance dependency, or procedure blocker.
- Preserve vendor, third_party, durable reports, Owner work/deletion, unrelated/protected rows, and all historical evidence.
- Owner `PAUSE` or `STOP` halts commands, edits, QA, commits, and lane control immediately.
- Commit after completed accepted P6-E is standing AGENTS Rule 12 authority; never ask Owner to grant it again.
- No push unless Owner explicitly requests push.

Outgoing Main terminates absolutely after official transfer dispatch.
