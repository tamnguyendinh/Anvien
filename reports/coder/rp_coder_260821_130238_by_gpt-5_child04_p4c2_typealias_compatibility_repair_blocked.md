# Coder Report: Child 04 P4-C2 TypeAlias compatibility repair

Status: **BLOCKED**

## Metadata

- Report time: `2026-08-21 13:02:38 +07:00` (`Asia/Bangkok`).
- Role: bounded Coder rejection-repair lane.
- Repository: `E:\Anvien`.
- Current Git basis at evidence capture: branch `master`, HEAD `310502a88849fe75f86a45a987ba21490d19dbe2` (`docs(orchestration): fix skill name frontmatter`).
- Accepted predecessor: P4-C commit `c99c4070b66e7a96be8c9fa2721a0335a1f94877`.
- Rejection authority: `reports/Supervisor/rp_supervisor_260821_123030_by_gpt-5_child04_p4c2_durable_retry_review1.md`, canonical SHA-256 `8E37F4B126ABB38A5DCE071C35937F59D9152EFA135B9245408CA138BEC781F7`.
- Active scope: only the rejected direct exported TypeAlias definition compatibility and downstream FileContext invariant.
- Target `E:\cheapapp.org`: not accessed.
- Next owner: Main for blocker routing. This report is not a QA or acceptance verdict.

## Outcome

The exact production-first repair and focused test update are present in the allowed two-file boundary, but the mandatory canonical full build did not pass. Three build attempts failed while replacing `E:\Anvien\anvien\bin\anvien.exe`; Windows Restart Manager attributed a handle on that path to two editor-owned global Anvien MCP processes. Their command/path are not build-related, so lane authority explicitly prohibited terminating them.

Because full build is a prerequisite for testing, no focused test or nearest real resolution/FileContext/persistence boundary was run. The lane therefore returns `BLOCKED`, not `READY_FOR_QA`.

## Exact rejected invariant

- P001-P014 and P018 are direct exported TypeAlias definitions. Definition compatibility must be `isExported=true`, and production FileContext must consequently return `exported=true`.
- Their Export facts must remain `typeOnly=true` with `meanings=[type]`; definition compatibility must not synthesize runtime-value eligibility.
- Six Function positives, 11 negative controls, access/export separation, Graph JSON/Ladybug parity, zero orphan/diagnostic state, and zero Child 05 terminal/resolved/public-API state must remain unchanged.

## Invariant Family Map

- Family: local source export fact -> definition compatibility -> FileContext exported-symbol result.
- Source of truth: a `ScopeIR.ExportFact` carrying `LocalDefID` after `exportProjectionNodes` has rejected source re-export (`TargetRaw != nil`), missing definition, and cross-file definition states.
- Sibling surfaces: value/Function direct exports, type-only TypeAlias direct exports, unexported definitions, anonymous defaults, source re-exports, access visibility, Export-node type/meaning fields, Ladybug persistence, and FileContext canonical-field precedence.
- Forbidden fallback: using runtime-value eligibility to decide whether a local source definition is directly exported.
- Forbidden expansion: FileContext/Ladybug changes, terminal/barrel/alias-chain/cycle/ambiguity/public-API state, target access, or Child 05 work.
- Stale artifact: `TestP4CProjectsExportFactsAndRuntimeCompatibility` encoded the rejected TypeOnly compatibility expectation as false.

## Confirmed root cause and production repair

`internal/resolution/emit.go::exportProjectionNodes` was already inside the guarded local-definition branch, but it populated `directExportDefIDs` only when `exportFactIsRuntime(fact)` returned true. That helper rejects `TypeOnly` and type-only meaning facts, conflating runtime eligibility with source-export membership.

Production-first change:

- `internal/resolution/emit.go:374` now records every validated local export definition in `directExportDefIDs` independent of `TypeOnly`/meaning.
- The now-unused `exportFactIsRuntime` helper was removed; it had no other caller.
- `exportGraphNode` is unchanged, so `typeOnly`, `meanings`, source provenance, access separation, and absence of Child 05 fields remain source-level preserved.
- `internal/filecontext/context.go`, `internal/lbugload/csv.go`, and `internal/lbugschema/schema.go` remain byte-untouched and continue to consume/persist the canonical fields.

