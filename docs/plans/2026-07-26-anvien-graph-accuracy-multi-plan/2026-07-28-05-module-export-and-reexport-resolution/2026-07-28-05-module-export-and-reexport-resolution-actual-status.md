# Anvien Module Export And Re-Export Resolution Actual Status

Title: Anvien Module Export And Re-Export Resolution
Date: 2026-07-28
Status: P0 Complete / P5-A committed and accepted / P5-B committed at c1559df9 / P5-C Supervisor PASS and detect recorded, commit pending / P5-D+ locked
Companion plan: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-05-module-export-and-reexport-resolution/2026-07-28-05-module-export-and-reexport-resolution-plan.md`
Companion evidence: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-05-module-export-and-reexport-resolution/2026-07-28-05-module-export-and-reexport-resolution-evidence.md`
Companion benchmark: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-05-module-export-and-reexport-resolution/2026-07-28-05-module-export-and-reexport-resolution-benchmark.md`
Roadmap: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-26-anvien-graph-accuracy-roadmap.md`
Predecessor: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-04-typescript-export-semantics/2026-07-28-04-typescript-export-semantics-plan.md`
Successor: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-plan.md`

## Purpose

This file records the current module-request, path lookup, export-surface, re-export traversal, terminal call-binding, and target-boundary state before Child 05 production work.

The Child 04 predecessor is accepted at closure commit `0aa49c87628c9e8b2041754515d6ebf0a930d55b`. P5-A fresh inventory is recorded at the same HEAD and this R2 refresh authorizes only the exact requested-meaning/type-only implementation defined below. Detailed proof belongs in the evidence ledger.

## Freshness / Refresh Rules

Refresh this file:

- after the Child 04 handoff changes the available import/export facts;
- before each P5 slice after fresh `anvien analyze --force`, file-detail, and impact;
- after each accepted slice and before opening the next one;
- whenever an affected persistence/reader or owner file differs from the current touch map.

Update affected rows with explicit transitions and append a refresh-log row. Do not delete earlier accepted refreshes.

## Scope

Target scope:

- repository-backed TypeScript module request/path inputs;
- export surfaces derived from accepted Child 04 facts;
- alias/re-export/star/cycle/ambiguity/meaning traversal;
- terminal Symbol binding and proof for the two bounded barrel calls;
- syntactic `IMPORTS` and physical path-resolution preservation;
- only persistence/readers proven affected by Child 02 inventory and fresh impact.

Out of scope:

- Child 04 export extraction;
- Child 06 ambient/external declaration authority;
- broad package-resolution or reader redesign without evidence;
- graph-output behavior, scanner behavior, target-source edits, and unrelated language semantics.

## Relationship / Impact Evidence

The 2026-08-10 P0 tuples remain historical. The P5-A graph was freshly rebuilt at HEAD `0aa49c87628c9e8b2041754515d6ebf0a930d55b`; every listed file-detail result reported `stale=false` and `changedSinceAnalyze=false`.

| Unit / File / Surface | File Detail Evidence | Related File Count | Relationship Summary | Impact Note |
|-----------------------|----------------------|-------------------:|----------------------|-------------|
| `internal/scopeir/facts.go` | `E5-P5A-IMPACT1` | 247 | shared `ImportFact` contract | medium file risk; file impact CRITICAL |
| `internal/providers/tsjs/imports.go` | `E5-P5A-IMPACT1` | 17 | source-written TS/JS imports plus compatibility re-export path facts | high file risk; exact provider methods LOW |
| `internal/scopeir/ir.go` | `E5-P5A-IMPACT1` | 245 | owning clone and canonical normalization | high file risk; file impact CRITICAL |
| `internal/scopeir/sort_keys.go` | `E5-P5A-IMPACT1` | 242 | deterministic `ImportFact` ordering | high file risk; file impact CRITICAL |
| `internal/resolution/indexes.go` | `E5-P5A-IMPACT1`, `E5-P5B-IMPACT1` | 51 | live module/file result, optional physical-definition result, P5-B storage/wiring owner | high file risk; preserve path behavior; minimal P5-B wiring only |
| `internal/resolution/import_resolution.go` | `E5-P5A-IMPACT1` | 34 | existing multi-language path strategies | high file risk; preserve in P5-A |
| `internal/resolution/export_tables.go` | `E5-P5B-IMPACT1` | new file; no pre-edit graph node | dedicated syntax-derived export-table owner selected by fresh source/graph inventory | implementation owner; create only after Main authorization |

