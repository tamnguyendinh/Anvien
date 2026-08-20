# Child 04 / P4-B1 Coder Handoff — TypeScript/JavaScript Re-export Syntax Facts

## Candidate State

- State: `READY_FOR_SUPERVISOR`.
- Scope: Child 04 `P4-B1` only — immutable source-bearing named/default/star/namespace/type-only re-export syntax facts at the `tsjs.Extract -> ScopeIR` boundary.
- This report is Coder evidence, not a Supervisor verdict and not a claim that P4-B1 is accepted or committed.
- Baseline/current HEAD: `11a37aa8ec0320dd93258c058b088d1070aa778d` (`feat(tsjs): extract direct export facts`).
- No `detect-changes`, stage, commit, push, graph/persistence projection, target validation, Child 05 terminal resolution, or `E:\cheapapp.org` access was performed.
- The pre-existing Main-owned untracked report `reports/Investigation/rp_main_260821_0044_orchestration_rotation_handoff.md` was preserved unchanged.

## Exact Boundary

Production changed:

- `internal/providers/tsjs/imports.go`

Tests changed only after production behavior compiled:

- `internal/providers/tsjs/extract_test.go`

Inspected and preserved:

- `internal/providers/tsjs/extract.go` — existing P4-B collection/result wiring is sufficient and has zero diff.
- `internal/providers/tsjs/definitions.go` — definition/access visibility behavior has zero diff.
- `internal/scopeir/{facts.go,kinds.go,ir.go,sort_keys.go}` — accepted P4-A contract consumed without redesign.
- `internal/resolution`, `internal/analyze`, graph, Ladybug, persistence, CLI/MCP/HTTP/Web, and Child 05 owners — no production diff.

Current source/test diff: `617` insertions / `24` deletions across the exact two files (`imports.go` `369/22`; `extract_test.go` `248/2`).

## Invariant Family Map

| Surface | Authority / SSOT | P4-B1 result | Verification |
|---|---|---|---|
| Source-bearing named/default re-export | P4-A `scopeir.ExportFact`; Child 04 P4-B1 plan | One `ExportReexport` per eligible `export_specifier`; exported name and source-side `TargetExportedName` stay distinct | TS/JS exact-field matrices |
| Star re-export | P4-B1 `export *` requirement | One `ExportStar` per star token; `Range`/`SelectionRange` identify `*`; `TargetRaw` is syntax-only module text | TS/JS star cases |
| Namespace re-export | P4-B1 `export * as ns` requirement | One `ExportNamespace` per namespace site; exported alias, namespace meaning, exact site/selection ranges | TS/JS namespace and default/string alias cases |
| Type-only re-export | Child 04 type-only invariant | Statement-level `export type {}`, inline `type`/`typeof`, `export type *`, and type-only namespace syntax emit `TypeOnly=true` with type meaning and no value meaning | Eight explicit TS type-only sites |
| Compatibility ImportFact | Existing P4-B compatibility boundary; no second export truth | `ImportReexport`/`ImportWildcard` are derived from each accepted source export fact | `24/24` compatibility matches, resolver/analyze regression |
| Unsupported/malformed source-bearing syntax | Hidden fallback forbidden | Missing source and malformed specifier sites emit structured, countable `ExportDiagnosticFact`; valid siblings still emit facts | Three missing-source cases + one mixed-specifier case |
| Access visibility | P4-A access/export separation | `DefinitionFact.Visibility` and definition owners unchanged | zero production diff outside `imports.go`; P4-B direct tests PASS |
| Child 05 terminal boundary | Child 04/05 ownership split | No target file/definition, link status, barrel reachability, ambiguity, cycle, or public-API state | forbidden-field assertion `0/7`; source scan |

Authority / SSOT:

- `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-04-typescript-export-semantics/` four ledgers and campaign roadmap.
- Accepted P4-A commit `479e8ac229a17f2f6f94be9a4d04e07d74ac4d43` and Supervisor PASS report.
- Accepted P4-B commit `11a37aa8ec0320dd93258c058b088d1070aa778d` and Supervisor resubmission PASS report.

