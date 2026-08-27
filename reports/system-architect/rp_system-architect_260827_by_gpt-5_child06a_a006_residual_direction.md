# Child 06A A006 Residual Architecture Decision

Current M2-or-stop governance decision is appended at EOF.

Accepted post-M1 production verdict preserved below: `ARCHITECT_A006_NO_SAFE_DIRECTION`

Historical pre-M1 verdict preserved below: `ARCHITECT_A006_NEEDS_MEASUREMENT_INPUT`

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

Historical pre-M1 terminal marker: `ARCHITECT_A006_NEEDS_MEASUREMENT_INPUT`

---

## Post-M1 Superseding Architecture Decision

Final verdict: `ARCHITECT_A006_NO_SAFE_DIRECTION`

This section consumes the completed `A006-M1-D001-DIRECT-CALLEE-ATTRIBUTION` packet and supersedes the historical pre-M1 request above. The historical request, its constraints, and its measurement contract remain preserved as provenance. This decision releases no production direction, does not terminalize D001, and does not request Planner or Coder.

### Current Scope, State, And Freshness

- Lane: fresh A006 System Architect, architecture only.
- Slice: `P2-A / A006 / B1-P1A-OP001 resolution / B2-P2A-A001-D001 resolve_calls`.
- Current checkpoint HEAD: `80ced4d7191fc70eb97187f626b89f64e2fd779e`; worktree clean; staged count `0` before this report-only update.
- Accepted implementation and measurement basis: A003 checkpoint `b6bf45bce95323aa6b53b182edfea8628bd8b463`.
- Entering campaign state: A004/A005 `SUPERVISOR_PASS / NO_KEEP / ROLLBACK_COMPLETE`; D001 unsuccessful streak `2`; parent/D001 unchecked; D002-D017 queued and unopened; P3 and Child 07 closed.
- Startup authority was reread through EOF in the mandated order: `AGENTS.md`, `working-rules`, `System-Architect`, `plan-rules.md`, all four standard Child 06A ledgers, this historical A006 report, the full M1 report, and the M1 handoff Supervisor report.
- `anvien --help` passed. Exactly one fresh `anvien analyze --force` then exited `0`: `2264` scanned, `766` parsed code, `0` failed; graph `124418` nodes / `171284` relationships.
- Fresh file-detail reports `internal/resolution/resolve.go` current, non-stale, unchanged since analyze, and `HIGH`: `200` symbols, inbound/outbound/local relationships `134 / 289 / 87`, `26` linked flows, and `40` linked tests.
- Fresh upstream impact including tests reports `resolveCall` `CRITICAL`: `31` impacted symbols, `1` direct caller, `13` affected files, `7` modules, and `32` processes. HIGH/CRITICAL is a blast-radius warning, not an edit prohibition; no source edit is authorized here.

### M1 Input Accepted Without Promotion

The M1 handoff is valid Architect input:

- the preserved initial `AttemptId=A006-M1` builder-input failure created no executable, target launch, value, attempt, or streak event;
- the one authorized recovery reused the unchanged overlay, manifest, builder, production, and tests with provenance `AttemptId=A006`;
- recovery build/root exits are `0/0`; provenance exposes exact overlay/candidate/native identities `2/4/3`;
- executable version/bytes/SHA-256 are `1.2.8 / 73,836,032 / 0362FE211C2E072988EF61B662FC4F0D2160437F520F4525CDA59DDE52FB57A7`;
- Cheapapp and Restaurant Manager each launched once and exited `0`, sequentially, with cross-target overlap `0`;
- each packet has `30/30` operations, `17/17` children, exact ten-group-to-D001 conservation, parent conservation, zero intra-target overlap, exact A003 Graph JSON/stdout/stderr identity, and exact non-timing graph/DB/resolution/diagnostic/outcome/evidence parity.

The instrumented D001, parent, analyzer, process, and resource values remain attribution-only. They are not a candidate improvement, accepted-baseline promotion, Supervisor result, production failure, or D001 streak event.

### Selection Gate Result

A production direction is eligible only if all of these are true at once:

1. M1 proves nonzero exclusive work on Cheapapp and Restaurant Manager separately.
2. The exact removable mechanism and exact current production owner are separated from A001-A005 ownership and from semantically mandatory carrier work.
3. One synchronous production algorithm removes that work rather than deferring it into finalization, another resolution child, a later analyzer phase, persistence, or CLI publication.
4. The direction is materially new rather than a cosmetic cache/index/byte/order variant of A001-A005.
5. The resulting resource lifetime and deterministic ordering can be bounded without global/cross-run state, concurrency, fallback behavior, output drift, or a second semantic authority.

No M1 group satisfies all five conditions. Nonzero aggregate timing alone is insufficient when the measured group combines several owners or when the proposed change would retain, defer, or merely rename required work.

### Exact Group-By-Group Disqualification

| M1 group | Cheapapp duration / invocations | Restaurant duration / invocations | Post-M1 architecture finding |
|---|---:|---:|---|
| `source_context` | `43,637,500 ns / 27,890` | `204,263,300 ns / 86,030` | Nonzero on both, but it measures `sourceForScopeOrFile` as a whole. Current source delegates to `callerForScope`, which walks parent scopes and owned definitions, then falls back to a file ref. M1 exposes no distinct-scope count, scope-hop count, repeated-key count, or caller-versus-file split. A cache/precomputed scope context could add a map and construction work without proving any repeated work is eliminated on either target. No exact removable mechanism is established. |
| `binding_receiver` | `78,596,800 ns / 26,079` | `185,310,600 ns / 105,656` | Composite of `repositoryReceiverClaimed`, binding-occurrence lookup, and binding-reference emission. The first may traverse scope/type/global claims and consumes accepted A001 import-claim semantics; the latter two create required reference/outcome/graph carriers. M1 does not split lookup from mandatory emission, report receiver keys/hops, or prove a repeated result safe to cache. |
| `scoped_same_file` | `53,094,400 ns / 29,686` | `266,276,000 ns / 65,948` | Composite of `resolveScopedName` and both `resolveSameFileName` sites. Source shows scope-chain binding lookup and a per-file definition traversal, but M1 gives neither per-callee elapsed/invocations nor binding hops, definitions visited, lookup-name distribution, or unique keys. A new memo/index direction would be an unproven A001-style variant and could move construction cost into `buildWorkspace`/resolution setup. |
| `member_import` | `186,935,000 ns / 47,123` | `308,333,200 ns / 167,208` | Composite of member resolution, imported-member/export proof, accepted A001 claim helpers, and five import-call-state sites. It does not isolate `defsByFile` traversal, semantic export resolution, receiver typing, inheritance lookup, or A001-preserved work. No candidate-visit or unique-key evidence proves one new scan/index mechanism on both targets. |
| `go_same_package` | `972,300 ns / 8,501` | `3,946,043,200 ns / 16,011` | Current source proves a whole `defsByFile` traversal in `resolveGoSamePackageFunction`, so the mechanism is concrete. It is nevertheless Restaurant-dominant and near-zero on Cheapapp. It cannot alone satisfy the mandatory two-target selection rule, and folding it into a broader definition-index proposal would rely on unmeasured portions of other composite groups. |
| `global_lookup` | `14,767,200 ns / 4,118` | `11,392,500 ns / 13,781` | Nonzero but small and already uses the existing `defsByName` index plus required label, arity, uniqueness, and ambiguity checks. M1 identifies no redundant scan, repeated key, or removable carrier work. A cache would add state without evidence of duplicate queries or net work removal. |
| `typescript_lookup_record` | `686,469,300 ns / 23,430` | `435,161,600 ns / 44,775` | Material and nonzero on both, but one interval deliberately combines catalog lookup, target-text argument evaluation, `recordTypeScriptLookup`, authority-result/external-site carriage, outcome recording, and conditional diagnostic emission. The record/diagnostic/outcome descendants include accepted A002/A003 and rejected A005 ownership. M1 does not separate catalog lookup from per-site recording, expose member/global/status splits, or report unique lookup keys/cache hits. A lookup cache might retain result state while mandatory per-site recording remains; a record-side change would repeat A002/A003/A005. No exact newly actionable owner is proven. |
| `evidence_emission` | `1,954,984,600 ns / 44,488` | `2,774,855,300 ns / 164,333` | Largest common group, but explicitly composite: `appendExportBindingEvidence`, `emitUnresolvedReference`, `emitReference`, `recordRepositoryUnresolvedOutcome`, `retainedExportResolutionForScopedBinding`, argument identity work, graph/reference-index mutation, outcome/diagnostic carriers, and relationship merge. Its descendants include accepted A002 appender, accepted A003 decoder, rejected A004 evidence dedupe/key/sort, and rejected A005 canonical outcome-byte lifecycle. The remainder creates semantically required graph/reference/outcome/diagnostic state. M1 provides no sub-owner timing that separates a new redundant synchronous operation from those attempted or mandatory paths. Deferral/batching would shift work downstream and alter write-through/order/failure timing. |
| `direct_site_identity` | `9,858,900 ns / 40,360` | `120,948,600 ns / 156,976` | Nonzero on both and source shows branch-local `callTargetText`/`sourceSiteID` construction. The group still combines required stable identity creation with any repeated target-text construction, and M1 does not separate the removable repeat. Stack-local target-text reuse would be a cosmetic micro-change, not a material new architecture direction, while source-site IDs and carrier target text must still be produced exactly. |
| `resolve_call_residual` | `19,683,300 ns / 27,890` | `44,540,800 ns / 86,030` | Exact nonnegative residual, but by definition it combines branch/control flow, assignments, metrics, recorder overhead, and any unlisted work. It has no exact production symbol or removable mechanism. Selecting it would be attribution by subtraction, not architecture evidence. |

