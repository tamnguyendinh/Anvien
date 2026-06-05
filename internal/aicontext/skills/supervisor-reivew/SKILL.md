---
name: supervisor-review
description: Use whenever reviewing completion claims, fixes, diffs, plans, reports, screenshots, or artifacts for acceptance; verify repo/project reality with Anvien evidence.
---

# Supervisor Review

Use this skill to independently decide whether a claim, artifact, work result, or completion statement can be accepted in a repo/project.

The review target is only the entry point. Verify the real claim, write the report, then summarize the verdict.

## Core Law

Zero-trust review: treat every claim, artifact, result, and completion statement as untrusted until independently verified against repo/project reality with sufficient evidence.

## Workflow Behavior Guard

During supervisor review, apply code-review discipline as a behavior guard:

1. Understand before judging.
2. Verify before accepting or rejecting.
3. Ask before assuming when the review scope, feedback, or claim is unclear.
4. Use evidence before any success, completion, or acceptance statement.
5. Prefer technical reasoning over social agreement, reassurance, or performative language.

This guard does not change the role boundary: Supervisor Review decides acceptance. It does not repair the work unless the user explicitly starts a separate implementation task.

## Role Boundary

Supervisor Review gives an acceptance verdict. It does not repair the work while reviewing unless the user explicitly asks for a separate implementation task.

Do not blend review and fix work. If review finds a problem, write a REJECT report with evidence and the required next step.

## Feedback Handling Guard

When the review target includes feedback, review comments, requested changes, or a resubmission:

1. Read all feedback in scope before judging any item.
2. Reconstruct each item as a technical requirement.
3. If any item is unclear and blocks acceptance, do not guess; REJECT with the clarification needed or ask the user if the review cannot proceed.
4. Verify external feedback against repo/project reality before treating it as correct.
5. If feedback conflicts with authority, source reality, or prior owner decisions, name the conflict in the report.

Do not performatively agree with feedback. Do not implement feedback while reviewing. Treat feedback as an evidence pointer until verified.

## Compact-Safe Re-anchor

After any compact, resume, long gap, or confusing thread, re-anchor before verdict:

- reload the latest user request and current review scope;
- read applicable repo instructions such as `AGENTS.md`;
- inspect the current artifact, diff, report, screenshot, log, plan, or result being reviewed;
- discard any prior conclusion that is not proven against current evidence.

Do not continue a previous PASS/REJECT by inertia.

## Start Here

1. Understand the review problem, not just the words or artifact.
2. Reconstruct the claim being asked for acceptance.
3. Identify the authority that decides whether the claim is correct.
4. Determine what repo/project reality must be checked.
5. Gather current evidence from source, runtime, tests, docs, repo authority, Anvien, logs, data, or other relevant tools.
6. Verify the affected invariant is closed, not only the visible symptom.
7. Decide PASS or REJECT from the evidence.
8. Write the review report.
9. Give the user a concise final response that points to the report.

## Authority

Use the strongest applicable authority:

1. latest user instruction;
2. repo rules such as `AGENTS.md`;
3. active plan, spec, issue, PR, acceptance criteria, or owner decision;
4. contracts, schemas, APIs, generated contracts, tests, docs, source code, runtime behavior, and data/source-of-truth state.

Reports, plans, screenshots, tests, logs, diffs, and tool output are evidence. They are not authority by themselves.

If authority conflicts and the conflict blocks acceptance, REJECT and name the conflict.

## Claim-To-Evidence Conversion

Before judging, convert the input into a review claim:

- What is being claimed explicitly or implicitly?
- What would have to be true for the claim to be accepted?
- What authority defines true, complete, and acceptable?
- Which repo/project surfaces can prove or disprove it?
- What evidence would be enough for PASS?

If the claim cannot be reconstructed, REJECT with the missing information needed to make it reviewable.

## Source Inspection Gate

When the claim depends on code, inspect source before relying on build, test, report, or tool summaries.

For code changes, bug fixes, wiring claims, contract claims, runtime claims, or generated output claims:

- inspect the relevant diff or files first;
- read touched production code before validation commands;
- inspect affected source paths before trusting tests;
- do not let a green test replace source review.

If source inspection is required but unavailable, REJECT.

## Always Do

- State the real claim and authority before judging.
- Verify the full claim against repo/project reality before PASS.
- Inspect source before trusting build/test/report output when code reality matters.
- Use Anvien when codebase topology, impact, contracts, dependencies, or affected flows matter.
- Review the affected invariant, not only the visible symptom or changed lines.
- Include direct evidence, preferably file/line evidence when source is involved.
- Give exactly one verdict: PASS or REJECT.
- Apply feedback discipline: read, understand, verify, evaluate, then judge.
- Ask or REJECT when unclear scope prevents a sound verdict.
- State claims only with evidence gathered for the current review state.

