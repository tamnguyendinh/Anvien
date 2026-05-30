# Anvien File-Centric Graph Projection Plan

Date: 2026-05-30

Status: Draft

Companion files:

- Evidence ledger: [2026-05-30-anvien-file-centric-graph-projection-evidence.md](2026-05-30-anvien-file-centric-graph-projection-evidence.md)
- Benchmark ledger: [2026-05-30-anvien-file-centric-graph-projection-benchmark.md](2026-05-30-anvien-file-centric-graph-projection-benchmark.md)

## Master Rules

1. Preserve the existing symbol-centric graph as the source of truth.
2. Implement file-centric behavior as a projection/view derived from existing graph facts unless a missing graph fact is proven necessary.
3. Use Anvien for codebase analysis and impact checks during implementation slices.
4. Run `anvien analyze --force` before graph-based work.
5. Run impact analysis before editing graph builders, analyzers, resolvers, API handlers, MCP tools, shared contracts, or exported symbols.
6. Record blast radius before implementation edits; HIGH or CRITICAL impact is a warning to scope carefully, not an automatic stop.
7. Run a full build before testing.
8. If Web UI behavior changes, include an e2e test.
9. Record evidence as each evidenced task is completed.
10. Record benchmark results as each benchmarkable task is completed.
11. Run `anvien detect-changes --repo Anvien --scope all` before each implementation commit.
12. Commit each completed implementation slice after evidence and benchmark ledgers are updated.

## Goal

Add a file-centric graph projection that lets users open a file and immediately see:

- what symbols the file contains;
- how symbols are nested;
- what the file depends on;
- what depends on the file;
- unresolved source sites grouped by file and source symbol;
- linked flows, routes, MCP tools, and tests;
- graph quality signals for that file.

## Problem

The current graph is strong at symbol-level operations:

```text
Symbol -> relationship -> Symbol
SourceSite -> resolved/unresolved target
Flow/API/tool -> linked symbols
```

This is the right source of truth for impact, rename, context, detect-changes, and query. The missing view is file-first inspection. Today, a user investigating one file must mentally combine symbol context, relationships, unresolved gaps, and flow overlays from separate views.

The new capability should answer file-level questions directly:

- Which symbols are declared in this file?
- Which nested methods or members belong under each top-level symbol?
- Which relationships are local to this file?
- Which external files does this file call/reference/import?
- Which files call/reference this file?
- Where are unresolved source sites inside this file?
- Which flows, API routes, MCP tools, and tests involve this file?
- Is this file graph data healthy, stale, generated, or noisy?

## Scope

In scope:

- Derived file summary data.
- File to symbol tree.
- File-level local, inbound, and outbound relationship grouping.
- Unresolved source-site grouping by file.
- Quality signals for parse, resolution, generated status, stale status, and changed-since-analyze status.
- Linked flows, routes, MCP tools, and tests where existing graph data can support the link.
- CLI/API surfaces for file context and file hotspots.
- Web UI File Map and File Detail views.
- A shared backend projection service/package consumed by CLI, API, MCP, and Web runtime code.
- A shared target resolver for file, symbol, route, tool, flow, and API target disambiguation.
- Projection cache/index behavior and invalidation tied to graph freshness.
- Additive file-layer rendering in existing Anvien commands where file context makes the existing command more useful.
- Anvien skills, AGENTS/CLAUDE guidance, setup docs, and command references updated so users and agents know how existing commands expose the file layer.
- Tests, benchmarks, and evidence for the new projection.

## Non-Goals

- Do not replace symbol-centric graph storage.
- Do not make `File -> File` edges the source of truth for impact.
- Do not remove or weaken symbol-level `context`, `impact`, `rename`, or `detect-changes`.
- Do not dump every edge by default in the CLI or Web UI.
- Do not add compatibility aliases or alternate graph schemas without a migration reason.
- Do not add speculative AI summaries as graph facts.

## Core Invariants

1. Symbol relationships remain canonical.
2. File dependencies are derived from symbol/source-site relationships.
3. Every file-level relationship must be traceable back to source symbols and source sites.
4. Default views must summarize first and expand into details on request.
5. Projection output must be deterministic for the same graph input.
6. JSON output must be complete enough for Web UI consumption.
7. Human output must be compact enough for terminal inspection.
8. Generated, test, config, docs, and source files must be classified so file-level signals do not mix unrelated semantics.
9. Existing commands must keep their current useful output and add file-layer sections; file projection must not replace symbol, flow, route, API, or quality details.
10. One shared projection service must own file-level derivation; CLI, MCP, API, and Web surfaces must not reimplement separate derivation logic.
11. One shared target resolver must own parent/child target disambiguation so file-vs-symbol behavior is consistent across command surfaces.
12. Projection caches must invalidate when the underlying graph data, graph hash, graph mtime, repo path, or analyze freshness changes.

