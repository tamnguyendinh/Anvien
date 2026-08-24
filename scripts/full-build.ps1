param(
  [string]$LaneRoot = ""
)

$ErrorActionPreference = "Stop"
Set-StrictMode -Version Latest

$RepoRoot = Split-Path -Parent $PSScriptRoot
$RepoTempRoot = Join-Path $RepoRoot ".tmp"
$NativeAuthorityRoot = Join-Path $RepoRoot "third_party\ladybugdb"
$NativeRuntimeScript = Join-Path $RepoRoot "scripts\ensure-ladybug-native.ps1"
$VendorVerifier = Join-Path $RepoRoot "scripts\verify-go-vendor.ps1"
$PinnedLadybugVersion = "v0.19.1"

function Assert-NativeSuccess($Step) {
  if ($LASTEXITCODE -ne 0) {
    throw "$Step failed with exit code $LASTEXITCODE"
  }
}

function Write-Step($Step) {
  Write-Host "[full-build] $Step"
}

function Resolve-RepoPath($Path) {
  if ([System.IO.Path]::IsPathRooted($Path)) {
    return [System.IO.Path]::GetFullPath($Path)
  }
  return [System.IO.Path]::GetFullPath((Join-Path $RepoRoot $Path))
}

function Assert-StrictChildPath($Parent, $Child, $Label) {
  $ParentFull = [System.IO.Path]::GetFullPath($Parent).TrimEnd([System.IO.Path]::DirectorySeparatorChar, [System.IO.Path]::AltDirectorySeparatorChar)
  $ChildFull = [System.IO.Path]::GetFullPath($Child).TrimEnd([System.IO.Path]::DirectorySeparatorChar, [System.IO.Path]::AltDirectorySeparatorChar)
  $Prefix = $ParentFull + [System.IO.Path]::DirectorySeparatorChar
  if (-not $ChildFull.StartsWith($Prefix, [System.StringComparison]::OrdinalIgnoreCase)) {
    throw "$Label must be a strict child of $ParentFull; received $ChildFull."
  }
  return $ChildFull
}

function Get-Sha256Hash($Path) {
  return (Get-FileHash -LiteralPath $Path -Algorithm SHA256).Hash
}

