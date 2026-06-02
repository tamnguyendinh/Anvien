# Anvien Test File Unresolved Separation Benchmark Ledger

Date: 2026-06-02

Status: Complete

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

## B5 - P2-B Explicit Test Unresolved Access Metrics

Status: superseded for Web production UI by B7

Source evidence: E6.

| Metric | Unit | Latest | Target | Result |
|---|---:|---:|---:|---|
| File map explicit unresolved sort options added for raw/test buckets | options | 2 | 2 | pass |
| File map default sort remains production/default-visible unresolved | available | 1 | 1 | pass |
| File detail raw unresolved toggle enabled by default | enabled_by_default | 0 | 0 | pass |
| File detail raw unresolved toggle path | available | 1 | 1 | pass |
| Raw test unresolved sample visible after explicit toggle | available | 1 | 1 | pass |
| Raw test unresolved `sourceSiteId` visible after explicit toggle | available | 1 | 1 | pass |
| Mocked browser e2e covering explicit raw/test drill-down | tests | 1 | 1 | pass |

## B6 - Final Metrics

Status: recorded

| Metric | Unit | Final | Target | Result |
|---|---:|---:|---:|---|
| Raw unresolved traceability | available | 1 | 1 | pass |
| Production/actionable unresolved count | available | 1 | 1 | pass |
| Test unresolved count | available | 1 | 1 | pass |
| Non-actionable unresolved count | available | 1 | 1 | pass |
| Unknown unresolved count | available | 1 | 1 | pass |
| Raw/default-visible risk separation | available | 1 | 1 | pass |
| Raw unresolved files | files | 591 | measured | pass |
| Default-visible unresolved files | files | 335 | measured | pass |
| Explicit test-unresolved files | files | 239 | measured | pass |
| Default top 5 hotspots that are test/e2e files | files | 0 | 0 | pass |
| Default top 5 hotspots that are production/source files | files | 5 | measured | pass |
| Raw/test top 5 hotspots that remain test/e2e files | files | 5 | traceable | pass |
| Default graph expansion of test unresolved child nodes | enabled_by_default | 0 | 0 | pass |
| CLI/API explicit test unresolved diagnostic access | available | 1 | 1 | pass |
| Web production raw/test unresolved controls visible | controls | 0 | 0 | pass |

Final graph inventory:

| Metric | Unit | Final |
|---|---:|---:|
| Files scanned | files | 819 |
| Parsed code files | files | 599 |
| Failed parses | files | 0 |
| Graph nodes | nodes | 96,858 |
| Graph relationships | relationships | 132,460 |
| File projection files | files | 819 |
| File projection dependency edges | edges | 15,932 |
| Raw unresolved files | files | 591 |
| Default-visible unresolved files | files | 335 |

Final default top hotspots:

| Rank | File | Kind | Default-visible unresolved | Raw unresolved | Risk |
|---:|---|---|---:|---:|---|
| 1 | `anvien-web/src/hooks/useAppState.local-runtime.tsx` | source | 537 | 548 | high |
| 2 | `anvien-web/src/components/GraphCanvas.tsx` | source | 451 | 451 | high |
| 3 | `internal/contracts/web_ui.go` | source | 440 | 624 | high |
| 4 | `anvien-web/src/components/FileTreePanel.tsx` | source | 424 | 435 | high |
| 5 | `anvien-web/src/hooks/useSigma.ts` | source | 407 | 431 | high |

Final raw/test top hotspots:

| Rank | File | Kind | Default-visible unresolved | Raw unresolved | Test unresolved | Risk |
|---:|---|---|---:|---:|---:|---|
| 1 | `internal/mcp/server_test.go` | test | 0 | 1445 | 1445 | low |
| 2 | `internal/cli/command_test.go` | test | 0 | 1121 | 1121 | low |
| 3 | `internal/analyze/analyze_test.go` | test | 0 | 1052 | 1052 | low |
| 4 | `internal/resolution/resolution_test.go` | test | 0 | 934 | 934 | low |
| 5 | `anvien-web/e2e/graph-orientation-labels.spec.ts` | test | 0 | 856 | 856 | low |

## B7 - Post-Closure Web Production Surface Metrics

Status: recorded

Source evidence: E10.

| Metric | Unit | Latest | Target | Result |
|---|---:|---:|---:|---|
| Web file map `raw-unresolved` sort option visible | options | 0 | 0 | pass |
| Web file map `test-unresolved` sort option visible | options | 0 | 0 | pass |
| Web file detail raw unresolved toggle visible for test files | controls | 0 | 0 | pass |
| Raw test unresolved samples rendered in Web test file detail | enabled_by_default | 0 | 0 | pass |
| Test file identity visible in Web file map/detail | available | 1 | 1 | pass |
| Test file tested-target relationship visible in Web file detail | available | 1 | 1 | pass |
| CLI/API raw/test unresolved diagnostics retained outside production Web UI | available | 1 | 1 | pass |