## Proposed Projection Shape

```text
File
  -> summary / metadata
  -> top-level symbols
      -> nested symbols
      -> methods
      -> exported members
      -> symbol-level relationship counts
  -> relationships
      -> local relationships
      -> inbound relationships
      -> outbound relationships
  -> unresolved source sites
  -> linked flows / routes / MCP tools / tests
  -> quality signals
```

## Proposed Derived Edges

The projection may expose derived file-level edges, but they must be built from symbol/source-site facts:

```text
(File)-[:CONTAINS]->(Symbol)
(Symbol)-[:CONTAINS]->(Symbol)
(Symbol)-[:DECLARED_IN]->(File)
(SourceSite)-[:IN_FILE]->(File)
(SourceSite)-[:IN_SYMBOL]->(Symbol)
(SourceSite)-[:RESOLVES_TO]->(Symbol)
(SourceSite)-[:UNRESOLVED_TARGET]->(ResolutionGap)
(File)-[:DEPENDS_ON {counts}]->(File)        // derived
(File)<-[:DEPENDED_ON_BY {counts}]-(File)    // derived
(File)-[:PARTICIPATES_IN]->(Flow)
(File)-[:IMPLEMENTS]->(Route|MCPTool)
(File)-[:TESTED_BY]->(File)
```

## Implementation Ownership Boundaries

The implementation should avoid spreading file projection logic across command handlers.

| Layer | Responsibility |
|---|---|
| Shared projection service/package | Build file summaries, symbol trees, relationship groups, unresolved groups, linked overlays, quality signals, hotspot lists, and trace samples from graph facts. |
| Shared target resolver | Resolve parent-command input into explicit target types, return ambiguity candidates, and recommend exact child commands. |
| CLI commands | Validate arguments, call shared resolver/projection services, format human and JSON output, and preserve existing parent command behavior. |
| MCP tools/API routes | Expose the same target semantics and projection contract in structured form. |
| Generated Web contracts | Carry the API response types into `anvien-web` through the existing generator. |
| Web UI | Render typed file projection data; do not derive graph semantics in frontend code. |
| AI context generator | Teach agents the same command hierarchy and file-layer workflow from source-owned embedded skill content. |

The shared service should also expose cache metrics or debug hooks for benchmark evidence. Cache behavior is part of the backend contract because file-list, hotspot, context, and command integration views will all depend on it at current repo scale.

## Command Integration Direction

The file projection is not only a pair of new commands. It is a new display and reasoning layer that existing Anvien commands should use when it improves traceability.

Existing commands must keep their current output and add file-layer sections:

| Command family | Required file-layer behavior |
|---|---|
| `anvien analyze` | After the existing scanned/parsed/node/relationship summary, show file projection build status, file inventory counts, file dependency edge counts, unresolved-file counts, and top file hotspots. |
| `anvien context <symbol-or-file>` | If the input is a file path, show full file context. If the input is a symbol, keep symbol context and add declared file, containing file summary, nearby symbol tree, file inbound/outbound counts, unresolved count, linked flows/tests, and trace samples. |
| `anvien impact <symbol-or-file>` | Keep symbol impact. Add impacted files, impacted file groups, affected flows/tests, changed file risk, and file-level inbound/outbound blast radius. If input is a file path, aggregate impact from contained symbols. |
| `anvien detect-changes` | Keep changed-symbol detail. Group changed symbols by file and show affected files, affected flows/tests, unresolved deltas, and file-level risk summaries. |
| `anvien query` | Keep semantic query results. Add relevant file hits with matched symbols, file summaries, relationship hints, and linked flows/routes/tools where available. |
| `anvien graph-health`, `query-health`, `resolution-inventory`, `source-site-accuracy` | Keep existing graph-quality details. Add top files by unresolved gaps, low confidence, source-site issues, generated-file noise, and file-level quality signals. |
| `anvien api route-map`, `tool-map`, `shape-check`, `api impact` | Keep route/tool/contract details. Add handler files, implementation symbol tree, file dependencies, linked tests, unresolved sites in handler files, and file-level blast radius. |
| group/cross-repo commands | Keep cross-repo behavior. Add file summaries only where repository-local file context is available and clearly labeled by repo. |
| MCP tools/resources/prompts | Keep MCP parity with CLI/API behavior and expose file-layer data in agent-friendly structured output where those tools already answer graph questions. |

Parent and child commands must both work:

```text
anvien context <target> --repo Anvien
anvien context symbol <symbol> --repo Anvien
anvien context file <path> --repo Anvien

anvien impact <target> --repo Anvien
anvien impact symbol <symbol> --repo Anvien
anvien impact file <path> --repo Anvien
anvien impact route <route> --repo Anvien
anvien impact tool <tool> --repo Anvien

anvien query <text> --repo Anvien
anvien query files <text> --repo Anvien
anvien query symbols <text> --repo Anvien
anvien query flows <text> --repo Anvien
anvien query api <text> --repo Anvien
```

Parent commands keep existing usability and smart dispatch. Child commands make the target type explicit for humans, scripts, and agents. Ambiguous parent-command input must report the matching target types and recommend the exact child command to run.

Dedicated file commands can still exist:

```text
anvien file-context <path> --repo Anvien
anvien file-context <path> --repo Anvien --json
anvien file-hotspots --repo Anvien --sort unresolved
anvien file-hotspots --repo Anvien --sort fan-in
anvien file-hotspots --repo Anvien --sort fan-out
```

Web/API should expose equivalent JSON for:

- file list;
- file context detail;
- file hotspot table;
- file relationship expansion.

But these commands are not enough by themselves. They are direct entrypoints for file-first inspection; the core requirement is that existing graph commands can show the same file layer in the context of their current job.

MCP and generated agent guidance must follow the same additive rule: preserve existing tool semantics, add file-layer context where it helps the agent immediately see which file, symbol layer, relationship, flow, route, test, or unresolved site matters.

## Acceptance Criteria

The plan is complete when:

- file context output exists in CLI/API and is backed by deterministic projection data;
- file projection derivation is owned by a shared backend service used by CLI, MCP, API, and Web runtime code;
- parent/child target disambiguation is owned by a shared resolver used by affected command surfaces;
- projection cache and invalidation behavior is tested and benchmarked;
- parent commands and explicit child commands both work for the command families where target type changes behavior;
- existing command outputs keep their current useful details and add file-layer sections where applicable;
- command, MCP/resource, prompt, generated skill, and setup/documentation guidance describes how existing commands expose file-layer evidence;
- file summary includes metadata, symbol counts, relationship counts, unresolved counts, linked flow/test counts, and quality signals;
- symbol tree shows top-level and nested symbols with ranges, kind, export status, signature where available, and relationship counts;
- local, inbound, and outbound relationships are grouped by target/source file and traceable to symbol/source-site samples;
- unresolved source sites include line/col, target text, source symbol when available, classification, and actionability;
- linked flows/routes/MCP tools/tests are shown when graph data supports them;
- Web UI has a file list and file detail view with filters for changed files, unresolved files, API files, generated files, high fan-in, and high fan-out;
- tests cover projection builder, CLI/API output shape, Web UI rendering, and e2e behavior for the file map if UI is changed;
- benchmarks record file inventory counts, projection build time, response size, hotspot counts, and graph inventory counts;
- `anvien detect-changes --repo Anvien --scope all` is recorded before each implementation commit.

## Phase Checklist

- [ ] [P0-A] Baseline current graph schema and file/symbol facts.

  Goal: Prove what file, symbol, source-site, resolution-gap, flow, route, MCP tool, and test facts already exist before adding new code.

  Work Steps:

  1. Run `anvien analyze --force` from the repo root.
  2. Inspect existing graph schema and persisted graph shape for File, Symbol, SourceSite, ResolutionGap, Flow, route, MCP tool, and test-related nodes/edges.
  3. Identify whether `File -> Symbol`, `SourceSite -> File`, and symbol containment already exist or must be derived from current node metadata.
  4. Inspect current CLI/API/MCP graph commands for reusable output and handler patterns.
  5. Record current graph counts and any missing facts in evidence and benchmark files.

  Implementation Gate: No code edits in this task unless missing documentation/test fixtures are discovered; if code edits become necessary, run impact first.

  Acceptance: Evidence names the exact existing graph facts the projection can reuse, lists missing facts, and records baseline graph/file/symbol/source-site counts.

- [ ] [P0-B] Freeze the file context JSON contract.

  Goal: Define the JSON contract before implementation so CLI, API, MCP, and Web UI can share one shape.

  Work Steps:

  1. Draft a `FileContext` response shape with `summary`, `symbolTree`, `relationships`, `unresolved`, `linked`, and `quality`.
  2. Define count fields and sample limits separately so summary counts are stable even when detail lists are truncated.
  3. Define relationship grouping keys: `local`, `outboundByFile`, `inboundByFile`.
  4. Define trace fields for each relationship sample: source file, source symbol, source line/col, relationship kind, target file, target symbol, target line/col when available.
  5. Define unresolved fields: line/col, target text, source symbol, gap kind, classification, actionability, proof kind, and source-site status.
  6. Define quality fields: parse status, generated flag, changed-since-analyze flag, resolution confidence, unresolved call/ref/import counts, and stale status.

  Implementation Gate: Contract changes after implementation starts require evidence explaining why the original contract was insufficient.

  Acceptance: A contract section or schema fixture exists and is referenced by evidence; every field has a source from existing graph data or a documented derivation.

