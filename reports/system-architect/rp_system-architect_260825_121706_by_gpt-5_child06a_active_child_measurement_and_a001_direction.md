# System Architect Report — Child 06A Active-Child Measurement Process and A001 Direction

- **Status:** `DECIDED / READY_FOR_PLANNER_TRANSLATION`
- **Date:** `2026-08-25` (`UTC+7`)
- **Attempt:** `A001`
- **Active parent:** `B1-P1A-OP001 resolution`
- **Selected child:** `B2-P2A-A001-D001 resolve_calls`
- **Planner target:** `E2-P2A-A001PLAN1`
- **Send to:** Main Orchestration, then a separate visible Planner lane
- **Not:** a Coder authorization, implementation result, benchmark result, disposition, or Supervisor PASS

## 1. Scope

This report finalizes the measurement and attribution process for Child 06A and binds it to the already-selected A001 technical direction.

The central distinction is:

1. **Child total elapsed time** controls ranking, before/after comparison, and retention decisions.
2. **Exclusive active-child breakdown** explains where graph-producing work is spent, including preparation, lookup, resolution, graph mutation, outcomes, diagnostics, and residual work.

“Graph-producing work” is broader than calls that directly add graph nodes or relationships. Any required preparation or resolution work that contributes to the same final graph belongs inside the child total even when it executes before graph mutation.

This report does not edit the plan or ledgers. Planner must translate it into the current Child 06A plan surface.

## 2. Authority and evidence basis

Authority read for this decision:

- `E:\Anvien\AGENTS.md`
- `E:\Anvien\.agents\skills\working-rules\SKILL.md`
- `E:\Anvien\.agents\skills\System-Architect\SKILL.md`
- Child 06A `plan-rules.md`, plan, evidence, benchmark, and actual-status ledgers
- `reports/Investigation/rp_main_260824_204615_child06a_optimization_method_discussion.md`
- `reports/Investigation/rp_huong_dan_lam_viec_de_slice_p6e_tien_hanh_nhanh_hon.md`
- Exact carried source-map, measurement, and CPU-profile evidence for D001
- Current `internal/resolution` source and direct consumers/tests

Anvien architecture evidence already obtained for A001:

- `anvien --help`: PASS
- `anvien analyze --force`: PASS, exit `0`, scanned `2209`, parsed `765`, failed `0`
- Read-only file detail for `internal/resolution/resolve.go`, `indexes.go`, and `export_resolution.go`
- `resolveCall`, `workspace`, `buildWorkspace`, and `resolveImports` have CRITICAL blast-radius evidence

The graph refresh was freshness/architecture evidence only. It is not a Child 06A benchmark capture, speed result, or output-equivalence proof.

## 3. Current numeric basis and child queue

Current selected-child basis:

- D001 `resolve_calls`: `243.794553800 s`
- Calls: `72976`
- Files: `765`
- Share of same-run parent: `44.697337%`
- Parent `resolution`: `545.434182000 s`
- Process/analyzer: `653.475797100 s / 630.160832200 s`
- Child sum/residual: `545.364656400 s / 0.069525600 s`
- Parent intervals: `2309`; overlaps: `0`
- Current no-KEEP streak: `0`

All 17 child rows and checkboxes remain unchecked:

| Rank | Child | Operation | Same-run elapsed |
|---:|---|---|---:|
| 1 | `B2-P2A-A001-D001` | `resolve_calls` | `243.794553800 s` |
| 2 | `B2-P2A-A001-D002` | `resolve_accesses` | `204.093359700 s` |
| 3 | `B2-P2A-A001-D003` | `resolve_type_annotations` | `92.316121800 s` |
| 4 | `B2-P2A-A001-D004` | `project_resolution_outcomes` | `4.067254700 s` |
| 5 | `B2-P2A-A001-D005` | `emit_definition_nodes` | `0.387684000 s` |
| 6 | `B2-P2A-A001-D006` | `finalize_resolution_outcomes` | `0.328439300 s` |
| 7 | `B2-P2A-A001-D007` | `build_binding_occurrence_index` | `0.172202300 s` |
| 8 | `B2-P2A-A001-D008` | `finalize_typescript_authority_results` | `0.096993100 s` |
| 9 | `B2-P2A-A001-D009` | `emit_import_edges` | `0.088903700 s` |
| 10 | `B2-P2A-A001-D010` | `emit_typescript_external_symbols` | `0.011376900 s` |
| 11 | `B2-P2A-A001-D011` | `emit_method_dispatch_edges` | `0.007767100 s` |
| 12 | `B2-P2A-A001-D012` | `assemble_resolution_result` | `0 s` |
| 13 | `B2-P2A-A001-D013` | `binding_accumulator_dispose` | `0 s` |
| 14 | `B2-P2A-A001-D014` | `emit_heritage_compatibility_edges` | `0 s` |
| 15 | `B2-P2A-A001-D015` | `emit_unresolved_heritage_diagnostics` | `0 s` |
| 16 | `B2-P2A-A001-D016` | `finalize_resolution_metadata` | `0 s` |
| 17 | `B2-P2A-A001-D017` | `runtime_setup` | `0 s` |

D002–D017 remain queued and are not implementation scope for A001.

## 4. Final measurement architecture decision

### 4.1 Ranking contract

`child_total_elapsed` is the sole timing used to:

- rank children;
- select the active child;
- compare the active child before and after an attempt;
- decide whether the implementation produced observable child improvement.

No internal sub-metric, CPU cumulative sample, microbenchmark, build duration, or graph-mutation-only duration may replace `child_total_elapsed` for ranking or retention.

### 4.2 Active-child exclusive breakdown

For the active child only, attribute elapsed time into mutually exclusive groups:

| Metric | Exact purpose |
|---|---|
| `prepare_lookup_elapsed` | Path/input normalization, index/import scanning, and candidate preparation or lookup before a semantic decision |
| `resolution_decision_elapsed` | Binding, member, package, authority, target, proof, and resolution-outcome decision work not already owned by lookup |
| `graph_mutation_elapsed` | Node, relationship, reference, property, and metadata creation or merge work |
| `diagnostic_outcome_elapsed` | Resolution outcome and unresolved/diagnostic recording |
| `metrics_bookkeeping_elapsed` | Metrics/counter bookkeeping not already owned by another category |
| `residual_elapsed` | Child elapsed not attributed to the five explicit groups |

Mandatory conservation rule:

```text
prepare_lookup_elapsed
+ resolution_decision_elapsed
+ graph_mutation_elapsed
+ diagnostic_outcome_elapsed
+ metrics_bookkeeping_elapsed
+ residual_elapsed
= child_total_elapsed

overlap_count = 0
```

Nested independent timers must not be summed. If a lookup helper executes inside a broader resolution operation, the elapsed interval belongs to exactly one exclusive category, normally the most specific category.

### 4.3 Work-count contract

Elapsed time must be accompanied by workload counters that demonstrate equivalent work. The exact relevant counters depend on the active child. For D001 they are:

- files and calls;
- import-claim helper invocations;
- global-import candidates visited;
- import-path normalizations at claim lookup sites;
- index entries and unique lookup keys;
- matching-bucket candidates visited;
- resolution targets considered;
- references and relationships emitted;
- outcomes and diagnostics recorded;
- graph node/relationship/property inventories at affected boundaries.

Counters describe work; they do not replace wall time for optimization acceptance.

### 4.4 Observer-overhead contract

- Do not emit a log, trace record, or standalone timer artifact for each of the `72976` calls.
- Use the smallest cumulative timers/counters that answer the active causal question.
- Prefer one local accumulator per helper/block and merge it once, rather than atomic or external recording per candidate.
- Preserve the same instrumentation identity before and after an attempt.
- Record instrumentation overhead/comparability. An invalid or mixed-instrumentation capture cannot rank or prove improvement.

