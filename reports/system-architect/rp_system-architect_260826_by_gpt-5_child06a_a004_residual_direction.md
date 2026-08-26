# System Architect Report — Child 06A A004 Residual Direction

Date: `2026-08-26`
Role: fresh visible A004 System Architect
Scope: `P2-A / A004 / B1-P1A-OP001 resolution / B2-P2A-A001-D001 resolve_calls`
Verdict: `ARCHITECT_A004_OWNER_APPROVED_READY_FOR_PLANNER`
Decision boundary: Owner-approved architecture authority released only for visible Planner translation; no source/test edit, build, test, target measurement, Supervisor, detect, stage, commit, P3, Child 07, or Coder authorization
Current repository HEAD / graph identity: `54b2bd9b8b10b84861abd81f1686874cbd23c16c`
Accepted performance checkpoint: A003 `b6bf45bce95323aa6b53b182edfea8628bd8b463`
Accepted storage checkpoint: `0f3a572331dd23d17688886fcbfebeb7d37ee35d`

## 1. Resume authority and current control basis

The intervening Ladybug force-cleanup/WAL defect is accepted as independently reviewed and committed at `0f3a572331dd23d17688886fcbfebeb7d37ee35d` (`fix(storage): reset Ladybug artifact family on force`). Current docs-sync HEAD is `54b2bd9b8b10b84861abd81f1686874cbd23c16c`. This lane did not audit or rerun the WAL diagnosis, fix, build, tests, or Supervisor review. The storage fix changes no A004 performance number, D001 streak, parent/child checkbox, or queue.

The one mandated post-resume graph refresh completed successfully at current HEAD:

- `anvien analyze --force`: exit `0`;
- files `2241` scanned / `766` parsed code / `0` failed;
- graph `124069` nodes / `170935` relationships;
- current file-detail identity: indexed commit and current commit both `54b2bd9b8b10b84861abd81f1686874cbd23c16c`, `stale=false`, `changedSinceAnalyze=false` for the selected owner.

Campaign controls remain unchanged:

| Control | Current state |
|---|---|
| Implementation slice | `P2-A`, unchecked and active |
| Active parent | `B1-P1A-OP001 resolution`, unchecked |
| Active child | `B2-P2A-A001-D001 resolve_calls`, unchecked |
| D001 unsuccessful-attempt streak | `0` |
| Remaining child queue | `D002-D017`, queued/unopened |
| A003 decoder | accepted preserve-only canonical single interpretation |
| Planner/Coder | Planner released for exact plan translation; Coder remains locked pending completed translation, Main verification, and separate release |

Accepted A003 elapsed bases remain separate and controlling:

| Target | D001 | Parent | Analyzer | Process | Denominator |
|---|---:|---:|---:|---:|---|
| Cheapapp | `3.447846300 s` | `20.472602300 s` | `93.531974900 s` | `95.630648200 s` | `calls=27890; files=887` |
| Restaurant Manager | `9.401585300 s` | `20.850792800 s` | `98.020546700 s` | `101.096911900 s` | `calls=86030; files=1234` |

These targets are never averaged or combined. Elapsed wall time controls future disposition; profile samples below are causal clues only.

## 2. Exact residual cause

### Verified current-source facts

The new residual owner is `internal/resolution/export_binding_proof.go`, not the accepted A003 diagnostic decoder.

Current export-binding evidence projection and merge perform the following work:

1. `appendExportBindingEvidence` constructs JSON-backed terminal, hop, and failure `graph.Evidence` records from already-typed proof data.
2. It sorts the newly projected evidence at line `164`.
3. It immediately calls `mergeExportBindingEvidence` at line `165`.
4. The merge exact-tuple deduplicates existing/incoming evidence, separates generic and export-binding evidence, then sorts the projected evidence again at line `218`.
5. `sortExportBindingEvidence` uses `sort.SliceStable`. Every comparator invocation calls `exportBindingEvidenceOrderFor` for both compared records.
6. `exportBindingEvidenceOrderFor` decodes the full JSON `Evidence.Note` to recover ordinals:
   - terminal note: `json.Unmarshal` at line `259`;
   - hop note: `json.Unmarshal` at line `265`;
   - failure note: `json.Unmarshal` at line `271`.

