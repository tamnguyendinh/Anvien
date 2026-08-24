[CmdletBinding()]
param(
    [Parameter(Mandatory)][string]$LaneRoot
)

$ErrorActionPreference = 'Stop'
Set-StrictMode -Version Latest

$ScriptRoot = Split-Path -Parent $MyInvocation.MyCommand.Path
$RepoRoot = Split-Path -Parent $ScriptRoot
$AuthorityRepoRoot = $RepoRoot
$RepoTempRoot = Join-Path $RepoRoot '.tmp'
$ProductionVerifier = Join-Path $ScriptRoot 'verify-go-vendor.ps1'
$ProductionMaterializer = Join-Path $ScriptRoot 'materialize-go-vendor.ps1'
$Utf8NoBom = [System.Text.UTF8Encoding]::new($false)
$Utf8NoBomStrict = [System.Text.UTF8Encoding]::new($false, $true)
$Results = [System.Collections.Generic.List[object]]::new()
$TrackedJunctions = [System.Collections.Generic.List[string]]::new()

function Assert-ContractChildPath {
    param([Parameter(Mandatory)][string]$Parent, [Parameter(Mandatory)][string]$Child, [Parameter(Mandatory)][string]$Label)
    $parentFull = [System.IO.Path]::GetFullPath($Parent).TrimEnd('\', '/')
    $childFull = [System.IO.Path]::GetFullPath($Child).TrimEnd('\', '/')
    $prefix = $parentFull + [System.IO.Path]::DirectorySeparatorChar
    if (-not $childFull.StartsWith($prefix, [System.StringComparison]::OrdinalIgnoreCase)) {
        throw "$Label must be a strict child of $parentFull; received $childFull."
    }
    return $childFull
}

function Add-ContractResult {
    param([Parameter(Mandatory)][string]$Suite, [Parameter(Mandatory)][string]$Name, [Parameter(Mandatory)][string]$Proof)
    $Results.Add([pscustomobject][ordered]@{ suite = $Suite; name = $Name; status = 'PASS'; proof = $Proof })
}

function Copy-ContractTree {
    param([Parameter(Mandatory)][string]$Source, [Parameter(Mandatory)][string]$Destination)
    $attributes = [System.IO.File]::GetAttributes($Source)
    if (($attributes -band [System.IO.FileAttributes]::ReparsePoint) -ne 0) {
        throw "Contract fixture source is a reparse point: $Source"
    }
    if (($attributes -band [System.IO.FileAttributes]::Directory) -ne 0) {
        [System.IO.Directory]::CreateDirectory($Destination) | Out-Null
        foreach ($entry in [System.IO.Directory]::EnumerateFileSystemEntries($Source)) {
            Copy-ContractTree $entry (Join-Path $Destination ([System.IO.Path]::GetFileName($entry)))
        }
        return
    }
    [System.IO.Directory]::CreateDirectory((Split-Path -Parent $Destination)) | Out-Null
    [System.IO.File]::Copy($Source, $Destination, $true)
}

function Write-ContractText {
    param([Parameter(Mandatory)][string]$Path, [Parameter(Mandatory)][AllowEmptyString()][string]$Text)
    [System.IO.Directory]::CreateDirectory((Split-Path -Parent $Path)) | Out-Null
    $normalized = $Text.Replace("`r`n", "`n").Replace("`r", "`n")
    [System.IO.File]::WriteAllText($Path, $normalized, $Utf8NoBom)
}

function Add-ContractSuffix {
    param([Parameter(Mandatory)][byte[]]$Bytes, [Parameter(Mandatory)][string]$Suffix)
    $suffixBytes = $Utf8NoBom.GetBytes($Suffix)
    $combined = New-Object byte[] ($Bytes.Length + $suffixBytes.Length)
    [System.Array]::Copy($Bytes, 0, $combined, 0, $Bytes.Length)
    [System.Array]::Copy($suffixBytes, 0, $combined, $Bytes.Length, $suffixBytes.Length)
    return [byte[]]$combined
}

function New-PackagedAuthorityFixture {
    param([Parameter(Mandatory)][string]$Root)
    [System.IO.Directory]::CreateDirectory($Root) | Out-Null
    foreach ($relative in @('go.mod', 'go.sum', 'scripts/verify-go-vendor.ps1', 'vendor', 'third_party/go-vendor')) {
        Copy-ContractTree (Join-Path $RepoRoot $relative.Replace('/', '\')) (Join-Path $Root $relative.Replace('/', '\'))
    }
}

function Invoke-ContractVerifier {
    param([Parameter(Mandatory)][string]$Root)
    $verifier = Join-Path $Root 'scripts\verify-go-vendor.ps1'
    if (-not (Test-Path -LiteralPath $verifier -PathType Leaf)) {
        throw "Packaged Go source verifier is absent: $verifier"
    }
    $raw = & $verifier -SourceRoot $Root -Json | Out-String
    return ($raw | ConvertFrom-Json)
}

function Invoke-ExpectedVerifierFailure {
    param(
        [Parameter(Mandatory)][string]$Name,
        [Parameter(Mandatory)][string]$Root,
        [Parameter(Mandatory)][string]$Expected
    )
    $unexpected = "Verifier unexpectedly accepted contract case: $Name"
    try {
        Invoke-ContractVerifier $Root | Out-Null
        throw $unexpected
    }
    catch {
        $actual = [string]$_.Exception.Message
        if ($actual -ceq $unexpected) { throw $actual }
        if (-not $actual.Contains($Expected)) {
            throw "Verifier case '$Name' expected '$Expected'; received '$actual'."
        }
        Add-ContractResult 'verifier' $Name $Expected
    }
}

function Invoke-FileMutationCase {
    param(
        [Parameter(Mandatory)][string]$Name,
        [Parameter(Mandatory)][string]$Root,
        [Parameter(Mandatory)][string]$RelativePath,
        [Parameter(Mandatory)][ValidateSet('missing', 'tamper')][string]$Mode,
        [Parameter(Mandatory)][string]$Expected
    )
    $path = Join-Path $Root $RelativePath.Replace('/', '\')
    if (-not (Test-Path -LiteralPath $path -PathType Leaf)) { throw "Mutation target is absent: $path" }
    $original = [System.IO.File]::ReadAllBytes($path)
    try {
        if ($Mode -ceq 'missing') { [System.IO.File]::Delete($path) }
        else { [System.IO.File]::WriteAllBytes($path, (Add-ContractSuffix $original "`ncontract-tamper")) }
        Invoke-ExpectedVerifierFailure $Name $Root $Expected
    }
    finally {
        [System.IO.Directory]::CreateDirectory((Split-Path -Parent $path)) | Out-Null
        [System.IO.File]::WriteAllBytes($path, $original)
    }
}

function Invoke-ExtraFileCase {
    param(
        [Parameter(Mandatory)][string]$Name,
        [Parameter(Mandatory)][string]$Root,
        [Parameter(Mandatory)][string]$RelativePath,
        [Parameter(Mandatory)][string]$Expected
    )
    $path = Join-Path $Root $RelativePath.Replace('/', '\')
    try {
        Write-ContractText $path "contract-extra`n"
        Invoke-ExpectedVerifierFailure $Name $Root $Expected
    }
    finally {
        if (Test-Path -LiteralPath $path -PathType Leaf) { [System.IO.File]::Delete($path) }
    }
}

function Invoke-ManifestMutationCase {
    param(
        [Parameter(Mandatory)][string]$Name,
        [Parameter(Mandatory)][string]$Root,
        [Parameter(Mandatory)][scriptblock]$Mutation,
        [Parameter(Mandatory)][string]$Expected
    )
    $manifestPath = Join-Path $Root 'third_party\go-vendor\manifest.v1.json'
    $original = [System.IO.File]::ReadAllBytes($manifestPath)
    try {
        $manifest = $Utf8NoBomStrict.GetString($original) | ConvertFrom-Json
        & $Mutation $manifest
        $json = ($manifest | ConvertTo-Json -Depth 100).Replace("`r`n", "`n").Replace("`r", "`n") + "`n"
        [System.IO.File]::WriteAllText($manifestPath, $json, $Utf8NoBom)
        Invoke-ExpectedVerifierFailure $Name $Root $Expected
    }
    finally {
        [System.IO.File]::WriteAllBytes($manifestPath, $original)
    }
}

function Invoke-ReparseCase {
    param([Parameter(Mandatory)][string]$Root)
    $target = Join-Path $TestRoot 'junction-target'
    $junction = Join-Path $Root 'vendor\zz-contract-reparse'
    [System.IO.Directory]::CreateDirectory($target) | Out-Null
    Write-ContractText (Join-Path $target 'target.txt') "target`n"
    try {
        New-Item -ItemType Junction -Path $junction -Target $target | Out-Null
        $TrackedJunctions.Add($junction)
        Invoke-ExpectedVerifierFailure 'reparse_point' $Root 'Vendor authority contains 1 reparse point(s).'
    }
    finally {
        if (Test-Path -LiteralPath $junction) {
            [System.IO.Directory]::Delete($junction)
        }
        [void]$TrackedJunctions.Remove($junction)
    }
}

function Import-MaterializerFunctions {
    $tokens = $null
    $errors = $null
    $ast = [System.Management.Automation.Language.Parser]::ParseFile($ProductionMaterializer, [ref]$tokens, [ref]$errors)
    if (@($errors).Count -ne 0) { throw "Materializer AST parse failed: $($errors[0].Message)" }
    $definitions = @($ast.FindAll({
        param($node)
        $node -is [System.Management.Automation.Language.FunctionDefinitionAst]
    }, $true))
    if ($definitions.Count -eq 0) { throw 'Materializer AST contains no functions.' }
    foreach ($definition in $definitions) {
        $body = [string]$definition.Body.Extent.Text
        $body = $body.Substring(1, $body.Length - 2)
        Set-Item -Path ("Function:\script:" + $definition.Name) -Value ([scriptblock]::Create($body))
    }
}

function New-ContractZip {
    param(
        [Parameter(Mandatory)][string]$Path,
        [Parameter(Mandatory)][string]$ModulePath,
        [Parameter(Mandatory)][string]$Version,
        [Parameter(Mandatory)][System.Collections.IDictionary]$Files,
        [string[]]$AdditionalEntries = @()
    )
    [System.IO.Directory]::CreateDirectory((Split-Path -Parent $Path)) | Out-Null
    if (Test-Path -LiteralPath $Path) { [System.IO.File]::Delete($Path) }
    $archive = [System.IO.Compression.ZipFile]::Open($Path, [System.IO.Compression.ZipArchiveMode]::Create)
    try {
        foreach ($relative in @($Files.Keys | Sort-Object -CaseSensitive)) {
            $entry = $archive.CreateEntry("$ModulePath@$Version/$relative")
            $stream = $entry.Open()
            try {
                $bytes = $Utf8NoBom.GetBytes([string]$Files[$relative])
                $stream.Write($bytes, 0, $bytes.Length)
            }
            finally { $stream.Dispose() }
        }
        foreach ($relative in $AdditionalEntries) {
            $entry = $archive.CreateEntry("$ModulePath@$Version/$relative")
            $stream = $entry.Open()
            try {
                $bytes = $Utf8NoBom.GetBytes("case-collision`n")
                $stream.Write($bytes, 0, $bytes.Length)
            }
            finally { $stream.Dispose() }
        }
    }
    finally { $archive.Dispose() }
}

function New-SyntheticClosureFixture {
    param([Parameter(Mandatory)][string]$Root, [hashtable]$Options = @{})
    $modulePath = 'github.com/tree-sitter/go-tree-sitter'
    $version = 'v0.25.0'
    $moduleRoot = Join-Path $Root 'module'
    $vendorRoot = Join-Path $Root 'vendor'
    $zipPath = Join-Path $Root 'module.zip'
    [System.IO.Directory]::CreateDirectory($moduleRoot) | Out-Null
    [System.IO.Directory]::CreateDirectory($vendorRoot) | Out-Null

    $cflags = '#cgo CFLAGS: -I../../src'
    if ($Options.ContainsKey('wasmDefine')) { $cflags += ' -DTREE_SITTER_FEATURE_WASM' }
    $includeLines = @('#include <runtime.h>', '#include <tree_sitter/parser.h>')
    if ($Options.ContainsKey('dynamic')) { $includeLines += '#include HEADER_MACRO' }
    if ($Options.ContainsKey('escape')) { $includeLines += '#include "../../../../outside.h"' }
    $binding = (@('package fixture', '', '/*', $cflags, '#cgo CPPFLAGS: -I../../generated/src') + $includeLines + @('*/', 'import "C"', '')) -join "`n"
    $files = [ordered]@{
        'bindings/go/binding.go' = $binding
        'src/runtime.h' = "#define CONTRACT_RUNTIME 1`n"
        'src/complete-root.txt' = "complete src root`n"
        'src/wasm_store.c' = "#ifdef TREE_SITTER_FEATURE_WASM`n#include `"./stdlib-symbols.txt`"`n#endif`n"
        'generated/src/tree_sitter/parser.h' = "#define CONTRACT_PARSER 1`n"
        'generated/src/scanner.c' = "#include `"../common/scanner.h`"`n"
        'generated/src/node-types.json' = "{}`n"
        'generated/common/scanner.h' = "#include <tree_sitter/parser.h>`n"
        'generated/common/NOTICE' = "complete common root`n"
    }
    if ($Options.ContainsKey('unsupportedConditional')) {
        $files['generated/src/guarded.c'] = "#if FEATURE_LEVEL > 1`n#include `"missing.h`"`n#endif`n"
    }
    if ($Options.ContainsKey('ambiguous')) {
        $files['src/tree_sitter/parser.h'] = "#define AMBIGUOUS_PARSER 1`n"
    }
    foreach ($relative in $files.Keys) {
        Write-ContractText (Join-Path $moduleRoot ([string]$relative).Replace('/', '\')) ([string]$files[$relative])
    }
    $additional = @()
    if ($Options.ContainsKey('caseCollision')) { $additional = @('SRC/runtime.h') }
    New-ContractZip $zipPath $modulePath $version $files $additional
    if ($Options.ContainsKey('sourceTamper')) {
        Write-ContractText (Join-Path $moduleRoot 'generated\src\tree_sitter\parser.h') "tampered after zip`n"
    }
    $vendorBinding = Join-Path $vendorRoot ($modulePath.Replace('/', '\') + '\bindings\go\binding.go')
    Write-ContractText $vendorBinding $binding
    $junction = ''
    if ($Options.ContainsKey('reparse')) {
        $target = Join-Path $Root 'junction-target'
        [System.IO.Directory]::CreateDirectory($target) | Out-Null
        Write-ContractText (Join-Path $target 'target.txt') "target`n"
        $junction = Join-Path $moduleRoot 'src\reparse'
        New-Item -ItemType Junction -Path $junction -Target $target | Out-Null
        $TrackedJunctions.Add($junction)
    }
    return [pscustomobject]@{
        vendorRoot = $vendorRoot
        module = [pscustomobject]@{ path = $modulePath; version = $version; dir = $moduleRoot; zipPath = $zipPath }
        junction = $junction
    }
}

function Invoke-SyntheticClosureFailure {
    param([Parameter(Mandatory)][string]$Name, [Parameter(Mandatory)][hashtable]$Options, [Parameter(Mandatory)][string]$Expected)
    $root = Join-Path $TestRoot ("closure-" + $Name)
    $fixture = New-SyntheticClosureFixture $root $Options
    $unexpected = "Synthetic closure unexpectedly accepted case: $Name"
    try {
        try {
            Invoke-SupplementalClosure $fixture.vendorRoot @($fixture.module) | Out-Null
            throw $unexpected
        }
        catch {
            $actual = [string]$_.Exception.Message
            if ($actual -ceq $unexpected) { throw $actual }
            if (-not $actual.Contains($Expected)) {
                throw "Synthetic closure case '$Name' expected '$Expected'; received '$actual'."
            }
            Add-ContractResult 'materializer-closure' $Name $Expected
        }
    }
    finally {
        if (-not [string]::IsNullOrWhiteSpace([string]$fixture.junction) -and (Test-Path -LiteralPath $fixture.junction)) {
            [System.IO.Directory]::Delete([string]$fixture.junction)
            [void]$TrackedJunctions.Remove([string]$fixture.junction)
        }
    }
}

function New-SyntheticPatchFixture {
    param([Parameter(Mandatory)][string]$Root, [hashtable]$Options = @{})
    $fixtureRepo = Join-Path $Root 'repo'
    $vendorRoot = Join-Path $Root 'vendor'
    $moduleRoot = Join-Path $Root 'module'
    $modulePath = 'github.com/tree-sitter/tree-sitter-go'
    $version = 'v0.25.0'
    $patchRelative = 'third_party\go-vendor\patches\tree-sitter-go-v0.25.0-remove-absent-scanner.patch'
    Copy-ContractTree (Join-Path $AuthorityRepoRoot $patchRelative) (Join-Path $fixtureRepo $patchRelative)
    if ($Options.ContainsKey('patchTamper')) {
        $path = Join-Path $fixtureRepo $patchRelative
        [System.IO.File]::WriteAllBytes($path, (Add-ContractSuffix ([System.IO.File]::ReadAllBytes($path)) "`n# tamper"))
    }
    if ($Options.ContainsKey('extraPatch')) {
        Write-ContractText (Join-Path $fixtureRepo 'third_party\go-vendor\patches\second.patch') "second patch`n"
    }
    [System.IO.Directory]::CreateDirectory((Join-Path $moduleRoot 'src')) | Out-Null
    Copy-ContractTree (Join-Path $AuthorityRepoRoot 'vendor\github.com\tree-sitter\tree-sitter-go\src\grammar.json') (Join-Path $moduleRoot 'src\grammar.json')
    Copy-ContractTree (Join-Path $AuthorityRepoRoot 'vendor\github.com\tree-sitter\tree-sitter-go\src\parser.c') (Join-Path $moduleRoot 'src\parser.c')
    if ($Options.ContainsKey('grammarDrift')) { Write-ContractText (Join-Path $moduleRoot 'src\grammar.json') "{}`n" }
    if ($Options.ContainsKey('parserDrift')) { Write-ContractText (Join-Path $moduleRoot 'src\parser.c') "#define EXTERNAL_TOKEN_COUNT 0`n" }
    if ($Options.ContainsKey('sourceScanner')) { Write-ContractText (Join-Path $moduleRoot 'src\scanner.c') "scanner`n" }

    $postimagePath = Join-Path $AuthorityRepoRoot 'vendor\github.com\tree-sitter\tree-sitter-go\bindings\go\binding.go'
    $postimageText = $Utf8NoBomStrict.GetString([System.IO.File]::ReadAllBytes($postimagePath))
    $needle = "// #include `"../../src/parser.c`"`n"
    $removedBlock = "// #if __has_include(`"../../src/scanner.c`")`n// #include `"../../src/scanner.c`"`n// #endif`n"
    if (-not $postimageText.Contains($needle)) { throw 'Unable to reconstruct reviewed patch preimage.' }
    $preimageText = $postimageText.Replace($needle, $needle + $removedBlock)
    $bindingPath = Join-Path $vendorRoot ($modulePath.Replace('/', '\') + '\bindings\go\binding.go')
    Write-ContractText $bindingPath $preimageText
    if ($Options.ContainsKey('preimageDrift')) {
        [System.IO.File]::WriteAllBytes($bindingPath, (Add-ContractSuffix ([System.IO.File]::ReadAllBytes($bindingPath)) "x"))
    }
    $zipFiles = [ordered]@{
        'bindings/go/binding.go' = $preimageText
        'src/grammar.json' = [System.IO.File]::ReadAllText((Join-Path $moduleRoot 'src\grammar.json'))
        'src/parser.c' = [System.IO.File]::ReadAllText((Join-Path $moduleRoot 'src\parser.c'))
    }
    if ($Options.ContainsKey('zipScanner')) { $zipFiles['src/scanner.c'] = "scanner`n" }
    $zipPath = Join-Path $Root 'module.zip'
    New-ContractZip $zipPath $modulePath $version $zipFiles
    return [pscustomobject]@{
        repoRoot = $fixtureRepo
        vendorRoot = $vendorRoot
        bindingPath = $bindingPath
        module = [pscustomobject]@{ path = $modulePath; version = $version; dir = $moduleRoot; zipPath = $zipPath }
    }
}

function Invoke-SyntheticPatchFailure {
    param([Parameter(Mandatory)][string]$Name, [Parameter(Mandatory)][hashtable]$Options, [Parameter(Mandatory)][string]$Expected)
    $fixture = New-SyntheticPatchFixture (Join-Path $TestRoot ("patch-" + $Name)) $Options
    $script:RepoRoot = $fixture.repoRoot
    $unexpected = "Synthetic patch unexpectedly accepted case: $Name"
    try {
        Assert-AndApplyReviewedPatch $fixture.vendorRoot @($fixture.module) | Out-Null
        throw $unexpected
    }
    catch {
        $actual = [string]$_.Exception.Message
        if ($actual -ceq $unexpected) { throw $actual }
        if (-not $actual.Contains($Expected)) {
            throw "Synthetic patch case '$Name' expected '$Expected'; received '$actual'."
        }
        Add-ContractResult 'materializer-patch' $Name $Expected
    }
}

[System.Reflection.Assembly]::LoadWithPartialName('System.IO.Compression') | Out-Null
[System.Reflection.Assembly]::LoadWithPartialName('System.IO.Compression.FileSystem') | Out-Null

$LaneRootFull = Assert-ContractChildPath $RepoTempRoot $LaneRoot 'Go-vendor contract lane root'
$TestRoot = Join-Path $LaneRootFull ("go-vendor-contract-" + [Guid]::NewGuid().ToString('N'))
Assert-ContractChildPath $LaneRootFull $TestRoot 'Go-vendor contract test root' | Out-Null
[System.IO.Directory]::CreateDirectory($TestRoot) | Out-Null

try {
    $rootResult = Invoke-ContractVerifier $RepoRoot
    if ([string]$rootResult.status -cne 'PASS' -or [int]$rootResult.moduleCount -ne 45 -or [int]$rootResult.fileCount -ne 1798 -or
        [int]$rootResult.supplementalRoots -ne 20 -or [int]$rootResult.reviewedPatches -ne 1) {
        throw "Durable root verifier denominator mismatch: $($rootResult | ConvertTo-Json -Compress)"
    }
    Add-ContractResult 'verifier' 'durable_root_positive' '45 modules / 1798 files / 20 roots / 1 patch'

    $fixtureRoot = Join-Path $TestRoot 'packaged-go-src'
    New-PackagedAuthorityFixture $fixtureRoot
    $fixtureResult = Invoke-ContractVerifier $fixtureRoot
    if ([string]$fixtureResult.status -cne 'PASS' -or [int]$fixtureResult.moduleCount -ne 45 -or [int]$fixtureResult.fileCount -ne 1798) {
        throw "Packaged fixture verifier denominator mismatch: $($fixtureResult | ConvertTo-Json -Compress)"
    }
    Add-ContractResult 'verifier' 'packaged_root_positive' '45 modules / 1798 files'

    $manifestSnapshot = Get-Content -LiteralPath (Join-Path $fixtureRoot 'third_party\go-vendor\manifest.v1.json') -Raw | ConvertFrom-Json
    $baselineSample = @($manifestSnapshot.standardBaseline.files | Sort-Object -Property bytes, path | Select-Object -First 1)[0]
    $supplementalSample = @($manifestSnapshot.supplementalClosure.files | Sort-Object -Property bytes, path | Select-Object -First 1)[0]
    $licenseModule = @($manifestSnapshot.modules | Where-Object { @($_.licensePaths).Count -gt 1 } | Select-Object -First 1)[0]
    $licenseSample = [string]$licenseModule.licensePaths[0]

    Invoke-FileMutationCase 'missing_standard_file' $fixtureRoot ('vendor/' + [string]$baselineSample.path) 'missing' 'Vendor file inventory mismatch: missing=1'
    Invoke-FileMutationCase 'tampered_standard_file' $fixtureRoot ('vendor/' + [string]$baselineSample.path) 'tamper' 'hashOrBytes=1'
    Invoke-ExtraFileCase 'extra_vendor_file' $fixtureRoot 'vendor/zz-contract-extra.txt' 'extra=1'
    Invoke-FileMutationCase 'missing_supplemental_file' $fixtureRoot ('vendor/' + [string]$supplementalSample.path) 'missing' 'Vendor file inventory mismatch: missing=1'
    Invoke-FileMutationCase 'tampered_patched_binding' $fixtureRoot 'vendor/github.com/tree-sitter/tree-sitter-go/bindings/go/binding.go' 'tamper' 'hashOrBytes=1'
    Invoke-FileMutationCase 'missing_recursive_php_common' $fixtureRoot 'vendor/github.com/tree-sitter/tree-sitter-php/common/scanner.h' 'missing' 'Vendor file inventory mismatch: missing=1'
    Invoke-FileMutationCase 'missing_cppflags_typescript_parser' $fixtureRoot 'vendor/github.com/tree-sitter/tree-sitter-typescript/tsx/src/tree_sitter/parser.h' 'missing' 'Vendor file inventory mismatch: missing=1'
    Invoke-FileMutationCase 'missing_reviewed_patch' $fixtureRoot 'third_party/go-vendor/patches/tree-sitter-go-v0.25.0-remove-absent-scanner.patch' 'missing' 'Vendor file inventory mismatch: missing=1'
    Invoke-FileMutationCase 'tampered_reviewed_patch' $fixtureRoot 'third_party/go-vendor/patches/tree-sitter-go-v0.25.0-remove-absent-scanner.patch' 'tamper' 'hashOrBytes=1'
    Invoke-ExtraFileCase 'unreferenced_patch' $fixtureRoot 'third_party/go-vendor/patches/second.patch' 'extra=1'
    Invoke-FileMutationCase 'missing_license' $fixtureRoot ('third_party/go-vendor/' + $licenseSample) 'missing' 'Vendor file inventory mismatch: missing=1'
    Invoke-ExtraFileCase 'unreferenced_license' $fixtureRoot 'third_party/go-vendor/licenses/contract-extra/LICENSE' 'extra=1'
    Invoke-FileMutationCase 'go_mod_drift' $fixtureRoot 'go.mod' 'tamper' 'go.mod identity is outside the fixed R42 authority'
    Invoke-FileMutationCase 'go_sum_drift' $fixtureRoot 'go.sum' 'tamper' 'go.sum identity is outside the fixed R42 authority'
    Invoke-FileMutationCase 'missing_manifest' $fixtureRoot 'third_party/go-vendor/manifest.v1.json' 'missing' 'Required vendor authority file is absent'
    Invoke-FileMutationCase 'missing_packaged_verifier' $fixtureRoot 'scripts/verify-go-vendor.ps1' 'missing' 'Packaged Go source verifier is absent'
    Invoke-FileMutationCase 'tampered_packaged_verifier' $fixtureRoot 'scripts/verify-go-vendor.ps1' 'tamper' 'Verifier identity differs from manifest'
    Invoke-ReparseCase $fixtureRoot

    Invoke-ManifestMutationCase 'legacy_manifest' $fixtureRoot { param($m) $m.closureContractVersion = 1 } 'Unsupported closure contract: 1'
    Invoke-ManifestMutationCase 'wrong_go_version' $fixtureRoot { param($m) $m.goVersion = 'go1.26.2' } 'Wrong Go version'
    Invoke-ManifestMutationCase 'source_content_sum_mismatch' $fixtureRoot { param($m) $m.sourceInput.modules[0].contentSum = 'h1:contract-mismatch=' } 'Source module checksum provenance differs from unchanged go.sum'
    Invoke-ManifestMutationCase 'graph_go_mod_sum_mismatch' $fixtureRoot { param($m) $m.sourceInput.graphModules[0].goModSum = 'h1:contract-mismatch=' } 'Graph module row does not match unchanged go.sum'
    Invoke-ManifestMutationCase 'standard_run_mismatch' $fixtureRoot { param($m) $m.standardBaseline.equal = $false } 'Standard baseline equality contract is invalid'
    Invoke-ManifestMutationCase 'standard_digest_mismatch' $fixtureRoot { param($m) $m.standardBaseline.firstRun.treeSha256 = ('0' * 64) } 'Standard baseline inventory/equality record does not reconstruct'
    Invoke-ManifestMutationCase 'supplemental_digest_mismatch' $fixtureRoot { param($m) $m.supplementalClosure.treeSha256 = ('0' * 64) } 'Supplemental file inventory/equality record does not reconstruct'
    Invoke-ManifestMutationCase 'unresolved_closure_edge' $fixtureRoot { param($m) $m.supplementalClosure.recursiveEdges[0].disposition = 'missing' } 'Unresolved/unsupported active closure edge: missing'
    Invoke-ManifestMutationCase 'active_wasm_define' $fixtureRoot {
        param($m)
        $control = @($m.supplementalClosure.effectiveDefines | Where-Object module -CEQ 'github.com/tree-sitter/go-tree-sitter@v0.25.0')[0]
        $control.defines = @($control.defines) + @('TREE_SITTER_FEATURE_WASM')
    } 'TREE_SITTER_FEATURE_WASM is active in sealed build controls'
    Invoke-ManifestMutationCase 'patch_record_absent' $fixtureRoot { param($m) $m.reviewedPatches = @() } 'Reviewed patch cardinality must be exactly one; actual=0'
    Invoke-ManifestMutationCase 'second_patch_record' $fixtureRoot {
        param($m)
        $clone = (($m.reviewedPatches[0] | ConvertTo-Json -Depth 20) | ConvertFrom-Json)
        $m.reviewedPatches = @($m.reviewedPatches) + @($clone)
    } 'Reviewed patch cardinality must be exactly one; actual=2'
    foreach ($guard in @('preimageSha256', 'postimageSha256', 'occurrenceCount', 'grammarExternals', 'externalTokenCount', 'sourceScannerEntries', 'zipScannerEntries', 'fuzz', 'offset')) {
        Invoke-ManifestMutationCase ("patch_guard_" + $guard) $fixtureRoot {
            param($m)
            if ($guard -ceq 'preimageSha256' -or $guard -ceq 'postimageSha256') { $m.reviewedPatches[0].$guard = ('0' * 64) }
            else { $m.reviewedPatches[0].$guard = 1 }
            if ($guard -ceq 'occurrenceCount') { $m.reviewedPatches[0].$guard = 2 }
        }.GetNewClosure() 'Reviewed patch guard record is invalid'
    }
    Invoke-ManifestMutationCase 'supplemental_overlaps_baseline' $fixtureRoot {
        param($m)
        $baseline = @($m.standardBaseline.files | Where-Object path -CEQ 'modules.txt')[0]
        $overlap = [pscustomobject][ordered]@{
            path = 'modules.txt'; modulePath = 'github.com/daulet/tokenizers'; moduleVersion = 'v1.27.0'
            moduleRelativePath = 'modules.txt'; root = '.'; bytes = [long]$baseline.bytes; sha256 = [string]$baseline.sha256
        }
        $m.supplementalClosure.files = @($m.supplementalClosure.files) + @($overlap)
    } 'Supplemental file overlaps standard baseline: modules.txt'
    Invoke-ManifestMutationCase 'unclassified_origin' $fixtureRoot { param($m) $m.files[0].origin = 'unclassified' } 'Manifest contains a malformed file inventory row'
    Invoke-ManifestMutationCase 'wrong_standard_origin' $fixtureRoot {
        param($m)
        $row = @($m.files | Where-Object origin -CEQ 'standard' | Select-Object -First 1)[0]
        $row.origin = 'supplemental'
    } 'Final standard partition drift'
    Invoke-ManifestMutationCase 'manifest_self_inclusion' $fixtureRoot {
        param($m)
        $row = [pscustomobject][ordered]@{ path = 'third_party/go-vendor/manifest.v1.json'; bytes = 1; sha256 = ('0' * 64); origin = 'unclassified' }
        $orderedRows = [System.Collections.Generic.SortedDictionary[string, object]]::new([System.StringComparer]::Ordinal)
        foreach ($candidate in @($m.files) + @($row)) { $orderedRows.Add([string]$candidate.path, $candidate) }
        $m.files = @($orderedRows.Values)
    } 'Manifest file row is outside the allowed authority trees'
    Invoke-ManifestMutationCase 'case_colliding_manifest_path' $fixtureRoot {
        param($m)
        $clone = (($m.files[0] | ConvertTo-Json -Depth 20) | ConvertFrom-Json)
        $clone.path = ([string]$clone.path).ToUpperInvariant()
        $m.files = @($m.files) + @($clone)
    } 'Duplicate or case-colliding inventory path'
    Invoke-ManifestMutationCase 'path_escape' $fixtureRoot { param($m) $m.supplementalClosure.files[0].path = '../escape' } 'Non-canonical relative path'
    Invoke-ManifestMutationCase 'unknown_license_disposition' $fixtureRoot { param($m) $m.modules[0].licenseDisposition = 'unknown' } 'Incomplete module license disposition count'
    Invoke-ManifestMutationCase 'missing_license_disposition' $fixtureRoot { param($m) $m.modules[0].licensePaths = @() } 'Incomplete module license disposition count'
    Invoke-ManifestMutationCase 'unreferenced_license_manifest' $fixtureRoot {
        param($m)
        $module = @($m.modules | Where-Object { @($_.licensePaths).Count -gt 1 } | Select-Object -First 1)[0]
        $module.licensePaths = @($module.licensePaths[0])
    } 'License inventory coverage is incomplete'

    $fixtureFinal = Invoke-ContractVerifier $fixtureRoot
    if ([string]$fixtureFinal.status -cne 'PASS') { throw 'Packaged fixture was not restored after verifier matrix.' }
    Add-ContractResult 'verifier' 'packaged_root_restored' 'PASS after all mutations'

    Import-MaterializerFunctions
    $script:PatchAuthorityRelativePath = 'third_party/go-vendor/patches/tree-sitter-go-v0.25.0-remove-absent-scanner.patch'
    $script:ExpectedPatchSha256 = '43BB195FFF439C7DCD0D057094EC1169304553E357F6D0B2781DC5713516BEEA'
    $script:PatchedModulePath = 'github.com/tree-sitter/tree-sitter-go'
    $script:PatchedModuleVersion = 'v0.25.0'
    $script:PatchedBindingRelativePath = 'github.com/tree-sitter/tree-sitter-go/bindings/go/binding.go'
    $script:ExpectedPatchPreimageBytes = 333
    $script:ExpectedPatchPreimageSha256 = '680E50820F2F7429FEAC88B69DDCEE67735E6E692DFD1F98147109D77D296096'
    $script:ExpectedPatchPostimageBytes = 245
    $script:ExpectedPatchPostimageSha256 = '5A456F352CFEF15430F3FCFB155A3E1BDE38F244E5F8537026553B1F21547713'
    $script:ExpectedGrammarBytes = 198042
    $script:ExpectedGrammarSha256 = '0DCC665AA521A1B73A200BE13B4E376780CCA037D764A165F12864BABFE28000'
    $script:ExpectedParserBytes = 1572685
    $script:ExpectedParserSha256 = '3DBF6ED1238B5DFCF2BE4D2F2D4CB27A14D34F34D7784ECCCCBFD532FD4A6D85'
    $script:Utf8NoBom = [System.Text.UTF8Encoding]::new($false)
    $script:Utf8NoBomStrict = [System.Text.UTF8Encoding]::new($false, $true)

    $closurePositive = New-SyntheticClosureFixture (Join-Path $TestRoot 'closure-positive')
    $closure = Invoke-SupplementalClosure $closurePositive.vendorRoot @($closurePositive.module)
    $rootNames = @($closure.roots | ForEach-Object { [string]$_.root } | Sort-Object -CaseSensitive)
    if (($rootNames -join ',') -cne 'generated/common,generated/src,src') { throw "Synthetic closure roots differ: $($rootNames -join ',')" }
    $controls = @($closure.effectiveDefines | Where-Object module -CEQ 'github.com/tree-sitter/go-tree-sitter@v0.25.0')[0]
    if (@($controls.includeRoots | Where-Object { $_ -ceq 'src' }).Count -ne 1 -or @($controls.includeRoots | Where-Object { $_ -ceq 'generated/src' }).Count -ne 1) {
        throw 'Synthetic closure did not preserve both CFLAGS and CPPFLAGS include roots.'
    }
    if (@($closure.recursiveEdges | Where-Object { $_.resolvedPath -ceq 'generated/common/scanner.h' }).Count -lt 1 -or
        @($closure.recursiveEdges | Where-Object { $_.referencer -ceq 'generated/common/scanner.h' -and $_.resolvedPath -ceq 'generated/src/tree_sitter/parser.h' -and $_.resolution -ceq 'include_root' }).Count -ne 1) {
        throw 'Synthetic closure did not reach the required fixed point.'
    }
    foreach ($relative in @('src/complete-root.txt', 'generated/src/node-types.json', 'generated/common/NOTICE')) {
        $path = Join-Path $closurePositive.vendorRoot ('github.com\tree-sitter\go-tree-sitter\' + $relative.Replace('/', '\'))
        if (-not (Test-Path -LiteralPath $path -PathType Leaf)) { throw "Complete supplemental root omitted $relative" }
    }
    $inactiveWasm = @($closure.inactiveConditionalReferences | Where-Object { $_.resolvedPath -ceq 'src/stdlib-symbols.txt' -and $_.state -ceq 'inactive' })
    if ($inactiveWasm.Count -ne 1) { throw 'Synthetic closure inactive WASM denominator differs.' }
    Add-ContractResult 'materializer-closure' 'positive_cppflags_fixed_point_complete_roots' '3 roots / CFLAGS+CPPFLAGS / recursive common / inactive WASM'

    Invoke-SyntheticClosureFailure 'dynamic_include' @{ dynamic = $true } 'Dynamic include is forbidden:'
    Invoke-SyntheticClosureFailure 'path_escape' @{ escape = $true } 'Module-local path escapes its module:'
    Invoke-SyntheticClosureFailure 'unsupported_conditional' @{ unsupportedConditional = $true } 'Unsupported conditional guards a module-local include:'
    Invoke-SyntheticClosureFailure 'ambiguous_include' @{ ambiguous = $true } 'Ambiguous module-local include'
    Invoke-SyntheticClosureFailure 'zip_case_collision' @{ caseCollision = $true } 'Module zip has duplicate/case-colliding entry:'
    Invoke-SyntheticClosureFailure 'source_zip_mismatch' @{ sourceTamper = $true } 'Extracted source differs from authenticated zip entry:'
    Invoke-SyntheticClosureFailure 'source_reparse' @{ reparse = $true } 'contains 1 reparse point(s)'
    Invoke-SyntheticClosureFailure 'active_wasm' @{ wasmDefine = $true } 'TREE_SITTER_FEATURE_WASM is forbidden in effective cgo flags'

    $patchPositive = New-SyntheticPatchFixture (Join-Path $TestRoot 'patch-positive')
    $script:RepoRoot = $patchPositive.repoRoot
    $patchRecord = Assert-AndApplyReviewedPatch $patchPositive.vendorRoot @($patchPositive.module)
    if ([string]$patchRecord.postimageSha256 -cne $script:ExpectedPatchPostimageSha256 -or
        (Get-Sha256 $patchPositive.bindingPath) -cne $script:ExpectedPatchPostimageSha256) {
        throw 'Synthetic reviewed patch positive transform differs.'
    }
    Add-ContractResult 'materializer-patch' 'positive_exact_transform' '333-byte preimage -> 245-byte postimage'
    $unexpectedSecondPatch = 'Second patch application unexpectedly succeeded.'
    try {
        Assert-AndApplyReviewedPatch $patchPositive.vendorRoot @($patchPositive.module) | Out-Null
        throw $unexpectedSecondPatch
    }
    catch {
        if ([string]$_.Exception.Message -ceq $unexpectedSecondPatch) { throw $unexpectedSecondPatch }
        if (-not ([string]$_.Exception.Message).Contains('Tree-sitter Go binding preimage drifted')) { throw }
        Add-ContractResult 'materializer-patch' 'second_application' 'Tree-sitter Go binding preimage drifted'
    }
    Invoke-SyntheticPatchFailure 'patch_tamper' @{ patchTamper = $true } 'Reviewed patch cardinality or identity is invalid'
    Invoke-SyntheticPatchFailure 'extra_patch' @{ extraPatch = $true } 'Reviewed patch cardinality or identity is invalid'
    Invoke-SyntheticPatchFailure 'preimage_drift' @{ preimageDrift = $true } 'Tree-sitter Go binding preimage drifted'
    Invoke-SyntheticPatchFailure 'grammar_drift' @{ grammarDrift = $true } 'tree-sitter-go grammar identity drifted'
    Invoke-SyntheticPatchFailure 'parser_drift' @{ parserDrift = $true } 'tree-sitter-go parser identity drifted'
    Invoke-SyntheticPatchFailure 'source_scanner' @{ sourceScanner = $true } 'Unexpected tree-sitter-go scanner source'
    Invoke-SyntheticPatchFailure 'zip_scanner' @{ zipScanner = $true } 'Unexpected tree-sitter-go scanner zip entry'

    $suiteCounts = @($Results | Group-Object suite | Sort-Object Name | ForEach-Object {
        [pscustomobject][ordered]@{ suite = $_.Name; passed = $_.Count }
    })
    [pscustomobject][ordered]@{
        status = 'PASS'
        shell = $PSVersionTable.PSVersion.ToString()
        cases = $Results.Count
        suites = $suiteCounts
        results = @($Results)
    } | ConvertTo-Json -Depth 8
}
finally {
    foreach ($junction in @($TrackedJunctions)) {
        if (Test-Path -LiteralPath $junction) { [System.IO.Directory]::Delete($junction) }
    }
    $verifiedTestRoot = Assert-ContractChildPath $LaneRootFull $TestRoot 'Go-vendor contract cleanup root'
    if (Test-Path -LiteralPath $verifiedTestRoot) {
        Remove-Item -LiteralPath $verifiedTestRoot -Recurse -Force
    }
}
