# Child 05 / P5-D Full Acceptance Review

Date: 2026-08-22 01:38:36 +07:00
Role: independent Supervisor, existing E-only lane
Authority checkout: `E:\Anvien`
Verdict: PASS

## Scope and decision

This review covers only P5-D: retain the accepted P5-C terminal export-resolution result and project its deterministic terminal, hop, and failure proof through CALLS/ACCESSES, relationship coalescing, Graph JSON, Ladybug, MCP context, and MCP impact, while preserving syntactic IMPORTS and all locked P5-B/P5-C/schema/reader boundaries.

The frozen eight-path candidate closes that invariant. No blocking source, proof, persistence, reader, target, or boundary defect remains. Planner/checklist advancement, detect-changes, staging, and commit remain Main-owned and were not performed here.

## Authority read and immutable inputs

Read in full before the decision:

- `E:\Anvien\AGENTS.md`.
- `.agents/skills/working-rules/SKILL.md`, `.agents/skills/supervisor/SKILL.md`, and `.agents/skills/Data-Integrity/SKILL.md`.
- All four living Child 05 ledgers under `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-05-module-export-and-reexport-resolution/`.
- Accepted P5-D inventory `reports/coder/rp_coder_260821_232052_by_gpt-5_child05_p5d_pre_implementation_inventory.md`, identity `32,125 bytes`, `393 LF`, SHA-256 `AECE504D0C447BE567C81D2651E47CC5FA2B6A012FE03809AD04F055E2C46269`.
- Work Step 1 report `reports/coder/rp_coder_260822_003300_by_gpt-5_child05_p5d_proof_projection_ready_for_supervisor.md`, independently remeasured as `18,358 bytes`, SHA-256 `97235D73735497328DD7DE41DB0B5023FC6A7ACC05CC46F362A965D0BDB0FB18`.
- Work Step 2 report `reports/coder/rp_coder_260822_011119_by_gpt-5_child05_p5d_target_validation_target_ready.md`, independently remeasured as `22,185 bytes`, SHA-256 `2740EDB71ADFFE1242E192DAEC86C70805F6B656568B49C95C089B981EFECDF2`.
- All four production owners and all four test owners in full, their exact tracked diffs, accepted P5-B table owner, accepted P5-C traversal/result owner, and the preserve-only persistence/reader owners needed to establish the contract.

`backend-development` was not invoked because this lane performed no implementation. Data-Integrity was used only as a review lens for proof ownership, persistence parity, and fail-closed behavior; it did not create another lane or artifact.

## Repository and candidate identity

The pre-report review boundary was rechecked after all focused commands:

- E branch/HEAD/parent: `master` / `26cb03eed3a72f1052f1af5de6a4de2f8326e794` / `fd6cb52f6258be2cbdaa622ee53c2d31d173566d`.
- E index entries: `0`; pre-report status entries: exactly `23`; `git diff --check`: clean.
- The status decomposition was exactly five tracked candidate files, three untracked candidate files, thirteen protected Main handoffs, and the two immutable Coder reports.
- Target branch/HEAD/parent: `master` / `a869876ab6262dacde6cd5d432d099a91852a646` / `79ab9b101cc21ec8da79dab724d435e87a6ea6f6`.
- Target index entries: `0`; target status entries: exactly the pre-existing `13`.

Final candidate bytes, rechecked after the Supervisor focused matrix:

| Path | Bytes | SHA-256 |
|---|---:|---|
| `internal/resolution/export_binding_proof.go` | 10,355 | `4BFE1976766FBF2E2257102070FF43CAC6D757E32E71DD583F8824781EAB6A8E` |
| `internal/resolution/indexes.go` | 34,478 | `CF26F90758C7D984423867B1122DE3076F2C541B4D7BB2C68EACFB676A6FD1B8` |
| `internal/resolution/resolve.go` | 21,949 | `57CE6C9DBDCF369ADD2F7B44F71ED7AF00B99736C11D40216A291DAA55F83493` |
| `internal/resolution/emit.go` | 26,748 | `E09EFD2A5601BE8D1143CC407916EB1206FA05DAEB995D2220A23EDF8B5869AF` |
| `internal/resolution/export_binding_proof_test.go` | 9,669 | `AD3DCA9E82EACFB31137560636B59D426A063AEB967613E724D5EE3017AD5812` |
| `internal/resolution/resolution_test.go` | 61,079 | `50B8B636AD9CA9DFC676179FF69D6DD2CA48517053117A9E4C70B65B181B0433` |
| `internal/lbugload/p5d_export_proof_persistence_test.go` | 4,963 | `C7FA2B925AC42DD0DBCCE7C86752A5A2CA5674D1D46EB816074DC49FF607F6EB` |
| `internal/analyze/analyze_test.go` | 38,998 | `83EF95668D161617DA21CE901B71A18850820197FC74CB7505EC62B815865802` |

