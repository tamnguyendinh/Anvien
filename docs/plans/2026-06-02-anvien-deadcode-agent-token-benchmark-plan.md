# Anvien Deadcode Agent Token Benchmark Plan

Date: 2026-06-02

Status: complete

Companion files:

- Evidence ledger: [2026-06-02-anvien-deadcode-agent-token-benchmark-evidence.md](2026-06-02-anvien-deadcode-agent-token-benchmark-evidence.md)
- Benchmark ledger: [2026-06-02-anvien-deadcode-agent-token-benchmark-benchmark.md](2026-06-02-anvien-deadcode-agent-token-benchmark-benchmark.md)

## Reopen Notice

The previous benchmark run is void.

Reason: the plan/report drifted from the real question. The benchmark must not measure how much data Anvien creates. Anvien is a local tool and does not spend AI tokens. The benchmark must measure how many tokens the AI agent spends when solving the same deadcode task with and without Anvien.

All previous evidence and benchmark values are discarded and must not be reused.

## Master Rules

1. This is an AI-agent token benchmark, not an Anvien output-size benchmark.
2. Anvien does not spend tokens. The agent spends tokens.
3. The primary score is `agent_session_tokens` for each mode.
4. The two modes are:
   - Native mode: the agent finds deadcode without Anvien.
   - Anvien mode: the agent finds deadcode while using Anvien as a local tool.
5. The benchmark target is one task: find deadcode currently present in this repo.
6. Do not delete, rename, refactor, or edit candidate deadcode during this benchmark.
7. Use one git commit as the benchmark baseline and record the exact hash before discovery starts.
8. If source code is dirty at baseline time, stop and record the blocker.
9. Run native mode first. Native mode must not use Anvien commands, Anvien MCP tools, Anvien resources, Anvien generated context, `.anvien` graph/cache/index data, or prior Anvien reports.
10. Run Anvien mode second. Anvien mode must use Anvien as the discovery aid before broad source reads.
11. Keep discovery modes separate. Do not leak candidate lists, files-read logs, commands, reports, or conclusions from one mode into the other until both discovery reports are closed.
12. Use the fixed task prompt in this plan unchanged as the task text for both modes. A mode-control envelope may be added to enforce Native-vs-Anvien tool policy in isolated Codex CLI sessions; the envelope must be recorded and its runtime tokens are counted.
13. Every candidate must be verified after both discovery modes using source-backed evidence.
14. Candidate verification must classify each candidate as `confirmed_deadcode`, `likely_deadcode`, `uncertain`, or `false_positive`.
15. Verification cost is shared and is not counted as either mode's discovery cost.
16. Do not impose artificial elapsed-time, command-call, file-read, source-byte, or token limits on either method. These are measured outcomes, not stop gates.
17. Do not accept a token result unless both modes have valid `agent_session_tokens` or a valid exact transcript proxy.
18. Record evidence immediately after each phase and quantitative measurements immediately in the benchmark ledger.

## Goal

Answer this exact question:

```text
How many AI-agent tokens are spent to find deadcode without Anvien?
How many AI-agent tokens are spent to find deadcode with Anvien?
Which mode spends fewer AI-agent tokens, and by how much?
```

The final comparison must answer:

- Native mode token spend.
- Anvien mode token spend.
- Token delta and reduction ratio.
- Which deadcode candidates each mode found.
- Which candidates are confirmed, likely, uncertain, or false positives.
- Whether Anvien changed file reads, search calls, elapsed time, and correctness.

## Fixed Task Prompt

Use this exact prompt for both modes:

```text
Find deadcode currently present in this repository. Do not edit, delete, rename, or refactor code. Return candidates only. For each candidate, provide path, symbol or file name, kind, why it appears unused, commands or source facts used, and uncertainty or dynamic-use risk.
```

Codex CLI execution note:

- The fixed task prompt above must appear unchanged inside both mode envelopes.
- Native envelope must prohibit Anvien commands, MCP tools, resources, generated context, `.anvien` graph/cache/index data, and prior Anvien reports.
- Anvien envelope must require Anvien graph use before broad source reads.
- Envelope text is part of the measured agent session and is counted by `codex exec --json` runtime usage.

## Token Model

Primary metric:

```text
agent_session_tokens
= model_input_tokens
+ model_output_tokens
```

If exact model token telemetry is available, use it as source of truth.

If exact telemetry is not available, an exact transcript proxy may be used:

```text
agent_session_token_proxy
= task_prompt_tokens
+ agent_tool_call_argument_tokens
+ tool_result_tokens_delivered_to_agent
+ file_content_tokens_delivered_to_agent
+ agent_response_tokens
+ retry_error_tokens_delivered_to_agent
```

