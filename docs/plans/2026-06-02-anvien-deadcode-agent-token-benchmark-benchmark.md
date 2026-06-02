# Anvien Deadcode Agent Token Benchmark Ledger

Date: 2026-06-02

Status: measurement harness ready; benchmark baseline pending

Companion files:

- Plan: [2026-06-02-anvien-deadcode-agent-token-benchmark-plan.md](2026-06-02-anvien-deadcode-agent-token-benchmark-plan.md)
- Evidence ledger: [2026-06-02-anvien-deadcode-agent-token-benchmark-evidence.md](2026-06-02-anvien-deadcode-agent-token-benchmark-evidence.md)

## Reset Notice

All prior benchmark numbers from the invalid run are discarded.

Reason: the old run did not produce a valid `agent_session_tokens` value for Anvien mode and incorrectly treated output-volume diagnostics as if they could answer the token question.

Do not reuse old token totals, reconstructed output volumes, file-read counts, candidate counts, or ratios.

## Benchmark Rules

1. Record quantitative data only.
2. Put command details, reasoning, token-measurement notes, and candidate proof in the evidence ledger.
3. The primary score is `agent_session_tokens`.
4. If runtime/model telemetry is unavailable, use `agent_session_token_proxy` only when the delivered transcript is exact and measurable.
5. Do not use Anvien output volume, graph size, cache size, index size, or redirected command body size as token score.
6. Separate discovery cost from shared verification cost.
7. Separate native and Anvien mode metrics.
8. Do not count shared verification cost as a discovery win for either mode.
9. If token usage cannot be measured for either mode, mark token comparison invalid before writing ratios.

## Metric Definitions

| Metric | Unit | Definition |
|---|---:|---|
| `agent_session_tokens` | tokens | Actual model input + model output tokens spent by the AI agent during the phase. |
| `agent_session_token_proxy` | tokens | Exact delivered-transcript proxy when telemetry is unavailable. |
| `token_measurement_valid` | yes/no | Whether token usage for the phase is valid for comparison. |
| `task_prompt_tokens` | tokens | Fixed task prompt delivered to the agent. |
| `tool_call_argument_tokens` | tokens | Tool-call command/argument text emitted by the agent. |
| `delivered_tool_result_tokens` | tokens | Tool result text delivered back into the agent session. |
| `delivered_file_content_tokens` | tokens | File content read into the agent session. |
| `agent_response_tokens` | tokens | Agent-visible response/report text for the phase. |
| `retry_error_tokens` | tokens | Error/retry text delivered into the agent session. |
| `local_tool_output_volume` | chars/bytes | Diagnostic only; local command data not necessarily delivered to the agent and not token score. |
| `unique_files_read` | files | Unique source/doc files opened into agent context during a phase. |
| `file_bytes_read` | bytes | Bytes for file content read into agent context. |
| `search_calls` | calls | Native list/search commands such as `rg`, `rg --files`, `Get-ChildItem`, or equivalent. |
| `anvien_calls` | calls | Anvien CLI/MCP/tool/resource calls used in Anvien mode. |
| `elapsed_seconds` | seconds | Wall-clock time for the phase. |
| `candidate_count` | candidates | Candidates reported by that discovery mode before shared verification. |
| `confirmed_deadcode` | candidates | Candidates verified as confirmed deadcode. |
| `likely_deadcode` | candidates | Candidates verified as likely deadcode. |
| `uncertain` | candidates | Candidates with insufficient proof or material dynamic/public risk. |
| `false_positive` | candidates | Candidates disproven during verification. |

## B0 - Baseline Metrics

Status: superseded by measurement-harness implementation; final benchmark baseline pending

| Metric | Unit | Value |
|---|---:|---:|
| Baseline commit | SHA | `b5568094184b92618e2627460c5a2f22b120a497` pre-harness; final pending |
| Dirty source files at baseline | files | 0 |
| Dirty benchmark docs/reports at baseline | files | 3 |
| CPU logical processors | count | 8 |
| Physical memory | bytes | 33238466560 |
| Git version | version | `2.54.0.windows.1` |
| Go version | version | `go1.26.3 windows/amd64` |
| Node version | version | `v24.15.0` |
| npm version | version | `11.12.1` |

## B1 - Token Measurement Setup Metrics

Status: complete

| Metric | Unit | Value |
|---|---:|---:|
| Measurement mechanism | name | `codex_exec_json_usage` primary; `tiktoken_transcript_proxy` fallback |
| Runtime token telemetry available | yes/no | yes |
| Exact transcript proxy available | yes/no | yes |
| Native token measurement eligible | yes/no | yes |
| Anvien token measurement eligible | yes/no | yes |
| Token comparison eligible before discovery | yes/no | yes |

## B2 - Discovery Cost Summary

Status: pending

| Method | elapsed_seconds | token_measurement_valid | agent_session_tokens | agent_session_token_proxy | unique_files_read | file_bytes_read | search_calls | anvien_calls | candidate_count |
|---|---:|---|---:|---:|---:|---:|---:|---:|---:|
| Native mode | pending | pending | pending | pending | pending | pending | pending | 0 | pending |
| Anvien mode | pending | pending | pending | pending | pending | pending | pending | pending | pending |

