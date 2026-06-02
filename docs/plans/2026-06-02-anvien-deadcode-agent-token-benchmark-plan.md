# Anvien Deadcode Agent Token Benchmark Plan

Date: 2026-06-02

Status: Ready to execute

Companion files:

- Evidence ledger: [2026-06-02-anvien-deadcode-agent-token-benchmark-evidence.md](2026-06-02-anvien-deadcode-agent-token-benchmark-evidence.md)
- Benchmark ledger: [2026-06-02-anvien-deadcode-agent-token-benchmark-benchmark.md](2026-06-02-anvien-deadcode-agent-token-benchmark-benchmark.md)

## Master Rules

1. This is a new agent-workflow benchmark, not a speed/analyze benchmark.
2. The benchmark target is one task: find deadcode currently present in this repo.
3. Do not delete, rename, refactor, or edit candidate deadcode during this benchmark.
4. Use one git commit as the benchmark baseline and record the exact hash before discovery starts.
5. If the worktree has source-code changes outside benchmark docs/reports at baseline time, stop and record the blocker before continuing.
6. Run the native-search phase first and do not use Anvien commands, Anvien MCP tools, Anvien resources, generated Anvien context, or Anvien graph output in that phase.
7. Run the Anvien-guided phase second and use Anvien as the discovery aid before reading source files.
8. Keep the two discovery phases separate. Do not leak candidate lists, files, or conclusions from the native phase into the Anvien-guided discovery phase until comparison/verification.
9. Measure the whole agent session, but break token and read cost down by source: task prompt, search output, file reads, graph tool output, agent response, validation output, and retry/error output.
10. Use an explicit token estimate when exact model token accounting is unavailable: `estimated_tokens = ceil(characters / 4)`.
11. A lower token count is not a win if the method finds fewer true candidates or produces more false positives.
12. Every candidate must be verified after both discovery phases using source-backed evidence.
13. Candidate verification must classify each candidate as `confirmed_deadcode`, `likely_deadcode`, `uncertain`, or `false_positive`.
14. Verification may use any appropriate source reads and commands; the verification phase is shared and is not counted as either discovery method unless explicitly recorded as shared validation cost.
15. Record evidence immediately after each phase and quantitative measurements immediately in the benchmark ledger.
16. Count only data the agent actually sees or reads into context. Anvien internal processing, graph files, JSON artifacts, or redirected command output do not cost agent tokens until the agent receives their stdout/stderr or reads those files.

## Goal

Benchmark whether an AI agent uses fewer tokens and reads fewer files when using Anvien to find deadcode, compared with native agentic search without Anvien, while preserving result correctness.

The final comparison must answer:

- Does Anvien-guided discovery reduce estimated token cost?
- Does Anvien-guided discovery reduce file reads and search calls?
- Does Anvien-guided discovery find the same, more, or fewer valid deadcode candidates?
- Does Anvien-guided discovery increase or reduce false positives?
- Which token sources dominate each method?

## Problem

Agentic code search can spend many tokens on broad file listing, repeated text search, and reading unrelated source files before it identifies relevant owners. A graph-guided method may reduce that search space, but this must be measured on a real task instead of assumed.

The chosen benchmark task is deadcode discovery because it naturally requires:

- finding declarations;
- proving reference absence or weak usage;
- checking dynamic or public entrypoint risks;
- reading enough surrounding code to avoid false positives.

## Scope

In scope:

- current local `E:\Anvien` repository;
- deadcode candidate discovery only;
- native-search phase without Anvien;
- Anvien-guided phase with Anvien;
- estimated token accounting by phase and source;
- file-read and command-count accounting;
- candidate verification from source evidence;
- final comparison report after execution.

Out of scope:

- deleting deadcode;
- implementing cleanup patches;
- optimizing lookup latency;
- benchmarking analyze speed;
- comparing Anvien with another code graph product;
- claiming exact model-token usage when only proxy accounting is available.

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

## Token And Read Accounting

Track whole-session cost with separate buckets:

```text
session_total_estimated_tokens
= task_prompt_tokens
+ search_output_tokens
+ file_read_tokens
+ graph_tool_output_tokens
+ agent_response_tokens
+ validation_output_tokens
+ retry_error_tokens
```

