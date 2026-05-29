# Anvien Hard Rebrand Plan

Date: 2026-05-29

Status: In progress

Companion files:

- Evidence ledger: [2026-05-29-anvien-rebrand-evidence.md](2026-05-29-anvien-rebrand-evidence.md)
- Benchmark ledger: [2026-05-29-anvien-rebrand-benchmark.md](2026-05-29-anvien-rebrand-benchmark.md)
- Inventory ledger: [2026-05-29-anvien-rebrand-inventory.csv](2026-05-29-anvien-rebrand-inventory.csv)

## Master Rules

1. Use AVmatrix for codebase analysis and impact checks while working on this plan, including documentation planning when the plan depends on actual codebase surfaces.
2. As each task is completed, update the corresponding checklist item immediately.
3. Run a full build before testing; the test suite must include an e2e test if Web UI behavior changes.
4. Record benchmark results as each benchmarkable task is completed. Benchmarkable means measured product/runtime performance, capacity, package/startup size, graph/DB throughput, graph inventory counts, or rename inventory counts; build/test/e2e timings are validation evidence unless the slice changes those systems.
5. Record evidence as each evidenced task is completed.
6. Before graph-based work, refresh the graph with `avmatrix analyze --force`.
7. Before editing any function, class, method, exported symbol, API handler, graph builder, resolver, analyzer, or shared contract, run AVmatrix impact analysis and record blast radius.
8. Before every implementation commit, run `avmatrix detect-changes --repo AVmatrix --scope all`.
9. After each completed implementation slice, commit the work, then continue until the full plan is complete.

## Problem

The current product/repository name, `AVmatrix`, conflicts with an external brand presence and a third-party website. The project must be renamed to `Anvien` on GitHub and in this local repository.

This is not a simple text replacement. Current codebase evidence shows the old name is embedded into command names, MCP setup, MCP resource URI schemes, generated AI context, storage paths, environment variables, package metadata, launcher paths, Web UI text, tests, snapshots, and generated distribution artifacts.

## Non-Ownership Constraint

`avmatrix.com` is owned by another party and is outside this project. This plan must not include DNS, redirect, certificate, hosting, analytics, or ownership work for that domain.

The only allowed domain-related work is removing repository references that could imply ownership, affiliation, or official status.

## Hard Rename Rule

No legacy runtime or compatibility alias is allowed.

Final implementation must not keep:

- `avmatrix` as an accepted CLI command name;
- `avmatrix` as an MCP server name in editor config;
- `avmatrix://` as an MCP resource URI scheme;
- `.avmatrix` as the active repo/global storage directory;
- `AVMATRIX_*` as active environment variable names;
- `AVmatrixLauncher.exe`, `avmatrix-server.exe`, or `avmatrix.exe` as active release artifact names;
- `avmatrix` as an npm package/bin name;
- `avmatrix-go` as a Go module/repository path;
- generated `.claude/skills/avmatrix/**` or embedded `avmatrix-*` skill ids as active generated output.

Historical mentions are allowed only inside this rebrand plan/evidence/benchmark set and any release-note sentence that identifies the previous name. They must not create a working alias or dual support path.

## Required Final Names

