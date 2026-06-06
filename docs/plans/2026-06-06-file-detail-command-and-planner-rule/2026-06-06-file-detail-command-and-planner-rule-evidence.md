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

- P0-A completed.
- `git status --short`
  Result: worktree was clean before implementation.
- `anvien analyze --force`
  Result: graph refreshed successfully before implementation. Graph reported `files.scanned=1409`, `parsed_code=683`, `failed=0`, `nodes=84004`, `relationships=122393`, `dependencyEdges=16549`, `unresolved=429`.
- `rg -n -g '!docs/plans/**' -g '!reports/**' "file-context" cmd internal README.md RUNBOOK.md AGENTS.md CLAUDE.md anvien-web contracts CHANGELOG.md`
  Result: active old-surface references remain in CLI, HTTP/API contracts, Web client/tests, generated AI-context output, README/RUNBOOK, generated root files, and active skill templates. Historical plan/report references remain out of scope.
- P1-A completed.
- Impact commands used `anvien impact ... --repo Anvien --direction upstream --include-tests --json` for the touched command, route, contract, Web client, and AI-context generator surfaces.
- Impact summary:
  - `newFileContextCommand`: `CRITICAL`, `impacted=62`, key files `cmd/anvien/main.go`, `internal/cli/command.go`, `internal/cli/command_test.go`, `internal/cli/file_context_command_test.go`.
  - `NewRootCommand`: `CRITICAL`, `impacted=61`, key files `cmd/anvien/main.go`, `internal/cli/command_test.go`.
  - `handleFileContext`: `LOW`, `impacted=0`.
  - `NewHandler`: `CRITICAL`, `impacted=63`, key files `internal/httpapi/file_context_test.go`, `internal/httpapi/handlers_test.go`, `internal/httpapi/listen.go`.
  - `WebUIContract`: `CRITICAL`, `impacted=6`, key files `cmd/generate-web-contracts/main.go`, `internal/contracts/web_ui.go`, `internal/contracts/web_ui_test.go`.
  - `renderMasterRulesBlock`: `HIGH`, `impacted=8`, key files `internal/aicontext/aicontext.go`, `internal/aicontext/aicontext_test.go`, `internal/cli/analyze_postrun.go`.
  - `GenerateAIContextFiles`: `CRITICAL`, `impacted=8`, key files `internal/aicontext/aicontext.go`, `internal/aicontext/aicontext_test.go`, `internal/cli/analyze_postrun.go`.
  - `fetchFileContext`: `UNKNOWN`, `impacted=0`; file-level web client impact used instead.
  - `anvien-web/src/services/backend-client.ts`: `CRITICAL`, `impacted=122`, key files include Web app components and tests that import backend-client helpers.
- Proceeding reason: CRITICAL/HIGH are blast-radius warnings. The scoped edits are a coordinated rename of one command/route surface plus generated guidance wording and matching tests/docs; no broad behavior or data contract shape change is intended.
- P2-A completed.
- CLI source changes:
  - Moved `internal/cli/file_context_command.go` to `internal/cli/file_detail_command.go`.
  - Moved `internal/cli/file_context_command_test.go` to `internal/cli/file_detail_command_test.go`.
  - Renamed `newFileContextCommand` to `newFileDetailCommand`.
  - Changed Cobra use/help from `file-context <path>` to `file-detail <path>`.
  - Updated root command registration to add only `newFileDetailCommand`; no alias command was registered.
  - Updated CLI help and command tests to expect `file-detail`.
  - Added `TestFileContextCommandIsNotRegistered` to lock the no-alias decision.
- `rg -n "file-context|newFileContextCommand|TestFileContext" internal/cli`
  Result: no matches.
- `go test ./internal/cli -run "TestFileDetailCommand|TestFileContextCommandIsNotRegistered|TestDirectToolHelpShowsCompatibilityFlags|TestHelpCommandPrintsStubHelp"`
  Result: passed.
- P3-A completed.
- API/Web/contract source changes:
  - HTTP server route changed from `/api/file-context` to `/api/file-detail`.
  - HTTP tests now call `/api/file-detail`.
  - Added `TestFileContextEndpointIsNotRegistered` to verify the old route returns `404`.
  - Web UI contract source route metadata changed to `/api/file-detail`.
  - Generated `contracts/web-ui/anvien-web-contract.schema.json` and `anvien-web/src/generated/anvien-contracts.ts` by running `go run ./cmd/generate-web-contracts`.
  - Web backend client and active Web unit/e2e tests now use `/api/file-detail`.
- `go test ./internal/httpapi ./internal/contracts -run "TestFileDetail|TestFileContextEndpointIsNotRegistered|TestWebUIContract"`
  Result: passed.
- `npm --prefix anvien-web run test -- server-connection.test.ts`
  Result: passed, `1` file / `21` tests.
- `rg -n "file-context|/api/file-detail|/api/file-context" internal/httpapi internal/contracts contracts/web-ui anvien-web/src anvien-web/test anvien-web/e2e`
  Result: active route/client/contract references use `/api/file-detail`; only remaining `/api/file-context` reference in these roots is the explicit negative HTTP test that proves the legacy route is not registered.
