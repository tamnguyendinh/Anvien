# Child 06A A005 Outcome Serialization Architecture

Verdict: `ARCHITECT_A005_READY_FOR_PLANNER`

## Scope And Authority

- Slice: `P2-A / A005 / B1-P1A-OP001 resolution / B2-P2A-A001-D001 resolve_calls`.
- Architect task: `01a03f5e-b844-7ac3-b23b-b9c9cdc0374d`.
- Authoritative clean docs/source HEAD: `0030c97bcd0fc6f94d4146842e6b0d21d8ac3f1d`.
- Accepted basis remains A003, target-separated:
  - Cheapapp D001/parent/analyzer/process: `3.447846300 / 20.472602300 / 93.531974900 / 95.630648200 s`; `calls=27890; files=887`.
  - Restaurant Manager D001/parent/analyzer/process: `9.401585300 / 20.850792800 / 98.020546700 / 101.096911900 s`; `calls=86030; files=1234`.
- State entering this decision: `A004_ROLLBACK_COMPLETE / A005_CURRENT_BASIS_RECORDED / A005_ATTRIBUTION_COMPLETE / ARCHITECT_ACTIVE / D001_STREAK_1`.
- This report owns architecture direction only. It performs no plan translation, source/test edit, build, target measurement, Supervisor review, disposition, detect, staging, or commit.
- Full raw authority was consumed in the required order: `AGENTS.md`; `working-rules`; `System-Architect`; binding `plan-rules.md`; all four standard Child 06A ledgers through EOF; `rp_child06a_a005_residual_attribution.md`; and the A004 Supervisor report solely to preserve its rejected-direction boundary.

## Current Graph And Blast Radius

Exactly one fresh `anvien analyze --force` ran at the authoritative HEAD before graph work. It passed with `2,250` scanned, `766` parsed code, `0` failed, and graph `124,189` nodes / `171,055` relationships. The graph is current and non-stale.

- `internal/resolution/outcome.go`: HIGH; `119` symbols, `87` inbound, `122` outbound, `117` local relationships, `148` unresolved sites, `11` linked flows, and `23` linked tests.
- `(*resolutionOutcomeCollector).record`: CRITICAL; `5` impacted symbols, `5` direct callers, `10` processes, `2` files.
- `marshalResolutionOutcome`: CRITICAL; `127` impacted symbols, `3` direct callers, `48` files, `25` modules, `52` processes.
- `(*resolutionOutcomeCollector).finalize`: CRITICAL; `124` impacted symbols, `2` direct callers, `50` files, `26` modules, `42` processes.
- `projectResolutionOutcomes`: CRITICAL; `124` impacted symbols, `2` direct callers, `26` modules, `42` processes.
- `internal/resolution/resolve.go`: HIGH; `200` symbols, `134` inbound, `289` outbound, `26` linked flows, and `40` linked tests.
- `ResolveBoundInto`: CRITICAL; `184` impacted symbols, `4` direct callers, `57` files, `28` modules, `31` processes.

These are blast-radius warnings, not edit prohibitions. They require a one-purpose implementation, exact owner limits, broad run-only regression coverage, and fail-closed STOP conditions.

## Preferred Direction

### Decision

Make the record-time encoding the **one canonical encoding event** for every successfully retained `ResolutionOutcome`. The run-scoped collector owns a private one-to-one sidecar from `SourceSiteID` to that canonical encoded string. Immediate diagnostics and final graph/reference projection consume the same string. `projectResolutionOutcomes` remains inside A005 production ownership, but only as a consumer and validator; it must stop calling `marshalResolutionOutcome` and must never regenerate missing bytes.

This converts the current record-time bytes from discarded/transient work into the canonical bytes used by all downstream carriers. It removes projection-time re-serialization and equal-duplicate re-serialization without moving validation or initial marshal failure timing.

There is one semantic authority and one byte authority per retained source site:

```text
validated cloned ResolutionOutcome
        +
record-time marshalResolutionOutcome bytes
        =
one private finalized outcome tuple
        -> immediate diagnostic when required
        -> final relationship/reference/diagnostic parity validation
        -> graph carriers
```

There is no status-specific fallback encoder, projection encoder, recovery encoder, second byte cache, or alternate serialized shape.

### Private Representation

The implementation must keep the existing semantic map and add one private sidecar on the same run-scoped collector:

```text
resolutionOutcomeCollector
  bySourceSite:        SourceSiteID -> cloned ResolutionOutcome
  encodedBySourceSite: SourceSiteID -> canonical JSON string
  err:                 first retained error
```

