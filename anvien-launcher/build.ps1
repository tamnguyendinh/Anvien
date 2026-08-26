$ErrorActionPreference = "Stop"
Set-StrictMode -Version Latest

$LauncherRoot = Split-Path -Parent $MyInvocation.MyCommand.Path
$RepoRoot = Split-Path -Parent $LauncherRoot
$LauncherSourceRoot = Join-Path $LauncherRoot "src"
$ServerSourceRoot = Join-Path $LauncherRoot "server-wrapper"
$ServerBundleRoot = Join-Path $LauncherRoot "server-bundle"
$WebDistRoot = Join-Path $LauncherRoot "web-dist"
$CanonicalCliBinRoot = Join-Path $RepoRoot "anvien\bin"
$CanonicalCliPath = Join-Path $CanonicalCliBinRoot "anvien.exe"
$CanonicalMetadataPath = Join-Path $CanonicalCliBinRoot "anvien-runtime.json"
$LauncherOutPath = Join-Path $LauncherRoot "AnvienLauncher.exe"
$ServerOutPath = Join-Path $ServerBundleRoot "anvien-server.exe"
$WebRoot = Join-Path $RepoRoot "anvien-web"
$WebSourceConfig = Join-Path $WebRoot "vite.config.ts"
$ReadmeSourcePath = Join-Path $RepoRoot "README.md"
$TypeScriptEntry = Join-Path $WebRoot "node_modules\typescript\bin\tsc"
$ViteEntry = Join-Path $WebRoot "node_modules\vite\bin\vite.js"
$NativeRuntimeScript = Join-Path $RepoRoot "scripts\ensure-ladybug-native.ps1"
$NativeAuthorityRoot = Join-Path $RepoRoot "third_party\ladybugdb"
$PinnedLadybugVersion = "v0.19.1"

function Assert-Command($Name) {
  if (-not (Get-Command $Name -ErrorAction SilentlyContinue)) {
    throw "$Name is required to build the packaged launcher."
  }
}

function Resolve-GoCommand {
  if ($env:ANVIEN_GO -and (Test-Path -LiteralPath $env:ANVIEN_GO -PathType Leaf)) {
    return $env:ANVIEN_GO
  }
  $GoCommand = Get-Command "go" -ErrorAction SilentlyContinue
  if ($GoCommand) {
    return $GoCommand.Source
  }
  throw "Go 1.26.3 is required to build the Go launcher runtime."
}

function Assert-NativeSuccess($Step) {
  if ($LASTEXITCODE -ne 0) {
    throw "$Step failed with exit code $LASTEXITCODE"
  }
}

function Get-Sha256Hash($Path) {
  return (Get-FileHash -LiteralPath $Path -Algorithm SHA256).Hash
}

function Copy-FileIfChanged($Source, $DestinationDirectory) {
  if (-not (Test-Path -LiteralPath $Source -PathType Leaf)) {
    throw "Required build input is missing: $Source"
  }
  New-Item -ItemType Directory -Path $DestinationDirectory -Force | Out-Null
  $Destination = Join-Path $DestinationDirectory (Split-Path -Leaf $Source)
  if ((Test-Path -LiteralPath $Destination -PathType Leaf) -and (Get-Sha256Hash $Source) -eq (Get-Sha256Hash $Destination)) {
    Write-Host "[build] up to date: $Destination"
    return
  }
  Copy-Item -LiteralPath $Source -Destination $Destination -Force
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
    Write-Host "[build] stopping output holder PID $($Holder.Id): $Target"
    Stop-Process -Id $Holder.Id -Force -ErrorAction Stop
    Wait-Process -Id $Holder.Id -ErrorAction SilentlyContinue
  }
}

Assert-Command "node"
$Go = Resolve-GoCommand
$Node = (Get-Command "node" -ErrorAction Stop).Source
foreach ($RequiredPath in @($CanonicalCliPath, $CanonicalMetadataPath, $WebSourceConfig, $ReadmeSourcePath, $TypeScriptEntry, $ViteEntry)) {
  if (-not (Test-Path -LiteralPath $RequiredPath -PathType Leaf)) {
    throw "Required build input is missing: $RequiredPath"
  }
}

$NativeDir = & $NativeRuntimeScript -Version $PinnedLadybugVersion -OutputRoot $NativeAuthorityRoot
$NativeDir = (Resolve-Path -LiteralPath $NativeDir).Path
Write-Host "[build] using LadybugDB native authority: $NativeDir"

foreach ($BuildOutput in @($LauncherOutPath, $ServerOutPath)) {
  Stop-BuildOutputProcess $BuildOutput
}
New-Item -ItemType Directory -Path $CanonicalCliBinRoot -Force | Out-Null
New-Item -ItemType Directory -Path $ServerBundleRoot -Force | Out-Null

$EnvironmentNames = @("GOENV", "GOWORK", "GOFLAGS", "GOPROXY", "GOSUMDB", "GOTOOLCHAIN", "PATH")
$PreviousEnvironment = @{}
foreach ($Name in $EnvironmentNames) {
  $PreviousEnvironment[$Name] = [Environment]::GetEnvironmentVariable($Name, "Process")
}

[Environment]::SetEnvironmentVariable("GOENV", "off", "Process")
[Environment]::SetEnvironmentVariable("GOWORK", "off", "Process")
[Environment]::SetEnvironmentVariable("GOFLAGS", "", "Process")
[Environment]::SetEnvironmentVariable("GOPROXY", "off", "Process")
[Environment]::SetEnvironmentVariable("GOSUMDB", "off", "Process")
[Environment]::SetEnvironmentVariable("GOTOOLCHAIN", "local", "Process")
[Environment]::SetEnvironmentVariable("PATH", "$NativeDir;$($PreviousEnvironment['PATH'])", "Process")

try {
  $GoVersion = & $Go version
  Assert-NativeSuccess "go version"
  Write-Host "[build] using Go: $GoVersion"

  Push-Location $WebRoot
  try {
    & $Node $TypeScriptEntry -b
    Assert-NativeSuccess "TypeScript validation"
    & $Node $ViteEntry build --config $WebSourceConfig --outDir $WebDistRoot --emptyOutDir
    Assert-NativeSuccess "Vite Web build"
    Copy-Item -LiteralPath $ReadmeSourcePath -Destination (Join-Path $WebDistRoot "README.md") -Force
  } finally {
    Pop-Location
  }

  Push-Location $LauncherSourceRoot
  try {
    & $Go build -trimpath -ldflags="-s -w -H=windowsgui" -o $LauncherOutPath .
    Assert-NativeSuccess "go build launcher"
  } finally {
    Pop-Location
  }

  Push-Location $ServerSourceRoot
  try {
    & $Go build -trimpath -ldflags="-s -w -H=windowsgui" -o $ServerOutPath .
    Assert-NativeSuccess "go build server wrapper"
  } finally {
    Pop-Location
  }

  Copy-FileIfChanged (Join-Path $NativeDir "lbug_shared.dll") $CanonicalCliBinRoot

  & $LauncherOutPath register
  Assert-NativeSuccess "launcher protocol registration"

  Write-Host "[build] launcher SHA-256: $(Get-Sha256Hash $LauncherOutPath)"
  Write-Host "[build] server SHA-256: $(Get-Sha256Hash $ServerOutPath)"
} finally {
  foreach ($Name in $EnvironmentNames) {
    [Environment]::SetEnvironmentVariable($Name, $PreviousEnvironment[$Name], "Process")
  }
}
