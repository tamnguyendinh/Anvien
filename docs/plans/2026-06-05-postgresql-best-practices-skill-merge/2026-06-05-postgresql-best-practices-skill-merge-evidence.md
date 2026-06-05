# Evidence Ledger

Title: PostgreSQL Best Practices Skill Merge

Date: 2026-06-05

Status: Planned

Companion files:

- Plan: [2026-06-05-postgresql-best-practices-skill-merge-plan.md](2026-06-05-postgresql-best-practices-skill-merge-plan.md)
- Benchmark ledger: [2026-06-05-postgresql-best-practices-skill-merge-benchmark.md](2026-06-05-postgresql-best-practices-skill-merge-benchmark.md)

## Evidence Rules

1. This file records planning and future implementation evidence.
2. Source-document facts must be separated from validation results.
3. Build/test pass-fail belongs here unless timing, size, or inventory count is the measured target.
4. Benchmark numbers and inventory counts belong in the benchmark ledger.
5. Append evidence as each checklist item is completed.

## E0 - User Request

User asked how to merge:

```text
E:\owner-tool\clockwork\skills\SKILL_POSTGRESQL.md
```

into:

```text
internal\aicontext\skills\databases
```

After comparison, the selected direction is:

- add a new PostgreSQL reference under `internal/aicontext/skills/databases/references`;
- update `internal/aicontext/skills/databases/SKILL.md` navigation;
- do not replace the database skill and do not add a `rules/` directory in this slice.

## E1 - Planning Commands

Commands run before creating this plan:

```text
anvien analyze --force
anvien query files "database skill postgresql references best practices" --repo Anvien
git status --short
```

Observed planning facts:

- `anvien analyze --force` completed successfully.
- Graph path: `E:\Anvien\.anvien\graph.json`.
- Graph inventory from analyze: 83,548 nodes and 122,297 relationships.
- `git status --short` returned no output before this plan was created.
- Anvien query identified AI context skill package code as high-risk source when changing generated skill packaging, but this plan intentionally scopes implementation to skill reference documentation instead.

## E2 - Source Inspection Facts

Existing Anvien database skill:

```text
internal/aicontext/skills/databases/SKILL.md
internal/aicontext/skills/databases/references/*.md
internal/aicontext/skills/databases/scripts/*.py
```

Observed structure:

- `SKILL.md` is the skill manifest and navigation entry point.
- `references/` contains 8 files.
- PostgreSQL currently has 4 reference files:
  - `postgresql-administration.md`
  - `postgresql-performance.md`
  - `postgresql-psql-cli.md`
  - `postgresql-queries.md`
- The skill also contains MongoDB references and Python utility scripts.

Clockwork PostgreSQL source document:

```text
E:\owner-tool\clockwork\skills\SKILL_POSTGRESQL.md
```

Observed content:

- It is a demo README, not a standard skill manifest.
- It is focused on PostgreSQL performance anti-patterns.
- It includes three high-value examples:
  - missing foreign key index;
  - JOIN behavior before/after foreign key index;
  - partial index for filtered workload.
- It includes benchmark-style before/after evidence.
- It includes demo-only content that should not be merged verbatim.

## E3 - Current Risk Evidence

Content risk:

- Copying the source file verbatim would introduce demo setup details and Claude-specific instructions that do not match the Anvien database skill.

Performance-claim risk:

- The source speedups are valid as example benchmark evidence from one workload, but they are not universal guarantees.

Skill-architecture risk:

- Creating `rules/` would change the current structure without a full rule catalog or rule-loading behavior.

Containment:

- Keep implementation inside `internal/aicontext/skills/databases/SKILL.md` and one new reference file.
- Do not touch generated context files directly.
- Use source-document validation checks to prevent demo-only strings from leaking into the merged reference.

## E4 - Implementation Evidence

Implementation started after user command:

```text
triên khai plan
```

P0-A approval and worktree hygiene:

- User explicitly asked to implement the plan.
- `git status --short` showed this plan directory as untracked.
- `git status --short` also showed unrelated untracked skill directories under `internal/aicontext/skills`, including `Architect-review`, `Data-Integrity`, `Edge-Case`, `System-Architect`, and `supervisor`.
- Those unrelated untracked skill directories were not created or edited by this slice and must not be staged.

P0-B graph refresh and impact:

```text
anvien analyze --force
```

Result:

- Failed before graph refresh completed because unrelated untracked skill package directories do not contain standard `SKILL.md` entries.
- First failure: `skill package "supervisor" has no SKILL.md entry`.
- Later failure after attempted excludes: `skill package "Architect-review" has no SKILL.md entry`.

Containment:

- The failed analyze is recorded as an external worktree blocker caused by unrelated untracked directories.
- `anvien status` reported the current index as up-to-date for commit `84af8e9`.
- File-level impact was run against the current up-to-date index before source edits.

```text
anvien impact file internal/aicontext/skills/databases/SKILL.md --repo Anvien --direction upstream
```

Impact result:

- Risk: LOW.
- Affected file count: 0.
- Affected process count: 0.
- Linked tests: 0.
- `SKILL.md` has outbound reference edges to the existing 8 database reference files.

Pending evidence items:

- P2-A content validation.
- P2-B build, tests, analyze, detect-changes.
- P3-A commit hash and final worktree state.

P1-A new reference file:

- Added `internal/aicontext/skills/databases/references/postgresql-best-practices.md`.
- New reference final length after concision pass: 142 physical lines, 115 non-empty lines, 658 words, 4133 characters.
- The reference covers:
  - indexing foreign key columns;
  - supporting JOIN-heavy queries with indexes;
  - using partial indexes for filtered workloads.
- The reference preserves the source benchmark values for the three Clockwork examples while labeling them as source workload benchmark evidence.

P1-B `SKILL.md` navigation update:

- Updated `internal/aicontext/skills/databases/SKILL.md`.
- Added a PostgreSQL References link to `references/postgresql-best-practices.md`.
- Existing MongoDB reference links and script navigation were not changed.

P2-A content validation:

```text
rg -n "~/.claude|Claude Code|demo-postgres-best-practices\.sql|postgres-best-practices Skill Structure|Copy vào|Cài đặt|Chạy Demo" internal/aicontext/skills/databases/SKILL.md internal/aicontext/skills/databases/references/postgresql-best-practices.md
```

Result:

- No matches.

Required rule and benchmark checks:

- `Rule 1: Index Foreign Key Columns` present.
- `Rule 2: Support JOIN-Heavy Queries` present.
- `Rule 3: Use Partial Indexes For Filtered Workloads` present.
- Benchmark values present: `9.17 ms`, `0.11 ms`, `11.86 ms`, `0.25 ms`, `10.39 ms`, `0.10 ms`.

Navigation and unchanged-area checks:

- `SKILL.md` contains `postgresql-best-practices.md`.
- `git diff --name-only -- internal/aicontext/skills/databases/references/mongodb-*.md internal/aicontext/skills/databases/scripts` returned no output.
- `git diff --check -- internal/aicontext/skills/databases docs/plans/2026-06-05-postgresql-best-practices-skill-merge` returned no output.

P2-B build, tests, analyze, and detect-changes:

Main worktree build:

| Command | Result |
|---|---|
| `go build ./...` | Failed on existing non-buildable fixtures under `anvien/test/fixtures` (`models`, `animal`, mixed package fixture, C fixture). Not caused by this change. |
| `go build ./cmd/... ./internal/...` | Passed. |

Main worktree focused test:

| Command | Result |
|---|---|
| `go test ./internal/aicontext` | Failed because unrelated untracked skill directories such as `Architect-review` do not contain standard `SKILL.md` entries. |

Clean validation worktree:

- Created a temporary git worktree under `E:\Anvien\.tmp\postgresql-best-practices-skill-merge-worktree`.
- Applied the staged slice diff into that clean worktree.
- This avoided unrelated untracked skill directories in the main worktree while keeping temporary files inside the repo as required.

| Command | Result |
|---|---|
| `go build ./cmd/... ./internal/...` | Passed in clean validation worktree. |
| `go test ./internal/aicontext` | Passed in clean validation worktree: `ok ... 17.765s`. |
| `anvien analyze --force` | Passed in clean validation worktree. |

Clean validation analyze inventory:

- Files scanned: 1369.
- Parsed code files: 697.
- Graph nodes: 83582.
- Graph relationships: 122338.
- Failed files: 0.

Detect-changes:

```text
anvien detect-changes --repo Anvien --scope all
```

Result:

- Failed because the clean validation worktree registered another `Anvien` repo name, making the registry alias ambiguous.

```text
anvien detect-changes --repo 'E:\Anvien' --scope all
```

Result:

- Passed.
- Changed files: 5.
- Affected files: 4.
- Risk level: low.
- Changed app layer: docs.
- Changed functional area: documentation.