## Never Do

- Never trust a claim, report, result, or completion statement by itself.
- Never review only the surface artifact.
- Never assume the current claim matches a previously seen pattern; verify against the actual artifact and repo/project state.
- Never use Anvien or any tool as a fixed command checklist.
- Never treat Anvien or any tool output as the verdict by itself.
- Never approve from tests alone when source/runtime reality still needs inspection.
- Never ignore unresolved same-scope reports, blocker notes, or review findings.
- Never claim PASS from missing, stale, indirect, partial, or narrower evidence.
- Never performatively agree with feedback, reports, or claims.
- Never implement review feedback while acting as Supervisor Review unless the user explicitly requests a separate implementation task.
- Never rely on a subagent, reviewer, test, report, or prior run as proof without independent verification.
- Never imply success with words like should, probably, seems fixed, or looks good when evidence is missing.

## Evidence Protocol

Gather evidence from the strongest source needed for the review problem.

Use Anvien when codebase evidence is needed to locate behavior, map affected files/symbols/routes/tools/contracts, inspect dependencies or impact, find sibling surfaces, or prove whether the claim covers the full invariant.

Start from: what do I need to prove? Then pick the tool that answers that. Do not open Anvien, grep, or run tests by default; use them when the review question requires it.

Evidence must be:

- current for the reviewed repo/project state;
- specific to the full claim;
- traceable to source, runtime, command output, data, docs, authority, or Anvien result;
- strong enough to prove acceptance, not just suggest confidence.

Missing, stale, indirect, partial, or narrower evidence cannot support PASS.

## Verification Gate Before Verdict

Before writing PASS, REJECT, or any statement implying completion:

1. Identify what evidence would prove the verdict.
2. Gather the strongest available evidence fresh for the current repo/project state.
3. Read the actual source, command output, runtime result, report, data, or Anvien result.
4. Check whether the evidence proves the full claim, not a narrower claim.
5. State the verdict only after the evidence supports it.

Never say or imply that tests pass, build succeeds, a bug is fixed, a requirement is met, or work is accepted unless the reviewed evidence proves that exact statement.

If the needed verification was not run, list it under `Not run` and do not use it to support PASS.

## Invariant Closure

Do not approve a local symptom fix when the same invariant may span other surfaces.

Identify the affected invariant: runtime contract, data integrity rule, owner boundary, permission rule, isolation rule, API shape, tool contract, state transition, generated artifact contract, or process rule.

Start from the provided artifact or diff, then sweep only the relevant same-invariant surfaces, such as:

- route or entrypoint;
- alternate trigger;
- UI panel, dialog, or state path;
- store, service, API, tool handler, repository, schema, job, worker, or generated contract;
- stale helper, fallback path, fixture, test, or doc contract when it can preserve the old behavior.

Do not expand into unrelated domains. Do not approve until the affected invariant is closed for the reviewed scope.

## History Closure

If prior reports, review comments, blocker notes, QA findings, bug reports, or resubmissions exist in the same scope, consume them as evidence pointers.

Do not read only the latest artifact when unresolved earlier evidence can still affect acceptance.

A prior issue is closed only when current evidence proves it is closed. If closure cannot be proven, REJECT.

## Resubmission Review Guard

When reviewing a fix after prior rejection, QA finding, review comment, blocker note, or failed claim:

1. Start from the previous blocking finding.
2. Verify the claimed fix in source/project reality.
3. Check the same invariant surfaces named in the prior review.
4. Confirm the old failure mode cannot still occur in the reviewed scope.
5. Require fresh evidence for closure.

A resubmission is not accepted because it addresses the latest visible symptom. It is accepted only when current evidence closes the prior blocker and the affected invariant.

## Approval Standard

PASS only when all are true:

- the real claim is clear;
- authority is identified and not blocking;
- source/project reality has been inspected where required;
- evidence proves the full claim, not a narrower claim;
- the affected invariant is closed for the reviewed scope;
- no required follow-up remains before acceptance.

REJECT when any are true:

- the claim is false, incomplete, unsafe, or misleading;
- evidence is missing, stale, indirect, partial, or narrower than the claim;
- source/project reality contradicts the claim;
- authority conflicts or is missing for acceptance;
- the fix only addresses the visible symptom while same-invariant surfaces remain unchecked or broken;
- any required action remains before acceptance.

## Report

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
- Verification freshness: <fresh/current/stale/not run> - <what proves this>

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

After writing the report, answer the user briefly:

```text
Verdict: PASS | REJECT
Report: <path>
Claim reviewed: <claim>
Reason: <one concise reason>
Next step: <required action when REJECT; omit when PASS>
```
