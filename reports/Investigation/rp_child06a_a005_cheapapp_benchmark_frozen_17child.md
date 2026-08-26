# Child 06A A005 Cheapapp Frozen 17-Child Benchmark

Status: `A005_CHEAPAPP_MEASUREMENT_PACKET_READY`.

This report preserves the original failed-input measurement-support event and records the one Main-authorized corrected `retry1` recovery launch. The failed root remains ineligible and was not rerun. The corrected root contains the sole eligible A005 Cheapapp candidate packet. This recovery is not a new architecture/implementation attempt, creates no D001 streak change, and makes no disposition, ledger, Supervisor, staging, or commit claim.

## Failed-input measurement-support event — preserved

The frozen identity and launch preconditions passed, and the candidate was started exactly once. It exited `1` after `33.387055100 s` without producing `candidate\benchmark.json`.

Exact blocker:

```text
resolution phase: ANVIEN_OP001_RESOLUTION_CPU_PROFILE must name the resolution CPU profile artifact
```

The authorized invocation did not include an `ANVIEN_OP001_RESOLUTION_CPU_PROFILE` environment assignment. The failure occurred after the emitted `resolution done` progress line. Per the execution contract, this lane stopped without retry or repair.

## Frozen identity preflight

| Check | Exact observed value | Result |
|---|---|---|
| Executable | `E:\Anvien\.tmp\child06a_a005_a00x_benchmark_build_retry1\anvien-a005-benchmark.exe`; `73,825,792` bytes; SHA-256 `0178DAC396E0DD4D0BECCE7B803378A3400013F0B508EC95F9221A0020CF96F3` | `PASS` |
| Executable version | provenance output version `1.2.8` | `PASS` |
| Adjacent DLL | `E:\Anvien\.tmp\child06a_a005_a00x_benchmark_build_retry1\lbug_shared.dll`; `20,230,656` bytes; SHA-256 `20CBD87840483A2053CFF3FC2DB23A86DD802B8915D86509D41A4B709624CDB7` | `PASS` |
| Provenance | `E:\Anvien\.tmp\child06a_a005_a00x_benchmark_build_retry1\anvien-a005-benchmark.provenance.json`; `13,322` bytes; SHA-256 `E4B9C7E9F9850A58356D41131C7CC177CFC6207D7D8CD3E158B633524B56244E` | `PASS` |
| Provenance schema / attempt / build exit / top-level exit | `1 / A005 / 0 / 0` | `PASS` |
| Provenance mapping / candidate / native counts | `2 / 4 / 3` | `PASS` |
| Provenance expected/actual pairs | all overlay replacement, candidate source, and native input SHA-256 pairs equal | `PASS` |
| Target | `E:\cheapapp.org`; HEAD `a869876ab6262dacde6cd5d432d099a91852a646`; complete ordered porcelain status equal to accepted A003 preflight | `PASS` |
| A003 accepted basis | Cheapapp packet `E:\Anvien\.tmp\child06a_a003_cheapapp_frozen_17child_20260826`; D001/parent/analyzer/process `3.447846300 / 20.472602300 / 93.531974900 / 95.630648200 s`; `calls=27890; files=887` | `PASS` |
| Pre-existing candidate output/report | output root absent; assigned report absent | `PASS` |
| Staged set before launch | empty | `PASS` |

## Exact launch and process identity

| Field | Exact value |
|---|---|
| Launch count | `1` |
| Process ID | `3440` |
| Executable | `E:\Anvien\.tmp\child06a_a005_a00x_benchmark_build_retry1\anvien-a005-benchmark.exe` |
| Working directory | `E:\cheapapp.org` |
| Target | `E:\cheapapp.org` |
| UTC start | `2026-08-26T21:34:53.1119547+00:00` |
| UTC end | `2026-08-26T21:35:26.4990098+00:00` |
| Exit code | `1` |
| Process wall | `33.387055100 s` |
| Total / user / kernel CPU | `16.156250000 / 13.687500000 / 2.468750000 s` |
| ProcessStartInfo | `FileName` exact frozen executable; `WorkingDirectory=E:\cheapapp.org`; `UseShellExecute=false`; `CreateNoWindow=true`; stdout/stderr redirected; every argument added with `ArgumentList.Add(...)` in the order below |

