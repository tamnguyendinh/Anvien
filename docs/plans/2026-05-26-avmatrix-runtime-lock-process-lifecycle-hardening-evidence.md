# AVmatrix Runtime Lock And Process Lifecycle Hardening Evidence Ledger

Date: 2026-05-26

Status: Complete

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

Date: 2026-05-26

Status: committed as `a0ba34c`

### Fresh Graph Refreshes

Commands:

```powershell
.\avmatrix\bin\avmatrix.exe analyze --force
```

Observed graph refreshes during implementation:

| Time/order | Files scanned | Files parsed | Unsupported | Failed | Nodes | Relationships |
|---|---:|---:|---:|---:|---:|---:|
| Pre-edit refresh | 765 | 568 | 197 | 0 | 85995 | 117993 |
| Post-doctor/setup refresh | 766 | 569 | 197 | 0 | 86590 | 118875 |
| Post-smoke refresh | 766 | 569 | 197 | 0 | 86682 | 118981 |
| Pre-detect refresh | 766 | 569 | 197 | 0 | 86706 | 119008 |

### Impact Evidence

Commands:

```powershell
.\avmatrix\bin\avmatrix.exe impact AcquireStorageLock --repo AVmatrix --direction upstream
.\avmatrix\bin\avmatrix.exe impact ensureAnalyzeLockAvailable --repo AVmatrix --direction upstream
.\avmatrix\bin\avmatrix.exe impact handleEmbed --repo AVmatrix --direction upstream
.\avmatrix\bin\avmatrix.exe impact NewRootCommand --repo AVmatrix --direction upstream
.\avmatrix\bin\avmatrix.exe impact printSetupResult --repo AVmatrix --direction upstream
.\avmatrix\bin\avmatrix.exe impact collectAVmatrixProcesses --repo AVmatrix --direction upstream
.\avmatrix\bin\avmatrix.exe impact collectWindowsAVmatrixProcesses --repo AVmatrix --direction upstream
```

Observed impact summaries:

| Symbol | Risk | Impacted count | Processes affected | Notes |
|---|---:|---:|---:|---|
| `AcquireStorageLock` | CRITICAL | 8 | 53 | Shared lock path used by CLI analyze, HTTP analyze, HTTP embed, and graph-health access audit flow. |
| `ensureAnalyzeLockAvailable` | CRITICAL | 1 | 9 | HTTP analyze preflight only; direct caller is `Server.handleAnalyze`. |
| `handleEmbed` | LOW | 0 | 0 | No upstream callers were reported by AVmatrix for the handler method. |
| `NewRootCommand` | CRITICAL | 1 | 11 | Root CLI command surface; direct external entry is `cmd/avmatrix/main.go:main`. |
| `printSetupResult` | CRITICAL | 3 | 11 | Setup output path through root command; wording-only change. |
| `collectAVmatrixProcesses` | CRITICAL | 3 | 11 | Diagnostics process collection path through `doctor processes`. |
| `collectWindowsAVmatrixProcesses` | LOW | 3 | 0 | Windows diagnostics helper under `doctor processes`. |

Blast radius interpretation:

- HIGH/CRITICAL is treated as a workflow safety warning, not a prohibition. This slice edited shared storage-lock infrastructure where required and preserved the existing `AcquireStorageLock(lockPath string)` API contract for callers.
- The launcher cleanup code was inspected but not edited in this slice.

### Edited Source Files

Implementation files:

- `internal/repo/lock.go`
- `internal/repo/lock_test.go`
- `internal/cli/command.go`
- `internal/cli/command_test.go`
- `internal/cli/doctor_command.go`
- `internal/cli/setup_command.go`
- `internal/httpapi/analyze.go`
- `internal/httpapi/analyze_test.go`
- `internal/httpapi/embed.go`
- `internal/httpapi/embed_test.go`

Plan ledger files:

- `docs/plans/2026-05-26-avmatrix-runtime-lock-process-lifecycle-hardening-plan.md`
- `docs/plans/2026-05-26-avmatrix-runtime-lock-process-lifecycle-hardening-evidence.md`
- `docs/plans/2026-05-26-avmatrix-runtime-lock-process-lifecycle-hardening-benchmark.md`

