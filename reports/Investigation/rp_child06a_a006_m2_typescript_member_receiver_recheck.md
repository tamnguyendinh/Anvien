# Child 06A A006-M2 TypeScript Member Receiver Recheck

Measurement ID: `A006-M2-D001-TYPESCRIPT-MEMBER-RECEIVER-RECHECK`

Scope: `P2-A / A006 / B1-P1A-OP001 resolution / B2-P2A-A001-D001 resolve_calls`.

This is attribution-only evidence. It is not a production attempt, candidate, accepted-baseline comparison or promotion, Supervisor event, streak/checklist transition, or disposition. Targets remain separate; instrumented D001, parent, analyzer, process, and resource values are observer facts only.

## Authority and starting gate

- Authority checkpoint: `7764ebf69ce4a155d11caa253b8b16e378915bf1`.
- Architect authority ended in `ARCHITECT_A006_M2_READY`; Main's handoff review was PASS.
- Work root `E:\Anvien\.tmp\child06a_a006_m2_typescript_member_recheck_20260827` and this report were absent before creation.
- Worktree and staged set were empty; `anvien --help` exited `0`.
- Architect blast radius was acknowledged: file `HIGH`, D001 `resolveCall` `CRITICAL`. Only isolated overlay copies were edited.
- No graph-query/impact command, production/test/script/plan/ledger/target-source edit, detect, stage, commit, M3, retry, or other lane occurred.

## Overlay identity and fixed lifetime

Both files were copied mechanically from the accepted M1 overlay. `gofmt` touched only the new copies.

| Owner | M1 seed SHA-256 | M2 overlay SHA-256 |
|---|---|---|
| `overlay/internal/resolution/resolve.go` | `6B77B55097664234AB9FA28102C84702F3A8B830D36C7449AF2F676242DB8A61` | `FFEAC3F676736D1439796B90009C0DDA47921A7DC4321F985B4C3EA502ED3F0B` |
| `overlay/internal/resolution/types.go` | `8F5F8E7B2B29FDDB7A65384970DAC60E7E15882608D9024FF6DF7414DD9F4E2D` | `13DFD438043789907714F356F05AD16117E03FB32F5534B3F2B9DDDBB62B14ED` |

Mapped production starts remained exact: `resolve.go=8CEEDBA1883314EE8883320D3647C25DEF6F19F043D57881A893FBA73BA210D9`; `types.go=7C5E113F5E50584665D6D9AED0BCB2B3C6F6085219A62CD5AE74DD7654CD5DC3`.

Manifest `overlay/optimized-overlay.json` is `354` bytes, SHA-256 `918693F924BCB6C4F028F5CC5F5F15604C5ED1E81264DBEBA16532B3A9FD02D9`, and contains only one `Replace` object with exactly the two canonical-to-overlay mappings.

The existing M1 ten-row metric remains intact. M2 adds only `resolveCallTypeScriptLookupRecordMeasurement`: one fixed `[7]` subgroup array and scalar work counters. D001 passes the run-local recorder; D002/D003 pass `nil`. The helper retains its nil/library/language guards and exact result path. One adjacent empty timer control and only the direct receiver predicate are timed. There is no slice growth, map/cache, per-site data, global/cross-run state, log, I/O, goroutine, lock, or target state. Lifetime is fixed O(1), run-local, and unreachable after exit.

## One-shot build identity

Builder `scripts/build-a00x-benchmark.ps1` remained unchanged at SHA-256 `ADA407C7496FCEA988276F03BAD5001ED139A4AEC9A16B9C32947DA440814EC5`. Output, DLL, and provenance were absent; build competitor count was zero. It was invoked exactly once, without retry:

```powershell
& 'E:\Anvien\scripts\build-a00x-benchmark.ps1' `
  -AttemptId 'A006' `
  -OverlayManifestPath 'E:\Anvien\.tmp\child06a_a006_m2_typescript_member_recheck_20260827\overlay\optimized-overlay.json' `
  -OutputExecutablePath 'E:\Anvien\.tmp\child06a_a006_m2_typescript_member_recheck_20260827\frozen-build\anvien-a006-m2-benchmark.exe' `
  -ExpectedOverlaySha256 '918693F924BCB6C4F028F5CC5F5F15604C5ED1E81264DBEBA16532B3A9FD02D9' `
  -ExpectedMappedSourceHash @(
    'E:\Anvien\internal\resolution\resolve.go=FFEAC3F676736D1439796B90009C0DDA47921A7DC4321F985B4C3EA502ED3F0B',
    'E:\Anvien\internal\resolution\types.go=13DFD438043789907714F356F05AD16117E03FB32F5534B3F2B9DDDBB62B14ED'
  ) `
  -ExpectedCandidateSourceHash @(
    'internal/graphhealth/diagnostics.go=6DE54D97A1B95E877B686DC3459238598F6060D28CE6329357A80BD7FC376D30',
    'internal/resolution/emit.go=73AE20E8A18AC25F8019E2A3223407621AE7ABC358DBAF4F058E1D9AC3136060',
    'internal/resolution/outcome.go=02092F9FE7DA2A4BDB49E13056FBA3C97DC24F416141E7F27866EE80F60C1F7E',
    'internal/resolution/export_binding_proof.go=4BFE1976766FBF2E2257102070FF43CAC6D757E32E71DD583F8824781EAB6A8E'
  ) `
  -ExpectedNativeHash @(
    'lbug.h=3D5114D0863B3DAB3B28BD2FEC97A52E6CF669213739921A01814A5BBF5525EB',
    'lbug_shared.lib=B18AAFC0B712DC1C4CB9DD25F76C3828282D7D460627980E3A4B16EFCD98A955',
    'lbug_shared.dll=20CBD87840483A2053CFF3FC2DB23A86DD802B8915D86509D41A4B709624CDB7'
  ) `
  -GoExecutable 'C:\Program Files\Go\bin\go.exe'
```

Build marker `A00X_BENCHMARK_BUILD_COMPLETE` and exits passed. Provenance schema/attempt/build/root exits are `1 / A006 / 0 / 0`; identities are `2 / 4 / 3`; all expected/actual hashes match.

| Artifact | Exact identity |
|---|---|
| Executable | version `1.2.8`; `73,846,784` bytes; `9455E26AD0C862D7D528B7E1FD37806DAEA668F5C05486575D24266A64868ABA` |
| Adjacent DLL | `20,230,656` bytes; `20CBD87840483A2053CFF3FC2DB23A86DD802B8915D86509D41A4B709624CDB7` |
| Provenance | `13,697` bytes; `FCBAF3C3A7511AF302ECBB1472E0C4250438BFBF7C06BF9C4C5988BB394BEC6D` |
| Build interval | `2026-08-27T02:18:46.0960126Z` to `2026-08-27T02:21:28.0509043Z` |

## Sequential process identity

Both immediate competitor lists were empty. Each target had exactly one direct `ProcessStartInfo` launch with `UseShellExecute=false`, `CreateNoWindow=true`, redirected streams, and ordered `ArgumentList` entries.

| Target | PID / exit | Start | End | Wall / CPU / user / kernel (ns) |
|---|---|---|---|---|
| Cheapapp | `10968 / 0` | `2026-08-27T02:23:43.6693914Z` | `2026-08-27T02:25:35.6024842Z` | `111,933,092,800 / 100,953,125,000 / 85,953,125,000 / 15,000,000,000` |
| Restaurant Manager | `14464 / 0` | `2026-08-27T02:29:09.2313354Z` | `2026-08-27T02:31:05.8457853Z` | `116,614,449,900 / 109,171,875,000 / 93,015,625,000 / 16,156,250,000` |

Restaurant Manager started `213,628,851,200 ns` after Cheapapp ended. Cross-target overlap is `0`. Both target HEAD/status pairs and repository HEAD were unchanged; staged counts remained zero.

