# Child 05 / P5-D Terminal Binding and Proof — Pre-Implementation Inventory

## Status and verdict

- Status: PRE_IMPLEMENTATION_READY / READY_FOR_MAIN_AUTHORIZATION.
- Verdict: P5-D is not a no-code slice. Accepted P5-C already selects the terminal definition for direct and barrel imports at the resolver boundary, but its exportResolutionResult and ordered proof hops are discarded before CALLS/ACCESSES emission. The graph therefore records the terminal endpoint where resolution succeeds, but cannot retain or expose the re-export proof chain required by P5-D.
- Required next authority: Main must planner-refresh and lock the exact proof-retention/projection encoding plus the bounded owner manifest below before any production edit. This inventory does not authorize implementation or target access.
- Current responsible Main: task 01a024ef-2740-7802-a9cb-71eb5d951496. Successor task 01a02518-d4a9-7081-860b-b2e5dddde93e remains routing-only until an OFFICIAL AUTHORITY TRANSFER.
- Lane state after handoff: IDLE / READY_FOR_MAIN_AUTHORIZATION.

## Authority and hard boundary

- Authoritative checkout: E:\Anvien only.
- Branch / HEAD: master / fd6cb52f6258be2cbdaa622ee53c2d31d173566d.
- Accepted P5-C implementation: 76899d45a21fce55f6328b4cb30a6a5cb8719a81; verified as an ancestor of current HEAD.
- P5-D is the sole open Child 05 slice. Pn-A+, Child 06, P5-C reopen, target execution, and E:\cheapapp.org remain locked.
- This gate was read-only outside this one Coder report. No production, test, fixture, ledger, existing report, target, or alternate-worktree file was edited.
- No build, QA/test, detect-changes, stage, commit, Supervisor, push, reset, checkout, cleanup, target command, C-worktree command, or subagent was run.
- Existing E:\Anvien\.tmp captures were neither created, edited, consumed as acceptance evidence, nor cleaned.

## Required reading completed

The following were read in full before the inventory:

1. E:\Anvien\AGENTS.md.
2. E:\Anvien\.agents\skills\working-rules\SKILL.md.
3. E:\Anvien\.agents\skills\coder\SKILL.md.
4. E:\Anvien\.agents\skills\backend-development\SKILL.md.
5. E:\Anvien\.agents\skills\Data-Integrity\SKILL.md.
6. All four living Child 05 ledgers under docs\plans\2026-07-26-anvien-graph-accuracy-multi-plan\2026-07-28-05-module-export-and-reexport-resolution.
7. All four Child 02 persistence/reader ledgers under docs\plans\2026-07-26-anvien-graph-accuracy-multi-plan\2026-07-28-02-current-graph-persistence-and-reader-consistency.
8. Accepted P5-C source reality and final Supervisor PASS report.

Current ledger identities:

| Ledger | Bytes | SHA-256 |
|---|---:|---|
| Child 05 actual-status | 35,004 | 90EC70F898CD649C00BF46D805CCDD429610612738201FEAEE133467C9DEC38D |
| Child 05 benchmark | 9,852 | 7F93E8D7AE8085CE81620E1ECE0C4FC0A3CF3C849A4A634778CB553E23E6FB57 |
| Child 05 evidence | 41,981 | 26B6EFE7121F7CEF4AA0E13A9372B57565AD920FF242CA4C0966B8CA9F80CA4C |
| Child 05 plan | 39,707 | 235E9AC1D5A63D9BE0660E1767B0590E447188ABF93D2EB1DE94348D062B2E25 |
| Child 02 actual-status | 40,370 | 7C5C8F419A04C4FB35F358B73D57D730D3313F3FBE4C41B6B454F1C12C0F100C |
| Child 02 benchmark | 8,374 | E5ED894809CFBF4D6A6BAB790BD6796B882F41D056AA4A33101BB69AE42AA2DC |
| Child 02 evidence | 57,718 | 40A944D6C3FD64A1DF32386DBF64F1D72B923A7F7E8C3DB9D5E1C0BB333F532C |
| Child 02 plan | 43,311 | 7B5DBB4BC0FEF221A74BB1E2F412194F2B7A8613BBA656141C624A0FE3AFFB4B |

Accepted P5-C PASS anchor:

- E:\Anvien\reports\Supervisor\rp_supervisor_260821_223035_by_gpt-5_child05_p5c_ambiguous_owner_member_resubmission_pass.md
- 8,607 bytes
- SHA-256 AC600E17A023E58261355C18647FC17674DCB1E2238258F9CE6941ABD49739DA

## Command-surface and fresh graph gate

Anvien command discovery was confirmed from cwd E:\Anvien with:

    anvien --help

Exactly one P5-D graph refresh was run before every graph query:

    cwd: E:\Anvien
    command: anvien analyze --force
    exit: 0
    scanned / parsed_code / failed: 1,969 / 740 / 0
    graph: E:\Anvien\.anvien\graph.json
    nodes / relationships: 115,947 / 159,626
    graph bytes: 462,444,449
    graph SHA-256: 014DC02974092E556FE270CC6BC524A70A282E2DB45CB423CAB8942D3A5A6B7E
    indexed commit: fd6cb52f6258be2cbdaa622ee53c2d31d173566d
    stale: false

No analyze, file-detail, impact, detect, or other graph gate was rerun after this accepted fresh corpus.

## Current graph facts

Relationship inventory:

| Relationship | Total | Evidence present |
|---|---:|---:|
| CALLS | 11,467 | 11,467 |
| ACCESSES | 5,974 | 5,974 |

Current proofKind distribution:

| Relationship | proofKind | Count |
|---|---|---:|
| CALLS | go-same-package | 2,863 |
| CALLS | import-member | 837 |
| CALLS | receiver-member | 1,364 |
| CALLS | same-file | 5,713 |
| CALLS | scope-binding | 690 |
| ACCESSES | import-member | 4 |
| ACCESSES | receiver-member | 5,584 |
| ACCESSES | scope-binding | 386 |

Language-specific findings:

- TypeScript/JavaScript import-member CALLS/ACCESSES: 0.
- TypeScript/JavaScript scope-binding CALLS: 629.
- TypeScript/JavaScript scope-binding ACCESSES: 381.
- Cross-file TypeScript/JavaScript scope-binding CALLS: 313.
- Sample evidence on these edges remains one generic record of the form kind=type-binding or scope-chain, weight=1, note=target name. No accepted P5-C export hop, terminal outcome, ambiguity branch, cycle branch, missing branch, or meaning-mismatch branch is represented.

Interpretation:

- A free imported alias is emitted as scope-binding rather than import-member, so proofKind cannot reconstruct whether the binding was direct, aliased, or reached through one or more barrels.
- Every current CALLS/ACCESSES relationship having Evidence proves transport presence only. It does not prove that P5-C proof data reached that Evidence.
- The target 2/2 result remains unknown because target access is intentionally locked. No target verdict is inferred from this repository graph.

## Current-reality pipeline classification

    accepted P5-C exportResolutionResult / exportResolutionProof / Hops
      -> indexes.go definition() reduction to defRef
      -> resolvedImport.TargetDef and scopeBindings
      -> resolveCall / resolveAccess
      -> resolution.Reference.Evidence
      -> emitter.emitReference
      -> graph.Relationship.Evidence
      -> Graph JSON and Ladybug evidence STRING
      -> transparent MCP / HTTP / Web transport

