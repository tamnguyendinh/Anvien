# Child 06A A005 Restaurant Manager Frozen Measurement

Status: `A005_RESTAURANT_MEASUREMENT_PACKET_READY`

## Result

`A005_RESTAURANT_MEASUREMENT_PACKET_READY`

The Main-authorized corrected recovery launch exited `0` and produced the complete assigned A005 Restaurant Manager `retry1` packet. All assigned structural, same-work, canonical Graph, public-output, graph/DB, semantic, counter, and resource-field checks were completed against accepted A003. No disposition is made here.

### Preserved failed-input event

The original launch exited `1` and did not create its `candidate\benchmark.json`. It remains recorded as a failed-input measurement-support event; it is not a candidate measurement, retry target, architecture/implementation attempt, or D001-streak change. Main explicitly authorized one corrected recovery launch under a distinct root.

Exact blocker:

```text
resolution phase: ANVIEN_OP001_RESOLUTION_CPU_PROFILE must name the resolution CPU profile artifact
```

## Frozen identity preflight

| Check | Result | Evidence |
|---|---|---|
| Executable exists | `PASS` | `E:\Anvien\.tmp\child06a_a005_a00x_benchmark_build_retry1\anvien-a005-benchmark.exe` |
| Executable bytes | `PASS` | expected/actual `73825792 / 73825792` |
| Executable SHA-256 | `PASS` | expected/actual `0178DAC396E0DD4D0BECCE7B803378A3400013F0B508EC95F9221A0020CF96F3 / 0178DAC396E0DD4D0BECCE7B803378A3400013F0B508EC95F9221A0020CF96F3` |
| Executable version | `PASS` | provenance output version `1.2.8`; PE file/product version fields are not exposed |
| Adjacent DLL | `PASS` | `E:\Anvien\.tmp\child06a_a005_a00x_benchmark_build_retry1\lbug_shared.dll`; `20230656` bytes; SHA-256 `20CBD87840483A2053CFF3FC2DB23A86DD802B8915D86509D41A4B709624CDB7` |
| Provenance SHA-256 | `PASS` | `E4B9C7E9F9850A58356D41131C7CC177CFC6207D7D8CD3E158B633524B56244E` |
| Provenance schema/attempt/exits | `PASS` | `1 / A005 / 0 / 0` |
| Provenance counts | `PASS` | mappings/candidate/native `2 / 4 / 3` |
| Provenance hash pairs | `PASS` | overlay manifest, both mapped replacements, all four candidate sources, and all three native inputs have identical expected/actual hashes |
| Target identity | `PASS` | requested and resolved target `E:\Restaurant_manager` |
| Output collision gate | `PASS` | output root, candidate directory, benchmark/stdout/stderr artifacts, and assigned report were absent before launch |

The version check consumed the already-frozen provenance identity; no separate executable invocation was added.

## Failed-input event execution

- FileName: `E:\Anvien\.tmp\child06a_a005_a00x_benchmark_build_retry1\anvien-a005-benchmark.exe`
- WorkingDirectory: `E:\Restaurant_manager`
- Process API: direct `System.Diagnostics.ProcessStartInfo`
- `UseShellExecute=false`; `CreateNoWindow=true`; stdout/stderr redirected
- Launch count: `1`
- Process ID: `15776`
- UTC start: `2026-08-26T21:34:53.6516411Z`
- UTC end: `2026-08-26T21:35:30.8686415Z`
- Exit code: `1` — `FAIL`
- Stopwatch wall: `37.214553900 s`
- Process CPU/user/kernel: `22.765625000 / 20.859375000 / 1.906250000 s`

Exact ordered `ArgumentList.Add(...)` values:

1. `analyze`
2. `E:\Restaurant_manager`
3. `--force`
4. `--json`
5. `--progress`
6. `--exclude`
7. `electron/renderer/src/api/userApi.ts`
8. `--benchmark-json`
9. `E:\Anvien\.tmp\child06a_a005_restaurant_manager_frozen_17child_20260827\candidate\benchmark.json`
10. `--benchmark-label`
11. `child06a-a005-restaurant-manager-frozen-17child-20260827`

Equivalent command rendering, for inspection only:

```text
& "E:\Anvien\.tmp\child06a_a005_a00x_benchmark_build_retry1\anvien-a005-benchmark.exe" analyze E:\Restaurant_manager --force --json --progress --exclude electron/renderer/src/api/userApi.ts --benchmark-json E:\Anvien\.tmp\child06a_a005_restaurant_manager_frozen_17child_20260827\candidate\benchmark.json --benchmark-label child06a-a005-restaurant-manager-frozen-17child-20260827
```

## Failed-input event artifacts

| Artifact | Result | Bytes | SHA-256 |
|---|---|---:|---|
| `E:\Anvien\.tmp\child06a_a005_restaurant_manager_frozen_17child_20260827\candidate\benchmark.json` | `FAIL` — not created | n/a | n/a |
| `E:\Anvien\.tmp\child06a_a005_restaurant_manager_frozen_17child_20260827\candidate\stdout.txt` | captured | `0` | `E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855` |
| `E:\Anvien\.tmp\child06a_a005_restaurant_manager_frozen_17child_20260827\candidate\stderr.txt` | captured | `234` | `AC4E9403D3ADA081B035B331FBFDD06158ECC8D043ECC271DDA641B938A3E164` |

Captured stderr, exactly as decoded from the redirected stream:

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

## Failed-input event versus A003

Accepted A003 Restaurant basis was not rerun: D001/parent/analyzer/process `9.401585300 / 20.850792800 / 98.020546700 / 101.096911900 s`, `calls=86030; files=1234`.

| Boundary | Accepted A003 | A005 candidate | Delta | Check |
|---|---:|---:|---:|---|
| D001 `resolve_calls` | `9.401585300 s` | not exposed | not exposed | `NOT EXPOSED` — benchmark absent |
| Parent `resolution` | `20.850792800 s` | not exposed | not exposed | `NOT EXPOSED` — benchmark absent |
| Analyzer | `98.020546700 s` | not exposed | not exposed | `NOT EXPOSED` — benchmark absent |
| Process wall | `101.096911900 s` | `37.214553900 s` | `-63.882358000 s` | `FAIL / INCOMPARABLE` — A005 exited `1` before completing the workload; arithmetic is recorded only as failed-run identity, never as performance evidence |

No A005 denominator was emitted. The accepted A003 denominator is recorded only as the comparison authority; it is not attributed to the failed candidate.

## Failed-input event validation matrix

| Required check | Result | Exact fact |
|---|---|---|
| Exact command/runtime/target/output/process identity | `PASS` | exact frozen FileName, working directory, ordered 11 arguments, redirected streams, PID, timestamps, and one launch recorded |
| Successful candidate exit | `FAIL` | exit `1` |
| Candidate benchmark artifact | `FAIL` | `benchmark.json` absent |
| Candidate D001, parent, analyzer elapsed | `NOT EXPOSED` | no benchmark JSON |
| Candidate process wall | `FAIL / EXPOSED` | `37.214553900 s`, but incomplete/non-comparable due exit `1` |
| Exact top-level operations | `NOT EXPOSED` | cannot establish `30/30` or uniqueness without benchmark JSON |
| Exact resolution children | `NOT EXPOSED` | cannot establish `17/17`, uniqueness, or mapping without benchmark JSON |
| Candidate denominator | `NOT EXPOSED` | cannot establish `calls=86030; files=1234` for A005 |
| Interval conservation | `NOT EXPOSED` | no interval rows |
| Zero overlap | `NOT EXPOSED` | no interval rows |
| Complete ordered evidence array | `NOT EXPOSED` | no candidate result/output packet |
| Canonical Graph JSON SHA/semantic equality | `NOT EXPOSED` | no completed candidate packet was available for comparison; target/index state was not cleaned or reinterpreted |
| Public stdout semantics | `FAIL` | candidate stdout is empty and cannot equal the completed A003 public output semantics |
| Public stderr semantics | `FAIL` | candidate stderr contains the fatal resolution-profile requirement and cannot equal completed A003 semantics |
| Raw stdout comparison | `FAIL` | A005 `E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855`; A003 `CD0028F3B3939451FF4D3623DF7DD7482A993E0934D2DFA773682D37D4FB0E94` |
| Raw stderr comparison | `FAIL` | A005 `AC4E9403D3ADA081B035B331FBFDD06158ECC8D043ECC271DDA641B938A3E164`; A003 `0FDCF15222C02C25771C98E17D9220C89FDCD0294137A6AE4D9B3AAFD4761ADC` |
| Graph/DB readback | `NOT EXPOSED` | no completed packet; no post-failure target/index inspection was performed |
| Resolution outcomes/diagnostics/counters | `NOT EXPOSED` | no benchmark/result packet |
| Files scanned/parsed/failed | `NOT EXPOSED` | no benchmark/result packet |
| Parser bytes | `NOT EXPOSED` | no benchmark/result packet |
| Graph nodes/relationships | `NOT EXPOSED` | no benchmark/result packet |
| `startAllocBytes` | `NOT EXPOSED` | no benchmark JSON |
| `endAllocBytes` | `NOT EXPOSED` | no benchmark JSON |
| `maxObservedSys` | `NOT EXPOSED` | no benchmark JSON |
| Private run-scoped `O(U+B)` one-canonical-payload lifecycle proof | `NOT EXPOSED` | no candidate evidence/resource packet; no inference from implementation or prior tests |
| Same-work/output equivalence against A003 | `FAIL` | required candidate evidence is missing and public streams show a failed run |

