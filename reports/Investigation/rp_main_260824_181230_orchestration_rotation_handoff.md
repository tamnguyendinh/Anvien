# Visible Main Orchestration Rotation Handoff

## Authority transfer record

- Campaign: Anvien Graph Accuracy / Analyze Performance.
- Outgoing Visible Main: `01a03345-22a5-7033-b25a-4e5b64a39d35`.
- Prepared clean successor: `01a03375-ef3c-73c1-972a-5532b262558d` on host `local`.
- Raw incoming official-authority UUIDv7 for the outgoing Main: `01a03349-f07b-74b2-9a7f-7253c8044d2e`.
- Outgoing authority start: `2026-08-24T17:21:18.843+07:00`.
- Outgoing warm target: `2026-08-24T18:06:18.843+07:00`.
- Outgoing hard transfer deadline: `2026-08-24T18:21:18.843+07:00`.
- First post-compact clock observation: `2026-08-24T18:06:55.469+07:00`.
- Exact warm lateness: `36.626s — HIGH — VERIFIED VIOLATION`.
- Preserve this value exactly. Do not normalize, omit, downgrade, justify, or replace it with a later tool time.
- The raw UUIDv7 of the follow-up headed exactly `OFFICIAL AUTHORITY TRANSFER` is the successor's incoming authority-time source of truth. The successor computes its warm target at `+45 minutes` and hard deadline at `+60 minutes` from that UUIDv7.
- Outgoing Main terminates absolutely after official dispatch.
- Never use `fork_thread` for Main rotation.

## Successor preparation and ACK

- Successor thread: `01a03375-ef3c-73c1-972a-5532b262558d`.
- Pre-transfer prompt prohibited every tool/read/lane action before official transfer.
- Successor ACK turn: `01a03375-f113-7c92-9b87-31af1ce332ea`.
- ACK: `UNDERSTOOD / WAITING_FOR_OFFICIAL_TRANSFER`.
- Successor restated the exact campaign goal, sole-open P6-E order, Main-only command authority, workspace/bans, old Worker STOP, replacement Coder preservation, Guard preservation, invalid-fork exclusion, and no-action boundary.
- No successor tool action occurred before this official transfer.

## Mandatory raw startup and role boundary

Before any orchestration action, successor must read FULL RAW sequentially through EOF:

1. `E:\Anvien\AGENTS.md`.
2. `E:\Anvien\.agents\skills\working-rules\SKILL.md`.
3. `E:\Anvien\.agents\skills\orchestration\SKILL.md`.
4. `E:\Anvien\.agents\skills\governance-rule-guard\SKILL.md`.
5. This entire handoff and verify the detached seal supplied with official transfer.

Before plan/ledger update or acceptance, read FULL RAW the applicable planner/supervisor skill, active roadmap, all four Child 06 ledgers, and the exact current lane handoff/evidence. After compaction, follow orchestration section 7 and continue from the first uncompleted gate. Re-anchor must not restart passed gates or incident audits.

Apply the FULL RAW, unmodified `AGENTS.md`, `working-rules`, and `orchestration` rules at all times. No summary, compact memory, handoff synopsis, filename, or prior verdict substitutes for those raw rules.

Main is Main only: design lanes, issue exact commands, monitor actual commands/writes, block deviation, receive durable handoffs, verify boundaries, route independent review, and decide transitions. Main must not do Coder, Worker, Supervisor, QA, Architect, or Planner work.

## Binding Owner correction

- Owner explicitly corrected that Main's job is to drive the unfinished campaign to completion, not obey Worker/lane reports or sit in passive waiting.
- Owner explicitly ordered that old Worker `01a0319a-62c9-7650-ba6a-ef59dfcf1973` must not work anymore and that a different visible Coder be opened.
- Main must issue commands to lanes. Lanes execute within their assignment and hand evidence back; lanes do not select Main's workflow or transition.
- Do not ask Owner for intermediate choices already determined by repository authority.
- Do not repeat audit/report/recovery loops on unchanged evidence.

Preserve these governance findings exactly:

- predecessor warm lateness `52.457s — HIGH — VERIFIED VIOLATION`;
- predecessor successor-creation lateness `3.488s — HIGH — VERIFIED VIOLATION`;
- this Main warm lateness `36.626s — HIGH — VERIFIED VIOLATION`;
- this Main's interrupted monitoring turn before warm: `HIGH — VERIFIED VIOLATION`;
- this Main's passive wait-only posture after the incident verdict: `HIGH — VERIFIED VIOLATION`.