## 5. Active-child-only sequencing decision

Deep drill-down is performed for exactly one active child at a time.

```text
Measure coarse totals for the complete 17-child list
→ select the largest current unchecked child
→ deep-measure only that active child
→ verify the cost center through source/profile/work counts
→ obtain a fresh attempt-scoped Architect direction
→ Planner translates the direction
→ Coder implements production first and tests second
→ run the canonical full build
→ remeasure active child + parent + end-to-end on the same work
→ run affected equivalence/resource checks
→ obtain a separate visible Supervisor result and disposition
→ continue the same child until terminal
→ rerank the complete current child list
→ only then deep-measure the next selected child
```

Reasons not to deep-instrument D002–D017 now:

- Their coarse ordering already exists.
- D001 currently owns `44.697337%` of parent elapsed.
- A retained D001 change may alter shared-helper costs and therefore later child totals.
- Deep evidence for inactive children can become stale before those children are selected.
- Deep instrumentation across all children increases observer overhead and attribution complexity.
- Rows currently measured as `0 s` are not treated as cost-free system characteristics; they are remeasured and drilled down only when selected by current ranking.

Coarse child-total instrumentation may remain enabled for all 17 children during the same parent capture. Only detailed attribution is limited to the active child.

## 6. D001-specific drill-down mapping

For `resolve_calls`, the exclusive diagnostic grouping is:

1. Explicit import-claim lookup and path comparison.
2. Scope/binding/member/package/global/TypeScript authority search and resolution decision.
3. Reference and relationship emission.
4. Resolution outcome and unresolved diagnostic recording.
5. Metrics bookkeeping.
6. Residual.

Existing source and profile evidence already proves that repeated import-claim scans are a current cost center:

- `explicitImportNameClaimed` scans all `w.imports` candidates and normalizes each visited source path.
- `explicitImportCallState` also scans all imports and normalizes candidate paths.
- `resolveCall` may invoke these helpers more than once depending on call form and fallback path.
- `repositoryReceiverClaimed` invokes the name-claim scan for member receivers.
- Current workload has `4887` imports and `72976` calls.
- One hypothetical full scan per call is `356633712` candidate visits; this is a structural scale calculation, not an observed visit counter.

CPU cumulative samples support the mechanism but overlap and must not be added:

- `resolveCall`: `227.45 s`
- `cleanPath`: `161.32 s`
- `explicitImportNameClaimed`: `134.42 s`
- `repositoryReceiverClaimed`: `72.02 s`
- `explicitImportCallState`: `49.27 s`

## 7. A001 technical direction remains selected

The measurement process does not replace the current A001 intervention. It validates and governs it.

Selected direction:

```text
Create a separate run-scoped all-import claim index:

{canonical source file path, exact local import name}
    → []original w.imports index in existing order
```

Required behavior:

- Build the index during import construction after source-path canonicalization.
- Include resolved, unresolved, semantic, nonsemantic, and duplicate imports.
- Preserve exact `w.imports` order by storing original indices in append order.
- Use the index only to replace global candidate selection in the two import-claim helpers.
- Keep semantic filtering, target checks, allowed/import-target sets, target-nil behavior, lexical-shadow behavior, and import-export resolution unchanged.
- Do not reuse or change the semantics of the resolved-only `importsByReceiver` index.
- Do not cache final resolution answers.
- Do not iterate map keys to determine resolution order.
- Do not change `resolveCall` branch ordering, canonical path rules, import-target resolution, graph emission, persistence, readers, or lifecycle behavior.

Expected algorithmic change:

```text
Before: repeated O(helper queries × all imports) scans
After:  O(imports) run-scoped construction
        + expected O(1) key lookup
        + O(matching bucket size) candidate traversal
```