Exact ordered arguments:

1. `analyze`
2. `E:\cheapapp.org`
3. `--force`
4. `--skip-git`
5. `--json`
6. `--progress`
7. `--benchmark-json`
8. `E:\Anvien\.tmp\child06a_a005_cheapapp_frozen_17child_20260827\candidate\benchmark.json`
9. `--benchmark-label`
10. `child06a-a005-cheapapp-frozen-17child-20260827`

## Captured run artifacts

| Artifact | Bytes | SHA-256 | Result |
|---|---:|---|---|
| `E:\Anvien\.tmp\child06a_a005_cheapapp_frozen_17child_20260827\candidate\benchmark.json` | absent | n/a | `FAIL` |
| `E:\Anvien\.tmp\child06a_a005_cheapapp_frozen_17child_20260827\candidate\stdout.txt` | `0` | `E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855` | `PASS` as truthful capture; `FAIL` for A003 stdout parity |
| `E:\Anvien\.tmp\child06a_a005_cheapapp_frozen_17child_20260827\candidate\stderr.txt` | `234` | `AC4E9403D3ADA081B035B331FBFDD06158ECC8D043ECC271DDA641B938A3E164` | `PASS` as truthful capture; `FAIL` for A003 stderr semantics/parity |

Exact captured stderr:

```text
scan done
structure done
documents done
cobol done
parse done
routes done
tools done
orm done
cross_file_binding done
resolution done
resolution phase: ANVIEN_OP001_RESOLUTION_CPU_PROFILE must name the resolution CPU profile artifact
```

Captured stdout is empty.

## Measurement and validation matrix

`NOT EXPOSED` means the failed run produced no benchmark packet or eligible completed output from which that check could be established. Missing data is not inferred.

| Required check | Candidate fact / A003 comparison | Result |
|---|---|---|
| Command/runtime/target/output/process identity | exact values above; launch count `1` | `PASS` |
| Successful candidate exit | exit `1` | `FAIL` |
| Candidate D001 `resolve_calls` | no benchmark artifact | `NOT EXPOSED` |
| D001 delta vs A003 `3.447846300 s` | no candidate D001 value | `NOT EXPOSED` |
| Candidate parent `resolution` | no benchmark artifact | `NOT EXPOSED` |
| Parent delta vs A003 `20.472602300 s` | no candidate parent value | `NOT EXPOSED` |
| Candidate analyzer total | no benchmark artifact | `NOT EXPOSED` |
| Analyzer delta vs A003 `93.531974900 s` | no candidate analyzer value | `NOT EXPOSED` |
| Process wall | observed failed-run wall `33.387055100 s`; arithmetic delta vs A003 `-62.243593100 s` (`-65.087495%`) | `PASS` as process observation; `FAIL` as comparable performance evidence because exit was nonzero |
| Top-level operation inventory | no benchmark artifact; cannot establish `30/30` | `NOT EXPOSED` |
| Resolution child inventory | no benchmark artifact; cannot establish `17/17` | `NOT EXPOSED` |
| D001 denominator | accepted A003 is `calls=27890; files=887`; candidate denominator unavailable | `NOT EXPOSED` |
| Child interval conservation | no child intervals available | `NOT EXPOSED` |
| Zero interval overlap | no child intervals available | `NOT EXPOSED` |
| Complete ordered evidence array | no eligible candidate result/evidence array | `NOT EXPOSED` |
| Canonical Graph JSON SHA/equality | stop contract forbids continuing output validation after nonzero exit; no candidate gate established | `NOT EXPOSED` |
| Canonical Graph JSON semantic equality | no eligible completed candidate output | `NOT EXPOSED` |
| Public stdout semantics vs A003 | empty candidate stdout differs from accepted A003 `8,824` bytes / SHA-256 `F52193AD16DD18C3DBD78CC161B503CF4091BDFAC1B363962FEBC2C50AD7FEF4` | `FAIL` |
| Public stderr semantics vs A003 | candidate contains the fatal profile-path error and differs from accepted A003 `213` bytes / SHA-256 `0FDCF15222C02C25771C98E17D9220C89FDCD0294137A6AE4D9B3AAFD4761ADC` | `FAIL` |
| Graph/DB readback | not performed after the mandatory nonzero-exit stop | `NOT EXPOSED` |
| Resolution outcomes/diagnostics/counters | no eligible completed packet | `NOT EXPOSED` |
| Files scanned/parsed/failed | accepted A003 `1,359 / 887 / 0`; candidate completed values unavailable | `NOT EXPOSED` |
| Parser bytes | accepted A003 `5,430,581`; candidate completed value unavailable | `NOT EXPOSED` |
| Graph nodes/relationships | accepted A003 `95,762 / 129,716`; candidate completed values unavailable | `NOT EXPOSED` |
| `startAllocBytes` | no benchmark artifact | `NOT EXPOSED` |
| `endAllocBytes` | no benchmark artifact | `NOT EXPOSED` |
| `maxObservedSys` | no benchmark artifact | `NOT EXPOSED` |
| Private run-scoped `O(U+B)` one-canonical-payload lifecycle | not exposed by an eligible completed runtime packet | `NOT EXPOSED` |

