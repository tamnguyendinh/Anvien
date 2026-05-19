# AVmatrix Go Supported-Language Property Access Accuracy Benchmark

Date: 2026-05-19

Status: open

Companion plan: [2026-05-19-avmatrix-go-supported-language-property-access-accuracy-plan.md](2026-05-19-avmatrix-go-supported-language-property-access-accuracy-plan.md)

Companion evidence: [2026-05-19-avmatrix-go-supported-language-property-access-accuracy-evidence.md](2026-05-19-avmatrix-go-supported-language-property-access-accuracy-evidence.md)

## Benchmark Rules

- Record product/runtime performance, graph fact counts, graph/database throughput, capacity, and inventory counts here.
- Record build/test/e2e timings in evidence, not here, unless the slice changes those systems.
- Use final graph facts for `ACCESSES` and `HAS_PROPERTY` quality metrics.
- Do not compare `avmatrix-go` final graph edge counts directly to `avmatrix-main` internal counters.
- Property/access gate rows must include all graph `Property` nodes, with language breakdowns.

## Existing Workload Baseline

Date: 2026-05-19

Status: recorded from existing benchmark payload

Sources:

- Combined benchmark: `reports\benchmark\2026-05-16-120908-avmatrix-go-self-analyze.json`
- Go-heavy scenario: `scenarios.avmatrixGoSelf`
- TypeScript-heavy scenario: `scenarios.websiteTypescriptHeavy`

No `avmatrix-main` rerun is required for the Website comparison because `E:\Website` is treated as unchanged for this benchmark.

| Scenario | avmatrix-go | avmatrix-main | Go speedup | Go nodes | main nodes | Go relationships | main relationships |
|---|---:|---:|---:|---:|---:|---:|---:|
| `E:\AVmatrix-GO` Go-heavy | 13,588.1 ms | 63,215.2 ms | 4.65x | 19,984 | 17,539 | 47,622 | 39,019 |
| `E:\Website` TypeScript-heavy | 20,763.6 ms | 70,428.1 ms | 3.39x | 26,081 | 18,607 | 48,163 | 34,055 |

Website detail from the existing payload:

| Metric | avmatrix-go | avmatrix-main | Note |
|---|---:|---:|---|
| Total analyze time | 20,763.6 ms | 70,428.1 ms | Go `3.39x` faster |
| Files scanned | 1,870 | 1,870 | same workload |
| Files parsed | 998 | 998 | same workload |
| Graph nodes | 26,081 | 18,607 | Go +7,474 |
| Graph relationships | 48,163 | 34,055 | Go +14,108 |
| Graph `Property` nodes | 5,222 | 3 | Go emits many standalone property nodes |
| Graph `HAS_PROPERTY` edges | 3 | 3 | minimal owner-link coverage |
| Graph `ACCESSES` edges | 3 | 3 | same final graph edge count |
| Go `resolvedAccesses` metric | 3 | n/a | final graph-equivalent for Go pipeline |
| main `scopeResolutionResolvedAccesses` counter | n/a | 755 | internal counter; not graph-equivalent |

Baseline conclusion:

- The `3` vs `755` access number is not a final graph `ACCESSES` delta.
- The final graph `ACCESSES` count is `3` for both engines on Website.
- The actionable benchmark gap is property ownership/access usefulness: `avmatrix-go` has `5,222` Website `Property` nodes but only `3` `HAS_PROPERTY` edges and `3` `ACCESSES` edges.
- This does not mean all `5,222` properties must be linked. The gate must separate true orphan properties from false orphan properties before setting a coverage target.

## Cross-Language Metrics To Record

Every property/access gate benchmark row must include:

| Metric | Meaning |
|---|---|
| `Property` nodes by language | Full graph inventory, not TypeScript-only. |
| Owner-linked properties | Property nodes with incoming `HAS_PROPERTY`. |
| Standalone properties | Property nodes without incoming `HAS_PROPERTY`. |
| True orphan properties | Property facts that correctly have no defensible owner link in the repo. |
| False orphan properties | Property facts that have a real owner but are missing `HAS_PROPERTY`. |
| Unknown property ownership | Property facts that need classification before any edge is emitted. |
| Artificial owner links rejected | Cases where an apparent owner link was intentionally not emitted to preserve graph truth. |
| Final `ACCESSES` edges | Graph relationship count, not an internal resolution counter. |

