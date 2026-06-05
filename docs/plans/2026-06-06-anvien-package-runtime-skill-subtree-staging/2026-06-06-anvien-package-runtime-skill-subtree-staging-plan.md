# Plan

Title: Anvien Package Runtime Skill Subtree Staging
Date: 2026-06-06
Status: Completed
Plan: docs/plans/2026-06-06-anvien-package-runtime-skill-subtree-staging/2026-06-06-anvien-package-runtime-skill-subtree-staging-plan.md
Evidence: docs/plans/2026-06-06-anvien-package-runtime-skill-subtree-staging/2026-06-06-anvien-package-runtime-skill-subtree-staging-evidence.md
Benchmark: docs/plans/2026-06-06-anvien-package-runtime-skill-subtree-staging/2026-06-06-anvien-package-runtime-skill-subtree-staging-benchmark.md

## Goal

Align package runtime source staging with the current skill package architecture: `internal/aicontext/skills` is a package subtree and must be staged as an opaque directory artifact, not filtered by file extension.

## Rules

- Write this plan before coding.
- Use Anvien for codebase analysis, impact checks, and change detection.
- Keep the implementation scoped to package runtime staging and its direct tests.
- Do not change the generated AI context or skill package installation architecture unless Anvien evidence proves it is part of the same invariant.
- Code behavior first, then update tests.
- Run the required build and validation sequence before handoff.
- Record evidence and benchmarkable inventory as each phase closes.

## Problem

Anvien evidence shows the current skill package architecture in `internal/aicontext/skill_packages.go` treats a skill package as a package/file snapshot with path and content hashes. Any regular file under the package belongs to the package.

The package runtime source staging surface in `internal/cli/package_runtime.go` still contains older logic in `copyPackageGoDir` that special-cases `internal/aicontext/skills/` and copies only `.md` files. That conflicts with the package invariant because it classifies skill payload by extension.

Deleting the `.md` branch alone is not sufficient, because package runtime staging would then omit the skill subtree entirely. The correct fix is to remove skill responsibility from `copyPackageGoDir` and stage the skill subtree through a package-subtree copy path.

## Scope

In scope:

- `internal/cli/package_runtime.go`
- `internal/cli/package_command_test.go`
- Direct package runtime staging behavior for `internal/aicontext/skills`
- Validation and evidence for package runtime staging

Out of scope:

- Rewriting `internal/aicontext/skill_packages.go`
- Changing generated `.agents` or `CLAUDE.md` output behavior
- Changing skill metadata, skill descriptions, or skill content
- Changing package command UX unless required by the staging invariant

## Requirements

- `copyPackageGoDir` must only stage Go source files.
- `copyPackageGoDir` must continue excluding Go test files from staged Go source.
- Skill staging must copy the whole `internal/aicontext/skills` subtree.
- Skill staging must not filter by extension.
- Skill staging must preserve relative paths and file content.
- Staged source manifest must include copied skill subtree files when present.
- Missing `internal/aicontext/skills` must not fail package source staging unless existing behavior already requires it.

## Validation Sequence

Full build means run the whole command sequence below from the repository root before testing:

```powershell
cd .\anvien
npm install
npm run build
npm install -g .
Get-Command anvien
anvien version
cd ..
powershell -ExecutionPolicy Bypass -File .\anvien-launcher\build.ps1
anvien version
anvien analyze . --force
```

After the full build sequence succeeds, run targeted Go tests for the changed package runtime surface, then run change detection before commit or handoff.

## Technical Direction

Use the existing package runtime staging flow as the integration point:

1. Keep `prepareGoSourcePackage` as the orchestration surface.
2. Keep `copyPackageGoDir` responsible for Go source directories only.
3. Add a dedicated subtree helper for non-Go package artifacts, scoped initially to `internal/aicontext/skills`.
4. Call the subtree helper from `prepareGoSourcePackage` after Go source copying.
5. Append copied skill paths to the existing copied-path inventory before writing the manifest.

The helper should copy regular files under a relative subtree root, create parent directories, preserve paths, and skip missing roots cleanly.

## Acceptance Criteria

