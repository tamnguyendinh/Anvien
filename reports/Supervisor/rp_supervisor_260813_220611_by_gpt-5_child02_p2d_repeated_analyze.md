# Supervisor Report: Child 02 P2-D repeated normal built analyze

Verdict: PASS

Disposition: `ACCEPT_P2D_AND_HAND_BACK`

## Metadata

- Report file: `reports/Supervisor/rp_supervisor_260813_220611_by_gpt-5_child02_p2d_repeated_analyze.md`
- Review time: `2026-08-13 22:06:11 +07:00` (`SE Asia Standard Time`)
- Reviewer: `gpt-5`
- Repository: `E:\Anvien`
- Scope reviewed: Child 02 slice P2-D only — repeated normal built analyze and clear failure behavior at C04 and C17
- Claim reviewed: repeated normal built analyze on one repository path exposes the current corrected facts deterministically at the semantic level; changed input replaces stale facts; owned analyze/read failures are clear and do not return substitute success; recovery restores readable accepted facts
- Authority: Owner delegation; `AGENTS.md`; `working-rules`; `supervisor`; `Data-Integrity`; `Edge-Case`; graph-accuracy contract; graph-accuracy roadmap; all four Child 02 ledgers; current source; current Git boundary; P2-D harness and QA evidence
- Review mode: zero-trust, independent, review-only

## Executive Summary

- Decision: unconditional `PASS` for P2-D only.
- P2-D is validation-only at the reviewed HEAD. No production or product-test repair is required.
- The official matrix is independently cleared as `7 PASS / 0 FAIL / 0 BLOCKED` after checking current source, raw JSON evidence, artifact hashes, invocation/artifact timestamps, exact accepted fact records, exact `DEFINES` pairs, native-read behavior, failure behavior, and recovery.
- Exact `ErrUnavailable -> same-repository graph.json` execution is not an acceptance blocker. In the normal packaged production binary, `ladybugdb` is compiled in and absence/corruption of its artifact returns a native operational error, not the `ErrUnavailable` build-capability sentinel. Manufacturing that sentinel would require an untagged/dev build or injection, both outside the normal-built-command authority and explicitly forbidden by the assigned boundary. The source-authorized branch is classified; all truthful production-binary outcomes needed by P2-D are dynamically evidenced.
- Required outcome: orchestration/main may accept P2-D, update only the roadmap/four Child 02 ledgers with the exact accepted evidence, run main-owned detect-changes, and commit the isolated P2-D boundary. P2-E must remain closed until that commit succeeds.

## Git and Worktree Boundary

- Branch: `master` — verified.
- HEAD: `927a676653963e8001d7789291010d5b819bac83` — verified exactly; subject `fix(readers): consume corrected graph facts`.
- HEAD contains the accepted P2-C boundary; tracked source diff at review time: zero.
- Staged paths before this report: `0`.
- Dirty boundary before this report: exactly four untracked paths:
  1. `scripts/qa-child02-p2d-repeated-analyze.ps1`
  2. `reports/QA/child02-p2d-repeated-analyze/qa_child02_p2d_repeated_analyze_260813_213658.json`
  3. `reports/QA/child02-p2d-repeated-analyze/qa_child02_p2d_repeated_analyze_260813_213658.md`
  4. `reports/QA/rp_qa_260813_214244_by_gpt-5_child02_p2d_repeated_analyze.md`
- `git diff --check`: clean.
- No file was staged, committed, pushed, cleaned up, or switched by this Supervisor lane.
- This report is the only new path introduced by this lane.

## Authority and Files Inspected

Read in full:

- `E:\Anvien\AGENTS.md`
- `E:\Anvien\.agents\skills\working-rules\SKILL.md`
- `E:\Anvien\.agents\skills\supervisor\SKILL.md`
- `E:\Anvien\.agents\skills\Data-Integrity\SKILL.md`
- `E:\Anvien\.agents\skills\Edge-Case\SKILL.md`
- `docs/contracts/graph-accuracy-contract.md`
- `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-26-anvien-graph-accuracy-roadmap.md`
- all four ledgers under `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-02-current-graph-persistence-and-reader-consistency/`
- `scripts/qa-child02-p2d-repeated-analyze.ps1`
- the official P2-D JSON and Markdown evidence
- the P2-D QA report

Direct source/contract surfaces inspected:

- `internal/analyze/analyze.go::{Run,prepareStorage,writeGraphSnapshot,writeGraphSnapshotJSON}`
- `internal/cli/command.go` normal analyze dispatch
- `internal/cli/analyze_postrun.go::{recordAnalyzeResult,generateAnalyzeAIContext}`
- `internal/mcp/tools.go::{cypherTool,Server.runCypherRead}`
- `internal/mcp/server.go::{Config,NewServer}`
- `internal/mcp/resources.go::storagePathForEntry`
- `internal/lbugnative/runner.go::ErrUnavailable`
- `internal/lbugnative/runner_default.go`
- `internal/lbugnative/native_ladybugdb.go::openNativeDatabase`
- `internal/cli/package_runtime.go` packaged runtime build tags
- directly relevant existing tests in `internal/analyze/analyze_test.go`, `internal/analyze/p2b_snapshot_persistence_test.go`, `internal/mcp/server_test.go`, and `internal/lbugnative/runner_default_test.go`

Accepted P2-B/P2-C reports were not rerun because current HEAD, tracked source, binary hash, and evidence boundary do not invalidate their predecessor invariants.

## Evidence and Hash Verification

All delegated hashes match the files present in the current worktree:

| Artifact | Verified SHA-256 | Result |
|---|---|---|
| reusable harness | `2CF8A230154029451343333CB206D0C1444D90FA988460EDB719CE34980733C7` | exact match |
| official JSON | `FEE128775578D1F267FAD124295BB854ACD53DBC4A098B42AEC76FCB54EACC57` | exact match |
| official Markdown | `8044924CE0F641EA6CA369DCEA4408B619221E3BE6F7FE38D1932873D1AAD41D` | exact match |
| QA report | `191222C39936DCABAF41DB1F54B96B5273408CCE94B451E45DC07894463E5E7D` | exact match |
| fresh normal binary | `C77D3E1F2C13BDC0BB3F089D476E9B340385E0385CA4F7DF7689D7ECFDD40BFB` | exact match at review time |

The binary is currently `67,040,768` bytes and reports version `1.2.8`. Its timestamp and hash match the build artifact row embedded in the JSON evidence.

The clean-holder gate is complete and internally correlated:

- six build artifacts were registered with Windows Restart Manager;
- only verified `anvien.exe mcp` holders PIDs `13464` and `24512` were killed;
- their parent PIDs `7776` and `24020` were recorded and not classified as verified launchers;
- no launcher parent and no unrelated process was killed;
- every artifact ended at zero Restart Manager holders;
- exclusive-open succeeded for all six artifacts before build;
- `npm run full-build` ran once, exited `0`, and took `130.822s` from `2026-08-13T14:31:26.1004704Z` to `2026-08-13T14:33:36.9349842Z`;
- no build retry occurred.

## Evidence Target Disposition

| Evidence target | Verdict | Independent Supervisor basis |
|---|---|---|
| `E2-P2D-IMPACT1` | PASS | Current QA graph/file-detail/impact evidence is tied to the reviewed HEAD and is not invalidated by tracked drift. C04 host/source blast radius is HIGH/CRITICAL; C17 exact symbol impact is LOW. These are scope warnings. No production edit exists in P2-D. |
| `E2-P2D-SOURCE1` | PASS | Current source independently proves C04 atomic Graph JSON selection and truthful dispatch, plus C17 Ladybug-primary behavior, exact-sentinel fallback, same-repository path selection, and non-sentinel fail-closed behavior. |
| `E2-P2D-BUILD1` | PASS | Raw JSON proves holder/launcher gate, zero holders, exclusive opens, one full build exit `0`, runtime version/hash, and post-build artifact inventory. Current binary still matches. |
| `E2-P2D-TEST1` | PASS | Existing behavior did not fail, so the plan does not require production test edits. The reusable harness executes the required repeat/change/failure/read/recovery behavior matrix through the fresh normal CLI/MCP boundaries; all seven rows pass. |
| `E2-P2D-REPEAT1` | PASS | Raw invocation, input, artifact, exact fact, fault, JSON-RPC, and recovery records close the repeated-analyze and clear-failure matrix. |
| `E2-P2D-REVIEW1` | PASS | This independent report is the durable unconditional Supervisor acceptance for P2-D. |

`E2-P2D-DETECT1` and `E2-P2D-COMMIT1` remain intentionally main-owned and are not part of this lane's verdict artifact creation.

## Runtime Matrix and Invariant Reasoning

