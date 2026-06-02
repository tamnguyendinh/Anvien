# Anvien Go Graph Accuracy 100% Evidence Ledger

Date: 2026-05-16

Source plan: [2026-05-16-anvien-graph-accuracy-100-plan.md](2026-05-16-anvien-graph-accuracy-100-plan.md)

This ledger records commands, artifacts, validation results, and decisions for the graph accuracy plan. Benchmark values belong in the benchmark ledger.

## Baseline Evidence

### Analyze Runs

- Status: completed.
- Date: 2026-05-16.
- Repo: `E:\Anvien`

Node/MCP command:

```powershell
anvien analyze E:\Anvien --force [redacted removed argument] --no-stats --benchmark-json .tmp\compare2-node-analyze-20260516-r2.json --benchmark-label node-mcp-r2
```

Go local command:

```powershell
.\anvien\bin\anvien.exe analyze E:\Anvien --force [redacted removed argument] --no-stats --benchmark-json .tmp\compare2-go-analyze-20260516-r2.json --benchmark-label go-local-r2
```

Artifacts:

- `.tmp\compare2-node-analyze-20260516-r2.json`
- `.tmp\compare2-go-analyze-20260516-r2.json`
- `.tmp\compare2-node-api-graph-20260516-r2.json`
- `.tmp\compare2-go-api-graph-20260516-r2.json`
- `.tmp\compare2-anvien-node-vs-go-20260516.summary.json`

Decision:

- API graph output was used for graph comparison.
- `.anvien\graph.json` was not used as comparison authority because it can reflect the most recent analyzer run rather than both analyzer outputs.

### Accuracy Probe

- Status: completed.
- Date: 2026-05-16.

Command:

```powershell
go run .tmp\accuracy_probe.go -repo . -node .tmp\compare2-node-api-graph-20260516-r2.json -go .tmp\compare2-go-api-graph-20260516-r2.json -out .tmp\graph-accuracy-node-vs-go-20260516.json
```

Output:

```text
wrote .tmp\graph-accuracy-node-vs-go-20260516.json
common_go_files=403 expected_defs=11673 expected_import_edges=3456 expected_direct_calls=5495
```

Artifacts:

- `.tmp\accuracy_probe.go`
- `.tmp\graph-accuracy-node-vs-go-20260516.json`

Ground-truth method:

- Go standard parser over the `403` Go files common to both API graphs.
- Local imports resolved through discovered `go.mod` module roots.
- Direct call ground truth limited to same-package direct function calls, dot-import function calls, and imported-package function calls.

## Baseline Miss Classification

### TypeAlias

Status: open.

Known Go local misses:

```text
TypeAlias|internal/providers/astro/extract.go|Request
TypeAlias|internal/providers/svelte/extract.go|Request
TypeAlias|internal/providers/vue/extract.go|Request
```

Source evidence:

```text
internal\providers\astro\extract.go:9:type Request = sfc.Request
internal\providers\svelte\extract.go:9:type Request = sfc.Request
internal\providers\vue\extract.go:9:type Request = sfc.Request
```

Initial diagnosis:

- Go local captures most named non-struct types.
- It misses alias declarations that use `type A = B` wrapper syntax.

### Variable

Status: open.

Sample Go local misses:

```text
Variable|internal/cli/benchmark_command.go|key
Variable|internal/cli/package_command_test.go|want
Variable|internal/embeddings/pipeline.go|index
Variable|internal/group/matching.go|link
Variable|internal/group/storage.go|entry
Variable|internal/group/topic_extractor_test.go|i
Variable|internal/httpapi/graph.go|value
Variable|internal/httpapi/panels.go|group
Variable|internal/lbugnative/native_ladybugdb.go|column
Variable|internal/providers/golang/extract_test.go|binding
```

Initial diagnosis:

- Missing coverage is concentrated in local variable binding forms.
- Expected families include short declarations, range variables, init variables, and table-test locals.

### Direct CALLS

Status: open.

Sample Go local misses:

```text
Function|internal/embeddings/search_test.go|TestSemanticSearchQueriesVectorIndexDedupsChunksAndHydratesMetadata->Function|internal/embeddings/pipeline_test.go|containsQuery
Function|internal/group/query.go|Query->Function|internal/group/storage.go|Load
Function|internal/group/status.go|Status->Function|internal/group/storage.go|Load
Function|internal/providers/python/imports.go|importNameAndAlias->Function|internal/providers/python/nodes.go|directChildOfKind
Function|internal/providers/swift/references.go|callArguments->Function|internal/providers/swift/nodes.go|directChildrenOfKind
Function|internal/providers/tsjs/nodes.go|descendantsOfType->Function|internal/providers/tsjs/extract.go|walk
Function|internal/repo/meta.go|FindIndexed->Function|internal/repo/paths.go|absClean
Function|internal/resolution/resolution_test.go|BenchmarkResolveTypeScriptGraphFixture->Function|internal/resolution/resolve.go|Resolve
Function|internal/structure/structure_test.go|TestApplyPreservesExistingFileNodeProperties->Function|internal/structure/structure.go|Apply
Method|internal/providers/c/references.go|collector.emitReference->Function|internal/providers/c/nodes.go|lastIdentifierLikeChild
```

Initial diagnosis:

- Missing coverage includes same-package cross-file helper calls.
- Test helper calls and provider helper calls need stronger package-level direct-call resolution.

## Validation Commands To Use Per Cluster

Focused tests:

```powershell
go test ./internal/providers/golang ./internal/resolution
```

Broader regression tests:

```powershell
go test ./...
```

Go local analyze:

```powershell
.\anvien\bin\anvien.exe analyze E:\Anvien --force [redacted removed argument] --no-stats --benchmark-json <artifact> --benchmark-label <label>
```

Accuracy gate report mode:

```powershell
go run ./cmd/graph-accuracy-probe -mode report -repo E:\Anvien -node .tmp\compare2-node-api-graph-20260516-r2.json -anvien .\anvien\bin\anvien.exe -fresh-go-graph <go-api-graph-json> -benchmark-json <analyze-benchmark-json> -benchmark-label <label> -out <accuracy-json>
```

Accuracy gate enforce mode:

```powershell
go run ./cmd/graph-accuracy-probe -mode enforce -repo E:\Anvien -node .tmp\compare2-node-api-graph-20260516-r2.json -go <go-api-graph-json> -out <accuracy-json>
```

## Phase Evidence Slots

### Phase 1 - Accuracy Gate Ownership

- Status: completed.
- Tracked gate owner: `internal/graphaccuracy` plus `cmd/graph-accuracy-probe`.
- Implementation files:
  - `internal/graphaccuracy/graphaccuracy.go`
  - `internal/graphaccuracy/graphaccuracy_test.go`
  - `cmd/graph-accuracy-probe/main.go`
- Report-mode command:

```powershell
go run ./cmd/graph-accuracy-probe -mode report -repo E:\Anvien -node .tmp\compare2-node-api-graph-20260516-r2.json -anvien .\anvien\bin\anvien.exe -fresh-go-graph .tmp\p1-go-api-graph-20260516.json -benchmark-json .tmp\p1-go-analyze-20260516.json -benchmark-label p1-tracked-gate-baseline -out .tmp\p1-tracked-gate-baseline-20260516.json
```

- Report-mode result:

```text
files: scanned=676 parsed=520 unsupported=156 failed=0
graph: nodes=19009 relationships=45385 path=E:\Anvien\.anvien\graph.json
common_go_files=403 expected_defs=11673 expected_import_edges=3456 expected_direct_calls=5495
definition.TypeAlias.goLocal=41/44 recall=93.18 precision=100.00 graphCandidates=41
definition.Variable.goLocal=6223/7074 recall=87.97 precision=100.00 graphCandidates=6223
imports.goLocal=3456/3456 recall=100.00 precision=100.00 graphCandidates=3456
calls.goLocal=4696/5495 recall=85.46 graphCandidates=6956
accuracy gate: 3 failure(s)
```

