# Anvien GitNexus Deep Comparison Evidence Ledger

Date: 2026-06-01

Status: completed

Companion files:

- Plan: [2026-06-01-benchmark-anvien-gitnexus-deep-comparison-plan.md](2026-06-01-benchmark-anvien-gitnexus-deep-comparison-plan.md)
- Benchmark ledger: [2026-06-01-benchmark-anvien-gitnexus-deep-comparison-benchmark.md](2026-06-01-benchmark-anvien-gitnexus-deep-comparison-benchmark.md)
- Final report: [2026-06-01-benchmark-anvien-gitnexus-deep-comparison-report.md](2026-06-01-benchmark-anvien-gitnexus-deep-comparison-report.md)

## Evidence Rules

1. Record facts observed from source code, commands, generated graph/index output, or build/test results.
2. Do not record GitHub stars, popularity, social proof, or README claims as evaluation evidence.
3. Separate measured facts from interpretation.
4. Link or name exact files, symbols, commands, and output artifacts wherever possible.
5. Record failed commands exactly enough that the failure can be reproduced.
6. Keep quantitative benchmark tables in the benchmark ledger and use this ledger for command context and interpretation.
7. Record Anvien graph freshness before any graph-based Anvien evidence.
8. Record temporary clone cleanup before closing the comparison.
9. Record the GitNexus temporary clone path and verify it is outside `E:\Anvien` before cloning.

## E0 - Environment Inventory

Date: 2026-06-01

Status: completed

Commands to run:

```powershell
Get-ComputerInfo | Select-Object OsName,OsVersion,OsArchitecture,CsProcessors,CsTotalPhysicalMemory
git --version
go version
node --version
npm --version
python --version
```

Result:

| Field | Value |
|---|---|
| OS | Microsoft Windows 10 Pro 10.0.19045, 64-bit |
| CPU | Intel(R) Core(TM) i7-3770 CPU @ 3.40GHz, 8 logical processors |
| RAM | 33238466560 bytes |
| Git | git version 2.54.0.windows.1 |
| Go | go version go1.26.3 windows/amd64 |
| Node | v24.15.0 |
| npm | 11.12.1 |
| Python | Python 3.14.5 |

## E1 - Repository Commit Inventory

Date: 2026-06-01

Status: completed

Commands to run:

```powershell
$anvienRoot = (Resolve-Path "E:\Anvien").Path
$tempRoot = [System.IO.Path]::GetFullPath($root)
if ($tempRoot.StartsWith($anvienRoot, [System.StringComparison]::OrdinalIgnoreCase)) {
    throw "GitNexus temp clone path must be outside E:\Anvien"
}
git -C E:\Anvien rev-parse HEAD
git -C E:\Anvien remote -v
git -C $gitnexus rev-parse HEAD
git -C $gitnexus remote -v
```

Result:

| Repo | Local path | Remote | Commit |
|---|---|---|---|
| Anvien | `E:\Anvien` | `https://github.com/tamnguyendinh/Anvien.git` | `7b4d48d9bf44b5aa0c6f394861a7d356929521cb` |
| GitNexus | `%TEMP%\anvien-gitnexus-comparison\GitNexus` | `https://github.com/abhigyanpatwari/GitNexus` | `ce7f45e18d8dceedbcecffad83e5ae23ca105149` |

Temporary clone safety:

| Check | Result |
|---|---|
| Temp root | `%TEMP%\anvien-gitnexus-comparison` |
| GitNexus clone path | `%TEMP%\anvien-gitnexus-comparison\GitNexus` |
| Outside `E:\Anvien` | true |

Clean benchmark targets:

| Target | Local path | Commit | Reason |
|---|---|---|---|
| AnvienTarget | `%TEMP%\anvien-gitnexus-comparison\AnvienTarget` | `7b4d48d9bf44b5aa0c6f394861a7d356929521cb` | Prevent GitNexus from writing `.gitnexus` into `E:\Anvien`; gives a clean comparable Anvien target. |
| GitNexusTarget | `%TEMP%\anvien-gitnexus-comparison\GitNexusTarget` | `ce7f45e18d8dceedbcecffad83e5ae23ca105149` | Prevent GitNexus from scanning Anvien-generated `.anvien` output in the build clone. |

Note:

- First local AnvienTarget clone attempt was blocked by Git safe-directory ownership checks. Retry succeeded with command-local `-c safe.directory=...`; no global Git config was changed.

## E2 - Anvien Graph Freshness

Date: 2026-06-01

Status: completed

Command:

```powershell
.\anvien\bin\anvien.exe analyze --force --name Anvien --json --benchmark-json $env:TEMP\anvien-gitnexus-comparison\anvien-self-analyze.json
```

Result:

| Field | Value |
|---|---:|
| Elapsed seconds | 43.7646606 |
| Files scanned | 813 |
| Files parsed | 596 |
| Files failed | 0 |
| Current aggregate unsupported | 217 |
| Graph nodes | 95889 |
| Graph relationships | 131238 |
| Execution flows | 700 on clean Anvien target |

