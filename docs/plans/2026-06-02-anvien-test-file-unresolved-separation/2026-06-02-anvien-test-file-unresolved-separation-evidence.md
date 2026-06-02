# Anvien Test File Unresolved Separation Evidence Ledger

Date: 2026-06-02

Status: Active

Companion files:

- Plan: [2026-06-02-anvien-test-file-unresolved-separation-plan.md](2026-06-02-anvien-test-file-unresolved-separation-plan.md)
- Benchmark ledger: [2026-06-02-anvien-test-file-unresolved-separation-benchmark.md](2026-06-02-anvien-test-file-unresolved-separation-benchmark.md)

## Evidence Rules

1. Record Anvien command evidence for code/graph plan writing, plan review, and implementation slices.
2. Do not run Anvien only for doc-only commit ceremony.
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
- No Anvien command was run for the initial doc-only file creation; the later code/graph plan review is recorded in E1 with Anvien evidence.

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

## E1 - Plan Review With Anvien

Date: 2026-06-02

Status: recorded

Scope:

- Reviewed whether the plan direction matches current code and graph behavior.
- Updated plan direction before implementation.
- No implementation source files changed.

Source / command evidence:

| Check | Result |
|---|---|
| `anvien analyze --force` | Pass. Graph refreshed with 818 files scanned, 598 parsed code files, 0 failed parses, 96,340 nodes, 131,828 relationships, 590 files with unresolved, and top 5 hotspots all test/e2e files. |
| `anvien query files "file projection unresolved hotspot ranking ResolutionGap" --repo Anvien` | Confirmed unresolved hotspot query is dominated by test/e2e files. |
| `anvien query files "test file classification kind appLayer backend_test e2e" --repo Anvien` | Confirmed file summaries already expose `kind=test` and test app layers such as `backend_test`, `api_test`, and `frontend_test`. |
| `anvien query files "web graph file map unresolved ResolutionGap node display" --repo Anvien` | Confirmed Web-facing file map/detail behavior depends on file unresolved summary fields. |
| `rg` source inspection | Found primary owners: `internal/semantic/app_layer.go`, `internal/filecontext/context.go`, CLI analyze/file-hotspots/graph-health commands, `internal/httpapi/file_context.go`, generated Web contracts, `FileMapPanel`, and `FileDetailPanel`. |

Plan review decisions:

| Decision | Evidence |
|---|---|
| Do not invent a new test-file detector first. | Existing `kind=test` and test app layers already exist in graph/file summaries. |
| P1-A should reuse/harden classification truth, not recreate it. | `filecontext.fileKind` derives `test` from app-layer values, and semantic app-layer tests already cover backend/API/frontend test paths. |
| Bucket separation must include default risk/warning semantics. | Web rows currently use `unresolvedSourceSiteCount` for warning icon, `Unres`, totals, and file detail unresolved display. |
| Web UI must not hard-code path checks. | Backend/file projection already owns classification; UI should consume backend fields. |
| Test-to-target relationships must remain visible. | `filecontext` already tracks reverse linked-test counts; plan must ensure test-file view can show tested targets too. |

Implementation evidence:

| File | Evidence |
|---|---|
| Plan | Updated Master Rules, Technical Direction, Requirements, P0-A, P1-A, P1-B, P1-C, P2-A, P2-B, P3-A, and P4-A. |
| Benchmark ledger | Updated B0 baseline to the latest analyze output and added raw/default risk separation as a target metric. |

Validation:

| Command | Result |
|---|---|
| Plan review | P0-A owner/baseline discovery is complete; implementation phases P1-A onward remain pending. |

Failures / handling:

- Initial review found baseline drift from the first plan draft; B0 was refreshed.
- Initial review found P1-A was too broad because classification already exists; P1-A was narrowed to reuse and harden existing backend truth.

## E2 - P1-A Reuse And Harden Test-File Classification Truth

Date: 2026-06-02

