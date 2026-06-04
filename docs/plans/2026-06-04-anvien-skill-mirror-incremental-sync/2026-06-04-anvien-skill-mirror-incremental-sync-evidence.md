# Evidence Ledger

Date: 2026-06-04

Status: Complete

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

P0 gate evidence recorded before implementation edits.

Commands:

```text
git rev-parse --short HEAD
git status --short
anvien analyze --force
anvien impact symbol "installSkillPackagesTo" --repo Anvien --direction upstream
anvien impact symbol "InstallSkillPackagesTo" --repo Anvien --direction upstream
anvien impact symbol "InstallSkillPackagesForRepoTo" --repo Anvien --direction upstream
anvien impact symbol "GenerateAIContextFiles" --repo Anvien --direction upstream
anvien impact symbol "installBaseSkills" --repo Anvien --direction upstream
anvien impact symbol "setupInstallSkillsTo" --repo Anvien --direction upstream
```

Result:

- implementation starts from plan commit `c1add09`;
- worktree was clean immediately before implementation edits;
- fresh analyze completed with `scanned=1345`, `parsed_code=703`, `graph nodes=84137`, `relationships=122901`;
- `installSkillPackagesTo` impact risk: HIGH; affected files include `internal/aicontext/aicontext.go`, `internal/aicontext/skill_packages.go`, `internal/cli/analyze_postrun.go`, and `internal/cli/setup_command.go`;
- `InstallSkillPackagesTo` impact risk: LOW; affected files include `internal/aicontext/skill_packages.go` and `internal/cli/setup_command.go`;
- `InstallSkillPackagesForRepoTo` impact risk: LOW with no upstream impacted symbols reported;
- `GenerateAIContextFiles` impact risk: CRITICAL; affected files include `internal/aicontext/aicontext.go`, `internal/cli/analyze_postrun.go`, and `internal/cli/command.go`;
- `installBaseSkills` impact risk: CRITICAL; affected files include `internal/aicontext/aicontext.go`, `internal/cli/analyze_postrun.go`, and `internal/cli/command.go`;
- `setupInstallSkillsTo` impact risk: LOW; affected file is `internal/cli/setup_command.go`;
- HIGH/CRITICAL is treated as blast-radius warning, not a blocker; changes remain scoped to skill sync/guidance behavior.

Implementation notes:

- changed `internal/aicontext/skill_packages.go`;
- changed `internal/aicontext/aicontext.go`;
- added explicit generated-output file snapshots with generated path, hash, package name, package path, and desired file content;
- desired snapshot is built from discovered `SkillPackage.Files`, keyed by normalized install path, with duplicate generated paths reported as collisions;
- actual snapshot walks `targetDir`, skips `.anvien-skill-manifest.json`, hashes regular files, and rejects non-regular filesystem entries as unsafe;
- diff classifies writes, overwrites, deletes, and unchanged files by hash;
- sync applies obsolete-file deletes, writes missing files, overwrites changed files, prunes empty dirs, verifies the post-apply target snapshot, then writes the manifest last;
- manifest is rebuilt only from current source packages; deleted source packages are not preserved as stale manifest entries;
- `SkillInstallResult.Summary()` now reports package counters plus `files_written`, `files_overwritten`, `files_deleted`, and `files_skipped`;
- generated guidance now states that `.claude/skills/anvien/**` mirrors the source snapshot and custom skills belong outside that generated namespace.
- obsolete stale-preservation, unmanaged-collision, and legacy-adoption helper functions were removed after confirming they had no remaining callers.

## E5 - Validation Evidence

Build gate completed before tests:

```text
powershell -ExecutionPolicy Bypass -File anvien-launcher\build.ps1
```

Result:

- PASS;
- Go version reported by build: `go version go1.26.3 windows/amd64`;
- final web build completed with Vite in `18.68s`;
- existing Vite chunk-size warnings were reported; no build failure.

Focused tests and real repo validation remain pending.

Manual synthetic verification before writing new tests:

```text
go run ./cmd/anvien analyze <temp-repo> --force --skip-git --no-stats
```

Result:

- first manual script failed only at the PowerShell manifest property-count check after file-level checks had already passed;
- rerun with explicit NoteProperty counting passed;
- verified source edit propagation, stale output deletion, stale package deletion, source package deletion, and manifest removal for a synthetic `internal/aicontext/skills/demo` package;
- pass marker: `manual-sync-pass`.

Focused validation:

```text
go test ./internal/aicontext
go test ./internal/cli -run "TestSetupInstallsEmbeddedSkillsInsteadOfPackageRootSkills|TestAnalyzeCommandGeneratesAIContextByDefault"
```

Result:

- first aicontext run failed because legacy tests hardcoded removed real catalog package `debugging`;
- tests were corrected to use synthetic behavior fixtures or dynamic catalog checks instead of treating the current real catalog as source of truth;
- final `go test ./internal/aicontext`: PASS;
- final focused `go test ./internal/cli ...`: PASS;
- no Web UI behavior changed; no e2e test required.

Real repo validation:

```text
anvien analyze --force
```

Result:

- PASS;
- analyze output: `scanned=1345`, `parsed_code=703`, `graph nodes=84196`, `relationships=123037`;
- inventory check: `source_packages=34`, `manifest_packages=34`, `source_payload_files=593`, `target_payload_files=593`;
- target diff check: `missing=0`, `extra=0`, `mismatch=0`;
- deleted source symptom check: `debugging_source=False`, `debugging_target=False`.

No Web UI behavior changed, so no e2e test was required.

## E6 - Detect Changes And Commit Evidence

Detect changes completed before commit:

```text
anvien analyze --force
anvien detect-changes --repo Anvien --scope all
```

Result:

- PASS;
- changed files reported: 7;
- affected files reported: 7;
- overall risk level: HIGH;
- high-risk changed files include `internal/aicontext/skill_packages.go` and `internal/aicontext/aicontext.go`;
- low-risk changed files include the plan/evidence/benchmark docs, `internal/aicontext/aicontext_test.go`, and `internal/cli/command_test.go`;
- affected processes include `InstallSkillPackagesTo -> CleanSkillInstallPath`, `InstallSkillPackagesTo -> HashBytes`, `InstallSkillPackagesTo -> RemoveEmptyDirsUpTo`, `InstallSkillPackagesTo -> RemoveEmptySkillDirs`, `InstallSkillPackagesTo -> SafeJoin`, and `InstallSkillPackagesTo -> SkillFileSnapshot`;
- HIGH risk is expected for this installer change and was handled with scoped implementation, full build, focused tests, manual synthetic validation, and real repo mirror inventory validation.

Commit evidence:

- implementation commit: `b4ebc76` (`fix: mirror generated Anvien skills`);
- closure summary: PA4 exact generated mirror sync is implemented, validated, benchmarked, and committed;
- doc-only closure commit follows this evidence update.
