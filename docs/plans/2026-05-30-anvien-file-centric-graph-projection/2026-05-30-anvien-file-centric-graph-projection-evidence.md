# Anvien File-Centric Graph Projection Evidence Ledger

Date: 2026-05-30

Status: Completed

Companion files:

- Plan: [2026-05-30-anvien-file-centric-graph-projection-plan.md](2026-05-30-anvien-file-centric-graph-projection-plan.md)
- Benchmark ledger: [2026-05-30-anvien-file-centric-graph-projection-benchmark.md](2026-05-30-anvien-file-centric-graph-projection-benchmark.md)

## Evidence Rules

1. Record facts that explain why each task is correct.
2. Keep benchmark tables in the benchmark ledger, not here.
3. For code changes, record impact/blast-radius before edits.
4. For graph-based validation, record the graph refresh command and graph inventory summary.
5. For API or contract changes, record route/tool/shape impact and contract regeneration evidence.
6. For Web UI changes, record full build, unit tests, e2e tests, and any screenshot/browser validation if used.
7. Record failures and the fix or decision that handled them.
8. Record `anvien detect-changes --repo Anvien --scope all` before each implementation commit.
9. Record commit hashes as closure evidence.

## Evidence Template

Use this template for each implementation slice:

```text
## E<n> - <Phase/Task Title>

Date:

Status:

Scope:

- ...

Impact / blast radius:

| Command | Result |
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

## E0 - User Problem And Direction

Date: 2026-05-30

Status: recorded

User direction:

- Keep the current symbol-centric graph model as the source of truth.
- Add a file-centric projection layer so users can inspect graph facts from a file-first perspective.
- The desired view is:

```text
File
  -> summary
  -> symbol tree
  -> relationships
  -> unresolved source sites
  -> linked flows/routes/tools/tests
  -> quality signals
```

Problem evidence from discussion:

- Current symbol graph is strong for exact symbol context, impact, rename, detect-changes, and source-site proof.
- Current inspection is weaker when the user starts from a file and asks what it contains, who depends on it, what it depends on, where unresolved sites are, and which flows/tests touch it.
- The proposed solution is a projection derived from existing graph facts, not a replacement for symbol-level graph ownership.

Planning evidence:

| Check | Result |
|---|---|
| Plan file naming | `2026-05-30-anvien-file-centric-graph-projection-plan.md` uses ISO date and lowercase kebab-case slug. |
| Evidence file naming | `2026-05-30-anvien-file-centric-graph-projection-evidence.md` shares the same slug. |
| Benchmark file naming | `2026-05-30-anvien-file-centric-graph-projection-benchmark.md` shares the same slug. |
| Doc-only planning rule | No Anvien graph command is required for creating this initial doc-only plan set. |

## E1 - Baseline Graph Schema Discovery

Date: 2026-05-30

Status: completed

Readiness review evidence:

| Check | Result |
|---|---|
| Graph refresh | `anvien analyze --force --name Anvien` completed. |
| Graph inventory | `files: scanned=819 parsed=584 unsupported=235 failed=0`; `nodes=91586 relationships=125053`; graph path `.anvien/graph.json`. |
| CLI ownership inspected | Existing command owners include `internal/cli/command.go`, `internal/cli/tool_command.go`, `internal/cli/api_command.go`, and graph-quality command files. |
| MCP ownership inspected | Existing tool owners include `internal/mcp/server.go`, `internal/mcp/context.go`, `internal/mcp/impact.go`, `internal/mcp/tools.go`, `internal/mcp/route_tool_map.go`, and `internal/mcp/route_shape_impact.go`. |
| Graph facts inspected | `internal/graph/types.go` already carries file path and source-site fields on graph nodes/relationships; graph-health inputs already include file/source-site metadata. |
| Web contract ownership inspected | Web contract source is owned by `internal/contracts/web_ui.go` and generated through `cmd/generate-web-contracts`; generated TypeScript lives under `anvien-web/src/generated`. |
| AI context ownership inspected | Generated guidance is owned by `internal/aicontext/aicontext.go` and embedded skill source files under `internal/aicontext/skills`. |

P0-A graph refresh:

| Command | Result |
|---|---|
| `.\anvien\bin\anvien.exe analyze --force --name Anvien` | Pass. `files: scanned=819 parsed=584 unsupported=235 failed=0`; `nodes=91587 relationships=125054`; graph path `.anvien/graph.json`. |
| `.\anvien\bin\anvien.exe graph-health summary --repo Anvien --json` | Pass. Indexed commit and current commit both `cdbd4af19b867b1ed4a3efc2d6c9779f25907ce3`; `resolutionGapNodeCount=65652`; `hasResolutionGapRelationshipCount=65652`; `sourceBackedUnresolvedReferenceCount=66555`; `unattributedUnresolvedReferenceCount=0`. |

P0-A graph facts:

| Fact | Evidence |
|---|---|
| File nodes exist | Graph contains `819` `File` nodes, all with `filePath`. |
| File classification exists | File nodes include `appLayer`, `functionalArea`, language, extension, document kind, and binary metadata where available. |
| File-to-symbol ownership exists | Graph contains `21334` `DEFINES` relationships from `File` nodes to symbol-like nodes. |
| Symbol containment exists | Graph contains `2784` `CONTAINS` relationships; `143` from file nodes and `2641` from non-file nodes for nesting/ownership. |
| Source-site trace fields exist | `83143` relationships carry `sourceSiteId`, `sourceSiteIds`, and `filePath`; distinct observed source-site ids: `95433`. |
| ResolutionGap trace fields exist | `65652` `ResolutionGap` nodes carry `sourceSiteId` and `filePath`. |
| Unresolved grouping by file is derivable | `576` files have unresolved source-site evidence through `ResolutionGap` file paths. |
| Relationship types are sufficient for first projection | Existing relationship types include `DEFINES`, `CONTAINS`, `CALLS`, `USES`, `IMPORTS`, `ACCESSES`, `MEMBER_OF`, `HAS_PROPERTY`, `HAS_METHOD`, `STEP_IN_PROCESS`, `ENTRY_POINT_OF`, and `HAS_RESOLUTION_GAP`. |
| Command surface owners are identifiable | CLI parent commands exist for `query`, `context`, `impact`, `detect-changes`, `graph-health`, `api`, and `group`; API and graph-health already use child command patterns. |

Plan additions from review:

- Add shared projection service/package as an explicit ownership boundary.
- Add shared target resolver for parent/child command dispatch and ambiguity handling.
- Add projection cache/index invalidation tied to graph freshness.
- Add exact Web/API route naming and generated-contract validation gates.
- Add MCP surface snapshot/tool schema validation gates.

Remaining implementation evidence:

- Existing File/Symbol/SourceSite/ResolutionGap/Flow/API/MCP/test graph facts.
- Current schema facts that support `File -> Symbol`, source-site ownership, symbol nesting, and relationship traceability.
- Missing facts that require implementation.
- Baseline graph inventory summary recorded in benchmark ledger.

## E2 - File Context Contract

Date: 2026-05-30

Status: completed

Contract evidence:

| Contract area | Result |
|---|---|
| Envelope | Added `File Context JSON Contract V0` to the plan. |
| Required top-level fields | `repo`, `repoPath`, `graph`, `target`, `summary`, `symbolTree`, `relationships`, `unresolved`, `linked`, `quality`, and `limits`. |
| Target dispatch fields | `type`, `input`, `normalizedPath`, `dispatchMode`, and `ambiguityCandidates`. |
| Summary fields | Path, language, kind, app layer, functional area, parse status, symbol counts, relationship counts, unresolved count, linked counts, and risk. |
| Relationship shape | `local`, `outboundByFile`, `inboundByFile`, total counts, samples, and trace fields. |
| Unresolved shape | Total, grouped counts, line/column, target text, source symbol, gap kind, classification, actionability, proof kind, source-site id, and source-site status. |
| Linked overlays | Flows, routes, MCP tools, and tests with source/confidence/trace metadata. |
| Quality shape | Parser, resolution confidence, unresolved counts, generated/stale/changed-since-analyze flags. |
| Sample limits | Relationship, unresolved, and linked samples have explicit limits; total counts must not be truncated. |
| Source rules | Contract documents field derivation from `File`, `DEFINES`, `CONTAINS`, symbol relationships, `ResolutionGap`, graph-health, process/route/tool/test facts, and git freshness data. |
| Compatibility | Contract is structured for CLI JSON, API, MCP, and Web; human output may summarize the same shape without changing counts. |

## E3 - Projection Builder

Date: 2026-05-30

Status: completed

Scope:

- Add the first shared backend file-context projection package.
- Do not connect CLI, MCP, API, or Web surfaces yet; those are later phases.
- Do not edit existing graph core storage or analyzer logic.

Impact / blast radius:

| Command | Result |
|---|---|
| `.\anvien\bin\anvien.exe impact "Graph" --repo Anvien --direction upstream` | Returned a broad ambiguous candidate set for `Graph`; treated as HIGH blast-radius signal for graph model changes. No existing `internal/graph` symbol was edited in this slice. |
| `.\anvien\bin\anvien.exe impact "Struct:internal/graph/types.go:Graph" --repo Anvien --direction upstream` | Target UID was not resolved by the current impact command. Scope stayed additive in a new package. |

Implementation evidence:

| File | Evidence |
|---|---|
| `internal/filecontext/context.go` | New shared `filecontext` package with `FileContext`, `FileSummary`, `SymbolTreeNode`, relationship grouping, unresolved grouping, linked/quality/limit structs, `Builder`, `NewBuilder`, and `BuildFileContext`. |
| `internal/filecontext/context.go` | Builder derives file context from existing `graph.Graph` facts: `File` node properties, `DEFINES`, `CONTAINS`, canonical symbol relationships, and `ResolutionGap` nodes with source-site metadata. |
| `internal/filecontext/context.go` | File-level local/outbound/inbound relationship groups preserve total counts while limiting samples. |
| `internal/filecontext/context.go` | Unresolved groups preserve total counts, classification/actionability/kind counts, line/column, target text, proof kind, source-site id, and source-site status. |
| `internal/filecontext/context_test.go` | Fixture tests cover path normalization, file summary, symbol tree nesting, relationship grouping, unresolved grouping, sample limits, missing-file behavior, and deterministic output across relationship insertion order. |

Traceability evidence:

| Projection field | Source fact |
|---|---|
| File summary path/language/layer/area | `File` node properties. |
| Symbol tree | `DEFINES` from File to symbol and `CONTAINS` between symbols. |
| Relationship groups | Canonical relationships such as `CALLS`, with source/target node file paths and source-site fields. |
| Unresolved groups | `ResolutionGap` node properties including `sourceNodeId`, `sourceSiteId`, `sourceSiteStatus`, `targetText`, `classification`, and `actionability`. |

Validation:

| Command | Result |
|---|---|
| `powershell -ExecutionPolicy Bypass -File anvien-launcher\build.ps1` | Pass. Existing Vite dynamic-import and chunk-size warnings only. |
| `go test ./internal/filecontext -count=1` | Pass; `1.052s` after full build gate. |
| `go test ./cmd/... ./internal/... -count=1` | Pass; all cmd/internal packages passed, including new `internal/filecontext`. |
| `.\anvien\bin\anvien.exe analyze --force --name Anvien` | Pass after P1-A code. `files: scanned=821 parsed=586 unsupported=235 failed=0`; `nodes=92352 relationships=126299`. |
| `.\anvien\bin\anvien.exe detect-changes --repo Anvien --scope all` | CRITICAL risk reported for staged P1-A changes: `changed_files=5`, `changed_count=789`, `affected_count=17`; changed layers were `backend`, `backend_test`, and `docs`. |
| `.\anvien\bin\anvien.exe context "BuildFileContext" --repo Anvien` | Found `Method:internal/filecontext/context.go:Builder.BuildFileContext#2`. Incoming edge is `HAS_METHOD` from `Builder`; outgoing edges are same-package projection helpers/structs. |

