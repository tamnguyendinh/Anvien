---
name: anvien-runtime-packaging
description: "Use when the user needs serve, mcp, setup, doctor diagnostics, launcher, package runtime, canonical executable, startup, or process lifecycle validation."
---

# Runtime And Packaging With Anvien

Use this skill for local runtime startup, MCP server startup, editor setup, doctor diagnostics, launcher behavior, package lifecycle helpers, canonical executable validation, and process lifecycle bugs.

## Command Choices

| Need | Use |
|---|---|
| Build the repo's product gate | `powershell -ExecutionPolicy Bypass -File anvien-launcher\build.ps1` |
| Start Web/API runtime | `anvien serve --host 127.0.0.1 --port <port>` |
| Start MCP server | `anvien mcp` |
| Configure editor/agent integrations | `anvien setup` |
| Inspect analyze lock state | `anvien doctor locks --repo <repo> --json` |
| Inspect runtime processes | `anvien doctor processes --json` |
| Generate shell completion | `anvien completion <shell>` |
| Check wiki capability | `anvien wiki`, `anvien wiki-mode [off|local]` |
| Verify package runtime | `anvien package ensure-runtime` |
| Build package runtime | `anvien package build-runtime` |
| Prepare/clean package source | `anvien package prepare-go-source`, `anvien package clean-go-source` |
| Print version | `anvien version` |

In this repository, product validation should use the canonical executable path `anvien\bin\anvien.exe` after the full build. Do not treat `anvien-launcher\server-bundle\anvien.exe` as a second production CLI authority.

## Workflow

1. Run the full build before runtime/package validation.
2. Use explicit host/port for local server smoke checks.
3. Use `anvien doctor processes --json` to inspect ownership before deciding whether a process is editor-owned, user-command-owned, launcher-owned, or diagnostic-command.
4. Treat MCP `Auth: Unsupported` as normal for the local stdio server. Validate actual availability through the tool list and a smoke tool/resource call when the client supports it.
5. Track started process ids and stop the exact process tree you started.
6. For setup/package skill behavior, verify installed/generated skill inventories and content hashes from embedded `internal/aicontext/skills/*.md`; generated `.claude/skills/anvien/**` and package-root `skills/` files are not source of truth.
7. For lifecycle bugs, validate that browser close, app exit, failed analyze, and reset paths do not leave orphan backend/analyze processes.

## Evidence To Record

- Built executable path, size, and version.
- Runtime command, host, port, pid, readiness output, and cleanup result.
- Doctor lock/process output when diagnosing stuck analyze or orphan runtime behavior.
- Setup/package target directories, file inventories, and embedded-source-to-installed-skill hash checks when skills are involved.
- Any process cleanup checks before and after validation.

## Current Limitations

Some package and hook commands are hidden lifecycle helpers, not normal repo-analysis commands. They can still be tested and documented for maintainers when package/setup behavior changes.
