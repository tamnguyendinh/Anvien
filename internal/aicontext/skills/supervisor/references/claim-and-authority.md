# Claim And Authority

> This file is part of Supervisor Skill v2. Read when: clarifying authority, reconstructing the review claim, or handling external review comments/feedback.

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

If the claim cannot be reconstructed, REJECT (unactionable - claim unclear, cannot proceed) with the missing information needed to make it reviewable.

## Feedback Handling Guard

When the review target includes feedback, review comments, requested changes, or a resubmission:

1. Read all feedback in scope before judging any item.
2. Reconstruct each item as a technical requirement.
3. If any item is unclear and blocks acceptance, do not guess; REJECT with the clarification needed or ask the user if the review cannot proceed.
4. Verify external feedback against repo/project reality before treating it as correct.
5. If feedback conflicts with authority, source reality, or prior owner decisions, name the conflict in the report.

Do not performatively agree with feedback. Do not implement feedback while reviewing. Treat feedback as an evidence pointer until verified.

## Workflow Behavior Guard

During supervisor review, apply code-review discipline as a behavior guard:

1. Understand before judging.
2. Verify before accepting or rejecting.
3. Ask before assuming when the review scope, feedback, or claim is unclear.
4. Use evidence before any success, completion, or acceptance statement.
5. Prefer technical reasoning over social agreement, reassurance, or performative language.

This guard does not change the role boundary: Supervisor Review decides acceptance. It does not repair the work unless the user explicitly starts a separate implementation task.
