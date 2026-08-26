# Child 06A A005 Residual Attribution

Verdict: `A005_ATTRIBUTION_COMPLETE`

## Lane And Scope

- Functional outcome: prove one current residual cause shared by Cheapapp and Restaurant Manager, its exact owner, and the complete call path needed for the next A005 Architect attempt.
- Slice: `P2-A / A005 / B1-P1A-OP001 resolution / B2-P2A-A001-D001 resolve_calls`.
- Authoritative clean docs/source HEAD used for graph evidence: `6ee70e29ebd8b0345a28a95eada5b77b5c0f737f`.
- Role boundary: read-only attribution. This report is the only write. No architecture direction, plan translation, code/test edit, build, target analyze, target measurement, Supervisor, disposition, detect, staging, or commit was performed.
- Current state entering this packet: `A004_ROLLBACK_COMPLETE / A005_CURRENT_BASIS_RECORDED / RESIDUAL_ATTRIBUTION_PENDING / ARCHITECT_LOCKED / D001_STREAK_1`.

## Accepted Current Basis

Elapsed wall time remains the controlling authority. Targets remain separate and are never averaged or combined.

| Target | Accepted D001 `resolve_calls` | Accepted parent `resolution` | Accepted analyzer | Accepted process wall | Denominator |
|---|---:|---:|---:|---:|---|
| Cheapapp | `3.447846300 s` | `20.472602300 s` | `93.531974900 s` | `95.630648200 s` | `calls=27890; files=887` |
| Restaurant Manager | `9.401585300 s` | `20.850792800 s` | `98.020546700 s` | `101.096911900 s` | `calls=86030; files=1234` |

These are the accepted A003 values. A004 values below remain rejected-candidate evidence only.

## Exact Profile Identities

| Target / attempt | Role | Exact profile | Bytes | SHA-256 |
|---|---|---|---:|---|
| Cheapapp A003 | accepted-current causal profile | `E:\Anvien\.tmp\child06a_a003_cheapapp_frozen_17child_20260826\candidate\resolution.cpu.pprof` | `65,363` | `806FAD79B12C567618778A8232E8521991882CE6C745AC2EBBE7B09ECC267107` |
| Restaurant A003 | accepted-current causal profile | `E:\Anvien\.tmp\child06a_a003_restaurant_manager_frozen_17child_20260826\candidate\resolution.cpu.pprof` | `56,397` | `4272FFAB0ABE4C14C470A139A31A601F31323B1373E0FA2986A09935456FA3E2` |
| Cheapapp A004 | rejected candidate, residual-exposure only | `E:\Anvien\.tmp\child06a_a004_cheapapp_frozen_17child_20260827\candidate\resolution.cpu.pprof` | `52,040` | `C1250E0F4F8655054967125F454FED6D9B7329EA8528726C04B7D65BA61F7616` |
| Restaurant A004 | rejected candidate, residual-exposure only | `E:\Anvien\.tmp\child06a_a004_restaurant_manager_frozen_17child_20260827\candidate\resolution.cpu.pprof` | `55,580` | `F41A0CB057B6F7F688343A565FC50AD8B9D947B51700C2C36BC824498B70A9E4` |

## Separate A003 Versus A004 CPU Evidence

Commands were bounded `go tool pprof -top`, `-cum`, `-focus`, `-ignore`, and `-list` reads over the four existing profiles. No profile was captured. CPU samples overlap, are non-additive, are not elapsed wall time, cannot prove speedup, and do not select priority over the accepted wall-time basis.

The known-family exclusion removed stacks containing the accepted A002 diagnostic appender/normalization family, the accepted A003 canonical decoder family, and the rejected A004 export-binding append/merge/sort/order-key family. It was used only to expose residual D001 work.

