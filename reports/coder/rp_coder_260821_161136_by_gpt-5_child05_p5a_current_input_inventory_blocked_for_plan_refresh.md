# Child 05 / P5-A Coder Input Inventory Handoff

## Candidate State

- Trạng thái: `BLOCKED_FOR_MAIN_PLAN_REFRESH`.
- Scope: chỉ P5-A — establish current module-request and path inputs.
- Đây là Coder evidence và precise blocker cho Main; không phải Supervisor verdict, không phải P5-A completion, và không mở P5-B/P5-C/P5-D.
- HEAD: `0aa49c87628c9e8b2041754515d6ebf0a930d55b`.
- Parent accepted boundary: `d1d8eb9002ce9c449c3713de0837ac8216d17a8d`.
- Không có production edit, test edit, plan/ledger edit, build/QA gate, target access, detect-changes, stage, commit, push, reset, hoặc checkout.
- Main-owned successor refresh trong Child 05 actual-status được giữ nguyên với SHA-256 `1DC9341F2D67974E3C659EDEFB6047E5AD55B84C7E7E175DA127C94D5B155488`.
- Ba Main handoff report có sẵn vẫn được giữ nguyên untracked.

## Authority Read

Đã đọc đầy đủ:

- `E:\Anvien\AGENTS.md`;
- `internal/aicontext/skills/working-rules/SKILL.md`;
- coder, backend-development và planner skill;
- bốn planner template;
- graph-accuracy roadmap;
- toàn bộ bốn Child 05 ledgers;
- `docs/contracts/graph-accuracy-contract.md`;
- fresh source owners được liệt kê bên dưới.

Không có `SPEC-MAP.md` hoặc `Docs/SPEC` cluster trong repo. `reports/coder/readme.md` cũng không tồn tại, nên report dùng cấu trúc durable của Coder handoff gần nhất mà không giả lập template bị thiếu.

## Invariant Family Map

| Surface | Owner / source | Current classification | P5-A consequence |
|---|---|---|---|
| Source-written TS/JS import request | `internal/providers/tsjs/imports.go` | named/default/alias/namespace name and raw module text exist; requested meaning/type-only does not | production input gap; plan decision required before edit |
| Accepted re-export syntax | `ScopeIR.ExportFact` from Child 04 | complete name, target name, three meaning lanes, type-only and provenance | consume as immutable input; do not regenerate or reinterpret |
| ScopeIR import contract | `internal/scopeir/facts.go::ImportFact` | carries names/module text plus legacy output-looking fields; no meaning/type-only/source range | partial contract; exact new representation is not authorized yet |
| Canonical clone/order | `internal/scopeir/ir.go`, `internal/scopeir/sort_keys.go` | only `TransitiveVia` is cloned/sorted for imports; `compareImport` has no meaning/type-only discriminator | conditional owners if P5-A adds a collection/scalar contract |
| Module/file result | `internal/resolution/indexes.go::resolvedImport` | module result is distinct from optional physical definition result | correct and preserve in P5-A |
| TS/JS path strategy | `resolveImportFiles` -> `resolveImportFile` | relative and index candidates resolve deterministically from current file set | correct for bounded path; preserve |
| Other-language path strategies | `internal/resolution/import_resolution.go` | shared multi-language behavior | preserve-only |
| Syntactic dependency projection | `emitImportEdges` | emits File -> File `IMPORTS` from `resolvedImport.TargetFiles` independently of `TargetDef` | correct separation; preserve |
| Export lookup / traversal / terminal binding | P5-B/P5-C/P5-D | not part of current tranche | locked |

Forbidden fallback status:

- Physical definitions must not become implicit export entries.
- Explicit import failure must not use repository-global same-name rescue.
- Compatibility re-export `ImportFact` must not replace accepted Child 04 `ExportFact` meaning/provenance.
- Target source and `E:\cheapapp.org` remain untouched.