Status: implemented

Scope:

- Reused existing backend app-layer classification as test-file truth.
- Added focused tests only; no runtime classification behavior was changed.
- Covered frontend e2e/spec/test paths, API test paths, backend test paths, and file summary `kind=test` output from app layers.

Source / command evidence:

| Check | Result |
|---|---|
| `anvien analyze --force` | Pass. Graph refreshed before graph-based work. Baseline top unresolved hotspots remain test/e2e files. |
| Source inspection | `internal/semantic/app_layer.go` maps test/e2e paths to `backend_test`, `api_test`, and `frontend_test`; `internal/filecontext/context.go:fileKind` maps test app layers to `kind=test`. |

Impact / blast radius:

| Target | Result |
|---|---|
| `anvien impact file internal/semantic/app_layer.go --repo Anvien --direction upstream` | HIGH/CRITICAL blast radius across analyze, contracts, CLI, MCP, semantic metadata, and Web consumers. P1-A avoided runtime source edits and added tests only. |
| `anvien impact file internal/filecontext/context.go --repo Anvien --direction upstream` | HIGH/CRITICAL blast radius across file projection, CLI/API, MCP resources, Web contracts, `FileMapPanel`, and `FileDetailPanel`. P1-A avoided runtime source edits and added tests only. |

Implementation evidence:

| File | Evidence |
|---|---|
| `internal/semantic/app_layer_test.go` | Added `TestClassifyAppLayerMapsTestPathBoundaries` for frontend e2e/unit/co-located tests, API tests, backend tests, fixture paths, and non-test source boundaries. |
| `internal/filecontext/context_test.go` | Added `TestBuildFileListUsesAppLayerAsTestClassificationTruth` proving `api_test`, `backend_test`, and `frontend_test` summaries expose `kind=test` while backend source remains `kind=source`. |
| Plan | Marked P1-A complete. |

Validation:

| Command | Result |
|---|---|
| `powershell -ExecutionPolicy Bypass -File anvien-launcher\build.ps1` | Pass; Vite emitted existing chunk-size/dynamic-import warnings. |
| `go test ./internal/semantic` | Pass. |
| `go test ./internal/filecontext` | Pass. |

Failures / handling:

- No runtime source change was needed for P1-A because classification truth already existed.
- No benchmark ledger update was required for P1-A because this slice added test coverage and did not change product/runtime/graph metrics.

Detect changes:

| Command | Result |
|---|---|
| `anvien detect-changes --repo Anvien --scope all` | Pass. Summary risk `low`; affected files: 4; changed files: 4; affected processes: 0; changed app layers: `backend_test`, `docs`. |

Commit:

- `9e42f3f test: lock test file classification truth`

## E3 - P1-B Separate Unresolved Metric Buckets

Date: 2026-06-02

Status: implemented

Scope:

- Added additive unresolved bucket fields to file summaries.
- Kept `unresolvedSourceSiteCount` as the raw compatibility count for this phase.
- Added raw/default-visible risk fields without changing the legacy `risk` field yet.
- Regenerated the Web TypeScript contract from the source contract.

Source / command evidence:

| Check | Result |
|---|---|
| `anvien analyze --force` | Pass. Graph refreshed before P1-B work; pre-edit graph had 818 files scanned, 598 parsed code files, 0 failed parses, 96,363 nodes, 131,889 relationships, and 590 files with unresolved. |
| `anvien impact file internal/filecontext/context.go --repo Anvien --direction upstream` | HIGH/CRITICAL blast radius across file projection, CLI/API, MCP, Web contracts, and file map/detail consumers. |
| `anvien impact file internal/contracts/web_ui.go --repo Anvien --direction upstream` | Generated-contract blast radius across `cmd/generate-web-contracts` and Web contract output. |
| Contract source inspection | `go run ./cmd/generate-web-contracts` emits TypeScript from `internal/contracts/web_ui.go`; the generated file did not change until the source contract interface was updated. |

