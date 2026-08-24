[CmdletBinding()]
param(
    [Parameter(Mandatory)][string]$LaneRoot,
    [string]$RepositoryRoot = (Split-Path -Parent $PSScriptRoot),
    [string]$SourceProxy = 'https://proxy.golang.org',
    [string]$InputProxy = ''
)

$ErrorActionPreference = 'Stop'
Set-StrictMode -Version Latest

$ExpectedGoVersion = 'go version go1.26.3 windows/amd64'
$ExpectedGoModSha256 = 'C203E25E0A83A583E89A9D8F8E65BC0881052B3075E46915F90A4FB6433BD249'
$ExpectedGoSumSha256 = '5434F39B08F50F1C53A71333E2E971CFCDE0E80BCA7DD028C00E31B9ACCAEC73'
$AllowedProxy = 'https://proxy.golang.org'
$PatchAuthorityRelativePath = 'third_party/go-vendor/patches/tree-sitter-go-v0.25.0-remove-absent-scanner.patch'
$ExpectedPatchSha256 = '43BB195FFF439C7DCD0D057094EC1169304553E357F6D0B2781DC5713516BEEA'
$PatchedModulePath = 'github.com/tree-sitter/tree-sitter-go'
$PatchedModuleVersion = 'v0.25.0'
$PatchedBindingRelativePath = 'github.com/tree-sitter/tree-sitter-go/bindings/go/binding.go'
$ExpectedPatchPreimageBytes = 333
$ExpectedPatchPreimageSha256 = '680E50820F2F7429FEAC88B69DDCEE67735E6E692DFD1F98147109D77D296096'
$ExpectedPatchPostimageBytes = 245
$ExpectedPatchPostimageSha256 = '5A456F352CFEF15430F3FCFB155A3E1BDE38F244E5F8537026553B1F21547713'
$ExpectedGrammarBytes = 198042
$ExpectedGrammarSha256 = '0DCC665AA521A1B73A200BE13B4E376780CCA037D764A165F12864BABFE28000'
$ExpectedParserBytes = 1572685
$ExpectedParserSha256 = '3DBF6ED1238B5DFCF2BE4D2F2D4CB27A14D34F34D7784ECCCCBFD532FD4A6D85'
$Utf8NoBom = [System.Text.UTF8Encoding]::new($false)
$Utf8NoBomStrict = [System.Text.UTF8Encoding]::new($false, $true)

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

