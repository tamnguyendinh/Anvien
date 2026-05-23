# AVmatrix Skill System Upgrade Plan

Date: 2026-05-23

Status: Planned

Companion files:

- Benchmark ledger: [2026-05-23-avmatrix-skill-system-upgrade-benchmark.md](2026-05-23-avmatrix-skill-system-upgrade-benchmark.md)
- Evidence ledger: [2026-05-23-avmatrix-skill-system-upgrade-evidence.md](2026-05-23-avmatrix-skill-system-upgrade-evidence.md)

## Master rules

1. Use AVmatrix for codebase analysis and impact checks while working on implementation slices in this plan.
2. As each task is completed, update the corresponding checklist item immediately.
3. Run the full build gate before testing; include focused backend/CLI/setup/package validation for generated skill behavior, and include Web/e2e validation only if Web behavior changes.
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

There is a query reliability bug. A broad intent query can return plausible but wrong code regions. During plan review, broad queries about AI context skill generation returned unrelated launcher, resolution-gap, and frontend/backend flows instead of the expected `internal/aicontext`, analyze post-run, setup, and package-skill surfaces. That behavior is dangerous for agent work because a query that cannot identify the right region can send the agent to edit or reason about the wrong code. This is not a minor documentation issue or normal behavior to accept. `query` is a core AVmatrix feature and must be able to locate the correct work area for broad intent. Until the bug is fixed, broad `query` output must be treated as candidate retrieval and verified by symbol-level `context`, exact file/symbol inspection, or `query-health` evidence.

## Scope

Implementation may touch:

- `internal/aicontext/aicontext.go`;
- `internal/aicontext/skills/*.md`;
- tests under `internal/aicontext`;
- CLI/setup/package tests under `internal/cli` if they assert skill counts, skill paths, generated output, or packaging behavior;
- query implementation, query user-facing command surfaces, query output formatting, and tests under `internal/mcp`, `internal/cli`, query-health suites, and any query ranking/scoring helpers if source inspection proves broad-intent query misses expected owners;
- MCP setup/resource guidance under `internal/mcp` if it contains tool, resource, setup, or command-surface reference text;
- README and user-facing docs that explain generated AVmatrix skills, AI context setup, or AVmatrix command surfaces;
- packaging/setup code only if source inspection proves the expanded embedded skill set is not installed or packaged correctly.

Out of scope unless source inspection proves it is required:

- changing the behavior of AVmatrix graph analysis commands;
- changing Web UI graph rendering;
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
| `avmatrix-cli` | Complete CLI command guide, including analyze/status/query/context/impact/detect-changes/cypher/rename/augment/group/setup/serve/mcp/package/wiki/hook/version and any current accuracy commands confirmed by source. |
| `avmatrix-graph-quality` | Query health, source-site inventory, resolution inventory, edge accuracy, ResolutionGap/UnresolvedSymbol review, and benchmark comparison. |
| `avmatrix-api-surface` | API routes, MCP tools, contract shape checks, API impact, generated Web contracts, handlers, and consumers. |
| `avmatrix-cross-repo` | Group repositories, cross-repo query/contracts/status/sync, and multi-repo analysis guidance. |
| `avmatrix-runtime-packaging` | `serve`, `mcp`, `setup`, launcher, packaged runtime, package preparation, runtime cleanup, and startup validation. |
| `avmatrix-ai-context` | Generated `AGENTS.md`, `CLAUDE.md`, `.claude/skills/avmatrix/**`, source-vs-generated rules, regeneration, and validation. |

The exact command list inside each skill must be based on current source/help output during implementation. If a command name in this plan is not available in the current codebase, the implementation must not document it as available; record the mismatch in evidence and update the skill wording accordingly.

Command names must match the surface that implements them. CLI commands use hyphenated names such as `query-health`, `resolution-inventory`, and `source-site-accuracy`. MCP tools use underscore names such as `route_map`, `tool_map`, `shape_check`, and `api_impact`. A skill may mention both surfaces, but it must not invent a CLI command just because an MCP tool exists.

