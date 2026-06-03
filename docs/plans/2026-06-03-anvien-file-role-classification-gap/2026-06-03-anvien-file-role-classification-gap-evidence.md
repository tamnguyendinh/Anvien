# Anvien File Group Classification Evidence Ledger

Date: 2026-06-03

Status: Detect-changes complete - implementation commit pending

Companion files:

- Plan: [2026-06-03-anvien-file-role-classification-gap-plan.md](2026-06-03-anvien-file-role-classification-gap-plan.md)
- Benchmark ledger: [2026-06-03-anvien-file-role-classification-gap-benchmark.md](2026-06-03-anvien-file-role-classification-gap-benchmark.md)

## Evidence Rules

1. Record Anvien commands used to discover the owner and baseline.
2. Keep quantitative benchmark tables in the benchmark ledger.
3. For code changes, record impact/blast-radius before edits.
4. Preserve the distinction between raw unresolved, default-visible unresolved, file-role classification, and first-class file-group classification.
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

## E9 - Legacy File Role Commit Closure

Date: 2026-06-03

Status: superseded by corrected file group plan

Scope:

- Close P4-A after implementation commit.
- Record the implementation commit hash in the plan artifact.

Source / command evidence:

| Check | Result |
|---|---|
| Implementation commit | `444dcdd feat: add file role classification` |
| Detect changes | Recorded in E8 before the implementation commit. |
| Remaining plan work | The `fileRole` implementation is useful taxonomy foundation, but it does not close the corrected `fileGroup` requirement. The plan was reopened to create and surface `fileGroup=backend_support_model_helper`. |

Commit:

- `444dcdd feat: add file role classification`

## E10 - Corrected File Group Plan Reopen

Date: 2026-06-03

Status: completed

Scope:

- Reopen the plan around first-class file grouping, not only file role labeling.
- Confirm planner compliance before product-code implementation starts.
- Re-record current graph inventory and the 17-file anchor sample.

Source / command evidence:

| Check | Result |
|---|---|
| Planner skill read | `.claude/skills/anvien/anvien-planner/SKILL.md` read. The skill requires a standard three-file plan set and checklist items that are complete mini-plans. |
| Plan file structure | Pass. The plan has metadata, rules, goal, problem, scope, requirements, invariants, technical direction, definition of done, phase checklist, and risk notes. |
| Checklist mini-plan audit | Pass. 11 checklist items (`G0-A` through `G10-A`) include Goal, Work Steps, Implementation Gate, and Acceptance. |
| `.\anvien\bin\anvien.exe analyze --force` | Pass. `files.scanned=826`, `parsed_code=603`, `failed=0`, `nodes=61132`, `relationships=96893`, `dependencyEdges=16039`, `unresolvedFiles=337`, `rawUnresolvedFiles=354`. |
| `.\anvien\bin\anvien.exe file-hotspots --repo Anvien --json --limit 0 --sort path` | Used to record current sample axes after analyze refresh. |

Corrected problem statement:

| Item | Evidence |
|---|---|
| Product issue | Files that Anvien can identify as backend support/model/helper files are still not placed into a concrete file group users can read directly. |
| Prior implementation status | `444dcdd feat: add file role classification` added `fileRole`; it did not create `fileGroup`. |
| Correct target field | `fileGroup` |
| Correct target key | `backend_support_model_helper` |
| Correct target label | `Backend support/model/helper files` |
| Role relationship | `fileRole` remains a subcategory inside the file group. |
| Count rule | The 17 files are the required current anchor sample. The full group total must be measured from backend identity rules after implementation and must not be inferred from `rawUnresolvedFiles - unresolvedFiles`. |

Current 17-file anchor sample:

| File | Kind | App layer | Functional area | File role | Default unresolved | Raw unresolved |
|---|---|---|---|---|---:|---:|
| `internal/cli/exit_error.go` | `source` | `backend` | `cli` | `helper` | 0 | 2 |
| `internal/cobol/copy_expander.go` | `source` | `backend` | `analyzer` | `analyzer_helper` | 0 | 9 |
| `internal/frameworks/frameworks.go` | `source` | `backend` | `analyzer` | `analyzer_helper` | 0 | 209 |
| `internal/group/types.go` | `source` | `backend` | `query` | `contract_model` | 0 | 16 |
| `internal/lbugnative/runner_default.go` | `source` | `backend` | `storage` | `fallback_adapter` | 0 | 1 |
| `internal/lbugnative/runner.go` | `source` | `backend` | `storage` | `adapter` | 0 | 1 |
| `internal/parser/metrics.go` | `source` | `backend` | `providers` | `parser_model` | 0 | 8 |
| `internal/repo/paths.go` | `source` | `backend` | `storage` | `storage_helper` | 0 | 13 |
| `internal/repo/runtime_config.go` | `source` | `backend` | `storage` | `config` | 0 | 10 |
| `internal/repo/settings.go` | `source` | `backend` | `storage` | `config` | 0 | 11 |
| `internal/resolution/source_site.go` | `source` | `backend` | `resolution` | `helper` | 0 | 4 |
| `internal/scopeir/facts.go` | `source` | `backend` | `providers` | `parser_model` | 0 | 4 |
| `internal/scopeir/range.go` | `source` | `backend` | `providers` | `parser_model` | 0 | 4 |
| `internal/scopeir/sort_keys.go` | `source` | `backend` | `providers` | `helper` | 0 | 63 |
| `internal/session/error.go` | `source` | `backend` | `session` | `runtime_model` | 0 | 6 |
| `internal/session/types.go` | `source` | `backend` | `session` | `runtime_model` | 0 | 3 |
| `internal/testutil/path.go` | `source` | `backend` | `unknown` | `test_helper` | 0 | 12 |

Sample summary:

| Metric | Value |
|---|---:|
| Anchor sample files | 17 |
| Anchor sample files with `kind=source` | 17 |
| Anchor sample files with `appLayer=backend` | 17 |
| Anchor sample files with non-unknown `fileRole` | 17 |
| Anchor sample default unresolved source sites | 0 |
| Anchor sample raw unresolved source sites | 376 |

Plan update evidence:

| File | Evidence |
|---|---|
| Plan | Status changed to ready for implementation; `G0-A` and `G1-A` marked complete; group aggregation wording corrected so 17 is the required anchor sample, not a forced full-group total. |
| Evidence | Reopened ledger around file group classification and recorded current planner/analyze/sample evidence. |
| Benchmark | Reopened ledger around file group metrics and added file-group target measurements. |

Readiness conclusion:

- The corrected plan is ready for implementation after this update.
- Product-code implementation still requires Anvien impact checks before editing shared classifiers, file summaries, contracts, CLI/API output, graph/file projection behavior, or Web file views.

## E11 - Corrected File Group Implementation And Validation

Date: 2026-06-03

Status: implementation validated; detect-changes complete; commit pending

Scope:

- Add first-class `fileGroup=backend_support_model_helper`.
- Keep `fileRole` as the subcategory under the group.
- Surface the group through backend file projection, CLI/API/MCP, generated Web contracts, and Web file views.
- Prove the 17-file anchor sample belongs to the group without deriving membership from `rawUnresolvedFiles - unresolvedFiles`.

Source / command evidence:

| Check | Result |
|---|---|
| Planner skill read | `.claude/skills/anvien/anvien-planner/SKILL.md` read before updating plan/evidence/benchmark. |
| `.\anvien\bin\anvien.exe analyze --force` before implementation work | Pass. `files.scanned=828`, `parsed_code=605`, `failed=0`, `nodes=61337`, `relationships=97219`, `unresolvedFiles=338`, `rawUnresolvedFiles=355`. |
| Anvien owner queries | Confirmed owners in `internal/semantic`, `internal/filecontext/context.go`, CLI output files, `internal/contracts/web_ui.go`, HTTP API, MCP hints/resources, and Web `FileMapPanel`/`FileDetailPanel`. |
| `.\anvien\bin\anvien.exe analyze --force` after implementation build and ledger updates | Pass. `files.scanned=828`, `parsed_code=605`, `failed=0`, `nodes=61340`, `relationships=97225`, `dependencyEdges=16099`, `unresolvedFiles=338`, `rawUnresolvedFiles=355`. |

