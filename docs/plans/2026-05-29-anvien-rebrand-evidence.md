# Anvien Hard Rebrand Evidence Ledger

Date: 2026-05-29

Status: Draft

Companion files:

- Plan: [2026-05-29-anvien-rebrand-plan.md](2026-05-29-anvien-rebrand-plan.md)
- Benchmark ledger: [2026-05-29-anvien-rebrand-benchmark.md](2026-05-29-anvien-rebrand-benchmark.md)

## Evidence Rules

Record commands, changed files, naming decisions, old-name inventories, AVmatrix impact output, build/test/e2e output, package smoke results, and concise observations needed to audit the hard rebrand later.

No inferred behavior is accepted as evidence. Every behavior claim needs source inspection, AVmatrix output, command output, test output, or an explicit recorded decision.

## E0 - Plan Creation Evidence

Date: 2026-05-29

Status: recorded

Created file set:

- `docs/plans/2026-05-29-anvien-rebrand-plan.md`
- `docs/plans/2026-05-29-anvien-rebrand-evidence.md`
- `docs/plans/2026-05-29-anvien-rebrand-benchmark.md`

Original plan issue found:

- The first draft treated this as a generic compatibility-preserving rename.
- User clarified that MCP must also be renamed to `anvien` and that no legacy alias should remain.
- The plan was rewritten as a hard rename with no active legacy runtime/config/MCP/package surface.

## E1 - AVmatrix Graph Refresh

Date: 2026-05-29

Status: recorded

Command:

```powershell
avmatrix analyze --force
```

Output summary:

```text
analyzed E:\AVmatrix-GO
files: scanned=800 parsed=583 unsupported=217 failed=0
graph: nodes=91223 relationships=124702 path=E:\AVmatrix-GO\.avmatrix\graph.json
```

Indexed repo discovery:

```text
AVmatrix
Path: E:\AVmatrix-GO
Commit: 3e37a3e
```

## E2 - AVmatrix Query Evidence

Date: 2026-05-29

Status: recorded

MCP query:

```powershell
avmatrix query "MCP server command setup configuration mcp" --repo AVmatrix
```

Key results:

- `internal/mcp/resources.go`: `Server.setupResource`
- `internal/mcp/prompts.go`: `generateMapPrompt`, `promptDefinitions`
- `internal/cli/setup_command.go`: `setupInstallEditorSkills`, `setupInstallSkillsTo`
- `internal/mcp/tools.go`: query/context/process matching functions
- process evidence included `Handle -> GenerateMapSteps` and MCP resource/tool flows

CLI/package query:

```powershell
avmatrix query "CLI command name binary package metadata launcher npm" --repo AVmatrix
```

Key results:

- `internal/cli/setup_command.go`: `newSetupCommand`, `runSetup`, setup installer functions
- `internal/cli/query_health_command.go`: command-surface example
- `internal/cli/package_command.go`: package command owner

Generated context query:

```powershell
avmatrix query "generated AGENTS CLAUDE skills AI context avmatrix" --repo AVmatrix
```

Key results:

- `internal/aicontext/aicontext.go`: `installBaseSkills`, `baseSkillContent`, `GenerateAIContextFiles`, `renderAVmatrixBlock`, `removeGeneratedSkills`
- `internal/cli/analyze_postrun.go`: `generateAnalyzeAIContext`
- process evidence included `NewAnalyzeCommand -> RenderAVmatrixBlock`

Web/launcher query:

```powershell
avmatrix query "web UI brand title onboarding start AVmatrix launcher" --repo AVmatrix
```

Key results:

- `avmatrix-web/src/components/LauncherStartScreen.tsx`: `LauncherStartScreen`
- `avmatrix-web/src/components/AnalyzeOnboarding.tsx`: `AnalyzeOnboarding`
- `avmatrix-launcher/src/main.go`: launcher lifecycle/process functions
- `avmatrix-launcher/server-wrapper/main.go`: server wrapper entrypoint

Storage query:

```powershell
avmatrix query "repo storage directory .avmatrix AVMATRIX_HOME registry paths" --repo AVmatrix
```

Key results:

- `internal/repo/paths.go`: `Paths`
- `internal/group/storage.go`: storage/registry readers
- `internal/cli/group_command.go`: group path normalization and config writers

MCP resource scheme query:

```powershell
avmatrix query "MCP resource scheme avmatrix:// resources prompts setupResource" --repo AVmatrix
```

Key results:

- `internal/mcp/resources.go`: `Server.setupResource`, `resourceDefinitions`, resource process helpers
- `internal/mcp/prompts.go`: `generateMapPrompt`, `promptDefinitions`
- `internal/cli/setup_command.go`: setup write flow

## E3 - AVmatrix Context Evidence

Date: 2026-05-29

Status: recorded

`setupWriteMCPJSON`:

```powershell
avmatrix context "setupWriteMCPJSON" --repo AVmatrix
```