- Enforce-mode command:

```powershell
go run ./cmd/graph-accuracy-probe -mode enforce -repo E:\Anvien -node .tmp\compare2-node-api-graph-20260516-r2.json -go .tmp\p1-go-api-graph-20260516.json -out .tmp\p1-tracked-gate-enforce-20260516.json
```

- Enforce-mode result: exited `1` while the known below-target gates remain open.
- Artifacts:
  - `.tmp\p1-go-analyze-20260516.json`
  - `.tmp\p1-go-api-graph-20260516.json`
  - `.tmp\p1-tracked-gate-baseline-20260516.json`
  - `.tmp\p1-tracked-gate-enforce-20260516.json`
- Full build:

```powershell
.\anvien-launcher\build.ps1
```

Result: passed after the final P1 code change. This built the Web UI, Go backend runtime, launcher, and server wrapper.

- Package runtime build check:

```powershell
Push-Location anvien; npm run build; Pop-Location
```

Result: failed while copying `anvien\bin\lbug_shared.dll` because four linked `anvien mcp` processes were holding the DLL open. `anvien\bin\anvien.exe` was updated before the DLL copy step; the DLL lock is environmental and not caused by the P1 gate code.

- Focused tests:

```powershell
go test ./internal/graphaccuracy ./cmd/graph-accuracy-probe ./internal/providers/golang ./internal/resolution
```

Result: passed.

- Broader regression test:

```powershell
go test ./...
```

Result: failed only in fixture packages that are intentionally not standalone-buildable under `go test ./...`:

```text
anvien/test/fixtures/lang-resolution/go-map-range: package models is not in std
anvien/test/fixtures/lang-resolution/go-method-enrichment: package animal is not in std; mixed packages animal/main
anvien/test/fixtures/sample-code: C source files not allowed when not using cgo or SWIG
anvien/test/fixtures/lang-resolution/go-make-builtin: cannot call pointer method Greet on User
anvien/test/fixtures/lang-resolution/go-type-assertion: impossible type assertion
```

All non-fixture Go packages reported `ok` in that run.

- CLI/MCP runtime e2e:

```powershell
.\anvien\bin\anvien.exe status
.\anvien\bin\anvien.exe context Run --repo E:\Anvien --file internal\analyze\analyze.go
```

Result: status reported the repo up to date at the fresh P1 analyze, and `context` returned the `internal/analyze/analyze.go:Run` symbol with incoming/outgoing graph relationships and processes.

- Anvien pre-commit refresh:

```powershell
.\anvien\bin\anvien.exe analyze E:\Anvien --force [redacted removed argument] --no-stats
```

Result: passed; graph refreshed before staged `detect-changes`.

- Detect changes:

```powershell
.\anvien\bin\anvien.exe detect-changes --scope staged --repo E:\Anvien
```

Result summary:

```text
changed_files=6
changed_count=246
affected_count=12
risk_level=high
```

Representative affected processes were the new tracked gate flows such as `Main -> GraphFacts`, `Main -> PackageInfo`, `Main -> ParseGraphPath`, `Main -> DefinitionLabels`, and `Main -> WriteResult`.

- Decision: P1 gate ownership is complete. The tracked gate matches the old baseline for the measured gates while ignoring Go graph candidates outside the `403` common baseline files, so new gate source files do not contaminate the baseline comparison.
- Implementation commit: `288c093`.

### Phase 2 - TypeAlias Fix

- Status: completed.
- Code changes:
  - `internal/providers/golang/definitions.go`: handle tree-sitter `type_alias` nodes in context collection and definition emission.
  - `internal/providers/golang/types.go`: emit type references for `type_alias` nodes.
  - `internal/providers/golang/extract_test.go`: add `TestExtractGoTypeAliasDeclarations`.
- Impact checks:

```powershell
.\anvien\bin\anvien.exe context emitTypeBindingKind --repo E:\Anvien --file internal\providers\golang\types.go
.\anvien\bin\anvien.exe context emitValueSpecDefinitions --repo E:\Anvien --file internal\providers\golang\definitions.go
.\anvien\bin\anvien.exe impact --uid "Function:internal/providers/golang/definitions.go:goTypeSpecLabel#1" --repo E:\Anvien --direction upstream --depth 3 --include-tests
```

