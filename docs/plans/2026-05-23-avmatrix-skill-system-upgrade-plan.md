# AVmatrix Skill System Upgrade Plan

Date: 2026-05-23

Status: Active

Companion files:

- Benchmark ledger: [2026-05-23-avmatrix-skill-system-upgrade-benchmark.md](2026-05-23-avmatrix-skill-system-upgrade-benchmark.md)
- Evidence ledger: [2026-05-23-avmatrix-skill-system-upgrade-evidence.md](2026-05-23-avmatrix-skill-system-upgrade-evidence.md)

## Master rules

1. Use AVmatrix for codebase analysis and impact checks while working on implementation slices in this plan.
2. As each task is completed, update the corresponding checklist item immediately.
3. Run the full build gate before testing; include focused backend/CLI/setup/package validation for generated skill behavior, and include Web unit/e2e/browser screenshot validation for the graph labeling phase.
4. Record benchmark results as each benchmarkable task is completed. Benchmarkable means measured product/runtime performance, capacity, package/startup size, graph/DB throughput, query hit rate, command output inventory, graph inventory counts, source-site inventory counts, generated skill inventory counts, setup/package file inventories, or resolved-edge accuracy; build/test/e2e timings are validation evidence unless the slice changes those systems.
5. Record evidence as each evidenced task is completed.
6. For doc-only commits, do not use AVmatrix.
7. After each completed implementation slice, commit the work, then continue until the full plan is complete.

## Rules of plan

1. Follow active workspace and repository instructions, including `AGENTS.md`; this plan records product work and validation, it does not replace repository rules.
2. Use AVmatrix according to active repository instructions for implementation slices; do not use AVmatrix for doc-only commits.
3. Never use `--skip-agents-md`. The generated `AGENTS.md`, `CLAUDE.md`, and `.claude/skills/avmatrix/**` files are part of the AVmatrix AI context output and must be validated through the normal generation path.
4. Do not edit generated `.claude/skills/avmatrix/**`, generated root `AGENTS.md`, or generated root `CLAUDE.md` as source files. Update the generator source, embedded skill Markdown, tests, and docs that produce those outputs.
5. Keep generated AVmatrix guidance repo-agnostic. The project-name/statistics line inside the generated managed block may be auto-filled for the current repository, but command descriptions and skill guidance must work for any indexed repository.
6. Treat MCP tools, CLI commands, resources, Web/API views, and generated skills as AVmatrix command surfaces. Do not narrow the guidance to only `analyze`, `query`, and `impact` when more precise AVmatrix operations exist.
7. Run a full build before testing. For this plan the full build gate is `powershell -ExecutionPolicy Bypass -File avmatrix-launcher\build.ps1`.
8. If implementation touches function, class, or method symbols, run impact first and record the blast radius. HIGH or CRITICAL risk is a warning to report and account for, not a reason to abandon the required work.
9. Run `detect-changes` before committing an implementation slice and record the expected changed scope.
10. Command inventory must be taken from the current source tree or the freshly built local binary. A stale `avmatrix` found in `PATH` may be recorded as evidence, but it must not define final skill content.

## Problem

The generated `AGENTS.md` and `CLAUDE.md` files point agents to `.claude/skills/avmatrix/**`, but the skill set and skill content have not kept pace with the current AVmatrix command surface.

The existing generated Skills table lists six skills:

- `avmatrix-exploring`
- `avmatrix-impact-analysis`
- `avmatrix-debugging`
- `avmatrix-refactoring`
- `avmatrix-guide`
- `avmatrix-cli`

Those skills are generated from source files, not manually authored under `.claude/skills/avmatrix/**`. Current source inspection shows the source of truth is in `internal/aicontext/aicontext.go` and `internal/aicontext/skills/*.md`, while `.claude/skills/avmatrix/**` is generated output. Updating generated files directly would be the wrong fix.

The skill content also under-describes current AVmatrix capability. It does not clearly route agents to newer or specialized surfaces such as API route/tool analysis, query health, resolution/source-site accuracy, cross-repo groups, runtime/packaging/setup flows, generated AI context maintenance, or the unified command surface that includes MCP tools and CLI equivalents.

There is also a source-distribution risk. Repository-local generated skills are produced from embedded Markdown under `internal/aicontext/skills/*.md`, but editor setup currently installs skills from a package-root `skills/` directory when one exists. If those two sources are not reconciled, users can receive different skills depending on whether they run `analyze` in a repository or `setup` from a packaged install.

There is also a build artifact drift problem. The repository can contain `avmatrix\bin\avmatrix.exe` and `avmatrix-launcher\server-bundle\avmatrix.exe` at the same time, and the full build path currently builds the launcher bundle binary independently from the package CLI binary. That allows validation, agents, launcher runtime, and users to execute different AVmatrix versions from the same checkout. The product must have one production/canonical executable path: `avmatrix\bin\avmatrix.exe`. Build, launcher, tests, docs, and command examples must use that path as the single source of truth. If the launcher needs to start the backend, it must use the canonical executable path rather than a second private copy in `server-bundle`.

There is also a graph-health command-surface gap. Graph-health computation and metadata exist in the codebase, and Web/API surfaces already use graph-health data for topology, confidence, diagnostics, explain/report behavior, and filters. However, `avmatrix --help` does not expose a first-class `graph-health` CLI command. Agents and users therefore cannot audit topology health directly from terminal workflows when they need to investigate orphan nodes, unknown connectivity, detached components, no-incoming/no-outgoing nodes, or dead-code candidates. This is a missing product surface, not missing graph-health engine capability.

There is also a broader CLI parity gap. Current source/help comparison shows several implemented MCP/API/Web capabilities are not exposed as clear user-facing CLI commands. `rename` exists as an MCP tool but does not have a CLI command. API-surface tools such as `route_map`, `tool_map`, `shape_check`, and `api_impact` exist as MCP tools but not as CLI commands. HTTP/Web graph support exposes graph explain/report, grep, process, cluster, and search endpoints, while the CLI only exposes a subset through `cypher`, `query`, or MCP resources. This plan must not silently normalize missing terminal surfaces when the product goal is a broad AVmatrix command system that agents and users can operate from CLI, MCP, resources, Web/API, and generated skills.

There is also an MCP prompt-template accuracy gap. README lists MCP prompts such as `generate_map`, and source inspection shows `generate_map` is a real MCP prompt in `internal/mcp/prompts.go`, not a CLI command. However, its current body is too loose for reliable agent execution in large repositories: it uses `{name}` as a placeholder when no repo is supplied, does not instruct the agent how to resolve the repo from `avmatrix://repos`, does not make freshness/staleness handling explicit, does not define how to choose "top 5 most important processes", and can encourage architecture diagrams that are inferred beyond graph evidence. MCP prompts are part of AVmatrix's command surface and must be accurate, executable templates rather than demo text.

There is a query reliability bug. A broad intent query can return plausible but wrong code regions. During plan review, broad queries about AI context skill generation returned unrelated launcher, resolution-gap, and frontend/backend flows instead of the expected `internal/aicontext`, analyze post-run, setup, and package-skill surfaces. That behavior is dangerous for agent work because a query that cannot identify the right region can send the agent to edit or reason about the wrong code. This is not a minor documentation issue or normal behavior to accept. `query` is a core AVmatrix feature and must be able to locate the correct work area for broad intent. Until the bug is fixed, broad `query` output must be treated as candidate retrieval and verified by symbol-level `context`, exact file/symbol inspection, or `query-health` evidence.

There is also a Web graph orientation problem shown by `reports/problem/screenshot_1779517751.png`. The graph can contain visible rings and node islands, but the canvas does not name the ring or island directly. Users can see colored clusters, but they cannot immediately tell whether a macro ring is Backend, Frontend, API, Docs, Config, Shared, Test, Unknown, or another top-level group, and they cannot tell whether an island is Function, Method, File, Route, ResolutionGap, External Reference, or another node/filter group without looking away from the graph. Color and side-panel filters are not enough; the graph itself must communicate what each ring and island represents.

## Scope

Implementation may touch:

- `internal/aicontext/aicontext.go`;
- `internal/aicontext/skills/*.md`;
- tests under `internal/aicontext`;
- CLI/setup/package tests under `internal/cli` if they assert skill counts, skill paths, generated output, or packaging behavior;
- graph-health CLI command implementation and tests under `internal/cli`, shared graph-health report/explain logic under `internal/graphhealth` or another source-verified shared package, and API reuse points under `internal/httpapi` if needed to avoid duplicate graph-health semantics;
- CLI parity command implementation and tests for implemented MCP/API/Web surfaces that are accepted as user-facing terminal commands, including `rename`, API-surface tools, graph/process/cluster panels, grep/search, or equivalent grouped command families if source inspection confirms they should be exposed;
- MCP prompt implementation and tests under `internal/mcp`, especially `internal/mcp/prompts.go` and prompt surface snapshots if they describe architecture-map or impact workflows;
- query implementation, query user-facing command surfaces, query output formatting, and tests under `internal/mcp`, `internal/cli`, query-health suites, and any query ranking/scoring helpers if source inspection proves broad-intent query misses expected owners;
- Web graph layout, graph adapter, Sigma/canvas overlay, graph canvas components, graph filter/legend integration, and Web tests under `avmatrix-web` for ring/island labeling and visual orientation;
- MCP setup/resource guidance under `internal/mcp` if it contains tool, resource, setup, or command-surface reference text;
- README and user-facing docs that explain generated AVmatrix skills, AI context setup, or AVmatrix command surfaces;
- packaging/setup code only if source inspection proves the expanded embedded skill set is not installed or packaged correctly.
- launcher build/runtime files if source inspection confirms the plan can produce or run more than one AVmatrix CLI executable, including `avmatrix-launcher/build.ps1`, `avmatrix-launcher/src/main.go`, `avmatrix-launcher/server-wrapper/main.go`, launcher tests, and docs that name runtime binary paths.

Out of scope unless source inspection proves it is required:

- changing the behavior of AVmatrix graph analysis commands;
- changing Web UI graph rendering outside the graph labeling and visual orientation layer required by this plan;
- changing MCP tool input schemas unless a backward-compatible query usability extension is required by this plan;
- editing generated `.claude/skills/avmatrix/**`, `AGENTS.md`, or `CLAUDE.md` directly as the source of truth;
- changing historical evidence ledgers only to make old records match the new skill set.

## Design Decisions

Base skill source files live under `internal/aicontext/skills/*.md` and are embedded by `internal/aicontext/aicontext.go`. The implementation must update those source files and the `baseSkills` registry rather than patching generated output.

The generated root Skills table in `AGENTS.md` and `CLAUDE.md` must be generated from the same intended skill set. If the table remains hard-coded, tests must protect it against drifting away from `baseSkills`.

The upgraded skill set should keep the existing six skills and add five focused skills:

| Skill | Purpose |
|---|---|
| `avmatrix-exploring` | Architecture exploration, execution-flow discovery, process/context/resource usage. |
| `avmatrix-impact-analysis` | Blast-radius work, impact interpretation, changed-scope checks, HIGH/CRITICAL warning handling. |
| `avmatrix-debugging` | Bug tracing with graph facts, runtime evidence, diagnostics, source-site and resolution health where relevant. |
| `avmatrix-refactoring` | Rename/extract/split/refactor workflows using graph guidance, impact, query/context, and detect-changes. |
| `avmatrix-guide` | Unified AVmatrix command surface and schema/resource reference across MCP, CLI, resources, Web/API, and generated skills. |
| `avmatrix-cli` | Complete CLI command guide, including analyze/status/query/context/impact/detect-changes/cypher/rename/augment/group/setup/serve/mcp/package/wiki/hook/version, graph-health, and any current accuracy commands confirmed by source. |
| `avmatrix-graph-quality` | Graph-health topology audit, query health, source-site inventory, resolution inventory, edge accuracy, ResolutionGap/UnresolvedSymbol review, and benchmark comparison. |
| `avmatrix-api-surface` | API routes, MCP tools, contract shape checks, API impact, generated Web contracts, handlers, and consumers. |
| `avmatrix-cross-repo` | Group repositories, cross-repo query/contracts/status/sync, and multi-repo analysis guidance. |
| `avmatrix-runtime-packaging` | `serve`, `mcp`, `setup`, launcher, packaged runtime, package preparation, runtime cleanup, and startup validation. |
| `avmatrix-ai-context` | Generated `AGENTS.md`, `CLAUDE.md`, `.claude/skills/avmatrix/**`, source-vs-generated rules, regeneration, and validation. |

The exact command list inside each skill must be based on current source/help output during implementation. If a command name in this plan is not available in the current codebase, the implementation must not document it as available; record the mismatch in evidence and update the skill wording accordingly.

Command names must match the surface that implements them. CLI commands use hyphenated names such as `query-health`, `resolution-inventory`, `source-site-accuracy`, and the planned `graph-health` command. MCP tools use underscore names such as `route_map`, `tool_map`, `shape_check`, and `api_impact`. A skill may mention both surfaces, but it must not invent a CLI command just because an MCP tool exists.

The graph-health CLI surface must be a first-class graph-quality command, not a Web-only feature and not a query-health substitute. The implementation should expose topology health through subcommands such as `avmatrix graph-health summary --repo <repo> --json`, `avmatrix graph-health report --repo <repo> --limit <n> --json`, `avmatrix graph-health explain "<node-id-or-name>" --repo <repo> --json`, and `avmatrix graph-health components --repo <repo> --json`, or an equivalent source-verified design. The command must start from fresh analyze output according to repository rules, read the same graph-health metadata/compute behavior used by Web/API, and return table output plus machine-readable JSON. It must not fork a second graph-health definition that can drift from Web/API semantics.

CLI parity must be decided deliberately, not by accident. Phase 1.6 must inventory MCP tools, MCP resources, HTTP/Web endpoints, and generated guidance against the CLI command tree, classify each gap as must-add, intentionally MCP/API/Web-only, hidden lifecycle-only, or already covered by an existing CLI command. Accepted must-add surfaces should prefer coherent command families over scattered top-level commands where that improves usability, for example `avmatrix api route-map`, `avmatrix api tool-map`, `avmatrix api shape-check`, `avmatrix api impact`, `avmatrix graph-health ...`, or `avmatrix graph process/cluster/search` if source inspection supports those shapes. Hidden lifecycle commands such as package or editor hook helpers must remain documented as internal only if they are not meant for normal users.

MCP prompts must be treated as executable agent templates, not static README examples. A prompt such as `generate_map` must resolve the target repo deterministically, treat `avmatrix://repos` only as a discovery/selection resource, never continue with `{name}` as a real URI, require fresh graph evidence before graph-based mapping, define how process candidates are selected, and require diagrams and architecture claims to be backed by resources or command output the agent actually read.

The package/editor setup skill source must be reconciled with the embedded AI-context skill source. The final implementation should make it hard for package-root `skills/`, embedded `internal/aicontext/skills/*.md`, and generated `.claude/skills/avmatrix/**` to drift away from each other.

The production executable path is `avmatrix\bin\avmatrix.exe`. There must not be a separate internal AVmatrix CLI executable with independent build output. `avmatrix-launcher\server-bundle` may contain launcher support binaries and native runtime files, but it must not define a second AVmatrix CLI authority. The full build must produce the canonical executable at `avmatrix\bin\avmatrix.exe`; launcher/runtime code must resolve and run that canonical executable; validation must prove `query --lanes`, `query --explain`, `query-health`, `analyze`, and runtime `serve` are available from that exact file after build.

Every base skill source file must include valid skill frontmatter with `name` and `description`. The fallback skill generator remains useful as defensive code, but final tests should prove no final base skill depends on fallback content because of a missing or empty embedded source file.

MCP setup/resource guidance is part of the same user-facing command guide. If `avmatrix://setup` or related resource output lists tools/resources/commands, it must be checked and updated with the same final command taxonomy instead of leaving a second stale guide in the codebase.

The upgraded skills must explain query reliability honestly. `query` is useful for finding candidate flows from concepts, but broad-intent query results are not proof that a region is the correct owner. When a symbol, file, command, or generated artifact is known, the skill guidance must prefer `context` and exact source inspection. When broad-intent query misses the expected owner, that is graph-quality evidence to record through `query-health` or a benchmark case, not a result to silently accept.

