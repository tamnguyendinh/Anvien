---
name: anvien-guide
description: "Use when the user asks about Anvien itself, available tools, MCP resources, graph schema, prompts, or workflow reference."
---

# Anvien Guide

Use this skill as the unified reference for Anvien command surfaces. CLI commands, MCP tools, resources, prompts, and Web/API views expose the same local repository intelligence through different interfaces.

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
| Refresh graph | `anvien analyze --force` | Run CLI from agent shell |
| List repos | `anvien list` | `list_repos`, `anvien://repos` |
| Broad discovery | `anvien query "<concept>" --repo <repo>` | `query` |
| File-first discovery | `anvien query files "<concept>" --repo <repo>` | `query` with `target_type=files` |
| Exact symbol view | `anvien context symbol "<symbol>" --repo <repo>` | `context` with `target_type=symbol` |
| Exact file view | `anvien context file <path> --repo <repo>` or `anvien file-context <path> --repo <repo>` | `context` with `target_type=file` |
| Symbol blast radius | `anvien impact symbol "<symbol>" --repo <repo> --direction upstream` | `impact` with `target_type=symbol` |
| File blast radius | `anvien impact file <path> --repo <repo> --direction upstream` | `impact` with `target_type=file` |
| Changed-scope review | `anvien detect-changes --repo <repo> --scope all`; file view: `detect-changes files` | `detect_changes` |
| Rename | `anvien rename <symbol> <newName> --repo <repo>` | `rename` |
| API route map | `anvien api route-map [route] --repo <repo>` | `route_map` |
| MCP/tool map | `anvien api tool-map [tool] --repo <repo>` | `tool_map` |
| Shape check | `anvien api shape-check [route] --repo <repo>` | `shape_check` |
| API impact | `anvien api impact [route] --repo <repo>` | `api_impact` |

Do not invent CLI spellings for MCP-only names. The MCP tool is `route_map`; the CLI command is `anvien api route-map`. The MCP tool is `api_impact`; the CLI command is `anvien api impact`.

File relationships are a projection from symbol/source-site graph facts. Use file-layer sections to navigate from overview to file to symbol to relationship/source-site samples, but keep canonical graph claims tied to the underlying symbol, route, tool, process, and source-site evidence.

## Graph Quality Commands

- `anvien graph-health summary|report|components|explain` audits topology, diagnostics, component membership, confidence, and resolution-health overlays.
- `anvien graph-health files` and `anvien file-hotspots` show file-level unresolved, fan-in, fan-out, linked flow/test, and risk signals.
- `anvien query-health` measures retrieval quality with threshold and exact pass modes.
- `anvien resolution-inventory` reports persisted ResolutionGap and Resolution Health inventory with file groups and nearest source symbols.
- `anvien source-site-accuracy` audits proof/source-site accuracy and resolved-edge quality with file groups and trace samples.
- `anvien benchmark-compare` compares analyze benchmark output files.

## Resources

| Resource | Use |
|---|---|
| `anvien://repos` | Discover indexed repo names |
| `anvien://setup` | Tool, resource, prompt, setup, and command reference |
| `anvien://repo/<repo>/context` | Overview and freshness |
| `anvien://repo/<repo>/clusters` | Functional areas |
| `anvien://repo/<repo>/processes` | Execution flow list |
| `anvien://repo/<repo>/process/{name}` | Step-by-step flow trace |
| `anvien://repo/<repo>/schema` | Graph schema for Cypher |

## Standard Rules

- Run `anvien analyze --force` before graph-based work when freshness is required.
- Run impact before editing important symbols or contracts.
- HIGH/CRITICAL impact is blast-radius evidence to report and account for, not an edit ban.
- Run `detect-changes` before committing implementation work.
- Preserve graph counts, samples, and traceability in evidence.

## Current Limitations

Some surfaces are intentionally not normal user CLI commands. Hidden lifecycle helpers such as `package` and `hook` are for Anvien maintenance. MCP prompts guide agents but do not fetch evidence themselves; the receiving agent must read the named tools/resources/commands.
