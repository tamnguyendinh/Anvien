[CmdletBinding()]
param(
  [Parameter(Mandatory = $true)]
  [ValidateSet("Prepare", "Build", "Repeat", "Readers", "Parity", "RuntimeStart", "RuntimeStatus", "RuntimeStop", "Cleanup")]
  [string]$Mode
)

$ErrorActionPreference = "Stop"
Set-StrictMode -Version Latest

$RepoRoot = Split-Path -Parent $PSScriptRoot
$LaneRoot = Join-Path $RepoRoot ".tmp\qa-child02-p2e"
$OwnerMarker = Join-Path $LaneRoot "owner.json"
$AcceptedRepeatHarness = Join-Path $RepoRoot "scripts\qa-child02-p2d-repeated-analyze.ps1"
$AcceptedRepeatHarnessSha256 = "2CF8A230154029451343333CB206D0C1444D90FA988460EDB719CE34980733C7"
$ExecutionRepeatHarness = Join-Path $LaneRoot "qa-child02-p2e-repeated-analyze.ps1"
$BuildEvidenceRoot = Join-Path $RepoRoot "reports\QA\child02-p2e-build"
$ReaderEvidenceRoot = Join-Path $RepoRoot "reports\QA\child02-p2e-readers"
$ParityEvidenceRoot = Join-Path $RepoRoot "reports\QA\child02-p2e-parity"
$LadybugRuntime = Join-Path $RepoRoot "third_party\ladybugdb\v0.19.1\windows-x86_64"
$RuntimeManifestPath = Join-Path $LaneRoot "runtime.json"

function Write-JsonFile([string]$Path, [object]$Value) {
  $parent = Split-Path -Parent $Path
  if ($parent) { New-Item -ItemType Directory -Path $parent -Force | Out-Null }
  $json = $Value | ConvertTo-Json -Depth 100
  [System.IO.File]::WriteAllText($Path, $json + [Environment]::NewLine, [System.Text.UTF8Encoding]::new($false))
}

function Write-TextFile([string]$Path, [string[]]$Lines) {
  $parent = Split-Path -Parent $Path
  if ($parent) { New-Item -ItemType Directory -Path $parent -Force | Out-Null }
  [System.IO.File]::WriteAllLines($Path, $Lines, [System.Text.UTF8Encoding]::new($false))
}

function Get-Sha256([string]$Path) {
  if (-not (Test-Path -LiteralPath $Path -PathType Leaf)) { return $null }
  return (Get-FileHash -LiteralPath $Path -Algorithm SHA256).Hash
}

function Format-Command([string]$FilePath, [string[]]$Arguments) {
  $parts = @($FilePath)
  foreach ($argument in $Arguments) {
    $parts += if ($argument -match '[\s"]') { '"' + $argument.Replace('"', '\"') + '"' } else { $argument }
  }
  return $parts -join " "
}

