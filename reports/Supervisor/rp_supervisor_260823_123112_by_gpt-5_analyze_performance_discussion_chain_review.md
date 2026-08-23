# Supervisor Report: Analyze Performance Strategy Feasibility

Verdict: PASS

## Metadata

- Report file: `rp_supervisor_260823_123112_by_gpt-5_analyze_performance_discussion_chain_review.md`
- Review time: `2026-08-23 12:31:12 +07:00`
- Reviewer: `gpt-5`
- Repo/project: `E:/Anvien`
- Scope reviewed: technical feasibility and safety of the already-proposed incremental analyze optimization strategy only
- Claim reviewed: whether Main can present canonical-path reuse, an all-import claim index, a run-scoped diagnostic accumulator, a conditional single-pass decoder, and the accompanying cost-map/A-B/rollback protocol as the final best strategy without weakening accuracy or graph contracts
- Authority used: latest Owner narrowing; `AGENTS.md`; full `working-rules` and `supervisor` skills; current source; the sealed Root Cause, Architect, Planner, and amended graph-health anchor reports
- Acceptance boundary: strategy presentation only. This report does not authorize a plan, implementation, build, benchmark, target access, staging, or commit, and it does not claim any achieved speedup.

## Executive Summary

The proposed strategy is technically feasible and safe enough for Main to present. Each initial unit removes work that is directly visible in current source and retained profile evidence, while preserving the current semantic decision path and output order. The workflow makes the difficult diagnostic changes conditional on a fresh rebaseline and a clean ownership disposition, and it rejects every optimization that cannot prove exact output, native-reader, determinism, freshness, failure, publication, and resource equivalence.

No exact technical blocker makes the strategy infeasible or unsafe. The remaining gaps are correctly expressed as implementation prerequisites or Owner decisions: E-only build/dependency authority, a comparable unprofiled baseline and noise rule, numeric materiality/resource budgets, and clean ownership of the dirty graph-health anchor before any diagnostic edit. None is being misrepresented as already complete.

## Feasibility Matrix

