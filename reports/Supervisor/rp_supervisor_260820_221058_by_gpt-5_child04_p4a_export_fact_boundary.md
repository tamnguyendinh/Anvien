# Supervisor Report: Child 04 P4-A export fact boundary

Verdict: PASS

## Metadata

- Report file: `reports/Supervisor/rp_supervisor_260820_221058_by_gpt-5_child04_p4a_export_fact_boundary.md`
- Review time: `2026-08-20 22:10:58 +07:00`
- Reviewer: `gpt-5`
- Repo/project: `E:\Anvien`
- HEAD được review: `ff2467bb92f94a9c53c4de030685686700051a98`
- Scope reviewed: Child 04 `P4-A` trên sáu candidate paths, hai report provenance được khai báo ban đầu, một Main rotation report được attribution trong final artifact gate, preserve-only siblings, build và fresh boundary validation.
- Claim reviewed: P4-A thiết lập smallest production contract cho one source-site `ExportFact`, export kind/names/ranges/canonical meaning/type-only/source provenance, structured countable diagnostic, deterministic deep-copy normalization/JSON, giữ nguyên access visibility và không mở AST extraction, compatibility projection/write, terminal resolution, barrel hoặc public-API state.
- Authority used: latest delegation; `AGENTS.md`; `working-rules`; `supervisor`; `backend-development` cùng `backend-code-quality.md`, `backend-testing.md`, `backend-architecture.md`; Child 04 roadmap và bốn living ledgers; `docs/contracts/graph-accuracy-contract.md`; current source/diff/runtime evidence.
- Related artifacts: `reports/Investigation/rp_main_260820_211734_orchestration_rotation_handoff.md`; `reports/Investigation/rp_main_260820_221051_orchestration_rotation_handoff.md`; `reports/coder/rp_coder_260820_214739_by_gpt-5_child04_p4a_export_fact_boundary.md`.

## Executive Summary

- Vấn đề nghiệm thu: xác định current worktree có đóng đúng P4-A first-class ScopeIR contract mà không lấn sang P4-B/P4-C/Child 05 hay không.
- Quyết định: toàn bộ invariant P4-A được chứng minh độc lập bằng exact source/diff, fresh excluded Anvien graph, source-level total-order/deep-copy audit, byte-exact golden, canonical full build và fresh nearest-boundary tests.
- Boundary: đúng bốn production files và hai test-after-code files; staged paths bằng `0`; hai initial report paths và Main-owned rotation report mới đều được attribution; preserve-only worktree blobs bằng đúng HEAD blobs.
- Kết quả cần tiếp theo: Main tiêu thụ report này như `E4-P4A-REVIEW1`, tự refresh living ledgers, chạy excluded graph/change detection, stage/commit exact accepted slice và không push. P4-B chỉ được mở sau commit P4-A thành công.

## Claim / Authority Reconstruction

P4-A chỉ chịu trách nhiệm tạo contract in-memory có thể lưu một fact cho từng binding/specifier site. Slice này phải biểu diễn source form, names, ranges, meanings, `TypeOnly`, source provenance và diagnostic; phải clone/normalize/serialize deterministically; không được tự extract AST, ghi compatibility state, project graph/persistence hoặc chứa derived terminal/public-API state.

Acceptance authority trực tiếp là P4-A trong Child 04 plan: editable `facts.go`, `kinds.go`, `ir.go`, `sort_keys.go`; test-after-code `scopeir_test.go` và golden; inspect/preserve-only TSJS extraction, graph/persistence consumers, `DefinitionFact.Visibility`, và Child 05 resolution. `CRITICAL/HIGH` Anvien impact là blast-radius warning, không phải edit prohibition.

## Source-Level Clearance Notes

- `internal/scopeir/facts.go`: clear — `ExportProvenance` tại line `114-118`; one-site `ExportFact` tại `121-144`; structured `ExportDiagnosticCode`/`ExportDiagnosticFact` tại `146-163`. Nested mutable values chỉ gồm `Meanings`, `SelectionRange`, `TargetRaw`, và đều có owning clone path. `LocalDefID`, `TargetRaw`, `TargetExportedName` được giới hạn rõ là local/source-side, không phải terminal target. `DefinitionFact.Visibility` tại line `19` không đổi.
- `internal/scopeir/kinds.go`: clear — `ExportKind` tại `78-99` giữ bảy source forms (`direct`, `named`, `default`, `alias`, `reexport`, `star`, `namespace`); `ExportMeaning` tại `101-109` giữ ba lanes (`value`, `type`, `namespace`) và cho phép dual meaning bằng set.
- `internal/scopeir/ir.go`: clear — collections tại `23-24`; outer copies tại `41-45` và `154-158`; meaning sort/dedup tại `99-112`; export/diagnostic ordering tại `125-130`; nested clones tại `173-179`; deterministic marshal/unmarshal-normalize tại `220-236`. `ExportDiagnosticFact` không có nested reference nên outer slice clone là đầy đủ.
- `internal/scopeir/sort_keys.go`: clear — `compareExport` tại `58-102` bao phủ mọi serialized `ExportFact` field; `TargetRaw=nil` khác non-nil, kể cả pointer-to-empty, tại `80-89`; diagnostic comparator tại `105-124` bao phủ mọi diagnostic field; canonical meaning sequence tại `127-143`; provenance comparator tại `145-150`. Comparator equality chỉ xảy ra khi serialized record tương đương.

