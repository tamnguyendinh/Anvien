# Child 06A A006-M1 Direct-Callee Attribution

Status: blocked before capture.

Measurement ID: `A006-M1-D001-DIRECT-CALLEE-ATTRIBUTION`

This was a measurement-only overlay attempt for `P2-A / A006-M1 / B1-P1A-OP001 resolution / B2-P2A-A001-D001 resolve_calls`. It was not a production attempt, candidate, baseline rerun, Supervisor event, disposition, or streak event. The sole authorized build invocation failed, so the stop contract prohibited any retry or target launch.

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

Label construction, branching, metric increments, recorder overhead, and other unlisted control work remain residual. Static coverage/non-overlap audit: PASS. Runtime conservation remains `NOT EXPOSED` because the build failed before capture.

## Sole builder invocation

Builder: `E:\Anvien\scripts\build-a00x-benchmark.ps1`; tracked diff empty; SHA-256 `ADA407C7496FCEA988276F03BAD5001ED139A4AEC9A16B9C32947DA440814EC5`.

The executable, adjacent DLL, and provenance were absent immediately before invocation. The builder was invoked exactly once from `E:\Anvien` with:

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

Post-failure state: the `frozen-build` directory, executable, adjacent DLL, and provenance JSON are all absent. Therefore completion marker, executable version/hash, provenance schema/attempt/exits, and exact `2/4/3` provenance input validation are `NOT EXPOSED`. No retry, patch, rebuild, or cleanup occurred.

## Target packets

Cheapapp launch count: `0`. Restaurant Manager launch count: `0`. Capture roots are absent, so there are no process identities, argv/env packets, capture intervals, benchmark JSON, CPU profiles, stdout/stderr, or target output mutations. Sequential/non-overlap proof is `NOT EXPOSED` rather than inferred.

### Cheapapp ten-group rows

| Order | Group | Duration (ns) | Invocation count | Result |
|---:|---|---:|---:|---|
| 1 | `source_context` | NOT EXPOSED | NOT EXPOSED | NOT RUN |
| 2 | `binding_receiver` | NOT EXPOSED | NOT EXPOSED | NOT RUN |
| 3 | `scoped_same_file` | NOT EXPOSED | NOT EXPOSED | NOT RUN |
| 4 | `member_import` | NOT EXPOSED | NOT EXPOSED | NOT RUN |
| 5 | `go_same_package` | NOT EXPOSED | NOT EXPOSED | NOT RUN |
| 6 | `global_lookup` | NOT EXPOSED | NOT EXPOSED | NOT RUN |
| 7 | `typescript_lookup_record` | NOT EXPOSED | NOT EXPOSED | NOT RUN |
| 8 | `evidence_emission` | NOT EXPOSED | NOT EXPOSED | NOT RUN |
| 9 | `direct_site_identity` | NOT EXPOSED | NOT EXPOSED | NOT RUN |
| 10 | `resolve_call_residual` | NOT EXPOSED | NOT EXPOSED | NOT RUN |

Cheapapp `30/30`, `17/17`, denominator, D001/parent conservation, direct-callee conservation, overlap count, ordered Evidence, Graph/stdout/stderr hashes, graph/DB, file/parser, resolution semantics, outcomes/diagnostics, and resources: `NOT EXPOSED`.

### Restaurant Manager ten-group rows

| Order | Group | Duration (ns) | Invocation count | Result |
|---:|---|---:|---:|---|
| 1 | `source_context` | NOT EXPOSED | NOT EXPOSED | NOT RUN |
| 2 | `binding_receiver` | NOT EXPOSED | NOT EXPOSED | NOT RUN |
| 3 | `scoped_same_file` | NOT EXPOSED | NOT EXPOSED | NOT RUN |
| 4 | `member_import` | NOT EXPOSED | NOT EXPOSED | NOT RUN |
| 5 | `go_same_package` | NOT EXPOSED | NOT EXPOSED | NOT RUN |
| 6 | `global_lookup` | NOT EXPOSED | NOT EXPOSED | NOT RUN |
| 7 | `typescript_lookup_record` | NOT EXPOSED | NOT EXPOSED | NOT RUN |
| 8 | `evidence_emission` | NOT EXPOSED | NOT EXPOSED | NOT RUN |
| 9 | `direct_site_identity` | NOT EXPOSED | NOT EXPOSED | NOT RUN |
| 10 | `resolve_call_residual` | NOT EXPOSED | NOT EXPOSED | NOT RUN |

Restaurant Manager `30/30`, `17/17`, denominator, D001/parent conservation, direct-callee conservation, overlap count, ordered Evidence, Graph/stdout/stderr hashes, graph/DB, file/parser, resolution semantics, outcomes/diagnostics, and resources: `NOT EXPOSED`.

## Limitation, checkpoint, and handoff

Instrumentation overhead would have remained attribution-only and could not have been compared with or promoted over accepted A003. Because the sole build failed, no instrumented timing exists at all.

Exact blocker: the unchanged accepted builder invocation terminated with the string-formatting error above and produced no completion identity. The mandatory nonzero-build stop condition is active.

Checkpoint: `A006_ARCHITECT_NEEDS_MEASUREMENT_INPUT / A006_M1_BLOCKED / D001_STREAK_2`. Parent/D001 remain unchecked; D002-D017, P3, and Child 07 remain closed. No plan, ledger, production, test, script, target-source, stage, commit, disposition, or Supervisor mutation occurred.

Next owner: Main Orchestration, to record this blocked packet and decide the next fresh A006 Architect handoff. No measurement rerun is authorized by this report.

`A006_M1_DIRECT_CALLEE_ATTRIBUTION_BLOCKED`
