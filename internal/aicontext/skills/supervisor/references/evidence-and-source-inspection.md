# Evidence And Source Inspection

> This file is part of Supervisor Skill v2. Read when: inspecting source diffs, gathering codebase evidence with Anvien, executing verification gates, or re-anchoring post-compaction.

## Source Inspection Gate

When the claim depends on code, inspect source before relying on build, test, report, or tool summaries.

For code changes, bug fixes, wiring claims, contract claims, runtime claims, or generated output claims:

- inspect the relevant diff or files first;
- read touched production code before validation commands;
- inspect affected source paths before trusting tests;
- do not let a green test replace source review.

If source inspection is required but unavailable, REJECT.

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

## Compact-Safe Re-anchor

After any compact, resume, long gap, or confusing thread, re-anchor before verdict:

- reload the latest user request and current review scope;
- read applicable repo instructions such as `AGENTS.md`;
- inspect the current artifact, diff, report, screenshot, log, plan, or result being reviewed;
- discard any prior conclusion that is not proven against current evidence.

Do not continue a previous PASS/REJECT by inertia.
