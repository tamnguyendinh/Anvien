# Runbook - Anvien

Copy-paste operations for the current local Anvien runtime: CLI, MCP, Web UI, packaged launcher, repo indexes, and recovery.

Current CLI package version: `1.2.3` (`anvien/package.json`).

---

## Prerequisites

Required for normal local development:

- Node.js 20+
- npm
- Git
- Go

Required for packaged Windows launcher work:

- PowerShell

Build and link the local CLI from source:

```powershell
cd anvien
npm install
npm run build
npm link

anvien --version
```

If you are using Codex CLI or Claude Code inside this repository, you can also ask the agent:

```text
Install Anvien from this repository and configure its MCP integration.
```

---

## Index A Repo

From the target repo:

```powershell
anvien analyze .
```

Force a full rebuild:

```powershell
anvien analyze . --force
```

Analyze a folder without requiring `.git`:

```powershell
anvien analyze C:\path\to\repo --skip-git
```

Register the repo under a custom name:

```powershell
anvien analyze C:\path\to\repo --name MyRepo
```

Check what Anvien knows:

```powershell
anvien status
anvien list
```

Index data is stored in the repo:

```text
<repo>\.anvien\
  lbug
  lbug.wal
  lbug.lock
  graph.json
  meta.json
  settings.json
```

The global registry is stored at:

```text
~\.anvien\registry.json
```

---

## Stale Index

Symptom:

- MCP/resources warn the index is behind `HEAD`
- graph/query results do not match recent code

Fix:

```powershell
cd C:\path\to\repo
anvien analyze .
```

If the index may be corrupt or the schema/runtime changed:

```powershell
anvien analyze . --force
```

---

## MCP Setup And Recovery

Configure detected editors/agents:

```powershell
anvien setup
```

Manual examples:

```powershell
claude mcp add anvien -- anvien mcp
codex mcp add anvien -- anvien mcp
```

Codex TOML:

```toml
[mcp_servers.anvien]
command = "anvien"
args = ["mcp"]
```

Start MCP manually:

```powershell
anvien mcp
```

If MCP says no repos are indexed:

```powershell
cd C:\path\to\repo
anvien analyze .
```

If multiple repos are indexed, pass `repo` explicitly in MCP tool calls or call `list_repos` first.

---

## Web UI Local Backend

Start the HTTP backend:

```powershell
go run ./cmd/anvien serve --host 127.0.0.1 --port 4848
```

Or through the package script:

```powershell
cd anvien
npm run serve
```

After `npm link`:

```powershell
anvien serve
```

Default backend:

```text
http://127.0.0.1:4848
```

Health checks:

```powershell
Invoke-WebRequest http://127.0.0.1:4848/api/info
Invoke-WebRequest http://127.0.0.1:4848/api/repos
```

Start the Web UI dev server:

```powershell
cd anvien-web
npm install
npm run dev
```

Open:

```text
http://127.0.0.1:5228
```

Important behavior:

- `anvien serve` is loopback-only by default.
- Web graph loading uses explicit repo-scoped read targets.
- The graph switch path is `/api/graph?repo=...&stream=true`.
- The Web UI is a frontend over the local backend, not a separate index/runtime.

---

## Web UI Troubleshooting

### Failed to fetch / backend disconnected

Check backend:

```powershell
Invoke-WebRequest http://127.0.0.1:4848/api/info
```

If the port is down, restart:

```powershell
go run ./cmd/anvien serve --host 127.0.0.1 --port 4848
```

If port `4848` is already in use:

```powershell
anvien serve --port 4748
```

The default Web UI expects `4848`; use the packaged/default path unless you are deliberately testing another backend URL.

### Loading graph hangs

Check the target repo is registered:

```powershell
anvien list
```

Check the graph endpoint directly:

```powershell
Invoke-WebRequest "http://127.0.0.1:4848/api/graph?repo=MyRepo&stream=true"
```

If the response contains LadybugDB/WAL errors, stop active runtimes and rebuild the repo index.

### Newly analyzed repo does not appear

Check the registry response:

```powershell
Invoke-WebRequest http://127.0.0.1:4848/api/repos
```

If the repo exists in CLI but not the Web UI, refresh the browser. If needed, restart `anvien serve`.

### File Map or File Detail is empty

Check the file projection endpoints against an indexed repo:

```powershell
Invoke-WebRequest "http://127.0.0.1:4848/api/file-hotspots?repo=MyRepo&limit=5"
Invoke-WebRequest "http://127.0.0.1:4848/api/file-context?repo=MyRepo&path=src%2Fapp.go"
```

If the responses report stale or missing graph data, rerun:

```powershell
anvien analyze --force
```

File relationship groups are derived projections from symbol and source-site graph facts. The canonical graph remains symbol/source-site based; stale or missing projection data usually means the repo graph needs a refresh or the requested file path is not in the indexed repo.

---

## Packaged Windows Launcher

Build the full packaged local runtime:

```powershell
powershell -ExecutionPolicy Bypass -File anvien-launcher\build.ps1
```

Artifacts:

```text
anvien-launcher\AnvienLauncher.exe
anvien\bin\anvien.exe
anvien-launcher\server-bundle\anvien-server.exe
anvien-launcher\web-dist\
```

`AnvienLauncher.exe` is rebuilt by `anvien-launcher\build.ps1` and is the packaged user entrypoint. `anvien\bin\anvien.exe` is the single production Anvien CLI/runtime executable; the launcher backend wrapper runs that file with `serve`. The start screen is served from the packaged Web UI; there is no separate root HTML launcher file.