Test/golden clearance:

- `internal/scopeir/scopeir_test.go:213-289` chứng minh canonical meaning set, nested deep copy cho `Normalized` và `NormalizeOwned`, diagnostic retention và `Visibility="private"` không đổi.
- `internal/scopeir/scopeir_test.go:291-378` chứng minh permutation-stable JSON, nil-vs-empty `TargetRaw`, diagnostic count/order, source non-mutation, round trip và negative Child 05 field assertions.
- `internal/scopeir/scopeir_test.go:380-421` chứng minh bảy source-form kinds, three meaning lanes, dual value/type và explicit type-only state được bảo toàn ở contract boundary.
- `internal/scopeir/testdata/scopeir.golden.json:101-187` là representative byte oracle cho direct/re-export facts và structured unsupported diagnostic. Seven-kind/type-only/malformed coverage nằm ở dedicated tests, không bị overclaim thành golden coverage.

## Evidence Checked

### Đạt

- Coder report identity: `16,864` bytes / `263` LF / SHA-256 `E4627C2426C1EF56AE7FA6A36FD1104041B9512B8C06A7B92D4FCBCE123F4EB6`; khớp delegation. Main report SHA-256 `3A4D74472EB90D329F6CC61C1656BB80A105ADA277794B50E9B9595679C33BC1`.
- Final artifact gate phát hiện `reports/Investigation/rp_main_260820_221051_orchestration_rotation_handoff.md` xuất hiện concurrent. Full read xác nhận đây là Main-owned 60-minute rotation handoff cho đúng active Supervisor thread `01a01fa8-e5b6-7e90-89d4-95f86226c760`; file tự khai báo exact six-file candidate, hai prior reports và forthcoming Supervisor report. Identity: `8,522` bytes / `121` LF / SHA-256 `26E414A9EE0EC3A455A006DA65AA362651643A6AD62DAC6A58B9CEE8C42C17AC`. Drift được attribution, không thay đổi candidate bytes và không còn unknown path.
- Fresh graph command trước mọi graph-based review:

  ```text
  anvien analyze . --force --exclude "internal/aicontext/skills/**" --exclude ".claude/skills/**"
  ```

  Exit `0`; scanned/parsed/failed `1,128/626/0`; graph `81,103/120,485`. Hai file tăng so với Coder graph là đúng hai declared report files; parsed-code count và candidate hashes không đổi. Full-build final unexcluded analyze counts không được dùng làm Child 04 semantic evidence.
- Fresh `file-detail` final-byte summaries: `facts.go` `192` symbols / `1127` inbound / `27` outbound / `174` local / `98` linked tests; `kinds.go` `99/619/0/10/98`; `ir.go` `57/1106/38/89/96`; `sort_keys.go` `122/252/128/54/95`; graph stale `false`, changedSinceAnalyze `false`.
- Fresh upstream file impact warnings: `facts.go` CRITICAL `347/110/76/1`; `kinds.go` CRITICAL `1358/143/182/1`; `ir.go` CRITICAL `124/55/30/1`; `sort_keys.go` HIGH `27/24/3/1` impacted/direct/files/flows. Scope được giữ đúng bốn production owners.
- Exact diff: sáu modified files, `570` insertions / `0` deletions; `git diff --check -- internal/scopeir` exit `0`; staged diff bằng `0`.
- Candidate bytes khớp Coder report:

  | Path | Bytes / LF | SHA-256 | Git blob |
  |---|---:|---|---|
  | `internal/scopeir/facts.go` | `11,518 / 272` | `7F2B51D878F1541995AA884C438E18F1B1E6C72E20597E2C171FB36A59BFCA6A` | `472c84c171f5ded85315ed3923f307d82334feba` |
  | `internal/scopeir/kinds.go` | `5,167 / 153` | `2D80ECABB63D07BAFDED161E000BFC1A412BC5B8CB242B29726271688E5D2878` | `91ed38e36ee257b16014705021ce72fa2a568c68` |
  | `internal/scopeir/ir.go` | `9,259 / 237` | `732EE7F8959F077FED5550962A5369A020F6C9EC5ABDAF4540054C00E46E728C` | `d062b6d3a68ad92628165a2122a4e4c44835450c` |
  | `internal/scopeir/sort_keys.go` | `12,528 / 446` | `5C155B4C151D8E11833015376C26979C50928425A169CABAB475A65F52A52DB5` | `db16c7565490259fd00a5893fd92f4b8359063e5` |
  | `internal/scopeir/scopeir_test.go` | `27,670 / 788` | `ADC375EBB590F28FF72347A9B16EF0A09B971941D8C22F3F0789E877E58E6D04` | `ea4cde217e8ddb98aa3b23ad13c542284e087f5b` |
  | `internal/scopeir/testdata/scopeir.golden.json` | `6,740 / 307` | `C92682219EB9BD6F518B28913C64276A2E6C44B1162215DB959EF3CD5254B306` | `1163a34804efad92710bfc2ea487151424c6bde8` |
