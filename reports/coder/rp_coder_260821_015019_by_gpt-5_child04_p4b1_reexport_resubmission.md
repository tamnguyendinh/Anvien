# Child 04 / P4-B1 Coder Resubmission — TypeScript Re-export Recovery

## Handoff State

- State: READY_FOR_SUPERVISOR (fresh resubmission; Supervisor re-review is still required).
- Scope: Child 04 P4-B1 only. This resubmission repairs the sole invariant rejected by REVIEW1.
- Prior Supervisor report: reports/Supervisor/rp_supervisor_260821_012743_by_gpt-5_child04_p4b1_reexport_syntax.md.
- Prior report identity: 13,798 bytes / 158 LF / SHA-256 54D5D6CD7E71DBCF65BD68C0E2BDC0DBF940BE1A48DA51A0B446FAAA0F9652CD.
- Prior verdict: REJECT.
- Baseline HEAD: 11a37aa8ec0320dd93258c058b088d1070aa778d.
- The Main-owned reports and historical Coder report were preserved unchanged.
- No detect-changes, stage, commit, push, plan/ledger/roadmap edit, target access, or reviewer creation was performed.

## Sole Rejected Invariant

The rejected source site was:

    export { Good, Broken as, AlsoGood } from "./mixed";

Tree-sitter recovery produces one export_specifier spanning Broken as, AlsoGood:

    name=Broken, anonymous as, ERROR text=",", alias=AlsoGood

The previous implementation discarded the whole malformed node, resulting in:

    exports=1 [Good], diagnostics=1, imports=1

The required behavior is:

    exports=2 [Good, AlsoGood], diagnostics=1, imports=2

Broken must not produce an ExportFact or ImportReexport. The diagnostic must remain countable and source-backed.

## Invariant Family Map

| Surface | Authority / owner | Resubmission result |
|---|---|---|
| Named/default source-bearing re-export facts | P4-A ScopeIR contract; internal/providers/tsjs/imports.go | One fact per valid specifier, including the recovered AlsoGood sibling |
| Malformed alias recovery | Child 04 P4-B1; emitReexportClauseFacts | Exactly one diagnostic on the dangling as token; malformed Broken is not emitted |
| Compatibility imports | Existing ImportReexport contract, derived by addSourceExportFact | Good and AlsoGood each have one derived import |
| Ranges/provenance | ScopeIR ExportFact / ExportDiagnosticFact | Valid facts use exact node ranges and full statement provenance |
| Later-slice boundary | Child 05 owns terminal resolution | No target file/definition, barrel, ambiguity, cycle, link, or public-API state |

Sibling surfaces checked and preserved: direct/default/local/type-only P4-B facts, star and namespace facts, missing-source diagnostics, ScopeIR normalization, Extract result wiring, resolution/analyze consumers, and access visibility.

Forbidden fallback status: no silent drop for the recovered valid sibling; no physical-definition lookup or terminal-resolution fallback was added.

## Production Change

Only internal/providers/tsjs/imports.go was changed in production.

At emitReexportClauseFacts (current line 166), malformed export_specifier nodes now pass through a source/AST-gated recovery check. The new recovery method recoveredReexportSiblingAfterMalformedAlias accepts only the exact shape with:

- four children in the grammar order name, as, ERROR containing only a comma, alias;
- valid source-side and exported-name nodes;
- no missing or nested malformed name/alias nodes;
- whitespace-only gaps between the AST ranges.

For that shape, the implementation:

- emits one ExportDiagnosticFact at the anonymous as node with statement provenance;
- uses the post-comma alias node as the valid sibling's source/exported name;
- uses that alias node for the exact Range and SelectionRange;
- derives one ImportReexport through the existing addSourceExportFact path;
- leaves all other malformed shapes fail-closed through the prior diagnostic branch.

The normal valid-specifier path is unchanged. No ScopeIR contract, extract wiring, definitions, visibility, graph, persistence, resolver, or Child 05 code changed.

## Regression Test

Only internal/providers/tsjs/extract_test.go was updated after production behavior passed the direct boundary probe.

Test: TestExtractExportDiagnosticsAndLaterSliceBoundaries at current line 698.

The regression asserts for the exact mixed source:

- 2 exports, 1 diagnostic, and 2 imports;
- Good and AlsoGood exact Range and SelectionRange text;
- empty LocalName and LocalDefID;
- TargetRaw ./mixed;
- value meaning and TypeOnly=false;
- full statement provenance and export_specifier site kind;
- diagnostic code, dangling as range, NodeKind, and export_statement provenance;
- no Broken fact or compatibility import;
- no terminal/resolution fields through assertNoTerminalExportState.

## Verification

