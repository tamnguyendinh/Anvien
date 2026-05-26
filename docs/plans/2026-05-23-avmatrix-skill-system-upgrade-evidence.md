# AVmatrix Skill System Upgrade Evidence Ledger

Date: 2026-05-23

Status: Active

Companion files:

- Plan: [2026-05-23-avmatrix-skill-system-upgrade-plan.md](2026-05-23-avmatrix-skill-system-upgrade-plan.md)
- Benchmark ledger: [2026-05-23-avmatrix-skill-system-upgrade-benchmark.md](2026-05-23-avmatrix-skill-system-upgrade-benchmark.md)

## Evidence Rules

Record evidence as each evidenced task is completed. Evidence should include commands, impacted files, test results, smoke artifacts, generated output inventory, and concise observations needed to audit the plan later.

For doc-only commits, do not use AVmatrix.

Do not record inferred generated-file behavior. Every behavior claim must include source inspection, test output, generation smoke output, setup/package output, or exact file measurement.

## E0 - Plan Creation Evidence

Date: 2026-05-23

Status: recorded

Created file set:

- `docs/plans/2026-05-23-avmatrix-skill-system-upgrade-plan.md`
- `docs/plans/2026-05-23-avmatrix-skill-system-upgrade-evidence.md`
- `docs/plans/2026-05-23-avmatrix-skill-system-upgrade-benchmark.md`

Plan creation scope:

- Identify the files responsible for generating `.claude/skills/avmatrix/**`.
- Plan a new and upgraded skill set.
- Plan source code, generated context, docs, setup, package, and validation updates.
- Keep generated `.claude/skills/avmatrix/**`, root `AGENTS.md`, and root `CLAUDE.md` as validation output rather than source files.

Convention inspection command:

```powershell
Get-ChildItem .\docs\plans -Filter "*.md" | Sort-Object LastWriteTime -Descending | Select-Object -First 12 Name,LastWriteTime
```

Observed planning convention:

- Plan files use a plan/evidence/benchmark trio.
- The plan file carries rules, problem, scope, design decisions, acceptance criteria, phases, and concrete checklist items.
- The evidence ledger records commands, observations, validation artifacts, and generated output facts.
- The benchmark ledger records measured inventories and command-output metrics separately from narrative evidence.

## E1 - Initial Source Trace Evidence

Date: 2026-05-23

Status: preliminary; implementation must re-verify before code edits

Command:

```powershell
rg -n "baseSkills|baseSkillContent|GenerateAIContextFiles|GenerateSkillFiles|go:embed skills|\.claude\\skills\\avmatrix|\.claude/skills/avmatrix|setupInstallEditorSkills|package.*skills" internal cmd avmatrix README.md docs -g "*.go" -g "*.md"
```

Observed source owners:

| Area | Observed path | Responsibility |
|---|---|---|
| Embedded skill source | `internal/aicontext/skills/*.md` | Source Markdown for packaged AVmatrix skills. |
| Embedded filesystem | `internal/aicontext/aicontext.go` | `//go:embed skills/*.md` embeds source skill Markdown. |
| Base skill registry | `internal/aicontext/aicontext.go` | `baseSkills` controls installed base skill ids/descriptions. |
| Base skill content loader | `internal/aicontext/aicontext.go` | `baseSkillContent` reads embedded skill Markdown. |
| Base skill installer | `internal/aicontext/aicontext.go` | `installBaseSkills` writes `.claude/skills/avmatrix/<skill>/SKILL.md`. |
| Root AI context generator | `internal/aicontext/aicontext.go` | `GenerateAIContextFiles` creates or updates root `AGENTS.md`, `CLAUDE.md`, and base skills. |
| Generated community skills | `internal/aicontext/aicontext.go` | `GenerateSkillFiles` writes `.claude/skills/generated/**`, separate from base AVmatrix skills. |
| Analyze post-run bridge | `internal/cli/analyze_postrun.go` | Calls AI context generation after analyze. |
| Editor setup skill copy | `internal/cli/setup_command.go` | `setupInstallEditorSkills` copies packaged skills into editor skill directories. |
| CLI tests | `internal/cli/command_test.go` | Contains generated output and package/setup assertions that may need update. |
| AI context tests | `internal/aicontext/aicontext_test.go`, `internal/aicontext/skill_gen_test.go` | Existing coverage for root context and generated skill behavior. |

Initial conclusion:

- `.claude/skills/avmatrix/**` is generated output.
- The implementation must update `internal/aicontext/skills/*.md`, `internal/aicontext/aicontext.go`, and tests/docs that rely on the old skill set.

## E2 - Current Generated Skill Gap Audit

Date: 2026-05-23

Status: preliminary; implementation must re-verify exact content before edits

Current generated skill ids observed from the generated Skills table and local generated directory:

- `avmatrix-exploring`
- `avmatrix-impact-analysis`
- `avmatrix-debugging`
- `avmatrix-refactoring`
- `avmatrix-guide`
- `avmatrix-cli`

Observed gaps to verify and fix:

| Skill area | Gap |
|---|---|
| `avmatrix-cli` | Describes a small subset of CLI usage and does not clearly cover runtime, setup, package, group, wiki, hook, version, benchmark, or accuracy command families. |
| `avmatrix-guide` | Separates MCP tools and CLI fallback in a way that can make agents treat them as separate incomplete systems instead of AVmatrix command surfaces. |
| `avmatrix-exploring` | Needs current guidance for execution flows, resources, query/context usage, App Layer/Functional Area metadata, and when to use more specific skills. |
| `avmatrix-impact-analysis` | Needs current guidance that HIGH/CRITICAL risk is blast-radius evidence to report and account for, not a blanket prohibition against required work. |
| `avmatrix-debugging` | Needs current graph-health, resolution-health, source-site, diagnostics, runtime evidence, and query-health guidance. |
| `avmatrix-refactoring` | Needs current rename/impact/detect-changes/API contract/source-site guidance and no find-and-replace symbol rename behavior. |
| Missing graph quality skill | Query health, source-site inventory, resolution inventory, resolved-edge precision, and benchmark comparison need a dedicated skill. |
| Missing API surface skill | API route maps, MCP tool maps, contract shape checks, API impact, generated contracts, handlers, and consumers need a dedicated skill. |
| Missing cross-repo skill | Group repositories, cross-repo query/contracts/status/sync, and multi-repo guidance need a dedicated skill. |
| Missing runtime/packaging skill | `serve`, `mcp`, `setup`, launcher, packaged runtime, package preparation, runtime cleanup, and startup validation need a dedicated skill. |
| Missing AI context skill | Generated `AGENTS.md`, `CLAUDE.md`, `.claude/skills/avmatrix/**`, source-vs-generated rules, regeneration, and validation need a dedicated skill. |

## E3 - Target Skill Taxonomy

Date: 2026-05-23

Status: planned

Target final skill set:

| Skill | Action | Primary command/resource coverage |
|---|---|---|
| `avmatrix-exploring` | Upgrade existing | `query`, `context`, resources, process/resource exploration, architecture navigation. |
| `avmatrix-impact-analysis` | Upgrade existing | `impact`, `detect-changes`, changed-scope review, blast-radius reporting. |
| `avmatrix-debugging` | Upgrade existing | `query`, `context`, diagnostics, graph health, runtime evidence, resolution/source-site facts. |
| `avmatrix-refactoring` | Upgrade existing | `rename`, `impact`, `context`, `detect-changes`, refactor validation. |
| `avmatrix-guide` | Upgrade existing | Unified MCP/CLI/resource/Web/API command surface and graph schema reference. |
| `avmatrix-cli` | Upgrade existing | Full CLI command guide based on current source/help output. |
| `avmatrix-graph-quality` | Add new | Query-health, resolution/source-site inventory, edge accuracy, benchmark comparison. |
| `avmatrix-api-surface` | Add new | Route/tool map, shape check, API impact, contracts, handlers, consumers. |
| `avmatrix-cross-repo` | Add new | Groups, multi-repo query/contracts/status/sync, cross-repo impact context. |
| `avmatrix-runtime-packaging` | Add new | `serve`, `mcp`, `setup`, launcher, package/runtime flows. |
| `avmatrix-ai-context` | Add new | AI context generation, embedded skills, generated output validation. |

Implementation note:

- The exact command list must be rechecked from code/help before skill content is written.
- A command named in discussion but absent from current source must be recorded as absent or future-facing, not documented as working behavior.

## E4 - Implementation Evidence

Date: pending

Status: pending

Record here:

- AVmatrix analyze/impact evidence for implementation slices.
- Edited source files.
- Generated output smoke commands.
- Test commands and pass/fail counts.
- Setup/package smoke output when applicable.
- `detect-changes` output before commit.
- Commit hashes.

## E5 - Codebase Review Before Implementation

Date: 2026-05-23

Status: recorded

Fresh graph command:

```powershell
avmatrix analyze --force
```

Result:

```text
analyzed E:\AVmatrix-GO
files: scanned=762 parsed=569 unsupported=193 failed=0
graph: nodes=24204 relationships=60607 path=E:\AVmatrix-GO\.avmatrix\graph.json
```

AVmatrix context commands used:

```powershell
avmatrix context GenerateAIContextFiles --repo AVmatrix
avmatrix context installBaseSkills --repo AVmatrix
avmatrix context setupInstallSkillsTo --repo AVmatrix
avmatrix context newRootCommand --repo AVmatrix
avmatrix context newPackageCommand --repo AVmatrix
avmatrix context newQueryHealthCommand --repo AVmatrix
avmatrix context newResolutionInventoryCommand --repo AVmatrix
avmatrix context newSourceSiteAccuracyCommand --repo AVmatrix
```

Confirmed code facts:

| Symbol | File | Finding |
|---|---|---|
| `GenerateAIContextFiles` | `internal/aicontext/aicontext.go` | Generates root `AGENTS.md`, root `CLAUDE.md`, and calls `installBaseSkills`. Incoming callers include `Generate`, `generateAnalyzeAIContext`, and AI context tests. |
| `installBaseSkills` | `internal/aicontext/aicontext.go` | Writes `.claude/skills/avmatrix/<skill>/SKILL.md` from embedded base skill content. |
| `setupInstallSkillsTo` | `internal/cli/setup_command.go` | Installs editor skills by reading package-root `skills/`, then copying flat `.md` or directory `SKILL.md` entries to editor skill directories. |
| `NewRootCommand` | `internal/cli/command.go` | Registers current CLI commands including package, group, query-health, resolution-inventory, and source-site-accuracy from source. |
| `newPackageCommand` | `internal/cli/package_command.go` | Owns package lifecycle subcommands. |
| `newQueryHealthCommand` | `internal/cli/query_health_command.go` | Source contains query-health command registration. |
| `newResolutionInventoryCommand` | `internal/cli/resolution_inventory_command.go` | Source contains resolution-inventory command registration. |
| `newSourceSiteAccuracyCommand` | `internal/cli/source_site_accuracy_command.go` | Source contains source-site-accuracy command registration. |

Command-surface mismatch discovered:

```powershell
avmatrix --help
avmatrix query-health --help
avmatrix resolution-inventory --help
avmatrix source-site-accuracy --help
go run .\cmd\avmatrix --help
go run .\cmd\avmatrix query-health --help
go run .\cmd\avmatrix resolution-inventory --help
go run .\cmd\avmatrix source-site-accuracy --help
```

Observed behavior:

- `avmatrix` from `PATH` did not list `query-health`, `resolution-inventory`, or `source-site-accuracy`.
- `go run .\cmd\avmatrix --help` from the current source did list `query-health`, `resolution-inventory`, and `source-site-accuracy`.
- Therefore command inventory for this plan must use current source or the freshly built local binary, not an older binary found in `PATH`.

Package/editor skill source finding:

- `setupInstallSkillsTo` reads from `setupResolvePackagePath("skills")`.
- Repository-local generation reads embedded files from `internal/aicontext/skills/*.md`.
- Initial filesystem inspection did not show a root-level `skills/` directory in the working tree.
- Phase 0, Phase 3, and Phase 4 must reconcile package/editor skill installation with embedded AI-context skill source so the packaged setup path does not drift from generated repository-local skills.

MCP/resource guidance finding:

- `internal/mcp/resources.go` contains setup/resource/tool guidance including MCP tool tables and setup reference output.
- This file must be part of the plan's stale-guidance search because updating only `internal/aicontext/aicontext.go` would leave another user-facing command guide that can drift.

Query reliability bug finding:

- Broad intent query was not reliable enough to identify the correct owner region for this plan.
- Query intents around AI context skill generation, setup/editor skills, package skill distribution, and command surface returned plausible but unrelated launcher, resolution-gap, and frontend/backend flows instead of consistently surfacing `internal/aicontext/aicontext.go`, `internal/aicontext/skills/*.md`, `internal/cli/analyze_postrun.go`, `internal/cli/setup_command.go`, package lifecycle code, and `internal/mcp/resources.go`.
- Symbol-level `context` calls on `GenerateAIContextFiles`, `installBaseSkills`, `setupInstallSkillsTo`, `NewRootCommand`, `newPackageCommand`, `newQueryHealthCommand`, `newResolutionInventoryCommand`, `newSourceSiteAccuracyCommand`, and `setupResource` did locate the correct owner surfaces.
- Classification: this is a core `query` feature reliability bug, not only a documentation issue. A query that cannot identify the target region can send an agent to edit or reason about the wrong code.
- Conclusion: broad `query` output must be treated as candidate retrieval until the bug is fixed. The implementation plan must reproduce the miss in query-health, root-cause ranking/scoring, fix retrieval where possible, and keep exact missed-target reporting visible even when threshold retrieval passes.

Query behavior analysis added to the plan:

- `query` is intended to be concept-to-code and concept-to-flow discovery. It should help an agent find likely work areas from broad intent.
- `query` is not meant to replace `context`. `context` remains the exact symbol inspection tool once a candidate symbol or owner file is known.
- The planned fix must improve retrieval/ranking and auditability without reducing `query` to grep-only matching or exact-symbol-only lookup.
- Execution-flow and process results remain valuable, but they must not outrank stronger owner evidence when they have weak overlap with the query intent.
- Query-health must separate usable retrieval from exact coverage. A usable pass means enough correct owner evidence exists to guide work; an exact pass means no expected target was missed.
- The plan now requires a broad-discovery regression check so the query reliability repair cannot accidentally remove the original broad concept discovery purpose.

Query umbrella-command correction:

- The plan must not define `query` through a narrow set of implementation symbols or through the AI-context bug case alone.
- `query` is a top-level command family. The plan now treats its behavior as multiple retrieval lanes under one umbrella command: owner discovery, concept discovery, execution-flow discovery, API surface discovery, graph-quality discovery, docs/setup/AI-context discovery, command-surface discovery, and cross-repo discovery when indexed data supports it.
- The AI-context skill-generation case is one benchmark lane. It is not the product definition of `query`.
- Implementation work must preserve broad query behavior while adding stronger structure, lane evidence, match reasons, and clearer ranking.
- Query-health evidence must record whether hit@5/hit@10 use a global rank or source-specific rank so users and agents can interpret the benchmark correctly.

Usability requirement added:

- A query capability is not complete if it only exists as hidden scoring code, unrendered fields, or internal tests.
- Query lanes must be discoverable and usable through AVmatrix command surfaces. The plan now requires CLI lane discovery, explainable query JSON output, normal `avmatrix query` compatibility, and MCP query output with machine-readable lane/rank/match evidence.
- If an existing Web/API query/search surface consumes query results, it must display or pass through the new evidence. If no Web UI surface is in scope, the evidence ledger must record that CLI/MCP/API output is the usable product surface for this plan.
- Validation must run actual commands or focused MCP tests that prove users and agents can invoke the feature without reading internal code.

## E6 - Graph Labeling Problem Evidence

Date: 2026-05-23

Status: recorded for planning; implementation must re-verify before code edits

Problem screenshot:

| Artifact | Size | Finding |
|---|---:|---|
| `reports/problem/screenshot_1779517751.png` | 341,738 bytes | The graph shows visible rings/islands, but the canvas does not directly name each macro ring or each island. |

Observed UI issue:

- The graph has visual grouping, but users still need to infer the meaning of groups from color, side-panel filters, memory, or hover behavior.
- Node islands are also too visually compressed in the reported graph. Nodes can appear too close together, island spiral bands do not leave enough breathing room, and neighboring islands/rings can read as dense masses instead of separate groups.
- The center of each macro ring should carry a readable name such as Backend, Frontend, API, Docs, Config, Shared, Test, Unknown, or the graph's current equivalent.
- Each visible node island should carry a readable label above or near the island, such as Function, Method, File, Route, ResolutionGap, External Reference, or the graph's current group label.
- Island radius and ring/orbit spacing should expand with visible node count so users can see individual nodes at near zoom and still understand the separated ring/island structure when zoomed out.
- Color and the left dashboard are not enough. The graph itself must communicate what the visible structure represents and leave enough whitespace for users to read that structure.

Preliminary Web source trace:

```powershell
Get-Content -Path avmatrix-web/package.json
rg -n "export .*GraphCanvas|function GraphCanvas|const GraphCanvas|useSigma\(|export const useSigma|layoutRings|buildGraph|adaptGraph|nodeLabel|appLayer|ring|island" avmatrix-web/src/components/GraphCanvas.tsx avmatrix-web/src/hooks/useSigma.ts avmatrix-web/src/lib/graph-adapter.ts avmatrix-web/src/lib/constants.ts avmatrix-web/e2e/server-connect.spec.ts avmatrix-web/test/unit/graph-adapter.edge-geometry.test.ts
```

Observed candidate owners:

| Area | Path | Finding |
|---|---|---|
| Graph canvas and runtime diagnostics | `avmatrix-web/src/components/GraphCanvas.tsx` | Already computes app-layer ring diagnostics, ring centers, island counts, and render ownership around `GraphCanvas`. |
| Sigma rendering hook | `avmatrix-web/src/hooks/useSigma.ts` | Owns Sigma lifecycle and custom rendering behavior that may need overlay/label integration. |
| Graph layout adapter | `avmatrix-web/src/lib/graph-adapter.ts` | Assigns `appLayerRing`, `islandKey`, ring centers, island placement, and node attributes used by the current layout. |
| Display constants | `avmatrix-web/src/lib/constants.ts` | Owns node colors, display labels, filterable labels, documentation display label, and relationship display helpers. |
| Unit geometry tests | `avmatrix-web/test/unit/graph-adapter.edge-geometry.test.ts` | Already tests ring/island geometry and is the likely place for deterministic label metadata/anchor tests. |
| Browser/e2e graph diagnostics | `avmatrix-web/e2e/server-connect.spec.ts` | Already validates layout rings and node-type islands in browser and captures graph screenshots. |

Validation commands available from `avmatrix-web/package.json`:

```powershell
cd avmatrix-web
npm run test
npm run test:e2e
```

## E7 - Zero-Trust Plan Readiness Review

Date: 2026-05-23

Status: recorded

Fresh graph command:

```powershell
avmatrix analyze --force
```

Result:

```text
analyzed E:\AVmatrix-GO
files: scanned=762 parsed=569 unsupported=193 failed=0
graph: nodes=24211 relationships=60614 path=E:\AVmatrix-GO\.avmatrix\graph.json
```

AVmatrix context checks:

```powershell
avmatrix context GenerateAIContextFiles --repo AVmatrix
avmatrix context installBaseSkills --repo AVmatrix
avmatrix context setupInstallSkillsTo --repo AVmatrix
avmatrix context GraphCanvas --repo AVmatrix
avmatrix context useSigma --repo AVmatrix
```

Codebase observations from the readiness review:

| Area | Observation | Plan effect |
|---|---|---|
| AI context generation | `GenerateAIContextFiles` calls `renderAVmatrixBlock`, `upsertSection`, and `installBaseSkills`; `installBaseSkills` reads embedded source skill Markdown through `baseSkillContent`. | Existing Phase 0/3/5 coverage is sufficient. |
| Setup/editor skill installation | `setupInstallSkillsTo` reads package-root `skills/` via `setupResolvePackagePath("skills")`, while repository-local generation reads `internal/aicontext/skills/*.md`. | Existing Phase 0/4 coverage is sufficient and must remain. |
| Query CLI surface | `internal/cli/tool_command.go` currently defines `avmatrix query <search_query>` as one positional argument and forwards to MCP `query`. | Added a compatibility requirement so lane/explain syntax cannot break normal broad query usage. |
| Query ranking owner | MCP `query` is implemented in `internal/mcp/tools.go`, with `rankedProcessMatches`, `matchingDefinitionRows`, token scoring, semantic boosts, and query-health consuming the same local query output. | Existing Phase 0/1 coverage is sufficient; implementation must inspect these real owners. |
| Web graph layout | `graph-adapter.ts` assigns `appLayerRing`, `islandKey`, ring centers, and node positions; `GraphCanvas.tsx` records ring diagnostics; `useSigma.ts` owns Sigma rendering and manual layout optimization. | Added label visibility requirements tied to the currently visible filtered/depth graph. |
| Web label validation | Existing e2e diagnostics already count rings/islands but not label entities. | Added requirement for runtime diagnostics or test selectors that make label counts machine-checkable in browser tests. |

Readiness conclusion:

- The plan has the right implementation phases and owner files.
- The plan needed two execution guardrails before implementation: `query` lane syntax must preserve the current positional query behavior, and graph labels must update with the visible graph after filters/depth changes instead of using stale initial conversion metadata.

## E8 - Phase 0 Implementation Baseline

Date: 2026-05-26

Status: recorded before implementation edits

Fresh graph command:

```powershell
.\avmatrix\bin\avmatrix.exe analyze --force
```

Result:

```text
files: scanned=761 parsed=568 unsupported=193 failed=0
graph: nodes=85683 relationships=117589 path=E:\AVmatrix-GO\.avmatrix\graph.json
```

Full build gate for Phase 0 command-surface comparison:

```powershell
powershell -ExecutionPolicy Bypass -File avmatrix-launcher\build.ps1
```

Result:

```text
Go: go version go1.26.3 windows/amd64
Web build: tsc -b && vite build
Build status: pass
```

Generator ownership trace:

| Owner | Path/symbol | Phase 0 finding |
|---|---|---|
| Embedded source skill files | `internal/aicontext/skills/*.md` | Canonical source for generated repository-local AVmatrix skills. |
| Embedded filesystem | `internal/aicontext/aicontext.go` `//go:embed skills/*.md` | Embeds base skill Markdown at compile time. |
| Base skill registry | `internal/aicontext/aicontext.go` `baseSkills` | Currently registers six base skills. |
| Skill content loader | `internal/aicontext/aicontext.go` `baseSkillContent` | Reads embedded Markdown; fallback remains defensive only. |
| Skill writer | `internal/aicontext/aicontext.go` `installBaseSkills` | Writes `.claude/skills/avmatrix/<skill>/SKILL.md`. |
| Root AI context generator | `internal/aicontext/aicontext.go` `GenerateAIContextFiles` | Writes managed `AGENTS.md`, `CLAUDE.md`, and base skills. |
| Analyze caller | `internal/cli/analyze_postrun.go` `generateAnalyzeAIContext` | Calls `GenerateAIContextFiles` after analyze. |
| Setup/editor skill path | `internal/cli/setup_command.go` `setupInstallSkillsTo` | Reads `setupResolvePackagePath("skills")`; repo root currently has no package-root `skills/` directory. |
| Package lifecycle | `internal/cli/package_command.go`, `internal/cli/package_runtime.go` | Hidden package commands exist; package/setup skill source must be reconciled with embedded skill source. |

Current source/generated skill inventory:

| Skill | Source bytes/lines | Generated bytes/lines | Generated matches source | Top heading |
|---|---:|---:|---|---|
| `avmatrix-cli` | 4,873 / 91 | 4,873 / 91 | yes | `# AVmatrix CLI Commands` |
| `avmatrix-debugging` | 2,168 / 57 | 2,168 / 57 | yes | `# Debugging With AVmatrix` |
| `avmatrix-exploring` | 2,274 / 57 | 2,274 / 57 | yes | `# Exploring Codebases With AVmatrix` |
| `avmatrix-guide` | 3,967 / 83 | 3,967 / 83 | yes | `# AVmatrix Tool And Resource Guide` |
| `avmatrix-impact-analysis` | 2,283 / 64 | 2,283 / 64 | yes | `# Impact Analysis With AVmatrix` |
| `avmatrix-refactoring` | 1,934 / 54 | 1,934 / 54 | yes | `# Refactoring With AVmatrix` |

All six current source skills contain `name`, `description`, and a non-empty body. The generated `.claude/skills/avmatrix/**/SKILL.md` files are validation output, not source files.

Command-surface inventory from current source and freshly built binary:

| Surface | Current commands/tools/resources |
|---|---|
| Built CLI `.\avmatrix\bin\avmatrix.exe --help` | `analyze`, `augment`, `benchmark-compare`, `clean`, `completion`, `context`, `cypher`, `detect-changes`, `group`, `help`, `impact`, `index`, `list`, `mcp`, `query`, `query-health`, `resolution-inventory`, `serve`, `setup`, `source-site-accuracy`, `status`, `version`, `wiki`, `wiki-mode` |
| Current source `go run .\cmd\avmatrix --help` | Same visible command list as freshly built binary. |
| PATH `avmatrix --help` | Same visible command list in this Phase 0 run; earlier evidence recorded PATH staleness before the local build/path state changed. |
| Hidden CLI lifecycle commands from source | `package build-runtime`, `package prepare-go-source`, `package ensure-runtime`, `package clean-go-source`, and `hook claude`. |
| MCP tools from `internal/mcp/tools.go` | `list_repos`, `query`, `cypher`, `context`, `detect_changes`, `rename`, `impact`, `route_map`, `tool_map`, `shape_check`, `api_impact`, `group_list`, `group_sync`, `group_contracts`, `group_query`, `group_status`. |
| MCP resources from `internal/mcp/resources.go` | `avmatrix://repos`, `avmatrix://setup`, `avmatrix://repo/{name}/context`, `/clusters`, `/processes`, `/schema`, `/cluster/{clusterName}`, `/process/{processName}`. |
| MCP prompts from `internal/mcp/prompts.go` | `detect_impact`, `generate_map`; `generate_map` still needs deterministic repo resolution and freshness/evidence rules in Phase 1.7. |
| HTTP/Web query/search surfaces | `POST /api/query` for Cypher-style panel queries, `POST /api/search` for BM25/semantic/hybrid search, Web `QueryFAB` for Cypher panel queries, Web backend search calls for semantic search. |

Skill routing matrix derived from the inventory:

| Command/tool/resource family | Primary skill | Secondary skills |
|---|---|---|
| `analyze`, `status`, `list`, `index`, `clean` | `avmatrix-cli` | `avmatrix-runtime-packaging`, `avmatrix-guide` |
| `query`, `context`, process/resources | `avmatrix-exploring` | `avmatrix-debugging`, `avmatrix-guide`, `avmatrix-graph-quality` |
| `impact`, `api_impact`, `detect_changes`, `rename` | `avmatrix-impact-analysis` | `avmatrix-refactoring`, `avmatrix-api-surface` |
| `cypher`, MCP resources, graph schema | `avmatrix-guide` | `avmatrix-exploring`, `avmatrix-graph-quality` |
| `query-health`, `resolution-inventory`, `source-site-accuracy`, `benchmark-compare`, planned `graph-health` | `avmatrix-graph-quality` | `avmatrix-debugging`, `avmatrix-cli` |
| `route_map`, `tool_map`, `shape_check`, API contracts/routes | `avmatrix-api-surface` | `avmatrix-impact-analysis`, `avmatrix-guide` |
| `group` CLI and `group_*` MCP tools | `avmatrix-cross-repo` | `avmatrix-cli`, `avmatrix-guide` |
| `serve`, `mcp`, `setup`, hidden `package`, hidden `hook`, launcher runtime | `avmatrix-runtime-packaging` | `avmatrix-cli`, `avmatrix-ai-context` |
| `AGENTS.md`, `CLAUDE.md`, `.claude/skills/avmatrix/**`, embedded skills | `avmatrix-ai-context` | `avmatrix-guide`, `avmatrix-runtime-packaging` |

MCP setup/resource guidance finding:

- `internal/mcp/resources.go` currently renders `avmatrix://setup`.
- The setup resource lists tools and resources but does not yet describe MCP prompts as executable templates and does not yet reflect the final command taxonomy required by this plan.
- `avmatrix://repos` is a concrete discovery resource for indexed repositories; prompt/template guidance must instruct agents to resolve the repo from this resource instead of continuing with placeholders such as `{name}`.

Baseline query-health suite created:

```powershell
.\avmatrix\bin\avmatrix.exe query-health --repo AVmatrix --suite .\docs\query-health\2026-05-23-avmatrix-skill-system-upgrade-suite.json --limit 10
```

Summary:

```text
cases=8 thresholdPassed=5 thresholdFailed=3 exactPassed=1 exactFailed=7 matchedTargets=33/54 missedTargets=21
```

Important baseline misses:

| Case | Threshold | Exact | Key finding |
|---|---|---|---|
| `ai-context-generated-skills-owner-discovery` | fail | fail | Hit@10 found only `internal/aicontext/aicontext.go`; missed embedded skill file, analyze post-run caller, and key generator symbols. Top results were generated contracts/search/backend surfaces. |
| `setup-editor-skill-installation-owner-discovery` | fail | fail | Found setup install owner symbols but missed `setupSkillTargetName`; package/runtime process symbols also ranked high. |
| `package-skill-distribution-owner-discovery` | pass | pass | Package/setup owner discovery works for this baseline case. |
| `mcp-setup-resource-prompt-guidance-owner-discovery` | pass | fail | Usable threshold passed, but missed `promptDefinitions` and `mcpTools`. |
| `query-command-surface-owner-discovery` | pass | fail | Usable threshold passed, but missed `internal/cli/query_health_command.go` and `runQueryHealth`. |
| `graph-quality-command-surface-owner-discovery` | fail | fail | Web graph-health and MCP query internals outranked CLI graph-quality commands. |
| `api-surface-tool-discovery` | pass | fail | Missed `toolMapTool`. |
| `cross-repo-command-surface-discovery` | pass | fail | Missed CLI/MCP group command owners. |

Query implementation/scoring trace:

- CLI `avmatrix query` is defined in `internal/cli/tool_command.go` as `query <search_query>` and forwards to MCP `query`.
- MCP `query` is implemented by `Server.queryTool` in `internal/mcp/tools.go`.
- `rankedProcessMatches` scores process labels, step names, step file paths, and half-weight semantic surface boosts.
- `matchingDefinitionRows` scores name/id/file path/label/app layer/functional area/content, applies semantic surface boosts and penalties, then caps results per file.
- `querySemanticSurfaceBoost` currently has specific boost lanes for graph health, layout/front-end graph surfaces, query internals, and API/contract terms, but not for AI-context generated skills, setup/editor skill installation, package-root skills, or MCP prompt guidance.
- Current query output does not expose lane/match-reason evidence to the CLI/MCP user beyond result fields and query-health top results. Phase 1 must add a usable lane/explain surface without breaking `avmatrix query "<intent>" --repo <repo>`.

## Phase 1 Query Reliability Implementation Evidence

Status: completed for Phase 1.

Blast-radius checks before editing query and launcher/runtime code:

| Target | Result |
|---|---|
| `querySemanticSurfaceBoost` | CRITICAL; affects query ranking surface and MCP/CLI query results, handled as blast-radius warning. |
| `matchingDefinitionRows` | CRITICAL; affects query result ranking/definition rows. |
| `rankedProcessMatches` | CRITICAL; affects process ranking in query output. |
| `newQueryCommand` | CRITICAL; affects CLI query command surface. |
| `queryTool` | LOW. |
| `queryHealthActualResults`, `scoreQueryHealthCase`, `queryHealthSummaryLines` | CRITICAL; affects query-health reporting semantics. |
| `avmatrix-launcher/build.ps1` | LOW. |
| `avmatrix-launcher/src/main.go` | LOW. |
| `avmatrix-launcher/server-wrapper/main.go` | LOW. |

Implemented query changes:

- Added Query Capability Lane metadata in `internal/mcp/tools.go`.
- Added lane evidence and match reasons to MCP `query` output through `queryCapabilities`, `queryLanes`, `matchReasons`, `rank`, `sourceRank`, and `processRank`.
- Added CLI `query --lanes`, `query --lanes --json`, and `query --explain`.
- Updated `query --help` so the help text names all query lanes and tells users to use `--lanes` and `--explain`.
- Preserved existing `avmatrix query "<intent>" --repo <repo>` behavior.
- Updated `query-health` output to keep threshold pass and exact pass as separate meanings and to include lane/rank/match evidence in matched/top results.

Canonical executable/build changes:

- `avmatrix\bin\avmatrix.exe` is now the only production AVmatrix CLI/runtime executable built by `avmatrix-launcher\build.ps1`.
- `avmatrix-launcher\server-bundle\avmatrix.exe` is no longer built and no longer exists after the full build.
- `avmatrix-launcher\server-bundle\avmatrix-server.exe` remains a launcher support wrapper and starts `avmatrix\bin\avmatrix.exe serve --host 127.0.0.1 --port 4848`.
- Launcher process cleanup now targets canonical `avmatrix\bin\avmatrix.exe` only when its command line is the launcher-owned `serve` process on port `4848`, instead of killing every canonical AVmatrix CLI/MCP process.
- The build script copies `lbug_shared.dll` into `avmatrix\bin` only when content differs. If the existing DLL is already identical but locked by another process, the build records it as up to date and continues.

Validation after full build:

| Command | Result |
|---|---|
| `powershell -ExecutionPolicy Bypass -File avmatrix-launcher\build.ps1` | pass; Web build completed, canonical CLI built to `avmatrix\bin\avmatrix.exe`, native DLL already up to date, launcher and server wrapper built. |
| `go test .\internal\mcp .\internal\cli -count=1` | pass: `internal/mcp`, `internal/cli`. |
| `go test . -count=1` in `avmatrix-launcher\src` | pass: `avmatrix-launcher`. |
| `go test . -count=1` in `avmatrix-launcher\server-wrapper` | pass: `avmatrix-server-wrapper`. |
| `Get-ChildItem avmatrix\bin, avmatrix-launcher\server-bundle -Filter *.exe` | only `avmatrix\bin\avmatrix.exe` and `avmatrix-launcher\server-bundle\avmatrix-server.exe` were present. |
| `Test-Path avmatrix-launcher\server-bundle\avmatrix.exe` | `False`. |
| `.\avmatrix\bin\avmatrix.exe analyze --force` | pass; final pre-commit run reported `files: scanned=762 parsed=568 unsupported=194 failed=0`, `nodes=85963 relationships=117955`. |
| `.\avmatrix\bin\avmatrix.exe version` | `1.2.2`. |
| `.\avmatrix\bin\avmatrix.exe query --help` | pass; help lists all eight query capability lanes plus `--lanes` and `--explain`. |
| `.\avmatrix\bin\avmatrix.exe query --lanes --json` | pass; JSON returned eight `queryCapabilities`. |
| `.\avmatrix\bin\avmatrix.exe query "generated AVmatrix skills AGENTS.md CLAUDE.md internal aicontext" --repo AVmatrix --limit 5 --explain` | pass; top definitions were `installBaseSkills`, `baseSkillContent`, `setupInstallEditorSkills`, `setupInstallSkillsTo`, `setupSkillTargetName`, with lane/match evidence. |
| `.\avmatrix\bin\avmatrix.exe query-health --repo AVmatrix --suite .\docs\query-health\2026-05-23-avmatrix-skill-system-upgrade-suite.json --limit 10 --out .\.tmp\2026-05-23-skill-system-query-health-phase1-canonical-final.json` | pass; threshold 8/8, exact 3/8, matched targets 46/54. |
| `.\avmatrix\bin\avmatrix.exe detect-changes --repo AVmatrix --scope all` | pass; final pre-commit run reported `changed_files=15`, `changed_count=391`, `affected_count=26`, `risk_level=critical`. Critical risk is expected for this slice because query/MCP/CLI and launcher runtime code changed. |

Pre-commit changed scope from `detect-changes`:

| Scope | Output |
|---|---|
| Changed app layers | `api=137`, `api_test=66`, `backend=110`, `backend_test=35`, `cli_launcher=18`, `docs=25`. |
| Affected app layers | `api=9`, `backend=6`, `cli_launcher=5`, `mixed=6`. |
| Changed functional areas | `cli=94`, `documentation=25`, `launcher=26`, `mcp=203`, `query=43`. |
| Affected functional areas | `cli=5`, `launcher=5`, `mcp=9`, `mixed=6`, `query=1`. |
| Resolution gap changes | `changedGapEntities=268`, `changedGapOccurrenceCount=269`; top changed targets included `strings.Contains`, `string`, `append`, and `t.Fatalf`. |

Phase 3 skill guidance requirements from Phase 1:

- `query` is the broad candidate-discovery command, not a replacement for exact `context`.
- `query` now exposes multiple capability lanes: owner, concept, execution-flow, API surface, graph-quality, docs/setup/AI-context, command-surface, and cross-repo discovery.
- Broad query results must be verified with `context` or exact source/file inspection before choosing edit surfaces.
- Query-health reports two separate meanings: threshold/usable pass and exact expected-target coverage. Agents must record missed exact targets instead of treating threshold pass as full coverage.
- `query --lanes --json` and `query --explain` are the user-facing CLI surfaces for lane and match evidence.

## Phase 1 P1-L Broad-Discovery Regression Evidence

Status: completed for P1-L.

Blast-radius checks before editing query-health regression code:

| Target | Result |
|---|---|
| `scoreQueryHealthCase` | CRITICAL; affects query-health scoring and reporting through `runQueryHealth`. This is blast-radius evidence, not a blocker. |
| `queryHealthActualResults` | CRITICAL; affects query-health conversion of query output into scored rows. This is blast-radius evidence, not a blocker. |
| `runQueryHealth` | CRITICAL; affects the CLI `query-health` command through `newQueryHealthCommand` and `NewRootCommand`. This is blast-radius evidence, not a blocker. |
| `readQueryHealthSuite` | CRITICAL; affects suite schema validation through `runQueryHealth`. This is blast-radius evidence, not a blocker. |

Implementation:

- Extended `query-health` suites with optional `expectedProcesses` so broad-discovery regression cases can assert returned execution-flow/process labels, not only files and symbols.
- Added process rows from MCP `query` output into query-health scoring, preserving process label, process rank, source rank, source type, and process match evidence.
- Added focused CLI tests proving process target matching and process-row conversion.
- Added `cross-repo-execution-flow-process-discovery` to `docs/query-health/2026-05-23-avmatrix-skill-system-upgrade-suite.json`.

Validation:

| Command | Result |
|---|---|
| `powershell -ExecutionPolicy Bypass -File avmatrix-launcher\build.ps1` | pass; full build completed before testing. |
| `go test .\internal\cli .\internal\mcp -count=1` | pass: `internal/cli` and `internal/mcp`. |
| `.\avmatrix\bin\avmatrix.exe analyze --force` | pass after implementation; `files: scanned=766 parsed=569 unsupported=197 failed=0`, `nodes=86837 relationships=119164`. |
| `.\avmatrix\bin\avmatrix.exe query-health --repo AVmatrix --suite .\docs\query-health\2026-05-23-avmatrix-skill-system-upgrade-suite.json --limit 10 --out .\.tmp\2026-05-23-skill-system-query-health-p1l-process-regression.json` | pass; `cases=9 thresholdPassed=9 thresholdFailed=0 exactPassed=4 exactFailed=5 matchedTargets=53/61 missedTargets=8`. |
| `.\avmatrix\bin\avmatrix.exe query "group query execution flow process sync contracts" --repo AVmatrix --limit 10 --explain --json` | pass; top definitions include `Query`, `Sync`, `groupQueryTool`, `Contracts`, and `newGroupCommand`; top process labels include `Sync -> NormalizeHTTPPath` and `Query -> GroupProcess`. |
| `.\avmatrix\bin\avmatrix.exe detect-changes --repo AVmatrix --scope all` | pass before commit; `changed_files=6`, `changed_count=84`, `affected_count=9`, `risk_level=high`. HIGH is expected because this slice changes shared CLI/query-health scoring. Changed app layers: `backend=44`, `backend_test=32`, `docs=8`; affected app layers: `backend=9`. |

P1-L regression case result:

| Case | Expected targets | Hit@5 | Hit@10 | Threshold | Exact | Matched | Missed |
|---|---:|---:|---:|---|---|---:|---:|
| `cross-repo-execution-flow-process-discovery` | 7 | 7 | 7 | PASS | PASS | 7 | 0 |

Matched process targets:

| Expected process | Rank | Source |
|---|---:|---|
| `Sync -> NormalizeHTTPPath` | 1 | `process` |
| `Query -> GroupProcess` | 2 | `process` |

Before/after interpretation:

- The Phase 1 benchmark already recorded the non-AI-context cross-repo case improving from `matched=4/8` and `missed=4` at baseline to `matched=8/8` and `missed=0` after the ranking fix.
- P1-L adds an execution-flow-specific regression case on top of that fixed behavior. It proves the broad query still returns meaningful concept-to-flow/process candidates while keeping strong owner surfaces at the top.

## Phase 1.5 Graph Health CLI Surface Evidence

### P1.5-A Current Implementation Trace

Status: recorded before graph-health CLI edits.

Graph refresh:

| Command | Result |
|---|---|
| `.\avmatrix\bin\avmatrix.exe analyze --force` | pass after P1-L commit; `files: scanned=766 parsed=569 unsupported=197 failed=0`, `nodes=86838 relationships=119165`. |

Source trace:

| Surface | Files and symbols |
|---|---|
| Core graph-health computation | `internal/graphhealth/compute.go`: `Compute`, `ComputeSummary`, `computeComponents`, `reachableNodes`, `largestDetachedComponents`, `componentSummary`, `addNodeHealthToSummary`. |
| Graph-health policy and data contract | `internal/graphhealth/policy.go`: `PolicyVersion`, `IsCounted`, `CountedEdgeTypes`, `StructuralEdgeTypes`, topology statuses, confidence levels, diagnostic classifications/actionability, resolution health buckets, `NodeHealth`, `ComponentSummary`, `Summary`. |
| Topology/count metadata | `ComputeSummary` builds counted incoming/outgoing maps from `IsCounted`, excluded edge counts from structural/other categories, component IDs/sizes/root reachability, root nodes, detached components, expected isolation reasons, diagnostics, confidence, resolution confidence, and summary inventory counts. |
| HTTP graph payload | `internal/httpapi/graph.go`: `graphPayload` calls `graphhealth.ComputeSummary` and returns `graphHealth` summary with graph nodes/relationships and semantic status. |
| HTTP graph-health report | `internal/httpapi/graph.go`: `handleGraphHealthReport`, `graphHealthReportCandidates`, `graphHealthReportPriority`, `graphHealthReportPriorityRank`, `graphHealthReportLimit`. Candidate priority order is topology `no_incoming`, `detached_component`, `true_isolated`, `no_outgoing`, `unknown_connectivity`, then diagnostic `unresolved_reference`. |
| HTTP graph-health explain | `internal/httpapi/graph.go`: `handleGraphHealthExplain`, `graphHealthNodeExplain`, `graphHealthComponentExplain`, `graphHealthNodeRelationships`, `graphHealthComponentRelationshipSamples`, `countComponentCountedEdges`, `graphHealthDiagnosticCount`. Explain supports exactly one of `nodeId` or `componentId`, strips content unless requested, and returns counted/excluded relationship samples. |
| Web filters and labels | `avmatrix-web/src/lib/graph-health-filters.ts`: `getNodeGraphHealth`, `graphHealthMatchesFilters`, `graphNodeMatchesHealthFilters`, topology/reason/diagnostic/confidence labels, descriptions, next-action text, and count helpers. |
| CLI root registration | `internal/cli/command.go`: `NewRootCommand` currently registers `query-health`, `resolution-inventory`, and `source-site-accuracy`, but not `graph-health`. |
| Existing CLI graph-health-adjacent surface | `internal/cli/resolution_inventory_command.go`: `runResolutionInventory` already calls `graphhealth.ComputeSummary`, but it reports resolution inventory, not graph-health topology/report/explain/components. |

AVmatrix context checks:

| Symbol | Key incoming/current ownership |
|---|---|
| `ComputeSummary` | Called by `graphhealth.Compute`, `httpapi.graphPayload`, `httpapi.handleGraphHealthReport`, `cli.runResolutionInventory`, and graphhealth tests. |
| `graphHealthReportCandidates` | Owned by `internal/httpapi/graph.go`; called by `handleGraphHealthReport` and HTTP tests. |
| `graphHealthNodeExplain` | Owned by `internal/httpapi/graph.go`; called by `handleGraphHealthExplain`. |
| `graphHealthComponentExplain` | Owned by `internal/httpapi/graph.go`; called by `handleGraphHealthExplain`. |
| `NewRootCommand` | Owns CLI command registration; no graph-health command is present. |

Trace conclusion:

- Graph-health computation is already shared under `internal/graphhealth`.
- Report/explain/component behavior is currently HTTP-owned even though CLI needs the same semantics. P1.5-C should move that behavior into a reusable `graphhealth` owner and leave HTTP as a thin adapter so CLI and API do not drift.
- `resolution-inventory` must remain separate: it reports resolution gap inventory and uses graph-health summary, but it is not a substitute for topology/report/explain/components graph-health triage.

### P1.5-B Graph Health CLI Contract

Status: defined before graph-health CLI edits.

Command family:

| Command | Purpose | Required/optional inputs | Output |
|---|---|---|---|
| `avmatrix graph-health summary --repo <repo> [--json]` | Compute and print graph-level topology, component, diagnostic, confidence, and resolution-health summary from the selected indexed graph. | `--repo` optional only when a single repo is indexed. Requires fresh analyze output. | Table by default; full `graphhealth.Summary` JSON with input metadata when `--json`. |
| `avmatrix graph-health report --repo <repo> [--limit <n>] [--include-expected] [--json]` | Print triage candidates using the same priority rules as HTTP report. | `--limit` defaults to HTTP default and is capped at the same max; `--include-expected` includes expected-isolated nodes. Requires fresh analyze output. | Table by default; JSON report with summary, total/returned candidate counts, candidates, priority/dimension, topology/confidence/diagnostics/component fields. |
| `avmatrix graph-health components --repo <repo> [--limit <n>] [--json]` | List component-level graph-health summaries for detached or otherwise notable components. | `--limit` defaults to a bounded list. Requires fresh analyze output. | Table by default; JSON component summaries with component ID, node count, counted edge count, detached/reachable flags, root/sample node IDs, and health/resolution counts. |
| `avmatrix graph-health explain <node-id-or-name> --repo <repo> [--json]` | Explain one graph node by exact ID or unique node name. | Positional selector; if ambiguous or missing, fail clearly. Requires fresh analyze output. | Table by default; JSON node explain with node, health, counted incoming/outgoing relationships, excluded relationships. |
| `avmatrix graph-health explain --component <component-id> --repo <repo> [--json]` | Explain one component by component ID. | `--component` must not be combined with a node selector. Requires fresh analyze output. | Table by default; JSON component explain with component aggregate counts, sample nodes, counted relationship samples, excluded relationship samples, and sample limit. |

Contract constraints:

- CLI must load the selected repo's indexed `graph.json` from the registry/storage path, not a separately supplied ad hoc graph path, because the user-facing contract is repo-oriented.
- Freshness behavior must match current query-health expectations: if indexed commit and current commit differ, fail and tell the user to run `avmatrix analyze --force`.
- Report/explain/components must reuse shared graph-health semantics moved from HTTP, including topology status, counted/excluded edge policy, confidence, diagnostics, resolution confidence, component membership, priority ordering, and sample limits.
- Missing graph returns a clear "Graph not found. Run: avmatrix analyze --force" style error.
- Unknown node selector returns "Graph node not found"; ambiguous node name returns an ambiguity error listing candidate IDs; unknown component returns "Graph component not found".
- Table output must include stable identifiers users can feed into `explain`, especially node IDs and component IDs.

### P1.5-C/D/E Shared Graph-Health CLI Implementation Evidence

Status: implementation and focused validation complete; guidance/smoke remain in P1.5-F/G.

Blast radius:

| Symbol | Risk | Interpretation |
|---|---:|---|
| `graphHealthReportCandidates` | CRITICAL | HTTP graph-health report priority path; refactored into shared `graphhealth.ReportCandidates` with HTTP wrapper kept for tests/parity. |
| `graphHealthNodeExplain` | CRITICAL | HTTP graph-health node explain path; moved to shared `graphhealth.ExplainNode`. |
| `graphHealthComponentExplain` | CRITICAL | HTTP graph-health component explain path; moved to shared `graphhealth.ExplainComponent`. |
| `graphHealthReportPriority` | CRITICAL | Shared triage ordering path; moved to `graphhealth.ReportPriority` / `ReportPriorityRank`. |
| `NewRootCommand` | CRITICAL | CLI launcher registration point; scoped edit only added `newGraphHealthCommand()` to the command tree. |
| `handleGraphHealthExplain` | LOW | API adapter now calls shared graphhealth explain and still strips response content/internal diagnostics through `graphNodeForResponse`. |
| `handleGraphHealthReport` | LOW | API adapter now calls shared graphhealth report builder. |

Implementation:

| Task | Evidence |
|---|---|
| P1.5-C shared owner | Added `internal/graphhealth/report.go` with `BuildReport`, `ReportCandidates`, `ReportPriority`, `ExplainNode`, `ExplainComponent`, `ComponentSummaries`, shared response structs, report limit constants, counted/excluded relationship sampling, component aggregation, and priority ordering. |
| HTTP parity | `internal/httpapi/graph.go` now aliases shared graphhealth response/candidate types, uses `graphhealth.BuildReport`, `graphhealth.ExplainNode`, and `graphhealth.ExplainComponent`, and keeps only HTTP-specific response sanitization. |
| P1.5-D CLI command surface | Added `internal/cli/graph_health_command.go` and registered `avmatrix graph-health` in `NewRootCommand`. Subcommands implemented: `summary`, `report`, `components`, `explain`; all support table output and `--json`, load selected repo graph from registry storage, and enforce stale commit/missing graph errors. |
| Node/component selectors | `explain` accepts exact node ID first, then unique node `name`; ambiguous names list matching IDs; `--component` is mutually exclusive with positional node selector. |
| P1.5-E tests | Added `internal/cli/graph_health_command_test.go` covering summary counts, report ordering/limit, table/JSON output, components output, explain by ID/name/component, missing node, ambiguous name, missing graph, stale commit, and help registration. Existing HTTP graph-health tests continue to validate API semantics through shared wrappers. |

Validation:

| Command | Result |
|---|---|
| `powershell -ExecutionPolicy Bypass -File avmatrix-launcher\build.ps1` | pass; Go build and Web production build completed. Vite emitted existing chunk-size/dynamic-import warnings only. |
| `go test .\internal\graphhealth .\internal\httpapi .\internal\cli -count=1` | pass; `graphhealth` 1.158s, `httpapi` 2.327s, `cli` 8.963s. |

### P1.5-F Graph-Health Guidance Evidence

Status: complete.

Blast radius:

| Symbol | Risk | Interpretation |
|---|---:|---|
| `renderAVmatrixBlock` | CRITICAL | Analyze-generated AGENTS/CLAUDE context path; scoped edit only adds `graph-health` to graph-refresh guidance and command selection. |

Updated guidance sources:

| File | Change |
|---|---|
| `internal/aicontext/aicontext.go` | Added `graph-health` to freshness-sensitive graph command list and command selection guide. |
| `internal/aicontext/skills/avmatrix-cli.md` | Added `graph-health` command examples for summary/report/components/explain and clarified graph-quality command boundaries. |
| `internal/aicontext/skills/avmatrix-guide.md` | Added `avmatrix graph-health` to CLI graph-quality diagnostic commands and clarified that it answers topology/component/diagnostic triage, not retrieval or source-site accuracy. |
| `README.md` | Added `avmatrix graph-health` to direct graph tools and semantic graph diagnostics with summary/report/components examples. |
| `internal/aicontext/aicontext_test.go` | Updated generated context expectations for `graph-health`. |

Validation:

| Command | Result |
|---|---|
| `powershell -ExecutionPolicy Bypass -File avmatrix-launcher\build.ps1` | pass after guidance changes; same existing Vite chunk-size/dynamic-import warnings only. |
| `go test .\internal\aicontext .\internal\cli -count=1` | pass; `aicontext` 0.769s, `cli` 8.412s. |

Note: dedicated `internal/aicontext/skills/avmatrix-graph-quality.md` does not exist yet in this phase; the new dedicated graph-quality skill remains part of the later Phase 3 skill expansion. The current generated guidance surfaces that exist today now mention the implemented `graph-health` CLI behavior.

### P1.5-G Real Graph-Health CLI Smoke Evidence

Status: complete.

Fresh graph:

| Command | Result |
|---|---|
| `.\avmatrix\bin\avmatrix.exe analyze --force` | pass; `files: scanned=769 parsed=572 unsupported=197 failed=0`, `nodes=87382 relationships=120022`, graph path `E:\AVmatrix-GO\.avmatrix\graph.json`. |

Smoke commands:

| Command | Representative output |
|---|---|
| `.\avmatrix\bin\avmatrix.exe graph-health --help` | Help lists `components`, `explain`, `report`, `summary`, and persistent `--repo`. |
| `.\avmatrix\bin\avmatrix.exe graph-health summary --repo AVmatrix` | `nodes=87382 relationships=120022 countedRelationships=25951 components=79089 detachedComponents=62 rootNodes=852`; topology counts include `connected:2897`, `detached_component:236`, `no_incoming:1743`, `no_outgoing:3482`, `true_isolated:79024`; diagnostics `unresolved_reference:63576`. |
| `.\avmatrix\bin\avmatrix.exe graph-health summary --repo AVmatrix --json` | JSON parse confirmed `summary.nodeCount=87382`, `totals.relationships=120022`, `summary.countedRelationshipCount=25951`, `summary.componentCount=79089`, `summary.detachedComponentCount=62`, `summary.unresolvedReferenceCount=63576`. |
| `.\avmatrix\bin\avmatrix.exe graph-health report --repo AVmatrix --limit 20 --json` | JSON parse confirmed `totalCandidates=47316`, `returnedCandidates=20`; first candidate `File:avmatrix-web/src/components/RightPanel.tsx`, priority `no_incoming`, dimension `topology`, topology `no_incoming`, confidence `candidate`, component `component_000001`. |
| `.\avmatrix\bin\avmatrix.exe graph-health components --repo AVmatrix --limit 20 --json` | JSON parse confirmed `totalComponents=79089`, `returnedComponents=20`; first listed component `component_000936`, `nodeCount=21`, `countedEdgeCount=53`, `detached=true`, `reachableFromRoot=false`. |
| `.\avmatrix\bin\avmatrix.exe graph-health explain "File:avmatrix-web/src/components/RightPanel.tsx" --repo AVmatrix` | Table output: `topology=no_incoming`, `confidence=candidate`, `incoming=0`, `outgoing=6`, `excluded=6`, `component="component_000001"`, `resolutionConfidence=unknown`, `resolutionGaps=0`. |
| `.\avmatrix\bin\avmatrix.exe graph-health explain "File:avmatrix-web/src/components/RightPanel.tsx" --repo AVmatrix --json` | JSON parse confirmed node explain with `kind=node`, same node ID, topology `no_incoming`, incoming `0`, outgoing `6`, component `component_000001`, excluded relationships `6`. |
| `.\avmatrix\bin\avmatrix.exe graph-health explain --component component_000001 --repo AVmatrix --json` | JSON parse confirmed component explain with `kind=component`, `nodeCount=7990`, `countedEdgeCount=25322`, `detached=false`, `reachableFromRoot=true`, `sampleNodes=20`. |
| `.\avmatrix\bin\avmatrix.exe graph-health explain "__missing_graph_health_node__" --repo AVmatrix` | Expected failure; exit code `1`, output `graph node not found: __missing_graph_health_node__`. |

Limitations:

- The CLI freshness check compares indexed commit to current commit, matching existing `query-health` behavior. It does not treat uncommitted working tree edits as stale by itself; the workflow rule still requires `avmatrix analyze --force` before graph-based work.

Pre-commit checks:

| Command | Result |
|---|---|
| `git diff --check` | pass. |
| `.\avmatrix\bin\avmatrix.exe analyze --force` | pass after evidence/benchmark updates; `files: scanned=769 parsed=572 unsupported=197 failed=0`, `nodes=87383 relationships=120023`. |
| `.\avmatrix\bin\avmatrix.exe detect-changes --repo AVmatrix --scope all` | pass; summary `changed_files=11`, `changed_count=79`, `affected_count=35`, `risk_level=critical`; changed App Layers `api=36`, `backend=5`, `backend_test=12`, `docs=26`; affected App Layers `api=1`, `backend=1`, `mixed=33`. Critical scope is expected for shared graph-health/API/CLI/generator changes. |

## Phase 1.6 CLI Parity Audit And Missing Command Surfaces Evidence

Status: complete.

Fresh graph and impact commands:

| Command | Result |
|---|---|
| `.\avmatrix\bin\avmatrix.exe analyze --force` | pass before P1.6 impact checks; `files: scanned=770 parsed=573 unsupported=197 failed=0`, `nodes=87524 relationships=120184`. |
| `.\avmatrix\bin\avmatrix.exe impact "NewRootCommand" --repo AVmatrix --direction upstream` | CRITICAL; root CLI registration path affects `cmd/avmatrix/main.go:main` and 11 execution processes. This is blast-radius evidence, not a blocker. |
| `.\avmatrix\bin\avmatrix.exe impact "printLocalMCPTool" --repo AVmatrix --direction upstream` | CRITICAL; shared direct CLI tool wrapper affects query/context/impact/cypher/detect-changes style commands through `NewRootCommand` and CLI launcher paths. |
| `.\avmatrix\bin\avmatrix.exe impact "callLocalMCPTool" --repo AVmatrix --direction upstream` | CRITICAL; shared local MCP invocation helper affects direct tool commands and query-health local query execution. |
| `.\avmatrix\bin\avmatrix.exe analyze --force` | pass before guidance impact checks; `files: scanned=771 parsed=574 unsupported=197 failed=0`, `nodes=87760 relationships=120503`. |
| `.\avmatrix\bin\avmatrix.exe impact "renderAVmatrixBlock" --repo AVmatrix --direction upstream` | CRITICAL; analyze-generated AGENTS/CLAUDE context path through `GenerateAIContextFiles` and analyze post-run. |
| `.\avmatrix\bin\avmatrix.exe impact "contextResource" --repo AVmatrix --direction upstream` | LOW; MCP repo context resource text surface. |
| `.\avmatrix\bin\avmatrix.exe impact "setupResource" --repo AVmatrix --direction upstream` | LOW; MCP setup resource text surface. |

### P1.6-A/B Command Surface Inventory

Source and binary inventory commands:

| Command | Result |
|---|---|
| `.\avmatrix\bin\avmatrix.exe --help` | pass; visible help lists 28 top-level commands including Cobra built-ins `completion` and `help`. New user-facing commands are `api` and `rename`. |
| `go run .\cmd\avmatrix --help` | pass; source-run help matches the freshly built canonical binary help for the visible command list. |
| `rg -n 'Use:\s+"' internal\cli -g '*.go'` | source registration inventory recorded 26 explicit visible top-level user commands plus 2 hidden lifecycle families, `hook` and `package`. |
| `rg -n 'Name:\s+"..." internal\mcp -g '*.go'` | source MCP inventory recorded 16 tools: `list_repos`, `query`, `cypher`, `context`, `detect_changes`, `rename`, `impact`, `route_map`, `tool_map`, `shape_check`, `api_impact`, `group_list`, `group_sync`, `group_contracts`, `group_query`, `group_status`. |
| `rg -n 'resourceDefinitions|resourceTemplates|promptDefinitions' internal\mcp -g '*.go'` | source MCP surface remains 2 fixed resources, 6 resource templates, and 2 prompt templates. |
| `rg -n 'HandleFunc\(|/api/' internal\httpapi avmatrix-web\src\services -g '*.go' -g '*.ts' -g '*.tsx'` | HTTP/Web inventory recorded `/api/heartbeat`, `/api/info`, `/api/repos`, `/api/repo`, `/api/local/folder-picker`, `/api/graph`, `/api/graph/explain`, `/api/graph/report`, `/api/file`, `/api/grep`, `/api/query`, `/api/processes`, `/api/process`, `/api/clusters`, `/api/cluster`, `/api/analyze`, `/api/analyze/`, `/api/search`, `/api/embed`, `/api/embed/`, `/api/mcp`, `/api/session/status`, `/api/session/chat`, and `/api/session/`. |

### P1.6-C/D Parity Matrix And Accepted Design

Accepted CLI additions use a grouped API command family plus one top-level rename command:

| Surface | Classification | P1.6 decision |
|---|---|---|
| MCP `rename` | `has_cli` after P1.6 | Add `avmatrix rename [symbol_name] <new_name>` with `--repo`, `--uid`, `--file`, `--apply`, and `--json`. Default remains dry run. |
| MCP `route_map` | `has_cli` after P1.6 | Add `avmatrix api route-map [route] --repo <repo> [--json]`, delegating to MCP `route_map`. |
| MCP `tool_map` | `has_cli` after P1.6 | Add `avmatrix api tool-map [tool] --repo <repo> [--json]`, delegating to MCP `tool_map`. |
| MCP `shape_check` | `has_cli` after P1.6 | Add `avmatrix api shape-check [route] --repo <repo> [--json]`, delegating to MCP `shape_check`. |
| MCP `api_impact` | `has_cli` after P1.6 | Add `avmatrix api impact [route] --repo <repo> [--file <path>] [--json]`, delegating to MCP `api_impact`. |
| Graph explain/report | `has_cli` | Already covered by `avmatrix graph-health report`, `components`, and `explain` from Phase 1.5. |
| HTTP/Web `grep` | `mcp_api_web_only_by_design` for P1.6 | Not promoted in this slice. CLI users have native `rg` for text grep and `avmatrix augment`/`query` for graph-context search. |
| HTTP/Web `search` | `covered_by_existing_cli` plus `follow_up` for semantic Web behavior | `avmatrix query` is the CLI graph search surface. Web/API semantic search remains runtime-oriented and can be revisited if a product CLI search command is required. |
| HTTP/Web `processes` and `process` | `mcp_api_web_only_by_design` for P1.6 | Exposed through MCP resources `avmatrix://repo/{name}/processes` and `avmatrix://repo/{name}/process/{processName}`; no CLI command added in this slice. |
| HTTP/Web `clusters` and `cluster` | `mcp_api_web_only_by_design` for P1.6 | Exposed through MCP resources `avmatrix://repo/{name}/clusters` and `avmatrix://repo/{name}/cluster/{clusterName}`; no CLI command added in this slice. |
| HTTP/Web `file` | `mcp_api_web_only_by_design` for P1.6 | Web/API source viewer endpoint only. CLI users can inspect files directly. |
| HTTP/Web `repo`, `repos`, `analyze`, `embed`, `session`, local folder picker | `has_cli` or `web_runtime_only` | Repo/analyze have CLI equivalents; embed/session/folder picker remain Web/runtime workflows in this slice. |
| Hidden `package` and `hook` | `hidden_lifecycle_only` | Kept hidden. Guidance now labels these as lifecycle helpers, not normal repo-analysis commands. |

Design constraints:

- The new CLI commands call existing MCP tool owners through `callLocalMCPTool`, avoiding a second implementation of rename/API semantics.
- `--json` strips the MCP next-step hint and prints only the primary JSON payload.
- Positional selectors and selector flags are mutually exclusive for API subcommands.
- Existing `query`, `group`, graph-health, and hidden lifecycle command parsing is preserved.

### P1.6-E/F Implementation And Tests

Implementation:

| File | Change |
|---|---|
| `internal/cli/api_command.go` | Added `avmatrix api` with `route-map`, `tool-map`, `shape-check`, and `impact` subcommands. |
| `internal/cli/tool_command.go` | Added `avmatrix rename`; added JSON-only wrapper behavior through `printLocalMCPToolWithJSON` and `primaryMCPToolPayload`. |
| `internal/cli/command.go` | Registered `newAPICommand()` and `newRenameCommand()` in the root CLI tree. |
| `internal/cli/api_command_test.go` | Added focused tests for API command JSON success paths, duplicate selector error handling, and rename dry-run JSON behavior through the local MCP runtime. |
| `internal/cli/command_test.go` | Updated root/help tests for `api` and `rename`. |

Validation:

| Command | Result |
|---|---|
| `powershell -ExecutionPolicy Bypass -File avmatrix-launcher\build.ps1` | pass before focused tests; Go build and Web production build completed. Existing Vite chunk-size/dynamic-import warnings only. |
| `go test .\internal\cli -count=1` | pass; `internal/cli` 8.287s. |
| `powershell -ExecutionPolicy Bypass -File avmatrix-launcher\build.ps1` | pass after guidance updates; same existing Vite warnings only. |
| `go test .\internal\cli .\internal\aicontext .\internal\mcp -count=1` | pass; `cli` 10.186s, `aicontext` 0.756s, `mcp` 6.279s. |

### P1.6-G Guidance Updates

Updated guidance sources:

| File | Change |
|---|---|
| `internal/aicontext/aicontext.go` | Generated command selection now lists CLI `rename` and `api route-map`/`api tool-map`/`api shape-check`/`api impact` equivalents beside MCP tools. Hidden `package`/`hook` rows are labeled lifecycle helpers. |
| `internal/aicontext/skills/avmatrix-cli.md` | Added `rename` and `api` command sections and clarified CLI vs MCP use. |
| `internal/aicontext/skills/avmatrix-guide.md` | Added CLI equivalent table for query/context/impact/detect-changes/cypher/rename/API parity commands. |
| `internal/mcp/resources.go` | Added CLI equivalents to repo context and setup resources. |
| `README.md` | Added `rename` and `api ...` commands to direct graph tools and documented that they delegate to MCP owners. |
| `internal/aicontext/aicontext_test.go`, `internal/mcp/resources_parity_test.go` | Updated generated guidance/resource expectations for the new command taxonomy. |

### P1.6-H User-Facing Smoke

Fresh graph before smoke:

| Command | Result |
|---|---|
| `.\avmatrix\bin\avmatrix.exe analyze --force` | pass after P1.6 guidance updates; `files: scanned=771 parsed=574 unsupported=197 failed=0`, `nodes=87762 relationships=120505`. |

Help and source smoke:

| Command | Representative output |
|---|---|
| `.\avmatrix\bin\avmatrix.exe --help` | Lists `api` and `rename`; visible help count is 28 including `completion` and `help`. |
| `go run .\cmd\avmatrix --help` | Matches canonical binary help for visible command list. |
| `.\avmatrix\bin\avmatrix.exe api --help` | Lists `impact`, `route-map`, `shape-check`, `tool-map`, and persistent `--repo`. |
| `.\avmatrix\bin\avmatrix.exe rename --help` | Lists `--apply`, `--file`, `--json`, `--repo`, and `--uid`. |

AVmatrix repo smoke:

| Command | Representative output |
|---|---|
| `.\avmatrix\bin\avmatrix.exe rename NewRootCommand NewRootCommand2 --repo AVmatrix --json` | Dry-run JSON parse succeeded: `status=success applied=false files=4 edits=4 graphEdits=4 textSearchEdits=0`. |
| `.\avmatrix\bin\avmatrix.exe cypher "MATCH (n:Route) RETURN n.id AS id, n.name AS name, n.filePath AS filePath LIMIT 10" --repo AVmatrix` | `_No rows_`; current AVmatrix self graph has no `Route` nodes, so positive API command smoke used a small analyzed fixture below. |
| `.\avmatrix\bin\avmatrix.exe cypher "MATCH (n:Tool) RETURN n.id AS id, n.name AS name, n.filePath AS filePath LIMIT 10" --repo AVmatrix` | `_No rows_`; current AVmatrix self graph has no `Tool` nodes. |

Positive API/rename smoke fixture:

| Command | Result |
|---|---|
| Created `.tmp\p1-6-cli-parity-smoke-20260526173638` | Fixture contains `app/api/users/route.ts`, `src/client.ts`, and `README.md`. |
| `.\avmatrix\bin\avmatrix.exe analyze .tmp\p1-6-cli-parity-smoke-20260526173638 --force --skip-git --no-stats --name p1-6-cli-parity-smoke-20260526173638` | pass; `files: scanned=3 parsed=2 unsupported=1 failed=0`, `nodes=22 relationships=20`. |
| `.\avmatrix\bin\avmatrix.exe api route-map "/api/users" --repo p1-6-cli-parity-smoke-20260526173638 --json` | JSON parse confirmed `total=1`, route `/api/users`, handler `app/api/users/route.ts`, `consumers=1`, `flows=0`. |
| `.\avmatrix\bin\avmatrix.exe api tool-map query --repo p1-6-cli-parity-smoke-20260526173638 --json` | JSON parse confirmed `total=1`, tool `query`, file `src/client.ts`, description `Query tool`, `flows=0`. |
| `.\avmatrix\bin\avmatrix.exe api shape-check "/api/users" --repo p1-6-cli-parity-smoke-20260526173638 --json` | JSON parse confirmed `total=1`, `mismatches=1`, route status `MISMATCH`. |
| `.\avmatrix\bin\avmatrix.exe api impact "/api/users" --repo p1-6-cli-parity-smoke-20260526173638 --json` | JSON parse confirmed route `/api/users`, handler `app/api/users/route.ts`, `consumers=1`, `flows=0`, `risk=MEDIUM`, `mismatches=1`. |
| `.\avmatrix\bin\avmatrix.exe rename loadUsers loadUsers2 --repo p1-6-cli-parity-smoke-20260526173638 --json` | Dry-run JSON parse confirmed `status=success`, `applied=false`, `files=1`, `edits=1`, `graphEdits=1`, `textSearchEdits=0`. |

Pre-commit checks:

| Command | Result |
|---|---|
| `git diff --check` | pass. |
| `.\avmatrix\bin\avmatrix.exe analyze --force` | pass after checklist/evidence/benchmark updates; `files: scanned=771 parsed=574 unsupported=197 failed=0`, `nodes=87768 relationships=120511`. |
| `.\avmatrix\bin\avmatrix.exe detect-changes --repo AVmatrix --scope all` | pass; summary `changed_files=13`, `changed_count=103`, `affected_count=20`, `risk_level=critical`; changed App Layers `api=4`, `api_test=2`, `backend=53`, `backend_test=16`, `docs=28`; affected App Layers `backend=3`, `mixed=17`. Critical scope is expected for root CLI registration, direct MCP wrapper, AI-context generator, MCP resource guidance, and docs/test changes. |

## Phase 1.7 MCP Prompt Template Accuracy Evidence

Status: complete.

Fresh graph and impact commands:

| Command | Result |
|---|---|
| `.\avmatrix\bin\avmatrix.exe analyze --force` | pass before P1.7 impact checks; `files: scanned=771 parsed=574 unsupported=197 failed=0`, `nodes=87768 relationships=120511`. |
| `.\avmatrix\bin\avmatrix.exe impact "promptDefinitions" --repo AVmatrix --direction upstream` | CRITICAL; prompt list registration affects the MCP `prompts/list` surface through `Server.handle` and 7 execution processes. This is blast-radius evidence, not an edit ban. |
| `.\avmatrix\bin\avmatrix.exe impact "detectImpactPrompt" --repo AVmatrix --direction upstream` | CRITICAL; generated `detect_impact` prompt body flows through `getPrompt` and `Server.handle` into MCP prompt execution. |
| `.\avmatrix\bin\avmatrix.exe impact "generateMapPrompt" --repo AVmatrix --direction upstream` | CRITICAL; generated `generate_map` prompt body flows through `getPrompt` and `Server.handle` into MCP prompt execution. |
| `.\avmatrix\bin\avmatrix.exe impact "getPrompt" --repo AVmatrix --direction upstream` | CRITICAL; shared prompt dispatcher affects MCP prompt retrieval. |
| `.\avmatrix\bin\avmatrix.exe impact "renderAVmatrixBlock" --repo AVmatrix --direction upstream` | CRITICAL; generated `AGENTS.md` / `CLAUDE.md` context path through AI-context generation. |
| `.\avmatrix\bin\avmatrix.exe impact "setupResource" --repo AVmatrix --direction upstream` | LOW; MCP setup resource guidance text surface. |
| `.\avmatrix\bin\avmatrix.exe impact "contextResource" --repo AVmatrix --direction upstream` | LOW; MCP repo context guidance text surface. |

### P1.7-A/B Prompt Inventory And Audit

Prompt inventory:

| Prompt | Arguments | Runtime role | Source/guidance references |
|---|---|---|---|
| `detect_impact` | `scope`, `base_ref` | Agent workflow template for pre-commit change detection, context, impact, freshness rules, and blast-radius interpretation. It produces analysis text, not files. | `internal/mcp/prompts.go`, README MCP prompts table, MCP setup resource, AI-context generated guidance, `avmatrix-guide` embedded skill. |
| `generate_map` | `repo` | Agent workflow template for evidence-backed architecture-map drafting. It may draft `ARCHITECTURE.md`-ready content, but tells the agent not to edit/create files unless the user explicitly asks. | `internal/mcp/prompts.go`, README MCP prompts table, MCP setup resource, repo context resource prompt list, AI-context generated guidance, `avmatrix-guide` embedded skill. |

Audit findings before the fix:

| Finding | Resolution |
|---|---|
| `generate_map` used `{name}` when `repo` was omitted, which could become a bogus actionable resource URI. | The prompt now reads `avmatrix://repos`, selects the single repo or workspace-path match, and stops to ask the user if multiple repos remain ambiguous. |
| Repo and process resource names need URI escaping. | The prompt now uses URL-escaped repo and process URI guidance and tests assert escaped repo resource paths. |
| Stale graph handling was not explicit. | The prompt now requires context/freshness checks and `avmatrix analyze --force` before graph-based mapping when stale or required by repo rules. |
| "Top 5 most important processes" had no deterministic selection rule. | The prompt now selects representative clusters/processes by user request, runtime/API/tool relevance, graph surface involvement, step count, and centrality if available, while recording the reason for each selection. |
| Architecture docs and Mermaid diagrams could be inferred beyond read graph evidence. | The prompt now requires claims, nodes, edges, dependencies, layers, and ownership to come only from resources/tools/commands the agent actually read, with uncertainty notes for incomplete evidence. |

### P1.7-C Through H Implementation

Implementation:

| File | Change |
|---|---|
| `internal/mcp/prompts.go` | Rewrote `generate_map` and `detect_impact` bodies; added repo discovery, URL-escaped resource URI guidance, freshness handling, deterministic selection criteria, evidence-only output rules, and current change-detection wording. |
| `internal/mcp/prompts_test.go` | Added focused prompt tests for `generate_map` with repo, `generate_map` without repo, and `detect_impact` workflow wording. |
| `internal/mcp/server_test.go` | Expanded MCP prompt server assertions for escaped resource URIs, freshness language, no actionable `{name}`, evidence-only architecture guidance, and process selection criteria. |
| `internal/mcp/resources.go` | Added MCP prompt guidance to setup and repo context resources. |
| `internal/mcp/resources_parity_test.go` | Updated setup-resource expectations for MCP prompt guidance. |
| `internal/aicontext/aicontext.go` | Added generated MCP prompt guidance and prompt/workflow distinction to AI-context output. |
| `internal/aicontext/aicontext_test.go` | Updated AI-context expectations for prompt guidance. |
| `internal/aicontext/skills/avmatrix-guide.md` | Added embedded guide section for MCP prompts and prompt-vs-command distinction. |
| `README.md` | Updated MCP prompt table and clarified that prompts are agent templates, not CLI commands. |

Validation:

| Command | Result |
|---|---|
| `powershell -ExecutionPolicy Bypass -File avmatrix-launcher\build.ps1` | pass before focused tests; Go build and Web production build completed. Existing Vite chunk-size/dynamic-import warnings only. |
| `go test .\internal\mcp .\internal\aicontext -count=1` | pass; `internal/mcp` 4.849s, `internal/aicontext` 0.660s. |
| `powershell -ExecutionPolicy Bypass -File avmatrix-launcher\build.ps1` | pass after final prompt-command wording patch; same existing Vite warnings only. |
| `go test .\internal\mcp .\internal\aicontext -count=1` | pass; `internal/mcp` 4.794s, `internal/aicontext` 0.758s. |

### P1.7-I Runtime MCP Prompt Smoke

Runtime smoke:

| Check | Result |
|---|---|
| `.\avmatrix\bin\avmatrix.exe mcp` JSON-RPC `prompts/list` | pass; response includes `detect_impact`, `generate_map`, `scope`, `base_ref`, and `repo`. |
| `prompts/get generate_map` with `repo=AVmatrix` | pass; body includes `avmatrix://repo/AVmatrix/context`, `avmatrix://repo/AVmatrix/clusters`, `avmatrix://repo/AVmatrix/processes`, freshness instruction `avmatrix analyze --force`, and evidence-only / do-not-invent guidance. No actionable `{name}` placeholder remains. |
| `prompts/get generate_map` without repo | pass; body includes `avmatrix://repos`, single-repo/workspace-match selection rules, ambiguity stop/ask-user behavior, and no actionable `{name}` placeholder. |
| `prompts/get detect_impact` with `scope=all`, `base_ref=HEAD~1` | pass; body includes MCP `detect_changes`, CLI fallback `avmatrix detect-changes`, and states HIGH/CRITICAL are blast-radius warnings, not an edit ban. |

Intentional limitation:

- MCP prompts are executable agent templates, but they still depend on the receiving agent actually reading the named resources/tools/commands. The templates now forbid unsupported claims; they do not themselves fetch graph evidence.

Pre-commit checks:

| Command | Result |
|---|---|
| `git diff --check` | pass. |
| `.\avmatrix\bin\avmatrix.exe analyze --force` | pass after checklist/evidence/benchmark updates; `files: scanned=772 parsed=575 unsupported=197 failed=0`, `nodes=87883 relationships=120603`. |
| `.\avmatrix\bin\avmatrix.exe detect-changes --repo AVmatrix --scope all` | pass; summary `changed_files=11`, `changed_count=49`, `affected_count=9`, `risk_level=high`; changed App Layers `api=15`, `api_test=8`, `backend=7`, `backend_test=2`, `docs=17`; affected App Layers `api=7`, `backend=1`, `mixed=1`. High scope is expected for shared MCP prompt dispatch, MCP resource text, AI-context generator, docs, and tests. |

## Phase 2 Graph Labeling And Visual Orientation Evidence

Status: active.

### P2-A Problem Screenshot Evidence

| Evidence | Result |
|---|---|
| `reports/problem/screenshot_1779517751.png` | File exists; `341,738` bytes; last modified `2026-05-26 14:54:54`. |
| Image dimensions | `1314x826`, `Format24bppRgb`. |
| Product issue recorded from screenshot | The graph view shows visible macro rings and node islands, but the rings and islands do not carry direct readable names on the graph canvas. Users must infer area meaning from color, spatial grouping, legend/filter state, or side-panel context instead of reading labels on the graph itself. |

### P2-B Web Graph Owner Trace

Fresh graph:

| Command | Result |
|---|---|
| `.\avmatrix\bin\avmatrix.exe analyze --force` | pass before Web graph trace; `files: scanned=772 parsed=575 unsupported=197 failed=0`, `nodes=87885 relationships=120605`. |

AVmatrix trace:

| Command | Finding |
|---|---|
| `.\avmatrix\bin\avmatrix.exe query "web graph layout ring island Sigma labels filters graph adapter" --repo AVmatrix --limit 10` | Top owners include `knowledgeGraphToGraphology`, `useSigma`, `filterGraphByLabels`, `applyFilterBasedClusteredLayout`, `graphHealthMatchesFilters`, `graphNodeMatchesHealthFilters`, `buildLayoutRingDiagnostics`, and `createLayoutRingBounds`. |
| `.\avmatrix\bin\avmatrix.exe context "GraphCanvas" --repo AVmatrix` | `GraphCanvas` lives in `avmatrix-web/src/components/GraphCanvas.tsx`; it is imported by `App.tsx` and an existing selection-performance unit test. |
| `.\avmatrix\bin\avmatrix.exe context "useSigma" --repo AVmatrix` | `useSigma` lives in `avmatrix-web/src/hooks/useSigma.ts`; it owns Sigma creation, label renderer settings, reducers, camera controls, graph replacement, and manual layout optimization. |
| `.\avmatrix\bin\avmatrix.exe context "filterGraphByLabels" --repo AVmatrix` | `filterGraphByLabels` lives in `avmatrix-web/src/lib/graph-adapter.ts`; it sets node `hidden` from node label, graph-health filters, and semantic filters. |
| `.\avmatrix\bin\avmatrix.exe context "useGraphState" --repo AVmatrix` | `useGraphState` lives in `avmatrix-web/src/hooks/app-state/graph.tsx`; it owns `visibleLabels`, `graphHealthFilters`, `semanticFilters`, depth filter state, selected node state, and graph-link visibility. |
| `.\avmatrix\bin\avmatrix.exe impact "GraphCanvas" --repo AVmatrix --direction upstream` | LOW; direct dependent `App.tsx`, depth-2 dependent `main.tsx`, no affected process listed. This is the edit blast-radius record for the Web overlay slice. |

Source owner map:

| Concern | Owner |
|---|---|
| Ring and island coordinate generation | `applyFilterBasedClusteredLayout` in `avmatrix-web/src/lib/graph-adapter.ts`; it groups nodes by `appLayerRing` and `islandKey`, places ring centers, island centers, and node spirals. |
| Node type colors and display labels | `getNodeColor`, `getNodeDisplayLabel`, `FILTERABLE_LABELS`, and related constants imported by `graph-adapter.ts`. |
| Per-node ring/island metadata | `getAppLayerRingKey`, `getNodeIslandKey`, and `applyFilterBasedClusteredLayout` write `appLayerRing`, `islandKey`, `appLayerRingCenterX`, and `appLayerRingCenterY` into Sigma node attributes. |
| Sigma rendering and camera projection | `useSigma` creates the Sigma renderer, sets label/camera options, replaces the graph, exposes `sigmaRef`, and refreshes after filtering and layout. |
| Overlay and diagnostics host | `GraphCanvas.tsx` hosts the Sigma container, top/selected-node overlays, zoom/layout controls, graph conversion, filter effects, and `buildLayoutRingDiagnostics`. |
| Filter/depth visibility | `filterGraphByLabels`, `filterGraphByDepth`, `graph-health-filters.ts`, `semantic-filters.ts`, and `useGraphState` determine which nodes have `hidden=true`. |

Design decision before edit:

- Keep coordinate generation in `graph-adapter.ts` unchanged for the first label slice.
- Add a pure orientation-label contract/helper so unit tests can verify ring/island label metadata without a browser.
- Render labels as a DOM overlay in `GraphCanvas` using Sigma `graphToViewport`, so browser tests can query `data-testid` and text content instead of relying only on pixels.

### P2-C Through H Implementation

Implementation:

| File | Change |
|---|---|
| `avmatrix-web/src/lib/graph-orientation-labels.ts` | Added the pure ring/island orientation label contract, ring/island label formatting, visible-node counting from non-hidden Sigma nodes, zoom presentation rules, viewport projection, clamping, collision-aware placement, and fallback text behavior for unknown/custom categories. |
| `avmatrix-web/src/components/GraphCanvas.tsx` | Added a DOM graph orientation label overlay backed by Sigma `graphToViewport`; recomputes after graph load, filter/depth/semantic/graph-health changes, camera updates, render, resize, and window resize. |
| `avmatrix-web/test/unit/graph-orientation-labels.test.ts` | Added deterministic unit coverage for label metadata, anchors, visible counts, hidden-node behavior, ResolutionGap/custom labels, zoom visibility, viewport clamping, and overlap guardrails. |
| `avmatrix-web/test/unit/GraphCanvas.selection-performance.test.tsx` | Extended the Sigma mock so the new overlay can recompute labels without changing the existing selection and manual-layout performance assertions. |
| `avmatrix-web/e2e/graph-orientation-labels.spec.ts` | Added mocked runtime graph e2e coverage that counts labels through selectors, checks representative ring/island text, validates filter updates, asserts label-overlap count is zero, and captures desktop/smaller-viewport screenshots. |

P2-F2 layout geometry verification:

| Source/test | Finding |
|---|---|
| `avmatrix-web/src/lib/graph-adapter.ts` lines around `getClusterNodeSpacing`, `getClusterIslandRadius`, `getIslandOffset`, `islandGap`, `largestAdjacentClusterSpan`, `ringGap`, and `largestAdjacentRingSpan` | Current layout already computes density-aware node spacing, island radius, spiral offsets, island gutter, adjacent island span, ring gutter, and adjacent ring span from cluster footprint instead of a fixed-size circle. No `graph-adapter.ts` edit was required for this slice. |
| `avmatrix-web/test/unit/graph-adapter.edge-geometry.test.ts` | Existing deterministic geometry tests cover separated node-type islands, imbalanced island gutters, dense large-repo pinwheel gutters, non-rail island shapes, App Layer ring ordering, ring/island grouping, ResolutionGap island separation, documentation island center placement, and deterministic spiral ordering. |

Impact and freshness:

| Command | Result |
|---|---|
| `.\avmatrix\bin\avmatrix.exe analyze --force` | pass before final placement edit; `files: scanned=775 parsed=578 unsupported=197 failed=0`, `nodes=88501 relationships=121393`. |
| `.\avmatrix\bin\avmatrix.exe impact --uid "Function:avmatrix-web/src/lib/graph-orientation-labels.ts:placeGraphOrientationLabels" --repo AVmatrix --direction upstream --include-tests` | CRITICAL; `impactedCount=7`, direct impacted nodes `4`, affected processes `12`, affected app layers `frontend=4` and `frontend_test=3`. This is blast-radius evidence for a shared Web label-placement helper, not an edit ban. |
| `.\avmatrix\bin\avmatrix.exe impact "GraphCanvas" --repo AVmatrix --direction upstream` | LOW; direct dependent `App.tsx`, depth-2 dependent `main.tsx`. This is the GraphCanvas overlay blast-radius record. |

Validation:

| Command | Result |
|---|---|
| `powershell -ExecutionPolicy Bypass -File avmatrix-launcher\build.ps1` | pass after final label-placement patch; Go build and Web production build completed. Existing Vite dynamic-import and chunk-size warnings only. |
| `cd avmatrix-web; npm run test -- graph-orientation-labels GraphCanvas.selection-performance graph-adapter.edge-geometry` | pass; `3` test files, `33` tests. |
| `cd avmatrix-web; npx playwright test graph-orientation-labels.spec.ts --project chromium` | pass; `2` Chromium e2e tests. |
| `cd avmatrix-web; npm run test` | pass after final label-placement patch; `46` test files, `377` tests. |
| Playwright DOM measurement script against the mocked orientation graph | pass; label counts, filter-update counts, and overlap counts recorded in the benchmark ledger. |

Screenshot artifacts:

| Artifact | Result |
|---|---|
| `avmatrix-web/test-results/graph-orientation-labels-G-5b603-labels-on-the-desktop-graph-chromium/graph-orientation-labels-desktop.png` | Captured by e2e after final placement patch. |
| `avmatrix-web/test-results/graph-orientation-labels-G-e5a0a-t-and-updates-after-filters-chromium/graph-orientation-labels-small-filtered.png` | Captured by e2e after final placement patch. |

Browser validation note:

- The Browser plugin skill was reviewed for local browser validation, but the expected Node REPL browser tool was not available through tool discovery in this session. The phase used repo Playwright e2e plus a Playwright DOM measurement script instead, which validates the same local runtime target and selector/screenshot behavior.

P2-H guidance decision:

- No user-facing instructional prose or label toggle was added. The UI change names the rings and islands directly on the graph.

Pre-commit checks:

| Command | Result |
|---|---|
| `git diff --check` | pass. |
| `.\avmatrix\bin\avmatrix.exe analyze --force` | pass after staging the P2 slice; `files: scanned=775 parsed=578 unsupported=197 failed=0`, `nodes=88580 relationships=121455`. |
| `.\avmatrix\bin\avmatrix.exe detect-changes --repo AVmatrix --scope all` | pass after staging the P2 slice; summary `changed_files=8`, `changed_count=737`, `affected_count=0`, `risk_level=low`; changed App Layers `frontend=434`, `frontend_test=292`, `docs=11`; changed Functional Areas `web_graph_ui=75`, `unknown=651`, `documentation=11`. |

## Phase 3 Embedded Skill Source Upgrade Evidence

Status: complete; pre-commit detect pending.

Fresh graph and owner trace:

| Command | Result |
|---|---|
| `.\avmatrix\bin\avmatrix.exe analyze --force` | pass before Phase 3 graph work; `files: scanned=775 parsed=578 unsupported=197 failed=0`, `nodes=88580 relationships=121455`. |
| `.\avmatrix\bin\avmatrix.exe query "embedded AVmatrix skills baseSkills AI context generated skills tests" --repo AVmatrix --limit 10` | Top owners included `installBaseSkills`, `baseSkillContent`, `GenerateAIContextFiles`, `renderAVmatrixBlock`, setup skill installation owners, and `internal/aicontext/aicontext_test.go`. |
| `.\avmatrix\bin\avmatrix.exe context "GenerateAIContextFiles" --repo AVmatrix` | Confirmed `GenerateAIContextFiles` writes root context, installs base skills, removes generated community skills, and is reached from analyze post-run. |
| `.\avmatrix\bin\avmatrix.exe context "baseSkillContent" --repo AVmatrix` | Confirmed embedded Markdown is loaded from `internal/aicontext/skills/<skill>.md`, with `fallbackBaseSkillContent` only as defensive behavior. |

Impact:

| Command | Result |
|---|---|
| `.\avmatrix\bin\avmatrix.exe impact "renderAVmatrixBlock" --repo AVmatrix --direction upstream --include-tests` | CRITICAL; `impactedCount=7`, direct impacted nodes `1`, affected processes `12`; affected App Layers `backend=4`, `backend_test=3`. This is blast-radius evidence for generated root context output, not an edit ban. |
| `.\avmatrix\bin\avmatrix.exe impact --uid "Variable:internal/aicontext/aicontext.go:baseSkills" --repo AVmatrix --direction upstream --include-tests` | LOW; no impacted processes. |
| `.\avmatrix\bin\avmatrix.exe impact "TestGenerateAIContextFilesCreatesAndUpdatesManagedContext" --repo AVmatrix --direction upstream --include-tests` | LOW; no impacted processes. |

Implementation:

| File | Change |
|---|---|
| `internal/aicontext/aicontext.go` | Expanded `baseSkills` from 6 to 11, added per-skill generated table task text, and made the generated Skills table iterate over `baseSkills` to prevent registry/table drift. |
| `internal/aicontext/skills/avmatrix-exploring.md` | Rewritten as a practical exploration guide with command choices, multi-lane query guidance, context verification, validation, and limitations. |
| `internal/aicontext/skills/avmatrix-impact-analysis.md` | Rewritten with impact workflow, command choices, risk interpretation, and explicit HIGH/CRITICAL blast-radius wording. |
| `internal/aicontext/skills/avmatrix-debugging.md` | Rewritten with symptom-to-owner workflow, graph-quality command choices, query reliability rule, and evidence expectations. |
| `internal/aicontext/skills/avmatrix-refactoring.md` | Rewritten with context/impact/rename/API-impact workflow and behavior-preserving refactor rules. |
| `internal/aicontext/skills/avmatrix-guide.md` | Rewritten as a unified CLI/MCP/resource/prompt/Web/API guide, with surface-correct command naming. |
| `internal/aicontext/skills/avmatrix-cli.md` | Rewritten as a current CLI command guide covering core graph navigation, API parity, graph quality, runtime/setup/package/groups/wiki/version surfaces. |
| `internal/aicontext/skills/avmatrix-graph-quality.md` | Added graph-health, query-health, resolution-inventory, source-site-accuracy, benchmark-compare, query reliability, and evidence guidance. |
| `internal/aicontext/skills/avmatrix-api-surface.md` | Added route/tool/shape/API-impact guidance with CLI-vs-MCP command naming rules. |
| `internal/aicontext/skills/avmatrix-cross-repo.md` | Added group status/sync/contracts/query guidance and cross-repo validation rules. |
| `internal/aicontext/skills/avmatrix-runtime-packaging.md` | Added serve/mcp/setup/package/launcher/canonical executable/process lifecycle guidance. |
| `internal/aicontext/skills/avmatrix-ai-context.md` | Added source-vs-generated rules for `AGENTS.md`, `CLAUDE.md`, `.claude/skills/avmatrix/**`, embedded skill source, and validation. |
| `internal/aicontext/aicontext_test.go` | Added final skill id assertions, generated Skills table coverage, embedded frontmatter/source-content tests, expanded command-surface regression checks, and surface-correct command naming tests. |

Validation:

| Command | Result |
|---|---|
| `powershell -ExecutionPolicy Bypass -File avmatrix-launcher\build.ps1` | pass after final skill wording; Go build and Web production build completed. Existing Vite dynamic-import and chunk-size warnings only. |
| `go test .\internal\aicontext .\internal\cli -count=1` | pass; `internal/aicontext` 0.695s, `internal/cli` 8.478s. |

Intermediate validation finding:

- The first focused Phase 3 test run failed because a skill used an example that contained the forbidden invented CLI spelling `avmatrix route_map`. The wording was changed to forbid MCP underscore names as top-level AVmatrix CLI commands without including a fake command example, and the final focused test passed.

Pre-commit checks:

| Command | Result |
|---|---|
| `git diff --check` | pass. |
| `.\avmatrix\bin\avmatrix.exe analyze --force` | pass after staging the P3 slice; `files: scanned=780 parsed=578 unsupported=202 failed=0`, `nodes=88644 relationships=121551`. |
| `.\avmatrix\bin\avmatrix.exe detect-changes --repo AVmatrix --scope all` | pass after staging the P3 slice; summary `changed_files=16`, `changed_count=158`, `affected_count=1`, `risk_level=medium`; changed App Layers `backend=7`, `backend_test=77`, `docs=74`; affected App Layers `backend=1`. |

## Phase 4 Setup, Package, And Documentation Integration Evidence

Status: complete; pre-commit detect pending.

Fresh graph and owner trace:

| Command | Result |
|---|---|
| `.\avmatrix\bin\avmatrix.exe analyze --force` | pass before Phase 4 graph work; `files: scanned=780 parsed=578 unsupported=202 failed=0`, `nodes=88644 relationships=121551`. |
| `.\avmatrix\bin\avmatrix.exe query "setup editor skills package embedded aicontext skills distribution" --repo AVmatrix --limit 10` | Top owners included `installBaseSkills`, `setupInstallEditorSkills`, `baseSkillContent`, `setupInstallSkillsTo`, `setupSkillTargetName`, `setupResource`, and package lifecycle functions. |
| `.\avmatrix\bin\avmatrix.exe context "setupInstallSkillsTo" --repo AVmatrix` | Confirmed the pre-change setup path read package-root `skills/` through `setupResolvePackagePath("skills")`. |
| `.\avmatrix\bin\avmatrix.exe context "setupInstallEditorSkills" --repo AVmatrix` | Confirmed editor setup callers for Cursor, Claude Code, OpenCode, and Codex all flow through `setupInstallSkillsTo`. |
| `.\avmatrix\bin\avmatrix.exe context "resolvePackageSourceRoot" --repo AVmatrix` | Confirmed package fallback builds use `go-src` when source is not available beside the package root. |

Impact:

| Target | Result |
|---|---|
| `setupInstallSkillsTo` | CRITICAL; `impactedCount=9`, affected App Layers `backend=6`, `backend_test=3`, affected Functional Area `cli=9`, affected setup processes `5`. This is setup/editor install blast-radius evidence, not an edit ban. |
| `setupInstallEditorSkills` | CRITICAL; `impactedCount=9`, direct impacted nodes `4`, affected setup processes `5`. This covers all supported editor setup callers. |
| `baseSkillContent` | HIGH; `impactedCount=9`, affected App Layers `backend=4`, `backend_test=5`; this is canonical embedded skill content loading. |
| `installBaseSkills` | CRITICAL; `impactedCount=7`, affected analyze AI-context processes `12`; this is analyze-generated root context/base skill output. |
| `setupResource` | LOW; no upstream impacted nodes; this is MCP setup/onboarding text. |
| `prepareGoSourcePackage` | CRITICAL; `impactedCount=6`, affected package/root CLI paths and `18` processes; this is package fallback source assembly. |
| `copyPackageGoDir` | CRITICAL; `impactedCount=4`, affected package source-copy path and `18` processes; this is package fallback source file selection. |
| Removed setup package-root helper path | `setupSkillTargetName`, `setupResolvePackagePath`, `setupCopyDir`, and `setupFileExists` were CRITICAL through the old `setupInstallSkillsTo` package-root flow; `setupCopyFile`, `setupAncestorRoots`, and `setupPathExists` were LOW. They were removed because setup no longer reads package-root skills. |
| `TestAnalyzeCommandGeneratesAIContextByDefault` | LOW; no upstream impacted nodes; test was tightened to assert all final embedded base skills are generated, not only `avmatrix-cli`. |

Implementation:

| File | Change |
|---|---|
| `internal/aicontext/aicontext.go` | Added exported `BaseSkillFile`, `BaseSkillFiles`, and `InstallBaseSkillsTo` so setup and tests materialize the same embedded base skill source as analyze. Added retired `avmatrix-pr-review` cleanup and made `installBaseSkills` delegate through the shared installer after cleaning the generated AVmatrix skill root. |
| `internal/cli/setup_command.go` | Replaced package-root `skills/` discovery/copy behavior with `aicontext.InstallBaseSkillsTo`. Removed the package-root skill target/copy/path helpers so editor setup cannot drift from embedded skill source. |
| `internal/cli/package_runtime.go` | Updated package fallback source assembly so `copyPackageGoDir` also copies `internal/aicontext/skills/*.md`, preserving `go:embed` inputs in `go-src`. |
| `internal/cli/command_test.go` | Added setup assertions that Cursor, Claude Code, OpenCode, and Codex receive all embedded base skills with exact content. Replaced the old package-root install test with a test proving package-root `flat-skill`, directory skills, and retired `avmatrix-pr-review` are not installed. Tightened default analyze coverage to assert all final generated base skill files. |
| `internal/cli/package_command_test.go` | Added package fallback coverage for embedded skill Markdown in `go-src`; updated package metadata expectations so `skills` is no longer shipped as a package-root authority. |
| `internal/aicontext/aicontext_test.go` | Added `BaseSkillFiles` inventory and metadata assertions. |
| `internal/mcp/resources.go` and `internal/mcp/resources_parity_test.go` | Updated `avmatrix://setup` guidance to explain generated AI context, embedded skill source, setup-installed embedded skills, and the non-authority status of package-root `skills/`. |
| `README.md` and `avmatrix/README.md` | Removed stale `analyze --skills`/old four-skill/package-root assumptions and documented generated `AGENTS.md`, `CLAUDE.md`, `.claude/skills/avmatrix/**`, embedded source ownership, and setup-installed embedded base skills. |
| `avmatrix/package.json` and `avmatrix/skills/*.md` | Removed `skills` from npm package files and deleted the stale package-root skill source files, including retired `avmatrix-pr-review`. |
| `docs/query-health/2026-05-23-avmatrix-skill-system-upgrade-suite.json` | Updated setup/package owner expectations from package-root skill helpers to embedded skill owners and package source-copy owners. |
| `internal/aicontext/skills/avmatrix-runtime-packaging.md` | Clarified that setup/package skill validation must use embedded skill source and that package-root `skills/` is not a source of truth. |

Validation:

| Command | Result |
|---|---|
| `powershell -ExecutionPolicy Bypass -File avmatrix-launcher\build.ps1` | pass after Phase 4 implementation; Go build and Web production build completed. Existing Vite dynamic-import and chunk-size warnings only. |
| `go test .\internal\aicontext .\internal\cli .\internal\mcp -count=1` | pass after the final analyze/setup/package test tightening; `internal/aicontext` 1.294s, `internal/cli` 12.076s, `internal/mcp` 6.530s. |
| Built binary setup smoke with temp `HOME`/`USERPROFILE`, temp Cursor/Claude/OpenCode/Codex dirs, and isolated `AVMATRIX_HOME` | pass; setup reported 11 skills for Cursor, Claude Code, OpenCode, and Codex. Each checked editor target contained 11 embedded skill directories totaling 30,984 bytes, with `avmatrix-pr-review` absent. |
| `avmatrix\bin\avmatrix.exe package prepare-go-source` from `avmatrix/`, followed by inventory and `package clean-go-source` | pass; prepare copied 293 files, inventory found 278 Go files and 11 embedded skill Markdown files under `go-src/internal/aicontext/skills`, totaling 30,984 bytes, then cleanup removed `go-src`. |
| Package source/hash comparison during temporary `go-src` preparation | pass; source skill Markdown count 11, prepared skill Markdown count 11, name mismatches 0, SHA-256 hash mismatches 0, prepared skill bytes 30,984. |
| `npm pack --dry-run --ignore-scripts` from `avmatrix/` | pass for skill distribution check; package dry-run listed 7 tarball files and no package-root `skills/` files. Existing bin backup artifacts remain in the dry-run inventory and are outside this skill-source reconciliation slice. |

Validation note:

- The first setup smoke script attempt failed before invoking AVmatrix because the PowerShell script tried to overwrite the read-only `$HOME` variable. The rerun used a separate `$tempHome` variable and passed.

Pre-commit checks:

| Command | Result |
|---|---|
| `git diff --check` | pass. |
| `.\avmatrix\bin\avmatrix.exe analyze --force` | pass after Phase 4 implementation and ledger updates; `files: scanned=773 parsed=578 unsupported=195 failed=0`, `nodes=88582 relationships=121473`. |
| `.\avmatrix\bin\avmatrix.exe detect-changes --repo AVmatrix --scope all` | pass after Phase 4 implementation and ledger updates; summary `changed_files=23`, `changed_count=145`, `affected_count=1`, `risk_level=medium`; changed App Layers `api=3`, `api_test=2`, `backend=55`, `backend_test=69`, `docs=16`; affected App Layers `api=1`; affected Functional Areas `mcp=1`. |

## Phase 5 Regeneration And Validation Evidence

Status: in progress.

Validation:

| Task | Command | Result |
|---|---|---|
| P5-A full build gate | `powershell -ExecutionPolicy Bypass -File avmatrix-launcher\build.ps1` | pass; Go build used `go version go1.26.3 windows/amd64`, native runtime path `.tmp\ladybug-native\v0.16.1\windows-x86_64`, Web production build transformed 2,931 modules and built in 21.34s. Existing Vite dynamic-import and chunk-size warnings only. |
| P5-B focused backend/CLI/MCP tests | `go test .\internal\mcp .\internal\cli .\internal\aicontext .\internal\graphhealth .\internal\httpapi -count=1` | pass; `internal/mcp` 8.579s, `internal/cli` 12.099s, `internal/aicontext` 1.270s, `internal/graphhealth` 1.132s, `internal/httpapi` 3.629s. |
| P5-C Web unit tests | `cd avmatrix-web; npm run test` | pass; 46 test files passed, 377 tests passed. Focused graph orientation label coverage included `test/unit/graph-orientation-labels.test.ts` with 7 tests passed. |
| P5-D Web e2e/browser validation | `cd avmatrix-web; npm run test:e2e` with Vite started directly on `127.0.0.1:5228` for the test run | pass for frontend/harness full suite; 16 passed, 31 skipped, exit code 0. Graph-health fixture specs passed 4/4 and graph-orientation label specs passed 2/2. Cleanup check reported `remaining_listeners=0` for ports 5228 and 4848. |
| P5-E normal AI-context generation | `.\avmatrix\bin\avmatrix.exe analyze --force` | pass with no `--skip-agents-md`; analyzed `E:\AVmatrix-GO`, `files: scanned=773 parsed=578 unsupported=195 failed=0`, `nodes=88583 relationships=121474`. |
| P5-F generated/source skill inventory comparison | PowerShell inventory comparing `internal\aicontext\skills\*.md` to `.claude\skills\avmatrix\*\SKILL.md` | pass; source count 11, generated count 11, name diff 0, SHA-256 mismatch 0, source bytes 30,984, generated bytes 30,984, source lines 556, generated lines 556. `AGENTS.md` and `CLAUDE.md` each contained 11 skill table rows, 28 command-selection rows, and all final skill paths. |
| P5-G setup smoke | Isolated `HOME`, `USERPROFILE`, and `AVMATRIX_HOME`, then `.\avmatrix\bin\avmatrix.exe setup` | pass; setup output configured Cursor, Claude Code, OpenCode, and Codex fallback, installed 11 skills to each supported target, added Claude Go hooks, preserved one Codex MCP fallback section, created isolated `AVMATRIX_HOME`, and left legacy `avmatrix-hook.cjs` absent. |
| P5-G setup installed skill inventory | PowerShell inventory of `.cursor\skills`, `.claude\skills`, `.config\opencode\skill`, and `.agents\skills` under the isolated home | pass; each target had 11 `SKILL.md` files, 30,984 bytes, SHA-256 mismatches 0, and retired `avmatrix-pr-review` absent. |
| P5-G package fallback source smoke | From `avmatrix\`, `.\bin\avmatrix.exe package prepare-go-source`, inventory/hash check, `npm pack --dry-run --ignore-scripts`, and `.\bin\avmatrix.exe package clean-go-source` | pass; prepare output copied 293 files to `avmatrix\go-src`; inventory found 294 files including 278 Go files and 11 embedded skill Markdown files totaling 30,984 bytes; source/prepared skill name diff 0 and SHA-256 mismatch 0; dry-run tarball total files 301 with no package-root `skills/`; cleanup removed `go-src`. |
| P5-H MCP setup resource smoke | JSON-RPC stdio `initialize`, `notifications/initialized`, then `resources/read` for `avmatrix://setup` | pass; response advertised server `avmatrix` version `1.2.2`, included `avmatrix://setup`, and the text contained `AI Context And Skills`, `internal/aicontext/skills/*.md`, `AGENTS.md`, `CLAUDE.md`, `.claude/skills/avmatrix/**`, `avmatrix setup` installed skills, and `package-root skills/ is not a source of truth`. |
| P5-I dedicated query-health suite | `.\avmatrix\bin\avmatrix.exe query-health --repo AVmatrix --suite docs\query-health\2026-05-23-avmatrix-skill-system-upgrade-suite.json --limit 10 --out docs\query-health\2026-05-26-avmatrix-skill-system-upgrade-p5-query-health.json` | pass by threshold; report written to `docs\query-health\2026-05-26-avmatrix-skill-system-upgrade-p5-query-health.json`. Summary: 9 cases, threshold passed 9, threshold failed 0, exact passed 4, exact failed 5, matched targets 54/63, missed targets 9. Matched query lanes: `owner_discovery`, `command_surface_discovery`, `api_surface_discovery`, `graph_quality_discovery`, `docs_setup_ai_context_discovery`, `cross_repo_discovery`, `execution_flow_discovery`, `concept_discovery`. Top-result sources covered `definition`, `process`, and `process_symbol`; matched target global ranks ranged 1-41 and source ranks 1-9. |