| Symbol | Impact Evidence | Risk | Impacted Symbols | Affected Files | Modules | Processes | Linked Flows / Tests |
|--------|-----------------|------|-----------------:|---------------:|--------:|----------:|----------------------:|
| `scopeir.ImportFact` | `E5-P5A-IMPACT1` | CRITICAL | 624 | 73 | 25 | 67 | 0 / 100 from containing file |
| `collector.emitImportStatement` | `E5-P5A-IMPACT1` | LOW | 0 | 0 | 0 | 0 | 1 / 4 from containing file |
| `collector.addImport` | `E5-P5A-IMPACT1` | LOW | 0 | 0 | 0 | 0 | 1 / 4 from containing file |
| `collector.addSourceExportFact` | `E5-P5A-IMPACT1` | LOW | 0 | 0 | 0 | 0 | 1 / 4 from containing file |
| `ScopeIR.Normalized` | `E5-P5A-IMPACT1` | CRITICAL | 21 | 6 | 4 | 15 | 3 / 98 from containing file |
| `ScopeIR.NormalizeInPlace` | `E5-P5A-IMPACT1` | LOW | 5 | 2 | 1 | 0 | 3 / 98 from containing file |
| `ScopeIR.NormalizeOwned` | `E5-P5A-IMPACT1` | LOW | 3 | 1 | 1 | 0 | 3 / 98 from containing file |
| `compareImport` | `E5-P5A-IMPACT1` | LOW | 6 | 2 | 1 | 2 | 2 / 97 from containing file |
| `resolvedImport` | `E5-P5A-IMPACT1` | CRITICAL | 95 | 16 | 2 | 43 | 13 / 28 from containing file |
| `buildWorkspace` | `E5-P5A-IMPACT1` | CRITICAL | 48 | 21 | 7 | 25 | 13 / 28 from containing file |
| `workspace.resolveImports` | `E5-P5A-IMPACT1` | CRITICAL | 27 | 15 | 3 | 19 | 13 / 28 from containing file |
| `workspace.resolveImportFiles` | `E5-P5A-IMPACT1` | CRITICAL | 18 | 11 | 1 | 5 | 13 / 28 from containing file |
| `workspace.resolveImportFile` | `E5-P5A-IMPACT1` | LOW | 3 | 1 | 1 | 1 | 13 / 28 from containing file |
| `workspace.resolveImportedDef` | `E5-P5A-IMPACT1` | CRITICAL | 18 | 11 | 1 | 5 | 13 / 28 from containing file |

## Status Rules

| Status | Meaning | Allowed next action |
|--------|---------|---------------------|
| `correct` | Current behavior satisfies this child's bounded requirement. | Preserve and regress. |
| `partial` | A usable part exists, but its contract or coverage is incomplete. | Change only the proved gap. |
| `wrong` | Current behavior conflicts with the accepted requirement. | Replace at the proved owner. |
| `missing` | Required behavior does not exist. | Add it at the evidence-selected owner. |
| `unbound` | The fact/result exists but is not connected to the terminal consumer. | Bind only that boundary. |
| `blocked` | Required predecessor or impact evidence is unavailable. | Do not edit until refreshed. |

## Current Status Matrix

