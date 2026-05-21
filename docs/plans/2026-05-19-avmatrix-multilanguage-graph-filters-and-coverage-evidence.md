# AVmatrix Multi-Language Graph Filters and Coverage Evidence Ledger

Date: 2026-05-19

Status: complete - zero-trust follow-up closure recorded

Companion files:

- Plan: [2026-05-19-avmatrix-multilanguage-graph-filters-and-coverage-plan.md](2026-05-19-avmatrix-multilanguage-graph-filters-and-coverage-plan.md)
- Benchmark ledger: [2026-05-19-avmatrix-multilanguage-graph-filters-and-coverage-benchmark.md](2026-05-19-avmatrix-multilanguage-graph-filters-and-coverage-benchmark.md)

## Evidence Rules

Record evidence as each evidenced task is completed. Evidence should include commands, impacted files, test results, e2e artifacts, and concise observations needed to audit the plan later.

For doc-only commits, do not use AVmatrix.

## E0 - Plan Creation Evidence

Date: 2026-05-19

Created file set:

- original file set: `docs/plans/2026-05-19-avmatrix-heritage-edge-semantics-and-coverage-{plan,benchmark,evidence}.md`
- current file set after scope correction: `docs/plans/2026-05-19-avmatrix-multilanguage-graph-filters-and-coverage-{plan,benchmark,evidence}.md`

Reason:

The previous Web UI graph display work exposed a deeper graph semantics question. `E:\Restaurant_manager` shows `EXTENDS=6` and `INHERITS=6`, but those counts are the same `6` source-target pairs emitted twice. The same repo also contains TypeScript heritage source sites that are not represented by TS heritage graph edges.

Doc-only note:

- This plan file creation did not run AVmatrix. It records evidence already gathered during the preceding investigation.

## E1 - Initial Graph Inventory Evidence

Date: 2026-05-19

Command:

```powershell
$g = Get-Content -Raw -LiteralPath 'E:\Restaurant_manager\.avmatrix\graph.json' | ConvertFrom-Json
$g.nodes | Group-Object label | Sort-Object Name
$g.relationships | Group-Object type | Sort-Object Name
```

Result summary:

```text
nodes: 78,350
relationships: 130,497

Class: 3
Constructor: 3
Interface: 587
Struct: 946
Function: 5,659
Method: 2,687
Section: 35,488

EXTENDS: 6
INHERITS: 6
IMPLEMENTS: 0
```

Conclusion:

- `Class=3` and `Constructor=3` are not UI size values. They are real graph counts.
- The counts are plausible after source audit because this repo uses mostly Go structs, TS interfaces, TS function components, and functions, not many TS classes.

## E2 - Class and Constructor Source Verification

Date: 2026-05-19

Command:

```powershell
rg -n --glob '*.ts' --glob '*.tsx' --glob '*.js' --glob '*.jsx' "^\s*export\s+(default\s+)?class\s+|^\s*class\s+" E:\Restaurant_manager\electron E:\Restaurant_manager\shared E:\Restaurant_manager\scripts
rg -n --glob '*.ts' --glob '*.tsx' --glob '*.js' --glob '*.jsx' "^\s*constructor\s*\(" E:\Restaurant_manager\electron E:\Restaurant_manager\shared E:\Restaurant_manager\scripts
```

Class results:

```text
electron/main/sync/sse-listener.ts:28: export class SSEListener
electron/renderer/src/api/client.ts:1: export class ApiError extends Error
electron/renderer/src/components/shared/ErrorState/ErrorBoundary.tsx:14: export class ErrorBoundary extends Component<Props, State>
```

Constructor results:

```text
electron/main/sync/sse-listener.ts:34: constructor(mainWindow: BrowserWindow)
electron/renderer/src/api/client.ts:5: constructor(message: string, status: number = 500, code: string | null = null)
electron/renderer/src/components/shared/ErrorState/ErrorBoundary.tsx:15: constructor(props: Props)
```

Conclusion:

- Current `Class=3` and `Constructor=3` match the audited source declarations.
- This is not the same problem as the `EXTENDS` / `INHERITS` heritage issue.

## E3 - Duplicate EXTENDS / INHERITS Pair Evidence

Date: 2026-05-19

Command:

```powershell
$g = Get-Content -Raw -LiteralPath 'E:\Restaurant_manager\.avmatrix\graph.json' | ConvertFrom-Json
$g.relationships |
  Where-Object { $_.type -in @('EXTENDS','INHERITS') } |
  Group-Object sourceId,targetId |
  Select-Object Count,
    @{n='Types';e={($_.Group.type|Sort-Object)-join','}},
    @{n='Source';e={$_.Group[0].sourceId}},
    @{n='Target';e={$_.Group[0].targetId}}
```

Result summary:

```text
6 source-target pairs.
Each pair has both EXTENDS and INHERITS.
No unique pair has only one of those two edge types.
```

Conclusion:

- The Web UI must not present `EXTENDS=6` and `INHERITS=6` as if they are `12` independent codebase relationships.
- Either the graph payload or the UI display layer needs a semantic grouping/collapse policy.

## E4 - Code Path Trace Evidence

Date: 2026-05-19

Commands:

```powershell
.\avmatrix-launcher\server-bundle\avmatrix.exe context RelExtends --repo AVmatrix
.\avmatrix-launcher\server-bundle\avmatrix.exe context HeritageExtends --repo AVmatrix
rg -n "emitHeritageCompatibilityEdges|ReferenceInherits|relationshipTypeForReference|HeritageExtends|extends_clause" internal
```

Relevant code paths:

- `internal/providers/tsjs/references.go`
  - `extends_clause` creates `scopeir.HeritageExtends`.
  - `implements_clause` creates `scopeir.HeritageImplements`.
- `internal/providers/golang/references.go`
  - embedded fields and interface type elements create `scopeir.HeritageExtends`.
- `internal/resolution/resolve.go`
  - `emitInherits := !options.DisableScopeInheritsCompatibility`
  - each `w.heritage` item is passed into `emitHeritageCompatibilityEdges`.
