---
name: antigravity-call-gpt
description: Dùng khi user yêu cầu gọi OpenAI GPT/Codex từ Antigravity để sinh code, dịch, tóm tắt, hoặc hỏi đáp qua API OpenAI.
---

# Tích hợp OpenAI GPT/Codex vào Antigravity

## Mục đích

Cho phép agent gọi OpenAI Responses API từ bên trong Antigravity thông qua script Python.
Áp dụng cho: sinh code, dịch thuật, tóm tắt, hỏi đáp, review, hoặc bất kỳ tác vụ nào cần model OpenAI.

Không dùng skill này khi tác vụ đã được Antigravity/Gemini xử lý tốt mà không cần model bên ngoài.

## Cấu hình (đã gắn sẵn trong script, chạy là chạy)

| Biến | Giá trị |
|------|---------|
| `OPENAI_API_KEY` | `agt_codex_pO9igWPgzTk07nO44VoSO2Bmvw8n9Tzb` |
| `OPENAI_BASE_URL` | `http://localhost:57334/v1` |
| `OPENAI_MODEL` | `gpt-6-astra` |
| Endpoint | `/responses` (Responses API) |

## Yêu cầu

- Python 3.10+.
- Package `openai`: `pip install openai`.

## Cách dùng

```bash
# Hỏi đáp
python internal/aicontext/skills/antigravity-call-gpt/scripts/call_gpt.py -p "Giải thích async/await"

# Truyền code context từ file
python internal/aicontext/skills/antigravity-call-gpt/scripts/call_gpt.py -p "Review code này" -c src/main.py

# Chế độ sinh code
python internal/aicontext/skills/antigravity-call-gpt/scripts/call_gpt.py -p "Viết hàm merge sort" --code-mode

# Custom system prompt
python internal/aicontext/skills/antigravity-call-gpt/scripts/call_gpt.py -p "Dịch sang tiếng Anh" -s "You are a translator."
```

## CLI flags

| Flag | Mô tả | Mặc định |
|------|--------|----------|
| `-p`, `--prompt` | Prompt (bắt buộc) | — |
| `-m`, `--model` | Model | `gpt-6-astra` |
| `-s`, `--system` | Developer/system prompt | — |
| `-c`, `--context-file` | File code làm context | — |
| `--code-mode` | Bật system prompt sinh code | `false` |
| `--api-key` | Override API key | — |
| `-t`, `--temperature` | Temperature | `0.7` |
| `--max-tokens` | Giới hạn output tokens | — |

## Tham khảo

- `scripts/call_gpt.py`: Script chính.
- OpenAI Responses API: https://platform.openai.com/docs/api-reference/responses
