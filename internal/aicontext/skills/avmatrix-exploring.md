---
name: avmatrix-exploring
description: "Use when the user asks how code works, wants to understand architecture, trace execution flows, or explore unfamiliar parts of the codebase. Examples: \"How does X work?\", \"What calls this function?\", \"Show me the auth flow\""
---

# Exploring Codebases With AVmatrix

Use this skill when you need to understand unfamiliar code before editing it.

## Workflow

1. Read `avmatrix://repos` to discover indexed repositories when the repo name is unknown.
2. Read `avmatrix://repo/{name}/context` for overview, stats, and stale-index warnings.
3. Use `query({query: "<concept>"})` to find process-grouped execution flows.
4. Use `context({name: "<symbol>"})` for a 360-degree symbol view.
5. Read `avmatrix://repo/{name}/process/{processName}` for step-by-step flow details.
6. Read source files only after the graph points to the relevant files.

## Resources

| Resource | Use |
|---|---|
| `avmatrix://repos` | All indexed repos |
| `avmatrix://repo/{name}/context` | Repo overview and freshness check |
| `avmatrix://repo/{name}/clusters` | Functional areas |
| `avmatrix://repo/{name}/cluster/{clusterName}` | Area members and files |
| `avmatrix://repo/{name}/processes` | Execution flow list |
| `avmatrix://repo/{name}/process/{processName}` | Detailed process trace |
| `avmatrix://repo/{name}/schema` | Graph schema for Cypher |

## Tool Patterns

`query({query: "authentication"})`

Returns processes and symbols related to authentication.

`context({name: "validateSession"})`

Returns incoming calls, outgoing calls, process participation, and nearby graph facts.

`cypher({query: "MATCH (n:Function) RETURN n.id LIMIT 10"})`

Use only for custom graph questions after checking the higher-level tools.

## Exploration Checklist

- [ ] Confirm repo name and freshness.
- [ ] Query by product concept or subsystem.
- [ ] Inspect processes before individual files.
- [ ] Use `context` on key symbols.
- [ ] Read source after graph navigation narrows the search.
- [ ] Record open questions and uncertain edges instead of guessing.

## When To Stop Exploring

Stop broad exploration once you can name the relevant entrypoints, core symbols, affected files, and tests. Switch to impact analysis before making code changes.
