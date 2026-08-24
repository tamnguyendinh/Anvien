# Coder Report: P6-D Durable Native / Canonical Build Recovery

Status: `BLOCKED`

## Metadata and authority

- Time: `2026-08-23 20:11:56 +07:00`.
- Role: Coder.
- Repository and sole filesystem boundary: `E:\Anvien`.
- Branch / HEAD / parent: `master` / `741335f7efa5ad215ec28361d88c462f431565ab` / `838f379ab00c93dfe7507603f6608bb08d61f038`.
- Unique lane root: `E:\Anvien\.tmp\p6d-native-recovery-01a02e4a-260823`.
- Active authority: existing P6-D native/current-runtime recovery only. This is not a new slice, gate, or commit.
- P6-D remains unchecked and `REJECT/BLOCKED`; P6-E remains one locked U1-U4 slice and was not accessed or changed.
- Next responsible task: Visible Main Orchestration `01a02e33-6974-7f30-951d-966a5fb7fec9`.
- Git authority: none. No stage, commit, push, reset, checkout, stash, branch change, clean, or restoration occurred.
- Target authority: none. `E:\cheapapp.org` was never accessed.
- Final disposition is Coder `BLOCKED`, never self-PASS.

## Outcome

The repository now has the exact durable LadybugDB v0.19.1 Windows x86_64 three-file input bundle, and the current production/build contract routes that authority through the canonical build, helper, package-runtime, launcher, and the one ledger-authorized QA consumer. The changed contract fails closed on missing, partial, extra, tampered, or alternate native authority; requires distinct unique lane/cache/runtime roots; prevents the Windows path from accepting shared `.tmp\ladybug-native`, silent downloads, package-script/install recovery, or a stale packaged runtime; and stages runtime/launcher outputs before publication.

The recovery cannot be accepted as complete. The real corrected canonical `scripts\full-build.ps1` invocation exited `1` at its first direct Go package-runtime build because its clean lane-owned module cache does not contain the exact Cobra, Hugot, and tree-sitter dependency closure and `GOPROXY=off` correctly blocks lookup. It produced no current runtime. Consequently, current-runtime provenance, canonical final analyze, native parity through the built runtime, and all affected built graph-health readers are absent. Target/oracle/pre-post proof is additionally forbidden by the current `E:\cheapapp.org` boundary.

There is also a procedural-integrity blocker: two wrappers used PowerShell's reserved case-insensitive `$Home`/`$HOME` variable before corrected E-rooted invocations. Both outputs were excluded, but proving or cleaning any inherited-home side effect would require crossing the E-only boundary. The final graph rerun was performed exactly once after PID/child/lock verification using non-reserved `$homeRoot`; this yields useful current-source graph evidence but cannot make the overall lane procedurally clean.

## Startup, active authority, and preservation re-anchor

The session read FULL RAW through EOF, in required order, before implementation continuation:

1. `E:\Anvien\AGENTS.md`.
2. `E:\Anvien\.agents\skills\working-rules\SKILL.md`.
3. `E:\Anvien\.agents\skills\coder\SKILL.md`.
4. `docs\plans\2026-07-26-anvien-graph-accuracy-multi-plan\2026-07-26-anvien-graph-accuracy-roadmap.md`.
5. All four Child 06 plan/evidence/benchmark/actual-status files.
6. Sealed prior blocker `reports\coder\rp_coder_260823_171110_by_gpt-5_p6d_real_recovery_blocked.md`.

The prior blocker is `12,008` bytes with expected SHA-256 `D335ADD9EF4B87928A5F5C9ED769A9FBEE91AAC8B7CB8424AA75C32BB93BBC1A`. The startup branch/HEAD/parent matched dispatch and the index was empty.

Latest Main correction is retained exactly: the roadmap and all four Child 06 ledgers are inspect-only/preserve-only for this Coder. No Coder ledger write occurred. Main alone performed R40 and authorized exactly `scripts/qa-child02-p2e-validation.ps1` as the additional same-invariant owner after the four primary production/build owners were correct.

The independently cleared six-field graph-health bytes remained preserve-only:

- `internal/graphhealth/diagnostics.go`: `26,884` bytes; SHA-256 `518BD06FF5F859BD42AD7B3B9FF8E3F9F4E7235720BC5E568E2C4B2DA72BD2AE`.
- `internal/graphhealth/p6d_outcome_projection_test.go`: `19,336` bytes; SHA-256 `6E90E655FFCCF5A5BD3D398D95EB5B6B7F77754CAAF35DAAF4D9609248A4E713`.

No P6-C3/six-field semantic branch, graph/Ladybug schema, public API, resolver, declaration model, scanner, target-specific behavior, or P6-E architecture was changed or re-audited.

## Invariant family map

- Input authority: exactly `third_party\ladybugdb\v0.19.1\windows-x86_64\{lbug.h,lbug_shared.lib,lbug_shared.dll}` with pinned identities.
- Canonical root authority: one unique strict child of `E:\Anvien\.tmp`; distinct `cache` and `runtime` strict children; E-only home/temp/Go/npm/Vite roots.
- Consumers: `scripts\full-build.ps1` -> `internal/cli/package_runtime.go` package runtime -> `anvien-launcher\build.ps1`; `scripts\ensure-ladybug-native.ps1` validates the same durable authority; the exact QA owner reads the same directory.
- Publication invariant: build from current source and pinned native inputs into lane staging; only publish after successful staged build; never accept a stale existing packaged binary.
- Failure invariant: missing, partial, extra, tampered, wrong-version, alternate-path, outside-lane, or shared cache/runtime inputs fail visibly.
- Validation boundary: current-source native Ladybug smoke is useful but is not built-runtime parity. Canonical runtime and every actually affected built reader remain mandatory.
- Forbidden fallbacks: shared/protected/quarantined cache; Windows `.tmp\ladybug-native` authority; runtime download; package install/script recovery; global/stale runtime; target-text inference; target access.

## Fresh Anvien pre-edit evidence and blast radius

`anvien --help` preceded graph work. The initial forced current-repository graph used only E-rooted session data and exited `0` with `2,096` scanned / `763` parsed / `0` failed, `121,216` nodes / `167,804` relationships, `20,323` dependency edges, and `478` unresolved projections.

Fresh full file-detail and upstream file/symbol impact were completed before edits:

- `scripts/full-build.ps1`, `scripts/ensure-ladybug-native.ps1`, and `anvien-launcher/build.ps1` each reported LOW zero graph impact because the PowerShell extractor emitted no symbols/edges. This is an extractor limitation, not proof of no consumers; direct source established their call/root flow.
- `internal/cli/package_runtime.go` file impact was CRITICAL: `19` impacted symbols / `7` files / `14` direct / `1` process / `2` linked tests.
- `buildGoRuntimePackage` symbol impact was CRITICAL: `5` impacted / `2` modules / `10` processes.
- `prepareGoSourcePackage` symbol impact was CRITICAL: `6` impacted / `2` modules / `10` processes.
- Prior `resolvePackageNativeDir` symbol impact was CRITICAL: `3` impacted / `1` module / `10` processes.
- After direct source identified the extra QA owner, Main's R40 graph was `2,100/765/0`, `122,147/168,747/20,335`, with `479` unresolved projections. `scripts/qa-child02-p2e-validation.ps1` was current but had `0` extracted symbols/edges/flows/tests and LOW `0` upstream file impact. Its direct `$LadybugRuntime` use in required-file checks, `CGO_CFLAGS`, `CGO_LDFLAGS`, `PATH`, and evidence metadata established same-invariant ownership despite that extractor limitation.

CRITICAL is retained as the actual Go package-runtime blast-radius warning. The edits stayed within the named native/root contract and its tests; Linux Docker/CI helpers and every other potential call site remained preserve-only.

## Official LadybugDB provenance and durable identities

Acquisition used the narrow native-input authority only. Source identity is the official `LadybugDB/ladybug` GitHub release/tag API:

- Version/tag: `v0.19.1`.
- Immutable release ID: `365199979`.
- Tag commit: `554c1e71158564c37a30c541a92bfc9eddc96430`.
- Published: `2026-08-04T23:34:22Z`.
- Asset ID/name: `501937224` / `liblbug-windows-x86_64.zip`.
- Archive: `8,292,678` bytes; SHA-256 `865E2C8765064BE76D41E4D786DFB0CD3AD0C258DDAF7B522FA3DA7159ECD3EF`.
- License/source identity: MIT License; copyright 2022-2025 Kùzu Inc.; LadybugDB source release.
- Archive content: four top-level files. `lbug.hpp` was verified in the extraction lane but intentionally excluded because the plan authorizes exactly three native inputs.

Durable authority `E:\Anvien\third_party\ladybugdb\v0.19.1\windows-x86_64` contains exactly `3` files / `0` subdirectories / `52,743,720` bytes:

| File | Bytes | SHA-256 |
|---|---:|---|
| `lbug.h` | 79,108 | `3D5114D0863B3DAB3B28BD2FEC97A52E6CF669213739921A01814A5BBF5525EB` |
| `lbug_shared.lib` | 32,433,956 | `B18AAFC0B712DC1C4CB9DD25F76C3828282D7D460627980E3A4B16EFCD98A955` |
| `lbug_shared.dll` | 20,230,656 | `20CBD87840483A2053CFF3FC2DB23A86DD802B8915D86509D41A4B709624CDB7` |

No provenance/license sidecar or unplanned bundle file was added to that durable directory. The acquisition archive remains debug/provenance material only at `E:\Anvien\.tmp\p6d-native-recovery-01a02e4a-260823\download\liblbug-windows-x86_64-v0.19.1.zip` with the same byte/hash identity.

## Production-first implementation chronology and source contract

Production/build behavior was changed before tests:

1. Materialized and verified the exact durable three-file bundle.
2. Corrected `scripts\ensure-ladybug-native.ps1` to accept only the repository-owned pinned authority and exact identities; it no longer downloads or treats an output cache as authority.
3. Corrected `scripts\full-build.ps1` to accept/generate a unique repo-local lane, require distinct strict-child cache/runtime roots, set E-only home/temp/Go/npm roots, validate the durable bundle, directly build the package runtime into the lane, pass all roots/native authority to the launcher, require current staged/canonical runtimes, record provenance, and run final analyze only through the current canonical runtime.
4. Corrected `internal/cli/package_runtime.go` to resolve unique strict-child build roots, validate the pinned Windows bundle by file count/size/SHA, include the exact durable bundle in prepared Go source, build into lane runtime staging, and fail closed rather than accept stale `anvien\bin` output.
5. Corrected `anvien-launcher\build.ps1` to consume explicit lane/cache/runtime/native inputs, use E-only Go/npm/Vite/temp roots, build into runtime staging, and publish only after successful staged artifacts.
6. After Main released the exact QA owner, changed only its `$LadybugRuntime` assignment from shared `.tmp\ladybug-native` to the durable authority. All unrelated Child 02 QA behavior and evidence destinations were preserved.
7. Only then updated/added the same-invariant tests.

Current Windows production/build/helper/package/launcher/QA owners contain no shared `.tmp\ladybug-native`, runtime downloader, `npm install`, `npm run`, lifecycle/package-script recovery, `Version=auto`, or stale-runtime fallback. The sole repository search hit retaining `.tmp/ladybug-native` is the unchanged non-Windows shell helper default; Windows canonical execution does not use it, and package-mode shell use passes an explicit lane-cache output root. The package code's `cache\ladybug-native` directory is a strict child of the unique lane cache and is not the prohibited shared root or authority.

## Exact Coder-owned changed manifest

Production/build/QA owners:

| Path | Bytes | SHA-256 |
|---|---:|---|
| `scripts/full-build.ps1` | 11,184 | `C4BC1DF9C1698D0FFFD3DE7BE294719D481C0E68FCA98D9E1A26944639918659` |
| `scripts/ensure-ladybug-native.ps1` | 3,479 | `E6BE6E0151C9B766D24A871A9A7C0AF3468DEF4D092551FE9A6FC8CD98A24819` |
| `internal/cli/package_runtime.go` | 21,707 | `27DC66AE94E8722D1C40A3D4751AB5F8F6466CEB9E3E6BCF3AC3DC2E4E9991A6` |
| `anvien-launcher/build.ps1` | 13,516 | `7E8825A9498E703D12344ED2EEF13A0F5A21CADA7222D9AFDF553C80C03A21FE` |
| `scripts/qa-child02-p2e-validation.ps1` | 21,761 | `55985C5587EED35B3FEF110DDF70DE8DD07DB026DDB47C3F47330148B9668C01` |

