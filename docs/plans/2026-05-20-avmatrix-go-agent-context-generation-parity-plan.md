# AVmatrix GO Agent Context Generation Parity Plan

Date: 2026-05-20

Status: active

Companion files:

- Benchmark ledger: [2026-05-20-avmatrix-go-agent-context-generation-parity-benchmark.md](2026-05-20-avmatrix-go-agent-context-generation-parity-benchmark.md)
- Evidence ledger: [2026-05-20-avmatrix-go-agent-context-generation-parity-evidence.md](2026-05-20-avmatrix-go-agent-context-generation-parity-evidence.md)

## Rules

1. Use AVmatrix for codebase analysis and impact checks before implementation slices that edit functions, classes, methods, or shared generation behavior.
2. For doc-only commits, do not use AVmatrix.
3. Compare AVmatrix-GO behavior against the original TypeScript implementation in `E:\avmatrix-main` before declaring parity.
4. Do not claim generated `AGENTS.md`, `CLAUDE.md`, MCP resource, MCP tool, or skill guidance is correct unless it is grounded in the Go implementation, original implementation, or tests.
5. Treat `AGENTS.md`, `CLAUDE.md`, `.claude/`, and `.avmatrix/` as generated or ignored local artifacts unless a phase explicitly changes tracked packaging inputs.
6. Update this plan, the evidence ledger, and the benchmark ledger as each implementation slice is completed.
7. Run focused tests after each implementation slice and final smoke commands before closure.
8. Commit each completed implementation slice after evidence and benchmark rows are updated.

## Problem

AVmatrix-GO did not generate useful root agent context files by default. In the reported repository state, `AGENTS.md` and `CLAUDE.md` existed but were empty, while the original TypeScript project under `E:\avmatrix-main` generated full AVmatrix context during analyze.

The conversion gap is serious because agents rely on these root files to learn:

- when to use AVmatrix tools;
- when impact analysis and `detect_changes` are required;
- which MCP tools and resources exist;
- which AVmatrix skills should be consulted;
- which CLI fallback commands exist when MCP is not available.

The generated content also needs an accuracy pass. A generated instruction block that is non-empty but stale, mislabeled, or incomplete can steer agents incorrectly.

## Scope Boundary

Implementation may touch:

- `internal/cli/command.go`;
- `internal/cli/analyze_postrun.go`;
- `internal/aicontext/aicontext.go`;
- AI context generation tests under `internal/aicontext`;
- analyze command tests under `internal/cli`;
- MCP surface tests under `internal/mcp` if generated resource/tool references are updated;
- packaged or embedded base skill sources if AVmatrix-GO needs rich skill content instead of placeholder skill files;
- documentation for generated agent context behavior.

Out of scope unless a later phase explicitly reopens it:

- changing MCP tool semantics;
- changing registry naming or graph storage semantics;
- making generated `AGENTS.md` or `CLAUDE.md` tracked source files;
- rewriting generated community skill selection logic beyond the `--skills` gate;
- adding product behavior unrelated to agent context generation;
- broad redesign of all agent policies outside the AVmatrix managed block.

## Parity Contract

The accepted Go behavior must match the original intent, not only the current port:

| Behavior | Required contract |
|---|---|
| Default `avmatrix analyze` | Generates or updates `AGENTS.md`, `CLAUDE.md`, and base AVmatrix skill files after a successful analyze. |
| `--skills` | Adds generated community skill files and generated-skill rows only; it must not be the gate for root `AGENTS.md`, `CLAUDE.md`, or base skill installation. |
| `--skip-agents-md` | Preserves manual `AGENTS.md` and `CLAUDE.md` content exactly while still allowing non-root generated artifacts that are part of the accepted contract. |
| `--no-stats` | Omits volatile symbol, relationship, and process counts from generated root context. |
| Empty root files | Replaced with a clean AVmatrix managed block, not left empty and not converted into malformed whitespace. |
| Existing AVmatrix block | Replaced in place between `<!-- avmatrix:start -->` and `<!-- avmatrix:end -->`. |
| Legacy managed block | Replaced when a prior managed Code Intelligence block exists under another marker name. |
| Manual root files without managed block | Preserve manual content and append one AVmatrix managed block. |
| Base skills | Install rich AVmatrix skill content, not thin fallback placeholders, unless the rich package source is genuinely unavailable and evidence records the fallback. |
| Generated instructions | Refer only to MCP tools, resources, skills, and CLI commands that exist in AVmatrix-GO. |

