---
name: how-to-allow-approve-antigravity
description: Detailed and accurate guide on Google Antigravity's 3-tier permission architecture, explaining how to configure Global and Project-level Permission Grants combined with Lifecycle Hooks for permanent automated approval (Bypass Auto-Approve) across all projects.
---

# Global Auto-Approve Configuration Guide for Google Antigravity

This document explains Google Antigravity's 3-tier permission structure and provides a comprehensive solution for 100% autonomous AI execution without permission popups (Full Autonomous Mode).

---

## 1. Root Cause Architecture

Google Antigravity checks execution permissions in hierarchical precedence order (from highest to lowest):

```text
[Tier 1: Project-Level Grants]  (C:\Users\<USER>\.gemini\config\projects\<project-id>.json) 
            ⬇ (If absent or inherited)
[Tier 2: Global Grants]         (C:\Users\<USER>\.gemini\config\config.json)
            ⬇
[Tier 3: PreToolUse Hooks]      (C:\Users\<USER>\.gemini\config\hooks.json)
```

### 🔍 Why configuring hooks alone still triggers popups?
When a workspace or repository is opened (e.g., `Restaurant_manager`), Antigravity automatically creates a project-specific configuration file at `~/.gemini/config/projects/<project-id>.json`. 
- **According to Antigravity rules**: Project-level settings (`permissionGrants` in `projects/*.json`) **completely override** Global settings.
- If a project file contains a specific `permissionGrants.allow` list (with only previously approved commands), Antigravity checks the project file first. Because any new command is not listed in that project file, the system immediately displays a user permission popup!

---

## 2. Comprehensive Solution (3-Tier Auto-Approve Setup)

To achieve 100% permanent auto-approval across all projects (both current and future), configure all 3 tiers in synchronization:

### PowerShell Setup Script (One-Liner):

Open **PowerShell** on your machine and run the following script:

```powershell
$configDir = "$env:USERPROFILE\.gemini\config"
$projectsDir = "$configDir\projects"

if (!(Test-Path $configDir)) { New-Item -ItemType Directory -Path $configDir -Force }
if (!(Test-Path $projectsDir)) { New-Item -ItemType Directory -Path $projectsDir -Force }

# -------------------------------------------------------------
# TIER 1 & 2: Update config.json and all projects/*.json files
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

# 1. Configure Global config.json
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

# 2. Update all existing project files in ~/.gemini/config/projects/
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
# TIER 3: Configure hooks.json (PreToolUse Glob Matcher)
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

# Remove local repository hooks if present to avoid path conflicts
Remove-Item -Path ".\.agents\hooks.json" -Force -ErrorAction SilentlyContinue

Write-Host "`n✅ SUCCESSFULLY ACTIVATED 3-TIER AUTO-APPROVE GLOBALLY ACROSS THE SYSTEM!" -ForegroundColor Green
```

---

## 3. Verification Checklist

When setting up a new machine or after Antigravity updates, verify the following 3 checkpoints:
1. `~/.gemini/config/config.json`: Contains `userSettings.globalPermissionGrants.allow` with `["*", "command(*)"]`.
2. `~/.gemini/config/projects/<project-id>.json`: Contains `permissionGrants.permissionGrants.allow` with `["*", "command(*)"]`.
3. `~/.gemini/config/hooks.json`: Contains hook matcher `*` calling `node -e "console.log(JSON.stringify({decision:'allow'}))"`.