Residual unverified surfaces: requested-meaning contract semantics and accepted count denominator remain unresolved by current plan authority; therefore `READY_FOR_SUPERVISOR` is not claimed.

## Fresh Graph Evidence

Commands:

```text
anvien --help
anvien analyze --force
anvien analyze --force --json
anvien analyze --force --benchmark-json .tmp\p5a-prechange-analyze.json --benchmark-label p5a-prechange-counts
```

All commands exited `0`. The repeated analyzes produced the same inventory:

```text
scanned=1,944
parsed_code=736
failed=0
nodes=114,738
relationships=157,553
```

Final graph identity:

- indexed/current commit: `0aa49c87628c9e8b2041754515d6ebf0a930d55b`;
- analyzed at: `2026-08-21T08:57:27Z`;
- graph path: `C:\Users\TAM NGUYEN\.codex\worktrees\a363\Anvien\.anvien\graph.json`;
- graph bytes: `456,926,068`;
- graph SHA-256: `F6F74BB49521795D0B1FBB8275E70C2B7E56DE8BCBDBBA1064AE7CE9FE5D8D5C`;
- every candidate file-detail result: `stale=false`, `changedSinceAnalyze=false`.

The debug-only benchmark capture was `9,382` bytes with SHA-256 `F6C0A4DE5DBD2D4E93CE3D5B7A0D1353FE25B0B35A20E34E7AF7DF32A9BC4E67`. Its relevant metrics are copied below; `.tmp` is not durable evidence authority.

## Fresh Input Manifest

### Source syntax -> `ImportFact`

| Source form | Current fact fields | Requested name today | Requested meaning today | Result |
|---|---|---|---|---|
| `import DefaultLocal from "m"` | `Kind=named`, `LocalName=DefaultLocal`, `ImportedName=default`, `TargetRaw=m` | exact `default` | absent | name correct; meaning missing |
| `import { Remote } from "m"` | `Kind=named`, `LocalName=Remote`, `ImportedName=Remote`, `TargetRaw=m` | exact `Remote` | absent | name correct; meaning missing |
| `import { Remote as Local } from "m"` | `Kind=alias`, `LocalName=Local`, `ImportedName=Remote`, `Alias=Local`, `TargetRaw=m` | exact `Remote` | absent | name/alias correct; meaning missing |
| `import * as LocalNS from "m"` | `Kind=namespace`, `LocalName=LocalNS`, `ImportedName=moduleNameFromTarget(m)` | module basename, not an exported name | only kind implies namespace; no explicit meaning | partial |
| `import type { T } from "m"` | same field shape as a non-type named import | exact `T` | type-only modifier is not retained | wrong for meaning |
| `import { type T } from "m"` | same field shape as a non-type named import | exact `T` | inline type-only modifier is not retained | wrong for meaning |
| `import "m"` | `emitImportStatement` returns when `import_clause` is absent | no fact | no fact | omitted sibling syntax; not proven necessary for bounded 2/2 and needs explicit plan disposition |

Current `ImportFact` fields are:

```text
ID, FilePath, FileHash, Kind, LocalName, ImportedName, Alias,
TargetRaw, TargetFile, TargetExportedName, TargetModuleScope,
TargetDefID, TransitiveVia, LinkStatus
```

There is no requested meaning/type-only field and no import source range/provenance field. Repository search found no production assignment to `ImportFact.TargetFile`, `TargetExportedName`, `TargetModuleScope`, `TargetDefID`, or `LinkStatus`; the live resolution result is the separate in-memory `resolvedImport`.

### Accepted Child 04 re-export fact -> compatibility `ImportFact`

| Accepted `ExportFact` | Compatibility import emitted today | Information retained/lost |
|---|---|---|
| `ExportReexport` | `ImportReexport`; local=exported name; imported=target exported name; alias when renamed | names/module retained; meanings/type-only/provenance remain only on `ExportFact` |
| `ExportStar` | empty-name `ImportWildcard` | target module retained; star meaning/type-only/provenance remain only on `ExportFact` |
| `ExportNamespace` | the same empty-name `ImportWildcard` | namespace exported name and namespace/type-only meaning are absent from compatibility import but remain on `ExportFact` |

