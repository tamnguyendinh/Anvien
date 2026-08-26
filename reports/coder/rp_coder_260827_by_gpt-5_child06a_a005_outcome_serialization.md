# Coder Report — Child 06A A005 Outcome Serialization

Verdict: `CODER_A005_IMPLEMENTATION_VALIDATED_READY_FOR_MAIN_HANDOFF`

## Metadata

- Date: `2026-08-27`
- Role: sole visible A005 Coder
- Slice: `P2-A / A005 / B1-P1A-OP001 resolution / B2-P2A-A001-D001 resolve_calls`
- Repository: `E:\Anvien`, shared saved-project checkout; no worktree, fork, branch, alternate checkout, stash, reset, checkout, or broad cleanup
- Required and actual base HEAD: `cf230eeee87e2cc212c42639c8c7e99eba084bc7` (`docs(plan): verify Child 06A A005 execution plan`)
- Initial worktree entries: `0`
- Initial staged entries: `0`
- Accepted basis preserved: A003 target-separated values; A004 remains `NO_KEEP / ROLLBACK_COMPLETE`; D001 streak remains `1`; D002-D017, P3, and Child 07 remain closed
- Authority: `E2-P2A-A005ATTRIB1/ARCH1/PLAN1/MAINVERIFY1`, the A005 Architect report, visible Planner report, and Main Planner-handoff `PASS`
- Ownership boundary: production first, then exactly the authorized tests, canonical full build before every test execution, focused/package validation, four-path final diff audit, and this report only
- Explicit non-ownership: no target measurement, frozen candidate/overlay work, Supervisor verdict, Main disposition, ledger edit, detect-changes, staging, commit, P3, or Child 07 action

## Outcome

A005 now has one semantic authority and one canonical byte authority per retained `SourceSiteID`:

```text
clone
-> validate
-> conflict/equal-duplicate decision
-> unique semantic retention
-> unchanged record-time marshal once
-> private SourceSiteID -> canonical JSON string sidecar
-> private finalized values-plus-sidecar bundle
-> immediate diagnostic and final relationship/reference/diagnostic parity consumers
```

`projectResolutionOutcomes` is now a strict consumer and validator. It does not call `marshalResolutionOutcome`, regenerate bytes, repair missing coverage, tolerate an extra payload, or create a second serialized shape. `ResolveBoundInto` passes the full private bundle to projection and unwraps only the sorted semantic values into public `Result.ResolutionOutcomes`.

## Invariant Family Map

- Family: A005 canonical outcome-byte ownership.
- Authority / SSOT: `E2-P2A-A005ATTRIB1/ARCH1/PLAN1/MAINVERIFY1`, `reports/system-architect/rp_system-architect_260827_by_gpt-5_child06a_a005_outcome_serialization.md`, `reports/Planner/rp_planner_260827_by_gpt-5_child06a_a005_outcome_serialization.md`, and `reports/Supervisor/rp_supervisor_260827_025255_by_gpt-5_child06a_a005_planner_handoff.md`.
- Primary semantic surface: `resolutionOutcomeCollector.bySourceSite` and its existing clone/validation/conflict/first-error behavior.
- Primary byte surface: new private `encodedBySourceSite`, populated only after the unchanged record-time marshaler succeeds.
- Finalized boundary: private `finalizedResolutionOutcomes`, SourceSiteID-sorted cloned values plus the run-scoped sidecar.
- Sibling runtime surfaces checked: repository resolved/unresolved, intrinsic resolved, TypeScript resolved-external/capability-unavailable/profile-excluded/meaning-mismatch, equal duplicates, conflicts, record-time marshal failure, immediate diagnostics, graph relationships, both reference-index maps, diagnostic parity, public result carriage, Graph JSON/graphhealth/analyze/Ladybug loading/native readback regressions.
- Forbidden legacy fallback: projection-time marshal, missing-byte repair, status-specific encoder, alternate serialized shape, public/persisted cached field, payload copy per carrier, or cross-run state. Status: absent.
- Stale tests/helpers/plans: only the existing P6C3 direct-private-call test needed mechanical bundle adaptation. No helper, plan, ledger, golden owner, A001-A004 owner, or excluded consumer was edited.
- Verify matrix: first add; equal duplicate; conflict; invalid/marshal error; missing/extra coverage; relationship/reference/diagnostic carriers; traversal/order; replay; focused cross-package regressions; canonical full build.
- Residual unverified surfaces inside the Coder boundary: none. Target measurements, later A005 Supervisor, and disposition are deliberately owned by Main’s subsequent chain.

## Mandatory Pre-Edit Gate

### Clean start

