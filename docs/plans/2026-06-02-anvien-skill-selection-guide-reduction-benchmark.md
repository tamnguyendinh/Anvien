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

Source evidence: E0, E1.

| Metric | Unit | Baseline | Target | Delta |
|---|---:|---:|---:|---:|
| Embedded Anvien source skills | files | 10 | 4 | -6 |
| Generated Anvien skill directories | directories | 10 | 4 | -6 |
| Retained Anvien workflow skills | skills | 3 | 4 | +1 |
| Removed broad/router Anvien skills | skills | 7 | 0 | -7 |
| Generated `Skill Selection Guide` rows | rows | 0 | 4 | +4 |

Retained skill target:

| Skill | Unit | Target |
|---|---:|---:|
| `anvien-api-surface` | retained | 1 |
| `anvien-refactoring` | retained | 1 |
| `anvien-debugging` | retained | 1 |
| `anvien-planner` | retained | 1 |

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

## B1 - Source Inventory After P2

Status: recorded

Source evidence: E2.

| Metric | Unit | Latest | Target | Result |
|---|---:|---:|---:|---|
| Embedded Anvien source skills | files | 4 | 4 | pass |
| Registered Anvien workflow skills | skills | 4 | 4 | pass |
| Removed broad/router embedded source skills | files | 0 | 0 | pass |
| Planner embedded source skill | files | 1 | 1 | pass |

## B2 - Final Inventory Snapshot

Status: recorded

Source evidence: E3.

| Metric | Unit | Final | Target | Result |
|---|---:|---:|---:|---|
| Embedded Anvien source skills | files | 4 | 4 | pass |
| Generated Anvien skill directories | directories | 4 | 4 | pass |
| Generated `Skill Selection Guide` rows | rows | 4 | 4 | pass |
| Removed skill references in generated current output | matches | 0 | 0 | pass |
| Removed broad/router embedded source skill files | files | 0 | 0 | pass |
| Post-regeneration graph nodes | nodes | 96,204 | measured | recorded |
| Post-regeneration graph relationships | relationships | 131,703 | measured | recorded |
| Post-regeneration scanned files | files | 815 | measured | recorded |