The ten rows sum exactly to each target's instrumented D001 (`3,048,999,300 ns` and `8,297,125,100 ns`). That conservation validates the inventory; it does not turn a composite row into a safe production owner.

### Rejected Post-M1 Directions

1. **Go package/function index:** exact source mechanism, but no material Cheapapp signal. Rejected by the two-target rule.
2. **General definition lookup index using existing `defsByName`:** could touch same-file, imported-member, and Go paths, but M1 does not isolate their definition-scan share or candidate visits. It would combine three measured groups to manufacture a direction and risk becoming a cosmetic A001-style index extension.
3. **TypeScript lookup-result cache:** authority lookups are synchronous and deterministic, but duplicate/unique keys, member/global mix, status mix, lookup-only elapsed, clone/alias requirements, and retained memory are unmeasured. Per-site authority/outcome/diagnostic carriage remains mandatory and overlaps A002/A003/A005.
4. **Scope/caller memoization:** M1 proves aggregate source/scope work, not repeated scope keys or hops. A run-scoped cache may add as much lookup/retention work as it removes and lacks an evidenced resource bound tied to actual uniqueness.
5. **Call target/source-site reuse:** source proves limited repeated construction, but M1 does not isolate the duplicate portion from required identity work; the measured direct group is not a material production direction and cannot plausibly justify another full attempt after two NO_KEEP results.
6. **Emission batching or deferred projection:** forbidden because it moves the largest local group into finalization/later phases and changes write-through visibility, ordering, failure timing, graph/reference/outcome/diagnostic ownership, or downstream cost.
7. **Parallel call resolution:** forbidden without deterministic ownership of shared graph, relationship merge state, reference indexes, metrics, outcomes, diagnostics, TypeScript result slices, and ordered carriers. No such ownership proof exists.

### Full Preserved Product Path

The unchanged path remains:

`CLI analyze -> analyze.Run -> runPhase(PhaseResolution) -> ResolveBoundInto -> w.files -> ir.Calls -> resolveCall -> resolution/evidence/outcome/diagnostic/reference/graph carriers -> finalizeTypeScriptAuthorityResults -> emitTypeScriptExternalSymbols -> resolutionOutcomeCollector.finalize -> projectResolutionOutcomes -> resolution.Result -> analyze.Result -> MRO -> communities -> processes -> semantic enrichment -> Graph.Compact -> Ladybug/native load and readback -> canonical Graph JSON -> CLI/public output`.

No local D001 direction may transfer cost into another resolution child, final projection, graph mutation, DB load, snapshot, semantic processing, Ladybug, Graph JSON, or CLI publication. Every semantic, order, failure, lifecycle, and resource invariant named in the historical decision remains preserve-only.

### State Effect, Rollback, And Mandatory STOP

