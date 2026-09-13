---
name: codex-call-gemini
description: Use when delegating repository tasks to Antigravity via Native Go MCP, instructing Gemini to open an Antigravity agent session to work directly on the repository.
---
# Delegate Tasks to Antigravity via Native Go MCP
## Purpose
Cho phép Codex Desktop ủy quyền tác vụ lập trình và quản lý mã nguồn cho **Antigravity** thực thi trực tiếp trên repository thông qua Native Go MCP `gemini-mcp.exe`.
---
## 1. Cấu hình MCP trong Codex Desktop (Cố định)
Được khai báo trong file `C:\Users\<USER>\.codex\config.toml`:
```toml
[mcp_servers.gemini]
command = 'e:\add-provider-to-codex\bridge\gemini-mcp.exe'
args = []
startup_timeout_sec = 60
```
---
## 2. Quy cách gọi duy nhất qua Native MCP Tool
Khi cần giao việc cho Antigravity, Codex **không tự chạy lệnh terminal** mà gọi công cụ MCP:
* **ServerName:** `gemini`
* **ToolName:** `ask_gemini`
* **Arguments:**
```json
{
  "prompt": "Dùng lệnh: & \"C:\\Users\\TAM NGUYEN\\.gemini\\antigravity\\bin\\agentapi.bat\" new-conversation \"<nội dung>\" để mở phiên\n\nNội dung làm việc trong phiên: <nhập nội dung công việc cụ thể cần Antigravity thực hiện trên repository>"
}
```
---
## 3. Cơ chế Thực thi & Tiếp nhận kết quả
1. **Gửi lệnh qua MCP:** Codex gọi tool `ask_gemini` với cấu trúc prompt ở trên.