Failures / handling:

- Initial targeted `go test ./internal/filecontext` was run before the full build gate. The full build gate was then run, and the focused test plus full cmd/internal test batch were rerun after the gate.
- `detect-changes` reported CRITICAL because this slice adds a new backend package with many new symbols and source-site gaps. The affected processes are generated around the new `internal/filecontext` package; the package is not wired into CLI/MCP/API/Web runtime behavior yet. Scope remains additive and tests passed.

Detect changes:

| Command | Result |
|---|---|
| `.\anvien\bin\anvien.exe detect-changes --repo Anvien --scope all` | `risk_level=critical`; `changed_files=5`; `changed_count=789`; `affected_count=17`; affected app layer `backend`; affected functional area `unknown`; no existing runtime command surface was changed in this slice. |

## E4 - File Hotspots And Aggregation

Date: 2026-05-30

Status: completed

Scope:

- Add repo-wide file list/hotspot aggregation to the shared `internal/filecontext` builder.
- Do not add CLI/API/MCP/Web command surfaces yet.
- Do not implement projection cache yet; cache and invalidation remain P1-C.

Impact / blast radius:

| Command | Result |
|---|---|
| `.\anvien\bin\anvien.exe analyze --force --name Anvien` | Pass before P1-B graph-based checks. `files: scanned=821 parsed=586 unsupported=235 failed=0`; `nodes=92352 relationships=126299`. |
| `.\anvien\bin\anvien.exe impact "BuildFileContext" --repo Anvien --direction upstream` | LOW. `impactedCount=0`, no affected modules or processes. |
| `.\anvien\bin\anvien.exe context "Builder" --repo Anvien` | Ambiguous across many builder symbols; exact P1-B edits stayed inside `internal/filecontext`. |

Implementation evidence:

| File | Evidence |
|---|---|
| `internal/filecontext/context.go` | Added `FileListOptions`, `FileList`, and `Builder.BuildFileList`. |
| `internal/filecontext/context.go` | `BuildFileList` builds summaries from all `File` nodes and uses one relationship pass to aggregate local, inbound, and outbound counts. |
| `internal/filecontext/context.go` | Added sort modes: `path`, `unresolved`, `fan-in`, `fan-out`, `symbols`, `flows`, and `tests`. |
| `internal/filecontext/context.go` | Added filters for kind, app layer, functional area, API-related files, unresolved-only, high fan-in, and high fan-out. |
| `internal/filecontext/context.go` | Added limit/offset pagination. |
| `internal/filecontext/context_test.go` | Added tests for sorting, filtering, pagination, high fan-in, and high fan-out behavior. |
| `internal/filecontext/context_test.go` | Added `BenchmarkBuildFileListCurrentScale` using `821` files and `126000` relationships. |

Validation:

| Command | Result |
|---|---|
| `powershell -ExecutionPolicy Bypass -File anvien-launcher\build.ps1` | Pass after P1-B code. Existing Vite dynamic-import and chunk-size warnings only. |
| `go test ./internal/filecontext -count=1` | Pass; `1.578s`. |
| `go test ./internal/filecontext -run '^$' -bench BenchmarkBuildFileListCurrentScale -benchmem -count=3` | Pass. Benchmark runs: `233345033 ns/op`, `106833581 ns/op`, `123577177 ns/op`; `490000 B/op`, `831 allocs/op`. |
| `go test ./cmd/... ./internal/... -count=1` | Pass; all cmd/internal packages passed. |
| `.\anvien\bin\anvien.exe analyze --force --name Anvien` | Pass after P1-B code. `files: scanned=821 parsed=586 unsupported=235 failed=0`; `nodes=92559 relationships=126624`. |
| `.\anvien\bin\anvien.exe detect-changes --repo Anvien --scope all` | MEDIUM risk reported for staged P1-B changes: `changed_files=5`, `changed_count=249`, `affected_count=2`; changed layers were `backend`, `backend_test`, and `docs`. |

Benchmark link:

- See B5 for file list aggregation timing and allocation metrics.

Remaining:

- Projection cache behavior remains pending for P1-C.

Detect changes:

| Command | Result |
|---|---|
| `.\anvien\bin\anvien.exe detect-changes --repo Anvien --scope all` | `risk_level=medium`; `changed_files=5`; `changed_count=249`; `affected_count=2`; affected app layer `backend`; affected functional area `unknown`. |

## E4A - Projection Cache, Index Reuse, And Invalidation

Date: 2026-05-30

Status: completed

Scope:

- Add reusable cache behavior for the shared file-context projection builder.
- Keep the cache inside `internal/filecontext`; do not wire CLI/API/MCP/Web surfaces yet.
- Preserve graph freshness by keying cache entries by repo identity, repo path, graph path, and graph hash/fingerprint.

Impact / blast radius:

| Command | Result |
|---|---|
| `.\anvien\bin\anvien.exe analyze --force --name Anvien` | Pass before P1-C graph-based checks. `files: scanned=821 parsed=586 unsupported=235 failed=0`; `nodes=92559 relationships=126624`. |
| `.\anvien\bin\anvien.exe impact "NewBuilder" --repo Anvien --direction upstream` | LOW. `impactedCount=0`; target `Function:internal/filecontext/context.go:NewBuilder#1`. |
| `.\anvien\bin\anvien.exe impact "BuildFileList" --repo Anvien --direction upstream` | LOW. `impactedCount=0`; target `Method:internal/filecontext/context.go:Builder.BuildFileList#1`. |

Implementation evidence:

| File | Evidence |
|---|---|
| `internal/filecontext/context.go` | Added `CacheKey` with repo, repo path, graph path, and graph hash fields. |
| `internal/filecontext/context.go` | Added `BuilderCache` with mutex-protected `Get`, `Invalidate`, `Clear`, and `Len` methods. |
| `internal/filecontext/context.go` | `Get` returns a warm cached `Builder` when key fields and graph freshness match; otherwise it creates and stores a new builder. |
| `internal/filecontext/context.go` | Added `GraphFingerprint` fallback for callers that do not yet supply an explicit graph hash. |
| `internal/filecontext/context.go` | `Invalidate` can remove exact entries or all entries matching repo/repo path/graph path so analyze refreshes and repo switches cannot reuse stale projection state. |
| `internal/filecontext/context_test.go` | Added tests for cold miss, warm hit, graph-change miss, explicit invalidation, repo isolation, and explicit graph-hash isolation. |
| `internal/filecontext/context_test.go` | Added benchmarks for file-list aggregation, cache warm hit, and cache cold build. |

Validation:

| Command | Result |
|---|---|
| `go test ./internal/filecontext -count=1` | Pass before full build gate; `1.092s`. |
| `powershell -ExecutionPolicy Bypass -File anvien-launcher\build.ps1` | Pass after P1-C code. Existing Vite dynamic-import and chunk-size warnings only. |
| `go test ./internal/filecontext -count=1` | Pass after full build gate; `1.363s`. |
| `go test ./internal/filecontext -run '^$' -bench 'Benchmark(BuildFileListCurrentScale|BuilderCacheHit|BuilderCacheColdBuild)$' -benchmem -count=3` | Pass. File-list median `131.3 ms/op`; cache warm-hit median `305.3 ns/op`, `0 B/op`, `0 allocs/op`; cache cold-build median `8.0 ms/op`, about `0.57 MB/op`. |
| `go test ./cmd/... ./internal/... -count=1` | Pass; all cmd/internal packages passed. |
| `.\anvien\bin\anvien.exe analyze --force --name Anvien --benchmark-label file-centric-p1c-final --benchmark-json .tmp\file-centric-p1c-final-analyze-benchmark.json` | Pass after P1-C code and ledger updates. `files: scanned=821 parsed=586 unsupported=235 failed=0`; `nodes=92652 relationships=126821`. |
| `.\anvien\bin\anvien.exe graph-health summary --repo Anvien --json` | Pass after P1-C code. `sourceBackedUnresolvedReferenceCount=67223`; `resolutionGapNodeCount=66322`; `resolvedReferenceCount=30440`. |

Graph generation benchmark:

| Command | Result |
|---|---|
| `.\anvien\bin\anvien.exe analyze --force --name Anvien --benchmark-label file-centric-p1c-final --benchmark-json .tmp\file-centric-p1c-final-analyze-benchmark.json` | Pass. Total graph generation time `34839.3 ms`; `821` scanned files; `92652` nodes; `126821` relationships. |

Benchmark link:

- See B0A for graph generation speed.
- See B5 for projection aggregation and cache performance.

Failures / handling:

- No code validation failure was observed in this slice.
- Graph generation speed was missing from the ledger during the P1-C review. It was added with `anvien analyze --benchmark-json` before closure.

Detect changes:

| Command | Result |
|---|---|
| `.\anvien\bin\anvien.exe detect-changes --repo Anvien --scope all` | `risk_level=low`; `changed_files=5`; `changed_count=114`; `affected_count=0`; changed app layers `backend`, `backend_test`, and `docs`; resolution gap changed entities `66`. |

## E5 - CLI Surface

Date: 2026-05-30

Status: completed

Scope:

- Add repo-agnostic CLI commands for file-centric graph projection.
- Add `anvien file-context <path> --repo <repo>` with compact human output and full `--json`.
- Add `anvien file-hotspots --repo <repo>` with sort/filter/limit options and `--json`.
- Do not add MCP/API/Web surfaces in this slice.

Repo-agnostic invariant:

- Product code must not hardcode the `Anvien` repository name. `Anvien` is only the validation repo for this implementation plan.
- P2-A commands resolve repositories through the normal registry store and `--repo` input using `repo.NewEnvStore()` and `repo.ResolveEntry`.
- CLI tests register a temporary indexed repo named `fixture` and validate against that repo to prove the commands work for arbitrary indexed repositories.

