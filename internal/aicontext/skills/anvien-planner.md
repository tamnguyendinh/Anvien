---
name: anvien-planner
description: "Use when the user needs to create or review docs/plans plan, evidence, benchmark, or checklist mini-plan work."
---

# Planner With Anvien

Use this skill when work involves creating, reviewing, or updating a plan under `docs/plans`.

This skill is a workflow gate for plan/evidence/benchmark authoring. It is not a command router. When a concrete Anvien command is needed for implementation evidence, choose it directly from the generated Command Selection Guide.

## Standard Plan Set

A standard plan is a three-file set with the same date and slug:

```text
docs/plans/YYYY-MM-DD-<slug>/
  YYYY-MM-DD-<slug>-plan.md
  YYYY-MM-DD-<slug>-evidence.md
  YYYY-MM-DD-<slug>-benchmark.md
```

Rules:

- Put all three standard files in `docs/plans/YYYY-MM-DD-<slug>/`.
- Use ISO date format: `YYYY-MM-DD`.
- Use lowercase ASCII kebab-case for the slug.
- Keep the same slug in all three standard files; only the suffix changes.
- Use the matching H1: `Plan`, `Evidence Ledger`, or `Benchmark Ledger`.
- Auxiliary files such as `*-remaining-files.md` can exist, but they are not part of the standard three-file set.

## Plan File

The plan file controls the work.

It should contain:

- metadata: title, date, status, and companion files;
- rules, master rules, or rules of plan;
- goal;
- problem;
- scope, scope boundary, and non-goals;
- requirements, invariants, and technical direction when needed;
- acceptance criteria or definition of done;
- phase checklist with IDs such as `P0-A`, `P1-B`, and so on;
- risk notes;
- task state through checkboxes.

Every checklist item must be a complete mini-plan by itself. Do not write generic checklist items. The plan must state what to do, in what order, and what condition proves the item is done.

Each checklist item must include:

- Goal: what the phase achieves.
- Work Steps: concrete ordered work, including the implementation sequence.
- Implementation Gate: the condition that must be true before editing or moving forward.
- Acceptance: the condition that proves the phase is done.

Do not use the plan file as a command log, benchmark ledger, changelog, or place to store long metric tables.

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

1. Read the existing nearby plan files before creating or changing a plan.
2. Confirm the date, slug, directory layout, and companion file links.
3. Keep phase checklists specific enough that another agent can implement them in order.
4. Update the matching checklist item as soon as a phase is completed.
5. Record implementation evidence in the evidence file as the work completes.
6. Record benchmarkable inventory or performance counts in the benchmark file as the measurements are taken.
7. Keep generated output evidence separate from source-of-truth source changes.
