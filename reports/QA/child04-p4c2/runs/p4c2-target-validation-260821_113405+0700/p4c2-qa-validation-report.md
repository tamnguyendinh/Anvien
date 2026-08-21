# Child 04 P4-C2 real-target QA validation

Status: **BLOCKED**

Đây là kết quả QA/fail-closed, không phải verdict nghiệm thu Supervisor. Slice được kiểm tra là Child 04 P4-C2, với `E4-P4C2-ORACLE1` đã seal; `E4-P4C2-TARGET1` bị chặn tại normal analyze boundary; phần preservation của `E4-P4C2-BOUNDARY1` đạt nhưng graph/persistence/read parity chưa thể đánh giá.

## Kết luận điều hành

Canonical full build PASS, built runtime được khóa chính xác, và đúng một invocation target đã được thực hiện:

```text
anvien analyze E:\cheapapp.org --force
```

Invocation duy nhất thoát `1` sau `0.17s` với output:

```text
not a git repository: E:\cheapapp.org; pass --skip-git to index any folder without a .git directory
```

Target thực tế có `.git`, HEAD và status đúng sealed basis. Git CLI chỉ đọc được repo trong user context hiện tại khi dùng override invocation-local `-c safe.directory=E:/cheapapp.org`; normal Anvien command bị khóa không có override đó và phân loại Git dubious-ownership thành “not a git repository”. QA không được sửa global/local Git config, không được thêm `--skip-git`, và không được chạy invocation thứ hai. Vì vậy không có fresh graph hợp lệ để so 21 positive + 11 negative; mọi artifact target cũ đều bị loại khỏi evidence.

Blocker chính xác: **current Windows identity không trust ownership của `E:\cheapapp.org`, trong khi exact normal analyzer invocation không có cơ chế trust override và đã tiêu thụ invocation duy nhất của run này.**

## Seal gate — PASS và đóng

- Bundle: `E:\Anvien\reports\QA\child04-p4c2\oracle\p4c2-oracle-v1-a869876ab626-260821_110849+0700`
- Oracle ID: `p4c2-oracle-v1-a869876ab626-260821_110849+0700`
- Status/sealed-at: `SEALED`, `2026-08-21T11:21:54.109+07:00`
- `seal.json`: `2,550` bytes, SHA-256 `00FFA78CAB1B584FB9290EEF8578CBF07B52A86351E9327DA60C1BE39956FE4F`
- Non-circular bundle digest computed/declared: `7749AB14E02FBF61CF7F81B3A7638888F8AD414064F9F969CFED06EB11355439` / same, match `true`.
- Construction checked: prefix `p4c2-oracle-v1\n`; ordinal filename order; UTF-8 `filename NUL byte_count NUL SHA256_UPPER_HEX LF`; `seal.json` excluded.
- PowerShell `Test-Json -SchemaFile` result for `expected-values.json`: `True`.
- Counts: positive `21`, negative `11`.
- Target basis: HEAD `a869876ab6262dacde6cd5d432d099a91852a646`, branch `master`, tracked-status canonical SHA-256 `FCB5AD9155C029FFA6B3D80B8AADD1EE40EDC3B408EA401992E8FEA048E4C5E1`.

Sealed inventory:

| File | Bytes | SHA-256 |
|---|---:|---|
| `authoring-report.md` | 3,293 | `DAFD82D01A6E5C70D269CEE21CCD42BDE8D1363F3E3C474F60AE5D9E258F38EA` |
| `expected-values.json` | 35,934 | `1DC54F773455A031C9FBC844F5ACEB7D107EF24848649FC2F1188EA05819A66B` |
| `oracle.schema.json` | 8,233 | `388E7E187DB876E48FC85E0AF71A8BBE989B38AE7D96CE6DB4B7D7DC8CD5E96A` |
| `provenance.json` | 10,366 | `2D3BFE247667B2FE47AF674E5BF49266CECA5696CAA27C1A458FAA10462B7639` |
| `seal.json` | 2,550 | `00FFA78CAB1B584FB9290EEF8578CBF07B52A86351E9327DA60C1BE39956FE4F` |
| `source-basis.json` | 8,769 | `1DD90BBFB664EEC3EDE43A1A7949DF59166DB7C7CA23732ABEA5C973D05213A8` |