| Surface | Current evidence | Final required name |
|---|---|---|
| Brand display | `AVmatrix` | `Anvien` |
| GitHub repository | `tamnguyendinh/AVmatrix` / `avmatrix-go` references | `tamnguyendinh/Anvien` or the approved Anvien slug |
| Local folder | `E:\AVmatrix-GO` | `E:\Anvien` unless a different exact path is approved |
| CLI command | `avmatrix` | `anvien` |
| Cobra command name | `internal/version.CommandName = "avmatrix"` | `anvien` |
| Go command entrypoint | `cmd/avmatrix` | `cmd/anvien` |
| Go module path | `github.com/tamnguyendinh/avmatrix-go` | approved Anvien module path |
| npm package | `avmatrix/package.json` name/bin `avmatrix` | `anvien` package/bin |
| Web package | `avmatrix-web` | `anvien-web` if package/folder names are renamed |
| Launcher package | `avmatrix-launcher` | `anvien-launcher` if package/folder names are renamed |
| MCP server name | `setupMCPServerName = "avmatrix"` | `anvien` |
| MCP start command | `avmatrix mcp` | `anvien mcp` |
| MCP resource scheme | `canonicalResourceScheme = "avmatrix"` | `anvien://` |
| MCP resources | `avmatrix://repos`, `avmatrix://setup`, `avmatrix://repo/...` | `anvien://repos`, `anvien://setup`, `anvien://repo/...` |
| Editor config keys | `mcpServers.avmatrix`, `[mcp_servers.avmatrix]`, OpenCode `mcp.avmatrix` | `anvien` keys only |
| Repo storage dir | `.avmatrix` | `.anvien` |
| Global storage dir | `~/.avmatrix` | `~/.anvien` |
| Home env var | `AVMATRIX_HOME` | `ANVIEN_HOME` |
| Analyze env vars | `AVMATRIX_MAX_PROCESSES`, related `AVMATRIX_*` | `ANVIEN_*` |
| Launcher env vars | `AVMATRIX_GO`, `AVMATRIX_LAUNCHER_NO_BROWSER`, etc. | `ANVIEN_*` |
| Launcher protocol | `avmatrix://reset` and registry key `HKCU\Software\Classes\avmatrix` | `anvien://...` and Anvien registry key |
| Launcher executables | `AVmatrixLauncher.exe`, `avmatrix-server.exe`, `avmatrix.exe` | `AnvienLauncher.exe`, `anvien-server.exe`, `anvien.exe` |
| Generated AI context block | `renderAVmatrixBlock` and generated AVmatrix tables | `renderAnvienBlock` and Anvien tables |
| Embedded skills | `internal/aicontext/skills/avmatrix-*.md` | `internal/aicontext/skills/anvien-*.md` |
| Generated skills | `.claude/skills/avmatrix/**` | `.claude/skills/anvien/**` |

## Actual Codebase Surface Evidence

AVmatrix graph refresh for this planning update:

- `avmatrix analyze --force`
- scanned `800` files, parsed `583`, unsupported `217`, failed `0`
- graph `91263` nodes, `124743` relationships after the Phase 6 AI context generator slice validation

AVmatrix query/context identified these owners before documentation rewrite:

- MCP server/resources/prompts/tools: `internal/mcp/server.go`, `internal/mcp/resources.go`, `internal/mcp/prompts.go`, `internal/mcp/tools.go`
- MCP editor setup: `internal/cli/setup_command.go`
- CLI root and subcommands: `internal/cli/command.go`
- command name constant: `internal/version/version.go`
- storage paths/env var: `internal/repo/paths.go`
- generated AI context and embedded skills: `internal/aicontext/aicontext.go`, `internal/aicontext/skills/*.md`
- package metadata/runtime packaging: `avmatrix/package.json`, root `package.json`, package lifecycle commands
- Web UI brand/onboarding: `avmatrix-web/index.html`, `avmatrix-web/src/components/LauncherStartScreen.tsx`, `avmatrix-web/src/components/AnalyzeOnboarding.tsx`
- launcher/protocol/process names: `avmatrix-launcher/src/main.go`, `avmatrix-launcher/server-wrapper/main.go`, `avmatrix-launcher/build.ps1`
- GitHub automation and repo metadata sources: `.github/workflows/*.yml`, `.github/actions/setup-avmatrix/action.yml`, `.github/actions/setup-avmatrix-web/action.yml`, `.github/ISSUE_TEMPLATE/*.yml`, `.github/PULL_REQUEST_TEMPLATE.md`, `.github/release-drafter.yml`, `.github/release.yml`, `.github/scripts/**`
- checked-in agent integration config and plugin surfaces: `.mcp.json`, `.grok/config.toml`, `avmatrix-claude-plugin/.mcp.json`, `avmatrix-claude-plugin/hooks/avmatrix-hook.js`, `avmatrix-claude-plugin/hooks/hooks.json`, `avmatrix-claude-plugin/skills/**`
- generated Web contract filenames and contents: `contracts/web-ui/avmatrix-web-contract.schema.json`, `avmatrix-web/src/generated/avmatrix-contracts.ts`
- tests and snapshots: `internal/mcp/*_test.go`, `internal/version/version_test.go`, `avmatrix-web/e2e/*.spec.ts`, launcher tests, baseline snapshots

