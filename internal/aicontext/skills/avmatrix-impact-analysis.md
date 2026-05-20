---
name: avmatrix-impact-analysis
description: "Use when the user wants to know what will break if they change something, or needs safety analysis before editing code. Examples: \"Is it safe to change X?\", \"What depends on this?\", \"What will break?\""
---

# Impact Analysis With AVmatrix

Use this skill before editing a function, class, method, API surface, or shared module.

## Workflow

1. Refresh the graph if stale with `avmatrix analyze --force`.
2. Run `impact({target: "<symbol>", direction: "upstream"})` before editing.
3. Review direct depth-1 dependents first; these are most likely to break.
4. Inspect affected processes and modules.
5. Warn the user before proceeding when risk is HIGH or CRITICAL.
6. After editing, run focused tests and `detect_changes({scope: "all"})`.

## Impact Directions

| Direction | Question |
|---|---|
| `upstream` | What callers, flows, or modules depend on this symbol? |
| `downstream` | What callees or dependencies does this symbol use? |

Use upstream for most edit safety checks. Use downstream when you need to understand data flow or dependency behavior.

## Risk Interpretation

| Risk | Meaning |
|---|---|
| LOW | Small or local blast radius |
| MEDIUM | Multiple dependents or flows need review |
| HIGH | Broad caller or flow impact; warn before edits |
| CRITICAL | Core path or many flows affected; narrow the change and validate heavily |

## Tool Patterns

`impact({target: "validateUser", direction: "upstream"})`

Finds callers and affected flows before changing `validateUser`.

`context({name: "validateUser"})`

Use after impact to inspect the target and disambiguate candidates.

`detect_changes({scope: "all"})`

Use after edits to verify the changed symbols and affected flows match the intended scope.

## Checklist

- [ ] Identify the exact target symbol.
- [ ] Run upstream impact before editing.
- [ ] Report direct callers, affected processes, and risk.
- [ ] Warn before HIGH or CRITICAL changes.
- [ ] Keep the implementation scoped to the stated blast radius.
- [ ] Run tests that cover affected flows.
- [ ] Run `detect_changes` before commit.

## Guardrail

Do not treat impact output as permission to skip source review. It is a map of likely affected code; source and tests decide correctness.
