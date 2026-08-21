# Child 05 / P5-C Ambiguous Owner Member — Coder Resubmission

## Trạng thái và authority

- Status: `READY_FOR_SUPERVISOR`.
- Scope duy nhất: sửa Supervisor-rejected invariant về distinct-owner ambiguity trên semantic namespace/member surface.
- Checkout: `E:\Anvien`; branch/HEAD: `master` / `861000cb6b6e36ce105623f0dc8c093b089f61fa`.
- Accepted predecessor: P5-B commit `c1559df953a277b099009f8489576d00ed25aa58`.
- Supervisor REJECT authority: `E:\Anvien\reports\Supervisor\rp_supervisor_260821_220516_by_gpt-5_child05_p5c_export_resolution_reject.md`, `8,918` bytes / `81` LF / `0` CR / UTF-8 không BOM / SHA-256 `1B8E32DA433782116A70BE427962A12DD799FF26B431BF7982A2D32FCE3121F1`.
- Accepted impact authority: `E5-P5C-IMPACT1`; không rerun analyze/file-detail/impact inventory.
- P5-C vẫn unchecked. P5-D, target, `E:\cheapapp.org`, ledger edit, detect, stage/commit và Supervisor call vẫn khóa.

## Invariant Family Map

- Family: proof-bearing semantic member composition sau repository export-owner resolution.
- Authority / SSOT: accepted Child 04 `ScopeIR.ExportFact`, accepted P5-A requested meanings, accepted P5-B export tables, P5-C deterministic result/proof model và exact Supervisor REJECT ở trên.
- Rejected path: ambiguous owner result A/B -> flatten mọi owner terminal proof -> chỉ A có member `run` -> branch B bị drop -> aggregate thấy một member terminal -> `definition()` chọn A.
- Required path: mỗi distinct owner candidate được compose riêng -> A tạo terminal-member proof -> B tạo owned missing-member proof -> owner ambiguity được giữ ở aggregate member result -> `definition()` false -> `resolveImportedMember` false -> không physical/global fallback và không emitted relationship.
- Sibling surfaces checked: semantic result outcome/candidates/failures; `definition()`; production `resolveImportedMember` consumer; CALLS emission; ACCESSES emission; complete legacy P5-C matrix; full resolution package; fixed-corpus path/IMPORTS preservation.
- Preserve-only: `indexes.go`, `resolve.go`, P5-B tables/tests, import path owner, emit owner, Child 04/P5-A facts, graph/persistence/readers, bốn ledgers và toàn bộ existing Main/Coder/Supervisor reports.
- Forbidden fallback: không physical definition scan hoặc repository-global rescue sau semantic receiver claim; không chọn sole surviving member từ ambiguous owners; không bỏ missing owner branch.
- Stale test gap closed: prior suite chỉ có direct terminal ambiguity và unambiguous namespace/member. New focused fixture owns ambiguous-owner member composition.
- Residual unverified surfaces trong rejected invariant: `none`. Supervisor re-review vẫn là independent acceptance gate; report này không self-accept P5-C và không mở P5-D.

## Bounded production fix

Editable production owner duy nhất: `internal/resolution/export_resolution.go`.

`resolveSemanticImportedMember` giờ:

1. Retain `ownerResult.Outcome == ambiguity` thay vì bỏ qua outcome.
2. Copy owner failures trực tiếp và iterate từng `ownerResult.Candidates`/proof branch thay vì flatten qua `ownerResult.allProofs()`.
3. Compose member terminal theo từng owner proof.
4. Khi một definition owner không có requested member, tạo `missingOwnedMemberProof` chứa member hop với exact `MemberOwnerDefID`, member name, owner file và inherited export hops; branch không còn bị drop.
5. Sau deterministic aggregation, force member result outcome về `ambiguity` nếu owner level đã ambiguous. Sole member candidate được giữ làm provenance nhưng `definition()` không thể select vì outcome không phải terminal.

Exact source locations trên final file:

- `ownerAmbiguous` tracking: lines `181`, `226`, `273`.
- per-candidate member composition và per-owner failure: lines `228-268`.
- `missingOwnedMemberProof`: line `283`.

Không sửa `indexes.go`: existing semantic dispatch vẫn fail-close vì `definition()` false và handled=true, do đó không rơi xuống physical scan. Không sửa `resolve.go` hoặc generic global helpers.

## Code-first chronology

Production code được sửa trước. Sau production patch + gofmt:

```text
export_resolution.go SHA-256: 566A69B965212D94E87E4D46703FC4E4C80E910FDC8657449D3914D5B868248F
export_resolution_test.go SHA-256 vẫn là rejected-candidate identity:
09797F5B6AAC7F424AA380C8E9F758F08FE3F45CFD7C53DC6CEA4E317674E82E
```

Chỉ sau đó focused regression được thêm vào test owner.

## Focused regression fixture

