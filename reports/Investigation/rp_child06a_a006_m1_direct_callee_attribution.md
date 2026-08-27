# Child 06A A006-M1 Direct-Callee Attribution

Status: measurement complete after one Main-authorized support recovery.

Measurement ID: `A006-M1-D001-DIRECT-CALLEE-ATTRIBUTION`

This is a measurement-only packet for `P2-A / A006-M1 / B1-P1A-OP001 resolution / B2-P2A-A001-D001 resolve_calls`. It is not a production attempt, candidate, baseline rerun, Supervisor event, disposition, or streak event. The initial input failure and the later Main-authorized measurement-support recovery are both preserved below.

## Starting gate

| Gate | Result | Evidence |
|---|---|---|
| Work root absent before creation | PASS | `E:\Anvien\.tmp\child06a_a006_m1_direct_callee_20260827` absent |
| Report absent before creation | PASS | This report path was absent |
| Repository HEAD | PASS | `3adc38109bf860c29ffb6484ea34e4ecc667157d` |
| Worktree / staged set | PASS | clean / empty |
| `anvien --help` | PASS | exit `0` |
| Architect impact authority | ACKNOWLEDGED | file `HIGH`; `resolveCall` `CRITICAL`; isolated overlay only |
| Build competitors | PASS | zero `go.exe`, `compile.exe`, `link.exe`, or other `build-a00x-benchmark.ps1` process |

No graph analyze, query, file-detail, impact, or detect command was run; the fresh owner/impact evidence was supplied by the A006 Architect and this lane changed only repo-local overlay copies.

## Seed, overlay, and manifest identity

The two seeds were copied mechanically with `Copy-Item -LiteralPath`; their originals were not changed. Instrumentation edits and manifest creation used `apply_patch`; `gofmt` touched only the two new overlay copies.

| File | Seed SHA-256 | Overlay SHA-256 | Diff |
|---|---|---|---|
| `resolve.go` | `92A89227E1B1B1C159DE8BE77A1F060361C9058C4F83C78BFC32E9AB40DADEA9` | `6B77B55097664234AB9FA28102C84702F3A8B830D36C7449AF2F676242DB8A61` | `184` insertions, `32` deletions; recorder, publication, and caller-site timing only |
| `types.go` | `A8DDCB1C61C312D3D3B164B1DD02BFD91C640F1FEB8B6DB511B73FD923C802B8` | `8F5F8E7B2B29FDDB7A65384970DAC60E7E15882608D9024FF6DF7414DD9F4E2D` | semantic diff is exactly one row type and one Metrics field; remaining churn is gofmt-only |

Canonical mapped-source start hashes passed: `resolve.go=8CEEDBA1883314EE8883320D3647C25DEF6F19F043D57881A893FBA73BA210D9`; `types.go=7C5E113F5E50584665D6D9AED0BCB2B3C6F6085219A62CD5AE74DD7654CD5DC3`.

Manifest: `E:\Anvien\.tmp\child06a_a006_m1_direct_callee_20260827\overlay\optimized-overlay.json`; SHA-256 `D9ED3E5B254CA31DD1CA20894A049F4E5B0329EB95644E0FD445E165AF68AB98`. It contains exactly one top-level `Replace` object and exactly these two mappings:

- `E:\Anvien\internal\resolution\resolve.go` -> `E:\Anvien\.tmp\child06a_a006_m1_direct_callee_20260827\overlay\internal\resolution\resolve.go`
- `E:\Anvien\internal\resolution\types.go` -> `E:\Anvien\.tmp\child06a_a006_m1_direct_callee_20260827\overlay\internal\resolution\types.go`

The measurement row exposes only `groupId`, `durationNs`, and `invocationCount`; Metrics exposes it only as `resolveCallDirectCalleeMeasurements`.

## Static group-coverage and non-overlap audit

