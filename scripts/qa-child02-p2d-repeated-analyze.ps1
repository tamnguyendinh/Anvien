[CmdletBinding()]
param(
  [Parameter(Mandatory = $true)]
  [ValidateSet("Prepare", "Build", "Run")]
  [string]$Mode
)

$ErrorActionPreference = "Stop"
Set-StrictMode -Version Latest

$RepoRoot = Split-Path -Parent $PSScriptRoot
$LaneRoot = Join-Path $RepoRoot ".tmp\qa-child02-p2d"
$OwnerMarker = Join-Path $LaneRoot "owner.json"
$FixtureRoot = Join-Path $LaneRoot "fixture"
$FixtureSource = Join-Path $FixtureRoot "src\identity.ts"
$AcceptedFixtureSource = Join-Path $RepoRoot "internal\resolution\testdata\p1c_identity_repo\src\identity.ts"
$IsolatedHome = Join-Path $LaneRoot "anvien-home"
$BuildEvidencePath = Join-Path $LaneRoot "build-evidence.json"
$OfficialEvidenceRoot = Join-Path $RepoRoot "reports\QA\child02-p2d-repeated-analyze"
$CliPath = Join-Path $RepoRoot "anvien\bin\anvien.exe"
$RegistryName = "qa-child02-p2d"

function Write-JsonFile([string]$Path, [object]$Value) {
  $parent = Split-Path -Parent $Path
  if ($parent) {
    New-Item -ItemType Directory -Path $parent -Force | Out-Null
  }
  $json = $Value | ConvertTo-Json -Depth 100
  [System.IO.File]::WriteAllText($Path, $json + [Environment]::NewLine, [System.Text.UTF8Encoding]::new($false))
}

function Get-Sha256([string]$Path) {
  if (-not (Test-Path -LiteralPath $Path -PathType Leaf)) {
    return $null
  }
  return (Get-FileHash -LiteralPath $Path -Algorithm SHA256).Hash
}

function Get-StringSha256([string]$Value) {
  $bytes = [System.Text.Encoding]::UTF8.GetBytes($Value)
  $sha = [System.Security.Cryptography.SHA256]::Create()
  try {
    return -join ($sha.ComputeHash($bytes) | ForEach-Object { $_.ToString("X2") })
  } finally {
    $sha.Dispose()
  }
}

function Assert-LaneOwnership {
  if (-not (Test-Path -LiteralPath $OwnerMarker -PathType Leaf)) {
    throw "Lane ownership marker is missing: $OwnerMarker. Run -Mode Prepare first."
  }
  $owner = Get-Content -LiteralPath $OwnerMarker -Raw | ConvertFrom-Json
  if ($owner.scope -ne "child02-p2d") {
    throw "Lane path is not owned by child02-p2d: $LaneRoot"
  }
}

function Format-Command([string]$FilePath, [string[]]$Arguments) {
  $parts = @($FilePath)
  foreach ($argument in $Arguments) {
    if ($argument -match '[\s"]') {
      $parts += '"' + $argument.Replace('"', '\"') + '"'
    } else {
      $parts += $argument
    }
  }
  return $parts -join " "
}

function Invoke-CapturedProcess {
  param(
    [Parameter(Mandatory = $true)][string]$FilePath,
    [string[]]$Arguments = @(),
    [Parameter(Mandatory = $true)][string]$WorkingDirectory,
    [hashtable]$Environment = @{},
    [AllowNull()][string]$InputText = $null
  )

  $startUtc = [DateTime]::UtcNow
  $stopwatch = [System.Diagnostics.Stopwatch]::StartNew()
  $startInfo = [System.Diagnostics.ProcessStartInfo]::new()
  $startInfo.FileName = $FilePath
  $startInfo.WorkingDirectory = $WorkingDirectory
  $startInfo.UseShellExecute = $false
  $startInfo.RedirectStandardOutput = $true
  $startInfo.RedirectStandardError = $true
  $startInfo.RedirectStandardInput = $true
  $startInfo.CreateNoWindow = $true
  $quotedArguments = foreach ($argument in $Arguments) {
    if ($argument -match '[\s"]') {
      '"' + $argument.Replace('"', '\"') + '"'
    } else {
      $argument
    }
  }
  $startInfo.Arguments = $quotedArguments -join " "
  foreach ($key in $Environment.Keys) {
    $startInfo.Environment[$key] = [string]$Environment[$key]
  }

  $process = [System.Diagnostics.Process]::new()
  $process.StartInfo = $startInfo
  if (-not $process.Start()) {
    throw "Failed to start: $(Format-Command $FilePath $Arguments)"
  }
  $stdoutTask = $process.StandardOutput.ReadToEndAsync()
  $stderrTask = $process.StandardError.ReadToEndAsync()
  if ($null -ne $InputText) {
    $process.StandardInput.Write($InputText)
  }
  $process.StandardInput.Close()
  $process.WaitForExit()
  $stdout = $stdoutTask.GetAwaiter().GetResult()
  $stderr = $stderrTask.GetAwaiter().GetResult()
  $stopwatch.Stop()
  $peakWorkingSet = 0
  try { $peakWorkingSet = $process.PeakWorkingSet64 } catch { }

  return [ordered]@{
    command = Format-Command $FilePath $Arguments
    filePath = $FilePath
    arguments = @($Arguments)
    workingDirectory = $WorkingDirectory
    startUtc = $startUtc.ToString("o")
    endUtc = [DateTime]::UtcNow.ToString("o")
    durationSeconds = [Math]::Round($stopwatch.Elapsed.TotalSeconds, 3)
    exitCode = $process.ExitCode
    peakWorkingSetBytes = $peakWorkingSet
    stdout = $stdout.TrimEnd()
    stderr = $stderr.TrimEnd()
  }
}