## Corrected recovery execution

- Output root: `E:\Anvien\.tmp\child06a_a005_restaurant_manager_frozen_17child_20260827_retry1`
- FileName: `E:\Anvien\.tmp\child06a_a005_a00x_benchmark_build_retry1\anvien-a005-benchmark.exe`
- WorkingDirectory: `E:\Anvien`
- Process API: direct `System.Diagnostics.ProcessStartInfo`
- `UseShellExecute=false`; `CreateNoWindow=true`; stdout/stderr redirected
- Corrected launch count: `1`
- Process ID: `17228`
- UTC start: `2026-08-26T21:41:56.7304748Z`
- UTC end: `2026-08-26T21:44:39.0222374Z`
- Exit code: `0` — `PASS`
- Process wall: `162.290202200 s`
- Process CPU/user/kernel: `112.046875000 / 96.359375000 / 15.687500000 s`

Exact ordered arguments:

1. `analyze`
2. `E:\Restaurant_manager`
3. `--force`
4. `--json`
5. `--progress`
6. `--exclude`
7. `electron/renderer/src/api/userApi.ts`
8. `--benchmark-json`
9. `E:\Anvien\.tmp\child06a_a005_restaurant_manager_frozen_17child_20260827_retry1\candidate\benchmark.json`
10. `--benchmark-label`
11. `child06a-a005-restaurant-manager-frozen-17child-20260827-retry1`

Benchmark label emitted by the packet: `child06a-a005-restaurant-manager-frozen-17child-20260827-retry1` — `PASS`.

### Exact child environment

All `17` specified overrides were set through `ProcessStartInfo.Environment` and are persisted verbatim in `candidate\process.json`:

| Variable | Exact value |
|---|---|
| `TEMP` | `E:\Anvien\.tmp\child06a_a005_restaurant_manager_frozen_17child_20260827_retry1\candidate\process-local\temp` |
| `TMP` | `E:\Anvien\.tmp\child06a_a005_restaurant_manager_frozen_17child_20260827_retry1\candidate\process-local\tmp` |
| `TMPDIR` | `E:\Anvien\.tmp\child06a_a005_restaurant_manager_frozen_17child_20260827_retry1\candidate\process-local\tmp` |
| `HOME` | `E:\Anvien\.tmp\child06a_a005_restaurant_manager_frozen_17child_20260827_retry1\candidate\process-local\home` |
| `USERPROFILE` | `E:\Anvien\.tmp\child06a_a005_restaurant_manager_frozen_17child_20260827_retry1\candidate\process-local\home` |
| `APPDATA` | `E:\Anvien\.tmp\child06a_a005_restaurant_manager_frozen_17child_20260827_retry1\candidate\process-local\home\AppData\Roaming` |
| `LOCALAPPDATA` | `E:\Anvien\.tmp\child06a_a005_restaurant_manager_frozen_17child_20260827_retry1\candidate\process-local\home\AppData\Local` |
| `XDG_CACHE_HOME` | `E:\Anvien\.tmp\child06a_a005_restaurant_manager_frozen_17child_20260827_retry1\candidate\process-local\home\.cache` |
| `XDG_CONFIG_HOME` | `E:\Anvien\.tmp\child06a_a005_restaurant_manager_frozen_17child_20260827_retry1\candidate\process-local\home\.config` |
| `XDG_DATA_HOME` | `E:\Anvien\.tmp\child06a_a005_restaurant_manager_frozen_17child_20260827_retry1\candidate\process-local\home\.local\share` |
| `GOCACHE` | `E:\Anvien\.tmp\child06a_a005_restaurant_manager_frozen_17child_20260827_retry1\candidate\process-local\gocache` |
| `GOMODCACHE` | `E:\Anvien\.tmp\child06a_a005_restaurant_manager_frozen_17child_20260827_retry1\candidate\process-local\gomodcache` |
| `GOTMPDIR` | `E:\Anvien\.tmp\child06a_a005_restaurant_manager_frozen_17child_20260827_retry1\candidate\process-local\gotmp` |
| `ANVIEN_OP001_RESOLUTION_CPU_PROFILE` | `E:\Anvien\.tmp\child06a_a005_restaurant_manager_frozen_17child_20260827_retry1\candidate\resolution.cpu.pprof` |
| `GIT_CONFIG_COUNT` | `1` |
| `GIT_CONFIG_KEY_0` | `safe.directory` |
| `GIT_CONFIG_VALUE_0` | `E:/Restaurant_manager` |

## Corrected packet artifacts

| Artifact | Bytes | SHA-256 | Check |
|---|---:|---|---|
| `candidate\benchmark.json` | `197467` | `87FE48DD910082DB9D674F5B40634F6F707258E490A905A3090D70ECC0B19C1F` | `PASS` |
| `candidate\stdout.json` | `8944` | `CD0028F3B3939451FF4D3623DF7DD7482A993E0934D2DFA773682D37D4FB0E94` | `PASS`; exact A003 hash |
| `candidate\stderr.log` | `213` | `0FDCF15222C02C25771C98E17D9220C89FDCD0294137A6AE4D9B3AAFD4761ADC` | `PASS`; exact A003 hash |
| `candidate\process.json` | `3652` | `246EBFE5D86916F207BBF8C2AB2A283273597FB4C8B89DF6841C5B9CCAB9310A` | `PASS` |
| `candidate\resolution.cpu.pprof` | `56273` | `5FC396A5DD96859B9E863F91CD029ACCB0E55C2B8C1744CD8E9388D5426F98B3` | `PASS`; nonempty |

The child process also materialized its runtime registry inside the isolated HOME at `candidate\process-local\home\.anvien\registry.json`: `379` bytes, SHA-256 `5A3FEAC2D77B47847E50343D23ED7E65DD344432259F5DA0D551144F8D5F418D`. It was not created by the pre-launch directory setup, is not a sixth assigned candidate output, and was preserved without cleanup as process-local run evidence.

## Corrected A005 versus accepted A003

The A003 Restaurant packet was not rerun. Candidate deltas are A005 minus accepted A003 and are never averaged with Cheapapp.

| Boundary | Accepted A003 | Corrected A005 | Delta | Percent delta |
|---|---:|---:|---:|---:|
| D001 `resolve_calls` | `9.401585300 s` | `9.142619400 s` | `-0.258965900 s` | `-2.754492%` |
| Parent `resolution` | `20.850792800 s` | `19.678482100 s` | `-1.172310700 s` | `-5.622380%` |
| Analyzer | `98.020546700 s` | `122.900035100 s` | `+24.879488400 s` | `+25.381911%` |
| Process wall | `101.096911900 s` | `162.290202200 s` | `+61.193290300 s` | `+60.529337%` |

These are measurement facts only. This lane makes no `KEEP`, `NO_KEEP`, streak, or disposition decision.

## Operation and child structure

