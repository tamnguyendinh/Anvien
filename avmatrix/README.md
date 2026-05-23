# AVmatrix

**Graph-powered code intelligence for AI agents.** Index any codebase into a knowledge graph, then query it via MCP or CLI.

Works with **Cursor**, **Claude Code**, **Codex**, **Windsurf**, **Cline**, **OpenCode**, and any MCP-compatible tool.

[![npm version](https://img.shields.io/npm/v/avmatrix.svg)](https://www.npmjs.com/package/avmatrix)
[![License: PolyForm Noncommercial](https://img.shields.io/badge/License-PolyForm%20Noncommercial-blue.svg)](https://polyformproject.org/licenses/noncommercial/1.0.0/)

---

## Why?

AI coding tools don't understand your codebase structure. They edit a function without knowing 47 other functions depend on it. AVmatrix fixes this by **precomputing every dependency, call chain, and relationship** into a queryable graph.

**Three commands to give your AI agent full codebase awareness.**

## Quick Start

```bash
# Index your repo (run from repo root)
avmatrix analyze
```

That's it. This indexes the codebase, installs agent skills, registers Claude Code hooks, and creates `AGENTS.md` / `CLAUDE.md` context files — all in one command.

To configure MCP for your editor, run `avmatrix setup` once — or set it up manually below.

`avmatrix setup` auto-detects your editors and writes the correct global MCP config. You only need to run it once.

### Editor Support

| Editor | MCP | Skills | Hooks (auto-augment) | Support |
|--------|-----|--------|---------------------|---------|
| **Claude Code** | Yes | Yes | Yes (PreToolUse) | **Full** |
| **Cursor** | Yes | Yes | — | MCP + Skills |
| **Codex** | Yes | Yes | — | MCP + Skills |
| **Windsurf** | Yes | — | — | MCP |
| **OpenCode** | Yes | Yes | — | MCP + Skills |

> **Claude Code** gets the deepest integration: MCP tools + agent skills + PreToolUse hooks that automatically enrich grep/glob/bash calls with knowledge graph context.

### Community Integrations

| Agent | Install | Source |
|-------|---------|--------|
| [pi](https://pi.dev) | `pi install npm:avmatrix` | [AVmatrix integration](https://github.com/tamnguyendinh/AVmatrix) |

## MCP Setup (manual)

If you prefer to configure manually instead of using `avmatrix setup`:

### Claude Code (full support — MCP + skills + hooks)

```bash
# First install the local CLI onto PATH (for example: `cd avmatrix && npm link`)
claude mcp add avmatrix -- avmatrix mcp
```

### Codex (full support — MCP + skills)

```bash
# First install the local CLI onto PATH (for example: `cd avmatrix && npm link`)
codex mcp add avmatrix -- avmatrix mcp
```

### Cursor / Windsurf

Add to `~/.cursor/mcp.json` (global — works for all projects):

```json
{
  "mcpServers": {
    "avmatrix": {
      "command": "avmatrix",
      "args": ["mcp"]
    }
  }
}
```

### OpenCode

Add to `~/.config/opencode/config.json`:

```json
{
  "mcp": {
    "avmatrix": {
      "command": "avmatrix",
      "args": ["mcp"]
    }
  }
}
```

## How It Works

AVmatrix builds a complete knowledge graph of your codebase through a multi-phase indexing pipeline:

1. **Structure** — Walks the file tree and maps folder/file relationships
2. **Parsing** — Extracts functions, classes, methods, and interfaces using Tree-sitter ASTs
3. **Resolution** — Resolves imports and function calls across files with language-aware logic
   - **Field & Property Type Resolution** — Tracks field types across classes and interfaces for deep chain resolution (e.g., `user.address.city.getName()`)
   - **Return-Type-Aware Variable Binding** — Infers variable types from function return types, enabling accurate call-result binding
4. **Clustering** — Groups related symbols into functional communities
5. **Processes** — Traces execution flows from entry points through call chains
6. **Semantic enrichment** — Adds App Layer, Functional Area, source-site proof metadata, ResolutionGap entities, and Resolution Health summaries
7. **Search** — Builds hybrid search indexes for fast retrieval

The result is a **LadybugDB graph database** stored locally for your repo in `.avmatrix/`, with full-text search and semantic embeddings.

### Semantic Graph Output

AVmatrix now persists graph meaning that users and agents can inspect consistently from CLI, MCP, and the Web UI:

- **App Layer** separates backend, frontend, API, shared contract, docs, tests, config, generated contract, mixed, and unknown graph regions.
- **Functional Area** records high-confidence ownership such as resolution, graph health, query, MCP, Web graph UI, layout, contracts, providers, runtime, analyzer, session, launcher, CLI, and storage.
- **Source-site proof** records why a `CALLS`, `ACCESSES`, type-reference, or heritage edge was accepted, including source-site IDs, proof kind, target role, target text, file/range, confidence, and resolution source.
- **ResolutionGap** persists unresolved or diagnostic references instead of inventing fake resolved in-repo edges. Non-actionable gaps are split into `builtin`, `standard_library`, and `test_framework` groups.
- **Resolution Health** exposes clear, degraded, unknown, in-repo analyzer gap, external unresolved, and non-actionable reference counts.

In the Web UI, ResolutionGap entities are diagnostic graph nodes, not normal code symbols. They render as small square nodes and can be filtered or grouped separately from real symbols.

## MCP Tools

Your AI agent gets these tools automatically:

| Tool | What It Does | `repo` Param |
|------|-------------|--------------|
| `list_repos` | Discover all indexed repositories | — |
| `query` | Process-grouped search with App Layer, Functional Area, Resolution Health, and gap summaries | Optional |
| `context` | 360-degree symbol view, categorized refs, process participation, source-site proof/status, and related ResolutionGap rows | Optional |
| `impact` | Blast radius analysis with affected App Layers, Functional Areas, resolution-health risks, depth grouping, and confidence | Optional |
| `detect_changes` | Git-diff impact with changed App Layers, affected App Layers, ResolutionGap changes, and resolution-health impact | Optional |
| `rename` | Multi-file coordinated rename with graph + text search | Optional |
| `cypher` | Raw Cypher graph queries | Optional |
| `route_map` | API route to handler/consumer mapping with semantic route, consumer, and flow fields | Optional |
| `tool_map` | MCP/RPC tool definition and handler mapping | Optional |
| `shape_check` | API response shape vs consumer access checks with semantic fields | Optional |
| `api_impact` | API route impact with consumer/flow layer summaries and resolution-health impact | Optional |
| `group_list` | List repo groups | — |
| `group_sync` | Build cross-repo contract registry | — |
| `group_contracts` | Inspect group contracts and cross-links | — |
| `group_query` | Search across a repo group | — |
| `group_status` | Check group/repo staleness | — |

> With one indexed repo, the `repo` param is optional. With multiple, specify which: `query({query: "auth", repo: "my-app"})`.

## MCP Resources

| Resource | Purpose |
|----------|---------|
| `avmatrix://repos` | List all indexed repositories (read first) |
| `avmatrix://repo/{name}/context` | Codebase stats, staleness check, and available tools |
| `avmatrix://repo/{name}/clusters` | All functional clusters with cohesion scores |
| `avmatrix://repo/{name}/cluster/{name}` | Cluster members and details |
| `avmatrix://repo/{name}/processes` | All execution flows |
| `avmatrix://repo/{name}/process/{name}` | Full process trace with steps |
| `avmatrix://repo/{name}/schema` | Graph schema for Cypher queries |

## MCP Prompts

| Prompt | What It Does |
|--------|-------------|
| `detect_impact` | Pre-commit change analysis — scope, affected processes, risk level |
| `generate_map` | Architecture documentation from the knowledge graph with mermaid diagrams |

## CLI Commands

```bash
avmatrix setup                      # Configure MCP for your editors (one-time)
avmatrix analyze [path]             # Index a repository (or update stale index)
avmatrix analyze --force            # Force full re-index
avmatrix analyze --embeddings       # Enable embedding generation (slower, better search)
avmatrix analyze --verbose          # Log skipped files when parsers are unavailable
avmatrix mcp                        # Start MCP server (stdio) — serves all indexed repos
avmatrix serve                      # Start local HTTP server (multi-repo) for web UI
avmatrix index                      # Register an existing local index folder into the global registry
avmatrix list                       # List all indexed repositories
avmatrix status                     # Show index status for current repo
avmatrix clean                      # Delete index for current repo
avmatrix clean --all --force        # Delete all indexes
avmatrix wiki [path]                # Generate LLM-powered docs from knowledge graph
avmatrix wiki --model <model>       # Wiki with custom LLM model (default: gpt-4o-mini)

# Graph quality and semantic diagnostics
avmatrix query <search_query>        # Search graph flows/symbols with semantic fields
avmatrix context [name]              # Inspect a symbol/node and related ResolutionGap rows
avmatrix impact [target]             # Inspect blast radius and resolution-health risks
avmatrix detect-changes              # Inspect changed symbols, affected flows, and gap impact
avmatrix query-health                # Benchmark query retrieval accuracy against a suite
avmatrix resolution-inventory        # Report persisted ResolutionGap and Resolution Health inventory
avmatrix source-site-accuracy        # Report source-site and resolved-edge accuracy metrics

# Repository groups (multi-repo / monorepo service tracking)
avmatrix group create <name>        # Create a repository group
avmatrix group add <name> <repo>    # Add a repo to a group
avmatrix group remove <name> <repo> # Remove a repo from a group
avmatrix group list [name]          # List groups, or show one group's config
avmatrix group sync <name>          # Extract contracts and match across repos/services
avmatrix group contracts <name>     # Inspect extracted contracts and cross-links
avmatrix group query <name> <q>     # Search execution flows across all repos in a group
avmatrix group status <name>        # Check staleness of repos in a group
```

## Remote Embeddings

Set these env vars to use a remote OpenAI-compatible `/v1/embeddings` endpoint instead of the local model:

```bash
export AVMATRIX_EMBEDDING_URL=http://your-server:8080/v1
export AVMATRIX_EMBEDDING_MODEL=BAAI/bge-large-en-v1.5
export AVMATRIX_EMBEDDING_DIMS=1024          # optional, default 384
export AVMATRIX_EMBEDDING_API_KEY=your-key   # optional, default: "unused"
avmatrix analyze . --embeddings
```

Works with Infinity, vLLM, TEI, llama.cpp, Ollama, LM Studio, or OpenAI. When unset, local embeddings are used unchanged.

## Multi-Repo Support

AVmatrix supports indexing multiple repositories. Each `avmatrix analyze` registers the repo in the global registry at `~/.avmatrix/registry.json`.

## Supported Languages

TypeScript, JavaScript, Python, Java, C, C++, C#, Go, Rust, PHP, Kotlin, Swift, Ruby

### Language Feature Matrix

| Language | Imports | Named Bindings | Exports | Heritage | Type Annotations | Constructor Inference | Config | Frameworks | Entry Points |
|----------|---------|----------------|---------|----------|-----------------|---------------------|--------|------------|-------------|
| TypeScript | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| JavaScript | ✓ | ✓ | ✓ | ✓ | — | ✓ | ✓ | ✓ | ✓ |
| Python | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| Java | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | — | ✓ | ✓ |
| Kotlin | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | — | ✓ | ✓ |
| C# | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| Go | ✓ | — | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| Rust | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | — | ✓ | ✓ |
| PHP | ✓ | ✓ | ✓ | — | ✓ | ✓ | ✓ | ✓ | ✓ |
| Ruby | ✓ | — | ✓ | ✓ | — | ✓ | — | ✓ | ✓ |
| Swift | — | — | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| C | — | — | ✓ | — | ✓ | ✓ | — | ✓ | ✓ |
| C++ | — | — | ✓ | ✓ | ✓ | ✓ | — | ✓ | ✓ |

**Imports** — cross-file import resolution · **Named Bindings** — `import { X as Y }` / re-export tracking · **Exports** — public/exported symbol detection · **Heritage** — class inheritance, interfaces, mixins · **Type Annotations** — explicit type extraction for receiver resolution · **Constructor Inference** — infer receiver type from constructor calls (`self`/`this` resolution included for all languages) · **Config** — language toolchain config parsing (tsconfig, go.mod, etc.) · **Frameworks** — AST-based framework pattern detection · **Entry Points** — entry point scoring heuristics

## Agent Skills

AVmatrix ships with skill files that teach AI agents how to use the tools effectively:

- **Exploring** — Navigate unfamiliar code using the knowledge graph
- **Debugging** — Trace bugs through call chains
- **Impact Analysis** — Analyze blast radius before changes
- **Refactoring** — Plan safe refactors using dependency mapping

Installed automatically by both `avmatrix analyze` (per-repo) and `avmatrix setup` (global).

## Requirements

- Node.js >= 18
- Git repository (uses git for commit tracking)

## Release candidates

Stable releases publish to the default `latest` dist-tag. When a pull request
with non-documentation changes merges into `main`, an automated workflow also
publishes a prerelease build under the `rc` dist-tag, so early adopters can
try in-flight fixes without waiting for the next stable cut. (Docs-only
merges are skipped.)

```bash
# Try the latest release candidate (pre-stable — may change at any time)
npm install -g avmatrix@rc
# — or —
npx avmatrix@rc analyze
```

Release-candidate versions follow the standard semver prerelease format
`X.Y.Z-rc.N`, where `X.Y.Z` is the next stable target (bumped from the
current `latest` by patch by default; `minor` or `major` when kicking off a
bigger cycle) and `N` increments per published rc. Example sequence:
`1.0.0-rc.1`, `1.0.0-rc.2`, …, then once `1.0.0` ships stable,
`1.0.1-rc.1`. See the [Releases page](https://github.com/tamnguyendinh/AVmatrix/releases)
for the full list; stable `latest` is unaffected.

## Troubleshooting

### `Cannot destructure property 'package' of 'node.target' as it is null`

This crash was caused by a dependency URL format that is incompatible with
certain npm/arborist versions ([npm/cli#8126](https://github.com/npm/cli/issues/8126)).
It is fixed in the current AVmatrix codebase. Upgrade to the latest version:

```bash
npm install -g avmatrix@latest
avmatrix analyze
```

If you still hit npm install issues after upgrading, these generic workarounds
may help:

```bash
npm install -g npm@latest            # update npm itself
npm cache clean --force              # clear a possibly corrupt cache
```

### Installation fails with native module errors

Some optional language grammars (Dart, Kotlin, Swift) require native compilation. If they fail, AVmatrix still works — those languages will be skipped.

If `npm install -g avmatrix` fails on native modules:

```bash
# Ensure build tools are available (Linux/macOS)
# Ubuntu/Debian: sudo apt install python3 make g++
# macOS: xcode-select --install

# Retry installation
npm install -g avmatrix
```

### Analysis runs out of memory

For very large repositories:

```bash
# Increase Node.js heap size
NODE_OPTIONS="--max-old-space-size=16384" avmatrix analyze

# Exclude large directories
echo "vendor/" >> .avmatrixignore
echo "dist/" >> .avmatrixignore
```

## Privacy

- All processing happens locally on your machine
- No code is sent to any server
- AVmatrix stores the local repo index in `.avmatrix/` inside your repo (gitignored)
- AVmatrix keeps the global registry in `~/.avmatrix/`

## Web UI

AVmatrix also has a browser-based UI for local development — your code never leaves the browser.

**Local Backend Mode:** Run `avmatrix serve` and open the web UI locally — it auto-detects the server and shows all your indexed repos, with full AI chat support. No need to re-upload or re-index. The agent's tools (Cypher queries, search, code navigation) route through the backend HTTP API automatically.

## License

[PolyForm Noncommercial 1.0.0](https://polyformproject.org/licenses/noncommercial/1.0.0/)

Free for non-commercial use. Contact for commercial licensing.
