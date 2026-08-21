# Child 05 / P5-C Export Resolution — Coder Handoff

## Trạng thái

- `READY_FOR_SUPERVISOR`.
- Slice: P5-C — resolve alias/re-export/star/cycle/ambiguity/meaning và chặn repository-global rescue sau explicit-import export failure.
- Checkout duy nhất: `E:\Anvien`.
- Authority HEAD: `861000cb6b6e36ce105623f0dc8c093b089f61fa` trên `master`.
- P5-B predecessor: `c1559df953a277b099009f8489576d00ed25aa58`.
- Inventory authority: `reports/coder/rp_coder_260821_202905_by_gpt-5_child05_p5c_pre_implementation_inventory.md`, `17,036` bytes / `234` LF / `0` CR / UTF-8 không BOM / SHA-256 `7E77E5010CC3D728753547D0EE5499C51ED38CFD9660130381274BBB3DDCA256`.
- Không sửa ledger, không chạy `detect-changes`, không commit, không gọi Supervisor, không truy cập target/`E:\cheapapp.org`, không dùng C-worktree hoặc lane/subagent khác.

## Invariant Family Map

- Family: repository-backed TypeScript module export lookup và proof-bearing terminal preparation trong P5-C.
- Authority / SSOT: accepted Child 04 `ScopeIR.ExportFact`; accepted P5-A `ImportFact.RequestedMeanings`/`TypeOnly`; accepted P5-B `exportTables` explicit entries + star adjacency.
- Primary path: source-written semantic import -> resolve toàn bộ module/file candidates -> build P5-B tables đúng một lần -> deterministic export traversal -> terminal definition hoặc explicit unresolved outcome -> import binding.
- Sibling surfaces đã đóng: direct/named/alias imports; explicit-over-star; star default exclusion; namespace/member; same-terminal dedupe; distinct-terminal ambiguity; terminal/pure cycles; value/type/namespace lanes; explicit-import call fallback; syntactic `IMPORTS`; non-semantic Go package/member path.
- Preserve-only surfaces: Child 04 facts/providers; P5-B table shape/construction/tests; `resolveImportFiles`, `resolveImportFile`, `import_resolution.go`, `resolvedImport.TargetFiles`; generic `resolveName`, `resolveGlobalName`, `resolveGlobalCallName`; `emitImportEdges`; graph/persistence/readers; P5-D; target.
- Forbidden fallback: physical definition không được tự trở thành export; không first-candidate selection; không star-derived default; không name-only meaning repair; không repository-global same-name rescue khi explicit import đã claim identifier nhưng export lookup thất bại.
- Stale tests/helpers/plans: không có artifact nào trong authorized P5-C owner set còn encode physical-definition traversal hoặc global rescue như acceptance; existing legacy tests được giữ làm regression, focused P5-C tests là authority mới cho slice.
- Residual unverified surfaces trong P5-C: `none`. Supervisor/detect/commit là các gate Main-owned còn pending, không phải behavior gap. P5-D/target vẫn là slice sau và không được mở bởi report này.

## Blast radius đã chấp nhận

`E5-P5C-IMPACT1` được dùng nguyên trạng; không rerun analyze/file-detail/impact vì không có invalidation của inventory owner set.

| Exact symbol | Risk | Impacted symbols | Files | Modules | Processes |
|---|---|---:|---:|---:|---:|
| `workspace.resolveImportedDef` | HIGH | 19 | 12 | 1 | 3 |
| `workspace.resolveImports` | CRITICAL | 28 | 16 | 3 | 17 |
| `workspace.resolveImportedMember` | CRITICAL | 18 | 8 | 3 | 34 |
| `buildWorkspace` | CRITICAL | 49 | 22 | 8 | 23 |
| `resolveCall` | CRITICAL | 27 | 11 | 7 | 32 |
| `workspace.resolveGlobalCallName` | CRITICAL / preserve-only | 6 | 4 | 2 | 23 |

Additional preserve-only impacts retained: `resolveName` CRITICAL `34/14/2/38`; `resolveGlobalName` CRITICAL `11/3/1/20`; `resolveImportFiles` HIGH `19/12/1/3`; `buildExportTables` LOW `0`. HIGH/CRITICAL được xử lý như blast-radius warning: implementation chỉ chạm exact authorized seams và regression bao phủ resolver/analyze/CLI boundary.

## Production implementation

