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

That's it. This indexes the codebase and writes managed `AGENTS.md`, `CLAUDE.md`, and `.claude/skills/avmatrix/**` context files from AVmatrix's embedded skill source.

To configure MCP for your editor, run `avmatrix setup` once — or set it up manually below.

`avmatrix setup` auto-detects your editors, writes the correct global MCP config, and installs the same embedded base skills into supported editor skill directories. You only need to run it once.

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
avmatrix setup                       # Configure MCP/editor access and embedded skills
avmatrix analyze [path]              # Index a repository and refresh AI context
avmatrix analyze --force             # Force full re-index
avmatrix analyze --embeddings        # Enable embedding generation
avmatrix analyze --no-stats          # Omit volatile stats from generated agent files
avmatrix analyze --skip-git          # Analyze a folder without requiring .git
avmatrix analyze --name <alias>      # Register repo under a custom name
avmatrix index [path...]             # Register existing local indexes
avmatrix list                        # List indexed repositories
avmatrix status                      # Show index status for current repo
avmatrix clean --all --force         # Delete all indexes
avmatrix mcp                         # Start MCP server over stdio
avmatrix serve                       # Start local HTTP backend for the Web UI
avmatrix doctor                      # Inspect local runtime locks and processes
avmatrix version                     # Print version/build information
avmatrix wiki                        # Show wiki capability status
avmatrix wiki-mode [off|local]       # Show or set local wiki capability mode
avmatrix completion <shell>          # Generate shell completion script
```

Direct graph, API, and quality tools:

```bash
avmatrix query <search_query>         # Multi-lane graph discovery
avmatrix query --lanes --json         # List query capability lanes
avmatrix context [name]               # Inspect a symbol/node and related graph facts
avmatrix impact [target]              # Inspect blast radius and resolution-health risks
avmatrix detect-changes               # Inspect changed symbols, affected flows, and gap impact
avmatrix cypher <query>               # Run read-only graph queries
avmatrix augment <pattern>            # Add graph context to text search
avmatrix rename <symbol> <newName>    # Graph-guided rename, dry-run by default
avmatrix api route-map [route]        # Route handlers, consumers, and linked flows
avmatrix api tool-map [tool]          # MCP/RPC tool handlers and linked flows
avmatrix api shape-check [route]      # Response shape drift against consumers
avmatrix api impact [route]           # Route/API blast radius and shape risk
avmatrix graph-health                 # Topology, diagnostics, components, explanations
avmatrix query-health                 # Query retrieval benchmark
avmatrix resolution-inventory         # Persisted ResolutionGap inventory
avmatrix source-site-accuracy         # Source-site and resolved-edge accuracy metrics
avmatrix benchmark-compare <a> <b>    # Compare analyze benchmark outputs
```

Repository groups:

```bash
avmatrix group create <name>
avmatrix group add <group> <groupPath> <registryName>
avmatrix group remove <group> <path>
avmatrix group list [name]
avmatrix group status <name>
avmatrix group sync <name>
avmatrix group contracts <name>
avmatrix group query <name> <query>
```

`query-health` reports two separate outcomes. The threshold result says whether
hit@5/hit@10 found enough expected targets to make retrieval usable for an
agent, while the exact result says whether every expected file/symbol was found.
Use `--fail-on-threshold` for usable-retrieval gates and `--fail-on-exact` for
strict target-coverage gates.

`graph-health` audits topology status, component membership, diagnostics,
confidence, resolution-health overlays, and prioritized candidates from the
indexed graph. `rename` and `api ...` commands are CLI equivalents for MCP
`rename`, `route_map`, `tool_map`, `shape_check`, and `api_impact`, so terminal
and agent workflows use the same local owner logic.

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

AVmatrix ships embedded base skills that teach AI agents how to use the tool surfaces effectively:

- **Exploring** — architecture, ownership, and execution flow discovery
- **Impact Analysis** — blast radius, HIGH/CRITICAL warnings, and changed-scope checks
- **Debugging** — failures, diagnostics, and graph-quality evidence
- **Refactoring** — rename, extract, split, move, and restructure work
- **Guide** — unified CLI, MCP, resource, prompt, and Web/API reference
- **CLI** — terminal command guide for AVmatrix CLI surfaces
- **Graph Quality** — graph-health, query-health, resolution inventory, and accuracy audits
- **API Surface** — API routes, MCP tools, shape checks, contracts, and consumers
- **Cross Repo** — repository groups, cross-repo query, contracts, status, and sync
- **Runtime Packaging** — runtime, setup, launcher, package, and process lifecycle workflows
- **AI Context** — generated `AGENTS.md`, `CLAUDE.md`, embedded skills, and validation

`avmatrix analyze` installs the per-repo generated output under `.claude/skills/avmatrix/**`. `avmatrix setup` installs the same embedded base skill content into supported editor skill directories. Package-root `skills/` files are not a source of truth.

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
