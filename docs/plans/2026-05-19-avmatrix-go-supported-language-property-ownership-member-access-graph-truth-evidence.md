# AVmatrix Go Supported-Language Property Ownership And Member Access Graph Truth Evidence

Date: 2026-05-19

Status: open

Companion plan: [2026-05-19-avmatrix-go-supported-language-property-ownership-member-access-graph-truth-plan.md](2026-05-19-avmatrix-go-supported-language-property-ownership-member-access-graph-truth-plan.md)

Companion benchmark: [2026-05-19-avmatrix-go-supported-language-property-ownership-member-access-graph-truth-benchmark.md](2026-05-19-avmatrix-go-supported-language-property-ownership-member-access-graph-truth-benchmark.md)

## Evidence Rules

- Record commands, artifacts, focused tests, full build, e2e proof, graph snapshots, and impact checks here.
- Keep benchmark measurements in the benchmark ledger.
- For doc-only commits, do not run AVmatrix solely for the commit.
- For implementation slices, use AVmatrix for codebase analysis and impact checks.

## Plan Creation Evidence

- Created a new plan because the 2026-05-16 graph accuracy plan is scoped to Go-local measured gates on `E:\AVmatrix-GO`.
- Initial follow-up scope was too narrow because it named TypeScript only.
- The corrected scope is supported-language property/access graph accuracy across every language that appears in AVmatrix-Go graph `Property` facts.
- `E:\Website` is retained as the TypeScript-heavy benchmark workload, not as the whole plan.
- The new plan does not reopen the completed 2026-05-16 plan.
- Plan creation commit: `b90a9de`.
- Scope correction commit: `0972c3d`.
- Access candidate audit commit: `73b1cf8`.
- P2 ownership graph truth commit: `b91c2cd`.
- P2-D Go anonymous struct truth commit: `0367bc1`.
- P2-E property ownership false-orphan closure commit: `de5f4d8`.
- P3-A post-ownership access taxonomy commit: `42ba98b`.
- P3-B/P3-C/P3-D awaited TypeScript property access commit: `08649e6`.

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
- The real follow-up is property ownership and member access semantics across supported languages.
- True orphan properties must remain visible as true orphans; the plan must not hide real repo structure problems by force-linking graph edges.

Current Website graph symptom:

```text
E:\Website avmatrix-go graph:
Property nodes:      5,222
HAS_PROPERTY edges:      3
ACCESSES edges:          3
```

Code pointers for the initial hypothesis:

- `internal/resolution/emit.go` emits `HAS_PROPERTY` only for `Property` definitions with owner metadata.
- `internal/resolution/indexes.go` indexes owner members only when `DefinitionFact.OwnerID` is non-empty.
- `internal/resolution` resolves `AccessFact` through owner/member lookup.
- Providers emit `Property` facts in multiple language families, including TS/JS, Go, Python, Java, C/C++, C#, Ruby, Rust, PHP, Kotlin, Swift, and Dart.

Initial hypothesis:

- Many properties are emitted without a stable owner link.
- Unowned properties do not enter owner/member indexes.
- Member access resolution cannot target unowned properties.
- Some unowned properties may be correct true orphans and should stay unlinked.
- The first implementation task must distinguish false orphans from true orphans by language before adding edges.

## Evidence Ledger

| Slice | Evidence | Result | Status |
|---|---|---|---|
| Plan creation | new plan/benchmark/evidence files | committed as `b90a9de` | done |
| P1-A cross-language audit gate | `internal/graphaccuracy/property_access.go`, `cmd/property-access-audit`, focused test | implemented and validated after full build | done |
| P1-B Website baseline gate | `.tmp\p1-property-access-website-baseline-20260519.json` | baseline recorded from explicit Website graph snapshot | done |
| P1-C AVmatrix-GO baseline gate | `.tmp\p1-property-access-avmatrix-go-baseline-20260519.json` | baseline recorded from explicit self-repo graph snapshot | done |
| P1-D language/category taxonomy | language and category counts in baseline artifacts | source categories classified across current graph languages | done |
| P1-E orphan truth classification | orphan status counts in baseline artifacts | true/false/unknown/external/intentionally-unmodeled owner status classified | done |
| P1-F graph truth classification | graph truth counts in baseline artifacts | real missing vs true no-edge classified | done |
| P1-G unresolved access reasons | `.tmp\p1-access-candidates-website-20260519.json`, `.tmp\p1-access-candidates-avmatrix-go-20260519.json` | unresolved access candidates classified by reason | done |
| P1-H Phase 2/3 targets | plan checklist update | measurable follow-up targets defined from P1 taxonomy | done |
| P2 ownership implementation | TS/JS direct type-alias property ownership and `TypeAlias -> Property` schema support | implemented and validated | done |
| P2 validation | full build, focused tests, analyze e2e, property gate e2e, benchmark update, impact check | recorded | done |
| P2-D Go ownership classification | Go anonymous struct fields classified as true no-edge | recorded | done |
| P2-E remaining ownership clusters | nested TS/JS shape ownership and inline type-literal no-edge classification | recorded | done |
| P3-A access taxonomy | post-ownership access candidate audit | recorded | done |
| P3-B access implementation | awaited TS/JS return binding, `TypeAlias` member owners, call-return enrichment guard | implemented and validated | done |
| P3-C access tests | provider, resolution, and access-audit focused tests | passed after full build | done |
| P3-D access validation | full build, focused tests, analyze/property/access e2e, benchmark update | recorded | done |
| P4 consumer checks | pending | pending | open |
| P5 final evidence | pending | pending | open |

