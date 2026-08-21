# Child 05 / P5-C Re-export Traversal — Pre-Implementation Inventory

## Status and authority

- Status: `PRE_IMPLEMENTATION_READY / READY_FOR_MAIN_AUTHORIZATION`.
- This is an inventory and impact handoff only. No production source, test, plan, evidence ledger, benchmark ledger, actual-status ledger, or existing report was edited.
- Authoritative checkout: `E:\Anvien`; current HEAD `cd35b48f5466117fa1348fdc71c52e1408685a1b` (`docs(plan): open P5-C after P5-B`).
- Predecessor: accepted P5-B implementation commit `c1559df953a277b099009f8489576d00ed25aa58`.
- P5-D, target, and `E:\cheapapp.org` remain locked and were not accessed. No C-worktree, push, reset, checkout, build, detect-changes, commit, Supervisor, or subagent action was run.

## Fresh graph gate

From cwd `E:\Anvien`, exactly one fresh graph refresh was run before graph queries:

```text
anvien analyze --force
exit: 0
scanned: 1959
parsed_code: 738
failed: 0
graph: E:\Anvien\.anvien\graph.json
nodes / relationships: 115134 / 158153
analyzedAt: 2026-08-21T13:18:34Z
graph SHA-256: 2D071FBC162E2CEE396C9A9E4DEC2376DD1313F4357ECDB0D43F3A8D3612B957
```

Every subsequent `file-detail` and `impact` result below reported `stale=false`, `changedSinceAnalyze=false`, and complete semantic schema (`115134/115134` nodes with `appLayer` and `functionalArea`; missing fields `0`; resolution-health gaps `0`).

## Invariant Family Map

- Invariant family: repository-backed TypeScript module export lookup and terminal proof preparation (P5-C only).
- Authority / SSOT: accepted Child 04 `ScopeIR.ExportFact` (source kind, exported/target names, meanings, type-only state, ranges, provenance) consumed by the accepted P5-B `exportTables` boundary. `ImportFact.RequestedMeanings`/`TypeOnly` is the accepted P5-A request contract.
- Sibling surfaces checked: P5-B table construction, workspace import orchestration, physical imported-definition lookup, namespace/member imported lookup, call resolution, generic global lookup helpers, import-path strategies, import-edge emission, existing resolver tests/benchmarks, and the four current P5 ledgers.
- Forbidden legacy fallback: no physical-definition export inference, no first-candidate selection, no implicit star default, no meaning repair by name-only lookup, no terminal result fabricated from a compatibility `ImportFact`, and no repository-global same-name rescue after an explicit import failure.
- Stale tests/helpers/plans: no files updated in this inventory. Existing tests listed below prove adjacent legacy behavior only; they do not yet prove the P5-C topology matrix.
- Verify matrix for the later authorized slice: alias/re-export chains; explicit-over-star precedence; star default exclusion; namespace/member traversal; same-terminal path dedupe; distinct-candidate ambiguity; terminal and pure cycles; value/type/namespace meaning mismatch; explicit-import miss versus global same-name definition; direct/path/`IMPORTS` preservation; legacy non-TS/JS regression.
- Residual unverified surfaces: implementation, focused P5-C tests, full build, runtime boundary, target/P5-D proof, Supervisor, detect-changes, and commit are intentionally pending Main authorization.

## Accepted P5-B boundary inspected

`E:\Anvien\internal\resolution\export_tables.go:12-248` is the accepted semantic table owner (SHA-256 `BA4C7F9F2D05F6715CA22195C1CFF64C1AACADC3D81E7CA3ED711C8EA8B63E19`). It stores copied accepted facts in `Explicit` and unexpanded `StarAdjacency`; it deliberately does not inspect physical definitions or resolve terminal Symbols. The accepted test owner `internal/resolution/export_tables_test.go` is byte-stable (SHA-256 `A0395EF9030CB054B09C52AFBE048D2F8244BAD8A924EE492BB7AB312CD08AD8`).

Current orchestration in `E:\Anvien\internal\resolution\indexes.go` (`26DE75A911B4B1E1471C61548C4153749E067671DB34852E51448D52D4C0C486`) is:

```text
buildWorkspace: w.resolveImports() -> w.buildExportTables() -> w.resolveHeritage()
resolveImports: resolveImportFiles -> resolveImportedDef -> scope binding
```

Thus P5-B tables are populated after the current physical import lookup. A P5-C implementation must make table availability and traversal a bounded two-phase/orchestration decision; it must not duplicate path resolution or alter the accepted table/Child 04 owner.

