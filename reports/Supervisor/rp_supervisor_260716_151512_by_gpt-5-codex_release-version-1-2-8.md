# Supervisor Report: Release Version 1.2.8

Verdict: PASS

## Metadata

- Report file: `reports/Supervisor/rp_supervisor_260716_151512_by_gpt-5-codex_release-version-1-2-8.md`
- Review time: `260716 151512 +07:00`
- Reviewer: `gpt-5-codex`
- Repo/project: `E:\Anvien`
- Scope reviewed: commit `73612979fd9300fbc36ff47731099c8110962dab` (`chore(release): bump version to 1.2.8`)
- Claim reviewed: Anvien release metadata and changelog were bumped from `1.2.7` to `1.2.8`, with validation and change-detection evidence.
- Authority used: latest user request; `AGENTS.md` release/build/change-detection rules; source files; runtime command output.
- Related artifacts: `reports/coder/rp_coder_260716_151122_by_gpt-5-codex_release-version-1-2-8.md`; commit `73612979fd9300fbc36ff47731099c8110962dab`.

## Executive Summary

- Problem: verify whether the release version bump and changelog update are complete and safely scoped.
- Decision: PASS. The commit aligns package metadata, runtime version constant, README, RUNBOOK, and root CHANGELOG on `1.2.8`; runtime smoke and validation evidence confirm the installed CLI reports `1.2.8`.
- Required outcome: accepted.

## Source-Level Clearance Notes

- `anvien/package.json`: clear - line 3 is `1.2.8`, matching the requested release bump.
- `anvien/package-lock.json`: clear - top-level version and root package version are both `1.2.8`.
- `internal/version/version.go`: clear - line 5 sets `Version = "1.2.8"`, and `anvien version` returned `1.2.8`.
- `README.md` and `RUNBOOK.md`: clear - current visible release references are updated to `1.2.8`.
- `CHANGELOG.md`: clear - top entry is `[1.2.8] - 2026-07-16` and records the Tree-sitter drift readiness fix plus the version bump.
- Reports/notes: clear - coder report and notes log record the scope and verification evidence.

## Evidence Checked

Passed:

- Source diff: `git show --no-ext-diff --unified=3 HEAD -- CHANGELOG.md README.md RUNBOOK.md anvien/package.json anvien/package-lock.json internal/version/version.go` shows only the intended version/changelog edits.
- Runtime smoke: `anvien version` -> `1.2.8`.
- Commit evidence: commit `73612979fd9300fbc36ff47731099c8110962dab` message records `npm run full-build PASS`, `go test ./internal/version ./internal/cli ./internal/mcp PASS`, `anvien version -> 1.2.8`, and `anvien detect-changes --repo Anvien --scope all -> risk_level=low`.
- Coder report evidence: `reports/coder/rp_coder_260716_151122_by_gpt-5-codex_release-version-1-2-8.md` lists the impact prechecks, full build, focused tests, runtime smoke, and detect-changes result.
- Worktree cleanliness before review report: `git status --short` returned no output after commit `73612979`.

Failed:

- None.

Not run:

- Full build was not rerun during supervisor review because it already ran immediately before commit on the reviewed staged diff; supervisor independently rechecked source and runtime version after commit.

## Invariant Closure

- Affected invariant: current Anvien release version metadata must be consistent across npm package metadata, packaged runtime version constant, user-facing docs, and changelog.
- Sibling surfaces checked: package manifest, lockfile, Go runtime version constant, README, RUNBOOK, CHANGELOG, coder report, notes log, runtime CLI version output.
- Residual unverified same-invariant surfaces: none.

## Overall Evaluation

The release slice is acceptable. The implementation is tightly scoped to release metadata and documentation, the runtime reports `1.2.8`, the full build and focused tests passed before commit, and Anvien change detection reported low risk with no affected processes.