- [ ] [P1-A] Implement the projection model and builder.

  Goal: Add a reusable backend projection builder that produces file context from the existing graph without changing symbol graph ownership.

  Work Steps:

  1. Locate the graph query/model layer that currently powers `context`, `detect-changes`, API maps, and Web graph views.
  2. Run impact on the builder/query functions that will be edited.
  3. Add or identify a single shared projection package/service that CLI, MCP, API, and Web runtime handlers will call.
  4. Add typed structs for file summary, symbol tree nodes, relationship groups, unresolved groups, linked overlays, and quality signals.
  5. Add contract tests that prevent command handlers from bypassing the shared projection service with one-off derivation logic.
  6. Build file lookup by normalized repo-relative path.
  7. Populate declared symbols and symbol nesting using existing symbol range, parent, owner, or containment metadata.
  8. Derive file-level relationship counts from symbol/source-site relationships.
  9. Preserve traceability from every derived file edge back to source symbols and source sites.
  10. Add unit tests with small fixture graphs covering source, test, generated, docs, and unresolved cases.

  Implementation Gate: Do not persist derived file edges as canonical graph facts unless baseline evidence proves current graph cannot support the projection efficiently.

  Acceptance: Unit tests prove the shared builder returns deterministic file context for known fixtures, including counts and relationship samples, and every non-UI surface consumes that shared builder.

- [ ] [P1-B] Add file-level dependency and hotspot aggregation.

  Goal: Add repo-wide aggregation for file list and hotspot views without loading every edge into default output.

  Work Steps:

  1. Implement file summary aggregation for all indexed files.
  2. Compute inbound fan-in, outbound fan-out, unresolved count, symbol count, linked flow count, and linked test count per file.
  3. Add sort modes for unresolved, fan-in, fan-out, symbol count, linked flow count, and recently changed files where data exists.
  4. Add filters for source, test, generated, docs, config, API-related, unresolved-only, high fan-in, and high fan-out.
  5. Add pagination or limit/offset support for large repos.
  6. Add tests that verify sorting, filtering, and limit behavior.

  Implementation Gate: Any new repository-wide scan must be benchmarked; reject designs that require repeated full graph traversal per row in the Web UI.

  Acceptance: A file hotspot/list function returns stable summaries, supports documented sort/filter modes, and has benchmark evidence for current repo scale.

- [ ] [P1-C] Add projection cache, index reuse, and invalidation.

  Goal: Keep file-context, file-hotspot, existing-command, API, MCP, and Web views responsive without duplicating graph traversal.

  Work Steps:

  1. Inspect existing context, query, graph, and Web/API cache/index patterns before adding a new cache.
  2. Define cache keys using repo identity, graph path, graph hash or mtime, graph version, and requested projection options.
  3. Cache reusable indexes such as file lookup, symbol-by-file, relationship-by-file, unresolved-by-file, and linked overlay maps.
  4. Invalidate projection cache when analyze refreshes graph data, graph storage changes, repo path changes, or stale-index status is detected.
  5. Expose cache hit/miss/build timing metrics for benchmark evidence.
  6. Add tests for cold build, warm hit, graph-change invalidation, repo-switch isolation, and stale-index behavior.

  Implementation Gate: Do not add cache behavior that can return file projection data from an older graph after `anvien analyze --force`.

  Acceptance: Projection cache tests prove warm reuse and invalidation, and benchmark entries record cold build, warm query, and invalidation behavior.

- [ ] [P2-A] Add CLI command surfaces.

  Goal: Let users inspect file context and file hotspots from the terminal.

  Work Steps:

  1. Add `anvien file-context <path> --repo <repo>` with compact human output.
  2. Add `anvien file-context <path> --repo <repo> --json` with the full contract.
  3. Add `anvien file-hotspots --repo <repo> --sort <mode>`.
  4. Add help text that explains projection semantics and says symbol graph remains source of truth.
  5. Add tests for argument validation, missing file behavior, path normalization, JSON output, and human output smoke.
  6. Record command output examples in evidence.

  Implementation Gate: Follow existing CLI command patterns; do not invent a separate graph loader if existing commands already have one.

  Acceptance: CLI commands pass tests, return useful errors for missing paths, and produce JSON that validates against the contract.

