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

- Pending until detect-changes passes.
