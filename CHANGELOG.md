# Changelog

All notable changes to anvien will be documented in this file.

## [1.2.6] - 2026-06-06

### Added

- Added the UI backend-binding skill package and expanded its execution rules for binding backend/API data into approved UI without changing approved layout, copy, visual design, component hierarchy, or interaction states.
- Added `anvien skill --help` so agents can read the generated Skill Selection Guide from the CLI, matching the generated `AGENTS.md` and `CLAUDE.md` skill routing content.
- Added a root `npm run full-build` script that runs the repository full-build sequence from `RUNBOOK.md`.

### Changed

- Generated master rules now include the Skill Selection Guide alongside the command selection guidance, keeping CLI help and generated agent context aligned.
- Renamed the active structured file-detail command and local API route from `file-context` to `file-detail` without keeping a legacy alias.
- Generated agent context now includes the planner-skill rule for real `docs/plans` plan creation and clarifies rule 4 as `Write plan (use planner skill) before coding`.
- Updated Coder workflow/reporting guidance for clearer process adherence and evidence expectations.
- Bumped the CLI package version from `1.2.5` to `1.2.6`.

### Fixed

- Fixed package runtime source staging so the skill subtree is included in the packaged runtime source fallback.
- Fixed generated rule help wording for the Anvien usage guidance.

### Removed

- Removed Repomix usage-pattern documentation and associated scripts from the generated skill surface.

### Documented

- Updated full-build instructions in `README.md` and `RUNBOOK.md`, including the new `scripts\full-build.ps1` entrypoint.
- Recorded UI backend-binding wiring audit and plan alignment notes for the new skill package.

## [1.2.5] - 2026-06-05

### Added

- Added recursive, namespace-preserving AI-context skill package installation so top-level packages under `internal/aicontext/skills/**` can include nested `SKILL.md` entries, scripts, references, templates, and assets without flattening or dropping package payload.
- Added deterministic skill package hashing, manifest ownership, and incremental file-level sync for generated skill output, including edit/add/delete/rename propagation, tamper repair, missing-output repair, stale-output deletion, and sync counters.
- Added surface-specific generated skill layouts for Codex and Claude Code: `AGENTS.md` now points to `.agents/skills/<package>/...`, `CLAUDE.md` points to `.claude/skills/<package>/...`, and analyze installs both generated skill surfaces.
- Added scoped generated-skill ownership handling for direct skill roots so Anvien-managed package roots can sync while unrelated repo-local custom skills remain protected.
- Added first-class file group metadata for backend support/model/helper files, including the `backend_support_model_helper` group key and `Backend support/model/helper files` label across file projection, CLI/API output, contracts, and Web file views.
- Added a PostgreSQL best-practices database skill reference covering foreign-key indexes, JOIN support indexes, partial indexes, benchmark examples, and production guardrails.

### Changed

- Bumped the CLI package version from `1.2.4` to `1.2.5`.
- Separated generated Anvien command routing from generated skill routing by keeping direct CLI/MCP command selection in the `Command Selection Guide` and routing workflow skills through a concise `Skill Selection Guide`.
- Shortened generated skill descriptions to trigger-only wording, removed the rejected `ai-multimodal` package from the generated catalog, and normalized multi-entry package routing for document and problem-solving skills.
- Made generated `.agents/skills/**` and `.claude/skills/**` output an exact projection of `internal/aicontext/skills/**` for Anvien-managed packages while removing the old generated `skills/anvien/` namespace layer.
- Removed the volatile indexed-project inventory sentence from generated agent context so ordinary `AGENTS.md` and `CLAUDE.md` output remains stable and repo-agnostic.
- Updated default unresolved, hotspot, risk, and Web file-display behavior so test files remain visible as test files with tested-target relationships while test-source unresolved details no longer dominate default product signals.

### Fixed

