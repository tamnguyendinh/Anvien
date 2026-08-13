# QA report — Child 02 P2-D repeated normal built analyze

## Handoff status

**READY_FOR_SUPERVISOR**

This is an independent QA handoff, not an acceptance verdict. P2-D remains unaccepted until the independent Supervisor reviews this report and its paired evidence. P2-E and later slices were not opened.

## Scope and repository state

- Role: validation-first QA for P2-D only.
- Owners measured:
  - C04: `internal/analyze/analyze.go::{writeGraphSnapshotJSON,Run}`.
  - C17 P2-D read/failure boundary: `internal/mcp/tools.go::Server.runCypherRead`.
- Inspect-only context: accepted P2-B persistence and P2-C affected-reader evidence.
- Production source, existing product tests, plans, ledgers, contract, rules, and accepted reports were not edited.
- Repository: `E:\Anvien`.
- Branch before and after: `master`.
- HEAD before and after: `927a676653963e8001d7789291010d5b819bac83` (`fix(readers): consume corrected graph facts`).
- Initial worktree: clean; staged paths: 0.
- Final permitted dirty paths are recorded in Final containment and Git state below; staged paths remain 0.

## Authority and QA lenses read in full

- `E:\Anvien\AGENTS.md`.
- `E:\Anvien\.agents\skills\working-rules\SKILL.md`.
- `E:\Anvien\.agents\skills\qa\SKILL.md`.
- `E:\Anvien\.agents\skills\Edge-Case\SKILL.md`.
- `E:\Anvien\.agents\skills\Data-Integrity\SKILL.md`.
- `E:\Anvien\.agents\skills\backend-development\SKILL.md` as a read-only source-review lens.
- `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-26-anvien-graph-accuracy-roadmap.md`.
- All four Child 02 ledgers under `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-02-current-graph-persistence-and-reader-consistency/`.
- `docs/contracts/graph-accuracy-contract.md`.
- Accepted P2-B Supervisor report `reports/Supervisor/rp_supervisor_260813_173541_by_gpt-5_child02_p2b_persistence.md`.
- Accepted P2-C re-review `reports/Supervisor/rp_supervisor_260813_203838_by_gpt-5_child02_p2c_affected_readers_rereview.md`; verified SHA-256 `F54C4225520FFC57A6438CE3F4C3B0816FE8095B31AFB55FD004E91A89808698`.
- Current source, tests, build/package scripts, and command dispatch needed for C04/C17 and the normal built runtime.

The QA, edge-case, and data-integrity lenses shaped the matrix around invocation-to-artifact correlation, stale/substitute exclusion, fail-closed behavior, fault recovery, exact identities/ranges/endpoints, and ownership-safe cleanup. The backend lens was used only to trace the source boundary; it did not authorize production edits.

## Anvien graph evidence and blast radius

The session ran `anvien --help`, then refreshed the repository graph with `anvien analyze --force` before graph-based inspection. The refreshed index was current at HEAD `927a676`; the full-build refresh later reported `files scanned=1604`, `parsed_code=686`, `failed=0`, `nodes=96692`, and `relationships=135854`.

Current file-detail/impact evidence:

| Boundary | File state / host risk | Upstream impact | Interpretation |
|---|---|---|---|
| `internal/analyze/analyze.go::writeGraphSnapshotJSON` | current, non-stale; host risk HIGH | CRITICAL; 22 impacted, 1 direct, 5 modules, 23 processes | Broad persistence blast-radius warning; covered at the normal built boundary. |
| `internal/analyze/analyze.go::Run` | current, non-stale; host risk HIGH | CRITICAL; 23 impacted, 17 direct, 7 modules, 27 processes | Broad analyze blast-radius warning; covered across success, failure, changed input, and recovery. |
| `internal/mcp/tools.go::Server.runCypherRead` | current, non-stale; host risk HIGH | LOW; 0 impacted/direct/modules/processes | Near read/failure owner covered directly through the built MCP command. |
| `scripts/full-build.ps1` (harness convention owner) | current, non-stale | LOW; 0 impacted | Supports a repo-native reusable PowerShell QA harness. |