## Failed-input event boundary

- Failed-root launch count remains exactly `1`; that root was not rerun.
- No target/index cleanup, repair, rebuild, process killing, A003 rerun, A004 audit, graph operation, detect, ledger update, source/test/script/plan mutation, staging, commit, disposition, Supervisor action, D002-D017/P3/Child 07 action, or Restaurant classification occurred.
- The concurrently running Restaurant measurement was neither waited for nor treated as a conflict.
- The failed-root packet is not performance evidence and cannot support `KEEP` or `NO_KEEP`.

## Failed-event report-only repository validation

| Check | Result |
|---|---|
| `git diff --check -- reports/Investigation/rp_child06a_a005_cheapapp_benchmark_frozen_17child.md` | `PASS` (exit `0`, no output) |
| Staged set remains empty | `PASS` (`0` paths) |

## Corrected `retry1` result

`A005_CHEAPAPP_MEASUREMENT_PACKET_READY`

The one corrected launch completed with exit `0`. It produced the exact five assigned candidate artifacts, a nonempty resolution CPU profile, `30/30` top-level operations, `17/17` resolution children, the unchanged D001 denominator `calls=27890; files=887`, exact interval conservation, zero overlap, and exact same-work/output parity against accepted A003.

This section reports measurements only. In particular, the higher analyzer/process totals are retained as facts without a `KEEP`, `NO_KEEP`, acceptance, rejection, or streak decision.

## Corrected frozen identity and launch contract

| Check | Exact value | Result |
|---|---|---|
| Executable | `E:\Anvien\.tmp\child06a_a005_a00x_benchmark_build_retry1\anvien-a005-benchmark.exe`; `73,825,792` bytes; SHA-256 `0178DAC396E0DD4D0BECCE7B803378A3400013F0B508EC95F9221A0020CF96F3`; version `1.2.8` | `PASS` |
| Adjacent DLL | `20,230,656` bytes; SHA-256 `20CBD87840483A2053CFF3FC2DB23A86DD802B8915D86509D41A4B709624CDB7` | `PASS` |
| Provenance | `13,322` bytes; SHA-256 `E4B9C7E9F9850A58356D41131C7CC177CFC6207D7D8CD3E158B633524B56244E`; schema/attempt/exits `1/A005/0/0`; mappings/candidate/native `2/4/3`; every expected/actual pair equal | `PASS` |
| Corrected output root | `E:\Anvien\.tmp\child06a_a005_cheapapp_frozen_17child_20260827_retry1`; absent before recovery | `PASS` |
| Working directory | `E:\Anvien` | `PASS` |
| Launch count | exactly `1` for corrected root; no further retry | `PASS` |
| ProcessStartInfo | direct start; `UseShellExecute=false`; `CreateNoWindow=true`; stdout/stderr redirected; exact ordered `ArgumentList` | `PASS` |
| Child environment | exactly the ten assigned overrides recorded in `process.json` | `PASS` |
| Target identity | `E:\cheapapp.org`; HEAD `a869876ab6262dacde6cd5d432d099a91852a646`; full ordered porcelain status equal to accepted A003 | `PASS` |
| Candidate label | `child06a-a005-cheapapp-frozen-17child-20260827-retry1` | `PASS` |

