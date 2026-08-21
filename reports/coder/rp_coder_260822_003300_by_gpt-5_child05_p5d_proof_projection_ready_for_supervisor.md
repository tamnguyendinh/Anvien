# Child 05 P5-D Work Step 1 — export-binding proof projection

## Status and verdict

`READY_FOR_SUPERVISOR`

P5-D Work Step 1 is implemented and validated locally in the authorized E-only checkout. The existing `graph.Evidence` shape carries deterministic terminal, hop, and failure proof records from accepted P5-C semantic results through CALLS/ACCESSES emission, coalescing, Graph JSON, Ladybug, MCP context, and MCP impact. No typed Evidence/schema/reader expansion was needed.

This is a Coder handoff, not self-acceptance. P5-D remains open until an independent Supervisor verdict. Target Work Step 2 and `E:\cheapapp.org` remain untouched and locked.

## Authority and boundary

- Checkout/cwd: `E:\Anvien` only.
- Branch/HEAD: `master` / `26cb03eed3a72f1052f1af5de6a4de2f8326e794` (`docs(plan): authorize P5-D proof projection`).
- Accepted predecessor: P5-C implementation commit `76899d45a21fce55f6328b4cb30a6a5cb8719a81`.
- Accepted inventory/impact: `E5-P5D-IMPACT1` in `reports/coder/rp_coder_260821_232052_by_gpt-5_child05_p5d_pre_implementation_inventory.md`, identity `32,125 bytes / 393 LF / SHA-256 AECE504D0C447BE567C81D2651E47CC5FA2B6A012FE03809AD04F055E2C46269`.
- The accepted impact stayed valid because authority/owner source bytes did not change before implementation; inventory analyze/file-detail/impact was not rerun.
- No ledger edit, planner action, Supervisor call, detect-changes, stage/commit, push/reset/checkout, target access, alternate worktree/drive, or broad cleanup occurred.

## Invariant Family Map and blast radius

Family: accepted P5-C export result/proof -> retained import binding -> CALLS/ACCESSES Evidence -> deterministic relationship coalescing -> unchanged persistence/readers.

| Exact surface | Accepted risk and complete upstream tuple | P5-D decision |
|---|---|---|
| `workspace.resolveImports` | CRITICAL `34 / 17 / 3 / 17` | retain one phase-three semantic result; no second traversal/path pass |
| `workspace.resolveImportedDef` | HIGH `20 / 13 / 1 / 3` | proof-returning seam plus preserved wrapper |
| `workspace.resolveImportedMember` | CRITICAL `20 / 9 / 3 / 34` | proof-returning seam plus preserved wrapper |
| `resolvedImport` | CRITICAL `115 / 20 / 2 / 62` | bounded result/proof storage only |
| `resolveCall` | CRITICAL `27 / 11 / 6 / 32` | attach proof to existing CALLS Reference |
| `resolveAccess` | CRITICAL `27 / 11 / 6 / 32` | attach proof to existing ACCESSES Reference |
| `mergeRelationship` | CRITICAL `11 / 2 / 1 / 25` | deterministic Evidence union/dedupe only |
| `emitter.emitReference` | CRITICAL `10 / 5 / 2 / 34` | byte-stable transparent projection |
| `graph.Evidence` | CRITICAL `494 / 124 / 24 / 286` | preserve `Kind`, `Weight`, `Note`; no schema change |
| `writeGraphSnapshotJSON` | CRITICAL `22 / 7 / 5 / 21` | validate-only |
| `relationshipCSVRow` | CRITICAL `24 / 10 / 2 / 12` | validate-only |
| `contextRefPayload` | CRITICAL `8 / 3 / 1 / 13` | affected transparent reader, validate-only |
| `impactItemPayload` | CRITICAL `11 / 5 / 1 / 17` | affected transparent reader, validate-only |

All production owners are HIGH/CRITICAL surfaces. The change stayed at the accepted seams; broad `graph.Evidence`, persistence, HTTP/Web, and reader owners were preserved.

Forbidden fallback status: no physical-definition export inference, no global-name rescue, no second semantic/path pass, no duplicate table construction, no endpoint/key rewrite, and no proof on nonsemantic/no-proof paths.

