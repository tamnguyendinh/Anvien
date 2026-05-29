# Anvien Hard Rebrand Benchmark Ledger

Date: 2026-05-29

Status: Draft

Companion files:

- Plan: [2026-05-29-anvien-rebrand-plan.md](2026-05-29-anvien-rebrand-plan.md)
- Evidence ledger: [2026-05-29-anvien-rebrand-evidence.md](2026-05-29-anvien-rebrand-evidence.md)

## Benchmark Rules

This file records quantitative data only: old-name counts, active legacy-surface counts, graph inventory counts, package/startup sizes or timings, and validation pass/fail counts.

Narrative evidence and command excerpts belong in the evidence ledger.

## B0 - Graph Baseline

Status: recorded

| Metric | Unit | Baseline | Latest | Delta | Notes |
|---|---:|---:|---:|---:|---|
| Files scanned | files | 800 | 800 | 0 | `avmatrix analyze --force` |
| Files parsed | files | 583 | 583 | 0 | `avmatrix analyze --force` |
| Unsupported files | files | 217 | 217 | 0 | `avmatrix analyze --force` |
| Failed files | files | 0 | 0 | 0 | `avmatrix analyze --force` |
| Graph nodes | nodes | 91223 | 91223 | 0 | Fresh graph |
| Graph relationships | relationships | 124702 | 124702 | 0 | Fresh graph |

## B1 - Old-Name Reference Baseline

Status: recorded

Search excluded `node_modules` and this rebrand file set.

| Pattern | Unit | Baseline | Latest | Delta | Final target |
|---|---:|---:|---:|---:|---:|
| `AVmatrix` | matches | 2238 | 2238 | 0 | 0 active |
| `avmatrix` | matches | 9291 | 9291 | 0 | 0 active |
| `AVMATRIX` | matches | 281 | 281 | 0 | 0 active |
| `AVmatrix-GO` | matches | 629 | 629 | 0 | 0 active |
| `avmatrix.com` | matches | 0 | 0 | 0 | 0 |
| `.avmatrix` | matches | 316 | 316 | 0 | 0 active |
| `AVMATRIX_` | matches | 281 | 281 | 0 | 0 active |
| `mcpServers` | matches | 9 | 9 | 0 | inspect/update keys |

## B2 - Old-Name File Group Baseline

Status: recorded

| Group | Unit | Baseline | Latest | Delta | Final target |
|---|---:|---:|---:|---:|---:|
| `internal` | files | 338 | 338 | 0 | 0 active old names |
| `reports` | files | 70 | 70 | 0 | delete/update/classify |
| `avmatrix-web` | files | 68 | 68 | 0 | rename/update |
| `docs` | files | 55 | 55 | 0 | update active docs |
| `baseline` | files | 19 | 19 | 0 | regenerate/update active baselines |
| `avmatrix-launcher` | files | 6 | 6 | 0 | rename/update |
| `avmatrix` | files | 5 | 5 | 0 | rename package |
| `cmd` | files | 5 | 5 | 0 | rename entrypoint/imports |

## B3 - Active Legacy Surface Count

Status: baseline recorded

| Surface | Unit | Baseline | Latest | Delta | Final target |
|---|---:|---:|---:|---:|---:|
| Old CLI command names accepted/generated | count | 1 | 1 | 0 | 0 |
| Old MCP server names generated | count | 1 | 1 | 0 | 0 |
| Old MCP resource schemes generated | count | 1 | 1 | 0 | 0 |
| Old repo/global storage dirs generated | count | 2 | 2 | 0 | 0 |
| Old env var prefixes read | count | 1+ | 1+ | 0 | 0 |
| Old package/bin names generated | count | 1+ | 1+ | 0 | 0 |
| Old launcher protocol/executable names generated | count | 3+ | 3+ | 0 | 0 |
| Old generated skill namespace generated | count | 1 | 1 | 0 | 0 |

## B3.1 - GitHub Automation Old-Name Baseline

Status: recorded

| Pattern | Unit | Baseline | Latest | Delta | Final target |
|---|---:|---:|---:|---:|---:|
| `.github` `AVmatrix` references | matches | 4 | 4 | 0 | 0 active |
| `.github` `avmatrix` references | matches | 94 | 94 | 0 | 0 active |
| `.github` `AVMATRIX` references | matches | 7 | 7 | 0 | 0 active |
| `.github` `AVmatrix-GO` references | matches | 0 | 0 | 0 | 0 |
| `.github` `setup-avmatrix` references | matches | 6 | 6 | 0 | 0 active |
| `.github` old GitHub URL references | matches | 0 | 0 | 0 | 0 |

## B4 - Future Runtime/Package Metrics

Status: pending implementation

| Metric | Unit | Baseline | Latest | Delta | Target |
|---|---:|---:|---:|---:|---:|
| `anvien.exe` size | bytes | pending | pending | pending | record |
| `AnvienLauncher.exe` size | bytes | pending | pending | pending | record |
| npm package tarball size | bytes | pending | pending | pending | record |
| CLI startup time | ms | pending | pending | pending | no unintended regression |
| MCP tools/list pass count | tests | pending | pending | pending | pass |
| MCP resources/list pass count | tests | pending | pending | pending | pass |
| MCP `anvien://setup` smoke count | tests | pending | pending | pending | pass |
| Web e2e Anvien branding checks | tests | pending | pending | pending | pass |