The package/editor setup skill source must be reconciled with the embedded AI-context skill source. The final implementation should make it hard for package-root `skills/`, embedded `internal/aicontext/skills/*.md`, and generated `.claude/skills/avmatrix/**` to drift away from each other.

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

The query reliability bug must be fixed with a measured retrieval path, not only worked around in skills. The implementation should first reproduce the miss in `query-health`, inspect the current query implementation and scoring reasons, then improve retrieval/ranking so expected owner files and symbols appear in the top results. A result with weak or no lexical/semantic overlap should not outrank exact path, symbol, package, generated-artifact, command-name, or resource-name matches. Query output should expose enough match reason evidence for agents to understand why a result was ranked.

The query fix must preserve and expand the original discovery value of `query`. It must not turn `query` into a pure grep command, a `context` alias, an exact-symbol-only lookup, or a tool tuned only for the AI-context plan case. Execution-flow and process results remain important, but they must be ranked behind stronger owner evidence when the process has weak overlap with the user intent. The accepted behavior change is structured broad discovery with lower wrong-owner noise, not loss of broad concept discovery.

The final query behavior must separate four roles clearly: candidate discovery through `query`, exact owner inspection through `context` or source inspection, retrieval-quality measurement through `query-health`, and validation evidence through the benchmark/evidence ledgers. Query-health output must distinguish usable retrieval from exact coverage: a usable pass means enough correct owner evidence appears to guide the agent, while an exact pass means no expected target was missed. Query-health must also make result ordering meaningful across sources; if process-symbol and definition ranks are merged, the report must define or expose a global rank/source rank so hit@5 and hit@10 are not ambiguous.

The target behavior for this plan's AI-context query case is one benchmark lane, not the whole definition of `query`: broad queries about generated AVmatrix skills and AI context must surface `internal/aicontext/aicontext.go`, embedded skill source files, analyze post-run AI context generation, setup/editor skill installation, package skill distribution, and MCP setup/resource guidance where relevant. If an exact target cannot be found, the query-health report must make the miss explicit through exact pass/fail, matched targets, missed targets, and noise reason.

The query result evidence schema should remain backward compatible while adding auditable fields. Prefer optional fields such as `queryLane`, `matchedFields`, `matchReasons`, `scoreClass`, `sourceRank`, `globalRank`, and `noiseReason` where appropriate. Do not require verbose raw scoring internals in normal output, but JSON/detail output must be strong enough for query-health and agents to explain why a result ranked.

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
- Full build, focused tests, generation smoke, setup/package validation if touched, and `detect-changes` pass before closure.

## Phase 0 - Generator Source Trace And Command Inventory

- [ ] [P0-A] Trace the generator ownership for `.claude/skills/avmatrix/**` and record the result in the evidence ledger. The trace must identify the source skill files, the embedded filesystem owner, the `baseSkills` registry, `baseSkillContent`, `installBaseSkills`, `GenerateAIContextFiles`, the generated root Skills table, the analyze post-run caller, and any setup/package paths that copy installed skills into editor-specific directories.

- [ ] [P0-B] Inventory the current source and generated skill set before implementation. Record each skill id, source file path, generated output path, source byte/line count, generated byte/line count, top headings, and whether the generated output matches the embedded source. This inventory must prove whether `.claude/skills/avmatrix/**` is source or generated validation output.

- [ ] [P0-C] Inventory the current AVmatrix command surface from code/tests/help output before writing skill content. Record actual available CLI commands, MCP tools, MCP resources, setup/package commands, runtime commands, group/cross-repo commands, API-surface commands, and graph-quality/accuracy commands. The evidence must distinguish implemented commands from planned or absent commands so skills do not document non-existent behavior as real.

- [ ] [P0-D] Build the skill routing matrix from the command inventory. Map every current command/tool/resource family to one primary skill and any secondary skill references, then record the final decision in evidence. The matrix must include the existing six skills and the proposed new skills for graph quality, API surface, cross-repo work, runtime/packaging, and AI context generation.