### Dedicated semantic owner

New `internal/resolution/export_resolution.go` là sole P5-C owner:

- immutable deterministic result model cho `terminal`, `ambiguity`, `cycle`, `missing`, `meaning-mismatch` tại lines `10-82`;
- terminal và namespace candidates, source-backed proof hops, consumer request, target files, failure file/name;
- alias/re-export traversal, explicit-over-star behavior và star default exclusion tại traversal owner bắt đầu line `352`;
- cycle detection bằng active key `(file,name,canonical meanings)`; terminal branch được giữ đồng thời với cycle failure;
- canonical meaning intersection/union xuyên mọi hop;
- same-terminal paths group theo terminal identity nhưng giữ mọi deterministic proof; distinct terminals trả ambiguity và `definition()` không chọn candidate;
- namespace/member traversal dùng cùng semantic resolver tại line `171`;
- explicit-import call-state guard owner tại line `290`;
- nested result ownership: request/proof/fact/meaning/hop và terminal `DefinitionFact` slices, ranges, pointer fields được deep-clone; line `610` là terminal clone owner.

### One two-phase `resolveImports`

`internal/resolution/indexes.go` chỉ đổi tại authorized seams:

1. `resolveImports` line `281`: phase 1 gọi path resolution đúng một lần cho từng import và retain toàn bộ `TargetFiles`.
2. Line `308`: `w.buildExportTables()` đúng một lần sau khi mọi candidate đã tồn tại.
3. Phase 3 dùng semantic export lookup để tạo `TargetDef`/scope binding; không lặp path lookup.
4. `resolveImportedDef` line `477` dispatch semantic TS/JS imports vào P5-C owner, giữ legacy physical path cho non-semantic imports.
5. `resolveImportedMember` line `629` consume shared semantic namespace/member result, fail closed khi semantic import đã claim receiver.
6. Redundant standalone `w.buildExportTables()` sau `w.resolveImports()` bị xóa; `buildWorkspace` vẫn gọi `w.resolveImports()` đúng tại line `174`.

Không có corrective second pass, call swap đơn giản, duplicate path resolution hoặc broad workspace refactor.

### Explicit-import no-global rescue

`internal/resolution/resolve.go` chỉ sửa trong `resolveCall` line `360`. Free/constructor calls kiểm tra `explicitImportCallState` trước repository-global fallback; một lexical inner/local shadow vẫn được phép, còn resolved target không thuộc semantic import hoặc explicit export miss đều không thể được global same-name target cứu. `resolveGlobalCallName`, `resolveGlobalName` và `resolveName` không đổi.

## Source/test manifest trên final bytes

| File | Role | Bytes | LF | CR | BOM | SHA-256 |
|---|---|---:|---:|---:|---|---|
| `internal/resolution/export_resolution.go` | new production semantic owner | 25,380 | 764 | 0 | no | `5494CAE992429F4E2599B51FD9E2B21A39D0AD3128FDC19A337E541580BE9F04` |
| `internal/resolution/indexes.go` | bounded orchestration/consumers | 33,417 | 1,213 | 0 | no | `A44C94FA545FFBAEF7017D49C6529B1B4CFDE82C3BB244EC6F4DD57EBE4283F6` |
| `internal/resolution/resolve.go` | bounded `resolveCall` guard | 20,799 | 668 | 0 | no | `047CCDAE8F451553FC159CCF5B9864FFCA60F597E25FF31BDF5494638EC5817C` |
| `internal/resolution/export_resolution_test.go` | new focused behavior owner | 25,477 | 523 | 0 | no | `09797F5B6AAC7F424AA380C8E9F758F08FE3F45CFD7C53DC6CEA4E317674E82E` |

Tracked diff is bounded to `indexes.go` and `resolve.go`: `65` insertions / `16` deletions. The two dedicated P5-C files are untracked until Main-owned acceptance/commit gates.

## Focused behavior/proof evidence

Production behavior was implemented before focused tests were added.