P5-I exact misses:

| Case | Noise reason |
|---|---|
| `ai-context-generated-skills-owner-discovery` | thresholds met; exact target miss `file:internal/aicontext/skills/avmatrix-cli.md`. |
| `package-skill-distribution-owner-discovery` | thresholds met; exact target misses `file:internal/cli/package_runtime.go`, `symbol:prepareGoSourcePackage`, `symbol:copyPackageGoDir`. |
| `mcp-setup-resource-prompt-guidance-owner-discovery` | thresholds met; exact target miss `symbol:mcpTools`. |
| `query-command-surface-owner-discovery` | thresholds met; exact target misses `symbol:queryTool`, `symbol:rankedProcessMatches`. |
| `graph-quality-command-surface-owner-discovery` | thresholds met; exact target misses `file:internal/graphhealth/compute.go`, `symbol:Compute`. |

P5-K broad-discovery regression:

| Command | Result |
|---|---|
| `.\avmatrix\bin\avmatrix.exe query "group query execution flow process sync contracts" --repo AVmatrix --limit 10 --explain --json` | pass; after parsing the JSON portion before the explanatory footer, output contained 10 definitions, 10 process candidates, and 10 definitions with match reasons. The first definition names were `Query`, `Sync`, `groupQueryTool`, `Contracts`, and `newGroupCommand`. Process labels included `Sync -> NormalizeHTTPPath`, `Query -> GroupProcess`, and `Query -> GroupProcessStep`. Definition query lanes included `execution_flow_discovery`, `concept_discovery`, `cross_repo_discovery`, `api_surface_discovery`, `graph_quality_discovery`, `owner_discovery`, and `command_surface_discovery`. |

P5-K2 graph-health CLI smoke:

| Command | Result |
|---|---|
| `.\avmatrix\bin\avmatrix.exe graph-health --help` | pass; help listed `summary`, `report`, `components`, and `explain`. |
| `.\avmatrix\bin\avmatrix.exe graph-health summary --repo AVmatrix --json` | pass; JSON contained `inputs`, `totals`, and `summary`; totals were 88,583 nodes and 121,474 relationships; summary included policy `graph-health-non-structural-v1`, 80,216 components, 64 detached components, topology counts, diagnostic counts, resolution bucket counts, app-layer counts, and functional-area counts. |
| `.\avmatrix\bin\avmatrix.exe graph-health report --repo AVmatrix --limit 5 --json` | pass; returned 5 candidates. First candidate was `File:avmatrix-web/src/components/RightPanel.tsx`, triage priority `no_incoming`, topology status `no_incoming`, component `component_000001`, component size 8,062. |
| `.\avmatrix\bin\avmatrix.exe graph-health components --repo AVmatrix --limit 5 --json` | pass; returned 5 components. First component `component_000936` had 21 nodes, 53 counted edges, detached `true`, reachableFromRoot `false`, unresolved reference diagnostics 86, and resolution health bucket counts. |
| `.\avmatrix\bin\avmatrix.exe graph-health explain "File:avmatrix-web/src/components/RightPanel.tsx" --repo AVmatrix --json` | pass; output contained `inputs`, `totals`, `kind`, `nodeId`, `componentId`, `node`, `health`, `countedOutgoingRelationships`, and `excludedRelationships`. |
| `.\avmatrix\bin\avmatrix.exe graph-health explain "missing-node-for-p5-k2" --repo AVmatrix --json` | expected failure; exit code 1 with `graph node not found: missing-node-for-p5-k2`. |

P5-K3 CLI parity smoke:

| Command | Result |
|---|---|
| `.\avmatrix\bin\avmatrix.exe --help` | pass; visible help includes `api`, `rename`, `graph-health`, `query`, `query-health`, `resolution-inventory`, `source-site-accuracy`, `setup`, `serve`, `mcp`, `group`, `wiki`, and `wiki-mode`. Root help does not list `package`, `hook`, `grep`, `processes`, or `clusters`. |
| `.\avmatrix\bin\avmatrix.exe api --help` | pass; help lists 4 accepted API parity subcommands: `impact`, `route-map`, `shape-check`, and `tool-map`. |
| `.\avmatrix\bin\avmatrix.exe rename --help` | pass; help exposes dry-run rename flags including `--apply`, `--file`, `--json`, `--repo`, and `--uid`. |
| `.\avmatrix\bin\avmatrix.exe api route-map "/api/users" --repo p1-6-cli-parity-smoke-20260526173638 --json` | pass; JSON parse confirmed `total=1`, route `/api/users`. |
| `.\avmatrix\bin\avmatrix.exe api tool-map query --repo p1-6-cli-parity-smoke-20260526173638 --json` | pass; JSON parse confirmed `total=1`, tool `query`. |
| `.\avmatrix\bin\avmatrix.exe api shape-check "/api/users" --repo p1-6-cli-parity-smoke-20260526173638 --json` | pass; JSON parse confirmed `total=1`, `mismatches=1`. |
| `.\avmatrix\bin\avmatrix.exe api impact "/api/users" --repo p1-6-cli-parity-smoke-20260526173638 --json` | pass; JSON parse confirmed route `/api/users`, handler `app/api/users/route.ts`, `directConsumers=1`, `affectedFlows=0`, `riskLevel=MEDIUM`, and one response-shape mismatch. |
| `.\avmatrix\bin\avmatrix.exe rename loadUsers loadUsers2 --repo p1-6-cli-parity-smoke-20260526173638 --json` | pass; dry-run JSON parse confirmed `status=success`, `applied=false`, `files_affected=1`, `total_edits=1`, `graph_edits=1`, `text_search_edits=0`. |
| `rg -n 'Hidden:\s*true|Use:\s+"(package|hook)' internal\cli -g '*.go'` | pass; found 9 source lines confirming `hook` and `package` families and package subcommands remain hidden lifecycle-only surfaces. |

P5-K4 MCP prompt smoke:

| Command | Result |
|---|---|
| JSON-RPC `prompts/list` | pass; returned 2 prompts: `detect_impact` and `generate_map`. `detect_impact` description says it analyzes current changes using change detection, context, impact, and freshness rules. `generate_map` description says it generates evidence-backed architecture documentation from resources and graph facts. |
| JSON-RPC `prompts/get` for `generate_map` with no arguments | pass; returned one user message. Text instructs repo selection through `avmatrix://repos`, URL-escaped repo resource URIs, freshness check with `avmatrix analyze --force` if stale, reading context/clusters/processes/process details, optional use of `context`, `impact`, API CLI commands, `graph-health`, `query-health`, `resolution-inventory`, and `cypher`, and output rules to write architecture claims only from resources, MCP tools, CLI commands, or Web/API output actually read. |
| JSON-RPC `prompts/get` for `generate_map` with `repo=AVmatrix` | pass; returned one user message with supplied repo `AVmatrix`, resolved resource URI segment `AVmatrix`, freshness instruction, evidence workflow, and evidence-only output rules. |
| JSON-RPC `prompts/get` for `detect_impact` | pass; returned one user message. Text instructs obeying repository freshness rules, MCP `detect_changes`, CLI fallback `avmatrix detect-changes`, MCP/CLI `context`, MCP/CLI `impact`, HIGH/CRITICAL as blast-radius warnings not edit bans, impact before implementation edits, and detect-changes before committing. |

P5-L pre-commit checks:

| Command | Result |
|---|---|
| `git diff --check` | pass. |
| `.\avmatrix\bin\avmatrix.exe analyze --force` | pass after Phase 5 checklist, evidence, benchmark, and query-health report updates; `files: scanned=774 parsed=578 unsupported=196 failed=0`, `nodes=88585 relationships=121476`. |
| `.\avmatrix\bin\avmatrix.exe detect-changes --repo AVmatrix --scope all` | pass; summary `changed_files=3`, `changed_count=10`, `affected_count=0`, `risk_level=low`; changed App Layers `docs=10`; changed Functional Areas `documentation=10`; no affected processes; semantic status complete for App Layer and Functional Area fields. |

## Phase 6 Zero-Trust Closure Review Evidence

Status: in progress.

P6-A stale-assumption review:

| Check | Result |
|---|---|
| Search active docs/source for six-skill-only, package-root skills, direct-edit generated output, stale CLI names, and color/filter-only graph guidance | pass with one active-doc fix. Historical plan/evidence ledgers still mention old six-skill/package-root states as baseline or superseded evidence and were intentionally left unchanged. |
| `README.md` | Current root README already listed final direct graph/API/quality command surfaces, generated AI-context source rules, and no direct-edit generated-output guidance. |
| `internal/aicontext/skills/*.md` and generated context source | Current embedded skills and generator describe 11 skills, CLI/MCP name distinctions, graph-health/query-health/resolution/source-site commands, and source-vs-generated rules. |
| `avmatrix/README.md` | Updated packaged README CLI section to include final setup/analyze/index/list/status/serve/mcp/version/wiki-mode commands, direct graph/API/quality commands including `rename`, `api route-map`, `api tool-map`, `api shape-check`, `api impact`, `graph-health`, `query-health`, `resolution-inventory`, `source-site-accuracy`, and `benchmark-compare`, plus notes that `graph-health` audits topology/diagnostics and `rename`/`api ...` delegate to the same MCP owner logic. |

P5-D screenshot evidence:

| Screenshot | Result |
|---|---|
| `avmatrix-web\test-results\graph-orientation-labels-G-5b603-labels-on-the-desktop-graph-chromium\graph-orientation-labels-desktop.png` | captured after desktop graph label assertions; 237,028 bytes. |
| `avmatrix-web\test-results\graph-orientation-labels-G-e5a0a-t-and-updates-after-filters-chromium\graph-orientation-labels-small-filtered.png` | captured after small viewport/filter assertions; 92,005 bytes. |

P5-D validation note:

- First `npm run test:e2e` attempt failed 16 tests because no frontend server was listening on `127.0.0.1:5228`; every failure was `ERR_CONNECTION_REFUSED`.
- A real-backend full e2e attempt with both Vite `5228` and `avmatrix serve` `4848` reached the app and passed the graph-label fixture specs, but failed 9 real-backend specs on very large indexed repos: readiness timeouts while loading 453-491 MB graph payloads and one diagnostic expectation on a non-AVmatrix repository. That run reported 37 passed, 1 skipped, 9 failed and `remaining_listeners=0` after cleanup. The accepted P5-D evidence is the full frontend/harness Playwright run above, which executes the label UI assertions and lets real-backend specs skip when no live backend is present.

P6-B final validation after closure review:

| Task | Command | Result |
|---|---|---|
| Full build gate | `powershell -ExecutionPolicy Bypass -File avmatrix-launcher\build.ps1` | pass; Go build and Web production build completed. Web build transformed 2,931 modules and built in 18.71s. Existing Vite dynamic-import and chunk-size warnings only. |
| Focused backend/CLI/MCP tests | `go test .\internal\mcp .\internal\cli .\internal\aicontext .\internal\graphhealth .\internal\httpapi -count=1` | pass; `internal/mcp` 9.697s, `internal/cli` 14.096s, `internal/aicontext` 1.751s, `internal/graphhealth` 1.111s, `internal/httpapi` 4.558s. |
| Web unit tests | `cd avmatrix-web; npm run test` | pass; 46 test files passed, 377 tests passed, duration 32.24s. |
| Web e2e/browser validation | `cd avmatrix-web; npm run test:e2e` with Vite started directly on `127.0.0.1:5228` | pass for frontend/harness full suite; 16 passed, 31 skipped, 0 failed, duration 56.4s. Cleanup check reported `remaining_listeners=0` for ports 5228 and 4848. |
| Normal AI-context generation | `.\avmatrix\bin\avmatrix.exe analyze --force` | pass with no `--skip-agents-md`; analyzed `E:\AVmatrix-GO`, `files: scanned=774 parsed=578 unsupported=196 failed=0`, `nodes=88584 relationships=121475`. |
| Generated/source skill inventory comparison | PowerShell inventory comparing `internal\aicontext\skills\*.md` to `.claude\skills\avmatrix\*\SKILL.md` | pass; source count 11, generated count 11, name diff 0, SHA-256 mismatch 0, source bytes 30,984, generated bytes 30,984, source lines 556, generated lines 556. `AGENTS.md` and `CLAUDE.md` each contained 11 skill table rows and 44 command-selection rows. |
| Package dry-run inventory after packaged README update | `cd avmatrix; npm pack --dry-run --ignore-scripts --json` | pass; package `avmatrix@1.2.2`, tarball `avmatrix-1.2.2.tgz`, total files 7, `README.md` included with size 18,635 bytes, package-root `skills/` files 0. Files were `README.md`, `bin/avmatrix-runtime.json`, `bin/avmatrix.exe`, `bin/avmatrix.exe~`, `bin/lbug_shared.dll`, `bin/lbug_shared.dll.locked-20260520202525`, and `package.json`. |
| Web label screenshots after final e2e run | `Get-ChildItem avmatrix-web\test-results -Recurse -Filter graph-orientation-labels-*.png` | pass; desktop screenshot `graph-orientation-labels-desktop.png` was 237,028 bytes, small filtered screenshot `graph-orientation-labels-small-filtered.png` was 92,012 bytes. |

P6-B remaining limitation:

- The accepted final e2e evidence remains the frontend/harness full suite because the real-backend full suite is not deterministic on the current machine with very large indexed repository payloads. The label fixture specs pass in the accepted run, screenshots are refreshed, and process cleanup reports no remaining listeners.

P6-C closure state:

| Check | Result |
|---|---|
| Source and generated output agreement | pass; source embedded skills, generated `.claude/skills/avmatrix/**`, generated `AGENTS.md`, and generated `CLAUDE.md` inventories agree on 11 final skills and current command guidance. |
| Validation agreement | pass; P6-B repeated full build, focused backend/CLI/MCP tests, Web unit tests, Web e2e/browser validation, normal generation, package dry-run inventory, and refreshed Web label screenshots. |
| Prior command-surface validations | pass; P5 recorded graph-health CLI validation, CLI parity validation, MCP prompt validation, Web graph label validation, and the dedicated AI-context query-health threshold/exact results. |
| Benchmark and evidence ledgers | pass; P6-B records final evidence and B8 records final closure-review inventory metrics. |
| Commit-state check | pass; `git diff --check` returned clean, `git diff --stat` showed only 4 doc/README files changed. AVmatrix `detect-changes` was not run for this closure commit because it is doc-only and repository rule 6 says not to use AVmatrix for doc-only commits. |
