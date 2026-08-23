# Main Orchestration Rotation Handoff — Anvien Graph Accuracy

## Authority

- Outgoing Main: task `01a0296e-3fab-7f53-babd-d819c3ac7190`.
- Waiting successor: task `01a0297e-a550-75c1-b0b4-f59a9a36f90a`.
- Exact internal createdAt: `2026-08-22 19:43:03 +07:00`.
- Successor absolute rotation deadline remains `2026-08-22 20:49:30 +07:00`.
- Successor required next warmup remains by `2026-08-22 20:34:30 +07:00`.
- Locked authoritative workspace remains only `E:\Anvien`.
- `C:\`, alternate worktree/checkouts, and `E:\cheapapp.org` remain forbidden. P6-D/target stays locked.
- External bytes/LF/CR/strict UTF-8/BOM/SHA-256/filesystem timestamps for this report are supplied in the official transfer message after the file is sealed. This report must not be edited after that measurement.

## Early Rotation Reason And Procedural Disclosure

This Main rotates early because its strict no-`C:\` observation boundary is no longer clean for any future P6-C3 acceptance/commit action.

The exact event occurred while the active P6-C3 coder was still implementing and before any coder handoff, Supervisor lane, acceptance, staging, or commit:

1. Main correctly resumed compact monitoring with `wait_threads`, which returned only lane state/commentary/tool markers and no host launcher path.
2. To inspect the active turn without command output, Main then called `read_thread` with `includeOutputs=false`.
3. Despite that explicit setting, the app response serialized historical `commandExecution` metadata containing the host PowerShell launcher path under `C:\`.
4. Main did not invoke, enumerate, inspect, hash, modify, or clean any `C:\` filesystem path. Every Main repository command and every coder command observed in scope targeted `E:\Anvien`.
5. No repository mutation resulted from reading thread metadata.
6. By the campaign's existing strict procedural precedent, launcher-path exposure prevents this Main from owning a later acceptance/commit boundary. Therefore it rotates immediately to a clean visible successor.

This disclosure is not a P6-C3 code blocker and must not be routed to the coder as a repair invariant. It does not invalidate P6-C2 or any accepted earlier commit. The successor must independently evaluate the eventual coder report and Supervisor evidence.

## Incoming Handoff Verification And Re-Anchoring Completed

Incoming sealed handoff was externally verified and read in full:

- Path: `reports/Investigation/rp_main_260822_192827_orchestration_rotation_handoff.md`.
- `15,099` bytes / `205` LF / `0` CR / strict UTF-8 without BOM.
- SHA-256 `DFD3D2BB1896A39CEDD3C5BE044B965F5AE9603416DC18CB9118F62A4DACD0E6`.
- Filesystem createdAt/lastWrite `2026-08-22T19:30:38.5182306+07:00`.

Main then read in full:

- `E:\Anvien\AGENTS.md`;
- `.agents/skills/working-rules/SKILL.md`;
- `.agents/skills/orchestration/SKILL.md`;
- `.agents/skills/planner/SKILL.md`;
- all four current Child 06 ledgers, including the entire plan/evidence/benchmark/actual-status rather than only P6-C3 sections.

The four Main-owned P6-C2-to-P6-C3 transition ledgers remained exactly at the sealed incoming baseline through the final dynamic snapshot:

1. plan: `68,961` bytes / SHA-256 `6F0B583C1CA7D85DD87897A182F2143B63DD52B58FB51ECE8E1E748920ED3D15`;
2. evidence: `77,961` bytes / SHA-256 `C57F6837AD82EFF3CE9CDE27EF772E5E861CE3F65A9AF60349E24F2C410391F6`;
3. benchmark: `17,507` bytes / SHA-256 `F1080934185E6A86446ACEF2353AC3C971904F836F92C3A020CFD9DB29C630A5`;
4. actual-status: `45,941` bytes / SHA-256 `2A6AEF6202702C8EF42E7218E688C3DF7FF20D9CBC437015C323A43E44D0ADF7`.

These diffs are the preserved Main transition baseline recording P6-C2 commit, `E6-P6C2-COMMIT1`, R18, and P6-C3 open. No coder P6-C3 ledger addition existed at the final snapshot.

## Accepted Campaign State

- P6-A remains committed at `b98131e44932a7bcac17b487ecb2914535927d01`.
- P6-B remains committed at `5c1584fef7153a7a331c8fedd1ce64176ddc873d`.
- P6-C1 remains committed at `8055f0a6860721e26462572e34469e0d708d4a52`.
- P6-C2 remains independently accepted and committed at `2acd357db32ac0c2bd3c3968a950c773acf3000d`, parent `8055f0a6860721e26462572e34469e0d708d4a52`.
- P6-C3 remains the sole open slice with one active visible coder.
- P6-D and `E:\cheapapp.org` remain locked.
- Do not invent another P6-C2 review, repair, analyze/detect, or commit loop.

## Active P6-C3 Coder Lane

Only active worker lane:

- Task `01a02946-f50c-7a63-ae34-a98cee8bac1a`.
- Turn `01a02946-f704-7150-8382-e4f0afc28003`.
- Latest compact wait cursor at handoff preparation: `15923456-b532-4f51-acc6-537a810baf21:133`.
- State `active/inProgress`, no error/attention, no deadline, no report.
- Do not STOP/restart because Main rotated.

Coder state at the latest compact status:

- production-first seam exists and remains format-clean;
- no test edit exists yet;
- fresh graph refresh required before touching the identified test functions has been running for roughly five minutes without output, stderr, lock failure, or competing writer;
- coder is correctly waiting on that exact process and has made no edits while it runs;
- if it later returns a concrete lock/process error, classify that evidence; otherwise do not kill it merely for elapsed time.

The exact current production owner set remains:

- `internal/resolution/types.go`;
- new `internal/resolution/outcome.go`;
- bounded hooks in `internal/resolution/resolve.go` and `internal/resolution/emit.go`;
- `internal/analyze/analyze.go`.

Inspect/validate-only locks remain `internal/resolution/external_symbol.go`, shared graph structs, Ladybug schema/CSV, graph-health, and generic graph transports. The current implementation records versioned final outcomes, finalization/exclusivity, repository/external/intrinsic/capability status, and adapters into relationship/reference Evidence, diagnostic Note, and `analyze.Result`; it has not opened graph-health policy or the P6-C2 materializer.

Main's source-diff inspection is monitoring evidence only, not acceptance. Tests/build/runtime/parity/detect/report remain absent and the candidate must not be treated as ready.

Exact production identities at `2026-08-22T19:43:26.7258544+07:00`:

- `internal/analyze/analyze.go`: `33,734` bytes / SHA-256 `CC9329FB8AD3510DFC5F643F39EE3CFBB749ABBB3F3CB130E519CE3DA2BBE608`;
- `internal/resolution/emit.go`: `27,319` bytes / SHA-256 `13A31C8EAE9F0665BBBA9552B3D75D1B2CA663DE3FA8594780ED0BCEACCDCB6E`;
- `internal/resolution/resolve.go`: `39,912` bytes / SHA-256 `D01322440AC0A2CF6D5F39402791F61B23AA6894707C1E350E5B56687080E919`;
- `internal/resolution/types.go`: `4,933` bytes / SHA-256 `7C5E113F5E50584665D6D9AED0BCB2B3C6F6085219A62CD5AE74DD7654CD5DC3`;
- `internal/resolution/outcome.go`: `19,944` bytes / SHA-256 `EF05764F15B2DD405E4F4CB9EE74688152F9A5C448AC322280E950EF3C9A2CA1`.

Required worker disposition remains `READY_FOR_INDEPENDENT_SUPERVISOR`, never self-PASS. Required evidence remains `E6-P6C3-IMPACT1`, `SRC1`, `STATUS1`, `BUILD1`, `TEST1`, `PARITY1`, and `DETECT1`; `REVIEW1`/`COMMIT1` remain Supervisor/Main owned.

## Git Boundary Before This Handoff Report

Exact read-only snapshot `2026-08-22T19:43:26.7258544+07:00`:

- branch `master`;
- HEAD `2acd357db32ac0c2bd3c3968a950c773acf3000d`;
- parent `8055f0a6860721e26462572e34469e0d708d4a52`;
- status `49 = 8 tracked + 41 untracked`;
- index empty;
- `git diff --check` PASS;
- non-protected current paths exactly `9`: four Main-owned Child 06 ledgers plus the five production paths listed above;
- protected current paths exactly `40 = 33 Main handoffs + 5 Supervisor reports + 2 old coder reports`;
- no test, new coder report, Supervisor report, staged path, or unexpected owner existed.

Creation of this handoff adds exactly one protected Main report. The official transfer message supplies the final post-report dynamic snapshot and exact report seal. Because the coder remains active, the successor must treat every snapshot as point-in-time and resnapshot immediately.

## Successor Initialization

- Valid waiting successor task: `01a0297e-a550-75c1-b0b4-f59a9a36f90a`.
- It returned `UNDERSTOOD / WAITING_FOR_OFFICIAL_TRANSFER`, confirmed the E-only boundary, and performed no command/repo/lane action.
- Earlier same-directory fork `01a0297d-31f8-7213-82fa-baabc85ce992` completed two empty turns with no assistant output, no tool, no repo action, and no ownership. It is idle and must not be used as Main.

## Exact Successor Ownership

1. Verify this sealed report identity externally; read it fully, then read `AGENTS.md`, `working-rules`, `orchestration`, applicable skills, and all four current Child 06 ledgers.
2. Re-snapshot Git immediately. Preserve the four Main ledger baseline hashes above; distinguish subsequent coder test/ledger/report additions from protected history; keep index empty.
3. Resume compact monitoring of coder task `01a02946-f50c-7a63-ae34-a98cee8bac1a` from cursor `15923456-b532-4f51-acc6-537a810baf21:133`. Use `wait_threads`; do not use `read_thread` because even `includeOutputs=false` exposed host launcher metadata in this Main.
4. Do not restart the current fresh analyze or impose a coder deadline. Enforce production-first ordering, exact owner/boundary locks, full build before validation, truthful baseline failures, persistence/reader parity, exact cleanup, fresh final analyze/detect, empty index, and sealed coder report.
5. If evidence proves Note/shared-struct carriage is not lossless or requires major boundary change, stop at a safe point and make a concrete planner decision; do not let the worker self-expand.
6. On coder handoff, externally measure/read the report, read the full source/test diff, verify exact candidate/protected boundary, and only then open one separate visible Supervisor lane. No Supervisor lane exists now.
7. If Supervisor REJECTs, route only concrete owning invariants to the same coder. If PASSes, use planner once, run required current analyze/detect on accepted bytes, stage the exact accepted manifest, and create one isolated P6-C3 commit.
8. Only after P6-C3 commit may P6-D open. Never access `E:\cheapapp.org` before P6-D.
9. Warm the next Main by `2026-08-22 20:34:30 +07:00` and transfer by `2026-08-22 20:49:30 +07:00` if the campaign remains active.

Owner correction remains binding: coder/Supervisor lanes have no deadlines; only Main has the absolute rotation deadline.