Native-search discovery must populate `search_output_tokens` and `file_read_tokens`, but `graph_tool_output_tokens` must be zero.

Anvien-guided discovery must populate `graph_tool_output_tokens`; it may also populate `search_output_tokens` and `file_read_tokens` for follow-up source checks during that discovery phase.

Important accounting rule:

```text
graph_tool_output_tokens
= Anvien command/tool/resource output actually returned to the agent context
  or Anvien-generated artifact content later read by the agent.

graph_tool_output_tokens
!= all data Anvien computes internally
!= .anvien/graph.json size
!= benchmark JSON/output files redirected to disk and never read by the agent
```

Examples:

```text
anvien query ... > .tmp/out.json
```

If the agent only sees process status, the `.tmp/out.json` content is not counted yet.

```text
Get-Content .tmp/out.json
```

Now the read content is counted because it enters the agent context.

Use byte and character counts as supporting metrics:

- command stdout/stderr characters;
- file bytes read into the agent context;
- number of unique files read;
- number of search/list commands;
- number of Anvien commands/tools;
- elapsed wall seconds per phase.

## Measurement Phases

Each method must be measured through the same phase shape.

| Phase | Native-search measurement | Anvien-guided measurement |
|---|---|---|
| Task prompt | Same task prompt where possible. | Same task prompt where possible. |
| Discovery/search | `list`, `grep`, `read`, search output. | Graph query/context/impact output plus any follow-up search output. |
| Context gathering | File content read by the agent. | File content read after graph guidance. |
| Decision/response | Agent response that states candidate owners and why they may be deadcode. | Agent response that states candidate owners and why they may be deadcode. |
| Validation | Source-backed checks that prove or disprove candidates; pass/fail and output tokens are recorded. | Source-backed checks that prove or disprove candidates; pass/fail and output tokens are recorded. |
| Retry/error | Failed commands, detours, or repeated reads. | Failed commands, detours, or repeated reads. |

Required result shape:

```text
Native search:
- total tokens:
- search output:
- file read:
- graph tool output: 0
- agent response:
- validation/retry:
- files read:
- correct:

Anvien-guided:
- total tokens:
- search output:
- file read:
- graph tool output:
- agent response:
- validation/retry:
- files read:
- correct:
```

The benchmark proves value only if the breakdown explains the total. For example, Anvien-guided discovery should show whether graph tool output is offset by lower `file_read_tokens` and `search_output_tokens`.

## Benchmark Methodology

The benchmark has four execution stages:

1. Baseline setup.
2. Native-search deadcode discovery.
3. Anvien-guided deadcode discovery.
4. Shared correctness verification and comparison.

Native-search allowed operations:

- `git` status/diff/log checks;
- `rg`, `rg --files`, `Get-ChildItem`, source reads;
- language/toolchain commands used only as native search or verification aids;
- no Anvien CLI, MCP, resource, generated context, or graph artifact reads.

Anvien-guided allowed operations:

- `anvien analyze --force` before graph-based work;
- Anvien `query`, `context`, `impact`, `detect-changes`, graph-quality commands if useful;
- source reads after Anvien narrows candidates;
- native search commands as follow-up checks, recorded separately.

The comparison must not reward a method for avoiding verification. Discovery cost and shared verification cost must be reported separately.

## Definition Of Done

This benchmark is complete when:

1. baseline commit and worktree state are recorded;
2. native-search discovery report is recorded with candidate list and cost breakdown;
3. Anvien-guided discovery report is recorded with candidate list and cost breakdown;
4. the union of candidates is verified and classified;
5. benchmark ledger contains per-phase token/read/time/candidate metrics;
6. final comparison states which method used fewer tokens, which method read fewer files, and whether correctness changed;
7. no deadcode is deleted or modified.

## Phase Checklist

- [ ] [P0-A] Establish benchmark baseline.
  - Goal: freeze the repo state used for both discovery methods.
  - Work Steps: record `git rev-parse HEAD`; record branch and worktree status; record excluded dirty files if any are benchmark docs/reports only; record current date, machine, shell, and tool versions needed for the benchmark.
  - Implementation Gate: do not start discovery until the baseline is recorded in evidence and benchmark ledgers.
  - Acceptance: evidence contains commit hash and worktree status; benchmark ledger contains baseline environment rows.