Sibling surfaces checked:

- P4-B direct/default/local alias/type-only facts remain passing.
- Existing named/wildcard ImportFact compatibility remains available to current consumers.
- `extract.go::result` continues to flush the same `Exports`, `ExportDiagnostics`, and `Imports` collections.
- `scopeir` normalization/serialization, provider parity, resolver, and analyze packages pass unchanged.

Legacy fallback status:

- No silent drop remains for malformed source-bearing export sites exercised by this slice.
- No `Visibility="exported"` writer or independent compatibility export truth was added.
- No terminal-resolution fallback or physical-definition lookup was added.

Stale tests/helpers/plans updated:

- The P4-B later-slice guard in `TestExtractExportDiagnosticsAndLaterSliceBoundaries` now asserts P4-B1 facts/diagnostics while continuing to reject later terminal state.
- No plan/ledger was edited by Coder; Main owns ledger refresh after independent Supervisor review.

## Fresh Anvien Evidence — `E4-P4B1-IMPACT1` Candidate

Pre-edit excluded graph:

```text
anvien analyze . --force --exclude 'internal/aicontext/skills/**' --exclude '.claude/skills/**' --json
exit 0; scanned/parsed/failed = 1,136/626/0; graph = 81,772 nodes / 121,285 relationships
```

Pre-edit file-detail and upstream file impact:

| File | Symbols | Related | In / Out / Local | Linked flows/tests | File-detail risk | Upstream impact |
|---|---:|---:|---:|---:|---|---|
| `internal/providers/tsjs/imports.go` | 129 | 17 | 7 / 82 / 27 | 3 / 4 | HIGH | MEDIUM; `10` impacted, `10` direct, `1` file, `1` flow |
| `internal/providers/tsjs/extract.go` | 39 | 25 | 24 / 35 / 32 | 4 / 6 | HIGH | CRITICAL; `24` impacted, `11` direct, `11` files, `1` flow |

Exact pre-edit symbol evidence constrained the change: existing re-export/direct-export methods were LOW/0; `Extract` was CRITICAL (`11` impacted, `3` direct, `7` modules, `35` processes), so no `Extract` edit was made and existing `result` wiring was reused.

Final excluded graph after canonical build:

```text
anvien analyze . --force --exclude 'internal/aicontext/skills/**' --exclude '.claude/skills/**' --json
exit 0; scanned/parsed/failed = 1,137/626/0; graph = 81,966 nodes / 121,621 relationships
```

Final-byte file evidence:

| File | Symbols | Related | In / Out / Local | Linked flows/tests | Upstream impact |
|---|---:|---:|---:|---:|---|
| `internal/providers/tsjs/imports.go` | 173 | 17 | 7 / 106 / 38 | 2 / 4 | HIGH; `17` impacted, `17` direct, `1` file, `1` flow |
| `internal/providers/tsjs/extract.go` | 39 | 25 | 24 / 35 / 32 | 3 / 6 | CRITICAL; `23` impacted, `11` direct, `11` files, `1` flow |

All nine exact changed/new syntax methods and helpers checked on final bytes (`emitExportStatement`, `emitSourceExportFacts`, `emitReexportClauseFacts`, `emitStarExportFact`, `emitNamespaceExportFact`, `addSourceExportFact`, `emitRecoveredExportDiagnostic`, `sourceExportStatementTypeOnly`, `sourceExportHasUnexpectedMalformedSyntax`) report LOW risk with zero upstream impacted symbols/files/processes. HIGH/CRITICAL file results are blast-radius warnings, not edit prohibitions; they justify the full build and named consumer regressions while the edit remains in the exact syntax owner.

## Production Result — `E4-P4B1-SRC1` Candidate