Result: direct impact stayed inside the Go provider module.

- Full build:

```powershell
.\anvien-launcher\build.ps1
```

Result: passed.

- Canonical package runtime build attempt:

```powershell
Push-Location anvien; npm run build; $code=$LASTEXITCODE; Pop-Location; Write-Output "npm_build_exit=$code"
```

Result: failed with `npm_build_exit=1` because active linked `anvien mcp` processes held `anvien\bin\anvien.exe` open. The P2 validation therefore used the freshly built runtime at `anvien-launcher\server-bundle\anvien.exe`, which was produced by the successful full build.

- Focused tests:

```powershell
go test ./internal/providers/golang -run TestExtractGoTypeAliasDeclarations -count=1 -v
go test ./internal/providers/golang ./internal/resolution
```

Result: passed.

- Broader regression test:

```powershell
go test ./...
```

Result: failed only in the same fixture packages that are intentionally not standalone-buildable under `go test ./...`; all non-fixture Go packages reported `ok`.

- Fresh analyze and accuracy gate:

```powershell
go run ./cmd/graph-accuracy-probe -mode report -repo E:\Anvien -node .tmp\compare2-node-api-graph-20260516-r2.json -anvien .\anvien-launcher\server-bundle\anvien.exe -fresh-go-graph .tmp\p2-go-api-graph-20260516.json -benchmark-json .tmp\p2-go-analyze-20260516.json -benchmark-label p2-typealias-fix -out .tmp\p2-typealias-accuracy-20260516.json
```

Result:

```text
files: scanned=676 parsed=520 unsupported=156 failed=0
graph: nodes=19024 relationships=45410 path=E:\Anvien\.anvien\graph.json
definition.TypeAlias.goLocal=44/44 recall=100.00 precision=100.00 graphCandidates=44
definition.Variable.goLocal=6225/7076 recall=87.97 precision=100.00 graphCandidates=6225
calls.goLocal=4700/5499 recall=85.47 graphCandidates=6960
accuracy gate: 2 failure(s)
```

- CLI/MCP runtime e2e:

```powershell
.\anvien-launcher\server-bundle\anvien.exe status
.\anvien-launcher\server-bundle\anvien.exe context Request --repo E:\Anvien --file internal\providers\astro\extract.go
```

Result: status reported the repo up to date, and `context Request` returned `kind: "TypeAlias"` for `internal/providers/astro/extract.go`.

- Anvien pre-commit refresh:

```powershell
.\anvien-launcher\server-bundle\anvien.exe analyze E:\Anvien --force [redacted removed argument] --no-stats
```

Result: passed; graph refreshed with the P2 runtime before staged `detect-changes`.

- Detect changes:

```powershell
.\anvien-launcher\server-bundle\anvien.exe detect-changes --scope staged --repo E:\Anvien
```

Result summary:

```text
changed_files=6
changed_count=19
affected_count=0
risk_level=low
```

- Artifacts:
  - `.tmp\p2-go-analyze-20260516.json`
  - `.tmp\p2-go-api-graph-20260516.json`
  - `.tmp\p2-typealias-accuracy-20260516.json`
- Decision: P2 is complete. The TypeAlias gate is at `100.00%`; remaining open gates are Variable and direct `CALLS`.
- Implementation commit: `c6d2921`.

### Phase 3 - Variable Fix

- Status: completed.
- Code changes:
  - `internal/providers/golang/definitions.go`: emit `Variable` definitions for Go `range_clause` declarations, `type_switch_statement` aliases, and `receive_statement` aliases; skip assignment-only range loops and blank identifiers.
  - `internal/providers/golang/extract_test.go`: add focused coverage for range variables, nested range variables, assignment-only range loops, type-switch aliases, and select receive aliases.