Exact ordered arguments:

1. `analyze`
2. `E:\cheapapp.org`
3. `--force`
4. `--skip-git`
5. `--json`
6. `--progress`
7. `--benchmark-json`
8. `E:\Anvien\.tmp\child06a_a005_cheapapp_frozen_17child_20260827_retry1\candidate\benchmark.json`
9. `--benchmark-label`
10. `child06a-a005-cheapapp-frozen-17child-20260827-retry1`

Exact child environment overrides:

```text
HOME=E:\Anvien\.tmp\child06a_a005_cheapapp_frozen_17child_20260827_retry1\home
USERPROFILE=E:\Anvien\.tmp\child06a_a005_cheapapp_frozen_17child_20260827_retry1\home
TEMP=E:\Anvien\.tmp\child06a_a005_cheapapp_frozen_17child_20260827_retry1\temp
TMP=E:\Anvien\.tmp\child06a_a005_cheapapp_frozen_17child_20260827_retry1\temp
APPDATA=E:\Anvien\.tmp\child06a_a005_cheapapp_frozen_17child_20260827_retry1\appdata
LOCALAPPDATA=E:\Anvien\.tmp\child06a_a005_cheapapp_frozen_17child_20260827_retry1\localappdata
XDG_CACHE_HOME=E:\Anvien\.tmp\child06a_a005_cheapapp_frozen_17child_20260827_retry1\xdg-cache
XDG_CONFIG_HOME=E:\Anvien\.tmp\child06a_a005_cheapapp_frozen_17child_20260827_retry1\xdg-config
XDG_DATA_HOME=E:\Anvien\.tmp\child06a_a005_cheapapp_frozen_17child_20260827_retry1\xdg-data
ANVIEN_OP001_RESOLUTION_CPU_PROFILE=E:\Anvien\.tmp\child06a_a005_cheapapp_frozen_17child_20260827_retry1\candidate\resolution.cpu.pprof
```

## Corrected process identity

| Field | Exact value |
|---|---|
| PID | `17616` |
| UTC start | `2026-08-26T21:41:20.9473060+00:00` |
| UTC end | `2026-08-26T21:44:44.5461379+00:00` |
| Exit | `0` |
| Process wall | `203.598831900 s` |
| Total CPU | `103.734375000 s` |
| User CPU | `87.781250000 s` |
| Kernel CPU | `15.953125000 s` |

## Primary A003-to-A005 measurements

| Boundary | Accepted A003 | Corrected A005 | A005 - A003 | Result |
|---|---:|---:|---:|---|
| D001 `resolve_calls` | `3.447846300 s` | `3.036901000 s` | `-0.410945300 s` (`-11.918898%`) | measured |
| B1-P1A-OP001 `resolution` parent | `20.472602300 s` | `18.160962900 s` | `-2.311639400 s` (`-11.291380%`) | measured |
| Analyzer total | `93.531974900 s` | `95.376559900 s` | `+1.844585000 s` (`+1.972144%`) | measured |
| Process wall | `95.630648200 s` | `203.598831900 s` | `+107.968183700 s` (`+112.901236%`) | measured; no disposition |

Cheapapp remains independent from Restaurant. No averaging or cross-target combination was performed.

## Complete `30/30` top-level operation inventory

Operation names, order, and every denominator object exactly equal accepted A003.

