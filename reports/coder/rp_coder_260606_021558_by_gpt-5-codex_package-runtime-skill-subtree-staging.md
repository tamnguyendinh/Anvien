# Coder Report: Package Runtime Skill Subtree Staging

Status: READY FOR SUPERVISOR REVIEW

## Metadata
- Report file: `reports/coder/rp_coder_260606_021558_by_gpt-5-codex_package-runtime-skill-subtree-staging.md`
- Time: 260606 021558 Asia/Bangkok
- Coder: gpt-5-codex
- Scope: implement `docs/plans/2026-06-06-anvien-package-runtime-skill-subtree-staging/2026-06-06-anvien-package-runtime-skill-subtree-staging-plan.md`
- Authority: `AGENTS.md`, `.agents/skills/coder/SKILL.md`, package-runtime source staging evidence, current repo evidence
- Implementation checkpoint: pending until git checkpoint is created.

## Invariant Family Map
- Family name: package runtime Go-source fallback staging for bundled Anvien runtime.
- SSOT / authority source: package source staging in `internal/cli/package_runtime.go`; skill package subtree architecture in `internal/aicontext/skill_packages.go`; plan/evidence/benchmark ledgers under `docs/plans/2026-06-06-anvien-package-runtime-skill-subtree-staging`.
- Sibling runtime surfaces checked: `prepareGoSourcePackage`, `copyPackageGoDir`, `copyPackageFileIfExists`, package command entrypoint, package command tests, package prepare smoke, source manifest inventory, full build lifecycle.
- Forbidden fallback / alternate path: no extension-based skill payload filtering and no deleting the old `.md` branch without staging the skill subtree through a separate subtree copy path.
- Stale tests/helpers/plans updated: `TestPrepareGoSourcePackageCopiesMinimalGoSource`, new idempotent copy test, plan/evidence/benchmark ledgers.
- Verify matrix: full build sequence, targeted Go tests, repo-local package prepare smoke, manifest inventory, Anvien detect-changes.

## Files Changed
- `internal/cli/package_runtime.go`: separated skill subtree staging from Go-source-only staging; added manifest path inventory; made identical native runtime destination copies idempotent.
- `internal/cli/package_command_test.go`: added non-`.md` skill subtree coverage, manifest path assertions, `_test.go` exclusion assertions, and idempotent destination-copy coverage.
- `docs/plans/2026-06-06-anvien-package-runtime-skill-subtree-staging/*`: recorded plan closure, validation evidence, and inventory benchmark.
- `reports/coder/rp_coder_260606_021558_by_gpt-5-codex_package-runtime-skill-subtree-staging.md`: handoff report.
- `docs/notes_decisions_log/notes_decisions_log_20260606.md`: report link and verification summary.

## Verify Outputs
| Command | Result |
|---|---|
| Full build sequence from repo root | Pass on fail-fast rerun. Version checks returned `1.2.5`; final analyze passed with 1402 files, 84437 nodes, 122926 relationships, 16570 dependency edges, 430 unresolved. |
| Final pre-commit `anvien analyze . --force` | Pass after report/notes docs were added. Final graph: 1404 files, 84448 nodes, 122937 relationships, 16570 dependency edges, 430 unresolved. |
| `go test ./internal/cli -count=1` | Pass. `ok github.com/tamnguyendinh/anvien/internal/cli 64.427s`. |
| Repo-local `go run ..\cmd\anvien package prepare-go-source` smoke | Pass. Copied 929 files to `E:\Anvien\anvien\go-src`; cleanup removed `go-src`. |
| Repo-local staged inventory | Pass. Skill files staged: 638; non-`.md` skill files staged: 295; manifest files: 929; manifest skill paths: 638. |
| `anvien detect-changes --repo Anvien --scope all` | Pass. Summary risk `medium`; affected process `BuildGoRuntimePackage -> CopyPackageFileIfExists`. |

## E2E Flow
- Not applicable. This slice changes CLI/package runtime staging behavior and no UI behavior.

## Closure Notes
- The `.md`-only embedded skill branch is removed from `copyPackageGoDir`.
- `internal/aicontext/skills` is staged as a complete subtree artifact and includes non-`.md` payloads.
- The native runtime DLL copy failure found during full build is closed by identical-byte idempotency.
- Residual unverified surfaces: none.
- Risks/open points: direct Supervisor acceptance is still required before this scope is DONE.