- Fixed generated skill installation missing nested skill entries and associated scripts/resources from multi-entry packages.
- Fixed stale generated skill output surviving after source package or file deletion.
- Fixed Codex-facing generated guidance pointing to Claude-shaped skill paths.
- Fixed default unresolved hotspot lists being dominated by test and e2e source unresolved details.
- Fixed backend support/model/helper files being identifiable only by unresolved-count differences instead of by a direct file group label.

### Documented

- Recorded an evidence-backed Anvien versus GitNexus deep comparison with benchmark, accuracy, feature, and maturity findings.
- Recorded benchmark ledgers for skill inventory reduction, recursive package installation, incremental skill mirror sync, generated skill layout split, file group counts, test-file unresolved separation, and PostgreSQL reference inventory.

## [1.2.4] - 2026-06-01

### 2026-06-01

#### Added

- Added a file-centric graph projection shared by CLI, MCP, API, and Web surfaces for file summaries, symbol trees, file relationship groups, unresolved source-site groups, linked flows/tests, quality signals, and hotspots.
- Added Web File Map and File Detail views backed by the same file projection contract used by backend command and API surfaces.
- Added target-aware command paths and file-layer output for existing graph workflows while preserving symbol, API, route, tool, and quality details.

#### Changed

- Replaced the repo-level analyze `unsupported` aggregate with causal file classification buckets for parsed code, documents, metadata-only inputs, dedicated analyzer inputs, scripts without extractors, static assets, unsupported languages, unknowns, and failures.
- Made `anvien analyze` build and report file projection inventory while keeping symbol/source-site graph facts as the canonical source of truth.
- Bumped the CLI package version from `1.2.3` to `1.2.4`.
- Retired the self-referential generated AI-context skill; generated root context and Anvien skill content now remain sourced from `internal/aicontext`.

#### Fixed

- Fixed healthy docs, manifests, reports, fixtures, COBOL/JCL analyzer inputs, scripts, and static assets being presented as one misleading analyze `unsupported` count.

### 2026-05-26

#### Added

- Added runtime lock metadata and diagnostics so users can distinguish one-shot `analyze` / `embed` work from long-running `serve` and editor-owned `mcp` processes.
- Added Web graph node-spacing diagnostics and dense-graph e2e screenshot coverage for desktop and smaller viewport validation.

#### Changed

- Hardened repository storage locks with PID, host, command, timestamp, and token ownership metadata.
- Bumped the CLI package version from `1.2.2` to `1.2.3`.
- Made stale same-host analyze/embed locks recoverable while preserving live-lock protection and foreign-host safety.
- Made Web graph dense-island placement enforce a deterministic minimum node gap based on rendered Sigma node size.
- Expanded dense island and macro-ring footprints from actual node offsets so large graphs prefer readable separation over compressed overlap.

#### Fixed

- Fixed stale `.anvien/analyze.lock` files blocking analyze after owner death, crash, or reboot.
- Fixed lock release so an older owner cannot remove a lock acquired by a newer process.
- Fixed dense Web graph node overlap by enforcing one rendered node diameter of empty edge-to-edge gap between rendered nodes.

### 2026-05-23

#### Added

- Added user-facing documentation for semantic graph output: App Layer, Functional Area, source-site proof metadata, persisted ResolutionGap entities, Resolution Health, and diagnostic square nodes in the Web UI.
- Documented the graph quality commands `query-health`, `resolution-inventory`, and `source-site-accuracy`.
- Added expanded generated Anvien skill guidance for graph quality, API surface work, cross-repo workflows, runtime packaging, and generated AI context maintenance.
- Added graph orientation labels so Web graph macro rings and node-type islands identify their App Layer and node/filter groups directly on the canvas.

#### Changed

- Bumped the CLI package version from `1.2.1` to `1.2.2`.
- Updated CLI/MCP documentation to describe semantic fields on `query`, `context`, `impact`, `detect-changes`, `route_map`, `shape_check`, and `api_impact`.
- Reconciled generated AI context skill sources, package/setup skill distribution, and root `AGENTS.md` / `CLAUDE.md` guidance around the current Anvien command surface.
- Standardized guidance around the canonical production executable path `anvien\bin\anvien.exe`.

