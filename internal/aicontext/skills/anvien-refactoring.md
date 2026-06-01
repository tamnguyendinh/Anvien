---
name: anvien-refactoring
description: "Use when the user wants to rename, extract, split, move, or restructure code safely."
---

# Refactoring With Anvien

Use this skill for behavior-preserving renames, extractions, moves, splits, and shared-code cleanup.

## Command Choices

| Need | Use |
|---|---|
| Understand current ownership | MCP/CLI `context`; prefer `context file` or `context symbol` once target type is known |
| Check callers and flows before editing | MCP/CLI `impact --direction upstream`; prefer `impact file` or `impact symbol` once target type is known |
| Rename a symbol | MCP `rename` or CLI `anvien rename <symbol> <newName> --repo <repo>` |
| Check API route/contract consumers | MCP `api_impact` or CLI `anvien api impact [route] --repo <repo>` |
| Check route/tool blast radius | `anvien impact route <route> --repo <repo>` or `anvien impact tool <tool> --repo <repo>` |
| Verify changed scope before commit | MCP `detect_changes` or CLI `anvien detect-changes --repo <repo> --scope all` |

## Workflow

1. Refresh the graph with `anvien analyze --force` before graph-based refactoring.
2. Use `context` to inspect the exact target. If multiple candidates exist, use UID/file disambiguation.
3. Run upstream impact and report the blast radius. HIGH/CRITICAL means proceed carefully, not stop automatically.
4. For renames, start with a graph-guided dry run. Do not use find-and-replace for symbols.
5. Make behavior-preserving changes first. Defer unrelated cleanup unless required.
6. Run focused tests, broader tests when shared contracts are touched, and `detect-changes`.

## Rename Flow

1. `anvien rename oldName newName --repo <repo> --json`
2. Inspect files, edit counts, graph edits, text-search edits, and ambiguity warnings.
3. Use `--uid` or `--file` if the dry run is ambiguous.
4. Apply with `--apply` only when the edit list matches the intended scope.

## Extract, Split, Or Move Flow

- Use `context` to identify callers/callees and nearby contracts.
- Use `impact` to identify consumers that must remain compatible.
- Move code in a small slice, update imports deliberately, and preserve behavior.
- Re-run tests that cover affected flows and any API shape tests if contracts moved.

## Current Limitations

Graph-guided rename is safer than text replacement, but generated files, dynamic references, string-based APIs, and external integrations still need source review and tests.