## Ledger

| Slice | Artifact | Key Metrics | Status |
|---|---|---|---|
| Existing workload baseline | `reports\benchmark\2026-05-16-120908-avmatrix-go-self-analyze.json` | Website: `Property=5,222`, `HAS_PROPERTY=3`, `ACCESSES=3`; Go speedups `4.65x` and `3.39x` | recorded |
| P1 Website baseline gate | `.tmp\p1-property-access-website-baseline-20260519.json` | `Property=5,222`, `HAS_PROPERTY=3`, `ACCESSES=3`, `typescript=5,222`, `invalidHasPropertyEdges=0` | recorded |
| P1 AVmatrix-GO baseline gate | `.tmp\p1-property-access-avmatrix-go-baseline-20260519.json` | `Property=3,046`, `HAS_PROPERTY=2,708`, `ACCESSES=2,724`, `go=2,469`, `typescript=577`, `invalidHasPropertyEdges=0` | recorded |
| P1 language taxonomy | baseline gate artifacts | Website `false_orphan=3,627`; AVmatrix-GO `false_orphan=21`; no invalid synthetic edges | recorded |
| P1 access candidate taxonomy | `.tmp\p1-access-candidates-website-20260519.json`, `.tmp\p1-access-candidates-avmatrix-go-20260519.json` | Website `total=24,542 resolved=3`; AVmatrix-GO `total=21,098 resolved=5,112` | recorded |
| P2 ownership validation | pending | false-orphan reduction and `HAS_PROPERTY` precision | open |
| P3 access validation | pending | `ACCESSES` expansion and precision | open |
| P5 final gate | pending | final workload matrix | open |

## P1 Cross-Language Baseline Gate

Date: 2026-05-19

Artifacts:

- Website graph snapshot: `.tmp\p1-property-access-website-go-graph-20260519.json`
- Website gate output: `.tmp\p1-property-access-website-baseline-20260519.json`
- AVmatrix-GO graph snapshot: `.tmp\p1-property-access-avmatrix-go-graph-20260519.json`
- AVmatrix-GO gate output: `.tmp\p1-property-access-avmatrix-go-baseline-20260519.json`

Commands:

```powershell
go run ./cmd/property-access-audit -repo E:\Website -graph .tmp\p1-property-access-website-go-graph-20260519.json -out .tmp\p1-property-access-website-baseline-20260519.json -max-examples 20
go run ./cmd/property-access-audit -repo E:\AVmatrix-GO -graph .tmp\p1-property-access-avmatrix-go-graph-20260519.json -out .tmp\p1-property-access-avmatrix-go-baseline-20260519.json -max-examples 20
```

Totals:

| Workload | `Property` | Owner-linked | Standalone | `HAS_PROPERTY` | `ACCESSES` | Invalid `HAS_PROPERTY` |
|---|---:|---:|---:|---:|---:|---:|
| `E:\Website` | 5,222 | 3 | 5,219 | 3 | 3 | 0 |
| `E:\AVmatrix-GO` | 3,046 | 2,708 | 338 | 2,708 | 2,724 | 0 |

Language breakdown:

| Workload | Language | `Property` | Owner-linked | Standalone |
|---|---|---:|---:|---:|
| `E:\Website` | TypeScript | 5,222 | 3 | 5,219 |
| `E:\AVmatrix-GO` | Go | 2,469 | 2,263 | 206 |
| `E:\AVmatrix-GO` | TypeScript | 577 | 445 | 132 |

Category taxonomy:

| Workload | Category | Count |
|---|---|---:|
| `E:\Website` | `tsjs_type_alias_object_literal_member` | 3,627 |
| `E:\Website` | `tsjs_typed_shape_or_binding_property` | 1,094 |
| `E:\Website` | `tsjs_runtime_object_literal_key` | 285 |
| `E:\Website` | `tsjs_destructuring_or_binding_pattern` | 145 |
| `E:\Website` | `tsjs_unclassified_property` | 68 |
| `E:\Website` | `tsjs_class_field` | 3 |
| `E:\AVmatrix-GO` | `go_owner_linked_struct` | 2,263 |
| `E:\AVmatrix-GO` | `tsjs_interface_property_signature` | 449 |
| `E:\AVmatrix-GO` | `go_typed_property_without_owner` | 206 |
| `E:\AVmatrix-GO` | `tsjs_destructuring_or_binding_pattern` | 58 |
| `E:\AVmatrix-GO` | `tsjs_typed_shape_or_binding_property` | 26 |
| `E:\AVmatrix-GO` | `tsjs_runtime_object_literal_key` | 24 |
| `E:\AVmatrix-GO` | `tsjs_type_alias_object_literal_member` | 9 |
| `E:\AVmatrix-GO` | `tsjs_class_field` | 8 |
| `E:\AVmatrix-GO` | `tsjs_unclassified_property` | 3 |

Orphan taxonomy:

| Workload | `owner_linked` | `false_orphan` | `true_orphan` | `unknown` | `external_library_owned` | `intentionally_unmodeled` |
|---|---:|---:|---:|---:|---:|---:|
| `E:\Website` | 3 | 3,627 | 430 | 1,162 | 0 | 0 |
| `E:\AVmatrix-GO` | 2,708 | 21 | 82 | 235 | 0 | 0 |

Graph truth taxonomy:

| Workload | `edge_present` | `real_edge_missing` | `true_no_edge` | `unknown_no_edge` | `invalid_synthetic_edge_risk` |
|---|---:|---:|---:|---:|---:|
| `E:\Website` | 3 | 3,627 | 430 | 1,162 | 0 |
| `E:\AVmatrix-GO` | 2,708 | 21 | 82 | 235 | 0 |

## P1 Access Candidate Taxonomy

Date: 2026-05-19

Artifacts:

- Website access candidate output: `.tmp\p1-access-candidates-website-20260519.json`
- AVmatrix-GO access candidate output: `.tmp\p1-access-candidates-avmatrix-go-20260519.json`

Commands:

```powershell
go run ./cmd/access-candidate-audit -repo E:\Website -out .tmp\p1-access-candidates-website-20260519.json -max-examples 20
go run ./cmd/access-candidate-audit -repo E:\AVmatrix-GO -out .tmp\p1-access-candidates-avmatrix-go-20260519.json -max-examples 20
```

Totals:

| Workload | Analyze runtime | Access candidates | Resolved | Unresolved | Final graph resolved accesses | Resolution unresolved references |
|---|---:|---:|---:|---:|---:|---:|
| `E:\Website` | 7,365 ms | 24,542 | 3 | 24,539 | 3 | 64,274 |
| `E:\AVmatrix-GO` | 5,862 ms | 21,098 | 5,112 | 15,986 | 5,112 | 50,282 |

Language breakdown:

| Workload | Language | Access candidates | Resolved | Unresolved |
|---|---|---:|---:|---:|
| `E:\Website` | JavaScript | 231 | 0 | 231 |
| `E:\Website` | TypeScript | 24,311 | 3 | 24,308 |
| `E:\AVmatrix-GO` | Go | 19,059 | 4,975 | 14,084 |
| `E:\AVmatrix-GO` | TypeScript | 2,039 | 137 | 1,902 |

Reason taxonomy:

| Workload | `resolved` | `missing_receiver_type` | `external_library_type` | `unsupported_syntax` | `missing_caller` | `missing_owner_link` | `false_positive_candidate` | `ambiguous_owner` |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| `E:\Website` | 3 | 13,512 | 9,660 | 1,313 | 53 | 1 | 0 | 0 |
| `E:\AVmatrix-GO` | 5,112 | 11,418 | 3,619 | 910 | 14 | 10 | 15 | 0 |