function Invoke-CapturedProcess {
  param(
    [Parameter(Mandatory = $true)][string]$FilePath,
    [string[]]$Arguments = @(),
    [Parameter(Mandatory = $true)][string]$WorkingDirectory,
    [hashtable]$Environment = @{}
  )
  $startUtc = [DateTime]::UtcNow
  $stopwatch = [System.Diagnostics.Stopwatch]::StartNew()
  $startInfo = [System.Diagnostics.ProcessStartInfo]::new()
  $startInfo.FileName = $FilePath
  $startInfo.WorkingDirectory = $WorkingDirectory
  $startInfo.UseShellExecute = $false
  $startInfo.RedirectStandardOutput = $true
  $startInfo.RedirectStandardError = $true
  $startInfo.CreateNoWindow = $true
  $startInfo.Arguments = (@($Arguments | ForEach-Object {
    if ($_ -match '[\s"]') { '"' + $_.Replace('"', '\"') + '"' } else { $_ }
  }) -join " ")
  foreach ($key in $Environment.Keys) { $startInfo.Environment[$key] = [string]$Environment[$key] }
  $process = [System.Diagnostics.Process]::new()
  $process.StartInfo = $startInfo
  if (-not $process.Start()) { throw "Failed to start $(Format-Command $FilePath $Arguments)" }
  $stdoutTask = $process.StandardOutput.ReadToEndAsync()
  $stderrTask = $process.StandardError.ReadToEndAsync()
  $process.WaitForExit()
  $stdout = $stdoutTask.GetAwaiter().GetResult()
  $stderr = $stderrTask.GetAwaiter().GetResult()
  $stopwatch.Stop()
  $peak = 0
  try { $peak = $process.PeakWorkingSet64 } catch { }
  return [ordered]@{
    command = Format-Command $FilePath $Arguments
    filePath = $FilePath
    arguments = @($Arguments)
    workingDirectory = $WorkingDirectory
    startUtc = $startUtc.ToString("o")
    endUtc = [DateTime]::UtcNow.ToString("o")
    durationSeconds = [Math]::Round($stopwatch.Elapsed.TotalSeconds, 3)
    exitCode = $process.ExitCode
    peakWorkingSetBytes = $peak
    stdout = $stdout.TrimEnd()
    stderr = $stderr.TrimEnd()
  }
}

function Assert-LaneOwnership {
  if (-not (Test-Path -LiteralPath $OwnerMarker -PathType Leaf)) { throw "Missing P2-E owner marker: $OwnerMarker" }
  $owner = Get-Content -LiteralPath $OwnerMarker -Raw | ConvertFrom-Json
  if ($owner.scope -ne "child02-p2e" -or $owner.owner -ne "independent-qa") {
    throw "Refusing to use non-P2-E temp root: $LaneRoot"
  }
}

function New-ExecutionRepeatHarness {
  $actual = Get-Sha256 $AcceptedRepeatHarness
  if ($actual -ne $AcceptedRepeatHarnessSha256) {
    throw "Accepted P2-D harness hash drifted: expected $AcceptedRepeatHarnessSha256, got $actual"
  }
  $source = Get-Content -LiteralPath $AcceptedRepeatHarness -Raw
  $source = $source.Replace('$RepoRoot = Split-Path -Parent $PSScriptRoot', '$RepoRoot = "' + $RepoRoot.Replace('"', '`"') + '"')
  $source = $source.Replace("qa-child02-p2d", "qa-child02-p2e")
  $source = $source.Replace("child02-p2d", "child02-p2e")
  $source = $source.Replace("P2DRestartManager", "P2ERestartManager")
  $source = $source.Replace("P2-D", "P2-E")
  $source = $source.Replace("p2d-qa@local.invalid", "p2e-qa@local.invalid")
  $source = $source.Replace("child02-p2d-repeated-analyze-v1", "child02-p2e-repeated-analyze-v1")
  [System.IO.File]::WriteAllText($ExecutionRepeatHarness, $source, [System.Text.UTF8Encoding]::new($false))
}

function Invoke-Prepare {
  if (Test-Path -LiteralPath $LaneRoot) {
    Assert-LaneOwnership
  } else {
    New-Item -ItemType Directory -Path $LaneRoot -Force | Out-Null
    Write-JsonFile $OwnerMarker ([ordered]@{
      scope = "child02-p2e"
      owner = "independent-qa"
      createdUtc = [DateTime]::UtcNow.ToString("o")
      authorizedBoundary = "Child 02 P2-E validation-only"
    })
  }
  New-ExecutionRepeatHarness
  $powershell = (Get-Command powershell.exe -ErrorAction Stop).Source
  $result = Invoke-CapturedProcess -FilePath $powershell -Arguments @("-NoProfile", "-ExecutionPolicy", "Bypass", "-File", $ExecutionRepeatHarness, "-Mode", "Prepare") -WorkingDirectory $RepoRoot
  if ($result.exitCode -ne 0) { throw "P2-E repeat harness prepare failed:`n$($result.stderr)" }
  $prepared = [ordered]@{
    schema = "child02-p2e-prepare-v1"
    generatedUtc = [DateTime]::UtcNow.ToString("o")
    acceptedHarness = [ordered]@{ path = $AcceptedRepeatHarness; expectedSha256 = $AcceptedRepeatHarnessSha256; actualSha256 = Get-Sha256 $AcceptedRepeatHarness }
    executionHarness = [ordered]@{ path = $ExecutionRepeatHarness; sha256 = Get-Sha256 $ExecutionRepeatHarness }
    command = $result
  }
  Write-JsonFile (Join-Path $LaneRoot "p2e-prepare.json") $prepared
  $prepared | ConvertTo-Json -Depth 100
}

