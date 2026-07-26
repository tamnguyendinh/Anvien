# Index Reader Matrix (P2-A1 Planned Artifact)

Status: `planned seed inventory at self-graph HEAD 1932359bee5d78fd8e6167ef94e7974f33e85bd0 / not P2-A completion evidence`

This matrix is the checked-in seed for the version/generation denominator. Every listed anchor must exist and its semantics must match source, but the seed is not a completeness claim: P2-A1 must re-run the source scan, classify every dispatcher/non-reader explicitly, assign every row to one exact later guard-owner slice, add one exact row for every newly discovered graph/index reader/backend, and attach the final matrix hash to `E2-P2A1-MATRIXREVIEW1`. A row is not proof merely because it is listed. Each row `RNN` reserves its row-specific classification evidence ID in the form `E2-P2A1-R<row-number>`; `pending` means that evidence has not yet been produced. Runtime guard PASS belongs to the assigned P2-A2..P2-A15 owner slice, not to the inventory proof.

Global guard contract used verbatim by every row:

- Request fields: `{readerProtocolVersion, readerBuild, supportedGraphSchemaVersions[], supportedIdentitySchemaVersions[], supportedScopeIrVersions[]}`.
- Manifest fields: `protocolVersion`, `minReaderProtocol`, `minReaderBuild`, `graphSchemaVersion`, `identitySchemaVersion`, `scopeIrVersion`, `generation`, `configHash`, and `catalogHash`.
- Validation occurs before opening graph body, database, cache, registry, group snapshot, embedding row, or stream. Any absent, stale, mixed, or unsupported value returns `{code:"INDEX_VERSION_MISMATCH", expected:{...}, actual:{...}, retryable:false}` through the row's native contract: loader/driver error for S0–S2, stderr/nonzero for S3, JSON-RPC `data.code` for S4, HTTP `409` before body for S5, mismatch UI for S6, cache miss/error for S7–S8, and typed no-data error for S9–S11.
- v2 is stored in a non-overlapping protocol/layout namespace; an unaware v1 reader cannot interpret v2 bytes as v1.
- P2-A1 inventory passes only when `rows_classified == rows_total`, `unassigned_rows == 0`, and the source scan reports `unlisted_readers == 0`. P2-E2 runtime acceptance later requires `rows_passed == rows_total`. The current rows are a seed, not a fixed upper bound and not permission to retain a catch-all/router row in place of exact children.