## Generated Content Accuracy Requirements

The root block must be audited against the Go codebase before closure.

Required content corrections or confirmations:

- The section currently labeled `## CLI` is a skill-file table. It must either be renamed to `## Skills` or split into separate `## Skills` and `## CLI` sections.
- Resource references must match `internal/mcp/resources.go`, including canonical resources and templates that are relevant to agents.
- Tool references must match `internal/mcp/tools.go` and MCP surface snapshot tests.
- Staleness wording must be internally consistent. If the policy says to refresh before graph-based work, the block should not also imply refresh is only needed after a stale warning unless the difference is intentional and documented.
- CLI fallback commands must be present if the block tells agents to use AVmatrix when MCP tools are unavailable.
- Skill file paths must point to files actually installed by the Go generator.
- Generated community skill rows must appear only when `--skills` produces generated skills.

## Acceptance Criteria

- A default analyze run creates non-empty `AGENTS.md` and `CLAUDE.md` with one valid AVmatrix managed block.
- A default analyze run installs base skill files under `.claude/skills/avmatrix/`.
- Running analyze without `--skills` does not create `.claude/skills/generated/`.
- Running analyze with `--skills` creates generated community skills when the graph has eligible communities and appends generated rows to the root context.
- `--skip-agents-md` preserves existing root agent files byte-for-byte.
- Empty root agent files are replaced cleanly.
- Legacy managed Code Intelligence blocks are replaced instead of duplicated.
- Generated root content uses accurate MCP tool names, accurate MCP resource URIs, accurate skill headings, and a separate CLI fallback section if CLI commands are listed.
- Base skill files are rich enough to be useful and are sourced from packaged content or embedded Go assets, not the current one-paragraph fallback content.
- Focused tests pass for `internal/aicontext` and `internal/cli`.
- MCP surface snapshot or equivalent validation passes if generated content references are updated.
- Final smoke on `E:\AVmatrix-GO` confirms generated root files are non-empty, base skills are installed, generated artifacts remain ignored by git, and analyze still completes successfully.

## Baseline Requirements

The following baseline must be recorded before implementation closure:

- pre-fix or reproduced default analyze behavior for root `AGENTS.md` and `CLAUDE.md`;
- default analyze behavior after the fix;
- `--skills` analyze behavior after the fix;
- `--skip-agents-md` behavior after the fix;
- empty file upsert behavior;
- legacy managed block replacement behavior;
- current and final base skill file sizes;
- current and final generated root file sizes;
- exact generated root block headings and resource/tool references;
- test commands and results;
- smoke command output and git-ignore status for generated artifacts.

## Codebase Findings Before Implementation

Initial source inspection identified these facts:

- The original TypeScript implementation imports and calls `generateAIContextFiles` from `E:\avmatrix-main\avmatrix\src\cli\ai-context.ts` during analyze in `E:\avmatrix-main\avmatrix\src\core\run-analyze.ts`.
- The original TypeScript generator replaces any managed Code Intelligence block using a marker-aware `MANAGED_SECTION_PATTERN`.
- The original TypeScript generator reads rich bundled skill Markdown from `E:\avmatrix-main\avmatrix\skills\avmatrix-*.md` before falling back to minimal generated skill content.
- AVmatrix-GO has AI context generation in `internal/aicontext/aicontext.go`.
- AVmatrix-GO's base skill installation currently emits thin generated fallback content instead of copying or embedding the rich source skill Markdown.
- AVmatrix-GO's MCP resources include more surfaces than the current generated root resource table shows, including `avmatrix://repos`, `avmatrix://setup`, `avmatrix://repo/{name}/schema`, and `avmatrix://repo/{name}/cluster/{clusterName}`.
- AVmatrix-GO's MCP tools include the core tool names used in the generated block: `query`, `cypher`, `context`, `detect_changes`, `rename`, and `impact`.
- The generated root block needs a naming and wording audit because its `## CLI` section currently points to skill files rather than CLI commands.

## Phase 0 - Plan Creation

- [x] [P0-A] Create the plan, evidence ledger, and benchmark ledger for this issue.
- [x] [P0-B] Ground the plan in observed conversion parity gaps and generated output risks.
- [x] [P0-C] Commit the plan file set when requested. Commit: `c7b427f docs: plan agent context generation parity fix`.

## Phase 1 - Parity Discovery