function Invoke-Build {
  Assert-LaneOwnership
  New-ExecutionRepeatHarness
  $powershell = (Get-Command powershell.exe -ErrorAction Stop).Source
  $command = Invoke-CapturedProcess -FilePath $powershell -Arguments @("-NoProfile", "-ExecutionPolicy", "Bypass", "-File", $ExecutionRepeatHarness, "-Mode", "Build") -WorkingDirectory $RepoRoot
  $innerPath = Join-Path $LaneRoot "build-evidence.json"
  $inner = if (Test-Path -LiteralPath $innerPath -PathType Leaf) { Get-Content -LiteralPath $innerPath -Raw | ConvertFrom-Json } else { $null }
  $head = (& git rev-parse HEAD).Trim()
  $branch = (& git branch --show-current).Trim()
  $timestamp = Get-Date -Format "yyMMdd_HHmmss"
  New-Item -ItemType Directory -Path $BuildEvidenceRoot -Force | Out-Null
  $jsonPath = Join-Path $BuildEvidenceRoot "qa_child02_p2e_build_$timestamp.json"
  $markdownPath = Join-Path $BuildEvidenceRoot "qa_child02_p2e_build_$timestamp.md"
  $status = if ($command.exitCode -eq 0 -and $null -ne $inner -and $inner.status -eq "PASS") { "PASS" } else { "FAIL" }
  $evidence = [ordered]@{
    schema = "child02-p2e-clean-holder-build-v1"
    generatedUtc = [DateTime]::UtcNow.ToString("o")
    status = $status
    repo = [ordered]@{ root = $RepoRoot; branch = $branch; head = $head }
    acceptedHarness = [ordered]@{ path = $AcceptedRepeatHarness; sha256 = Get-Sha256 $AcceptedRepeatHarness }
    executionHarness = [ordered]@{ path = $ExecutionRepeatHarness; sha256 = Get-Sha256 $ExecutionRepeatHarness }
    invocation = $command
    cleanHolderAndBuild = $inner
  }
  Write-JsonFile $jsonPath $evidence
  Write-TextFile $markdownPath @(
    "# Child 02 P2-E clean-holder full build",
    "",
    "- Status: **$status**",
    "- Generated UTC: ``$($evidence.generatedUtc)``",
    "- Repository: ``$RepoRoot``",
    "- Branch / HEAD: ``$branch`` / ``$head``",
    "- Command: ``$($command.command)``",
    "- Exit / duration: ``$($command.exitCode)`` / ``$($command.durationSeconds)s``",
    "- Restart Manager zero: ``$($inner.holderGate.restartManagerZero)``",
    "- Exclusive open all: ``$($inner.holderGate.exclusiveOpenAll)``",
    "",
    "Full artifact paths, holder PIDs, launcher classification, killed processes, exclusive-open rows, command output, timings, and hashes are in the paired JSON."
  )
  [ordered]@{ status = $status; json = $jsonPath; jsonSha256 = Get-Sha256 $jsonPath; markdown = $markdownPath; markdownSha256 = Get-Sha256 $markdownPath } | ConvertTo-Json -Depth 20
  if ($status -ne "PASS") { exit 1 }
}

