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

Source evidence: E1.

| Metric | Unit | Baseline |
|---|---:|---:|
| Files scanned | files | 818 |
| Parsed code files | files | 598 |
| Failed parses | files | 0 |
| Graph nodes | nodes | 96,340 |
| Graph relationships | relationships | 131,828 |
| File projection files | files | 818 |
| File projection dependency edges | edges | 15,902 |
| Files with unresolved | files | 590 |
| Top 5 hotspots that are test/e2e files | files | 5 |
| Top 5 hotspots that are production files | files | 0 |

Baseline top hotspots:

| Rank | File | Kind | Unresolved | Risk |
|---:|---|---|---:|---|
| 1 | `internal/mcp/server_test.go` | test | 1445 | high |
| 2 | `internal/cli/command_test.go` | test | 1121 | high |
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
| Raw/default-visible risk separation | available | 1 |
| Default top 5 hotspots dominated by test/e2e files | files | 0 |
| Default graph expansion of test unresolved child nodes | enabled_by_default | 0 |
| Explicit test unresolved drill-down/filter | available | 1 |

## B2 - P1-B Bucket Field Availability

Status: recorded

Source evidence: E3.

| Metric | Unit | Latest | Target | Result |
|---|---:|---:|---:|---|
| Raw unresolved traceability | available | 1 | 1 | pass |
| Production/actionable unresolved count | available | 1 | 1 | pass |
| Test unresolved count | available | 1 | 1 | pass |
| Non-actionable unresolved count | available | 1 | 1 | pass |
| Unknown unresolved count | available | 1 | 1 | pass |
| Default-visible unresolved count | available | 1 | 1 | pass |
| Raw/default-visible risk separation | available | 1 | 1 | pass |

P1-B field inventory:

| Field | Surface |
|---|---|
| `unresolvedSourceSiteCount` | raw compatibility count |
| `rawUnresolvedSourceSiteCount` | raw traceability count |
| `productionUnresolvedSourceSiteCount` | production/actionable bucket |
| `testUnresolvedSourceSiteCount` | test-file bucket |
| `nonActionableUnresolvedSourceSiteCount` | non-actionable bucket |
| `unknownUnresolvedSourceSiteCount` | unknown metadata bucket |
| `defaultVisibleUnresolvedSourceSiteCount` | default-visible bucket |
| `rawRisk` | raw risk |
| `defaultVisibleRisk` | default-visible risk |

P1-B graph inventory after implementation:

| Metric | Unit | Latest |
|---|---:|---:|
| Files scanned | files | 818 |
| Parsed code files | files | 598 |
| Failed parses | files | 0 |
| Graph nodes | nodes | 96,521 |
| Graph relationships | relationships | 132,071 |
| File projection files | files | 818 |
| File projection dependency edges | edges | 15,918 |
| Files with unresolved | files | 590 |
| Default top 5 hotspots that are test/e2e files | files | 5 |

## B3 - P1-C Hotspot Ranking Metrics

Status: recorded

Source evidence: E4.

P1-C graph inventory:

| Metric | Unit | Latest |
|---|---:|---:|
| Files scanned | files | 818 |
| Parsed code files | files | 598 |
| Failed parses | files | 0 |
| Graph nodes | nodes | 96,594 |
| Graph relationships | relationships | 132,208 |
| File projection files | files | 818 |
| File projection dependency edges | edges | 15,929 |
| Raw unresolved files | files | 590 |
| Default-visible unresolved files | files | 335 |
| Test-unresolved files from explicit `test-unresolved` view | files | 238 |

P1-C hotspot composition:

| Metric | Unit | Baseline | Latest | Target | Result |
|---|---:|---:|---:|---:|---|
| Default top 5 hotspots that are test/e2e files | files | 5 | 0 | 0 | pass |
| Default top 5 hotspots that are non-test files | files | 0 | 5 | measured | pass |
| Raw top 5 hotspots that are test/e2e files | files | 5 | 5 | traceable | pass |
| Explicit raw/test unresolved view available in command/API sort mode | available | 0 | 1 | 1 | pass |

P1-C default top hotspots:

| Rank | File | Kind | Default-visible unresolved | Raw unresolved | Risk |
|---:|---|---|---:|---:|---|
| 1 | `anvien-web/src/hooks/useAppState.local-runtime.tsx` | source | 537 | 548 | high |
| 2 | `anvien-web/src/components/GraphCanvas.tsx` | source | 451 | 451 | high |
| 3 | `internal/contracts/web_ui.go` | source | 440 | 624 | high |
| 4 | `anvien-web/src/components/FileTreePanel.tsx` | source | 424 | 435 | high |
| 5 | `anvien-web/src/hooks/useSigma.ts` | source | 407 | 431 | high |

P1-C raw/test top hotspots:

| Rank | File | Kind | Default-visible unresolved | Raw unresolved | Risk |
|---:|---|---|---:|---:|---|
| 1 | `internal/mcp/server_test.go` | test | 0 | 1445 | low |
| 2 | `internal/cli/command_test.go` | test | 0 | 1121 | low |
| 3 | `internal/analyze/analyze_test.go` | test | 0 | 1052 | low |
| 4 | `internal/resolution/resolution_test.go` | test | 0 | 934 | low |
| 5 | `anvien-web/e2e/graph-orientation-labels.spec.ts` | e2e test | 0 | 856 | low |

## B4 - P2-A Web Default Visibility Metrics

Status: recorded

Source evidence: E5.

| Metric | Unit | Latest | Target | Result |
|---|---:|---:|---:|---|
| File map default unresolved count uses default-visible bucket | available | 1 | 1 | pass |
| File detail default unresolved count uses default-visible bucket | available | 1 | 1 | pass |
| Test file identity visible in Web file map/detail | available | 1 | 1 | pass |
| Raw test unresolved samples rendered by default in file detail | enabled_by_default | 0 | 0 | pass |
| Test file tested-target relationship visible in file detail | available | 1 | 1 | pass |
| Test/non-actionable ResolutionGap nodes visible in graph semantic defaults | enabled_by_default | 0 | 0 | pass |
| Explicit graph filter path can re-enable hidden test/non-actionable gaps | available | 1 | 1 | pass |
| Mocked browser e2e covering file map/detail default behavior | tests | 1 | 1 | pass |

## B5 - Final Metrics

Status: pending

| Metric | Unit | Final | Target | Result |
|---|---:|---:|---:|---|
| Raw unresolved traceability | available | pending | 1 | pending |
| Production/actionable unresolved count | available | pending | 1 | pending |
| Test unresolved count | available | pending | 1 | pending |
| Non-actionable unresolved count | available | pending | 1 | pending |
| Unknown unresolved count | available | pending | 1 | pending |
| Raw/default-visible risk separation | available | pending | 1 | pending |
| Default top 5 hotspots that are test/e2e files | files | pending | 0 | pending |
| Default top 5 hotspots that are production files | files | pending | measured | pending |
| Default graph expansion of test unresolved child nodes | enabled_by_default | pending | 0 | pending |
| Explicit test unresolved drill-down/filter | available | pending | 1 | pending |