| Unit | Current State | Required State | Status | Relationship Count | Evidence | Next Plan Decision |
|------|---------------|----------------|--------|-------------------:|----------|--------------------|
| Problem authority | Originating report records the barrel symptom; causal and Supervisor reports verify the bounded C6 cause without prescribing a fix | findings and target acceptance separated from proposed design | correct | N/A | `E0-P0A-ORIGIN1`, `E0-P0A-VERIFY1`, `E0-P0A-VERIFY2` | preserve this evidence hierarchy |
| Child 04 fact handoff | accepted predecessor provides one immutable `ScopeIR.ExportFact` source of truth with seven source-form kinds, three meaning lanes, type-only/name/range/provenance/diagnostic state, `414/414` graph conservation, `11,592/0` persistence parity, and no terminal traversal/ambiguity/cycle/resolved/public-API state | consume accepted syntax/direct-export facts without regenerating them; compatibility `ImportFact` records remain path-only | correct | accepted `414/414` Export/File→Export boundary | `E4-PNC-CLOSE1`, `E4-PNC-COMMIT1`, `E4-P4C2-REVIEW2`, `E5-P5A-INPUT1` | preserve as the sole re-export semantic authority |
| Module request and file result | TS/JS facts retain source file, raw module text, local/imported names, deterministic relative/index candidates, and a live `resolvedImport` module/file result distinct from optional `TargetDef`; other-language strategies are separate | preserve exact module/file lookup and separation from export lookup | correct | `51` workspace / `34` path-owner related files | `E5-P5A-INPUT1`, `E5-P5A-IMPACT1`, `E5-P5A-COUNT1` | preserve `indexes.go` and `import_resolution.go` in P5-A |
| Requested exported name | default/named/alias names remain exact; namespace imports now retain only `LocalName`, kind, and raw module with an empty `ImportedName` | exact default/named/alias name; namespace has no exported-name request and is represented by kind + raw module | correct | `17` provider related files | `E5-P5A-INPUT1`, `E5-P5A-SRC1`, `E5-P5A-TEST1`, `E5-P5A-REVIEW1`, `E5-P5A-COMMIT1` | preserve accepted P5-A commit `2560f914` |
| Requested meaning / type-only | source-written TS/JS imports now carry canonical requested meanings and explicit type-only state; compatibility re-export and non-TS/JS facts remain empty | `RequestedMeanings` canonical allowed-set plus explicit `TypeOnly`; normal default/named/alias `{value,type,namespace}`, type-only `{type}`, namespace `{namespace}` or `{type}` when type-only | correct | `247` contract / `17` provider / `245` normalization / `242` sort-key related files | `E5-P5A-INPUT1`, `E5-P5A-SRC1`, `E5-P5A-TEST1`, `E5-P5A-REVIEW1`, `E5-P5A-COMMIT1` | preserve accepted P5-A commit `2560f914` |
| Dormant `ImportFact.Target*` versus live result | no production writer assigns `TargetFile`, `TargetExportedName`, `TargetModuleScope`, `TargetDefID`, or `LinkStatus`; live resolution is the separate in-memory `resolvedImport` | do not activate/remove dormant fields in P5-A; preserve live result separation | correct | `247` contract / `51` workspace related files | `E5-P5A-INPUT1`, `E5-P5A-IMPACT1` | preserve; later cleanup requires its own accepted evidence |
| P5-A count authority | fixed-corpus validation on the same `736` parsed-code corpus records `5,072` physical target-file resolutions, `5,072` resolver-emitted syntactic `IMPORTS`, and `5,088` final persisted graph-wide `IMPORTS` | retain all three denominators and prove delta `0` for each | correct | accepted P5-A commit plus P5-B fixed-corpus preservation run | `E5-P5A-COUNT1`, `E5-P5B-COUNT1` | preserve `0 / 0 / 0` |
| Module export surface | dedicated `internal/resolution/export_tables.go` now builds deterministic explicit entries and star adjacency from accepted facts, with only minimal `workspace/buildWorkspace` wiring | deterministic table derived only from Child 04 facts; no physical-definition inference or terminal traversal | correct | post-build new owner: 30 related files / 62 symbols / HIGH risk; workspace/buildWorkspace CRITICAL; detect graph saw `indexes.go` high risk and one expected pre-commit analyzer gap for `w.buildExportTables`; isolated commit is now present | `E5-P5B-IMPACT1`, `E5-P5B-SRC1`, `E5-P5B-ZEROBARREL1`, `E5-P5B-REVIEW1`, `E5-P5B-DETECT1`, `E5-P5B-COMMIT1` | P5-B closed; refresh P5-C traversal impact before code; keep P5-D/target locked |
| Re-export traversal | final uncommitted P5-C candidate uses dedicated proof-bearing traversal and one file-candidates -> tables-once -> terminal-bind sequence; after the source-backed REJECT it now composes ambiguous owners independently, retains an owned missing-member branch, preserves aggregate ambiguity, and selects no sole surviving member | terminal traversal with alias/star/cycle/ambiguity/meaning proof, including complete per-owner member provenance | correct | accepted impact remains `resolveImportedDef` HIGH `19/12/1/3`; `resolveImports` CRITICAL `28/16/3/17`; `resolveImportedMember` CRITICAL `18/8/3/34`; `buildWorkspace` CRITICAL `49/22/8/23`; fresh detect records high changed-file risk with no affected process | `E5-P5C-IMPACT1`, `E5-P5C-SRC1`, `E5-P5C-PROOF1`, `E5-P5C-REPAIR1`, `E5-P5C-TEST1`, `E5-P5C-REVIEW1`, `E5-P5C-DETECT1` | stage, verify, and commit only the isolated P5-C manifest before opening P5-D |
| Explicit-import global-name-rescue boundary | final uncommitted candidate gates repository-global fallback only at `resolveCall`; generic global helpers remain unchanged; focused replay and final Supervisor review record zero false global calls | no repository-global same-name rescue; explicit export failure retained | correct | `resolveCall` CRITICAL `27/11/7/32`; `resolveGlobalCallName` CRITICAL `6/4/2/23` and preserve-only; fresh detect records no affected process | `E5-P5C-IMPACT1`, `E5-P5C-NOGLOBAL1`, `E5-P5C-REVIEW1`, `E5-P5C-DETECT1` | preserve the accepted guard; isolated P5-C commit remains pending and P5-D remains locked |
| Terminal call/proof emission | the two accepted target calls lack terminal `CALLS` relationships | both sites bind to the expected terminal Symbols with complete proof | unbound | resolver/emission impact must be refreshed in P5-D | `E0-P0A-VERIFY1` | P5-D after P5-C |
| Affected persistence/readers | no current Child 05 evidence identifies which corrected proof fields require adapters | only consumers named by Child 02 inventory and fresh P5-D impact are editable | blocked | unknown until Child 02/P5-D evidence | `E0-P0A-SCOPE1` | do not assume a fixed reader denominator |
| Target boundary | accepted target is analyzed in place and its source is not a fixture or edit surface | preserve source/worktree; regenerate only normal target index during P5-D validation | correct | accepted target graph: 114,125 relationships | `E0-P0A-BOUNDARY1` | preserve until P5-D pre/post evidence |

