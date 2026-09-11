---
name: antigravity-call-codex
description: Use when delegating tasks to Codex to automatically execute code, edit files, and run tests directly on the repository via native MCP.
---

# Integrate Codex into Antigravity via Native MCP

## Purpose

Enable Antigravity and other AI agents to delegate end-to-end technical tasks to Codex to execute directly on the repository via the Model Context Protocol (MCP).

Codex possesses full tools to read/write files and execute shell commands, autonomously closing the self-repair loop (edit → build → test → pass) directly on the repository without requiring the host agent to act as an intermediary for copying and pasting code.

---

## 1. Local Global MCP Configuration (One-time Setup)

To make the MCP Server active across all workspaces on this machine, declare it in Antigravity's Global configuration file:

* **Windows Path:** `C:\Users\<USER>\.gemini\config\mcp_config.json` (shorthand: `~/.gemini/config/mcp_config.json`)
* **Configuration Content (YOLO Mode - Full execution permissions for Codex):**

```json
{
  "mcpServers": {
    "codex": {
      "command": "codex.cmd",
      "args": [
        "--dangerously-bypass-approvals-and-sandbox",
        "--dangerously-bypass-hook-trust",
        "mcp-server"
      ]
    }
  }
}
```

* **Verify Status:** In the Antigravity UI, navigate to **Additional Options (`...`) > MCP Servers** to confirm `codex` shows a connected state with 2 tools: `codex` and `codex-reply`.

---

## 2. Execution via Dispatcher Script

To trigger Codex immediately without manual protocol handling, run the dispatcher script from the repository root:

```bash
# In Antigravity / Codex:
python .agents/skills/antigravity-call-codex/scripts/call_codex.py -p "<task description>"

# In Claude Code:
python .claude/skills/antigravity-call-codex/scripts/call_codex.py -p "<task description>"
```

### Flags:
* `-p`, `--prompt`: Detailed task description for Codex (required).
* `-c`, `--cwd`: Working directory root for Codex (optional, defaults to current working directory).

---

## 3. Direct MCP Native Tool Call (Alternative)

If the active environment provides native MCP tool calling directly:

* **ServerName:** `codex`
* **ToolName:** `codex`
* **Arguments:**

```json
{
  "prompt": "<task description>",
  "cwd": "<root_directory_path_of_the_target_repository>"
}
```
