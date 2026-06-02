# Anvien Test File Unresolved Separation Benchmark Ledger

Date: 2026-06-02

Status: Active

Companion files:

- Plan: [2026-06-02-anvien-test-file-unresolved-separation-plan.md](2026-06-02-anvien-test-file-unresolved-separation-plan.md)
- Evidence ledger: [2026-06-02-anvien-test-file-unresolved-separation-evidence.md](2026-06-02-anvien-test-file-unresolved-separation-evidence.md)

## Benchmark Rules

1. Record quantitative inventory and before/after counts only.
2. Put interpretation and command context in the evidence ledger.
3. Track raw unresolved separately from default-visible unresolved.
4. Track test unresolved separately from production/actionable unresolved.
5. Track top hotspot composition before and after the implementation.
6. Build/test/e2e pass-fail belongs in evidence unless timing/count/size is the measured target.

## B0 - Baseline From Current Session

Status: recorded

Source evidence: E0.

| Metric | Unit | Baseline |
|---|---:|---:|
| Files scanned | files | 815 |
| Parsed code files | files | 598 |
| Failed parses | files | 0 |
| Graph nodes | nodes | 96,206 |
| Graph relationships | relationships | 131,705 |
| File projection files | files | 815 |
| File projection dependency edges | edges | 15,898 |
| Files with unresolved | files | 590 |
| Top 5 hotspots that are test/e2e files | files | 5 |
| Top 5 hotspots that are production files | files | 0 |

Baseline top hotspots:

| Rank | File | Kind | Unresolved | Risk |
|---:|---|---|---:|---|
| 1 | `internal/mcp/server_test.go` | test | 1445 | high |
| 2 | `internal/cli/command_test.go` | test | 1119 | high |
| 3 | `internal/analyze/analyze_test.go` | test | 1052 | high |
| 4 | `internal/resolution/resolution_test.go` | test | 934 | high |
| 5 | `anvien-web/e2e/graph-orientation-labels.spec.ts` | e2e test | 856 | high |

## B1 - Target Metrics

Status: planned

| Metric | Unit | Target |
|---|---:|---:|
| Raw unresolved traceability | available | 1 |
| Separated production/actionable unresolved count | available | 1 |
| Separated test unresolved count | available | 1 |
| Separated non-actionable unresolved count | available | 1 |
| Separated unknown unresolved count | available | 1 |
| Default top 5 hotspots dominated by test/e2e files | files | 0 |
| Default graph expansion of test unresolved child nodes | enabled_by_default | 0 |
| Explicit test unresolved drill-down/filter | available | 1 |

## B2 - Final Metrics

Status: pending

| Metric | Unit | Final | Target | Result |
|---|---:|---:|---:|---|
| Raw unresolved traceability | available | pending | 1 | pending |
| Production/actionable unresolved count | available | pending | 1 | pending |
| Test unresolved count | available | pending | 1 | pending |
| Non-actionable unresolved count | available | pending | 1 | pending |
| Unknown unresolved count | available | pending | 1 | pending |
| Default top 5 hotspots that are test/e2e files | files | pending | 0 | pending |
| Default top 5 hotspots that are production files | files | pending | measured | pending |
| Default graph expansion of test unresolved child nodes | enabled_by_default | pending | 0 | pending |
| Explicit test unresolved drill-down/filter | available | pending | 1 | pending |
