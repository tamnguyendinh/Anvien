# P2-B — Resolver/module first-divergence report

## Scope and verdict

This is a read-only, bounded root-cause investigation for the two fixed P1-C reproductions in `E:\cheapapp.org`. It does not modify Anvien production code, tests, the target worktree, or the target graph. The target was not copied into `E:\Anvien`.

The bounded verdict is **wrong at the resolver input/binding boundary**, with two distinct first-divergence families:

1. TypeScript ambient declarations are outside the ScopeIR workspace. `Promise`, `Math.max`, and `Math.min` are extracted as source facts, but the resolver has no TypeScript standard-library declaration/binding index to search. The unresolved result is therefore produced by the resolver, not by parsing the target source.
2. A consumer import whose target is a barrel is file-resolved, but `resolveImportedDef` only searches definitions physically owned by the target file. It does not follow the target module's re-export binding. The consumer's `TargetDef` stays nil, the scope binding is skipped, and the free call becomes an unresolved/low-confidence fallback. A direct-import in-memory control resolves both calls without changing source or Anvien code.

The `in_repo_unresolved` / `analyzer_gap` labels on the ambient cases are a separate downstream classification problem: graph-health classification recognizes Go builtins/qualifiers, not TypeScript ambient/global declarations. This report does not merge that projection issue into the resolver's first divergence.

No remediation is proposed or authorized by this report.

## Evidence boundary

| Item | Captured value |
|---|---|
| Target | `E:\cheapapp.org` |
| Target HEAD | `a869876ab6262dacde6cd5d432d099a91852a646` |
| Target graph | `E:\cheapapp.org\.anvien\graph.json` |
| Graph SHA-256 | `DD3A3A987E95C91787E7C70EF68BA5D2F953D24A32D35DEFE38B8E8E606FF090` |
| Graph inventory | 84,807 nodes; 114,125 relationships |
| Graph freshness | `stale=false`; hash unchanged after probes |
| Target worktree | Existing user changes remained; probes added no target artifacts |

The target graph has **zero File nodes ending in `.d.ts`**. `git ls-files '*.d.ts'` for the target also returned zero paths. The TypeScript oracle instead resolved ambient declarations in `E:\Anvien\node_modules\typescript\lib\lib.es*.d.ts`; those external declaration files are not target source and are not in the Anvien target graph.

Evidence IDs used below:

- `E2-P2B-BOUNDARY1`: graph hash/count/freshness and target worktree boundary.
- `E2-P2B-TARGET1`: fresh target `file-detail` results for all fixed files.
- `E2-P2B-ORACLE1`: TypeScript 5.9.3 ambient oracle; zero target-line diagnostics.
- `E2-P2B-ORACLE2`: TypeScript barrel-alias oracle; both aliases point to the declaration at line 10; zero line diagnostics.
- `E2-P2B-IR1`: direct in-memory extraction of the real target files.
- `E2-P2B-RESOLVE1`: actual barrel-chain resolution versus direct-declaration control.
- `E2-P2B-GRAPH1`: raw target graph selected-node/edge projection.
- `E2-P2B-IMPACT1`: self-graph `file-detail` and `impact` for every candidate owner listed in this report.

## Fixed source cases

### Case A — ambient TypeScript names

Independent source/oracle facts:

- `modules/commercial-config/server/admin-commercial-config/read-admin-commercial-config.ts:13` contains `Promise<AdminCommercialConfigReadModel>`.
- `modules/email/server/operations/email-operations-observability.ts:191` contains `Math.max(1, Math.min(...))`.
- The TypeScript 5.9.3 program built from the target `tsconfig.json` resolved `Promise` to declarations in `lib.es5.d.ts`, `lib.es2015.promise.d.ts`, and related ES libraries; `Math.max` and `Math.min` resolved to `lib.es5.d.ts`; `targetDiagnostics=[]`.

Fresh target graph observations:

- `file-detail` for `read-admin-commercial-config.ts`: parsed, one function symbol, one unresolved type-reference at line 13/column 3, `targetText=Promise`, `sourceSiteStatus=unresolved_local_binding`, `classification=in_repo_unresolved`, `actionability=analyzer_gap`.
- `file-detail` for `email-operations-observability.ts`: parsed; the exact two line-191 rows are `unresolved_call` for `Math.max` at column 9 and `Math.min` at column 21, with the same unresolved-local/analyzer-gap metadata.
- Raw graph nodes are `ResolutionGap:SourceSite:...#type-reference#Promise#13#3#13#10`, `...#call#Math.max#191#9#191#71`, and `...#call#Math.min#191#21#191#70`; each is connected by `HAS_RESOLUTION_GAP` from its real owning function.

Extraction evidence from the real target (no fixture or copy):

- `internal/providers/tsjs/types.go:163-186` walks identifier-like nodes and skips only the primitive `builtinTypeNames` map. `Promise` is not in that list, so a `TypeAnnotationFact` is emitted at line 13/column 3.
- `internal/providers/tsjs/references.go:15-32` recognizes a member call and emits `CallSiteFact{Name=max/min, ExplicitReceiver=Math, CallForm=CallMember}` with the exact line/range.
- The probe's IR ledger (`E2-P2B-IR1`) contains those exact facts and no parser error.

### Case B — barrel re-export binding

Independent source/oracle facts:

- `modules/commercial-config/server/admin-commercial-config/index.ts:21` re-exports `readAdminCommercialConfig` from `./read-admin-commercial-config`.
- `save-admin-commercial-config-mutation.ts:4-10` and `read-admin-commercial-config-route-view.ts:6-9` import the name from the barrel.
- The calls are at `save-admin-commercial-config-mutation.ts:142` and `read-admin-commercial-config-route-view.ts:32`.
- The TypeScript oracle follows both aliases to `read-admin-commercial-config.ts:10`, with no diagnostics on either line.

Fresh target graph observations:

- `index.ts` is parsed with zero definition symbols. It has a correct `IMPORTS` edge to the declaration and a correct barrel `USES` edge to the declaration (reason `scope-finalize import-use reexport readAdminCommercialConfig`).
- Each consumer has a correct `IMPORTS` edge to the barrel, but neither consumer has a `USES` or `CALLS` edge to the declaration. Each has a call resolution gap at the fixed call site; the save consumer also has the corresponding type-reference gap at the call expression, as exposed by the current graph projection.
- The selected raw graph edge ledger is in `E2-P2B-GRAPH1`; the full raw probe is retained under the P2-B temporary directory.

The real-source extraction probe recorded:

```text
barrel: ImportReexport local=readAdminCommercialConfig imported=readAdminCommercialConfig target=./read-admin-commercial-config
save consumer: ImportNamed local=readAdminCommercialConfig imported=readAdminCommercialConfig target=../../../commercial-config/server/admin-commercial-config
route consumer: ImportNamed local=readAdminCommercialConfig imported=readAdminCommercialConfig target=../../../commercial-config/server/admin-commercial-config
save call: CallFree readAdminCommercialConfig at 142:29
route call: CallFree readAdminCommercialConfig at 32:28
```

## First-divergence trace A: ambient declarations

### Pipeline trace

1. `internal/analyze/analyze.go:281-293,799-863` parses only the scanner's target files into `[]scopeir.ScopeIR` and does not add TypeScript compiler library declarations.
2. `internal/resolution/indexes.go:140-150` builds `defsByName` solely by iterating `ir.Definitions` from those parsed ScopeIRs. There is no ambient/lib.d.ts input or TypeScript language-service lookup in this path.
3. For `Promise`, `internal/resolution/resolve.go:321-334` does not treat it as a primitive (`isBuiltinType` at lines 380-386 excludes `Promise`), calls `w.resolveName`, and emits `type target not resolved` when the workspace index has no declaration.
4. For `Math.max/min`, `resolveCall` takes the `CallMember` branch at `internal/resolution/resolve.go:165-183`. `resolveMember` enters `resolveReceiverType`; for a non-`this` receiver, `internal/resolution/indexes.go:650-653` requires `lookupTypeBinding("Math", ...)`. The workspace has no ambient Math binding, so the member path returns false. There is no import receiver and the non-empty receiver prevents the global-call fallback; the call is emitted unresolved at lines 211-217.
5. `internal/resolution/emit.go:95-134` records the generic unresolved diagnostic as `sourceSiteStatus=unresolved_local_binding` with no proof kind. This is a faithful record of the resolver miss, but it does not distinguish “missing ambient declaration” from “missing in-repo declaration.”
6. `internal/graphhealth/diagnostics.go:223-258` fills missing metadata with `classifyDiagnostic`. Its classification tables/functions are Go-oriented (`isGoBuiltinOrPredeclared`, Go standard-library qualifiers). `Promise` and `Math.max/min` therefore fall through to `in_repo_unresolved` and `analyzer_gap`.

