# AVmatrix Multi-Language Graph Filters and Coverage Benchmark Ledger

Date: 2026-05-19

Status: active

Companion files:

- Plan: [2026-05-19-avmatrix-multilanguage-graph-filters-and-coverage-plan.md](2026-05-19-avmatrix-multilanguage-graph-filters-and-coverage-plan.md)
- Evidence ledger: [2026-05-19-avmatrix-multilanguage-graph-filters-and-coverage-evidence.md](2026-05-19-avmatrix-multilanguage-graph-filters-and-coverage-evidence.md)

## Benchmark Rules

Record benchmarkable results when they measure product/runtime performance, capacity, package/startup size, graph/DB throughput, or conversion-inventory counts. Build/test/e2e timings are validation evidence unless the slice changes those systems.

For this plan, benchmarkable measurements include:

- raw graph relationship counts for `EXTENDS`, `INHERITS`, and `IMPLEMENTS`;
- unique semantic heritage source-target pair counts;
- duplicate compatibility edge counts;
- full node label and relationship type filter inventory counts;
- supported-language graph coverage matrix counts;
- TypeScript heritage source-site counts, preferring parser/ScopeIR inventory over regex after implementation starts;
- per-language source-site/extraction/resolution/display coverage for each supported graph fact family;
- resolved, unresolved, and missing heritage coverage counts;
- Go embedded-struct heritage counts and final display labels;
- Web dashboard displayed edge counts versus raw graph counts;
- Web dashboard displayed node counts versus raw graph counts;
- graph adapter label/type classification coverage;
- graph adapter relationship preservation/collapse counts;
- edge visibility and focus-depth filter behavior;
- top bar Back navigation target and stale connection-loss banner behavior during intentional navigation;
- left dashboard resize min/max width bounds and graph/canvas usable width after resize;
- analyze runtime if analyzer behavior is changed.

## B0 - Initial Restaurant_manager Graph And Heritage Baseline

Date: 2026-05-19

Graph snapshot:

- repo path: `E:\Restaurant_manager`
- graph path: `E:\Restaurant_manager\.avmatrix\graph.json`
- nodes: `78,350`
- relationships: `130,497`

### Graph Filter Contract Baseline

Generated contract counts come from `internal/contracts/web_ui.go` and the generated Web contract consumed by `avmatrix-web/src/lib/constants.ts`.

| Metric | Baseline |
|---|---:|
| Generated Web contract node labels | `37` |
| Generated Web contract relationship types | `22` |
| Supported code languages | `18` |
| Loaded graph node label inventory | pending measured |
| Loaded graph relationship type inventory | pending measured |
| UI node filter rows vs graph payload labels | pending measured |
| UI edge filter rows vs graph payload relationship types | pending measured |
| Graph adapter label/type classification coverage | pending measured |

### Node Inventory Relevant To The Question

| Node label | Count | Source audit result |
|---|---:|---|
| Class | `3` | Matches source app class declarations |
| Constructor | `3` | Matches constructors on the 3 source app classes |
| Interface | `587` | Many TS interface nodes exist |
| Struct | `946` | Go-heavy domain/backend model surface |
| Function | `5,659` | Function-heavy TS/Go app style |
| Method | `2,687` | Method-heavy Go/backend plus TS class methods |
| Section | `35,488` | Documentation/report sections dominate graph volume |

Verified `Class` nodes:

| Class | File |
|---|---|
| `SSEListener` | `electron/main/sync/sse-listener.ts` |
| `ApiError` | `electron/renderer/src/api/client.ts` |
| `ErrorBoundary` | `electron/renderer/src/components/shared/ErrorState/ErrorBoundary.tsx` |

Verified `Constructor` nodes:

| Constructor owner | File |
|---|---|
| `SSEListener` | `electron/main/sync/sse-listener.ts` |
| `ApiError` | `electron/renderer/src/api/client.ts` |
| `ErrorBoundary` | `electron/renderer/src/components/shared/ErrorState/ErrorBoundary.tsx` |

### Heritage Relationship Inventory

| Metric | Count |
|---|---:|
| Raw `EXTENDS` relationships | `6` |
| Raw `INHERITS` relationships | `6` |
| Raw `IMPLEMENTS` relationships | `0` |
| Unique `EXTENDS`/`INHERITS` source-target pairs | `6` |
| Source-target pairs duplicated as both `EXTENDS` and `INHERITS` | `6` |
| TS class/interface heritage edges found in graph | `0` |

Current duplicate pairs:

| Source | Target | Raw edge types |
|---|---|---|
| `Struct:backend/internal/domain/event.go:EventStoreEntry` | `Struct:backend/internal/domain/event.go:EventEnvelope` | `EXTENDS`, `INHERITS` |
| `Struct:backend/internal/domain/shift/models.go:AssignmentWithUser` | `Struct:backend/internal/domain/shift/models.go:AssignmentWithAttendance` | `EXTENDS`, `INHERITS` |
| `Struct:backend/internal/repo/cash_repo.go:CashRepository` | `Struct:backend/internal/repo/cash_repo.go:CashSessionRepoSQL` | `EXTENDS`, `INHERITS` |
| `Struct:backend/internal/repo/cash_repo.go:CashRepository` | `Struct:backend/internal/repo/cash_repo.go:CashLedgerRepoSQL` | `EXTENDS`, `INHERITS` |
| `Struct:backend/internal/repo/cash_repo.go:CashRepository` | `Struct:backend/internal/repo/cash_repo.go:TransferLedgerRepoSQL` | `EXTENDS`, `INHERITS` |
| `Struct:backend/internal/repo/cash_repo.go:CashRepository` | `Struct:backend/internal/repo/cash_repo.go:RevenueTotalsRepoSQL` | `EXTENDS`, `INHERITS` |

