---
name: avmatrix-debugging
description: "Use when the user is debugging a bug, tracing an error, or asking why something fails. Examples: \"Why is X failing?\", \"Where does this error come from?\", \"Trace this bug\""
---

# Debugging With AVmatrix

Use this skill to trace a bug through execution flows, callers, callees, and graph-backed context before changing code.

## Workflow

1. Refresh the graph if stale with `avmatrix analyze --force`.
2. Run `query({query: "<error or symptom>"})` to find relevant execution flows.
3. Run `context({name: "<suspect symbol>"})` for callers, callees, and process membership.
4. Read `avmatrix://repo/{name}/process/{processName}` for the full step trace when the failure sits inside a known flow.
5. Use `cypher` only for targeted graph questions that `query` and `context` do not answer.

## Evidence To Collect

- The failing symptom or error text.
- The process or route/tool flow that reaches the failing code.
- Incoming callers that can trigger the path.
- Outgoing calls or external dependencies that can return bad data.
- Tests or fixtures that already cover the failing path.

## Tool Patterns

`query({query: "timeout while saving invoice"})`

Finds processes, symbols, and modules related to the symptom.

`context({name: "saveInvoice"})`

Shows what calls the suspect symbol, what it calls, and which flows include it.

`impact({target: "saveInvoice", direction: "upstream"})`

Use before editing the suspect symbol to understand what may break.

`detect_changes({scope: "all"})`

Run after the fix and before commit to verify the affected scope matches the bug fix.

## Debugging Checklist

- [ ] Query by symptom, error text, and domain concept.
- [ ] Inspect the most relevant process trace.
- [ ] Inspect the suspect symbol with `context`.
- [ ] Identify direct callers and external dependencies.
- [ ] Run impact before edits.
- [ ] Add or update focused tests for the failing path.
- [ ] Run `detect_changes` before commit.

## Interpretation

Do not assume the first matching symbol is the root cause. Treat AVmatrix output as navigation and impact evidence, then confirm behavior in source and tests.
