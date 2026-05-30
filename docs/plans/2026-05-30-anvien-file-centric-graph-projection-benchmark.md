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
| Files scanned | files | 819 | 821 | +2 | record | Latest after P1-C `anvien analyze --force --name Anvien --benchmark-json ...`. |
| Files parsed | files | 584 | 586 | +2 | record | Latest after P1-C `anvien analyze --force --name Anvien --benchmark-json ...`. |
| Unsupported files | files | 235 | 235 | 0 | record | Latest after P1-C `anvien analyze --force --name Anvien --benchmark-json ...`. |
| Failed files | files | 0 | 0 | 0 | `0` | From readiness `anvien analyze --force --name Anvien`. |
| Graph nodes | nodes | 91587 | 92652 | +1065 | record | Current graph inventory after P1-C final analyze benchmark. |
| Graph relationships | relationships | 125054 | 126821 | +1767 | record | Current graph inventory after P1-C final analyze benchmark. |
| SourceSite count | distinct ids | 95433 | 96762 | +1329 | record | Distinct source-site ids from relationship trace fields. |
| ResolutionGap count | nodes | 65652 | 66322 | +670 | record | Required for quality projection coverage. |

## B0A - Graph Generation Speed

Status: P1-C analyze benchmark recorded

| Metric | Unit | Baseline | Latest | Delta | Target | Notes |
|---|---:|---:|---:|---:|---:|---|
| Analyze benchmark runs | runs | pending | 1 | pending | record | Final recorded run: `anvien analyze --force --name Anvien --benchmark-label file-centric-p1c-final --benchmark-json .tmp\file-centric-p1c-final-analyze-benchmark.json`. |
| Total graph generation time | ms | pending | 34839.3 | pending | track | Includes scan, parse, resolution, semantic enrichment, and DB load. |
| Scan phase time | ms | pending | 158.4 | pending | track | Benchmark phase `scan`. |
| Parse phase time | ms | pending | 8779.1 | pending | track | Benchmark phase `parse`. |
| Resolution phase time | ms | pending | 3743.7 | pending | track | Benchmark phase `resolution`. |
| Semantic enrichment phase time | ms | pending | 774.5 | pending | track | Benchmark phase `semantic_enrichment`. |
| DB load phase time | ms | pending | 13783.2 | pending | track | Benchmark phase `db_load`; largest recorded phase. |
| Files processed per second | files/s | pending | 23.6 | pending | track | `821` scanned files / total duration. |
| Nodes generated per second | nodes/s | pending | 2659.4 | pending | track | `92652` nodes / total duration. |
| Relationships generated per second | relationships/s | pending | 3640.2 | pending | track | `126821` relationships / total duration. |
| Analyze end heap allocation | MB | pending | 510.0 | pending | track | `memory.endAllocBytes`. |
| Analyze max observed sys memory | MB | pending | 811.7 | pending | track | `memory.maxObservedSys`. |

## B1 - File Inventory Counts

Status: P0-A partial baseline recorded

| Metric | Unit | Baseline | Latest | Delta | Target | Notes |
|---|---:|---:|---:|---:|---:|---|
| Indexed files | files | 819 | 821 | +2 | record | All files represented by projection. |
| Source files | files | 368 | 369 | +1 | record | `backend`, `api`, `frontend`, launcher/client/contract source app layers. |
| Test files | files | 231 | 232 | +1 | record | `backend_test`, `api_test`, `frontend_test`. |
| Generated files | files | 2 | 2 | 0 | record | `generated_contract` app layer. |
| Docs files | files | 178 | 178 | 0 | record | `docs` app layer. |
| Config files | files | 14 | 14 | 0 | record | `config` app layer. |
| Files with symbols | files | 575 | 577 | +2 | record | Distinct `File` nodes with outgoing `DEFINES`. |
| Files with unresolved source sites | files | 576 | 578 | +2 | track | Distinct `ResolutionGap.filePath` values. |
| Files linked to flows | files | pending | pending | pending | track | Flow overlay coverage. |
| Files linked to tests | files | pending | pending | pending | track | Test overlay coverage. |

## B2 - Symbol Tree Coverage