- HEAD: `cf230eeee87e2cc212c42639c8c7e99eba084bc7`, exact match.
- Worktree: clean, `0` entries.
- Staged set: empty, `0` entries.
- `anvien --help`: exit `0`.
- Exactly one Coder-owned fresh `anvien analyze --force`: exit `0`; `2,254` scanned, `766` parsed code, `0` failed; graph `124,251` nodes / `171,117` relationships.

The later canonical full build ran its own mandatory maintenance analyze; it is build validation, not another Coder pre-edit graph refresh or product benchmark.

### File-detail

| File | Risk | Symbols | Inbound | Outbound | Local relationships | Unresolved | Linked flows | Linked tests | Current/stale |
|---|---|---:|---:|---:|---:|---:|---:|---:|---|
| `internal/resolution/outcome.go` | HIGH | 119 | 87 | 122 | 117 | 148 | 11 | 23 | current / false |
| `internal/resolution/resolve.go` | HIGH | 200 | 134 | 289 | 87 | 401 | 26 | 40 | current / false |

Both file-detail records resolved indexed/current commit `cf230eeee87e2cc212c42639c8c7e99eba084bc7` and `changedSinceAnalyze=false`.

### Exact upstream symbol impact

| Symbol | Risk | Impacted symbols | Direct | Files | Modules | Processes | Functional areas |
|---|---|---:|---:|---:|---:|---:|---|
| `resolutionOutcomeCollector` | CRITICAL | 33 | 5 | 5 | 2 | 49 | analyzer 1, resolution 32 |
| `newResolutionOutcomeCollector` | CRITICAL | 4 | 1 | 3 | 2 | 23 | analyzer 1, resolution 3 |
| `(*resolutionOutcomeCollector).record` | CRITICAL | 4 | 4 | 1 | 1 | 10 | resolution 4 |
| `(*resolutionOutcomeCollector).finalize` | CRITICAL | 6 | 1 | 4 | 4 | 32 | analyzer 1, CLI 1, graph health 1, resolution 3 |
| `projectResolutionOutcomes` | CRITICAL | 6 | 1 | 4 | 4 | 32 | analyzer 1, CLI 1, graph health 1, resolution 3 |
| `ResolveBoundInto` | CRITICAL | 7 | 2 | 5 | 5 | 31 | analyzer 1, CLI 3, graph health 1, resolution 2 |

Name-only method resolution was ambiguous, so `record` and `finalize` were rebound through the exact current file-detail UIDs before impact. The new private `finalizedResolutionOutcomes` representation was queried after creation against the single pre-edit graph and was not addressable (`UNKNOWN`, impacted count `0`); no second graph refresh was run. HIGH/CRITICAL were treated as blast-radius warnings and the edit remained inside the exact two-file boundary.

## Production-First Implementation

### `internal/resolution/outcome.go`

- `resolutionOutcomeCollector` at line 74 adds only `encodedBySourceSite map[string]string`.
- `newResolutionOutcomeCollector` at line 80 initializes only the existing semantic map and the new sidecar.
- `record` at line 87 preserves its signature and clone/validation/retention/first-error/return behavior:
  - a unique valid outcome is semantically retained before the unchanged record-time marshal;
  - canonical bytes are stored only on marshal success;
  - an equal duplicate returns the prior immutable clone/string with `added=false` and no second marshal;
  - a conflicting duplicate returns the prior tuple, preserves the first error, and does not encode the conflict;
  - a prior marshal failure stays record-time, retains empty bytes, returns `added=false`, is never retried, and fails finalization through the first error.
- `finalizedResolutionOutcomes` at line 115 is private and contains sorted semantic values plus the run-scoped sidecar.
- `finalize` at line 120 preserves first-error precedence, semantic revalidation, deep cloning, SourceSiteID sorting, and fails closed on missing or extra sidecar coverage.
- `projectResolutionOutcomes` at line 416 consumes the bundle, preserves duplicate finalized-site, resolved/non-resolved, relationship, reference-index, overlap, and diagnostic parity checks, and contains no encoder/fallback.
- `marshalResolutionOutcome`, `validateResolutionOutcome`, `cloneResolutionOutcome`, outcome constructors, diagnostic callers, reference helpers, and diagnostic-site validation bodies were not edited.

### `internal/resolution/resolve.go`

- Only the existing `ResolveBoundInto` finalize/project/result block changed.
- Projection receives the private bundle.
- Error and success results expose only `outcomes.values` as public `ResolutionOutcomes`.
- Exported signature, loop/branch order, authority processing, metrics, graph/reference ownership, and every other line remain unchanged.

### Production diff gate before tests

