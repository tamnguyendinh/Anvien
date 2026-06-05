# Evidence Ledger

Title: Anvien Package Runtime Skill Subtree Staging
Date: 2026-06-06
Status: Completed
Plan: docs/plans/2026-06-06-anvien-package-runtime-skill-subtree-staging/2026-06-06-anvien-package-runtime-skill-subtree-staging-plan.md
Evidence: docs/plans/2026-06-06-anvien-package-runtime-skill-subtree-staging/2026-06-06-anvien-package-runtime-skill-subtree-staging-evidence.md
Benchmark: docs/plans/2026-06-06-anvien-package-runtime-skill-subtree-staging/2026-06-06-anvien-package-runtime-skill-subtree-staging-benchmark.md

## Evidence Rules

- Record Anvien evidence before implementation.
- Record implementation evidence as each phase closes.
- Record build, test, analyze, and detect-changes results before handoff.
- Keep metric tables in the benchmark file.

## E0 - Pre-Implementation Anvien Scope Evidence

These commands are investigation evidence used to write the plan. The implementation agent must refresh Anvien and rerun the P0-A scope checks before editing because the worktree or graph can change after plan creation.

- `anvien analyze . --force`
  Result before plan creation: completed successfully. Graph reported 1399 files, 84382 nodes, 122854 relationships, 16569 file projection dependency edges, and 430 unresolved items.

- `anvien analyze . --force`
  Result after adding this plan set: completed successfully. Graph reported 1402 files, 84408 nodes, 122880 relationships, 16569 file projection dependency edges, and 430 unresolved items.

- `anvien file-context internal/cli/package_runtime.go --repo Anvien --json`
  Result: target file is parsed, backend/cli, risk high, inbound refs 8, linked flows 6, linked tests 2, changedSinceAnalyze false.

- `anvien context symbol "copyPackageGoDir" --repo Anvien`
  Result: symbol found in `internal/cli/package_runtime.go`, range 316-352, parameter count 4, return type `error`.

- `anvien context symbol "prepareGoSourcePackage" --repo Anvien`
  Result: symbol found in `internal/cli/package_runtime.go`, range 132-180, parameter count 2, return type `error`.

- `anvien impact symbol "copyPackageGoDir" --repo Anvien --direction upstream`
  Result: CRITICAL blast radius. Affected files were `internal/cli/package_runtime.go`, `internal/cli/package_command.go`, and `internal/cli/command.go`. Direct caller chain included `prepareGoSourcePackage`, `newPackageCommand`, and `NewRootCommand`.

Required before implementation:

- Rerun `anvien analyze . --force`.
- Rerun scope discovery for the package runtime staging surface.
- Run `anvien impact symbol "copyPackageGoDir" --repo Anvien --direction upstream`.
- Run `anvien impact symbol "prepareGoSourcePackage" --repo Anvien --direction upstream`.
- Record the fresh outputs in this evidence file before editing code.

## E1 - P0-A Fresh Scope Evidence

- Supervisor report scan:
  Result: old supervisor reject/problem files exist, but no open supervisor report was found for the current package-runtime skill-subtree staging scope.

- `anvien analyze . --force`
  Result: completed successfully before implementation. Graph reported 1402 files, 84409 nodes, 122881 relationships, 16569 file projection dependency edges, and 430 unresolved items.

- `anvien query files "package runtime prepare go source skill subtree staging copy skills manifest" --repo Anvien`
  Result: owner discovery identified `internal/cli/package_runtime.go` for package runtime staging and `internal/aicontext/skill_packages.go` as the skill package architecture/source surface.

- `anvien context symbol "copyPackageGoDir" --repo Anvien`
  Result: symbol found in `internal/cli/package_runtime.go`, range 316-352, parameter count 4, return type `error`.

- `anvien context symbol "prepareGoSourcePackage" --repo Anvien`
  Result: symbol found in `internal/cli/package_runtime.go`, range 132-180, parameter count 2, return type `error`.

- `anvien impact symbol "copyPackageGoDir" --repo Anvien --direction upstream`
  Result: CRITICAL blast radius. Direct chain: `copyPackageGoDir` -> `prepareGoSourcePackage` -> `newPackageCommand` -> `NewRootCommand`. Affected files: `internal/cli/package_runtime.go`, `internal/cli/package_command.go`, `internal/cli/command.go`.

- `anvien impact symbol "prepareGoSourcePackage" --repo Anvien --direction upstream`
  Result: CRITICAL blast radius. Direct chain: `prepareGoSourcePackage` -> `newPackageCommand` -> `NewRootCommand` -> `cmd/anvien/main.go`. Affected files: `internal/cli/package_command.go`, `internal/cli/command.go`, `cmd/anvien/main.go`.

## E2 - Architecture Facts

- `internal/aicontext/skill_packages.go` owns the current skill package discovery, package hash, snapshot, sync, and manifest behavior.
- `SkillPackage` contains package files and package hash data; `packageHash` is derived from per-file package paths and file hashes.
- `installSkillPackagesTo` uses desired and actual snapshots, diff planning, apply, verify, and manifest write.
- The package runtime surface is separate from skill package install/sync, but must stage skill source consistently for packaged runtime source output.

## E3 - Implementation Evidence

