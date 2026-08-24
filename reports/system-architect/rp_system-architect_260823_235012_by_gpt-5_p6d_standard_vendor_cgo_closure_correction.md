# Báo cáo Kiến trúc Hệ thống: Hiệu chỉnh P6-D Standard-Vendor CGO Closure

Trạng thái cuối của lane: `READY_FOR_MAIN`

Quyết định duy nhất: `STANDARD_VENDOR_PLUS_CHECKSUM_SOURCE_ROOT_OVERLAY_AND_ONE_REVIEWED_PATCH`

Lane: Kiến trúc hệ thống read-only cho P6-D/R42

Người chịu trách nhiệm tiếp theo: Main Orchestration hiển thị `01a02f61-5384-7410-94e3-6421a2817ab0`

## 1. Kết luận niêm phong

Kiến trúc cũ “output nguyên vẹn của `go mod vendor` là toàn bộ dependency source authority” bị bằng chứng compile thực tế bác bỏ. `go mod vendor` vẫn là baseline bắt buộc, nhưng không truy vết các C preprocessor include từ vendored Go binding sang các thư mục nằm ngoài package directory. Vì vậy baseline 45 module/1.578 file có thể vượt qua verifier hiện tại nhưng vẫn không đóng được CGO compile.

Correction tối thiểu, fail-closed và có thể triển khai là:

1. Sinh deterministic standard-vendor baseline từ exact unchanged `go.mod`/`go.sum`, Go `1.26.3`, hai clean E-only lane, equality `2/2`.
2. Từ đúng checksum-verified module zip/source của baseline, machine-scan CGO preamble và C/C++ include theo fixed-point; với mỗi active module-local target bị baseline bỏ sót, copy toàn bộ parent source root của target vào đúng module-relative location dưới `vendor/`.
3. Áp dụng đúng một reviewed deterministic patch cho `github.com/tree-sitter/tree-sitter-go@v0.25.0`: xóa block include scanner đang tham chiếu một file không tồn tại, chỉ sau khi toàn bộ guard về binding, grammar, parser và zip/source đều khớp chính xác.
4. Manifest phải tách riêng standard baseline, supplemental overlay, reviewed patch và final tree; verifier phải chứng minh cả bốn partition trước first Go process.

Không đổi `go.mod`/`go.sum`; không cần file-proxy ở build time; không có cache/network/runtime fallback; không có lựa chọn còn lại cho Owner. `READY_FOR_MAIN` chỉ có nghĩa correction kiến trúc đã đóng và sẵn sàng để Main refresh plan rồi trả delta cho Coder hiện hữu. P6-D vẫn `REJECT/BLOCKED` cho đến khi implementation, full build và toàn bộ acceptance gate hiện hữu PASS. P6-E vẫn `LOCKED` đúng một U1-U4 slice.

## 2. Phạm vi, quyền hạn và phương pháp

- Đã đọc FULL RAW đến EOF: `E:\Anvien\AGENTS.md`, `working-rules`, `System-Architect`, báo cáo kiến trúc cũ, diagnosis JSON và seal Main.
- `System-Architect` áp dụng vì blocker làm sai giả định supply-chain/build authority và cần một quyết định trade-off, invariant cùng machine-checkable contract mới.
- Đây không phải Mode 1 SPEC authoring hoặc Mode 2 execution planning. Quyền hạn hiện tại cấm sửa SPEC/plan và chỉ cho phép đúng một lane report.
- Không dùng `governance-rule-guard`.
- Không dùng Anvien graph; không có graph claim mới. Do không sửa production symbol/file, không cần chạy `anvien analyze --force`, `file-detail`, `impact` hay `detect-changes` trong lane này.
- Không chạy build, offline compile mới, test, QA, stage, commit hoặc push.
- Không sửa production, test, vendor, manifest, script, plan, ledger, roadmap hay báo cáo hiện hữu.
- Commit reference: `N/A`; exact user authority cấm stage/commit/push, ưu tiên hơn commit rule mặc định của skill.

## 3. Nguồn và identity đã đọc

### 3.1 Authority bền vững