- The `.md`-filtered embedded skill branch is removed from `copyPackageGoDir`.
- `internal/aicontext/skills` is staged as a complete subtree artifact.
- A representative non-`.md` file under a synthetic skill subtree is copied in tests.
- Existing Go source staging behavior remains intact, including `_test.go` exclusion.
- Package runtime native DLL copy remains idempotent when npm lifecycle runs build-runtime repeatedly against an identical destination file.
- Targeted Go tests for `internal/cli` pass.
- Full build sequence passes.
- `anvien analyze . --force` completes after implementation.
- `anvien detect-changes --repo Anvien --scope all` is recorded before commit.

## Phase Checklist

- [x] P0-A: Confirm graph-owned scope before editing.
  Goal: Establish the exact implementation and test files before code changes.
  Work Steps: Refresh Anvien graph; use owner discovery for the package runtime staging scope; inspect `copyPackageGoDir` and `prepareGoSourcePackage`; confirm callers and tests; run impact for both symbols and record blast radius.
  Implementation Gate: Do not edit until Anvien identifies the affected implementation and test surfaces.
  Acceptance: Evidence file records the scoped files, symbols, linked tests, and risk level.

- [x] P1-A: Split skill staging responsibility out of `copyPackageGoDir`.
  Goal: Make `copyPackageGoDir` a Go-source-only copier.
  Work Steps: Remove the `embeddedSkillSource` branch; keep `.go` inclusion and `_test.go` exclusion; ensure return inventory still reports copied Go paths.
  Implementation Gate: Impact for `copyPackageGoDir` has been reviewed and recorded.
  Acceptance: `copyPackageGoDir` no longer references `internal/aicontext/skills` or `.md` as a skill-copy rule.

- [x] P1-B: Add subtree staging for `internal/aicontext/skills`.
  Goal: Stage skill packages as opaque directory artifacts.
  Work Steps: Add a helper that copies all regular files under a relative subtree; call it from `prepareGoSourcePackage`; append copied paths to the source manifest inventory.
  Implementation Gate: Impact for `prepareGoSourcePackage` has been reviewed and recorded; helper must use path containment checks and preserve existing package-root safety behavior.
  Acceptance: Missing subtree is skipped safely; present subtree is copied with all regular files and preserved relative paths.

- [x] P2-A: Update package runtime staging tests after code behavior exists.
  Goal: Prove the package invariant without extension-based assumptions.
  Work Steps: Update `TestPrepareGoSourcePackageCopiesMinimalGoSource` to create a synthetic skill package subtree with `SKILL.md` and at least one nested non-`.md` file; assert both contents are staged; assert the staged source manifest includes both skill paths; assert Go `_test.go` remains excluded.
  Implementation Gate: Code behavior from P1-A and P1-B is implemented first.
  Acceptance: Targeted test fails on old `.md`-only logic and passes on subtree-copy behavior, including manifest/inventory coverage.

- [x] P3-A: Validate, analyze, detect changes, and hand off.
  Goal: Close the implementation slice with build and graph evidence.
  Work Steps: Run `gofmt`; run the full build sequence exactly as listed in this plan; run targeted Go tests for `internal/cli`; run `anvien detect-changes --repo Anvien --scope all`; update evidence and benchmark files.
  Implementation Gate: No handoff before validation and detect-changes evidence are recorded.
  Acceptance: Evidence file contains pass/fail results, detect-changes output summary, and final changed-file list.

- [x] P3-B: Handle package-runtime full-build DLL overwrite failure found during validation.
  Goal: Keep package runtime build idempotent when npm lifecycle invokes build-runtime more than once and the native DLL destination already matches the source.
  Work Steps: Run Anvien impact for `copyPackageFileIfExists`; update the helper so identical destination bytes skip overwrite; add a direct unit test for the identical-destination path.
  Implementation Gate: This phase is allowed only because the exact full build sequence failed on `anvien/bin/lbug_shared.dll` overwrite during validation.
  Acceptance: The helper avoids unnecessary writes for identical files and the test proves the behavior.

## Risks

- Package runtime impact is CRITICAL by Anvien blast radius because it reaches package command and root command surfaces. This is a scope warning, not a prohibition.
- A naive deletion of the `.md` branch would silently drop staged skills.
- Extension-based tests would preserve the wrong architecture, so tests must assert subtree equivalence, not file-type allowance.