Observed:

- symbol: `Function:internal/cli/setup_command.go:setupWriteMCPJSON#1`
- lines: `207-222`
- incoming calls: `setupClaudeCode` at line `158`, `setupCursor` at line `143`
- outgoing calls: `setupReadJSONObject`, `setupWriteJSONObject`
- process membership: `RunSetup -> SetupReadJSONObject`, `RunSetup -> SetupWriteJSONObject`

`runSetup`:

```powershell
avmatrix context "runSetup" --repo AVmatrix
```

Observed:

- symbol: `Function:internal/cli/setup_command.go:runSetup#0`
- incoming call: `newSetupCommand`
- outgoing calls: `setupCursor`, `setupClaudeCode`, `setupOpenCode`, `setupCodex`, `repo.GlobalDir`

`renderAVmatrixBlock`:

```powershell
avmatrix context "renderAVmatrixBlock" --repo AVmatrix
```

Observed:

- symbol: `Function:internal/aicontext/aicontext.go:renderAVmatrixBlock#3`
- lines: `99-212`
- incoming call: `GenerateAIContextFiles`
- process membership: `NewAnalyzeCommand -> RenderAVmatrixBlock`

## E4 - Source Inspection Evidence

Date: 2026-05-29

Status: recorded

`internal/cli/setup_command.go` currently hardcodes:

```go
setupBrand         = "AVmatrix"
setupCommandName   = "avmatrix"
setupMCPServerName = "avmatrix"
```

The same file writes:

- Cursor/Claude config under `mcpServers[setupMCPServerName]`
- OpenCode config under `mcp[setupMCPServerName]`
- Codex command `codex mcp add <server> -- <command> mcp`
- Codex TOML `[mcp_servers.<server>]`
- hook commands/status messages containing the old brand

`internal/version/version.go` currently hardcodes:

```go
CommandName = "avmatrix"
Version     = "1.2.3"
```

`internal/repo/paths.go` currently hardcodes:

```go
StorageDirName = ".avmatrix"
HomeEnvName    = "AVMATRIX_HOME"
```

`internal/mcp/resources.go` currently hardcodes:

```go
canonicalResourceScheme = "avmatrix"
```

`cmd/avmatrix/main.go` is the current CLI entrypoint and imports the old Go module path:

```go
github.com/tamnguyendinh/avmatrix-go/internal/cli
```

`avmatrix/package.json` currently exposes:

- package name `avmatrix`
- description beginning `AVmatrix graph-powered...`
- GitHub URLs under `tamnguyendinh/AVmatrix`
- keyword `avmatrix`
- bin mapping `"avmatrix": "bin/avmatrix.exe"`
- lifecycle scripts using `cmd/avmatrix` and `bin/avmatrix.exe`

`avmatrix-web/index.html` currently has:

```html
<title>avmatrix</title>
```

Launcher source currently includes old protocol, executable, env, and package names:

- `avmatrix://reset`
- `AVmatrixLauncher.exe`
- `avmatrix-server.exe`
- `avmatrix.exe`
- `avmatrix-web`
- `avmatrix-launcher`
- `HKCU\Software\Classes\avmatrix`
- `AVMATRIX_LAUNCHER_NO_BROWSER`
- `AVMATRIX_GO`

## E5 - Initial Reference Inventory

Date: 2026-05-29

Status: recorded

Command shape:

```powershell
rg -o <pattern> --glob '!node_modules/**' --glob '!docs/plans/2026-05-29-anvien-rebrand-*'
```

Counts:

| Pattern | Count |
|---|---:|
| `AVmatrix` | 2238 |
| `avmatrix` | 9291 |
| `AVMATRIX` | 281 |
| `AVmatrix-GO` | 629 |
| `avmatrix.com` | 0 |
| `.avmatrix` | 316 |
| `AVMATRIX_` | 281 |
| `mcpServers` | 9 |

File-group count command:

```powershell
rg -l -i "avmatrix" --glob '!node_modules/**' --glob '!docs/plans/2026-05-29-anvien-rebrand-*'
```

Top groups:

| Group | Files |
|---|---:|
| `internal` | 338 |
| `reports` | 70 |
| `avmatrix-web` | 68 |
| `docs` | 55 |
| `baseline` | 19 |
| `avmatrix-launcher` | 6 |
| `avmatrix` | 5 |
| `cmd` | 5 |

## E5.1 - GitHub Automation Inventory

Date: 2026-05-29

Status: recorded

`.github` files present during plan update:

- `.github/dependabot.yml`
- `.github/FUNDING.yml`
- `.github/PULL_REQUEST_TEMPLATE.md`
- `.github/release-drafter.yml`
- `.github/release.yml`
- `.github/actions/setup-avmatrix/action.yml`
- `.github/actions/setup-avmatrix-web/action.yml`
- `.github/ISSUE_TEMPLATE/bug_report.yml`
- `.github/ISSUE_TEMPLATE/config.yml`
- `.github/ISSUE_TEMPLATE/feature_request.yml`
- `.github/scripts/check-tree-sitter-upgrade-readiness.py`
- `.github/scripts/check-workflow-concurrency.py`
- `.github/scripts/triage/**`
- `.github/workflows/ci-e2e.yml`
- `.github/workflows/ci-quality.yml`
- `.github/workflows/ci-report.yml`
- `.github/workflows/ci-tests.yml`
- `.github/workflows/ci.yml`
- `.github/workflows/claude-code-review.yml`
- `.github/workflows/claude.yml`
- `.github/workflows/docker.yml`
- `.github/workflows/pr-description-check.yml`
- `.github/workflows/pr-labeler.yml`
- `.github/workflows/publish.yml`
- `.github/workflows/release-candidate.yml`
- `.github/workflows/tree-sitter-upgrade-readiness.yml`
- `.github/workflows/triage-sweep.yml`

`.github` reference counts:

| Pattern | Count |
|---|---:|
| `AVmatrix` | 4 |
| `avmatrix` | 94 |
| `AVMATRIX` | 7 |
| `AVmatrix-GO` | 0 |
| `setup-avmatrix` | 6 |
| `github.com/tamnguyendinh/AVmatrix` | 0 |

Observed `.github` rename surfaces:

- composite action directories: `.github/actions/setup-avmatrix`, `.github/actions/setup-avmatrix-web`;
- workflow references to those composite actions;
- workflow working directories and cache paths for `avmatrix`, `avmatrix-web`, and `avmatrix-launcher`;
- E2E workflow build path `.tmp/avmatrix`, `cmd/avmatrix`, `.avmatrix` index checks, and `avmatrix serve`;
- Docker/GHCR workflow image slugs `avmatrix` and `avmatrix-web`;
- publish and release-candidate workflows using package name/directory `avmatrix`;
- issue and PR templates listing `avmatrix` and `avmatrix-web` components;
- CI report heredoc marker `AVMATRIX_CI_REPORT_EOF_7f3a`;
- scripts using `AVMATRIX_DIR = REPO_ROOT / "avmatrix"`.

Plan amendment:

- Phase 2 was expanded from a single GitHub rename item into a detailed GitHub execution checklist covering repo Settings, metadata, branch/ruleset checks, Actions secrets/variables/environments, integrations, release config, templates, `.github/actions`, `.github/workflows`, publish/Docker/GHCR, badges, local remote, fresh clone validation, and redirect evidence.

## E5.2 - Plan Review Gap Inventory

Date: 2026-05-29

Status: recorded

Graph refresh before this review:

```powershell
avmatrix analyze --force
```

Output summary:

```text
files: scanned=800 parsed=583 unsupported=217 failed=0
graph: nodes=91230 relationships=124709
```

Additional AVmatrix queries used:

- `avmatrix query "agent integration config mcp json codex claude grok plugin setup avmatrix" --repo AVmatrix`
- `avmatrix query "docker deploy compose release package publish ghcr npm installer avmatrix" --repo AVmatrix`
- `avmatrix query "contract generated web contracts package lock reports baseline snapshots avmatrix" --repo AVmatrix`
- `avmatrix query "configuration files dotfiles mcp.json agents codex claude grok avmatrix" --repo AVmatrix`

Additional checked-in old-name areas found by source inventory:

| Area | Old-name matches | Files |
|---|---:|---:|
| `.mcp.json` | 2 | 1 |
| `.grok` | 21 | 1 |
| `avmatrix-claude-plugin` | 113 | 16 |
| `contracts` | 4 | 1 |
| `avmatrix-web/src/generated` | 1 | 1 |

Specific files requiring plan coverage:

- `.mcp.json`
- `.grok/config.toml`
- `avmatrix-claude-plugin/.mcp.json`
- `avmatrix-claude-plugin/hooks/hooks.json`
- `avmatrix-claude-plugin/hooks/avmatrix-hook.js`
- `avmatrix-claude-plugin/skills/**/SKILL.md`
- `avmatrix-claude-plugin/skills/**/mcp.json`
- `contracts/web-ui/avmatrix-web-contract.schema.json`
- `avmatrix-web/src/generated/avmatrix-contracts.ts`

Current local git remote still points to the old repository URL:

```text
origin https://github.com/tamnguyendinh/AVmatrix.git
```

Plan amendments from this review:

- Added explicit coverage for checked-in `.mcp.json` and `.grok/config.toml`.
- Added a dedicated Claude plugin and agent integration package phase.
- Added generated Web contract filename/output rename requirements.
- Added file/directory-name inventory requirements beyond text content inventory.
- Added tracked historical report handling so old filenames do not silently remain active product surfaces.

## E5.3 - Generator Source Review

Date: 2026-05-29

