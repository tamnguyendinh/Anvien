# Plan

Title: Codex And Claude Skill Layout Split

Date: 2026-06-05

Status: In Progress - implementation validation complete; detect-changes and commit pending

Companion files:

- Evidence ledger: [2026-06-05-codex-claude-skill-layout-split-evidence.md](2026-06-05-codex-claude-skill-layout-split-evidence.md)
- Benchmark ledger: [2026-06-05-codex-claude-skill-layout-split-benchmark.md](2026-06-05-codex-claude-skill-layout-split-benchmark.md)

## Master Rules

1. Follow active repository instructions and generated `AGENTS.md`.
2. Use `anvien-planner` for this plan artifact.
3. Use Anvien before graph-based code work and record impact before implementation edits.
4. HIGH or CRITICAL blast radius is a warning, not a ban.
5. Do not edit generated `AGENTS.md`, `CLAUDE.md`, `.agents/skills/**`, or `.claude/skills/**` as source of truth.
6. Source of truth remains `internal/aicontext/skills/**`.
7. `AGENTS.md` is the Codex-facing context file.
8. `CLAUDE.md` is the Claude Code-facing context file.
9. Code first; tests validate behavior after implementation.
10. Run full build before tests when implementation starts.
11. Record evidence and benchmarkable counts as each implementation slice completes.
12. Run `anvien detect-changes --repo Anvien --scope all` before implementation commits.
13. Commit only approved implementation slices after validation.

## Goal

Split generated skill layout and generated guide paths by agent surface:

```text
internal/aicontext/skills/<package>/**
  -> .agents/skills/<package>/**
  -> .claude/skills/<package>/**
```

Generated guide paths must be surface-specific:

```text
AGENTS.md -> .agents/skills/<package>/...
CLAUDE.md -> .claude/skills/<package>/...
```

The generated path must not include the old `anvien/` namespace layer:

```text
old: .claude/skills/anvien/<package>/...
new: .claude/skills/<package>/...
new: .agents/skills/<package>/...
```

## Problem

Current analyze output is Claude-shaped for both agents. `GenerateAIContextFiles` renders one shared Anvien block and writes the same block into both `AGENTS.md` and `CLAUDE.md`. The `Skill Selection Guide` path renderer hardcodes `.claude/skills/anvien/`, so Codex sees Claude skill paths in `AGENTS.md`.

Current repo skill installation also installs generated skills only into `.claude/skills/anvien/**` during analyze. Codex-native repo skill discovery uses `.agents/skills/**`, so Codex does not get a native repo skill layout from analyze.

The old `anvien/` namespace kept generated output safe because the installer could exact-mirror that namespace without touching user custom skills. Removing the namespace creates a new ownership problem: the target roots become `.agents/skills` and `.claude/skills`, which can also contain user or team custom skills. The sync engine must not exact-mirror the whole target root.

## Approved Direction

Use direct package roots per surface:

```text
AGENTS.md  -> .agents/skills/<package>/...
CLAUDE.md  -> .claude/skills/<package>/...
```

After analyze, a package such as `problem-solving` should produce:

```text
.agents/skills/problem-solving/problem-solving-parent-skill/SKILL.md
.claude/skills/problem-solving/problem-solving-parent-skill/SKILL.md
```

And the generated guide rows should be:

```text
AGENTS.md: .agents/skills/problem-solving/problem-solving-parent-skill/SKILL.md
CLAUDE.md: .claude/skills/problem-solving/problem-solving-parent-skill/SKILL.md
```

## Scope

In scope:

- `internal/aicontext/aicontext.go` AI context rendering and analyze-time skill install targets.
- Removing the volatile indexed-project sentence from generated `AGENTS.md` and `CLAUDE.md`.
- `internal/aicontext/skill_packages.go` install path rendering, manifest, snapshot, diff, verify, delete, and collision behavior.
- `internal/cli/setup_command.go` validation that Codex and Claude Code setup install compatible direct package layouts.
- Other setup skill targets that share `setupInstallSkillsTo`, including Cursor and OpenCode, because installer ownership behavior is shared even when their guide files are not part of this plan.
- `internal/gitignore/**` managed ignore entries if direct repo skill roots need more precise ignore behavior.
- `internal/mcp/resources.go` or other generated guidance text that still claims only `.claude/skills/anvien/**`.
- Tests for generated `AGENTS.md`, generated `CLAUDE.md`, dual repo skill install, setup skill install, legacy cleanup, and custom-skill preservation.
- Evidence and benchmark ledgers for package counts, generated path counts, token counts, build/test validation, and detect-changes.

Out of scope:

- Changing skill bodies or skill descriptions.
- Changing the Anvien command selection guide.
- Changing Codex or Claude product behavior.
- Publishing or packaging plugins.
- Web UI behavior.
- Manual edits to generated context or generated skill output.

## Invariants

1. `internal/aicontext/skills/**` is the only skill source of truth.
2. `AGENTS.md` `Skill Selection Guide` `Use` paths must never point to `.claude/skills/**`.
3. `CLAUDE.md` `Skill Selection Guide` `Use` paths must never point to `.agents/skills/**`.
4. Generated `Use` paths must omit `anvien/` between `skills/` and `<package>`.
5. Analyze must install both Codex and Claude repo skill outputs.
6. Setup must keep installing Codex skills to `$HOME/.agents/skills` and Claude Code skills to `$HOME/.claude/skills`.
7. Direct package layout must not delete or overwrite unrelated custom skills in `.agents/skills` or `.claude/skills`.
8. A foreign package root with the same name as an Anvien package is a collision unless it satisfies the exact-match adoption rule. The implementation must not silently overwrite it.
9. Stale Anvien-managed package roots must be deleted when the source package disappears.
10. Legacy `.claude/skills/anvien/**` and `.agents/skills/anvien/**` generated output must be removed or migrated after successful direct-root install.
11. Manifest data is ownership and sync evidence, not source of truth for skill content.
12. Tests must use synthetic fixtures for sync behavior; they must not hardcode the current real skill catalog as behavioral source of truth.
13. The shared installer must preserve unrelated custom skills for every setup target it touches, not only Codex and Claude.
14. Generated `AGENTS.md` and `CLAUDE.md` must not include volatile graph inventory counts in ordinary guidance.
15. A package root without manifest ownership can be adopted only when its file set exactly matches the desired Anvien package snapshot and contains no extra files.

## Technical Direction

### 1. Surface-specific rendering

Refactor AI context rendering so the path prefix is surface-specific:

```text
Codex surface path prefix: .agents/skills/
Claude surface path prefix: .claude/skills/
```

Expected direction:

- either add a `SkillGuideSurface` / `AgentSurface` argument to `renderAnvienBlock`;
- or add a smaller `skillGuideUseForPrefix(pkg, prefix)` helper and render two blocks from the same shared content template.

Do not duplicate the whole Anvien guide manually. Keep one shared source of guide text and parameterize only the skill path surface.

### 2. Remove volatile indexed-project sentence

Remove the generated sentence currently shaped like:

```text
This project is indexed by Anvien as **<repo>** (<symbols> symbols, <relationships> relationships, <flows> execution flows). Use Anvien to understand code, assess impact, navigate, audit graph quality, inspect resolution gaps, run query/accuracy benchmarks, and manage local/indexed repository intelligence.
```

Do not replace it with another sentence. The next generated content after the title should be the existing repo-agnostic guidance.

Required follow-up:

- review the `--no-stats` option because removing this sentence may make the option obsolete or reduce it to no-op behavior;
- update tests that currently assert stats/no-stats behavior in generated context;
- record token and stability impact in the benchmark ledger.

### 3. Scoped ownership sync

Current snapshot sync exact-mirrors the whole target directory. That is safe inside `.claude/skills/anvien`, but unsafe inside `.agents/skills` or `.claude/skills`.

Implementation must change the actual snapshot scope so it includes only:

- desired Anvien package roots from current source packages;
- previously managed package roots from `.anvien-skill-manifest.json`;
- explicitly known legacy generated roots such as `.claude/skills/anvien/**` and `.agents/skills/anvien/**`.

It must exclude unrelated custom roots:

```text
.agents/skills/my-team-skill/**
.claude/skills/my-team-skill/**
```

unless they are known Anvien-managed roots or collide with desired Anvien roots.

This must be implemented before analyze targets `.agents/skills` or direct `.claude/skills`.

New direct-root manifest locations:

```text
.agents/skills/.agents-skill-manifest.json
.claude/skills/.claude-skill-manifest.json
```

Legacy manifest locations, when present, are ownership evidence only:

```text
.claude/skills/anvien/.anvien-skill-manifest.json
.agents/skills/anvien/.anvien-skill-manifest.json
```

Do not copy stale legacy manifest entries blindly into the new manifest. The new manifest must be written from the current desired source snapshot after sync verification.

### 4. Collision handling

Before writing a direct package root, detect whether the root already exists without Anvien ownership evidence.

For a collision:

- do not overwrite the package;
- increment/report `collisions`;
- return a clear error;
- record the exact package root in evidence.

Fail clearly. Silent skip can make the generated guide point to missing or stale skill content.

Adoption rule:

- if a same-name package root has no manifest ownership but exactly matches the desired Anvien package file set and hashes, with no extra files, it may be adopted and recorded as `adopted`;
- if it differs in any path, hash, or extra file, it is a foreign collision and must not be overwritten silently.

### 5. Dual repo skill install during analyze

Change analyze-time installation from one Claude namespace:

```text
.claude/skills/anvien/**
```

to two direct target roots:

```text
.agents/skills/**
.claude/skills/**
```

The created/result summary should mention both surfaces separately, for example:

```text
.agents/skills/ (...)
.claude/skills/ (...)
```

Do not switch analyze to direct roots until scoped ownership sync and collision handling are in place.

### 6. Legacy namespace cleanup

After direct-root outputs are verified, clean old generated namespaces:

```text
.claude/skills/anvien/**
.agents/skills/anvien/**
```

Only remove these if they are generated Anvien namespaces or match the old managed output contract. Do not delete arbitrary user directories outside these legacy generated roots.

### 7. Managed gitignore precision

Review `internal/gitignore/managed.go`.

Current managed entries ignore `.claude/` and `.agents/` entirely. With direct repo skill layout, this may hide custom repo skills that users intentionally keep under `.agents/skills` or `.claude/skills`.

Preferred direction:

- keep generated Anvien outputs ignored;
- avoid using ignore rules as ownership rules;
- ignore only Anvien-generated package roots and manifest files when feasible, not all `.agents/` or `.claude/`;
- if broad `.agents/` or `.claude/` ignores are retained, record the reason in evidence because that choice hides user-owned repo skills unless users force-add them.

This item must be decided before the implementation slice is committed.

### 8. Secondary guidance/resource text

Update internal guidance surfaces that describe generated AI context and skills, including `internal/mcp/resources.go`.

Replace stale claims shaped like:

```text
.claude/skills/anvien/**
```

with surface-aware language:

```text
AGENTS.md, CLAUDE.md, .agents/skills/**, and .claude/skills/**
```

This is not a source-of-truth change for skills. It is a consistency update so MCP resources and generated guidance do not teach the old Claude-only namespace.

## Acceptance Criteria

1. `AGENTS.md` generated `Skill Selection Guide` paths use `.agents/skills/<package>/...`.
2. `CLAUDE.md` generated `Skill Selection Guide` paths use `.claude/skills/<package>/...`.
3. Neither generated guide contains `.claude/skills/anvien/`.
4. Neither generated guide contains `.agents/skills/anvien/`.
5. Analyze installs direct Codex repo skills under `.agents/skills/<package>/**`.
6. Analyze installs direct Claude repo skills under `.claude/skills/<package>/**`.
7. Setup continues to install Codex user skills under `$HOME/.agents/skills/<package>/**`.
8. Setup continues to install Claude Code user skills under `$HOME/.claude/skills/<package>/**`.
9. Shared setup installer behavior remains safe for Cursor and OpenCode skill targets.
10. Generated direct outputs mirror `internal/aicontext/skills/**` for Anvien-managed packages.
11. Custom non-Anvien skill roots under `.agents/skills` and `.claude/skills` survive analyze/setup sync.
12. Custom non-Anvien skill roots under shared setup targets such as Cursor/OpenCode survive setup sync.
13. A same-name custom package collision is not overwritten silently.
14. Deleted source packages remove corresponding Anvien-managed direct outputs on both surfaces.
15. Legacy `.claude/skills/anvien/**` and `.agents/skills/anvien/**` generated output is removed after migration.
16. Generated `AGENTS.md` and `CLAUDE.md` no longer contain the sentence beginning `This project is indexed by Anvien`.
17. Generated `AGENTS.md` and `CLAUDE.md` no longer contain volatile symbol, relationship, or execution-flow counts.
18. `--no-stats` behavior is removed, repurposed, or documented as no-op according to implementation findings.
19. Direct-root manifests are written to `.agents/skills/.agents-skill-manifest.json` and `.claude/skills/.claude-skill-manifest.json`.
20. Same-name unowned roots are adopted only when their file set exactly matches the desired package snapshot and has no extra files.
21. MCP resources and other internal guidance no longer describe generated skills as only `.claude/skills/anvien/**`.
22. Full build passes before tests.
23. Focused aicontext, setup, and analyze tests pass.
24. `anvien detect-changes --repo Anvien --scope all` runs before any implementation commit.

## Risk Notes