The current `query` command must be treated as a concept-to-code and concept-to-flow discovery command. The CLI `avmatrix query` surface is a wrapper into the MCP query implementation, which combines graph data, process matches, definition matches, semantic fields, resolution-gap summaries, and warnings. Its product purpose is not the same as `context`: `query` should help an agent find likely work areas from broad intent, while `context` should inspect a known symbol or exact owner after the candidate area is identified.

`query` is a broad command family, not one narrow lookup. The implementation must define a Query Capability Taxonomy under the top-level `query` behavior instead of treating one bug case as the full product definition. The taxonomy should include at least:

| Query capability | Purpose |
|---|---|
| owner discovery | Find the file, symbol, command, resource, generated artifact, or package that owns a problem. |
| concept discovery | Find likely code areas from broad natural-language intent. |
| execution-flow discovery | Find processes, flows, and process steps related to the intent. |
| API surface discovery | Find route/tool handlers, contracts, generated API types, and consumers. |
| graph-quality discovery | Find query-health, source-site, resolution, ResolutionGap, graph-health, and accuracy surfaces. |
| docs/setup/AI-context discovery | Find generated guidance, skill source, setup, package, and AI context generation surfaces. |
| command-surface discovery | Find CLI, MCP, resource, Web/API, package, and runtime command owners. |
| cross-repo discovery | Find group/cross-repo query, contracts, sync, status, and multi-repo surfaces when indexed data supports them. |

These capabilities are not separate product commands unless the codebase already exposes them that way. They are retrieval lanes inside the umbrella `query` behavior. A query result may be produced by one or more lanes, and the output should expose the lane or match reason so users and agents can understand why it ranked.

Query capability work must produce usable product surfaces, not hidden internal scoring. A capability is not complete because a struct field or scoring branch exists in code. It is complete only when a user or agent can discover and use it through AVmatrix command surfaces. The CLI must expose clear query subcommands or flags for lane discovery and explainable query output, such as `avmatrix query lanes`, `avmatrix query explain "<intent>" --repo <repo> --json`, or an equivalent source-verified design. The existing `avmatrix query "<intent>" --repo <repo>` behavior must remain available for normal broad discovery. MCP `query` output must expose the same lane/rank/match evidence in a machine-readable way for agents. If an existing Web/API query/search surface consumes query results, it must display or pass through the lane/rank/match evidence; if no such UI surface exists, the evidence ledger must record why CLI/MCP/API output is the usable surface for this plan.

The current CLI `query` command accepts a single positional search query. Any lane/explain surface added in this plan must account for that existing parsing behavior instead of accidentally treating reserved words as the only supported query use. The implementation may use subcommands, flags, or another source-verified design, but tests must prove normal `avmatrix query "<intent>" --repo <repo>` still works and users can still search ordinary terms that resemble lane/explain command words unless the new syntax explicitly documents a non-ambiguous escape path.

The query reliability bug must be fixed with a measured retrieval path, not only worked around in skills. The implementation should first reproduce the miss in `query-health`, inspect the current query implementation and scoring reasons, then improve retrieval/ranking so expected owner files and symbols appear in the top results. A result with weak or no lexical/semantic overlap should not outrank exact path, symbol, package, generated-artifact, command-name, or resource-name matches. Query output should expose enough match reason evidence for agents to understand why a result was ranked.

The query fix must preserve and expand the original discovery value of `query`. It must not turn `query` into a pure grep command, a `context` alias, an exact-symbol-only lookup, or a tool tuned only for the AI-context plan case. Execution-flow and process results remain important, but they must be ranked behind stronger owner evidence when the process has weak overlap with the user intent. The accepted behavior change is structured broad discovery with lower wrong-owner noise, not loss of broad concept discovery.

The final query behavior must separate four roles clearly: candidate discovery through `query`, exact owner inspection through `context` or source inspection, retrieval-quality measurement through `query-health`, and validation evidence through the benchmark/evidence ledgers. Query-health output must distinguish usable retrieval from exact coverage: a usable pass means enough correct owner evidence appears to guide the agent, while an exact pass means no expected target was missed. Query-health must also make result ordering meaningful across sources; if process-symbol and definition ranks are merged, the report must define or expose a global rank/source rank so hit@5 and hit@10 are not ambiguous.

The target behavior for this plan's AI-context query case is one benchmark lane, not the whole definition of `query`: broad queries about generated AVmatrix skills and AI context must surface `internal/aicontext/aicontext.go`, embedded skill source files, analyze post-run AI context generation, setup/editor skill installation, package skill distribution, and MCP setup/resource guidance where relevant. If an exact target cannot be found, the query-health report must make the miss explicit through exact pass/fail, matched targets, missed targets, and noise reason.

The query result evidence schema should remain backward compatible while adding auditable fields. Prefer optional fields such as `queryLane`, `matchedFields`, `matchReasons`, `scoreClass`, `sourceRank`, `globalRank`, and `noiseReason` where appropriate. Do not require verbose raw scoring internals in normal output, but JSON/detail output must be strong enough for query-health and agents to explain why a result ranked.

The Web graph must have a visual orientation layer. Macro rings and node islands are not complete if they only exist as coordinates or colors. Each visible macro ring must have an on-canvas name that identifies the top-level group, such as Backend, Frontend, API, Docs, Config, Shared, Test, Unknown, or the current graph's equivalent category. Each visible island must have an on-canvas label that identifies the cluster/filter/node group, such as Function, Method, File, Route, ResolutionGap, External Reference, or the current graph's equivalent group.

Ring and island labels are layout output, not optional hover help. Users must be able to understand the graph structure without hovering individual nodes or reading only the left dashboard. Macro labels should sit near the ring center or an equivalent stable ring anchor. Island labels should sit above or near each island. Labels must remain readable during normal zoom/pan, avoid incoherent overlap with dense nodes, edges, or controls, and may simplify at far zoom levels only if users can recover the names by zooming or selecting the area.

Graph label data must come from existing graph metadata where possible: app layer, node type/filter label, semantic group, layout ring key, island key, and existing display labels. The implementation should not invent a separate hidden taxonomy when graph metadata already carries the category. Unknown/Unresolved/ResolutionGap-style areas are especially important to label because they represent investigation zones.

Graph label visibility must follow the current visible graph, not only the original graph conversion. If node type filters, graph-health filters, semantic filters, or depth filtering hide every node in a ring or island, that ring/island label must hide or update with the visible subset. Label counts and anchors should be recomputed from the Sigma graph attributes that are visible after filtering, or from an equivalent state that is proven to match the currently rendered graph.

The Phase 2 layout work must also improve visual spacing, not only add labels. Node islands are not acceptable if nodes look stacked, marched into rails, or compressed into a dense mass. Island radius and internal spiral spacing must expand from visible node count, rendered node size, minimum node gap, and spiral band gap so individual nodes remain distinguishable. Neighboring island placement must use each island's expanded radius plus a clear island gutter, and macro-ring radius must then expand from the expanded islands plus ring gutters. The layout must prefer a larger graph/canvas footprint with fit-to-view camera behavior over compressing nodes into overlap. Multipliers such as 2x or 3x may be used only as tunable spacing factors; the core behavior must be density-aware adaptive geometry. The goal is that a user can visually see each island as a separate group, see the ring structure clearly, and inspect node positions without nodes sitting on top of each other.

## Acceptance Criteria