| # | Operation | A003 s | A005 s | Delta s | Denominator parity |
|---:|---|---:|---:|---:|---|
| 1 | `cli_startup` | `0.017113300` | `0.016998500` | `-0.000114800` | `PASS` |
| 2 | `cli_preparation` | `0.050773100` | `0.045623900` | `-0.005149200` | `PASS` |
| 3 | `analyze_setup` | `0.228877000` | `0.200606000` | `-0.028271000` | `PASS` |
| 4 | `scan` | `0.329323800` | `0.539106100` | `+0.209782300` | `PASS` |
| 5 | `structure` | `0.003751900` | `0.006965300` | `+0.003213400` | `PASS` |
| 6 | `documents` | `0.115228800` | `0.098979600` | `-0.016249200` | `PASS` |
| 7 | `cobol` | `0.002517600` | `0.000000000` | `-0.002517600` | `PASS` |
| 8 | `parse` | `10.798086200` | `9.371085800` | `-1.427000400` | `PASS` |
| 9 | `routes` | `1.471013400` | `1.301305600` | `-0.169707800` | `PASS` |
| 10 | `tools` | `0.460391300` | `0.421589800` | `-0.038801500` | `PASS` |
| 11 | `orm` | `0.088622300` | `0.080896300` | `-0.007726000` | `PASS` |
| 12 | `cross_file_binding` | `0.596224000` | `0.566897800` | `-0.029326200` | `PASS` |
| 13 | `resolution` | `20.472602300` | `18.160962900` | `-2.311639400` | `PASS` |
| 14 | `mro` | `0.003037800` | `0.000000000` | `-0.003037800` | `PASS` |
| 15 | `communities` | `0.039984300` | `0.031086400` | `-0.008897900` | `PASS` |
| 16 | `processes` | `0.576625900` | `0.596409600` | `+0.019783700` | `PASS` |
| 17 | `semantic_enrichment` | `3.516770500` | `3.879986000` | `+0.363215500` | `PASS` |
| 18 | `graph_compact` | `0.011348800` | `0.005543100` | `-0.005805700` | `PASS` |
| 19 | `db_runner_resolve` | `1.005832600` | `2.764960100` | `+1.759127500` | `PASS` |
| 20 | `db_load` | `39.490259400` | `42.070223400` | `+2.579964000` | `PASS` |
| 21 | `db_runner_close` | `0.202554800` | `0.261822800` | `+0.059268000` | `PASS` |
| 22 | `graph_snapshot` | `14.158001400` | `15.049875800` | `+0.891874400` | `PASS` |
| 23 | `benchmark_write` | `0.007725800` | `0.002507700` | `-0.005218100` | `PASS` |
| 24 | `analyzer_orchestration` | `0.182072000` | `0.166355800` | `-0.015716200` | `PASS` |
| 25 | `memory_profile` | `0.000000000` | `0.000000000` | `0.000000000` | `PASS` |
| 26 | `registry_meta` | `0.098595900` | `0.106917600` | `+0.008321700` | `PASS` |
| 27 | `ai_context` | `1.241777700` | `107.295071200` | `+106.053293500` | `PASS` |
| 28 | `file_projection` | `0.247998400` | `0.241303100` | `-0.006695300` | `PASS` |
| 29 | `output_publication` | `0.000000000` | `0.001074100` | `+0.001074100` | `PASS` |
| 30 | `cpu_profile_completion` | `0.000000000` | `0.000000000` | `0.000000000` | `PASS` |

## Complete `17/17` resolution-child inventory

The ordered child IDs, schemas, denominator expressions and objects, source owners/ranges, call paths, and overlap groups exactly equal accepted A003. D001 remains `calls=27890; files=887`.

| # | Resolution child | A003 s | A005 s | Delta s | Intervals | Conservation |
|---:|---|---:|---:|---:|---:|---|
| 1 | `build_binding_occurrence_index` | `0.077439500` | `0.065904600` | `-0.011534900` | `1` | `PASS` |
| 2 | `runtime_setup` | `0.000000000` | `0.000023600` | `+0.000023600` | `1` | `PASS` |
| 3 | `emit_definition_nodes` | `0.389693800` | `0.356197800` | `-0.033496000` | `1` | `PASS` |
| 4 | `emit_unresolved_heritage_diagnostics` | `0.000000000` | `0.000000000` | `0.000000000` | `1` | `PASS` |
| 5 | `emit_import_edges` | `0.037951300` | `0.035539500` | `-0.002411800` | `1` | `PASS` |
| 6 | `emit_heritage_compatibility_edges` | `0.000000000` | `0.000000000` | `0.000000000` | `1` | `PASS` |
| 7 | `resolve_calls` | `3.447846300` | `3.036901000` | `-0.410945300` | `887` | `PASS` |
| 8 | `resolve_accesses` | `9.380783200` | `8.524050200` | `-0.856733000` | `887` | `PASS` |
| 9 | `resolve_type_annotations` | `2.262894300` | `1.829571700` | `-0.433322600` | `887` | `PASS` |
| 10 | `emit_method_dispatch_edges` | `0.000000000` | `0.000000000` | `0.000000000` | `1` | `PASS` |
| 11 | `finalize_typescript_authority_results` | `0.028755700` | `0.030978500` | `+0.002222800` | `1` | `PASS` |
| 12 | `emit_typescript_external_symbols` | `0.015196500` | `0.017271300` | `+0.002074800` | `1` | `PASS` |
| 13 | `finalize_resolution_outcomes` | `0.156782300` | `0.215744100` | `+0.058961800` | `1` | `PASS` |
| 14 | `project_resolution_outcomes` | `4.652523600` | `4.023686900` | `-0.628836700` | `1` | `PASS` |
| 15 | `finalize_resolution_metadata` | `0.000000000` | `0.000000000` | `0.000000000` | `1` | `PASS` |
| 16 | `assemble_resolution_result` | `0.000000000` | `0.000000000` | `0.000000000` | `1` | `PASS` |
| 17 | `binding_accumulator_dispose` | `0.000000000` | `0.000000000` | `0.000000000` | `1` | `PASS` |