The graph inspection reported zero degraded nodes and zero active resolution gaps for these boundary lookups. HIGH/CRITICAL were treated as blast-radius warnings, not edit bans; this lane remained read-only for production code.

## Source trace and backend classification

### C04 analyze/persistence

- `Run` resolves/acquires the repository storage boundary and fails immediately if storage preparation fails.
- It loads the current graph into Ladybug first. A load or close error returns before the Graph JSON write.
- When enabled by the normal command, Graph JSON is written through `<graph path>.tmp`, buffered output is flushed, the file is closed, and `os.Rename` atomically selects the completed file as `graph.json`.
- Any Graph JSON write/rename error returns from `Run`.
- Normal CLI dispatch calls `analyze.Run`; only after it returns successfully does command dispatch save metadata, register the repository, generate AI context, and print successful output. Thus the injected analyze fault is a real non-success rather than hidden registration/success.

### C17 read/failure

- `runCypherRead` opens `<same repo storage>/lbug` first.
- A successful open/query returns Ladybug rows immediately.
- A query error or any open error other than exact `errors.Is(err, lbugnative.ErrUnavailable)` returns non-success.
- Only exact `ErrUnavailable` authorizes fallback to `<same repo storage>/graph.json`; Graph JSON load or unsupported-query errors return non-success.
- The normal packaged binary is built with the Ladybug native tag/runtime. Current source exposes no truthful normal-runtime switch that forces only the `ErrUnavailable` sentinel. Therefore this run source-classifies the supported fallback but does not fabricate it using an untagged/dev binary. Native success and no-backend native failure were both measured with the fresh normal binary.

This distinguishes:

- current native artifact: success with `graph.json` absent throughout the built MCP invocation;
- supported fallback: only the exact source-authorized sentinel and the same repository's `graph.json`;
- forbidden stale/substitute success: not observed;
- external/unowned availability: no VECTOR dependency was required or retried.

## Reusable harness and isolated fixture

- Harness: `scripts/qa-child02-p2d-repeated-analyze.ps1`.
- Modes exercised: `Prepare`, `Build`, and `Run`.
- Harness SHA-256: `2CF8A230154029451343333CB206D0C1444D90FA988460EDB719CE34980733C7`.
- Execution-only lane root: `E:\Anvien\.tmp\qa-child02-p2d`.
- Same isolated repository path for every matched analyze sequence: `E:\Anvien\.tmp\qa-child02-p2d\fixture`.
- Isolated registry/home: `E:\Anvien\.tmp\qa-child02-p2d\anvien-home`.
- Ownership marker during execution: `owner.json`, `scope=child02-p2d`, `owner=independent-qa`.
- Minimal fixture source was copied from accepted Child 01 source `internal/resolution/testdata/p1c_identity_repo/src/identity.ts`.
- Accepted/initial/recovered fixture SHA-256: `327F6BD02FCB341078902FE4302CAB03A2077DDFE98B379164C17E96B547A078`.

No browser/UI run was added: C04 and C17 were completely measurable at nearer normal built CLI/MCP boundaries.

## Clean-holder full build

Before the only build, the harness used Windows Restart Manager and exclusive opens over all six relevant artifacts:

- repository `anvien/bin/anvien.exe` and `anvien/bin/lbug_shared.dll`;
- global installed `anvien.exe` and `lbug_shared.dll`;
- `anvien-launcher/AnvienLauncher.exe`;
- `anvien-launcher/server-bundle/anvien-server.exe`.

Restart Manager found only two verified holders: global `anvien.exe mcp` PIDs `13464` and `24512`. Their parent PIDs were recorded (`7776`, `24020`), but no parent was classified as a verified launcher holder, so only PIDs `13464` and `24512` were killed. The final gate recorded:

- `restartManagerZero=true` for all six artifacts;
- `exclusiveOpenAll=true` for all six artifacts;
- verified launcher parents killed: none;
- unrelated processes killed: none.

The root build then ran exactly `npm run full-build`:

- start: `2026-08-13T14:31:26.1004704Z`;
- end: `2026-08-13T14:33:36.9349842Z`;
- duration: `130.822s`;
- exit code: `0`.

Fresh normal CLI used for the matrix:

- path: `E:\Anvien\anvien\bin\anvien.exe`;
- version: `1.2.8`;
- SHA-256: `C77D3E1F2C13BDC0BB3F089D476E9B340385E0385CA4F7DF7689D7ECFDD40BFB`;
- bytes: `67040768`.

There was no build retry. Complete holder, process, command output, warning, artifact, timestamp, and hash records are embedded in the official JSON evidence.

## Validation matrix

| ID | Scenario | Verdict | Evidence and invariant proved |
|---|---|---|---|
| M1 | Two unchanged normal built analyze invocations | PASS | Both used the fresh binary and the same fixture path; exits `0/0`; equal input manifest `BBA28BFE274C863E8DAFA1DBF961F2E13ED1B3DBAE3D825952EFC14EFD016E7B`; equal accepted fact signature `FCE1E085BFE0352FF219C2FC33E5B5DFBEE2D49A3B73EAFE89CF6490CFF1BAAD`; definitions `10/10`, DEFINES `10/10`, missing endpoints `0/0`; exact IDs and ranges retained. |
| M2 | Changed source, next analyze, and current read | PASS | Changed input manifest `590BC39FCE5841B3D148E24C420902D681BDB9BF5A509704475DAFE9DC0BE3F8`; exit `0`; current fact signature `20C6FF5CDBEBC0D25E05603FD625E9C445EBC8D4449EFAA48B48DD17B836AEC6`; names `time=2`, `now=1`, `later=1`; the old second-`now` identity was absent. The immediate built MCP read returned `later, now, time, time`. |
| M3 | Clear analyze-command failure | PASS | Replaced the exact owned `.anvien` storage path with a regular file; normal built analyze exited `1`; stderr: `mkdir E:\Anvien\.tmp\qa-child02-p2d\fixture\.anvien: The system cannot find the path specified.`; expected success `graph.json` path was absent. No hidden success was counted. |
| M4 | Ladybug/native availability at C17 | PASS | Removed `graph.json` for the full built MCP invocation while changed `lbug` remained; process exit `0`, JSON-RPC result rows `later, now, time, time`. With Graph JSON unavailable, success is attributable to the current Ladybug artifact, not fallback or a stale substitute. |
| M5 | No readable backend | PASS | Held both the same repo's `lbug` and `graph.json` absent for the full call; MCP host process exited normally (`0`) but the requested JSON-RPC operation returned error `-32602`, `lbug_database_init failed with state 1`, with no result rows. This is a clear graph-bound non-success and no cross-artifact success. |
| M6 | Subsequent analyze/read recovery | PASS | Restored the baseline source/storage, ran the same normal built analyze successfully (`0`), restored the original accepted fact signature, then read `now, now, time, time` successfully through C17. |
| M7 | Accepted corrected identity/range/endpoint facts | PASS | Compared exact IDs, construct ranges, SelectionRanges, and DEFINES endpoints, rather than aggregate graph totals. `time` start lines were `5,10`; `now` start lines were `15,20`; definitions/DEFINES were `10/10`; missing endpoints were zero. |

Matrix total: **7 PASS / 0 FAIL / 0 BLOCKED**.

## Invocation-to-artifact correlation and determinism

All normal analyze commands were the fresh built executable with arguments equivalent to:

`E:\Anvien\anvien\bin\anvien.exe analyze E:\Anvien\.tmp\qa-child02-p2d\fixture --force --json --name qa-child02-p2d`

| Invocation | UTC interval | Input/fact state | `graph.json` SHA-256 | Ladybug SHA-256 |
|---|---|---|---|---|
| unchanged 1 | `14:36:43.366`–`14:36:46.734` | baseline / accepted fact signature | `C15E8075FC1FCA1849D5798CCC5DEF948E4C5210E9BFD62A35A23D4D623DFD1A` | `0631AB63590D991C1C9C08EDD0F7D93064A7E7F6B14044596775FF4610F8E38C` |
| unchanged 2 | `14:36:46.951`–`14:36:50.167` | same baseline / same accepted fact signature | `C15E8075FC1FCA1849D5798CCC5DEF948E4C5210E9BFD62A35A23D4D623DFD1A` | `F7D338B9BE3E388411EA40A9FCA72D606733FC3941071D887F595EAC2ABF947F` |
| changed | `14:36:50.377`–`14:36:53.636` | changed manifest / changed fact signature | `20395C774D56781197FAD25779F9B6AA2F36AF28767D4167C81731A95B9AF5C1` | `3AEA03E8913973607F39E5D9AD6583996297ABE6FF817920D32D6697327EA539` |
| recovery | `14:36:54.594`–`14:36:57.945` | baseline restored / accepted fact signature restored | `C15E8075FC1FCA1849D5798CCC5DEF948E4C5210E9BFD62A35A23D4D623DFD1A` | `C490D0F67923DA14BBD60FA847F9FB7E27C2AF0ED4EBC4127544DFDDB2B7371C` |