- Báo cáo kiến trúc cần hiệu chỉnh: `E:\Anvien\reports\system-architect\rp_system-architect_260823_210034_by_gpt-5_p6d_clean_e_only_go_vendor_closure.md`; `20.868` bytes; SHA-256 `C2A20A72F94A971317E54A4DF77973E5CA4DC8516221FB72D43DD589BC75C776`.
- Main seal/context: `E:\Anvien\reports\Investigation\rp_main_260823_232232_orchestration_rotation_handoff.md`; `10.939` bytes; SHA-256 `CA153525295BC34B17CC17D690B3E886781D86EC1DA58A7857334486B551A22C`.
- `E:\Anvien\go.mod`: `2.392` bytes; SHA-256 `C203E25E0A83A583E89A9D8F8E65BC0881052B3075E46915F90A4FB6433BD249`.
- `E:\Anvien\go.sum`: `16.065` bytes; SHA-256 `5434F39B08F50F1C53A71333E2E971CFCDE0E80BCA7DD028C00E31B9ACCAEC73`.

### 3.2 Scratch evidence được đọc, không được nâng thành runtime authority

- `E:\Anvien\.tmp\p6d-go-vendor-materialize-01a02ef3-260823\evidence\standard-vendor-incomplete-diagnosis.json`; `23.861` bytes; SHA-256 `B0B6C0A281BA41EC7AF13105204029F3F8521C2D7EEBFF7344B3DED308FB1EA5`.
- `E:\Anvien\.tmp\p6d-go-vendor-materialize-01a02ef3-260823\evidence\standard-vendor-direct-includes.json`; `14.876` bytes; SHA-256 `B45DAF1DEE863BB137A23AFA6B74DB00E06C1506C7D7FAB9C99F198ECC680467`.
- Clean checksum-verified source root: `E:\Anvien\.tmp\p6d-go-vendor-materialize-clean-01a02ef3-260823-09\acquire\state\gomodcache`.
- Newest compile root: `E:\Anvien\.tmp\p6d-go-vendor-materialize-01a02ef3-260823\offline-production-compile-20260823T162547835Z-d15a0471247346cb9a5f2c097d8ad8d8`.
  - `01-verifier.log`: `648` bytes; SHA-256 `289DBB9704D5732D72CB4A9DD337D9D0D7BC460D2573CB5D324B85AE175C8532`.
  - `02-go-version.log`: `35` bytes; SHA-256 `1A29B12D3D0ABD08CA1431FC930D17C3AB5220BB116622806D17EEA91D74E0CB`.
  - `03-offline-compile.log`: `5.480` bytes; SHA-256 `7B9828E608FAAE246ACC3FFB83C482AA682BD4F62B2EA205E5B25A93DFD68374`.

### 3.3 Current implementation surface đã kiểm tra

- `scripts/materialize-go-vendor.ps1`: `38.852` bytes; SHA-256 `4B11DE923C65CF5348E937CDD215F34B98D1B64AF87EB44A5D81A74B9A45EE9B`.
- `scripts/verify-go-vendor.ps1`: `20.634` bytes; SHA-256 `1A8A17521A2AC7F98299C8782B8D7A135BAE039BDB57A8F62B88B6262B23AFFC`.
- `scripts/full-build.ps1`: `12.104` bytes; SHA-256 `9E655CA833E5E3190587B85FC161096BE4CB492BA151E85DAB4638FEE6885CC7`.
- `internal/cli/package_runtime.go`: `24.817` bytes; SHA-256 `25B7519C7222A4D3BAC2C8D6FD55A4B44F7893F7DF715B924A6A8CA37E23D0EC`.
- `anvien-launcher/build.ps1`: `14.339` bytes; SHA-256 `C8B0BCB482745D18539806BBBC70EF7D0B56ED3D8D6F47F136B3930316E92799`.

Các production diff hiện hành đã có đúng hướng: verifier-before-first-Go, explicit `-mod=vendor`, offline controls, packaged `go-src` copy toàn bộ `vendor/**` và `third_party/go-vendor/**`, cùng manifest identity trong packaged runtime metadata. Correction này không thiết kế lại ba owner đó; nó làm cho authority mà chúng đang consume thực sự compile-closed.

## 4. Sự thật đã xác minh

### 4.1 Standard vendor không đủ trên source hiện tại

Materialization chuẩn đã chứng minh:

- `45/45` vendored module có content sum và `/go.mod` sum khớp unchanged `go.sum`.
- Standard `go mod vendor -o <output>` sinh `1.578` file và equality `2/2`; standard tree SHA-256 `34FE3F5C...ABD65B`.
- License disposition `45/45`, exception `0`.
- Durable manifest hiện tại cover `1.634` file / `20.332.100` bytes; tree SHA-256 `8A501CC65AB4910284488803973127BA7B74D74EAF75C91F4C916ECE0B2BB6AC`.