### First divergence

The earliest wrong boundary is **the resolver workspace's declaration universe**, before graph emission: TypeScript ambient declarations are not represented in the ScopeIR/`defsByName`/type-binding indexes. The exact surface where the miss becomes observable is `resolveTypeAnnotation -> resolveName` for `Promise`, and `resolveCall -> resolveMember -> resolveReceiverType -> lookupTypeBinding` for `Math`.

This is not an extractor omission: the IR has the source sites, and the compiler oracle proves the targets exist in external TypeScript library declarations. It is also not a stale-index issue: the target graph is fresh and hash-bound.

## First-divergence trace B: barrel re-export

### Pipeline trace

1. `internal/providers/tsjs/imports.go:73-114` correctly converts `export { readAdminCommercialConfig } from ...` into an `ImportReexport` fact. It does not materialize a definition node in the barrel (the barrel has no declaration of its own).
2. `internal/resolution/indexes.go:280-300` resolves the consumer's relative path to the barrel and the barrel's relative path to the declaration. This is why the graph contains both expected file-level `IMPORTS` edges and why the actual probe reports `ImportsResolved=10`.
3. For each import, `resolveImports` calls `resolveImportedDef` at lines 291-293. `resolveImportedDef` (`internal/resolution/indexes.go:460-483`) searches only `w.defsByFile[targetFile]` for a matching definition. For the consumer import, `targetFile` is the barrel and `w.defsByFile[barrel]` is empty. It does not consult the barrel's already-resolved `scopeBindings`/`ImportReexport` fact to continue to the declaration. Thus `resolved.TargetDef` remains nil.
4. The guard at `internal/resolution/indexes.go:302-304` skips adding the consumer's `readAdminCommercialConfig` to `scopeBindings` whenever `TargetDef == nil`. `importsByReceiver` still records the file link, which explains the apparently-correct consumer-to-barrel `IMPORTS` edge without a symbol binding.
5. The consumer call is a free call. `internal/resolution/resolve.go:184-209` looks through scoped, same-file, Go-package, and global name tiers; it never follows a re-export chain. The unresolved/fallback result is emitted at lines 211-217.
6. The barrel itself is not broken: its own re-export gets a declaration `TargetDef`, so `resolveImports` adds a `BindingReexport` and emits the barrel `USES` edge. The missing link is specifically **consumer → barrel re-export → declaration**.

### First divergence

The earliest wrong boundary is **`resolveImportedDef`'s target-definition lookup**, coupled to the `TargetDef == nil` scope-binding guard. File path resolution and extraction are already correct. The resolver has a re-export binding in the target module but treats the target module as if only its physical definitions were importable.

### Controlled in-memory differential

The P2-B probe changed no target file and no Anvien source. It cloned the extracted IR in memory and changed only the two consumer `TargetRaw` values from the barrel path to the declaration path:

| Run | Imports resolved | Import-use edges | Resolved calls | Unresolved references |
|---|---:|---:|---:|---:|
| Actual barrel chain | 10 | 1 | 37 | 542 |
| Direct declaration control | 10 | 3 | 39 | 540 |

In the control, both fixed calls resolve to the unique declaration with `proofKind=scope-binding`, `confidence=1.0`, and `CALLS` edges. The ambient gaps remain, as expected, because the control changes no ambient inputs. This isolates the barrel fault to re-export binding rather than parser, path, or call syntax.

## Candidate owners and blast-radius evidence