| Target / attempt | Profile duration / total CPU samples | `resolveCall`-focused cumulative | `resolveCall` after known-family exclusion | `(*resolutionOutcomeCollector).record` | `marshalResolutionOutcome` | `recordRepositoryResolvedOutcome` |
|---|---:|---:|---:|---:|---:|---:|
| Cheapapp A003 | `20.46 / 25.80 s` | `3.12 s` | `0.71 s` | `0.26 s` | `0.16 s` | `0.11 s` |
| Cheapapp A004 | `13.26 / 16.81 s` | `1.60 s` | `0.67 s` | `0.27 s` | `0.22 s` | `0.10 s` |
| Restaurant A003 | `20.84 / 23.96 s` | `8.29 s` | `6.18 s` | `0.63 s` | `0.49 s` | `0.32 s` |
| Restaurant A004 | `19.41 / 24.34 s` | `8.07 s` | `5.93 s` | `0.74 s` | `0.57 s` | `0.37 s` |

Source-mapped `pprof -list` attributes the selected helper samples to the unique-outcome path at `internal/resolution/outcome.go:100` and its `json.Marshal` at line `221`:

| Target / attempt | `record` line 100 cumulative | `marshalResolutionOutcome` cumulative | Helper line 221 `json.Marshal` | Helper line 225 byte-to-string |
|---|---:|---:|---:|---:|
| Cheapapp A003 | `0.16 s` | `0.16 s` | `0.13 s` | `0.03 s` |
| Cheapapp A004 | `0.22 s` | `0.22 s` | `0.20 s` | `0.02 s` |
| Restaurant A003 | `0.49 s` | `0.49 s` | `0.41 s` | `0.08 s` |
| Restaurant A004 | `0.57 s` | `0.57 s` | `0.50 s` | `0.07 s` |

The A004 candidate profiles therefore expose the same outcome-recording/encoding family after the A004 ordering work was removed. The table does not compare CPU values as a speed result and does not promote A004.

## Selected Residual Mechanism

Selected cause: **eager record-time JSON serialization of each accepted `ResolutionOutcome`, including resolved-outcome paths whose returned encoded bytes have no caller, followed by another serialization of every finalized outcome during graph/reference projection**.

Exact primary owner:

- file: `E:\Anvien\internal\resolution\outcome.go`;
- method that owns the decision to serialize at record time: `(*resolutionOutcomeCollector).record`, lines `83-106`, specifically unique-outcome line `100` and duplicate-outcome line `93`;
- serialization helper: `marshalResolutionOutcome`, lines `220-226`, specifically `json.Marshal` at line `221`;
- current file SHA-256: `02092F9FE7DA2A4BDB49E13056FBA3C97DC24F416141E7F27866EE80F60C1F7E`.

Why this mechanism is selected:

1. It is sampled under `resolveCall` in all four profiles, including both A004 profiles after the rejected export-ordering optimization removed that local family.
2. Current source proves that a unique outcome is cloned, validated, stored by `SourceSiteID`, and then serialized immediately at line `100`.
3. `recordRepositoryResolvedOutcome` calls `e.outcomes.record(...)` as a statement and consumes none of its three return values. Thus, for repository-resolved calls, the encoded string produced at record time has no downstream caller.
4. The resolved TypeScript branch similarly obtains `encoded` in `recordTypeScriptLookup` but uses it only when the outcome is not `ResolutionResolvedExternal`; the resolved-external path does not consume it.
5. After call/access/type processing, `projectResolutionOutcomes` serializes every finalized outcome again at line `406` to create the canonical encoded map used by graph/reference/diagnostic carriers.
6. This is materially distinct from:
   - A001: global import-claim scanning/indexing;
   - A002: repeated diagnostic normalization and the run-scoped diagnostic appender;
   - A003: decoding/interpreting structured diagnostic Notes;
   - A004: export-binding evidence pre-sort/comparator order-key decoding and cached-key final sort.

This attribution does **not** decide that serialization must be delayed, cached, removed, or moved. Record-time encoding also participates in current error timing and immediate unresolved-diagnostic construction. Those semantics belong to Architect.

## Verified Current-Source Facts, Profile Inference, And Missing Evidence

### Verified current-source facts