Initial exact-reference inventory excluding `node_modules` and this rebrand file set:

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

Initial `.github` reference inventory:

| Pattern | Count |
|---|---:|
| `AVmatrix` | 4 |
| `avmatrix` | 94 |
| `AVMATRIX` | 7 |
| `AVmatrix-GO` | 0 |
| `setup-avmatrix` | 6 |
| `github.com/tamnguyendinh/AVmatrix` | 0 |

Additional checked-in integration/contract inventory found during plan review:

| Area | Old-name matches | Files |
|---|---:|---:|
| `.mcp.json` | 2 | 1 |
| `.grok` | 21 | 1 |
| `avmatrix-claude-plugin` | 113 | 16 |
| `contracts` | 4 | 1 |
| `avmatrix-web/src/generated` | 1 | 1 |

Generator/source-of-truth inventory requiring explicit audit before output edits:

| Generator area | Source files | Generated or served output to verify |
|---|---|---|
| AI context and embedded skills | `internal/aicontext/aicontext.go`, `internal/aicontext/skills/*.md` | `AGENTS.md`, `CLAUDE.md`, `.claude/skills/anvien/**`, editor-installed skills |
| Setup/editor config | `internal/cli/setup_command.go` | Cursor/Claude/OpenCode/Codex MCP config, Codex TOML, Claude hooks, installed skills |
| Repo/global storage | `internal/repo/paths.go`, `internal/repo/meta.go`, `internal/repo/settings.go`, `internal/repo/registry.go`, `internal/repo/runtime_config.go` | `<repo>/.anvien/{lbug,graph.json,meta.json,settings.json,analyze.lock,analyze.tmp}`, `~/.anvien/{registry.json,runtime.json}` |
| Group registry storage | `internal/group/storage.go`, `internal/cli/group_command.go`, `internal/mcp/group_tools.go` | `~/.anvien/groups/<name>/group.yaml`, `~/.anvien/groups/<name>/contracts.json`, group command/help output |
| MCP served setup and prompts | `internal/mcp/resources.go`, `internal/mcp/prompts.go`, `internal/mcp/tools.go` | `anvien://setup`, next-step hints, prompt templates, tool descriptions |
| Web contract generation | `internal/contracts/web_ui.go`, `cmd/generate-web-contracts/main.go` | Web contract schema and TypeScript generated adapter filenames/imports |
| Launcher and package outputs | `avmatrix-launcher/build.ps1`, `avmatrix-launcher/src/main.go`, `avmatrix-launcher/server-wrapper/main.go`, `internal/cli/package_runtime.go` | `web-dist`, launcher state/log files, packaged binaries, native runtime copies |
| Hook/plugin integration | `internal/cli/hook_command.go`, `avmatrix-claude-plugin/**` | hook config/status text, plugin skill dirs, plugin `mcp.json`, package install output |
| Graph-quality helper defaults | `cmd/graph-accuracy-probe/main.go`, `internal/cli/resolution_inventory_command.go`, `internal/cli/source_site_accuracy_command.go` | default graph path examples and generated reports that read `.anvien/graph.json` |

Top old-name file groups by file count:

| Group | Files containing old name |
|---|---:|
| `internal` | 338 |
| `reports` | 70 |
| `avmatrix-web` | 68 |
| `docs` | 55 |
| `baseline` | 19 |
| `avmatrix-launcher` | 6 |
| `avmatrix` | 5 |
| `cmd` | 5 |

## Acceptance Criteria

