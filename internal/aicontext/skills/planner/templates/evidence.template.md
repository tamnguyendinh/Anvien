# {{TITLE}} Evidence Ledger

## Metadata

- Date: `{{YYYY-MM-DD}}`
- Plan: `{{PLAN_PATH}}`
- Evidence: `{{EVIDENCE_PATH}}`
- Benchmark: `{{BENCHMARK_PATH}}`
- Actual status: `{{ACTUAL_STATUS_PATH}}`

## Evidence Rules

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

Evidence sections must follow the plan phases:

- `E0` corresponds to `P0`.
- `E1` corresponds to `P1`.
- `E2` corresponds to `P2`.
- Use item-level IDs such as `E-P1-A` when a checklist item needs separate evidence.
- Each evidence section must name the plan phase or checklist item it supports.
- Do not invent fixed evidence categories; record the evidence required by the matching plan phase.

## E0 - P0 Evidence

Matching plan item(s): `P0-A`

{{P0_EVIDENCE}}

## E1 - P1 Evidence

Matching plan item(s): `P1-A`

{{P1_EVIDENCE}}

## E2 - P2 Evidence

Matching plan item(s): `P2-A`

{{P2_EVIDENCE}}

## Closure Evidence

Use this section for final detect-changes, commit hash, and closure evidence when the plan reaches completion.

{{CLOSURE_EVIDENCE}}