Required source identities locked by the seal:

| Path | Bytes | SHA-256 |
|---|---:|---|
| `modules/email/server/operations/email-operations-observability.ts` | 19,328 | `A9247972D44D0A74C411FB7CF83D681FEBA371C7A85EEB83F8986AFD6446E40C` |
| `modules/release-distribution/server/release-distribution-publication-state.ts` | 405 | `B53F966F82AEDAC58EC15A892B9B71BA0A5716638703105BB35B126CB398D474` |
| `modules/commercial-config/server/admin-commercial-config/read-admin-commercial-config.ts` | 828 | `44467B1D6F1778AE4728B4941136321F7831D8FD67EBE2C93A430A3D132F90AC` |

## Target pre-state — PASS

- Resolved target: `E:\cheapapp.org`.
- HEAD/branch: `a869876ab6262dacde6cd5d432d099a91852a646` / `master`.
- Tracked status: exact `7` pre-existing modified paths, canonical SHA-256 `FCB5AD9155C029FFA6B3D80B8AADD1EE40EDC3B408EA401992E8FEA048E4C5E1`.
- Index changed paths: `0`.
- Sealed source hashes: `3/3` exact.
- Git-visible non-`.anvien` manifest: `2,377` entries, SHA-256 `CAF9F5798C53F5FAD491E31DFD2295B9C49FA32B25705EE45D1136609A8DF702`.
- `.anvien` pre-state: `4` files; canonical manifest SHA-256 `90B264B6C927D4E32BAD276C0E984CD1F0EB4714FFF772B6273E931E381998AF`.

The four pre-existing analyzer-owned files were:

| Path | Bytes | SHA-256 |
|---|---:|---|
| `.anvien/graph.json` | 423,734,194 | `5D82C464C291ED7734E9BF32C3F79CF2350008C62F83F06B67517FD52EF97E05` |
| `.anvien/lbug` | 144,744,448 | `B3173C66D0273149B3AD63D11D529A5B15B9FDA098408922704A32E938550EE8` |
| `.anvien/meta.json` | 259 | `226BB1FDC696C15153B650EDE607DF067C1C78EA218BA7C5FF72CB8B939ECFEC` |
| `.anvien/settings.json` | 31 | `6C831C0AC49625D09EE54E76042612B92F47ABC47E4A90D1310F0EA45E824B3C` |

## Build-holder gate — PASS

- `anvien doctor locks --repo E:\Anvien --json`: exit `0`, lock status `free`, lock file absent.
- `anvien doctor processes --json`: identified Anvien Playwright test-server PID `17920` as the non-editor-owned runtime/build holder.
- PID `17920` was terminated; absence was verified.
- Post-cleanup locks/processes: exits `0/0`, lock remains free, build-process pattern count `0`.
- Editor-owned MCP processes were preserved as instructed.

## Canonical full build — PASS

- Command: `npm run full-build` from `E:\Anvien`.
- Exit: `0`; duration `180.018s`.
- Packaged CLI: `1.2.8`.
- Web build: `2,943` modules; Vite build `38.13s`; retained chunk/dynamic-import warnings were non-failing.
- Canonical build's fresh Anvien self-analysis: scanned/parsed/failed `1,881/735/0`; graph `113,684/156,191` nodes/relationships.

Built runtime identity used by the target invocation:

- Resolved command: `C:\Users\TAM NGUYEN\AppData\Roaming\npm\anvien.ps1` (`482` bytes; SHA-256 `29ACAEB5A69B36A9FD34D3F782C6852C50777C394DF30B3EF3326C7C9C841560`).
- Local built runtime: `E:\Anvien\anvien\bin\anvien.exe`.
- Global invoked runtime: `C:\Users\TAM NGUYEN\AppData\Roaming\npm\node_modules\anvien\bin\anvien.exe`.
- Both binaries: `71,371,776` bytes; SHA-256 `23D647FC27BDDE841EC3F94316A6BFA7ADA2F24E87BF15D236FB86743BD6E8C6`.