## Status Refresh Log

| Refresh | Date | Repo Basis | Changed Scope | Status Changes | Evidence | Next Phase Update |
|---------|------|------------|---------------|----------------|----------|-------------------|
| R0 | 2026-08-10 | HEAD `238aec06d28286acc0fca0f4e6a69f9eb4ff6a49`; graph file-detail `stale=false`, analyzed `2026-08-09T19:19:54Z` | Child 05 source/report/ledger reset | removed campaign-wide assumptions; path behavior classified separately from export lookup; P5-A blocked only by Child 04 handoff | `E0-P0A-GRAPH1`, `E0-P0A-SRC1..E0-P0A-SRC4`, `E0-P0A-FD1..E0-P0A-FD4`, `E0-P0A-IMPACT1..E0-P0A-IMPACT4` | open P5-A only after predecessor refresh; inventory before code |
| R1 | 2026-08-21 | Child 04 closure commit `0aa49c87628c9e8b2041754515d6ebf0a930d55b`; Pn-B commit `d1d8eb9002ce9c449c3713de0837ac8216d17a8d` | accepted Child 04 syntax/direct-export predecessor and P5-A opening | predecessor `blocked -> accepted`; P0 remains complete; 2026-08-10 owner/file-detail/impact tuples remain historical and cannot authorize production edits | `E4-PNC-CLOSE1`, `E4-PNC-COMMIT1` | open only P5-A for fresh source/file-detail/impact/input/count inventory; update status/next action/work steps before implementation; keep P5-B+ and target locked |
| R2 | 2026-08-21 | HEAD `0aa49c87628c9e8b2041754515d6ebf0a930d55b`; fresh graph `114,738` nodes / `157,553` relationships; Coder inventory report SHA-256 `82D9F651A0BF6CF13CD66F0EEF6DC310F9DAA69A4782E3769383A3294F8672DE` | P5-A input contract, exact owners, count denominators, and implementation work steps | module/path result stays `correct`; requested names split to `correct/partial`; requested meaning/type-only `missing`; inventory block resolved by one planner refresh | `E5-P5A-IMPACT1`, `E5-P5A-INPUT1`, `E5-P5A-COUNT1` | authorize only P5-A requested-meaning/type-only implementation; P5-B+ and target remain locked |
| R3 | 2026-08-21 | HEAD `0aa49c87628c9e8b2041754515d6ebf0a930d55b` plus uncommitted P5-A candidate; canonical CLI `1.2.8`; authoritative `E:\Anvien` full-build/analyze completed; graph `114,842 / 157,744` | requested name/meaning/type-only contract, canonical ownership, build/boundary/regression, and three-count preservation | requested namespace name `partial -> correct`; requested meaning/type-only `missing -> correct`; module/path result and all three count denominators remain `correct` pending Supervisor recheck on the canonical E graph | `E5-P5A-SRC1`, `E5-P5A-BUILD1`, `E5-P5A-TEST1`, `E5-P5A-COUNT1` | hand P5-A candidate to Supervisor; keep P5-B+, target, detect, and commit locked/pending |
| R4 | 2026-08-21 | HEAD `0aa49c87628c9e8b2041754515d6ebf0a930d55b` plus uncommitted P5-A candidate; canonical CLI `1.2.8`; authoritative `E:\Anvien` full-build/analyze completed; graph `114,842 / 157,744` | Supervisor acceptance of P5-A source, build, boundary, regressions, and three-count preservation | P5-A Supervisor gate `pending -> PASS`; source invariant remains correct; detect/commit remain pending | `E5-P5A-BUILD1`, `E5-P5A-COUNT1`, `E5-P5A-REVIEW1` | run fresh analyze refresh, then `detect-changes`; stage/commit only isolated P5-A manifest; keep P5-B+, target, and later gates locked |
| R5 | 2026-08-21 | HEAD `40ea0095a79084a3c6805cf5d5f46108926d1dca`; fresh E graph `114,852 / 157,754`; P5-B inventory report SHA-256 `CF01A36C2AAAF7B85574EFB8F69BE931939AE4F4704AE5D5E0CB2F30835A4056` | P5-A acceptance/commit is now the predecessor; P5-B exact owner and CRITICAL blast radius are established | P5-B export surface remains `missing`; owner selected as dedicated `export_tables.go` with minimal `workspace/buildWorkspace` wiring; no source edit yet | `E5-P5B-IMPACT1` | authorize only the bounded P5-B production tranche; keep P5-C/P5-D/target locked |
| R6 | 2026-08-21 | HEAD `17f3dad2587f2785d02ad266e188c7a8feca499b` plus uncommitted P5-B candidate; exact absolute E build PASS; full graph `115,099 / 158,118`; Supervisor resubmission PASS report SHA-256 `749AEB5699CC191C16CFC5EED1561B2717937BC0D6F94F9C9E5EB9A4E940988A` | dedicated export table, minimal workspace wiring, focused tests, fixed-corpus preservation, and two evidence-blocker repairs | module export surface `missing -> correct`; P5-B Supervisor gate `pending -> PASS`; detect/commit pending | `E5-P5B-SRC1`, `E5-P5B-ZEROBARREL1`, `E5-P5B-BUILD1`, `E5-P5B-TEST1`, `E5-P5B-COUNT1`, `E5-P5B-REVIEW1` | refresh graph, run detect-changes, commit isolated P5-B manifest; only then open P5-C |
| R7 | 2026-08-21 | HEAD `17f3dad2587f2785d02ad266e188c7a8feca499b` plus the same uncommitted P5-B candidate; one fresh `anvien analyze --force` exit `0`, graph `115,134 / 158,153`; `anvien detect-changes --repo E:\Anvien --scope all` exit `0` | Main-owned pre-commit impact transition for the accepted P5-B candidate | detect gate `pending -> recorded`; summary `18` changed symbols / `5` changed files / `1` affected process, `backend=4` and `docs=14`, risk `medium`, changed-file risk `high`; one expected analyzer gap for `w.buildExportTables`, persisted gap total `0` | `E5-P5B-DETECT1` | stage exactly the 11-path P5-B manifest, verify cached diff/protected handoffs, then create the isolated P5-B commit; P5-C remains locked |
| R8 | 2026-08-21 | P5-B commit `c1559df953a277b099009f8489576d00ed25aa58` on parent `17f3dad2587f2785d02ad266e188c7a8feca499b`; docs-only refresh follows the accepted 11-path commit | P5-B closure and P5-C opening transition | P5-B `open -> committed`; P5-C `locked -> open`; P5-D/target remain locked; no fresh P5-C graph evidence exists yet | `E5-P5B-COMMIT1` | refresh graph, file-detail, and upstream impact for the P5-C traversal owner before any production edit; then hand off the existing E-only Coder lane for inventory |
| R9 | 2026-08-21 | HEAD `cd35b48f5466117fa1348fdc71c52e1408685a1b`; fresh graph `115,134 / 158,153`; immutable Coder inventory SHA-256 `7E77E5010CC3D728753547D0EE5499C51ED38CFD9660130381274BBB3DDCA256` | P5-C traversal/global-rescue owner and sequencing inventory | inventory `pending -> accepted`; exact owner set becomes dedicated `export_resolution.go` plus bounded `indexes.go` and `resolve.go`; two-phase `resolveImports` selected because tables require completed file candidates; all HIGH/CRITICAL tuples recorded | `E5-P5C-IMPACT1` | authorize only P5-C production after this docs commit; code first, tests after code; keep generic global/path/table/emission/P5-D/target surfaces preserve-only |
| R10 | 2026-08-21 | HEAD `861000cb6b6e36ce105623f0dc8c093b089f61fa` plus exact uncommitted four-file P5-C candidate; immutable Coder handoff SHA-256 `3AA88C36D8AF4E839BADB31B4388EAB183EEE80ED1D70E6F56311C5019B4D28D`; final candidate graph `115,800 / 159,445` | traversal/result/proof owner, two-phase orchestration, namespace/member consumer, no-global guard, final-byte build/tests, and fixed-corpus preservation | re-export traversal `wrong -> correct` in candidate; global-rescue boundary `wrong -> correct` in candidate; P5-C remains unchecked and Supervisor pending | `E5-P5C-SRC1`, `E5-P5C-PROOF1`, `E5-P5C-BUILD1`, `E5-P5C-TEST1`, `E5-P5C-NOGLOBAL1` | resume only existing E-only Supervisor; keep review/detect/commit, P5-D, and target locked |
| R11 | 2026-08-21 | same HEAD `861000cb6b6e36ce105623f0dc8c093b089f61fa` plus repaired four-file P5-C candidate; REJECT SHA-256 `1B8E32DA433782116A70BE427962A12DD799FF26B431BF7982A2D32FCE3121F1`; reject-only Coder resubmission SHA-256 `14A9ADB407527F12CF5F27484FF1E8C3E731AF174573D63D6895CDC607E4CC5F`; final Supervisor PASS SHA-256 `AC600E17A023E58261355C18647FC17674DCB1E2238258F9CE6941ABD49739DA` | distinct-owner ambiguity/member composition, per-owner missing provenance, fail-closed definition/member consumers, final source/test/build identities, and reject-only re-review | Supervisor gate `pending -> REJECT -> PASS`; `export_resolution.go` final `566A69B9...`, `export_resolution_test.go` final `97FF4990...`; cleared `indexes.go`/`resolve.go` remain byte-identical; P5-C stays unchecked until detect and isolated commit | `E5-P5C-REPAIR1`, `E5-P5C-BUILD1`, `E5-P5C-TEST1`, `E5-P5C-REVIEW1` | run exactly one fresh analyze and detect-changes; then stage/commit only the authorized P5-C manifest; keep P5-D and target locked |
| R12 | 2026-08-21 | same HEAD plus final candidate/ledger/report worktree; one fresh `anvien analyze --force` produced `115,947 / 159,626`; `anvien detect-changes --repo E:\Anvien --scope all` exited `0` | pre-commit graph refresh and full changed-symbol/file/gap/schema classification | detect `pending -> recorded`: `55` changed symbols / `6` changed files / `0` affected processes; summary risk `low`, changed-file risk `high`; `17` changed gap entities with persisted total `0`; semantic schema complete | `E5-P5C-DETECT1` | stage exactly the authorized 12-path P5-C manifest, verify cached diff and protected handoffs, then create the isolated commit; P5-D/target remain locked |

