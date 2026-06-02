# Anvien Deadcode Agent Token Benchmark Ledger

Date: 2026-06-02

Status: complete

Companion files:

- Plan: [2026-06-02-anvien-deadcode-agent-token-benchmark-plan.md](2026-06-02-anvien-deadcode-agent-token-benchmark-plan.md)
- Evidence ledger: [2026-06-02-anvien-deadcode-agent-token-benchmark-evidence.md](2026-06-02-anvien-deadcode-agent-token-benchmark-evidence.md)

## Reset Notice

Old benchmark numbers from the invalidated run have been removed. This ledger starts clean for the rerun.

## Benchmark Rules

1. Record quantitative data only.
2. Put command details, reasoning, token-accountant notes, and candidate proof in the evidence ledger.
3. Use `estimated_tokens = ceil(characters / 4)` when exact model tokenizer accounting is unavailable.
4. Separate discovery cost from shared verification cost.
5. Separate native-search and Anvien-guided metrics.
6. Record both totals and token-source breakdowns.
7. Do not count shared verification cost as a discovery win for either method.
8. Token accounting is observation-based: count only context the token accountant observes the main agent receive, read, or emit.
9. Do not include or exclude content by source category alone. Anvien command output, graph files, redirected files, cache/index artifacts, summaries, and source files are counted only when the accountant observes them entering main-agent context.
10. Count observed output from every Anvien command called by the main agent.
11. If observed character counts cannot be measured for a phase, mark that phase's token comparison invalid.

## Metric Definitions

| Metric | Unit | Definition |
|---|---:|---|
| `observed_characters` | chars | Characters the token accountant observes entering or leaving main-agent context. |
| `estimated_tokens` | tokens | `ceil(observed_characters / 4)` unless exact tokenizer data is available. |
| `tool_call_argument_tokens` | tokens | Estimated tokens for command/tool-call text emitted by the main agent. |
| `tool_result_tokens` | tokens | Estimated tokens for observed stdout/stderr/tool-result text received by the main agent. |
| `anvien_observed_tokens` | tokens | Observed Anvien-related context in the Anvien-guided phase. |
| `search_output_tokens` | tokens | Observed output tokens from native list/search commands. |
| `source_read_tokens` | tokens | Observed source/file content tokens read by the main agent. |
| `agent_response_tokens` | tokens | Main-agent text output tokens for reports/conclusions. |
| `retry_error_tokens` | tokens | Observed failed command/error/retry output tokens. |
| `unique_files_read` | files | Unique source/doc files opened into main-agent context during a phase. |
| `file_bytes_read` | bytes | Bytes for file content read into main-agent context. |
| `search_calls` | calls | Native list/search commands such as `rg`, `rg --files`, `Get-ChildItem`, or equivalent. |
| `anvien_calls` | calls | Anvien CLI/MCP/tool/resource calls used in Anvien-guided discovery. |
| `elapsed_seconds` | seconds | Wall-clock time for the phase. |
| `candidate_count` | candidates | Candidates reported by that discovery method before shared verification. |
| `confirmed_deadcode` | candidates | Candidates verified as confirmed deadcode. |
| `likely_deadcode` | candidates | Candidates verified as likely deadcode. |
| `uncertain` | candidates | Candidates with insufficient proof or material dynamic/public risk. |
| `false_positive` | candidates | Candidates disproven during verification. |

## Token Observation Rules

