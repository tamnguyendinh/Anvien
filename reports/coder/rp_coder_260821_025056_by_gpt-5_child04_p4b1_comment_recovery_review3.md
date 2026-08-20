# Child 04 / P4-B1 Coder Resubmission REVIEW3 — Comment-bearing Re-export Recovery

## Handoff State

- State: `READY_FOR_SUPERVISOR`.
- Scope: Child 04 P4-B1 only; this resubmission repairs the sole blocker from Supervisor REVIEW2.
- Supervisor authority: `reports/Supervisor/rp_supervisor_260821_022347_by_gpt-5_child04_p4b1_reexport_resubmission_review2.md`.
- REVIEW2 identity: `14,671` bytes / `157` LF / SHA-256 `B06061C6A765AEC40CDFD43B29C7AC91AB2EB2B6197A8C116CFCF3B1A82084AF`.
- REVIEW2 verdict: `REJECT`.
- Authorized implementation baseline: `11a37aa8ec0320dd93258c058b088d1070aa778d`.
- Current HEAD: `ce0e200c55bd96c4374cc6e84bd99a3c82bef641`.
- No detect-changes, stage, commit, push, target access, plan/ledger/roadmap edit, docs-tree edit, or new reviewer was performed.

## Git / External Drift Boundary

Current HEAD contains exactly two external docs-only commits after the authorized baseline:

- `84a354940aea8240c99bf4868e721209e7248830` — `docs(orchestration): mark session rotation mandatory`.
- `ce0e200c55bd96c4374cc6e84bd99a3c82bef641` — `docs(orchestration): enforce session rotation steps`.

Both commits change only `internal/aicontext/skills/orchestration/SKILL.md`. The command
`git diff --quiet 11a37aa8... HEAD -- internal/providers/tsjs/imports.go internal/providers/tsjs/extract_test.go`
exited `0`; candidate ownership did not drift. These commits were preserved and were not reset, checked out, or edited.

The staged index is empty. The tracked worktree diff contains exactly:

- `internal/providers/tsjs/imports.go`
- `internal/providers/tsjs/extract_test.go`

Main-owned Investigation reports, both Supervisor reports, and both historical Coder reports remain preserved and unmodified.

## Rejected Invariant

REVIEW2 proved that the prior recovery helper accepted only a four-child recovered `export_specifier` with whitespace-only gaps. Valid comments are represented as non-error `comment` children inside the same recovered specifier, so these sources lost `AlsoGood`:

    export { Good, Broken as, /*keep*/ AlsoGood } from "./mixed";
    export { Good, Broken as /*bad*/, AlsoGood } from "./mixed";

Required result for each source:

    exports=2 [Good, AlsoGood]
    diagnostics=1 [dangling as]
    imports=2
    Broken facts/imports=0

## Invariant Family Map

| Surface | Authority / owner | REVIEW3 result |
|---|---|---|
| Comment-bearing malformed alias recovery | `imports.go::recoveredReexportSiblingAfterMalformedAlias` | Comments/extra trivia may surround the recovered comma; valid `AlsoGood` survives |
| Diagnostic | P4-A `ExportDiagnosticFact` contract | Exactly one malformed diagnostic remains on the dangling `as` site |
| Valid sibling facts | P4-A `ExportFact`; P4-B1 syntax owner | `Good` and `AlsoGood` each emit one exact source-bearing re-export fact |
| Compatibility | Existing `ImportReexport`, derived from source facts | Exactly two imports; no independently parsed second truth |
| Fail-closed behavior | Existing malformed-specifier path | Missing/malformed name or alias, duplicate required tokens, non-comma errors, and non-trivia children are rejected |
| Later-slice boundary | Child 05 owns terminal resolution | No graph/persistence/terminal/barrel/ambiguity/cycle/public-API state |

Sibling surfaces checked: no-comment recovery, newline recovery, direct/default/local/type-only P4-B facts, TS/JS named/default/star/namespace/type-only P4-B1 facts, missing-source diagnostics, compatibility imports, ScopeIR normalization, provider boundary, resolution/analyze consumers, and access visibility.

Legacy fallback status: the valid comment-bearing sibling is no longer silently dropped; malformed `Broken` is not fabricated; no terminal-resolution fallback was added.

Stale tests/helpers/plans updated: only the focused extraction test was extended. Plans and ledgers were explicitly preserve-only.

## Production Change

Production changed only `internal/providers/tsjs/imports.go`.