P5-A/P5-B must therefore treat accepted `ScopeIR.Exports` as the sole re-export semantic input. Extending the compatibility import into a second re-export truth would violate the plan boundary.

### `ImportFact` -> module/file result

```text
ScopeIR.ImportFact
  -> NormalizeInPlace / compareImport
  -> buildWorkspace
  -> source Scope = moduleScopeByFile[ImportFact.FilePath]
  -> preprocessImportTarget(TargetRaw)
  -> resolveImportFiles(Language, FilePath, TargetRaw)
     -> Go package strategy, or language-specific strategy
     -> generic TS/JS resolveImportFile fallback
  -> resolvedImport {
       Fact,
       SourceScope,
       TargetFiles,
       TargetFile = TargetFiles[0],
       TargetDef = optional physical declaration,
       LinkStatus = "" when a module/file is found
     }
```

For TS/JS generic lookup, candidate order is exact and current-source-backed:

```text
base
base.ts
base.tsx
base.js
base.jsx
base/index.ts
base/index.tsx
base/index.js
base/index.jsx
```

Relative requests derive `base` from `path.Dir(sourceFile) + targetRaw`; non-relative requests use cleaned raw module text. The first file present in `workspace.fileSet` wins. Current source reads no `tsconfig` paths, package exports, or project-profile input in this TS/JS path. No evidence in P5-A authorizes a broad package-resolution rewrite.

`resolveImportedDef` is downstream of the valid module result. It searches only `defsByFile[TargetFile]` by imported/local name, and its default fallback selects the first class/function/interface physical declaration. That behavior is wrong for export lookup but belongs to P5-C after P5-B, not P5-A.

## Absolute Pre-Change Counts

Benchmark/runtime metric:

```text
resolution.ImportsResolved = 5,072
resolution.FinalizedImportsEmitted = 5,072
resolution.ImportUsesEmitted = 1,261
```

Persisted graph query:

```text
anvien cypher "MATCH ()-[r]->() WHERE r.type = 'IMPORTS' RETURN count(r) AS imports" --repo <isolated-worktree>
row_count=1
imports=5,088
```

Interpretation that must be retained, not collapsed:

- `5,072` is the physical target-file resolution count from the ScopeIR resolver and the number of resolver-emitted syntactic edges before graph-wide aggregation/deduplication.
- `5,088` is the absolute final persisted graph-wide `IMPORTS` relationship count across all graph producers.
- The current benchmark ledger has one ambiguous row named only `syntactic IMPORTS`. Main must either choose one authority explicitly or record both denominators. Recommended gate: P5-D proves delta `0` for both `resolution.FinalizedImportsEmitted` and final persisted graph `IMPORTS`.

## Fresh File Detail

| File | SHA-256 | Related files | Symbols | In / out / local | Unresolved | Linked flows / tests | File risk |
|---|---|---:|---:|---:|---:|---:|---|
| `internal/scopeir/facts.go` | `7F2B51D878F1541995AA884C438E18F1B1E6C72E20597E2C171FB36A59BFCA6A` | 247 | 192 | 1,232 / 27 / 174 | 5 | 0 / 100 | MEDIUM |
| `internal/providers/tsjs/imports.go` | `52AFBA2FC9A7ACEA314D5B39043A054F2671B2EF987161D8F1C3D641B2382749` | 17 | 185 | 7 / 107 / 40 | 497 | 1 / 4 | HIGH |
| `internal/scopeir/ir.go` | `732EE7F8959F077FED5550962A5369A020F6C9EC5ABDAF4540054C00E46E728C` | 245 | 57 | 1,147 / 38 / 89 | 159 | 3 / 98 | HIGH |
| `internal/scopeir/sort_keys.go` | `5C155B4C151D8E11833015376C26979C50928425A169CABAB475A65F52A52DB5` | 242 | 122 | 254 / 128 / 54 | 140 | 2 / 97 | HIGH |
| `internal/resolution/indexes.go` | `AA19B9D543012309A90974089BACBD0A122594C7481FFB0790DE5C01F3D3D76B` | 51 | 319 | 207 / 95 / 371 | 439 | 13 / 28 | HIGH |
| `internal/resolution/import_resolution.go` | `67C5F10E40E2DCD5850C56622A4D0ED3CBDFE15D84A9419971E97F8799876413` | 34 | 116 | 22 / 59 / 7 | 237 | 0 / 18 | HIGH |

