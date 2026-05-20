# AVmatrix Orphan Node Connectivity Lens Evidence Ledger

Date: 2026-05-20

Status: active

Companion files:

- Plan: [2026-05-20-avmatrix-orphan-node-connectivity-lens-plan.md](2026-05-20-avmatrix-orphan-node-connectivity-lens-plan.md)
- Benchmark ledger: [2026-05-20-avmatrix-orphan-node-connectivity-lens-benchmark.md](2026-05-20-avmatrix-orphan-node-connectivity-lens-benchmark.md)

## Evidence Rules

Record evidence as each evidenced task is completed. Evidence should include commands, impacted files, test results, e2e artifacts, and concise observations needed to audit the plan later.

For doc-only commits, do not use AVmatrix.

Do not record inferred graph counts. Every count must include the command, source graph, repo path, commit or graph timestamp when available, and interpretation.

## E0 - Plan Creation Evidence

Date: 2026-05-20

Status: recorded

Created file set:

- `docs/plans/2026-05-20-avmatrix-orphan-node-connectivity-lens-plan.md`
- `docs/plans/2026-05-20-avmatrix-orphan-node-connectivity-lens-evidence.md`
- `docs/plans/2026-05-20-avmatrix-orphan-node-connectivity-lens-benchmark.md`

User requirement:

- Create a proper repo-standard plan, not a chat-only outline.
- Do not create a speculative plan.
- Follow the repository's established planning rules and required companion ledgers.

Convention inspection commands:

```powershell
Get-ChildItem docs\plans | Sort-Object Name | Select-Object -Last 20 | Format-Table -AutoSize
rg -n "^# |^Status:|^## |Acceptance|Validation|Closure|Evidence|Benchmark|Zero-Trust|Phase" docs\plans
Get-Content docs\plans\2026-05-19-avmatrix-multilanguage-graph-filters-and-coverage-plan.md -TotalCount 280
Get-Content docs\plans\2026-05-19-avmatrix-multilanguage-graph-filters-and-coverage-evidence.md -TotalCount 160
Get-Content docs\plans\2026-05-19-avmatrix-multilanguage-graph-filters-and-coverage-benchmark.md -TotalCount 160
```

Observed planning standard:

- Plan file has date, status, companion files, rules, problem/scope, acceptance guardrails, phase checklist, ledger, and closure definition.
- Evidence ledger records commands, files, status, observations, and validation artifacts.
- Benchmark ledger records measured counts and benchmarkable inventory; unmeasured values remain pending.
- Doc-only plan creation does not require AVmatrix.

Plan creation decisions:

- Status is `active` because the user requested a formal plan for future implementation.
- No graph counts were invented.
- Baseline graph-health counts remain `pending measurement` until commands are run and recorded.
- "Orphan node" is defined as derived connectivity status, not a primary semantic node label.

## E1 - Initial Product Reasoning Evidence

Date: 2026-05-20

Status: recorded

Discussion summary:

- A user asked whether lonely/orphan nodes should be classified or mapped into a separate node/filter type to manage code, buggy functions, and unwired code.
- The accepted planning direction is to classify this as a graph-health/connectivity lens, not a semantic node label.
- The plan requires taxonomy and evidence before presenting any orphan status as a bug.

Non-speculative boundary:

- No current orphan counts are claimed.
- No current analyzer defect is claimed.
- No dead-code count is claimed.
- No UI implementation detail is considered accepted until inspected during implementation phases.

## E2 - Pending Baseline Evidence

Date: 2026-05-20

Status: pending

Required before implementation claims:

- measured connectivity inventory for `E:\AVmatrix-GO`;
- measured connectivity inventory for one large indexed repo when available;
- recorded edge policy used for the measurements;
- expected-isolated policy and count by reason;
- comparison of raw graph connectivity versus Web-visible connectivity.