function Invoke-Repeat {
  Assert-LaneOwnership
  New-ExecutionRepeatHarness
  $powershell = (Get-Command powershell.exe -ErrorAction Stop).Source
  $result = Invoke-CapturedProcess -FilePath $powershell -Arguments @("-NoProfile", "-ExecutionPolicy", "Bypass", "-File", $ExecutionRepeatHarness, "-Mode", "Run") -WorkingDirectory $RepoRoot
  $result | ConvertTo-Json -Depth 100
  if ($result.exitCode -ne 0) { exit $result.exitCode }
}

function Get-NativeEnvironment {
  foreach ($required in @("lbug.h", "lbug_shared.lib", "lbug_shared.dll")) {
    if (-not (Test-Path -LiteralPath (Join-Path $LadybugRuntime $required) -PathType Leaf)) {
      throw "Missing repository-local Ladybug prerequisite: $(Join-Path $LadybugRuntime $required)"
    }
  }
  return @{
    CGO_ENABLED = "1"
    CGO_CFLAGS = "-I$($LadybugRuntime.Replace('\', '/'))"
    CGO_LDFLAGS = "-L$($LadybugRuntime.Replace('\', '/')) -llbug_shared"
    PATH = "$LadybugRuntime;$env:PATH"
  }
}

function Invoke-Readers {
  Assert-LaneOwnership
  $nativeEnvironment = Get-NativeEnvironment
  $vitest = Join-Path $RepoRoot "anvien-web\node_modules\.bin\vitest.cmd"
  $go = (Get-Command go.exe -ErrorAction Stop).Source
  $commands = @()
  $commands += Invoke-CapturedProcess -FilePath $vitest -Arguments @("run", "test/unit/useAppState.local-runtime.test.tsx", "test/unit/ChatPanel.grounding-links.test.tsx", "test/unit/CodeReferencesPanel.graph-health.test.tsx") -WorkingDirectory (Join-Path $RepoRoot "anvien-web")
  $commands += Invoke-CapturedProcess -FilePath $go -Arguments @("test", "./internal/filecontext", "-count=1", "-v") -WorkingDirectory $RepoRoot
  $commands += Invoke-CapturedProcess -FilePath $go -Arguments @("test", "./internal/mcp", "-count=1", "-v") -WorkingDirectory $RepoRoot
  $commands += Invoke-CapturedProcess -FilePath $go -Arguments @("test", "./internal/embeddings", "./internal/httpapi", "-count=1", "-v") -WorkingDirectory $RepoRoot
  $commands += Invoke-CapturedProcess -FilePath $go -Arguments @("test", "-tags", "ladybugdb", "./internal/embeddings", "-run", "^TestNativeLadybugSemanticSearchHydratesPersistedExplicitLabels$", "-count=1", "-v") -WorkingDirectory $RepoRoot -Environment $nativeEnvironment
  $commands += Invoke-CapturedProcess -FilePath $go -Arguments @("test", "-tags", "ladybugdb", "./internal/lbugnative", "-run", "^TestNativeLadybugPersistenceReadbackAndStream$", "-count=1", "-v") -WorkingDirectory $RepoRoot -Environment $nativeEnvironment
  $allPass = @($commands | Where-Object { $_.exitCode -ne 0 }).Count -eq 0
  $timestamp = Get-Date -Format "yyMMdd_HHmmss"
  New-Item -ItemType Directory -Path $ReaderEvidenceRoot -Force | Out-Null
  $jsonPath = Join-Path $ReaderEvidenceRoot "qa_child02_p2e_readers_$timestamp.json"
  $markdownPath = Join-Path $ReaderEvidenceRoot "qa_child02_p2e_readers_$timestamp.md"
  $rows = @(
    [ordered]@{ id = "C09"; scenario = "exact opaque-ID resolution"; boundary = "built frontend unit + mounted Playwright"; commandIndexes = @(0); verdict = $(if ($commands[0].exitCode -eq 0) { "PASS_PENDING_FRESH_UI" } else { "FAIL" }) },
    [ordered]@{ id = "C10"; scenario = "unique grounding fail-closed"; boundary = "built frontend unit + mounted Playwright"; commandIndexes = @(0); verdict = $(if ($commands[0].exitCode -eq 0) { "PASS_PENDING_FRESH_UI" } else { "FAIL" }) },
    [ordered]@{ id = "C11"; scenario = "one-based lines / zero-based UTF-8 byte columns / exclusive ends"; boundary = "built frontend unit + mounted Playwright"; commandIndexes = @(0); verdict = $(if ($commands[0].exitCode -eq 0) { "PASS_PENDING_FRESH_UI" } else { "FAIL" }) },
    [ordered]@{ id = "C12"; scenario = "nodeRange"; boundary = "internal/filecontext"; commandIndexes = @(1); verdict = $(if ($commands[1].exitCode -eq 0) { "PASS" } else { "FAIL" }) },
    [ordered]@{ id = "C13"; scenario = "Definition MCP context"; boundary = "internal/mcp"; commandIndexes = @(2); verdict = $(if ($commands[2].exitCode -eq 0) { "PASS" } else { "FAIL" }) },
    [ordered]@{ id = "C14"; scenario = "detectChangedSymbols"; boundary = "internal/mcp"; commandIndexes = @(2); verdict = $(if ($commands[2].exitCode -eq 0) { "PASS" } else { "FAIL" }) },
    [ordered]@{ id = "C15"; scenario = "collectRenameChanges"; boundary = "internal/mcp"; commandIndexes = @(2); verdict = $(if ($commands[2].exitCode -eq 0) { "PASS" } else { "FAIL" }) },
    [ordered]@{ id = "C16"; scenario = "persisted explicit embedding label and semantic-search hydration"; boundary = "embeddings + HTTP + repo-local native Ladybug"; commandIndexes = @(3,4,5); verdict = $(if ($commands[3].exitCode -eq 0 -and $commands[4].exitCode -eq 0 -and $commands[5].exitCode -eq 0) { "PASS" } else { "FAIL" }); vectorPolicy = "No remote VECTOR install/retry/substitution; repository-local persistence/readback/hydration seam only" }
  )
  $evidence = [ordered]@{
    schema = "child02-p2e-readers-v1"
    generatedUtc = [DateTime]::UtcNow.ToString("o")
    status = $(if ($allPass) { "PASS_PENDING_FRESH_UI" } else { "FAIL" })
    denominator = [ordered]@{ passedBackendOrUnit = @($rows | Where-Object { $_.verdict -in @("PASS", "PASS_PENDING_FRESH_UI") }).Count; total = 8; freshMountedUIRequired = @("C09", "C10", "C11") }
    nativeRuntime = [ordered]@{ path = $LadybugRuntime; source = "repository-local existing prerequisite"; remoteVectorAttempted = $false }
    rows = $rows
    commands = $commands
  }
  Write-JsonFile $jsonPath $evidence
  $rowLines = @($rows | ForEach-Object { "| $($_.id) | $($_.scenario) | $($_.boundary) | $($_.verdict) |" })
  Write-TextFile $markdownPath (@(
    "# Child 02 P2-E affected-reader command evidence", "", "- Status: **$($evidence.status)**", "- Denominator: ``8``", "- C09-C11 require the separate fresh mounted visible Playwright evidence before final 8/8 disposition.", "- Remote VECTOR attempts: ``0``", "", "| Row | Scenario | Boundary | Result |", "|---|---|---|---|"
  ) + $rowLines + @("", "Full commands, exits, UTC intervals, stdout, stderr, durations, and peak working sets are in the paired JSON."))
  [ordered]@{ status = $evidence.status; json = $jsonPath; jsonSha256 = Get-Sha256 $jsonPath; markdown = $markdownPath; markdownSha256 = Get-Sha256 $markdownPath } | ConvertTo-Json -Depth 20
  if (-not $allPass) { exit 1 }
}

