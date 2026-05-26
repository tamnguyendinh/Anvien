---
name: avmatrix-guide
description: "Use when the user asks about AVmatrix itself, available tools, MCP resources, graph schema, prompts, or workflow reference."
---

# AVmatrix Guide

Use this skill as the unified reference for AVmatrix command surfaces. CLI commands, MCP tools, resources, prompts, and Web/API views expose the same local repository intelligence through different interfaces.

## Surfaces

| Surface | Use |
|---|---|
| CLI | Terminal workflows, smoke checks, build/package/setup validation, scripts |
| MCP tools | Agent-native graph operations such as `query`, `context`, `impact`, `detect_changes`, `rename`, `route_map`, `tool_map`, `shape_check`, and `api_impact` |
| MCP resources | Stable repo context, clusters, processes, process traces, setup reference, and schema |
| MCP prompts | Agent templates such as `detect_impact` and `generate_map`; these are not CLI commands |
| Web/API | Local runtime graph, panels, route/search views, and browser validation |

## Command Selection

| Need | CLI | MCP/resource |
|---|---|---|
| Refresh graph | `avmatrix analyze --force` | Run CLI from agent shell |
| List repos | `avmatrix list` | `list_repos`, `avmatrix://repos` |
| Broad discovery | `avmatrix query "<concept>" --repo <repo>` | `query` |
| Exact symbol view | `avmatrix context "<symbol>" --repo <repo>` | `context` |
| Blast radius | `avmatrix impact "<symbol>" --repo <repo> --direction upstream` | `impact` |
| Changed-scope review | `avmatrix detect-changes --repo <repo> --scope all` | `detect_changes` |
| Rename | `avmatrix rename <symbol> <newName> --repo <repo>` | `rename` |
| API route map | `avmatrix api route-map [route] --repo <repo>` | `route_map` |
| MCP/tool map | `avmatrix api tool-map [tool] --repo <repo>` | `tool_map` |
| Shape check | `avmatrix api shape-check [route] --repo <repo>` | `shape_check` |
| API impact | `avmatrix api impact [route] --repo <repo>` | `api_impact` |

Do not invent CLI spellings for MCP-only names. The MCP tool is `route_map`; the CLI command is `avmatrix api route-map`. The MCP tool is `api_impact`; the CLI command is `avmatrix api impact`.

## Graph Quality Commands

- `avmatrix graph-health summary|report|components|explain` audits topology, diagnostics, component membership, confidence, and resolution-health overlays.
- `avmatrix query-health` measures retrieval quality with threshold and exact pass modes.
- `avmatrix resolution-inventory` reports persisted ResolutionGap and Resolution Health inventory.
- `avmatrix source-site-accuracy` audits proof/source-site accuracy and resolved-edge quality.
- `avmatrix benchmark-compare` compares analyze benchmark output files.

## Resources

| Resource | Use |
|---|---|
| `avmatrix://repos` | Discover indexed repo names |
| `avmatrix://setup` | Tool, resource, prompt, setup, and command reference |
| `avmatrix://repo/<repo>/context` | Overview and freshness |
| `avmatrix://repo/<repo>/clusters` | Functional areas |
| `avmatrix://repo/<repo>/processes` | Execution flow list |
| `avmatrix://repo/<repo>/process/{name}` | Step-by-step flow trace |
| `avmatrix://repo/<repo>/schema` | Graph schema for Cypher |

## Standard Rules

- Run `avmatrix analyze --force` before graph-based work when freshness is required.
- Run impact before editing important symbols or contracts.
- HIGH/CRITICAL impact is blast-radius evidence to report and account for, not an edit ban.
- Run `detect-changes` before committing implementation work.
- Preserve graph counts, samples, and traceability in evidence.

## Current Limitations

Some surfaces are intentionally not normal user CLI commands. Hidden lifecycle helpers such as `package` and `hook` are for AVmatrix maintenance. MCP prompts guide agents but do not fetch evidence themselves; the receiving agent must read the named tools/resources/commands.
