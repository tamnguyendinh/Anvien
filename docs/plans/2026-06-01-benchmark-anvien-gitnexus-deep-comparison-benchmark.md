# Anvien GitNexus Deep Comparison Benchmark Ledger

Date: 2026-06-01

Status: rerun completed after Anvien update

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
| Anvien commit | SHA | `97a45525820c609410796b1f11fa38239e31cbfa` |
| GitNexus commit | SHA | `ce7f45e18d8dceedbcecffad83e5ae23ca105149` |
| Restaurant_manager commit | SHA | `fdfacba78e5445522dd09cca98fa27d39e0e22c8` |
| Anvien build artifact size | bytes | 50898432 |
| GitNexus CLI entry artifact size | bytes | 12153 |
| GitNexus `dist` artifact size | bytes | 6884862 |

## B1 - Build and Install Metrics

Status: completed

| Tool | Command | Cold/warm | Elapsed seconds | Artifact size bytes | Result |
|---|---|---|---:|---:|---|
| Anvien | `powershell -ExecutionPolicy Bypass -File .\anvien-launcher\build.ps1` | cold | 42.996 | 50898432 | pass |
| GitNexus shared | `npm ci` in `E:\avgn-rerun\tools\GitNexus\gitnexus-shared` | cold | 3.230302 | n/a | pass |
| GitNexus core | `npm ci` in `E:\avgn-rerun\tools\GitNexus\gitnexus` | cold | 179.394383 | 6884862 | pass |

## B2 - Analyze Performance Matrix

Status: completed

| Tool | Target repo | Run | Elapsed seconds | Files scanned | Files parsed/indexed | Nodes | Relationships | Output size bytes | Result |
|---|---|---|---:|---:|---:|---:|---:|---:|---|
| Anvien | Anvien | cold | 37.985012 | 816 | 598 | 96211 | 131684 | 327984313 | measured on clean target `E:\avgn-rerun\targets\anvien-a`; benchmark JSON totalDuration 37.238526s |
| Anvien | Anvien | warm | n/a | n/a | n/a | n/a | n/a | n/a | not run |
| Anvien | GitNexus | cold | 100.934158 | 1339 | 1221 | 225455 | 245957 | 823484355 | measured on clean target `E:\avgn-rerun\targets\gitnexus-a`; benchmark JSON totalDuration 99.7532442s |
| Anvien | GitNexus | warm | n/a | n/a | n/a | n/a | n/a | n/a | not run |
| GitNexus | Anvien | cold | 117.262267 | 815 | 815 | 23264 | 60687 | 240978915 | measured on clean target `E:\avgn-rerun\targets\anvien-g`; CLI reported 113.8s |
| GitNexus | Anvien | warm | n/a | n/a | n/a | n/a | n/a | n/a | not run |
| GitNexus | GitNexus | cold | 215.510712 | 1339 | 1339 | 31622 | 50171 | 339209809 | measured on clean target `E:\avgn-rerun\targets\gitnexus-g`; CLI reported 210.3s |
| GitNexus | GitNexus | warm | n/a | n/a | n/a | n/a | n/a | n/a | not run |
| Anvien | Restaurant_manager | cold | 93.133052 | 6198 | 1228 | 202810 | 253342 | 657873587 | measured on clean target `E:\avgn-rerun\targets\rm-a`; benchmark JSON totalDuration 91.4707017s |
| GitNexus | Restaurant_manager | cold | 173.754737 | 6198 | 6198 | 72792 | 143910 | 641866383 | measured on clean target `E:\avgn-rerun\targets\rm-g`; CLI reported 171.6s |

## B3 - Graph Inventory Comparison

Status: completed

