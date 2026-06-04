# Plan

Title: AI Context Skill Description Trigger Shortening

Date: 2026-06-04

Status: Draft - awaiting user approval

Companion files:

- Evidence ledger: [2026-06-04-aicontext-skill-description-triggers-evidence.md](2026-06-04-aicontext-skill-description-triggers-evidence.md)
- Benchmark ledger: [2026-06-04-aicontext-skill-description-triggers-benchmark.md](2026-06-04-aicontext-skill-description-triggers-benchmark.md)

## Master Rules

1. Follow active repository instructions and generated `AGENTS.md`.
2. Use `anvien-planner` for this plan artifact.
3. Do not edit generated `AGENTS.md`, `CLAUDE.md`, or `.claude/skills/anvien/**` as source of truth.
4. Source of truth is `internal/aicontext/skills/**/SKILL.md`.
5. Treat these skills as Anvien-provided skills for any target repository using Anvien, not as skills only for modifying the Anvien product repository.
6. Run `anvien analyze --force` before graph-based Anvien commands.
7. Run impact analysis before changing generator code; for frontmatter-only skill Markdown edits, run file-level impact when useful and record the result.
8. HIGH or CRITICAL blast radius is a warning, not a ban.
9. Code/source first; tests are validation only and must not define the desired description text.
10. Run full build before tests when implementation changes are made.
11. Regenerate AI context with the current source command before validating generated output.
12. Record evidence and benchmarkable counts as work completes.
13. Run `anvien detect-changes --repo Anvien --scope all` before committing implementation work.
14. Commit each approved completed implementation slice after validation.

## Goal

Shorten generated `Skill Selection Guide` descriptions so they are repo-agnostic trigger rules, not capability summaries, and remove catalog skills that the user explicitly rejects as unnecessary. Each retained description should tell the agent when to load the skill for the user's current target repository, while detailed workflow, tool, provider, and validation rules remain inside the target `SKILL.md` body.

The target style is:

```text
Use when the user asks to <do the skill's job>.
```

Example approved direction:

```text
Use when the user asks to debug.
```

## Problem

Current skill descriptions are often long summaries of workflow internals, supported tools, provider lists, validation gates, or capability catalogs. That makes `AGENTS.md` and `CLAUDE.md` heavier than necessary and can obscure the trigger that matters.

The `anvien-debugging` row is the clearest example. Its current generated description summarizes systematic debugging, root-cause tracing, fix planning, defense-in-depth, graph evidence, impact checks, full-build gates, and completion verification. That belongs in the skill body, not in the guide row. The guide row should only say when to use the skill.

This is not a request to remove capability from skills. It is a request to move capability detail out of the always-loaded selection guide and keep it inside the lazily opened skill.

The catalog must also avoid implying that the skill is only for fixing or operating the Anvien product repository. `internal/aicontext/skills/**` is edited in this repository because this repository owns the generator and bundled catalog, but generated skills are installed into any target repo that uses Anvien. Descriptions must speak from the target-repo user's task: debug, run QA, build backend code, integrate payments, write a plan, and so on.

During plan review, the user explicitly rejected `ai-multimodal` as unnecessary and without enough value for the generated catalog. That package is no longer a description-shortening target; it is a removal target.

## Scope

In scope:

- Frontmatter `description` fields in `internal/aicontext/skills/**/SKILL.md`.
- Generated `Skill Selection Guide` rows in `AGENTS.md` and `CLAUDE.md`, validated only through regeneration.
- Generated `.claude/skills/anvien/**` output, validated only through regeneration.
- Removing the `ai-multimodal` source package from the generated catalog and validating that generated output removes it.
- Tests that validate AI-context generation behavior if source changes reveal a generator or parser issue.
- Evidence and benchmark ledgers for before/after description length, token estimate, generated row count, build/test validation, and detect-changes.

Out of scope:

- Changing skill bodies except where a body contains obsolete trigger wording that conflicts with the new description contract.
- Changing command selection guide content.
- Changing generated output manually.
- Removing skills from the catalog, except the approved removal of `ai-multimodal`.
- Renaming skills.
- Editing generator code unless source-only description changes cannot produce the desired output.

## Description Contract

1. A description is a trigger, not a capability summary.
2. Prefer one short sentence.
3. Use the phrase shape `Use when the user asks to ...`.
4. Do not list every framework, provider, operation, model limit, or validation phase in the description.
5. Include a technology name only when it disambiguates the trigger, such as `Better Auth`, `Bunny.net`, `Shopify`, `Three.js`, or `Mermaid`.
6. Put workflow detail, validation gates, and tool procedures in the skill body.
7. For multi-entry packages, keep the package-level trigger concise and verify which entry controls the generated package row before editing.
8. The generated guide must remain understandable without loading the skill body.
9. Descriptions must be target-repo generic. Do not write them as if the skill is only for changing the Anvien product repository.
10. Avoid phrases such as `in Anvien-indexed repositories` unless the phrase is needed to distinguish the trigger. Most rows should simply state the user task.

## Repository Boundary

This plan is implemented in the Anvien product repository only because this repository owns the bundled skill catalog and generator. Validation commands that use `--repo Anvien` refer to this implementation repository.

The generated skills themselves are repo-agnostic. When installed into another repository, the descriptions must route work for that target repository. Do not encode Anvien-product-only assumptions into skill descriptions.

## Target Descriptions

| Skill package | Target description |
|---|---|
| `aesthetic` | `Use when the user asks to improve UI aesthetics.` |
| `anvien-api-surface` | `Use when the user asks to inspect API or MCP surfaces.` |
| `anvien-debugging` | `Use when the user asks to debug.` |
| `anvien-planner` | `Use when the user asks to create, write, or review a docs/plans plan.` |
| `anvien-qa` | `Use when the user asks to run QA without fixing code.` |
| `anvien-refactoring` | `Use when the user asks to refactor code.` |
| `backend-development` | `Use when the user asks to build or change backend code.` |
| `better-auth` | `Use when the user asks to implement Better Auth.` |
| `bunny` | `Use when the user asks to integrate Bunny.net.` |
| `chrome-devtools` | `Use when the user asks to automate or inspect a browser.` |
| `code-review` | `Use when the user asks to review code or handle review feedback.` |
| `context-engineering` | `Use when the user asks to design or improve AI-agent context.` |
| `databases` | `Use when the user asks to work with databases.` |
| `devops` | `Use when the user asks to deploy or operate infrastructure.` |
| `docs-seeker` | `Use when the user asks to find current technical documentation.` |
| `document-skills` | `Use when the user asks to create, edit, or analyze documents.` |
| `frontend-design` | `Use when the user asks to design a frontend UI.` |
| `frontend-development` | `Use when the user asks to build or change frontend code.` |
| `google-adk-python` | `Use when the user asks to build Python agents with Google ADK.` |
| `mcp-builder` | `Use when the user asks to build an MCP server.` |
| `mcp-management` | `Use when the user asks to manage MCP integrations.` |
| `media-processing` | `Use when the user asks to process media files.` |
| `mermaidjs-v11` | `Use when the user asks to create Mermaid diagrams.` |
| `payment-integration` | `Use when the user asks to integrate payments.` |
| `problem-solving` | `Use when the user asks to solve a hard problem.` |
| `repo-bootstrap` | `Use when the user asks to bootstrap a new repo or project.` |
| `repomix` | `Use when the user asks to package a repository with Repomix.` |
| `sequential-thinking` | `Use when the user asks to reason step by step.` |
| `shopify` | `Use when the user asks to build Shopify apps, themes, or extensions.` |
| `skill-creator` | `Use when the user asks to create or update a skill.` |
| `threejs` | `Use when the user asks to build 3D web experiences with Three.js.` |
| `ui-styling` | `Use when the user asks to style UI.` |
| `web-frameworks` | `Use when the user asks to build with Next.js, Turborepo, or web frameworks.` |
| `web-testing` | `Use when the user asks to test web behavior.` |