Interpretation:

- The graph refresh succeeded with no failed files.
- Counts differ from the pre-plan baseline because the three GitNexus comparison plan files are now part of the repo inventory.

Clean Anvien target benchmark:

```powershell
.\anvien\bin\anvien.exe analyze $env:TEMP\anvien-gitnexus-comparison\AnvienTarget --force --name AnvienTargetBenchmark --allow-duplicate-name --json --benchmark-json $env:TEMP\anvien-gitnexus-comparison\anvien-anvientarget-analyze.json
.\anvien\bin\anvien.exe graph-health summary --repo AnvienTargetBenchmark --json
```

| Field | Value |
|---|---:|
| Elapsed seconds | 41.5269639 |
| Files scanned | 810 |
| Files parsed | 596 |
| Files failed | 0 |
| Current aggregate unsupported | 214 |
| Graph nodes | 95845 |
| Graph relationships | 131188 |
| Resolved references | 31584 |
| Unresolved references | 69807 |
| Source-backed unresolved references | 69807 |
| In-repo analyzer gaps | 41091 |
| Graph snapshot size | 326917002 bytes |
| `.anvien` directory size | 464231665 bytes |

Additional Anvien benchmark target:

```powershell
.\anvien\bin\anvien.exe analyze $env:TEMP\anvien-gitnexus-comparison\GitNexus --force --name GitNexusBenchmark --allow-duplicate-name --json --benchmark-json $env:TEMP\anvien-gitnexus-comparison\anvien-gitnexus-analyze.json
.\anvien\bin\anvien.exe graph-health summary --repo GitNexusBenchmark --json
.\anvien\bin\anvien.exe source-site-accuracy --graph $env:TEMP\anvien-gitnexus-comparison\GitNexus\.anvien\graph.json
```

Result:

| Field | Value |
|---|---:|
| Target repo | GitNexus temp clone |
| Elapsed seconds | 85.377302 |
| Files scanned | 1339 |
| Files parsed | 1221 |
| Files failed | 0 |
| Current aggregate unsupported | 118 |
| Graph nodes | 225455 |
| Graph relationships | 245957 |
| Resolved references | 35247 |
| Unresolved references | 191224 |
| Source-backed unresolved references | 191224 |
| In-repo analyzer gaps | 178924 |
| False-resolved edge candidates | 0 |
| Resolved edges without source-site proof | 0 |
| Graph snapshot size | 823484355 bytes |
| `.anvien` directory size | 1075507495 bytes |

Interpretation:

- Anvien successfully analyzed the full GitNexus monorepo clone without failed files.
- The GitNexus target produced a much larger graph than Anvien's own repo on this machine: 225,455 nodes and 245,957 relationships.
- Source-site auditing reported no false-resolved edge candidates and no resolved edges missing source-site proof, but unresolved volume was high: 191,224 occurrences, including 178,924 in-repo analyzer gaps.

## E3 - Anvien Architecture Evidence

Date: 2026-06-01

Status: completed

Evidence targets:

| Area | Files/symbols/commands | Findings |
|---|---|---|
| Scanner and analyzer pipeline | `internal/analyze/analyze.go`, `internal/scanner/language.go`, `internal/scanner/registry_primary.go` | `analyze.Run` is a staged pipeline: scan, structure, documents, COBOL, parse, routes, tools, ORM, cross-file binding, resolution, MRO, communities, processes, semantic enrichment, DB load, optional embeddings. Scanner maps code/document extensions and tracks primary languages. |
| Parser/extractor model | `internal/parser/registry.go`, `internal/providers/*`, `internal/analyze/analyze.go:parseFiles` | Tree-sitter grammars cover JavaScript, TypeScript/TSX, Go, Python, Java, Kotlin, C, C#, C++, Rust, PHP, Dart, Swift, Ruby. Extractor layer additionally handles Vue/Svelte/Astro script containers. COBOL is handled as a separate phase. |
| Resolver and relationship builder | `internal/resolution/resolve.go`, `internal/resolution/indexes.go`, `internal/resolution/import_resolution.go`, `internal/graph/types.go` | Resolver builds a cross-file workspace, emits definitions, import edges, inheritance, calls, accesses, type annotations, method dispatch, and source-backed `ResolutionGap` diagnostics. Relationship types include `CALLS`, `IMPORTS`, `USES`, `ACCESSES`, `HAS_METHOD`, `HAS_PROPERTY`, route/tool edges, process edges, and resolution gaps. |
| Graph schema and persistence | `internal/graph/types.go`, `.anvien/graph.json`, `internal/lbugload`, `internal/lbugruntime` | Canonical graph is JSON snapshot plus LadybugDB-native runtime load. Nodes carry labels/properties; relationships carry confidence, reason, source-site proof, source range, proof kind, role, and target text. |
| Query and context commands | `internal/cli/tool_command.go`, `internal/mcp/tools.go`, `internal/mcp/context.go`, `internal/mcp/resources.go` | CLI query/context/cypher delegate to local MCP runtime. Query has lanes for owner, concept, execution flow, API surface, graph quality, docs/setup/AI context, command surface, and cross-repo discovery. |
| Impact analysis | `internal/mcp/impact.go`, `internal/mcp/detect_changes.go`, `internal/cli/tool_command.go` | Impact uses BFS over selected semantic relationship types and reports affected files/processes/risk. Detect-changes parses git hunks, maps changed symbols/files, affected flows, app layers, functional areas, and resolution health effects. |
| API/Web surfaces | `internal/httpapi/server.go`, `anvien-web/src`, `contracts/web-ui` | HTTP API exposes repo, graph, graph health explain/report, file context, grep, query, processes/clusters, analyze jobs, search/embed, MCP-over-HTTP, and session routes. Web UI is React/Vite/Sigma/Graphology over local HTTP API. |
| MCP/agent integration | `internal/mcp/server.go`, `internal/mcp/tools.go`, `internal/mcp/resources.go`, `internal/mcp/prompts.go` | MCP stdio implements initialize, tools, resources, resource templates, and prompts. Tools include list/query/cypher/context/impact/detect-changes/rename plus API/group surfaces. |
| Graph health and diagnostics | `internal/graphhealth/compute.go`, `internal/cli/graph_health_command.go`, `internal/graphaccuracy/*` | Graph health computes topology, expected isolation, components, resolution-health buckets, source-site diagnostics, and actionable/non-actionable ResolutionGap classifications. |
| Benchmarks and reports | `internal/analyze/analyze.go:WriteBenchmark`, `internal/cli/benchmark_command.go`, `cmd/graph-accuracy-probe`, `internal/lbugload/benchmark_test.go` | Analyze can emit benchmark JSON. Dedicated commands compare benchmark artifacts, report source-site accuracy, and inspect resolution inventory. |
| Tests | `rg --files -g '*_test.go'`; `anvien-web/package.json` | 174 Go test files and 65 Web unit/e2e test/spec files were present at measurement time. Web scripts include unit tests, coverage, and Playwright e2e. |
| Packaging/runtime | `anvien-launcher/build.ps1`, `anvien-launcher`, `internal/cli/package_*`, `internal/httpapi/listen.go` | Full build produced `anvien/bin/anvien.exe` version 1.2.4 and Web dist. Launcher/runtime packaging includes Windows launcher and native LadybugDB runtime handling. |

Graph-health evidence from the fresh Anvien index:

| Metric | Value |
|---|---:|
| Nodes | 95889 |
| Relationships | 131238 |
| Counted relationships | 27752 |
| Files in file layer | 813 |
| File unresolved count | 588 |
| Resolved references | 31584 |
| Unresolved references | 69807 |
| Source-backed unresolved references | 69807 |
| In-repo analyzer gaps | 41091 |
| Non-actionable unresolved diagnostics | 28422 |
| External-library review diagnostics | 294 |
| False-resolved edge candidates from source-site accuracy | 0 |
| Resolved edges without source-site proof | 0 |
| Graph snapshot size | 326974292 bytes |
| `.anvien` directory size | 464313456 bytes |

Interpretation:

- Anvien is a broad, multi-surface graph system rather than only a repository visualizer.
- The core strength is breadth of graph facts and agent-facing workflows: source-site proof, ResolutionGap diagnostics, graph health, impact analysis, detect-changes, file context, MCP resources, local HTTP, Web UI, and group support.
- The main accuracy risk visible in current evidence is unresolved-reference volume: 69,807 unresolved occurrences, including 41,091 classified as in-repo analyzer gaps.

## E4 - GitNexus Setup and Architecture Evidence

Date: 2026-06-01

Status: completed

Evidence targets:

| Area | Files/symbols/commands | Findings |
|---|---|---|
| Dependency manifests | root `package.json`, `package-lock.json`; `gitnexus/package.json`; `gitnexus-web/package.json`; `gitnexus-shared/tsconfig.json`; `eval/pyproject.toml` | GitNexus is primarily a Node/TypeScript monorepo with CLI/core in `gitnexus`, Web UI in `gitnexus-web`, shared scope-resolution helpers in `gitnexus-shared`, and Python eval tooling. |
| Build/install commands | `gitnexus/package.json`, `gitnexus/scripts/build.js`, `gitnexus-shared/package.json`, `gitnexus-web/package.json` | Core package script `build` runs `node scripts/build.js`, which compiles `gitnexus-shared`, compiles `gitnexus`, copies shared dist into `gitnexus/dist/_shared`, rewrites imports, makes CLI executable, builds Web UI, and copies Web dist into `gitnexus/web`. `npm ci` runs postinstall grammar materialization/build scripts plus prepare/prepack paths. |
| Scanner and analyzer pipeline | `gitnexus/src/core/run-analyze.ts`, `gitnexus/src/core/ingestion/pipeline.ts`, `gitnexus/src/core/ingestion/pipeline-phases/*` | `runFullAnalysis` owns end-to-end analyze, metadata, LadybugDB load, FTS, embeddings, registry, cache, and AI-context generation. Pipeline uses a typed phase DAG: scan, structure, markdown, cobol, parse, routes, tools, orm, crossFile, scopeResolution, mro, communities, processes. |
| Parser/extractor model | `gitnexus/src/core/tree-sitter/parser-loader.ts`, `gitnexus/src/core/ingestion/pipeline-phases/parse.ts`, `parse-impl.ts`, `gitnexus-shared/src/languages.ts` | Parser loader maps language keys to tree-sitter grammars. Required grammars include JS/TS/Python/Java/C#/C++/Go/Rust/PHP/Ruby/Vue; C is optional-severity error because of known ABI issues; Swift/Dart/Kotlin are optional. `parse-impl` chunks files by byte budget, uses worker pool or sequential path, parse cache, AST cache, and pre-extracted `ParsedFile` artifacts. |
| Resolver and relationship builder | `gitnexus/src/core/ingestion/call-processor.ts`, `import-processor.ts`, `scope-resolution/pipeline/phase.ts`, `cross-file.ts`, `model/*` | GitNexus has a legacy call-resolution DAG and a newer registry-primary scope-resolution path. Scope-resolution consumes the semantic model and emits IMPORTS/CALLS/ACCESSES/INHERITS/USES for migrated languages while legacy gates prevent duplicate edge emission. Cross-file propagation owns the binding accumulator lifecycle. |
| Graph schema and persistence | `gitnexus/src/core/graph/graph.ts`, `gitnexus/src/core/graph/types.ts`, `gitnexus/src/core/lbug/schema.ts`, `lbug-adapter.ts` | In-memory graph uses indexed maps for nodes, relationships, relationship type buckets, reverse adjacency, and file-to-node ids. Persistence uses LadybugDB with many node tables and one `CodeRelation` relationship table carrying `type`, confidence/reason/step fields. |
| Query surfaces | `gitnexus/src/cli/index.ts`, `gitnexus/src/cli/tool.ts`, `gitnexus/src/mcp/local/local-backend.ts`, `gitnexus/src/core/search/*` | CLI direct commands call LocalBackend. Query combines LadybugDB FTS/BM25 and semantic vectors with RRF where embeddings exist; FTS failures degrade to semantic-only rather than crashing. Cypher, context, impact, detect-changes, and rename are exposed as direct CLI and MCP tools. |
| API/Web surfaces | `gitnexus/src/server/api.ts`, `gitnexus-web/package.json`, `gitnexus-web/src` | Express server exposes local API, graph streaming, query/search, repo operations, analyze/embed jobs, process/cluster views, MCP-over-HTTP, and SPA serving. Web UI is Vite/React/Sigma/Graphology and also includes LangChain dependencies for chat-like features. |
| Diagnostics and graph health | `gitnexus/src/mcp/local/local-backend.ts`, `gitnexus/src/mcp/staleness.ts`, `gitnexus/src/storage/repo-manager.ts`, `gitnexus/src/core/lbug/*` | GitNexus has staleness checks, runtime doctor, WAL/sidecar recovery, read-only query error handling, FTS degradation flags, and query timing logs. It does not expose Anvien-style source-site accuracy, ResolutionGap inventory, or graph-health topology command in the inspected command surface. |
| Benchmarks and reports | `gitnexus/bench/*`, `gitnexus/scripts/bench-scope-resolution.ts`, `gitnexus/test/integration/*benchmark*.test.ts`, `gitnexus/bench/parse-throughput.md` | Repo has parse/scope benchmark scripts and benchmark tests, but the CLI analyze command does not expose an `--benchmark-json` flag comparable to Anvien. Runtime timing must be captured externally with `Measure-Command`. |
| Tests | `rg --files gitnexus/test -g '*.test.ts'`, `gitnexus-web/package.json`, `.github/workflows/*.yml` | 429 GitNexus TS test files and 34 Web test/e2e/spec files were present. CI workflow inventory includes tests, e2e, quality, report, scope parity, CodeQL, dependency review, Docker, trivy, scorecard, tree-sitter readiness, and publish workflows. |
| Packaging/runtime | `gitnexus/package.json`, `Dockerfile.cli`, `Dockerfile.web`, `docker-compose.yaml`, `gitnexus/scripts/build.js` | GitNexus is published as npm package `gitnexus` with bin `dist/cli/index.js`; build bundles Web UI into package. Dockerfiles exist for CLI/Web. |

GitNexus source inventory:

| Metric | Value |
|---|---:|
| Files under `gitnexus/` | 3059 |
| Core TS test files | 429 |
| Web test/e2e/spec files | 34 |
| GitHub workflow files | 21 |
| Supported language enum entries | 16 including Vue and standalone COBOL |

