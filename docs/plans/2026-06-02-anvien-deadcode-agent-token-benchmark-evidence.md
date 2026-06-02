# Anvien Deadcode Agent Token Benchmark Evidence Ledger

Date: 2026-06-02

Status: measurement harness ready; benchmark baseline pending

Companion files:

- Plan: [2026-06-02-anvien-deadcode-agent-token-benchmark-plan.md](2026-06-02-anvien-deadcode-agent-token-benchmark-plan.md)
- Benchmark ledger: [2026-06-02-anvien-deadcode-agent-token-benchmark-benchmark.md](2026-06-02-anvien-deadcode-agent-token-benchmark-benchmark.md)

## Reset Notice

All prior evidence from the invalid run is discarded.

Reason: the old run did not correctly measure AI-agent token spend in Anvien mode. Anvien local output volume is not the same thing as agent token usage.

Do not reuse old command logs, old candidate lists, old token totals, old verification tables, or old conclusions.

## Evidence Rules

1. Evidence explains why the benchmark result is valid.
2. Keep quantitative metric tables in the benchmark ledger.
3. Native mode evidence must not include Anvien commands, outputs, resources, generated context, graph artifacts, or prior Anvien reports.
4. Anvien mode evidence must record graph freshness before graph-based commands.
5. Token evidence must prove the AI agent's token usage measurement mechanism, not Anvien output size.
6. Record delivered tool results and file reads only when they enter the AI agent session or exact transcript proxy.
7. Record unobserved local tool artifacts as unobserved; do not turn them into token usage.
8. Record failures, retries, and abandoned paths because they can consume agent tokens when delivered to the agent.
9. Record candidate verification from source-backed facts, not intuition.
10. Do not record deadcode deletion or cleanup patches because this benchmark only finds candidates.

## Evidence Template

Use this template for each phase:

```text
## E<n> - <Phase>

Date:

Status:

Scope:

- ...

Commands / reads:

| Step | Command or file | Purpose | Delivered to agent? | Token-measurement note | Result |
|---|---|---|---|---|---|
| ... | ... | ... | ... | ... | ... |

Candidate evidence:

| Candidate id | Path | Symbol/name | Kind | Discovery evidence | Notes |
|---|---|---|---|---|---|
| ... | ... | ... | ... | ... | ... |

Failures / retries:

- ...

Completion:

| Item | Result |
|---|---|
| Declared procedure recorded before discovery | pending |
| Token measurement active | pending |
| Completion condition met | pending |
| Open leads remaining | pending |
| Blocker or incomplete reason | pending |
| Confidence | pending |
```

## E0 - Baseline

Date: 2026-06-02T13:06:24.1784982+07:00

Status: superseded by measurement-harness implementation; final benchmark baseline pending

Required evidence:

| Check | Result |
|---|---|
| Baseline commit | `b5568094184b92618e2627460c5a2f22b120a497` pre-harness baseline; not the final benchmark discovery baseline. |
| Branch | `master` |
| Worktree status | Dirty only in benchmark docs at pre-harness baseline: `M docs/plans/2026-06-02-anvien-deadcode-agent-token-benchmark-benchmark.md`, `M docs/plans/2026-06-02-anvien-deadcode-agent-token-benchmark-evidence.md`, `M docs/plans/2026-06-02-anvien-deadcode-agent-token-benchmark-plan.md` |
| Source-code dirty state | Clean at pre-harness baseline. Final benchmark baseline must be captured after committing the measurement harness. |
| Benchmark docs/reports dirty state | 3 dirty benchmark docs/reports at pre-harness baseline. |
| Shell | PowerShell 7.6.2 |
| OS / CPU / RAM | Microsoft Windows 10 Pro 10.0.19045 64-bit; Intel(R) Core(TM) i7-3770 CPU @ 3.40GHz; 4 cores / 8 logical processors; 32,459,440 KiB visible memory. |
| Go / Node / npm / Git versions if used | `go version go1.26.3 windows/amd64`; `node v24.15.0`; `npm 11.12.1`; `git version 2.54.0.windows.1` |

Commands / reads:

| Step | Command or file | Purpose | Delivered to agent? | Token-measurement note | Result |
|---|---|---|---|---|---|
| E0-1 | `git rev-parse HEAD` | Record baseline commit. | yes | Baseline evidence only; discovery not started. | `b5568094184b92618e2627460c5a2f22b120a497` |
| E0-2 | `git branch --show-current` | Record branch. | yes | Baseline evidence only; discovery not started. | `master` |
| E0-3 | `git status --porcelain=v1` | Record dirty state. | yes | Baseline evidence only; discovery not started. | Three modified benchmark docs; no dirty source files. |
| E0-4 | `git --version` | Record tool version. | yes | Baseline evidence only; discovery not started. | `git version 2.54.0.windows.1` |
| E0-5 | `go version` | Record tool version. | yes | Baseline evidence only; discovery not started. | `go version go1.26.3 windows/amd64` |
| E0-6 | `node --version` | Record tool version. | yes | Baseline evidence only; discovery not started. | `v24.15.0` |
| E0-7 | `npm --version` | Record tool version. | yes | Baseline evidence only; discovery not started. | `11.12.1` |
| E0-8 | `$PSVersionTable.PSVersion.ToString()` | Record shell version. | yes | Baseline evidence only; discovery not started. | `7.6.2` |
| E0-9 | `Get-Date -Format o` | Record local timestamp. | yes | Baseline evidence only; discovery not started. | `2026-06-02T13:06:24.1784982+07:00` |
| E0-10 | `Get-CimInstance Win32_OperatingSystem ...` | Record OS/RAM. | yes | Baseline evidence only; discovery not started. | Windows 10 Pro 10.0.19045 64-bit; 32,459,440 KiB visible memory. |
| E0-11 | `Get-CimInstance Win32_Processor ...` | Record CPU. | yes | Baseline evidence only; discovery not started. | Intel(R) Core(TM) i7-3770 CPU @ 3.40GHz; 4 cores / 8 logical processors. |

## E1 - Token Measurement Setup

Date: 2026-06-02

Status: complete

Required evidence:

| Item | Result |
|---|---|
| Measurement mechanism | Primary: isolated `codex exec --json` sessions, using `turn.completed.usage.input_tokens + turn.completed.usage.output_tokens` as `agent_session_tokens`. Fallback/audit: `scripts/measure-agent-token-proxy.ps1` transcript proxy with Python `tiktoken` using `o200k_base` unless overridden. |
| Exact model/runtime token telemetry available | Yes for Codex CLI sessions. Probe output included `usage.input_tokens`, `usage.cached_input_tokens`, `usage.output_tokens`, and `usage.reasoning_output_tokens`. |
| Exact transcript proxy available if telemetry is not available | Yes as a fallback for harnessed commands/files/responses: `scripts/measure-agent-token-proxy.ps1` records NDJSON events and counts delivered content through `tiktoken`. It is not used to infer hidden model telemetry. |
| Can measure native mode with this mechanism | Yes, by running native discovery in its own `codex exec --json` session and parsing the usage event. |
| Can measure Anvien mode with this mechanism | Yes, by running Anvien-guided discovery in a separate `codex exec --json` session and parsing the usage event. |
| Can distinguish agent-session tokens from Anvien local output volume | Yes. Codex CLI usage is runtime model usage, and the proxy script marks undelivered/local output separately from delivered transcript content. |
| Can exclude hidden graph/cache/index/output not delivered to agent | Yes. Hidden local artifacts are not part of Codex CLI model usage unless the measured agent reads/receives them; the proxy script records undelivered output under `local_tool_output_volume`. |
| Blocker if token measurement is unavailable | No current blocker. Discovery has not started because the harness must be committed and the final benchmark baseline recaptured first. |

Commands / reads:

| Step | Command or file | Purpose | Delivered to agent? | Token-measurement note | Result |
|---|---|---|---|---|---|
| E1-1 | `get_goal` tool | Check whether the active runtime exposes goal/token usage telemetry. | yes | This confirmed no usable runtime token telemetry for the benchmark phases. | `goal: null`; `remainingTokens: null`; `completionBudgetReport: null` |
| E1-2 | `codex exec --json --cd E:\Anvien --sandbox danger-full-access "Reply with exactly: TOKEN_PROBE_OK"` | Check isolated Codex CLI usage telemetry. | yes | Primary benchmark mechanism; measured usage belongs to the isolated Codex CLI session. | `input_tokens=17510`, `cached_input_tokens=3456`, `output_tokens=26`, `reasoning_output_tokens=16`. |
| E1-3 | `python -m pip install --target .tmp\tokenizer-python tiktoken` | Install tokenizer for transcript proxy fallback. | yes | Enables exact proxy counts for delivered transcript text. | Installed `tiktoken 0.13.0` plus dependencies under `.tmp\tokenizer-python`. |
| E1-4 | `scripts\measure-agent-token-proxy.ps1` smoke sequence | Verify transcript proxy with `tiktoken`. | yes | Fallback/audit mechanism produced exact token counts. | Smoke summary reported `agent_session_token_proxy=56` and `token_measurement_valid=true`. |

Completion:

| Item | Result |
|---|---|
| Declared procedure recorded before discovery | yes |
| Token measurement active | yes, for future isolated Codex CLI discovery sessions |
| Completion condition met | yes |
| Open leads remaining | not applicable; discovery did not start |
| Blocker or incomplete reason | none for token measurement; final benchmark baseline still pending after harness commit |
| Confidence | high |

## E2 - Native Discovery Without Anvien

Date:

Status: pending

Rules for this section:

- Do not use Anvien.
- Record every delivered command result and source file read that enters the agent session/proxy.
- Record candidates before any Anvien-mode work starts.

Native command/read log:

| Step | Command or file | Purpose | Delivered to agent? | Token-measurement note | Result |
|---|---|---|---|---|---|
| pending | pending | pending | pending | pending | pending |

Native candidates:

| Candidate id | Path | Symbol/name | Kind | Discovery evidence | Notes |
|---|---|---|---|---|---|
| pending | pending | pending | pending | pending | pending |

Native completion:

| Item | Result |
|---|---|
| Declared native procedure recorded before first search | pending |
| Token measurement closed native phase | pending |
| Completion condition met | pending |
| Open native leads remaining | pending |
| Blocker or incomplete reason | pending |
| Confidence | pending |

## E3 - Native Discovery Report

Date:

Status: pending

Required evidence:

| Item | Result |
|---|---|
| Native candidate count | pending |
| Native unique files read | pending |
| Native command/search count | pending |
| Native `agent_session_tokens` or exact proxy tokens | pending |
| Native token validity | pending |
| Native completion status | pending |
| Native unresolved questions | pending |

## E4 - Anvien-Guided Discovery

Date:

Status: pending

Rules for this section:

- Record graph freshness before graph-based work.
- Do not seed discovery from the native candidate list.
- Record delivered Anvien command output separately from Anvien local output volume.
- Do not count Anvien graph/cache/index/output unless delivered into the agent session/proxy.
- Record source reads after Anvien narrows candidate leads.

Graph freshness:

| Check | Result |
|---|---|
| Analyze command | pending |
| Analyze local-output token status | local tool work; not agent tokens unless delivered |
| Delivered analyze output to agent | pending |
| Indexed commit | pending |
| Current commit | pending |
| Fresh/stale result | pending |

Anvien command/read log:

| Step | Command or file | Purpose | Delivered to agent? | Token-measurement note | Result |
|---|---|---|---|---|---|
| pending | pending | pending | pending | pending | pending |

Anvien candidates:

| Candidate id | Path | Symbol/name | Kind | Discovery evidence | Notes |
|---|---|---|---|---|---|
| pending | pending | pending | pending | pending | pending |

Anvien completion:

| Item | Result |
|---|---|
| Declared Anvien procedure recorded before first graph command | pending |
| Token measurement closed Anvien phase | pending |
| Completion condition met | pending |
| Open Anvien leads remaining | pending |
| Blocker or incomplete reason | pending |
| Confidence | pending |

## E5 - Anvien Discovery Report

Date:

Status: pending

Required evidence:

| Item | Result |
|---|---|
| Anvien candidate count | pending |
| Anvien unique files read | pending |
| Anvien command count | pending |
| Anvien follow-up native search count | pending |
| Anvien `agent_session_tokens` or exact proxy tokens | pending |
| Anvien token validity | pending |
| Anvien completion status | pending |
| Anvien unresolved questions | pending |

## E6 - Candidate Union And Verification

Date:

Status: pending

Verification rules:

- Verify the union of candidates from both modes.
- Check static references, dynamic/public entrypoint risk, generated-code status, test/build/runtime hooks, and external contract hints.
- Do not delete or edit candidate code.

Candidate verdicts:

| Candidate id | Found by native | Found by Anvien | Path | Symbol/name | Verdict | Verification evidence | Dynamic/public risk |
|---|---|---|---|---|---|---|---|
| pending | pending | pending | pending | pending | pending | pending | pending |

False positives:

| Candidate id | Method source | Reason |
|---|---|---|
| pending | pending | pending |

Uncertain candidates:

| Candidate id | Method source | Uncertainty reason | Follow-up needed |
|---|---|---|---|
| pending | pending | pending | pending |

## E7 - Final Comparison Evidence

Date:

Status: pending

Required comparison facts:

| Question | Evidence |
|---|---|
| How many tokens did the agent spend without Anvien? | pending |
| How many tokens did the agent spend with Anvien? | pending |
| Which mode used fewer agent-session tokens? | pending |
| Which mode read fewer files? | pending |
| Which mode used fewer search/tool calls? | pending |
| Which mode found more confirmed/likely deadcode? | pending |
| Which mode produced fewer false positives? | pending |
| Which candidates were found by both/native-only/Anvien-only? | pending |
| Was token measurement valid for both modes? | pending |

Required summary shape:

```text
Native mode:
- agent_session_tokens or proxy:
- token validity:
- search/tool calls:
- file reads:
- candidates:

Anvien mode:
- agent_session_tokens or proxy:
- token validity:
- Anvien calls:
- follow-up search/tool calls:
- file reads:
- candidates:

Shared verification:
- candidates verified:
- confirmed/likely/uncertain/false-positive:
```

## E8 - Closure

Date:

Status: pending

Closure checks:

| Check | Result |
|---|---|
| No deadcode deletion/edit was made | pending |
| Token measurement valid for native mode | pending |
| Token measurement valid for Anvien mode | pending |
| Plan checklist updated | pending |
| Benchmark ledger complete | pending |
| Final comparison written | pending |
| Commit hash for documentation update, if committed | pending |
