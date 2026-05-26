---
name: avmatrix-ai-context
description: "Use when the user changes generated AGENTS.md, CLAUDE.md, embedded AVmatrix skills, AI context generation, or source-vs-generated validation."
---

# AI Context Generation With AVmatrix

Use this skill for changes to generated `AGENTS.md`, `CLAUDE.md`, `.claude/skills/avmatrix/**`, embedded skill Markdown, or AI context generation behavior.

## Source Of Truth

- Root `AGENTS.md` and `CLAUDE.md` managed AVmatrix blocks are generated output.
- `.claude/skills/avmatrix/**/SKILL.md` files are generated output.
- Embedded source skill Markdown lives under `internal/aicontext/skills/*.md`.
- The base skill registry and generated root Skills table are owned by `internal/aicontext/aicontext.go`.
- Analyze post-run generation is reached through `internal/cli/analyze_postrun.go`.

Do not patch generated files as the permanent fix. Update source Markdown/generator code, then regenerate through normal analyze/setup/package paths.

## Command Choices

| Need | Use |
|---|---|
| Regenerate root context and skills | `avmatrix analyze --force` |
| Inspect generation owner | `avmatrix context "GenerateAIContextFiles" --repo <repo>` |
| Inspect base skill install owner | `avmatrix context "installBaseSkills" --repo <repo>` |
| Check generated changed scope | `avmatrix detect-changes --repo <repo> --scope all` |
| Validate query discovery of this area | `avmatrix query-health --repo <repo> --suite docs/query-health/2026-05-23-avmatrix-skill-system-upgrade-suite.json` |

## Workflow

1. Refresh graph evidence before graph-based work.
2. Use `query` for broad AI-context discovery only as candidate retrieval; verify with `context` on exact owners such as `GenerateAIContextFiles`, `installBaseSkills`, `baseSkillContent`, and setup/package owners.
3. Update embedded skill source and generator code, not generated `.claude` output.
4. Add tests that assert final skill ids, generated file paths, Skills table links, frontmatter, command naming by surface, and no fallback placeholder content.
5. Regenerate through `avmatrix analyze --force` before final closure and compare source/generated inventories.

## Validation

- Every registered base skill has a matching embedded Markdown file.
- Frontmatter `name` matches the registry id and `description` is present.
- Generated `.claude/skills/avmatrix/<skill>/SKILL.md` files match the final skill set.
- Generated root Skills table lists every final skill.
- No final skill relies on `fallbackBaseSkillContent`.
- CLI/MCP command names are surface-correct: CLI uses `query-health` and `api route-map`; MCP uses `route_map` and `api_impact`.

## Current Limitations

Setup/package installation may use additional distribution paths. When those paths change, verify inventories and hashes separately instead of assuming embedded analyze output and editor setup output are identical.
