---
name: antigravity-call-codex
description: Dùng khi cần giao việc cho Codex tự động thực thi code, sửa file và chạy test trên repo qua native MCP.
---

# Tích hợp Codex vào Antigravity qua Native MCP

## Mục đích

Cho phép Antigravity giao trọn gói nhiệm vụ kỹ thuật cho Codex thực thi trực tiếp trên repository thông qua giao thức MCP (Model Context Protocol).

Codex có đầy đủ công cụ đọc/ghi file và chạy shell, tự động khép kín vòng lặp sửa sai (edit → build → test → pass) mà không cần Antigravity làm trung gian copy-paste code.

---

## 1. Cấu hình MCP Local Global (Thiết lập một lần)

Để MCP Server có hiệu lực trên toàn bộ các workspace của Antigravity trên máy tính, khai báo vào file cấu hình Global của Antigravity:

* **Đường dẫn:** `C:\Users\<USER>\.gemini\config\mcp_config.json` (hoặc `~/.gemini/config/mcp_config.json`)
* **Nội dung cấu hình (Chế độ YOLO - Full quyền thực thi):**

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

* **Xác nhận trạng thái:** Trên giao diện Antigravity, vào **Additional Options (`...`) > MCP Servers** sẽ thấy `codex` ở trạng thái kết nối với 2 tools: `codex` và `codex-reply`.

---

## 2. Cách Agent gọi Codex qua MCP

Antigravity gọi trực tiếp công cụ MCP native (thông qua `call_mcp_tool` hoặc function call gốc):

### A. Khởi tạo một phiên làm việc mới (`Tool: codex`)

* **ServerName:** `codex`
* **ToolName:** `codex`
* **Arguments:**

```json
{
  "prompt": "Mô tả chi tiết nhiệm vụ cần thực hiện",
  "cwd": "<đường_dẫn_thư_mục_gốc_của_repository_đang_làm_việc>"
}
```
