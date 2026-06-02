# Anvien Deadcode Agent Token Benchmark Plan

Date: 2026-06-02

Status: reopened; not started

Companion files:

- Evidence ledger: [2026-06-02-anvien-deadcode-agent-token-benchmark-evidence.md](2026-06-02-anvien-deadcode-agent-token-benchmark-evidence.md)
- Benchmark ledger: [2026-06-02-anvien-deadcode-agent-token-benchmark-benchmark.md](2026-06-02-anvien-deadcode-agent-token-benchmark-benchmark.md)

## Reset Notice

The previous benchmark run is discarded for token-comparison purposes. Its old evidence and benchmark values must not be reused.

Reason: the old run confused full command-output potential size with the context actually received by the agent. A rerun is required with a dedicated token accountant that measures only agent-visible input/output context.

## Master Rules

1. This is an agent-workflow benchmark, not an Anvien analyze-speed benchmark.
2. The benchmark target is one task: find deadcode currently present in this repo.
3. Do not delete, rename, refactor, or edit candidate deadcode during this benchmark.
4. Use one git commit as the benchmark baseline and record the exact hash before discovery starts.
5. If the worktree has source-code changes outside benchmark docs/reports at baseline time, stop and record the blocker before continuing.
6. Run the native-search phase first and do not use Anvien commands, Anvien MCP tools, Anvien resources, generated Anvien context, Anvien graph output, or prior Anvien reports in that phase.
7. Run the Anvien-guided phase second and use Anvien as the discovery aid before reading source files.
8. Keep discovery phases separate. Do not leak candidate lists, files-read logs, commands, reports, or conclusions from one phase into the other until both discovery reports are closed.
9. Use the fixed task prompt in this plan for both discovery methods. If the prompt changes, mark the run invalid and restart both phases.
10. Every candidate must be verified after both discovery phases using source-backed evidence.
11. Candidate verification must classify each candidate as `confirmed_deadcode`, `likely_deadcode`, `uncertain`, or `false_positive`.
12. Verification may use any appropriate source reads and commands; its cost is shared and is not counted as either discovery method's discovery cost.
13. Count only context actually received or emitted by the agent. Tool internal work, graph files, redirected output, capture-only output, and disk artifacts are not agent tokens unless their content is returned to or read by the agent.
14. Any command output from any Anvien command that is returned to the agent is agent-visible context and must be counted. This includes `analyze`, `query`, `context`, `impact`, `cypher`, `graph-health`, `query-health`, `resolution-inventory`, `source-site-accuracy`, `detect-changes`, `api ...`, `group ...`, and any other Anvien command.
15. Anvien internal graph/index/cache work is not agent-visible context and must not be counted as agent token cost.
16. If the token accountant cannot determine exact agent-visible context for a phase, that phase cannot claim a token winner.
17. Do not impose artificial elapsed-time, command/tool-call, file-read, source-byte, or token limits on either method. These are measured outcomes, not stop gates.
18. End a discovery phase only when its declared procedure is complete, or when a blocker is recorded and the run is marked incomplete.
19. Record evidence immediately after each phase and quantitative measurements immediately in the benchmark ledger.
20. No benchmark result is accepted without the token accountant's ledger.

## Goal

Benchmark whether an AI agent uses fewer tokens and reads fewer files when using Anvien to find deadcode, compared with native agentic search without Anvien, while preserving result correctness.

The final comparison must answer:

- Which deadcode candidates did native-search find?
- Which deadcode candidates did Anvien-guided search find?
- Which candidates were found by both, native-only, or Anvien-only?
- Which candidates are confirmed, likely, uncertain, or false positives?
- How much agent-visible context did each method consume and emit?
- Did Anvien reduce file reads, search calls, or total agent-visible tokens?
- Did Anvien improve, match, or reduce correctness?

## Fixed Task Prompt

Use this exact prompt for both discovery methods:

```text
Find deadcode currently present in this repository. Do not edit, delete, rename, or refactor code. Return candidates only. For each candidate, provide path, symbol or file name, kind, why it appears unused, commands or source facts used, and uncertainty or dynamic-use risk.
```

## Core Token Model

The benchmark measures agent workflow cost, not tool internal computation.

```text
agent_visible_total_tokens
= task_prompt_tokens
+ tool_call_argument_tokens
+ tool_result_tokens
+ source_read_tokens
+ search_output_tokens
+ agent_response_tokens
+ retry_error_tokens
```

For Anvien-guided discovery:

```text
anvien_visible_tokens
= visible stdout/stderr/tool-result text from every Anvien command the agent calls
```

Do not count:

- Anvien graph/index/cache data generated internally;
- `.anvien/graph.json` unless the agent reads its content;
- full command output redirected to a file and never read;
- output captured into a variable where only a count/summary is printed;
- output hidden by truncation beyond the part actually visible to the agent.

If the agent sees this text:

```text
full_stdout_proxy_tokens: 228504
```

then only that visible text is counted, not `228504` tokens.

If the agent receives the full output of an Anvien command through a tool result, that full visible output is counted.