| Stage | Classification | Evidence and consequence |
|---|---|---|
| P5-C traversal and terminal selection | already-correct at synthetic resolver boundary | export_resolution.go owns deterministic terminal, ambiguity, cycle, missing, meaning mismatch, candidates, failures, and ordered Hops. Accepted export_resolution_test.go direct/barrel fixture emits two CALLS to the same terminal. |
| Import result retention | unbound | resolvedImport at indexes.go:25-32 stores TargetDef but no exportResolutionResult/proof. |
| Direct/free import binding | partial | resolveImports phase 3 at indexes.go:281-332 calls resolveImportedDef and stores only the definition. BindingRef.Via retains a resolvedImport pointer, but that object has no semantic proof. |
| Imported definition seam | unbound proof | resolveImportedDef at indexes.go:477-508 calls resolveImportExport(...).definition(); definition() drops candidates/failures/hops. |
| Namespace/member seam | unbound proof | resolveImportedMember at indexes.go:629-667 calls result.definition(); the P5-C result is discarded after selection. |
| Call endpoint emission | endpoint already-correct when binding exists; proof missing | resolveCall at resolve.go:360-509 emits generic scope-binding/import-member evidence only. Accepted explicit-import no-global-rescue guard remains intact. |
| Access endpoint emission | endpoint already-correct when binding exists; proof missing | resolveAccess at resolve.go:534-593 emits generic binding evidence only. |
| Relationship projection | transport-ready | emitReference at emit.go:130-172 copies Reference.Evidence unchanged into Relationship.Evidence. No adapter is needed here. |
| Coalesced-edge proof conservation | partial / unsafe for multi-proof paths | mergeRelationship at emit.go:669-720 replaces existing Evidence with incoming Evidence. It merges source-site IDs but does not deterministically union proof records. |
| Graph JSON persistence | already transparent | writeGraphSnapshotJSON generically serializes graph.Graph, including Relationship.Evidence. |
| Ladybug persistence | already transparent | relationshipCSVRow and relationshipEvidence marshal Evidence as JSON into the existing evidence STRING column; COPY and fallback query already carry that column. |
| MCP context and impact | already transparent affected readers | both payloads expose the full graph.Relationship, including Evidence. They require validation, not a production adapter, if the existing Evidence shape is retained. |
| HTTP/Web graph transport | transparent | graph payload and generated contract already carry graph evidence kind/weight/note. No semantic interpretation currently occurs. |
| File-detail and FileDetailPanel | not an Evidence consumer | these expose proofKind/source-site status, not Relationship.Evidence. They are preserve-only unless Owner explicitly opens a new display requirement. |
| Child 02 C09-C16 | not the automatic P5-D denominator | those accepted Definition identity/selection/semantic-search readers do not interpret CALLS/ACCESSES export proof. Preserve their accepted behavior. |

## Source identities

| Source | Bytes | SHA-256 | P5-D disposition |
|---|---:|---|---|
| internal\resolution\export_tables.go | 7,213 | BA4C7F9F2D05F6715CA22195C1CFF64C1AACADC3D81E7CA3ED711C8EA8B63E19 | accepted P5-B preserve-only |
| internal\resolution\export_resolution.go | 26,015 | 566A69B965212D94E87E4D46703FC4E4C80E910FDC8657449D3914D5B868248F | accepted P5-C inspect/preserve |
| internal\resolution\indexes.go | 33,417 | A44C94FA545FFBAEF7017D49C6529B1B4CFDE82C3BB244EC6F4DD57EBE4283F6 | proposed bounded proof-retention owner |
| internal\resolution\resolve.go | 20,799 | 047CCDAE8F451553FC159CCF5B9864FFCA60F597E25FF31BDF5494638EC5817C | proposed bounded call/access projection owner |
| internal\resolution\emit.go | 26,772 | B1B8F22AEFE94F72483CD284B7E267FE3897412C04D581AFAB7F9FE30F0B2867 | proposed bounded deterministic Evidence-union owner |
| internal\graph\types.go | 8,113 | 9AA3AD89B0FB59F0AC295EF80D170EB17B26693014C473D9CC7E142E1163468D | preserve existing Evidence shape |
| internal\lbugload\csv.go | 23,176 | 48AFC8DD8828EC9BCFBB754BC8144F906C9C7D62DC9770826E802E2E93FC1C18 | validate-only persistence |
| internal\lbugload\queries.go | 5,394 | 5DF71E74B0A3E514C973CE03B0C3714B407BE231022CE9995BB1887F0A8C961C | validate-only persistence |
| internal\lbugschema\schema.go | 18,933 | 5E163EB06DCA7C6BB6BC1DFF94FE9DAF56532037385D8A61F53583261A6715E5 | validate-only schema |
| internal\analyze\analyze.go | 32,560 | 646F4D62CDB4D2F6D7D2F1BED41F2043E3A3E750AA13B50C75752E0F01A25F14 | validate-only Graph JSON writer |
| internal\contracts\web_ui.go | 67,305 | CE4D76C678FFA99C43EAEEC053E530DF5FB0BFEE8801CC5684D18E213233A8BA | preserve existing contract |
| internal\mcp\context.go | 21,530 | 6358048275386AE5AB0B97565FB0A4E9D78462D27146D191FA39D3211428804A | affected transparent reader; validate |
| internal\mcp\impact.go | 36,860 | 5AF783341D3CD2D37D78CABBBA679E93F62F67E0DA42FEB63D75F34070CB6377 | affected transparent reader; validate |
| internal\filecontext\context.go | 51,554 | 9AA49A4E0784CCD84857E472F35CF3FA80385C22DC0F32277C944BFC5668B25E | preserve-only non-Evidence reader |
| internal\httpapi\graph.go | 7,627 | 2B9831B4EAC8CF14D8082394541940B69F5CC1BE4E2CFFCF5904F0223D9D142C | transparent transport; validate if exercised |
| anvien-web\src\services\file-detail-adapter.ts | 11,792 | 2D4E622637D8EA3E4E7547D29025D222ED57D3B2B95C6B5D257D3F65BB783570 | preserve-only |
| anvien-web\src\components\FileDetailPanel.tsx | 26,774 | BCA99FD0D0EF54752FB5BCBF5C78F63C9BEE218E4D8BBEB66DA98766A36702F7 | preserve-only |