Impact / blast radius:

| Command | Result |
|---|---|
| `.\anvien\bin\anvien.exe analyze --force --name Anvien` | Pass before P2-A graph-based checks. `files: scanned=821 parsed=586 unsupported=235 failed=0`; `nodes=92652 relationships=126821`. |
| `.\anvien\bin\anvien.exe query "CLI command registration cobra root command file context command" --repo Anvien` | Found CLI command owners under `internal/cli/*_command.go`; root registration is in `internal/cli/command.go`; child-command style examples are in `internal/cli/group_command.go` and `internal/cli/api_command.go`. |
| `.\anvien\bin\anvien.exe impact "NewRootCommand" --repo Anvien --direction upstream` | CRITICAL, `impactedCount=1`, direct affected caller `cmd/anvien/main.go`, `processes_affected=11`. This slice makes an additive root command registration and preserves existing commands. |
| `.\anvien\bin\anvien.exe impact "loadGraphHealthGraph" --repo Anvien --direction upstream` | CRITICAL, `impactedCount=6`, `processes_affected=40`. Existing graph-health helper was not edited; P2-A added a separate file-projection graph loader to avoid changing graph-health behavior. |

Implementation evidence:

| File | Evidence |
|---|---|
| `internal/cli/command.go` | Registered `newFileContextCommand()` and `newFileHotspotsCommand()` on the root CLI. |
| `internal/cli/file_context_command.go` | Added `file-context <path>` with repo registry resolution, stale-index guard, graph snapshot loading, shared `filecontext.Builder.BuildFileContext`, compact human output, and full JSON output. |
| `internal/cli/file_context_command.go` | Added `file-hotspots` with sort, limit, offset, kind, app-layer, functional-area, API-only, unresolved-only, high fan-in, and high fan-out filters; implementation uses shared `filecontext.Builder.BuildFileList`. |
| `internal/cli/file_context_command.go` | Default human output is bounded; JSON output preserves the full file context contract. |
| `internal/cli/file_context_command_test.go` | Added tests for JSON and human `file-context`, JSON and human `file-hotspots`, missing-file errors, and repo-agnostic fixture registry behavior. |
| `internal/cli/command_test.go` | Added root/help assertions so `file-context` and `file-hotspots` are discoverable and expose compatibility flags. |

Validation:

| Command | Result |
|---|---|
| `go test ./internal/cli -run 'Test(FileContext|FileHotspots|DirectToolHelp|HelpCommand)' -count=1` | Pass before full build gate; `1.485s`. |
| `powershell -ExecutionPolicy Bypass -File anvien-launcher\build.ps1` | Pass after P2-A code. Existing Vite dynamic-import and chunk-size warnings only. |
| `go test ./internal/cli -count=1` | Pass after full build gate; `50.604s`. |
| `go test ./cmd/... ./internal/... -count=1` | Pass; all cmd/internal packages passed. |
| `.\anvien\bin\anvien.exe file-context internal/cli/command.go --repo Anvien --json` | Pass. Smoke summary: repo `Anvien`, path `internal/cli/command.go`, symbols `51`, outbound `131`, unresolved `168`, symbol tree roots `51`. |
| `.\anvien\bin\anvien.exe file-hotspots --repo Anvien --sort unresolved --limit 3 --json` | Pass. Smoke summary: total files `823`, returned `3`, first hotspot `internal/mcp/server_test.go` with unresolved `1422`. |
| `.\anvien\bin\anvien.exe analyze --force --name Anvien --benchmark-label file-centric-p2a-final --benchmark-json .tmp\file-centric-p2a-final-analyze-benchmark.json` | Pass after P2-A code and ledger updates. `files: scanned=823 parsed=588 unsupported=235 failed=0`; `nodes=93072 relationships=127394`. |
| `.\anvien\bin\anvien.exe graph-health summary --repo Anvien --json` | Pass after P2-A code. `sourceBackedUnresolvedReferenceCount=67517`; `resolutionGapNodeCount=66613`; `resolvedReferenceCount=30625`. |

Graph generation benchmark:

| Command | Result |
|---|---|
| `.\anvien\bin\anvien.exe analyze --force --name Anvien --benchmark-label file-centric-p2a-final --benchmark-json .tmp\file-centric-p2a-final-analyze-benchmark.json` | Pass. Total graph generation time `38342.7 ms`; `823` scanned files; `93072` nodes; `127394` relationships. |

Benchmark link:

- See B0/B1/B2/B3/B4 for graph inventory after CLI surface changes.
- See B6 for CLI response size metrics.
- See B8/B10 for command/test counts.

Failures / handling:

- First response-size check showed unbounded human `file-context` output for a large CLI file: `225` lines and `13096` bytes. Default human output was bounded while keeping `--json` full, reducing the representative human output to `76` lines and `4465` bytes.
- No validation failures were observed after bounding output.

Detect changes:

| Command | Result |
|---|---|
| `.\anvien\bin\anvien.exe detect-changes --repo Anvien --scope all` | `risk_level=critical`; `changed_files=7`; `changed_count=409`; `affected_count=17`; changed app layers `backend`, `backend_test`, and `docs`; affected app layers `backend` and `mixed`; resolution gap changed entities `295`. CRITICAL is expected because root CLI command surface changed; behavior is additive and repo-agnostic. |

## E6 - Web/API Surface

Date: 2026-05-30

Status: completed

Route and contract decisions:

| Decision | Evidence |
|---|---|
| Use `GET /api/file-context` for one file detail | Avoids changing the existing `GET /api/file` source-content endpoint. |
| Use `GET /api/file-hotspots` for file list/hotspot rows | Keeps file projection list separate from raw file reads and graph dump routes. |
| Do not add a separate relationship-expansion endpoint in P2-B | `FileContextResponse.relationships` already carries grouped local/inbound/outbound relationship samples with total counts; expansion can be added later if Web UI needs lazy detail. |
| Keep routes repo-agnostic | Endpoints resolve `repo` through the existing registry resolver; `Anvien` is only the validation repo name. |

API / contract impact:

| Command | Result |
|---|---|
| `.\anvien\bin\anvien.exe analyze --force --name Anvien` | Pass before impact work. `files: scanned=823 parsed=588 unsupported=235 failed=0`; `nodes=93072 relationships=127394`. |
| `.\anvien\bin\anvien.exe impact "NewHandler" --repo Anvien --direction upstream` | `risk=CRITICAL`; `impactedCount=3`; direct `ListenAndServe`; `processes_affected=11`; affected app layers `api`, `backend`. Proceeded because route registration is additive. |
| `.\anvien\bin\anvien.exe impact "WebUIContractTypeScript" --repo Anvien --direction upstream` | `risk=CRITICAL`; `impactedCount=1`; direct `cmd/generate-web-contracts/main.go`; `processes_affected=5`. Proceeded because generated TypeScript contract output is additive. |
| `.\anvien\bin\anvien.exe impact "WebUIContractManifest" --repo Anvien --direction upstream` | `risk=CRITICAL`; `impactedCount=4`; `processes_affected=10`. Proceeded because manifest field is additive. |
| `.\anvien\bin\anvien.exe impact "WebUIContract" --repo Anvien --direction upstream` | `risk=CRITICAL`; `impactedCount=3`; `processes_affected=10`. Proceeded because contract metadata population is additive. |
| `.\anvien\bin\anvien.exe api route-map "/api/file" --repo Anvien --json` | Returned no matching routes even though source has `/api/file`; recorded as HTTP route extraction limitation for current graph. |
| `.\anvien\bin\anvien.exe api route-map "/api/graph" --repo Anvien --json` | Returned no matching routes even though source has `/api/graph`; recorded as HTTP route extraction limitation for current graph. |

Implementation evidence:

| File | Evidence |
|---|---|
| `internal/httpapi/server.go` | Registered `GET /api/file-context` and `GET /api/file-hotspots`; added server-side file projection caches. |
| `internal/httpapi/file_context.go` | Added repo-registry based graph loading, file context response, hotspot/list response, filters, sort/limit/offset handling, stale graph metadata, and projection snapshot cache keyed by repo/path/graph mtime/size/commit. |
| `internal/httpapi/file_context_test.go` | Added API tests for repo-name resolution, file context JSON shape, relationship/unresolved data, hotspot sorting/filtering, and missing-file errors. |
| `internal/contracts/web_ui.go` | Added generated route metadata for file projection API endpoints and TypeScript response interfaces for file context/hotspots. |
| `internal/contracts/web_ui_test.go` | Added assertions for route metadata and generated TypeScript file projection interfaces. |
| `contracts/web-ui/anvien-web-contract.schema.json` | Regenerated from Go contract source; includes `/api/file-context` and `/api/file-hotspots` route metadata. |
| `anvien-web/src/generated/anvien-contracts.ts` | Regenerated from Go contract source; includes `FILE_PROJECTION_API_ROUTES`, `FileContextResponse`, and `FileHotspotsResponse`. |

Validation:

| Command | Result |
|---|---|
| `go run .\cmd\generate-web-contracts` | Pass; regenerated schema and TypeScript contracts from source. |
| `powershell -ExecutionPolicy Bypass -File anvien-launcher\build.ps1` | Pass after P2-B code. Existing Vite dynamic-import and chunk-size warnings only. |
| `go test ./internal/httpapi ./internal/contracts -count=1` | First run hit `TestEmbedEndpointRecoversStaleRepoLock` timing failure; focused rerun passed and package rerun passed. |
| `go test ./internal/httpapi -run "TestFile(Context|Hotspots)|TestFileContextEndpoint|TestFileHotspotsEndpoint" -count=1 -v` | Pass; file projection endpoint tests passed. |
| `go test ./internal/httpapi -run TestEmbedEndpointRecoversStaleRepoLock -count=1 -v` | Pass on rerun; confirms earlier package failure was unrelated timing behavior. |
| `go test ./internal/httpapi ./internal/contracts -count=1` | Pass after rerun; `internal/httpapi` and `internal/contracts` passed. |
| `go run .\cmd\generate-web-contracts --check` | Pass; generated contract output matches source. |
| `go test ./cmd/... ./internal/... -count=1` | Pass; all cmd/internal packages passed after P2-B implementation. |
| HTTP runtime smoke via `.\anvien\bin\anvien.exe serve --host 127.0.0.1 --port <temp>` | Pass. `GET /api/file-context?repo=Anvien&path=internal/httpapi/file_context.go` returned `200`; `GET /api/file-hotspots?repo=Anvien&sort=unresolved&limit=5` returned `200`. |
| `.\anvien\bin\anvien.exe analyze --force --name Anvien --benchmark-label file-centric-p2b-final-cache --benchmark-json .tmp\file-centric-p2b-final-cache-analyze-benchmark.json` | Pass after P2-B cache edit. `files: scanned=825 parsed=590 unsupported=235 failed=0`; `nodes=93434 relationships=127940`. |
| `.\anvien\bin\anvien.exe graph-health summary --repo Anvien --json` | Pass after P2-B. `sourceBackedUnresolvedReferenceCount=67790`; `resolutionGapNodeCount=66883`; `resolvedReferenceCount=30790`. |
| `.\anvien\bin\anvien.exe detect-changes --repo Anvien --scope all` | `risk_level=high`; `changed_files=8`; `changed_count=83`; `affected_count=10`; changed app layers `api`, `api_contract`, `api_test`, and `docs`; affected app layers `api_contract` and `mixed`; resolution gap changed entities `31`. HIGH is expected because API route registration and generated Web contracts changed; behavior is additive and repo-agnostic. |