### 2026-05-22

#### Added

- Added App Layer and Functional Area semantic metadata to graph output and consumer surfaces.
- Added persisted ResolutionGap / unresolved-symbol records with source-site inventory and actionability metadata.
- Added graph quality inventory commands and reporting for query health, resolution gaps, source-site accuracy, and resolved-edge accuracy.
- Added Web UI App Layer rings and node-type islands so graph layout separates Backend, API, Frontend, Docs, Config, Test, Generated, Mixed, and related layers when present.

#### Changed

- Made resolved `CALLS` and `ACCESSES` graph edges proof-based, while preserving unresolved, ambiguous, external, dynamic, and unsupported source sites as inspectable graph evidence.
- Separated Resolution Health from topology Graph Health so connected nodes can still expose degraded resolution confidence without being mislabeled as topology unknown.

### 2026-05-21

#### Added

- Added the exe-served in-app launcher start screen as the single packaged entry surface.
- Added explicit folder-picker cancellation and request cancellation propagation for local repository selection.

#### Changed

- Moved the Windows launcher start surface into the Web UI served by the rebuilt `AnvienLauncher.exe`, making the exe the single packaged user entrypoint.
- Removed the separate root HTML start-file flow from active launcher docs and validation expectations.
- Kept the graph-shell Back action as an in-app navigation path back to the served start screen.

#### Fixed

- Fixed the packaged Start action entering the manual `anvien serve` bridge guide instead of the repository chooser/analyze surface.
- Fixed local folder picker UX so a pending native picker no longer traps the Analyze Repository screen or blocks manual path analysis.
- Fixed reset/runtime helper process handling so reset does not rely on visible terminal windows in the packaged flow.
- Fixed Graph Health `Unknown` semantics by separating topology status from unresolved-reference diagnostics.

### 2026-05-20

#### Added

- Added default generation of non-empty root `AGENTS.md` and `CLAUDE.md` Anvien managed blocks from the Go implementation.
- Added base Anvien skill installation during analyze and preserved generated community skills behind the explicit `--skills` path.
- Added the Graph Health connectivity lens taxonomy for connected, isolated, no-incoming, no-outgoing, detached-component, and unknown-connectivity states.
- Added a dedicated Documentation display filter and centered documentation island in the Web graph layout.

#### Changed

- Made generated agent context refresh on analyze without preserving the old root-context bypass behavior.
- Reworked Web graph initial placement into deterministic filter/color islands and kept layout optimization as an explicit manual action.
- Classified graph-health diagnostics, expected-isolated reasons, and topology as separate overlays instead of one ambiguous orphan-node label.

#### Fixed

- Fixed empty or stale generated agent context files by replacing managed blocks cleanly and preserving manual content outside the managed region.
- Fixed automatic graph layout optimization after graph load so normal loading uses deterministic initial layout instead of post-load movement.

### 2026-05-19

#### Added

- Added multi-language graph coverage classification across supported scanner/provider languages and graph fact families.
- Added source-site and graph-truth auditing for property ownership, `HAS_PROPERTY`, and member `ACCESSES` facts.
- Added Web dashboard completeness coverage for graph-present node labels, relationship types, color legend rows, and representative uncommon graph payloads.

#### Changed

- Made the Web graph dashboard enumerate loaded graph node labels and relationship types instead of relying on a small fixed subset.
- Grouped or explained compatibility heritage relationships so `EXTENDS` and `INHERITS` duplicates do not look like independent source facts.
- Preserved parallel relationships between the same source and target instead of collapsing them into a single canvas edge.
- Bounded rendered node sizes so structural nodes remain readable without visually dominating large graphs.

#### Fixed

