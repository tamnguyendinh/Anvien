# System Architect Report: P6-D Clean E-only Go Vendor Closure

Final architecture decision: `CHECKED_IN_EXACT_GO_VENDOR_CLOSURE`

Final status: `BLOCKED_NO_COMPLIANT_SOURCE`

Next responsible: Visible Main task `01a02ec4-11a9-7e93-b7f9-4a066d0db128`

## 1. Kết luận

Kiến trúc tối thiểu được chọn là một Go vendor tree chính xác, được check in tại root module `E:\Anvien\vendor`, sinh từ `go.mod`/`go.sum` hiện tại và bắt buộc được dùng bằng explicit `-mod=vendor`.

Đây là quyết định kỹ thuật cuối cùng, không phải menu. Build runtime không dùng module cache làm authority, không dùng file-proxy ở runtime, và không có fallback. `GOMODCACHE` tiếp tục là derived cache riêng của từng lane; `GOPROXY=off` tiếp tục giữ nguyên.

Một blocker vật lý duy nhất còn lại: trong toàn bộ surface E-only được phép hiện không có source bytes của Go modules để materialize `vendor/`. `go.sum` có checksum nhưng không thể tái tạo source bytes. Vì vậy kiến trúc đã được quyết định, nhưng Coder không thể tạo vendor tree hợp lệ từ bytes hiện có.

Smallest non-technical authority action: Main phải cung cấp một checksum-verifiable E-only module-source input cho đúng unchanged module set dưới lane materialization riêng trong `E:\Anvien\.tmp`; sau đó Planner refresh bốn ledger và route một fresh Coder theo handoff ở cuối báo cáo. Không cần Owner chọn lại cơ chế.

## 2. Boundary và mode

- Active slice: P6-D; P6-E giữ nguyên LOCKED.
- Role: advice-only System Architect; không implementation, test edit, ledger edit, stage/commit/push.
- Sole read/write workspace: `E:\Anvien`.
- Output duy nhất: báo cáo này.
- Graph claim mới: không có. Graph đã stale sau R41; `anvien analyze --force` bị boundary hiện tại cấm vì mutate product state.
- Commit reference: N/A — user authority cấm stage/commit.

## 3. Verified facts

### 3.1 Git và module authority hiện tại

- Read-only Git evidence: branch `master`, HEAD `741335f7efa5ad215ec28361d88c462f431565ab`, parent `838f379ab00c93dfe7507603f6608bb08d61f038`, index rỗng. Owner deletion `CONTRIBUTING.md` được giữ nguyên.
- `git diff -- go.mod go.sum` exit `0`; hai module authority files không đổi.
- `E:\Anvien\go.mod` là root module, yêu cầu Go `1.26.3`, có `6` direct và `39` indirect requirements (`45` total): `go.mod:1-53`.
- File identities:
  - `go.mod`: `2,392` bytes; SHA-256 `C203E25E0A83A583E89A9D8F8E65BC0881052B3075E46915F90A4FB6433BD249`.
  - `go.sum`: `16,065` bytes / `172` lines / `88` unique module-version pairs; SHA-256 `5434F39B08F50F1C53A71333E2E971CFCDE0E80BCA7DD028C00E31B9ACCAEC73`.
- Cả `17/17` module versions xuất hiện trong failed canonical lookup đều có đúng một content `h1` và một `/go.mod h1` record trong `go.sum`. Đây là checksum/provenance evidence, không phải source closure.

### 3.2 Không có compliant module bytes

Direct read-only inventory:

- `E:\Anvien\vendor`: absent.
- `E:\Anvien\third_party\go`: absent.
- `E:\Anvien\anvien\go-src`: absent; tracked packaged Go source paths: `0`.
- Git object/history search cho `vendor/modules.txt`, `third_party/go`, `go-proxy`, hoặc gomod bundle: `0` hits.
- Tracked dependency artifacts ngoài `go.mod`/`go.sum`: không có `vendor/**`, module `.zip`, `.mod`, `.info`, hoặc repository-owned Go proxy/cache bundle.
- Exact active lane cache `E:\Anvien\.tmp\p6d-native-recovery-01a02e4a-260823\cache\go-mod`: `17` files, `17` zero-byte `.lock`, `0` non-lock files, `0` module bytes.
- Exact active lane dependency-candidate inventory chỉ có LadybugDB archive `liblbug-windows-x86_64-v0.19.1.zip`; không có Go module archive/source.