- [ ] [P2-B] Add Web/API contract surfaces.

  Goal: Expose file list and file context data through existing local Web/API runtime patterns.

  Work Steps:

  1. Identify the current API route registration and generated Web contract flow.
  2. Run API impact before editing route handlers or contracts.
  3. Define exact route names before implementation, including file list/hotspot/detail endpoints and file relationship expansion endpoints if needed.
  4. Add route(s) for file list/hotspots and file context detail.
  5. Add generated TypeScript contract updates through the existing generator, not by hand-editing generated output.
  6. Add API tests for success, missing repo, missing file, filters, sort modes, and JSON shape.
  7. Run shape-check or equivalent API contract validation where available.
  8. Run generated contract check or diff validation so `anvien-web/src/generated` matches the source contract.

  Implementation Gate: Generated Web contracts must be regenerated from source.

  Acceptance: API routes are tested, contracts are regenerated, and Web code can consume typed file context data.

- [ ] [P3-A] Add unresolved source-site grouping and quality signals.

  Goal: Make file-level graph quality visible and actionable.

  Work Steps:

  1. Group ResolutionGap entities by file and nearest source symbol.
  2. Count unresolved calls, refs, imports, member/property accesses, and unknown categories separately where data supports it.
  3. Attach classification, actionability, proof kind, source-site status, and target text to samples.
  4. Compute quality fields for parse status, generated status, stale status, changed-since-analyze status, and resolution confidence.
  5. Add tests for analyzer-gap, external/dynamic, generated-file, and test-file examples.
  6. Ensure default output shows counts first and samples second.

  Implementation Gate: Do not hide unresolved data to make output smaller; use limits with total counts.

  Acceptance: File context shows enough unresolved detail to debug graph quality while preserving total counts and trace samples.

- [ ] [P3-B] Add linked flows, routes, MCP tools, and tests.

  Goal: Connect file context to product workflows and validation coverage, not only code dependencies.

  Work Steps:

  1. Derive linked flows from symbols or source sites already attached to execution flows.
  2. Derive linked API routes and MCP tools from existing route/tool mapping data.
  3. Derive linked tests from test files that reference symbols in the target file or are already connected by graph relationships.
  4. Keep each link traceable to the symbol or source-site relationship that caused it.
  5. Add tests for files with no links, files with multiple flows, API handler files, MCP tool files, and files tested only indirectly.
  6. Add sample limits while preserving total counts.

  Implementation Gate: If flow/route/tool/test links are incomplete, expose confidence/source metadata rather than pretending coverage is complete.

  Acceptance: File context can explain which workflows and tests touch a file, with clear totals and trace samples.

- [ ] [P4-A] Build the Web UI File Map list.

  Goal: Add a scannable file list that helps users find suspicious or important files quickly.

  Work Steps:

  1. Add a `Files` or `File Map` entry using the existing Web navigation style.
  2. Render a table/list with path, layer, functional area, symbol count, inbound count, outbound count, unresolved count, linked flow count, linked test count, and risk/quality signal.
  3. Add filters for changed files, unresolved > 0, API files, generated files, high fan-in, high fan-out, source/test/docs/config.
  4. Add sort controls for unresolved, fan-in, fan-out, symbols, flows, and tests.
  5. Add loading, empty, error, and stale-index states.
  6. Add component tests for rendering and filter behavior.

  Implementation Gate: Run full build before Web tests; include e2e coverage because this changes Web UI behavior.

  Acceptance: Users can open the File Map, sort/filter files, and identify unresolved/high-dependency files without reading raw graph JSON.

- [ ] [P4-B] Build the Web UI File Detail view.

  Goal: Let users click a file and inspect summary, symbol tree, relationships, unresolved gaps, linked overlays, and raw source-site samples.

  Work Steps:

  1. Add file detail routing or selected-row panel according to existing Web UI patterns.
  2. Render summary and quality signals at the top.
  3. Render symbol tree with expandable top-level and nested symbols.
  4. Render relationship groups as local, inbound, and outbound sections grouped by file.
  5. Render unresolved source-site samples with line/col and actionability.
  6. Render linked flows/routes/MCP tools/tests with counts and trace samples.
  7. Add tests for long paths, empty sections, many relationships, and unresolved samples.
  8. Add e2e coverage that opens a file and verifies the major sections.

  Implementation Gate: Avoid nested cards and dense unbounded dumps; default view must be summary-first with expandable detail.

  Acceptance: File Detail is readable at current repo scale and exposes all contract sections without layout overlap.

