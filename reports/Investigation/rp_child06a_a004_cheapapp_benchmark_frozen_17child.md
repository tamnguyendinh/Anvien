# Child 06A A004 Cheapapp Frozen 17-Child Benchmark

Status: MEASUREMENT_COMPLETE.

This is measurement data for E:\cheapapp.org only. It does not make a KEEP, no-KEEP, Supervisor, acceptance, rejection, disposition, baseline-promotion, checklist, or Restaurant Manager claim.

## Result

The sole authorized A004 launch completed with exit 0 and produced a valid same-work packet: 30/30 unique top-level operations, 17/17 unique resolution children, D001 calls=27,890/files=887, exactly 2,675 exclusive intervals, zero overlap, exact parent conservation, and a nonempty resolution CPU profile. Exact canonical graph.json equality proves the full ordered export-binding evidence sequence is equal to accepted A003.

| Boundary | Accepted A003 before | A004 candidate | Candidate - before |
|---|---:|---:|---:|
| D001 resolve_calls | 3.447846300 s | 2.074182500 s | -1.373663800 s (-39.841213%) |
| B1-P1A-OP001 resolution parent | 20.472602300 s | 13.265999200 s | -7.206603100 s (-35.201207%) |
| Analyzer total | 93.531974900 s | 107.287054400 s | +13.755079500 s (+14.706286%) |
| Process wall | 95.630648200 s | 144.975972400 s | +49.345324200 s (+51.599906%) |

These are measured values for Main handoff; no performance disposition is made here. The complete 30-operation and 17-child A003/A004/delta/rank/share/denominator inventories are in comparison.json.

## Exact launch and frozen identity

| Field | Exact value |
|---|---|
| Target | E:\cheapapp.org |
| Target HEAD | a869876ab6262dacde6cd5d432d099a91852a646 |
| Anvien HEAD before / after measurement | 1d6fa0c203d96fd6913938b0c67c44e578224453 / 1d6fa0c203d96fd6913938b0c67c44e578224453 |
| Executable | E:\Anvien\.tmp\child06a_a004_a00x_benchmark_build\anvien-a004-benchmark.exe |
| Version / bytes / SHA-256 | 1.2.8 / 73,825,280 / 6D319467D198B8BBA2375339CDF6BD7634FA97E7C503D3DFC0C8C315D965352C |
| Adjacent DLL bytes / SHA-256 | 20,230,656 / 20CBD87840483A2053CFF3FC2DB23A86DD802B8915D86509D41A4B709624CDB7 |
| Provenance | E:\Anvien\.tmp\child06a_a004_a00x_benchmark_build\anvien-a004-benchmark.provenance.json |
| Provenance bytes / SHA-256 | 12,994 / 4EBF871D0008B6BB1A7FDB433DE07CE437F859DFB34221C94D360E866711A20D |
| Provenance schema / attempt / mappings / candidate / native / exit | 1 / A004 / 2 / 4 / 3 / 0 |
| Launch count / PID / exit | 1 / 12012 / 0 |
| Start / end | 2026-08-27T00:13:48.1277488+07:00 / 2026-08-27T00:16:12.9238146+07:00 |
| Process wall / CPU | 144.9759724 s / 98.5781250 s |
| User / kernel CPU | 83.0000000 s / 15.5781250 s |
| Competing analyze/benchmark processes | 0 unapproved; 0 present |
| Exact argv | E:\Anvien\.tmp\child06a_a004_a00x_benchmark_build\anvien-a004-benchmark.exe analyze E:\cheapapp.org --force --skip-git --json --progress --benchmark-json E:\Anvien\.tmp\child06a_a004_cheapapp_frozen_17child_20260827\candidate\benchmark.json --benchmark-label child06a-a004-cheapapp-frozen-17child-20260827 |

Executable, DLL, provenance, and all four required candidate source hashes matched their frozen identities before launch:

- diagnostics.go: 6DE54D97A1B95E877B686DC3459238598F6060D28CE6329357A80BD7FC376D30
- emit.go: 73AE20E8A18AC25F8019E2A3223407621AE7ABC358DBAF4F058E1D9AC3136060
- outcome.go: 02092F9FE7DA2A4BDB49E13056FBA3C97DC24F416141E7F27866EE80F60C1F7E
- export_binding_proof.go: 36D41BC7336A5A04B4C77D1EE9DEB45DE4E5E2BA120EEFBF98419BD2D846C4D2

## Conservation and resources

| Conservation field | Accepted A003 before | A004 candidate |
|---|---:|---:|
| Resolution parent | 20.472602300 s | 13.265999200 s |
| 17-child sum | 20.449866500 s | 13.251820500 s |
| Parent residual | 0.022735800 s | 0.014178700 s |
| Exclusive intervals | 2,675 | 2,675 |
| Overlap count | 0 | 0 |

Every A004 child duration equals the exact sum of its intervalEnd - intervalStart pairs. The 17-child sum plus residual equals the same-run parent exactly: 13,251,820,500 ns + 14,178,700 ns = 13,265,999,200 ns.

| Resource field | Accepted A003 before | A004 candidate | Delta |
|---|---:|---:|---:|
| startAllocBytes | 1,315,192 | 1,314,424 | -768 |
| endAllocBytes | 886,583,856 | 656,525,488 | -230,058,368 |
| maxObservedSys | 1,995,295,224 | 1,984,702,968 | -10,592,256 |

## Same-work, output, and ordered-evidence equivalence

| Field | Accepted A003 | A004 candidate | Result |
|---|---:|---:|---|
| Files scanned / parsed / failed | 1,359 / 887 / 0 | 1,359 / 887 / 0 | exact |
| Parser bytes | 5,430,581 | 5,430,581 | exact |
| Graph nodes / relationships | 95,762 / 129,716 | 95,762 / 129,716 | exact |
| Dependency edges / projection unresolved | 13,360 / 735 | 13,360 / 735 | exact |
| DB rows nodes / relationships | 95,762 / 129,716 | 95,762 / 129,716 | exact |
| Top-level operations / resolution children | 30 / 17 | 30 / 17 | exact inventories and denominators |
| Resolution semantic counters, diagnostics, outcomes | accepted A003 object | A004 object | exact non-timing JSON equality |
| Full ordered export-binding evidence | canonical graph SHA-256 8763AA2D089328F3D782A5F01E06178D171922E3A5CE96A4EE6707510976C920 | same SHA-256 | exact ordered-evidence proof |
| canonical graph.json | 840,614,023 bytes; 8763AA2D089328F3D782A5F01E06178D171922E3A5CE96A4EE6707510976C920 | 840,614,023 bytes; 8763AA2D089328F3D782A5F01E06178D171922E3A5CE96A4EE6707510976C920 | exact hash match |
| stdout.json | 8,824 bytes; F52193AD16DD18C3DBD78CC161B503CF4091BDFAC1B363962FEBC2C50AD7FEF4 | 8,824 bytes; F52193AD16DD18C3DBD78CC161B503CF4091BDFAC1B363962FEBC2C50AD7FEF4 | exact hash match |
| stderr.log | 213 bytes; 0FDCF15222C02C25771C98E17D9220C89FDCD0294137A6AE4D9B3AAFD4761ADC | 213 bytes; 0FDCF15222C02C25771C98E17D9220C89FDCD0294137A6AE4D9B3AAFD4761ADC | exact hash match |
| Ladybug lbug | recorded A003 boundary | 597,942,272 bytes; 556C43571D265DAA699BDEDD05C3EF06F4583ED843E0BC59B649B77408D337B3 | recorded; raw-byte equality is not the semantic gate |
| meta.json | recorded A003 boundary | 219 bytes; 8A16875F4D25DA959570C06690883979B07F333CB4B71FCF63C5A9DD2EE8945A | recorded; raw-byte equality is not the semantic gate |