## Removal Targets

| Skill package | Removal reason | Required result |
|---|---|---|
| `ai-multimodal` | User rejected it as unnecessary and not valuable enough for the catalog. | Remove from source catalog and generated `AGENTS.md`, `CLAUDE.md`, and `.claude/skills/anvien/**`. |

## Implementation Direction

Each skill is a separate implementation slice. Do not batch-edit multiple skill descriptions in one unreviewable patch unless the user explicitly approves batching.

For each skill slice:

1. Read the source `SKILL.md`.
2. Confirm which frontmatter description controls the generated package row.
3. Run file-level impact when the file is represented in Anvien.
4. Edit only the `description` field.
5. Regenerate generated AI context with the current source command.
6. Verify the exact generated row in both `AGENTS.md` and `CLAUDE.md`.
7. Verify `.claude/skills/anvien/**/SKILL.md` mirrors source.
8. Run the build/test validation required by the slice.
9. Record evidence and benchmark data.
10. Run `anvien detect-changes --repo Anvien --scope all`.
11. Commit the slice after approval and validation.

## Acceptance Criteria

1. Every targeted package row in `Skill Selection Guide` uses the approved trigger-only description.
2. `ai-multimodal` is absent from source catalog, generated skill output, `AGENTS.md`, and `CLAUDE.md`.
3. Generated `AGENTS.md` and `CLAUDE.md` contain no long capability-summary descriptions for targeted rows.
4. Generated `AGENTS.md` and `CLAUDE.md` still point only to the correct `.claude/skills/anvien/**/SKILL.md` paths in the `Use` column.
5. Generated `.claude/skills/anvien/**` mirrors `internal/aicontext/skills/**`.
6. Multi-entry packages keep correct routing to all entries.
7. Retained skill bodies still preserve detailed workflow instructions.
8. Source edits are validated through build, focused AI-context tests, regeneration checks, and detect-changes.
9. Benchmark ledger records before/after guide length, description-length reduction, and removal counts.

## Risk Notes

- Short descriptions can become too vague. Keep disambiguating names such as `Shopify`, `Better Auth`, `Bunny.net`, `Mermaid`, `Three.js`, `Repomix`, and `Google ADK`.
- Multi-entry packages can have one entry controlling the package row. Inspect before editing them.
- `ai-multimodal` removal must be exact. Do not remove unrelated media-processing, document, browser, UI, or Gemini references unless they belong to that package's files.
- The installed `anvien` on PATH may lag source code. For generated output validation after generator changes, use the current source command and record which command produced the output.
- There are currently uncommitted pre-plan edits to `internal/aicontext/skills/anvien-planner/SKILL.md` and `internal/aicontext/skills/anvien-debugging/SKILL.md` made before this plan artifact was created. Reconcile them before any implementation commit.

## Phase Checklist

- [x] [P0-A] Reconcile pre-plan source edits.
  - Goal: restore process integrity before implementation continues.
  - Work Steps: inspect the current diff for `anvien-planner` and `anvien-debugging`; decide whether to keep them as the first approved slices or revert them; do not commit them until the plan is approved; record the decision in evidence.
  - Implementation Gate: user approval confirms whether the pending edits are accepted as part of this plan.
  - Acceptance: worktree state is documented and no accidental edit is committed outside the approved plan.

- [x] [P0-B] Record baseline generated guide size and row inventory.
  - Goal: measure current token/character pressure before shortening descriptions.
  - Work Steps: count generated `Skill Selection Guide` rows; measure `AGENTS.md` total characters and words; measure description character totals; record current generated rows with long-description examples.
  - Implementation Gate: graph has been refreshed and current generated files are understood as validation output, not source.
  - Acceptance: benchmark ledger contains baseline guide size, row count, and top long descriptions.