- The source files responsible for generating `.claude/skills/avmatrix/**` are identified in the evidence ledger with exact paths and responsibilities.
- The plan records which generated outputs are validation artifacts and which files are source of truth.
- `internal/aicontext/aicontext.go` registers the final base skill set and generates a root Skills table that matches it.
- `internal/aicontext/skills/*.md` contains upgraded content for the existing six skills and source content for the new skills.
- Generated `AGENTS.md` and `CLAUDE.md` point to all final skills with clear task routing.
- A normal generation path creates every expected `.claude/skills/avmatrix/<skill>/SKILL.md` file.
- Skill content explains AVmatrix as a broad command/tool system, not a tiny workflow limited to analyze/query/impact.
- Skill content distinguishes generated files from source files and states that generated AI context files must be regenerated through AVmatrix rather than manually patched.
- README and relevant docs explain the skill system accurately enough for users and agents to know where skills come from and how to regenerate them.
- Tests protect base skill registration, generated root Skills table content, generated skill file creation, and key command-surface coverage.
- The command inventory is generated from current source or the freshly built local binary, and any stale `PATH` binary mismatch is recorded instead of used as truth.
- Final skill docs use correct CLI hyphen names and MCP underscore names for each surface.
- Package/editor setup installs the same final skill set as repository-local AI context generation, or the evidence ledger records an explicit design decision for any intentional difference.
- MCP setup/resource guidance is updated or explicitly verified as already consistent with the final command taxonomy.
- Tests prove every final base skill has non-empty embedded source content and valid frontmatter, so no final skill silently falls back to minimal placeholder content.
- Skill guidance treats broad-intent `query` output as candidate evidence and instructs agents to verify it with `context`, exact file/symbol inspection, or `query-health` before using it to select an edit surface.
- The query-health/graph-quality skill includes at least one case or documented benchmark path for AI-context skill-generation intent so this plan's observed query noise can be measured instead of normalized.
- The broad-intent query reliability bug is reproduced, root-caused, and fixed or left with an explicit failing query-health artifact and follow-up blocker. Closing this plan as complete requires the AI-context query-health case to report threshold pass and a documented exact result.
- Query result ranking favors exact owner evidence over unrelated process flows for AI context skill-generation intent, and query output or query-health evidence explains match/noise reasons.
- Query still works as a broad concept and execution-flow discovery command after the fix. Validation must prove the implementation did not collapse `query` into exact symbol lookup, grep-only search, or a `context` alias.
- Query-health reports usable pass and exact pass separately, and any exact misses are visible to agents and users instead of being hidden by a threshold pass.
- The plan has a dedicated query-health suite for this work, such as `docs/query-health/2026-05-23-avmatrix-skill-system-upgrade-suite.json`, so skill-system query reliability is not mixed into an older suite with a different purpose.
- Query-health ranking semantics are clear. Hit@5 and hit@10 must use a documented global result order or explicitly separated source-rank/global-rank fields.
- Query capability lanes are validated by representative cases so the query repair does not become an AI-context-only fix.
- Query capability lanes are usable through product surfaces. CLI help/output and MCP JSON output must expose lane discovery and explainable query results; Web/API display or pass-through must be validated if an existing UI/API query surface is in scope.
- No query capability is counted complete if it only exists as internal code, unrendered fields, or tests that users and agents cannot invoke.
- `avmatrix graph-health` is exposed in CLI help as a first-class graph-quality command, with source-verified summary, report, explain, and components behavior or an explicitly documented equivalent command design.
- Graph-health CLI output reuses the same topology, confidence, diagnostic, component, resolution-confidence, and priority semantics as the graph-health engine/Web/API surfaces instead of implementing an incompatible duplicate.
- Graph-health CLI table and JSON outputs are usable by users and agents for orphan/dead-code/unknown-connectivity investigation, including counts, candidate node identifiers, labels, paths, topology status, confidence, diagnostics, component details, and explain evidence where applicable.
- CLI parity gaps are inventoried and classified with exact source/help evidence. The plan must state which missing terminal surfaces are implemented now, which are intentionally MCP/API/Web-only, which are hidden lifecycle-only, and which remain tracked follow-up work.
- Accepted CLI parity additions expose user-facing help, table output, JSON output where appropriate, tests, and generated skill/docs guidance. No implemented capability should be counted as user-usable from CLI if it only exists as MCP/API/Web internals.
- MCP prompt templates are inventoried, classified, and updated where needed. `generate_map` must be executable without ambiguous placeholders, must use `avmatrix://repos` only to select an exact repo when no repo argument is supplied, must require stale/freshness handling, and must constrain generated architecture docs and Mermaid diagrams to graph evidence read by the agent.
- Web graph macro rings have readable on-canvas labels so users can identify each top-level ring directly on the graph.
- Web graph node islands have readable on-canvas labels so users can identify each cluster/filter/node group directly on the graph.
- Web graph node islands use adaptive spacing so nodes do not visually stack, island spiral bands are readable, neighboring islands have clear gutters, island radius grows from visible node count and minimum node gap, and macro-ring radius grows from expanded island footprints instead of forcing large islands into a fixed ring.
- Label placement, zoom behavior, and overlap behavior are validated with unit geometry checks and real browser screenshots. A graph that still requires users to infer island meaning only from color or side-panel filters is not accepted.
- Web unit tests and Web e2e/browser screenshot validation pass for graph labeling behavior.
- Full build, focused tests, generation smoke, setup/package validation if touched, and `detect-changes` pass before closure.

## Phase 0 - Generator Source Trace And Command Inventory

- [x] [P0-A] Trace the generator ownership for `.claude/skills/avmatrix/**` and record the result in the evidence ledger. The trace must identify the source skill files, the embedded filesystem owner, the `baseSkills` registry, `baseSkillContent`, `installBaseSkills`, `GenerateAIContextFiles`, the generated root Skills table, the analyze post-run caller, and any setup/package paths that copy installed skills into editor-specific directories.

- [x] [P0-B] Inventory the current source and generated skill set before implementation. Record each skill id, source file path, generated output path, source byte/line count, generated byte/line count, top headings, and whether the generated output matches the embedded source. This inventory must prove whether `.claude/skills/avmatrix/**` is source or generated validation output.

- [x] [P0-C] Inventory the current AVmatrix command surface from code/tests/help output before writing skill content. Record actual available CLI commands, MCP tools, MCP resources, setup/package commands, runtime commands, group/cross-repo commands, API-surface commands, and graph-quality/accuracy commands. The evidence must distinguish implemented commands from planned or absent commands so skills do not document non-existent behavior as real.

- [x] [P0-D] Build the skill routing matrix from the command inventory. Map every current command/tool/resource family to one primary skill and any secondary skill references, then record the final decision in evidence. The matrix must include the existing six skills and the proposed new skills for graph quality, API surface, cross-repo work, runtime/packaging, and AI context generation.

- [x] [P0-E] Compare the command surface exposed by the current `PATH` binary, `go run .\cmd\avmatrix --help`, and the binary produced by the full build gate. Record mismatches in evidence and use only the current source or freshly built local binary as the source of truth for skill content.

- [x] [P0-F] Trace the package/editor skill source path from `setupInstallSkillsTo` and package lifecycle code. Record whether package-root `skills/` exists, how packaged installs are expected to contain skills, and what code/test change is needed so packaged/editor skills cannot drift from embedded AI-context skills.

- [x] [P0-G] Inspect MCP resource/setup guidance such as `avmatrix://setup` and the source that renders it. Record whether it already matches the final command taxonomy or must be updated alongside `internal/aicontext/aicontext.go`.

- [x] [P0-H] Audit embedded base skill source validity before editing. Record each source file's frontmatter `name`, `description`, non-empty body status, and whether `baseSkillContent` would read real content or fall back.

- [x] [P0-I] Record the broad-intent query reliability bug as baseline evidence. Run and record query results for AI context skill generation, setup/editor skill installation, package skill distribution, and MCP setup/resource guidance; then verify the correct owner surfaces with `context` and exact file inspection. The evidence must show when query output is only a candidate, when it misses expected owner files, and why that miss is dangerous for agent edit-surface selection.

- [x] [P0-J] Create a dedicated query-health suite for this plan, for example `docs/query-health/2026-05-23-avmatrix-skill-system-upgrade-suite.json`, and add a case that reproduces the AI-context skill-generation query miss before any retrieval fix. The case must list expected files/symbols, actual top results, hit@5/hit@10, exact target coverage, missed targets, and noise reason so the bug is measurable without mixing this plan into an older suite with a different purpose.

- [x] [P0-K] Trace the current query implementation and scoring path before changing it. Record the exact files and symbols that rank definitions, process symbols, execution flows, lexical tokens, path/name matches, semantic fields, filters, and result limits. The trace must explain why unrelated launcher/resolution/frontend flows outrank the expected AI-context owners.

- [x] [P0-L] Inventory the current user-facing query surfaces before changing them. Record CLI help/output for `avmatrix query`, MCP `query` tool schema/output, `query-health` output, and any existing Web/API query/search surface. The evidence must state which surfaces users and agents can actually invoke today and which query-lane evidence is currently hidden or absent.

## Phase 1 - Core Query Reliability Repair