Kết luận verified: current allowed bytes không thể materialize vendor source tree. Build cache objects không phải module source authority, không cung cấp licenses/source review, phụ thuộc toolchain/action IDs và không thể thay cho Go module closure.

### 3.3 Canonical build/runtime contract

- `scripts/full-build.ps1:117-156` tạo một unique strict-child lane với cache/runtime roots tách biệt và durable Ladybug authority.
- `scripts/full-build.ps1:172-201` đặt HOME/temp/Go state dưới lane, `GOPROXY=off`, `GOSUMDB=off`, `GOTOOLCHAIN=local`.
- First build boundary là `go run ..\cmd\anvien package build-runtime` tại `scripts/full-build.ps1:206-221`; runtime phải tồn tại trong lane trước khi launcher/publication chạy.
- Canonical launcher và final current-runtime analyze chỉ chạy sau package runtime thành công: `scripts/full-build.ps1:226-242`.
- `internal/cli/package_runtime.go:74-140` build `./cmd/anvien` bằng `-tags ladybugdb`, CGO và lane-local caches; `package_runtime.go:141-168` chỉ publish sau staged build thành công.
- `internal/cli/package_runtime.go:172-233` hiện package `go.mod`, `go.sum`, `cmd`, `internal`, helper và Ladybug inputs, nhưng không copy dependency closure.
- `anvien-launcher/build.ps1:215-235` cũng ép clean lane cache và offline Go environment; root CLI build nằm tại `build.ps1:268-277`. Hai launcher submodules ở `anvien-launcher/src/go.mod` và `anvien-launcher/server-wrapper/go.mod` chỉ dùng standard library, không có external requirements.
- `anvien/package.json:37-49` publish `bin` và `go-src`; do đó packaged-source fallback cũng phải mang cùng vendor authority, không được chỉ làm canonical repo build PASS.

### 3.4 Existing P6-D plan authority

- P6-D hiện sở hữu `scripts/full-build.ps1`, `scripts/ensure-ladybug-native.ps1`, `internal/cli/package_runtime.go`, `anvien-launcher/build.ps1` và exact QA consumer: canonical plan `:468-472`.
- Plan khóa build tại clean dependency closure và cấm rerun cho tới khi có explicit E-only authority cho unchanged `go.mod`/`go.sum`: plan `:503`.
- Plan xác nhận `17` locks / `0` module bytes / no runtime: plan `:519-520`.
- New roots `vendor/**`, `third_party/go-vendor/**` và new verification helper/test chưa nằm trong editable manifest. Vì vậy cần four-ledger Planner refresh trước code, nhưng không cần phase/slice/architecture redesign: toàn bộ vẫn là một phần của existing P6-D và cùng commit boundary ở plan `:543`.

### 3.5 Blast radius

- Existing sealed impact evidence cho `internal/cli/package_runtime.go`: CRITICAL, `19` impacted symbols / `7` files / `14` direct / `1` process / `2` linked tests.
- `buildGoRuntimePackage`, `prepareGoSourcePackage`, và native/package path owners đều có CRITICAL upstream impact theo sealed Coder report.
- PowerShell graph impact `LOW 0` là extractor limitation, không phải no-consumer proof.
- New vendor/artifact/helper paths chưa có graph nodes. Fresh graph/file-detail/impact phải được thực hiện bởi implementation lane sau khi Planner refresh và khi authorized `anvien analyze --force` có thể chạy E-only.

## 4. Architecture rule và implementation suggestion

### Architecture / invariant

`E:\Anvien\vendor` là duy nhất Go dependency source authority cho root module build. Build phải fail closed nếu authority thiếu, thừa, drift, không khớp `go.mod`/`go.sum`, chứa reparse/symlink, hoặc được yêu cầu từ path khác.

`GOMODCACHE` và `GOCACHE` chỉ là derived lane state. Không cache nào — kể cả cache đầy đủ — được dùng làm input authority hoặc fallback.

### Exact durable layout

```text
E:\Anvien\
├── go.mod                                      # unchanged
├── go.sum                                      # unchanged
├── vendor\
│   ├── modules.txt
│   └── <Go-generated vendored package trees...>
├── third_party\go-vendor\
│   ├── manifest.v1.json
│   └── licenses\
│       └── <escaped-module>@<version>\<LICENSE|NOTICE|COPYING...>
├── scripts\verify-go-vendor.ps1
└── scripts\test-go-vendor-contract.ps1
```