`finalize` returns one private `finalizedResolutionOutcomes` bundle containing:

```text
values:                SourceSiteID-sorted cloned []ResolutionOutcome
encodedBySourceSite:   the sealed collector sidecar, read-only after finalize
```

The bundle is private to `internal/resolution`; it changes no exported type, public function signature, persisted schema, JSON tag, or analyzer result shape. `ResolveBoundInto` unwraps `values` only when assigning `Result.ResolutionOutcomes`; it passes the complete private bundle to projection.

The sidecar must be sealed by control flow: no record call occurs after `finalize`. `finalize` must fail closed if the retained semantic inventory and canonical-byte inventory are not exactly one-to-one. `projectResolutionOutcomes` must fail closed on a missing canonical byte; it must not call the marshaler as a fallback.

## Canonical Ownership By Status Family

| Status family | Record-time behavior | Immediate carrier | Final carrier | Canonical byte owner |
|---|---|---|---|---|
| Repository `resolved_internal` | clone, validate, conflict-check, store, encode once | none; existing caller may ignore returned string | relationship and reference evidence | collector sidecar; projection consumes it |
| Intrinsic `resolved_internal` | same one encoding event | none | relationship/reference evidence where applicable | collector sidecar; D003 remains queued and preserve-only as a control row |
| TypeScript `resolved_external` | same one encoding event | none | relationship/reference evidence | collector sidecar; the current unused returned string becomes non-redundant rather than removed |
| Repository `unresolved` | same one encoding event | on first successful add, the exact string becomes `Diagnostic.Note` immediately | diagnostic/outcome byte-parity validation | the same collector string; diagnostic assignment shares it |
| TypeScript `capability_unavailable` | same one encoding event | on first successful add, the exact string becomes `Diagnostic.Note` immediately | diagnostic/outcome byte-parity validation | the same collector string |
| TypeScript profile-excluded or meaning-mismatch mapped to `unresolved` | same one encoding event | on first successful add, the exact string becomes `Diagnostic.Note` immediately | diagnostic/outcome byte-parity validation | the same collector string |
| Equal duplicate for any family | validate and compare with the stored clone; return the stored clone/string and `added=false` | no duplicate diagnostic | no duplicate outcome/carrier | existing tuple; no second marshal |
| Conflicting duplicate | retain the current first error rule, return the previous tuple and `added=false` | none | `finalize` fails closed | previous tuple remains immutable; conflict is not encoded as a new outcome |

Repository-resolved and resolved-external encodings are therefore not delayed or removed. They are made non-redundant: the bytes created at the current record boundary are retained for the final projection that already requires them.

## Exact Behavioral Contract

### Clone, validation, conflict, and `added`

1. Clone the incoming outcome before any validation or retention, exactly as now.
2. Run `validateResolutionOutcome` before duplicate lookup, exactly as now. Invalid input is not retained, returns `added=false`, and records the first error.
3. For an existing `SourceSiteID`, compare the stored cloned semantic outcome with the incoming cloned outcome exactly as now.
4. On conflict, call the existing first-error retention path and return the prior cloned outcome, prior canonical bytes (or empty bytes when the prior marshal already failed), and `added=false`.
5. On an equal duplicate, return the prior cloned outcome, prior canonical bytes, and `added=false`. Re-marshaling is forbidden because the stored clone is immutable, the object graph has no custom `MarshalJSON`, and the first successful marshal is deterministic.
6. On a unique valid outcome, retain the cloned semantic outcome before attempting the marshal, preserving the current internal lifecycle. Marshal through the unchanged `marshalResolutionOutcome` helper at the same record boundary. Store canonical bytes only on success. Return `added=true` only on that success, exactly as now.

### Marshal failure timing and error retention

- `marshalResolutionOutcome` remains the sole encoder and retains its exact body, signature, contextual error text, JSON field order, and byte format.
- The first marshal attempt remains inside `record`, after semantic validation and unique retention, so record-time error timing is not moved to projection.
- A marshal failure continues to return an empty string and `added=false`, suppressing immediate diagnostic emission. The collector retains the first error, and `finalize` returns it before projection.
- If the first marshal failed, later duplicates do not retry or hide it; the already retained first error remains authoritative.
- If the first marshal succeeded, equal-duplicate re-marshaling cannot discover a new deterministic encoding error because the stored clone is immutable and none of its types has a custom marshaler. Skipping that redundant call changes no error precedence.
- Missing or extra sidecar state is an internal invariant violation and must fail closed. Production must not regenerate bytes to recover.