- [ ] [P0-E] Compare the command surface exposed by the current `PATH` binary, `go run .\cmd\avmatrix --help`, and the binary produced by the full build gate. Record mismatches in evidence and use only the current source or freshly built local binary as the source of truth for skill content.

- [ ] [P0-F] Trace the package/editor skill source path from `setupInstallSkillsTo` and package lifecycle code. Record whether package-root `skills/` exists, how packaged installs are expected to contain skills, and what code/test change is needed so packaged/editor skills cannot drift from embedded AI-context skills.

- [ ] [P0-G] Inspect MCP resource/setup guidance such as `avmatrix://setup` and the source that renders it. Record whether it already matches the final command taxonomy or must be updated alongside `internal/aicontext/aicontext.go`.

- [ ] [P0-H] Audit embedded base skill source validity before editing. Record each source file's frontmatter `name`, `description`, non-empty body status, and whether `baseSkillContent` would read real content or fall back.

- [ ] [P0-I] Record the broad-intent query reliability bug as baseline evidence. Run and record query results for AI context skill generation, setup/editor skill installation, package skill distribution, and MCP setup/resource guidance; then verify the correct owner surfaces with `context` and exact file inspection. The evidence must show when query output is only a candidate, when it misses expected owner files, and why that miss is dangerous for agent edit-surface selection.

- [ ] [P0-J] Create a dedicated query-health suite for this plan, for example `docs/query-health/2026-05-23-avmatrix-skill-system-upgrade-suite.json`, and add a case that reproduces the AI-context skill-generation query miss before any retrieval fix. The case must list expected files/symbols, actual top results, hit@5/hit@10, exact target coverage, missed targets, and noise reason so the bug is measurable without mixing this plan into an older suite with a different purpose.

- [ ] [P0-K] Trace the current query implementation and scoring path before changing it. Record the exact files and symbols that rank definitions, process symbols, execution flows, lexical tokens, path/name matches, semantic fields, filters, and result limits. The trace must explain why unrelated launcher/resolution/frontend flows outrank the expected AI-context owners.

- [ ] [P0-L] Inventory the current user-facing query surfaces before changing them. Record CLI help/output for `avmatrix query`, MCP `query` tool schema/output, `query-health` output, and any existing Web/API query/search surface. The evidence must state which surfaces users and agents can actually invoke today and which query-lane evidence is currently hidden or absent.

## Phase 1 - Core Query Reliability Repair

This is a large blocking phase. The rest of the skill-system upgrade depends on agents being able to use `query` without being sent to plausible but wrong code regions. This phase must treat broad-intent query misses as a product reliability bug, not as a documentation caveat. Skill text may still teach verification discipline, but the product fix must improve query itself.

- [ ] [P1-A] Build the query reliability baseline for this plan before fixing ranking. The baseline must run broad-intent queries for AI context skill generation, setup/editor skill installation, package skill distribution, MCP setup/resource guidance, and command-surface discovery. For each query, record expected owner files/symbols, actual top results, unrelated top results before first expected owner, hit@5, hit@10, exact matched targets, missed targets, and noise reason in the evidence and benchmark ledgers.

- [ ] [P1-B] Root-cause the current `query` retrieval and ranking pipeline. Trace the exact implementation path for CLI/MCP query from input text to returned definitions/processes, including tokenization, lexical matching, process scoring, definition scoring, path/name scoring, App Layer/Functional Area boosts, docs/test filtering, result limits, and result diversification. The evidence must explain why launcher/resolution/frontend flows can outrank `internal/aicontext` owners for the AI-context intent.

- [ ] [P1-C] Define the Query Capability Taxonomy and relevance contract for broad-intent repository work. The contract must describe `query` as an umbrella command with retrieval lanes such as owner discovery, concept discovery, execution-flow discovery, API surface discovery, graph-quality discovery, docs/setup/AI-context discovery, command-surface discovery, and cross-repo discovery. It must also define that `context` performs exact symbol inspection after a candidate owner is found. Exact file path, exact symbol name, generated artifact name, command name, resource URI/name, package/module name, and high lexical overlap are strong evidence; unrelated execution-flow names with weak overlap must not outrank those owners. The contract must also define how docs/tests/examples are ranked when the query is about product source versus documentation.