`recoveredReexportSiblingAfterMalformedAlias` now scans the recovered specifier children rather than requiring `ChildCount()==4`. It accepts exactly one of each required semantic token in source order:

1. supported, non-malformed `name`;
2. anonymous `as`;
3. one `ERROR` whose exact trimmed source text is `,`;
4. supported, non-malformed `alias`.

Only `comment` or parser-extra nodes may occur between those tokens. Missing nodes, duplicate semantic tokens, a malformed name/alias, any other error text, or any non-trivia child returns `false` and preserves the prior fail-closed diagnostic branch.

The recovered alias node remains the fact `Range`, `SelectionRange`, source-side name, and exported name. `addSourceExportFact` remains the sole compatibility derivation path.

No ScopeIR contract, extract wiring, definitions, visibility, graph, persistence, resolver, or later-slice source changed.

## Regression Test

Test owner: `internal/providers/tsjs/extract_test.go`.

Test: `TestExtractExportDiagnosticsAndLaterSliceBoundaries`, current lines `605-756`.

Recovery subcases:

- `no-comment`
- `comment-after-comma`
- `comment-before-comma`
- `newline-after-comma`

Every subcase asserts:

- `Exports/ExportDiagnostics/Imports = 2/1/2`;
- exact `Good` and `AlsoGood` `Range` and `SelectionRange`;
- exact full `StatementRange` and `SiteKind=export_specifier`;
- `TargetRaw="./mixed"`;
- value meaning, `TypeOnly=false`;
- empty `LocalName` and `LocalDefID`;
- diagnostic code, `NodeKind=as`, exact `as` range, and `export_statement` provenance;
- no `Broken` fact/import;
- one derived compatibility import per valid fact;
- zero terminal/resolution JSON fields.

## Canonical Build and Validation

Pre-build holder gate:

- analyze lock: `free`;
- eight editor-owned global `anvien.exe mcp` children and their eight verified `cmd.exe` parents were stopped because `npm install -g .` replaces the held package;
- build-holder recount: `0`;
- Codex parent/runtime processes were preserved.

Canonical build, after production and test changes:

    npm run full-build
    exit 0
    runtime 1.2.8
    Web modules transformed: 2,943
    Vite build: 22.35s

Existing mixed static/dynamic import and chunk-size messages were non-failing warnings. The build-internal unexcluded analyze is operational output and is not used as semantic graph evidence.

Validation after the full build, repeated successfully after the Codex app update/restart:

    go test -v ./internal/providers/tsjs -run '^(TestExtractTypeScriptDirectAndDefaultExportFacts|TestExtractTypeScriptLocalAliasAndTypeOnlyExportFacts|TestExtractJavaScriptDirectDefaultAndLocalExportFacts|TestExtractTypeScriptReexportSyntaxFacts|TestExtractJavaScriptReexportSyntaxFacts|TestExtractExportDiagnosticsAndLaterSliceBoundaries)$' -count=1
    exit 0; 6/6 top-level tests PASS; all four recovery subcases PASS

    go test ./internal/providers/tsjs -count=1
    exit 0

    go test ./internal/providers/tsjs ./internal/scopeir ./internal/providers -count=1
    exit 0; 3/3 packages PASS

    go test ./internal/resolution ./internal/analyze -count=1
    exit 0; 2/2 packages PASS

Independent parser → `tsjs.Extract` → ScopeIR exact-field boundary output:

    PASS source="export { Good, Broken as, /*keep*/ AlsoGood } from \"./mixed\";" exports=2 names=[Good AlsoGood] diagnostics=1 nodeKind=as range=as imports=2 targetRaw=./mixed meanings=[value] typeOnly=false local=empty terminalFields=0
    PASS source="export { Good, Broken as /*bad*/, AlsoGood } from \"./mixed\";" exports=2 names=[Good AlsoGood] diagnostics=1 nodeKind=as range=as imports=2 targetRaw=./mixed meanings=[value] typeOnly=false local=empty terminalFields=0

The boundary probe asserted exact fact ranges, selection ranges, statement provenance, site kinds, target/name fields, meanings/type-only state, compatibility contents, absence of `Broken`, and seven forbidden terminal fields before printing PASS. It was deleted after execution.

E2E Verification:

    [PASS] Compiled: npm run full-build -> exit 0
    [PASS] Runtime: parser -> tsjs.Extract -> ScopeIR Exports/Diagnostics/Imports
    [PASS] Happy path: Good + AlsoGood -> two exact facts and two derived imports
    [PASS] Edge case: comments on either side of recovered comma -> one dangling-as diagnostic, no Broken fact/import

