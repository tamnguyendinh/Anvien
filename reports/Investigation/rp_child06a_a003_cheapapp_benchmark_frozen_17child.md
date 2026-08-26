# Child 06A A003 Cheapapp Frozen 17-Child Benchmark

Status: MEASUREMENT_COMPLETE.

This is measurement data for E:\cheapapp.org only. It does not make a KEEP, no-KEEP, Supervisor, acceptance, rejection, disposition, baseline-promotion, checklist, or Restaurant Manager claim.

## Result

The sole authorized A003 launch completed with exit 0 and produced a valid same-work packet: 30/30 unique top-level operations, 17/17 unique resolution children, D001 calls=27,890/files=887, exactly 2,675 exclusive intervals, zero overlap, exact parent conservation, and a nonempty resolution CPU profile.

| Boundary | Accepted A002 before | A003 candidate | Candidate - before |
|---|---:|---:|---:|
| D001 resolve_calls | 3.090914200 s | 3.447846300 s | +0.356932100 s (+11.547784%) |
| B1-P1A-OP001 resolution parent | 19.040468000 s | 20.472602300 s | +1.432134300 s (+7.521529%) |
| Analyzer total | 100.843249000 s | 93.531974900 s | -7.311274100 s (-7.250137%) |
| Process wall | 136.729876000 s | 95.630648200 s | -41.099227800 s (-30.058703%) |

These are measured values for Main handoff; no performance disposition is made here. The complete 30-operation and 17-child before/candidate/delta/rank/share/denominator inventories are in comparison.json.

## Exact launch and frozen identity

| Field | Exact value |
|---|---|
| Target | E:\cheapapp.org |
| Target HEAD | a869876ab6262dacde6cd5d432d099a91852a646 |
| Anvien HEAD | 90edf7fe99cd9600b99c1947c2483f8c5fe2c67c |
| Executable | E:\Anvien\.tmp\child06a_a003_a00x_benchmark_build\anvien-a003-benchmark.exe |
| Version / bytes / SHA-256 | 1.2.8 / 73,823,232 / 2DBBD7AC70C04FB62C3D8AB4A90F50E7F90C51251E82B67883A58B61D42B426B |
| Adjacent DLL bytes / SHA-256 | 20,230,656 / 20CBD87840483A2053CFF3FC2DB23A86DD802B8915D86509D41A4B709624CDB7 |
| Provenance | E:\Anvien\.tmp\child06a_a003_a00x_benchmark_build\anvien-a003-benchmark.provenance.json |
| Provenance bytes / SHA-256 | 13,012 / CB91D5F7FC2EB2810E6A02AED36B5C7E0791F3973487356336B67A1BC8A66F76 |
| Provenance schema / attempt / overlay / candidate / native / exit | 1 / A003 / 2/2 / 3/3 / 3/3 / 0 |
| Launch count / PID / exit | 1 / 15988 / 0 |
| Start / end | 2026-08-26T16:24:38.4264266+07:00 / 2026-08-26T16:26:14.0570748+07:00 |
| Process wall / CPU | 95.630648200 s / 101.687500000 s |
| User / kernel CPU | 91.000000000 s / 10.687500000 s |
| Competing analyze/benchmark processes | 0 |
| Exact argv | E:\Anvien\.tmp\child06a_a003_a00x_benchmark_build\anvien-a003-benchmark.exe analyze E:\cheapapp.org --force --skip-git --json --progress --benchmark-json E:\Anvien\.tmp\child06a_a003_cheapapp_frozen_17child_20260826\candidate\benchmark.json --benchmark-label child06a-a003-cheapapp-frozen-17child-20260826 |

Executable, DLL, provenance, and all five recorded production/overlay source hashes match the frozen preflight identities. Target HEAD/full status remained exact and the Anvien staged path count remained zero. Per Owner/Main correction, concurrent deltas confined to the four Child 06A ledgers or Main/Coder/measurement reports are excluded from the functional mutation gate.

## Conservation and resources

| Conservation field | Accepted A002 before | A003 candidate |
|---|---:|---:|
| Resolution parent | 19.040468000 s | 20.472602300 s |
| 17-child sum | 19.024756200 s | 20.449866500 s |
| Parent residual | 0.015711800 s | 0.022735800 s |
| Exclusive intervals | 2,675 | 2,675 |
| Overlap count | 0 | 0 |

Every A003 child duration equals the exact sum of its intervalEnd - intervalStart pairs; child sum plus residual equals the same-run parent exactly.

| Resource field | Accepted A002 before | A003 candidate | Delta |
|---|---:|---:|---:|
| startAllocBytes | 1,315,328 | 1,315,192 | -136 |
| endAllocBytes | 1,055,121,264 | 886,583,856 | -168,537,408 |
| maxObservedSys | 1,972,115,960 | 1,995,295,224 | +23,179,264 |

## Same-work and output equivalence

