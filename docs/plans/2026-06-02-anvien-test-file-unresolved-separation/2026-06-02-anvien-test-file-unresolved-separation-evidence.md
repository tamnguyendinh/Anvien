# Anvien Test File Unresolved Separation Evidence Ledger

Date: 2026-06-02

Status: Active

Companion files:

- Plan: [2026-06-02-anvien-test-file-unresolved-separation-plan.md](2026-06-02-anvien-test-file-unresolved-separation-plan.md)
- Benchmark ledger: [2026-06-02-anvien-test-file-unresolved-separation-benchmark.md](2026-06-02-anvien-test-file-unresolved-separation-benchmark.md)

## Evidence Rules

1. Record Anvien command evidence for implementation slices.
2. Do not use Anvien for doc-only planning commits.
3. Keep quantitative inventory and before/after counts in the benchmark ledger.
4. Record impact/blast-radius before editing graph builders, file projection, hotspot ranking, API contracts, or Web graph/file views.
5. Record generated output and UI checks only after the normal generation/build path creates them.
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

## E0 - Plan Creation

Date: 2026-06-02

Status: recorded

Scope:

- Created the standard three-file plan set for separating test-file unresolved from default production unresolved signal.
- No implementation source files changed.
- No Anvien command was run for this doc-only plan creation.

Source / command evidence:

| Check | Result |
|---|---|
| User problem statement | Test files only need to display as `Test File` and show what they test; unresolved details inside test files do not help the default production graph. |
| Prior analyze output from the current session | Top 5 unresolved hotspots were all test/e2e files, with unresolved counts from 856 to 1445. |
| Plan convention | This plan uses the standard `docs/plans/YYYY-MM-DD-<slug>/` directory with matching plan, evidence, and benchmark files. |

Impact / blast radius:

| Target | Result |
|---|---|
| Implementation code | Not run; no implementation edits in this planning step. |

Validation:

| Command | Result |
|---|---|
| File creation | Plan, evidence, and benchmark ledgers created under `docs/plans/2026-06-02-anvien-test-file-unresolved-separation/`. |

Failures / handling:

- None.
