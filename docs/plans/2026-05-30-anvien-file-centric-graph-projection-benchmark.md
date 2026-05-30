# Anvien File-Centric Graph Projection Benchmark Ledger

Date: 2026-05-30

Status: In progress

Companion files:

- Plan: [2026-05-30-anvien-file-centric-graph-projection-plan.md](2026-05-30-anvien-file-centric-graph-projection-plan.md)
- Evidence ledger: [2026-05-30-anvien-file-centric-graph-projection-evidence.md](2026-05-30-anvien-file-centric-graph-projection-evidence.md)

## Benchmark Rules

1. Record quantitative data only.
2. Put command pass/fail narrative in the evidence ledger.
3. Use units for every metric.
4. Preserve baseline, latest, target, and delta where useful.
5. Record graph inventory counts after graph-affecting implementation.
6. Record projection build/query timing and response sizes for CLI/API/Web-relevant outputs.
7. Record file inventory counts and hotspot counts because this feature is about file-level graph projection.
8. Build/test/e2e pass/fail belongs in evidence unless timing/count/size is the measured target.

## B0 - Graph Inventory Baseline

Status: P0-A baseline recorded

| Metric | Unit | Baseline | Latest | Delta | Target | Notes |
|---|---:|---:|---:|---:|---:|---|
| Files scanned | files | 819 | 829 | +10 | record | Latest after P5 parent/child command analyze refresh. |
| Files parsed | files | 584 | 594 | +10 | record | Latest after P5 parent/child command analyze refresh. |
| Unsupported files | files | 235 | 235 | 0 | record | Latest after P4-A Web File Map analyze benchmark. |
| Failed files | files | 0 | 0 | 0 | `0` | From readiness `anvien analyze --force --name Anvien`. |
| Graph nodes | nodes | 91587 | 94838 | +3251 | record | Current graph inventory after final P5 parent/child command analyze refresh. |
| Graph relationships | relationships | 125054 | 129861 | +4807 | record | Current graph inventory after final P5 parent/child command analyze refresh. |
| SourceSite count | distinct ids | 95433 | 99292 | +3859 | record | Distinct source-site ids from relationship trace fields after final P5 refresh. |
| ResolutionGap count | nodes | 65652 | 68035 | +2383 | record | Required for quality projection coverage after final P5 refresh. |

## B0A - Graph Generation Speed

Status: P4-A analyze benchmark recorded

| Metric | Unit | Baseline | Latest | Delta | Target | Notes |
|---|---:|---:|---:|---:|---:|---|
| Analyze benchmark runs | runs | pending | 4 | pending | record | Latest recorded run: `anvien analyze --force --name Anvien --benchmark-label file-centric-p4a-web-file-map --benchmark-json .tmp\file-centric-p4a-web-file-map-analyze-benchmark.json`. |
| Total graph generation time | ms | pending | 35958.2 | pending | track | Includes scan, parse, resolution, semantic enrichment, and DB load. |
| Scan phase time | ms | pending | 166.2 | pending | track | Benchmark phase `scan`. |
| Parse phase time | ms | pending | 8817.1 | pending | track | Benchmark phase `parse`. |
| Resolution phase time | ms | pending | 4057.1 | pending | track | Benchmark phase `resolution`. |
| Semantic enrichment phase time | ms | pending | 746.8 | pending | track | Benchmark phase `semantic_enrichment`. |
| DB load phase time | ms | pending | 13688.3 | pending | track | Benchmark phase `db_load`; largest recorded phase. |
| Files processed per second | files/s | pending | 23.0 | pending | track | `827` scanned files / total duration. |
| Nodes generated per second | nodes/s | pending | 2618.9 | pending | track | `94171` nodes / total duration. |
| Relationships generated per second | relationships/s | pending | 3582.9 | pending | track | `128835` relationships / total duration. |
| Analyze end heap allocation | MB | pending | 493.6 | pending | track | `memory.endAllocBytes`. |
| Analyze max observed sys memory | MB | pending | 875.1 | pending | track | `memory.maxObservedSys`. |

