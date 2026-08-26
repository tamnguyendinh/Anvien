# Child 06A A004 Restaurant Manager Frozen Measurement

Status: `MEASUREMENT_A004_RESTAURANT_COMPLETE`

## Execution

- Target: `E:\Restaurant_manager`
- FileName: `E:\Anvien\.tmp\child06a_a004_a00x_benchmark_build\anvien-a004-benchmark.exe`
- WorkingDirectory: `E:\Restaurant_manager`
- ArgumentList, in exact order: `analyze`, `E:\Restaurant_manager`, `--force`, `--json`, `--progress`, `--exclude`, `electron/renderer/src/api/userApi.ts`, `--benchmark-json`, `E:\Anvien\.tmp\child06a_a004_restaurant_manager_frozen_17child_20260827\candidate\benchmark.json`, `--benchmark-label`, `child06a-a004-restaurant-manager-frozen-17child-20260827`
- Launch count / PID / exit: `1 / 15116 / 0`
- Start / end: `2026-08-27T00:20:12.7071885+07:00 / 2026-08-27T00:22:28.2780141+07:00`
- Process wall / CPU / user / kernel: `135.569489100 / 111.828125000 / 95.500000000 / 16.328125000 s`

## A003 → A004

| Boundary | Accepted A003 | A004 candidate | Delta |
|---|---:|---:|---:|
| D001 `resolve_calls` | `9.401585300 s` | `8.975767700 s` | `-0.425817600 s` |
| Parent `resolution` | `20.850792800 s` | `19.416099500 s` | `-1.434693300 s` |
| Analyzer | `98.020546700 s` | `101.406172300 s` | `+3.385625600 s` |
| Process | `101.096911900 s` | `135.569489100 s` | `+34.472577200 s` |

## Validation

- Operations / unique: `30 / 30`; exact order, boundaries, full inventories, denominators, ranks, and shares recorded: `PASS`
- Resolution children / unique / mapped: `17 / 17 / 17`; exact order, full inventories, denominators, ranks, and shares recorded: `PASS`
- D001 denominator: `calls=86030; files=1234`
- Benchmark and CPU profile are nonempty; CPU profile: `55,580 bytes`
- Files scanned / parsed / failed: `1556 / 1234 / 0`
- Parser bytes: `8,546,179`; graph/DB rows: `137,229 / 190,644`
- Dependency/projection counts and resolution semantic counters/diagnostics/outcomes: `PASS`
- Target HEAD/full status unchanged: `PASS`; four candidate source hashes unchanged: `PASS`; Anvien staged set empty: `PASS`

## Conservation

- Exclusive intervals / overlap: `3716 / 0`; every interval shape and duration sum exact: `PASS`
- Parent / child sum / residual: `19.416099500 / 19.396949800 / 0.019149700 s`
- Child sum plus residual equals parent exactly; maximum interval end remains within the parent: `PASS`

## Resources

- `startAllocBytes`: `1,474,592 -> 1,470,056`; delta `-4,536`
- `endAllocBytes`: `1,306,049,888 -> 1,029,351,808`; delta `-276,698,080`
- `maxObservedSys`: `2,634,607,224 -> 2,757,036,664`; delta `+122,429,440`

## Equivalence

- Full non-timing benchmark semantics, including same-work files/parser/dependency/projection/semantic/diagnostic/outcome fields: `PASS`
- Canonical `graph.json`: `PASS; 900,212,685 bytes; 09708B038AD63B659A95BAC34304A752F2C6350952B91CF435EBD5C5F9D94B7C / 09708B038AD63B659A95BAC34304A752F2C6350952B91CF435EBD5C5F9D94B7C`
- Full ordered export-binding Evidence arrays: `PASS`; the serialized canonical graph is byte/size/SHA-identical to accepted A003, so ordered arrays are identical rather than merely set-equal.
- Stdout: `PASS; CD0028F3B3939451FF4D3623DF7DD7482A993E0934D2DFA773682D37D4FB0E94 / CD0028F3B3939451FF4D3623DF7DD7482A993E0934D2DFA773682D37D4FB0E94`
- Stderr: `PASS; 0FDCF15222C02C25771C98E17D9220C89FDCD0294137A6AE4D9B3AAFD4761ADC / 0FDCF15222C02C25771C98E17D9220C89FDCD0294137A6AE4D9B3AAFD4761ADC`
- Ladybug/meta recorded, raw-byte equality non-governing alone: `lbug 503,070,720 bytes SHA-256 35F862E2026BEFC1CDC1F1F0B7D2894B8E90D46E032022ECC3CD95F8A8A41C; meta 267 bytes SHA-256 4B25BF19F401F099C3051142405D243B3A176DC9DD20B59FACB6BA96CC857322`

## Artifacts

- Raw root: `E:\Anvien\.tmp\child06a_a004_restaurant_manager_frozen_17child_20260827`
- Report: `E:\Anvien\reports\Investigation\rp_child06a_a004_restaurant_manager_benchmark_frozen_17child.md`
- Preflight: `E:\Anvien\.tmp\child06a_a004_restaurant_manager_frozen_17child_20260827\preflight.json`
- Comparison: `E:\Anvien\.tmp\child06a_a004_restaurant_manager_frozen_17child_20260827\candidate\comparison.json`
- Validation: `E:\Anvien\.tmp\child06a_a004_restaurant_manager_frozen_17child_20260827\candidate\validation.json`
- Manifest: `E:\Anvien\.tmp\child06a_a004_restaurant_manager_frozen_17child_20260827\candidate\output-manifest.json`

## Boundary

Measurement-only. No disposition, baseline rerun/promotion, architecture/code/build/test, Supervisor, ledger, detect, stage, commit, cleanup, or retry. Next owner: Main Orchestration.

`MEASUREMENT_A004_RESTAURANT_COMPLETE`