Exact Cheapapp command:

```text
E:\Anvien\.tmp\child06a_a006_m2_typescript_member_recheck_20260827\frozen-build\anvien-a006-m2-benchmark.exe analyze E:\cheapapp.org --force --skip-git --json --progress --benchmark-json E:\Anvien\.tmp\child06a_a006_m2_typescript_member_recheck_20260827\captures\cheapapp\candidate\benchmark.json --benchmark-label child06a-a006-m2-cheapapp-ts-member-recheck-20260827
```

Exact Restaurant Manager command:

```text
E:\Anvien\.tmp\child06a_a006_m2_typescript_member_recheck_20260827\frozen-build\anvien-a006-m2-benchmark.exe analyze E:\Restaurant_manager --force --json --progress --exclude electron/renderer/src/api/userApi.ts --benchmark-json E:\Anvien\.tmp\child06a_a006_m2_typescript_member_recheck_20260827\captures\restaurant_manager\candidate\benchmark.json --benchmark-label child06a-a006-m2-restaurant-manager-ts-member-recheck-20260827
```

All process-local `TEMP/TMP/TMPDIR`, home/profile, AppData, XDG, Go cache/module/temp, and profile paths were under the corresponding capture root. Restaurant also used only the exact run-local safe-directory environment `1 / safe.directory / E:/Restaurant_manager`.

## Cheapapp packet

### Parent and retained M1 conservation

| Gate | Exact value | Result |
|---|---:|---|
| Operations | `30 / 30`, ordered exact | PASS |
| Resolution children | `17 / 17`, ordered exact | PASS |
| Child intervals / overlap | `2,675 / 0` | PASS |
| Parent / child sum / residual | `18,251,352,700 / 18,239,607,800 / 11,744,900 ns` | PASS |
| D001 denominator | `calls=27,890; files=887` | PASS |
| D001 `resolve_calls` | `2,990,982,500 ns` | PASS |
| Analyzer / process | `92,522,196,800 / 111,933,092,800 ns` | attribution only |

| Order | M1 group | Duration (ns) | Invocations |
|---:|---|---:|---:|
| 1 | `source_context` | `48,748,500` | `27,890` |
| 2 | `binding_receiver` | `51,230,600` | `26,079` |
| 3 | `scoped_same_file` | `40,383,000` | `29,686` |
| 4 | `member_import` | `190,904,800` | `47,123` |
| 5 | `go_same_package` | `4,610,800` | `8,501` |
| 6 | `global_lookup` | `5,661,700` | `4,118` |
| 7 | `typescript_lookup_record` | `764,972,500` | `23,430` |
| 8 | `evidence_emission` | `1,854,661,700` | `44,488` |
| 9 | `direct_site_identity` | `22,892,400` | `40,360` |
| 10 | `resolve_call_residual` | `6,916,500` | `27,890` |

M1 sum is exactly `2,990,982,500 ns`, equal to D001; all durations/remainders are nonnegative and overlap is `0`.

### Seven-group split and work equations

| Order | M2 subgroup | Duration (ns) | Invocations |
|---:|---|---:|---:|
| 1 | `member_receiver_recheck` | `10,466,100` | `3,549` |
| 2 | `member_receiver_recheck_timer_control` | `1,000,600` | `3,549` |
| 3 | `member_lookup_remainder` | `9,434,200` | `3,692` |
| 4 | `global_lookup` | `15,879,200` | `4,118` |
| 5 | `site_target_text` | `0` | `7,810` |
| 6 | `record_typescript_lookup` | `728,192,400` | `7,810` |
| 7 | `typescript_lookup_record_residual` | `0` | `23,430` |

