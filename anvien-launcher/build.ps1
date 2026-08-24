param(
  [string]$LaneRoot = "",
  [string]$CacheRoot = "",
  [string]$RuntimeRoot = "",
  [string]$NativeDir = ""
)

$ErrorActionPreference = "Stop"
Set-StrictMode -Version Latest

$LauncherRoot = Split-Path -Parent $MyInvocation.MyCommand.Path
$RepoRoot = Split-Path -Parent $LauncherRoot
$RepoTempRoot = Join-Path $RepoRoot ".tmp"
$LauncherSourceRoot = Join-Path $LauncherRoot "src"
$ServerSourceRoot = Join-Path $LauncherRoot "server-wrapper"
$ServerBundleRoot = Join-Path $LauncherRoot "server-bundle"
$WebDistRoot = Join-Path $LauncherRoot "web-dist"
$CanonicalCliBinRoot = Join-Path $RepoRoot "anvien\bin"
$LauncherOutPath = Join-Path $LauncherRoot "AnvienLauncher.exe"
$WebRoot = Join-Path $RepoRoot "anvien-web"
$WebSourceConfig = Join-Path $WebRoot "vite.config.ts"
$ReadmeSourcePath = Join-Path $RepoRoot "README.md"
$TypeScriptEntry = Join-Path $WebRoot "node_modules\typescript\bin\tsc"
$ViteEntry = Join-Path $WebRoot "node_modules\vite\bin\vite.js"
$NativeRuntimeScript = Join-Path $RepoRoot "scripts\ensure-ladybug-native.ps1"
$VendorVerifier = Join-Path $RepoRoot "scripts\verify-go-vendor.ps1"
$NativeAuthorityRoot = Join-Path $RepoRoot "third_party\ladybugdb"
$PinnedLadybugVersion = "v0.19.1"

function Assert-Command($Name) {
  if (-not (Get-Command $Name -ErrorAction SilentlyContinue)) {
    throw "$Name is required to build the packaged launcher."
  }
}

function Resolve-GoCommand {
  if ($env:ANVIEN_GO -and (Test-Path -LiteralPath $env:ANVIEN_GO)) {
    return $env:ANVIEN_GO
  }

  $SelectedGo = Join-Path $env:USERPROFILE "go\bin\go1.26.3.exe"
  if (Test-Path -LiteralPath $SelectedGo) {
    return $SelectedGo
  }

  $GoCommand = Get-Command "go" -ErrorAction SilentlyContinue
  if ($GoCommand) {
    return $GoCommand.Source
  }

  throw "Go 1.26.3 is required to build the Go launcher runtime. Install it or set ANVIEN_GO."
}

function Resolve-RepoPath($Path) {
  if ([System.IO.Path]::IsPathRooted($Path)) {
    return [System.IO.Path]::GetFullPath($Path)
  }
  return [System.IO.Path]::GetFullPath((Join-Path $RepoRoot $Path))
}