Residual unverified surfaces inside Work Step 1: `none`. Target behavior is intentionally a separate locked Work Step 2.

## Production implementation

1. `internal/resolution/export_binding_proof.go`
   - Sole adapter from accepted `exportResolutionResult`/proofs to `graph.Evidence`.
   - Uses concrete versioned JSON structs, never maps.
   - Emits only `export-terminal-v1`, `export-hop-v1`, and `export-failure-v1`, all with `Weight=1` and `sourceSiteId`.
   - Terminal Note carries outcome/kind/requested name/canonical meanings/canonical target files/terminal def and Graph IDs/terminal and namespace files.
   - Hop Note carries deterministic proof/hop ordinals, outcome, full accepted `ExportFact`, hop kind/file/requested/target/member owner/member name/terminal ID.
   - Failure Note carries deterministic proof ordinal, outcome, failure file/name, namespace file, and canonical meanings.
   - Exact dedupe key is `(Kind, Weight, Note)`; generic evidence is stable-unioned first; P5-D records sort by kind rank, proof ordinal, hop ordinal, then canonical Note.
2. `internal/resolution/indexes.go`
   - Retains one owned semantic result on `resolvedImport` from the existing phase-three lookup.
   - Adds proof-returning definition/member seams while preserving existing wrappers and unaffected consumers.
   - Adds no traversal, path resolution, table construction, or second pass.
3. `internal/resolution/resolve.go`
   - `resolveCall` and `resolveAccess` preserve existing generic Evidence as item zero, then append proof from the exact retained semantic result that produced the selected target.
   - Endpoint, confidence, proof kind, target role/text, and source-site identity remain unchanged.
4. `internal/resolution/emit.go`
   - Only `mergeRelationship` changes, from Evidence replacement to the deterministic stable union above.
   - `emitter.emitReference` and `emitImportEdges` remain unchanged.

### Final production identities

| Path | Bytes | SHA-256 |
|---|---:|---|
| `internal/resolution/export_binding_proof.go` | 10,355 | `4BFE1976766FBF2E2257102070FF43CAC6D757E32E71DD583F8824781EAB6A8E` |
| `internal/resolution/indexes.go` | 34,478 | `CF26F90758C7D984423867B1122DE3076F2C541B4D7BB2C68EACFB676A6FD1B8` |
| `internal/resolution/resolve.go` | 21,949 | `57CE6C9DBDCF369ADD2F7B44F71ED7AF00B99736C11D40216A291DAA55F83493` |
| `internal/resolution/emit.go` | 26,748 | `E09EFD2A5601BE8D1143CC407916EB1206FA05DAEB995D2220A23EDF8B5869AF` |

## Focused test implementation

Tests were added only after production behavior existed.

| Path / direct anchors | Bytes | SHA-256 | Coverage |
|---|---:|---|---|
| `internal/resolution/export_binding_proof_test.go:13,109,160` | 9,669 | `AD3DCA9E82EACFB31137560636B59D426A063AEB967613E724D5EE3017AD5812` | direct/barrel retention and exact encoding; owned namespace/member hop; terminal/cycle failure ownership; ordering/dedupe/source-site conservation |
| `internal/resolution/resolution_test.go:1357` | 61,079 | `50B8B636AD9CA9DFC676179FF69D6DD2CA48517053117A9E4C70B65B181B0433` | emitted CALLS/ACCESSES proof, two-site coalescing, generic-first, endpoint/key conservation |
| `internal/lbugload/p5d_export_proof_persistence_test.go:16` | 4,963 | `C7FA2B925AC42DD0DBCCE7C86752A5A2CA5674D1D46EB816074DC49FF607F6EB` | Graph JSON -> relationship CSV -> Ladybug load endpoint/Evidence parity |
| `internal/analyze/analyze_test.go:24` | 38,998 | `83EF95668D161617DA21CE901B71A18850820197FC74CB7505EC62B815865802` | built analyzer Graph JSON proof survival |

The P5-D matrix also proves nonsemantic/no-proof paths add no export records and that distinct owner branches and source-site paths survive coalescing.

