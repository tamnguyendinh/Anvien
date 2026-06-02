# Anvien Skill Selection Guide Reduction Benchmark Ledger

Date: 2026-06-02

Status: Active

Companion files:

- Plan: [2026-06-02-anvien-skill-selection-guide-reduction-plan.md](2026-06-02-anvien-skill-selection-guide-reduction-plan.md)
- Evidence ledger: [2026-06-02-anvien-skill-selection-guide-reduction-evidence.md](2026-06-02-anvien-skill-selection-guide-reduction-evidence.md)

## Benchmark Rules

1. Record quantitative inventory data only.
2. Put interpretation and command context in the evidence ledger.
3. Track embedded source skill inventory and generated skill inventory before and after implementation.
4. Track retained and removed skill counts explicitly.
5. Build/test pass/fail belongs in the evidence ledger unless timing, size, or inventory is the measured target.

## B0 - Skill Inventory Baseline And Target

Status: recorded

Source evidence: E0.

| Metric | Unit | Baseline | Target | Delta |
|---|---:|---:|---:|---:|
| Embedded Anvien source skills | files | 10 | 3 | -7 |
| Generated Anvien skill directories | directories | 10 | 3 | -7 |
| Retained Anvien workflow skills | skills | 3 | 3 | 0 |
| Removed broad/router Anvien skills | skills | 7 | 0 | -7 |
| Generated `Skill Selection Guide` rows | rows | 0 | 3 | +3 |

Retained skill target:

| Skill | Unit | Target |
|---|---:|---:|
| `anvien-api-surface` | retained | 1 |
| `anvien-refactoring` | retained | 1 |
| `anvien-debugging` | retained | 1 |

Removed skill target:

| Skill | Unit | Target |
|---|---:|---:|
| `anvien-cli` | removed | 1 |
| `anvien-cross-repo` | removed | 1 |
| `anvien-exploring` | removed | 1 |
| `anvien-graph-quality` | removed | 1 |
| `anvien-guide` | removed | 1 |
| `anvien-impact-analysis` | removed | 1 |
| `anvien-runtime-packaging` | removed | 1 |

## B1 - Final Inventory Snapshot

Status: pending

| Metric | Unit | Final | Target | Result |
|---|---:|---:|---:|---|
| Embedded Anvien source skills | files | pending | 3 | pending |
| Generated Anvien skill directories | directories | pending | 3 | pending |
| Generated `Skill Selection Guide` rows | rows | pending | 3 | pending |
| Removed skill references in generated current output | matches | pending | 0 | pending |