Tracked numstat remains `97+/0-`, `1+/3-`, `38+/18-`, `70+/0-`, and `50+/14-` on the five tracked owners. No hidden candidate path was found.

## Source-level clearance by production owner

### `internal/resolution/export_binding_proof.go`

- It is the sole adapter from the accepted immutable `exportResolutionResult`/proof structures to the unchanged `graph.Evidence{Kind,Weight,Note}` shape.
- Concrete versioned JSON structs encode `export-terminal-v1`, `export-hop-v1`, and `export-failure-v1`; every projected record carries the exact source-site identity and weight `1`.
- Terminal records retain requested name, canonical meanings and target files, terminal definition/Graph identity, terminal file, and namespace file. Hop records retain deterministic proof/hop ordinals and the complete cloned accepted `ExportFact`, including owner/member provenance. Failure records retain the owned failure branch and canonical requested meanings.
- Generic evidence is retained before projected evidence. Exact tuple identity is `(Kind, Weight, Note)`. P5-D records sort deterministically by kind rank, proof ordinal, hop ordinal, and canonical Note. Distinct source sites and owner branches therefore survive while exact duplicates collapse.
- `retainedExportResolutionForScopedBinding` only retrieves the result already owned by the exact selected binding and verifies the same target Graph ID. It does not resolve a module path, rebuild an export table, or traverse exports again.

### `internal/resolution/indexes.go`

- `resolvedImport` retains one `SemanticResult` plus its presence bit from the existing phase-three semantic lookup.
- Phase one still resolves file candidates once; phase two still builds accepted P5-B tables once; phase three calls the proof-returning definition seam once and creates the same binding. No second path lookup, table build, or export pass was introduced.
- Existing nonsemantic wrappers remain available. Semantic member resolution returns the exact accepted P5-C result that selected the member; handled semantic failures remain fail-closed before physical fallback.

### `internal/resolution/resolve.go`

- The production diff is confined to `resolveCall` and `resolveAccess` proof attachment.
- Each source-site ID is computed once. Existing generic evidence is constructed first; proof is appended only when the retained semantic result resolves to the exact selected endpoint.
- Endpoint, confidence, relationship kind, proof kind, target role/text, and source-site contract remain unchanged.
- Constructor/free-call explicit-import guards remain in place. Generic `resolveName`, `resolveGlobalName`, `resolveGlobalCallName`, same-file, Go same-package, and low-confidence global behavior are unchanged.

### `internal/resolution/emit.go`

- The only production change is `mergeRelationship` replacing incoming-Evidence loss with deterministic stable union/dedupe.
- `emitter.emitReference`, `emitImportEdges`, semantic edge identity, endpoint selection, and source-site accumulation remain unchanged.
- Generic evidence remains first, preserving `relationshipCallName`; P5-D evidence cannot synthesize a new IMPORTS endpoint.

The HIGH/CRITICAL blast radius from the accepted inventory was treated as a scope warning. The actual edit remains confined to these four owners.

## Preserve-only clearance

Independent hashes matched the accepted identities for P5-B `export_tables.go`, P5-C `export_resolution.go` and `export_resolution_test.go`, `import_resolution.go`, `graph/types.go`, Ladybug CSV/query/schema owners, the Graph JSON writer, Web contract, MCP context/impact, filecontext, HTTP graph transport, and both file-detail UI owners. Representative exact anchors include:

- `export_tables.go` `BA4C7F9F...B63E19`.
- `export_resolution.go` `566A69B9...6248F` and its test `97FF4990...1620`.
- `import_resolution.go` `67C5F10E...76413`.
- `graph/types.go` `9AA3AD89...3468D`.
- `lbugschema/schema.go` `5E163EB0...15E5`.
- MCP context `63580482...804A` and impact `5AF78334...6377`.

There is no typed Evidence expansion, schema column change, new reader adapter, HTTP/Web/UI widening, P5-B mutation, or P5-C traversal mutation.

