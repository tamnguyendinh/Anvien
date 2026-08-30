# Reporting

> This file is part of Supervisor Skill v2. Read when: formatting supervisor reports, structuring verdicts, or writing durable acceptance artifacts.

## Report Rules And Conventions

A supervisor review is not complete until a report is written. The report is the durable evidence artifact; the chat response is only a summary.

Use the repo's required report convention when one exists. Otherwise use:

- Review report: `rp_supervisor_<YYMMDD>_<HHMMSS>_by_<model_slug>_<scope>.md`

Filename rules:

- time uses the repo-required timezone, or local review time if none is specified;
- `model_slug` uses the model identifier, such as `gpt-5-codex`, `gpt-4o`, or `claude-sonnet-4-6`;
- `scope` must be stable, lowercase, ASCII, and descriptive;
- use `_` between filename fields, and `-` inside `model_slug` only when needed;
- do not use spaces or non-ASCII characters;
- do not rename legacy reports just to fit this convention.

Write the report in the repo's required report location, or the nearest appropriate review/report area if none exists. The report must let a future reader understand the reviewed claim, the problem, the evidence, and the path to acceptance without reading the chat.

## Report Template

```text
# Supervisor Report: <short readable title>

Verdict: PASS | REJECT

## Metadata
- Report file: <filename>
- Review time: <YYMMDD HHMMSS and timezone>
- Reviewer: <model_slug>
- Repo/project: <repo or project name>
- Scope reviewed: <plan/diff/fix/report/artifact/worktree/commit window>
- Claim reviewed: <claim being accepted or rejected>
- Authority used: <user request, repo rules, plan/spec, contract, runtime, source, etc.>
- Related artifacts: <reports, screenshots, logs, PRs, commits, or none>

## Executive Summary
- Problem: <what issue or acceptance question this review is about>
- Decision: <why the verdict is PASS or REJECT>
- Required outcome: <what must happen next when REJECT, or "accepted" when PASS>

## Blocking Findings
Use this section for REJECT. Omit it for PASS if there are no blocking findings.

### [SEVERITY] <finding title>
File: <path:line, or "N/A" if not source-backed>
Issue: <clear explanation of the defect, gap, unsafe claim, or missing proof>
Evidence: <source/tool/command/runtime/doc/data evidence and what each item proves>
Why this blocks acceptance: <tie the finding to authority, invariant, risk, or acceptance criteria>
Fix direction: <how to close the issue>
Re-review evidence required: <what evidence must be supplied for the next review>

## Source-Level Clearance Notes
For source-involved reviews, include at least one direct finding or explicit clearance note for each touched production file group.

- <file group or path>: <clear / blocked / not applicable> - <file:line evidence and reason>

## Evidence Checked
Passed:
- <command/source/runtime/doc/data/tool evidence that passed>
- Verification freshness: <fresh/current/stale/not run> - <what proves this>

Failed:
- <command/source/runtime/doc/data/tool evidence that failed>
- Verification freshness: <fresh/current/stale/not run> - <what proves this>

Not run:
- <evidence not gathered and why>

## Invariant Closure
- affected invariant: <runtime contract, data rule, API shape, process rule, etc.>
- sibling surfaces checked: <routes, handlers, stores, tools, docs, tests, generated contracts, etc.>
- residual unverified same-invariant surfaces: <none, or list with reason>

## Required Fix List For Resubmission
Use this section for REJECT.

1. <specific action required>
2. <specific evidence required>

## Overall Evaluation
<short assessment of why the work is acceptable or not, distinguishing implementation quality, evidence quality, authority conflict, and remaining risk>
```

For REJECT, the next step must explain how to close the affected invariant, not just the isolated symptom.

For PASS, state that residual same-invariant unverified surfaces are none, or why no sibling sweep was needed.

## Final Summary Response Format

After writing the report, answer the user briefly:

```text
Verdict: PASS | REJECT
Report: <path>
Claim reviewed: <claim>
Reason: <one concise reason>
Next step: <required action when REJECT; omit when PASS>
```
