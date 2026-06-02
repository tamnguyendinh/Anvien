param(
  [Parameter(Mandatory = $true, Position = 0)]
  [ValidateSet("init", "prompt", "exec", "read", "response", "note", "summary")]
  [string]$Action,

  [string]$RunDir = ".tmp\agent-token-proxy",
  [string]$Phase = "",
  [string]$Text = "",
  [string]$TextFile = "",
  [string]$Path = "",
  [string]$Exec = "",

  [ValidateSet("powershell", "pwsh", "cmd")]
  [string]$Shell = "powershell",

  [string]$TokenizerEncoding = "o200k_base",
  [string]$TokenizerPath = ".tmp\tokenizer-python",
  [string]$TokenizerCommand = "",
  [switch]$DisableTokenizer,
  [switch]$StrictTokens,
  [switch]$Quiet
)

$ErrorActionPreference = "Stop"

$ScriptRoot = Split-Path -Parent $MyInvocation.MyCommand.Path
$RepoRoot = Split-Path -Parent $ScriptRoot

function Resolve-RepoPath($InputPath) {
  if ([string]::IsNullOrWhiteSpace($InputPath)) {
    return [System.IO.Path]::GetFullPath($RepoRoot)
  }
  if ([System.IO.Path]::IsPathRooted($InputPath)) {
    return [System.IO.Path]::GetFullPath($InputPath)
  }
  return [System.IO.Path]::GetFullPath((Join-Path $RepoRoot $InputPath))
}

function ConvertTo-Hex($Bytes) {
  $Builder = [System.Text.StringBuilder]::new()
  foreach ($Byte in $Bytes) {
    [void]$Builder.Append($Byte.ToString("x2"))
  }
  return $Builder.ToString()
}

