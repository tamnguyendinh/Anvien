# Planner Report — Child 06A A004 Export-Binding Evidence Ordering

Date: `2026-08-26`
Role: visible Planner for Child 06A A004
Scope: `P2-A / A004 / B1-P1A-OP001 resolution / B2-P2A-A001-D001 resolve_calls`
Verdict: `PLANNER_A004_READY_FOR_MAIN_VERIFY`
State transition: `A004_ARCHITECT_OWNER_APPROVED -> A004_PLAN_READY_FOR_MAIN_VERIFY / CODER_LOCKED`
Decision boundary: exact architecture-to-plan translation only; no architecture ownership, source/test/script/target edit, build, test, benchmark capture, profile, Supervisor, measurement disposition, detect, stage, commit, P3, Child 07, or Coder authorization
Repository basis: HEAD `54b2bd9b8b10b84861abd81f1686874cbd23c16c`
Accepted performance checkpoint: A003 `b6bf45bce95323aa6b53b182edfea8628bd8b463`
Accepted storage checkpoint: `0f3a572331dd23d17688886fcbfebeb7d37ee35d`

## 1. Authority consumed

The Planner read the full binding sources to EOF in the required order: `AGENTS.md`, `working-rules`, `planner`, `plan-rules.md`, all four Child 06A ledgers, and the finalized Architect report:

`E:\Anvien\reports\system-architect\rp_system-architect_260826_by_gpt-5_child06a_a004_residual_direction.md`

The report verdict is `ARCHITECT_A004_OWNER_APPROVED_READY_FOR_PLANNER`. It releases only exact Planner translation. No A001, A002, A003, or WAL gate was audited, rerun, or reinterpreted.

Planner workflow evidence:

- `anvien --help`: PASS;
- exactly one fresh `anvien analyze --force`: exit `0`, `2241` scanned / `766` parsed / `0` failed, graph `124071 / 170937`;
- indexed/current commit `54b2bd9b8b10b84861abd81f1686874cbd23c16c`, stale `false`;
- file-detail for all four ledgers: PASS; every ledger is current/non-stale LOW-risk documentation with `2` inbound references, `1` outbound reference, `0` unresolved sites, and `2` related files (`plan-rules.md` and the roadmap).

This evidence proves plan freshness and relationships only. It is not product, source-impact, validation, or acceptance evidence.

## 2. Current campaign state preserved

- P2-A, parent `B1-P1A-OP001`, and child `B2-P2A-A001-D001` remain active and unchecked.
- D001 unsuccessful-attempt streak remains `0`.
- D002-D017 remain queued, unchecked, and unopened.
- Cheapapp and Restaurant Manager accepted A003 bases and denominators remain separate in `benchmark.md`; they are not copied into a new numeric control surface, averaged, or combined.
- No A004 candidate, measurement, allocation result, speed claim, Supervisor result, disposition, baseline promotion, or rollback exists.
- Coder is locked. Planner does not open or command Coder.

## 3. Locked A004 architecture translated

The living plan now binds this exact pipeline without reinterpretation:

```text
exact-tuple dedupe
-> derive the canonical order key from final serialized Evidence.Note exactly once per unique projected export-binding evidence record
-> exactly one final stable sort using cached keys
-> discard transient keys
-> return byte/order-identical []graph.Evidence
```

The current cause and ownership are preserved:

- `appendExportBindingEvidence` serializes notes and performs a redundant projection pre-sort;
- `mergeExportBindingEvidence` exact-tuple deduplicates/partitions and owns the final projected ordering;
- the current stable-sort comparator repeatedly calls `exportBindingEvidenceOrderFor`, which decodes final serialized JSON `Evidence.Note`;
- the primary path is `analyze.Run -> runPhase(PhaseResolution) -> ResolveBoundInto -> w.files/ir.Calls -> resolveCall -> appendExportBindingEvidence -> mergeExportBindingEvidence -> sortExportBindingEvidence -> exportBindingEvidenceOrderFor -> json.Unmarshal`;
- resolved relationship coalescing, final outcome/reference-index projection, and sibling `resolveAccess` are preserve/run-only consumers, not additional edit owners.