The complete non-timing workload objects, including all recorded counts and classification samples, compare exactly. There are no workload, operation-denominator, child-denominator, resolution-semantic, output, or ordered-evidence mismatches.

## Raw packet

| Artifact | Bytes | SHA-256 |
|---|---:|---|
| E:\Anvien\.tmp\child06a_a004_cheapapp_frozen_17child_20260827\preflight.json | 6,815 | C0444AB98308E7EDBDD2B7DFC704DA6150379E2C8553CB666D5747DA28763E56 |
| E:\Anvien\.tmp\child06a_a004_cheapapp_frozen_17child_20260827\candidate\benchmark.json | 148,232 | BD67C0B764083137CB2AD4B473E9281230FB9DF73A5AE699F533FAFA22598833 |
| E:\Anvien\.tmp\child06a_a004_cheapapp_frozen_17child_20260827\candidate\process.json | 2,198 | 0104BB88F96C4F87C393B7E50195F1B0F10BF4672C5CB4394146F4E48AC1A628 |
| E:\Anvien\.tmp\child06a_a004_cheapapp_frozen_17child_20260827\candidate\stdout.json | 8,824 | F52193AD16DD18C3DBD78CC161B503CF4091BDFAC1B363962FEBC2C50AD7FEF4 |
| E:\Anvien\.tmp\child06a_a004_cheapapp_frozen_17child_20260827\candidate\stderr.log | 213 | 0FDCF15222C02C25771C98E17D9220C89FDCD0294137A6AE4D9B3AAFD4761ADC |
| E:\Anvien\.tmp\child06a_a004_cheapapp_frozen_17child_20260827\candidate\resolution.cpu.pprof | 52,040 | C1250E0F4F8655054967125F454FED6D9B7329EA8528726C04B7D65BA61F7616 |
| E:\Anvien\.tmp\child06a_a004_cheapapp_frozen_17child_20260827\comparison.json | 93,750 | EE5AA56C278513045A30DE2E69F0B95B5B7B9677D654B13B6B558CC138DC7FC0 |
| E:\Anvien\.tmp\child06a_a004_cheapapp_frozen_17child_20260827\candidate\validation.json | 20,594 | 974E37335F816498EFC56643B1E3ADD0B7A8E0FEB608283A4E24D28622FB727C |
| E:\Anvien\.tmp\child06a_a004_cheapapp_frozen_17child_20260827\output-manifest.json | 2,819 | C7DEF16725430694A1FC9B3601A289651C7829BB148E209642D5DFCFEA8B46A0 |
| E:\cheapapp.org\.anvien\graph.json | 840,614,023 | 8763AA2D089328F3D782A5F01E06178D171922E3A5CE96A4EE6707510976C920 |
| E:\cheapapp.org\.anvien\lbug | 597,942,272 | 556C43571D265DAA699BDEDD05C3EF06F4583ED843E0BC59B649B77408D337B3 |
| E:\cheapapp.org\.anvien\meta.json | 219 | 8A16875F4D25DA959570C06690883979B07F333CB4B71FCF63C5A9DD2EE8945A |

## Unchanged boundary and handoff

- Target HEAD and complete porcelain-v2 status were byte-for-byte unchanged across measurement.
- The four frozen candidate production hashes and export_binding_proof_test.go hash were unchanged across measurement.
- Anvien HEAD and complete status were unchanged during measurement; the staged set remained empty.
- No source, test, script, or target source was edited; no candidate rebuild, retry, detect, stage, commit, cleanup, Supervisor, acceptance, or disposition action occurred.
- Raw root: E:\Anvien\.tmp\child06a_a004_cheapapp_frozen_17child_20260827
- Next owner: Main Orchestration.

The packet is measurement-only. Main Orchestration owns any later KEEP/no-KEEP decision, Supervisor review, acceptance, baseline promotion, ledger update, staging, or commit.