- Miss classification:
  - Primary miss family: `range_clause` declarations, including value-only `for _, name := range ...`, key-only `for key := range map`, and key/value `for key, value := range ...`. Examples included `node`, `relationship`, `skill`, `dd`, `line`, `link`, `process`, `repoPath`, `filePath`, `index`, `leftValue`, and `relType`.
  - Remaining misses after the range fix: type-switch aliases such as `typed`, `raw`, and `value`, plus receive aliases `event` and `ok`.
  - Existing `short_var_declaration` coverage already handled ordinary local helper variables and table-test locals; the extractor changes focused on the missing AST families instead of rewriting working short-var behavior.
- Impact checks:

```powershell
.\anvien-launcher\server-bundle\anvien.exe context emitShortVarDefinitions --repo E:\Anvien --file internal\providers\golang\definitions.go
.\anvien-launcher\server-bundle\anvien.exe impact --uid "Method:internal/providers/golang/definitions.go:collector.emitShortVarDefinitions#1" --repo E:\Anvien --direction upstream --depth 2 --include-tests
```

Result: direct impact stayed inside the Go provider extraction path; no affected runtime process was reported for the staged variable extractor change.

- Full build:

```powershell
.\anvien-launcher\build.ps1
```

Result: passed. This built the Web UI, Go backend runtime, launcher, and server wrapper before validation tests.

- Focused tests:

```powershell
go test ./internal/providers/golang -run "TestExtractGo(Range|SwitchAndReceive)VariableDeclarations" -count=1 -v
go test ./internal/providers/golang ./internal/resolution
```

Result: passed.

- Broader regression test:

```powershell
go test ./...
```

Result: failed only in the same fixture packages that are intentionally not standalone-buildable under `go test ./...`; all non-fixture Go packages reported `ok`.

```text
anvien/test/fixtures/lang-resolution/go-map-range: package models is not in std
anvien/test/fixtures/lang-resolution/go-method-enrichment: package animal is not in std; mixed packages animal/main
anvien/test/fixtures/sample-code: C source files not allowed when not using cgo or SWIG
anvien/test/fixtures/lang-resolution/go-make-builtin: cannot call pointer method Greet on User
anvien/test/fixtures/lang-resolution/go-type-assertion: impossible type assertion
```

- Fresh analyze and accuracy gate:

```powershell
go run ./cmd/graph-accuracy-probe -mode report -repo E:\Anvien -node .tmp\compare2-node-api-graph-20260516-r2.json -anvien .\anvien-launcher\server-bundle\anvien.exe -fresh-go-graph .tmp\p3-go-api-graph-20260516-final.json -benchmark-json .tmp\p3-go-analyze-20260516-final.json -benchmark-label p3-variable-fix -out .tmp\p3-variable-accuracy-20260516-final.json
```

Result:

```text
files: scanned=676 parsed=520 unsupported=156 failed=0
graph: nodes=19916 relationships=46367 path=E:\Anvien\.anvien\graph.json
definition.TypeAlias.goLocal=44/44 recall=100.00 precision=100.00 graphCandidates=44
definition.Variable.goLocal=7080/7080 recall=100.00 precision=100.00 graphCandidates=7080
imports.goLocal=3456/3456 recall=100.00 precision=100.00 graphCandidates=3456
calls.goLocal=4711/5514 recall=85.44 graphCandidates=6972
accuracy gate: 1 failure(s)
```

- CLI/MCP runtime e2e:

```powershell
.\anvien-launcher\server-bundle\anvien.exe context emitRangeClauseDefinitions --repo E:\Anvien --file internal\providers\golang\definitions.go
.\anvien-launcher\server-bundle\anvien.exe context emitTypeSwitchStatementDefinitions --repo E:\Anvien --file internal\providers\golang\definitions.go
```

Result: both symbols were found in the built Go runtime graph. `context` returned the new methods with outgoing `CALLS` to `definitionNameNodes`, `rangeClauseDefines`, and `definesBefore`, plus `USES` evidence for `sitter.Node`.