function ConvertTo-CommandArgument($Value) {
  if ($null -eq $Value) {
    return '""'
  }
  return '"' + ([string]$Value).Replace('\', '\\').Replace('"', '\"') + '"'
}

function Get-TextSha256($Content) {
  if ($null -eq $Content) {
    $Content = ""
  }
  $Bytes = [System.Text.Encoding]::UTF8.GetBytes($Content)
  $Sha = [System.Security.Cryptography.SHA256]::Create()
  try {
    return ConvertTo-Hex ($Sha.ComputeHash($Bytes))
  } finally {
    $Sha.Dispose()
  }
}

function Invoke-Tokenizer($Content, $TokenizerCommand, $TokenizerEncoding, $TokenizerPath, $DisableTokenizer, $RunRoot) {
  if ($DisableTokenizer) {
    return [ordered]@{
      tokens = $null
      token_status = "tokenizer_disabled"
      tokenizer = ""
      tokenizer_error = ""
    }
  }

  $TempDir = Join-Path $RunRoot "tokenizer"
  New-Item -ItemType Directory -Path $TempDir -Force | Out-Null
  $TempPath = Join-Path $TempDir ([System.Guid]::NewGuid().ToString("n") + ".txt")
  [System.IO.File]::WriteAllText($TempPath, $Content, [System.Text.UTF8Encoding]::new($false))

  try {
    $Output = $null
    $Exit = 0
    $TokenizerLabel = $TokenizerCommand

    if (-not [string]::IsNullOrWhiteSpace($TokenizerCommand)) {
      $QuotedPath = "'" + $TempPath.Replace("'", "''") + "'"
      $Command = $TokenizerCommand.Replace("{file}", $QuotedPath)
      if ($Command -eq $TokenizerCommand) {
        $Command = "$TokenizerCommand $QuotedPath"
      }
      $Output = & powershell.exe -NoProfile -ExecutionPolicy Bypass -Command $Command 2>&1
      $Exit = if ($null -eq $LASTEXITCODE) { 0 } else { $LASTEXITCODE }
    } else {
      $TokenizerPathFull = Resolve-RepoPath $TokenizerPath
      $Python = @"
import os
import sys

tokenizer_path = r'''$TokenizerPathFull'''
if os.path.isdir(tokenizer_path):
    sys.path.insert(0, tokenizer_path)

import tiktoken

encoding = tiktoken.get_encoding(r'''$TokenizerEncoding''')
with open(r'''$TempPath''', 'r', encoding='utf-8') as handle:
    text = handle.read()
print(len(encoding.encode(text)))
"@
      $Output = $Python | python - 2>&1
      $Exit = if ($null -eq $LASTEXITCODE) { 0 } else { $LASTEXITCODE }
      $TokenizerLabel = "python:tiktoken:$TokenizerEncoding"
    }
    $LastLine = (($Output | Select-Object -Last 1) -join "").Trim()

    if ($Exit -eq 0 -and $LastLine -match "^\d+$") {
      return [ordered]@{
        tokens = [int64]$LastLine
        token_status = "exact"
        tokenizer = $TokenizerLabel
        tokenizer_error = ""
      }
    }

    return [ordered]@{
      tokens = $null
      token_status = "tokenizer_error"
      tokenizer = $TokenizerLabel
      tokenizer_error = (($Output | Out-String).Trim())
    }
  } finally {
    Remove-Item -LiteralPath $TempPath -Force -ErrorAction SilentlyContinue
  }
}

function Measure-Content($Content, $TokenizerCommand, $TokenizerEncoding, $TokenizerPath, $DisableTokenizer, $RunRoot) {
  if ($null -eq $Content) {
    $Content = ""
  }
  $Utf8Bytes = [System.Text.Encoding]::UTF8.GetByteCount($Content)
  $TokenResult = Invoke-Tokenizer $Content $TokenizerCommand $TokenizerEncoding $TokenizerPath $DisableTokenizer $RunRoot
  return [ordered]@{
    chars = $Content.Length
    utf8_bytes = $Utf8Bytes
    sha256 = Get-TextSha256 $Content
    tokens = $TokenResult.tokens
    token_status = $TokenResult.token_status
    tokenizer = $TokenResult.tokenizer
    tokenizer_error = $TokenResult.tokenizer_error
  }
}

function Add-TranscriptEvent($RunRoot, $Phase, $Bucket, $EventType, $Content, $Delivered, $TokenizerCommand, $TokenizerEncoding, $TokenizerPath, $DisableTokenizer, $Extra) {
  if ($null -eq $Content) {
    $Content = ""
  }
  $LogPath = Join-Path $RunRoot "transcript.ndjson"
  $Measure = Measure-Content $Content $TokenizerCommand $TokenizerEncoding $TokenizerPath $DisableTokenizer $RunRoot
  $Event = [ordered]@{
    schema_version = "agent_token_proxy_v1"
    timestamp_utc = [DateTime]::UtcNow.ToString("o")
    phase = $Phase
    bucket = $Bucket
    event_type = $EventType
    delivered_to_agent = [bool]$Delivered
    chars = $Measure.chars
    utf8_bytes = $Measure.utf8_bytes
    sha256 = $Measure.sha256
    tokens = $Measure.tokens
    token_status = $Measure.token_status
    tokenizer = $Measure.tokenizer
    tokenizer_error = $Measure.tokenizer_error
    content = $Content
  }

  if ($Extra) {
    foreach ($Key in $Extra.Keys) {
      $Event[$Key] = $Extra[$Key]
    }
  }

  ($Event | ConvertTo-Json -Depth 20 -Compress) | Add-Content -LiteralPath $LogPath -Encoding UTF8
}

function Get-InputText($Text, $TextFile) {
  if (-not [string]::IsNullOrWhiteSpace($TextFile)) {
    $FullPath = Resolve-RepoPath $TextFile
    return [System.IO.File]::ReadAllText($FullPath)
  }
  return $Text
}

function Invoke-CapturedCommand($Shell, $Exec) {
  if ([string]::IsNullOrWhiteSpace($Exec)) {
    throw "The exec action requires -Exec."
  }

  $ProcessInfo = [System.Diagnostics.ProcessStartInfo]::new()
  switch ($Shell) {
    "pwsh" {
      $ProcessInfo.FileName = "pwsh"
      $ProcessInfo.Arguments = "-NoProfile -Command " + (ConvertTo-CommandArgument $Exec)
    }
    "cmd" {
      $ProcessInfo.FileName = "cmd.exe"
      $ProcessInfo.Arguments = "/d /c " + (ConvertTo-CommandArgument $Exec)
    }
    default {
      $ProcessInfo.FileName = "powershell.exe"
      $ProcessInfo.Arguments = "-NoProfile -ExecutionPolicy Bypass -Command " + (ConvertTo-CommandArgument $Exec)
    }
  }

  $ProcessInfo.WorkingDirectory = $RepoRoot
  $ProcessInfo.RedirectStandardOutput = $true
  $ProcessInfo.RedirectStandardError = $true
  $ProcessInfo.UseShellExecute = $false
  $ProcessInfo.CreateNoWindow = $true

  $Stopwatch = [System.Diagnostics.Stopwatch]::StartNew()
  $Process = [System.Diagnostics.Process]::new()
  $Process.StartInfo = $ProcessInfo
  [void]$Process.Start()
  $Stdout = $Process.StandardOutput.ReadToEnd()
  $Stderr = $Process.StandardError.ReadToEnd()
  $Process.WaitForExit()
  $Stopwatch.Stop()

  return [ordered]@{
    stdout = $Stdout
    stderr = $Stderr
    exit_code = $Process.ExitCode
    elapsed_ms = $Stopwatch.ElapsedMilliseconds
  }
}

function Read-TranscriptEvents($LogPath) {
  if (-not (Test-Path -LiteralPath $LogPath)) {
    return @()
  }
  return Get-Content -LiteralPath $LogPath | Where-Object {
    -not [string]::IsNullOrWhiteSpace($_)
  } | ForEach-Object {
    $_ | ConvertFrom-Json
  }
}

function New-BucketSummary() {
  return [ordered]@{
    events = 0
    chars = 0
    utf8_bytes = 0
    tokens = 0
    token_status = "exact"
  }
}

function Write-Summary($RunRoot, $StrictTokens) {
  $LogPath = Join-Path $RunRoot "transcript.ndjson"
  $Events = @(Read-TranscriptEvents $LogPath)
  $Buckets = [ordered]@{}
  $Delivered = @($Events | Where-Object { $_.delivered_to_agent -eq $true })
  $Undelivered = @($Events | Where-Object { $_.delivered_to_agent -ne $true })
  $TokenValid = $true
  $TokenInvalidReasons = New-Object System.Collections.Generic.List[string]

  foreach ($Event in $Delivered) {
    $Bucket = [string]$Event.bucket
    if (-not $Buckets.Contains($Bucket)) {
      $Buckets[$Bucket] = New-BucketSummary
    }
    $Buckets[$Bucket].events += 1
    $Buckets[$Bucket].chars += [int64]$Event.chars
    $Buckets[$Bucket].utf8_bytes += [int64]$Event.utf8_bytes
    if ($Event.token_status -eq "exact" -and $null -ne $Event.tokens) {
      $Buckets[$Bucket].tokens += [int64]$Event.tokens
    } else {
      $Buckets[$Bucket].token_status = "invalid"
      $TokenValid = $false
      [void]$TokenInvalidReasons.Add("$($Event.event_type):$($Event.bucket):$($Event.token_status)")
    }
  }

  $TotalTokens = 0
  foreach ($BucketName in $Buckets.Keys) {
    $TotalTokens += [int64]$Buckets[$BucketName].tokens
  }

  $UndeliveredChars = 0
  $UndeliveredBytes = 0
  foreach ($Event in $Undelivered) {
    $UndeliveredChars += [int64]$Event.chars
    $UndeliveredBytes += [int64]$Event.utf8_bytes
  }

  $Summary = [ordered]@{
    schema_version = "agent_token_proxy_summary_v1"
    generated_utc = [DateTime]::UtcNow.ToString("o")
    run_dir = $RunRoot
    transcript = $LogPath
    delivered_event_count = $Delivered.Count
    undelivered_event_count = $Undelivered.Count
    delivered_chars = ($Delivered | Measure-Object -Property chars -Sum).Sum
    delivered_utf8_bytes = ($Delivered | Measure-Object -Property utf8_bytes -Sum).Sum
    undelivered_chars = $UndeliveredChars
    undelivered_utf8_bytes = $UndeliveredBytes
    agent_session_token_proxy = if ($TokenValid) { $TotalTokens } else { $null }
    token_measurement_valid = $TokenValid
    token_invalid_reasons = @($TokenInvalidReasons)
    buckets = $Buckets
  }

  $SummaryPath = Join-Path $RunRoot "summary.json"
  ($Summary | ConvertTo-Json -Depth 30) | Set-Content -LiteralPath $SummaryPath -Encoding UTF8
  $Summary | ConvertTo-Json -Depth 30

  if ($StrictTokens -and -not $TokenValid) {
    exit 2
  }
}

$RunRoot = Resolve-RepoPath $RunDir
New-Item -ItemType Directory -Path $RunRoot -Force | Out-Null
$LogPath = Join-Path $RunRoot "transcript.ndjson"

switch ($Action) {
  "init" {
    if (Test-Path -LiteralPath $LogPath) {
      Remove-Item -LiteralPath $LogPath -Force
    }
    $Meta = [ordered]@{
      repo_root = $RepoRoot
      run_dir = $RunRoot
      initialized_utc = [DateTime]::UtcNow.ToString("o")
      tokenizer_command = $TokenizerCommand
      tokenizer_encoding = $TokenizerEncoding
      tokenizer_path = Resolve-RepoPath $TokenizerPath
    }
    ($Meta | ConvertTo-Json -Depth 10) | Set-Content -LiteralPath (Join-Path $RunRoot "run.json") -Encoding UTF8
    Write-Output "initialized $RunRoot"
  }
  "prompt" {
    $Content = Get-InputText $Text $TextFile
    Add-TranscriptEvent $RunRoot $Phase "task_prompt" "prompt" $Content $true $TokenizerCommand $TokenizerEncoding $TokenizerPath $DisableTokenizer $null
    Write-Output "recorded prompt chars=$($Content.Length)"
  }
  "exec" {
    Add-TranscriptEvent $RunRoot $Phase "tool_call_argument" "exec_command" $Exec $true $TokenizerCommand $TokenizerEncoding $TokenizerPath $DisableTokenizer @{
      shell = $Shell
    }
    $Result = Invoke-CapturedCommand $Shell $Exec
    $StatusLine = "exit_code=$($Result.exit_code) elapsed_ms=$($Result.elapsed_ms)"
    Add-TranscriptEvent $RunRoot $Phase "delivered_tool_result" "exec_status" $StatusLine $true $TokenizerCommand $TokenizerEncoding $TokenizerPath $DisableTokenizer $null

    if ($Quiet) {
      Add-TranscriptEvent $RunRoot $Phase "local_tool_output_volume" "exec_stdout" $Result.stdout $false $TokenizerCommand $TokenizerEncoding $TokenizerPath $DisableTokenizer @{
        exit_code = $Result.exit_code
      }
      Add-TranscriptEvent $RunRoot $Phase "local_tool_output_volume" "exec_stderr" $Result.stderr $false $TokenizerCommand $TokenizerEncoding $TokenizerPath $DisableTokenizer @{
        exit_code = $Result.exit_code
      }
      Write-Output $StatusLine
      Write-Output "quiet_stdout_chars=$($Result.stdout.Length) quiet_stderr_chars=$($Result.stderr.Length)"
    } else {
      if ($Result.stdout.Length -gt 0) {
        Add-TranscriptEvent $RunRoot $Phase "delivered_tool_result" "exec_stdout" $Result.stdout $true $TokenizerCommand $TokenizerEncoding $TokenizerPath $DisableTokenizer @{
          exit_code = $Result.exit_code
        }
        [Console]::Out.Write($Result.stdout)
      }
      if ($Result.stderr.Length -gt 0) {
        $Bucket = if ($Result.exit_code -eq 0) { "delivered_tool_result" } else { "retry_error" }
        Add-TranscriptEvent $RunRoot $Phase $Bucket "exec_stderr" $Result.stderr $true $TokenizerCommand $TokenizerEncoding $TokenizerPath $DisableTokenizer @{
          exit_code = $Result.exit_code
        }
        [Console]::Error.Write($Result.stderr)
      }
      if ($Result.stdout.Length -eq 0 -and $Result.stderr.Length -eq 0) {
        Write-Output $StatusLine
      }
    }

    exit $Result.exit_code
  }
  "read" {
    if ([string]::IsNullOrWhiteSpace($Path)) {
      throw "The read action requires -Path."
    }
    $FullPath = Resolve-RepoPath $Path
    $Content = [System.IO.File]::ReadAllText($FullPath)
    Add-TranscriptEvent $RunRoot $Phase "delivered_file_content" "file_read" $Content (-not $Quiet) $TokenizerCommand $TokenizerEncoding $TokenizerPath $DisableTokenizer @{
      path = $FullPath
    }
    if ($Quiet) {
      $StatusLine = "quiet_file_read path=$FullPath chars=$($Content.Length)"
      Add-TranscriptEvent $RunRoot $Phase "delivered_tool_result" "file_read_status" $StatusLine $true $TokenizerCommand $TokenizerEncoding $TokenizerPath $DisableTokenizer $null
      Write-Output $StatusLine
    } else {
      [Console]::Out.Write($Content)
    }
  }
  "response" {
    $Content = Get-InputText $Text $TextFile
    Add-TranscriptEvent $RunRoot $Phase "agent_response" "response" $Content $true $TokenizerCommand $TokenizerEncoding $TokenizerPath $DisableTokenizer $null
    Write-Output "recorded response chars=$($Content.Length)"
  }
  "note" {
    $Content = Get-InputText $Text $TextFile
    Add-TranscriptEvent $RunRoot $Phase "delivered_tool_result" "note" $Content $true $TokenizerCommand $TokenizerEncoding $TokenizerPath $DisableTokenizer $null
    Write-Output $Content
  }
  "summary" {
    Write-Summary $RunRoot $StrictTokens
  }
}
