# System Architect Report — Child 06A A002 Run-Scoped Diagnostic Appender

Date: 2026-08-25  
Cursor: `A002 / D001 resolve_calls`  
Decision: `ARCHITECT_A002_READY_FOR_PLANNER`

## 1. One selected direction

Apply the successful A001 method—not its import-index mechanism—to A002: remove the largest measured repeated work with one private, reversible, run-scoped structure. Add a write-through diagnostic appender for one `ResolveBoundInto` execution so each node's existing diagnostics are normalized once on first touch and each incoming diagnostic is normalized once, while the graph continues to receive the same `[]graphhealth.Diagnostic` value after every append.

No second optimization is authorized. In particular, A002 does not redesign diagnostic policy, eliminate the two JSON unmarshals inside one structured-outcome decode, change bucket lookup/sorting, or open D002-D017.

## 2. Current authority and measured residual

Current source/index authority is commit `56090ac3c396b11b7a7aa0e9240ec5acbe562f01`. Fresh `anvien analyze --force` completed with `2216` files scanned, `765` parsed, and `0` failures.

The accepted repositories remain independent:

| Target | D001 `resolve_calls` | Parent `resolution` | Process wall | Calls / files | Diagnostics / outcomes |
|---|---:|---:|---:|---:|---:|
| Cheapapp | `25.045225300 s` | `184.481061700 s` | `279.105934600 s` | `27,890 / 887` | `57,683 / 86,742` |
| Restaurant Manager | `40.769294200 s` | `136.436879300 s` | `218.680628900 s` | `86,030 / 1,234` | `129,009 / 186,251` |

The CPU profiles were filtered to stacks containing `resolution.resolveCall`. These are cumulative CPU samples and causal attribution, not elapsed-time claims and not additive:

| Function | Cheapapp cumulative CPU | Restaurant cumulative CPU |
|---|---:|---:|
| `resolution.resolveCall` | `23.15 s` | `37.47 s` |
| `graphhealth.AppendDiagnosticToNode` | `20.30 s` | `30.18 s` |
| `graphhealth.normalizeDiagnosticMetadata` | `20.09 s` | `29.56 s` |
| `graphhealth.decodeStructuredResolutionOutcome` | `20.06 s` | `29.40 s` |
| `graphhealth.appendDiagnosticsFromProperties` | `18.97 s` | `28.11 s` |
| `graphhealth.normalizeDiagnosticSlice` | `17.60 s` | `26.07 s` |
| `resolution.(*emitter).emitUnresolvedReference` | `9.13 s` | `23.45 s` |
| `resolution.(*emitter).emitTypeScriptOutcomeDiagnostic` | `11.41 s` | `7.37 s` |

Raw evidence:

- Cheapapp benchmark SHA-256 `BD598F4D5BCC52E7FB248988889431E7DF0CEC6788D583A4BF3637DE00CDC3BA`; profile SHA-256 `81606B16916225E5FCDAABEA7BE542F680626BF6093D1B026342608026289484`.
- Restaurant benchmark SHA-256 `72B9B441A84ABBA427F844C728B85DF1842C7A25187B785D5F7344CAFE9C1666`; profile SHA-256 `3A4C1BA5C535124DE3A5E898FA316C612AA7697523CF6B8819275DF468A1B961`.
- Roots: `E:\Anvien\.tmp\child06a_a001_cheapapp_benchmark\owner-final-run\optimized` and `E:\Anvien\.tmp\child06a_a001_restaurant_manager_benchmark\fresh-supported-exclude-userapi-20260825\runs\optimized`.

## 3. Exact attribution: measured fact versus inference

Measured/profile fact: the diagnostic append/normalization stack accounts for most sampled CPU beneath `resolveCall` in both repositories.

Source fact in `internal/graphhealth/diagnostics.go`:

- `AppendDiagnosticToNode` normalizes the incoming diagnostic at line 39.
- `appendDiagnosticsFromProperties` obtains all existing diagnostics at lines 92-94 and normalizes the incoming diagnostic again.
- A stored `[]Diagnostic` enters `normalizeDiagnosticSlice` at lines 151-156; its loop at lines 216-220 allocates a new slice and re-normalizes every prior entry.
- Each structured normalization enters `decodeStructuredResolutionOutcome`; lines 334-346 unmarshal the same JSON `Note` first into an envelope and then into the outcome.
- Production resolution has exactly two `AppendDiagnosticToNode` call sites: `internal/resolution/emit.go:218` and `internal/resolution/outcome.go:392`.