- Production direction: none.
- Production owner/test owner/algorithm: none released.
- Expected observable production gain: none claimed.
- Planner/Coder/Supervisor/measurement transition: none requested by this lane.
- Accepted state: exact A003 checkpoint remains authoritative.
- A006-M1 effect: Architect input only; not a failed production attempt.
- D001 streak: remains `2`; it is not incremented or terminalized.
- Checklist/queue: P2-A, parent, and D001 remain unchecked; D002-D017 remain queued/unopened; P3 and Child 07 remain closed.
- Production rollback: none, because this decision changes no production/test/script/plan/ledger/target byte. Report rollback, if Owner rejects the architecture record, is limited to this appended post-M1 section and the two verdict-label lines at the top/history boundary.
- Next owner: Main Orchestration. Main decides governance without treating M1 as an unsuccessful production attempt and without opening Planner/Coder from this verdict.

Mandatory STOP: do not infer a third A006 production attempt, `SYSTEM_CHARACTERISTIC`, child check, parent transition, Planner/Coder handoff, new measurement request, source/test edit, or downstream phase from this report. Any later production edit would require new evidence and a separately authorized fresh attempt chain; this A006 lane releases none.

`ARCHITECT_A006_NO_SAFE_DIRECTION`

---

## A006 M2-Or-Stop Governance Decision

Governance result: exactly one bounded M2 is justified; the terminal marker is recorded at EOF.

This section preserves both prior decisions above. `ARCHITECT_A006_NEEDS_MEASUREMENT_INPUT` remains the historical pre-M1 decision, and `ARCHITECT_A006_NO_SAFE_DIRECTION` remains the accepted post-M1 production-direction verdict. This follow-up releases exactly one measurement-only M2 packet; it does not release production work, Planner, Coder, Supervisor, a candidate comparison, a streak event, or a checklist transition.

### Current Boundary And Freshness

- Slice: `P2-A / A006 / B1-P1A-OP001 resolution / B2-P2A-A001-D001 resolve_calls`.
- Current clean checkpoint: `d6b5b954d02ea2f908e7319bebdfe5a29c6a9fd7`; staged count `0` before this report-only update.
- Accepted implementation and measurement baseline: A003 checkpoint `b6bf45bce95323aa6b53b182edfea8628bd8b463`.
- Campaign state remains `A004/A005 SUPERVISOR_PASS / NO_KEEP / ROLLBACK_COMPLETE / A006_M2_OR_STOP_PENDING / D001_STREAK_2`; parent/D001 are unchecked and D002-D017 are queued/unopened.
- Startup authority, all four updated ledgers, this report, accepted M1, the M1 handoff PASS, and Main's post-M1 handoff PASS were read through EOF.
- `anvien --help` passed. Exactly one fresh `anvien analyze --force` exited `0`: `2265` scanned, `766` parsed code, `0` failed; graph `124436 / 171302`; indexed/current commit both equal `d6b5b954d02ea2f908e7319bebdfe5a29c6a9fd7`, stale `false`.
- `internal/resolution/resolve.go` is current/non-stale/unchanged-since-analyze and `HIGH`: `200` symbols, `134 / 289 / 87` inbound/outbound/local relationships, `26` linked flows, and `40` linked tests. Fresh upstream impact including tests keeps `resolveCall` `CRITICAL`: `31` impacted symbols, `1` direct caller, `7` modules, and `32` processes. Exact helper LOW/zero graph impacts are non-assurance because their call sites include unresolved graph gaps; the containing-file/D001 warning controls.

### The One Eligible Composite Family And Exact Question

M2 selects only M1 family `typescript_lookup_record`, which is material and nonzero on both targets:

- Cheapapp: `686,469,300 ns / 23,430` invocations.
- Restaurant Manager: `435,161,600 ns / 44,775` invocations.

The exact unresolved question is whether one already-proven duplicate predicate inside that family owns nonzero exclusive work beyond matched timer overhead on **both** targets:

`internal/resolution/resolve.go::(*workspace).lookupTypeScriptMember:858-870` recomputes `repositoryReceiverClaimed(receiver, startScope)` at current line `862` even though the D001 caller has already computed the identical predicate at `resolveCall:405`, assigned it to the monotonic `externalBlocked` gate at line `406`, and can reach the TypeScript member call at lines `516-519` only when that earlier result is false.

No repository-claim input (`scopeBindings`, `typeBindings`, accepted A001 `importClaimsByReceiver`, parent scopes, or `defsByName`) is mutated between those two synchronous reads. The later predicate therefore has one exact candidate owner: the D001-only `resolveCall -> lookupTypeScriptMember -> repositoryReceiverClaimed` recheck. The other current helper callers at `resolveAccess:687` and `lookupTypeScriptAnnotation:816` remain preserve-only and are not attributed by M2.