Implementation evidence:

| File | Evidence |
|---|---|
| `internal/filecontext/context.go` | Added raw, production/actionable, test, non-actionable, unknown, and default-visible unresolved counts to `FileSummary`; added `rawRisk` and `defaultVisibleRisk`; list and detail summaries now use the same bucket helper. |
| `internal/filecontext/context.go` | Bucket rule: test files put all raw unresolved into the test bucket; non-test files split `non_actionable`, unknown/blank metadata, and production/actionable; default-visible currently equals production/actionable plus unknown. |
| `internal/filecontext/context_test.go` | Added `TestBuildFileSummariesSeparateUnresolvedBuckets`, proving raw counts remain traceable and test unresolved does not count as production/default-visible for both list and detail summaries. |
| `internal/contracts/web_ui.go` | Added the new fields to the Web `FileSummary` contract source. |
| `anvien-web/src/generated/anvien-contracts.ts` | Regenerated TypeScript contract now exposes the new `FileSummary` fields. |
| Plan | Marked P1-B complete. |
| Benchmark ledger | Recorded P1-B bucket field availability. |

Validation:

| Command | Result |
|---|---|
| `go run ./cmd/generate-web-contracts --check` | Pass. Generated Web contract is current. |
| `powershell -ExecutionPolicy Bypass -File anvien-launcher\build.ps1` | Pass; Vite emitted existing chunk-size/dynamic-import warnings. |
| `go test ./internal/filecontext ./internal/contracts ./internal/httpapi ./internal/cli` | Pass. |
| `anvien analyze --force` | Pass before detect-changes. Post-edit graph had 818 files scanned, 598 parsed code files, 0 failed parses, 96,521 nodes, 132,071 relationships, and 590 files with unresolved. Default top hotspots remain test files because P1-C has not changed ranking yet. |

Failures / handling:

- The first generator run made no changes because `internal/contracts/web_ui.go` is the source of the generated TypeScript interface; updated the source contract and regenerated.
- P1-B intentionally did not change default sort/filter/risk behavior; P1-C owns that behavior change.

Detect changes:

| Command | Result |
|---|---|
| `anvien detect-changes --repo Anvien --scope all` | Pass. Summary risk `high`; changed files: 7; affected files: 6; affected processes: 8; changed app layers: `api_contract`, `backend`, `backend_test`, `docs`; generated summaries include the new raw/default-visible bucket fields. |

Commit:

- `06e9c2f feat: expose unresolved bucket metrics`

## E4 - P1-C Change Hotspot Ranking To Actionable Default Signal

Date: 2026-06-02

Status: implemented

Scope:

- Changed default unresolved sorting, filtering, risk, analyze summary, file-hotspots, graph-health, MCP context resource, and MCP target payloads to use default-visible unresolved.
- Kept raw unresolved diagnostics available through `raw-unresolved` sorting and preserved explicit resolution inventory as raw.
- Added `production-unresolved` and `test-unresolved` sort modes for targeted investigation.

Source / command evidence:

| Check | Result |
|---|---|
| `anvien analyze --force` | Pass before P1-C work. Graph had 818 files scanned, 598 parsed code files, 0 failed parses, 96,521 nodes, 132,071 relationships, 590 raw unresolved files, and default hotspots still test-dominated before ranking change. |
| `anvien query files "file hotspot unresolved ranking graph health file context" --repo Anvien` | Confirmed owners in file projection, file-hotspots, graph-health, analyze projection, MCP resources, and resolution inventory. |
| `anvien query files "unresolvedSourceSiteCount sort unresolved FileSummary graph health hotspots" --repo Anvien` | Confirmed P1-B bucket fields existed, but callers still used raw/default legacy ranking before P1-C. |
| Source inspection | Confirmed `resolution_inventory_command.go` is explicit raw diagnostic output and must not inherit default-visible filtering. |

Impact / blast radius:

| Target | Result |
|---|---|
| `internal/filecontext/context.go` | HIGH/CRITICAL file projection surface; affects list/detail summaries, sorting, filtering, risk, CLI/API/MCP/Web consumers. |
| `internal/cli/command.go` | CRITICAL analyze output surface; changed fileProjection summary and hotspot lines. |
| `internal/cli/file_context_command.go` | CRITICAL command surface; changed default unresolved display and added raw column/sort modes. |
| `internal/cli/graph_health_command.go` | CRITICAL graph-health surface; changed file-layer unresolved summary and hotspot rendering. |
| `internal/cli/resolution_inventory_command.go` | CRITICAL diagnostic surface; intentionally kept explicit inventory raw via `raw-unresolved`. |
| `internal/mcp/resources.go` and `internal/mcp/target_dispatch.go` | CRITICAL MCP resource/target payload surface; default `unresolved` now means default-visible and raw fields remain explicit. |

Implementation evidence:

| File | Evidence |
|---|---|
| `internal/filecontext/context.go` | `risk` now follows `defaultVisibleRisk`; default `unresolved` sort/filter uses `defaultVisibleUnresolvedSourceSiteCount`; added `raw-unresolved`, `production-unresolved`, and `test-unresolved` sort/filter modes. |
| `internal/cli/command.go` | Analyze file projection now reports `unresolvedFiles` as default-visible and adds `rawUnresolvedFiles` plus `defaultVisibleUnresolvedFiles`; hotspot lines include raw unresolved separately. |
| `internal/cli/file_context_command.go` | File hotspot output now displays default unresolved and raw unresolved separately; help lists raw/production/test unresolved sort modes. |
| `internal/cli/graph_health_command.go` | Graph-health file output now separates default-visible and raw unresolved file counts and hotspot counts. |
| `internal/cli/resolution_inventory_command.go` | Resolution inventory now explicitly requests `raw-unresolved` so diagnostic inventory still includes test/raw unresolved. |
| `internal/mcp/resources.go` | Repo context resource now reports `unresolved_files`, `raw_unresolved_files`, and `default_visible_unresolved_files`; top hotspots use default-visible unresolved plus raw side channel. |
| `internal/mcp/target_dispatch.go` | MCP selected-file relationship hints now expose default unresolved, raw unresolved, bucket counts, raw risk, and default-visible risk. |
| `internal/filecontext/context_test.go` | Added ranking/filter assertions proving default unresolved excludes test-file unresolved while raw/test/production sort modes expose the correct bucket. |
| `internal/cli/file_context_command_test.go` and `internal/httpapi/file_context_test.go` | Added fixtures with test-file unresolved; default unresolved-only hides the test file, while `test-unresolved` returns it with raw/test counts and low risk. |
| Plan | Marked P1-C complete. |
| Benchmark ledger | Added P1-C ranking and inventory metrics. |

Validation:

| Command | Result |
|---|---|
| `go test ./internal/filecontext ./internal/cli ./internal/httpapi ./internal/mcp` | Pass. |
| `powershell -ExecutionPolicy Bypass -File anvien-launcher\build.ps1` | Pass; Vite emitted existing chunk-size/dynamic-import warnings. |
| `go run ./cmd/generate-web-contracts --check` | Pass. |
| `go test ./cmd/... ./internal/filecontext ./internal/cli ./internal/httpapi ./internal/mcp ./internal/contracts` | Pass. |
| `go test ./...` | Fails for existing repo-wide fixture/baseline issues outside this slice: invalid fixture packages under `anvien/test/fixtures`, C source fixture without cgo, deliberate type/pointer assertion failures, and missing `baseline/phase-1-contract-freeze/ladybugdb-graph-contract.json` for `internal/lbugschema`. Relevant package tests above pass. |
| `anvien analyze --force` | Pass after implementation. Graph had 818 files scanned, 598 parsed code files, 0 failed parses, 96,594 nodes, 132,208 relationships, 15,929 dependency edges, 590 raw unresolved files, and 335 default-visible unresolved files. Default top 5 hotspots are no longer test/e2e files. |
| `anvien file-hotspots --repo Anvien --sort unresolved --limit 5` | Pass. Top 5 default hotspots are `useAppState.local-runtime.tsx`, `GraphCanvas.tsx`, `internal/contracts/web_ui.go`, `FileTreePanel.tsx`, and `useSigma.ts`; none are test/e2e files. |
| `anvien file-hotspots --repo Anvien --sort raw-unresolved --limit 5` | Pass. Raw view still exposes the previous test/e2e hotspots with raw counts 1445, 1121, 1052, 934, and 856; default unresolved for those rows is 0 and risk is low. |
| `anvien file-hotspots --repo Anvien --sort test-unresolved --unresolved-only --limit 5` | Pass. Explicit test view reports 238 test-unresolved files and the same top raw test/e2e files. |
| `anvien graph-health files --repo Anvien --sort unresolved --limit 5 --json` | Pass. JSON rows expose raw, production, test, non-actionable, unknown, and default-visible counts; default rows are non-test. |

Failures / handling:

- `go test ./...` remains unsuitable as a repo-wide validation command because it includes intentionally invalid fixture packages and a missing frozen baseline file. This slice records the failure and uses full build plus affected package tests as validation.
- No Web UI source was changed in P1-C, so browser/e2e validation remains for P2/P3-B.

Detect changes:

| Command | Result |
|---|---|
| `anvien detect-changes --repo Anvien --scope all` | Pass. Summary risk `critical`; changed files: 13; affected files: 13; changed count: 216; affected count: 36; changed app layers: `api`, `api_test`, `backend`, `backend_test`, `docs`; affected app layers: `api`, `backend`, `mixed`; changed functional areas: `api`, `cli`, `documentation`, `mcp`, `unknown`; affected functional areas: `cli`, `mcp`, `mixed`, `unknown`; changed ResolutionGap entities: 157. Critical blast radius is expected because the slice changes central file projection, CLI, graph-health, MCP payloads, and tests/docs. |

Commit:

- `6c0a7f7 feat: rank unresolved hotspots by default-visible bucket`

## E5 - P2-A Web File Map And Graph Default Display

Date: 2026-06-02

Status: implemented

Scope:

- Updated Web file map and file detail defaults to use backend-provided `kind`, `defaultVisibleUnresolvedSourceSiteCount`, and risk fields instead of raw unresolved.
- Rendered test file identity as `Test File` in file map/detail.
- Prevented raw test unresolved samples from expanding in file detail by default.
- Updated graph semantic defaults so test/non-actionable ResolutionGap nodes are hidden by default while test file nodes remain visible.

Source / command evidence:

| Check | Result |
|---|---|
| `anvien analyze --force` | Pass before P2-A inspection and again after implementation. Latest graph had 819 files scanned, 599 parsed code files, 0 failed parses, 96,771 nodes, 132,391 relationships, 591 raw unresolved files, and 335 default-visible unresolved files. |
| `anvien query files "Web file map unresolved defaultVisibleUnresolved file detail graph test file" --repo Anvien` | Confirmed Web surfaces still had raw/test unresolved evidence and file summaries expose `kind=test`, raw, test, and default-visible bucket fields. |
| `anvien query files "FileMapPanel FileDetailPanel unresolvedSourceSiteCount defaultVisibleRisk" --repo Anvien` | Confirmed `FileMapPanel` and `FileDetailPanel` were direct owners for raw unresolved display. |
| `anvien query files "Web graph ResolutionGap node filter test unresolved default visibility" --repo Anvien` | Confirmed graph default visibility is governed by semantic filters and graph filtering, not by file map code. |
| Source inspection | `FileMapPanel` used `unresolvedSourceSiteCount` for totals, warning icon, and `Unres` column; `FileDetailPanel` used raw summary and raw unresolved groups; `semantic-filters.ts` defaulted ResolutionGap actionability/source app-layer filters to all values. |

