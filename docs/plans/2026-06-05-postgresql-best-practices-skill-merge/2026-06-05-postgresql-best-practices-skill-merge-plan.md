# Plan

Title: PostgreSQL Best Practices Skill Merge

Date: 2026-06-05

Status: Complete

Companion files:

- Evidence ledger: [2026-06-05-postgresql-best-practices-skill-merge-evidence.md](2026-06-05-postgresql-best-practices-skill-merge-evidence.md)
- Benchmark ledger: [2026-06-05-postgresql-best-practices-skill-merge-benchmark.md](2026-06-05-postgresql-best-practices-skill-merge-benchmark.md)

## Master Rules

1. Follow active repository instructions and generated `AGENTS.md`.
2. Use Anvien before graph-based plan or implementation work and refresh the graph before graph reads.
3. Run impact analysis before editing shared code, handlers, generated contracts, functions, classes, or methods.
4. For this documentation-focused slice, run file-level Anvien impact before editing existing skill source files.
5. HIGH or CRITICAL blast radius is a scope warning, not an edit ban.
6. Do not edit generated `AGENTS.md` or `CLAUDE.md` as the permanent source of truth.
7. Source of truth for this change is `internal/aicontext/skills/databases/**`.
8. Keep temporary directories, if any are needed, inside the working repo.
9. Code or source-document content is updated before validation artifacts.
10. Run a build before tests if tests are run.
11. No Playwright e2e is required because this plan does not change UI behavior.
12. Record evidence and benchmarkable inventory as each implementation slice completes.
13. Run `anvien detect-changes --repo Anvien --scope all` before any implementation commit.
14. Commit the completed implementation slice after validation, then record the commit hash in evidence.

## Goal

Merge the useful PostgreSQL performance guidance from:

```text
E:\owner-tool\clockwork\skills\SKILL_POSTGRESQL.md
```

into the existing Anvien database skill:

```text
internal/aicontext/skills/databases
```

The merged result should preserve the current `databases` skill architecture while adding a focused PostgreSQL best-practices reference with benchmark-backed rules for common indexing anti-patterns.

## Problem

The existing `databases` skill is a broad MongoDB and PostgreSQL skill. Its PostgreSQL references cover queries, psql CLI usage, performance, and administration, but the high-impact index anti-pattern guidance is spread out and lacks the concrete before/after benchmark evidence from the Clockwork document.

The Clockwork file contains valuable PostgreSQL performance examples, but it is shaped as a Vietnamese demo README. It includes Claude Code installation instructions, demo environment details, and claims about a separate `postgres-best-practices` skill structure that do not match the current `databases` skill layout.

Merging it verbatim would pollute the skill manifest and confuse the reference model. The content needs to be normalized into the existing `references/` pattern.

## Approved Direction

Use a new reference file:

```text
internal/aicontext/skills/databases/references/postgresql-best-practices.md
```

Update the PostgreSQL reference list in:

```text
internal/aicontext/skills/databases/SKILL.md
```

Do not create a `rules/` directory in this slice. The current database skill uses `references/` plus `scripts/`, and this merge should stay within that structure.

## Scope

In scope:

- Add `references/postgresql-best-practices.md`.
- Convert the Clockwork content from demo README style into a production reference.
- Keep the strongest rules:
  - index foreign key columns when they are used for joins, lookups, or cascading actions;
  - support JOIN-heavy queries with the right foreign key and filter indexes;
  - use partial indexes for filtered workloads such as soft deletes, active records, or pending queues.
- Preserve the before/after benchmark evidence in compact tables.
- Add practical guardrails: verify with `EXPLAIN ANALYZE`, check selectivity, account for write overhead, and prefer safe production index creation patterns where appropriate.
- Update `SKILL.md` so agents can discover the new reference.
- Record implementation evidence, benchmark inventory, validation, detect-changes, and commit hash.

Out of scope:

- Replacing the entire `databases` skill.
- Adding a standalone `postgres-best-practices` skill.
- Adding a `rules/` directory before a real rule catalog exists.
- Changing MongoDB references.
- Changing Python utility scripts.
- Changing generated `AGENTS.md` or `CLAUDE.md` directly.
- Adding UI, browser, or Playwright validation.
- Reproducing the Clockwork demo database.

## Content Requirements

1. The new reference must be written in English to match the surrounding skill files.
2. The document must read as a durable reference, not a demo pitch.
3. Remove demo-only content:
   - Claude Code install instructions;
   - `~/.claude/skills` setup instructions;
   - `demo-postgres-best-practices.sql` run instructions;
   - separate skill-structure claims unless reframed as future optional work;
   - Vietnamese marketing or demo framing.
4. Keep source benchmark values, but label them as example benchmark evidence from the source workload, not universal guarantees.
5. Include direct SQL examples for before/after rules.
6. Mention that indexes can slow writes and increase storage, so they should be validated against workload and query plans.
7. Keep the file concise enough for skill use, with links from `SKILL.md` doing navigation.

## Acceptance Criteria

