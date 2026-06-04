# Evidence Ledger

Title: AI Context Skill Description Trigger Shortening

Date: 2026-06-04

Status: Draft - awaiting implementation

Companion files:

- Plan: [2026-06-04-aicontext-skill-description-triggers-plan.md](2026-06-04-aicontext-skill-description-triggers-plan.md)
- Benchmark ledger: [2026-06-04-aicontext-skill-description-triggers-benchmark.md](2026-06-04-aicontext-skill-description-triggers-benchmark.md)

## Evidence Rules

1. Evidence records facts read or commands run.
2. Generated `AGENTS.md`, `CLAUDE.md`, and `.claude/skills/anvien/**` are validation output, not source of truth.
3. Tests validate behavior after source changes; tests do not define the desired description text.
4. Record failures and process corrections, including accidental pre-plan edits.
5. Build/test pass-fail belongs here; measured counts and sizes belong in the benchmark ledger.

## E0 - User Report

Status: recorded

The user identified that generated skill descriptions are too long and should be trigger-only. The clearest stated example:

```text
Use when the user asks to debug.
```

The user also clarified that a request to write a plan is a continuation of the current planning task and must activate `anvien-planner` to create a plan artifact, not merely print a plan in chat.

The user further clarified that these skills are for every repository that uses Anvien, not only for fixing or operating the Anvien product repository. The plan must therefore treat descriptions as target-repo task triggers. Any wording that reads like "this skill is for working on Anvien itself" is wrong unless the task is specifically about Anvien product code.

During plan review, the user rejected `ai-multimodal` and requested complete removal:

```text
Remove this skill completely; it is not necessary and has no value.
```

## E1 - Planner Skill Loaded

Status: recorded

`anvien-planner` was loaded from:

```text
.claude/skills/anvien/anvien-planner/SKILL.md
```

Relevant planner requirements read:

- standard three-file plan set under `docs/plans/YYYY-MM-DD-<slug>/`;
- plan file controls work;
- evidence file explains why work is known correct;
- benchmark file records measured numbers;
- checklist items must be complete mini-plans.

## E2 - Nearby Plan Inspection

Status: recorded

Nearby plans inspected:

```text
docs/plans/2026-06-02-anvien-skill-selection-guide-reduction/2026-06-02-anvien-skill-selection-guide-reduction-plan.md
docs/plans/2026-06-04-anvien-skill-mirror-incremental-sync/2026-06-04-anvien-skill-mirror-incremental-sync-plan.md
```

Finding:

- The 2026-06-02 plan reduced Anvien workflow skills too narrowly for the current catalog.
- The current plan keeps the current skill catalog and changes description shape only.
- The 2026-06-04 mirror-sync plan confirms `internal/aicontext/skills/**` is source of truth and `.claude/skills/anvien/**` is generated output.

## E3 - Anvien Graph And Query Evidence

Status: recorded

Commands run before creating this plan:

```text
anvien analyze --force
anvien query files "skill description Skill Selection Guide aicontext generated AGENTS CLAUDE" --repo Anvien
```

`--repo Anvien` is used here because this implementation changes the bundled catalog in the Anvien product repository. It is not part of the generated skill trigger contract for target repositories.

Observed relevant files from query:

- `internal/aicontext/skill_packages.go`
- `internal/aicontext/aicontext_test.go`

Source inspection confirmed:

- `readSkillPackage` reads `description` from `SKILL.md` frontmatter.
- `pkg.Description` is set from `primarySkillEntry(pkg).Description`.
- `skillGuideNeed(pkg)` renders the package description into the generated `Skill Selection Guide`.
- `skillGuideUse(pkg)` renders generated skill entry paths.

## E4 - Current Process Correction

Status: recorded

Before this plan artifact was written, two source description edits were made prematurely:

```text
internal/aicontext/skills/anvien-planner/SKILL.md
internal/aicontext/skills/anvien-debugging/SKILL.md
```

Current required correction:

- do not commit these edits before the plan is approved;
- reconcile them in P0-A;
- if accepted, treat them as the first two implementation slices and validate normally;
- if rejected, revert only these agent-created edits before proceeding.

## E5 - Pending Validation

Status: pending