Impact / blast radius:

| Target | Result |
|---|---|
| `anvien impact file anvien-web/src/components/FileMapPanel.tsx --repo Anvien --direction upstream` | HIGH blast radius. Affected files include `App.tsx`, `FileTreePanel.tsx`, `main.tsx`, and the component tests. |
| `anvien impact file anvien-web/src/components/FileDetailPanel.tsx --repo Anvien --direction upstream` | HIGH blast radius. Affected files include `App.tsx`, `CodeReferencesPanel.tsx`, `main.tsx`, and the component tests. |

Implementation evidence:

| File | Evidence |
|---|---|
| `anvien-web/src/components/FileMapPanel.tsx` | Default unresolved totals, warning icon, and `Unres` column now use `defaultVisibleUnresolvedSourceSiteCount`; test rows show a `Test File` badge; raw/test counts remain only in row metadata title. |
| `anvien-web/src/components/FileDetailPanel.tsx` | Summary unresolved count uses default-visible unresolved; test files show `Kind: Test File`; test-file raw unresolved call/ref/import quality pills and raw unresolved sample groups are not rendered by default. |
| `anvien-web/src/lib/semantic-filters.ts` | Default ResolutionGap semantic filters exclude `non_actionable` gaps and `*_test` source app layers; normal test file nodes stay visible because app-layer visibility still includes test layers. |
| `anvien-web/test/unit/FileMapPanel.test.tsx` | Fixture now includes raw/test/default-visible bucket fields and asserts `Test File` identity while default unresolved totals ignore test raw unresolved. |
| `anvien-web/test/unit/FileDetailPanel.test.tsx` | Added test-file fixture proving raw test unresolved sample text is not rendered by default while tested-target relationship remains visible. |
| `anvien-web/test/unit/semantic-filters.test.ts` | Added graph filter test proving test file nodes stay visible, source analyzer gaps stay visible, and test/non-actionable ResolutionGap nodes are hidden by default but can be re-enabled explicitly. |
| `anvien-web/e2e/file-map-test-unresolved.spec.ts` | Added mocked Playwright e2e for file map/detail default behavior with source and test file fixtures. |
| Plan | Marked P2-A complete. |
| Benchmark ledger | Added P2-A UI/default-visibility metrics. |

Validation:

| Command | Result |
|---|---|
| `npm --prefix anvien-web test -- FileMapPanel FileDetailPanel` | Pass before graph-filter change; 2 files, 6 tests. |
| `npm --prefix anvien-web test -- semantic-filters FileMapPanel FileDetailPanel` | Pass after graph-filter change; 3 files, 10 tests. |
| `powershell -ExecutionPolicy Bypass -File anvien-launcher\build.ps1` | Pass; Vite emitted existing chunk-size/dynamic-import warnings. |
| `npm --prefix anvien-web test` | Pass; 52 files, 410 tests. |
| `npm --prefix anvien-web run test:e2e -- file-map-test-unresolved.spec.ts` | Pass; 1 Chromium e2e. |

Failures / handling:

- First e2e attempt failed because `page.getByText("src/app.test.ts").click()` resolved a hidden text node; selector was scoped to the visible `file-map-row` and the spec passed.
- Vite dev server was started at `http://127.0.0.1:5228` for e2e validation.

Detect changes:

| Command | Result |
|---|---|
| `anvien detect-changes --repo Anvien --scope all` | Pass after final graph refresh. Scope contained 9 changed/affected files across Web graph UI, frontend tests, and ledger files; summary risk was low, file-layer changed risk was high from the touched Web UI files, and no affected processes were reported. |

Commit:

- `f7de70d feat: hide test unresolved in web defaults`

## E6 - P2-B Explicit Test Unresolved Drill-Down

