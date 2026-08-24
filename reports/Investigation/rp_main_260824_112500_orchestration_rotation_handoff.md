# Visible Main Orchestration Rotation Handoff

## Authority transfer record

- Campaign: Anvien Graph Accuracy / Analyze Performance.
- Outgoing Visible Main: `01a031c8-db4f-7460-8093-2f6669fde256`.
- Incoming prepared successor: `01a031fa-6286-7d11-a4d2-d1b8363ffe17`.
- Raw outgoing official authority UUIDv7: `01a031d1-2eb6-7d23-b6e3-479514b9c842`.
- Transport-source authority start: `2026-08-24T10:29:47.702+07:00`.
- Exact warm successor target: `2026-08-24T11:14:47.702+07:00`.
- Hard transfer deadline: `2026-08-24T11:29:47.702+07:00`.
- Successor creation UUIDv7: `01a031fa-6286-7d11-a4d2-d1b8363ffe17`.
- Successor creation time: `2026-08-24T11:14:47.942+07:00`.
- Exact warm lateness: `0.240s — HIGH`. Preserve exactly; do not normalize or downgrade.
- Successor returned `UNDERSTOOD / WAITING_FOR_OFFICIAL_TRANSFER` and remains idle without file, skill, tool, workspace, or lane action.
- Exact incoming authority time must be derived from the raw UUIDv7 of the follow-up headed exactly `OFFICIAL AUTHORITY TRANSFER`; do not backfill it into this immutable handoff.
- Outgoing Main terminates absolutely immediately after dispatch.

## Mandatory iron startup for successor

Before every orchestration action, read FULL RAW sequentially through EOF:

1. `E:\Anvien\AGENTS.md`.
2. `E:\Anvien\.agents\skills\working-rules\SKILL.md`.
3. `E:\Anvien\.agents\skills\orchestration\SKILL.md`.
4. This entire sealed handoff, then verify its detached identity from the official transfer message.

Before plan, ledger, lane-control, corpus, baseline, or implementation authorization action, continue FULL RAW through EOF:

1. `E:\Anvien\.agents\skills\planner\SKILL.md`.
2. `E:\Anvien\docs\plans\2026-07-26-anvien-graph-accuracy-multi-plan\2026-07-26-anvien-graph-accuracy-roadmap.md`.
3. All four Child 06 ledgers under `E:\Anvien\docs\plans\2026-07-26-anvien-graph-accuracy-multi-plan\2026-07-28-06-ambient-external-resolution-and-diagnostics\`.
4. The current P6-E worker reports/checkpoint named below.

Do not use summaries, compacted context, prior memory, or file names instead of raw source. After any auto-compaction, perform orchestration section 7 FULL RAW re-anchor; this restores context only and never reopens a PASSed gate without proved invalidation.

The independent governance Guard remains task `01a0307e-fe39-73c2-83df-e2bb234143a5`. Main must not read/use `governance-rule-guard`, create a replacement Guard, or self-assume Guard authority.

## Binding Owner correction: no intermediate Owner approval

- Owner must never be asked to approve, choose, review, or comment on numeric materiality/resource rules or any other intermediate P6-E decision.
- After exactly 10 valid A/A pairs, the worker stops before U1 and hands its raw summary plus numeric proposal to Main.
- Main opens one separate visible Architect lane. Architect alone owns and seals the numeric materiality/resource decision.
- Main relays the sealed Architect decision to the existing worker; that decision opens U1.
- Owner receives only the final completed campaign result unless Owner independently issues `PAUSE`, `STOP`, or a scope override.
- Do not preserve the obsolete plan wording that says Owner/Main approval is required. Current worker-owned lifecycle/protocol/cost-map wording already reflects the correction; the four living plan ledgers have not yet been changed for this wording-only synchronization.

## Current campaign state

- Sole workspace: `E:\Anvien`, saved project local, no worktree.
- Branch: `master`.
- Current HEAD: `599e401548bd69e5214148a9d2ffd515a7e3305d`.
- Parent: `81163e39718b94a509e41114cada224e8f269e36`.
- Tree: `cf35c431962627167f83619c7512ded78705c202`.
- P6-D is CLOSED. Its accepted implementation commit is `81163e39718b94a509e41114cada224e8f269e36`; do not update its checklist or rerun accepted gates absent proved invalidation.
- P6-E is the sole open Child 06 slice and exactly one isolated commit boundary containing ordered `U1 -> U2 -> rebaseline/ownership -> U3 -> conditional U4 -> Pareto -> final equivalence`.
- U1-U4 are internal units, not separate slices, gates, or commits. No intermediate U2/U3 commit.
- Pn-A and later Child 06 gates remain locked.
- No P6-E production/test edit, stage, commit, push, target access, A/A run, or comparable baseline exists at this handoff.
- Approximate whole-slice progress remains about `15%`: authority/corpus/current graph/impact/cost-map/protocol preflight is substantially established, while all 20 A/A invocations, Architect decision, U1-U4 implementation, rebaselines, final equivalence, Supervisor, detect, cleanup, and commit remain ahead. Treat the percentage only as a status estimate, never acceptance evidence.

## Active visible P6-E implementation/performance lane

- Task: `01a0319a-62c9-7650-ba6a-ef59dfcf1973`.
- Saved project: `E:\Anvien`; environment local; no worktree.
- Durable worker-owned root: `E:\Anvien\reports\coder\p6e_analyze_performance_260824_093855_raw`.
- Scratch root: `E:\Anvien\.tmp\p6e-worker-01a0319a`; every `.tmp` path is disposable reproducible scratch and never evidence/authority/blocker.
- Latest direct task snapshot at this handoff: `ACTIVE`. Worker reports that governance wording was corrected in four worker-owned artifacts; roadmap/four ledgers remain unchanged. It is designing the A/A harness with unique E-only scratch per run, retained PID/handle, heartbeat, full-row corpus checks, output hashes, and process-resource series. It opened one bounded read-only discovery subtask to identify the exact reusable native-reader digest path. No source edit or A/A was opened.
- Preserve and continuously monitor actual commands, files, scope, and loops. Do not duplicate, replace, stop, pressure, archive, transfer, or disrupt this worker because Main rotated.
- Worker owns source inspection, current impact refresh, harness/cost-map, A/A/baseline, ordered U1-U4 implementation, per-unit A/B/equivalence/resource/rollback proof, and complete no-commit Coder handoff.
- Main is not the worker. Main updates the four ledgers with planner only when actual state changes.
- Independent P6-E Supervisor opens only after a complete no-commit candidate and durable Coder handoff.

## Current worker durable checkpoint

Read these FULL RAW before directing the worker:

- `README.md`.
- `preflight/lifecycle.md`.
- `preflight/frozen_boundary.json`.
- `preflight/post-handoff-corpus-freeze.seal.json`.
- `preflight/post-handoff-graph-authority-corpus.seal.json`.
- `preflight/aa-baseline-protocol.md`.
- `costmap/living-cost-map-v1.md`.
- `graph/impact-summary.md`.

All paths are below the durable worker root above.

Current facts:

- First profiled run remains `ATTRIBUTION_ATTEMPT_ONLY`, exit `0`, wall `571,249.2723 ms`; it observed worker-owned artifact rows and is forbidden as graph authority, A/A, comparable baseline, threshold, or speedup proof.
- Initial helper output with lost handle remains `INVALID_NONAUTHORITATIVE`.
- Exact worker-root scanner exclusion is proved.
- Final corpus is `SEALED_STABLE`: `2,173` ordered rows, manifest SHA-256 `c81d33e3acdf21ca7689b351ad53230cf693b1203d9b72f87e9fcd20c821d853`, zero worker-root rows, exactly one required prior rotation-report row, all identity comparisons true.
- Corpus seal identity: `preflight/post-handoff-corpus-freeze.seal.json`, `2,078` bytes, SHA-256 `799C0298712374273D4D66E6A71905175EA703CBDA45D2A7451276431C57FB4D`.
- Current graph/storage authority is PASS and explicitly not an A/A run: `2,173/765/0`, graph `122,696/169,347`, file projection `20,351/479`, wall `608,132.3919 ms`, benchmark total `605,079,802,800 ns`, benchmark SHA-256 `2bcb601798d94849c63e7e21a49740543f877ad89b1cd95f8d2d01d47f965b18`.
- Post-graph inventory is full-row equal to the sealed corpus; graph-authority corpus seal SHA-256 is `8D4DE09E301E51A2EBC1B00FF766D616AEEA222CFC682AB47252EB8A0B7DDDC9`.
- One unwrapped direct version probe is explicitly invalid/unpinned negative history; do not cite it. Valid pinned runtime remains `E:\Anvien\anvien\bin\anvien.exe`, version `1.2.8`, `73,749,504` bytes, SHA-256 `64EECEBAF459F23EA11DA38B29C5E290F6F6B22D5EE4749E068BB927A8EEDEBB`.
- A/A protocol is frozen: exactly `10` valid pairs / `20` unprofiled invocations, alternating `L/R` labels, no early stop, no statistical outlier exclusion, continue until 10 valid pairs, with predeclared invalidation reasons and exact raw/statistical/resource records.
- No A/A invocation has started yet.
- Worker must stop after A/A and deliver the numeric proposal to Main for the separate visible Architect decision. It must not ask Owner and must not start U1 before that sealed decision is relayed.

## Cost and impact checkpoint

- Living cost map is attribution context only, not baseline. Profiled disjoint phase wall was resolution `499,895.9298 ms`, DB load `30,635.1092 ms`, parse `14,599.2618 ms`, semantic enrichment `3,928.1245 ms`, cross-file binding `1,627.2942 ms`, other named phases `2,021.7028 ms`, explicit phase residual `14,818.9314 ms`, benchmark total `567,526.3537 ms`, wrapper residual `3,722.9186 ms`, wrapper wall `571,249.2723 ms`.
- Inclusive CPU/allocation rows overlap and must never be summed into wall savings. Primary measured families remain path cleaning/import claims; secondary measured families remain diagnostic normalization/decode.
- Current file impacts are all CRITICAL: `indexes.go 253/38`, `resolve.go 155/48`, `export_resolution.go 163/23`, `diagnostics.go 107/31` (`impacted symbols / affected files`).
- Exact pre-freeze symbol rows remain lineage only and require current rerun before first edit. Important rows include `workspace 145/30/7/71/92`, `buildWorkspace 62/26/7/23/18`, `cleanPath 129/30/5/66/49`, `ResolveBoundInto 120/47/25/31/4`, `AppendDiagnosticToNode 11/7/2/24/6`, and `appendDiagnosticsFromProperties 11/8/2/13/1`.
- HIGH/CRITICAL is a blast-radius warning requiring surgical change and exact affected-boundary regression, not an edit prohibition.
- U1/U2 initial candidate files remain `internal/resolution/indexes.go`, `internal/resolution/resolve.go`, and `internal/resolution/export_resolution.go`. Diagnostics, DB, snapshot, parse, and post-run remain inspect/measure-only until their ordered gates.

## Immediate continuation for incoming Main

1. Complete mandatory FULL RAW startup and detached seal verification before orchestration action.
2. Snapshot the existing worker without interrupting it; preserve its active harness design and bounded read-only discovery.
3. Do not issue another corpus-freeze instruction: corpus and current graph/storage authority are already sealed/current and have not been invalidated by worker-root artifacts because that exact root is excluded. This new Main handoff is outside the excluded root, so before the first A/A invocation the worker must include this new protected row in its already-required protected/corpus identity check; if full-row corpus identity changes, follow the worker's predeclared invalidation/reseal rules rather than reusing old identity blindly.
4. Permit only completion/sealing of the A/A harness/schema and required current impact refresh. Then let the existing worker run exactly 10 valid A/A pairs under the frozen protocol.
5. After A/A, receive the raw summary and numeric proposal; open one separate visible Architect lane with full authority to decide materiality/resource rules. Do not ask Owner.
6. Relay the sealed Architect decision to the existing worker; preserve U1 -> U2 -> rebaseline/ownership -> U3 -> conditional U4 -> Pareto -> final equivalence with one final P6-E commit only.
7. Open one independent P6-E Supervisor only after the complete no-commit candidate and durable Coder handoff.
8. After Supervisor PASS: planner refresh, fresh current E-only detect, exact cleanup disposition, exact isolated staging, one P6-E commit under standing AGENTS Rule 12 authority, and no push.

## Git/worktree snapshot before this handoff

- Snapshot time: `2026-08-24T11:24:15.3184537+07:00`.
- Branch/HEAD/parent/tree: `master` / `599e401548bd69e5214148a9d2ffd515a7e3305d` / `81163e39718b94a509e41114cada224e8f269e36` / `cf35c431962627167f83619c7512ded78705c202`.
- Index: empty, `0` staged paths.
- Status before this handoff: exactly `201` rows = `173` exact worker-root rows + `28` protected outside-worker rows.
- `git diff --check`: PASS, zero output rows.
- This handoff becomes one additional protected outside-worker row. Worker artifact count can increase while active; classify by exact ownership rather than assuming a permanent total.
- Preserve Owner deletion `CONTRIBUTING.md`, Microsoft PowerShell artifact, every historical Main handoff, P6-E plan-acceptance report, all six protected P6-D raw outputs, and every worker artifact.
- Never normalize, stash, reset, checkout, broad-stage, broad-clean, or broad-commit this worktree.

## Existing visible lanes to preserve

- P6-D Coder `01a02f10-c175-7162-ac32-6df30a4d604a`: idle after accepted handoff; preserve.
- P6-D independent Supervisor `01a030d2-3ae2-7b33-a57a-c23503703b1f`: idle after PASS; preserve.
- Independent governance Guard `01a0307e-fe39-73c2-83df-e2bb234143a5`: preserve independently.
- Active P6-E Coder `01a0319a-62c9-7650-ba6a-ef59dfcf1973`: preserve and monitor.

## Governance findings and immutable corrections

- Current successor warm creation was late exactly `0.240s — HIGH`; preserve exactly and do not create a duplicate successor or worker.
- Prior successor warm lateness `0.154s — HIGH` remains immutable history.
- Preserve all historical default/shared-registry, source-grep, Rule-12, final-yield, commit-delay, stage/commit-after-compaction, warm/hard lateness, and correction findings from earlier sealed handoffs.
- The withdrawn claim that a fresh graph invocation did not exist must never be resurrected.
- Every `E:\Anvien\.tmp` path is disposable scratch, deletable wholesale by Owner, and never evidence, review/handoff/recovery authority, source of truth, acceptance dependency, or procedure blocker.

## Workspace bans and standing authority

- Sole workspace is `E:\Anvien`, local saved project, no worktree.
- No C-drive workspace/read/write, alternate worktree, target access, implicit network, install/package-script side effects, or shared/global/protected/quarantined cache authority.
- Preserve vendor, third_party, durable reports, Owner work, Owner deletion, and unrelated/protected rows.
- Owner `PAUSE` or `STOP` halts commands, edits, QA, commits, and lane control immediately.
- Commit after completed accepted P6-E is standing AGENTS Rule 12 authority; never ask Owner to grant it again.
- Do not push unless Owner explicitly requests push.
- Incoming Main must initialize its next visible successor at the exact raw-UUID warm target and transfer no later than its hard deadline unless Owner gives pause/override.

Outgoing Main terminates absolutely after official transfer dispatch.
