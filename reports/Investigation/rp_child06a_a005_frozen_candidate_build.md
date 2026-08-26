# A005 frozen candidate recovery build

## Result

`A005_FROZEN_CANDIDATE_READY_FOR_MEASUREMENT`

The corrected recovery patched only the copied overlay, invoked the unchanged A00X benchmark build script exactly once, and produced a validated frozen candidate under the new `retry1` root. The executable was not run.

## Historical handoff record

The following history is recorded from the Main handoff and was intentionally not re-audited in this recovery lane:

- Preparation stopped before any build invocation (`zero-build preparation stop`).
- The original invocation 1 reached compilation and failed; its root remains `E:\Anvien\.tmp\child06a_a005_a00x_benchmark_build`.
- The prior recovery attempt then stopped with zero action before patch/build because Main supplied the wrong long-form HEAD suffix. No repair or investigation was performed in that stopped attempt.
- Corrected HEAD authority for this recovery was `6ccc5134055b4dff5f80c925fee89ea4ae456949`.

## Recovery preconditions

The one permitted precondition check passed:

- HEAD matched `6ccc5134055b4dff5f80c925fee89ea4ae456949`.
- The worktree contained exactly the four expected A005 source/test paths and the index was empty.
- All four source/test SHA-256 values matched the frozen handoff values.
- Copied overlay resolve SHA-256 was `90FAA94983BC2BAB6CBF26D88F9DE9E8AA56CCE34E9246EEB27763D1571BA13E` before recovery.
- Copied types SHA-256 was `A8DDCB1C61C312D3D3B164B1DD02BFD91C640F1FEB8B6DB511B73FD923C802B8`.
- Manifest SHA-256 was `FD555BCE64FB63C81526614B123CAEAB5CD96E7250C554F0833DD405B1F2EA53`.
- The failed root existed and was treated as preserve-only.
- The new `E:\Anvien\.tmp\child06a_a005_a00x_benchmark_build_retry1` root did not exist.

## Exact copied-overlay recovery

One `apply_patch` invocation changed exactly three uses in `E:\Anvien\.tmp\child06a_a005_overlay\optimized-overlay\internal\resolution\resolve.go`:

1. First projection call: `resolutionOutcomeProjectionDenominators(g, e.referenceIndex, outcomes)` became `resolutionOutcomeProjectionDenominators(g, e.referenceIndex, outcomes.values)`.
2. Second projection call: `resolutionOutcomeProjectionDenominators(g, e.referenceIndex, outcomes)` became `resolutionOutcomeProjectionDenominators(g, e.referenceIndex, outcomes.values)`.
3. Result denominator: `int64(len(outcomes))` became `int64(len(outcomes.values))` on the exact `resolutionChildAssembleResolutionResult` line.

Post-patch verification passed without a second patch:

- Resolve SHA-256: `304F65E40629AE9B32803BA3D61ECD093022481381C948D8CF0B09AD75BFF788`.
- Old/new projection-call counts: `0/2`.
- Old/new exact result-count-line counts: `0/1`.
- Helper-local `outcomeCount := int64(len(outcomes))` count: `1`.
- Copied types and manifest hashes remained unchanged.

## Single recovery build invocation

The unchanged `E:\Anvien\scripts\build-a00x-benchmark.ps1` was invoked exactly once from `E:\Anvien` with:

- Attempt: `A005`.
- Overlay manifest: `E:\Anvien\.tmp\child06a_a005_overlay\optimized-overlay.json`.
- Output executable: `E:\Anvien\.tmp\child06a_a005_a00x_benchmark_build_retry1\anvien-a005-benchmark.exe`.
- Expected overlay SHA-256: `FD555BCE64FB63C81526614B123CAEAB5CD96E7250C554F0833DD405B1F2EA53`.
- Expected mapped source hashes: resolve `304F65E40629AE9B32803BA3D61ECD093022481381C948D8CF0B09AD75BFF788`; types `A8DDCB1C61C312D3D3B164B1DD02BFD91C640F1FEB8B6DB511B73FD923C802B8`.
- Expected candidate source count/hash set: `4`, exactly as validated below.
- Expected native input count/hash set: `3`, exactly as validated below.
- Go executable: `C:\Program Files\Go\bin\go.exe`.

The invocation exited `0` and emitted `A00X_BENCHMARK_BUILD_COMPLETE`. Provenance timestamps were `2026-08-26T21:10:56.0917828Z` through `2026-08-26T21:13:06.7588728Z`. Compilation emitted the existing tree-sitter-swift shift-count warning; it did not change the successful exit or completion marker. There was no retry.

## Output packet

