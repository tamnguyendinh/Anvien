# P4-C Coder implementation report — graph and affected persistence projection

Report ID: `E4-P4C-CODER-260821-graph-persistence-projection`

Status: `READY_FOR_SUPERVISOR`

## Authority and boundary

- Lane: Child 04 P4-C only, opened by `E4-P4C-AUTH1`.
- Starting HEAD: `871189b8c6a4e4bb9ff538407232c913b8cf4db6`.
- Accepted P4-B1 at `42d167aaf28446ac0b3de479a8afefabb8d06736` was treated as closed and was not reviewed or reopened.
- Read before action: `E:\Anvien\AGENTS.md`; `E:\Anvien\.agents\skills\working-rules\SKILL.md`; `coder`, `backend-development`, and `databases` skills; orchestration handoff `reports/Investigation/rp_main_260821_0631_orchestration_rotation_handoff.md`; roadmap and all four Child 04 ledger files at their exact authority paths.
- No internal subagent was created. No access, read, stat, hash, analyze, or write was performed against the locked real target repository.
- Locked/non-claims: P4-C2, Child 05, terminal resolution, export-table/barrel traversal, alias-chain resolution, cycles, ambiguity, package public API, target oracle, ambient/external, scanner, UI/Playwright, and unrelated transport/cache/CLI/MCP changes.

## Initial Git boundary and protected ownership

The initial index was empty. Main-owned pre-existing changes were preserved without edit, stage, delete, overwrite, reset, checkout, commit, or push:

- Four modified planning documents: roadmap, P4-C plan, actual-status, and evidence.
- Main-owned orchestration handoff reports: `rp_main_260821_0631_orchestration_rotation_handoff.md` and later `rp_main_260821_0721_orchestration_rotation_handoff.md`.
- The benchmark ledger was read at its authority path and remained untouched.

Protected identity checks (bytes / LF / SHA-256) remained unchanged:

| file | bytes | LF | SHA-256 |
|---|---:|---:|---|
| roadmap | 24565 | 178 | `855058C75CD1C705F36ED17D850DA3B7DA18CDF693C6EA0CE48C48F57091EDF8` |
| P4-C plan | 38403 | 384 | `1FE0573CC0D7A895032C7B9190C8E814FB5DC68E085DFF18CA19866E680C8444` |
| actual-status | 24001 | 188 | `AE87E216334E63ECDD1D03E1DE19101DD8C8BCA52AAF5D7E5D2E5B886950AE59` |
| evidence | 33487 | 269 | `EA0C5DE109E40923B548EC888E50DAD16ECD103E1B592C0F993C8A0D8AE81615` |

Preserved handoff identities:

- `rp_main_260821_0631...`: 6542 bytes / 65 LF / `623FDC57BAC97F4C1F86F6A39C463E11F6BC0FFDA7DB8E9E661F0B0C1FFCC9EB`.
- `rp_main_260821_0721...`: 6542 bytes / 58 LF / `FDDAFFA421D64B10B2BBF6DDA11B8705E1E478BB358A4EF0EE3C497F3A5F019B`.

## E4-P4C-IMPACT1 — exact owner map and blast radius

After the required fresh `anvien analyze --force`, file-detail and upstream impact were run before edits for `internal/resolution/emit.go`, `internal/lbugload/csv.go`, `internal/lbugschema/schema.go`, `internal/filecontext/context.go`, and symbols `emitDefinitionNodes`, `nodeCSVRow`, `NodeSchema`, `exportedSymbol`, `emitDefinitionNode`, and `definitionNodeCompatible`.

Fresh file-detail summaries:

| owner | symbols | related files | inbound / outbound / local | unresolved | flows | tests | risk |
|---|---:|---:|---:|---:|---:|---:|---|
| `internal/resolution/emit.go` | 108 | 43 | 43 / 166 / 63 | 310 | 12 | 17 | HIGH |
| `internal/lbugload/csv.go` | 192 | 21 | 55 / 53 / 92 | 200 | 2 | 8 | HIGH |
| `internal/lbugschema/schema.go` | 41 | 21 | 52 / 8 / 29 | 50 | 1 | 6 | HIGH |
| `internal/filecontext/context.go` | 552 | 44 | 337 / 88 / 565 | 542 | 51 | 8 | HIGH |

Upstream file blast-radius tuples: `emit.go` CRITICAL `26/20/5/1`; `csv.go` CRITICAL `45/28/16/1`; `schema.go` CRITICAL `55/26/27/1`; `context.go` CRITICAL `252/134/32/1`. Exact symbol tuples: `emitDefinitionNodes` CRITICAL `6/1/4/33`; `nodeCSVRow` CRITICAL `3/1/1/12`; `NodeSchema` LOW `3/1/2/0`; `exportedSymbol` CRITICAL `9/3/3/24`; `emitDefinitionNode` CRITICAL `4/1/2/25`; `definitionNodeCompatible` CRITICAL `3/1/1/13`. HIGH/CRITICAL is a scope warning, not an edit prohibition.