Interpretation:

- GitNexus is a mature TypeScript/Node implementation with heavy investment in worker-pool parsing, parse caching, LadybugDB persistence, hybrid search, group contracts, and agent integrations.
- It has broader test-file volume than Anvien in raw count, especially for integration and resolver coverage.
- Its inspected diagnostics focus on runtime reliability, staleness, search degradation, and DB recovery; it does not expose the same source-site proof and ResolutionGap audit surface that Anvien exposes.

## E5 - Command Surface Verification

Date: 2026-06-01

Status: completed

| Tool | Command | Result | Notes |
|---|---|---|---|
| Anvien | `powershell -ExecutionPolicy Bypass -File .\anvien-launcher\build.ps1` | pass | Built local Web UI and `anvien/bin/anvien.exe`. Vite emitted chunk-size and dynamic-import warnings only. |
| Anvien | `.\anvien\bin\anvien.exe version` | `1.2.4` | Local binary used for Anvien benchmark commands. |
| Anvien | `.\anvien\bin\anvien.exe --help` | 28 top-level commands | Includes analyze, API, query, context, impact, detect-changes, graph-health, source-site-accuracy, resolution-inventory, group, MCP, serve, package, doctor, setup. |
| Anvien | `.\anvien\bin\anvien.exe analyze --help` | pass | Supports `--benchmark-json`, `--json`, `--force`, embeddings, include/exclude, pprof profiles, and compatibility diagnostic flag. |
| GitNexus | `npm ci` in `gitnexus-shared` | pass in 2.1135874s | Installed TypeScript dependency for shared package build. |
| GitNexus | `npm ci` in `gitnexus` with `GITNEXUS_BUILD_TIMEOUT_MS=900000` | pass in 211.2149352s | Ran postinstall grammar materialization/build scripts and prepare build. Vite emitted chunk-size/dynamic-import warnings only. |
| GitNexus | `node .\dist\cli\index.js --version` | `1.6.5` | Built CLI used for benchmark commands. |
| GitNexus | `node .\dist\cli\index.js --help` | at least 20 top-level commands | Includes setup, analyze, index, serve, MCP, list, status, doctor, clean/remove, wiki, augment, publish, query, context, impact, cypher, detect-changes, eval-server, group. |
| GitNexus | `node .\dist\cli\index.js analyze --help` | pass | Supports `--force`, `--repair-fts`, embeddings, skills/AI context flags, `--index-only`, git/name flags, worker controls, max file size, WAL checkpoint threshold, and embedding controls. No `--benchmark-json` equivalent. |

Record command help snippets or summarized flags here after discovery.

GitNexus benchmark commands:

```powershell
node $cli analyze $env:TEMP\anvien-gitnexus-comparison\AnvienTarget --force --index-only --skip-agents-md --skip-skills --name GitNexusBenchAnvien --allow-duplicate-name
node $cli analyze $env:TEMP\anvien-gitnexus-comparison\GitNexusTarget --force --index-only --skip-agents-md --skip-skills --name GitNexusBenchSelf --allow-duplicate-name
```

Results:

| Tool | Target | Elapsed seconds | CLI reported seconds | Files | Nodes | Edges | Clusters | Flows | Index bytes | Result |
|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---|
| GitNexus | AnvienTarget | 69.3610866 | 68.0 | 809 | 23121 | 60428 | 598 | 300 | 240083880 | pass |
| GitNexus | GitNexusTarget | 227.4896233 | 225.6 | 1339 | 31622 | 50171 | 1167 | 300 | 339238569 | pass |

Observed capabilities in both GitNexus `meta.json` files:

- LadybugDB graph provider: available.
- LadybugDB FTS provider: available.
- Vector search: unavailable on this platform; exact-scan fallback would apply if embeddings existed.

## E6 - Accuracy Audit Method and Samples

Date: 2026-06-01

Status: completed

Sampling rule:

- Common target: `AnvienTarget` at commit `7b4d48d9bf44b5aa0c6f394861a7d356929521cb`.
- Source evidence came from `rg` against the clean target clone.
- Graph evidence came from `anvien cypher`/`anvien context` on `AnvienTargetBenchmark` and `gitnexus cypher`/`gitnexus context` on `GitNexusBenchAnvien`.
- This is a targeted deterministic audit, not an exhaustive language-by-language audit.

Sample inventory:

| Category | Target repo | Tool | Sample count | Selection method | Artifact |
|---|---|---|---:|---|---|
| File nodes | AnvienTarget | both | 10 | fixed paths across Go, TS, e2e, docs | `cypher` file-node lookup |
| Top-level symbols | AnvienTarget | both | 10 | declarations in analyzer, CLI, graphaccuracy, HTTP API | `rg`, `cypher` symbol lookup |
| Methods/functions | AnvienTarget | both | 6 | analyzer/API functions and methods | `rg`, `cypher`, `context` |
| Imports/dependencies | AnvienTarget | both | 5 | `internal/analyze/analyze.go` import targets | `rg`, `cypher` import-edge lookup |
| Calls/references | AnvienTarget | both | 10 | outgoing calls from `internal/analyze/analyze.go:Run` | `context`, `cypher` |
| Unresolved/gaps | AnvienTarget | Anvien | exhaustive graph-health/source-site counters | all source-site diagnostics | `graph-health`, `source-site-accuracy` |
| Unresolved/gaps | AnvienTarget | GitNexus | unsupported as equivalent command | none | `meta.json`, CLI help/source read |

Source facts checked:

| Fact group | Source evidence |
|---|---|
| `internal/analyze/analyze.go` declarations | `Run` line 196, `WriteBenchmark` line 762, `parseFiles` line 785, `currentAlloc` line 1015, `currentSys` line 1021. |
| Call sites into analyzer | `internal/cli/command.go` calls `analyze.Run` at line 225; `internal/graphaccuracy/access_candidate.go` calls `analyze.Run` at line 60. |
| Imports from analyzer | `context`, `runtime`, `internal/graph`, `internal/parser`, `internal/resolution`, `internal/scanner`. |

Graph facts checked:

| Tool | Evidence |
|---|---|
| Anvien | 10/10 selected file nodes found. Selected declarations found with 1-based source lines. `Run` had `CALLS` edges to `WriteBenchmark`, `parseFiles`, `runPhase`, `currentAlloc`, `currentSys`, `resolveDBRunner`, `prepareStorage`, `loadGraph`, `runEmbeddings`, and `writeGraphSnapshot`. |
| GitNexus | 10/10 selected file nodes found. Selected declarations found; CLI/Cypher reports `startLine` one less than source grep, consistent with a 0-based line convention. `Run` context and Cypher showed the same sampled outgoing call targets. |

## E7 - Accuracy Findings

Date: 2026-06-01

Status: completed

| Tool | Target repo | Category | Correct | Partial | Missing | False positives | Unsupported | Notes |
|---|---|---|---:|---:|---:|---:|---:|---|
| Anvien | AnvienTarget | File nodes | 10 | 0 | 0 | 0 | 0 | Selected paths all present with expected `File:` ids and file paths. |
| GitNexus | AnvienTarget | File nodes | 10 | 0 | 0 | 0 | 0 | Selected paths all present with expected `File:` ids and file paths. |
| Anvien | AnvienTarget | Top-level symbols | 10 | 0 | 0 | 0 | 0 | Selected declarations found; line numbers matched source grep. |
| GitNexus | AnvienTarget | Top-level symbols | 10 | 0 | 0 | 0 | 0 | Selected declarations found; line numbers are 0-based in CLI/Cypher output. |
| Anvien | AnvienTarget | Methods/functions | 6 | 0 | 0 | 0 | 0 | Selected function/method declarations and ownership matched source. |
| GitNexus | AnvienTarget | Methods/functions | 6 | 0 | 0 | 0 | 0 | Selected function/method declarations and ownership matched source. |
| Anvien | AnvienTarget | Imports/dependencies | 5 | 0 | 0 | 0 | 0 | Internal dependency file edges for sampled analyzer imports were present. |
| GitNexus | AnvienTarget | Imports/dependencies | 5 | 0 | 0 | 0 | 0 | Internal dependency file edges for sampled analyzer imports were present. |
| Anvien | AnvienTarget | Calls/references | 10 | 0 | 0 | 0 | 0 | Sampled `Run` outgoing calls were present; exhaustive source-site audit reported 0 false-resolved edge candidates. |
| GitNexus | AnvienTarget | Calls/references | 10 | 0 | 0 | 0 | 0 | Sampled `Run` outgoing calls were present; no equivalent false-positive audit command was exposed. |

Representative cases:

| Tool | Target repo | Case | Classification | Source evidence | Graph/index evidence |
|---|---|---|---|---|---|
| Anvien | AnvienTarget | `Run -> parseFiles` | correct | `internal/analyze/analyze.go` calls `parseFiles`; `parseFiles` declaration at line 785 | `CALLS` edge to `Function:internal/analyze/analyze.go:parseFiles#4`. |
| GitNexus | AnvienTarget | `Run -> parseFiles` | correct | same source fact | `CALLS` edge to `Function:internal/analyze/analyze.go:parseFiles`. |
| Anvien | AnvienTarget | source line for `Run` | correct | `rg` reports line 196 | `startLine` 196. |
| GitNexus | AnvienTarget | source line for `Run` | correct with convention note | `rg` reports line 196 | `startLine` 195, interpreted as 0-based. |
| Anvien | AnvienTarget | unresolved/gap observability | correct diagnostic surface | source-site audit covers all source-site diagnostics | 69,807 unresolved references; 41,091 in-repo analyzer gaps. |
| GitNexus | AnvienTarget | unresolved/gap observability | unsupported | no equivalent CLI/source-site command found | `meta.json` exposes aggregate graph stats but no unresolved/gap inventory. |

