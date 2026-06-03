# Anvien File Role Classification Gap Evidence Ledger

Date: 2026-06-03

Status: Complete

Companion files:

- Plan: [2026-06-03-anvien-file-role-classification-gap-plan.md](2026-06-03-anvien-file-role-classification-gap-plan.md)
- Benchmark ledger: [2026-06-03-anvien-file-role-classification-gap-benchmark.md](2026-06-03-anvien-file-role-classification-gap-benchmark.md)

## Evidence Rules

1. Record Anvien commands used to discover the owner and baseline.
2. Keep quantitative benchmark tables in the benchmark ledger.
3. For code changes, record impact/blast-radius before edits.
4. Preserve the distinction between raw unresolved, default-visible unresolved, and file-role classification.
5. Record failures and their handling.
6. Record `anvien detect-changes --repo Anvien --scope all` before each implementation commit.

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

## E0 - Baseline Analyze Refresh

Date: 2026-06-03

Status: completed

Command:

```powershell
.\anvien\bin\anvien.exe analyze --force
```

Result:

| Field | Value |
|---|---:|
| Files scanned | 821 |
| Parsed code files | 601 |
| Failed files | 0 |
| Indexed documents | 114 |
| Indexed metadata | 99 |
| Indexed analyzers | 0 |
| Indexed scripts | 4 |
| Indexed static | 3 |
| Unsupported language gaps | 0 |
| Unknown gaps | 0 |
| Graph nodes | 60974 |
| Graph relationships | 96624 |
| File projection files | 821 |
| File projection dependency edges | 15965 |
| Default-visible unresolved files | 336 |
| Raw unresolved files | 353 |
| Raw-only file difference | 17 |

Evidence interpretation:

- Analyze succeeded with no failed files.
- The `353 - 336 = 17` difference is not parse failure or unknown language inventory.
- The issue is classification expressiveness for known raw-only files.

## E1 - Raw-Only File Evidence

Date: 2026-06-03

Status: completed

Commands:

```powershell
.\anvien\bin\anvien.exe file-hotspots --repo Anvien --json --limit 0 --sort raw-unresolved --unresolved-only
.\anvien\bin\anvien.exe file-context internal/frameworks/frameworks.go --repo Anvien --json
.\anvien\bin\anvien.exe file-context internal/scopeir/sort_keys.go --repo Anvien --json
.\anvien\bin\anvien.exe file-context internal/group/types.go --repo Anvien --json
.\anvien\bin\anvien.exe file-context internal/repo/paths.go --repo Anvien --json
```

Raw-only file summary:

| Metric | Value |
|---|---:|
| Raw unresolved files | 353 |
| Default-visible unresolved files | 336 |
| Raw-only files | 17 |
| Raw-only source sites | 376 |
| Raw-only production source sites | 0 |
| Raw-only non-actionable source sites | 376 |

Representative samples:

| File | Raw | Classification | Sample targets |
|---|---:|---|---|
| `internal/frameworks/frameworks.go` | 209 | `builtin=2`, `standard_library=207` | `strings.ReplaceAll`, `strings.ToLower`, `strings.HasPrefix`, `float64` |
| `internal/scopeir/sort_keys.go` | 63 | `builtin=63` | `int`, `string` |
| `internal/group/types.go` | 16 | `builtin=16` | `int`, `map[string]string`, `map[string]RepoSnapshot` |
| `internal/repo/paths.go` | 13 | `standard_library=13` | `os.Getenv`, `os.UserHomeDir`, `filepath.Join` |

Evidence interpretation:

- The raw-only files are recognized source files.
- Their unresolved sites are all non-actionable builtin, standard library, or test-framework targets.
- The user-facing gap is that Anvien lacks a concise role label explaining these files as backend support/model/helper files.

## E2 - Owner Discovery

Date: 2026-06-03

Status: completed

Commands:

```powershell
.\anvien\bin\anvien.exe query "file classification appLayer functionalArea file role unresolved file projection" --repo Anvien
.\anvien\bin\anvien.exe context "ClassifyAppLayer" --repo Anvien
.\anvien\bin\anvien.exe context "ClassifyFunctionalArea" --repo Anvien
```