- [ ] [P5-A] Design the parent/child command hierarchy for target-aware file-layer usage.

  Goal: Make command usage explicit and scriptable while keeping existing parent commands working.

  Work Steps:

  1. Inventory existing commands whose behavior depends on target type: `context`, `impact`, `query`, `detect-changes`, graph quality commands, API route/tool commands, and group commands.
  2. Define parent-command behavior for each family. Parent commands must preserve current usage and smart dispatch.
  3. Define child commands for each family where target type changes behavior. Required candidates include `context symbol`, `context file`, `impact symbol`, `impact file`, `impact route`, `impact tool`, `query files`, `query symbols`, `query flows`, and `query api`.
  4. Implement or specify a shared target resolver that returns resolved target type, ambiguity candidates, confidence, and exact child-command suggestions.
  5. Define which commands should not have child commands and why. For example, `analyze` should remain repo/path oriented unless a real target-specific analyze mode exists, and `rename` should remain symbol-first unless a separate file rename workflow is designed.
  6. Define ambiguity handling. If a target could be a file, symbol, route, or tool, the parent command must report matches and recommend exact child commands.
  7. Define help text layout so parent help shows common usage and child help shows target-specific output and JSON contract.
  8. Define JSON contract parity between parent smart-dispatch output and explicit child output.
  9. Add compatibility/golden tests for parent help, child help, ambiguous target errors, and existing flat command syntax.

  Implementation Gate: Do not remove existing parent-command syntax. Existing scripts using parent commands must keep working unless a breaking change is separately approved.

  Acceptance: The hierarchy document lists parent commands, child commands, unsupported child commands, shared resolver behavior, ambiguity rules, help behavior, and JSON parity requirements.

- [ ] [P5-B] Implement explicit child commands for context and impact.

  Goal: Make the two highest-value target-aware commands explicit for file and symbol workflows.

  Work Steps:

  1. Add or update `context symbol <symbol>` to force symbol context and include containing file summary.
  2. Add or update `context file <path>` to force full file context and avoid ambiguity with symbol names.
  3. Keep `context <target>` as smart dispatch and add ambiguity suggestions.
  4. Add or update `impact symbol <symbol>` to force symbol blast radius plus file-layer affected evidence.
  5. Add or update `impact file <path>` to aggregate impact from contained symbols and show impacted files/flows/tests.
  6. Add or update `impact route <route>` and `impact tool <tool>` if existing API/tool impact data supports stable target-specific behavior.
  7. Add tests for parent compatibility, child command behavior, ambiguous target suggestions, missing target errors, and JSON output.

  Implementation Gate: Child command output must be target-specific, not a thin alias that hides ambiguity without adding clarity.

  Acceptance: Users and agents can explicitly ask for symbol or file context/impact and get deterministic output with file-layer evidence.

- [ ] [P5-C] Implement explicit child commands for query and change/quality workflows.

  Goal: Let users narrow discovery and diagnostics by target layer while preserving broad parent commands.

  Work Steps:

  1. Keep `query <text>` as multi-lane search.
  2. Add or update `query files <text>` for file-first results with matched symbols, file summaries, linked flows, and relationship hints.
  3. Add or update `query symbols <text>` for symbol-first results with containing file summaries.
  4. Add or update `query flows <text>` and `query api <text>` where existing lanes can support stable filtered output.
  5. Define whether `detect-changes files`, `detect-changes symbols`, and `detect-changes flows` are commands or flags based on existing command architecture; implement the chosen form with evidence.
  6. Define whether graph quality child commands such as `graph-health files` or `resolution-inventory files` are commands or flags; implement the chosen form with evidence.
  7. Add tests for parent output preservation, child narrowing, sorting, filtering, and JSON output.

  Implementation Gate: Do not split commands so far that users must know internal graph taxonomy before getting useful output; parent commands remain the default on-ramp.

  Acceptance: Users can start broad with parent commands or force a file/symbol/flow/API view with child commands, and both paths expose the file layer.

- [ ] [P5-D] Align MCP/API/generated contracts with parent/child command semantics.

  Goal: Keep agent/API contracts consistent with the CLI hierarchy.

  Work Steps:

  1. Map each parent and child CLI command to existing or new API/MCP equivalents.
  2. Use the same target type names across CLI, API, MCP tools, resources, and generated TypeScript contracts.
  3. Add agent-friendly structured fields for target type, dispatch mode, ambiguity candidates, selected file, selected symbol, and file-layer sections.
  4. Regenerate generated Web contracts from source where API shape changes.
  5. Update MCP tool surface tests/snapshots when tool schemas or result payloads change.
  6. Add parity tests so parent smart dispatch and explicit child commands produce compatible JSON for the same resolved target.
  7. Add shape checks or tool-map/route-map validation where command/API/MCP contracts are affected.
  8. Validate generated Web contract source and output are in sync after regeneration.

  Implementation Gate: Do not create MCP-only or API-only target type names that do not exist in CLI behavior.

  Acceptance: CLI, API, MCP, and generated contracts share the same parent/child target semantics and ambiguity behavior.