| ID | Supervisor result | Closure reasoning |
|---|---|---|
| M1 unchanged x2 | PASS | Same fresh binary, same fixture path, equal input signatures, exits `0/0`, matching artifact timestamps inside each invocation, and exact equality of all 10 normalized Definition identity/range rows and all 10 normalized `DEFINES` pairs. |
| M2 changed input | PASS | Input signature changes; next successful graph signature changes; `later@20` appears; `now@20` and its old opaque identity disappear; immediate native read returns `later, now, time, time`. |
| M3 analyze storage failure | PASS | Replacing the exact lane-owned `.anvien` path with a file makes normal analyze exit `1` with a concrete `mkdir` error. The expected `graph.json` path is absent during failure; CLI registration/success output cannot run after `Run` returns error. |
| M4 native read, Graph JSON absent | PASS | `graph.json` is absent for the entire fresh packaged MCP call, current Ladybug exists, and the read succeeds with the changed facts. No Graph JSON or stale cross-artifact success is possible in this scenario. |
| M5 both backends absent | PASS | Both current same-repository artifacts are absent; the MCP process remains a valid JSON-RPC host but the requested operation returns `-32602`, native init error text, no `result`, and no rows. This is clear graph-bound non-success. |
| M6 recovery | PASS | Baseline source and artifacts are restored; normal analyze exits `0`; exact baseline Definition rows and exact baseline `DEFINES` pairs return; native read returns `now, now, time, time`. |
| M7 corrected facts | PASS | Exact opaque IDs, labels/names, construct ranges, SelectionRanges, 10/10 Definition-to-`DEFINES` pairs, and zero missing endpoints are checked. Counts are not used as the sole proof. |

Invocation-to-artifact correlation is valid: each `graph.json` and Ladybug write timestamp falls within its owning invocation interval. Baseline Graph JSON SHA-256 is stable across warmup, unchanged runs, and recovery; the changed run has a distinct Graph JSON hash. The fact signature is baseline-stable, changes for the controlled source edit, and returns exactly on recovery.

Semantic determinism is the correct acceptance unit. The two unchanged Ladybug whole-file hashes differ, but the authority forbids traversal/order/randomness from changing otherwise identical canonical facts; it does not require byte-for-byte identity of the entire Ladybug storage file. All accepted Definition identities/ranges and exact endpoint pairs are identical across unchanged runs and recovery. Therefore the whole-Ladybug hash difference is informational and non-blocking.

## ErrUnavailable Decision

Decision: absence of a dynamic exact-sentinel fallback run does **not** block P2-D acceptance.

Source facts:

1. `internal/lbugnative/runner_default.go` returns `ErrUnavailable` only under build constraint `!ladybugdb`.
2. `internal/lbugnative/native_ladybugdb.go` is selected under `ladybugdb`; native open failure returns an operational error such as `lbug_database_init failed with state 1`, not `ErrUnavailable`.
3. `internal/cli/package_runtime.go` builds the normal packaged binary with `go build -tags ladybugdb` and records `ladybugdb` in runtime metadata.
4. `Server.runCypherRead` returns native rows on success, returns query errors, returns every non-sentinel open error, and falls back only when `errors.Is(openErr, lbugnative.ErrUnavailable)` is true.
5. Both Ladybug and Graph JSON paths are derived from the same resolved registry entry/storage path.

Dynamic evidence present:

- successful current native read while `graph.json` is absent for the full call;
- clear non-success/no result while both backends are absent and the tagged native open returns a non-sentinel error;
- changed-input current read and baseline recovery read;
- analyze storage failure with no artifact at the normal success path.

Dynamic evidence absent:

- exact `ErrUnavailable -> same-repository graph.json` fallback through the fresh normal packaged production binary.

Why the set is sufficient:

- P2-D authority tests the normal built command and the mechanisms actually implemented by that build. In this build, native Ladybug is an available capability; exact `ErrUnavailable` is unreachable by deleting or corrupting a runtime artifact because those actions produce operational errors.
- The acceptance requirement is truthful success/current facts or clear failure with no substitute success. The evidence executes both truthful outcomes available to the production binary.
- Using an untagged/dev binary, injecting a fake opener, adding a test hook, monkey-patching, or fabricating an artifact would test a different command boundary and violate the assigned non-goals.
- Source classification is sufficient for the unreachable compile-time capability branch because no acceptance wording requires every internal conditional to be dynamically covered, and the accepted C17 inventory classifies the boundary as validate-only.
- `ErrUnavailable` is capability absence, not failure of the tagged build's current Ladybug artifact. Therefore the source-authorized Graph JSON path does not excuse or mask the native artifact failure dynamically tested by M5.

