# Child 06A A006 Residual Architecture Decision

Verdict: `ARCHITECT_A006_NEEDS_MEASUREMENT_INPUT`

## Scope And Authority

- Role: fresh visible A006 System Architect.
- Slice: `P2-A / A006 / B1-P1A-OP001 resolution / B2-P2A-A001-D001 resolve_calls`.
- Repository: `E:\Anvien`, direct shared checkout only.
- Current clean HEAD: `4b65bda273774b4b840325c14e848d7ddb6b9aef`.
- Accepted implementation and measurement baseline: A003 checkpoint `b6bf45bce95323aa6b53b182edfea8628bd8b463`.
- Entering state: `A004 SUPERVISOR_PASS / NO_KEEP / ROLLBACK_COMPLETE`; `A005 SUPERVISOR_PASS / NO_KEEP / ROLLBACK_COMPLETE`; `A006_ARCHITECT_PENDING / D001_STREAK_2`.
- Parent and D001 remain unchecked. D002-D017 remain queued and unopened. P3 and Child 07 remain closed.
- This lane owns architecture evidence only. It performs no plan translation, source/test/ledger edit, build, test, target analyze, new profile capture, benchmark, Supervisor review, disposition, detect, stage, commit, or terminalization.

This report intentionally selects no production direction. Existing evidence does not prove one materially new, safe D001 mechanism with a plausible retained benefit on both targets. The requested input below is measurement-only, is not A006 production work, and must not increment the D001 unsuccessful-attempt streak.

## Current Accepted Numeric Basis

Targets remain separate and are never averaged or combined.

| Target | D001 `resolve_calls` | Parent `resolution` | Analyzer | Process | Denominator |
|---|---:|---:|---:|---:|---|
| Cheapapp | `3.447846300 s` | `20.472602300 s` | `93.531974900 s` | `95.630648200 s` | `calls=27890; files=887` |
| Restaurant Manager | `9.401585300 s` | `20.850792800 s` | `98.020546700 s` | `101.096911900 s` | `calls=86030; files=1234` |

Initial-to-A003 process reduction remains objective accepted history:

- Cheapapp: `890.314783200 -> 95.630648200 s`.
- Restaurant Manager: `1178.391336900 -> 101.096911900 s`.

No A004 or A005 candidate value is promoted into this basis.

## Authority And Freshness Consumed

The following were read fully through EOF in the mandated order before graph/source conclusions:

1. `AGENTS.md`.
2. `.agents/skills/working-rules/SKILL.md`.
3. `.agents/skills/System-Architect/SKILL.md`.
4. Child 06A `plan-rules.md`.
5. All four standard Child 06A ledgers.
6. A004 Architect, Supervisor, and both target reports.
7. A005 Architect, Supervisor, exact rollback, and both target reports.
8. Both accepted A003 target reports and both retained A003 CPU profiles.

Mandated discovery then completed:

- `anvien --help`: PASS.
- Exactly one fresh `anvien analyze --force`: exit `0`; `2261` scanned, `766` parsed code, `0` failed; graph `124365` nodes / `171231` relationships.
- No index cleanup, move, deletion, or recovery action occurred.

## Verified Current-Source And Graph Facts

### Exact D001 owner

The current measured child owner remains:

- file: `internal/resolution/resolve.go`;
- symbol: `resolveCall`;
- current source range: lines `385-605`;
- caller in the active phase: `ResolveBoundInto` at the `w.files -> ir.Calls` loop, lines `91-94`.

Fresh file-detail for `internal/resolution/resolve.go` reports:

- current and non-stale; `changedSinceAnalyze=false`;
- risk `HIGH`;
- `200` symbols;
- inbound/outbound/local relationships `134 / 289 / 87`;
- `401` unresolved source sites;
- `26` linked flows and `40` linked tests.

Fresh exact upstream impact for `resolveCall`, including tests, reports:

- risk `CRITICAL`;
- `31` impacted symbols;
- `1` direct caller;
- `13` affected files;
- `7` modules;
- `32` processes.

HIGH/CRITICAL are blast-radius warnings. They do not prohibit a future scoped edit, but they prevent authorizing an unmeasured residual hypothesis as A006 production work.