The overlay owns one fixed `[10]` row array for the run. It has no per-call slice/map/event/log, global, goroutine, lock, I/O, finalizer, or TTL. All product callees are timed synchronously at their direct `resolveCall` caller sites. One group timer always ends before another begins; nested descendants are included in the outer direct-call elapsed time and are not timed again. The enclosing accepted D001 timer is not a second direct-callee group.

| Order | Group | Direct-site coverage |
|---:|---|---|
| 1 | `source_context` | `sourceForScopeOrFile` |
| 2 | `binding_receiver` | `repositoryReceiverClaimed`, `bindingOccurrences.resolve`, `emitBindingOccurrenceReference` |
| 3 | `scoped_same_file` | both `resolveScopedName` and both `resolveSameFileName` branches |
| 4 | `member_import` | `resolveMember`, `resolveImportedMemberWithProof`, `explicitImportNameClaimed`, all five `explicitImportCallState` sites |
| 5 | `go_same_package` | `resolveGoSamePackageFunction` |
| 6 | `global_lookup` | all three `resolveGlobalCallName` sites |
| 7 | `typescript_lookup_record` | exactly one lookup branch plus `callTargetText` argument evaluation and `recordTypeScriptLookup` in one non-overlapping interval |
| 8 | `evidence_emission` | both `appendExportBindingEvidence`, all four direct `emitUnresolvedReference`, `emitReference`, `recordRepositoryUnresolvedOutcome`, `retainedExportResolutionForScopedBinding`; emission argument identity evaluation stays inside this group |
| 9 | `direct_site_identity` | both standalone `callTargetText` + `sourceSiteID` evaluations not owned by groups 1-8 |
| 10 | `resolve_call_residual` | exact D001 child duration minus groups 1-9; invocation count increments once per `resolveCall` |

Label construction, branching, metric increments, recorder overhead, and other unlisted control work remain residual. Static coverage/non-overlap audit: PASS. Both runtime packets below prove exact conservation and zero overlap.

## Initial builder invocation — preserved failure evidence

Builder: `E:\Anvien\scripts\build-a00x-benchmark.ps1`; tracked diff empty; SHA-256 `ADA407C7496FCEA988276F03BAD5001ED139A4AEC9A16B9C32947DA440814EC5`.

The original executable, adjacent DLL, and provenance were absent immediately before invocation. At the initial checkpoint, the builder was invoked exactly once from `E:\Anvien` with:

```powershell
& 'E:\Anvien\scripts\build-a00x-benchmark.ps1' `
  -AttemptId 'A006-M1' `
  -OverlayManifestPath 'E:\Anvien\.tmp\child06a_a006_m1_direct_callee_20260827\overlay\optimized-overlay.json' `
  -OutputExecutablePath 'E:\Anvien\.tmp\child06a_a006_m1_direct_callee_20260827\frozen-build\anvien-a006-m1-benchmark.exe' `
  -ExpectedOverlaySha256 'D9ED3E5B254CA31DD1CA20894A049F4E5B0329EB95644E0FD445E165AF68AB98' `
  -ExpectedMappedSourceHash @(
    'E:\Anvien\internal\resolution\resolve.go=6B77B55097664234AB9FA28102C84702F3A8B830D36C7449AF2F676242DB8A61',
    'E:\Anvien\internal\resolution\types.go=8F5F8E7B2B29FDDB7A65384970DAC60E7E15882608D9024FF6DF7414DD9F4E2D'
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

Result: FAIL, exit `1`.

Exact failure:

```text
build-a00x-benchmark.ps1:
Line |
  17 |  & 'E:\Anvien\scripts\build-a00x-benchmark.ps1' `
     |  ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
     | Error formatting a string: Index (zero based) must be greater than or equal to zero and less than the size of the argument list..
```

Post-failure state at that checkpoint: the `frozen-build` directory, executable, adjacent DLL, and provenance JSON were all absent. Completion marker, executable version/hash, provenance schema/attempt/exits, and exact `2/4/3` inputs were therefore `NOT EXPOSED`. No retry, patch, rebuild, or cleanup occurred in that gate. Main later classified the blocker: the unchanged builder requires `AttemptId` matching `^A[0-9]{3}$`; `A006-M1` hit that guard, whose unescaped `{3}` error text surfaced as the format exception above.

## Main-authorized measurement-support recovery build

Main authorized one recovery within the same A006-M1 ownership. Overlay, manifest, script, production, tests, plans, and ledgers remained unchanged. The retry1 executable, adjacent DLL, and provenance were absent before the one recovery invocation. Only these two inputs changed from the preserved failed command:

- provenance `AttemptId`: `A006`;
- output: `E:\Anvien\.tmp\child06a_a006_m1_direct_callee_20260827\frozen-build-retry1\anvien-a006-m1-benchmark.exe`.

The unchanged builder was invoked once and returned exit `0` with marker `A00X_BENCHMARK_BUILD_COMPLETE`. No further build recovery was attempted.

| Recovery identity | Exact result |
|---|---|
| Repository HEAD | `046eda35f380ba43995c39ea579f5014b6e1cfac` |
| Executable | `73,836,032` bytes; version `1.2.8`; SHA-256 `0362FE211C2E072988EF61B662FC4F0D2160437F520F4525CDA59DDE52FB57A7` |
| Adjacent DLL | `20,230,656` bytes; SHA-256 `20CBD87840483A2053CFF3FC2DB23A86DD802B8915D86509D41A4B709624CDB7` |
| Provenance | `E:\Anvien\.tmp\child06a_a006_m1_direct_callee_20260827\frozen-build-retry1\anvien-a006-m1-benchmark.provenance.json`; `13,502` bytes; SHA-256 `80BCC23E20518815CB5B7A94C21B2524B7984F6B5C32475CFD45444AA746D887` |
| Schema / attempt / build exit / root exit | `1 / A006 / 0 / 0` |
| Overlay / candidate-source / native inputs | `2 / 4 / 3` |
| Overlay manifest | expected/actual `D9ED3E5B254CA31DD1CA20894A049F4E5B0329EB95644E0FD445E165AF68AB98` |
| Overlay replacements | expected/actual `6B77B55097664234AB9FA28102C84702F3A8B830D36C7449AF2F676242DB8A61`; `8F5F8E7B2B29FDDB7A65384970DAC60E7E15882608D9024FF6DF7414DD9F4E2D` |
| Candidate/native hash pairs | all four candidate-source and all three native expected/actual pairs exact |

Provenance build interval: `2026-08-26T23:45:08.1621643Z` to `2026-08-26T23:48:23.9659308Z`. The builder did not launch analyze.

## Sequential capture identity

Both targets ran with zero competitors and exactly one `Process.Start`. Cheapapp exited before Restaurant Manager preflight. A preliminary Cheapapp wrapper identity query encountered Git's dubious-ownership guard before `ProcessStartInfo` construction; it created no assigned output and made no target launch. The successful preflight used only command-local `git -c safe.directory=E:/cheapapp.org`; no Git configuration changed.

| Target | PID / exit | Start | End | Wall / CPU / user / kernel (s) |
|---|---|---|---|---|
| Cheapapp | `19288 / 0` | `2026-08-27T06:51:47.8464038+07:00` | `2026-08-27T06:53:53.5496640+07:00` | `125.7032602 / 104.0156250 / 87.9687500 / 16.0468750` |
| Restaurant Manager | `17244 / 0` | `2026-08-27T06:58:14.5609437+07:00` | `2026-08-27T07:00:35.3795083+07:00` | `140.8185646 / 110.1718750 / 93.0781250 / 17.0937500` |

Restaurant Manager started exactly `261,011,279,700 ns` after Cheapapp ended. Target capture overlap count: `0`. Both target HEADs and full statuses were identical before/after; repository HEAD stayed `046eda35f380ba43995c39ea579f5014b6e1cfac`; staged counts were `0 -> 0` for both runs.

## Cheapapp packet

Exact argv:

```text
E:\Anvien\.tmp\child06a_a006_m1_direct_callee_20260827\frozen-build-retry1\anvien-a006-m1-benchmark.exe analyze E:\cheapapp.org --force --skip-git --json --progress --benchmark-json E:\Anvien\.tmp\child06a_a006_m1_direct_callee_20260827\captures\cheapapp\candidate\benchmark.json --benchmark-label child06a-a006-m1-cheapapp-direct-callee-20260827
```

Working directory: `E:\cheapapp.org`. Process-local environment identity:

```text
HOME=...\captures\cheapapp\process-local\home
USERPROFILE=...\captures\cheapapp\process-local\home
TEMP=...\captures\cheapapp\process-local\temp
TMP=...\captures\cheapapp\process-local\tmp
APPDATA=...\captures\cheapapp\process-local\appdata
LOCALAPPDATA=...\captures\cheapapp\process-local\localappdata
XDG_CACHE_HOME=...\captures\cheapapp\process-local\xdg-cache
XDG_CONFIG_HOME=...\captures\cheapapp\process-local\xdg-config
XDG_DATA_HOME=...\captures\cheapapp\process-local\xdg-data
ANVIEN_OP001_RESOLUTION_CPU_PROFILE=...\captures\cheapapp\candidate\resolution.cpu.pprof
```

### Cheapapp conservation

| Gate | Exact result | Status |
|---|---:|---|
| Top-level operations / unique | `30 / 30` | PASS |
| Resolution children / unique | `17 / 17` | PASS |
| Exclusive intervals / overlap | `2,675 / 0` | PASS |
| Parent / child sum / residual | `19,018,779,500 / 19,001,125,700 / 17,653,800 ns` | PASS |
| D001 denominator | `calls=27,890; files=887` | PASS |
| D001 `resolve_calls` | `3,048,999,300 ns` | PASS |
| Analyzer / process | `96,633,102,400 / 125,703,260,200 ns` | attribution only |

Every child duration equals its interval sum; child sum plus residual equals the parent exactly.

### Cheapapp ten direct-callee groups

| Order | Group | Duration (ns) | Invocation count |
|---:|---|---:|---:|
| 1 | `source_context` | `43,637,500` | `27,890` |
| 2 | `binding_receiver` | `78,596,800` | `26,079` |
| 3 | `scoped_same_file` | `53,094,400` | `29,686` |
| 4 | `member_import` | `186,935,000` | `47,123` |
| 5 | `go_same_package` | `972,300` | `8,501` |
| 6 | `global_lookup` | `14,767,200` | `4,118` |
| 7 | `typescript_lookup_record` | `686,469,300` | `23,430` |
| 8 | `evidence_emission` | `1,954,984,600` | `44,488` |
| 9 | `direct_site_identity` | `9,858,900` | `40,360` |
| 10 | `resolve_call_residual` | `19,683,300` | `27,890` |

Direct-callee sum: `3,048,999,300 ns`, exactly equal to D001. Residual is nonnegative, every invocation count is explicit, and direct-group overlap count is `0`: PASS.

### Cheapapp parity and resources

| Gate | Exact result | Status |
|---|---|---|
| Files scanned / parsed / failed | `1,359 / 887 / 0` | PASS |
| Parser bytes | `5,430,581` | PASS |
| Graph / DB rows | `95,762 / 129,716` | PASS |
| Non-timing benchmark semantics | `565 / 565` comparable A003 leaves; `0` differences | PASS |
| Resolution counters, diagnostics, outcomes | exact after excluding timers and the authorized measurement field | PASS |
| Ordered Evidence and canonical Graph JSON | `840,614,023` bytes; `8763AA2D089328F3D782A5F01E06178D171922E3A5CE96A4EE6707510976C920` | PASS |
| stdout | `8,824` bytes; `F52193AD16DD18C3DBD78CC161B503CF4091BDFAC1B363962FEBC2C50AD7FEF4` | PASS |
| stderr | `213` bytes; `0FDCF15222C02C25771C98E17D9220C89FDCD0294137A6AE4D9B3AAFD4761ADC` | PASS |
| Ladybug / meta | `597,946,368` bytes / `F85563604F687D50104F67E6A3FF30B11ECA833D2D376DA2F8C3FB0005D73105`; `219` bytes / `FBCB3E554AF37EA9F6467A28FAECFAB271F1C14C78523A38674FE4F2CC74DF95` | recorded; raw bytes non-governing |

Resource field set exactly matches A003 and values are recorded honestly:

| Field | A003 | A006-M1 | Delta |
|---|---:|---:|---:|
| `startAllocBytes` | `1,315,192` | `1,316,960` | `+1,768` |
| `endAllocBytes` | `886,583,856` | `1,138,887,304` | `+252,303,448` |
| `maxObservedSys` | `1,995,295,224` | `2,022,459,896` | `+27,164,672` |

Artifacts: benchmark `150,626` bytes / `98836884377FA3AF8FC8380076A06B7A7BD559525531CFB314A0055ADA553D0F`; process `5,168` bytes / `0F79229A8AC2B193FA55B4A6788EF9BA1BB9C1CF742134ADE949C4879B6BC10C`; profile `61,432` bytes / `28FB0A69E96385ED2FC42E0E5A33CA1C19307BBACA2B2AC704FBDC871F1310A8`.

## Restaurant Manager packet

Exact argv:

```text
E:\Anvien\.tmp\child06a_a006_m1_direct_callee_20260827\frozen-build-retry1\anvien-a006-m1-benchmark.exe analyze E:\Restaurant_manager --force --json --progress --exclude electron/renderer/src/api/userApi.ts --benchmark-json E:\Anvien\.tmp\child06a_a006_m1_direct_callee_20260827\captures\restaurant_manager\candidate\benchmark.json --benchmark-label child06a-a006-m1-restaurant-manager-direct-callee-20260827
```

Working directory: `E:\Anvien`. Process-local environment identity:

```text
TEMP=...\captures\restaurant_manager\process-local\temp
TMP=...\captures\restaurant_manager\process-local\tmp
TMPDIR=...\captures\restaurant_manager\process-local\tmp
HOME=...\captures\restaurant_manager\process-local\home
USERPROFILE=...\captures\restaurant_manager\process-local\home
APPDATA=...\process-local\home\AppData\Roaming
LOCALAPPDATA=...\process-local\home\AppData\Local
XDG_CACHE_HOME=...\process-local\home\.cache
XDG_CONFIG_HOME=...\process-local\home\.config
XDG_DATA_HOME=...\process-local\home\.local\share
GOCACHE=...\captures\restaurant_manager\process-local\gocache
GOMODCACHE=...\captures\restaurant_manager\process-local\gomodcache
GOTMPDIR=...\captures\restaurant_manager\process-local\gotmp
ANVIEN_OP001_RESOLUTION_CPU_PROFILE=...\captures\restaurant_manager\candidate\resolution.cpu.pprof
GIT_CONFIG_COUNT=1; GIT_CONFIG_KEY_0=safe.directory; GIT_CONFIG_VALUE_0=E:/Restaurant_manager
```

### Restaurant Manager conservation

| Gate | Exact result | Status |
|---|---:|---|
| Top-level operations / unique | `30 / 30` | PASS |
| Resolution children / unique | `17 / 17` | PASS |
| Exclusive intervals / overlap | `3,716 / 0` | PASS |
| Parent / child sum / residual | `18,605,489,700 / 18,582,446,900 / 23,042,800 ns` | PASS |
| D001 denominator | `calls=86,030; files=1,234` | PASS |
| D001 `resolve_calls` | `8,297,125,100 ns` | PASS |
| Analyzer / process | `111,172,345,800 / 140,818,564,600 ns` | attribution only |

Every child duration equals its interval sum; child sum plus residual equals the parent exactly.

### Restaurant Manager ten direct-callee groups

| Order | Group | Duration (ns) | Invocation count |
|---:|---|---:|---:|
| 1 | `source_context` | `204,263,300` | `86,030` |
| 2 | `binding_receiver` | `185,310,600` | `105,656` |
| 3 | `scoped_same_file` | `266,276,000` | `65,948` |
| 4 | `member_import` | `308,333,200` | `167,208` |
| 5 | `go_same_package` | `3,946,043,200` | `16,011` |
| 6 | `global_lookup` | `11,392,500` | `13,781` |
| 7 | `typescript_lookup_record` | `435,161,600` | `44,775` |
| 8 | `evidence_emission` | `2,774,855,300` | `164,333` |
| 9 | `direct_site_identity` | `120,948,600` | `156,976` |
| 10 | `resolve_call_residual` | `44,540,800` | `86,030` |

Direct-callee sum: `8,297,125,100 ns`, exactly equal to D001. Residual is nonnegative, every invocation count is explicit, and direct-group overlap count is `0`: PASS.

### Restaurant Manager parity and resources

| Gate | Exact result | Status |
|---|---|---|
| Files scanned / parsed / failed | `1,556 / 1,234 / 0` | PASS |
| Parser bytes | `8,546,179` | PASS |
| Graph / DB rows | `137,229 / 190,644` | PASS |
| Non-timing benchmark semantics | `565 / 565` comparable A003 leaves; `0` differences | PASS |
| Resolution counters, diagnostics, outcomes | exact after excluding timers and the authorized measurement field | PASS |
| Ordered Evidence and canonical Graph JSON | `900,212,685` bytes; `09708B038AD63B659A95BAC34304A752F2C6350952B91CF435EBD5C5F9D94B7C` | PASS |
| stdout | `8,944` bytes; `CD0028F3B3939451FF4D3623DF7DD7482A993E0934D2DFA773682D37D4FB0E94` | PASS |
| stderr | `213` bytes; `0FDCF15222C02C25771C98E17D9220C89FDCD0294137A6AE4D9B3AAFD4761ADC` | PASS |
| Ladybug / meta | `503,070,720` bytes / `E9C4B85C0FB743564A84ED4709FDDA81CC7C32A597C4833F938A79DCE156CE57`; `267` bytes / `ADA05DD404B7F9525FE99820DF7F0D8EB9252B2F1D7794F6652E0463768DAD7D` | recorded; raw bytes non-governing |

Resource field set exactly matches A003 and values are recorded honestly:

| Field | A003 | A006-M1 | Delta |
|---|---:|---:|---:|
| `startAllocBytes` | `1,474,592` | `1,470,696` | `-3,896` |
| `endAllocBytes` | `1,306,049,888` | `1,269,417,312` | `-36,632,576` |
| `maxObservedSys` | `2,634,607,224` | `2,998,655,608` | `+364,048,384` |

Artifacts: benchmark `197,487` bytes / `6918E3AA395B872E585496A81B9DDF993D60E77DB131B8D340C378F3A9A82E88`; process `4,125` bytes / `0F6A64A470547FDA4A22876A04C4C4ECE613803485FEDF297EE7BA59D5C11263`; profile `56,558` bytes / `D8887426FCB8226E367E5CE94A9EAD3E9A996103C92F953688E64A0FC4FC7AB7`.

## Limitation, checkpoint, and handoff

Instrumentation overhead is recorded honestly. All instrumented D001, parent, analyzer, process, and resource values are attribution evidence only; none is compared or promoted as candidate performance against accepted A003.

All mandatory build, sequential launch, denominator, conservation, overlap, output, graph/DB, semantic, ordering, and packet-presence gates pass. The initial failed-input evidence remains preserved above and is not a streak event.

Checkpoint: `A006_ARCHITECT_NEEDS_MEASUREMENT_INPUT / A006_M1_READY / D001_STREAK_2`. Parent/D001 remain unchecked; D002-D017, P3, and Child 07 remain closed. No plan, ledger, production, test, script, target-source, stage, commit, disposition, or Supervisor mutation occurred.

Next owner: Main Orchestration, to record this packet and reopen one fresh A006 Architect. No further measurement rerun is authorized by this report.

`A006_M1_DIRECT_CALLEE_ATTRIBUTION_READY`