Therefore ordering metadata that existed as typed values during projection is serialized, discarded, then repeatedly reconstructed from JSON inside the stable-sort comparator. Comparator work grows with comparisons rather than evidence count, and the newly projected set is sorted twice.

`graph.Evidence` remains the public persisted shape `{Kind, Weight, Note}`; it has no ordering fields. A004 does not change that contract.

### Profile inference, kept separate by target

Bounded `go tool pprof` reads used the already accepted A003 CPU artifacts and focused only stacks containing `resolution.resolveCall`.

| Target | `resolveCall` cumulative sample | `sortExportBindingEvidence` cumulative sample | `exportBindingEvidenceOrderFor` cumulative sample | Comparator split |
|---|---:|---:|---:|---:|
| Cheapapp | `3.12 s` | `1.23 s` | `1.22 s` | first sort `0.50 s`; merge sort `0.73 s` |
| Restaurant Manager | `8.29 s` | `0.21 s` | `0.20 s` | first sort `0.04 s`; merge sort `0.16 s` |

These samples overlap and are non-additive. They are not D001 elapsed values, do not predict an exact saving, and cannot prove future improvement. They do verify that the same repeated-ordering stack occurs beneath D001 on both targets independently.

Accepted profile identities remain separate:

- Cheapapp: `E:\Anvien\.tmp\child06a_a003_cheapapp_frozen_17child_20260826\candidate\resolution.cpu.pprof`, `65,363` bytes, SHA-256 `806FAD79B12C567618778A8232E8521991882CE6C745AC2EBBE7B09ECC267107`;
- Restaurant Manager: `E:\Anvien\.tmp\child06a_a003_restaurant_manager_frozen_17child_20260826\candidate\resolution.cpu.pprof`, `56,397` bytes, SHA-256 `4272FFAB0ABE4C14C470A139A31A601F31323B1373E0FA2986A09935456FA3E2`.

### Missing evidence

- No exact count of export-binding evidence records or comparator invocations per target is available.
- No exact split exists between first-sort and later relationship-coalescing invocations beyond sampled stacks.
- No A004 candidate allocation, elapsed, output-equivalence, or resource result exists.
- No numeric gain is predicted.

Those missing counters do not prevent a safe architecture decision because the repeated decode/double-sort mechanism is directly proven in current source and appears in both retained profiles. They do prevent any percentage or threshold claim.

## 3. Exact owner and complete call path

### Primary D001 upstream path

`internal/analyze.Run -> runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> w.files -> ir.Calls -> resolveCall`.

Within `resolveCall`, export-binding evidence is attached through two current branches:

1. Unresolved semantic-result branch:

   `resolveCall:536-560 -> appendExportBindingEvidence:540 -> export evidence projection -> ordering/merge -> emitUnresolvedReference:560`.

2. Resolved semantic-result branch:

   `resolveCall:567-602 -> retainedExportResolutionForScopedBinding:578-583 when needed -> appendExportBindingEvidence:586 -> export evidence projection -> ordering/merge -> emitReference:588-602`.

### Exact residual stack

`appendExportBindingEvidence:97-166 -> marshal terminal/hop/failure notes -> sortExportBindingEvidence:164 -> mergeExportBindingEvidence:165/195-220 -> sortExportBindingEvidence:218/231-246 -> comparator -> exportBindingEvidenceOrderFor:248-277 -> json.Unmarshal`.

### Downstream/shared consumers

- Resolved relationship coalescing:

  `resolveCall -> emitter.emitReference -> emitter.emitRelationship -> mergeRelationship:683-732 -> mergeExportBindingEvidence:731`.

- Final resolution-outcome projection:

  `ResolveBoundInto:110-116 -> projectResolutionOutcomes -> each resolved graph relationship -> mergeExportBindingEvidence at outcome.go:426-430`, plus `projectReferenceIndexOutcomes -> each BySourceScope/ByTargetDef reference -> mergeExportBindingEvidence at outcome.go:467-471`.

- `resolveAccess` uses `appendExportBindingEvidence` at `resolve.go:680/722`. D002 remains unopened; this is a required preserve/regression surface, not a second A004 target.

### Publication/persistence tail

