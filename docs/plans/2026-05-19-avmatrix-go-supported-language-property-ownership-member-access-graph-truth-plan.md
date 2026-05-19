# AVmatrix Go Supported-Language Property Ownership And Member Access Graph Truth Plan

Date: 2026-05-19

Status: complete

Companion files:

- Benchmark ledger: [2026-05-19-avmatrix-go-supported-language-property-ownership-member-access-graph-truth-benchmark.md](2026-05-19-avmatrix-go-supported-language-property-ownership-member-access-graph-truth-benchmark.md)
- Evidence ledger: [2026-05-19-avmatrix-go-supported-language-property-ownership-member-access-graph-truth-evidence.md](2026-05-19-avmatrix-go-supported-language-property-ownership-member-access-graph-truth-evidence.md)

## Rules

1. Use AVmatrix for codebase analysis and impact checks while working on implementation slices in this plan.
2. As each task is completed, update the corresponding checklist item immediately.
3. Run a full build before testing; the test suite must include an e2e test.
4. Record benchmark results as each benchmarkable task is completed. Benchmarkable means measured product/runtime performance, capacity, package/startup size, graph/DB throughput, or conversion-inventory counts; build/test/e2e timings are validation evidence unless the slice changes those systems.
5. Record evidence as each evidenced task is completed.
6. For doc-only commits, do not use AVmatrix.
7. After each completed implementation slice, commit the work, then continue until the full plan is complete.

## Scope Boundary

This is a new follow-up plan. It does not reopen [2026-05-16-avmatrix-go-graph-accuracy-100-plan.md](2026-05-16-avmatrix-go-graph-accuracy-100-plan.md).

The 2026-05-16 plan completed its intended scope: Go-local measured gates on `E:\AVmatrix-GO`.

Canonical file set for this plan:

- plan: `docs\plans\2026-05-19-avmatrix-go-supported-language-property-ownership-member-access-graph-truth-plan.md`
- benchmark: `docs\plans\2026-05-19-avmatrix-go-supported-language-property-ownership-member-access-graph-truth-benchmark.md`
- evidence: `docs\plans\2026-05-19-avmatrix-go-supported-language-property-ownership-member-access-graph-truth-evidence.md`

This plan covers a broader product contract:

- repos/workloads: `E:\AVmatrix-GO`, `E:\Website`, and provider fixture graphs for every supported language family;
- graph facts: `Property`, `HAS_PROPERTY`, and `ACCESSES`;
- language scope: every language that AVmatrix-Go can scan/parse/extract and every language that appears in graph `Property` nodes;
- problem area: property ownership and member access semantics across the supported-language graph, not TypeScript only.

`E:\Website` remains important because it is a TypeScript-heavy benchmark workload. It is not the whole problem.

## Goal

Make property graph facts useful and auditable across all AVmatrix-Go supported languages by distinguishing true orphan properties from false orphan properties, connecting only the properties that have real owners, and resolving member accesses only where semantics are defensible.

The outcome must be measured with final graph facts, not with non-equivalent internal counters from another engine.

## Graph Truth Rule

The graph must reflect the real repository state. Do not create artificial owner links just to increase `HAS_PROPERTY` or `ACCESSES` counts.

If a property, file, object shape, binding, or language construct is truly orphaned in the repo, it must remain visibly orphaned in the graph or in the audit output. That orphan status is useful signal: a reader should be able to see that the source has no stable owner/consumer relation instead of seeing a fabricated edge.

Graph meaning for this plan:

- an edge means AVmatrix has evidence that a real relationship exists in the repo;
- no edge means AVmatrix has no defensible relationship for that pair under the current source model;
- a node with no owner/consumer edges is still meaningful because it exposes the actual disconnected state;
- a missing real edge is a false negative and should be fixed;
- a fabricated edge is worse than a missing edge because it hides the real repository state.

For example, if file or symbol `A` really refers to `B`, `C`, and `D`, the graph should show those edges. If `A` does not really refer to `E`, the graph must not create `A -> E` to improve counts. If `A` has no real links, the graph must leave `A` disconnected so readers can see that condition.

This plan fixes only false orphans:

- a false orphan is a property that has a real static owner in source code, but AVmatrix-Go failed to connect it;
- a true orphan is a property-like fact with no defensible owner in the source model, and it must not be force-linked;
- unknown cases must stay classified as unknown until evidence proves they are safe to connect.

## Baseline