## What Counts

Count only tokens that are part of the AI agent's session.

| Item | Count? | Rule |
|---|---|---|
| User/task prompt delivered to agent | yes | Part of agent session. |
| Agent reasoning/output text if available in telemetry | yes | Part of agent session. |
| Agent visible response/report text | yes | Part of agent session. |
| Tool-call arguments emitted by agent | yes | Part of agent session/proxy. |
| Tool result text delivered back to agent | yes | Part of agent session/proxy. |
| File content read into agent context | yes | Part of agent session/proxy. |
| Error/retry output delivered to agent | yes | Part of agent session/proxy. |
| Anvien analyze internals | no | Local tool work; not agent tokens. |
| Anvien graph/cache/index files not read by agent | no | Local artifacts; not agent tokens. |
| Anvien generated output not delivered to agent | no | Not in agent session. |
| Command stdout redirected to file and not read by agent | no | Not in agent session. |
| A printed summary/count of hidden output | summary only | Count only the delivered summary/count text. |
| Output volume estimate such as `full_stdout_tokens: 228504` | text only | Count the printed text, not the hidden body. |

Important:

```text
Anvien output volume != AI agent token spend.
Graph size != AI agent token spend.
Cache/index size != AI agent token spend.
Only delivered agent-session content can be counted.
```

## Measurement Gate

Before discovery starts, choose one valid measurement mechanism:

1. Runtime/model token telemetry for the agent session.
2. Exact transcript proxy where every delivered prompt, tool call, tool result, file read, retry/error, and response can be measured.

If neither is available for both native and Anvien modes, stop before discovery and record the blocker. Do not run another invalid token benchmark.

## Deadcode Candidate Definition

A candidate is any source declaration, file, package-level object, helper, route/tool handler, exported surface, test helper, generated directive, or frontend symbol that appears unused or unreachable under normal repository behavior.

Minimum candidate fields:

| Field | Required |
|---|---|
| Candidate id | yes |
| Path | yes |
| Symbol/name | yes when symbol-level |
| Kind | yes |
| Discovery method | native, anvien, or both |
| Discovery evidence | yes |
| Verification verdict | yes after verification |
| Verification evidence | yes after verification |
| Dynamic/public risk | yes |

Verification verdicts:

| Verdict | Meaning |
|---|---|
| `confirmed_deadcode` | Strong source evidence shows no live reference, no dynamic/public entrypoint risk, and no known external contract dependency. |
| `likely_deadcode` | Static evidence strongly suggests deadcode, but one low-probability dynamic or external-use risk remains. |
| `uncertain` | Evidence is insufficient or dynamic/public usage risk is material. |
| `false_positive` | Candidate is referenced, generated, required by tests/build/runtime, or intentionally public. |

## Discovery Procedures

### Native Mode

1. Start native-mode token measurement.
2. Deliver the fixed task prompt.
3. Use native filesystem/list/search/read/static-analysis commands only.
4. Exclude dependency folders, build output, cache/index output, generated package output, and benchmark artifacts from candidate discovery; record exclusions.
5. Identify likely declaration surfaces in Go, TypeScript/JavaScript, scripts, generated-context sources, runnable config/spec files, command wrappers, API surfaces, and integration surfaces.
6. For each possible candidate, run native reference searches for exact symbol/name, exported aliases, route/tool names, command names, filenames, and generated references where applicable.
7. Read only source files needed to classify a candidate lead.
8. Reject leads with evidence when they are referenced, generated, test fixtures, public contracts, runtime hooks, CLI/API/MCP entrypoints, or intentionally retained surfaces.
9. Record unresolved leads when dynamic use, reflection, external contract use, or generated references cannot be ruled out.
10. Close native discovery only after every lead is classified as candidate, rejected, or unresolved and every planned source surface has been checked.
11. Close native-mode token measurement and record `agent_session_tokens` or exact proxy tokens.

### Anvien Mode

