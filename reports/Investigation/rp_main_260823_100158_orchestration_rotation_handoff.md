# Anvien Analyze Performance — Main Orchestration Rotation Handoff

Status: `SEALED_CANDIDATE / OFFICIAL_TRANSFER_PENDING / FOUR_LANES_ACTIVE / ROOT_CAUSE_REQUIRED / P6-D OPEN`

Internal createdAt: `2026-08-23 10:01:58 +07:00`

Outgoing Main task: `01a02c2d-f372-70a3-bb11-93e9aa4ba59f`

Warmed successor task: `01a02c90-7232-7531-8393-33eda9af7fe8`, host `local`, saved project/workspace `E:\Anvien`, cursor `a4de4a3a-6716-4982-86f4-277c9c65aa0a:1`; it answered `UNDERSTOOD / WAITING_FOR_OFFICIAL_TRANSFER` and has performed no action.

The outgoing Main inherited a warmup target `2026-08-23 09:04:07 +07:00` and absolute transfer deadline `2026-08-23 09:19:07 +07:00`. Both were missed. At the first explicit clock audit this task observed `2026-08-23 10:00:23 +07:00`. This violation is disclosed and must remain in successor handoffs; it is not normalized away.

If official transfer follows immediately from this handoff, successor next warmup target is `2026-08-23 10:46:58 +07:00` and next absolute transfer deadline is `2026-08-23 11:01:58 +07:00` while the campaign remains active.

## Raw-rule inheritance — mandatory, not a summary

- Successor MUST read and apply the complete raw, unmodified contents of `E:\Anvien\AGENTS.md`, `E:\Anvien\.agents\skills\working-rules\SKILL.md`, and `E:\Anvien\.agents\skills\orchestration\SKILL.md` before action.
- It is forbidden to use this handoff, a checkpoint, memory, chat paraphrase, or any shortened/compacted rendering as a replacement for any line of those three raw files.
- Successor must also read the complete visible-session template, the complete active skills it directly uses, the active Child 06 plan and all four ledgers, and the latest durable reports before deciding.
- This exact anti-summarization inheritance requirement must be transmitted unchanged in every later Main handoff.

## Owner-mandated orchestration behavior — transmit through every Main handoff

- Main is the meeting chair and orchestration authority, not a worker and not a report courier.
- Main must understand the full request, plan, phase/slice functions, pipeline, evidence, and acceptance before giving lane commands.
- Main must assign the correct visible lane/capability for each actual job, monitor actual commands/files/scope, and intervene immediately on wrong command, scope drift, loop, vague claim, or wrong target.
- A reminder, question, rebuttal, or status request is NOT PAUSE. Main must not final/yield or stop campaign work merely to answer. Only explicit `PAUSE` or `STOP` halts.
- Supervisor runs continuously as the later review lane. Do not interrupt it except for infinite loop, scope drift, or a concrete unsafe/wrong command. Supervisor does not govern Main.
- No hidden/internal lane may replace visible Planner, Architect, Supervisor, coder, QA, security, or long-running specialist work.
- Every finding and conclusion requires cause plus evidence. Owner statements and lane reports are evidence pointers, never automatic truth.
- Main must not answer a lane's challenge on behalf of another lane. Every Planner question for Architect must be relayed verbatim to Architect; every Architect question for Planner must be relayed verbatim to Planner; their answers must be relayed back. Root Cause evidence must be sent to both, and Supervisor challenges must be routed to the owning lane. Main evaluates the exchange but does not perform the employees' work.
- Final output to Owner must be the best evidence-backed option for making `analyze` faster without sacrificing accuracy, semantic completeness, deterministic output, freshness, persistence/reader parity, or accepted graph invariants.
- These prior Main violations and reminders are mandatory handoff content: stopping to reply; waiting on Owner instead of coordinating; treating Supervisor as authority over Main; opening hidden lanes; using the wrong full-build command; failing to question vague reports; reducing raw rules/skills to summaries; assigning too few/wrong capabilities; failing to relay lane questions; and doing worker analysis instead of chairing the meeting.

## Workspace, target, and campaign boundary