Status: recorded

User review note:

- Generated files and folders must be checked through their source generators. Manually editing generated output is not enough, because the next `analyze`, `setup`, contract generation, build, package, or plugin install can recreate old `AVmatrix`/`avmatrix` names.

Graph refresh before this review:

```powershell
avmatrix analyze --force
```

Output summary:

```text
files: scanned=800 parsed=583 unsupported=217 failed=0
graph: nodes=91233 relationships=124712
```

Additional AVmatrix query used:

```powershell
avmatrix query "content generators write generated files folders AGENTS CLAUDE skills contracts web-dist setup mcp storage .avmatrix" --repo AVmatrix
```

Generator/source-of-truth evidence:

| Area | Source evidence | Generated or served output risk |
|---|---|---|
| AI context | `internal/aicontext/aicontext.go` defines `startMarker`, `endMarker`, `GenerateAIContextFiles`, `renderAVmatrixBlock`, `baseSkills`, `installBaseSkills`, and embeds `skills/*.md` | Recreates `AGENTS.md`, `CLAUDE.md`, `.claude/skills/avmatrix/**`, old markers, old skill ids, and command/resource tables |
| Embedded skills | `internal/aicontext/skills/avmatrix-*.md` | Reinstalled into repo/editor skill directories by analyze/setup |
| Setup/editor config | `internal/cli/setup_command.go` defines `setupBrand`, `setupCommandName`, `setupMCPServerName`, `setupWriteMCPJSON`, `setupWriteOpenCodeJSON`, `setupUpsertCodexToml`, `setupMergeClaudeHookSettings` | Recreates old MCP server keys, commands, Codex TOML, Claude hooks, and skill installs |
| Repo/global storage | `internal/repo/paths.go`, `meta.go`, `settings.go`, `registry.go`, `runtime_config.go` | Recreates `.avmatrix`, `AVMATRIX_HOME`, `graph.json`, `meta.json`, `settings.json`, `registry.json`, `runtime.json` paths |
| Group registry | `internal/group/storage.go`, `internal/cli/group_command.go`, `internal/mcp/group_tools.go` | Recreates group config/contract registry under the global storage root |
| MCP served text | `internal/mcp/resources.go`, `internal/mcp/prompts.go`, `internal/mcp/tools.go` | Recreates old resource schemes, setup resource text, prompt templates, tool descriptions, and next-step hints |
| Web contracts | `internal/contracts/web_ui.go`, `cmd/generate-web-contracts/main.go` | Recreates old generated schema/TypeScript filenames and imports |
| Launcher/package | `avmatrix-launcher/build.ps1`, `avmatrix-launcher/src/main.go`, `avmatrix-launcher/server-wrapper/main.go`, `internal/cli/package_runtime.go` | Recreates old package folders, `web-dist`, launcher state/log/output names, and executable artifact names |
| Hook/plugin | `internal/cli/hook_command.go`, `avmatrix-claude-plugin/**` | Recreates or ships old hook names, `.avmatrix` stale checks, plugin skill dirs, and plugin MCP config |
| Graph-quality helper defaults | `cmd/graph-accuracy-probe/main.go`, `internal/cli/resolution_inventory_command.go`, `internal/cli/source_site_accuracy_command.go` | Defaults to `.avmatrix/graph.json` in helper commands and generated report workflows |

Local generated storage observed in this workspace after analyze:

```text
.avmatrix/graph.json
.avmatrix/lbug
.avmatrix/meta.json
.avmatrix/settings.json
```

Code also defines or documents these storage artifacts:

- `.avmatrix/analyze.lock`
- `.avmatrix/analyze.tmp`
- `.avmatrix/lbug.wal`
- `.avmatrix/lbug.lock`
- `~/.avmatrix/registry.json`
- `~/.avmatrix/runtime.json`
- `~/.avmatrix/groups/<name>/group.yaml`
- `~/.avmatrix/groups/<name>/contracts.json`

Old-name generator baseline from source inventory:

| Generator area | Old-name matches | Files |
|---|---:|---:|
| AI context generator and embedded skills | 299 | 12 |
| Setup/editor config generator | 8 | 1 |
| Repo/global storage generators | 4 | 2 |
| MCP served setup/resources/prompts | 59 | 3 |
| Web contract generator | 11 | 2 |
| Group registry generators | 13 | 6 |
| Launcher/package generated outputs | 48 | 6 |
| Hook/plugin generated integration | 122 | 17 |
| Graph accuracy/default graph path helpers | 21 | 3 |

Plan amendments from this review:

- Added a generator/source-of-truth inventory table to the plan.
- Added Phase 1.5 for generator source audit before output edits.
- Expanded storage validation to include concrete `.anvien` artifacts, not only the top-level directory.
- Added generated marker validation for `AGENTS.md`, `CLAUDE.md`, and skill output so old markers and old skill ids do not remain as steady-state compatibility.
- Added benchmark rows for old-name counts inside generator source areas.