- `internal/resolution/emit.go`
  - `emitHeritageCompatibilityEdges` emits `EXTENDS` or `IMPLEMENTS`.
  - when `emitInherits` is true, it emits `ReferenceInherits`.
  - `relationshipTypeForReference(ReferenceInherits)` maps to `INHERITS`.
- `internal/mcp/resources.go`
  - docs currently describe `INHERITS` as normalized scope-resolved inheritance or heritage dependency.
  - docs describe `EXTENDS` as class inheritance.

Conclusion:

- `EXTENDS` and `INHERITS` are intentionally both emitted today for compatibility/normalized dependency reasons.
- That may be valid for raw graph consumers, but it is misleading in a user-facing graph dashboard unless grouped, labelled, or otherwise explained.

## E5 - TypeScript Heritage Coverage Evidence

Date: 2026-05-19

Commands:

```powershell
rg -n --glob '*.ts' --glob '*.tsx' --glob '!node_modules/**' --glob '!.git/**' --glob '!.avmatrix/**' --glob '!reports/**' --glob '!Docs/**' --glob '!dist/**' --glob '!build/**' "^\s*(export\s+)?interface\s+\w+\s+extends\s+" E:\Restaurant_manager

rg -n --glob '*.ts' --glob '*.tsx' --glob '!node_modules/**' --glob '!.git/**' --glob '!.avmatrix/**' --glob '!reports/**' --glob '!Docs/**' --glob '!dist/**' --glob '!build/**' "^\s*(export\s+)?class\s+\w+\s+extends\s+" E:\Restaurant_manager
```

Result summary:

```text
interface extends sites: 14
class extends sites: 2
current graph TS heritage edges involving audited Class/Interface endpoints: 0
```

Observed TS class extends sites:

```text
electron/renderer/src/api/client.ts:1: ApiError extends Error
electron/renderer/src/components/shared/ErrorState/ErrorBoundary.tsx:14: ErrorBoundary extends Component<Props, State>
```

Observed TS interface extends examples:

```text
electron/renderer/src/types/area.ts:14: AreaWithTableCount extends Area
electron/renderer/src/features/tables/types.ts:65: TableWithUser extends Table
electron/renderer/src/types/table.ts:21: TableWithUser extends Table
electron/renderer/src/features/shifts/types.ts:86: ShiftWithCounts extends Shift
electron/renderer/src/features/shifts/types.ts:134: AssignmentWithUser extends ShiftAssignment
electron/renderer/src/features/shifts/types.ts:329: ShiftWithCountsDTO extends ShiftDTO
```

Conclusion:

- TS provider has code to emit heritage facts, but `E:\Restaurant_manager` graph does not show TS heritage relationships for the audited TS sites.
- The implementation plan must determine whether the loss is in extraction, owner-scope resolution, import/name resolution, external target handling, or graph payload/UI filtering.

## E6 - Implementation Slice Evidence

Date: 2026-05-19

Status: recorded

### AVmatrix-Assisted Checks

Commands:

```powershell
avmatrix analyze --force
avmatrix context FileTreePanel --repo AVmatrix
avmatrix impact FileTreePanel --repo AVmatrix --direction upstream --depth 2 --include-tests
avmatrix context knowledgeGraphToGraphology --repo AVmatrix
avmatrix impact knowledgeGraphToGraphology --repo AVmatrix --direction upstream --depth 2 --include-tests
avmatrix context Header --repo AVmatrix
avmatrix impact Header --repo AVmatrix --direction upstream --depth 2 --include-tests
avmatrix context resolveHeritage --repo AVmatrix
avmatrix impact resolveHeritage --repo AVmatrix --direction upstream --depth 2 --include-tests
avmatrix context emitReferenceKind --repo AVmatrix
avmatrix impact emitReferenceKind --repo AVmatrix --direction upstream --depth 2 --include-tests
```

Observed impact summary:

- `FileTreePanel` changes affect `App.tsx` and dashboard completeness tests.
- `knowledgeGraphToGraphology` changes affect `GraphCanvas.tsx` and graph-adapter geometry tests.
- `Header` changes are local to app shell wiring and branding tests.
- `resolveHeritage` is high-impact because it feeds relationship emission and downstream graph consumers.
- `emitReferenceKind` is the TS/JS extraction entry point for heritage target collection.

### Implementation Files

Backend and contract files:

- `internal/providers/tsjs/references.go`
- `internal/providers/tsjs/extract_test.go`
- `internal/resolution/indexes.go`
- `internal/resolution/resolve.go`
- `internal/resolution/types.go`
- `internal/resolution/resolution_test.go`
- `internal/contracts/web_ui.go`
- `internal/contracts/web_ui_test.go`
- `contracts/web-ui/avmatrix-web-contract.schema.json`

Web files:

- `avmatrix-web/src/generated/avmatrix-contracts.ts`
- `avmatrix-web/src/lib/constants.ts`
- `avmatrix-web/src/lib/graph-adapter.ts`
- `avmatrix-web/src/lib/lucide-icons.tsx`
- `avmatrix-web/src/components/FileTreePanel.tsx`
- `avmatrix-web/src/components/Header.tsx`
- `avmatrix-web/src/App.tsx`
- `avmatrix-web/vite.config.ts`

Tests:

- `avmatrix-web/test/unit/constants.test.ts`
- `avmatrix-web/test/unit/FileTreePanel.dashboard-completeness.test.tsx`
- `avmatrix-web/test/unit/graph-adapter.edge-geometry.test.ts`
- `avmatrix-web/test/unit/Branding.local-only.test.tsx`
- `avmatrix-web/e2e/shell-interactions.spec.ts`

### Analyzer Changes

Implementation summary:

- TS/JS extraction now handles `extends_type_clause` and collects multiple heritage targets from interface `extends` clauses.
- Resolution now records `HeritageFactsIndexed` and `UnresolvedInheritance`.
- Resolution now uses `baseTypeName` for heritage target lookup.
- Resolution now falls back to a same-file target search when global heritage lookup is ambiguous. This fixes real TS files where names such as `Area` or `Table` are defined in multiple files but the local `interface ... extends ...` target is in the same file.
- Raw graph compatibility behavior is preserved: resolved heritage still emits both the specific edge (`EXTENDS` or `IMPLEMENTS`) and `INHERITS` when compatibility mode is enabled.