## Phase Touch Map

| Unit / File / Surface | Plan-Relevant Relationship File | Relationship to Target | Plan Item | Touch Mode | Evidence | Constraint |
|-----------------------|---------------------------------|------------------------|-----------|------------|----------|------------|
| source-written import contract | `internal/scopeir/facts.go` | owns `ImportFact` requested semantic inputs | P5-A | edit | `E5-P5A-INPUT1`, `E5-P5A-IMPACT1` | add only `RequestedMeanings` and `TypeOnly`; dormant `Target*` fields unchanged |
| TS/JS import syntax | `internal/providers/tsjs/imports.go` | populates default/named/alias/namespace and type-only requests; also emits compatibility re-export path facts | P5-A | edit source-written import path only | `E5-P5A-INPUT1`, `E5-P5A-IMPACT1` | accepted `ExportFact` remains sole re-export semantic source; compatibility facts keep requested fields empty |
| canonical clone/normalization | `internal/scopeir/ir.go` | owns nested requested-meaning collection and canonical set | P5-A | edit | `E5-P5A-INPUT1`, `E5-P5A-IMPACT1` | deep clone, sort, and deduplicate without changing export normalization |
| deterministic import ordering | `internal/scopeir/sort_keys.go` | compares import semantic inputs | P5-A | edit | `E5-P5A-INPUT1`, `E5-P5A-IMPACT1` | include requested meanings and `TypeOnly`; preserve existing ordering inputs |
| current workspace/import binding | `internal/resolution/indexes.go` | live module/file result, current physical-definition lookup, P5-B table storage, and P5-C orchestration/member consumers | P5-A/P5-B/P5-C | bounded P5-C edit at `resolveImports`, `resolveImportedDef`, `resolveImportedMember`, and redundant table call only | `E5-P5A-INPUT1`, `E5-P5B-IMPACT1`, `E5-P5C-IMPACT1` | keep path candidates exact; implement one two-phase pass and shared export lookup; no broad workspace refactor |
| dedicated export-table owner | `internal/resolution/export_tables.go` | owns deterministic syntax-derived explicit entries and star adjacency | P5-B | edit | `E5-P5B-IMPACT1` | derive only from accepted `ExportFact`; no physical-definition inference or terminal traversal |
| dedicated export-resolution owner | `internal/resolution/export_resolution.go` | owns deterministic terminal/ambiguity/cycle/missing/meaning outcomes and proof hops | P5-C | create/edit | `E5-P5C-IMPACT1` | consume P5-B tables only; no path, graph, persistence, or target behavior |
| current path strategies | `internal/resolution/import_resolution.go` | module/file resolution and other-language strategies | P5-A | preserve-only | `E5-P5A-INPUT1`, `E5-P5A-IMPACT1` | no broad package/path rewrite; unaffected languages must regress unchanged |
| terminal call resolution | `internal/resolution/resolve.go` | consumes scope/import bindings and contains global fallback | P5-C/P5-D | bounded P5-C edit at `resolveCall` only | `E5-P5C-IMPACT1` | explicit import failure guard only; `resolveGlobalCallName` and other generic helpers preserve-only |
| Child 04 facts | predecessor four-file set at closure commit `0aa49c87628c9e8b2041754515d6ebf0a930d55b` | accepted immutable syntax/direct-export input | P5-A/P5-B | accepted dependency / inspect-only | `E4-PNC-CLOSE1`, `E4-PNC-COMMIT1` | consume accepted facts; do not regenerate syntax; refresh current source owner before any edit |
| Child 02 affected readers | current reader-impact inventory | preservation consumer | P5-D | inspect; edit only named affected rows | `E0-P0A-SCOPE1` | edit only named affected rows |
| `E:\cheapapp.org` source | target `.anvien` output | real integration subject | P5-D | source preserve; normal analyze/read only | `E0-P0A-BOUNDARY1` | no copy, fixture, report, or source edit in target |

