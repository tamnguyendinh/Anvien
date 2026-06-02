# Anvien Deadcode Agent Token Benchmark Ledger

Date: 2026-06-02

Status: Not started

Companion files:

- Plan: [2026-06-02-anvien-deadcode-agent-token-benchmark-plan.md](2026-06-02-anvien-deadcode-agent-token-benchmark-plan.md)
- Evidence ledger: [2026-06-02-anvien-deadcode-agent-token-benchmark-evidence.md](2026-06-02-anvien-deadcode-agent-token-benchmark-evidence.md)

## Benchmark Rules

1. Record quantitative data only.
2. Put command details, reasoning, and candidate proof in the evidence ledger.
3. Use `estimated_tokens = ceil(characters / 4)` when exact token accounting is unavailable.
4. Separate discovery cost from shared verification cost.
5. Separate native-search and Anvien-guided metrics.
6. Record both totals and token-source breakdowns.
7. Do not count shared verification cost as a discovery win for either method.
8. Build/test pass/fail belongs in evidence unless timing/count/size is part of a measured metric.
9. Count only content that the agent actually sees or reads into context. Do not count Anvien internal processing, `.anvien/graph.json`, or redirected artifacts unless their content is returned to or read by the agent.

## Metric Definitions

| Metric | Unit | Definition |
|---|---:|---|
| `characters` | chars | Characters returned to or read by the agent for a bucket. |
| `estimated_tokens` | tokens | `ceil(characters / 4)`. |
| `unique_files_read` | files | Unique source/doc files opened into the agent context during a phase. |
| `file_bytes_read` | bytes | Sum of bytes for files read into the agent context. |
| `search_calls` | calls | Native list/search commands such as `rg`, `rg --files`, `Get-ChildItem`, or equivalent. |
| `anvien_calls` | calls | Anvien CLI/MCP/tool/resource calls used in Anvien-guided discovery. |
| `elapsed_seconds` | seconds | Wall-clock time for the phase. |
| `candidate_count` | candidates | Candidates reported by that discovery method before shared verification. |
| `confirmed_deadcode` | candidates | Candidates verified as confirmed deadcode. |
| `likely_deadcode` | candidates | Candidates verified as likely deadcode. |
| `uncertain` | candidates | Candidates with insufficient proof or material dynamic/public risk. |
| `false_positive` | candidates | Candidates disproven during verification. |

## Token Inclusion Rules

| Data source | Count tokens? | Rule |
|---|---|---|
| Anvien internal graph computation | no | Not visible to the agent. |
| `.anvien/graph.json` generated on disk | no | Count only if the agent reads file content. |
| Anvien command stdout/stderr shown to agent | yes | Count returned characters. |
| Anvien output redirected to `.tmp/*.json` | no | Count only process status unless the file is later read. |
| Later `Get-Content .tmp/anvien-output.json` | yes | Count the read content. |
| Native `rg` output shown to agent | yes | Count returned characters. |
| Source file opened/read by agent | yes | Count file content read into context. |

## B0 - Baseline Metrics

Status: pending

| Metric | Unit | Value |
|---|---:|---:|
| Baseline commit | SHA | pending |
| Dirty source files at baseline | files | pending |
| Dirty benchmark docs/reports at baseline | files | pending |
| CPU logical processors | count | pending |
| Physical memory | bytes | pending |
| Git version | version | pending |
| Go version | version | pending |
| Node version | version | pending |
| npm version | version | pending |

## B1 - Discovery Cost Summary

Status: pending

| Method | elapsed_seconds | total_characters | total_estimated_tokens | unique_files_read | file_bytes_read | search_calls | anvien_calls | candidate_count |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| Native-search | pending | pending | pending | pending | pending | pending | 0 | pending |
| Anvien-guided | pending | pending | pending | pending | pending | pending | pending | pending |

## B2 - Native-Search Token Breakdown

Status: pending

