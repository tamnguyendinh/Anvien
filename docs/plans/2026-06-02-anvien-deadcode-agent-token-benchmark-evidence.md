# Anvien Deadcode Agent Token Benchmark Evidence Ledger

Date: 2026-06-02

Status: Not started

Companion files:

- Plan: [2026-06-02-anvien-deadcode-agent-token-benchmark-plan.md](2026-06-02-anvien-deadcode-agent-token-benchmark-plan.md)
- Benchmark ledger: [2026-06-02-anvien-deadcode-agent-token-benchmark-benchmark.md](2026-06-02-anvien-deadcode-agent-token-benchmark-benchmark.md)

## Evidence Rules

1. Record commands, source reads, and candidate evidence as they happen.
2. Keep quantitative metric tables in the benchmark ledger.
3. Keep discovery phases separate until union verification.
4. Native-search evidence must not include Anvien commands, outputs, resources, generated context, or graph artifacts.
5. Anvien-guided evidence must record graph freshness before graph-based commands.
6. Record failures, retries, and abandoned paths because they consume agent work.
7. Record candidate verification from source-backed facts, not intuition.
8. Do not record deadcode deletion or cleanup patches because this benchmark only finds candidates.
9. For token evidence, count only output/content that entered the agent context. Do not count Anvien internal processing, graph files on disk, redirected artifacts, or command output files unless the agent reads them.

## Evidence Template

Use this template for each phase:

```text
## E<n> - <Phase>

Date:

Status:

Scope:

- ...

Commands / reads:

| Step | Command or file | Purpose | Output artifact | Notes |
|---|---|---|---|---|
| ... | ... | ... | ... | ... |

Candidate evidence:

| Candidate id | Path | Symbol/name | Kind | Discovery evidence | Notes |
|---|---|---|---|---|---|
| ... | ... | ... | ... | ... | ... |

Failures / retries:

- ...
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

## E1 - Native-Search Discovery Without Anvien

Date:

Status: pending

Rules for this section:

- Do not use Anvien.
- Record every command and every source file read.
- Record candidates before any Anvien-guided work starts.

Native command/read log:

| Step | Command or file | Purpose | Output artifact | Notes |
|---|---|---|---|---|
| pending | pending | pending | pending | pending |

Native candidates:

| Candidate id | Path | Symbol/name | Kind | Discovery evidence | Notes |
|---|---|---|---|---|---|
| pending | pending | pending | pending | pending | pending |

Native stop condition:

- pending

## E2 - Native-Search Discovery Report

Date:

Status: pending

Required evidence:

| Item | Result |
|---|---|
| Native candidate count | pending |
| Native unique files read | pending |
| Native command/search count | pending |
| Native estimated token total | pending |
| Native unresolved questions | pending |

## E3 - Anvien-Guided Discovery

Date:

Status: pending

Rules for this section:

- Record graph freshness before Anvien graph use.
- Do not seed discovery from the native candidate list.
- Record every Anvien command/output and every source file read.

Graph freshness:

| Check | Result |
|---|---|
| Analyze command | pending |
| Indexed commit | pending |
| Current commit | pending |
| Fresh/stale result | pending |

Anvien command/read log:

| Step | Command or file | Purpose | Output artifact | Notes |
|---|---|---|---|---|
| pending | pending | pending | pending | pending |

Anvien candidates:

| Candidate id | Path | Symbol/name | Kind | Discovery evidence | Notes |
|---|---|---|---|---|---|
| pending | pending | pending | pending | pending | pending |

Anvien stop condition:

- pending

## E4 - Anvien-Guided Discovery Report

Date:

Status: pending

Required evidence:

| Item | Result |
|---|---|
| Anvien candidate count | pending |
| Anvien unique files read | pending |
| Anvien command count | pending |
| Anvien follow-up native search count | pending |
| Anvien estimated token total | pending |
| Anvien unresolved questions | pending |

## E5 - Candidate Union And Verification

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

## E6 - Final Comparison Evidence

Date:

Status: pending

Required comparison facts:

| Question | Evidence |
|---|---|
| Which method used fewer estimated tokens? | pending |
| Which method read fewer files? | pending |
| Which method found more confirmed/likely deadcode? | pending |
| Which method produced fewer false positives? | pending |
| Which token bucket dominated native-search cost? | pending |
| Which token bucket dominated Anvien-guided cost? | pending |
| Was graph tool output cost offset by fewer file reads/search outputs? | pending |
| Did validation prove or disprove the discovered candidates? | pending |
| Were redirected Anvien artifacts excluded until read by the agent? | pending |

Required summary shape:

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

## E7 - Closure

Date:

Status: pending

Closure checks:

| Check | Result |
|---|---|
| No deadcode deletion/edit was made | pending |
| Plan checklist updated | pending |
| Benchmark ledger complete | pending |
| Final comparison written | pending |
| Commit hash for documentation update, if committed | pending |