## E5.4 - Phase 1 Tracked File Inventory

Date: 2026-05-29

Status: recorded

Inventory artifact:

- `docs/plans/2026-05-29-anvien-rebrand-inventory.csv`

Inventory command classifies tracked files with old-name matches, excluding this rebrand plan/evidence/benchmark/inventory file set:

```powershell
git ls-files | Select-String -Pattern 'AVmatrix|avmatrix|AVMATRIX|AVmatrix-GO|\.avmatrix|avmatrix://|avmatrix-'
```

Inventory summary:

| Category | Files | Matches |
|---|---:|---:|
| `source-or-config` | 316 | 1546 |
| `report` | 70 | 923 |
| `web-ui` | 67 | 244 |
| `docs` | 55 | 8094 |
| `generator-source` | 31 | 183 |
| `baseline` | 19 | 250 |
| `generator-ai-context` | 13 | 384 |
| `launcher` | 6 | 56 |
| `npm-package` | 5 | 156 |
| `cli-command` | 3 | 4 |
| `generated-contract` | 2 | 5 |
| `github-automation` | 1 | 17 |

Phase/checklist impact:

- P1-A completed: tracked file inventory exists as CSV.
- P1-B completed: every inventory row has a classification such as `edit-generator`, `regenerate-from-generator`, `rename-in-place`, `hard-rename-or-remove`, `regenerate-or-update`, `preserve-or-delete-stale`, or `update-active-or-preserve-history`.
- P1-C/P1-H completed from generator/source-of-truth matrix.
- P1-F completed for tracked file paths because the CSV records paths containing old-name folder/file names.
- P1-G completed for local generated/cache/temp classification from E5.3.
- P1.5-A through P1.5-E completed from source reads and generator matrix evidence.

## E6 - Naming Decisions

Date: pending

Status: pending

| Surface | Decision | Evidence |
|---|---|---|
| Brand display name | `Anvien` | user request |
| CLI command | `anvien` | user clarified MCP/server name must not remain old name |
| MCP server name | `anvien` | user clarified MCP/server name must not remain old name |
| MCP resource scheme | `anvien://` | derived from hard-rename requirement |
| GitHub owner/repo slug | pending | pending |
| Local folder path | pending | pending |
| Go module path | pending | pending |
| npm package name | `anvien` unless unavailable | pending registry check |
| Storage dir | `.anvien` | derived from no-legacy rule |
| Env prefix | `ANVIEN_` | derived from no-legacy rule |
| Domain strategy | no action on third-party `avmatrix.com` | user clarification |

## E7 - Phase 6 AI Context Generator Slice

Date: 2026-05-29

Status: recorded

Scope:

- Updated the AI context generator source in `internal/aicontext/aicontext.go`.
- Renamed embedded skill source files from `internal/aicontext/skills/avmatrix-*.md` to `internal/aicontext/skills/anvien-*.md`.
- Updated generated-context tests in `internal/aicontext/aicontext_test.go`.

Graph and impact evidence:

- Fresh graph before validation: `.\avmatrix\bin\avmatrix.exe analyze --force`.
- Analyze result after the final pre-commit refresh: scanned `801`, parsed `583`, unsupported `218`, failed `0`; graph `91263` nodes, `124743` relationships.
- `impact GenerateAIContextFiles --repo AVmatrix --direction upstream`: CRITICAL; `4` impacted nodes, `23` affected processes, app layer `backend`, functional area mostly `cli`.
- `impact renderAnvienBlock --repo AVmatrix --direction upstream`: CRITICAL; `4` impacted nodes, `12` affected processes.
- `impact InstallBaseSkillsTo --repo AVmatrix --direction upstream`: CRITICAL; `10` impacted nodes, `9` affected processes, covers analyze and setup skill installation paths.
- `impact BaseSkillFiles --repo AVmatrix --direction upstream`: HIGH; `5` impacted nodes, `3` affected processes.

Implementation notes:

- Used graph-guided rename for `renderAVmatrixBlock` to `renderAnvienBlock`.
- Generated block markers are now `<!-- anvien:start -->` and `<!-- anvien:end -->`.
- Generated command/resource examples now use `anvien`, `anvien://`, and `.anvien`.
- Generated skill namespace is now `.claude/skills/anvien/**`.
- Steady-state generated output no longer creates `.claude/skills/avmatrix/**`.
- The implementation keeps only constructed cleanup strings for removing stale installed old skill directories. They are not a supported alias or active output.
- One old-name source match remains in this slice: the Go module import path `github.com/tamnguyendinh/avmatrix-go/internal/analyze`; it belongs to the later module-path rename slice.

Validation:

| Command | Result |
|---|---|
| `powershell -ExecutionPolicy Bypass -File avmatrix-launcher\build.ps1` | first run failed while copying `avmatrix\bin\lbug_shared.dll` because two global `avmatrix.exe mcp` processes held runtime state; stopped PIDs `7436` and `12356`, then reran successfully. Existing Vite dynamic-import and chunk-size warnings only. |
| `go test .\internal\aicontext -count=1` | pass; `1.038s`. |
| `.\avmatrix\bin\avmatrix.exe analyze --force` | pass; regenerated AI context and local skill output. |
| `rg -n "AVmatrix|avmatrix|AVMATRIX|\.avmatrix|avmatrix://|avmatrix-" internal\aicontext` | `1` match in `internal\aicontext\aicontext.go` for the old Go module import path only. |
| `rg -n "AVmatrix|avmatrix|AVMATRIX|\.avmatrix|avmatrix://|avmatrix-" AGENTS.md CLAUDE.md .claude\skills\anvien` | `6` matches, all in `AGENTS.md`/`CLAUDE.md` outside generated skill files: the user-maintained top coding rules and the indexed repo name `AVmatrix`. Generated `.claude\skills\anvien` had no old-name matches. |
| `Test-Path .claude\skills\avmatrix` | `False`. |

E2E status:

- Not run for this slice because it changed backend-generated AI context/skill files and no Web UI behavior.

Pre-commit change detection:

- Command: `.\avmatrix\bin\avmatrix.exe detect-changes --repo AVmatrix --scope all`.
- Result: pass; summary risk `medium`.
- Changed scope: `16` files, `161` changed symbols; app layers `backend` `100`, `backend_test` `48`, `docs` `13`.
- Affected runtime processes: `3`, all in the analyze AI context flow: `NewAnalyzeCommand -> RemoveGeneratedSkills`, `NewAnalyzeCommand -> RenderAnvienBlock`, and `NewAnalyzeCommand -> UpsertSection`.
- Resolution health impact: `0` degraded nodes and `0` nodes with gaps.

## E8 - Phase 4/5 Setup And Storage Generator Slice

Date: 2026-05-29

Status: recorded

Scope:

- Updated repo/global storage source in `internal/repo/paths.go` from `.avmatrix`/`AVMATRIX_HOME` to `.anvien`/`ANVIEN_HOME`.
- Updated storage consumers and diagnostics in lock handling, status/index guidance, graph-health/query-health messages, source-site/resolution-inventory defaults, HTTP graph/embed errors, and Claude hook stale-index checks.
- Updated setup generator source in `internal/cli/setup_command.go` so editor MCP config, Codex TOML, Codex MCP add command, Claude hook command, and setup output generate `Anvien`/`anvien`.
- Updated `.gitignore` from `.avmatrix/` to `.anvien/`.
- Removed `478` untracked generated `.avmatrix` directories from the local workspace after verifying no tracked files were under `.avmatrix/`.

Graph and impact evidence:

- Fresh graph before storage impact checks: `.\avmatrix\bin\avmatrix.exe analyze --force`.
- `impact Paths --repo AVmatrix --direction upstream`: CRITICAL; affected app layers `api:6`, `backend:18`, `cli_launcher:2`; affected functional areas include analyzer, api, cli, graph_health, session, and storage; `71` affected processes.
- `impact GlobalDir --repo AVmatrix --direction upstream`: CRITICAL; affected app layers `api:3`, `backend:51`; affected functional areas include api, cli, mcp, query, and storage; `80` affected processes.
- `impact StorageDirName --repo AVmatrix --direction upstream`: LOW; `0` affected processes.
- `impact HomeEnvName --repo AVmatrix --direction upstream`: LOW; `0` affected processes.
- `impact lockHeldNextAction`: LOW; `1` impacted node. `impact lockInfoRepoHint`: LOW; `2` impacted nodes.
- Additional edited command/API helpers were checked before edits. CRITICAL examples: `newIndexCommand`, `resolveIndexPath`, `newStatusCommand`, `loadGraphHealthGraph`, `verifyQueryHealthFreshRepo`, `newSourceSiteAccuracyCommand`, `newResolutionInventoryCommand`, `handleClaudePostToolUse`, `resolveRepoQuery`, and `loadGraphSnapshot`.
- Hook helper rename impact: `findClaudeHookAVmatrixDir` LOW, then renamed to `findClaudeHookAnvienDir`.
- Fresh graph before setup generator edits: `.\avmatrix\bin\avmatrix.exe analyze --force`.
- `impact newSetupCommand --repo AVmatrix --direction upstream`: CRITICAL; affected app layers `backend:1`, `cli_launcher:1`; `11` affected processes through `NewRootCommand` and CLI `main`.
- `impact runSetup --repo AVmatrix --direction upstream`: CRITICAL; affected app layers `backend:2`, `cli_launcher:1`; `11` affected processes.
- `impact setupMCPServerName` and `impact setupCommandName`: LOW; `0` affected processes.
- `impact TestFindClaudeHookAVmatrixDirWalksParents`: LOW; `0` affected processes, then test renamed to `TestFindClaudeHookAnvienDirWalksParents`.

