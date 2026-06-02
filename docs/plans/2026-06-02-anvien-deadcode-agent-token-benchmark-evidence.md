# Anvien Deadcode Agent Token Benchmark Evidence Ledger

Date: 2026-06-02

Status: reset; not started

Companion files:

- Plan: [2026-06-02-anvien-deadcode-agent-token-benchmark-plan.md](2026-06-02-anvien-deadcode-agent-token-benchmark-plan.md)
- Benchmark ledger: [2026-06-02-anvien-deadcode-agent-token-benchmark-benchmark.md](2026-06-02-anvien-deadcode-agent-token-benchmark-benchmark.md)

## Reset Notice

Old evidence from the invalidated benchmark run has been removed. This ledger starts clean for the rerun.

## Evidence Rules

1. Record commands, source reads, token-accountant observations, and candidate evidence as they happen.
2. Keep quantitative metric tables in the benchmark ledger.
3. Keep native-search and Anvien-guided discovery separate until union verification.
4. Native-search evidence must not include Anvien commands, outputs, resources, generated context, graph artifacts, or prior Anvien reports.
5. Anvien-guided evidence must record graph freshness before graph-based commands.
6. Token evidence must record only context actually visible to the main agent.
7. Record every visible Anvien command output when Anvien is used.
8. Record hidden, redirected, capture-only, and truncated output as non-counted or partially counted with the reason.
9. Record failures, retries, and abandoned paths because they consume agent work.
10. Record candidate verification from source-backed facts, not intuition.
11. Do not record deadcode deletion or cleanup patches because this benchmark only finds candidates.

## Evidence Template

Use this template for each phase:

```text
## E<n> - <Phase>

Date:

Status:

Scope:

- ...

Commands / reads:

| Step | Command or file | Purpose | Visible to main agent? | Token-accountant note | Result |
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
| Token accountant active | pending |
| Completion condition met | pending |
| Open leads remaining | pending |
| Blocker or incomplete reason | pending |
| Confidence | pending |
```

## E0 - Baseline

Date:

Status: pending

Required evidence:

| Check | Result |
|---|---|
| Baseline commit | pending |
| Branch | pending |
| Worktree status | pending |
| Source-code dirty state | pending |
| Benchmark docs/reports dirty state | pending |
| Shell | pending |
| OS / CPU / RAM | pending |
| Go / Node / npm / Git versions if used | pending |

## E1 - Token Accountant Setup

Date:

Status: pending

Required evidence:

| Item | Result |
|---|---|
| Accountant identity/mechanism | pending |
| Can observe main-agent visible tool results | pending |
| Can distinguish full stdout proxy vs visible output | pending |
| Can record tool-call argument text | pending |
| Can record source/file reads | pending |
| Can record agent response text | pending |
| Truncation handling rule | pending |
| Blocker if exact visible-context accounting is unavailable | pending |

## E2 - Native-Search Discovery Without Anvien

Date:

Status: pending

Rules for this section:

- Do not use Anvien.
- Record every visible command output and source file read.
- Record candidates before any Anvien-guided work starts.

Native command/read log:

| Step | Command or file | Purpose | Visible to main agent? | Token-accountant note | Result |
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
| Token accountant closed native phase | pending |
| Completion condition met | pending |
| Open native leads remaining | pending |
| Blocker or incomplete reason | pending |
| Confidence | pending |

## E3 - Native-Search Discovery Report

Date:

Status: pending

Required evidence:

| Item | Result |
|---|---|
| Native candidate count | pending |
| Native unique files read | pending |
| Native command/search count | pending |
| Native agent-visible token total | pending |
| Native completion status | pending |
| Native unresolved questions | pending |

## E4 - Anvien-Guided Discovery

Date:

Status: pending

Rules for this section:

- Record graph freshness before graph-based work.
- Do not seed discovery from the native candidate list.
- Record every Anvien command and every visible Anvien output.
- Record source reads after Anvien narrows candidate leads.

Graph freshness:

| Check | Result |
|---|---|
| Analyze command | pending |
| Analyze output visible to main agent | pending |
| Indexed commit | pending |
| Current commit | pending |
| Fresh/stale result | pending |

Anvien command/read log:

| Step | Command or file | Purpose | Visible to main agent? | Token-accountant note | Result |
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
| Token accountant closed Anvien phase | pending |
| Completion condition met | pending |
| Open Anvien leads remaining | pending |
| Blocker or incomplete reason | pending |
| Confidence | pending |

## E5 - Anvien-Guided Discovery Report

Date:

Status: pending

Required evidence:

| Item | Result |
|---|---|
| Anvien candidate count | pending |
| Anvien unique files read | pending |
| Anvien command count | pending |
| Anvien follow-up native search count | pending |
| Anvien agent-visible token total | pending |
| Anvien completion status | pending |
| Anvien unresolved questions | pending |

## E6 - Candidate Union And Verification

Date:

Status: pending

Verification rules:

- Verify the union of candidates from both methods.
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
| Which method used fewer agent-visible tokens? | pending |
| Which method read fewer files? | pending |
| Which method used fewer search/tool calls? | pending |
| Which method found more confirmed/likely deadcode? | pending |
| Which method produced fewer false positives? | pending |
| Which candidates were found by both/native-only/Anvien-only? | pending |
| Was the token accountant able to measure exact visible context? | pending |
| Were hidden/internal/redirected/capture-only outputs excluded? | pending |

Required summary shape:

```text
Native search:
- agent-visible total tokens:
- task prompt:
- tool call arguments:
- search output:
- file reads:
- agent response:
- retry/error:
- files read:
- candidates:

Anvien-guided:
- agent-visible total tokens:
- task prompt:
- tool call arguments:
- Anvien visible output:
- follow-up search output:
- file reads:
- agent response:
- retry/error:
- files read:
- candidates:

Shared verification:
- agent-visible total tokens:
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
| Token accountant ledger complete | pending |
| Plan checklist updated | pending |
| Benchmark ledger complete | pending |
| Final comparison written | pending |
| Commit hash for documentation update, if committed | pending |