Tests added/updated after production behavior:

| Path | Bytes | SHA-256 |
|---|---:|---|
| `internal/cli/package_command_test.go` | 12,432 | `A78B5796A1ED76D04C3750356460FF70AC042473DBE8FDE883EBA5C214BFCEA7` |
| `internal/cli/package_runtime_p6d_test.go` | 8,419 | `82E4873713C5820611895B704F5C05CE44A08FA4F1DCC27897CE18F62407DE1D` |
| `scripts/test-ladybug-native-contract.ps1` | 4,591 | `E774D0A50C84B99A3C9E1B9078E31387F0FBDFF84D51103A4016FE693B233595` |

Native inputs are the three exact files in the provenance table. This report is the sole additional durable Coder handoff. No other path is claimed as Coder-owned.

## Static and focused validation

Results after production/source correctness:

- PowerShell parser: five affected/new PowerShell files, `0` parser errors, PASS.
- `gofmt -d` over the affected Go production/test files: zero output, PASS.
- `git diff --check`: exit `0` at the pre-build static gate; a final post-report recheck is external to this report's final-byte seal.
- Native helper contract command `powershell -NoProfile -ExecutionPolicy Bypass -File scripts\test-ladybug-native-contract.ps1 -LaneRoot <lane>`: valid bundle plus missing/partial/extra/alternate-authority cases, `5/5`, exit `0`.
- Current-production file-list Go test: distinct configured roots; two generated lanes do not share; outside-lane root rejection; durable bundle acceptance; alternate authority rejection; partial/tampered/extra rejection; stale-runtime fallback rejection. All named rows PASS, package output `ok`, `1.568s`. Raw log: `1,480` bytes; SHA-256 `2977809642BA04ECC710258723F3ADD926DEB64D238BE50BC825DD27390F46F1`.
- Focused normal `internal/cli` package-mode test could not enter tests: setup failed on the same absent offline module closure. Raw log: `3,415` bytes; SHA-256 `3417AE681DD88DF32822E51136D45FF4B417B1823A2C8E2976809DDE70E00B4A`. It is BLOCKED, not counted as regression PASS.

The missing/partial/extra/tampered/alternate cases prove fail-closed behavior. The configured distinct-root row and two independently generated lane rows prove unique-root propagation and non-sharing at the production function boundary.

## Rule 13 holder/lock preflight

Before the canonical build, the accepted E-only lock check reported the analyze lock free/absent. Exact executable holder inspection identified only current Anvien build/runtime holders:

- Confirmed `anvien` PIDs: `1988`, `3020`, `10032`, `12960`, `13308`, `13716`.
- Former PID `12668` was already absent.
- Only those six confirmed build-related holders were terminated. No broad process kill occurred.
- PID-gone/name-reuse/respawn checks confirmed all six absent afterward.
- Restart Manager plus exclusive-open then showed zero holders for `anvien-launcher\AnvienLauncher.exe`, `anvien-launcher\server-bundle\anvien-server.exe`, `anvien\bin\anvien.exe`, and `anvien\bin\lbug_shared.dll`.

Only this corrected holder/lock result authorized the build.

## Canonical full build: exact blocker

