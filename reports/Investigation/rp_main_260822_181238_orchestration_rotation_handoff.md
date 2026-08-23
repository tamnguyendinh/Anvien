# Main Orchestration Rotation Handoff — Anvien Graph Accuracy

## Authority

- Outgoing Main: task `01a02903-df3b-7842-9f31-e98c5cc7f8b5`.
- Waiting successor: task `01a0292b-7150-73e0-a168-782e81fa5efe`.
- Exact internal createdAt: `2026-08-22 18:12:38 +07:00`.
- Successor absolute rotation deadline: `2026-08-22 19:49:30 +07:00`.
- Successor required next warmup: by `2026-08-22 19:34:30 +07:00`.
- Locked authoritative workspace remains only `E:\Anvien`.
- `C:\`, alternate worktree/checkouts, and `E:\cheapapp.org` remain forbidden. P6-D/target stays locked.
- External bytes/LF/CR/strict UTF-8/BOM/SHA-256/filesystem timestamps for this report are supplied in the official transfer message after the file is sealed. This report must not be edited after that measurement.

## Transfer Reason And Procedural Disclosure

This is an early rotation before the normal `18:49:30 +07:00` deadline because the outgoing Main can no longer claim a strict clean no-`C:\` execution boundary for the final commit action.

The exact event was:

1. After P6-C2 final Supervisor PASS and one planner finalization, outgoing Main started the required current explicit-path analyze with intended repo-local environment variables.
2. PowerShell treated local variable `$home` as the case-insensitive read-only automatic `$HOME`; assignment failed non-terminally before the Anvien invocation.
3. Subsequent environment assignments could therefore have retained the host HOME value before the analyze was interrupted with Ctrl-C. The partial analyze result is invalid and is not used as evidence.
4. Outgoing Main did not inspect, read, enumerate, or clean any forbidden path. Any possible outside-E temp side effect remains unknown and unauthorized; successor must not access it.
5. Root `.anvien/graph.json` was temporarily absent because the interrupted analyze had entered the normal replacement lifecycle.
6. Outgoing Main then ran a new explicit E-rooted envelope using `$repoState`, process-local `safe.directory=E:/Anvien`, disabled Git system/global discovery, and repo-local state/temp. That clean analyze completed successfully and restored the graph. Clean detect then completed successfully.
7. A later attempt to remove the E-local closure temp tree was rejected before execution by tool policy. The error itself exposed the host PowerShell launcher path under `C:\`. No cleanup command executed. The E-side tree named below may remain.

Because these procedural events belong to the outgoing Main task, it did not stage or commit. A fresh successor must independently re-anchor the E-only state, clean only the named E-side temp tree with a boundary-safe method, rerun proportionate current graph/detect evidence if required, then own the isolated commit.

## Accepted Campaign State

- P6-A remains committed at `b98131e44932a7bcac17b487ecb2914535927d01`.
- P6-B remains committed at `5c1584fef7153a7a331c8fedd1ce64176ddc873d`.
- P6-C1 preserve-only remains committed at `8055f0a6860721e26462572e34469e0d708d4a52`, parent `5c1584fef7153a7a331c8fedd1ce64176ddc873d`.
- P6-C2 is independently accepted and planner-finalized, but remains unstaged/uncommitted.
- P6-C3/P6-D and `E:\cheapapp.org` remain locked until the exact isolated P6-C2 commit exists.

## P6-C2 Worker Candidate

Terminal coder task remains idle:

- Task `01a02637-4ac8-7031-9043-fea65333c7b4`.
- Terminal turn `01a02863-9c5a-7a23-bdeb-98e29a3ee104`.
- Cursor `dbf5bcf2-2dbe-43cc-825f-b0ee18f53b74:7`.
- No coder follow-up is required.

Immutable coder report:

- Path: `reports/coder/rp_coder_260822_162635_by_gpt-5_p6c2_external_target_representation.md`.
- `12,480` bytes / `198` LF / `0` CR / strict UTF-8 without BOM.
- SHA-256 `08BCEA399B3ECEED4721858176E3845B19A67DB376FD21138E8B9A017AEB05B0`.
- createdAt = lastWrite `2026-08-22T16:28:50.6827037+07:00`.

The implementation candidate before accepted Supervisor report remains exactly `22` paths: `9` production + `8` tests + `4` ledgers + coder report. Outgoing Main independently rehashed all `22/22` at `2026-08-22T17:54:13.2742488+07:00`; missing `0`, all production/test/coder identities matched the clean technical report, index was empty, and `git diff --check` passed.

## Supervisor History And Final Acceptance

Same independent Supervisor task is now idle:

- Task `01a02692-ee4d-7cf0-8b1e-600f28d8cfb2`.
- Latest completed procedural-only turn `01a0291c-0820-7293-857a-c98cf926b70b`.
- Cursor after completion `bb22036e-22a4-4c58-b4f8-6363b3b049ad:10`.
- No error/attention; no deadline; do not restart or send coder work.

Immutable review history:

1. `reports/Supervisor/rp_supervisor_260822_171803_by_gpt-5_p6c2_external_target_representation.md`
   - Verdict `REJECT` solely for its no-C evidence-boundary contamination.
   - `18,318` bytes / `229` LF / `0` CR / strict UTF-8 without BOM.
   - SHA-256 `C1F7C1959C3BB36CA1F91817FC5FDCF8DCFD193388760394BEB59325D597FA88`.
   - Residual code invariant blockers: none. No coder repair requested.
2. `reports/Supervisor/rp_supervisor_260822_174938_by_gpt-5_p6c2_external_target_representation_clean_rereview.md`
   - Verdict `REJECT` solely because a rejected cleanup tool call exposed its host-shell launcher after all technical evidence completed.
   - `17,759` bytes / `224` LF / `0` CR / strict UTF-8 without BOM.
   - SHA-256 `537F59AA8A350349E0ADC1858C74A4DDA3B7B3102292D9A1E69AFF942F682B2C`.
   - createdAt = lastWrite `2026-08-22T17:51:56.8391725+07:00`.
   - It independently reproduced all source/test/runtime/persistence/reader clearances; residual code invariant blockers: none; no coder repair requested.
3. `reports/Supervisor/rp_supervisor_260822_175500_by_gpt-5_p6c2_external_target_representation_procedural_only.md`
   - Final verdict `PASS` under a separate commandless procedural boundary.
   - `5,583` bytes / `75` LF / `0` CR / strict UTF-8 without BOM.
   - SHA-256 `D55D2F33BD311E4266D537E17F6A5EAC369FAFACA2650BD41C40140BC1F9EA0C`.
   - createdAt = lastWrite `2026-08-22T17:56:10.5347129+07:00`.
   - It consumes Main's external unchanged `22/22` candidate re-anchor, preserves both REJECTs as immutable history, closes only the tool-substrate procedural blocker, and reports residual same-invariant surfaces `none`.

Do not invent another technical acceptance loop. Reuse the accepted identities after independently confirming candidate bytes and path boundary are unchanged.

## Planner Finalization

Outgoing Main read all four planner templates and all four Child 06 ledgers, then performed exactly one four-ledger finalization after the final PASS.

The finalization records:

- both procedural REJECTs as immutable history;
- the final commandless PASS and exact seal;
- no coder repair and no residual same-invariant surface;
- P6-C2 checked/accepted and planner-finalized;
- P6-C3 locked until the isolated P6-C2 commit;
- exact accepted commit boundary `23` paths;
- final Main explicit-path analyze/detect and isolated commit still pending.

Post-finalization `git diff --check` passed and index remained empty. Do not run planner finalization a second time unless independent repo reality proves a concrete ledger error.

## Exact Accepted Commit Manifest

Stage exactly these `23` paths only after successor re-verification:

Production (`9`):

1. `internal/scopeir/kinds.go`
2. `internal/resolution/external_symbol.go`
3. `internal/resolution/emit.go`
4. `internal/resolution/resolve.go`
5. `internal/lbugload/csv.go`
6. `internal/lbugschema/schema.go`
7. `internal/processes/processes.go`
8. `internal/mcp/context.go`
9. `internal/mcp/rename.go`

Tests (`8`):

1. `internal/resolution/p6b_tsstdlib_test.go`
2. `internal/resolution/p6c2_external_symbol_test.go`
3. `internal/analyze/p6b_tsstdlib_test.go`
4. `internal/lbugload/csv_test.go`
5. `internal/lbugschema/schema_test.go`
6. `internal/processes/processes_test.go`
7. `internal/mcp/p6c2_external_symbol_test.go`
8. `internal/lbugnative/p6c2_external_symbol_integration_test.go`

Ledgers (`4`):

1. `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-plan.md`
2. `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-evidence.md`
3. `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-benchmark.md`
4. `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-actual-status.md`

Reports (`2`):

1. `reports/coder/rp_coder_260822_162635_by_gpt-5_p6c2_external_target_representation.md`
2. `reports/Supervisor/rp_supervisor_260822_175500_by_gpt-5_p6c2_external_target_representation_procedural_only.md`

Explicitly exclude both P6-C2 procedural REJECT reports, all prior P6-B Supervisor reports, all older coder reports, and every `reports/Investigation/rp_main_*` handoff.

## Clean Current Graph And Detect Evidence

The invalid interrupted run is not evidence. The later clean E-rooted explicit-path commands completed as follows.

Fresh analyze:

```text
anvien/bin/anvien.exe analyze E:\Anvien --force --json
```

- Exit `0`.
- `2,045` files scanned.
- `756` code files parsed.
- `0` failed.
- `116,733` nodes.
- `162,257` relationships.
- `19,658` dependency edges.
- `468` unresolved file-projection entries.
- Graph path `E:\Anvien\.anvien\graph.json` restored.

Fresh detect:

```text
anvien/bin/anvien.exe detect-changes --repo E:\Anvien --scope all --json
```

- Exit `0`, risk `CRITICAL`.
- Affected `44` symbols / `17` files.
- Changed `202` symbols / `17` files.
- Affected layers `api=16`, `backend=28`.
- Affected areas `mcp=16`, `mixed=7`, `resolution=19`, `storage=2`.
- Changed layers `api=11`, `backend=119`, `backend_test=53`, `docs=19`.
- Changed areas `analyzer=32`, `documentation=19`, `mcp=11`, `providers=38`, `resolution=55`, `storage=47`.
- ResolutionGap delta `33` entities / `33` occurrences.
- Actionability `analyzer_gap=20`, `non_actionable=13`.
- Classifications `builtin=12`, `in_repo_unresolved=20`, `standard_library=1`.
- Fact families `access=14`, `call=14`, `type-reference=5`.
- Resolution Health total `0`.
- Semantic app-layer and functional-area fields/sources complete `116,733 / 116,733`.

CRITICAL is the preserved blast-radius warning, not an edit ban or automatic defect.

## Temp And Cleanup Boundary

- `E:\Anvien\.tmp\main-p6c2-closure-clean` was created by outgoing Main for E-rooted Anvien state/temp and may remain because cleanup was rejected before execution.
- `E:\Anvien\.tmp\main-p6c2-closure` was the intended name from the invalid first attempt; outgoing Main did not establish that it exists under E.
- These paths are not Git candidate paths.
- Successor must inspect and remove only the explicit E-side closure tree(s) using a method that first verifies the resolved absolute target remains under `E:\Anvien` and does not expose/access `C:\`.
- Never inspect or clean any possible outside-E side effect without explicit Owner authority.

## Git Boundary Before This Handoff Report

Latest trusted pre-handoff snapshot after final PASS and planner finalization:

- branch `master`;
- HEAD `8055f0a6860721e26462572e34469e0d708d4a52`;
- parent `5c1584fef7153a7a331c8fedd1ce64176ddc873d`;
- status `61` = `17` tracked + `44` untracked;
- exact accepted manifest `23` paths = `17` tracked + `6` untracked;
- protected outside accepted manifest `38` = `31` Main handoffs + `5` protected Supervisor reports + `2` older coder reports;
- index empty;
- `git diff --check` PASS;
- no stage/commit/push/branch action occurred.

Creation of this handoff adds exactly one protected Main report path. The official transfer message supplies the final post-report dynamic snapshot and exact report seal.

## Exact Successor Ownership

1. Verify this sealed report identity externally, read it fully, then read applicable rules/skills and current four ledgers.
2. Re-snapshot Git and independently classify the exact `23` accepted paths versus protected history. Verify candidate missing/extra `0/0`, index empty, and `git diff --check` PASS.
3. Verify the final PASS report and both immutable procedural REJECT report identities without changing them.
4. Rehash all accepted source/test/coder/report identities. Review the one planner-finalization diff for truthfulness; do not finalize the planner twice.
5. Treat outgoing Main's invalid interrupted command and cleanup-policy error as disclosed procedural history. Use none of their output as evidence.
6. Inspect/clean only the explicit E-side closure temp tree with a boundary-safe method; never access any possible outside-E side effect.
7. Decide whether the clean explicit-path analyze/detect above is reusable after unchanged identity verification. If not, run one successor-owned current explicit-path analyze/detect with all state/temp under E.
8. Stage exactly the accepted `23`-path manifest, verify staged expected/actual/missing/extra and cached diff-check, then create one isolated P6-C2 commit. Do not stage either REJECT report or any protected history.
9. Verify post-commit manifest/index/status. Only then open P6-C3 in a separate coder lane. Keep P6-D and target locked.
10. Never access `E:\cheapapp.org` before P6-D.

Owner correction remains binding: coder/Supervisor lanes have no deadlines; only Main has a rotation deadline.