Focused tests added:

- TS provider extraction for interface extends, multiple extends, and generic extends.
- TS resolution for same-file interface extends plus unresolved external generic heritage.
- TS resolution for same-file target preference when another file defines the same interface name.

### Web Display Changes

Implementation summary:

- Generated Web contract now includes:
  - `RELATIONSHIP_DISPLAY_POLICY`
  - `LANGUAGE_GRAPH_COVERAGE`
- Web constants now include full generated relationship size policy coverage and explicit structural/community node label classifications.
- `INHERITS` is labelled `Normalized Heritage`.
- Dashboard counts and legend titles group duplicate compatibility `INHERITS` edges when the same source-target pair also has `EXTENDS` or `IMPLEMENTS`.
- Graph adapter conversion uses the display relationship set so duplicate compatibility `INHERITS` does not draw as a second unrelated edge.

Shell implementation summary:

- Header has an icon-first Back button labelled `Back to Start screen`.
- Back target resolves to `/Start-AVmatrix.html` on the current origin.
- Vite dev/build serves or emits `Start-AVmatrix.html`.
- App suppresses the reconnect banner during intentional navigation to the Start screen.
- Left dashboard has a drag handle, persists width in `localStorage`, and clamps width to `192px` through `480px`.

### Graph Measurements

Commands:

```powershell
go run ./cmd/avmatrix analyze --force [redacted removed argument] --no-stats
go run ./cmd/avmatrix analyze E:\Restaurant_manager --force [redacted removed argument] --no-stats
```

Results:

```text
AVmatrix-GO analyze runtime: 17.89s
AVmatrix-GO graph: nodes=20,771 relationships=51,854
AVmatrix-GO graph-present labels=16 relationship types=11
AVmatrix-GO heritage: raw=0 uniquePairs=0 duplicateCompatibilityPairs=0

Restaurant_manager analyze runtime: 28.93s
Restaurant_manager graph: nodes=78,358 relationships=130,588
Restaurant_manager raw heritage: EXTENDS=19 INHERITS=19 IMPLEMENTS=0
Restaurant_manager unique semantic heritage pairs=19
Restaurant_manager duplicate compatibility pairs=19 raw, 0 misleading UI display duplicates
Restaurant_manager TS heritage: raw=16 relationships, 8 unique resolved source-target pairs
Restaurant_manager display relationships after grouping=130,569
```

Restaurant_manager TS trigger conclusion:

- Resolved in-repo TS source-target pairs now appear in the graph:
  - `AssignmentWithUser -> ShiftAssignment`
  - `ShiftWithCounts -> Shift`
  - `ShiftWithCountsDTO -> ShiftDTO`
  - both `TableWithUser -> Table` declarations
  - `AreaWithTableCount -> Area`
  - `DateTimeOptions -> DateOptions`
  - `DateTimeOptions -> TimeOptions`
- External TS targets such as `Error`, `React.*HTMLAttributes`, DOM `Performance*`, and React `Component<Props, State>` remain unresolved/external by policy.

### Validation Commands

Passed:

```powershell
go run .\cmd\generate-web-contracts
go test ./internal/providers/tsjs ./internal/resolution ./internal/contracts
go build ./cmd/... ./internal/...
go test ./cmd/... ./internal/...
npm --prefix avmatrix-web run build
npm --prefix avmatrix-web test -- --run
npm --prefix avmatrix-web run test:e2e -- shell-interactions.spec.ts -g "back button|resizes the left dashboard" --workers=1 --timeout=120000
```

Result summary:

- Focused Go tests passed.
- Full applicable Go build passed.
- Full applicable Go tests for `cmd` and `internal` passed.
- Web production build passed and emitted `dist/Start-AVmatrix.html`.
- Full Vitest suite passed: `41` files, `322` tests.
- Focused Playwright e2e passed: `2/2`.

Additional note:

- `go test ./...` was also attempted. It fails on pre-existing non-buildable fixture packages under `avmatrix/test/fixtures` such as fixture imports from `models`/`animal`, C source files without cgo, and intentionally invalid Go examples. The applicable Go validation scope for this slice is `./cmd/... ./internal/...`.

## E7 - Plan Review Evidence

Date: 2026-05-19

Review result:

- The original plan pointed at the correct high-level problem: duplicated heritage semantics plus missing TS heritage coverage.
- The original plan needed sharper ordering before implementation. It could jump from baseline to "fix provider/resolution" without first forcing a source-to-graph trace for each missing TS heritage class.
- The plan now requires tracing every missing TS site through:

```text
source AST -> ScopeIR HeritageFact -> workspace heritage resolution -> graph relationship -> Web payload -> dashboard/canvas display
```

Additional corrections added:

- final source-site accuracy must prefer parser/ScopeIR inventory over regex counts;
- Go embedded structs must be classified explicitly, because showing Go embedding as `EXTENDS` may be misleading even if the underlying heritage edge is useful;
- MCP/context/impact/MRO compatibility must be tested if `INHERITS` raw graph or graph payload semantics change;
- generated contracts/schema docs must be updated if the selected policy adds relationship metadata, external target facts, or display-group fields;
- UI tests must cover both duplicate `EXTENDS` + `INHERITS` pairs and TS resolved/unresolved heritage display.

## E8 - Web UI Usability Requirements Moved Into Active Plan

Date: 2026-05-19

Correction:

- the top bar Back navigation and left dashboard resize requirements were originally attached to the completed Web UI dashboard plan by mistake;
- those requirements belong in this active multi-language graph filters/display plan because the current implementation work is already touching Web UI relationship display and graph-inspection workflow;
- the completed Web UI dashboard plan remains historical closure and is not reopened here.

Requirements now carried by this plan:

- add a Back arrow/button beside the `AVmatrix` top bar title to return to `Start-AVmatrix.html`;
- avoid showing a stale connection-loss banner during intentional Back navigation;
- make the left dashboard resizable by dragging its right boundary;
- enforce min/max width bounds and verify the graph canvas remains usable after resizing.