## Test-after-code repair

Only after the production edit, `internal/resolution/p4c_export_projection_test.go` was updated:

- the runtime-positive fixture now uses `Function`;
- the rejected type-only fixture now uses `TypeAlias`;
- direct TypeAlias definition compatibility expects `isExported=true`;
- the Export node must still expose `typeOnly=true` and exactly `meanings=[type]`;
- private/public visibility separation checks and zero later-slice field checks remain.

The test was not executed because the canonical build gate failed first.

## Exact candidate diff and identities

| Path | Diff | Bytes | LF | SHA-256 | Git blob |
|---|---:|---:|---:|---|---|
| `internal/resolution/emit.go` | `+1/-15` | `26,772` | `815` | `B1B8F22AEFE94F72483CD284B7E267FE3897412C04D581AFAB7F9FE30F0B2867` | `3c6ede9c93531a634db32b8b0100c38bde0ffaeb` |
| `internal/resolution/p4c_export_projection_test.go` | `+23/-6` | `9,247` | `203` | `F575F16D20D8F95C347E142BCBBEE96E18DA84D23E1433E00B6C2545132E0162` | `ee9076e20adea437222e3c2df8cc28e9ad61e0ae` |

- Total: `24` insertions / `21` deletions across exactly two tracked repair files.
- `gofmt -d` produced no output.
- `git diff --check -- <two repair files>` exited `0`.
- Before this lane edited them, both target files had zero working-tree diff and zero drift from accepted P4-C through then-current HEAD.

## Anvien freshness, ownership, and blast radius

- Repo-local `anvien --help`: PASS.
- Initial fresh self graph: exit `0`; scanned/parsed/failed `1,915/736/0`; `114,487` nodes / `157,336` relationships.
- A concurrent unrelated orchestration-skill commit advanced HEAD from graph basis `107d7ae9...` to `310502a8...`; exact diff was only `internal/aicontext/skills/orchestration/SKILL.md`, with zero target-file drift.
- Required candidate refresh then passed: `1,915/736/0`; `114,511` nodes / `157,326` relationships; indexed/current commit `310502a88849fe75f86a45a987ba21490d19dbe2`.
- `emit.go` file-detail: `150` symbols, `43` inbound, `226` outbound, `77` local relationships, `13` linked flows, `17` linked tests, file risk `HIGH`, non-stale.
- `emit.go` upstream file impact: `CRITICAL`; `30` impacted symbols, `24` direct, `5` affected files, `1` affected flow.
- `emitDefinitionNodes`: `CRITICAL`; `6` impacted symbols, `1` direct, `4` modules, `34` processes.
- `exportProjectionNodes`: `CRITICAL`; `4` impacted symbols, `1` direct, `2` modules, `26` processes.
- `exportFactIsRuntime`: `CRITICAL`; `3` impacted symbols, `1` direct, `1` module, `14` processes.
- Focused test file-detail: `39` symbols, zero inbound/unresolved, risk `LOW`; test function impact `LOW` with zero affected symbols/modules/processes.
- HIGH/CRITICAL results were treated as blast-radius warnings, not edit prohibitions; the candidate remained two files.

## Canonical full-build attempts

All three attempts reached the same failure in both the package postinstall build and prepare build paths.

| Attempt | Command | Exit | npm debug log |
|---:|---|---:|---|
| 1 | `npm run full-build` | `1` | `C:\Users\TAM NGUYEN\AppData\Local\npm-cache\_logs\2026-08-21T05_56_02_433Z-debug-0.log` |
| 2 | `npm run full-build` | `1` | `C:\Users\TAM NGUYEN\AppData\Local\npm-cache\_logs\2026-08-21T05_57_36_280Z-debug-0.log` |
| 3 | delayed diagnostic wrapper invoking canonical `npm run full-build` once | `1` | `C:\Users\TAM NGUYEN\AppData\Local\npm-cache\_logs\2026-08-21T05_59_07_460Z-debug-0.log` |

Exact failing destination:

```text
E:\Anvien\anvien\bin\anvien.exe
```

Exact repeated error shape:

```text
go build ... copying ...\a.out.exe: open E:\Anvien\anvien\bin\anvien.exe:
The process cannot access the file because it is being used by another process.
```