| Artifact | Bytes | SHA-256 | Version |
|---|---:|---|---|
| `E:\Anvien\.tmp\child06a_a005_a00x_benchmark_build_retry1\anvien-a005-benchmark.exe` | 73,825,792 | `0178DAC396E0DD4D0BECCE7B803378A3400013F0B508EC95F9221A0020CF96F3` | `1.2.8` |
| `E:\Anvien\.tmp\child06a_a005_a00x_benchmark_build_retry1\lbug_shared.dll` | 20,230,656 | `20CBD87840483A2053CFF3FC2DB23A86DD802B8915D86509D41A4B709624CDB7` | n/a |
| `E:\Anvien\.tmp\child06a_a005_a00x_benchmark_build_retry1\anvien-a005-benchmark.provenance.json` | 13,322 | `E4B9C7E9F9850A58356D41131C7CC177CFC6207D7D8CD3E158B633524B56244E` | schema `1` |

Provenance validation passed with schema/attempt/build-exit/top-level-exit `1/A005/0/0` and exact mapped/candidate/native counts `2/4/3`.

### Overlay hash pairs

| Input | Expected SHA-256 | Actual SHA-256 |
|---|---|---|
| Overlay manifest | `FD555BCE64FB63C81526614B123CAEAB5CD96E7250C554F0833DD405B1F2EA53` | `FD555BCE64FB63C81526614B123CAEAB5CD96E7250C554F0833DD405B1F2EA53` |
| `internal/resolution/resolve.go` replacement | `304F65E40629AE9B32803BA3D61ECD093022481381C948D8CF0B09AD75BFF788` | `304F65E40629AE9B32803BA3D61ECD093022481381C948D8CF0B09AD75BFF788` |
| `internal/resolution/types.go` replacement | `A8DDCB1C61C312D3D3B164B1DD02BFD91C640F1FEB8B6DB511B73FD923C802B8` | `A8DDCB1C61C312D3D3B164B1DD02BFD91C640F1FEB8B6DB511B73FD923C802B8` |

### Candidate source hash pairs

| Input | Expected/actual SHA-256 |
|---|---|
| `internal/graphhealth/diagnostics.go` | `6DE54D97A1B95E877B686DC3459238598F6060D28CE6329357A80BD7FC376D30` |
| `internal/resolution/emit.go` | `73AE20E8A18AC25F8019E2A3223407621AE7ABC358DBAF4F058E1D9AC3136060` |
| `internal/resolution/outcome.go` | `18203DFAB9A227B526F8F7478B516AE6673F635BABC02D9463975E428A3983AF` |
| `internal/resolution/export_binding_proof.go` | `4BFE1976766FBF2E2257102070FF43CAC6D757E32E71DD583F8824781EAB6A8E` |

### Native input hash pairs

| Input | Expected/actual SHA-256 |
|---|---|
| `lbug.h` | `3D5114D0863B3DAB3B28BD2FEC97A52E6CF669213739921A01814A5BBF5525EB` |
| `lbug_shared.lib` | `B18AAFC0B712DC1C4CB9DD25F76C3828282D7D460627980E3A4B16EFCD98A955` |
| `lbug_shared.dll` | `20CBD87840483A2053CFF3FC2DB23A86DD802B8915D86509D41A4B709624CDB7` |

## Frozen source/test boundary after build

| Path | SHA-256 |
|---|---|
| `internal/resolution/outcome.go` | `18203DFAB9A227B526F8F7478B516AE6673F635BABC02D9463975E428A3983AF` |
| `internal/resolution/resolve.go` | `76B7B62A060B36EE2438E76689E858544358AD681DB42F6A4FC47D271F1749A1` |
| `internal/resolution/p6c3_structured_outcome_test.go` | `E561280B3F8420D2288001179431A37FC39E6E2B8564863C9AD644EEE005A2E6` |
| `internal/resolution/outcome_serialization_test.go` | `89B168B2764B9A9B8EACDAB37A071BE0E1C51A8F791FE158B111651E9640C957` |

The worktree still contained exactly those four paths and staged count remained `0`.

## Boundary and next owner

- The failed root `E:\Anvien\.tmp\child06a_a005_a00x_benchmark_build` was preserved without cleanup.
- This lane did not edit repository source/tests, the build script, build interfaces, plans, or prior reports.
- This lane did not run the produced executable and did not run target analyze, tests, graph operations, detect-changes, staging, commits, cleanup, alternate builds, or retries.
- The only code-like mutation was the exact three-use copied-overlay patch above; the only repository report mutation was this file.
- Next owner: Main verifies the packet before opening target measurement lanes.