## Token Accountant

A dedicated token accountant must run for this benchmark.

Token accountant responsibility:

1. Observe the main agent's visible context for each phase.
2. Record every tool call that the main agent makes.
3. Record which tool outputs were visible to the main agent.
4. Count visible characters and estimated tokens per bucket.
5. Mark any truncated, hidden, redirected, or capture-only output explicitly.
6. State whether exact agent-visible token accounting is valid for the phase.

Token accountant non-goals:

- It must not search for deadcode.
- It must not decide candidate verdicts.
- It must not replace source-backed verification.
- It must not count tool internal data as agent tokens.

Implementation gate:

- Before P1 discovery starts, create or assign the token accountant.
- If no mechanism exists for the accountant to measure agent-visible context, stop and record the blocker instead of running an invalid token benchmark.

## Deadcode Candidate Definition

A candidate is any source declaration, file, package-level object, helper, route/tool handler, or exported surface that appears unused or unreachable under normal repository behavior.

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

### Native-Search Procedure

1. Token accountant starts native phase.
2. Main agent receives the fixed task prompt.
3. Build a source inventory with native file listing/search commands only.
4. Exclude dependency folders, build output, cache/index output, generated package output, and benchmark artifacts from candidate discovery; record exclusions.
5. Identify likely declaration surfaces in Go, TypeScript/JavaScript, scripts, generated-context sources, runnable config/spec files, command wrappers, API surfaces, and integration surfaces.
6. For each possible candidate, run native reference searches for exact symbol/name, exported aliases, route/tool names, command names, filenames, and generated references where applicable.
7. Read only source files needed to classify a candidate lead.
8. Reject leads with evidence when they are referenced, generated, test fixtures, public contracts, runtime hooks, CLI/API/MCP entrypoints, or intentionally retained surfaces.
9. Record unresolved leads when dynamic use, reflection, external contract use, or generated references cannot be ruled out.
10. Close native discovery only after every lead is classified as candidate, rejected, or unresolved and every planned source surface has been checked.
11. Token accountant closes native phase and records visible-context totals.

### Anvien-Guided Procedure

1. Token accountant starts Anvien-guided phase.
2. Main agent receives the fixed task prompt.
3. Refresh graph evidence with `anvien analyze --force` and record visible stdout/stderr.
4. Use Anvien graph commands before broad source reads. Candidate-finding commands may include any Anvien CLI/MCP/resource command appropriate to deadcode discovery.
5. Count every visible Anvien command result as agent-visible context.
6. Do not count Anvien internal graph/index/cache work.
7. For each graph lead, inspect Anvien output first, then read source only where needed to classify the lead.
8. Run native follow-up reference searches only for graph-surfaced candidates or dynamic/public-risk checks; record this cost inside the Anvien-guided phase.
9. Reject graph leads with evidence when they are referenced, generated, test fixtures, public contracts, runtime hooks, CLI/API/MCP entrypoints, or intentionally retained surfaces.
10. Record unresolved leads when graph evidence is insufficient or dynamic/external use cannot be ruled out.
11. Close Anvien-guided discovery only after every graph/source lead is classified as candidate, rejected, or unresolved and every planned graph/source surface has been checked.
12. Token accountant closes Anvien-guided phase and records visible-context totals.

### Shared Verification Procedure

1. Merge native and Anvien candidate lists only after both discovery reports are closed.
2. Assign stable candidate ids.
3. Check references, dynamic entrypoint risk, public API risk, generated-code status, tests/build/runtime hooks, and external contract hints.
4. Classify each candidate as `confirmed_deadcode`, `likely_deadcode`, `uncertain`, or `false_positive`.
5. Token accountant records shared verification visible-context totals separately.

## Discovery Isolation Protocol

Preferred execution:

1. Start from the same clean baseline commit.
2. Run native-search discovery in an isolated transcript/session.
3. Save native discovery report and token accountant ledger.
4. Run Anvien-guided discovery in a separate isolated transcript/session that cannot see native candidates.
5. Save Anvien discovery report and token accountant ledger.
6. Run shared verification/comparison after both reports are closed.

Invalid cases:

- Native-search discovery uses Anvien commands, graph artifacts, resources, generated context, or Anvien output.
- Anvien-guided discovery reads native candidates before its discovery report is closed.
- Either method receives a different task prompt, procedure, or completion standard.
- Token accountant cannot distinguish visible output from hidden/capture-only/internal output.

## Definition Of Done

This benchmark is complete when:

1. baseline commit and worktree state are recorded;
2. token accountant is assigned before discovery;
3. native-search discovery report is recorded with candidate list and visible-context cost breakdown;
4. Anvien-guided discovery report is recorded with candidate list and visible-context cost breakdown;
5. the union of candidates is verified and classified;
6. benchmark ledger contains per-phase agent-visible token/read/time/candidate metrics;
7. final comparison states token result, file-read result, search/tool-call result, and correctness result;
8. no deadcode is deleted or modified.

## Phase Checklist