Status: P0-A partial baseline recorded

| Metric | Unit | Baseline | Latest | Delta | Target | Notes |
|---|---:|---:|---:|---:|---:|---|
| Declared symbols grouped by file | symbols | 21334 | 21746 | +412 | record | `DEFINES` relationships from `File` nodes. |
| Top-level symbols | symbols | pending | pending | pending | record | Root nodes in symbol trees. |
| Nested symbols/methods | relationships | 2641 | 2647 | +6 | record | Existing non-file `CONTAINS` relationships available for tree derivation. |
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
| Source-site-backed relationships | relationships | 83143 | 84237 | +1094 | record | Relationships carrying source-site/file trace fields. |
| Resolved source-site-backed relationships | relationships | 17491 | 17915 | +424 | record | Relationships with resolved source-site status. |

## B4 - Unresolved And Quality Counts

Status: P0-A partial baseline recorded

| Metric | Unit | Baseline | Latest | Delta | Target | Notes |
|---|---:|---:|---:|---:|---:|---|
| Unresolved source sites grouped by file | sites | 66555 | 67223 | +668 | record | Graph-health source-backed unresolved references. |
| Unresolved calls | sites | 34959 | 35242 | +283 | record | Graph-health `resolutionGapFactFamilyCounts.call`. |
| Unresolved refs | sites | 31596 | 31981 | +385 | record | Access + type-reference + heritage gap counts. |
| Unresolved imports | sites | pending | pending | pending | record | Gap kind count. |
| Analyzer-gap classified sites | sites | 39535 | 39780 | +245 | record | Graph-health actionability bucket. |
| External/dynamic classified sites | sites | 250 | 250 | 0 | record | Graph-health `external_library` classification bucket. |
| Files with high unresolved count | files | pending | pending | pending | track | Hotspot metric. |

## B5 - Projection Performance

Status: pending

| Metric | Unit | Baseline | Latest | Delta | Target | Notes |
|---|---:|---:|---:|---:|---:|---|
| Full file summary aggregation time | ms | pending | 131.3 | pending | no unacceptable regression | Median of 3 `BenchmarkBuildFileListCurrentScale` runs over 821 files and 126000 relationships. |
| Projection cache cold build time | ms | pending | 8.0 | pending | track | Median of 3 `BenchmarkBuilderCacheColdBuild` runs. |
| Projection cache warm hit time | ms | pending | 0.000305 | pending | faster than cold | Median of 3 `BenchmarkBuilderCacheHit` runs with explicit graph hash. |
| Projection cache invalidation count | events | pending | 2 | pending | record | Unit tests cover graph-change miss and explicit repo/path invalidation. |
| Single file context cold query time | ms | pending | pending | pending | responsive | Representative large source file. |
| Single file context warm query time | ms | pending | pending | pending | responsive | Representative large source file. |
| Hotspot list query time | ms | pending | 131.3 | pending | responsive | Median of 3 `BenchmarkBuildFileListCurrentScale` runs over 821 files and 126000 relationships. |
| Peak memory during projection benchmark | MB | pending | 0.49 | pending | track | `490000 B/op` from `BenchmarkBuildFileListCurrentScale`. |
| Projection cache cold build memory | MB | pending | 0.57 | pending | track | Median allocation from `BenchmarkBuilderCacheColdBuild`. |
| Projection cache warm hit allocations | allocs/op | pending | 0 | pending | `0` | `BenchmarkBuilderCacheHit` reports zero allocations. |

## B6 - Response Size Metrics

Status: pending

| Metric | Unit | Baseline | Latest | Delta | Target | Notes |
|---|---:|---:|---:|---:|---:|---|
| Compact human file-context output | lines | pending | pending | pending | readable | Default CLI output. |
| Full JSON file-context response | bytes | pending | pending | pending | bounded | Representative large source file. |
| File hotspot JSON response | bytes | pending | pending | pending | bounded | Default limit. |
| Web file list response | bytes | pending | pending | pending | bounded | Default page size. |
| Relationship samples per file default | samples | pending | pending | pending | bounded | Counts must preserve totals. |

