# Main Orchestration Rotation Handoff — Child 04 P4-A

Date: 2026-08-20
Created at: 2026-08-20 21:17:34 +07:00
Outgoing Main task: `01a01f30-8caf-76b1-bc9f-049a50f7b52a`
Incoming authority: successor Main must be created immediately from this handoff
Resolved cwd: `E:\Anvien`

## Rotation Status

- Outgoing authority task was created at `2026-08-20 19:41:10 +07:00`.
- Required 60-minute handoff deadline was `2026-08-20 20:41:10 +07:00`.
- This handoff was created about `36` minutes late. The overrun is a process deviation and must not be repeated.
- Successor Main must record its own creation timestamp and schedule handoff at exactly `+60 minutes`, preparing the next successor before the deadline even if a Child lane is still active.

## Accepted Campaign State

- Child 03 Pn-C closes at `3e25f9ca75c9cb7d59bf228e6cc9aa0b81d738b4`.
- Child 04 P0-A closes at `ff2467bb92f94a9c53c4de030685686700051a98`, parent `3e25f9ca75c9cb7d59bf228e6cc9aa0b81d738b4`.
- P0-A manifest is exact roadmap plus four Child 04 living ledgers; post-commit worktree was clean; no push.
- Sole open slice: Child 04 `P4-A`.
- `P4-B`, `P4-B1`, `P4-C`, `P4-C2`, Child 05, and every later lane remain locked.

## Active Visible Coder Task

- Title: `Child 04 P4-A — Coder`
- threadId: `01a01f84-9127-7a10-b589-c2b4721597ae`
- hostId: `local`
- Latest known cursor: `9721fd12-8edf-44b0-9de7-69838ed499c0:22`
- Snapshot state: active / first turn in progress.
- Latest progress: fresh graph on the current draft reported `1,126/626/0`, `81,034/120,392`; Coder is reading the exact six owner/test files and draft diff before symbol impact. It has not yet edited production.
- Continue this exact visible task. Do not interrupt or create a duplicate Coder lane unless it reports a blocker or terminates unexpectedly.
- Hidden subagent execution ownership was explicitly stopped. Hidden/read-only helpers inside the visible Coder task are allowed, but the visible task owns all edits, build, tests, and handoff.

## Current Git / Worktree Boundary at Snapshot

- HEAD: `ff2467bb92f94a9c53c4de030685686700051a98`.
- Parent: `3e25f9ca75c9cb7d59bf228e6cc9aa0b81d738b4`.
- Index: clean.
- Exactly four unstaged production paths existed before this handoff report:
  - `internal/scopeir/facts.go`
  - `internal/scopeir/ir.go`
  - `internal/scopeir/kinds.go`
  - `internal/scopeir/sort_keys.go`
- Draft stat at snapshot: `4 files changed, 187 insertions`, with no deletion.
- This handoff report is the only additional Main-owned path authorized concurrently.
- Any other drift must be attributed to the active visible Coder task before acceptance; unknown drift is a stop condition.

## P4-A Exact Boundary

Production editable:

- `internal/scopeir/facts.go`
- `internal/scopeir/kinds.go`
- `internal/scopeir/ir.go`
- `internal/scopeir/sort_keys.go`

Test-after-code editable:

- `internal/scopeir/scopeir_test.go`
- `internal/scopeir/testdata/scopeir.golden.json`

Inspect/preserve-only:

- TSJS extraction owners;
- graph projection/persistence consumers;
- `DefinitionFact.Visibility`;
- Child 05 terminal resolution, barrels, ambiguity, cycles, and public API.

Required invariant:

- one immutable `ExportFact` per exported binding/specifier;
- export kind, names, ranges, meaning lanes, type-only state, and source provenance are representable;
- dual value/type meaning is representable without guessing;
- unsupported/malformed export syntax has a structured, countable diagnostic;
- `ScopeIR` clone/normalize/JSON behavior is deterministic and deep-copies nested mutable values;
- access visibility is unchanged;
- no AST extraction, compatibility projection, terminal target, barrel, or public-API state is introduced in P4-A.

## Draft Provenance and Non-Evidence

- The four-file draft was started by outgoing Main before role correction and transferred to the visible Coder for zero-trust review.
- The Coder must accept, change, or replace the draft based on the P4-A contract; outgoing Main does not claim it is correct.
- An outgoing-Main `npm run full-build` attempt was interrupted during package runtime construction and exited `1`; it is not build evidence.
- No P4-A focused test, product test, Coder report, Supervisor verdict, detect result, or commit exists yet.
- No P4-A path is staged and no push is authorized.

## Successor Main Responsibilities

1. Read `AGENTS.md`, `working-rules`, orchestration skill, this handoff, roadmap, and all four Child 04 living ledgers completely.
2. Verify this report identity and the Git/task snapshot once; do not repeat implementation-grade audit while Coder is active.
3. Monitor visible Coder task `01a01f84-9127-7a10-b589-c2b4721597ae` from cursor `9721fd12-8edf-44b0-9de7-69838ed499c0:22`.
4. When Coder returns `READY_FOR_SUPERVISOR`, read its durable report and exact boundary, then create one visible independent Supervisor task. Do not self-accept and do not use a hidden Supervisor as lane owner.
5. If Supervisor rejects, route only the exact rejected invariant back to the same visible Coder task.
6. If Supervisor passes, use planner once to refresh roadmap plus four Child 04 ledgers; run required excluded graph/change detection; stage the exact accepted P4-A boundary; create one isolated commit; do not push.
7. Only after P4-A commit may P4-B open automatically as a new visible Coder task.
8. Preserve `E:\cheapapp.org` until P4-C2 and never use/stage `internal/aicontext/skills/**` or `.claude/skills/**` as Child 04 evidence.
9. Maintain the 60-minute Main rotation independently of Child completion. Prepare and create the next Main successor before the deadline.

## Immediate Next Action

Monitor the existing visible Coder task from the transferred cursor. In parallel, record the successor rotation deadline; do not edit the Coder-owned worktree or open a duplicate implementation lane.