Proven fact-to-owner map:

```text
ScopeIR.Exports / accepted ExportFact
  -> resolution.emitDefinitionNodes / exportProjectionNodes
  -> one Export graph node per fact + File CONTAINS Export
  -> lbugload.nodeCSVRow / export.csv
  -> lbugschema.NodeSchema / Export table + File->Export relation pair

same fact's local runtime meaning and LocalDefID
  -> definition node isExported compatibility field
  -> filecontext.exportedSymbol canonical read (legacy fallback only when absent)
```

No extraction owner, access-visibility owner, module-dependency edge, or unaffected consumer was edited.

## E4-P4C-SRC1 — implementation

Production behavior was implemented before tests:

- `internal/resolution/emit.go`: projects every accepted `ScopeIR.Exports` fact, preserves kind/name/meaning/type-only/ranges/provenance/raw target/local fact IDs, emits File→Export containment, validates same-file local references, and derives definition `isExported` only from the same runtime local fact. Legacy definition enrichment is retained while missing fact-derived `isExported` is inserted safely.
- `internal/lbugload/csv.go`: adds the complete Export node column contract and adds `isExported` to default definition persistence columns.
- `internal/lbugschema/schema.go`: adds Export node schema, File→Export relation pair, complete Export fields, and default definition `isExported` schema.
- `internal/filecontext/context.go`: canonical `isExported` takes precedence; legacy `exported`/visibility fallback is used only when canonical field is absent; malformed canonical data fails closed.

The CSV fixture was corrected to use a valid direct record and a separate valid re-export record; a direct export no longer carries `targetRaw`.

Focused tests added after production code:

- `internal/resolution/p4c_export_projection_test.go`
- `internal/lbugload/p4c_export_persistence_test.go`
- `internal/lbugschema/p4c_export_schema_test.go`
- `internal/filecontext/p4c_export_reader_test.go`

Existing owner tests/goldens updated only for the proven Export table/field contract, schema pair, and definition compatibility enrichment:

- `internal/lbugload/load_test.go`, `internal/lbugload/p2b_definition_persistence_test.go`
- `internal/lbugschema/p2b_persistence_schema_test.go`, `internal/lbugschema/schema_test.go`
- `internal/resolution/definition_collision_test.go`, `internal/resolution/testdata/typescript_graph_baseline_counts.golden.json`, `internal/resolution/testdata/typescript_graph_signature.golden.json`

Diff statistics (candidate boundary, excluding Main-owned planning/report files): production `+236/-10`; existing test/golden updates `+36/-17`; new focused tests `+383/-0`; aggregate candidate `+655/-27`. `git diff --check` passed.

Candidate file identities (bytes / LF / SHA-256):

| file | bytes | LF | SHA-256 |
|---|---:|---:|---|
| `internal/resolution/emit.go` | 27077 | 829 | `6CD67341CBB4778FC6693D0733B88AEF1E26D0775EBBBC99C375BCE65A2D37AB` |
| `internal/lbugload/csv.go` | 23176 | 668 | `48AFC8DD8828EC9BCFBB754BC8144F906C9C7D62DC9770826E802E2E93FC1C18` |
| `internal/lbugschema/schema.go` | 18933 | 541 | `5E163EB06DCA7C6BB6BC1DFF94FE9DAF56532037385D8A61F53583261A6715E5` |
| `internal/filecontext/context.go` | 51554 | 1741 | `9AA49A4E0784CCD84857E472F35CF3FA80385C22DC0F32277C944BFC5668B25E` |
| `internal/resolution/p4c_export_projection_test.go` | 8578 | 186 | `45D8B78ADAB2DF4D5897C4F001C9DAAAD9048760BF0E7DF2C44E537C4DDCBDCD` |
| `internal/lbugload/p4c_export_persistence_test.go` | 5376 | 113 | `93E26E5E57C3FE29FBDB64756A66015058A9F4DC90FFF406FD8E7854B9970BEC` |
| `internal/lbugschema/p4c_export_schema_test.go` | 1296 | 36 | `77C222B607BD68FEFDFAC9915C6DC9C51CE96A7280338C8EB1A89220E44258C6` |
| `internal/filecontext/p4c_export_reader_test.go` | 1387 | 48 | `8F16D938527AB1D72F5D2E893D6234DF50FEA7C2D9EE238EAD4D9E5EDD103F4C` |
| `internal/lbugload/load_test.go` | 16465 | 307 | `EDB45B063FBF577E744803B6BCE9B80C213511051FE5C3D0A543FCF3340CD285` |
| `internal/lbugload/p2b_definition_persistence_test.go` | 9012 | 240 | `6A1F12329DE3253AD24AA32E8AC83416E8FD22D9F4D58ED9508B32C26BF97799` |
| `internal/lbugschema/p2b_persistence_schema_test.go` | 3387 | 90 | `F43D94134B4D609D235634C146D5DA8EBF7342774560C81B8785F014E2FA8357` |
| `internal/lbugschema/schema_test.go` | 9337 | 311 | `AE93078221CDF1B9A062EDF200FD7A4B812C92546748775444F6C13505BAE853` |
| `internal/resolution/definition_collision_test.go` | 6989 | 206 | `7AF9B9E23AE0D0266739425420692F072330B237E987C65F6886F0CFA9C025A4` |
| `internal/resolution/testdata/typescript_graph_baseline_counts.golden.json` | 3056 | 110 | `0FD95205E651D67C50663DBEB5472512F12870100791F682175EAA19484CE536` |
| `internal/resolution/testdata/typescript_graph_signature.golden.json` | 8844 | 120 | `AA5EE9AAA52FA692D38D7378B96773932D09F1D64701D2740D03E090E0A7AE25` |

