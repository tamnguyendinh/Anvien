# Anvien GitNexus Deep Comparison Benchmark Ledger

Date: 2026-06-01

Status: completed

Companion files:

- Plan: [2026-06-01-benchmark-anvien-gitnexus-deep-comparison-plan.md](2026-06-01-benchmark-anvien-gitnexus-deep-comparison-plan.md)
- Evidence ledger: [2026-06-01-benchmark-anvien-gitnexus-deep-comparison-evidence.md](2026-06-01-benchmark-anvien-gitnexus-deep-comparison-evidence.md)
- Final report: [2026-06-01-benchmark-anvien-gitnexus-deep-comparison-report.md](2026-06-01-benchmark-anvien-gitnexus-deep-comparison-report.md)

## Benchmark Rules

1. Record quantitative data only in benchmark tables.
2. Put command context, interpretation, and failure explanations in the evidence ledger.
3. Use the same machine and pinned commits for all comparable runs.
4. Distinguish `measured`, `unsupported`, `failed`, `not exposed`, and `not run`.
5. Do not infer graph counts from README claims.
6. Record cold and warm runs separately when possible.
7. Keep raw output artifacts in a stable report path if they are generated during execution.
8. Keep benchmark clones and generated GitNexus outputs outside `E:\Anvien` unless a later step explicitly copies final report artifacts into `docs\plans`.

## B0 - Environment Metrics

Status: completed

| Metric | Unit | Value |
|---|---:|---:|
| CPU logical processors | count | 8 |
| Physical memory | bytes | 33238466560 |
| Anvien commit | SHA | `7b4d48d9bf44b5aa0c6f394861a7d356929521cb` |
| GitNexus commit | SHA | `ce7f45e18d8dceedbcecffad83e5ae23ca105149` |
| Anvien build artifact size | bytes | 50882560 |
| GitNexus CLI entry artifact size | bytes | 12153 |
| GitNexus `dist` artifact size | bytes | 6884862 |

## B1 - Build and Install Metrics

Status: completed

| Tool | Command | Cold/warm | Elapsed seconds | Artifact size bytes | Result |
|---|---|---|---:|---:|---|
| Anvien | `powershell -ExecutionPolicy Bypass -File .\anvien-launcher\build.ps1` | cold | 49.5 observed wall time from command output/session | 50882560 | pass |
| GitNexus shared | `npm ci` in `gitnexus-shared` | cold | 2.1135874 | n/a | pass |
| GitNexus core | `npm ci` in `gitnexus` | cold | 211.2149352 | 6884862 | pass |
| GitNexus shared | `npm ci` in `C:\rmbench\GitNexus\gitnexus-shared` | cold rerun for Restaurant_manager benchmark | 2.59429 | n/a | pass |
| GitNexus core | `npm ci` in `C:\rmbench\GitNexus\gitnexus` | cold rerun for Restaurant_manager benchmark | 187.4168997 | n/a | pass |

## B2 - Analyze Performance Matrix

Status: completed

| Tool | Target repo | Run | Elapsed seconds | Files scanned | Files parsed/indexed | Nodes | Relationships | Output size bytes | Result |
|---|---|---|---:|---:|---:|---:|---:|---:|---|
| Anvien | Anvien | cold | 41.5269639 | 810 | 596 | 95845 | 131188 | 326917002 | measured on clean temp target |
| Anvien | Anvien | warm | n/a | n/a | n/a | n/a | n/a | n/a | not run |
| Anvien | GitNexus | cold | 85.377302 | 1339 | 1221 | 225455 | 245957 | 823484355 | measured |
| Anvien | GitNexus | warm | n/a | n/a | n/a | n/a | n/a | n/a | not run |
| GitNexus | Anvien | cold | 69.3610866 | 809 | 809 | 23121 | 60428 | 240083880 | measured on clean temp target |
| GitNexus | Anvien | warm | n/a | n/a | n/a | n/a | n/a | n/a | not run |
| GitNexus | GitNexus | cold | 227.4896233 | 1339 | 1339 | 31622 | 50171 | 339238569 | measured on clean temp target |
| GitNexus | GitNexus | warm | n/a | n/a | n/a | n/a | n/a | n/a | not run |
| Anvien | Restaurant_manager | cold | 79.8917692 | 6198 | 1228 | 202810 | 253342 | 657873587 | measured on clean target `C:\rmbench\rm-a` |
| GitNexus | Restaurant_manager | cold | 156.9328417 | 6198 | 6198 | 72792 | 143910 | 643298557 | measured on clean target `C:\rmbench\rm-g` |

## B3 - Graph Inventory Comparison

Status: completed