No justification or normalization is authorized.

## Current campaign reality

- Sole workspace: `E:\Anvien`, saved project `local-dc76d679b5f0e3c8ddbb8f9218f46f6d`, environment `local`, no worktree.
- P6-D remains CLOSED at accepted isolated commit `81163e39718b94a509e41114cada224e8f269e36`; do not reopen without actual invalidation.
- P6-E is the sole open Child 06 slice and exactly one eventual commit boundary:
  `U1 -> U2 -> rebaseline/ownership -> U3 -> conditional U4 -> Pareto -> final equivalence`.
- P6-E goal: improve end-to-end `anviens analyze` runtime while preserving accuracy, semantic completeness, graph correctness, deterministic output, freshness, failure/publication behavior, and persistence/reader parity.
- Current product state at transfer: A/A `0/20`; harness unsealed; corpus/graph performance run not started; no comparable baseline; U1 locked; no P6-E implementation, Supervisor acceptance, detect, cleanup, commit, push, or target access.
- Exact first active gate now delegated: takeover plus fresh pre-edit graph/file-detail/impact authority `E6-P6E-IMPACT1` only.
- After a valid `E6-P6E-IMPACT1` handoff, Main decides whether and how to open living cost-map plus a clean harness/A/A protocol. Coder may not self-transition.
- After exactly `10` valid A/A pairs, Coder must stop before U1. Main then opens one separate visible Architect lane to seal numeric materiality/resource rules, relays that seal, and only then may U1 open.
- P6-E has exactly one final commit boundary. No interim checkpoint commit is authorized despite generic Coder defaults.

## Replacement Coder lane

- New visible Coder: `01a03375-f616-7651-ba81-99bad6536746`, host `local`.
- Pre-assignment ACK turn: `01a03375-fa2b-7ae3-a37b-cf83c7f5c5e8`.
- ACK: `UNDERSTOOD / WAITING_FOR_OFFICIAL_CODER_ASSIGNMENT`.
- Official assignment turn was dispatched after ACK.
- New sole durable root: `E:\Anvien\reports\coder\p6e_takeover_260824_181055_raw`.
- At the pre-handoff snapshot, the Coder is active and is reading FULL RAW `AGENTS.md`, `working-rules`, and `coder`; it has not yet written its durable root.
- Coder assignment is bounded to:
  1. raw startup and exact P6-E authority reads;
  2. read-only Git/worktree preservation inventory;
  3. direct E-only runtime help;
  4. direct E-only `analyze --force` only if E-only cache/runtime boundaries are proved before launch;
  5. fresh file-detail and upstream impact for `internal/resolution/indexes.go`, `internal/resolution/resolve.go`, `internal/resolution/export_resolution.go`, plus exact U1/U2 symbols/callers proved by file-detail;
  6. one durable `E6-P6E-IMPACT1` report/checkpoint and STOP.
- Coder is forbidden to run A/A, open U1, edit production/tests, build, commit, review, delete/clean, or self-transition in this assignment.
- Monitor actual commands and writes. Intervene immediately on scope drift, forbidden cache/path access, old-root use, production edit, or self-transition.

## Old Worker termination and provenance boundary

- Old Worker: `01a0319a-62c9-7650-ba6a-ef59dfcf1973`.
- It received absolute STOP and returned `STOPPED`.
- It reported no active process; last command was a read-only `rg`; last write was its prior canonical handoff.
- It must never resume or receive work again under current Owner authority.
- Old durable root remains read-only provenance and excluded from the replacement assignment:
  `E:\Anvien\reports\coder\p6e_analyze_performance_260824_093855_raw`.
- Canonical old report remains protected at
  `handoff\rp_coder_260824_170000_p6e_canonical_worker_handoff.md`,
  identity `23,298` bytes / `290` LF / `0` CR / SHA-256 `AB22D39B3280BAB3038099E2654F9973E42D490366E26F993B8DAAC878E28721`.
- Do not write, reuse harness/scripts, resume, message, archive, or transfer the old Worker/root.

## Protected-row incident boundary

- Historical deleted path: `E:\Anvien\Microsoft\Windows\PowerShell\ModuleAnalysisCache`.
- Last observed identity: regular file / `4,641` bytes / SHA-256 `D3452BF999F2C34D83B183D6FAC2F73FC5A1225BC0755559B8D03BAC6FAB06FE`.
- Exact writer remains `UNKNOWN_NOT_CAPTURED`.
- Sole incident Supervisor report remains terminal historical evidence:
  `E:\Anvien\reports\Supervisor\rp_supervisor_260824_171447_by_gpt-5_p6e_protected_row_deletion_incident.md`,
  SHA-256 `C8252B768A94C4B5590D41D71F5EAA8FDB6DE25DC560B3AE9311904231964FAA`, verdict `REJECT` under the prior continuation route.