### Current `resolveCall` decision path

For each `scopeir.CallSiteFact`, current source performs these synchronous families, depending on call form and branch result:

1. Resolve the source through `sourceForScopeOrFile` / `callerForScope`.
2. For member calls, test repository receiver ownership and binding-occurrence resolution.
3. Resolve constructor/default/member targets through combinations of:
   - `resolveScopedName`;
   - `resolveSameFileName`;
   - `resolveMember`;
   - `resolveImportedMemberWithProof`;
   - the accepted A001 import-claim helpers;
   - `resolveGoSamePackageFunction` for default Go calls;
   - `resolveGlobalCallName`.
4. When repository resolution is absent or low-confidence, run TypeScript standard-library lookup and `recordTypeScriptLookup`.
5. Emit one of the resolved/unresolved carriers through `appendExportBindingEvidence`, `emitUnresolvedReference`, or `emitReference`, with source-site identity, evidence, diagnostic, outcome, graph, and reference-index effects.

Current source also proves that `resolveGoSamePackageFunction` performs a whole `defsByFile` map traversal for each eligible Go call, filtering every candidate file by language and `path.Dir`, then checking function labels and lookup names. That is a concrete algorithmic mechanism, but the retained profiles prove it is dominant only on Restaurant Manager, not on Cheapapp.

### Complete upstream and downstream call path

The active product path is:

`CLI analyze -> internal/analyze.Run -> runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> w.files -> ir.Calls -> resolveCall`.

Within D001, terminal branches feed:

- resolved repository path: `emitReference -> emitRelationship -> graph.Relationship / ReferenceIndex`, plus repository resolved outcome recording;
- unresolved repository path: `emitUnresolvedReference -> recordRepositoryUnresolvedOutcome -> diagnosticAppender -> Graph.AddNode`;
- TypeScript path: `lookupTypeScriptMember/lookupTypeScriptGlobal -> recordTypeScriptLookup -> recordTypeScriptOutcome -> optional emitTypeScriptOutcomeDiagnostic`;
- export-proof path: `appendExportBindingEvidence -> mergeExportBindingEvidence`, carried into relationships/references.

After all calls/accesses/type annotations:

`finalizeTypeScriptAuthorityResults -> emitTypeScriptExternalSymbols -> resolutionOutcomeCollector.finalize -> projectResolutionOutcomes -> resolution.Result`.

`analyze.Run` then installs `Result.Graph`, `TypeScriptAuthorityResults`, `ResolutionOutcomes`, and resolution metrics. The same graph continues through MRO, communities, processes, semantic enrichment, `Graph.Compact`, Ladybug/native `loadGraph`, canonical Graph JSON `writeGraphSnapshot`, and CLI/public output publication.

Any future A006 production direction must preserve this entire synchronous result and publication tail; reducing D001 by transferring work into final projection, a later phase, DB load, snapshot, AI context, or CLI publication is not a net benefit.

## Retained Profile Inference — Not Elapsed Proof

The retained A003 artifacts are:

- Cheapapp: `E:\Anvien\.tmp\child06a_a003_cheapapp_frozen_17child_20260826\candidate\resolution.cpu.pprof`, `65363` bytes, SHA-256 `806FAD79B12C567618778A8232E8521991882CE6C745AC2EBBE7B09ECC267107`.
- Restaurant Manager: `E:\Anvien\.tmp\child06a_a003_restaurant_manager_frozen_17child_20260826\candidate\resolution.cpu.pprof`, `56397` bytes, SHA-256 `4272FFAB0ABE4C14C470A139A31A601F31323B1373E0FA2986A09935456FA3E2`.

Bounded `go tool pprof` reads focused on stacks containing `resolution.resolveCall`. Values are cumulative CPU samples. They overlap, are non-additive, are not wall time, and cannot be used as a speedup claim.