- Parent/subgroup/M1 equality: `764,972,500 = 764,972,500 = 764,972,500 ns`.
- Parent invocation equation: `23,430 = 3,692 + 4,118 + 7,810 + 7,810`.
- Site equation: `7,810 = 3,692 + 4,118 = targetTextCount = recordCount`.
- Prior-unclaimed equation: `3,692 = memberLookupCount 3,692`.
- Eligible/recheck equation: `3,549 = 3,549`.
- False/true equation: `3,549 + 0 = 3,549`; timer-control invocations are `3,549`.
- Subgroup sum and overlap: `764,972,500 ns / 0`. Every remainder is nonnegative.
- Net receiver-recheck duration: `10,466,100 - 1,000,600 = 9,465,500 ns`, positive.

### Cheapapp parity, resources, and artifacts

- Files `1,359 / 887 / 0`; parser bytes `5,430,581`; graph/DB `95,762 / 129,716`: PASS.
- Non-timing A003 comparison: `565 / 565` comparable leaves, zero differences. Resolution counters, diagnostics, outcomes, complete ordered Evidence, workload, file/parser, graph/DB semantics are exact after excluding only the authorized measurement/timing/resource fields.
- Copied canonical Graph JSON: `840,614,023` bytes; `8763AA2D089328F3D782A5F01E06178D171922E3A5CE96A4EE6707510976C920`.
- stdout/stderr: `8,824 / 213` bytes; `F52193AD16DD18C3DBD78CC161B503CF4091BDFAC1B363962FEBC2C50AD7FEF4` / `0FDCF15222C02C25771C98E17D9220C89FDCD0294137A6AE4D9B3AAFD4761ADC`.
- Resources: `startAllocBytes=1,325,984`; `endAllocBytes=779,787,072`; `maxObservedSys=2,021,476,856`.
- Benchmark/process/profile: `153,280 / 5,893 / 61,331` bytes; SHA-256 `EBD6C5F0DD8869E64B6A3A69E5E82D106DDFC026C9886BEE5941FE3A7139B6BD` / `B1FE1CAD96D2FAD8784984E2FF224185461716F9E57697E7C23D420C37E81C06` / `4C1A30068F88DB017E2E21C5B3EDA5F63292447DE91BCD8E546699530BD2327E`.

Cheapapp independently satisfies every exposure predicate.

## Restaurant Manager packet

### Parent and retained M1 conservation

| Gate | Exact value | Result |
|---|---:|---|
| Operations | `30 / 30`, ordered exact | PASS |
| Resolution children | `17 / 17`, ordered exact | PASS |
| Child intervals / overlap | `3,716 / 0` | PASS |
| Parent / child sum / residual | `18,778,365,600 / 18,758,990,600 / 19,375,000 ns` | PASS |
| D001 denominator | `calls=86,030; files=1,234` | PASS |
| D001 `resolve_calls` | `8,511,039,500 ns` | PASS |
| Analyzer / process | `100,274,133,800 / 116,614,449,900 ns` | attribution only |

| Order | M1 group | Duration (ns) | Invocations |
|---:|---|---:|---:|
| 1 | `source_context` | `138,497,200` | `86,030` |
| 2 | `binding_receiver` | `121,796,600` | `105,656` |
| 3 | `scoped_same_file` | `237,167,000` | `65,948` |
| 4 | `member_import` | `244,756,800` | `167,208` |
| 5 | `go_same_package` | `4,209,333,300` | `16,011` |
| 6 | `global_lookup` | `21,303,500` | `13,781` |
| 7 | `typescript_lookup_record` | `500,153,900` | `44,775` |
| 8 | `evidence_emission` | `2,878,459,900` | `164,333` |
| 9 | `direct_site_identity` | `92,659,300` | `156,976` |
| 10 | `resolve_call_residual` | `66,912,000` | `86,030` |

M1 sum is exactly `8,511,039,500 ns`, equal to D001; all durations/remainders are nonnegative and overlap is `0`.

### Seven-group split and work equations

| Order | M2 subgroup | Duration (ns) | Invocations |
|---:|---|---:|---:|
| 1 | `member_receiver_recheck` | `0` | `638` |
| 2 | `member_receiver_recheck_timer_control` | `0` | `638` |
| 3 | `member_lookup_remainder` | `438,200` | `1,144` |
| 4 | `global_lookup` | `19,925,000` | `13,781` |
| 5 | `site_target_text` | `0` | `14,925` |
| 6 | `record_typescript_lookup` | `479,691,300` | `14,925` |
| 7 | `typescript_lookup_record_residual` | `99,400` | `44,775` |

