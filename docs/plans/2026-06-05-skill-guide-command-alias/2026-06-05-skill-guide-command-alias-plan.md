# Plan

Title: Skill Guide Command Alias Column

Date: 2026-06-05

Status: Complete

Companion files:

- Evidence ledger: [2026-06-05-skill-guide-command-alias-evidence.md](2026-06-05-skill-guide-command-alias-evidence.md)
- Benchmark ledger: [2026-06-05-skill-guide-command-alias-benchmark.md](2026-06-05-skill-guide-command-alias-benchmark.md)

## Master Rules

1. Follow active repository instructions and generated `AGENTS.md`.
2. Write plan before coding and keep this plan updated as phases complete.
3. Use Anvien before graph-based implementation work; refresh with `anvien analyze --force`.
4. Run impact analysis before editing `renderAnvienBlock`, skill guide helpers, or tests around generated AI context.
5. HIGH or CRITICAL blast radius is a scope warning, not an edit prohibition.
6. Do not edit generated `AGENTS.md`, `CLAUDE.md`, `.agents/skills/**`, or `.claude/skills/**` as source of truth.
7. Source of truth remains `internal/aicontext/skills/**/SKILL.md` plus the generator code.
8. The command alias column must be generated projection only; do not add a manual alias registry.
9. Code first; update tests after behavior is implemented.
10. Run a full build before tests.
11. Record evidence and benchmarkable inventory counts as each phase completes.
12. Run `anvien detect-changes --repo Anvien --scope all` before implementation commit.

## Goal

Add a generated `Command` column to the generated `Skill Selection Guide` in both `AGENTS.md` and `CLAUDE.md`.

Target table shape:

```text
| When you need to... | Command | Use |
|---------------------|---------|-----|
| use when user ask to review spec | `/architect-review` | `.agents/skills/Architect-review/SKILL.md` |
```

The command alias must be derived from the current primary skill entry name. It must appear and disappear with the skill entry that generated it.

## Problem

The generated `Skill Selection Guide` currently has only two columns:

```text
| When you need to... | Use |
```

This forces agents to infer the slash-style invocation from the skill path or skill name. The user wants an explicit command alias column so each skill row advertises the command for that skill, while preserving current generated-output semantics:

- if a skill package disappears from `internal/aicontext/skills/**`, its row and command alias disappear;
- if a skill name changes in `SKILL.md`, the alias changes on the next generation;
- if a skill description changes, the `When you need to...` cell changes;
- if a skill path changes, the `Use` cell changes.

The alias must not become an independent source of truth.

## Scope

In scope:

- `internal/aicontext/aicontext.go` generated `Skill Selection Guide` table header and row rendering.
- `internal/aicontext/skill_packages.go` generated skill guide helper logic.
- `internal/aicontext/aicontext_test.go` tests for table shape, command aliases, primary-entry behavior, and generated guide stability.
- Generated `AGENTS.md` and `CLAUDE.md` output as validation artifacts only.
- Evidence and benchmark records for generated skill rows, alias counts, duplicate aliases, and validation commands.

Out of scope:

- Manual edits to generated `AGENTS.md` or `CLAUDE.md` as source of truth.
- Any slash-command execution system.
- Any manually maintained `commands.json`, alias registry, or per-skill alias file.
- Expanding nested child skills into the main generated guide.
- Changing skill install layout.
- Changing skill metadata parsing.
- Changing skill bodies or descriptions except as separately assigned work.

## Invariants

1. `internal/aicontext/skills/**/SKILL.md` remains the skill source of truth.
2. The generated command alias is derived from `primarySkillEntry(pkg).Name`.
3. The generated `Use` path remains derived from `primarySkillEntry(pkg).InstallPath`.
4. The generated `When you need to...` text remains derived from the primary package description.
5. A removed skill must remove its row and alias during regeneration.
6. A renamed skill must change its alias during regeneration.
7. The main generated guide must continue to show only each package primary entry.
8. Nested child skill entries must not appear as extra rows unless a future plan changes that invariant.
9. `AGENTS.md` aliases and `CLAUDE.md` aliases must be identical for the same skill package; only the `Use` path prefix differs by surface.
10. The table renderer must escape Markdown table cells consistently with existing `skillGuideNeed` and `skillGuideUse` behavior.
11. Duplicate generated aliases must be detectable in tests or benchmark evidence.
12. Alias generation must be deterministic across platforms.