- [ ] [P6-A] Define the existing-command file-layer integration matrix.

  Goal: Decide exactly how every existing graph-related Anvien command should keep its current output and add the file layer.

  Work Steps:

  1. Inventory all existing graph-related commands and MCP/API equivalents: `analyze`, `status`, `list`, `query`, `context`, `impact`, `detect-changes`, `rename`, `augment`, `cypher`, `graph-health`, `query-health`, `resolution-inventory`, `source-site-accuracy`, `api route-map`, `api tool-map`, `api shape-check`, `api impact`, and group commands.
  2. Classify each command as `must add file layer`, `may add file layer`, or `no file layer` with a reason.
  3. For each `must add file layer` command, define the default human output section, JSON fields, sample limits, and total counts.
  4. Define how file path inputs are detected without breaking symbol-name inputs.
  5. Define how command output progresses from overview to file to symbol to source-site detail.
  6. Define common wording for derived file edges so users know file relationships are projections from symbol/source-site facts.
  7. Record commands intentionally left unchanged and why.

  Implementation Gate: No command may lose existing symbol, flow, API, route, MCP tool, or graph-quality detail as part of this integration.

  Acceptance: The matrix states what every existing command and child command will display, what stays unchanged, and how the new file layer connects overview, file, symbol, relationship, and source-site evidence.

- [ ] [P6-B] Add file-layer output to analyze, query, context, impact, and detect-changes.

  Goal: Make the main daily graph commands show file context directly inside their existing workflows.

  Work Steps:

  1. Update `analyze` so file projection build and file hotspot counts appear after the current graph inventory summary.
  2. Update `query` so relevant files appear alongside symbol/process/docs results, with matched symbols and relationship hints.
  3. Update `context` so a file path opens full file context and a symbol context includes its containing file summary, file relationships, unresolved count, and linked flows/tests.
  4. Update `impact` so a file path aggregates contained-symbol impact and a symbol impact includes file-level blast radius.
  5. Update `detect-changes` so changed symbols are grouped by changed file, with affected files, affected flows/tests, unresolved deltas, and file-level risk.
  6. Add JSON output fields for each command without requiring consumers to parse human text.
  7. Add tests proving old command details still exist and new file-layer sections appear where expected.

  Implementation Gate: Do not replace current command output with file output; file sections are additive and must preserve current symbols, process, and impact details.

  Acceptance: A user can run the same existing command and immediately see the relevant file layer without losing the previous symbol-centric details.

- [ ] [P6-C] Add file-layer output to graph quality and API/MCP command families.

  Goal: Make graph quality, route, and tool diagnostics show where relationships are connected or broken at file level.

  Work Steps:

  1. Update `graph-health` summaries/reports so top file hotspots show unresolved gaps, low confidence, generated-file noise, and source-site quality issues.
  2. Update `resolution-inventory` so unresolved gaps can be grouped by file and nearest source symbol.
  3. Update `source-site-accuracy` so source-site failures include file-level grouping and trace samples.
  4. Update `query-health` where relevant so retrieval gaps can identify whether missing results cluster by file, symbol layer, or app area.
  5. Update `api route-map`, `api tool-map`, `api shape-check`, and `api impact` so handler files, symbol tree, file dependencies, linked tests, and unresolved handler-file sites are visible.
  6. Update MCP tool/resource output for equivalent command families so agents receive the same file-layer facts in structured form.
  7. Add tests or snapshots that prove command/API/MCP parity for file-layer fields.

  Implementation Gate: File-layer diagnostics must preserve total counts and samples; do not hide graph quality problems to shorten output.

  Acceptance: Quality/API/MCP commands show not only what symbol/route/tool is involved, but which files connect, which files have broken links, and where source-site evidence lives.

