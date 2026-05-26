---
name: avmatrix-runtime-packaging
description: "Use when the user needs serve, mcp, setup, launcher, package runtime, canonical executable, startup, or process lifecycle validation."
---

# Runtime And Packaging With AVmatrix

Use this skill for local runtime startup, MCP server startup, editor setup, launcher behavior, package lifecycle helpers, canonical executable validation, and process lifecycle bugs.

## Command Choices

| Need | Use |
|---|---|
| Build the repo's product gate | `powershell -ExecutionPolicy Bypass -File avmatrix-launcher\build.ps1` |
| Start Web/API runtime | `avmatrix serve --host 127.0.0.1 --port <port>` |
| Start MCP server | `avmatrix mcp` |
| Configure editor/agent integrations | `avmatrix setup` |
| Verify package runtime | `avmatrix package ensure-runtime` |
| Build package runtime | `avmatrix package build-runtime` |
| Prepare/clean package source | `avmatrix package prepare-go-source`, `avmatrix package clean-go-source` |
| Print version | `avmatrix version` |

In this repository, product validation should use the canonical executable path `avmatrix\bin\avmatrix.exe` after the full build. Do not treat `avmatrix-launcher\server-bundle\avmatrix.exe` as a second production CLI authority.

## Workflow

1. Run the full build before runtime/package validation.
2. Use explicit host/port for local server smoke checks.
3. Track started process ids and stop the exact process tree you started.
4. For setup/package skill behavior, verify installed/generated skill inventories and content hashes rather than assuming package-root files match embedded skills.
5. For lifecycle bugs, validate that browser close, app exit, failed analyze, and reset paths do not leave orphan backend/analyze processes.

## Evidence To Record

- Built executable path, size, and version.
- Runtime command, host, port, pid, readiness output, and cleanup result.
- Setup/package target directories and file inventories.
- Any process cleanup checks before and after validation.

## Current Limitations

Some package and hook commands are hidden lifecycle helpers, not normal repo-analysis commands. They can still be tested and documented for maintainers when package/setup behavior changes.