- Only workspace: `E:\Anvien`.
- `C:\`, alternate worktree, undelegated path, network/install/package action, stale/shared/quarantined cache substitution, protected-root reuse/cleanup workaround, and `E:\cheapapp.org` remain forbidden unless later exact Owner authority changes a named boundary.
- Target `E:\cheapapp.org` remains locked and must not be accessed.
- P6-D is the sole open campaign slice and remains unchecked, unstaged, and uncommitted.
- P6-A/B/C1/C2/C3 remain accepted/committed. Pn-A/Pn-B/Pn-C and Child 07 remain locked.
- Exact local P6-D six-field projector repair remains independently cleared. Full P6-D remains REJECT/BLOCKED for current-runtime/native/built-reader/target/oracle/pre-post evidence and procedural cleanliness.
- The committed P6-C3 same-invariant omission remains out of P6-D authority and must be routed separately.
- No discussion lane is authorized to edit code, tests, active plan/ledgers, target, AGENTS/skills, stage, commit, clean, or use protected roots. Each may create at most its assigned new report/artifacts under its exact boundary, uncommitted.

## Current Git, candidate, and report truth

- Measured `2026-08-23 10:01:37.6396628 +07:00`.
- Branch `master`; HEAD `c28660118fc606ea3e19ad2f6cfb206a768f46f1`; parent `1838be3074a508d442481fc8720ce3afa6745404`.
- HEAD adds only committed orchestration-skill history above P6-C3. `git diff 0d980377..HEAD` changes only `internal/aicontext/skills/orchestration/SKILL.md`; this is outside the P6-D candidate.
- Status `64 = 5 tracked + 59 untracked`; index/staged `0`.
- Tracked paths are exactly the four Child 06 ledgers plus `internal/graphhealth/diagnostics.go`. Sole non-report untracked path is `internal/graphhealth/p6d_outcome_projection_test.go`.
- Exact seven-path P6-D candidate remains `377,945` bytes, ordinal/case-sensitive canonical SHA-256 `E4FF42DE8AA1941010AFFF5DD38B6789AE71168E3EBF9F59D8EECF735D5BD044`.
- The old sealed handoff identity for the P6-D resubmission Supervisor report is stale. Current file `reports/Supervisor/rp_supervisor_260823_064551_by_gpt-5_p6d_graph_health_projection_resubmission.md` was Owner-amended after seal: `33,046` bytes, SHA-256 `3837059F852CF44F20732F2BB0F7D88234BC6612AAB9C7CD439AEC3288A768DA`, lastWrite `2026-08-23T09:33:10.5645217+07:00`. It still returns REJECT, corrects canonical full build to `scripts/full-build.ps1`, and records analyze performance/liveness as unresolved.
- The four Child 06 ledger identities remain the prior exact candidate identities and still end at R37. They do not yet contain the Owner's new performance request, the canonical full-build rerun, or a comparable analyze-duration benchmark. Do not update them during architecture discussion.

## New performance requirement and current evidence grade

- Owner reports operational regression: analyze was about two minutes, then five, now above seven; graph generation delay blocks all repo work.
- Completion now requires an evidence-backed root cause, not merely choosing P6-D versus P6-E.
- Accuracy cannot be traded for speed. Forbidden shortcuts include scanning/parsing fewer eligible inputs, suppressing edges/outcomes/proofs, stale graph/cache reuse, weaker validation, or nondeterministic concurrency.
- Direct visible-task evidence was recovered from task `01a02c5b-333a-7783-a074-d6a6a106991d`:
  - retained successful command was exactly `.\scripts\full-build.ps1`, exit `0`, tool duration `678,182 ms` (`11m18.182s`);
  - the task turn duration `1,042,071 ms` included one first full-build run whose terminal handle was lost plus one complete rerun, so `17m20s` is not a single-command timing;
  - successful rerun reports Anvien `1.2.8`, `2,074` scanned / `763` parsed / `0` failed, graph `120,843` nodes / `167,418` relationships, `20,310` dependency edges, `478` unresolved file projections;
  - the lane's later `~8 minutes` analyze statement was inferred from polling and has no exact analyze start/end, benchmark JSON, or phase timing;
  - canonical run used a global wrapper located on C according to its output, so it is historical canonical-build evidence, not E-only runtime/profile authority.
- The active Child 06 benchmark ledger contains graph inventory but no comparable analyze wall-time protocol/baseline/threshold. Historical E-resident benchmark artifacts show much smaller prior totals, but no causal regression claim is accepted without same-condition comparison.
- Current source supports `--benchmark-json`, per-phase timings, `--cpuprofile`, and `--memprofile`; a dedicated visible debugging lane owns the single controlled profiling run. Main must not perform that worker run.

## Active visible lanes and meeting protocol

All four lanes are active in saved project `E:\Anvien`, environment local, no worktree. They received complete goals, authority, scope/non-goals, evidence, stop/completion conditions, and mandatory `UNDERSTOOD/NOT UNDERSTOOD` protocol.

1. Planner `01a02c87-884d-7a93-be54-488b832085a2`, cursor `f5fecf50-cf2d-454a-9273-1db97d8feb8a:3`, active.
   - Owns plan topology, exact phase/slice IDs, benchmark/evidence gates, and commit boundaries; no ledger edit.
   - Round 1 requested: provisional P6-D/P6-E choice, three exact questions for Architect, and evidence that would change its position.
2. Architect `01a02c87-8ccc-7482-94a6-79be7bdb587a`, cursor `93bc712c-9c55-4f42-ac6c-e6fdf0b77e96:3`, active.
   - Owns architecture/causal boundary, accuracy-preserving performance contract, invalidation/rollback analysis; no SPEC/plan/code edit.
   - Round 1 requested: provisional choice, three exact questions for Planner, and causal evidence that would change its position.
3. Supervisor `01a02c87-961d-71e1-a3d2-dd1121d66b4c`, cursor `6feacb17-3b74-46f7-9e84-d9e2c4a6fb99:4`, active.
   - Review-only, continuous. Current preliminary verdict is REJECT for claim that current evidence already authorizes optimization implementation inside P6-D.
   - Must review the complete Planner/Architect exchange and Root Cause artifacts after they exist; it intentionally has not sealed the final report yet.
4. Root Cause / Debugging `01a02c8b-37d5-7b72-ab97-12d48326332e`, cursor `6fca6aa6-c21e-4ab8-ae07-45b4c7f8c4a4:6`, active.
   - Owns diagnosis only. It is reading raw rules/skills/ledgers/source before one controlled E-only current-runtime `analyze . --force` with benchmark JSON and CPU/memory profiles, all artifacts under `E:\Anvien\.tmp\analyze-performance-root-cause-01a02c2d`.
   - No second analyze without exact reason and Main approval; no code/plan/target/network/build/cleanup/Git mutation.

Meeting sequence successor must enforce:

1. Get Planner and Architect Round 1 messages.
2. Relay Planner's three questions verbatim to Architect and Architect's three questions verbatim to Planner; do not answer for either.
3. Relay each response back to the question owner. Require explicit disposition of every disagreement.
4. Monitor Root Cause actual invocation before it starts; verify it is the E-only current binary, exact repo-local artifact root, retained handle, one run, correct profile flags, no target/network/protected root. Intervene immediately on deviation.
5. Relay Root Cause phase timings/profile/source attribution to Planner, Architect, and Supervisor. Require Planner/Architect to revise or retain their positions explicitly from the measured cause.
6. Send the resulting proposals/answers to Supervisor for final zero-trust review. Relay any Supervisor challenge to the owning lane; Supervisor must not be interrupted merely for being long.
7. Main self-verifies reports/artifacts/current Git boundary and chooses the final phase/plan structure. Do not decide by voting or role title.
8. Report the chosen option, verified root cause, exact workflow/slices, performance protocol, accuracy-equivalence gates, risks, and next ownership to Owner. No code/plan edit or implementation begins during this discussion without later authority.

## Exact next ownership

1. Officially accept this transfer only after zero-trust seal verification and raw-rule reload.
2. Continue the live meeting immediately; do not stop merely to answer status/reminders.
3. Relay cross-lane questions and answers exactly as specified; Main does not answer employee questions for them.
4. Monitor the Root Cause long run and update Owner before it starts, during long execution, and when evidence arrives.
5. Preserve P6-D unchecked/unstaged/uncommitted and target locked.
6. Do not create another Supervisor or Root Cause lane; current visible lanes remain the owners unless actual scope/loop evidence requires intervention.
7. Warm next Main by `2026-08-23 10:46:58 +07:00` and officially transfer by `2026-08-23 11:01:58 +07:00` if campaign remains active.