- Fixed TypeScript and multi-language heritage graph coverage gaps found during the Restaurant_manager audit.
- Fixed misleading standalone property ownership/access graph facts by distinguishing true orphans, false orphans, unknowns, and intentionally unmodeled cases.

### 2026-05-16

#### Fixed

- Closed the Go graph-accuracy gate for the local Anvien graph, including TypeAlias, Variable, and direct `CALLS` recall targets.
- Fixed Go graph extraction/resolution gaps that prevented measured local accuracy categories from reaching the plan target.

### 2026-05-14

#### Changed

- Completed the Go full-conversion cutover for non-Web Anvien runtime authority.
- Kept TypeScript/React as the Web UI display/build surface while moving CLI, backend, analyzer, graph, MCP, setup, packaging, and runtime authority to Go.
- Removed or replaced non-Web TypeScript/Node implementation authority after converted Go implementations and tests passed.

#### Performance

- Preserved the Go runtime as the accepted faster large-repo analyze path while keeping deeper optimization work outside the conversion cutover.

#### Documented

- Closed the TypeScript/Node-to-Go conversion planning thread through the final Go cutover evidence and benchmark ledgers.

### 2026-05-08

#### Changed

- Made Web UI analyze entry points behave as full analyze actions: landing repo cards, local-path analyze, header re-analyze, and header analyze-new now run `/api/analyze` before graph loading.
- Made Web/API analyze force a rebuild at the backend worker boundary so clients cannot accidentally trigger the up-to-date shortcut for user-facing analyze actions.
- Changed post-analyze graph loading to route by selected physical `repoPath` and then by canonical `repoInfo.repoPath` from `/api/repo`, rather than by display name or basename.
- Kept header repository dropdown switching as a load/open-existing-graph action, separate from analyze, while still routing selected repos by path.

#### Fixed

- Fixed Web UI analyze/load drift where a completed analyze could reload a graph by `repoName` and show a different same-name repository.
- Fixed backend repo resolution so absolute paths are matched before name/basename lookup and remain path-first after registry refresh while analyze jobs complete.
- Fixed header re-analyze failure handling so start/SSE failures surface the selected repo path instead of failing silently on the prior graph.

#### Documented

- Documented the Web UI full-analyze path contract and repoPath-first graph loading behavior in `ARCHITECTURE.md`.

### 2026-05-07

#### Added

- Added the Accurate Single-Pass Graph path: parse workers now emit AST-reused `ParsedFile` scope facts, finalize scope/import indexes once, and feed `resolutionPhase` in the default analyze run.
- Added scope-aware graph emission for `CALLS`, `ACCESSES`, `USES`, `INHERITS`, finalized `IMPORTS`, and import-use edges with audit metadata.
- Added Python AST-reused scope capture coverage for imports, classes, functions/methods, `self` bindings, typed `self.field` properties, member calls/accesses, type references, and inheritance.
- Added per-language benchmark coverage reporting and explicit benchmark target git-state metadata, including `repoGitUnavailable` when a target repo is not a git checkout.

#### Changed

- Narrowed legacy `crossFilePhase` so it skips source reread/reprocess work when parse metrics prove complete AST-reused scope coverage.
- Moved useful cross-file receiver/type propagation into scope facts and `resolutionPhase` for TypeScript/JavaScript patterns such as call returns, awaited calls, aliases, field/method-derived receivers, destructuring, for-of elements, JSDoc params, chained fields, and callable properties.
- Made method dispatch strategy-aware through finalized scope inheritance facts and provider MRO strategy inputs.
- Updated `ARCHITECTURE.md` to document the new 13-phase pipeline, `resolutionPhase`, scope fact boundary, crossFile narrowing, resolution workers, and benchmark artifact protocol.

#### Fixed