function Write-BuildProvenance($Status, $LaneCommandPath, $CanonicalRuntimePath, $Failure = "") {
  $SourceFiles = @(
    "go.mod",
    "go.sum",
    "vendor\modules.txt",
    "third_party\go-vendor\manifest.v1.json",
    "scripts\verify-go-vendor.ps1",
    "scripts\full-build.ps1",
    "scripts\ensure-ladybug-native.ps1",
    "internal\cli\package_runtime.go",
    "anvien-launcher\build.ps1",
    "internal\graphhealth\diagnostics.go"
  ) | ForEach-Object {
    $Path = Join-Path $RepoRoot $_
    $Item = Get-Item -LiteralPath $Path
    [ordered]@{
      path = $_.Replace("\", "/")
      bytes = $Item.Length
      sha256 = Get-Sha256Hash $Path
    }
  }
  $NativeFiles = @("lbug.h", "lbug_shared.lib", "lbug_shared.dll") | ForEach-Object {
    $Path = Join-Path $NativeDir $_
    $Item = Get-Item -LiteralPath $Path
    [ordered]@{
      name = $_
      bytes = $Item.Length
      sha256 = Get-Sha256Hash $Path
    }
  }
  $RuntimeFiles = @()
  foreach ($Path in @($LaneCommandPath, $CanonicalRuntimePath)) {
    if ($Path -and (Test-Path -LiteralPath $Path -PathType Leaf)) {
      $Item = Get-Item -LiteralPath $Path
      $BuildInfo = @(& go version -m $Path 2>&1 | ForEach-Object { $_.ToString() })
      $BuildInfoExitCode = $LASTEXITCODE
      $RuntimeFiles += [ordered]@{
        path = [System.IO.Path]::GetFullPath($Path)
        bytes = $Item.Length
        sha256 = Get-Sha256Hash $Path
        buildInfoExitCode = $BuildInfoExitCode
        buildInfo = @($BuildInfo)
      }
    }
  }
  $GitHead = (& git -C $RepoRoot rev-parse HEAD).Trim()
  Assert-NativeSuccess "git rev-parse HEAD"
  $GitStatus = @(& git -C $RepoRoot status --porcelain=v2 --untracked-files=all)
  Assert-NativeSuccess "git status --porcelain=v2"
  $Payload = [ordered]@{
    schemaVersion = 1
    status = $Status
    recordedAtUtc = [DateTime]::UtcNow.ToString("o")
    repoRoot = [System.IO.Path]::GetFullPath($RepoRoot)
    gitHead = $GitHead
    gitStatus = @($GitStatus)
    sourceFiles = @($SourceFiles)
    laneRoot = $LaneRootFull
    cacheRoot = $CacheRoot
    runtimeRoot = $RuntimeRoot
    nativeAuthority = [ordered]@{
      version = $PinnedLadybugVersion
      platform = "windows-x86_64"
      path = $NativeDir
      files = @($NativeFiles)
    }
    runtimeFiles = @($RuntimeFiles)
    failure = $Failure
  }
  $Json = $Payload | ConvertTo-Json -Depth 8
  $Encoding = New-Object System.Text.UTF8Encoding($false)
  [System.IO.File]::WriteAllText($ProvenancePath, $Json + [Environment]::NewLine, $Encoding)
  Write-Step "provenance: $ProvenancePath"
}

if ([string]::IsNullOrWhiteSpace($LaneRoot)) {
  $UniqueName = "{0}-{1}-{2}" -f [DateTime]::UtcNow.ToString("yyyyMMddTHHmmssfffZ"), $PID, [Guid]::NewGuid().ToString("N")
  $LaneRoot = Join-Path $RepoTempRoot (Join-Path "full-build" $UniqueName)
}

$LaneRootFull = Assert-StrictChildPath $RepoTempRoot (Resolve-RepoPath $LaneRoot) "Full-build lane root"
$CacheRoot = Assert-StrictChildPath $LaneRootFull (Join-Path $LaneRootFull "cache") "Full-build cache root"
$RuntimeRoot = Assert-StrictChildPath $LaneRootFull (Join-Path $LaneRootFull "runtime") "Full-build runtime root"
if ($CacheRoot.Equals($RuntimeRoot, [System.StringComparison]::OrdinalIgnoreCase)) {
  throw "Full-build cache and runtime roots must be distinct."
}

$HomeRoot = Join-Path $LaneRootFull "home"
$TempRoot = Join-Path $LaneRootFull "temp"
$ProvenancePath = Join-Path $RuntimeRoot "full-build-provenance.json"
$LaneCommandPath = Join-Path $RuntimeRoot "package-runtime\anvien.exe"
$CanonicalRuntime = Join-Path $RepoRoot "anvien\bin\anvien.exe"
foreach ($Path in @(
  $LaneRootFull,
  $CacheRoot,
  $RuntimeRoot,
  $HomeRoot,
  $TempRoot,
  (Join-Path $HomeRoot "AppData\Roaming"),
  (Join-Path $HomeRoot "AppData\Local"),
  (Join-Path $CacheRoot "go-build"),
  (Join-Path $CacheRoot "go-mod"),
  (Join-Path $CacheRoot "go-path"),
  (Join-Path $CacheRoot "go-tmp"),
  (Join-Path $CacheRoot "npm")
)) {
  New-Item -ItemType Directory -Path $Path -Force | Out-Null
}

$NativeDir = & $NativeRuntimeScript -Version $PinnedLadybugVersion -OutputRoot $NativeAuthorityRoot
$NativeDir = (Resolve-Path -LiteralPath $NativeDir).Path
Write-Step "lane root: $LaneRootFull"
Write-Step "cache root: $CacheRoot"
Write-Step "runtime root: $RuntimeRoot"
Write-Step "LadybugDB native authority: $NativeDir"

$EnvironmentNames = @(
  "HOME", "USERPROFILE", "HOMEDRIVE", "HOMEPATH", "APPDATA", "LOCALAPPDATA",
  "TEMP", "TMP", "XDG_CONFIG_HOME", "XDG_CACHE_HOME",
  "GOCACHE", "GOMODCACHE", "GOPATH", "GOTMPDIR", "GOENV", "GOWORK", "GOFLAGS", "GOPROXY", "GOSUMDB", "GOTOOLCHAIN",
  "GOPRIVATE", "GONOPROXY", "GONOSUMDB", "GOINSECURE", "GOVCS",
  "NPM_CONFIG_CACHE", "NPM_CONFIG_OFFLINE",
  "ANVIEN_BUILD_REPO_ROOT", "ANVIEN_BUILD_LANE_ROOT", "ANVIEN_BUILD_CACHE_ROOT", "ANVIEN_BUILD_RUNTIME_ROOT",
  "ANVIEN_LADYBUGDB_NATIVE_DIR", "ANVIEN_LADYBUGDB_VERSION",
  "GIT_CONFIG_COUNT", "GIT_CONFIG_KEY_0", "GIT_CONFIG_VALUE_0", "PATH"
)
$PreviousEnvironment = @{}
foreach ($Name in $EnvironmentNames) {
  $PreviousEnvironment[$Name] = [Environment]::GetEnvironmentVariable($Name, "Process")
}

$HomeDrive = Split-Path -Qualifier $HomeRoot
[Environment]::SetEnvironmentVariable("HOME", $HomeRoot, "Process")
[Environment]::SetEnvironmentVariable("USERPROFILE", $HomeRoot, "Process")
[Environment]::SetEnvironmentVariable("HOMEDRIVE", $HomeDrive, "Process")
[Environment]::SetEnvironmentVariable("HOMEPATH", $HomeRoot.Substring($HomeDrive.Length), "Process")
[Environment]::SetEnvironmentVariable("APPDATA", (Join-Path $HomeRoot "AppData\Roaming"), "Process")
[Environment]::SetEnvironmentVariable("LOCALAPPDATA", (Join-Path $HomeRoot "AppData\Local"), "Process")
[Environment]::SetEnvironmentVariable("TEMP", $TempRoot, "Process")
[Environment]::SetEnvironmentVariable("TMP", $TempRoot, "Process")
[Environment]::SetEnvironmentVariable("XDG_CONFIG_HOME", (Join-Path $HomeRoot ".config"), "Process")
[Environment]::SetEnvironmentVariable("XDG_CACHE_HOME", $CacheRoot, "Process")
[Environment]::SetEnvironmentVariable("GOCACHE", (Join-Path $CacheRoot "go-build"), "Process")
[Environment]::SetEnvironmentVariable("GOMODCACHE", (Join-Path $CacheRoot "go-mod"), "Process")
[Environment]::SetEnvironmentVariable("GOPATH", (Join-Path $CacheRoot "go-path"), "Process")
[Environment]::SetEnvironmentVariable("GOTMPDIR", (Join-Path $CacheRoot "go-tmp"), "Process")
[Environment]::SetEnvironmentVariable("GOENV", "off", "Process")
[Environment]::SetEnvironmentVariable("GOWORK", "off", "Process")
[Environment]::SetEnvironmentVariable("GOFLAGS", "", "Process")
[Environment]::SetEnvironmentVariable("GOPROXY", "off", "Process")
[Environment]::SetEnvironmentVariable("GOSUMDB", "off", "Process")
[Environment]::SetEnvironmentVariable("GOTOOLCHAIN", "local", "Process")
[Environment]::SetEnvironmentVariable("GOPRIVATE", "", "Process")
[Environment]::SetEnvironmentVariable("GONOPROXY", "none", "Process")
[Environment]::SetEnvironmentVariable("GONOSUMDB", "none", "Process")
[Environment]::SetEnvironmentVariable("GOINSECURE", "", "Process")
[Environment]::SetEnvironmentVariable("GOVCS", "*:off", "Process")
[Environment]::SetEnvironmentVariable("NPM_CONFIG_CACHE", (Join-Path $CacheRoot "npm"), "Process")
[Environment]::SetEnvironmentVariable("NPM_CONFIG_OFFLINE", "true", "Process")
[Environment]::SetEnvironmentVariable("ANVIEN_BUILD_REPO_ROOT", $RepoRoot, "Process")
[Environment]::SetEnvironmentVariable("ANVIEN_BUILD_LANE_ROOT", $LaneRootFull, "Process")
[Environment]::SetEnvironmentVariable("ANVIEN_BUILD_CACHE_ROOT", $CacheRoot, "Process")
[Environment]::SetEnvironmentVariable("ANVIEN_BUILD_RUNTIME_ROOT", $RuntimeRoot, "Process")
[Environment]::SetEnvironmentVariable("ANVIEN_LADYBUGDB_NATIVE_DIR", $NativeDir, "Process")
[Environment]::SetEnvironmentVariable("ANVIEN_LADYBUGDB_VERSION", $PinnedLadybugVersion, "Process")
[Environment]::SetEnvironmentVariable("GIT_CONFIG_COUNT", "1", "Process")
[Environment]::SetEnvironmentVariable("GIT_CONFIG_KEY_0", "safe.directory", "Process")
[Environment]::SetEnvironmentVariable("GIT_CONFIG_VALUE_0", "E:/Anvien", "Process")
[Environment]::SetEnvironmentVariable("PATH", "$NativeDir;$($PreviousEnvironment['PATH'])", "Process")

try {
  Write-Step "verify Go vendor authority"
  & $VendorVerifier -SourceRoot $RepoRoot -Json | Out-Host

  Write-BuildProvenance "preflight" $null $null

  Push-Location (Join-Path $RepoRoot "anvien")
  try {
    Write-Step "direct Go package runtime build"
    go run -mod=vendor ..\cmd\anvien package build-runtime
    Assert-NativeSuccess "direct Go package runtime build"

    if (-not (Test-Path -LiteralPath $LaneCommandPath -PathType Leaf)) {
      throw "Lane-staged package runtime is missing: $LaneCommandPath"
    }
    Assert-StrictChildPath $RuntimeRoot $LaneCommandPath "Lane-staged package runtime" | Out-Null

    Write-Step "lane-staged package runtime version"
    & $LaneCommandPath version
    Assert-NativeSuccess "lane-staged package runtime version"

    Write-BuildProvenance "package-runtime-built" $LaneCommandPath $CanonicalRuntime
  } finally {
    Pop-Location
  }

  Push-Location $RepoRoot
  try {
    Write-Step "anvien-launcher build"
    powershell -NoProfile -ExecutionPolicy Bypass -File .\anvien-launcher\build.ps1 -LaneRoot $LaneRootFull -CacheRoot $CacheRoot -RuntimeRoot $RuntimeRoot -NativeDir $NativeDir
    Assert-NativeSuccess "anvien-launcher build"

    Write-Step "canonical runtime version"
    & $CanonicalRuntime version
    Assert-NativeSuccess "canonical runtime version"

    Write-BuildProvenance "runtime-built" $LaneCommandPath $CanonicalRuntime

    Write-Step "canonical runtime analyze . --force"
    & $CanonicalRuntime analyze . --force
    Assert-NativeSuccess "canonical runtime analyze . --force"

    Write-BuildProvenance "complete" $LaneCommandPath $CanonicalRuntime
  } finally {
    Pop-Location
  }
} catch {
  $BuildError = $_
  $FailureLaneCommand = if (Test-Path -LiteralPath $LaneCommandPath -PathType Leaf) { $LaneCommandPath } else { $null }
  $FailureCanonicalRuntime = if ($FailureLaneCommand -and (Test-Path -LiteralPath $CanonicalRuntime -PathType Leaf)) { $CanonicalRuntime } else { $null }
  try {
    Write-BuildProvenance "failed" $FailureLaneCommand $FailureCanonicalRuntime $BuildError.Exception.Message
  } catch {
    Write-Warning "Failed to write full-build failure provenance: $($_.Exception.Message)"
  }
  throw $BuildError
} finally {
  foreach ($Name in $EnvironmentNames) {
    [Environment]::SetEnvironmentVariable($Name, $PreviousEnvironment[$Name], "Process")
  }
}
