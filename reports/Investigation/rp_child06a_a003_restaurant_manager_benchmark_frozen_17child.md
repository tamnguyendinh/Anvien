# Child 06A A003 Restaurant Manager Frozen Measurement

Status: `MEASUREMENT_A003_RESTAURANT_COMPLETE`

## Execution
- Target: `E:\Restaurant_manager`
- Command: `& "E:\Anvien\.tmp\child06a_a003_a00x_benchmark_build\anvien-a003-benchmark.exe" analyze E:\Restaurant_manager --force --json --progress --exclude electron/renderer/src/api/userApi.ts --benchmark-json E:\Anvien\.tmp\child06a_a003_restaurant_manager_frozen_17child_20260826\candidate\benchmark.json --benchmark-label child06a-a003-restaurant-manager-frozen-17child-20260826`
- Launch count / PID / exit: `1 / 9732 / 0`
- Start / end: `08/26/2026 17:14:36 / 08/26/2026 17:16:17`
- Process wall / CPU / user / kernel: `101.096911900 / 110.656250000 / 98.515625000 / 12.140625000 s`

## A002 → A003
| Boundary | A002 before | A003 candidate | Delta |
|---|---:|---:|---:|
| D001 `resolve_calls` | `9.909636600 s` | `9.401585300 s` | `-0.508051300 s` |
| Parent `resolution` | `21.242055400 s` | `20.850792800 s` | `-0.391262600 s` |
| Analyzer | `109.339859600 s` | `98.020546700 s` | `-11.319312900 s` |
| Process | `145.066210900 s` | `101.096911900 s` | `-43.969299000 s` |

## Validation
- Operations / unique: `30 / 30`; denominators equal A002: `PASS`
- Resolution children / unique / mapped: `17 / 17 / 17`; denominators equal A002: `PASS`
- D001 denominator: `calls=86030; files=1234`
- Exclusive intervals / overlap: `3716 / 0`; every interval sum exact: `PASS`
- Parent / child sum / residual: `20.850792800 / 20.834583800 / 0.016209000 s`; conservation: `PASS`
- CPU profile: `56397 bytes`; nonempty: `PASS`
- Files scanned / parsed / failed: `1556 / 1234 / 0`
- Parser bytes: `8546179`; graph/DB rows: `137229 / 190644`
- Dependency/projection counts and resolution semantic counters/diagnostics/outcomes: `PASS`
- Canonical Graph JSON: `PASS; 09708B038AD63B659A95BAC34304A752F2C6350952B91CF435EBD5C5F9D94B7C / 09708B038AD63B659A95BAC34304A752F2C6350952B91CF435EBD5C5F9D94B7C`
- Stdout: `PASS; CD0028F3B3939451FF4D3623DF7DD7482A993E0934D2DFA773682D37D4FB0E94 / CD0028F3B3939451FF4D3623DF7DD7482A993E0934D2DFA773682D37D4FB0E94`
- Stderr: `PASS; 0FDCF15222C02C25771C98E17D9220C89FDCD0294137A6AE4D9B3AAFD4761ADC / 0FDCF15222C02C25771C98E17D9220C89FDCD0294137A6AE4D9B3AAFD4761ADC`
- Ladybug/meta: `lbug 503070720 bytes SHA-256 91FD255BEFC4668255F21F6AB862ABA575D709C099F655ECF5ADCB7C4C335E18; meta 267 bytes SHA-256 717DFD2FE48BADADD3B538DBA09C755C938F1ACE2EA122D03FF5F705F5377FFA; recorded, raw-byte equality non-governing alone`
- Target HEAD/full status unchanged: `PASS`; candidate production identity unchanged: `PASS`

## Resources
- `startAllocBytes`: `1475192 -> 1474592; delta -600`
- `endAllocBytes`: `1119052736 -> 1306049888; delta 186997152`
- `maxObservedSys`: `2911591032 -> 2634607224; delta -276983808`

## Artifacts
- Raw root: `E:\Anvien\.tmp\child06a_a003_restaurant_manager_frozen_17child_20260826`
- Report: `E:\Anvien\reports\Investigation\rp_child06a_a003_restaurant_manager_benchmark_frozen_17child.md`
- Comparison: `E:\Anvien\.tmp\child06a_a003_restaurant_manager_frozen_17child_20260826\candidate\comparison.json`
- Validation: `E:\Anvien\.tmp\child06a_a003_restaurant_manager_frozen_17child_20260826\candidate\validation.json`
- Manifest: `E:\Anvien\.tmp\child06a_a003_restaurant_manager_frozen_17child_20260826\candidate\output-manifest.json`

## Boundary
Measurement-only. No disposition claim, baseline rerun/promotion, architecture/code/build/test, Supervisor, ledger, detect, stage, commit, cleanup, or retry. Next owner: Main.

`MEASUREMENT_A003_RESTAURANT_COMPLETE`