- The active GitHub repository, local workspace naming, package metadata, Web UI, launcher, CLI, MCP server name, MCP resource URI scheme, generated AI context, config paths, and environment variables use `Anvien`/`anvien`.
- No active runtime path accepts or generates the old CLI name, MCP server name, MCP URI scheme, storage path, or environment variable prefix.
- All active tests, snapshots, docs, package files, generated artifacts, and e2e checks expect Anvien names.
- Any remaining old-name reference is confined to this plan/evidence/benchmark set, release-note historical wording, or unavoidable git history.
- `avmatrix.com` remains treated as a third-party website and no plan step attempts to control it.
- AVmatrix impact analysis is recorded before every implementation edit to functions/classes/methods/exported symbols/API handlers/shared contracts.
- Full build runs before tests.
- Web UI brand changes include e2e coverage.
- MCP rename changes include MCP protocol tests for tools/list, resources/list, resources/read, prompts/list, setup output, and editor config generation.
- `avmatrix detect-changes --repo AVmatrix --scope all` runs before every implementation commit until the tool itself is renamed; after the rename slice, use the new equivalent command and record the switch.

## Phase 0 - Baseline And Decisions

- [x] [P0-A] Create the rebrand plan file set under `docs/plans`.
- [x] [P0-B] Refresh the graph with `avmatrix analyze --force` before codebase-dependent planning.
- [x] [P0-C] Record initial AVmatrix query/context evidence for MCP, CLI, generated context, storage, Web UI, package, and launcher surfaces.
- [x] [P0-D] Record initial old-name reference counts.
- [x] [P0-E] Record that `avmatrix.com` is third-party and not controlled by this project.
- [ ] [P0-F] Confirm final GitHub owner/repo slug and local folder path.
- [ ] [P0-G] Confirm whether folders are renamed in-place in the same PR or through a filesystem step before commit.

## Phase 1 - Full Inventory And Edit Map

- [x] [P1-A] Build a full file inventory for every old-name reference, grouped by active source, test, generated artifact, docs, baseline, package output, report, GitHub automation, and temporary output.
- [x] [P1-B] Classify each file as rename-in-place, regenerate, delete stale output, or preserve only as rebrand evidence.
- [x] [P1-C] Identify all generated outputs and their source generators. Do not edit generated `AGENTS.md`, `CLAUDE.md`, `.claude/skills/**`, `avmatrix-launcher/web-dist/**`, or generated Web contracts as permanent source; update the generator first, then regenerate.
- [ ] [P1-D] Run AVmatrix impact analysis for `NewRootCommand`, `newMCPCommand`, `runSetup`, `setupWriteMCPJSON`, `setupWriteOpenCodeJSON`, `setupRunCodexMCPAdd`, `setupUpsertCodexToml`, `setupMergeClaudeHookSettings`, `GenerateAIContextFiles`, `renderAVmatrixBlock`, `repo.Paths`, `repo.GlobalDir`, MCP resource/prompt handlers, launcher startup/reset/cleanup functions, and every other edited symbol found during inventory.
- [ ] [P1-E] Record blast radius and HIGH/CRITICAL warnings before code edits.
- [x] [P1-F] Inventory file and directory names, not only file contents. Rename or remove checked-in paths containing old names, including package folders, command folders, generated contract filenames, plugin folders, action folders, and executable artifacts.
- [x] [P1-G] Classify local-only generated/cache/temp paths such as `.avmatrix`, `.tmp`, `.codex-tmp`, and `.history` separately from tracked files. Do not carry old local cache names into release artifacts.
- [x] [P1-H] Build a generator-to-output matrix for every file-writing or served-content generator listed above. Each row must name the source, generated output, regeneration command, and old-name validation command.
- [ ] [P1-I] Treat a generator still emitting `AVmatrix`, `avmatrix`, `AVMATRIX`, `.avmatrix`, `avmatrix://`, or `avmatrix-*` as a blocker for completing its implementation slice, even if the checked-in output was manually edited.

## Phase 1.5 - Generator Source Audit