- Parent/subgroup/M1 equality: `500,153,900 = 500,153,900 = 500,153,900 ns`.
- Parent invocation equation: `44,775 = 1,144 + 13,781 + 14,925 + 14,925`.
- Site equation: `14,925 = 1,144 + 13,781 = targetTextCount = recordCount`.
- Prior-unclaimed equation: `1,144 = memberLookupCount 1,144`.
- Eligible/recheck equation: `638 = 638`.
- False/true equation: `638 + 0 = 638`; timer-control invocations are `638`.
- Subgroup sum and overlap: `500,153,900 ns / 0`. Every remainder is nonnegative.
- Net receiver-recheck duration: `0 - 0 = 0 ns`, not positive.

### Restaurant Manager parity, resources, and artifacts

- Files `1,556 / 1,234 / 0`; parser bytes `8,546,179`; graph/DB `137,229 / 190,644`: PASS.
- Non-timing A003 comparison: `565 / 565` comparable leaves, zero differences. Resolution counters, diagnostics, outcomes, complete ordered Evidence, workload, file/parser, graph/DB semantics are exact after excluding only the authorized measurement/timing/resource fields.
- Copied canonical Graph JSON: `900,212,685` bytes; `09708B038AD63B659A95BAC34304A752F2C6350952B91CF435EBD5C5F9D94B7C`.
- stdout/stderr: `8,944 / 213` bytes; `CD0028F3B3939451FF4D3623DF7DD7482A993E0934D2DFA773682D37D4FB0E94` / `0FDCF15222C02C25771C98E17D9220C89FDCD0294137A6AE4D9B3AAFD4761ADC`.
- Resources: `startAllocBytes=1,475,952`; `endAllocBytes=1,207,002,712`; `maxObservedSys=2,776,259,192`.
- Benchmark/process/profile: `200,715 / 4,390 / 54,519` bytes; SHA-256 `E8E17CE5D9FAB88CC8B9D5960532E7E1E1420BFC554C6F2F6C971EC57A63214E` / `D0661238448084CF3BC200601F90E33AD174C43242B7523F4C78B79A8552A99B` / `94FE7B1C35859B06AA7307E07095B0FDBC90C97CEF7F124C69884EC488833B68`.

Restaurant Manager independently passes every packet, count, conservation, parity, zero-true, and prior-unclaimed gate, but fails the required positive net-duration predicate.

## Validation artifacts and decision

| Validation | Bytes | SHA-256 | Status |
|---|---:|---|---|
| `validation/cheapapp.json` | `18,617` | `57F6FB4BD0B43AA809E6F50A3ED4BCFC42EB874A5E401E0E89FD52EAFDC06EE5` | PASS, zero failed gates |
| `validation/restaurant_manager.json` | `18,068` | `2E0E98125C0FB852FF2612584F08554E3E54D59850A572DF554EC60F1EF784AB` | PASS, zero failed gates |
| `validation/cross-target.json` | `2,220` | `E3412C9B8CDDDC4D380238B7BB73453D0423E7114D78C6F29D48783094CC70C0` | PASS, overlap `0` |

Both target packets are valid. Cheapapp has positive recheck duration beyond control; Restaurant Manager has `638` rechecks but records `0 ns` for both the predicate and matched control, so its required exclusive-duration-minus-control predicate is not positive. Under the Architect's exact decision rule, this falsifies the proposed two-target production owner. It does not authorize a production change or another attribution packet.

Checkpoint remains attribution-only with D001 streak `2`; parent/D001 are unchecked and downstream work remains closed. Next owner: Main Orchestration. No M3 or rerun is authorized.

`A006_M2_RECEIVER_RECHECK_FALSIFIED`
