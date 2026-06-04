# Evidence Ledger

Date: 2026-06-04

Status: Approved - implementation pending

Companion files:

- Plan: [2026-06-04-anvien-skill-mirror-incremental-sync-plan.md](2026-06-04-anvien-skill-mirror-incremental-sync-plan.md)
- Benchmark ledger: [2026-06-04-anvien-skill-mirror-incremental-sync-benchmark.md](2026-06-04-anvien-skill-mirror-incremental-sync-benchmark.md)

## Evidence Rules

1. Record evidence as each evidenced task completes.
2. Use Anvien for codebase analysis, impact checks, and implementation-slice evidence.
3. Record build and validation pass/fail results after implementation.
4. Record failures and how they were handled.
5. Record `anvien detect-changes --repo Anvien --scope all` before commit.
6. Do not treat generated `.claude/skills/anvien/**`, `AGENTS.md`, or `CLAUDE.md` as source of truth.

## E0 - User Report And Policy Clarification

User reported that deleting `internal\aicontext\skills\debugging` did not remove `.claude\skills\anvien\debugging` after `anvien analyze`.

Clarified requirement:

- the bug is broader than deletion;
- every change under `internal/aicontext/skills/**` must propagate to `.claude/skills/anvien/**`;
- tests are verification tools, not source of truth;
- implementation should follow PA4: exact generated mirror semantics with incremental diff application.

Approval:

- 2026-06-04: user requested `commit` and then `trien khai plan`, approving the plan for implementation.

## E1 - Initial Anvien And Source Inspection

Commands already run before plan creation:

```text
anvien analyze --force
anvien query "sync skills internal aicontext .claude skills anvien analyze" --repo Anvien
anvien context symbol "installBaseSkills" --repo Anvien
anvien context symbol "InstallBaseSkillsTo" --repo Anvien
anvien file-context internal/aicontext/skill_packages.go --repo Anvien --json
```

Initial findings:

- `internal/aicontext/skill_packages.go` owns the skill package install and manifest behavior.
- Existing behavior preserves missing source packages by marking manifest entries stale.
- `.claude/skills/anvien/.anvien-skill-manifest.json` can contain entries with `"stale": true`.
- Current behavior does not enforce `.claude/skills/anvien/**` as a mirror of source.

## E2 - Worktree Hygiene Before Plan Creation

An accidental implementation patch was applied during discussion and then reverted before this plan set was created.

Verification:

```text
git diff -- internal/aicontext/skill_packages.go
git status --short
```

Result:

- no remaining diff in `internal/aicontext/skill_packages.go`;
- worktree was clean before adding this plan set.

## E3 - Plan Review Evidence

Commands run while reviewing this plan:

```text
anvien analyze --force
anvien file-context docs/plans/2026-06-04-anvien-skill-mirror-incremental-sync/2026-06-04-anvien-skill-mirror-incremental-sync-plan.md --repo Anvien --json
```

Review adjustments made:

- clarified that the generated output root is the installer `targetDir`, normally `.claude/skills/anvien/**`;
- clarified that the snapshot walker must be generic for both `InstallSkillPackagesTo(targetDir)` and `InstallSkillPackagesForRepoTo(targetDir, repoPath)`;
- added required sync counters for result summary and benchmark evidence;
- added acceptance criteria for result-summary behavior without stale preservation;
- tightened `P2-A` to require an explicit `SkillInstallResult` / `Summary()` compatibility decision;
- tightened `P3-D` to verify no generated payload remains outside the desired snapshot;
- clarified that acceptance covers both `anvien analyze` and generated-skill setup installation;
- reframed the problem around source-change propagation instead of deletion as the central symptom;
- added the rule to update checklist items immediately when completed.

## E4 - Implementation Evidence

Pending. Fill after plan approval and implementation.

Expected entries:

- files changed;
- summary of snapshot structures;
- summary of diff/apply behavior;
- manifest behavior changes;
- generated guidance source changes.

## E5 - Validation Evidence

Pending. Fill after implementation.

Expected entries:

- full build result;
- focused `internal/aicontext` tests;
- focused CLI analyze/setup tests;
- real repo `anvien analyze --force` validation;
- no Web UI e2e required unless UI scope changes.

## E6 - Detect Changes And Commit Evidence

Pending. Fill after validation.

Expected entries:

- `anvien detect-changes --repo Anvien --scope all`;
- changed files and affected flows;
- commit hash;
- closure summary.