## Technical Direction

### Command alias source

Add a generated helper near existing skill guide helpers:

```go
func skillGuideCommand(pkg SkillPackage) string
```

The helper must:

1. read `primarySkillEntry(pkg).Name`;
2. normalize the name into a slash command;
3. return the command in backticks for table rendering;
4. fall back to `pkg.Name` if the primary entry name normalizes to empty.

Normalization rules:

- lowercase ASCII letters;
- keep `a-z`, `0-9`, and `-`;
- convert spaces and underscores to `-`;
- treat other punctuation as a separator unless there is a stronger local reason to drop it;
- collapse repeated `-`;
- trim leading/trailing `-`;
- prefix `/`.

Required examples:

| Skill entry name | Command |
|---|---|
| `architect-review` | `/architect-review` |
| `System-architect` | `/system-architect` |
| `UI_taste skill` | `/ui-taste-skill` |
| `When Stuck - Problem-Solving Dispatch` | `/when-stuck-problem-solving-dispatch` |

### Table render

Update `renderAnvienBlock` so the generated skill guide table is:

```go
builder.WriteString("| When you need to... | Command | Use |\n")
builder.WriteString("|---------------------|---------|-----|\n")
for _, pkg := range packages {
    fmt.Fprintf(&builder, "| %s | %s | %s |\n",
        skillGuideNeed(pkg),
        skillGuideCommand(pkg),
        skillGuideUse(pkg, skillPathPrefix),
    )
}
```

Do not duplicate separate Codex/Claude table logic. Keep the command helper surface-independent.

### Tests

Update focused tests in `internal/aicontext/aicontext_test.go`:

- assert the new header `| When you need to... | Command | Use |`;
- assert the separator row has three columns;
- assert an existing real skill row includes a generated command, for example `/architect-review`;
- assert the command is derived from skill `name`, not from folder casing;
- assert the `problem-solving` package still shows only the primary package entry command and does not expose child entry commands in the main guide;
- assert generated `AGENTS.md` uses `.agents/skills/...` and generated `CLAUDE.md` uses `.claude/skills/...` while command aliases stay the same.

Add unit coverage for normalization if local test style allows a small helper test. If helper visibility stays package-private, test it in the same Go package.

## Blast Radius

Anvien impact checks during planning:

- `renderAnvienBlock`: CRITICAL, affects `internal/aicontext/aicontext.go`, CLI analyze postrun/command flows, generated `AGENTS.md`/`CLAUDE.md`.
- `skillGuideNeed`: HIGH, affects generated skill guide rows through `renderAnvienBlock`.
- `skillGuideUse`: HIGH, affects generated skill guide rows through `renderAnvienBlock`.
- `primarySkillEntry`: central source for primary package entry selection; command alias must preserve this invariant.

Risk handling:

- Treat HIGH/CRITICAL as blast-radius warnings.
- Keep edits scoped to `aicontext.go`, `skill_packages.go`, and focused tests unless implementation evidence proves another file is required.
- Do not change package discovery, installation, or metadata parsing behavior.

## Phase Checklist

- [x] P0-A: Confirm codebase facts and impact before editing.
  - Goal: establish exact source surfaces and blast radius.
  - Work Steps:
    1. Run `anvien analyze --force`.
    2. Run impact checks for `renderAnvienBlock`, `skillGuideNeed`, `skillGuideUse`, and `primarySkillEntry`.
    3. Read the current table renderer and tests around `Skill Selection Guide`.
  - Implementation Gate: impact evidence is recorded in the evidence ledger.
  - Acceptance: plan evidence names the touched source files, linked tests, and risk level.

