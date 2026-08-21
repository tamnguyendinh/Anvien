# Child 05 / P5-B Export Tables — Evidence Blocker Resubmission

## Status and scope

- Status: `READY_FOR_SUPERVISOR` (handoff only; not a Supervisor verdict, self-acceptance, detect-changes result, or commit).
- Scope of this resubmission: close the two P0 evidence blockers identified by Supervisor. Production source, tests, ledgers, and the prior Coder report were not edited.
- Repository and required working directory: `E:\Anvien`.
- Git basis: `E:\Anvien`, HEAD `17f3dad2587f2785d02ad266e188c7a8feca499b`.
- Authorized P5-B source/test files remain uncommitted exactly as handed off. The six Main orchestration handoff reports remain untracked and untouched.

## Blocker 1 — absolute canonical build evidence

### Preflight and lock remediation

1. `anvien doctor locks --repo E:\Anvien --json` returned `status=free`; `E:\Anvien\.anvien\analyze.lock` did not exist.
2. `anvien doctor processes --json` initially showed only editor-owned MCP processes. The first absolute build probe then reported that `E:\Anvien\anvien\bin\anvien.exe` was held by another process.
3. A Restart Manager resource query for that exact output identified PID `15816` (`C:\Users\TAM NGUYEN\AppData\Roaming\npm\node_modules\anvien\bin\anvien.exe mcp`) as the direct holder. That build-related holder was terminated; its parent PID `15176` was already gone. The target handle list was then empty and the Anvien lock remained free.

The initial probe is recorded as a failed preflight attempt (`exit 1`, Go linker sharing violation). It is not used as build evidence. After the holder was cleared, the required command below completed successfully.

### Successful canonical run

Literal command (absolute path, exactly as executed):

```text
powershell -ExecutionPolicy Bypass -File E:\Anvien\scripts\full-build.ps1
```

- cwd: `E:\Anvien`
- exit code: `0`
- script stages observed: package/runtime build, launcher/server build, web bundle build, CLI version `1.2.8`, and final `anvien analyze . --force`.
- analyze result: `scanned=1956`, `parsed_code=738`, `failed=0`, graph `nodes=115099`, `relationships=158118`.

Fresh canonical E outputs after that successful run:

| Output | Bytes | Last write (UTC) | SHA-256 |
|---|---:|---|---|
| `E:\Anvien\anvien\bin\anvien.exe` | 71,395,328 | `2026-08-21T12:41:16.3504227Z` | `CB8A79C4873E87C03F39FBB0A72116DDC599469E3E46661B3767022E349558B1` |
| `E:\Anvien\anvien\bin\lbug_shared.dll` | 20,230,656 | `2026-08-09T12:39:05.9469203Z` | `20CBD87840483A2053CFF3FC2DB23A86DD802B8915D86509D41A4B709624CDB7` |
| `E:\Anvien\.anvien\graph.json` | 458,436,615 | `2026-08-21T12:42:04.6643519Z` | `B15832386D74773CF447D0A9DC396A441CA9E52C2BC0959BA6FF3EAFA715E8B7` |

Post-build lock verification again reported `status=free`; no build process was left running.

## Blocker 1 — source/test chronology and unchanged bytes

The baseline snapshot was recorded before the successful rerun after lock remediation. The final snapshot was captured at `2026-08-21T12:43:58.3865675Z`. Each authorized source/test file has the same size, mtime, and SHA-256 before and after the build:

| Source/test file | Bytes | Last write (UTC) | SHA-256 |
|---|---:|---|---|
| `E:\Anvien\internal\resolution\export_tables.go` | 7,213 | `2026-08-21T12:18:48.1624491Z` | `BA4C7F9F2D05F6715CA22195C1CFF64C1AACADC3D81E7CA3ED711C8EA8B63E19` |
| `E:\Anvien\internal\resolution\indexes.go` | 32,345 | `2026-08-21T12:05:59.2930580Z` | `26DE75A911B4B1E1471C61548C4153749E067671DB34852E51448D52D4C0C486` |
| `E:\Anvien\internal\resolution\export_tables_test.go` | 9,751 | `2026-08-21T12:24:43.0991903Z` | `A0395EF9030CB054B09C52AFBE048D2F8244BAD8A924EE492BB7AB312CD08AD8` |

