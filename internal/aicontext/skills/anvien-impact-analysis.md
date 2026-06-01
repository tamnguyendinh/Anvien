---
name: anvien-impact-analysis
description: "Use when the user wants to know what will break if they change something, or needs safety analysis before editing code."
---

# Impact Analysis With Anvien

Use this skill before editing functions, classes, methods, exported symbols, API handlers, graph builders, resolvers, analyzers, shared contracts, or generated-context owners.

## Command Choices

| Need | Use |
|---|---|
| Blast radius for one symbol | MCP `impact` or CLI `anvien impact symbol "<symbol>" --repo <repo> --direction upstream` |
| Blast radius for one file | CLI `anvien impact file <path> --repo <repo> --direction upstream` or MCP `impact` with `target_type=file` |
| Disambiguate a target | MCP/CLI `context`; use explicit `context symbol` or `context file` when input could be both |
| Route or contract impact | MCP `api_impact` or CLI `anvien api impact [route] --repo <repo>` |
| Route blast radius | CLI `anvien impact route <route> --repo <repo> --json` or MCP `impact` with `target_type=route` |
| Tool blast radius | CLI `anvien impact tool <tool> --repo <repo> --json` or MCP `impact` with `target_type=tool` |
| Changed-scope review before commit | MCP `detect_changes` or CLI `anvien detect-changes --repo <repo> --scope all`; use `detect-changes files` for grouped file risk |
| Resolution/topology risk context | `graph-health`, `resolution-inventory`, or `source-site-accuracy` when relevant |

## Workflow

1. Run `anvien analyze --force` before graph-based impact work when required by repo rules.
2. Identify the exact symbol, file, route, or tool. If impact returns candidates, use `context symbol`, `context file`, `route-map`, `tool-map`, or `--uid` to disambiguate.
3. Run upstream impact before editing. For file edits, prefer `impact file`; for single owner edits, prefer `impact symbol`; for route/tool targets, use `api impact`, `impact route`, or `impact tool` according to the output shape you need.
4. Report affected files, App Layers, Functional Areas, execution flows, linked tests, route/tool consumers, unresolved deltas, and risk.
5. Proceed carefully on HIGH or CRITICAL impact. Those levels are blast-radius warnings to account for, not a ban on required code changes.
6. Keep the implementation scoped to the reason for the edit.
7. Run focused validation and `detect-changes` before commit.

## Risk Interpretation

| Risk | Meaning |
|---|---|
| LOW | Local or small blast radius. Still run relevant tests. |
| MEDIUM | Multiple dependents or flows need review. |
| HIGH | Important shared code or broad callers; warn clearly and validate carefully. |
| CRITICAL | Core path, command surface, generated context, runtime, or many flows; work is allowed but evidence and tests must be strong. |

HIGH/CRITICAL is not a prohibition. If the requested fix requires editing that area, edit it deliberately, avoid unrelated refactors, and record why the blast radius is acceptable.

## Validation

- Impact command and target UID/name recorded.
- Direct dependents reviewed.
- Affected process names reviewed when present.
- File-layer evidence reviewed: containing file summary, affected files, file risk, linked flows/tests, and derived relationship note.
- Route/tool evidence reviewed when the target is an API route or MCP/RPC tool.
- Tests cover the affected behavior.
- `detect-changes --scope all` matches the intended scope before commit.

## Current Limitations

Impact relies on current graph relationships and source-site resolution. ResolutionGap rows can show unresolved references or analyzer gaps; treat them as risk evidence, not fake resolved dependencies.