## B3 - Native Mode Token Breakdown

Status: pending

| Token source | tokens | Notes |
|---|---:|---|
| Task prompt | pending | fixed task prompt |
| Tool-call arguments | pending | native command/tool text |
| Delivered tool/search results | pending | native command output delivered to agent |
| Delivered file content | pending | source/doc files read into agent context |
| Agent response | pending | native conclusion/report text |
| Retry/error output | pending | failed commands or detours delivered to agent |
| Total `agent_session_tokens` or proxy | pending | primary native score |
| Token measurement valid | pending | yes/no |

## B4 - Anvien Mode Token Breakdown

Status: pending

| Token source | tokens | Notes |
|---|---:|---|
| Task prompt | pending | fixed task prompt |
| Tool-call arguments | pending | Anvien and follow-up command/tool text |
| Delivered Anvien tool results | pending | only Anvien output delivered into agent session |
| Delivered follow-up search results | pending | native search after Anvien guidance |
| Delivered file content | pending | source/doc files read into agent context |
| Agent response | pending | Anvien-mode conclusion/report text |
| Retry/error output | pending | failed commands or detours delivered to agent |
| Total `agent_session_tokens` or proxy | pending | primary Anvien-mode score |
| Token measurement valid | pending | yes/no |

## B5 - Local Tool Output Diagnostics

Status: pending

These metrics are diagnostic only. They are not token score unless the content was delivered into the agent session and counted in B3/B4.

| Method | diagnostic | unit | value | Counted as token score? |
|---|---|---:|---:|---|
| Native mode | local command output volume not delivered to agent | chars/bytes | pending | no |
| Anvien mode | Anvien local output/graph/cache/index volume not delivered to agent | chars/bytes | pending | no |

## B6 - File Read And Command Counts

Status: pending

| Method | unique_files_read | file_bytes_read | source_files_read | docs_read | search_calls | anvien_calls | failed_or_retry_calls |
|---|---:|---:|---:|---:|---:|---:|---:|
| Native mode | pending | pending | pending | pending | pending | 0 | pending |
| Anvien mode | pending | pending | pending | pending | pending | pending | pending |
| Shared verification | pending | pending | pending | pending | pending | pending | pending |

## B7 - Candidate Discovery Counts

Status: pending

| Method | candidate_count | symbol_level | file_level | package_or_module_level | route_or_tool_surface | other |
|---|---:|---:|---:|---:|---:|---:|
| Native mode | pending | pending | pending | pending | pending | pending |
| Anvien mode | pending | pending | pending | pending | pending | pending |
| Union | pending | pending | pending | pending | pending | pending |

## B8 - Candidate Verification Counts

Status: pending

| Method source | candidates | confirmed_deadcode | likely_deadcode | uncertain | false_positive |
|---|---:|---:|---:|---:|---:|
| Native mode | pending | pending | pending | pending | pending |
| Anvien mode | pending | pending | pending | pending | pending |
| Found by both | pending | pending | pending | pending | pending |
| Native only | pending | pending | pending | pending | pending |
| Anvien only | pending | pending | pending | pending | pending |
| Union | pending | pending | pending | pending | pending |

## B9 - Efficiency Ratios

Status: pending

| Metric | Formula | Value |
|---|---|---:|
| Token delta | `anvien_agent_session_tokens - native_agent_session_tokens` | pending |
| Token reduction from Anvien | `(native_agent_session_tokens - anvien_agent_session_tokens) / native_agent_session_tokens` | pending |
| File-read reduction from Anvien | `(native_unique_files - anvien_unique_files) / native_unique_files` | pending |
| Search-call reduction from Anvien | `(native_search_calls - anvien_search_calls) / native_search_calls` | pending |
| True-candidate delta | `(anvien_confirmed + anvien_likely) - (native_confirmed + native_likely)` | pending |
| False-positive delta | `anvien_false_positive - native_false_positive` | pending |
| Verification burden delta | `anvien_uncertain - native_uncertain` | pending |

## B10 - Final Outcome Matrix

Status: pending

| Question | Result | Numeric support |
|---|---|---|
| Did Anvien mode spend fewer AI-agent tokens? | pending | pending |
| Did Anvien mode read fewer files? | pending | pending |
| Did Anvien mode use fewer native search calls? | pending | pending |
| Did Anvien mode find at least as many confirmed/likely candidates? | pending | pending |
| Did Anvien mode avoid increasing false positives? | pending | pending |
| Was token comparison valid? | pending | pending |
| Overall result | pending | pending |

## B11 - Shared Verification Token Shape

Status: pending

| Metric | tokens | Notes |
|---|---:|---|
| Verification prompt/context | pending | shared across union candidates |
| Verification tool-call arguments | pending | verification commands/tools |
| Verification delivered tool/search results | pending | source-backed checks |
| Verification delivered file content | pending | source/doc files read for proof |
| Verification agent response | pending | verdict narrative |
| Verification retry/error output | pending | failed commands or detours |
| Shared verification total | pending | not counted as either discovery mode cost |
