# Anvien File-Centric Graph Projection Evidence Ledger

Date: 2026-05-30

Status: In progress

Companion files:

- Plan: [2026-05-30-anvien-file-centric-graph-projection-plan.md](2026-05-30-anvien-file-centric-graph-projection-plan.md)
- Benchmark ledger: [2026-05-30-anvien-file-centric-graph-projection-benchmark.md](2026-05-30-anvien-file-centric-graph-projection-benchmark.md)

## Evidence Rules

1. Record facts that explain why each task is correct.
2. Keep benchmark tables in the benchmark ledger, not here.
3. For code changes, record impact/blast-radius before edits.
4. For graph-based validation, record the graph refresh command and graph inventory summary.
5. For API or contract changes, record route/tool/shape impact and contract regeneration evidence.
6. For Web UI changes, record full build, unit tests, e2e tests, and any screenshot/browser validation if used.
7. Record failures and the fix or decision that handled them.
8. Record `anvien detect-changes --repo Anvien --scope all` before each implementation commit.
9. Record commit hashes as closure evidence.

## Evidence Template

Use this template for each implementation slice:

```text
## E<n> - <Phase/Task Title>

Date:

Status:

Scope:

- ...

Impact / blast radius:

| Command | Result |
|---|---|
| ... | ... |

Implementation evidence:

| File | Evidence |
|---|---|
| ... | ... |

Validation:

| Command | Result |
|---|---|
| ... | ... |

Failures / handling:

- ...

Detect changes:

| Command | Result |
|---|---|
| `anvien detect-changes --repo Anvien --scope all` | ... |

Commit:

- `<hash> <subject>`
```

## E0 - User Problem And Direction

Date: 2026-05-30

Status: recorded

User direction:

- Keep the current symbol-centric graph model as the source of truth.
- Add a file-centric projection layer so users can inspect graph facts from a file-first perspective.
- The desired view is:

```text
File
  -> summary
  -> symbol tree
  -> relationships
  -> unresolved source sites
  -> linked flows/routes/tools/tests
  -> quality signals
```

Problem evidence from discussion:

- Current symbol graph is strong for exact symbol context, impact, rename, detect-changes, and source-site proof.
- Current inspection is weaker when the user starts from a file and asks what it contains, who depends on it, what it depends on, where unresolved sites are, and which flows/tests touch it.
- The proposed solution is a projection derived from existing graph facts, not a replacement for symbol-level graph ownership.

Planning evidence:

| Check | Result |
|---|---|
| Plan file naming | `2026-05-30-anvien-file-centric-graph-projection-plan.md` uses ISO date and lowercase kebab-case slug. |
| Evidence file naming | `2026-05-30-anvien-file-centric-graph-projection-evidence.md` shares the same slug. |
| Benchmark file naming | `2026-05-30-anvien-file-centric-graph-projection-benchmark.md` shares the same slug. |
| Doc-only planning rule | No Anvien graph command is required for creating this initial doc-only plan set. |

## E1 - Baseline Graph Schema Discovery

Date: 2026-05-30

Status: completed

Readiness review evidence:

| Check | Result |
|---|---|
| Graph refresh | `anvien analyze --force --name Anvien` completed. |
| Graph inventory | `files: scanned=819 parsed=584 unsupported=235 failed=0`; `nodes=91586 relationships=125053`; graph path `.anvien/graph.json`. |
| CLI ownership inspected | Existing command owners include `internal/cli/command.go`, `internal/cli/tool_command.go`, `internal/cli/api_command.go`, and graph-quality command files. |
| MCP ownership inspected | Existing tool owners include `internal/mcp/server.go`, `internal/mcp/context.go`, `internal/mcp/impact.go`, `internal/mcp/tools.go`, `internal/mcp/route_tool_map.go`, and `internal/mcp/route_shape_impact.go`. |
| Graph facts inspected | `internal/graph/types.go` already carries file path and source-site fields on graph nodes/relationships; graph-health inputs already include file/source-site metadata. |
| Web contract ownership inspected | Web contract source is owned by `internal/contracts/web_ui.go` and generated through `cmd/generate-web-contracts`; generated TypeScript lives under `anvien-web/src/generated`. |
| AI context ownership inspected | Generated guidance is owned by `internal/aicontext/aicontext.go` and embedded skill source files under `internal/aicontext/skills`. |

P0-A graph refresh:

| Command | Result |
|---|---|
| `.\anvien\bin\anvien.exe analyze --force --name Anvien` | Pass. `files: scanned=819 parsed=584 unsupported=235 failed=0`; `nodes=91587 relationships=125054`; graph path `.anvien/graph.json`. |
| `.\anvien\bin\anvien.exe graph-health summary --repo Anvien --json` | Pass. Indexed commit and current commit both `cdbd4af19b867b1ed4a3efc2d6c9779f25907ce3`; `resolutionGapNodeCount=65652`; `hasResolutionGapRelationshipCount=65652`; `sourceBackedUnresolvedReferenceCount=66555`; `unattributedUnresolvedReferenceCount=0`. |

P0-A graph facts:

| Fact | Evidence |
|---|---|
| File nodes exist | Graph contains `819` `File` nodes, all with `filePath`. |
| File classification exists | File nodes include `appLayer`, `functionalArea`, language, extension, document kind, and binary metadata where available. |
| File-to-symbol ownership exists | Graph contains `21334` `DEFINES` relationships from `File` nodes to symbol-like nodes. |
| Symbol containment exists | Graph contains `2784` `CONTAINS` relationships; `143` from file nodes and `2641` from non-file nodes for nesting/ownership. |
| Source-site trace fields exist | `83143` relationships carry `sourceSiteId`, `sourceSiteIds`, and `filePath`; distinct observed source-site ids: `95433`. |
| ResolutionGap trace fields exist | `65652` `ResolutionGap` nodes carry `sourceSiteId` and `filePath`. |
| Unresolved grouping by file is derivable | `576` files have unresolved source-site evidence through `ResolutionGap` file paths. |
| Relationship types are sufficient for first projection | Existing relationship types include `DEFINES`, `CONTAINS`, `CALLS`, `USES`, `IMPORTS`, `ACCESSES`, `MEMBER_OF`, `HAS_PROPERTY`, `HAS_METHOD`, `STEP_IN_PROCESS`, `ENTRY_POINT_OF`, and `HAS_RESOLUTION_GAP`. |
| Command surface owners are identifiable | CLI parent commands exist for `query`, `context`, `impact`, `detect-changes`, `graph-health`, `api`, and `group`; API and graph-health already use child command patterns. |

Plan additions from review:

- Add shared projection service/package as an explicit ownership boundary.
- Add shared target resolver for parent/child command dispatch and ambiguity handling.
- Add projection cache/index invalidation tied to graph freshness.
- Add exact Web/API route naming and generated-contract validation gates.
- Add MCP surface snapshot/tool schema validation gates.

Remaining implementation evidence:

- Existing File/Symbol/SourceSite/ResolutionGap/Flow/API/MCP/test graph facts.
- Current schema facts that support `File -> Symbol`, source-site ownership, symbol nesting, and relationship traceability.
- Missing facts that require implementation.
- Baseline graph inventory summary recorded in benchmark ledger.

## E2 - File Context Contract

Date: 2026-05-30

Status: completed

Contract evidence:

| Contract area | Result |
|---|---|
| Envelope | Added `File Context JSON Contract V0` to the plan. |
| Required top-level fields | `repo`, `repoPath`, `graph`, `target`, `summary`, `symbolTree`, `relationships`, `unresolved`, `linked`, `quality`, and `limits`. |
| Target dispatch fields | `type`, `input`, `normalizedPath`, `dispatchMode`, and `ambiguityCandidates`. |
| Summary fields | Path, language, kind, app layer, functional area, parse status, symbol counts, relationship counts, unresolved count, linked counts, and risk. |
| Relationship shape | `local`, `outboundByFile`, `inboundByFile`, total counts, samples, and trace fields. |
| Unresolved shape | Total, grouped counts, line/column, target text, source symbol, gap kind, classification, actionability, proof kind, source-site id, and source-site status. |
| Linked overlays | Flows, routes, MCP tools, and tests with source/confidence/trace metadata. |
| Quality shape | Parser, resolution confidence, unresolved counts, generated/stale/changed-since-analyze flags. |
| Sample limits | Relationship, unresolved, and linked samples have explicit limits; total counts must not be truncated. |
| Source rules | Contract documents field derivation from `File`, `DEFINES`, `CONTAINS`, symbol relationships, `ResolutionGap`, graph-health, process/route/tool/test facts, and git freshness data. |
| Compatibility | Contract is structured for CLI JSON, API, MCP, and Web; human output may summarize the same shape without changing counts. |

## E3 - Projection Builder

Date: 2026-05-30

Status: completed

Scope:

- Add the first shared backend file-context projection package.
- Do not connect CLI, MCP, API, or Web surfaces yet; those are later phases.
- Do not edit existing graph core storage or analyzer logic.

Impact / blast radius:

| Command | Result |
|---|---|
| `.\anvien\bin\anvien.exe impact "Graph" --repo Anvien --direction upstream` | Returned a broad ambiguous candidate set for `Graph`; treated as HIGH blast-radius signal for graph model changes. No existing `internal/graph` symbol was edited in this slice. |
| `.\anvien\bin\anvien.exe impact "Struct:internal/graph/types.go:Graph" --repo Anvien --direction upstream` | Target UID was not resolved by the current impact command. Scope stayed additive in a new package. |

Implementation evidence:

| File | Evidence |
|---|---|
| `internal/filecontext/context.go` | New shared `filecontext` package with `FileContext`, `FileSummary`, `SymbolTreeNode`, relationship grouping, unresolved grouping, linked/quality/limit structs, `Builder`, `NewBuilder`, and `BuildFileContext`. |
| `internal/filecontext/context.go` | Builder derives file context from existing `graph.Graph` facts: `File` node properties, `DEFINES`, `CONTAINS`, canonical symbol relationships, and `ResolutionGap` nodes with source-site metadata. |
| `internal/filecontext/context.go` | File-level local/outbound/inbound relationship groups preserve total counts while limiting samples. |
| `internal/filecontext/context.go` | Unresolved groups preserve total counts, classification/actionability/kind counts, line/column, target text, proof kind, source-site id, and source-site status. |
| `internal/filecontext/context_test.go` | Fixture tests cover path normalization, file summary, symbol tree nesting, relationship grouping, unresolved grouping, sample limits, missing-file behavior, and deterministic output across relationship insertion order. |

Traceability evidence:

| Projection field | Source fact |
|---|---|
| File summary path/language/layer/area | `File` node properties. |
| Symbol tree | `DEFINES` from File to symbol and `CONTAINS` between symbols. |
| Relationship groups | Canonical relationships such as `CALLS`, with source/target node file paths and source-site fields. |
| Unresolved groups | `ResolutionGap` node properties including `sourceNodeId`, `sourceSiteId`, `sourceSiteStatus`, `targetText`, `classification`, and `actionability`. |

Validation:

| Command | Result |
|---|---|
| `powershell -ExecutionPolicy Bypass -File anvien-launcher\build.ps1` | Pass. Existing Vite dynamic-import and chunk-size warnings only. |
| `go test ./internal/filecontext -count=1` | Pass; `1.052s` after full build gate. |
| `go test ./cmd/... ./internal/... -count=1` | Pass; all cmd/internal packages passed, including new `internal/filecontext`. |
| `.\anvien\bin\anvien.exe analyze --force --name Anvien` | Pass after P1-A code. `files: scanned=821 parsed=586 unsupported=235 failed=0`; `nodes=92352 relationships=126299`. |
| `.\anvien\bin\anvien.exe detect-changes --repo Anvien --scope all` | CRITICAL risk reported for staged P1-A changes: `changed_files=5`, `changed_count=789`, `affected_count=17`; changed layers were `backend`, `backend_test`, and `docs`. |
| `.\anvien\bin\anvien.exe context "BuildFileContext" --repo Anvien` | Found `Method:internal/filecontext/context.go:Builder.BuildFileContext#2`. Incoming edge is `HAS_METHOD` from `Builder`; outgoing edges are same-package projection helpers/structs. |

Failures / handling:

- Initial targeted `go test ./internal/filecontext` was run before the full build gate. The full build gate was then run, and the focused test plus full cmd/internal test batch were rerun after the gate.
- `detect-changes` reported CRITICAL because this slice adds a new backend package with many new symbols and source-site gaps. The affected processes are generated around the new `internal/filecontext` package; the package is not wired into CLI/MCP/API/Web runtime behavior yet. Scope remains additive and tests passed.

Detect changes:

| Command | Result |
|---|---|
| `.\anvien\bin\anvien.exe detect-changes --repo Anvien --scope all` | `risk_level=critical`; `changed_files=5`; `changed_count=789`; `affected_count=17`; affected app layer `backend`; affected functional area `unknown`; no existing runtime command surface was changed in this slice. |

## E4 - File Hotspots And Aggregation

Date: 2026-05-30

Status: completed

Scope:

- Add repo-wide file list/hotspot aggregation to the shared `internal/filecontext` builder.
- Do not add CLI/API/MCP/Web command surfaces yet.
- Do not implement projection cache yet; cache and invalidation remain P1-C.

Impact / blast radius:

| Command | Result |
|---|---|
| `.\anvien\bin\anvien.exe analyze --force --name Anvien` | Pass before P1-B graph-based checks. `files: scanned=821 parsed=586 unsupported=235 failed=0`; `nodes=92352 relationships=126299`. |
| `.\anvien\bin\anvien.exe impact "BuildFileContext" --repo Anvien --direction upstream` | LOW. `impactedCount=0`, no affected modules or processes. |
| `.\anvien\bin\anvien.exe context "Builder" --repo Anvien` | Ambiguous across many builder symbols; exact P1-B edits stayed inside `internal/filecontext`. |

Implementation evidence:

| File | Evidence |
|---|---|
| `internal/filecontext/context.go` | Added `FileListOptions`, `FileList`, and `Builder.BuildFileList`. |
| `internal/filecontext/context.go` | `BuildFileList` builds summaries from all `File` nodes and uses one relationship pass to aggregate local, inbound, and outbound counts. |
| `internal/filecontext/context.go` | Added sort modes: `path`, `unresolved`, `fan-in`, `fan-out`, `symbols`, `flows`, and `tests`. |
| `internal/filecontext/context.go` | Added filters for kind, app layer, functional area, API-related files, unresolved-only, high fan-in, and high fan-out. |
| `internal/filecontext/context.go` | Added limit/offset pagination. |
| `internal/filecontext/context_test.go` | Added tests for sorting, filtering, pagination, high fan-in, and high fan-out behavior. |
| `internal/filecontext/context_test.go` | Added `BenchmarkBuildFileListCurrentScale` using `821` files and `126000` relationships. |

Validation:

| Command | Result |
|---|---|
| `powershell -ExecutionPolicy Bypass -File anvien-launcher\build.ps1` | Pass after P1-B code. Existing Vite dynamic-import and chunk-size warnings only. |
| `go test ./internal/filecontext -count=1` | Pass; `1.578s`. |
| `go test ./internal/filecontext -run '^$' -bench BenchmarkBuildFileListCurrentScale -benchmem -count=3` | Pass. Benchmark runs: `233345033 ns/op`, `106833581 ns/op`, `123577177 ns/op`; `490000 B/op`, `831 allocs/op`. |
| `go test ./cmd/... ./internal/... -count=1` | Pass; all cmd/internal packages passed. |
| `.\anvien\bin\anvien.exe analyze --force --name Anvien` | Pass after P1-B code. `files: scanned=821 parsed=586 unsupported=235 failed=0`; `nodes=92559 relationships=126624`. |
| `.\anvien\bin\anvien.exe detect-changes --repo Anvien --scope all` | MEDIUM risk reported for staged P1-B changes: `changed_files=5`, `changed_count=249`, `affected_count=2`; changed layers were `backend`, `backend_test`, and `docs`. |

Benchmark link:

- See B5 for file list aggregation timing and allocation metrics.

Remaining:

- Projection cache behavior remains pending for P1-C.

Detect changes:

| Command | Result |
|---|---|
| `.\anvien\bin\anvien.exe detect-changes --repo Anvien --scope all` | `risk_level=medium`; `changed_files=5`; `changed_count=249`; `affected_count=2`; affected app layer `backend`; affected functional area `unknown`. |

## E5 - CLI Surface

Date: pending

Status: pending

Expected evidence:

- Help output or command docs for `file-context` and `file-hotspots`.
- Human output examples.
- JSON output examples.
- Missing file and invalid repo behavior.
- CLI tests and smoke commands.

## E6 - Web/API Surface

Date: pending

Status: pending

Expected evidence:

- API impact analysis before route/contract edits.
- Exact route names for file list/hotspots, file context detail, and file relationship expansion if implemented.
- Route implementation evidence.
- Generated Web contract regeneration evidence.
- API tests and shape validation.
- Web consumer integration evidence.
- Generated contract source/output sync validation.

## E7 - Unresolved And Quality Signals

Date: pending

Status: pending

Expected evidence:

- ResolutionGap grouping by file and source symbol.
- Counts by gap kind.
- Classification/actionability/proof/source-site examples.
- Quality fields for parsed/generated/stale/changed/resolution confidence.
- Tests for dynamic, generated, test, and normal source files.

## E8 - Linked Flows, Routes, MCP Tools, And Tests

Date: pending

Status: pending

Expected evidence:

- How links are derived.
- Trace samples for each link type.
- Confidence/source metadata for partial links.
- Tests for files with no links, multiple flows, API handlers, MCP tools, and indirect tests.

## E9 - Web UI File Map And File Detail

Date: pending

Status: pending

Expected evidence:

- File list UI behavior.
- Sort/filter behavior.
- File Detail sections: summary, symbol tree, relationships, unresolved, linked overlays, source-site samples.
- Loading/empty/error/stale states.
- Web unit tests and e2e validation.

## E10 - Parent/Child Command Hierarchy

Date: pending

Status: pending

Expected evidence:

- Inventory of commands whose behavior depends on target type.
- Parent command behavior for smart dispatch and backward compatibility.
- Child command definitions for explicit file, symbol, route, tool, flow, API, and quality workflows where applicable.
- Shared target resolver behavior, ambiguity candidate shape, confidence fields, and exact child-command suggestions.
- Ambiguity handling examples and exact child-command suggestions.
- Help text examples for parent and child commands.
- JSON parity evidence between parent resolved output and explicit child output.
- Commands intentionally left without child commands with reason.

## E11 - Context And Impact Child Commands

Date: pending

Status: pending

Expected evidence:

- `context <target>` remains backward-compatible.
- `context symbol <symbol>` forces symbol context with containing file summary.
- `context file <path>` forces full file context.
- `impact <target>` remains backward-compatible.
- `impact symbol <symbol>` forces symbol impact with file-layer blast radius.
- `impact file <path>` aggregates contained-symbol impact.
- `impact route <route>` and `impact tool <tool>` support target-specific output if implemented.
- Tests for ambiguity suggestions, missing targets, JSON output, and parent/child parity.

## E12 - Query, Change, And Quality Child Commands

Date: pending

Status: pending

Expected evidence:

- `query <text>` remains broad multi-lane search.
- `query files`, `query symbols`, `query flows`, and `query api` behavior where implemented.
- Decision for `detect-changes files/symbols/flows` as child commands or flags.
- Decision for graph quality file/symbol child commands or flags.
- Tests for parent output preservation, child narrowing, sorting/filtering, and JSON output.
- Compatibility/golden tests for parent help, child help, ambiguous target errors, and existing flat command syntax.

## E13 - Existing Command Integration Matrix

Date: pending

Status: pending

Expected evidence:

- Inventory of existing graph-related commands and MCP/API equivalents.
- Classification of each command as `must add file layer`, `may add file layer`, or `no file layer`.
- For each included command, the current output that must be preserved.
- For each included command, the added file-layer section, JSON fields, sample limits, and total counts.
- Evidence that file path inputs do not break symbol-name inputs.
- Commands intentionally left unchanged with reason.

## E14 - Existing Command File-Layer Behavior

Date: pending

Status: pending

Expected evidence:

- `analyze` output keeps current graph inventory and adds file projection build/count/hotspot evidence.
- `query` output keeps current result lanes and adds relevant file hits with matched symbols and relationships.
- `context` output keeps symbol context and adds file summary for symbols; file path input opens full file context.
- `impact` output keeps symbol blast radius and adds impacted file/file-group/flow/test evidence.
- `detect-changes` output keeps changed-symbol detail and groups changed/affected evidence by file.
- Graph quality commands show file-level unresolved, confidence, source-site, and generated-file hotspots.
- API/MCP map commands show handler files, symbol trees, file dependencies, linked tests, and unresolved handler sites.
- Tests or snapshots prove old details are preserved and file-layer sections are additive.

## E15 - Generated Skills And AI Context

Date: pending

Status: pending

Expected evidence:

- Inventory of embedded skill source files that need parent/child file-layer command workflow updates.
- Generated context owners and source-of-truth files updated.
- Generated `.claude/skills/anvien/**`, `AGENTS.md`, and `CLAUDE.md` regenerated through normal analyze/setup paths.
- Workflow examples showing overview -> file -> symbol -> relationship/source-site -> impact/test/flow tracing.
- Source-vs-generated parity validation.
- Tests for skill ids, parent/child command spellings, resource URIs, guidance wording, and absence of placeholder/fallback content.

## E16 - Final Validation And Closure

Date: pending

Status: pending

Expected evidence:

- Full build result.
- Backend/CLI/API/MCP/contracts/AI context tests.
- Generated Web contract check or generator diff validation.
- MCP surface snapshot/tool schema validation if MCP output changes.
- Web unit and e2e tests if applicable.
- File-context smoke outputs for representative files.
- File-hotspots smoke outputs for sort modes.
- Parent/child command smoke checks proving file-layer sections appear without removing current details.
- Projection cache validation smoke for cold build, warm hit, and graph-change invalidation if cache is implemented.
- Command/MCP/resource/prompt/skill guidance validation.
- Final `anvien detect-changes --repo Anvien --scope all`.
- Commit hashes for completed slices.