## Validation chronology

### Pre-build resolver boundaries

```text
go test ./internal/resolution -run '^TestP5D' -count=1 -v
PASS — package 0.562s

go test ./internal/resolution -count=1
PASS — package 0.234s
```

These commands ran on the final resolution package bytes. The later test-only correction touched only `internal/analyze/analyze_test.go`.

### Build-lock and canonical build chronology

The first canonical invocation failed only while replacing `E:\Anvien\anvien\bin\anvien.exe`. Windows Restart Manager identified exact holders PIDs `11440`, `14876`, `15440`, and `15788`; only those holders were terminated. The next canonical invocation passed on the then-current candidate (`1,975 / 743 / 0`, graph `116,322 / 160,446`).

Post-build validation then exposed one real test-fixture error: `TestP5DGraphSnapshotPreservesTerminalExportProof` passed a relative repo-local workspace to `analyze.Run`, whose existing contract requires an absolute path. Only `internal/analyze/analyze_test.go` changed, from SHA-256 `DAA4AADC3F1EFAF0F4B26652D31355468702268F4ECA76D5A19608A30E76D7CD` to final `83EF95668D161617DA21CE901B71A18850820197FC74CB7505EC62B815865802`, by resolving the same `E:\Anvien\.tmp` root with `filepath.Abs`. Production bytes and the other seven authorized test/source files did not change. This concrete test-byte invalidation required the final rebuild; it was not a redundant confirmation.

Before that final invocation:

- `anvien doctor locks --repo E:\Anvien --json`: analyze lock `free`.
- `anvien doctor processes --json`: only editor-owned global npm MCP was present.
- Exact Restart Manager inspection across six build artifacts found only PID `4156`, global npm `anvien.exe mcp`, holding the local/global `lbug_shared.dll` resources.
- Only PID `4156` was terminated. Its parent and all non-holders were preserved.
- Post-remediation Restart Manager returned `holderPids=[]` for all six artifacts and exclusive read/write no-share open succeeded `6/6`.

Final literal invocation:

```text
cwd E:\Anvien
powershell -ExecutionPolicy Bypass -File E:\Anvien\scripts\full-build.ps1
exit 0
```

Final built analyze result:

```text
scanned / parsed_code / failed = 1,975 / 743 / 0
nodes / relationships = 116,323 / 160,447
```

| Canonical output | Bytes | SHA-256 |
|---|---:|---|
| `E:\Anvien\anvien\bin\anvien.exe` | 71,509,504 | `BC060FDDF46394B72859B86930409B1B7E91C996BA304A86F0A633C37E043532` |
| `E:\Anvien\anvien\bin\lbug_shared.dll` | 20,230,656 | `20CBD87840483A2053CFF3FC2DB23A86DD802B8915D86509D41A4B709624CDB7` |
| `E:\Anvien\.anvien\graph.json` | 465,795,127 | `AAFAA260D9E2DBB49CF2D2503DFE3BA59477E5FB63E1F80469CD6AF3BD8365A1` |

All eight source/test hashes recorded above were rechecked after the final build and post-build tests; none changed. No further build or analyze was run.

### Post-build tests on final bytes

```text
go test ./internal/analyze -run '^TestP5DGraphSnapshotPreservesTerminalExportProof$' -count=1 -v
PASS — package 2.845s

go test ./internal/lbugload -run '^TestP5DExportProofPreservesGraphJSONAndLadybugRelationshipParity$' -count=1 -v
PASS — package 0.250s

go test ./internal/resolution -run '^TestP5C' -count=1 -v
PASS — package 0.245s; complete accepted P5-C matrix including no-global rescue and nonsemantic Go preservation

$env:TEMP='E:\Anvien\.tmp'; $env:TMP='E:\Anvien\.tmp'; go test ./internal/analyze -run '^(TestParseFilesRoutesGoFilesToGoProvider|TestParseFilesRoutesPythonFilesToPythonProvider|TestParseFilesRoutesJavaFilesToJavaProvider|TestParseFilesRoutesCSharpFilesToCSharpProvider|TestParseFilesRoutesRustFilesToRustProvider)$' -count=1 -v
PASS — package 0.528s; Go/Python/Java/C#/Rust provider routing unchanged
```