function Get-LauncherRelativePath($BaseDirectory, $TargetPath) {
  $BaseFull = [System.IO.Path]::GetFullPath($BaseDirectory).TrimEnd([System.IO.Path]::DirectorySeparatorChar, [System.IO.Path]::AltDirectorySeparatorChar)
  $TargetFull = [System.IO.Path]::GetFullPath($TargetPath)
  if (-not [System.IO.Path]::GetPathRoot($BaseFull).Equals([System.IO.Path]::GetPathRoot($TargetFull), [System.StringComparison]::OrdinalIgnoreCase)) {
    throw "Launcher paths must share one filesystem root: base=$BaseFull target=$TargetFull"
  }
  [System.Uri]$BaseUri = $BaseFull + [System.IO.Path]::DirectorySeparatorChar
  [System.Uri]$TargetUri = $TargetFull
  $RelativeUri = $BaseUri.MakeRelativeUri($TargetUri)
  if ($RelativeUri.IsAbsoluteUri) {
    throw "Launcher paths must share one filesystem root: base=$BaseFull target=$TargetFull"
  }
  return [System.Uri]::UnescapeDataString($RelativeUri.ToString()).Replace('/', [System.IO.Path]::DirectorySeparatorChar)
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

function Assert-NativeSuccess($Step) {
  if ($LASTEXITCODE -ne 0) {
    throw "$Step failed with exit code $LASTEXITCODE"
  }
}

function Reset-ChildDirectory($Parent, $Path, $Label) {
  Assert-StrictChildPath $Parent $Path $Label | Out-Null
  if (Test-Path -LiteralPath $Path) {
    Remove-Item -LiteralPath $Path -Recurse -Force
  }
  New-Item -ItemType Directory -Path $Path -Force | Out-Null
}

function Get-Sha256Hash($Path) {
  $Stream = [System.IO.File]::OpenRead($Path)
  try {
    $Sha256 = [System.Security.Cryptography.SHA256]::Create()
    try {
      $HashBytes = $Sha256.ComputeHash($Stream)
      return -join ($HashBytes | ForEach-Object { $_.ToString("x2") })
    } finally {
      $Sha256.Dispose()
    }
  } finally {
    $Stream.Dispose()
  }
}

function Copy-FileIfChanged($Source, $DestinationDirectory) {
  if (-not (Test-Path -LiteralPath $Source -PathType Leaf)) {
    throw "Required build output is missing: $Source"
  }
  New-Item -ItemType Directory -Path $DestinationDirectory -Force | Out-Null
  $Destination = Join-Path $DestinationDirectory (Split-Path -Leaf $Source)
  if (Test-Path -LiteralPath $Destination) {
    $SourceHash = Get-Sha256Hash $Source
    $DestinationHash = Get-Sha256Hash $Destination
    if ($SourceHash -eq $DestinationHash) {
      Write-Host "[build] up to date: $Destination"
      return
    }
  }
  Copy-Item -LiteralPath $Source -Destination $DestinationDirectory -Force
}

if ([string]::IsNullOrWhiteSpace($LaneRoot)) {
  $LaneRoot = $env:ANVIEN_BUILD_LANE_ROOT
}
if ([string]::IsNullOrWhiteSpace($LaneRoot)) {
  $UniqueName = "{0}-{1}-{2}" -f [DateTime]::UtcNow.ToString("yyyyMMddTHHmmssfffZ"), $PID, [Guid]::NewGuid().ToString("N")
  $LaneRoot = Join-Path $RepoTempRoot (Join-Path "launcher-build" $UniqueName)
}
$LaneRootFull = Assert-StrictChildPath $RepoTempRoot (Resolve-RepoPath $LaneRoot) "Launcher lane root"

if ([string]::IsNullOrWhiteSpace($CacheRoot)) {
  $CacheRoot = $env:ANVIEN_BUILD_CACHE_ROOT
}
if ([string]::IsNullOrWhiteSpace($CacheRoot)) {
  $CacheRoot = Join-Path $LaneRootFull "cache"
}
$CacheRootFull = Assert-StrictChildPath $LaneRootFull (Resolve-RepoPath $CacheRoot) "Launcher cache root"

if ([string]::IsNullOrWhiteSpace($RuntimeRoot)) {
  $RuntimeRoot = $env:ANVIEN_BUILD_RUNTIME_ROOT
}
if ([string]::IsNullOrWhiteSpace($RuntimeRoot)) {
  $RuntimeRoot = Join-Path $LaneRootFull "runtime"
}
$RuntimeRootFull = Assert-StrictChildPath $LaneRootFull (Resolve-RepoPath $RuntimeRoot) "Launcher runtime root"
if ($CacheRootFull.Equals($RuntimeRootFull, [System.StringComparison]::OrdinalIgnoreCase)) {
  throw "Launcher cache and runtime roots must be distinct."
}

$ResolvedNativeDir = & $NativeRuntimeScript -Version $PinnedLadybugVersion -OutputRoot $NativeAuthorityRoot
$ResolvedNativeDir = (Resolve-Path -LiteralPath $ResolvedNativeDir).Path
if (-not [string]::IsNullOrWhiteSpace($NativeDir)) {
  $RequestedNativeDir = Resolve-RepoPath $NativeDir
  if (-not $RequestedNativeDir.Equals($ResolvedNativeDir, [System.StringComparison]::OrdinalIgnoreCase)) {
    throw "Launcher LadybugDB native authority must be $ResolvedNativeDir; requested $RequestedNativeDir."
  }
}
$NativeDir = $ResolvedNativeDir

$LauncherStageRoot = Join-Path $RuntimeRootFull "launcher"
$CanonicalCliStageRoot = Join-Path $LauncherStageRoot "cli"
$LauncherBinaryStageRoot = Join-Path $LauncherStageRoot "launcher"
$ServerStageRoot = Join-Path $LauncherStageRoot "server-bundle"
$CanonicalCliStagePath = Join-Path $CanonicalCliStageRoot "anvien.exe"
$LauncherStagePath = Join-Path $LauncherBinaryStageRoot "AnvienLauncher.exe"
$ServerStagePath = Join-Path $ServerStageRoot "anvien-server.exe"
$WebStageRoot = Join-Path $LauncherStageRoot "web-dist"
$ViteLaneConfig = Join-Path $LauncherStageRoot "vite.config.mjs"
$ViteCacheRoot = Join-Path $CacheRootFull "vite"
$BuildTempRoot = Join-Path $LaneRootFull "temp"
Reset-ChildDirectory $RuntimeRootFull $LauncherStageRoot "Launcher staging root"
foreach ($Path in @(
  $CacheRootFull,
  $RuntimeRootFull,
  $CanonicalCliStageRoot,
  $LauncherBinaryStageRoot,
  $ServerStageRoot,
  $WebStageRoot,
  $BuildTempRoot,
  (Join-Path $CacheRootFull "go-build"),
  (Join-Path $CacheRootFull "go-mod"),
  (Join-Path $CacheRootFull "go-path"),
  (Join-Path $CacheRootFull "go-tmp"),
  (Join-Path $CacheRootFull "npm"),
  $ViteCacheRoot
)) {
  New-Item -ItemType Directory -Path $Path -Force | Out-Null
}

$ViteConfigImport = (Get-LauncherRelativePath $LauncherStageRoot $WebSourceConfig).Replace("\", "/")
if (-not $ViteConfigImport.StartsWith(".")) {
  $ViteConfigImport = "./$ViteConfigImport"
}
$ViteConfigImportJson = $ViteConfigImport | ConvertTo-Json -Compress
$ViteCacheRootJson = $ViteCacheRoot.Replace("\", "/") | ConvertTo-Json -Compress
$ViteLaneConfigSource = @"
import baseConfig from $ViteConfigImportJson;

export default async function laneConfig(environment) {
  const resolved = typeof baseConfig === "function" ? await baseConfig(environment) : baseConfig;
  const plugins = (resolved.plugins ?? [])
    .flat(Infinity)
    .filter((plugin) => plugin?.name !== "anvien-root-readme");
  return { ...resolved, cacheDir: $ViteCacheRootJson, plugins };
}
"@
[System.IO.File]::WriteAllText($ViteLaneConfig, $ViteLaneConfigSource, [System.Text.UTF8Encoding]::new($false))

$EnvironmentNames = @(
  "GOCACHE", "GOMODCACHE", "GOPATH", "GOTMPDIR", "GOENV", "GOWORK", "GOFLAGS", "GOPROXY", "GOSUMDB", "GOTOOLCHAIN",
  "GOPRIVATE", "GONOPROXY", "GONOSUMDB", "GOINSECURE", "GOVCS",
  "NPM_CONFIG_CACHE", "NPM_CONFIG_OFFLINE", "TEMP", "TMP",
  "ANVIEN_BUILD_REPO_ROOT", "ANVIEN_BUILD_LANE_ROOT", "ANVIEN_BUILD_CACHE_ROOT", "ANVIEN_BUILD_RUNTIME_ROOT",
  "ANVIEN_LADYBUGDB_NATIVE_DIR", "ANVIEN_LADYBUGDB_VERSION",
  "GIT_CONFIG_COUNT", "GIT_CONFIG_KEY_0", "GIT_CONFIG_VALUE_0", "PATH"
)
$PreviousEnvironment = @{}
foreach ($Name in $EnvironmentNames) {
  $PreviousEnvironment[$Name] = [Environment]::GetEnvironmentVariable($Name, "Process")
}

[Environment]::SetEnvironmentVariable("GOCACHE", (Join-Path $CacheRootFull "go-build"), "Process")
[Environment]::SetEnvironmentVariable("GOMODCACHE", (Join-Path $CacheRootFull "go-mod"), "Process")
[Environment]::SetEnvironmentVariable("GOPATH", (Join-Path $CacheRootFull "go-path"), "Process")
[Environment]::SetEnvironmentVariable("GOTMPDIR", (Join-Path $CacheRootFull "go-tmp"), "Process")
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
[Environment]::SetEnvironmentVariable("NPM_CONFIG_CACHE", (Join-Path $CacheRootFull "npm"), "Process")
[Environment]::SetEnvironmentVariable("NPM_CONFIG_OFFLINE", "true", "Process")
[Environment]::SetEnvironmentVariable("TEMP", $BuildTempRoot, "Process")
[Environment]::SetEnvironmentVariable("TMP", $BuildTempRoot, "Process")
[Environment]::SetEnvironmentVariable("ANVIEN_BUILD_REPO_ROOT", $RepoRoot, "Process")
[Environment]::SetEnvironmentVariable("ANVIEN_BUILD_LANE_ROOT", $LaneRootFull, "Process")
[Environment]::SetEnvironmentVariable("ANVIEN_BUILD_CACHE_ROOT", $CacheRootFull, "Process")
[Environment]::SetEnvironmentVariable("ANVIEN_BUILD_RUNTIME_ROOT", $RuntimeRootFull, "Process")
[Environment]::SetEnvironmentVariable("ANVIEN_LADYBUGDB_NATIVE_DIR", $NativeDir, "Process")
[Environment]::SetEnvironmentVariable("ANVIEN_LADYBUGDB_VERSION", $PinnedLadybugVersion, "Process")
[Environment]::SetEnvironmentVariable("GIT_CONFIG_COUNT", "1", "Process")
[Environment]::SetEnvironmentVariable("GIT_CONFIG_KEY_0", "safe.directory", "Process")
[Environment]::SetEnvironmentVariable("GIT_CONFIG_VALUE_0", "E:/Anvien", "Process")
[Environment]::SetEnvironmentVariable("PATH", "$NativeDir;$($PreviousEnvironment['PATH'])", "Process")

try {
  Assert-Command "node"
  foreach ($RequiredWebTool in @($WebSourceConfig, $ReadmeSourcePath, $TypeScriptEntry, $ViteEntry)) {
    if (-not (Test-Path -LiteralPath $RequiredWebTool -PathType Leaf)) {
      throw "Required offline Web build tool is missing: $RequiredWebTool"
    }
  }

  Write-Host "[build] verifying Go vendor authority"
  & $VendorVerifier -SourceRoot $RepoRoot -Json | Out-Host

  $Go = Resolve-GoCommand
  $Node = (Get-Command "node" -ErrorAction Stop).Source
  $GoVersion = & $Go version
  Assert-NativeSuccess "go version"
  Write-Host "[build] using Go: $GoVersion"
  Write-Host "[build] lane root: $LaneRootFull"
  Write-Host "[build] cache root: $CacheRootFull"
  Write-Host "[build] runtime root: $RuntimeRootFull"
  Write-Host "[build] using LadybugDB native authority: $NativeDir"

  Push-Location $WebRoot
  try {
    & $Node $TypeScriptEntry -p .\tsconfig.app.json --noEmit --incremental false
    Assert-NativeSuccess "TypeScript app validation"
    & $Node $TypeScriptEntry -p .\tsconfig.node.json --noEmit --tsBuildInfoFile (Join-Path $CacheRootFull "tsconfig.node.tsbuildinfo")
    Assert-NativeSuccess "TypeScript build-config validation"
    & $Node $ViteEntry build --config $ViteLaneConfig --outDir $WebStageRoot --emptyOutDir
    Assert-NativeSuccess "Vite Web build"
    Copy-Item -LiteralPath $ReadmeSourcePath -Destination (Join-Path $WebStageRoot "README.md") -Force
  } finally {
    Pop-Location
  }

  Push-Location $RepoRoot
  try {
    $PreviousCgoEnabled = $env:CGO_ENABLED
    $PreviousCgoCflags = $env:CGO_CFLAGS
    $PreviousCgoLdflags = $env:CGO_LDFLAGS
    $env:CGO_ENABLED = "1"
    $env:CGO_CFLAGS = "-I$NativeDir"
    $env:CGO_LDFLAGS = "-L$NativeDir -llbug_shared"
    & $Go build -mod=vendor -tags ladybugdb -trimpath -ldflags="-s -w" -o $CanonicalCliStagePath .\cmd\anvien
    Assert-NativeSuccess "go build cmd/anvien"
  } finally {
    $env:CGO_ENABLED = $PreviousCgoEnabled
    $env:CGO_CFLAGS = $PreviousCgoCflags
    $env:CGO_LDFLAGS = $PreviousCgoLdflags
    Pop-Location
  }

  Push-Location $LauncherSourceRoot
  try {
    & $Go build -trimpath -ldflags="-s -w -H=windowsgui" -o $LauncherStagePath .
    Assert-NativeSuccess "go build launcher"
  } finally {
    Pop-Location
  }

  Push-Location $ServerSourceRoot
  try {
    & $Go build -trimpath -ldflags="-s -w -H=windowsgui" -o $ServerStagePath .
    Assert-NativeSuccess "go build server wrapper"
  } finally {
    Pop-Location
  }

  Copy-FileIfChanged $CanonicalCliStagePath $CanonicalCliBinRoot
  Copy-FileIfChanged (Join-Path $NativeDir "lbug_shared.dll") $CanonicalCliBinRoot
  Copy-FileIfChanged $LauncherStagePath $LauncherRoot
  Reset-ChildDirectory $LauncherRoot $ServerBundleRoot "Canonical server bundle"
  Copy-FileIfChanged $ServerStagePath $ServerBundleRoot

  if (Test-Path -LiteralPath $WebDistRoot) {
    Assert-StrictChildPath $LauncherRoot $WebDistRoot "Canonical launcher web distribution" | Out-Null
    Remove-Item -LiteralPath $WebDistRoot -Recurse -Force
  }
  Copy-Item -LiteralPath $WebStageRoot -Destination $WebDistRoot -Recurse -Force

  & $LauncherOutPath register
  Assert-NativeSuccess "launcher protocol registration"
} finally {
  foreach ($Name in $EnvironmentNames) {
    [Environment]::SetEnvironmentVariable($Name, $PreviousEnvironment[$Name], "Process")
  }
}