Existing benchmark source: `reports\benchmark\2026-05-16-120908-avmatrix-go-self-analyze.json`.

Known workload-level result:

| Scenario | avmatrix-go | avmatrix-main | Go speedup | Note |
|---|---:|---:|---:|---|
| `E:\AVmatrix-GO` Go-heavy | 13,588.1 ms | 63,215.2 ms | 4.65x | useful Go-heavy workload |
| `E:\Website` TypeScript-heavy | 20,763.6 ms | 70,428.1 ms | 3.39x | useful TypeScript-heavy workload |

Known `E:\Website` graph symptom from the existing benchmark and current Go graph:

| Metric | avmatrix-go | avmatrix-main | Note |
|---|---:|---:|---|
| Files scanned | 1,870 | 1,870 | same workload |
| Files parsed | 998 | 998 | same workload |
| Graph nodes | 26,081 | 18,607 | Go emits larger graph |
| Graph relationships | 48,163 | 34,055 | Go emits larger graph |
| Graph `Property` nodes | 5,222 | 3 | Go emits many standalone property facts |
| Graph `HAS_PROPERTY` edges | 3 | 3 | current owner-link coverage is minimal |
| Graph `ACCESSES` edges | 3 | 3 | final graph edge count is equal |
| Access resolution counter | 3 | 755 | not graph-equivalent; do not use as quality delta |

Current diagnosis:

- `resolvedAccessesDeltaGoMinusMain=-752` is not a valid final graph quality conclusion.
- The real issue is property ownership/access usefulness in the final graph.
- `resolveAccess` can only emit `ACCESSES` when the property can be found through owner-linked members.
- Some standalone properties may be true orphans and must remain unlinked.
- Owner semantics must be audited by language, because each provider has different property constructs.

## Accuracy Method

The measured gate for this plan must be repo-owned and repeatable. It must report at least:

- count of `Property` nodes across the full graph;
- count of `Property` nodes by language;
- count of `Property` nodes by source category;
- count of `Property` nodes with incoming `HAS_PROPERTY`;
- count of true orphan properties that must remain unlinked;
- count of false orphan properties that should be fixed;
- count of unknown properties that need more evidence before any link is emitted;
- count of invalid/synthetic links rejected or removed;
- count of final graph `ACCESSES` edges;
- sampled precision for newly emitted `HAS_PROPERTY` and `ACCESSES`;
- examples of unresolved member accesses grouped by reason.

Candidate language families:

- TS/JS/SFC: JavaScript, TypeScript, Vue, Svelte, Astro;
- static OO/member languages: Java, Kotlin, C#, PHP, Swift, Dart;
- systems/member languages: Go, C, C++, Rust;
- dynamic languages: Python, Ruby;
- enterprise/legacy: COBOL and JCL-related graph facts where emitted.

Candidate property categories:

- class/struct/interface fields;
- interface property signatures;
- type-alias object literal members;
- nested object literal/type literal members;
- destructuring/binding pattern properties;
- runtime object literal keys;
- imported type/member access patterns;
- qualified member properties without owner links;
- language-specific unknown properties.

Do not set a numeric `100%` target until Phase 1 defines a defensible denominator per language family. The first hard target is a trustworthy cross-language gate and taxonomy.

The eventual target is not `HAS_PROPERTY = Property`. The target is:

- `100%` owner-link coverage for properties classified as false orphans with a defensible owner;
- `0` artificial links for properties classified as true orphans;
- `0` synthetic edges introduced only to improve graph counts;
- a shrinking unknown bucket with examples and reasons recorded by language.

## Phase 1 - Cross-Language Baseline Gate and Taxonomy

