# A004 Frozen Benchmark Candidate Build

## Result

- Status: A004_FROZEN_CANDIDATE_READY_FOR_MEASUREMENT
- Invocation count: exactly one
- Script exit code: 0
- Terminal completion marker: A00X_BENCHMARK_BUILD_COMPLETE
- Build working directory: E:\Anvien
- The executable was not run and no target analyze, measurement, tests, detect, stage, commit, source/test edit, cleanup, retry, or alternate build was performed.

## Exact invocation inputs

- Script: E:\Anvien\scripts\build-a00x-benchmark.ps1
- AttemptId: A004
- OverlayManifestPath: E:\Anvien\.tmp\child06a_a001_cheapapp_benchmark\optimized-overlay.json
- OutputExecutablePath: E:\Anvien\.tmp\child06a_a004_a00x_benchmark_build\anvien-a004-benchmark.exe
- ExpectedOverlaySha256: 7B138FBA06B41CBD4C6709E6DF00C80C03E4015B615D1E42F787DE2A65D378F7
- ExpectedMappedSourceHash:
  - E:\Anvien\internal\resolution\resolve.go=92A89227E1B1B1C159DE8BE77A1F060361C9058C4F83C78BFC32E9AB40DADEA9
  - E:\Anvien\internal\resolution\types.go=A8DDCB1C61C312D3D3B164B1DD02BFD91C640F1FEB8B6DB511B73FD923C802B8
- ExpectedCandidateSourceHash:
  - internal/graphhealth/diagnostics.go=6DE54D97A1B95E877B686DC3459238598F6060D28CE6329357A80BD7FC376D30
  - internal/resolution/emit.go=73AE20E8A18AC25F8019E2A3223407621AE7ABC358DBAF4F058E1D9AC3136060
  - internal/resolution/outcome.go=02092F9FE7DA2A4BDB49E13056FBA3C97DC24F416141E7F27866EE80F60C1F7E
  - internal/resolution/export_binding_proof.go=36D41BC7336A5A04B4C77D1EE9DEB45DE4E5E2BA120EEFBF98419BD2D846C4D2
- ExpectedNativeHash:
  - lbug.h=3D5114D0863B3DAB3B28BD2FEC97A52E6CF669213739921A01814A5BBF5525EB
  - lbug_shared.lib=B18AAFC0B712DC1C4CB9DD25F76C3828282D7D460627980E3A4B16EFCD98A955
  - lbug_shared.dll=20CBD87840483A2053CFF3FC2DB23A86DD802B8915D86509D41A4B709624CDB7
- GoExecutable: C:\Program Files\Go\bin\go.exe

## Frozen output identities

| Artifact | Path | Bytes | SHA-256 | Version |
|---|---|---:|---|---|
| Executable | E:\Anvien\.tmp\child06a_a004_a00x_benchmark_build\anvien-a004-benchmark.exe | 73,825,280 | 6D319467D198B8BBA2375339CDF6BD7634FA97E7C503D3DFC0C8C315D965352C | 1.2.8 |
| Native DLL | E:\Anvien\.tmp\child06a_a004_a00x_benchmark_build\lbug_shared.dll | 20,230,656 | 20CBD87840483A2053CFF3FC2DB23A86DD802B8915D86509D41A4B709624CDB7 | n/a |
| Provenance | E:\Anvien\.tmp\child06a_a004_a00x_benchmark_build\anvien-a004-benchmark.provenance.json | 12,994 | 4EBF871D0008B6BB1A7FDB433DE07CE437F859DFB34221C94D360E866711A20D | schema 1 |

## Provenance validation

- schemaVersion: 1
- attemptId: A004
- build.exitCode: 0
- top-level exitCode: 0
- Repository HEAD recorded by provenance: 6e961f5a4b8bb5fce0f09c722619519b47503e61
- Overlay manifest expected/actual SHA-256 match:
  - 7B138FBA06B41CBD4C6709E6DF00C80C03E4015B615D1E42F787DE2A65D378F7
- Mapping count: exactly 2. Every expectedReplacementSha256 equals actualReplacementSha256:
  - resolve.go: 92A89227E1B1B1C159DE8BE77A1F060361C9058C4F83C78BFC32E9AB40DADEA9
  - types.go: A8DDCB1C61C312D3D3B164B1DD02BFD91C640F1FEB8B6DB511B73FD923C802B8
- Candidate source count: exactly 4. Every expectedSha256 equals actualSha256:
  - internal/graphhealth/diagnostics.go: 6DE54D97A1B95E877B686DC3459238598F6060D28CE6329357A80BD7FC376D30
  - internal/resolution/emit.go: 73AE20E8A18AC25F8019E2A3223407621AE7ABC358DBAF4F058E1D9AC3136060
  - internal/resolution/outcome.go: 02092F9FE7DA2A4BDB49E13056FBA3C97DC24F416141E7F27866EE80F60C1F7E
  - internal/resolution/export_binding_proof.go: 36D41BC7336A5A04B4C77D1EE9DEB45DE4E5E2BA120EEFBF98419BD2D846C4D2
- Native input count: exactly 3. Every expectedSha256 equals actualSha256:
  - lbug.h: 3D5114D0863B3DAB3B28BD2FEC97A52E6CF669213739921A01814A5BBF5525EB
  - lbug_shared.lib: B18AAFC0B712DC1C4CB9DD25F76C3828282D7D460627980E3A4B16EFCD98A955
  - lbug_shared.dll: 20CBD87840483A2053CFF3FC2DB23A86DD802B8915D86509D41A4B709624CDB7

## Unchanged worktree boundary

- Start HEAD: 6e961f5a4b8bb5fce0f09c722619519b47503e61
- End HEAD: 6e961f5a4b8bb5fce0f09c722619519b47503e61
- Dirty tracked paths at start and end were exactly:
  - internal/resolution/export_binding_proof.go
  - internal/resolution/export_binding_proof_test.go
- Staged set at start and end: empty
- Boundary file SHA-256 values were unchanged:
  - export_binding_proof.go start/end: 36D41BC7336A5A04B4C77D1EE9DEB45DE4E5E2BA120EEFBF98419BD2D846C4D2
  - export_binding_proof_test.go start/end: 99C6F6FC5FD0AE4D1BFDFE547D5F67C958544AA61FFD89895DC0BAC79C839BBC
- Provenance statusAtStart and statusAtEnd both contain only those two modified tracked paths.

## Handoff

Next owner: Main Orchestration. Main opens separate target measurement lanes using the frozen executable, DLL, and provenance recorded above.
