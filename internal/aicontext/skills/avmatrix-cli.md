---
name: avmatrix-cli
description: "Use when the user needs to run AVmatrix CLI commands like analyze/index a repo, check status, clean the index, inspect wiki capability mode, or list indexed repos. Examples: \"Index this repo\", \"Reanalyze the codebase\", \"Check wiki mode\""
---

# AVmatrix CLI Commands

Use this skill when a task needs terminal access to the local AVmatrix index. Canonical commands use `avmatrix` when the CLI is on `PATH`; in development, `go run .\cmd\avmatrix ...` exercises the same command surface.

## Workflow

1. Run `avmatrix status` to check whether the current repo is indexed and fresh.
2. Run `avmatrix analyze --force` before graph-based work that depends on current source facts.
3. Use `avmatrix list` when a command needs an explicit `--repo` name.
4. Use `avmatrix query`, `avmatrix context`, `avmatrix impact`, `avmatrix detect-changes`, `avmatrix rename`, or `avmatrix api ...` when MCP tools are unavailable or when a terminal smoke check is required.

## Commands

### analyze

`avmatrix analyze --force`

Builds or refreshes the graph, writes `.avmatrix/graph.json`, registers the repo, updates generated AI context files, and installs AVmatrix skills. Analyze always updates the managed AVmatrix sections in root `AGENTS.md` and `CLAUDE.md`. Use `--no-stats` when volatile symbol and relationship counts should be omitted from generated context.

### status

`avmatrix status`

Shows whether the current repo has an index, when it was indexed, and whether the current commit matches the indexed commit.

### list

`avmatrix list`

Lists registered repositories. Use the listed name with `--repo` when multiple repos are indexed.

### query

`avmatrix query "payment flow" --repo MyRepo`

Searches process-grouped code intelligence for a concept or symptom. Current output includes semantic graph fields when the analyzed graph provides them: App Layer, Functional Area, topology status, resolution confidence, ResolutionGap counts, and gap summaries.

### context

`avmatrix context "validateUser" --repo MyRepo --content`

Shows callers, callees, process participation, source-site proof/status metadata, related ResolutionGap rows, and optionally source content for one symbol.

### impact

`avmatrix impact "validateUser" --repo MyRepo --direction upstream --depth 3`

Runs blast-radius analysis before editing a symbol. Review direct callers first. HIGH and CRITICAL risk are blast-radius warnings, not command-output blockers; the output includes affected App Layers, affected Functional Areas, and resolution-health risks when available.

### detect-changes

`avmatrix detect-changes --repo MyRepo --scope all`

Maps current git changes to affected symbols and execution flows. Run before committing. Current output includes changed/affected App Layers, changed/affected Functional Areas, ResolutionGap changes, and resolution-health impact when persisted graph data provides them.

### rename

`avmatrix rename oldName newName --repo MyRepo --json`

Runs the same graph-guided rename engine exposed by MCP `rename`. The default is a dry run; inspect `files_affected`, `total_edits`, and edit confidence before using `--apply`. Use `--uid` or `--file` when a symbol name is ambiguous.

### api

`avmatrix api route-map /api/users --repo MyRepo --json`

CLI parity for API-surface MCP tools. Use `api route-map` for route handlers, consumers, middleware, and linked flows; `api tool-map` for MCP/RPC tool definitions and linked flows; `api shape-check` for response-shape drift against consumers; and `api impact` before changing route handlers or API contracts. These commands delegate to MCP `route_map`, `tool_map`, `shape_check`, and `api_impact` so terminal and agent workflows do not drift.

### query-health

`avmatrix query-health --repo MyRepo --out .tmp/query-health.json`

Runs a query retrieval benchmark suite and reports hit@5, hit@10, expected files/symbols, actual top results, noise reason, and two separate pass results. `thresholdPassed` means the query returned enough expected targets to be usable for navigation; `exactPassed` means no expected file/symbol target was missed. Use `--fail-on-threshold` for usable-retrieval gates and `--fail-on-exact` for strict target-coverage gates.

### graph-health

`avmatrix graph-health summary --repo MyRepo --json`

Audits computed topology health, component membership, diagnostic counts, confidence buckets, and resolution-health overlays for an indexed repo. Use `graph-health report --repo MyRepo --limit 20 --json` for prioritized topology/diagnostic triage candidates, `graph-health components --repo MyRepo --json` for component summaries, and `graph-health explain <node-id-or-name> --repo MyRepo --json` or `graph-health explain --component <component-id> --repo MyRepo --json` for traceable evidence.

### resolution-inventory

`avmatrix resolution-inventory --graph .avmatrix/graph.json --out .tmp/resolution-inventory.json`

Reports full persisted ResolutionGap and Resolution Health inventory: gap nodes, gap occurrences, resolved references, App Layer counts, Functional Area counts, fact families, target roles, classifications, actionability, topology, and the non-actionable `builtin` / `standard_library` / `test_framework` breakdown.

### source-site-accuracy

`avmatrix source-site-accuracy --graph .avmatrix/graph.json --out .tmp/source-site-accuracy.json`

Reports proof-based source-site and resolved-edge accuracy metrics, including missing source-site IDs, resolved edges without proof, false resolved edge candidates, non-property ACCESSES targets, duplicate/merged source-site evidence, and optional golden fixture validation with `--golden`.

### cypher

`avmatrix cypher "MATCH (n:Function) RETURN n.id LIMIT 5" --repo MyRepo`

Runs read-only graph queries against the indexed graph.

## Safety Notes

- If an AVmatrix command reports a stale index, refresh with `avmatrix analyze --force`.
- Prefer MCP tools when available inside an MCP-capable agent; use CLI commands for terminal workflows, smoke validation, or environments without MCP.
- Smoke tests that validate generated root context files must run `avmatrix analyze --force` normally so `AGENTS.md` and `CLAUDE.md` are refreshed.
- Do not treat unresolved or diagnostic references as proven topology. Persisted ResolutionGap entities are diagnostic graph facts, not fake resolved symbols.
- Keep graph-quality commands distinct: `graph-health` audits topology, components, confidence, and diagnostics; `query-health` measures retrieval quality; `resolution-inventory` reports persisted ResolutionGap/Resolution Health inventory; `source-site-accuracy` checks proof/source-site accuracy.
- Hidden lifecycle helpers such as `avmatrix package ...` and `avmatrix hook claude` are for AVmatrix packaging/setup maintenance, not normal repo analysis.