## Complete HIGH / CRITICAL Blast-Radius Inventory

HIGH/CRITICAL is a scope warning, not an edit prohibition. Every fresh HIGH/CRITICAL result observed for the exact candidate owner set is included below.

### File-level upstream impact

| Owner | Risk | Impacted symbols | Direct | Affected files | Affected flows | Linked flows | Linked tests |
|---|---|---:|---:|---:|---:|---:|---:|
| `internal/scopeir/facts.go` | CRITICAL | 867 | 244 | 123 | 1 | 0 | 100 |
| `internal/providers/tsjs/imports.go` | HIGH | 18 | 18 | 1 | 0 | 1 | 4 |
| `internal/scopeir/ir.go` | CRITICAL | 610 | 442 | 76 | 1 | 3 | 98 |
| `internal/scopeir/sort_keys.go` | CRITICAL | 32 | 24 | 5 | 1 | 2 | 97 |
| `internal/resolution/indexes.go` | CRITICAL | 191 | 74 | 29 | 1 | 13 | 28 |
| `internal/resolution/import_resolution.go` | HIGH | 7 | 5 | 3 | 1 | 0 | 18 |

Exact file-impact path boundaries:

- `imports.go` affects only itself.
- `import_resolution.go` affects `internal/resolution/import_resolution.go`, `internal/resolution/indexes.go`, and `internal/resolution/legacy_import_language_conversion_test.go`.
- `sort_keys.go` affects `internal/scopeir/ir.go`, `internal/scopeir/position_index.go`, `internal/scopeir/scope_indexes_test.go`, `internal/scopeir/scopeir_test.go`, and `internal/scopeir/sort_keys.go`.
- `indexes.go` affects 29 files: `cmd/access-candidate-audit/main.go`; `internal/analyze/{analyze.go,analyze_test.go,legacy_resolver_conversion_test.go,pipeline_parity_test.go}`; `internal/cli/command.go`; `internal/graphaccuracy/access_candidate.go`; and 22 files under `internal/resolution` (`access_audit.go`, `access_audit_test.go`, `binding_accumulator_test.go`, `definition_collision_test.go`, `emit.go`, `identity_occurrence_test.go`, `import_resolution.go`, `indexes.go`, `legacy_heritage_map_conversion_test.go`, `legacy_import_language_conversion_test.go`, `legacy_p7_conversion_test.go`, `legacy_parity_test.go`, `legacy_scope_indexes_test.go`, `legacy_scope_symbol_semantics_conversion_test.go`, `legacy_suffix_index_ambiguity_conversion_test.go`, `p2b_persistence_test.go`, `p3c_binding_occurrence_test.go`, `proof_accuracy_golden_test.go`, `resolution_test.go`, `resolve.go`, `source_site.go`, `types.go`).
- `facts.go` and `ir.go` are shared ScopeIR owners. Their exact file-mode impacts span every production provider family, analyzer, resolution, selected graphaccuracy/framework/storage consumers, ScopeIR canonicalization, CLI probes, and their linked tests; the exact totals are `123` and `76` files respectively. The narrower exact-symbol results below are the edit gate, not the whole-file totals.

