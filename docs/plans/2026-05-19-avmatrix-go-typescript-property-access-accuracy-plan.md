# AVmatrix Go TypeScript Property Access Accuracy Plan

Date: 2026-05-19

Status: planned

Companion files:

- Benchmark ledger: [2026-05-19-avmatrix-go-typescript-property-access-accuracy-benchmark.md](2026-05-19-avmatrix-go-typescript-property-access-accuracy-benchmark.md)
- Evidence ledger: [2026-05-19-avmatrix-go-typescript-property-access-accuracy-evidence.md](2026-05-19-avmatrix-go-typescript-property-access-accuracy-evidence.md)

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

This plan covers a different scope:

- repo: `E:\Website`;
- language/workload: TypeScript-heavy / Next.js-style;
- graph facts: `Property`, `HAS_PROPERTY`, and `ACCESSES`;
- problem area: TypeScript property ownership and member access semantics.

## Goal

Make TypeScript property graph facts useful and auditable for `E:\Website` by distinguishing true orphan properties from false orphan properties, connecting only the properties that have real owners, and resolving member accesses only where semantics are defensible.

The outcome must be measured with final graph facts, not with non-equivalent internal counters from another engine.

## Graph Truth Rule

The graph must reflect the real repository state. Do not create artificial owner links just to increase `HAS_PROPERTY` or `ACCESSES` counts.

If a property, file, object shape, or binding is truly orphaned in the repo, it must remain visibly orphaned in the graph or in the audit output. That orphan status is useful signal: a reader should be able to see that the source has no stable owner/consumer relation instead of seeing a fabricated edge.

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

Baseline source: `reports\benchmark\2026-05-16-120908-avmatrix-go-self-analyze.json`, scenario `websiteTypescriptHeavy`, plus current `E:\Website\.avmatrix\graph.json` after `avmatrix-go` analyze.

`avmatrix-main` must not be rerun for this baseline unless `E:\Website` changes or the embedded benchmark payload is proven invalid.

| Metric | avmatrix-go | avmatrix-main | Note |
|---|---:|---:|---|
| Files scanned | 1,870 | 1,870 | same workload |
| Files parsed | 998 | 998 | same workload |
| Graph nodes | 26,081 | 18,607 | Go emits larger graph |
| Graph relationships | 48,163 | 34,055 | Go emits larger graph |
| Graph `Property` nodes | 5,222 | 3 | Go emits many standalone TS property facts |
| Graph `HAS_PROPERTY` edges | 3 | 3 | current owner-link coverage is minimal |
| Graph `ACCESSES` edges | 3 | 3 | final graph edge count is equal |
| Access resolution counter | 3 | 755 | not graph-equivalent; do not use as quality delta |

Current diagnosis:

- `resolvedAccessesDeltaGoMinusMain=-752` is not a valid final graph quality conclusion.
- The real issue is that many TypeScript `Property` nodes are standalone.
- `resolveAccess` can only emit `ACCESSES` when the property can be found through owner-linked members.
- Some standalone properties may be true orphans and must remain unlinked.
- TypeScript object/type-literal and shape-like properties need explicit ownership semantics before access resolution can improve.

## Accuracy Method

The measured gate for this plan must be repo-owned and repeatable. It must report at least:

- count of `Property` nodes by source category;
- count of `Property` nodes with incoming `HAS_PROPERTY`;
- count of true orphan properties that must remain unlinked;
- count of false orphan properties that should be fixed;
- count of unknown properties that need more evidence before any link is emitted;
- count of invalid/synthetic links rejected or removed;
- count of final graph `ACCESSES` edges;
- sampled precision for newly emitted `HAS_PROPERTY` and `ACCESSES`;
- examples of unresolved member accesses grouped by reason.

Candidate categories:

- class fields;
- interface property signatures;
- type-alias object literal members;
- nested object literal/type literal members;
- destructuring/binding pattern properties;
- runtime object literal keys;
- imported type/member access patterns.

Do not set a numeric `100%` target until Phase 1 defines a defensible denominator. The first hard target is a trustworthy gate and taxonomy.

The eventual target is not `HAS_PROPERTY = Property`. The target is:

- `100%` owner-link coverage for properties classified as false orphans with a defensible owner;
- `0` artificial links for properties classified as true orphans;
- `0` synthetic edges introduced only to improve graph counts;
- a shrinking unknown bucket with examples and reasons recorded.

## Phase 1 - Baseline Gate and Taxonomy

- [ ] [P1-A] Build a tracked TypeScript property/access audit gate. Owner: `internal/graphaccuracy` for reusable audit logic and a command wrapper under `cmd/` if needed. The gate must consume a graph snapshot and emit JSON with property ownership/access metrics.
- [ ] [P1-B] Run the baseline gate on `E:\Website` using the current `avmatrix-go` graph and record the artifact in the benchmark and evidence ledgers.
- [ ] [P1-C] Classify `Property` nodes by source category and write sample examples into the evidence ledger. Required categories: class/interface, type-alias object literal, nested shape, destructuring/binding pattern, runtime object literal, and unknown.
- [ ] [P1-D] Classify standalone properties by orphan status. Required statuses: true orphan, false orphan, unknown, external/library-owned, and intentionally unmodeled.
- [ ] [P1-E] Classify missing and absent edges by graph truth status. Required statuses: real edge missing, true no-edge, unknown no-edge, and invalid synthetic edge risk.
- [ ] [P1-F] Classify unresolved member access candidates by reason. Required reasons: missing receiver type, missing owner link, ambiguous owner, external/library type, unsupported syntax, and false-positive candidate.
- [ ] [P1-G] Define the Phase 2 and Phase 3 measurable targets after the taxonomy is known. The targets must be written as checklist updates, not loose notes.

