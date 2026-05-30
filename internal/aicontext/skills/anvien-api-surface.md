---
name: anvien-api-surface
description: "Use when the user needs to inspect API routes, MCP tools, contract shape drift, generated Web contracts, handlers, consumers, or route/tool impact."
---

# API Surface With Anvien

Use this skill for route handlers, MCP/RPC tool handlers, generated contracts, response shape drift, route consumers, and API-impact decisions.

## Command Choices

| Need | CLI | MCP |
|---|---|---|
| Route handlers, consumers, middleware, flows | `anvien api route-map [route] --repo <repo> --json` | `route_map` |
| MCP/tool definitions and linked flows | `anvien api tool-map [tool] --repo <repo> --json` | `tool_map` |
| Response shape drift against consumers | `anvien api shape-check [route] --repo <repo> --json` | `shape_check` |
| Route/API blast radius | `anvien api impact [route] --repo <repo> --json` | `api_impact` |
| Exact handler file context | `anvien context file <path> --repo <repo>` | `context` with `target_type=file` |
| Exact handler symbol context | `anvien context symbol "<handler>" --repo <repo>` | `context` with `target_type=symbol` |
| Broad API discovery | `anvien query "API route <concept>" --repo <repo>` | `query` |

MCP tool names use underscores: `route_map`, `tool_map`, `shape_check`, `api_impact`. CLI commands use hyphenated subcommands under `anvien api`. Do not invent CLI commands by reusing MCP underscore names as top-level Anvien commands.

## Workflow

1. Refresh the graph with `anvien analyze --force` before graph-based API work.
2. Use route-map/tool-map to find handlers, consumers, flows, and handler-file `handlerFile` projection data.
3. Open handler files with `context file` when you need symbol tree, file dependencies, linked tests, or unresolved handler-file sites.
4. Use shape-check before changing response contracts or generated Web contracts.
5. Use API impact before editing handlers, schemas, contracts, or shared API helpers.
6. Validate with focused backend tests and Web contract/client tests when consumers are affected.
7. Run `detect-changes --scope all` before commit.

## Evidence To Record

- Route/tool selector used and whether it was ambiguous.
- Handler file, handler-file summary, symbol tree/dependency counts, unresolved handler-file sites, consumer count, middleware, flow count, linked tests, and shape-check mismatches.
- API impact risk and affected App Layers/Functional Areas.
- Contract or generated-client validation commands.

## Current Limitations

API-surface graph quality depends on route/tool extraction and source-site resolution. If a route or tool is missing, record it as graph-quality evidence and verify source manually before concluding the API does not exist.
