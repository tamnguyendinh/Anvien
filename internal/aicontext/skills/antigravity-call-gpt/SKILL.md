---
name: antigravity-call-gpt
description: Dùng khi user yêu cầu gọi OpenAI GPT/Codex từ Antigravity để làm việc.
---

# Tích hợp OpenAI GPT/Codex vào Antigravity

## Mục đích

Cho phép agent gọi OpenAI Responses API từ bên trong Antigravity thông qua script Python.
Áp dụng cho: subagent lane GPT supervisor hoặc bất kỳ tác vụ nào cần model OpenAI.

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

Cách sử dụng:
1.	Script: Dùng duy nhất script chuẩn có sẵn của skill (scripts/call_gpt.py).
2.	Prompt: Điền nội dung vào đúng Template Universal Lane Contract chuẩn (Role, Lifecycle, 5 Prohibitions, Goal, Scope, Checklist, Handoff Data).
3.	Thực thi: Gọi trực tiếp:
```bash
python .agents/skills/antigravity-call-gpt/scripts/call_gpt.py --prompt-file <prompt_file> --model gpt-6-astra --reasoning-effort high/xhigh/max
```
high/xhigh/max: mức độ reasoning effort (mức độ suy luận) của model, mặc định là `high`. Chọn `xhigh` hoặc `max` tuỳ theo thực tế của task, không nhất thiết phải chọn high theo mặc định nếu đây là task quan trọng, phức tạp, hoặc cần reasoning effort cao hơn.

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