## Interval conservation and overlap

| Check | Exact candidate fact | Result |
|---|---:|---|
| Resolution parent | `18.160962900 s` | measured |
| Sum of 17 children | `18.135869200 s` | measured |
| Parent residual | `0.025093700 s` | `PASS` nonnegative |
| Exclusive intervals | `2,675` | `PASS` |
| Per-child duration equals interval-end minus interval-start sum | `17/17` | `PASS` |
| Parent equals child sum plus residual | exact integer-nanosecond equality | `PASS` |
| Positive-width interval overlaps | `0` | `PASS` |

## Same-work, output, DB, and semantic parity

The normalized benchmark comparison removes only labels, elapsed fields, interval timestamps, and memory observations. It retains phase order; all operation names/denominators; every non-timing subsystem payload; parser work counters; every resolution semantic counter; and every child schema, denominator, source owner/range, call path, and overlap group. The normalized objects are exact JSON-equal.

| Check | Accepted A003 | Corrected A005 | Result |
|---|---:|---:|---|
| Normalized same-work payload | accepted normalized object | exact JSON equality | `PASS` |
| Resolution semantic counters | accepted counter object | exact JSON equality | `PASS` |
| Files scanned / parsed / failed | `1,359 / 887 / 0` | `1,359 / 887 / 0` | `PASS` |
| Parser bytes | `5,430,581` | `5,430,581` | `PASS` |
| Graph nodes / relationships | `95,762 / 129,716` | `95,762 / 129,716` | `PASS` |
| DB node / relationship rows | `95,762 / 129,716` | `95,762 / 129,716` | `PASS` |
| Dependency edges / projection unresolved | `13,360 / 735` | `13,360 / 735` | `PASS` |
| `meta.json` semantic stats | files/nodes/edges/communities/processes `1,359/95,762/129,716/804/700` | `1,359/95,762/129,716/804/700` | `PASS` |
| Canonical Graph JSON | `840,614,023` bytes; `8763AA2D089328F3D782A5F01E06178D171922E3A5CE96A4EE6707510976C920` | same bytes and SHA-256 | `PASS` exact byte and semantic equality |
| Complete ordered evidence arrays | serialized inside the accepted canonical Graph JSON | full Graph JSON byte equality preserves every element and its order | `PASS` |
| Public stdout | `8,824` bytes; `F52193AD16DD18C3DBD78CC161B503CF4091BDFAC1B363962FEBC2C50AD7FEF4` | same bytes and SHA-256 | `PASS` |
| Public stderr | `213` bytes; `0FDCF15222C02C25771C98E17D9220C89FDCD0294137A6AE4D9B3AAFD4761ADC` | same bytes and SHA-256 | `PASS` |
| Resolution outcome/diagnostic/carrier semantics | accepted canonical graph and normalized counter objects | exact graph bytes plus normalized counter equality | `PASS` |
| Ladybug raw storage | A003 `DAB5EA00BA28728E23340CAE3BAA470DBC1B6937972B14C6686FC3E281374885` | A005 `4FC410680111475FDD0B65EE63AFB582B06BEDDE520D795CCD41A7B86472E572` | recorded; raw-byte equality is not the semantic gate |
| `meta.json` raw bytes | A003 `5A1C2066069845C41B8B0A1E4656B298BBA017C45EC377A317C65EAC055D4366` | A005 `8483FBE4C364386EB5A54D29D478682EC89F7E0345B64FF0B3DDCC1EEF043110` | recorded; timestamped raw bytes are not the semantic gate |