- Top-level operations: `30/30`, unique `30`, exact A003 order `PASS`, and all `30` name/boundary/denominator mappings `PASS`.
- Exact ordered top-level operations: `cli_startup`, `cli_preparation`, `analyze_setup`, `scan`, `structure`, `documents`, `cobol`, `parse`, `routes`, `tools`, `orm`, `cross_file_binding`, `resolution`, `mro`, `communities`, `processes`, `semantic_enrichment`, `graph_compact`, `db_runner_resolve`, `db_load`, `db_runner_close`, `graph_snapshot`, `benchmark_write`, `analyzer_orchestration`, `memory_profile`, `registry_meta`, `ai_context`, `file_projection`, `output_publication`, `cpu_profile_completion`.
- Resolution children: `17/17`, unique `17`, exact A003 order `PASS`, and all `17` child/source/call-path/denominator mappings `PASS`.
- Exact ordered resolution children: `build_binding_occurrence_index`, `runtime_setup`, `emit_definition_nodes`, `emit_unresolved_heritage_diagnostics`, `emit_import_edges`, `emit_heritage_compatibility_edges`, `resolve_calls`, `resolve_accesses`, `resolve_type_annotations`, `emit_method_dispatch_edges`, `finalize_typescript_authority_results`, `emit_typescript_external_symbols`, `finalize_resolution_outcomes`, `project_resolution_outcomes`, `finalize_resolution_metadata`, `assemble_resolution_result`, `binding_accumulator_dispose`.
- D001 denominator: `calls=86030; files=1234` — exact A003 equality `PASS`.

## Interval conservation

| Check | Corrected A005 | Result |
|---|---:|---|
| Exclusive intervals | `3716` | `PASS` |
| Overlap count | `0` | `PASS` |
| Interval shapes | all start/end arrays paired and nonnegative | `PASS` |
| Per-child interval sums | every sum equals its `durationNs` | `PASS` |
| Parent | `19678482100 ns` | measured |
| Child sum | `19618545400 ns` | measured |
| Residual | `59936700 ns` | measured |
| Recomposed | `19678482100 ns` | `PASS`; exact parent conservation |
| Maximum interval end | `19618931900 ns` | `PASS`; within parent |

## Same-work, output, graph, DB, and semantic parity

| Required check | Result | Exact evidence |
|---|---|---|
| Non-timing benchmark semantics | `PASS` | canonical normalized A003/A005 SHA-256 `2B42F7D908F0F119D10C8E106D0F3CBBE8AC6C40E89C469CB41284D9C7FF6E81 / 2B42F7D908F0F119D10C8E106D0F3CBBE8AC6C40E89C469CB41284D9C7FF6E81`; normalization excludes only label, timing/interval, and separately reported memory fields |
| Complete ordered evidence array | `PASS` | the complete `900212685`-byte canonical Graph JSON is byte-identical to A003; therefore every serialized Evidence element and its order are identical, without sampling or reserialization |
| Canonical Graph JSON | `PASS` | A003/A005 SHA-256 `09708B038AD63B659A95BAC34304A752F2C6350952B91CF435EBD5C5F9D94B7C / 09708B038AD63B659A95BAC34304A752F2C6350952B91CF435EBD5C5F9D94B7C` |
| Public stdout bytes and semantics | `PASS` | exact SHA-256 `CD0028F3B3939451FF4D3623DF7DD7482A993E0934D2DFA773682D37D4FB0E94`; parsed JSON objects equal |
| Public stderr bytes and semantics | `PASS` | exact SHA-256 `0FDCF15222C02C25771C98E17D9220C89FDCD0294137A6AE4D9B3AAFD4761ADC` |
| Files | `PASS` | scanned/parsed/failed `1556 / 1234 / 0`; complete file classification object equals A003 |
| Parser | `PASS` | total/succeeded/failed/unsupported/timed-out/created/bytes `1234 / 1234 / 0 / 0 / 0 / 4 / 8546179`; all non-timing parser fields equal A003 |
| Public graph counts | `PASS` | nodes/relationships `137229 / 190644`, exact A003 equality |
| DB readback | `PASS` | node/relationship rows `137229 / 190644`; copy counts `19 / 92`; fallback inserts/failures/skipped relationships `0 / 0 / 0`; complete `dbLoad` object equals A003 |
| Resolution scalar counters | `PASS` | complete scalar object equals A003 |
| Outcome cardinality across lifecycle rows | `PASS` | finalized outcome map/projected outcomes/assembled outcomes `186251 / 186251 / 186251`; exact A003 denominator equality |
| Diagnostic counters | `PASS` | unresolved references/diagnostics/unattributed `124409 / 129009 / 0` |
| Reference and graph-emission counters | `PASS` | resolved references/calls/accesses/types `49240 / 20278 / 13925 / 15037`; graph emitted nodes/relationships `66421 / 112049` |
| Dependency/projection public semantics | `PASS` | stdout file projection is object-identical to A003: files/dependency edges/unresolved `1556 / 17822 / 853` |
| Ladybug database presence/readback boundary | `PASS` | `503070720` bytes; semantic readback/count gates above pass; raw SHA-256 `3DC1A1D597E86A17091955803EB9BF993E75CEB68242AD5EC7B74CF2642A822B` differs from A003 and is non-governing alone |
| Ladybug meta presence | `PASS` | `267` bytes; raw SHA-256 `C74885C295F2FC928150948A58FA3DC3D9E8871D6D219139D6F7E5D126880519` differs from A003 and is non-governing alone |

