---
name: avmatrix-cli
description: "Use when the user needs to run AVmatrix CLI commands for analysis, query, graph quality, API parity, groups, setup, runtime diagnostics, completion, package, wiki, hook, or version workflows."
---

# AVmatrix CLI Commands

Use this skill for terminal access to AVmatrix. In a development checkout, prefer the freshly built canonical binary at `avmatrix\bin\avmatrix.exe` when validating product behavior.

## Core Repository Commands

| Command | Use |
|---|---|
| `avmatrix analyze --force` | Rebuild graph evidence and generate AI context. |
| `avmatrix status` | Check index and freshness. |
| `avmatrix list` | List indexed repositories. |
| `avmatrix index [path...]` | Register existing local index folders. |
| `avmatrix clean --force` | Remove the current repo index. |
| `avmatrix clean --all --force` | Remove all indexed repo data. |
| `avmatrix version` | Print version/build information. |
| `avmatrix completion <shell>` | Generate shell completion scripts. |

## Graph Navigation Commands

| Command | Use |
|---|---|
| `avmatrix query "<concept>" --repo <repo>` | Broad multi-lane discovery for concepts, flows, owners, command surfaces, API areas, docs/setup, and graph quality. |
| `avmatrix query --lanes --json` | Discover query capability lanes available to users and agents. |
| `avmatrix context "<symbol>" --repo <repo>` | Exact symbol callers, callees, refs, process membership, source-site proof, and ResolutionGap context. |
| `avmatrix impact "<symbol>" --repo <repo> --direction upstream` | Blast radius before editing. HIGH/CRITICAL are warnings, not blockers. |
| `avmatrix detect-changes --repo <repo> --scope all` | Changed-symbol and affected-flow review before commit. |
| `avmatrix cypher "<query>" --repo <repo>` | Read-only graph query for custom questions. |
| `avmatrix augment "<pattern>"` | Add graph context to text search. |
| `avmatrix rename <symbol> <newName> --repo <repo>` | Graph-guided rename, dry-run by default; use `--apply` only after review. |

## API Surface Commands

CLI API commands delegate to MCP owners so terminal and agent behavior stay aligned.

| Command | MCP owner |
|---|---|
| `avmatrix api route-map [route] --repo <repo>` | `route_map` |
| `avmatrix api tool-map [tool] --repo <repo>` | `tool_map` |
| `avmatrix api shape-check [route] --repo <repo>` | `shape_check` |
| `avmatrix api impact [route] --repo <repo>` | `api_impact` |

## Graph Quality Commands

| Command | Use |
|---|---|
| `avmatrix graph-health summary --repo <repo> --json` | Topology, confidence, diagnostic, component, and resolution-health summary. |
| `avmatrix graph-health report --repo <repo> --limit 20 --json` | Prioritized topology/diagnostic candidates. |
| `avmatrix graph-health components --repo <repo> --json` | Component inventory. |
| `avmatrix graph-health explain <node-or-name> --repo <repo> --json` | Node/component explanation. |
| `avmatrix query-health --repo <repo> --suite <file>` | Retrieval benchmark with threshold and exact pass modes. |
| `avmatrix resolution-inventory --graph .avmatrix/graph.json` | Persisted ResolutionGap and Resolution Health inventory. |
| `avmatrix source-site-accuracy --graph .avmatrix/graph.json` | Source-site proof and resolved-edge accuracy audit. |
| `avmatrix benchmark-compare <before> <after>` | Compare analyze benchmark outputs. |

## Runtime, Setup, Package, And Groups

| Command | Use |
|---|---|
| `avmatrix serve --host 127.0.0.1 --port <port>` | Start local Web/API runtime. |
| `avmatrix mcp` | Start MCP server. |
| `avmatrix setup` | Configure editor/agent integrations. |
| `avmatrix doctor locks --repo <repo> --json` | Inspect analyze lock state for a repo. |
| `avmatrix doctor processes --json` | Inspect AVmatrix runtime processes without stopping them. |
| `avmatrix group list|status|sync|contracts|query` | Multi-repo groups, contracts, and cross-repo query. |
| `avmatrix wiki` and `avmatrix wiki-mode [off|local]` | Wiki capability status and mode. |
| `avmatrix package ...` | Hidden lifecycle package/runtime maintenance. |
| `avmatrix hook claude` | Hidden lifecycle hook integration. |

## Validation

- Check exact flags with `avmatrix <command> --help`.
- Use `--json` for machine-readable smoke output when available.
- Run the full build before command validation in this repo: `powershell -ExecutionPolicy Bypass -File avmatrix-launcher\build.ps1`.
- Do not use a stale `avmatrix` from `PATH` to define final command behavior when a freshly built local binary exists.
