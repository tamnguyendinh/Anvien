---
name: planner
description: "Use when the user asks to create, write, or review a docs/plans plan."
---

# Planner With Anvien

Use this skill when work involves creating, reviewing, or updating a plan under `docs/plans`.

This skill is a workflow gate for plan/evidence/benchmark/actual-status authoring. It is not a command router. When a concrete Anvien command is needed for implementation evidence, choose it directly from the generated Command Selection Guide.

## Standard Plan Set

A standard plan is a four-file set with the same date and slug:

```text
docs/plans/YYYY-MM-DD-<slug>/
  YYYY-MM-DD-<slug>-plan.md
  YYYY-MM-DD-<slug>-evidence.md
  YYYY-MM-DD-<slug>-benchmark.md
  YYYY-MM-DD-<slug>-actual-status.md
```

Rules:

- Put all four standard files in `docs/plans/YYYY-MM-DD-<slug>/`.
- Use ISO date format: `YYYY-MM-DD`.
- Use lowercase ASCII kebab-case for the slug.
- Keep the same slug in all four standard files; only the suffix changes.
- Use the matching H1: `Plan`, `Evidence Ledger`, `Benchmark Ledger`, or `Actual Status`.
- Auxiliary files such as `*-remaining-files.md` can exist, but they are not part of the standard four-file set.

## Template Files

Use the bundled templates before writing a new standard plan set:

- `templates/plan.template.md`
- `templates/evidence.template.md`
- `templates/benchmark.template.md`
- `templates/actual-status.template.md`

Rules:

- Start from these templates instead of inferring structure from older nearby plans.
- Read nearby plans only for local naming, scope, or evidence conventions that are clearly still active.
- Replace template placeholders with the concrete date, slug, title, companion file paths, and task scope.
- Keep template sections that protect discipline unless a user explicitly asks for a smaller review-only artifact.

## Plan File

The plan file controls the work.

It should contain:
- metadata: title, date, status, and companion files;
- goal;
- rules, master rules, or rules of plan;
- problem;
- scope, scope boundary, and non-goals;
- requirements, invariants, and technical direction when needed;
- acceptance criteria or definition of done;
- phase checklist with IDs such as `P0-A`, `P1-B`, and so on;
- risk notes;
- task state through checkboxes.

For implementation plans, the first phase must be P0 actual status:

- create and fill `YYYY-MM-DD-<slug>-actual-status.md`;
- classify current reality before implementation work;
- update later phase status assumptions, next actions, and work steps from that evidence before implementation starts.
- After completing a phase or implementation slice and refreshing `actual- status.md`, update the next affected phase's work steps as needed to match the latest repo reality, while preserving that phase's original goal, scope,  acceptance criteria, and major phase order.

Every checklist item must be a complete mini-plan by itself. Do not write generic checklist items. The plan must state what to do, in what order, and what condition proves the item is done, don't hide the logic in the outer section.**

Each checklist item must include:

- Goal: what the phase achieves.
- Work Steps: concrete ordered work, including the implementation sequence.
- Implementation Gate: the condition that must be true before editing or moving forward.
- Acceptance: the condition that proves the phase is done.

Do not use the plan file as a command log, benchmark ledger, changelog, or place to store long metric tables.

## Actual Status File

The actual-status file records the true current state before implementation work. It prevents planning from assumptions.

Do not treat P0 actual status as frozen after implementation begins. Keep it as a living current-state record:

- P0 records the baseline before implementation.
- After each completed implementation slice, update the affected status matrix rows to match the new repo reality.
- Before starting the next phase, check whether completed work or repo changes made previous current-state classifications stale.
- If actual status changes, update the next phase's status assumptions, next action, and work steps before implementation continues.
- Keep detailed proof in `evidence.md`; actual-status stores classification, evidence IDs, touch mode, and plan consequences.

The file has two layers:

- `Current Status Matrix`: the latest known state that the next agent should trust.
- `Status Refresh Log`: a short history of refreshes, status changes, evidence IDs, and next-phase updates.

It must answer:

- What is the real current state?
- Which surfaces, files, or slots are already correct?
- What is partial?
- What is fake, demo-only, stubbed, or placeholder?
- What is missing, unbound, or not wired?
- What is blocked?
- For file targets, how many related files does `anvien file-detail <path> --repo <repo> --json` report?
- Which related files/surfaces are preserve-only, inspect-only, editable, generated, validation-only, blocked, or out of scope?
- From that evidence, what status, next-action, or work-step updates are needed for the next phase?

