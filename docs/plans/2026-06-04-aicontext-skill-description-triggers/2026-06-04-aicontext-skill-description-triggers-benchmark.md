# Benchmark Ledger

Title: AI Context Skill Description Trigger Shortening

Date: 2026-06-04

Status: Draft - awaiting implementation

Companion files:

- Plan: [2026-06-04-aicontext-skill-description-triggers-plan.md](2026-06-04-aicontext-skill-description-triggers-plan.md)
- Evidence ledger: [2026-06-04-aicontext-skill-description-triggers-evidence.md](2026-06-04-aicontext-skill-description-triggers-evidence.md)

## Benchmark Rules

1. Record measured counts, sizes, and before/after inventory only.
2. Do not store command logs here; command context belongs in the evidence ledger.
3. Build/test pass-fail belongs in the evidence ledger unless timing is the measured target.
4. Generated-output inventory is benchmarkable because this plan changes always-loaded AI context size.

## B0 - Baseline AI Context Size

Status: recorded

Source evidence: E0, E3.

Measured on current `AGENTS.md` in the catalog implementation repository before full description shortening:

| Metric | Unit | Baseline |
|---|---:|---:|
| `AGENTS.md` total lines | lines | 175 |
| `AGENTS.md` total words | words | 3,364 |
| `AGENTS.md` total characters | chars | 26,434 |
| Rough token estimate at 4 chars/token | tokens | 6,608 |
| Rough token estimate at 3.5 chars/token | tokens | 7,553 |
| Rough token estimate at 3 chars/token | tokens | 8,811 |

## B1 - Skill Selection Guide Inventory

Status: recorded

| Metric | Unit | Baseline | Target |
|---|---:|---:|---:|
| Generated skill package rows | rows | 35 | 34 |
| Skills removed by this plan | skills | 0 | 1 |
| Targeted package descriptions | descriptions | 35 | 34 |
| Generated `Use` paths per single-entry package | paths | 1 | 1 |
| Generated `Use` paths for multi-entry packages | paths | measured during implementation | unchanged |

Removal target:

| Skill | Baseline | Target |
|---|---:|---:|
| `ai-multimodal` package | present | absent |

## B2 - Long Description Baseline Examples

Status: recorded

| Skill | Baseline description chars | Target description |
|---|---:|---|
| `aesthetic` | 769 | `Use when the user asks to improve UI aesthetics.` |
| `media-processing` | 659 | `Use when the user asks to process media files.` |
| `shopify` | 651 | `Use when the user asks to build Shopify apps, themes, or extensions.` |
| `backend-development` | 644 | `Use when the user asks to build or change backend code.` |
| `better-auth` | 607 | `Use when the user asks to implement Better Auth.` |
| `repomix` | 604 | `Use when the user asks to package a repository with Repomix.` |
| `anvien-debugging` | 222 | `Use when the user asks to debug.` |

## B3 - Final Measurements

Status: pending

To be recorded after implementation:

| Metric | Unit | Final | Delta |
|---|---:|---:|---:|
| `AGENTS.md` total characters | chars | pending | pending |
| `AGENTS.md` rough token estimate at 3 chars/token | tokens | pending | pending |
| Sum of generated skill descriptions | chars | pending | pending |
| Longest generated skill description | chars | pending | pending |
| Generated skill package rows | rows | pending | pending |
| Generated stale package/root path matches in `Use` column | matches | pending | pending |
| `ai-multimodal` generated references | matches | pending | pending |