This is a large blocking phase. The rest of the skill-system upgrade depends on agents being able to use `query` without being sent to plausible but wrong code regions. This phase must treat broad-intent query misses as a product reliability bug, not as a documentation caveat. Skill text may still teach verification discipline, but the product fix must improve query itself.

- [x] [P1-A] Build the query reliability baseline for this plan before fixing ranking. The baseline must run broad-intent queries for AI context skill generation, setup/editor skill installation, package skill distribution, MCP setup/resource guidance, and command-surface discovery. For each query, record expected owner files/symbols, actual top results, unrelated top results before first expected owner, hit@5, hit@10, exact matched targets, missed targets, and noise reason in the evidence and benchmark ledgers.

- [x] [P1-B] Root-cause the current `query` retrieval and ranking pipeline. Trace the exact implementation path for CLI/MCP query from input text to returned definitions/processes, including tokenization, lexical matching, process scoring, definition scoring, path/name scoring, App Layer/Functional Area boosts, docs/test filtering, result limits, and result diversification. The evidence must explain why launcher/resolution/frontend flows can outrank `internal/aicontext` owners for the AI-context intent.

- [x] [P1-C] Define the Query Capability Taxonomy and relevance contract for broad-intent repository work. The contract must describe `query` as an umbrella command with retrieval lanes such as owner discovery, concept discovery, execution-flow discovery, API surface discovery, graph-quality discovery, docs/setup/AI-context discovery, command-surface discovery, and cross-repo discovery. It must also define that `context` performs exact symbol inspection after a candidate owner is found. Exact file path, exact symbol name, generated artifact name, command name, resource URI/name, package/module name, and high lexical overlap are strong evidence; unrelated execution-flow names with weak overlap must not outrank those owners. The contract must also define how docs/tests/examples are ranked when the query is about product source versus documentation.

- [x] [P1-D] Implement the query retrieval/ranking fix through the capability taxonomy rather than through an AI-context-only patch. The implementation must improve owner discovery using proven signals such as exact symbol/file/path matches, command/resource names, generated artifact names, package/module names, lexical token overlap, and process/definition relevance. It must preserve concept and execution-flow discovery, demote unrelated process flows that lack meaningful overlap with the query intent, and keep useful process/API/graph-quality/docs-command results when they have clear evidence. The implementation must keep broad discovery behavior intact and must not replace `query` with grep-only matching or exact-symbol-only lookup.

- [x] [P1-E] Add result reason output where needed so query results are auditable. Query or query-health output must expose enough match evidence for an agent/user to see why a result ranked highly, such as query lane, matched tokens, matched path/symbol/command/resource fields, score class, source rank, global rank, and noise reason for expected misses. This must not turn into verbose raw internals by default; use summary fields or JSON detail where appropriate.

- [x] [P1-F] Add the user-facing query lane command surface. Preserve the existing `avmatrix query "<intent>" --repo <repo>` command, then add clear lane discovery and explainable query usage through subcommands or flags such as `avmatrix query lanes`, `avmatrix query explain "<intent>" --repo <repo> --json`, `--lane <lane>`, or an equivalent source-verified design. Help output must describe each lane and show how a user or agent can get lane/rank/match evidence.

- [x] [P1-F2] Protect `query` CLI compatibility while adding lane/explain syntax. Current source uses one positional `query <search_query>`, so the implementation must choose a non-ambiguous Cobra design, flag design, or escape behavior, then add tests proving ordinary broad query strings still work, lane/explain help is discoverable, JSON evidence is invokable, and a search term that looks like a lane/explain command is not silently routed to the wrong behavior.

- [x] [P1-G] Add focused unit tests for query scoring, lane assignment, and filtering. Tests must cover exact owner file/symbol/path matches, generated artifact names such as `AGENTS.md` and `.claude/skills/avmatrix`, command/resource-name matches such as `setupResource` and `query-health`, execution-flow results with real overlap, API/graph-quality/docs-command lane examples where supported by current data, and a negative case where unrelated launcher/resolution/frontend flows must not outrank stronger owner evidence.

- [x] [P1-H] Add or update CLI/MCP query and query-health integration tests. Tests must prove the AI-context intent returns expected owner surfaces, representative non-AI-context query lanes still return useful results, threshold pass is reported separately from exact pass, missed-target reporting works, source/global ranking semantics are clear, CLI lane/explain commands are invokable, MCP output exposes equivalent machine-readable evidence, and unrelated high-scoring results cannot silently be accepted as a passing exact result.

- [x] [P1-I] Run the updated query-health suite after the query fix and record threshold pass/fail, exact pass/fail, expected targets, matched targets, missed targets, query lane coverage, unrelated top-result count, source/global rank behavior, user-facing command outputs, and remaining noise reason in the benchmark ledger.

- [x] [P1-J] Finalize skill-guidance requirements only after the query behavior and command surface are measured. Do not edit embedded skill content in this phase. Record the exact guidance that Phase 3 must apply to `avmatrix-exploring`, `avmatrix-debugging`, `avmatrix-graph-quality`, and `avmatrix-ai-context`: `query` is the right broad discovery command with multiple usable capability lanes, and broad results still need verification with `context` or exact source inspection when selecting edit surfaces.

- [x] [P1-K] Add a closure gate for query reliability before moving to the rest of the skill-system implementation. The gate is satisfied only when the AI-context query-health case has recorded threshold and exact results, the remaining missed targets if any are explicit, the usable CLI/MCP query lane surfaces are validated, and the plan/evidence/benchmark ledgers state whether the query bug is fixed or still has a tracked blocker.

- [x] [P1-L] Add broad-discovery regression checks for `query` after the ranking fix. Use representative non-AI-context intents across the capability taxonomy, including at least one intent that should naturally return execution-flow or process candidates. Record the before/after top results and prove the fix reduced wrong-owner noise without removing the original concept-to-flow discovery capability.

## Phase 1.5 - Graph Health CLI Surface

This phase closes the gap where graph-health exists as engine metadata and Web/API behavior but cannot be invoked as a first-class CLI command. It must create a usable graph-quality surface for terminal workflows without duplicating or redefining graph-health semantics separately from the existing engine/API/Web behavior.

- [x] [P1.5-A] Trace the current graph-health implementation before editing. Record the files and symbols that compute topology status, counted/excluded edges, confidence, diagnostics, resolution confidence, component IDs/sizes, report priority, explain payloads, Web filters, and API routes. The trace must include `internal/graphhealth/compute.go`, `internal/httpapi/graph.go`, `avmatrix-web/src/lib/graph-health-filters.ts`, the CLI root command registration in `internal/cli/command.go`, and any better owner found by source inspection.

- [x] [P1.5-B] Define the `graph-health` CLI contract from the traced source behavior. The contract must specify table and `--json` output for summary, report, explain, and components behavior or an explicitly justified equivalent design. It must define required inputs such as `--repo`, node id/name selection for explain, `--limit`, graph freshness expectations, output fields, ordering/priority, and how unknown or missing nodes are reported.

- [x] [P1.5-C] Refactor graph-health report/explain logic into a shared reusable owner if the current API code owns behavior that the CLI needs. The shared owner must preserve Web/API semantics for topology status, confidence, diagnostics, component membership, resolution health buckets, counted/excluded relationship samples, and triage priority. Do not copy-paste a second CLI-only implementation that can drift.

- [x] [P1.5-D] Implement the user-facing CLI command surface. Add `avmatrix graph-health` to the CLI command tree with discoverable help and subcommands or flags for summary, report, explain, and components. The command must support human-readable table output and `--json`, use the current indexed graph for the selected repo, and fail clearly when no fresh graph exists or a requested node cannot be found.

- [x] [P1.5-E] Add focused CLI/unit tests for graph-health. Tests must cover command registration/help output, summary counts, report ordering/limit behavior, explain by node id or source-verified selector, components output, JSON shape, missing-node errors, stale/missing graph handling, and parity with shared graph-health semantics used by API/Web where that can be tested deterministically.