### Exact symbol upstream impact

| Symbol | Risk | Impacted symbols | Direct | Affected files | Modules | Processes |
|---|---|---:|---:|---:|---:|---:|
| `scopeir.ImportFact` | CRITICAL | 624 | 18 | 73 | 25 | 67 |
| `collector.emitImportStatement` | LOW | 0 | 0 | 0 | 0 | 0 |
| `collector.addImport` | LOW | 0 | 0 | 0 | 0 | 0 |
| `collector.addSourceExportFact` | LOW | 0 | 0 | 0 | 0 | 0 |
| `ScopeIR.Normalized` | CRITICAL | 21 | 11 | 6 | 4 | 15 |
| `ScopeIR.NormalizeInPlace` | LOW | 5 | 2 | 2 | 1 | 0 |
| `ScopeIR.NormalizeOwned` | LOW | 3 | 3 | 1 | 1 | 0 |
| `compareImport` | LOW | 6 | 1 | 2 | 1 | 2 |
| `resolvedImport` | CRITICAL | 95 | 6 | 16 | 2 | 43 |
| `buildWorkspace` | CRITICAL | 48 | 16 | 21 | 7 | 25 |
| `workspace.resolveImports` | CRITICAL | 27 | 1 | 15 | 3 | 19 |
| `workspace.resolveImportFiles` | CRITICAL | 18 | 1 | 11 | 1 | 5 |
| `workspace.resolveImportFile` | LOW | 3 | 1 | 1 | 1 | 1 |
| `workspace.resolveImportedDef` | CRITICAL | 18 | 1 | 11 | 1 | 5 |

All symbol-impact results report persisted resolution health `0` degraded nodes, `0` nodes with gaps, and `0` total ResolutionGap count for the selected graph evidence.

Critical exact-symbol affected-file sets:

- `ImportFact` — 73 files: two binding-contract probe files; three analyze files; two framework files; `internal/graphaccuracy/access_candidate.go`; `internal/lbugload/p3c_binding_occurrence_persistence_test.go`; provider extract/extract-test surfaces for Astro, C, C++, C#, Dart, Go, Java, Kotlin, PHP, Python, Ruby, Rust, SFC/Svelte, Swift, TSJS and Vue plus provider parity; resolution access/binding/definition/emit/identity/import/index/parity/persistence/proof/resolve/type surfaces; and `internal/scopeir/{ir.go,scopeir_test.go,sort_keys.go}`.
- `ScopeIR.Normalized` — `cmd/binding-contract-probe/{main.go,main_test.go}`, `internal/providers/tsjs/definition_position_inputs_test.go`, `internal/resolution/p3c_binding_occurrence_test.go`, `internal/scopeir/{ir.go,scopeir_test.go}`.
- `resolvedImport` — `internal/analyze/analyze.go` plus 15 resolution files: `access_audit.go`, `emit.go`, `identity_occurrence_test.go`, `import_resolution.go`, `indexes.go`, `legacy_heritage_map_conversion_test.go`, `legacy_import_language_conversion_test.go`, `legacy_p7_conversion_test.go`, `legacy_scope_indexes_test.go`, `legacy_scope_symbol_semantics_conversion_test.go`, `legacy_suffix_index_ambiguity_conversion_test.go`, `p2b_persistence_test.go`, `resolution_test.go`, `resolve.go`, `types.go`.
- `buildWorkspace` — `cmd/access-candidate-audit/main.go`; four analyze files; `internal/cli/command.go`; `internal/graphaccuracy/access_candidate.go`; and 14 resolution source/test files (`access_audit*`, definition/identity, five legacy regression families, P2-B/P3-C, `resolution_test.go`, `resolve.go`).
- `resolveImports` — `internal/analyze/analyze.go`, `internal/graphaccuracy/access_candidate.go`, and 13 resolution source/test files from access audit through `resolve.go`.
- `resolveImportFiles` and `resolveImportedDef` each affect the same 11-file set: `internal/resolution/access_audit.go`, `identity_occurrence_test.go`, `indexes.go`, `legacy_heritage_map_conversion_test.go`, `legacy_import_language_conversion_test.go`, `legacy_scope_indexes_test.go`, `legacy_scope_symbol_semantics_conversion_test.go`, `legacy_suffix_index_ambiguity_conversion_test.go`, `p2b_persistence_test.go`, `resolution_test.go`, `resolve.go`.