## Fresh file-detail evidence

All rows came from the one fresh P5-D graph and reported HIGH file risk. HIGH is a blast-radius warning, not an edit prohibition.

| File | Symbols | In / Out / Local | Related | Unresolved | Flows / Tests | Decision |
|---|---:|---:|---:|---:|---:|---|
| resolution\resolve.go | 120 | 94 / 168 / 57 | 57 | 206 | 22 / 33 | bounded editable |
| resolution\emit.go | 148 | 43 / 223 / 76 | 47 | 388 | 13 / 17 | bounded editable |
| graph\types.go | 115 | 2,714 / 10 / 115 | 253 | 88 | 24 / 90 | preserve-only |
| resolution\indexes.go | 328 | 280 / 97 / 374 | 55 | 445 | 14 / 30 | bounded editable |
| resolution\export_resolution.go | 186 | 28 / 102 / 190 | 31 | 267 | 18 / 18 | accepted P5-C preserve-only |
| lbugload\csv.go | 193 | 56 / 53 / 92 | 22 | 201 | 2 / 9 | validate-only |
| lbugload\queries.go | 56 | 15 / 5 / 52 | 10 | 39 | 0 / 4 | validate-only |
| lbugschema\schema.go | 41 | 53 / 8 / 29 | 22 | 50 | 1 / 7 | validate-only |
| analyze\analyze.go | 326 | 87 / 310 / 172 | 186 | 509 | 20 / 7 | validate-only |
| contracts\web_ui.go | 193 | 41 / 59 / 160 | 31 | 636 | 8 / 2 | preserve-only |
| mcp\context.go | 165 | 48 / 111 / 50 | 28 | 261 | 21 / 2 | affected transparent reader |
| mcp\impact.go | 286 | 47 / 129 / 115 | 41 | 497 | 32 / 7 | affected transparent reader |
| filecontext\context.go | 556 | 345 / 89 / 565 | 46 | 542 | 51 / 9 | preserve-only |
| httpapi\graph.go | 77 | 20 / 58 / 15 | 22 | 119 | 16 / 3 | transparent transport |
| file-detail-adapter.ts | 118 | 22 / 4 / 40 | 6 | 141 | 0 / 2 | preserve-only |
| FileDetailPanel.tsx | 164 | 4 / 17 / 31 | 5 | 303 | 0 / 1 | preserve-only |

## Complete fresh upstream impact tuples

Every row is upstream impact from the same fresh graph. Every result had active resolution gaps / degraded nodes = 0 / 0. Counts are impacted symbols / affected files / modules / processes.

