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

## B1 - Implementation Slice Benchmark

Status: recorded

Measurement date: 2026-05-19

Analyzer command used for final graph measurements:

```powershell
go run ./cmd/avmatrix analyze --force --skip-agents-md --no-stats
go run ./cmd/avmatrix analyze E:\Restaurant_manager --force --skip-agents-md --no-stats
```

Analyze runtime:

| Repo | Runtime |
|---|---:|
| `E:\AVmatrix-GO` | `17.89s` |
| `E:\Restaurant_manager` | `28.93s` |

### Contract And Coverage Matrix

| Metric | Final |
|---|---:|
| Generated Web contract node labels | `37` |
| Generated Web contract relationship types | `22` |
| Relationship display policy entries | `22` |
| Supported code languages | `18` |
| Language graph coverage entries | `18` |
| Provider-backed languages | `14` |
| Script-container-backed languages | `3` |
| Dedicated analyzer-phase languages | `1` |

Provider-backed languages are JavaScript, TypeScript, Python, Java, C, C++, C#, Go, Ruby, Rust, PHP, Kotlin, Swift, and Dart. Vue, Svelte, and Astro are script-container-backed. COBOL is classified as a dedicated analyzer phase.

### AVmatrix-GO Final Graph Inventory

| Metric | Count |
|---|---:|
| Nodes | `20,771` |
| Relationships | `51,854` |
| Graph-present node labels | `16` |
| Graph-present relationship types | `11` |
| Raw heritage relationships | `0` |
| Unique semantic heritage pairs | `0` |
| Duplicate compatibility pairs | `0` |
| Display relationships after compatibility grouping | `51,854` |
| Display relationship types after compatibility grouping | `11` |

Node counts:

| Node label | Count |
|---|---:|
| Class | `4` |
| Community | `912` |
| Const | `323` |
| Constructor | `5` |
| File | `691` |
| Folder | `112` |
| Function | `3,402` |
| Interface | `100` |
| Method | `809` |
| Package | `413` |
| Process | `639` |
| Property | `3,205` |
| Section | `988` |
| Struct | `503` |
| TypeAlias | `74` |
| Variable | `8,591` |

Relationship counts:

| Relationship type | Raw count | Display count |
|---|---:|---:|
| ACCESSES | `5,078` | `5,078` |
| CALLS | `8,523` | `8,523` |
| CONTAINS | `1,766` | `1,766` |
| DEFINES | `17,429` | `17,429` |
| ENTRY_POINT_OF | `639` | `639` |
| HAS_METHOD | `339` | `339` |
| HAS_PROPERTY | `2,862` | `2,862` |
| IMPORTS | `3,733` | `3,733` |
| MEMBER_OF | `3,884` | `3,884` |
| STEP_IN_PROCESS | `2,364` | `2,364` |
| USES | `5,237` | `5,237` |

### Restaurant_manager Final Trigger Inventory

| Metric | Baseline | Final |
|---|---:|---:|
| Nodes | `78,350` | `78,358` |
| Relationships | `130,497` | `130,588` |
| Raw `EXTENDS` count | `6` | `19` |
| Raw `INHERITS` count | `6` | `19` |
| Raw `IMPLEMENTS` count | `0` | `0` |
| Unique semantic heritage pair count | `6` | `19` |
| Duplicate compatibility pair count | `6` | `19` raw, `0` misleading display duplicates |
| TS heritage raw relationships | `0` | `16` |
| TS heritage unique source-target pairs | `0` | `8` |
| Go embedded struct unique semantic pairs | `6` | `11` |
| Display relationship count after compatibility grouping | pending | `130,569` |
| Display relationship type count after compatibility grouping | pending | `13` |

Final graph-present node labels: `17`.

Final graph-present relationship types: `14` raw, `13` displayed after grouping duplicate compatibility `INHERITS`.

Final Restaurant_manager TS heritage pairs:

| Source | Target | Raw edge types |
|---|---|---|
| `Interface:electron/renderer/src/features/shifts/types.ts:AssignmentWithUser` | `Interface:electron/renderer/src/features/shifts/types.ts:ShiftAssignment` | `EXTENDS`, `INHERITS` |
| `Interface:electron/renderer/src/features/shifts/types.ts:ShiftWithCounts` | `Interface:electron/renderer/src/features/shifts/types.ts:Shift` | `EXTENDS`, `INHERITS` |
| `Interface:electron/renderer/src/features/shifts/types.ts:ShiftWithCountsDTO` | `Interface:electron/renderer/src/features/shifts/types.ts:ShiftDTO` | `EXTENDS`, `INHERITS` |
| `Interface:electron/renderer/src/features/tables/types.ts:TableWithUser` | `Interface:electron/renderer/src/features/tables/types.ts:Table` | `EXTENDS`, `INHERITS` |
| `Interface:electron/renderer/src/types/area.ts:AreaWithTableCount` | `Interface:electron/renderer/src/types/area.ts:Area` | `EXTENDS`, `INHERITS` |
| `Interface:electron/renderer/src/types/table.ts:TableWithUser` | `Interface:electron/renderer/src/types/table.ts:Table` | `EXTENDS`, `INHERITS` |
| `Interface:electron/renderer/src/utils/dateUtils.ts:DateTimeOptions` | `Interface:electron/renderer/src/utils/dateUtils.ts:DateOptions` | `EXTENDS`, `INHERITS` |
| `Interface:electron/renderer/src/utils/dateUtils.ts:DateTimeOptions` | `Interface:electron/renderer/src/utils/dateUtils.ts:TimeOptions` | `EXTENDS`, `INHERITS` |

Final Restaurant_manager Go embedded struct heritage pairs: `11`, displayed as `EXTENDS` in the Web dashboard with grouped `INHERITS` raw compatibility counts.

Restaurant_manager raw relationship counts:

| Relationship type | Raw count | Display count |
|---|---:|---:|
| ACCESSES | `7,670` | `7,670` |
| CALLS | `15,109` | `15,109` |
| CONTAINS | `42,016` | `42,016` |
| DEFINES | `34,335` | `34,335` |
| ENTRY_POINT_OF | `510` | `510` |
| EXTENDS | `19` | `19` |
| HANDLES_ROUTE | `50` | `50` |
| HAS_METHOD | `2,566` | `2,566` |
| HAS_PROPERTY | `7,520` | `7,520` |
| IMPORTS | `1,612` | `1,612` |
| INHERITS | `19` | `0` when grouped with same-pair `EXTENDS` |
| MEMBER_OF | `6,610` | `6,610` |
| STEP_IN_PROCESS | `2,017` | `2,017` |
| USES | `10,535` | `10,535` |

### Shell Interaction Measurements

| Metric | Final |
|---|---|
| Top bar Back navigation target | `/Start-AVmatrix.html` on the current origin |
| Back navigation false reconnect banner | `0` banners in focused Playwright e2e |
| Left dashboard min width | `192px` |
| Left dashboard default width | `248px` |
| Left dashboard max width | `480px` |
| Resize persistence | `localStorage` key `avmatrix.leftPanelWidth` |
| Canvas usability after resize | focused Playwright e2e opened filters and confirmed canvas visibility |

### Original Target Table Status