Runtime benchmark evidence:

| Endpoint | Round | Status | Bytes | Elapsed |
|---|---:|---:|---:|---:|
| `/api/file-context?repo=Anvien&path=internal/httpapi/file_context.go` | 1 | 200 | 57803 | `7017.5 ms` |
| `/api/file-context?repo=Anvien&path=internal/httpapi/file_context.go` | 2 | 200 | 57803 | `107.7 ms` |
| `/api/file-hotspots?repo=Anvien&sort=unresolved&limit=5` | 1 | 200 | 2086 | `97.6 ms` |
| `/api/file-hotspots?repo=Anvien&sort=unresolved&limit=5` | 2 | 200 | 2086 | `93.7 ms` |

Notes:

- The cold file-context request includes graph JSON load and projection cache population.
- The warm request validates the HTTP projection snapshot cache; the second file-context request dropped from `7017.5 ms` to `107.7 ms`.
- The product behavior remains repo-agnostic: tests use fixture repo names `alpha`, `workspace`, and `beta`; runtime smoke uses repo `Anvien` only as the local validation target.

## E7 - Unresolved And Quality Signals

Date: 2026-05-30

Status: completed

Impact:

| Command | Result |
|---|---|
| `.\anvien\bin\anvien.exe analyze --force --name Anvien` | Pass before P3-A impact checks. `files: scanned=825 parsed=590 unsupported=235 failed=0`; `nodes=93434 relationships=127940`. |
| `.\anvien\bin\anvien.exe impact "buildQuality" --repo Anvien --direction upstream` | `risk=HIGH`; `impactedCount=1`; affected `BuildFileContext`; `processes_affected=4`. Proceeded because quality changes are additive and test-covered. |
| `.\anvien\bin\anvien.exe impact "attachFileProjectionMetadata" --repo Anvien --direction upstream` | `risk=CRITICAL`; `impactedCount=3`; affected CLI root flow through `file-context`; `processes_affected=11`. Proceeded because behavior only centralizes metadata attachment. |
| `.\anvien\bin\anvien.exe impact "handleFileContext" --repo Anvien --direction upstream` | `risk=LOW`; `impactedCount=0`. |

Implementation evidence:

| File | Evidence |
|---|---|
| `internal/filecontext/context.go` | Added `AttachMetadata` helper so CLI/API attach repo, graph metadata, `quality.stale`, and `quality.changedSinceAnalyze` consistently. Existing builder already grouped ResolutionGap nodes by file/source symbol and carried kind/classification/actionability/proof/source-site fields. |
| `internal/cli/file_context_command.go` | Uses shared metadata helper for `file-context` JSON/human output. |
| `internal/httpapi/file_context.go` | Uses shared metadata helper for `GET /api/file-context`. |
| `internal/filecontext/context_test.go` | Added coverage for unresolved call/import/type buckets, classification/actionability counts, stale/changed quality propagation, generated file quality, and test file classification. |

Validation:

| Command | Result |
|---|---|
| `powershell -ExecutionPolicy Bypass -File anvien-launcher\build.ps1` | Pass after P3-A code. Existing Vite dynamic-import and chunk-size warnings only. |
| `go test ./internal/filecontext ./internal/cli ./internal/httpapi -count=1` | Pass. |
| `go test ./internal/filecontext -run TestBuildFileContextQualitySignalsAndUnresolvedBuckets -count=1 -v` | Pass; new P3-A focused test passed. |
| `go test ./cmd/... ./internal/... -count=1` | Pass; all cmd/internal packages passed. |
| `.\anvien\bin\anvien.exe detect-changes --repo Anvien --scope all` | `risk_level=medium`; `changed_files=7`; `changed_count=78`; `affected_count=2`. MEDIUM is expected because shared filecontext metadata and CLI/API consumers changed. |

## E8 - Linked Flows, Routes, MCP Tools, And Tests

Date: 2026-05-30

Status: completed

Scope:

- Add linked overlay counts and samples to the shared file context projection.
- Derive flow links from `STEP_IN_PROCESS`, route links from `HANDLES_ROUTE`, MCP tool links from `HANDLES_TOOL`, and test links from test-file relationships into the target file.
- Preserve sample limits while exposing full linked counts.
- Keep the implementation repo-agnostic; `Anvien` is only the local validation repo name.

Impact / blast radius:

| Command | Result |
|---|---|
| `.\anvien\bin\anvien.exe analyze --force --name Anvien` | Pass before P3-B impact checks. `files: scanned=825 parsed=590 unsupported=235 failed=0`; `nodes=93522 relationships=128035`. |
| `.\anvien\bin\anvien.exe impact "BuildFileContext" --repo Anvien --direction upstream` | `risk=LOW`; `impactedCount=0`. |
| `.\anvien\bin\anvien.exe impact "LinkedSummary" --repo Anvien --direction upstream` | `risk=CRITICAL`; `impactedCount=7`; affected app layers `api` and `backend`; `processes_affected=6`. Proceeded because the JSON shape change is additive and generated contract output was regenerated. |
| `.\anvien\bin\anvien.exe impact "WebUIContractTypeScript" --repo Anvien --direction upstream` | `risk=CRITICAL`; `impactedCount=1`; `processes_affected=5`. Proceeded because the TypeScript contract field is additive. |
| `.\anvien\bin\anvien.exe impact "buildFileSummaries" --repo Anvien --direction upstream` | `risk=LOW`; `impactedCount=1`; affected `BuildFileList`. |

Implementation evidence:

| File | Evidence |
|---|---|
| `internal/filecontext/context.go` | Added `LinkedCounts` and populated `LinkedSummary.counts` so truncated samples do not hide total flow/route/tool/test counts. |
| `internal/filecontext/context.go` | `BuildFileContext` now derives linked overlays through shared graph facts: `STEP_IN_PROCESS`, `HANDLES_ROUTE`, `HANDLES_TOOL`, and inbound test-file relationships. |
| `internal/filecontext/context.go` | `BuildFileList` now aggregates linked flow/test counts per file for hotspot and File Map rows. |
| `internal/filecontext/context.go` | Link samples include kind, source relationship type, confidence when present, and a trace string back to source/target graph node ids. |
| `internal/filecontext/context_test.go` | Fixture coverage proves flow, route, MCP tool, and test links are emitted, and that sample limits preserve full counts. |
| `internal/contracts/web_ui.go` | Added `FileLinkedCounts` to generated TypeScript contract source. |
| `internal/contracts/web_ui_test.go` | Added contract generator expectation for `FileLinkedCounts`. |
| `anvien-web/src/generated/anvien-contracts.ts` | Regenerated generated Web TypeScript contract output. |

Validation:

| Command | Result |
|---|---|
| `powershell -ExecutionPolicy Bypass -File anvien-launcher\build.ps1` | Pass after P3-B code. Existing Vite dynamic-import and chunk-size warnings only. |
| `go test ./internal/filecontext ./internal/contracts ./internal/cli ./internal/httpapi -count=1` | Pass. |
| `go run .\cmd\generate-web-contracts --check` | Pass; generated Web contracts match source. |
| `go test ./cmd/... ./internal/... -count=1` | Pass; all cmd/internal packages passed. |
| `.\anvien\bin\anvien.exe analyze --force --name Anvien --benchmark-label file-centric-p3b-linked-overlays --benchmark-json .tmp\file-centric-p3b-linked-overlays-analyze-benchmark.json` | Pass. `files: scanned=825 parsed=590 unsupported=235 failed=0`; `nodes=93728 relationships=128296`; total graph generation time `34530.4 ms`. |
| `.\anvien\bin\anvien.exe graph-health summary --repo Anvien --json` | Pass after P3-B. `sourceBackedUnresolvedReferenceCount=68028`; `resolutionGapNodeCount=67120`; `resolvedReferenceCount=30896`. |
| `.\anvien\bin\anvien.exe file-context internal\httpapi\file_context.go --repo Anvien --json` | Pass. Summary: `symbols=62`, `unresolved=144`, linked counts `flows=7 routes=0 mcpTools=0 tests=3`, samples `flows=5 tests=3`, `stale=false`. |
| `.\anvien\bin\anvien.exe file-context internal\mcp\tools.go --repo Anvien --json --linked 10` | Pass. Summary: `symbols=152`, `unresolved=618`, linked counts `flows=18 routes=0 mcpTools=0 tests=2`, samples `flows=10 tests=2`, `stale=false`. |
| `.\anvien\bin\anvien.exe file-hotspots --repo Anvien --json --limit 0` | Pass. `total=825`; files with linked flows `162`; files with linked tests `348`; top flow file `internal/httpapi/response.go` with `55`; top test-linked file `internal/scopeir/scope_tree.go` with `86`. |
| `go test ./internal/filecontext -run '^$' -bench BenchmarkBuildFileListCurrentScale -benchmem -count=3` | Pass. Runs: `68.5 ms/op`, `58.7 ms/op`, `54.2 ms/op`; `490000 B/op`, `831 allocs/op`. |
| `.\anvien\bin\anvien.exe api route-map --repo Anvien --json` | Returned no routes in the current validation graph. |
| `.\anvien\bin\anvien.exe api tool-map context --repo Anvien --json` | Returned no matching tools in the current validation graph. |

Failures / handling:

- The current validation graph has no route/tool map facts available for runtime route/MCP tool samples, so real-repo `file-context` smoke shows `routes=0` and `mcpTools=0`.
- P3-B still implements route/tool projection from `HANDLES_ROUTE` and `HANDLES_TOOL`, and fixture tests prove those links are emitted when the graph supplies the facts.
- `linked.counts` was added to the contract because sample arrays are intentionally bounded; without full counts, a Web or CLI consumer could mistake a truncated linked sample list for total coverage.

Benchmark link:

- See B0A for graph generation speed.
- See B1 for file inventory and linked flow/test coverage counts.
- See B5 and B6 for projection timing and response-size metrics.

Detect changes:

| Command | Result |
|---|---|
| `.\anvien\bin\anvien.exe detect-changes --repo Anvien --scope all` | `risk_level=high`; `changed_files=8`; `changed_count=234`; `affected_count=10`; changed app layers `api_contract`, `api_test`, `backend`, `backend_test`, and `docs`; affected app layer `backend`; resolution gap changed entities `175`. HIGH is expected because the shared projection contract, projection builder, generated Web contract, and tests changed; behavior remains additive and repo-agnostic. |

## E9 - Web UI File Map And File Detail

Date: 2026-05-30

Status: P4-A completed; P4-B pending

Scope:

- Add the Web File Map list to the existing left dashboard navigation.
- Keep Web behavior repo-agnostic by passing the selected repo name into `/api/file-hotspots`; `Anvien` is only the local validation repo.
- Add typed Web client support for file hotspot queries.
- Add changed-file support to the shared file list projection and HTTP API so the File Map can expose the full P4-A filter set.
- Do not implement File Detail in this slice; P4-B remains pending.

Impact / blast radius:

| Command | Result |
|---|---|
| `.\anvien\bin\anvien.exe analyze --force --name Anvien` | Pass before P4-A impact checks. Initial P4-A graph: `files=825`, `nodes=93728`, `relationships=128296`. |
| `.\anvien\bin\anvien.exe impact "FileTreePanel" --repo Anvien --direction upstream` | `risk=LOW`; `impactedCount=2`; affected Web entry points `App.tsx` and `main.tsx`. |
| `.\anvien\bin\anvien.exe impact --uid "Function:anvien-web/src/services/backend-client.ts:fetchGraph" --repo Anvien --direction upstream` | `risk=CRITICAL`; `impactedCount=31`; `processes_affected=46`. Proceeded because `fetchGraph` behavior was not changed; the shared client file only gained additive file-hotspot client code. |
| `.\anvien\bin\anvien.exe analyze --force --name Anvien` | Pass before adding changed-file API support. `files=827`, `nodes=94053`, `relationships=128700`. |
| `.\anvien\bin\anvien.exe impact "BuildFileList" --repo Anvien --direction upstream` | `risk=LOW`; `impactedCount=0`. |
| `.\anvien\bin\anvien.exe impact "handleFileHotspots" --repo Anvien --direction upstream` | `risk=LOW`; `impactedCount=0`. |
| `.\anvien\bin\anvien.exe impact --uid "Struct:internal/filecontext/context.go:FileListOptions" --repo Anvien --direction upstream` | `risk=LOW`; `impactedCount=2`; direct users are `BuildFileList` and `filterSummaries`. |
| `.\anvien\bin\anvien.exe impact "WebUIContractTypeScript" --repo Anvien --direction upstream` | `risk=CRITICAL`; `impactedCount=1`; `processes_affected=5`. Proceeded because generated TypeScript contract changes are additive. |

Implementation evidence:

| File | Evidence |
|---|---|
| `anvien-web/src/components/FileMapPanel.tsx` | New File Map panel renders summary counts, search, sort, kind filters, changed/unresolved/API/high fan-in/high fan-out filters, loading/error/empty states, file rows, and changed-file row badges. |
| `anvien-web/src/components/FileTreePanel.tsx` | Added `Map` tab and collapsed-panel icon; selecting a File Map row focuses the matching graph file node and opens the existing code panel. |
| `anvien-web/src/services/backend-client.ts` | Added repo-parameterized `fetchFileHotspots` client. Query params include `repo`, `sort`, pagination, kind/layer filters, `changedOnly`, `unresolvedOnly`, `apiOnly`, `highFanIn`, and `highFanOut`. |
| `internal/filecontext/context.go` | Added `ChangedOnly`, `ChangedPaths`, and `Stale` to `FileListOptions`; `FileSummary` now includes `stale` and `changedSinceAnalyze`. Filtering remains in the shared projection builder. |
| `internal/httpapi/file_context.go` | `GET /api/file-hotspots` now accepts `changedOnly`; changed files are derived from the selected repo's git diff, including untracked files, and passed into the shared projection. |
| `internal/contracts/web_ui.go` | Added `changedOnly` to `/api/file-hotspots` route metadata and added `stale`/`changedSinceAnalyze` to `FileSummary`. |
| `contracts/web-ui/anvien-web-contract.schema.json` | Regenerated from contract source. |
| `anvien-web/src/generated/anvien-contracts.ts` | Regenerated from contract source. |
| `internal/filecontext/context_test.go` | Added changed-file filter and changed metadata test coverage. |
| `anvien-web/test/unit/FileMapPanel.test.tsx` | Added File Map rendering, filter/sort request state, and row-open tests. |
| `anvien-web/test/unit/server-connection.test.ts` | Added `fetchFileHotspots` URL/query-param test coverage. |
| `anvien-web/e2e/shell-interactions.spec.ts` | Added File Map e2e smoke and repo-agnostic graph-ready timeout override for large local validation repos. |

Validation:

| Command | Result |
|---|---|
| `go run .\cmd\generate-web-contracts` | Pass; regenerated Web schema and TypeScript contracts. |
| `go run .\cmd\generate-web-contracts --check` | Pass. |
| `go test ./internal/filecontext ./internal/httpapi ./internal/contracts -count=1` | Pass. |
| `powershell -ExecutionPolicy Bypass -File anvien-launcher\build.ps1` | First rerun failed because the previously started local `anvien serve` process held `anvien\bin\anvien.exe`; that launcher-owned serve process was stopped, then the build passed. Existing Vite dynamic-import and chunk-size warnings only. |
| `npm --prefix anvien-web test -- FileMapPanel.test.tsx server-connection.test.ts` | Pass after final build; `2` files, `23` tests. |
| `npm --prefix anvien-web run test:e2e -- shell-interactions.spec.ts -g "file map"` with `E2E_REPO_NAME=Anvien` | Pass after final build/runtime restart; `1` Chromium test, `2.9m`. |
| `go test ./cmd/... ./internal/... -count=1` | Pass; all cmd/internal packages passed after P4-A changes. |
| `.\anvien\bin\anvien.exe analyze --force --name Anvien --benchmark-label file-centric-p4a-web-file-map --benchmark-json .tmp\file-centric-p4a-web-file-map-analyze-benchmark.json` | Pass. `files=827`, `nodes=94171`, `relationships=128835`; total graph generation time `35958.2 ms`. |
| `.\anvien\bin\anvien.exe graph-health summary --repo Anvien --json` | Pass after P4-A. `sourceBackedUnresolvedReferenceCount=68414`; `resolutionGapNodeCount=67506`; `resolvedReferenceCount=30964`. |
| HTTP runtime smoke `GET /api/file-hotspots?repo=Anvien&changedOnly=true&limit=5` | Pass from restarted local binary. `status=200`; `total=11`; `rows=5`; first row `changedSinceAnalyze=true`; repo selector remains an input. |
| Browser validation | In-app Browser control tool was not exposed after loading the Browser skill; standalone Playwright smoke was used as fallback. Manual smoke loaded File Map on `Anvien`, rendered `200` rows, toggled unresolved, and changed sort to `fan-out`. |

Failures / handling:

- Large local graph load for `Anvien` takes about `132s` before `status-ready` in headless Playwright. The File Map e2e now uses `test.slow()` and a repo-agnostic `E2E_GRAPH_READY_TIMEOUT` override, defaulting to `180s` only when `E2E_REPO_NAME` is explicitly set.
- Build failed once because a local validation `anvien serve` process locked the binary output path. Only that launcher-owned serve process was stopped; editor-owned MCP processes were left alone.
- The changed-file filter uses git diff for the selected repo path. It is repo-agnostic and does not special-case `Anvien`.

Benchmark link:

- See B0/B0A for P4-A graph inventory and graph generation speed.
- See B6/B7 for Web/API response size and File Map UI metrics.
- See B10 for validation counts.

Detect changes:

| Command | Result |
|---|---|
| `.\anvien\bin\anvien.exe detect-changes --repo Anvien --scope all` | `risk_level=high`; `changed_files=13`; `changed_count=260`; `affected_count=14`; changed app layers `api`, `api_contract`, `backend`, `backend_test`, `docs`, `frontend`, `frontend_api_client`, and `frontend_test`; affected app layers `api`, `api_contract`, `backend`, and `mixed`; resolution gap changed entities `194`. HIGH is expected because the slice touches shared projection data, HTTP API, generated contracts, Web UI, Web client, and tests. |

Commit:

- `e4efd84 feat: add web file map list`

## E10 - Parent/Child Command Hierarchy

Date: 2026-05-30

Status: P5-A completed

Scope:

- Target-aware parent commands inventoried: `context`, `impact`, `query`, `detect-changes`, `graph-health`, `api route-map`, `api tool-map`, `api shape-check`, `api impact`, and `group query`.
- Parent commands keep current usage. `context <target>` and `impact <target>` now pass `target_type=auto` for smart dispatch; `query <text>` remains broad multi-lane search; `detect-changes` remains broad changed-symbol/affected-flow output; `graph-health summary/report/components/explain` remain unchanged.
- Child commands defined and implemented for target-specific views: `context symbol`, `context file`, `impact symbol`, `impact file`, `impact route`, `impact tool`, `query files`, `query symbols`, `query flows`, `query api`, `detect-changes files`, `detect-changes symbols`, `detect-changes flows`, and `graph-health files`.
- Shared target semantics use the same field names across CLI/MCP payloads: `target_type`, `dispatch_mode`, `targetType`, `dispatchMode`, `selectedFile`, `selectedSymbol`, and explicit `suggestedCommand` values on ambiguity candidates.
- Commands intentionally left without Phase 5 child commands: `analyze` remains repo/path oriented; `rename` remains symbol-first; `cypher` remains raw graph-query oriented; `resolution-inventory` remains graph-snapshot oriented while file quality rows are exposed through `graph-health files`; group commands remain group/repo oriented until cross-repo file projection is designed.

Impact / blast radius:

| Command | Result |
|---|---|
| `.\anvien\bin\anvien.exe analyze --force --name Anvien` | Pass before P5 graph work. `files=827`, `nodes=94171`, `relationships=128835`. |
| `.\anvien\bin\anvien.exe impact "newContextCommand" --repo Anvien --direction upstream` | `risk=CRITICAL`; direct dependent `NewRootCommand`, then CLI `main`; proceeded because syntax is additive and parent command still exists. |
| `.\anvien\bin\anvien.exe impact "newImpactCommand" --repo Anvien --direction upstream` | `risk=CRITICAL`; same CLI root blast radius; additive child commands and JSON flag only. |
| `.\anvien\bin\anvien.exe impact "newQueryCommand" --repo Anvien --direction upstream` | `risk=CRITICAL`; same CLI root blast radius; parent query remains broad. |
| `.\anvien\bin\anvien.exe impact "newDetectChangesCommand" --repo Anvien --direction upstream` | `risk=CRITICAL`; same CLI root blast radius; parent detect output remains available. |
| `.\anvien\bin\anvien.exe impact "newGraphHealthCommand" --repo Anvien --direction upstream` | `risk=CRITICAL`; same CLI root blast radius; existing graph-health subcommands unchanged. |
| `.\anvien\bin\anvien.exe impact "contextToolInternal" --repo Anvien --direction upstream` | `risk=LOW`; no upstream impacted nodes reported. |
| `.\anvien\bin\anvien.exe impact "impactToolInternal" --repo Anvien --direction upstream` | `risk=LOW`; no upstream impacted nodes reported. |
| `.\anvien\bin\anvien.exe impact "queryTool" --repo Anvien --direction upstream` | `risk=LOW`; no upstream impacted nodes reported. |
| `.\anvien\bin\anvien.exe impact "detectChangesTool" --repo Anvien --direction upstream` | `risk=LOW`; no upstream impacted nodes reported. |
| `.\anvien\bin\anvien.exe impact --uid "Function:internal/mcp/tools.go:mcpTools#0" --repo Anvien --direction upstream` | `risk=CRITICAL`; direct dependent `Server.handle`; proceeded because MCP schema changes are additive optional fields. |

Implementation evidence:

| File | Evidence |
|---|---|
| `internal/mcp/target_dispatch.go` | Added shared target constants, target/dispatch normalization, file context helpers, ambiguity suggestion helpers, file impact aggregation helpers, and query/detect target-type allowlists. |
| `internal/cli/tool_command.go` | Added child commands and inherited flags for `query`, `context`, `impact`, and `detect-changes`; added `--json` stripping for parent/child MCP command output where needed. |
| `internal/mcp/tools.go` | Added optional MCP schema fields `target_type` and `dispatch_mode` for `query`, `context`, `impact`, and `detect_changes`. |
| `internal/mcp/testdata/typescript_baseline_surface.json` | Updated MCP surface snapshot to include optional target dispatch fields. |
| `internal/mcp/surface_snapshot_test.go` | Added schema assertions for target dispatch fields. |

## E11 - Context And Impact Child Commands

Date: 2026-05-30

Status: P5-B completed

Scope:

- `context symbol <symbol>` forces symbol lookup and adds `fileSummary`, `selectedFile`, and `selectedSymbol` to the MCP result.
- `context file <path>` forces full file projection under `fileContext`, with `targetType=file` and `dispatchMode=explicit`.
- Parent `context <target>` uses smart dispatch and returns an ambiguity payload with exact `anvien context file ...` and `anvien context symbol ...` suggestions when a target matches both layers.
- `impact symbol <symbol>` keeps symbol blast radius and adds `affectedFiles` / file summary evidence.
- `impact file <path>` aggregates impact across contained symbols and reports `containedSymbols`, `symbolImpacts`, `impacted`, `affectedFiles`, `affected_processes`, linked flow/test counts, and an aggregate risk.
- `impact route <route>` delegates to existing API impact data with `targetType=route`.
- `impact tool <tool>` reports matched MCP tool definitions and linked flow count with `targetType=tool`.

Implementation evidence:

| File | Evidence |
|---|---|
| `internal/mcp/context.go` | Added `target_type=auto|symbol|file`, smart file-vs-symbol dispatch, explicit file context payloads, symbol file summaries, and ambiguity suggestions. |
| `internal/mcp/impact.go` | Added `target_type=auto|symbol|file|route|tool`, file aggregate impact, route impact wrapping, tool impact summaries, symbol file summaries, and affected-file grouping. |
| `internal/cli/target_command_test.go` | Added CLI tests for `context file`, `context symbol`, `impact file`, JSON output, file summaries, affected file evidence, and ambiguous parent suggestions. |

## E12 - Query, Change, And Quality Child Commands

Date: 2026-05-30

Status: P5-C/P5-D completed

Scope:

- `query <text>` remains the broad multi-lane parent behavior.
- `query files <text>` returns file-first rows with `summary` and bounded `matchedSymbols`.
- `query symbols <text>` returns symbol-first rows and adds containing `fileSummary`.
- `query flows <text>` narrows output to execution-flow rows.
- `query api <text>` narrows output to route-map and tool-map rows.
- `detect-changes files|symbols|flows` was implemented as child commands rather than flags because the parent command already has scope/base-ref flags and the result layer changes the primary payload shape.
- `graph-health files` was implemented as the file-level quality child command. `resolution-inventory` remains a graph snapshot inventory command; file-oriented quality triage belongs to `graph-health files` for this phase.
- MCP/API/generated alignment: CLI and MCP use the same target names. Existing HTTP/API file projection routes already use file target semantics; no Web contract source shape changed in this slice.

Implementation evidence:

| File | Evidence |
|---|---|
| `internal/mcp/tools.go` | Added target-specific query payloads for `files`, `symbols`, `flows`, and `api`; parent broad query payload remains intact. |
| `internal/mcp/detect_changes.go` | Added target-specific changed `files`, `symbols`, and `flows` payloads with shared `targetType`/`dispatchMode` fields. |
| `internal/cli/graph_health_command.go` | Added `graph-health files` with JSON and human output for file-level unresolved/fan-in/fan-out/risk rows. |
| `internal/cli/target_command_test.go` | Added tests for `query files`, `query symbols`, `graph-health files`, and `detect-changes files`. |

Validation so far:

| Command | Result |
|---|---|
| `powershell -ExecutionPolicy Bypass -File anvien-launcher\build.ps1` | Pass before final tests. Existing Vite dynamic-import and chunk-size warnings only. |
| `go test ./cmd/... ./internal/... -count=1` | Pass after full build. |
| `go run .\cmd\generate-web-contracts --check` | Pass; no Web generated contract drift from P5. |
| `.\anvien\bin\anvien.exe analyze --force --name Anvien` | Pass after P5 implementation. Final P5 graph: `files=829`, `nodes=94838`, `relationships=129861`. |
| Built binary smoke: `context file internal\cli\tool_command.go --repo Anvien --json` | Pass. Parsed summary: `targetType=file`, `dispatchMode=explicit`, `path=internal/cli/tool_command.go`, `symbols=69`. |
| Built binary smoke: `query files "target dispatch" --repo Anvien --limit 3 --json` | Pass. Parsed summary: `targetType=files`, `total=3`, first row `internal/mcp/impact.go`. |
| Built binary smoke: `impact file internal\cli\tool_command.go --repo Anvien --depth 1 --json` | Pass. Parsed summary: `targetType=file`, `impactedCount=21`, `affectedFiles=5`, `risk=CRITICAL`. CRITICAL is expected for CLI command-surface owners. |
| Built binary smoke: `graph-health files --repo Anvien --limit 3 --json` | Pass. Parsed summary: `total=829`, `returned=3`, first row `internal/mcp/server_test.go`. |
| Built binary smoke: `detect-changes files --repo Anvien --scope all --json` | Pass before final evidence edit. Parsed summary: `targetType=files`, `total=11`, `changed_count=539`, `affected_count=59`. |
| Compatibility smoke: `query files --repo Anvien --limit 1 --json` | Pass. Missing child query argument falls back to parent broad query for `files`; parsed `query=files`, `targetType=null`, `definitions=1`. |

Benchmark link:

- See B0/B1 for P5 graph/file inventory counts.
- See B8 for parent/child command hierarchy counts.
- See B10 for validation test counts.

## E13 - Existing Command Integration Matrix

Date: 2026-05-30

Status: P6-A completed

Matrix:

| Command family | Classification | Preserved output | Added file-layer behavior |
|---|---|---|---|
| `analyze` | must add file layer | Scanned/parsed/unsupported/failed counts and graph node/relationship inventory. | Human `fileProjection` summary and top hotspots; `--json` now includes `fileProjection.status`, `files`, `dependencyEdges`, `unresolvedFiles`, `hotspots`, and `derivedEdgesNote`. |
| `query` / `query files` | must add file layer | Existing definitions/processes/docs result lanes. | Parent JSON adds `files` and `fileLayer`; file rows include matched symbols and relationship hints. |
| `context` / `context symbol` / `context file` | must add file layer | Existing symbol context, references, and target dispatch behavior. | File paths open full file context; symbol context includes `fileLayer` with containing file summary, relationship counts, unresolved counts, and linked evidence. |
| `impact` / `impact symbol` / `impact file` / `impact route` / `impact tool` | must add file layer | Existing symbol blast radius, process/test impact, route/tool impact payloads. | Symbol/file output adds file-level blast radius; tool output adds handler file rows; route/tool impacts keep target-specific semantics. |
| `detect-changes` / `detect-changes files` | must add file layer | Changed symbols, affected processes, semantic status, app-layer and functional-area summaries. | Parent JSON adds `changed_files`, `affected_files`, and `fileLayer`; file child rows include relationship hints, linked flows/tests, unresolved delta, changed symbols, and file risk. |
| `graph-health summary/report/files` | must add file layer | Existing topology, resolution, source-site, and component diagnostics. | Summary/report add `fileLayer` and top `fileHotspots`; `graph-health files` remains the detailed file quality child command. |
| `query-health` | must add file layer | Existing query case scoring, expected/matched/missed targets, threshold/exact checks. | Actual results include file rows; misses include `missedClusters` grouped by file, app layer, functional area, and target layer. |
| `resolution-inventory` | must add file layer | Existing graph snapshot totals and gap buckets. | Adds top `fileGroups` with gap totals, buckets, nearest source symbols, and samples. |
| `source-site-accuracy` | must add file layer | Existing source-site policy issues and summary lines. | Adds `fileGroups` with issue counts and source-site samples including `filePath` and `startLine`. |
| `api route-map`, `api tool-map`, `api shape-check`, `api impact` | must add file layer | Existing route/tool/consumer/shape/impact details. | Route/tool records include `handlerFile`; shape and impact summaries propagate handler-file evidence. |
| MCP `query`, `context`, `impact`, `detect_changes`, API map equivalents, resources | must add file layer | Existing tool schemas and result semantics. | Structured payloads now expose the same file-layer facts, handler files, and resource guidance. |
| `status`, `list` | no file layer | Repository freshness and registry inventory. | Left unchanged because they do not answer graph-content questions. |
| `rename` | no default file layer | Symbol rename edits and graph edits. | Left unchanged in P6 to avoid expanding mutation output; file-level impact remains available through `impact`. |
| `augment`, `cypher` | no file layer | Raw/search augmentation and raw graph query behavior. | Left unchanged because reshaping raw results would hide caller-selected graph facts. |
| Group commands | may add file layer later | Cross-repo group status/query/contracts. | Deferred until cross-repo file projection semantics are designed and labeled per repo. |

Common wording:

- All added file relationship output uses: `File relationship groups are projections derived from symbol and source-site graph facts; canonical graph relationships remain symbol/source-site facts.`
- Parent commands keep broad discovery behavior. Child commands remain the exact target-specific path when scripts or agents need `target_type=file`, `target_type=symbol`, route, tool, flow, or quality-specific output.