- Preserve-only blob equality with HEAD and diff exit `0`: `internal/providers/tsjs/{definitions.go,imports.go,extract.go}`, `internal/resolution/{emit.go,indexes.go,import_resolution.go}`, `internal/lbugload/csv.go`, `internal/lbugschema/schema.go`. Entire TSJS extraction/resolution/projection/persistence surfaces have no candidate drift.
- New-type usage scan over Go source finds `ExportFact`, `ExportDiagnosticFact`, `ExportKind`, `ExportMeaning`, `Exports`, `ExportDiagnostics` only in the four ScopeIR owners and `scopeir_test.go`; không có AST producer, graph projector, persistence writer hoặc resolver consumer.
- Pre-build holder gate: `anvien doctor locks --repo E:\Anvien --json` trả `status=free`, `exists=false`, `alive=false`; ba global editor-owned `anvien.exe ... mcp` holders và ba parent `cmd.exe` được dừng; independent recount `buildRelatedHolderCount=0`.
- Canonical build:

  ```text
  npm run full-build
  ```

  Exit `0`; packaged CLI `1.2.8`; Web build `2,943` modules; Vite `23.05s`. npm allow-scripts, mixed static/dynamic import và large chunk là non-failing warnings. Post-build lock free và holder count `0`.
- Fresh focused matrix:

  ```text
  go test -v ./internal/scopeir -run '^(TestMarshalDeterministicMatchesGolden|TestUnmarshalNormalizesScopeIR|TestNormalizeInPlaceMatchesNormalized|TestNormalizeOwnedMatchesNormalized|TestExportCollectionsAreDeepCopiedByNormalization|TestExportCollectionsMarshalDeterministically|TestExportKindsAndMeaningsRepresentSourceForms)$' -count=1
  ```

  Exit `0`; đúng `7/7` named tests.
- Fresh package boundary: `go test ./internal/scopeir -count=1` exit `0`.
- Fresh nearest product boundary:

  ```text
  go test ./internal/scopeir ./internal/providers ./internal/providers/tsjs ./internal/resolution ./internal/analyze ./cmd/binding-contract-probe -count=1
  ```

  Exit `0` cho cả sáu packages.

### Failed, retained accurately, không thuộc P4-A

```text
go test ./... -count=1
```

Exit `1`; không được gọi là successful evidence. Năm compile/setup-negative fixture packages là `anvien/test/fixtures/lang-resolution/go-map-range`, `go-method-enrichment`, `sample-code`, `go-make-builtin`, `go-type-assertion`. Hai out-of-slice parity failures là `internal/providers/csharp::TestResolveCSharpGraphParityCounts` (expected `ACCESSES:2`, got none) và `internal/providers/dart::TestResolveDartGraphParityCounts` (expected `ACCESSES:2`, got `1`). Fresh focused/nearest commands ở trên là closure evidence; baseline failures này không do sáu P4-A paths và không che khuất same-slice failure.

### Not run

- `anvien detect-changes`: không chạy theo explicit Supervisor authority; `E4-P4A-DETECT1` vẫn Main-owned và pending.
- Stage/commit/push: không thực hiện; `E4-P4A-COMMIT1` vẫn Main-owned và pending.
- P4-B, P4-B1, P4-C, P4-C2, Child 05: không mở.
- `E:\cheapapp.org`: không access; real-target validation thuộc P4-C2.
- Browser/Playwright: N/A cho non-UI in-memory contract boundary.
- Không dùng `internal/aicontext/skills/**` hoặc `.claude/skills/**` làm evidence. Unexcluded analyze nằm bên trong canonical full-build chỉ được tính vào exit của build script, không được dùng cho semantic graph claim.

