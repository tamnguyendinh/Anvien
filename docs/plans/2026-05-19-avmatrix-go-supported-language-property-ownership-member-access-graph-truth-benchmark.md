# AVmatrix Go Supported-Language Property Ownership And Member Access Graph Truth Benchmark

Date: 2026-05-19

Status: open

Companion plan: [2026-05-19-avmatrix-go-supported-language-property-ownership-member-access-graph-truth-plan.md](2026-05-19-avmatrix-go-supported-language-property-ownership-member-access-graph-truth-plan.md)

Companion evidence: [2026-05-19-avmatrix-go-supported-language-property-ownership-member-access-graph-truth-evidence.md](2026-05-19-avmatrix-go-supported-language-property-ownership-member-access-graph-truth-evidence.md)

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
| P2 ownership validation | `.tmp\p2-property-access-website-20260519.json`, `.tmp\p2-property-access-avmatrix-go-20260519.json` | Website `HAS_PROPERTY=5,129`, `real_edge_missing=392`, invalid edges `0`; AVmatrix-GO `HAS_PROPERTY=2,762`, `real_edge_missing=12`, invalid edges `0` | recorded |
| P2-D Go anonymous struct classification | `.tmp\p2d-property-access-avmatrix-go-20260519.json` | Go `go_anonymous_struct_field=206`, Go `true_no_edge=206`, Go `unknown=0`; no new `HAS_PROPERTY` emitted | recorded |
| P2-E remaining ownership clusters | `.tmp\p2e-property-access-website-20260519.json`, `.tmp\p2e-property-access-avmatrix-go-20260519.json` | Website `HAS_PROPERTY=5,922`, `real_edge_missing=0`; AVmatrix-GO `HAS_PROPERTY=2,769`, `real_edge_missing=0`; invalid edges `0` | recorded |
| P3-A post-ownership access taxonomy | `.tmp\p3a-access-candidates-website-20260519.json`, `.tmp\p3a-access-candidates-avmatrix-go-20260519.json` | Website `total=24,542 resolved=3`; AVmatrix-GO `total=21,132 resolved=5,128`; biggest bucket remains `missing_receiver_type` | recorded |
| P3-B/P3-D access validation | `.tmp\p3b-property-access-website-20260519.json`, `.tmp\p3b-access-candidates-website-20260519.json`, `.tmp\p3b-property-access-avmatrix-go-20260519.json`, `.tmp\p3b-access-candidates-avmatrix-go-20260519.json` | Website final `ACCESSES=2,769`, candidate `resolved=4,978`; AVmatrix-GO final `ACCESSES=2,746`, candidate `resolved=5,110`; invalid owner edges `0` | recorded |
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

## P2 Ownership Validation

Date: 2026-05-19

Artifacts:

- Website graph snapshot: `.tmp\p2-property-access-website-go-graph-20260519.json`
- Website gate output: `.tmp\p2-property-access-website-20260519.json`
- AVmatrix-GO graph snapshot: `.tmp\p2-property-access-avmatrix-go-graph-20260519.json`
- AVmatrix-GO gate output: `.tmp\p2-property-access-avmatrix-go-20260519.json`

Commands:

```powershell
.\avmatrix-launcher\server-bundle\avmatrix.exe analyze E:\Website --force --skip-agents-md --no-stats
Copy-Item -LiteralPath 'E:\Website\.avmatrix\graph.json' -Destination '.tmp\p2-property-access-website-go-graph-20260519.json' -Force
go run ./cmd/property-access-audit -repo E:\Website -graph .tmp\p2-property-access-website-go-graph-20260519.json -out .tmp\p2-property-access-website-20260519.json -max-examples 20

.\avmatrix-launcher\server-bundle\avmatrix.exe analyze E:\AVmatrix-GO --force --skip-agents-md --no-stats
Copy-Item -LiteralPath 'E:\AVmatrix-GO\.avmatrix\graph.json' -Destination '.tmp\p2-property-access-avmatrix-go-graph-20260519.json' -Force
go run ./cmd/property-access-audit -repo E:\AVmatrix-GO -graph .tmp\p2-property-access-avmatrix-go-graph-20260519.json -out .tmp\p2-property-access-avmatrix-go-20260519.json -max-examples 20
```