| Test | File line | What it proves | Result |
|---|---:|---|---|
| `TestP5CResolveAliasChainMatchesDirectTerminalAndRetainsProof` | `export_resolution_test.go:18` | 3-hop alias/re-export proof, direct/barrel terminal identity, `2` CALLS to same terminal, immutable nested terminal data | PASS |
| `TestP5CExplicitExportPrecedesStarAndStarExcludesDefault` | `:81` | explicit entry wins over star; star cannot synthesize default | PASS |
| `TestP5CSameTerminalDedupeDistinctAmbiguityAndCycles` | `:118` | same terminal -> one candidate/two proofs; distinct terminals -> ambiguity/no selection; pure cycle; terminal branch plus retained cycle proof | PASS |
| `TestP5CMeaningMismatchAndNamespaceMemberResolution` | `:223` | value/type mismatch remains explicit; type lane resolves; named namespace/direct namespace converge; type-only namespace cannot resolve value member | PASS |
| `TestP5CExplicitImportMissBlocksGlobalRescueAndPreservesImports` | `:271` | zero false CALLS to global same-name definition; one unresolved diagnostic; `ImportsResolved=1`; syntactic `IMPORTS=1` | PASS |
| `TestP5CPreservesNonSemanticGoPackageMemberResolution` | `:304` | unchanged non-semantic Go package/member import and CALLS behavior | PASS |

Final-byte commands:

```text
cwd: E:\Anvien
go test ./internal/resolution -run '^TestP5C' -count=1 -v
PASS; package duration 0.181s

cwd: E:\Anvien
go test ./internal/resolution -count=1
PASS; package duration 0.212s
```

Adjacent boundary regression run before the final immutable-clone strengthening:

```text
cwd: E:\Anvien
go test ./internal/analyze ./internal/cli -count=1
PASS; internal/analyze 1.935s; internal/cli 139.006s
```

Không dùng adjacent run này như final-byte proof. Final-byte analyze/CLI runtime proof đến từ canonical build script bên dưới, vốn build canonical CLI rồi chạy analyze bằng chính output đó.

## Canonical full-build chronology

Chronology được giữ trung thực:

1. Một invocation đầu FAIL compile vì local `memberProof` nằm ngoài scope; production source được sửa tại đúng local owner.
2. Invocation kế FAIL trước build completion vì canonical executable bị lock.
3. E-only Restart Manager/doctor evidence xác nhận PID `5492` và `9772` là actual build-output owners; chỉ hai PID này bị terminate.
4. Canonical full build PASS trên production+test candidate.
5. Sau PASS đó, `export_resolution.go` được strengthen để deep-clone nested terminal `DefinitionFact`; focused test thêm anti-alias mutation assertions. Đây là source/test byte invalidation thật, không phải formatting hoặc report-only change.
6. Vì invalidation trên, đúng một final canonical rerun là bắt buộc. Không phân loại invocation này là redundant confirmation.
7. Sau final PASS không có thêm source/test byte change và không có thêm build/analyze gate nào được chạy. Main intervention dừng audit/build loop được tuân thủ.

Final build command literal:

```text
cwd: E:\Anvien
powershell -ExecutionPolicy Bypass -File E:\Anvien\scripts\full-build.ps1
exit: 0
scanned / parsed_code / failed: 1,963 / 740 / 0
graph nodes / relationships: 115,800 / 159,445
```

Fresh chronology by final bytes/output timestamps (UTC):

```text
export_resolution_test.go  2026-08-21T14:20:25.4242801Z
export_resolution.go       2026-08-21T14:20:34.1384573Z
anvien.exe                 2026-08-21T14:22:22.2762152Z
graph.json                 2026-08-21T14:24:06.6685078Z
```

Canonical outputs:

| Output | SHA-256 |
|---|---|
| `E:\Anvien\anvien\bin\anvien.exe` | `A3AFBA37318E4762E4F7045FB650A457F7A7ADE1ED4BE78C0D8A157841742A37` |
| `E:\Anvien\anvien\bin\lbug_shared.dll` | `20CBD87840483A2053CFF3FC2DB23A86DD802B8915D86509D41A4B709624CDB7` |
| `E:\Anvien\.anvien\graph.json` | `ACB222619F0E15C888FEB22D4F6382A27B784E253F831679A90E6B751C5BDFE5` |

## Fixed-corpus và preservation evidence

Accepted fixed-corpus artifact remains:

```text
E:\Anvien\.tmp\p5b-fixed-corpus.json
SHA-256 7687C5197AE76DC74294FB948AB9E25124BCEF26E444230EB29A2771EC373796
fixed baseline: physical target-file resolutions 5,072
                resolver-emitted syntactic IMPORTS 5,072
                persisted graph-wide IMPORTS 5,088
```

