# Anvien Deadcode Agent Token Benchmark Ledger

Date: 2026-06-02

Status: reset; not started

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

## B1 - Token Accountant Setup Metrics

Status: pending

| Metric | Unit | Value |
|---|---:|---:|
| Token accountant active before native phase | yes/no | pending |
| Token accountant active before Anvien phase | yes/no | pending |
| Exact observed output counting available | yes/no | pending |
| Truncated-output handling valid | yes/no | pending |
| Token comparison valid | yes/no | pending |

## B2 - Discovery Cost Summary

Status: pending

| Method | elapsed_seconds | observed_total_tokens | observed_characters | unique_files_read | file_bytes_read | search_calls | anvien_calls | candidate_count |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| Native-search | pending | pending | pending | pending | pending | pending | 0 | pending |
| Anvien-guided | pending | pending | pending | pending | pending | pending | pending | pending |

## B3 - Native-Search Token Breakdown

Status: pending

| Token source | observed_characters | estimated_tokens | Notes |
|---|---:|---:|---|
| Task prompt | pending | pending | fixed task prompt |
| Tool call arguments | pending | pending | native command/tool text |
| Search/list command output | pending | pending | native search only |
| File reads | pending | pending | source/doc files read |
| Anvien observed output | 0 | 0 | prohibited in native phase |
| Agent response | pending | pending | native conclusion/report text |
| Retry/error output | pending | pending | failed commands or detours |
| Total | pending | pending | sum |

## B4 - Anvien-Guided Token Breakdown

Status: pending

| Token source | observed_characters | estimated_tokens | Notes |
|---|---:|---:|---|
| Task prompt | pending | pending | fixed task prompt |
| Tool call arguments | pending | pending | Anvien and follow-up command/tool text |
| Anvien observed context | pending | pending | all observed output/artifact content from Anvien work |
| Follow-up search/list output | pending | pending | native search after Anvien guidance |
| File reads | pending | pending | source/doc files read |
| Agent response | pending | pending | Anvien-guided conclusion/report text |
| Retry/error output | pending | pending | failed commands or detours |
| Total | pending | pending | sum |

## B5 - File Read And Command Counts

Status: pending

| Method | unique_files_read | file_bytes_read | source_files_read | docs_read | search_calls | anvien_calls | failed_or_retry_calls |
|---|---:|---:|---:|---:|---:|---:|---:|
| Native-search | pending | pending | pending | pending | pending | 0 | pending |
| Anvien-guided | pending | pending | pending | pending | pending | pending | pending |
| Shared verification | pending | pending | pending | pending | pending | pending | pending |

## B6 - Candidate Discovery Counts

Status: pending

| Method | candidate_count | symbol_level | file_level | package_or_module_level | route_or_tool_surface | other |
|---|---:|---:|---:|---:|---:|---:|
| Native-search | pending | pending | pending | pending | pending | pending |
| Anvien-guided | pending | pending | pending | pending | pending | pending |
| Union | pending | pending | pending | pending | pending | pending |

## B7 - Candidate Verification Counts

Status: pending

| Method source | candidates | confirmed_deadcode | likely_deadcode | uncertain | false_positive |
|---|---:|---:|---:|---:|---:|
| Native-search | pending | pending | pending | pending | pending |
| Anvien-guided | pending | pending | pending | pending | pending |
| Found by both | pending | pending | pending | pending | pending |
| Native only | pending | pending | pending | pending | pending |
| Anvien only | pending | pending | pending | pending | pending |
| Union | pending | pending | pending | pending | pending |

## B8 - Efficiency Ratios

Status: pending

| Metric | Formula | Value |
|---|---|---:|
| Token reduction from Anvien | `(native_observed_total_tokens - anvien_observed_total_tokens) / native_observed_total_tokens` | pending |
| File-read reduction from Anvien | `(native_unique_files - anvien_unique_files) / native_unique_files` | pending |
| Search-call reduction from Anvien | `(native_search_calls - anvien_search_calls) / native_search_calls` | pending |
| True-candidate delta | `(anvien_confirmed + anvien_likely) - (native_confirmed + native_likely)` | pending |
| False-positive delta | `anvien_false_positive - native_false_positive` | pending |
| Verification burden delta | `anvien_uncertain - native_uncertain` | pending |

## B9 - Final Outcome Matrix

Status: pending

| Question | Result | Numeric support |
|---|---|---|
| Did Anvien-guided discovery use fewer observed-context tokens? | pending | pending |
| Did Anvien-guided discovery read fewer files? | pending | pending |
| Did Anvien-guided discovery use fewer native search calls? | pending | pending |
| Did Anvien-guided discovery find at least as many confirmed/likely candidates? | pending | pending |
| Did Anvien-guided discovery avoid increasing false positives? | pending | pending |
| Was token comparison valid? | pending | pending |
| Overall result | pending | pending |

## B10 - Shared Verification Token Shape

Status: pending

| Metric | tokens | observed_characters | Notes |
|---|---:|---:|---|
| Verification prompt/context | pending | pending | shared across union candidates |
| Verification tool call arguments | pending | pending | verification commands/tools |
| Verification search output | pending | pending | source-backed checks |
| Verification file reads | pending | pending | source/doc files read for proof |
| Verification agent response | pending | pending | verdict narrative |
| Verification retry/error output | pending | pending | failed commands or detours |
| Shared verification total | pending | pending | not counted as either discovery method cost |