- [x] [P1.5-F] Add graph-health command guidance to generated command surfaces after the CLI behavior exists. Update `internal/aicontext/aicontext.go`, `internal/aicontext/skills/avmatrix-graph-quality.md` or its planned source, `avmatrix-cli`/`avmatrix-guide` content, MCP setup/resource guidance if it lists CLI command families, and README/user docs that describe graph-quality commands. The guidance must explain that `graph-health` audits topology/diagnostics, while `query-health`, `resolution-inventory`, and `source-site-accuracy` answer different graph-quality questions.

- [x] [P1.5-G] Validate graph-health CLI with real command smoke output after the full build. Run representative commands such as `avmatrix graph-health summary --repo AVmatrix`, `avmatrix graph-health report --repo AVmatrix --limit 20 --json`, `avmatrix graph-health components --repo AVmatrix --json`, and one `graph-health explain` command against a real node discovered from the report. Record counts, candidate examples, exact command output summaries, and any limitation in the evidence and benchmark ledgers.

## Phase 1.6 - CLI Parity Audit And Missing Command Surfaces

This phase makes the broader command-surface parity decision explicit. It must not assume every MCP/API/Web capability needs a CLI command, but it must prevent accidental gaps where implemented product capabilities are unavailable to terminal users and AI agents using the CLI.

- [x] [P1.6-A] Inventory the current CLI command tree from both the freshly built local binary and source registration. Record top-level commands, subcommands, hidden commands, visible help text, and command families from `internal/cli/command.go`, `internal/cli/*_command.go`, `avmatrix --help`, and `go run .\cmd\avmatrix --help`.

- [x] [P1.6-B] Inventory non-CLI command surfaces from source. Record MCP tools such as `rename`, `route_map`, `tool_map`, `shape_check`, `api_impact`, group tools, and graph tools; MCP resources such as repo context, clusters, processes, schema, cluster detail, and process detail; HTTP/Web endpoints such as graph explain/report, grep, search, processes, process, clusters, cluster, file, repo, analyze, embed, and session; and any Web-only controls that represent product capability.

- [x] [P1.6-C] Build a CLI parity matrix. For each MCP/API/Web/resource surface, classify it as `has_cli`, `must_add_cli`, `covered_by_existing_cli`, `mcp_api_web_only_by_design`, `hidden_lifecycle_only`, or `follow_up`. The matrix must include exact evidence for `rename`, `route_map`, `tool_map`, `shape_check`, `api_impact`, graph explain/report, grep, search, processes, process detail, clusters, cluster detail, and graph-health.

- [x] [P1.6-D] Define the accepted CLI command design for parity gaps. The design must decide whether to use grouped commands or top-level commands, for example `avmatrix api route-map`, `avmatrix api tool-map`, `avmatrix api shape-check`, `avmatrix api impact`, `avmatrix rename`, `avmatrix graph grep/search/processes/process/clusters/cluster`, or another source-verified structure. The design must preserve current commands and avoid ambiguous parsing with existing `query`, `group`, and hidden lifecycle commands.

- [x] [P1.6-E] Implement the CLI parity commands accepted for this plan, or explicitly record why a surface is intentionally not implemented in this plan. For each implemented command, reuse the existing MCP/API/Web owner logic where possible, support repo selection, provide readable table output, add `--json` where the output is structured, and avoid copying divergent semantics.

- [x] [P1.6-F] Add tests for accepted CLI parity commands and for intentionally hidden/internal commands. Tests must cover help output, command registration, representative success output, JSON shape, error handling, and non-regression that hidden lifecycle helpers such as `package` or `hook` are not promoted accidentally if they are not user-facing.

- [x] [P1.6-G] Update generated skills, setup/resource guidance, README/docs, and command-surface tables after the parity decision. Guidance must explain which surfaces are available through CLI, which remain MCP/API/Web/resource-only by design, and how agents should choose the right surface without falling back to incomplete command lists.

- [x] [P1.6-H] Validate accepted CLI parity commands with real smoke output after the full build. Record representative commands, outputs, and any intentionally deferred surface in the evidence and benchmark ledgers. At minimum, validation must include each command family accepted in P1.6-D and a source/help check proving the final CLI list matches the documented command surface.

## Phase 1.7 - MCP Prompt Template Accuracy

This phase fixes MCP prompts as product command surfaces. It belongs after CLI parity inventory because prompts must be classified beside tools/resources/CLI/Web/API surfaces, and before skill/docs work because README, generated AI context, and skills must document the final prompt behavior rather than stale prompt text.

- [x] [P1.7-A] Inventory the current MCP prompt surface from source and runtime behavior. Record every prompt name, description, argument, generated prompt body, tests, README references, MCP resource references, and whether the prompt is meant to create files, produce analysis, guide command use, or only explain a workflow. The inventory must include `detect_impact` and `generate_map`.

- [x] [P1.7-B] Audit `generate_map` for executable correctness. The audit must verify each referenced resource exists, identify the placeholder behavior when `repo` is omitted, check repo and process URI escaping requirements, verify whether stale graph handling is stated, identify how "top 5 most important processes" is selected, and list every instruction that could cause architecture claims or Mermaid edges to be invented without graph evidence.

- [x] [P1.7-C] Redesign `generate_map` as a precise architecture-map prompt template. If `repo` is provided, the prompt must use that exact repo name with safe resource URI guidance. If `repo` is omitted, it must first read `avmatrix://repos`, select the repo whose indexed path matches the current workspace when possible, use the single listed repo when only one exists, and stop to ask the user when multiple repos exist and no match is provable. It must never use `{name}` as a real resource URI.

- [x] [P1.7-D] Add freshness and evidence rules to `generate_map`. The prompt must tell the agent to verify graph freshness from repo context or active instructions, run `avmatrix analyze --force` when graph work requires a refresh, and write architecture documentation only from resources, tools, or command output it actually read. It must require uncertainty notes when clusters, processes, unresolved references, graph-health, or query-health evidence is incomplete.

- [x] [P1.7-E] Define deterministic process and cluster selection rules for `generate_map`. The prompt must tell the agent how to choose representative processes and clusters, such as by user-facing runtime relevance, process type, step count, route/tool/API involvement, graph centrality if available, or explicit user request. The prompt must record the reason for each selected process instead of saying only "top 5 most important processes".

- [x] [P1.7-F] Update `detect_impact` if the inventory shows the prompt is stale. It must align with the current AVmatrix rules: refresh before graph-based work when required, use change detection before implementation commits, treat HIGH/CRITICAL as blast-radius warnings rather than an edit ban, and point to the correct CLI/MCP surface names without replacing repository rules.

- [x] [P1.7-G] Add or update MCP prompt tests. Tests must cover `prompts/list`, `prompts/get` for `generate_map` with and without `repo`, no literal `{name}` continuation as an actionable URI, presence of `avmatrix://repos` discovery instructions, repo selection ambiguity handling, freshness language, evidence-only architecture claims, process selection criteria, and snapshot updates for expected prompt text.

- [x] [P1.7-H] Update README, MCP setup/resource guidance, generated skills, and AI-context generator text that mention MCP prompts. Documentation must label `generate_map` and `detect_impact` as MCP prompts, not CLI commands, and must explain what they automate, what they require from the agent, and how they relate to CLI/MCP tools/resources.

- [x] [P1.7-I] Validate MCP prompt behavior after the full build. Run focused MCP tests and one real `prompts/get` smoke check for `generate_map` and `detect_impact`; record the prompt names, argument schemas, key generated instruction fragments, and any intentional limitation in the evidence ledger.

## Phase 2 - Graph Labeling And Visual Orientation Layer

- [ ] [P2-A] Record the graph orientation problem from `reports/problem/screenshot_1779517751.png` in the evidence ledger. The evidence must state that visible rings and islands currently lack direct names on the graph, making users infer meaning from color or side-panel filters instead of reading the graph itself.

- [ ] [P2-B] Trace the current Web graph layout and rendering owners before editing. Record the files and symbols that produce ring/island coordinates, node type colors, app-layer rings, island keys, Sigma rendering, overlays, and graph filter/legend state. The trace must include `avmatrix-web/src/lib/graph-adapter.ts`, `avmatrix-web/src/components/GraphCanvas.tsx`, `avmatrix-web/src/hooks/useSigma.ts`, and any better owner discovered by source inspection.

- [ ] [P2-C] Define the visual label data contract. Macro ring labels must be derived from app-layer/top-level ring metadata where available. Island labels must be derived from node type/filter/island metadata where available. The contract must include label kind, display text, source key, anchor coordinates, visible node count, and fallback text for unknown/custom categories. It must not rely on color alone.