- [x] P1-A: Add deterministic command alias generation.
  - Goal: derive a command alias from current primary skill entry name.
  - Work Steps:
    1. Add `skillGuideCommand(pkg SkillPackage) string`.
    2. Add a small normalizer helper if needed to keep `skillGuideCommand` readable.
    3. Ensure empty normalized names fall back to `pkg.Name`.
    4. Escape the command cell consistently with table helper behavior.
  - Implementation Gate: no package discovery or install code is changed.
  - Acceptance: normalization examples pass in focused tests.

- [x] P2-A: Add the `Command` column to generated guide rendering.
  - Goal: render `When you need to... | Command | Use` for both `AGENTS.md` and `CLAUDE.md`.
  - Work Steps:
    1. Update the generated table header.
    2. Update the separator row.
    3. Update the row formatter to include `skillGuideCommand(pkg)`.
    4. Confirm only path prefix remains surface-specific.
  - Implementation Gate: `skillGuideNeed` and `skillGuideUse` behavior remains unchanged except for table placement.
  - Acceptance: generated blocks contain the three-column table with one command per package row.

- [x] P3-A: Update focused tests for shape, aliases, and primary-entry behavior.
  - Goal: prevent drift between skill source, command alias, and generated output.
  - Work Steps:
    1. Update existing generated guide assertions.
    2. Add/extend normalization test cases.
    3. Assert `/architect-review` appears for the `Architect-review` skill row.
    4. Assert `/problem-solving` appears while child commands such as `/collision-zone-thinking` do not appear in the main generated guide.
    5. Assert duplicate command aliases are detected or at least counted in benchmark evidence.
  - Implementation Gate: tests must not hardcode all current skill rows as the behavioral source of truth.
  - Acceptance: focused tests fail before the implementation and pass after it.

- [x] P4-A: Validate generated output and command alias inventory.
  - Goal: prove normal generation produces correct `AGENTS.md` and `CLAUDE.md` content.
  - Work Steps:
    1. Run full build before tests.
    2. Run focused and full Go tests.
    3. Run normal analyze/generation path.
    4. Inspect generated `Skill Selection Guide` sections in both files.
    5. Count skill rows, command aliases, and duplicate aliases.
    6. Run `anvien detect-changes --repo Anvien --scope all`.
  - Implementation Gate: generated files are not manually edited as source of truth.
  - Acceptance: evidence and benchmark ledgers contain validation commands, row counts, alias counts, duplicate count, and detect-changes summary.

## Validation Matrix

| Surface | Required evidence |
|---|---|
| `internal/aicontext/skill_packages.go` | helper tests for command normalization and primary-entry alias source |
| `internal/aicontext/aicontext.go` | generated table header and row render tests |
| `AGENTS.md` generated output | `.agents/skills/...` use paths and generated command aliases |
| `CLAUDE.md` generated output | `.claude/skills/...` use paths and same generated command aliases |
| Nested packages | primary-entry command only in main guide |
| Deleted/renamed skill behavior | source-derived invariant described and covered by package-driven tests or fixture |
| Duplicate aliases | measured duplicate count recorded; test if feasible |

## Definition Of Done

This plan is complete when:

1. generated `Skill Selection Guide` has columns `When you need to...`, `Command`, and `Use`;
2. each skill package row has exactly one generated command alias;
3. command aliases are derived from current primary skill entry names;
4. no manual alias registry exists;
5. generated aliases disappear or change when source skill entries disappear or change;
6. generated `Use` paths remain surface-specific for `AGENTS.md` and `CLAUDE.md`;
7. primary-entry behavior for nested packages remains unchanged;
8. focused tests and full validation commands pass;
9. benchmark ledger records skill row count, command alias count, and duplicate alias count;
10. `anvien detect-changes --repo Anvien --scope all` is recorded before commit.
