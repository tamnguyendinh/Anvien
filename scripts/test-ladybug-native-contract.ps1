param(
  [Parameter(Mandatory = $true)]
  [string]$LaneRoot
)

$ErrorActionPreference = "Stop"
Set-StrictMode -Version Latest

$ScriptRoot = Split-Path -Parent $MyInvocation.MyCommand.Path
$RepoRoot = Split-Path -Parent $ScriptRoot
$RepoTempRoot = Join-Path $RepoRoot ".tmp"
$ProductionHelper = Join-Path $ScriptRoot "ensure-ladybug-native.ps1"
$ProductionNativeDir = Join-Path $RepoRoot "third_party\ladybugdb\v0.19.1\windows-x86_64"

function Assert-StrictChildPath($Parent, $Child, $Label) {
  $ParentFull = [System.IO.Path]::GetFullPath($Parent).TrimEnd([System.IO.Path]::DirectorySeparatorChar, [System.IO.Path]::AltDirectorySeparatorChar)
  $ChildFull = [System.IO.Path]::GetFullPath($Child).TrimEnd([System.IO.Path]::DirectorySeparatorChar, [System.IO.Path]::AltDirectorySeparatorChar)
  $Prefix = $ParentFull + [System.IO.Path]::DirectorySeparatorChar
  if (-not $ChildFull.StartsWith($Prefix, [System.StringComparison]::OrdinalIgnoreCase)) {
    throw "$Label must be a strict child of $ParentFull; received $ChildFull."
  }
  return $ChildFull
}

function New-SyntheticRepository($Root, $Mode) {
  $Scripts = Join-Path $Root "scripts"
  $Native = Join-Path $Root "third_party\ladybugdb\v0.19.1\windows-x86_64"
  New-Item -ItemType Directory -Path $Scripts -Force | Out-Null
  Copy-Item -LiteralPath $ProductionHelper -Destination (Join-Path $Scripts "ensure-ladybug-native.ps1")
  if ($Mode -eq "missing") {
    return
  }
  New-Item -ItemType Directory -Path $Native -Force | Out-Null
  if ($Mode -eq "partial") {
    Copy-Item -LiteralPath (Join-Path $ProductionNativeDir "lbug.h") -Destination $Native
    return
  }
  foreach ($Name in @("lbug.h", "lbug_shared.lib", "lbug_shared.dll")) {
    Copy-Item -LiteralPath (Join-Path $ProductionNativeDir $Name) -Destination $Native
  }
  if ($Mode -eq "extra") {
    [System.IO.File]::WriteAllText((Join-Path $Native "lbug.hpp"), "unauthorized")
  }
}

function Invoke-ExpectedFailure($Script, $AuthorityRoot, $ExpectedText) {
  try {
    & $Script -Version "v0.19.1" -OutputRoot $AuthorityRoot | Out-Null
  } catch {
    if (-not $_.Exception.Message.Contains($ExpectedText)) {
      throw "Expected failure containing '$ExpectedText'; received '$($_.Exception.Message)'."
    }
    return
  }
  throw "Expected helper failure containing '$ExpectedText', but the command succeeded."
}

$LaneRootFull = Assert-StrictChildPath $RepoTempRoot $LaneRoot "Native-contract test lane root"
$TestRoot = Join-Path $LaneRootFull ("helper-contract-" + [Guid]::NewGuid().ToString("N"))
Assert-StrictChildPath $LaneRootFull $TestRoot "Native-contract test root" | Out-Null
New-Item -ItemType Directory -Path $TestRoot -Force | Out-Null

try {
  $ValidRoot = Join-Path $TestRoot "valid"
  New-SyntheticRepository $ValidRoot "valid"
  $ValidScript = Join-Path $ValidRoot "scripts\ensure-ladybug-native.ps1"
  $ValidAuthority = Join-Path $ValidRoot "third_party\ladybugdb"
  $ValidResult = & $ValidScript -Version "v0.19.1" -OutputRoot $ValidAuthority
  $ExpectedValid = [System.IO.Path]::GetFullPath((Join-Path $ValidAuthority "v0.19.1\windows-x86_64"))
  if (-not ([System.IO.Path]::GetFullPath($ValidResult)).Equals($ExpectedValid, [System.StringComparison]::OrdinalIgnoreCase)) {
    throw "Valid helper result was $ValidResult; expected $ExpectedValid."
  }

  $MissingRoot = Join-Path $TestRoot "missing"
  New-SyntheticRepository $MissingRoot "missing"
  Invoke-ExpectedFailure (Join-Path $MissingRoot "scripts\ensure-ladybug-native.ps1") (Join-Path $MissingRoot "third_party\ladybugdb") "native bundle is missing"

  $PartialRoot = Join-Path $TestRoot "partial"
  New-SyntheticRepository $PartialRoot "partial"
  Invoke-ExpectedFailure (Join-Path $PartialRoot "scripts\ensure-ladybug-native.ps1") (Join-Path $PartialRoot "third_party\ladybugdb") "must contain exactly"

  $ExtraRoot = Join-Path $TestRoot "extra"
  New-SyntheticRepository $ExtraRoot "extra"
  Invoke-ExpectedFailure (Join-Path $ExtraRoot "scripts\ensure-ladybug-native.ps1") (Join-Path $ExtraRoot "third_party\ladybugdb") "must contain exactly"

  Invoke-ExpectedFailure $ProductionHelper (Join-Path $TestRoot "alternate-authority") "repository-owned durable root"

  [pscustomobject]@{
    valid = 1
    missingFailedClosed = 1
    partialFailedClosed = 1
    extraFailedClosed = 1
    alternateAuthorityFailedClosed = 1
  } | ConvertTo-Json
} finally {
  $VerifiedTestRoot = Assert-StrictChildPath $LaneRootFull $TestRoot "Native-contract cleanup root"
  if (Test-Path -LiteralPath $VerifiedTestRoot) {
    Remove-Item -LiteralPath $VerifiedTestRoot -Recurse -Force
  }
}