| Metric | Baseline | Target |
|---|---:|---:|
| Raw `EXTENDS` count | `6` | `19` in Restaurant_manager final graph |
| Raw `INHERITS` count | `6` | `19` in Restaurant_manager final graph |
| Unique semantic heritage pair count | `6` | `19` in Restaurant_manager final graph |
| Duplicate compatibility pair count | `6` | `19` raw, `0` misleading display duplicates |
| Generated node labels in Web contract | `37` | `37`, with generated dashboard rows and adapter fallback/classification |
| Generated graph relationship types in Web contract | `22` | `22`, with generated display policy and size policy test coverage |
| Supported code languages in scanner/Web contracts | `18` | `18`, all classified in generated graph coverage matrix |
| Supported languages with extractor/provider status | pending | `14` provider-backed, `3` script-container-backed, `1` dedicated analyzer-phase |
| Supported languages with graph fact coverage | pending | classified in generated coverage matrix |
| Languages with provider parity fixtures | current provider parity subset | all provider-backed languages covered by claimed fact families or explicitly documented as not applicable |
| Per-language source-to-graph contract | pending | no unclassified generated language entries |
| UI node filter rows vs graph payload labels | pending | graph-present labels covered by passing unit tests |
| UI edge filter rows vs graph payload relationships | pending | graph-present relationship types covered by passing unit tests |
| Graph adapter label/type classification coverage | pending | generated relationship types have size policy; node classifications have fallback |
| AVmatrix-GO full graph filter inventory | pending | recorded |
| Restaurant_manager full graph filter inventory | partial heritage baseline | recorded |
| TS heritage source sites | `16` | `16` source sites, `17` target facts |
| TS parser/ScopeIR heritage source sites | pending | focused parser/ScopeIR extraction covered in tests |
| TS resolved heritage edges | `0` | `8` unique source-target pairs, `16` raw compatibility edges |
| TS unresolved/external heritage sites represented | `0` | `9` external target facts audited as unresolved/external policy |
| TS missing heritage source sites | `16` | `0` for resolvable in-repo target pairs; external targets remain unresolved by policy |
| Go embedded struct unique semantic pairs | `6` | `11` |
| Go embedded struct user-facing duplicate pairs | `6` through `EXTENDS` + `INHERITS` | `0` misleading display duplicates |
| Top bar Back navigation to `Start-AVmatrix.html` | not implemented in current product | visible and tested |
| Back navigation stale connection-loss banner behavior | pending | no false connection-loss banner during intentional navigation |
| Left dashboard drag resize | not implemented in current product | `192px` to `480px`, tested |
| Canvas usability after dashboard resize | pending | validated by focused Playwright e2e |

## B2 - MCP Policy And Provider Heritage Parity Coverage

Status: recorded

Measurement date: 2026-05-19

This slice did not change analyzer output counts or Web runtime behavior. Benchmarkable inventory is coverage count expansion for policy and provider heritage parity.

### Policy Coverage Inventory

| Metric | Count |
|---|---:|
| MCP schema raw/display heritage policy sections | `3` |
| MCP schema unresolved/external policy sections | `2` |
| MCP schema language heritage terminology groups | `5` |
| Relationship compatibility consumers explicitly preserved | `4` |

Compatibility consumers explicitly preserved: Cypher, MCP context, MCP impact, and MRO.

### Provider Heritage Parity Inventory

| Metric | Count |
|---|---:|
| Languages with representative heritage extraction parity | `12` |
| Languages with representative heritage graph-resolution parity | `12` |
| Resolution cases asserting specific `EXTENDS`/`IMPLEMENTS` relationship | `16` |
| Resolution cases asserting compatibility `INHERITS` relationship | `16` |

Languages with representative heritage extraction parity: TypeScript, Go, Python, C++, Ruby, Java, C#, Kotlin, Rust, PHP, Dart, and Swift.

Languages with representative heritage graph-resolution parity: TypeScript, Go, Python, Java, C#, Kotlin, C++, PHP, Ruby, Rust, Dart, and Swift.

### Validation Inventory

| Command | Result |
|---|---|
| `go test ./internal/providers ./internal/mcp ./internal/mro ./internal/resolution` | passed |
| `go build ./cmd/... ./internal/...` | passed |
| `go test ./cmd/... ./internal/...` | passed |

## B3 - Final Benchmark

Status: pending broader provider-parity/focus-depth expansion

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
