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

## E7 - Future Implementation Evidence

Date: pending

Status: pending

Record during implementation:

- impact outputs for every edited symbol;
- changed files by slice;
- build/test/e2e output;
- MCP protocol smoke output after the MCP scheme/server rename;
- package/install smoke output for `anvien`;
- final old-name inventory and every remaining exception.