## Invariant Closure

| Invariant | Closure |
|---|---|
| One source-site fact contract | `ExportFact` là one-site value record với site range + statement provenance; collection giữ nguyên từng record và không dedup facts. Dedicated `7/7` test chứng minh representability/conservation, không được diễn đạt thành AST extraction proof. |
| Canonical meaning set | `NormalizeInPlace` sort+dedup; owned paths clone trước mutation; dual value/type và namespace/type-only state round-trip đúng. |
| Nil-vs-empty `TargetRaw` | Pointer shape + comparator + JSON `omitempty` giữ nil khác pointer-to-empty; permutation bytes bằng nhau. |
| Nested deep copy | Mọi nested mutable export member (`Meanings`, `SelectionRange`, `TargetRaw`) được clone trong cả `Normalized` và `NormalizeOwned`; source mutation assertions thành công. |
| Diagnostic count/order | Typed code/range/node/reason/provenance; collection chỉ sort, không filter/dedup; count bảo toàn và comparator bao phủ mọi field. |
| Deterministic JSON/unmarshal | Total ordering không bỏ serialized field; fixed struct JSON, normalized marshal, normalized unmarshal, byte-exact golden và remarshal equality. |
| Visibility independence | `DefinitionFact.Visibility` diff bằng `0`; private fixture vẫn private; không có `Visibility="exported"` hay `isExported` writer. |
| Zero later-slice state | Production shape và usage scan không có `TargetFile`, terminal `TargetDefID`, `TargetModuleScope`, `LinkStatus`, `TransitiveVia`, ambiguity/cycle/barrel/public-API state; preserve-only owners bằng HEAD. |

Affected invariant: first-class TypeScript export syntax fact/meaning/diagnostic/deterministic ScopeIR boundary.

Sibling surfaces checked: TSJS extraction owners; graph projection; Ladybug persistence; current import-resolution owners; existing normalization consumer; `DefinitionFact.Visibility`; exact test/golden; worktree provenance.

Residual unverified same-invariant surfaces: none trong P4-A. AST extraction (P4-B/P4-B1), graph/compatibility/persistence projection (P4-C), target `21/21` (P4-C2), terminal resolution/barrels/ambiguity/cycles/public API (Child 05) là locked successor work, không phải residual P4-A.

## Coder Evidence Calibration

- Coder report đúng khi tự giới hạn là evidence-only và giữ `E4-P4A-REVIEW1`, `E4-P4A-DETECT1`, `E4-P4A-COMMIT1` pending.
- General normalization tests tại line `66`/`77` không tự chứa export fixtures; export coverage thực nằm ở dedicated tests `213-421` và golden round trip. Đây là wording correction, không phải source/test blocker.
- `7/7` one-fact matrix chứng minh contract representability/conservation, không chứng minh provider extraction hoặc multi-specifier AST behavior; các hành vi đó vẫn khóa đúng ở P4-B/P4-B1.
- Forbidden-field loop trong test không exhaustive theo mọi tên có thể tưởng tượng; exact production shape, all-field comparator audit, Go usage scan và preserve-only blob equality đóng phần còn lại.
- Coder Anvien counts là pre-final implementation evidence pointer; fresh Supervisor excluded graph trên exact final hashes nêu ở trên là authority cho review này.

## Overall Evaluation

Implementation quality: contract nhỏ, đúng owner, không tạo second export truth, clone đủ nested mutable state và total-order đủ serialized fields.

Evidence quality: canonical build và exact boundary commands được rerun fresh; Coder wording overclaims đã được thu hẹp; repo-wide exit `1` được giữ nguyên là baseline failure ngoài slice.

Authority/boundary: exact six-file candidate + three attributed Main/Coder provenance reports, no staged/unknown drift, no later slice opened, no target access.

Report này cung cấp independent Supervisor clearance cho `E4-P4A-REVIEW1`; living ledgers chưa được sửa vì Supervisor chỉ được tạo đúng report này.

## Handoff cho Main

1. Xác minh report hash/path và tiêu thụ report này làm `E4-P4A-REVIEW1`.
2. Dùng planner đúng một lần để refresh roadmap + bốn Child 04 living ledgers theo current evidence.
3. Trước graph-based closure, chạy lại excluded analyze; sau đó Main chạy required `anvien detect-changes`, kiểm tra exact manifest, stage và commit isolated P4-A.
4. Không push. Không mở P4-B trước khi P4-A commit thành công và worktree provenance được xác nhận.
5. Không cần nhắn Coder từ Supervisor lane; nếu Main phát hiện post-report drift thì dừng và tái route theo ownership.