## E4-P4C-BUILD1 — canonical full build

Required pre-build gates:

```text
anvien doctor locks --repo E:\Anvien --json       -> status=free
anvien doctor processes --json                    -> no build process; editor/VS Code runtime only
```

One retry initially failed because the globally installed `anvien.exe` is a hardlink to `E:\Anvien\anvien\bin\anvien.exe` and editor-owned MCP children held that same file. The file ID was verified equal before remediation. Only those build-related MCP child/parent `cmd.exe` holders were terminated; `codex.exe` and VS Code were not terminated. A second lock/process check was free, and the final canonical command completed with exit code `0`:

```text
npm run full-build                                      -> exit 0
  npm install                                           -> passed
  npm run build / tsc -b / vite build                   -> passed
  anvien-launcher build                                 -> passed
  anvien analyze . --force                              -> passed
```

Final full-build analyze output: `scanned=1854`, `parsed_code=735`, `failed=0`, graph `113483` nodes / `155990` relationships. Vite emitted only existing chunk-size/dynamic-import warnings; no build failure.

## E4-P4C-TEST1 — focused and owner tests plus invariant measurements

Focused commands and results:

- `go test ./internal/resolution -run '^TestP4C' -count=1` -> PASS.
- `go test ./internal/lbugload -run '^TestP4C' -count=1` -> PASS.
- `go test ./internal/lbugschema -run '^TestP4C' -count=1` -> PASS.
- `go test ./internal/filecontext -run '^TestP4C' -count=1` -> PASS.
- `go test ./internal/resolution ./internal/lbugload ./internal/lbugschema ./internal/filecontext -count=1` -> PASS (all four packages).

The isolated two-file TypeScript fixture covered direct, alias, type-only, default, named re-export, star, namespace, and private negative-control forms. Independent parser→ScopeIR fact extraction and Graph JSON comparison measured:

| invariant | result |
|---|---:|
| accepted export facts | 15 |
| Graph Export nodes | 15 |
| export-fact conservation | `1.0` |
| canonical fact-key mismatches | 0 |
| direct runtime export count (`kind=direct`, value/namespace, non-type-only) | 2 |
| runtime local export facts | 3 |
| definition nodes marked `isExported=true` | 2 |
| compatibility drift | 0 |
| orphan local-definition references | 0 |
| negative-control failures (`privateValue`, `TypeOnly`, `RemoteShape`) | 0 |
| forbidden terminal/public-API/resolved-target state | 0 |
| export projection access/visibility fields | 0 |

## E4-P4C-BOUNDARY1 — real Graph JSON, Ladybug/schema/loader, and CLI reader

The boundary was run only on an isolated repo-local fixture under `.tmp\\p4c\\boundary-fixture` and was removed after evidence capture. Fresh command:

```text
anvien analyze E:\\Anvien\\.tmp\\p4c\\boundary-fixture --force --skip-git --no-gitignore --exclude AGENTS.md --exclude CLAUDE.md --exclude .agents/** --exclude .claude/** --json --benchmark-json E:\\Anvien\\.tmp\\p4c\\boundary-analyze-benchmark-final.json --benchmark-label p4c-boundary-final
-> exit 0; scanned=2, parsed=2, failed=0; graph=25 nodes / 29 relationships
```