Use clear status labels:

- `correct` or `bound-correct`: already works and should be preserved.
- `partial`: exists but does not fully satisfy the requirement.
- `wrong`: exists but conflicts with the requirement.
- `missing` or `unbound`: required surface does not exist or is not wired.
- `fake-or-stub`: demo-only, placeholder, or fake behavior.
- `blocked`: cannot proceed without external input or unavailable dependency.

If actual status finds `correct`/`bound-correct`, `partial`, `missing`/`unbound`, `fake-or-stub`, or `blocked`, update the affected next-phase status assumptions, next actions, and work steps before implementation work. Do not execute a stale phase after actual-status evidence changes the work.

After implementation begins, use status transitions such as `missing -> correct`, `partial -> correct`, `fake-or-stub -> removed`, or `unbound -> bound-correct` when evidence proves the new state. Append a refresh-log row instead of deleting the history.

## Evidence File

The evidence file explains why the work is known to be correct.

It should contain:

- metadata and companion files;
- evidence rules or evidence template;
- evidence sections such as `E0`, `E1`, or sections by phase/task;
- user report or problem evidence;
- source inspection, codebase facts, and document facts;
- commands run and pass/fail result;
- impact or blast-radius evidence when code/graph behavior changes;
- implementation evidence: files changed and behavior changed;
- validation evidence: build, tests, e2e, screenshots, or traces;
- failures encountered and how they were handled;
- detect-changes before commit;
- commit hash and closure evidence.

Evidence can reference short metric traces, but long metric tables belong in the benchmark file.

Evidence IDs must be exact and stable so `plan.md`, `actual-status.md`, `benchmark.md`, and later agents can reference proof without ambiguity.

Use this format:

```text
E<phase>-<item>-<kind><n>
```

Rules:

- `E<phase>` matches the plan phase number: `E0` for `P0`, `E1` for `P1`, `E2` for `P2`, and so on.
- `<item>` matches the checklist item without the dash: `P0A`, `P1A`, `P2B`.
- `<kind>` is plan-local. Choose a short uppercase token that is meaningful for the current repo and plan.
- `<n>` is a 1-based sequence number within that phase item and kind.
- Keep the same `<kind>` meaning stable inside one plan.
- Do not reuse an evidence ID for different facts.
- Reference exact evidence IDs from `actual-status.md` and `benchmark.md`; avoid referencing only broad section IDs such as `E1`.
- Use ranges such as `E0-P0A-FD1..E0-P0A-FD17` only for compact inventory summaries; use exact IDs when a specific status decision depends on a specific fact.
- If nearby plans already use a clear local evidence naming style, follow that style instead of inventing a new one.

Examples only:

- `E0-P0A-SRC1`
- `E0-P0A-GRAPH1`
- `E1-P1A-ROUTE1`
- `E2-P2B-KEYBOARD1`
- `E2-P2B-DETECT1`

## Benchmark File

The benchmark file records measurements.

It should contain:

- metadata and companion files;
- benchmark rules;
- benchmark sections such as `B0`, `B1`, or sections by phase/task;
- metric tables with unit, baseline, latest, final, target, and delta when needed;
- inventory count;
- runtime or performance metric;
- graph, coverage, or accuracy metric;
- package, bundle, file size, or hash metric;
- before/after numbers;
- UI, layout, or browser metric when the plan involves UI;
- command-surface or generated-output inventory when the plan involves generated documentation.

Benchmark records measured numbers only. Do not put command logs, design decisions, or validation narrative here. Build/test/e2e pass-fail belongs in evidence unless the timing, count, or size is the measured target.

## Workflow

1. Read the bundled templates for the four standard files.
2. Read nearby plan files only for active local conventions, not as the source of the standard format.
3. Confirm the date, slug, directory layout, and companion file links.
4. Create all four standard files before implementation work.
5. Complete `actual-status.md` as P0 before implementation work and update later phase status assumptions, next actions, and work steps from its evidence.
6. Keep phase checklists specific enough that another agent can implement them in order.
7. Record implementation evidence in the evidence file as the work completes.
8. Record benchmarkable inventory or performance counts in the benchmark file as the measurements are taken.
9. After each completed implementation slice, refresh affected rows in `actual-status.md` and append a `Status Refresh Log` row.
10. Before starting the next phase, compare the latest actual-status classifications with the next plan item; update only its stale status assumptions, next action, or work steps if repo reality changed.
11. Update the matching checklist item as soon as a phase is completed.
12. Keep generated output evidence separate from source-of-truth source changes.
