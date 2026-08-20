# Child 04 / P4-B Coder Handoff

## Candidate State

- State: `READY_FOR_SUPERVISOR` (resubmission after `reports/Supervisor/rp_supervisor_260821_001125_by_gpt-5_child04_p4b_export_facts.md` REVIEW1 REJECT).
- Scope: Child 04 P4-B only — truthful first-class ScopeIR export facts for direct/default/local alias/type-only TypeScript and JavaScript syntax.
- This is Coder evidence, not a Supervisor acceptance verdict.
- No stage, commit, push, final detect-changes, ledger edit, roadmap edit, graph/projection/persistence edit, Child 05 work, or target access was performed.
- Current HEAD before handoff: `479e8ac229a17f2f6f94be9a4d04e07d74ac4d43` (accepted P4-A boundary).

## Exact Boundary

Production files changed:

- `internal/providers/tsjs/imports.go`
- `internal/providers/tsjs/extract.go`

Tests changed only after production behavior:

- `internal/providers/tsjs/extract_test.go`

No `internal/scopeir`, graph, projection, persistence, definitions, Child 05, target, roadmap, or ledger bytes were changed. The pre-existing untracked Main-owned handoff `reports/Investigation/rp_main_260820_225428_orchestration_rotation_handoff.md` was preserved unchanged.

## Invariant Family Map

| Surface | Authority / invariant | Result | Evidence |
|---|---|---|---|
| Direct/default export extraction | Child 04 P4-B requirements; one fact per eligible binding site | Emits `ExportDirect` / `ExportDefault` for variable declarators, functions, generators, classes, enums, interfaces, type aliases, namespaces, ambient declarations, named/anonymous default expressions | `E4-P4B-SRC1`, `E4-P4B-TEST1` |
| Local export clauses and aliases | One fact per export specifier; exported/local names stay distinct | Emits named, alias, default-alias, `export type`, and inline `type` specifier facts with exact range/selection/provenance and local definition evidence when unique | `E4-P4B-SRC1`, `E4-P4B-TEST1` |
| Meaning/type-only lane | Export meaning is independent from access visibility | Preserves value/type/namespace meanings and explicit `TypeOnly`; dual-meaning class/enum facts retain both lanes; `DefinitionFact.Visibility` remains untouched | `E4-P4B-SRC1`, `E4-P4B-BOUNDARY1` |
| Unsupported/malformed syntax | Hidden fallback is forbidden | Emits structured `ExportDiagnosticFact` for unsupported or malformed source-less direct/local forms; valid sibling sites in mixed statements remain countable | `E4-P4B-SRC1`, `E4-P4B-TEST1` |
| Source-bearing re-export compatibility | P4-B must preserve P4-B1/Child 05 boundary | Existing `ImportReexport` and `ImportWildcard` behavior remains; namespace-star compatibility wildcard is preserved; no first-class re-export/terminal fact is emitted in P4-B | `E4-P4B-TEST1` |
| Extraction dispatch | `Extract` walk dispatches export collection and flushes facts at ScopeIR result boundary | Thin dispatch only adds export state to the existing collector and calls `emitPendingExportFacts` before normalization | `E4-P4B-SRC1`, `E4-P4B-BOUNDARY1` |

Sibling surfaces checked and preserved:

- `internal/providers/tsjs/definitions.go`: definition extraction and visibility are preserve-only.
- `internal/scopeir/*`: accepted P4-A contract consumed without redesign.
- `internal/resolution`, `internal/lbugload`, `internal/lbugschema`, graph/projection, compatibility writers, and semantic readers: preserve-only; P4-C remains locked.
- Child 05 terminal resolution, barrel traversal, ambiguity/cycle handling, and public-API reachability: untouched.
- `E:\cheapapp.org`: not accessed.

Legacy fallback status: no `Visibility="exported"` writer, no independent compatibility export truth, no terminal resolution state, and no silent drop for unsupported source-less export syntax. Source-bearing malformed/re-export semantics remain outside P4-B.

## `E4-P4B-IMPACT1` — Fresh Graph / File Detail / Upstream Impact

Fresh graph before the final evidence pass:

```text
anvien analyze . --force --exclude 'internal/aicontext/skills/**' --exclude '.claude/skills/**' --json
exit 0; scanned=1,131; parsed=626; failed=0; graph=81,722 nodes / 121,234 relationships
```

Final file-detail summaries (`stale=false`, `changedSinceAnalyze=false`):

| File | Symbols | Inbound / outbound / local | Unresolved | Risk |
|---|---:|---:|---:|---|
| `internal/providers/tsjs/imports.go` | 129 | 7 / 82 / 27 | 353 | HIGH |
| `internal/providers/tsjs/extract.go` | 39 | 24 / 35 / 32 | 71 | HIGH |

The HIGH file risks are blast-radius warnings. Exact edited symbols (`emitImportKind`, `emitExportStatement`, `emitPendingExportFacts`, `emitExportFacts`, `emitExportDeclaration`, `emitVariableExportDeclarations`, `emitDefaultExportExpression`, `emitLocalExportClause`, `Extract`, and `collector.result`) each report upstream impact risk `LOW` with `0` direct affected files/modules/processes. Linked provider flows/tests remain visible in file-detail and were covered by the nearest package regression.

