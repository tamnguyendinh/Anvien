$ErrorActionPreference = "Stop"
Set-StrictMode -Version Latest

$RepoRoot = Split-Path -Parent $PSScriptRoot

function Assert-NativeSuccess($Step) {
  if ($LASTEXITCODE -ne 0) {
    throw "$Step failed with exit code $LASTEXITCODE"
  }
}

function Write-Step($Step) {
  Write-Host "[full-build] $Step"
}

Push-Location (Join-Path $RepoRoot "anvien")
try {
  Write-Step "npm install"
  npm install
  Assert-NativeSuccess "npm install"

  Write-Step "npm run build"
  npm run build
  Assert-NativeSuccess "npm run build"

  Write-Step "npm install -g ."
  npm install -g .
  Assert-NativeSuccess "npm install -g ."

  Write-Step "Get-Command anvien"
  Get-Command anvien -ErrorAction Stop | Out-Host

  Write-Step "anvien version"
  anvien version
  Assert-NativeSuccess "anvien version"
} finally {
  Pop-Location
}

Push-Location $RepoRoot
try {
  Write-Step "anvien-launcher build"
  powershell -ExecutionPolicy Bypass -File .\anvien-launcher\build.ps1
  Assert-NativeSuccess "anvien-launcher build"

  Write-Step "anvien version"
  anvien version
  Assert-NativeSuccess "anvien version"

  Write-Step "anvien analyze . --force"
  anvien analyze . --force
  Assert-NativeSuccess "anvien analyze . --force"
} finally {
  Pop-Location
}