- Artifacts:
  - `.tmp\p3-go-analyze-20260516-final.json`
  - `.tmp\p3-go-api-graph-20260516-final.json`
  - `.tmp\p3-variable-accuracy-20260516-final.json`
- Anvien pre-commit refresh:

```powershell
.\anvien-launcher\server-bundle\anvien.exe analyze E:\Anvien --force [redacted removed argument] --no-stats
```

Result: passed; graph refreshed with the P3 runtime before staged `detect-changes`.

- Detect changes:

```powershell
.\anvien-launcher\server-bundle\anvien.exe detect-changes --scope staged --repo E:\Anvien
```

Result summary:

```text
changed_files=5
changed_count=26
affected_count=0
risk_level=low
```

- Decision: P3 is complete. The Variable gate is at `100.00%`; remaining open gate is direct `CALLS`.
- Implementation commit: `cffce45`.

### Phase 4 - Direct CALLS Fix

- Status: completed.
- Code changes:
  - `internal/resolution/indexes.go`: track each file's language in the resolution workspace and add a Go-only same-package function lookup for direct calls.
  - `internal/resolution/resolve.go`: try the Go same-package function lookup after lexical/same-file resolution and before global fallback.
  - `internal/resolution/resolution_test.go`: add same-package cross-file direct-call coverage that avoids broad global ambiguity and covers variadic metadata mismatch.
- Miss classification:
  - Full classification artifact: `.tmp\p4-calls-full-missing-20260516.json`.
  - Total current misses before the P4 fix: `803`.
  - `803 / 803` misses were same-directory Go cross-file direct function calls.
  - `0` imported-package direct-call misses and `0` dot-import direct-call misses were present in the current artifact.
  - Source/target labels: `463` `Method -> Function`, `340` `Function -> Function`.
  - Test split: `632` production-to-production, `142` test-to-production, `29` test-to-test.
  - Largest groups: provider package helpers `658`, resolution helpers `51`, enrichment `Apply` calls `36`, analyze `Run` / test helpers `22`, package `Load` / storage helpers `11`, repo path helpers `10`, misc same-package helpers `15`.
- Representative misses closed:

```text
Method|internal/providers/golang/definitions.go|collector.addReturnType->Function|internal/providers/golang/nodes.go|child
Function|internal/providers/java/definitions.go|formalParameters->Function|internal/providers/java/nodes.go|directChildOfKind
Function|internal/providers/dart/scopes.go|callableRange->Function|internal/providers/dart/nodes.go|nodeRangeWithEnd
Function|internal/resolution/emit.go|emitImportEdges->Function|internal/resolution/indexes.go|cleanPath
Function|internal/analyze/analyze_test.go|TestRunRejectsConcurrentWriterLock->Function|internal/analyze/analyze.go|Run
Function|internal/group/config_test.go|TestLoadReadsGroupConfigFromDisk->Function|internal/group/storage.go|Load
Function|internal/cobol/jcl.go|parseJCL->Function|internal/cobol/cobol.go|firstNonEmpty
```

- Impact checks:

```powershell
.\anvien-launcher\server-bundle\anvien.exe context resolveCall --repo E:\Anvien --file internal\resolution\resolve.go
.\anvien-launcher\server-bundle\anvien.exe context resolveGlobalCallName --repo E:\Anvien --file internal\resolution\indexes.go
.\anvien-launcher\server-bundle\anvien.exe impact --uid "Function:internal/resolution/resolve.go:resolveCall#3" --repo E:\Anvien --direction upstream --depth 2 --include-tests
```

Result: `resolveCall` impact is broad by design (`CRITICAL`) because it is resolution core; the implementation keeps the new branch Go-only, same-directory, function-only, and after lexical/same-file resolution to avoid changing receiver dispatch or non-Go behavior.

- Full build:

```powershell
.\anvien-launcher\build.ps1
```

Result: passed after the final P4 code change. This built the Web UI, Go backend runtime, launcher, and server wrapper before validation tests.

- Focused tests:

```powershell
go test ./internal/resolution -run "TestResolveGoSamePackageDirectCallAcrossFilesBeforeGlobalAmbiguity|TestResolveImportedPackageMemberCall|TestResolveGlobalCallFallbackUsesArityToAvoidAmbiguity" -count=1 -v
go test ./internal/resolution ./internal/providers/golang ./internal/graphaccuracy ./cmd/graph-accuracy-probe
```

Result: passed.

- Broader regression test:

```powershell
go test ./...
```

Result: failed only in the same fixture packages that are intentionally not standalone-buildable under `go test ./...`; all non-fixture Go packages reported `ok`.

```text
anvien/test/fixtures/lang-resolution/go-map-range: package models is not in std
anvien/test/fixtures/lang-resolution/go-method-enrichment: package animal is not in std; mixed packages animal/main
anvien/test/fixtures/sample-code: C source files not allowed when not using cgo or SWIG
anvien/test/fixtures/lang-resolution/go-make-builtin: cannot call pointer method Greet on User
anvien/test/fixtures/lang-resolution/go-type-assertion: impossible type assertion
```

- Fresh analyze and accuracy gate:

```powershell
go run ./cmd/graph-accuracy-probe -mode report -repo E:\Anvien -node .tmp\compare2-node-api-graph-20260516-r2.json -anvien .\anvien-launcher\server-bundle\anvien.exe -fresh-go-graph .tmp\p4-go-api-graph-20260516-final.json -benchmark-json .tmp\p4-go-analyze-20260516-final.json -benchmark-label p4-direct-calls-fix-final -out .tmp\p4-calls-accuracy-20260516-final.json -max-examples 10000
```

Result:

```text
files: scanned=676 parsed=520 unsupported=156 failed=0
graph: nodes=19984 relationships=47622 path=E:\Anvien\.anvien\graph.json
definition.Function.goLocal=2964/2964 recall=100.00 precision=100.00 graphCandidates=2964
definition.Method.goLocal=791/791 recall=100.00 precision=100.00 graphCandidates=791
definition.Struct.goLocal=472/472 recall=100.00 precision=100.00 graphCandidates=472
definition.Interface.goLocal=21/21 recall=100.00 precision=100.00 graphCandidates=21
definition.TypeAlias.goLocal=44/44 recall=100.00 precision=100.00 graphCandidates=44
definition.Const.goLocal=321/321 recall=100.00 precision=100.00 graphCandidates=321
definition.Variable.goLocal=7088/7088 recall=100.00 precision=100.00 graphCandidates=7088
imports.goLocal=3456/3456 recall=100.00 precision=100.00 graphCandidates=3456
calls.goLocal=5520/5520 recall=100.00 graphCandidates=7791
accuracy gate: PASS
```

- CLI/MCP runtime e2e:

```powershell
.\anvien-launcher\server-bundle\anvien.exe context parseJCL --repo E:\Anvien --file internal\cobol\jcl.go
.\anvien-launcher\server-bundle\anvien.exe context firstNonEmpty --repo E:\Anvien --file internal\cobol\cobol.go
```

Result: `context parseJCL` returned outgoing `CALLS` to `firstNonEmpty` with confidence `0.95`; `context firstNonEmpty` returned incoming calls from both `processJCL` and `parseJCL`.

- Artifacts:
  - `.tmp\p4-calls-full-missing-20260516.json`
  - `.tmp\p4-go-analyze-20260516-final.json`
  - `.tmp\p4-go-api-graph-20260516-final.json`
  - `.tmp\p4-calls-accuracy-20260516-final.json`
- Anvien pre-commit refresh:

```powershell
.\anvien-launcher\server-bundle\anvien.exe analyze E:\Anvien --force [redacted removed argument] --no-stats
```

Result: passed; graph refreshed with the P4 runtime before staged `detect-changes`.

- Detect changes:

```powershell
.\anvien-launcher\server-bundle\anvien.exe detect-changes --scope staged --repo E:\Anvien
```

Result summary:

```text
changed_files=6
changed_count=27
affected_count=10
risk_level=high
```

