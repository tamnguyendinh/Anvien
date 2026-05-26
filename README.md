# AVmatrix (version: `1.2.2`)

```md
## **AVmatrix is a code map for AI coding agents.**
```

## Why Use AVmatrix?
```md
> **AVmatrix builds one of the fastest and most accurate code intelligence graphs for large repositories.**
```

When you vibe code with AI, the context window is always limited. The longer a session runs, the easier it is for the agent to lose track of information it already read, trace the same code again, or misunderstand how files, functions, and modules relate to each other.

AVmatrix solves this by building a connected map of your codebase: which files relate to which files, which functions call each other, and which modules belong to each execution flow. This helps AI agents understand the project structure faster, navigate relationships across the codebase, and spend less time rediscovering context.


---

## Important notice: **AVmatrix has no official cryptocurrency, token, or coin. Any token/coin using the AVmatrix name is not affiliated with, endorsed by, or created by this project or its maintainers.**

AVmatrix indexes a local codebase into a knowledge graph, then exposes that graph to AI coding agents, CLI commands, and a local Web UI.

The core product is still local code intelligence:

- `avmatrix analyze` builds a repo-local graph index in `<repo>/.avmatrix/`.
- the graph stores semantic layers such as App Layer, Functional Area, source-site proof metadata, ResolutionGap entities, and Resolution Health summaries.
- `avmatrix mcp` exposes indexed repos to MCP clients such as Claude Code, Codex, Cursor, and OpenCode.
- `avmatrix serve` exposes the same local runtime over HTTP for the browser UI.
- `avmatrix-launcher/` packages the local backend and Web UI for a Windows `AVmatrixLauncher.exe` flow.

No AVmatrix-hosted cloud service is involved in the active local runtime path.

You do not need to put API keys into AVmatrix. Indexing, graph storage, repo switching, and graph queries run locally on your machine. For chat, AVmatrix uses the local Codex or Claude Code session/account you already use on this machine; AVmatrix does not store provider keys or route chat through an AVmatrix cloud service.

---

## Current Runtime Model

| Surface | Purpose | Entry point |
|---------|---------|-------------|
| CLI | Analyze repos, query the graph, inspect impact, manage indexes/groups | `avmatrix ...` |
| MCP stdio | Agent-facing graph tools and resources | `avmatrix mcp` |
| Local HTTP API | Web UI backend, graph streaming, analyze jobs, session bridge | `avmatrix serve` |
| Web UI | Browser graph explorer, repo picker/analyze UI, Codex/Claude Code style session chat | `avmatrix-web/` or packaged launcher |
| Windows launcher | Starts packaged Web UI and backend on `127.0.0.1` | `avmatrix-launcher\AVmatrixLauncher.exe` |

The Web UI is a frontend over the local HTTP backend. Repo switching and graph loading use explicit repo-scoped read targets; they do not depend on one mutable process-global active repo.

The Web chat does not run an AI model inside AVmatrix. The shared session contract supports `codex` and `claude-code`; the current backend mounts the Codex CLI adapter. AVmatrix keeps repo binding, streaming, cancellation, and UI state local.

---

  ## How to use AVmatrix

  Requires Node.js 20+, npm, and Go.

  1. Clone or download the AVmatrix repository.

  2. Open Codex CLI or Claude Code in the AVmatrix repository folder.

  3. Paste this prompt:

     ```text
     Install AVmatrix from this repository and configure its MCP integration.
     ```


     Then run:

     ```
     powershell -ExecutionPolicy Bypass -File avmatrix-launcher\build.ps1
     ```

  4. Use avmatrix-launcher\AVmatrixLauncher.exe to open the visual Web UI.
  5. After AVmatrix MCP is configured, your AI agent can use AVmatrix tools
     for codebase analysis, impact checks, graph queries, and navigation.

  The agent should build the Go-backed avmatrix package, install or link the
  local CLI, run avmatrix setup, and verify avmatrix --version.

### Manual install

```bash
git clone https://github.com/tamnguyendinh/AVmatrix.git
cd AVmatrix

cd avmatrix
npm install
npm link

avmatrix --version
```

Index a local repository:

```bash
cd /path/to/your/repo
avmatrix analyze .
```

This creates `<repo>/.avmatrix/` and registers the repo in `~/.avmatrix/registry.json`.

Configure MCP/editor integration:

```bash
avmatrix setup
```

Manual MCP examples:

```bash
claude mcp add avmatrix -- avmatrix mcp
codex mcp add avmatrix -- avmatrix mcp
```

Codex TOML:

```toml
[mcp_servers.avmatrix]
command = "avmatrix"
args = ["mcp"]
```

### Grok (xAI)

This repository provides a **Grok-only** MCP configuration at `.grok/config.toml`.

When you open the AVmatrix-GO folder with Grok, the AVmatrix tools are automatically available (this file has higher priority than `.mcp.json` and does not affect Claude, Cursor, Codex, or other agents).

**For contributors working inside this repo:**

- Start Grok (recommended: `grok --model grok-build --effort high` or `xhigh`)
- The MCP server will be started via `go run ./cmd/avmatrix mcp`
- Verify with `/mcps` or `grok mcp list`

**For other projects or daily use:**

Build once and register with an explicit path:

```bash
go build -o avmatrix-stable.exe ./cmd/avmatrix
grok mcp add avmatrix -- "E:\\path\\to\\avmatrix-stable.exe" mcp
```

You can also create a `.grok/config.toml` in any of your own repositories to enable AVmatrix tools there.

This approach keeps the public MCP contract (used by all other agents) completely unchanged.

---

## Quick Start: Web UI

Development flow:

```bash
# terminal 1, from the repo root
go run ./cmd/avmatrix serve --host 127.0.0.1 --port 4848

# or build the local Go CLI once, then run it
go build -trimpath -o .tmp/avmatrix.exe ./cmd/avmatrix
.\.tmp\avmatrix.exe serve --host 127.0.0.1 --port 4848

# terminal 2, from the repo root
cd avmatrix-web
npm install
npm run dev
```

Open:

```text
http://127.0.0.1:5228
```

The browser connects to:

```text
http://127.0.0.1:4848
```

From the Web UI you can:

- choose an indexed local repo
- analyze another local repo
- remove a repo from the landing list
- switch repos from the header dropdown
- browse graph nodes, links, files, processes, and search results
- use the local session bridge for Codex/Claude Code style chat

---

## Packaged Windows Launcher

The packaged launcher is a convenience layer around the same local backend and Web UI.

Build it:

```powershell
powershell -ExecutionPolicy Bypass -File avmatrix-launcher\build.ps1
```

Important artifacts:

```text
avmatrix-launcher\AVmatrixLauncher.exe
avmatrix\bin\avmatrix.exe
avmatrix-launcher\server-bundle\avmatrix-server.exe
avmatrix-launcher\web-dist\
```

Runtime behavior:

- `AVmatrixLauncher.exe` is rebuilt by `avmatrix-launcher\build.ps1` and is the packaged user entrypoint.
- `avmatrix\bin\avmatrix.exe` is the single production AVmatrix CLI/runtime executable built by the full build.
- `AVmatrixLauncher.exe` serves the packaged Web UI on `127.0.0.1:5228` and opens the in-app start screen.
- `avmatrix-server.exe` starts `avmatrix\bin\avmatrix.exe serve`.
- backend health is checked at `http://127.0.0.1:4848/api/info`.
- reset/stop use the launcher state file plus process path sweep for the packaged runtime.

The launcher must remain optional. `avmatrix serve` is still the direct backend entry point.

---

## Main CLI Commands

```bash
avmatrix setup                     # Configure local MCP/editor access
avmatrix analyze [path]            # Full local repo analysis
avmatrix analyze --force           # Force full re-index
avmatrix analyze --embeddings      # Generate semantic embeddings
avmatrix analyze --no-stats        # Omit volatile stats from generated agent files
avmatrix analyze --skip-git        # Analyze a folder without requiring .git
avmatrix analyze --name <alias>    # Register repo under a custom name
avmatrix index [path...]           # Register an existing local index
avmatrix list                      # List indexed repos
avmatrix status                    # Show index status for current repo
avmatrix clean                     # Delete current repo index
avmatrix clean --all --force       # Delete all indexes
avmatrix mcp                       # Start MCP server over stdio
avmatrix serve                     # Start local HTTP backend on 127.0.0.1:4848
avmatrix doctor                    # Inspect local runtime locks and processes
avmatrix version                   # Print version/build information
avmatrix wiki                      # Show wiki capability status
avmatrix wiki-mode [off|local]     # Show or set local wiki capability mode
avmatrix completion <shell>        # Generate shell completion script
```