- [x] [P1.5-A] Audit `internal/aicontext/aicontext.go` and `internal/aicontext/skills/*.md` as the source of truth for generated `AGENTS.md`, `CLAUDE.md`, `.claude/skills/**`, and editor-installed skills.
- [x] [P1.5-B] Audit `internal/cli/setup_command.go` as the source of truth for generated editor MCP configs, Codex TOML, Claude hooks, and skill installation paths.
- [x] [P1.5-C] Audit repo/global storage generators so `.anvien` output covers `lbug`, `graph.json`, `meta.json`, `settings.json`, `analyze.lock`, `analyze.tmp`, `registry.json`, `runtime.json`, and `groups/**`.
- [x] [P1.5-D] Audit MCP served-content generators so `anvien://setup`, resources, prompts, tool descriptions, and next-step hints cannot regenerate old names.
- [x] [P1.5-E] Audit Web contract and launcher/package generators before touching their generated outputs.
- [x] [P1.5-F] After each generator edit, run the generator or owning command and search the regenerated outputs for old names. Record results immediately in evidence and benchmark ledgers. Satisfied for the AI context generator slice; repeat the same validation inside later generator-specific phases.

## Phase 2 - GitHub Repository Rename Execution

This phase covers work that must happen on GitHub itself, not only in local source files.

- [ ] [P2-A] Confirm the final GitHub owner and repository slug are available and approved. Record the exact final URL in the evidence ledger.
- [ ] [P2-B] Put a short implementation freeze on release/publish activity before the GitHub rename so tags, package publishes, and workflow runs do not race the rename.
- [ ] [P2-C] Rename the GitHub repository in repository Settings from the old AVmatrix slug to the approved Anvien slug.
- [ ] [P2-D] Update the GitHub repository display metadata: description, website/homepage field if any, topics, social preview, and pinned repository references.
- [ ] [P2-E] Update GitHub branch protection and rulesets if any rule names, status-check names, path filters, or required checks include old `avmatrix` naming.
- [ ] [P2-F] Audit and update GitHub Actions repository secrets, variables, and environments whose names or values contain `AVMATRIX`, `avmatrix`, old package names, old image names, old command paths, or old release artifact names.
- [ ] [P2-G] Audit and update GitHub webhooks, deploy keys, GitHub Apps, Pages settings, environments, package permissions, and repository integrations that point at old repo URLs or old package/image names.
- [ ] [P2-H] Update GitHub Releases and release-drafter configuration so generated release titles, notes, assets, and links use Anvien names.
- [ ] [P2-I] Update GitHub issue templates, PR template, labels, CODEOWNERS if present, funding metadata, and support/security contact text that still says AVmatrix.
- [ ] [P2-J] Update `.github/actions/setup-avmatrix` to an Anvien action directory/name and update all workflow references to it.
- [ ] [P2-K] Update `.github/actions/setup-avmatrix-web` to an Anvien action directory/name and update all workflow references to it.
- [ ] [P2-L] Update `.github/workflows/**` path filters, working directories, cache dependency paths, artifact paths, test commands, build commands, workflow comments, and workflow env markers from old names to Anvien names.
- [ ] [P2-M] Update publish workflows for npm package name, package directory, tarball names, release asset names, and authentication scopes.
- [ ] [P2-N] Update Docker/GHCR workflows from `avmatrix`/`avmatrix-web` image slugs to Anvien image slugs.
- [ ] [P2-O] Update GitHub automation scripts under `.github/scripts/**` that reference the old package directory or old constants.
- [ ] [P2-P] Update badges in README/docs to use the new repository slug and workflow names.
- [ ] [P2-Q] Update local `origin` remote to the new GitHub URL and record `git remote -v`.
- [ ] [P2-R] Verify a fresh clone from the new GitHub URL works.
- [ ] [P2-S] Verify old GitHub URL behavior only as GitHub redirect evidence. Do not rely on the redirect as a supported runtime or documentation path.
- [ ] [P2-T] After GitHub and `.github` updates, run the affected workflow-equivalent local commands where possible and record the validation evidence.

## Phase 2.5 - Module, Package, And Folder Names

