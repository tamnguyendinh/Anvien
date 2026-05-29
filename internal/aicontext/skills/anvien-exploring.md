---
name: anvien-exploring
description: "Use when the user asks how code works, wants to understand architecture, trace execution flows, or explore unfamiliar parts of the codebase."
---

# Exploring Codebases With Anvien

Use this skill when the task is to understand a repo before editing. Exploration should produce concrete owners, entrypoints, flows, and open questions, not guesses.

## Command Choices

| Need | Use |
|---|---|
| Confirm repo names and freshness | `anvien://repos`, `anvien://repo/<repo>/context`, or `anvien status` |
| Broad concept discovery | MCP `query` or CLI `anvien query "<concept>" --repo <repo>` |
| Exact symbol inspection | MCP `context` or CLI `anvien context "<symbol>" --repo <repo>` |
| Execution flow detail | `anvien://repo/<repo>/processes` and `anvien://repo/<repo>/process/{name}` |
| Functional area inventory | `anvien://repo/<repo>/clusters` and `anvien://repo/<repo>/cluster/{name}` |
| Custom graph questions | MCP/CLI `cypher` after higher-level tools narrow the target |

## Workflow

1. Refresh graph evidence with `anvien analyze --force` before graph-based work when repo rules require it or freshness is unclear.
2. Resolve the repo name from `anvien://repos` or `anvien list`.
3. Use `query` for broad concepts, product terms, symptoms, command surfaces, and flow discovery.
4. Treat broad `query` output as candidate retrieval. Verify owner regions with `context`, process resources, and exact source inspection before selecting edit targets.
5. When a symbol or file is already known, prefer `context` and source inspection over another broad query.
6. Record missed or noisy query results as graph-quality/query-health evidence instead of silently accepting them.

## Query Interpretation

`query` is a multi-lane discovery command. It can return owner, concept, execution-flow, API-surface, graph-quality, docs/setup/AI-context, command-surface, and cross-repo evidence when the indexed graph supports it. A threshold query-health pass means the result is usable for navigation; exact target coverage is a stricter claim and must be checked separately.

Use `context` to inspect the exact symbol once candidates are known. Use process resources when the question is about runtime order, route/tool flow, or how a subsystem hangs together.

## Validation

- Name the repo and freshness source used.
- Record the command/resource that identified each owner.
- Include process names or symbol names that support the explanation.
- State uncertainty when the graph has ResolutionGap rows or unresolved source-site evidence.
- Before editing any symbol discovered during exploration, switch to impact analysis.

## Current Limitations

Anvien graph facts are only as fresh as the last analyze run. Broad query can still return plausible but incomplete regions, especially for docs/setup or generated-context work. Do not use color, UI grouping, or query rank alone as proof of ownership; confirm with symbol context, source, tests, or query-health evidence.