The real analyze path exercised Graph JSON writing and the native Ladybug load phase. Captured benchmark summary (`p4c-boundary-final`): total duration `2,187,154,300 ns` (~2.187 s), DB node rows `25`, DB relationship rows `29`, node COPY tables `6`, relationship COPY pairs `7`, fallback inserts `0`, skipped relationships `0`.

Production-owner boundary harness commands (run before cleanup):

```text
go run .\\.tmp\\p4c\\boundary_measure.go .\\.tmp\\p4c\\boundary-fixture\\.anvien\\graph.json .\\.tmp\\p4c\\boundary-fixture
-> accepted=15, graph=15, conservation=1, compatibilityDrift=0, orphan=0, negativeControls=0, forbiddenState=0

go run .\\.tmp\\p4c\\boundary_loader.go .\\.tmp\\p4c\\boundary-fixture\\.anvien\\graph.json .tmp\\p4c\\boundary-csv-final
-> Export CSV data rows=15; persistenceFieldDifferences=0; Graph Export containment=15; nodeCopyCount=6; relationshipCopyCount=7; fallbackInsertCount=0; skippedRelationships=0; COPY Export=true; File->Export pair=true; relation declared=true
```

The generated Export CSV header contained all accepted fact/provenance fields (`kind`, `exportedName`, `localName`, `localDefId`, `localDefinitionNodeId`, `targetRaw`, `targetExportedName`, `meanings`, `typeOnly`, source/selection/statement ranges, `siteKind`, `appLayer`, `functionalArea`). `NodeSchema("Export")` matched the same fields and declared `PRIMARY KEY (id)`; definition schemas retained `isExported`.

CLI/read-owner commands against the same fresh fixture:

```text
anvien query runtimeValue --repo E:\\Anvien\\.tmp\\p4c\\boundary-fixture --json --limit 20
-> exit 0; Export matches include runtimeValue and renamedRuntime

anvien query remote --repo E:\\Anvien\\.tmp\\p4c\\boundary-fixture --json --limit 20
-> exit 0; re-export/namespace Export records are returned without terminal target resolution

anvien file-detail src/exports.ts --repo E:\\Anvien\\.tmp\\p4c\\boundary-fixture --json --relationships 20 --linked 20 --unresolved 20
-> stale=false, parser=parsed, unresolvedRefs=0; runtime exported=true, private=false, type-only=false
```

Boundary artifact identities before cleanup: Graph JSON 37786 bytes / 969 LF / `629A25A2448FD45DE523CF6E60FE509359191D4882BEE7E4053A3D3E164A836E`; benchmark JSON 4665 bytes / 209 LF / `AD10110163878C2D92777B8B422FEBDF525792D20B07B4B8E706D6CA5D503077`; Export CSV 5916 bytes / 16 LF / `A28AC9921DEB93B07B38AFF1E2BDD08BAA873B79C99BD676CB5B89616D688B0A`.

## Cleanup and current Git state

- Removed the complete debug/fixture trees `E:\Anvien\.tmp\\p4c` and `E:\Anvien\.tmp\\p4c-tests` after capturing measurements. Existing unrelated `.tmp` directories were left untouched.
- `git diff --check`: PASS.
- Index remains empty (`git diff --cached --name-only` produced no paths).
- No stage, commit, push, reset, or checkout was performed.
- Current worktree contains only the candidate production/test changes listed above, the four preserved Main planning modifications, the two preserved Main handoff reports, and this new untracked Coder report.

## Blockers, non-claims, and handoff

- No implementation blocker remains inside P4-C authority.
- The initial hardlink holder failure was resolved at the build gate and the canonical build subsequently passed.
- This report does not claim Supervisor review (`E4-P4C-REVIEW1`), Main-owned change detection (`E4-P4C-DETECT1`), or commit (`E4-P4C-COMMIT1`). Main owns ledger updates, final detect-changes, staging, commit, and later acceptance gates.
- No terminal module resolution, barrel traversal, alias-chain/cycle/ambiguity resolution, package public-API state, or real-target result is claimed.

## Report identity

- Path: `E:\Anvien\reports\Implementation\rp_coder_p4c_260821_graph_persistence_projection.md`
- Encoding/line ending: UTF-8 / LF
- Final identity (bytes / LF / canonical SHA-256): `16037` bytes / `223` LF / `1944C72E51FCCFBAF5BB77BDEC319EDEE11716B0292A2790AF623187EEAC1B3F`.
- SHA basis: the final report with only the 64-character value after `canonical SHA-256:` replaced by 64 ASCII zeroes (self-reference-safe and reproducible).

READY_FOR_SUPERVISOR