Fresh post-build test chronology:

```text
2026-08-21T12:42:52.1179813Z -> 2026-08-21T12:42:54.9189738Z
go test ./internal/resolution -run '^(TestBuildExportTablesZeroPhysicalBarrelAndDeterministicEntries|TestBuildWorkspaceWiresExportTablesWithoutPhysicalDefinitions|TestBuildExportTablesRetainsExplicitFactKindsAndMeanings)$' -count=1
PASS, package duration 0.190s, exit 0

2026-08-21T12:43:00.3750208Z -> 2026-08-21T12:43:03.1897552Z
go test ./internal/resolution -count=1
PASS, package duration 0.238s, exit 0
```

The earlier nearest-boundary evidence remains valid and unchanged: `go test ./internal/resolution ./internal/analyze ./internal/cli -count=1` passed (`resolution 0.259s; analyze 2.904s; cli 119.566s`) on the same byte-identical source state. The canonical build itself also passed the real CLI `version` and analyze boundary.

## P5-B scope preservation

- `internal/resolution/export_tables.go` remains the deterministic semantic owner over accepted Child 04 `ScopeIR.ExportFact` (explicit entries and star adjacency, exact provenance/meaning).
- `internal/resolution/indexes.go` contains only the authorized storage field and `buildWorkspace` wiring call.
- No Child 04 facts/provider/ScopeIR, import resolver logic, graph/persistence/readers, P5-C traversal, P5-D emission, target, or other drive/worktree was changed.
- Existing fixed-corpus evidence remains the accepted result: 736 parsed files, `ImportsResolved=5,072`, `FinalizedImportsEmitted=5,072`, persisted `IMPORTS=5,088`, delta `0/0/0`, corpus artifact SHA-256 `7687C5197AE76DC74294FB948AB9E25124BCEF26E444230EB29A2771EC373796`.
- `anvien detect-changes`, commit, Supervisor, target, and `E:\cheapapp.org` were not run or touched in this resubmission.

## Invariant Family Map

- Authority / SSOT: accepted Child 04 `ScopeIR.ExportFact` plus existing P5-A resolved target-file candidates.
- Sibling surfaces checked: `workspace`, `buildWorkspace`, deterministic explicit-entry/star-table construction, focused P5-B tests, the full `internal/resolution` package, and the real CLI/analyze boundary. They remain byte-identical to the authorized implementation state.
- Legacy fallback status: no physical-definition export inference, global-name rescue, target special case, second semantic source, side-effect expansion, or terminal/proof outcome was introduced.
- Residual unverified surfaces: `none` inside the authorized P5-B invariant family. `detect-changes`, commit, and Supervisor are external workflow gates intentionally pending Main control.

## E2E Verification

```text
E2E Verification:
  [PASS] Compiled: powershell -ExecutionPolicy Bypass -File E:\Anvien\scripts\full-build.ps1 -> exit 0, canonical E outputs present
  [PASS] Runtime: CLI version 1.2.8 and final analyze boundary -> PASS
  [PASS] Happy path: accepted ExportFact -> explicit exportTable/star adjacency -> zero-physical barrel -> PASS
  [PASS] Edge case: deterministic ordering, exact provenance/meaning, no implicit default, owning-copy isolation -> PASS
```

## Blocker 2 — immutable fresh report handoff

The rejected prior report was not edited:

```text
E:\Anvien\reports\coder\rp_coder_260821_192131_by_gpt-5_child05_p5b_export_tables_ready_for_supervisor.md
current bytes: 9696
current SHA-256: D41F6704EEFFBD616B7EFDBD25521A0C2FBC87E7B209FC760BDD68204E1CAD99
prior handoff digest (for mismatch traceability): E6CB4195F7535151D053A07CFD0FE94DD9FFDD336BB1B66BC48A2997815CB9E9
```

This new report is the unique resubmission artifact. Its final byte identity (byte length, LF count, UTF-8 BOM flag, and SHA-256) is supplied in the Main handoff and must be treated as immutable; no edits are made after identity calculation.

`READY_FOR_SUPERVISOR`
