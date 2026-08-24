$ErrorActionPreference = "Stop"
Set-StrictMode -Version Latest

$RepoRoot = Split-Path -Parent $PSScriptRoot
$CanonicalRuntime = Join-Path $RepoRoot "anvien\bin\anvien.exe"
$LauncherRuntime = Join-Path $RepoRoot "anvien-launcher\AnvienLauncher.exe"
$ServerRuntime = Join-Path $RepoRoot "anvien-launcher\server-bundle\anvien-server.exe"
$LauncherBuildScript = Join-Path $RepoRoot "anvien-launcher\build.ps1"

function Assert-NativeSuccess($Step) {
  if ($LASTEXITCODE -ne 0) {
    throw "$Step failed with exit code $LASTEXITCODE"
  }
}

function Write-Step($Step) {
  Write-Host "[full-build] $Step"
}

function Stop-BuildOutputProcess($Path) {
  $Target = [System.IO.Path]::GetFullPath($Path)
  $Holders = @(Get-Process -ErrorAction SilentlyContinue | Where-Object {
    try {
      $_.Path -and [System.IO.Path]::GetFullPath($_.Path).Equals($Target, [System.StringComparison]::OrdinalIgnoreCase)
    } catch {
      $false
    }
  })
  foreach ($Holder in $Holders) {
    Write-Step "stopping output holder PID $($Holder.Id): $Target"
    Stop-Process -Id $Holder.Id -Force -ErrorAction Stop
    Wait-Process -Id $Holder.Id -ErrorAction SilentlyContinue
  }
}

foreach ($BuildOutput in @($CanonicalRuntime, $LauncherRuntime, $ServerRuntime)) {
  Stop-BuildOutputProcess $BuildOutput
}

Push-Location (Join-Path $RepoRoot "anvien")
try {
  Write-Step "build canonical runtime"
  go run -mod=vendor ..\cmd\anvien package build-runtime
  Assert-NativeSuccess "canonical runtime build"
} finally {
  Pop-Location
}

if (-not (Test-Path -LiteralPath $CanonicalRuntime -PathType Leaf)) {
  throw "Canonical runtime is missing after build: $CanonicalRuntime"
}

Write-Step "build canonical launcher assets"
& pwsh -NoLogo -NoProfile -NonInteractive -ExecutionPolicy Bypass -File $LauncherBuildScript
Assert-NativeSuccess "canonical launcher build"

Write-Step "canonical runtime version"
& $CanonicalRuntime version
Assert-NativeSuccess "canonical runtime version"

$GitHead = (& git -C $RepoRoot rev-parse HEAD).Trim()
Assert-NativeSuccess "git rev-parse HEAD"
$RuntimeHash = (Get-FileHash -LiteralPath $CanonicalRuntime -Algorithm SHA256).Hash
Write-Step "Git HEAD: $GitHead"
Write-Step "canonical runtime SHA-256: $RuntimeHash"

Push-Location $RepoRoot
try {
  Write-Step "canonical runtime analyze . --force"
  & $CanonicalRuntime analyze . --force
  Assert-NativeSuccess "canonical runtime analyze . --force"
} finally {
  Pop-Location
}