## E8 - Feature and Maturity Evidence

Date: 2026-06-01

Status: completed

| Dimension | Anvien evidence | GitNexus evidence | Interpretation |
|---|---|---|---|
| Language coverage | 17 parser/extractor languages plus COBOL and document indexing from source reads | 16 enum entries, required/optional grammar tiers, Vue and COBOL support | Both are broad. GitNexus makes optional grammar degradation explicit; Anvien has broader graph-diagnostic language surfaces. |
| CLI quality | 28 top-level commands; query/context/impact/cypher/graph-health/source-site/benchmark/group/setup/package | at least 20 top-level commands plus group, doctor, query/context/impact/cypher, eval-server | Both are mature CLI products. Anvien exposes more graph-quality commands; GitNexus direct Node CLI was faster for context/impact in the sample. |
| API/Web UI | 25 HTTP route handlers; 76 Web source files; graph/file/process/chat/local runtime panels | Express API plus SPA; 84 Web source files; language switcher/settings/help/chat-like surfaces | Both have usable local Web apps. GitNexus UI source count is slightly larger; Anvien has more graph-health/file-context-oriented API routes. |
| Graph query | Query lanes, Cypher, context, file-context, file-hotspots, graph health | Query, Cypher, context, impact through LocalBackend and LadybugDB | Both support runtime graph exploration. Anvien returns richer resolution metadata; GitNexus was lower latency in the small command sample. |
| Impact analysis | Impact plus detect-changes, affected flows/files/app layers/functional areas | Impact plus detect-changes, depth/limit/summary controls | Both present. Anvien reports richer app-layer/functional-area detail; GitNexus summary output is concise and quick. |
| Health diagnostics | graph-health, source-site-accuracy, resolution-inventory, query-health, benchmark-compare | doctor, staleness, WAL/sidecar recovery, FTS degradation, query timing | Anvien is stronger for graph correctness diagnostics. GitNexus is strong on runtime/storage operational diagnostics. |
| Multi-repo support | group create/add/status/sync/contracts/query | group create/add/status/sync/contracts/query in source and CLI | Comparable feature family; both are agent-oriented. |
| Agent/MCP integration | MCP stdio, MCP-over-HTTP, resources, prompts, setup, generated skills | MCP stdio/HTTP, setup for editors/agents, bundled skills/hooks | Comparable breadth. GitNexus has more visible multi-agent/plugin packaging assets; Anvien has richer MCP graph-quality prompts/resources. |
| Tests | 174 Go test files; 65 Web unit/e2e/spec files; 14 GitHub workflows | 429 core TS test files; 34 Web test/e2e/spec files; 22 GitHub workflows | GitNexus has larger raw core test inventory. Anvien has stronger Go/Web split plus e2e coverage. |
| Build reliability | Full Windows build passed in about 49.5s; local binary version 1.2.4 | `npm ci`/build passed, but cold install/build took 211.2s | Both build. GitNexus setup is materially slower on this machine. |
| Packaging/release | Windows launcher, local binary, package runtime helpers, Docker workflow | npm bin package, Dockerfiles, bundled Web dist, package scripts | Both are package-oriented. GitNexus has a standard npm distribution path; Anvien has stronger native launcher/runtime packaging. |
| Error handling | analyze locks, stale lock recovery, graph freshness, source-site diagnostics | detailed CLI recovery hints, OOM/native binding guidance, WAL/FTS recovery, vector fallback status | Both mature. GitNexus CLI recovery text is especially explicit; Anvien graph diagnostics are deeper. |
| Persistence/index clarity | JSON graph snapshot plus LadybugDB load/runtime | LadybugDB `.gitnexus/lbug`, parse-cache, `meta.json` with stats/capabilities | GitNexus `meta.json` is compact and useful; Anvien JSON graph is inspectable but much larger. |

## E9 - Failure and Unsupported Cases

Date: 2026-06-01

Status: completed

| Tool | Target repo | Command | Outcome | Technical implication |
|---|---|---|---|---|
| GitNexus | both | `analyze --benchmark-json` | not exposed | Analyze timing required external `Measure-Command`; no first-class benchmark artifact comparable to Anvien. |
| GitNexus | both | source-site accuracy / resolution inventory equivalent | not exposed | Relationship false-positive and unresolved-gap assessment is less auditable from CLI. |
| GitNexus | both | vector search capability | unavailable in `meta.json` on this platform | Graph and FTS worked; semantic vector lane unavailable unless platform support changes. |
| GitNexus | GitNexusTarget, AnvienTarget | analyze | pass | P10 had no analyze failure to report. |
| Git | AnvienTarget clone | `git clone --no-hardlinks E:\Anvien` | first attempt blocked by safe-directory ownership | Retried with command-local `-c safe.directory=...`; no global config change. |
| Both | both | warm analyze | not run | Cold full rebuild numbers are comparable; warm/incremental behavior was left out rather than mixing no-op/incremental semantics. |

