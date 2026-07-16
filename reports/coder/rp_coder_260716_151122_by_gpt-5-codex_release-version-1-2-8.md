# Coder Report: Release Version 1.2.8

Status: READY FOR SUPERVISOR REVIEW

## Metadata

- Report file: `reports/coder/rp_coder_260716_151122_by_gpt-5-codex_release-version-1-2-8.md`
- Report time: `260716_151122 +07:00`
- Coder: `gpt-5-codex`
- Repo: `E:\Anvien`
- Scope: bump Anvien CLI/runtime package version from `1.2.7` to `1.2.8` and update release changelog.
- Git reference: pending release commit from base `e1ec17886ee00c5834d69b6b8c8a929af68b650f`.

## Invariant Family Map

- Family name: Anvien release version metadata.
- Authority / SSOT: user request to bump version and update changelog; `anvien/package.json`; `internal/version/version.go`; repo rules in `AGENTS.md`.
- Sibling surfaces checked:
  - `anvien/package.json`
  - `anvien/package-lock.json`
  - `internal/version/version.go`
  - `README.md`
  - `RUNBOOK.md`
  - `CHANGELOG.md`
- Forbidden legacy fallback / alternate path:
  - Do not leave the package manifest, Go runtime version constant, README, or runbook split across different current versions.
  - Do not rewrite historical report/plan evidence that correctly records older `1.2.7` runs.
  - Do not treat unrelated dependency version `11.2.7` as an Anvien release version.
- Verify matrix:
  - Package path: npm manifest and lockfile report `1.2.8`.
  - Runtime path: installed `anvien version` prints `1.2.8`.
  - CLI/API/MCP import path: version-adjacent Go package tests pass.
  - Release docs path: README, RUNBOOK, and CHANGELOG show current release `1.2.8`.
  - Graph/change path: full build refreshes graph and `detect-changes` reports low risk.

## Implementation Summary

- Bumped `anvien/package.json` and `anvien/package-lock.json` from `1.2.7` to `1.2.8`.
- Bumped `internal/version/version.go` runtime constant from `1.2.7` to `1.2.8`.
- Updated README and RUNBOOK current version references to `1.2.8`.
- Added `CHANGELOG.md` entry for `1.2.8` dated `2026-07-16`, covering the Tree-sitter Go module drift readiness fix and the version bump.

## Files Changed

- `CHANGELOG.md`
- `README.md`
- `RUNBOOK.md`
- `anvien/package-lock.json`
- `anvien/package.json`
- `internal/version/version.go`
- `reports/coder/rp_coder_260716_151122_by_gpt-5-codex_release-version-1-2-8.md`
- `docs/notes_decisions_log/notes_decisions_log_20260716.md`

## Verify Outputs

E2E Verification:
  [PASS] Impact precheck: `anvien file-detail internal/version/version.go --repo Anvien --json` + `anvien impact file internal/version/version.go --repo Anvien --direction upstream` -> file-level blast radius `CRITICAL`, 6 affected files, 2 linked tests; treated as scope warning and handled surgically.
  [PASS] Symbol impact precheck: `anvien impact symbol "Version" --uid "Const:internal/version/version.go:Version" --repo Anvien --direction upstream` -> risk `LOW`, impactedCount `0`.
  [PASS] Package manifest precheck: `anvien file-detail anvien/package.json --repo Anvien --json` + `anvien impact file anvien/package.json --repo Anvien --direction upstream` -> risk `LOW`, impactedCount `0`.
  [PASS] Full build: `npm run full-build` -> completed successfully, installed `anvien@1.2.8`, built runtime and Web UI, built launcher, printed `anvien version` as `1.2.8`, and ran `anvien analyze . --force`.
  [PASS] Focused tests: `go test ./internal/version ./internal/cli ./internal/mcp` -> all packages passed.
  [PASS] Runtime version smoke: `anvien version` -> `1.2.8`.
  [PASS] Detect changes: `anvien detect-changes --repo Anvien --scope all` -> `risk_level=low`, changed files scoped to version/config/docs.

## E2E Flow

Trigger -> process -> observable result:

1. Trigger: release metadata bump request.
2. Process: package manifest, lockfile, Go runtime version constant, README, RUNBOOK, and root CHANGELOG are updated together.
3. Observable result: full build installs the local package and `anvien version` prints `1.2.8`.
4. Change detection confirms the diff is low-risk and scoped to release metadata/runtime version constant.

## Residual Unverified Surfaces

none

## Risks / Open Points

- `npm run full-build` emitted npm `allow-scripts` and Vite chunk-size/import warnings, but no `EALLOWSCRIPTS` failure occurred and the build exited successfully.
- `anvien/CHANGELOG.md` remains untouched because it is not aligned with the current root release series and was not the active `1.2.7` changelog owner for this bump.
