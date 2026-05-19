# AVmatrix Go TypeScript Property Access Accuracy Benchmark

Date: 2026-05-19

Status: open

Companion plan: [2026-05-19-avmatrix-go-typescript-property-access-accuracy-plan.md](2026-05-19-avmatrix-go-typescript-property-access-accuracy-plan.md)

Companion evidence: [2026-05-19-avmatrix-go-typescript-property-access-accuracy-evidence.md](2026-05-19-avmatrix-go-typescript-property-access-accuracy-evidence.md)

## Benchmark Rules

- Record product/runtime performance, graph fact counts, graph/database throughput, capacity, and inventory counts here.
- Record build/test/e2e timings in evidence, not here, unless the slice changes those systems.
- Use final graph facts for `ACCESSES` and `HAS_PROPERTY` quality metrics.
- Do not compare `avmatrix-go` final graph edge counts directly to `avmatrix-main` internal counters.

## Baseline - Website TypeScript Property/Access Audit

Date: 2026-05-19

Status: recorded from existing benchmark payload and current Go graph

Sources:

- Combined benchmark: `reports\benchmark\2026-05-16-120908-avmatrix-go-self-analyze.json`
- Scenario: `scenarios.websiteTypescriptHeavy`
- Current Go graph snapshot source: `E:\Website\.avmatrix\graph.json`

No `avmatrix-main` rerun was required because `E:\Website` is treated as unchanged for this benchmark.

| Metric | avmatrix-go | avmatrix-main | Note |
|---|---:|---:|---|
| Total analyze time | 20,763.6 ms | 70,428.1 ms | Go `3.39x` faster |
| Files scanned | 1,870 | 1,870 | same workload |
| Files parsed | 998 | 998 | same workload |
| Graph nodes | 26,081 | 18,607 | Go +7,474 |
| Graph relationships | 48,163 | 34,055 | Go +14,108 |
| Graph `Property` nodes | 5,222 | 3 | Go emits many standalone TS property nodes |
| Graph `HAS_PROPERTY` edges | 3 | 3 | minimal owner-link coverage |
| Graph `ACCESSES` edges | 3 | 3 | same final graph edge count |
| Go `resolvedAccesses` metric | 3 | n/a | final graph-equivalent for Go pipeline |
| main `scopeResolutionResolvedAccesses` counter | n/a | 755 | internal counter; not graph-equivalent |

Baseline conclusion:

- The `3` vs `755` access number is not a final graph `ACCESSES` delta.
- The final graph `ACCESSES` count is `3` for both engines.
- The actionable benchmark gap is TypeScript property ownership/access usefulness: `avmatrix-go` has `5,222` `Property` nodes but only `3` `HAS_PROPERTY` edges and `3` `ACCESSES` edges.
- This does not mean all `5,222` properties must be linked. The next benchmark must separate true orphan properties from false orphan properties before setting a coverage target.

## Orphan Classification Metrics To Add

Future benchmark rows must include:

| Metric | Meaning |
|---|---|
| True orphan properties | Property facts that correctly have no defensible owner link in the repo. |
| False orphan properties | Property facts that have a real owner but are missing `HAS_PROPERTY`. |
| Unknown property ownership | Property facts that need classification before any edge is emitted. |
| Artificial owner links rejected | Cases where an apparent owner link was intentionally not emitted to preserve graph truth. |

## Ledger

| Slice | Artifact | Key Metrics | Status |
|---|---|---|---|
| Baseline | `reports\benchmark\2026-05-16-120908-avmatrix-go-self-analyze.json` | Website: `Property=5,222`, `HAS_PROPERTY=3`, `ACCESSES=3` | recorded |
| P1 baseline gate | pending | pending | open |
| P1 orphan taxonomy | pending | true/false/unknown orphan counts | open |
| P2 ownership validation | pending | pending | open |
| P3 access validation | pending | pending | open |
| P5 final gate | pending | pending | open |
