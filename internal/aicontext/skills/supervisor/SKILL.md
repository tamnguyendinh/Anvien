---
name: supervisor
description: Use whenever reviewing completion claims, fixes, diffs, reports, screenshots, or artifacts for acceptance; verify repo/project reality with Anvien evidence. Always use this skill before accepting any agent output, closing any task, or merging any result.
---

# Supervisor Review

Use this skill to independently decide whether a claim, artifact, work result, or completion statement can be accepted in a repo/project.

The review target is only the entry point. Verify the real claim, write the report, then summarize the verdict.

## Core Law

Zero-trust review: treat every claim, artifact, result, and completion statement as untrusted until independently verified against repo/project reality with sufficient evidence.

## Role Boundary

Supervisor Review gives an acceptance verdict. It does not repair the work while reviewing unless the user explicitly asks for a separate implementation task.

Do not blend review and fix work. If review finds a problem, write a REJECT report with evidence and the required next step.

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

## Core Workflow

1. **Reconstruct & Authorize** → Claim, authority, feedback (→ read `references/claim-and-authority.md`)
2. **Inspect & Verify** → Source inspection, evidence, Anvien (→ read `references/evidence-and-source-inspection.md`)
3. **Close Invariants** → Sibling sweep, history & resubmission (→ read `references/invariant-and-resubmission.md`)
4. **Report & Decide** → PASS/REJECT report & handoff (→ read `references/reporting.md`)

## Quick Decision Tree

- Unclear claim, conflicting authority, or reviewing feedback? → `references/claim-and-authority.md`
- Inspecting source diff, gathering evidence, or post-compaction re-anchor? → `references/evidence-and-source-inspection.md`
- Checking invariant closure, sweeping sibling surfaces, or resubmission after fix? → `references/invariant-and-resubmission.md`
- Writing the supervisor report or formatting final verdict? → `references/reporting.md`

## Reference Index

| Need | File |
|------|------|
| Claim conversion, authority hierarchy, feedback discipline | `references/claim-and-authority.md` |
| Evidence standards, source inspection gate, re-anchor | `references/evidence-and-source-inspection.md` |
| Invariant closure, sibling surface sweep, resubmissions | `references/invariant-and-resubmission.md` |
| Report template, naming conventions, verdict format | `references/reporting.md` |
