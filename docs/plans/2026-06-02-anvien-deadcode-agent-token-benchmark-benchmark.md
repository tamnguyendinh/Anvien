# Anvien Deadcode Agent Token Benchmark Ledger

Date: 2026-06-02

Status: complete

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

Status: complete

| Metric | Unit | Value |
|---|---:|---:|
| Baseline commit | SHA | `6516020c323b54c74583fbaa2caf81dad4475036` |
| Dirty source files at baseline | files | 0 |
| Dirty benchmark docs/reports at baseline | files | 0 |
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

Status: discovery complete; verification pending

| Method | elapsed_seconds | token_measurement_valid | agent_session_tokens | agent_session_token_proxy | unique_files_read | file_bytes_read | search_calls | anvien_calls | candidate_count |
|---|---:|---|---:|---:|---:|---:|---:|---:|---:|
| Native mode | 391 | yes | 856990 | n/a | 10 | 55619 | 39 | 0 | 12 |
| Anvien mode | 429 | yes | 1524993 | n/a | 11 | 69934 | 11 | 30 | 6 |

## B3 - Native Mode Token Breakdown

Status: complete

| Token source | tokens | Notes |
|---|---:|---|
| Task prompt | included in runtime total | fixed prompt plus native mode envelope; Codex CLI telemetry does not split this bucket |
| Tool-call arguments | included in runtime total | 39 completed command executions; bucket not split by runtime telemetry |
| Delivered tool/search results | included in runtime total | command outputs delivered inside measured Codex CLI session |
| Delivered file content | included in runtime total | 10 targeted file reads, 55,619 bytes |
| Agent response | included in runtime total | final native candidate report, 6,347 chars |
| Retry/error output | included in runtime total | 9 nonzero command exits/no-match searches |
| Total `agent_session_tokens` or proxy | 856990 | `input_tokens=850171` + `output_tokens=6819` |
| Token measurement valid | yes | Codex CLI `turn.completed.usage` telemetry |

## B4 - Anvien Mode Token Breakdown

Status: complete

| Token source | tokens | Notes |
|---|---:|---|
| Task prompt | included in runtime total | fixed prompt plus Anvien mode envelope; Codex CLI telemetry does not split this bucket |
| Tool-call arguments | included in runtime total | 41 completed total commands, including 30 Anvien commands |
| Delivered Anvien tool results | included in runtime total | Anvien command output delivered inside measured Codex CLI session |
| Delivered follow-up search results | included in runtime total | 11 non-Anvien completed commands after graph guidance |
| Delivered file content | included in runtime total | 11 targeted repo file reads, 69,934 bytes |
| Agent response | included in runtime total | final Anvien candidate report, 3,558 chars |
| Retry/error output | included in runtime total | 6 failed/retry commands plus one pre-discovery Codex CLI usage-limit attempt recorded separately in evidence |
| Total `agent_session_tokens` or proxy | 1524993 | `input_tokens=1519088` + `output_tokens=5905` |
| Token measurement valid | yes | Codex CLI `turn.completed.usage` telemetry |

## B5 - Local Tool Output Diagnostics

Status: complete

These metrics are diagnostic only. They are not token score unless the content was delivered into the agent session and counted in B3/B4.

| Method | diagnostic | unit | value | Counted as token score? |
|---|---|---:|---:|---|
| Native mode | delivered command output observed by Codex CLI session | chars | 1196125 | yes, through runtime telemetry rather than as separate proxy score |
| Anvien mode | delivered command output observed by Codex CLI session | chars | 2541841 | yes, through runtime telemetry rather than as separate proxy score |
| Anvien mode | clean-worktree `.anvien/graph.json` local artifact size | bytes | 328064552 | no |

## B6 - File Read And Command Counts

Status: discovery complete; verification pending

| Method | unique_files_read | file_bytes_read | source_files_read | docs_read | search_calls | anvien_calls | failed_or_retry_calls |
|---|---:|---:|---:|---:|---:|---:|---:|
| Native mode | 10 | 55619 | 10 | 0 | 39 | 0 | 9 |
| Anvien mode | 11 | 69934 | 11 | 0 | 11 | 30 | 6 |
| Shared verification | pending | pending | pending | pending | pending | pending | pending |