1. Start Anvien-mode token measurement.
2. Deliver the fixed task prompt.
3. Refresh graph evidence with `anvien analyze --force`.
4. Use Anvien graph commands before broad source reads. Candidate-finding commands may include any Anvien CLI/MCP/resource command appropriate to deadcode discovery.
5. Remember that Anvien local work does not spend tokens. Count only Anvien command output that is delivered to the agent session.
6. If an Anvien command writes large output to disk and the agent does not read it, do not count that hidden output as agent tokens.
7. For each graph lead, inspect Anvien-delivered output first, then read source only where needed to classify the lead.
8. Run native follow-up reference searches only for graph-surfaced candidates or dynamic/public-risk checks; record this cost inside Anvien mode.
9. Reject graph leads with evidence when they are referenced, generated, test fixtures, public contracts, runtime hooks, CLI/API/MCP entrypoints, or intentionally retained surfaces.
10. Record unresolved leads when graph evidence is insufficient or dynamic/external use cannot be ruled out.
11. Close Anvien discovery only after every graph/source lead is classified as candidate, rejected, or unresolved and every planned graph/source surface has been checked.
12. Close Anvien-mode token measurement and record `agent_session_tokens` or exact proxy tokens.

### Shared Verification

1. Merge native and Anvien candidate lists only after both discovery reports are closed.
2. Assign stable candidate ids.
3. Check references, dynamic entrypoint risk, public API risk, generated-code status, tests/build/runtime hooks, and external contract hints.
4. Classify each candidate as `confirmed_deadcode`, `likely_deadcode`, `uncertain`, or `false_positive`.
5. Record shared verification token cost separately. It is not assigned to either discovery mode.

## Definition Of Done

This benchmark is complete only when:

1. baseline commit and worktree state are recorded;
2. token measurement mechanism is selected before discovery;
3. native mode has valid `agent_session_tokens` or exact proxy tokens;
4. Anvien mode has valid `agent_session_tokens` or exact proxy tokens;
5. native candidate list is closed before Anvien discovery starts;
6. Anvien candidate list is closed before shared verification starts;
7. the union of candidates is verified and classified;
8. benchmark ledger contains per-mode token/read/time/candidate metrics;
9. final comparison states token winner/tie/invalid with numeric support;
10. no deadcode is deleted or modified.

## Phase Checklist

- [x] [P0-A] Establish benchmark baseline.
  - Goal: freeze the repo state used for both modes.
  - Work Steps: record `git rev-parse HEAD`; record branch and worktree status; record source-code dirty state; record benchmark docs/reports dirty state; record current date, machine, shell, and tool versions needed for the benchmark.
  - Implementation Gate: do not start token measurement or discovery until baseline is recorded in evidence and benchmark ledgers.
  - Acceptance: evidence contains commit hash and worktree status; benchmark ledger contains baseline environment rows.
  - Current result: complete on 2026-06-02. Final discovery baseline commit is `6516020c323b54c74583fbaa2caf81dad4475036`; branch `master`; worktree clean before baseline evidence was written.

- [x] [P0-B] Establish token measurement mechanism.
  - Goal: guarantee the run measures agent token spend, not Anvien output size.
  - Work Steps: choose runtime token telemetry or exact transcript proxy; define how prompt, tool-call args, delivered tool results, file reads, agent responses, retries, and errors are measured; prove that hidden local Anvien graph/index/cache/output is not counted unless delivered to the agent.
  - Implementation Gate: if both native and Anvien modes cannot be measured with the same valid mechanism, stop and record blocker.
  - Acceptance: evidence contains measurement setup; benchmark ledger marks token comparison eligible before discovery.
  - Current result: complete on 2026-06-02. Primary mechanism is isolated `codex exec --json` sessions, whose `turn.completed.usage` includes `input_tokens`, `output_tokens`, and `reasoning_output_tokens`. Fallback/audit mechanism is `scripts/measure-agent-token-proxy.ps1`, which records delivered transcript events and counts them with Python `tiktoken` using `o200k_base` unless a tokenizer override is provided. Discovery has not started; final benchmark baseline must be recaptured after the measurement harness commit.

- [x] [P1-A] Run native deadcode discovery without Anvien.
  - Goal: measure agent token spend and result quality for native agentic deadcode discovery.
  - Work Steps: run the Native Mode procedure; keep Anvien fully excluded; record commands, delivered outputs, source reads, rejected leads, unresolved leads, candidates, and token metrics.
  - Implementation Gate: no Anvien command, resource, generated context, graph artifact, Anvien output, or prior Anvien report may be used.
  - Acceptance: evidence records native procedure, command/read log, candidate list, completion condition, and native token total.
  - Current result: complete on 2026-06-02. Native discovery ran in isolated `codex exec --json`; `agent_session_tokens=856990`; 39 completed command executions; 12 candidates reported; completed command audit found no Anvien command/resource/graph use.

