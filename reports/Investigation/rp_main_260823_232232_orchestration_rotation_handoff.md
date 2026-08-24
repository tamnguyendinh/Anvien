# Visible Main Orchestration Rotation Handoff — 2026-08-23 23:22:32 +07

## Authority

- Outgoing Main: `01a02f0c-9eda-7a31-a56a-b60600e33392`.
- Prepared visible successor: `01a02f61-5384-7410-94e3-6421a2817ab0`, host `local`, direct saved-project checkout `E:\Anvien`.
- Outgoing Main retains authority until it sends a follow-up whose exact heading is `OFFICIAL AUTHORITY TRANSFER` with this report's detached seal. Authority then transfers immediately and outgoing Main terminates.
- Current authority began from raw official user-message UUIDv7 `01a02f37-5dce-7113-beb2-df8f004f48c6` at `2026-08-23T22:22:32.782+07:00`; warm target was `2026-08-23T23:07:32.782+07:00`; hard deadline is `2026-08-23T23:22:32.782+07:00`.
- Successor UUIDv7 creation time is `2026-08-23T23:08:22.660+07:00`, exactly `49.878 seconds` after the warm target. This is a verified `HIGH` rotation violation. Guard warning was accepted and the successor was immediately confirmed `WAITING_FOR_OFFICIAL_TRANSFER` with no authority or action.
- Historical verified HIGH violation from an earlier rotation also remains recorded: one successor was created at least `6 minutes 16.3529916 seconds` late while hard transfer remained on time.
- The successor must calculate its next warm/hard targets from the actual UUIDv7 timestamp of the incoming official authority user message.

## Iron startup

Before any orchestration action after transfer, read FULL RAW sequentially through EOF:

1. `E:\Anvien\AGENTS.md`
2. `E:\Anvien\.agents\skills\working-rules\SKILL.md`
3. `E:\Anvien\.agents\skills\orchestration\SKILL.md`
4. this sealed handoff

Before any plan/ledger update, additionally read FULL RAW:

1. `E:\Anvien\.agents\skills\planner\SKILL.md`
2. the roadmap
3. all four Child 06 ledgers

No summary, compacted context, or memory may replace these raw files. Apply 100% of the raw, unmodified `AGENTS.md`, `working-rules`, and `orchestration` rules. `governance-rule-guard` remains an independent Guard; Main must not read or use that skill.

## Owner authority and operating corrections

- Drive P6-D to completion without assigning bytes, choices, instructions, or waiting responsibility to Owner.
- Translate Owner warnings into the minimum operational delta; never forward reprimands to workers.
- Do not turn implementation into documentation/audit/review loops. Documentation follows verified progress once, using planner when required.
- `.tmp` is disposable debug/scratch only and may be deleted wholesale by Owner at any time. Reusable logic belongs under `scripts/`; dependency authority belongs under root `vendor/` and `third_party/go-vendor/`; durable evidence belongs under `reports/`.
- Sole campaign workspace/build authority is `E:\Anvien`. No alternate worktree, no `E:\cheapapp.org`, no target access before the required current-runtime gate, no C-drive workspace/read/write, no shared/global/protected/quarantined cache, and no stage/commit/push without later exact authority.
- The active Coder's historical narrow explicit network authority applied only to exact Go-module acquisition through `GOPROXY=https://proxy.golang.org`, no direct/VCS fallback. Acquisition/materialization is now complete. Runtime/build architecture remains checked-in vendor with `GOPROXY=off`, `GOSUMDB=off`, and no cache/network fallback.

## Active slice and fixed architecture

