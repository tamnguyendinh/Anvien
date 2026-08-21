# Child 05 / P5-B Export Tables Coder Report

## Status and Git Basis

- Trạng thái: `READY_FOR_SUPERVISOR / MAIN_HANDOFF`.
- Đây là Coder handoff, không phải Supervisor verdict, self-acceptance, detect-changes, hay commit.
- Scope duy nhất: build deterministic syntax-derived export tables from accepted Child 04 `ScopeIR.ExportFact` and accepted P5-A module/file results.
- Git basis: `E:\Anvien`, HEAD `17f3dad2587f2785d02ad266e188c7a8feca499b` (`docs(plan): authorize P5-B export table owner`).
- P5-B source/test changes remain uncommitted by instruction. Six Main orchestration handoff reports under `reports/Investigation/` remain untouched and untracked.
- No changes were made to Child 04 providers/facts/ScopeIR, `resolveImports`, `resolveImportedDef`, `import_resolution.go`, graph/persistence/readers, P5-C/P5-D, target, or any other drive/worktree.

## Authority and Invariant Family Map

| Surface | Authority / result | Boundary |
|---|---|---|
| Export syntax/meaning/provenance | accepted immutable Child 04 `ScopeIR.ExportFact` | copied exactly; no reinterpretation |
| Module/file candidates | existing P5-A `resolvedImport.TargetFiles` | consumed after existing `resolveImports`; path logic unchanged |
| Explicit export entries | `exportTable.Explicit` keyed by source-written `ExportedName` | all non-star facts, including direct/default/alias/re-export/namespace/type-only |
| Star adjacency | `exportTable.StarAdjacency` | `ExportStar` fact plus existing target-file candidates; no expansion/default synthesis |
| Physical definitions | preserve-only | never inferred as exports |
| Terminal traversal, precedence, ambiguity, cycles, meanings at lookup | P5-C | untouched and still locked |
| Terminal binding/projection/target | P5-D | untouched and still locked |
| Syntactic `IMPORTS` and physical path counts | existing resolver/analyze path | preserved; measured on fixed corpus |

Forbidden fallback status: no physical-definition export inference, no global-name rescue, no target special case, no second re-export semantic source, no side-effect-only fact expansion, and no terminal/proof outcome invented in P5-B.

Residual unverified surfaces: `none` inside the authorized P5-B invariant family. Supervisor, detect-changes, and commit are workflow gates intentionally pending Main control.

## Production Changes

1. `internal/resolution/export_tables.go` (dedicated semantic owner)
   - `exportTableEntry` stores an owning copy of the accepted `ExportFact` plus existing resolved module/file candidates.
   - `exportStarAdjacency` stores `ExportStar` provenance/meaning and target-file candidates without traversing.
   - `exportTable`/`exportTables` provide per-file syntax-derived surfaces.
   - `buildExportTables` indexes only accepted facts and existing `resolvedImport` results; it does not inspect `defsByFile` or resolve terminal Symbols.
   - Facts, ranges, target strings, meanings, and target paths are copied; entries and star edges are deterministically ordered.
2. `internal/resolution/indexes.go`
   - Added only `workspace.exportTables` storage.
   - Added only `w.buildExportTables()` immediately after existing `w.resolveImports()`; existing path and physical-definition code is unchanged.
3. `internal/resolution/export_tables_test.go`
   - Added after production code for zero-physical barrel, explicit/default/alias/re-export/namespace/star/type-only metadata, deterministic input-order behavior, owning-copy isolation, and `buildWorkspace` wiring.

Source line anchors:

- `export_tables.go:16-39` table data types;
- `export_tables.go:45-106` table construction and workspace seam;
- `export_tables.go:110-245` import association, deterministic ordering, and owning copies;
- `indexes.go:64` storage field and `indexes.go:175` wiring;
- `export_tables_test.go:11` zero-physical/determinism test, `export_tables_test.go:141` workspace wiring test, and `export_tables_test.go:204` explicit fact-kind/meaning test.

Final source hashes:

- `internal/resolution/export_tables.go` — `BA4C7F9F2D05F6715CA22195C1CFF64C1AACADC3D81E7CA3ED711C8EA8B63E19`;
- `internal/resolution/indexes.go` — `26DE75A911B4B1E1471C61548C4153749E067671DB34852E51448D52D4C0C486`;
- `internal/resolution/export_tables_test.go` — `A0395EF9030CB054B09C52AFBE048D2F8244BAD8A924EE492BB7AB312CD08AD8`.

## Fresh Graph and Blast-Radius Evidence

Pre-edit inventory was performed after the authorized E graph refresh. Existing-owner upstream impact was:

| Target | Risk | Impacted symbols | Files | Modules | Processes |
|---|---|---:|---:|---:|---:|
| `workspace` in `internal/resolution/indexes.go` | CRITICAL | 70 | 9 | 4 | 49 |
| `buildWorkspace` in `internal/resolution/indexes.go` | CRITICAL | 8 | 6 | 5 | 25 |

Post-build graph evidence for the new owner:

- `anvien file-detail internal/resolution/export_tables.go --repo E:\Anvien --json` → parsed, `stale=false`, `changedSinceAnalyze=false`, 62 symbols, 30 related files, 22 inbound / 39 outbound / 41 local relationships, 3 linked flows, 18 linked tests, file risk HIGH.
- Graph path: `E:\Anvien\.anvien\graph.json`; restored full graph after benchmark: `115,058` nodes / `158,074` relationships; full current corpus `738` parsed-code files / `0` failures.

The HIGH/CRITICAL values are scope warnings. The implementation remains limited to the dedicated table owner plus the two minimal `workspace/buildWorkspace` lines.

## Verification Evidence

### Focused behavior

```text
go test ./internal/resolution -run '^(TestBuildExportTablesZeroPhysicalBarrelAndDeterministicEntries|TestBuildWorkspaceWiresExportTablesWithoutPhysicalDefinitions|TestBuildExportTablesRetainsExplicitFactKindsAndMeanings)$' -count=1
PASS — resolution 0.182s (final focused matrix)

go test ./internal/resolution -count=1
PASS — resolution 0.212s (final source state)
```

The zero-physical barrel has no definitions but exposes explicit `run`/`ns` entries and one `ExportStar` adjacency with exact target files. The tests also prove no implicit `default`, deterministic input-order output, exact meaning/provenance, deep-copy isolation, and `buildWorkspace` wiring.

### Nearest real resolver/CLI boundary

```text
go test ./internal/resolution ./internal/analyze ./internal/cli -count=1
PASS — resolution 0.259s; analyze 2.904s; cli 119.566s
```

This is the nearest non-UI boundary for workspace construction, analyze, and the real CLI package. The final production source state is byte-identical to the source used by this run; the additional explicit fact-kind test was then run in the final focused/full resolution test commands below.

### Canonical E full build

Preflight: `anvien doctor locks --repo E:\Anvien --json` reported `status=free`; process inspection found only editor-owned MCP processes, no build holder.

Exact command, run once from the required cwd:

```text
cwd: E:\Anvien
powershell -ExecutionPolicy Bypass -File .\scripts\full-build.ps1
PASS — package/runtime, launcher/server, web bundle, CLI version 1.2.8, and final E analyze completed without error.
```

Canonical outputs:

- `E:\Anvien\anvien\bin\anvien.exe` — SHA-256 `CB8A79C4873E87C03F39FBB0A72116DDC599469E3E46661B3767022E349558B1`;
- `E:\Anvien\anvien\bin\lbug_shared.dll` — SHA-256 `20CBD87840483A2053CFF3FC2DB23A86DD802B8915D86509D41A4B709624CDB7`;
- `E:\Anvien\.anvien\graph.json` — restored full graph, SHA-256 `B724458DE88B997B3280CCB5BC9A407A6E45CC40548BEB35D97CF2B9B4F37115`.

### P5-A count preservation and benchmark

The full current corpus naturally includes the two new P5-B Go files and reports `ImportsResolved=5,112`, `FinalizedImportsEmitted=5,112`, persisted `IMPORTS=5,128`. The acceptance denominator is the same fixed `736` parsed-code corpus used by P5-A; the two new P5-B files were excluded only for this controlled measurement and then the full graph was restored.

Artifact: `E:\Anvien\.tmp\p5b-fixed-corpus.json` (9,386 bytes, SHA-256 `7687C5197AE76DC74294FB948AB9E25124BCEF26E444230EB29A2771EC373796`).

| Metric | P5-A baseline | P5-B fixed-corpus result | Delta |
|---|---:|---:|---:|
| parsed-code files | 736 | 736 | 0 |
| physical target-file resolutions (`resolution.ImportsResolved`) | 5,072 | 5,072 | 0 |
| resolver-emitted syntactic `IMPORTS` (`resolution.FinalizedImportsEmitted`) | 5,072 | 5,072 | 0 |
| persisted graph-wide `IMPORTS` | 5,088 | 5,088 | 0 |

Persisted query on the fixed-corpus graph:

```text
MATCH ()-[r]->() WHERE r.type = 'IMPORTS' RETURN count(r) AS imports
imports = 5,088
```

This proves P5-B table construction does not alter path lookup or syntactic dependency counts for the P5-A corpus. Full-corpus `+40` is attributable to the two newly parsed P5-B source/test files and is not used as the acceptance denominator.

## E2E Verification

```text
E2E Verification:
  [PASS] Compiled: powershell -ExecutionPolicy Bypass -File .\scripts\full-build.ps1 -> canonical E build completed
  [PASS] Runtime: built CLI/analyze and nearest resolver/analyze/CLI boundary -> PASS
  [PASS] Happy path: accepted ExportFact -> exportTable explicit entry/star adjacency -> zero-physical barrel table -> PASS
  [PASS] Edge case: no physical definitions, no implicit default, deterministic ordering, exact provenance/meaning, owning-copy isolation -> PASS
```

## Handoff Gates

- Production source and focused tests are ready for independent Supervisor review.
- P5-C/P5-D, target, and `E:\cheapapp.org` remain locked.
- No plan/evidence/benchmark/actual-status ledger was self-edited in this lane; Main owns the ledger refresh and Supervisor orchestration.
- `anvien detect-changes` and commit were intentionally not run per Main instruction.

`READY_FOR_SUPERVISOR`
