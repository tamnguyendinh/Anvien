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
| File-first discovery | CLI `anvien query files "<concept>" --repo <repo>` or MCP `query` with `target_type=files` |
| Symbol, flow, or API discovery | CLI `anvien query symbols|flows|api "<concept>" --repo <repo>` or MCP `query` with the matching `target_type` |
| Exact symbol inspection | MCP `context` or CLI `anvien context symbol "<symbol>" --repo <repo>` |
| Exact file inspection | CLI `anvien context file <path> --repo <repo>` or `anvien file-context <path> --repo <repo> --json` |
| Execution flow detail | `anvien://repo/<repo>/processes` and `anvien://repo/<repo>/process/{name}` |
| Functional area inventory | `anvien://repo/<repo>/clusters` and `anvien://repo/<repo>/cluster/{name}` |
| File hotspot inventory | `anvien graph-health files --repo <repo> --json` or `anvien file-hotspots --repo <repo> --json` |
| Custom graph questions | MCP/CLI `cypher` after higher-level tools narrow the target |

## Workflow

1. Refresh graph evidence with `anvien analyze --force` before graph-based work when repo rules require it or freshness is unclear.
2. Resolve the repo name from `anvien://repos` or `anvien list`.
3. Use the narrowest query child that matches the exploration question when known: `query files`, `query symbols`, `query flows`, or `query api`; use parent `query` for broad or mixed intent.
4. Treat broad `query` output as candidate retrieval. Verify owner regions with `context`, process resources, file projection, and exact source inspection before selecting edit targets.
5. When a file path is already known, use `context file` or `file-context` first, then drill into the symbol tree and relationship/source-site samples. When a symbol is already known, use `context symbol`.
6. Use `graph-health files` or `file-hotspots` when exploration starts from graph risk, unresolved source sites, high fan-in/fan-out, or linked flow/test density.
7. Record missed or noisy query results as graph-quality/query-health evidence instead of silently accepting them.

## Overview To Detail

Start broad with `query` or `analyze` hotspot output, then move to `query files` or `graph-health files` for the file layer. Open the target with `context file`, inspect the symbol tree and relationship/source-site samples, then use `context symbol` or `impact symbol` for the exact function/class/method before editing. For change-set review, use `detect-changes files` to confirm changed-file risk and linked flows/tests.

## Query Interpretation

`query` is a multi-lane discovery command. It can return owner, concept, execution-flow, API-surface, graph-quality, docs/setup/AI-context, command-surface, and cross-repo evidence when the indexed graph supports it. A threshold query-health pass means the result is usable for navigation; exact target coverage is a stricter claim and must be checked separately.

Use `context symbol` to inspect the exact symbol once candidates are known. Use `context file` when the candidate is a file path or when you need containing symbols, derived file relationships, unresolved source sites, linked flows/tests, and quality signals. Use process resources when the question is about runtime order, route/tool flow, or how a subsystem hangs together.

## Validation

- Name the repo and freshness source used.
- Record the command/resource that identified each owner.
- Include process names or symbol names that support the explanation.
- State uncertainty when the graph has ResolutionGap rows or unresolved source-site evidence.
- Before editing any symbol discovered during exploration, switch to impact analysis.

## Current Limitations

Anvien graph facts are only as fresh as the last analyze run. Broad query can still return plausible but incomplete regions, especially for docs/setup or generated-context work. Do not use color, UI grouping, or query rank alone as proof of ownership; confirm with symbol context, source, tests, or query-health evidence.