Direct graph tools:

```bash
avmatrix query <search_query>
avmatrix context [name]
avmatrix impact [target]
avmatrix rename <symbol> <newName>
avmatrix cypher <query>
avmatrix detect-changes
avmatrix augment <pattern>
avmatrix api route-map [route]
avmatrix api tool-map [tool]
avmatrix api shape-check [route]
avmatrix api impact [route]
avmatrix graph-health
avmatrix query-health
avmatrix resolution-inventory
avmatrix source-site-accuracy
avmatrix benchmark-compare <before> <after>
```

AI context and skills:

`avmatrix analyze` refreshes managed `AGENTS.md`, `CLAUDE.md`, and `.claude/skills/avmatrix/**` from embedded source files under `internal/aicontext/skills/*.md`. `avmatrix setup` installs the same embedded base skill set into supported editor skill directories. Do not edit generated root context or `.claude/skills/avmatrix/**` as source; change the embedded skill source or generator and regenerate through analyze.

Semantic graph diagnostics:

```bash
avmatrix graph-health summary --repo <repo> --json
avmatrix graph-health report --repo <repo> --limit 20 --json
avmatrix graph-health components --repo <repo> --json
avmatrix query-health --repo <repo> --out .tmp/query-health.json
avmatrix resolution-inventory --graph .avmatrix/graph.json --out .tmp/resolution-inventory.json
avmatrix source-site-accuracy --graph .avmatrix/graph.json --out .tmp/source-site-accuracy.json
```

These commands are for checking graph quality, not for replacing `analyze`. `analyze` remains the source of truth that refreshes the graph. `graph-health` audits computed topology health, diagnostics, component membership, confidence, resolution-health overlays, and prioritized candidate reports from the indexed repo graph. `query-health` measures query retrieval with two separate outcomes: threshold pass/fail for usable retrieval, and exact pass/fail for complete expected target coverage. Use `--fail-on-threshold` to fail when hit@5/hit@10 thresholds are missed, or `--fail-on-exact` to fail when any expected file/symbol is still missing. `resolution-inventory` reports persisted ResolutionGap and Resolution Health counts, including non-actionable breakdowns such as `builtin`, `standard_library`, and `test_framework`. `source-site-accuracy` reports proof-based CALLS/ACCESSES inventory, missing source-site IDs, false resolved edge candidates, and other graph accuracy gates.

`avmatrix rename` and `avmatrix api ...` are CLI equivalents for the MCP `rename`, `route_map`, `tool_map`, `shape_check`, and `api_impact` tools. Use them for terminal workflows and smoke validation; they delegate to the same local MCP tool logic so API/rename semantics stay consistent across command surfaces.

Repository groups:

```bash
avmatrix group create <name>
avmatrix group add <group> <groupPath> <registryName>
avmatrix group remove <group> <path>
avmatrix group list [name]
avmatrix group sync <name>
avmatrix group contracts <name>
avmatrix group query <name> <query>
avmatrix group status <name>
```

Repo-local settings live in `.avmatrix/settings.json`. The current repo-local setting is `maxExecutionFlows`, which controls the cap used while materializing execution flows during `analyze`. `AVMATRIX_MAX_PROCESSES` remains available as a temporary override.

---

## MCP Tools And Resources

AVmatrix exposes 16 MCP tools:

| Tool | Purpose |
|------|---------|
| `list_repos` | Discover indexed repos |
| `query` | Hybrid search over execution flows and symbols, with semantic App Layer, Functional Area, and Resolution Health fields when available |
| `cypher` | Raw Cypher against the graph |
| `context` | 360-degree symbol context, source-site proof/status metadata, and related ResolutionGap rows |
| `detect_changes` | Map git diffs to affected symbols/processes, changed App Layers, ResolutionGap changes, and resolution-health impact |
| `rename` | Graph-assisted multi-file rename preview/application |
| `impact` | Upstream/downstream blast radius with affected App Layers, Functional Areas, and resolution-health risks |
| `route_map` | API route to handler/consumer mapping |
| `tool_map` | MCP/RPC tool definition and handler mapping |
| `shape_check` | API response shape vs consumer access checks |
| `api_impact` | API route pre-change impact report |
| `group_list` | List repo groups |
| `group_sync` | Build cross-repo contract registry |
| `group_contracts` | Inspect group contracts and cross-links |
| `group_query` | Search across a repo group |
| `group_status` | Check group/repo staleness |

Common resources:

| Resource | Purpose |
|----------|---------|
| `avmatrix://repos` | All indexed repos |
| `avmatrix://setup` | Generated setup/onboarding content |
| `avmatrix://repo/{name}/context` | Repo overview and stats |
| `avmatrix://repo/{name}/clusters` | Functional clusters |
| `avmatrix://repo/{name}/cluster/{name}` | Cluster detail |
| `avmatrix://repo/{name}/processes` | Execution flows |
| `avmatrix://repo/{name}/process/{name}` | Process trace |
| `avmatrix://repo/{name}/schema` | Graph schema |

MCP prompts:

| Prompt | Purpose |
|--------|---------|
| `detect_impact` | Agent template for pre-commit impact analysis with `detect_changes`, `context`, `impact`, freshness checks, and HIGH/CRITICAL blast-radius interpretation |
| `generate_map` | Agent template for evidence-backed architecture documentation from `avmatrix://repos`, repo context, clusters, processes, selected process details, and any extra tools/commands the agent actually reads |

MCP prompts are workflow templates for MCP-capable agents, not CLI commands. `generate_map` must resolve an exact repo before reading repo resources, URL-escape repo and process names in resource URIs, refresh stale graph evidence with `avmatrix analyze --force` when required, and avoid architecture claims or Mermaid edges that are not backed by graph evidence the agent actually read.

When only one repo is indexed, most repo-scoped tool calls can omit `repo`. With multiple indexed repos, pass the repo name or path explicitly.

---

## How Indexing Works

`avmatrix analyze` runs a full local pipeline:

```text
scan -> structure -> [markdown, cobol] -> parse -> [routes, tools, orm]
  -> crossFile -> mro -> communities -> processes
  -> semantic enrichment -> LadybugDB load -> FTS
  -> optional embeddings -> metadata/registry/agent files
```

The graph is stored in LadybugDB under `<repo>/.avmatrix/`.

Semantic enrichment adds user-facing graph meaning on top of raw code symbols:

- **App Layer**: backend, frontend, API, shared contract, docs, tests, config, generated contract, mixed, or unknown.
- **Functional Area**: high-confidence ownership such as resolution, graph health, query, MCP, Web graph UI, layout, contracts, providers, runtime, analyzer, session, launcher, CLI, storage, or unknown.
- **Source-site proof**: resolved relationships keep source-site IDs, proof kind, target role, target text, file/range, confidence, and resolution source.
- **ResolutionGap**: unresolved, external, ambiguous, unsupported, or non-actionable references are persisted as diagnostic graph entities instead of being silently dropped or converted into fake resolved edges.
- **Resolution Health**: graph readers can separate resolved references, in-repo analyzer gaps, external unresolved references, non-actionable builtins/standard-library/test-framework references, and unclassified unknowns.

In the Web UI, ResolutionGap entities are diagnostic nodes rather than real code symbols. They are rendered as small square nodes and can be filtered or grouped separately from normal symbol nodes.

Storage:

```text
<repo>/.avmatrix/
  lbug
  lbug.wal
  lbug.lock
  meta.json
  settings.json

~/.avmatrix/
  registry.json
```

Supported language detection currently covers:

```text
JavaScript, TypeScript, Python, Java, C, C++, C#, Go, Ruby, Rust,
PHP, Kotlin, Swift, Dart, Vue, COBOL
```

