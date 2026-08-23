# SEALED MAIN ORCHESTRATION ROTATION HANDOFF

## 1. Authority envelope

- Campaign: Anvien Graph Accuracy.
- Outgoing Main: task `01a02637-a47d-7ba0-b6a9-b643f9fa415b`.
- Designated successor Main: task `01a02668-89e7-7642-b55b-d5c6f41c2daf`, host `local`, same saved project and authoritative checkout.
- Successor PRE-TRANSFER ACK: `UNDERSTOOD / WAITING_FOR_OFFICIAL_TRANSFER`; it confirmed the sole boundary `E:\Anvien`, the ban on `C:\`, alternate worktrees, and `E:\cheapapp.org`, and no action before official follow-up.
- This report is the immutable handoff payload. Authority transfers only when outgoing Main sends the exact phrase `OFFICIAL AUTHORITY TRANSFER` to the designated successor together with this report identity.
- Exact report `createdAt`: `2026-08-22 05:24:26 +07:00`.
- Successor absolute rotation deadline: `2026-08-22 06:24:26 +07:00`.

## 2. Locked workspace and safety boundary

- Authoritative checkout and only permitted repository boundary: `E:\Anvien`.
- Do not access `C:\`, any alternate checkout/worktree, or `E:\cheapapp.org`.
- Target `E:\cheapapp.org` remains locked until P6-D.
- Active P6-B still forbids network access, dependency installation, package-script execution, stage, commit, push, target access, and hidden/internal subagents.
- Direct offline `node anvien-web/scripts/generate-tsstdlib-catalog.mjs` is the accepted generation-time authority; runtime Node/`tsc` is forbidden. No package script was invoked.
- Preserve all protected untracked Main handoffs under `reports/Investigation/`; never edit, delete, stage, or commit them.
- Preserve all active P6-B source/test/ledger/report diff, repo-local `.tmp` evidence, packaged-runtime artifacts, and any later lane-owned output. Do not overwrite or absorb lane ownership.

## 3. Incoming handoff verification completed by outgoing Main

Outgoing Main independently read the full root `AGENTS.md` and `CLAUDE.md`, full `working-rules`, `orchestration`, and `supervisor` skills, the four Child 06 ledgers, the P6-A Architect report, both P6-A Supervisor reports, and historical commits `b98131e4` / `ec765deb`.

Verified prior sealed handoff:

- Path: `E:\Anvien\reports\Investigation\rp_main_260822_042700_orchestration_rotation_handoff.md`.
- Identity: `12,946 bytes / 172 LF / 0 CR / strict UTF-8 without BOM`.
- SHA-256: `27CBB0F093BC1723D32F0C818DBD2A6413657B807596062A4D98EF8AE58D9BCC`.
- Exact createdAt: `2026-08-22 04:27:00 +07:00`.
- Incoming Git boundary at `2026-08-22 04:33:22 +07:00` matched HEAD `b98131e44932a7bcac17b487ecb2914535927d01`, parent `ec765debff335540c77d409ebb2c9f45e4a0a77d`, ahead/behind `54/0`, empty index, clean tracked worktree, and both diff checks PASS.
- Exactly seventeen protected Main handoffs existed. Three legacy filenames use a four-digit time while fourteen use six digits; all seventeen were identified from full porcelain inventory and preserved.

## 4. Closed P6-A state

P6-A remains independently accepted and committed by `b98131e44932a7bcac17b487ecb2914535927d01`, exact seven-path manifest, `599 insertions / 63 deletions`.

Closed decision remains unchanged:

1. Offline TypeScript `5.9.3` compiler-API generation from official `lib*.d.ts` inputs.
2. Checked-in compact provenance catalog embedded into the Go binary with `go:embed`.
3. Runtime has no Node, `tsc`, network, dependency installation, or package-script requirement.
4. Supported config is TypeScript-only zero/one root JSONC `tsconfig.json`, with only `target`, `lib`, and `noLib`; unsupported topology fails closed as `capability_unavailable`.
5. Repository/P5 resolution precedes external lookup; explicit-import failure is terminal.
6. P6-C1 is `PRESERVE_ONLY`.
7. P6-C2 is `ACTIVE`, referenced-only `ExternalSymbol`, and remains locked until P6-B plus P6-C1 closure.

Do not reopen P6-A gates unless its source/evidence boundary is invalidated.

## 5. Active P6-B Coder lane

- Task: `01a02637-4ac8-7031-9043-fea65333c7b4`.
- Host: `local`.
- Turn: `01a02637-4d6d-76f1-b3a7-1a5c717b23dd`.
- Latest observed cursor at sealing: `bfa88881-8500-4623-b486-d2af38779a76:116`.
- State: `ACTIVE`; no durable coder handoff yet and no acceptance claim.
- The turn auto-compacted after the packaged-runtime analyze completed. It re-read `AGENTS.md`, `working-rules`, and `coder`, re-anchored the CRITICAL blast radius, and is continuing validation/ledger/report only; do not restart passed gates merely because of compaction.
- Boundary update was ACKed: exactly seventeen inherited protected Main handoffs before creation of this report; this new report becomes the eighteenth.

### 5.1 Pre-edit graph and blast radius

- Fresh graph after P6-B ledger transition: `1,990 scanned / 743 parsed-code / 0 failed / 116,535 nodes / 160,659 relationships`.
- Existing owner file-detail related-file counts: analyze `187`, resolution types `34`, indexes `58`, resolve `60`.
- File impacts: `internal/analyze/analyze.go` CRITICAL `100` impacted; `types.go` CRITICAL `62`; `indexes.go` CRITICAL `234`; `resolve.go` CRITICAL `112`.
- Symbol impacts: `analyze.Run` CRITICAL `24 impacted / 8 modules / 23 processes`; `workspace` CRITICAL `134 / 7 / 69`; `BuildCrossFileBinding` CRITICAL `95 / 24 / 31`; `resolveCall`, `resolveAccess`, and `resolveTypeAnnotation` each CRITICAL `28 / 7 / 32`. `Options` and `Metrics` were LOW individually, but their containing file remained CRITICAL.
- Several first symbol-impact commands were invalid because `impact symbol` does not accept `--file`; the lane did not treat null results as evidence and reran with exact symbol UIDs.

### 5.2 Current production/test mechanism

Production owners are confined to P6-B:

- new `anvien-web/scripts/generate-tsstdlib-catalog.mjs`;
- new `internal/tsstdlib/catalog.v1.json`, `catalog.go`, and `profile.go`;
- minimum wiring in `internal/analyze/analyze.go`, `internal/resolution/types.go`, `internal/resolution/indexes.go`, and `internal/resolution/resolve.go`;
- new P6-B tests/fixtures under `internal/tsstdlib`, `internal/resolution/p6b_tsstdlib_test.go`, and `internal/analyze/p6b_tsstdlib_test.go` plus its repo-local fixture.

No manifest/lock, shared graph struct, persistence, graph-health, P6-C, or P6-D file is changed.

Implementation facts observed:

- Generator hard-locks TypeScript `5.9.3` and the accepted package-lock integrity; package manifest/lock remain zero-diff.
- First generator implementation flattened inherited members and produced `7,901,468 bytes / 80,689 members`; this was debug/benchmark finding, then overwritten.
- Current compact artifact has `100 inputs / 3,141,835 input bytes / 2,030 symbols / 11,802 direct/base-linked members / 1,388,709 artifact bytes`.
- Two clean direct generation runs were byte-equal: artifact SHA-256 `78C9F6F49896D2F11E7E45B58C9D7CB90E1CA20413F06FFE9FAC95310C1AF054`; internal catalog hash `22ab8f6986bfa8efd01addffac64b4265fff09244391ecb5d385b80dfa187119`.
- Initial generator `lanes is not defined` failure was fixed before current artifact.
- Initial Go loader canonical hash mismatch caused fail-closed `catalog_hash_mismatch`; fixed by matching non-HTML-escaped canonical JSON while preserving validation.
- Real analyze test found a sibling dotted `Math.max` TypeAnnotation gap; fixed with a general dotted owner/member lookup by TypeRef meaning, not a target-name branch.
- Focused `tsstdlib`, `resolution`, and `analyze` suites passed after those fixes.

### 5.3 Build, regression, runtime, and benchmark evidence so far

Build authority conflict is explicit:

- README official full build invokes `npm install` and package scripts, forbidden by the active lane authority. It was not run and must not be called PASS.
- `go build ./...` was attempted and failed only on intentionally non-buildable fixture packages under `anvien/test/fixtures`; it is not PASS evidence.
- Production package build `go build ./cmd/... ./internal/...` passed.
- Both launcher wrapper Go modules passed.
- Direct `go run ../cmd/anvien package build-runtime` from `E:\Anvien\anvien` passed and wrote `anvien/bin/anvien.exe`.
- Packaged binary final: `72,961,536 bytes`, SHA-256 `6E28F94BE0CBF8A2DAE1A17C8D6DA6296EB7C828D82299029F9D4DC953B4A040`.
- Pre-existing backup baseline `anvien/bin/anvien.exe~`: `71,509,504 bytes`, SHA-256 `BC060FDDF46394B72859B86930409B1B7E91C996BA304A86F0A633C37E043532`; measured size delta `+1,452,032 bytes`.
- PATH restricted to the two E-local binary/native directories failed `0xC0000135` because the Ladybug native dependency chain needs additional runtime directories. The lane did not chase dependencies on `C:\` and did not mislabel the probe as PASS or Node evidence.
- Packaged binary with inherited runtime environment passed `anvien.exe version` and reported `1.2.8`.
- Built packaged binary `analyze . --force --benchmark-json .tmp/p6b-analyze-benchmark.json --benchmark-label p6b-final --json` passed in `230.031s` on authoritative `E:\Anvien`; observed file projection reports `2,005` files, `19,264` dependency edges, `457` unresolved, with stale flags false in returned details. No P6-D target access occurred.
- Catalog load benchmark across three runs: about `63.7-64.5 ms/op`, `17,623,402-17,623,713 B/op`, `56,709 allocs/op`.
- Global lookup benchmark: about `252.7-297.5 ns/op`, `192 B/op`, `1 alloc/op`.
- Inherited-member lookup benchmark: about `613.9-667.9 ns/op`, `216 B/op`, `4 allocs/op`.

Regression truth:

- `go test ./cmd/... ./internal/... -count=1` exited `1`; do not call it a full PASS.
- Exactly two failures reproduce known accepted baseline mismatches already recorded by Child 04/05: C# expected two `ACCESSES` and got none; Dart expected two and got one. Repeated focused runs reproduced the same values. All other packages, including current P6-B focused surfaces, passed.
- Do not fix those two out-of-scope parity baselines inside P6-B. Supervisor must still decide whether the current evidence sufficiently proves no P6-B regression.

### 5.4 Pending coder gates

The active lane still owns and must finish:

1. Consume the completed packaged-runtime analyze result without rerunning the same passed gate.
2. Inspect final exact source/diff and prove no target-name allowlist, no runtime process spawn/Node dependency, repository/P5 precedence, explicit-import terminal behavior, config fail-closed behavior, language isolation, and no silent loss of eligible/unavailable sites within P6-B's interim contract.
3. Record final evidence and benchmark values in all required Child 06 living ledgers.
4. Remove only lane-created dead/debug artifacts, especially untracked `anvien-launcher/server-wrapper/anvien-server-wrapper.exe` and `anvien-launcher/src/anvien-launcher.exe`; preserve accepted catalog/tests and every user/pre-existing artifact.
5. Refresh graph only if required by later file changes, then run `anvien detect-changes --repo Anvien --scope all` before coder handoff.
6. Run final diff-check/status, seal an immutable coder report with exact bytes/LF/CR/strict UTF-8/SHA-256/createdAt, list the exact owned diff, and stop for independent review.

A coder report is evidence input, not acceptance.

## 6. Bounded Git/worktree snapshot

Snapshot time: `2026-08-22 05:24:12 +07:00`.

- Branch: `master`.
- HEAD: `b98131e44932a7bcac17b487ecb2914535927d01`.
- Parent: `ec765debff335540c77d409ebb2c9f45e4a0a77d`.
- `origin/master...HEAD`: left/right `0 / 54`, therefore ahead 54, behind 0.
- Index: empty.
- Cached and worktree diff checks: PASS, no output.
- Tracked worktree: seven P6-B owned modified paths, total `210 insertions / 14 deletions`:
  - actual-status `6+/5-`;
  - evidence `13+/3-`;
  - plan `15+/6-`;
  - analyze `18+/0-`;
  - indexes `3+/0-`;
  - resolve `149+/0-`;
  - resolution types `6+/0-`.
- Untracked count at snapshot: `35` = `17` protected Main handoffs + `18` lane-owned paths.
- The 18 lane-owned untracked paths consist of 16 intended new generator/catalog/profile/test/fixture paths plus two known debug wrapper executables pending cleanup.
- Creation of this sealed report adds one protected Main handoff, so expected protected count becomes exactly `18` and total untracked becomes `36` unless the active lane changes its owned artifacts after the snapshot.
- The lane remains active and may legitimately mutate its owned diff after this bounded snapshot. Preserve and attribute later changes; do not assume this snapshot is its final manifest.

## 7. Visible lane inventory

### Active

- P6-B Coder task `01a02637-4ac8-7031-9043-fea65333c7b4`, cursor `bfa88881-8500-4623-b486-d2af38779a76:116`, ACTIVE.

### Waiting successor

- New Main task `01a02668-89e7-7642-b55b-d5c6f41c2daf`, PRE-TRANSFER ACK complete, IDLE / `WAITING_FOR_OFFICIAL_TRANSFER` until exact official follow-up.

### Idle / do not reuse

- P6-A Architect/Planner `01a025f0-4cbb-7f90-b663-bd9f0bfc954c`: IDLE.
- P6-A Supervisor `01a0261a-345a-7212-9b01-008e18ae7367`: IDLE/PASS; do not reuse for P6-B.
- Child 05 lanes remain CLOSED/IDLE per the prior sealed report.

## 8. Exact successor actions

1. After official authority, read this full sealed report, `AGENTS.md`/rules, applicable skills, current four Child 06 ledgers, and active lane's durable output when it exists.
2. Verify this report identity and current Git boundary exactly once. Preserve all eighteen Main handoffs and all current/later P6-B owned diff/evidence.
3. Immediately monitor only the existing P6-B lane from cursor `bfa88881-8500-4623-b486-d2af38779a76:116`; do not open another coder.
4. Ensure compaction re-anchor continues from the first uncompleted gate; do not rerun fresh analyze/build/focused tests/benchmarks solely because context compacted.
5. When coder seals its handoff, independently verify immutable report identity, exact source/diff, catalog provenance/determinism, config inventory and fail-closed topology, repository/P5 and explicit-import precedence, general lookup/oracle cases, packaged runtime, build authority conflict, baseline regression classification, benchmarks, cleanup, graph freshness, detect-changes, and ledger consistency.
6. Pay special attention to whether resolved/unavailable external lookups remain observable rather than silently disappearing before P6-C2/C3; do not infer acceptance from metrics alone.
7. Only after Main verification, open exactly one fresh visible independent P6-B Supervisor task. It must use repository rules plus `supervisor`, use current Anvien evidence, and remain read-only/no-stage/no-commit. Do not reuse prior Supervisor tasks.
8. Supervisor PASS alone authorizes P6-B acceptance/ledger-finalization/detect/stage/commit. REJECT returns only the exact rejected invariant to the correct repair owner.
9. Do not open P6-C1 or P6-C2 before P6-B is accepted and committed. Keep target locked until P6-D.
10. Rotate again before `2026-08-22 06:24:26 +07:00`.

## 9. Non-negotiable acceptance rules

- Reports are evidence inputs, never self-validating acceptance.
- Independent P6-B Supervisor PASS is mandatory before Main acceptance or commit.
- Fresh graph precedes every graph-based query and detect-changes when source has changed after the last analyze.
- File-detail and impact are mandatory before editing existing functions/classes/methods/exported/shared symbols; HIGH/CRITICAL are scope warnings, not bans.
- Code behavior precedes tests.
- Build/runtime failures and prohibited gates must remain labeled accurately; never convert a constrained or partial build into official full-build PASS.
- Never stage protected handoffs or active lane output before acceptance.
- Never access `E:\cheapapp.org` before P6-D.

## 10. Transfer terminal condition

Once outgoing Main sends the successor an official follow-up containing exact `OFFICIAL AUTHORITY TRANSFER` plus this report's measured identity, outgoing Main terminates immediately. The successor becomes sole Main authority and owns P6-B monitoring plus the next rotation deadline.