`ResolveBoundInto` returns the mutated graph and projected reference/outcome carriers -> `analyze.Run:374-377` installs the result -> later MRO/communities/processes/semantic phases consume the same graph -> `Graph.Compact` -> `loadGraph` at `analyze.go:472-474` persists the graph through the configured Ladybug runner -> `writeGraphSnapshot` at `analyze.go:503-506` publishes canonical Graph JSON -> existing CLI publication emits the accepted public result. A004 edits none of those downstream owners; ordered evidence parity is what preserves them.

The helper is therefore shared. A004 remains selected by the measured D001 stack, but exact parity must cover the sibling access and final-projection consumers.

## 4. Anvien blast radius

Current graph evidence for `internal/resolution/export_binding_proof.go`:

- containing-file risk: `HIGH`;
- `98` symbols;
- fan-in `68`, fan-out `48`, local relationships `70`;
- unresolved source sites `101`, linked flows `1`, linked tests `23`;
- functional area `resolution`;
- graph current/non-stale and `changedSinceAnalyze=false`.

Exact upstream impacts with tests included are all `CRITICAL` scope warnings:

| Exact helper | Impacted symbols | Direct callers | Affected files | Modules | Processes |
|---|---:|---:|---:|---:|---:|
| `appendExportBindingEvidence` | `36` | `5` | `14` | `7` | `39` |
| `mergeExportBindingEvidence` | `51` | `5` | `19` | `7` | `49` |
| `sortExportBindingEvidence` | `23` | `2` | `10` | `2` | `34` |
| `exportBindingEvidenceOrderFor` | `14` | `1` | `6` | `1` | `19` |

The widest exact owner, `mergeExportBindingEvidence`, reaches these `19` files: `internal/analyze/analyze.go`, `internal/analyze/analyze_test.go`, `internal/analyze/legacy_resolver_conversion_test.go`, `internal/analyze/p6b_tsstdlib_test.go`, `internal/analyze/pipeline_parity_test.go`, `internal/cli/command.go`, `internal/graphaccuracy/access_candidate.go`, `internal/resolution/definition_collision_test.go`, `internal/resolution/emit.go`, `internal/resolution/export_binding_proof.go`, `internal/resolution/export_binding_proof_test.go`, `internal/resolution/external_symbol.go`, `internal/resolution/legacy_p7_conversion_test.go`, `internal/resolution/outcome.go`, `internal/resolution/p3c_binding_occurrence_test.go`, `internal/resolution/p6c3_structured_outcome_test.go`, `internal/resolution/parser_integration_test.go`, `internal/resolution/resolution_test.go`, and `internal/resolution/resolve.go`.

Named affected process samples include `ResolveBoundInto -> RelationshipCallName`, `AppendExportBindingEvidence -> CleanExportTablePath`, `ResolveAccess` flows, `ResolveTypeAnnotation` flows, analyzer `Run` flows, and CLI `Main`/`NewAnalyzeCommand` flows. These are regression/consumer evidence, not additional edit ownership.

The upstream helper family therefore reaches:

- `resolveCall` and sibling `resolveAccess`;
- resolution outcome/reference projection;
- the analyzer result path;
- existing export-binding proof, parser-integration, and legacy-conversion tests.

The file-level `HIGH` and helper-level `CRITICAL` results are scope warnings, not edit prohibitions. They require the one-file production boundary and full output/order regressions below. Leaf-helper locality must not be mistaken for system-level isolation because merge ordering is consumed by resolved relationships, reference indexes, analyzer/CLI flows, and final Graph JSON.

## 5. One preferred A004 direction

### Merge-owned decorate–sort–undecorate with one canonical key extraction

Architecture rule:

`one unique export-binding evidence record -> at most one order-key extraction per merge boundary -> one final stable sort -> unchanged graph.Evidence output`.

The preferred direction is to make `mergeExportBindingEvidence` the sole final-order owner and replace decode-in-comparator ordering with one private transient decorated representation inside `internal/resolution/export_binding_proof.go`:

- exact-tuple deduplication and generic/projected partitioning occur first, preserving current first-seen behavior;
- each remaining unique projected record is passed through the existing semantic order extractor exactly once, after deduplication;
- each transient record carries the unchanged `graph.Evidence` plus that cached `kindRank / proofOrdinal / hopOrdinal / note` key;
- generic evidence retains its current first-seen stable order;
- projected evidence receives exactly one final stable sort using only cached scalar/string keys;
- the transient key is discarded before return; returned values remain exactly `[]graph.Evidence`;
- the redundant pre-sort in `appendExportBindingEvidence` is removed because the merge owns the single final ordering decision.