- [ ] [P0-A] Establish benchmark baseline.
  - Goal: freeze the repo state used for both discovery methods.
  - Work Steps: record `git rev-parse HEAD`; record branch and worktree status; record source-code dirty state; record benchmark docs/reports dirty state; record current date, machine, shell, and tool versions needed for the benchmark.
  - Implementation Gate: do not start discovery until baseline is recorded in evidence and benchmark ledgers.
  - Acceptance: evidence contains commit hash and worktree status; benchmark ledger contains baseline environment rows.

- [ ] [P0-B] Start token accountant.
  - Goal: guarantee the run measures agent-visible context, not tool internal output.
  - Work Steps: create/assign token accountant; define how visible tool outputs, tool-call arguments, source reads, search outputs, Anvien outputs, truncation, and agent responses will be counted; record the counting method before discovery.
  - Implementation Gate: if visible-context accounting cannot be measured, stop and record blocker.
  - Acceptance: evidence contains token accountant setup; benchmark ledger contains token-accounting method row.

- [ ] [P1-A] Run native-search deadcode discovery without Anvien.
  - Goal: measure visible context and result quality for native agentic deadcode discovery.
  - Work Steps: run the Native-Search Procedure; keep Anvien fully excluded; record commands, outputs, reads, rejected leads, unresolved leads, and candidates as observed by token accountant.
  - Implementation Gate: no Anvien command, resource, generated context, graph artifact, or Anvien report may be used.
  - Acceptance: evidence records native procedure, command/read log, candidate list, and completion condition; benchmark ledger records native visible-context metrics.

- [ ] [P1-B] Close native-search discovery report.
  - Goal: freeze native output before Anvien-guided discovery starts.
  - Work Steps: write native candidate table, unresolved questions, rejected-lead summary, visible-token buckets, file reads, search calls, and confidence notes.
  - Implementation Gate: do not start Anvien-guided phase until native report and accountant ledger are closed.
  - Acceptance: native report is complete in evidence and benchmark ledger.

- [ ] [P2-A] Run Anvien-guided deadcode discovery.
  - Goal: measure visible context and result quality when Anvien guides the agent.
  - Work Steps: run the Anvien-Guided Procedure; run `anvien analyze --force`; use Anvien commands to surface candidates; read source only where needed; record all visible Anvien output and follow-up source/search output through the token accountant.
  - Implementation Gate: graph freshness must be recorded before graph-based work; native candidates must not be read as input.
  - Acceptance: evidence records Anvien command/read log, candidate list, rejected leads, unresolved leads, and completion condition; benchmark ledger records Anvien visible-context metrics.

- [ ] [P2-B] Close Anvien-guided discovery report.
  - Goal: freeze Anvien-guided output before shared verification.
  - Work Steps: write Anvien candidate table, unresolved questions, rejected-lead summary, visible-token buckets, file reads, Anvien command count, follow-up search count, and confidence notes.
  - Implementation Gate: do not start union verification until Anvien report and accountant ledger are closed.
  - Acceptance: Anvien report is complete in evidence and benchmark ledger.

- [ ] [P3-A] Verify the union of candidates.
  - Goal: determine true/likely/uncertain/false-positive outcomes for all candidates.
  - Work Steps: merge both candidate lists; dedupe; assign stable ids; check references, public/dynamic risk, generated-code status, tests/build/runtime hooks, and external contract hints; classify every candidate.
  - Implementation Gate: verification cost is shared and must not be assigned to either discovery method.
  - Acceptance: evidence contains verdict and proof for every candidate; benchmark ledger contains correctness counts by method.

- [ ] [P4-A] Compare native-search vs Anvien-guided discovery.
  - Goal: answer whether Anvien reduces token/read/search cost without reducing correctness.
  - Work Steps: compare agent-visible tokens, file reads, source bytes, search calls, Anvien calls, elapsed time, candidates found, confirmed/likely/uncertain/false-positive counts, and method-only finds.
  - Implementation Gate: comparison cannot be written until every candidate has a verification verdict and the token accountant has closed all phases.
  - Acceptance: final comparison states winner/tie/invalid per axis with numeric support.

- [ ] [P5-A] Close benchmark docs.
  - Goal: keep plan, evidence, and benchmark ledgers synchronized after execution.
  - Work Steps: update checklist statuses; ensure benchmark tables contain quantitative data only; ensure evidence contains command facts and interpretations; record final summary.
  - Implementation Gate: no deadcode edits were made; token accountant ledger is complete.
  - Acceptance: plan status is complete; evidence and benchmark ledgers are synchronized.

## Risk Notes

- Deadcode detection is false-positive-prone because dynamic calls, CLI entrypoints, tests, reflection, generated code, and external contracts can hide usage.
- Native and Anvien phases can contaminate each other if candidate lists are reused before comparison.
- Token accounting is invalid if it counts tool internal work or full-output proxies that the agent did not receive.
- Token accounting is invalid if visible output is truncated and the visible character count is not captured.
- A method that reads fewer tokens but finds fewer true candidates is not better.
- A method that finds more candidates but produces many false positives may increase verification cost.