| Claim | Independent finding | Evidence and safety boundary |
|---|---|---|
| Canonical-path reuse | Cleared | `resolveImports` canonicalizes every stored import path before append at `internal/resolution/indexes.go:287-307`; `explicitImportNameClaimed` and `explicitImportCallState` then normalize each visited stored path again at `internal/resolution/resolve.go:900-911` and `internal/resolution/export_resolution.go:306-350`. Direct comparison removes redundant allocation while retaining the same scan, order, early exit, and match predicate. Differential path cases and exact graph/native parity remain mandatory before acceptance. |
| All-import claim index | Cleared | The exact safe key is current source file plus exact local name. The proposal stores original `w.imports` indices and includes every import, including unresolved rows. This is essential because current `importsByReceiver` excludes `LinkStatus == unresolved` at `internal/resolution/indexes.go:305-307`. Name claims become membership checks; call-state still applies the current semantic-import, target, allowed-label, nil-target, and shadow rules over the original-order bucket. No resolved-only shortcut is accepted. |
| Run-scoped diagnostic accumulator | Cleared with the stated hard prerequisite | Current append behavior normalizes the incoming diagnostic at `internal/graphhealth/diagnostics.go:39` and again at line 94, re-normalizes the complete existing slice through lines 89 and 216-221, merges the first matching bucket at lines 95-104, and performs a full stable sort only after a new bucket at lines 107-122. The Architect/Planner contract preserves each of those exact behaviors, rebuilds the first-match index after every full sort, materializes the node property after every append, retains missing-node/fail-closed behavior, and reruns the six-field matrix. Current production search shows the diagnostic property is written only through this owner; resolution supplies the two production callers. The dirty anchor must first become a clean committed baseline or receive an independently reviewed supersession. |
| Conditional single-pass decoder | Cleared as conditional | Current `decodeStructuredResolutionOutcome` performs one whole-note unmarshal into an envelope and a second into the typed outcome at `internal/graphhealth/diagnostics.go:327-349`. A one-traversal presence-aware decoder is feasible, but it is not pre-authorized: it opens only if the post-accumulator rebaseline still measures the second pass as material. Missing versus null, malformed and wrong-type input, unknown fields, duplicate-key last-value behavior, nested raw values, and the exact `(outcome, structured, valid)` triple must match differentially. Any mismatch makes this unit N/A or rollback; marker-present fail-closed and six-field validation cannot be relaxed. |
| Cost-map and A/B protocol | Cleared | The proposed schema records run/runtime/corpus/cache identity, disjoint wall spans, CPU samples, allocations, live/peak memory, I/O, wait/block/mutex data, work units, call paths, artifact hashes, and explicit measured/estimated/not-instrumented status. It reconciles process wall with explicit residual, sums only disjoint leaves, and keeps profile attribution separate from unprofiled timing. A/A noise, predeclared trial count, alternating A/B order, identical cache regime, direct-cost reduction, whole-wall result, and resource budgets prevent opportunistic timing claims. |
| Accuracy and equivalence gates | Cleared | Counts alone are explicitly insufficient. A/B must match ordered outcomes and diagnostics, source sites, evidence, ResolutionGap data, graph bytes/SHA-256, canonical in-memory/Graph JSON/Ladybug digests, affected built readers, two forced deterministic replays, controlled freshness invalidation/restore, failure class, transaction rollback, temp cleanup, and publication identity. No workload reduction, stale cache, fallback relationship, skipped row, or weaker malformed-marker behavior can pass. |
| Rollback and dirty-anchor ownership | Cleared | Units are causally and commit-isolated: Unit 1 removes path transformations without an index; Unit 2 adds only the all-import index; rebaseline occurs before diagnostic work; Unit 3 precedes the conditional decoder. The verified seven-path anchor remains exactly `377,945` bytes / canonical SHA-256 `E4FF42DE8AA1941010AFFF5DD38B6789AE71168E3EBF9F59D8EECF735D5BD044`. M0/U1/U2 exclude all seven paths from staging. U3/U4 cannot open while `diagnostics.go` remains rejected and uncommitted. Each unit has mismatch, nondeterminism, stale-state, failure, publication, resource, direct-cost, and whole-wall rollback triggers plus an isolated commit boundary. |

## Causal Sufficiency

The retained Root Cause bundle is current-run causal evidence, not historical-regression evidence. Its primary artifact identities match the sealed report:

- benchmark JSON: `9,711` bytes / SHA-256 `66CBC05732B4AA2E55F6C3C84718EC802D82CB9CD608F6F529873704DD21187E`;
- CPU profile: `392,687` bytes / SHA-256 `35C6D4DD35284B67D467B4ADBBBAAFE71F5B134FB41F968F88921563E4E4D9FB`;
- heap profile: `107,005` bytes / SHA-256 `7874869AFE9F5B79F171F43F0FF9899CDC8CC95C57C2605AF80ABBBE97BD870A`.

The benchmark independently reconciles `596,568.8634 ms` total, `581,657.5740 ms` named phases, and `14,911.2894 ms` in-run residual. Resolution is `533,746.9551 ms` (`89.469462%`). Raw pprof text reports:

- `explicitImportNameClaimed`: `39.43 s` flat / `319.95 s` cumulative;
- `explicitImportCallState`: `3.90 s` flat / `47.51 s` cumulative;
- `cleanPath`: `0.88 s` flat / `324.75 s` cumulative, `149,590.52 MB` allocated, and `3,056,971,217` allocated objects;
- `AppendDiagnosticToNode`: `113.20 s` cumulative;
- `decodeStructuredResolutionOutcome`: `112.93 s` cumulative and `20,314.18 MB` allocated.

These inclusive contexts overlap and are not summed into promised savings. Source and profile evidence instead establish the removable mechanisms. Historical `~8 minute`, `17m20`, May, and P1-B timings remain non-comparable and are not used for attribution, threshold, or acceptance.

