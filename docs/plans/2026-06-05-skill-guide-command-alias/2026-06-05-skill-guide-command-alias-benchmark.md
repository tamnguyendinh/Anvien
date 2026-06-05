# Benchmark Ledger

Title: Skill Guide Command Alias Column

Date: 2026-06-05

Status: Planned

Companion files:

- Plan: [2026-06-05-skill-guide-command-alias-plan.md](2026-06-05-skill-guide-command-alias-plan.md)
- Evidence ledger: [2026-06-05-skill-guide-command-alias-evidence.md](2026-06-05-skill-guide-command-alias-evidence.md)

## Benchmark Rules

1. Record measured inventory and generated-output counts only.
2. Build/test pass-fail belongs in the evidence ledger unless timing or count is the measured target.
3. Update this file as each benchmarkable implementation slice completes.
4. Generated files are measured output, not source of truth.

## B0 - Baseline Source Skill Inventory

Measured on 2026-06-05 before implementation, against the current working tree.

| Metric | Baseline | Unit | Source |
|---|---:|---|---|
| Top-level source skill directories | 37 | directories | `internal/aicontext/skills` |
| Recursive `SKILL.md` files | 58 | files | `internal/aicontext/skills/**/SKILL.md` |
| Top-level directories with direct `SKILL.md` | 35 | directories | `internal/aicontext/skills/<dir>/SKILL.md` |

## B1 - Baseline Generated Guide Inventory

Measured on 2026-06-05 before implementation.

| File | Skill guide rows | Command column | Command alias cells |
|---|---:|---:|---:|
| `AGENTS.md` | 37 | 0 | 0 |

Current generated header:

```text
| When you need to... | Use |
|---------------------|-----|
```

## B2 - Target Generated Guide Inventory

To measure after implementation and normal regeneration.

| File | Skill guide rows | Command column | Command alias cells | Duplicate aliases |
|---|---:|---:|---:|---:|
| `AGENTS.md` | TBD | 1 | TBD | 0 |
| `CLAUDE.md` | TBD | 1 | TBD | 0 |

Expected invariant:

```text
command alias cells == skill guide rows
```

## B3 - Target Alias Examples

To verify after implementation.

| Skill entry name | Expected command |
|---|---|
| `architect-review` | `/architect-review` |
| `System-architect` | `/system-architect` |
| `UI_taste skill` | `/ui-taste-skill` |
| `When Stuck - Problem-Solving Dispatch` | `/when-stuck-problem-solving-dispatch` |

## B4 - Generated Surface Path Inventory

To measure after implementation and normal regeneration.

| Surface | Expected command behavior | Expected path behavior |
|---|---|---|
| `AGENTS.md` | same aliases as `CLAUDE.md` | `.agents/skills/...` |
| `CLAUDE.md` | same aliases as `AGENTS.md` | `.claude/skills/...` |

## B5 - Nested Skill Primary-Entry Inventory

To measure after implementation.

| Package | Expected main guide command | Child commands in main guide |
|---|---|---:|
| `problem-solving` | `/problem-solving` | 0 |

## B6 - Final Inventory Template

Fill after implementation:

| Metric | Baseline | Final | Delta |
|---|---:|---:|---:|
| `AGENTS.md` skill guide rows | 37 | TBD | TBD |
| `AGENTS.md` command alias cells | 0 | TBD | TBD |
| `CLAUDE.md` skill guide rows | TBD | TBD | TBD |
| `CLAUDE.md` command alias cells | 0 | TBD | TBD |
| duplicate generated command aliases | TBD | 0 | TBD |