New test: `TestP5CAmbiguousOwnerMemberPreservesNoSelectionAndEveryBranch`, `internal/resolution/export_resolution_test.go:271`.

Fixture dùng accepted export-table path:

```text
consumer imports named api from ./barrel
./barrel export-star -> ./left and ./right
./left exports LeftAPI as api; LeftAPI owns method run
./right exports RightAPI as api; RightAPI has no run
```

Assertions PASS:

- semantic member result outcome là `ambiguity`;
- sole branch member vẫn có candidate/proof provenance nhưng `definition()` trả false;
- `resolveImportedMember` trả false;
- không `CALLS` từ consumer tới `LeftAPI.run`;
- không `ACCESSES` từ consumer tới `LeftAPI.run`;
- metrics ghi `ResolvedCalls=0`, `ResolvedAccesses=0`, hai unresolved source sites;
- proof/failure member hops chứa cả owner IDs: Left owner -> terminal, Right owner -> missing.

## Final source/test manifest

| File | Final bytes | LF | CR | BOM | Final SHA-256 | Rejected identity -> final delta |
|---|---:|---:|---:|---|---|---|
| `internal/resolution/export_resolution.go` | 26,015 | 780 | 0 | no | `566A69B965212D94E87E4D46703FC4E4C80E910FDC8657449D3914D5B868248F` | 25,380 bytes / 764 LF / `5494CAE9...` -> `+635` bytes / `+16` LF |
| `internal/resolution/export_resolution_test.go` | 29,480 | 612 | 0 | no | `97FF4990066179AE779C9354D7A59859F3CA437826062DA0DA19C54776BC1620` | 25,477 bytes / 523 LF / `09797F5B...` -> `+4,003` bytes / `+89` LF |

Follow-up editable manifest là đúng hai file trên. Production delta chỉ nằm trong bounded member composition + helper; test delta chỉ là một focused regression fixture. Không file khác được lane này sửa.

Preserve-only identities vẫn exact:

| File | Bytes / LF | SHA-256 |
|---|---|---|
| `internal/resolution/indexes.go` | 33,417 / 1,213 | `A44C94FA545FFBAEF7017D49C6529B1B4CFDE82C3BB244EC6F4DD57EBE4283F6` |
| `internal/resolution/resolve.go` | 20,799 / 668 | `047CCDAE8F451553FC159CCF5B9864FFCA60F597E25FF31BDF5494638EC5817C` |
| `internal/resolution/export_tables.go` | 7,213 / 248 | `BA4C7F9F2D05F6715CA22195C1CFF64C1AACADC3D81E7CA3ED711C8EA8B63E19` |
| `internal/resolution/export_tables_test.go` | 9,751 / 249 | `A0395EF9030CB054B09C52AFBE048D2F8244BAD8A924EE492BB7AB312CD08AD8` |
| `internal/resolution/import_resolution.go` | 13,091 / 427 | `67C5F10E40E2DCD5850C56622A4D0ED3CBDFE15D84A9419971E97F8799876413` |
| `internal/resolution/emit.go` | 26,772 / 815 | `B1B8F22AEFE94F72483CD284B7E267FE3897412C04D581AFAB7F9FE30F0B2867` |

## Final-byte tests

```text
cwd: E:\Anvien
go test ./internal/resolution -run '^TestP5C' -count=1 -v
exit: 0
result: all seven top-level P5-C tests plus four cycle/dedupe subtests PASS
package: ok github.com/tamnguyendinh/anvien/internal/resolution 0.261s
```

Matrix giữ PASS: alias/direct terminal identity và proof; explicit-over-star; star default exclusion; same-terminal dedupe; direct distinct-terminal ambiguity/no selection; pure and terminal cycles; meaning mismatch; namespace/member; new ambiguous-owner member fixture; explicit-import no-global rescue; syntactic IMPORTS; non-semantic Go regression.

```text
cwd: E:\Anvien
go test ./internal/resolution -count=1
exit: 0
result: ok github.com/tamnguyendinh/anvien/internal/resolution 0.316s
```

Source/test không đổi sau hai commands này.

## Canonical build và lock chronology

Pre-build E-only evidence:

- `E:\Anvien\anvien\bin\anvien.exe doctor locks --repo E:\Anvien --json`: analyze lock `free`, file không tồn tại.
- `doctor processes --json`: chỉ global npm MCP editor-owned processes.

First canonical attempt dùng đúng literal command nhưng FAIL vì canonical executable bị lock trong `npm install`; đây không phải source/test failure. Restart Manager query trực tiếp trên `E:\Anvien\anvien\bin\anvien.exe` xác nhận actual owners:

```text
PID 11504  C:\Users\TAM NGUYEN\AppData\Roaming\npm\node_modules\anvien\bin\anvien.exe mcp
PID 12932  C:\Users\TAM NGUYEN\AppData\Roaming\npm\node_modules\anvien\bin\anvien.exe mcp
exclusiveOpen=false
```