## `E4-P4B-SRC1` — Production Source Result

- `emitExportStatement` now queues source-less top-level export statements for first-class extraction while preserving source-bearing compatibility imports.
- `emitExportFacts` dispatches direct declarations, default expressions, local clauses, and structured diagnostics.
- Direct/default extraction emits one fact per eligible variable declarator or binding-pattern leaf and preserves exact declaration/selection ranges, file/hash, local definition ID when unique, meaning lanes, type-only state, and statement provenance.
- Local clauses emit one fact per specifier, including aliases, `as default`, `export type`, and inline `type` forms. Ambiguous merged declarations intentionally leave `LocalDefID` empty.
- Unsupported/malformed sites emit countable diagnostics; valid siblings continue to emit facts.
- `extract.go` contains only the thin collector fields/result wiring and dispatch flush required for the provider boundary.
- Existing wildcard compatibility is preserved for both `export * from ...` and `export * as ns from ...`; no P4-B1 first-class re-export state is introduced.
- `DefinitionFact.Visibility` and all access extraction behavior are unchanged.

## `E4-P4B-BUILD1` — Build / Holder Gate

Holder gate:

- `anvien doctor locks --repo E:\Anvien --json`: `status=free`, lock absent/not alive.
- `anvien doctor processes --json`: only editor-owned MCP/runtime helpers; no build-related holder required termination.

Canonical build:

```text
npm run full-build
exit 0
```

The build completed npm/runtime packaging, Go launcher build, web Vite production build (`2,943` modules transformed), CLI version `1.2.8`, and final self-analyze. Warnings about mixed dynamic/static import and large chunks were non-failing.

## `E4-P4B-TEST1` — Focused Tests

```text
go test ./internal/providers/tsjs -count=1
PASS (0.206s)
```

Direct tests in `internal/providers/tsjs/extract_test.go` cover:

- `TestExtractTypeScriptDirectAndDefaultExportFacts` — direct/default declarations, multi-declarator bindings, ambient forms, dual meanings, negative unexported control.
- `TestExtractTypeScriptLocalAliasAndTypeOnlyExportFacts` — local aliases, default aliases, type-only statement/specifier forms, unique/ambiguous local definition evidence, nested negative control.
- `TestExtractJavaScriptDirectDefaultAndLocalExportFacts` — JavaScript parity for direct/default/alias forms.
- `TestExtractExportDiagnosticsAndLaterSliceBoundaries` — unsupported/malformed diagnostics, mixed valid/invalid sibling sites, and preserved source-bearing re-export compatibility.

## `E4-P4B-BOUNDARY1` — Nearest Real Boundary

```text
go test ./internal/providers/tsjs ./internal/scopeir ./internal/providers -count=1
PASS
```

All three packages passed. The observable boundary is:

```text
TypeScript/JavaScript source
  -> tsjs.Extract
  -> ScopeIR.Exports / ScopeIR.ExportDiagnostics
  -> owned normalization
```

The focused matrix proves one fact per eligible direct/local/default site, explicit meaning/type-only state, preserved definition visibility, structured unsupported diagnostics, and unchanged source-bearing compatibility imports. P4-B does not claim graph projection, persistence parity, terminal resolution, or the locked `21/21` target validation (P4-C/P4-C2).

## Cleanup / Residuals / Handoff

- REVIEW1 finding closure: the exact empty directory `E:\Anvien\.tmp\p4b_ast_probe` was verified (`exists=true`, `childCount=0`) and then deleted with an exact non-recursive directory operation. Fresh verification now reports `artifactExists=false`; the parent `.tmp` directory remains intact.
- Fresh `.tmp` probe census reports `0` directories matching `p4b` or `ast_probe`; no new debug/probe artifact was created.
- Repo-local debug probe `.tmp/p4b_ast_probe/main.go` and its containing directory are absent before resubmission.
- `git diff --check` passes for the changed files.
- No temporary directory was created outside `E:\Anvien`.
- Residual unverified surfaces inside P4-B: `none`.
- Intentionally locked later surfaces: P4-B1, P4-C, P4-C2, Child 05, graph/projection/persistence, and target validation.
- Blocker: none known inside the authorized P4-B boundary.

## REVIEW1 Resubmission Evidence

- Fresh filesystem cleanup: `Test-Path E:\Anvien\.tmp\p4b_ast_probe` -> `False`; `.tmp` parent -> present; `Get-ChildItem E:\Anvien\.tmp -Directory` probe census -> `0` matching directories.
- Fresh status: `git status --short --untracked-files=all` still contains the same three candidate paths (`internal/providers/tsjs/imports.go`, `extract.go`, `extract_test.go`) plus pre-existing report/note/handoff artifacts; no production/test path was added or removed by cleanup.
- Fresh diff check: `git diff --check` exit `0`.
- Fresh nearest boundary: `go test ./internal/providers/tsjs ./internal/scopeir ./internal/providers -count=1` exit `0`; all three packages PASS.
- No stage, commit, push, detect-changes, P4-B1/P4-C/P4-C2, Child 05, graph/projection/persistence, or target action was performed.

## Required Next Action

Main should open an independent Supervisor re-review against this exact worktree and updated report, then own any ledger refresh, final detect-changes, staging, isolated commit, and no-push closure.