### Finalization, ordering, and determinism

- `finalize` continues to revalidate retained semantic outcomes, clone them, and sort them only by `SourceSiteID`.
- Map traversal never defines output order.
- The canonical byte sidecar is keyed only by exact `SourceSiteID`; it adds no sort key, evidence order key, or A004-style ordering machinery.
- `projectResolutionOutcomes` keeps its existing duplicate-finalized-site check, resolved/unresolved status checks, graph relationship traversal, reference-index traversal, resolved-vs-diagnostic overlap check, and diagnostic/outcome byte-parity check.
- Relationship/reference evidence continues through the existing `mergeExportBindingEvidence` behavior. A004 code and ordering remain fully restored and preserve-only.
- Repeated runs on the same accepted input must produce identical structured outcome order, `Diagnostic.Note`, `Evidence.Note`, complete ordered evidence slices, Graph JSON, and public output.

## Why `projectResolutionOutcomes` Is Inside A005

It is inside production ownership because leaving its current marshal call would preserve the duplicate serialization A005 is meant to remove. Its architecture role changes narrowly from **encoder + carrier validator** to **canonical-byte consumer + carrier validator**.

It is not allowed to own or create canonical bytes. It may only:

- verify one-to-one finalized outcome/encoding coverage;
- build its current cloned semantic lookup;
- attach the supplied bytes to relationship/reference evidence;
- validate exact diagnostic byte parity and resolved/unresolved exclusivity.

All other projection behavior stays unchanged. D004 `project_resolution_outcomes` remains a queued unchecked benchmark child; A005 does not open D004 as a separate optimization or alter its general algorithms. Any observed D004 timing change is a recorded sibling effect of closing the D001 encoding lifecycle, not a D004 disposition.

## Exact Allowed Production And Test Ownership

### Production

Only these production surfaces may change:

1. `internal/resolution/outcome.go`
   - private `resolutionOutcomeCollector` state: add only the canonical encoded sidecar;
   - `newResolutionOutcomeCollector`: initialize that sidecar;
   - `(*resolutionOutcomeCollector).record`: store/reuse one canonical string while preserving its signature and all return semantics;
   - one new private `finalizedResolutionOutcomes` representation;
   - `(*resolutionOutcomeCollector).finalize`: return the private bundle, preserve validation/cloning/sorting/error precedence, and verify one-to-one coverage;
   - `projectResolutionOutcomes`: accept the private bundle, remove only its marshal call/local encoded-map construction, and consume the sealed canonical sidecar without fallback.
2. `internal/resolution/resolve.go`
   - `ResolveBoundInto` only in the existing `finalize -> project -> error/result ResolutionOutcomes` wiring block. Its exported signature, loop order, branch order, metrics, graph/reference ownership, error return shapes, and all other lines are preserve-only.

`marshalResolutionOutcome`, `validateResolutionOutcome`, `cloneResolutionOutcome`, all four outcome construction methods, `recordTypeScriptLookup`, `emitTypeScriptOutcomeDiagnostic`, `emitUnresolvedReference`, `projectReferenceIndexOutcomes`, and `resolutionOutcomeDiagnosticSites` are call/run-only and must not be edited. In particular, keeping the existing `record` return signature avoids any immediate-diagnostic caller edit.

### Tests, only after production is correct

Allowed test owners are exactly:

1. New `internal/resolution/outcome_serialization_test.go` for A005-only focused tests.
2. Existing `internal/resolution/p6c3_structured_outcome_test.go` only for the smallest mechanical adaptation of its three direct private `finalize` calls and one direct private `projectResolutionOutcomes` call to the finalized private bundle. Its existing assertions and every other byte remain unchanged.

All other source/test/script/report/ledger/target files are preserve-only or run-only. No edit is allowed in `internal/resolution/types.go`, `internal/resolution/emit.go`, graph or graphhealth types/policy/decoder, analyzer result types, persistence/readers, CLI, instrumentation, the reusable A00x script, A001-A003 owners, restored A004 owners, D002-D017 owners, P3, or Child 07.

## Focused Test Contract

Production must be implemented and inspected first. Only then may the two authorized test owners change.

The new focused test file must prove:

- one canonical byte string for repository-resolved, intrinsic-resolved, TypeScript resolved-external, repository-unresolved, TypeScript capability-unavailable, profile-excluded, and meaning-mismatch cases;
- repository-unresolved and non-resolved TypeScript `Diagnostic.Note` bytes equal the pre-A005 `json.Marshal` bytes exactly and also equal final projection parity bytes;
- resolved relationship and both reference-index carrier bytes equal the same canonical string;
- first-add and equal-duplicate `added` behavior, no duplicate diagnostic, and prior canonical tuple reuse;
- conflicting `SourceSiteID` fails closed with unchanged first-error precedence;
- input, returned, finalized, target, proof evidence, authority, and declaration-range mutation cannot alter the retained semantic clone or canonical bytes;
- marshal failure at record time using a naturally unsupported JSON float such as non-finite `graph.Evidence.Weight`: empty bytes, `added=false`, no diagnostic, retained contextual error, and `finalize` failure before projection;
- missing/extra finalized encoding coverage fails closed with no re-encoding fallback;
- duplicate finalized source sites, resolved-plus-unresolved overlap, diagnostic/outcome payload drift, and non-resolved reference carrier remain rejected;
- SourceSiteID-sorted structured outcomes, graph/reference traversal behavior, complete ordered evidence, and repeated-run determinism remain exact.

The test-local expected-byte/oracle logic must not become a production fallback or second production encoder.

## Preserved Cross-Surface Invariants

- Exact `ResolutionOutcome` schema, JSON tags, field order, status/stage/site/target/reason/proof/authority shapes, and contextual marshal errors.
- Exact `SourceSiteID` identity, duplicate/conflict detection, first-error retention, immutable clone semantics, and `added` behavior.
- Exact immediate `Diagnostic.Note` and final relationship/reference `Evidence.Note` byte parity.
- Exact diagnostic bucket semantics, A002 write-through appender behavior, count merge, order, `[]Diagnostic` graph-property representation, and run-scoped `O(T)` appender lifecycle.
- Exact A003 canonical single-interpretation decoder tuple/policy/fail-closed behavior; no decoder or diagnostic policy edit.
- Exact graph nodes, relationships, IDs, labels, properties, counts, reference indexes, resolution outcome inventory, and ordered evidence.
- Exact analyzer `Result.ResolutionOutcomes` and `Result.Graph` carriage, Graph JSON, Ladybug/native persistence/readback, affected readers, normal CLI stdout/stderr and public output, and benchmark metrics.
- Deterministic replay, source freshness/invalidation, failure propagation, transaction rollback, temporary-file flush/close/rename, publication visibility, and known preserve-only golden-failure truth.
- A001 import-claim index/lookups, A002 diagnostic appender, A003 decoder, WAL cleanup, P1 timing instrumentation, accepted target workloads/options/denominators, and restored A004 code/test identities.

## Resource Boundary

Let `U` be the number of unique valid retained outcomes and `B` be the sum of the byte lengths of their canonical JSON strings.

- Collector sidecar complexity: `O(U + B)` retained state, with exactly one map entry and at most one encoded backing payload per retained `SourceSiteID`.
- The semantic inventory remains the existing `O(U)` cloned outcome map. The sidecar does not duplicate outcome objects.
- Assigning a Go string to an immediate diagnostic or later evidence copies only a string header; the canonical byte backing storage is shared. Repeated relationship/reference carriers may add headers but must not allocate another encoded payload.
- Lifetime: created with the emitter, populated during one sequential `ResolveBoundInto`, sealed at `finalize`, consumed during projection, and the map becomes unreachable when `ResolveBoundInto` returns. Canonical byte backing storage survives only where the returned graph carriers retain it.
- The design retains bytes earlier than current projection, but the retained payload is bounded by the exact final unique outcome payload that projection already needs. There is no global/cross-run cache, I/O, goroutine, lock, concurrency, flush, finalizer, TTL, or unbounded duplication.

Any implementation retaining more than one encoded payload per unique outcome, copying payload bytes per carrier, or keeping the sidecar beyond the run is outside this architecture.

## Qualitative Expected Effect

No percentage, tolerance, or exact saving is predicted. Profile samples are overlapping causal evidence only.