Every candidate owner used in the trace was checked with fresh self-graph `file-detail` and upstream `impact`. HIGH/CRITICAL values below are workflow warnings, not edit prohibitions; no edit is authorized here.

| Candidate | Source role | Impact result |
|---|---|---|
| `internal/resolution/indexes.go:buildWorkspace` | constructs `defsByName`, file/scope indexes | **CRITICAL**; 5 modules, 28 processes |
| `internal/resolution/indexes.go:workspace.resolveName` | type/global name lookup | **CRITICAL**; 2 modules, 41 processes |
| `internal/resolution/indexes.go:workspace.resolveGlobalName` | global declaration lookup | **CRITICAL**; 1 module, 25 processes |
| `internal/resolution/indexes.go:workspace.lookupTypeBinding` | receiver type lookup for Math/member calls | **CRITICAL**; 1 module, 14 processes |
| `internal/resolution/indexes.go:workspace.resolveImports` | import file/link and scope-binding assembly | **CRITICAL**; 3 modules, 18 processes |
| `internal/resolution/indexes.go:workspace.resolveImportedDef` | physical target-file definition lookup | **CRITICAL**; 1 module, 16 processes |
| `internal/resolution/indexes.go:workspace.resolveMember` / `resolveImportedMember` | member-call resolution fallbacks | **CRITICAL**; 2/3 modules, 23/33 processes |
| `internal/resolution/resolve.go:resolveTypeAnnotation` | Promise miss emission | **CRITICAL**; 4 modules, 35 processes |
| `internal/resolution/resolve.go:resolveCall` / `resolveAccess` | call/member miss emission | **CRITICAL**; 4 modules, 35 processes each |
| `internal/resolution/emit.go:emitter.emitUnresolvedReference` | diagnostic materialization | **CRITICAL**; 2 modules, 29 processes |
| `internal/graphhealth/diagnostics.go:classifyDiagnostic` | ambient classification projection | **HIGH**; 3 modules, 3 processes |
| `internal/graphhealth/diagnostics.go:normalizeDiagnosticMetadata` | applies classification/actionability | **CRITICAL**; 3 modules, 33 processes |
| `internal/providers/tsjs/types.go:collector.emitTypeReferences` | source type-fact extraction | LOW; no graph impact path found |
| `internal/providers/tsjs/references.go:collector.emitReferenceKind` | source call-fact extraction | LOW; no graph impact path found |
| `internal/providers/tsjs/imports.go:collector.emitExportStatement` | source re-export-fact extraction | LOW; no graph impact path found |

The LOW collector impact is itself only a graph traversal limitation (interface/dispatch edges are not visible in the self-graph); direct source inspection plus the in-memory IR proves those collectors emitted the needed facts. It is not evidence that the collectors were skipped at runtime.

Fresh self-file summaries used for the owner review:

- `internal/providers/tsjs/references.go`: parsed, 24 symbols, 70 unresolved, high file risk.
- `internal/providers/tsjs/types.go`: parsed, 49 symbols, 131 unresolved, high file risk.
- `internal/providers/tsjs/imports.go`: parsed, 21 symbols, 52 unresolved, high file risk.
- `internal/resolution/resolve.go`: parsed, 40 symbols, 121 unresolved, high file risk.
- `internal/resolution/indexes.go`: parsed, 192 symbols, 437 unresolved, high file risk.
- `internal/resolution/import_resolution.go`: parsed, 68 symbols, 237 unresolved, high file risk.
- `internal/resolution/emit.go`: parsed, 75 symbols, 289 unresolved, high file risk.
- `internal/graphhealth/diagnostics.go`: parsed, 47 symbols, 91 unresolved, high file risk.

## Alternatives ruled out

