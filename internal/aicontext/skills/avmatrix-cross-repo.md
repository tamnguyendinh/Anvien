---
name: avmatrix-cross-repo
description: "Use when the user works across indexed repository groups, cross-repo query, contracts, status, sync, or multi-repo ownership."
---

# Cross-Repo Work With AVmatrix

Use this skill when a task spans multiple indexed repositories or a repository group.

## Command Choices

| Need | CLI | MCP |
|---|---|---|
| List/inspect groups | `avmatrix group list [name]` | `group_list` |
| Add repo to a group | `avmatrix group add <group> <groupPath> <registryName>` | CLI |
| Remove repo from a group | `avmatrix group remove <group> <path>` | CLI |
| Check freshness across group repos | `avmatrix group status <name>` | `group_status` |
| Sync contract registry | `avmatrix group sync <name>` | `group_sync` |
| Search execution flows across repos | `avmatrix group query <name> "<query>"` | `group_query` |
| Inspect group contracts | `avmatrix group contracts <name>` | `group_contracts` |

## Workflow

1. Confirm individual repo freshness with `avmatrix analyze --force` where needed.
2. Use group status before relying on cross-repo evidence.
3. Run group sync before contract checks when contracts may have changed.
4. Use group query for cross-repo execution-flow discovery, then inspect exact repo/symbol context in the owning repo.
5. Record which repo supplied each fact; do not collapse multi-repo evidence into a single local assumption.

## Query Guidance

Cross-repo query can return process and contract evidence from several repos. Treat it as candidate discovery. When a target repo/symbol is known, switch to that repo's `context`, `impact`, and tests.

## Validation

- Group name and member repos recorded.
- Status freshness recorded.
- Contract sync output recorded when contracts matter.
- Matched repo, file, symbol, and process evidence recorded for cross-repo decisions.
- `detect-changes` still runs in the repo being edited.

## Current Limitations

Cross-repo impact is only as complete as the indexed group and synced contracts. Missing or stale group members must be handled as uncertainty, not ignored.