Exact lexical/direct denominator từ CGO binding là:

- `15` module identities.
- `33` literal module-local include targets.
- `1/33` đã có trong cả hai standard-vendor output: `github.com/tree-sitter/go-tree-sitter@v0.25.0/allocator.h`.
- `32/33` vắng trong cả hai standard-vendor output.
- Trong 32 target vắng đó, `31` target / `162.822.185` bytes tồn tại trong exact checksum-verified source và authorized module zip.
- `1` target vắng cả source lẫn zip: `github.com/tree-sitter/tree-sitter-go@v0.25.0/src/scanner.c`.

Direct-root diagnosis ban đầu tìm được `18` source roots / `157` files / `169.903.194` bytes, source-root digest `1C21748DC98C376DD6EEAD63F5F21EF225F7F0C5F5F7E5A66267E73F3A64CE72`; `0/157` có trong standard run 1, standard run 2 hoặc durable vendor hiện tại.

Bounded fixed-point remeasurement còn mở rộng root set hiện tại thành `20` roots / `163` files / `170.018.014` bytes vì scanner sources của PHP và TypeScript tham chiếu `common/scanner.h`. `20/163` là evidence snapshot để kiểm tra implementation, không phải denominator hardcode vĩnh viễn. Denominator bền vững phải được máy derive lại từ exact inputs ở mỗi generation.

### 4.2 Cơ chế compile failure

Verifier hiện tại PASS trước compile với `45` modules, `1.634` manifest-covered files, `0` reparse, `0` missing/extra/mismatch. Go version là exact `go1.26.3 windows/amd64` và lane `GOMODCACHE` không cung cấp module source fallback.

Sau đó cgo preprocesses preamble trong vendored Go files. Các binding như `bindings/go/binding.go` chứa `#include "../../src/parser.c"`; PHP dùng `../../php/src/parser.c`; TypeScript dùng `../../tsx/src/parser.c`; `go-tree-sitter` dùng `#cgo CFLAGS: -Iinclude -Isrc` rồi `#include <tree_sitter/api.h>`. Go package vendoring không truy vết các preprocessor target nằm ngoài package directory, nên compiler resolve path tại đúng module-relative location dưới `vendor/` và nhận `No such file or directory`.

Compile log ghi failure cho toàn bộ `15` affected module identities, sau verifier PASS. Vì vậy verifier hiện tại chứng minh integrity của một incomplete tree, không chứng minh compile closure.

### 4.3 Recursive và conditional facts

- PHP và TypeScript scanner sources cần `common/scanner.h`; các file đó có trong authorized module source/zip nhưng nằm ngoài 18 direct roots.
- TypeScript `common/scanner.h` cần `tree_sitter/parser.h`; binding có `#cgo CPPFLAGS: -I../../tsx/src` hoặc source-root tương ứng, và complete `tsx/src`/`typescript/src` chứa header này. Derivation bắt buộc parse cả `CPPFLAGS`, không chỉ `CFLAGS`.
- `go-tree-sitter/src/wasm_store.c` có include `stdlib-symbols.txt` bên trong `#ifdef TREE_SITTER_FEATURE_WASM`. Canonical build không define macro này; branch bất hoạt. Nếu macro xuất hiện trong effective build flags sau này, verifier/generator phải fail vì target không có authority.

## 5. Adjudication cho `tree-sitter-go/src/scanner.c`

Không synthesize scanner bytes và không đổi dependency version.

Exact preconditions đã xác minh:

- Module: `github.com/tree-sitter/tree-sitter-go@v0.25.0`.
- Authorized `bindings/go/binding.go`: `333` bytes; SHA-256 `680E50820F2F7429FEAC88B69DDCEE67735E6E692DFD1F98147109D77D296096`.
- Binding có một block duy nhất:

```c
#if __has_include("../../src/scanner.c")
#include "../../src/scanner.c"
#endif
```

