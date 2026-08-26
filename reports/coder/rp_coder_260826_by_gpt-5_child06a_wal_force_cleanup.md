# Coder Report — Child 06A Ladybug Force Artifact Cleanup

Date: `2026-08-26`
Base Git HEAD: `c5f364978ade62effe61c34a6d1a1605de755a06`
Status: `READY_FOR_SUPERVISOR`
Disposition: uncommitted Coder handoff; staged set empty

## Root Cause

`prepareStorage(force)` removed only the main Ladybug database path and Graph JSON. `lbugruntime` modeled only `.wal` and `.lock`, while the physical database generation also owns `.wal.checkpoint`, `.checkpoint.apply.lock`, and `.checkpoint.intent.lock`. The mixed old-sidecar/new-main state caused `anvien analyze --force` to fail with `lbug_database_init failed with state 1`.

The first mandatory refresh failed with that exact state. After `.wal` and `.lock` were already absent, another refresh still failed. The three remaining checkpoint sidecars were then found at the explicit sibling paths and preserved, not deleted, under `E:\Anvien\.tmp\wal-fix-preedit-sidecars`; the next fresh refresh passed. Protected `E:\Anvien\.tmp\lbug-recovery-20260826-a003` was untouched.

## Invariant Family Map

- Authority/owner: `internal/lbugruntime` owns the explicit Ladybug base database and five-sidecar physical artifact family.
- Force reset: `internal/analyze::prepareStorage` decides when to reset and delegates full database artifact cleanup to `lbugruntime`.
- WAL recovery: remains gated only by `IsWALCorruptionError`, removes the same five sidecars, preserves the main database, and retries once.
- Forbidden fallback: no `lbug.*` glob, no native wrapper/API change, no generic `state 1` recovery classification, and no non-force lifecycle change.
- Sibling surfaces checked: force reset, database-family cleanup, WAL-corruption recovery, non-WAL error classification, and rebuilt CLI runtime.
- Legacy fallback status: incomplete two-sidecar helpers were replaced; no alternate cleanup list remains in the scoped packages.
- Residual unverified surfaces: none inside the approved artifact-family scope.

## Graph And Impact Evidence

- Fresh `anvien analyze --force`: PASS after preserving the three stale checkpoint sidecars; indexed commit matches HEAD and graph is fresh.
- `internal/analyze/analyze.go`: file risk `HIGH`; `prepareStorage` impact reaches `Run`, CLI analyze, and access-candidate audit paths. Change kept to one delegated call.
- `internal/lbugruntime/wal_recovery.go`: file risk `HIGH`, fan-in `21`; direct storage blast radius reaches `RunWithWALRecovery` and `lbugnative.OpenWriteRunnerWithEmbeddingDims`. Symbol-level impact for the sidecar path chain is `LOW` across two affected production files.
- `TestRunForceRemovesPreviousLbugOutput` and `TestRunWithWALRecoveryRemovesSidecarsAndRetries`: symbol impact `LOW`, no upstream impacted symbols.

## Exact Diff

- `internal/lbugruntime/wal_recovery.go`: adds explicit `DatabaseSidecarPaths`, `RemoveDatabaseSidecars`, and `RemoveDatabaseArtifacts`; recovery reuses sidecar-only cleanup.
- `internal/lbugruntime/extensions_retry_test.go`: locks the exact five-path contract; proves full reset removes base plus all sidecars; proves WAL recovery removes all sidecars while preserving the base DB; explicitly keeps generic native `state 1` non-WAL.
- `internal/analyze/analyze.go::prepareStorage`: delegates force database reset to `lbugruntime.RemoveDatabaseArtifacts` and keeps Graph JSON behavior unchanged.
- `internal/analyze/analyze_test.go::TestRunForceRemovesPreviousLbugOutput`: seeds and asserts all five sidecars while retaining main DB and Graph JSON assertions.
- Scoped diff: `4 files changed, 81 insertions(+), 11 deletions(-)`; `gofmt` complete; scoped `git diff --check` PASS.

## Verification

- Holder preflight: `analyze.lock` free. The only reported runtime process was the VS Code Playwright test server; no proven build holder was terminated.
- Canonical full build on final tree: PASS, exit `0`; runtime version `1.2.8`; `anvien.exe` SHA-256 `EAC01B85E54FFA76502E122FDBF99C5221D4AE22D1CA823EAA48415D09A54A58`; build-owned `analyze . --force` PASS.
- `go test ./internal/lbugruntime -run '^(TestIsWALCorruptionErrorMatchesLegacyRecoverySignals|TestDatabaseSidecarPaths|TestRemoveDatabaseArtifactsRemovesDatabaseAndSidecars|TestRunWithWALRecoveryRemovesSidecarsAndRetries|TestRunWithWALRecoveryDoesNotRetryNonWALErrors)$' -count=1 -v`: PASS.
- `go test ./internal/analyze -run '^TestRunForceRemovesPreviousLbugOutput$' -count=1 -v`: PASS.
- Rebuilt-binary runtime repro at `E:\Anvien\.tmp\lbug-force-coder-validation`: seeded main `lbug` plus all five explicit stale sidecars; `anvien.exe analyze <root> --force --skip-git` exited `0`, reported `scanned=0`, produced main DB/Graph JSON/meta, and left all five seeded sidecars absent. Generated agent/publication files appeared only after successful analysis and are not input source files or performance evidence.

E2E Verification:
  [PASS] Compiled: canonical `scripts/full-build.ps1` -> exit `0`
  [PASS] Runtime: stale five-sidecar generation -> force reset -> new DB/graph/meta, no stale sidecar
  [PASS] Happy path: explicit full-generation cleanup -> base DB and five sidecars removed
  [PASS] Edge case: WAL-classified error removes five sidecars but preserves main DB; generic `state 1` remains non-WAL

## Preserve Boundary And Residual

- Preserved the paused A004 report and all other lane-owned plan/report changes; none were edited, staged, deleted, or included.
- Preserved `E:\Anvien\.tmp\lbug-recovery-20260826-a003` and the original three checkpoint sidecars in `E:\Anvien\.tmp\wal-fix-preedit-sidecars`.
- Native state-1 diagnostic/error plumbing remains an explicit residual. This fix does not interpret state `1`, change native handles, or broaden WAL recovery.
- No detect, stage, commit, target benchmark/profile, A003 rerun, or A004 work was performed.

Next owner: Main opens a fresh Supervisor review of this exact uncommitted four-file implementation plus this report.

`CODER_CHILD06A_WAL_FORCE_FIX_READY_FOR_SUPERVISOR`
