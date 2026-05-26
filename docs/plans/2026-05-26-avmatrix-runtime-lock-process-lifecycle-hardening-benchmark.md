# AVmatrix Runtime Lock And Process Lifecycle Hardening Benchmark Ledger

Date: 2026-05-26

Status: Planned

Companion files:

- Plan: [2026-05-26-avmatrix-runtime-lock-process-lifecycle-hardening-plan.md](2026-05-26-avmatrix-runtime-lock-process-lifecycle-hardening-plan.md)
- Evidence ledger: [2026-05-26-avmatrix-runtime-lock-process-lifecycle-hardening-evidence.md](2026-05-26-avmatrix-runtime-lock-process-lifecycle-hardening-evidence.md)

## Benchmark Rules

This file records quantitative data only: lock metadata field counts, recovery pass/fail counts, process inventory counts, command inventory counts, diagnostics output counts, test pass/fail counts, and measured runtime behavior.

Narrative evidence, commands, logs, and source observations belong in the evidence ledger.

Use `pending` only when a future phase has not measured that value yet.

## B0 - Lock Capability Inventory

Status: baseline recorded; final pending.

| Metric | Unit | Baseline | Latest | Delta | Final target |
|---|---:|---:|---:|---:|---:|
| Lock metadata fields written | fields | 2 | 2 | 0 | >= 6 |
| Metadata includes `pid` | boolean | 1 | 1 | 0 | 1 |
| Metadata includes `acquiredAt` | boolean | 1 | 1 | 0 | 1 |
| Metadata includes `version` | boolean | 0 | 0 | 0 | 1 |
| Metadata includes `host` | boolean | 0 | 0 | 0 | 1 |
| Metadata includes `command` | boolean | 0 | 0 | 0 | 1 |
| Metadata includes ownership `token` | boolean | 0 | 0 | 0 | 1 |
| Existing lock metadata parser | count | 0 | 0 | 0 | >= 1 |
| PID liveness checks during acquire | count | 0 | 0 | 0 | >= 1 |
| Same-host dead-PID stale recovery paths | count | 0 | 0 | 0 | >= 1 |
| Foreign-host safety checks | count | 0 | 0 | 0 | >= 1 |
| Token-safe release checks | count | 0 | 0 | 0 | >= 1 |

## B1 - Stale Lock Recovery Metrics

Status: baseline recorded; final pending.

| Scenario | Baseline behavior | Final behavior target |
|---|---:|---:|
| Live same-host PID lock blocks second writer | pass | pass |
| Dead same-host PID lock recovers automatically | fail | pass |
| Old-format dead PID lock recovers automatically | fail | pass |
| Malformed fresh lock is not removed immediately | pending | pass |
| Malformed stale lock recovers under documented policy | fail | pass |
| Foreign-host lock is not removed by local PID check | pending | pass |
| Token mismatch release does not delete new owner lock | fail | pass |

## B2 - Runtime Process Inventory

Status: baseline recorded from local trace; final pending.

| Metric | Unit | Baseline | Latest | Delta | Final target |
|---|---:|---:|---:|---:|---:|
| Live `avmatrix.exe` processes observed during baseline trace | processes | 1 | 1 | 0 | record |
| Live editor-owned `avmatrix mcp` processes observed | processes | 1 | 1 | 0 | record |
| Live `avmatrix analyze` processes observed after analyze completed | processes | 0 | 0 | 0 | 0 |
| Live launcher-owned `serve --port 4848` processes observed | processes | 0 | 0 | 0 | record |
| Process classes discoverable through AVmatrix diagnostics | classes | 0 | 0 | 0 | >= 5 |
| Parent process chain levels captured in diagnostics | levels | 0 | 0 | 0 | >= 2 where OS supports it |
| Diagnostics JSON output modes | modes | 0 | 0 | 0 | >= 1 |

## B3 - User-Facing Error And Diagnostics Inventory

Status: baseline recorded; final pending.

| Metric | Unit | Baseline | Latest | Delta | Final target |
|---|---:|---:|---:|---:|---:|
| Lock conflict error includes PID | boolean | 0 | 0 | 0 | 1 |
| Lock conflict error includes command | boolean | 0 | 0 | 0 | 1 |
| Lock conflict error includes lock path | boolean | 0 | 0 | 0 | 1 |
| Lock conflict error includes age | boolean | 0 | 0 | 0 | 1 |
| Lock conflict error includes live/stale classification | boolean | 0 | 0 | 0 | 1 |
| Commands that report lock status | commands | 0 | 0 | 0 | >= 1 |
| Commands that report AVmatrix process status | commands | 0 | 0 | 0 | >= 1 |
| Setup docs/guidance mentions editor-owned MCP lifecycle | checked surfaces | 0 | 0 | 0 | record |

## B4 - Test Coverage Metrics

Status: final pending.

| Metric | Unit | Baseline | Final | Delta | Target |
|---|---:|---:|---:|---:|---:|
| `internal/repo` lock tests for stale recovery | tests | 0 | pending | pending | >= 4 |
| Token-safe release tests | tests | 0 | pending | pending | >= 1 |
| HTTP analyze stale-lock tests | tests | 0 | pending | pending | >= 1 |
| HTTP embed stale-lock tests | tests | 0 | pending | pending | >= 1 |
| CLI diagnostics tests | tests | 0 | pending | pending | >= 2 if CLI diagnostics added |
| Launcher MCP non-kill tests | tests | pending | pending | pending | >= 1 if launcher touched |
| Focused test packages passing | packages | pending | pending | pending | record |

## B5 - Validation Metrics

Status: final pending.

| Metric | Unit | Baseline | Final | Delta | Target |
|---|---:|---:|---:|---:|---:|
| Full build gate result | pass/fail | pending | pending | pending | pass |
| Stale lock smoke recovery result | pass/fail | fail | pending | pending | pass |
| Live lock conflict smoke result | pass/fail | pending | pending | pending | pass |
| Diagnostics table smoke result | pass/fail | pending | pending | pending | pass if diagnostics added |
| Diagnostics JSON smoke result | pass/fail | pending | pending | pending | pass if diagnostics added |
| `detect-changes` pre-commit result | pass/fail | pending | pending | pending | pass |
