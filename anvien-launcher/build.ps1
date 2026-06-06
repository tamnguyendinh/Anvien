$ErrorActionPreference = "Stop"

$LauncherRoot = Split-Path -Parent $MyInvocation.MyCommand.Path
$RepoRoot = Split-Path -Parent $LauncherRoot
$LauncherSourceRoot = Join-Path $LauncherRoot "src"
$ServerSourceRoot = Join-Path $LauncherRoot "server-wrapper"
$ServerBundleRoot = Join-Path $LauncherRoot "server-bundle"
$WebDistRoot = Join-Path $LauncherRoot "web-dist"
$CanonicalCliBinRoot = Join-Path $RepoRoot "anvien\bin"
$CanonicalCliOutPath = Join-Path $CanonicalCliBinRoot "anvien.exe"
$LauncherOutPath = Join-Path $LauncherRoot "AnvienLauncher.exe"
$ServerOutPath = Join-Path $ServerBundleRoot "anvien-server.exe"
$WebRoot = Join-Path $RepoRoot "anvien-web"
$WebBuildRoot = Join-Path $WebRoot "dist"
$NativeRuntimeScript = Join-Path $RepoRoot "scripts\ensure-ladybug-native.ps1"
$LadybugVersion = if ($env:ANVIEN_LADYBUGDB_VERSION) { $env:ANVIEN_LADYBUGDB_VERSION } else { "auto" }

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

function Assert-ChildPath($Parent, $Child) {
  $ParentFull = [System.IO.Path]::GetFullPath($Parent).TrimEnd([System.IO.Path]::DirectorySeparatorChar)
  $ChildFull = [System.IO.Path]::GetFullPath($Child).TrimEnd([System.IO.Path]::DirectorySeparatorChar)
  if (-not $ChildFull.StartsWith($ParentFull, [System.StringComparison]::OrdinalIgnoreCase)) {
    throw "Refusing to delete path outside launcher root: $ChildFull"
  }
}

function Assert-NativeSuccess($Step) {
  if ($LASTEXITCODE -ne 0) {
    throw "$Step failed with exit code $LASTEXITCODE"
  }
}

function Reset-Directory($Path) {
  Assert-ChildPath $LauncherRoot $Path
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

Assert-Command "npm"
Assert-Command "node"

$Go = Resolve-GoCommand
$GoVersion = & $Go version
Assert-NativeSuccess "go version"
Write-Host "[build] using Go: $GoVersion"

$NativeDir = & $NativeRuntimeScript -Version $LadybugVersion -OutputRoot (Join-Path $RepoRoot ".tmp\ladybug-native")
$NativeDir = (Resolve-Path -LiteralPath $NativeDir).Path
Write-Host "[build] using LadybugDB native runtime: $NativeDir"

Push-Location $WebRoot
try {
  npm run build
  Assert-NativeSuccess "npm run build"
} finally {
  Pop-Location
}

Reset-Directory $ServerBundleRoot
New-Item -ItemType Directory -Path $CanonicalCliBinRoot -Force | Out-Null

Push-Location $RepoRoot
try {
  $PreviousCgoEnabled = $env:CGO_ENABLED
  $PreviousCgoCflags = $env:CGO_CFLAGS
  $PreviousCgoLdflags = $env:CGO_LDFLAGS
  $PreviousPath = $env:PATH
  $env:CGO_ENABLED = "1"
  $env:CGO_CFLAGS = "-I$NativeDir"
  $env:CGO_LDFLAGS = "-L$NativeDir -llbug_shared"
  $env:PATH = "$NativeDir;$env:PATH"
  & $Go build -tags ladybugdb -ldflags="-s -w" -o $CanonicalCliOutPath .\cmd\anvien
  Assert-NativeSuccess "go build cmd/anvien"
} finally {
  $env:CGO_ENABLED = $PreviousCgoEnabled
  $env:CGO_CFLAGS = $PreviousCgoCflags
  $env:CGO_LDFLAGS = $PreviousCgoLdflags
  $env:PATH = $PreviousPath
  Pop-Location
}

Copy-FileIfChanged (Join-Path $NativeDir "lbug_shared.dll") $CanonicalCliBinRoot

Push-Location $LauncherSourceRoot
try {
  & $Go build -ldflags="-s -w -H=windowsgui" -o $LauncherOutPath .
  Assert-NativeSuccess "go build launcher"
} finally {
  Pop-Location
}

Push-Location $ServerSourceRoot
try {
  & $Go build -ldflags="-s -w -H=windowsgui" -o $ServerOutPath .
  Assert-NativeSuccess "go build server wrapper"
} finally {
  Pop-Location
}

if (Test-Path -LiteralPath $WebDistRoot) {
  Assert-ChildPath $LauncherRoot $WebDistRoot
  Remove-Item -LiteralPath $WebDistRoot -Recurse -Force
}
Copy-Item -LiteralPath $WebBuildRoot -Destination $WebDistRoot -Recurse -Force

& $LauncherOutPath register
Assert-NativeSuccess "launcher protocol registration"