- `resolutionOutcomeCollector.record` clones and validates every input before checking `bySourceSite`.
- Both duplicate and unique valid paths call `marshalResolutionOutcome`; the unique path stores a cloned outcome first and then encodes it.
- Four production methods call `record`: `recordRepositoryResolvedOutcome`, `recordRepositoryUnresolvedOutcome`, `recordIntrinsicTypeOutcome`, and `recordTypeScriptOutcome`. D001 reaches the repository-resolved, repository-unresolved, and TypeScript callers; the intrinsic caller belongs to type-annotation work.
- `recordRepositoryResolvedOutcome` discards every returned value.
- Repository-unresolved outcomes use the returned encoded string as `Diagnostic.Note` through `emitUnresolvedReference`.
- TypeScript outcomes use the encoded string for a diagnostic only when an added outcome is not `ResolutionResolvedExternal`.
- `finalize` validates/clones/sorts the retained outcome inventory by `SourceSiteID`.
- `projectResolutionOutcomes` re-encodes every finalized outcome and supplies those bytes to relationship evidence, both reference-index projections, and diagnostic/outcome consistency validation.
- The analyzer carries the structured slice as `analyze.Result.ResolutionOutcomes` and the graph carriers as `analyze.Result.Graph`.
- The graph, including resolution-outcome evidence/diagnostic carriers, flows into Ladybug DB load and Graph JSON snapshot. Normal CLI JSON/text publication exposes repository/path/count/file-projection fields, not the `ResolutionOutcomes` slice. `WriteBenchmark` serializes `result.Metrics` only.

### Profile-supported inference

- The record-time outcome encoding family consumes CPU under D001 on both targets.
- The family remains visible in both rejected A004 candidate profiles after removing A004's ordering work.
- The repository-resolved branch itself has samples on every target/attempt, making the source-proven discarded-return path credible for D001.
- The profile evidence does not prove how much elapsed wall time a future design can retain.

### Missing evidence not required to open Architect

- Exact record invocations split by repository-resolved, repository-unresolved, TypeScript status, unique, and duplicate paths.
- Exact number and byte size of record-time versus projection-time encodings.
- Exclusive allocation/GC cost owned only by this family.
- Exclusive D001 wall-time share or a predicted child/parent/process improvement.
- Whether changing encoding time/location preserves the exact current failure timing without additional design work.

No new target capture is needed to prove the attribution. Any counters or measurements needed to validate a future architecture must be explicitly selected by Architect/Planner and cannot be inferred here.

## Complete Upstream And Downstream Call Path

### Entry through D001

```text
internal/cli.newAnalyzeCommand.func1
-> internal/analyze.Run
-> runPhase(PhaseResolution)
-> resolution.ResolveBoundInto
-> for each w.files IR
-> for each ir.Calls
-> resolution.resolveCall
```

Exact source anchors are `internal/analyze/analyze.go:365-376` and `internal/resolution/resolve.go:57-128`, with the D001 loop at `resolve.go:91-94` and `resolveCall` at `385-605`.

### D001 branch paths into the owner

Repository-resolved call:

```text
resolveCall:567-602
-> (*emitter).emitReference:137-180
-> (*emitter).recordRepositoryResolvedOutcome:256-283
-> (*resolutionOutcomeCollector).record:83-106
-> marshalResolutionOutcome:220-226
-> encoding/json.Marshal
```

The encoded return is discarded by `recordRepositoryResolvedOutcome`.

Repository-unresolved call:

```text
resolveCall:388/392/548-560/563-565
-> (*emitter).emitUnresolvedReference:182-225
-> (*emitter).recordRepositoryUnresolvedOutcome:285-305
-> (*resolutionOutcomeCollector).record
-> marshalResolutionOutcome
-> encoded Diagnostic.Note
-> emitter-owned diagnosticAppender
-> graph node DiagnosticPropertyKey
```

TypeScript call:

```text
resolveCall:516-535
-> recordTypeScriptLookup:926-977
-> (*emitter).recordTypeScriptOutcome:331-370
-> (*resolutionOutcomeCollector).record
-> marshalResolutionOutcome
-> when added and non-resolved-external:
   emitTypeScriptOutcomeDiagnostic:372-397
   -> diagnosticAppender
   -> graph node DiagnosticPropertyKey
```

### Finalization and graph/output consumers

```text
ResolveBoundInto:110
-> (*resolutionOutcomeCollector).finalize:114-129
-> projectResolutionOutcomes:399-455
   -> marshalResolutionOutcome again at line 406 for every finalized outcome
   -> graph.Relationship.Evidence kind resolution-outcome-v1
   -> ReferenceIndex.BySourceScope evidence
   -> ReferenceIndex.ByTargetDef evidence
   -> resolutionOutcomeDiagnosticSites verifies Diagnostic.Note byte parity
-> resolution.Result.ResolutionOutcomes + resolution.Result.Graph
-> analyze.Run assigns analyze.Result.ResolutionOutcomes + Graph
-> later graph consumers:
   MRO / communities / processes / semantic enrichment
   Ladybug DB load
   Graph JSON snapshot
-> CLI consumers:
   registry/meta record
   generated AI context
   file projection
   public repo/path/count output
   benchmark Metrics serialization
```

The normal CLI output and benchmark do not directly serialize the structured outcome slice; the user-visible/persisted semantic effect is carried through Graph evidence/diagnostic bytes, while `analyze.Result.ResolutionOutcomes` remains an in-memory/API result field.

## Current Anvien Evidence And Blast Radius

One and only one fresh `anvien analyze --force` ran at clean authoritative HEAD before graph work. It passed with `2,249` scanned, `766` parsed code, `0` failed, and graph `124,165` nodes / `171,031` relationships. Pre/post HEAD was exact and the worktree remained clean.

### Containing file

- `file-detail internal/resolution/outcome.go`: current/non-stale, `changedSinceAnalyze=false`, risk `HIGH`.
- Complete file-detail counts: `119` symbols, `87` inbound, `122` outbound, `117` local relationships, `148` unresolved sites, `11` linked flows, `23` linked tests.
- Upstream file impact, depth `5`, tests included: `CRITICAL`; `177` impacted symbols; `119` contained symbols; `7` direct; `55` affected files; `1` affected flow; `11` linked flows; `23` linked tests.
- Exact affected process sample: `ResolveBoundInto -> ResolutionOutcomeCollector`, `4` total hits, earliest broken step `1`.
- Affected file samples with full reported per-file symbol counts include `internal/resolution/outcome.go (21)`, `internal/resolution/emit.go (12)`, `internal/resolution/resolve.go (9)`, `internal/analyze/analyze.go (2)`, `internal/cli/command.go (2)`, `internal/resolution/resolution_test.go (15)`, `internal/resolution/parser_integration_test.go (9)`, and persistence/provider test surfaces. Total affected-file count remains `55`.

### Primary method `(*resolutionOutcomeCollector).record`

- Exact-UID upstream impact, depth `5`, tests included: `CRITICAL`.
- Complete counts: `5` impacted symbols, `5` direct, `1` module, `10` affected processes, `2` affected files.
- Affected app layers: `backend=4`, `backend_test=1`; functional area: `resolution=5`.
- Complete affected files: `internal/resolution/outcome.go (4)` and `internal/resolution/p6c3_structured_outcome_test.go (1)`.
- Graph context resolves exactly five callers: four production callers (`recordRepositoryResolvedOutcome`, `recordRepositoryUnresolvedOutcome`, `recordIntrinsicTypeOutcome`, `recordTypeScriptOutcome`) and one test caller.
- Process samples include `RecordIntrinsicTypeOutcome -> MarshalResolutionOutcome`, `RecordIntrinsicTypeOutcome -> SetError`, `RecordRepositoryUnresolvedOutcome -> CleanPath`, and `RecordRepositoryUnresolvedOutcome -> GenerateID`.

### Helper `marshalResolutionOutcome`