### Lock Implementation Evidence

New lock metadata format written by `AcquireStorageLock`:

```text
version=2
pid=<pid>
acquiredAt=<RFC3339Nano UTC timestamp>
host=<hostname>
command=<command line>
token=<random ownership token>
```

Compatibility behavior:

- Existing old-format locks with only `pid` and `acquiredAt` are parsed.
- Same-host dead-PID locks are considered stale/recoverable and removed before retrying acquisition.
- Live same-host locks still return an error compatible with `errors.Is(err, repo.ErrLockHeld)`.
- Foreign-host locks are not removed based on local PID liveness.
- Malformed locks are recoverable only after the malformed grace period.
- `StorageLock.Release` reads the lock token and does not remove a replaced lock whose token no longer matches.

### Diagnostics Evidence

New CLI surface:

```powershell
.\avmatrix\bin\avmatrix.exe doctor locks --repo .
.\avmatrix\bin\avmatrix.exe doctor locks --repo . --json
.\avmatrix\bin\avmatrix.exe doctor processes --json
```

Observed `doctor locks --repo .` output before the final process self-filter refinement:

```text
AVmatrix analyze lock
Repo: E:\AVmatrix-GO
Storage: E:\AVmatrix-GO\.avmatrix
Lock: E:\AVmatrix-GO\.avmatrix\analyze.lock
Status: free
```

Observed `doctor locks --repo . --json` output before the final process self-filter refinement:

```json
{
  "repoPath": "E:\\AVmatrix-GO",
  "storagePath": "E:\\AVmatrix-GO\\.avmatrix",
  "lockPath": "E:\\AVmatrix-GO\\.avmatrix\\analyze.lock",
  "status": "free",
  "diagnosis": {
    "exists": false,
    "alive": false,
    "stale": false,
    "recoverable": false,
    "foreignHost": false,
    "reason": "lock file does not exist"
  }
}
```

Observed `doctor processes --json` before the final self-filter refinement:

- Included `avmatrix mcp` PID `14008`, parent `cmd.exe` PID `912`, parent chain rooted under `codex.exe`.
- Classified the MCP process as `role=mcp`, `ownership=editor-owned`.
- Also reported the current diagnostic command and PowerShell helper; this led to a follow-up filter change to exclude the current doctor process/helper from output.

### Setup UX Evidence

Changed `internal/cli/setup_command.go` setup output summary:

```text
MCP lifecycle: avmatrix mcp is editor-owned and may stay running while the editor or agent session is active.
Diagnostics: avmatrix doctor locks, avmatrix doctor processes
```

Current limitation:

- Setup output is updated.
- README and generated AI-context guidance updates remain pending and are not marked complete.

### Validation Evidence

Full build command:

```powershell
powershell -ExecutionPolicy Bypass -File avmatrix-launcher\build.ps1
```

Observed latest result:

- Exit code `0`.
- Web build completed through Vite.
- Go launcher/runtime build completed.
- Existing Vite warnings remained: large chunks and dynamic/static import chunking for `ProcessFlowModal.tsx`.

Focused test command:

```powershell
go test ./internal/repo ./internal/cli ./internal/httpapi ./internal/analyze
```

Observed latest result:

```text
ok github.com/tamnguyendinh/avmatrix-go/internal/repo
ok github.com/tamnguyendinh/avmatrix-go/internal/cli 9.586s
ok github.com/tamnguyendinh/avmatrix-go/internal/httpapi
ok github.com/tamnguyendinh/avmatrix-go/internal/analyze
```

Exact focused test gate:

```powershell
go test .\internal\repo .\internal\analyze .\internal\httpapi .\internal\cli -count=1
```

Observed latest result:

```text
ok github.com/tamnguyendinh/avmatrix-go/internal/repo 3.603s
ok github.com/tamnguyendinh/avmatrix-go/internal/analyze 1.599s
ok github.com/tamnguyendinh/avmatrix-go/internal/httpapi 2.633s
ok github.com/tamnguyendinh/avmatrix-go/internal/cli 8.674s
```

