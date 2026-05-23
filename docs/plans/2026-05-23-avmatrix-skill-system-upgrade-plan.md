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

## Scope

Implementation may touch:

- `internal/aicontext/aicontext.go`;
- `internal/aicontext/skills/*.md`;
- tests under `internal/aicontext`;
- CLI/setup/package tests under `internal/cli` if they assert skill counts, skill paths, generated output, or packaging behavior;
- README and user-facing docs that explain generated AVmatrix skills, AI context setup, or AVmatrix command surfaces;
- packaging/setup code only if source inspection proves the expanded embedded skill set is not installed or packaged correctly.

Out of scope unless source inspection proves it is required:

- changing the behavior of AVmatrix graph analysis commands;
- changing Web UI graph rendering;
- changing MCP tool schemas;
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
- Full build, focused tests, generation smoke, setup/package validation if touched, and `detect-changes` pass before closure.

## Phase 0 - Generator Source Trace And Command Inventory

- [ ] [P0-A] Trace the generator ownership for `.claude/skills/avmatrix/**` and record the result in the evidence ledger. The trace must identify the source skill files, the embedded filesystem owner, the `baseSkills` registry, `baseSkillContent`, `installBaseSkills`, `GenerateAIContextFiles`, the generated root Skills table, the analyze post-run caller, and any setup/package paths that copy installed skills into editor-specific directories.

- [ ] [P0-B] Inventory the current source and generated skill set before implementation. Record each skill id, source file path, generated output path, source byte/line count, generated byte/line count, top headings, and whether the generated output matches the embedded source. This inventory must prove whether `.claude/skills/avmatrix/**` is source or generated validation output.

- [ ] [P0-C] Inventory the current AVmatrix command surface from code/tests/help output before writing skill content. Record actual available CLI commands, MCP tools, MCP resources, setup/package commands, runtime commands, group/cross-repo commands, API-surface commands, and graph-quality/accuracy commands. The evidence must distinguish implemented commands from planned or absent commands so skills do not document non-existent behavior as real.

- [ ] [P0-D] Build the skill routing matrix from the command inventory. Map every current command/tool/resource family to one primary skill and any secondary skill references, then record the final decision in evidence. The matrix must include the existing six skills and the proposed new skills for graph quality, API surface, cross-repo work, runtime/packaging, and AI context generation.

## Phase 1 - Embedded Skill Source Upgrade

- [ ] [P1-A] Upgrade the six existing embedded skill Markdown files in `internal/aicontext/skills/`. Each file must become a practical task guide with command choices, when to use each AVmatrix surface, validation expectations, and current limitations. `avmatrix-impact-analysis` must explain that HIGH/CRITICAL is blast-radius evidence to report and account for, not a blanket prohibition against required work.

- [ ] [P1-B] Add the new embedded source skill files under `internal/aicontext/skills/`: `avmatrix-graph-quality.md`, `avmatrix-api-surface.md`, `avmatrix-cross-repo.md`, `avmatrix-runtime-packaging.md`, and `avmatrix-ai-context.md`. Each new skill must contain concrete usage guidance, command examples based on implemented commands, expected outputs, and validation notes.

- [ ] [P1-C] Update the base skill registry and generated Skills table in `internal/aicontext/aicontext.go`. The registry and generated table must include all final skills, use repo-agnostic descriptions, and avoid splitting AVmatrix into misleading MCP-only versus CLI-only capability lists.

- [ ] [P1-D] Add or update `internal/aicontext` tests so generated root files and generated base skills are protected. Tests must assert the final skill ids, generated `.claude/skills/avmatrix/<skill>/SKILL.md` paths, generated Skills table links, and representative key phrases for the new command surfaces.

- [ ] [P1-E] Add coverage tests that prevent the guide from regressing back to a six-skill or analyze/query/impact-only view. The tests should check the generated guidance for the AI context skill, graph-quality skill, API-surface skill, cross-repo skill, runtime/packaging skill, and a current command-surface fragment confirmed in Phase 0.

## Phase 2 - Setup, Package, And Documentation Integration

- [ ] [P2-A] Verify the analyze post-run path installs the expanded base skill set through the same normal generation path that creates `AGENTS.md` and `CLAUDE.md`. If tests currently assert the old six-skill count or specific old table rows, update them to assert the new final set.

- [ ] [P2-B] Verify setup/editor installation behavior for the expanded embedded skill set. Inspect and test `setupInstallEditorSkills` and related setup command behavior so supported editor skill directories receive the same final skill content without relying on generated repository-local `.claude/skills/avmatrix/**` as source.

- [ ] [P2-C] Verify package/runtime distribution behavior for the expanded embedded skill set. If packaging tests or package assembly code enumerate skills, update them so the packaged tool can generate and install the final skill set from embedded source files.

- [ ] [P2-D] Update README and relevant user-facing docs that describe AVmatrix skills, AI context generation, setup, or usage. The docs must tell users that `AGENTS.md`, `CLAUDE.md`, and `.claude/skills/avmatrix/**` are generated by AVmatrix, and that source changes belong in the embedded skill source and generator code.

- [ ] [P2-E] Search the active documentation for stale six-skill-only guidance or stale wording that treats MCP and CLI as separate incomplete command lists. Update current guides and README-style docs; leave historical ledgers untouched unless they are actively reused as user guidance.

## Phase 3 - Regeneration And Validation

- [ ] [P3-A] Run the full build gate before tests: `powershell -ExecutionPolicy Bypass -File avmatrix-launcher\build.ps1`. Record the command result in evidence.

- [ ] [P3-B] Run focused backend/CLI tests for AI context generation, setup, and packaging surfaces touched by the implementation. The minimum expected test scope is `go test .\internal\aicontext .\internal\cli -count=1`, expanded as needed if package/setup code outside those packages changes.

- [ ] [P3-C] Run the normal generation path with `avmatrix analyze --force` and no `--skip-agents-md`. Verify generated `AGENTS.md`, generated `CLAUDE.md`, and `.claude/skills/avmatrix/**` contain the final skill set and expected content fragments.

- [ ] [P3-D] Compare source and generated skill inventories after regeneration. Record final skill count, generated file paths, byte/line counts, and any intentional generated differences in the benchmark ledger.

- [ ] [P3-E] Validate setup/package behavior if Phase 2 changed setup or package code. Record the exact command outputs and installed/packaged skill file inventories in evidence and benchmark ledgers.

- [ ] [P3-F] Run `detect-changes` before commit and record the affected scope. Commit the implementation slice after checklist items and ledgers are updated.

## Phase 4 - Zero-Trust Closure Review

- [ ] [P4-A] Review the codebase and documentation for old assumptions after implementation. Search for old six-skill-only tables, old `avmatrix-cli` descriptions that omit current command families, stale generated-output instructions, and any direct-edit guidance for `.claude/skills/avmatrix/**`; fix active docs that would mislead users or agents.

- [ ] [P4-B] Re-run the final validation commands required by this plan after the closure review changes. Record final pass/fail counts, generated inventory, setup/package inventory if applicable, and any remaining limitation in the evidence and benchmark ledgers.

- [ ] [P4-C] Mark the plan complete only after source files, generated validation output, tests, docs, benchmark ledger, evidence ledger, and commit state all agree on the final skill set.