| Stack under D001 | Cheapapp | Restaurant Manager | Authority effect |
|---|---:|---:|---|
| `resolveCall` cumulative | `3.12 s` | `8.29 s` | D001 CPU sample envelope only |
| diagnostic appender | `1.08 s` | `1.88 s` | accepted A002 owner; cannot repeat |
| canonical diagnostic decoder | `1.02 s` | `1.58 s` | accepted A003 owner; cannot repeat |
| export-evidence sort/extractor | `1.23 / 1.22 s` | `0.21 / 0.20 s` | rejected A004 owner; cannot repeat |
| outcome collector record / marshal | `0.26 / 0.16 s` | `0.63 / 0.49 s` | rejected A005 owner; cannot repeat |
| `resolveGoSamePackageFunction` | no retained sample match | `4.17 s` | dominant unattempted stack on Restaurant only |
| `sourceForScopeOrFile` | `0.07 s` | `0.15 s` | common but sparse/overlapping |
| `resolveSameFileName` | `0.07 s` | `0.18 s` | common but sparse/overlapping |
| `sourceSiteID` | `0.01 s` | `0.21 s` | common but sparse/overlapping |
| `callTargetText` | no retained sample match | `0.04 s` | no two-target material signal |

The unresolved-emission stack further shows why a renamed variant of A002/A003/A005 is not valid A006:

- Cheapapp `emitUnresolvedReference` is `0.49 s`, of which the appender is `0.37 s` and repository outcome recording is `0.12 s`.
- Restaurant Manager `emitUnresolvedReference` is `2.01 s`, of which the appender is `1.50 s`, repository outcome recording is `0.39 s`, and the remaining sampled path/site work is small.

Those cumulative children overlap, but they show that nearly all sampled common unresolved-emission work is already owned by accepted or rejected prior directions.

## Why Existing Evidence Cannot Safely Select A006

### No one materially new common cause is proven

After excluding A001-A005 ownership, the largest clear unattempted stack is `resolveGoSamePackageFunction` on Restaurant Manager. Cheapapp has no retained sample match for it. Authorizing a Go same-package index would therefore predict no evidenced Cheapapp D001 benefit, while A006 retention still requires target-separated child, parent, and process improvement.

The helpers that do appear on both targets (`sourceForScopeOrFile`, `resolveSameFileName`, `sourceSiteID`) have small, overlapping cumulative samples. Current evidence does not expose:

- their exact invocation counts by D001 branch;
- their exclusive wall time;
- scope-chain hops or repeated site/target construction counts;
- how much of their sampled work is already nested in prior A002-A005 stacks;
- whether eliminating any one family is large enough to survive at D001, parent, analyzer, and process boundaries on both targets.

### Local savings have already failed to retain at process wall

A004 and A005 are objective warnings against inferring net process benefit from a locally plausible CPU mechanism:

- A004 lowered D001 and parent on both targets, but process wall increased by `49.345324200 s` on Cheapapp and `34.472577200 s` on Restaurant Manager.
- A005 lowered D001 and parent on both targets, but process wall increased by `107.968183700 s` on Cheapapp and `61.193290300 s` on Restaurant Manager. The corrected runs overlapped as recorded, which limits isolated-machine interpretation without changing the measured facts or NO_KEEP disposition.

A006 cannot responsibly authorize a still smaller common-per-call cleanup without exclusive current attribution. That would be a guess, not an evidence-backed architecture direction.

## Alternatives Considered And Rejected For A006

1. **Go same-package function index.** Rejected now because the retained causal signal is Restaurant-only. It cannot plausibly satisfy the independent Cheapapp child requirement from existing evidence.
2. **Call-scoped target-text/source-site reuse.** Not selected. Current source proves repeated deterministic construction, but retained common samples are too small and no exact dynamic repetition or exclusive wall cost exists on both targets.
3. **Scope/caller or same-file lookup memoization/indexing.** Not selected. It risks becoming a cosmetic extension of A001's run-scoped index pattern without evidence that this distinct lookup family is materially responsible on both targets.
4. **Outcome clone/byte lifecycle variation.** Forbidden as an A005 cosmetic variant. A005 already owned record-time canonical byte lifecycle, strict projection consumption, and private sidecar/finalized representation; it was rolled back after NO_KEEP.
5. **Export-evidence key/order/sort variation.** Forbidden as an A004 cosmetic variant. A004 already owned exact-tuple dedupe, cached canonical order keys, and one final stable sort; it was rolled back after NO_KEEP.
6. **Another diagnostic decoder/appender path.** Forbidden by accepted A002/A003. No fallback, recovery, fast-path, or second interpretation may be introduced.
7. **Parallel call resolution.** Rejected without further authority. Current emission mutates shared graph, reference index, metrics, outcome collector, diagnostic appender, TypeScript result slices, and ordering-sensitive carriers. Existing architecture proves sequential ownership, not safe deterministic concurrency.