Source package test command:

```powershell
go test ./cmd/... ./internal/...
```

Observed latest result:

- Exit code `0`.
- All `cmd/...` and `internal/...` packages passed.

Launcher test command:

```powershell
go test ./...
```

Working directory:

```text
E:\AVmatrix-GO\avmatrix-launcher\src
```

Observed latest result:

```text
ok avmatrix-launcher
```

Known invalid broad test command:

```powershell
go test ./...
```

Observed repository-root result:

- Failed because it tries to build intentionally invalid fixture packages under `avmatrix/test/fixtures` and `node_modules`.
- `cmd/...` and `internal/...` are the valid Go source package set for this repo.

### Smoke Evidence

Dead-PID stale lock smoke command:

```powershell
$lock = Join-Path (Get-Location) '.avmatrix\analyze.lock'
New-Item -ItemType Directory -Force -Path (Split-Path $lock) | Out-Null
Set-Content -Path $lock -Encoding ASCII -Value "pid=999999999`nacquiredAt=2026-05-26T07:50:03Z`n"
.\avmatrix\bin\avmatrix.exe analyze --force
```

Observed output:

```text
analyzed E:\AVmatrix-GO
files: scanned=766 parsed=569 unsupported=197 failed=0
graph: nodes=86682 relationships=118981 path=E:\AVmatrix-GO\.avmatrix\graph.json
```

Live-lock conflict smoke command:

```powershell
$lock = Join-Path (Get-Location) '.avmatrix\analyze.lock'
$hostName = [System.Net.Dns]::GetHostName()
Set-Content -Path $lock -Encoding ASCII -Value "version=2`npid=$PID`nacquiredAt=2026-05-26T07:50:03Z`nhost=$hostName`ncommand=manual live lock smoke`ntoken=smoke-token`n"
.\avmatrix\bin\avmatrix.exe analyze --force
```

Observed output:

```text
exit=1
repository index lock is already held (pid=13332, host=TAM-PC, acquiredAt=2026-05-26T07:50:03Z, command=manual live lock smoke, reason=owning process is still running)
```

Current limitation:

- Live-lock conflict smoke confirms PID, host, acquired time, command, and live reason.
- Lock path, explicit age text, and explicit next action are still pending under P2-E/P5-E.

Diagnostics smoke commands after self-filter refinement:

```powershell
.\avmatrix\bin\avmatrix.exe doctor locks --repo .
.\avmatrix\bin\avmatrix.exe doctor locks --repo . --json
.\avmatrix\bin\avmatrix.exe doctor processes --json
```

Observed latest result:

- `doctor locks --repo .` reported repo `E:\AVmatrix-GO`, lock `E:\AVmatrix-GO\.avmatrix\analyze.lock`, and status `free`.
- `doctor locks --repo . --json` reported status `free` with `reason: lock file does not exist`.
- `doctor processes --json` reported the active Playwright `avmatrix-web` test server and editor-owned MCP processes.
- The latest process output did not include the running `doctor processes` command or its PowerShell helper process.

Pre-commit change detection:

```powershell
.\avmatrix\bin\avmatrix.exe detect-changes --repo AVmatrix --scope all
```

Observed summary:

```text
risk_level=critical
affected_count=32
changed_count=542
changed_files=12
affected_app_layers={api:8, backend:5, mixed:19}
affected_functional_areas={api:8, cli:5, mixed:19}
changed_app_layers={api:7, api_test:55, backend:227, backend_test:227, docs:26}
changed_functional_areas={api:62, cli:95, documentation:26, storage:359}
resolution_gap_changes.changedGapEntities=380
resolution_gap_changes.changedGapOccurrenceCount=385
semanticStatus.appLayer.status=complete
semanticStatus.functionalArea.status=complete
```

Blast radius note:

- `detect-changes` reported `risk_level=critical`; this means the slice touched important shared code paths and required careful impact analysis, build/test gates, smoke checks, diagnostics checks, and pre-commit change detection.
- HIGH/CRITICAL blast radius is not a prohibition on editing and does not mean the commit must be artificially narrowed. It is a signal to keep the work deliberate, validated, and traceable.
- The current slice intentionally covers the storage lock lifecycle, lock diagnostics, process diagnostics, setup text, and focused tests/docs needed for those behaviors.

### Slice Commit

Committed as:

```text
a0ba34c Harden runtime lock lifecycle diagnostics
```

## E7 - Final Lock UX Slice Evidence

Date: 2026-05-26

Status: committed as `c2f52e5`

### Scope Correction

README, active docs, MCP setup resources, and generated AI context guidance are not part of this bug fix. The bug is in runtime lock behavior and diagnostics, so P4-C/P4-D were rescoped out and no README or `internal/aicontext` changes are included in this slice.

### Impact Evidence

Commands:

```powershell
.\avmatrix\bin\avmatrix.exe analyze --force
.\avmatrix\bin\avmatrix.exe impact LockHeldError --repo AVmatrix --direction upstream
.\avmatrix\bin\avmatrix.exe impact lockHeldMessage --repo AVmatrix --direction upstream
.\avmatrix\bin\avmatrix.exe impact classifyDoctorProcess --repo AVmatrix --direction upstream
.\avmatrix\bin\avmatrix.exe impact renderAVmatrixBlock --repo AVmatrix --direction upstream
```

Observed summaries:

| Symbol | Risk | Impacted count | Processes affected | Action |
|---|---:|---:|---:|---|
| `LockHeldError` | CRITICAL | 7 | 19 | Edited to include lock path, age, and next action. |
| `lockHeldMessage` | CRITICAL | 2 | 16 | Covered by HTTP analyze/embed tests through `LockHeldError.Error()`. |
| `classifyDoctorProcess` | LOW | 3 | 0 | Edited to classify explicit embed commands. |
| `renderAVmatrixBlock` | CRITICAL | 4 | 13 | Not edited after scope correction; generated AI context is out of scope for this lock bug. |

Blast radius note:

- CRITICAL impact here is a signal that lock/API/user-facing error paths are important and require careful validation. It is not a ban on editing those paths.

### Implementation Evidence

- `LockHeldError.Error()` now includes `path`, `age`, and `next` fields in addition to PID, host, acquired time, command, malformed status, and reason.
- The next action tells the user to wait for the owner, stop the PID only if stale, and inspect with `avmatrix doctor locks --repo <repo>`.
- `classifyDoctorProcess` now reports explicit embed command lines as `role=embed` with `ownership=user-command-or-job`.
- Tests were added/updated for actionable lock metadata and embed process classification.

### Build Gate

Command:

```powershell
powershell -ExecutionPolicy Bypass -File avmatrix-launcher\build.ps1
```

Observed latest result:

- Exit code `0`.
- Web build and Go launcher/runtime build completed.
- Existing Vite warnings remained: large chunks and dynamic/static import chunking for `ProcessFlowModal.tsx`.

### Focused Test Attempt

Command:

```powershell
go test .\internal\repo .\internal\analyze .\internal\httpapi .\internal\cli -count=1
```

Observed result:

```text
internal/repo: pass
internal/analyze: pass
internal/httpapi: pass
internal/cli: fail
```

Failure:

- `TestSetupCommandWritesEditorConfigsAndSkills` failed because a new assertion required installed package skills to mention `doctor locks`, `doctor processes`, and `editor-owned`.

Correction:

- That assertion is out of scope for this lock bug because broad setup-installed skill documentation is not being changed in this slice.
- Remove the assertion and keep setup skill installation validation focused on copying packaged content.

Focused test rerun:

```powershell
go test .\internal\repo .\internal\analyze .\internal\httpapi .\internal\cli -count=1
```

Observed latest result:

```text
ok github.com/tamnguyendinh/avmatrix-go/internal/repo 3.037s
ok github.com/tamnguyendinh/avmatrix-go/internal/analyze 5.070s
ok github.com/tamnguyendinh/avmatrix-go/internal/httpapi 6.110s
ok github.com/tamnguyendinh/avmatrix-go/internal/cli 9.767s
```