Doc-only note:

- no AVmatrix analysis was run for this documentation correction.

## E9 - Supported-Language Contract Scope Correction

Date: 2026-05-19

Correction:

- the graph filter/fact fix cannot be bounded to TypeScript, Go, or heritage relationships because AVmatrix declares a multi-language code surface;
- TypeScript missing heritage and Go duplicate `EXTENDS` / `INHERITS` are trigger examples, not the full acceptance boundary;
- the plan now requires a supported-language graph coverage matrix before implementation can be considered complete.

Supported code-language surface identified from scanner/Web contracts:

- `javascript`
- `typescript`
- `python`
- `java`
- `c`
- `cpp`
- `csharp`
- `go`
- `ruby`
- `rust`
- `php`
- `kotlin`
- `swift`
- `dart`
- `vue`
- `svelte`
- `astro`
- `cobol`

Required classification for each language:

- provider/extractor status;
- node labels and relationship types supported by the provider or dedicated analyzer phase;
- source fact families supported by the language and provider, including definitions, imports, calls, uses/type refs, member/property/method relationships, accesses, heritage-like facts, routes/tools/process facts where applicable;
- ScopeIR extraction behavior;
- graph resolution behavior;
- unresolved/external representation policy;
- Web filter/display label/grouping policy;
- fixture/e2e evidence status.

Doc-only note:

- no AVmatrix analysis was run for this documentation correction.

## E10 - AVmatrix-Assisted Codebase Audit For Multi-Language Graph Filters

Date: 2026-05-19

Commands:

```powershell
avmatrix status
avmatrix analyze
avmatrix analyze --force
```

Result:

- `avmatrix status` reported the index was stale.
- `avmatrix analyze` first failed during native LadybugDB schema initialization with `File already exists in catalog`.
- `avmatrix analyze --force` succeeded:

```text
files: scanned=691 parsed=530 unsupported=161 failed=0
graph: nodes=19714 relationships=47169 path=E:\AVmatrix-GO\.avmatrix\graph.json
```

AVmatrix context/impact checks:

- `FileTreePanel` is used by `avmatrix-web/src/App.tsx` and `avmatrix-web/test/unit/FileTreePanel.dashboard-completeness.test.tsx`.
- `knowledgeGraphToGraphology` is used by `avmatrix-web/src/components/GraphCanvas.tsx` and `avmatrix-web/test/unit/graph-adapter.edge-geometry.test.ts`.
- `getGraphEdgeVisibilityMode` is used by `avmatrix-web/src/hooks/useSigma.ts` and `avmatrix-web/test/unit/graph-edge-visibility-mode.test.ts`.
- `WebUIContract` feeds generated Web graph/language constants and is covered by contract tests.
- `parseFiles` routes supported language extraction and is covered by parse routing tests for many provider languages.

Code audit findings:

- `internal/contracts/web_ui.go` is the generated Web contract source for `NODE_LABELS`, `GRAPH_RELATIONSHIP_TYPES`, and `codeLanguages`.
- `avmatrix-web/src/lib/constants.ts` imports generated graph contracts and defines display behavior: default visible labels, colors, sizes, edge labels/colors, sort order, and unknown fallbacks.
- `avmatrix-web/src/components/FileTreePanel.tsx` builds loaded-graph node/edge filter rows from graph-present labels/types, not from every contract row; no-graph fallback uses the full generated lists.
- `avmatrix-web/src/lib/graph-edge-visibility-mode.ts` and `avmatrix-web/src/hooks/useSigma.ts` do apply `visibleEdgeTypes` to rendered edges through the Sigma edge reducer.
- `avmatrix-web/src/lib/graph-adapter.ts` still contains hand-classified graph behavior: `structuralTypes`, `symbolTypes`, hierarchy relationship priorities, node mass rules, edge size multipliers, node size caps, and community-color symbol handling. Those sets are not proven complete for all generated node labels and relationship types.
- `internal/analyze/analyze.go` declares `hasExtractor` for every code language except `cobol`; Vue/Svelte/Astro use script-container extraction; Cobol is scanned and processed by a separate analyzer phase rather than ScopeIR provider extraction.
- `internal/providers/provider_parity_test.go` has broad but uneven provider parity coverage. Some fact families are covered across many languages, while owner/member and heritage-like forms have clear subsets and must be expanded or explicitly classified.
- Current e2e coverage includes uncommon `Property` / `Accesses` dashboard toggles, but does not prove the full generated node-label/relationship-type contract or per-language provider matrix.

Conclusion:

- The plan must be driven by the concrete chain `contract -> provider/dedicated analyzer -> resolution -> graph payload -> FileTreePanel -> graph adapter/useSigma -> e2e`, not by a generic statement that the tool is multi-language.
- The central acceptance target is no unclassified language, node label, relationship type, or UI filter behavior.

## E11 - Plan Alignment Correction After Review

Date: 2026-05-19

Correction:

- the plan rules were not changed;
- scope now includes `avmatrix-web/src/lib/graph-edge-visibility-mode.ts` and `avmatrix-web/src/hooks/useSigma.ts`, because edge visibility is enforced through the graph edge visibility mode and Sigma reducer, not only through dashboard row generation;
- Phase 3 now starts from the provider/fact-family matrix for every supported language before TS/Go trigger fixes;
- TS missing heritage and Go duplicate heritage remain concrete trigger cases, but they no longer define the full analyzer coverage boundary;
- Phase 5 and Phase 6 now require full graph filter inventory and final graph/filter verification for both `E:\AVmatrix-GO` and `E:\Restaurant_manager`.

Doc-only note:

- no AVmatrix analysis was run for this documentation correction; it uses the AVmatrix-assisted audit evidence recorded in E10.

## E12 - MCP Policy And Provider Heritage Parity Slice

Date: 2026-05-19

Status: recorded

### AVmatrix-Assisted Checks

Commands:

```powershell
go run ./cmd/avmatrix context schemaResource --repo AVmatrix
go run ./cmd/avmatrix impact schemaResource --repo AVmatrix --direction upstream --depth 2 --include-tests
```