| Tool | Target repo | Files | Symbols | Relationships | Relationship types | Unresolved/gaps | Execution flows | Languages | Notes |
|---|---|---:|---:|---:|---:|---:|---:|---:|---|
| Anvien | Anvien | 810 | 95845 nodes | 131188 | 14 | 69807 | 700 | see B6 | Clean target. Relationship types led by HAS_RESOLUTION_GAP, DEFINES, CALLS, USES, MEMBER_OF, IMPORTS. |
| Anvien | GitNexus | 1339 | 225455 nodes | 245957 | 21 | 191224 | 381 | see B6 | Source-site accuracy: 35247 resolved refs, 0 false-resolved candidates, 0 resolved edges without source-site proof. |
| GitNexus | Anvien | 809 | 23121 nodes | 60428 | 10 | not exposed | 300 | see B6 | 598 clusters; embeddings 0; vector search unavailable on this platform. |
| GitNexus | GitNexus | 1339 | 31622 nodes | 50171 | 13 | not exposed | 300 | see B6 | 1167 clusters; embeddings 0; vector search unavailable on this platform. |
| Anvien | Restaurant_manager | 6198 | 202810 nodes | 253342 | 15 | 129135 | 508 | see B6 | Source-site accuracy: 44493 resolved refs, 0 false-resolved candidates, 0 resolved edges without source-site proof. |
| GitNexus | Restaurant_manager | 6198 | 72792 nodes | 143910 | 9 | not exposed | 300 | see B6 | 1105 clusters; embeddings 0; vector search unavailable on this platform. |

## B4 - Query and Diagnostic Performance

Status: completed

| Tool | Target repo | Operation | Query/command | Elapsed seconds | Result count | Result |
|---|---|---|---|---:|---:|---|
| Anvien | Anvien | graph health | `graph-health summary --repo AnvienTargetBenchmark --json` | 7.9661617 | 95845 nodes / 69807 unresolved refs | measured |
| Anvien | Anvien | concept query | `query "analyze pipeline" --repo AnvienTargetBenchmark --json` | 7.6748876 | 5 files, 5 processes, ranked definitions | measured |
| Anvien | Anvien | symbol context | `context symbol "Run" --uid Function:internal/analyze/analyze.go:Run#3` | 7.8181696 | 16 flows, 5 linked tests in sample output | measured |
| Anvien | Anvien | impact analysis | `impact symbol "Run" --uid Function:internal/analyze/analyze.go:Run#3 --direction upstream` | 7.5394923 | 4 affected files in sample output | measured |
| GitNexus | Anvien | equivalent graph health | no equivalent graph-health/source-site command | n/a | n/a | not exposed |
| GitNexus | Anvien | concept query | `query "analyze pipeline" -r GitNexusBenchAnvien --limit 5` | 4.1021882 | 5 processes plus definitions | measured |
| GitNexus | Anvien | symbol context | `context Run -r GitNexusBenchAnvien -f internal/analyze/analyze.go` | 2.6881255 | 1 symbol with incoming/outgoing calls | measured |
| GitNexus | Anvien | impact analysis | `impact Run -r GitNexusBenchAnvien -f internal/analyze/analyze.go --summary-only` | 2.9984849 | impactedCount 5, risk LOW | measured |

## B5 - Accuracy Audit Scores

Status: completed

| Tool | Target repo | Category | Samples | Correct | Partial | Missing | False positives | Unsupported | Accuracy percent |
|---|---|---|---:|---:|---:|---:|---:|---:|---:|
| Anvien | AnvienTarget | File nodes | 10 | 10 | 0 | 0 | 0 | 0 | 100 |
| GitNexus | AnvienTarget | File nodes | 10 | 10 | 0 | 0 | 0 | 0 | 100 |
| Anvien | AnvienTarget | Top-level symbols | 10 | 10 | 0 | 0 | 0 | 0 | 100 |
| GitNexus | AnvienTarget | Top-level symbols | 10 | 10 | 0 | 0 | 0 | 0 | 100 |
| Anvien | AnvienTarget | Methods/functions | 6 | 6 | 0 | 0 | 0 | 0 | 100 |
| GitNexus | AnvienTarget | Methods/functions | 6 | 6 | 0 | 0 | 0 | 0 | 100 |
| Anvien | AnvienTarget | Imports/dependencies | 5 | 5 | 0 | 0 | 0 | 0 | 100 |
| GitNexus | AnvienTarget | Imports/dependencies | 5 | 5 | 0 | 0 | 0 | 0 | 100 |
| Anvien | AnvienTarget | Calls/references | 10 | 10 | 0 | 0 | 0 | 0 | 100 |
| GitNexus | AnvienTarget | Calls/references | 10 | 10 | 0 | 0 | 0 | 0 | 100 |

Accuracy percent formula:

```text
accuracy_percent = (correct + 0.5 * partial) / samples * 100
```

Unsupported samples remain visible and should not be silently dropped from the denominator unless the final report explicitly separates capability coverage from accuracy within claimed scope.

## B6 - Feature Coverage Matrix

Status: completed