No production alternative is released to Planner.

## Smallest Exact Missing Measurement

### Measurement ID and purpose

Request one measurement-only input: `A006-M1-D001-DIRECT-CALLEE-ATTRIBUTION`.

Its sole purpose is to identify whether one unattempted synchronous `resolveCall` family has exclusive wall cost on **both** targets. It is not a candidate, production attempt, benchmark promotion, Supervisor input, or no-KEEP event.

### Exact measurement owner and allowed surface

- Semantic owner under measurement: `internal/resolution/resolve.go::resolveCall` only.
- Allowed instrumentation surface: the measurement-overlay copy of `internal/resolution/resolve.go` at the direct call sites inside `resolveCall`.
- The existing overlay-mapped `internal/resolution/types.go` may carry the new metric fields only; it owns no behavior and must not change product types or production source.
- No production source, test, script, ledger, report other than the assigned measurement report, graphhealth, outcome, export-ordering, persistence, reader, CLI, or target source file may change.
- If direct-callee attribution cannot be represented in those existing overlay owners, STOP and return the exact missing owner; do not broaden instrumentation silently.

### Required non-overlapping D001 groups

Measure caller-side elapsed time and invocation count for these exact direct-call families inside `resolveCall`:

1. `source_context`: `sourceForScopeOrFile`.
2. `binding_receiver`: `repositoryReceiverClaimed`, `bindingOccurrences.resolve`, and `emitBindingOccurrenceReference`.
3. `scoped_same_file`: `resolveScopedName` plus both `resolveSameFileName` call sites.
4. `member_import`: `resolveMember`, `resolveImportedMemberWithProof`, `explicitImportNameClaimed`, and `explicitImportCallState`.
5. `go_same_package`: `resolveGoSamePackageFunction`.
6. `global_lookup`: all `resolveGlobalCallName` call sites.
7. `typescript_lookup_record`: `lookupTypeScriptMember` or `lookupTypeScriptGlobal` plus `recordTypeScriptLookup`.
8. `evidence_emission`: direct `appendExportBindingEvidence`, `emitUnresolvedReference`, and `emitReference` calls.
9. `direct_site_identity`: direct `callTargetText` and `sourceSiteID` evaluations in `resolveCall` that are not already inside a timed callee above.
10. `resolve_call_residual`: D001 elapsed not attributed to groups 1-9, including branch/control/metrics work.

The packet must prove, in integer nanoseconds:

```text
source_context
+ binding_receiver
+ scoped_same_file
+ member_import
+ go_same_package
+ global_lookup
+ typescript_lookup_record
+ evidence_emission
+ direct_site_identity
+ resolve_call_residual
= child_total_elapsed

overlap_count = 0
```

Nested work stays owned by its outer direct-call group. Do not double-count A002-A005 descendants as separate elapsed groups. Report the named invocation count beside every group so a zero or unsampled path is explicit rather than inferred.

### Exact build and invocation contract

1. Start from current production bytes, which are the accepted A003 implementation plus accepted WAL behavior and exact A004/A005 rollback.
2. Reuse unchanged `scripts/build-a00x-benchmark.ps1` and the accepted `17`-child overlay/native/runtime/provenance contract. Build one identifiable **measurement-only** executable under repository-local `E:\Anvien\.tmp`.
3. Run no accepted A003 baseline again and promote no instrumented timing. The accepted A003 values above remain the control basis.
4. Run the two instrumented captures **sequentially**, never overlapping, after zero-competitor preflight:
   - Cheapapp: `analyze E:\cheapapp.org --force --skip-git --json --progress` plus its new benchmark/profile/report paths.
   - Restaurant Manager: `analyze E:\Restaurant_manager --force --json --progress --exclude electron/renderer/src/api/userApi.ts` plus its new benchmark/profile/report paths.
