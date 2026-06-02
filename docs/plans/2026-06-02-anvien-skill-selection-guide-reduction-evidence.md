# Anvien Skill Selection Guide Reduction Evidence Ledger

Date: 2026-06-02

Status: Active

Companion files:

- Plan: [2026-06-02-anvien-skill-selection-guide-reduction-plan.md](2026-06-02-anvien-skill-selection-guide-reduction-plan.md)
- Benchmark ledger: [2026-06-02-anvien-skill-selection-guide-reduction-benchmark.md](2026-06-02-anvien-skill-selection-guide-reduction-benchmark.md)

## Evidence Rules

1. Record Anvien command evidence for implementation slices.
2. Do not use Anvien for doc-only planning commits.
3. Keep quantitative inventory counts in the benchmark ledger.
4. Record impact/blast-radius before editing generator functions or retained workflow owners.
5. Record generated output checks after regeneration.
6. Record `anvien detect-changes --repo Anvien --scope all` before implementation commits.

## Evidence Template

Use this template for implementation phases:

```text
## E<n> - <Phase/Task>

Date:

Status:

Scope:

- ...

Source / command evidence:

| Check | Result |
|---|---|
| ... | ... |

Impact / blast radius:

| Target | Result |
|---|---|
| ... | ... |

Implementation evidence:

| File | Evidence |
|---|---|
| ... | ... |

Validation:

| Command | Result |
|---|---|
| ... | ... |

Failures / handling:

- ...

Detect changes:

| Command | Result |
|---|---|
| `anvien detect-changes --repo Anvien --scope all` | ... |

Commit:

- `<hash> <subject>`
```

## E0 - Plan Structure Review

Date: 2026-06-02

Status: recorded

Scope:

- Created plan, evidence ledger, and benchmark ledger using the existing `docs/plans` structure.
- No implementation source files changed in this planning step.
- No Anvien command was used because this is doc-only planning.

Source / command evidence:

| Check | Result |
|---|---|
| Reviewed `docs/plans/2026-06-01-anvien-analyze-file-classification-metrics-plan.md` | Confirmed plan convention: Date, Status, companion ledgers, Master Rules, Goal, Problem, Scope, Requirements, Invariants, Technical Direction, Definition Of Done, Phase Checklist, Risk Notes. |
| Reviewed `docs/plans/2026-05-23-anvien-skill-system-upgrade-plan.md` | Confirmed generated AI-context plans treat generated files as validation artifacts and source files as `internal/aicontext/aicontext.go` plus embedded Markdown. |
| `Get-ChildItem internal\aicontext\skills -Filter 'anvien-*.md'` | Current embedded Anvien skill inventory has 10 files. |
| `Get-ChildItem .claude\skills\anvien` | Current generated Anvien skill inventory has 10 directories. |

Impact / blast radius:

| Target | Result |
|---|---|
| Implementation symbols | Not run; no implementation edits in this planning step. |

Validation:

| Command | Result |
|---|---|
| `git status --short` | Shows only `docs/plans` doc-only planning files changed for this work. |

Failures / handling:

- None.