This mechanism is materially new versus A001-A005: it removes a repeated read-only ownership traversal made redundant by an existing control-flow fact. It is not an index, cache, diagnostic path, decoder, evidence-order variant, outcome-byte lifecycle, deferral, batching, or concurrency proposal. M2 must measure it before any production direction can be judged.

### Exact Overlay Surface, Subgroups, And Counters

Measurement ID: `A006-M2-D001-TYPESCRIPT-MEMBER-RECEIVER-RECHECK`.

Allowed surface is exactly two new overlay copies under the one M2 root:

- `overlay/internal/resolution/resolve.go` for the existing ten-group recorder plus this one family split;
- `overlay/internal/resolution/types.go` for measurement-only metric carriers.

The executor must seed them from the accepted M1 overlay identities `resolve.go=6B77B55097664234AB9FA28102C84702F3A8B830D36C7449AF2F676242DB8A61` and `types.go=8F5F8E7B2B29FDDB7A65384970DAC60E7E15882608D9024FF6DF7414DD9F4E2D`. Current mapped production starts must remain `resolve.go=8CEEDBA1883314EE8883320D3647C25DEF6F19F043D57881A893FBA73BA210D9` and `types.go=7C5E113F5E50584665D6D9AED0BCB2B3C6F6085219A62CD5AE74DD7654CD5DC3`.

Keep the accepted M1 ten-group field unchanged and add exactly one metric object named `resolveCallTypeScriptLookupRecordMeasurement` with this schema:

```json
{
  "parentGroupId": "typescript_lookup_record",
  "parentDurationNs": 0,
  "parentInvocationCount": 0,
  "subgroups": [
    {"groupId":"member_receiver_recheck","durationNs":0,"invocationCount":0},
    {"groupId":"member_receiver_recheck_timer_control","durationNs":0,"invocationCount":0},
    {"groupId":"member_lookup_remainder","durationNs":0,"invocationCount":0},
    {"groupId":"global_lookup","durationNs":0,"invocationCount":0},
    {"groupId":"site_target_text","durationNs":0,"invocationCount":0},
    {"groupId":"record_typescript_lookup","durationNs":0,"invocationCount":0},
    {"groupId":"typescript_lookup_record_residual","durationNs":0,"invocationCount":0}
  ],
  "work": {
    "siteCount": 0,
    "memberLookupCount": 0,
    "globalLookupCount": 0,
    "priorUnclaimedMemberLookupCount": 0,
    "memberCatalogEligibleCount": 0,
    "receiverRecheckCount": 0,
    "receiverRecheckFalseCount": 0,
    "receiverRecheckTrueCount": 0,
    "targetTextCount": 0,
    "recordCount": 0
  },
  "subgroupSumNs": 0,
  "overlapCount": 0
}
```

Instrumentation sites and ownership are exact:

1. Preserve the existing M1 outer interval at `resolveCall:516-535`; it supplies `parentDurationNs` and `parentInvocationCount`.
2. At the D001 member call site `resolveCall:519`, record `memberLookupCount` and the already-held false predicate as `priorUnclaimedMemberLookupCount`; pass the fixed recorder only through this measurement-overlay call. D002/D003 callers pass no recorder and retain their exact behavior.
3. Inside overlay `lookupTypeScriptMember`, preserve the nil/library/language guards. When the D001 recorder is present and those guards pass, record `memberCatalogEligibleCount`, measure one empty adjacent `time.Now/time.Since` pair as the matched timer control, then measure only the direct `repositoryReceiverClaimed` call and record its true/false result. Recorder accumulation occurs after both local elapsed values are captured and is outside both intervals.
4. Measure the complete D001 member helper at its caller; publish `member_lookup_remainder = member helper elapsed - receiver recheck elapsed - matched control elapsed`. Measure `lookupTypeScriptGlobal`, the single `callTargetText` evaluation used by the TypeScript source-site record, and `recordTypeScriptLookup` separately at their existing caller sites and in their existing evaluation order.
5. Compute `typescript_lookup_record_residual` only after the run as the parent duration minus the other six subgroup durations. No subgroup may overlap another.