`vendor/` phải là output nguyên vẹn của standard `go mod vendor` trên exact unchanged module inputs; không được hand-prune chỉ còn 17 first failures. `17` chỉ là observed first-failure set, không phải complete closure denominator.

### `manifest.v1.json` minimum contract

Manifest không tự-hash chính nó. Nó phải chứa:

- `schemaVersion = 1`, `authorityKind = "go_vendor"`;
- exact `goVersion = "go1.26.3"` và generation command identity;
- exact `goModSha256` và `goSumSha256` nêu trên;
- `vendorModulesTxtSha256`;
- canonical ordered module rows: module path, version, `go.sum` content hash, `go.mod` hash;
- canonical ordered file rows cho toàn bộ `vendor/**` và `third_party/go-vendor/licenses/**`: relative path, bytes, SHA-256;
- canonical tree SHA-256 trên ordered file rows;
- license disposition cho `100%` modules trong `vendor/modules.txt`: copied license/notice paths hoặc một explicit reviewed exception. Missing/unknown disposition fail closed;
- source-materialization bundle identity và two-run equality result.

Timestamps, absolute lane paths và machine-specific values không được tham gia deterministic manifest bytes.

## 5. Vì sao exact checked-in vendor thắng

Vendor là Go-native path ngắn nhất cho source hiện tại: root module đã có exact `go.mod`/`go.sum`, canonical commands đều build local packages, launcher submodules không có external deps, và build đã ép `GOPROXY=off`. Chỉ cần materialize standard vendor tree, validate nó trước first Go process, và truyền explicit `-mod=vendor`. Không cần module proxy server, file-URL normalization trên Windows, module-cache seeding/copy, hoặc runtime extraction phase.

Vendor cũng giữ source/CGO grammar files và licenses visible trong repository, phù hợp review/security hơn một opaque cache. Its cost là repository working-tree size và việc Go không dùng `go.sum` để rehash từng vendored file; manifest per-file SHA + generation from checksum-verified module input closes đúng gap đó. Exact size hiện unavailable vì source bytes không tồn tại; Coder phải record it at materialization, không estimate.

Rejected alternative `repository-owned file proxy`: technically valid, và Go có thể verify module zips bằng unchanged `go.sum`, nhưng nó lưu whole module archives, cần Windows file-URI plumbing ở full-build/package/launcher, repopulates each lane cache, hides source/license review inside zips, và làm packaged `go-src` lớn/phức tạp hơn. Không tối thiểu cho boundary hiện tại.

Rejected alternative `immutable GOMODCACHE seed`: cache layout là derived/mutable implementation detail, có locks/extracted trees/ziphash state, khó read-only/lane-isolate và provenance yếu hơn. Direct repository cache reuse hoặc copy từ shared/protected/quarantined cache đều bị cấm.

Local replacements/go.work, precompiled archives, stale runtime và global caches bị loại: chúng bypass hoặc làm mờ `go.sum`, phụ thuộc toolchain/host, thay contract, hoặc vi phạm explicit bans.

## 6. Canonical flags và environment

Mọi root-module Go invocation phải sanitize inherited Go controls và dùng explicit vendor mode:

```text
GOENV=off
GOWORK=off
GOFLAGS=
GOTOOLCHAIN=local
GOPROXY=off
GOSUMDB=off
GOPRIVATE=
GONOPROXY=none
GONOSUMDB=none
GOINSECURE=
GOVCS=*:off
```

Exact command changes:

- `scripts/full-build.ps1`: trước first Go command, chạy `scripts/verify-go-vendor.ps1`; sau đó dùng:
  - `go run -mod=vendor ..\cmd\anvien package build-runtime`.
- `internal/cli/package_runtime.go`: runtime build dùng:
  - `go build -mod=vendor -tags ladybugdb -trimpath -ldflags="-s -w" -o <lane-stage> ./cmd/anvien`.
- `anvien-launcher/build.ps1`: root CLI build dùng:
  - `go build -mod=vendor -tags ladybugdb -trimpath -ldflags="-s -w" -o <lane-stage> .\cmd\anvien`.