- P1-A source change:
  `copyPackageGoDir` in `internal/cli/package_runtime.go` no longer contains the `.md`-filtered `embeddedSkillSource` branch. The function now only copies regular `.go` files and still excludes `_test.go`.

- P1-B source change:
  `prepareGoSourcePackage` now excludes `internal/aicontext/skills` from Go source traversal and calls a dedicated `copyPackageSubtree` helper for that subtree. The helper skips missing subtrees, copies all regular files under a present subtree, preserves relative paths, and reuses `copyPackageFile` so destination writes remain guarded by `assertPackageChild`.

- Manifest source change:
  `anvien-go-source.json` keeps the existing numeric `files` field and adds `paths` as a sorted copied-path inventory so tests and validation can prove staged skill package paths are present.

- P3-B validation-failure source change:
  The exact full build sequence failed during `npm install` because package runtime build attempted to overwrite `anvien/bin/lbug_shared.dll` while Windows denied access. `anvien impact symbol "copyPackageFileIfExists" --repo Anvien --direction upstream` reported CRITICAL blast radius limited to package runtime build flow (`copyPackageFileIfExists` -> `copyPackageNativeRuntime` -> `buildGoRuntimePackage` -> `newPackageCommand`). `copyPackageFileIfExists` now reads the destination and skips overwrite when source and destination bytes are identical. This keeps repeated npm lifecycle build-runtime calls idempotent without changing the native runtime payload.

- P2-A test change:
  `TestPrepareGoSourcePackageCopiesMinimalGoSource` now creates `internal/aicontext/skills/planner/assets/config.json`, verifies that the non-`.md` payload is staged, verifies the generated manifest contains both skill paths, and still verifies Go `_test.go` paths are excluded.

- P3-B test change:
  `TestCopyPackageFileIfExistsSkipsIdenticalDestination` verifies that an identical destination file is accepted without rewrite.

## E4 - Validation Evidence

- `gofmt internal/cli/package_runtime.go internal/cli/package_command_test.go`
  Result: completed successfully before validation.

- Full build sequence, first execution:
  Result: failed during `npm install` / package runtime build when Windows denied overwrite of `E:\Anvien\anvien\bin\lbug_shared.dll`. The plain multi-line PowerShell run continued after the npm failure, so this result was treated as failed validation, not a pass.

- Runtime-lock investigation:
  `anvien doctor processes --json` showed a Vite dev server, editor-owned global Anvien MCP processes, and a Playwright server process. `anvien doctor locks --repo Anvien --json` reported the analyze lock free. A loaded-module scan showed editor MCP processes using the global `node_modules\anvien\bin\lbug_shared.dll`, not the local `E:\Anvien\anvien\bin\lbug_shared.dll`.

- Validation failure fix:
  `copyPackageFileIfExists` now skips overwrite when the destination bytes already match the source bytes. This resolved the repeated package-runtime lifecycle build failure without changing native runtime content.

- Full build sequence, fail-fast rerun from repository root:
  Result: completed successfully.
  Commands executed in order:
  `cd .\anvien`; `npm install`; `npm run build`; `npm install -g .`; `Get-Command anvien`; `anvien version`; `cd ..`; `powershell -ExecutionPolicy Bypass -File .\anvien-launcher\build.ps1`; `anvien version`; `anvien analyze . --force`.
  Version checks returned `1.2.5`.
  Final analyze reported 1402 files, 84437 nodes, 122926 relationships, 16570 file projection dependency edges, and 430 unresolved items.

- `go test ./internal/cli -count=1`
  Result: passed. Package result: `ok github.com/tamnguyendinh/anvien/internal/cli 64.427s`.

- Final pre-commit graph refresh:
  `anvien analyze . --force` completed successfully after adding report/notes docs. Graph reported 1404 files, 84448 nodes, 122937 relationships, 16570 file projection dependency edges, and 430 unresolved items.

- Repo-local package prepare smoke:
  Commands executed from `.\anvien` through `go run ..\cmd\anvien package prepare-go-source`, inventory check, and `go run ..\cmd\anvien package clean-go-source`.
  Result: passed. Prepare output copied 929 files to `E:\Anvien\anvien\go-src`; cleanup removed the staged `go-src` directory.

- Repo-local package prepare inventory:
  Skill subtree files staged: 638.
  Non-`.md` skill files staged: 295.
  Manifest file count: 929.
  Manifest skill-path count: 638.
  Representative non-`.md` staged paths included `internal/aicontext/skills/better-auth/scripts/better_auth_init.py`, `internal/aicontext/skills/better-auth/scripts/requirements.txt`, `internal/aicontext/skills/better-auth/scripts/tests/test_better_auth_init.py`, and `internal/aicontext/skills/databases/scripts/db_backup.py`.

## E5 - Detect Changes Evidence

- `anvien detect-changes --repo Anvien --scope all`
  Result: completed successfully.
  Summary risk: medium.
  Affected process reported: `BuildGoRuntimePackage -> CopyPackageFileIfExists`.
  Final changed app layers: backend 51, backend_test 10, docs 8.
  Final changed functional areas: cli 61, documentation 8.
  Changed implementation/test files for this slice: `internal/cli/package_runtime.go` and `internal/cli/package_command_test.go`.
  Detect output also listed `internal/aicontext/skills/coder/SKILL.md`; that file is not part of this package-runtime slice and is not included in the implementation commit.

## E6 - Commit Evidence

Pending until git checkpoint is created.
