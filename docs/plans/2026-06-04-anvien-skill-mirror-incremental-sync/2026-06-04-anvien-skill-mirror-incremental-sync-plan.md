# Anvien Skill Mirror Incremental Sync Plan

Date: 2026-06-04

Status: Approved - implementation pending

Companion files:

- Evidence ledger: [2026-06-04-anvien-skill-mirror-incremental-sync-evidence.md](2026-06-04-anvien-skill-mirror-incremental-sync-evidence.md)
- Benchmark ledger: [2026-06-04-anvien-skill-mirror-incremental-sync-benchmark.md](2026-06-04-anvien-skill-mirror-incremental-sync-benchmark.md)

## Master Rules

1. Follow active repository instructions and generated `AGENTS.md`.
2. Use Anvien for codebase analysis, impact checks, and implementation-slice evidence.
3. Run `anvien analyze --force` before graph-based Anvien commands.
4. Run impact analysis before editing `internal/aicontext/aicontext.go`, `internal/aicontext/skill_packages.go`, setup/install owners, generated context owners, or shared skill data contracts.
5. HIGH or CRITICAL impact is a blast-radius warning, not a blocker. Report it clearly and keep changes scoped.
6. Do not edit generated `AGENTS.md`, `CLAUDE.md`, or `.claude/skills/anvien/**` as source of truth.
7. `internal/aicontext/skills/**` is the source of truth for Anvien skills.
8. `.claude/skills/anvien/**` is generated output and must mirror the current source snapshot after `anvien analyze` and generated-skill setup installation.
9. Run the full build before tests. For this repo, the full build gate is `powershell -ExecutionPolicy Bypass -File anvien-launcher\build.ps1`.
10. Add e2e tests only if Web UI behavior changes.
11. Record evidence as each evidenced task completes.
12. Record benchmarkable inventory, size, and sync counters as each benchmarkable task completes.
13. Update each checklist item immediately when its task is completed.
14. Run `anvien detect-changes --repo Anvien --scope all` before implementation commits.
15. Commit completed implementation slices only after build, validation, evidence, benchmark, and detect-changes are complete.

## Goal

Implement PA4: **exact generated mirror semantics with incremental file-level diff application** for Anvien skill installation.

After `anvien analyze`, the target generated skill root, normally `.claude/skills/anvien/**`, must be the generated projection of `internal/aicontext/skills/**`. Every source change must propagate:

- edited source file updates the corresponding output file;
- added source file or package creates output;
- deleted source file or package removes output;
- renamed source path behaves as delete old output plus create new output;
- tampered or missing output is repaired from source;
- stale generated output not present in source is removed.

The implementation must avoid rewriting every file when only a subset changed. It should compute source and output snapshots, diff them, apply only required writes/deletes, then write the new manifest.

## Problem

The current installer uses package hashes for existing source packages, but it does not enforce full source-to-output change propagation. A source change can leave generated output out of sync instead of making `.claude/skills/anvien/**` become the current projection of `internal/aicontext/skills/**`.

Deleted source packages are one observed example: the installer can mark the manifest entry stale and preserve the generated output. The broader bug is that source change is not treated as a required output sync. The installer needs a snapshot sync model, not a stale-preservation or partial-update model.

## Scope

In scope:

- `internal/aicontext/skill_packages.go` skill package discovery, install, manifest, hashing, and sync behavior.
- `internal/aicontext/aicontext.go` generated AI context install path if needed.
- `anvien analyze` AI-context generation path.
- `anvien setup` editor skill installation path.
- Generated Anvien skill output under `.claude/skills/anvien/**`.
- Tests and synthetic fixtures for source edits, additions, deletions, renames, output tampering, and missing output.
- Evidence and benchmark ledgers for blast radius, build/test results, sync counters, and generated-output inventory.

Out of scope:

- Web UI behavior.
- Editing generated `.claude/skills/anvien/**` as source.
- Preserving custom files inside `.claude/skills/anvien/**`; that namespace is generated output for this plan.
- Redesigning the skill package format beyond mirror sync behavior.

## Invariants

1. Source of truth is `internal/aicontext/skills/**`.
2. Generated output root is the installer `targetDir`, normally `.claude/skills/anvien/**`.
3. The generated output root must be an exact mirror/projection of the current source package snapshot after analyze/setup skill installation completes.
4. Manifest data is metadata and sync evidence, not source of truth.
5. The sync implementation must use file hashes to avoid unnecessary rewrites.
6. The sync implementation must remove generated output paths that no longer exist in source.
7. The sync implementation must never use tests as source of truth for the current catalog contents.
8. Tests must use synthetic source fixtures for behavior that should survive frequent skill catalog changes.

## Technical Direction

Use a snapshot-diff-apply model:

1. Build `desiredSnapshot` from source packages:
   - generated relative file path;
   - source relative file path;
   - SHA-256 hash;
   - file content or deferred content reader;
   - package metadata and entry metadata.
