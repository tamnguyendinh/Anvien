# Child 06A A002 Restaurant Manager Frozen Measurement

Status: `MEASUREMENT_COMPLETE`

## Execution
- Task: `01a03c72-5f6d-7fd0-948a-88b06a1ebb65`
- Target: `E:\Restaurant_manager`
- Command: `& "E:\Anvien\.tmp\child06a_a002_a00x_benchmark_build\anvien-a002-benchmark.exe" analyze E:\Restaurant_manager --force --json --progress --exclude electron/renderer/src/api/userApi.ts --benchmark-json E:\Anvien\.tmp\child06a_a002_restaurant_manager_frozen_17child_20260826\candidate\benchmark.json --benchmark-label child06a-a002-restaurant-manager-frozen-17child-20260826`
- Launch count: `1`
- PID: `15688`
- Start / end: `2026-08-26T05:12:53.9496932+00:00 / 2026-08-26T05:15:19.0185984+00:00`
- Exit: `0`
- Process wall / CPU: `145.066210900 / 114.203125000 s`

## Before / After
| Boundary | A001 before | A002 candidate | Delta |
|---|---:|---:|---:|
| D001 `resolve_calls` | `40.769294200 s` | `9.909636600 s` | `-30.859657600 s` |
| Parent `resolution` | `136.436879300 s` | `21.242055400 s` | `-115.194823900 s` |
| Analyzer | `215.972455200 s` | `109.339859600 s` | `-106.632595600 s` |
| Process | `218.680628900 s` | `145.066210900 s` | `-73.614418000 s` |

## Required Packet Checks
- Operations: `30/30`
- Resolution children: `17/17`, exact order: `PASS`
- D001 denominator: `calls=86030; files=1234`
- Exclusive intervals / overlap: `3716 / 0`
- Parent / child sum / residual: `21.242055400 / 21.217471500 / 0.024583900 s`
- Conservation: `PASS`
- CPU profile: `59197 bytes; nonempty=true`

## Equivalence
- Files scanned / parsed / failed: `1556 / 1234 / 0`
- Graph nodes / relationships: `137229 / 190644`
- DB nodes / relationships: `137229 / 190644`
- Operation/child denominators: `PASS`
- Resolution semantic counters / diagnostics / outcomes: `PASS`
- Graph JSON: `PASS; 09708B038AD63B659A95BAC34304A752F2C6350952B91CF435EBD5C5F9D94B7C / 09708B038AD63B659A95BAC34304A752F2C6350952B91CF435EBD5C5F9D94B7C`
- Stdout: `PASS; CD0028F3B3939451FF4D3623DF7DD7482A993E0934D2DFA773682D37D4FB0E94 / CD0028F3B3939451FF4D3623DF7DD7482A993E0934D2DFA773682D37D4FB0E94`
- Stderr: `PASS; 0FDCF15222C02C25771C98E17D9220C89FDCD0294137A6AE4D9B3AAFD4761ADC / 0FDCF15222C02C25771C98E17D9220C89FDCD0294137A6AE4D9B3AAFD4761ADC`
- Ladybug/meta: `lbug 503070720 bytes SHA-256 5CDE4541F687E515CF04C2824F89EE9E6C701B38EE7D090AB75438B9572E1E48; meta 267 bytes SHA-256 B45157FC1B230FE1522E0A395226997ED071534F4BC73DDED8C870F5CE03078A; recorded, non-governing alone`
- Anvien HEAD/status unchanged: `PASS`
- Restaurant HEAD/status unchanged: `PASS`

## Resources
- `startAllocBytes`: `1461312 -> 1475192; delta +13880`
- `endAllocBytes`: `876051256 -> 1119052736; delta +243001480`
- `maxObservedSys`: `3050756728 -> 2911591032; delta -139165696`
- Structural appender boundary: `O(T)` slice-header map, shared graph/cache slice, no duplicate retained diagnostic objects.

## Artifacts
- Raw root: `E:\Anvien\.tmp\child06a_a002_restaurant_manager_frozen_17child_20260826`
- Report: `E:\Anvien\reports\Investigation\rp_child06a_a002_restaurant_manager_benchmark_frozen_17child.md`
- Benchmark SHA-256: `C2E94CD1939C8A0435A970D00341EDF4F366E9418B53284058D92355195E51FF`
- Process SHA-256: `2045359544842E3E110C32B1E93CDB19BAFA115F4B64FE4B16C77CF3566B7E16`
- Graph SHA-256: `09708B038AD63B659A95BAC34304A752F2C6350952B91CF435EBD5C5F9D94B7C`
- Stdout SHA-256: `CD0028F3B3939451FF4D3623DF7DD7482A993E0934D2DFA773682D37D4FB0E94`
- Stderr SHA-256: `0FDCF15222C02C25771C98E17D9220C89FDCD0294137A6AE4D9B3AAFD4761ADC`
- Profile SHA-256: `97F0BB7E24D5DF9585F0B0221B352E63B2E603E26BE9CB959371885CE646F4DB`

## Boundary
Measurement-only. No KEEP/REWORK/ROLLBACK, baseline promotion, checklist/streak change, Supervisor, detect, stage, commit, build, test, or source edit.

`MEASUREMENT_COMPLETE`