- Campaign: Anvien Graph Accuracy / Analyze Performance.
- `P6-D` remains unchecked and `REJECT/BLOCKED` until complete current-runtime, native/built-reader, target/oracle/boundary, independent Supervisor, detect, cleanup, and isolated commit proof exists.
- `P6-E` remains exactly one U1-U4 slice and `LOCKED`; no implementation, target, stage, commit, or lane authority exists for it.
- Final architecture is closed: exact checked-in root `E:\Anvien\vendor`, standard `go mod vendor`, unchanged `go.mod`/`go.sum`, explicit `-mod=vendor`, offline sanitized Go controls, deterministic `third_party\go-vendor\manifest.v1.json`, complete licenses, reusable verifier and negative tests. No architecture reopening.

## Verified durable materialization

The active Coder completed materialization from clean root `E:\Anvien\.tmp\p6d-go-vendor-materialize-clean-01a02ef3-260823-09` with exit `0` before production edits:

- durable scripts: `scripts/materialize-go-vendor.ps1`, `scripts/verify-go-vendor.ps1`;
- obsolete scratch source removed: `01-acquire-proxy.ps1`, `02-build-file-proxy.ps1` under the original lane root;
- `88/88` unchanged-go.sum-authorized `/go.mod` hashes verified;
- `45/45` vendor source modules verified for both content and `/go.mod`; actual content/go.mod mismatches `0/0`;
- offline module differences `0`, checksum mismatches `0`, Go-native verify PASS `2/2`;
- standard vendor output `1,578` files, exact equality `2/2`, tree SHA-256 prefix/suffix `34FE3F5C...ABD65B`;
- licenses `45/45` copied, explicit exceptions `0`;
- manifest-covered durable tree `1,634` files / `20,332,100` bytes, tree SHA-256 prefix/suffix `8A501CC6...B6AC`;
- physical Main remeasurement after publication: `vendor/` = `1,578` files / `20,135,220` bytes; `third_party/go-vendor/` = `57` files / `585,897` bytes. The physical one-file/manifest-byte difference from the manifest-covered denominator is the self-excluded `manifest.v1.json`; this still requires later zero-trust verification, not assumption.
- `go.mod` remains SHA-256 `C203E25E0A83A583E89A9D8F8E65BC0881052B3075E46915F90A4FB6433BD249`.
- `go.sum` remains SHA-256 `5434F39B08F50F1C53A71333E2E971CFCDE0E80BCA7DD028C00E31B9ACCAEC73`.

## Active visible Coder

- Task: `01a02f10-c175-7162-ac32-6df30a4d604a`, host `local`, direct checkout `E:\Anvien`.
- State at seal preparation: `RUNNING`; do not stop, duplicate, replace, or hand off its worker authority during Main rotation.
- Pre-edit file/symbol impact evidence:
  - `internal/cli/package_runtime.go`: `CRITICAL`; file batch affected `8` files and process `BuildGoRuntimePackage -> CommandPackageOutput`;
  - `buildGoRuntimePackage`: `CRITICAL`, `6` impacted, `2` direct, `10` process hits;
  - `prepareGoSourcePackage`: `CRITICAL`, `6` impacted, `2` direct, `10` process hits;
  - `hasPackageGoSource`: `HIGH`, `4` impacted, `1` direct, `3` process hits;
  - `ensurePackagedRuntime`: `CRITICAL`, `3` upstream files, `3` flows, `3` linked tests;
  - PowerShell owner `LOW 0` results are extractor limitation, not no-consumer proof.
- Production-first patch at `2026-08-23T23:10:47+07:00` touched exactly the expected three existing production owners: `scripts/full-build.ps1`, `internal/cli/package_runtime.go`, and `anvien-launcher/build.ps1`.
- Current production intent: verifier-before-first-Go, explicit `-mod=vendor`, sanitized offline Go environment, exact vendor/manifest/licenses/verifier package carriage, metadata manifest SHA-256, and fail-closed packaged-runtime identity acceptance.
- First uncompleted gate at seal preparation is Windows PowerShell 5.1 compatibility of `scripts/verify-go-vendor.ps1` before tests:
  - first runtime failure was missing `[System.IO.Path]::GetRelativePath`; Coder replaced it with a compatible helper;
  - second runtime failure is missing `[System.Security.Cryptography.SHA256]::HashData` / `Convert.ToHexString`; Coder is replacing these with Windows PowerShell 5.1-compatible hashing primitives;
  - no test edit is accepted before production verifier and offline compile pass.