Observed impact summary:

- `schemaResource` is consumed by `Server.readResourceText`.
- upstream impact is limited to the MCP module with risk reported as `LOW`.
- a context/impact lookup for the newly added provider parity test name was attempted before reindexing and returned `Target not found`; this is expected because the test symbol was not in the existing graph snapshot yet.

### Implementation Files

Backend files:

- `internal/mcp/resources.go`
- `internal/mcp/server_test.go`
- `internal/providers/provider_parity_test.go`

Plan files:

- `docs/plans/2026-05-19-avmatrix-multilanguage-graph-filters-and-coverage-plan.md`
- `docs/plans/2026-05-19-avmatrix-multilanguage-graph-filters-and-coverage-benchmark.md`
- `docs/plans/2026-05-19-avmatrix-multilanguage-graph-filters-and-coverage-evidence.md`

### Policy Changes

MCP schema resource now records:

- raw graph compatibility policy: preserve `EXTENDS`/`IMPLEMENTS` and `INHERITS` for Cypher, MCP context/impact, and MRO consumers;
- user display policy: group same-pair compatibility `INHERITS` with `EXTENDS` or `IMPLEMENTS`;
- standalone `INHERITS` policy: draw/count only when no matching `EXTENDS` or `IMPLEMENTS` exists for the same source-target pair;
- language heritage terminology for TypeScript, Go, Java/C#/Kotlin/Dart/PHP, and Python/C++/Swift/Ruby/Rust forms;
- unresolved/external policy: do not synthesize graph nodes/edges for missing package, DOM, React, standard-library, or otherwise external targets.

### Provider Parity Changes

`internal/providers/provider_parity_test.go` now covers heritage extraction for:

- TypeScript
- Go
- Python
- C++
- Ruby
- Java
- C#
- Kotlin
- Rust
- PHP
- Dart
- Swift

It also covers parser/ScopeIR-to-resolution graph heritage emission for:

- TypeScript
- Go
- Python
- Java
- C#
- Kotlin
- C++
- PHP
- Ruby
- Rust
- Dart
- Swift

Each graph-resolution case asserts both the language-specific relationship (`EXTENDS` or `IMPLEMENTS`) and the compatibility `INHERITS` relationship.

### Validation Commands

Passed:

```powershell
gofmt -w internal\providers\provider_parity_test.go internal\mcp\resources.go internal\mcp\server_test.go
go test ./internal/providers ./internal/mcp ./internal/mro ./internal/resolution
go build ./cmd/... ./internal/...
go test ./cmd/... ./internal/...
```

Result summary:

- focused provider/MCP/MRO/resolution tests passed;
- full applicable Go build passed;
- full applicable Go tests for `cmd` and `internal` passed;
- no Web files changed in this slice, so Web build/e2e were not rerun for this policy/provider-parity-only change.

Conclusion:

- raw graph compatibility behavior is preserved for MCP/context/impact/MRO consumers;
- the user-facing duplicate heritage grouping policy is now explicitly documented in MCP schema resources;
- provider heritage parity is no longer TS/Go-only in representative extraction and graph-resolution fixtures;
- broader non-heritage provider fact-family parity and full UI e2e filter/focus-depth expansion were later reopened and closed by E18/P8 with explicit representative proof-level wording.

## E13 - UI Filter, Legend, Focus-Depth, And Large-Graph Smoke Coverage Slice

Date: 2026-05-19

Status: recorded

### AVmatrix-Assisted Checks

Commands:

```powershell
go run ./cmd/avmatrix context FileTreePanel --repo AVmatrix
go run ./cmd/avmatrix impact FileTreePanel --repo AVmatrix --direction upstream --depth 2 --include-tests
```

Observed impact summary:

- `FileTreePanel` is used by `avmatrix-web/src/App.tsx` and `avmatrix-web/test/unit/FileTreePanel.dashboard-completeness.test.tsx`.
- upstream depth-2 impact also reaches `avmatrix-web/src/main.tsx` through `App.tsx`;
- risk reported as `LOW`.

### Implementation Files

Web test files:

- `avmatrix-web/test/unit/FileTreePanel.dashboard-completeness.test.tsx`
- `avmatrix-web/test/unit/constants.test.ts`
- `avmatrix-web/e2e/shell-interactions.spec.ts`

Plan files:

- `docs/plans/2026-05-19-avmatrix-multilanguage-graph-filters-and-coverage-plan.md`
- `docs/plans/2026-05-19-avmatrix-multilanguage-graph-filters-and-coverage-benchmark.md`
- `docs/plans/2026-05-19-avmatrix-multilanguage-graph-filters-and-coverage-evidence.md`

### Coverage Added

Unit coverage added:

- loaded-graph mode hides zero-count generated contract labels/types and displays only graph-present rows;
- graph-present loaded-mode rows still include counts and legend rows;
- generated `LANGUAGE_GRAPH_COVERAGE` records resolved in-repo and unresolved/external policies for representative provider-backed languages;
- generated coverage records script-container heritage for Vue and dedicated analyzer-phase behavior for COBOL;
- `RELATIONSHIP_DISPLAY_POLICY` records `INHERITS` as `Normalized Heritage` with grouped display policy;
- focus-depth controls call app state and show the no-selection warning.

E2E coverage added:

- opens the Filters tab;
- verifies Node Types, Edge Types, Focus Depth, and Color Legend sections;
- verifies graph-present `File` and `Calls` count rows;
- verifies legend rows for `File` and `Calls`;
- toggles the `Calls` relationship type off and on;
- selects `2 hops`, verifies the focus-depth warning with no selected node, then clears it with `All`.

Large-graph smoke context:

```powershell
Invoke-WebRequest -UseBasicParsing -Uri 'http://127.0.0.1:4848/api/repos'
```

Result summary:

```text
Restaurant_manager: nodes=78,358 edges=130,588
AVmatrix: nodes=20,771 edges=51,854
```

The focused Playwright shell e2e selects the first repo returned by the backend. In this environment that repo is `Restaurant_manager`, so the new filter/legend/focus-depth smoke test ran against the large Restaurant_manager graph.