Implementation notes:

- `repo.StorageDirName` is now `.anvien`.
- `repo.HomeEnvName` is now `ANVIEN_HOME`; no `AVMATRIX_HOME` fallback was added.
- `repo.GlobalDir()` falls back to `~/.anvien`.
- Default graph helper paths now use `.anvien/graph.json`.
- Hook stale-index messages now say `Anvien index is stale` and suggest `anvien analyze`.
- Setup now writes MCP server key `anvien` and command `anvien` for Cursor, Claude Code, OpenCode, and Codex.
- Setup still removes stale old hook entries by using a constructed cleanup needle. This is cleanup only, not a supported legacy hook path or alias.
- Remaining old-name matches in touched source are Go module import paths or later-slice package/plugin surfaces, not active storage/setup output.

Validation:

| Command | Result |
|---|---|
| `powershell -ExecutionPolicy Bypass -File avmatrix-launcher\build.ps1` | pass after setup/storage edits; existing Vite dynamic-import and chunk-size warnings only. |
| `go test .\internal\repo .\internal\cli .\internal\httpapi .\internal\mcp .\internal\group -count=1` | pass; repo `3.196s`, cli `13.009s`, httpapi `3.591s`, mcp `8.032s`, group `2.053s`. |
| Temp repo smoke with `ANVIEN_HOME=<temp>\home` and `.\avmatrix\bin\avmatrix.exe analyze <temp>\repo --force --skip-git --name storage-smoke` | pass; generated `<repo>\.anvien\graph.json`, `meta.json`, `settings.json`, `lbug`, and `<home>\registry.json`; `<repo>\.avmatrix` was `False`. |
| `git ls-files \| rg '(^\|/)\.avmatrix/'` before local cleanup | no tracked `.avmatrix/` files. |
| `Get-ChildItem -Path . -Directory -Recurse -Force -Filter .avmatrix` after cleanup | no `.avmatrix` directories remained in the workspace. |

E2E status:

- Not run for this slice because no Web UI behavior changed.

Pre-commit change detection:

- Command: `.\avmatrix\bin\avmatrix.exe detect-changes --repo AVmatrix --scope all`.
- Result: pass; summary risk `critical`.
- Changed scope: `28` files, `311` changed symbols.
- Changed app layers: `api` `9`, `api_test` `41`, `backend` `67`, `backend_test` `182`, `docs` `12`.
- Changed functional areas: `api` `50`, `cli` `187`, `documentation` `12`, `storage` `62`.
- Affected scope: `73` affected symbols/process nodes; affected app layers `api` `3`, `backend` `24`, `mixed` `46`.
- Resolution health impact: `0` degraded nodes and `0` nodes with gaps.
- The CRITICAL summary is expected for this slice because storage path constants and setup generator output sit under CLI/API startup and graph-loading flows.

## E9 - Phase 4 MCP Served Resource Slice

Date: 2026-05-29

Status: recorded

Scope:

- Updated root `.mcp.json` from server key/command `avmatrix` to `anvien`.
- Updated `internal/mcp/resources.go` so `canonicalResourceScheme` and served resource URIs use `anvien://`.
- Updated `internal/mcp/prompts.go`, `internal/mcp/server.go`, and `internal/mcp/tools.go` so served prompt text, next-step hints, `serverInfo.name`, tool descriptions, and stale-index guidance use Anvien names.
- Updated MCP tests and TypeScript baseline surface testdata to expect `anvien://` and MCP server name `anvien`.

Graph and impact evidence:

- Fresh graph before MCP impact checks: `.\avmatrix\bin\avmatrix.exe analyze --force`.
- Analyze result before impact checks: scanned `801`, parsed `583`, unsupported `218`, failed `0`; graph `91274` nodes, `124754` relationships.
- Final graph refresh before change detection: scanned `801`, parsed `583`, unsupported `218`, failed `0`; graph `91276` nodes, `124756` relationships.
- `impact canonicalResourceScheme --repo AVmatrix --direction upstream`: LOW; `0` affected processes.
- `impact resourceDefinitions --repo AVmatrix --direction upstream`: CRITICAL; affected app layer `api`, functional area `mcp`, `9` affected processes through `Server.handle`.
- `impact readResourceText --repo AVmatrix --direction upstream`: LOW; `0` affected processes.
- `impact setupResource --repo AVmatrix --direction upstream`: LOW; `0` affected processes.
- `impact promptDefinitions --repo AVmatrix --direction upstream`: CRITICAL; affected app layer `api`, functional area `mcp`, `9` affected processes through `Server.handle`.
- `impact detectImpactPrompt --repo AVmatrix --direction upstream`: CRITICAL; affected app layer `api`, functional area `mcp`, `9` affected processes through `getPrompt` and `Server.handle`.
- `impact generateMapPrompt --repo AVmatrix --direction upstream`: CRITICAL; affected app layer `api`, functional area `mcp`, `9` affected processes through `getPrompt` and `Server.handle`.
- `impact nextStepHint --repo AVmatrix --direction upstream`: LOW; direct caller `Server.callTool`.
- `impact mcpTools --repo AVmatrix --direction upstream`: CRITICAL; affected app layer `api`, functional area `mcp`, `9` affected processes through `Server.handle`.
- `impact querySemanticWarning --repo AVmatrix --direction upstream`: CRITICAL; affected app layer `api`, functional area `mcp`, `37` affected processes across query/context/impact/API response flows.
- `impact --uid "Method:internal/mcp/server.go:Server.initialize#1" --repo AVmatrix --direction upstream`: LOW; `0` affected processes.