## B1 - File Inventory Counts

Status: P0-A partial baseline recorded

| Metric | Unit | Baseline | Latest | Delta | Target | Notes |
|---|---:|---:|---:|---:|---:|---|
| Indexed files | files | 819 | 829 | +10 | record | All files represented by projection after P5 refresh. |
| Source files | files | 368 | 399 | +31 | record | File kind count from `file-hotspots --repo Anvien --json --limit 0`. |
| Test files | files | 231 | 236 | +5 | record | File kind count from `file-hotspots --repo Anvien --json --limit 0`. |
| Generated files | files | 2 | 2 | 0 | record | `generated_contract` app layer. |
| Docs files | files | 178 | 178 | 0 | record | `docs` app layer. |
| Config files | files | 14 | 14 | 0 | record | `config` app layer. |
| Files with symbols | files | 575 | 585 | +10 | record | File summaries with `symbolCount > 0`. |
| Files with unresolved source sites | files | 576 | 586 | +10 | track | File summaries with `unresolvedSourceSiteCount > 0`. |
| Files linked to flows | files | pending | 161 | pending | track | File summaries with `linkedFlowCount > 0` from `file-hotspots --limit 0`. |
| Files linked to tests | files | pending | 349 | pending | track | File summaries with `linkedTestCount > 0` from `file-hotspots --limit 0`. |

## B2 - Symbol Tree Coverage

Status: P0-A partial baseline recorded

| Metric | Unit | Baseline | Latest | Delta | Target | Notes |
|---|---:|---:|---:|---:|---:|---|
| Declared symbols grouped by file | symbols | 21334 | 21957 | +623 | record | `DEFINES` relationships from `File` nodes. |
| Top-level symbols | symbols | pending | pending | pending | record | Root nodes in symbol trees. |
| Nested symbols/methods | relationships | 2641 | 2651 | +10 | record | Existing non-file `CONTAINS` relationships available for tree derivation. |
| Exported symbols | symbols | pending | pending | pending | record | Language-dependent export/public count. |
| Symbols with line ranges | symbols | pending | pending | pending | maximize | Required for useful navigation. |
| Symbols with signatures | symbols | pending | pending | pending | track | Useful but language-dependent. |

## B3 - Relationship Projection Counts

Status: P0-A partial baseline recorded

| Metric | Unit | Baseline | Latest | Delta | Target | Notes |
|---|---:|---:|---:|---:|---:|---|
| Derived file dependency edges | edges | pending | pending | pending | record | `File -> File` derived from symbol/source-site relationships. |
| Local file relationships | relationships | pending | pending | pending | record | Source and target in same file. |
| Outbound file relationships | relationships | pending | pending | pending | record | Source file depends on another file. |
| Inbound file relationships | relationships | pending | pending | pending | record | Other files depend on source file. |
| Relationship samples retained per group | samples | pending | pending | pending | bounded | Default output limit. |
| Relationship total counts preserved | percent | pending | pending | pending | 100 | Counts must not be truncated with samples. |
| Source-site-backed relationships | relationships | 83143 | 85312 | +2169 | record | Relationships carrying source-site/file trace fields. |
| Resolved source-site-backed relationships | relationships | 17491 | 18192 | +701 | record | Relationships with resolved source-site status. |

## B4 - Unresolved And Quality Counts

Status: P0-A partial baseline recorded