| Data source | Count tokens? | Rule |
|---|---|---|
| Task prompt shown to main agent | observed => yes | Count fixed prompt text once per discovery method when the accountant observes it. |
| Tool-call command/argument text emitted by main agent | observed => yes | Count command/tool arguments when the accountant observes them. |
| Native `rg`/list/search output | observed => yes | Count stdout/stderr text observed by the accountant. |
| Source file content | observed => yes | Count file content the accountant observes the main agent read. |
| Any Anvien command stdout/stderr/tool result | observed => yes | Count observed output from every Anvien command. |
| Anvien graph file, output file, cache/index artifact, or generated artifact | observed => yes | Count it if the accountant observes the main agent read/receive it; otherwise record it as unobserved and exclude it from token totals. |
| Output redirected to `.tmp/*.json` | observed => yes | The accountant records observed redirected content when the main agent receives/reads it; unobserved redirected bodies are recorded as unobserved and excluded from token totals. |
| Captured command body where only a summary/count is printed | summary observed => summary only | Count the observed summary/count text; count the captured body only if the accountant observes it entering main-agent context. |
| Tool output truncated by environment | observed portion only | Count the observed portion if measurable; otherwise mark token comparison invalid. |
| Agent-written report/final answer | observed => yes | Count main-agent response text for the phase. |

## B0 - Baseline Metrics

Status: complete

| Metric | Unit | Value |
|---|---:|---:|
| Baseline commit | SHA | `6564d7d5f5f7d53767a4afbc1028cda26535b977` |
| Discovery start commit | SHA | `e0b2f336b37a7d65eec6ddf5d6f20ac7dfd40900` |
| Dirty source files at baseline | files | 0 |
| Dirty benchmark docs/reports at baseline | files | 0 |
| CPU logical processors | count | 8 |
| Physical memory | bytes | 33238466560 |
| Git version | version | 2.54.0.windows.1 |
| Go version | version | 1.26.3 windows/amd64 |
| Node version | version | v24.15.0 |
| npm version | version | 11.12.1 |

## B1 - Token Accountant Setup Metrics

Status: complete

| Metric | Unit | Value |
|---|---:|---:|
| Token accountant active before native phase | yes/no | yes |
| Token accountant active before Anvien phase | yes/no | yes |
| Exact observed output counting available | yes/no | yes |
| Truncated-output handling valid | yes/no | yes |
| Token comparison valid | yes/no | no |

## B2 - Discovery Cost Summary

Status: complete

| Method | elapsed_seconds | observed_total_tokens | observed_characters | unique_files_read | file_bytes_read | search_calls | anvien_calls | candidate_count |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| Native-search | 174 | 27088 | 108350 | 10 | 10300 | 29 | 0 | 8 |
| Anvien-guided | 180 approx | invalid; reconstructed >=215431 | invalid; reconstructed >=861723 | 11 | 6160 | 17 | 24 | 14 |

## B3 - Native-Search Token Breakdown

Status: complete

| Token source | observed_characters | estimated_tokens | Notes |
|---|---:|---:|---|
| Task prompt | included in total | included in total | fixed task prompt |
| Tool call arguments | 15900 | 3975 | native command/tool text |
| Search/list command output | 23800 | 5950 | native search only |
| File reads | 10300 | 2575 | source/config files visibly read |
| Anvien observed output | 0 | 0 | prohibited in native phase |
| Agent response | 6200 | 1550 | native conclusion/report text |
| Retry/error output | 2650 | 663 | failed commands or detours |
| Other command stdout/stderr | 49500 | 12375 | command output not separately bucketed as source/search |
| Total | 108350 | 27088 | native self-accountant ledger |

## B4 - Anvien-Guided Token Breakdown

Status: complete; invalid token axis

| Token source | observed_characters | estimated_tokens | Notes |
|---|---:|---:|---|
| Task prompt | invalid | invalid | included in phase but exact total invalid |
| Tool call arguments | invalid | invalid | not exact after transcript truncation |
| Anvien observed context | 777694 reconstructed | 194424 reconstructed | large `context` outputs were transcript-truncated |
| Follow-up search/list output | 77869 reconstructed | 19468 reconstructed | one broad native search output was transcript-truncated |
| File reads | 6160 | 1540 | source snippets visibly read |
| Agent response | invalid | invalid | not exact after transcript truncation |
| Retry/error output | included in reconstructed categories | included in reconstructed categories | 6 failed/retry calls |
| Total | invalid; >=861723 reconstructed | invalid; >=215431 reconstructed | no token winner may be claimed |

## B5 - File Read And Command Counts

Status: complete