## Current source findings and exact owner candidates

### Re-export traversal

- `indexes.go:462-486`, `workspace.resolveImportedDef(targetFile, item)`: current first divergence. It scans `w.defsByFile[targetFile]` by imported/local name and has a default fallback that selects the first class/function/interface. It never reads `w.exportTables`, follows no alias/star hop, carries no cycle/ambiguity/meaning proof, and can treat a physical declaration as an export.
- `indexes.go:282-317`, `workspace.resolveImports`: calls `resolveImportedDef` while building `resolvedImport`, then creates `scopeBindings` only when `TargetDef` exists. This is the minimum import-binding orchestration seam.
- `indexes.go:607-638`, `workspace.resolveImportedMember`: namespace/member lookup independently scans physical definitions in each target file. It is a sibling P5-C surface for namespace/star traversal and must either consume the same proof-bearing export lookup or be explicitly proven unaffected.
- `indexes.go:73-180`, `buildWorkspace`: the call order makes table availability a real orchestration dependency. The fresh impact marks this entrypoint CRITICAL; Main must authorize either a minimal call-order adapter here or an equivalent two-phase implementation contained by `resolveImports`.

### Explicit-import global-name-rescue boundary

- `resolve.go:360-489`, `resolveCall`: after scoped/same-file/imported-member attempts fail, free/constructor/member calls call `w.resolveGlobalCallName`; a low-confidence result is converted to a diagnostic. When `resolveImportedDef` failed, no scoped import binding exists, so an explicit import can reach this fallback.
- `indexes.go:528-541`, `workspace.resolveGlobalCallName`: generic fallback helper used by multiple call forms. It must remain generic/preserve-only unless later source evidence proves a narrowly scoped change; the explicit-import guard belongs at the `resolveCall` boundary or an equivalent import-state adapter.
- `indexes.go:488-526`, `resolveName`/`resolveGlobalName`: broad scope/type/heritage/member lookup helpers. They are preserve-only for this slice; changing them would widen P5-C into unrelated resolution behavior.

## Fresh file-detail evidence

| File | Symbols | Inbound / outbound / local | Unresolved | Flows / tests | File risk |
|---|---:|---:|---:|---:|---|
| `internal/resolution/indexes.go` | 320 | 215 / 96 / 372 | 440 | 13 / 29 | HIGH |
| `internal/resolution/resolve.go` | 113 | 90 / 168 / 57 | 201 | 22 / 32 | HIGH |
| `internal/resolution/export_tables.go` | 62 | 23 / 39 / 41 | 70 | 3 / 18 | HIGH (P5-B preserve-only) |
| `internal/resolution/import_resolution.go` | 116 | 22 / 59 / 7 | 237 | 0 / 18 | HIGH (path preserve-only) |

Repo-local debug captures were retained unchanged for traceability:

- `.tmp/p5c-file-detail-indexes-20260821.json` — 306,523 bytes, SHA-256 `0D17D70EC94880D30CAD40C55B7AB6674712E149FAE5554385A9F433FBB1FC63`.
- `.tmp/p5c-file-detail-resolve-20260821.json` — 180,676 bytes, SHA-256 `3AE1C29385DDFCEC8A75C0C31BE32FC18E85BDC43BF1DA1C4751D9A4447C13C5`.
- `.tmp/p5c-file-detail-export_tables-20260821.json` — 61,714 bytes, SHA-256 `6E76B5CA8C0CFE33DD3D914AB63293E997B789A2EE9038C95D032A8823A71927`.
- `.tmp/p5c-file-detail-import_resolution-20260821.json` — 100,296 bytes, SHA-256 `C3015F6D1975F37CD43DD369C637B040B062ED4C8BB57853E35C723DF7F73F39`.

## Complete fresh upstream impact (include-tests)

All rows below are upstream symbol impact from the same fresh graph. HIGH/CRITICAL is a scope warning, not an edit prohibition.