| Target | D001 expectation | Parent expectation | Analyzer/process expectation |
|---|---|---|---|
| Cheapapp | A modest reduction is expected from removing equal-duplicate marshals and making resolved-path record bytes useful rather than transient. The sampled family is present but smaller than Restaurant's, so the retained D001 effect may be limited. | `resolution` should inherit the D001 saving and also avoid the final projection re-marshal, without transferring work to another resolution child. | Allocation/GC pressure should decrease or stay bounded; analyzer and process benefit may be small. A005 is retainable only if the independent packet shows lower child, lower parent, and retained lower process time with no unexplained sibling regression. |
| Restaurant Manager | A clearer D001 reduction is expected because the record/marshal family is more prominent in the retained profile and the accepted D001 workload is larger. This remains qualitative, not a promised delta. | `resolution` should inherit the D001 saving and the same no-re-encode projection benefit. | The larger causal family offers a better chance of observable analyzer/process retention, but process improvement is not assumed. The same independent lower child/parent/process rule applies. |

## Rejected Alternatives

1. **Delay all encoding to final projection.** Rejected because repository-unresolved and non-resolved TypeScript diagnostics require bytes immediately, and resolved-family marshal failure would move from record time to projection.
2. **Delay only resolved-family encoding while keeping unresolved encoding at record time.** Rejected because it creates two lifecycle/encoding authorities by status, changes resolved error timing, and leaves byte parity dependent on separate paths.
3. **Let immediate callers marshal locally and projection marshal again.** Rejected because it duplicates encoding authority and can drift byte/error behavior.
4. **Keep current record-time marshal but continue discarding resolved bytes.** Rejected because it preserves the attributed redundant work.
5. **Embed cached bytes in exported `ResolutionOutcome`, add a global/cross-run cache, or attach bytes to public graph/reference types.** Rejected because it changes public/private ownership, extends lifetime, and risks unbounded or externally visible state.
6. **Revive A004 exact-tuple/cached-order/single-sort work.** Rejected by authority: A004 passed correctness but increased analyzer/process elapsed on both targets, received `NO_KEEP`, and was rolled back. A005 contains no ordering optimization.

## Required Future Execution And Validation Sequence

1. Main verifies this report and, if accepted, releases one visible Planner to translate it without redesign. Planner cannot open Coder.
2. A later separate Coder repeats `anvien --help`, exactly one fresh `anvien analyze --force`, `file-detail` for both production files, and exact upstream impact for every symbol it will edit. The HIGH/CRITICAL warnings above must be recorded, not treated as bans.
3. Coder changes only the authorized production surfaces first and inspects the exact production diff. If any additional owner is needed, STOP before tests.
4. Only after production behavior is correct may Coder add the new focused test file and mechanically adapt the one existing direct private projection call.
5. Before build, run holder/lock/process preflight, including `anvien doctor locks --repo E:\Anvien --json` and `anvien doctor processes --json`; prove and terminate only actual build-output holders, and start no build until all such holders are gone.
6. Run the canonical full build and require exit `0` before executing tests:

```text
pwsh -NoLogo -NoProfile -NonInteractive -ExecutionPolicy Bypass -File E:\Anvien\scripts\full-build.ps1
```

7. After full-build PASS, run the new A005 focused tests plus these existing focused regressions:
   - `TestP6C3P5ProofNestingConflictAndImmutableReplay`;
   - `TestP6C3FinalOutcomeStatusPrecedenceAndCarriageMatrix`;
   - `TestResolveAttachesSourceBackedUnresolvedDiagnostics`;
   - `TestP6BResolverCarriesCatalogValidationFailuresPerSourceSite`;
   - `TestP6DStructuredOutcomeProjectionPreservesGraphJSONAndHealthParity`;
   - `TestP6C3AnalyzeResultPreservesFinalOutcomesAndGraphCarriage`;
   - `TestP6C3AnalyzeCapabilityOutcomesRetainAcceptedAuthorityStatus`;
   - `TestP6C3ResolutionOutcomesPreserveGraphJSONAndLadybugParity`;
   - `TestP6C3NativeLadybugResolutionOutcomeReadback`.
8. Run package regressions for `internal/resolution`, `internal/graphhealth`, `internal/analyze`, `internal/lbugload`, and `internal/lbugnative`. The known `TestProofBasedCallAccessGoldenCorpus` mismatch may be recorded only if its payload and exact preserve-only baseline remain unchanged; the package is never called PASS. Any new or changed failure blocks A005.
9. Run scoped diff-check. No detect, stage, or commit occurs until after later measurements, Supervisor, and Main disposition.

## Unchanged A00x Candidate And Measurement Contract