- Authorized `src/grammar.json`: `198.042` bytes; SHA-256 `0DCC665AA521A1B73A200BE13B4E376780CCA037D764A165F12864BABFE28000`; `externals` có cardinality `0`.
- Authorized `src/parser.c`: `1.572.685` bytes; SHA-256 `3DBF6ED1238B5DFCF2BE4D2F2D4CB27A14D34F34D7784ECCCCBFD532FD4A6D85`; chứa exact `#define EXTERNAL_TOKEN_COUNT 0`.
- Authorized module zip: `232.860` bytes; raw SHA-256 `E4501D10DB1AEE5F0B15354F4AE31296D8D8A93282A85D12302E6C682944438A`; có binding và parser, không có `src/scanner.c`, `.cc` hoặc `.cpp`.
- Các build metadata khác của module chỉ add scanner khi file tồn tại.

Reviewed patch phải xóa đúng block ba dòng trên, không sửa byte nào khác. Exact postimage dự kiến là `245` bytes; SHA-256 `5A456F352CFEF15430F3FCFB155A3E1BDE38F244E5F8537026553B1F21547713`; delta `-88` bytes.

Patch chỉ được áp dụng nếu toàn bộ preconditions trên cùng đúng. Preimage mismatch, block count khác `1`, scanner xuất hiện, `externals != 0`, `EXTERNAL_TOKEN_COUNT != 0`, hoặc patch cần offset/fuzz đều phải fail closed. Patch application phải là exact byte transform, không dùng fuzzy patching.

Durable reviewed patch authority duy nhất:

`third_party/go-vendor/patches/tree-sitter-go-v0.25.0-remove-absent-scanner.patch`

Không được có patch thứ hai. Không được thay block bằng stub scanner. Patch này không thay grammar semantics; nó loại dormant reference mà chính module source, generated parser và grammar đều xác nhận không có external scanner contract.

## 6. Các phương án đã đánh giá

### 6.1 Standard baseline + supplemental complete-root fixed-point + một reviewed patch

Quyết định: chọn.

Ưu điểm:

- Giữ unchanged `go.mod`/`go.sum` và standard vendor làm baseline Go-native.
- Runtime/build tiếp tục `GOPROXY=off`, `GOSUMDB=off`, explicit `-mod=vendor`.
- Chỉ bổ sung source mà active cgo closure cần, từ exact checksum-verified module authority.
- Complete parent roots tránh hand-picking individual parser/header files; fixed-point tự mở rộng khi nested include thay đổi.
- Một patch minh bạch, guarded và reviewable giải quyết duy nhất reference không có source authority.
- Packaged `go-src` đã copy nguyên tree, nên không phát sinh một runtime/proxy path mới.

Chi phí:

- Vendor tree tăng khoảng denominator source hiện tại `170 MB`; final exact size phải đo sau materialization.
- Generator/verifier phải hiểu include roots, CPPFLAGS và conservative conditional guards.

### 6.2 Repository-owned immutable module/file proxy dùng ở build time

Quyết định: loại.

Phương án này có thể giữ zip bytes và để Go verify bằng `go.sum`, nhưng buộc canonical build đổi từ `GOPROXY=off` sang file proxy, rehydrate lane cache, mang toàn bộ proxy vào packaged source và thêm Windows file-URI/runtime plumbing. Nó mở một build-time acquisition path mới, phức tạp hơn correction đã chọn và không cần thiết.

### 6.3 Chỉ copy 31 direct target hoặc một danh sách file thủ công

Quyết định: loại.

Direct list đã bỏ sót recursive `common/scanner.h`; nó có thể âm thầm drift khi upstream binding/header thay đổi. Đây không phải fail-closed closure.

### 6.4 Copy toàn bộ 15 module roots

Quyết định: loại.

Đây là complete nhưng thêm test, examples, JS metadata, nested module files và surface không tham gia build. Fixed-point complete parent roots đạt cùng tính đóng với ít byte và ít hành vi phụ hơn.

### 6.5 Giữ nguyên conditional scanner include, không patch

Quyết định: không chọn.

Về C preprocessor, `__has_include` có thể làm reference này bất hoạt. Tuy nhiên closure inventory vẫn phải mang một special unresolved lexical exception. Exact reviewed patch với guard grammar/parser/zip mạnh hơn, làm final vendor không còn dormant reference và đơn giản hóa fail-closed verification.

### 6.6 Đổi `tree-sitter-go` version hoặc local replacement

Quyết định: loại.

Vi phạm unchanged `go.mod`/`go.sum`, cần authority/version review mới và chạm stop condition. Không có bằng chứng bắt buộc mở rộng như vậy.

### 6.7 Sinh scanner stub, dùng cache, precompiled archive hoặc network fallback

