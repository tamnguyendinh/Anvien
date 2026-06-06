# {{TITLE}} Plan

## Metadata

- Date: `{{YYYY-MM-DD}}`
- Status: `draft`
- Plan: `{{PLAN_PATH}}`
- Evidence: `{{EVIDENCE_PATH}}`
- Benchmark: `{{BENCHMARK_PATH}}`
- Actual status: `{{ACTUAL_STATUS_PATH}}`

## Goal

{{GOAL}}

## Rules

- Complete P0 actual status before implementation work.
- Update each checklist item immediately when it is completed.
- Record evidence as work completes.
- Record benchmarkable counts or measurements when they are taken.
- Rewrite later phases when actual-status evidence changes the scope.

## Problem

{{PROBLEM}}

## Scope

{{SCOPE}}

## Non-Goals

{{NON_GOALS}}

## Requirements

{{REQUIREMENTS}}

## Acceptance Criteria

{{ACCEPTANCE_CRITERIA}}

## Checklist

- [ ] P0-A: Complete actual status before implementation work.
  - Goal: establish the real current state.
  - Work Steps: inspect source-of-truth files, classify each surface, record blocked or missing pieces, and rewrite later phases from evidence.
  - Implementation Gate: no implementation or editing starts until `{{ACTUAL_STATUS_PATH}}` has a final P0 decision.
  - Acceptance: actual status identifies correct, partial, missing/unbound, fake-or-stub, and blocked surfaces for this scope.
- [ ] P1-A: {{PHASE_1_TITLE}}
  - Goal: {{PHASE_1_GOAL}}
  - Work Steps: {{PHASE_1_WORK_STEPS}}
  - Implementation Gate: {{PHASE_1_GATE}}
  - Acceptance: {{PHASE_1_ACCEPTANCE}}

## Risk Notes

{{RISK_NOTES}}