To be filled during implementation:

- per-skill source diff;
- generated row verification in `AGENTS.md` and `CLAUDE.md`;
- generated mirror verification under `.claude/skills/anvien/**`;
- full build result;
- focused AI-context test result;
- final `anvien detect-changes --repo Anvien --scope all`;
- commit hashes.

## E6 - P0-A Reconcile Pre-Plan Source Edits

Status: recorded

Decision:

- Keep the pre-plan edits to `internal/aicontext/skills/anvien-debugging/SKILL.md` and `internal/aicontext/skills/anvien-planner/SKILL.md`.
- Treat them as the first two implementation slices because both match the approved target descriptions in the plan.

Source diffs:

- `anvien-debugging`: description changed to `Use when the user asks to debug.`
- `anvien-planner`: description changed to `Use when the user asks to create, write, or review a docs/plans plan.`

Generated-row verification after regeneration:

```text
AGENTS.md: anvien-debugging -> Use when the user asks to debug.
AGENTS.md: anvien-planner -> Use when the user asks to create, write, or review a docs/plans plan.
CLAUDE.md: anvien-debugging -> Use when the user asks to debug.
CLAUDE.md: anvien-planner -> Use when the user asks to create, write, or review a docs/plans plan.
.claude/skills/anvien/anvien-debugging/SKILL.md mirrors source description.
.claude/skills/anvien/anvien-planner/SKILL.md mirrors source description.
```

## E7 - P0-B Baseline And P1 Validation

Status: recorded

Commands run:

```text
go run ./cmd/anvien analyze --force
go build ./cmd/... ./internal/...
go test ./internal/aicontext -count=1
go test ./internal/cli -run "TestAnalyzeCommand|AIContext|Aicontext" -count=1
```

Results:

- Analyze passed and regenerated AI context from current source.
- Full product build passed.
- `internal/aicontext` focused tests passed.
- `internal/cli` focused analyze/AI-context tests passed.

## E8 - P1 Detect Changes

Status: recorded

Command run before P1 commit:

```text
go run ./cmd/anvien detect-changes --repo Anvien --scope all
```

Result summary:

- `risk_level`: `low`
- `changed_files`: 5
- `affected_files`: 3
- `affected_count`: 0
- `resolutionHealthImpact.degradedNodes`: 0
- changed source skill paths were reported as `internal/aicontext/skills/anvien-debugging/skill.md` and `internal/aicontext/skills/anvien-planner/skill.md` due to path case normalization in the detector output.

## E9 - P1-C Context Engineering

Status: recorded

Source change:

- `internal/aicontext/skills/context-engineering/SKILL.md`
- Replaced folded capability-summary description with `Use when the user asks to design or improve AI-agent context.`

Impact:

- `anvien impact file internal/aicontext/skills/context-engineering/SKILL.md --repo Anvien --direction upstream`
- Risk: `LOW`
- Affected files: 0
- Affected processes: 0

Generated-row verification after regeneration:

```text
AGENTS.md: context-engineering -> Use when the user asks to design or improve AI-agent context.
CLAUDE.md: context-engineering -> Use when the user asks to design or improve AI-agent context.
.claude/skills/anvien/context-engineering/SKILL.md mirrors source description.
```

Validation:

```text
go run ./cmd/anvien analyze --force
go build ./cmd/... ./internal/...
go test ./internal/aicontext -count=1
go test ./internal/cli -run "TestAnalyzeCommand|AIContext|Aicontext" -count=1
```

Results:

- Analyze/regeneration passed.
- Full product build passed.
- Focused AI-context and CLI tests passed.

Detect changes:

- Command: `go run ./cmd/anvien detect-changes --repo Anvien --scope all`
- `risk_level`: `low`
- `changed_files`: 14
- `affected_files`: 5
- `affected_count`: 0

Detect changes:

- Command: `go run ./cmd/anvien detect-changes --repo Anvien --scope all`
- `risk_level`: `low`
- `changed_files`: 7
- `affected_files`: 6
- `affected_count`: 0

## E16 - P3-A Concise Single-Entry Group

Status: recorded

Source changes:

- Updated frontmatter `description` only for 11 packages:
  `anvien-api-surface`, `anvien-refactoring`, `chrome-devtools`, `devops`, `frontend-design`, `mcp-builder`, `payment-integration`, `sequential-thinking`, `skill-creator`, `threejs`, and `web-testing`.
- Skill bodies were not changed.

Impact:

- Command pattern: `go run ./cmd/anvien impact file internal/aicontext/skills/<skill>/SKILL.md --repo Anvien --direction upstream`
- All 11 files reported `LOW`.
- All 11 files reported 0 affected files and 0 affected flows.

Generated-row verification after regeneration:

```text
anvien-api-surface -> Use when the user asks to inspect API or MCP surfaces.
anvien-refactoring -> Use when the user asks to refactor code.
chrome-devtools -> Use when the user asks to automate or inspect a browser.
devops -> Use when the user asks to deploy or operate infrastructure.
frontend-design -> Use when the user asks to design a frontend UI.
mcp-builder -> Use when the user asks to build an MCP server.
payment-integration -> Use when the user asks to integrate payments.
sequential-thinking -> Use when the user asks to reason step by step.
skill-creator -> Use when the user asks to create or update a skill.
threejs -> Use when the user asks to build 3D web experiences with Three.js.
web-testing -> Use when the user asks to test web behavior.
```

- Each source description matched the generated `.claude/skills/anvien/<skill>/SKILL.md` mirror.

Validation:

```text
go run ./cmd/anvien analyze --force
go build ./cmd/... ./internal/...
go test ./internal/aicontext -count=1
go test ./internal/cli -run "TestAnalyzeCommand|AIContext|Aicontext" -count=1
```

Results:

- Analyze/regeneration passed.
- Full product build passed.
- Focused AI-context and CLI tests passed.

Detect changes:

- Command: `go run ./cmd/anvien detect-changes --repo Anvien --scope all`
- `risk_level`: `low`
- `changed_files`: 7
- `affected_files`: 4
- `affected_count`: 0

## E15 - P2-E Document Skills Multi-Entry Package

Status: recorded

Source changes:

- Updated frontmatter `description` for `document-skills/docx`, `document-skills/pdf`, `document-skills/pptx`, and `document-skills/xlsx`.
- `docx` controls the generated package row, so it was set to the generic package trigger: `Use when the user asks to create, edit, or analyze documents.`
- `pdf`, `pptx`, and `xlsx` were shortened with type-aware triggers.

Impact:

- Command pattern: `go run ./cmd/anvien impact file internal/aicontext/skills/document-skills/<entry>/SKILL.md --repo Anvien --direction upstream`
- All 4 files reported `LOW`.
- All 4 files reported 0 affected files and 0 affected flows.

Generated-row verification after regeneration:

```text
document-skills -> Use when the user asks to create, edit, or analyze documents.
paths -> docx/SKILL.md, pdf/SKILL.md, pptx/SKILL.md, xlsx/SKILL.md
```

- Each nested source description matched the generated `.claude/skills/anvien/document-skills/<entry>/SKILL.md` mirror.

Validation:

```text
go run ./cmd/anvien analyze --force
go build ./cmd/... ./internal/...
go test ./internal/aicontext -count=1
go test ./internal/cli -run "TestAnalyzeCommand|AIContext|Aicontext" -count=1
```

Results:

- Analyze/regeneration passed.
- Full product build passed.
- Focused AI-context and CLI tests passed.

Detect changes:

- Command: `go run ./cmd/anvien detect-changes --repo Anvien --scope all`
- `risk_level`: `low`
- `changed_files`: 15
- `affected_files`: 4
- `affected_count`: 0

## E14 - P2-D Workflow And Role Group

Status: recorded

Source changes:

- Updated frontmatter `description` only for 4 packages:
  `repo-bootstrap`, `anvien-qa`, `frontend-development`, and `docs-seeker`.
- Skill bodies were not changed.

Impact:

- Command pattern: `go run ./cmd/anvien impact file internal/aicontext/skills/<skill>/SKILL.md --repo Anvien --direction upstream`
- All 4 files reported `LOW`.
- All 4 files reported 0 affected files and 0 affected flows.

Generated-row verification after regeneration:

```text
repo-bootstrap -> Use when the user asks to bootstrap a new repo or project.
anvien-qa -> Use when the user asks to run QA without fixing code.
frontend-development -> Use when the user asks to build or change frontend code.
docs-seeker -> Use when the user asks to find current technical documentation.
```

- Each source description matched the generated `.claude/skills/anvien/<skill>/SKILL.md` mirror.

Validation:

```text
go run ./cmd/anvien analyze --force
go build ./cmd/... ./internal/...
go test ./internal/aicontext -count=1
go test ./internal/cli -run "TestAnalyzeCommand|AIContext|Aicontext" -count=1
```

Results:

- Analyze/regeneration passed.
- Full product build passed.
- Focused AI-context and CLI tests passed.

Detect changes:

- Command: `go run ./cmd/anvien detect-changes --repo Anvien --scope all`
- `risk_level`: `low`
- `changed_files`: 4
- `affected_files`: 4
- `affected_count`: 0

## E10 - P1-D Google ADK Python

Status: recorded

Source change:

- `internal/aicontext/skills/google-adk-python/SKILL.md`
- Added standard `name` and `description` frontmatter because the file previously had no YAML frontmatter and generated the vague row `Google ADK Python Skill`.
- New description: `Use when the user asks to build Python agents with Google ADK.`

Impact:

- `anvien impact file internal/aicontext/skills/google-adk-python/SKILL.md --repo Anvien --direction upstream`
- Risk: `LOW`
- Affected files: 0
- Affected processes: 0

Generated-row verification after regeneration:

```text
AGENTS.md: google-adk-python -> Use when the user asks to build Python agents with Google ADK.
CLAUDE.md: google-adk-python -> Use when the user asks to build Python agents with Google ADK.
.claude/skills/anvien/google-adk-python/SKILL.md mirrors source frontmatter.
```

Validation:

```text
go run ./cmd/anvien analyze --force
go build ./cmd/... ./internal/...
go test ./internal/aicontext -count=1
go test ./internal/cli -run "TestAnalyzeCommand|AIContext|Aicontext" -count=1
```

Results:

- Analyze/regeneration passed.
- Full product build passed.
- Focused AI-context and CLI tests passed.

## E11 - P2-A Aesthetic

Status: recorded

Source change:

- `internal/aicontext/skills/aesthetic/SKILL.md`
- Replaced long capability-summary description with `Use when the user asks to improve UI aesthetics.`
- Body references to `ai-multimodal` remain for P2-B removal dependency review.

Impact:

- `anvien impact file internal/aicontext/skills/aesthetic/SKILL.md --repo Anvien --direction upstream`
- Risk: `LOW`
- Affected files: 0
- Affected processes: 0

Generated-row verification after regeneration:

```text
AGENTS.md: aesthetic -> Use when the user asks to improve UI aesthetics.
CLAUDE.md: aesthetic -> Use when the user asks to improve UI aesthetics.
.claude/skills/anvien/aesthetic/SKILL.md mirrors source description.
```

Validation:

```text
go run ./cmd/anvien analyze --force
go build ./cmd/... ./internal/...
go test ./internal/aicontext -count=1
go test ./internal/cli -run "TestAnalyzeCommand|AIContext|Aicontext" -count=1
```

Results:

- Analyze/regeneration passed.
- Full product build passed.
- Focused AI-context and CLI tests passed.

Detect changes:

- Command: `go run ./cmd/anvien detect-changes --repo Anvien --scope all`
- `risk_level`: `low`
- `changed_files`: 4
- `affected_files`: 3
- `affected_count`: 0

## E12 - P2-B Remove AI Multimodal

Status: recorded

Source change:

- Removed `internal/aicontext/skills/ai-multimodal/**`.
- Removal included hidden package-local files `.env.example` and `scripts/tests/.coverage`.
- Removed retained-package references from `internal/aicontext/skills/aesthetic/SKILL.md`.
- Removed retained-package reference from `internal/aicontext/skills/aesthetic/references/design-resources.md`.
- The direct `aesthetic` cleanup is part of P2-B because generated output would otherwise still contain the rejected package name.

Impact:

- `anvien impact file internal/aicontext/skills/ai-multimodal/SKILL.md --repo Anvien --direction upstream`: `LOW`, 0 affected files, 0 affected flows.
- Package impact loop covered all 13 non-hidden package files before deletion: every file was `LOW`, 0 affected flows; the three Python scripts each reported 1 affected file and no affected flows.
- Hidden package-local `.env.example` and `scripts/tests/.coverage` were removed as non-code package artifacts.
- `anvien impact file internal/aicontext/skills/aesthetic/SKILL.md --repo Anvien --direction upstream`: `LOW`, 0 affected files, 0 affected flows.

Regeneration and sync verification:

```text
go run ./cmd/anvien analyze --force
```

- First regeneration attempt failed because the empty source directory still existed without `SKILL.md`: `skill package "ai-multimodal" has no SKILL.md entry`.
- Removed the empty `internal/aicontext/skills/ai-multimodal` directory with a workspace path guard.
- Second regeneration passed: files scanned 1342, parsed code 697, graph nodes 82948, relationships 121591.

Generated-output verification after regeneration:

```text
rg -n "ai-multimodal" internal/aicontext/skills .claude/skills/anvien AGENTS.md CLAUDE.md
Test-Path .claude/skills/anvien/ai-multimodal -> False
Test-Path internal/aicontext/skills/ai-multimodal -> False
```

Results:

- No `ai-multimodal` matches remain in source skills, generated `.claude/skills/anvien`, `AGENTS.md`, or `CLAUDE.md`.
- `.claude/skills/anvien/.anvien-skill-manifest.json` no longer lists the package.
- Generated `Skill Selection Guide` rows decreased from 35 to 34.

Validation:

```text
go build ./cmd/... ./internal/...
go test ./internal/aicontext -count=1
go test ./internal/cli -run "TestAnalyzeCommand|AIContext|Aicontext" -count=1
```

Results:

- Full product build passed.
- Focused AI-context and CLI tests passed.

Detect changes:

- Command: `go run ./cmd/anvien detect-changes --repo Anvien --scope all`
- `risk_level`: `low`
- `changed_files`: 19
- `affected_files`: 5
- `affected_count`: 0

## E13 - P2-C Long Single-Entry Group

Status: recorded

Process correction:

- User directed grouping remaining description edits by skill clusters instead of one commit per skill.
- Updated the plan checklist from per-skill pending items to grouped items before continuing implementation.

Source changes:

- Updated frontmatter `description` only for 12 single-entry packages:
  `media-processing`, `shopify`, `backend-development`, `better-auth`, `repomix`, `web-frameworks`, `code-review`, `databases`, `ui-styling`, `mermaidjs-v11`, `mcp-management`, and `bunny`.
- Skill bodies were not changed.

Impact:

- Command pattern: `go run ./cmd/anvien impact file internal/aicontext/skills/<skill>/SKILL.md --repo Anvien --direction upstream`
- All 12 files reported `LOW`.
- All 12 files reported 0 affected files and 0 affected flows.

Generated-row verification after regeneration:

```text
media-processing -> Use when the user asks to process media files.
shopify -> Use when the user asks to build Shopify apps, themes, or extensions.
backend-development -> Use when the user asks to build or change backend code.
better-auth -> Use when the user asks to implement Better Auth.
repomix -> Use when the user asks to package a repository with Repomix.
web-frameworks -> Use when the user asks to build with Next.js, Turborepo, or web frameworks.
code-review -> Use when the user asks to review code or handle review feedback.
databases -> Use when the user asks to work with databases.
ui-styling -> Use when the user asks to style UI.
mermaidjs-v11 -> Use when the user asks to create Mermaid diagrams.
mcp-management -> Use when the user asks to manage MCP integrations.
bunny -> Use when the user asks to integrate Bunny.net.
```

- Each source description matched the generated `.claude/skills/anvien/<skill>/SKILL.md` mirror.

Validation:

```text
go run ./cmd/anvien analyze --force
go build ./cmd/... ./internal/...
go test ./internal/aicontext -count=1
go test ./internal/cli -run "TestAnalyzeCommand|AIContext|Aicontext" -count=1
```

Results:

- Analyze/regeneration passed.
- Full product build passed.
- Focused AI-context and CLI tests passed.
