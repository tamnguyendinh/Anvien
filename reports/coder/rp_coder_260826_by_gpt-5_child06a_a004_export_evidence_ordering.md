# Coder Report — Child 06A A004 Export-Binding Evidence Ordering

Date: `2026-08-26`
Role: sole visible Coder
Scope: `P2-A / A004 / B1-P1A-OP001 resolution / B2-P2A-A001-D001 resolve_calls`
Verdict: `CODER_A004_READY_FOR_MEASUREMENT_HANDOFF`
Base HEAD: `4695f7c706c5c627d06c60e12d1b3c2217738599` (`docs(plan): authorize Child 06A A004 ordering attempt`)
Candidate state: uncommitted, unstaged, and limited to the two authorized source/test paths plus this report

## Invariant Family Map

- Family: export-binding evidence projection, exact-tuple coalescing, and canonical deterministic ordering.
- Authority / SSOT: `E2-P2A-A004ARCH1` and `E2-P2A-A004PLAN1`; Architect report `reports/system-architect/rp_system-architect_260826_by_gpt-5_child06a_a004_residual_direction.md`; Planner report `reports/Planner/rp_planner_260826_by_gpt-5_child06a_a004_export_evidence_ordering.md`.
- Exact production owner: `internal/resolution/export_binding_proof.go`, limited to `appendExportBindingEvidence`, `mergeExportBindingEvidence`, `sortExportBindingEvidence`, `exportBindingEvidenceOrderFor`, existing private key machinery, and one private transient decorated type.
- Sibling runtime surfaces checked without edits: `resolveCall`, sibling `resolveAccess`, resolved relationship coalescing through `mergeRelationship`, resolution outcome/reference-index projection, graph-health projection, analyzer result carriage, and Graph JSON carriage.
- Forbidden alternate/fallback path: no producer-carried typed order key, private overload, second key authority, comparator JSON decode, second final projected sort, public/persisted order field, retained cache, concurrency, I/O, or lifecycle owner was introduced.
- Stale validation boundary: `TestProofBasedCallAccessGoldenCorpus` remains the already-recorded preserve-only golden failure and was not edited or called PASS.
- Verify matrix: full ordered oracle parity; terminal/hop/failure across multiple proofs/hops; generic/projected mixing; permutations; repeated coalescing; exact duplicates; same ordinals/different Notes; multiple SourceSiteIDs; malformed JSON for all three export kinds; unknown/non-export kinds; equal-key stability; empty/no-proof; caller-input non-mutation/aliasing; idempotency; call/access/outcome/Graph JSON and A003 diagnostic parity.
- Residual unverified surfaces: none inside the Coder source/build/test boundary. Target timing/resource measurement, candidate packaging, Supervisor review, disposition, detect, stage, and commit belong to later owners.

## Fresh Anvien Pre-Edit Gate

Execution order and result:

1. `anvien --help` — PASS, exit `0`.
2. Base HEAD verification — `4695f7c706c5c627d06c60e12d1b3c2217738599`.
3. Exactly one fresh pre-edit `anvien analyze --force` — PASS, exit `0`; `2243` scanned / `766` parsed code / `0` failed; graph `124096` nodes / `170962` relationships.
4. `anvien file-detail internal/resolution/export_binding_proof.go --repo E:\Anvien --json` — PASS; indexed/current commit both `4695f7c706c5c627d06c60e12d1b3c2217738599`; `stale=false`; `changedSinceAnalyze=false`; file risk HIGH; `98` symbols; fan-in `68`; fan-out `48`; `70` local relationships; `101` unresolved sites; `1` linked flow; `23` linked tests.
5. Exact upstream helper impacts — all PASS and CRITICAL. Fresh current counts from the mandated CLI commands are:

| Helper | Impacted symbols | Direct | Affected files | Modules | Processes |
|---|---:|---:|---:|---:|---:|
| `appendExportBindingEvidence` | `5` | `2` | `2` | `2` | `30` |
| `mergeExportBindingEvidence` | `16` | `4` | `6` | `2` | `34` |
| `sortExportBindingEvidence` | `9` | `2` | `4` | `1` | `19` |
| `exportBindingEvidenceOrderFor` | `8` | `1` | `4` | `1` | `8` |

