# AVmatrix Web Filter-Based Clustered Layout And Manual Optimizer Evidence Ledger

Date: 2026-05-20

Status: active

Companion files:

- Plan: [2026-05-20-avmatrix-web-filter-based-clustered-layout-manual-optimizer-plan.md](2026-05-20-avmatrix-web-filter-based-clustered-layout-manual-optimizer-plan.md)
- Benchmark ledger: [2026-05-20-avmatrix-web-filter-based-clustered-layout-manual-optimizer-benchmark.md](2026-05-20-avmatrix-web-filter-based-clustered-layout-manual-optimizer-benchmark.md)

## Evidence Rules

Record evidence as each evidenced task is completed. Evidence should include commands, impacted files, test results, e2e artifacts, screenshots when useful, and concise observations needed to audit the plan later.

For doc-only commits, do not use AVmatrix.

Do not record inferred performance or behavior as fact. Every behavior claim must include a command, inspected file, test result, browser/e2e artifact, or runtime diagnostic.

## E0 - Plan Creation Evidence

Date: 2026-05-20

Status: recorded

Created file set:

- `docs/plans/2026-05-20-avmatrix-web-filter-based-clustered-layout-manual-optimizer-plan.md`
- `docs/plans/2026-05-20-avmatrix-web-filter-based-clustered-layout-manual-optimizer-evidence.md`
- `docs/plans/2026-05-20-avmatrix-web-filter-based-clustered-layout-manual-optimizer-benchmark.md`

Plan creation scope:

- Use the repository's established plan, evidence, and benchmark ledger format.
- Capture the user-approved design direction before implementation.
- Keep this as a Web UI layout plan, not a backend graph or analyzer plan.

Convention inspection commands:

```powershell
git status --short
Get-Content -Path docs\plans\2026-05-20-avmatrix-orphan-node-connectivity-lens-plan.md -TotalCount 80
Get-Content -Path docs\plans\2026-05-20-avmatrix-orphan-node-connectivity-lens-evidence.md -TotalCount 80
Get-Content -Path docs\plans\2026-05-20-avmatrix-orphan-node-connectivity-lens-benchmark.md -TotalCount 80
```

Observed planning standard:

- Plan file has date, status, companion files, rules, problem/scope, design decision, checklist, acceptance criteria, and closure definition.
- Evidence ledger records commands, inspected files, decisions, and validation artifacts.
- Benchmark ledger records product/runtime measurements and leaves unmeasured values pending.

## E1 - Product Decision Evidence

Date: 2026-05-20

Status: recorded

Accepted direction from discussion:

- Node clustering should be the default layout foundation.
- Clustering should be based only on existing Node Type filters.
- Each node type/filter should map to one visual cluster.
- Do not compute node importance, hubs, degree, centrality, or "command" status for placement.
- The purpose is visual clarity, not ranking code by connectivity.
- Layout optimization is a separate manual cleanup feature.
- Layout optimization must not run automatically after graph load.

Rejected directions:

- Do not add a new taxonomy for layout clusters.
- Do not move "important" or highly connected nodes to the center.
- Do not use runtime optimizer as the primary way to make the graph readable.
- Do not use elapsed-time budget as the layout correctness mechanism.

## E2 - Current Codebase Facts From Initial Inspection

Date: 2026-05-20

Status: recorded

Initial code inspection commands:

```powershell
rg --files -g "*graph*" -g "*Graph*" -g "*.tsx" -g "*.ts" avmatrix-web internal cmd docs | Select-Object -First 120
rg -n "layout|optimiz|force|simulation|dagre|elk|cytoscape|sigma|cluster|community|graph" avmatrix-web/src avmatrix-web/e2e avmatrix-web/test internal -S
Get-Content -Path avmatrix-web\src\hooks\useSigma.ts -TotalCount 260
Get-Content -Path avmatrix-web\src\lib\graph-adapter.ts -TotalCount 260
Get-Content -Path avmatrix-web\src\hooks\app-state\graph.tsx -TotalCount 320
Get-Content -Path avmatrix-web\src\components\GraphCanvas.tsx -TotalCount 320
Get-Content -Path avmatrix-web\src\lib\constants.ts -TotalCount 180
```

Observed current implementation:

- `GraphCanvas` converts the loaded `KnowledgeGraph` into a Sigma graph through `knowledgeGraphToGraphology`.
- `GraphCanvas` builds `communityMemberships` from `MEMBER_OF` relationships before conversion.
- `knowledgeGraphToGraphology` already owns initial node position assignment.
- `knowledgeGraphToGraphology` currently mixes structural parent placement, community placement, and random jitter.
- `useSigma.setGraph` currently starts layout after assigning the Sigma graph.
- `useSigma` imports and uses `graphology-layout-forceatlas2/worker`.
- `useSigma` also runs `noverlap.assign` during layout stop/cleanup paths.
- Existing node filter order and known labels are exposed through Web constants such as `FILTERABLE_LABELS`; implementation should not use `DEFAULT_VISIBLE_LABELS` as the cluster order source.

Implementation implication:

- The likely smallest code change is Web-only:
  - change initial placement policy in `graph-adapter.ts`;
  - stop auto-starting layout in `useSigma.ts`;
  - keep optimizer as a manual GraphCanvas action.

## E3 - Required AVmatrix Impact Evidence

Date: 2026-05-20

Status: pending

Before implementation edits, run:

```powershell
go run .\cmd\avmatrix analyze --force
go run .\cmd\avmatrix impact knowledgeGraphToGraphology --repo AVmatrix --direction upstream --depth 2 --include-tests
go run .\cmd\avmatrix impact useSigma --repo AVmatrix --direction upstream --depth 2 --include-tests
go run .\cmd\avmatrix impact GraphCanvas --repo AVmatrix --direction upstream --depth 2 --include-tests
```

Record:

- risk level;
- impacted callers;
- impacted tests;
- any HIGH or CRITICAL warning before editing.

## E4 - Implementation Evidence Placeholder

Date: 2026-05-20

Status: pending

Record implementation evidence here as slices complete:

- files changed;
- behavior changed;
- commands run;
- screenshots or browser observations if used;
- validation outputs.

## E5 - Validation Evidence Placeholder

Date: 2026-05-20

Status: pending

Minimum expected validation:

```powershell
npm --prefix avmatrix-web run test -- test/unit/graph-adapter.edge-geometry.test.ts test/unit/GraphCanvas.selection-performance.test.tsx test/unit/runtime-diagnostics.test.ts
npm --prefix avmatrix-web run test:e2e -- server-connect.spec.ts -g "loads graph" --workers=1 --timeout=120000
```

The `--timeout=120000` flag above is a Playwright runner guard for validation execution only. It is not an accepted product or layout correctness mechanism.

Final commands may differ if implementation adds more focused tests. Record the actual commands and results.
