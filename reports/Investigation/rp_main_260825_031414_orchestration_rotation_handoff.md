# Main Orchestration Rotation Handoff — 2026-08-25 03:14:14 +07:00

## Authority

- Outgoing Main: `01a034d7-07f8-7310-b224-48a90ba8c791`.
- Clean visible successor: `01a03569-775d-7cb2-b302-c44b4f2802fe` on host `local`.
- Successor authority start: `2026-08-25T03:14:14.223+07:00`.
- Warm target: `2026-08-25T03:59:14.223+07:00`.
- Hard deadline: `2026-08-25T04:14:14.223+07:00`.
- Exact official transfer message sent after this file is sealed is the authority source of truth.

## Mandatory Raw Startup

Read each source FULL RAW sequentially through EOF; summaries, this handoff, memory, and prior context are not substitutes:

1. `E:\Anvien\AGENTS.md`
2. `E:\Anvien\.agents\skills\working-rules\SKILL.md`
3. `E:\Anvien\.agents\skills\orchestration\SKILL.md`
4. `E:\Anvien\.agents\skills\governance-rule-guard\SKILL.md`
5. `E:\Anvien\docs\plans\2026-07-26-anvien-graph-accuracy-multi-plan\2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy\2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy-plan.md`
6. `E:\Anvien\docs\plans\2026-07-26-anvien-graph-accuracy-multi-plan\2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy\2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy-evidence.md`
7. `E:\Anvien\docs\plans\2026-07-26-anvien-graph-accuracy-multi-plan\2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy\2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy-benchmark.md`
8. `E:\Anvien\docs\plans\2026-07-26-anvien-graph-accuracy-multi-plan\2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy\2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy-actual-status.md`
9. `E:\Anvien\docs\plans\2026-07-26-anvien-graph-accuracy-multi-plan\2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy\plan-rules.md`
10. This sealed handoff.

Apply the raw rules in full. Do not restart accepted capture, passed documentation review, graph-refresh history, completed impacts, or the P6-D doc update.

## Exact Campaign State

- Goal: complete Child 06A using exhaustive measured-bottleneck optimization of real `anvien analyze`, preserving final equivalence; exactly one P2-A implementation slice, one final Supervisor, cleanup/detect, exactly one implementation commit, then handoff Child 07.
- First incomplete gate remains Child 06A `P1-A`, with open `E1-P1A-INSTR1`. No complete `OPnnn` list/matching parent checklist, no P2 attempt, no KEEP/speedup, and no P3-C implementation commit.
- Accepted baseline remains process wall `605.732722s`, analyzer internal `602.5278811s`, and 15-phase sum `587.2385976s`.
- Production instrumentation remains in `internal/analyze/analyze.go`, `internal/cli/command.go`, and `internal/cli/command_test.go`. Inspect current source/diff and writer transcript once; do not duplicate the writer.
- Latest one-shot lane snapshot: instrumentation writer thread status is `notLoaded` with its latest turn completed; treat it as idle/preserved until exact transcript/diff inspection. Measurement executor is idle/preserved. Production Coder is explicitly `STOPPED` and must remain HOLD until P2 opens. Planner is preserved. Guard lineage is preserved and must be retargeted, not duplicated or controlled as a functional lane.

## Owner-Corrected Build Contract

- Binding future flow: `source -> E:\Anvien\anvien\bin\anvien.exe -> launcher -> analyze`.
- `scripts/full-build.ps1` stops exact output holders, runs `go run -mod=vendor ..\cmd\anvien package build-runtime`, invokes launcher build with `pwsh`, prints version/Git HEAD/SHA-256, then runs the canonical binary `analyze . --force`.
- `internal/cli/package_runtime.go` builds directly to the canonical binary, preserves checked-in vendor, durable LadybugDB native authority, offline guards, and `anvien\bin\anvien-runtime.json`.
- `anvien-launcher/build.ps1` builds launcher/server/web directly into canonical outputs and does not compile the canonical CLI again.
- Use default Go/TypeScript/Vite caches. No per-run lane/cache/runtime/staging/provenance, custom HOME/TEMP, or `GOCACHE`/`GOMODCACHE`/`GOPATH`/`GOTMPDIR`/`NPM_CONFIG_CACHE` override. `anvien clean` is graph lifecycle only.
- Canonical runtime build after this correction PASSed in `25.6760326s`, version `1.2.8`, SHA-256 `9645600C7CA9436EA31B3D0C8C24210B4CFA5A22A1DD02A6DA87AAAF9C148EED`.
- PowerShell syntax PASS, scoped `git diff --check` PASS, and `go test ./internal/cli -run '^TestP6D' -count=1` PASS in `157.458s`.
- A full build/analyze has not run after the simplification. Therefore current source requires a new canonical full build before instrumentation capture evidence may be accepted; after current source changes, do not use the older graph for new graph-based decisions.