| Token source | characters | estimated_tokens | Notes |
|---|---:|---:|---|
| Task prompt | pending | pending | same task prompt as Anvien-guided where possible |
| Search/list command output | pending | pending | native search only |
| File reads | pending | pending | source/doc files read |
| Graph tool output | 0 | 0 | prohibited in native phase |
| Agent response | pending | pending | native conclusion/report text |
| Validation output | pending | pending | source-backed proof/disproof output |
| Retry/error output | pending | pending | failed commands or detours |
| Total | pending | pending | sum |

## B3 - Anvien-Guided Token Breakdown

Status: pending

| Token source | characters | estimated_tokens | Notes |
|---|---:|---:|---|
| Task prompt | pending | pending | same task prompt as native where possible |
| Search/list command output | pending | pending | follow-up native search after Anvien guidance |
| File reads | pending | pending | source/doc files read |
| Graph tool output | pending | pending | Anvien query/context/impact/etc. |
| Agent response | pending | pending | Anvien-guided conclusion/report text |
| Validation output | pending | pending | source-backed proof/disproof output |
| Retry/error output | pending | pending | failed commands or detours |
| Total | pending | pending | sum |

## B4 - File Read And Command Counts

Status: pending

| Method | unique_files_read | file_bytes_read | source_files_read | docs_read | search_calls | anvien_calls | failed_or_retry_calls |
|---|---:|---:|---:|---:|---:|---:|---:|
| Native-search | pending | pending | pending | pending | pending | 0 | pending |
| Anvien-guided | pending | pending | pending | pending | pending | pending | pending |
| Shared verification | pending | pending | pending | pending | pending | pending | pending |

## B5 - Candidate Discovery Counts

Status: pending

| Method | candidate_count | symbol_level | file_level | package_or_module_level | route_or_tool_surface | other |
|---|---:|---:|---:|---:|---:|---:|
| Native-search | pending | pending | pending | pending | pending | pending |
| Anvien-guided | pending | pending | pending | pending | pending | pending |
| Union | pending | pending | pending | pending | pending | pending |

## B6 - Candidate Verification Counts

Status: pending

| Method source | candidates | confirmed_deadcode | likely_deadcode | uncertain | false_positive |
|---|---:|---:|---:|---:|---:|
| Native-search | pending | pending | pending | pending | pending |
| Anvien-guided | pending | pending | pending | pending | pending |
| Found by both | pending | pending | pending | pending | pending |
| Native only | pending | pending | pending | pending | pending |
| Anvien only | pending | pending | pending | pending | pending |
| Union | pending | pending | pending | pending | pending |

## B7 - Efficiency Ratios

Status: pending

| Metric | Formula | Value |
|---|---|---:|
| Token reduction from Anvien | `(native_total_tokens - anvien_total_tokens) / native_total_tokens` | pending |
| File-read reduction from Anvien | `(native_unique_files - anvien_unique_files) / native_unique_files` | pending |
| Search-call reduction from Anvien | `(native_search_calls - anvien_search_calls) / native_search_calls` | pending |
| True-candidate delta | `(anvien_confirmed + anvien_likely) - (native_confirmed + native_likely)` | pending |
| False-positive delta | `anvien_false_positive - native_false_positive` | pending |
| Verification burden delta | `anvien_uncertain - native_uncertain` | pending |

## B8 - Final Outcome Matrix

Status: pending

| Question | Result | Numeric support |
|---|---|---|
| Did Anvien-guided discovery use fewer estimated tokens? | pending | pending |
| Did Anvien-guided discovery read fewer files? | pending | pending |
| Did Anvien-guided discovery use fewer native search calls? | pending | pending |
| Did Anvien-guided discovery find at least as many confirmed/likely candidates? | pending | pending |
| Did Anvien-guided discovery avoid increasing false positives? | pending | pending |
| Overall result | pending | pending |

## B9 - Required Final Token Shape

Status: pending

| Method | total_tokens | search_output_tokens | file_read_tokens | graph_tool_output_tokens | agent_response_tokens | validation_retry_tokens | files_read | correct |
|---|---:|---:|---:|---:|---:|---:|---:|---|
| Native-search | pending | pending | pending | 0 | pending | pending | pending | pending |
| Anvien-guided | pending | pending | pending | pending | pending | pending | pending | pending |