Owner evidence:

| Owner | Evidence |
|---|---|
| `internal/semantic/app_layer.go` | Query rank 1 for app-layer classification; `ClassifyAppLayer` owns existing app-layer path classification. |
| `internal/semantic/functional_area.go` | Query rank 2 for functional-area classification; `ClassifyFunctionalArea` owns existing functional-area path classification. |
| `internal/filecontext/context.go` | Query rank 2 file-layer owner; `FileSummary` owns file summary fields and unresolved buckets. |
| `anvien-web/src/components/FileMapPanel.tsx` | Query result and search evidence show file-summary Web consumer if role labels surface in UI. |
| `anvien-web/src/components/FileDetailPanel.tsx` | Query result and search evidence show file-detail consumer if role labels surface in UI. |

## E3 - Impact Baseline For Likely Classifier Edits

Date: 2026-06-03

Status: completed

Commands:

```powershell
.\anvien\bin\anvien.exe impact symbol "ClassifyAppLayer" --uid "Function:internal/semantic/app_layer.go:ClassifyAppLayer#1" --repo Anvien --direction upstream
.\anvien\bin\anvien.exe impact symbol "ClassifyFunctionalArea" --uid "Function:internal/semantic/functional_area.go:ClassifyFunctionalArea#1" --repo Anvien --direction upstream
```

Blast radius:

| Target | Risk | Impacted count | Affected files | Affected app layers | Affected functional areas |
|---|---|---:|---:|---|---|
| `ClassifyAppLayer` | CRITICAL | 6 | 4 | `backend=6` | `analyzer=1`, `cli=1`, `graph_health=1`, `unknown=3` |
| `ClassifyFunctionalArea` | CRITICAL | 4 | 4 | `backend=4` | `analyzer=1`, `cli=1`, `graph_health=1`, `unknown=1` |

Directly affected files from impact output:

- `internal/analyze/analyze.go`
- `internal/cli/command.go`
- `internal/graphaccuracy/access_candidate.go`
- `internal/semantic/app_layer.go`

Impact interpretation:

- Classifier edits are allowed but must be narrow.
- The plan should prefer additive file-role classification over broad rewrites of existing app-layer or functional-area behavior.

## E4 - Plan Creation

Date: 2026-06-03

Status: completed

Files created:

| File | Evidence |
|---|---|
| `docs/plans/2026-06-03-anvien-file-role-classification-gap/2026-06-03-anvien-file-role-classification-gap-plan.md` | Plan controller with goal, scope, requirements, phase checklist, and risk notes. |
| `docs/plans/2026-06-03-anvien-file-role-classification-gap/2026-06-03-anvien-file-role-classification-gap-evidence.md` | Evidence ledger seeded with baseline, raw-only file evidence, owner discovery, and impact baseline. |
| `docs/plans/2026-06-03-anvien-file-role-classification-gap/2026-06-03-anvien-file-role-classification-gap-benchmark.md` | Benchmark ledger seeded with raw/default unresolved and role-coverage baseline. |

Implementation evidence:

- No product code was edited in P0-A.
- Plan status is ready for implementation.

## E5 - Web UI Scope Refinement

Date: 2026-06-03

Status: completed

User report:

- The file-role classification gap affects Web UI display and needs explicit handling in the plan.

Commands:

```powershell
.\anvien\bin\anvien.exe analyze --force
.\anvien\bin\anvien.exe query "Web UI file role FileMapPanel FileDetailPanel FileTreePanel file summary display" --repo Anvien
rg -n 'FileSummary|file\.kind|file\.appLayer|file\.functionalArea|summary\.appLayer|summary\.functionalArea' anvien-web/src/components anvien-web/test -S
rg -n 'FileSummary|fileHotspots|file-context|file context|FileHotspotsResponse' internal/contracts/web_ui.go anvien-web/src/generated/anvien-contracts.ts -S
```

Source / command evidence:

| Check | Result |
|---|---|
| Analyze refresh | Pass. `files.scanned=824`, `parsed_code=601`, `failed=0`, `nodes=61005`, `relationships=96661`, `unresolvedFiles=336`, `rawUnresolvedFiles=353`. |
| Anvien Web owner query | Identified `FileTreePanel`, `FileDetailPanel`, and `FileMapPanel` as Web UI/file summary display owners. |
| `FileMapPanel` search | Uses `FileSummary`; displays `file.kind`, `file.appLayer`, and `file.functionalArea`. |
| `FileDetailPanel` search | Uses `FileSummary`; displays Layer and Area pills from `summary.appLayer` and `summary.functionalArea`. |
| Contract search | `internal/contracts/web_ui.go` defines `FileSummary`; generated Web contract mirrors it in `anvien-web/src/generated/anvien-contracts.ts`. |

Plan update evidence:

| File | Evidence |
|---|---|
| Plan | Added `Web UI Direction`; replaced conditional Web phase with explicit API/generated-type phase and Web display phase. |
| Evidence | Added this Web UI refinement evidence. |
| Benchmark | Added Web UI consumer coverage and refreshed latest analyze inventory counts. |

Implementation evidence:

- No product code was edited.
- The plan now requires Web contract and Web display handling when `fileRole` is added to `FileSummary`.

## E6 - Plan Readiness Review

Date: 2026-06-03

Status: completed

Scope:

- Re-read plan/evidence/benchmark after Web UI refinement.
- Check for contradictions between requirements, phase gates, Definition of Done, evidence, and benchmark targets.

Commands:

```powershell
.\anvien\bin\anvien.exe analyze --force
Get-Content docs\plans\2026-06-03-anvien-file-role-classification-gap\2026-06-03-anvien-file-role-classification-gap-plan.md
Get-Content docs\plans\2026-06-03-anvien-file-role-classification-gap\2026-06-03-anvien-file-role-classification-gap-evidence.md
Get-Content docs\plans\2026-06-03-anvien-file-role-classification-gap\2026-06-03-anvien-file-role-classification-gap-benchmark.md
```

Source / command evidence:

| Check | Result |
|---|---|
| Analyze refresh | Pass. `files.scanned=824`, `parsed_code=601`, `failed=0`, `nodes=61008`, `relationships=96664`, `unresolvedFiles=336`, `rawUnresolvedFiles=353`. |
| Plan consistency | Found that the role mapping table was incomplete and still had ambiguous `model or parser_model` targets while the benchmark table had concrete targets. |
| Web validation rule | Found that the Web validation wording allowed browser evidence to look interchangeable with e2e coverage; this was tightened to require Web/e2e tests when visible UI behavior changes. |
| Phase overlap | Clarified P2-A as CLI/graph-quality command output and P2-B as API/generated Web contract work. |

Plan update evidence:

| File | Evidence |
|---|---|
| Plan | Added checklist-update master rule; replaced the partial mapping table with all 17 raw-only files; tightened Web/e2e validation language; clarified P2-A/P2-B responsibility split. |
| Benchmark | Updated latest analyze graph node/relationship counts to the readiness-review analyze output. |

Readiness conclusion:

- The plan is ready for implementation after this review.
- Remaining risks are normal implementation risks: CRITICAL classifier blast radius, additive `FileSummary` contract propagation, and Web role display validation.

## E7 - Implementation Impact And Scope

Date: 2026-06-03

Status: completed

Scope:

- Add backend-owned `fileRole` classification for file summaries.
- Surface role labels through CLI/MCP/file projection output and generated Web contracts.
- Update Web file map/detail display and tests.

Source / command evidence:

| Check | Result |
|---|---|
| `.\anvien\bin\anvien.exe analyze --force` | Pass before graph-based implementation work. `files.scanned=824`, `parsed_code=601`, `failed=0`, `nodes=61009`, `relationships=96665`, `unresolvedFiles=336`, `rawUnresolvedFiles=353`. |
| Anvien query for file role owners | Identified `internal/filecontext/context.go`, `internal/contracts/web_ui.go`, `FileMapPanel`, and `FileDetailPanel` as direct owners. |
| `anvien api shape-check --repo Anvien --json` | No routes with both response shapes and consumers found for this contract check. |
| `anvien api impact --repo Anvien --json` | Command requires a route or file; no broad route impact was available from the no-target invocation. |