Representative affected processes were expected resolution-core flows such as `BuildWorkspace -> GenerateID`, `BuildWorkspace -> CleanPath`, `ResolveCall -> CleanPath`, `ResolveCall -> DefRef`, and `ResolveCall -> DefinitionFact`.

- Decision: P4 is complete. All measured Go local gates are now at `100.00%`; continue to final cutover.
- Implementation commit: `7cefe83`.

### Phase 5 - Final Accuracy Cutover

- Status: completed.
- Tracked accuracy gate location: `internal/graphaccuracy` plus `cmd/graph-accuracy-probe`.
- Final analyze artifact: `.tmp\p5-final-go-analyze-20260516.json`.
- Final API graph artifact: `.tmp\p5-final-go-api-graph-20260516.json`.
- Final accuracy artifact: `.tmp\p5-final-accuracy-20260516.json`.
- Final validation commit before doc-only close: `70fcb1c`.
- Full build:

```powershell
.\anvien-launcher\build.ps1
```

Result: passed. This built the Web UI, Go backend runtime, launcher, and server wrapper before final focused tests and final enforce gate.

- Focused tests:

```powershell
go test ./internal/resolution ./internal/providers/golang ./internal/graphaccuracy ./cmd/graph-accuracy-probe
```

Result: passed.

- Final enforce gate:

```powershell
go run ./cmd/graph-accuracy-probe -mode enforce -repo E:\Anvien -node .tmp\compare2-node-api-graph-20260516-r2.json -anvien .\anvien-launcher\server-bundle\anvien.exe -fresh-go-graph .tmp\p5-final-go-api-graph-20260516.json -benchmark-json .tmp\p5-final-go-analyze-20260516.json -benchmark-label p5-final-accuracy -out .tmp\p5-final-accuracy-20260516.json -max-examples 10000
```

Result:

```text
files: scanned=676 parsed=520 unsupported=156 failed=0
graph: nodes=19984 relationships=47622 path=E:\Anvien\.anvien\graph.json
definition.Function.goLocal=2964/2964 recall=100.00 precision=100.00 graphCandidates=2964
definition.Method.goLocal=791/791 recall=100.00 precision=100.00 graphCandidates=791
definition.Struct.goLocal=472/472 recall=100.00 precision=100.00 graphCandidates=472
definition.Interface.goLocal=21/21 recall=100.00 precision=100.00 graphCandidates=21
definition.TypeAlias.goLocal=44/44 recall=100.00 precision=100.00 graphCandidates=44
definition.Const.goLocal=321/321 recall=100.00 precision=100.00 graphCandidates=321
definition.Variable.goLocal=7088/7088 recall=100.00 precision=100.00 graphCandidates=7088
imports.goLocal=3456/3456 recall=100.00 precision=100.00 graphCandidates=3456
calls.goLocal=5520/5520 recall=100.00 graphCandidates=7791
accuracy gate: PASS
```

- CLI/MCP runtime e2e:

```powershell
.\anvien-launcher\server-bundle\anvien.exe status
.\anvien-launcher\server-bundle\anvien.exe context resolveGoSamePackageFunction --repo E:\Anvien --file internal\resolution\indexes.go
```

Result: status reported repo indexed and up to date at commit `70fcb1c`; `context resolveGoSamePackageFunction` returned the symbol, incoming `CALLS` from `resolveCall`, and outgoing relationships to `cleanPath`, `definitionLookupNameMatches`, and `uniqueDefAccumulator` helpers.

- Broad regression carried forward from P4:

```powershell
go test ./...
```

Result: all non-fixture Go packages reported `ok`; the only failures were the known intentionally non-standalone fixture packages recorded in P4 evidence.

- Detect changes:

P5 has no implementation diff after P4; the final code-impact detect output is the P4 staged detect-changes result:

```text
changed_files=6
changed_count=27
affected_count=10
risk_level=high
```

No additional Anvien run was made solely for the doc-only close commit, per rule 6.

- Decision: plan complete. Benchmark and evidence ledgers agree with `.tmp\p5-final-accuracy-20260516.json`, and every measured Go local gate is `100.00%`.
- Final evidence commit: `3bca698`.