## Git and Worktree

- Sole workspace: `E:\Anvien`, saved project local, no worktree.
- Current HEAD: doc-only commit `50cbeb24` (`docs(plan): supersede P6-D build machinery`). It contains only the isolated P6-D future-build plan correction and does not count as the one P3-C implementation commit.
- Earlier doc-only checkpoint `7b3bccb3b1abd09d1d2d3bdc9d864f55d9bd4965` remains provenance and also does not count as P3-C.
- Owner-authorized build simplification is uncommitted in `scripts/full-build.ps1`, `anvien-launcher/build.ps1`, `internal/cli/package_runtime.go`, and `internal/cli/package_runtime_p6d_test.go`. Preserve it; do not create an extra implementation commit without reconciling the single P3-C boundary.
- Instrumentation changes are also uncommitted. Many plan/report/user changes remain dirty/protected. Preserve them. No reset/stash/checkout/broad cleanup/broad stage/push.
- The Child 06 plan still has pre-existing protected unstaged changes after the isolated hunk commit; this is expected.

## First Required Continuation

1. Complete Mandatory Raw Startup and verify this detached seal once.
2. Retarget continuous Guard to successor ID; do not duplicate Guard.
3. Inspect exact current task/process/diff once. Do not reopen an audit loop.
4. Continue P1-A only: establish current full-build readiness, prove/clear exact output holders, run the canonical full build on current source, then open exactly one like-for-like instrumented capture through the preserved measurement executor when the gate is ready.
5. Capture must use real canonical `anvien analyze`, exact accepted workload/options including `--benchmark-json`, and produce the complete non-overlapping measured operation inventory, denominators, descending `OPnnn` list, and matching parent checklist before P2.
6. Make exact carry-or-remove disposition for instrumentation after capture. No P2 optimization before P1-A acceptance.

Do not conflate the maintenance analyze inside full build with the like-for-like accepted instrumented capture if it lacks the exact benchmark/options tuple.

## Binding Workflow

- Main has complete command authority below User and must not ask User to authorize continuation.
- Elapsed wall-clock authority; complete parent/child lists and matching checklists; maximum three ACTIVE read-only analysts.
- Each production attempt: Architect -> Planner -> Coder -> build/remeasure child+parent+E2E -> Supervisor.
- KEEP requires child elapsed lower, parent end-to-end elapsed lower, E2E elapsed lower, and equivalence/accuracy PASS. Three consecutive no-KEEP terminalizes only the child.
- Current Child 06 has no current Pn-A/B/C; legacy E6-PNA/PNB/PNC are provenance only. Child 06A P3-A/B/C is closure.
- No hidden functional agents and never `fork_thread`.

## Workspace Bans

No C-drive workspace/read/write, alternate worktree/target, network, install/package scripts, protected/shared/global/quarantined cache, reset/stash/checkout/broad cleanup, broad stage, or push. Repo-local `.tmp` only for scratch. Kill only exact verified build-output holders; if an observed process is legitimately completing current work, wait, otherwise terminate it before build.

## Preserved Violations

Do not normalize, downgrade, omit, or justify any of these:

- predecessor hard miss `4.646s — HIGH — VERIFIED VIOLATION`;
- continuous-monitoring gap at least `255s — HIGH — VERIFIED VIOLATION`;
- Guard warm warning `40.891s — HIGH — VERIFIED VIOLATION`;
- successor-creation warm miss `146.891s — HIGH — VERIFIED VIOLATION`;
- outgoing Main hard miss `5216.875s — HIGH — VERIFIED VIOLATION`;
- Guard finding that graph queries were used after `status` without the mandatory preceding `anvien analyze --force`: `HIGH — VERIFIED VIOLATION`; those results are not decision evidence.

## Handoff Completion

Outgoing Main terminates after the official authority message, Guard retarget message, and successor navigation. Successor continues from the first incomplete P1-A gate without restarting passed work.