Quyết định: cấm.

Các lựa chọn này tạo bytes không có authority, làm mờ source provenance, phụ thuộc host/toolchain hoặc vi phạm offline/lane invariant.

## 7. Durable layout sau correction

```text
E:\Anvien\
├── go.mod                                      # byte-identical
├── go.sum                                      # byte-identical
├── vendor\
│   ├── modules.txt                             # standard baseline file
│   └── <standard files + supplemental roots + one patched binding>
├── third_party\go-vendor\
│   ├── manifest.v1.json                        # closureContractVersion = 2
│   ├── licenses\                               # unchanged complete 45/45 authority
│   └── patches\
│       └── tree-sitter-go-v0.25.0-remove-absent-scanner.patch
├── scripts\materialize-go-vendor.ps1
├── scripts\verify-go-vendor.ps1
└── scripts\test-go-vendor-contract.ps1
```

Giữ stable manifest path `third_party/go-vendor/manifest.v1.json` để không mở rộng package/runtime consumers. Manifest tiếp tục `schemaVersion = 1` nhưng bắt buộc thêm `closureContractVersion = 2` và `authorityKind = "go_vendor_cgo_closure"`. Updated verifier phải reject manifest cũ thiếu closure contract.

## 8. Machine-checkable generation contract

### 8.1 Source/checksum authority

- Exact Go: `go1.26.3 windows/amd64`.
- Exact unchanged `go.mod`/`go.sum` hashes ở Mục 3.
- Standard vendor module denominator: derive từ `vendor/modules.txt` và exact build/module list; hiện là `45`, không hardcode để che drift.
- Mỗi source module phải có content sum và `/go.mod` sum khớp unchanged `go.sum`; module zip phải pass Go-native checksum verification trước extraction.
- Zip là byte authority cho supplemental files. Extracted GOMODCACHE tree chỉ là derived view; mỗi copied file phải map tới exact normalized zip entry và có cùng bytes/hash.
- Không đọc shared/global/protected/quarantined cache. Current clean E-only source root được phép dùng cho correction generation; nó không trở thành runtime authority.

### 8.2 Standard baseline

1. Tạo hai clean unique E-only generation lanes.
2. Chạy exact `go mod vendor -o <lane-output>` từ identical inputs.
3. Normalize path bằng `/`; reject absolute path, drive/UNC, `.`/`..`, escape, backslash trong manifest, NUL, reparse/symlink và case-insensitive collision.
4. Inventory mọi baseline file theo `path`, `bytes`, uppercase SHA-256.
5. Sort canonical bằng ordinal path; reject exact hoặc ordinal-ignore-case duplicate trước sort.
6. Tree digest là SHA-256 của UTF-8/no-BOM rows `path<TAB>bytes<TAB>SHA256<LF>`.
7. Run 1 và run 2 phải equal về module rows, file rows, bytes, hashes và tree digest.

Manifest section `standardBaseline` phải record command, Go identity, module count/list digest, run counts, per-run file/byte/tree denominator, difference count `0` và ordered file rows.

### 8.3 Supplemental root derivation

Derivation chạy độc lập trên cả hai baseline và phải cho cùng output:

1. Map mỗi vendored package về exact module identity bằng `modules.txt` và source module rows.
2. Machine-scan vendored Go files có `import "C"`; parse literal `#include` và literal relative `-I` flags từ cả `CFLAGS` lẫn `CPPFLAGS` theo cgo preamble.
3. Resolve quoted include trước theo directory chứa source/preamble, sau đó theo effective relative include roots; resolve angle include theo include roots. Mọi target phải normalize trong module root.
4. Nếu active target đã có trong baseline, không supplement.
5. Nếu active target thiếu baseline nhưng có trong authenticated zip, add parent directory của target làm complete supplemental root và copy toàn bộ directory recursively, byte-for-byte, vào exact module-relative vendor location.
6. Scan C/C++/header/include files trong các root mới; lặp fixed-point cho mọi active module-local literal include mới. Không dừng ở direct binding list.
7. Unsupported dynamic/macro include có thể trỏ module-local, path escape, unresolved active target hoặc ambiguous include-root resolution phải fail closed; không đoán.
8. Conservative conditional evaluator chỉ chấp nhận simple `#if 0`, `#ifdef`, `#ifndef`, `defined(...)`, phủ định và boolean combination trên exact effective macro set. Unsupported condition chứa missing module-local target phải fail.
9. `TREE_SITTER_FEATURE_WASM` phải absent trong canonical effective defines. Reference `stdlib-symbols.txt` được record dưới `inactiveConditionalReferences`; nếu macro được define, generation/verifier fail trước Go.
10. Sau fixed-point, ordered module/root/file inventory của hai runs phải equal `2/2`.

