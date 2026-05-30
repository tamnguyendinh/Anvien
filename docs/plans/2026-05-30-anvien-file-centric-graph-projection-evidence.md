# Anvien File-Centric Graph Projection Evidence Ledger

Date: 2026-05-30

Status: Draft

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

Status: readiness review recorded; full implementation baseline still pending

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

Plan additions from review:

- Add shared projection service/package as an explicit ownership boundary.
- Add shared target resolver for parent/child command dispatch and ambiguity handling.
- Add projection cache/index invalidation tied to graph freshness.
- Add exact Web/API route naming and generated-contract validation gates.
- Add MCP surface snapshot/tool schema validation gates.

Expected implementation evidence:

- Existing File/Symbol/SourceSite/ResolutionGap/Flow/API/MCP/test graph facts.
- Current schema facts that support `File -> Symbol`, source-site ownership, symbol nesting, and relationship traceability.
- Missing facts that require implementation.
- Baseline graph inventory summary recorded in benchmark ledger.

## E2 - File Context Contract

Date: pending

Status: pending

Expected evidence:

- Final JSON shape for file context.
- Field-by-field source or derivation notes.
- Human CLI output shape.
- Default sample limits and total count behavior.
- Compatibility notes for Web UI and API consumers.

## E3 - Projection Builder

Date: pending

Status: pending

Expected evidence:

- Impact analysis before editing graph/query/model code.
- Files changed for projection structs and builder logic.
- Shared projection service/package name and consumer list across CLI, MCP, API, and Web runtime code.
- Fixture tests proving deterministic output.
- Traceability proof from file-level derived edge back to symbol/source-site facts.
- Guard evidence that command handlers do not reimplement separate projection derivation.

## E4 - File Hotspots And Aggregation

Date: pending

Status: pending

Expected evidence:

- Repo-wide aggregation behavior.
- Sort/filter behavior.
- Pagination or limit behavior.
- Representative hotspot command outputs.
- Performance notes linked to benchmark entries.
- Projection cache behavior for cold build, warm hit, graph-change invalidation, and repo-switch isolation.

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
