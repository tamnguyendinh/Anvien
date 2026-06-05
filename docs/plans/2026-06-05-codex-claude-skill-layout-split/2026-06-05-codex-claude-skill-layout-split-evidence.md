# Evidence Ledger

Title: Codex And Claude Skill Layout Split

Date: 2026-06-05

Status: Complete

Companion files:

- Plan: [2026-06-05-codex-claude-skill-layout-split-plan.md](2026-06-05-codex-claude-skill-layout-split-plan.md)
- Benchmark ledger: [2026-06-05-codex-claude-skill-layout-split-benchmark.md](2026-06-05-codex-claude-skill-layout-split-benchmark.md)

## Evidence Rules

1. This file records planning and future implementation evidence.
2. Tests are validation evidence only, not source of truth for desired behavior.
3. Generated files are validation output, not permanent source of truth.
4. Product-surface facts must be recorded separately from codebase facts.
5. Implementation evidence must be appended as phases complete.

## E0 - Planning Context

User-approved discussion outcome:

```text
AGENTS.md  -> .agents/skills/<package>/...
CLAUDE.md  -> .claude/skills/<package>/...
```

Interpretation:

- `AGENTS.md` is Codex-facing.
- `CLAUDE.md` is Claude Code-facing.
- The generated skill path in each file must match that agent's native skill layout.
- The old `anvien/` namespace layer should be removed from generated guide paths and generated skill output.

## E1 - Product Surface Evidence

Codex manual evidence gathered in this session:

- Codex uses `AGENTS.md` for project guidance.
- Codex repo skills are discovered under `.agents/skills`.
- Codex starts skill discovery from skill metadata and loads full `SKILL.md` only after selecting a skill.

Claude Code surface evidence for this plan is treated as user-approved requirement for this slice:

- `CLAUDE.md` is the Claude Code-facing guidance file.
- `.claude/skills` is the Claude skill layout to generate for Claude Code.

## E2 - Codebase Facts

Commands run:

```text
go run ./cmd/anvien analyze --force
go run ./cmd/anvien query files "skill package install generated skills claude agents" --repo Anvien
rg -n "\.claude/skills|claude.*skills|InstallSkillPackages|skillGuideUse|skillPackage|anvien/" internal/aicontext internal/cli
```

Observed facts:

- `internal/aicontext/aicontext.go` writes one shared rendered block to both `AGENTS.md` and `CLAUDE.md`.
- `internal/aicontext/aicontext.go` currently installs analyze-time generated skills to `.claude/skills/anvien`.
- `internal/aicontext/skill_packages.go` currently renders guide paths with `.claude/skills/anvien/`.
- `internal/cli/setup_command.go` already installs Codex setup skills to `$HOME/.agents/skills`.
- `internal/cli/setup_command.go` already installs Claude Code setup skills to `$HOME/.claude/skills`.
- `internal/aicontext/skill_packages.go` currently exact-mirrors the whole `targetDir`, which is safe for a generated namespace but unsafe for direct `.agents/skills` or `.claude/skills` roots unless scoped.
- `internal/gitignore/managed.go` currently ignores `.claude/` and `.agents/` entirely.

## E3 - Current Risk Evidence

Primary risk:

- Direct package layout removes the generated `anvien/` namespace, so Anvien package names can collide with user custom skills.

Sync risk:

- Existing `actualSkillSnapshot(targetDir)` walks every file under `targetDir`.
- Existing `diffSkillSnapshots` deletes actual paths missing from desired paths.
- Therefore using `.agents/skills` or `.claude/skills` as `targetDir` without scoping would delete unrelated custom skills.

Required mitigation:

- Scope actual snapshot and deletion to Anvien-managed package roots only.
- Add collision detection before overwriting a direct package root.

## E4 - Planning Commands

Planning commands completed:

```text
go run ./cmd/anvien analyze --force
go run ./cmd/anvien query files "skill package install generated skills claude agents" --repo Anvien
```

Result:

- Graph refreshed successfully.
- Query identified `internal/aicontext/skill_packages.go` and `internal/aicontext/aicontext_test.go` as primary files.