- Changed production paths: exactly `internal/resolution/outcome.go` and `internal/resolution/resolve.go`.
- Production numstat at that gate: `outcome.go +43/-24`; `resolve.go +2/-2`.
- Scoped `git diff --check`: exit `0`.
- `gofmt -d`: empty.
- No additional production owner was needed, so tests were authorized only after this gate.

## Authorized Tests

### New A005 tests

All new names use the mandated `TestA005OutcomeSerialization` prefix.

| Test | File:line | Proof | Focused result |
|---|---|---|---|
| `TestA005OutcomeSerializationStatusFamiliesUseCanonicalRecordBytes` | `internal/resolution/outcome_serialization_test.go:17` | all seven status families, pre-A005 JSON byte identity, first-add/equal-duplicate tuple reuse, one-to-one finalization | PASS |
| `TestA005OutcomeSerializationImmediateDiagnosticsReuseCanonicalBytesWithoutDuplicates` | `internal/resolution/outcome_serialization_test.go:75` | repository and TypeScript non-resolved immediate diagnostic parity; duplicate stable sites emit one outcome/diagnostic | PASS |
| `TestA005OutcomeSerializationDuplicateConflictFirstErrorAndMutationIsolation` | `internal/resolution/outcome_serialization_test.go:137` | equal-duplicate sidecar reuse without marshal, conflict/prior tuple, first error, input/returned/finalized/target/proof/authority/declaration deep isolation | PASS |
| `TestA005OutcomeSerializationProjectionParityTraversalOrderingAndDeterminism` | `internal/resolution/outcome_serialization_test.go:225` | SourceSiteID order, relationship and both reference-index byte parity, unchanged traversal, complete ordered evidence, diagnostic parity, replay determinism | PASS |
| `TestA005OutcomeSerializationFailClosedCoverageMarshalAndCarrierDrift` | `internal/resolution/outcome_serialization_test.go:274` | record-time non-finite `Evidence.Weight` marshal failure, no diagnostic/no bytes, missing/extra coverage without repair, duplicate finalized sites, resolved/unresolved overlap, diagnostic drift, non-resolved/missing reference carrier rejection | PASS |

### Existing P6C3 adaptation

`internal/resolution/p6c3_structured_outcome_test.go::TestP6C3P5ProofNestingConflictAndImmutableReplay` retains all assertions. Only its three direct `finalize` result uses and one direct `projectResolutionOutcomes` construction were mechanically adapted to the private bundle. Numstat is `+8/-5`; every other test byte remains unchanged.

## Build-Holder Preflight And Canonical Full Build

### Exact preflight

1. `anvien doctor locks --repo E:\Anvien --json`
   - exit `0`;
   - `status=free`;
   - analyze lock file absent;
   - no recovery/termination required.
2. `anvien doctor processes --json`
   - exit `0`;
   - listed editor-owned global-npm Anvien MCP processes and one VS Code-owned Playwright test server;
   - none was proven to hold `E:\Anvien` canonical outputs or build locks;
   - terminated processes: none.

### Exact canonical build

Command:

```text
pwsh -NoLogo -NoProfile -NonInteractive -ExecutionPolicy Bypass -File E:\Anvien\scripts\full-build.ps1
```

Result: exit `0`.

What it proves:

- canonical vendor verification PASS: `1,798` files, `0` missing, `0` extra, `0` mismatched;
- Go runtime build PASS with pinned Ladybug native authority;
- launcher/server/Vite asset build PASS; only the known chunk-size and mixed static/dynamic-import warnings were emitted;
- canonical runtime version `1.2.8`, SHA-256 `CFDDA72DBB448AB9BE20FC58A7C20A883353E7F0B3EE7DF60A6DDD0968163AC7`;
- required maintenance analyze PASS with `2,255` scanned, `767` parsed, `0` failed, graph `124,463 / 171,506`.

The build boundary is validation evidence only and is not recorded as product benchmark data.

## Post-Build Validation

### Exact focused command

```text
go test ./internal/resolution -count=1 -run '^(TestA005OutcomeSerialization.*|TestP6C3P5ProofNestingConflictAndImmutableReplay|TestP6C3FinalOutcomeStatusPrecedenceAndCarriageMatrix|TestResolveAttachesSourceBackedUnresolvedDiagnostics|TestP6BResolverCarriesCatalogValidationFailuresPerSourceSite|TestP6DStructuredOutcomeProjectionPreservesGraphJSONAndHealthParity|TestP6C3AnalyzeResultPreservesFinalOutcomesAndGraphCarriage|TestP6C3AnalyzeCapabilityOutcomesRetainAcceptedAuthorityStatus|TestP6C3ResolutionOutcomesPreserveGraphJSONAndLadybugParity|TestP6C3NativeLadybugResolutionOutcomeReadback)$'
```