## E10 - Temporary Clone Cleanup

Date: 2026-06-01

Status: completed

Command:

```powershell
Remove-Item -LiteralPath $root -Recurse -Force
Test-Path -LiteralPath $root
```

Result:

| Field | Value |
|---|---|
| Temp root | `C:\Users\TAM NGUYEN\AppData\Local\Temp\anvien-gitnexus-comparison` |
| Existed before cleanup | true |
| Exists after cleanup | false |
| Outside `E:\Anvien` | true |
| GitNexus registry cleanup | Removed benchmark-only entries `GitNexusBenchSelf`, `GitNexusBenchAnvien`, and `RestaurantManagerGitNexusBench`; final `~\.gitnexus\registry.json` content is `[]`. |
| Restaurant_manager temp cleanup | Removed `C:\rmbench`; removed GitNexus registry entry `RestaurantManagerGitNexusBench`; `C:\rmbench` exists after cleanup: false. |

## E11 - Final Report Evidence Map

Date: 2026-06-01

Status: completed

| Report section | Evidence sections | Benchmark sections |
|---|---|---|
| Methodology | E0, E1, E2, E5, E6 | B0 |
| Architecture | E3, E4 | B2, B3 |
| Performance | E5, E9 | B1, B2, B4 |
| Accuracy | E6, E7 | B5 |
| Functionality | E3, E4, E8 | B6 |
| Maturity | E3, E4, E8, E9 | B7 |
| Cleanup | E10 | none |

## E12 - Restaurant_manager Large Target Benchmark

Date: 2026-06-01

Status: completed

Target:

| Field | Value |
|---|---|
| Source repo | `E:\Restaurant_manager` |
| Commit | `fdfacba78e5445522dd09cca98fa27d39e0e22c8` |
| `rg --files` count | 6517 |
| Clean Anvien target | `C:\rmbench\rm-a` |
| Clean GitNexus target | `C:\rmbench\rm-g` |
| Clone safety | `C:\rmbench` is outside `E:\Anvien` and outside `E:\Restaurant_manager` |
| Windows path-length handling | First `%TEMP%` clone hit `Filename too long`; retry succeeded at short path `C:\rmbench` with command-local `core.longpaths=true`. |

Anvien command:

```powershell
E:\Anvien\anvien\bin\anvien.exe analyze C:\rmbench\rm-a --force --name RestaurantManagerAnvienBench --allow-duplicate-name --json --benchmark-json C:\rmbench\anvien-restaurant-manager-analyze.json
E:\Anvien\anvien\bin\anvien.exe graph-health summary --repo RestaurantManagerAnvienBench --json
E:\Anvien\anvien\bin\anvien.exe source-site-accuracy --graph C:\rmbench\rm-a\.anvien\graph.json
```

Anvien result:

| Field | Value |
|---|---:|
| Elapsed seconds | 79.8917692 |
| Files scanned | 6198 |
| Files parsed | 1228 |
| Files failed | 0 |
| Unsupported files | 4970 |
| Graph nodes | 202810 |
| Graph relationships | 253342 |
| Relationship types | 15 |
| Execution flows | 508 |
| Resolved references | 44493 |
| Unresolved references | 129135 |
| In-repo analyzer gaps | 80306 |
| False-resolved edge candidates | 0 |
| Resolved edges without source-site proof | 0 |
| Graph snapshot size | 657873587 bytes |
| `.anvien` directory size | 883305432 bytes |

GitNexus setup and command:

```powershell
git clone https://github.com/abhigyanpatwari/GitNexus C:\rmbench\GitNexus
npm ci # in C:\rmbench\GitNexus\gitnexus-shared, 2.59429s
npm ci # in C:\rmbench\GitNexus\gitnexus, 187.4168997s
node C:\rmbench\GitNexus\gitnexus\dist\cli\index.js analyze C:\rmbench\rm-g --force --index-only --skip-agents-md --skip-skills --name RestaurantManagerGitNexusBench --allow-duplicate-name
```

GitNexus result:

| Field | Value |
|---|---:|
| Elapsed seconds | 156.9328417 |
| CLI reported seconds | 155.7 |
| Files | 6198 |
| Nodes | 72792 |
| Edges | 143910 |
| Relationship types | 9 |
| Clusters | 1105 |
| Processes | 300 |
| Embeddings | 0 |
| Vector search | unavailable |
| `.gitnexus` index size | 643298557 bytes |

Interpretation:

- On this large target, Anvien was about 1.96x faster than GitNexus while producing 2.79x as many nodes and 1.76x as many relationships.
- The larger target confirms Anvien's cold analyze speed advantage, but it also shows the speed gap is workload-sensitive: GitNexus was 2.66x slower on the GitNexus target and 1.96x slower on Restaurant_manager.
- GitNexus produced more relationships than on the smaller GitNexus target and stayed below Anvien's graph volume.
