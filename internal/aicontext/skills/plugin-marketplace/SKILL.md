---
name: plugin-marketplace
description: Use when creating, reviewing, or troubleshooting Claude Code plugin marketplaces, marketplace.json files, plugin sources, hosting, team allowlists, or install/update failures.
license: MIT
---

# Plugin Marketplace

Operational guide for building and maintaining Claude Code plugin marketplaces.

## Activation Boundary

Use this skill only when the task involves plugin marketplace distribution:
- Creating or editing `.claude-plugin/marketplace.json`
- Publishing a marketplace to GitHub, GitLab, Bitbucket, a git URL, or a local path
- Defining plugin source entries for marketplace installation
- Reviewing team marketplace allowlists or managed marketplace restrictions
- Debugging marketplace validation, install, update, or authentication failures

Do not use this skill for ordinary `SKILL.md` creation or skill quality review unless the user is also asking to distribute the skill through a plugin marketplace.

## Core Rule

Inspect the target repository or marketplace folder before giving instructions.

Confirm:
- Existing `.claude-plugin/marketplace.json`
- Existing plugin folders and `.claude-plugin/plugin.json` manifests
- Whether sources are relative paths, GitHub repositories, git URLs, npm packages, or local paths
- Whether the marketplace is public, private, team-managed, or local-only
- Which command surface the user is using: Claude Code slash commands, CLI validation, or repository configuration

Do not invent marketplace fields. Follow the existing manifest shape or the schema reference.

## Workflow

1. Determine whether the user needs a marketplace catalog, a plugin entry, hosting guidance, or troubleshooting.
2. Inspect existing marketplace files before editing.
3. Use the schema reference for field names and required/optional structure.
4. Use the source-types reference before choosing where plugin code is loaded from.
5. Use hosting guidance when the task involves sharing, private repositories, or team policy.
6. Use troubleshooting guidance only for observed validation, install, update, path, or auth failures.

## Reference Navigation

- Marketplace setup: `references/plugin-marketplace-setup-overview.md`
- Marketplace JSON schema: `references/plugin-marketplace-json-schema.md`
- Plugin source types: `references/plugin-marketplace-source-types.md`
- Hosting and distribution: `references/plugin-marketplace-hosting-distribution.md`
- Errors and fixes: `references/plugin-marketplace-errors-and-fixes.md`

## Output Rule

For marketplace changes, produce the smallest valid manifest or config change that fits the repository's current distribution model.

For troubleshooting, report the observed failure, likely cause, verification step, and minimal fix.

