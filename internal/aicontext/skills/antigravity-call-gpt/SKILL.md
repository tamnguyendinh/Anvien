---
name: antigravity-call-gpt
description: Dùng khi user yêu cầu gọi OpenAI GPT/Codex từ Antigravity để làm việc.
---

# Tích hợp OpenAI GPT/Codex vào Antigravity

## Mục đích

Cho phép agent gọi OpenAI GPT/Codex API từ bên trong Antigravity thông qua script Python.
Áp dụng cho: sinh code, dịch thuật, tóm tắt, hỏi đáp, review, hoặc bất kỳ tác vụ nào cần model OpenAI.

Không dùng skill này khi tác vụ đã được Antigravity/Gemini xử lý tốt mà không cần model bên ngoài.

## Yêu cầu trước khi dùng

- Python 3.10+ đã cài trên máy.
- Package `openai` đã cài: `pip install openai` (hoặc `uv pip install openai`).
- Biến môi trường `OPENAI_API_KEY` đã được set, hoặc truyền trực tiếp qua flag `--api-key`.
- Biến môi trường `OPENAI_BASE_URL` (tuỳ chọn) — set khi dùng proxy/gateway local (vd: `http://localhost:57334/v1`).

## Quy tắc vận hành

- Luôn ưu tiên dùng biến môi trường `OPENAI_API_KEY`; không hardcode key vào lệnh terminal.
- Nếu `OPENAI_API_KEY` chưa set, hỏi user trước khi tiếp tục.
- Model mặc định: `gpt-6-astra`. Đổi model bằng flag `--model`.
- Kết quả trả về là text thuần; agent tự parse và xử lý tiếp.
- Nếu cần truyền code context dài, dùng flag `--context-file` thay vì paste inline.
- Không gọi API OpenAI cho tác vụ mà Antigravity đã xử lý đủ tốt.

## Cách dùng ví dụ:

### Cách 1: Gọi script trực tiếp

```bash
# Hỏi đáp đơn giản
python internal/aicontext/skills/antigravity-call-gpt/scripts/call_gpt.py --prompt "Giải thích async/await trong Python"

# Chọn model khác
python internal/aicontext/skills/antigravity-call-gpt/scripts/call_gpt.py --prompt "Viết unit test cho hàm sort" --model gpt-4o-mini

# Truyền code context từ file
python internal/aicontext/skills/antigravity-call-gpt/scripts/call_gpt.py --prompt "Review code này" --context-file src/main.py

# Chế độ sinh code (thêm system prompt chuyên code)
python internal/aicontext/skills/antigravity-call-gpt/scripts/call_gpt.py --prompt "Viết hàm merge sort" --code-mode

# Tuỳ chỉnh system prompt
python internal/aicontext/skills/antigravity-call-gpt/scripts/call_gpt.py --prompt "Dịch sang tiếng Anh" --system "You are a professional translator."
```

### Cách 2: Gọi inline bằng Python one-liner

```bash
python -c "
from openai import OpenAI; import os
c = OpenAI(api_key=os.environ['OPENAI_API_KEY'])
r = c.chat.completions.create(model='gpt-6-astra', messages=[{'role':'user','content':'Hello'}])
print(r.choices[0].message.content)
"
```

## Workflow

1. Xác định tác vụ cần model OpenAI (sinh code, dịch, tóm tắt, review…).
2. Kiểm tra `OPENAI_API_KEY` đã set chưa. Nếu chưa → hỏi user.
3. Chọn cách gọi phù hợp (script hoặc one-liner).
4. Chạy lệnh, đọc output, xử lý tiếp theo yêu cầu user.
5. Nếu output quá dài hoặc cần parse → dùng `--context-file` hoặc pipe output.

## Tham khảo

- `scripts/call_gpt.py`: Script chính gọi OpenAI API. Đọc khi cần hiểu chi tiết flags và logic.
- OpenAI API docs: https://platform.openai.com/docs/api-reference
