---
name: avmatrix-api-surface
description: "Use when the user needs to inspect API routes, MCP tools, contract shape drift, generated Web contracts, handlers, consumers, or route/tool impact."
---

# API Surface With AVmatrix

Use this skill for route handlers, MCP/RPC tool handlers, generated contracts, response shape drift, route consumers, and API-impact decisions.

## Command Choices

| Need | CLI | MCP |
|---|---|---|
| Route handlers, consumers, middleware, flows | `avmatrix api route-map [route] --repo <repo> --json` | `route_map` |
| MCP/tool definitions and linked flows | `avmatrix api tool-map [tool] --repo <repo> --json` | `tool_map` |
| Response shape drift against consumers | `avmatrix api shape-check [route] --repo <repo> --json` | `shape_check` |
| Route/API blast radius | `avmatrix api impact [route] --repo <repo> --json` | `api_impact` |
| Exact handler context | `avmatrix context "<handler>" --repo <repo>` | `context` |
| Broad API discovery | `avmatrix query "API route <concept>" --repo <repo>` | `query` |

MCP tool names use underscores: `route_map`, `tool_map`, `shape_check`, `api_impact`. CLI commands use hyphenated subcommands under `avmatrix api`. Do not invent CLI commands by reusing MCP underscore names as top-level AVmatrix commands.

## Workflow

1. Refresh the graph with `avmatrix analyze --force` before graph-based API work.
2. Use route-map/tool-map to find handlers and consumers.
3. Use shape-check before changing response contracts or generated Web contracts.
4. Use API impact before editing handlers, schemas, contracts, or shared API helpers.
5. Validate with focused backend tests and Web contract/client tests when consumers are affected.
6. Run `detect-changes --scope all` before commit.

## Evidence To Record

- Route/tool selector used and whether it was ambiguous.
- Handler file, consumer count, middleware, flow count, and shape-check mismatches.
- API impact risk and affected App Layers/Functional Areas.
- Contract or generated-client validation commands.

## Current Limitations

API-surface graph quality depends on route/tool extraction and source-site resolution. If a route or tool is missing, record it as graph-quality evidence and verify source manually before concluding the API does not exist.