| ID | Exact source path and function/entrypoint | Surface | Backend/layout | Guard / fixture / expected failure | Evidence |
|----|--------------------------------------------|---------|----------------|------------------------------------|----------|
| R01 | `internal/graphaccuracy/graphaccuracy.go:ReadGraph` | S0/audit | graph JSON | manifest guard; absent/old/mixed JSON -> `INDEX_VERSION_MISMATCH` | pending |
| R02 | `cmd/graph-accuracy-probe/main.go:runFreshAnalyze` | S0/audit | newly analyzed Go graph copied to `-fresh-go-graph` | validate the generated graph manifest before copy/return; this does not guard the pre-existing `-node`/`-go` inputs | pending |
| R03 | `cmd/graph-accuracy-probe/main.go:main` | S0/audit | two independent input graphs from `-node` and `-go`; `-out` is report JSON only | validate both input manifests independently before comparison; either mismatch -> `INDEX_VERSION_MISMATCH` | pending |
| R04 | `cmd/property-access-audit/main.go:main` | S0/audit | graph JSON | `-graph` guard before audit | pending |
| R05 | `cmd/access-candidate-audit/main.go:main` | producer/non-reader classification | analyzer + in-memory ScopeIR audit | explicit exclusion: it does not open a persisted graph/index; P2-A source scan records `not_reader` rather than inventing a manifest guard | pending |
| R06 | `internal/repo/meta.go:LoadMeta` | metadata | repo-local meta | manifest/schema guard | pending |
| R07 | `internal/repo/meta.go:LoadIndexed` | metadata | repo-local `.anvien/meta.json` and derived storage paths | manifest/version/generation guard before returning `Indexed` | pending |
| R08 | `internal/repo/meta.go:FindIndexed` | metadata | ancestor traversal calling repo-local `LoadIndexed` | every discovered ancestor meta uses the same guard; unsupported meta -> `INDEX_VERSION_MISMATCH` | pending |
| R09 | `internal/repo/meta.go:HasIndex` | metadata | repo-local index | unsupported layout -> `INDEX_VERSION_MISMATCH` | pending |
| R10 | `internal/repo/meta.go:HasLegacyKuzuIndex` | metadata | legacy layout | explicit legacy rejection/compatibility fixture | pending |
| R11 | `internal/session/resolver.go:(StoreResolver).Resolve` | startup/session | registry/index | handshake before repo binding | pending |
| R12 | `internal/session/resolver.go:(StoreResolver).resolveName` | startup/session | registry | ambiguous/stale registry -> `INDEX_VERSION_MISMATCH` or typed lookup error | pending |
| R13 | `internal/session/resolver.go:(StoreResolver).resolvePath` | startup/session | repo metadata | stale generation fixture | pending |
| R14 | `internal/cli/command.go:NewRootCommand` | S3 startup | CLI root | old binary handshake before dispatch | pending |
| R15 | `internal/cli/command.go:newMCPCommand` | S3 startup | MCP process startup | protocol negotiation before server body | pending |
| R16 | `internal/cli/command.go:newServeCommand` | S3 startup | HTTP/Web server startup | manifest guard before listener serves graph | pending |
| R17 | `internal/cli/command.go:newAnalyzeCommand` | S3 analyze | analyze lifecycle | active epoch guard | pending |
| R18 | `internal/cli/admin_command.go:newIndexCommand` | S3 index | repo index | unsupported protocol -> `INDEX_VERSION_MISMATCH` | pending |
| R19 | `internal/cli/status_command.go:newStatusCommand` | S3 status | registry/meta | stale/mixed registry fixture | pending |
| R20 | `internal/cli/file_detail_command.go:newFileDetailCommand` | S3 file-detail | Graph JSON through `loadFileProjectionGraph`; no Ladybug path | manifest/body guard before projection | pending |
| R21 | `internal/cli/file_detail_command.go:loadFileProjectionGraph` | S3 file-detail | graph JSON | mixed generation fixture | pending |
| R22 | `internal/cli/file_detail_command.go:newFileHotspotsCommand` | S3 file-detail | graph JSON | guard before hotspots | pending |
| R23 | `internal/cli/graph_health_command.go:newGraphHealthCommand` | S3 dispatcher/non-reader | graph-health parent command | explicit dispatcher classification; exact summary/report/components/explain/files children have separate rows | pending |
| R24 | `internal/cli/graph_health_command.go:loadGraphHealthGraph` | S3 health | graph JSON | absent/old graph -> `INDEX_VERSION_MISMATCH` | pending |
| R25 | `internal/cli/source_site_accuracy_command.go:newSourceSiteAccuracyCommand` | S3 audit | graph JSON | `--graph` guard | pending |
| R26 | `internal/cli/resolution_inventory_command.go:newResolutionInventoryCommand` | S3 audit | graph JSON | catalog/generation guard | pending |
| R27 | `internal/cli/resolution_inventory_command.go:readResolutionInventoryGraph` | S3 audit | graph JSON | mixed graph fixture | pending |
| R28 | `internal/cli/query_health_command.go:newQueryHealthCommand` | S3 benchmark | graph/query backend | guard before query | pending |
| R29 | `internal/cli/query_health_command.go:runLocalQueryHealthQuery` | S3 benchmark | local MCP `query` tool path | MCP query guard/generation fixture; it is not the native/fallback Cypher selector | pending |
| R30 | `internal/cli/benchmark_command.go:newBenchmarkCompareCommand` | S3 benchmark | benchmark artifacts | generation-bound benchmark metadata | pending |
| R31 | `internal/cli/benchmark_command.go:readBenchmarkMetrics` | S3 benchmark | benchmark JSON | stale baseline -> `INDEX_VERSION_MISMATCH` | pending |
| R32 | `internal/cli/tool_command.go:newQueryCommand` | S3 query | graph/query backend | guard before query | pending |
| R33 | `internal/cli/tool_command.go:runQueryToolCommand` | S3 query | in-process MCP query tool over Graph JSON/resource cache; no Ladybug path | generation equality fixture before local MCP dispatch | pending |
| R34 | `internal/cli/tool_command.go:newContextCommand` | S3 context | graph/file-context | cache/manifest guard | pending |
| R35 | `internal/cli/tool_command.go:newImpactCommand` | S3 impact | graph | guard before traversal | pending |
| R36 | `internal/cli/tool_command.go:runImpactTargetCommand` | S3 impact | graph | stale target fixture | pending |
| R37 | `internal/cli/tool_command.go:newRenameCommand` | S3 rename | graph/source files | generation/source commit guard | pending |
| R38 | `internal/cli/tool_command.go:newCypherCommand` | S3 Cypher | native/fallback | separate backend guard | pending |
| R39 | `internal/cli/tool_command.go:newDetectChangesCommand` | S3 detect | graph/registry | stale generation fixture | pending |
| R40 | `internal/cli/group_command.go:newGroupCommand` | S3 dispatcher/non-reader | group parent command | explicit dispatcher classification; create/add/remove/list/status/sync/query/contracts children have separate rows | pending |
| R41 | `internal/cli/group_command.go:newGroupStatusCommand` | S3 group | group contracts | member-generation vector mismatch | pending |
| R42 | `internal/cli/group_command.go:newGroupSyncCommand` | S3 group | group contracts + repo graphs | CAS conflict -> `INDEX_VERSION_MISMATCH` | pending |
| R43 | `internal/cli/group_command.go:newGroupQueryCommand` | S3 group | member graph snapshots | pinned vector guard | pending |
| R44 | `internal/cli/group_command.go:newGroupContractsCommand` | S3 group | contracts registry | stale group epoch fixture | pending |
| R45 | `internal/cli/group_command.go:loadGroupConfigAndDir` | S3 group | group config/storage | path/epoch guard | pending |
| R46 | `internal/cli/doctor_command.go:newDoctorCommand` | S3 dispatcher/non-reader | doctor parent command | explicit dispatcher classification; locks/processes children have separate rows | pending |
| R47 | `internal/cli/doctor_command.go:newDoctorLocksCommand` | S3 doctor | lock/meta | generation-aware lock fixture | pending |
| R48 | `internal/cli/doctor_command.go:newDoctorProcessesCommand` | S3 doctor/non-reader | operating-system process scan only | explicit exclusion: no persisted graph/index is opened, so record `not_reader` rather than inventing an epoch guard | pending |
| R49 | `internal/lbugload/load.go:LoadCSVExportWithOptions` | S1 persistence | Ladybug CSV load | schema/generation guard; skipped/partial load fails closed | pending |
| R50 | `internal/lbugnative/runner_ladybugdb.go:OpenReadRunner` | S1 persistence | native Ladybug DB (`ladybugdb`) | manifest guard before open | pending |
| R51 | `internal/lbugnative/runner_ladybugdb.go:OpenWriteRunner` | S1 persistence | native Ladybug DB (`ladybugdb`) | active epoch guard | pending |
| R52 | `internal/lbugnative/native_ladybugdb.go:openNativeDatabase` | S1 persistence | native Ladybug DB | schema/protocol guard | pending |
| R53 | `internal/lbugruntime/embedding_cache.go:FetchExistingEmbeddingHashes` | S9 embeddings | embedding table | generation/node/hash guard | pending |
| R54 | `internal/mcp/server.go:Serve` | S4 MCP | stdio MCP | handshake before dispatch | pending |
| R55 | `internal/mcp/server.go:HandleJSONRPC` | S4 MCP | JSON-RPC | protocol/schema negotiation | pending |
| R56 | `internal/mcp/resources.go:(Server).readResource` | S4 MCP | resource registry/cache | epoch guard | pending |
| R57 | `internal/mcp/resources.go:(Server).contextResource` | S4 MCP | graph/context cache | generation-bound cache | pending |
| R58 | `internal/mcp/resources.go:(Server).reposResource` | S4 MCP | global repo registry | registry epoch guard | pending |
| R59 | `internal/mcp/resources.go:(Server).clustersResource` | S4/S11 MCP | graph-derived projection | generation/provenance guard | pending |
| R60 | `internal/mcp/resources.go:(Server).processesResource` | S4/S11 MCP | graph-derived projection | generation/provenance guard | pending |
| R61 | `internal/mcp/resources.go:(Server).graphForResource` | S4 MCP | Graph JSON directly or through the MCP resource cache; no Ladybug path | manifest/body and cache-generation guard | pending |
| R62 | `internal/mcp/resources.go:loadResourceGraphSnapshot` | S4 MCP | graph JSON | absent/old/mixed -> `INDEX_VERSION_MISMATCH` | pending |
| R63 | `internal/mcp/resources.go:(Server).routeIndexForResource` | S4 MCP | route index/cache | generation cache key | pending |
| R64 | `internal/mcp/resource_cache.go:(resourceGraphCache).graph` | S8 cache | MCP resource cache | lease/generation guard | pending |
| R65 | `internal/mcp/resource_cache.go:(resourceGraphCache).routeIndex` | S8 cache | MCP route cache | lease/generation guard | pending |
| R66 | `internal/mcp/tools.go:(Server).runCypherRead` | S4 MCP Cypher selector | dispatcher selecting native Ladybug or the separately inventoried Graph JSON + Go fallback | validate selected generation before dispatch; no fallback on version mismatch; R154/R155 prove the two concrete backends separately | pending |
| R67 | `internal/mcp/impact.go:(Server).impactTool` | S4 MCP | graph | generation guard | pending |
| R68 | `internal/mcp/detect_changes.go:(Server).detectChangesTool` | S4 MCP | repo resolution + Graph JSON + git diff | generation guard before graph load; target payload helper is not the reader | pending |
| R69 | `internal/mcp/rename.go:(Server).renameTool` | S4 MCP | graph/source | source commit + generation guard | pending |
| R70 | `internal/httpapi/server.go:NewHandler` | S5 transport boundary/non-reader | HTTP router registration | protocol negotiation middleware before dispatch; does not substitute for exact child-handler rows | pending |
| R71 | `internal/httpapi/graph.go:(Server).handleGraph` | S5 HTTP | graph JSON/stream | handshake before body | pending |
| R72 | `internal/httpapi/graph.go:streamGraphNDJSON` | S5 HTTP | graph stream | generation in stream header/records | pending |
| R73 | `internal/httpapi/query.go:(Server).handleQuery` | S5 HTTP query | Graph JSON loader plus separately inventoried in-memory Go fallback adapter | graph manifest guard before R156 `runGraphPanelQuery`; no native Ladybug claim | pending |
| R74 | `internal/httpapi/file_context.go:(Server).handleFileContext` | S5/S7 HTTP | file-context cache | generation cache key | pending |
| R75 | `internal/httpapi/file_context.go:(Server).loadFileProjection` | S5 HTTP | graph JSON | body guard | pending |
| R76 | `internal/httpapi/file.go:(Server).handleFile` | S5 HTTP | source/file projection | repo generation guard | pending |
| R77 | `internal/httpapi/grep.go:(Server).handleGrep` | S5 HTTP | graph/source | generation guard | pending |
| R78 | `internal/httpapi/repos.go:(Server).handleRepos` | S5/S10 HTTP | global registry | registry epoch guard | pending |
| R79 | `internal/httpapi/repos.go:(Server).handleRepoInfo` | S5/S10 HTTP | meta/registry | generation guard | pending |
| R80 | `internal/httpapi/analyze.go:(Server).handleAnalyze` | S5 HTTP | analyze lifecycle | active epoch guard | pending |
| R81 | `internal/httpapi/analyze.go:(Server).handleAnalyzeJob` | S5 HTTP | analyze job state | generation/job guard | pending |
| R82 | `internal/httpapi/embed.go:(Server).handleEmbed` | S5/S9 HTTP | embedding pipeline | generation/hash guard | pending |
| R83 | `internal/httpapi/embed.go:loadGraphSnapshot` | S5 HTTP | graph JSON | body guard | pending |
| R84 | `internal/httpapi/mcp.go:(Server).handleMCP` | S5/S4 HTTP-MCP | MCP session | handshake before JSON-RPC body | pending |
| R85 | `internal/httpapi/session.go:(Server).handleSessionStatus` | S5 HTTP | session status/context | generation/session guard before returning repo-bound state | pending |
| R86 | `internal/httpapi/search.go:(Server).handleSearch` | S5/S9 HTTP | graph/embedding search | generation/catalog/hash guard | pending |
| R87 | `internal/processes/processes.go:Apply` | S11 derived | graph-derived processes | generation/provenance guard | pending |
| R88 | `internal/communities/communities.go:Apply` | S11 derived | graph-derived communities | generation/provenance guard | pending |
| R89 | `internal/group/query.go:Query` | S10 group | member graph snapshots | member-generation vector guard | pending |
| R90 | `internal/group/status.go:Status` | S10 group | group contracts | member-generation vector guard | pending |
| R91 | `internal/group/storage.go:ReadRegistry` | S10 group | global group registry | group epoch/CAS guard | pending |
| R92 | `internal/group/storage.go:WriteRegistry` | S10 group | global group registry | staged write + CAS pointer | pending |
| R93 | `internal/group/sync.go:Sync` | S10 group | multi-repo group snapshots | vector pin/CAS conflict | pending |
| R94 | `internal/repo/registry.go:(Store).ReadRegistry` | S10 registry | global repo registry | active epoch guard | pending |
| R95 | `internal/repo/registry.go:(Store).ListRegistered` | S10 registry | global repo registry | active epoch guard | pending |
| R96 | `internal/repo/registry.go:(Store).Register` | S10 registry | global repo registry | staged write + CAS pointer | pending |
| R97 | `internal/repo/registry.go:(Store).WriteRegistry` | S10 registry | global repo registry | staged write + fsync | pending |
| R98 | `internal/cli/admin_command.go:newCleanCommand` | S3 clean | repo-local meta/index | validate resolved index protocol before destructive cleanup; unsupported layout -> `INDEX_VERSION_MISMATCH` | pending |
| R99 | `internal/cli/list_command.go:newListCommand` | S3 list | global repo registry | registry epoch/protocol guard before listing | pending |
| R100 | `internal/cli/api_command.go:newAPIRouteMapCommand` | S3 API route-map | MCP route-map graph path | complete MCP/graph guard before output | pending |
| R101 | `internal/cli/api_command.go:newAPIToolMapCommand` | S3 API tool-map | MCP tool-map graph path | complete MCP/graph guard before output | pending |
| R102 | `internal/cli/api_command.go:newAPIShapeCheckCommand` | S3 API shape-check | MCP route/consumer graph path | complete MCP/graph guard before output | pending |
| R103 | `internal/cli/api_command.go:newAPIImpactCommand` | S3 API impact | MCP route/consumer graph path | complete MCP/graph guard before output | pending |
| R104 | `internal/cli/tool_command.go:newAugmentCommand` | S3 augment | local MCP query/context graph path | generation/protocol guard before augmentation | pending |
| R105 | `internal/cli/hook_command.go:newHookCommand` | S3 dispatcher/non-reader | hook parent command registration only | explicit dispatcher classification; `runClaudeHook`, pre-tool graph augmentation, and post-tool metadata freshness branches have separate rows | pending |
| R106 | `internal/cli/graph_health_command.go:newGraphHealthSummaryCommand` | S3 health | graph JSON/outcomes | guard before summary projection | pending |
| R107 | `internal/cli/graph_health_command.go:newGraphHealthReportCommand` | S3 health | graph JSON/outcomes | guard before report projection | pending |
| R108 | `internal/cli/graph_health_command.go:newGraphHealthComponentsCommand` | S3 health | graph JSON/components | guard before component projection | pending |
| R109 | `internal/cli/graph_health_command.go:newGraphHealthExplainCommand` | S3 health | graph JSON/node outcome | guard before explain projection | pending |
| R110 | `internal/cli/graph_health_command.go:newGraphHealthFilesCommand` | S3 health | graph JSON/file outcomes | guard before file projection | pending |
| R111 | `internal/cli/group_command.go:newGroupCreateCommand` | S3 group | group directory plus `group.yaml`; no repo/group registry read | group config schema/protocol guard before overwrite; do not claim a registry generation | pending |
| R112 | `internal/cli/group_command.go:newGroupAddCommand` | S3 group | existing `group.yaml` membership map only; `registryName` is stored but not resolved here | group config schema/epoch guard before membership write; no false repo-registry guard claim | pending |
| R113 | `internal/cli/group_command.go:newGroupRemoveCommand` | S3 group | existing `group.yaml` membership map only | group config schema/epoch guard before membership removal; no false registry claim | pending |
| R114 | `internal/cli/group_command.go:newGroupListCommand` | S3 group | group directory and `group.yaml` only | group config schema/epoch guard before list; no false registry claim | pending |
| R115 | `internal/mcp/server.go:(Server).listRepos` | S4 MCP | global repo registry | registry epoch/protocol guard | pending |
| R116 | `internal/mcp/tools.go:(Server).queryTool` | S4 MCP | Graph JSON/resource cache query | generation guard before graph/cache use | pending |
| R117 | `internal/mcp/tools.go:(Server).cypherTool` | S4 MCP | `runCypherRead` native/fallback selector | validate request/manifest before dispatch | pending |
| R118 | `internal/mcp/context.go:(Server).contextTool` | S4 MCP | graph/file-context | generation/cache guard before context | pending |
| R119 | `internal/mcp/route_tool_map.go:(Server).routeMapTool` | S4 MCP | route index + graph | generation-bound route-index guard | pending |
| R120 | `internal/mcp/route_tool_map.go:(Server).toolMapTool` | S4 MCP | tool/route index + graph | generation-bound route-index guard | pending |
| R121 | `internal/mcp/route_shape_impact.go:(Server).shapeCheckTool` | S4 MCP | route/consumer graph | generation guard before shape comparison | pending |
| R122 | `internal/mcp/route_shape_impact.go:(Server).apiImpactTool` | S4 MCP | route/consumer graph | generation guard before API impact | pending |
| R123 | `internal/mcp/group_tools.go:(Server).groupListTool` | S4 MCP | group registry | group epoch/vector guard | pending |
| R124 | `internal/mcp/group_tools.go:(Server).groupStatusTool` | S4 MCP | group contracts | member-generation vector guard | pending |
| R125 | `internal/mcp/group_tools.go:(Server).groupSyncTool` | S4 MCP | group/repo snapshots | pinned vector + CAS conflict guard | pending |
| R126 | `internal/mcp/group_tools.go:(Server).groupContractsTool` | S4 MCP | group contract registry | group epoch/vector guard | pending |
| R127 | `internal/mcp/group_tools.go:(Server).groupQueryTool` | S4 MCP | member graph snapshots | pinned member-vector guard | pending |
| R128 | `internal/mcp/resources.go:(Server).clusterDetailResource` | S4/S11 MCP | graph-derived cluster detail | generation/provenance guard | pending |
| R129 | `internal/mcp/resources.go:(Server).processDetailResource` | S4/S11 MCP | graph-derived process detail | generation/provenance guard | pending |
| R130 | `internal/httpapi/graph.go:(Server).handleGraphHealthExplain` | S5 HTTP | graph JSON/outcomes | handshake and manifest guard before projection | pending |
| R131 | `internal/httpapi/graph.go:(Server).handleGraphHealthReport` | S5 HTTP | graph JSON/outcomes | handshake and manifest guard before projection | pending |
| R132 | `internal/httpapi/file_context.go:(Server).handleFileHotspots` | S5/S7 HTTP | graph JSON/file-context cache | generation/cache guard before response | pending |
| R133 | `internal/httpapi/panels.go:(Server).handleProcesses` | S5/S11 HTTP | graph-derived process panel | generation/provenance guard | pending |
| R134 | `internal/httpapi/panels.go:(Server).handleProcess` | S5/S11 HTTP | graph-derived process detail | generation/provenance guard | pending |
| R135 | `internal/httpapi/panels.go:(Server).handleClusters` | S5/S11 HTTP | graph-derived cluster panel | generation/provenance guard | pending |
| R136 | `internal/httpapi/panels.go:(Server).handleCluster` | S5/S11 HTTP | graph-derived cluster detail | generation/provenance guard | pending |
| R137 | `internal/httpapi/panels.go:(Server).graphForPanelRequest` | S5 HTTP | Graph JSON shared panel loader | manifest guard before graph open | pending |
| R138 | `internal/httpapi/analyze.go:(Server).handleAnalyzeProgress` | S5 HTTP | analyze job/generation state | job generation/protocol guard | pending |
| R139 | `internal/httpapi/embed.go:(Server).handleEmbedJob` | S5/S9 HTTP | embedding job/generation state | job generation/catalog/hash guard | pending |
| R140 | `internal/httpapi/embed.go:(Server).handleEmbedProgress` | S5/S9 HTTP | embedding progress/generation state | job generation/catalog/hash guard | pending |
| R141 | `internal/httpapi/mcp.go:(Server).handleMCPWithSession` | S5/S4 HTTP-MCP | MCP session transport | handshake/session generation guard before dispatch | pending |
| R142 | `internal/httpapi/mcp.go:(Server).handleMCPGet` | S5/S4 HTTP-MCP | MCP session stream | handshake/session generation guard before stream | pending |
| R143 | `internal/httpapi/mcp.go:(Server).handleMCPPost` | S5/S4 HTTP-MCP | MCP JSON-RPC request | handshake/session generation guard before body dispatch | pending |
| R144 | `internal/httpapi/session.go:(Server).handleSessionChat` | S5 HTTP | session chat/context | repo generation/session guard before context use | pending |
| R145 | `internal/httpapi/search.go:(SearchService).Search` | S5/S9 HTTP selector | dispatcher selecting the concrete semantic/BM25 runners in R152/R153 | generation/catalog/hash guard before runner query; concrete backend rows remain separate | pending |
| R146 | `internal/embeddings/search.go:SemanticSearch` | S9 embeddings | Ladybug embedding/vector rows | generation/catalog/dimension/hash guard | pending |
| R147 | `internal/embeddings/search.go:collectBestChunks` | S9 embeddings | embedding chunk query | generation/catalog/dimension/hash guard | pending |
| R148 | `internal/embeddings/search.go:hydrateSearchResults` | S9 embeddings | metadata/node hydration query | node/generation guard; orphan -> typed no-data error | pending |
| R149 | `internal/group/query.go:loadGroupGraphSnapshot` | S10 group | member Graph JSON snapshot | member-generation vector guard before graph open | pending |
| R150 | `internal/analyze/analyze.go:runEmbeddings` | S9 analyze | existing embedding hashes + embedding store | active generation/catalog/hash guard before reuse | pending |
| R151 | `internal/httpapi/repos.go:(Server).handleRepoDelete` | S5/S10 HTTP | repo-local storage + global registry | active generation/lease guard before delete/unregister | pending |
| R152 | `internal/httpapi/search.go:(SearchService).searchSemantic` | S5/S9 HTTP | embedding runner/vector rows | generation/catalog/dimension/hash guard before query | pending |
| R153 | `internal/httpapi/search.go:searchBM25FromRunner` | S5/S9 HTTP | Ladybug/BM25 runner rows | generation/index guard before query | pending |
| R154 | `internal/lbugnative/runner_ladybugdb.go:(readRunner).QueryRows` | S1 native Cypher | native Ladybug query rows | pinned generation/schema guard before query; mismatch is never converted to `ErrUnavailable` | pending |
| R155 | `internal/mcp/tools.go:runMCPGraphQuery` | S4/S2 MCP fallback Cypher | in-memory Go adapter over the Graph JSON loaded by R66 | fallback-generation guard and explicit unsupported-query error; no native-backend claim | pending |
| R156 | `internal/httpapi/query.go:runGraphPanelQuery` | S5/S2 HTTP fallback Cypher | in-memory Go panel adapter over the Graph JSON loaded by R73 | fallback-generation guard and explicit unsupported-query error; no native-backend claim | pending |
| R157 | `internal/cli/api_command.go:newAPICommand` | S3 dispatcher/non-reader | CLI API parent command registration only | explicit dispatcher classification; route-map/tool-map/shape-check/impact children remain exact rows R100-R103 | pending |
| R158 | `internal/mcp/server.go:(Server).handle` | S4 dispatcher/non-reader | MCP method dispatcher | handshake is established before dispatch; exact tools/resources have their own reader rows | pending |
| R159 | `internal/mcp/server.go:(Server).callTool` | S4 dispatcher/non-reader | MCP tool-name dispatcher | preserve negotiated reader protocol while dispatching; child rows cannot be replaced by this row | pending |
| R160 | `internal/cli/hook_command.go:runClaudeHook` | S3 dispatcher/non-reader | Claude hook event dispatcher | explicit branch classification; pre-tool and post-tool readers are R161/R162 | pending |
| R161 | `internal/cli/hook_command.go:handleClaudePreToolUse` | S3 hook | child `anvien augment` graph-query process for eligible local indexes | validate discovered index protocol/generation before augmentation; mismatch is visible and no stale augmentation text is returned | pending |
| R162 | `internal/cli/hook_command.go:handleClaudePostToolUse` | S3 hook | discovered repo-local `.anvien/meta.json` freshness read | validate metadata protocol/schema/generation before comparing commits; mismatch returns the typed compatibility result | pending |
| R163 | `anvien-web/src/services/backend-client.ts:fetchFromBackend` | S6 transport boundary | shared browser fetch transport | attach the reader handshake to graph/index-dependent requests and preserve HTTP `409` `INDEX_VERSION_MISMATCH`; exact response readers remain separate rows | pending |
| R164 | `anvien-web/src/services/backend-client.ts:fetchServerInfo` | S6 startup/handshake | `/api/info` protocol/manifest response | negotiate reader build/protocol/schema arrays before any graph body or stream is requested | pending |
| R165 | `anvien-web/src/services/backend-client.ts:connectToServer` | S6 startup | repo-info then graph-download orchestration | require R164 handshake and one pinned generation across repo info and graph; mismatch renders the Web error state without graph records | pending |
| R166 | `anvien-web/src/services/backend-client.ts:fetchRepoInfo` | S6/S10 Web | `/api/repo` registry/meta response | negotiated generation/registry epoch guard before accepting repo identity or analysis state | pending |
| R167 | `anvien-web/src/services/backend-client.ts:fetchGraph` | S6 Web graph | `/api/graph` JSON or NDJSON response | validate manifest/generation before JSON decode or delegation to R168; no partial graph on mismatch | pending |
| R168 | `anvien-web/src/services/backend-client.ts:parseNdjsonGraphResponse` | S6 Web graph stream | NDJSON node/relationship/semantic-status stream | require a manifest handshake record before the first graph record; mixed/late/missing manifest aborts and discards partial arrays | pending |
| R169 | `anvien-web/src/services/backend-client.ts:runQuery` | S6/S2 Web query | `/api/query` Go-fallback rows | handshake and response-generation guard before accepting any row | pending |
| R170 | `anvien-web/src/services/backend-client.ts:search` | S6/S9 Web search | `/api/search` native BM25/semantic results | generation/catalog/dimension/hash guard before accepting results | pending |
| R171 | `anvien-web/src/services/backend-client.ts:grep` | S6 Web grep | `/api/grep` indexed/source response | generation and source-commit guard before accepting results | pending |
| R172 | `anvien-web/src/services/backend-client.ts:readFile` | S6 Web file | `/api/file` source response | repo generation/source-commit guard before rendering content | pending |
| R173 | `anvien-web/src/services/backend-client.ts:fetchFileHotspots` | S6/S7 Web | `/api/file-hotspots` Graph JSON/file-context projection | generation/cache-key guard before accepting hotspot rows | pending |
| R174 | `anvien-web/src/services/backend-client.ts:fetchFileContext` | S6/S7 Web | `/api/file-detail` Graph JSON/file-context projection | generation/cache-key guard before accepting context rows | pending |
| R175 | `anvien-web/src/services/backend-client.ts:fetchProcesses` | S6/S11 Web | `/api/processes` graph-derived projection | generation/provenance guard before accepting process membership/order | pending |
| R176 | `anvien-web/src/services/backend-client.ts:fetchProcessDetail` | S6/S11 Web | `/api/processes/{name}` graph-derived detail | generation/provenance guard before accepting the trace | pending |
| R177 | `anvien-web/src/services/backend-client.ts:fetchClusters` | S6/S11 Web | `/api/clusters` graph-derived projection | generation/provenance guard before accepting cluster membership/order | pending |
| R178 | `anvien-web/src/services/backend-client.ts:fetchClusterDetail` | S6/S11 Web | `/api/clusters/{name}` graph-derived detail | generation/provenance guard before accepting cluster nodes | pending |
| R179 | `anvien-web/src/services/backend-client.ts:startAnalyze` | S6 analyze | `/api/analyze` generation lifecycle request | negotiated writer/reader protocol guard before creating a job; returned job is generation-qualified | pending |
| R180 | `anvien-web/src/services/backend-client.ts:getAnalyzeStatus` | S6 analyze | `/api/analyze/{job}` state | job generation/protocol guard before accepting status | pending |
| R181 | `anvien-web/src/services/backend-client.ts:streamAnalyzeProgress` | S6 analyze stream | `/api/analyze/{job}/progress` SSE | handshake/job-generation event precedes progress; mismatch closes the stream without stale completion | pending |
| R182 | `anvien-web/src/services/backend-client.ts:startEmbeddings` | S6/S9 embeddings | `/api/embed` generation lifecycle request | negotiated generation/catalog/hash guard before creating an embedding job | pending |
| R183 | `anvien-web/src/services/backend-client.ts:getEmbedStatus` | S6/S9 embeddings | `/api/embed/{job}` state | job generation/catalog/hash guard before accepting status | pending |
| R184 | `anvien-web/src/services/backend-client.ts:streamEmbeddingProgress` | S6/S9 embeddings stream | `/api/embed/{job}/progress` SSE | handshake/job-generation/catalog event precedes progress; mismatch closes the stream | pending |
| R185 | `anvien-web/src/services/backend-client.ts:deleteRepo` | S6/S10 Web destructive path | `/api/repo` delete plus registry update | active-generation/lease/registry-epoch guard before delete; mismatch cannot be flattened to a generic client error | pending |
| R186 | `anvien-web/src/services/backend-client.ts:fetchRepos` | S6/S10 Web | `/api/repos` global registry response | registry epoch/protocol guard before accepting the list | pending |
| R187 | `anvien-web/src/services/backend-client.ts:probeBackend` | S6/S10 Web probe | `/api/repos` reachability probe | protocol mismatch is distinct from unreachable; HTTP 200 alone cannot mark an incompatible backend usable | pending |
| R188 | `anvien-web/src/services/backend-client.ts:cancelAnalyze` | S6 analyze | `/api/analyze/{job}` cancellation | job generation/protocol guard before cancellation | pending |
| R189 | `anvien-web/src/services/backend-client.ts:cancelEmbeddings` | S6/S9 embeddings | `/api/embed/{job}` cancellation | job generation/catalog/protocol guard before cancellation | pending |
| R190 | `anvien-web/src/services/backend-client.ts:connectHeartbeat` | S6 transport/non-reader | `/api/heartbeat` liveness only | explicit exclusion: heartbeat never proves index compatibility and cannot clear an existing mismatch state | pending |
| R191 | `anvien-web/src/services/backend-client.ts:streamSSE` | S6 stream transport | shared analyze/embed SSE parser | each graph/index stream must receive and validate its typed handshake event before data events; child rows R181/R184 remain authoritative | pending |
| R192 | `anvien-web/src/services/backend-client.ts:pickLocalFolder` | S6 local-control/non-reader | local runtime folder picker only | explicit exclusion: no graph/index body is read and no compatibility success is inferred | pending |
| R193 | `internal/httpapi/repos.go:(Server).handleRepo` | S5/S10 dispatcher/non-reader | HTTP repo route dispatcher selecting repo-info or repo-delete child handlers | validate transport/epoch before dispatch; R79 and R151 remain the concrete read/delete rows | pending |
| R194 | `internal/httpapi/session.go:(Server).handleSession` | S5 session dispatcher | HTTP session-cancel route dispatcher | validate session/repo generation before cancellation; child status/chat rows remain separate | pending |
| R195 | `internal/httpapi/mcp.go:(Server).handleMCPPostRaw` | S5/S4 HTTP-MCP response boundary | JSON-RPC response body or SSE envelope after MCP server dispatch | negotiate protocol/session and generation before writing any response body; no partial stream on mismatch | pending |

P2-A1 must attach the source-scan command/output, matrix SHA-256, exact `rows_classified/rows_total`, `unassigned_rows`, and `unlisted_readers` counts to `E2-P2A1-MATRIXREVIEW1`. If the source scan finds a reader not represented above, add a new exact row before the slice can pass; do not restore a generic catch-all row. P2-A1 must also record one exact later guard-owner slice for every row. P2-E2 records the final `rows_passed/rows_total` result. The S3/S4/S5/S6 surface definitions in the plan and parity ledgers must equal the union of these exact rows.