## Fresh Status Decision

P5-A does not proceed unchanged.

Fresh source preserves the original boundary between module/file lookup and export lookup, but it changes the executable work steps because the current plan assumes requested meaning is available while current `ImportFact` has no such data. Implementing now would force an unauthorized representation and semantic decision:

1. whether a normal named/default import requests one lane or an allowed set of value/type/namespace lanes;
2. how statement-level `import type` and inline `type` specifiers are represented;
3. whether namespace import meaning is explicit or only inferred from `ImportKind`;
4. whether the dormant `ImportFact.Target*` fields remain legacy/preserve-only or are replaced by a new module result contract;
5. whether side-effect-only imports are explicitly out of P5-A or must become facts;
6. whether the benchmark `syntactic IMPORTS` authority is resolver-emitted `5,072`, final persisted `5,088`, or both.

These are plan/work-step decisions, not safe Coder assumptions. The user gate requires stopping before code when fresh evidence changes work steps; therefore no production behavior was implemented.

## Exact Main-Owned Plan Refresh Required

Main should update only the stale P5-A portions while preserving the slice goal/order:

1. Actual status:
   - mark module text/file candidate/path result `correct / preserve-only`;
   - split requested-name state from requested-meaning state;
   - mark named/default/alias requested names `correct`, namespace requested name `partial`, and requested meaning/type-only `missing`;
   - record dormant `ImportFact.Target*` fields versus live `resolvedImport` result;
   - add fresh R2 graph/file-detail/impact/count evidence and replace the 2026-08-10 relationship counts.
2. P5-A touch map:
   - candidate edit: `internal/scopeir/facts.go` and `internal/providers/tsjs/imports.go`;
   - conditional canonicalization edit after representation choice: `internal/scopeir/ir.go` and `internal/scopeir/sort_keys.go`;
   - preserve-only: `internal/resolution/indexes.go` module/file behavior and `internal/resolution/import_resolution.go` other-language strategies;
   - keep `resolveImportedDef` deferred to P5-C and all export-table implementation deferred to P5-B.
3. P5-A work steps:
   - first authorize the exact requested-meaning/type-only contract and side-effect-import disposition;
   - explicitly require accepted `ExportFact` as the re-export meaning source and compatibility `ImportFact` as path compatibility only;
   - record both count denominators and require zero delta for both unless Main deliberately selects one with rationale;
   - only then authorize production-first implementation, focused tests, full build, and nearest real non-UI boundary in the later tranche.
4. Evidence/benchmark rows:
   - `E5-P5A-IMPACT1`: use the fresh six-file and fourteen-symbol inventory above;
   - `E5-P5A-INPUT1`: use the exact source-fact -> module-result manifest above;
   - `E5-P5A-COUNT1`: record `5,072` physical/resolver-emitted and `5,088` final persisted graph-wide `IMPORTS` without collapsing them.

## Handoff / Required Next Action

Main owns the plan/ledger refresh and exact re-authorization. After Main records the requested-meaning contract, count authority, exact editable owners and updated P5-A work steps, a fresh Coder lane may resume P5-A production-first implementation.

Until then:

- P5-A: open but `BLOCKED_FOR_MAIN_PLAN_REFRESH` before production code;
- P5-B/P5-C/P5-D: locked;
- build/QA/Supervisor/detect/commit: not reached;
- target access: forbidden and not performed;
- worktree: Main-owned changes preserved; this report is the only new durable Child 05 artifact.