## E14 - Existing Command File-Layer Behavior

Date: 2026-05-30

Status: P6-B/P6-C completed

Blast radius:

- Impact checks before edits reported CRITICAL blast radius for CLI root command owners (`newAnalyzeCommand`, graph-health/report, resolution inventory, source-site accuracy, query-health, generated AI context) because they are reachable from the root CLI and generated guidance.
- MCP tool internals for `query`, `context`, `impact`, `detect_changes`, route/tool maps, shape check, and API impact were LOW to CRITICAL depending on shared helper reuse.
- Changes were kept additive: existing symbol, route/tool, graph-quality, and changed-symbol payloads remain present while file-layer fields are appended.

Implementation evidence:

| Area | Files | Evidence |
|---|---|---|
| Analyze | `internal/cli/command.go`, `internal/cli/command_test.go` | Human output adds `fileProjection`; `analyze --json` adds machine-readable `fileProjection`. |
| Query/context/impact/detect | `internal/mcp/tools.go`, `internal/mcp/context.go`, `internal/mcp/impact.go`, `internal/mcp/detect_changes.go`, `internal/mcp/target_dispatch.go`, `internal/cli/target_command_test.go` | Parent and child payloads expose file summaries, relationship hints, file-layer blast radius, changed files, and affected files. |
| Graph quality | `internal/cli/graph_health_command.go`, `internal/cli/resolution_inventory_command.go`, `internal/graphaccuracy/source_site_accuracy.go`, `internal/cli/query_health_command.go`, related tests | Graph-health, resolution inventory, source-site accuracy, and query-health now include file-level grouping or miss clustering. |
| API/MCP maps | `internal/mcp/route_tool_map.go`, `internal/mcp/route_shape_impact.go`, `internal/mcp/server_test.go` | Route/tool/shape/impact fixture tests assert `handlerFile` propagation. Live Anvien graph has no route/tool rows, so parity evidence uses fixture graph tests. |
| MCP resources | `internal/mcp/resources.go` | `anvien://repo/<repo>/context` and setup guidance describe file projection, child commands, and graph-quality file outputs. |

Validation:

| Command | Result |
|---|---|
| `powershell -ExecutionPolicy Bypass -File anvien-launcher\build.ps1` | Pass after P6 code changes. Existing Vite dynamic-import and chunk-size warnings only. |
| `go test ./... -count=1` | Expected fixture failure only: non-buildable analyzer fixtures under `anvien/test/fixtures` import local packages or include invalid examples. |
| `go test ./cmd/... ./internal/... -count=1` | Pass after full build and after adding `analyze --json`. |
| `go run .\cmd\generate-web-contracts --check` | Pass; no generated Web contract drift from P6. |
| `.\anvien\bin\anvien.exe analyze --force --name Anvien` | Pass. `files=829`, `parsed=594`, `nodes=95419`, `relationships=130701`, `dependencyEdges=15806`, `unresolvedFiles=586`, `hotspots=5`. |
| `.\anvien\bin\anvien.exe analyze --force --name Anvien --json` | Pass. Parsed summary: `fileProjection.files=829`, `dependencyEdges=15806`, `hotspots=5`, `graph.nodes=95419`. |
| `query "target dispatch" --repo Anvien --limit 2 --json` | Pass. Existing definitions lane preserved (`definitions=2`), added `files=2`, `fileLayer=true`, first file `internal/mcp/impact.go`. |
| `query files "target dispatch" --repo Anvien --limit 3 --json` | Pass. `files=3`, first file `internal/mcp/impact.go`, relationship hints returned. |
| `context symbol "newAnalyzeCommand" --repo Anvien --json` | Pass. Symbol payload includes `fileLayer.path=internal/cli/command.go`, unresolved count `217`. |
| `context file internal\cli\command.go --repo Anvien --json` | Pass. Full file context: `symbols=69`, `fanIn=6`, `fanOut=143`, `risk=high`. |
| `impact symbol "newAnalyzeCommand" --repo Anvien --direction upstream --json` | Pass. Existing symbol risk remains `CRITICAL`; file-layer blast radius reports `affectedFiles=2`. |
| `impact file internal\mcp\route_tool_map.go --repo Anvien --direction upstream --depth 1 --json` | Pass. `target.normalizedPath=internal/mcp/route_tool_map.go`, `risk=CRITICAL`, file-layer `affectedFiles=6`. |
| `graph-health summary --repo Anvien --json` | Pass. `fileLayer.totalFiles=829`, `unresolvedFiles=586`, `highFanOutFiles=245`, `fileHotspots=5`. |
| `graph-health files --repo Anvien --limit 3 --json` | Pass. `files=3`, first hotspot `internal/mcp/server_test.go`, `risk=high`. |
| `resolution-inventory --graph .anvien\graph.json --json` | Pass. `fileGroups=20`, first file `internal/mcp/server_test.go`, first total `1445`. |
| `source-site-accuracy --graph .anvien\graph.json --json` | Pass. `fileGroups=586`, first file `internal/mcp/server_test.go`, first total `1445`. |
| `query-health --repo Anvien --json` | Pass at command level. Suite summary remains `cases=7`, `passed=6`, `failed=1`; file/layer `missedClusters` present. |
| Final `detect-changes --repo Anvien --scope all --json` | Pass before commit. `changed_symbols=745`, `changed_files=31`, `affected_files=31`, summary `affected_count=117`, `risk_level=critical`, `fileLayer=true`. CRITICAL is expected because the slice touches CLI root output, MCP/API payloads, graph-quality commands, AI context generation, docs, and tests. |

## E15 - Generated Skills And AI Context

Date: 2026-05-30

Status: P6-D completed

Source-of-truth updates:

| File | Evidence |
|---|---|
| `internal/aicontext/aicontext.go` | Generated AGENTS/CLAUDE command guidance now lists `file-context`, `file-hotspots`, `query files`, `context file`, `context symbol`, `impact file`, `impact symbol`, `detect-changes files`, and `graph-health files/file-hotspots`. |
| `internal/aicontext/skills/anvien-cli.md` | CLI skill now documents parent/child file-layer commands and validation notes. |
| `internal/aicontext/skills/anvien-exploring.md` | Exploration workflow now starts from query/analyze hotspots, then drills into file context, symbol tree, relationships, and source-site evidence. |
| `internal/aicontext/skills/anvien-impact-analysis.md` | Impact workflow now distinguishes `impact symbol` and `impact file`, and uses `detect-changes files` for changed-scope evidence. |
| `internal/aicontext/skills/anvien-graph-quality.md` | Graph quality workflow now routes unresolved/source-site issues through file groups and `graph-health files`. |
| `internal/aicontext/skills/anvien-api-surface.md` | API surface workflow now checks handler files, `handlerFile`, route/tool maps, and `context file` before edits. |
| `internal/aicontext/skills/anvien-debugging.md` | Debugging workflow now uses `query files`, `context file`, and `impact file` when failures are file-scoped. |
| `internal/aicontext/skills/anvien-guide.md` | Unified guide now includes file-layer command spelling. The former self-referential `anvien-ai-context` skill is retired and is no longer part of the generated skill set. |

Regeneration / parity:

- `.\anvien\bin\anvien.exe analyze --force --name Anvien` ran after generator-owned source updates, using the normal `generateAnalyzeAIContext` path. Generated root context files and `.claude/skills/anvien/**` are not tracked in this repository, so tracked diffs remain in generator-owned source and tests.
- `internal/aicontext/aicontext_test.go` now asserts generated guidance for command spellings, generated skill content, `handlerFile`, `target_type=file`, and file-layer workflow text.
- `go test ./internal/aicontext -count=1` passed as part of the full `go test ./cmd/... ./internal/... -count=1` run.

## E15A - Web File Detail View

Date: 2026-06-01

Status: P4-B completed

Scope:

- Add Web UI File Detail rendering for the existing `/api/file-context` contract.
- Keep file projection semantics backend-owned; the frontend only fetches and renders typed contract sections.
- Extend File Map e2e so clicking a file opens File Detail and verifies major sections.

Impact / blast radius:

| Command | Result |
|---|---|
| `.\anvien\bin\anvien.exe analyze --force --name Anvien` | Pass before P4-B graph-based checks. `files=829`, `parsed=594`, `nodes=95419`, `relationships=130701`; file projection `dependencyEdges=15806`, `unresolvedFiles=586`. |
| `.\anvien\bin\anvien.exe query "File Map File Detail web component file context" --repo Anvien --limit 8` | Relevant Web owners identified: `FileTreePanel`, `FileMapPanel`, `CodeReferencesPanel`, and `backend-client`. |
| `.\anvien\bin\anvien.exe impact --uid "Function:anvien-web/src/components/FileTreePanel.tsx:FileTreePanel" --repo Anvien --direction upstream` | LOW; impacted `App.tsx` and `main.tsx`. |
| `.\anvien\bin\anvien.exe impact --uid "Function:anvien-web/src/components/FileMapPanel.tsx:FileMapPanel" --repo Anvien --direction upstream` | LOW; impacted `FileTreePanel.tsx`, `App.tsx`, and `main.tsx`. |
| `.\anvien\bin\anvien.exe impact --uid "Function:anvien-web/src/services/backend-client.ts:fetchFileHotspots" --repo Anvien --direction upstream` | LOW symbol impact; file-level note: shared API client has broad inbound Web usage, so edits stayed additive. |
| `.\anvien\bin\anvien.exe impact --uid "Function:anvien-web/src/components/CodeReferencesPanel.tsx:CodeReferencesPanel" --repo Anvien --direction upstream` | LOW; impacted `App.tsx` and `main.tsx`. |

Implementation evidence:

| File | Evidence |
|---|---|
| `anvien-web/src/components/FileDetailPanel.tsx` | New File Detail component fetches `fetchFileContext`, renders summary, quality, expandable symbol tree, local/outbound/inbound relationship groups, unresolved source-site samples, linked flows/routes/MCP tools/tests, loading, error, empty states, and graph-focus buttons for symbol rows. |
| `anvien-web/src/services/backend-client.ts` | Added typed `fetchFileContext(path, { repo, relationships, unresolved, linked })` for `GET /api/file-context`. |
| `anvien-web/src/components/CodeReferencesPanel.tsx` | Renders File Detail inside the selected file viewer when the selected graph node is a `File`. |
| `anvien-web/test/unit/FileDetailPanel.test.tsx` | Covers successful rendering of all major sections and symbol focus behavior, plus backend error state. |
| `anvien-web/test/unit/server-connection.test.ts` | Covers `/api/file-context` URL/query construction and sample-limit params. |
| `anvien-web/e2e/shell-interactions.spec.ts` | File Map e2e now opens the first file row through its path button and verifies File Detail summary, symbol tree, relationships, unresolved, and linked sections. |
| `README.md` | User-facing docs now describe File Map/File Detail, file-aware CLI commands, graph-health files, MCP file-layer wording, and `/api/file-context` / `/api/file-hotspots`. |