Final full source corpus reports physical/emitted `IMPORTS = 5,161 / 5,161` and persisted graph-wide `IMPORTS = 5,177`. New repository-owned P5-B/P5-C code/test corpus explains exactly `89` relationships:

```text
P5-B outgoing: export_tables.go 8 + export_tables_test.go 12 = 20
P5-B inbound new-production target = 20
P5-B growth = 40

P5-C outgoing: export_resolution.go 8 + export_resolution_test.go 21 = 29
P5-C inbound new-production target = 20
P5-C growth = 49

5,177 - 40 - 49 = 5,088 persisted fixed baseline
5,161 - 40 - 49 = 5,072 physical/emitted fixed baseline
```

Delta trên fixed `736` parsed-code corpus là `0 / 0 / 0`. Source sequencing confirms `resolveImportFiles` chỉ ở phase 1, `buildExportTables()` chỉ một lần ở phase 2, và phase 3 không thực hiện physical path lookup. P5-C không đổi syntactic `emitImportEdges` behavior.

Preserve-only file identities:

| File | SHA-256 |
|---|---|
| `internal/resolution/export_tables.go` | `BA4C7F9F2D05F6715CA22195C1CFF64C1AACADC3D81E7CA3ED711C8EA8B63E19` |
| `internal/resolution/export_tables_test.go` | `A0395EF9030CB054B09C52AFBE048D2F8244BAD8A924EE492BB7AB312CD08AD8` |
| `internal/resolution/import_resolution.go` | `67C5F10E40E2DCD5850C56622A4D0ED3CBDFE15D84A9419971E97F8799876413` |
| `internal/resolution/emit.go` | `B1B8F22AEFE94F72483CD284B7E267FE3897412C04D581AFAB7F9FE30F0B2867` |

Child 04 fact owners, generic global helpers, graph/persistence/readers, P5-D and target have no worktree change from this lane.

## E2E verification summary

```text
[PASS] Compiled: literal absolute canonical full build -> exit 0 on final source/test bytes
[PASS] Runtime: canonical built CLI analyze inside full-build.ps1 -> 740 parsed-code, 0 failed
[PASS] Happy path: alias/direct/star/namespace -> deterministic terminal identity plus source-backed proofs
[PASS] Edge cases: ambiguity/no selection, pure/terminal cycles, meaning mismatch, explicit-import no-global rescue
[PASS] Preservation: fixed corpus delta 0/0/0; syntactic IMPORTS retained; non-semantic Go regression retained
```

## Git/worktree boundary

- Branch/HEAD: `master` / `861000cb6b6e36ce105623f0dc8c093b089f61fa`.
- Index/staged paths: none.
- Tracked modified: `internal/resolution/indexes.go`, `internal/resolution/resolve.go`.
- Untracked implementation: `internal/resolution/export_resolution.go`, `internal/resolution/export_resolution_test.go`.
- `git diff --check`: PASS.
- Tám protected Main handoffs remain untracked and untouched:
  - `reports/Investigation/rp_main_260821_0631_orchestration_rotation_handoff.md`
  - `reports/Investigation/rp_main_260821_0721_orchestration_rotation_handoff.md`
  - `reports/Investigation/rp_main_260821_1518_orchestration_rotation_handoff.md`
  - `reports/Investigation/rp_main_260821_155017_orchestration_rotation_handoff.md`
  - `reports/Investigation/rp_main_260821_163855_orchestration_rotation_handoff.md`
  - `reports/Investigation/rp_main_260821_172833_orchestration_rotation_handoff.md`
  - `reports/Investigation/rp_main_260821_195827_orchestration_rotation_handoff.md`
  - `reports/Investigation/rp_main_260821_204245_orchestration_rotation_handoff.md`
- Không stage, cleanup, detect, commit, push, reset hoặc checkout.

## Risks/open points và next handoff

- Slice-local residual behavior risk: none identified by the required matrix.
- Accepted HIGH/CRITICAL blast radius remains the reason Supervisor must independently inspect source, proof semantics, two-phase orchestration, no-global guard, fixed-corpus arithmetic, and current Git boundary.
- Report này không self-accept P5-C và không mở P5-D.
- Exact next action: Main resume đúng existing Supervisor lane với report bất biến này; chỉ sau Supervisor verdict mới quyết định detect/commit/ledger transition.

`READY_FOR_SUPERVISOR`
