---
name: anvien-debugging
description: "Use when the user is debugging a bug, tracing an error, or asking why something fails."
---

# Debugging With Anvien

Use this skill to trace a bug from symptom to owner using graph facts, runtime evidence, source inspection, and tests.

## Command Choices

| Need | Use |
|---|---|
| Find candidate owners from a symptom | MCP `query` or CLI `anvien query "<symptom>" --repo <repo>` |
| Inspect suspect files | CLI `anvien query files "<symptom>" --repo <repo>` then `anvien context file <path> --repo <repo>` |
| Inspect suspect symbol callers/callees | MCP/CLI `context symbol` |
| Trace a known flow | `anvien://repo/<repo>/process/{name}` |
| Check topology or diagnostic candidates | CLI `anvien graph-health ...` |
| Inspect unresolved references | CLI `anvien resolution-inventory --graph .anvien/graph.json` |
| Check source-site proof quality | CLI `anvien source-site-accuracy --graph .anvien/graph.json` |
| Measure query reliability for a bug class | CLI `anvien query-health --repo <repo> --suite <file>` |

## Workflow

1. Reproduce or capture the symptom first: command, input, log, screenshot, failing test, or runtime trace.
2. Refresh graph evidence with `anvien analyze --force` before graph-based debugging when needed.
3. Use `query` for broad symptom and domain discovery, then verify the candidate owner with file-layer output, `context file`, `context symbol`, and source.
4. Follow process resources when the failure sits in a route/tool/CLI/runtime flow.
5. Use graph-quality commands when the bug may be missing topology, stale graph data, unresolved references, source-site proof, or poor query retrieval.
6. Run `impact file` or `impact symbol` before editing the suspected owner.
7. Add or update focused regression tests for the failure.

## Query Reliability Rule

Broad `query` is useful but not proof. It has multiple lanes and can surface plausible but wrong regions. If broad query misses expected owners, record the miss as graph-quality/query-health evidence. When the exact symbol or file is known, use `context` and source inspection directly.

## Evidence To Record

- Symptom and reproduction command.
- Query/context/process commands used to locate the owner.
- File-layer evidence, ResolutionGap rows, graph-health file hotspots, or source-site proof if relevant.
- Impact result before edits.
- Regression test or validation command after the fix.
- `detect-changes` before commit.

## Current Limitations

Anvien can identify likely owners and dependency paths, but it does not prove runtime state by itself. Confirm with tests, logs, browser/e2e evidence, or CLI smoke commands that exercise the failing behavior.