### TypeScript Heritage Source-Site Baseline

Source-site search scope:

- include: `*.ts`, `*.tsx`
- exclude: `node_modules`, `.git`, `.avmatrix`, `reports`, `Docs`, `dist`, `build`

| Source-site type | Count |
|---|---:|
| `interface ... extends ...` | `14` |
| `class ... extends ...` | `2` |
| Total audited TS heritage source sites | `16` |
| Current TS heritage graph relationships found | `0` |

Observed `class extends` sites:

| Source site | Target text |
|---|---|
| `electron/renderer/src/api/client.ts:1` | `ApiError extends Error` |
| `electron/renderer/src/components/shared/ErrorState/ErrorBoundary.tsx:14` | `ErrorBoundary extends Component<Props, State>` |

Observed `interface extends` examples:

| Source site | Target text |
|---|---|
| `electron/renderer/src/types/area.ts:14` | `AreaWithTableCount extends Area` |
| `electron/renderer/src/features/tables/types.ts:65` | `TableWithUser extends Table` |
| `electron/renderer/src/types/table.ts:21` | `TableWithUser extends Table` |
| `electron/renderer/src/features/shifts/types.ts:86` | `ShiftWithCounts extends Shift` |
| `electron/renderer/src/features/shifts/types.ts:134` | `AssignmentWithUser extends ShiftAssignment` |
| `electron/renderer/src/features/shifts/types.ts:329` | `ShiftWithCountsDTO extends ShiftDTO` |

## B1 - Target Metrics To Record During Implementation

Status: pending

Record after each relevant slice:

| Metric | Baseline | Target |
|---|---:|---:|
| Raw `EXTENDS` count | `6` | measured |
| Raw `INHERITS` count | `6` | measured |
| Unique semantic heritage pair count | `6` | measured |
| Duplicate compatibility pair count | `6` | `0` for UI display, or clearly marked if raw graph preserves both |
| Generated node labels in Web contract | `37` | all classified for filter/display/adapter behavior |
| Generated graph relationship types in Web contract | `22` | all classified for filter/display/adapter behavior |
| Supported code languages in scanner/Web contracts | `18` | all classified in graph coverage matrix |
| Supported languages with extractor/provider status | pending | all classified as provider-backed, script-container backed, scan-only, or not applicable |
| Supported languages with graph fact coverage | pending | all classified by fact family and graph policy |
| Languages with provider parity fixtures | current provider parity subset | all provider-backed languages covered by claimed fact families or explicitly documented as not applicable |
| Per-language source-to-graph contract | pending | no unclassified language entries |
| UI node filter rows vs graph payload labels | pending | no graph-present label missing from dashboard/legend |
| UI edge filter rows vs graph payload relationships | pending | no graph-present relationship type missing from dashboard/legend |
| Graph adapter label/type classification coverage | pending | all labels/types covered or intentional fallback documented |
| AVmatrix-GO full graph filter inventory | pending | node/edge/filter/adapter counts recorded |
| Restaurant_manager full graph filter inventory | partial heritage baseline | node/edge/filter/adapter counts recorded |
| TS heritage source sites | `16` | measured |
| TS parser/ScopeIR heritage source sites | pending | measured |
| TS resolved heritage edges | `0` | all resolvable in-repo sites |
| TS unresolved/external heritage sites represented | `0` | all unresolved/external sites represented or audited |
| TS missing heritage source sites | `16` | `0` |
| Go embedded struct unique semantic pairs | `6` | measured |
| Go embedded struct user-facing duplicate pairs | `6` through `EXTENDS` + `INHERITS` | `0` misleading duplicates |
| Top bar Back navigation to `Start-AVmatrix.html` | not implemented in current product | visible and tested |
| Back navigation stale connection-loss banner behavior | pending | no false connection-loss banner during intentional navigation |
| Left dashboard drag resize | not implemented in current product | measured min/max width bounds |
| Canvas usability after dashboard resize | pending | measured/validated |

## B2 - Final Benchmark

Status: pending

Record final:

- AVmatrix-GO analyze counts and runtime;
- Restaurant_manager analyze counts and runtime;
- raw relationship counts;
- semantic unique heritage counts;
- duplicate compatibility counts;
- full Web graph filter inventory counts;
- supported-language graph coverage matrix;
- per-language provider/extractor status;
- per-language source-site, extraction, resolution, graph relationship/node, and Web display/filter classification;
- TS parser/ScopeIR heritage source-site coverage;
- Go embedded struct heritage display behavior;
- Web dashboard displayed node and edge counts;
- graph adapter classification coverage;
- edge visibility and focus-depth behavior;
- top bar Back navigation target and stale connection-loss banner behavior;
- left dashboard min/max resize bounds and canvas usable width;
- e2e observed display behavior.
