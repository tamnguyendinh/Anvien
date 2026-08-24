[CmdletBinding()]
param(
    [string]$SourceRoot = (Split-Path -Parent $PSScriptRoot),
    [switch]$Json
)

$ErrorActionPreference = 'Stop'
Set-StrictMode -Version Latest

$ExpectedGoModSha256 = 'C203E25E0A83A583E89A9D8F8E65BC0881052B3075E46915F90A4FB6433BD249'
$ExpectedGoSumSha256 = '5434F39B08F50F1C53A71333E2E971CFCDE0E80BCA7DD028C00E31B9ACCAEC73'
$ExpectedPatchAuthorityPath = 'third_party/go-vendor/patches/tree-sitter-go-v0.25.0-remove-absent-scanner.patch'
$ExpectedPatchSha256 = '43BB195FFF439C7DCD0D057094EC1169304553E357F6D0B2781DC5713516BEEA'
$ExpectedPatchedDestination = 'vendor/github.com/tree-sitter/tree-sitter-go/bindings/go/binding.go'
$ExpectedPatchPreimageBytes = 333
$ExpectedPatchPreimageSha256 = '680E50820F2F7429FEAC88B69DDCEE67735E6E692DFD1F98147109D77D296096'
$ExpectedPatchPostimageBytes = 245
$ExpectedPatchPostimageSha256 = '5A456F352CFEF15430F3FCFB155A3E1BDE38F244E5F8537026553B1F21547713'
$Utf8NoBomStrict = [System.Text.UTF8Encoding]::new($false, $true)
$Utf8NoBom = [System.Text.UTF8Encoding]::new($false)

