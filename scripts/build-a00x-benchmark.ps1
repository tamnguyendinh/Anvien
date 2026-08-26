[CmdletBinding()]
param(
    [Parameter(Mandatory = $true)]
    [string]$AttemptId,

    [Parameter(Mandatory = $true)]
    [string]$OverlayManifestPath,

    [Parameter(Mandatory = $true)]
    [string]$OutputExecutablePath,

    [Parameter(Mandatory = $true)]
    [string]$ExpectedOverlaySha256,

    [Parameter(Mandatory = $true)]
    [string[]]$ExpectedMappedSourceHash,

    [Parameter(Mandatory = $true)]
    [string[]]$ExpectedCandidateSourceHash,

    [Parameter(Mandatory = $true)]
    [string[]]$ExpectedNativeHash,

    [string]$GoExecutable
)

Set-StrictMode -Version Latest
$ErrorActionPreference = 'Stop'

function Throw-ContractError {
    param([Parameter(Mandatory = $true)][string]$Message)

    throw [System.InvalidOperationException]::new(
        "A00X_BENCHMARK_BUILD_CONTRACT: $Message"
    )
}

function Get-NormalizedFullPath {
    param(
        [Parameter(Mandatory = $true)][string]$PathValue,
        [Parameter(Mandatory = $true)][string]$BasePath
    )

    if ([string]::IsNullOrWhiteSpace($PathValue)) {
        Throw-ContractError 'path value must not be blank'
    }

    return [System.IO.Path]::GetFullPath($PathValue, $BasePath)
}

function Test-PathAtOrBelow {
    param(
        [Parameter(Mandatory = $true)][string]$CandidatePath,
        [Parameter(Mandatory = $true)][string]$RootPath
    )

    $candidate = [System.IO.Path]::GetFullPath($CandidatePath)
    $root = [System.IO.Path]::GetFullPath($RootPath)
    if ($candidate.Equals($root, [System.StringComparison]::OrdinalIgnoreCase)) {
        return $true
    }

    $rootWithSeparator = $root.TrimEnd(
        [System.IO.Path]::DirectorySeparatorChar,
        [System.IO.Path]::AltDirectorySeparatorChar
    ) + [System.IO.Path]::DirectorySeparatorChar

    return $candidate.StartsWith(
        $rootWithSeparator,
        [System.StringComparison]::OrdinalIgnoreCase
    )
}