## Phase 2 - TypeScript Property Ownership

- [ ] [P2-A] Implement owner-link semantics only for defensible false-orphan TypeScript property definitions. Start with interface properties and type-alias object literal properties when the owner is stable in source code.
- [ ] [P2-B] Add focused tests for TypeScript owner-linked properties. Coverage must include interface properties, type alias object members, nested object shape behavior, class fields, and cases that must remain unowned.
- [ ] [P2-C] Run the ownership validation slice and record it. Required evidence: full build before tests, focused tests, CLI/runtime e2e, fresh `E:\Website` analyze, benchmark update, evidence update, and detect-changes or equivalent impact record.

## Phase 3 - TypeScript Member Access Resolution

- [ ] [P3-A] Classify member access samples that should resolve after Phase 2 ownership is available. Include `this.x`, typed parameter access, typed local variable access, constructor assignment, return-value/initializer-derived receiver type, and imported type access.
- [ ] [P3-B] Implement the smallest defensible access-resolution expansion for TypeScript. Prefer cases with explicit receiver type bindings before inferred or heuristic cases.
- [ ] [P3-C] Add focused tests for every access-resolution family closed in [P3-B].
- [ ] [P3-D] Run the access validation slice and record it. Required evidence: full build before tests, focused tests, CLI/runtime e2e, fresh `E:\Website` analyze, benchmark update, evidence update, and impact record.

## Phase 4 - Consumer Impact Checks

- [ ] [P4-A] Verify `context` output includes representative new `HAS_PROPERTY` and `ACCESSES` facts for selected Website symbols.
- [ ] [P4-B] Verify impact analysis behavior on a property owner and on a property consumer. The result must show whether new edges improve affected-symbol discovery without noisy unrelated expansion.
- [ ] [P4-C] Verify graph API/readback preserves the new relationships, including `reason`, `confidence`, `resolutionSource`, and evidence fields.
- [ ] [P4-D] Record consumer-impact evidence and any precision concerns. If noisy edges are found, classify them before final cutover.

## Phase 5 - Final Cutover

- [ ] [P5-A] Run the final TypeScript property/access gate on `E:\Website`.
- [ ] [P5-B] Record final graph size, analyze performance, `Property`, `HAS_PROPERTY`, and `ACCESSES` metrics in the benchmark ledger.
- [ ] [P5-C] Record final evidence: commands, artifacts, focused tests, full build, e2e proof, graph snapshot, gate output, and impact evidence.
- [ ] [P5-D] Close the plan only after benchmark and evidence ledgers agree with the final tracked artifact and all completed targets are satisfied.

## Ledger

| ID | Area | Scope | Target | Benchmark | Evidence | Commit | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| P1-A | Gate | tracked TS property/access audit | repeatable graph fact gate exists | pending | pending | pending | open |
| P1-B | Baseline | Website Go graph | baseline gate recorded | pending | pending | pending | open |
| P1-C | Taxonomy | Property node categories | source categories classified | pending | pending | pending | open |
| P1-D | Taxonomy | standalone properties | true/false/unknown orphan status classified | pending | pending | pending | open |
| P1-E | Taxonomy | missing/absent edges | real missing vs true no-edge classified | pending | pending | pending | open |
| P1-F | Taxonomy | unresolved member accesses | miss reasons classified | pending | pending | pending | open |
| P1-G | Targets | measurable follow-up gates | Phase 2/3 targets defined | pending | pending | pending | open |
| P2-A | Ownership | TS property owner links | defensible `HAS_PROPERTY` expansion | pending | pending | pending | open |
| P2-B | Tests | ownership focused tests | relevant owner cases covered | n/a | pending | pending | open |
| P2-C | Validation | ownership slice | analyze/test/e2e recorded | pending | pending | pending | open |
| P3-A | Access | access sample taxonomy | resolvable families classified | pending | pending | pending | open |
| P3-B | Access | TS member access resolution | defensible `ACCESSES` expansion | pending | pending | pending | open |
| P3-C | Tests | access focused tests | resolved families covered | n/a | pending | pending | open |
| P3-D | Validation | access slice | analyze/test/e2e recorded | pending | pending | pending | open |
| P4-A | Consumer | context | new facts visible in context | n/a | pending | pending | open |
| P4-B | Consumer | impact | affected-symbol behavior checked | n/a | pending | pending | open |
| P4-C | Consumer | graph API/readback | new relationships preserved | n/a | pending | pending | open |
| P4-D | Consumer | precision/noise | concerns classified | pending | pending | pending | open |
| P5-A | Final gate | Website TS property/access | final gate run | pending | pending | pending | open |
| P5-B | Final benchmark | graph facts/performance | final metrics recorded | pending | pending | pending | open |
| P5-C | Final evidence | proof set | final evidence recorded | n/a | pending | pending | open |
| P5-D | Final closure | ledger consistency | plan closed | pending | pending | pending | open |

## Definition Of Done

- The tracked TypeScript property/access gate exists outside `.tmp` and is documented.
- Baseline and final metrics are recorded in the benchmark ledger.
- Evidence contains commands, artifacts, focused tests, full build, e2e proof, graph snapshot, and impact checks.
- Final report uses final graph `ACCESSES`/`HAS_PROPERTY` facts as the quality metric, not non-equivalent internal counters.
- All completed graph expansions have sampled precision evidence.
- True orphans remain visible and are not force-linked.
- True no-edge cases remain without edges and are counted as correct graph truth, not failures.
- Any synthetic edge added only for count inflation is a failure.
- The plan is closed only when the ledger, benchmark, and evidence files agree.
