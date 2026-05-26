# AVmatrix Runtime Lock And Process Lifecycle Hardening Evidence Ledger

Date: 2026-05-26

Status: Planned

Companion files:

- Plan: [2026-05-26-avmatrix-runtime-lock-process-lifecycle-hardening-plan.md](2026-05-26-avmatrix-runtime-lock-process-lifecycle-hardening-plan.md)
- Benchmark ledger: [2026-05-26-avmatrix-runtime-lock-process-lifecycle-hardening-benchmark.md](2026-05-26-avmatrix-runtime-lock-process-lifecycle-hardening-benchmark.md)

## Evidence Rules

Record evidence as each evidenced task is completed. Evidence should include commands, impacted files, lock file samples, process inventory, parent process chains, test results, smoke artifacts, and concise observations needed to audit the plan later.

For doc-only commits, do not use AVmatrix.

Do not record inferred runtime behavior. Every behavior claim must include source inspection, command output, test output, or exact process/lock evidence.

## E0 - Plan Creation Evidence

Date: 2026-05-26

Status: recorded

Created file set:

- `docs/plans/2026-05-26-avmatrix-runtime-lock-process-lifecycle-hardening-plan.md`
- `docs/plans/2026-05-26-avmatrix-runtime-lock-process-lifecycle-hardening-evidence.md`
- `docs/plans/2026-05-26-avmatrix-runtime-lock-process-lifecycle-hardening-benchmark.md`

Plan creation scope:

- Address stale `.avmatrix/analyze.lock` behavior after crash, killed process, or reboot.
- Distinguish one-shot `analyze` work from long-running `mcp` and `serve` processes.
- Plan diagnostics so users and agents can see lock owner and process ownership without manual OS inspection.
- Preserve mutual exclusion for real concurrent writers.
- Avoid killing editor-owned MCP processes during browser or launcher cleanup.

Doc-only note:

- This plan creation is documentation-only, so AVmatrix was not used for this commit slice.

## E1 - Initial Runtime Trace From User Report

Date: 2026-05-26

Status: preliminary; implementation must re-verify before code edits

Observed user issue:

- User reported that after analyze runs, AVmatrix appears to keep a process alive even after browser close and even after PC restart.
- The behavior is considered a severe UX risk because it can make the tool appear stuck globally.

Local process inventory command:

```powershell
Get-CimInstance Win32_Process | Where-Object { $_.Name -like '*avmatrix*' } | Select-Object ProcessId,Name,CommandLine,CreationDate | Format-List
```

Observed process:

```text
ProcessId: 14008
Name: avmatrix.exe
CommandLine: "C:\Users\TAM NGUYEN\AppData\Roaming\npm\node_modules\avmatrix\bin\avmatrix.exe" mcp
CreationDate: 2026-05-26 14:53:15 local time
```

Parent chain command:

```powershell
$ids = 14008,912,532
foreach ($id in $ids) {
  Get-CimInstance Win32_Process -Filter "ProcessId = $id" |
    Select-Object ProcessId,ParentProcessId,Name,ExecutablePath,CommandLine,CreationDate |
    Format-List
}
```

Observed parent chain:

| PID | Process | Parent | Command summary |
|---:|---|---:|---|
| 14008 | `avmatrix.exe` | 912 | `avmatrix.exe mcp` from npm package path. |
| 912 | `cmd.exe` | 532 | Runs `avmatrix.cmd mcp`. |
| 532 | `codex.exe` | 1616 | Codex session process. |

Conclusion:

- The currently observed live process is editor/agent-owned `avmatrix mcp`, not an `analyze` process.
- Closing a browser is not expected to stop this process because it is owned by Codex through MCP setup.
- UX still needs improvement because users can reasonably misinterpret any long-running `avmatrix.exe` as a stuck analyze process.

## E2 - Initial Stale Lock Finding

Date: 2026-05-26

Status: preliminary; implementation must re-verify before code edits

Earlier stale lock observation during investigation:

```text
E:\AVmatrix-GO\.avmatrix\analyze.lock
pid=3108
acquiredAt=2026-05-26T07:50:03.5665076Z
```

PID check:

```powershell
Get-Process -Id 3108 -ErrorAction SilentlyContinue
```

Observed result:

- PID `3108` did not exist.
- Removing the stale lock allowed `.\avmatrix\bin\avmatrix.exe analyze --force` to run successfully.

Conclusion:

- A stale lock file can block analyze even when no analyze process is alive.
- This matches a reboot/crash/killed-process failure mode because process lifetime and lock file lifetime are not tied together.

## E3 - Initial Source Trace