- [ ] [P1-A] Run native-search deadcode discovery without Anvien.
  - Goal: measure how much token/read/search work an agent spends finding deadcode without Anvien.
  - Work Steps: use native list/search/source-read commands only; record every command, output-size estimate, files read, and candidate found; stop after a defined discovery budget or after no new candidates appear for the recorded stop condition.
  - Implementation Gate: no Anvien command, MCP tool, resource, generated context, or graph artifact may be used in this phase.
  - Acceptance: evidence records the native process and candidate list; benchmark ledger records native token/read/search/time/candidate counts.

- [ ] [P1-B] Write native-search discovery report.
  - Goal: make the native result auditable before the Anvien-guided phase starts.
  - Work Steps: summarize native commands, files read, token buckets, candidate table, and unresolved questions; do not verify against Anvien or compare with Anvien yet.
  - Implementation Gate: do not start the Anvien-guided phase until the native report exists.
  - Acceptance: native report section is complete in evidence; benchmark rows for native discovery are filled.

- [ ] [P2-A] Run Anvien-guided deadcode discovery.
  - Goal: measure how much token/read/search work an agent spends finding deadcode when Anvien guides discovery.
  - Work Steps: refresh graph with `anvien analyze --force`; use Anvien query/context/impact or related commands to identify candidates; read source only where needed; record `graph_tool_output` size, source file reads, native follow-up searches, elapsed time, and candidates found.
  - Implementation Gate: Anvien graph freshness must be recorded before any graph-based command; do not use native phase candidate lists as input.
  - Acceptance: evidence records Anvien commands and candidate list; benchmark ledger records Anvien token/read/search/time/candidate counts.

- [ ] [P2-B] Write Anvien-guided discovery report.
  - Goal: make the Anvien-guided result auditable before shared verification and comparison.
  - Work Steps: summarize Anvien commands, files read, token buckets, candidate table, and unresolved questions; do not merge native candidates yet except in the later verification phase.
  - Implementation Gate: do not start union verification until the Anvien-guided report exists.
  - Acceptance: Anvien report section is complete in evidence; benchmark rows for Anvien discovery are filled.

- [ ] [P3-A] Verify the union of deadcode candidates.
  - Goal: determine which candidates from either method are true, likely, uncertain, or false positives.
  - Work Steps: merge native and Anvien candidate lists; assign stable candidate ids; check references, dynamic entrypoint risk, public API risk, generated-code status, tests/build/runtime hooks, and external contract hints; classify each candidate.
  - Implementation Gate: verification may use any source-backed method, but its cost must be recorded as shared verification, not attributed to one discovery method unless explicitly required by that method's discovery report.
  - Acceptance: evidence contains verdict and proof for every candidate; benchmark ledger contains correctness counts by method.

- [ ] [P4-A] Compare native-search vs Anvien-guided discovery.
  - Goal: answer whether Anvien reduces token/read cost for this deadcode task without reducing correctness.
  - Work Steps: compare total and bucketed estimated tokens, file reads, bytes read, search calls, tool calls, elapsed time, candidates found, true candidates, false positives, uncertain candidates, and method-only finds.
  - Implementation Gate: comparison cannot be written until every candidate has a verification verdict.
  - Acceptance: final comparison section states winner/loser/mixed result with numeric support and correctness caveats.

- [ ] [P5-A] Close benchmark docs.
  - Goal: make the plan, evidence, and benchmark files consistent after execution.
  - Work Steps: update checklist statuses; ensure benchmark tables contain only numeric/status data; ensure evidence contains commands and interpretations; add final summary links if a separate report is created.
  - Implementation Gate: do not mark complete until no deadcode edits were made and all outputs are recorded.
  - Acceptance: plan status is complete; evidence and benchmark ledgers are synchronized; final report path is recorded if created.

## Risk Notes

- Deadcode detection is inherently false-positive-prone because dynamic calls, CLI entrypoints, tests, reflection, generated code, and external contracts can hide usage.
- Native and Anvien phases can contaminate each other if candidate lists are reused before the comparison stage.
- Exact model-token accounting may be unavailable; the benchmark must label proxy token estimates clearly.
- A method that reads fewer tokens but finds fewer true candidates is not better.
- A method that finds more candidates but produces many false positives may increase verification cost.