## Solution-to-Cause Match

### Unit 1: canonical path reuse

The unit removes exactly the repeated `cleanPath(item.Fact.FilePath)` work identified by source and profile while retaining `O(Q x I)` traversal. Stored import paths have a single construction owner and are canonicalized before entering `w.imports`; source-scope paths are returned canonically. The unit therefore has a narrow semantic surface and `O(1)` retained-memory change.

### Unit 2: all-import index

The unit removes the remaining whole-import traversal rather than making it merely cheaper. The proposed bucket payload is original import indices, not derived resolution answers, so all consumer checks remain current and ordered. Including unresolved and non-semantic imports in the bucket is necessary: name claims consider all imports, while call-state applies the semantic filter after path/name match. This distinction is preserved explicitly.

### Unit 3: diagnostic accumulator

The accumulator removes repeated normalization/decode and linear bucket search without deferring graph publication. The exact first-duplicate and later-full-sort corner cases are not hand-waved: they are named acceptance cases, with per-append legacy/new node-property comparison. Run-scoped ownership avoids stale cross-run state, and second-writer/property drift aborts rather than silently publishing mixed state.

### Unit 4: single-pass decoder

This unit removes one whole-note traversal only after Unit 3 establishes the new call count. It is safe precisely because the strategy allows it to be skipped. A custom presence-aware wire decode can preserve field presence and last-value behavior while retaining nested authority validation, but implementation acceptance depends on adversarial old/new differential evidence, not on the architecture claim alone.

## Cost Map, Equivalence, and No-Double-Count Clearance

The protocol is sufficient for strategy-level readiness:

1. Frozen pair identity covers runtime, source delta, dirty manifest, corpus, options, dependencies, native library, environment, storage prestate, instrumentation, command, timestamps, stdout, and stderr.
2. Cost centers carry parent/call-path identity plus wall, CPU, allocation, memory, I/O, wait, work-unit, call-count, and artifact fields.
3. `exclusive wall` is derived from the union of child spans; caller/callee inclusive samples, concurrent GC, and profiled versus unprofiled runs are never added together.
4. Every run retains an explicit unreconciled residual instead of assigning missing wall time by guess.
5. Direct cost-center improvement and whole-analyze behavior are both required; a faster whole wall alone cannot prove the mechanism.
6. Exact graph/output/native equivalence, deterministic replay, freshness invalidation, failure injection, rollback, publication, and resource gates prevent speed gained through semantic loss.
7. After each accepted unit, fresh absolute durations and counters rerank the next cost center. Old percentages lose priority authority.

## Dirty Anchor and Build Authority

The current amended graph-health report remains a historical `REJECT` anchor and is not promoted by this review. Its current bytes and seal are `33,046` / SHA-256 `3837059F852CF44F20732F2BB0F7D88234BC6612AAB9C7CD439AEC3288A768DA`. The exact seven-path aggregate and every per-path hash match the Architect/Planner reports.

Root `package.json` maps `full-build` to `scripts/full-build.ps1`; that script performs install/build/global discovery, launcher build, version checks, and final `anvien analyze . --force`. The workflow correctly blocks implementation until Owner/Main supplies explicit E-only authority for that canonical contract or an approved hermetic equivalent with sealed dependencies, runtime/native identity, cache policy, and name/PID-only holder preflight. This is a prerequisite, not evidence that the proposed algorithms are infeasible.

## No Blocking Finding

No exact source contradiction, unsafe semantic shortcut, missing algorithmic prerequisite, or ownership topology makes the proposed strategy infeasible or unsafe to present.

The following remain mandatory future gates and must not be presented as completed:

1. Owner acceptance of the campaign topology and later creation/review/commit of a real performance plan before code.
2. Canonical E-only build, dependency, runtime, cache, holder, and cleanup authority.
3. Comparable unprofiled A/A noise distribution, predeclared repetition/stopping rule, materiality floor, and resource budgets before candidate results.
4. Fresh corpus/dirty-anchor seals for every pair; any durable artifact that changes the scanned corpus forces a new pair identity.
5. A clean committed or independently superseded diagnostic baseline before U3/U4.
6. Fresh rebaseline after U1/U2 and after U3; U4 and all DB/snapshot/parse/post-run work remain conditional on the new absolute Pareto map.
7. Fresh implementation-time source, build, reader, output, failure, Supervisor, detect, cleanup, and commit evidence for each unit.

## Source-Level Clearance Notes

- `internal/resolution/indexes.go`: clear for Units 1-2. Single import construction normalizes before append; current receiver index is resolved-only and therefore explicitly rejected as the new index.
- `internal/resolution/resolve.go`: clear for Units 1-2. The hot name-claim scan and all call/access/type/receiver callers match the Root Cause attribution.
- `internal/resolution/export_resolution.go`: clear for Units 1-2. The independent call-state scan and its ordered post-match rules match the proposed bucket consumer.
- `internal/graphhealth/diagnostics.go`: clear for strategy feasibility only. Current append, first-match, stable-sort, double-decode, and six-field behavior match Units 3-4 contracts. Its dirty ownership still blocks implementation.
- `internal/resolution/emit.go`, `internal/resolution/outcome.go`, and `internal/graph/types.go`: clear for the run-scoped lifetime/materialization design; exact implementation impact remains future evidence.
- `internal/analyze/analyze.go` and `internal/cli/command.go`: clear for the cost-map need. Current benchmark timing omits pre-timer and several in-run/post-run substeps, so the proposed named spans and explicit residual are justified.
- `internal/lbugload/csv.go` and `internal/lbugload/load.go`: clear as conditional later evidence only. Aggregate relationship CSV is written while the loader consumes pair files; no DB optimization is activated until fresh substep timing and a separate exact contract exist.

## Evidence Checked

Passed:

- All five required chain/anchor reports match current bytes, LF/CR, strict UTF-8 without BOM, and supplied SHA-256 seals.
- All nineteen retained Root Cause artifacts exist; every reported primary and derived artifact size/hash matches.
- Benchmark phase math, workload counts, runtime/native identities, pprof hot paths, and allocation magnitudes match the Root Cause report.
- Current source directly confirms every algorithmic and semantic premise used by Units 1-4.
- Current seven-path anchor aggregate matches `377,945` bytes / SHA-256 `E4FF42DE8AA1941010AFFF5DD38B6789AE71168E3EBF9F59D8EECF735D5BD044`.
- Canonical full-build/analyze linkage in `package.json` and `scripts/full-build.ps1` matches the Planner prerequisite.

Failed:

- None for the strategy-feasibility claim.

Not run:

- No build, analyze, profile, benchmark, test, graph, target, Git, reader, cleanup, or implementation command was run in this review.
- No achieved speedup, comparable baseline, numeric threshold, implementation output equivalence, or implementation resource result exists yet; none supports this strategy-level decision.

## Invariant Closure

- Affected invariant: performance improvement must remove measured computation without changing semantic work, output, order, freshness, persistence/readers, failure behavior, or publication.
- Same-invariant surfaces reviewed: import construction and both claim scans; resolution call/access/type owners; diagnostic append/merge/sort/decode; graph node materialization; analyze phase/residual boundaries; CLI post-run owners; Graph JSON/Ladybug parity contract; dirty-anchor identity and canonical build linkage.
- Residual unverified same-invariant surfaces: implementation bytes and implementation-time runtime evidence do not exist by design. They are fully gated and cannot be reused or inferred from this discussion review.

## Overall Evaluation

The Root Cause, Architect, and Planner chain is technically traceable and internally consistent for the narrowed feasibility claim. The first three units directly remove measured work; the fourth is correctly conditional; the measurement protocol prevents double counting and accuracy shortcuts; and ownership/rollback boundaries keep every unit independently rejectable. Main may present this as the final recommended incremental strategy, provided the presentation states that it is a strategy acceptance, not implementation authorization or a speedup result.