- [x] [P1-B] Close native discovery report.
  - Goal: freeze native output before Anvien discovery starts.
  - Work Steps: write native candidate table, unresolved questions, rejected-lead summary, token buckets, file reads, search calls, confidence notes, and token validity.
  - Implementation Gate: do not start Anvien mode until native report and native token ledger are closed.
  - Acceptance: native report is complete in evidence and benchmark ledger.
  - Current result: evidence `E2` and `E3` have been filled. Benchmark ledger rows for native mode are being synchronized from the same Codex CLI usage event.

- [x] [P2-A] Run Anvien-guided deadcode discovery.
  - Goal: measure agent token spend and result quality when the agent uses Anvien as a local tool.
  - Work Steps: run the Anvien Mode procedure; run `anvien analyze --force`; use Anvien commands to surface candidates; read source only where needed; record delivered Anvien command output, delivered artifact reads if any, follow-up searches, source reads, candidates, and token metrics.
  - Implementation Gate: graph freshness must be recorded before graph-based work; native candidates must not be read as input; Anvien local graph/index/cache/output must not be counted unless delivered to the agent.
  - Acceptance: evidence records Anvien command/read log, candidate list, rejected leads, unresolved leads, completion condition, and Anvien-mode token total.
  - Current result: complete on 2026-06-02. Successful retry ran in clean worktree `E:\Anvien-benchmark-anvien` at baseline commit to avoid native-result contamination. `agent_session_tokens=1524993`; 30 Anvien commands within 41 completed commands; 6 candidates reported.

- [x] [P2-B] Close Anvien discovery report.
  - Goal: freeze Anvien output before shared verification.
  - Work Steps: write Anvien candidate table, unresolved questions, rejected-lead summary, token buckets, file reads, Anvien command count, follow-up search count, confidence notes, and token validity.
  - Implementation Gate: do not start union verification until Anvien report and Anvien token ledger are closed.
  - Acceptance: Anvien report is complete in evidence and benchmark ledger.
  - Current result: evidence `E4` and `E5` have been filled. Benchmark ledger rows for Anvien mode are being synchronized from the same Codex CLI usage event.

- [x] [P3-A] Verify the union of candidates.
  - Goal: determine true/likely/uncertain/false-positive outcomes for all candidates.
  - Work Steps: merge both candidate lists; dedupe; assign stable ids; check references, public/dynamic risk, generated-code status, tests/build/runtime hooks, and external contract hints; classify every candidate.
  - Implementation Gate: verification cost is shared and must not be assigned to either discovery mode.
  - Acceptance: evidence contains verdict and proof for every candidate; benchmark ledger contains correctness counts by method.
  - Current result: complete on 2026-06-02. Union has 18 candidates: 7 `confirmed_deadcode`, 10 `likely_deadcode`, 0 `uncertain`, 1 `false_positive`.

- [x] [P4-A] Compare native vs Anvien mode.
  - Goal: answer whether using Anvien reduces agent token spend without reducing correctness.
  - Work Steps: compare `agent_session_tokens`, token delta, token reduction ratio, file reads, source bytes, search calls, Anvien calls, elapsed time, candidates found, confirmed/likely/uncertain/false-positive counts, and method-only finds.
  - Implementation Gate: comparison cannot be written until both modes have valid token metrics and every candidate has a verification verdict.
  - Acceptance: final comparison states winner/tie/invalid per axis with numeric support.
  - Current result: complete on 2026-06-02. Native mode wins this run: 856,990 tokens vs Anvien 1,524,993; Anvien spent 668,003 more tokens and found 7 fewer confirmed/likely candidates.

- [x] [P5-A] Close benchmark docs.
  - Goal: keep plan, evidence, and benchmark ledgers synchronized after execution.
  - Work Steps: update checklist statuses; ensure benchmark tables contain quantitative data only; ensure evidence contains command facts and interpretations; record final summary.
  - Implementation Gate: no deadcode edits were made; token ledger is complete for both modes.
  - Acceptance: plan status is complete; evidence and benchmark ledgers are synchronized.
  - Current result: complete on 2026-06-02. No deadcode was edited; benchmark/evidence ledgers are synchronized; final doc commit hash is recorded in the final response because a commit cannot contain its own hash.

## Risk Notes

- A token benchmark without valid agent-session token measurement is invalid.
- Anvien output volume can be useful diagnostic data, but it is not the benchmark score.
- Deadcode detection is false-positive-prone because dynamic calls, CLI entrypoints, tests, reflection, generated code, and external contracts can hide usage.
- Native and Anvien phases can contaminate each other if candidate lists are reused before comparison.
- A method that spends fewer tokens but finds fewer true candidates is not automatically better.
- A method that finds more candidates but produces many false positives may increase verification cost.