The recorder is one fixed `[7]` row array plus scalar integer counters. No per-call event, slice growth, map, key cache, log, trace, goroutine, lock, I/O, global/cross-run state, or target-repository state is permitted.

For each target, the packet must prove these integer equations and predicates:

```text
member_receiver_recheck
+ member_receiver_recheck_timer_control
+ member_lookup_remainder
+ global_lookup
+ site_target_text
+ record_typescript_lookup
+ typescript_lookup_record_residual
= parentDurationNs

parentDurationNs
= resolveCallDirectCalleeMeasurements["typescript_lookup_record"].durationNs

parentInvocationCount
= memberLookupCount + globalLookupCount + targetTextCount + recordCount

siteCount = memberLookupCount + globalLookupCount
siteCount = targetTextCount = recordCount
priorUnclaimedMemberLookupCount = memberLookupCount
memberCatalogEligibleCount = receiverRecheckCount
receiverRecheckFalseCount + receiverRecheckTrueCount = receiverRecheckCount
member_receiver_recheck_timer_control.invocationCount = receiverRecheckCount
subgroupSumNs = parentDurationNs
overlapCount = 0
```

Every remainder must be nonnegative. The existing ten M1 groups must still sum exactly to D001 with overlap `0`; M2 does not re-accept or replace M1.

### One Reusable Build And Execution Contract

Use exactly one new repo-local output root, currently absent:

`E:\Anvien\.tmp\child06a_a006_m2_typescript_member_recheck_20260827`

All overlay, process-local home/cache/temp, build, capture, comparison, and validation artifacts must remain beneath that root. The only durable result outside it is:

`E:\Anvien\reports\Investigation\rp_child06a_a006_m2_typescript_member_receiver_recheck.md`

The overlay manifest is exactly `overlay\optimized-overlay.json` with exactly the two mappings above. Reuse unchanged builder `E:\Anvien\scripts\build-a00x-benchmark.ps1`, current SHA-256 `ADA407C7496FCEA988276F03BAD5001ED139A4AEC9A16B9C32947DA440814EC5`, once and without retry. Its invocation is fixed as:

```powershell
& 'E:\Anvien\scripts\build-a00x-benchmark.ps1' `
  -AttemptId 'A006' `
  -OverlayManifestPath 'E:\Anvien\.tmp\child06a_a006_m2_typescript_member_recheck_20260827\overlay\optimized-overlay.json' `
  -OutputExecutablePath 'E:\Anvien\.tmp\child06a_a006_m2_typescript_member_recheck_20260827\frozen-build\anvien-a006-m2-benchmark.exe' `
  -ExpectedOverlaySha256 '<literal validated M2 manifest SHA-256>' `
  -ExpectedMappedSourceHash @(
    'E:\Anvien\internal\resolution\resolve.go=<literal validated M2 overlay resolve.go SHA-256>',
    'E:\Anvien\internal\resolution\types.go=<literal validated M2 overlay types.go SHA-256>'
  ) `
  -ExpectedCandidateSourceHash @(
    'internal/graphhealth/diagnostics.go=6DE54D97A1B95E877B686DC3459238598F6060D28CE6329357A80BD7FC376D30',
    'internal/resolution/emit.go=73AE20E8A18AC25F8019E2A3223407621AE7ABC358DBAF4F058E1D9AC3136060',
    'internal/resolution/outcome.go=02092F9FE7DA2A4BDB49E13056FBA3C97DC24F416141E7F27866EE80F60C1F7E',
    'internal/resolution/export_binding_proof.go=4BFE1976766FBF2E2257102070FF43CAC6D757E32E71DD583F8824781EAB6A8E'
  ) `
  -ExpectedNativeHash @(
    'lbug.h=3D5114D0863B3DAB3B28BD2FEC97A52E6CF669213739921A01814A5BBF5525EB',
    'lbug_shared.lib=B18AAFC0B712DC1C4CB9DD25F76C3828282D7D460627980E3A4B16EFCD98A955',
    'lbug_shared.dll=20CBD87840483A2053CFF3FC2DB23A86DD802B8915D86509D41A4B709624CDB7'
  ) `
  -GoExecutable 'C:\Program Files\Go\bin\go.exe'
```