Chỉ hai confirmed PIDs bị terminate. Post-remediation Restart Manager result: `ownerPids=[]`, `exclusiveOpen=true`. Không terminate parent/editor hoặc MCP PID không giữ resource.

Required retry:

```text
cwd: E:\Anvien
powershell -ExecutionPolicy Bypass -File E:\Anvien\scripts\full-build.ps1
exit: 0
version: 1.2.8
scanned / parsed_code / failed: 1,966 / 740 / 0
graph nodes / relationships: 115,902 / 159,581
graph path: E:\Anvien\.anvien\graph.json
```

Retry là recovery sau concrete lock FAIL, không phải duplicate confirmation sau PASS. Không có source/test byte invalidation sau PASS nên không rerun build/analyze.

Final chronology UTC:

```text
export_resolution.go       2026-08-21T15:13:19.5963808Z
export_resolution_test.go  2026-08-21T15:14:32.7154315Z
anvien.exe                 2026-08-21T15:19:05.2555923Z
graph.json                 2026-08-21T15:20:36.6854798Z
```

Canonical outputs:

| Output | Bytes | SHA-256 |
|---|---:|---|
| `E:\Anvien\anvien\bin\anvien.exe` | 71,478,272 | `A55D9CE575EAD60EE07612B882448B69FC0272712B5AEE08EF2971C4577A4E62` |
| `E:\Anvien\anvien\bin\lbug_shared.dll` | 20,230,656 | `20CBD87840483A2053CFF3FC2DB23A86DD802B8915D86509D41A4B709624CDB7` |
| `E:\Anvien\.anvien\graph.json` | 462,380,776 | `0A16E835901EB134023CC85C14F7900BBAC1B7DCD7E15E8F329FAB59227847B2` |

## Fixed-corpus preservation

Accepted artifact remains byte-identical:

```text
E:\Anvien\.tmp\p5b-fixed-corpus.json
9,386 bytes / 365 LF / 0 CR / no BOM
SHA-256 7687C5197AE76DC74294FB948AB9E25124BCEF26E444230EB29A2771EC373796
parsed_code=736
physical target-file resolutions=5,072
resolver-emitted syntactic IMPORTS=5,072
accepted persisted graph-wide IMPORTS=5,088
```

Fresh post-build graph query:

```text
MATCH ()-[r]->() WHERE r.type = 'IMPORTS' RETURN count(r) AS imports
imports=5,177
```

Exact current decomposition was rechecked:

```text
P5-B outgoing: export_tables.go 8 + export_tables_test.go 12 = 20
P5-B incoming production owner: 20
P5-B growth: 40

P5-C outgoing: export_resolution.go 8 + export_resolution_test.go 21 = 29
P5-C incoming production owner: 20
P5-C growth: 49

5,177 - 40 - 49 = 5,088
```

Follow-up adds no import statement, path lookup, table construction hoặc emit change. Fixed-corpus physical/emitted/persisted delta remains `0 / 0 / 0`.

## E2E verification

```text
[PASS] Compiled: canonical absolute E-only build -> exit 0 on final bytes
[PASS] Primary: ambiguous owner A/B, only A owns run -> semantic member outcome ambiguity
[PASS] No-selection: definition() false and resolveImportedMember false
[PASS] Proof completeness: A terminal member hop + B owned missing-member hop
[PASS] Runtime boundary: Resolve emits neither CALLS nor ACCESSES to sole A.run branch
[PASS] Fallback: handled semantic ambiguity does not enter physical/global rescue
[PASS] Regression: complete P5-C matrix and full internal/resolution package
[PASS] Preservation: fixed-corpus IMPORTS/path delta 0/0/0 and locked owner hashes exact
```

## Git/worktree boundary

- Branch/HEAD: `master` / `861000cb6b6e36ce105623f0dc8c093b089f61fa`.
- Staged/index paths: none.
- Lane-owned follow-up edits: only the two untracked P5-C source/test files identified above.
- Pre-existing tracked candidate files `indexes.go` and `resolve.go` remain modified with Supervisor-cleared hashes.
- Four Main-owned Child 05 ledgers were already modified when this reject repair resumed; they remain preserve-only and untouched by this lane.
- Nine Main rotation handoffs, prior immutable Coder report và immutable Supervisor REJECT remain untracked and untouched.
- `git diff --check`: PASS.
- Không ledger edit, detect, stage/commit, push/reset/checkout, target access, P5-D action hoặc broad cleanup.

## Handoff

- Rejected invariant đã được sửa và verified, nhưng chưa được self-accepted.
- Exact next action: Main task `01a024bc-cd0e-7521-942b-4a7049d8b41e` resume đúng existing Supervisor lane để re-review duy nhất invariant này.
- Sau handoff lane state: `IDLE / READY_FOR_SUPERVISOR`.

`READY_FOR_SUPERVISOR`