- Launcher/server-wrapper standard-library-only module builds không dùng root `-mod=vendor`; chúng vẫn chạy với proxy/VCS disabled.
- `prepareGoSourcePackage` phải copy exact `vendor/**`, `third_party/go-vendor/**` và verifier cùng với existing `go.mod`/`go.sum`/source/native files. `buildGoRuntimePackage` phải resolve/verify vendor relative to its actual `sourceRoot`, nên repo root và packaged `go-src` dùng cùng contract.

`GOENV`, `GOWORK`, `GOFLAGS`, `GOPRIVATE`, `GONOPROXY`, `GONOSUMDB`, `GOINSECURE`, `GOVCS` phải được thêm vào save/restore lists ở both PowerShell owners. Không được kế thừa hidden host configuration.

## 7. Failure/publication semantics

- Vendor verifier chạy trước first `go run`; invalid authority không được compile bất kỳ candidate runtime nào.
- Missing/extra/tampered file; changed `go.mod`/`go.sum`; mismatched `modules.txt`; wrong Go version; missing license disposition; reparse point; alternate root; hoặc manifest drift => visible nonzero exit.
- Không tự chạy `go mod vendor`, `go mod download`, network, package install hoặc package script trong canonical build.
- Không `-mod=mod`, không `-mod=readonly` fallback, không `GOPROXY` list có `direct`, dấu phẩy hoặc pipe fallback.
- Build outputs ở lane staging; canonical/package/launcher destinations chỉ publish sau complete staged success, giữ current P6-D semantics.
- Existing/stale runtime không bao giờ là fallback nếu vendor/build fail.
- Failed/partial vendor materialization không được đặt tại durable root; generation dùng lane temp và publish atomically only after complete verification.

## 8. One-time materialization lifecycle

Đây là prerequisite, không phải runtime architecture.

1. Main cung cấp exact checksum-verifiable module-source input tại một unique root theo pattern:
   `E:\Anvien\.tmp\p6d-go-vendor-materialize-<unique>\input-proxy`.
2. Fresh Coder dùng exact Go `1.26.3`, isolated E-only HOME/TEMP/GOCACHE/GOMODCACHE/GOPATH, no shared cache, no alternate worktree. Materialization source chỉ được dùng trong lane này.
3. Verify every required module `.zip`/`.mod` against unchanged `go.sum`; generate vendor bằng standard `go mod vendor` vào lane staging.
4. Run the generation twice from clean lane roots; canonical file list, bytes và tree digest phải equal `2/2`.
5. Extract/copy license and notice files; require `100%` module disposition; generate `manifest.v1.json`.
6. Prove `go.mod`/`go.sum` hashes unchanged and staged vendor passes contract tests.
7. Atomically materialize durable `vendor/**` and `third_party/go-vendor/**`. Lane input/cache/output remain derived evidence and are removed or explicitly dispositioned after seal; they never become build authority.

Current physical blocker is exactly Step 1. No allowed current byte source can satisfy it.

## 9. Exact Coder write scope after prerequisite + Planner refresh

Add:

- `vendor/**`;
- `third_party/go-vendor/manifest.v1.json`;
- `third_party/go-vendor/licenses/**`;
- `scripts/verify-go-vendor.ps1`;
- `scripts/test-go-vendor-contract.ps1`.

Edit:

- `scripts/full-build.ps1`;
- `internal/cli/package_runtime.go`;
- `internal/cli/package_command_test.go`;
- `internal/cli/package_runtime_p6d_test.go`;
- `anvien-launcher/build.ps1`.

Preserve-only unless fresh evidence expands authority:

- `go.mod`, `go.sum`;
- `scripts/ensure-ladybug-native.ps1`;
- `scripts/qa-child02-p2e-validation.ps1`;
- durable Ladybug bundle;
- cleared graph-health production/test bytes;
- P6-E owners, target, ledgers, roadmap and all protected/user-owned paths.

Existing P6-D owns the three production/build owners being edited, but not the new vendor/manifest/helper/test roots. Main must perform a four-ledger Planner refresh before Coder edit; this is an owner-map extension inside P6-D, not a new slice or commit boundary.

## 10. Verification matrix

### Materialization/integrity

- Exact input: all module `.zip` and `.mod` hashes match unchanged `go.sum`.
- Determinism: two clean materializations equal by ordered paths, bytes, file hashes and tree digest.
- `vendor/modules.txt` is consistent with unchanged `go.mod`.
- Manifest covers every durable file, every module, and every license disposition.
- Negative rows: missing, extra, tampered, wrong path/version/hash, reparse point, changed mod/sum, missing license, alternate vendor root all fail before Go execution.

