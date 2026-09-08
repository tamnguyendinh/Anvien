#!/usr/bin/env python3
"""
call_gpt.py — Script cầu nối gọi OpenAI GPT/Codex API từ Antigravity.

Cách dùng:
    python call_gpt.py --prompt "Câu hỏi của bạn"
    python call_gpt.py --prompt "Review code" --context-file src/main.py
    python call_gpt.py --prompt "Viết hàm sort" --code-mode
    python call_gpt.py --prompt "Dịch sang tiếng Anh" --system "You are a translator."
    python call_gpt.py --prompt "Hello" --model gpt-4o-mini
    python call_gpt.py --prompt "Hello" --api-key sk-...

Biến môi trường:
    OPENAI_API_KEY  — API key OpenAI (bắt buộc nếu không truyền --api-key).
"""

from __future__ import annotations

import argparse
import os
import sys
from pathlib import Path



# ┌─────────────────────────────────────────────────┐
# │  CẤU HÌNH MẶC ĐỊNH                              │
# └─────────────────────────────────────────────────┘
DEFAULT_API_KEY = "agt_codex_pO9igWPgzTk07nO44VoSO2Bmvw8n9Tzb"
DEFAULT_BASE_URL = "http://localhost:57334/v1"


def _resolve_api_key(cli_key: str | None) -> str:
    """Lấy API key theo thứ tự: CLI flag → env var → hardcoded."""
    key = cli_key or os.environ.get("OPENAI_API_KEY", "") or DEFAULT_API_KEY
    if not key:
        print(
            "LỖI: Chưa có OPENAI_API_KEY.\n"
            "Set biến môi trường OPENAI_API_KEY hoặc truyền --api-key.",
            file=sys.stderr,
        )
        sys.exit(1)
    return key


def _build_messages(
    prompt: str,
    system: str | None,
    context_file: str | None,
    code_mode: bool,
) -> list[dict[str, str]]:
    """Xây danh sách messages cho ChatCompletion."""
    messages: list[dict[str, str]] = []

    # System prompt
    if code_mode and not system:
        system = (
            "You are a code generation assistant. "
            "Return clean, production-ready code with brief comments."
        )
    if system:
        messages.append({"role": "system", "content": system})

    # User prompt (có thể kèm code context)
    user_content = prompt
    if context_file:
        path = Path(context_file)
        if not path.is_file():
            print(f"LỖI: File không tồn tại: {context_file}", file=sys.stderr)
            sys.exit(1)
        code_text = path.read_text(encoding="utf-8", errors="replace")
        user_content = f"Code context ({path.name}):\n```\n{code_text}\n```\n\n{prompt}"

    messages.append({"role": "user", "content": user_content})
    return messages


def main() -> None:
    # Force UTF-8 output on Windows to avoid cp1252 encoding errors
    if sys.stdout.encoding != "utf-8":
        sys.stdout.reconfigure(encoding="utf-8")  # type: ignore[attr-defined]
    if sys.stderr.encoding != "utf-8":
        sys.stderr.reconfigure(encoding="utf-8")  # type: ignore[attr-defined]

    parser = argparse.ArgumentParser(
        description="Call OpenAI GPT/Codex API from Antigravity.",
    )
    parser.add_argument(
        "--prompt", "-p",
        required=True,
        help="Prompt content to send to GPT.",
    )
    parser.add_argument(
        "--model", "-m",
        default="gpt-6-astra",
        help="OpenAI model name (default: gpt-6-astra).",
    )
    parser.add_argument(
        "--system", "-s",
        default=None,
        help="Custom system prompt (optional).",
    )
    parser.add_argument(
        "--context-file", "-c",
        default=None,
        help="Path to a code file to include as context.",
    )
    parser.add_argument(
        "--code-mode",
        action="store_true",
        help="Enable code generation mode (adds code-focused system prompt).",
    )
    parser.add_argument(
        "--api-key",
        default=None,
        help="OpenAI API key (overrides OPENAI_API_KEY env var).",
    )
    parser.add_argument(
        "--temperature", "-t",
        type=float,
        default=0.7,
        help="Temperature for response (default: 0.7).",
    )
    parser.add_argument(
        "--max-tokens",
        type=int,
        default=None,
        help="Max tokens limit for response.",
    )

    args = parser.parse_args()

    # --- Check openai package ---
    try:
        from openai import OpenAI  # noqa: E402
    except ImportError:
        print(
            "ERROR: Package 'openai' is not installed.\n"
            "Run: pip install openai",
            file=sys.stderr,
        )
        sys.exit(1)

    # --- Call API ---
    api_key = _resolve_api_key(args.api_key)
    base_url = os.environ.get("OPENAI_BASE_URL") or DEFAULT_BASE_URL
    client = OpenAI(api_key=api_key, base_url=base_url)

    messages = _build_messages(
        prompt=args.prompt,
        system=args.system,
        context_file=args.context_file,
        code_mode=args.code_mode,
    )

    create_kwargs: dict = {
        "model": args.model,
        "messages": messages,
        "temperature": args.temperature,
    }
    if args.max_tokens is not None:
        create_kwargs["max_tokens"] = args.max_tokens

    try:
        response = client.chat.completions.create(**create_kwargs)
    except Exception as exc:
        print(f"ERROR calling OpenAI API: {exc}", file=sys.stderr)
        sys.exit(1)

    # --- Print result ---
    content = response.choices[0].message.content or ""
    print(content)

    # Print usage info to stderr (so it doesn't mix with main output)
    if response.usage:
        print(
            f"\n--- Usage: prompt={response.usage.prompt_tokens} "
            f"completion={response.usage.completion_tokens} "
            f"total={response.usage.total_tokens} ---",
            file=sys.stderr,
        )


if __name__ == "__main__":
    main()