## Target analyze — FAIL-CLOSED

- Invocation count: exactly `1`.
- Exact command: `anvien analyze E:\cheapapp.org --force`.
- Runtime: freshly built `1.2.8`, binary identity above.
- Exit/duration: `1` / `0.17s`.
- No retry, `--skip-git`, custom analyzer, stale output substitution, config mutation, or graph-backed follow-up was performed.

Required result matrix:

| Requirement | Result |
|---|---|
| Positive oracle rows | **NOT EVALUATED** (`0/21` evaluated; no fresh output) |
| Negative controls | **NOT EVALUATED** (`0/11` evaluated; no fresh output) |
| Exact kind/name/meaning/typeOnly/access | **NOT EVALUATED** |
| Compatibility/access separation | **NOT EVALUATED** |
| Graph JSON ↔ Ladybug ↔ affected reader parity | **NOT EVALUATED** |
| Diagnostics/orphans | **NO CLAIM** |
| Forbidden terminal/resolved-target/public-API state | **NO CLAIM** |

These are deliberately not marked contract failures: the analyzer never produced a fresh graph. P4-C2 acceptance remains unadjudicated.

## Target post-state and contamination — PASS

- HEAD and branch preserved.
- Exact seven-line tracked status preserved; index remains empty.
- Required source files preserved `3/3`.
- Non-`.anvien` pre/post entry counts: `2,377/2,377`; manifest hashes identical; diff count `0`.
- `.anvien` pre/post manifest hashes identical; diff count `0`. The failed command created no analyzer output.
- No target-side report, fixture, probe, debug file, source edit, Git config edit, stage, commit, reset, checkout, or push occurred.

## Repository boundary and provenance

- Current Anvien HEAD/branch at handoff preparation: `e32a412b289453a530bc71b93320ef2b97b3a97a` / `master`.
- Tracked boundary remains exactly five unstaged Child 04 living documents; index is empty.
- The sealed oracle bundle and all pre-existing untracked provenance were preserved. This run added only the durable QA run directory under `reports/QA/child04-p4c2/runs/`.
- No evidence-bearing artifact was created in `.tmp` or under the target. Canonical build used its normal repo-local `.tmp` runtime dependency path; it was not used as QA evidence. No dead debug artifact was created by this lane.

Non-mutating execution deviations retained for provenance:

1. At re-anchor, `working-rules` was read before `AGENTS.md`; `AGENTS.md` immediately followed and no task action occurred.
2. Before the explicit PowerShell-validator correction, a missing-AJV resolution check and a read-only Python schema validation were run; no package was installed and no artifact was created. The authoritative schema result is PowerShell `Test-Json=True`.
3. `source-basis.json` was read once after the seal gate was declared locked; this was read-only and was not used to mutate expected values.
4. Initial run-parent creation and initial target Git capture attempts failed before writes (missing parent; Git dubious ownership). The approved run directory and successful pre-state capture are the artifacts retained here.
5. One holder-cleanup command had a PowerShell parse error before process action; the corrected command then terminated PID `17920` and recorded the post-cleanup PASS evidence.

## Handoff

- State: **BLOCKED**.
- Active lane: Child 04 P4-C2 QA only; no Supervisor lane is open.
- Next owner: **Main** for self-verification.
- Completion condition still required: a fresh successful exact normal target analyze, complete immutable `21/21` + `11/11` comparison, record-level Graph JSON/Ladybug/affected-reader parity, and required zero-state checks.
- Exact next action: Main/Owner must resolve repository trust/ownership for `E:\cheapapp.org` outside this QA lane without changing its sealed HEAD/status/source basis. Only then may Main open a new P4-C2 QA run with one fresh exact `anvien analyze E:\cheapapp.org --force`. Do not open P4-C2 Supervisor or Child 05 before that run completes.