- Removing the namespace increases collision risk with user skills.
- Exact-mirror sync must be narrowed before targeting `.agents/skills` or `.claude/skills`.
- `verifySkillSnapshot` must verify only managed scope, not the entire target root.
- `removeEmptySkillDirs` must not remove custom directories or parent roots that still contain custom files.
- Generated `.gitignore` behavior may need to change so direct skill layout does not hide user-owned repo skill sources unintentionally.
- The shared installer is used by setup for more than Codex and Claude; ownership scoping must be generic.
- Removing the indexed-project sentence may make `Options.NoStats`, `--no-stats`, and no-stats tests obsolete. Do not leave misleading option behavior in place.
- Implicit adoption can be unsafe. Adoption is allowed only for exact path/hash matches with no extra files; otherwise fail as collision.
- Internal MCP resources can keep teaching the old layout if not updated with the generator change.
- Migration cleanup must not run before new direct outputs are verified.
- Tests must validate behavior with synthetic package sources and synthetic custom skill roots.

## Phase Checklist

- [x] [P0-A] Confirm plan approval and worktree hygiene.
  - Goal: start implementation only after this plan is approved.
  - Work Steps: get explicit approval; run `git status --short`; inspect any modified files; classify user changes versus implementation changes; do not revert unrelated user changes.
  - Implementation Gate: user has approved this plan and the worktree state is understood.
  - Acceptance: evidence ledger records approval and worktree status before code edits.

- [x] [P0-B] Refresh graph and map blast radius.
  - Goal: establish current graph evidence before editing generator or sync code.
  - Work Steps: run `anvien analyze --force`; run impact for `GenerateAIContextFiles`, `renderAnvienBlock`, `installBaseSkills`, `installSkillPackagesTo`, `actualSkillSnapshot`, `diffSkillSnapshots`, `applySkillSyncPlan`, `verifySkillSnapshot`, `setupInstallSkillsTo`, and managed gitignore helpers if edited.
  - Implementation Gate: impact output is reviewed and HIGH/CRITICAL is treated as scope warning.
  - Acceptance: evidence ledger records affected files, risks, and planned containment.

- [x] [P1-A] Introduce surface-specific guide path rendering.
  - Goal: make `AGENTS.md` and `CLAUDE.md` render different skill path prefixes from one guide source.
  - Work Steps: add a small surface/prefix abstraction; update `skillGuideUse` or equivalent helper to accept `.agents/skills/` or `.claude/skills/`; keep primary-entry routing unchanged.
  - Implementation Gate: impact for edited render helpers is complete.
  - Acceptance: unit test can render Codex and Claude blocks with different `Use` paths.

- [x] [P1-B] Generate distinct AGENTS and CLAUDE managed blocks.
  - Goal: write Codex-specific content to `AGENTS.md` and Claude-specific content to `CLAUDE.md`.
  - Work Steps: call the renderer once for Codex and once for Claude; upsert each file with its own content; preserve start/end markers and non-skill guide content.
  - Implementation Gate: P1-A helper exists and does not duplicate full guide text.
  - Acceptance: generated `AGENTS.md` contains `.agents/skills/`; generated `CLAUDE.md` contains `.claude/skills/`.

- [x] [P1-C] Remove volatile indexed-project sentence.
  - Goal: stop generated context from changing only because graph inventory counts changed.
  - Work Steps: remove the sentence beginning `This project is indexed by Anvien`; do not add a replacement sentence; inspect `Options.NoStats` and CLI `--no-stats`; update behavior/tests so no obsolete stats path remains.
  - Implementation Gate: impact for `renderAnvienBlock`, no-stats tests, and CLI flag owners is complete.
  - Acceptance: generated `AGENTS.md` and `CLAUDE.md` start with `# Anvien - Code Intelligence` followed by the existing repo-agnostic guidance, with no symbol/relationship/flow counts.

- [x] [P1-D] Scope snapshot sync to Anvien-managed package roots.
  - Goal: prevent direct-root sync from deleting custom skills.
  - Work Steps: change actual snapshot collection to include desired package roots, manifest-managed package roots, and known legacy generated roots only; exclude unrelated roots; update verify/delete/empty-dir cleanup to operate within managed scope.
  - Implementation Gate: synthetic fixtures cover custom skill roots before implementation is considered complete.
  - Acceptance: custom roots survive sync while stale Anvien-managed roots are removed.

- [x] [P1-E] Add collision detection for direct package roots.
  - Goal: avoid overwriting user skills with the same package name.
  - Work Steps: before writing a desired package root, detect existing root ownership; use manifest and exact snapshot hash evidence to classify Anvien-managed, adoptable, or foreign; fail clearly on foreign same-name collision.
  - Implementation Gate: P1-D ownership model exists.
  - Acceptance: test fixture with `.agents/skills/anvien-planner/SKILL.md` not owned by Anvien is not overwritten silently; exact matching unowned roots are adopted only when they contain no extra files.

