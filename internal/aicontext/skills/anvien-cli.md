---
name: anvien-cli
description: "Use when the user needs to run Anvien CLI commands for analysis, query, graph quality, API parity, groups, setup, runtime diagnostics, completion, package, wiki, hook, or version workflows."
---

# Anvien CLI Commands

Use this skill for terminal access to Anvien. In a development checkout, prefer the freshly built canonical binary at `anvien\bin\anvien.exe` when validating product behavior.

## Core Repository Commands

| Command | Use |
|---|---|
| `anvien analyze --force` | Rebuild graph evidence, generate AI context, and print file projection inventory/hotspots after graph counts. |
| `anvien status` | Check index and freshness. |
| `anvien list` | List indexed repositories. |
| `anvien index [path...]` | Register existing local index folders. |
| `anvien clean --force` | Remove the current repo index. |
| `anvien clean --all --force` | Remove all indexed repo data. |
| `anvien version` | Print version/build information. |
| `anvien completion <shell>` | Generate shell completion scripts. |

## Graph Navigation Commands

| Command | Use |
|---|---|
| `anvien query "<concept>" --repo <repo>` | Broad multi-lane discovery for concepts, flows, owners, files, command surfaces, API areas, docs/setup, and graph quality. |
| `anvien query files "<concept>" --repo <repo>` | File-first discovery with matched symbols and relationship hints. |
| `anvien query --lanes --json` | Discover query capability lanes available to users and agents. |
| `anvien context symbol "<symbol>" --repo <repo>` | Exact symbol callers, callees, refs, process membership, containing file summary, source-site proof, and ResolutionGap context. |
| `anvien context file <path> --repo <repo>` | Full file context with summary, symbol tree, relationships, unresolved sites, linked flows/tests, and quality signals. |
| `anvien impact symbol "<symbol>" --repo <repo> --direction upstream` | Symbol blast radius plus containing-file evidence before editing. HIGH/CRITICAL are warnings, not blockers. |
| `anvien impact file <path> --repo <repo> --direction upstream` | Aggregate impact from contained symbols and affected file groups. |
| `anvien detect-changes --repo <repo> --scope all` | Changed-symbol, changed-file, affected-file, and affected-flow review before commit. |
| `anvien detect-changes files --repo <repo> --scope all` | File-grouped change review with file risk and unresolved deltas. |
| `anvien cypher "<query>" --repo <repo>` | Read-only graph query for custom questions. |
| `anvien augment "<pattern>"` | Add graph context to text search. |
| `anvien rename <symbol> <newName> --repo <repo>` | Graph-guided rename, dry-run by default; use `--apply` only after review. |
| `anvien file-context <path> --repo <repo> --json` | Dedicated file projection detail for source, test, generated, docs, or config files. |
| `anvien file-hotspots --repo <repo> --sort unresolved --json` | Dedicated file hotspot list for unresolved, fan-in, fan-out, symbol, flow, or test sorting. |

## API Surface Commands

CLI API commands delegate to MCP owners so terminal and agent behavior stay aligned.

| Command | MCP owner |
|---|---|
| `anvien api route-map [route] --repo <repo>` | `route_map` |
| `anvien api tool-map [tool] --repo <repo>` | `tool_map` |
| `anvien api shape-check [route] --repo <repo>` | `shape_check` |
| `anvien api impact [route] --repo <repo>` | `api_impact` |

## Graph Quality Commands

| Command | Use |
|---|---|
| `anvien graph-health summary --repo <repo> --json` | Topology, confidence, diagnostic, component, and resolution-health summary. |
| `anvien graph-health report --repo <repo> --limit 20 --json` | Prioritized topology/diagnostic candidates. |
| `anvien graph-health files --repo <repo> --json` | File-level graph-health rows for unresolved gaps, fan-in, fan-out, linked flows/tests, and risk. |
| `anvien graph-health components --repo <repo> --json` | Component inventory. |
| `anvien graph-health explain <node-or-name> --repo <repo> --json` | Node/component explanation. |
| `anvien query-health --repo <repo> --suite <file>` | Retrieval benchmark with threshold and exact pass modes. |
| `anvien resolution-inventory --graph .anvien/graph.json` | Persisted ResolutionGap and Resolution Health inventory. |
| `anvien source-site-accuracy --graph .anvien/graph.json` | Source-site proof and resolved-edge accuracy audit. |
| `anvien benchmark-compare <before> <after>` | Compare analyze benchmark outputs. |

## Runtime, Setup, Package, And Groups

| Command | Use |
|---|---|
| `anvien serve --host 127.0.0.1 --port <port>` | Start local Web/API runtime. |
| `anvien mcp` | Start MCP server. |
| `anvien setup` | Configure editor/agent integrations. |
| `anvien doctor locks --repo <repo> --json` | Inspect analyze lock state for a repo. |
| `anvien doctor processes --json` | Inspect Anvien runtime processes without stopping them. |
| `anvien group list|status|sync|contracts|query` | Multi-repo groups, contracts, and cross-repo query. |
| `anvien wiki` and `anvien wiki-mode [off|local]` | Wiki capability status and mode. |
| `anvien package ...` | Hidden lifecycle package/runtime maintenance. |
| `anvien hook claude` | Hidden lifecycle hook integration. |

## Validation

- Check exact flags with `anvien <command> --help`.
- Use `--json` for machine-readable smoke output when available.
- For file-path inputs, prefer explicit child commands (`context file`, `impact file`, `query files`, `detect-changes files`) when scripts or agents need deterministic shape.
- Run the full build before command validation in this repo: `powershell -ExecutionPolicy Bypass -File anvien-launcher\build.ps1`.
- Do not use a stale `anvien` from `PATH` to define final command behavior when a freshly built local binary exists.