Manifest section `supplementalClosure` phải record algorithm version, affected module rows, root rows, active seed target rows, recursive edge rows, inactive conditional rows, file rows, counts/bytes/tree digest và run equality. Current `15 modules / 20 roots / 163 files` chỉ là expected remeasurement assertion cho lần correction đầu; manifest values phải đến từ derivation, không từ source hardcode.

### 8.4 Reviewed patch

Patch application order là: standard baseline -> supplemental overlay -> reviewed patch.

Manifest section `reviewedPatches` phải có cardinality đúng `1` và record:

- patch authority path và patch SHA-256;
- exact module path/version;
- destination `vendor/github.com/tree-sitter/tree-sitter-go/bindings/go/binding.go`;
- preimage bytes/hash `333` / `680E50820F2F7429FEAC88B69DDCEE67735E6E692DFD1F98147109D77D296096`;
- exact removed block and occurrence count `1`;
- grammar bytes/hash, `externals = 0`;
- parser bytes/hash, `EXTERNAL_TOKEN_COUNT = 0`;
- source/zip scanner absence proof;
- postimage bytes/hash `245` / `5A456F352CFEF15430F3FCFB155A3E1BDE38F244E5F8537026553B1F21547713`;
- additions `0`, deletions `3`, byte delta `-88`, fuzz/offset `0`.

Patch file, script transform và final postimage phải agree. Missing/extra patch, modified patch, precondition drift hoặc second application đều fail.

### 8.5 Final tree, licenses và publication

- Manifest `files` covers every final `vendor/**`, `third_party/go-vendor/licenses/**` và `third_party/go-vendor/patches/**` file; manifest tự loại chính nó khỏi tree digest.
- Mỗi final vendor row có origin class đúng một trong `standard`, `supplemental`, `patched_postimage`.
- `standard` rows equal baseline rows; supplemental rows equal authenticated zip-derived rows; patched path links exact standard preimage, patch và postimage.
- Extra, missing, duplicate-origin hoặc unclassified row fail.
- License module denominator giữ `45/45`; supplemental files không tạo module identity mới. Missing/unknown license disposition fail.
- Hai complete final candidates phải equal `2/2` theo paths, bytes, hashes và final tree digest.
- Chỉ publish atomically sau candidate verifier PASS. Partial candidate không được chạm durable roots.
- Timestamps, absolute lane paths và machine-specific values không tham gia deterministic manifest bytes.

## 9. Runtime verifier contract

`scripts/verify-go-vendor.ps1` phải chạy bằng cả `pwsh` và Windows PowerShell 5.1, không gọi Go, network, VCS hoặc cache acquisition.

Verifier phải:

1. Xác nhận exact source root, `go.mod`/`go.sum`, manifest contract version, materializer/verifier/patch identities và allowed authority layout.
2. Reject mọi reparse point/symlink, path escape, malformed/case-colliding/duplicate path.
3. Recompute actual complete inventory và final tree digest.
4. Reconstruct baseline/supplemental/patched partitions từ manifest và enforce one-origin-per-file.
5. Enforce standard baseline equality records `2/2`, supplemental derivation equality `2/2`, final equality `2/2` và patch cardinality `1`.
6. Enforce current actual standard rows equal recorded baseline except exact patched destination; exact patch postimage phải match.
7. Enforce supplemental root/file rows, module identity, zip content/go.mod checksum provenance và all recursive edges recorded at generation.
8. Enforce no active unresolved target. Inactive `TREE_SITTER_FEATURE_WASM` row chỉ hợp lệ khi macro không có trong sealed build controls.
9. Enforce complete license disposition và no unreferenced license/patch file.
10. Trả visible nonzero exit trên missing, extra, tamper, drift, stale legacy manifest hoặc unsupported condition; không repair và không fallback.

Verifier phải chạy trước first `go run`, `go build`, `go test` hoặc packaged-runtime acceptance. Verifier PASS không tự đủ để tuyên bố compile closure; exact offline compile là gate riêng bắt buộc.

## 10. Invariant phải giữ nguyên