- [ ] [P1-D] Implement the query retrieval/ranking fix through the capability taxonomy rather than through an AI-context-only patch. The implementation must improve owner discovery using proven signals such as exact symbol/file/path matches, command/resource names, generated artifact names, package/module names, lexical token overlap, and process/definition relevance. It must preserve concept and execution-flow discovery, demote unrelated process flows that lack meaningful overlap with the query intent, and keep useful process/API/graph-quality/docs-command results when they have clear evidence. The implementation must keep broad discovery behavior intact and must not replace `query` with grep-only matching or exact-symbol-only lookup.

- [ ] [P1-E] Add result reason output where needed so query results are auditable. Query or query-health output must expose enough match evidence for an agent/user to see why a result ranked highly, such as query lane, matched tokens, matched path/symbol/command/resource fields, score class, source rank, global rank, and noise reason for expected misses. This must not turn into verbose raw internals by default; use summary fields or JSON detail where appropriate.

- [ ] [P1-F] Add the user-facing query lane command surface. Preserve the existing `avmatrix query "<intent>" --repo <repo>` command, then add clear lane discovery and explainable query usage through subcommands or flags such as `avmatrix query lanes`, `avmatrix query explain "<intent>" --repo <repo> --json`, `--lane <lane>`, or an equivalent source-verified design. Help output must describe each lane and show how a user or agent can get lane/rank/match evidence.

- [ ] [P1-G] Add focused unit tests for query scoring, lane assignment, and filtering. Tests must cover exact owner file/symbol/path matches, generated artifact names such as `AGENTS.md` and `.claude/skills/avmatrix`, command/resource-name matches such as `setupResource` and `query-health`, execution-flow results with real overlap, API/graph-quality/docs-command lane examples where supported by current data, and a negative case where unrelated launcher/resolution/frontend flows must not outrank stronger owner evidence.

- [ ] [P1-H] Add or update CLI/MCP query and query-health integration tests. Tests must prove the AI-context intent returns expected owner surfaces, representative non-AI-context query lanes still return useful results, threshold pass is reported separately from exact pass, missed-target reporting works, source/global ranking semantics are clear, CLI lane/explain commands are invokable, MCP output exposes equivalent machine-readable evidence, and unrelated high-scoring results cannot silently be accepted as a passing exact result.

- [ ] [P1-I] Run the updated query-health suite after the query fix and record threshold pass/fail, exact pass/fail, expected targets, matched targets, missed targets, query lane coverage, unrelated top-result count, source/global rank behavior, user-facing command outputs, and remaining noise reason in the benchmark ledger.

- [ ] [P1-J] Finalize skill-guidance requirements only after the query behavior and command surface are measured. Do not edit embedded skill content in this phase. Record the exact guidance that Phase 2 must apply to `avmatrix-exploring`, `avmatrix-debugging`, `avmatrix-graph-quality`, and `avmatrix-ai-context`: `query` is the right broad discovery command with multiple usable capability lanes, and broad results still need verification with `context` or exact source inspection when selecting edit surfaces.

- [ ] [P1-K] Add a closure gate for query reliability before moving to the rest of the skill-system implementation. The gate is satisfied only when the AI-context query-health case has recorded threshold and exact results, the remaining missed targets if any are explicit, the usable CLI/MCP query lane surfaces are validated, and the plan/evidence/benchmark ledgers state whether the query bug is fixed or still has a tracked blocker.

- [ ] [P1-L] Add broad-discovery regression checks for `query` after the ranking fix. Use representative non-AI-context intents across the capability taxonomy, including at least one intent that should naturally return execution-flow or process candidates. Record the before/after top results and prove the fix reduced wrong-owner noise without removing the original concept-to-flow discovery capability.