## Data-Integrity and Edge-Case Clearance

- Repeated/ordering: two identical analyzes preserve exact accepted facts and endpoints; changed input immediately replaces stale identity.
- Stale-state isolation: C17 reads the same registered repository storage; native success is proved with Graph JSON absent, and no-backend returns no result.
- Partial failure: analyze storage failure produces non-success before registration/success output and leaves no normal success artifact at the injected path.
- Recovery: held artifacts and source are restored, baseline facts return exactly, and no held fault path remains.
- Endpoint integrity: 10/10 exact `DEFINES` pairs across unchanged runs and recovery; zero missing endpoints.
- Cleanup/containment: `faultArtifactsRestored=true`, `heldPathsRemaining=[]`, the owned `.tmp/qa-child02-p2d` tree is absent, shared/unknown temp areas were not touched, and no target source/worktree contamination is reported or observed.
- No cross-repository or alternate-artifact success is present.

## Source-Level Clearance Notes

- `internal/analyze/analyze.go`: clear for P2-D. `prepareStorage` failures return; Ladybug load/close completes before Graph JSON; Graph JSON uses temp/write/flush/close/rename and returns every error.
- `internal/cli/command.go` and `internal/cli/analyze_postrun.go`: clear. Registration, metadata, AI context, and success output occur only after `analyze.Run` returns successfully.
- `internal/mcp/tools.go`: clear for the reviewed C17 contract. Native success is primary; query and non-sentinel open errors fail; exact capability sentinel alone selects same-repository Graph JSON.
- `internal/lbugnative/*` and `internal/cli/package_runtime.go`: clear. Build constraints explain why exact sentinel fallback cannot be honestly exercised through the normal tagged binary.
- P2-B persistence and P2-C readers: predecessor invariants remain accepted and unmodified at this HEAD; no rerun was needed.

## Findings

Blocking findings: none.

Non-blocking findings: none.

Evidence limitation, explicitly non-blocking: exact compile-time `ErrUnavailable` fallback is source-classified but not dynamically executed through the tagged production binary for the reasons above. This is not residual acceptance work for P2-D.

Residual unverified same-invariant surfaces: none within P2-D's C04/C17 normal-built-command scope.

## Not Run

- No second full build, QA matrix rerun, accepted P2-B/P2-C gate, graph/impact batch, or product test batch was run. Current raw evidence, matching binary, unchanged tracked source, and exact HEAD supplied sufficient proof; no concrete invalidation was found.
- No `anvien detect-changes` was run because the Owner assigned it to orchestration/main after the verdict.
- No exact-sentinel fallback was manufactured through an untagged/dev binary, source edit, hook, fake opener, monkey patch, or fake artifact.
- No browser/Playwright run was added because C04/C17 are non-UI and were exercised at their nearer real CLI/MCP boundaries.

## Exact Handoff to Orchestration/Main

1. Accept this unconditional P2-D `PASS` as `E2-P2D-REVIEW1`.
2. Update only the roadmap and four Child 02 ledgers to record the accepted P2-D facts, measurements, hashes, matrix, semantic determinism rule, and explicit non-blocking `ErrUnavailable` decision.
3. Preserve the harness, official JSON/Markdown, QA report, and this Supervisor report; do not alter or replace their evidence content.
4. Run main-owned `anvien detect-changes --repo Anvien --scope all` under current repository rules, record `E2-P2D-DETECT1`, stage only the exact P2-D boundary, and create the isolated P2-D commit as `E2-P2D-COMMIT1`.
5. Keep P2-E and every later Child closed until Git confirms that isolated commit succeeds. After confirmation, orchestration/main may open P2-E only.

## Overall Evaluation

P2-D's full outcome contract is closed at the current HEAD: repeated normal built analyzes are semantically deterministic for the accepted facts; changed input becomes current without stale identity; native reads use current Ladybug facts; owned analyze and graph-bound failures are clear and fail closed; and recovery restores exact baseline facts and endpoints. Evidence quality is sufficient and internally correlated. The unexecuted exact capability-sentinel branch is not a missing acceptance proof for the tagged normal production binary.