## Detailed Findings

### Module found does not imply export found

Current state:

- `ImportFact` carries source file, local/imported names, raw target, and optional resolved fields.
- TS re-export syntax is currently represented through the same import-shaped fact.
- `resolveImportFiles` can resolve the bounded source imports to the barrel.
- `resolveImportedDef` immediately scans physical definitions in that target file.

Required state:

```text
source import -> module/file result -> export-table lookup -> terminal Symbol or explicit unresolved result
```

Classification: path result is correct for the bounded case; export lookup is missing/wrong.

Allowed next action: P5-B adds the export boundary at the evidence-selected dedicated owner and minimal workspace wiring; P5-C later performs terminal traversal.

Forbidden next action: treating every physical definition as exported or rebuilding all path/package handling before evidence requires it.

### Requested meaning and type-only contract

Current state:

- Default/named/alias imported names and raw module text remain exact.
- Namespace imports retain no exported-name request and use kind plus raw module text.
- Source-written TS/JS imports carry canonical `RequestedMeanings` and `TypeOnly`; compatibility re-export and non-TS/JS imports retain empty requested fields while accepted `ExportFact` remains the re-export semantic SSOT.

Required state:

```text
source-written TS/JS ImportFact
  -> RequestedMeanings is a canonical allowed-set
  -> plain default/named/alias: {value,type,namespace}
  -> statement-level or inline type-only: {type}, TypeOnly=true
  -> plain namespace: no exported name, {namespace}
  -> type-only namespace: no exported name, {type}, TypeOnly=true
compatibility re-export ImportFact
  -> requested fields empty; accepted ExportFact owns name/meaning/type-only/provenance
```

