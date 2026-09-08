#!/usr/bin/env python3
"""
call_gpt.py — Script cầu nối gọi OpenAI Responses API từ Antigravity.

Cách dùng:
    python call_gpt.py --prompt "Câu hỏi của bạn"
    python call_gpt.py --prompt "Review code" --context-file src/main.py
    python call_gpt.py --prompt "Viết hàm sort" --code-mode
    python call_gpt.py --prompt "Dịch sang tiếng Anh" --system "You are a translator."
    python call_gpt.py --prompt "Hello" --model gpt-6-astra

Biến môi trường:
    OPENAI_API_KEY   — API key (bắt buộc nếu không truyền --api-key).
    OPENAI_BASE_URL  — Base URL (mặc định: http://localhost:57334/v1).
    OPENAI_MODEL     — Model (mặc định: gpt-6-astra).
"""

from __future__ import annotations

import argparse
import os
import sys
from pathlib import Path


# ┌─────────────────────────────────────────────────┐
# │  CẤU HÌNH MẶC ĐỊNH                              │
# └─────────────────────────────────────────────────┘
DEFAULT_API_KEY  = "agt_codex_pO9igWPgzTk07nO44VoSO2Bmvw8n9Tzb"
DEFAULT_BASE_URL = "http://localhost:57334/v1"
DEFAULT_MODEL    = "gpt-6-astra"


def _resolve(cli_val: str | None, env_name: str, default: str) -> str:
    return cli_val or os.environ.get(env_name, "") or default


def _resolve_api_key(cli_key: str | None) -> str:
    key = _resolve(cli_key, "OPENAI_API_KEY", DEFAULT_API_KEY)
    if not key:
        print(
            "LỖI: Chưa có OPENAI_API_KEY.\n"
            "Set biến môi trường OPENAI_API_KEY hoặc truyền --api-key.",
            file=sys.stderr,
        )
        sys.exit(1)
    return key


def _build_user_content(prompt: str, context_file: str | None) -> str:
    if not context_file:
        return prompt
    path = Path(context_file)
    if not path.is_file():
        print(f"LỖI: File không tồn tại: {context_file}", file=sys.stderr)
        sys.exit(1)
    code_text = path.read_text(encoding="utf-8", errors="replace")
    return f"Code context ({path.name}):\n```\n{code_text}\n```\n\n{prompt}"


def main() -> None:
    if sys.stdout.encoding != "utf-8":
        sys.stdout.reconfigure(encoding="utf-8")  # type: ignore[attr-defined]
    if sys.stderr.encoding != "utf-8":
        sys.stderr.reconfigure(encoding="utf-8")  # type: ignore[attr-defined]

    parser = argparse.ArgumentParser(
        description="Call OpenAI Responses API from Antigravity.",
    )
    parser.add_argument("--prompt", "-p", required=True,
                        help="Prompt gửi tới GPT.")
    parser.add_argument("--model", "-m", default=None,
                        help=f"Model (default: {DEFAULT_MODEL}).")
    parser.add_argument("--system", "-s", default=None,
                        help="Custom system/developer prompt.")
    parser.add_argument("--context-file", "-c", default=None,
                        help="Path to code file to include as context.")
    parser.add_argument("--code-mode", action="store_true",
                        help="Bật system prompt chuyên sinh code.")
    parser.add_argument("--api-key", default=None,
                        help="OpenAI API key (override env var).")
    parser.add_argument("--temperature", "-t", type=float, default=0.7,
                        help="Temperature (default: 0.7).")
    parser.add_argument("--max-tokens", type=int, default=None,
                        help="Giới hạn output tokens.")

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

    # --- Resolve config ---
    api_key  = _resolve_api_key(args.api_key)
    base_url = _resolve(None, "OPENAI_BASE_URL", DEFAULT_BASE_URL)
    model    = _resolve(args.model, "OPENAI_MODEL", DEFAULT_MODEL)

    client = OpenAI(api_key=api_key, base_url=base_url)

    # --- Build input ---
    user_content = _build_user_content(args.prompt, args.context_file)

    system = args.system
    if not system and args.code_mode:
        system = (
            "You are a code generation assistant. "
            "Return clean, production-ready code with brief comments."
        )

    input_items = []
    if system:
        input_items.append({"role": "developer", "content": system})
    input_items.append({"role": "user", "content": user_content})

    kwargs: dict = {
        "model": model,
        "input": input_items,
        "temperature": args.temperature,
    }
    if args.max_tokens is not None:
        kwargs["max_output_tokens"] = args.max_tokens

    # --- Call Responses API ---
    try:
        response = client.responses.create(**kwargs)
    except Exception as exc:
        print(f"ERROR calling OpenAI Responses API: {exc}", file=sys.stderr)
        sys.exit(1)

    # --- Print result ---
    content = ""
    if hasattr(response, "output_text"):
        content = response.output_text
    elif hasattr(response, "output") and response.output:
        for item in response.output:
            if hasattr(item, "content") and item.content:
                for part in item.content:
                    if hasattr(part, "text"):
                        content += part.text
    print(content)

    # Usage → stderr
    usage = getattr(response, "usage", None)
    if usage:
        inp = getattr(usage, "input_tokens", 0)
        out = getattr(usage, "output_tokens", 0)
        print(
            f"\n--- Usage: input={inp} output={out} total={inp + out} ---",
            file=sys.stderr,
        )


if __name__ == "__main__":
    main()