## P1-A Implementation Evidence

Changed files:

- `internal/graphaccuracy/property_access.go`
- `internal/graphaccuracy/property_access_test.go`
- `cmd/property-access-audit/main.go`

The gate reads a graph snapshot JSON, uses the repo source only to classify source-line context, and emits JSON with:

- final graph `Property`, `HAS_PROPERTY`, and `ACCESSES` counts;
- property counts by language;
- property source categories;
- orphan statuses;
- graph truth statuses;
- invalid `HAS_PROPERTY` edge count;
- examples per bucket.

AVmatrix use for analysis before scope correction:

```powershell
.\avmatrix-launcher\server-bundle\avmatrix.exe status
.\avmatrix-launcher\server-bundle\avmatrix.exe context Run --repo E:\AVmatrix-GO --file internal\graphaccuracy\graphaccuracy.go
.\avmatrix-launcher\server-bundle\avmatrix.exe context main --repo E:\AVmatrix-GO --file cmd\graph-accuracy-probe\main.go
.\avmatrix-launcher\server-bundle\avmatrix.exe context <first attempted property gate symbol> --repo E:\AVmatrix-GO --file <first attempted property gate file>
```

Result:

```text
status initially reported stale at commit c233131 vs current c1c688e.
context Run returned graphaccuracy callers and callees for the existing gate pattern.
context main disambiguated existing graph-accuracy command wrapper symbols.
context on the first attempted property gate showed the first gate was wired only through a TypeScript-specific command and API.
```

Full build before final validation:

```powershell
.\avmatrix-launcher\build.ps1
```

Result: passed. Vite reported chunk-size and dynamic-import warnings only.

Focused tests after full build:

```powershell
go test ./internal/graphaccuracy ./cmd/property-access-audit
```

Result:

```text
ok  	github.com/tamnguyendinh/avmatrix-go/internal/graphaccuracy
?   	github.com/tamnguyendinh/avmatrix-go/cmd/property-access-audit	[no test files]
```

AVmatrix refresh after the cross-language correction:

```powershell
.\avmatrix-launcher\server-bundle\avmatrix.exe analyze E:\AVmatrix-GO --force --skip-agents-md --no-stats
```

Result:

```text
analyzed E:\AVmatrix-GO
files: scanned=678 parsed=523 unsupported=155 failed=0
graph: nodes=20089 relationships=48012 path=E:\AVmatrix-GO\.avmatrix\graph.json
```

CLI/runtime e2e after full build:

```powershell
go run ./cmd/property-access-audit -repo E:\Website -graph .tmp\p1-property-access-website-go-graph-20260519.json -out .tmp\p1-property-access-website-baseline-20260519.json -max-examples 20
go run ./cmd/property-access-audit -repo E:\AVmatrix-GO -graph .tmp\p1-property-access-avmatrix-go-graph-20260519.json -out .tmp\p1-property-access-avmatrix-go-baseline-20260519.json -max-examples 20
```

Website result:

```text
wrote .tmp\p1-property-access-website-baseline-20260519.json
properties.total=5222 ownerLinked=3 standalone=5219 hasPropertyEdges=3 accessesEdges=3 invalidHasPropertyEdges=0
language.typescript.properties=5222 ownerLinked=3 standalone=5219
orphan.external_library_owned=0
orphan.false_orphan=3627
orphan.intentionally_unmodeled=0
orphan.owner_linked=3
orphan.true_orphan=430
orphan.unknown=1162
graphTruth.edge_present=3
graphTruth.invalid_synthetic_edge_risk=0
graphTruth.real_edge_missing=3627
graphTruth.true_no_edge=430
graphTruth.unknown_no_edge=1162
```

AVmatrix-GO result:

```text
wrote .tmp\p1-property-access-avmatrix-go-baseline-20260519.json
properties.total=3046 ownerLinked=2708 standalone=338 hasPropertyEdges=2708 accessesEdges=2724 invalidHasPropertyEdges=0
language.go.properties=2469 ownerLinked=2263 standalone=206
language.typescript.properties=577 ownerLinked=445 standalone=132
orphan.external_library_owned=0
orphan.false_orphan=21
orphan.intentionally_unmodeled=0
orphan.owner_linked=2708
orphan.true_orphan=82
orphan.unknown=235
graphTruth.edge_present=2708
graphTruth.invalid_synthetic_edge_risk=0
graphTruth.real_edge_missing=21
graphTruth.true_no_edge=82
graphTruth.unknown_no_edge=235
```

Staged AVmatrix impact check:

```powershell
.\avmatrix-launcher\server-bundle\avmatrix.exe detect-changes --scope staged --repo E:\AVmatrix-GO
```

Result summary:

```text
changed_files=9
changed_count=176
affected_count=13
risk_level=high
affected processes include the new property-access audit command flow:
- Main -> GraphFile
- Main -> AddLanguageStats
- Main -> SortedBucketKeys
- Main -> SortedLanguageKeys
- Main -> WritePropertyAccessAuditResult
- BuildPropertyAccessAudit -> property classification helpers
```

Interpretation: the high risk level is expected for this slice because it adds a new command/API and many new graphaccuracy symbols. The changed runtime surface is isolated to the new audit gate and docs; existing product analysis behavior is not modified by this slice.

## P1-G Access Candidate Evidence

Changed files:

- `internal/resolution/access_audit.go`
- `internal/resolution/access_audit_test.go`
- `internal/graphaccuracy/access_candidate.go`
- `cmd/access-candidate-audit/main.go`

The gate runs analyze with ScopeIR retained and classifies access candidates before final graph edges hide unresolved cases.

Full build before validation:

```powershell
.\avmatrix-launcher\build.ps1
```

Result: passed. Vite reported chunk-size and dynamic-import warnings only.

Focused tests after full build:

```powershell
go test ./internal/resolution ./internal/graphaccuracy ./cmd/access-candidate-audit
```

Result:

```text
ok  	github.com/tamnguyendinh/avmatrix-go/internal/resolution
ok  	github.com/tamnguyendinh/avmatrix-go/internal/graphaccuracy
?   	github.com/tamnguyendinh/avmatrix-go/cmd/access-candidate-audit	[no test files]
```

CLI/runtime e2e after full build:

```powershell
go run ./cmd/access-candidate-audit -repo E:\Website -out .tmp\p1-access-candidates-website-20260519.json -max-examples 20
go run ./cmd/access-candidate-audit -repo E:\AVmatrix-GO -out .tmp\p1-access-candidates-avmatrix-go-20260519.json -max-examples 20
```

Website result:

```text
wrote .tmp\p1-access-candidates-website-20260519.json
accessCandidates.total=24542 resolved=3 unresolved=24539 analyzeMillis=7365 resolvedAccesses=3 unresolvedReferences=64274
language.javascript.accessCandidates=231 resolved=0 unresolved=231
language.typescript.accessCandidates=24311 resolved=3 unresolved=24308
reason.ambiguous_owner=0
reason.external_library_type=9660
reason.false_positive_candidate=0
reason.missing_caller=53
reason.missing_owner_link=1
reason.missing_receiver_type=13512
reason.resolved=3
reason.unsupported_syntax=1313
```

AVmatrix-GO result:

```text
wrote .tmp\p1-access-candidates-avmatrix-go-20260519.json
accessCandidates.total=21098 resolved=5112 unresolved=15986 analyzeMillis=5862 resolvedAccesses=5112 unresolvedReferences=50282
language.go.accessCandidates=19059 resolved=4975 unresolved=14084
language.typescript.accessCandidates=2039 resolved=137 unresolved=1902
reason.ambiguous_owner=0
reason.external_library_type=3619
reason.false_positive_candidate=15
reason.missing_caller=14
reason.missing_owner_link=10
reason.missing_receiver_type=11418
reason.resolved=5112
reason.unsupported_syntax=910
```

AVmatrix refresh after the P1-G implementation:

```powershell
.\avmatrix-launcher\server-bundle\avmatrix.exe analyze E:\AVmatrix-GO --force --skip-agents-md --no-stats
```

Result:

```text
analyzed E:\AVmatrix-GO
files: scanned=682 parsed=527 unsupported=155 failed=0
graph: nodes=20223 relationships=48368 path=E:\AVmatrix-GO\.avmatrix\graph.json
```

P1-H target update:

```text
Phase 2 first target: TS/JS/SFC static shape ownership, especially Website tsjs_type_alias_object_literal_member real_edge_missing=3,627.
Phase 2 guardrail: keep runtime object keys and destructuring true-no-edge cases unlinked.
Phase 2 second target: classify AVmatrix-GO go_typed_property_without_owner=206 before any Go ownership expansion.
Phase 3 first target: reduce missing_receiver_type using explicit receiver type bindings.
Phase 3 guardrail: do not convert external_library_type or unsupported_syntax buckets into noisy edges.
Phase 3 small bucket: close or reclassify Website missing_owner_link=1 and AVmatrix-GO missing_owner_link=10 after ownership fixes.
```

Staged AVmatrix impact check:

```powershell
.\avmatrix-launcher\server-bundle\avmatrix.exe detect-changes --scope staged --repo E:\AVmatrix-GO
```

Result summary:

```text
changed_files=7
changed_count=132
affected_count=21
risk_level=critical
affected processes include the new access-candidate audit command flow:
- Main -> UnsupportedAccessReceiver
- Main -> AddAccessReason
- Main -> WriteAccessCandidateAuditResult
- Main -> SortedAccessLanguageKeys
- Main -> SortedAccessReasonKeys
- RunAccessCandidateAudit -> AuditAccessCandidates
- AuditAccessCandidates -> resolver helper paths
```

Interpretation: the critical risk level is expected for this slice because the new audit command runs analyze and reuses resolution internals to classify access candidates. It does not change the production resolver output path; it adds an audit/readout path plus focused tests.

## P2-A/P2-C Ownership Evidence

Changed files:

- `internal/providers/tsjs/definitions.go`
- `internal/providers/tsjs/extract_test.go`
- `internal/lbugschema/schema.go`
- `internal/lbugschema/schema_test.go`
- `internal/mcp/resources.go`

Implementation summary:

- Direct `property_signature` members inside a TypeScript `type_alias_declaration` now resolve their owner to the containing `TypeAlias`.
- Nested property signatures remain unowned until nested shape ownership has a defensible model.
- LadybugDB relation schema now includes `TypeAlias -> Property`, which is required for the new final graph `HAS_PROPERTY` edges to load.
- MCP schema text now describes `HAS_PROPERTY` as supported by `Class`, `Struct`, `Interface`, `TypeAlias`, and other supported owners.

Initial validation blocker found and fixed:

```text
db_load phase: copy relationships TypeAlias->Property: schema pair unsupported
```

Root cause: the extractor correctly emitted `TypeAlias -> Property`, but `internal/lbugschema.RelationPairs` did not allow that source/target pair. The schema pair was added and covered by `internal/lbugschema/schema_test.go`.

Full build before tests:

```powershell
.\avmatrix-launcher\build.ps1
```

Result: passed. Vite reported chunk-size and dynamic/static import warnings only.

Focused tests after full build:

```powershell
go test ./internal/lbugschema ./internal/lbugload ./internal/providers/tsjs ./internal/resolution ./internal/graphaccuracy ./cmd/property-access-audit
```

Result:

```text
ok  	github.com/tamnguyendinh/avmatrix-go/internal/lbugschema	0.938s
ok  	github.com/tamnguyendinh/avmatrix-go/internal/lbugload	1.155s
ok  	github.com/tamnguyendinh/avmatrix-go/internal/providers/tsjs	(cached)
ok  	github.com/tamnguyendinh/avmatrix-go/internal/resolution	(cached)
ok  	github.com/tamnguyendinh/avmatrix-go/internal/graphaccuracy	0.528s
?   	github.com/tamnguyendinh/avmatrix-go/cmd/property-access-audit	[no test files]
```

Analyze e2e after full build:

```powershell
.\avmatrix-launcher\server-bundle\avmatrix.exe analyze E:\Website --force --skip-agents-md --no-stats
.\avmatrix-launcher\server-bundle\avmatrix.exe analyze E:\AVmatrix-GO --force --skip-agents-md --no-stats
```

Website result:

```text
analyzed E:\Website
files: scanned=1870 parsed=998 unsupported=872 failed=0
graph: nodes=27764 relationships=54972 path=E:\Website\.avmatrix\graph.json
ANALYZE_MS=18,684.0
```

AVmatrix-GO result:

```text
analyzed E:\AVmatrix-GO
files: scanned=682 parsed=527 unsupported=155 failed=0
graph: nodes=20241 relationships=48415 path=E:\AVmatrix-GO\.avmatrix\graph.json
ANALYZE_MS=15,755.6
```

Graph snapshots:

```powershell
Copy-Item -LiteralPath 'E:\Website\.avmatrix\graph.json' -Destination '.tmp\p2-property-access-website-go-graph-20260519.json' -Force
Copy-Item -LiteralPath 'E:\AVmatrix-GO\.avmatrix\graph.json' -Destination '.tmp\p2-property-access-avmatrix-go-graph-20260519.json' -Force
```

CLI/runtime property gate e2e after full build:

```powershell
go run ./cmd/property-access-audit -repo E:\Website -graph .tmp\p2-property-access-website-go-graph-20260519.json -out .tmp\p2-property-access-website-20260519.json -max-examples 20
go run ./cmd/property-access-audit -repo E:\AVmatrix-GO -graph .tmp\p2-property-access-avmatrix-go-graph-20260519.json -out .tmp\p2-property-access-avmatrix-go-20260519.json -max-examples 20
```

Website result:

```text
wrote .tmp\p2-property-access-website-20260519.json
properties.total=6905 ownerLinked=5129 standalone=1776 hasPropertyEdges=5129 accessesEdges=3 invalidHasPropertyEdges=0
language.typescript.properties=6905 ownerLinked=5129 standalone=1776
orphan.false_orphan=392
orphan.owner_linked=5129
orphan.true_orphan=469
orphan.unknown=915
graphTruth.edge_present=5129
graphTruth.real_edge_missing=392
graphTruth.true_no_edge=469
graphTruth.unknown_no_edge=915
```

AVmatrix-GO result:

```text
wrote .tmp\p2-property-access-avmatrix-go-20260519.json
properties.total=3089 ownerLinked=2762 standalone=327 hasPropertyEdges=2762 accessesEdges=2751 invalidHasPropertyEdges=0
language.go.properties=2508 ownerLinked=2302 standalone=206
language.typescript.properties=581 ownerLinked=460 standalone=121
orphan.false_orphan=12
orphan.owner_linked=2762
orphan.true_orphan=82
orphan.unknown=233
graphTruth.edge_present=2762
graphTruth.real_edge_missing=12
graphTruth.true_no_edge=82
graphTruth.unknown_no_edge=233
```

Interpretation:

- P2-A reduced Website `real_edge_missing` from `3,627` to `392` and increased Website `HAS_PROPERTY` from `3` to `5,129`.
- P2-A did not fabricate links for true-no-edge classes: invalid synthetic edge risk remains `0`.
- AVmatrix-GO `go_typed_property_without_owner=206` was intentionally unchanged by P2-A and is classified by P2-D below.
- Website `ACCESSES` is unchanged at `3`; access resolution remains Phase 3.

Staged AVmatrix impact check:

```powershell
.\avmatrix-launcher\server-bundle\avmatrix.exe detect-changes --scope staged --repo E:\AVmatrix-GO
```

Result summary:

```text
changed_files=8
changed_count=17
affected_count=0
risk_level=low
affected_processes=[]
changed symbols include:
- RelationPairs
- TestSchemaQueriesPreserveDDLShape
- TestSchemaSurfaceCoversLegacyCoreAndModernNodeTypes
- schemaResource
- directTypeAliasObjectOwner
- collector.ownerDeclarationFor
- collector.ownerDeclarationNameFor
- collector.ownerDefIDFor
- TestExtractTypeAliasObjectPropertiesHaveDirectOwner
```