Analyze runtime and graph size:

| Workload | Analyze runtime | Files scanned | Files parsed | Graph nodes | Graph relationships |
|---|---:|---:|---:|---:|---:|
| `E:\Website` | 18,684.0 ms | 1,870 | 998 | 27,764 | 54,972 |
| `E:\AVmatrix-GO` | 15,755.6 ms | 682 | 527 | 20,241 | 48,415 |

Ownership gate delta:

| Workload | Baseline `Property` | P2 `Property` | Baseline owner-linked | P2 owner-linked | Baseline `HAS_PROPERTY` | P2 `HAS_PROPERTY` | Baseline `ACCESSES` | P2 `ACCESSES` | Baseline real missing | P2 real missing | Invalid P2 edges |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| `E:\Website` | 5,222 | 6,905 | 3 | 5,129 | 3 | 5,129 | 3 | 3 | 3,627 | 392 | 0 |
| `E:\AVmatrix-GO` | 3,046 | 3,089 | 2,708 | 2,762 | 2,708 | 2,762 | 2,724 | 2,751 | 21 | 12 | 0 |

Language breakdown after P2:

| Workload | Language | `Property` | Owner-linked | Standalone |
|---|---|---:|---:|---:|
| `E:\Website` | TypeScript | 6,905 | 5,129 | 1,776 |
| `E:\AVmatrix-GO` | Go | 2,508 | 2,302 | 206 |
| `E:\AVmatrix-GO` | TypeScript | 581 | 460 | 121 |

P2 category taxonomy:

| Workload | Category | Count |
|---|---|---:|
| `E:\Website` | `tsjs_type_alias_object_literal_member` | 5,518 |
| `E:\Website` | `tsjs_typed_shape_or_binding_property` | 891 |
| `E:\Website` | `tsjs_runtime_object_literal_key` | 320 |
| `E:\Website` | `tsjs_destructuring_or_binding_pattern` | 149 |
| `E:\Website` | `tsjs_unclassified_property` | 24 |
| `E:\Website` | `tsjs_class_field` | 3 |
| `E:\AVmatrix-GO` | `go_owner_linked_struct` | 2,302 |
| `E:\AVmatrix-GO` | `go_typed_property_without_owner` | 206 |
| `E:\AVmatrix-GO` | `tsjs_interface_property_signature` | 449 |
| `E:\AVmatrix-GO` | `tsjs_destructuring_or_binding_pattern` | 58 |
| `E:\AVmatrix-GO` | `tsjs_typed_shape_or_binding_property` | 24 |
| `E:\AVmatrix-GO` | `tsjs_runtime_object_literal_key` | 24 |
| `E:\AVmatrix-GO` | `tsjs_type_alias_object_literal_member` | 15 |
| `E:\AVmatrix-GO` | `tsjs_class_field` | 8 |
| `E:\AVmatrix-GO` | `tsjs_unclassified_property` | 3 |

Orphan and graph-truth after P2:

| Workload | `owner_linked` | `false_orphan` | `true_orphan` | `unknown` | `edge_present` | `real_edge_missing` | `true_no_edge` | `unknown_no_edge` |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| `E:\Website` | 5,129 | 392 | 469 | 915 | 5,129 | 392 | 469 | 915 |
| `E:\AVmatrix-GO` | 2,762 | 12 | 82 | 233 | 2,762 | 12 | 82 | 233 |

P2 interpretation:

- The slice fixed the first large supported-language cluster by linking direct TypeScript type-alias object members to their `TypeAlias` owner.
- The schema now permits the corresponding `TypeAlias -> Property` relationship, so the graph loads through LadybugDB instead of failing at `db_load`.
- Runtime object keys, destructuring/binding pattern properties, and nested shape properties remain unlinked unless a stable owner model exists.
- `ACCESSES` was not the main P2 target. Website stayed at `3`; AVmatrix-GO increased from `2,724` to `2,751` as a downstream effect of improved owner-member indexing.

## P2-D Go Anonymous Struct Classification

Date: 2026-05-19

Artifacts:

- AVmatrix-GO gate output: `.tmp\p2d-property-access-avmatrix-go-20260519.json`
- Website guard output: `.tmp\p2d-property-access-website-20260519.json`

Commands:

```powershell
go run ./cmd/property-access-audit -repo E:\AVmatrix-GO -graph .tmp\p2-property-access-avmatrix-go-graph-20260519.json -out .tmp\p2d-property-access-avmatrix-go-20260519.json -max-examples 20
go run ./cmd/property-access-audit -repo E:\Website -graph .tmp\p2-property-access-website-go-graph-20260519.json -out .tmp\p2d-property-access-website-20260519.json -max-examples 20
```

AVmatrix-GO Go-language taxonomy after P2-D:

| Language | `Property` | Owner-linked | Standalone | `go_owner_linked_struct` | `go_anonymous_struct_field` | Go `true_no_edge` | Go `unknown_no_edge` |
|---|---:|---:|---:|---:|---:|---:|---:|
| Go | 2,508 | 2,302 | 206 | 2,302 | 206 | 206 | 0 |

Whole-workload graph-truth after P2-D:

| Workload | `Property` | `HAS_PROPERTY` | `ACCESSES` | `edge_present` | `real_edge_missing` | `true_no_edge` | `unknown_no_edge` | Invalid edges |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| `E:\AVmatrix-GO` | 3,089 | 2,762 | 2,751 | 2,762 | 12 | 288 | 27 | 0 |
| `E:\Website` | 6,905 | 5,129 | 3 | 5,129 | 392 | 469 | 915 | 0 |

P2-D interpretation:

- The previous AVmatrix-GO `go_typed_property_without_owner=206` bucket is not a missing owner bucket.
- The source samples are anonymous struct fields such as table-test rows, local JSON decode shapes, and nested anonymous response structs.
- Those fields have no stable named graph owner under the current model, so the correct graph truth is `true_no_edge`.
- No `HAS_PROPERTY` edges were added for this bucket.

## P2-E Remaining Ownership Clusters

Date: 2026-05-19

Artifacts:

- Website graph snapshot: `.tmp\p2e-property-access-website-go-graph-20260519.json`
- Website gate output: `.tmp\p2e-property-access-website-20260519.json`
- AVmatrix-GO graph snapshot: `.tmp\p2e-property-access-avmatrix-go-graph-20260519.json`
- AVmatrix-GO gate output: `.tmp\p2e-property-access-avmatrix-go-20260519.json`

Commands:

```powershell
.\avmatrix-launcher\server-bundle\avmatrix.exe analyze E:\Website --force --skip-agents-md --no-stats
Copy-Item -LiteralPath 'E:\Website\.avmatrix\graph.json' -Destination '.tmp\p2e-property-access-website-go-graph-20260519.json' -Force
go run ./cmd/property-access-audit -repo E:\Website -graph .tmp\p2e-property-access-website-go-graph-20260519.json -out .tmp\p2e-property-access-website-20260519.json -max-examples 20

.\avmatrix-launcher\server-bundle\avmatrix.exe analyze E:\AVmatrix-GO --force --skip-agents-md --no-stats
Copy-Item -LiteralPath 'E:\AVmatrix-GO\.avmatrix\graph.json' -Destination '.tmp\p2e-property-access-avmatrix-go-graph-20260519.json' -Force
go run ./cmd/property-access-audit -repo E:\AVmatrix-GO -graph .tmp\p2e-property-access-avmatrix-go-graph-20260519.json -out .tmp\p2e-property-access-avmatrix-go-20260519.json -max-examples 20
```