| Exact owner | Risk | Complete tuple | Meaning and representative affected-file trace |
|---|---|---|---|
| resolveCall | CRITICAL | 27 / 11 / 6 / 32 | CALLS emission seam. Files include analyze.go, analyze_test.go, legacy_resolver_conversion_test.go, pipeline_parity_test.go, cli\command.go, graphaccuracy\access_candidate.go, definition_collision_test.go, legacy_p7_conversion_test.go, p3c_binding_occurrence_test.go, resolution_test.go, resolve.go. |
| resolveAccess | CRITICAL | 27 / 11 / 6 / 32 | ACCESSES emission seam; same complete 11-file set as resolveCall. |
| emitter.emitReference | CRITICAL | 10 / 5 / 2 / 34 | Existing transparent Reference-to-Relationship projection. Files: analyze.go, emit.go, legacy_p7_conversion_test.go, resolution_test.go, resolve.go. Preserve implementation; use as validation boundary. |
| mergeRelationship | CRITICAL | 11 / 2 / 1 / 25 | Edge coalescing and Evidence replacement. Files: emit.go and resolve.go. Bounded union change is required only for proof conservation. |
| workspace.resolveImports | CRITICAL | 34 / 17 / 3 / 17 | Import orchestration/proof retention. Files include analyze.go, graphaccuracy\access_candidate.go, access_audit.go/test, export_resolution_test.go, export_tables_test.go, identity_occurrence_test.go, indexes.go, legacy heritage/import/P7/scope/suffix tests, p2b_persistence_test.go, resolution_test.go, resolve.go. |
| workspace.resolveImportedDef | HIGH | 20 / 13 / 1 / 3 | P5-C result is reduced to definition. Resolution package only; representative tests include export_resolution_test.go, export_tables_test.go, identity_occurrence_test.go, p2b_persistence_test.go, resolution_test.go and legacy resolver suites. |
| workspace.resolveImportedMember | CRITICAL | 20 / 9 / 3 / 34 | Namespace/member result reduction. Representative files: indexes.go, resolve.go, analyze.go, graphaccuracy\access_candidate.go, export_resolution_test.go, resolution_test.go and legacy resolver suites. |
| resolvedImport | CRITICAL | 115 / 20 / 2 / 62 | Shared import-state carrier. Complete affected family includes analyze.go; resolution access_audit, emit, export_resolution and test, export_tables and test, identity, import_resolution, indexes, legacy heritage/import/P7/scope/suffix tests, p2b persistence, resolution_test, resolve, and types. |
| graph.Evidence | CRITICAL | 494 / 124 / 24 / 286 | Shared machine contract. This blast radius is the reason to preserve Kind/Weight/Note and encode P5-D proof through the existing shape rather than add fields. |
| relationshipCSVRow | CRITICAL | 24 / 10 / 2 / 12 | Ladybug CSV relationship projection; validate current Evidence JSON transport. |
| relationshipEvidence | LOW | 24 / 10 / 2 / 0 | JSON marshal helper; no reader process directly depends on the helper node. Validate-only. |
| FallbackRelationshipInsertQuery | LOW | 6 / 3 / 2 / 0 | Fallback database insert already carries evidence. Validate-only. |
| relationColumns | LOW | 7 / 4 / 2 / 0 | Existing evidence STRING schema column. Preserve. |
| writeGraphSnapshotJSON | CRITICAL | 22 / 7 / 5 / 21 | Generic graph snapshot persistence. Validate-only. |
| contextRefPayload | CRITICAL | 8 / 3 / 1 / 13 | MCP context reader payload path. It carries Relationship; validate proof visibility. |
| impactItemPayload | CRITICAL | 11 / 5 / 1 / 17 | MCP impact reader payload path. It carries Relationship; validate proof visibility. |
| filecontext.relationshipSample | CRITICAL | 7 / 6 / 4 / 19 | File-detail sampling does not expose Evidence. Preserve-only absent a separate display decision. |
| WebUIContractTypeScript | CRITICAL | 2 / 2 / 2 / 8 | Generated contract already models Evidence Kind/Weight/Note. Preserve if shape stays unchanged. |
| httpapi.graphPayload | CRITICAL | 2 / 2 / 2 / 9 | Transparent graph HTTP transport; no semantic adapter discovered. |

Blast-radius warning:

- All proposed production owners are HIGH or CRITICAL surfaces. Edits must remain at the named seams and carry focused regression coverage.
- graph.Evidence is especially broad at 494 symbols / 124 files / 24 modules / 286 processes. A typed-field or schema-column expansion would materially widen P5-D and requires a new plan/impact gate. It is not authorized by this inventory.

## Affected-reader matrix

| Surface | Reads terminal endpoint | Reads full Evidence | Current state | Proposed P5-D action |
|---|---|---|---|---|
| Graph JSON | yes | yes, generic JSON | already transparent | byte-stable production; focused parity test |
| Ladybug CSV/schema/COPY/fallback | yes | yes, evidence STRING carrying JSON array | already transparent | byte-stable production; persistence round-trip test |
| MCP context | yes | yes via Relationship | affected transparent reader | real-reader validation |
| MCP impact | yes | yes via Relationship | affected transparent reader | real-reader validation |
| HTTP graph API | yes | yes via graph payload | transparent | validate only if selected by Main as nearest public boundary |
| Generated Web TypeScript contract | yes | existing kind/weight/note only | compatible with existing shape | preserve byte-identical |
| Web graph consumer | yes | transport only | no semantic export-hop interpretation found | no edit |
| file-detail / filecontext | yes for relationship samples | no Evidence field | not an Evidence reader | preserve; not in acceptance denominator |
| FileDetailPanel | renders file-detail data | no Evidence field | not an Evidence reader | preserve; UI/Playwright N/A |
| source-site-accuracy / resolution-health | proofKind/status, not Evidence | no | unaffected | preserve |
| Server.runCypherRead | generic relationship properties | transparent | usable inspection boundary | validate, no edit |
| Child 02 C09-C16 | Definition identity/selection/embedding semantics | no P5-D export-proof interpretation | unaffected accepted readers | preserve; do not import the 8-row denominator |

