# Anvien Hard Rebrand Benchmark Ledger

Date: 2026-05-29

Status: Draft

Companion files:

- Plan: [2026-05-29-anvien-rebrand-plan.md](2026-05-29-anvien-rebrand-plan.md)
- Evidence ledger: [2026-05-29-anvien-rebrand-evidence.md](2026-05-29-anvien-rebrand-evidence.md)

## Benchmark Rules

This file records quantitative data only: old-name counts, active legacy-surface counts, graph inventory counts, package/startup sizes or timings, and validation pass/fail counts.

Narrative evidence and command excerpts belong in the evidence ledger.

## B0 - Graph Baseline

Status: recorded

| Metric | Unit | Baseline | Latest | Delta | Notes |
|---|---:|---:|---:|---:|---|
| Files scanned | files | 800 | 801 | +1 | `avmatrix analyze --force` after MCP served-resource slice |
| Files parsed | files | 583 | 583 | 0 | `avmatrix analyze --force` |
| Unsupported files | files | 217 | 218 | +1 | `avmatrix analyze --force` after setup/storage slice |
| Failed files | files | 0 | 0 | 0 | `avmatrix analyze --force` |
| Graph nodes | nodes | 91223 | 91276 | +53 | Fresh graph after MCP served-resource slice |
| Graph relationships | relationships | 124702 | 124756 | +54 | Fresh graph after MCP served-resource slice |

## B1 - Old-Name Reference Baseline

Status: recorded

Search excluded `node_modules` and this rebrand file set.

| Pattern | Unit | Baseline | Latest | Delta | Final target |
|---|---:|---:|---:|---:|---:|
| `AVmatrix` | matches | 2238 | 2238 | 0 | 0 active |
| `avmatrix` | matches | 9291 | 9291 | 0 | 0 active |
| `AVMATRIX` | matches | 281 | 281 | 0 | 0 active |
| `AVmatrix-GO` | matches | 629 | 629 | 0 | 0 active |
| `avmatrix.com` | matches | 0 | 0 | 0 | 0 |
| `.avmatrix` | matches | 316 | 316 | 0 | 0 active |
| `AVMATRIX_` | matches | 281 | 281 | 0 | 0 active |
| `mcpServers` | matches | 9 | 9 | 0 | inspect/update keys |

## B2 - Old-Name File Group Baseline

Status: recorded

| Group | Unit | Baseline | Latest | Delta | Final target |
|---|---:|---:|---:|---:|---:|
| `internal` | files | 338 | 338 | 0 | 0 active old names |
| `reports` | files | 70 | 70 | 0 | delete/update/classify |
| `avmatrix-web` | files | 68 | 68 | 0 | rename/update |
| `docs` | files | 55 | 55 | 0 | update active docs |
| `baseline` | files | 19 | 19 | 0 | regenerate/update active baselines |
| `avmatrix-launcher` | files | 6 | 6 | 0 | rename/update |
| `avmatrix` | files | 5 | 5 | 0 | rename package |
| `cmd` | files | 5 | 5 | 0 | rename entrypoint/imports |

## B2.1 - Tracked Inventory Classification Baseline

Status: recorded

Source artifact: `docs/plans/2026-05-29-anvien-rebrand-inventory.csv`

| Category | Unit | Baseline | Latest | Delta | Final target |
|---|---:|---:|---:|---:|---:|
| `source-or-config` | files | 316 | 316 | 0 | 0 active old names |
| `source-or-config` | matches | 1546 | 1546 | 0 | 0 active old names |
| `report` | files | 70 | 70 | 0 | classify/delete/update |
| `report` | matches | 923 | 923 | 0 | classify/delete/update |
| `web-ui` | files | 67 | 67 | 0 | 0 active old names |
| `web-ui` | matches | 244 | 244 | 0 | 0 active old names |
| `docs` | files | 55 | 55 | 0 | update active or preserve historical |
| `docs` | matches | 8094 | 8094 | 0 | update active or preserve historical |
| `generator-source` | files | 31 | 31 | 0 | 0 active old names |
| `generator-source` | matches | 183 | 183 | 0 | 0 active old names |
| `baseline` | files | 19 | 19 | 0 | regenerate/update active |
| `baseline` | matches | 250 | 250 | 0 | regenerate/update active |
| `generator-ai-context` | files | 13 | 1 | -12 | 0 active old names |
| `generator-ai-context` | matches | 384 | 1 | -383 | 0 active old names |
| `launcher` | files | 6 | 6 | 0 | 0 active old names |
| `launcher` | matches | 56 | 56 | 0 | 0 active old names |
| `npm-package` | files | 5 | 5 | 0 | 0 active old names |
| `npm-package` | matches | 156 | 156 | 0 | 0 active old names |
| `cli-command` | files | 3 | 3 | 0 | 0 active old names |
| `cli-command` | matches | 4 | 4 | 0 | 0 active old names |
| `generated-contract` | files | 2 | 2 | 0 | regenerate from generator |
| `generated-contract` | matches | 5 | 5 | 0 | regenerate from generator |
| `github-automation` | files | 1 | 1 | 0 | 0 active old names |
| `github-automation` | matches | 17 | 17 | 0 | 0 active old names |

