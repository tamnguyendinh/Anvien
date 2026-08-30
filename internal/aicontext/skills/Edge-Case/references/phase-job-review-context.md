# Phase & Job Review Context

> This file is part of Edge-Case Review Skill v2. Read when: evaluating edge cases for an active phase/job in `Docs/execution/` before completion.

---

## Context Overview

Use this context guide when the active review target is a declared phase or job in `Docs/execution/*` that still lacks full verification or supervisor clearance.

**Core Mission:** Break the active job's failure paths before it can be marked ready for review or merged.

---

## Preflight Checklist

1. Reload `SKILL.md` plus `AGENTS.md`.
2. Anchor strictly to the declared phase/job and read its exact `Docs/SPEC/*` family.
3. Check `git status --short`.
4. Inspect stale or generated outputs that can mask edge failures:
   - `dist/`
   - `playwright-report/`
   - `test-results/`
   - `.tmp/`
5. Construct a fresh perturbation matrix for the declared job scope before running attacks.

---

## Execution Workflow

1. **Identify Highest-Risk Invariants**: What is the most fragile failure boundary in this job? (locks, money, sync, permissions, state).
2. **Select Domain Guides**: Read the corresponding attack guides in `references/` (e.g. `lock-and-concurrent-write.md`, `crash-recovery-and-partial-apply.md`).
3. **Choose Attack Vehicle**:
   - Pick the runtime attack vehicle that most directly forces the breakage to surface.
   - Use Playwright only when browser/operator sequencing is truly the necessary vehicle.
   - Otherwise use direct API calls, CLI commands, network interruptions, or service restarts.
   - Never rely on passive code inspection alone to close a runnable breakage path.
4. **Execute Hostile Perturbations**: Run bad timing, bad ordering, bad state, and bad inputs against the live path.
5. **Evaluate & Report**: Record confirmed failures or risky gaps following `references/reporting-and-evidence.md`.