1. `internal/aicontext/skills/databases/references/postgresql-best-practices.md` exists.
2. `internal/aicontext/skills/databases/SKILL.md` links to the new reference in the PostgreSQL References section.
3. The new reference contains sections for foreign key indexes, JOIN support indexes, and partial indexes.
4. The new reference contains the source benchmark improvements for the three rules.
5. The new reference does not contain demo-only install strings such as `~/.claude/skills` or `demo-postgres-best-practices.sql`.
6. Existing MongoDB references and scripts remain unchanged.
7. Anvien graph is refreshed before graph-based validation.
8. File-level Anvien impact for the existing skill manifest/reference entry point is recorded before editing.
9. Validation records at least:
   - link presence in `SKILL.md`;
   - new reference file presence;
   - demo-only string absence checks;
   - `anvien analyze --force`;
   - build/test command results selected for this doc-only slice.
10. `anvien detect-changes --repo Anvien --scope all` runs before commit.
11. The implementation slice is committed and the commit hash is recorded in evidence.

## Risk Notes

- Verbatim merge would introduce stale or misleading Claude-specific setup guidance into a Codex/Anvien skill.
- Benchmark speedups from one demo workload are useful evidence but not universal performance guarantees.
- Adding indexes blindly can hurt write-heavy systems; the reference must include workload validation guidance.
- The current `databases` skill is shared by MongoDB and PostgreSQL workflows, so `SKILL.md` should stay compact and navigational.
- The new reference should not imply that Anvien has a rule engine for these checks unless that is implemented separately.
- Because this is source documentation under the skill catalog, generated `AGENTS.md` and `CLAUDE.md` should be regenerated by normal analyze flow only if implementation needs it; they should not be edited directly.

## Phase Checklist

- [x] [P0-A] Confirm approval and worktree hygiene.
  - Goal: start implementation only after this plan is accepted.
  - Work Steps: get explicit approval to implement; run `git status --short`; inspect any modified or untracked files; classify unrelated user changes; do not revert unrelated changes.
  - Implementation Gate: user approval and cleanly understood worktree state.
  - Acceptance: evidence ledger records approval and worktree status before source edits.

- [x] [P0-B] Refresh graph and record file-level impact.
  - Goal: establish Anvien evidence before editing existing skill source.
  - Work Steps: run `anvien analyze --force`; run `anvien impact file internal/aicontext/skills/databases/SKILL.md --repo Anvien --direction upstream`; optionally run file context or query commands if the relevant skill structure is unclear.
  - Implementation Gate: impact output has been reviewed and any HIGH/CRITICAL blast radius is treated as a scope warning.
  - Acceptance: evidence ledger records graph refresh, impact output summary, and containment notes.

- [x] [P1-A] Draft the PostgreSQL best-practices reference.
  - Goal: convert the Clockwork demo into a durable reference file.
  - Work Steps: create `references/postgresql-best-practices.md`; write sections for when to use the reference, rule 1 foreign key indexes, rule 2 JOIN support indexes, rule 3 partial indexes, benchmark summary, review checklist, and cautions; translate and normalize source examples to English.
  - Implementation Gate: no changes to `SKILL.md` yet; the reference content must be coherent as a standalone file.
  - Acceptance: the new reference exists and contains the three required rules with compact before/after SQL and benchmark tables.

- [x] [P1-B] Update database skill navigation.
  - Goal: make the new reference discoverable from the skill entry point.
  - Work Steps: update `internal/aicontext/skills/databases/SKILL.md` PostgreSQL References with a link to `postgresql-best-practices.md`; keep the quick-start and existing references unchanged except for the new navigation row; optionally add one concise PostgreSQL best-practice bullet if needed.
  - Implementation Gate: P1-A reference exists and has the final filename.
  - Acceptance: `SKILL.md` links to the new reference and the link path matches the actual file.

- [x] [P2-A] Validate content shape.
  - Goal: prove the merge is clean and not a pasted demo.
  - Work Steps: search the new reference and `SKILL.md` for `~/.claude`, `Claude Code`, `demo-postgres-best-practices.sql`, and other demo-only strings; verify rule headings and benchmark values; verify no MongoDB reference files changed.
  - Implementation Gate: source edits are complete.
  - Acceptance: evidence ledger records pass/fail results for content checks and any intentional retained source wording.

- [x] [P2-B] Run build, tests, analyze, and detect changes.
  - Goal: validate repo health for the doc-only skill update.
  - Work Steps: run a full or approved build command before tests; run focused relevant tests if available; run `anvien analyze --force`; run `anvien detect-changes --repo Anvien --scope all`; record failures, existing unrelated failures, and final pass/fail state.
  - Implementation Gate: content validation from P2-A is complete.
  - Acceptance: evidence ledger records all commands and outcomes; benchmark ledger records updated inventory counts.

- [x] [P3-A] Commit the implementation slice.
  - Goal: close the completed merge with a traceable commit.
  - Work Steps: inspect `git diff`; stage only files for this slice; commit with a concise message; record the commit hash; confirm post-commit worktree state.
  - Implementation Gate: P2-B validation and detect-changes are complete.
  - Acceptance: evidence ledger records commit hash and final worktree status.