Each artifact timestamp fell inside its invocation interval. The changed artifact hashes/facts differed from the unchanged state and the recovery facts returned to baseline. Ladybug whole-artifact hashes legitimately differed across unchanged runs; byte identity of the whole backend was not treated as the contract. Determinism was judged using the source-backed canonical fact manifest, identities, construct/selection ranges, and endpoint integrity.

## Fault lifecycle and containment

- Analyze-failure injection touched only the lane-owned fixture storage boundary.
- Native-read/no-backend injections held only the fixture's current `graph.json`/`lbug` and restored them before recovery.
- Evidence recorded `faultArtifactsRestored=true` and `heldPathsRemaining=[]` before cleanup.
- Fixture source recovered to SHA-256 `327F6BD02FCB341078902FE4302CAB03A2077DDFE98B379164C17E96B547A078`.
- Final containment check found 0 lane-owned runtime processes and 0 listeners.
- The execution-only tree `E:\Anvien\.tmp\qa-child02-p2d` was removed after its build/matrix records had been embedded in official evidence. Four read-only Git objects required removal of their read-only file attribute inside this exact owned subtree; no shared path was affected.
- Shared/ownership-unknown `.tmp/ladybug-home`, `.tmp/ladybug-native`, and `.tmp/runtime-p2c` remained present and untouched by cleanup.
- No target source/worktree was contaminated.

## Official evidence

- JSON: `reports/QA/child02-p2d-repeated-analyze/qa_child02_p2d_repeated_analyze_260813_213658.json`
  - SHA-256: `FEE128775578D1F267FAD124295BB854ACD53DBC4A098B42AEC76FCB54EACC57`
  - Contains full build holder gate, build/runtime outputs and exits, invocation manifests, artifact hashes/timestamps, exact fact records, JSON-RPC payloads, failure/restoration details, and assertions.
- Markdown: `reports/QA/child02-p2d-repeated-analyze/qa_child02_p2d_repeated_analyze_260813_213658.md`
  - SHA-256: `8044924CE0F641EA6CA369DCEA4408B619221E3BE6F7FE38D1932873D1AAD41D`.

## Final containment and Git state

Final checks performed after execution-tree cleanup:

- `git diff --check`: clean result.
- branch: `master`.
- HEAD: `927a676653963e8001d7789291010d5b819bac83`.
- staged paths: `0`.
- expected dirty paths only:
  - `scripts/qa-child02-p2d-repeated-analyze.ps1`;
  - `reports/QA/child02-p2d-repeated-analyze/qa_child02_p2d_repeated_analyze_260813_213658.json`;
  - `reports/QA/child02-p2d-repeated-analyze/qa_child02_p2d_repeated_analyze_260813_213658.md`;
  - this report.
- lane-owned processes/listeners: `0/0`.
- `.tmp/qa-child02-p2d`: absent.

No files were staged or committed. Per the assigned QA boundary, `anvien detect-changes` was not run.

## QA conclusion and handoff

Measured current behavior satisfies every assigned P2-D invariant: repeated unchanged normal built analyzes preserve accepted facts; changed input is immediately current and stale identity is excluded; analyze and graph-bound faults fail clearly; no-backend cannot return substitute success; and subsequent normal analyze/read recovers truthfully. The supported Graph JSON fallback is narrowly classified from current source and was not falsely represented as a dynamically executed normal-runtime case.

**Status: READY_FOR_SUPERVISOR.** Orchestration main should hand this report and the paired evidence to the independent Supervisor. This QA lane does not call P2-D accepted and does not authorize P2-E.