- `emitExportStatement` now routes a source-bearing statement into first-class syntax extraction while source-less P4-B behavior remains unchanged.
- Named/default re-export specifiers emit one `ExportReexport` with:
  - consumer-facing `ExportedName`;
  - source-side `TargetExportedName`;
  - empty `LocalName`/`LocalDefID`;
  - stripped, non-resolved `TargetRaw`;
  - exact specifier `Range`, exported-token `SelectionRange`, and full statement provenance;
  - value or explicit type-only meaning.
- Plain star emits one `ExportStar`; namespace star emits one `ExportNamespace`. Namespace aliases include identifier, `default`, and string-literal export-name forms.
- The installed TypeScript grammar does not natively accept `export type *`; production recognizes only the exact top-level `type` marker recovery node and rejects any other grammar error. Both type-star and type-namespace forms are covered directly.
- Valid siblings survive a malformed source-bearing clause; the malformed site produces a structured diagnostic and does not fabricate a compatibility import.
- Existing `ImportReexport`/`ImportWildcard` facts are created only from an emitted first-class source export fact, preserving compatibility without a separately parsed second truth.
- No graph, persistence, terminal target, barrel traversal, ambiguity, cycle, or public-API state exists in this change.

## Tests — `E4-P4B1-TEST1` Candidate

Direct tests in `internal/providers/tsjs/extract_test.go`:

- line 488: `TestExtractTypeScriptReexportSyntaxFacts` — `16/16` TS facts across named, aliases, source/default aliases, statement/inline type-only, `typeof`, star, namespace, type-star recovery, type-namespace recovery, and string export names; exact names/ranges/provenance/meaning/`TargetRaw`; compatibility and zero-terminal assertions.
- line 565: `TestExtractJavaScriptReexportSyntaxFacts` — `8/8` JS facts across named/default/star/namespace/default-namespace forms; compatibility and zero-terminal assertions.
- line 605: `TestExtractExportDiagnosticsAndLaterSliceBoundaries` — three valid source-bearing facts, three missing-source diagnostics, one eligible sibling beside one malformed alias diagnostic, preserved compatibility, and no terminal state.

P4-B regression tests retained and rerun:

- line 285: `TestExtractTypeScriptDirectAndDefaultExportFacts` — PASS.
- line 369: `TestExtractTypeScriptLocalAliasAndTypeOnlyExportFacts` — PASS.
- line 433: `TestExtractJavaScriptDirectDefaultAndLocalExportFacts` — PASS.

Focused command after full build:

```text
go test -v ./internal/providers/tsjs -run '^(TestExtractTypeScriptDirectAndDefaultExportFacts|TestExtractTypeScriptLocalAliasAndTypeOnlyExportFacts|TestExtractJavaScriptDirectDefaultAndLocalExportFacts|TestExtractTypeScriptReexportSyntaxFacts|TestExtractJavaScriptReexportSyntaxFacts|TestExtractExportDiagnosticsAndLaterSliceBoundaries)$' -count=1
exit 0; 6/6 top-level named tests PASS; all diagnostic subtests PASS
```

Full package:

```text
go test ./internal/providers/tsjs -count=1
exit 0; PASS
```

## Full Build and Real Boundary — `E4-P4B1-BUILD1` / `E4-P4B1-BOUNDARY1` Candidates

Pre-build holder gate:

- `anvien doctor locks --repo E:\Anvien --json`: free; lock absent/not alive.
- Five editor-owned global `anvien.exe mcp` processes and their remaining command parents, plus one repo Playwright test-server, were stopped before build.
- Recount before build: build-related holder count `0`.

Canonical full build:

```text
npm run full-build
exit 0
```

- Packaged/global CLI version `1.2.8`.
- Launcher/native runtime build PASS.
- Web production build PASS; `2,943` modules transformed; Vite completed in `23.43s`.
- Existing mixed static/dynamic import and chunk-size messages were non-failing warnings.
- The build's final unexcluded analyze is not semantic evidence; the final excluded analyze recorded above supersedes it.

Nearest real provider/ScopeIR boundary:

```text
go test ./internal/providers/tsjs ./internal/scopeir ./internal/providers -count=1
exit 0; 3/3 packages PASS
```