| Exact symbol | Risk | Impacted symbols | Affected files | Modules | Processes | App layers | Functional areas |
|---|---|---:|---:|---:|---:|---|---|
| `workspace.resolveImportedDef` (`indexes.go:462`) | HIGH | 19 | 12 | 1 | 3 | backend 4, backend_test 15 | resolution 19 |
| `workspace.resolveImports` (`indexes.go:282`) | CRITICAL | 28 | 16 | 3 | 17 | backend 6, backend_test 22 | analyzer 1, graph_health 1, resolution 26 |
| `workspace.resolveImportedMember` (`indexes.go:607`) | CRITICAL | 18 | 8 | 3 | 34 | backend 11, backend_test 7 | analyzer 1, graph_health 1, resolution 16 |
| `buildWorkspace` (`indexes.go:73`) | CRITICAL | 49 | 22 | 8 | 23 | backend 7, backend_test 41, cli_launcher 1 | analyzer 16, cli 2, graph_health 1, resolution 30 |
| `resolveCall` (`resolve.go:360`) | CRITICAL | 27 | 11 | 7 | 32 | backend 6, backend_test 21 | analyzer 16, cli 1, graph_health 1, resolution 9 |
| `workspace.resolveGlobalCallName` (`indexes.go:528`) | CRITICAL | 6 | 4 | 2 | 23 | backend 4, backend_test 2 | analyzer 1, resolution 5 |

Complete affected-file sets for the six HIGH/CRITICAL rows:

```text
workspace.resolveImportedDef:
  internal/resolution/access_audit.go
  internal/resolution/export_tables_test.go
  internal/resolution/identity_occurrence_test.go
  internal/resolution/indexes.go
  internal/resolution/legacy_heritage_map_conversion_test.go
  internal/resolution/legacy_import_language_conversion_test.go
  internal/resolution/legacy_scope_indexes_test.go
  internal/resolution/legacy_scope_symbol_semantics_conversion_test.go
  internal/resolution/legacy_suffix_index_ambiguity_conversion_test.go
  internal/resolution/p2b_persistence_test.go
  internal/resolution/resolution_test.go
  internal/resolution/resolve.go

workspace.resolveImports:
  internal/analyze/analyze.go
  internal/graphaccuracy/access_candidate.go
  internal/resolution/access_audit.go
  internal/resolution/access_audit_test.go
  internal/resolution/export_tables_test.go
  internal/resolution/identity_occurrence_test.go
  internal/resolution/indexes.go
  internal/resolution/legacy_heritage_map_conversion_test.go
  internal/resolution/legacy_import_language_conversion_test.go
  internal/resolution/legacy_p7_conversion_test.go
  internal/resolution/legacy_scope_indexes_test.go
  internal/resolution/legacy_scope_symbol_semantics_conversion_test.go
  internal/resolution/legacy_suffix_index_ambiguity_conversion_test.go
  internal/resolution/p2b_persistence_test.go
  internal/resolution/resolution_test.go
  internal/resolution/resolve.go

workspace.resolveImportedMember:
  internal/analyze/analyze.go
  internal/graphaccuracy/access_candidate.go
  internal/resolution/access_audit.go
  internal/resolution/access_audit_test.go
  internal/resolution/indexes.go
  internal/resolution/legacy_p7_conversion_test.go
  internal/resolution/resolution_test.go
  internal/resolution/resolve.go

buildWorkspace:
  cmd/access-candidate-audit/main.go
  internal/analyze/analyze.go
  internal/analyze/analyze_test.go
  internal/analyze/legacy_resolver_conversion_test.go
  internal/analyze/pipeline_parity_test.go
  internal/cli/command.go
  internal/graphaccuracy/access_candidate.go
  internal/resolution/access_audit.go
  internal/resolution/access_audit_test.go
  internal/resolution/definition_collision_test.go
  internal/resolution/export_tables_test.go
  internal/resolution/identity_occurrence_test.go
  internal/resolution/legacy_heritage_map_conversion_test.go
  internal/resolution/legacy_import_language_conversion_test.go
  internal/resolution/legacy_p7_conversion_test.go
  internal/resolution/legacy_scope_indexes_test.go
  internal/resolution/legacy_scope_symbol_semantics_conversion_test.go
  internal/resolution/legacy_suffix_index_ambiguity_conversion_test.go
  internal/resolution/p2b_persistence_test.go
  internal/resolution/p3c_binding_occurrence_test.go
  internal/resolution/resolution_test.go
  internal/resolution/resolve.go

resolveCall:
  internal/analyze/analyze.go
  internal/analyze/analyze_test.go
  internal/analyze/legacy_resolver_conversion_test.go
  internal/analyze/pipeline_parity_test.go
  internal/cli/command.go
  internal/graphaccuracy/access_candidate.go
  internal/resolution/definition_collision_test.go
  internal/resolution/legacy_p7_conversion_test.go
  internal/resolution/p3c_binding_occurrence_test.go
  internal/resolution/resolution_test.go
  internal/resolution/resolve.go

workspace.resolveGlobalCallName:
  internal/analyze/analyze.go
  internal/resolution/legacy_p7_conversion_test.go
  internal/resolution/resolution_test.go
  internal/resolution/resolve.go
```