| Capability | Unit | Anvien | GitNexus | Evidence section |
|---|---:|---|---|---|
| Supported languages | count/list | 17 parser/extractor languages plus COBOL phase and document indexing | 16 enum entries including Vue and standalone COBOL; optional Swift/Dart/Kotlin and optional-degradation C parser handling | E3, E4, E8 |
| CLI commands | count/list | 28 top-level commands | at least 20 top-level commands plus group subcommands from Commander source | E3, E4, E5 |
| API routes | count/list | 25 HTTP route handlers registered in `internal/httpapi/server.go` | at least 22 Express API routes plus MCP-over-HTTP and SPA routes | E3, E4, E8 |
| Web UI source files | count/list | 76 under `anvien-web/src`; components include graph canvas, file tree/map/detail, process modal/panel, chat, query FAB, analyzer onboarding/progress | 84 under `gitnexus-web/src`; components include graph canvas, file tree, process panel/modal, help, language switcher, settings, query FAB | E8 |
| Graph query support | yes/no/details | CLI/MCP query, cypher, context, file-context, file-hotspots | CLI/MCP query, cypher, context, hybrid FTS/semantic search path | E3, E8 |
| Impact analysis | yes/no/details | CLI/MCP impact plus detect-changes | CLI/MCP impact plus detect-changes, summary/depth/limit controls | E3, E8 |
| Graph health diagnostics | yes/no/details | graph-health, source-site-accuracy, resolution-inventory, query-health | doctor, staleness, FTS/vector degradation and DB recovery; no source-site graph-health equivalent found | E3, E8, E9 |
| Benchmark output | yes/no/details | analyze `--benchmark-json`, benchmark-compare | benchmark scripts/tests exist, but analyze CLI has no `--benchmark-json` equivalent | E3, E8, E9 |
| Multi-repo support | yes/no/details | group create/add/status/sync/contracts/query | group create/add/status/sync/contracts/query | E3, E8 |
| Agent/MCP support | yes/no/details | MCP stdio, MCP-over-HTTP, resources, prompts, setup | MCP stdio, MCP-over-HTTP, setup, skills/hooks | E3, E8 |
| Tests | count | 174 Go test files; 65 Web test/e2e/spec files | 429 core TS test files; 34 Web test/e2e/spec files | E3, E4, E8 |
| CI workflows | count | 14 GitHub workflow files | 22 GitHub workflow files | E4, E8 |
| Release/package artifacts | count/list | Windows launcher/build script, local binary, packaged runtime path | npm package bin, package build script, Dockerfiles, Web dist bundling | E3, E4, E8 |

## B7 - Maturity Scoring Worksheet

Status: completed

Use scores only after evidence is recorded. A score without a source-backed note is invalid.

Scale:

| Score | Meaning |
|---:|---|
| 0 | Not present or not runnable. |
| 1 | Prototype-level, fragile, or mostly manual. |
| 2 | Works for narrow cases with visible gaps. |
| 3 | Usable with tests or diagnostics, but incomplete. |
| 4 | Production-leaning, repeatable, documented, and tested. |
| 5 | Mature, robust, well-tested, observable, and packaged. |

| Dimension | Anvien score | GitNexus score | Evidence section |
|---|---:|---:|---|
| Analyzer architecture | 4 | 4 | E3, E4 |
| Resolver accuracy design | 4 | 3 | E3, E4, E7 |
| Runtime reliability | 4 | 4 | E5, E9 |
| Test quality | 4 | 4 | E8 |
| Build and packaging | 4 | 4 | E5, E8 |
| Diagnostics and observability | 5 | 3 | E8 |
| Query usability | 4 | 4 | E5, E8 |
| Documentation quality | 3 | 3 | E8 |
| Cross-platform readiness | 4 | 3 | E5, E9 |
| Maintainability | 4 | 4 | E3, E4 |

## B8 - Final Summary Metrics

Status: completed

| Category | Winner | Margin | Confidence | Notes |
|---|---|---|---|---|
| Analyze speed | Anvien | 41.53s vs 69.36s on Anvien; 85.38s vs 227.49s on GitNexus; 79.89s vs 156.93s on Restaurant_manager | high | Cold full rebuilds only. |
| Graph completeness | Anvien | More nodes/relationships and exposes ResolutionGap/source-site inventory | medium | Higher graph volume is useful only with accuracy diagnostics, which Anvien exposes. |
| Relationship accuracy | Anvien | Exhaustive source-site audit available; reduced sample tied at 100 percent | medium | GitNexus passed the reduced sample but lacks equivalent false-positive audit surface. |
| Feature breadth | Anvien | Stronger graph-health, benchmark, source-site, file-context surfaces | high | GitNexus is still broad and close on CLI/Web/MCP/groups. |
| Operational maturity | tie | Different strengths | medium | GitNexus has strong Node/npm/CI/runtime recovery; Anvien has faster build/analyze and stronger graph diagnostics. |
| Developer usability | mixed | GitNexus faster for sampled context/impact; Anvien richer diagnostics | medium | Choose by workflow: fast lookup vs graph-quality/impact evidence. |
| Best ideas for Anvien | n/a | n/a | high | Borrow ideas conceptually: parse-cache/worker ergonomics, compact `meta.json`, explicit capability degradation, concise CLI summary output. |