- Made scope graph emission fail closed when graph node endpoints cannot be mapped.
- Merged scope audit metadata into existing semantic duplicate edges instead of emitting overlapping relationships.
- Fixed ambiguous same-file method/function graph-node aliasing so file-level functions are not hidden by same-name methods.
- Fixed Windows full-suite validation by using Vitest v4 `--no-isolate` for selected single-process forked tests and keeping native DB/API impact tests out of unstable hidden worker paths.

#### Performance

- Added auto/force/off reference-resolution worker mode with deterministic chunk parity checks and worker usage/count metrics.
- Removed avoidable worker-index byte-measurement overhead and added benchmark comparison coverage for resolution timing, graph parity, semantic duplicates, and per-language scope coverage.

### 2026-04-30

#### Changed

- Updated `ARCHITECTURE.md` to match the current codebase: CLI/MCP/HTTP backend, Web UI, shared contracts, Windows local launcher, repo-scoped graph reads, and the Codex/Claude Code session bridge contract.
- Polished the Web UI header, repo dropdown placement, graph loading copy, runtime controls, chat panel spacing, and status/footer presentation.
- Simplified the AI runtime settings panel around the actual local session state instead of legacy provider/model configuration.
- Added global button press states and refined button borders while keeping the repo tree and repo dropdown visually lighter.

#### Fixed

- Unified external graph node selection so choosing a file from the dashboard or selecting a result from search follows the same graph selection/display path as clicking a node in the graph canvas.

### 2026-04-29

#### Added

- Added the repo-scoped HTTP read engine for repo switching and graph loading.
- Added shared repo resolution/runtime ID helpers, a repo-scoped read executor, a graph read/stream service, and a graph streaming HTTP adapter.
- Added focused unit coverage for repo resolution, repo read execution, graph streaming, MCP alignment, and repeated dropdown repo-switching.
- Added an OS-backed local folder picker for analyzing local repositories from the Web UI.
- Added repo removal actions on the repository landing screen.
- Added a placeholder `user_guide.md` entry and removed the placeholder logo from active UI surfaces.

#### Changed

- Migrated `/api/graph`, `/api/query`, `/api/search`, and `/api/grep` away from the old process-global LadybugDB retargeting path for repo-switch graph-load scope.
- Kept Anvien local-only: no cloud control plane, no managed workspace import model, and no required always-on daemon.
- Refreshed the repository list after analyzing a repository from the dropdown flow.
- Reworked loading UI copy and removed the fake hard-coded graph download percentage.

#### Fixed

- Fixed the Web UI dropdown repo-switch failure path by replacing the broken global DB switch/read path with explicit repo-scoped graph reads.
- Fixed dropdown analyze flow so newly analyzed repositories appear in the repository list.

#### Release Notes

- Marked the repo-switch graph engine plan complete for the graph-load scope.
- Bumped the CLI package version to mark the repo-switch engine completion point.

### 2026-04-25

#### Added

- Added the initial packaged launcher flow for Windows.
- Added `AnvienLauncher.exe`, the Go launcher source, packaged Web UI output, backend server wrapper, and protocol registration for the original launcher start action.
- Added launcher reset/start/stop behavior around the same local backend runtime.

#### Changed

- Moved launcher sources and build output from hidden `.anvien-launcher/` to visible `anvien-launcher/`.
- Updated launcher path references, protocol registration, and build packaging after the directory move.

#### Documented

- Added local launcher entry and root HTML launcher plans.
- Documented reset/analyze lifecycle risks around stale runtime processes, repo locks, WAL corruption, and clean rebuild boundaries.
- Paused the full analyze optimization plan after the delivered speedup and measured LadybugDB/FTS bottleneck.

### 2026-04-24

#### Performance

- Added full analyze phase timing and deep persistence instrumentation.
- Made the parse worker path canonical for full analyze and removed the old silent whole-repo sequential fallback behavior from the optimized production path.
- Implemented a deterministic dynamic parse worker scheduler.
- Reduced large-repo parse time on the benchmark path from the prior sequential-fallback profile to the worker-canonical profile while preserving output counts.
- Cached cross-file parser query work after measuring crossFile internals.

