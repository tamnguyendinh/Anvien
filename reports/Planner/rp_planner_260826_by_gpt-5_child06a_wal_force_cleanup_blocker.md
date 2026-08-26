# Planner Report — Child 06A WAL Force Cleanup Blocker

Date: `2026-08-26`
Lane: bounded Planner refresh
Verdict: `PLANNER_CHILD06A_WAL_FORCE_FIX_READY_FOR_CODER`
Next owner: Main Orchestration

## Current Cursor

`A003_CHECKPOINT_COMPLETE / A004_ARCHITECT_PAUSED_WAL_FORCE_BUG / WAL_FIX_READY_FOR_CODER / D001_STREAK_0`

A003 remains closed at implementation checkpoint `b6bf45bce95323aa6b53b182edfea8628bd8b463` and docs sync `2ea8c9486ad629f86f36fc5f4ce17574fdcb65c8`. No A003 gate, metric, disposition, checkbox, or queue is changed. Parent/D001 remain unchecked; D002-D017 remain queued. A004 task `01a03e09-77e0-7c03-b144-bed526838f37` is Owner-paused.

## Root Cause and Fix Contract

Current `internal/analyze/analyze.go::prepareStorage` force cleanup removes the main `lbug` database and Graph JSON but leaves sibling `lbug.wal`/`lbug.lock`. The exact stale-WAL repro fails `lbug_database_init failed with state 1`; the control without stale WAL succeeds.

- Production only: after removing `paths.LbugPath`, call existing `lbugruntime.RemoveWALSidecars(paths.LbugPath)` and propagate error.
- Test only after production: extend `internal/analyze/analyze_test.go::TestRunForceRemovesPreviousLbugOutput` to seed `.wal` and `.lock`, assert both are removed, and preserve existing main DB/Graph JSON assertions.
- Do not change native wrappers, native error plumbing, `RunWithWALRecovery` patterns, public contracts, non-force behavior, or A004 source. Generic state-1 diagnostics remain a later residual.

## Coder Validation

Run: `anvien --help` -> fresh healthy-index `anvien analyze --force` -> file-detail for both files -> impact for `prepareStorage` and exact test symbol -> production first -> test -> gofmt/scoped diff-check -> holder preflight -> canonical `scripts/full-build.ps1` -> focused test -> exact repo-local stale-WAL repro with rebuilt canonical runtime -> report and stop.

No target benchmark/profile/A003 rerun. STOP if another production owner, native API/error plumbing, public contract, non-force behavior, or A004 source is required. Preserve `E:\Anvien\.tmp\lbug-recovery-20260826-a003`; clean repro/control debug roots only after acceptance. A later Supervisor reviews only this bug fix and exact repro/full-build behavior. A004 remains paused until Supervisor `PASS` and accepted commit.

`PLANNER_CHILD06A_WAL_FORCE_FIX_READY_FOR_CODER`
