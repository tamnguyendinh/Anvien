# Main Orchestration Rotation Handoff — Anvien Graph Accuracy

## Authority

- Outgoing Main: task `01a0292b-7150-73e0-a168-782e81fa5efe`.
- Waiting successor: task `01a0296e-3fab-7f53-babd-d819c3ac7190`.
- Exact internal createdAt: `2026-08-22 19:28:27 +07:00`.
- Successor absolute rotation deadline: `2026-08-22 20:49:30 +07:00`.
- Successor required next warmup: by `2026-08-22 20:34:30 +07:00`.
- Locked authoritative workspace remains only `E:\Anvien`.
- `C:\`, alternate worktree/checkouts, and `E:\cheapapp.org` remain forbidden. P6-D/target stays locked.
- External bytes/LF/CR/strict UTF-8/BOM/SHA-256/filesystem timestamps for this report are supplied in the official transfer message after the file is sealed. This report must not be edited after that measurement.

## Early Rotation Reason And Procedural Disclosure

This Main rotates early, before the absolute `19:49:30 +07:00` deadline, because its strict no-`C:\` observation boundary is no longer clean for any future P6-C3 acceptance/commit action.

The exact event occurred after P6-C2 was already independently accepted, staged, committed, and post-commit verified:

1. Main opened a visible same-directory coder fork for P6-C3. That first fork remained active for more than seven minutes without ACK, assistant item, or tool marker.
2. Main sent a scoped reminder and then an absolute STOP. The stuck task ended idle with no assistant output, no tool, and no repo action. It is not an active lane.
3. Main opened a clean project-local replacement coder task and used `codex_app__read_thread(includeOutputs=true)` once to verify its mandatory ACK and actual first command behavior.
4. The app response serialized command-execution metadata that included the host PowerShell launcher path under `C:\` even though the coder command itself targeted only explicit `E:\Anvien` files.
5. Main did not invoke, enumerate, inspect, hash, modify, or clean any `C:\` filesystem path. No repository mutation resulted from reading the thread metadata.
6. By the campaign's existing strict procedural precedent, exposure of a host launcher path is enough to prevent this Main from owning a later acceptance/commit boundary. Therefore it does not wait for P6-C3 completion and rotates to the already waiting clean successor.

This disclosure does not invalidate P6-C2: the event happened after its exact isolated commit and post-commit verification. It also does not establish a P6-C3 code blocker and must not be routed to the coder as a repair invariant. The successor must independently evaluate the eventual coder report and Supervisor evidence.

One later read-only Git snapshot script had a PowerShell parser error before execution because `$LASTEXITCODE` was embedded inside an object expression. Main immediately reran the read-only snapshot with the exit code separated; repo/index were unchanged. The failed invocation is not evidence.

## Accepted Campaign State

- P6-A remains committed at `b98131e44932a7bcac17b487ecb2914535927d01`.
- P6-B remains committed at `5c1584fef7153a7a331c8fedd1ce64176ddc873d`.
- P6-C1 preserve-only remains committed at `8055f0a6860721e26462572e34469e0d708d4a52`, parent `5c1584fef7153a7a331c8fedd1ce64176ddc873d`.
- P6-C2 is independently accepted and committed at `2acd357db32ac0c2bd3c3968a950c773acf3000d`, parent `8055f0a6860721e26462572e34469e0d708d4a52`.
- P6-C3 is the sole open slice and has an active visible coder lane.
- P6-D and `E:\cheapapp.org` remain locked.

## P6-C2 Closure Completed By This Main

This Main re-read the sealed incoming handoff, `AGENTS.md`, applicable `working-rules`/`orchestration`/`supervisor`/`planner` skills, all four planner templates, and the current Child 06 ledger sections.

Incoming sealed handoff identity was verified unchanged:

- Path: `reports/Investigation/rp_main_260822_181238_orchestration_rotation_handoff.md`.
- `13,602` bytes / `223` LF / `0` CR / strict UTF-8 without BOM.
- SHA-256 `DB125178C20D08BA072BD5CFAC6D2ED596897C77448756F87B6A0E988C4EA3E6`.
- createdAt/lastWrite `2026-08-22T18:14:27.8456173+07:00`.

P6-C2 final commandless PASS remains immutable:

- `reports/Supervisor/rp_supervisor_260822_175500_by_gpt-5_p6c2_external_target_representation_procedural_only.md`.
- `5,583` bytes / `75` LF / `0` CR / strict UTF-8 without BOM.
- SHA-256 `D55D2F33BD311E4266D537E17F6A5EAC369FAFACA2650BD41C40140BC1F9EA0C`.

Both earlier P6-C2 procedural REJECT reports remain protected immutable history outside the commit, with SHA-256 values `C1F7C1959C3BB36CA1F91817FC5FDCF8DCFD193388760394BEB59325D597FA88` and `537F59AA8A350349E0ADC1858C74A4DDA3B7B3102292D9A1E69AFF942F682B2C`. Neither established a code blocker or requested coder repair.

Successor-owned pre-stage re-anchor at `2026-08-22T18:25:06.7772180+07:00` verified:

- branch `master`, HEAD `8055f0a6860721e26462572e34469e0d708d4a52`;
- accepted manifest `23/23`, missing `0`, unexpected `0`;
- protected outside boundary `39 = 32 Main + 5 Supervisor + 2 coder`;
- empty index and `git diff --check` PASS;
- all known report and four-ledger identities unchanged;
- both explicit E-side closure temp trees absent.

The clean E-rooted post-finalization analyze/detect evidence from the prior Main was reused under the orchestration no-rerun rule because accepted identities and boundary were unchanged:

- analyze: `2,045` scanned / `756` parsed / `0` failed / `116,733` nodes / `162,257` relationships / `19,658` dependency edges / `468` unresolved file-projection entries;
- graph: `462,474,189` bytes;
- detect: exit `0`, risk CRITICAL, `44` affected symbols / `17` files, `202` changed symbols / `17` files, ResolutionGap `33/33`, Resolution Health `0`, semantic completeness `116,733/116,733`.

Exact staging then verified `23/23`, missing/extra `0/0`, cached diff-check PASS, no unstaged tracked diff, and all `39` protected reports unstaged. Isolated commit:

- hash `2acd357db32ac0c2bd3c3968a950c773acf3000d`;
- subject `feat(resolution): add external target representation`;
- parent `8055f0a6860721e26462572e34469e0d708d4a52`;
- exact manifest `23` paths;
- `1,491` insertions / `137` deletions;
- post-commit index empty;
- post-commit remaining status exactly `39` protected untracked reports;
- commit missing/extra `0/0`, unexpected outside classification `0`, diff-check PASS.

P6-C2 commit categories are `9` production + `8` tests + `4` Child 06 ledgers + the accepted coder report + final PASS report. Both P6-C2 REJECT reports and all other history remain outside it.

## Main-Owned P6-C2-to-P6-C3 Ledger Transition

After the isolated commit, Main used the planner skill and all four templates to record the truthful transition exactly once. This is a post-commit transition record, not a second P6-C2 acceptance finalization.

The four tracked ledger diffs record:

- P6-C2 commit `2acd357db32ac0c2bd3c3968a950c773acf3000d` and exact manifest/index/protected-boundary evidence;
- clean analyze/detect reuse on unchanged identities;
- `E6-P6C2-COMMIT1` closure;
- actual-status refresh `R18`;
- P6-C3 `LOCKED -> OPEN`;
- P6-D/target remain locked.

Exact Main baseline identities at `2026-08-22T19:28:18.0582096+07:00`, before any coder production/test edit:

1. `2026-07-28-06-ambient-external-resolution-and-diagnostics-plan.md` — `68,961` bytes — SHA-256 `6F0B583C1CA7D85DD87897A182F2143B63DD52B58FB51ECE8E1E748920ED3D15`.
2. `2026-07-28-06-ambient-external-resolution-and-diagnostics-evidence.md` — `77,961` bytes — SHA-256 `C57F6837AD82EFF3CE9CDE27EF772E5E861CE3F65A9AF60349E24F2C410391F6`.
3. `2026-07-28-06-ambient-external-resolution-and-diagnostics-benchmark.md` — `17,507` bytes — SHA-256 `F1080934185E6A86446ACEF2353AC3C971904F836F92C3A020CFD9DB29C630A5`.
4. `2026-07-28-06-ambient-external-resolution-and-diagnostics-actual-status.md` — `45,941` bytes — SHA-256 `2A6AEF6202702C8EF42E7218E688C3DF7FF20D9CBC437015C323A43E44D0ADF7`.

These four paths are the Main-owned P6-C3 worktree baseline. The coder may extend them only with truthful P6-C3 worker evidence/status/benchmark while preserving P6-C2/R18. Successor must distinguish this baseline from coder-authored additions.

## Active P6-C3 Coder Lane

Only active worker lane:

- Task `01a02946-f50c-7a63-ae34-a98cee8bac1a`.
- Title `P6-C3 Structured Resolution Outcomes — Coder (Clean)`.
- Current turn `01a02946-f704-7150-8382-e4f0afc28003`.
- Wait cursor at final Main snapshot `15923456-b532-4f51-acc6-537a810baf21:127`.
- State `active/inProgress`, no error/attention, no deadline, no report.
- Do not STOP/restart because Main rotated.

The original same-directory fork `01a0293c-f83b-7f91-9400-0ae0934dee9c` was STOPped after stuck initialization and is idle. It produced no ACK/tool/repo action and owns nothing.

Coder mandatory boundary:

- only `E:\Anvien` project-local checkout;
- no `C:\`, alternate worktree, target, network/install/external side effects;
- no stage/commit/push/branch/stash/reset/checkout;
- P6-D, graph-health policy/rendering, and `E:\cheapapp.org` locked;
- 39 pre-existing protected reports untouched;
- Main-owned four-ledger transition baseline preserved;
- disposition may be only `READY_FOR_INDEPENDENT_SUPERVISOR`, never self-PASS.

Coder fresh graph completed after two invalid pre-index invocations:

1. doubled path escaping caused an immediate path invocation failure; repo unchanged;
2. repo-local HOME hid global safe.directory, causing Git fail-closed before indexing; repo unchanged;
3. process-local `safe.directory=E:/Anvien` under the E-only envelope succeeded.

Valid fresh analyze evidence:

- `2,046` scanned;
- `756` parsed code;
- `0` failed;
- `116,746` nodes;
- `162,270` relationships;
- `19,658` dependency edges.

P6-C3 selected design boundary before edit:

- Goal: exactly one immutable/versioned structured resolver outcome per affected stable source site, with zero resolved/unresolved overlap and deterministic repository/external/intrinsic/capability precedence.
- Editable owner set: `internal/resolution/types.go`, new isolated `internal/resolution/outcome.go`, bounded hooks in `internal/resolution/resolve.go` and `internal/resolution/emit.go`, and projection in `internal/analyze/analyze.go`.
- Inspect/validate only: `internal/resolution/external_symbol.go`, shared graph structs, graph-health, Ladybug schema/CSV, generic graph JSON.
- `graph.Evidence.Note` / diagnostic Note is lossless through relationship Evidence and ResolutionGap/Ladybug. Intrinsic/unattributed outcomes have no existing persisted carrier and remain at the immutable `resolution.Result` boundary; shared graph structs are not opened.

File impacts recorded:

- `resolve.go`: CRITICAL, `5` affected files / `31` affected symbols;
- `emit.go`: CRITICAL, `6` files / `33` symbols;
- `analyze.go`: CRITICAL, `14` files / `51` symbols.

Seam impact inventory is CRITICAL:

- `ResolveBoundInto`: `7` symbols / `5` modules / `31` processes;
- `resolveTypeAnnotation`: `6 / 4 / 32`;
- `newEmitter`: `6 / 4 / 32`;
- `emitReference`: `8 / 2 / 36`;
- `emitUnresolvedReference`: `7 / 2 / 36`;
- `analyze.Result`: `5 / 4 / 31`;
- `analyze.Run`: `5 / 3 / 23`.

At final Main lane snapshot, coder reported direct UID impacts had closed the pre-edit gate, HEAD/parent and index remained correct, and only the four Main ledgers plus `39` protected untracked reports existed. It was beginning production-first work with the isolated outcome owner and bounded resolver/emitter/analyze hooks. No production/test edit existed at the exact Git snapshot below; dynamic state may advance after this report is sealed, so successor must re-snapshot immediately and monitor the same active turn.

Required coder evidence remains `E6-P6C3-IMPACT1`, `SRC1`, `STATUS1`, `BUILD1`, `TEST1`, `PARITY1`, and `DETECT1`. `REVIEW1`/`COMMIT1` remain pending for Supervisor/Main. Expected report convention is `reports/coder/rp_coder_<YYMMDD>_<HHMMSS>_by_<model_slug>_p6c3_structured_resolution_outcomes.md` with external seal and exact candidate/protected boundary.

No Supervisor lane is open yet. Open one separate visible Supervisor only after the coder seals a stable report and Main independently verifies source, diff, evidence, report identity, and Git boundary.

## Git Boundary Before This Handoff Report

Exact read-only snapshot `2026-08-22T19:28:18.0582096+07:00`:

- branch `master`;
- HEAD `2acd357db32ac0c2bd3c3968a950c773acf3000d`;
- parent `8055f0a6860721e26462572e34469e0d708d4a52`;
- status `43 = 4 tracked + 39 untracked`;
- non-protected paths exactly the four Main-owned Child 06 ledgers;
- protected `39 = 32 Main + 5 Supervisor + 2 coder`;
- index empty;
- `git diff --check` PASS;
- no unexpected path;
- no coder production/test edit yet at this exact snapshot.

Creation of this handoff adds exactly one protected Main report path. The official transfer message supplies the final post-report dynamic snapshot and exact report seal. Because the coder remains active, successor must treat the official snapshot as a point-in-time boundary and re-snapshot immediately.

## Exact Successor Ownership

1. Verify this sealed report identity externally, read it fully, then read `AGENTS.md`, `working-rules`, `orchestration`, applicable skills, and all four current Child 06 ledgers.
2. Re-snapshot Git immediately. Separate the four Main ledger baseline hashes above, any coder production/test/ledger changes created after `19:28:18`, the eventual coder report, and all protected history. Keep index empty.
3. Resume monitoring active coder task `01a02946-f50c-7a63-ae34-a98cee8bac1a` from cursor `15923456-b532-4f51-acc6-537a810baf21:127`. Do not STOP/restart due Main rotation and do not impose a lane deadline.
4. Enforce production-first ordering, exact editable owner set, graph-health/shared graph/P6-D/target locks, full build before validation, truthful baseline failure classification, affected persistence/reader parity, exact cleanup, fresh final analyze/detect, and sealed coder report.
5. If coder evidence proves Note/shared-struct carriage is not lossless or requires a major boundary change, stop implementation at a safe point and perform a concrete planner decision; do not let the worker self-expand.
6. When coder hands off, externally measure/read its report, read full source/test diff, verify candidate manifest and protected boundary, and only then open a separate visible Supervisor lane with zero-trust authority. Main's launcher-metadata disclosure is procedural history, not a coder repair invariant.
7. If Supervisor REJECTs, route only concrete owning invariants to the same P6-C3 coder. If PASSes, Main uses planner once, runs required current analyze/detect on accepted bytes, stages exact accepted manifest, and creates one isolated P6-C3 commit.
8. Only after P6-C3 commit may P6-D open. `E:\cheapapp.org` remains forbidden until P6-D explicitly opens.
9. Initialize the next Main successor by `2026-08-22 20:34:30 +07:00` and transfer by `2026-08-22 20:49:30 +07:00` if the campaign remains active.

Owner correction remains binding: coder/Supervisor lanes have no deadlines; only Main has a rotation deadline.
