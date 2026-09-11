#!/usr/bin/env python3
"""
call_codex.py — Dispatcher script to trigger Codex via native MCP stdio.
Automatically connects to Codex in YOLO mode with full permissions.
"""

from __future__ import annotations

import argparse
import json
import os
import shutil
import subprocess
import sys


def _resolve_codex_command() -> str:
    """Find codex executable (prefers codex.cmd on Windows)."""
    if sys.platform == "win32":
        cmd = shutil.which("codex.cmd")
        if cmd:
            return cmd
        # Fallback to standard npm roaming path
        appdata = os.environ.get("APPDATA", "")
        npm_cmd = os.path.join(appdata, "npm", "codex.cmd")
        if os.path.isfile(npm_cmd):
            return npm_cmd
    cmd = shutil.which("codex")
    if cmd:
        return cmd
    return "codex.cmd" if sys.platform == "win32" else "codex"


def main() -> None:
    # Ensure UTF-8 output on Windows
    if sys.stdout.encoding != "utf-8":
        sys.stdout.reconfigure(encoding="utf-8")  # type: ignore[attr-defined]
    if sys.stderr.encoding != "utf-8":
        sys.stderr.reconfigure(encoding="utf-8")  # type: ignore[attr-defined]

    parser = argparse.ArgumentParser(
        description="Dispatch tasks to Codex via MCP in YOLO mode.",
    )
    parser.add_argument(
        "--prompt", "-p",
        required=True,
        help="Task description / prompt for Codex.",
    )
    parser.add_argument(
        "--cwd", "-c",
        default=None,
        help="Working directory root for Codex (defaults to current working directory).",
    )
    parser.add_argument(
        "--model", "-m",
        default=None,
        help="Model override for Codex (e.g. gpt-6-astra).",
    )
    parser.add_argument(
        "--thinking", "-t",
        default=None,
        help="Thinking / reasoning effort override (e.g. low, high, xhigh).",
    )

    args = parser.parse_args()

    cwd = os.path.abspath(args.cwd) if args.cwd else os.getcwd()
    codex_bin = _resolve_codex_command()

    cmd = [
        codex_bin,
        "--dangerously-bypass-approvals-and-sandbox",
        "--dangerously-bypass-hook-trust",
    ]
    if args.model:
        cmd.extend(["-c", f'model="{args.model}"'])
    if args.thinking:
        cmd.extend(["-c", f'model_reasoning_effort="{args.thinking}"'])
    cmd.append("mcp-server")

    try:
        proc = subprocess.Popen(
            cmd,
            stdin=subprocess.PIPE,
            stdout=subprocess.PIPE,
            stderr=subprocess.PIPE,
            encoding="utf-8",
            errors="replace",
            cwd=cwd,
        )
    except Exception as exc:
        print(f"ERROR launching Codex MCP server: {exc}", file=sys.stderr)
        sys.exit(1)

    assert proc.stdin is not None
    assert proc.stdout is not None

    try:
        # 1. initialize
        init_req = {
            "jsonrpc": "2.0",
            "id": 1,
            "method": "initialize",
            "params": {
                "protocolVersion": "2024-11-05",
                "capabilities": {},
                "clientInfo": {"name": "antigravity-runner", "version": "1.0"},
            },
        }
        proc.stdin.write(json.dumps(init_req) + "\n")
        proc.stdin.flush()
        proc.stdout.readline()

        # 2. notifications/initialized
        notif = {"jsonrpc": "2.0", "method": "notifications/initialized", "params": {}}
        proc.stdin.write(json.dumps(notif) + "\n")
        proc.stdin.flush()

        # 3. tools/call
        arguments = {
            "prompt": args.prompt,
            "cwd": cwd,
        }
        if args.model:
            arguments["model"] = args.model
        if args.thinking:
            arguments["config"] = {"model_reasoning_effort": args.thinking}

        call_req = {
            "jsonrpc": "2.0",
            "id": 2,
            "method": "tools/call",
            "params": {
                "name": "codex",
                "arguments": arguments,
            },
        }
        proc.stdin.write(json.dumps(call_req) + "\n")
        proc.stdin.flush()

        # Wait for result
        while True:
            line = proc.stdout.readline()
            if not line:
                break
            try:
                data = json.loads(line)
                if data.get("id") == 2:
                    result = data.get("result", {})
                    content = result.get("content", "")
                    if content:
                        print(content)
                    break
                if "error" in data and data.get("id") == 2:
                    err = data.get("error", {})
                    print(f"Codex error: {err.get('message', err)}", file=sys.stderr)
                    sys.exit(1)
            except json.JSONDecodeError:
                continue

    finally:
        try:
            proc.terminate()
            proc.wait(timeout=3)
        except Exception:
            proc.kill()


if __name__ == "__main__":
    main()