Inference, explicitly not an operation-counter measurement: because each successful append writes `[]Diagnostic` back to the node, the next append re-normalizes all prior entries. For `k` appends to one node, that repeated-normalization component has the triangular `0 + 1 + ... + (k-1)` shape. The raw profiles prove the cost is material; they do not measure the exact number of repeated decodes per node.

## 4. Anvien owner and blast-radius evidence

All three candidate files are fresh and HIGH file risk:

| File | Symbols | Inbound / outbound | Linked flows / tests |
|---|---:|---:|---:|
| `internal/graphhealth/diagnostics.go` | `154` | `50 / 70` | `1 / 19` |
| `internal/resolution/emit.go` | `154` | `81 / 231` | `13 / 24` |
| `internal/resolution/outcome.go` | `119` | `87 / 124` | `11 / 23` |

Exact symbol impact warnings:

| Symbol | Risk | Impacted / direct | Modules / processes |
|---|---:|---:|---:|
| `AppendDiagnosticToNode` | CRITICAL | `7 / 2` | `1 / 24` |
| `emitter` | CRITICAL | `28 / 25` | `2 / 49` |
| `newEmitter` | CRITICAL | `6 / 1` | `4 / 32` |
| `emitUnresolvedReference` | CRITICAL | `7 / 4` | `2 / 36` |
| `emitTypeScriptOutcomeDiagnostic` | LOW | `0 / 0` | `0 / 0` |

HIGH/CRITICAL are scope warnings, not edit prohibitions. The LOW method result does not cancel the HIGH `outcome.go` file risk or its measured profile evidence.

## 5. Exact implementation boundary

Planner may authorize only these production owners:

1. `internal/graphhealth/diagnostics.go`
   - Add a run-scoped appender (recommended names: `DiagnosticAppender`, `NewDiagnosticAppender`, `AppendToNode`).
   - Cache only `map[nodeID][]Diagnostic` slice headers/state.
   - On first successful touch, read the current node and normalize its existing supported property representation once.
   - Apply the same source-node/count defaults and normalize each incoming diagnostic once.
   - Merge through one private helper that assumes normalized inputs but preserves current bucket equality, count merge, empty-target fill, earliest-line rule, and stable sort.
   - On every successful append, read the current node, preserve all other properties, set the unchanged `DiagnosticPropertyKey` representation to `[]Diagnostic`, and call `Graph.AddNode` as today.
   - Keep `AppendDiagnosticToNode` available for all other callers with exactly its current normalization and observable semantics; a private-helper refactor is allowed only to prevent rule duplication.
2. `internal/resolution/emit.go`
   - Add the appender field to `emitter`.
   - Initialize it in `newEmitter` from the same graph.
   - Route `(*emitter).emitUnresolvedReference` through it.
3. `internal/resolution/outcome.go`
   - Route `(*emitter).emitTypeScriptOutcomeDiagnostic` through the same appender.

The appender is owned by the existing per-run `emitter`; `internal/resolution/resolve.go` needs no edit and execution order remains unchanged.

## 6. Preserved contracts

The implementation must preserve exactly:

- false-return behavior for nil graph, blank node ID, blank diagnostic kind, or absent source node;
- `SourceNodeID` and `Count` defaults;
- structured fail-closed and legacy classification/actionability policy;
- diagnostic bucket identity, merged counts, target fill, earliest source line, stable order, and write-through visibility after each call;
- `[]Diagnostic` graph property, Graph JSON, persistence/readback, health projections, public output, and deterministic replay;
- every resolution outcome, encoded `Note`, authority/proof/source-site field, metric, node, relationship, ID, label, property, and resolution branch/order;
- pre-existing diagnostic representations on an input graph, normalized once on first touch exactly as the current function would normalize them.

The cache is valid only under the existing sequential `ResolveBoundInto` ownership. It must not cache whole nodes or tolerate an unproven concurrent/out-of-band diagnostic writer by silently overwriting it.

## 7. Memory and lifecycle

Let `T` be the number of source nodes touched by resolution diagnostics. Additional retained state is `O(T)` map entries and slice headers. The cache and graph must point to the same normalized slice; it must not retain a second copy of diagnostic objects. First-touch normalization may allocate the replacement slice once, after which the previous property value can be collected.

The map is created by `newEmitter`, performs no serialization, I/O, goroutine, lock, or global caching, and becomes unreachable when `ResolveBoundInto` returns. The diagnostic slices remain owned only by the returned graph. No explicit flush/finalizer is permitted because every append is write-through.

