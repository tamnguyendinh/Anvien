# System Architect Report — Child 06A A003 Residual Direction

Date: `2026-08-26 14:12:00 +07:00`
Role: visible System Architect
Scope: `P2-A / A003 / B1-P1A-OP001 resolution / B2-P2A-A001-D001 resolve_calls`
Decision boundary: architecture direction for Planner translation only; no implementation, Coder authorization, `KEEP`, detect, stage, commit, P3, or Child 07 authority
Graph-evidence HEAD: `273788864301fd063c64c9b0e46b43dd117f3fee`; the later docs-only `f072a8f9` plan-rules wording checkpoint changes no production/test bytes and does not invalidate owner/impact evidence
Current graph evidence: refreshed once only after Anvien explicitly reported the earlier index stale; `2230` scanned / `767` parsed / `0` failed, graph `123929 / 170690`, non-stale at the graph-evidence HEAD

## 1. Authority and current control rows

The prior attribution outputs and current source/profile packet already exist as the current campaign input. The Owner did not personally validate the technical attribution result. The Owner explicitly requires separating the earlier process/ownership violation from the existing result and forbids rerunning already available attribution, pprof, or source-discovery work. This decision consumes `E2-P2A-A003CURRENT1/ATTRIB1` as supplied campaign input without reopening it.

| Control | Current row/state |
|---|---|
| Campaign state | `A001 KEEP / A001_COMMIT_COMPLETE / A002 KEEP / A003 RESIDUAL_ATTRIBUTION_COMPLETE / ARCHITECT_ACTIVE` |
| Implementation slice | `P2-A`, unchecked; the only implementation slice |
| Active parent | `B1-P1A-OP001 resolution`, unchecked |
| Active child | `B2-P2A-A001-D001 resolve_calls`, unchecked |
| D001 consecutive no-KEEP streak | `0` |
| Remaining child queue | `D002-D017`, all unchecked, queued, and unopened |
| A002 accepted production delta | uncommitted and preserved: `internal/graphhealth/diagnostics.go +52/-2`, `internal/resolution/emit.go +3/-1`, `internal/resolution/outcome.go +1/-1` |
| A002 accepted test state | new `internal/graphhealth/diagnostics_test.go`, preserved; staged set empty |
| Commit boundary | no A003 detect/stage/commit; P3-C alone owns final detect/commit |

Accepted A002 baselines remain independent and must never be averaged or combined:

| Target | D001 `resolve_calls` | Parent `resolution` | Analyzer / process | Denominator |
|---|---:|---:|---:|---|
| Cheapapp | `3.090914200 s` | `19.040468000 s` | `100.843249000 / 136.729876000 s` | `calls=27890; files=887` |
| Restaurant Manager | `9.909636600 s` | `21.242055400 s` | `109.339859600 / 145.066210900 s` | `calls=86030; files=1234` |

Complete active-parent child queue, with no transition or implicit terminalization:

1. `[ ] B2-P2A-A001-D001 resolve_calls` — active.
2. `[ ] B2-P2A-A001-D002 resolve_accesses` — queued/unopened.
3. `[ ] B2-P2A-A001-D003 resolve_type_annotations` — queued/unopened.
4. `[ ] B2-P2A-A001-D004 project_resolution_outcomes` — queued/unopened.
5. `[ ] B2-P2A-A001-D005 emit_definition_nodes` — queued/unopened.
6. `[ ] B2-P2A-A001-D006 finalize_resolution_outcomes` — queued/unopened.
7. `[ ] B2-P2A-A001-D007 build_binding_occurrence_index` — queued/unopened.
8. `[ ] B2-P2A-A001-D008 finalize_typescript_authority_results` — queued/unopened.
9. `[ ] B2-P2A-A001-D009 emit_import_edges` — queued/unopened.
10. `[ ] B2-P2A-A001-D010 emit_typescript_external_symbols` — queued/unopened.
11. `[ ] B2-P2A-A001-D011 emit_method_dispatch_edges` — queued/unopened.
12. `[ ] B2-P2A-A001-D012 assemble_resolution_result` — queued/unopened.
13. `[ ] B2-P2A-A001-D013 binding_accumulator_dispose` — queued/unopened.
14. `[ ] B2-P2A-A001-D014 emit_heritage_compatibility_edges` — queued/unopened.
15. `[ ] B2-P2A-A001-D015 emit_unresolved_heritage_diagnostics` — queued/unopened.
16. `[ ] B2-P2A-A001-D016 finalize_resolution_metadata` — queued/unopened.
17. `[ ] B2-P2A-A001-D017 runtime_setup` — queued/unopened.

## 2. Retained cause and profile caveats

The existing two profiles retain the same stack under `resolveCall`, independently:

| Target | `resolveCall` | `diagnosticAppender.appendToNode` | normalization/policy | decoder | `encoding/json.Unmarshal` |
|---|---:|---:|---:|---:|---:|
| Cheapapp | `2.76 s` | `0.92 s` | `0.87 s` cumulative | included in the same `0.87 s` cumulative stack | `1.84 s` cumulative |
| Restaurant Manager | `9.19 s` | `2.35 s` | `2.11 s` cumulative | `2.10 s` cumulative | `2.28 s` cumulative |

These are cumulative CPU samples only. They overlap, are non-additive, are not elapsed wall time, are not averaged across targets, and do not predict a speedup. Not all `encoding/json.Unmarshal` samples are attributed exclusively to this decoder.

Retained source cause:

- Per diagnostic, `diagnosticAppender.appendToNode:60-88` calls normalization at line `77`.
- The chain is `normalizeDiagnosticMetadata:274-287 -> structuredResolutionDiagnosticPolicy:353-375 -> decodeStructuredResolutionOutcome:377-400`.
- `decodeStructuredResolutionOutcome` currently decodes the same trimmed `Diagnostic.Note` into an envelope map at line `385` and again into `structuredResolutionOutcome` at line `396` whenever the note is treated as structured.
- `statusMarker := isStructuredResolutionStatus(...)` is already computed before either decode. When it is true, structured status is already proven; the envelope map contributes no additional classification decision for that branch.
- The conditional authority decode at `validStructuredResolutionAuthority:460-491`, line `465`, is separate and must not be conflated with or changed by this direction.

Complete retained upstream path:

`internal/analyze.Run:365-370 -> runPhase:1134-1155 -> resolution.ResolveBoundInto:57-128 -> w.files / ir.Calls:91-94 -> resolveCall:385-605`.

Repository-unresolved family:

`resolveCall:388/392/560/564 -> emitter.emitUnresolvedReference:182-225 -> added guard -> e.diagnosticAppender:220`.

TypeScript family:

`resolveCall:516-535 -> recordTypeScriptLookup:926-977 -> recordTypeScriptOutcome -> added && non-resolved-external -> emitTypeScriptOutcomeDiagnostic:372-397 -> e.diagnosticAppender:392`.

Dynamic binding remains:

`newEmitter:31-42 -> graphhealth.NewDiagnosticAppender:52-58 -> diagnosticAppender.appendToNode`.

## 3. Exactly one selected technical direction

### Canonical single-interpretation diagnostic decoder

The architecture rule is:

`one Diagnostic -> one canonical full-Note interpretation -> one semantic decode result -> one policy decision -> one graph write`.

`decodeStructuredResolutionOutcome` becomes the sole authority for interpreting the outer `Diagnostic.Note`. It must traverse/decode the outer note exactly once and produce one internal semantic result containing all three facts needed downstream:

- whether structured evidence exists, derived deterministically from the existing `SourceSiteStatus` evidence plus exact top-level `schemaVersion`/`status` presence observed during that same interpretation;
- the typed `structuredResolutionOutcome` when decodable; and
- validity/error state sufficient to preserve fail-closed behavior.

The semantic result has exactly three states:

1. `UNSTRUCTURED`: no structured evidence exists.
2. `STRUCTURED_VALID`: structured evidence exists and the complete outcome contract validates.
3. `STRUCTURED_INVALID`: structured evidence exists but syntax, type, required-field, proof, authority, or contract validation fails.

These are result states, not competing execution paths. There is no primary decoder, recovery decoder, legacy decoder, retry decoder, or fallback decoder. `SourceSiteStatus` is evidence consumed by the same decision; it never selects another decoding route.

If any structured evidence exists, every decode or validation error must produce `STRUCTURED_INVALID` and retain current `unclassified/review` fail-closed behavior. The decoder must never reinterpret an invalid structured payload as unstructured merely to continue graph generation.

The outer note is interpreted once. Exact marker presence and typed outcome are products of that same interpretation, so the current envelope-map interpretation and second full-note typed interpretation cease to exist as separate authorities. The conditional decode of the explicitly nested `Authority` subdocument remains unchanged: it validates a nested serialized contract and is not a second interpretation of the outer note.

Implementation freedom is limited to the private representation necessary to produce this one semantic result inside `internal/graphhealth/diagnostics.go`. A private local/wire result is permitted only as state owned by the canonical decoder. It must not become a second decoder, shared contract, cache, public type, or alternate path. The existing external function signature and downstream policy contract remain unchanged.

This updated direction supersedes and withdraws the earlier status-marker fast-path/fallback wording. Planner must translate only this canonical single-authority architecture.

## 4. Exact owners and Anvien blast radius

