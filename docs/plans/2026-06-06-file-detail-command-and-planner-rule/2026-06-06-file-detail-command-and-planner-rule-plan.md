# Plan

Title: File Detail Command And Planner Rule
Date: 2026-06-06
Status: Ready for Implementation
Companion evidence: `docs/plans/2026-06-06-file-detail-command-and-planner-rule/2026-06-06-file-detail-command-and-planner-rule-evidence.md`
Companion benchmark: `docs/plans/2026-06-06-file-detail-command-and-planner-rule/2026-06-06-file-detail-command-and-planner-rule-benchmark.md`

## Goal

Implement the user-approved command and generated AI-context rule changes without creating legacy aliases or manual generated-file drift.

## User-Approved Requirements

1. Replace the generated Command Selection Guide row label with:

```text
Read detailed information about a file and its relationships with other files.
```

2. Rename the active file detail command surface from `file-context` to `file-detail`.
3. Do not keep `file-context` as an alias, compatibility path, or deprecated command.
4. Add this generated Master iron rule:

```text
When the user asks to "write/create a plan", the AI agent must immediately use the planner skill and create a real docs/plans plan.
```

5. Change generated `# AGENTS Rules` rule 4 to:

```text
4. **Write plan (use planner skill) before coding.**
```

## Scope

- CLI command surface for the current `file-context` command.
- HTTP/API, Web contract, and Web client surfaces only where they expose an active `file-context` route or command contract.
- Generated AI-context source for `AGENTS.md` and `CLAUDE.md`.
- Active tests and active user-facing docs that reference the renamed command.
- `CHANGELOG.md` entry for the command rename and generated-rule change.

## Non-Goals

- Do not rewrite historical reports, old plan evidence, or old benchmark ledgers that mention `file-context` as past evidence.
- Do not introduce an alias, redirect, compatibility registration, or fallback path for `file-context`.
- Do not manually edit generated `AGENTS.md` or `CLAUDE.md` as the source of truth.
- Do not rename unrelated internal package concepts unless they expose the old command/API surface or keep active old naming visible.
- Do not bump version unless explicitly requested.

## Invariants

- There is exactly one active file detail command name: `file-detail`.
- `file-context` is not discoverable through CLI help, generated command guidance, active docs, active skill guidance, API contract metadata, or Web client code.
- Generated `AGENTS.md` and `CLAUDE.md` are produced from source generators.
- Any CRITICAL/HIGH Anvien blast radius is treated as a scope warning, not a prohibition.
- Checklist items are updated immediately as implementation phases complete.

## Checklist

- [x] P0-A: Confirm source owners and baseline.
  Goal: establish the exact source-of-truth files before editing.
  Work Steps: run `anvien analyze --force`; query the command and AI-context generator surfaces; inspect `internal/cli/file_context_command.go`, `internal/cli/command.go`, `internal/aicontext/aicontext.go`, active API route/contract/client files, active docs, and active skill files that mention `file-context`.
  Implementation Gate: do not edit until evidence records the current owners, baseline `file-context` active-reference count, and generated-file source-of-truth rule.
  Acceptance: evidence lists each active surface to change and separates active references from historical reports/plans.

- [x] P1-A: Run impact before touching command, API, and generator symbols.
  Goal: make the blast radius explicit before code changes.
  Work Steps: run Anvien impact for the CLI command constructor, root command registration, file detail API handler or route registration, Web contract generator when touched, and `renderMasterRulesBlock` / command-guide generation symbols.
  Implementation Gate: do not edit any function, class, method, exported symbol, API handler, graph builder, resolver, analyzer, or shared contract before its impact evidence is recorded.
  Acceptance: evidence records risk level, affected files, affected flows/tests, and the scoped reason for proceeding.

- [x] P2-A: Rename the CLI command surface from `file-context` to `file-detail`.
  Goal: make `anvien file-detail` the only CLI command for structured file details.
  Work Steps: rename the command use string/help/examples; rename command constructor/test names where they carry active command naming; update root help tests; remove all command registration for `file-context`; update CLI tests from `file-context` to `file-detail`.
  Implementation Gate: no alias or compatibility command may be registered.
  Acceptance: `anvien file-detail --help` works, `anvien file-detail <path> --repo Anvien --json` returns structured file data, and `anvien file-context --help` fails as an unknown command.