- Reuse `scripts/build-a00x-benchmark.ps1` unchanged with explicit A005 attempt, overlay, output, mapped-source, candidate-source, native, and hash inputs. Do not repeat general build-interface discovery or edit either canonical build script.
- Because A005 narrowly changes `ResolveBoundInto` in the overlay-mapped `resolve.go`, the A005 `resolve.go` measurement overlay must be mechanically regenerated/verified from the exact A005 production source plus the already accepted 17-child instrumentation. A stale A003/A004 overlay that erases the A005 wiring is a mandatory STOP. The `types.go` instrumentation mapping remains unchanged unless exact compilation evidence proves otherwise.
- The frozen candidate identity is accepted A003/WAL state plus only A005-authorized production/test bytes and the accepted measurement overlay. No rejected A004 bytes may enter it.
- Measure Cheapapp once with its accepted command/options and Restaurant Manager once with its accepted command/options and exact one `electron/renderer/src/api/userApi.ts` exclusion. Each uses its own accepted A003 packet as `before`; never rerun, average, or combine the accepted bases.
- Each independent packet must record D001, parent, analyzer, process wall, `30/30` top-level operations, `17/17` children, exact calls/files, interval conservation/zero overlap, workload, diagnostics/outcomes, full ordered evidence, Graph JSON, graph/DB readback, stdout/stderr, semantic counters, `startAllocBytes`, `endAllocBytes`, `maxObservedSys`, and the `O(U+B)` one-payload lifecycle proof.
- Only after both valid packets exist may one fresh visible A005 Supervisor review the exact candidate and affected correctness/equivalence/output/lifecycle/resource boundary. Main alone decides `KEEP`, `REWORK`, `ROLLBACK`, or subsequent streak effect.

## Exact Rollback

Rollback only A005-owned bytes:

- remove the collector canonical-byte sidecar and its initialization;
- restore the prior `record` duplicate/unique marshal behavior;
- remove the private finalized bundle and restore the prior `finalize` return;
- restore projection-time `marshalResolutionOutcome` and its local encoded map;
- restore only the narrow `ResolveBoundInto` finalize/project/result wiring;
- remove the new focused test file and restore only the mechanical existing-test `finalize`/projection call adaptations;
- remove only the A005 frozen/overlay packet if the attempt is rejected, preserving accepted campaign tooling and historical evidence.

Do not reset, checkout, stash, overwrite whole files, or disturb A001-A003, WAL, P1 instrumentation, reports, ledgers, target work, or protected/user changes.

## Mandatory STOP Conditions

STOP and return to Main for a fresh decision if any of the following occurs:

- any production/test file or production symbol outside the exact allowed list is required;
- `marshalResolutionOutcome`, immediate diagnostic callers, schema/type/JSON tags, diagnostic policy/appender/decoder, graph/reference/public types, persistence/readers, CLI, instrumentation, or the A00x script must change;
- record-time semantic validation or initial marshal timing moves later;
- conflict detection, first-error retention, `added`, clone immutability, or SourceSiteID behavior changes;
- unresolved or non-resolved TypeScript diagnostic bytes are not immediately byte-identical;
- project has a re-encoding fallback, a second encoder, missing-byte tolerance, or altered carrier validation;
- outcome/evidence/diagnostic ordering or A004-restored behavior changes;
- retained resources exceed `O(U+B)`, duplicate payload bytes per carrier/source site, survive the run, or require global state, concurrency, I/O, locks, flush, or finalizers;
- the canonical full build fails, any new/changed test failure appears, or the known golden failure changes;
- the A005 overlay does not contain the exact production wiring, the unchanged script contract cannot build it, or candidate identity is mixed;
- either target packet is missing, incomparable, uses a changed workload/denominator, lacks exact equivalence, or shows an unexplained resource/sibling regression;
- D002-D017, another parent, P3, or Child 07 is opened;
- a later unsuccessful A005 attempt is incorrectly treated as terminal streak `3`; the current streak is `1`, so one unsuccessful A005 would make it `2` only.

## Handoff

- Architecture decision: one record-time canonical JSON encoding per retained `SourceSiteID`, owned by a private run-scoped collector sidecar; immediate diagnostics and final projection share it; projection never encodes or falls back.
- Missing measurement input: none. Existing attribution is sufficient to authorize Planner translation.
- Residual open architecture questions: none inside A005.
- Next owner: Main Orchestration for report verification. Planner remains locked until Main accepts this report. Coder remains locked until a later visible Planner translation and a separate Main release.
- Commit reference: none; this lane was explicitly forbidden to stage or commit.