- [x] [P1-A] Record the original TypeScript analyze-to-agent-context flow with source paths and command evidence.
- [x] [P1-B] Record the current Go analyze-to-agent-context flow with source paths and command evidence.
- [x] [P1-C] Identify every flag that affects AI context generation: `--skills`, `--skip-agents-md`, and `--no-stats`.
- [x] [P1-D] Compare original packaged skill sources against Go generated skill output.
- [x] [P1-E] Compare generated root resource/tool references against `internal/mcp` definitions and tests.
- [x] [P1-F] Record all findings in the evidence ledger before editing implementation.

## Phase 2 - Analyze Invocation Contract

- [x] [P2-A] Ensure successful default analyze calls AI context generation unconditionally after graph registration.
- [x] [P2-B] Ensure `--skills` only controls generated community skills and generated rows.
- [x] [P2-C] Ensure `--skip-agents-md` skips only the accepted root-file write behavior and preserves manual content exactly.
- [x] [P2-D] Ensure `--no-stats` continues to omit volatile stats in root context.
- [x] [P2-E] Add focused CLI tests for default analyze, `--skills`, `--skip-agents-md`, and `--no-stats`.

## Phase 3 - Managed Block Upsert Correctness

- [x] [P3-A] Replace empty `AGENTS.md` and `CLAUDE.md` cleanly.
- [x] [P3-B] Replace existing AVmatrix managed blocks in place.
- [x] [P3-C] Replace legacy managed Code Intelligence blocks in place.
- [x] [P3-D] Append to manual root files that do not contain a managed block.
- [x] [P3-E] Add tests for all four upsert modes.

## Phase 4 - Generated Root Content Accuracy

- [x] [P4-A] Rename or split the mislabeled `## CLI` section so skill-file guidance and CLI commands are not conflated.
- [x] [P4-B] Add or confirm resource references for all agent-relevant Go MCP resources.
- [x] [P4-C] Confirm tool references against MCP surface tests.
- [x] [P4-D] Rewrite staleness and refresh wording so it is consistent and enforceable.
- [x] [P4-E] Add a CLI fallback section with real `avmatrix` commands if generated content references terminal fallback behavior.
- [x] [P4-F] Add tests that assert the generated block contains the accepted headings, resources, tools, and skill paths.

## Phase 5 - Base Skill Content Parity

- [x] [P5-A] Decide where rich base skill Markdown should live in AVmatrix-GO: tracked `skills/`, embedded assets, or another package-local source.
- [x] [P5-B] Port the six accepted base skills from the original package or document intentional differences.
- [x] [P5-C] Install rich skill content during AI context generation.
- [x] [P5-D] Keep fallback placeholder generation only as an explicit fallback path with evidence.
- [x] [P5-E] Add tests that fail when installed base skills collapse to one-paragraph placeholders.

## Phase 6 - Validation

- [x] [P6-A] Run focused Go tests for `internal/aicontext` and `internal/cli`.
- [x] [P6-B] Run MCP surface tests if generated resource or tool references change.
- [x] [P6-C] Run full applicable Go tests before closure. `go test ./internal/... ./cmd/...` passed; `go test ./...` still fails on intentionally non-buildable fixture packages.
- [x] [P6-D] Run default analyze smoke on `E:\AVmatrix-GO`.
- [x] [P6-E] Run analyze smoke with `--skills` when generated community skill behavior is touched.
- [x] [P6-F] Run analyze smoke with `--skip-agents-md` against a temp repo or controlled fixture.
- [x] [P6-G] Verify generated artifacts are ignored by git and no unintended tracked files are staged.
- [x] [P6-H] Update evidence and benchmark ledgers with all validation results.

## Phase 7 - Closure

- [x] [P7-A] Update this plan checklist after each completed implementation slice.
- [x] [P7-B] Update benchmark ledger with root file sizes, skill sizes, and smoke graph counts.
- [x] [P7-C] Update evidence ledger with commands, files changed, tests, and conclusions.
- [ ] [P7-D] Commit implementation and plan updates together only after validation is recorded.
- [ ] [P7-E] Confirm no unrelated dirty files are included in the commit.

## Definition Of Done

The plan is complete only when default AVmatrix-GO analyze reliably generates accurate, non-empty, useful agent context files and skill files; all relevant flags behave intentionally; generated content matches actual Go MCP and CLI capabilities; focused tests pass; final smoke evidence is recorded; benchmark rows are updated; and the final commit excludes unrelated workspace changes.
