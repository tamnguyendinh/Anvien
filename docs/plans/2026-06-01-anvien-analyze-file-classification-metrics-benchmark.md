# Anvien Analyze File Classification Metrics Benchmark Ledger

Date: 2026-06-01

Status: Planned

Companion files:

- Plan: [2026-06-01-anvien-analyze-file-classification-metrics-plan.md](2026-06-01-anvien-analyze-file-classification-metrics-plan.md)
- Evidence ledger: [2026-06-01-anvien-analyze-file-classification-metrics-evidence.md](2026-06-01-anvien-analyze-file-classification-metrics-evidence.md)

## Benchmark Rules

1. Record quantitative data only.
2. Put command interpretation in the evidence ledger.
3. Preserve current, target, and delta where useful.
4. Track analyze inventory counts after every implementation slice.
5. Track classification bucket counts after every implementation slice.
6. Build/test pass/fail belongs in the evidence ledger unless timing/count/size is the measured target.

## B0 - Current Analyze Inventory Baseline

Status: recorded

Source evidence: E0 and E6.

| Metric | Unit | Initial | Latest | Delta | Target |
|---|---:|---:|---:|---:|---:|
| Files scanned | files | 807 | 813 | +6 | record |
| Files parsed | files | 596 | 596 | 0 | no unintended decrease |
| Current aggregate unsupported | files | 211 | 217 | +6 | 0 user-facing aggregate |
| Files failed | files | 0 | 0 | 0 | 0 |
| Graph nodes | nodes | 95814 | 95889 | +75 | record |
| Graph relationships | relationships | 131151 | 131238 | +87 | record |
| File projection files | files | 807 | 813 | +6 | record |
| File projection dependency edges | edges | 15816 | 15828 | +12 | record |
| File projection unresolved files | files | 588 | 588 | 0 | record |

## B1 - Initial 211-File Bucket Baseline

Status: recorded

| Current bucket | Unit | Count |
|---|---:|---:|
| `document:markdown` | files | 105 |
| `document:spreadsheet` | files | 1 |
| `.json` | files | 92 |
| `.mod` | files | 3 |
| `.html` | files | 2 |
| `.ps1` | files | 2 |
| `.cli` | files | 1 |
| `.conf` | files | 1 |
| `.css` | files | 1 |
| `.sh` | files | 1 |
| `.web` | files | 1 |
| `.yaml` | files | 1 |
| Total | files | 211 |

## B2 - Latest 217-File Bucket Snapshot

Status: recorded

| Current bucket | Unit | Count |
|---|---:|---:|
| `document:markdown` | files | 111 |
| `document:spreadsheet` | files | 1 |
| `.json` | files | 92 |
| `.mod` | files | 3 |
| `.html` | files | 2 |
| `.ps1` | files | 2 |
| `.cli` | files | 1 |
| `.conf` | files | 1 |
| `.css` | files | 1 |
| `.sh` | files | 1 |
| `.web` | files | 1 |
| `.yaml` | files | 1 |
| Total | files | 217 |

## B3 - Target Classification Baseline For Latest Inventory

Status: target defined

| Target bucket | Unit | Latest equivalent | Target after implementation |
|---|---:|---:|---:|
| Parsed code files | files | 596 | 596 |
| Document files | files | 112 | 112 |
| Metadata/config/report/fixture files | files | 99 | 99 |
| Script files without ScopeIR extractor | files | 3 | 3 |
| Static Web/assets | files | 3 | 3 |
| True unsupported analyzer inputs | files | not separated | 0 unless real unsupported code inputs exist |
| Unknown/unclassified files | files | not separated | 0 for current repo baseline |
| Failed files | files | 0 | 0 |
| Sum | files | 813 | 813 |

## B4 - Initial Directory Concentration Baseline

Status: recorded

| Top path | Unit | Count |
|---|---:|---:|
| `reports` | files | 79 |
| `docs` | files | 62 |
| `internal` | files | 28 |
| `anvien-claude-plugin` | files | 14 |
| `anvien-web` | files | 7 |
| `anvien` | files | 3 |
| `anvien-launcher` | files | 3 |
| `scripts` | files | 2 |
| remaining root/singleton paths | files | 13 |
| Total | files | 211 |

## B5 - Implementation Measurement Targets

Status: pending

| Metric | Unit | Baseline | Latest | Target |
|---|---:|---:|---:|---:|
| Human analyze output aggregate `unsupported` | occurrences | 1 | pending | 0 |
| Human analyze output causal file buckets | buckets | 0 | pending | at least 6 |
| JSON causal file bucket fields | fields | 0 | pending | at least 7 |
| Unknown/unclassified current repo files | files | not separated | pending | 0 |
| True unsupported current repo files | files | not separated | pending | 0 unless real unsupported code inputs are found |
| Failed files | files | 0 | pending | 0 |
| Analyze scanned files | files | 813 | pending | record |
| Analyze parsed code files | files | 596 | pending | no unintended decrease |

## B6 - Validation Runs

Status: pending

| Run | Unit | Latest | Target |
|---|---:|---:|---:|
| Full build | result | pending | pass |
| Targeted Go tests | result | pending | pass |
| Analyze smoke | result | pending | pass |
| Graph-health summary | result | pending | pass |
| Detect changes | result | pending | recorded before implementation commit |