- P4-A completed.
  - `internal/aicontext/aicontext.go` now emits `file-detail` in the graph-work freshness list.
  - Command Selection Guide now uses `Read detailed information about a file and its relationships with other files.` and `anvien file-detail <path> --repo <repo> --json`.
  - Master iron rules now include the planner-skill docs/plans rule.
  - AGENTS rule 4 now emits `**Write plan (use planner skill) before coding.**`.
  - `go test ./internal/aicontext -run TestGenerateAIContextFilesCreatesManagedContextAndSkillPackages`
    Result: passed.
  - `go run ./cmd/anvien analyze . --force`
    Result: regenerated root AI-context output from source and refreshed the graph: `files.scanned=1409`, `nodes=84018`, `relationships=122400`, `dependencyEdges=16551`.
  - `rg -n "Master iron rules|write/create a plan|Write plan \\(use planner skill\\)|Read detailed information about a file and its relationships|file-detail|file-context" AGENTS.md CLAUDE.md`
    Result: both generated files contain the approved planner rule, rule 4 wording, `file-detail`, and the approved file-detail row wording; no generated `file-context` matches remain in those files.
- P5-A completed.
- Active docs/skills changed:
  - `RUNBOOK.md` file projection CLI/API examples now use `file-detail` and `/api/file-detail`.
  - `internal/aicontext/skills/debugging/SKILL.md`, `qa/SKILL.md`, `ui-be-binding-skill/SKILL.md`, and the UI-BE binding evidence template now use `file-detail`.
  - `internal/aicontext/skills/coder/SKILL.md` rule 1 now says `Write plan (use planner skill) before coding`.
  - `CHANGELOG.md` documents the command/API rename and generated planner rule change under `1.2.6`.
- `rg -n -g '!docs/plans/**' -g '!reports/**' "file-context|/api/file-context|anvien file-context|Write plan before coding" README.md RUNBOOK.md CHANGELOG.md internal/aicontext/skills internal/aicontext/aicontext.go internal/aicontext/aicontext_test.go internal/cli internal/httpapi internal/contracts contracts anvien-web AGENTS.md CLAUDE.md`
  Result before regeneration: remaining old mentions are generated `AGENTS.md` / `CLAUDE.md`, changelog change note, and explicit negative tests proving no CLI/API legacy surface.
- `go test ./internal/aicontext`
  Result: passed.

## E4 - Validation Evidence

- P6-A completed.
- `npm run full-build`
  Result: passed. Full build installed package version `1.2.6`, built the Web UI, built the Windows launcher/runtime, and ran `anvien analyze . --force` successfully.
- `anvien file-detail --help`
  Result: passed; usage contains `anvien file-detail <path>`.
- `anvien file-detail internal/aicontext/aicontext.go --repo Anvien --json`
  Result: passed; JSON smoke returned repo `Anvien`, path `internal/aicontext/aicontext.go`, `symbolCount=56`, combined relationship count `54`, response size `72831` bytes.
- `anvien file-context --help`
  Result: failed as expected with `unknown command "file-context"` and exit code `1`.
- API smoke with temporary `anvien\bin\anvien.exe serve --host 127.0.0.1 --port 4861`
  Result: `/api/file-detail` returned `200` for `internal/aicontext/aicontext.go`; `/api/file-context` returned `404`; temporary server process was stopped.
- `go test ./internal/cli ./internal/httpapi ./internal/contracts ./internal/aicontext -run "TestFileDetailCommand|TestFileContextCommandIsNotRegistered|TestDirectToolHelpShowsCompatibilityFlags|TestHelpCommandPrintsStubHelp|TestFileDetail|TestFileContextEndpointIsNotRegistered|TestWebUIContract|TestGenerateAIContextFilesCreatesManagedContextAndSkillPackages|TestSkillPackage"`
  Result: passed.
- `npm --prefix anvien-web run test -- server-connection.test.ts`
  Result: passed, `1` file / `21` tests.
- `npm --prefix anvien-web run test:e2e -- file-map-test-unresolved.spec.ts`
  Result: passed, `1` Chromium Playwright test.
- Root help check:
  Result: `anvien --help` contains `file-detail`, does not contain `file-context`, and still contains `skill`.
- Final generated context check:
  Result: `AGENTS.md` and `CLAUDE.md` contain the approved planner rule, rule 4 wording, `file-detail`, and the approved file-detail row wording.
- Final active old-name inventory:
  `rg -n -g '!docs/plans/**' -g '!reports/**' "file-context" README.md RUNBOOK.md CHANGELOG.md internal/aicontext/skills internal/aicontext/aicontext.go internal/aicontext/aicontext_test.go internal/cli internal/httpapi internal/contracts contracts anvien-web AGENTS.md CLAUDE.md .agents .claude`
  Result: `8` lines across `3` files. All are justified: `CHANGELOG.md` describes the rename, `internal/cli/file_detail_command_test.go` verifies the old CLI command is not registered, and `internal/httpapi/file_context_test.go` verifies the old API route is not registered.

## E5 - Commit Evidence

- P7 detect-changes completed before implementation commit.
- `anvien analyze --force`
  Result: graph refreshed after final evidence/doc edits: `files.scanned=1409`, `nodes=84018`, `relationships=122400`, `dependencyEdges=16551`.
- `anvien detect-changes --repo Anvien --scope all`
  Result: completed with `risk_level=critical`, `changed_files=25`, `affected_files=18`, `changed_count=49`, `affected_count=16`. CRITICAL is expected because the slice intentionally touches CLI root registration, HTTP route registration, Web contract metadata, AI-context generator output, and the Web backend client.
- High-risk changed files reported by detect-changes include `anvien-web/src/services/backend-client.ts`, `internal/aicontext/aicontext.go`, `internal/cli/command.go`, `internal/contracts/web_ui.go`, and `internal/httpapi/server.go`.
- Implementation commit: `793d49b25d18a572ec3a5651682c821d70c0661f` (`feat: rename file-context to file-detail`).
- P7-A completed with implementation work committed.