Result: exit `0`, PASS.

This proves the complete new A005 matrix and the named P6B/P6C3/P6D/Graph JSON/Ladybug carriage regressions at their focused real boundary.

### Separate package commands

| Exact command | Exit | Result |
|---|---:|---|
| `go test ./internal/resolution -count=1` | 1 | Truthful known preserve-only FAIL only at `TestProofBasedCallAccessGoldenCorpus`; package is not called PASS |
| `go test ./internal/graphhealth -count=1` | 0 | PASS |
| `go test ./internal/analyze -count=1` | 0 | PASS |
| `go test ./internal/lbugload -count=1` | 0 | PASS |
| `go test ./internal/lbugnative -count=1` | 0 | PASS |

The sole resolution failure payload is unchanged: source site `SourceSite:src/golden.ts#call#callback#14#2#14#12`, classification/actionability `unclassified/review`, and the same structured resolution-outcome `Note`. The preserve-only golden owner is byte-identical to HEAD: current and HEAD Git blob are both `00d0ae042f109cd7d3c8aeadf855d192a5aa8172`. No golden file was edited. No new or changed failure exists.

## Final Scope And Identity Audit

Final source/test scope before this report:

| Path | Numstat | SHA-256 |
|---|---:|---|
| `internal/resolution/outcome.go` | `+43/-24` | `18203DFAB9A227B526F8F7478B516AE6673F635BABC02D9463975E428A3983AF` |
| `internal/resolution/resolve.go` | `+2/-2` | `76B7B62A060B36EE2438E76689E858544358AD681DB42F6A4FC47D271F1749A1` |
| `internal/resolution/outcome_serialization_test.go` | `+652/-0` | `89B168B2764B9A9B8EACDAB37A071BE0E1C51A8F791FE158B111651E9640C957` |
| `internal/resolution/p6c3_structured_outcome_test.go` | `+8/-5` | `E561280B3F8420D2288001179431A37FC39E6E2B8564863C9AD644EEE005A2E6` |

Total source/test numstat: `+705/-31`.

- Exact four-path `git diff --check`: exit `0`.
- `gofmt -d` across all four paths: empty.
- Source/test changed-path count: exactly `4`.
- Production/test path outside the authorized boundary: none.
- Staged set: empty, `0` entries.
- Detect-changes: not run, by explicit Coder boundary.
- Stage/commit: not performed, by explicit Coder boundary.
- Ledger/benchmark/actual-status edit: none.
- Target measurement or frozen A005 build: none.

## E2E Verification

```text
[PASS] Compiled: canonical full build -> exit 0
[PASS] Runtime boundary: maintenance analyze inside canonical build -> 2255 scanned / 767 parsed / 0 failed
[PASS] Happy path: record -> canonical sidecar -> finalized bundle -> relationship + both reference indexes + public values
[PASS] Edge cases: equal duplicate, conflict, first error, deep mutation, non-finite marshal, missing/extra coverage, duplicate site, overlap, diagnostic/reference drift
[PASS] Cross-surface: focused P6B/P6C3/P6D/Graph JSON/Ladybug regressions and four sibling packages
[KNOWN PRESERVE-ONLY FAIL] internal/resolution full package -> only unchanged TestProofBasedCallAccessGoldenCorpus; package not labeled PASS
```

## Risks And Open Points

- Fresh graph blast radius is HIGH/CRITICAL and broad. The implementation remains deliberately private, two-file, run-scoped, and fail-closed; broad focused/package regressions passed within the authorized outcome/graph/reference/reader boundary.
- Retained A005 state is `O(U+B)`: one semantic map already existed; the new sidecar holds one map entry and one canonical string backing payload per retained source site; carrier assignments copy string headers only. No global state, I/O, goroutine, lock, concurrency, flush, finalizer, TTL, or public/persisted field was introduced.
- The known resolution golden remains a truthful unrelated validation failure and is not A005 evidence or authorization to modify its owner.
- Performance eligibility is intentionally unresolved here. Main owns frozen candidate generation, target-separated measurements, later A005 Supervisor review, and disposition.

## Handoff

- Coder state: implementation, authorized tests, canonical full build, focused validation, allowed package outcomes, scope audit, hashes, and report complete.
- Next owner: Main Orchestration.
- Main must independently verify this source/test/report boundary before any frozen-candidate or measurement action.
- No Supervisor verdict, Main disposition, detect, stage, or commit is claimed.

`CODER_A005_IMPLEMENTATION_VALIDATED_READY_FOR_MAIN_HANDOFF`