Analyze runtime and graph size:

| Workload | Analyze runtime | Files scanned | Files parsed | Graph nodes | Graph relationships |
|---|---:|---:|---:|---:|---:|
| `E:\Website` | 17,093.2 ms | 1,870 | 998 | 27,956 | 55,957 |
| `E:\AVmatrix-GO` | 14,805.4 ms | 682 | 527 | 20,268 | 48,460 |

Ownership gate delta from P2-D to P2-E:

| Workload | P2-D/P2 `Property` | P2-E `Property` | P2-D/P2 `HAS_PROPERTY` | P2-E `HAS_PROPERTY` | P2-D/P2 `ACCESSES` | P2-E `ACCESSES` | P2-D/P2 real missing | P2-E real missing | Invalid P2-E edges |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| `E:\Website` | 6,905 | 7,097 | 5,129 | 5,922 | 3 | 3 | 392 | 0 | 0 |
| `E:\AVmatrix-GO` | 3,089 | 3,096 | 2,762 | 2,769 | 2,751 | 2,754 | 12 | 0 | 0 |

Owner labels after P2-E:

| Workload | Owner label | `HAS_PROPERTY` |
|---|---|---:|
| `E:\Website` | Class | 3 |
| `E:\Website` | Property | 793 |
| `E:\Website` | TypeAlias | 5,126 |
| `E:\AVmatrix-GO` | Class | 8 |
| `E:\AVmatrix-GO` | Interface | 419 |
| `E:\AVmatrix-GO` | Property | 25 |
| `E:\AVmatrix-GO` | Struct | 2,302 |
| `E:\AVmatrix-GO` | TypeAlias | 15 |

P2-E category taxonomy:

| Workload | Category | Count |
|---|---|---:|
| `E:\Website` | `tsjs_type_alias_object_literal_member` | 5,126 |
| `E:\Website` | `tsjs_owner_linked_property` | 793 |
| `E:\Website` | `tsjs_inline_type_literal_property` | 807 |
| `E:\Website` | `tsjs_runtime_object_literal_key` | 313 |
| `E:\Website` | `tsjs_destructuring_or_binding_pattern` | 36 |
| `E:\Website` | `tsjs_typed_shape_or_binding_property` | 17 |
| `E:\Website` | `tsjs_class_field` | 3 |
| `E:\Website` | `tsjs_unclassified_property` | 2 |
| `E:\AVmatrix-GO` | `go_owner_linked_struct` | 2,302 |
| `E:\AVmatrix-GO` | `go_anonymous_struct_field` | 206 |
| `E:\AVmatrix-GO` | `tsjs_interface_property_signature` | 419 |
| `E:\AVmatrix-GO` | `tsjs_inline_type_literal_property` | 92 |
| `E:\AVmatrix-GO` | `tsjs_owner_linked_property` | 25 |
| `E:\AVmatrix-GO` | `tsjs_runtime_object_literal_key` | 24 |
| `E:\AVmatrix-GO` | `tsjs_type_alias_object_literal_member` | 15 |
| `E:\AVmatrix-GO` | `tsjs_class_field` | 8 |
| `E:\AVmatrix-GO` | `tsjs_typed_shape_or_binding_property` | 3 |
| `E:\AVmatrix-GO` | `tsjs_destructuring_or_binding_pattern` | 2 |

Graph-truth after P2-E:

| Workload | `owner_linked` | `false_orphan` | `true_orphan` | `unknown` | `edge_present` | `real_edge_missing` | `true_no_edge` | `unknown_no_edge` |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| `E:\Website` | 5,922 | 0 | 1,156 | 19 | 5,922 | 0 | 1,156 | 19 |
| `E:\AVmatrix-GO` | 2,769 | 0 | 324 | 3 | 2,769 | 0 | 324 | 3 |

P2-E interpretation:

- Nested TS/JS object-shape members now use defensible `Property -> Property` ownership.
- Inline anonymous TS/JS type literals are classified as true no-edge, not false missing owner.
- Phase 2 ownership false-orphan denominator is now closed for the measured workloads: both have `real_edge_missing=0`.

## P3-A Post-Ownership Access Taxonomy

Date: 2026-05-19

Artifacts:

- Website access candidate output: `.tmp\p3a-access-candidates-website-20260519.json`
- AVmatrix-GO access candidate output: `.tmp\p3a-access-candidates-avmatrix-go-20260519.json`

Commands:

```powershell
go run ./cmd/access-candidate-audit -repo E:\Website -out .tmp\p3a-access-candidates-website-20260519.json -max-examples 20
go run ./cmd/access-candidate-audit -repo E:\AVmatrix-GO -out .tmp\p3a-access-candidates-avmatrix-go-20260519.json -max-examples 20
```

Totals:

| Workload | Analyze runtime | Access candidates | Resolved | Unresolved | Final graph resolved accesses | Resolution unresolved references |
|---|---:|---:|---:|---:|---:|---:|
| `E:\Website` | 6,839 ms | 24,542 | 3 | 24,539 | 3 | 64,274 |
| `E:\AVmatrix-GO` | 5,227 ms | 21,132 | 5,128 | 16,004 | 5,128 | 50,360 |

Reason taxonomy:

| Workload | `resolved` | `missing_receiver_type` | `external_library_type` | `unsupported_syntax` | `missing_caller` | `missing_owner_link` | `false_positive_candidate` | `ambiguous_owner` |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| `E:\Website` | 3 | 13,512 | 9,660 | 1,313 | 53 | 1 | 0 | 0 |
| `E:\AVmatrix-GO` | 5,128 | 11,436 | 3,619 | 910 | 14 | 10 | 15 | 0 |

P3-A interpretation:

- Phase 2 ownership fixes did not materially change Website access resolution; Website remains blocked by receiver type inference.
- AVmatrix-GO resolved accesses increased from `5,112` to `5,128`, but `missing_receiver_type` remains the largest bucket.
- `missing_owner_link` is small after Phase 2: Website `1`, AVmatrix-GO `10`.
- P3-B should focus on explicit receiver type binding and locally inferable receiver types, while leaving `external_library_type` and `unsupported_syntax` out of scope.

## P3-B/P3-D Access Validation

Date: 2026-05-19

Artifacts:

- Website graph snapshot: `.tmp\p3b-property-access-website-go-graph-20260519.json`
- Website property/access gate output: `.tmp\p3b-property-access-website-20260519.json`
- Website access candidate output: `.tmp\p3b-access-candidates-website-20260519.json`
- AVmatrix-GO graph snapshot: `.tmp\p3b-property-access-avmatrix-go-graph-20260519.json`
- AVmatrix-GO property/access gate output: `.tmp\p3b-property-access-avmatrix-go-20260519.json`
- AVmatrix-GO access candidate output: `.tmp\p3b-access-candidates-avmatrix-go-20260519.json`

Commands:

```powershell
.\avmatrix-launcher\server-bundle\avmatrix.exe analyze E:\Website --force --skip-agents-md --no-stats
Copy-Item -LiteralPath 'E:\Website\.avmatrix\graph.json' -Destination '.tmp\p3b-property-access-website-go-graph-20260519.json' -Force
go run ./cmd/property-access-audit -repo E:\Website -graph .tmp\p3b-property-access-website-go-graph-20260519.json -out .tmp\p3b-property-access-website-20260519.json -max-examples 20
go run ./cmd/access-candidate-audit -repo E:\Website -out .tmp\p3b-access-candidates-website-20260519.json -max-examples 20

.\avmatrix-launcher\server-bundle\avmatrix.exe analyze E:\AVmatrix-GO --force --skip-agents-md --no-stats
Copy-Item -LiteralPath 'E:\AVmatrix-GO\.avmatrix\graph.json' -Destination '.tmp\p3b-property-access-avmatrix-go-graph-20260519.json' -Force
go run ./cmd/property-access-audit -repo E:\AVmatrix-GO -graph .tmp\p3b-property-access-avmatrix-go-graph-20260519.json -out .tmp\p3b-property-access-avmatrix-go-20260519.json -max-examples 20
go run ./cmd/access-candidate-audit -repo E:\AVmatrix-GO -out .tmp\p3b-access-candidates-avmatrix-go-20260519.json -max-examples 20
```