- Coder retains no plan/ledger, Supervisor, P6-E, target, Git stage/commit/push, or independent acceptance authority.

Required Coder continuation, in order:

1. Complete Windows PowerShell 5.1 verifier compatibility and require fail-closed verifier PASS against exact root authority.
2. Compile/check the production changes offline with explicit vendor mode before updating tests.
3. Update only exact R42 tests and reusable negative contract test after production behavior is correct.
4. Run Rule 13, terminate build-related lock holders before full build, then canonical full build.
5. Prove current runtime, native-through-runtime, packaged/native and every affected built reader with exact vendor/manifest identity.
6. Capture target pre-state and validate target/oracle/boundary only after current-runtime gates exist.
7. Run affected regressions, fresh `anvien analyze --force`, fresh detect, exact disposable artifact disposition, and one sealed Coder report with `READY_FOR_MAIN` only if all gates pass.

## Main continuation after Coder handoff

1. Continuously monitor the active Coder's actual commands, files, scope, and gates; immediately correct drift or loops.
2. On `READY_FOR_MAIN`, Main performs zero-trust source/diff/runtime/evidence verification without restarting passed gates.
3. Before the plan update, read FULL RAW planner skill, roadmap, and all four Child 06 ledgers; then use planner once to refresh exactly those four ledgers from verified reality.
4. Open one fresh visible independent Supervisor only after current runtime, native/built readers, target/oracle/boundary, and clean procedure evidence exist.
5. On Supervisor PASS: fresh detect, exact cleanup disposition, exact staging, isolated P6-D commit; then and only then open P6-E.
6. On Supervisor REJECT: return only the exact rejected invariant to the appropriate worker boundary.

## Durable plan/checkpoint authority

- Roadmap: `E:\Anvien\docs\plans\2026-07-26-anvien-graph-accuracy-multi-plan\2026-07-26-anvien-graph-accuracy-roadmap.md`.
- Child 06 ledgers: the four files under `E:\Anvien\docs\plans\2026-07-26-anvien-graph-accuracy-multi-plan\2026-07-28-06-ambient-external-resolution-and-diagnostics\`.
- Current Child 06 plan identity after Owner's `.tmp` rule: `119,472` bytes / `670` LF / `0` CR / SHA-256 `BA666BBCCED2A4783CE3AF2453658388245005D58E67F9318D89692BD0EAB82D`.
- Architect report: `E:\Anvien\reports\system-architect\rp_system-architect_260823_210034_by_gpt-5_p6d_clean_e_only_go_vendor_closure.md`; SHA-256 `C2A20A72F94A971317E54A4DF77973E5CA4DC8516221FB72D43DD589BC75C776`.
- Prior recovery report: `E:\Anvien\reports\coder\rp_coder_260823_201156_by_gpt-5_p6d_native_canonical_build_recovery_blocked.md`; SHA-256 `7019234AE8BBE030CADEA73B91760001EC9E9089C6B7A775D23CBE12BEF58A22`.

## Workspace and Git boundary

- Preserve branch `master`, HEAD `741335f7efa5ad215ec28361d88c462f431565ab`, parent `838f379ab00c93dfe7507603f6608bb08d61f038`, Owner deletion `CONTRIBUTING.md`, every unrelated/protected path, and empty index.
- No stage/commit/push authority transfers yet.
- Current worktree contains earlier P6-D/native/projector/plan evidence plus the active vendor/R42 production work; do not misclassify or clean unrelated/protected changes.
- Coder owns active production/test/generated-root modifications. Main must not edit the same files in parallel.

Outgoing Main terminates immediately after official transfer and performs no further filesystem, Git, graph, build, lane-control, or orchestration action.