## Resources and private lifecycle exposure

| Resource field | Accepted A003 | Corrected A005 | Delta | Result |
|---|---:|---:|---:|---|
| `startAllocBytes` | `1,315,192` | `1,317,600` | `+2,408` | `PASS` exposed |
| `endAllocBytes` | `886,583,856` | `908,981,224` | `+22,397,368` | `PASS` exposed |
| `maxObservedSys` | `1,995,295,224` | `1,820,223,992` | `-175,071,232` | `PASS` exposed |
| Private run-scoped `O(U+B)` one-canonical-payload lifecycle | not present as a distinct runtime field in `benchmark.json`, stdout, stderr, process metadata, or profile metadata | `NOT EXPOSED` |

No missing private lifecycle field is inferred. The packet records `NOT EXPOSED` exactly as required; it does not reuse implementation/test proof as runtime measurement evidence.

## Corrected packet artifacts

| Artifact | Bytes | SHA-256 |
|---|---:|---|
| `E:\Anvien\.tmp\child06a_a005_cheapapp_frozen_17child_20260827_retry1\candidate\benchmark.json` | `149,143` | `FF93F805150C482EEDA31A1CDCADC5CC70B0BA3C4C21653C705638909B9A8C15` |
| `E:\Anvien\.tmp\child06a_a005_cheapapp_frozen_17child_20260827_retry1\candidate\stdout.json` | `8,824` | `F52193AD16DD18C3DBD78CC161B503CF4091BDFAC1B363962FEBC2C50AD7FEF4` |
| `E:\Anvien\.tmp\child06a_a005_cheapapp_frozen_17child_20260827_retry1\candidate\stderr.log` | `213` | `0FDCF15222C02C25771C98E17D9220C89FDCD0294137A6AE4D9B3AAFD4761ADC` |
| `E:\Anvien\.tmp\child06a_a005_cheapapp_frozen_17child_20260827_retry1\candidate\process.json` | `2,500` | `2E6FA27290E4578472B0ACAC02FCC2B147778DF181EB4D8333A5B8C94B7C6201` |
| `E:\Anvien\.tmp\child06a_a005_cheapapp_frozen_17child_20260827_retry1\candidate\resolution.cpu.pprof` | `61,024` | `0DA6425130E12F9169BCC990C349710D4BA4277E9C30612B9509CE61D9EF9389` |
| `E:\cheapapp.org\.anvien\graph.json` | `840,614,023` | `8763AA2D089328F3D782A5F01E06178D171922E3A5CE96A4EE6707510976C920` |
| `E:\cheapapp.org\.anvien\lbug` | `597,946,368` | `4FC410680111475FDD0B65EE63AFB582B06BEDDE520D795CCD41A7B86472E572` |
| `E:\cheapapp.org\.anvien\meta.json` | `219` | `8483FBE4C364386EB5A54D29D478682EC89F7E0345B64FF0B3DDCC1EEF043110` |

## Corrected recovery boundary and next owner

- The failed root remains preserved and was not rerun.
- The corrected root was launched exactly once; no further retry occurred.
- The concurrently running corrected Restaurant process was permitted and was not killed, waited for, or classified as a conflict.
- No disposition, ledger update, Supervisor, detect, stage, commit, source/test/script/plan mutation, A003 rerun, A004 audit, target cleanup, D002-D017/P3/Child 07 action, or cross-target averaging occurred.
- D001 streak remains unchanged.
- Next owner: Main Orchestration.

## Final report-only repository validation

| Check | Result |
|---|---|
| `git diff --check -- reports/Investigation/rp_child06a_a005_cheapapp_benchmark_frozen_17child.md` | `PASS` (exit `0`, no output) |
| Staged set remains empty | `PASS` (`0` paths) |