2. Build `actualSnapshot` from the installer target root, normally `.claude/skills/anvien/**`:
   - generated relative file path;
   - SHA-256 hash;
   - exclude `.anvien-skill-manifest.json`;
   - keep target-root walking generic so `InstallSkillPackagesTo(targetDir)` and `InstallSkillPackagesForRepoTo(targetDir, repoPath)` share the same sync engine.
3. Diff:
   - desired path missing in actual -> write;
   - desired path exists with different hash -> overwrite;
   - actual path missing in desired -> delete;
   - actual path exists with same hash -> skip.
4. Apply the diff in a safe order:
   - create parent dirs before writes;
   - write/overwrite desired files;
   - delete obsolete output files;
   - remove empty dirs below `.claude/skills/anvien`;
   - write the manifest last.
5. Return sync counters for result summary and benchmark evidence:
   - discovered packages;
   - written new files;
   - overwritten changed files;
   - deleted obsolete files;
   - skipped unchanged files;
   - collisions or unsafe filesystem entries if encountered.
6. Manifest should describe the final desired snapshot:
   - package roots;
   - entry count;
   - file count;
   - package hash;
   - per-file hashes.

## Acceptance Criteria

1. `anvien analyze` and generated-skill setup installation propagate edits, additions, deletions, renames, missing output, and tampered output from `internal/aicontext/skills/**` to `.claude/skills/anvien/**`.
2. `.claude/skills/anvien/**` contains no stale generated files after sync, excluding `.anvien-skill-manifest.json`.
3. Sync uses hashes and skips unchanged files.
4. The manifest is rewritten to match the current source snapshot and does not preserve deleted source packages as stale entries.
5. Result summary exposes useful sync counters without using `Stale` as deleted-output preservation behavior.
6. Behavior tests use synthetic source fixtures and do not hardcode the current real skill catalog.
7. Full build passes before tests are run.
8. Focused tests for `internal/aicontext` and CLI analyze/setup paths pass.
9. `anvien detect-changes --repo Anvien --scope all` runs before commit.

## Phase Checklist

- [ ] [P0-A] Confirm plan approval and worktree hygiene.
  - Goal: ensure implementation starts only after this plan is approved and no accidental implementation patch remains.
  - Work Steps: confirm user approval; run `git status --short`; inspect any modified source files; revert only agent-created accidental implementation changes if present and approved by ownership history; leave unrelated user changes untouched.
  - Implementation Gate: user has explicitly approved this plan and worktree state is understood.
  - Acceptance: evidence ledger records approval and worktree status before implementation.

- [ ] [P0-B] Refresh graph and map blast radius.
  - Goal: establish fresh Anvien graph evidence and affected paths before code edits.
  - Work Steps: run `anvien analyze --force`; run impact for `installSkillPackagesTo`, `InstallSkillPackagesTo`, `InstallSkillPackagesForRepoTo`, `GenerateAIContextFiles`, and setup/analyze callers as needed; summarize affected files, flows, tests, and risk.
  - Implementation Gate: graph is fresh and impact output has been reviewed.
  - Acceptance: evidence ledger records impact risk, affected files, and that HIGH/CRITICAL is treated as a warning, not a blocker.

- [ ] [P1-A] Define source and output snapshot structures.
  - Goal: create explicit data structures for desired source state and actual generated output state.
  - Work Steps: inspect existing `SkillPackage`, `SkillPackageFile`, manifest structs, and hash helpers; design minimal structs or helpers for file snapshots; keep package discovery as the source snapshot input; avoid catalog-name-specific logic.
  - Implementation Gate: impact for edited symbols is complete and the design preserves the package boundary rule.
  - Acceptance: code can represent desired and actual file paths with hashes independent of the current real skill catalog.

- [ ] [P1-B] Build desired snapshot from source packages.
  - Goal: convert discovered source packages into the exact generated file set that should exist.
  - Work Steps: walk each `SkillPackage.Files`; map `file.InstallPath` to desired output path; include file content and hash; build package manifest metadata from the same source data; detect duplicate desired paths and fail explicitly.
  - Implementation Gate: source package discovery behavior remains unchanged except for snapshot projection.
  - Acceptance: desired snapshot contains every source payload file exactly once and exposes package metadata needed for the manifest.

- [ ] [P1-C] Build actual snapshot from generated output.
  - Goal: inspect `.claude/skills/anvien/**` as the current output state.
  - Work Steps: walk target root recursively; skip directories; skip `.anvien-skill-manifest.json`; hash regular files; reject or handle non-regular filesystem entries safely; normalize paths with slash separators.
  - Implementation Gate: target root is resolved through existing safe path logic and cannot escape the generated output root.
  - Acceptance: actual snapshot reports every current generated-output file path and hash without using manifest as source truth.

- [ ] [P1-D] Implement file-level diff classification.
  - Goal: classify required operations without applying them yet.
  - Work Steps: compare desired and actual snapshots; classify writes, overwrites, deletes, and skips; count operations; ensure source deletion is represented as output delete; ensure source rename becomes delete old plus write new.
  - Implementation Gate: diff code has no side effects and can be validated with synthetic snapshots.
  - Acceptance: diff output is deterministic and covers edit, add, delete, rename, tamper, missing-output, and unchanged cases.