| Metric | Unit | Baseline | Latest | Delta | Target | Notes |
|---|---:|---:|---:|---:|---:|---|
| Unresolved source sites grouped by file | sites | 66555 | 68028 | +1473 | record | Graph-health source-backed unresolved references. |
| Unresolved calls | sites | 34959 | 35615 | +656 | record | Graph-health `resolutionGapFactFamilyCounts.call`. |
| Unresolved refs | sites | 31596 | 32413 | +817 | record | Access + type-reference + heritage gap counts. |
| Unresolved imports | sites | pending | 0 | pending | record | No import gap bucket in latest graph-health output. |
| Analyzer-gap classified sites | sites | 39535 | 40140 | +605 | record | Graph-health actionability bucket. |
| External/dynamic classified sites | sites | 250 | 266 | +16 | record | Graph-health `external_library` classification bucket. |
| Files with high unresolved count | files | pending | 582 | pending | track | Files with unresolved source sites after P3-B linked overlay analyze. |

## B5 - Projection Performance

Status: P3-B projection benchmark recorded

| Metric | Unit | Baseline | Latest | Delta | Target | Notes |
|---|---:|---:|---:|---:|---:|---|
| Full file summary aggregation time | ms | pending | 58.7 | pending | no unacceptable regression | Median of 3 `BenchmarkBuildFileListCurrentScale` runs after linked flow/test aggregation; fixture uses 821 files and 126000 relationships. |
| Projection cache cold build time | ms | pending | 8.0 | pending | track | Median of 3 `BenchmarkBuilderCacheColdBuild` runs. |
| Projection cache warm hit time | ms | pending | 0.000305 | pending | faster than cold | Median of 3 `BenchmarkBuilderCacheHit` runs with explicit graph hash. |
| Projection cache invalidation count | events | pending | 2 | pending | record | Unit tests cover graph-change miss and explicit repo/path invalidation. |
| Single file context cold query time | ms | pending | 7017.5 | pending | responsive | HTTP runtime first request includes graph snapshot load for `internal/httpapi/file_context.go`. |
| Single file context warm query time | ms | pending | 107.7 | pending | responsive | HTTP runtime second request for `internal/httpapi/file_context.go` after projection snapshot cache. |
| Hotspot list query time | ms | pending | 93.7 | pending | responsive | HTTP runtime warm `GET /api/file-hotspots?repo=Anvien&sort=unresolved&limit=5`; builder list benchmark median remains `131.3 ms`. |
| Peak memory during projection benchmark | MB | pending | 0.49 | pending | track | `490000 B/op` from `BenchmarkBuildFileListCurrentScale`. |
| Projection cache cold build memory | MB | pending | 0.57 | pending | track | Median allocation from `BenchmarkBuilderCacheColdBuild`. |
| Projection cache warm hit allocations | allocs/op | pending | 0 | pending | `0` | `BenchmarkBuilderCacheHit` reports zero allocations. |

## B6 - Response Size Metrics

Status: pending

| Metric | Unit | Baseline | Latest | Delta | Target | Notes |
|---|---:|---:|---:|---:|---:|---|
| Compact human file-context output | lines | pending | 76 | pending | readable | `file-context internal/cli/command.go --repo Anvien`; bounded default human output. |
| Full JSON file-context response | bytes | pending | 140289 | pending | bounded | `file-context internal/cli/command.go --repo Anvien --json`; full contract output. |
| File hotspot JSON response | bytes | pending | 10099 | pending | bounded | `file-hotspots --repo Anvien --json`; default limit returns 20 rows from the then-current file set. |
| Full file hotspot JSON response | bytes | pending | 399214 | pending | bounded | `file-hotspots --repo Anvien --json --limit 0`; returns all 825 file rows with linked flow/test counts. |
| Linked file-context JSON response | bytes | pending | 85584 | pending | bounded | `file-context internal/httpapi/file_context.go --repo Anvien --json`; includes linked counts and samples. |
| Web file list response | bytes | pending | 2086 | pending | bounded | HTTP `GET /api/file-hotspots?repo=Anvien&sort=unresolved&limit=5`. |
| Web file context response | bytes | pending | 57803 | pending | bounded | HTTP `GET /api/file-context?repo=Anvien&path=internal/httpapi/file_context.go`. |
| Web File Map default response | bytes | pending | 79058 | pending | bounded | HTTP `GET /api/file-hotspots?repo=Anvien&sort=unresolved&limit=200`; `827` total, `200` rows. |
| Web File Map changed response | bytes | pending | 2357 | pending | bounded | HTTP `GET /api/file-hotspots?repo=Anvien&changedOnly=true&limit=5`; `11` total, `5` rows. |
| Relationship samples per file default | samples | pending | 5 | pending | bounded | `file-context` default `--relationships 5`; counts preserve totals. |
| Compact file-hotspots output | lines | pending | 5 | pending | readable | `file-hotspots --repo Anvien --sort unresolved --limit 3`. |