### Validation Commands

Passed:

```powershell
npm --prefix avmatrix-web test -- --run test/unit/constants.test.ts test/unit/FileTreePanel.dashboard-completeness.test.tsx
npm --prefix avmatrix-web run test:e2e -- shell-interactions.spec.ts -g "displays graph filters" --workers=1 --timeout=120000
npm --prefix avmatrix-web run build
npm --prefix avmatrix-web test -- --run
npm --prefix avmatrix-web run test:e2e -- shell-interactions.spec.ts -g "back button|resizes the left dashboard|displays graph filters" --workers=1 --timeout=120000
```

Result summary:

- focused Vitest passed: `2` files, `32` tests;
- new focused Playwright e2e passed: `1/1`;
- Web production build passed and emitted `dist/Start-AVmatrix.html`;
- full Vitest suite passed: `41` files, `325` tests;
- focused Playwright shell coverage passed: `3/3`.

Conclusion:

- zero-count contract rows are intentionally hidden in loaded-graph mode because graph-present rows come directly from the loaded graph payload;
- graph policy and unresolved/external coverage are visible in generated contract tests;
- filter sections, legend rows, relationship toggles, focus-depth warning/clear behavior, Back navigation, and dashboard resize now have focused e2e coverage on the local large graph.

## E14 - Restaurant_manager TypeScript Heritage ScopeIR Trace Slice

Date: 2026-05-19

Status: recorded

### Implementation Files

Backend test files:

- `internal/providers/tsjs/extract_test.go`

Plan files:

- `docs/plans/2026-05-19-avmatrix-multilanguage-graph-filters-and-coverage-plan.md`
- `docs/plans/2026-05-19-avmatrix-multilanguage-graph-filters-and-coverage-benchmark.md`
- `docs/plans/2026-05-19-avmatrix-multilanguage-graph-filters-and-coverage-evidence.md`

### Trace Commands

Source-site inventory:

```powershell
rg -n --glob '*.ts' --glob '*.tsx' --glob '!node_modules/**' --glob '!.git/**' --glob '!.avmatrix/**' --glob '!reports/**' --glob '!Docs/**' --glob '!dist/**' --glob '!build/**' "^\s*(export\s+)?interface\s+\w+\s+extends\s+|^\s*(export\s+)?class\s+\w+\s+extends\s+" E:\Restaurant_manager
```

ScopeIR trace validation:

```powershell
$env:AVMATRIX_RESTAURANT_MANAGER_ROOT='E:\Restaurant_manager'
go test ./internal/providers/tsjs -run TestExtractRestaurantManagerTypeScriptHeritageSites -count=1 -v
```

Default-suite safety check:

```powershell
Remove-Item Env:\AVMATRIX_RESTAURANT_MANAGER_ROOT -ErrorAction SilentlyContinue
go test ./internal/providers/tsjs -run TestExtractRestaurantManagerTypeScriptHeritageSites -count=1 -v
go test ./internal/providers/tsjs ./internal/resolution
```

Graph output trace:

```powershell
$g = Get-Content -Raw -LiteralPath 'E:\Restaurant_manager\.avmatrix\graph.json' | ConvertFrom-Json
$g.relationships |
  Where-Object { $_.type -in @('EXTENDS','INHERITS','IMPLEMENTS') -and $_.sourceId -like 'Interface:electron/renderer/src/*' } |
  Sort-Object sourceId,targetId,type |
  ForEach-Object { "$($_.type)`t$($_.sourceId)`t$($_.targetId)" }
```

Import/source context trace:

```powershell
rg -n "^\s*import\s+|extends\s+" <each audited Restaurant_manager TS/TSX file>
```

### ScopeIR Findings

The optional TS provider trace test reads the audited Restaurant_manager source files and verifies `17` raw `HeritageFact` target facts across `13` files:

| File | Heritage targets |
|---|---|
| `electron/renderer/src/utils/performance.ts` | `PerformanceEntry`, `Performance` |
| `electron/renderer/src/utils/dateUtils.ts` | `DateOptions`, `TimeOptions` |
| `electron/renderer/src/types/table.ts` | `Table` |
| `electron/renderer/src/types/area.ts` | `Area` |
| `electron/renderer/src/features/tables/types.ts` | `Table` |
| `electron/renderer/src/components/shared/Form/FormTextarea.tsx` | `React.TextareaHTMLAttributes<HTMLTextAreaElement>` |
| `electron/renderer/src/components/shared/Form/FormSelect.tsx` | `React.SelectHTMLAttributes<HTMLSelectElement>` |
| `electron/renderer/src/components/shared/Form/FormInput.tsx` | `React.InputHTMLAttributes<HTMLInputElement>` |
| `electron/renderer/src/components/shared/Form/FormCheckbox.tsx` | `React.InputHTMLAttributes<HTMLInputElement>` |
| `electron/renderer/src/components/shared/ErrorState/ErrorBoundary.tsx` | `Component` |
| `electron/renderer/src/components/shared/Button/Button.tsx` | `React.ButtonHTMLAttributes<HTMLButtonElement>` |
| `electron/renderer/src/api/client.ts` | `Error` |
| `electron/renderer/src/features/shifts/types.ts` | `Shift`, `ShiftAssignment`, `ShiftDTO` |

### Graph Findings

Resolved TS heritage pairs in the final Restaurant_manager graph:

```text
AssignmentWithUser -> ShiftAssignment
ShiftWithCounts -> Shift
ShiftWithCountsDTO -> ShiftDTO
TableWithUser -> Table
AreaWithTableCount -> Area
TableWithUser -> Table
DateTimeOptions -> DateOptions
DateTimeOptions -> TimeOptions
```

Each resolved pair appears as raw `EXTENDS` plus compatibility `INHERITS`.

### Cross-File / Imported Target Classification

Result:

- no audited TS heritage source site has an in-repo cross-file/imported target that failed import binding;
- resolved in-repo targets are same-file interface targets;
- imported or global heritage targets are external/built-in symbols from React, DOM, JavaScript, or browser performance APIs:
  - `React.*HTMLAttributes<...>`
  - `Component<Props, State>`
  - `Error`
  - `PerformanceEntry`
  - `Performance`
