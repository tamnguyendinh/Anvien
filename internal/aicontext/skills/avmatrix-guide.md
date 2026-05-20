---
name: avmatrix-guide
description: "Use when the user asks about AVmatrix itself, available tools, MCP resources, graph schema, or workflow reference. Examples: \"What AVmatrix tools are available?\", \"How do I use AVmatrix?\""
---

# AVmatrix Tool And Resource Guide

Use this skill as the reference for the local AVmatrix MCP and CLI surfaces.

## MCP Tools

| Tool | Use |
|---|---|
| `list_repos` | List indexed repositories |
| `query` | Process-grouped search for a concept |
| `context` | Symbol callers, callees, refs, and processes |
| `impact` | Blast radius analysis before edits |
| `detect_changes` | Git-diff impact before commit |
| `rename` | Graph-guided symbol rename |
| `cypher` | Read-only graph query |
| `route_map` | Route surface inspection |
| `tool_map` | Tool surface inspection |
| `shape_check` | Contract and shape checks |
| `api_impact` | API route impact analysis |

## Resources

| Resource | Use |
|---|---|
| `avmatrix://repos` | All indexed repos |
| `avmatrix://setup` | Setup and surface reference |
| `avmatrix://repo/{name}/context` | Repo overview and freshness |
| `avmatrix://repo/{name}/clusters` | Functional areas |
| `avmatrix://repo/{name}/cluster/{clusterName}` | Functional area detail |
| `avmatrix://repo/{name}/processes` | Execution flow list |
| `avmatrix://repo/{name}/process/{processName}` | Step-by-step trace |
| `avmatrix://repo/{name}/schema` | Graph schema for Cypher |

## Standard Workflows

### Explore

Use `query` for concepts, then `context` on important symbols, then process resources for traces.

### Edit

Run `impact({target: "<symbol>", direction: "upstream"})` before editing functions, classes, or methods. Warn on HIGH or CRITICAL risk and review direct callers first.

### Commit

Run `detect_changes({scope: "all"})` before committing. Confirm changed symbols and affected flows match the intended scope.

### Rename

Use `rename` instead of text replacement. Start with a dry run, inspect confidence, apply only when the edits make sense, then run `detect_changes`.

## CLI Fallback

When MCP tools are unavailable, use `avmatrix query`, `avmatrix context`, `avmatrix impact`, `avmatrix detect-changes`, and `avmatrix cypher` with `--repo <name>`.

## Stale Index Policy

If any AVmatrix surface reports stale data, refresh with `avmatrix analyze --force` before relying on graph facts.