## B7 - Web UI Metrics

Status: pending

| Metric | Unit | Baseline | Latest | Delta | Target | Notes |
|---|---:|---:|---:|---:|---:|---|
| File Map rendered rows default | rows | pending | pending | pending | usable | Default page size or virtualized count. |
| File Map filters | count | pending | pending | pending | record | Changed, unresolved, API, generated, fan-in, fan-out, etc. |
| File Detail sections rendered | sections | pending | pending | pending | `6` | Summary, symbol tree, relationships, unresolved, linked overlays, source-site samples. |
| E2E file-map assertions | assertions | pending | pending | pending | record | Count only; pass/fail goes to evidence. |
| Visual overlap failures | count | pending | pending | pending | `0` | If screenshot/browser validation is used. |

## B8 - Parent/Child Command Hierarchy Counts

Status: pending

| Metric | Unit | Baseline | Latest | Delta | Target | Notes |
|---|---:|---:|---:|---:|---:|---|
| Target-aware parent commands inventoried | commands | pending | pending | pending | record | Commands where target type can change behavior. |
| Child commands proposed | commands | pending | pending | pending | record | Explicit file/symbol/route/tool/flow/API commands. |
| Child commands implemented | commands | pending | pending | pending | record | Actual implemented child commands. |
| Shared target resolver cases | cases | pending | pending | pending | record | File/symbol/route/tool/flow/API disambiguation cases. |
| Parent commands kept backward-compatible | commands | pending | pending | pending | match affected parents | Existing syntax still works. |
| Ambiguous target cases tested | cases | pending | pending | pending | record | Parent command ambiguity suggestions. |
| Parent/child JSON parity tests | tests | pending | pending | pending | pass | Same resolved target has compatible JSON shape. |
| Child command help entries | entries | pending | pending | pending | record | Help discoverability. |
| Parent/child help golden tests | tests | pending | pending | pending | pass | Help and flat syntax compatibility. |

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
| Shared projection service consumers | surfaces | pending | pending | pending | CLI+MCP+API+Web runtime | Avoid per-surface derivation drift. |
| MCP/API equivalents with file-layer parity | surfaces | pending | pending | pending | record | Agent/API parity where equivalent commands exist. |
| Embedded Anvien skills updated | skills | pending | pending | pending | record | Source-of-truth skill Markdown. |
| Generated Anvien skills updated | skills | pending | pending | pending | match source | Generated output after analyze/setup path. |
| Root generated context files updated | files | pending | pending | pending | record | `AGENTS.md`, `CLAUDE.md` if changed. |
| Command integration tests | tests | pending | pending | pending | pass | Existing details preserved plus file layer added. |
| Skill/context parity tests | tests | pending | pending | pending | pass | Source/generated guidance parity. |

## B10 - Final Validation Counts

Status: P1-A partial validation counts recorded

| Metric | Unit | Baseline | Latest | Delta | Target | Notes |
|---|---:|---:|---:|---:|---:|---|
| Projection unit tests | tests | 3 | 7 | +4 | pass | `internal/filecontext` tests through P1-C. |
| CLI tests | tests | pending | pending | pending | pass | Count of relevant tests. |
| API tests | tests | pending | pending | pending | pass | Count of relevant tests. |
| MCP surface snapshot tests | tests | pending | pending | pending | pass | If MCP output/schema changes. |
| Contract generator checks | checks | pending | pending | pending | pass | Source and generated Web contracts in sync. |
| AI context tests | tests | pending | pending | pending | pass | Generated guidance and embedded skill parity. |
| Projection cache tests | tests | pending | 2 | pending | pass | Cold/warm, graph-change, repo-switch, explicit hash, and explicit invalidation coverage. |
| Web unit tests | tests | pending | pending | pending | pass | Count if Web UI is implemented. |
| Web e2e tests | tests | pending | pending | pending | pass | Count if Web UI is implemented. |
| Detect-changes changed files | files | 5 | 5 | 0 | record | Latest P1-C staged implementation scope. |
| Detect-changes affected count | symbols/processes | 17 | 0 | -17 | record | Latest P1-C staged affected process/symbol count. |