COBOL/JCL is handled through the dedicated COBOL phase rather than the normal tree-sitter worker path.

---

## Local HTTP API

`avmatrix serve` exposes the local backend used by the Web UI:

| Endpoint | Purpose |
|----------|---------|
| `/api/info` | Finite backend liveness/readiness |
| `/api/heartbeat` | Long-lived SSE heartbeat stream |
| `/api/repos`, `/api/repo` | List/select/remove indexed repos |
| `/api/graph` | Repo-scoped graph load/stream |
| `/api/query`, `/api/search`, `/api/file`, `/api/grep` | Repo-scoped read/search helpers |
| `/api/process*`, `/api/cluster*` | Derived graph views |
| `/api/local/folder-picker` | Native local folder picker bridge |
| `/api/analyze`, `/api/embed` | Background analyze/embed jobs |
| `/api/mcp` | MCP-over-HTTP bridge |
| `/api/session/*` | Session bridge for chat runtime |

The graph loading path uses:

```text
repo-resolver -> repo-read-executor -> graph-read-service -> graph-stream-http
```

That is the replacement for the older process-global DB retargeting path used during Web repo switching.

---

## Docker

Docker support remains available as an advanced deployment path. It is separate from the primary local CLI/launcher flow.

Files:

- `Dockerfile.cli`
- `Dockerfile.web`
- `docker-compose.yaml`
- `.env.example`

Compose expects exact image tags:

```bash
cp .env.example .env
# set SERVER_IMAGE and WEB_IMAGE
docker compose --env-file .env up -d
```

Default ports:

```text
server: http://127.0.0.1:4848
web:    http://127.0.0.1:4173
```

To make host repos visible to the container, set `WORKSPACE_DIR` to a local folder that contains the repos you want to analyze. It is mounted read-only at `/workspace`.

---

## Repository Layout

| Path | Role |
|------|------|
| `cmd/`, `internal/` | Go CLI, MCP server, HTTP API, ingestion, LadybugDB, embeddings, contracts, session/runtime code |
| `avmatrix/` | npm packaging and Go runtime distribution glue |
| `avmatrix-web/` | React/Vite Web UI and local runtime client |
| `contracts/web-ui/` | Go-generated Web UI contract manifest |
| `avmatrix-launcher/` | Windows launcher, server wrapper, packaged Web UI/backend assets |
| `.claude/`, `avmatrix-claude-plugin/`, `avmatrix-cursor-integration/` | Generated agent context output and plugin metadata |
| `docs/plans/` | Implementation plans and investigation records |
| `.github/` | CI workflows |

See [ARCHITECTURE.md](ARCHITECTURE.md) for the detailed system map.

---

## Development

Build core packages:

```bash
cd avmatrix
npm install
npm run build
```

Build Web UI:

```bash
go run ./cmd/generate-web-contracts --check

cd avmatrix-web
npm install
npm run build
```

Build full Windows launcher package:

```powershell
powershell -ExecutionPolicy Bypass -File avmatrix-launcher\build.ps1
```

Useful docs:

- [ARCHITECTURE.md](ARCHITECTURE.md)
- [CHANGELOG.md](CHANGELOG.md)
- [RUNBOOK.md](RUNBOOK.md)
- [GUARDRAILS.md](GUARDRAILS.md)
- [CONTRIBUTING.md](CONTRIBUTING.md)
- [TESTING.md](TESTING.md)

---

## Security And Privacy

- Index data is stored locally in `<repo>/.avmatrix/`.
- The global registry is local under `~/.avmatrix/`.
- The Web UI talks to the local backend at `127.0.0.1:4848`.
- AVmatrix does not store AI provider API keys in the browser.
- AVmatrix does not route chat through an AVmatrix cloud service.
- Codex/Claude Code style chat depends on the local session/provider already available on the machine.

---

## Acknowledgments

- [Model Context Protocol](https://modelcontextprotocol.io/)
- [Tree-sitter](https://tree-sitter.github.io/)
- [LadybugDB](https://ladybugdb.com/)
- [Sigma.js](https://www.sigmajs.org/)
- [Graphology](https://graphology.github.io/)
- [Transformers.js](https://huggingface.co/docs/transformers.js)