## Test-owner and focused final-byte evidence

The four test owners are source-backed rather than pass-by-default:

- Adapter tests decode exact terminal/hop/failure Notes, verify direct/barrel and owned-member paths, deterministic ordering, exact dedupe, cycle/failure ownership, and two-source-site conservation.
- Resolution tests verify emitted CALLS and ACCESSES endpoints, generic item zero, two-site coalescing, terminal/hop cardinality, and absence of false failures.
- Ladybug persistence test verifies Graph JSON to relationship CSV/load parity with zero skips/fallbacks.
- Analyzer test runs the real analyzer on a repo-local temporary workspace and compares the emitted relationship before and after Graph JSON persistence.

Supervisor lock preflight reported `E:\Anvien\.anvien\analyze.lock` free. Only editor-owned global npm MCP processes were present; none was terminated. With `TEMP` and `TMP` rooted at `E:\Anvien\.tmp`, these final-byte focused commands all exited `0`:

```text
go test ./internal/resolution -run '^TestP5D' -count=1 -v
go test ./internal/lbugload -run '^TestP5DExportProofPreservesGraphJSONAndLadybugRelationshipParity$' -count=1 -v
go test ./internal/analyze -run '^TestP5DGraphSnapshotPreservesTerminalExportProof$' -count=1 -v
```

All four P5-D resolver tests succeeded; the Ladybug and analyzer boundary tests also succeeded. Candidate hashes remained exact afterward.

## Work Step 1 build, artifact, and fixed-corpus evidence

The immutable chronology records the final literal canonical invocation, after the one analyzer-test path correction:

```text
cwd E:\Anvien
powershell -ExecutionPolicy Bypass -File E:\Anvien\scripts\full-build.ps1
exit 0
```

The final rebuild produced `1,975 / 743 / 0` scanned/parsed-code/failed and graph `116,323 / 160,447`. Current artifacts independently match those final bytes:

| Artifact | Bytes | SHA-256 |
|---|---:|---|
| `E:\Anvien\anvien\bin\anvien.exe` | 71,509,504 | `BC060FDDF46394B72859B86930409B1B7E91C996BA304A86F0A633C37E043532` |
| `E:\Anvien\anvien\bin\lbug_shared.dll` | 20,230,656 | `20CBD87840483A2053CFF3FC2DB23A86DD802B8915D86509D41A4B709624CDB7` |
| `E:\Anvien\.anvien\graph.json` | 465,795,127 | `AAFAA260D9E2DBB49CF2D2503DFE3BA59477E5FB63E1F80469CD6AF3BD8365A1` |

An independent streaming audit of that exact E Graph JSON found `313` proof-bearing relationships, `479` terminal records, `536` hop records, `0` current-corpus failure records, `0` generic-first violations, and `0` exact-tuple duplicate violations. Focused tests cover failure projection and ACCESSES, which the current repository corpus does not naturally exercise.

The fixed artifact `E:\Anvien\.tmp\p5b-fixed-corpus.json` remains `9,386 bytes`, SHA-256 `7687C5197AE76DC74294FB948AB9E25124BCEF26E444230EB29A2771EC373796`, and directly records `736` parsed-code files, `5,072` physical resolutions, and `5,072` resolver-emitted IMPORTS. Accepted persisted baseline is `5,088`.

The current Graph JSON contains `5,240` persisted IMPORTS. The accepted P5-B/P5-C corpus growth is `40 + 49`; the independently scanned P5-D growth is `9 + 13 + 21` outgoing from the three new files plus `20` pre-existing incoming edges to the new production file, or `63`. Therefore the total growth is `152`, and the fixed-corpus decomposition closes exactly:

```text
5,224 - (40 + 49 + 63) = 5,072 physical resolutions
5,224 - (40 + 49 + 63) = 5,072 resolver-emitted syntactic IMPORTS
5,240 - (40 + 49 + 63) = 5,088 persisted graph-wide IMPORTS
```

Fixed-corpus delta is `0 / 0 / 0`.

## Work Step 2 real-target evidence

The immutable chronology and Main trace establish exactly two target analyze invocations: the literal first invocation failed the Git dubious-ownership precondition before indexing; the sole process-scoped `safe.directory` retry exited `0` and produced `1,359` files, `93,562` nodes, and `127,516` relationships. No `--skip-git`, persistent config, ownership change, or third analyze occurred.