## B7 - Candidate Discovery Counts

Status: discovery complete; verification pending

| Method | candidate_count | symbol_level | file_level | package_or_module_level | route_or_tool_surface | other |
|---|---:|---:|---:|---:|---:|---:|
| Native mode | 12 | 8 | 1 | 2 | 0 | 1 |
| Anvien mode | 6 | 5 | 1 | 0 | 0 | 0 |
| Union | 18 | 13 | 2 | 2 | 0 | 1 |

## B8 - Candidate Verification Counts

Status: complete

| Method source | candidates | confirmed_deadcode | likely_deadcode | uncertain | false_positive |
|---|---:|---:|---:|---:|---:|
| Native mode | 12 | 5 | 7 | 0 | 0 |
| Anvien mode | 6 | 2 | 3 | 0 | 1 |
| Found by both | 0 | 0 | 0 | 0 | 0 |
| Native only | 12 | 5 | 7 | 0 | 0 |
| Anvien only | 6 | 2 | 3 | 0 | 1 |
| Union | 18 | 7 | 10 | 0 | 1 |

## B9 - Efficiency Ratios

Status: complete

| Metric | Formula | Value |
|---|---|---:|
| Token delta | `anvien_agent_session_tokens - native_agent_session_tokens` | 668003 |
| Token reduction from Anvien | `(native_agent_session_tokens - anvien_agent_session_tokens) / native_agent_session_tokens` | -77.95% |
| File-read reduction from Anvien | `(native_unique_files - anvien_unique_files) / native_unique_files` | -10.00% |
| Search-call reduction from Anvien | `(native_search_calls - anvien_search_calls) / native_search_calls` | 71.79% native follow-up search reduction; total Anvien-mode completed commands were 41 vs native 39 |
| True-candidate delta | `(anvien_confirmed + anvien_likely) - (native_confirmed + native_likely)` | -7 |
| False-positive delta | `anvien_false_positive - native_false_positive` | 1 |
| Verification burden delta | `anvien_uncertain - native_uncertain` | 0 |

## B10 - Final Outcome Matrix

Status: complete

| Question | Result | Numeric support |
|---|---|---|
| Did Anvien mode spend fewer AI-agent tokens? | no | native 856,990; Anvien 1,524,993; Anvien spent 668,003 more |
| Did Anvien mode read fewer files? | no | native 10; Anvien 11 |
| Did Anvien mode use fewer native search calls? | yes for native follow-up searches only; no for total commands | native search/commands 39; Anvien follow-up native commands 11 plus 30 Anvien commands |
| Did Anvien mode find at least as many confirmed/likely candidates? | no | native 12 confirmed/likely; Anvien 5 confirmed/likely |
| Did Anvien mode avoid increasing false positives? | no | native 0; Anvien 1 |
| Was token comparison valid? | yes | both modes used Codex CLI `turn.completed.usage` telemetry |
| Overall result | native wins this run on token spend, file reads, confirmed/likely candidates, and false positives | Anvien spent 77.95% more tokens and found 7 fewer confirmed/likely candidates |

## B11 - Shared Verification Token Shape

Status: not token-scored; excluded from discovery comparison

| Metric | tokens | Notes |
|---|---:|---|
| Verification prompt/context | n/a | shared verification was performed after both discovery reports closed and was not assigned to either mode |
| Verification tool-call arguments | n/a | verification commands are recorded in evidence, not as discovery cost |
| Verification delivered tool/search results | n/a | source-backed checks are recorded in E6 |
| Verification delivered file content | n/a | source/doc proof is recorded in E6 |
| Verification agent response | n/a | verdict narrative is recorded in E6/E7 |
| Verification retry/error output | n/a | no verification blocker; one malformed PowerShell search was rerun with corrected quoting |
| Shared verification total | n/a | not counted as either discovery mode cost |
