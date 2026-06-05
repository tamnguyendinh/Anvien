# Benchmark Ledger

Title: Codex And Claude Skill Layout Split

Date: 2026-06-05

Status: Complete

Companion files:

- Plan: [2026-06-05-codex-claude-skill-layout-split-plan.md](2026-06-05-codex-claude-skill-layout-split-plan.md)
- Evidence ledger: [2026-06-05-codex-claude-skill-layout-split-evidence.md](2026-06-05-codex-claude-skill-layout-split-evidence.md)

## Benchmark Rules

1. Record measured inventory, token, size, and generated-output counts only.
2. Build/test pass-fail belongs in the evidence ledger unless timing or count is the measured target.
3. Update this file as each benchmarkable implementation slice completes.

## B0 - Baseline Skill Catalog Inventory

Measured on 2026-06-05 before implementation.

| Metric | Baseline | Unit | Source |
|---|---:|---|---|
| Source skill packages | 34 | packages | `internal/aicontext/skills` directory count |
| Source `SKILL.md` entries | 43 | files | recursive `SKILL.md` count |
| Source skill payload files | 584 | files | recursive source file count |

## B1 - Baseline Generated Guide Inventory

Measured on 2026-06-05 before implementation.

| File | Total tokens (`o200k_base`) | Skill guide tokens (`o200k_base`) | Skill rows | `.claude/skills/anvien/` references | `.agents/skills/` references |
|---|---:|---:|---:|---:|---:|
| `AGENTS.md` | 3561 | 1055 | 34 | 34 | 0 |
| `CLAUDE.md` | 3559 | 1055 | 34 | 34 | 0 |

## B2 - Target Generated Guide Inventory

Measured after implementation on 2026-06-05 with `go run ./cmd/anvien analyze --force`.

| File | `.claude/skills/anvien/` references | `.agents/skills/` references | `.claude/skills/` direct references |
|---|---:|---:|---:|
| `AGENTS.md` | 0 | 34 | 0 |
| `CLAUDE.md` | 0 | 0 | 34 |

## B2A - Target Volatile Intro Inventory

Measured after implementation on 2026-06-05.

| File | `This project is indexed by Anvien` sentences | Volatile symbol/relationship/flow count sentence |
|---|---:|---:|
| `AGENTS.md` | 0 | 0 |
| `CLAUDE.md` | 0 | 0 |

Post-change token inventory:

| File | Total tokens (`o200k_base`) | Skill guide tokens (`o200k_base`) | Skill rows |
|---|---:|---:|---:|
| `AGENTS.md` | 3390 | 953 | 34 |
| `CLAUDE.md` | 3422 | 987 | 34 |

## B3 - Target Generated Output Inventory

Measured after implementation on 2026-06-05.

| Surface | Expected package roots | Expected `SKILL.md` entries | Expected source payload files |
|---|---:|---:|---:|
| `.agents/skills` | 34 | 43 | 584 |
| `.claude/skills` | 34 | 43 | 584 |

Expected generated metadata:

| Surface | Expected direct-root manifest files | Expected legacy namespace roots after migration |
|---|---:|---:|
| `.agents/skills` | 1 | 0 |
| `.claude/skills` | 1 | 0 |

Expected direct-root manifest paths:

| Surface | Manifest path |
|---|---|
| Codex repo skills | `.agents/skills/.agents-skill-manifest.json` |
| Claude repo skills | `.claude/skills/.claude-skill-manifest.json` |

Measured direct-root manifest package counts:

| Surface | Manifest package count |
|---|---:|
| `.agents/skills` | 34 |
| `.claude/skills` | 34 |

## B4 - Collision And Custom Skill Preservation Inventory

Measured through focused fixtures in `go test ./internal/aicontext` and `go test ./internal/cli`.

| Fixture | Expected custom roots before sync | Expected custom roots after sync | Expected collisions | Expected adopted |
|---|---:|---:|---:|---:|
| unrelated custom root under `.agents/skills` | 1 | 1 | 0 | 0 |
| unrelated custom root under `.claude/skills` | 1 | 1 | 0 | 0 |
| same-name foreign root under `.agents/skills` | 1 | 1 | 1 | 0 |
| same-name foreign root under `.claude/skills` | 1 | 1 | 1 | 0 |
| unrelated custom root under a shared setup target such as Cursor/OpenCode | 1 | 1 | 0 | 0 |
| same-name exact desired snapshot without manifest | 1 | 1 | 0 | 1 |
| same-name root with extra file without manifest | 1 | 1 | 1 | 0 |