## B3 - Active Legacy Surface Count

Status: baseline recorded

| Surface | Unit | Baseline | Latest | Delta | Final target |
|---|---:|---:|---:|---:|---:|
| Old CLI command names accepted/generated | count | 1 | 0 | -1 | 0 |
| Old MCP server names generated | count | 1 | 0 | -1 | 0 |
| Old MCP resource schemes generated | count | 1 | 0 | -1 | 0 |
| Old repo/global storage dirs generated | count | 2 | 0 | -2 | 0 |
| Old env var prefixes read | count | 1+ | 1+ | 0 | 0 |
| Old package/bin names generated | count | 1+ | 0 | -1+ | 0 |
| Old launcher protocol/executable names generated | count | 3+ | 0 | -3+ | 0 |
| Old Go module path in active Go imports/module declaration | count | 696 | 0 | -696 | 0 |
| Old generated skill namespace generated | count | 1 | 0 | -1 | 0 |

## B3.1 - GitHub Automation Old-Name Baseline

Status: recorded

| Pattern | Unit | Baseline | Latest | Delta | Final target |
|---|---:|---:|---:|---:|---:|
| `.github` `AVmatrix` references | matches | 4 | 4 | 0 | 0 active |
| `.github` `avmatrix` references | matches | 94 | 94 | 0 | 0 active |
| `.github` `AVMATRIX` references | matches | 7 | 7 | 0 | 0 active |
| `.github` `AVmatrix-GO` references | matches | 0 | 0 | 0 | 0 |
| `.github` `setup-avmatrix` references | matches | 6 | 6 | 0 | 0 active |
| `.github` old GitHub URL references | matches | 0 | 0 | 0 | 0 |

## B3.2 - Agent Integration And Generated Contract Old-Name Baseline

Status: recorded

| Area | Unit | Baseline | Latest | Delta | Final target |
|---|---:|---:|---:|---:|---:|
| `.mcp.json` old-name matches | matches | 2 | 0 | -2 | 0 |
| `.mcp.json` files with old names | files | 1 | 0 | -1 | 0 |
| `.grok` old-name matches | matches | 21 | 21 | 0 | 0 |
| `.grok` files with old names | files | 1 | 1 | 0 | 0 |
| `avmatrix-claude-plugin` old-name matches | matches | 113 | 113 | 0 | 0 active |
| `avmatrix-claude-plugin` files with old names | files | 16 | 16 | 0 | 0 active |
| `contracts` old-name matches | matches | 4 | 4 | 0 | 0 active |
| `contracts` files with old names | files | 1 | 1 | 0 | 0 active |
| `avmatrix-web/src/generated` old-name matches | matches | 1 | 1 | 0 | 0 active |
| `avmatrix-web/src/generated` files with old names | files | 1 | 1 | 0 | 0 active |

## B3.3 - Generator Source Old-Name Baseline

Status: recorded

| Generator area | Unit | Baseline | Latest | Delta | Final target |
|---|---:|---:|---:|---:|---:|
| AI context generator and embedded skills | matches | 299 | 0 | -299 | 0 active |
| AI context generator and embedded skills | files | 12 | 0 | -12 | 0 active |
| Setup/editor config generator | matches | 8 | 0 | -8 | 0 active |
| Setup/editor config generator | files | 1 | 0 | -1 | 0 active |
| Repo/global storage generators | matches | 4 | 0 | -4 | 0 active |
| Repo/global storage generators | files | 2 | 0 | -2 | 0 active |
| MCP served setup/resources/prompts | matches | 59 | 9 | -50 | 0 active |
| MCP served setup/resources/prompts | files | 3 | 1 | -2 | 0 active |
| Web contract generator | matches | 11 | 11 | 0 | 0 active |
| Web contract generator | files | 2 | 2 | 0 | 0 active |
| Group registry generators | matches | 13 | 13 | 0 | 0 active |
| Group registry generators | files | 6 | 6 | 0 | 0 active |
| Launcher/package generated outputs | matches | 48 | 48 | 0 | 0 active |
| Launcher/package generated outputs | files | 6 | 6 | 0 | 0 active |
| Hook/plugin generated integration | matches | 122 | 122 | 0 | 0 active |
| Hook/plugin generated integration | files | 17 | 17 | 0 | 0 active |
| Graph accuracy/default graph path helpers | matches | 21 | 21 | 0 | 0 active |
| Graph accuracy/default graph path helpers | files | 3 | 3 | 0 | 0 active |

