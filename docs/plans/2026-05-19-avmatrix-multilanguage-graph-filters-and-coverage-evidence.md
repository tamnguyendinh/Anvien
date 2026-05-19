# AVmatrix Multi-Language Graph Filters and Coverage Evidence Ledger

Date: 2026-05-19

Status: active

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

## E6 - Implementation Evidence

Status: pending

Record each implementation slice here:

- files changed;
- AVmatrix context/impact results;
- build/test/e2e commands;
- Restaurant_manager analyze results;
- AVmatrix-GO analyze results;
- final graph and Web UI behavior.

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