- [ ] [P2-D] Implement macro ring labels as a visible graph orientation layer. Each currently visible macro ring must show a stable name near the ring center or equivalent anchor, remain readable during normal zoom/pan, and avoid covering controls or dense node areas. Ring labels must update when node type filters, graph-health filters, semantic filters, or depth filtering remove every visible node in a ring.

- [ ] [P2-E] Implement island labels as a visible graph orientation layer. Each major currently visible node island must show a stable name above or near the island, such as Function, Method, File, Route, ResolutionGap, External Reference, or the graph's current group label. Island labels must update when filters/depth change the visible subset. Labels may be simplified when zoomed far out, but users must be able to recover names by zooming or selecting the area.

- [ ] [P2-F] Add label placement and visibility tests. Unit tests must verify ring/island label metadata, stable anchors, visible-node counts, filter/depth visibility behavior, zoom visibility rules, and overlap guardrails where they can be tested deterministically.

- [ ] [P2-F2] Improve adaptive island and ring spacing in the Web graph layout. Update the layout geometry so each island computes a minimum node gap, spiral band gap, rendered-node-size allowance, visible-node-count footprint, and resulting island radius before island placement. Neighboring islands must be spaced by `radiusA + radiusB + islandGutter` or an equivalent collision-free rule, and each macro ring must compute its radius from the expanded island footprints plus a ring gutter instead of pushing large islands into a fixed-size circle. The implementation may expose tunable spacing multipliers such as 2x or 3x, but those multipliers must feed the density-aware geometry rather than replace it with a constant diameter. The graph/canvas is allowed to become larger with fit-to-view camera behavior when repo size demands it; it must not solve density by compressing nodes, shrinking gaps below the minimum, or creating rail-like/marched-node shapes. Record deterministic geometry tests proving node gap, spiral band gap, island gutter, macro-ring expansion, and no-overlap behavior, plus browser screenshot evidence at desktop and smaller viewports.

- [ ] [P2-G] Add Web UI validation for readable labels. Extend runtime diagnostics or test selectors so browser tests can count visible ring/island labels and inspect representative label text without relying only on screenshot pixels. Use browser/e2e validation and screenshots to prove that macro ring labels and island labels are visible, readable, update after filters, and do not incoherently overlap nodes, edges, or controls on representative desktop and smaller viewports.

- [ ] [P2-H] Update user-facing graph guidance only if the UI needs a short label toggle or explanation. Do not add in-app instructional prose as a substitute for labels; the graph itself must name the rings and islands.

## Phase 3 - Embedded Skill Source Upgrade

- [ ] [P3-A] Upgrade the six existing embedded skill Markdown files in `internal/aicontext/skills/`. Each file must become a practical task guide with command choices, when to use each AVmatrix surface, validation expectations, and current limitations. `avmatrix-impact-analysis` must explain that HIGH/CRITICAL is blast-radius evidence to report and account for, not a blanket prohibition against required work.

- [ ] [P3-B] Add the new embedded source skill files under `internal/aicontext/skills/`: `avmatrix-graph-quality.md`, `avmatrix-api-surface.md`, `avmatrix-cross-repo.md`, `avmatrix-runtime-packaging.md`, and `avmatrix-ai-context.md`. Each new skill must contain concrete usage guidance, command examples based on implemented commands, expected outputs, and validation notes. The graph-quality skill must include the `graph-health` CLI command from Phase 1.5 after that command exists.

- [ ] [P3-C] Update the base skill registry and generated Skills table in `internal/aicontext/aicontext.go`. The registry and generated table must include all final skills, use repo-agnostic descriptions, and avoid splitting AVmatrix into misleading MCP-only versus CLI-only capability lists.

- [ ] [P3-D] Add or update `internal/aicontext` tests so generated root files and generated base skills are protected. Tests must assert the final skill ids, generated `.claude/skills/avmatrix/<skill>/SKILL.md` paths, generated Skills table links, and representative key phrases for the new command surfaces.

- [ ] [P3-E] Add coverage tests that prevent the guide from regressing back to a six-skill or analyze/query/impact-only view. The tests should check the generated guidance for the AI context skill, graph-quality skill, API-surface skill, cross-repo skill, runtime/packaging skill, and a current command-surface fragment confirmed in Phase 0.

- [ ] [P3-F] Add or update tests that validate command naming by surface. The test should protect at least one CLI-only hyphenated command such as `query-health`, one MCP underscore tool such as `route_map`, and one statement that does not invent a CLI spelling for an MCP-only tool.

- [ ] [P3-G] Add frontmatter/source-content tests for every final embedded base skill. The test must fail if a registered base skill is missing its embedded Markdown file, has empty content, has mismatched `name`, lacks `description`, or would rely on `fallbackBaseSkillContent`.

- [ ] [P3-H] Update `avmatrix-exploring`, `avmatrix-debugging`, `avmatrix-graph-quality`, and `avmatrix-ai-context` skill content using the measured Phase 1 guidance. The guidance must state that broad intent query is a multi-lane candidate discovery command, explain the user-facing CLI/MCP query lane commands or flags, state that `context` is preferred when an exact symbol is known, and tell agents that noisy/missed query results must be recorded as graph-quality/query-health evidence.

- [ ] [P3-I] Add or extend skill-facing query-health guidance for this plan's AI-context intent. The skill should point to the query-health suite/case created in Phase 0/Phase 1, explain threshold versus exact pass, and tell agents to record missed targets instead of treating partial retrieval as complete.

## Phase 4 - Setup, Package, And Documentation Integration

- [ ] [P4-A] Verify the analyze post-run path installs the expanded base skill set through the same normal generation path that creates `AGENTS.md` and `CLAUDE.md`. If tests currently assert the old six-skill count or specific old table rows, update them to assert the new final set.

- [ ] [P4-B] Verify setup/editor installation behavior for the expanded embedded skill set. Inspect and test `setupInstallEditorSkills` and related setup command behavior so supported editor skill directories receive the same final skill content without relying on generated repository-local `.claude/skills/avmatrix/**` as source.

- [ ] [P4-C] Verify package/runtime distribution behavior for the expanded embedded skill set. If packaging tests or package assembly code enumerate skills, update them so the packaged tool can generate and install the final skill set from embedded source files.

- [ ] [P4-D] Reconcile package-root `skills/` with embedded `internal/aicontext/skills/*.md`. Either make the package/setup path materialize or copy from the same canonical skill source, or document and test a deliberately equivalent packaged `skills/` directory. The output must prove `avmatrix setup` installs the same final base skill ids and content family as `avmatrix analyze --force` generates in `.claude/skills/avmatrix/**`.

- [ ] [P4-E] Update MCP setup/resource guidance if Phase 0 finds stale command/tool/resource text. This includes the source that renders `avmatrix://setup`, MCP tool reference tables, and any setup onboarding text used by agents.

- [ ] [P4-F] Update README and relevant user-facing docs that describe AVmatrix skills, AI context generation, setup, or usage. The docs must tell users that `AGENTS.md`, `CLAUDE.md`, and `.claude/skills/avmatrix/**` are generated by AVmatrix, and that source changes belong in the embedded skill source and generator code.

- [ ] [P4-G] Search the active documentation for stale six-skill-only guidance, stale package-root skill assumptions, or stale wording that treats MCP and CLI as separate incomplete command lists. Update current guides and README-style docs; leave historical ledgers untouched unless they are actively reused as user guidance.

- [x] [P4-H] Eliminate AVmatrix executable build drift. Update the full build and launcher/runtime path resolution so `avmatrix\bin\avmatrix.exe` is the only production AVmatrix CLI executable built and used from this checkout. The work must remove or stop relying on an independently built `avmatrix-launcher\server-bundle\avmatrix.exe`, update launcher/server-wrapper code to run the canonical executable for `serve`, update process cleanup to target the canonical executable path plus launcher support processes, and update tests/docs that still describe the bundle copy as an executable authority.

## Phase 5 - Regeneration And Validation

- [ ] [P5-A] Run the full build gate before tests: `powershell -ExecutionPolicy Bypass -File avmatrix-launcher\build.ps1`. Record the command result in evidence.