- Latest Owner authority replaces the old Worker and commands campaign advancement; it does not authorize touching that path.
- Replacement Coder assignment explicitly excludes the path: no read/hash polling, restore, recreate, regenerate, write, or baseline adoption action.
- Do not open another unchanged-byte incident inventory, recovery audit, Supervisor, or wording review. The functional next action is the separately owned fresh impact gate, not an incident re-review.

## Independent runtime identity boundary

- Runtime identity report:
  `E:\Anvien\reports\Supervisor\rp_supervisor_260824_154625_by_gpt-5_p6e_runtime_identity_independent_review.md`.
- SHA-256: `467D9C75E2BA73FD70D8A1CC42A140400859AE751384790C464F8FE6B9A655A1`.
- Identity-only PASS covers `E:\Anvien\anvien\bin\anvien.exe`, version `1.2.8`, `73,750,016` bytes, SHA-256 `D4EC8B58C41B9F0A95359CFC014DB07D11653F25736728549555F74D204ABD19`.
- It accepts only runtime identity, not harness/corpus/graph/A-A/baseline/U1/P6-E.

## Governance lineage

- Preserve continuous Guard lineage; latest warning transport observed from `01a03373-96ae-70d2-94c4-ac5e631a2a72`, with prior transport `01a03339-144d-7383-8cd1-72fcd5e5417c`.
- Guard is governance-only. Do not duplicate it or control it as a functional lane.
- Invalid fork `01a032d9-44a3-7fe2-b97b-8a8fe89fe8f3` has no authority and must never be used.

## Pre-handoff Git/worktree snapshot

- Observed: `2026-08-24T18:12:18.160+07:00`.
- Branch: `master`.
- HEAD: `1a9b5310bc3204c610e9e6c673afc214cd049991`.
- Parent: `d02e4eb76ee6e42848ea0ffb627449ac7d9df092`.
- Tree: `f124bc5811ae26b2751c1b09fc8da673b9f4ab1c`.
- Status: `237` rows before this new handoff.
- Old Worker-root rows: `196`.
- New Coder-root rows: `0`.
- Index: empty, `0` staged paths.
- `git diff --check`: PASS.
- This handoff adds exactly one protected outside-Worker row. Detached verification after creation must confirm the only expected delta is this handoff unless the active replacement Coder independently creates its authorized new-root checkpoint meanwhile.

## Workspace bans and standing authority

- Sole workspace is `E:\Anvien`, local, no worktree.
- No C-drive workspace/read/write, alternate worktree, target access, network, install/package scripts, shared/global/protected/quarantined cache, reset/stash/checkout/broad cleanup, unscoped discovery, broad staging, or push.
- Every `E:\Anvien\.tmp` path is disposable scratch only, never evidence, authority, blocker, recovery source, or source of truth.
- Preserve Owner dirty work/deletion, all ledgers, vendor, third_party, durable reports, handoffs, old Worker artifacts, replacement Coder artifacts, Supervisor reports, and protected rows.
- Commit only after completed, independently accepted P6-E with final detect/cleanup under standing AGENTS Rule 12. No push unless Owner explicitly requests it.
- Owner PAUSE/STOP halts commands, edits, QA, commit, and lane control immediately.

## Immediate successor actions

1. Complete FULL RAW startup and verify this handoff's detached seal.
2. Record exact incoming authority time from the official transfer message UUIDv7; calculate new warm and hard targets.
3. Preserve Guard continuity and all verified violations without justification.
4. Monitor replacement Coder `01a03375-f616-7651-ba81-99bad6536746` actual commands/writes. Do not merely wait for its prose report.
5. Enforce its exact `E6-P6E-IMPACT1` boundary and STOP conditions.
6. On handoff, Main verifies report/source/diff/worktree/evidence, then decides and commands the next P6-E gate. Coder does not decide the transition.
7. Keep campaign moving: after accepted impact authority, issue a bounded cost-map/harness/A-A assignment; after `10` valid A/A pairs, route one Architect decision; then command U1 and subsequent ordered gates.
8. Initialize the next clean visible Main successor at the exact new warm target and transfer no later than the new hard deadline.

Outgoing Main terminates absolutely after official transfer dispatch.
