param(
  [string]$Version = "v0.19.1",
  [string]$OutputRoot = ""
)

$ErrorActionPreference = "Stop"
Set-StrictMode -Version Latest

$ScriptRoot = Split-Path -Parent $MyInvocation.MyCommand.Path
$RepoRoot = Split-Path -Parent $ScriptRoot
$PinnedVersion = "v0.19.1"
$PinnedPlatform = "windows-x86_64"
$ExpectedIdentities = @{
  "lbug.h" = @{ Bytes = 79108; SHA256 = "3D5114D0863B3DAB3B28BD2FEC97A52E6CF669213739921A01814A5BBF5525EB" }
  "lbug_shared.lib" = @{ Bytes = 32433956; SHA256 = "B18AAFC0B712DC1C4CB9DD25F76C3828282D7D460627980E3A4B16EFCD98A955" }
  "lbug_shared.dll" = @{ Bytes = 20230656; SHA256 = "20CBD87840483A2053CFF3FC2DB23A86DD802B8915D86509D41A4B709624CDB7" }
}
$ExpectedFiles = @($ExpectedIdentities.Keys | Sort-Object)
$AuthorityRoot = Join-Path $RepoRoot "third_party\ladybugdb"

function Resolve-RepoPath($Path) {
  if ([string]::IsNullOrWhiteSpace($Path)) {
    return [System.IO.Path]::GetFullPath($AuthorityRoot)
  }
  if ([System.IO.Path]::IsPathRooted($Path)) {
    return [System.IO.Path]::GetFullPath($Path)
  }
  return [System.IO.Path]::GetFullPath((Join-Path $RepoRoot $Path))
}

function Assert-ExactNativeBundle($NativeDir) {
  if (-not (Test-Path -LiteralPath $NativeDir -PathType Container)) {
    throw "LadybugDB native bundle is missing: $NativeDir"
  }

  $Items = @(Get-ChildItem -LiteralPath $NativeDir -Force)
  $ActualFiles = @($Items | Where-Object { -not $_.PSIsContainer } | ForEach-Object { $_.Name } | Sort-Object)
  $ExpectedSorted = @($ExpectedFiles | Sort-Object)
  $Differences = @(Compare-Object -ReferenceObject $ExpectedSorted -DifferenceObject $ActualFiles)
  $Directories = @($Items | Where-Object { $_.PSIsContainer })
  if ($Differences.Count -ne 0 -or $Directories.Count -ne 0 -or $Items.Count -ne $ExpectedFiles.Count) {
    $Actual = if ($Items.Count -eq 0) { "<empty>" } else { ($Items.Name | Sort-Object) -join ", " }
    throw "LadybugDB native bundle must contain exactly {$($ExpectedFiles -join ', ')}; found {$Actual} in $NativeDir"
  }

  foreach ($Name in $ExpectedFiles) {
    $Path = Join-Path $NativeDir $Name
    $Item = Get-Item -LiteralPath $Path
    if (($Item.Attributes -band [System.IO.FileAttributes]::ReparsePoint) -ne 0) {
      throw "LadybugDB native input must be a regular durable file, not a reparse point: $Path"
    }
    $Expected = $ExpectedIdentities[$Name]
    $Hash = (Get-FileHash -LiteralPath $Path -Algorithm SHA256).Hash
    if ($Item.Length -ne $Expected.Bytes -or $Hash -ne $Expected.SHA256) {
      throw "LadybugDB native input identity mismatch: $Path"
    }
  }
}

if ($env:PROCESSOR_ARCHITECTURE -notin @("AMD64", "x86_64")) {
  throw "LadybugDB Windows native runtime is wired for x86_64; current architecture is $env:PROCESSOR_ARCHITECTURE."
}

$VersionTag = if ($Version.StartsWith("v")) { $Version } else { "v$Version" }
if ($VersionTag -ne $PinnedVersion) {
  throw "LadybugDB native authority is pinned to $PinnedVersion; requested $VersionTag."
}

$OutputRootFull = Resolve-RepoPath $OutputRoot
$AuthorityRootFull = [System.IO.Path]::GetFullPath($AuthorityRoot)
if (-not $OutputRootFull.Equals($AuthorityRootFull, [System.StringComparison]::OrdinalIgnoreCase)) {
  throw "LadybugDB native authority must be the repository-owned durable root $AuthorityRootFull; requested $OutputRootFull."
}

$NativeDir = Join-Path $OutputRootFull (Join-Path $PinnedVersion $PinnedPlatform)
Assert-ExactNativeBundle $NativeDir
Write-Output ([System.IO.Path]::GetFullPath($NativeDir))