- these targets follow the unresolved/external policy: do not synthesize graph nodes or resolved graph edges to external packages or platform APIs.

Conclusion:

- the original missing TS graph edges were not extraction losses after the fix: every audited source site now emits ScopeIR `HeritageFact` data;
- graph relationships are emitted for resolved in-repo targets;
- unresolved imported/global targets are policy-classified external targets, not silent missing graph facts.

## E15 - Provider Non-Heritage Fact-Family Coverage Audit

Date: 2026-05-19

Status: recorded - scope limited by zero-trust reopen

### Audit Commands

```powershell
rg "func Test.*ScopeIRParityFixture|func TestResolve.*GraphParity|func Test.*GraphParity|func Test.*Provider.*Parity" internal\providers -g "*_test.go" -n
rg --files internal\providers | Select-String -Pattern "scopeir_signature\.golden\.json$"
rg "func TestResolve.*GraphParity|func Test.*GraphParity" internal\providers -g "*_test.go" -n
```

### Findings

Provider-specific ScopeIR golden fixtures exist for:

- C
- C++
- C#
- Dart
- Go
- Java
- Kotlin
- PHP
- Python
- Ruby
- Rust
- Swift
- TypeScript/JavaScript
- Vue

Provider/script-container graph parity count tests exist for:

- C
- C++
- C#
- Dart
- Go
- Java
- Kotlin
- PHP
- Python
- Ruby
- Rust
- Swift
- Vue
- Svelte
- Astro

Centralized provider parity tests cover:

- calls/forms/receivers/arity across representative provider-backed languages;
- imports across TS/JS, C#, Go, Java, Kotlin, Python, Ruby, Rust, C++, Dart, and Swift;
- owner/member extraction across TS, Python, Java, C#, Go, Rust, and C++;
- heritage extraction and graph resolution across TS, Go, Python, Java, C#, Kotlin, C++, PHP, Ruby, Rust, Dart, and Swift.

Source fact families covered by these suites:

- definitions;
- imports;
- calls;
- accesses;
- type refs/type bindings/type annotations;
- member/property/method ownership;
- heritage relationships;
- graph resolution counts and representative relationship emission.

Route/tool/process relationships are intentionally not provider parity fixtures. They are covered by route, process, tool, MCP, framework, and graph-resolution tests because those facts are framework/analyzer-derived rather than direct language-provider extraction.

Conclusion:

- non-TS/Go provider facts are not limited to a small subset; they are represented by provider-specific golden fixtures plus graph parity count tests;
- script-container providers have graph parity count tests for embedded JS/TS extraction;
- the original audit treated the remaining work as closure validation and commit, but E17 reopens the proof level because some evidence is representative/count-level rather than explicit per-language/per-fact classification.

## E16 - Final Closure Evidence

Date: 2026-05-19

Status: recorded - superseded by zero-trust reopen

### Final Validation Commands

Passed:

```powershell
go build ./cmd/... ./internal/...
go test ./cmd/... ./internal/...
```

Previously passed in the active implementation slices and included in final benchmark:

```powershell
$env:AVMATRIX_RESTAURANT_MANAGER_ROOT='E:\Restaurant_manager'
go test ./internal/providers/tsjs -run TestExtractRestaurantManagerTypeScriptHeritageSites -count=1 -v
npm --prefix avmatrix-web run build
npm --prefix avmatrix-web test -- --run
npm --prefix avmatrix-web run test:e2e -- shell-interactions.spec.ts -g "back button|resizes the left dashboard|displays graph filters" --workers=1 --timeout=120000
```

Result summary:

- Go build passed for `./cmd/... ./internal/...`;
- Go tests passed for `./cmd/... ./internal/...`;
- optional Restaurant_manager TS ScopeIR trace passed when `AVMATRIX_RESTAURANT_MANAGER_ROOT` was set;
- Web production build passed;
- full Vitest suite passed with `41` files and `325` tests;
- focused Playwright shell e2e passed `3/3`;
- final graph and UI benchmark inventories are recorded in B5.

Superseded closure conclusion:

- graph-present filters, legends, relationship counts, and display rows are aligned with generated contracts and graph payload reality;
- raw graph compatibility semantics are preserved for downstream consumers, while Web display groups duplicate same-pair heritage compatibility edges;
- `Restaurant_manager` TS heritage source sites emit ScopeIR `HeritageFact` data, resolved in-repo targets become graph edges, and external/imported platform targets are classified by unresolved/external policy;
- supported-language coverage matrix and provider fact-family evidence were previously treated as having no remaining unclassified language entries, but E17 reopens the fact-family detail proof;
- Back navigation, left dashboard resize, graph filter display, legend display, relationship toggles, and focus-depth behavior have unit/e2e evidence;
- all completed implementation slices have been committed or are included in the final closure commit.

## E17 - Zero-Trust Reopen Review

Date: 2026-05-19

Status: recorded

Doc-only note:

- no AVmatrix analysis was run for this review/update, per plan rule 6;
- this section reopens closure claims using local file inspection and previously recorded evidence.

### Review Finding Summary

High:

- `LANGUAGE_GRAPH_COVERAGE` is too generic for provider-backed languages. `internal/contracts/web_ui.go` uses shared provider-backed `SourceFactFamilies` values such as definitions/imports/calls/accesses/type-references/members/heritage-where-language-supports-it instead of explicit per-language/per-fact statuses. This does not fully satisfy the original P1-J/P3-B/P3-C/P7-E closure criteria.

Medium:

- The `Restaurant_manager` large-graph e2e smoke is not deterministic. `avmatrix-web/e2e/shell-interactions.spec.ts` chooses the first repository returned by `/api/repos`; the previous evidence only proves that `Restaurant_manager` happened to be first in that environment.
- `TestExtractRestaurantManagerTypeScriptHeritageSites` is useful audit evidence, but it is optional unless `AVMATRIX_RESTAURANT_MANAGER_ROOT` is set. It is not currently a default regression gate.
- The plan was marked complete while benchmark/evidence statuses and historical text still contained active/pending closure language.
- Provider graph parity proof is partly count-level and representative. That may be acceptable if stated honestly, but it is not endpoint-level proof for every claimed source fact kind.