Date: 2026-05-26

Status: preliminary; implementation must re-verify with fresh AVmatrix graph before code edits

Source search command:

```powershell
rg -n "analyze\.lock|index lock|lock is already held|acquiredAt|repository index lock|With.*Lock|lock file|stale lock|Remove.*lock|pid=" internal cmd avmatrix-launcher -g "*.go" -g "*.ps1"
```

Observed owner files:

| Area | Path | Finding |
|---|---|---|
| Lock implementation | `internal/repo/lock.go` | `AcquireStorageLock` creates lock with `O_CREATE|O_EXCL`; `Release` removes the lock path. |
| Lock tests | `internal/repo/lock_test.go` | Tests concurrent exclusion and release, but not stale lock recovery. |
| CLI analyze | `internal/analyze/analyze.go` | `Run` acquires `repo.Paths(resolvedPath).AnalyzeLockPath` and defers release. |
| HTTP analyze preflight | `internal/httpapi/analyze.go` | `ensureAnalyzeLockAvailable` acquires and releases a lock to check availability. |
| HTTP embed | `internal/httpapi/embed.go` | Embed acquires the same analyze lock and releases it after job completion. |
| Repo paths | `internal/repo/paths.go` | Lock path is `.avmatrix/analyze.lock`. |
| Setup MCP config | `internal/cli/setup_command.go` | Setup writes editor MCP entries with command `avmatrix` and args `["mcp"]`. |
| Launcher cleanup | `avmatrix-launcher/src/main.go` | Cleanup targets launcher-owned `serve --port 4848` and support processes, not editor-owned MCP. |

Key lock implementation observed:

```go
file, err := os.OpenFile(lockPath, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0o644)
if err != nil {
    if errors.Is(err, os.ErrExist) {
        return nil, ErrLockHeld
    }
    return nil, err
}
fmt.Fprintf(file, "pid=%d\nacquiredAt=%s\n", os.Getpid(), time.Now().UTC().Format(time.RFC3339Nano))
```

Conclusion:

- The current lock implementation cannot recover stale locks because it never reads existing lock metadata.
- The current metadata is enough to start a compatibility path, but not enough for robust ownership-safe release.

## E4 - Initial MCP Setup Finding

Date: 2026-05-26

Status: preliminary; implementation must re-verify before code edits

Config inspection command:

```powershell
$userHome=$env:USERPROFILE
$paths=@(
  (Join-Path $userHome ".codex\config.toml"),
  (Join-Path $userHome ".cursor\mcp.json"),
  (Join-Path $userHome ".claude.json"),
  (Join-Path $userHome ".config\opencode\opencode.json")
)
foreach ($pathItem in $paths) {
  if (Test-Path $pathItem) {
    Select-String -LiteralPath $pathItem -Pattern "avmatrix|mcp_servers|mcpServers|command|args" -Context 1,2
  }
}
```

Observed local Codex config:

```toml
[mcp_servers.avmatrix]
command = "avmatrix"
args = ["mcp"]
```

Observed local Claude config:

```json
{
  "mcpServers": {
    "avmatrix": {
      "args": ["mcp"],
      "command": "avmatrix"
    }
  }
}
```

Conclusion:

- `avmatrix mcp` is intentionally configured as an editor/agent MCP server.
- Process diagnostics and docs should distinguish this expected long-running process from stuck analyze work.

## E5 - Initial Impact Evidence

Date: 2026-05-26

Status: preliminary; implementation must re-run before code edits

Impact command:

```powershell
.\avmatrix\bin\avmatrix.exe impact AcquireStorageLock --repo AVmatrix --direction upstream
```

Observed summary:

```text
risk: CRITICAL
impactedCount: 8
affected_app_layers: api=3, backend=4, cli_launcher=1
affected_functional_areas: analyzer=1, api=3, cli=3, graph_health=1
processes_affected: 53
```

Direct affected symbols included:

- `internal/analyze/analyze.go:Run`
- `internal/httpapi/analyze.go:ensureAnalyzeLockAvailable`
- `internal/httpapi/embed.go:Server.handleEmbed`

Conclusion:

- Lock hardening is high-impact shared infrastructure.
- CRITICAL risk is expected and should be handled with focused tests, not avoided.

## E6 - Implementation Evidence

Date: pending

Status: pending

Record here:

- Fresh AVmatrix analyze and impact evidence for implementation slices.
- Edited source files.
- Lock metadata examples before and after implementation.
- Unit and integration test results.
- Process diagnostics smoke output.
- Full build result.
- `detect-changes` output before commit.
- Commit hashes.
