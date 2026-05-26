---
name: avmatrix-debugging
description: "Use when the user is debugging a bug, tracing an error, or asking why something fails."
---

# Debugging With AVmatrix

Use this skill to trace a bug from symptom to owner using graph facts, runtime evidence, source inspection, and tests.

## Command Choices

| Need | Use |
|---|---|
| Find candidate owners from a symptom | MCP `query` or CLI `avmatrix query "<symptom>" --repo <repo>` |
| Inspect suspect symbol callers/callees | MCP/CLI `context` |
| Trace a known flow | `avmatrix://repo/<repo>/process/{name}` |
| Check topology or diagnostic candidates | CLI `avmatrix graph-health ...` |
| Inspect unresolved references | CLI `avmatrix resolution-inventory --graph .avmatrix/graph.json` |
| Check source-site proof quality | CLI `avmatrix source-site-accuracy --graph .avmatrix/graph.json` |
| Measure query reliability for a bug class | CLI `avmatrix query-health --repo <repo> --suite <file>` |

## Workflow

1. Reproduce or capture the symptom first: command, input, log, screenshot, failing test, or runtime trace.
2. Refresh graph evidence with `avmatrix analyze --force` before graph-based debugging when needed.
3. Use `query` for broad symptom and domain discovery, then verify the candidate owner with `context` and source.
4. Follow process resources when the failure sits in a route/tool/CLI/runtime flow.
5. Use graph-quality commands when the bug may be missing topology, stale graph data, unresolved references, or poor query retrieval.
6. Run impact before editing the suspected owner.
7. Add or update focused regression tests for the failure.

## Query Reliability Rule

Broad `query` is useful but not proof. It has multiple lanes and can surface plausible but wrong regions. If broad query misses expected owners, record the miss as graph-quality/query-health evidence. When the exact symbol or file is known, use `context` and source inspection directly.

## Evidence To Record

- Symptom and reproduction command.
- Query/context/process commands used to locate the owner.
- ResolutionGap or graph-health evidence if relevant.
- Impact result before edits.
- Regression test or validation command after the fix.
- `detect-changes` before commit.

## Current Limitations

AVmatrix can identify likely owners and dependency paths, but it does not prove runtime state by itself. Confirm with tests, logs, browser/e2e evidence, or CLI smoke commands that exercise the failing behavior.