No elapsed-time threshold or fabricated saving is authorized. The attempt must prove improvement through fresh child, parent, and end-to-end measurements.

## 8. Exact implementation ownership for Planner translation

Planner may authorize only these production owners:

- `internal/resolution/indexes.go`
  - `workspace`: one private all-import claim index field
  - `buildWorkspace`: initialize that field
  - `(*workspace).resolveImports`: populate original import indices in existing order
- `internal/resolution/resolve.go`
  - body of `(*workspace).explicitImportNameClaimed`
- `internal/resolution/export_resolution.go`
  - candidate-selection loop of `(*workspace).explicitImportCallState`

Preserve-only production surfaces:

- `resolveCall`
- `repositoryReceiverClaimed`
- `cleanPath`
- `scopeFilePath`
- import target/file resolution
- graph emitters
- persistence and readers
- every exported/shared contract

If implementation needs another production owner or helper, Planner must stop and return for a fresh architecture decision.

Tests may be updated only after production behavior is complete, and only in:

- `internal/resolution/export_resolution_test.go`
- `internal/resolution/resolution_test.go`

Other affected tests remain run-only/preserve-only.

## 9. Preserved invariants

Planner must copy these invariants into `E2-P2A-A001PLAN1`:

- Exact resolution outcomes, confidence, proof kind, unresolved notes, and call-branch order.
- Exact query trimming and case-sensitive import-local-name semantics.
- Original import ordering, duplicates, first-match behavior, and traversal order.
- Current canonical path behavior without lowercase, absolutization, or symlink resolution.
- Relative, absolute, language-specific, and Go/package import handling.
- Unresolved import claim and global-rescue blocking behavior.
- Nonsemantic import name claims and semantic call-state filtering.
- Graph node, relationship, ID, label, property, count, and ordered-output equivalence.
- Resolution outcome/index equivalence.
- Graph JSON and Ladybug/native persistence/readback parity.
- Deterministic replay and output.
- Freshness, failure propagation, transaction rollback, temporary-file lifecycle, and publication visibility.
- Private run-scoped in-memory lifecycle; no new serialization, I/O, goroutine, global cache, or public schema.

## 10. Expected observable result

For a conforming implementation:

- Global `w.imports` scans inside the two claim helpers disappear.
- Repeated `cleanPath(item.Fact.FilePath)` calls at those two helper sites disappear.
- Claim lookup candidate visits become the sum of matching bucket sizes rather than repeated whole-import scans.
- Index construction remains bounded by current import count and unique keys.
- Files, calls, imports, targets, outcomes, diagnostics, nodes, relationships, and output content remain equivalent.
- D001 child elapsed should improve; parent resolution and end-to-end analyze should receive corresponding net benefit if construction/allocation overhead remains smaller than removed work.

The child, parent, and end-to-end wall times—not the internal lookup timer—control performance acceptance.

## 11. Validation and resource boundary

Planner must require:

1. Fresh Coder pre-edit `anvien --help`, `anvien analyze --force`, and exact file-detail/impact evidence. The Architect refresh does not replace this gate.
2. Production implementation before test edits.
3. Focused correctness tests and unchanged affected equivalence tests.
4. Canonical full build:

   ```text
   pwsh -NoLogo -NoProfile -NonInteractive -ExecutionPolicy Bypass -File E:\Anvien\scripts\full-build.ps1
   ```

5. Same-instrumentation/same-work measurement in this order:
   - D001 `resolve_calls`
   - parent `resolution`
   - end-to-end graph generation/analyzer
6. Work-count and output equivalence at affected graph, outcome, persistence, and reader boundaries.
7. Allocation/resource evidence showing memory remains `O(imports + unique keys)` and the index stores indices rather than duplicate import objects.
8. Preservation of carried P1 instrumentation ownership; those bytes are not A001 production work.
9. A separate visible Supervisor review before any disposition.

Full-build, test, CPU-profile, or microbenchmark duration cannot substitute for the required wall-time measurements.

