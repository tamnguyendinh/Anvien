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
4. Use `avmatrix query`, `avmatrix context`, `avmatrix impact`, or `avmatrix detect-changes` when MCP tools are unavailable.

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

Searches process-grouped code intelligence for a concept or symptom.

### context

`avmatrix context "validateUser" --repo MyRepo --content`

Shows callers, callees, process participation, and optionally source content for one symbol.

### impact

`avmatrix impact "validateUser" --repo MyRepo --direction upstream --depth 3`

Runs blast-radius analysis before editing a symbol. Review direct callers first.

### detect-changes

`avmatrix detect-changes --repo MyRepo --scope all`

Maps current git changes to affected symbols and execution flows. Run before committing.

### cypher

`avmatrix cypher "MATCH (n:Function) RETURN n.id LIMIT 5" --repo MyRepo`

Runs read-only graph queries against the indexed graph.

## Safety Notes

- If an AVmatrix command reports a stale index, refresh with `avmatrix analyze --force`.
- Prefer MCP tools when available; use CLI commands as a fallback or for smoke validation.
- Smoke tests that validate generated root context files must run `avmatrix analyze --force` normally so `AGENTS.md` and `CLAUDE.md` are refreshed.
