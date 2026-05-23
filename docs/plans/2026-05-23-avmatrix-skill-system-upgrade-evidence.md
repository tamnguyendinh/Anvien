# AVmatrix Skill System Upgrade Evidence Ledger

Date: 2026-05-23

Status: Planned

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
- Phase 0 and Phase 2 must reconcile package/editor skill installation with embedded AI-context skill source so the packaged setup path does not drift from generated repository-local skills.

MCP/resource guidance finding:

- `internal/mcp/resources.go` contains setup/resource/tool guidance including MCP tool tables and setup reference output.
- This file must be part of the plan's stale-guidance search because updating only `internal/aicontext/aicontext.go` would leave another user-facing command guide that can drift.
