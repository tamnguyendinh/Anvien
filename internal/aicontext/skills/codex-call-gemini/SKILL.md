---
name: codex-call-gemini
description: Use when delegating reasoning, algorithmic analysis, architecture design, refactoring, or complex code generation to Google Gemini 3.8 Flash (Google Antigravity High Thinking) via native MCP.
---

# Codex Call Gemini via Native MCP (`codex-call-gemini`)

## Purpose

Cho phép **Codex Desktop** và các AI Agent ủy quyền các tác vụ lập trình phức tạp, phân tích giải thuật hóc búa, hoặc thẩm định mã nguồn cho **Google Gemini 3.8 Flash (High Thinking - Google Antigravity)** thông qua công cụ MCP `ask_gemini`.

Gemini 3.8 Flash sở hữu năng lực suy luận sâu (High Thinking Mode) và ngữ cảnh lớn, là đối tác lý tưởng để Codex tham vấn khi gặp các vấn đề vượt quá khả năng xử lý nhanh hoặc cần góc nhìn kiến trúc chuyên sâu.

---

## 1. Khi nào nên gọi Gemini?

Codex nên chủ động kích hoạt công cụ `ask_gemini` trong các trường hợp sau:

* **Thuật toán & Toán học phức tạp:** Quy hoạch động (DP), đồ thị, hình học không gian, tối ưu hóa tổ hợp.
* **Kiến trúc & Phân tích hệ thống:** Đánh giá luồng dữ liệu, phân rã microservice, phân tích deadlock / race condition trong lập trình đồng thời (concurrency).
* **Refactor mã nguồn lớn:** Tái cấu trúc các module phức tạp hoặc chuyển đổi ngôn ngữ/framework.
* **Thẩm định & Bắt lỗi tiềm ẩn (Cross-validation):** Cần một "bộ não thứ hai" độc lập để soi lỗi logic, rò rỉ bộ nhớ (memory leak), hoặc lỗ hổng bảo mật.
* **Khi người dùng yêu cầu rõ ràng:** Người dùng nói *"hỏi Gemini"*, *"nhờ Gemini giải quyết"*, *"dùng Gemini"*.

---

## 2. Đặc tả giao thức gọi công cụ (Tool Call Specification)

* **MCP Server Name:** `gemini`
* **Tool Name:** `ask_gemini`
* **Cơ chế truyền dữ liệu:** `stdio` qua binary `gemini-mcp.exe`

### Cấu trúc tham số (Arguments):

```json
{
  "prompt": "<Mô tả chi tiết câu hỏi, yêu cầu hoặc nhiệm vụ cần Gemini xử lý>",
  "context": "<(Tùy chọn) Đoạn code liên quan, nội dung file hoặc ngữ cảnh dự án>"
}
```

---

## 3. Quy trình thực hiện mẫu của Codex (Best Practices Workflow)

Khi nhận được yêu cầu cần tham vấn Gemini, Codex cần tuân theo 3 bước:

```
[1. Thu thập ngữ cảnh] ──> [2. Gọi ask_gemini] ──> [3. Tiếp nhận & Áp dụng]
  Codex đọc các file          Gửi prompt rõ ràng        Codex tổng hợp câu
  liên quan trong repo        kèm đoạn code vào         trả lời, chỉnh sửa file
                              tham số context           và chạy test kiểm thử
```

### Bước 1: Chuẩn bị dữ liệu
Codex dùng các công cụ đọc file có sẵn để trích xuất đoạn code hoặc interface cần xử lý.

### Bước 2: Gọi Tool `ask_gemini`
Soạn prompt mạch lạc và truyền ngữ cảnh vào `context`:
```json
{
  "prompt": "Hãy phân tích đoạn mã xử lý socket dưới đây và chỉ ra nguy cơ race condition khi có nhiều kết nối đồng thời. Đề xuất bản vá tối ưu bằng Go.",
  "context": "// Nội dung file server.go\npackage main\n..."
}
```

### Bước 3: Đọc kết quả và hành động
Nhận câu trả lời từ Gemini, áp dụng giải pháp vào codebase của dự án và chạy build/test để đảm bảo mã nguồn hoạt động chính xác.

---

## 4. Mẫu lệnh tương tác nhanh cho Người dùng

Người dùng có thể kích hoạt skill này trong Codex Desktop bằng các câu lệnh tự nhiên:

1. *"Dùng Gemini 3.8 Flash thẩm định xem giải thuật này có tối ưu không: [code]"*
2. *"Nhờ Gemini viết giúp tôi hàm xử lý phân trang bằng Go kèm unit test."*
3. *"Hỏi Gemini kiến trúc phù hợp nhất để mở rộng module thanh toán này."*