Start:

```powershell
.\anvien-launcher\AnvienLauncher.exe
```

Reset runtime:

```powershell
.\anvien-launcher\AnvienLauncher.exe reset
```

Stop runtime:

```powershell
.\anvien-launcher\AnvienLauncher.exe stop
```

Register protocol:

```powershell
.\anvien-launcher\AnvienLauncher.exe register
```

Logs:

```text
anvien-launcher\logs\launcher.log
anvien-launcher\logs\backend.log
anvien-launcher\logs\server-wrapper.log
```

The launcher is optional. It starts the same local backend that `anvien serve` starts directly.

---

## Session Chat Runtime

The Web chat does not run an AI model inside Anvien. It sends chat requests through the local session bridge and currently executes them with the Codex CLI session available on this machine.

Implemented chat provider:

```text
codex -> Codex CLI adapter
```

Reserved provider identifier:

```text
claude-code -> UI/status slot only until a backend adapter is added
```

Check Codex locally:

```powershell
codex --version
codex login status
```

If chat is unavailable:

- verify the local provider CLI is installed
- verify the provider account is signed in
- verify the selected repo is indexed
- check `http://127.0.0.1:4848/api/session/status`

Anvien does not store provider API keys in the browser and does not route chat through an Anvien cloud service.

---

## Embeddings

Generate embeddings during full analyze:

```powershell
anvien analyze . --embeddings
```

If you want to preserve/regenerate embeddings on future rebuilds, keep passing `--embeddings`.

Current behavior:

- embeddings are optional
- embeddings are stored in the `CodeEmbedding` table
- embedding generation is skipped for repos over `100000` nodes
- search still works without embeddings through BM25/FTS paths

Backend embedding endpoint:

```text
POST /api/embed
GET /api/embed/:jobId
GET /api/embed/:jobId/progress
DELETE /api/embed/:jobId
```

Analyze and embed jobs are locked per repo to avoid concurrent writes to the same LadybugDB store.

---

## Generated Agent Context And Skills

`anvien analyze` refreshes generated `AGENTS.md`, `CLAUDE.md`, and `.claude/skills/anvien/**`.

Source of truth:

```text
internal\aicontext\aicontext.go
internal\aicontext\skills\*.md
```

Do not edit generated root context or `.claude\skills\anvien\**` as the permanent source. Change the generator registry or embedded skill source, then regenerate through analyze/setup as appropriate.

The generated Anvien skill set is task-facing only. It must not include `anvien-ai-context` or any generated skill that instructs agents to modify the AI-context generator itself.

---

## Clean And Recover A Repo Index

Current repo:

```powershell
anvien clean --force
```

All indexed repos:

```powershell
anvien clean --all --force
```

Then rebuild:

```powershell
anvien analyze . --force
```

If you see WAL/checksum corruption:

1. Stop the packaged launcher or `anvien serve`.
2. Stop MCP sessions that may have the repo open.
3. Clean the repo index.
4. Re-run `anvien analyze . --force`.

Do not manually delete `.anvien\lbug*` while launcher/backend/MCP is running.

---

## Lock And Concurrent Job Errors

Symptom:

```text
Another job is already active for this repository
```

Meaning:

- analyze, delete, or embed is already active for that repo
- the backend lock is protecting the repo-local LadybugDB store

Fix:

- wait for the current job to finish
- cancel the Web UI job if available
- restart the local backend if the job is stale

---

## CLI Equivalents Of MCP Tools

Useful when debugging without an editor:

```powershell
anvien query "authentication flow" --repo MyRepo
anvien context SomeSymbol --repo MyRepo
anvien impact SomeSymbol --direction upstream --repo MyRepo
anvien cypher "MATCH (n) RETURN count(n) LIMIT 1" --repo MyRepo
anvien detect-changes --repo MyRepo
```

---

## Wiki Capability

The active local-only build has no remote wiki fallback.

```powershell
anvien wiki
anvien wiki-mode
anvien wiki-mode off
anvien wiki-mode local
```

`wiki-mode local` is reserved for a future local wiki engine.

---

## Docker

Docker is an advanced deployment path, separate from the primary local CLI/launcher flow.

Required files:

```text
Dockerfile.cli
Dockerfile.web
docker-compose.yaml
.env.example
```

Run:

```powershell
Copy-Item .env.example .env
# edit SERVER_IMAGE and WEB_IMAGE
docker compose --env-file .env up -d
```

Default ports:

```text
server: http://127.0.0.1:4848
web:    http://127.0.0.1:4173
```

To analyze host repos inside Docker, set `WORKSPACE_DIR` to a local folder that contains those repos. It is mounted read-only at `/workspace`.

---

## Build And Validation Commands

Core:

```powershell
cd anvien
npm run build
npm test
npx tsc --noEmit
```

Web:

```powershell
cd anvien-web
npm run build
npm test
```

Web e2e requires both backend and frontend running:

```powershell
cd anvien
anvien serve

cd ..\anvien-web
npm run dev
npm run test:e2e
```

Launcher package:

```powershell
powershell -ExecutionPolicy Bypass -File anvien-launcher\build.ps1
```

---

## Where To Dig Deeper

- [ARCHITECTURE.md](ARCHITECTURE.md)
- [README.md](README.md)
- [CHANGELOG.md](CHANGELOG.md)
- [GUARDRAILS.md](GUARDRAILS.md)
- [TESTING.md](TESTING.md)