## Phase 2 - Embedded Skill Source Upgrade

- [ ] [P2-A] Upgrade the six existing embedded skill Markdown files in `internal/aicontext/skills/`. Each file must become a practical task guide with command choices, when to use each AVmatrix surface, validation expectations, and current limitations. `avmatrix-impact-analysis` must explain that HIGH/CRITICAL is blast-radius evidence to report and account for, not a blanket prohibition against required work.

- [ ] [P2-B] Add the new embedded source skill files under `internal/aicontext/skills/`: `avmatrix-graph-quality.md`, `avmatrix-api-surface.md`, `avmatrix-cross-repo.md`, `avmatrix-runtime-packaging.md`, and `avmatrix-ai-context.md`. Each new skill must contain concrete usage guidance, command examples based on implemented commands, expected outputs, and validation notes.

- [ ] [P2-C] Update the base skill registry and generated Skills table in `internal/aicontext/aicontext.go`. The registry and generated table must include all final skills, use repo-agnostic descriptions, and avoid splitting AVmatrix into misleading MCP-only versus CLI-only capability lists.

- [ ] [P2-D] Add or update `internal/aicontext` tests so generated root files and generated base skills are protected. Tests must assert the final skill ids, generated `.claude/skills/avmatrix/<skill>/SKILL.md` paths, generated Skills table links, and representative key phrases for the new command surfaces.

- [ ] [P2-E] Add coverage tests that prevent the guide from regressing back to a six-skill or analyze/query/impact-only view. The tests should check the generated guidance for the AI context skill, graph-quality skill, API-surface skill, cross-repo skill, runtime/packaging skill, and a current command-surface fragment confirmed in Phase 0.

- [ ] [P2-F] Add or update tests that validate command naming by surface. The test should protect at least one CLI-only hyphenated command such as `query-health`, one MCP underscore tool such as `route_map`, and one statement that does not invent a CLI spelling for an MCP-only tool.

- [ ] [P2-G] Add frontmatter/source-content tests for every final embedded base skill. The test must fail if a registered base skill is missing its embedded Markdown file, has empty content, has mismatched `name`, lacks `description`, or would rely on `fallbackBaseSkillContent`.

- [ ] [P2-H] Update `avmatrix-exploring`, `avmatrix-debugging`, `avmatrix-graph-quality`, and `avmatrix-ai-context` skill content using the measured Phase 1 guidance. The guidance must state that broad intent query is a multi-lane candidate discovery command, explain the user-facing CLI/MCP query lane commands or flags, state that `context` is preferred when an exact symbol is known, and tell agents that noisy/missed query results must be recorded as graph-quality/query-health evidence.

- [ ] [P2-I] Add or extend skill-facing query-health guidance for this plan's AI-context intent. The skill should point to the query-health suite/case created in Phase 0/Phase 1, explain threshold versus exact pass, and tell agents to record missed targets instead of treating partial retrieval as complete.

## Phase 3 - Setup, Package, And Documentation Integration

- [ ] [P3-A] Verify the analyze post-run path installs the expanded base skill set through the same normal generation path that creates `AGENTS.md` and `CLAUDE.md`. If tests currently assert the old six-skill count or specific old table rows, update them to assert the new final set.

- [ ] [P3-B] Verify setup/editor installation behavior for the expanded embedded skill set. Inspect and test `setupInstallEditorSkills` and related setup command behavior so supported editor skill directories receive the same final skill content without relying on generated repository-local `.claude/skills/avmatrix/**` as source.

- [ ] [P3-C] Verify package/runtime distribution behavior for the expanded embedded skill set. If packaging tests or package assembly code enumerate skills, update them so the packaged tool can generate and install the final skill set from embedded source files.

- [ ] [P3-D] Reconcile package-root `skills/` with embedded `internal/aicontext/skills/*.md`. Either make the package/setup path materialize or copy from the same canonical skill source, or document and test a deliberately equivalent packaged `skills/` directory. The output must prove `avmatrix setup` installs the same final base skill ids and content family as `avmatrix analyze --force` generates in `.claude/skills/avmatrix/**`.