The comparator must perform no JSON decoding, marshaling, validation, or mutation. `exportBindingEvidenceOrderFor` remains the one order-key authority for both newly projected and previously stored evidence; A004 does not add a producer-side typed-key fast path, private overload carrying separate keys, or second ordering interpretation.

Implementation freedom is limited to a private transient record/helper arrangement in the same file. A decorate–stable-sort–undecorate implementation is compliant only if it calls the canonical extractor no more than once for each deduplicated projected record and preserves the exact current comparison tuple. No public type, persisted field, cross-call cache, producer-carried alternate key, or secondary ordering authority is allowed.

## 6. Why this direction is materially new

- A001 indexed import claims to remove global import traversal.
- A002 introduced a run-scoped write-through diagnostic appender to avoid repeated prior-diagnostic normalization.
- A003 made the outer diagnostic Note decoder canonical and single-interpretation.
- A004 changes export-binding evidence ordering complexity and ownership. It neither changes import lookup, diagnostic appending, nor diagnostic decoding.

A004 does not edit, wrap, bypass, duplicate, or cosmetically re-express `decodeStructuredResolutionOutcome`.

## 7. Alternatives rejected

### Go same-package function index

Restaurant's profile has a large `resolveGoSamePackageFunction` stack, but Cheapapp has no corresponding material retained stack. This may be a later residual direction, but it cannot be the preferred A004 direction because current evidence does not support improvement on both targets independently.

### Another diagnostic decoder fast path or structured-metadata bypass

Rejected because A003 already owns and accepted the canonical single-interpretation decoder. A second decoder, fallback, producer-side bypass, or cosmetic decoder rewrite would repeat or undermine A003.

### Remove sorting and trust construction order

Rejected. Projection emits proofs in proof-local order while the contract groups by evidence kind and ordinals. Relationship coalescing merges independently created slices. Eliminating the canonical final ordering would break deterministic output.

### Only remove the first redundant sort

Rejected as incomplete. It removes one sort but leaves full JSON decoding inside every comparison of the remaining sort, which is the directly profiled core mechanism.

### Producer-carried typed order keys

Rejected for A004. Projection already has ordinals in typed form, but passing those keys through a private overload while arbitrary existing evidence still uses decoded keys would create two order-key derivation paths and additional choreography across the shared merge. The selected post-dedup decoration removes comparator amplification with one canonical extractor, keeps malformed-note fallback identical, and is sufficient without a second authority.

### Add order fields to public `graph.Evidence` or retain a run-wide cache

Rejected. Either changes a persisted/public graph contract or introduces lifecycle/invalidation ownership beyond the selected local helper. Neither is needed.

## 8. Exact allowed and preserve-only surfaces

### Allowed production surface

Only `internal/resolution/export_binding_proof.go`:

- `appendExportBindingEvidence`;
- `mergeExportBindingEvidence`;
- `sortExportBindingEvidence`;
- `exportBindingEvidenceOrderFor`;
- existing private order/key types; and
- at most the private transient helper/type required to implement one-extraction/one-final-sort.

All existing function signatures remain unchanged. No private overload carrying alternate producer-computed keys and no exported symbol are authorized.

### Authorized test surface, after production is correct

Only `internal/resolution/export_binding_proof_test.go`, appended with A004-specific parity/adversarial cases.

### Run-only regression surfaces

- existing export-binding proof tests;
- focused call/access resolution tests that consume semantic export proofs;
- outcome projection and Graph JSON parity tests;
- full `internal/resolution`;
- existing A003 graphhealth parity tests as needed to prove no diagnostic/output drift.

### Preserve-only production/test surfaces

- `internal/graph/types.go` and public `graph.Evidence`;
- `internal/resolution/resolve.go`, `emit.go`, `outcome.go`, `indexes.go`, export traversal/result schemas, and resolution instrumentation;
- `internal/graphhealth/diagnostics.go` and `diagnostics_test.go`, including all accepted A002/A003 bytes;
- persistence, Ladybug/native readers, Graph JSON writers, CLI/public output;
- `scripts/build-a00x-benchmark.ps1`;
- D002-D017 implementation ownership;
- every A001/A002/A003 source/test/report/measurement gate.