The current target graph independently rehashes to `440,884,239 bytes`, SHA-256 `061B6263A9D05D9C4702E290DECA22E526C06FE7B2F303E03342984E6D75857C`. All eight target oracle/config paths match the immutable manifest, including terminal implementation `44467B1D...F90AC`, barrel `0DECC0F2...A989F8`, route consumer `08FFD6F3...AC98`, and mutation consumer `C17330FF...0788`. Source lines independently confirm the terminal declaration at line 10, barrel re-export at line 21, and calls at route line 32 and mutation lines 142-144.

Direct raw-Graph inspection found exactly the two oracle CALLS, both targeting the same terminal Graph ID. Each has confidence `1`, source-site count `1`, status `resolved`, proof kind `scope-binding`, role `callable`, and exactly this ordered Evidence sequence:

```text
scope-chain, export-terminal-v1, export-hop-v1, export-hop-v1
```

The terminal Note retains requested name `readAdminCommercialConfig`, meaning `[value]`, the barrel target file, and the exact terminal definition/Graph identity. Hop zero is the complete barrel re-export fact; hop one is the complete direct export fact. Matching ResolutionGap count is `0` for both sites.

Four-reader parity was independently rechecked against the live persisted target artifact:

| Reader | Rows | Route Evidence SHA-256 | Mutation Evidence SHA-256 |
|---|---:|---|---|
| raw Graph JSON | 2 | `5D60F0037D40EB1C5ECED41271DE9F3AB7E13CA3EEEDEBDF8654930D186AA45B` | `18B3FC3F1046BE9C66CA06C1EA78A6A6F3D3E14DFA2563FB828CE7DB09FCA9B5` |
| Ladybug `CodeRelation` | 2 | same | same |
| MCP `context.incoming.calls` | 2 | same | same |
| MCP `impact.byDepth[1]` | 2 | same | same |

MCP context returned `status=found`. MCP impact returned `risk=MEDIUM`, `impactedCount=5`. Both exposed the exact two source sites and complete ordered Evidence arrays.

The target raw graph contains `2,625` persisted IMPORTS: `2,346` scope-finalize plus `279` markdown-link relationships. The only oracle consumer IMPORTS edges are the two consumer-to-barrel edges; consumer-to-implementation count is `0`. Together with unchanged target source/config bytes and unchanged physical-path/emitter owners, the immutable pre/post table closes physical/emitted/persisted delta at `0 / 0 / 0`.

Fresh-shell `GIT_CONFIG_*` count is `0`; target local `safe.directory` entry count is `0`.

## Exact invariant closure

1. One file-candidate resolution sequence feeds P5-B tables once and terminal binding; no second path lookup, table construction, or export traversal was added.
2. The accepted P5-C result remains dedicated and immutable through terminal/failure projection with complete owned proof provenance.
3. Generic Evidence is first; versioned encoding, order, and exact tuple dedupe are deterministic; distinct owner and source-site branches survive.
4. CALLS/ACCESSES attach proof only to the exact selected target; explicit-import no-global rescue and generic/global/name helpers remain preserve-only.
5. Syntactic IMPORTS, physical resolution, and nonsemantic behavior retain their accepted boundary.
6. Graph JSON, Ladybug, MCP context, and MCP impact expose byte-identical proof arrays.
7. Real target terminal equality is `2/2`, complete chains are `2/2`, matching false gaps are `0`, and no synthetic physical IMPORTS edge exists.
8. No schema/typed Evidence/new-reader/UI/P5-B/P5-C widening, hidden path, stale acceptance fixture, or unresolved same-invariant surface was found.

## Commands intentionally not run

- No full build was repeated because final source/test hashes and all canonical artifacts remained exact; there was no concrete invalidation.
- No E analyze and absolutely no target analyze was run.
- No accepted inventory/file-detail/impact refresh was repeated.
- No full-package or unrelated test sweep was repeated beyond the focused final-byte matrix above.
- No detect-changes, planner/checklist edit, source/test/ledger edit, stage, commit, push, reset, checkout, cleanup, target mutation, persistent Git config, HTTP/Web/UI/Playwright action, or alternate worktree/drive action occurred.

## Handoff

P5-D is source- and evidence-complete for Supervisor acceptance. Main may now perform only its separately owned ledger/detect/isolated-commit transition. The Supervisor lane returns to `IDLE / accepted` after this report is sealed.