- `go.mod` và `go.sum` byte-identical.
- Checked-in `vendor/**` + `third_party/go-vendor/**` là runtime/build source and integrity authority.
- Canonical/package/launcher root Go commands dùng explicit `-mod=vendor`.
- `GOPROXY=off`, `GOSUMDB=off`, `GOENV=off`, `GOWORK=off`, `GOTOOLCHAIN=local`, `GOVCS=*:off`; inherited Go proxy/private/insecure controls được sanitize.
- `GOMODCACHE`, `GOCACHE`, GOPATH, HOME và temp là isolated E-only derived lane state; không authority/fallback.
- Verifier trước first Go; failure không dùng stale runtime.
- License disposition complete `45/45`.
- Deterministic baseline, overlay và final generation `2/2`.
- `prepare-go-source` mang exact vendor, manifest, licenses, patches và verifier; packaged source dùng cùng contract, không dựa repo parent.
- Current runtime/native/built-reader gates và target/oracle/boundary gates giữ nguyên; correction không mở P6-E hay target sớm.
- Publication vẫn staged/atomic; canonical destination chỉ đổi sau complete success.

## 11. Exact implementation handoff

### 11.1 Planner gate

Bốn Child 06 ledgers bắt buộc được Main refresh đúng một lần bằng planner skill trước khi code resume. Lý do: standard-vendor-only assumption đã sai, correction thêm supplemental derivation, patch authority/layout, manifest/verifier fields và negative gates. Đây là Patch-mode owner-map/acceptance update bên trong P6-D/R42, không tạo phase, slice, campaign hay commit boundary mới.

Planner refresh phải giữ:

- P6-D unchecked `REJECT/BLOCKED`;
- P6-E exactly one U1-U4 slice `LOCKED`;
- Coder hiện hữu `01a02f10-c175-7162-ac32-6df30a4d604a`; không thay thế;
- code-first, tests sau production behavior;
- full build và mọi existing gate/commit boundary không đổi.

### 11.2 Correction write scope cho Coder hiện hữu

Được thêm/sửa đúng correction delta:

- `scripts/materialize-go-vendor.ps1`;
- `scripts/verify-go-vendor.ps1`;
- `scripts/test-go-vendor-contract.ps1`;
- `vendor/**` dưới deterministic regeneration/overlay contract;
- `third_party/go-vendor/manifest.v1.json`;
- `third_party/go-vendor/patches/tree-sitter-go-v0.25.0-remove-absent-scanner.patch`;
- `third_party/go-vendor/licenses/**` chỉ được regenerate nếu output byte-identical và denominator vẫn `45/45`.

Được hoàn tất trong original P6-D owner scope, không redesign:

- `scripts/full-build.ps1`;
- `internal/cli/package_runtime.go`;
- `anvien-launcher/build.ps1`.

Chỉ sau production materializer/verifier/authority và bounded offline compile đúng:

- `internal/cli/package_command_test.go`;
- `internal/cli/package_runtime_p6d_test.go`;
- reusable negative contract test nêu trên.

Preserve-only:

- `go.mod`, `go.sum`;
- `scripts/ensure-ladybug-native.ps1` và durable Ladybug 3/3 bundle;
- P6-E owners, roadmap ngoài exact Planner refresh, target trước gate, Owner deletion và mọi unrelated/protected path.

Coder không có ledger, Supervisor, target, stage/commit/push hoặc acceptance authority.

## 12. Negative contract gates bắt buộc

- Drift `go.mod` hoặc `go.sum`.
- Wrong Go version, source zip content sum hoặc `/go.mod` sum mismatch.
- Standard baseline run mismatch, missing/extra/tampered baseline file.
- Supplemental root/file missing, extra, tampered, wrong destination/module, order/digest mismatch.
- Recursive `common/scanner.h` hoặc TypeScript `tree_sitter/parser.h` missing.
- Deriver bỏ qua `CPPFLAGS`, không đạt fixed-point, gặp dynamic include, path escape, case collision hoặc reparse.
- Define `TREE_SITTER_FEATURE_WASM` khi `stdlib-symbols.txt` không có authority.
- Patch absent, extra, second patch, patch tamper, preimage mismatch, occurrence != 1, fuzz/offset, wrong postimage.
- Scanner xuất hiện trong source/zip; grammar `externals` khác `0`; parser `EXTERNAL_TOKEN_COUNT` khác `0`; binding identity drift.
- Manifest legacy thiếu `closureContractVersion = 2`, origin overlap/unclassified file hoặc self-inclusion.
- Missing/unknown license disposition hay unreferenced license/patch file.
- Root verifier PASS nhưng packaged `go-src` missing/tampered vendor, manifest, patch hoặc verifier.
- Invalid authority phải fail trước first Go process; stale packaged/canonical runtime không được chấp nhận.