E2E Verification:

  [PASS] Compiled: npm run full-build -> exit 0; packaged/runtime version 1.2.8; Web build transformed 2,943 modules.
  [PASS] Runtime boundary: parser -> tsjs.Extract -> ScopeIR Exports/ExportDiagnostics/Imports -> NormalizeOwned.
  [PASS] Happy path: Good and AlsoGood -> 2 exact immutable ExportFacts and 2 derived ImportReexport facts.
  [PASS] Edge case: Broken as recovery -> 1 structured malformed diagnostic, no Broken fact/import, valid post-error sibling retained.

Commands after the canonical build:

- Focused matrix: go test -v ./internal/providers/tsjs -run ^(TestExtractTypeScriptDirectAndDefaultExportFacts|TestExtractTypeScriptLocalAliasAndTypeOnlyExportFacts|TestExtractJavaScriptDirectDefaultAndLocalExportFacts|TestExtractTypeScriptReexportSyntaxFacts|TestExtractJavaScriptReexportSyntaxFacts|TestExtractExportDiagnosticsAndLaterSliceBoundaries)$ -count=1 -> 6/6 top-level tests PASS, including all diagnostic subtests.
- Full package: go test ./internal/providers/tsjs -count=1 -> PASS.
- Nearest boundary: go test ./internal/providers/tsjs ./internal/scopeir ./internal/providers -count=1 -> 3/3 packages PASS.
- Compatibility consumers: go test ./internal/resolution ./internal/analyze -count=1 -> 2/2 packages PASS.
- Fresh excluded graph: anvien analyze . --force --exclude internal/aicontext/skills/** --exclude .claude/skills/** --json -> scanned/parsed/failed 1,140/626/0; graph 82,009 nodes / 121,710 relationships.

Final Anvien evidence:

- imports.go file-detail: 184 symbols, 7 inbound, 107 outbound, 40 local, 2 linked flows, 4 linked tests, stale=false, changedSinceAnalyze=false, risk HIGH.
- extract_test.go file-detail: 709 symbols, 13 inbound, 366 outbound, 245 local, stale=false, changedSinceAnalyze=false, risk LOW.
- imports.go file impact: HIGH, 18 impacted/direct symbols, 1 affected file, 1 affected flow.
- emitReexportClauseFacts symbol impact: LOW, 0 upstream impacted symbols/files/processes.
- recoveredReexportSiblingAfterMalformedAlias symbol impact: LOW, 0 upstream impacted symbols/files/processes.
- TestExtractExportDiagnosticsAndLaterSliceBoundaries symbol impact: LOW, 0 upstream impacted symbols/files/processes.
- extract_test.go file impact: CRITICAL, 63 impacted symbols, 60 direct, 4 affected files, 0 affected flows.
- gofmt -d on both candidate files: empty; git diff --check: exit 0.
- Production diff scan: 0 terminal/barrel/public-API field names and 0 visibility writes.

## Benchmark / Cardinality

| Measurement | Rejected candidate | Resubmission | Target |
|---|---:|---:|---:|
| Valid facts for Good + AlsoGood | 1/2 | 2/2 | 2/2 |
| Structured diagnostics for malformed Broken as | 1/1 | 1/1 | 1/1 |
| Derived ImportReexport facts | 1/2 | 2/2 | 2/2 |
| Broken facts/imports | 0 (valid AlsoGood was lost) | 0 | 0 |
| Child 05 terminal fields | 0/7 | 0/7 | 0/7 |

This fix does not intentionally change product/runtime performance, capacity, package size, or graph throughput. Build and test timings are validation evidence only.

## Candidate Identity

| Path | Lines | Bytes | LF | SHA-256 |
|---|---:|---:|---:|---|
| internal/providers/tsjs/imports.go | 1,297 | 36,055 | 1,297 | 0B5E7E8F596CAD6B53CEB17D08FB86046BD071FD0D34DC94FD1FF260EF806A44 |
| internal/providers/tsjs/extract_test.go | 3,076 | 134,250 | 3,076 | A05951319B835D8231E828FC0E172B40E20D594FF0B096D36643BE93C019A757 |

Git state at handoff:

- HEAD remains 11a37aa8ec0320dd93258c058b088d1070aa778d.
- Staged index is empty.
- Tracked candidate diff is exactly the two files above.
- Pre-existing/shared untracked reports were preserved.
- Exact P4-B1 probe census is 0; p4b1/re-export/AST probe files are absent.
- No temporary directory was created outside E:/Anvien.
- E:/cheapapp.org was not accessed.

## Closure / Next Action

Residual unverified surfaces inside the rejected invariant family: none. Later surfaces remain intentionally locked: P4-C graph/persistence projection, P4-C2 target validation, Child 05 terminal/barrel/ambiguity/cycle/public-API behavior.

Main should send this exact worktree and report to the existing Supervisor lane for independent re-review. Coder does not open another reviewer and does not run Main-owned detect-changes, stage, commit, or push.