The HIGH/CRITICAL results were treated as blast-radius warnings. Current graph/source identity was fresh, and the exact allowed one-file production boundary was sufficient, so no STOP condition applied.

## Production-First Implementation

Production was changed before tests.

- Removed the redundant projected-evidence pre-sort from `appendExportBindingEvidence`.
- Preserved `mergeExportBindingEvidence` exact-tuple dedupe by `(Kind, Weight, Note)`, existing-first/incoming-second traversal, and generic/projected partitioning before decoration.
- Added private transient `orderedExportBindingEvidence`, which copies a `graph.Evidence` value/string headers and its existing canonical `exportBindingEvidenceOrder` tuple; it does not copy retained Note bytes or survive the call.
- `sortExportBindingEvidence` now decorates each unique projected record exactly once by calling unchanged `exportBindingEvidenceOrderFor` once in a linear loop.
- The one stable comparator reads only cached `kindRank`, `proofOrdinal`, `hopOrdinal`, and `note`; it performs no JSON decode, marshal, validation, or mutation.
- After the single stable sort, unchanged evidence values are copied back and transient keys are discarded.
- The extractor, public `graph.Evidence`, and every existing function signature remain unchanged.

Static final-source proof:

- exactly one production call to `sortExportBindingEvidence(projected)`, owned by `mergeExportBindingEvidence`;
- exactly one production call site to `exportBindingEvidenceOrderFor`, inside the decoration loop;
- `json.Unmarshal` remains only inside the canonical extractor for terminal/hop/failure notes, never inside the comparator;
- no second projected sort or alternate key authority exists.

## Exact Source/Test Diff And Hashes

| Path | Diff | Final SHA-256 |
|---|---:|---|
| `internal/resolution/export_binding_proof.go` | `+18/-4` | `36D41BC7336A5A04B4C77D1EE9DEB45DE4E5E2BA120EEFBF98419BD2D846C4D2` |
| `internal/resolution/export_binding_proof_test.go` | `+477/-0` | `99C6F6FC5FD0AE4D1BFDFE547D5F67C958544AA61FFD89895DC0BAC79C839BBC` |

Total source/test diff: `+495/-4`.

Test additions begin at:

- `export_binding_proof_test.go:245` — `TestA004ProjectionAndCoalescingMatchPreOptimizationOracle`;
- `export_binding_proof_test.go:374` — `TestA004MergeOrderingMatchesPreOptimizationOracleAdversarially`;
- `export_binding_proof_test.go:488` — `TestA004EmptyAndNoProofPathsMatchPreOptimizationOracle`.

The oracle is test-local, preserves the pre-A004 pre-sort/merge/comparator behavior, and is never a production fallback. Every differential assertion compares the complete ordered `[]graph.Evidence` value, not a semantic set.

## Build-Holder Clearance And Canonical Full Build

Preflight ran before the build:

- `anvien doctor locks --repo E:\Anvien --json` — analyze lock `free`; lock file absent; no stale/recoverable/foreign holder.
- `anvien doctor processes --json` — listed editor-owned npm MCP processes and a VS Code-owned Playwright test server; no repo-local analyze process, Go/Node build process, or proven holder of canonical `E:\Anvien\anvien\bin` outputs existed.
- Terminated processes: none, because no build-related holder was proven.

Canonical command:

```text
pwsh -NoLogo -NoProfile -NonInteractive -ExecutionPolicy Bypass -File E:\Anvien\scripts\full-build.ps1
```

Result: PASS, exit `0`, before all test execution.

- Vendor verification: `1798` files, `0` missing, `0` extra, `0` mismatched.
- Go runtime and launcher/Vite production build: PASS.
- Runtime version: `1.2.8`.
- Runtime SHA-256: `7C077F8555ABDD9D1A76ED19931B448B208622A3B13F170BB71131A463E826F6`.
- Embedded maintenance analyze: exit `0`; `2243/766/0`; graph `124258/171186`.
- Vite emitted only existing dynamic-import/chunk-size warnings; they did not fail the build.

The maintenance analyze is build validation, not target measurement or an additional Coder benchmark.

## Post-Build Validation

### A004 differential and export-binding proof suite

Command:

```text
go test ./internal/resolution -run '^(TestP5DProjectionEncodesRetainedDirectAndBarrelProofs|TestP5DProjectionEncodesOwnedMemberProof|TestP5DProjectionAndCoalescingRetainFailureAndSourceSiteProofs|TestP5DResolveEmitsCallAccessProofAndConservesCoalescedSites|TestA004ProjectionAndCoalescingMatchPreOptimizationOracle|TestA004MergeOrderingMatchesPreOptimizationOracleAdversarially|TestA004EmptyAndNoProofPathsMatchPreOptimizationOracle)$' -count=1 -v
```

Result: PASS, exit `0`; all A004 and existing export-binding proof tests passed.

### Focused call/access/outcome regressions

Command:

```text
go test ./internal/resolution -run '^(TestP5DResolveEmitsCallAccessProofAndConservesCoalescedSites|TestResolveAttachesSourceBackedUnresolvedDiagnostics|TestP6BResolverCarriesCatalogValidationFailuresPerSourceSite|TestP6BExternalMemberLookupRequiresExternalReceiverForCallAndAccess|TestP6C3FinalOutcomeStatusPrecedenceAndCarriageMatrix|TestLegacyEmitReferencesParityCoversMemberCallsAccessesAndUnresolved)$' -count=1 -v
```

Result: PASS, exit `0`; all six named consumer regressions passed.

### A003 diagnostic and Graph JSON parity

Command:

```text
go test ./internal/graphhealth -run '^(TestDiagnosticAppenderMatchesLegacySemantics|TestDecodeStructuredResolutionOutcomeCanonicalSingleInterpretationParity|TestP6DStructuredResolutionOutcomesDriveDiagnosticPolicy|TestP6DStructuredResolutionOutcomesFailClosedOnMalformedOrDriftedCarriers|TestP6DStructuredOutcomeProjectionPreservesGraphJSONAndHealthParity)$' -count=1 -v
```

Result: PASS, exit `0`; appender, canonical decoder, fail-closed policy, and Graph JSON/health parity passed.

### Analyzer outcome/graph carriage

Command:

```text
go test ./internal/analyze -run '^(TestP6C3AnalyzeResultPreservesFinalOutcomesAndGraphCarriage|TestP6C3AnalyzeCapabilityOutcomesRetainAcceptedAuthorityStatus)$' -count=1 -v
```

Result: PASS, exit `0`; both named analyzer carriage tests passed.

### Full `internal/resolution` truth

Command:

```text
go test ./internal/resolution -count=1
```

Result: exit `1`, solely `TestProofBasedCallAccessGoldenCorpus` at `proof_accuracy_golden_test.go:166`.

The observed callback diagnostic remains exactly the recorded preserve-only mismatch: actual `Classification:"unclassified"` and `Actionability:"review"`, with the same `SourceSiteID`, range, Note payload, and other fields. No A004 or other new/changed failure appeared. The package is not labeled PASS, and the golden owner remains untouched.

E2E verification:

```text
[PASS] Compiled: canonical full build -> exit 0
[PASS] Runtime boundary: full-build maintenance analyze -> exit 0, 2243/766/0
[PASS] Happy path: projected/coalesced terminal-hop-failure evidence -> full ordered oracle equality
[PASS] Edge cases: malformed notes, duplicates, equal keys, permutations, multiple sites, no-proof, non-mutation, idempotency -> full ordered oracle equality
```

## Preserve And Forbidden-Surface Proof

- `git diff --name-only` before report creation contained only the two authorized source/test paths.
- No edit was made to `resolve.go`, `emit.go`, `outcome.go`, graphhealth, graph types, public `graph.Evidence`, persistence/readers, Graph JSON/CLI output, instrumentation, scripts, target repositories, D002-D017, plan rules, ledgers, Architect/Planner/Supervisor reports, or any other production/test file.
- Function signatures and public/persisted shapes are unchanged.
- No target Cheapapp/Restaurant measurement, A00x candidate build, post-measurement Supervisor, disposition, detect-changes, staging, or commit was performed.
- Scoped `git diff --check` passed for the two source/test paths and this report.
- Staged set is empty.

## Handoff

Candidate is ready for Main's independent source/build/test verification and later target-separated A004 measurement ownership. Coder stops here and does not open or direct measurement, Supervisor, disposition, detect, staging, commit, P3, or Child 07 lanes.

`CODER_A004_READY_FOR_MEASUREMENT_HANDOFF`
