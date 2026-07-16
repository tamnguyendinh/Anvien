# Coder Report: Tree-sitter Go Module Drift

Status: READY FOR SUPERVISOR REVIEW

## Metadata

- Report file: `reports/coder/rp_coder_260716_144412_by_gpt-5-codex_tree-sitter-go-module-drift.md`
- Report time: `260716_144412 +07:00`
- Coder: `gpt-5-codex`
- Repo: `E:\Anvien`
- Scope: implement `docs/plans/2026-07-16-tree-sitter-go-module-drift`
- Git reference: implementation commit `22fee7030dae322f0589907000ca45f067383e05`; base commit `46e7e2e6a8aecadad5eee75b6be38c59e147190a`
- Plan: `docs/plans/2026-07-16-tree-sitter-go-module-drift/2026-07-16-tree-sitter-go-module-drift-plan.md`
- Evidence: `docs/plans/2026-07-16-tree-sitter-go-module-drift/2026-07-16-tree-sitter-go-module-drift-evidence.md`

## Invariant Family Map

- Family name: Tree-sitter dependency drift monitoring.
- Authority / SSOT: user request to implement the accepted plan; plan acceptance criteria; `go.mod` as current parser dependency source of truth; repo rules in `AGENTS.md`.
- Sibling surfaces checked:
  - `.github/scripts/check-tree-sitter-upgrade-readiness.py`
  - `.github/scripts/test_check_tree_sitter_upgrade_readiness.py`
  - `.github/workflows/tree-sitter-upgrade-readiness.yml`
  - `.github/dependabot.yml`
  - `go.mod`
  - `.gitignore`
  - `docs/plans/2026-07-16-tree-sitter-go-module-drift/*`
- Forbidden legacy fallback / alternate path:
  - Do not read `anvien/package.json["dependencies"]["tree-sitter"]`.
  - Do not report npm `tree-sitter@0.25` readiness as the current model.
  - Do not treat upstream Tree-sitter core `v0.26.11` as a runtime failure when no newer Go binding tag exists.
- Verify matrix:
  - Primary path: live checker command returns a Markdown Go module drift report.
  - Failure path: missing upstream metadata is classified as `UNKNOWN_FETCH_FAILED`, not hidden.
  - Actionable drift path: grammar module updates produce `GRAMMAR_UPDATE_AVAILABLE` and exit code `1`.
  - Informational path: upstream core ahead of Go binding produces `UPSTREAM_CORE_AHEAD_GO_BINDING` without making the report fail by itself.
  - Config path: workflow and Dependabot YAML parse successfully.
  - Runtime preservation path: full build and `anvien analyze --force` pass.

## Implementation Summary

- Rewrote `.github/scripts/check-tree-sitter-upgrade-readiness.py` from npm peer-dependency readiness to Go module drift detection.
- Added `.github/scripts/test_check_tree_sitter_upgrade_readiness.py` with stdlib `unittest` coverage.
- Updated `.github/workflows/tree-sitter-upgrade-readiness.yml` wording, issue title, report row parsing, and PR path filters.
- Updated `.github/dependabot.yml` with root `gomod` monitoring and Tree-sitter Go module grouping.
- Preserved `/anvien` and `/anvien-web` npm Dependabot entries while removing stale Tree-sitter npm runtime comments.
- Added `__pycache__/` to `.gitignore` after `.github/` was removed from ignore so Python test cache does not pollute Git.
- Created and maintained the standard four-file plan set under `docs/plans/2026-07-16-tree-sitter-go-module-drift/`.

## Files Changed

- `.gitignore`
- `.github/scripts/check-tree-sitter-upgrade-readiness.py`
- `.github/scripts/test_check_tree_sitter_upgrade_readiness.py`
- `.github/workflows/tree-sitter-upgrade-readiness.yml`
- `.github/dependabot.yml`
- `docs/plans/2026-07-16-tree-sitter-go-module-drift/2026-07-16-tree-sitter-go-module-drift-plan.md`
- `docs/plans/2026-07-16-tree-sitter-go-module-drift/2026-07-16-tree-sitter-go-module-drift-evidence.md`
- `docs/plans/2026-07-16-tree-sitter-go-module-drift/2026-07-16-tree-sitter-go-module-drift-benchmark.md`
- `docs/plans/2026-07-16-tree-sitter-go-module-drift/2026-07-16-tree-sitter-go-module-drift-actual-status.md`
- `reports/coder/rp_coder_260716_144412_by_gpt-5-codex_tree-sitter-go-module-drift.md`

## Verify Outputs

E2E Verification:
  [PASS] Focused tests: `python .github\scripts\test_check_tree_sitter_upgrade_readiness.py` -> `Ran 5 tests in 0.001s OK`
  [PASS] Live checker expected-drift path: wrapper around `python .github\scripts\check-tree-sitter-upgrade-readiness.py` -> exit code `1` expected, `Actionable drift rows: **2**`, `UPSTREAM_CORE_AHEAD_GO_BINDING`, no `KeyError`
  [PASS] YAML parse: PyYAML parsed `.github/workflows/tree-sitter-upgrade-readiness.yml` and `.github/dependabot.yml`
  [PASS] Full build: `npm run full-build` -> completed successfully, built runtime and web, ran `anvien analyze . --force`
  [PASS] Final graph refresh: `anvien analyze --force` -> `files: scanned=1468 parsed_code=676 failed=0`, `nodes=84126`, `relationships=122966`
  [PASS] Detect changes: `anvien detect-changes --repo Anvien --scope all` -> `risk_level=low`

## E2E Flow

Trigger -> process -> observable result:

1. Trigger: scheduled/PR workflow runs `.github/scripts/check-tree-sitter-upgrade-readiness.py`.
2. Process: checker reads `go list -m -u -json all`, filters Tree-sitter Go modules, fetches upstream Tree-sitter core latest release, classifies statuses.
3. Observable result: Markdown report lists 19 Tree-sitter modules, 2 actionable drift rows, and upstream core ahead as informational.
4. Workflow consumes the report and opens/updates a `Tree-sitter Go module drift` issue when actionable drift exists.
5. Dependabot monitors root Go modules so future parser dependency updates are surfaced as dependency PRs.

## Residual Unverified Surfaces

none

## Risks / Open Points

- The checker currently reports two real actionable module drifts:
  - `github.com/tree-sitter/tree-sitter-embedded-template` has update `v0.25.0`.
  - `github.com/UserNobody14/tree-sitter-dart` has update `v0.0.0-20260707040301-be07cf7118d3`.
- This plan intentionally does not bump those dependencies. It repairs detection and CI reporting only.
- `.github` was previously ignored; the user removed `.github/` from `.gitignore`. Final staging must include only target `.github` files for this plan unless the owner separately asks to track the full `.github` directory.