The real canonical command was run exactly once after accepted Rule 13 preflight:

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File E:\Anvien\scripts\full-build.ps1 -LaneRoot E:\Anvien\.tmp\p6d-native-recovery-01a02e4a-260823
```

Result: exit `1`; started `2026-08-23T19:41:27.7482288+07:00`; finished `2026-08-23T19:41:49.3307205+07:00`; elapsed `21,577 ms`.

The script first proved and printed the exact lane/cache/runtime/native roots and wrote current source/native/root provenance. It then failed at `direct Go package runtime build`. The clean lane `GOMODCACHE` had no module bytes, while exact required classes included Cobra `v1.10.2`, Hugot `v0.7.2`, go-tree-sitter `v0.25.0`, Dart/Swift/Kotlin grammars, and C/C#/C++/Go/Java/JavaScript/PHP/Python/Ruby/Rust/TypeScript tree-sitter grammars. Each lookup failed visibly with `module lookup disabled by GOPROXY=off`.

Artifacts:

- Raw log: `3,567` bytes; SHA-256 `3E6BAC1E4E7CEC2EBDA32F71B2B2AE25041CB517E2148D852E1C9115BDEB0F6C`.
- Result JSON: `432` bytes; SHA-256 `5DFA57242EC3C0AE29CDCF8665A74305AD79C2E7CDB5B770B9CB76055514C8BD`.
- Failure provenance: `7,991` bytes; SHA-256 `22082559D37D06C31179413B0EDBA07C14DDCECC7FAED3A7DEB2AD18CF1E7E02`.
- Package-runtime stage: absent; provenance `runtimeFiles` is empty.
- Lane module cache: `17` zero-byte lookup lock files and `0` module bytes.

No network/package dependency acquisition, global/shared cache, package install/script, alternative build workflow, or stale-runtime substitution was authorized or attempted. The build failure is the precise remaining dependency-authority blocker.

## Current runtime provenance and built-reader disposition

The existing canonical binary is rejected as product evidence:

- `E:\Anvien\anvien\bin\anvien.exe`: `73,756,672` bytes; SHA-256 `6BB431D00C20A6E9954039A54563855561E4A77B5D0AF323A4A38A54C7070B80`.
- Embedded VCS revision: `838f379ab00c93dfe7507603f6608bb08d61f038`; `vcs.modified=true`, not current HEAD.
- Embedded CGO paths still reference forbidden shared `E:\Anvien\.tmp\ladybug-native\v0.19.1\windows-x86_64` include/link roots.

Because the corrected build produced no binary:

- current canonical runtime: BLOCKED / absent;
- canonical final analyze through a current runtime: BLOCKED / not reached;
- built `graph-health summary/report/components/explain/files`: BLOCKED / not run;
- built `resolution-inventory` and every other actually affected reader: BLOCKED / not run;
- package/launcher publication parity: source/static contract present, runtime proof BLOCKED.

The stale executable was used only to satisfy mandatory current-source graph refresh/change detection after the build blocker. It is never promoted to current-runtime, native-parity, or built-reader evidence.

## Nearest real E:\Anvien native boundary

A direct current-source native smoke compiled `E:\Anvien\.tmp\p6d-native-recovery-01a02e4a-260823\native-smoke\main.go` against the durable header/import-library/DLL using only lane-owned cache/temp/runtime roots. It initialized a writable Ladybug database, executed `RETURN 1 AS value`, closed it, reopened read-only, and executed the same query.

Accepted confirmation: exit `0`, `2,240 ms`, with both `writeRows=[{"value":"1"}]` and `readRows=[{"value":"1"}]`. Raw confirmation: `185` bytes; SHA-256 `6F97DDF3738F5929F21EAA18A20406EAA3E56D5A0D4A3D6E324D682AA150382C`.

This proves the exact durable native inputs link/load and the repository's current `internal/lbugnative` write/read boundary works against a real E-only Ladybug database. It does not prove the absent canonical built runtime or its readers.

## Final current-source analyze and detect-changes

The only accepted final graph rerun used non-reserved `$homeRoot`, with HOME/USERPROFILE/TEMP and all relevant runtime roots under the declared lane. It invoked the stale executable only as a graph engine over current repository bytes:

```text
E:\Anvien\anvien\bin\anvien.exe analyze E:\Anvien --force --json
```

Result: exit `0`; `595,490 ms`; `2,101` scanned / `765` parsed code / `0` failed; `122,163` nodes / `168,763` relationships / `20,335` dependency edges / `479` unresolved file projections. Raw log: `8,678` bytes; SHA-256 `42282FDF2526512A681BB2027EFE7D339EE50329CA5E69B857E79D64325D2953`.

Mandatory final command after final code/ledger bytes:

```text
E:\Anvien\anvien\bin\anvien.exe detect-changes --repo E:\Anvien --scope all --json
```

Result: exit `0`; `23,990 ms`.

- Overall risk: MEDIUM; file-layer risk: HIGH.
- Affected: `3` symbols / `8` files / `3` processes.
- Changed: `407` symbols / `13` files.
- Changed functional areas: CLI `181`, documentation `46`, graph-health `180`.
- ResolutionGap delta: `206` entities / `206` occurrences (`53` unresolved access, `131` unresolved call, `22` unresolved type reference).
- Current Resolution Health: total gaps `0`, nodes with gaps `0`.
- Semantic app-layer completeness: `122,163/122,163`; functional-area completeness: `122,163/122,163`.
- Raw log: `1,268,753` bytes; SHA-256 `AB22D76261943FCF588B83FC5DD13AD5497170F7EE53EF1062F91DB9A2099DB8`.

The `595,490 ms` analyze duration, `23,990 ms` detect duration, `21,577 ms` failed-build duration, and `2,240 ms` native confirmation are validation timings, not a P6-E performance baseline or speedup claim. The graph inventory counts are benchmarkable capacity observations handed to Main; this Coder had no ledger write authority.

## Reserved-variable incidents and evidence exclusion

Two case-insensitive PowerShell reserved-variable mistakes make procedural integrity non-clean:

1. An earlier holder/lock wrapper assigned `$Home`. Its inherited-home invocation and all output were excluded. The accepted Rule 13 evidence came only from the corrected E-rooted holder/lock invocation.
2. The first final-analyze wrapper assigned `$HOME`. Session/PID `50850` and child PID `2328` were terminated and verified gone; no descendant remained. The E analyze lock was present but exclusive-open proved it free. Its output and graph were excluded completely. Exactly one corrected rerun used `$homeRoot` and produced the final graph above.

No path outside `E:\Anvien` was opened, read, followed, or cleaned to investigate either invocation. Therefore absence of any inherited-home side effect outside E cannot be proved inside this authority. No excluded output was used as source, runtime, graph, build, test, or detect evidence.

## Target/oracle/boundary gate

`E:\cheapapp.org` remained prohibited and was never accessed. Therefore exact target pre-state, `Promise` / `Math.max` / `Math.min` `3/3`, independent TypeScript oracle, target graph, and target post-state are explicitly `BLOCKED / UNPROVED`, never inferred from Anvien source, aggregate graph counts, or the native smoke.

## Cleanup and lane inventory

No protected/shared/quarantined root was used or cleaned. The recovery lane is intentionally retained as debug/acquisition/build-failure evidence for Main/Supervisor; it is not product runtime or dependency authority.

Current lane inventory before this report: `1,942` files / `333` directories / `174,910,370` bytes / `19` zero-byte files / `0` reparse points. Key roots:

- `cache\go-build`: `1,901` files / `110,185,601` bytes.
- `cache\go-mod`: `17` files / `0` bytes, all lookup locks and no module closure.
- `download`: `1` file / `8,292,678` bytes.
- `extract`: `4` files / `53,086,494` bytes; includes excluded `lbug.hpp` only in the lane.
- `home`: `4` files / `16,729` bytes.
- `native-smoke`: `1` source file / `1,062` bytes.
- `runtime`: `3` files / `2,039,607` bytes: failure provenance plus two native smoke databases; no package runtime.
- `temp`: empty.

Test-owned synthetic helper roots were removed by their narrowly verified `finally` cleanup. No broad or alternate cleanup was attempted. Main must decide later disposition after independent verification; the current lane must not be mistaken for a reusable cache.

## Exact preserved dirty boundary

Current non-Coder-owned dirty paths were preserved without overwrite, restoration, deletion, staging, or absorption:

- Owner deletion: `CONTRIBUTING.md`.
- Main/Planner: roadmap and four Child 06 ledgers.
- Cleared semantic repair: `internal/graphhealth/diagnostics.go`, `internal/graphhealth/p6d_outcome_projection_test.go`.
- Main handoffs: `reports/Investigation/rp_main_260823_164621_orchestration_rotation_handoff.md`, `rp_main_260823_174805_orchestration_rotation_handoff.md`, `rp_main_260823_184842_orchestration_rotation_handoff.md`, `rp_main_260823_194346_orchestration_rotation_handoff.md`.
- Supervisor report: `reports/Supervisor/rp_supervisor_260823_164652_by_gpt-5_p6e_current_plan_acceptance_rereview.md`.
- Sealed prior Coder blocker: `reports/coder/rp_coder_260823_171110_by_gpt-5_p6d_real_recovery_blocked.md`.

Post-R40 inspect-only plan identities at handoff:

- roadmap: `38,623` bytes / `FD3542BC9B92C8BB8A763D87C9E25CD39BFB859455861E88115F31DFA54DD6B2`;
- Child 06 plan: `115,051` bytes / `8094D0E737188FFE845F9FB2269DC3F273161922D507376336878535569FD765`;
- evidence: `150,711` bytes / `77AE9E6FE3FB19F7F12B4BEED435082B47DB1ACB29B68599F2F7DA3EAC57F1B7`;
- benchmark: `31,905` bytes / `4C1173FEF0A2C76F9C2AE04F84E58632A487918D0E0DA51979498793ED165736`;
- actual status: `75,627` bytes / `98540E388F748990BC12C11B24537CF9A78E781CADA203FA4DF17E9B543018D0`.

The final Git recheck must retain branch/HEAD/parent above, empty index, the full dirty manifest plus this one report, and zero `git diff --check` errors.

## Blocking conditions and smallest next action

Acceptance remains blocked by four exact gates:

1. A clean repository- or lane-owned dependency closure for the exact existing `go.mod`/`go.sum` set is absent; current authority forbids network/package acquisition and shared/global cache reuse.
2. The corrected canonical build therefore produced no current runtime, so package/launcher publication, final current-runtime analyze, native-through-runtime, and affected built-reader parity are unproved.
3. Target/oracle/pre-post proof is forbidden under the current `E:\cheapapp.org` boundary.
4. Reserved `$Home`/`$HOME` incidents prevent a clean procedural-integrity claim even though excluded outputs were not reused and the corrected invocations passed their local gates.

Smallest next responsible action for Main/Owner: independently verify this report and source/artifact identities, then decide an explicit clean E-only dependency-closure authority for the unchanged module set. Do not authorize shared/global cache reuse implicitly. Once a clean dependency closure exists, run a fresh Coder recovery lane through Rule 13, the corrected canonical build, current runtime/native/built-reader validation, and a new clean final graph/detect without reserved-variable wrappers. Target work still requires separate explicit authority. Main, using planner, owns any ledger update; after complete current-runtime proof, route a fresh independent Supervisor. If resolving the dependency closure would require architecture/scope redesign, Main must route that exact blocker to a visible Architect lane rather than expanding P6-D here.

No P6-D checkbox, P6-E state, stage, or commit is authorized by this handoff.

## E2E verification

```text
E2E Verification:
  [PASS] Source/native authority: exact durable 3/3 bundle; helper/package/root fail-closed matrix
  [BLOCKED] Compiled: canonical scripts/full-build.ps1 -> exit 1 on absent clean offline Go module closure
  [BLOCKED] Runtime: no current package/canonical runtime was produced
  [PASS, bounded] Native happy path: current internal/lbugnative write + read-only reopen against real E-only Ladybug DB
  [BLOCKED] Built readers: current graph-health/resolution-inventory reader parity cannot run without current runtime
  [BLOCKED] Target edge/oracle/boundary: E:\cheapapp.org forbidden and never accessed
```

Final handoff disposition: `BLOCKED`.

## External final-byte seal protocol

This report is the single durable Coder handoff. Its final bytes, LF/CR counts, strict UTF-8 validity, BOM state, SHA-256, and created/last-write timestamps are measured only after the final write and supplied externally to Main with the handoff. The final SHA-256 is intentionally not embedded in its own byte stream because doing so would mutate the artifact being sealed; no sidecar is created.
