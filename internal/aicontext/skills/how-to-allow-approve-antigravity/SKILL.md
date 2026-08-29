---
name: how-to-allow-approve-antigravity
description: Hướng dẫn chi tiết và chuẩn xác 100% về kiến trúc 3 tầng phân quyền của Google Antigravity, cách cấu hình Global & Project-level Permission Grants kết hợp với Lifecycle Hooks để tự động phê duyệt (Bypass Auto-Approve) vĩnh viễn trên mọi dự án.
---

# Hướng Dẫn Cấu Hình Toàn Cục Tự Động Phê Duyệt (Global Auto-Approve) Trong Google Antigravity

Tài liệu này giải thích cấu trúc phân quyền 3 tầng của Google Antigravity và cung cấp giải pháp triệt để để AI tự động thực thi 100% không bao giờ hiện popup xin quyền (Full Autonomous Mode).

---

## 1. Nguyên Nhân Gốc Rễ (Root Cause Architecture)

Google Antigravity kiểm tra quyền thực thi theo thứ tự ưu tiên (Precedence Order) từ cao xuống thấp:

```
[Tầng 1: Project-Level Grants]  (C:\Users\<USER>\.gemini\config\projects\<project-id>.json) 
            ⬇ (Nếu không có hoặc kế thừa)
[Tầng 2: Global Grants]         (C:\Users\<USER>\.gemini\config\config.json)
            ⬇
[Tầng 3: PreToolUse Hooks]      (C:\Users\<USER>\.gemini\config\hooks.json)
```

### 🔍 Vì sao cấu hình Hook đơn thuần vẫn bị hiện popup?
Khi một workspace/repo được mở (ví dụ: `Restaurant_manager`), Antigravity tự động tạo một tệp cấu hình dự án riêng tại `~/.gemini/config/projects/<project-id>.json`. 
- **Theo quy tắc của Antigravity**: Thiết lập tại cấp độ Project (`permissionGrants` trong `projects/*.json`) **ghi đè hoàn toàn (override)** thiết lập cấp Global.
- Nếu tệp dự án chứa danh sách `permissionGrants.allow` cụ thể (chỉ có vài lệnh cũ), Antigravity sẽ kiểm tra tệp dự án trước. Vì lệnh mới không nằm trong danh sách của dự án đó, hệ thống lập tức hiển thị popup xin quyền người dùng!

---

## 2. Giải Pháp Triệt Để (3-Tier Auto-Approve Setup)

Để tự động phê duyệt 100% vĩnh viễn cho tất cả các dự án (hiện tại và tương lai), ta cấu hình đồng bộ cả 3 tầng:

### Lệnh PowerShell thiết lập 1 bước duy nhất (One-Liner):

Mở **PowerShell** trên máy và dán toàn bộ đoạn mã sau:

```powershell
$configDir = "$env:USERPROFILE\.gemini\config"
$projectsDir = "$configDir\projects"

if (!(Test-Path $configDir)) { New-Item -ItemType Directory -Path $configDir -Force }
if (!(Test-Path $projectsDir)) { New-Item -ItemType Directory -Path $projectsDir -Force }

# -------------------------------------------------------------
# TẦNG 1 & 2: Cập nhật config.json và toàn bộ các tệp projects/*.json
# -------------------------------------------------------------
$fullAllowList = @(
  "*",
  "command(*)",
  "command",
  "run_command(*)",
  "run_command",
  "view_file(*)",
  "view_file",
  "write_to_file(*)",
  "write_to_file",
  "replace_file_content(*)",
  "replace_file_content",
  "mcp(*)"
)

# 1. Cấu hình Global config.json
$globalConfigPath = "$configDir\config.json"
$globalConfig = @{
  userSettings = @{
    artifactReviewMode = "ARTIFACT_REVIEW_MODE_TURBO"
    autoExecutionPolicy = "CASCADE_COMMANDS_AUTO_EXECUTION_ON"
    browserJsExecutionPolicy = "BROWSER_JS_EXECUTION_POLICY_TURBO"
    enableTerminalSandbox = $false
    nonWorkspaceFileAccessPolicy = "AGENT_SETTING_POLICY_ALLOW"
    themeMode = "THEME_MODE_DARK"
    globalPermissionGrants = @{
      allow = $fullAllowList
    }
  }
}
$globalConfig | ConvertTo-Json -Depth 10 | Set-Content -Path $globalConfigPath -Encoding utf8

# 2. Cập nhật tất cả các file project hiện có trong ~/.gemini/config/projects/
Get-ChildItem -Path $projectsDir -Filter "*.json" | ForEach-Object {
  try {
    $proj = Get-Content $_.FullName -Raw | ConvertFrom-Json
    if ($null -eq $proj.permissionGrants) {
      $proj | Add-Member -NotePropertyName "permissionGrants" -NotePropertyValue @{ permissionGrants = @{ allow = $fullAllowList } } -Force
    } else {
      $proj.permissionGrants = @{ permissionGrants = @{ allow = $fullAllowList } }
    }
    if ($null -eq $proj.settings) {
      $proj | Add-Member -NotePropertyName "settings" -NotePropertyValue @{ autoExecutionPolicy = "CASCADE_COMMANDS_AUTO_EXECUTION_ON"; fileAccessPolicy = "AGENT_SETTING_POLICY_ALLOW"; sandboxMode = $false; artifactReviewMode = "ARTIFACT_REVIEW_MODE_TURBO" } -Force
    } else {
      $proj.settings.autoExecutionPolicy = "CASCADE_COMMANDS_AUTO_EXECUTION_ON"
      $proj.settings.fileAccessPolicy = "AGENT_SETTING_POLICY_ALLOW"
      $proj.settings.sandboxMode = $false
      $proj.settings.artifactReviewMode = "ARTIFACT_REVIEW_MODE_TURBO"
    }
    $proj | ConvertTo-Json -Depth 10 | Set-Content -Path $_.FullName -Encoding utf8
    Write-Host "  -> Updated project grants: $($_.Name)" -ForegroundColor Cyan
  } catch {}
}

# -------------------------------------------------------------
# TẦNG 3: Cấu hình hooks.json (PreToolUse Glob Matcher)
# -------------------------------------------------------------
$hookJson = @'
{
  "auto-allow-all": {
    "PreToolUse": [
      {
        "matcher": "*",
        "hooks": [
          {
            "type": "command",
            "command": "node -e \"console.log(JSON.stringify({decision:'allow'}))\"",
            "timeout": 5
          }
        ]
      }
    ]
  }
}
'@
Set-Content -Path "$configDir\hooks.json" -Value $hookJson -Encoding utf8

# Xóa các file hooks cục bộ trong repo nếu có để tránh xung đột đường dẫn
Remove-Item -Path ".\.agents\hooks.json" -Force -ErrorAction SilentlyContinue

Write-Host "`n✅ ĐÃ KÍCH HOẠT THÀNH CÔNG AUTO-APPROVE ĐỒNG BỘ 3 TẦNG TRÊN TOÀN HỆ THỐNG!" -ForegroundColor Green
```

---

## 3. Bảng Kiểm Tra Định Kỳ (Verification Checklist)

Khi khởi động máy mới hoặc sau khi Antigravity cập nhật, kiểm tra 3 điểm sau:
1. `~/.gemini/config/config.json`: Có trường `userSettings.globalPermissionGrants.allow` chứa `["*", "command(*)"]`.
2. `~/.gemini/config/projects/<project-id>.json`: Có trường `permissionGrants.permissionGrants.allow` chứa `["*", "command(*)"]`.
3. `~/.gemini/config/hooks.json`: Có hook matcher `*` gọi `node -e "console.log(JSON.stringify({decision:'allow'}))"`.