- [ ] [P5-B] Run focused backend/CLI/MCP tests for AI context generation, setup, packaging, query, query-health, graph-health, CLI parity surfaces, and MCP prompt templates touched by the implementation. The minimum expected test scope is `go test .\internal\mcp .\internal\cli .\internal\aicontext .\internal\graphhealth .\internal\httpapi -count=1`, expanded as needed if package/setup code outside those packages changes.

- [ ] [P5-C] Run Web unit tests after graph labeling work: `cd avmatrix-web; npm run test`. Record exact pass/fail counts and any focused label/layout tests in evidence.

- [ ] [P5-D] Run Web e2e/browser validation after graph labeling work: `cd avmatrix-web; npm run test:e2e` or a focused Playwright command expanded to full e2e before closure. Capture desktop and smaller-viewport screenshots that prove macro ring labels and island labels are readable and not incoherently overlapping the graph.

- [ ] [P5-E] Run the normal generation path with `avmatrix analyze --force` and no `--skip-agents-md`. Verify generated `AGENTS.md`, generated `CLAUDE.md`, and `.claude/skills/avmatrix/**` contain the final skill set and expected content fragments.

- [ ] [P5-F] Compare source and generated skill inventories after regeneration. Record final skill count, generated file paths, byte/line counts, and any intentional generated differences in the benchmark ledger.

- [ ] [P5-G] Validate setup/package behavior if Phase 4 changed setup or package code. Record the exact command outputs and installed/packaged skill file inventories in evidence and benchmark ledgers.

- [x] [P5-G2] Validate canonical executable behavior after the full build. Record `Get-ChildItem avmatrix\bin\avmatrix.exe`, `avmatrix\bin\avmatrix.exe version`, `avmatrix\bin\avmatrix.exe query --help`, `avmatrix\bin\avmatrix.exe query --lanes --json`, `avmatrix\bin\avmatrix.exe query "<representative intent>" --repo AVmatrix --limit 5 --explain`, and `avmatrix\bin\avmatrix.exe query-health --repo AVmatrix --suite docs\query-health\2026-05-23-avmatrix-skill-system-upgrade-suite.json --limit 10`. Also record that no independent `avmatrix-launcher\server-bundle\avmatrix.exe` remains as a production/runtime command source.

- [ ] [P5-H] Validate MCP setup/resource output if Phase 4 touched MCP resources. Record the exact `avmatrix://setup` or equivalent resource output check in evidence so the generated guidance and MCP-facing guide are proven consistent.

- [ ] [P5-I] Run the dedicated query-health suite updated for this plan's AI-context skill-generation intent and representative query capability lanes. Record threshold pass/fail, exact pass/fail, expected targets, matched targets, missed targets, source/global rank behavior, query lane coverage, and noise reason in benchmark/evidence.

- [x] [P5-J] Run user-facing query lane smoke commands after the full build. Validate normal query behavior, lane discovery, explainable JSON output, MCP query JSON evidence, and the dedicated query-health suite command; record exact commands and representative output in evidence and benchmark.

- [ ] [P5-K] Run the query broad-discovery regression check from Phase 1 and record whether execution-flow candidates still appear with meaningful match evidence after the ranking fix.

- [ ] [P5-K2] Run user-facing graph-health CLI smoke commands after the full build. Validate help output, summary, report, components, explain, `--json`, missing-node error behavior, and parity with Web/API graph-health semantics where applicable. Record exact commands and representative output in evidence and benchmark.

- [ ] [P5-K3] Run user-facing CLI parity smoke commands after the full build. Validate every command family accepted in Phase 1.6, plus the source/help inventory that proves intentionally MCP/API/Web-only or hidden lifecycle surfaces are classified correctly. Record exact commands and representative output in evidence and benchmark.

- [ ] [P5-K4] Run MCP prompt smoke validation after the full build. Validate `prompts/list` and `prompts/get` for `generate_map` and `detect_impact`, including prompt descriptions, arguments, repo-selection instructions, freshness instructions, evidence-only architecture guidance, and correct distinction between MCP prompts, MCP tools, MCP resources, CLI commands, and Web/API surfaces. Record representative output in evidence and benchmark.

- [x] [P5-L] Run `detect-changes` before commit and record the affected scope. Commit the implementation slice after checklist items and ledgers are updated.

## Phase 6 - Zero-Trust Closure Review

- [ ] [P6-A] Review the codebase and documentation for old assumptions after implementation. Search for old six-skill-only tables, old `avmatrix-cli` descriptions that omit current command families, stale generated-output instructions, any direct-edit guidance for `.claude/skills/avmatrix/**`, and any graph guidance that expects users to infer rings/islands only from color or side-panel filters; fix active docs that would mislead users or agents.

- [ ] [P6-B] Re-run the final validation commands required by this plan after the closure review changes. Record final pass/fail counts, generated inventory, setup/package inventory if applicable, Web label screenshot evidence, and any remaining limitation in the evidence and benchmark ledgers.

- [ ] [P6-C] Mark the plan complete only after source files, generated validation output, tests, graph-health CLI validation, CLI parity validation, MCP prompt validation, Web graph label validation, docs, benchmark ledger, evidence ledger, and commit state all agree on the final skill set, graph orientation labels, graph-health command surface, CLI parity decisions, MCP prompt behavior, and the AI-context query-health case has recorded threshold and exact results.

## Phase 7 - Final README And AI Context Command Documentation Sync

This phase runs after all implementation, CLI parity, graph-health, query, Web graph, and validation phases have finished enough to know the final command surface. Its purpose is to prevent new commands from existing only in code while README, generated `AGENTS.md` / `CLAUDE.md` content, and generated skills still describe an older or incomplete AVmatrix command set.

- [ ] [P7-A] Inventory the final user-facing command surface from the freshly built binary, MCP tool registration, MCP prompt registration, MCP setup/resource guidance, and any accepted Web/API-only surfaces from Phase 1.6. Record the final command families, exact command names, prompt names, intentionally hidden lifecycle commands, and intentionally non-CLI surfaces in the evidence ledger before editing documentation.

- [ ] [P7-B] Update `README.md` so new commands and command families added or accepted by this plan are visible to normal users. The README must explain the practical purpose of the new graph-quality, query-health, source-site accuracy, resolution inventory, CLI parity, setup, MCP, runtime, and Web/API command surfaces without presenting AVmatrix as only analyze/query/impact.

- [ ] [P7-C] Update `internal/aicontext/aicontext.go` so generated `AGENTS.md` and `CLAUDE.md` command guidance includes the final command surface. The generated text must stay repo-agnostic, preserve the auto-updated project/index summary between `<!-- avmatrix:start -->` and `<!-- avmatrix:end -->`, and avoid splitting MCP tools and CLI commands into misleading incomplete lists.

- [ ] [P7-D] Update any embedded skill source files under `internal/aicontext/skills/*.md` that reference command selection, graph quality, API surface inspection, query behavior, setup, packaging, or runtime usage. Each changed skill must point users and agents to the actual implemented command names and must explain when a command is CLI, MCP, Web/API, resource-only, or intentionally hidden.

- [ ] [P7-E] Update tests that protect README-adjacent guidance and generated AI context output. The tests must fail if new accepted commands are missing from generated command tables, if old six-skill/analyze-query-impact-only wording returns, or if generated output invents CLI names for MCP/API/Web-only surfaces.

- [ ] [P7-F] Regenerate the AI context with `avmatrix analyze --force` and no `--skip-agents-md`. Verify generated `AGENTS.md`, generated `CLAUDE.md`, and generated `.claude/skills/avmatrix/**` contain the final command guidance from `internal/aicontext/aicontext.go` and embedded skill sources rather than stale generated output.

- [ ] [P7-G] Re-run the full build gate, focused docs/generator tests, and final command-surface smoke checks after the README and AI-context updates. Record exact command output summaries, generated file checks, and any remaining intentionally deferred documentation surface in the evidence and benchmark ledgers.

- [ ] [P7-H] Close the plan only after README, `internal/aicontext/aicontext.go`, embedded skills, generated `AGENTS.md` / `CLAUDE.md`, validation evidence, benchmark evidence, and the final CLI/MCP prompt/MCP tool/resource/Web/API command inventory all describe the same command surface.
