# Benchmark Ledger

Title: File Detail Command And Planner Rule
Date: 2026-06-06
Status: Initialized
Companion plan: `docs/plans/2026-06-06-file-detail-command-and-planner-rule/2026-06-06-file-detail-command-and-planner-rule-plan.md`
Companion evidence: `docs/plans/2026-06-06-file-detail-command-and-planner-rule/2026-06-06-file-detail-command-and-planner-rule-evidence.md`

## Benchmark Rules

- Record measured inventory, response-size, runtime, or graph counts only.
- Build/test pass-fail belongs in the evidence ledger unless timing/count/size is the measured target.
- Historical reports/plans are excluded from active-reference cleanup counts unless explicitly stated.

## B0 - Baseline Inventory

| Metric | Unit | Baseline | Latest | Final | Target | Evidence |
|--------|------|----------|--------|-------|--------|----------|
| Graph scanned files after plan precheck | files | 1406 | 1406 | pending | no unexpected shrink | `anvien analyze --force` |
| Graph nodes after plan precheck | nodes | 83978 | 83978 | pending | no unexpected shrink | `anvien analyze --force` |
| Graph relationships after plan precheck | relationships | 122367 | 122367 | pending | no unexpected shrink | `anvien analyze --force` |
| File projection dependency edges after plan precheck | edges | 16549 | 16549 | pending | no unexpected shrink | `anvien analyze --force` |
| Active files containing `file-context`, excluding `docs/plans/**` and `reports/**` | files | 22 | 22 | pending | 0 active old-name files, except justified non-active/historical residues | `rg -l ... 'file-context' \| Measure-Object` |
| Active `file-context` line hits in selected active roots | lines | 41 | 41 | pending | 0 active old-name instructions or registrations | `rg -n ... 'file-context' cmd internal README.md RUNBOOK.md AGENTS.md CLAUDE.md \| Measure-Object` |

## B1 - Pending Rename Measurements

| Metric | Unit | Baseline | Latest | Final | Target | Evidence |
|--------|------|----------|--------|-------|--------|----------|
| CLI command help entries for `file-detail` | entries | pending | pending | pending | at least 1 | future `anvien --help` / `anvien file-detail --help` |
| CLI command help entries for `file-context` | entries | pending | pending | pending | 0 | future `anvien file-context --help` |
| Active `/api/file-detail` route registrations, if API route is renamed | routes | pending | pending | pending | 1 | future source/contract/API smoke |
| Active `/api/file-context` route registrations, if API route is renamed | routes | pending | pending | pending | 0 | future source/contract/API smoke |
| Generated root files with approved planner rule | files | 0 | pending | pending | 2 | future `rg` over `AGENTS.md` and `CLAUDE.md` |

## B2 - Pending Runtime Measurements

| Metric | Unit | Baseline | Latest | Final | Target | Evidence |
|--------|------|----------|--------|-------|--------|----------|
| `file-detail` JSON smoke response size | bytes | pending | pending | pending | record | future command smoke |
| `file-detail` JSON smoke symbol count | symbols | pending | pending | pending | record | future command smoke |
| `file-detail` JSON smoke relationship count | relationships | pending | pending | pending | record | future command smoke |
| `/api/file-detail` smoke status, if API route is renamed | HTTP status | pending | pending | pending | 200 | future API smoke |