function Get-StrictRelativePath {
    param(
        [Parameter(Mandatory)][string]$BasePath,
        [Parameter(Mandatory)][string]$ChildPath
    )
    $baseFull = [System.IO.Path]::GetFullPath($BasePath).TrimEnd('\', '/')
    $childFull = [System.IO.Path]::GetFullPath($ChildPath)
    if ($childFull.Equals($baseFull, [System.StringComparison]::OrdinalIgnoreCase)) {
        return '.'
    }
    $prefix = $baseFull + [System.IO.Path]::DirectorySeparatorChar
    if (-not $childFull.StartsWith($prefix, [System.StringComparison]::OrdinalIgnoreCase)) {
        throw "Path is not a child of the expected root: root=$baseFull path=$childFull"
    }
    return $childFull.Substring($prefix.Length)
}

function Get-BytesSha256Hex {
    param([Parameter(Mandatory)][byte[]]$Bytes)
    $sha256 = [System.Security.Cryptography.SHA256]::Create()
    try {
        $hash = $sha256.ComputeHash($Bytes)
    }
    finally {
        $sha256.Dispose()
    }
    return ([System.BitConverter]::ToString($hash)).Replace('-', '')
}

function Sort-InventoryRows {
    param([Parameter(Mandatory)][object[]]$Rows)
    $caseInsensitive = [System.Collections.Generic.HashSet[string]]::new([System.StringComparer]::OrdinalIgnoreCase)
    $sorted = [System.Collections.Generic.SortedDictionary[string, object]]::new([System.StringComparer]::Ordinal)
    foreach ($row in $Rows) {
        $path = [string]$row.path
        Assert-CanonicalRelativePath $path
        if (-not $caseInsensitive.Add($path)) { throw "Duplicate or case-colliding inventory path: $path" }
        if ($sorted.ContainsKey($path)) { throw "Duplicate inventory path: $path" }
        $sorted.Add($path, $row)
    }
    return @($sorted.Values)
}

function Assert-CanonicalRelativePath {
    param([Parameter(Mandatory)][string]$Path)
    if ([string]::IsNullOrWhiteSpace($Path) -or $Path.Contains([char]0) -or $Path.Contains('\') -or
        $Path.StartsWith('/', [System.StringComparison]::Ordinal) -or $Path.StartsWith('//', [System.StringComparison]::Ordinal) -or
        $Path -match '^[A-Za-z]:' -or $Path.EndsWith('/', [System.StringComparison]::Ordinal)) {
        throw "Non-canonical relative path: $Path"
    }
    foreach ($segment in ($Path -split '/')) {
        if ([string]::IsNullOrEmpty($segment) -or $segment -ceq '.' -or $segment -ceq '..') {
            throw "Non-canonical relative path segment in: $Path"
        }
    }
}

function Get-CanonicalStringDigest {
    param([Parameter(Mandatory)][string[]]$Rows)
    $text = if ($Rows.Count -eq 0) { '' } else { ($Rows -join "`n") + "`n" }
    return Get-BytesSha256Hex -Bytes $Utf8NoBom.GetBytes($text)
}

function Sort-OrdinalStrings {
    param([Parameter(Mandatory)][string[]]$Values)
    $copy = [string[]]@($Values)
    [System.Array]::Sort($copy, [System.StringComparer]::Ordinal)
    return $copy
}

function ConvertTo-GoCacheEscapedPath {
    param([Parameter(Mandatory)][string]$Value)
    $builder = [System.Text.StringBuilder]::new()
    foreach ($character in $Value.ToCharArray()) {
        if ($character -ge [char]'A' -and $character -le [char]'Z') {
            [void]$builder.Append('!')
            [void]$builder.Append([char]([int]$character + 32))
        }
        else {
            [void]$builder.Append($character)
        }
    }
    return $builder.ToString()
}

function Get-Sha256 {
    param([Parameter(Mandatory)][string]$LiteralPath)
    return (Get-FileHash -Algorithm SHA256 -LiteralPath $LiteralPath).Hash.ToUpperInvariant()
}

function Read-StrictUtf8Text {
    param([Parameter(Mandatory)][string]$LiteralPath)
    $bytes = [System.IO.File]::ReadAllBytes($LiteralPath)
    if ($bytes.Length -ge 3 -and $bytes[0] -eq 0xEF -and $bytes[1] -eq 0xBB -and $bytes[2] -eq 0xBF) {
        throw "UTF-8 BOM is forbidden: $LiteralPath"
    }
    return $Utf8NoBomStrict.GetString($bytes)
}

function Get-ReparsePoints {
    param([Parameter(Mandatory)][string]$LiteralPath)
    return @(
        Get-Item -LiteralPath $LiteralPath -Force
        Get-ChildItem -LiteralPath $LiteralPath -Force -Recurse
    ) | Where-Object {
        ($_.Attributes -band [System.IO.FileAttributes]::ReparsePoint) -ne 0
    }
}

function Get-InventoryRows {
    param(
        [Parameter(Mandatory)][string]$Root,
        [Parameter(Mandatory)][string[]]$RelativeRoots
    )
    $rows = [System.Collections.Generic.List[object]]::new()
    foreach ($relativeRoot in $RelativeRoots) {
        $absoluteRoot = Join-Path $Root $relativeRoot
        if (-not (Test-Path -LiteralPath $absoluteRoot -PathType Container)) {
            throw "Required directory is absent: $absoluteRoot"
        }
        foreach ($file in @(Get-ChildItem -LiteralPath $absoluteRoot -File -Force -Recurse | Sort-Object FullName)) {
            $relative = (Get-StrictRelativePath -BasePath $Root -ChildPath $file.FullName).Replace('\', '/')
            $rows.Add([pscustomobject][ordered]@{
                path = $relative
                bytes = [long]$file.Length
                sha256 = Get-Sha256 $file.FullName
            })
        }
    }
    return @(Sort-InventoryRows -Rows @($rows))
}

function Get-TreeDigest {
    param([Parameter(Mandatory)][object[]]$Rows)
    $lines = @($Rows | ForEach-Object { "$($_.path)`t$($_.bytes)`t$($_.sha256)" })
    $text = if ($lines.Count -eq 0) { '' } else { ($lines -join "`n") + "`n" }
    return Get-BytesSha256Hex -Bytes $Utf8NoBom.GetBytes($text)
}

function Get-VendorModuleRows {
    param([Parameter(Mandatory)][string]$ModulesText)
    $rows = [System.Collections.Generic.List[object]]::new()
    foreach ($line in ($ModulesText -split "`n")) {
        $trimmed = $line.TrimEnd("`r")
        if ($trimmed -match '^#\s+(\S+)\s+(\S+)(?:\s+=>.*)?$') {
            if ($trimmed -match '\s=>\s') {
                throw "Module replacements are outside the durable vendor contract: $trimmed"
            }
            $rows.Add([pscustomobject][ordered]@{ path = $Matches[1]; version = $Matches[2] })
        }
    }
    return @($rows | Sort-Object path, version -CaseSensitive)
}

$resolvedSourceRoot = [System.IO.Path]::GetFullPath($SourceRoot).TrimEnd('\')
$goModPath = Join-Path $resolvedSourceRoot 'go.mod'
$goSumPath = Join-Path $resolvedSourceRoot 'go.sum'
$vendorRoot = Join-Path $resolvedSourceRoot 'vendor'
$authorityRoot = Join-Path $resolvedSourceRoot 'third_party\go-vendor'
$manifestPath = Join-Path $authorityRoot 'manifest.v1.json'
$modulesPath = Join-Path $vendorRoot 'modules.txt'
$licensesRoot = Join-Path $authorityRoot 'licenses'
$patchesRoot = Join-Path $authorityRoot 'patches'

foreach ($path in @($goModPath, $goSumPath, $manifestPath, $modulesPath)) {
    if (-not (Test-Path -LiteralPath $path -PathType Leaf)) {
        throw "Required vendor authority file is absent: $path"
    }
}
if (-not (Test-Path -LiteralPath $licensesRoot -PathType Container)) {
    throw "Required license authority directory is absent: $licensesRoot"
}
if (-not (Test-Path -LiteralPath $patchesRoot -PathType Container)) {
    throw "Required patch authority directory is absent: $patchesRoot"
}

$reparsePoints = @(
    Get-ReparsePoints $vendorRoot
    Get-ReparsePoints $authorityRoot
)
if ($reparsePoints.Count -ne 0) {
    throw "Vendor authority contains $($reparsePoints.Count) reparse point(s)."
}
$authorityEntries = @(Get-ChildItem -LiteralPath $authorityRoot -Force)
$unauthorizedAuthorityEntries = @($authorityEntries | Where-Object {
    $_.Name -cne 'manifest.v1.json' -and $_.Name -cne 'licenses' -and $_.Name -cne 'patches'
})
if ($authorityEntries.Count -ne 3 -or $unauthorizedAuthorityEntries.Count -ne 0) {
    throw "third_party/go-vendor must contain exactly manifest.v1.json, licenses, and patches; entries=$($authorityEntries.Count), unauthorized=$($unauthorizedAuthorityEntries.Count)"
}

$manifestText = Read-StrictUtf8Text $manifestPath
$manifest = $manifestText | ConvertFrom-Json
if ([int]$manifest.schemaVersion -ne 1) { throw "Unsupported manifest schema: $($manifest.schemaVersion)" }
if ([int]$manifest.closureContractVersion -ne 2) { throw "Unsupported closure contract: $($manifest.closureContractVersion)" }
if ([string]$manifest.authorityKind -cne 'go_vendor_cgo_closure') { throw "Wrong authority kind: $($manifest.authorityKind)" }
if ([string]$manifest.goVersion -cne 'go1.26.3') { throw "Wrong Go version: $($manifest.goVersion)" }

$goModSha256 = Get-Sha256 $goModPath
$goSumSha256 = Get-Sha256 $goSumPath
$modulesSha256 = Get-Sha256 $modulesPath
if ($goModSha256 -cne $ExpectedGoModSha256 -or [string]$manifest.goModSha256 -cne $ExpectedGoModSha256) {
    throw "go.mod identity is outside the fixed R42 authority: actual=$goModSha256 manifest=$($manifest.goModSha256)"
}
if ($goSumSha256 -cne $ExpectedGoSumSha256 -or [string]$manifest.goSumSha256 -cne $ExpectedGoSumSha256) {
    throw "go.sum identity is outside the fixed R42 authority: actual=$goSumSha256 manifest=$($manifest.goSumSha256)"
}
if ($modulesSha256 -cne [string]$manifest.vendorModulesTxtSha256) { throw 'vendor/modules.txt identity does not match manifest.' }

if ([string]$manifest.generation.script -cne 'scripts/materialize-go-vendor.ps1' -or
    [string]$manifest.generation.verifier -cne 'scripts/verify-go-vendor.ps1' -or
    [string]$manifest.generation.command -cne 'pwsh -File scripts/materialize-go-vendor.ps1 -LaneRoot <repo-local-ephemeral-root> [-InputProxy <checksum-verified-ephemeral-file-proxy>]' -or
    [string]$manifest.generation.sourceProxy -cne 'https://proxy.golang.org' -or
    [string]$manifest.generation.sourceFallback -cne 'none' -or
    [string]$manifest.generation.vendorCommand -cne 'go mod vendor -o <ephemeral-output>' -or
    [string]$manifest.generation.scriptSha256 -cnotmatch '^[A-F0-9]{64}$' -or
    [string]$manifest.generation.verifierSha256 -cnotmatch '^[A-F0-9]{64}$') {
    throw 'Manifest generation provenance is invalid.'
}
$actualVerifierSha256 = Get-Sha256 $PSCommandPath
if ([string]$manifest.generation.verifierSha256 -cne $actualVerifierSha256) {
    throw "Verifier identity differs from manifest: actual=$actualVerifierSha256 manifest=$($manifest.generation.verifierSha256)"
}
$materializerPath = Join-Path $resolvedSourceRoot 'scripts\materialize-go-vendor.ps1'
if (Test-Path -LiteralPath $materializerPath -PathType Leaf) {
    $actualMaterializerSha256 = Get-Sha256 $materializerPath
    if ([string]$manifest.generation.scriptSha256 -cne $actualMaterializerSha256) {
        throw "Materializer identity differs from manifest: actual=$actualMaterializerSha256 manifest=$($manifest.generation.scriptSha256)"
    }
}
$sourceInput = $manifest.sourceInput
if ([string]$sourceInput.role -cne 'ephemeral reproducible acquisition input; never runtime, commit, handoff, or recovery authority' -or
    [string]$sourceInput.proxy -cne 'https://proxy.golang.org' -or
    [string]$sourceInput.checksumAuthority -cne 'unchanged go.sum with GOSUMDB=off' -or
    [int]$sourceInput.materializationListModulesIncludingMain -le 1 -or
    [int]$sourceInput.selectedDependencyModules -ne ([int]$sourceInput.materializationListModulesIncludingMain - 1) -or
    [int]$sourceInput.materializationListMissingContentAuthority -lt 0 -or
    [int]$sourceInput.materializationListMissingGoModAuthority -lt 0 -or
    [int]$sourceInput.materializationListActualContentMismatches -ne 0 -or
    [int]$sourceInput.materializationListActualGoModMismatches -ne 0 -or
    [int]$sourceInput.verifiedVendorClosureModules -ne @($manifest.modules).Count -or
    [int]$sourceInput.goSumAuthorizedGraphModules -le 0 -or
    [int]$sourceInput.authorizedGraphGoModChecks -ne [int]$sourceInput.goSumAuthorizedGraphModules -or
    [int]$sourceInput.authorizedGraphGoModMismatches -ne 0 -or
    [int]$sourceInput.contentChecks -ne @($manifest.modules).Count -or
    [int]$sourceInput.goModChecks -ne @($manifest.modules).Count -or
    [int]$sourceInput.vendorMissingContentAuthority -ne 0 -or
    [int]$sourceInput.vendorMissingGoModAuthority -ne 0 -or
    [string]$sourceInput.goNativeModuleVerify -cne 'all modules verified' -or
    [string]$sourceInput.offlineProxyRehydration -cne 'PASS' -or
    [int]$sourceInput.offlineVendorModuleListDifferences -ne 0 -or
    [int]$sourceInput.offlineVendorGoModVerifyRuns -ne 2 -or
    [int]$sourceInput.offlineChecksumMismatches -ne 0 -or
    [int]$sourceInput.fileProxyFiles -le 0 -or
    [long]$sourceInput.fileProxyBytes -le 0 -or
    [string]$sourceInput.fileProxyTreeSha256 -cnotmatch '^[A-F0-9]{64}$') {
    throw 'Manifest source-input proof is incomplete or outside the fixed R42 boundary.'
}

$actualRows = @(Get-InventoryRows -Root $resolvedSourceRoot -RelativeRoots @('vendor', 'third_party\go-vendor\licenses', 'third_party\go-vendor\patches'))
$rawManifestRows = @($manifest.files | ForEach-Object {
    [pscustomobject][ordered]@{
        path = [string]$_.path
        bytes = [long]$_.bytes
        sha256 = ([string]$_.sha256).ToUpperInvariant()
        origin = [string]$_.origin
    }
})
$manifestRows = @(Sort-InventoryRows $rawManifestRows)
if ((@($rawManifestRows | ForEach-Object path) -join "`n") -cne (@($manifestRows | ForEach-Object path) -join "`n")) {
    throw 'Manifest final file rows are not in canonical ordinal order.'
}

$actualByPath = [System.Collections.Generic.Dictionary[string, object]]::new([System.StringComparer]::Ordinal)
foreach ($row in $actualRows) {
    if ($actualByPath.ContainsKey($row.path)) { throw "Duplicate actual path: $($row.path)" }
    $actualByPath[$row.path] = $row
}
$manifestByPath = [System.Collections.Generic.Dictionary[string, object]]::new([System.StringComparer]::Ordinal)
foreach ($row in $manifestRows) {
    if ($manifestByPath.ContainsKey($row.path)) { throw "Duplicate manifest path: $($row.path)" }
    if (-not ($row.path.StartsWith('vendor/', [System.StringComparison]::Ordinal) -or
              $row.path.StartsWith('third_party/go-vendor/licenses/', [System.StringComparison]::Ordinal) -or
              $row.path.StartsWith('third_party/go-vendor/patches/', [System.StringComparison]::Ordinal))) {
        throw "Manifest file row is outside the allowed authority trees: $($row.path)"
    }
    if ($row.bytes -lt 0 -or $row.sha256 -cnotmatch '^[A-F0-9]{64}$' -or
        $row.origin -notin @('standard', 'supplemental', 'patched_postimage', 'license', 'reviewed_patch')) {
        throw "Manifest contains a malformed file inventory row: $($row.path)"
    }
    $manifestByPath[$row.path] = $row
}

$missing = @($manifestRows | Where-Object { -not $actualByPath.ContainsKey($_.path) })
$extra = @($actualRows | Where-Object { -not $manifestByPath.ContainsKey($_.path) })
$mismatches = @($actualRows | Where-Object {
    $expected = $manifestByPath[$_.path]
    $null -ne $expected -and ($_.bytes -ne $expected.bytes -or $_.sha256 -cne $expected.sha256)
})
if ($missing.Count -ne 0 -or $extra.Count -ne 0 -or $mismatches.Count -ne 0) {
    throw "Vendor file inventory mismatch: missing=$($missing.Count), extra=$($extra.Count), hashOrBytes=$($mismatches.Count)"
}

$treeSha256 = Get-TreeDigest $actualRows
if ($treeSha256 -cne [string]$manifest.treeSha256) {
    throw "Vendor authority tree digest does not match manifest: actual=$treeSha256 expected=$($manifest.treeSha256)"
}

$modulesText = Read-StrictUtf8Text $modulesPath
$actualModules = @(Get-VendorModuleRows $modulesText)
$manifestModules = @($manifest.modules | ForEach-Object {
    [pscustomobject][ordered]@{ path = [string]$_.path; version = [string]$_.version; row = $_ }
})
$actualModuleIds = @(Sort-OrdinalStrings @($actualModules | ForEach-Object { "$($_.path)@$($_.version)" }))
$rawManifestModuleIds = @($manifestModules | ForEach-Object { "$($_.path)@$($_.version)" })
$manifestModuleIds = @(Sort-OrdinalStrings $rawManifestModuleIds)
if (($rawManifestModuleIds -join "`n") -cne ($manifestModuleIds -join "`n")) { throw 'Manifest module rows are not in canonical ordinal order.' }
$duplicateManifestModuleIds = @($manifestModuleIds | Group-Object | Where-Object Count -gt 1)
if ($duplicateManifestModuleIds.Count -ne 0) { throw "Manifest contains duplicate module identities: $($duplicateManifestModuleIds.Count)" }
$moduleDiff = @(Compare-Object -ReferenceObject $actualModuleIds -DifferenceObject $manifestModuleIds -CaseSensitive)
if ($moduleDiff.Count -ne 0) { throw "vendor/modules.txt module denominator differs from manifest: $($moduleDiff.Count) difference(s)." }

$standard = $manifest.standardBaseline
if ([string]$standard.pathDomain -cne 'vendor-relative' -or
    [string]$standard.command -cne 'go mod vendor -o <ephemeral-output>' -or
    [string]$standard.goVersion -cne 'go1.26.3' -or
    [int]$standard.runs -ne 2 -or -not [bool]$standard.equal -or [int]$standard.differenceCount -ne 0) {
    throw 'Standard baseline equality contract is invalid.'
}
$rawBaselineRows = @($standard.files | ForEach-Object {
    [pscustomobject][ordered]@{ path = [string]$_.path; bytes = [long]$_.bytes; sha256 = ([string]$_.sha256).ToUpperInvariant() }
})
$baselineRows = @(Sort-InventoryRows $rawBaselineRows)
if ((@($rawBaselineRows | ForEach-Object path) -join "`n") -cne (@($baselineRows | ForEach-Object path) -join "`n")) { throw 'Standard baseline rows are not canonical ordinal.' }
$baselineTreeSha256 = Get-TreeDigest $baselineRows
$baselineBytes = [long](($baselineRows | Measure-Object bytes -Sum).Sum)
if ([int]$standard.moduleCount -ne $actualModuleIds.Count -or
    [string]$standard.moduleListSha256 -cne (Get-CanonicalStringDigest $actualModuleIds) -or
    [int]$standard.firstRun.files -ne $baselineRows.Count -or [int]$standard.secondRun.files -ne $baselineRows.Count -or
    [long]$standard.firstRun.bytes -ne $baselineBytes -or [long]$standard.secondRun.bytes -ne $baselineBytes -or
    [string]$standard.firstRun.treeSha256 -cne $baselineTreeSha256 -or [string]$standard.secondRun.treeSha256 -cne $baselineTreeSha256) {
    throw 'Standard baseline inventory/equality record does not reconstruct.'
}
$standardModuleIds = @(Sort-OrdinalStrings @($standard.modules | ForEach-Object { "$([string]$_.path)@$([string]$_.version)" }))
if (@(Compare-Object -ReferenceObject $actualModuleIds -DifferenceObject $standardModuleIds -CaseSensitive).Count -ne 0) {
    throw 'Standard baseline module list differs from vendor/modules.txt.'
}

$supplemental = $manifest.supplementalClosure
if ([int]$supplemental.algorithmVersion -ne 2 -or [string]$supplemental.pathDomain -cne 'vendor-relative' -or
    [int]$supplemental.runs -ne 2 -or -not [bool]$supplemental.equal -or [int]$supplemental.differenceCount -ne 0) {
    throw 'Supplemental closure equality contract is invalid.'
}
$rawSupplementalRows = @($supplemental.files | ForEach-Object {
    [pscustomobject][ordered]@{
        path = [string]$_.path; modulePath = [string]$_.modulePath; moduleVersion = [string]$_.moduleVersion
        moduleRelativePath = [string]$_.moduleRelativePath; root = [string]$_.root
        bytes = [long]$_.bytes; sha256 = ([string]$_.sha256).ToUpperInvariant()
    }
})
$supplementalRows = @(Sort-InventoryRows $rawSupplementalRows)
if ((@($rawSupplementalRows | ForEach-Object path) -join "`n") -cne (@($supplementalRows | ForEach-Object path) -join "`n")) { throw 'Supplemental file rows are not canonical ordinal.' }
$baselinePathSet = [System.Collections.Generic.HashSet[string]]::new([System.StringComparer]::Ordinal)
foreach ($row in $baselineRows) { [void]$baselinePathSet.Add([string]$row.path) }
foreach ($row in $supplementalRows) {
    if ($baselinePathSet.Contains([string]$row.path)) { throw "Supplemental file overlaps standard baseline: $($row.path)" }
    $moduleID = "$($row.modulePath)@$($row.moduleVersion)"
    if ($moduleID -notin $actualModuleIds -or -not ([string]$row.path).StartsWith(([string]$row.modulePath) + '/', [System.StringComparison]::Ordinal)) {
        throw "Supplemental file has invalid module attribution: $($row.path)"
    }
}
$supplementalTreeSha256 = Get-TreeDigest $supplementalRows
$supplementalBytes = [long](($supplementalRows | Measure-Object bytes -Sum).Sum)
if ([int]$supplemental.fileCount -ne $supplementalRows.Count -or [long]$supplemental.bytes -ne $supplementalBytes -or
    [string]$supplemental.treeSha256 -cne $supplementalTreeSha256 -or [string]$supplemental.secondTreeSha256 -cne $supplementalTreeSha256) {
    throw 'Supplemental file inventory/equality record does not reconstruct.'
}
$rootKeys = [System.Collections.Generic.HashSet[string]]::new([System.StringComparer]::Ordinal)
foreach ($root in @($supplemental.roots)) {
    $rootPath = [string]$root.root
    if ($rootPath -cne '.') { Assert-CanonicalRelativePath $rootPath }
    $rootKey = "$([string]$root.modulePath)@$([string]$root.moduleVersion)`t$rootPath"
    if (-not $rootKeys.Add($rootKey)) { throw "Duplicate supplemental root: $rootKey" }
    $rootRows = @($supplementalRows | Where-Object {
        $_.modulePath -ceq [string]$root.modulePath -and $_.moduleVersion -ceq [string]$root.moduleVersion -and $_.root -ceq $rootPath
    })
    if ([int]$root.files -ne $rootRows.Count -or [long]$root.bytes -ne [long](($rootRows | Measure-Object bytes -Sum).Sum) -or
        [string]$root.treeSha256 -cne (Get-TreeDigest $rootRows)) { throw "Supplemental root inventory mismatch: $rootKey" }
}
if ([int]$supplemental.rootCount -ne $rootKeys.Count) { throw 'Supplemental root denominator mismatch.' }
$affectedModuleIds = @(Sort-OrdinalStrings @($supplemental.roots | ForEach-Object { "$([string]$_.modulePath)@$([string]$_.moduleVersion)" } | Select-Object -Unique))
$recordedAffectedModuleIds = @(Sort-OrdinalStrings @($supplemental.affectedModules | ForEach-Object { [string]$_ }))
if ([int]$supplemental.affectedModuleCount -ne $affectedModuleIds.Count -or
    @(Compare-Object -ReferenceObject $affectedModuleIds -DifferenceObject $recordedAffectedModuleIds -CaseSensitive).Count -ne 0) {
    throw 'Supplemental affected-module denominator mismatch.'
}
foreach ($control in @($supplemental.effectiveDefines)) {
    if (@($control.defines | ForEach-Object { [string]$_ }) -contains 'TREE_SITTER_FEATURE_WASM') { throw 'TREE_SITTER_FEATURE_WASM is active in sealed build controls.' }
}
$inactiveRows = @($supplemental.inactiveConditionalReferences)
$wasmInactive = @($inactiveRows | Where-Object {
    [string]$_.modulePath -ceq 'github.com/tree-sitter/go-tree-sitter' -and [string]$_.resolvedPath -ceq 'src/stdlib-symbols.txt' -and [string]$_.state -ceq 'inactive'
})
if ($wasmInactive.Count -ne 1) { throw "Inactive WASM authority row denominator mismatch: $($wasmInactive.Count)" }
foreach ($edge in @($supplemental.activeSeedTargets) + @($supplemental.recursiveEdges)) {
    $disposition = [string]$edge.disposition
    if ($disposition -notin @('present', 'supplement', 'patch_exception')) { throw "Unresolved/unsupported active closure edge: $disposition" }
    if ($disposition -ceq 'patch_exception') {
        if ([string]$edge.modulePath -cne 'github.com/tree-sitter/tree-sitter-go' -or [string]$edge.resolvedPath -cne 'src/scanner.c') { throw 'Unexpected reviewed-patch include exception.' }
        continue
    }
    $resolvedVendorPath = 'vendor/' + [string]$edge.modulePath + '/' + [string]$edge.resolvedPath
    if (-not $actualByPath.ContainsKey($resolvedVendorPath)) { throw "Active closure edge target is absent: $resolvedVendorPath" }
}

$patches = @($manifest.reviewedPatches)
if ($patches.Count -ne 1) { throw "Reviewed patch cardinality must be exactly one; actual=$($patches.Count)" }
$patch = $patches[0]
if ([string]$patch.authorityPath -cne $ExpectedPatchAuthorityPath -or [string]$patch.patchSha256 -cne $ExpectedPatchSha256 -or
    [string]$patch.modulePath -cne 'github.com/tree-sitter/tree-sitter-go' -or [string]$patch.moduleVersion -cne 'v0.25.0' -or
    [string]$patch.destination -cne $ExpectedPatchedDestination -or
    [int]$patch.preimageBytes -ne $ExpectedPatchPreimageBytes -or [string]$patch.preimageSha256 -cne $ExpectedPatchPreimageSha256 -or
    [int]$patch.postimageBytes -ne $ExpectedPatchPostimageBytes -or [string]$patch.postimageSha256 -cne $ExpectedPatchPostimageSha256 -or
    [int]$patch.grammarBytes -ne 198042 -or [string]$patch.grammarSha256 -cne '0DCC665AA521A1B73A200BE13B4E376780CCA037D764A165F12864BABFE28000' -or [int]$patch.grammarExternals -ne 0 -or
    [int]$patch.parserBytes -ne 1572685 -or [string]$patch.parserSha256 -cne '3DBF6ED1238B5DFCF2BE4D2F2D4CB27A14D34F34D7784ECCCCBFD532FD4A6D85' -or [int]$patch.externalTokenCount -ne 0 -or
    [int]$patch.sourceScannerEntries -ne 0 -or [int]$patch.zipScannerEntries -ne 0 -or
    [int]$patch.occurrenceCount -ne 1 -or [int]$patch.additions -ne 0 -or [int]$patch.deletions -ne 3 -or [int]$patch.byteDelta -ne -88 -or [int]$patch.fuzz -ne 0 -or [int]$patch.offset -ne 0) {
    throw 'Reviewed patch guard record is invalid.'
}
$patchAuthorityFullPath = Join-Path $resolvedSourceRoot $ExpectedPatchAuthorityPath.Replace('/', '\')
if (-not (Test-Path -LiteralPath $patchAuthorityFullPath -PathType Leaf) -or (Get-Sha256 $patchAuthorityFullPath) -cne $ExpectedPatchSha256) { throw 'Reviewed patch authority bytes are invalid.' }
$patchedBinding = $actualByPath[$ExpectedPatchedDestination]
if ($null -eq $patchedBinding -or $patchedBinding.bytes -ne $ExpectedPatchPostimageBytes -or $patchedBinding.sha256 -cne $ExpectedPatchPostimageSha256) { throw 'Patched binding postimage is invalid.' }
$baselinePatchedRow = @($baselineRows | Where-Object path -CEQ $ExpectedPatchedDestination.Substring('vendor/'.Length))
if ($baselinePatchedRow.Count -ne 1 -or $baselinePatchedRow[0].bytes -ne $ExpectedPatchPreimageBytes -or $baselinePatchedRow[0].sha256 -cne $ExpectedPatchPreimageSha256) { throw 'Patched binding baseline preimage is invalid.' }

$supplementalPathSet = [System.Collections.Generic.HashSet[string]]::new([System.StringComparer]::Ordinal)
foreach ($row in $supplementalRows) { [void]$supplementalPathSet.Add([string]$row.path) }
foreach ($row in $baselineRows) {
    $repoPath = 'vendor/' + [string]$row.path
    $finalRow = $manifestByPath[$repoPath]
    if ($null -eq $finalRow) { throw "Final manifest omits standard row: $repoPath" }
    if ($repoPath -ceq $ExpectedPatchedDestination) {
        if ($finalRow.origin -cne 'patched_postimage' -or $finalRow.sha256 -cne $ExpectedPatchPostimageSha256) { throw 'Patched destination origin is invalid.' }
    }
    elseif ($finalRow.origin -cne 'standard' -or $finalRow.bytes -ne $row.bytes -or $finalRow.sha256 -cne $row.sha256) {
        throw "Final standard partition drift: $repoPath"
    }
}
foreach ($row in $supplementalRows) {
    $repoPath = 'vendor/' + [string]$row.path
    $finalRow = $manifestByPath[$repoPath]
    if ($null -eq $finalRow -or $finalRow.origin -cne 'supplemental' -or $finalRow.bytes -ne $row.bytes -or $finalRow.sha256 -cne $row.sha256) {
        throw "Final supplemental partition drift: $repoPath"
    }
}
$actualVendorRows = @($actualRows | Where-Object { $_.path.StartsWith('vendor/', [System.StringComparison]::Ordinal) })
if ($actualVendorRows.Count -ne ($baselineRows.Count + $supplementalRows.Count)) { throw 'Final vendor file denominator does not equal standard plus supplemental partitions.' }
foreach ($row in $manifestRows) {
    if ($row.path.StartsWith('third_party/go-vendor/licenses/', [System.StringComparison]::Ordinal) -and $row.origin -cne 'license') { throw "Invalid license origin: $($row.path)" }
    if ($row.path.StartsWith('third_party/go-vendor/patches/', [System.StringComparison]::Ordinal) -and $row.origin -cne 'reviewed_patch') { throw "Invalid patch origin: $($row.path)" }
}

$goSumEntries = [System.Collections.Generic.Dictionary[string, string]]::new([System.StringComparer]::Ordinal)
foreach ($line in [System.IO.File]::ReadAllLines($goSumPath)) {
    if ([string]::IsNullOrWhiteSpace($line)) { continue }
    $parts = $line -split ' '
    if ($parts.Count -ne 3) { throw "Malformed go.sum line: $line" }
    $sumKey = "$($parts[0]) $($parts[1])"
    if ($goSumEntries.ContainsKey($sumKey)) { throw "Duplicate go.sum row: $sumKey" }
    $goSumEntries.Add($sumKey, $parts[2])
}
$authorizedGoModEntries = @($goSumEntries.Keys | Where-Object { $_.EndsWith('/go.mod', [System.StringComparison]::Ordinal) })
$graphModules = @($sourceInput.graphModules)
if ($authorizedGoModEntries.Count -ne $graphModules.Count -or
    [int]$sourceInput.goSumAuthorizedGraphModules -ne $graphModules.Count -or
    [int]$sourceInput.authorizedGraphGoModChecks -ne $graphModules.Count) {
    throw "Authorized graph module denominator is invalid: goSum=$($authorizedGoModEntries.Count) manifest=$($graphModules.Count)"
}
$graphModuleIds = [System.Collections.Generic.HashSet[string]]::new([System.StringComparer]::Ordinal)
foreach ($graphModule in $graphModules) {
    $graphID = "$([string]$graphModule.path)@$([string]$graphModule.version)"
    if (-not $graphModuleIds.Add($graphID)) { throw "Duplicate graph module identity: $graphID" }
    $expectedGraphGoModSum = $goSumEntries["$([string]$graphModule.path) $([string]$graphModule.version)/go.mod"]
    if ($null -eq $expectedGraphGoModSum -or [string]$graphModule.goModSum -cne $expectedGraphGoModSum) {
        throw "Graph module row does not match unchanged go.sum: $graphID"
    }
}
foreach ($goModKey in $authorizedGoModEntries) {
    $pair = $goModKey.Substring(0, $goModKey.Length - '/go.mod'.Length) -split ' ', 2
    if ($pair.Count -ne 2 -or -not $graphModuleIds.Contains("$($pair[0])@$($pair[1])")) {
        throw "Manifest graph modules omit unchanged go.sum authority: $goModKey"
    }
}

$sourceModules = @($sourceInput.modules)
if ($sourceModules.Count -ne $manifestModules.Count) { throw 'Source module provenance denominator differs from vendored modules.' }
$sourceModuleById = [System.Collections.Generic.Dictionary[string, object]]::new([System.StringComparer]::Ordinal)
foreach ($sourceModule in $sourceModules) {
    $sourceID = "$([string]$sourceModule.path)@$([string]$sourceModule.version)"
    if ($sourceModuleById.ContainsKey($sourceID)) { throw "Duplicate source module provenance row: $sourceID" }
    if ([long]$sourceModule.infoBytes -le 0 -or [string]$sourceModule.infoSha256 -cnotmatch '^[A-F0-9]{64}$' -or
        [long]$sourceModule.modBytes -le 0 -or [string]$sourceModule.modSha256 -cnotmatch '^[A-F0-9]{64}$' -or
        [long]$sourceModule.zipBytes -le 0 -or [string]$sourceModule.zipSha256 -cnotmatch '^[A-F0-9]{64}$') {
        throw "Incomplete source module byte provenance: $sourceID"
    }
    $expectedContent = $goSumEntries["$([string]$sourceModule.path) $([string]$sourceModule.version)"]
    $expectedGoMod = $goSumEntries["$([string]$sourceModule.path) $([string]$sourceModule.version)/go.mod"]
    if ($null -eq $expectedContent -or $null -eq $expectedGoMod -or [string]$sourceModule.contentSum -cne $expectedContent -or [string]$sourceModule.goModSum -cne $expectedGoMod) {
        throw "Source module checksum provenance differs from unchanged go.sum: $sourceID"
    }
    $sourceModuleById.Add($sourceID, $sourceModule)
}
foreach ($supplementalModuleID in $affectedModuleIds) {
    if (-not $sourceModuleById.ContainsKey($supplementalModuleID)) { throw "Supplemental module lacks authenticated source provenance: $supplementalModuleID" }
}

$licenseDispositionFailures = 0
$copiedDispositionCount = 0
$reviewedExceptionCount = 0
$referencedLicensePaths = [System.Collections.Generic.HashSet[string]]::new([System.StringComparer]::Ordinal)
foreach ($module in $manifestModules) {
    $row = $module.row
    $expectedContentSum = $goSumEntries["$($module.path) $($module.version)"]
    $expectedGoModSum = $goSumEntries["$($module.path) $($module.version)/go.mod"]
    if ($null -eq $expectedContentSum -or $null -eq $expectedGoModSum -or
        [string]$row.contentSum -cne [string]$expectedContentSum -or
        [string]$row.goModSum -cne [string]$expectedGoModSum) {
        throw "Module checksum row does not match unchanged go.sum: $($module.path)@$($module.version)"
    }
    $disposition = [string]$row.licenseDisposition
    $licensePaths = @($row.licensePaths | ForEach-Object { [string]$_ })
    if ($disposition -ceq 'copied') {
        $copiedDispositionCount++
        if ($licensePaths.Count -eq 0) { $licenseDispositionFailures++; continue }
        $escapedModuleIdentity = "$(ConvertTo-GoCacheEscapedPath $module.path)@$(ConvertTo-GoCacheEscapedPath $module.version)"
        $expectedLicensePrefix = "licenses/$escapedModuleIdentity/"
        foreach ($licensePath in $licensePaths) {
            if ($licensePath.Contains('\') -or -not $licensePath.StartsWith($expectedLicensePrefix, [System.StringComparison]::Ordinal) -or
                -not $referencedLicensePaths.Add($licensePath)) {
                $licenseDispositionFailures++
                continue
            }
            $fullPath = [System.IO.Path]::GetFullPath((Join-Path $authorityRoot $licensePath.Replace('/', '\')))
            $licensePrefix = [System.IO.Path]::GetFullPath($licensesRoot).TrimEnd('\') + '\'
            if (-not $fullPath.StartsWith($licensePrefix, [System.StringComparison]::OrdinalIgnoreCase) -or
                -not (Test-Path -LiteralPath $fullPath -PathType Leaf)) {
                $licenseDispositionFailures++
            }
        }
    } elseif ($disposition -ceq 'reviewed_exception') {
        $reviewedExceptionCount++
        if ($licensePaths.Count -ne 0 -or [string]::IsNullOrWhiteSpace([string]$row.licenseReview)) {
            $licenseDispositionFailures++
        }
    } else {
        $licenseDispositionFailures++
    }
}
if ($licenseDispositionFailures -ne 0) {
    throw "Incomplete module license disposition count: $licenseDispositionFailures"
}
$actualLicensePaths = @($actualRows | Where-Object {
    $_.path.StartsWith('third_party/go-vendor/licenses/', [System.StringComparison]::Ordinal)
} | ForEach-Object {
    $_.path.Substring('third_party/go-vendor/'.Length)
})
$unreferencedLicensePaths = @($actualLicensePaths | Where-Object { -not $referencedLicensePaths.Contains($_) })
if ($referencedLicensePaths.Count -ne $actualLicensePaths.Count -or $unreferencedLicensePaths.Count -ne 0) {
    throw "License inventory coverage is incomplete: referenced=$($referencedLicensePaths.Count) actual=$($actualLicensePaths.Count) unreferenced=$($unreferencedLicensePaths.Count)"
}
if ([int]$manifest.licenseDisposition.modules -ne $manifestModules.Count -or
    [int]$manifest.licenseDisposition.copied -ne $copiedDispositionCount -or
    [int]$manifest.licenseDisposition.reviewedExceptions -ne $reviewedExceptionCount -or
    $copiedDispositionCount + $reviewedExceptionCount -ne $manifestModules.Count -or
    [int]$manifest.licenseDisposition.incomplete -ne 0) {
    throw 'Manifest license disposition summary is inconsistent.'
}

if (-not [bool]$manifest.finalDeterminism.equal -or [int]$manifest.finalDeterminism.runs -ne 2 -or
    [int]$manifest.finalDeterminism.differenceCount -ne 0) {
    throw 'Manifest does not prove deterministic final generation 2/2.'
}
$actualVendorRelativeRows = @($actualVendorRows | ForEach-Object {
    [pscustomobject][ordered]@{
        path = ([string]$_.path).Substring('vendor/'.Length)
        bytes = [long]$_.bytes
        sha256 = [string]$_.sha256
    }
})
$actualVendorTreeSha256 = Get-TreeDigest $actualVendorRelativeRows
$actualVendorBytes = [long](($actualVendorRelativeRows | Measure-Object bytes -Sum).Sum)
$actualCompleteBytes = [long](($actualRows | Measure-Object bytes -Sum).Sum)
foreach ($run in @($manifest.finalDeterminism.firstRun, $manifest.finalDeterminism.secondRun)) {
    if ([int]$run.vendorFiles -ne $actualVendorRelativeRows.Count -or [long]$run.vendorBytes -ne $actualVendorBytes -or
        [string]$run.vendorTreeSha256 -cne $actualVendorTreeSha256 -or [int]$run.completeFiles -ne $actualRows.Count -or
        [long]$run.completeBytes -ne $actualCompleteBytes -or [string]$run.completeTreeSha256 -cne $treeSha256) {
        throw "Manifest final determinism does not match durable bytes: vendor=$actualVendorTreeSha256 complete=$treeSha256"
    }
}

$result = [ordered]@{
    status = 'PASS'
    sourceRoot = $resolvedSourceRoot
    goModSha256 = $goModSha256
    goSumSha256 = $goSumSha256
    vendorModulesTxtSha256 = $modulesSha256
    moduleCount = $actualModules.Count
    fileCount = $actualRows.Count
    bytes = [long](($actualRows | Measure-Object bytes -Sum).Sum)
    treeSha256 = $treeSha256
    reparsePoints = $reparsePoints.Count
    missingFiles = $missing.Count
    extraFiles = $extra.Count
    mismatchedFiles = $mismatches.Count
    incompleteLicenseDispositions = $licenseDispositionFailures
    standardFiles = $baselineRows.Count
    supplementalFiles = $supplementalRows.Count
    supplementalRoots = $rootKeys.Count
    reviewedPatches = $patches.Count
    finalVendorTreeSha256 = $actualVendorTreeSha256
}

if ($Json) {
    $result | ConvertTo-Json -Depth 5
} else {
    Write-Output "PASS: verified Go vendor authority ($($actualModules.Count) modules, $($actualRows.Count) files, tree $treeSha256)"
}