- [x] [P1-F] Install repo skills to both direct surface roots.
  - Goal: make analyze produce native repo skill directories for both agents.
  - Work Steps: replace analyze-time `.claude/skills/anvien` install target with two installs: `.agents/skills` and `.claude/skills`; return both summaries; keep setup user-skill install behavior consistent.
  - Implementation Gate: scoped sync and collision handling are implemented.
  - Acceptance: analyze creates `.agents/skills/<package>/...` and `.claude/skills/<package>/...`.

- [x] [P1-G] Remove legacy generated namespaces after verified migration.
  - Goal: delete old `.claude/skills/anvien/**` and `.agents/skills/anvien/**` output after direct outputs are correct.
  - Work Steps: install and verify direct outputs first; then remove known legacy generated namespaces; handle `.claude/skills/anvien/**` and `.agents/skills/anvien/**` if they exist; keep errors explicit.
  - Implementation Gate: direct outputs verified on both surfaces.
  - Acceptance: analyze result has no stale `.claude/skills/anvien/**` or `.agents/skills/anvien/**` output.

- [x] [P1-H] Reconcile managed `.gitignore`.
  - Goal: ensure generated local output is ignored without making user-owned repo skills impossible to track.
  - Work Steps: inspect current managed entries; decide whether to keep broad `.agents/` and `.claude/` ignores or replace them with generated package-root ignores; update tests accordingly.
  - Implementation Gate: decision is recorded in evidence because this affects target repo developer workflow.
  - Acceptance: managed `.gitignore` behavior matches the chosen ownership model and tests explain it.

- [x] [P1-I] Update secondary guidance and MCP resource text.
  - Goal: remove stale Claude-only generated skill namespace claims from internal guidance surfaces.
  - Work Steps: update `internal/mcp/resources.go` and any other source text found by search that still says generated skills live only under `.claude/skills/anvien/**`; use surface-aware wording for Codex and Claude.
  - Implementation Gate: generator and install target direction is stable.
  - Acceptance: source search finds no stale `.claude/skills/anvien/**` guidance except explicit legacy-cleanup references and tests.

- [x] [P2-A] Add generated guide tests.
  - Goal: validate surface-specific paths in generated files.
  - Work Steps: update `internal/aicontext` tests to assert `AGENTS.md` uses `.agents/skills/<package>/...`, `CLAUDE.md` uses `.claude/skills/<package>/...`, neither guide uses `skills/anvien/`, and neither guide contains the removed indexed-project sentence or volatile counts.
  - Implementation Gate: code behavior exists first; tests are validation only.
  - Acceptance: focused aicontext tests pass.

- [x] [P2-B] Add sync behavior tests with synthetic packages.
  - Goal: validate direct-root sync safety independent of the real catalog.
  - Work Steps: create synthetic package fixtures; test edit/add/delete/rename/tamper/missing-output for managed packages; add custom-skill preservation; add same-name collision test; add exact-match adoption and extra-file non-adoption tests; cover at least one non-Codex/Claude setup target through the shared installer.
  - Implementation Gate: P1-D and P1-E implementation exists.
  - Acceptance: tests prove managed output syncs while foreign custom skills survive.

- [x] [P2-C] Add analyze and setup command tests.
  - Goal: validate CLI surfaces match real user workflows.
  - Work Steps: update analyze tests to expect `.agents/skills/<package>` and `.claude/skills/<package>`; update setup tests for Codex and Claude direct user skill roots; verify summaries.
  - Implementation Gate: implementation passes focused aicontext tests.
  - Acceptance: focused CLI analyze/setup tests pass.

- [x] [P3-A] Regenerate and record generated-output evidence.
  - Goal: confirm real repo output matches source after implementation.
  - Work Steps: run current source `anvien analyze --force`; inspect `AGENTS.md`, `CLAUDE.md`, `.agents/skills`, `.claude/skills`, direct-root manifests, absence of legacy namespace, and absence of the indexed-project sentence; record generated path counts and examples.
  - Implementation Gate: build has passed before validation tests.
  - Acceptance: evidence and benchmark ledgers record generated output inventory.

- [ ] [P3-B] Build, test, detect changes, and commit.
  - Goal: close the approved implementation slice safely.
  - Work Steps: run full build; run focused tests; run broader relevant Go tests; run `anvien detect-changes --repo Anvien --scope all`; record failures and fixes; commit only after validation passes.
  - Implementation Gate: all prior phases are complete.
  - Acceptance: commit hash is recorded in evidence and the worktree is clean except intentional untracked artifacts.