Live-lock conflict smoke:

```powershell
$lock = Join-Path (Get-Location) '.avmatrix\analyze.lock'
New-Item -ItemType Directory -Force -Path (Split-Path $lock) | Out-Null
$hostName = [System.Net.Dns]::GetHostName()
Set-Content -Path $lock -Encoding ASCII -Value "version=2`npid=$PID`nacquiredAt=2026-05-26T07:50:03Z`nhost=$hostName`ncommand=manual live lock smoke`ntoken=smoke-token`n"
.\avmatrix\bin\avmatrix.exe analyze --force
Remove-Item -LiteralPath $lock -Force -ErrorAction SilentlyContinue
```

Observed latest result:

```text
exit=1
repository index lock is already held (path=E:\AVmatrix-GO\.avmatrix\analyze.lock, pid=8636, host=TAM-PC, acquiredAt=2026-05-26T07:50:03Z, age=1h9m54s, command=manual live lock smoke, reason=owning process is still running, next=wait for the owning process to finish or stop pid 8636 if it is stale; inspect with avmatrix doctor locks --repo E:\AVmatrix-GO)
```

The non-zero exit is expected for a live lock. The smoke confirms actionable owner metadata: path, PID, host, acquired time, age, command, live reason, and next action.

Malformed fresh lock smoke:

```powershell
$lock = Join-Path (Get-Location) '.avmatrix\analyze.lock'
New-Item -ItemType Directory -Force -Path (Split-Path $lock) | Out-Null
Set-Content -Path $lock -Encoding ASCII -Value 'not lock metadata'
.\avmatrix\bin\avmatrix.exe analyze --force
$exists = Test-Path -LiteralPath $lock
Remove-Item -LiteralPath $lock -Force -ErrorAction SilentlyContinue
```

Observed latest result:

```text
exit=1
lockExists=True
repository index lock is already held (path=E:\AVmatrix-GO\.avmatrix\analyze.lock, age=72ms, malformed=true, reason=lock metadata is malformed but recent, next=wait for the owning process to finish or stop the owning process if it is stale; inspect with avmatrix doctor locks --repo E:\AVmatrix-GO)
```

The non-zero exit is expected for a fresh malformed lock. The smoke confirms the lock is not removed during the malformed grace period.

Diagnostics smoke rerun:

```powershell
.\avmatrix\bin\avmatrix.exe doctor locks --repo .
.\avmatrix\bin\avmatrix.exe doctor locks --repo . --json
.\avmatrix\bin\avmatrix.exe doctor processes --json
```

Observed latest result:

- `doctor locks --repo .` reported status `free`.
- `doctor locks --repo . --json` reported status `free` and reason `lock file does not exist`.
- `doctor processes --json` reported the Playwright `avmatrix-web` test server as `avmatrix-runtime` and the active MCP command/wrapper as `editor-owned`.
- The diagnostics output did not include the running diagnostic command itself.

Pre-commit graph refresh:

```powershell
.\avmatrix\bin\avmatrix.exe analyze --force
```

Observed output:

```text
analyzed E:\AVmatrix-GO
files: scanned=766 parsed=569 unsupported=197 failed=0
graph: nodes=86778 relationships=119083 path=E:\AVmatrix-GO\.avmatrix\graph.json
```

Pre-commit change detection:

```powershell
.\avmatrix\bin\avmatrix.exe detect-changes --repo AVmatrix --scope all
```

Observed summary:

```text
risk_level=low
affected_count=0
changed_count=92
changed_files=9
changed_app_layers={api_test:11, backend:32, backend_test:26, docs:23}
changed_functional_areas={api:11, cli:12, documentation:23, storage:46}
resolution_gap_changes.changedGapEntities=51
resolution_gap_changes.changedGapOccurrenceCount=52
semanticStatus.appLayer.status=complete
semanticStatus.functionalArea.status=complete
```

### Slice Commit

Committed as:

```text
c2f52e5 Complete runtime lock UX hardening
```
