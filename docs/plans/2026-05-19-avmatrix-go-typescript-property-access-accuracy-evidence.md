# AVmatrix Go TypeScript Property Access Accuracy Evidence

Date: 2026-05-19

Status: open

Companion plan: [2026-05-19-avmatrix-go-typescript-property-access-accuracy-plan.md](2026-05-19-avmatrix-go-typescript-property-access-accuracy-plan.md)

Companion benchmark: [2026-05-19-avmatrix-go-typescript-property-access-accuracy-benchmark.md](2026-05-19-avmatrix-go-typescript-property-access-accuracy-benchmark.md)

## Evidence Rules

- Record commands, artifacts, focused tests, full build, e2e proof, graph snapshots, and impact checks here.
- Keep benchmark measurements in the benchmark ledger.
- For doc-only commits, do not run AVmatrix solely for the commit.
- For implementation slices, use AVmatrix for codebase analysis and impact checks.

## Plan Creation Evidence

- Created a new plan because the 2026-05-16 graph accuracy plan is scoped to Go-local measured gates on `E:\AVmatrix-GO`.
- New scope is TypeScript-heavy property/access graph accuracy on `E:\Website`.
- The new plan does not reopen the completed 2026-05-16 plan.

## Baseline Audit Evidence

Existing benchmark reviewed:

- `reports\benchmark\2026-05-16-120908-avmatrix-go-self-analyze.json`
- `reports\benchmark\2026-05-16-120908-avmatrix-go-self-analyze.md`

Relevant result:

```text
Website final graph ACCESSES edges:
avmatrix-go:   3
avmatrix-main: 3

Website internal access counters:
avmatrix-go resolvedAccesses:              3
avmatrix-main scopeResolutionResolvedAccesses: 755
```

Conclusion:

- The `3` vs `755` access number is not a final graph-edge comparison.
- `avmatrix-main` value `755` is an internal counter and must not be used as the final graph `ACCESSES` target.
- The real follow-up is TypeScript property ownership and member access semantics.
- True orphan properties must remain visible as true orphans; the plan must not hide real repo structure problems by force-linking graph edges.

Current graph symptom:

```text
E:\Website avmatrix-go graph:
Property nodes:      5,222
HAS_PROPERTY edges:      3
ACCESSES edges:          3
```

Code pointers for the initial hypothesis:

- `internal/providers/tsjs/definitions.go:48` emits TypeScript `public_field_definition` and `property_signature` as `Property`.
- `internal/resolution/indexes.go:147` indexes owner members only when `DefinitionFact.OwnerID` is non-empty.
- `internal/resolution/resolve.go:200` resolves `AccessFact` through `resolveMember(... propertyLabels())`.

Initial hypothesis:

- Many TypeScript properties are emitted without a stable owner link.
- Unowned properties do not enter `ownerMembers`.
- Member access resolution cannot target unowned properties.
- Some unowned properties may be correct true orphans and should stay unlinked.
- The first implementation task must distinguish false orphans from true orphans before adding edges.

## Evidence Ledger

| Slice | Evidence | Result | Status |
|---|---|---|---|
| Plan creation | new plan/benchmark/evidence files | pending commit | open |
| P1 baseline gate | pending | pending | open |
| P1 taxonomy | pending | pending | open |
| P1 orphan truth classification | pending | pending | open |
| P2 ownership implementation | pending | pending | open |
| P2 validation | pending | pending | open |
| P3 access implementation | pending | pending | open |
| P3 validation | pending | pending | open |
| P4 consumer checks | pending | pending | open |
| P5 final evidence | pending | pending | open |