The executor must replace the three angle-bracket values with the freshly computed literal uppercase hashes before invocation; placeholders in the real command are forbidden. Provenance must report schema/attempt/build/root exits `1/A006/0/0`, exact identities `2/4/3`, and matching expected/actual hashes. No canonical script or production file is edited.

After zero-competitor preflight, run exactly one launch per target, sequential Cheapapp then Restaurant Manager, with no retry:

```text
E:\Anvien\.tmp\child06a_a006_m2_typescript_member_recheck_20260827\frozen-build\anvien-a006-m2-benchmark.exe analyze E:\cheapapp.org --force --skip-git --json --progress --benchmark-json E:\Anvien\.tmp\child06a_a006_m2_typescript_member_recheck_20260827\captures\cheapapp\candidate\benchmark.json --benchmark-label child06a-a006-m2-cheapapp-ts-member-recheck-20260827

E:\Anvien\.tmp\child06a_a006_m2_typescript_member_recheck_20260827\frozen-build\anvien-a006-m2-benchmark.exe analyze E:\Restaurant_manager --force --json --progress --exclude electron/renderer/src/api/userApi.ts --benchmark-json E:\Anvien\.tmp\child06a_a006_m2_typescript_member_recheck_20260827\captures\restaurant_manager\candidate\benchmark.json --benchmark-label child06a-a006-m2-restaurant-manager-ts-member-recheck-20260827
```

Use working directories `E:\cheapapp.org` and `E:\Anvien`, respectively. Set `ANVIEN_OP001_RESOLUTION_CPU_PROFILE` respectively to `E:\Anvien\.tmp\child06a_a006_m2_typescript_member_recheck_20260827\captures\cheapapp\candidate\resolution.cpu.pprof` and `E:\Anvien\.tmp\child06a_a006_m2_typescript_member_recheck_20260827\captures\restaurant_manager\candidate\resolution.cpu.pprof`; redirect every process-local home/cache/temp path beneath the corresponding capture root. Beside each target's `benchmark.json`, record exact files `process.json`, `stdout.txt`, `stderr.txt`, and a copied canonical `graph.json`. Write machine-readable validation exactly to `E:\Anvien\.tmp\child06a_a006_m2_typescript_member_recheck_20260827\validation\cheapapp.json`, `E:\Anvien\.tmp\child06a_a006_m2_typescript_member_recheck_20260827\validation\restaurant_manager.json`, and `E:\Anvien\.tmp\child06a_a006_m2_typescript_member_recheck_20260827\validation\cross-target.json`.

Cheapapp must finish before Restaurant preflight and start. Both immediate competitor lists must be empty; cross-target capture overlap must equal `0`.

### Validation, Resource, Completion, And Falsification Rules

Each target must independently retain exact `30/30` operations, ordered `17/17` children, denominators `27,890/887` and `86,030/1,234`, parent/child/residual conservation, ten-group D001 conservation, the new TypeScript-family conservation/count equations, overlap `0`, and nonempty packet/profile/process artifacts. Validate exact A003 workload, file/parser, graph/DB, resolution counter, diagnostic, outcome, complete ordered Evidence, canonical Graph JSON, stdout, stderr, target HEAD/status, and source/native/provenance parity after excluding only authorized M2 timing/counter and ordinary timing/resource fields.

Record `startAllocBytes`, `endAllocBytes`, `maxObservedSys`, process wall/CPU/user/kernel, and the fixed `O(1)` recorder lifetime. These are M2 observer/resource facts only. No M2 D001, parent, analyzer, process, or resource value is compared with or promoted over A003.

M2 exposes the one exact production-owner candidate without another attribution packet only when, on Cheapapp and Restaurant Manager separately, all packet gates pass and:

```text
receiverRecheckCount > 0
member_receiver_recheck.durationNs - member_receiver_recheck_timer_control.durationNs > 0
receiverRecheckTrueCount = 0
receiverRecheckFalseCount = receiverRecheckCount
priorUnclaimedMemberLookupCount = memberLookupCount
```

If exposed, a fresh production Architect can judge immediately one synchronous call-only algorithm: preserve the existing nil/library/language, receiver trim, catalog lookup, result, order, and failure behavior, but route only D001's already-unclaimed member fallback through one private helper/path that omits the second `repositoryReceiverClaimed`; keep the existing checked helper for D002/D003. This removes the measured read-only traversal in place, adds no retained state, and shifts no work downstream. This paragraph is a post-M2 decision hypothesis, not production authorization.