Reader denominator decision:

- Required P5-D affected readers are Graph JSON, Ladybug relationship evidence, MCP context, and MCP impact because they can transport the changed CALLS/ACCESSES Evidence.
- HTTP/Web transport is validate-only and production-preserve. UI is N/A because no current UI consumer interprets Relationship.Evidence.
- Child 02 C09-C16 remain accepted and preserve-only; they are not silently counted as P5-D readers.

## Minimum production owner manifest proposed for planner refresh

1. New internal\resolution\export_binding_proof.go
   - Sole P5-D adapter from immutable accepted P5-C exportResolutionResult/proofs to deterministic graph.Evidence records.
   - Own canonical ordering, dedupe, terminal/failure/hop projection, and no-proof behavior.
   - Consume accepted P5-C structures; do not alter P5-C traversal, P5-B tables, physical path resolution, or syntactic IMPORTS.

2. Bounded internal\resolution\indexes.go
   - Extend resolvedImport only enough to retain the immutable semantic result/proof produced during the existing phase-three lookup.
   - Retain semantic member proof at the accepted resolveImportedMember seam without adding a second traversal or path-resolution pass.
   - Preserve the existing one-sequence invariant: resolve file candidates once, build export tables once, resolve semantic terminals/bind once.

3. Bounded internal\resolution\resolve.go
   - At resolveCall and resolveAccess only, append the retained export binding proof to the existing generic Evidence while preserving endpoint, confidence, proofKind, target role, source-site ID, explicit-import no-global-rescue, and unaffected-language behavior.
   - Generic resolveName, resolveGlobalName, resolveGlobalCallName, path helpers, and unrelated reference families remain unchanged.

4. Bounded internal\resolution\emit.go
   - At mergeRelationship only, replace incoming-wins Evidence loss with deterministic stable union/dedupe so all coalesced source-site proof paths survive.
   - emitter.emitReference remains byte-stable because it already copies Reference.Evidence correctly.

Recommended encoding decision for Main to lock:

- Preserve graph.Evidence exactly as Kind, Weight, Note.
- Preserve current baseline evidence entry.
- Append ordered P5-D entries with versioned stable kinds such as export-terminal, export-hop, and export-failure.
- Encode each Note as canonical deterministic content with no map-order dependence; include source file, requested/exported/target names, target file, fact kind/provenance, terminal identity/outcome, and owner branch where applicable.
- Do not overload Relationship.ProofKind; it describes the call/access resolution mode, not the complete export traversal.
- Same logical evidence records dedupe; distinct owner branches and distinct source-site proof paths must survive coalescing.

If Main instead selects typed Evidence fields, new relationship columns, or a UI-specific decoded model, this four-file production manifest is invalid. Required action would be BLOCKED_FOR_PLAN_REFRESH with fresh impact for graph types, contract generation, Ladybug schema/load, and newly affected readers.

## Proposed focused test and report manifest after authorization

Production code must be implemented before tests.

| Proposed path | Purpose |
|---|---|
| new internal\resolution\export_binding_proof.go | deterministic proof adapter owner |
| internal\resolution\indexes.go | retain accepted semantic result/proof only |
| internal\resolution\resolve.go | CALLS/ACCESSES proof attachment only |
| internal\resolution\emit.go | deterministic coalesced Evidence union only |
| new internal\resolution\export_binding_proof_test.go | direct/barrel equality, alias/star/namespace/member proof, cycle/ambiguity/failure ownership, deterministic ordering/dedupe |
| bounded internal\resolution\resolution_test.go | emitted terminal CALLS/ACCESSES, two-site coalescing, no false physical/global fallback, unchanged unrelated languages |
| new internal\lbugload\p5d_export_proof_persistence_test.go | exact Evidence Graph-to-CSV-to-Ladybug round-trip and endpoint conservation |
| bounded internal\analyze\analyze_test.go | Graph JSON proof survival at built analyzer boundary |
| one new reports\coder P5-D implementation handoff | source/test manifest, build, reader parity, fixed-corpus and target-ready evidence |