## B8 - Post-Closure Semantic Test Gap Suppression Metrics

Status: recorded

Source evidence: E11.

| Metric | Unit | Before E11 | Latest | Result |
|---|---:|---:|---:|---|
| Graph nodes | nodes | 96,790 | 60,878 | pass |
| Graph relationships | relationships | 132,418 | 96,510 | pass |
| Raw unresolved files | files | 591 | 352 | pass |
| Default-visible unresolved files | files | 335 | 335 | pass |
| Explicit `test-unresolved` files | files | 239 | 0 | pass |
| Test files returned by `--kind test` | files | 239 | 239 | pass |
| Sampled test-file unresolved buckets | counts | nonzero before semantic suppression | 0 | pass |
| Default top 5 hotspots that are source files | files | 5 | 5 | pass |
| Source-file linked test relationship signal | available | 1 | 1 | pass |

Semantic persistence metrics from benchmark JSON:

| Metric | Unit | Latest |
|---|---:|---:|
| Resolution gap inputs | inputs | 69,553 |
| Resolution gap nodes persisted | nodes | 33,579 |
| Resolution gap relationships persisted | relationships | 33,579 |
| Test-source resolution gap inputs skipped | inputs | 35,974 |

Latest graph/file projection inventory:

| Metric | Unit | Latest |
|---|---:|---:|
| Files scanned | files | 819 |
| Parsed code files | files | 599 |
| Failed parses | files | 0 |
| File projection files | files | 819 |
| File projection dependency edges | edges | 15,939 |
| Raw unresolved files | files | 352 |
| Default-visible unresolved files | files | 335 |

Latest default top hotspots:

| Rank | File | Kind | Default-visible unresolved | Raw unresolved | Test unresolved | Risk |
|---:|---|---|---:|---:|---:|---|
| 1 | `anvien-web/src/hooks/useAppState.local-runtime.tsx` | source | 537 | 548 | 0 | high |
| 2 | `anvien-web/src/components/GraphCanvas.tsx` | source | 451 | 451 | 0 | high |
| 3 | `internal/contracts/web_ui.go` | source | 440 | 624 | 0 | high |
| 4 | `anvien-web/src/components/FileTreePanel.tsx` | source | 424 | 435 | 0 | high |
| 5 | `anvien-web/src/hooks/useSigma.ts` | source | 407 | 431 | 0 | high |

## B9 - Retired Test-Unresolved Surface Metrics

Status: recorded

Source evidence: E12.

| Metric | Unit | Before E12 | Latest | Result |
|---|---:|---:|---:|---|
| `testUnresolvedSourceSiteCount` in Web contract | fields | 1 | 0 | pass |
| `test-unresolved` production code/contract/fixture matches | matches | present | 0 | pass |
| `test-unresolved` CLI help entries | entries | 1 | 0 | pass |
| `file-hotspots --sort test-unresolved` accepted | accepted | 1 | 0 | pass |
| `graph-health files --sort test-unresolved` accepted | accepted | 1 | 0 | pass |
| Test files returned by `--kind test` | files | 239 | 239 | pass |
| Sampled test-file unresolved/raw/default-visible counts | counts | 0 | 0 | pass |
| Default-visible unresolved files | files | 335 | 335 | pass |
| Raw unresolved files | files | 352 | 352 | pass |
| Production/default top 5 hotspots that are source files | files | 5 | 5 | pass |

Latest graph/file projection inventory:

| Metric | Unit | Latest |
|---|---:|---:|
| Files scanned | files | 821 |
| Parsed code files | files | 599 |
| Failed parses | files | 0 |
| Graph nodes | nodes | 60,911 |
| Graph relationships | relationships | 96,559 |
| ResolutionGap nodes | nodes | 33,585 |
| HAS_RESOLUTION_GAP relationships | relationships | 33,585 |
| ResolutionGap source-site count | sites | 34,218 |
| File projection dependency edges | edges | 15,947 |
| Raw unresolved files | files | 352 |
| Default-visible unresolved files | files | 335 |

Latest sampled test files:

| File | Kind | Raw unresolved | Default-visible unresolved | Risk |
|---|---|---:|---:|---|
| `anvien-launcher/server-wrapper/main_test.go` | test | 0 | 0 | low |
| `anvien-launcher/src/main_test.go` | test | 0 | 0 | low |
| `anvien-web/e2e/debug-issues.spec.ts` | test | 0 | 0 | low |
| `anvien-web/e2e/file-map-test-unresolved.spec.ts` | test | 0 | 0 | low |
| `anvien-web/e2e/graph-health-ui.spec.ts` | test | 0 | 0 | low |