- [x] [P1-A] Update `anvien-debugging`.
  - Goal: make debugging trigger concise.
  - Work Steps: set the description to `Use when the user asks to debug.`; regenerate; verify `AGENTS.md`, `CLAUDE.md`, and generated skill mirror.
  - Implementation Gate: P0-A is resolved and the source path is confirmed.
  - Acceptance: generated row exactly matches the target description.

- [x] [P1-B] Update `anvien-planner`.
  - Goal: make plan-writing trigger concise and artifact-specific.
  - Work Steps: set the description to `Use when the user asks to create, write, or review a docs/plans plan.`; regenerate; verify generated rows and mirror output.
  - Implementation Gate: P0-A is resolved and pending edits are accepted or reapplied cleanly.
  - Acceptance: asking to write a plan clearly routes to `anvien-planner`.

- [x] [P1-C] Update `context-engineering`.
  - Goal: replace invalid `>-` description with a meaningful trigger.
  - Work Steps: set the description to `Use when the user asks to design or improve AI-agent context.`; regenerate; verify generated rows and mirror output.
  - Implementation Gate: source file is inspected and no unrelated body edits are needed.
  - Acceptance: generated guide no longer contains `>-`.

- [x] [P1-D] Update `google-adk-python`.
  - Goal: replace vague description with a Google ADK trigger.
  - Work Steps: set the description to `Use when the user asks to build Python agents with Google ADK.`; regenerate; verify generated rows and mirror output.
  - Implementation Gate: source file confirms the skill is ADK-specific.
  - Acceptance: generated guide explains when to use the skill without a generic label.

- [x] [P2-A] Update `aesthetic`.
  - Goal: shorten visual-aesthetics routing.
  - Work Steps: set the description to `Use when the user asks to improve UI aesthetics.`; regenerate; verify generated rows and mirror output.
  - Implementation Gate: source body keeps detailed aesthetic workflow.
  - Acceptance: long design workflow summary is absent from the generated guide.

- [x] [P2-B] Remove `ai-multimodal`.
  - Goal: remove the rejected package from the generated skill catalog.
  - Work Steps: inspect `internal/aicontext/skills/ai-multimodal/**`; remove direct retained-package references to the rejected package; remove the `ai-multimodal` source package; regenerate; verify `AGENTS.md`, `CLAUDE.md`, `.claude/skills/anvien/**`, and the generated manifest contain no `ai-multimodal` entry.
  - Implementation Gate: user approval for this plan includes removal of `ai-multimodal`.
  - Acceptance: `ai-multimodal` is absent from source and generated output, while all retained packages remain present.

- [x] [P2-C] Update long single-entry package triggers as one group.
  - Goal: shorten the highest-token single-entry skill descriptions without changing skill bodies.
  - Work Steps: run file impact for each touched `SKILL.md`; set descriptions to these exact values: `media-processing` -> `Use when the user asks to process media files.`, `shopify` -> `Use when the user asks to build Shopify apps, themes, or extensions.`, `backend-development` -> `Use when the user asks to build or change backend code.`, `better-auth` -> `Use when the user asks to implement Better Auth.`, `repomix` -> `Use when the user asks to package a repository with Repomix.`, `web-frameworks` -> `Use when the user asks to build with Next.js, Turborepo, or web frameworks.`, `code-review` -> `Use when the user asks to review code or handle review feedback.`, `databases` -> `Use when the user asks to work with databases.`, `ui-styling` -> `Use when the user asks to style UI.`, `mermaidjs-v11` -> `Use when the user asks to create Mermaid diagrams.`, `mcp-management` -> `Use when the user asks to manage MCP integrations.`, `bunny` -> `Use when the user asks to integrate Bunny.net.`; regenerate once; verify generated rows and mirror output for all skills in this group.
  - Implementation Gate: each source body remains intact and only frontmatter description is edited.
  - Acceptance: generated guide no longer contains the long capability catalog text for any skill in this group.