| Alternative | Evidence against it |
|---|---|
| Stale target graph | `file-detail` reports `stale=false`; graph hash/count are unchanged and bound to the recorded HEAD. |
| Target files were omitted by scanning | All five fixed files are present and parsed in the target graph; their exact gaps are attached to real function nodes. The P1-A eight-file omission set is a separate scanner issue. |
| TypeScript parser failed on the syntax | The real-source IR contains the exact Promise annotation, Math member calls, re-export, consumer imports, and call ranges. |
| Relative module path resolution failed | The graph has consumer→barrel and barrel→declaration `IMPORTS`; the in-memory control keeps `ImportsResolved=10`. |
| Barrel file itself lacks a usable re-export fact | The barrel `ImportReexport` is extracted, resolves to the declaration, and emits its own `USES` edge. |
| Consumer call form is unsupported | Calls are represented as `CallFree`; the direct-declaration control resolves both with normal scope binding. |
| A same-name declaration ambiguity caused the miss | The target declaration is unique in the bounded workspace; direct control resolves to that exact function. The actual failure occurs before a declaration is attached to the consumer binding. |
| Graph persistence dropped a valid resolver edge | The standalone resolver probe and persisted target graph agree: actual consumer calls are unresolved, while the barrel's own re-export edge is present. |
| The ambient names are in-repo declarations | Target Git has zero tracked `.d.ts`; target graph has zero `.d.ts` File nodes; the oracle declarations are external TypeScript library files. |

## Evidence strength and limitations

**Strong:** fresh target graph identity; exact source ranges; raw graph nodes/edges; direct parser/extractor run against the real target path; differential in-memory resolver control; TypeScript oracle with zero target-line diagnostics; source line trace through the resolver.

**Bounded:** the probe uses five selected target files rather than replaying the complete 1,359-file analyze pipeline. It is intended to isolate the fixed cases, not to estimate global TypeScript or re-export accuracy. The full target graph observations are still from the fresh 84,807-node graph.

**Unresolved boundaries:** this report does not determine the correct product policy for external/ambient declarations, whether all TypeScript lib names should become graph nodes, how namespace/static-member modeling should work for `Math`, or how arbitrary multi-hop/wildcard re-export cycles should be handled. It also does not assess downstream context/impact/process projections (P2-C scope), remediation design, build/test behavior, or product SPEC intent.

## Reproducibility artifacts and commands

All artifacts below are under `E:\Anvien`; none are in the target:

- `.tmp/cheapapp-graph-root-cause-restart/p2b/trace_resolver.go` — direct real-source extraction/resolution differential probe.
- `.tmp/cheapapp-graph-root-cause-restart/p2b/trace_resolver_output.json` — raw probe output.
- `.tmp/cheapapp-graph-root-cause-restart/p2b/target_selected_graph.mjs` and `target-selected-graph.json` — selected raw target graph records.
- `.tmp/cheapapp-graph-root-cause-restart/p2b/target-raw-graph-probe.json` — full fixed-site raw graph probe output.
- `.tmp/cheapapp-graph-root-cause-restart/p2b/typescript-oracle.json` — ambient declaration oracle rerun.
- `.tmp/cheapapp-graph-root-cause-restart/p2b/barrel-oracle.json` — barrel alias oracle rerun.

Key read-only commands:

```text
anvien file-detail <fixed-target-file> --repo cheapapp-accuracy-direct --json
anvien file-detail <candidate-owner-file> --repo Anvien --json
anvien impact symbol --uid <candidate-owner-uid> --repo Anvien --direction upstream --json
node .tmp/cheapapp-graph-root-cause-restart/p2b/target_selected_graph.mjs
go run .tmp/cheapapp-graph-root-cause-restart/p2b/trace_resolver.go
node .tmp/cheapapp-graph-root-cause-restart/p2b/../p1c-typescript-oracle.mjs
node .tmp/cheapapp-graph-root-cause-restart/p2b/../p1c-barrel-oracle.mjs
```

The two final `node` commands above refer to the existing P1-C scripts by their actual paths (`.tmp/cheapapp-graph-root-cause-restart/p1c-typescript-oracle.mjs` and `p1c-barrel-oracle.mjs`); their P2-B rerun outputs are copied only as command output into the P2-B temporary directory.

## Slice decision

`P2-B` is **bounded root-cause confirmed** for the fixed ambient and barrel cases. The findings remain untrusted until the root agent independently rechecks this report and the Supervisor accepts the evidence. No implementation change, test update, build, detect-changes run, or commit was performed or authorized.