- [x] [P1-A] Build a tracked cross-language property/access audit gate. Owner: `internal/graphaccuracy` with command wrapper `cmd/property-access-audit`. The gate consumes a graph snapshot, includes every graph `Property` node, and emits JSON with language, ownership, access, category, orphan, and graph-truth metrics.
- [x] [P1-B] Run the baseline gate on `E:\Website` using explicit graph snapshot artifact `.tmp\p1-property-access-website-go-graph-20260519.json`; baseline output `.tmp\p1-property-access-website-baseline-20260519.json` is recorded in the benchmark and evidence ledgers.
- [x] [P1-C] Run the baseline gate on `E:\AVmatrix-GO` using explicit graph snapshot artifact `.tmp\p1-property-access-avmatrix-go-graph-20260519.json`; baseline output `.tmp\p1-property-access-avmatrix-go-baseline-20260519.json` is recorded in the benchmark and evidence ledgers.
- [x] [P1-D] Classify `Property` nodes by language and source category. Website baseline: `typescript=5,222`. AVmatrix-GO baseline: `go=2,469`, `typescript=577`. Required category families are represented in the gate output.
- [x] [P1-E] Classify standalone properties by orphan status. Website: `false_orphan=3,627`, `true_orphan=430`, `unknown=1,162`, `owner_linked=3`, `external_library_owned=0`, `intentionally_unmodeled=0`. AVmatrix-GO: `false_orphan=21`, `true_orphan=82`, `unknown=235`, `owner_linked=2,708`, `external_library_owned=0`, `intentionally_unmodeled=0`.
- [x] [P1-F] Classify missing and absent edges by graph truth status. Website: `real_edge_missing=3,627`, `true_no_edge=430`, `unknown_no_edge=1,162`, `edge_present=3`, `invalid_synthetic_edge_risk=0`. AVmatrix-GO: `real_edge_missing=21`, `true_no_edge=82`, `unknown_no_edge=235`, `edge_present=2,708`, `invalid_synthetic_edge_risk=0`.
- [x] [P1-G] Classify unresolved member access candidates by reason. Website: `total=24,542`, `resolved=3`, `missing_receiver_type=13,512`, `external_library_type=9,660`, `unsupported_syntax=1,313`, `missing_caller=53`, `missing_owner_link=1`, `ambiguous_owner=0`, `false_positive_candidate=0`. AVmatrix-GO: `total=21,098`, `resolved=5,112`, `missing_receiver_type=11,418`, `external_library_type=3,619`, `unsupported_syntax=910`, `missing_caller=14`, `missing_owner_link=10`, `false_positive_candidate=15`, `ambiguous_owner=0`.
- [x] [P1-H] Define the Phase 2 and Phase 3 measurable targets after the taxonomy is known. Targets are written below and must be updated after each validation gate.

## Phase 2 - Cross-Language Property Ownership

- [x] [P2-A] Implement owner-link semantics for the first large false-orphan cluster: TS/JS/SFC static shape properties. Result: direct TypeScript type-alias object members now emit defensible `TypeAlias -> Property` ownership, while nested shape properties remain unowned until nested ownership is modeled. Website `HAS_PROPERTY` increased from `3` to `5,129`; Website `real_edge_missing` decreased from `3,627` to `392`; invalid synthetic edges remain `0`.
- [x] [P2-B] Add focused tests for every language/provider family changed in [P2-A]. Result: TS/JS provider coverage now asserts direct type-alias members are owner-linked and nested shape members remain unowned.
- [x] [P2-C] Run ownership validation and record it. Result: full build passed before tests, focused tests passed, CLI/runtime e2e passed on `E:\Website` and `E:\AVmatrix-GO`, fresh graph snapshots and benchmark/evidence ledgers were updated, and AVmatrix impact was recorded.
- [x] [P2-D] Refine Go unknown property ownership before linking. Result: AVmatrix-GO `go_typed_property_without_owner=206` was reclassified as `go_anonymous_struct_field=206`, all `true_orphan`/`true_no_edge`, with no new `HAS_PROPERTY` emitted. Go ownership now has `owner_linked=2,302`, `true_orphan=206`, `unknown=0`.
- [x] [P2-E] Repeat large ownership clusters until all false-orphan categories with defensible owners are either fixed or explicitly deferred with evidence. Result: nested TS/JS object-shape members now link to their parent `Property`; inline TS/JS anonymous type literals are classified as true no-edge. Website `real_edge_missing=0`; AVmatrix-GO `real_edge_missing=0`; invalid synthetic edges remain `0`.

## Phase 3 - Cross-Language Member Access Resolution