Impact / blast radius:

| Target | Result |
|---|---|
| `FileSummary` | CRITICAL. Broad API/backend/frontend consumer blast radius; file-summary contract change kept additive. |
| `FileList` | CRITICAL. File list consumers affected by new `fileGroups` aggregation; output shape kept additive. |
| `BuildFileContext` | LOW. Direct detail summary construction updated to attach `fileGroup`. |
| `BuildFileList` / `buildFileSummaries` | CRITICAL. Shared file projection path updated once so CLI/API/Web consume backend-owned group. |
| `SemanticTermDefinitions` | CRITICAL. Metadata taxonomy extended with `file_group`. |
| `WebUIContract` / `WebUIContractTypeScript` | CRITICAL. Generated Web contract updated and regenerated. |
| `buildAnalyzeFileProjection` / `analyzeFileProjectionLines` | CRITICAL. Analyze output now emits direct group lines and group field on hotspots. |
| `renderFileHotspots` / `graphHealthFileLayerLines` | CRITICAL. Human and machine CLI output now includes group identity. |
| `handleFileHotspots` | LOW. API response shape now includes `fileGroups`. |
| `mcpFileRelationshipHints` / `addMCPSymbolTargetFields` / context resource | CRITICAL/LOW mixed. MCP file payloads include backend group where `FileSummary` is available. |
| `FileMapPanel` / `FileDetailPanel` | LOW. Web display reads generated group labels and backend `fileGroup`; no Web path inference added. |

Implementation evidence:

| File | Evidence |
|---|---|
| `internal/semantic/file_group.go` | Added `FileGroupBackendSupportModelHelper`, exact key/label metadata, deterministic classifier, and boundary source strings. |
| `internal/semantic/file_group_test.go` | Freezes key/label, proves all 17 anchor files enter the group, and proves unknown role/frontend/test-kind/docs/config-kind/generated/missing-path boundaries stay out. |
| `internal/semantic/metadata.go` | Added `file_group` semantic term. |
| `internal/filecontext/context.go` | Added `FileSummary.FileGroup`, `FileList.FileGroups`, `FileGroupSummary`, shared group classification, and group aggregation from backend summaries. |
| `internal/filecontext/context_test.go` | Proves anchor sample `fileGroup=backend_support_model_helper`, sample default unresolved `0`, sample raw unresolved `376`, and role breakdown. |
| `internal/contracts/web_ui.go`, generated schema, generated TypeScript | Added `FILE_GROUPS`, `FILE_GROUP_LABELS`, `FileGroup`, `FileGroupLabel`, `FileSummary.fileGroup`, and `FileHotspotsResponse.fileGroups`. |
| `internal/cli/command.go`, `internal/cli/file_context_command.go`, `internal/cli/graph_health_command.go` | Analyze/file-hotspots/graph-health output now surfaces group lines and group columns. |
| `internal/httpapi/file_context.go` | `/api/file-hotspots` includes backend-computed `fileGroups`. |
| `internal/mcp/target_dispatch.go`, `internal/mcp/resources.go` | MCP file hints and repo context hotspot lines include `fileGroup`. |
| `anvien-web/src/components/FileMapPanel.tsx` | File rows show group label first, then role, then layer/area. |
| `anvien-web/src/components/FileDetailPanel.tsx` | File detail summary shows `Group` before `Role`, `Layer`, and `Area`. |
| `anvien-web/e2e/file-map-test-unresolved.spec.ts` | E2E fixture and assertion prove the Web shows `Backend support/model/helper files`. |

Graph File-node enrichment outcome:

- Current analyze File nodes already carry `appLayer` and `functionalArea`.
- `fileRole` and `fileGroup` are authoritative in the shared backend `FileSummary` / file projection path for this slice.
- CLI/API/Web/MCP read the backend summary field; no Web-only or display-only path classifier was added.
- Direct graph-node `fileGroup` property enrichment remains a follow-up only if the graph model later requires file identity fields as persisted node properties.

Validation:

| Command | Result |
|---|---|
| `powershell -ExecutionPolicy Bypass -File anvien-launcher\build.ps1` | First run failed on TypeScript literal map key type; fixed with `Map<string, string>`. Second run passed. Vite reported existing chunk-size/dynamic-import warnings. |
| `go test ./internal/semantic ./internal/filecontext ./internal/contracts ./internal/cli ./internal/httpapi ./internal/mcp` | Pass. |
| `npm test -- FileMapPanel.test.tsx FileDetailPanel.test.tsx` | First run failed because the updated UI has two legitimate `Unknown` labels in the unknown-group/unknown-role test; test was corrected. Second run passed: 2 files, 8 tests. |
| `npm run test:e2e -- file-map-test-unresolved.spec.ts` | Pass. 1 Chromium test passed with Vite dev server started temporarily for the run. |
| `.\anvien\bin\anvien.exe analyze --force` | Pass. `files.scanned=828`, `parsed_code=605`, `failed=0`, `nodes=61340`, `relationships=97225`. Analyze output includes direct `fileProjection.group key="backend_support_model_helper" label="Backend support/model/helper files" files=42 defaultUnresolved=1073 rawUnresolved=2087 ...`. |
| `.\anvien\bin\anvien.exe file-hotspots --repo Anvien --json --sort path --limit 0` | Pass. `fileGroups[backend_support_model_helper]` reports `files=42`, `defaultUnresolved=1073`, `rawUnresolved=2087`; anchor missing count is `0`. |
| `.\anvien\bin\anvien.exe graph-health summary --repo Anvien --json` | Pass. `fileLayer.totalFiles=828`, `unresolvedFiles=338`, `rawUnresolvedFiles=355`, `fileGroups.Count=1`, target group files `42`. |
| `.\anvien\bin\anvien.exe file-context internal/repo/runtime_config.go --repo Anvien --json` | Pass. Summary has `kind=source`, `appLayer=backend`, `functionalArea=storage`, `fileRole=config`, `fileGroup=backend_support_model_helper`, default unresolved `0`, raw unresolved `10`. |

Measured group coverage:

| Metric | Result |
|---|---:|
| Full group files from backend identity rules | 42 |
| Full group default-visible unresolved source sites | 1073 |
| Full group raw unresolved source sites | 2087 |
| Anchor sample files missing from group | 0 |
| Anchor sample default-visible unresolved source sites | 0 |
| Anchor sample raw unresolved source sites | 376 |

Failures / handling:

- TypeScript build initially rejected `Map.get(value)` because generated `FILE_GROUP_LABELS` inferred the literal group key; fixed by widening the label map to `Map<string, string>`.
- Web unit test for unknown role initially assumed only one `Unknown` label; corrected to expect both unknown group and unknown role labels.

Detect changes:

| Command | Result |
|---|---|
| `.\anvien\bin\anvien.exe detect-changes --repo Anvien --scope all` | Pass. Summary risk `critical`; `affected_count=57`, `affected_files=23`, `changed_count=324`, `changed_files=25`. Affected app layers: `api=2`, `api_contract=5`, `backend=37`, `mixed=13`. Affected functional areas: `api=2`, `cli=19`, `contracts=10`, `mcp=7`, `mixed=11`, `unknown=8`. File layer changed risk `high`; `affectedFiles=23`, `changedFiles=25`. Resolution gap change inventory: `changedGapEntities=176`, `analyzer_gap=106`, `non_actionable=66`, `review=4`. Semantic app-layer and functional-area status remained complete. |

Commit:

- Pending implementation commit.
