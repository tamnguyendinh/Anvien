# Visible Main Orchestration Rotation Handoff

## Chuyển giao authority

- Campaign: Anvien Graph Accuracy / Analyze Performance.
- Outgoing Visible Main: `01a03162-ccd2-76a1-910b-88ebdb8fec9c`.
- Incoming visible successor: `01a0319b-58be-7150-a84b-1fecb90aa7da`.
- Raw outgoing official authority user-message UUIDv7: `01a03171-b1c1-74e0-8bb9-8f87ee889057`.
- Transport-source authority start: `2026-08-24T08:45:29.793+07:00`.
- Nominal warm successor target: `2026-08-24T09:30:29.793+07:00`.
- Hard transfer deadline: `2026-08-24T09:45:29.793+07:00`.
- Successor creation UUIDv7: `01a0319b-58be-7150-a84b-1fecb90aa7da`.
- Successor creation: `2026-08-24T09:30:59.518+07:00`.
- Exact warm lateness: `29.725s` — verified `HIGH`; preserve exactly, do not normalize or downgrade.
- Successor returned `UNDERSTOOD / WAITING_FOR_OFFICIAL_TRANSFER`, confirmed all boundaries/bans, and is idle without tools/actions.
- Exact incoming authority time is the timestamp encoded by the raw UUIDv7 of the follow-up with exact heading `OFFICIAL AUTHORITY TRANSFER`; do not backfill it into this immutable report.
- Outgoing Main terminates absolutely immediately after official dispatch.

## Mandatory iron startup for successor

Before any orchestration action, read each file FULL RAW sequentially through EOF:

1. `E:\Anvien\AGENTS.md`.
2. `E:\Anvien\.agents\skills\working-rules\SKILL.md`.
3. `E:\Anvien\.agents\skills\orchestration\SKILL.md`.
4. This entire sealed handoff, then verify its detached identity from the official transfer message.

Before any plan/ledger action, continue FULL RAW through EOF:

1. `E:\Anvien\.agents\skills\planner\SKILL.md`.
2. `E:\Anvien\docs\plans\2026-07-26-anvien-graph-accuracy-multi-plan\2026-07-26-anvien-graph-accuracy-roadmap.md`.
3. All four Child 06 ledgers under `E:\Anvien\docs\plans\2026-07-26-anvien-graph-accuracy-multi-plan\2026-07-28-06-ambient-external-resolution-and-diagnostics\`.
4. The current P6-E reports and checkpoint identified below.

Do not use summaries, compacted context, prior-session memory, or file names instead of raw source. After any auto-compaction, perform the orchestration section 7 raw re-anchor before action. `governance-rule-guard` is the existing independent Guard; Main must not read or use that skill and must not act as Guard.

## Current plan state

- P6-D is `[x]`, `CLOSED`, and must not be reopened or rechecked absent proved invalidation.
- P6-D implementation commit: `81163e39718b94a509e41114cada224e8f269e36`.
  - Parent: `741335f7efa5ad215ec28361d88c462f431565ab`.
  - Tree: `b6bc216172886d5de9ef80f82c0da3ed28880f13`.
  - Exact `1,883` paths.
- Five-document P6-D closure/P6-E opening commit and current HEAD: `599e401548bd69e5214148a9d2ffd515a7e3305d`.
  - Parent: `81163e39718b94a509e41114cada224e8f269e36`.
  - Tree: `cf35c431962627167f83619c7512ded78705c202`.
  - Exact boundary: roadmap plus four Child 06 ledgers.
- P6-E is the sole open Child 06 slice.
- P6-E is exactly one checklist slice and one isolated commit boundary containing ordered internal units `U1 -> U2 -> rebaseline/ownership -> U3 -> conditional U4 -> fresh Pareto -> final equivalence`.
- U1-U4 are not separate slices, acceptance gates, or commits. There is no intermediate U2/U3 commit.
- Pn-A and later gates remain locked.
- Do not update the P6-D checklist again.

## Active visible P6-E implementation/performance lane

- Task: `01a0319a-62c9-7650-ba6a-ef59dfcf1973`.
- Saved project: `E:\Anvien`.
- Environment: `local`, no worktree.
- State at handoff capture: `ACKNOWLEDGED / RUNNING MANDATORY RAW STARTUP`; no source/plan edit, stage, commit, push, target access, or implementation action has occurred.
- Mandatory first response was correct: `UNDERSTOOD`, exact goal/boundary/first action stated.
- Latest observed activity: FULL RAW rules/skills and roadmap/Child 06 authority reading. The lane has not crossed into cost-map or source work yet.
- Lane authority: source inspection, durable full-row graph evidence, exact run identity, granular cost map, A/A protocol, alternating unprofiled baseline, U1-U4 implementation under ordered gates, A/B/equivalence/resource/rollback proof, durable Coder handoff.
- Lane prohibition: it cannot update roadmap/four ledgers, stage, commit, push, access `E:\cheapapp.org`, use C-drive workspace/read/write, use default/shared/global registry/cache, or create alternate worktrees.
- Evidence must be born directly under an exact worker-owned durable `E:\Anvien\reports\coder\...` root. `.tmp` is scratch only and cannot be evidence or acceptance authority.
- The lane must stop after A/A with a concise proposed numeric materiality/resource rule and wait for Main/Owner approval before U1. It must also stop for any boundary expansion requiring planner refresh.
- Final lane result is `READY_FOR_MAIN` or a precise `BLOCKED`/rejected invariant with index empty and no commit/push.
- Next owner after complete candidate is Visible Main, which opens a new independent P6-E Supervisor only then.
- Preserve and continuously monitor this exact lane. Do not duplicate, replace, pressure, stop, archive, or transfer it merely because Main rotates.

## Fresh E-only graph and direct runtime authority

- Direct runtime: `E:\Anvien\anvien\bin\anvien.exe`.
- Version: `1.2.8`.
- Size: `73,749,504` bytes.
- SHA-256: `64EECEBAF459F23EA11DA38B29C5E290F6F6B22D5EE4749E068BB927A8EEDEBB`.
- Real fresh E-only `anvien analyze --force` exited `0` at current HEAD `599e401548bd69e5214148a9d2ffd515a7e3305d`.
- Analyze result: `2,171` scanned / `765` parsed / `0` failed.
- Graph: `122,665` nodes / `169,316` relationships.
- File projection: `20,351` dependency edges / `479` unresolved.
- E-only `anvien status` after the graph run reports indexed `599e401`, current `599e401`, `up-to-date`.
- The first attempted `anvien status --json` returned only `unknown flag: --json`; it is not evidence and did not alter graph state.
- Every Anvien invocation must pin process-local `ANVIEN_HOME`, `HOME`, `USERPROFILE`, `APPDATA`, `LOCALAPPDATA`, `TEMP/TMP`, XDG config/cache/data, and process-local Git `safe.directory` to an E-only lane root. Do not touch default/global registry again.
- The outgoing operational pin existed below `E:\Anvien\.tmp\p6e-main-fresh-graph-01a03162`; because every `.tmp` path is disposable scratch, this path is not handoff authority or source of truth. Successor/worker must verify or recreate its own E-only process-local pin from durable workspace/runtime facts.

## Valid current P6-E file-detail preflight

All rows below come from the current E-only graph at `599e401`; the stale default-registry result identified later is excluded absolutely.

| Owner | File risk | Related | Symbols | Inbound | Outbound | Local | Unresolved | Linked flows | Linked tests |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| `internal/resolution/indexes.go` | HIGH | current full payload retained in command record | 344 | 339 | 108 | 388 | 453 | 23 | 38 |
| `internal/resolution/resolve.go` | HIGH | 77 | 200 | 134 | 289 | 87 | 403 | 26 | 40 |
| `internal/resolution/export_resolution.go` | HIGH | 42 | 186 | 59 | 102 | 190 | 267 | 18 | 24 |
| `internal/graphhealth/diagnostics.go` | HIGH | 40 | 154 | 50 | 70 | 149 | 187 | 1 | 19 |

The worker must retain formal full rows durably before the first edit. The Main summaries do not replace full evidence.

## Valid current upstream file impact preflight

All four initial owner files are `CRITICAL`. These are blast-radius warnings, not edit bans.

| Owner | Risk | Impacted | Contained symbols | Direct | Affected files | Affected processes | Linked flows | Linked tests |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| `internal/resolution/indexes.go` | CRITICAL | 253 | 344 | 88 | 38 | 1 | 23 | 38 |
| `internal/resolution/resolve.go` | CRITICAL | 41 | 200 | 30 | 6 | 1 | 26 | 40 |
| `internal/resolution/export_resolution.go` | CRITICAL | 128 | 186 | 13 | 10 | 1 | 18 | 24 |
| `internal/graphhealth/diagnostics.go` | CRITICAL | 61 | 154 | 24 | 16 | 1 | 1 | 19 |

The initial three long file-impact commands were accidentally launched once without exposing their exec session IDs to Main; they completed read-only. Main reran them on the same E-only graph solely to capture exact summaries. No state changed; only the captured reruns above are used.

## Valid current upstream symbol impact preflight

| Symbol request | Risk | Impacted | Files | Modules | Processes | Direct | Disposition |
|---|---:|---:|---:|---:|---:|---:|---|
| `buildWorkspace` | CRITICAL | 8 | 6 | 5 | 23 | 2 | exact function matched in `indexes.go` |
| `explicitImportNameClaimed` | LOW | 0 | 1 | 0 | 0 | 0 | exact method matched in `resolve.go` |
| `typeScriptAnnotationImportClaimed` | LOW | 0 | 1 | 0 | 0 | 0 | exact method matched in `resolve.go` |
| `explicitImportCallState` | LOW | 0 | 1 | 0 | 0 | 0 | exact method matched in `export_resolution.go` |
| `AppendDiagnosticToNode` | CRITICAL | 7 | 3 | 1 | 24 | 2 | exact function matched in `diagnostics.go` |
| `appendDiagnosticsFromProperties` | CRITICAL | 7 | 4 | 2 | 13 | 1 | exact function matched in `diagnostics.go` |
| `normalizeDiagnosticMetadata` | CRITICAL | 13 | 6 | 2 | 28 | 4 | exact function matched in `diagnostics.go` |
| `decodeStructuredResolutionOutcome` | LOW | 6 | 1 | 1 | 0 | 1 | exact function matched in `diagnostics.go` |
| `workspace.resolveImports` | UNKNOWN | 0 | 1 | N/A | 1 | N/A | no exact graph target; UNKNOWN is not no-impact |
| `cleanPath` | UNKNOWN | 0 | 1 | N/A | 1 | N/A | no exact graph target; UNKNOWN is not no-impact |
| `repositoryReceiverClaimed` | UNKNOWN | 0 | 1 | N/A | 1 | N/A | no exact graph target; UNKNOWN is not no-impact |

The worker must resolve the three UNKNOWN rows through full Anvien file-detail/context and rerun qualified upstream impact. Keyword/source grep is forbidden for that resolution.

## P6-E immediate continuation

1. Monitor the existing visible worker's actual commands, files, durable roots, scope, and loops continuously.
2. Verify that it completes FULL RAW startup and then freezes exact E-only repo/dirty/corpus/runtime/native/dependency/cache/instrumentation identity.
3. Require durable full-row file-detail/symbol/file impact and explicit CRITICAL blast-radius reporting before edit.
4. Require a living granular cost map with disjoint wall plus residual, CPU/allocation/memory/I/O/wait/work units/calls/bytes/histograms; profiled attribution and unprofiled wall remain separate.
5. Require predeclared A/A and alternating unprofiled baseline repetitions/stopping/cache/median/dispersion/CI.
6. At the A/A decision boundary, inspect the worker's proposed numeric materiality/resource rule, update only actual current P6-E ledger state with planner when warranted, and obtain the required Main/Owner approval before U1.
7. Execute `U1 -> U2 -> rebaseline/ownership -> U3 -> conditional U4 -> Pareto -> final equivalence`; no intermediate U2/U3 commit.
8. Any later DB/snapshot/parse/post-run owner requires a four-ledger planner refresh naming exact current cost, owner, architecture contract, benchmark row, failure/publication/resource rules, and rollback condition before edit.
9. Open one new independent P6-E Supervisor only after a complete no-commit candidate and durable Coder handoff.
10. After Supervisor PASS: planner refresh, fresh current E-only detect, exact cleanup disposition, exact isolated staging, one P6-E commit under standing AGENTS Rule 12 authority, and no push.

## Current Git/worktree reality at pre-handoff capture

- Sole workspace: `E:\Anvien`.
- Branch: `master`.
- HEAD: `599e401548bd69e5214148a9d2ffd515a7e3305d`.
- Parent: `81163e39718b94a509e41114cada224e8f269e36`.
- Tree: `cf35c431962627167f83619c7512ded78705c202`.
- Index: empty, `0` staged paths.
- Status before this new handoff: exactly `26` protected rows.
- `git diff --check`: PASS.
- This new handoff is uncommitted and becomes the expected `27th` protected row if the worker has not yet created its authorized durable artifacts. After worker activity begins, classify additional rows against the exact worker-owned durable manifest rather than assuming a fixed total.
- Preserve Owner deletion `CONTRIBUTING.md`, the Microsoft PowerShell artifact, all historical Main handoffs, the P6-E plan-acceptance report, and the six protected P6-D raw outputs.
- Never normalize, stash, reset, checkout, broad-stage, broad-clean, or broad-commit this worktree. Preserve `vendor/`, `third_party/`, durable reports, Owner work, and unrelated/protected rows.

## Existing visible lanes to preserve

- P6-D Coder: `01a02f10-c175-7162-ac32-6df30a4d604a`, idle after accepted handoff. Do not stop, pressure, duplicate, replace, or archive.
- P6-D independent Supervisor: `01a030d2-3ae2-7b33-a57a-c23503703b1f`, idle after PASS. Do not stop, pressure, duplicate, replace, or archive.
- Independent governance Guard: `01a0307e-fe39-73c2-83df-e2bb234143a5`. Preserve; do not stop, replace, archive, read/use its skill, or self-assume its role.
- Active P6-E Coder: `01a0319a-62c9-7650-ba6a-ef59dfcf1973`, described above.

## New verified governance findings and corrections

1. `HIGH` — default/shared registry boundary violation.
   - One `file-detail` process did not reapply E-only HOME/registry, touched the default/shared registry, and returned stale anchor `741335f`.
   - That entire result is invalid absolutely and must never be cited or used for lane design.
   - Correction: every later Anvien invocation was pinned to the E-only process environment; status verified indexed/current `599e401`, and valid file-detail/impact was rerun there.
2. `HIGH` — Main source-grep/worker-discovery violation.
   - Main ran two source-level `rg` investigations for runtime/env mechanics after re-anchor.
   - All keyword hits and any conclusions from them are invalid and excluded. Main stopped source discovery, returned to direct verified runtime/E-only status, finished only Main-owned graph preflight, and delegated all source/cost/implementation work to the visible worker.
3. `HIGH` — successor warm creation late `29.725s`.
   - Required warm target: `2026-08-24T09:30:29.793+07:00`.
   - Raw successor UUIDv7 creation: `2026-08-24T09:30:59.518+07:00`.
   - Preserve exact lateness and severity. No duplicate successor or worker was created.

Keep every historical finding/correction from the prior sealed handoff immutable, including the Rule-12/final-yield/commit-delay HIGH findings, stage/commit-after-compaction HIGH, historical warm/hard lateness findings, and their completed corrections. The former claim that a fresh graph invocation did not exist was formally withdrawn after the real exit-0 invocation appeared; never resurrect it.

## Binding `.tmp` correction

- Every path under literal `E:\Anvien\.tmp` is disposable reproducible scratch and may be deleted wholesale by Owner at any time.
- `.tmp` is never evidence, review input, handoff/recovery authority, source of truth, acceptance dependency, or procedure blocker.
- Do not recreate false per-root ownership or broad-cleanup findings.
- A lane still names exact roots to avoid collisions, but loss/drift of `.tmp` cannot invalidate durable evidence or become a blocker by itself.

## Workspace, authority, and stop conditions

- Sole repository workspace: `E:\Anvien`, saved project local, no worktree.
- No C-drive workspace/read/write, alternate worktree, implicit network, install/package-script side effects, or shared/global/protected/quarantined cache authority.
- Stage/isolated commit after the completed accepted P6-E slice is standing campaign authority under AGENTS Rule 12. Never ask Owner to grant it again.
- Do not push unless Owner explicitly says push.
- Owner `PAUSE` or `STOP` halts every command/edit/QA/commit/lane-control action immediately.
- Incoming Main must initialize its next visible successor at the exact new raw-UUID warm target and transfer no later than its hard deadline unless Owner gives a new pause/override.

Outgoing Main terminates absolutely after official transfer dispatch.