| Tool | Target repo | Files | Symbols | Relationships | Relationship types | Unresolved/gaps | Execution flows | Languages | Notes |
|---|---|---:|---:|---:|---:|---:|---:|---:|---|
| Anvien | Anvien | 816 | 96211 nodes | 131684 | 14 | 69988 | 700 | see B6 | Clean target. Source-site accuracy: 31713 graph-health resolved refs, 0 false-resolved candidates, 0 resolved edges without source-site proof. |
| Anvien | GitNexus | 1339 | 225455 nodes | 245957 | 21 | 191224 | 381 | see B6 | Source-site accuracy: 35247 resolved refs, 0 false-resolved candidates, 0 resolved edges without source-site proof. |
| GitNexus | Anvien | 815 | 23264 nodes | 60687 | 10 | not exposed | 300 | see B6 | 600 clusters; embeddings 0; vector search unavailable on this platform. |
| GitNexus | GitNexus | 1339 | 31622 nodes | 50171 | 13 | not exposed | 300 | see B6 | 1167 clusters; embeddings 0; vector search unavailable on this platform. |
| Anvien | Restaurant_manager | 6198 | 202810 nodes | 253342 | 15 | 129135 | 508 | see B6 | Source-site accuracy: 44493 resolved refs, 0 false-resolved candidates, 0 resolved edges without source-site proof. |
| GitNexus | Restaurant_manager | 6198 | 72792 nodes | 143910 | 9 | not exposed | 300 | see B6 | 1105 clusters; embeddings 0; vector search unavailable on this platform. |

## B4 - Query and Diagnostic Performance

Status: completed

| Tool | Target repo | Operation | Query/command | Elapsed seconds | Result count | Result |
|---|---|---|---|---:|---:|---|
| Anvien | Anvien | graph health | `graph-health summary --repo BenchRerunAnvienTarget --json` | 7.9293809 | 96211 nodes / 69988 unresolved refs | measured |
| Anvien | Anvien | concept query | `query "analyze pipeline" --repo BenchRerunAnvienTarget --json` | 7.5428505 | JSON query result | measured |
| Anvien | Anvien | symbol context | `context symbol "Run" --uid Function:internal/analyze/analyze.go:Run#3` | 7.8183634 | JSON context result | measured |
| Anvien | Anvien | impact analysis | `impact symbol "Run" --uid Function:internal/analyze/analyze.go:Run#3 --direction upstream` | 10.8241073 | JSON impact result | measured |
| GitNexus | Anvien | equivalent graph health | no equivalent graph-health/source-site command | n/a | n/a | not exposed |
| GitNexus | Anvien | concept query | `query "analyze pipeline" -r BenchRerunGitNexusOnAnvien --limit 5` | 4.73722 | text query result | measured |
| GitNexus | Anvien | symbol context | `context Run -r BenchRerunGitNexusOnAnvien -f internal/analyze/analyze.go` | 2.9433826 | 1 symbol with incoming/outgoing calls | measured |
| GitNexus | Anvien | impact analysis | `impact Run -r BenchRerunGitNexusOnAnvien -f internal/analyze/analyze.go --summary-only` | 2.3872532 | summary output | measured |

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
| Analyze speed | Anvien | 37.99s vs 117.26s on Anvien; 100.93s vs 215.51s on GitNexus; 93.13s vs 173.75s on Restaurant_manager | high | Cold full rebuild wall-clock runs only. |
| Graph completeness | Anvien | More nodes/relationships because the graph model is more detailed, plus ResolutionGap/source-site inventory | high | The higher count is interpreted as richer modeled detail, not duplicate inflation, because Anvien also exposes source-site proof and false-positive audit output. |
| Relationship accuracy | Anvien | Exhaustive source-site audit available; reduced sample tied at 100 percent | medium | GitNexus passed the reduced sample but lacks equivalent false-positive audit surface. |
| Feature breadth | Anvien | Stronger graph-health, benchmark, source-site, file-context surfaces | high | GitNexus is still broad and close on CLI/Web/MCP/groups. |
| Operational maturity | tie | Different strengths | medium | GitNexus has strong Node/npm/CI/runtime recovery; Anvien has faster build/analyze and stronger graph diagnostics. |
| Developer usability | mixed | GitNexus faster for sampled context/impact; Anvien richer diagnostics | medium | Choose by workflow: fast lookup vs graph-quality/impact evidence. |
| Best ideas for Anvien | n/a | n/a | high | Borrow ideas conceptually: parse-cache/worker ergonomics, compact `meta.json`, explicit capability degradation, concise CLI summary output. |