Impact / blast radius:

| Target | Result |
|---|---|
| `Struct:internal/filecontext/context.go:FileSummary` | CRITICAL. Impacted count 62; affected app layers included `api=23`, `backend=37`, `frontend=2`; affected areas included CLI, MCP, API, and Web graph UI consumers. |
| `Method:internal/filecontext/context.go:Builder.BuildFileContext#2` | LOW. Direct file-context summary construction path. |
| `Method:internal/filecontext/context.go:Builder.buildFileSummaries#0` | CRITICAL. Shared file list/file-hotspots summary path. |
| `Function:internal/semantic/metadata.go:SemanticTermDefinitions#0` | CRITICAL. Generated semantic metadata contract path. |
| `Function:internal/contracts/web_ui.go:WebUIContract#0` | CRITICAL. Generated Web manifest contract path. |
| `Function:internal/contracts/web_ui.go:WebUIContractTypeScript#0` | CRITICAL. Generated Web TypeScript contract path. |
| `Function:anvien-web/src/components/FileMapPanel.tsx:FileMapPanel` | LOW. Web row display; direct affected files included `App.tsx`, `FileTreePanel.tsx`, and unit tests. |
| `Function:anvien-web/src/components/FileDetailPanel.tsx:FileDetailPanel` | LOW. Web detail display; direct affected files included `CodeReferencesPanel.tsx`, `App.tsx`, and unit tests. |

Implementation evidence:

| File | Evidence |
|---|---|
| `internal/semantic/file_role.go` | Added deterministic role taxonomy and classifier with `model`, `contract_model`, `helper`, `storage_helper`, `config`, `adapter`, `fallback_adapter`, `test_helper`, `analyzer_helper`, `parser_model`, `runtime_model`, and `unknown`. |
| `internal/semantic/file_role_test.go` | Encodes all 17 raw-only target mappings and boundary behavior, including unknown fallback for analyzer area alone. |
| `internal/filecontext/context.go` | Added `FileSummary.FileRole` and classification in file-context detail/list builders without changing unresolved bucket calculations. |
| `internal/cli/command.go`, `internal/cli/file_context_command.go`, `internal/cli/graph_health_command.go` | Added role display in analyze/file-hotspots/graph-health file projection output. |
| `internal/mcp/target_dispatch.go`, `internal/mcp/resources.go` | Added role to file summary hints/selected file payloads and repo resource hotspot lines when a `FileSummary` exists. |
| `internal/contracts/web_ui.go`, `contracts/web-ui/anvien-web-contract.schema.json`, `anvien-web/src/generated/anvien-contracts.ts` | Added file role enum/labels and `FileSummary.fileRole` generated contract shape. |
| `anvien-web/src/components/FileMapPanel.tsx` | Shows backend-provided role in the compact metadata column; no Web path-pattern classification added. |
| `anvien-web/src/components/FileDetailPanel.tsx` | Adds `Role` pill near Layer/Area/Kind. |
| `anvien-web/src/components/FileTreePanel.tsx` | Inspected. It embeds `FileMapPanel` and does not directly render `FileSummary` role metadata, so no direct edit was needed. |

Failures / handling:

- Avoided a broad `analyzer_helper` rule that would classify all analyzer-area files as helper files. The classifier now requires analyzer area plus framework/COBOL helper paths for that role.

## E8 - Validation And Role Coverage

Date: 2026-06-03

Status: completed

Validation:

| Command | Result |
|---|---|
| `powershell -ExecutionPolicy Bypass -File anvien-launcher\build.ps1` | Pass. Go build and Web `tsc -b && vite build` completed. Vite reported existing chunk-size/dynamic-import warnings. |
| `go test ./internal/semantic ./internal/filecontext ./internal/contracts ./internal/cli ./internal/httpapi ./internal/mcp` | Pass. All 6 focused packages passed. |
| `npm test -- FileMapPanel.test.tsx FileDetailPanel.test.tsx` in `anvien-web` | Pass. 2 files, 8 tests passed. |
| `npm run test:e2e -- file-map-test-unresolved.spec.ts` in `anvien-web` | Pass. Command returned exit code 0. |
| Playwright screenshot smoke | Pass after scoping duplicate text locator. Saved `.tmp/file-role-web-validation.png`; confirmed File Map and File Detail showed `Test Helper`/`Test File`. |
| `.\anvien\bin\anvien.exe analyze --force` | Pass after implementation. `files.scanned=826`, `parsed_code=603`, `failed=0`, `nodes=61130`, `relationships=96891`, `unresolvedFiles=337`, `rawUnresolvedFiles=354`. |
| `.\anvien\bin\anvien.exe file-hotspots --repo Anvien --sort raw-unresolved --limit 5` | Pass. Human output includes `Path Role Layer Area ...`; analyze hotspot lines include `role=...`. |
| `.\anvien\bin\anvien.exe graph-health summary --repo Anvien --json` | Pass. `summary.nodeCount=61130`, `fileLayer.totalFiles=826`, `fileLayer.unresolvedFiles=337`, `fileLayer.rawUnresolvedFiles=354`, `summary.resolutionGapCount=34271`. |

Role coverage:

| Metric | Result |
|---|---:|
| Raw-only files | 17 |
| Raw-only files with unknown or empty role | 0 |
| Raw-only production unresolved files | 0 |
| Raw-only raw source sites | 376 |
| Raw-only non-actionable source sites | 376 |

Raw-only file role inventory:

| File | Role | Raw sites |
|---|---|---:|
| `internal/cli/exit_error.go` | `helper` | 2 |
| `internal/cobol/copy_expander.go` | `analyzer_helper` | 9 |
| `internal/frameworks/frameworks.go` | `analyzer_helper` | 209 |
| `internal/group/types.go` | `contract_model` | 16 |
| `internal/lbugnative/runner_default.go` | `fallback_adapter` | 1 |
| `internal/lbugnative/runner.go` | `adapter` | 1 |
| `internal/parser/metrics.go` | `parser_model` | 8 |
| `internal/repo/paths.go` | `storage_helper` | 13 |
| `internal/repo/runtime_config.go` | `config` | 10 |
| `internal/repo/settings.go` | `config` | 11 |
| `internal/resolution/source_site.go` | `helper` | 4 |
| `internal/scopeir/facts.go` | `parser_model` | 4 |
| `internal/scopeir/range.go` | `parser_model` | 4 |
| `internal/scopeir/sort_keys.go` | `helper` | 63 |
| `internal/session/error.go` | `runtime_model` | 6 |
| `internal/session/types.go` | `runtime_model` | 3 |
| `internal/testutil/path.go` | `test_helper` | 12 |

Failures / handling:

- First Playwright screenshot smoke failed because `getByText("Test File")` matched both the file row and detail pill. Reran with the locator scoped to `file-detail-section-summary`; validation passed.

Detect changes:

| Command | Result |
|---|---|
| `.\anvien\bin\anvien.exe detect-changes --repo Anvien --scope all` | Pass. Summary risk `critical`, affected_count 34, affected_files 20, changed_count 86, changed_files 22. Affected app layers: `api=1`, `api_contract=5`, `backend=16`, `mixed=12`. Affected functional areas: `cli=4`, `contracts=10`, `mcp=7`, `mixed=5`, `unknown=8`. File layer risk `high`; semantic status complete for app layer and functional area. |

Commit:

- `444dcdd feat: add file role classification`

## E9 - Commit Closure

Date: 2026-06-03

Status: completed

Scope:

- Close P4-A after implementation commit.
- Record the implementation commit hash in the plan artifact.

Source / command evidence:

| Check | Result |
|---|---|
| Implementation commit | `444dcdd feat: add file role classification` |
| Detect changes | Recorded in E8 before the implementation commit. |
| Remaining plan work | None. |

Commit:

- `444dcdd feat: add file role classification`