## Fresh Anvien Evidence

Fresh final excluded graph after app restart and after this resubmission report was present:

    anvien analyze . --force --exclude 'internal/aicontext/skills/**' --exclude '.claude/skills/**' --json
    exit 0
    scanned/parsed/failed = 1,144/626/0
    graph = 82,059 nodes / 121,760 relationships

The inventory includes the preserved Main-owned `0228` rotation handoff report and this Coder resubmission report; both are evidence artifacts, not production scope.

Final file-detail:

- `imports.go`: 185 symbols; inbound/outbound/local `7/107/40`; unresolved `497`; linked flows/tests `2/4`; risk HIGH; stale=false; changedSinceAnalyze=false.
- `extract_test.go`: 713 symbols; inbound/outbound/local `13/366/245`; unresolved `0`; linked flows/tests `0/3`; risk LOW; stale=false; changedSinceAnalyze=false.

Final upstream impact:

- `imports.go`: HIGH, `18` impacted / `18` direct / `1` affected file / `1` affected flow.
- `extract_test.go`: CRITICAL, `63` impacted / `60` direct / `4` affected files / `0` affected flows.
- `emitReexportClauseFacts`: LOW, `0` upstream impacted.
- `recoveredReexportSiblingAfterMalformedAlias`: LOW, `0` upstream impacted.
- `TestExtractExportDiagnosticsAndLaterSliceBoundaries`: LOW, `0` upstream impacted.

HIGH/CRITICAL are recorded blast-radius warnings, not edit prohibitions. The change remains surgical in the exact LOW/0 syntax owner and focused test.

## Benchmark / Cardinality

| Metric | REVIEW2 candidate | REVIEW3 candidate | Target |
|---|---:|---:|---:|
| Comment-after-comma valid facts | 1/2 | 2/2 | 2/2 |
| Comment-before-comma valid facts | 1/2 | 2/2 | 2/2 |
| Diagnostics per comment case | 1 | 1 | 1 |
| Compatibility imports per comment case | 1/2 | 2/2 | 2/2 |
| Broken facts/imports | 0 | 0 | 0 |
| No-comment/newline regression | 2/1/2 | 2/1/2 | 2/1/2 |
| Terminal fields | 0/7 | 0/7 | 0/7 |

This slice does not intentionally change product/runtime performance, capacity, package size, startup time, or graph throughput. Build/test timing is validation evidence only.

## Exact Candidate Identity

| Path | Lines | Bytes | LF | SHA-256 |
|---|---:|---:|---:|---|
| `internal/providers/tsjs/imports.go` | 1,318 | 36,488 | 1,318 | `52AFBA2FC9A7ACEA314D5B39043A054F2671B2EF987161D8F1C3D641B2382749` |
| `internal/providers/tsjs/extract_test.go` | 3,105 | 135,246 | 3,105 | `07CF7D49715CA0398DA9485086A37E395FCB8E2E695AF2B649B6BC05074C604D` |

Diff relative to the authorized implementation baseline/current HEAD working tree:

- `imports.go`: `440` insertions / `22` deletions.
- `extract_test.go`: `296` insertions / `2` deletions.
- Total: `736` insertions / `24` deletions.

## Final Gates / Cleanup

- `gofmt -d internal/providers/tsjs/imports.go internal/providers/tsjs/extract_test.go`: empty, exit `0`.
- `git diff --check`: empty, exit `0`.
- Production forbidden terminal/barrel/public-API token matches: `0`.
- Production visibility writes: `0`.
- Exact P4-B1/re-export/AST probe census under repo-local `.tmp`: `0`.
- Analyze lock: `free`.
- HEAD remains `ce0e200c55bd96c4374cc6e84bd99a3c82bef641`.
- Staged index count: `0`.
- Tracked candidate path count: `2`.
- No temporary directory was created outside `E:\Anvien`.
- `E:\cheapapp.org` was not accessed.

Residual unverified surfaces inside the rejected invariant family: none.

Intentionally locked: P4-C graph/persistence projection, P4-C2 target validation, Child 05 terminal/barrel/ambiguity/cycle/public-API behavior.

Next action: Main task `01a02074-9de3-7072-a2fa-c2ef10db6358` should send this exact candidate/report to existing Supervisor lane `01a02059-d255-78c0-b608-b199440bcf18` for REVIEW3. Coder stops after handoff.
