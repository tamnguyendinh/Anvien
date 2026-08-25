# Child 06A A001 Coder Data Handoff — Import Claim Index

- Recorded: `2026-08-25 17:57:10 +07:00`
- Scope: `B2-P2A-A001-D001 resolve_calls`
- Purpose: concise implementation/build/test/measurement handoff; this is not an audit.

## Goal and implemented design

A001 removes repeated whole-workspace `w.imports` scans and repeated source-path normalization from D001 call resolution without changing resolution semantics, ordering, graph output, persistence, or lifecycle behavior.

The implementation adds a private, run-scoped in-memory index keyed by `{canonical source file path, exact case-sensitive local import name}`. Each value stores original `w.imports` indices in existing order. `resolveImports` populates the index after source-path canonicalization and includes resolved and unresolved, semantic and nonsemantic, duplicate imports. `explicitImportNameClaimed` uses bucket presence, while `explicitImportCallState` traverses only the indexed original candidates. It never derives semantic order from map-key iteration and does not reuse the resolved-only `importsByReceiver` index.

## Exact source/test boundary and hunks

The five-file candidate is `+174/-10`:

| File | Diff | Hunk summary |
|---|---:|---|
| `internal/resolution/indexes.go` | `+8/-1` | Adds and initializes the private claim index; appends each qualifying original import index during `resolveImports` while leaving `importsByReceiver` semantics intact. |
| `internal/resolution/resolve.go` | `+2/-6` | Replaces the whole-import scan in `explicitImportNameClaimed` with exact indexed-bucket presence. |
| `internal/resolution/export_resolution.go` | `+2/-3` | Replaces only `explicitImportCallState` candidate acquisition with ordered indexed traversal; semantic filters and target checks remain unchanged. |
| `internal/resolution/export_resolution_test.go` | `+64/-0` | Adds `TestExplicitImportCallStatePreservesIndexedCandidateSemantics` for semantic filtering, `allowed`/`importTargets`, target-nil, duplicate/order traversal, and unresolved/nonsemantic candidates. |
| `internal/resolution/resolution_test.go` | `+98/-0` | Adds `TestBuildWorkspaceIndexesAllImportClaimsByCanonicalSourceAndExactName` for canonical source paths, exact case-sensitive local names, original order, duplicates, unresolved/nonsemantic claims, and separation from resolved-only `importsByReceiver`. |

## Pre-edit impact warning

Fresh pre-edit analysis exited `0` with `2211` files scanned, `765` parsed, and `0` failed. All three production files were `HIGH`. Exact symbol impact was `CRITICAL` for:

- `workspace`: `145` symbols / `30` files / `71` processes / `7` modules;
- `buildWorkspace`: `62 / 26 / 23 / 7`;
- `resolveImports`: `37 / 19 / 17 / 3`.

`explicitImportNameClaimed` and `explicitImportCallState` reported `LOW/0`, but that was explicitly treated as non-assurance because unresolved method-call edges remain. The `HIGH/CRITICAL` results were blast-radius warnings, not edit bans, and the candidate stayed within the Architect-authorized owner boundary.

## Canonical build and post-build tests

Canonical full build ran before all validation tests:

- Command: `pwsh -NoLogo -NoProfile -NonInteractive -ExecutionPolicy Bypass -File E:\Anvien\scripts\full-build.ps1`
- Start: `2026-08-25T13:40:12.1770424+07:00`
- End: `2026-08-25T13:46:27.4898463+07:00`
- Exit: `0` (`PASS`)
- Build-holder preflight: analyze lock free; no repo-local runtime or repo-scoped Go/Node build holder; no process termination required.

Post-build focused validation:

- Command: `go test ./internal/resolution -run '^(TestBuildWorkspaceIndexesAllImportClaimsByCanonicalSourceAndExactName|TestExplicitImportCallStatePreservesIndexedCandidateSemantics)$' -count=1 -v`
- Result: exit `0`; both focused tests passed. Main independently reran the same command and observed both passing.

Post-build package validation:

- Command: `go test ./internal/resolution -count=1`
- Result: exit `1`, solely at `TestProofBasedCallAccessGoldenCorpus` in `proof_accuracy_golden_test.go:166`.
- Exact mismatch: expected callback metadata `in_repo_unresolved/analyzer_gap`; actual `unclassified/review`.
- Preserve-only proof: an isolated overlay replaced all five A001 files with their exact HEAD blobs (`5/5` blob IDs matched), and the targeted golden still failed with the identical diagnostic. `proof_accuracy_golden_test.go` was unchanged and outside the editable boundary, proving this is a pre-existing failure rather than an A001 regression or A001 validation pass.

## Independent target benchmarks

Each target retained its own baseline; results were never averaged or combined.

| Target | D001 `resolve_calls` | Parent `resolution` | Process E2E | D001 denominator |
|---|---:|---:|---:|---:|
| `E:\cheapapp.org` | `209.823705100 -> 25.045225300 s` | `733.384225100 -> 184.481061700 s` | `890.314783200 -> 279.105934600 s` | `calls=27890; files=887` |
| `E:\Restaurant_manager` | `501.638742800 -> 40.769294200 s` | `1090.085959900 -> 136.436879300 s` | `1178.391336900 -> 218.680628900 s` | `calls=86030; files=1234` |

Both pairs had `30/30` top-level operations, `17/17` resolution children, zero denominator mismatches, process exits `0`, and canonical Graph JSON/output/workload equivalence. Supervisor returned `SUPERVISOR_A001_PASS`; Main recorded `E2-P2A-A001DECISION1` with disposition `KEEP` and promoted the optimized values separately per repository.

## Change detection and commit

Required change detection:

- Command: `anvien detect-changes --repo E:\Anvien --scope all`
- Exit: `0`
- Overall risk: `CRITICAL`; file-layer risk: `HIGH`
- Changed: `545` symbols across `24` files
- Affected: `25` symbols across `22` files
- Affected layers: backend `18`, mixed `7`
- Affected areas: analyzer `8`, CLI `5`, mixed `8`, resolution `4`
- Resolution-gap changes: `272`; degraded nodes: `0`; total resolution gaps: `0`

Accepted commit:

- Commit: `17a1f3af37dcb61f9d389345822b6470a8f772cc`
- Message: `perf(resolution): index import claims for call resolution`
- Manifest: exactly these 12 paths:

  1. `internal/resolution/indexes.go`
  2. `internal/resolution/resolve.go`
  3. `internal/resolution/export_resolution.go`
  4. `internal/resolution/export_resolution_test.go`
  5. `internal/resolution/resolution_test.go`
  6. `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy/2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy-plan.md`
  7. `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy/2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy-evidence.md`
  8. `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy/2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy-benchmark.md`
  9. `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy/2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy-actual-status.md`
  10. `reports/Investigation/rp_child06a_a001_cheapapp_benchmark.md`
  11. `reports/Investigation/rp_child06a_a001_restaurant_manager_benchmark.md`
  12. `reports/Supervisor/rp_supervisor_260825_172308_by_gpt-5_a001_measurement_equivalence_acceptance.md`

## Preserved invariants, residuals, and next owner

Preserved invariants include exact resolution outcomes, confidence/proof/unresolved notes and call-branch order; trimmed case-sensitive import names; original import order, duplicates, first-match/traversal; unchanged `cleanPath`, target resolution, lexical shadowing, semantic/nonsemantic and unresolved-claim behavior; unchanged graph nodes, relationships, IDs, labels, properties, counts, ordering, Graph JSON, persistence/readback, determinism, invalidation, failure propagation, rollback, temporary-file publication, and public contracts. The index is private and run-scoped, performs no serialization/I/O/goroutine/global caching, and uses `O(imports + unique keys)` memory while storing indices rather than duplicate import objects.

Known residuals are the proven pre-existing preserve-only golden failure above and a Cheapapp Ladybug/meta byte difference accepted as non-blocking because canonical graph, DB readback counts, public output, and affected semantics match and A001 did not modify persistence or metadata code. P2-A, its parent, and D001 remain open; D002-D017 remain queued. No A002 source/test work is authorized by this handoff.

Next owner: Main Orchestration, which must open a fresh visible A002 Architect on the same active D001 before any further Planner or Coder transition.