Every test temp root used for this final gate resolved under `E:\Anvien\.tmp`.

## Graph JSON, Ladybug, and affected-reader parity

The final Graph JSON contains:

```text
export-terminal-v1 items = 479
export-hop-v1 items      = 536
export-failure-v1 items  = 0 on the current repository corpus
```

The absence of a current-corpus failure edge is expected; focused tests prove cycle/missing/meaning/ambiguity failure projection. The current repository corpus has resolved proof-bearing CALLS only; focused tests prove the ACCESSES surface.

Exact current sample:

- Source: `AppContent` in `anvien-web/src/App.tsx`.
- Target: `useAppState` in `anvien-web/src/hooks/useAppState.local-runtime.tsx`.
- Relationship: CALLS, confidence `1`, proof kind `scope-binding`, target role/text `callable / useAppState`.
- Generic `scope-chain/useAppState` is item zero.
- Two distinct source sites (`34:32-34:45` and `68:6-68:19`) retain two terminal and two hop Notes after edge coalescing.

Ladybug query over `CodeRelation` returned `313` proof-bearing relationships, all current-corpus CALLS, with exactly the same `479 terminal / 536 hop / 0 failure` records. Runtime audits reported:

```text
generic-first violations = 0
exact (Kind,Weight,Note) dedupe violations = 0
```

The exact sample round-tripped the same source/target IDs, sourceSiteId/sourceSiteIds, proof kind, target role/text, and complete ordered Evidence string.

For actual MCP reader validation, the editor MCP transport had been intentionally stopped as the confirmed build DLL holder. A fresh canonical global npm MCP server (`C:\Users\TAM NGUYEN\AppData\Roaming\npm\node_modules\anvien\bin\anvien.exe mcp`) was therefore invoked directly via MCP JSON-RPC `initialize` and `tools/call`, from cwd `E:\Anvien` with temp rooted at `E:\Anvien\.tmp`:

- MCP `context` on the exact `useAppState` Function UID: exit `0`, status `found`, exposed `30 terminal / 30 hop / 0 failure` records; its first incoming CALLS entry is the exact generic-first two-site sample above.
- MCP `impact` on the same UID, upstream depth 1: exit `0`, exposed `12 terminal / 12 hop / 0 failure` records in `byDepth[1]`, including the same endpoints, source-site Notes, meanings, target file, terminal identity, and ExportFact.

No MCP/HTTP/Web reader production file changed.

## Fixed-corpus and IMPORTS preservation

Accepted artifact:

```text
E:\Anvien\.tmp\p5b-fixed-corpus.json
9,386 bytes
SHA-256 7687C5197AE76DC74294FB948AB9E25124BCEF26E444230EB29A2771EC373796
fixed parsed-code corpus = 736
physical target-file resolutions = 5,072
resolver-emitted syntactic IMPORTS = 5,072
persisted graph-wide IMPORTS = 5,088
```

The final full corpus has `743` parsed-code files: fixed `736` plus exactly two P5-B files, two P5-C files, and three new P5-D files. Accepted P5-C full-corpus counts were `5,161 / 5,161 / 5,177`.

Current persisted decomposition for the three P5-D files:

```text
outgoing internal/resolution/export_binding_proof.go                 = 9
outgoing internal/resolution/export_binding_proof_test.go            = 13
outgoing internal/lbugload/p5d_export_proof_persistence_test.go       = 21
incoming pre-existing corpus -> export_binding_proof.go               = 20
P5-D corpus growth                                                      = 63
```

The new production file has `21` current incoming edges; the complete source list shows exactly one comes from the new Lbugload test, leaving `20` pre-existing importers. The exact new-test-to-new-owner edge is `scope-finalize import named resolution`, so it is counted once in the new test's outgoing `21`, not twice.

Therefore:

```text
current full physical/emitted = 5,161 + 63 = 5,224 / 5,224
current persisted IMPORTS     = 5,177 + 63 = 5,240  (confirmed by Ladybug count query)

P5-B growth = 40
P5-C growth = 49
P5-D growth = 63
total added corpus growth = 152

5,224 - 152 = 5,072 physical target-file resolutions
5,224 - 152 = 5,072 resolver-emitted syntactic IMPORTS
5,240 - 152 = 5,088 persisted graph-wide IMPORTS
```

Fixed-corpus delta is exactly `0 / 0 / 0`. No physical import owner or `emitImportEdges` behavior changed.

## Preserve-only identities and behavior

| Preserve-only path | SHA-256 / evidence |
|---|---|
| `internal/resolution/export_resolution.go` | `566A69B965212D94E87E4D46703FC4E4C80E910FDC8657449D3914D5B868248F` |
| `internal/resolution/export_resolution_test.go` | `97FF4990066179AE779C9354D7A59859F3CA437826062DA0DA19C54776BC1620` |
| `internal/resolution/export_tables.go` | `BA4C7F9F2D05F6715CA22195C1CFF64C1AACADC3D81E7CA3ED711C8EA8B63E19` |
| `internal/resolution/import_resolution.go` | `67C5F10E40E2DCD5850C56622A4D0ED3CBDFE15D84A9419971E97F8799876413` |
| `internal/graph/types.go` | `9AA3AD89B0FB59F0AC295EF80D170EB17B26693014C473D9CC7E142E1163468D` |
| `resolveImportFiles`, `resolveImportFile`, `resolvedImport.TargetFiles` | no authorized diff; one existing path pass retained |
| `emitImportEdges`, `emitter.emitReference` | no diff; syntactic edge and transparent projection contracts retained |
| generic global/name helpers and P5-C no-global guard | no diff; complete P5-C test matrix PASS |
| Child 04 facts/providers, P5-B tables, P5-C traversal | unchanged and not reopened |
| graph schema/persistence/readers | no production diff; parity evidence above |

## Git/worktree boundary

- Branch/HEAD remained `master` / `26cb03eed3a72f1052f1af5de6a4de2f8326e794`, ahead of `origin/master` by 48 commits.
- Exactly eight authorized source/test paths are modified or untracked: the four production and four test paths listed above.
- `git diff --check` on the exact eight-path manifest: PASS.
- Staged/index paths: none.
- Twelve pre-existing untracked Main rotation handoffs under `reports/Investigation/` remained byte-untouched and unstaged; the latest is `rp_main_260822_000630_orchestration_rotation_handoff.md`.
- This immutable Coder report is the only additional lane-owned artifact.
- No existing Coder/Supervisor/Main report was modified.
- No ledger, detect-changes, stage/commit, push/reset/checkout, target, P5-D Work Step 2, `E:\cheapapp.org`, C-worktree, or alternate-drive action occurred.

## E2E verification summary

```text
[PASS] Compiled: final literal absolute canonical E-only build -> exit 0
[PASS] Primary: retained P5-C result -> concrete versioned terminal/hop/failure Evidence -> CALLS/ACCESSES
[PASS] Coalescing: generic first, exact tuple dedupe, deterministic sort, distinct source paths/owner branches retained
[PASS] Persistence: Graph JSON and Ladybug endpoint/Evidence parity
[PASS] Readers: actual MCP context and impact expose the same proof records
[PASS] Preservation: P5-C/no-global/unaffected-language regressions and fixed-corpus delta 0/0/0
[PASS] Boundary: target and all locked owners remained untouched
```

## Handoff

- Slice-local residual behavior risk: none identified by the authorized Work Step 1 matrix.
- Independent Supervisor must still inspect the HIGH/CRITICAL seams, exact encoding/order/dedupe behavior, build chronology, reader parity, fixed-corpus arithmetic, and current Git boundary.
- Exact next action: successor Main task `01a02542-a408-7643-88cf-cb0c14488b0b` submits this frozen report and the current eight-path candidate to the existing Supervisor lane. Target Work Step 2 remains locked until a separate Main authorization after Supervisor closure.
- Lane state after handoff: `IDLE / READY_FOR_SUPERVISOR`.

`READY_FOR_SUPERVISOR`