- [x] [P2-D] Update workflow and role-specific single-entry triggers as one group.
  - Goal: shorten routing for repo setup, QA, frontend development, and documentation search while keeping each body-specific workflow intact.
  - Work Steps: run file impact for each touched `SKILL.md`; set descriptions to these exact values: `repo-bootstrap` -> `Use when the user asks to bootstrap a new repo or project.`, `anvien-qa` -> `Use when the user asks to run QA without fixing code.`, `frontend-development` -> `Use when the user asks to build or change frontend code.`, `docs-seeker` -> `Use when the user asks to find current technical documentation.`; regenerate once; verify generated rows and mirror output for all skills in this group.
  - Implementation Gate: QA retains the no-code-fix rule, repo-bootstrap remains repo-type-aware, and docs-seeker keeps current-doc search workflow in the body.
  - Acceptance: generated guide routes these workflows with trigger-only descriptions.

- [x] [P2-E] Update `document-skills` multi-entry package trigger.
  - Goal: shorten document-work routing without dropping nested DOCX/PDF/PPTX/XLSX entries.
  - Work Steps: inspect which nested `SKILL.md` descriptions control the package row; set the package-row controlling description to `Use when the user asks to create, edit, or analyze documents.`; keep nested entry descriptions concise and type-aware if they are edited; regenerate once; verify all document skill paths remain listed.
  - Implementation Gate: multi-entry package behavior is confirmed before editing.
  - Acceptance: generated guide routes document work without a long DOCX/PDF/PPTX/XLSX capability summary.

- [x] [P3-A] Normalize concise single-entry triggers as one group.
  - Goal: make the remaining single-entry package descriptions use the same trigger-only style.
  - Work Steps: run file impact for each touched `SKILL.md`; set descriptions to these exact values: `anvien-api-surface` -> `Use when the user asks to inspect API or MCP surfaces.`, `anvien-refactoring` -> `Use when the user asks to refactor code.`, `chrome-devtools` -> `Use when the user asks to automate or inspect a browser.`, `devops` -> `Use when the user asks to deploy or operate infrastructure.`, `frontend-design` -> `Use when the user asks to design a frontend UI.`, `mcp-builder` -> `Use when the user asks to build an MCP server.`, `payment-integration` -> `Use when the user asks to integrate payments.`, `sequential-thinking` -> `Use when the user asks to reason step by step.`, `skill-creator` -> `Use when the user asks to create or update a skill.`, `threejs` -> `Use when the user asks to build 3D web experiences with Three.js.`, `web-testing` -> `Use when the user asks to test web behavior.`; regenerate once; verify generated rows and mirror output for all skills in this group.
  - Implementation Gate: each source body keeps its detailed workflow and payment routing remains explicit.
  - Acceptance: generated guide uses concise trigger-only descriptions for all skills in this group.

- [ ] [P3-B] Update `problem-solving` multi-entry package trigger.
  - Goal: normalize problem-solving routing without dropping nested problem-solving entries.
  - Work Steps: inspect the multi-entry package primary description; set the generated package trigger to `Use when the user asks to solve a hard problem.`; regenerate once; verify all problem-solving entry paths remain listed.
  - Implementation Gate: multi-entry package behavior is confirmed before editing.
  - Acceptance: generated guide is concise and all problem-solving entries remain reachable.

- [ ] [P4-A] Final regeneration, validation, and inventory check.
  - Goal: prove all generated outputs reflect source and no old long descriptions remain.
  - Work Steps: regenerate with the approved current source command; check `AGENTS.md`, `CLAUDE.md`, and `.claude/skills/anvien/**`; run full build; run focused AI-context tests; measure final guide length and description totals.
  - Implementation Gate: all per-skill slices are complete and evidence is updated.
  - Acceptance: final generated guide matches the target matrix and validation passes.

- [ ] [P4-B] Detect changes and commit closure.
  - Goal: close the plan with traceable change evidence.
  - Work Steps: run `anvien detect-changes --repo Anvien --scope all`; record affected files and risk; commit approved completed slices; record commit hashes in evidence.
  - Implementation Gate: validation has passed and user approval scope is clear.
  - Acceptance: repository has committed, validated source changes and updated evidence/benchmark ledgers.