If another production file is required, A004 must STOP and return to Owner for a fresh architecture decision.

## 9. Expected observable gain

| Target | Expected D001 effect | Expected parent effect | Expected process effect |
|---|---|---|---|
| Cheapapp | material reduction in repeated order-key JSON decode and double-sort work beneath `resolve_calls`; D001 should be lower if the retained profile signal survives wall measurement | the D001 saving should flow into a lower `resolution` parent because the work is synchronous inside that phase | only the retained parent saving may lower process wall; no unrelated process saving is attributed to A004 |
| Restaurant Manager | smaller reduction in the same ordering mechanism; D001 should be lower if the smaller sampled signal survives wall measurement | a smaller corresponding reduction should flow into `resolution` | only that retained parent reduction may lower process wall; no cross-target subsidy or tolerance applies |

Both targets may also show lower transient CPU/allocation work in export-binding evidence ordering, but those secondary metrics cannot substitute for elapsed proof. No exact saving, percentage, minimum threshold, tolerance, or cross-target aggregate is claimed. A004 is retainable only from actual separate D001/parent/process measurements and a later independent Supervisor PASS. The A003-specific Owner KEEP exception creates no A004 precedent.

## 10. Correctness, output, determinism, and failure invariants

Preserve exactly:

- byte-for-byte `Evidence.Kind`, `Weight`, and JSON `Note`;
- exact `[]graph.Evidence` length, contents, and order;
- generic evidence first-seen stable order;
- exact-tuple deduplication by `(Kind, Weight, Note)`;
- export evidence kind rank: terminal, hop, failure, then default;
- proof ordinal, hop ordinal, and lexical Note tie-breaking;
- successful terminal/failure ordering with hop ordinal `-1`;
- failed/malformed Note decode fallback to maximum ordinals plus lexical Note ordering;
- stable order for equal keys;
- distinct SourceSiteID evidence, including coalesced per-site failures;
- idempotent merge behavior;
- no proof, hop, failure, generic record, or source site added, lost, or collapsed;
- existing marshal-failure behavior: an unencodable typed note produces no evidence rather than partial output;
- existing parse-failure behavior: malformed existing evidence remains present and is fallback-ordered rather than dropped or repaired;
- input slices are not unexpectedly mutated or aliased differently;
- relationship identity/coalescing, reference indexes, resolution outcomes, diagnostic carriage, graph nodes/relationships/properties, and execution order;
- canonical in-memory graph, Graph JSON, Ladybug/native persistence/readback, and public stdout/stderr;
- deterministic replay, source freshness/invalidation, failure propagation, transaction rollback, temporary-file flush/close/rename, and publication visibility.

No output/schema/format change is authorized.

## 11. Resource and lifecycle boundary

Let `P` be the number of unique projected export-binding evidence records at one merge.

- transient decorated ordering state must be `O(P)`;
- it may copy `graph.Evidence` values and string headers, but must not duplicate retained Note byte storage;
- order extraction occurs after exact-tuple deduplication and is at most once per unique projected record per merge boundary;
- no retained state survives the merge call;
- no run-wide/global cache, map keyed by Note, goroutine, lock, I/O, serialization side channel, explicit flush, or finalizer;
- current caller-owned input slices and returned graph ownership remain unchanged;
- transient allocation/resource fields must be observed later, not assumed improved.

## 12. Production-first validation contract for a later Owner-released plan

This section is Owner-approved authority for Planner translation only; it does not release Coder.

1. After direct Owner agreement and explicit Planner release, Planner must translate only this direction.
2. A future Coder repeats the current Anvien pre-edit sequence and records file-detail plus exact impacts for every edited helper immediately before editing.
3. Production is changed first inside the one allowed file: remove the projection pre-sort; deduplicate/partition; decorate each unique projected record through the canonical extractor once; stable-sort the decorated slice using cached keys only; undecorate to unchanged `[]graph.Evidence`.
4. Only after production is correct may A004-specific tests be appended to `export_binding_proof_test.go`.
5. The test-local oracle must preserve the pre-A004 algorithm and compare full `[]graph.Evidence` equality. It is never a production fallback.
6. Required adversarial matrix:
   - terminal/hop/failure combinations across multiple proofs and hops;
   - mixed generic/projected evidence;
   - input permutations and repeated coalescing;
   - exact duplicates and same ordinals with different Notes;
   - multiple SourceSiteIDs;
   - malformed JSON for each export evidence kind;
   - unknown kinds and non-export evidence;
   - equal sort keys/stability;
   - empty inputs and no-proof path;
   - caller-input non-mutation and idempotency.