Interpretation: AVmatrix reports low impact and no affected process traces for the staged slice. The risk is covered by full build, schema/load tests, TS/JS provider tests, analyze e2e on both workloads, and the property/access graph gate outputs recorded above.

## P2-D Go Anonymous Struct Evidence

Changed files:

- `internal/graphaccuracy/property_access.go`
- `internal/graphaccuracy/property_access_test.go`

Implementation summary:

- The property/access gate now classifies unowned Go fields inside anonymous `struct { ... }` blocks as `go_anonymous_struct_field`.
- `go_anonymous_struct_field` maps to `true_orphan` and `true_no_edge`.
- This does not emit or change final graph `HAS_PROPERTY` edges; it corrects the audit denominator so true orphan facts are not treated as missing graph links.

Representative source samples from `.tmp\p2d-property-access-avmatrix-go-20260519.json`:

```text
avmatrix-launcher/src/main_test.go:16 args :: args []string
avmatrix-launcher/src/main_test.go:17 want :: want string
internal/cli/hook_command.go:134 LastCommit :: LastCommit string `json:"lastCommit"`
internal/cli/hook_command.go:135 Stats :: Stats      struct {
internal/cli/hook_command.go:136 Embeddings :: Embeddings int `json:"embeddings"`
internal/cli/hook_command_test.go:83 name :: name      string
internal/cli/hook_command_test.go:84 toolName :: toolName  string
internal/cli/package_command_test.go:161 Bin :: Bin     map[string]string `json:"bin"`
internal/cli/tool_command.go:236 Content :: Content []struct {
```

Interpretation:

- These are table-test row fields, local anonymous JSON shapes, or nested anonymous response shapes.
- They expose useful field facts, but they do not have a stable named owner that should receive `HAS_PROPERTY`.
- Leaving them unowned is correct graph truth; the gate now counts them as true no-edge instead of unknown.

Full build before tests:

```powershell
.\avmatrix-launcher\build.ps1
```

Result: passed. Vite reported chunk-size and dynamic/static import warnings only.

Focused tests after full build:

```powershell
go test ./internal/graphaccuracy ./cmd/property-access-audit
```

Result:

```text
ok  	github.com/tamnguyendinh/avmatrix-go/internal/graphaccuracy	0.609s
?   	github.com/tamnguyendinh/avmatrix-go/cmd/property-access-audit	[no test files]
```

CLI/runtime property gate e2e after full build:

```powershell
go run ./cmd/property-access-audit -repo E:\AVmatrix-GO -graph .tmp\p2-property-access-avmatrix-go-graph-20260519.json -out .tmp\p2d-property-access-avmatrix-go-20260519.json -max-examples 20
go run ./cmd/property-access-audit -repo E:\Website -graph .tmp\p2-property-access-website-go-graph-20260519.json -out .tmp\p2d-property-access-website-20260519.json -max-examples 20
```

AVmatrix-GO result:

```text
wrote .tmp\p2d-property-access-avmatrix-go-20260519.json
properties.total=3089 ownerLinked=2762 standalone=327 hasPropertyEdges=2762 accessesEdges=2751 invalidHasPropertyEdges=0
language.go.properties=2508 ownerLinked=2302 standalone=206
language.typescript.properties=581 ownerLinked=460 standalone=121
orphan.false_orphan=12
orphan.owner_linked=2762
orphan.true_orphan=288
orphan.unknown=27
graphTruth.edge_present=2762
graphTruth.real_edge_missing=12
graphTruth.true_no_edge=288
graphTruth.unknown_no_edge=27
```

Go-only P2-D result:

```text
go_owner_linked_struct=2302
go_anonymous_struct_field=206
go.owner_linked=2302
go.true_orphan=206
go.edge_present=2302
go.true_no_edge=206
```

Website guard result:

```text
wrote .tmp\p2d-property-access-website-20260519.json
properties.total=6905 ownerLinked=5129 standalone=1776 hasPropertyEdges=5129 accessesEdges=3 invalidHasPropertyEdges=0
orphan.false_orphan=392
orphan.owner_linked=5129
orphan.true_orphan=469
orphan.unknown=915
graphTruth.real_edge_missing=392
graphTruth.true_no_edge=469
graphTruth.unknown_no_edge=915
```

Staged AVmatrix impact check:

```powershell
.\avmatrix-launcher\server-bundle\avmatrix.exe detect-changes --scope staged --repo E:\AVmatrix-GO
```

Result summary:

```text
changed_files=5
changed_count=10
affected_count=6
risk_level=high
affected process names:
- Main -> LooksDestructuringOrBinding
- Main -> IsTSJSLanguage
- BuildPropertyAccessAudit -> ContextContainsBlockHeader
- BuildPropertyAccessAudit -> LooksInsideTypeAliasObject
- BuildPropertyAccessAudit -> PropString
- Main -> PropString
```

Interpretation: high risk is expected because the slice changes the property/access audit classifier used by `cmd/property-access-audit`. The production analyzer graph is not changed by this slice; validation is covered by full build, focused graphaccuracy tests, and property gate e2e on both `E:\AVmatrix-GO` and `E:\Website`.

## P2-E Remaining Ownership Evidence

Changed files:

- `internal/providers/tsjs/definitions.go`
- `internal/providers/tsjs/extract_test.go`
- `internal/graphaccuracy/property_access.go`
- `internal/graphaccuracy/property_access_test.go`

Implementation summary:

- Nested TS/JS `property_signature` definitions now use their nearest parent `Property` as owner.
- Direct type-alias object members keep `TypeAlias` as owner.
- Inline anonymous TS/JS type literals, such as `useRef<{ ... }>` and `useState<{ ... }>`, remain unowned and are classified as `true_no_edge`.
- This adds defensible `Property -> Property` `HAS_PROPERTY` edges without linking anonymous inline shapes to unrelated interfaces or variables.

Full build before tests:

```powershell
.\avmatrix-launcher\build.ps1
```

Result: passed. Vite reported chunk-size and dynamic/static import warnings only.

Focused tests after full build:

```powershell
go test ./internal/providers/tsjs ./internal/resolution ./internal/graphaccuracy ./cmd/property-access-audit
```

Result:

```text
ok  	github.com/tamnguyendinh/avmatrix-go/internal/providers/tsjs
ok  	github.com/tamnguyendinh/avmatrix-go/internal/resolution
ok  	github.com/tamnguyendinh/avmatrix-go/internal/graphaccuracy
?   	github.com/tamnguyendinh/avmatrix-go/cmd/property-access-audit	[no test files]
```

Analyze e2e after full build:

```powershell
.\avmatrix-launcher\server-bundle\avmatrix.exe analyze E:\Website --force --skip-agents-md --no-stats
.\avmatrix-launcher\server-bundle\avmatrix.exe analyze E:\AVmatrix-GO --force --skip-agents-md --no-stats
```

Website result:

```text
analyzed E:\Website
files: scanned=1870 parsed=998 unsupported=872 failed=0
graph: nodes=27956 relationships=55957 path=E:\Website\.avmatrix\graph.json
ANALYZE_MS=17,093.2
```

AVmatrix-GO result:

```text
analyzed E:\AVmatrix-GO
files: scanned=682 parsed=527 unsupported=155 failed=0
graph: nodes=20268 relationships=48460 path=E:\AVmatrix-GO\.avmatrix\graph.json
ANALYZE_MS=14,805.4
```

Graph snapshots:

```powershell
Copy-Item -LiteralPath 'E:\Website\.avmatrix\graph.json' -Destination '.tmp\p2e-property-access-website-go-graph-20260519.json' -Force
Copy-Item -LiteralPath 'E:\AVmatrix-GO\.avmatrix\graph.json' -Destination '.tmp\p2e-property-access-avmatrix-go-graph-20260519.json' -Force
```

CLI/runtime property gate e2e after full build:

```powershell
go run ./cmd/property-access-audit -repo E:\Website -graph .tmp\p2e-property-access-website-go-graph-20260519.json -out .tmp\p2e-property-access-website-20260519.json -max-examples 20
go run ./cmd/property-access-audit -repo E:\AVmatrix-GO -graph .tmp\p2e-property-access-avmatrix-go-graph-20260519.json -out .tmp\p2e-property-access-avmatrix-go-20260519.json -max-examples 20
```

Website result:

```text
wrote .tmp\p2e-property-access-website-20260519.json
properties.total=7097 ownerLinked=5922 standalone=1175 hasPropertyEdges=5922 accessesEdges=3 invalidHasPropertyEdges=0
language.typescript.properties=7097 ownerLinked=5922 standalone=1175
orphan.false_orphan=0
orphan.owner_linked=5922
orphan.true_orphan=1156
orphan.unknown=19
graphTruth.edge_present=5922
graphTruth.real_edge_missing=0
graphTruth.true_no_edge=1156
graphTruth.unknown_no_edge=19
```

AVmatrix-GO result:

```text
wrote .tmp\p2e-property-access-avmatrix-go-20260519.json
properties.total=3096 ownerLinked=2769 standalone=327 hasPropertyEdges=2769 accessesEdges=2754 invalidHasPropertyEdges=0
language.go.properties=2508 ownerLinked=2302 standalone=206
language.typescript.properties=588 ownerLinked=467 standalone=121
orphan.false_orphan=0
orphan.owner_linked=2769
orphan.true_orphan=324
orphan.unknown=3
graphTruth.edge_present=2769
graphTruth.real_edge_missing=0
graphTruth.true_no_edge=324
graphTruth.unknown_no_edge=3
```

Interpretation:

- Website `real_edge_missing` moved from `392` to `0`.
- AVmatrix-GO `real_edge_missing` moved from `12` to `0`.
- Invalid synthetic edge risk remains `0`.
- The remaining `unknown_no_edge` buckets are small and intentionally not linked without more evidence: Website `19`, AVmatrix-GO `3`.

Staged AVmatrix impact check:

```powershell
.\avmatrix-launcher\server-bundle\avmatrix.exe detect-changes --scope staged --repo E:\AVmatrix-GO
```

Result summary:

```text
changed_files=7
changed_count=34
affected_count=4
risk_level=medium
affected process names:
- BuildPropertyAccessAudit -> GoFieldLineHasType
- BuildPropertyAccessAudit -> ContextContainsBlockHeader
- BuildPropertyAccessAudit -> LooksDestructuringOrBinding
- BuildPropertyAccessAudit -> LooksInsideTypeAliasObject
```

Interpretation: medium risk is expected because this slice changes TS/JS provider ownership and the property/access audit classifier. The slice is covered by full build, TS/JS provider tests, graphaccuracy tests, analyze e2e on both workloads, and final property gate outputs showing `real_edge_missing=0` with invalid edges still `0`.

## P3-A Access Taxonomy Evidence