Accepted P5-C export_resolution.go and export_resolution_test.go should remain byte-stable. Reuse their fixture helpers only through new P5-D tests if possible; do not rewrite accepted semantics.

## Required validation sequence after production authorization

1. Re-anchor HEAD/worktree and verify the authorized owner manifest. Do not rerun this accepted inventory unless source/HEAD invalidates it.
2. Implement production behavior first in only the authorized owners.
3. Add the focused tests above after production behavior exists.
4. Run focused resolution proof/emission tests and the complete internal/resolution package on final source/test bytes.
5. Before build, use E-only anvien doctor lock/process evidence and terminate only confirmed build-related owners.
6. Run the literal canonical full build from cwd E:\Anvien:

       powershell -ExecutionPolicy Bypass -File E:\Anvien\scripts\full-build.ps1

7. On unchanged post-build source/test bytes, validate:
   - focused P5-D resolution proof matrix;
   - full internal/resolution regression;
   - Graph JSON terminal/proof survival;
   - Ladybug endpoint and Evidence round-trip parity;
   - MCP context and impact proof visibility;
   - unchanged P5-C matrix and unaffected-language regression;
   - zero physical target-file, resolver-emitted syntactic IMPORTS, and accepted persisted IMPORTS deltas against the fixed corpus.
8. Do not rerun build/analyze gates without concrete byte or graph invalidation.
9. Produce an immutable READY_FOR_SUPERVISOR Coder report and return to Main. Supervisor, detect, ledger update, commit, and target action remain Main-governed.

What each future boundary proves:

- Focused resolution tests: terminal endpoint plus complete deterministic proof composition.
- Full internal/resolution: no resolver-family regression.
- Canonical build: repository compiles and canonical E-only outputs are generated from final bytes.
- Graph JSON test/analyze: Evidence reaches normal graph persistence.
- Ladybug parity: endpoint and every proof record survive CSV/schema/load/readback with zero loss.
- MCP context/impact: actual affected readers expose the same terminal/proof.
- Fixed corpus: physical module resolution and source-written IMPORTS semantics remain unchanged.

## Exact target preflight/work-step authority for later Main opening

No command below was run in this inventory. E:\cheapapp.org was not accessed.

Only after Main explicitly opens the target work step:

1. Record target source/worktree status and hashes without modifying source.
2. Record the target pre-state for exactly the two bounded direct/barrel call sites and their false-gap rows.
3. Invoke the freshly built canonical E:\Anvien runtime from the target root for one authorized fresh analyze.
4. Query the two source sites by source-site identity/range and inspect terminal CALLS plus complete ordered export proof.
5. Prove direct/barrel terminal equality, resolved calls 2/2, matching false gaps 0, and complete proof chains 2/2.
6. Compare physical target-file resolution, resolver-emitted syntactic IMPORTS, and persisted graph-wide IMPORTS to the accepted baselines; required delta is 0 / 0 / 0.
7. Read the same relationships through Graph JSON, Ladybug, and only the affected readers named above.
8. Verify target source/worktree unchanged after analyze; all official reports remain under E:\Anvien.
9. UI/Playwright remains N/A unless new impact proves a semantic UI consumer. Do not manufacture UI scope.

## Preservation and out-of-scope manifest

Preserve-only:

- Child 04 ExportFact facts/providers and all upstream ScopeIR semantics.
- Accepted P5-B export_tables.go table shape, construction, and zero-physical semantics.
- Accepted P5-C export_resolution.go traversal/outcome/proof semantics and export_resolution_test.go matrix.
- resolveImportFiles, resolveImportFile, import_resolution.go, TargetFiles, and all non-TS/JS path strategies.
- Explicit-import no-global-rescue guard and generic global/name helpers.
- emitImportEdges and every syntactic IMPORTS behavior/count.
- graph.Evidence field shape unless a new planner/impact gate explicitly widens it.
- Graph JSON writer, Ladybug schema/load/query production, MCP/HTTP/Web production readers under the recommended existing-shape encoding.
- Child 02 C09-C16 accepted behavior.
- All four Child 05 ledgers and all existing Coder/Supervisor/Main reports in this inventory.