| Field | Accepted A002 | A003 candidate | Result |
|---|---:|---:|---|
| Files scanned / parsed / failed | 1,359 / 887 / 0 | 1,359 / 887 / 0 | equal |
| Parser bytes | 5,430,581 | 5,430,581 | equal |
| Graph nodes / relationships | 95,762 / 129,716 | 95,762 / 129,716 | equal |
| Dependency edges / projection unresolved | 13,360 / 735 | 13,360 / 735 | equal |
| DB rows nodes / relationships | 95,762 / 129,716 | 95,762 / 129,716 | equal |
| Top-level operations / resolution children | 30 / 17 | 30 / 17 | equal denominators |
| Resolution semantic counters, diagnostics, outcomes | accepted A002 object | A003 object | exact JSON equality excluding timers |
| canonical graph.json | 840,614,023 bytes; 8763AA2D089328F3D782A5F01E06178D171922E3A5CE96A4EE6707510976C920 | 840,614,023 bytes; 8763AA2D089328F3D782A5F01E06178D171922E3A5CE96A4EE6707510976C920 | exact hash match |
| stdout.json | 8,824 bytes; F52193AD16DD18C3DBD78CC161B503CF4091BDFAC1B363962FEBC2C50AD7FEF4 | 8,824 bytes; F52193AD16DD18C3DBD78CC161B503CF4091BDFAC1B363962FEBC2C50AD7FEF4 | exact hash match |
| stderr.log | 213 bytes; 0FDCF15222C02C25771C98E17D9220C89FDCD0294137A6AE4D9B3AAFD4761ADC | 213 bytes; 0FDCF15222C02C25771C98E17D9220C89FDCD0294137A6AE4D9B3AAFD4761ADC | exact hash match |
| Ladybug lbug | 597,946,368 bytes; 2FEC9F35F7662B45789865DC9EB40A97555AAA1A68DB243BBC3A8D9B778C5AF4 | 597,946,368 bytes; DAB5EA00BA28728E23340CAE3BAA470DBC1B6937972B14C6686FC3E281374885 | recorded; raw-byte equality is not the semantic gate |
| meta.json | 219 bytes; 2FC9B672D97543DF42BED79F7CF19F113A70E1AA3A0561A871DA263557A9D43E | 219 bytes; 5A1C2066069845C41B8B0A1E4656B298BBA017C45EC377A317C65EAC055D4366 | recorded; raw-byte equality is not the semantic gate |

## Raw packet

| Artifact | Bytes | SHA-256 |
|---|---:|---|
| E:\Anvien\.tmp\child06a_a003_cheapapp_frozen_17child_20260826\preflight.json | 7,815 | FA9DA38A38834623672160D1403025F67BCBAE57F099F2D6E3770BA012E1DE64 |
| E:\Anvien\.tmp\child06a_a003_cheapapp_frozen_17child_20260826\candidate\benchmark.json | 149,592 | E98D5CF621FE330392A66ED7B75A86F9009A88E1BB480398E5A2CE783E0F9C02 |
| E:\Anvien\.tmp\child06a_a003_cheapapp_frozen_17child_20260826\candidate\process.json | 2,236 | 05FFFA450A2C5889E7A66885B3A35604B0606DBA6A5DD39B6122A5D747C4358A |
| E:\Anvien\.tmp\child06a_a003_cheapapp_frozen_17child_20260826\candidate\stdout.json | 8,824 | F52193AD16DD18C3DBD78CC161B503CF4091BDFAC1B363962FEBC2C50AD7FEF4 |
| E:\Anvien\.tmp\child06a_a003_cheapapp_frozen_17child_20260826\candidate\stderr.log | 213 | 0FDCF15222C02C25771C98E17D9220C89FDCD0294137A6AE4D9B3AAFD4761ADC |
| E:\Anvien\.tmp\child06a_a003_cheapapp_frozen_17child_20260826\candidate\resolution.cpu.pprof | 65,363 | 806FAD79B12C567618778A8232E8521991882CE6C745AC2EBBE7B09ECC267107 |
| E:\cheapapp.org\.anvien\graph.json | 840,614,023 | 8763AA2D089328F3D782A5F01E06178D171922E3A5CE96A4EE6707510976C920 |
| E:\cheapapp.org\.anvien\lbug | 597,946,368 | DAB5EA00BA28728E23340CAE3BAA470DBC1B6937972B14C6686FC3E281374885 |
| E:\cheapapp.org\.anvien\meta.json | 219 | 5A1C2066069845C41B8B0A1E4656B298BBA017C45EC377A317C65EAC055D4366 |
| E:\Anvien\.tmp\child06a_a003_cheapapp_frozen_17child_20260826\validation.json | 21,705 | 11C14AC7E98A6F98D49E468BC7B01829AA7A346ED33E29B245BD62E1BE9E1C3B |
| E:\Anvien\.tmp\child06a_a003_cheapapp_frozen_17child_20260826\comparison.json | 40,167 | 8D64F905A4413E375DF8CA75E6465EE09BBBEA777DC755DF304F09BB67691C2F |
| E:\Anvien\.tmp\child06a_a003_cheapapp_frozen_17child_20260826\output-manifest.json | 2,348 | C82844F388E33FFE3F468F8BEBD36E570E4C069A4022AFFE03A8CA613B25E245 |

validation.json records the complete gate results and source/status exclusions authorized by Owner/Main. comparison.json contains the full machine-readable 30-operation and 17-child inventories. No analyze retry, build, test, repair, Supervisor, disposition, stage, commit, cleanup, or baseline rerun occurred.