7. Perform the exact build-holder/lock preflight. The canonical full build must PASS before any test execution.
8. After the build, run the new differential test, all existing export-binding proof tests, focused call/access/outcome/Graph JSON regressions, then full `internal/resolution`. The known preserve-only golden must be reported truthfully and may not be edited or called PASS.
9. Build one identifiable A004 candidate using the unchanged accepted A00x overlay/native/runtime/provenance contract; no build-script audit or edit.
10. Remeasure against each target's accepted A003 basis independently:
    - same commands/options/exclusion and target-local graph;
    - D001, parent, analyzer, and process;
    - all `30/30` operations and `17/17` children;
    - exact denominators, interval conservation/zero overlap, workload, graph/DB, semantic counters, evidence ordering, Graph JSON, stdout/stderr, and resource fields.
11. Only after both packets exist may a fresh visible A004 Supervisor review the exact candidate. Main/Owner controls disposition.

Build/test/profile time cannot substitute for child, parent, or process elapsed proof.

## 13. Rollback and mandatory STOP

Exact A004 rollback owns only:

- the A004 ordering/merge hunk and private transient helper/type in `internal/resolution/export_binding_proof.go`; and
- A004-appended bytes in `internal/resolution/export_binding_proof_test.go`.

Rollback the exact A004 bytes if:

- full-array or serialized evidence parity changes;
- generic/export order, exact dedupe, malformed-note fallback, stability, idempotency, or input ownership changes;
- any graph/output/persistence/reader/determinism/failure/lifecycle invariant changes;
- a new validation failure appears;
- either target loses workload/denominator/equivalence identity;
- either target fails to retain lower D001, parent, and process elapsed time under the future disposition rule; or
- the later Supervisor rejects.

Mandatory STOP and return to Owner if:

- implementation requires another production file, public `graph.Evidence` field, shared contract, retained cache, concurrency, persistence/reader change, instrumentation change, or A003 decoder/appender change;
- comparator JSON decoding, a producer-side alternate key authority, or more than one final projected-evidence sort remains;
- malformed evidence would be dropped, repaired, or reclassified;
- D002-D017 is opened as implementation scope;
- exact parity cannot be proven; or
- independent two-target evidence is unavailable.

## 14. Owner decisions — resolved

Owner explicitly resolves the prior discussion points as follows:

1. **Shared-helper boundary approved:** `resolveAccess` and final outcome projection remain preserve/run-only regression surfaces; no edits to those consumers are authorized.
2. **Parity strength approved:** require exact ordered `[]graph.Evidence` equality, including byte-for-byte `Kind`, `Weight`, and `Note`, not semantic-set equality.
3. **Single key authority approved:** extract the order key after dedupe from the final serialized `Evidence.Note` exactly once per unique projected record. Producer-carried typed alternate keys remain forbidden.
4. **Transient memory approved:** one private `O(P)` decorated slice is allowed, with no retained Note-byte copy, public field, or cross-call state.
5. **Target separation approved:** Cheapapp and Restaurant Manager remain independent; no tolerance or aggregation is invented. A004 retention still requires the future governed child/parent/process result on each target plus Supervisor acceptance.
6. **Planner released:** Main may now open the visible Planner lane to translate this exact report into the living A004 plan. Coder remains locked.

## 15. Output and commit boundary

- Updated output: `reports/system-architect/rp_system-architect_260826_by_gpt-5_child06a_a004_residual_direction.md` only.
- Source, test, plan, ledger, benchmark, evidence, actual-status, script, target, and other report edits: none.
- Stage/commit: none, by direct Owner authority.
- Next owner: Main Orchestration, solely to verify this finalized report and open the visible Planner translation lane. Main must keep Coder locked until the translated A004 plan is complete and separately released.

`ARCHITECT_A004_OWNER_APPROVED_READY_FOR_PLANNER`