## B7 - Web UI Metrics

Status: P4-A File Map recorded

| Metric | Unit | Baseline | Latest | Delta | Target | Notes |
|---|---:|---:|---:|---:|---:|---|
| File Map rendered rows default | rows | pending | 200 | pending | usable | Default Web File Map API limit. |
| File Map filters | count | pending | 11 | pending | record | `6` kind filters plus changed, unresolved, API, high fan-in, and high fan-out toggles. |
| File Detail sections rendered | sections | pending | pending | pending | `6` | Summary, symbol tree, relationships, unresolved, linked overlays, source-site samples. |
| File Map sort modes | modes | pending | 7 | pending | record | Unresolved, fan-in, fan-out, symbols, flows, tests, path. |
| Changed File Map rows | rows | pending | 11 | pending | record | Runtime `changedOnly=true` total before P4-A commit. |
| File Map cold API time | ms | pending | 7093.7 | pending | track | First `limit=200` request after serve restart includes graph snapshot load. |
| File Map warm API time | ms | pending | 501.2 | pending | responsive | Second `limit=200` request after projection cache warmup. |
| Changed File Map API time | ms | pending | 543.2 | pending | responsive | Warm `changedOnly=true&limit=5` request, includes git changed-file collection. |
| E2E file-map assertions | assertions | pending | 8 | pending | record | Count only; pass/fail goes to evidence. |
| Visual overlap failures | count | pending | 0 | pending | `0` | No overlap failures observed in Playwright e2e/manual smoke. |

## B8 - Parent/Child Command Hierarchy Counts

Status: P5 parent/child command hierarchy recorded

| Metric | Unit | Baseline | Latest | Delta | Target | Notes |
|---|---:|---:|---:|---:|---:|---|
| Target-aware parent commands inventoried | commands | pending | 10 | pending | record | `context`, `impact`, `query`, `detect-changes`, `graph-health`, `api route-map`, `api tool-map`, `api shape-check`, `api impact`, and `group query`. |
| Child commands proposed | commands | pending | 16 | pending | record | Includes existing `file-context` and `file-hotspots`, plus P5 explicit child commands. |
| Child commands implemented | commands | 2 | 16 | +14 | record | Added `context` 2, `impact` 4, `query` 4, `detect-changes` 3, and `graph-health` 1 child command. |
| Shared target resolver cases | cases | pending | 11 | pending | record | `context` file/symbol/auto; `impact` file/symbol/route/tool/auto; `query` files/symbols/flows/api; `detect-changes` files/symbols/flows. |
| Parent commands kept backward-compatible | commands | pending | 5 | pending | match affected parents | `context`, `impact`, `query`, `detect-changes`, and existing `graph-health` subcommands remain available. |
| Ambiguous target cases tested | cases | pending | 1 | pending | record | Parent `context <target>` file-vs-symbol ambiguity suggestions. |
| Parent/child JSON parity tests | tests | pending | 3 | pending | pass | `context file`, `context symbol`, and `impact file` JSON target fields. |
| Child command help entries | entries | pending | 14 | pending | record | New P5 CLI child commands excluding existing `file-context` / `file-hotspots`. |
| Parent/child help golden tests | tests | pending | 1 | pending | pass | MCP surface schema snapshot updated for optional target dispatch fields. |