- [x] P3-A: Rename active API/Web/contract surfaces if they expose `file-context`.
  Goal: remove active non-CLI `file-context` surfaces that would keep the old name alive.
  Work Steps: inspect and rename `/api/file-context` route metadata, server registration, API tests, generated Web contract metadata, Web backend client references, and active Web tests to `/api/file-detail` where applicable.
  Implementation Gate: only rename active route/contract/client surfaces; do not rewrite historical plan/report evidence.
  Acceptance: `/api/file-detail` returns the same file detail contract, `/api/file-context` is not registered, contract tests pass, and Web/API callers use the new route.

- [x] P4-A: Update AI-context generator wording and generated planner rules.
  Goal: generate the approved guidance into both `AGENTS.md` and `CLAUDE.md`.
  Work Steps: update `internal/aicontext/aicontext.go` so the Command Selection Guide row uses the approved wording and `file-detail`; add the approved Master iron rule; update `# AGENTS Rules` rule 4 to `**Write plan (use planner skill) before coding.**`; update generator tests after source behavior is changed.
  Implementation Gate: do not manually edit generated managed sections as permanent source.
  Acceptance: regenerated `AGENTS.md` and `CLAUDE.md` contain the approved wording, the new Master iron rule, rule 4 wording, and no active `file-context` command guidance.

- [x] P5-A: Update active docs and skill guidance.
  Goal: keep active user/agent guidance aligned with the renamed command.
  Work Steps: update `README.md`, `RUNBOOK.md`, active skill files, and active templates that instruct agents to run `file-context`; update `CHANGELOG.md` with the rename and planner-rule change; leave historical reports and old plan evidence untouched.
  Implementation Gate: every active doc edit must be traceable to current command/rule guidance, not historical cleanup.
  Acceptance: active docs/skills reference `file-detail`; `rg "file-context"` returns only historical reports/plans or intentionally named source files that cannot be renamed without breaking package semantics, with each residual justified in evidence.

- [x] P6-A: Regenerate and validate.
  Goal: prove runtime behavior and generated output match the plan.
  Work Steps: run formatting for touched code; run targeted Go tests for CLI, HTTP API, contracts, AI context, and any touched Web client tests; run command smokes for `file-detail`; run API smoke for `/api/file-detail` if route is renamed; run `npm run full-build`.
  Implementation Gate: full build must pass before final testing evidence is accepted.
  Acceptance: validation evidence shows `file-detail` works, `file-context` is gone from active command/API guidance, generated files are updated from source, and full build passes.

- [ ] P7-A: Run final Anvien change detection and commit.
  Goal: close the implementation slice with graph-aware change evidence.
  Work Steps: run `anvien analyze --force`; run `anvien detect-changes --repo Anvien --scope all`; review dirty files; stage only scope files; commit with evidence in the commit body.
  Implementation Gate: do not commit implementation work until detect-changes has completed.
  Acceptance: detect-changes evidence is recorded, commit hash is recorded in evidence, and worktree status is clean or only explicitly out-of-scope files remain.

## Acceptance Criteria

- `anvien file-detail --help` documents the renamed command.
- `anvien file-detail <path> --repo Anvien --json` returns structured file-level data and relationships.
- `anvien file-context --help` fails because no alias or legacy command exists.
- Generated `AGENTS.md` and `CLAUDE.md` contain the approved file-detail wording and planner rules.
- Active docs and active skill guidance no longer instruct agents to use `file-context`.
- Full build passes.
- `anvien detect-changes --repo Anvien --scope all` is recorded before commit.

## Risks

- The old name may exist across CLI, API, Web, contract, docs, and skill guidance. The implementation must treat this as a surface rename, not a local CLI text change.
- If `/api/file-context` is renamed, Web and contract tests may need coordinated updates.
- Historical reports/plans contain many legitimate old mentions; active-reference checks must not rewrite past evidence.