- [ ] [P2.5-A] Rename Go module path from the old `avmatrix-go` path to the approved Anvien path.
- [x] [P2.5-B] Rename `cmd/avmatrix` to `cmd/anvien` and update all build scripts.
- [ ] [P2.5-C] Rename package folder `avmatrix` to `anvien` if npm package layout is kept.
- [ ] [P2.5-D] Rename `avmatrix-web` to `anvien-web` or record an approved exception only if folder renaming is deferred for filesystem reasons, not for runtime compatibility.
- [ ] [P2.5-E] Rename `avmatrix-launcher` to `anvien-launcher` or record an approved exception only if folder renaming is deferred for filesystem reasons, not for runtime compatibility.
- [ ] [P2.5-F] Update imports, workspace scripts, Docker files, compose files, deploy scripts, CI, release automation, package-lock, and docs links.

## Phase 3 - CLI Hard Rename

- [x] [P3-A] Change `internal/version.CommandName` from `avmatrix` to `anvien`.
- [x] [P3-B] Update root command short/help/version text to `Anvien`.
- [x] [P3-C] Update every command example from `avmatrix ...` to `anvien ...`.
- [x] [P3-D] Update package `bin` mapping to expose only `anvien`.
- [x] [P3-E] Remove old `avmatrix` binary output path from build/package/runtime scripts.
- [x] [P3-F] Update CLI tests and snapshots so the old command name is not accepted or expected.
- [x] [P3-G] Validate `anvien --help`, `anvien version`, `anvien analyze --help`, `anvien mcp`, `anvien serve --help`, `anvien setup --help`, and package lifecycle commands.

Interim Phase 3 validation completed the command entrypoint, Cobra command name, package bin mapping, package runtime metadata, build/package scripts, help/version smoke, MCP initialize smoke through `anvien mcp`, and package `npm run build`/`package ensure-runtime`. Folder names such as `avmatrix`, `avmatrix-web`, and `avmatrix-launcher` remain open under Phase 2.5 and later slices, not as compatibility aliases.

## Phase 4 - MCP Hard Rename

- [x] [P4-A] Change setup constants in `internal/cli/setup_command.go`: brand `Anvien`, command `anvien`, MCP server name `anvien`.
- [x] [P4-B] Update Cursor/Claude JSON setup generation from `mcpServers.avmatrix` to `mcpServers.anvien`.
- [x] [P4-C] Update OpenCode setup generation from `mcp.avmatrix` to `mcp.anvien`.
- [x] [P4-D] Update Codex setup command from `codex mcp add avmatrix -- avmatrix mcp` to `codex mcp add anvien -- anvien mcp`.
- [x] [P4-E] Update Codex TOML output from `[mcp_servers.avmatrix]` to `[mcp_servers.anvien]`.
- [x] [P4-F] Remove old hook matching that preserves `avmatrix-hook` or `avmatrix hook claude`; replace with Anvien-only hook cleanup/generation.
- [x] [P4-G] Change MCP `canonicalResourceScheme` from `avmatrix` to `anvien`.
- [x] [P4-H] Update all MCP next-step hints, setup resources, prompt text, tests, and snapshots from `avmatrix://...` to `anvien://...`.
- [x] [P4-I] Validate MCP stdio and HTTP bridge: `tools/list`, `resources/list`, `resources/read` for `anvien://repos` and `anvien://setup`, `prompts/list`, and representative tool calls.
- [x] [P4-J] Update checked-in root `.mcp.json` from server key/command `avmatrix` to `anvien`.
- [ ] [P4-K] Update `.grok/config.toml` from `[mcp_servers.avmatrix]`, `cmd/avmatrix`, `avmatrix-stable`, and old examples to Anvien-only names.
- [ ] [P4-L] Update every checked-in `mcp.json` under plugin or skill directories so no generated or packaged MCP config uses the old server name.

Interim Phase 4 MCP served-resource validation completed `canonicalResourceScheme`, served resource URIs, prompt text, next-step hints, `serverInfo.name`, tests, MCP baseline testdata, and root `.mcp.json`. P4-K stays open because `.grok/config.toml` is local-untracked in this workspace, not a checked-in commit target. P4-L stays open because plugin/package MCP configs remain under the later plugin/package slice.

