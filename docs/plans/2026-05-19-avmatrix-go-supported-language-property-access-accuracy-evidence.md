# AVmatrix Go Supported-Language Property Access Accuracy Evidence

Date: 2026-05-19

Status: open

Companion plan: [2026-05-19-avmatrix-go-supported-language-property-access-accuracy-plan.md](2026-05-19-avmatrix-go-supported-language-property-access-accuracy-plan.md)

Companion benchmark: [2026-05-19-avmatrix-go-supported-language-property-access-accuracy-benchmark.md](2026-05-19-avmatrix-go-supported-language-property-access-accuracy-benchmark.md)

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
- Scope correction commit: pending current slice.

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
| P1-G unresolved access reasons | pending | pending | open |
| P2 ownership implementation | pending | pending | open |
| P2 validation | pending | pending | open |
| P3 access implementation | pending | pending | open |
| P3 validation | pending | pending | open |
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