function Assert-NoReparsePointInExistingPath {
    param(
        [Parameter(Mandatory = $true)][string]$PathValue,
        [Parameter(Mandatory = $true)][string]$Label
    )

    $fullPath = [System.IO.Path]::GetFullPath($PathValue)
    $volumeRoot = [System.IO.Path]::GetPathRoot($fullPath)
    $current = $volumeRoot
    $relative = $fullPath.Substring($volumeRoot.Length)
    $segments = $relative.Split(
        [char[]]@('\', '/'),
        [System.StringSplitOptions]::RemoveEmptyEntries
    )

    foreach ($segment in $segments) {
        $current = Join-Path -Path $current -ChildPath $segment
        if (-not (Test-Path -LiteralPath $current)) {
            break
        }

        $item = Get-Item -LiteralPath $current -Force
        if (($item.Attributes -band [System.IO.FileAttributes]::ReparsePoint) -ne 0) {
            Throw-ContractError "$Label traverses reparse point: $current"
        }
    }
}

function Assert-RegularFile {
    param(
        [Parameter(Mandatory = $true)][string]$PathValue,
        [Parameter(Mandatory = $true)][string]$Label
    )

    if (-not (Test-Path -LiteralPath $PathValue -PathType Leaf)) {
        Throw-ContractError "$Label is not a regular file: $PathValue"
    }

    $item = Get-Item -LiteralPath $PathValue -Force
    if (($item.Attributes -band [System.IO.FileAttributes]::ReparsePoint) -ne 0) {
        Throw-ContractError "$Label must not be a reparse point: $PathValue"
    }
}

function Assert-RegularDirectory {
    param(
        [Parameter(Mandatory = $true)][string]$PathValue,
        [Parameter(Mandatory = $true)][string]$Label
    )

    if (-not (Test-Path -LiteralPath $PathValue -PathType Container)) {
        Throw-ContractError "$Label is not a regular directory: $PathValue"
    }

    $item = Get-Item -LiteralPath $PathValue -Force
    if (($item.Attributes -band [System.IO.FileAttributes]::ReparsePoint) -ne 0) {
        Throw-ContractError "$Label must not be a reparse point: $PathValue"
    }
}

function Normalize-Sha256 {
    param(
        [Parameter(Mandatory = $true)][string]$HashValue,
        [Parameter(Mandatory = $true)][string]$Label
    )

    if ($HashValue -notmatch '^[A-Fa-f0-9]{64}$') {
        Throw-ContractError "$Label must be exactly 64 hexadecimal characters"
    }

    return $HashValue.ToUpperInvariant()
}

function Get-Sha256 {
    param([Parameter(Mandatory = $true)][string]$PathValue)

    return (Get-FileHash -LiteralPath $PathValue -Algorithm SHA256).Hash.ToUpperInvariant()
}

function Assert-ExpectedSha256 {
    param(
        [Parameter(Mandatory = $true)][string]$PathValue,
        [Parameter(Mandatory = $true)][string]$ExpectedHash,
        [Parameter(Mandatory = $true)][string]$Label
    )

    $actualHash = Get-Sha256 -PathValue $PathValue
    if (-not $actualHash.Equals($ExpectedHash, [System.StringComparison]::Ordinal)) {
        Throw-ContractError (
            "$Label SHA-256 mismatch: expected $ExpectedHash, actual $actualHash, path $PathValue"
        )
    }

    return $actualHash
}

function Split-HashContractEntry {
    param(
        [Parameter(Mandatory = $true)][string]$Entry,
        [Parameter(Mandatory = $true)][string]$Label
    )

    if ([string]::IsNullOrWhiteSpace($Entry)) {
        Throw-ContractError "$Label contains a blank entry"
    }

    $separatorIndex = $Entry.LastIndexOf('=')
    if ($separatorIndex -le 0 -or $separatorIndex -ge ($Entry.Length - 1)) {
        Throw-ContractError "$Label entry must use key=sha256 syntax: $Entry"
    }

    $key = $Entry.Substring(0, $separatorIndex).Trim()
    $hash = Normalize-Sha256 -HashValue $Entry.Substring($separatorIndex + 1).Trim() -Label $Label
    if ([string]::IsNullOrWhiteSpace($key)) {
        Throw-ContractError "$Label contains a blank key"
    }

    return [pscustomobject]@{
        Key = $key
        Hash = $hash
    }
}

function Invoke-CapturedProcess {
    param(
        [Parameter(Mandatory = $true)][string]$FilePath,
        [Parameter(Mandatory = $true)][string[]]$Arguments,
        [Parameter(Mandatory = $true)][string]$WorkingDirectory,
        [Parameter(Mandatory = $true)]
        [System.Collections.IDictionary]$EnvironmentOverrides
    )

    $startInfo = [System.Diagnostics.ProcessStartInfo]::new()
    $startInfo.FileName = $FilePath
    $startInfo.WorkingDirectory = $WorkingDirectory
    $startInfo.UseShellExecute = $false
    $startInfo.CreateNoWindow = $true
    $startInfo.RedirectStandardOutput = $true
    $startInfo.RedirectStandardError = $true

    foreach ($argument in $Arguments) {
        [void]$startInfo.ArgumentList.Add($argument)
    }

    foreach ($key in $EnvironmentOverrides.Keys) {
        $startInfo.Environment[[string]$key] = [string]$EnvironmentOverrides[$key]
    }

    $process = [System.Diagnostics.Process]::new()
    $process.StartInfo = $startInfo
    $startedAtUtc = [System.DateTime]::UtcNow
    if (-not $process.Start()) {
        Throw-ContractError "failed to start child process: ${FilePath}"
    }

    $stdoutTask = $process.StandardOutput.ReadToEndAsync()
    $stderrTask = $process.StandardError.ReadToEndAsync()
    $process.WaitForExit()
    $stdout = $stdoutTask.GetAwaiter().GetResult()
    $stderr = $stderrTask.GetAwaiter().GetResult()
    $endedAtUtc = [System.DateTime]::UtcNow

    return [pscustomobject]@{
        ExitCode = $process.ExitCode
        StandardOutput = $stdout
        StandardError = $stderr
        StartedAtUtc = $startedAtUtc
        EndedAtUtc = $endedAtUtc
    }
}

$invocationStartedAtUtc = [System.DateTime]::UtcNow

if ($AttemptId -notmatch '^A[0-9]{3}$') {
    Throw-ContractError ('AttemptId must match ^A[0-9]{3}$: {0}' -f $AttemptId)
}

$scriptPath = [System.IO.Path]::GetFullPath($PSCommandPath)
$scriptsDirectory = [System.IO.Path]::GetFullPath($PSScriptRoot)
$repoRoot = [System.IO.Directory]::GetParent($scriptsDirectory).FullName
$expectedScriptPath = Join-Path -Path $scriptsDirectory -ChildPath 'build-a00x-benchmark.ps1'
if (-not $scriptPath.Equals(
    [System.IO.Path]::GetFullPath($expectedScriptPath),
    [System.StringComparison]::OrdinalIgnoreCase
)) {
    Throw-ContractError "script path is not the canonical repository scripts path: $scriptPath"
}

Assert-NoReparsePointInExistingPath -PathValue $scriptPath -Label 'script path'
Assert-RegularFile -PathValue $scriptPath -Label 'script path'

$gitRootOutput = @(& git -C $repoRoot rev-parse --show-toplevel 2>&1)
$gitRootExitCode = $LASTEXITCODE
if ($gitRootExitCode -ne 0) {
    Throw-ContractError (
        'failed to resolve repository root: ' +
        [string]::Join([System.Environment]::NewLine, @($gitRootOutput | ForEach-Object { $_.ToString() }))
    )
}

$gitRoot = Get-NormalizedFullPath -PathValue $gitRootOutput[0].ToString().Trim() -BasePath $repoRoot
if (-not $gitRoot.Equals($repoRoot, [System.StringComparison]::OrdinalIgnoreCase)) {
    Throw-ContractError "script-derived root $repoRoot does not match Git root $gitRoot"
}

Assert-NoReparsePointInExistingPath -PathValue $repoRoot -Label 'repository root'
Assert-RegularDirectory -PathValue $repoRoot -Label 'repository root'

$temporaryRoot = Join-Path -Path $repoRoot -ChildPath '.tmp'
$outputFullPath = Get-NormalizedFullPath -PathValue $OutputExecutablePath -BasePath $repoRoot
if (-not [System.IO.Path]::IsPathFullyQualified($OutputExecutablePath)) {
    Throw-ContractError "OutputExecutablePath must be absolute: $OutputExecutablePath"
}
if (-not (Test-PathAtOrBelow -CandidatePath $outputFullPath -RootPath $temporaryRoot)) {
    Throw-ContractError "output executable must be beneath repository .tmp: $outputFullPath"
}
if ($outputFullPath.Equals(
    [System.IO.Path]::GetFullPath($temporaryRoot),
    [System.StringComparison]::OrdinalIgnoreCase
)) {
    Throw-ContractError 'output executable cannot be the .tmp directory itself'
}
if (-not [System.IO.Path]::GetExtension($outputFullPath).Equals(
    '.exe',
    [System.StringComparison]::OrdinalIgnoreCase
)) {
    Throw-ContractError "output executable must have an .exe extension: $outputFullPath"
}

$outputDirectory = [System.IO.Path]::GetDirectoryName($outputFullPath)
if ([string]::IsNullOrWhiteSpace($outputDirectory)) {
    Throw-ContractError "output executable has no parent directory: $outputFullPath"
}

$outputBaseName = [System.IO.Path]::GetFileNameWithoutExtension($outputFullPath)
if ([string]::IsNullOrWhiteSpace($outputBaseName)) {
    Throw-ContractError "output executable has no base name: $outputFullPath"
}
if ($outputBaseName.Equals('lbug_shared', [System.StringComparison]::OrdinalIgnoreCase)) {
    Throw-ContractError 'output executable name would collide with the pinned native DLL'
}

$outputDllPath = Join-Path -Path $outputDirectory -ChildPath 'lbug_shared.dll'
$provenancePath = Join-Path -Path $outputDirectory -ChildPath "$outputBaseName.provenance.json"

$protectedOutputRoots = @(
    (Join-Path -Path $repoRoot -ChildPath 'anvien\bin'),
    (Join-Path -Path $repoRoot -ChildPath 'anvien-launcher'),
    (Join-Path -Path $repoRoot -ChildPath 'anvien-server'),
    (Join-Path -Path $repoRoot -ChildPath 'anvien-web'),
    (Join-Path -Path $repoRoot -ChildPath '.anvien')
)
foreach ($protectedRoot in $protectedOutputRoots) {
    if (Test-PathAtOrBelow -CandidatePath $outputFullPath -RootPath $protectedRoot) {
        Throw-ContractError "output executable targets a canonical output surface: $outputFullPath"
    }
}

Assert-NoReparsePointInExistingPath -PathValue $outputDirectory -Label 'output path'
if (Test-Path -LiteralPath $outputFullPath) {
    Throw-ContractError "refusing to overwrite existing output executable: $outputFullPath"
}
if (Test-Path -LiteralPath $outputDllPath) {
    Throw-ContractError "refusing to overwrite existing output DLL: $outputDllPath"
}
if (Test-Path -LiteralPath $provenancePath) {
    Throw-ContractError "refusing to overwrite existing provenance JSON: $provenancePath"
}

$goModPath = Join-Path -Path $repoRoot -ChildPath 'go.mod'
$goSumPath = Join-Path -Path $repoRoot -ChildPath 'go.sum'
$vendorPath = Join-Path -Path $repoRoot -ChildPath 'vendor'
foreach ($requiredFile in @(
    [pscustomobject]@{ Path = $goModPath; Label = 'go.mod' },
    [pscustomobject]@{ Path = $goSumPath; Label = 'go.sum' }
)) {
    Assert-NoReparsePointInExistingPath -PathValue $requiredFile.Path -Label $requiredFile.Label
    Assert-RegularFile -PathValue $requiredFile.Path -Label $requiredFile.Label
}
Assert-NoReparsePointInExistingPath -PathValue $vendorPath -Label 'vendor directory'
Assert-RegularDirectory -PathValue $vendorPath -Label 'vendor directory'

$manifestFullPath = Get-NormalizedFullPath -PathValue $OverlayManifestPath -BasePath $repoRoot
if (-not [System.IO.Path]::IsPathFullyQualified($OverlayManifestPath)) {
    Throw-ContractError "OverlayManifestPath must be absolute: $OverlayManifestPath"
}
if (-not (Test-PathAtOrBelow -CandidatePath $manifestFullPath -RootPath $temporaryRoot)) {
    Throw-ContractError "overlay manifest must be beneath repository .tmp: $manifestFullPath"
}
Assert-NoReparsePointInExistingPath -PathValue $manifestFullPath -Label 'overlay manifest'
Assert-RegularFile -PathValue $manifestFullPath -Label 'overlay manifest'

$normalizedExpectedOverlayHash = Normalize-Sha256 -HashValue $ExpectedOverlaySha256 -Label 'ExpectedOverlaySha256'
$actualOverlayHash = Assert-ExpectedSha256 -PathValue $manifestFullPath -ExpectedHash $normalizedExpectedOverlayHash -Label 'overlay manifest'

$expectedMappings = [System.Collections.Generic.Dictionary[string, string]]::new(
    [System.StringComparer]::OrdinalIgnoreCase
)
if (@($ExpectedMappedSourceHash).Count -eq 0) {
    Throw-ContractError 'ExpectedMappedSourceHash must contain at least one mapping'
}
foreach ($entry in $ExpectedMappedSourceHash) {
    $parsed = Split-HashContractEntry -Entry $entry -Label 'ExpectedMappedSourceHash'
    if (-not [System.IO.Path]::IsPathFullyQualified($parsed.Key)) {
        Throw-ContractError "mapped source contract key must be absolute: $($parsed.Key)"
    }

    $sourcePath = Get-NormalizedFullPath -PathValue $parsed.Key -BasePath $repoRoot
    if (-not (Test-PathAtOrBelow -CandidatePath $sourcePath -RootPath $repoRoot)) {
        Throw-ContractError "mapped source contract key escapes repository root: $sourcePath"
    }
    if (Test-PathAtOrBelow -CandidatePath $sourcePath -RootPath $temporaryRoot) {
        Throw-ContractError "mapped source contract key must refer to repository source, not .tmp: $sourcePath"
    }
    Assert-NoReparsePointInExistingPath -PathValue $sourcePath -Label 'mapped source contract key'
    Assert-RegularFile -PathValue $sourcePath -Label 'mapped source contract key'

    if (-not $expectedMappings.TryAdd($sourcePath, $parsed.Hash)) {
        Throw-ContractError "duplicate mapped source contract key: $sourcePath"
    }
}

$mappingRecords = [System.Collections.Generic.List[object]]::new()
$manifestMappings = [System.Collections.Generic.Dictionary[string, string]]::new(
    [System.StringComparer]::OrdinalIgnoreCase
)
$jsonOptions = [System.Text.Json.JsonDocumentOptions]::new()
$jsonOptions.AllowTrailingCommas = $false
$jsonOptions.CommentHandling = [System.Text.Json.JsonCommentHandling]::Disallow
try {
    $manifestDocument = [System.Text.Json.JsonDocument]::Parse(
        [System.IO.File]::ReadAllText($manifestFullPath),
        $jsonOptions
    )
}
catch {
    Throw-ContractError "overlay manifest is not valid JSON: $($_.Exception.Message)"
}

try {
    if ($manifestDocument.RootElement.ValueKind -ne [System.Text.Json.JsonValueKind]::Object) {
        Throw-ContractError 'overlay manifest root must be an object'
    }

    $replaceProperty = $null
    $rootPropertyCount = 0
    foreach ($property in $manifestDocument.RootElement.EnumerateObject()) {
        $rootPropertyCount++
        if ($property.Name -cne 'Replace') {
            Throw-ContractError "overlay manifest contains unsupported root property: $($property.Name)"
        }
        if ($null -ne $replaceProperty) {
            Throw-ContractError 'overlay manifest contains duplicate Replace properties'
        }
        $replaceProperty = $property
    }

    if ($rootPropertyCount -ne 1 -or $null -eq $replaceProperty) {
        Throw-ContractError 'overlay manifest must contain exactly one Replace object'
    }
    if ($replaceProperty.Value.ValueKind -ne [System.Text.Json.JsonValueKind]::Object) {
        Throw-ContractError 'overlay manifest Replace value must be an object'
    }

    foreach ($mappingProperty in $replaceProperty.Value.EnumerateObject()) {
        if ($mappingProperty.Value.ValueKind -ne [System.Text.Json.JsonValueKind]::String) {
            Throw-ContractError "overlay replacement must be a string: $($mappingProperty.Name)"
        }
        if (-not [System.IO.Path]::IsPathFullyQualified($mappingProperty.Name)) {
            Throw-ContractError "overlay source path must be absolute: $($mappingProperty.Name)"
        }

        $replacementValue = $mappingProperty.Value.GetString()
        if ([string]::IsNullOrWhiteSpace($replacementValue)) {
            Throw-ContractError "overlay replacement path is blank: $($mappingProperty.Name)"
        }
        if (-not [System.IO.Path]::IsPathFullyQualified($replacementValue)) {
            Throw-ContractError "overlay replacement path must be absolute: $replacementValue"
        }

        $sourcePath = Get-NormalizedFullPath -PathValue $mappingProperty.Name -BasePath $repoRoot
        $replacementPath = Get-NormalizedFullPath -PathValue $replacementValue -BasePath $repoRoot
        if (-not (Test-PathAtOrBelow -CandidatePath $sourcePath -RootPath $repoRoot)) {
            Throw-ContractError "overlay source path escapes repository root: $sourcePath"
        }
        if (Test-PathAtOrBelow -CandidatePath $sourcePath -RootPath $temporaryRoot) {
            Throw-ContractError "overlay source path must refer to repository source, not .tmp: $sourcePath"
        }
        if (-not (Test-PathAtOrBelow -CandidatePath $replacementPath -RootPath $temporaryRoot)) {
            Throw-ContractError "overlay replacement path must be beneath repository .tmp: $replacementPath"
        }

        Assert-NoReparsePointInExistingPath -PathValue $sourcePath -Label 'overlay source path'
        Assert-RegularFile -PathValue $sourcePath -Label 'overlay source path'
        Assert-NoReparsePointInExistingPath -PathValue $replacementPath -Label 'overlay replacement path'
        Assert-RegularFile -PathValue $replacementPath -Label 'overlay replacement path'

        if (-not $manifestMappings.TryAdd($sourcePath, $replacementPath)) {
            Throw-ContractError "overlay contains duplicate canonical source path: $sourcePath"
        }
    }
}
finally {
    $manifestDocument.Dispose()
}

if ($manifestMappings.Count -ne $expectedMappings.Count) {
    Throw-ContractError (
        "overlay mapping count mismatch: expected $($expectedMappings.Count), actual $($manifestMappings.Count)"
    )
}

foreach ($sourcePath in $manifestMappings.Keys) {
    if (-not $expectedMappings.ContainsKey($sourcePath)) {
        Throw-ContractError "overlay contains unexpected source mapping: $sourcePath"
    }

    $replacementPath = $manifestMappings[$sourcePath]
    $expectedReplacementHash = $expectedMappings[$sourcePath]
    $actualReplacementHash = Assert-ExpectedSha256 -PathValue $replacementPath -ExpectedHash $expectedReplacementHash -Label 'overlay replacement'
    $sourceHash = Get-Sha256 -PathValue $sourcePath
    $mappingRecords.Add([pscustomobject]@{
        sourcePath = $sourcePath
        sourceSha256 = $sourceHash
        replacementPath = $replacementPath
        expectedReplacementSha256 = $expectedReplacementHash
        actualReplacementSha256 = $actualReplacementHash
    })
}

$candidateRecords = [System.Collections.Generic.List[object]]::new()
$candidatePaths = [System.Collections.Generic.HashSet[string]]::new(
    [System.StringComparer]::OrdinalIgnoreCase
)
if (@($ExpectedCandidateSourceHash).Count -eq 0) {
    Throw-ContractError 'ExpectedCandidateSourceHash must contain at least one source'
}
foreach ($entry in $ExpectedCandidateSourceHash) {
    $parsed = Split-HashContractEntry -Entry $entry -Label 'ExpectedCandidateSourceHash'
    if ([System.IO.Path]::IsPathFullyQualified($parsed.Key) -or [System.IO.Path]::IsPathRooted($parsed.Key)) {
        Throw-ContractError "candidate source contract key must be repository-relative: $($parsed.Key)"
    }

    $candidateSegments = $parsed.Key.Replace('/', '\').Split(
        [char[]]@('\'),
        [System.StringSplitOptions]::RemoveEmptyEntries
    )
    if (@($candidateSegments | Where-Object { $_ -eq '.' -or $_ -eq '..' }).Count -ne 0) {
        Throw-ContractError "candidate source contract key must not contain dot segments: $($parsed.Key)"
    }

    $candidatePath = Get-NormalizedFullPath -PathValue $parsed.Key -BasePath $repoRoot
    if (-not (Test-PathAtOrBelow -CandidatePath $candidatePath -RootPath $repoRoot)) {
        Throw-ContractError "candidate source contract key escapes repository root: $candidatePath"
    }
    if (Test-PathAtOrBelow -CandidatePath $candidatePath -RootPath $temporaryRoot) {
        Throw-ContractError "candidate source contract key must not refer to .tmp: $candidatePath"
    }
    Assert-NoReparsePointInExistingPath -PathValue $candidatePath -Label 'candidate source'
    Assert-RegularFile -PathValue $candidatePath -Label 'candidate source'
    if (-not $candidatePaths.Add($candidatePath)) {
        Throw-ContractError "duplicate candidate source contract key: $candidatePath"
    }

    $actualCandidateHash = Assert-ExpectedSha256 -PathValue $candidatePath -ExpectedHash $parsed.Hash -Label 'candidate source'
    $candidateRecords.Add([pscustomobject]@{
        relativePath = $parsed.Key.Replace('\', '/')
        path = $candidatePath
        expectedSha256 = $parsed.Hash
        actualSha256 = $actualCandidateHash
    })
}

$nativeDirectory = Join-Path -Path $repoRoot -ChildPath 'third_party\ladybugdb\v0.19.1\windows-x86_64'
Assert-NoReparsePointInExistingPath -PathValue $nativeDirectory -Label 'Ladybug native directory'
Assert-RegularDirectory -PathValue $nativeDirectory -Label 'Ladybug native directory'

$requiredNativeNames = @('lbug.h', 'lbug_shared.lib', 'lbug_shared.dll')
$expectedNativeHashes = [System.Collections.Generic.Dictionary[string, string]]::new(
    [System.StringComparer]::OrdinalIgnoreCase
)
foreach ($entry in $ExpectedNativeHash) {
    $parsed = Split-HashContractEntry -Entry $entry -Label 'ExpectedNativeHash'
    if ([System.IO.Path]::GetFileName($parsed.Key) -cne $parsed.Key) {
        Throw-ContractError "native hash contract key must be a filename only: $($parsed.Key)"
    }
    if (-not ($requiredNativeNames -contains $parsed.Key)) {
        Throw-ContractError "native hash contract contains unsupported filename: $($parsed.Key)"
    }
    if (-not $expectedNativeHashes.TryAdd($parsed.Key, $parsed.Hash)) {
        Throw-ContractError "duplicate native hash contract key: $($parsed.Key)"
    }
}
if ($expectedNativeHashes.Count -ne $requiredNativeNames.Count) {
    Throw-ContractError (
        "native hash contract must contain exactly $($requiredNativeNames -join ', ')"
    )
}

$nativeRecords = [System.Collections.Generic.List[object]]::new()
foreach ($nativeName in $requiredNativeNames) {
    if (-not $expectedNativeHashes.ContainsKey($nativeName)) {
        Throw-ContractError "native hash contract is missing: $nativeName"
    }
    $nativePath = Join-Path -Path $nativeDirectory -ChildPath $nativeName
    Assert-NoReparsePointInExistingPath -PathValue $nativePath -Label "native input $nativeName"
    Assert-RegularFile -PathValue $nativePath -Label "native input $nativeName"
    $actualNativeHash = Assert-ExpectedSha256 -PathValue $nativePath -ExpectedHash $expectedNativeHashes[$nativeName] -Label "native input $nativeName"
    $nativeRecords.Add([pscustomobject]@{
        fileName = $nativeName
        path = $nativePath
        expectedSha256 = $expectedNativeHashes[$nativeName]
        actualSha256 = $actualNativeHash
    })
}

if ([string]::IsNullOrWhiteSpace($GoExecutable)) {
    $goCommand = Get-Command -Name 'go' -CommandType Application -ErrorAction Stop |
        Select-Object -First 1
    $goPath = [System.IO.Path]::GetFullPath($goCommand.Source)
}
else {
    if (-not [System.IO.Path]::IsPathFullyQualified($GoExecutable)) {
        Throw-ContractError "GoExecutable must be absolute when supplied: $GoExecutable"
    }
    $goPath = Get-NormalizedFullPath -PathValue $GoExecutable -BasePath $repoRoot
}
Assert-NoReparsePointInExistingPath -PathValue $goPath -Label 'Go executable'
Assert-RegularFile -PathValue $goPath -Label 'Go executable'

$attemptWorkRoot = Join-Path -Path $outputDirectory -ChildPath ".a00x-$AttemptId-work"
$workDirectories = [ordered]@{
    root = $attemptWorkRoot
    goCache = (Join-Path -Path $attemptWorkRoot -ChildPath 'gocache')
    goModCache = (Join-Path -Path $attemptWorkRoot -ChildPath 'gomodcache')
    goTmp = (Join-Path -Path $attemptWorkRoot -ChildPath 'gotmp')
    temp = (Join-Path -Path $attemptWorkRoot -ChildPath 'temp')
    tmp = (Join-Path -Path $attemptWorkRoot -ChildPath 'tmp')
    home = (Join-Path -Path $attemptWorkRoot -ChildPath 'home')
    appData = (Join-Path -Path $attemptWorkRoot -ChildPath 'home\AppData\Roaming')
    localAppData = (Join-Path -Path $attemptWorkRoot -ChildPath 'home\AppData\Local')
    xdgCache = (Join-Path -Path $attemptWorkRoot -ChildPath 'home\.cache')
    xdgConfig = (Join-Path -Path $attemptWorkRoot -ChildPath 'home\.config')
}

foreach ($workDirectoryName in $workDirectories.Keys) {
    $workDirectoryPath = [string]$workDirectories[$workDirectoryName]
    if (-not (Test-PathAtOrBelow -CandidatePath $workDirectoryPath -RootPath $outputDirectory)) {
        Throw-ContractError (
            "attempt-local work directory escapes the verified output boundary: $workDirectoryPath"
        )
    }
    if (-not (Test-PathAtOrBelow -CandidatePath $workDirectoryPath -RootPath $temporaryRoot)) {
        Throw-ContractError (
            "attempt-local work directory escapes repository .tmp: $workDirectoryPath"
        )
    }
    Assert-NoReparsePointInExistingPath -PathValue $workDirectoryPath -Label "attempt-local $workDirectoryName directory"
    if (Test-Path -LiteralPath $workDirectoryPath) {
        Throw-ContractError (
            "refusing to reuse existing attempt-local $workDirectoryName directory: $workDirectoryPath"
        )
    }
}

$repoHeadOutput = @(& git -C $repoRoot rev-parse HEAD 2>&1)
if ($LASTEXITCODE -ne 0 -or $repoHeadOutput.Count -ne 1) {
    Throw-ContractError 'failed to read repository HEAD'
}
$repoHead = $repoHeadOutput[0].ToString().Trim()
if ($repoHead -notmatch '^[A-Fa-f0-9]{40}$') {
    Throw-ContractError "repository HEAD is not a full commit hash: $repoHead"
}
$repoStatusAtStart = @(& git -C $repoRoot status --porcelain=v1 --untracked-files=all)
if ($LASTEXITCODE -ne 0) {
    Throw-ContractError 'failed to read repository status'
}

[void][System.IO.Directory]::CreateDirectory($outputDirectory)
Assert-NoReparsePointInExistingPath -PathValue $outputDirectory -Label 'created output directory'
Assert-RegularDirectory -PathValue $outputDirectory -Label 'created output directory'
foreach ($workDirectoryName in $workDirectories.Keys) {
    $workDirectoryPath = [string]$workDirectories[$workDirectoryName]
    [void][System.IO.Directory]::CreateDirectory($workDirectoryPath)
    Assert-NoReparsePointInExistingPath -PathValue $workDirectoryPath -Label "created attempt-local $workDirectoryName directory"
    Assert-RegularDirectory -PathValue $workDirectoryPath -Label "created attempt-local $workDirectoryName directory"
}
if (Test-Path -LiteralPath $outputFullPath) {
    Throw-ContractError "output executable appeared before build: $outputFullPath"
}
if (Test-Path -LiteralPath $outputDllPath) {
    Throw-ContractError "output DLL appeared before build: $outputDllPath"
}
if (Test-Path -LiteralPath $provenancePath) {
    Throw-ContractError "provenance JSON appeared before build: $provenancePath"
}

$existingPath = [System.Environment]::GetEnvironmentVariable(
    'PATH',
    [System.EnvironmentVariableTarget]::Process
)
$scopedEnvironment = [ordered]@{
    GOENV = 'off'
    GOWORK = 'off'
    GOFLAGS = ''
    GOPROXY = 'off'
    GOSUMDB = 'off'
    GOTOOLCHAIN = 'local'
    GOPRIVATE = ''
    GONOPROXY = 'none'
    GONOSUMDB = 'none'
    GOINSECURE = ''
    GOVCS = '*:off'
    CGO_ENABLED = '1'
    CGO_CFLAGS = "-I$nativeDirectory"
    CGO_LDFLAGS = "-L$nativeDirectory -llbug_shared"
    GOCACHE = $workDirectories.goCache
    GOMODCACHE = $workDirectories.goModCache
    GOTMPDIR = $workDirectories.goTmp
    TEMP = $workDirectories.temp
    TMP = $workDirectories.tmp
    HOME = $workDirectories.home
    USERPROFILE = $workDirectories.home
    APPDATA = $workDirectories.appData
    LOCALAPPDATA = $workDirectories.localAppData
    XDG_CACHE_HOME = $workDirectories.xdgCache
    XDG_CONFIG_HOME = $workDirectories.xdgConfig
    PATH = "$nativeDirectory$([System.IO.Path]::PathSeparator)$existingPath"
}

$goVersionResult = Invoke-CapturedProcess -FilePath $goPath -Arguments @('version') -WorkingDirectory $repoRoot -EnvironmentOverrides $scopedEnvironment
if ($goVersionResult.ExitCode -ne 0) {
    Throw-ContractError (
        "Go version command failed with exit $($goVersionResult.ExitCode): " +
        $goVersionResult.StandardError.Trim()
    )
}
$goVersion = (
    $goVersionResult.StandardOutput + $goVersionResult.StandardError
).Trim()
if ([string]::IsNullOrWhiteSpace($goVersion)) {
    Throw-ContractError 'Go version command returned no version text'
}

$buildArguments = @(
    'build',
    '-mod=vendor',
    '-tags',
    'ladybugdb',
    '-trimpath',
    '-buildvcs=false',
    '-ldflags=-s -w',
    '-overlay',
    $manifestFullPath,
    '-o',
    $outputFullPath,
    './cmd/anvien'
)

$buildResult = Invoke-CapturedProcess -FilePath $goPath -Arguments $buildArguments -WorkingDirectory $repoRoot -EnvironmentOverrides $scopedEnvironment
if (-not [string]::IsNullOrEmpty($buildResult.StandardOutput)) {
    [System.Console]::Out.Write($buildResult.StandardOutput)
}
if (-not [string]::IsNullOrEmpty($buildResult.StandardError)) {
    [System.Console]::Error.Write($buildResult.StandardError)
}
if ($buildResult.ExitCode -ne 0) {
    Throw-ContractError (
        "Go build failed with exit $($buildResult.ExitCode); stdout: " +
        $buildResult.StandardOutput.Trim() +
        '; stderr: ' +
        $buildResult.StandardError.Trim()
    )
}

Assert-NoReparsePointInExistingPath -PathValue $outputFullPath -Label 'built executable'
Assert-RegularFile -PathValue $outputFullPath -Label 'built executable'

$nativeDllSourcePath = Join-Path -Path $nativeDirectory -ChildPath 'lbug_shared.dll'
[System.IO.File]::Copy($nativeDllSourcePath, $outputDllPath, $false)
Assert-NoReparsePointInExistingPath -PathValue $outputDllPath -Label 'copied native DLL'
Assert-RegularFile -PathValue $outputDllPath -Label 'copied native DLL'
$copiedDllHash = Assert-ExpectedSha256 -PathValue $outputDllPath -ExpectedHash $expectedNativeHashes['lbug_shared.dll'] -Label 'copied native DLL'

$versionResult = Invoke-CapturedProcess -FilePath $outputFullPath -Arguments @('version') -WorkingDirectory $outputDirectory -EnvironmentOverrides $scopedEnvironment
if ($versionResult.ExitCode -ne 0) {
    Throw-ContractError (
        "built executable version command failed with exit $($versionResult.ExitCode): " +
        $versionResult.StandardError.Trim()
    )
}
$outputVersion = (
    $versionResult.StandardOutput + $versionResult.StandardError
).Trim()
if ([string]::IsNullOrWhiteSpace($outputVersion)) {
    Throw-ContractError 'built executable version command returned no version text'
}

$outputExecutableItem = Get-Item -LiteralPath $outputFullPath -Force
$outputDllItem = Get-Item -LiteralPath $outputDllPath -Force
$outputExecutableHash = Get-Sha256 -PathValue $outputFullPath
$outputDllVersion = [System.Diagnostics.FileVersionInfo]::GetVersionInfo(
    $outputDllPath
).FileVersion
$outputExecutableFileVersion = [System.Diagnostics.FileVersionInfo]::GetVersionInfo(
    $outputFullPath
).FileVersion

$repoStatusAtEnd = @(& git -C $repoRoot status --porcelain=v1 --untracked-files=all)
if ($LASTEXITCODE -ne 0) {
    Throw-ContractError 'failed to read repository status after build'
}

$invocationEndedAtUtc = [System.DateTime]::UtcNow
$provenance = [ordered]@{
    schemaVersion = 1
    attemptId = $AttemptId
    startedAtUtc = $invocationStartedAtUtc.ToString('O')
    endedAtUtc = $invocationEndedAtUtc.ToString('O')
    repository = [ordered]@{
        root = $repoRoot
        head = $repoHead
        statusAtStart = @($repoStatusAtStart)
        statusAtEnd = @($repoStatusAtEnd)
    }
    go = [ordered]@{
        path = $goPath
        version = $goVersion
    }
    workDirectories = $workDirectories
    build = [ordered]@{
        workingDirectory = $repoRoot
        argv = @($goPath) + $buildArguments
        flags = [ordered]@{
            mod = 'vendor'
            tags = @('ladybugdb')
            trimpath = $true
            buildvcs = $false
            ldflags = '-s -w'
            overlay = $manifestFullPath
            output = $outputFullPath
            package = './cmd/anvien'
        }
        startedAtUtc = $buildResult.StartedAtUtc.ToString('O')
        endedAtUtc = $buildResult.EndedAtUtc.ToString('O')
        exitCode = $buildResult.ExitCode
    }
    scopedEnvironment = $scopedEnvironment
    overlay = [ordered]@{
        manifestPath = $manifestFullPath
        expectedSha256 = $normalizedExpectedOverlayHash
        actualSha256 = $actualOverlayHash
        mappings = @($mappingRecords)
    }
    candidateSources = @($candidateRecords)
    nativeInputs = @($nativeRecords)
    outputs = [ordered]@{
        executable = [ordered]@{
            path = $outputFullPath
            bytes = $outputExecutableItem.Length
            sha256 = $outputExecutableHash
            version = $outputVersion
            fileVersion = $outputExecutableFileVersion
        }
        dll = [ordered]@{
            path = $outputDllPath
            bytes = $outputDllItem.Length
            sha256 = $copiedDllHash
            fileVersion = $outputDllVersion
        }
        provenance = [ordered]@{
            path = $provenancePath
        }
    }
    exitCode = 0
}

$utf8NoBom = [System.Text.UTF8Encoding]::new($false)
$provenanceJson = $provenance | ConvertTo-Json -Depth 12
[System.IO.File]::WriteAllText(
    $provenancePath,
    $provenanceJson + [System.Environment]::NewLine,
    $utf8NoBom
)
Assert-NoReparsePointInExistingPath -PathValue $provenancePath -Label 'provenance JSON'
Assert-RegularFile -PathValue $provenancePath -Label 'provenance JSON'

[pscustomobject]@{
    status = 'A00X_BENCHMARK_BUILD_COMPLETE'
    attemptId = $AttemptId
    executablePath = $outputFullPath
    dllPath = $outputDllPath
    provenancePath = $provenancePath
    executableSha256 = $outputExecutableHash
    version = $outputVersion
    exitCode = 0
} | ConvertTo-Json -Compress