If either target has zero rechecks, no positive elapsed above the matched control, any true recheck, or a failed count predicate, this exact owner is falsified. Main must not request M3, another family split, a cache-key packet, or broad instrumentation. The valid M2 result then returns directly to governance with no production direction from this packet.

STOP and return to Main if M2 needs a third overlay owner, production/test/script/canonical-build-interface/target edit, another M1 family, per-site retained data, a map/cache, a retry, overlapping targets, a competitor, a missing field, negative remainder, failed conservation, denominator/workload/output/order/semantic drift, or another attribution measurement before production can be judged.

Completion is exactly two valid target packets plus the one Investigation report and the predicates above. M2 remains attribution-only: no candidate comparison, baseline promotion, D001 streak/checklist/disposition effect, Supervisor, detect, stage, or commit. Next owner is Main Orchestration; after valid M2, Main may return it to one fresh A006 production Architect for the immediate expose-or-falsify decision, never to Planner/Coder directly.

`ARCHITECT_A006_M2_READY`

---

## Post-M2 Production/No-Production Architecture Decision

Final post-M2 result: no further measurement is justified.

This decision consumes the accepted M2 falsification without reopening the post-M1 production-direction review. The slice remains `P2-A / A006 / B1-P1A-OP001 resolution / B2-P2A-A001-D001 resolve_calls`; checkpoint `506490011c190fbeeaf81db8c2a13da09d87420a` is the authority basis, and A003 checkpoint `b6bf45bce95323aa6b53b182edfea8628bd8b463` remains accepted.

### Decisive Evidence

- Both M2 target packets pass every identity, count, conservation, parity, resource-lifetime, sequentiality, and overlap gate.
- Cheapapp records receiver recheck/control/net `10,466,100 / 1,000,600 / 9,465,500 ns`, with false/true counts `3,549 / 0`.
- Restaurant Manager records receiver recheck/control/net `0 / 0 / 0 ns`, with false/true counts `638 / 0`.
- Cross-target gap is `213,628,851,200 ns` and overlap is `0`.
- The binding exposure rule required positive recheck-minus-control elapsed on both targets. Restaurant Manager's valid `0 ns` result therefore establishes `A006_M2_RECEIVER_RECHECK_FALSIFIED`; the duplicate receiver-claim read is not a two-target production owner.

The accepted post-M1 disqualifications remain controlling for every other M1 family. `go_same_package` is Restaurant-dominant and near-zero on Cheapapp. `evidence_emission` remains a composite of mandatory graph/reference/outcome/diagnostic carriers plus accepted or already attempted A002-A005 ownership. `typescript_lookup_record` was the only material common composite family eligible for one final bounded split, and M2 falsified its sole exact new synchronous owner. The remaining groups are small, composite, already indexed, mandatory identity/carrier work, or lack an exact removable owner and synchronous cost-removal algorithm. M2 supplies no evidence that overturns those findings.

### Hard Loop Stop And Campaign Blocker

No materially new synchronous mechanism is now proven to have nonzero removable exclusive work on both targets. Any attempt to select another owner would require a forbidden M3, another family split, cache-key packet, profile/source discovery, or broad instrumentation and would still need another measurement before production could be judged. That would create the attribution/audit loop this boundary exists to stop.

The concrete campaign blocker for Main is therefore the absence of an evidence-proven, two-target, synchronously removable D001 owner outside A001-A005 after the sole eligible final split was falsified. No production owner, algorithm, test surface, expected gain, Planner, Coder, Supervisor, or rollback is released.

### State Effect And Handoff

- M1 and M2 remain attribution-only. Neither is a production attempt, candidate, no-KEEP event, `SYSTEM_CHARACTERISTIC`, or disposition.
- Accepted A003 remains unchanged; D001 streak remains `2`.
- P2-A, parent, and D001 remain unchecked. D002-D017 remain queued/unopened; P3 and Child 07 remain closed.
- M3 and every further A006 measurement/discovery packet are stopped. No graph/source refresh was needed because this decision releases no new production direction and canonical source is unchanged.
- Next owner: Main Orchestration, for governance only. Main must not infer a Planner/Coder release or terminalize D001 from this report.

`ARCHITECT_A006_NO_FURTHER_MEASUREMENT_JUSTIFIED`