- [x] [P3-A] Classify member access samples that should resolve after Phase 2 ownership is available. Result: post-Phase-2 access audit recorded Website `total=24,542`, `resolved=3`, `missing_receiver_type=13,512`, `missing_owner_link=1`; AVmatrix-GO `total=21,132`, `resolved=5,128`, `missing_receiver_type=11,436`, `missing_owner_link=10`. The next large target remains receiver type binding, not owner-link expansion.
- [x] [P3-B] Implement the first large defensible access-resolution cluster around explicit receiver type bindings. Result: awaited TS/JS local call-return bindings now unwrap `Promise<T>` when the source expression is `await`, resolver member owners include `TypeAlias` for property/member lookup, and call-return enrichment no longer overwrites provider-derived return bindings. Website final graph `ACCESSES` increased from `3` to `2,769`; Website access candidates resolved increased from `3` to `4,978`; Website `missing_receiver_type` decreased from `13,512` to `12,209`.
- [x] [P3-C] Add focused tests for every access-resolution family closed in [P3-B]. Result: provider test covers awaited `Promise<T>` local binding, resolution test covers nested `result.model.invoices` `ACCESSES`, and access audit test covers `TypeAlias` member-owner resolution.
- [x] [P3-D] Run access validation and record it. Result: full build passed before tests, focused tests passed, analyze/property-gate/access-candidate e2e ran on `E:\Website` and `E:\AVmatrix-GO`, fresh graph snapshots and benchmark/evidence ledgers were updated, and AVmatrix impact was recorded.
- [x] [P3-E] Close or reclassify the `missing_owner_link` access bucket after receiver-type expansion. Result: Website `missing_owner_link=768 -> 0`; AVmatrix-GO `missing_owner_link=10 -> 0`. The slice does not add owner links; it rejects cross-language global owner collisions and reclassifies same-name standalone-property guesses as false positives unless a real owner-member relation exists.
- [x] [P3-F] Repeat large access clusters until all target families are fixed or explicitly deferred with evidence. Result: imported workspace member accesses now resolve to real `ACCESSES` edges, unresolved imported receivers are classified as external/no-workspace-model instead of missing receiver type, and LadybugDB schema accepts the real access source/target pairs revealed by the self-repo graph. Website final `ACCESSES=2,770`; AVmatrix-GO final `ACCESSES=5,018`; `missing_owner_link=0` and invalid owner edges remain `0`. Remaining buckets are explicitly deferred with evidence: untyped receiver/dataflow inference, external library or unresolved import target modeling, and unsupported computed/call/index receiver syntax.

## Phase 4 - Consumer Impact Checks

- [x] [P4-A] Verify `context` output includes representative new `HAS_PROPERTY` and `ACCESSES` facts for selected symbols in multiple language families. Result: `context parseFiles` shows imported `ACCESSES` to `parser.ErrUnsupportedLanguage`; `context RawSettingsShape` shows TypeScript `TypeAlias -> Property` `HAS_PROPERTY`; `context LLMSettings.intelligentClustering` shows both owner and consumer.
- [x] [P4-B] Verify impact analysis behavior on property owners and property consumers. Result: impact default traversal now includes `HAS_PROPERTY` and `ACCESSES`; `ErrUnsupportedLanguage` impact reaches `parseFiles`, and `LLMSettings.intelligentClustering` impact reaches both `updateLocalRuntimeProviderSettings` and owner `LLMSettings`.
- [x] [P4-C] Verify graph API/readback preserves the new relationships, including `reason`, `confidence`, `resolutionSource`, and evidence fields. Result: `cypher` readback for `parseFiles -> ErrUnsupportedLanguage` returns `type=ACCESSES`, `confidence=0.9`, `reason=read`, `resolutionSource=scope-resolution`, and `import-binding` evidence.
- [x] [P4-D] Record consumer-impact evidence and any precision concerns. Result: initial P4 check found that `ACCESSES`/`HAS_PROPERTY` were allowed but omitted from default impact traversal; fixed in consumer layer and covered by regression test. No noisy unrelated expansion was observed in the sampled checks.

## Phase 5 - Final Cutover

- [x] [P5-A] Run the final property/access gate on `E:\AVmatrix-GO`, `E:\Website`, and the supported-language fixture matrix. Result: final gates ran on both workloads; supported-language fixture matrix is covered by `go test ./internal/providers/...`.
- [x] [P5-B] Record final graph size, analyze performance, `Property`, `HAS_PROPERTY`, and `ACCESSES` metrics in the benchmark ledger. Result: Website `16,436.3 ms`, `Property=7,097`, `HAS_PROPERTY=5,922`, `ACCESSES=2,770`; AVmatrix-GO `15,689.4 ms`, `Property=3,096`, `HAS_PROPERTY=2,769`, `ACCESSES=5,024`.
- [x] [P5-C] Record final evidence: commands, artifacts, focused tests, full build, e2e proof, graph snapshots, gate outputs, and impact evidence. Result: evidence ledger records full build, provider fixture tests, final analyze, graph snapshots, property/access gate outputs, and access-candidate outputs.
- [x] [P5-D] Close the plan only after benchmark and evidence ledgers agree with the final tracked artifacts and all completed targets are satisfied. Result: plan closed with `real_edge_missing=0`, `missing_owner_link=0`, invalid owner edges `0`, and remaining no-edge buckets explicitly classified.

