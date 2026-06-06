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
| Graph scanned files after plan precheck | files | 1406 | 1409 | 1409 | no unexpected shrink | `npm run full-build` final `anvien analyze . --force` |
| Graph nodes after plan precheck | nodes | 83978 | 84018 | 84018 | no unexpected shrink | `npm run full-build` final `anvien analyze . --force` |
| Graph relationships after plan precheck | relationships | 122367 | 122400 | 122400 | no unexpected shrink | `npm run full-build` final `anvien analyze . --force` |
| File projection dependency edges after plan precheck | edges | 16549 | 16551 | 16551 | no unexpected shrink | `npm run full-build` final `anvien analyze . --force` |
| Active files containing `file-context`, excluding `docs/plans/**` and `reports/**` | files | 22 | 3 | 3 | 0 active old-name files, except justified non-active/historical residues | final `rg` inventory: changelog plus negative tests only |
| Active `file-context` line hits in selected active roots | lines | 41 | 8 | 8 | 0 active old-name instructions or registrations | final `rg` inventory: changelog plus negative tests only |

## B1 - Pending Rename Measurements

| Metric | Unit | Baseline | Latest | Final | Target | Evidence |
|--------|------|----------|--------|-------|--------|----------|
| CLI command help entries for `file-detail` | entries | pending | 1 | 1 | at least 1 | `anvien --help` / `anvien file-detail --help` |
| CLI command help entries for `file-context` | entries | pending | 0 | 0 | 0 | `anvien --help` and `anvien file-context --help` |
| Active `/api/file-detail` route registrations, if API route is renamed | routes | pending | 1 | 1 | 1 | source/contract/API smoke |
| Active `/api/file-context` route registrations, if API route is renamed | routes | pending | 0 | 0 | 0 | source/contract/API smoke; old route appears only in negative test |
| Generated root files with approved planner rule | files | 0 | 2 | 2 | 2 | `rg` over `AGENTS.md` and `CLAUDE.md` |

## B2 - Pending Runtime Measurements

| Metric | Unit | Baseline | Latest | Final | Target | Evidence |
|--------|------|----------|--------|-------|--------|----------|
| `file-detail` JSON smoke response size | bytes | pending | 72831 | 72831 | record | `anvien file-detail internal/aicontext/aicontext.go --repo Anvien --json` |
| `file-detail` JSON smoke symbol count | symbols | pending | 56 | 56 | record | `anvien file-detail internal/aicontext/aicontext.go --repo Anvien --json` |
| `file-detail` JSON smoke relationship count | relationships | pending | 54 | 54 | record | `anvien file-detail internal/aicontext/aicontext.go --repo Anvien --json` |
| `/api/file-detail` smoke status, if API route is renamed | HTTP status | pending | 200 | 200 | 200 | API smoke with temporary `anvien serve` |