Analyze runtime and graph size:

| Workload | Analyze runtime | Files scanned | Files parsed | Graph nodes | Graph relationships |
|---|---:|---:|---:|---:|---:|
| `E:\Website` | 18,825.7 ms | 1,870 | 998 | 27,956 | 58,731 |
| `E:\AVmatrix-GO` | 14,784.4 ms | 682 | 527 | 20,282 | 48,550 |

Final graph property/access gate:

| Workload | `Property` | `HAS_PROPERTY` | `ACCESSES` | `real_edge_missing` | `true_no_edge` | `unknown_no_edge` | Invalid edges |
|---|---:|---:|---:|---:|---:|---:|---:|
| `E:\Website` | 7,097 | 5,922 | 2,769 | 0 | 1,156 | 19 | 0 |
| `E:\AVmatrix-GO` | 3,096 | 2,769 | 2,746 | 0 | 324 | 3 | 0 |

Final graph delta from P2-E:

| Workload | P2-E `ACCESSES` | P3-B `ACCESSES` | Delta |
|---|---:|---:|---:|
| `E:\Website` | 3 | 2,769 | +2,766 |
| `E:\AVmatrix-GO` | 2,754 | 2,746 | -8 |

Access candidate taxonomy:

| Workload | Analyze runtime | Access candidates | Resolved | Unresolved | Final graph resolved accesses | Resolution unresolved references |
|---|---:|---:|---:|---:|---:|---:|
| `E:\Website` | 7,072 ms | 24,542 | 4,978 | 19,564 | 4,978 | 59,289 |
| `E:\AVmatrix-GO` | 5,279 ms | 21,165 | 5,110 | 16,055 | 5,110 | 50,445 |

Reason taxonomy:

| Workload | `resolved` | `missing_receiver_type` | `external_library_type` | `unsupported_syntax` | `missing_caller` | `missing_owner_link` | `false_positive_candidate` | `ambiguous_owner` |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| `E:\Website` | 4,978 | 12,209 | 4,706 | 1,313 | 53 | 768 | 406 | 109 |
| `E:\AVmatrix-GO` | 5,110 | 11,464 | 3,622 | 910 | 14 | 10 | 35 | 0 |

Language breakdown:

| Workload | Language | Access candidates | Resolved | Unresolved |
|---|---|---:|---:|---:|
| `E:\Website` | JavaScript | 231 | 0 | 231 |
| `E:\Website` | TypeScript | 24,311 | 4,978 | 19,333 |
| `E:\AVmatrix-GO` | Go | 19,126 | 4,970 | 14,156 |
| `E:\AVmatrix-GO` | TypeScript | 2,039 | 140 | 1,899 |

P3-B interpretation:

- Website final graph `ACCESSES` increased from `3` to `2,769`, proving the first receiver-type cluster now emits real graph edges.
- Website candidate-level resolved accesses increased from `3` to `4,978`.
- `missing_receiver_type` remains the largest bucket, so P3-F still has work.
- `missing_owner_link`, `false_positive_candidate`, and `ambiguous_owner` increased on Website because many receiver owners are now known; those are newly visible follow-up buckets, not edges to force-link.
- AVmatrix-GO is not the main beneficiary of this TS/JS-heavy cluster. Its small final `ACCESSES` decrease is recorded for audit and will be watched in P3-E/P3-F instead of hidden.