#### Fixed

- Fixed Go receiver source attribution during crossFile reprocessing so receiver method calls are not incorrectly attributed to same-name interface methods.
- Made community detection deterministic by fixing randomization/input ordering behavior.

#### Documented

- Recorded full analyze Phase 0 through Phase 4 findings.
- Closed Phase 2 without resolver/linker optimization after measurement showed parse main-thread resolve was not a material bottleneck.
- Paused deeper LadybugDB/FTS optimization until there is a safer strategy that preserves stored graph/search behavior exactly.
- Added the Web graph load/render performance plan as a separate concern from full analyze performance.

### 2026-04-23

#### Added

- Added repo-local execution flow settings.
- Added a master graph links visibility toggle for the Web UI.

#### Changed

- Refined the Web UI shell and controls toward the `The Press` editorial design direction while keeping the graph workspace dark.
- Switched the graph viewport to black and refined graph rendering/selection behavior.
- Reordered graph guidance tools and aligned package metadata/docs with the current Anvien product surface.

#### Fixed

- Unblocked graph stream startup.
- Hardened the `impact` tool contract direction with schema/runtime alignment planning and implementation.

#### Documented

- Added plans for impact contract hardening, smart re-analyze delta graph constraints, graph links behavior, selected graph context, layout/scroll/resize hardening, click-node performance hardening, straight edges/thinner links, and the Web UI `The Press` migration.

### 2026-04-22

#### Performance

- Improved MCP startup performance for stdio clients.

#### Added

- Added configurable process flow caps.

#### Changed

- Updated AI context generation and README content for the local Anvien runtime direction.

#### Documented

- Added MCP startup optimization planning.
- Added Web UI Codex chat runtime optimization planning.

### 2026-04-21

#### Added

- Added a direct `detect-changes` CLI command.

#### Changed

- Completed the Anvien rename rollout across user-facing surfaces, CLI/package metadata, MCP integration, docs, setup, and active Web UI paths.
- Removed SWE-bench, sponsor, eval-server, and remote/eval product surfaces from the active local tool.
- Separated chat runtime state from graph state in the Web UI.
- Retired the legacy Web agent build path.
- Completed local runtime cleanup phases: provider/API-key paths were removed or wrapped out of the active build, and the Web runtime moved to the local session model.

#### Fixed

- Restored local Codex chat and portable onboarding.
- Hardened local runtime chat and WAL recovery.
- Removed stale clone states from local analyze flow.
- Fixed grounding links mock typing.

#### Documented

- Added and refined the local-only runtime migration plan.
- Added Anvien rename rollout planning.
- Added chat/graph separation reviews and supervisor review notes.
- Restored license notice and refreshed agent context/guidance.

### 2026-04-20

#### Added

- Added Phase 1 of the local session runtime bridge.
- Added HTTP session bridge support for local chat sessions, including status, streaming chat, cancellation, and indexed-repo binding.

#### Changed

- Migrated Web/local runtime behavior to the session bridge.
- Locked onboarding and CLI runtime behavior to the local-only flow.
- Enforced local path analyze semantics and began removing clone/remote URL assumptions from the active runtime path.

#### Fixed

- Hardened Phase 1 runtime repo resolution.
- Closed Phase 1 session runtime gaps.
- Enforced the WSL2 default recommendation for full Codex runtime mode.

#### Documented

- Refined the local runtime migration plan, preserved feature parity expectations, and gated wiki behavior for local runtime.
- Snapshotted the repository state before the local runtime migration.

### 2026-03-26

#### Planned

- Documented the complete COBOL language coverage plan, including high-value data-flow edges, EXEC SQL/DLI handling, DECLARATIVES support, CALL USING/RETURNING extraction, STRING/UNSTRING/INSPECT/SET/INITIALIZE access edges, and metadata completeness fixes.