## Phase 5 - Storage, Registry, Env Vars, And Local Data

- [x] [P5-A] Rename repo/global storage constants from `.avmatrix` to `.anvien`.
- [x] [P5-B] Rename `AVMATRIX_HOME` to `ANVIEN_HOME` with no fallback.
- [ ] [P5-C] Rename all `AVMATRIX_*` variables to `ANVIEN_*` with no fallback.
- [ ] [P5-D] Update settings examples, lock paths, graph paths, registry docs, group storage docs, and tests.
- [ ] [P5-E] Decide whether existing local `.avmatrix` data is moved once by release instructions or discarded/rebuilt. Do not implement dual-read support.
- [x] [P5-F] Validate analyze creates `<repo>/.anvien/graph.json` and global registry under `~/.anvien/registry.json`.
- [ ] [P5-G] Validate the full generated storage shape: `<repo>/.anvien/lbug`, optional `lbug.wal`/`lbug.lock`, `graph.json`, `meta.json`, `settings.json`, `analyze.lock`, `analyze.tmp`, plus `~/.anvien/runtime.json` and `~/.anvien/groups/**` where relevant.
- [ ] [P5-H] Validate `analyze`, `status`, `index`, `clean`, `doctor locks`, `doctor processes`, `serve`, MCP resources/tools, graph-health, query-health, `resolution-inventory`, `source-site-accuracy`, and graph-accuracy helpers do not recreate or default to `.avmatrix`.

Interim Phase 5 validation completed the core repo/global storage slice: `repo.Paths`, `repo.GlobalDir`, `ANVIEN_HOME`, hook stale-index checks, status/index/graph-health/query-health/source-site/resolution-inventory/httpapi graph guidance, and `.gitignore` now use `.anvien`. P5-C, P5-D, P5-G, and P5-H stay open because broader `AVMATRIX_*` launcher/package vars, group/runtime storage, and full command matrix validation remain in later slices.

Phase 3/7 launcher package work additionally renamed package/launcher env vars already touched by the active build/runtime flow: `ANVIEN_GO`, `ANVIEN_LADYBUGDB_VERSION`, and `ANVIEN_LAUNCHER_NO_BROWSER`. P5-C remains open until every remaining `AVMATRIX_*` source, test, docs, plugin, workflow, and external-helper variable is audited.

## Phase 6 - Generated AI Context And Skills

- [x] [P6-A] Rename `renderAVmatrixBlock` and generated section names to Anvien equivalents.
- [x] [P6-B] Rename embedded skill files from `avmatrix-*.md` to `anvien-*.md`.
- [x] [P6-C] Rename generated skill output from `.claude/skills/avmatrix/**` to `.claude/skills/anvien/**`.
- [x] [P6-D] Update skill descriptions, command tables, MCP resource references, and setup instructions to Anvien names.
- [x] [P6-E] Regenerate `AGENTS.md`, `CLAUDE.md`, and generated skills from source.
- [x] [P6-F] Update generated-context tests to assert Anvien names and absence of active old names.
- [x] [P6-G] Replace generated markers from `<!-- avmatrix:start -->` / `<!-- avmatrix:end -->` to Anvien markers, remove old generated blocks during the rename slice, and do not leave steady-state old-marker compatibility.
- [ ] [P6-H] Verify a fresh `anvien analyze --force` regenerates `AGENTS.md`, `CLAUDE.md`, and `.claude/skills/anvien/**` without recreating `.claude/skills/avmatrix/**` or `avmatrix-*` skill ids.

Interim Phase 6 validation was run with the current local binary path `.\avmatrix\bin\avmatrix.exe analyze --force` because the CLI/module/package rename slices are not complete yet. Final P6-H stays open until the command itself is `anvien`.

## Phase 6.5 - Claude Plugin And Agent Integration Package

