---
name: codex-call-gemini
description: Use when delegating reasoning, algorithmic analysis, architecture design, refactoring, or complex code generation to Google Gemini 3.8 Flash (Google Antigravity High Thinking) via native MCP.
---

# Integrate Gemini 3.8 Flash into Codex via Native MCP

## Purpose

Cho phép Codex Desktop và các AI agent ủy quyền trực tiếp các bài toán phân tích sâu, thuật toán phức tạp, hoặc tái cấu trúc mã nguồn cho **Google Gemini 3.8 Flash High Thinking** thông qua giao thức Model Context Protocol (MCP).

---

## 1. Thông số Kỹ thuật & Model Backend

* **MCP Server Name:** `gemini`
* **Tool Name:** `ask_gemini`
* **Transport:** Native `stdio` qua binary `gemini-mcp.exe`
* **Upstream Model:** `gemini-3.8-flash-high`
* **Reasoning Mode:** `High Thinking` (Tích hợp mặc định từ Google Antigravity backend)
* **API Endpoint:** `https://daily-cloudcode-pa.googleapis.com/v1internal:streamGenerateContent?alt=sse`
* **Project ID:** `aicode-consumers`
* **Authentication:** Tự động nạp OAuth token từ Windows Credential Manager (`gemini:antigravity`). Không yêu cầu API key hay đăng nhập thủ công trong Codex.

---

## 2. Cấu hình MCP trong Codex Desktop (One-time Setup)

Được khai báo trong file cấu hình `C:\Users\<USER>\.codex\config.toml`:

```toml
[mcp_servers.gemini]
command = 'e:\add-provider-to-codex\bridge\gemini-mcp.exe'
args = []
startup_timeout_sec = 60
```

---

## 3. Quy cách gọi Tool (Direct MCP Tool Call)

Codex thực hiện gọi tool theo schema sau:

* **ServerName:** `gemini`
* **ToolName:** `ask_gemini`
* **Arguments:**

```json
{
  "prompt": "<Khuôn mẫu chỉ thị kích hoạt Antigravity + Yêu cầu kỹ thuật chi tiết>",
  "context": "<Nội dung file code liên quan, stack trace, hoặc cấu trúc schema của dự án>"
}
```

### Chi tiết các tham số:

1. `prompt` (string, **bắt buộc**):
   * **Khuôn mẫu bắt buộc:** Codex **PHẢI** soạn prompt theo đúng cấu trúc 2 dòng sau để kích hoạt Gemini hoạt động với toàn bộ năng lực của Antigravity:
     ```
     [Môi trường: Google Antigravity | Model: gemini-3.8-flash-high]
     Yêu cầu bạn hoạt động trong Antigravity để thực hiện nhiệm vụ kỹ thuật sau:
     <Mô tả chi tiết bài toán, yêu cầu trả về mã nguồn hoàn chỉnh 100%, không dùng mã giả>
     ```
2. `context` (string, **tùy chọn nhưng khuyến nghị**):
   * Đường dẫn file và nội dung mã nguồn của các file liên quan trích xuất từ repository.
   * Log lỗi compile hoặc stack trace khi debug.
   * *Lưu ý:* `gemini-mcp.exe` không tự truy cập filesystem của Codex; toàn bộ ngữ cảnh cần thiết phải được Codex đọc từ repo và truyền trực tiếp vào tham số này.

---

## 4. Khi nào Codex PHẢI gọi Gemini?

Codex bắt buộc hoặc chủ động gọi `ask_gemini` trong các tình huống:

1. **Thuật toán & Tối ưu hóa:**
   * Dynamic Programming, đồ thị (Dijkstra, Tarjan), hình học tính toán, xử lý chuỗi nâng cao.
   * Tối ưu độ phức tạp thời gian từ $O(N^2)$ xuống $O(N \log N)$ hoặc $O(N)$.
2. **Concurrency & Race Condition:**
   * Phân tích mutex locks, channels, thread safety, deadlock trong Go hoặc Rust.
3. **Thẩm định chéo (Cross-Validation / Second Opinion):**
   * Đánh giá rủi ro trước khi thực hiện refactor kiến trúc lớn.
   * Tìm lỗ hổng bảo mật, tràn bộ nhớ hoặc edge-case tiềm ẩn mà model hiện tại chưa chắc chắn.
4. **Người dùng yêu cầu trực tiếp:**
   * Khi prompt của người dùng có chứa các từ khóa: `gemini`, `hỏi gemini`, `nhờ gemini`, `dùng gemini 3.8`.

---

## 5. Mẫu gọi thực tế (Execution Examples)

### Ví dụ 1: Phân tích Concurrency / Race Condition trong Go

```json
{
  "prompt": "[Môi trường: Google Antigravity | Model: gemini-3.8-flash-high]\nYêu cầu bạn hoạt động trong Antigravity để phân tích và xử lý dứt điểm rủi ro race condition trong Cache. Viết lại mã nguồn hoàn chỉnh sử dụng sync.RWMutex.",
  "context": "File: cache.go\n\ntype Cache struct {\n\tdata map[string]any\n}\n\nfunc (c *Cache) Get(k string) any {\n\treturn c.data[k]\n}\n\nfunc (c *Cache) Set(k string, v any) {\n\tc.data[k] = v\n}"
}
```

### Ví dụ 2: Tối ưu hóa Thuật toán & Cấu trúc Dữ liệu

```json
{
  "prompt": "[Môi trường: Google Antigravity | Model: gemini-3.8-flash-high]\nYêu cầu bạn hoạt động trong Antigravity để viết hàm tìm chu kỳ ngắn nhất trong đồ thị có hướng bằng Go với độ phức tạp tối ưu O(V * E). Kèm unit test đầy đủ.",
  "context": "type Edge struct { To, Weight int }\ntype Graph struct { Nodes map[int][]Edge }"
}
```

---

## 6. Xử lý kết quả trả về (Handling Output)

1. **Tiếp nhận phản hồi:** Kết quả từ `ask_gemini` là văn bản phản hồi hoàn chỉnh từ Gemini 3.8 Flash High Thinking (chứa phân tích lập luận và khối mã nguồn giải pháp).
2. **Thực thi:** Codex đọc kết quả, dùng các công cụ chỉnh sửa file nội bộ (`apply_patch`, `edit_file`) để áp dụng mã nguồn vào dự án.
3. **Kiểm thử:** Luôn chạy lệnh build/test của dự án để nghiệm thu tính chính xác trước khi kết thúc tác vụ.