## B4 - Future Runtime/Package Metrics

Status: pending implementation

| Metric | Unit | Baseline | Latest | Delta | Target |
|---|---:|---:|---:|---:|---:|
| `anvien.exe` size | bytes | pending | 50512384 | recorded | record |
| `AnvienLauncher.exe` size | bytes | pending | 6993408 | recorded | record |
| npm package tarball size | bytes | pending | pending | pending | record |
| CLI startup time | ms | pending | pending | pending | no unintended regression |
| MCP tools/list pass count | tests | pending | pending | pending | pass |
| MCP resources/list pass count | tests | pending | pending | pending | pass |
| MCP `anvien://setup` smoke count | tests | pending | pending | pending | pass |
| Web e2e Anvien branding checks | tests | pending | pending | pending | pass |

## B5 - Phase 6 AI Context Generator Counts

Status: recorded

Date: 2026-05-29

| Metric | Command | Latest | Note |
|---|---|---:|---|
| AI context generator old-name matches | `rg -n "AVmatrix|avmatrix|AVMATRIX|\.avmatrix|avmatrix://|avmatrix-" internal\aicontext` | `0` | Module rename slice removed the last deferred old import path from `internal\aicontext\aicontext.go`. |
| AI context generator files with old-name matches | same command, unique file count | `0` | No old-name matches remain in AI context generator sources. |
| Generated `.claude\skills\anvien` old-name matches | `rg -n "AVmatrix|avmatrix|AVMATRIX|\.avmatrix|avmatrix://|avmatrix-" .claude\skills\anvien` | `0` | Generated skill output is clean for this slice. |
| Old generated skill namespace exists | `Test-Path .claude\skills\avmatrix` | `0` | `False`; analyze regeneration did not recreate the old skill namespace. |
| Generated context old-name matches | `rg -n "AVmatrix|avmatrix|AVMATRIX|\.avmatrix|avmatrix://|avmatrix-" AGENTS.md CLAUDE.md .claude\skills\anvien` | `6` | All matches are outside generated skill files: top coding rules and indexed repo name still `AVmatrix` until repo/module rename. |

## B6 - Phase 4/5 Setup And Storage Counts

Status: recorded

Date: 2026-05-29

| Metric | Command | Latest | Note |
|---|---|---:|---|
| Setup generator old-name matches | `rg -n "AVmatrix|avmatrix|AVMATRIX|\.avmatrix|avmatrix://|avmatrix-" internal\cli\setup_command.go` | `0` | Module rename slice removed the last deferred old import paths from setup generator source. |
| Repo/global storage source old-name matches | `rg -n "AVmatrix|avmatrix|AVMATRIX|\.avmatrix|avmatrix://|avmatrix-" internal\repo\paths.go internal\repo\meta.go internal\repo\settings.go internal\repo\registry.go internal\repo\runtime_config.go internal\repo\lock.go` | `0` | Core storage generators no longer emit `.avmatrix` or `AVMATRIX_HOME`. |
| Temp smoke required `.anvien` artifacts present | analyze smoke on temp repo with `ANVIEN_HOME` | `6` | `.anvien`, `graph.json`, `meta.json`, `settings.json`, `lbug`, and global `registry.json`. |
| Temp smoke old `.avmatrix` artifact present | same smoke | `0` | `<repo>\.avmatrix` was `False`. |
| Local generated `.avmatrix` dirs removed | `Get-ChildItem -Directory -Recurse -Force -Filter .avmatrix` before cleanup | `478` | Verified no tracked files under `.avmatrix/` before removal. |

## B7 - Phase 4 MCP Served Resource Counts

Status: recorded

Date: 2026-05-29