Resolution scalar equality includes definitions/imports/import uses `60854 / 2259 / 2199`; resolved/unresolved references `49240 / 124409`; diagnostics `129009`; resolved calls/accesses/type references `20278 / 13925 / 15037`; external capability unavailable/profile excluded/meaning mismatch `9136 / 0 / 0`; heritage indexed/resolved/unresolved `43 / 19 / 24`; duplicate edges/finalized imports `18736 / 2259`; binding accumulator files/entries/finalized/disposed `346 / 1220 / true / true`; and every remaining scalar flag/count.

## Resources and lifecycle exposure

| Field | Accepted A003 | Corrected A005 | Delta | Result |
|---|---:|---:|---:|---|
| `startAllocBytes` | `1474592` | `1476832` | `+2240` | `EXPOSED` |
| `endAllocBytes` | `1306049888` | `785046824` | `-521003064` | `EXPOSED` |
| `maxObservedSys` | `2634607224` | `2651937400` | `+17330176` | `EXPOSED` |
| CPU profile | `56397 bytes` | `56273 bytes` | `-124 bytes` | `PASS`; candidate nonempty |
| Private run-scoped `O(U+B)` one-canonical-payload lifecycle counter/proof | n/a | no canonical-payload, sidecar, or encoding-count field is present in the corrected benchmark/output/process packet | n/a | `NOT EXPOSED`; no inference from implementation or prior tests |

The packet does expose equal semantic outcome cardinality at finalize/project/assemble and the three aggregate memory fields. Those facts are not promoted into an unexposed payload-duplication or lifetime proof.

## Corrected validation result

- Successful corrected exit and all five assigned artifacts: `PASS`.
- `30/30`, `17/17`, denominator, interval conservation, and zero overlap: `PASS`.
- Complete ordered evidence, canonical Graph JSON, stdout/stderr, graph/DB readback, resolution outcomes/diagnostics/counters, file/parser counts, and non-timing semantics against A003: `PASS`.
- Required aggregate resource fields: `EXPOSED` and recorded.
- Private lifecycle counter/proof: `NOT EXPOSED`, explicitly not inferred.
- Packet completeness for Main review: `PASS`.

## Stop-boundary compliance

- Failed-input event launches: `1`; it remains preserved under the original root and was not rerun.
- Main-authorized corrected recovery launches: `1`; further retries: `0`.
- The failed-input event does not change the D001 streak and is not a new architecture/implementation attempt.
- No A003 rerun, Cheapapp averaging, A004 audit, disposition, ledger/plan/source/test/script edit, graph operation, detect, stage, commit, rebuild, repair, D002-D017/P3/Child 07 action, or unrelated-process action occurred.
- The concurrently running corrected Cheapapp process was permitted and was not waited for, killed, or classified as a conflict.
- Only the existing report was updated; no second Restaurant report was created.

Next owner: Main Orchestration.

`A005_RESTAURANT_MEASUREMENT_PACKET_READY`