- Exact-symbol upstream impact, depth `5`, tests included: `CRITICAL`.
- Complete counts: `127` impacted symbols, `3` direct, `25` modules, `52` affected processes, `48` affected files.
- Affected app layers: `backend=13`, `backend_test=113`, `cli_launcher=1`.
- Affected functional areas: `analyzer=25`, `cli=3`, `graph_health=1`, `providers=20`, `resolution=75`, `storage=3`.
- File samples include `internal/resolution/outcome.go (6)`, `internal/resolution/resolve.go (3)`, `internal/analyze/analyze.go (1)`, `internal/cli/command.go (2)`, `internal/resolution/resolution_test.go (15)`, `internal/resolution/parser_integration_test.go (9)`, and graph/persistence/provider tests.
- Exact graph-resolved incoming callers are `(*resolutionOutcomeCollector).record` at lines `93` and `100`, `projectResolutionOutcomes` at line `406`, and one same-package test caller.
- Direct flow sample: `RecordIntrinsicTypeOutcome -> MarshalResolutionOutcome`.

`HIGH` and `CRITICAL` are blast-radius warnings, not edit prohibitions. They also do not authorize any edit; a future Coder must repeat the fresh pre-edit gate for the exact Architect-approved symbols.

## Why A004 Is Excluded Despite Local Gain

| Target | Boundary | Accepted A003 | Rejected A004 | Direction |
|---|---:|---:|---:|---|
| Cheapapp | D001 | `3.447846300 s` | `2.074182500 s` | lower by `1.373663800 s` |
| Cheapapp | parent | `20.472602300 s` | `13.265999200 s` | lower by `7.206603100 s` |
| Cheapapp | analyzer | `93.531974900 s` | `107.287054400 s` | higher by `13.755079500 s` |
| Cheapapp | process | `95.630648200 s` | `144.975972400 s` | higher by `49.345324200 s` |
| Restaurant | D001 | `9.401585300 s` | `8.975767700 s` | lower by `0.425817600 s` |
| Restaurant | parent | `20.850792800 s` | `19.416099500 s` | lower by `1.434693300 s` |
| Restaurant | analyzer | `98.020546700 s` | `101.406172300 s` | higher by `3.385625600 s` |
| Restaurant | process | `101.096911900 s` | `135.569489100 s` | higher by `34.472577200 s` |

The A004 Supervisor returned `SUPERVISOR_A004_PASS` for correctness/equivalence/output/lifecycle only. Main correctly issued `NO_KEEP`: Child 06A requires lower child, lower parent, and retained lower end-to-end process time on each target, plus Supervisor PASS. A004 failed the end-to-end requirement on both targets and also had higher analyzer time on both.

The A003 Owner exception created no numeric tolerance and no A004+ precedent. A004 production/test bytes were mechanically rolled back; the restored hashes are `4BFE1976766FBF2E2257102070FF43CAC6D757E32E71DD583F8824781EAB6A8E` and `AD3DCA9E82EACFB31137560636B59D426A063AEB967613E724D5EE3017AD5812`, both with zero diff against HEAD. Therefore:

- A004 values cannot become accepted elapsed proof or ranking input.
- A004's exact-tuple/cached-key/single-sort export-binding direction cannot be selected, redesigned cosmetically, or reused for A005.
- A004 profiles are used only to prove that the selected outcome-recording residual remains visible after A004 removed its local ordering work.

## Smallest A005 Architect Decision Boundary

Architect must decide only the boundary around **record-time outcome encoding versus immediate diagnostic needs and final projection encoding**. This report chooses no implementation.

Exact questions Architect must resolve:

1. Which exact symbols may change: only `(*resolutionOutcomeCollector).record` and `marshalResolutionOutcome`, or also specific caller/projection symbols in `outcome.go`? Any additional file must be named and justified before Planner/Coder.
2. Which component owns the canonical encoded bytes for each status family, and at what lifecycle point, while keeping repository-unresolved and non-resolved TypeScript diagnostics immediately byte-identical?
3. How are record-time marshal failure timing, collector fail-closed error retention, duplicate `SourceSiteID` conflict detection, immutable cloning, and the `added` boolean preserved exactly if encoding timing changes?
4. How are exact `Diagnostic.Note`, relationship/reference `Evidence.Note`, `ResolutionOutcome` inventory, SourceSiteID mapping, order, and repeated-run determinism proven byte-for-byte identical?
5. What private transient/run-scoped resource is allowed, what is its complexity and lifetime, and how is retained duplication of encoded payloads prevented or bounded?
6. Which focused tests and downstream Graph JSON/Ladybug/native-reader/analyzer-result checks are mandatory, including resolved, unresolved, resolved-external, capability-unavailable, duplicate, conflict, malformed/failure, and replay cases?
7. What exact child/parent/analyzer/process measurements on both independent targets and what rollback condition govern the future attempt?

Architect must explicitly decide whether changing `projectResolutionOutcomes` is inside or outside A005. Attribution alone authorizes neither option.

## Preserve-Only Surfaces And Mandatory STOP

Preserve exactly:

- A001 import-claim index and its two indexed lookup paths; do not reintroduce global scans or cosmetically repeat A001.
- A002 run-scoped diagnostic appender, its two resolution routes, write-through semantics, `[]Diagnostic` representation, sequential `O(T)` lifetime, and every accepted A002 test byte. `outcome.go:392` is part of that accepted route and is preserve-only unless a fresh Architect explicitly owns an equivalent change.
- A003 canonical one-interpretation decoder and fail-closed tuple/policy semantics; do not add a second decoder/path or cosmetically repeat A003.
- Restored pre-A004 export-binding code/test bytes and exact observable ordering; do not revive the A004 exact-tuple/cached-key/single-sort direction.
- `ResolutionOutcome` schema/JSON tags/status/stage/target/proof/authority shapes, `graph.Evidence`, diagnostic contract, SourceSiteID identity, reference-index shapes, Graph JSON, Ladybug/native persistence/readback, analyzer result carriage, CLI output, benchmark metrics, determinism, freshness, failure propagation, transaction/temp/publication behavior, and the known preserve-only golden failure.
- P1 timing instrumentation, WAL cleanup fix, target workloads/options/denominators, D002-D017, remaining parent queue, P3, and Child 07.

Mandatory STOP and return to Main for a fresh decision if any future proposal:

- changes a file/symbol not explicitly approved by the A005 Architect/Planner;
- alters accepted A001/A002/A003 behavior or reopens the A004 ordering path;
- changes public/persisted shapes, diagnostic policy, decoder semantics, Graph/Ladybug/readers, normal CLI output, timing instrumentation, or target workload;
- weakens fail-closed validation/conflict/error behavior or changes encoded bytes/order;
- needs a global/cross-run cache, unbounded retained payload duplication, concurrency, I/O, lock, flush, or finalizer;
- opens D002-D017 or another parent;
- lacks independent Cheapapp and Restaurant D001/parent/analyzer/process packets or cannot prove exact affected equivalence.

## Handoff

- Attribution result: `A005_ATTRIBUTION_COMPLETE`.
- Selected cause: eager `ResolutionOutcome` JSON serialization owned by `(*resolutionOutcomeCollector).record -> marshalResolutionOutcome`, with resolved-path returned bytes discarded and all finalized outcomes serialized again by `projectResolutionOutcomes`.
- Exact owner: `E:\Anvien\internal\resolution\outcome.go::(*resolutionOutcomeCollector).record` (`83-106`, sampled unique path line `100`) and helper `marshalResolutionOutcome` (`220-226`).
- Complete call path: recorded above from `analyze.Run` through `ResolveBoundInto -> resolveCall`, all D001 outcome branches, collector/marshaler, final projection, Graph/reference/diagnostic carriers, analyzer result, persistence, snapshot, and CLI consumers.
- Next owner: Main Orchestration. Main alone records the packet and opens one fresh visible A005 Architect. No Planner or Coder may open from this report alone.