## Ledger

| ID | Area | Scope | Target | Benchmark | Evidence | Commit | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| P1-A | Gate | cross-language property/access audit | repeatable graph fact gate exists | n/a | recorded | `0972c3d` | done |
| P1-B | Baseline | Website graph snapshot | cross-language baseline gate artifact recorded | recorded | recorded | `0972c3d` | done |
| P1-C | Baseline | AVmatrix-GO graph snapshot | cross-language baseline gate artifact recorded | recorded | recorded | `0972c3d` | done |
| P1-D | Taxonomy | property language/categories | source categories classified | recorded | recorded | `0972c3d` | done |
| P1-E | Taxonomy | standalone properties | true/false/unknown orphan status classified | recorded | recorded | `0972c3d` | done |
| P1-F | Taxonomy | missing/absent edges | real missing vs true no-edge classified | recorded | recorded | `0972c3d` | done |
| P1-G | Taxonomy | unresolved member accesses | miss reasons classified | recorded | recorded | `73b1cf8` | done |
| P1-H | Targets | measurable follow-up gates | Phase 2/3 targets defined | recorded | recorded | `73b1cf8` | done |
| P2-A | Ownership | TS/JS/SFC static shape cluster | defensible `TypeAlias -> Property` expansion | recorded | recorded | `b91c2cd` | done |
| P2-B | Tests | TS/JS provider ownership focused tests | direct owner and nested no-owner covered | n/a | recorded | `b91c2cd` | done |
| P2-C | Validation | ownership slice | analyze/test/e2e recorded | recorded | recorded | `b91c2cd` | done |
| P2-D | Ownership | Go anonymous struct property ownership | classify before linking | recorded | recorded | `0367bc1` | done |
| P2-E | Ownership | remaining false-orphan ownership clusters | fixed or deferred with evidence | recorded | recorded | `de5f4d8` | done |
| P3-A | Access | post-ownership access sample taxonomy | resolvable families classified | recorded | recorded | `42ba98b` | done |
| P3-B | Access | awaited TS/JS call-return and TypeAlias member owners | defensible `ACCESSES` expansion | recorded | recorded | `08649e6` | done |
| P3-C | Tests | access focused tests | resolved families covered | n/a | recorded | `08649e6` | done |
| P3-D | Validation | access slice | analyze/test/e2e recorded | recorded | recorded | `08649e6` | done |
| P3-E | Access | post-receiver missing-owner-link bucket | close or reclassify bucket | recorded | recorded | `a908b2d` | done |
| P3-F | Access | imported member receivers and remaining clusters | fixed or deferred with evidence | recorded | recorded | `9b58dea` | done |
| P4-A | Consumer | context | new facts visible in context | n/a | recorded | `ba2a0da` | done |
| P4-B | Consumer | impact | affected-symbol behavior checked | n/a | recorded | `ba2a0da` | done |
| P4-C | Consumer | graph API/readback | new relationships preserved | n/a | recorded | `ba2a0da` | done |
| P4-D | Consumer | precision/noise | concerns classified | n/a | recorded | `ba2a0da` | done |
| P5-A | Final gate | workload matrix | final gate run | recorded | recorded | `d45c4ca` | done |
| P5-B | Final benchmark | graph facts/performance | final metrics recorded | recorded | recorded | `d45c4ca` | done |
| P5-C | Final evidence | proof set | final evidence recorded | n/a | recorded | `d45c4ca` | done |
| P5-D | Final closure | ledger consistency | plan closed | recorded | recorded | `d45c4ca` | done |

## Definition Of Done

- The tracked property/access gate exists outside `.tmp` and is documented.
- The gate includes every graph `Property` node, not only TypeScript.
- Baseline and final metrics are recorded in the benchmark ledger.
- Evidence contains commands, artifacts, focused tests, full build, e2e proof, graph snapshots, and impact checks.
- Final report uses final graph `ACCESSES`/`HAS_PROPERTY` facts as the quality metric, not non-equivalent internal counters.
- All completed graph expansions have sampled precision evidence.
- True orphans remain visible and are not force-linked.
- True no-edge cases remain without edges and are counted as correct graph truth, not failures.
- Any synthetic edge added only for count inflation is a failure.
- The plan is closed only when the ledger, benchmark, and evidence files agree.
