# Planner Four-Template Standard Actual Status

## Purpose

Record the real current state before implementation work so the plan is based on evidence instead of guesses.

## Status Rules

| Status | Meaning |
|--------|---------|
| `correct` | Already works and should be preserved. |
| `partial` | Some required behavior exists, but the new requirement is not fully satisfied. |
| `wrong` | Existing behavior conflicts with the new requirement. |
| `missing` | Required surface does not exist yet. |
| `fake-or-stub` | Surface exists but is demo-only, placeholder, or not actually wired. |
| `blocked` | Cannot proceed without external input or unavailable dependency. |

## Current Status Matrix

| Surface / File / Slot | Status | Evidence | Required rewrite |
|-----------------------|--------|----------|------------------|
| `internal/aicontext/skills/planner/SKILL.md` | `partial` | The skill defines a three-file `Standard Plan Set` and does not require `actual-status.md`. | Rewrite to four files and make actual status a P0 gate before implementation work. |
| `internal/aicontext/skills/planner/templates/*` | `missing` | Planner directory currently has no template folder or template assets. | Add four templates: plan, evidence, benchmark, actual-status. |
| `internal/aicontext/aicontext_test.go` planner protection | `correct` | Existing tests already cover skill package discovery and nested payload behavior. User explicitly rejected additional content-locking tests for planner wording/templates. | Preserve; run focused existing tests only. |
| `internal/aicontext/skill_packages.go` payload discovery | `correct` | Existing package discovery walks skill files recursively and package tests already cover nested planner assets. | Preserve; no implementation rewrite unless tests prove a gap. |
| Generated `.agents` / `.claude` skill mirrors | `partial` | Generated content will not include new templates until regeneration. | Regenerate from source after edits. |
| Root `AGENTS.md` / `CLAUDE.md` generated blocks | `correct` | Rules say these are generated outputs and must not be edited as permanent source. | Do not edit directly. |

## Fake / Demo / Unbound / Blocked

- Fake/demo: none found.
- Unbound/missing: planner templates and actual-status standard are missing.
- Blocked: none.

## Rewrite Decisions

- P1 must change planner source instructions and add template assets.
- P2 must validate with existing tests and must not add brittle content-locking assertions for planner text.
- P3 must verify generation copies the new templates to agent mirrors.
- Package discovery code should remain unchanged unless validation fails.

## P0 Decision

P0 is complete. Proceed to coding with a narrowed scope: planner skill markdown, planner template files, and focused tests.