- [ ] [P3-E] Update MCP setup/resource guidance if Phase 0 finds stale command/tool/resource text. This includes the source that renders `avmatrix://setup`, MCP tool reference tables, and any setup onboarding text used by agents.

- [ ] [P3-F] Update README and relevant user-facing docs that describe AVmatrix skills, AI context generation, setup, or usage. The docs must tell users that `AGENTS.md`, `CLAUDE.md`, and `.claude/skills/avmatrix/**` are generated by AVmatrix, and that source changes belong in the embedded skill source and generator code.

- [ ] [P3-G] Search the active documentation for stale six-skill-only guidance, stale package-root skill assumptions, or stale wording that treats MCP and CLI as separate incomplete command lists. Update current guides and README-style docs; leave historical ledgers untouched unless they are actively reused as user guidance.

## Phase 4 - Regeneration And Validation

- [ ] [P4-A] Run the full build gate before tests: `powershell -ExecutionPolicy Bypass -File avmatrix-launcher\build.ps1`. Record the command result in evidence.

- [ ] [P4-B] Run focused backend/CLI/MCP tests for AI context generation, setup, packaging, query, and query-health surfaces touched by the implementation. The minimum expected test scope is `go test .\internal\mcp .\internal\cli .\internal\aicontext -count=1`, expanded as needed if package/setup code outside those packages changes.

- [ ] [P4-C] Run the normal generation path with `avmatrix analyze --force` and no `--skip-agents-md`. Verify generated `AGENTS.md`, generated `CLAUDE.md`, and `.claude/skills/avmatrix/**` contain the final skill set and expected content fragments.

- [ ] [P4-D] Compare source and generated skill inventories after regeneration. Record final skill count, generated file paths, byte/line counts, and any intentional generated differences in the benchmark ledger.

- [ ] [P4-E] Validate setup/package behavior if Phase 3 changed setup or package code. Record the exact command outputs and installed/packaged skill file inventories in evidence and benchmark ledgers.

- [ ] [P4-F] Validate MCP setup/resource output if Phase 3 touched MCP resources. Record the exact `avmatrix://setup` or equivalent resource output check in evidence so the generated guidance and MCP-facing guide are proven consistent.

- [ ] [P4-G] Run the dedicated query-health suite updated for this plan's AI-context skill-generation intent and representative query capability lanes. Record threshold pass/fail, exact pass/fail, expected targets, matched targets, missed targets, source/global rank behavior, query lane coverage, and noise reason in benchmark/evidence.

- [ ] [P4-H] Run user-facing query lane smoke commands after the full build. Validate normal query behavior, lane discovery, explainable JSON output, MCP query JSON evidence, and the dedicated query-health suite command; record exact commands and representative output in evidence and benchmark.

- [ ] [P4-I] Run the query broad-discovery regression check from Phase 1 and record whether execution-flow candidates still appear with meaningful match evidence after the ranking fix.

- [ ] [P4-J] Run `detect-changes` before commit and record the affected scope. Commit the implementation slice after checklist items and ledgers are updated.

## Phase 5 - Zero-Trust Closure Review

- [ ] [P5-A] Review the codebase and documentation for old assumptions after implementation. Search for old six-skill-only tables, old `avmatrix-cli` descriptions that omit current command families, stale generated-output instructions, and any direct-edit guidance for `.claude/skills/avmatrix/**`; fix active docs that would mislead users or agents.

- [ ] [P5-B] Re-run the final validation commands required by this plan after the closure review changes. Record final pass/fail counts, generated inventory, setup/package inventory if applicable, and any remaining limitation in the evidence and benchmark ledgers.

- [ ] [P5-C] Mark the plan complete only after source files, generated validation output, tests, docs, benchmark ledger, evidence ledger, and commit state all agree on the final skill set and the AI-context query-health case has recorded threshold and exact results.