## 8. Required implementation and validation order

1. Coder repeats fresh `anvien --help`, `anvien analyze --force`, file-detail, and exact symbol impacts; this report does not replace the pre-edit gate.
2. Edit the three production owners first.
3. Only after production behavior is complete, add focused test bytes in one new file: `internal/graphhealth/diagnostics_test.go`. The test must compare repeated appender writes with the legacy append semantics for first-touch decoded input, multiple unique entries, duplicate-bucket merging, stable order, policy fields, write-through after every append, and absent-node failure.
4. Perform build-holder/lock preflight, then run the canonical full build before any test validation:

   `pwsh -NoLogo -NoProfile -NonInteractive -ExecutionPolicy Bypass -File E:\Anvien\scripts\full-build.ps1`

5. Post-build focused tests must include:
   - the new appender-equivalence test;
   - `TestP6DStructuredOutcomeProjectionPreservesGraphJSONAndHealthParity`;
   - `TestResolveAttachesSourceBackedUnresolvedDiagnostics`;
   - `TestP6C3FinalOutcomeStatusPrecedenceAndCarriageMatrix`;
   - `TestP6BResolverCarriesCatalogValidationFailuresPerSourceSite`.
6. Run `go test ./internal/graphhealth -count=1` and `go test ./internal/resolution -count=1`. The known preserve-only `TestProofBasedCallAccessGoldenCorpus` baseline failure may be reported only with exact unchanged-baseline proof; it is not a pass, and any new failure blocks A002.
7. Run Anvien change detection before any later implementation commit.

## 9. Two-target measurement and retention contract

Use the accepted A001 artifacts above as each repository's own `before`; do not average or combine them. Build one identifiable A002 candidate from the accepted A001 state plus only the authorized A002 diff, with identical 17-child instrumentation and runtime/native/build contract.

- Cheapapp candidate command must retain target `E:\cheapapp.org` and flags `--force --skip-git --json --progress`; only executable identity, benchmark path, and label may differ.
- Restaurant candidate command must retain target `E:\Restaurant_manager`, flags `--force --json --progress`, and exactly one `--exclude electron/renderer/src/api/userApi.ts`; only executable identity, benchmark path, and label may differ.
- Keep each graph under its target-local `.anvien`; preserve target HEAD/status/workload and record process/executable/artifact hashes.
- For each repository separately, record D001 `resolve_calls`, parent `resolution`, analyzer total, and process wall. Full-build, test, microbenchmark, or CPU-profile duration cannot substitute for these wall-time boundaries.
- Calls/files, all 30 operation rows, all 17 child rows, files scanned/parsed/failed, graph nodes/relationships, DB readback counts, resolution semantic metrics, diagnostics/outcomes, canonical Graph JSON, and public stdout must remain equivalent. Raw Ladybug/meta byte identity alone is not the semantic gate.
- Record `startAllocBytes`, `endAllocBytes`, and `maxObservedSys` plus proof that retained appender state follows `O(T)` and does not duplicate diagnostics.

No speed percentage or threshold is predicted. KEEP requires a valid same-work D001 improvement and corresponding net parent/process benefit on both repositories independently, with all correctness/resource gates passing. Otherwise roll back the A002-owned hunks and leave D001 open under the plan rules.

## 10. Rollback and mandatory STOP

Rollback only the A002 appender/helper, the `emitter` field/initialization, the two call-site substitutions, and the focused new test file if any build, correctness, equivalence, determinism, lifecycle, memory, D001, parent, or process gate fails. Do not use broad reset/checkout/stash operations or disturb protected worktree changes.

STOP and return to Main for a fresh architecture decision if implementation requires any additional production owner, including `internal/graph/types.go`, `internal/graphhealth/policy.go`/`Diagnostic`, `internal/resolution/resolve.go`, resolution outcome/metric schemas, persistence/readers, public contracts, global/concurrent caching, or timing instrumentation. A newly exposed post-A002 residual belongs to a later attempt; it must not be folded into A002.

## 11. Handoff

`ARCHITECT_A002_READY_FOR_PLANNER`

Selected direction: introduce one per-`ResolveBoundInto`, write-through diagnostic appender shared by the two resolution diagnostic emitters so previously stored diagnostics are normalized once per touched node instead of on every append.

This step wrote only this architecture report; it did not edit source, tests, ledgers, or plan rules and did not build, test, benchmark, stage, commit, or change disposition.
