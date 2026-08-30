# Worktree & Incident Review Context

> This file is part of Edge-Case Review Skill v2. Read when: evaluating current worktree, hunting bugs, reviewing resubmissions, or analyzing post-completion incidents.

---

## Context Overview

Use this context guide when all backlog jobs are complete or when explicitly assigned to review the `current worktree`, perform a bug hunt, follow up on an incident, or stress a resubmission.

**Core Mission:** Independently attack live runtime paths to uncover hidden edge-case regressions or unhandled edge boundaries in the working tree.

---

## Anchor Rules For Post-Completion Work

- When the phase/job backlog is exhausted, old phase/job documents become **historical context only**, not the primary review anchor.
- You must anchor directly to:
  1. The exact `Docs/SPEC/*` family governing the assigned scope.
  2. The mounted runtime path of the current worktree.
- Old phase/job order must not be pulled back in as review context.

---

## Preflight Checklist

1. Reload `SKILL.md` plus `AGENTS.md`.
2. Read the relevant SPEC family directly from the assigned post-completion scope.
3. Check `git status --short`.
4. Inspect and clear stale or generated build/test outputs:
   - `dist/`
   - `playwright-report/`
   - `test-results/`
   - `.tmp/`
5. Build a fresh, scope-specific perturbation matrix.

---

## Execution Workflow

1. **Map Scope to SPEC Authority**: Resolve the exact invariants governing this incident or worktree slice.
2. **Select Domain Guides**: Read the corresponding attack guides in `references/` based on the technical risk.
3. **Execute Active Attack**:
   - Drive the live flow using the most direct attack vehicle (API, CLI, process interrupt, or Playwright when operator flow is required).
   - Test compounded scenarios (e.g. reconnect during retry while a stale store is mounted).
4. **Evaluate Fail-Safe Response**: Verify that the system fails closed and preserves all invariants.
5. **Report Findings**: Record results in a new timestamped report per `references/reporting-and-evidence.md`.