function Invoke-Parity {
  Assert-LaneOwnership
  $nativeEnvironment = Get-NativeEnvironment
  $go = (Get-Command go.exe -ErrorAction Stop).Source
  $graphPath = Join-Path $RepoRoot ".anvien\graph.json"
  $ladybugPath = Join-Path $RepoRoot ".anvien\lbug"
  $utility = Join-Path $RepoRoot "scripts\qa-child02-p2e-parity.go"
  foreach ($required in @($graphPath, $ladybugPath, $utility)) {
    if (-not (Test-Path -LiteralPath $required -PathType Leaf)) { throw "Missing parity input: $required" }
  }
  $result = Invoke-CapturedProcess -FilePath $go -Arguments @("run", "-tags", "ladybugdb", $utility, "-graph", $graphPath, "-ladybug", $ladybugPath, "-out", $ParityEvidenceRoot) -WorkingDirectory $RepoRoot -Environment $nativeEnvironment
  $result | ConvertTo-Json -Depth 100
  if ($result.exitCode -ne 0) { exit $result.exitCode }
}

function Get-ProcessRecord([int]$ProcessID) {
  $process = Get-CimInstance Win32_Process -Filter "ProcessId = $ProcessID" -ErrorAction SilentlyContinue
  if ($null -eq $process) { return $null }
  return [ordered]@{ pid = [int]$process.ProcessId; parentPid = [int]$process.ParentProcessId; name = [string]$process.Name; executablePath = [string]$process.ExecutablePath; commandLine = [string]$process.CommandLine }
}