## B9 - Existing Command File-Layer Coverage

Status: pending

| Metric | Unit | Baseline | Latest | Delta | Target | Notes |
|---|---:|---:|---:|---:|---:|---|
| Existing graph commands inventoried | commands | pending | pending | pending | record | Command matrix coverage. |
| Commands classified `must add file layer` | commands | pending | pending | pending | record | Required additive file sections. |
| Commands classified `may add file layer` | commands | pending | pending | pending | record | Optional/contextual file sections. |
| Commands classified `no file layer` | commands | pending | pending | pending | record | Must have evidence-backed reasons. |
| Existing command outputs with old details preserved | commands | pending | pending | pending | match included commands | Regression guard. |
| Existing command outputs with file sections added | commands | pending | pending | pending | match `must add` commands | Additive behavior guard. |
| JSON command outputs with file-layer fields | commands | pending | pending | pending | record | Machine-readable parity. |
| Shared projection service consumers | surfaces | pending | 3 | pending | CLI+MCP+API+Web runtime | CLI and HTTP API consume `internal/filecontext`; Web UI consumes the generated API contract; MCP pending. |
| MCP/API equivalents with file-layer parity | surfaces | pending | 1 | pending | record | HTTP API exposes file context and hotspot equivalents. |
| Embedded Anvien skills updated | skills | pending | pending | pending | record | Source-of-truth skill Markdown. |
| Generated Anvien skills updated | skills | pending | pending | pending | match source | Generated output after analyze/setup path. |
| Root generated context files updated | files | pending | pending | pending | record | `AGENTS.md`, `CLAUDE.md` if changed. |
| Command integration tests | tests | pending | pending | pending | pass | Existing details preserved plus file layer added. |
| Skill/context parity tests | tests | pending | pending | pending | pass | Source/generated guidance parity. |

## B10 - Final Validation Counts

Status: P1-A partial validation counts recorded

| Metric | Unit | Baseline | Latest | Delta | Target | Notes |
|---|---:|---:|---:|---:|---:|---|
| Projection unit tests | tests | 3 | 10 | +7 | pass | `internal/filecontext` tests through P4-A changed-file filter support. |
| CLI tests | tests | pending | 8 | pending | pass | P2-A relevant tests plus P5 target-command tests for context/impact/query/detect/graph-health child views. |
| API tests | tests | pending | 3 | pending | pass | File-context success, hotspot/filter, and missing-file endpoint tests. |
| MCP surface snapshot tests | tests | pending | 1 | pending | pass | MCP tool schema snapshot updated for optional target dispatch fields. |
| Contract generator checks | checks | pending | 1 | pending | pass | `go run .\cmd\generate-web-contracts --check`. |
| AI context tests | tests | pending | pending | pending | pass | Generated guidance and embedded skill parity. |
| Projection cache tests | tests | pending | 2 | pending | pass | Cold/warm, graph-change, repo-switch, explicit hash, and explicit invalidation coverage. |
| Web unit tests | tests | pending | 23 | pending | pass | `FileMapPanel.test.tsx` and `server-connection.test.ts` after P4-A. |
| Web e2e tests | tests | pending | 1 | pending | pass | File Map e2e in `shell-interactions.spec.ts`. |
| Detect-changes changed files | files | 5 | 13 | +8 | record | P4-A final detect-changes scope. |
| Detect-changes affected count | symbols/processes | 17 | 14 | -3 | record | P4-A final affected process/symbol count; HIGH due shared projection/API/contract/Web UI surface. |
| Detect-changes changed symbols | symbols | pending | 260 | pending | record | P4-A final detect-changes summary. |
| Detect-changes resolution gap changed entities | entities | pending | 194 | pending | record | P4-A final detect-changes resolution gap summary. |