- [ ] [P1-E] Apply sync operations safely.
  - Goal: apply the diff so target output becomes the desired mirror.
  - Work Steps: create parent directories for writes; write new/changed files; delete obsolete actual files; prune empty dirs under target root; write manifest last; surface errors with target paths.
  - Implementation Gate: operation list has been produced and target paths pass safe join checks.
  - Acceptance: after apply, rescanning actual output matches desired snapshot by path and hash.

- [ ] [P1-F] Replace stale manifest behavior.
  - Goal: remove stale-preservation semantics from generated Anvien skill output.
  - Work Steps: remove or bypass `entry.Stale = true` behavior for deleted source packages; rewrite manifest from desired snapshot only; keep stale fields only if retained for backward-compatible decoding, not for preserving deleted packages.
  - Implementation Gate: snapshot diff/apply handles deleted source packages and files.
  - Acceptance: deleted source packages do not remain in manifest and their output files are gone after sync.

- [ ] [P2-A] Wire PA4 sync into analyze/setup install paths.
  - Goal: ensure both `anvien analyze` and `anvien setup` use the same sync behavior.
  - Work Steps: route `InstallSkillPackagesTo`, `InstallSkillPackagesForRepoTo`, and internal analyze generation through the snapshot-diff installer; update `SkillInstallResult` or its `Summary()` output to report discovered packages plus write, overwrite, delete, skip, and collision counters; document any backward-compatible interpretation of existing fields.
  - Implementation Gate: call path impact has been reviewed and result-summary compatibility has an explicit decision before editing callers.
  - Acceptance: analyze and setup both produce exact generated mirror output for their selected skill source and report sync counters that match benchmark evidence.

- [ ] [P2-B] Update generated guidance policy.
  - Goal: make docs/guidance state that `.claude/skills/anvien/**` is generated mirror output.
  - Work Steps: update source guidance generator text if it currently promises preservation inside `.claude/skills/anvien`; clarify that custom skills belong outside this generated namespace; do not edit generated `AGENTS.md` or `CLAUDE.md` directly.
  - Implementation Gate: source text location is identified and generated output is not used as source.
  - Acceptance: regenerated guidance no longer describes stale preservation in the generated Anvien namespace.

- [ ] [P3-A] Add synthetic behavior tests after implementation is correct.
  - Goal: verify behavior without hardcoding the current real skill catalog.
  - Work Steps: create temp repo source under `internal/aicontext/skills/**`; run install/generation; mutate synthetic source for edit, add, delete, rename, tamper, and missing-output cases; assert generated output mirrors synthetic source after sync.
  - Implementation Gate: code behavior has been manually verified with synthetic inputs before tests are written.
  - Acceptance: tests fail against stale-preservation behavior and pass with PA4 sync behavior.

- [ ] [P3-B] Run full build before tests.
  - Goal: satisfy repository build gate before validation tests.
  - Work Steps: run `powershell -ExecutionPolicy Bypass -File anvien-launcher\build.ps1`; record result; only then run focused tests.
  - Implementation Gate: implementation code compiles far enough to attempt full build.
  - Acceptance: evidence ledger records build pass/fail and any failure handling.

- [ ] [P3-C] Run focused validation.
  - Goal: validate aicontext and CLI integration paths.
  - Work Steps: run focused Go tests for `internal/aicontext`; run relevant CLI tests for analyze/setup skill installation; add e2e only if Web UI behavior changes, which is not expected.
  - Implementation Gate: full build has completed.
  - Acceptance: evidence ledger records focused validation results and no Web UI e2e is required unless scope changes.

- [ ] [P3-D] Validate current repository generated output.
  - Goal: prove real repo output syncs with the current real source catalog.
  - Work Steps: run `anvien analyze --force`; inspect `.claude/skills/anvien/**`; confirm removed source packages/files are absent; confirm manifest package/file counts match source snapshot; confirm the target root contains no non-manifest generated payload file outside the desired snapshot; record sync counters and inventory counts.
  - Implementation Gate: build and focused tests pass or failures are explicitly resolved.
  - Acceptance: benchmark ledger records source package count, generated output file count, and sync operation counters.

- [ ] [P4-A] Final evidence, benchmark, detect-changes, and commit.
  - Goal: close the implementation slice with traceable proof.
  - Work Steps: update evidence ledger; update benchmark ledger; run `anvien detect-changes --repo Anvien --scope all`; inspect risk output; commit scoped implementation and docs changes.
  - Implementation Gate: all acceptance criteria are met or explicitly documented as blocked before commit.
  - Acceptance: final response includes commit hash, validation summary, benchmark highlights, and residual risk.

## Risk Notes

- The main behavioral change is that `.claude/skills/anvien/**` becomes a strict generated mirror. Files placed manually inside that namespace will be removed by sync if they are not present in source.
- The installer affects analyze and setup paths, so blast radius is expected to be HIGH or CRITICAL. That warning requires careful validation, not avoidance.
- Tests must not depend on a specific real catalog package such as `debugging`; the real catalog can change frequently.
- Manifest compatibility must be handled carefully so older manifests do not break sync, but old stale entries must not preserve deleted output.