Validation:

| Command | Result |
|---|---|
| `powershell -ExecutionPolicy Bypass -File anvien-launcher\build.ps1` | Pass. Existing Vite dynamic-import and chunk-size warnings only. |
| `npm --prefix anvien-web test -- FileDetailPanel.test.tsx FileMapPanel.test.tsx server-connection.test.ts CodeReferencesPanel.graph-health.test.tsx` | Pass. 4 files, 27 tests. |
| `.\anvien\bin\anvien.exe analyze --force --name Anvien` | Pass after P4-B implementation. `files=831`, `parsed=596`, `nodes=95884`, `relationships=131220`; file projection `dependencyEdges=15816`, `unresolvedFiles=588`. |
| `npm --prefix anvien-web run test:e2e -- shell-interactions.spec.ts -g "file map"` with `E2E=1`, `E2E_REPO_NAME=Anvien` | Pass after selector fix. 1 Chromium test, 3.1m. |
| `go test ./cmd/... ./internal/... -count=1` | Pass. |
| `go run .\cmd\generate-web-contracts --check` | Pass; no generated Web contract drift. |
| `npm --prefix anvien-web test` | Pass. 52 files, 408 tests. |

Command/API smoke evidence:

| Command | Result |
|---|---|
| `file-context internal/httpapi/file_context.go --repo Anvien --json` | `symbols=68`, `in=9`, `out=34`, `unresolved=164`, `flows=6`, `tests=3`, `risk=high`. |
| `file-context internal/httpapi/file_context_test.go --repo Anvien --json` | `symbols=14`, `out=57`, `unresolved=105`, `risk=high`. |
| `file-context README.md --repo Anvien --json` | Docs/config-style smoke: `symbols=0`, `unresolved=0`, `risk=low`. |
| `file-context contracts/web-ui/anvien-web-contract.schema.json --repo Anvien --json` | Generated file smoke: `kind=generated`, `appLayer=generated_contract`, `symbols=0`, `risk=low`. |
| `file-context internal/mcp/server.go --repo Anvien --json` | API/MCP file smoke: `symbols=91`, `in=66`, `out=45`, `unresolved=206`, `flows=12`, `tests=13`, `risk=high`. |
| `file-context internal/mcp/server_test.go --repo Anvien --json` | Unresolved-heavy file smoke: `symbols=238`, `out=197`, `unresolved=1445`, `tests=4`, `risk=high`. |
| `file-hotspots --repo Anvien --sort unresolved|fan-in|fan-out|symbols|flows|tests|path --limit 3 --json` | Pass for all seven sort modes; total `831` files. |
| `query files "file detail" --repo Anvien --limit 3 --json` | Pass. `targetType=files`, `total=3`, first row `internal/cli/command_test.go`. |
| `context symbol BuildFileContext --repo Anvien --json` | Pass. Symbol context includes file-layer path `internal/filecontext/context.go`, unresolved `479`. |
| `graph-health files --repo Anvien --limit 3 --json` | Pass. `total=831`, first hotspot `internal/mcp/server_test.go`. |

Failures / handling:

- First focused e2e attempt timed out waiting for File Map rows while the backend file projection endpoint was cold after server startup. A manual warm request returned the expected `831` files.
- Second e2e attempt opened the file list but clicked the table row, while the UI action is attached to the path button. The e2e was corrected to click the row button directly and then passed.
- Browser plugin navigation was requested through tool discovery, but only Figma/GitHub/Drive/Anvien deferred tools were exposed. Playwright e2e is the UI validation evidence for this slice.

Benchmark link:

- See B0/B1/B6/B7/B10 for updated graph inventory, file inventory, response size, Web UI, and validation counts.

## E16 - Final Validation And Closure

Date: 2026-06-01

Status: P7-A/P7-B completed

Scope:

- Close the file-centric projection plan with README updates, P4-B Web UI evidence, validation, benchmark metrics, and final change detection.
- Keep generated contracts unchanged; Web consumes the existing `FileContextResponse` contract.

Docs evidence:

| File | Evidence |
|---|---|
| `README.md` | Added File Map/File Detail capability notes, file-aware CLI command examples, `graph-health files`, MCP file-layer wording, and file projection HTTP endpoints. |
| `docs/plans/2026-05-30-anvien-file-centric-graph-projection-plan.md` | Marked P4-B, P7-A, and P7-B complete. |
| `docs/plans/2026-05-30-anvien-file-centric-graph-projection-evidence.md` | Added P4-B Web File Detail implementation, validation, failure handling, and closure evidence. |
| `docs/plans/2026-05-30-anvien-file-centric-graph-projection-benchmark.md` | Updated graph/file inventory, response-size, Web UI, and final validation counts. |

Final validation:

| Command | Result |
|---|---|
| `powershell -ExecutionPolicy Bypass -File anvien-launcher\build.ps1` | Pass after implementation code. Existing Vite dynamic-import and chunk-size warnings only. |
| `go test ./cmd/... ./internal/... -count=1` | Pass. Covers backend, CLI, API, MCP, contracts, AI context, graph quality, and file-context packages. |
| `go run .\cmd\generate-web-contracts --check` | Pass. |
| `npm --prefix anvien-web test` | Pass. 52 files, 408 tests. |
| `npm --prefix anvien-web run test:e2e -- shell-interactions.spec.ts -g "file map"` with `E2E=1`, `E2E_REPO_NAME=Anvien` | Pass. File Map opens File Detail and validates major sections. |
| Final `.\anvien\bin\anvien.exe analyze --force --name Anvien` | Pass. `files=831`, `parsed=596`, `nodes=95887`, `relationships=131223`; file projection `dependencyEdges=15816`, `unresolvedFiles=588`. |
| File-context representative smokes | Pass for source, test, docs/config-style, generated, API/MCP, and unresolved-heavy files. |
| File-hotspots sort smokes | Pass for `unresolved`, `fan-in`, `fan-out`, `symbols`, `flows`, `tests`, and `path`. |
| Parent/child command smokes | Pass for `query files`, `context symbol`, and `graph-health files`; file-aware command surfaces remain available. |

Runtime/UI availability:

- Backend was started at `http://127.0.0.1:4848` with `anvien serve`.
- Vite Web UI was started at `http://127.0.0.1:5228`.
- Browser plugin navigation was not exposed by tool discovery in this session; Playwright e2e is the recorded UI validation.

Final detect changes:

| Command | Result |
|---|---|
| `.\anvien\bin\anvien.exe detect-changes --repo Anvien --scope all --json` | Pass. `risk_level=low`; `changed_count=108`; `changed_files=8`; `affected_count=0`; `affected_files=10`; file layer `changedFiles=8`, `affectedFiles=10`, `changedFileRisk=high`; resolution gap changed entities `71`. |

## E17 - Retire Self-Referential AI Context Skill

Date: 2026-06-01

Status: completed

Scope:

- Remove only the self-referential `anvien-ai-context` skill from the embedded base skill registry.
- Keep `.claude/skills/anvien/**` as generated output for the remaining Anvien skills.
- Add cleanup coverage so stale generated `anvien-ai-context` directories are removed by normal analyze/setup generation.

Impact / blast radius:

| Command | Result |
|---|---|
| `.\anvien\bin\anvien.exe analyze --force --name Anvien` | Pass before edits. `files=831`, `parsed=596`, `nodes=95887`, `relationships=131223`; file projection `dependencyEdges=15816`, `unresolvedFiles=588`. |
| `.\anvien\bin\anvien.exe impact file "internal/aicontext/aicontext.go" --repo Anvien --direction upstream --include-tests` | HIGH/CRITICAL file-level blast radius through analyze/setup AI-context generation, so edits were scoped to the base skill registry, retired-skill cleanup, embedded source deletion, and tests. |
| `.\anvien\bin\anvien.exe impact --uid "Variable:internal/aicontext/aicontext.go:baseSkills" --repo Anvien --direction upstream --include-tests` | LOW symbol impact, `impactedCount=0`; linked flows/tests still show the containing generator file participates in analyze AI-context flows. |
| `.\anvien\bin\anvien.exe impact --uid "Variable:internal/aicontext/aicontext.go:retiredBaseSkillNames" --repo Anvien --direction upstream --include-tests` | LOW symbol impact, `impactedCount=0`. |

Implementation evidence:

| File | Evidence |
|---|---|
| `internal/aicontext/aicontext.go` | Removed `anvien-ai-context` from `baseSkills` and added it to `retiredBaseSkillNames` so normal generation removes stale output while preserving all other generated Anvien skills. |
| `internal/aicontext/skills/anvien-ai-context.md` | Deleted the embedded source of the self-referential skill. |
| `internal/aicontext/aicontext_test.go` | Expected base skill inventory now has 10 skills, asserts the generated root Skills table omits `anvien-ai-context`, and asserts the self-referential skill directory is not installed. |
| `internal/httpapi/analyze_test.go` | Hardened the existing analyze/embed lock lifecycle test timeout after full-suite Windows runs exposed a timing race while validating this slice. |

Regeneration / validation:

| Command | Result |
|---|---|
| `powershell -ExecutionPolicy Bypass -File anvien-launcher\build.ps1` | Initial run was blocked by an existing `anvien\bin\anvien.exe` process holding the binary. After stopping the workspace process, rerun passed; existing Vite dynamic-import/chunk-size warnings only. |
| `.\anvien\bin\anvien.exe analyze --force --name Anvien` | Pass with the rebuilt binary. `files=830`, `parsed=596`, `unsupported=234`, `failed=0`, `nodes=95885`, `relationships=131221`; file projection `dependencyEdges=15816`, `unresolvedFiles=588`. |
| Generated output checks | Pass. `.claude/skills/anvien/anvien-ai-context` and its `SKILL.md` no longer exist; generated `AGENTS.md` and `CLAUDE.md` no longer reference `anvien-ai-context`. |
| `go test ./internal/aicontext -count=1` | Pass. |
| `go test ./internal/httpapi -run TestAnalyzeLockBlocksEmbedAndReleasesAfterCancel -count=1 -v` | Pass after timeout hardening. |
| `go test ./cmd/... ./internal/... -count=1` | Pass after rerun. |
| `.\anvien\bin\anvien.exe detect-changes --repo Anvien --scope all --json` | Pass. `risk_level=low`; `changed_count=27`; `changed_files=6`; `affected_count=0`; file layer `changedFiles=6`, `affectedFiles=5`, `changedFileRisk=high`; resolution gap changed entities `11`. |

Notes:

- Web UI behavior did not change, so no Web e2e test was required for this slice.