| Metric | Command | Latest | Note |
|---|---|---:|---|
| MCP active old served names/schemes | `rg -n 'avmatrix://|AVmatrix|avmatrix analyze|avmatrix api|avmatrix rename|avmatrix setup|mcpServers.*avmatrix|\[mcp_servers\.avmatrix\]|"avmatrix"' internal\mcp\resources.go internal\mcp\prompts.go internal\mcp\server.go internal\mcp\tools.go internal\mcp\resources_parity_test.go internal\mcp\prompts_test.go internal\mcp\server_test.go internal\mcp\testdata\typescript_baseline_surface.json .mcp.json` | `0` | No old served MCP scheme, server key, command examples, or server name remains in the edited MCP surfaces. |
| MCP broad old-name source matches | `rg -n "AVmatrix|avmatrix|AVMATRIX|\.avmatrix|avmatrix://|avmatrix-" internal\mcp\resources.go internal\mcp\prompts.go internal\mcp\server.go internal\mcp\tools.go` | `9` | Remaining matches are deferred folder-path heuristics for `avmatrix-web` and `cmd/avmatrix`; module import paths are clean. |
| MCP broad old-name source files | same command, unique file count | `1` | `tools.go`. |
| Root `.mcp.json` old-name matches | same active old served-name search including `.mcp.json` | `0` | Root MCP config now uses key and command `anvien`. |
| Tracked `.grok` config files | `git ls-files .grok .mcp.json` | `0` | `.grok/config.toml` is local-untracked in this workspace; P4-K remains open outside this commit. |
| MCP protocol/resource/prompt smoke tests | targeted `go test .\internal\mcp -run ... -count=1 -v` | `11` | All targeted MCP protocol/resource/prompt tests passed. |

## B8 - Phase 3 CLI/Package/Launcher Runtime Counts

Status: recorded

Date: 2026-05-29

| Metric | Command | Latest | Note |
|---|---|---:|---|
| Active old CLI/package/launcher runtime-name matches | targeted `rg` over touched CLI/package/launcher/config surfaces | `5` | All `5` are cleanup-only stale artifact variables in `avmatrix-launcher\build.ps1`; non-cleanup matches were `0`. |
| Non-cleanup active old runtime-name matches | same search excluding `Legacy*` cleanup variables | `0` | No old CLI command, binary path, MCP scheme, launcher protocol, or artifact name is generated by touched runtime paths. |
| Positive Anvien runtime-name matches | targeted Anvien `rg` over touched surfaces | `60` | Confirms `anvien.exe`, `AnvienLauncher.exe`, `anvien-server.exe`, `anvien://`, and `ANVIEN_*` replacements are present. |
| Old generated artifact files present locally | `Test-Path` for `avmatrix.exe`, `avmatrix.exe~`, `avmatrix-runtime.json`, `AVmatrixLauncher.exe`, `avmatrix-server.exe` | `0` | All checks returned `False` after build and cleanup. |
| Anvien generated artifact files present locally | `Test-Path` for `anvien.exe`, `anvien-runtime.json`, `AnvienLauncher.exe`, `anvien-server.exe` | `4` | All required generated artifacts are present. |
| `anvien.exe` size | `Get-Item avmatrix\bin\anvien.exe` | `50495488` | Built by package lifecycle validation. |
| `anvien-runtime.json` size | `Get-Item avmatrix\bin\anvien-runtime.json` | `136` | Written by `npm run build` package lifecycle. |
| `AnvienLauncher.exe` size | `Get-Item avmatrix-launcher\AnvienLauncher.exe` | `6993408` | Built by launcher build script. |
| `anvien-server.exe` size | `Get-Item avmatrix-launcher\server-bundle\anvien-server.exe` | `2053632` | Built by launcher build script. |
| MCP initialize smoke through `anvien mcp` | piped initialize frame to `.\avmatrix\bin\anvien.exe mcp` | `1` | Response had `serverInfo.name = "anvien"`. |

## B9 - Phase 2.5 Go Module Path Counts

Status: recorded

Date: 2026-05-29

| Metric | Command | Latest | Note |
|---|---|---:|---|
| Old Go module path matches in active Go imports and `go.mod` | `rg -o "github.com/tamnguyendinh/avmatrix-go" --glob "*.go" .` plus `go.mod` | `0` | No active Go import or module declaration uses the old path. |
| New Go module path matches in active Go imports and `go.mod` | `rg -o "github.com/tamnguyendinh/anvien" --glob "*.go" .` plus `go.mod` | `696` | Includes `695` Go import matches and `1` module declaration. |
| Go files changed by module import rewrite | `git diff --name-only` filtered to `*.go` before commit | `302` | Mechanical exact-string import rewrite plus `gofmt`. |
| Historical old module-path matches outside this rebrand ledger | `git grep -n "github.com/tamnguyendinh/avmatrix-go" -- ':!docs/plans/2026-05-29-anvien-rebrand-*'` | `71` | Historical docs/reports only; no runtime alias. |
| Package runtime `anvien.exe` size after module rename | `Get-Item avmatrix\bin\anvien.exe` | `50512384` | Rebuilt by `npm run build` after module rename. |
| Package runtime metadata size after module rename | `Get-Item avmatrix\bin\anvien-runtime.json` | `136` | Rebuilt by package lifecycle. |