## 4. Exact future implementation boundary

Allowed production surface, only after separate Main release:

- `internal/resolution/export_binding_proof.go`;
- `appendExportBindingEvidence`;
- `mergeExportBindingEvidence`;
- `sortExportBindingEvidence`;
- `exportBindingEvidenceOrderFor`;
- existing private order/key types; and
- at most one private transient decorated representation/helper.

All existing signatures remain unchanged. Production order is locked:

1. Remove the projection pre-sort.
2. Exact-tuple deduplicate and partition generic/projected evidence first, preserving first-seen behavior.
3. For each remaining unique projected record, call the existing canonical extractor exactly once against final serialized `Evidence.Note` and cache the current `kindRank / proofOrdinal / hopOrdinal / note` tuple.
4. Perform exactly one final stable sort using cached scalar/string keys only.
5. Undecorate, discard transient keys, and return unchanged `[]graph.Evidence`.

Forbidden: producer-carried typed alternate keys or private overloads, comparator JSON decoding/marshaling/validation/mutation, a second final projected-evidence sort, public/persisted order fields, retained cross-call/run-wide/global caches, concurrency/I/O/lifecycle ownership, or any other production file.

## 5. Future Coder pre-edit gate

Before editing, the future separately released Coder must run and record:

1. `anvien --help`;
2. one fresh `anvien analyze --force`;
3. file-detail for `internal/resolution/export_binding_proof.go`; and
4. exact upstream impact for every helper it will edit.

Current architecture evidence classifies the containing file HIGH and every named helper CRITICAL. Exact current warnings are:

- `appendExportBindingEvidence`: `36` impacted symbols / `5` direct callers / `14` files / `7` modules / `39` processes;
- `mergeExportBindingEvidence`: `51 / 5 / 19 / 7 / 49`;
- `sortExportBindingEvidence`: `23 / 2 / 10 / 2 / 34`;
- `exportBindingEvidenceOrderFor`: `14 / 1 / 6 / 1 / 19`.

HIGH/CRITICAL are scope warnings, not edit bans. Failure to prove current graph/source identity or the exact one-file owner boundary is a STOP before edit.

## 6. Tests after production

Only after production is correct may A004-specific bytes be appended to:

`internal/resolution/export_binding_proof_test.go`

The tests use a test-local pre-A004 oracle, never a production fallback, and compare complete ordered `[]graph.Evidence` equality. The required matrix covers:

- terminal/hop/failure combinations across multiple proofs and hops;
- mixed generic/projected evidence;
- input permutations and repeated coalescing;
- exact duplicates and equal ordinals with different Notes;
- multiple SourceSiteIDs;
- malformed JSON for every export evidence kind;
- unknown kinds and non-export evidence;
- equal sort keys and stable order;
- empty inputs and the no-proof path;
- caller-input non-mutation; and
- idempotency.

## 7. Build and validation order

1. Inspect build holders/processes and terminate every proven build-output holder completely.
2. Run canonical `pwsh -NoLogo -NoProfile -NonInteractive -ExecutionPolicy Bypass -File E:\Anvien\scripts\full-build.ps1` and require PASS.
3. Only after that PASS, run the new differential test and all existing export-binding proof tests.
4. Run focused call, access, outcome, and Graph JSON regressions that consume export-binding evidence ordering.
5. Run full `go test ./internal/resolution -count=1` truthfully.
6. Run existing A003 graphhealth parity checks as needed to prove no diagnostic/output drift.

The known preserve-only `TestProofBasedCallAccessGoldenCorpus` failure may be recorded only if unchanged. The package is never called PASS, its owner is never edited, and any new or changed failure blocks A004.

## 8. Independent two-target measurement

Reuse unchanged `scripts/build-a00x-benchmark.ps1` and the accepted overlay/native/runtime/provenance contract. Do not audit or edit the script. Build one identifiable A004 candidate from the accepted A003 state plus only A004-owned production bytes.