5. Exactly one launch per target. No automatic retry.
6. Each target packet must retain `30/30` top-level operations, `17/17` children, the exact call/file denominator, parent/child conservation, zero overlap, workload identity, graph/DB counts, resolution counters, diagnostics/outcomes, complete ordered Evidence, canonical Graph JSON, stdout/stderr, target status, executable/overlay/native identity, and aggregate resource fields.
7. Instrumentation overhead is recorded honestly. These instrumented child/parent/process values are attribution evidence only and never replace or compare as a candidate against accepted A003.

### Decision rule for the next Architect

The next fresh A006 Architect may select a production direction only if `A006-M1` proves one unattempted group with nonzero exclusive elapsed on Cheapapp and Restaurant Manager separately, binds its exact dynamic work shape and owner, and supports a design that removes work synchronously without adding finalization, cross-run state, output drift, or downstream cost.

If the only material unattempted group remains `go_same_package` on Restaurant Manager, or the common groups remain too small/ambiguous, A006 still has no two-target production direction. Main then decides the next governance action; this report itself does not terminalize D001.

## Preserved Invariants For The Measurement And Any Future Direction

Preserve exactly:

- resolution branch order, confidence, proof kind, status, reason, and failure timing;
- SourceSiteID, target text, source-site status, evidence content/order, duplicate and first-error behavior;
- A001 import-claim index semantics and traversal order;
- A002 write-through diagnostic appender and run-scoped lifecycle;
- A003 canonical single-interpretation decoder and fail-closed tuple/policy behavior;
- restored pre-A004 evidence ordering;
- restored pre-A005 outcome encoding/projection behavior;
- graph nodes, relationships, IDs, labels, properties, reference indexes, outcomes, diagnostics, metrics, and ordering;
- canonical in-memory graph, Graph JSON, Ladybug/native persistence and readback, affected product/native readers, stdout, stderr, and public output;
- deterministic replay, freshness/invalidation, failure propagation, transaction rollback, temporary-file flush/close/rename, and publication visibility;
- accepted WAL force cleanup behavior;
- target commands, exclusions, denominators, and target separation;
- no global/cross-run cache, hidden persistence, goroutine, concurrency, I/O side channel, lock, flush, finalizer, or public schema change.

## Resource And Lifecycle Boundary

`A006-M1` instrumentation is run-local and measurement-only:

- fixed-size per-run counters/duration accumulators only;
- no per-call log, trace file, retained event list, or unbounded map;
- merged into the existing benchmark packet once;
- unreachable after the process exits;
- no state in target repositories beyond the normal accepted analyze outputs and assigned repo-local `.tmp` artifacts.

Any future production direction needs its own fresh resource boundary after this measurement; none is pre-authorized here.

## Rollback And Mandatory STOP

Measurement rollback removes only the A006-M1 overlay additions and its assigned `.tmp` packet/report if invalid. Production source remains untouched, so there is no A006 production rollback.

STOP the measurement and return to Main if:

- any production/test/ledger/plan/script/canonical build interface or target source edit is required;
- instrumentation requires an owner beyond overlay `resolve.go` and metric-carrier `types.go`;
- groups overlap, do not conserve D001, or omit a direct-call family;
- either target command, denominator, workload, output, graph/DB, semantic, ordering, or status boundary differs;
- captures overlap or a competitor is present;
- a launch fails, required packet field is absent, or a retry would be needed;
- A001-A005/WAL behavior is re-audited, renamed, bypassed, or altered;
- D002-D017, another parent, P3, or Child 07 is opened;
- anyone treats A006-M1 as a production attempt, candidate measurement, baseline promotion, Supervisor result, streak increment, or `SYSTEM_CHARACTERISTIC` disposition.

## Handoff

- Architecture direction released: none.
- Smallest missing input: `A006-M1-D001-DIRECT-CALLEE-ATTRIBUTION`, exactly as specified above.
- Missing-input owner: a separate visible read-only measurement executor using only the existing measurement-overlay boundary; not Planner or Coder.
- Planner and Coder remain locked.
- Next owner: Main Orchestration to verify this handoff and decide whether to authorize the measurement-only input.
- Stage/commit: none.

`ARCHITECT_A006_NEEDS_MEASUREMENT_INPUT`