## 12. Rollback condition and attempt-owned bytes

Rollback A001-owned bytes to the accepted baseline if:

- any build, test, outcome, order, graph, persistence, reader, deterministic, failure, or lifecycle invariant changes;
- memory becomes unstable or exceeds the selected linear resource model;
- a valid same-work/same-instrumentation result does not show D001 improvement or shows no corresponding net parent/end-to-end benefit;
- implementation crosses the exact allowed owner boundary.

Attempt-owned bytes are limited to:

- the private index field, initialization, and population hunks in `indexes.go`;
- the two claim-helper lookup/traversal hunks;
- focused test bytes in the two authorized test files.

Rollback must reverse only those hunks. It must not use broad reset/checkout/stash operations or disturb protected worktree changes and carried P1 instrumentation.

## 13. Planner translation checklist for `E2-P2A-A001PLAN1`

- [ ] Bind A001, active parent, full current 17-child queue, selected D001, current timing, calls, files, and no-KEEP streak.
- [ ] State that child total elapsed controls ranking and acceptance.
- [ ] Add the exclusive active-child metric contract and conservation/zero-overlap rule.
- [ ] Add D001 work counters and observer-overhead rules.
- [ ] State that deep instrumentation is active-child-only while coarse totals remain for all children.
- [ ] Bind the proven repeated import-claim/path-normalization cause and overlapping-profile caveat.
- [ ] Bind the one selected all-import claim-index direction.
- [ ] Copy exact production/test owners and preserve-only surfaces.
- [ ] Copy all resolution, graph, output, persistence, reader, determinism, and lifecycle invariants.
- [ ] Require the fresh Coder Anvien pre-edit gate.
- [ ] Require production first and tests second.
- [ ] Record the canonical full-build command.
- [ ] Require same-work D001, parent, and end-to-end measurements.
- [ ] Require allocation/resource evidence and exact output equivalence.
- [ ] Bind rollback conditions and exact A001-owned bytes.
- [ ] Keep Coder locked until Main verifies the Planner refresh as complete and internally consistent.

## 14. Residual NO EVIDENCE items

- Exact helper invocation counts, actual import visits, and bucket-size distribution are not yet measured.
- Exact allocation delta is not yet measured.
- CPU samples are cumulative and overlapping; they cannot predict elapsed savings.
- No A001 implementation, test result, full-build result, fresh remeasurement, equivalence result, disposition, or Supervisor result exists yet.
- Graph `LOW/0` results for the two unexported helper methods are not impact assurance because unresolved method-call edges remain; CRITICAL owner/file blast radius governs.

These uncertainties do not block Planner translation because the selected direction is source-proven, minimal, reversible, and has an exact validation/rollback boundary.

## 15. Output files created

- `reports/system-architect/rp_system-architect_260825_121706_by_gpt-5_child06a_active_child_measurement_and_a001_direction.md`

No production, test, plan, ledger, SPEC, benchmark, measurement, or generated-rule file was edited by this report step.

## 16. Decisions made

1. Use complete child total elapsed time for ranking and performance acceptance.
2. Use exclusive, conserved, zero-overlap breakdown only for the active child.
3. Keep coarse timing for the complete child list and defer deep inactive-child measurement until selection.
4. Retain the separate all-import claim index as the single A001 production direction.
5. Require same-work child, parent, and end-to-end remeasurement plus exact equivalence/resource validation.

## 17. Commit reference

This report is intended for a report-only commit. The resulting commit hash is supplied in the visible handoff after creation; resolve it locally with:

```text
git log -1 --format=%H -- reports/system-architect/rp_system-architect_260825_121706_by_gpt-5_child06a_active_child_measurement_and_a001_direction.md
```

## 18. Next owner

**Main Orchestration** must independently verify this report, then open a separate visible Planner lane to translate it into `E2-P2A-A001PLAN1`. This report does not authorize Coder directly.