### Clean canonical build

- Rule-13 holder/lock preflight passes.
- Start with new unique lane, empty `GOMODCACHE`, `GOPROXY=off`, VCS disabled and no inherited Go config.
- First `go run -mod=vendor` succeeds without module-cache authority/network.
- Package runtime and root launcher CLI builds use `-mod=vendor`; standard-library launcher modules remain independent.
- `GOMODCACHE` contains no acquired module source needed by success; no C/global/shared/protected/quarantined path appears in logs/provenance/binary build info.
- Runtime VCS/source/native/vendor manifest identities are current and exact.

### Packaging/publication

- `prepare-go-source` contains exact vendor + manifest/licenses/verifier.
- Build from packaged `go-src` succeeds in a clean lane with no repository parent dependency.
- Missing/tampered packaged vendor fails closed and never accepts stale `bin/anvien.exe`.
- Launcher/package/canonical outputs publish only after staged success.

### P6-D continuation gates

- Current-runtime Ladybug write/read and built graph-health/readers pass with zero fallback/skips.
- Full canonical build completes before validation.
- Clean final graph/detect and fresh independent Supervisor follow existing plan.
- Target remains locked until Main separately releases its existing target boundary; this architecture report grants no target access.

## 11. Artifact disposition

- Durable/commit candidate: `vendor/**`, `third_party/go-vendor/**`, verifier/tests, scoped build/package edits, existing P6-D candidate, four refreshed ledgers and accepted reports at the single P6-D commit boundary.
- Derived/temporary: materialization input proxy, lane GOMODCACHE/GOCACHE/GOPATH/TEMP, generation copies, failed vendor candidates and logs. Never runtime authority; cleanup/disposition must be exact.
- Preserved: Ladybug 3/3 bundle, Owner deletion, protected reports/roots, stale runtime only as rejected historical artifact, unchanged `go.mod`/`go.sum`.

## 12. Bounded fresh-Coder handoff

```text
P6-D GO VENDOR MATERIALIZATION + CANONICAL BUILD UNBLOCK

Preconditions supplied by Main:
1. Four Child 06 ledgers refresh exact write scope named in the Architect report.
2. A checksum-verifiable Go module source input exists under one unique
   E:\Anvien\.tmp\p6d-go-vendor-materialize-<unique>\input-proxy root.
3. No target authority; P6-E stays locked.

Implement exactly the selected checked-in vendor architecture:
- keep go.mod/go.sum byte-identical;
- materialize standard `go mod vendor` output at root vendor/ from verified input;
- add third_party/go-vendor/manifest.v1.json and complete license inventory;
- add reusable vendor verifier + contract test;
- update full-build, package-runtime/packaged-source, and launcher root CLI commands to
  sanitize Go env and use explicit -mod=vendor;
- preserve GOPROXY=off/GOSUMDB=off/GOTOOLCHAIN=local and all unique E-only lane roots;
- do not use shared/global/protected/quarantined cache, C:, alternate worktree,
  E:\cheapapp.org, stale runtime, implicit network, installs/package scripts, or fallback.

Code/build contract first, tests after. Before edits, satisfy fresh authorized Anvien
analyze/file-detail/impact requirements; report CRITICAL/HIGH blast radius. Run Rule 13,
the canonical full build, current-runtime/native/built-reader gates, fresh detect and exact
cleanup disposition. Do not stage/commit/target-access; return one sealed Coder report to Main.

If the exact module-source input is absent or one required zip/mod fails go.sum verification,
stop immediately with BLOCKED_NO_COMPLIANT_SOURCE. Do not search banned caches or fetch network.
```

## 13. Decision ledger

- Selected: exact checked-in Go vendor closure.
- Runtime authority: root `vendor/` only.
- Integrity authority: unchanged `go.mod`/`go.sum` + `third_party/go-vendor/manifest.v1.json` + complete license inventory.
- Cache authority: none; caches are derived and lane-local.
- Plan consequence: four-ledger owner-map refresh inside existing P6-D; no new slice/gate/commit.
- Next implementation owner: one fresh Coder after Main satisfies the single physical materialization prerequisite and Planner refresh.
- Residual architecture question: none.
- Physical blocker: module source bytes absent from all currently allowed E-only surfaces.
- Final status: `BLOCKED_NO_COMPLIANT_SOURCE`.

