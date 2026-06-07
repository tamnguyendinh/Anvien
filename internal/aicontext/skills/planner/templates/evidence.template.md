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

### Evidence ID Naming

Use stable, phase-scoped evidence IDs so `plan.md`, `actual-status.md`, `benchmark.md`, and later agents can reference exact proof without ambiguity.

Format:

```text
E<phase>-<item>-<kind><n>
```

Rules:

- `E<phase>` matches the plan phase number: `E0` for `P0`, `E1` for `P1`, `E2` for `P2`, and so on.
- `<item>` matches the checklist item without the dash: `P0A`, `P1A`, `P2B`.
- `<kind>` is plan-local. Choose a short uppercase token that is meaningful for this repo and this plan.
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

Evidence sections must follow the plan phases:

- `E0` corresponds to `P0`.
- `E1` corresponds to `P1`.
- `E2` corresponds to `P2`.
- Use exact evidence IDs inside each section, not broad section IDs as proof.
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