Classification: requested name and requested meaning/type-only are `correct` in the P5-A Coder candidate; Supervisor acceptance remains pending.

Allowed next action: preserve the exact candidate and hand it to Supervisor; do not open P5-B until P5-A is accepted and committed by Main.

Forbidden next action: infer re-export semantics from compatibility imports, add side-effect-only facts, activate dormant `Target*` fields, change path candidates, or open P5-B+.

### Terminal binding and global-name rescue

Current state:

- Absence of a barrel physical definition prevents the import binding.
- The two bounded calls therefore have no terminal `CALLS` relationship.
- Current call resolution contains a repository-global name branch after scoped/same-file paths.

Required state:

```text
explicit import -> exact exported-name/meaning result -> terminal call proof
explicit export failure -> explicit unresolved result, never a global same-name target
```

Classification: re-export traversal is wrong; terminal call binding is unbound; the global-name-rescue boundary is wrong for explicit imports.

Allowed next action: P5-C resolves proof-bearing terminal results; P5-D binds only their affected graph/reader consumers.

Forbidden next action: adding a target-name special case or accepting `2/2` without direct/barrel identity and proof equality.

## Next Phase Status Decisions

| Plan Item | Actual Status Finding | Required Status / Next-Action Update |
|-----------|-----------------------|--------------------------------------|
| P5-A | exact four-owner production candidate plus two focused tests passed build/boundary/regression; Supervisor PASS and isolated commits `2560f914` + `40ea0095` are present | committed/accepted; preserve the three `0` deltas and do not reopen P5-A |
| P5-B | dedicated table owner plus minimal wiring and focused tests pass exact absolute E build, fixed-corpus counts, and Supervisor resubmission review | detect `E5-P5B-DETECT1` and isolated commit `E5-P5B-COMMIT1` are recorded; P5-B is closed and P5-C is the sole open slice |
| P5-C | final four-file candidate implements the authorized traversal/proof owner, two-phase import orchestration, namespace sibling, narrow `resolveCall` guard, and the reject-only distinct-owner member repair; final Supervisor report records PASS and fresh detect is recorded | keep P5-C unchecked while commit remains pending; create only the isolated P5-C commit before opening P5-D |
| P5-D | terminal graph/persistence/readers are not yet bound | refresh impact and edit only actual affected consumers; prove exact two sites |