## E5 - Implementation Evidence

Implementation started after user command `triển khai plan`.

Worktree hygiene before implementation:

- `git status --short` initially showed only the plan directory as untracked.
- During implementation, `README.md` was found modified but was not changed or staged by this slice.
- `ai_intelligence_graph_knowledge_map_skill_kit.svg` was untracked and not touched.
- `go build ./cmd/anvien` created root `anvien.exe`; it was removed as an implementation artifact.

Graph refresh and impact:

```text
go run ./cmd/anvien analyze --force
```

Pre-edit graph refresh result:

- scanned files: 1345
- parsed code files: 697
- graph nodes: 83005
- graph relationships: 121656

Impact checks recorded before code edits:

| Symbol | Risk | Affected source containment |
|---|---|---|
| `GenerateAIContextFiles` | CRITICAL | `internal/aicontext/aicontext.go`, CLI analyze postrun/command paths |
| `renderAnvienBlock` | CRITICAL | `internal/aicontext/aicontext.go`, CLI analyze postrun/command paths |
| `installBaseSkills` | CRITICAL | `internal/aicontext/aicontext.go`, CLI analyze postrun/command paths |
| `installSkillPackagesTo` | HIGH | `internal/aicontext/skill_packages.go`, `internal/aicontext/aicontext.go`, setup/analyze callers |
| `setupInstallSkillsTo` | LOW | `internal/cli/setup_command.go` |
| `actualSkillSnapshot` | CRITICAL | skill package sync, analyze/setup callers |
| `diffSkillSnapshots` | CRITICAL | skill package sync, analyze/setup callers |
| `applySkillSyncPlan` | CRITICAL | skill package sync, analyze/setup callers |
| `verifySkillSnapshot` | CRITICAL | skill package sync, analyze/setup callers |

Containment:

- HIGH/CRITICAL was treated as blast-radius scope warning, not an edit ban.
- Edits stayed inside AI context generation, skill installer ownership sync, CLI flag text/tests, managed `.gitignore` tests, MCP setup resource text, and this plan set.

Implementation changes:

- `AGENTS.md` rendering now uses `.agents/skills/<package>/...` paths.
- `CLAUDE.md` rendering now uses `.claude/skills/<package>/...` paths.
- The generated sentence beginning `This project is indexed by Anvien` was removed from generated context.
- `--no-stats` remains accepted for compatibility but is documented as a deprecated no-op because generated context no longer includes volatile stats.
- Analyze installs repo skills to both `.agents/skills` and `.claude/skills`.
- Direct-root manifest names are `.agents-skill-manifest.json` and `.claude-skill-manifest.json`.
- The installer now scopes actual snapshot/diff/delete/verify to desired package roots and manifest-managed package roots.
- Unrelated custom roots are preserved.
- Same-name unowned package roots collide unless their full file set and hashes exactly match the desired package snapshot.
- Exact unowned matches are adopted and recorded in result summary.
- Legacy `.claude/skills/anvien` and `.agents/skills/anvien` roots are removed only when they contain an Anvien-managed legacy manifest.
- `internal/mcp/resources.go` now describes managed output as `AGENTS.md`, `CLAUDE.md`, `.agents/skills/**`, and `.claude/skills/**`.

Managed `.gitignore` decision:

- Broad `.agents/` and `.claude/` ignores were retained for this slice because that is the current generated local-state policy.
- Ownership safety is enforced by the installer, not by ignore rules.
- This means custom repo skills under those roots are preserved on disk by sync but still hidden by `.gitignore` unless users force-add or later policy changes narrow the ignore entries.

Generated output evidence after running current implementation:

```text
go run ./cmd/anvien analyze --force
```

Result:

- graph nodes: 83117
- graph relationships: 121856
- `AGENTS.md`: 34 `.agents/skills/` references, 0 `.claude/skills/` references, 0 legacy `skills/anvien/` references, 0 indexed-project sentences.
- `CLAUDE.md`: 34 `.claude/skills/` references, 0 `.agents/skills/` references, 0 legacy `skills/anvien/` references, 0 indexed-project sentences.
- `.agents/skills`: 34 package roots, 43 `SKILL.md`, 584 payload files, manifest present at `.agents/skills/.agents-skill-manifest.json`.
- `.claude/skills`: 34 package roots, 43 `SKILL.md`, 584 payload files, manifest present at `.claude/skills/.claude-skill-manifest.json`.
- `.claude/skills/anvien`: absent after migration cleanup.
- `.agents/skills/anvien`: absent after migration cleanup.

Build/test evidence:

| Command | Result |
|---|---|
| `go build ./...` | Failed on existing non-buildable fixtures under `anvien/test/fixtures` (`models`, `animal`, mixed package fixture, C fixture). Not caused by this change. |
| `go build ./cmd/anvien` | Passed; root `anvien.exe` artifact was removed. |
| `go build ./cmd/... ./internal/...` | Passed. |
| `go test ./internal/aicontext` | Passed. |
| `go test ./internal/cli` | Passed. |
| `go test ./internal/gitignore` | Passed. |
| `go test ./internal/mcp` | Passed. |
| `go test ./cmd/... ./internal/...` | Failed only at existing `internal/lbugschema` missing baseline file `baseline/phase-1-contract-freeze/ladybugdb-graph-contract.json`; all changed packages passed. |

Detect-changes evidence:

```text
go run ./cmd/anvien detect-changes --repo Anvien --scope all
```

Result:

- command passed;
- summary risk: critical;
- changed files reported: 8;
- affected files reported: 15;
- affected app layer: backend;
- affected functional areas: cli, mixed, unknown;
- high-risk changed source files included `internal/aicontext/aicontext.go`, `internal/aicontext/skill_packages.go`, `internal/cli/command.go`, and `internal/mcp/resources.go`;
- output also reported pre-existing/user README/readme changes, which were not staged for this implementation slice.

Implementation commit:

```text
45f109d Split Codex and Claude skill layouts
```

## E6 - Plan Review On 2026-06-05

Review result:

- No conflict with the approved high-level direction.
- One execution-order issue was found: direct-root install was listed before scoped ownership sync was implemented. The plan now requires scoped ownership sync and collision handling before analyze targets `.agents/skills` or direct `.claude/skills`.
- One wording issue was found: the invariant said `AGENTS.md` must never point to `.claude/skills/**`, which was too broad. The plan now scopes that rule to `Skill Selection Guide` `Use` paths.
- One coverage gap was found: `setupInstallSkillsTo` is shared by Codex, Claude Code, Cursor, and OpenCode setup paths. The plan now requires custom-skill preservation for every shared installer target, not only Codex and Claude.
- One migration gap was found: legacy cleanup mentioned `.claude/skills/anvien/**` more strongly than `.agents/skills/anvien/**`. The plan now covers both legacy namespaces.
- One `.gitignore` ambiguity was found: broad `.agents/` and `.claude/` ignores can hide user-owned repo skills. The plan now requires an explicit evidence-backed decision before commit.
- User approved removing the generated sentence beginning `This project is indexed by Anvien` entirely, with no replacement sentence, because it makes `AGENTS.md` and `CLAUDE.md` drift with low practical value.
- The plan now requires reviewing `Options.NoStats` and CLI `--no-stats`, because removing the volatile sentence may make that option obsolete.
- Second review found that direct-root manifests must be specified explicitly. The final approved direct-root manifest names are `.agents/skills/.agents-skill-manifest.json` and `.claude/skills/.claude-skill-manifest.json`.
- Second review found adoption rules were under-specified. The plan now allows adoption only when an unowned same-name root exactly matches the desired Anvien package snapshot and contains no extra files; otherwise it is a collision.
- Second review found `internal/mcp/resources.go` still teaches the old `.claude/skills/anvien/**` generated namespace. The plan now includes a secondary guidance update phase.
- Second review clarified that foreign same-name roots are collisions unless they satisfy exact-match adoption, and collision behavior must fail clearly rather than skip silently.
- Second review fixed the benchmark table so collision and adoption counts are separate measured columns.