- [ ] [P6-D] Update Anvien skills and generated agent context for parent/child file-layer command usage.

  Goal: Teach agents to use parent and child Anvien commands with the new file layer, not to treat file-context as a disconnected feature.

  Work Steps:

  1. Inventory embedded skill source files under the AI context generator and generated `.claude/skills/anvien/**` output that teach graph exploration, impact analysis, debugging, graph quality, API surface inspection, CLI usage, and guide workflows.
  2. Update embedded source skill Markdown so agents know which parent or child command to use when the user starts from a file path, a symbol, a broken relationship, a route/tool, a graph quality problem, or a change set.
  3. Add workflow examples showing overview-to-detail tracing: query/analyze hotspot -> file summary -> symbol tree -> relationship/source-site evidence -> impact/test/flow confirmation.
  4. Update generated root `AGENTS.md` / `CLAUDE.md` guidance so file paths trigger file-layer command usage while symbol questions still use symbol context.
  5. Regenerate generated skills and root context through normal analyze/setup paths.
  6. Add tests for generated skill ids, command spellings, resource URIs, guidance wording, and absence of placeholder/fallback content.
  7. Validate source-vs-generated parity so embedded skill source and generated skill output match.

  Implementation Gate: Never patch generated `AGENTS.md`, `CLAUDE.md`, or `.claude/skills/anvien/**` as the source of truth; update generator-owned source and regenerate.

  Acceptance: Skills teach agents to use parent commands for broad discovery and child commands for explicit file/symbol/route/tool workflows, so any trace can move from overview to file to symbol to source-site evidence.

- [ ] [P7-A] Add validation, evidence, and benchmark coverage.

  Goal: Prove the projection is correct, performant enough, and does not degrade symbol graph behavior.

  Work Steps:

  1. Run full build.
  2. Run backend tests for projection, CLI, API, MCP, contracts, AI context, and graph quality grouping.
  3. Run generated Web contract check or generator diff validation.
  4. Run MCP surface snapshot/tool schema tests if MCP tool output changes.
  5. Run Web unit tests and e2e tests if Web UI changed.
  6. Run file-context CLI smoke on representative files: source, test, generated, docs/config, API/MCP file, unresolved-heavy file.
  7. Run file-hotspots smoke for unresolved, fan-in, fan-out, and symbol count sort modes.
  8. Run parent/child command smoke checks proving file-layer additions appear in `analyze`, `query`, `context`, `impact`, `detect-changes`, graph quality commands, and API/MCP map commands without removing existing details.
  9. Run cache validation smoke for cold build, warm hit, and graph-change invalidation if cache was implemented.
  10. Record graph counts, file inventory counts, projection timing, response sizes, hotspot counts, command coverage counts, and Web validation evidence.
  11. Run `anvien detect-changes --repo Anvien --scope all`.

  Implementation Gate: Benchmarkable outputs go in the benchmark ledger; command validation and pass/fail evidence go in the evidence ledger.

  Acceptance: Evidence and benchmark ledgers are complete enough to explain what changed, how it was validated, and what scale/performance numbers were observed.

- [ ] [P7-B] Close the implementation with docs and commits.

  Goal: Finish the feature without leaving stale contracts, docs, or uncommitted implementation slices.

  Work Steps:

  1. Update user-facing docs for the new commands and Web UI view.
  2. Update MCP/resource/setup docs if MCP surfaces are added.
  3. Update any generated contracts through the generator.
  4. Re-run final old/stale command names or contract path checks if any new docs/config were generated.
  5. Run final full build and required tests.
  6. Run final `anvien detect-changes --repo Anvien --scope all`.
  7. Commit each completed implementation slice with evidence and benchmark updates.

  Implementation Gate: Do not mark the plan complete if generated output is hand-edited or evidence/benchmark ledgers are stale.

  Acceptance: Working tree is clean after final commit, all relevant checks are recorded, and the feature is discoverable from CLI/Web docs.

## Risk Notes

- Large repositories may make naive file aggregation expensive.
- File-level edges can mislead users if they are not clearly marked as derived.
- Unresolved source-site data can be noisy in test and dynamic-language files.
- Generated files can dominate counts unless classified and filterable.
- Web UI can become too dense if every relationship is shown by default.
- Flow/test links may be partial; confidence/source metadata must be visible.
- Existing command output can regress if file sections replace rather than extend current symbol/flow/API/quality details.
- CLI, MCP, API, Web, and generated skills can drift if file projection and target resolution are not shared backend contracts.
- Projection cache can serve stale graph data if invalidation is not tied to analyze output and graph storage metadata.

## Definition Of Done

- The symbol graph remains canonical.
- File projection is available through CLI/API and Web UI.
- File projection and target resolution are implemented as shared backend contracts, not separate per-command logic.
- Existing graph commands preserve current details and add file-layer evidence where it helps trace from overview to file to symbol to source site.
- File projection guidance is available through the appropriate Anvien command, MCP/resource/prompt, generated context, and skill surfaces.
- File context includes summary, symbol tree, grouped relationships, unresolved source sites, linked overlays, and quality signals.
- Tests, evidence, and benchmarks cover the new projection.
- Full build and relevant tests pass.
- Detect-changes is recorded before implementation commits.