- [ ] [P6.5-A] Decide whether `avmatrix-claude-plugin` remains a supported package. If it remains, hard rename the directory/package to Anvien; if it is obsolete, remove it with evidence.
- [ ] [P6.5-B] Rename `avmatrix-claude-plugin/hooks/avmatrix-hook.js` and all hook function names, status messages, env vars, `.avmatrix` path checks, CLI/npx invocations, and stale-index messages to Anvien.
- [ ] [P6.5-C] Update `avmatrix-claude-plugin/hooks/hooks.json` so hook commands and status messages use Anvien only.
- [ ] [P6.5-D] Rename plugin skill ids/directories from `avmatrix-*` to `anvien-*` and update every `SKILL.md` and `mcp.json` inside the plugin.
- [ ] [P6.5-E] Validate plugin installation or packaging behavior if this plugin is still shipped.

## Phase 7 - Web UI And Launcher

- [ ] [P7-A] Update Web UI package name, HTML title, onboarding, start screen, status messages, README rendering, diagnostics globals, and e2e assertions.
- [x] [P7-B] Update launcher executable names, server wrapper paths, temp state names, heartbeat/closed paths, lifecycle data attributes, process cleanup filters, registry protocol key, and tests.
- [x] [P7-C] Rename launcher protocol from `avmatrix://` to `anvien://`.
- [ ] [P7-D] Rebuild packaged Web distribution from source so `web-dist` does not retain old names.
- [ ] [P7-E] Add/update e2e test coverage for visible `Anvien` branding and launcher start behavior.

Interim Phase 7 launcher validation completed binary/protocol/process/lifecycle rename for the packaged launcher path. P7-A, P7-D, and P7-E remain open because the visible Web UI package/source branding and rebuilt `web-dist` old-name content must be fixed from source before e2e branding checks can be final.

## Phase 8 - Documentation, Baselines, And Reports

- [ ] [P8-A] Update README, ARCHITECTURE, RUNBOOK, TESTING, CONTRIBUTING, CHANGELOG current entries, Docker docs, install instructions, setup docs, and badges.
- [ ] [P8-B] Regenerate or rewrite active baseline contract snapshots that still assert old names.
- [ ] [P8-C] Delete stale generated reports or update them if they are active validation artifacts. For tracked historical reports, record whether old-name filenames are preserved only as historical evidence.
- [ ] [P8-D] Keep old-name wording only in this rebrand plan/evidence/benchmark set and final release note history sentence.
- [ ] [P8-E] Run a final old-name search and record every remaining occurrence with a reason.
- [ ] [P8-F] Rename generated Web contract output paths from `contracts/web-ui/avmatrix-web-contract.schema.json` and `avmatrix-web/src/generated/avmatrix-contracts.ts` to Anvien paths, then update the generator and imports rather than hand-editing generated outputs.
- [ ] [P8-G] Update root `package-lock.json`, per-package lockfiles, and any lockfile package names after package/folder renames.

## Phase 9 - Validation And Commits

- [ ] [P9-A] Run the full build gate before tests.
- [ ] [P9-B] Run backend/CLI/MCP tests covering renamed command, setup, storage, resources, prompts, package, and generated context behavior.
- [ ] [P9-C] Run Web unit tests and e2e tests for visible branding and launcher/onboarding behavior.
- [ ] [P9-D] Run package/install/startup smoke tests for `anvien`, `anvien mcp`, and `anvien serve`.
- [ ] [P9-E] Run final reference inventory and require zero active old runtime names.
- [ ] [P9-F] Run `avmatrix detect-changes --repo AVmatrix --scope all` before commits until the rename slice changes the command; then record and use the Anvien equivalent.
- [ ] [P9-G] Commit each completed implementation slice with evidence and benchmark ledgers updated.

## Benchmark Requirements

Record:

- old-name reference counts before and after each slice;
- active old runtime-name count, target `0`;
- package/binary/launcher artifact sizes after rename;
- CLI startup timing if command/package dispatch changes;
- graph inventory counts after generated context and storage rename;
- Web e2e pass/fail counts for Anvien visible text.

## Notes

This plan intentionally rejects a compatibility-preserving rename. The target is a clean Anvien identity with no active AVmatrix runtime, MCP, package, config, or generated-context surface left behind.
