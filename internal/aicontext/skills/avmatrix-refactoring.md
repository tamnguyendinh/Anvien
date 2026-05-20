---
name: avmatrix-refactoring
description: "Use when the user wants to rename, extract, split, move, or restructure code safely. Examples: \"Rename this function\", \"Extract this into a module\", \"Refactor this class\", \"Move this to a separate file\""
---

# Refactoring With AVmatrix

Use this skill for symbol renames, extraction, moves, and shared-code restructuring.

## Workflow

1. Refresh the graph if stale with `avmatrix analyze --force`.
2. Use `context({name: "<target>"})` to understand the symbol and disambiguate candidates.
3. Use `impact({target: "<target>", direction: "upstream"})` before editing.
4. For renames, use `rename` instead of find-and-replace.
5. Make the smallest refactor that satisfies the goal.
6. Run focused tests and `detect_changes({scope: "all"})`.

## Rename Flow

1. Run `rename({symbol_name: "oldName", new_name: "newName", dry_run: true})`.
2. Inspect files, line edits, and confidence.
3. Apply only when the edit list is correct.
4. Run tests and `detect_changes`.

## Extract Or Split Flow

1. Use `context` to identify callers and callees.
2. Use `impact` to identify consumers that must stay compatible.
3. Extract without changing behavior first.
4. Update imports and call sites deliberately.
5. Run tests for affected flows.

## Move Flow

1. Check current module ownership with `context`.
2. Review affected modules from `impact`.
3. Move code and update references.
4. Confirm no unrelated files changed.

## Checklist

- [ ] Exact target symbol is known.
- [ ] `context` has been reviewed.
- [ ] `impact` has been run before edits.
- [ ] HIGH or CRITICAL impact has been reported before proceeding.
- [ ] Rename operations use `rename`, not text replacement.
- [ ] Tests cover changed flows.
- [ ] `detect_changes` confirms expected scope.

## Guardrail

Avoid opportunistic cleanup. Refactoring should preserve behavior unless the user explicitly asked for behavior change.