function Get-InputManifest {
  $git = Get-Command git -ErrorAction Stop
  $result = Invoke-CapturedProcess -FilePath $git.Source -Arguments @("-C", $FixtureRoot, "ls-files", "-co", "--exclude-standard") -WorkingDirectory $RepoRoot
  if ($result.exitCode -ne 0) {
    throw "Failed to enumerate fixture inputs: $($result.stderr)"
  }
  $rows = @()
  foreach ($relative in ($result.stdout -split "`r?`n" | Where-Object { $_ })) {
    $normalized = $relative.Replace("\", "/")
    if ($normalized -eq ".anvien" -or $normalized.StartsWith(".anvien/")) {
      continue
    }
    $full = Join-Path $FixtureRoot $relative
    if (Test-Path -LiteralPath $full -PathType Leaf) {
      $rows += [ordered]@{
        path = $normalized
        sha256 = Get-Sha256 $full
        bytes = (Get-Item -LiteralPath $full).Length
      }
    }
  }
  $rows = @($rows | Sort-Object path)
  $canonical = $rows | ConvertTo-Json -Depth 10 -Compress
  return [ordered]@{
    count = $rows.Count
    signature = Get-StringSha256 $canonical
    files = $rows
  }
}

function Get-PropertyValue([object]$Properties, [string]$Name) {
  if ($null -eq $Properties) { return $null }
  $property = $Properties.PSObject.Properties[$Name]
  if ($null -eq $property) { return $null }
  return $property.Value
}

function Get-GraphFactManifest([string]$GraphPath) {
  $graph = Get-Content -LiteralPath $GraphPath -Raw | ConvertFrom-Json
  $nodeByID = @{}
  foreach ($node in $graph.nodes) { $nodeByID[[string]$node.id] = $node }
  $defines = @($graph.relationships | Where-Object { $_.type -eq "DEFINES" })
  $definitionIDs = @($defines | ForEach-Object { [string]$_.targetId } | Sort-Object -Unique)
  $rows = @()
  foreach ($id in $definitionIDs) {
    if (-not $nodeByID.ContainsKey($id)) { continue }
    $node = $nodeByID[$id]
    $properties = $node.properties
    if ((Get-PropertyValue $properties "filePath") -ne "src/identity.ts") { continue }
    $rows += [ordered]@{
      id = [string]$node.id
      label = [string]$node.label
      name = [string](Get-PropertyValue $properties "name")
      qualifiedName = [string](Get-PropertyValue $properties "qualifiedName")
      filePath = [string](Get-PropertyValue $properties "filePath")
      startLine = Get-PropertyValue $properties "startLine"
      startCol = Get-PropertyValue $properties "startCol"
      endLine = Get-PropertyValue $properties "endLine"
      endCol = Get-PropertyValue $properties "endCol"
      selectionStartLine = Get-PropertyValue $properties "selectionStartLine"
      selectionStartCol = Get-PropertyValue $properties "selectionStartCol"
      selectionEndLine = Get-PropertyValue $properties "selectionEndLine"
      selectionEndCol = Get-PropertyValue $properties "selectionEndCol"
    }
  }
  $rows = @($rows | Sort-Object id)
  $rowJSON = $rows | ConvertTo-Json -Depth 20 -Compress
  $missingEndpoints = @($graph.relationships | Where-Object {
    -not $nodeByID.ContainsKey([string]$_.sourceId) -or -not $nodeByID.ContainsKey([string]$_.targetId)
  })
  $fixtureDefines = @($defines | Where-Object { $definitionIDs -contains [string]$_.targetId } | ForEach-Object {
    [ordered]@{ id = [string]$_.id; sourceId = [string]$_.sourceId; targetId = [string]$_.targetId; type = [string]$_.type }
  } | Sort-Object targetId)
  $nameCounts = [ordered]@{}
  foreach ($name in @($rows | ForEach-Object { $_.name } | Sort-Object -Unique)) {
    $nameCounts[$name] = @($rows | Where-Object { $_.name -eq $name }).Count
  }
  return [ordered]@{
    graphNodes = @($graph.nodes).Count
    graphRelationships = @($graph.relationships).Count
    definitionCount = $rows.Count
    definesCount = $fixtureDefines.Count
    missingEndpointCount = $missingEndpoints.Count
    definitionSignature = Get-StringSha256 $rowJSON
    nameCounts = $nameCounts
    timeStartLines = @($rows | Where-Object { $_.name -eq "time" } | ForEach-Object { $_.startLine } | Sort-Object)
    nowStartLines = @($rows | Where-Object { $_.name -eq "now" } | ForEach-Object { $_.startLine } | Sort-Object)
    laterStartLines = @($rows | Where-Object { $_.name -eq "later" } | ForEach-Object { $_.startLine } | Sort-Object)
    definitions = $rows
    defines = $fixtureDefines
  }
}

function Get-ArtifactManifest {
  $storage = Join-Path $FixtureRoot ".anvien"
  $rows = [ordered]@{}
  foreach ($name in @("graph.json", "lbug", "meta.json")) {
    $path = Join-Path $storage $name
    if (Test-Path -LiteralPath $path -PathType Leaf) {
      $item = Get-Item -LiteralPath $path
      $rows[$name] = [ordered]@{
        path = $path
        sha256 = Get-Sha256 $path
        bytes = $item.Length
        lastWriteUtc = $item.LastWriteTimeUtc.ToString("o")
      }
    } else {
      $rows[$name] = [ordered]@{ path = $path; exists = $false }
    }
  }
  return $rows
}

function Invoke-NormalAnalyze([string]$Label) {
  $before = Get-InputManifest
  $environment = @{ ANVIEN_HOME = $IsolatedHome }
  $process = Invoke-CapturedProcess -FilePath $CliPath -Arguments @("analyze", $FixtureRoot, "--force", "--json", "--name", $RegistryName) -WorkingDirectory $RepoRoot -Environment $environment
  $after = Get-InputManifest
  $payload = $null
  if ($process.exitCode -eq 0 -and $process.stdout) {
    try { $payload = $process.stdout | ConvertFrom-Json } catch { }
  }
  $artifacts = Get-ArtifactManifest
  $facts = $null
  $graphPath = Join-Path $FixtureRoot ".anvien\graph.json"
  if ($process.exitCode -eq 0 -and (Test-Path -LiteralPath $graphPath -PathType Leaf)) {
    $facts = Get-GraphFactManifest $graphPath
  }
  return [ordered]@{
    label = $Label
    inputBefore = $before
    process = $process
    inputAfter = $after
    json = $payload
    artifacts = $artifacts
    facts = $facts
  }
}

function Invoke-McpCypher([string]$Label, [string]$Query) {
  $request = [ordered]@{
    jsonrpc = "2.0"
    id = 1
    method = "tools/call"
    params = [ordered]@{
      name = "cypher"
      arguments = [ordered]@{ repo = $RegistryName; query = $Query }
    }
  }
  $inputText = ($request | ConvertTo-Json -Depth 20 -Compress) + "`n"
  $process = Invoke-CapturedProcess -FilePath $CliPath -Arguments @("mcp") -WorkingDirectory $RepoRoot -Environment @{ ANVIEN_HOME = $IsolatedHome } -InputText $inputText
  $response = $null
  $toolPayload = $null
  if ($process.stdout) {
    try { $response = $process.stdout | ConvertFrom-Json } catch { }
  }
  $hasResponseError = $null -ne $response -and $null -ne $response.PSObject.Properties['error']
  if ($null -ne $response -and -not $hasResponseError) {
    try {
      $text = [string]$response.result.content[0].text
      $jsonText = ($text -split "`r?`n`r?`n---", 2)[0]
      $toolPayload = $jsonText | ConvertFrom-Json
    } catch { }
  }
  return [ordered]@{ label = $Label; query = $Query; process = $process; response = $response; toolPayload = $toolPayload }
}

function Get-McpVariableNames([object]$Read) {
  if ($null -eq $Read.toolPayload) { return @() }
  return @($Read.toolPayload.rows | ForEach-Object { [string]$_.name } | Sort-Object)
}

function Add-RestartManagerType {
  if ("P2DRestartManager" -as [type]) { return }
  Add-Type -TypeDefinition @'
using System;
using System.Collections.Generic;
using System.Runtime.InteropServices;
using System.Text;

public static class P2DRestartManager {
    private const int CCH_RM_SESSION_KEY = 32;
    private const int CCH_RM_MAX_APP_NAME = 255;
    private const int CCH_RM_MAX_SVC_NAME = 63;
    private const int ERROR_MORE_DATA = 234;

    [StructLayout(LayoutKind.Sequential)]
    private struct RM_UNIQUE_PROCESS {
        public int dwProcessId;
        public System.Runtime.InteropServices.ComTypes.FILETIME ProcessStartTime;
    }

    private enum RM_APP_TYPE {
        RmUnknownApp = 0, RmMainWindow = 1, RmOtherWindow = 2,
        RmService = 3, RmExplorer = 4, RmConsole = 5, RmCritical = 1000
    }

    [StructLayout(LayoutKind.Sequential, CharSet = CharSet.Unicode)]
    private struct RM_PROCESS_INFO {
        public RM_UNIQUE_PROCESS Process;
        [MarshalAs(UnmanagedType.ByValTStr, SizeConst = CCH_RM_MAX_APP_NAME + 1)]
        public string strAppName;
        [MarshalAs(UnmanagedType.ByValTStr, SizeConst = CCH_RM_MAX_SVC_NAME + 1)]
        public string strServiceShortName;
        public RM_APP_TYPE ApplicationType;
        public uint AppStatus;
        public uint TSSessionId;
        [MarshalAs(UnmanagedType.Bool)] public bool bRestartable;
    }

    [DllImport("rstrtmgr.dll", CharSet = CharSet.Unicode)]
    private static extern int RmStartSession(out uint handle, int flags, StringBuilder sessionKey);
    [DllImport("rstrtmgr.dll", CharSet = CharSet.Unicode)]
    private static extern int RmRegisterResources(uint handle, uint fileCount, string[] fileNames, uint appCount, IntPtr apps, uint serviceCount, string[] serviceNames);
    [DllImport("rstrtmgr.dll")]
    private static extern int RmGetList(uint handle, out uint needed, ref uint count, [In, Out] RM_PROCESS_INFO[] affectedApps, ref uint rebootReasons);
    [DllImport("rstrtmgr.dll")]
    private static extern int RmEndSession(uint handle);

    public static int[] Query(string path) {
        uint handle;
        var key = new StringBuilder(CCH_RM_SESSION_KEY + 1);
        int result = RmStartSession(out handle, 0, key);
        if (result != 0) throw new InvalidOperationException("RmStartSession failed: " + result);
        try {
            result = RmRegisterResources(handle, 1, new[] { path }, 0, IntPtr.Zero, 0, null);
            if (result != 0) throw new InvalidOperationException("RmRegisterResources failed: " + result);
            uint needed = 0, count = 0, reasons = 0;
            result = RmGetList(handle, out needed, ref count, null, ref reasons);
            if (result == 0) return new int[0];
            if (result != ERROR_MORE_DATA) throw new InvalidOperationException("RmGetList(size) failed: " + result);
            var info = new RM_PROCESS_INFO[needed];
            count = needed;
            result = RmGetList(handle, out needed, ref count, info, ref reasons);
            if (result != 0) throw new InvalidOperationException("RmGetList(data) failed: " + result);
            var output = new List<int>();
            for (int i = 0; i < count; i++) output.Add(info[i].Process.dwProcessId);
            return output.ToArray();
        } finally {
            RmEndSession(handle);
        }
    }
}
'@
}

function Get-ProcessRecord([int]$ProcessID) {
  $process = Get-CimInstance Win32_Process -Filter "ProcessId = $ProcessID" -ErrorAction SilentlyContinue
  if ($null -eq $process) { return $null }
  return [ordered]@{
    pid = [int]$process.ProcessId
    parentPid = [int]$process.ParentProcessId
    name = [string]$process.Name
    executablePath = [string]$process.ExecutablePath
    commandLine = [string]$process.CommandLine
  }
}

function Test-VerifiedLauncher([object]$ProcessRecord) {
  if ($null -eq $ProcessRecord) { return $false }
  $name = ([string]$ProcessRecord.name).ToLowerInvariant()
  if ($name -notin @("cmd.exe", "powershell.exe", "pwsh.exe", "node.exe")) { return $false }
  $commandLine = [string]$ProcessRecord.commandLine
  return $commandLine -match '(?i)anvien' -and $commandLine -match '(?i)(mcp|serve|analyze|anvien\.cmd)'
}

function Test-ExclusiveOpen([string]$Path) {
  if (-not (Test-Path -LiteralPath $Path -PathType Leaf)) {
    return [ordered]@{ path = $Path; exists = $false; success = $true; note = "artifact absent before build" }
  }
  try {
    $stream = [System.IO.File]::Open($Path, [System.IO.FileMode]::Open, [System.IO.FileAccess]::ReadWrite, [System.IO.FileShare]::None)
    $stream.Dispose()
    return [ordered]@{ path = $Path; exists = $true; success = $true }
  } catch {
    return [ordered]@{ path = $Path; exists = $true; success = $false; error = $_.Exception.Message }
  }
}

function Invoke-CleanHolderGate {
  Add-RestartManagerType
  $globalRoot = Join-Path $env:APPDATA "npm\node_modules\anvien\bin"
  $artifacts = @(
    (Join-Path $RepoRoot "anvien\bin\anvien.exe"),
    (Join-Path $RepoRoot "anvien\bin\lbug_shared.dll"),
    (Join-Path $RepoRoot "anvien-launcher\AnvienLauncher.exe"),
    (Join-Path $RepoRoot "anvien-launcher\server-bundle\anvien-server.exe"),
    (Join-Path $globalRoot "anvien.exe"),
    (Join-Path $globalRoot "lbug_shared.dll")
  )
  $initial = @()
  $holderIDs = [System.Collections.Generic.HashSet[int]]::new()
  foreach ($artifact in $artifacts) {
    $ids = @()
    if (Test-Path -LiteralPath $artifact -PathType Leaf) {
      $ids = @([P2DRestartManager]::Query($artifact) | Sort-Object -Unique)
      foreach ($id in $ids) { [void]$holderIDs.Add([int]$id) }
    }
    $initial += [ordered]@{ path = $artifact; exists = Test-Path -LiteralPath $artifact -PathType Leaf; holderPids = $ids }
  }
  $holders = @()
  $launchers = @()
  foreach ($id in @($holderIDs)) {
    $record = Get-ProcessRecord $id
    if ($null -ne $record) {
      $holders += $record
      $parent = Get-ProcessRecord $record.parentPid
      if (Test-VerifiedLauncher $parent) { $launchers += $parent }
    }
  }
  $launchers = @($launchers | Sort-Object pid -Unique)
  $killed = @()
  foreach ($id in @($holderIDs | Sort-Object -Unique)) {
    $record = Get-ProcessRecord ([int]$id)
    if ($null -eq $record) { continue }
    $live = Get-Process -Id ([int]$id) -ErrorAction SilentlyContinue
    if ($null -ne $live) {
      Stop-Process -Id ([int]$id) -Force -ErrorAction Stop
      $killed += [ordered]@{ pid = [int]$id; name = $record.name; commandLine = $record.commandLine; reason = "restart-manager-holder" }
    }
  }
  foreach ($record in $launchers) {
    $id = [int]$record.pid
    $live = Get-Process -Id $id -ErrorAction SilentlyContinue
    if ($null -ne $live) {
      Stop-Process -Id $id -Force -ErrorAction Stop
      $killed += [ordered]@{ pid = $id; name = $record.name; commandLine = $record.commandLine; reason = "verified-launcher-parent" }
    }
  }
  $final = @()
  foreach ($artifact in $artifacts) {
    $ids = @()
    if (Test-Path -LiteralPath $artifact -PathType Leaf) {
      $ids = @([P2DRestartManager]::Query($artifact) | Sort-Object -Unique)
    }
    $final += [ordered]@{ path = $artifact; exists = Test-Path -LiteralPath $artifact -PathType Leaf; holderPids = $ids; holderCount = $ids.Count }
  }
  $exclusive = @($artifacts | ForEach-Object { Test-ExclusiveOpen $_ })
  $remaining = @($final | ForEach-Object { $_.holderPids }).Count
  $exclusiveFailures = @($exclusive | Where-Object { -not $_.success }).Count
  return [ordered]@{
    timestampUtc = [DateTime]::UtcNow.ToString("o")
    artifacts = $artifacts
    initialRestartManager = $initial
    verifiedHolders = $holders
    verifiedLauncherParents = $launchers
    killed = $killed
    finalRestartManager = $final
    restartManagerZero = ($remaining -eq 0)
    exclusiveOpen = $exclusive
    exclusiveOpenAll = ($exclusiveFailures -eq 0)
  }
}

function Invoke-Prepare {
  if (Test-Path -LiteralPath $LaneRoot) {
    if (-not (Test-Path -LiteralPath $OwnerMarker -PathType Leaf)) {
      throw "Existing lane root has unknown ownership: $LaneRoot"
    }
    Assert-LaneOwnership
    foreach ($owned in @($FixtureRoot, $IsolatedHome, $BuildEvidencePath)) {
      if (Test-Path -LiteralPath $owned) { Remove-Item -LiteralPath $owned -Recurse -Force }
    }
  } else {
    New-Item -ItemType Directory -Path $LaneRoot -Force | Out-Null
    Write-JsonFile $OwnerMarker ([ordered]@{ scope = "child02-p2d"; owner = "independent-qa"; createdUtc = [DateTime]::UtcNow.ToString("o") })
  }
  New-Item -ItemType Directory -Path (Split-Path -Parent $FixtureSource) -Force | Out-Null
  New-Item -ItemType Directory -Path $IsolatedHome -Force | Out-Null
  Copy-Item -LiteralPath $AcceptedFixtureSource -Destination $FixtureSource -Force

  $git = (Get-Command git -ErrorAction Stop).Source
  $commands = @(
    @("init", "-b", "master", $FixtureRoot),
    @("-C", $FixtureRoot, "config", "user.name", "P2-D QA"),
    @("-C", $FixtureRoot, "config", "user.email", "p2d-qa@local.invalid"),
    @("-C", $FixtureRoot, "add", "src/identity.ts"),
    @("-C", $FixtureRoot, "commit", "-m", "fixture: accepted child01 identity source")
  )
  $results = @()
  foreach ($arguments in $commands) {
    $result = Invoke-CapturedProcess -FilePath $git -Arguments $arguments -WorkingDirectory $RepoRoot
    $results += $result
    if ($result.exitCode -ne 0) { throw "Fixture preparation failed: $($result.command)`n$($result.stderr)" }
  }
  $prepared = [ordered]@{
    laneRoot = $LaneRoot
    fixtureRoot = $FixtureRoot
    isolatedAnvienHome = $IsolatedHome
    acceptedSource = $AcceptedFixtureSource
    acceptedSourceSha256 = Get-Sha256 $AcceptedFixtureSource
    fixtureSourceSha256 = Get-Sha256 $FixtureSource
    commands = $results
    inputManifest = Get-InputManifest
  }
  Write-JsonFile (Join-Path $LaneRoot "prepare-evidence.json") $prepared
  $prepared | ConvertTo-Json -Depth 20
}

function Invoke-Build {
  Assert-LaneOwnership
  $gate = Invoke-CleanHolderGate
  if (-not $gate.restartManagerZero -or -not $gate.exclusiveOpenAll) {
    Write-JsonFile $BuildEvidencePath ([ordered]@{ holderGate = $gate; build = $null; status = "BLOCKED_PRE_BUILD" })
    throw "Clean-holder gate failed; build was not invoked."
  }
  $npm = (Get-Command npm.cmd -ErrorAction Stop).Source
  $build = Invoke-CapturedProcess -FilePath $npm -Arguments @("run", "full-build") -WorkingDirectory $RepoRoot
  $artifactRows = @()
  foreach ($path in $gate.artifacts) {
    $exists = Test-Path -LiteralPath $path -PathType Leaf
    $artifactRows += [ordered]@{
      path = $path
      exists = $exists
      sha256 = $(if ($exists) { Get-Sha256 $path } else { $null })
      bytes = $(if ($exists) { (Get-Item -LiteralPath $path).Length } else { $null })
      lastWriteUtc = $(if ($exists) { (Get-Item -LiteralPath $path).LastWriteTimeUtc.ToString("o") } else { $null })
    }
  }
  $evidence = [ordered]@{
    holderGate = $gate
    commandRequiredByRepo = "npm run full-build"
    build = $build
    artifactsAfterBuild = $artifactRows
    status = $(if ($build.exitCode -eq 0) { "PASS" } else { "FAIL" })
  }
  Write-JsonFile $BuildEvidencePath $evidence
  $evidence | ConvertTo-Json -Depth 100
  if ($build.exitCode -ne 0) { exit $build.exitCode }
}

function Invoke-Run {
  Assert-LaneOwnership
  if (-not (Test-Path -LiteralPath $BuildEvidencePath -PathType Leaf)) { throw "Build evidence is missing. Run -Mode Build first." }
  $buildEvidence = Get-Content -LiteralPath $BuildEvidencePath -Raw | ConvertFrom-Json
  if ($buildEvidence.status -ne "PASS") { throw "Latest build evidence is not PASS." }
  if (-not (Test-Path -LiteralPath $CliPath -PathType Leaf)) { throw "Fresh built CLI is missing: $CliPath" }

  $baselineSource = Get-Content -LiteralPath $AcceptedFixtureSource -Raw
  [System.IO.File]::WriteAllText($FixtureSource, $baselineSource, [System.Text.UTF8Encoding]::new($false))
  $version = Invoke-CapturedProcess -FilePath $CliPath -Arguments @("version") -WorkingDirectory $RepoRoot
  $warmup = Invoke-NormalAnalyze "warmup"
  $unchanged1 = Invoke-NormalAnalyze "unchanged-1"
  $unchanged2 = Invoke-NormalAnalyze "unchanged-2"

  $oldNowIDs = @($unchanged2.facts.definitions | Where-Object { $_.name -eq "now" } | ForEach-Object { $_.id })
  $oldSecondNowID = @($unchanged2.facts.definitions | Where-Object { $_.name -eq "now" -and $_.startLine -eq 20 } | ForEach-Object { $_.id })
  $changedBlock = @'
function secondReport() {
  const later = Date.now();
  return later;
}
'@
  $baselineBlock = @'
function secondReport() {
  const now = Date.now();
  return now;
}
'@
  $changedSource = $baselineSource.Replace($baselineBlock, $changedBlock)
  if ($changedSource -eq $baselineSource) { throw "Could not apply the controlled source change." }
  [System.IO.File]::WriteAllText($FixtureSource, $changedSource, [System.Text.UTF8Encoding]::new($false))
  $changed = Invoke-NormalAnalyze "changed-source"

  $variableQuery = "MATCH (n:Variable) WHERE n.filePath = 'src/identity.ts' RETURN n.id AS id, n.name AS name, n.filePath AS filePath, n.startLine AS startLine, n.startCol AS startCol, n.endLine AS endLine, n.endCol AS endCol, n.selectionStartLine AS selectionStartLine, n.selectionStartCol AS selectionStartCol, n.selectionEndLine AS selectionEndLine, n.selectionEndCol AS selectionEndCol ORDER BY n.startLine"
  $storage = Join-Path $FixtureRoot ".anvien"
  $graphPath = Join-Path $storage "graph.json"
  $lbugPath = Join-Path $storage "lbug"
  $graphHold = Join-Path $LaneRoot "held-native-proof-graph.json"
  $nativeRead = $null
  Move-Item -LiteralPath $graphPath -Destination $graphHold
  try {
    $nativeRead = Invoke-McpCypher "native-read-with-graph-json-absent" $variableQuery
    $nativeRead["graphJsonAbsentDuringRead"] = -not (Test-Path -LiteralPath $graphPath)
    $nativeRead["lbugPresentDuringRead"] = Test-Path -LiteralPath $lbugPath -PathType Leaf
    $nativeRead["lbugSha256"] = Get-Sha256 $lbugPath
  } finally {
    if (Test-Path -LiteralPath $graphHold) { Move-Item -LiteralPath $graphHold -Destination $graphPath -Force }
  }

  $storageHold = Join-Path $LaneRoot "held-storage-analyze-fault"
  if (Test-Path -LiteralPath $storageHold) { Remove-Item -LiteralPath $storageHold -Recurse -Force }
  Move-Item -LiteralPath $storage -Destination $storageHold
  [System.IO.File]::WriteAllText($storage, "P2-D intentional storage fault", [System.Text.UTF8Encoding]::new($false))
  $analyzeFailure = $null
  try {
    $failureProcess = Invoke-CapturedProcess -FilePath $CliPath -Arguments @("analyze", $FixtureRoot, "--force", "--json", "--name", $RegistryName) -WorkingDirectory $RepoRoot -Environment @{ ANVIEN_HOME = $IsolatedHome }
    $analyzeFailure = [ordered]@{
      label = "analyze-storage-path-is-file"
      injection = "moved lane-owned .anvien directory outside fixture and created a regular file at the exact storage path"
      process = $failureProcess
      expectedGraphPathPresentDuringFailure = Test-Path -LiteralPath (Join-Path $storage "graph.json")
      heldCurrentGraphSha256 = Get-Sha256 (Join-Path $storageHold "graph.json")
      heldCurrentLbugSha256 = Get-Sha256 (Join-Path $storageHold "lbug")
    }
  } finally {
    if (Test-Path -LiteralPath $storage -PathType Leaf) { Remove-Item -LiteralPath $storage -Force }
    if (Test-Path -LiteralPath $storageHold) { Move-Item -LiteralPath $storageHold -Destination $storage -Force }
  }

  $backendHold = Join-Path $LaneRoot "held-no-backend"
  if (Test-Path -LiteralPath $backendHold) { Remove-Item -LiteralPath $backendHold -Recurse -Force }
  New-Item -ItemType Directory -Path $backendHold -Force | Out-Null
  Move-Item -LiteralPath $graphPath -Destination (Join-Path $backendHold "graph.json")
  Move-Item -LiteralPath $lbugPath -Destination (Join-Path $backendHold "lbug")
  $noBackendRead = $null
  try {
    $noBackendRead = Invoke-McpCypher "no-readable-backend" $variableQuery
    $noBackendRead["graphJsonAbsentDuringRead"] = -not (Test-Path -LiteralPath $graphPath)
    $noBackendRead["lbugAbsentDuringRead"] = -not (Test-Path -LiteralPath $lbugPath)
    $noBackendRead["heldGraphSha256"] = Get-Sha256 (Join-Path $backendHold "graph.json")
    $noBackendRead["heldLbugSha256"] = Get-Sha256 (Join-Path $backendHold "lbug")
  } finally {
    if (Test-Path -LiteralPath (Join-Path $backendHold "graph.json")) { Move-Item -LiteralPath (Join-Path $backendHold "graph.json") -Destination $graphPath -Force }
    if (Test-Path -LiteralPath (Join-Path $backendHold "lbug")) { Move-Item -LiteralPath (Join-Path $backendHold "lbug") -Destination $lbugPath -Force }
    if (Test-Path -LiteralPath $backendHold) { Remove-Item -LiteralPath $backendHold -Force }
  }

  [System.IO.File]::WriteAllText($FixtureSource, $baselineSource, [System.Text.UTF8Encoding]::new($false))
  $recovery = Invoke-NormalAnalyze "recovery-baseline"
  $recoveryRead = Invoke-McpCypher "recovery-native-read" $variableQuery

  $nativeNames = Get-McpVariableNames $nativeRead
  $recoveryNames = Get-McpVariableNames $recoveryRead
  $unchangedInputStable = $unchanged1.inputBefore.signature -eq $unchanged1.inputAfter.signature -and $unchanged1.inputAfter.signature -eq $unchanged2.inputBefore.signature -and $unchanged2.inputBefore.signature -eq $unchanged2.inputAfter.signature
  $unchangedFactsStable = $unchanged1.facts.definitionSignature -eq $unchanged2.facts.definitionSignature
  $unchangedExpected = $unchanged1.facts.definitionCount -eq 10 -and $unchanged1.facts.definesCount -eq 10 -and $unchanged1.facts.missingEndpointCount -eq 0 -and $unchanged1.facts.nameCounts.time -eq 2 -and $unchanged1.facts.nameCounts.now -eq 2
  $changedOldMissing = @($oldSecondNowID | Where-Object { $changed.facts.definitions.id -contains $_ }).Count -eq 0
  $changedExpected = $changed.process.exitCode -eq 0 -and $changed.facts.nameCounts.time -eq 2 -and $changed.facts.nameCounts.now -eq 1 -and $changed.facts.nameCounts.later -eq 1 -and $changed.facts.missingEndpointCount -eq 0 -and $changedOldMissing
  $nativeHasError = $null -ne $nativeRead.response -and $null -ne $nativeRead.response.PSObject.Properties['error']
  $noBackendHasError = $null -ne $noBackendRead.response -and $null -ne $noBackendRead.response.PSObject.Properties['error']
  $nativeExpected = $nativeRead.process.exitCode -eq 0 -and -not $nativeHasError -and $nativeRead.graphJsonAbsentDuringRead -and $nativeRead.lbugPresentDuringRead -and (@($nativeNames | Where-Object { $_ -eq "time" }).Count -eq 2) -and (@($nativeNames | Where-Object { $_ -eq "now" }).Count -eq 1) -and (@($nativeNames | Where-Object { $_ -eq "later" }).Count -eq 1)
  $analyzeFailureExpected = $analyzeFailure.process.exitCode -ne 0 -and -not $analyzeFailure.expectedGraphPathPresentDuringFailure -and [bool]$analyzeFailure.process.stderr
  $noBackendExpected = $noBackendRead.process.exitCode -eq 0 -and $noBackendHasError -and $noBackendRead.graphJsonAbsentDuringRead -and $noBackendRead.lbugAbsentDuringRead
  $recoveryExpected = $recovery.process.exitCode -eq 0 -and $recovery.facts.definitionSignature -eq $unchanged1.facts.definitionSignature -and $recovery.facts.missingEndpointCount -eq 0 -and (@($recoveryNames | Where-Object { $_ -eq "time" }).Count -eq 2) -and (@($recoveryNames | Where-Object { $_ -eq "now" }).Count -eq 2) -and (@($recoveryNames | Where-Object { $_ -eq "later" }).Count -eq 0)

  $matrix = @(
    [ordered]@{ id = "M1"; scenario = "two unchanged normal built analyze runs"; verdict = $(if ($unchanged1.process.exitCode -eq 0 -and $unchanged2.process.exitCode -eq 0 -and $unchangedInputStable -and $unchangedFactsStable -and $unchangedExpected) { "PASS" } else { "FAIL" }); proves = "same path + equal input manifests preserve accepted fact identity/ranges/endpoints; whole-artifact hashes are informational" },
    [ordered]@{ id = "M2"; scenario = "changed source then current artifact/read"; verdict = $(if ($changedExpected -and $nativeExpected) { "PASS" } else { "FAIL" }); proves = "next successful artifact/read includes later, reduces now to one, and excludes the prior second-now identity" },
    [ordered]@{ id = "M3"; scenario = "clear analyze-command failure"; verdict = $(if ($analyzeFailureExpected) { "PASS" } else { "FAIL" }); proves = "owned storage failure returns nonzero and no graph at the normal expected artifact path" },
    [ordered]@{ id = "M4"; scenario = "C17 native availability with Graph JSON absent"; verdict = $(if ($nativeExpected) { "PASS" } else { "FAIL" }); proves = "successful current read used Ladybug, because graph.json was absent for the full MCP invocation" },
    [ordered]@{ id = "M5"; scenario = "C17 no-readable-backend failure"; verdict = $(if ($noBackendExpected) { "PASS" } else { "FAIL" }); proves = "native open failure is not ErrUnavailable fallback; with both backends absent MCP returns JSON-RPC non-success and no stale result" },
    [ordered]@{ id = "M6"; scenario = "subsequent successful analyze/read recovery"; verdict = $(if ($recoveryExpected) { "PASS" } else { "FAIL" }); proves = "restored baseline input produces the original fact signature and readable current rows after both faults" },
    [ordered]@{ id = "M7"; scenario = "accepted corrected-fact semantics"; verdict = $(if ($unchangedExpected -and $changedExpected -and $recoveryExpected) { "PASS" } else { "FAIL" }); proves = "exact IDs, construct/selection ranges, 10/10 DEFINES, and zero missing endpoints were compared rather than aggregate totals alone" }
  )
  $overall = if (@($matrix | Where-Object { $_.verdict -ne "PASS" }).Count -eq 0) { "PASS" } else { "FAIL" }
  $timestamp = Get-Date -Format "yyMMdd_HHmmss"
  New-Item -ItemType Directory -Path $OfficialEvidenceRoot -Force | Out-Null
  $jsonPath = Join-Path $OfficialEvidenceRoot "qa_child02_p2d_repeated_analyze_$timestamp.json"
  $markdownPath = Join-Path $OfficialEvidenceRoot "qa_child02_p2d_repeated_analyze_$timestamp.md"
  $evidence = [ordered]@{
    schema = "child02-p2d-repeated-analyze-v1"
    generatedUtc = [DateTime]::UtcNow.ToString("o")
    overall = $overall
    repo = [ordered]@{ root = $RepoRoot; fixture = $FixtureRoot; sameFixturePath = $true; isolatedAnvienHome = $IsolatedHome }
    builtRuntime = [ordered]@{ path = $CliPath; sha256 = Get-Sha256 $CliPath; bytes = (Get-Item -LiteralPath $CliPath).Length; version = $version; buildEvidence = $buildEvidence }
    sourceClassification = [ordered]@{
      c04 = "Run writes Ladybug before graph snapshot; graph snapshot uses temp+rename; every returned error prevents CLI registration/success output"
      c17 = "Ladybug primary; only errors.Is(openErr, ErrUnavailable) authorizes same-repository graph.json fallback; other open/query errors return non-success"
      normalTaggedBinary = "fresh packaged runtime metadata and build source use ladybugdb tag; ErrUnavailable fallback cannot be induced by deleting/corrupting a native artifact because those are non-sentinel native errors"
      supportedFallback = "source-authorized only for the exact ErrUnavailable sentinel; not executed as a fake normal-binary success in this run"
    }
    acceptedFixture = [ordered]@{ source = $AcceptedFixtureSource; sha256 = Get-Sha256 $AcceptedFixtureSource; expectedDefinitions = 10; expectedTime = 2; expectedNow = 2 }
    matrix = $matrix
    invocations = [ordered]@{ warmup = $warmup; unchanged1 = $unchanged1; unchanged2 = $unchanged2; changed = $changed; nativeRead = $nativeRead; analyzeFailure = $analyzeFailure; noBackendRead = $noBackendRead; recovery = $recovery; recoveryRead = $recoveryRead }
    comparisons = [ordered]@{
      unchangedInputStable = $unchangedInputStable
      unchangedFactsStable = $unchangedFactsStable
      graphJsonWholeArtifactHashes = @($unchanged1.artifacts.'graph.json'.sha256, $unchanged2.artifacts.'graph.json'.sha256)
      lbugWholeArtifactHashes = @($unchanged1.artifacts.lbug.sha256, $unchanged2.artifacts.lbug.sha256)
      wholeArtifactHashContractClaimed = $false
      oldNowIDs = $oldNowIDs
      oldSecondNowID = $oldSecondNowID
      changedOldSecondNowExcluded = $changedOldMissing
      nativeReadNames = $nativeNames
      recoveryReadNames = $recoveryNames
    }
    lifecycle = [ordered]@{
      faultArtifactsRestored = (Test-Path -LiteralPath $graphPath -PathType Leaf) -and (Test-Path -LiteralPath $lbugPath -PathType Leaf)
      heldPathsRemaining = @($graphHold, $storageHold, $backendHold | Where-Object { Test-Path -LiteralPath $_ })
      laneOwnedRuntimeProcessesExpected = 0
    }
  }
  Write-JsonFile $jsonPath $evidence

  $lines = @(
    "# Child 02 P2-D repeated analyze evidence",
    "",
    "- Overall: **$overall**",
    "- Built CLI: ``$CliPath``",
    "- Built CLI SHA-256: ``$(Get-Sha256 $CliPath)``",
    "- Same isolated repository path: ``$FixtureRoot``",
    "- Isolated ANVIEN_HOME: ``$IsolatedHome``",
    "- Accepted fixture SHA-256: ``$(Get-Sha256 $AcceptedFixtureSource)``",
    "",
    "## Matrix",
    "",
    "| ID | Scenario | Verdict | What it proves |",
    "|---|---|---|---|"
  )
  foreach ($row in $matrix) { $lines += "| $($row.id) | $($row.scenario) | $($row.verdict) | $($row.proves) |" }
  $lines += @(
    "",
    "## Artifact and backend classification",
    "",
    "- C04: Ladybug load completes before atomic Graph JSON temp+rename; errors prevent normal CLI registration/success output.",
    "- C17: Ladybug is primary. Only the exact ``ErrUnavailable`` sentinel permits fallback to the same repository's ``graph.json``. Other native open/query errors fail clearly.",
    "- Native proof removed ``graph.json`` for the entire MCP call; current changed facts were still returned, proving Ladybug use.",
    "- No-backend proof removed both ``lbug`` and ``graph.json``; MCP returned a JSON-RPC error and no stale/substitute rows.",
    "- The freshly built normal binary includes the ``ladybugdb`` tag. The supported fallback is classified from source but was not fabricated through a non-normal/dev build.",
    "- Whole Graph JSON/Ladybug hashes are recorded only as artifact identity. Determinism verdict uses the accepted canonical fact signature, ranges, identities, and endpoints.",
    "",
    "## Exact corrected facts",
    "",
    "- Unchanged runs: definitions ``$($unchanged1.facts.definitionCount)/$($unchanged2.facts.definitionCount)``, DEFINES ``$($unchanged1.facts.definesCount)/$($unchanged2.facts.definesCount)``, missing endpoints ``$($unchanged1.facts.missingEndpointCount)/$($unchanged2.facts.missingEndpointCount)``.",
    "- Unchanged fact signature: ``$($unchanged1.facts.definitionSignature)``.",
    "- Changed run: time ``$($changed.facts.nameCounts.time)``, now ``$($changed.facts.nameCounts.now)``, later ``$($changed.facts.nameCounts.later)``, old second-now excluded ``$changedOldMissing``.",
    "- Recovery fact signature equals baseline: ``$($recovery.facts.definitionSignature -eq $unchanged1.facts.definitionSignature)``.",
    "",
    "Full command output, exit codes, timings, peak working sets, artifact hashes/timestamps, fact rows, reads, fault injection/restoration, and clean-holder/full-build evidence are in the paired JSON."
  )
  [System.IO.File]::WriteAllLines($markdownPath, $lines, [System.Text.UTF8Encoding]::new($false))
  $result = [ordered]@{ overall = $overall; json = $jsonPath; jsonSha256 = Get-Sha256 $jsonPath; markdown = $markdownPath; markdownSha256 = Get-Sha256 $markdownPath; matrix = $matrix }
  Write-JsonFile (Join-Path $LaneRoot "latest-evidence.json") $result
  $result | ConvertTo-Json -Depth 20
  if ($overall -ne "PASS") { exit 1 }
}

switch ($Mode) {
  "Prepare" { Invoke-Prepare }
  "Build" { Invoke-Build }
  "Run" { Invoke-Run }
}