## 13. Validation và acceptance gates sau implementation

Thứ tự bắt buộc:

1. Main Planner refresh bốn ledgers; trả đúng delta cho Coder hiện hữu.
2. Coder làm production authority/generator/verifier trước; tests sau.
3. Hai clean materializations chứng minh equality riêng cho standard baseline, supplemental roots, one patch và final tree.
4. Verifier PASS trên `pwsh` và Windows PowerShell 5.1 ở repo root và packaged `go-src`.
5. Lặp exact bounded offline compile đã fail trước đó với empty lane GOMODCACHE, `GOPROXY=off`, no network/cache fallback. Compile phải PASS; record command, environment, acquired GOMODCACHE source file count và final tree identity.
6. Record final vendor bytes/file count, packaged `go-src` bytes/file count và delta so với standard baseline như benchmarkable package/source-size evidence.
7. Trước full build, Rule 13 terminate toàn bộ build-related holder/lock process; sau đó chạy full canonical build.
8. Prove verifier-before-first-Go, packaged source parity, packaged-runtime manifest identity và no stale-runtime publication/fallback.
9. Prove current runtime, Ladybug native-through-runtime, packaged/native và mọi affected built reader.
10. Chỉ sau current-runtime gates mới capture target pre-state và chạy existing target/oracle/boundary gates.
11. Chạy affected regression, fresh `anvien analyze --force`, fresh `detect-changes`, exact disposable artifact disposition và independent Supervisor.
12. Main chỉ stage/commit isolated P6-D sau Supervisor PASS và exact authority; correction report này không cấp Git authority.

## 14. Blast radius và rủi ro

Sealed pre-edit evidence hiện hữu cho `internal/cli/package_runtime.go` là `CRITICAL`; `buildGoRuntimePackage`, `prepareGoSourcePackage` và `ensurePackagedRuntime` nằm trên nhiều package/runtime flows. PowerShell `LOW 0` là extractor limitation, không phải no-consumer proof. Correction supply-chain chạm source bytes của `15` dependency modules và mọi root CLI compile, nên blast radius vận hành là HIGH/CRITICAL dù patch semantics duy nhất rất hẹp.

Rủi ro còn phải kiểm soát:

- Fixed-point scanner thiếu `CPPFLAGS` hoặc xử lý preprocessor quá rộng/hẹp: fail closed trên unsupported/ambiguous construct, có negative tests.
- Repository/package size tăng lớn: đo exact sau generation; không tối ưu bằng file list thủ công.
- Conditional WASM về sau được bật: current authority phải fail và cần architecture refresh; không tự tạo `stdlib-symbols.txt`.
- Upstream module version/content drift: one patch guard fail, không fuzzy apply.
- Windows path/case và PowerShell 5.1 behavior: enforce normalized inventory, collision rejection và dual-shell tests.

## 15. Sự thật và suy luận

Đã xác minh trực tiếp: identities, denominators, standard output omission, compiler failures, recursive common headers, TypeScript CPPFLAGS, inactive WASM guard, grammar/parser/scanner facts và exact patch pre/postimage.

Suy luận kiến trúc: standard baseline + fixed-point complete roots + guarded patch sẽ đóng exact offline compile mà không đổi module authority. Suy luận này hợp lý và implementable, nhưng chỉ exact offline compile PASS mới chuyển nó thành runtime evidence. Báo cáo này không tự accept implementation.

Residual architecture question: `0`.

Residual Owner choice/work: `0`.

## 16. Handoff niêm phong cho Main

Main phải:

1. Xác minh báo cáo này và identity seal.
2. Dùng planner một lần để refresh đúng bốn Child 06 ledgers theo Mục 11.
3. Trả duy nhất approved correction delta cho Coder hiện hữu `01a02f10-c175-7162-ac32-6df30a4d604a`.
4. Không mở Coder thay thế, P6-E, target, network, shared cache hoặc Git authority.

Trạng thái cuối của lane kiến trúc: `READY_FOR_MAIN`.