Each attempt ended with `npm install failed with exit code 1`. No source compile diagnostic was emitted before the file replacement failure.

## Lock/process evidence and bounded stop

- Pre-build Anvien analyze lock: `free`; lock file absent.
- `anvien doctor processes --json`: three Anvien MCP pairs, all global npm runtime and editor-owned.
- Verified build-process census: none.
- Repo-local executable process census immediately before retry: none.
- Exclusive no-share read-open immediately before retry: PASS, confirming the lock appeared during the build window.
- One bounded Windows Restart Manager read-only probe at `2026-08-21 13:01:37-13:01:38 +07:00` reported exact holders:
  - PID `11656`: `C:\Users\TAM NGUYEN\AppData\Roaming\npm\node_modules\anvien\bin\anvien.exe mcp`, parent PID `14616` / `cmd.exe`.
  - PID `14924`: `C:\Users\TAM NGUYEN\AppData\Roaming\npm\node_modules\anvien\bin\anvien.exe mcp`, parent PID `14120` / `cmd.exe`.
- Both are editor-owned MCP processes, not build-related holders. They were preserved and not terminated.
- A concurrent process-path monitor observed no process executing the repo-local binary during the diagnostic build.
- Per Main's bounded intervention, three repeated failures plus no terminable build-related holder end the diagnosis loop. No fourth build, package rebuild, extra probe, or audit is authorized.
- Disposable `.tmp/p4c2_lock_monitor.ps1` was deleted after the single Restart Manager probe.

## E2E Verification

```text
E2E Verification:
  [FAIL] Compiled: npm run full-build -> exit 1, repo-local anvien.exe file lock
  [NOT RUN] Runtime/nearest boundary: blocked by mandatory full-build gate
  [NOT RUN] Happy path: Function + direct TypeAlias compatibility projection
  [NOT RUN] Edge case: typeOnly/meaning preservation, negatives, access separation, persistence, Child 05 exclusion
```

No benchmark row is created: this repair changes no benchmarkable performance/capacity system, and build duration is validation evidence rather than a product benchmark.

## Worktree and boundary preservation

- Index is empty.
- Current tracked modifications are exactly the five pre-existing Child 04 living documents plus this lane's two repair files.
- Concurrent orchestration-skill work was committed by its owner and was not modified by this lane.
- All pre-existing untracked Oracle/QA/Supervisor/Architect/Planner/Main provenance remains preserved.
- This Coder report is the only durable report created by this lane.
- No edit occurred in FileContext, Ladybug CSV/schema/loader, ScopeIR/provider owners, other tests/goldens, plan/ledger documents, target files, or Child 05/later-slice owners.
- No target analyze/comparison, QA, Supervisor, Anvien detect-changes, stage, commit, push, reset, checkout, or broad cleanup occurred.
- No evidence-bearing artifact originated in `.tmp`.

## Residual unverified surfaces

- Canonical full build: **blocked**.
- Focused rejection regression: **not run**.
- Nearest real resolution -> FileContext -> persistence boundary: **not run**.
- Six target Function positives, 11 target negative controls, Graph JSON/Ladybug parity, and target zero-state: source-level paths were preserved but **not freshly revalidated by this Coder lane**.

Residual unverified surfaces are therefore not `none`; `READY_FOR_QA` is prohibited.

## Handoff

Status: **BLOCKED**.

Main must retain P4-C2 as open and Child 05 as locked. The exact external blocker is the non-terminable editor-owned MCP handle on `E:\Anvien\anvien\bin\anvien.exe`, which prevents the mandatory canonical full build from replacing the repo-local runtime. After Main coordinates a safe MCP/handle release under appropriate authority, the next Coder continuation should begin at the pre-build holder gate, run canonical `npm run full-build` once, then—only on PASS—run the focused rejection test and nearest resolution/FileContext/Ladybug boundary. It must not redo production/test edits unless their identities drift.

## Report identity

- Encoding/line ending: UTF-8 / LF.
- Canonical SHA-256 basis: replace the 64-character value after `canonical SHA-256:` with 64 ASCII zeroes.
- Canonical SHA-256: `D0190CBAC22716412867205309E7BA75FFD5A7690BC881FF74E9ACF38F5CF931`.