Implementation notes:

- MCP `resources/list`, resource templates, and `resources/read` now expose `anvien://repos`, `anvien://setup`, and `anvien://repo/...`.
- MCP prompt text now tells agents to read `anvien://repos` and Anvien repo resources.
- MCP initialize response now reports `serverInfo.name = "anvien"`.
- Root `.mcp.json` now uses server key `anvien`, command `anvien`, and args `["mcp"]`.
- Remaining old-name matches in touched MCP source are Go module import paths and path heuristics for still-unrenamed folders such as `avmatrix-web` and `cmd/avmatrix`; they are deferred to package/module/folder rename slices, not served MCP aliases.
- `.grok/config.toml` was not updated in this commit because `git ls-files .grok .mcp.json` showed only `.mcp.json` as tracked.

Validation:

| Command | Result |
|---|---|
| `powershell -ExecutionPolicy Bypass -File avmatrix-launcher\build.ps1` | pass after MCP served-resource edits; existing Vite dynamic-import and chunk-size warnings only. |
| `go test .\internal\mcp .\internal\httpapi .\internal\cli .\internal\aicontext -count=1` | pass; mcp `9.242s`, httpapi `4.098s`, cli `15.184s`, aicontext `1.061s`. |
| `go test .\internal\mcp -run "TestServeHandlesInitializeAndToolsList|TestServeReadsRepoContextResource|TestServeReadsRepoClustersAndProcessesResources|TestServeReadsSchemaDetailResourcesAndPrompts|TestGenerateMapPrompt|TestDetectImpactPrompt|TestResourceDefinitionsAndTemplatesParity|TestReadResourceText" -count=1 -v` | pass; MCP protocol/resource/prompt smoke `3.459s`. |
| `rg -n 'avmatrix://|AVmatrix|avmatrix analyze|avmatrix api|avmatrix rename|avmatrix setup|mcpServers.*avmatrix|\[mcp_servers\.avmatrix\]|"avmatrix"' internal\mcp\resources.go internal\mcp\prompts.go internal\mcp\server.go internal\mcp\tools.go internal\mcp\resources_parity_test.go internal\mcp\prompts_test.go internal\mcp\server_test.go internal\mcp\testdata\typescript_baseline_surface.json .mcp.json` | no matches. |
| `rg -n "anvien://" internal\mcp\resources.go internal\mcp\prompts.go internal\mcp\server.go internal\mcp\tools.go internal\mcp\resources_parity_test.go internal\mcp\prompts_test.go internal\mcp\server_test.go internal\mcp\testdata\typescript_baseline_surface.json .mcp.json` | positive matches in MCP prompt tests, baseline testdata, server tests, resource tests, `prompts.go`, and `server.go`. |
| `git ls-files .grok .mcp.json` | only `.mcp.json` is tracked. |

E2E status:

- Not run for this slice because no Web UI behavior changed.

Pre-commit change detection:

- Command: `.\avmatrix\bin\avmatrix.exe detect-changes --repo AVmatrix --scope all`.
- Result: pass; summary risk `high`.
- Changed scope: `12` files, `106` changed symbols.
- Changed app layers: `api` `28`, `api_test` `66`, `docs` `12`.
- Changed functional areas: `mcp` `94`, `documentation` `12`.
- Affected scope: `10` affected symbols/process nodes; affected app layers `api` `9`, `mixed` `1`; affected functional areas `mcp` `9`, `mixed` `1`.
- Resolution health impact: `0` degraded nodes and `0` nodes with gaps.
- The HIGH summary is expected for this slice because served MCP resources/prompts/tools flow through `Server.handle` and agent-facing MCP protocol responses.

## E10 - Future Implementation Evidence

Date: pending

Status: pending

Record during implementation:

- impact outputs for every edited symbol;
- changed files by slice;
- build/test/e2e output;
- MCP protocol smoke output after the MCP scheme/server rename;
- package/install smoke output for `anvien`;
- final old-name inventory and every remaining exception.