Measure Cheapapp and Restaurant Manager independently against their own accepted A003 bases. Preserve each target's accepted command/options/exclusion and target-local graph. Each packet must record:

- D001, parent, analyzer, and process elapsed wall time;
- all `30/30` operations and `17/17` children;
- exact denominators;
- interval conservation and zero overlap;
- workload, graph/DB, semantic counters, diagnostics/outcomes;
- complete export-binding evidence content and order;
- canonical Graph JSON;
- public stdout/stderr; and
- resource fields, including the private transient `O(P)` lifecycle boundary.

Targets are never averaged or combined. Build/test/profile time and secondary resource metrics cannot substitute for D001, parent, or process elapsed proof. The A003-specific Owner exception creates no A004 tolerance or precedent.

## 9. Supervisor and disposition

Only after both complete measurement packets exist may one fresh visible A004 Supervisor review the exact candidate and the affected evidence-ordering, graph/output, persistence/reader, determinism/failure/lifecycle, and resource boundary.

Main/Owner alone controls `KEEP`, rollback, or another governed disposition. Planner makes no performance claim and changes no benchmark value, baseline, streak, checkbox, or queue.

## 10. Exact invariants and resource boundary

Preserve exactly:

- byte-for-byte `Evidence.Kind`, `Weight`, and JSON `Note`;
- full evidence length, contents, and order;
- generic first-seen stable order;
- exact-tuple deduplication by `(Kind, Weight, Note)`;
- terminal, hop, failure, then default kind rank;
- proof ordinal, hop ordinal, and lexical Note tie-breaking;
- successful terminal/failure hop ordinal `-1`;
- malformed-note fallback to maximum ordinals plus lexical Note order;
- stability for equal keys;
- distinct SourceSiteIDs and coalesced per-site failures;
- merge idempotency;
- marshal-failure and parse-failure behavior;
- caller-input non-mutation and existing aliasing/ownership;
- relationship/reference/outcome/diagnostic carriage;
- graph nodes/relationships/properties and execution order;
- in-memory graph, Graph JSON, Ladybug/native persistence/readback, and stdout/stderr; and
- deterministic replay, freshness/invalidation, failure propagation, transaction rollback, temp flush/close/rename, and publication visibility.

For `P` unique projected records at one merge, transient decoration is `O(P)`, may copy evidence values and string headers, must not duplicate retained Note byte storage, and must become unreachable when the merge returns. No retained cache, goroutine, lock, I/O, side channel, flush, or finalizer is authorized.

## 11. Rollback and STOP

Exact rollback owns only:

- the A004 ordering/merge hunk and private transient helper/type in `internal/resolution/export_binding_proof.go`; and
- A004-appended bytes in `internal/resolution/export_binding_proof_test.go`.

Rollback exact A004 bytes on any full-array/serialized-evidence parity change, ordering/dedupe/fallback/stability/idempotency/non-mutation regression, graph/output/persistence/reader/determinism/failure/lifecycle regression, new validation failure, target identity mismatch, loss of governed D001/parent/process retention, or Supervisor rejection.

Mandatory STOP and return to Owner on another production file, public/shared contract, changed signature, public `graph.Evidence` field, alternate producer key authority, retained cache/concurrency, persistence/reader/instrumentation/A003 change, comparator JSON decoding, more than one final projected-evidence sort, malformed evidence drop/repair/reclassification, D002-D017 implementation scope, unprovable exact parity, or unavailable independent two-target evidence.

## 12. Owned output and next owner

Planner changed only the four existing Child 06A ledgers and created this report. `plan-rules.md`, Architect report, source, tests, scripts, benchmark artifacts, target repositories, and all other reports remain outside Planner write ownership. Planner did not stage or commit.

Next owner: Main Orchestration for independent verification of the exact five owned artifacts and a separate future Coder-release decision.

`PLANNER_A004_READY_FOR_MAIN_VERIFY`