Artifacts:

- `.tmp\p3a-access-candidates-website-20260519.json`
- `.tmp\p3a-access-candidates-avmatrix-go-20260519.json`

Commands:

```powershell
go run ./cmd/access-candidate-audit -repo E:\Website -out .tmp\p3a-access-candidates-website-20260519.json -max-examples 20
go run ./cmd/access-candidate-audit -repo E:\AVmatrix-GO -out .tmp\p3a-access-candidates-avmatrix-go-20260519.json -max-examples 20
```

Website result:

```text
wrote .tmp\p3a-access-candidates-website-20260519.json
accessCandidates.total=24542 resolved=3 unresolved=24539 analyzeMillis=6839 resolvedAccesses=3 unresolvedReferences=64274
language.javascript.accessCandidates=231 resolved=0 unresolved=231
language.typescript.accessCandidates=24311 resolved=3 unresolved=24308
reason.external_library_type=9660
reason.missing_owner_link=1
reason.missing_receiver_type=13512
reason.resolved=3
reason.unsupported_syntax=1313
```

AVmatrix-GO result:

```text
wrote .tmp\p3a-access-candidates-avmatrix-go-20260519.json
accessCandidates.total=21132 resolved=5128 unresolved=16004 analyzeMillis=5227 resolvedAccesses=5128 unresolvedReferences=50360
language.go.accessCandidates=19093 resolved=4991 unresolved=14102
language.typescript.accessCandidates=2039 resolved=137 unresolved=1902
reason.external_library_type=3619
reason.false_positive_candidate=15
reason.missing_owner_link=10
reason.missing_receiver_type=11436
reason.resolved=5128
reason.unsupported_syntax=910
```

Representative samples:

```text
Website missing_receiver_type:
- app/(account)/dashboard/billing/invoices/[invoiceId]/document/route.ts:123 receiver=result.model name=invoices
- app/(account)/dashboard/billing/invoices/[invoiceId]/document/route.ts:137 receiver=invoice name=invoiceId

Website missing_owner_link:
- modules/account-portal/server/profile/account-settings-write-rejection.ts:13 receiver=this name=name

AVmatrix-GO missing_receiver_type:
- avmatrix-launcher/server-wrapper/main.go:29 receiver=cmd name=Dir
- avmatrix-launcher/server-wrapper/main.go:54 receiver=os name=O_CREATE

AVmatrix-GO missing_owner_link:
- avmatrix-web/src/components/FileTreePanel.tsx:291 receiver=treeNode.graphNode name=id
- avmatrix-web/src/core/llm/session-client.ts:18 receiver=this name=name
```

Interpretation:

- The next large access blocker is `missing_receiver_type`.
- `missing_owner_link` is now small enough to handle after the receiver-type slice or as a focused cleanup.
- `external_library_type` and `unsupported_syntax` remain out of scope for P3-B because they need import/library modeling or richer receiver expression parsing.

## P3-B/P3-C/P3-D Access Resolution Evidence

Changed files:

- `internal/providers/tsjs/types.go`
- `internal/providers/tsjs/legacy_scope_captures_test.go`
- `internal/resolution/indexes.go`
- `internal/resolution/access_audit.go`
- `internal/resolution/access_audit_test.go`
- `internal/resolution/resolution_test.go`

Implementation summary:

- `const result = await readResult()` now binds `result` to the unwrapped `ReadResult` when `readResult(): Promise<ReadResult>`.
- Resolver member owner lookup now includes `TypeAlias` for member/property resolution, matching the Phase 2 `TypeAlias -> Property` ownership graph.
- Call-return enrichment no longer overwrites provider-derived `return-annotation` type bindings with less precise fallback call-return bindings.
- Access candidate audit uses the same `TypeAlias` member-owner lookup as production resolution.

Full build before tests:

```powershell
.\avmatrix-launcher\build.ps1
```

Result: passed. Vite reported chunk-size and dynamic/static import warnings only.

Focused tests after full build:

```powershell
go test ./internal/providers/tsjs ./internal/resolution ./internal/graphaccuracy ./cmd/access-candidate-audit
```

Result:

```text
ok  	github.com/tamnguyendinh/avmatrix-go/internal/providers/tsjs
ok  	github.com/tamnguyendinh/avmatrix-go/internal/resolution
ok  	github.com/tamnguyendinh/avmatrix-go/internal/graphaccuracy
?   	github.com/tamnguyendinh/avmatrix-go/cmd/access-candidate-audit	[no test files]
```

Focused test coverage added:

```text
TestExtractTypeScriptAwaitedPromiseReturnTypeBinding
TestResolveAwaitedPromiseReturnMemberAccess
TestAuditAccessCandidatesResolvesTypeAliasMembers
```

Analyze e2e after full build:

```powershell
.\avmatrix-launcher\server-bundle\avmatrix.exe analyze E:\Website --force --skip-agents-md --no-stats
.\avmatrix-launcher\server-bundle\avmatrix.exe analyze E:\AVmatrix-GO --force --skip-agents-md --no-stats
```

Website result:

```text
analyzed E:\Website
files: scanned=1870 parsed=998 unsupported=872 failed=0
graph: nodes=27956 relationships=58731 path=E:\Website\.avmatrix\graph.json
ANALYZE_MS=18,825.7
```

AVmatrix-GO result:

```text
analyzed E:\AVmatrix-GO
files: scanned=682 parsed=527 unsupported=155 failed=0
graph: nodes=20282 relationships=48550 path=E:\AVmatrix-GO\.avmatrix\graph.json
ANALYZE_MS=14,784.4
```

Graph snapshots:

```powershell
Copy-Item -LiteralPath 'E:\Website\.avmatrix\graph.json' -Destination '.tmp\p3b-property-access-website-go-graph-20260519.json' -Force
Copy-Item -LiteralPath 'E:\AVmatrix-GO\.avmatrix\graph.json' -Destination '.tmp\p3b-property-access-avmatrix-go-graph-20260519.json' -Force
```

CLI/runtime property gate e2e:

```powershell
go run ./cmd/property-access-audit -repo E:\Website -graph .tmp\p3b-property-access-website-go-graph-20260519.json -out .tmp\p3b-property-access-website-20260519.json -max-examples 20
go run ./cmd/property-access-audit -repo E:\AVmatrix-GO -graph .tmp\p3b-property-access-avmatrix-go-graph-20260519.json -out .tmp\p3b-property-access-avmatrix-go-20260519.json -max-examples 20
```

Website property/access gate result:

```text
wrote .tmp\p3b-property-access-website-20260519.json
properties.total=7097 ownerLinked=5922 standalone=1175 hasPropertyEdges=5922 accessesEdges=2769 invalidHasPropertyEdges=0
language.typescript.properties=7097 ownerLinked=5922 standalone=1175
graphTruth.real_edge_missing=0
graphTruth.true_no_edge=1156
graphTruth.unknown_no_edge=19
```

AVmatrix-GO property/access gate result:

```text
wrote .tmp\p3b-property-access-avmatrix-go-20260519.json
properties.total=3096 ownerLinked=2769 standalone=327 hasPropertyEdges=2769 accessesEdges=2746 invalidHasPropertyEdges=0
language.go.properties=2508 ownerLinked=2302 standalone=206
language.typescript.properties=588 ownerLinked=467 standalone=121
graphTruth.real_edge_missing=0
graphTruth.true_no_edge=324
graphTruth.unknown_no_edge=3
```

CLI/runtime access candidate e2e:

```powershell
go run ./cmd/access-candidate-audit -repo E:\Website -out .tmp\p3b-access-candidates-website-20260519.json -max-examples 20
go run ./cmd/access-candidate-audit -repo E:\AVmatrix-GO -out .tmp\p3b-access-candidates-avmatrix-go-20260519.json -max-examples 20
```

Website access candidate result:

```text
wrote .tmp\p3b-access-candidates-website-20260519.json
accessCandidates.total=24542 resolved=4978 unresolved=19564 analyzeMillis=7072 resolvedAccesses=4978 unresolvedReferences=59289
language.javascript.accessCandidates=231 resolved=0 unresolved=231
language.typescript.accessCandidates=24311 resolved=4978 unresolved=19333
reason.ambiguous_owner=109
reason.external_library_type=4706
reason.false_positive_candidate=406
reason.missing_caller=53
reason.missing_owner_link=768
reason.missing_receiver_type=12209
reason.resolved=4978
reason.unsupported_syntax=1313
```

AVmatrix-GO access candidate result:

```text
wrote .tmp\p3b-access-candidates-avmatrix-go-20260519.json
accessCandidates.total=21165 resolved=5110 unresolved=16055 analyzeMillis=5279 resolvedAccesses=5110 unresolvedReferences=50445
language.go.accessCandidates=19126 resolved=4970 unresolved=14156
language.typescript.accessCandidates=2039 resolved=140 unresolved=1899
reason.ambiguous_owner=0
reason.external_library_type=3622
reason.false_positive_candidate=35
reason.missing_caller=14
reason.missing_owner_link=10
reason.missing_receiver_type=11464
reason.resolved=5110
reason.unsupported_syntax=910
```

Interpretation:

- Website now has a large final graph `ACCESSES` expansion: `3 -> 2,769`.
- Website access candidate resolved count moved `3 -> 4,978`, and `missing_receiver_type` moved `13,512 -> 12,209`.
- The Website `missing_owner_link=768`, `false_positive_candidate=406`, and `ambiguous_owner=109` buckets are newly exposed because receivers now resolve to concrete owners. They are P3-E/P3-F work, not a reason to fabricate edges.
- AVmatrix-GO is recorded with final graph `ACCESSES=2,746` and access candidate `resolved=5,110`; this TS/JS-heavy slice primarily benefits Website.

Staged AVmatrix impact check:

```powershell
.\avmatrix-launcher\server-bundle\avmatrix.exe detect-changes --scope staged --repo E:\AVmatrix-GO
```

Result summary:

```text
changed_files=9
changed_count=39
affected_count=20
risk_level=critical
affected process names include:
- ResolveAccess -> DefRef
- ResolveAccess -> DispatchOwnerLabels
- ResolveAccess -> ParentScope
- ResolveAccess -> IsAnyLabel
- EnrichCallReturnTypeBindings -> CallableLabels
- EnrichCallReturnTypeBindings -> DispatchOwnerLabels
- RunAccessCandidateAudit -> DispatchOwnerLabels
- Main -> HasStandalonePropertyCandidate
```

Interpretation: critical risk is expected because this slice changes production member access resolution and call-return type enrichment, plus the audit classifier that reads the same resolver model. The risk is covered by full build, focused provider/resolution/graphaccuracy tests, analyze e2e on both workloads, property gate e2e with invalid edges still `0`, and access-candidate e2e metrics recorded above.