Out of scope:

- P5-C reopen or semantic traversal changes.
- P5-D production before Main authorization.
- Pn-A+, Child 06, target source changes, ambient/external resolution, unrelated relationship types/readers, UI redesign, target execution, detect, commit, and Supervisor.

## Git/worktree boundary

Pre-report reality:

- cwd/toplevel: E:\Anvien.
- branch: master.
- HEAD: fd6cb52f6258be2cbdaa622ee53c2d31d173566d.
- tracked worktree diff: none.
- staging/index diff: none.
- protected untracked Main rotation handoffs: 11.

The opening delegation described 10 handoffs. Current Git reality contains one additional later rotation handoff, reports\Investigation\rp_main_260821_231548_orchestration_rotation_handoff.md. It is treated as protected Main-owned state; no authority transfer is inferred from its presence.

| Protected untracked handoff | Bytes | SHA-256 |
|---|---:|---|
| rp_main_260821_0631_orchestration_rotation_handoff.md | 6,542 | 623FDC57BAC97F4C1F86F6A39C463E11F6BC0FFDA7DB8E9E661F0B0C1FFCC9EB |
| rp_main_260821_0721_orchestration_rotation_handoff.md | 6,542 | FDDAFFA421D64B10B2BBF6DDA11B8705E1E478BB358A4EF0EE3C497F3A5F019B |
| rp_main_260821_1518_orchestration_rotation_handoff.md | 12,050 | 921C8528C6A3A20A584AECFD479CF6540CBACCC91DD8FFAF7643676F9C8173F1 |
| rp_main_260821_155017_orchestration_rotation_handoff.md | 7,691 | E4EFBB0C81C18E4991FD3C3080496B53CD31DE022CADE4E97CB5B8DE6CE0AD1B |
| rp_main_260821_163855_orchestration_rotation_handoff.md | 12,088 | 38864FE124A331E48F2AFBA363CD7C1422CB5E1E518EFA5DF6E271502B5D713B |
| rp_main_260821_172833_orchestration_rotation_handoff.md | 16,793 | F998CF97FF09C39CF11F402A707372F7F77D5369C2A6FC782FB40ADD9DC5270C |
| rp_main_260821_195827_orchestration_rotation_handoff.md | 12,395 | 9AC992ED4FDF0B800CC4ED7A3D53D1D82947657CC9ADB09B7E9AE93085BF0DA2 |
| rp_main_260821_204245_orchestration_rotation_handoff.md | 15,695 | B5D5FF7DC1E77DD4A42B9989ACF5488CDF92E0404410127210492495EA1DD907 |
| rp_main_260821_213709_orchestration_rotation_handoff.md | 16,404 | 85DA4EBA574AFDEA36CFBE86134834E0212C94A5591D7261F6B647A293E4A6AC |
| rp_main_260821_222919_orchestration_rotation_handoff.md | 16,326 | 0A2EE9D8C60212C12A71C6E3B3FF41EA96F41A9BE7D119CDD67D8F15EA8E857E |
| rp_main_260821_231548_orchestration_rotation_handoff.md | 15,868 | 667E5984731DA84C11D411173BE2FE9FF17C69ACD0A788FFE2DF638055755716 |

Post-report expected boundary is those same 11 protected handoffs plus this one new Coder report, with staging still empty and no tracked diff.

## Handoff

- Inventory verdict: production proof projection is required; target acceptance remains pending and unclaimed.
- Exact next action: Main 01a024ef-2740-7802-a9cb-71eb5d951496 reads this immutable report, planner-refreshes the existing-Evidence encoding and exact four-file production owner manifest, commits that docs authorization if accepted, then resumes this same Coder lane for production. If OFFICIAL AUTHORITY TRANSFER occurs first, the same report identity/verdict routes to successor 01a02518-d4a9-7081-860b-b2e5dddde93e.
- Artifact path: E:\Anvien\reports\coder\rp_coder_260821_232052_by_gpt-5_child05_p5d_pre_implementation_inventory.md.
- Final bytes/LF/CR/BOM/SHA-256 are measured from the frozen file and supplied in the immutable handoff message. They cannot be embedded as a self-hash without changing the bytes being identified; this file must not be edited after that handoff.

PRE_IMPLEMENTATION_READY / READY_FOR_MAIN_AUTHORIZATION