### Reopen Decision

The previous final closure is rescinded. The implementation evidence remains valid for the original trigger fixes and representative coverage, but the plan cannot claim full zero-trust multi-language graph coverage until Phase 8 closes these gaps.

Checklist items reopened in the plan:

- P1-J
- P3-B
- P3-C
- P3-D
- P3-J
- P4-K
- P5-E
- P6-K
- P7-E

New follow-up requirements are tracked under Phase 8:

- explicit per-language/per-fact graph coverage matrix;
- contract tests preventing generic provider-backed coverage inheritance;
- deterministic `Restaurant_manager` or equivalent large-graph e2e;
- default or committed-fixture regression coverage for the `17` audited TypeScript heritage facts;
- provider parity proof strengthened or wording narrowed to representative/count-level evidence;
- benchmark/evidence drift cleanup;
- final validation and zero-trust closure.

## E18 - Zero-Trust Follow-Up Implementation Slice

Date: 2026-05-19

Status: recorded

### AVmatrix-Assisted Checks

Commands:

```powershell
go run .\cmd\avmatrix status
go run .\cmd\avmatrix analyze --force [redacted removed argument] --no-stats
go run .\cmd\avmatrix context providerCoverage --repo AVmatrix
go run .\cmd\avmatrix impact providerCoverage --repo AVmatrix --direction upstream --depth 3 --include-tests
go run .\cmd\avmatrix context WebUIContract --repo AVmatrix
go run .\cmd\avmatrix impact WebUIContract --repo AVmatrix --direction upstream --depth 2 --include-tests
```

Result summary:

- initial AVmatrix status was stale after the doc-only reopen commit;
- re-analyze succeeded with `files=691`, `parsed=530`, `nodes=20,881`, and `relationships=51,998`;
- `providerCoverage` impact was `LOW` and limited to the contracts module;
- `WebUIContract` impact was `CRITICAL` because it feeds generated schema/TypeScript artifacts and contract tests.

### Implementation Files

Backend and contracts:

- `internal/contracts/web_ui.go`
- `internal/contracts/web_ui_test.go`
- `contracts/web-ui/avmatrix-web-contract.schema.json`
- `internal/providers/provider_parity_test.go`
- `internal/providers/tsjs/extract_test.go`

Web:

- `avmatrix-web/src/generated/avmatrix-contracts.ts`
- `avmatrix-web/test/unit/constants.test.ts`
- `avmatrix-web/e2e/shell-interactions.spec.ts`

Plan files:

- `docs/plans/2026-05-19-avmatrix-multilanguage-graph-filters-and-coverage-plan.md`
- `docs/plans/2026-05-19-avmatrix-multilanguage-graph-filters-and-coverage-benchmark.md`
- `docs/plans/2026-05-19-avmatrix-multilanguage-graph-filters-and-coverage-evidence.md`

### Coverage Changes

- `LANGUAGE_GRAPH_COVERAGE` now includes explicit `factFamilies`, `supportedNodeLabels`, `supportedRelationshipTypes`, fixture coverage, proof-level metadata, default regression gate, and optional external audit fields.
- Provider-backed languages no longer use the generic `heritage-where-language-supports-it` source fact marker.
- Contract tests now require explicit fact-family status for generated graph coverage and fail if provider-backed languages fall back to the old generic heritage marker.
- The generated contract records `18` language entries and `141` explicit fact-family rows.
- Provider parity wording is narrowed to representative endpoint/count-level proof where that is the real evidence level.
- `TestProviderGraphParityEndpointProofCoversRepresentativeNonTSGoFacts` adds endpoint assertions for C and Java definitions, members, calls, accesses, and type-use relationships.
- `TestExtractRestaurantManagerTypeScriptHeritageFixture` adds a committed/default fixture covering all `17` audited Restaurant_manager TypeScript heritage target facts.
- The external `AVMATRIX_RESTAURANT_MANAGER_ROOT` trace remains available as audit evidence and passed in this environment.
- Playwright shell e2e now selects `E2E_REPO_NAME` or default `Restaurant_manager` by stable repo name/path and skips with an explicit reason if that deterministic target is unavailable.

### Validation Commands

Passed:

```powershell
go test ./internal/contracts ./internal/providers/tsjs ./internal/providers -run "TestWebUIContract|TestExtractRestaurantManagerTypeScriptHeritageFixture|TestProviderGraphParityEndpointProof" -count=1
go run .\cmd\generate-web-contracts
npm --prefix avmatrix-web test -- --run test/unit/constants.test.ts
go build ./cmd/... ./internal/...
go test ./cmd/... ./internal/...
npm --prefix avmatrix-web run build
npm --prefix avmatrix-web test -- --run
npm --prefix avmatrix-web run test:e2e -- shell-interactions.spec.ts -g "back button|resizes the left dashboard|displays graph filters" --workers=1 --timeout=120000
$env:AVMATRIX_RESTAURANT_MANAGER_ROOT='E:\Restaurant_manager'; go test ./internal/providers/tsjs -run TestExtractRestaurantManagerTypeScriptHeritageSites -count=1 -v
```

Result summary:

- focused Go tests passed for contracts, TS heritage fixture, and provider endpoint proof;
- full applicable Go build passed;
- full applicable Go tests for `cmd` and `internal` passed;
- Web production build passed;
- full Vitest suite passed: `41` files and `325` tests;
- focused Playwright shell e2e passed `3/3` against deterministic `Restaurant_manager`;
- optional external Restaurant_manager TS heritage trace passed and verified all `17` target facts.

Final conclusion:

- the zero-trust reopen issues are closed at the stated proof level;
- language graph coverage is explicit rather than generic;
- deterministic regression gates now cover the large graph e2e selection and the Restaurant_manager TS heritage target facts;
- provider parity evidence no longer overclaims every endpoint where only representative/count-level evidence exists;
- historical baseline `pending` rows remain benchmark history, while active closure status is recorded in B6 and this section.
