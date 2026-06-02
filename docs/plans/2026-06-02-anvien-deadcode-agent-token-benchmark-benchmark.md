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
8. Count only context visible to the main agent.
9. Do not count Anvien internal graph computation, `.anvien/graph.json`, redirected files, or capture-only command bodies unless their content enters main-agent context.
10. Count visible output from every Anvien command called by the main agent.
11. If visible character counts cannot be measured for a phase, mark that phase's token comparison invalid.

## Metric Definitions

| Metric | Unit | Definition |
|---|---:|---|
| `visible_characters` | chars | Characters actually visible to the main agent. |
| `estimated_tokens` | tokens | `ceil(visible_characters / 4)` unless exact tokenizer data is available. |
| `tool_call_argument_tokens` | tokens | Estimated tokens for command/tool-call text emitted by the main agent. |
| `tool_result_tokens` | tokens | Estimated tokens for visible stdout/stderr/tool-result text received by the main agent. |
| `anvien_visible_tokens` | tokens | Visible output tokens from all Anvien commands in the Anvien-guided phase. |
| `search_output_tokens` | tokens | Visible output tokens from native list/search commands. |
| `source_read_tokens` | tokens | Visible source/file content tokens read by the main agent. |
| `agent_response_tokens` | tokens | Main-agent text output tokens for reports/conclusions. |
| `retry_error_tokens` | tokens | Visible failed command/error/retry output tokens. |
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

## Token Inclusion Rules

| Data source | Count tokens? | Rule |
|---|---|---|
| Task prompt shown to main agent | yes | Count fixed prompt text once per discovery method. |
| Tool-call command/argument text emitted by main agent | yes | Count visible command/tool arguments when measurable. |
| Native `rg`/list/search output shown to main agent | yes | Count visible stdout/stderr text. |
| Source file opened/read by main agent | yes | Count visible file content. |
| Anvien command stdout/stderr shown to main agent | yes | Count visible output from every Anvien command. |
| Anvien internal graph/index/cache computation | no | Not visible to main agent. |
| `.anvien/graph.json` generated on disk | no | Count only if main agent reads file content. |
| Output redirected to `.tmp/*.json` and not read | no | Count only visible process status/summary. |
| Later `Get-Content .tmp/anvien-output.json` | yes | Count the visible file content read by main agent. |
| Captured command body where only a summary/count is printed | no for captured body | Count only the visible summary/count text. |
| Tool output truncated by environment | partial | Count only the visible portion if visible character count is known; otherwise mark token comparison invalid. |
| Agent-written report/final answer | yes | Count main-agent response text for the phase. |

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
| Exact visible output counting available | yes/no | pending |
| Truncated-output handling valid | yes/no | pending |
| Token comparison valid | yes/no | pending |

## B2 - Discovery Cost Summary

Status: pending

| Method | elapsed_seconds | agent_visible_total_tokens | visible_characters | unique_files_read | file_bytes_read | search_calls | anvien_calls | candidate_count |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| Native-search | pending | pending | pending | pending | pending | pending | 0 | pending |
| Anvien-guided | pending | pending | pending | pending | pending | pending | pending | pending |

## B3 - Native-Search Token Breakdown

Status: pending

| Token source | visible_characters | estimated_tokens | Notes |
|---|---:|---:|---|
| Task prompt | pending | pending | fixed task prompt |
| Tool call arguments | pending | pending | native command/tool text |
| Search/list command output | pending | pending | native search only |
| File reads | pending | pending | source/doc files read |
| Anvien visible output | 0 | 0 | prohibited in native phase |
| Agent response | pending | pending | native conclusion/report text |
| Retry/error output | pending | pending | failed commands or detours |
| Total | pending | pending | sum |

## B4 - Anvien-Guided Token Breakdown

Status: pending

| Token source | visible_characters | estimated_tokens | Notes |
|---|---:|---:|---|
| Task prompt | pending | pending | fixed task prompt |
| Tool call arguments | pending | pending | Anvien and follow-up command/tool text |
| Anvien visible output | pending | pending | all visible output from all Anvien commands |
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
| Token reduction from Anvien | `(native_agent_visible_total_tokens - anvien_agent_visible_total_tokens) / native_agent_visible_total_tokens` | pending |
| File-read reduction from Anvien | `(native_unique_files - anvien_unique_files) / native_unique_files` | pending |
| Search-call reduction from Anvien | `(native_search_calls - anvien_search_calls) / native_search_calls` | pending |
| True-candidate delta | `(anvien_confirmed + anvien_likely) - (native_confirmed + native_likely)` | pending |
| False-positive delta | `anvien_false_positive - native_false_positive` | pending |
| Verification burden delta | `anvien_uncertain - native_uncertain` | pending |

## B9 - Final Outcome Matrix

Status: pending

| Question | Result | Numeric support |
|---|---|---|
| Did Anvien-guided discovery use fewer agent-visible tokens? | pending | pending |
| Did Anvien-guided discovery read fewer files? | pending | pending |
| Did Anvien-guided discovery use fewer native search calls? | pending | pending |
| Did Anvien-guided discovery find at least as many confirmed/likely candidates? | pending | pending |
| Did Anvien-guided discovery avoid increasing false positives? | pending | pending |
| Was token comparison valid? | pending | pending |
| Overall result | pending | pending |

## B10 - Shared Verification Token Shape

Status: pending

| Metric | tokens | visible_characters | Notes |
|---|---:|---:|---|
| Verification prompt/context | pending | pending | shared across union candidates |
| Verification tool call arguments | pending | pending | verification commands/tools |
| Verification search output | pending | pending | source-backed checks |
| Verification file reads | pending | pending | source/doc files read for proof |
| Verification agent response | pending | pending | verdict narrative |
| Verification retry/error output | pending | pending | failed commands or detours |
| Shared verification total | pending | pending | not counted as either discovery method cost |