### Production owner

- File: `internal/graphhealth/diagnostics.go` only.
- Existing semantic owner: `decodeStructuredResolutionOutcome`, lines `377-400`.
- Allowed implementation surface: refactor this decoder and, only if necessary, add one private non-exported decode-result/wire representation in the same file that is exclusively owned by it.
- No other existing production function, method, struct contract, constant, exported surface, or file may change.

Fresh Anvien file-detail at HEAD `273788864301fd063c64c9b0e46b43dd117f3fee`:

- graph non-stale; `changedSinceAnalyze=false`;
- file risk `HIGH`;
- `164` symbols, `51` inbound refs, `81` outbound refs, `161` local relationships, `2` linked flows, `20` linked tests;
- functional area `graph_health`, backend layer.

Exact upstream symbol impact for `decodeStructuredResolutionOutcome` with tests included:

- symbol risk `LOW`, but the containing file remains `HIGH`;
- `12` impacted symbols, `1` direct caller, `1` affected module, `0` affected processes;
- affected layers: backend `7`, backend-test `5`;
- affected files: `internal/graphhealth/diagnostics.go` (`7` symbols), `internal/graphhealth/compute_test.go` (`1`), and `internal/graphhealth/p6d_outcome_projection_test.go` (`4`).

The file-level HIGH blast radius is a scope warning, not an edit ban. It requires the single-owner boundary and full graph-health regression checks below. The LOW exact-symbol result does not authorize broad edits elsewhere in the HIGH file.

### Authorized test owner

- `internal/graphhealth/diagnostics_test.go` only, and only after the production body is correct.
- Add A003-owned differential cases; preserve every accepted A002 test byte and do not replace or delete the file.
- `internal/graphhealth/compute_test.go` and `internal/graphhealth/p6d_outcome_projection_test.go` are run-only regression surfaces, not edit owners.

## 5. Expected observable gain

- D001 boundary: lower retained `resolve_calls` elapsed time on Cheapapp and Restaurant Manager independently.
- Parent boundary: lower retained `resolution` elapsed time on each target independently.
- Process boundary: lower retained process wall time on each target independently.

No percentage, threshold, absolute saving, or cross-target aggregate is predicted. A candidate is retainable only if all three elapsed boundaries improve on both targets independently and the later per-attempt Supervisor passes.

## 6. Preserved invariants and lifecycle/resource boundary

Preserve exactly:

- diagnostic structured/legacy classification and actionability, including fail-closed invalid structured notes;
- the full `(outcome, structured, valid)` return tuple for every current input class;
- exact, case-sensitive top-level `schemaVersion`/`status` presence semantics and existing `SourceSiteStatus` semantics, now resolved by one authority;
- `UNSTRUCTURED`, `STRUCTURED_VALID`, and `STRUCTURED_INVALID` must project to the exact same externally observed tuple and policy results as the accepted pre-A003 behavior;
- bucket identity, count merge, target fill, earliest-line behavior, stable order, defaults, and per-append write-through visibility;
- `[]Diagnostic` graph representation, Graph JSON, persistence/readback, health projections, public output, and determinism;
- every resolution outcome, encoded `Note`, authority/proof/source-site field, metric, node, relationship, ID, label, property, branch, and order;
- the conditional authority decode and validation/failure paths;
- existing sequential per-`ResolveBoundInto` lifecycle and the accepted A002 `O(T)` shared-slice/no-duplicate-diagnostic-object topology.

Resource boundary:

- only per-call local decode state owned by the canonical decoder;
- no retained cache, map, diagnostic copy, goroutine, lock, I/O, serialization change, global state, flush, or finalizer;
- the A002 appender map/slice ownership and lifetime remain unchanged.

## 7. Exact production-first, build, test, and measurement sequence

1. Planner translates only this direction into A003 `PLAN1`. Coder remains unauthorized until that translation and Main verification complete.
2. Immediately before editing, the future Coder confirms a current graph and records fresh file-detail for `internal/graphhealth/diagnostics.go` plus exact impact for `decodeStructuredResolutionOutcome` if source/graph identity has changed. This report does not waive the pre-edit gate.
3. Edit production first: replace the two independent full-note interpretations with one canonical presence-aware interpretation that produces structure state, typed outcome, and validity together. Do not preserve or introduce a second production decoding route.
4. After production behavior is correct, append A003-only cases to `internal/graphhealth/diagnostics_test.go`. Use a test-local pre-A003 reference oracle to compare the complete externally observed tuple across at least: structured evidence from status, structured evidence from exact note markers, both evidence sources together, valid markerless structured notes, unstructured notes, malformed JSON, non-object JSON, null, typed-field errors, missing/exact-case/case-variant markers, duplicate keys, unknown fields, conflicting evidence, and invalid authority/proof data. The reference exists only in tests and is never a production execution path. Do not add runtime instrumentation or edit an existing golden.
5. Perform exact build-holder/lock preflight and terminate proven build-output holders completely. Then run the canonical full build before any test command:

   `pwsh -NoLogo -NoProfile -NonInteractive -ExecutionPolicy Bypass -File E:\Anvien\scripts\full-build.ps1`