function ConvertTo-CanonicalModulePath {
    param([string]$BaseDirectory = '', [Parameter(Mandatory)][string]$RelativePath)
    if ([string]::IsNullOrWhiteSpace($RelativePath) -or $RelativePath.Contains([char]0) -or
        $RelativePath.Contains('\') -or $RelativePath.StartsWith('/', [System.StringComparison]::Ordinal) -or
        $RelativePath -match '^[A-Za-z]:' -or $RelativePath.StartsWith('//', [System.StringComparison]::Ordinal)) {
        throw "Unsafe module-local path: $RelativePath"
    }
    $segments = [System.Collections.Generic.List[string]]::new()
    if (-not [string]::IsNullOrWhiteSpace($BaseDirectory) -and $BaseDirectory -cne '.') {
        foreach ($segment in ($BaseDirectory -split '/')) {
            if ([string]::IsNullOrEmpty($segment) -or $segment -ceq '.' -or $segment -ceq '..') {
                throw "Unsafe module base directory: $BaseDirectory"
            }
            $segments.Add($segment)
        }
    }
    foreach ($segment in ($RelativePath -split '/')) {
        if ([string]::IsNullOrEmpty($segment) -or $segment -ceq '.') { continue }
        if ($segment -ceq '..') {
            if ($segments.Count -eq 0) { throw "Module-local path escapes its module: $RelativePath" }
            $segments.RemoveAt($segments.Count - 1)
            continue
        }
        $segments.Add($segment)
    }
    if ($segments.Count -eq 0) { throw "Module-local path resolved to the module root: $RelativePath" }
    return ($segments -join '/')
}

function Get-Sha256 {
    param([Parameter(Mandatory)][string]$LiteralPath)
    return (Get-FileHash -Algorithm SHA256 -LiteralPath $LiteralPath).Hash.ToUpperInvariant()
}

function Write-Utf8Lf {
    param([Parameter(Mandatory)][string]$LiteralPath, [AllowEmptyString()][string]$Text)
    $parent = Split-Path -Parent $LiteralPath
    if ($parent) { [System.IO.Directory]::CreateDirectory($parent) | Out-Null }
    $normalized = $Text.Replace("`r`n", "`n").Replace("`r", "`n")
    [System.IO.File]::WriteAllText($LiteralPath, $normalized, $Utf8NoBom)
}

function Read-StrictUtf8Text {
    param([Parameter(Mandatory)][string]$LiteralPath)
    $bytes = [System.IO.File]::ReadAllBytes($LiteralPath)
    if ($bytes.Length -ge 3 -and $bytes[0] -eq 0xEF -and $bytes[1] -eq 0xBB -and $bytes[2] -eq 0xBF) {
        throw "UTF-8 BOM is forbidden: $LiteralPath"
    }
    return $Utf8NoBomStrict.GetString($bytes)
}

function Assert-ModuleIdentity {
    param([Parameter(Mandatory)][string]$Root)
    $goModHash = Get-Sha256 (Join-Path $Root 'go.mod')
    $goSumHash = Get-Sha256 (Join-Path $Root 'go.sum')
    if ($goModHash -cne $ExpectedGoModSha256) {
        throw "go.mod drift: expected $ExpectedGoModSha256, actual $goModHash"
    }
    if ($goSumHash -cne $ExpectedGoSumSha256) {
        throw "go.sum drift: expected $ExpectedGoSumSha256, actual $goSumHash"
    }
}

function Copy-ModuleAuthority {
    param([Parameter(Mandatory)][string]$Destination)
    [System.IO.Directory]::CreateDirectory($Destination) | Out-Null
    [System.IO.File]::WriteAllBytes(
        (Join-Path $Destination 'go.mod'),
        [System.IO.File]::ReadAllBytes((Join-Path $script:RepoRoot 'go.mod'))
    )
    [System.IO.File]::WriteAllBytes(
        (Join-Path $Destination 'go.sum'),
        [System.IO.File]::ReadAllBytes((Join-Path $script:RepoRoot 'go.sum'))
    )
}

function Initialize-StateRoot {
    param([Parameter(Mandatory)][string]$StateRoot)
    foreach ($name in @(
        'home', 'userprofile', 'temp', 'tmp', 'tmpdir', 'appdata', 'localappdata',
        'xdg-config', 'xdg-cache', 'gocache', 'gomodcache', 'gopath'
    )) {
        [System.IO.Directory]::CreateDirectory((Join-Path $StateRoot $name)) | Out-Null
    }
}

function Set-GoEnvironment {
    param([Parameter(Mandatory)][string]$StateRoot, [Parameter(Mandatory)][string]$GoProxy)
    Initialize-StateRoot $StateRoot
    $env:HOME = Join-Path $StateRoot 'home'
    $env:USERPROFILE = Join-Path $StateRoot 'userprofile'
    $env:TEMP = Join-Path $StateRoot 'temp'
    $env:TMP = Join-Path $StateRoot 'tmp'
    $env:TMPDIR = Join-Path $StateRoot 'tmpdir'
    $env:APPDATA = Join-Path $StateRoot 'appdata'
    $env:LOCALAPPDATA = Join-Path $StateRoot 'localappdata'
    $env:XDG_CONFIG_HOME = Join-Path $StateRoot 'xdg-config'
    $env:XDG_CACHE_HOME = Join-Path $StateRoot 'xdg-cache'
    $env:GOCACHE = Join-Path $StateRoot 'gocache'
    $env:GOMODCACHE = Join-Path $StateRoot 'gomodcache'
    $env:GOPATH = Join-Path $StateRoot 'gopath'
    $env:GOENV = 'off'
    $env:GOWORK = 'off'
    $env:GOTOOLCHAIN = 'local'
    $env:GOTELEMETRY = 'off'
    $env:GO111MODULE = 'on'
    $env:GOFLAGS = ''
    $env:GOPROXY = $GoProxy
    $env:GOSUMDB = 'off'
    $env:GONOPROXY = 'none'
    $env:GONOSUMDB = 'none'
    $env:GOPRIVATE = ''
    $env:GOINSECURE = ''
    $env:GOVCS = '*:off'
    $env:HTTP_PROXY = ''
    $env:HTTPS_PROXY = ''
    $env:ALL_PROXY = ''
    $env:NO_PROXY = ''
    $env:GIT_TERMINAL_PROMPT = '0'
    $env:GIT_CONFIG_NOSYSTEM = '1'
}

function Invoke-GoLogged {
    param(
        [Parameter(Mandatory)][string]$Id,
        [Parameter(Mandatory)][string]$WorkingDirectory,
        [Parameter(Mandatory)][string[]]$Arguments
    )
    $stdoutPath = Join-Path $script:LogsRoot ($Id + '.stdout.log')
    $stderrPath = Join-Path $script:LogsRoot ($Id + '.stderr.log')
    $commandPath = Join-Path $script:LogsRoot ($Id + '.command.json')
    $resultPath = Join-Path $script:LogsRoot ($Id + '.result.json')
    $command = [ordered]@{
        executable = 'go'
        arguments = $Arguments
        workingDirectory = $WorkingDirectory
        environment = [ordered]@{
            HOME = $env:HOME; USERPROFILE = $env:USERPROFILE; TEMP = $env:TEMP; TMP = $env:TMP
            APPDATA = $env:APPDATA; LOCALAPPDATA = $env:LOCALAPPDATA
            GOCACHE = $env:GOCACHE; GOMODCACHE = $env:GOMODCACHE; GOPATH = $env:GOPATH
            GOENV = $env:GOENV; GOWORK = $env:GOWORK; GOTOOLCHAIN = $env:GOTOOLCHAIN
            GOPROXY = $env:GOPROXY; GOSUMDB = $env:GOSUMDB; GONOPROXY = $env:GONOPROXY
            GONOSUMDB = $env:GONOSUMDB; GOPRIVATE = $env:GOPRIVATE; GOINSECURE = $env:GOINSECURE
            GOVCS = $env:GOVCS; GOFLAGS = $env:GOFLAGS
        }
    }
    Write-Utf8Lf $commandPath (($command | ConvertTo-Json -Depth 6) + "`n")
    $started = [DateTimeOffset]::UtcNow
    Push-Location $WorkingDirectory
    & go @Arguments 1> $stdoutPath 2> $stderrPath
    $exitCode = $LASTEXITCODE
    Pop-Location
    $ended = [DateTimeOffset]::UtcNow
    $result = [ordered]@{
        exitCode = $exitCode
        startedUtc = $started.ToString('o')
        endedUtc = $ended.ToString('o')
        durationMilliseconds = [math]::Round(($ended - $started).TotalMilliseconds, 3)
        stdout = (Get-StrictRelativePath -BasePath $script:Lane -ChildPath $stdoutPath).Replace('\', '/')
        stderr = (Get-StrictRelativePath -BasePath $script:Lane -ChildPath $stderrPath).Replace('\', '/')
    }
    Write-Utf8Lf $resultPath (($result | ConvertTo-Json) + "`n")
    if ($exitCode -ne 0) {
        $stderr = if (Test-Path -LiteralPath $stderrPath) { Read-StrictUtf8Text $stderrPath } else { '' }
        throw "Go command failed ($Id), exit $exitCode. $stderr"
    }
    return [pscustomobject]@{
        ExitCode = $exitCode
        StdoutPath = $stdoutPath
        StderrPath = $stderrPath
        Stdout = Read-StrictUtf8Text $stdoutPath
        Result = $result
    }
}

function ConvertFrom-JsonSequence {
    param([Parameter(Mandatory)][string]$Text)
    $trimmed = $Text.Trim()
    if ([string]::IsNullOrWhiteSpace($trimmed)) { return @() }
    $arrayText = '[' + [regex]::Replace($trimmed, '(?m)^}\r?\n^{', "},`n{") + ']'
    return @($arrayText | ConvertFrom-Json)
}

function Get-InventoryRows {
    param([Parameter(Mandatory)][string]$Root, [string[]]$RelativeRoots = @('.'))
    $rows = [System.Collections.Generic.List[object]]::new()
    foreach ($relativeRoot in $RelativeRoots) {
        $absoluteRoot = if ($relativeRoot -ceq '.') { $Root } else { Join-Path $Root $relativeRoot }
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

function Assert-NoReparsePoints {
    param([Parameter(Mandatory)][string]$Root, [Parameter(Mandatory)][string]$Label)
    $points = @(@(
        Get-Item -LiteralPath $Root -Force
        Get-ChildItem -LiteralPath $Root -Force -Recurse
    ) | Where-Object { ($_.Attributes -band [System.IO.FileAttributes]::ReparsePoint) -ne 0 })
    if ($points.Count -ne 0) { throw "$Label contains $($points.Count) reparse point(s)." }
}

function Get-VendorModules {
    param([Parameter(Mandatory)][string]$ModulesText)
    $rows = [System.Collections.Generic.List[object]]::new()
    foreach ($line in ($ModulesText -split "`n")) {
        $trimmed = $line.TrimEnd("`r")
        if ($trimmed -match '^#\s+(\S+)\s+(\S+)(?:\s+=>.*)?$') {
            if ($trimmed -match '\s=>\s') { throw "Module replacement is forbidden: $trimmed" }
            $rows.Add([pscustomobject][ordered]@{ path = $Matches[1]; version = $Matches[2] })
        }
    }
    return @(Sort-ModuleRowsOrdinal @($rows))
}

function Sort-ModuleRowsOrdinal {
    param([Parameter(Mandatory)][object[]]$Rows)
    $sorted = [System.Collections.Generic.SortedDictionary[string, object]]::new([System.StringComparer]::Ordinal)
    foreach ($row in $Rows) {
        $key = "$([string]$row.path)@$([string]$row.version)"
        if ($sorted.ContainsKey($key)) { throw "Duplicate module identity: $key" }
        $sorted.Add($key, $row)
    }
    return @($sorted.Values)
}

function Convert-ToFileProxyUri {
    param([Parameter(Mandatory)][string]$Path)
    $normalized = [System.IO.Path]::GetFullPath($Path).Replace('\', '/')
    if ($normalized -notmatch '^[A-Za-z]:/') { throw "Expected Windows absolute path: $Path" }
    return 'file:///' + $normalized
}

function Get-CanonicalStringDigest {
    param([Parameter(Mandatory)][string[]]$Rows)
    $text = if ($Rows.Count -eq 0) { '' } else { ($Rows -join "`n") + "`n" }
    return Get-BytesSha256Hex -Bytes $Utf8NoBom.GetBytes($text)
}

function ConvertTo-GoCacheEscapedPath {
    param([Parameter(Mandatory)][string]$Value)
    $builder = [System.Text.StringBuilder]::new()
    foreach ($character in $Value.ToCharArray()) {
        if ($character -ge [char]'A' -and $character -le [char]'Z') {
            [void]$builder.Append('!')
            [void]$builder.Append([char]([int]$character + 32))
        }
        else { [void]$builder.Append($character) }
    }
    return $builder.ToString()
}

function Get-ModuleForVendorPath {
    param([Parameter(Mandatory)][string]$VendorRelativePath, [Parameter(Mandatory)][object[]]$Modules)
    $matches = @($Modules | Where-Object {
        $prefix = ([string]$_.path).Replace('\', '/')
        $VendorRelativePath -ceq $prefix -or $VendorRelativePath.StartsWith($prefix + '/', [System.StringComparison]::Ordinal)
    } | Sort-Object { ([string]$_.path).Length } -Descending)
    if ($matches.Count -eq 0) { return $null }
    if ($matches.Count -gt 1 -and ([string]$matches[0].path).Length -eq ([string]$matches[1].path).Length) {
        throw "Ambiguous module owner for vendor path: $VendorRelativePath"
    }
    return $matches[0]
}

function New-ModuleZipIndex {
    param([Parameter(Mandatory)][object]$Module)
    $archive = [System.IO.Compression.ZipFile]::OpenRead([string]$Module.zipPath)
    try {
        $entries = [System.Collections.Generic.Dictionary[string, object]]::new([System.StringComparer]::Ordinal)
        $caseInsensitive = [System.Collections.Generic.Dictionary[string, string]]::new([System.StringComparer]::OrdinalIgnoreCase)
        $archivePrefix = "$([string]$Module.path)@$([string]$Module.version)"
        $archivePrefixWithSlash = $archivePrefix + '/'
        foreach ($entry in @($archive.Entries)) {
            $fullName = [string]$entry.FullName
            if ($fullName.EndsWith('/', [System.StringComparison]::Ordinal)) { continue }
            if ($fullName.Contains('\') -or $fullName.Contains([char]0)) { throw "Malformed module zip entry: $fullName" }
            if (-not $fullName.StartsWith($archivePrefixWithSlash, [System.StringComparison]::Ordinal) -or $fullName.Length -eq $archivePrefixWithSlash.Length) {
                throw "Malformed module zip entry root: $fullName"
            }
            $relative = $fullName.Substring($archivePrefixWithSlash.Length)
            Assert-CanonicalRelativePath $relative
            if ($caseInsensitive.ContainsKey($relative)) { throw "Module zip has duplicate/case-colliding entry: $relative" }
            $caseInsensitive.Add($relative, $relative)
            $entries.Add($relative, $entry)
        }
        if ($entries.Count -eq 0) { throw "Module zip is empty: $($Module.path)@$($Module.version)" }
        return [pscustomobject]@{ archive = $archive; entries = $entries; entriesIgnoreCase = $caseInsensitive; prefix = $archivePrefix }
    }
    catch {
        $archive.Dispose()
        throw
    }
}

function Get-ZipEntryBytes {
    param([Parameter(Mandatory)][object]$ZipIndex, [Parameter(Mandatory)][string]$RelativePath)
    if (-not $ZipIndex.entries.ContainsKey($RelativePath)) { return $null }
    $stream = $ZipIndex.entries[$RelativePath].Open()
    $memory = [System.IO.MemoryStream]::new()
    try {
        $stream.CopyTo($memory)
        return $memory.ToArray()
    }
    finally {
        $stream.Dispose()
        $memory.Dispose()
    }
}

function Get-CgoPreambleLines {
    param([Parameter(Mandatory)][string]$Text, [Parameter(Mandatory)][string]$Label)
    $normalized = $Text.Replace("`r`n", "`n").Replace("`r", "`n")
    $lines = @($normalized -split "`n")
    $importIndexes = @()
    for ($index = 0; $index -lt $lines.Count; $index++) {
        if ($lines[$index] -cmatch '^\s*import\s+"C"\s*$') { $importIndexes += $index }
    }
    if ($importIndexes.Count -ne 1) { throw "Expected exactly one import C boundary in $Label; found $($importIndexes.Count)" }
    $end = $importIndexes[0] - 1
    if ($end -lt 0) { return @() }
    $result = [System.Collections.Generic.List[string]]::new()
    if ($lines[$end].TrimEnd().EndsWith('*/', [System.StringComparison]::Ordinal)) {
        $start = $end
        while ($start -ge 0 -and -not $lines[$start].Contains('/*')) { $start-- }
        if ($start -lt 0) { throw "Unterminated cgo preamble block in $Label" }
        $block = ($lines[$start..$end] -join "`n")
        $open = $block.IndexOf('/*', [System.StringComparison]::Ordinal)
        $close = $block.LastIndexOf('*/', [System.StringComparison]::Ordinal)
        if ($open -lt 0 -or $close -lt $open) { throw "Malformed cgo preamble block in $Label" }
        foreach ($line in @($block.Substring($open + 2, $close - $open - 2) -split "`n")) {
            $result.Add(([regex]::Replace($line, '^\s*\* ?', '')).TrimEnd())
        }
    }
    else {
        $start = $end
        while ($start -ge 0 -and $lines[$start] -match '^\s*//') { $start-- }
        $start++
        if ($start -gt $end) { return @() }
        foreach ($line in $lines[$start..$end]) {
            $result.Add(([regex]::Replace($line, '^\s*// ?', '')).TrimEnd())
        }
    }
    return @($result)
}

function Test-CgoConstraint {
    param([string]$Constraint)
    if ([string]::IsNullOrWhiteSpace($Constraint)) { return $true }
    $activeTags = [System.Collections.Generic.HashSet[string]]::new([System.StringComparer]::Ordinal)
    foreach ($tag in @('windows', 'amd64', 'cgo', 'gc')) { [void]$activeTags.Add($tag) }
    foreach ($orGroup in ($Constraint -split '\s+')) {
        if ([string]::IsNullOrWhiteSpace($orGroup)) { continue }
        $groupMatches = $true
        foreach ($term in ($orGroup -split ',')) {
            if ($term -cnotmatch '^!?[A-Za-z0-9_.]+$') { throw "Unsupported #cgo constraint: $Constraint" }
            $negated = $term.StartsWith('!', [System.StringComparison]::Ordinal)
            $tag = if ($negated) { $term.Substring(1) } else { $term }
            $present = $activeTags.Contains($tag)
            if (($negated -and $present) -or (-not $negated -and -not $present)) { $groupMatches = $false }
        }
        if ($groupMatches) { return $true }
    }
    return $false
}

function Read-PPPrimary {
    param([string[]]$Tokens, [ref]$Index, [System.Collections.Generic.HashSet[string]]$Defines)
    if ($Index.Value -ge $Tokens.Count) { throw 'Unexpected end of preprocessor expression.' }
    $token = $Tokens[$Index.Value]
    if ($token -ceq '(') {
        $Index.Value++
        $value = Read-PPOr $Tokens $Index $Defines
        if ($Index.Value -ge $Tokens.Count -or $Tokens[$Index.Value] -cne ')') { throw 'Unbalanced preprocessor expression.' }
        $Index.Value++
        return $value
    }
    if ($token -ceq 'defined') {
        $Index.Value++
        $parenthesized = $Index.Value -lt $Tokens.Count -and $Tokens[$Index.Value] -ceq '('
        if ($parenthesized) { $Index.Value++ }
        if ($Index.Value -ge $Tokens.Count -or $Tokens[$Index.Value] -cnotmatch '^[A-Za-z_][A-Za-z0-9_]*$') {
            throw 'Malformed defined(...) expression.'
        }
        $name = $Tokens[$Index.Value]
        $Index.Value++
        if ($parenthesized) {
            if ($Index.Value -ge $Tokens.Count -or $Tokens[$Index.Value] -cne ')') { throw 'Malformed defined(...) expression.' }
            $Index.Value++
        }
        return $Defines.Contains($name)
    }
    if ($token -ceq '0' -or $token -ceq '1') {
        $Index.Value++
        return $token -ceq '1'
    }
    if ($token -cmatch '^[A-Za-z_][A-Za-z0-9_]*$') {
        $Index.Value++
        return $Defines.Contains($token)
    }
    throw "Unsupported preprocessor token: $token"
}

function Read-PPUnary {
    param([string[]]$Tokens, [ref]$Index, [System.Collections.Generic.HashSet[string]]$Defines)
    if ($Index.Value -lt $Tokens.Count -and $Tokens[$Index.Value] -ceq '!') {
        $Index.Value++
        return -not (Read-PPUnary $Tokens $Index $Defines)
    }
    return Read-PPPrimary $Tokens $Index $Defines
}

function Read-PPAnd {
    param([string[]]$Tokens, [ref]$Index, [System.Collections.Generic.HashSet[string]]$Defines)
    $value = Read-PPUnary $Tokens $Index $Defines
    while ($Index.Value -lt $Tokens.Count -and $Tokens[$Index.Value] -ceq '&&') {
        $Index.Value++
        $right = Read-PPUnary $Tokens $Index $Defines
        $value = $value -and $right
    }
    return $value
}

function Read-PPOr {
    param([string[]]$Tokens, [ref]$Index, [System.Collections.Generic.HashSet[string]]$Defines)
    $value = Read-PPAnd $Tokens $Index $Defines
    while ($Index.Value -lt $Tokens.Count -and $Tokens[$Index.Value] -ceq '||') {
        $Index.Value++
        $right = Read-PPAnd $Tokens $Index $Defines
        $value = $value -or $right
    }
    return $value
}

function Test-PreprocessorExpression {
    param([Parameter(Mandatory)][string]$Expression, [Parameter(Mandatory)][System.Collections.Generic.HashSet[string]]$Defines)
    $matches = @([regex]::Matches($Expression, 'defined|&&|\|\||!|\(|\)|0|1|[A-Za-z_][A-Za-z0-9_]*'))
    $tokens = @($matches | ForEach-Object { $_.Value })
    $compact = [regex]::Replace($Expression, '\s+', '')
    if ($tokens.Count -eq 0 -or ($tokens -join '') -cne $compact) { throw "Unsupported preprocessor expression: $Expression" }
    $index = 0
    $value = Read-PPOr $tokens ([ref]$index) $Defines
    if ($index -ne $tokens.Count) { throw "Unsupported preprocessor expression tail: $Expression" }
    return [bool]$value
}

function Get-IncludeDirectives {
    param(
        [Parameter(Mandatory)][AllowEmptyString()][string[]]$Lines,
        [Parameter(Mandatory)][System.Collections.Generic.HashSet[string]]$BaseDefines,
        [Parameter(Mandatory)][string]$Label
    )
    $defines = [System.Collections.Generic.HashSet[string]]::new($BaseDefines, [System.StringComparer]::Ordinal)
    $stack = [System.Collections.Generic.List[object]]::new()
    $rows = [System.Collections.Generic.List[object]]::new()
    for ($lineIndex = 0; $lineIndex -lt $Lines.Count; $lineIndex++) {
        $line = $Lines[$lineIndex]
        $trimmed = $line.Trim()
        $parentState = if ($stack.Count -eq 0) { 'active' } else { [string]$stack[$stack.Count - 1].currentState }
        if ($trimmed -match '^#\s*(if|ifdef|ifndef)\b(.*)$') {
            $kind = $Matches[1]
            $expression = $Matches[2].Trim()
            $known = $true
            $condition = $false
            try {
                if ($kind -ceq 'ifdef') { $condition = $defines.Contains($expression) }
                elseif ($kind -ceq 'ifndef') { $condition = -not $defines.Contains($expression) }
                else { $condition = Test-PreprocessorExpression $expression $defines }
            }
            catch { $known = $false }
            $state = if ($parentState -ceq 'inactive') { 'inactive' } elseif ($parentState -ceq 'unknown' -or -not $known) { 'unknown' } elseif ($condition) { 'active' } else { 'inactive' }
            $stack.Add([pscustomobject]@{
                parentState = $parentState; currentState = $state; branchTaken = ($state -ceq 'active')
                unknownSeen = (-not $known -or $parentState -ceq 'unknown'); expression = $expression
            })
            continue
        }
        if ($trimmed -match '^#\s*elif\b(.*)$') {
            if ($stack.Count -eq 0) { throw "Unmatched #elif in $Label" }
            $frame = $stack[$stack.Count - 1]
            $expression = $Matches[1].Trim()
            if ($frame.parentState -ceq 'inactive' -or $frame.branchTaken) { $frame.currentState = 'inactive' }
            elseif ($frame.parentState -ceq 'unknown' -or $frame.unknownSeen) { $frame.currentState = 'unknown' }
            else {
                try {
                    $condition = Test-PreprocessorExpression $expression $defines
                    $frame.currentState = if ($condition) { 'active' } else { 'inactive' }
                    if ($condition) { $frame.branchTaken = $true }
                }
                catch { $frame.currentState = 'unknown'; $frame.unknownSeen = $true }
            }
            $frame.expression = $expression
            continue
        }
        if ($trimmed -match '^#\s*else\b') {
            if ($stack.Count -eq 0) { throw "Unmatched #else in $Label" }
            $frame = $stack[$stack.Count - 1]
            if ($frame.parentState -ceq 'inactive' -or $frame.branchTaken) { $frame.currentState = 'inactive' }
            elseif ($frame.parentState -ceq 'unknown' -or $frame.unknownSeen) { $frame.currentState = 'unknown' }
            else { $frame.currentState = 'active'; $frame.branchTaken = $true }
            $frame.expression = 'else'
            continue
        }
        if ($trimmed -match '^#\s*endif\b') {
            if ($stack.Count -eq 0) { throw "Unmatched #endif in $Label" }
            $stack.RemoveAt($stack.Count - 1)
            continue
        }
        $state = if ($stack.Count -eq 0) { 'active' } else { [string]$stack[$stack.Count - 1].currentState }
        $conditionText = if ($stack.Count -eq 0) { '' } else { [string]$stack[$stack.Count - 1].expression }
        if ($trimmed -match '^#\s*define\s+([A-Za-z_][A-Za-z0-9_]*)(?:\s|$)' -and $state -ceq 'active') {
            [void]$defines.Add($Matches[1]); continue
        }
        if ($trimmed -match '^#\s*undef\s+([A-Za-z_][A-Za-z0-9_]*)' -and $state -ceq 'active') {
            [void]$defines.Remove($Matches[1]); continue
        }
        if ($trimmed -match '^#\s*include\s*([<"])([^>"]+)[>"]') {
            $rows.Add([pscustomobject][ordered]@{
                line = $lineIndex + 1
                style = if ($Matches[1] -ceq '"') { 'quote' } else { 'angle' }
                token = $Matches[2]
                state = $state
                condition = $conditionText
                dynamic = $false
            })
            continue
        }
        if ($trimmed -match '^#\s*include\b') {
            $rows.Add([pscustomobject][ordered]@{
                line = $lineIndex + 1; style = 'dynamic'; token = $trimmed
                state = $state; condition = $conditionText; dynamic = $true
            })
        }
    }
    if ($stack.Count -ne 0) { throw "Unterminated preprocessor conditional in $Label" }
    return @($rows)
}

function Get-CgoClosureModels {
    param([Parameter(Mandatory)][string]$VendorRoot, [Parameter(Mandatory)][object[]]$Modules)
    $models = [System.Collections.Generic.List[object]]::new()
    $controlsByModule = @{}
    foreach ($goFile in @(Get-ChildItem -LiteralPath $VendorRoot -File -Force -Recurse -Filter '*.go' | Sort-Object FullName)) {
        $text = Read-StrictUtf8Text $goFile.FullName
        if ($text -cnotmatch '(?m)^\s*import\s+"C"\s*$') { continue }
        $vendorRelative = (Get-StrictRelativePath -BasePath $VendorRoot -ChildPath $goFile.FullName).Replace('\', '/')
        $module = Get-ModuleForVendorPath $vendorRelative $Modules
        if ($null -eq $module) { throw "No vendored module owns import-C file: $vendorRelative" }
        $modulePrefix = ([string]$module.path).Replace('\', '/')
        $moduleRelative = if ($vendorRelative -ceq $modulePrefix) { '' } else { $vendorRelative.Substring($modulePrefix.Length + 1) }
        Assert-CanonicalRelativePath $moduleRelative
        $preambleLines = @(Get-CgoPreambleLines $text $vendorRelative)
        $moduleID = "$($module.path)@$($module.version)"
        if (-not $controlsByModule.ContainsKey($moduleID)) {
            $defines = [System.Collections.Generic.HashSet[string]]::new([System.StringComparer]::Ordinal)
            foreach ($define in @('_WIN32', 'WIN32', '__x86_64__', '_M_X64')) { [void]$defines.Add($define) }
            $controlsByModule[$moduleID] = [pscustomobject]@{
                includeRoots = [System.Collections.Generic.HashSet[string]]::new([System.StringComparer]::Ordinal)
                defines = $defines
            }
        }
        $moduleControls = $controlsByModule[$moduleID]
        $fileDefines = [System.Collections.Generic.HashSet[string]]::new([System.StringComparer]::Ordinal)
        foreach ($define in @('_WIN32', 'WIN32', '__x86_64__', '_M_X64')) { [void]$fileDefines.Add($define) }
        $controls = [pscustomobject]@{
            includeRoots = [System.Collections.Generic.HashSet[string]]::new([System.StringComparer]::Ordinal)
            defines = $fileDefines
        }
        $fileDirectory = [System.IO.Path]::GetDirectoryName($moduleRelative).Replace('\', '/')
        foreach ($line in $preambleLines) {
            if ($line.Trim() -notmatch '^#cgo\s+(?:(.*?)\s+)?(CFLAGS|CPPFLAGS):\s*(.*)$') { continue }
            $constraint = [string]$Matches[1]
            $kind = [string]$Matches[2]
            $flags = [string]$Matches[3]
            if ($kind -cne 'CFLAGS' -and $kind -cne 'CPPFLAGS') { continue }
            if (-not (Test-CgoConstraint $constraint)) { continue }
            if ($flags.Contains('$') -or $flags.Contains('`') -or $flags.Contains("'")) {
                throw "Dynamic/quoted cgo flags are unsupported in ${vendorRelative}: $flags"
            }
            $tokens = @([regex]::Matches($flags, '"[^"]*"|\S+') | ForEach-Object { $_.Value.Trim('"') })
            for ($index = 0; $index -lt $tokens.Count; $index++) {
                $token = $tokens[$index]
                $includeValue = $null
                if ($token -ceq '-I') {
                    $index++
                    if ($index -ge $tokens.Count) { throw "Missing -I operand in $vendorRelative" }
                    $includeValue = $tokens[$index]
                }
                elseif ($token.StartsWith('-I', [System.StringComparison]::Ordinal) -and $token.Length -gt 2) {
                    $includeValue = $token.Substring(2)
                }
                if ($null -ne $includeValue) {
                    $includeRoot = ConvertTo-CanonicalModulePath $fileDirectory $includeValue
                    [void]$controls.includeRoots.Add($includeRoot)
                    [void]$moduleControls.includeRoots.Add($includeRoot)
                    continue
                }
                if ($token -cmatch '^-D([A-Za-z_][A-Za-z0-9_]*)(?:=.*)?$') {
                    [void]$controls.defines.Add($Matches[1]); [void]$moduleControls.defines.Add($Matches[1]); continue
                }
                if ($token -cmatch '^-U([A-Za-z_][A-Za-z0-9_]*)$') {
                    [void]$controls.defines.Remove($Matches[1]); [void]$moduleControls.defines.Remove($Matches[1]); continue
                }
            }
        }
        if ($controls.defines.Contains('TREE_SITTER_FEATURE_WASM')) {
            throw "TREE_SITTER_FEATURE_WASM is forbidden in effective cgo flags for $moduleID"
        }
        $models.Add([pscustomobject]@{
            module = $module
            moduleID = $moduleID
            moduleRelativeFile = $moduleRelative
            preambleLines = $preambleLines
            controls = $controls
        })
    }
    return [pscustomobject]@{ models = @($models); controlsByModule = $controlsByModule }
}

function Resolve-ModuleInclude {
    param(
        [Parameter(Mandatory)][object]$Module,
        [Parameter(Mandatory)][object]$ZipIndex,
        [Parameter(Mandatory)][object]$Controls,
        [Parameter(Mandatory)][string]$VendorRoot,
        [Parameter(Mandatory)][string]$Referencer,
        [Parameter(Mandatory)][object]$Directive
    )
    if ([bool]$Directive.dynamic) {
        if ([string]$Directive.state -cne 'inactive') { throw "Dynamic include is forbidden: $($Module.path)/${Referencer}:$($Directive.line) $($Directive.token)" }
        return [pscustomobject]@{ kind = 'inactive_dynamic'; path = ''; root = ''; via = '' }
    }
    $candidateByPath = [System.Collections.Generic.Dictionary[string, object]]::new([System.StringComparer]::Ordinal)
    $referencerDirectory = [System.IO.Path]::GetDirectoryName($Referencer).Replace('\', '/')
    if ([string]$Directive.style -ceq 'quote') {
        $candidate = ConvertTo-CanonicalModulePath $referencerDirectory ([string]$Directive.token)
        $root = [System.IO.Path]::GetDirectoryName($candidate).Replace('\', '/')
        if ([string]::IsNullOrWhiteSpace($root)) { $root = '.' }
        $candidateByPath[$candidate] = [pscustomobject]@{ path = $candidate; root = $root; via = 'source_directory' }
    }
    foreach ($includeRoot in @($Controls.includeRoots | Sort-Object -CaseSensitive)) {
        $candidate = ConvertTo-CanonicalModulePath ([string]$includeRoot) ([string]$Directive.token)
        if (-not $candidateByPath.ContainsKey($candidate)) {
            $candidateByPath[$candidate] = [pscustomobject]@{ path = $candidate; root = [string]$includeRoot; via = 'include_root' }
        }
    }
    $matches = [System.Collections.Generic.List[object]]::new()
    foreach ($candidate in @($candidateByPath.Values)) {
        if (-not $ZipIndex.entries.ContainsKey($candidate.path) -and $ZipIndex.entriesIgnoreCase.ContainsKey($candidate.path)) {
            throw "Case-mismatched module include: requested=$($candidate.path) actual=$($ZipIndex.entriesIgnoreCase[$candidate.path])"
        }
        $vendorPath = Join-Path (Join-Path $VendorRoot ([string]$Module.path).Replace('/', '\')) $candidate.path.Replace('/', '\')
        $sourceExists = $ZipIndex.entries.ContainsKey($candidate.path)
        $vendorExists = Test-Path -LiteralPath $vendorPath -PathType Leaf
        if ($sourceExists -or $vendorExists) {
            $matches.Add([pscustomobject]@{
                path = $candidate.path; root = $candidate.root; via = $candidate.via
                sourceExists = $sourceExists; vendorExists = $vendorExists
            })
        }
    }
    $distinctMatches = @($matches | Sort-Object path -Unique -CaseSensitive)
    if ($distinctMatches.Count -gt 1) {
        throw "Ambiguous module-local include $($Directive.token) from $($Module.path)/${Referencer}: $($distinctMatches.path -join ', ')"
    }
    if ($distinctMatches.Count -eq 0) {
        if ([string]$Directive.style -ceq 'angle') { return [pscustomobject]@{ kind = 'system'; path = ''; root = ''; via = '' } }
        $requested = ConvertTo-CanonicalModulePath $referencerDirectory ([string]$Directive.token)
        $isPatchScanner = [string]$Module.path -ceq $PatchedModulePath -and [string]$Module.version -ceq $PatchedModuleVersion -and
            $Referencer -ceq 'bindings/go/binding.go' -and $requested -ceq 'src/scanner.c' -and
            [string]$Directive.condition -ceq '__has_include("../../src/scanner.c")'
        if ($isPatchScanner) { return [pscustomobject]@{ kind = 'patch_exception'; path = $requested; root = 'src'; via = 'source_directory' } }
        if ([string]$Directive.state -ceq 'inactive') { return [pscustomobject]@{ kind = 'inactive_missing'; path = $requested; root = ''; via = 'source_directory' } }
        if ([string]$Directive.state -ceq 'unknown') { throw "Unsupported conditional guards a module-local include: $($Module.path)/${Referencer}:$($Directive.line)" }
        throw "Active module-local include is absent from authenticated source: $($Module.path)@$($Module.version)/$requested"
    }
    $match = $distinctMatches[0]
    $effectiveState = [string]$Directive.state
    if ($effectiveState -ceq 'unknown' -and [string]$Directive.condition -cmatch '^__has_include\("([^"]+)"\)$' -and $Matches[1] -ceq [string]$Directive.token) {
        $effectiveState = 'active'
    }
    if ($effectiveState -ceq 'inactive') { return [pscustomobject]@{ kind = 'inactive'; path = $match.path; root = $match.root; via = $match.via } }
    if ($effectiveState -ceq 'unknown') { throw "Unsupported conditional guards module-local include $($match.path) in $($Module.path)/$Referencer" }
    if ($match.vendorExists) { return [pscustomobject]@{ kind = 'present'; path = $match.path; root = $match.root; via = $match.via } }
    if (-not $match.sourceExists) { throw "Resolved module-local include has no authenticated zip entry: $($match.path)" }
    return [pscustomobject]@{ kind = 'supplement'; path = $match.path; root = $match.root; via = $match.via }
}

function Copy-SupplementalRoot {
    param(
        [Parameter(Mandatory)][object]$Module,
        [Parameter(Mandatory)][object]$ZipIndex,
        [Parameter(Mandatory)][string]$Root,
        [Parameter(Mandatory)][string]$VendorRoot,
        [Parameter(Mandatory)][System.Collections.Generic.Dictionary[string, object]]$SupplementalFiles
    )
    $sourceRoot = if ($Root -ceq '.') { [string]$Module.dir } else { Join-Path ([string]$Module.dir) $Root.Replace('/', '\') }
    if (-not (Test-Path -LiteralPath $sourceRoot -PathType Container)) { throw "Supplemental source root is absent: $($Module.path)@$($Module.version)/$Root" }
    Assert-NoReparsePoints $sourceRoot "Supplemental source root $($Module.path)@$($Module.version)/$Root"
    $rootRows = [System.Collections.Generic.List[object]]::new()
    foreach ($sourceFile in @(Get-ChildItem -LiteralPath $sourceRoot -File -Force -Recurse | Sort-Object FullName)) {
        $moduleRelative = (Get-StrictRelativePath -BasePath ([string]$Module.dir) -ChildPath $sourceFile.FullName).Replace('\', '/')
        Assert-CanonicalRelativePath $moduleRelative
        if (-not $ZipIndex.entries.ContainsKey($moduleRelative)) { throw "Supplemental source file is absent from authenticated zip: $($Module.path)/$moduleRelative" }
        $zipBytes = Get-ZipEntryBytes $ZipIndex $moduleRelative
        $sourceBytes = [System.IO.File]::ReadAllBytes($sourceFile.FullName)
        if ($zipBytes.Length -ne $sourceBytes.Length -or (Get-BytesSha256Hex $zipBytes) -cne (Get-BytesSha256Hex $sourceBytes)) {
            throw "Extracted source differs from authenticated zip entry: $($Module.path)/$moduleRelative"
        }
        $vendorRelative = (([string]$Module.path).Replace('\', '/') + '/' + $moduleRelative)
        Assert-CanonicalRelativePath $vendorRelative
        $destination = Join-Path $VendorRoot $vendorRelative.Replace('/', '\')
        if (Test-Path -LiteralPath $destination -PathType Leaf) {
            $existingBytes = [System.IO.File]::ReadAllBytes($destination)
            if ($existingBytes.Length -ne $zipBytes.Length -or (Get-BytesSha256Hex $existingBytes) -cne (Get-BytesSha256Hex $zipBytes)) {
                throw "Supplemental root overlaps baseline with different bytes: $vendorRelative"
            }
            continue
        }
        [System.IO.Directory]::CreateDirectory((Split-Path -Parent $destination)) | Out-Null
        [System.IO.File]::WriteAllBytes($destination, $zipBytes)
        $row = [pscustomobject][ordered]@{
            path = $vendorRelative
            modulePath = [string]$Module.path
            moduleVersion = [string]$Module.version
            moduleRelativePath = $moduleRelative
            root = $Root
            bytes = [long]$zipBytes.Length
            sha256 = Get-BytesSha256Hex $zipBytes
        }
        if ($SupplementalFiles.ContainsKey($vendorRelative)) { throw "Duplicate supplemental destination: $vendorRelative" }
        $SupplementalFiles.Add($vendorRelative, $row)
        $rootRows.Add($row)
    }
    return @($rootRows)
}

function Invoke-SupplementalClosure {
    param([Parameter(Mandatory)][string]$VendorRoot, [Parameter(Mandatory)][object[]]$Modules)
    $model = Get-CgoClosureModels $VendorRoot $Modules
    $zipIndexes = @{}
    try {
        foreach ($module in $Modules) {
            $moduleID = "$($module.path)@$($module.version)"
            $zipIndexes[$moduleID] = New-ModuleZipIndex $module
        }
        $rootQueue = [System.Collections.Generic.List[object]]::new()
        $rootKeys = [System.Collections.Generic.HashSet[string]]::new([System.StringComparer]::Ordinal)
        $supplementalFiles = [System.Collections.Generic.Dictionary[string, object]]::new([System.StringComparer]::Ordinal)
        $seedRows = [System.Collections.Generic.List[object]]::new()
        $edgeRows = [System.Collections.Generic.List[object]]::new()
        $inactiveRows = [System.Collections.Generic.List[object]]::new()
        $seedKeys = [System.Collections.Generic.HashSet[string]]::new([System.StringComparer]::Ordinal)
        $edgeKeys = [System.Collections.Generic.HashSet[string]]::new([System.StringComparer]::Ordinal)
        $inactiveKeys = [System.Collections.Generic.HashSet[string]]::new([System.StringComparer]::Ordinal)

        foreach ($preamble in @($model.models | Sort-Object moduleID, moduleRelativeFile -CaseSensitive)) {
            $controls = $preamble.controls
            if (@($preamble.preambleLines | Where-Object { -not [string]::IsNullOrWhiteSpace([string]$_) }).Count -eq 0) { continue }
            $directives = @(Get-IncludeDirectives $preamble.preambleLines $controls.defines "$($preamble.module.path)/$($preamble.moduleRelativeFile)")
            foreach ($directive in $directives) {
                $resolved = Resolve-ModuleInclude $preamble.module $zipIndexes[$preamble.moduleID] $controls $VendorRoot $preamble.moduleRelativeFile $directive
                if ($resolved.kind -ceq 'system' -or $resolved.kind -ceq 'inactive_dynamic') { continue }
                $seedKey = "$($preamble.moduleID)`t$($preamble.moduleRelativeFile)`t$($resolved.path)`t$($resolved.kind)"
                if ($seedKeys.Add($seedKey)) {
                    $seedRows.Add([pscustomobject][ordered]@{
                        modulePath = [string]$preamble.module.path; moduleVersion = [string]$preamble.module.version
                        referencer = [string]$preamble.moduleRelativeFile; line = [int]$directive.line
                        include = [string]$directive.token; style = [string]$directive.style
                        resolvedPath = [string]$resolved.path; resolution = [string]$resolved.via
                        disposition = [string]$resolved.kind
                    })
                }
                if ($resolved.kind -ceq 'inactive' -or $resolved.kind -ceq 'inactive_missing') {
                    $inactiveKey = "$($preamble.moduleID)`t$($preamble.moduleRelativeFile)`t$($resolved.path)`t$($directive.condition)"
                    if ($inactiveKeys.Add($inactiveKey)) {
                        $inactiveRows.Add([pscustomobject][ordered]@{
                            modulePath = [string]$preamble.module.path; moduleVersion = [string]$preamble.module.version
                            referencer = [string]$preamble.moduleRelativeFile; resolvedPath = [string]$resolved.path
                            condition = [string]$directive.condition; state = 'inactive'
                        })
                    }
                }
                elseif ($resolved.kind -ceq 'supplement') {
                    $rootKey = "$($preamble.moduleID)`t$($resolved.root)"
                    if ($rootKeys.Add($rootKey)) {
                        $rootQueue.Add([pscustomobject]@{ module = $preamble.module; moduleID = $preamble.moduleID; root = [string]$resolved.root; controls = $controls })
                    }
                }
            }
        }

        $queueIndex = 0
        while ($queueIndex -lt $rootQueue.Count) {
            $rootRow = $rootQueue[$queueIndex]
            $queueIndex++
            $newFiles = @(Copy-SupplementalRoot $rootRow.module $zipIndexes[$rootRow.moduleID] $rootRow.root $VendorRoot $supplementalFiles)
            $controls = $rootRow.controls
            if ($null -eq $controls) {
                $defines = [System.Collections.Generic.HashSet[string]]::new([System.StringComparer]::Ordinal)
                foreach ($define in @('_WIN32', 'WIN32', '__x86_64__', '_M_X64')) { [void]$defines.Add($define) }
                $controls = [pscustomobject]@{
                    includeRoots = [System.Collections.Generic.HashSet[string]]::new([System.StringComparer]::Ordinal)
                    defines = $defines
                }
                $model.controlsByModule[$rootRow.moduleID] = $controls
            }
            foreach ($fileRow in $newFiles) {
                $extension = [System.IO.Path]::GetExtension([string]$fileRow.moduleRelativePath).ToLowerInvariant()
                if ($extension -notin @('.c', '.cc', '.cpp', '.cxx', '.h', '.hh', '.hpp', '.hxx', '.inc', '.inl')) { continue }
                $filePath = Join-Path (Join-Path $VendorRoot ([string]$rootRow.module.path).Replace('/', '\')) ([string]$fileRow.moduleRelativePath).Replace('/', '\')
                $lines = @((Read-StrictUtf8Text $filePath).Replace("`r`n", "`n").Replace("`r", "`n") -split "`n")
                $directives = @(Get-IncludeDirectives $lines $controls.defines "$($rootRow.module.path)/$($fileRow.moduleRelativePath)")
                foreach ($directive in $directives) {
                    $resolved = Resolve-ModuleInclude $rootRow.module $zipIndexes[$rootRow.moduleID] $controls $VendorRoot ([string]$fileRow.moduleRelativePath) $directive
                    if ($resolved.kind -ceq 'system' -or $resolved.kind -ceq 'inactive_dynamic') { continue }
                    if ($resolved.kind -ceq 'patch_exception') { throw "Reviewed patch exception is only valid in the exact Go binding preamble." }
                    if ($resolved.kind -ceq 'inactive' -or $resolved.kind -ceq 'inactive_missing') {
                        $inactiveKey = "$($rootRow.moduleID)`t$($fileRow.moduleRelativePath)`t$($resolved.path)`t$($directive.condition)"
                        if ($inactiveKeys.Add($inactiveKey)) {
                            $inactiveRows.Add([pscustomobject][ordered]@{
                                modulePath = [string]$rootRow.module.path; moduleVersion = [string]$rootRow.module.version
                                referencer = [string]$fileRow.moduleRelativePath; resolvedPath = [string]$resolved.path
                                condition = [string]$directive.condition; state = 'inactive'
                            })
                        }
                        continue
                    }
                    if ($resolved.kind -ceq 'supplement') {
                        $nextRootKey = "$($rootRow.moduleID)`t$($resolved.root)"
                        if ($rootKeys.Add($nextRootKey)) {
                            $rootQueue.Add([pscustomobject]@{ module = $rootRow.module; moduleID = $rootRow.moduleID; root = [string]$resolved.root; controls = $controls })
                        }
                    }
                    $edgeKey = "$($rootRow.moduleID)`t$($fileRow.moduleRelativePath)`t$($resolved.path)"
                    if ($edgeKeys.Add($edgeKey)) {
                        $edgeRows.Add([pscustomobject][ordered]@{
                            modulePath = [string]$rootRow.module.path; moduleVersion = [string]$rootRow.module.version
                            referencer = [string]$fileRow.moduleRelativePath; line = [int]$directive.line
                            include = [string]$directive.token; resolvedPath = [string]$resolved.path
                            resolution = [string]$resolved.via; disposition = [string]$resolved.kind
                        })
                    }
                }
            }
        }
        Assert-NoReparsePoints $VendorRoot 'Supplemented vendor tree'
        $fileRows = @(Sort-InventoryRows @($supplementalFiles.Values))
        $rootManifestRows = [System.Collections.Generic.List[object]]::new()
        foreach ($queuedRoot in @($rootQueue | Sort-Object { "$($_.moduleID)`t$($_.root)" } -CaseSensitive)) {
            $rowsForRoot = @($fileRows | Where-Object {
                $_.modulePath -ceq [string]$queuedRoot.module.path -and $_.moduleVersion -ceq [string]$queuedRoot.module.version -and $_.root -ceq [string]$queuedRoot.root
            })
            $rootManifestRows.Add([pscustomobject][ordered]@{
                modulePath = [string]$queuedRoot.module.path; moduleVersion = [string]$queuedRoot.module.version
                root = [string]$queuedRoot.root; files = $rowsForRoot.Count
                bytes = [long](($rowsForRoot | Measure-Object bytes -Sum).Sum)
                treeSha256 = Get-TreeDigest $rowsForRoot
            })
        }
        $affectedModules = @($rootManifestRows | ForEach-Object { "$($_.modulePath)@$($_.moduleVersion)" } | Sort-Object -Unique -CaseSensitive)
        $effectiveDefines = @($model.controlsByModule.Keys | Sort-Object -CaseSensitive | ForEach-Object {
            $moduleID = $_
            [pscustomobject][ordered]@{
                module = $moduleID
                defines = @($model.controlsByModule[$moduleID].defines | Sort-Object -CaseSensitive)
                includeRoots = @($model.controlsByModule[$moduleID].includeRoots | Sort-Object -CaseSensitive)
            }
        })
        $wasmRows = @($inactiveRows | Where-Object {
            $_.modulePath -ceq 'github.com/tree-sitter/go-tree-sitter' -and $_.resolvedPath -ceq 'src/stdlib-symbols.txt'
        })
        if ($wasmRows.Count -ne 1) { throw "Expected exactly one inactive TREE_SITTER_FEATURE_WASM reference; found $($wasmRows.Count)" }
        return [pscustomobject]@{
            affectedModules = $affectedModules
            roots = @($rootManifestRows)
            activeSeedTargets = @($seedRows | Sort-Object modulePath, moduleVersion, referencer, resolvedPath -CaseSensitive)
            recursiveEdges = @($edgeRows | Sort-Object modulePath, moduleVersion, referencer, resolvedPath -CaseSensitive)
            inactiveConditionalReferences = @($inactiveRows | Sort-Object modulePath, moduleVersion, referencer, resolvedPath -CaseSensitive)
            effectiveDefines = $effectiveDefines
            files = $fileRows
            treeSha256 = Get-TreeDigest $fileRows
        }
    }
    finally {
        foreach ($zipIndex in @($zipIndexes.Values)) { if ($null -ne $zipIndex) { $zipIndex.archive.Dispose() } }
    }
}

function Assert-AndApplyReviewedPatch {
    param([Parameter(Mandatory)][string]$VendorRoot, [Parameter(Mandatory)][object[]]$Modules)
    $patchPath = Join-Path $script:RepoRoot $PatchAuthorityRelativePath.Replace('/', '\')
    if (-not (Test-Path -LiteralPath $patchPath -PathType Leaf)) { throw "Reviewed patch authority is absent: $PatchAuthorityRelativePath" }
    $patchFiles = @(Get-ChildItem -LiteralPath (Split-Path -Parent $patchPath) -File -Force -Recurse)
    if ($patchFiles.Count -ne 1 -or (Get-Sha256 $patchPath) -cne $ExpectedPatchSha256) {
        throw "Reviewed patch cardinality or identity is invalid: count=$($patchFiles.Count)"
    }
    $module = @($Modules | Where-Object { $_.path -ceq $PatchedModulePath -and $_.version -ceq $PatchedModuleVersion })
    if ($module.Count -ne 1) { throw "Patched module denominator is invalid: $($module.Count)" }
    $module = $module[0]
    $bindingPath = Join-Path $VendorRoot $PatchedBindingRelativePath.Replace('/', '\')
    $preimage = [System.IO.File]::ReadAllBytes($bindingPath)
    if ($preimage.Length -ne $ExpectedPatchPreimageBytes -or (Get-BytesSha256Hex $preimage) -cne $ExpectedPatchPreimageSha256) {
        throw "Tree-sitter Go binding preimage drifted."
    }
    $grammarPath = Join-Path ([string]$module.dir) 'src\grammar.json'
    $parserPath = Join-Path ([string]$module.dir) 'src\parser.c'
    if ((Get-Item -LiteralPath $grammarPath).Length -ne $ExpectedGrammarBytes -or (Get-Sha256 $grammarPath) -cne $ExpectedGrammarSha256) { throw 'tree-sitter-go grammar identity drifted.' }
    if ((Get-Item -LiteralPath $parserPath).Length -ne $ExpectedParserBytes -or (Get-Sha256 $parserPath) -cne $ExpectedParserSha256) { throw 'tree-sitter-go parser identity drifted.' }
    $grammar = Read-StrictUtf8Text $grammarPath | ConvertFrom-Json
    if (@($grammar.externals).Count -ne 0) { throw 'tree-sitter-go grammar externals must remain empty.' }
    $parserText = Read-StrictUtf8Text $parserPath
    if (@([regex]::Matches($parserText, '(?m)^#define EXTERNAL_TOKEN_COUNT 0$')).Count -ne 1) { throw 'tree-sitter-go EXTERNAL_TOKEN_COUNT guard failed.' }
    foreach ($scanner in @('src\scanner.c', 'src\scanner.cc', 'src\scanner.cpp')) {
        if (Test-Path -LiteralPath (Join-Path ([string]$module.dir) $scanner)) { throw "Unexpected tree-sitter-go scanner source: $scanner" }
    }
    $zipIndex = New-ModuleZipIndex $module
    try {
        foreach ($scanner in @('src/scanner.c', 'src/scanner.cc', 'src/scanner.cpp')) {
            if ($zipIndex.entries.ContainsKey($scanner)) { throw "Unexpected tree-sitter-go scanner zip entry: $scanner" }
        }
    }
    finally { $zipIndex.archive.Dispose() }
    $removedBlock = "// #if __has_include(`"../../src/scanner.c`")`n// #include `"../../src/scanner.c`"`n// #endif`n"
    $preimageText = $Utf8NoBomStrict.GetString($preimage)
    $occurrences = @([regex]::Matches($preimageText, [regex]::Escape($removedBlock))).Count
    if ($occurrences -ne 1) { throw "Reviewed patch block occurrence mismatch: $occurrences" }
    $postimageText = $preimageText.Replace($removedBlock, '')
    $postimage = $Utf8NoBom.GetBytes($postimageText)
    if ($postimage.Length -ne $ExpectedPatchPostimageBytes -or (Get-BytesSha256Hex $postimage) -cne $ExpectedPatchPostimageSha256) {
        throw 'Reviewed patch postimage guard failed.'
    }
    [System.IO.File]::WriteAllBytes($bindingPath, $postimage)
    return [pscustomobject][ordered]@{
        authorityPath = $PatchAuthorityRelativePath
        patchSha256 = $ExpectedPatchSha256
        modulePath = $PatchedModulePath
        moduleVersion = $PatchedModuleVersion
        destination = 'vendor/' + $PatchedBindingRelativePath
        preimageBytes = $ExpectedPatchPreimageBytes
        preimageSha256 = $ExpectedPatchPreimageSha256
        removedBlock = $removedBlock
        occurrenceCount = 1
        grammarBytes = $ExpectedGrammarBytes
        grammarSha256 = $ExpectedGrammarSha256
        grammarExternals = 0
        parserBytes = $ExpectedParserBytes
        parserSha256 = $ExpectedParserSha256
        externalTokenCount = 0
        sourceScannerEntries = 0
        zipScannerEntries = 0
        postimageBytes = $ExpectedPatchPostimageBytes
        postimageSha256 = $ExpectedPatchPostimageSha256
        additions = 0
        deletions = 3
        byteDelta = -88
        fuzz = 0
        offset = 0
    }
}

function New-CompleteCandidate {
    param(
        [Parameter(Mandatory)][string]$CandidateRoot,
        [Parameter(Mandatory)][string]$VendorSource,
        [Parameter(Mandatory)][object[]]$SourceModules
    )
    [System.IO.Directory]::CreateDirectory($CandidateRoot) | Out-Null
    $candidateVendor = Join-Path $CandidateRoot 'vendor'
    $candidateAuthority = Join-Path $CandidateRoot 'third_party\go-vendor'
    $candidateLicenses = Join-Path $candidateAuthority 'licenses'
    $candidatePatches = Join-Path $candidateAuthority 'patches'
    [System.IO.Directory]::CreateDirectory((Split-Path -Parent $candidateVendor)) | Out-Null
    [System.IO.Directory]::CreateDirectory($candidateLicenses) | Out-Null
    [System.IO.Directory]::CreateDirectory($candidatePatches) | Out-Null
    Copy-Item -LiteralPath $VendorSource -Destination $candidateVendor -Recurse
    $patchSource = Join-Path $script:RepoRoot $PatchAuthorityRelativePath.Replace('/', '\')
    $patchDestination = Join-Path $CandidateRoot $PatchAuthorityRelativePath.Replace('/', '\')
    [System.IO.File]::WriteAllBytes($patchDestination, [System.IO.File]::ReadAllBytes($patchSource))
    $modulesText = Read-StrictUtf8Text (Join-Path $candidateVendor 'modules.txt')
    $vendorModules = @(Get-VendorModules $modulesText)
    $sourceModuleById = @{}
    foreach ($module in $SourceModules) { $sourceModuleById["$($module.path)@$($module.version)"] = $module }
    $manifestModuleRows = [System.Collections.Generic.List[object]]::new()
    $copiedDispositionCount = 0
    $exceptionDispositionCount = 0
    foreach ($vendorModule in $vendorModules) {
        $identity = "$($vendorModule.path)@$($vendorModule.version)"
        $sourceModule = $sourceModuleById[$identity]
        if ($null -eq $sourceModule) { throw "Vendored module is absent from verified source closure: $identity" }
        $moduleDirectory = [System.IO.Path]::GetFullPath([string]$sourceModule.dir)
        $escapedIdentity = Get-StrictRelativePath -BasePath (Join-Path $acquireState 'gomodcache') -ChildPath $moduleDirectory
        $licenseFiles = @(Get-ChildItem -LiteralPath $moduleDirectory -File -Force -Recurse | Where-Object {
            $_.Name -match '^(?i:licen[cs]e|notice|copying|copyright|patents|unlicense)(?:[._-].*)?$'
        } | Sort-Object FullName)
        $licensePaths = [System.Collections.Generic.List[string]]::new()
        $disposition = 'reviewed_exception'
        $review = 'No LICENSE, LICENCE, NOTICE, COPYING, COPYRIGHT, PATENTS, or UNLICENSE file exists in the checksum-verified module archive.'
        if ($licenseFiles.Count -gt 0) {
            $disposition = 'copied'; $review = ''; $copiedDispositionCount++
            foreach ($licenseFile in $licenseFiles) {
                $moduleRelative = Get-StrictRelativePath -BasePath $moduleDirectory -ChildPath $licenseFile.FullName
                $relativeDestination = (Join-Path (Join-Path 'licenses' $escapedIdentity) $moduleRelative).Replace('\', '/')
                Assert-CanonicalRelativePath $relativeDestination
                $destination = Join-Path $candidateAuthority $relativeDestination.Replace('/', '\')
                [System.IO.Directory]::CreateDirectory((Split-Path -Parent $destination)) | Out-Null
                [System.IO.File]::WriteAllBytes($destination, [System.IO.File]::ReadAllBytes($licenseFile.FullName))
                $licensePaths.Add($relativeDestination)
            }
        }
        else { $exceptionDispositionCount++ }
        $manifestModuleRows.Add([pscustomobject][ordered]@{
            path = $vendorModule.path; version = $vendorModule.version
            contentSum = $sourceModule.contentSum; goModSum = $sourceModule.goModSum
            licenseDisposition = $disposition
            licensePaths = @($licensePaths | Sort-Object -CaseSensitive)
            licenseReview = $review
        })
    }
    if (($copiedDispositionCount + $exceptionDispositionCount) -ne $vendorModules.Count) { throw 'License disposition denominator is incomplete.' }
    Assert-NoReparsePoints $candidateVendor 'Candidate vendor'
    Assert-NoReparsePoints $candidateAuthority 'Candidate Go vendor authority'
    return [pscustomobject]@{
        root = $CandidateRoot
        vendor = $candidateVendor
        authority = $candidateAuthority
        licenses = $candidateLicenses
        modulesText = $modulesText
        vendorModules = $vendorModules
        manifestModuleRows = @(Sort-ModuleRowsOrdinal @($manifestModuleRows))
        copiedDispositionCount = $copiedDispositionCount
        exceptionDispositionCount = $exceptionDispositionCount
    }
}

function Publish-AtomicVendorAuthority {
    param([Parameter(Mandatory)][string]$CandidateRoot)
    $candidateVendor = Join-Path $CandidateRoot 'vendor'
    $candidateAuthority = Join-Path $CandidateRoot 'third_party\go-vendor'
    $durableVendor = Join-Path $script:RepoRoot 'vendor'
    $durableAuthority = Join-Path $script:RepoRoot 'third_party\go-vendor'
    $backupRoot = Join-Path $script:Lane 'publish-backup'
    $backupVendor = Join-Path $backupRoot 'vendor'
    $backupAuthority = Join-Path $backupRoot 'go-vendor'
    [System.IO.Directory]::CreateDirectory((Split-Path -Parent $durableAuthority)) | Out-Null
    [System.IO.Directory]::CreateDirectory($backupRoot) | Out-Null
    $hadVendor = Test-Path -LiteralPath $durableVendor
    $hadAuthority = Test-Path -LiteralPath $durableAuthority
    $publishedVendor = $false
    $publishedAuthority = $false
    try {
        if ($hadVendor) { [System.IO.Directory]::Move($durableVendor, $backupVendor) }
        if ($hadAuthority) { [System.IO.Directory]::Move($durableAuthority, $backupAuthority) }
        [System.IO.Directory]::Move($candidateVendor, $durableVendor)
        $publishedVendor = $true
        [System.IO.Directory]::Move($candidateAuthority, $durableAuthority)
        $publishedAuthority = $true
        & (Join-Path $PSScriptRoot 'verify-go-vendor.ps1') -SourceRoot $script:RepoRoot -Json | Out-Null
        if ($LASTEXITCODE -ne 0) { throw "Durable vendor verification exited $LASTEXITCODE" }
    } catch {
        if ($publishedAuthority -and (Test-Path -LiteralPath $durableAuthority)) {
            Remove-Item -LiteralPath $durableAuthority -Recurse -Force
        }
        if ($publishedVendor -and (Test-Path -LiteralPath $durableVendor)) {
            Remove-Item -LiteralPath $durableVendor -Recurse -Force
        }
        if ($hadVendor -and (Test-Path -LiteralPath $backupVendor)) {
            [System.IO.Directory]::Move($backupVendor, $durableVendor)
        }
        if ($hadAuthority -and (Test-Path -LiteralPath $backupAuthority)) {
            [System.IO.Directory]::Move($backupAuthority, $durableAuthority)
        }
        throw
    }
}

$script:RepoRoot = [System.IO.Path]::GetFullPath($RepositoryRoot).TrimEnd('\')
$script:Lane = [System.IO.Path]::GetFullPath($LaneRoot).TrimEnd('\')
$repoTmpRoot = [System.IO.Path]::GetFullPath((Join-Path $script:RepoRoot '.tmp')).TrimEnd('\')
if (-not $script:Lane.StartsWith($repoTmpRoot + '\', [System.StringComparison]::OrdinalIgnoreCase)) {
    throw "LaneRoot must be a strict child of $repoTmpRoot"
}
if ($SourceProxy -cne $AllowedProxy) {
    throw "Only exact source proxy $AllowedProxy is authorized."
}
if ($SourceProxy.Contains(',') -or $SourceProxy.Contains('|')) {
    throw 'Proxy fallback syntax is forbidden.'
}
if (-not (Test-Path -LiteralPath $script:RepoRoot -PathType Container)) { throw "Repository root is absent: $script:RepoRoot" }
if (Test-Path -LiteralPath $script:Lane) {
    $existing = @(Get-ChildItem -LiteralPath $script:Lane -Force)
    if ($existing.Count -ne 0) { throw "LaneRoot must be new or empty: $script:Lane" }
} else {
    [System.IO.Directory]::CreateDirectory($script:Lane) | Out-Null
}
Assert-NoReparsePoints $script:Lane 'Lane root'
Assert-ModuleIdentity $script:RepoRoot

$acquisitionProxy = $AllowedProxy
$acquisitionMode = 'exact_proxy'
if (-not [string]::IsNullOrWhiteSpace($InputProxy)) {
    $resolvedInputProxy = [System.IO.Path]::GetFullPath($InputProxy).TrimEnd('\')
    if (-not $resolvedInputProxy.StartsWith($repoTmpRoot + '\', [System.StringComparison]::OrdinalIgnoreCase)) {
        throw "InputProxy must be a strict child of $repoTmpRoot"
    }
    if (-not (Test-Path -LiteralPath $resolvedInputProxy -PathType Container)) { throw "InputProxy is absent: $resolvedInputProxy" }
    Assert-NoReparsePoints $resolvedInputProxy 'Checksum-verified input proxy'
    $acquisitionProxy = Convert-ToFileProxyUri $resolvedInputProxy
    $acquisitionMode = 'checksum_verified_file_proxy'
}

$script:LogsRoot = Join-Path $script:Lane 'logs'
$evidenceRoot = Join-Path $script:Lane 'evidence'
$acquireRoot = Join-Path $script:Lane 'acquire'
$inputProxy = Join-Path $script:Lane 'input-proxy'
$offlineRoot = Join-Path $script:Lane 'offline-rehydrate'
$vendorRun1Root = Join-Path $script:Lane 'vendor-run-1'
$vendorRun2Root = Join-Path $script:Lane 'vendor-run-2'
$candidateRoot = Join-Path $script:Lane 'publish-candidate'
foreach ($path in @($script:LogsRoot, $evidenceRoot, $acquireRoot)) {
    [System.IO.Directory]::CreateDirectory($path) | Out-Null
}

$acquireModule = Join-Path $acquireRoot 'module'
$acquireState = Join-Path $acquireRoot 'state'
Copy-ModuleAuthority $acquireModule
Set-GoEnvironment $acquireState $acquisitionProxy
$goVersion = Invoke-GoLogged '001-go-version' $acquireModule @('version')
if ($goVersion.Stdout.Trim() -cne $ExpectedGoVersion) { throw "Unexpected Go toolchain: $($goVersion.Stdout.Trim())" }
Assert-ModuleIdentity $script:RepoRoot
$initialGoModEditResult = Invoke-GoLogged '002-go-mod-edit-json' $acquireModule @('mod', 'edit', '-json')
$goModModel = $initialGoModEditResult.Stdout | ConvertFrom-Json
$vendorDiscoveryModules = @($goModModel.Require | ForEach-Object {
    [pscustomobject][ordered]@{ path = [string]$_.Path; version = [string]$_.Version }
} | Sort-Object path, version -CaseSensitive)
$downloadArguments = if ($acquisitionMode -ceq 'checksum_verified_file_proxy') {
    @('mod', 'download', '-json') + @($vendorDiscoveryModules | ForEach-Object { "$($_.path)@$($_.version)" })
} else {
    @('mod', 'download', '-json', 'all')
}
$downloadResult = Invoke-GoLogged '003-go-mod-download-source-closure' $acquireModule $downloadArguments
$listResult = if ($acquisitionMode -ceq 'checksum_verified_file_proxy') { $null } else { Invoke-GoLogged '004-go-list-modules' $acquireModule @('list', '-mod=readonly', '-m', '-json', 'all') }
$verifyResult = Invoke-GoLogged '005-go-mod-verify-source-closure' $acquireModule @('mod', 'verify')
if ($verifyResult.Stdout.Trim() -cne 'all modules verified') { throw "Unexpected Go verification output: $($verifyResult.Stdout.Trim())" }
Assert-ModuleIdentity $script:RepoRoot

$downloads = @(ConvertFrom-JsonSequence $downloadResult.Stdout)
$buildList = if ($acquisitionMode -ceq 'checksum_verified_file_proxy') {
    @([pscustomobject]@{ Path = [string]$goModModel.Module.Path; Main = $true }) + @($vendorDiscoveryModules | ForEach-Object { [pscustomobject]@{ Path = $_.path; Version = $_.version } })
} else {
    @(ConvertFrom-JsonSequence $listResult.Stdout)
}
$dependencies = @($buildList | Where-Object { -not ($_.PSObject.Properties.Name -contains 'Main') } | Sort-Object Path, Version)
if ($dependencies.Count -eq 0 -or ($acquisitionMode -ceq 'checksum_verified_file_proxy' -and $downloads.Count -ne $vendorDiscoveryModules.Count) -or
    ($acquisitionMode -cne 'checksum_verified_file_proxy' -and $downloads.Count -ne $dependencies.Count)) {
    throw "Unexpected module denominator: buildList=$($buildList.Count), dependencies=$($dependencies.Count), downloads=$($downloads.Count)"
}
$goSumEntries = @{}
$authorizedGoModPairs = [System.Collections.Generic.List[object]]::new()
foreach ($line in [System.IO.File]::ReadAllLines((Join-Path $script:RepoRoot 'go.sum'))) {
    if ([string]::IsNullOrWhiteSpace($line)) { continue }
    $parts = $line -split ' '
    if ($parts.Count -ne 3) { throw "Malformed go.sum line: $line" }
    $goSumEntries["$($parts[0]) $($parts[1])"] = $parts[2]
    if ($parts[1].EndsWith('/go.mod', [System.StringComparison]::Ordinal)) {
        $authorizedGoModPairs.Add([pscustomobject][ordered]@{
            path = $parts[0]
            version = $parts[1].Substring(0, $parts[1].Length - 7)
            goModSum = $parts[2]
        })
    }
}
$downloadById = @{}
foreach ($download in $downloads) { $downloadById["$($download.Path)@$($download.Version)"] = $download }
$authorizedGoModPairs = @($authorizedGoModPairs | Sort-Object path, version -CaseSensitive)
$authorizedGoModQueryArguments = @('list', '-mod=mod', '-m', '-json') + @($authorizedGoModPairs | ForEach-Object { "$($_.path)@$($_.version)" })
$authorizedGoModDownloadResult = Invoke-GoLogged '006-go-list-authorized-graph-metadata' $acquireModule $authorizedGoModQueryArguments
$authorizedGoModDownloads = @(ConvertFrom-JsonSequence $authorizedGoModDownloadResult.Stdout | ForEach-Object {
    $infoPath = if ([string]::IsNullOrWhiteSpace([string]$_.GoMod)) { '' } else { ([string]$_.GoMod) -replace '\.mod$', '.info' }
    [pscustomobject]@{
        Path = [string]$_.Path; Version = [string]$_.Version; Info = $infoPath
        GoMod = [string]$_.GoMod; GoModSum = [string]$_.GoModSum
    }
})
if ($authorizedGoModDownloads.Count -ne $authorizedGoModPairs.Count) {
    throw "Authorized go.mod download denominator mismatch: expected=$($authorizedGoModPairs.Count), actual=$($authorizedGoModDownloads.Count)"
}
$authorizedGoModDownloadById = @{}
foreach ($download in $authorizedGoModDownloads) {
    $identity = "$($download.Path)@$($download.Version)"
    $authorizedGoModDownloadById[$identity] = $download
}
$graphModules = [System.Collections.Generic.List[object]]::new()
$authorizedGoModMismatch = 0
$authorizedGoModMissing = 0
foreach ($pair in $authorizedGoModPairs) {
    $identity = "$($pair.path)@$($pair.version)"
    $download = $authorizedGoModDownloadById[$identity]
    if ($null -eq $download -or
        -not (Test-Path -LiteralPath ([string]$download.Info) -PathType Leaf) -or
        -not (Test-Path -LiteralPath ([string]$download.GoMod) -PathType Leaf)) {
        $authorizedGoModMissing++
        continue
    }
    if ($download.GoModSum -cne $pair.goModSum) { $authorizedGoModMismatch++ }
    $graphModules.Add([pscustomobject][ordered]@{
        path = $pair.path
        version = $pair.version
        goModSum = $pair.goModSum
        infoBytes = [long](Get-Item -LiteralPath $download.Info).Length
        infoSha256 = Get-Sha256 $download.Info
        modBytes = [long](Get-Item -LiteralPath $download.GoMod).Length
        modSha256 = Get-Sha256 $download.GoMod
    })
}
if ($authorizedGoModMissing -ne 0 -or $authorizedGoModMismatch -ne 0) {
    throw "Authorized go.mod metadata verification failed: missing=$authorizedGoModMissing, mismatch=$authorizedGoModMismatch"
}
$allSourceModules = [System.Collections.Generic.List[object]]::new()
$missingFiles = 0
$missingContentAuthority = 0
$missingGoModAuthority = 0
$actualContentMismatch = 0
$actualGoModMismatch = 0
$downloadRoot = Join-Path $acquireState 'gomodcache\cache\download'
foreach ($dependency in $dependencies) {
    $expectedContent = $goSumEntries["$($dependency.Path) $($dependency.Version)"]
    $expectedGoMod = $goSumEntries["$($dependency.Path) $($dependency.Version)/go.mod"]
    if ($null -eq $expectedContent) { $missingContentAuthority++ }
    if ($null -eq $expectedGoMod) { $missingGoModAuthority++ }
}
foreach ($download in $downloads) {
    $expectedContent = $goSumEntries["$($download.Path) $($download.Version)"]
    $expectedGoMod = $goSumEntries["$($download.Path) $($download.Version)/go.mod"]
    if ($null -ne $expectedContent -and $download.Sum -cne $expectedContent) { $actualContentMismatch++ }
    if ($null -ne $expectedGoMod -and $download.GoModSum -cne $expectedGoMod) { $actualGoModMismatch++ }
    foreach ($sourcePath in @([string]$download.Info, [string]$download.GoMod, [string]$download.Zip)) {
        if ([string]::IsNullOrWhiteSpace($sourcePath) -or -not (Test-Path -LiteralPath $sourcePath -PathType Leaf)) {
            $missingFiles++
        }
    }
    $allSourceModules.Add([pscustomobject][ordered]@{
        path = [string]$download.Path
        version = [string]$download.Version
        contentSum = [string]$download.Sum
        goModSum = [string]$download.GoModSum
        hasContentAuthority = $null -ne $expectedContent
        hasGoModAuthority = $null -ne $expectedGoMod
        infoBytes = [long](Get-Item -LiteralPath $download.Info).Length
        infoSha256 = Get-Sha256 $download.Info
        modBytes = [long](Get-Item -LiteralPath $download.GoMod).Length
        modSha256 = Get-Sha256 $download.GoMod
        zipBytes = [long](Get-Item -LiteralPath $download.Zip).Length
        zipSha256 = Get-Sha256 $download.Zip
        zipPath = [string]$download.Zip
        dir = [string]$download.Dir
    })
}
if ($missingFiles -ne 0 -or $actualContentMismatch -ne 0 -or $actualGoModMismatch -ne 0) {
    throw "Full-list acquisition failed: missingFiles=$missingFiles, contentMismatch=$actualContentMismatch, goModMismatch=$actualGoModMismatch"
}

Set-GoEnvironment $acquireState 'off'
Assert-ModuleIdentity $script:RepoRoot
$allSourceById = @{}
foreach ($module in $allSourceModules) { $allSourceById["$($module.path)@$($module.version)"] = $module }
$sourceModules = [System.Collections.Generic.List[object]]::new()
$vendorMissingContentAuthority = 0
$vendorMissingGoModAuthority = 0
foreach ($vendorModule in $vendorDiscoveryModules) {
    $identity = "$($vendorModule.path)@$($vendorModule.version)"
    $sourceModule = $allSourceById[$identity]
    if ($null -eq $sourceModule) { throw "Discovered vendor module is absent from acquisition: $identity" }
    if (-not [bool]$sourceModule.hasContentAuthority) { $vendorMissingContentAuthority++ }
    if (-not [bool]$sourceModule.hasGoModAuthority) { $vendorMissingGoModAuthority++ }
    $sourceModules.Add($sourceModule)
}
if ($vendorMissingContentAuthority -ne 0 -or $vendorMissingGoModAuthority -ne 0) {
    throw "Vendor closure lacks unchanged go.sum authority: content=$vendorMissingContentAuthority, goMod=$vendorMissingGoModAuthority"
}
$vendorDiscoveryIds = @($vendorDiscoveryModules | ForEach-Object { "$($_.path)@$($_.version)" } | Sort-Object -CaseSensitive)

[System.IO.Directory]::CreateDirectory($inputProxy) | Out-Null
$versionsByDirectory = @{}
foreach ($module in $graphModules) {
    $download = $authorizedGoModDownloadById["$($module.path)@$($module.version)"]
    foreach ($sourcePath in @([string]$download.Info, [string]$download.GoMod)) {
        $relative = Get-StrictRelativePath -BasePath $downloadRoot -ChildPath $sourcePath
        if ($relative.StartsWith('..')) { throw "Proxy source escaped download root: $sourcePath" }
        $destination = Join-Path $inputProxy $relative
        [System.IO.Directory]::CreateDirectory((Split-Path -Parent $destination)) | Out-Null
        [System.IO.File]::WriteAllBytes($destination, [System.IO.File]::ReadAllBytes($sourcePath))
    }
    $versionDirectory = Split-Path -Parent (Get-StrictRelativePath -BasePath $downloadRoot -ChildPath ([string]$download.Info))
    if (-not $versionsByDirectory.ContainsKey($versionDirectory)) {
        $versionsByDirectory[$versionDirectory] = [System.Collections.Generic.List[string]]::new()
    }
    $versionsByDirectory[$versionDirectory].Add($module.version)
}
foreach ($module in $sourceModules) {
    $download = $downloadById["$($module.path)@$($module.version)"]
    $sourcePath = [string]$download.Zip
    $relative = Get-StrictRelativePath -BasePath $downloadRoot -ChildPath $sourcePath
    if ($relative.StartsWith('..')) { throw "Proxy source escaped download root: $sourcePath" }
    $destination = Join-Path $inputProxy $relative
    [System.IO.Directory]::CreateDirectory((Split-Path -Parent $destination)) | Out-Null
    [System.IO.File]::WriteAllBytes($destination, [System.IO.File]::ReadAllBytes($sourcePath))
}
foreach ($versionDirectory in @($versionsByDirectory.Keys | Sort-Object -CaseSensitive)) {
    $versions = @($versionsByDirectory[$versionDirectory] | Sort-Object -Unique -CaseSensitive)
    Write-Utf8Lf (Join-Path (Join-Path $inputProxy $versionDirectory) 'list') (($versions -join "`n") + "`n")
}
Assert-NoReparsePoints $inputProxy 'Ephemeral file proxy'
$proxyRows = @(Get-InventoryRows $inputProxy)
$proxyTreeSha256 = Get-TreeDigest $proxyRows
Write-Utf8Lf (Join-Path $evidenceRoot 'input-proxy-inventory.json') (($proxyRows | ConvertTo-Json -Depth 5) + "`n")

$offlineModule = Join-Path $offlineRoot 'module'
$offlineState = Join-Path $offlineRoot 'state'
Copy-ModuleAuthority $offlineModule
$fileProxyUri = Convert-ToFileProxyUri $inputProxy
Set-GoEnvironment $offlineState $fileProxyUri
$offlineDownloadArguments = @('mod', 'download', '-json') + @($sourceModules | ForEach-Object { "$($_.path)@$($_.version)" })
$offlineDownloadResult = Invoke-GoLogged '101-go-mod-download-verified-closure-offline' $offlineModule $offlineDownloadArguments
$offlineDownloads = @(ConvertFrom-JsonSequence $offlineDownloadResult.Stdout)
$offlineById = @{}; foreach ($download in $offlineDownloads) { $offlineById["$($download.Path)@$($download.Version)"] = $download }
$offlineChecksumMismatch = 0
foreach ($module in $sourceModules) {
    $download = $offlineById["$($module.path)@$($module.version)"]
    if ($null -eq $download -or $download.Sum -cne $module.contentSum -or $download.GoModSum -cne $module.goModSum) {
        $offlineChecksumMismatch++
    }
}
if ($offlineDownloads.Count -ne $sourceModules.Count -or $offlineChecksumMismatch -ne 0) {
    throw "Offline rehydration mismatch: downloads=$($offlineDownloads.Count), expected=$($sourceModules.Count), checksums=$offlineChecksumMismatch"
}
$offlineListArguments = @('list', '-mod=mod', '-m', '-json') + @($sourceModules | ForEach-Object { "$($_.path)@$($_.version)" })
$offlineListResult = Invoke-GoLogged '102-go-list-source-modules-offline' $offlineModule $offlineListArguments
$offlineBuildList = @(ConvertFrom-JsonSequence $offlineListResult.Stdout)
$onlineBuildListIds = @($sourceModules | ForEach-Object { "$($_.path)@$($_.version)" } | Sort-Object -CaseSensitive)
$offlineBuildListIds = @($offlineBuildList | ForEach-Object { "$($_.Path)@$($_.Version)" } | Sort-Object -CaseSensitive)
$offlineBuildListDiff = @(Compare-Object -ReferenceObject $onlineBuildListIds -DifferenceObject $offlineBuildListIds -CaseSensitive)
if ($offlineBuildListDiff.Count -ne 0) { throw "Offline module-list difference count: $($offlineBuildListDiff.Count)" }
$offlineGraphArguments = @('list', '-mod=mod', '-m', '-json') + @($authorizedGoModPairs | ForEach-Object { "$($_.path)@$($_.version)" })
$offlineGraphResult = Invoke-GoLogged '103-go-list-authorized-graph-metadata-offline' $offlineModule $offlineGraphArguments
$offlineGraphRows = @(ConvertFrom-JsonSequence $offlineGraphResult.Stdout)
$offlineGraphMismatch = 0
foreach ($pair in $authorizedGoModPairs) {
    $row = @($offlineGraphRows | Where-Object { $_.Path -ceq $pair.path -and $_.Version -ceq $pair.version })
    if ($row.Count -ne 1 -or [string]$row[0].GoModSum -cne [string]$pair.goModSum) { $offlineGraphMismatch++ }
}
if ($offlineGraphRows.Count -ne $authorizedGoModPairs.Count -or $offlineGraphMismatch -ne 0) { throw "Offline graph metadata mismatch: rows=$($offlineGraphRows.Count), checksums=$offlineGraphMismatch" }
$offlineVerifyResult = Invoke-GoLogged '104-go-mod-verify-offline' $offlineModule @('mod', 'verify')
if ($offlineVerifyResult.Stdout.Trim() -cne 'all modules verified') { throw 'Offline go mod verify did not pass.' }
Assert-ModuleIdentity $offlineModule
Assert-ModuleIdentity $script:RepoRoot

$vendorOutputs = @()
$vendorModuleListDifferenceCount = 0
$vendorGoModVerifyPassCount = 0
$runNumber = 0
foreach ($vendorRunRoot in @($vendorRun1Root, $vendorRun2Root)) {
    $runNumber++
    [System.IO.Directory]::CreateDirectory($vendorRunRoot) | Out-Null
    $vendorModuleRoot = Join-Path $vendorRunRoot 'module'
    Copy-ModuleAuthority $vendorModuleRoot
    $vendorOutput = Join-Path $vendorModuleRoot 'vendor'
    $vendorState = Join-Path $vendorRunRoot 'state'
    Set-GoEnvironment $vendorState $fileProxyUri
    Invoke-GoLogged ("20$runNumber-go-mod-vendor") $script:RepoRoot @('mod', 'vendor', '-o', $vendorOutput) | Out-Null
    if (-not (Test-Path -LiteralPath (Join-Path $vendorOutput 'modules.txt') -PathType Leaf)) {
        throw "Vendor run $runNumber did not create modules.txt."
    }
    Assert-NoReparsePoints $vendorOutput "Vendor run $runNumber"
    Set-GoEnvironment $vendorState 'off'
    $vendorListArguments = @('list', '-mod=vendor', '-m', '-json') + @($vendorDiscoveryModules | ForEach-Object { $_.path })
    $vendorListResult = Invoke-GoLogged ("21$runNumber-go-list-vendor") $vendorModuleRoot $vendorListArguments
    $vendorList = @(ConvertFrom-JsonSequence $vendorListResult.Stdout)
    $vendorListIds = @($vendorList | Where-Object { -not ($_.PSObject.Properties.Name -contains 'Main') } | ForEach-Object { "$($_.Path)@$($_.Version)" } | Sort-Object -CaseSensitive)
    $vendorListDiff = @(Compare-Object -ReferenceObject $vendorDiscoveryIds -DifferenceObject $vendorListIds -CaseSensitive)
    $vendorModuleListDifferenceCount += $vendorListDiff.Count
    if ($vendorListDiff.Count -ne 0) { throw "Vendor run $runNumber module-list difference count: $($vendorListDiff.Count)" }
    $env:GOFLAGS = '-mod=vendor'
    $vendorVerifyResult = Invoke-GoLogged ("22$runNumber-go-mod-verify-vendor") $vendorModuleRoot @('mod', 'verify')
    if ($vendorVerifyResult.Stdout.Trim() -cne 'all modules verified') { throw "Vendor run $runNumber go mod verify did not pass." }
    $vendorGoModVerifyPassCount++
    $env:GOFLAGS = ''
    Assert-ModuleIdentity $vendorModuleRoot
    Assert-ModuleIdentity $script:RepoRoot
    $baselineRows = @(Get-InventoryRows $vendorOutput)
    $baselineTreeSha256 = Get-TreeDigest $baselineRows
    $closure = Invoke-SupplementalClosure $vendorOutput $sourceModules
    $reviewedPatch = Assert-AndApplyReviewedPatch $vendorOutput $sourceModules
    $finalRows = @(Get-InventoryRows $vendorOutput)
    $vendorOutputs += [pscustomobject][ordered]@{
        root = $vendorOutput
        baselineRows = $baselineRows
        baselineTreeSha256 = $baselineTreeSha256
        closure = $closure
        reviewedPatch = $reviewedPatch
        finalRows = $finalRows
        finalTreeSha256 = Get-TreeDigest $finalRows
    }
}
$baselineDiff = @(Compare-Object -ReferenceObject @($vendorOutputs[0].baselineRows | ForEach-Object { "$($_.path)`t$($_.bytes)`t$($_.sha256)" }) -DifferenceObject @($vendorOutputs[1].baselineRows | ForEach-Object { "$($_.path)`t$($_.bytes)`t$($_.sha256)" }) -CaseSensitive)
if ($baselineDiff.Count -ne 0 -or $vendorOutputs[0].baselineTreeSha256 -cne $vendorOutputs[1].baselineTreeSha256) {
    throw "Standard vendor determinism failed: rows=$($baselineDiff.Count), first=$($vendorOutputs[0].baselineTreeSha256), second=$($vendorOutputs[1].baselineTreeSha256)"
}
$closureFirst = [ordered]@{
    roots = $vendorOutputs[0].closure.roots; files = $vendorOutputs[0].closure.files
    activeSeedTargets = $vendorOutputs[0].closure.activeSeedTargets; recursiveEdges = $vendorOutputs[0].closure.recursiveEdges
    inactiveConditionalReferences = $vendorOutputs[0].closure.inactiveConditionalReferences
    effectiveDefines = $vendorOutputs[0].closure.effectiveDefines; treeSha256 = $vendorOutputs[0].closure.treeSha256
}
$closureSecond = [ordered]@{
    roots = $vendorOutputs[1].closure.roots; files = $vendorOutputs[1].closure.files
    activeSeedTargets = $vendorOutputs[1].closure.activeSeedTargets; recursiveEdges = $vendorOutputs[1].closure.recursiveEdges
    inactiveConditionalReferences = $vendorOutputs[1].closure.inactiveConditionalReferences
    effectiveDefines = $vendorOutputs[1].closure.effectiveDefines; treeSha256 = $vendorOutputs[1].closure.treeSha256
}
$closureDifferenceCount = if (($closureFirst | ConvertTo-Json -Depth 30 -Compress) -ceq ($closureSecond | ConvertTo-Json -Depth 30 -Compress)) { 0 } else { 1 }
if ($closureDifferenceCount -ne 0) { throw 'Supplemental closure derivation differs across the two clean standard baselines.' }
$patchDifferenceCount = if (($vendorOutputs[0].reviewedPatch | ConvertTo-Json -Depth 10 -Compress) -ceq ($vendorOutputs[1].reviewedPatch | ConvertTo-Json -Depth 10 -Compress)) { 0 } else { 1 }
if ($patchDifferenceCount -ne 0) { throw 'Reviewed patch application differs across runs.' }
$finalVendorDiff = @(Compare-Object -ReferenceObject @($vendorOutputs[0].finalRows | ForEach-Object { "$($_.path)`t$($_.bytes)`t$($_.sha256)" }) -DifferenceObject @($vendorOutputs[1].finalRows | ForEach-Object { "$($_.path)`t$($_.bytes)`t$($_.sha256)" }) -CaseSensitive)
if ($finalVendorDiff.Count -ne 0 -or $vendorOutputs[0].finalTreeSha256 -cne $vendorOutputs[1].finalTreeSha256) {
    throw "Final vendor determinism failed: rows=$($finalVendorDiff.Count), first=$($vendorOutputs[0].finalTreeSha256), second=$($vendorOutputs[1].finalTreeSha256)"
}

Copy-ModuleAuthority $candidateRoot
$candidateRoot2 = Join-Path $script:Lane 'publish-candidate-2'
Copy-ModuleAuthority $candidateRoot2
$candidate1 = New-CompleteCandidate $candidateRoot $vendorOutputs[0].root $sourceModules
$candidate2 = New-CompleteCandidate $candidateRoot2 $vendorOutputs[1].root $sourceModules
$candidateVendor = $candidate1.vendor
$candidateAuthority = $candidate1.authority
$candidateLicenses = $candidate1.licenses
$modulesText = $candidate1.modulesText
$vendorModules = @($candidate1.vendorModules)
$manifestModuleRows = @($candidate1.manifestModuleRows)
$copiedDispositionCount = [int]$candidate1.copiedDispositionCount
$exceptionDispositionCount = [int]$candidate1.exceptionDispositionCount
$candidateRows1 = @(Get-InventoryRows -Root $candidateRoot -RelativeRoots @('vendor', 'third_party\go-vendor\licenses', 'third_party\go-vendor\patches'))
$candidateRows2 = @(Get-InventoryRows -Root $candidateRoot2 -RelativeRoots @('vendor', 'third_party\go-vendor\licenses', 'third_party\go-vendor\patches'))
$completeCandidateDiff = @(Compare-Object -ReferenceObject @($candidateRows1 | ForEach-Object { "$($_.path)`t$($_.bytes)`t$($_.sha256)" }) -DifferenceObject @($candidateRows2 | ForEach-Object { "$($_.path)`t$($_.bytes)`t$($_.sha256)" }) -CaseSensitive)
if ($completeCandidateDiff.Count -ne 0 -or (Get-TreeDigest $candidateRows1) -cne (Get-TreeDigest $candidateRows2)) {
    throw "Complete candidate determinism failed: $($completeCandidateDiff.Count) difference(s)."
}
if (($candidate1.manifestModuleRows | ConvertTo-Json -Depth 10 -Compress) -cne ($candidate2.manifestModuleRows | ConvertTo-Json -Depth 10 -Compress)) {
    throw 'License/module dispositions differ across complete candidates.'
}

$baselineByPath = [System.Collections.Generic.Dictionary[string, object]]::new([System.StringComparer]::Ordinal)
foreach ($row in $vendorOutputs[0].baselineRows) { $baselineByPath.Add([string]$row.path, $row) }
$supplementalByPath = [System.Collections.Generic.Dictionary[string, object]]::new([System.StringComparer]::Ordinal)
foreach ($row in $vendorOutputs[0].closure.files) {
    if ($baselineByPath.ContainsKey([string]$row.path)) { throw "Supplemental partition overlaps standard baseline: $($row.path)" }
    $supplementalByPath.Add([string]$row.path, $row)
}
$fileRows = @($candidateRows1 | ForEach-Object {
    $origin = ''
    if ($_.path.StartsWith('vendor/', [System.StringComparison]::Ordinal)) {
        $vendorRelative = $_.path.Substring('vendor/'.Length)
        if ($vendorRelative -ceq $PatchedBindingRelativePath) { $origin = 'patched_postimage' }
        elseif ($supplementalByPath.ContainsKey($vendorRelative)) { $origin = 'supplemental' }
        elseif ($baselineByPath.ContainsKey($vendorRelative)) { $origin = 'standard' }
        else { throw "Unclassified final vendor file: $vendorRelative" }
    }
    elseif ($_.path.StartsWith('third_party/go-vendor/licenses/', [System.StringComparison]::Ordinal)) { $origin = 'license' }
    elseif ($_.path.StartsWith('third_party/go-vendor/patches/', [System.StringComparison]::Ordinal)) { $origin = 'reviewed_patch' }
    else { throw "Unclassified complete candidate file: $($_.path)" }
    [pscustomobject][ordered]@{ path = $_.path; bytes = [long]$_.bytes; sha256 = [string]$_.sha256; origin = $origin }
})
$fileRows = @(Sort-InventoryRows $fileRows)
$treeSha256 = Get-TreeDigest $fileRows
$scriptPath = Join-Path $PSScriptRoot 'materialize-go-vendor.ps1'
$sourceInputRows = @($sourceModules | Sort-Object path, version -CaseSensitive | ForEach-Object {
    [ordered]@{
        path = $_.path; version = $_.version; contentSum = $_.contentSum; goModSum = $_.goModSum
        infoBytes = $_.infoBytes; infoSha256 = $_.infoSha256
        modBytes = $_.modBytes; modSha256 = $_.modSha256
        zipBytes = $_.zipBytes; zipSha256 = $_.zipSha256
    }
})
$graphInputRows = @($graphModules | Sort-Object path, version -CaseSensitive | ForEach-Object {
    [ordered]@{
        path = $_.path; version = $_.version; goModSum = $_.goModSum
        infoBytes = $_.infoBytes; infoSha256 = $_.infoSha256
        modBytes = $_.modBytes; modSha256 = $_.modSha256
    }
})
$vendorModuleIds = @($vendorModules | ForEach-Object { "$($_.path)@$($_.version)" })
$vendorModuleListSha256 = Get-CanonicalStringDigest $vendorModuleIds
$baselineBytes1 = [long](($vendorOutputs[0].baselineRows | Measure-Object bytes -Sum).Sum)
$baselineBytes2 = [long](($vendorOutputs[1].baselineRows | Measure-Object bytes -Sum).Sum)
$supplementalBytes = [long](($vendorOutputs[0].closure.files | Measure-Object bytes -Sum).Sum)
$finalVendorBytes1 = [long](($vendorOutputs[0].finalRows | Measure-Object bytes -Sum).Sum)
$finalVendorBytes2 = [long](($vendorOutputs[1].finalRows | Measure-Object bytes -Sum).Sum)
$verifierPath = Join-Path $PSScriptRoot 'verify-go-vendor.ps1'
$manifest = [ordered]@{
    schemaVersion = 1
    closureContractVersion = 2
    authorityKind = 'go_vendor_cgo_closure'
    goVersion = 'go1.26.3'
    generation = [ordered]@{
        script = 'scripts/materialize-go-vendor.ps1'
        scriptSha256 = Get-Sha256 $scriptPath
        verifier = 'scripts/verify-go-vendor.ps1'
        verifierSha256 = Get-Sha256 $verifierPath
        command = 'pwsh -File scripts/materialize-go-vendor.ps1 -LaneRoot <repo-local-ephemeral-root> [-InputProxy <checksum-verified-ephemeral-file-proxy>]'
        sourceProxy = $AllowedProxy
        sourceFallback = 'none'
        vendorCommand = 'go mod vendor -o <ephemeral-output>'
    }
    goModSha256 = $ExpectedGoModSha256
    goSumSha256 = $ExpectedGoSumSha256
    vendorModulesTxtSha256 = Get-Sha256 (Join-Path $candidateVendor 'modules.txt')
    sourceInput = [ordered]@{
        role = 'ephemeral reproducible acquisition input; never runtime, commit, handoff, or recovery authority'
        proxy = $AllowedProxy
        checksumAuthority = 'unchanged go.sum with GOSUMDB=off'
        materializationListModulesIncludingMain = $buildList.Count
        selectedDependencyModules = $dependencies.Count
        materializationListMissingContentAuthority = $missingContentAuthority
        materializationListMissingGoModAuthority = $missingGoModAuthority
        materializationListActualContentMismatches = $actualContentMismatch
        materializationListActualGoModMismatches = $actualGoModMismatch
        verifiedVendorClosureModules = $sourceModules.Count
        goSumAuthorizedGraphModules = $graphModules.Count
        authorizedGraphGoModChecks = $graphModules.Count
        authorizedGraphGoModMismatches = $authorizedGoModMismatch
        contentChecks = $sourceModules.Count
        goModChecks = $sourceModules.Count
        vendorMissingContentAuthority = $vendorMissingContentAuthority
        vendorMissingGoModAuthority = $vendorMissingGoModAuthority
        goNativeModuleVerify = 'all modules verified'
        offlineProxyRehydration = 'PASS'
        offlineVendorModuleListDifferences = $vendorModuleListDifferenceCount
        offlineVendorGoModVerifyRuns = $vendorGoModVerifyPassCount
        offlineChecksumMismatches = $offlineChecksumMismatch
        fileProxyFiles = $proxyRows.Count
        fileProxyBytes = [long](($proxyRows | Measure-Object bytes -Sum).Sum)
        fileProxyTreeSha256 = $proxyTreeSha256
        modules = $sourceInputRows
        graphModules = $graphInputRows
    }
    standardBaseline = [ordered]@{
        pathDomain = 'vendor-relative'
        command = 'go mod vendor -o <ephemeral-output>'
        goVersion = 'go1.26.3'
        runs = 2; equal = $true
        moduleCount = $vendorModules.Count
        moduleListSha256 = $vendorModuleListSha256
        modules = @($vendorModules)
        firstRun = [ordered]@{ files = $vendorOutputs[0].baselineRows.Count; bytes = $baselineBytes1; treeSha256 = $vendorOutputs[0].baselineTreeSha256 }
        secondRun = [ordered]@{ files = $vendorOutputs[1].baselineRows.Count; bytes = $baselineBytes2; treeSha256 = $vendorOutputs[1].baselineTreeSha256 }
        differenceCount = $baselineDiff.Count
        files = @($vendorOutputs[0].baselineRows)
    }
    supplementalClosure = [ordered]@{
        algorithmVersion = 2
        pathDomain = 'vendor-relative'
        runs = 2; equal = $true
        affectedModuleCount = $vendorOutputs[0].closure.affectedModules.Count
        affectedModules = @($vendorOutputs[0].closure.affectedModules)
        rootCount = $vendorOutputs[0].closure.roots.Count
        roots = @($vendorOutputs[0].closure.roots)
        activeSeedTargets = @($vendorOutputs[0].closure.activeSeedTargets)
        recursiveEdges = @($vendorOutputs[0].closure.recursiveEdges)
        inactiveConditionalReferences = @($vendorOutputs[0].closure.inactiveConditionalReferences)
        effectiveDefines = @($vendorOutputs[0].closure.effectiveDefines)
        fileCount = $vendorOutputs[0].closure.files.Count
        bytes = $supplementalBytes
        treeSha256 = $vendorOutputs[0].closure.treeSha256
        secondTreeSha256 = $vendorOutputs[1].closure.treeSha256
        differenceCount = $closureDifferenceCount
        files = @($vendorOutputs[0].closure.files)
    }
    reviewedPatches = @($vendorOutputs[0].reviewedPatch)
    finalDeterminism = [ordered]@{
        runs = 2; equal = $true
        firstRun = [ordered]@{ vendorFiles = $vendorOutputs[0].finalRows.Count; vendorBytes = $finalVendorBytes1; vendorTreeSha256 = $vendorOutputs[0].finalTreeSha256; completeFiles = $candidateRows1.Count; completeBytes = [long](($candidateRows1 | Measure-Object bytes -Sum).Sum); completeTreeSha256 = Get-TreeDigest $candidateRows1 }
        secondRun = [ordered]@{ vendorFiles = $vendorOutputs[1].finalRows.Count; vendorBytes = $finalVendorBytes2; vendorTreeSha256 = $vendorOutputs[1].finalTreeSha256; completeFiles = $candidateRows2.Count; completeBytes = [long](($candidateRows2 | Measure-Object bytes -Sum).Sum); completeTreeSha256 = Get-TreeDigest $candidateRows2 }
        differenceCount = $finalVendorDiff.Count + $completeCandidateDiff.Count + $patchDifferenceCount
    }
    modules = @($manifestModuleRows)
    licenseDisposition = [ordered]@{
        modules = $vendorModules.Count
        copied = $copiedDispositionCount
        reviewedExceptions = $exceptionDispositionCount
        incomplete = 0
    }
    files = $fileRows
    treeSha256 = $treeSha256
}
$manifestText = ($manifest | ConvertTo-Json -Depth 100) + "`n"
Write-Utf8Lf (Join-Path $candidateAuthority 'manifest.v1.json') $manifestText
Write-Utf8Lf (Join-Path $candidate2.authority 'manifest.v1.json') $manifestText
$sealedCandidateRows1 = @(Get-InventoryRows -Root $candidateRoot -RelativeRoots @('vendor', 'third_party\go-vendor'))
$sealedCandidateRows2 = @(Get-InventoryRows -Root $candidateRoot2 -RelativeRoots @('vendor', 'third_party\go-vendor'))
$sealedCandidateDiff = @(Compare-Object -ReferenceObject @($sealedCandidateRows1 | ForEach-Object { "$($_.path)`t$($_.bytes)`t$($_.sha256)" }) -DifferenceObject @($sealedCandidateRows2 | ForEach-Object { "$($_.path)`t$($_.bytes)`t$($_.sha256)" }) -CaseSensitive)
if ($sealedCandidateDiff.Count -ne 0) { throw "Sealed complete candidates differ: $($sealedCandidateDiff.Count)" }

& (Join-Path $PSScriptRoot 'verify-go-vendor.ps1') -SourceRoot $candidateRoot -Json 1> (Join-Path $script:LogsRoot '301-candidate-verifier.stdout.json') 2> (Join-Path $script:LogsRoot '301-candidate-verifier.stderr.log')
if ($LASTEXITCODE -ne 0) { throw "Staged candidate verifier failed with exit $LASTEXITCODE" }
& (Join-Path $PSScriptRoot 'verify-go-vendor.ps1') -SourceRoot $candidateRoot2 -Json 1> (Join-Path $script:LogsRoot '302-candidate-2-verifier.stdout.json') 2> (Join-Path $script:LogsRoot '302-candidate-2-verifier.stderr.log')
if ($LASTEXITCODE -ne 0) { throw "Second staged candidate verifier failed with exit $LASTEXITCODE" }
Publish-AtomicVendorAuthority $candidateRoot
Assert-ModuleIdentity $script:RepoRoot

$durableRows = @(Get-InventoryRows -Root $script:RepoRoot -RelativeRoots @('vendor', 'third_party\go-vendor\licenses', 'third_party\go-vendor\patches'))
$summary = [ordered]@{
    status = 'PASS'
    sourceProxy = $AllowedProxy
    goVersion = 'go1.26.3'
    moduleBuildListIncludingMain = $buildList.Count
    dependencyModules = $dependencies.Count
    fullBuildListMissingContentAuthority = $missingContentAuthority
    fullBuildListMissingGoModAuthority = $missingGoModAuthority
    fullBuildListActualContentMismatches = $actualContentMismatch
    fullBuildListActualGoModMismatches = $actualGoModMismatch
    verifiedVendorClosureModules = $sourceModules.Count
    goSumAuthorizedGraphModules = $graphModules.Count
    authorizedGraphGoModChecks = $graphModules.Count
    authorizedGraphGoModMismatches = $authorizedGoModMismatch
    inputProxyFiles = $proxyRows.Count
    inputProxyBytes = [long](($proxyRows | Measure-Object bytes -Sum).Sum)
    inputProxyTreeSha256 = $proxyTreeSha256
    offlineRehydration = 'PASS'
    offlineVendorModuleListDifferences = $vendorModuleListDifferenceCount
    offlineVendorGoModVerifyRuns = $vendorGoModVerifyPassCount
    offlineChecksumMismatches = $offlineChecksumMismatch
    standardVendorRunsEqual = $true
    vendorModules = $vendorModules.Count
    standardVendorFiles = $vendorOutputs[0].baselineRows.Count
    standardVendorBytes = $baselineBytes1
    standardVendorTreeSha256 = $vendorOutputs[0].baselineTreeSha256
    supplementalModules = $vendorOutputs[0].closure.affectedModules.Count
    supplementalRoots = $vendorOutputs[0].closure.roots.Count
    supplementalFiles = $vendorOutputs[0].closure.files.Count
    supplementalBytes = $supplementalBytes
    supplementalTreeSha256 = $vendorOutputs[0].closure.treeSha256
    reviewedPatches = 1
    finalVendorFiles = $vendorOutputs[0].finalRows.Count
    finalVendorBytes = $finalVendorBytes1
    finalVendorTreeSha256 = $vendorOutputs[0].finalTreeSha256
    completeCandidatesEqual = ($sealedCandidateDiff.Count -eq 0)
    licenseCopiedModules = $copiedDispositionCount
    licenseReviewedExceptions = $exceptionDispositionCount
    durableFiles = $durableRows.Count
    durableBytes = [long](($durableRows | Measure-Object bytes -Sum).Sum)
    durableTreeSha256 = Get-TreeDigest $durableRows
    durableVendor = 'vendor/'
    durableAuthority = 'third_party/go-vendor/'
    ephemeralDisposition = 'disposable; never runtime, commit, handoff, or recovery authority'
    executionAcquisitionMode = $acquisitionMode
}
Write-Utf8Lf (Join-Path $evidenceRoot 'materialization-summary.json') (($summary | ConvertTo-Json -Depth 8) + "`n")
$summary | ConvertTo-Json -Depth 8