Date: 2026-06-02

Status: implemented

Scope:

- Added explicit Web access to raw/test unresolved without changing default production-focused visibility.
- Added file map sort choices for raw unresolved and test unresolved.
- Added a file detail raw unresolved toggle that is off by default and resets when the selected file changes.
- Exposed unresolved sample `sourceSiteId` in file detail sample rows for traceability.

Source / command evidence:

| Check | Result |
|---|---|
| `anvien analyze --force` | Pass after P2-A commits and again after P2-B implementation. Latest graph had 819 files scanned, 599 parsed code files, 0 failed parses, 96,857 nodes, 132,459 relationships, 591 raw unresolved files, and 335 default-visible unresolved files. |
| Source inspection | `FileMapPanel` already passed sort strings directly to `/api/file-hotspots`; backend already supports `raw-unresolved` and `test-unresolved`. `FileDetailPanel` had raw unresolved groups in the response but rendered only default-visible groups after P2-A. |

Impact / blast radius:

| Target | Result |
|---|---|
| `anvien impact file anvien-web/src/components/FileMapPanel.tsx --repo Anvien --direction upstream` | HIGH blast radius. Affected files include `App.tsx`, `FileTreePanel.tsx`, `main.tsx`, and `FileMapPanel` tests; no affected processes were reported. |
| `anvien impact file anvien-web/src/components/FileDetailPanel.tsx --repo Anvien --direction upstream` | HIGH blast radius. Affected files include `App.tsx`, `CodeReferencesPanel.tsx`, `main.tsx`, and `FileDetailPanel` tests; no affected processes were reported. |

Implementation evidence:

| File | Evidence |
|---|---|
| `anvien-web/src/components/FileMapPanel.tsx` | Added explicit `raw-unresolved` and `test-unresolved` sort options; default sort remains `unresolved`. |
| `anvien-web/src/components/FileDetailPanel.tsx` | Added a default-off `Raw` toggle for raw unresolved groups; raw test quality counts and raw samples only render after the toggle is pressed; sample rows now show `sourceSiteId` when present. |
| `anvien-web/test/unit/FileMapPanel.test.tsx` | Asserts selecting `test-unresolved` passes that exact sort to the backend request. |
| `anvien-web/test/unit/FileDetailPanel.test.tsx` | Asserts test-file raw sample and source-site ID are hidden by default, then visible after pressing the raw unresolved toggle. |
| `anvien-web/e2e/file-map-test-unresolved.spec.ts` | Extended mocked browser flow to select `test-unresolved`, open a test file, verify default hiding, then press Raw and verify raw sample and `sourceSiteId`. |
| Plan | Marked P2-B complete. |
| Benchmark ledger | Added P2-B explicit access metrics. |

Validation:

| Command | Result |
|---|---|
| `npm --prefix anvien-web test -- FileMapPanel FileDetailPanel` | Pass as a pre-build smoke check; 2 files, 6 tests. |
| `powershell -ExecutionPolicy Bypass -File anvien-launcher\build.ps1` | Pass; Vite emitted existing chunk-size/dynamic-import warnings. |
| `npm --prefix anvien-web test -- FileMapPanel FileDetailPanel` | Pass after full build; 2 files, 6 tests. |
| `npm --prefix anvien-web run test:e2e -- file-map-test-unresolved.spec.ts` | Pass after full build; 1 Chromium e2e. |

Failures / handling:

- No failure after final P2-B implementation. The first unit run was treated as a smoke check because the repository rule requires the formal validation sequence to run after the full build.

Detect changes:

| Command | Result |
|---|---|
| `anvien detect-changes --repo Anvien --scope all` | Pass after final graph refresh. Scope contained 8 changed/affected files across Web graph UI, frontend tests, and ledger files; summary risk was low, file-layer changed risk was low, and no affected processes were reported. |

Commit:

- Pending until detect-changes passes.