6. Only after full-build PASS, run the new focused decoder-parity test, `TestDiagnosticAppenderMatchesLegacySemantics`, `TestP6DStructuredOutcomeProjectionPreservesGraphJSONAndHealthParity`, the existing resolution-focused tests from A002, `go test ./internal/graphhealth -count=1`, and `go test ./internal/resolution -count=1`.
7. `TestProofBasedCallAccessGoldenCorpus` remains a truthful pre-existing preserve-only FAIL. It is never a package PASS, never A003 validation evidence, and never editable by this direction. Any new or changed failure blocks the candidate.
8. Build one identifiable A003 measurement candidate with the already accepted, unchanged `scripts/build-a00x-benchmark.ps1` contract: accepted A002 plus only the A003 production delta, retaining the same 17-child overlay/native/runtime/build identity. No script edit or renewed build-interface audit is authorized.
9. Measure Cheapapp against its own accepted A002 baseline with the same target/options/work denominator; record D001, parent, process, all 30 operations, all 17 children, output/equivalence/resource identity, and graph/readback semantics.
10. Separately measure Restaurant Manager against its own accepted A002 baseline with the same target/options, exactly one existing `userApi.ts` exclusion, and its own denominator; record the same boundaries. Never average or combine the targets.
11. A fresh visible per-attempt Supervisor reviews the exact A003 candidate. Main alone later owns disposition. No detect, stage, commit, P3, or Child 07 action occurs in A003 execution.

## 8. Attempt-owned rollback and mandatory STOP

Exact A003 rollback owns only:

- the canonical decoder hunk and any exclusively owned private decode-result/wire representation added in `internal/graphhealth/diagnostics.go`; and
- the A003-specific test additions appended to `internal/graphhealth/diagnostics_test.go`.

Rollback must preserve the entire accepted A002 production delta, the pre-existing A002 test content, A001/P1 work, ledgers, reports, scripts, and unrelated/user/protected changes. Do not delete the whole A002-created test file and do not use broad reset, checkout, stash, or cleanup.

Rollback the exact A003 bytes if any build, tuple-parity, fail-closed, classification/actionability, graph/output, persistence/readback, determinism, lifecycle/resource, D001, parent, process, or later Supervisor gate fails. A non-retainable timing result remains an unsuccessful A003 attempt under plan rules; it does not authorize a second direction inside A003.

Mandatory STOP and return to Main for a new architecture attempt if any of the following is required:

- a second production direction;
- any production file or existing semantic owner beyond `diagnostics.go::decodeStructuredResolutionOutcome`;
- a second decoder, recovery route, legacy route, retry route, fallback route, retained cache, shared/public decode contract, or edits to policy, authority decode, appender, resolution emitters, schemas, persistence/readers, public contracts, timing instrumentation, or D002-D017;
- broadening the workload/output/failure contract or weakening exact single-authority/fail-closed semantics;
- inability to prove parity within the one-canonical-interpretation boundary;
- a new validation failure, denominator/workload mismatch, or unavailable independent two-target evidence.

## 9. NO EVIDENCE

- `NO EVIDENCE`: exact invocation count of `decodeStructuredResolutionOutcome` or exact count of diagnostics where `statusMarker == true`.
- `NO EVIDENCE`: exact split between repository-unresolved and TypeScript diagnostic path families.
- `NO EVIDENCE`: elapsed-time attribution of the decoder or envelope decode inside the retained D001 residual.
- `NO EVIDENCE`: that every `encoding/json.Unmarshal` CPU sample belongs to this decoder.
- `NO EVIDENCE`: numeric D001, parent, or process gain from this direction.
- `NO EVIDENCE`: allocation or memory improvement; only the no-new-retained-state boundary is authorized.
- `NO EVIDENCE`: the exact implementation technique required to realize the canonical single interpretation; Planner/Coder may choose only an in-owner technique that satisfies the architecture contract.

## 10. Handoff

Durable output is this updated report only. No source, test, ledger, plan-rules, staged set, or commit was changed by this lane. The earlier fast-path/fallback direction is withdrawn. Next owner is Main Orchestration for report verification, followed by Planner translation only if this canonical single-authority READY verdict is accepted.

ARCHITECT_A003_READY_FOR_PLANNER