| Method | unique_files_read | file_bytes_read | source_files_read | docs_read | search_calls | anvien_calls | failed_or_retry_calls |
|---|---:|---:|---:|---:|---:|---:|---:|
| Native-search | 10 | 10300 | 10 | 0 | 29 | 0 | 4 |
| Anvien-guided | 11 | 6160 | 11 | 0 | 17 | 24 | 6 |
| Shared verification | 19 approx | 27500 approx | 17 approx | 0 | 7 approx | 0 | 2 |

## B6 - Candidate Discovery Counts

Status: complete

| Method | candidate_count | symbol_level | file_level | package_or_module_level | route_or_tool_surface | other |
|---|---:|---:|---:|---:|---:|---:|
| Native-search | 8 | 6 | 0 | 0 | 0 | 2 |
| Anvien-guided | 14 | 14 | 0 | 0 | 0 | 0 |
| Union | 21 | 19 | 0 | 0 | 0 | 2 |

## B7 - Candidate Verification Counts

Status: complete

| Method source | candidates | confirmed_deadcode | likely_deadcode | uncertain | false_positive |
|---|---:|---:|---:|---:|---:|
| Native-search | 8 | 4 | 4 | 0 | 0 |
| Anvien-guided | 14 | 13 | 1 | 0 | 0 |
| Found by both | 1 | 0 | 1 | 0 | 0 |
| Native only | 7 | 4 | 3 | 0 | 0 |
| Anvien only | 13 | 13 | 0 | 0 | 0 |
| Union | 21 | 17 | 4 | 0 | 0 |

## B8 - Efficiency Ratios

Status: complete

| Metric | Formula | Value |
|---|---|---:|
| Token reduction from Anvien | `(native_observed_total_tokens - anvien_observed_total_tokens) / native_observed_total_tokens` | invalid |
| File-read reduction from Anvien | `(native_unique_files - anvien_unique_files) / native_unique_files` | -10.0% |
| Search-call reduction from Anvien | `(native_search_calls - anvien_search_calls) / native_search_calls` | 41.4% native follow-up search reduction; total tool calls increased |
| True-candidate delta | `(anvien_confirmed + anvien_likely) - (native_confirmed + native_likely)` | +6 |
| False-positive delta | `anvien_false_positive - native_false_positive` | 0 |
| Verification burden delta | `anvien_uncertain - native_uncertain` | 0 |

## B9 - Final Outcome Matrix

Status: complete

| Question | Result | Numeric support |
|---|---|---|
| Did Anvien-guided discovery use fewer observed-context tokens? | invalid | native 27,088 valid; Anvien reconstructed >=215,431 but exact invalid due truncation |
| Did Anvien-guided discovery read fewer files? | no | native 10; Anvien 11 |
| Did Anvien-guided discovery use fewer native search calls? | yes for native follow-up only; no for total calls | native 29; Anvien follow-up 17 plus 24 Anvien calls |
| Did Anvien-guided discovery find at least as many confirmed/likely candidates? | yes | native 8; Anvien 14 |
| Did Anvien-guided discovery avoid increasing false positives? | yes | native 0; Anvien 0 |
| Was token comparison valid? | no | Anvien phase had transcript-truncated outputs |
| Overall result | correctness favors Anvien; token winner invalid; file-read count favors native | union 21 candidates; Anvien-only 13 confirmed |

## B10 - Shared Verification Token Shape

Status: complete

| Metric | tokens | observed_characters | Notes |
|---|---:|---:|---|
| Verification prompt/context | approx 250 | approx 1000 | shared across union candidates |
| Verification tool call arguments | approx 900 | approx 3600 | verification commands/tools |
| Verification search output | approx 2500 | approx 10000 | source-backed checks |
| Verification file reads | approx 3400 | approx 13600 | source/config files read for proof |
| Verification agent response | not finalized here | not finalized here | verdict narrative lives in evidence ledger |
| Verification retry/error output | approx 250 | approx 1000 | failed commands or detours |
| Shared verification total | approx 7300 plus evidence text | approx 29200 plus evidence text | not counted as either discovery method cost |
