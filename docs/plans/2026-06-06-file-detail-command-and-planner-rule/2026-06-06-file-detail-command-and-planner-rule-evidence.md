# Evidence Ledger

Title: File Detail Command And Planner Rule
Date: 2026-06-06
Status: Initialized
Companion plan: `docs/plans/2026-06-06-file-detail-command-and-planner-rule/2026-06-06-file-detail-command-and-planner-rule-plan.md`
Companion benchmark: `docs/plans/2026-06-06-file-detail-command-and-planner-rule/2026-06-06-file-detail-command-and-planner-rule-benchmark.md`

## Evidence Rules

- Record evidence as each checklist item is completed.
- Keep generated output evidence separate from source-of-truth source changes.
- Historical reports/plans that mention old command names are evidence history, not active guidance.
- Record command failures and the handling decision when they affect the plan.

## E0 - Plan Authoring Evidence

- Planner skill read:
  `.agents/skills/planner/SKILL.md`
  Result: confirmed the required three-file standard plan set under `docs/plans/YYYY-MM-DD-<slug>/`, complete mini-plan checklist items, evidence ledger, and benchmark ledger.

- Repo rules read:
  `AGENTS.md`
  Result: confirmed generated content must not be edited as permanent source; `anvien analyze --force` is required before graph work; impact must run before editing functions/classes/methods/API/shared contracts; detect-changes must run before committing implementation work.

- Nearby plan read:
  `docs/plans/2026-06-06-aicontext-rule1-help-wording/2026-06-06-aicontext-rule1-help-wording-plan.md`
  Result: confirmed prior AI-context generator plan style and source/generator boundary.

- Nearby evidence read:
  `docs/plans/2026-06-06-aicontext-rule1-help-wording/2026-06-06-aicontext-rule1-help-wording-evidence.md`
  Result: confirmed `renderMasterRulesBlock` and generated `AGENTS.md` / `CLAUDE.md` were handled through generator source and validated with full build plus AI-context tests.

- Missing nearby benchmark:
  `docs/plans/2026-06-06-aicontext-rule1-help-wording/2026-06-06-aicontext-rule1-help-wording-benchmark.md`
  Result: file was absent. This plan follows the current planner skill and creates the full three-file set.

## E1 - Initial Graph And Source Evidence

- `anvien analyze --force`
  Result: completed before writing this plan. Graph reported `files.scanned=1406`, `parsed_code=683`, `failed=0`, `nodes=83978`, `relationships=122367`, `dependencyEdges=16549`, `unresolved=429`.

- `anvien query "file-context command and file detail command surface" --repo Anvien`
  Result: command-surface discovery identified `internal/cli/tool_command.go`, `internal/cli/command.go`, and file detail Web surfaces. Follow-up text search identified the concrete command owner as `internal/cli/file_context_command.go`.

- `anvien query "generated AGENTS CLAUDE master iron rules aicontext" --repo Anvien`
  Result: docs/setup/AI-context discovery identified `internal/aicontext/aicontext.go` and `internal/aicontext/aicontext_test.go` as relevant generator/test surfaces.

- `rg -n --hidden --glob '!node_modules/**' --glob '!.git/**' 'file-context|Command Selection Guide|Master iron rules|write/create a plan|planner skill' cmd internal docs README.md RUNBOOK.md AGENTS.md CLAUDE.md package.json`
  Result: active old-name surfaces include:
  `AGENTS.md`, `CLAUDE.md`, `README.md`, `RUNBOOK.md`, `internal/aicontext/aicontext.go`, `internal/aicontext/aicontext_test.go`, `internal/cli/file_context_command.go`, `internal/cli/file_context_command_test.go`, `internal/cli/command_test.go`, `internal/httpapi/server.go`, `internal/httpapi/file_context_test.go`, `internal/contracts/web_ui.go`, `internal/contracts/web_ui_test.go`, and active skill guidance files.

## E2 - User Decisions Captured

- Approved Command Selection Guide label:

```text
Read detailed information about a file and its relationships with other files.
```

- Approved command rename:
  `file-context` becomes `file-detail`.

- Alias decision:
  no alias, no deprecated command, no compatibility fallback.

- Approved generated Master iron rule:

```text
When the user asks to "write/create a plan", the AI agent must immediately use the planner skill and create a real docs/plans plan.
```

- Approved generated `# AGENTS Rules` rule 4:

```text
4. **Write plan (use planner skill) before coding.**
```

## E3 - Pending Implementation Evidence

- P0-A through P7-A are not implemented yet.
- No source code has been changed as part of this plan authoring step.
- Future implementation evidence must be appended under this section as each checklist item is completed.

## E4 - Pending Validation Evidence

- Required validation pending:
  `anvien file-detail --help`
- Required validation pending:
  `anvien file-detail <path> --repo Anvien --json`
- Required validation pending:
  `anvien file-context --help` fails as unknown command.
- Required validation pending:
  generated `AGENTS.md` and `CLAUDE.md` contain the approved wording and planner rules.
- Required validation pending:
  targeted tests and `npm run full-build`.
- Required validation pending:
  `anvien detect-changes --repo Anvien --scope all`.

## E5 - Commit Evidence

- Pending.