Additional preserve-only impact checks:

- `workspace.resolveName`: CRITICAL, 34 symbols / 14 files / 2 modules / 38 processes; broad type/heritage/member/call surface, so not a P5-C edit owner.
- `workspace.resolveGlobalName`: CRITICAL, 11 symbols / 3 files / 1 module / 20 processes; generic scope/type/heritage fallback, preserve-only.
- `workspace.resolveImportFiles`: HIGH, 19 symbols / 12 files / 1 module / 3 processes; path lookup remains preserve-only.
- `workspace.buildExportTables`: LOW, 0 impacted symbols/processes; accepted P5-B semantic owner remains preserve-only.

## Minimum editable owner set for Main authorization

Evidence selects two production files and one bounded orchestration decision:

1. `E:\Anvien\internal\resolution\indexes.go`:
   - required traversal owner: `workspace.resolveImportedDef`;
   - required import-binding/table orchestration seam: `workspace.resolveImports`;
   - required namespace/member sibling owner: `workspace.resolveImportedMember`;
   - `buildWorkspace` is the CRITICAL call-order adapter candidate because tables are currently built after imports. Main's planner refresh must choose either one minimal sequencing change at `buildWorkspace` or an equivalent two-phase implementation contained inside `resolveImports`; no broader workspace refactor is justified.
2. `E:\Anvien\internal\resolution\resolve.go`:
   - required explicit-import boundary owner: `resolveCall`; preserve generic `resolveGlobalCallName` behavior for non-import calls and gate only an explicit import failure.

No edit is authorized yet. A new traversal helper file may be proposed only by Main's planner refresh if it preserves this same owner boundary; this report does not preselect speculative infrastructure.

## Preserve-only surfaces and forbidden widening

- Child 04 `ScopeIR.ExportFact` providers/facts/normalization/projection: immutable semantic authority; do not regenerate or reinterpret facts.
- P5-B `export_tables.go`, table data shape, explicit/star provenance, and zero-physical behavior: accepted and preserve-only.
- `resolveImportFiles`, `resolveImportFile`, `import_resolution.go`, all non-TS/JS path strategies, `resolvedImport.TargetFiles`, physical path counts, and syntactic `IMPORTS`: preserve-only.
- `emitImportEdges` (`emit.go:507-552`), graph/persistence/readers, P5-D terminal emission, target, and `E:\cheapapp.org`: locked/preserve-only.
- Generic `resolveName`, `resolveGlobalName`, and `resolveGlobalCallName`: preserve broad fallback behavior; add no global-name rescue special case.

## Read-only invariant verification and current gaps

- Accepted P5-B tests (`export_tables_test.go:11`, `141`, `204`) prove zero-physical table construction, explicit/star storage, deterministic ordering, copied provenance/meaning, and no implicit default. They intentionally do not traverse a terminal Symbol.
- `legacy_scope_symbol_semantics_conversion_test.go:11` proves cyclic module dependency conversion and import metrics, but its fixtures do not exercise accepted `ExportFact` traversal.
- `resolution_test.go:940-1011` proves an unqualified missing call does not emit a resolved global edge; it does not prove an explicit-import miss cannot bind a same-name repository-global Symbol.
- `resolution_test.go:1089-1154` benchmarks the current namespace/member physical-definition scan; it is not a semantic export-table proof.
- Therefore alias/star/cycle/ambiguity/meaning/namespace terminal vectors and the explicit-import no-global-rescue fixture remain genuinely pending P5-C implementation tests, not stale PASS evidence.

## Verification gates for the next authorized slice

```text
E2E Verification:
  [NOT RUN] Compiled: full repository build -> intentionally deferred until Main authorizes production code
  [NOT RUN] Runtime: nearest built resolver/CLI boundary -> intentionally deferred until code exists
  [NOT RUN] Happy path: accepted ExportFact -> terminal Symbol through alias/star/namespace traversal -> pending
  [NOT RUN] Edge case: cycle termination, ambiguity retention, meaning mismatch, and explicit-import no-global rescue -> pending
```

No checklist/ledger row was updated by this lane. Main's next action is to planner-refresh the P5-C owner/scope using this report, explicitly choose the `buildWorkspace` versus `resolveImports` sequencing adapter, then issue production authorization. Only after that authorization may the existing Coder lane edit, add tests, build, and continue the normal review gates.

`PRE_IMPLEMENTATION_READY / READY_FOR_MAIN_AUTHORIZATION`