function Get-RuntimeStatus {
  $listeners = @(Get-NetTCPConnection -State Listen -ErrorAction SilentlyContinue | Where-Object { $_.LocalPort -in @(4848, 5228) } | ForEach-Object {
    [ordered]@{ port = $_.LocalPort; pid = $_.OwningProcess; process = Get-ProcessRecord $_.OwningProcess }
  })
  $backendHealth = $null
  $frontendHealth = $null
  try { $response = Invoke-WebRequest -UseBasicParsing -Uri "http://127.0.0.1:4848/api/info" -TimeoutSec 2; $backendHealth = [ordered]@{ statusCode = $response.StatusCode; content = $response.Content } } catch { $backendHealth = [ordered]@{ error = $_.Exception.Message } }
  try { $response = Invoke-WebRequest -UseBasicParsing -Uri "http://127.0.0.1:5228/" -TimeoutSec 2; $frontendHealth = [ordered]@{ statusCode = $response.StatusCode; contentLength = $response.Content.Length } } catch { $frontendHealth = [ordered]@{ error = $_.Exception.Message } }
  return [ordered]@{ generatedUtc = [DateTime]::UtcNow.ToString("o"); listeners = $listeners; backendHealth = $backendHealth; frontendHealth = $frontendHealth }
}

function Invoke-RuntimeStart {
  Assert-LaneOwnership
  $before = Get-RuntimeStatus
  if (@($before.listeners).Count -ne 0) { throw "Ports 4848/5228 already have listeners; refusing to start P2-E runtime." }
  $backendPath = Join-Path $RepoRoot "anvien\bin\anvien.exe"
  $nodePath = (Get-Command node.exe -ErrorAction Stop).Source
  $vitePath = Join-Path $RepoRoot "anvien-web\node_modules\vite\bin\vite.js"
  foreach ($required in @($backendPath, $nodePath, $vitePath)) { if (-not (Test-Path -LiteralPath $required -PathType Leaf)) { throw "Missing runtime artifact: $required" } }
  $backendOut = Join-Path $LaneRoot "backend.stdout.log"
  $backendErr = Join-Path $LaneRoot "backend.stderr.log"
  $frontendOut = Join-Path $LaneRoot "frontend.stdout.log"
  $frontendErr = Join-Path $LaneRoot "frontend.stderr.log"
  $backend = Start-Process -FilePath $backendPath -ArgumentList @("serve", "--host", "127.0.0.1", "--port", "4848") -WorkingDirectory $RepoRoot -WindowStyle Hidden -RedirectStandardOutput $backendOut -RedirectStandardError $backendErr -PassThru
  $frontend = Start-Process -FilePath $nodePath -ArgumentList @($vitePath, "preview", "--host", "127.0.0.1", "--port", "5228", "--strictPort") -WorkingDirectory (Join-Path $RepoRoot "anvien-web") -WindowStyle Hidden -RedirectStandardOutput $frontendOut -RedirectStandardError $frontendErr -PassThru
  $deadline = (Get-Date).AddSeconds(30)
  do {
    Start-Sleep -Milliseconds 500
    $status = Get-RuntimeStatus
    if ($status.backendHealth.statusCode -eq 200 -and $status.frontendHealth.statusCode -eq 200) { break }
  } while ((Get-Date) -lt $deadline)
  $manifest = [ordered]@{
    schema = "child02-p2e-mounted-runtime-v1"
    generatedUtc = [DateTime]::UtcNow.ToString("o")
    backend = [ordered]@{ pid = $backend.Id; executable = $backendPath; arguments = @("serve", "--host", "127.0.0.1", "--port", "4848"); stdout = $backendOut; stderr = $backendErr }
    frontend = [ordered]@{ pid = $frontend.Id; executable = $nodePath; arguments = @($vitePath, "preview", "--host", "127.0.0.1", "--port", "5228", "--strictPort"); stdout = $frontendOut; stderr = $frontendErr }
    status = $status
  }
  Write-JsonFile $RuntimeManifestPath $manifest
  $manifest | ConvertTo-Json -Depth 20
  if ($status.backendHealth.statusCode -ne 200 -or $status.frontendHealth.statusCode -ne 200) { exit 1 }
}