## Implementation Gate

- [x] Target scope is limited to Child 05 responsibilities.
- [x] Each current unit has a status and exact P0 evidence IDs.
- [x] Current target files have file-detail related-file counts.
- [x] CRITICAL symbol impacts and blast-radius counts are recorded.
- [x] Correct bounded path behavior is marked preserve-only.
- [x] Missing/wrong/unbound behaviors have one owning slice.
- [x] Target and scanner boundaries are explicit.
- [x] R0 records the current repo/graph basis.
- [x] Child 04 export-fact handoff is accepted at `0aa49c87628c9e8b2041754515d6ebf0a930d55b` and reflected in refresh row R1.
- [x] P5-A editable owners, requested-meaning/type-only representation, side-effect disposition, and all three absolute count denominators are refreshed before implementation.
- [x] P5-B exact owner, preserve-only surfaces, CRITICAL blast radius, and E-only boundary are recorded in `E5-P5B-IMPACT1` before implementation.
- [x] P5-B source/build/test/count evidence and the exact evidence-blocker resubmission are accepted by Supervisor in `E5-P5B-REVIEW1`.
- [x] P5-B fresh graph refresh and full detect-changes result are recorded in `E5-P5B-DETECT1`; the expected pre-commit analyzer gap is classified and does not widen scope.
- [x] P5-B exact 11-path isolated commit is recorded in `E5-P5B-COMMIT1` at `c1559df953a277b099009f8489576d00ed25aa58`.
- [x] P5-C is opened only after the P5-B implementation and docs transition commits; P5-D and target remain locked.
- [x] P5-C fresh graph/file-detail/upstream impact, exact owners, and two-phase sequencing are recorded in `E5-P5C-IMPACT1` before production edits.
- [x] P5-C REJECT, exact reject-only repair, final source/test/build identities, and Supervisor PASS are recorded as `E5-P5C-REPAIR1` and `E5-P5C-REVIEW1`; the slice remains unchecked until detect and isolated commit.
- [x] P5-C one fresh graph refresh and full detect result are recorded as `E5-P5C-DETECT1`; no accepted inventory/file-detail/impact gate was rerun.

## Final P0 Decision

Choose one:

- [ ] P0 actual-status incomplete. Implementation is blocked.
- [x] P0 complete. Next phase can proceed unchanged.
- [ ] P0 complete. Next phase status, next action, or work steps must be updated before implementation.
- [ ] P0 complete. Target scope is preserve-only.
- [ ] P0 complete. Implementation is blocked by missing predecessor evidence.

Decision note:

Child 04 is closed at `0aa49c87628c9e8b2041754515d6ebf0a930d55b`. P5-A is accepted and committed at `2560f914334e65961f755febdda6585840a4260e`; P5-B has Supervisor PASS, detect evidence, and isolated commit `c1559df953a277b099009f8489576d00ed25aa58`. P5-C has a final repaired four-file candidate with the initial Coder-ready evidence, source-backed REJECT, reject-only repair, final identities, Supervisor PASS, and fresh detect recorded through `E5-P5C-DETECT1`. P5-C remains unchecked while only `E5-P5C-COMMIT1` is pending. P5-D and target access remain locked.