Compatibility consumer regression (no Child 05 terminal claim):

```text
go test ./internal/resolution ./internal/analyze -count=1
exit 0; 2/2 packages PASS
```

Formatting/source gates:

- `gofmt -d internal/providers/tsjs/imports.go internal/providers/tsjs/extract_test.go`: empty.
- `git diff --check`: exit `0`.
- Production diff contains none of `TargetFile`, `TargetDefID`, `TargetModuleScope`, `LinkStatus`, `TransitiveVia`, barrel/public-API fields, or `Visibility` writes.

E2E Verification:

```text
[PASS] Compiled: npm run full-build -> exit 0
[PASS] Runtime boundary: TypeScript/JavaScript source -> parser -> tsjs.Extract -> ScopeIR.Exports/ExportDiagnostics/Imports -> NormalizeOwned
[PASS] Happy path: 24 eligible TS/JS source-bearing sites -> 24 exact immutable facts + 24 derived compatibility facts
[PASS] Edge case: 4 malformed source-bearing sites -> 4 structured diagnostics; valid mixed-clause sibling retained
```

## Benchmark / Inventory Results

| Metric | Baseline | Candidate | Target | Result |
|---|---:|---:|---:|---|
| Requested P4-B1 forms represented (named, default, star, namespace, type-only) | `0/5` first-class paths | `5/5` | `5/5` | PASS |
| Facts per eligible focused source site | `0/24` | `24/24 = 1.0` | `1.0` | PASS |
| Compatibility facts derived from focused source facts | two legacy syntax classes only | `24/24` exact matches | `0` drift | PASS |
| Explicit type-only sites with value meaning | N/A | `0/8` | `0` | PASS |
| Malformed source-bearing sites with structured diagnostics | `0/4` in the P4-B later-slice guard | `4/4` | `100%` | PASS |
| Child 05 terminal/resolution JSON fields present | `0/7` | `0/7` | `0/7` | PASS |
| P4-B1 debug/probe artifacts remaining | N/A | `0` | `0` | PASS |

This slice does not intentionally change measured product/runtime performance, graph/DB throughput, capacity, startup/package size, or target graph inventory. Build/test timings are validation evidence rather than a product benchmark.

## Exact Candidate Identity

| Path | Lines | Bytes | SHA-256 |
|---|---:|---:|---|
| `internal/providers/tsjs/imports.go` | 1,247 | 34,219 | `4D6A796F305D4CB9812B6600385E0215DEFF8A27F557B41559A7BF634A95C850` |
| `internal/providers/tsjs/extract_test.go` | 3,057 | 133,138 | `7BC5C215414DFECC23F5E6B26EDC9F90BAE2826F8F8AC03A9E1BAD34A5CC9AEA` |

Git reference for the uncommitted candidate: baseline HEAD `11a37aa8ec0320dd93258c058b088d1070aa778d` plus the exact unstaged two-file diff above. The staged index is empty.

## Cleanup / Residuals / Handoff

- Repo-local debug probes `.tmp/p4b1_ast_probe.go` and `.tmp/p4b1_extract_probe.go` were deleted.
- Fresh recursive `.tmp` census for `p4b1`, re-export probes, and AST probes: `0` matching artifacts.
- No temporary directory was created outside `E:\Anvien`.
- No `E:\cheapapp.org` path was accessed.
- Residual unverified surfaces inside the assigned P4-B1 invariant family: `none`.
- Intentionally locked later surfaces: P4-C graph/persistence projection, P4-C2 target validation, Child 05 export-table/terminal resolution, barrels, ambiguity, cycles, and package public API.
- Risks/open points inside P4-B1: none known. The file-layer blast radius remains HIGH/CRITICAL as recorded, while exact changed syntax symbols are LOW/0 and all named boundaries pass.
- Required next action: Main should open one independent Supervisor review against this exact report/worktree. Only after Supervisor PASS should Main refresh living ledgers, run `anvien detect-changes`, stage the exact accepted boundary, and create the isolated no-push commit.