function Invoke-RuntimeStatus {
  Assert-LaneOwnership
  $manifest = if (Test-Path -LiteralPath $RuntimeManifestPath -PathType Leaf) { Get-Content -LiteralPath $RuntimeManifestPath -Raw | ConvertFrom-Json } else { $null }
  [ordered]@{ manifest = $manifest; current = Get-RuntimeStatus } | ConvertTo-Json -Depth 30
}

function Invoke-RuntimeStop {
  Assert-LaneOwnership
  if (-not (Test-Path -LiteralPath $RuntimeManifestPath -PathType Leaf)) { throw "P2-E runtime manifest is missing." }
  $manifest = Get-Content -LiteralPath $RuntimeManifestPath -Raw | ConvertFrom-Json
  $stopped = @()
  foreach ($owned in @($manifest.backend, $manifest.frontend)) {
    $ownedPid = [int]$owned.pid
    $record = Get-ProcessRecord $ownedPid
    if ($null -eq $record) { continue }
    if ([string]$record.executablePath -ne [string]$owned.executable) { throw "PID $ownedPid executable drifted; refusing to stop it." }
    Stop-Process -Id $ownedPid -Force -ErrorAction Stop
    $stopped += $record
  }
  Start-Sleep -Milliseconds 500
  [ordered]@{ generatedUtc = [DateTime]::UtcNow.ToString("o"); stopped = $stopped; final = Get-RuntimeStatus } | ConvertTo-Json -Depth 20
}

function Invoke-Cleanup {
  Assert-LaneOwnership
  Remove-Item -LiteralPath $LaneRoot -Recurse -Force
  [ordered]@{ removed = $LaneRoot; existsAfter = Test-Path -LiteralPath $LaneRoot; generatedUtc = [DateTime]::UtcNow.ToString("o") } | ConvertTo-Json
}

switch ($Mode) {
  "Prepare" { Invoke-Prepare }
  "Build" { Invoke-Build }
  "Repeat" { Invoke-Repeat }
  "Readers" { Invoke-Readers }
  "Parity" { Invoke-Parity }
  "RuntimeStart" { Invoke-RuntimeStart }
  "RuntimeStatus" { Invoke-RuntimeStatus }
  "RuntimeStop" { Invoke-RuntimeStop }
  "Cleanup" { Invoke-Cleanup }
}
