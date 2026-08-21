# Anvien Ambient And External Resolution Actual Status

Title: Anvien Ambient And External Resolution
Date: 2026-07-28
Status: P0 Complete / P6-A Supervisor PASS / Exact Decision Commit Pending
Companion plan: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-plan.md`
Companion evidence: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-evidence.md`
Companion benchmark: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-benchmark.md`
Roadmap: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-26-anvien-graph-accuracy-roadmap.md`
Predecessor: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-05-module-export-and-reexport-resolution/2026-07-28-05-module-export-and-reexport-resolution-plan.md`
Successor: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-07-cross-surface-acceptance-and-target-validation/2026-07-28-07-cross-surface-acceptance-and-target-validation-plan.md`

## Purpose

This file records the current declaration inputs, resolver workspace, standard-library/external lookup, outcome, graph-health, target, and affected-reader state before Child 06 work.

P6-A must turn unknown design questions into a source-backed decision before production implementation. Detailed proof belongs in the evidence ledger.

## Freshness / Refresh Rules

Refresh this file:

- after Child 05 changes the repository resolution result consumed by Child 06;
- during P6-A when source/config/runtime evidence selects the declaration-authority design;
- before every production slice after fresh graph, file-detail, and impact evidence;
- after each accepted slice and whenever affected readers or owner files change.

Use explicit transitions and append refresh-log rows. Do not preserve superseded assumptions as active authority.

## Scope

Target scope:

- source-backed declaration-universe behavior and implementation decision;
- TypeScript standard-library authority selected by P6-A;
- project/package declaration lookup only when P6-A proves it required;
- necessary external target representation and structured outcomes;
- graph-health projection from resolver outcomes;
- exact target validation for `Promise`, `Math.max`, and `Math.min`;
- only persistence/readers proven affected by Child 02 inventory and fresh impact.

Out of scope:

- preselected authority/storage/runtime architecture;
- unevidenced declaration-source coverage or public traversal options;
- Child 05 module/export resolution;
- graph-output behavior, scanner behavior, target-source edits, and unrelated language semantics.

## Relationship / Impact Evidence

The graph used for this P0 inspection reported current commit parity and `stale=false` for every listed file-detail result.

| Unit / File / Surface | File Detail Evidence | Related File Count | Relationship Summary | Impact Note |
|-----------------------|----------------------|-------------------:|----------------------|-------------|
| `internal/resolution/indexes.go` | `E0-P0A-FD1` | 46 | workspace built from repository ScopeIR, name/type/member/import indexes | high file risk; `buildWorkspace` CRITICAL |
| `internal/resolution/resolve.go` | `E0-P0A-FD2` | 50 | call/member/type resolution and unresolved emission | high file risk; call/type symbols CRITICAL |
| `internal/graphhealth/diagnostics.go` | `E0-P0A-FD3` | 29 | diagnostic normalization and target-text classification | high file risk; classifier HIGH |

| Symbol | Impact Evidence | Risk | Impacted Symbols | Affected Files | Modules | Processes | Linked Flows / Tests |
|--------|-----------------|------|-----------------:|---------------:|--------:|----------:|----------------------:|
| `buildWorkspace` | `E0-P0A-IMPACT1` | CRITICAL | 8 | 6 | 5 | 28 | 29 / 23 from containing file |
| `resolveCall` | `E0-P0A-IMPACT2` | CRITICAL | 6 | 4 | 4 | 35 | 21 / 26 from containing file |
| `resolveTypeAnnotation` | `E0-P0A-IMPACT3` | CRITICAL | 6 | 4 | 4 | 35 | 21 / 26 from containing file |
| `classifyDiagnostic` | `E0-P0A-IMPACT4` | HIGH | 7 | 2 | 3 | 3 | 8 / 14 from containing file |

Fresh P6-A graph basis is HEAD `ec765debff335540c77d409ebb2c9f45e4a0a77d`, analyzed with `anvien analyze --force`: `1,985` scanned, `743` parsed, `0` failed, `116,467` nodes, `160,591` relationships, `stale=false`.

| P6-A owner / consumer | Related files | Fresh impact warning |
|-----------------------|--------------:|----------------------|
| `internal/analyze/analyze.go` / `analyze.Run` | 187 | CRITICAL `24` symbols / `9` files / `8` modules / `23` processes |
| `internal/resolution/types.go` | 34 | options/outcome shared owner; exact edited symbol must be refreshed in P6-B/C3 |
| `internal/resolution/indexes.go` / `buildWorkspace` | 58 | CRITICAL `59 / 24 / 8 / 23` |
| `internal/resolution/resolve.go` / call-access-type owners | 60 | each CRITICAL `28 / 11 / 7 / 32` |
| `internal/resolution/emit.go` / `emitUnresolvedReference` | 50 | CRITICAL `9 / 4 / 2 / 34` |
| `internal/graph/types.go` / `graph.Node` | 256 | CRITICAL `1,717 / 273 / 48 / 428`; validate/preserve-only unless later evidence requires edit |
| `internal/scopeir/kinds.go` / `scopeir.NodeLabel` | 250 | CRITICAL `1,846 / 291 / 47 / 422` |
| Ladybug CSV / schema | 23 / 22 | `nodeCSVRow` CRITICAL `25 / 11 / 1 / 12`; `NodeSchema` MEDIUM `8 / 6 / 2 / 0` |
| graph-health policy / diagnostics / ResolutionGap inputs | 31 / 30 / 41 | Diagnostic CRITICAL `70 / 17 / 5 / 77`; classifier HIGH `10 / 4 / 3 / 0`; inputs CRITICAL `13 / 7 / 1 / 15` |
| processes | 14 | `Apply` CRITICAL `29 / 8 / 7 / 31`; call graph CRITICAL `28 / 8 / 6 / 25` |
| MCP context / impact / rename | 28 / 41 / 9 | payload/collector CRITICAL `1 / 1 / 1 / 7`, `11 / 5 / 2 / 17`, `1 / 1 / 1 / 8` |
| HTTP graph / filecontext | 22 / 46 | generic transport/fallback; validate/preserve-only pending later impact |

HIGH/CRITICAL is a blast-radius warning requiring narrow edits and affected-boundary regression, not an edit ban.

## Status Rules

| Status | Meaning | Allowed next action |
|--------|---------|---------------------|
| `correct` | Current behavior satisfies this child's bounded requirement. | Preserve and regress. |
| `partial` | A relevant input/behavior exists but the accepted contract is incomplete. | Change only the proved gap. |
| `wrong` | Current behavior conflicts with the accepted requirement. | Replace at the proved owner. |
| `missing` | Required behavior does not exist. | Add it after owner/design evidence. |
| `unbound` | A result exists but is not connected to the truthful downstream consumer. | Bind only that boundary. |
| `blocked` | A predecessor, design decision, or impact inventory is missing. | Do not implement until refreshed. |

## Current Status Matrix

| Unit | Current State | Required State | Status | Relationship Count | Evidence | Next Plan Decision |
|------|---------------|----------------|--------|-------------------:|----------|--------------------|
| Problem authority | originating report records the false external gaps but labels its architecture DRAFT; bounded reports verify C5 only | findings/acceptance separated from design; current source decides implementation | correct | N/A | `E0-P0A-ORIGIN1`, `E0-P0A-VERIFY1`, `E0-P0A-VERIFY2` | preserve evidence hierarchy |
| Child 05 handoff | exact four-ledger closure commit `ec765debff335540c77d409ebb2c9f45e4a0a77d` is current HEAD; P5-A/P5-B/P5-C/P5-D and Pn-A/Pn-B are accepted and committed | immutable repository module/file, syntax-derived export-table, deterministic traversal, fail-closed explicit-import, and generic-first proof contract available before external lookup | correct | exact closure commit; four living ledgers | `E0-P0A-DEPEND1`, `E6-P6A-DEPEND1` | consume as immutable predecessor; open only P6-A |
| P6-A declaration-universe decision | offline generated/check-in/embedded TypeScript `5.9.3` catalog, narrow root config profile, final outcome contract, owner/reader map, P6-C1 preserve-only, and P6-C2 active are recorded in an immutable report; initial review rejected only the contract boundary, exact repair restored contract to HEAD, and reject-only review passed | accepted decision with no residual same-invariant surface; exact decision commit required before production | correct / Supervisor PASS | report `25,326` bytes / `289` LF / SHA-256 `77D5E9AC8D76D98C76D1816C8D6E69265D4AFB30367E3DA50DF3EAA3445D7BA2`; final review `6,011` bytes / `83` LF / SHA-256 `39CE249E9D13F1C77FE6F61DD6E9B1D2E4000B004CE3EBBF461C762F3FA28384` | `E6-P6A-DECISION1`, `E6-P6A-REVIEW1` | commit exact seven-path decision manifest, then open P6-B with the resulting anchor |
| Resolver workspace | `buildWorkspace` indexes only supplied repository ScopeIR definitions/scopes/bindings | declaration authority available through a source-backed P6-A design | missing | 46 related files | `E0-P0A-SRC1`, `E0-P0A-FD1`, `E0-P0A-IMPACT1` | P6-A traces; P6-B implements accepted authority |
| Supported TypeScript config inputs | production reads no declaration profile; P6-A selects zero/one root JSONC `tsconfig.json` with only `target`, `lib`, `noLib`; absent config uses manifest default; invalid/unreadable/unsupported topology is capability-unavailable | implement exactly the selected profile without changing scanner corpus or guessing project ownership | missing | analyze/config owner plus new isolated `internal/tsstdlib` profile owner | `E6-P6A-SRC1`, `E6-P6A-ORACLE1`, `E6-P6A-DECISION1` | P6-B after Supervisor/commit |
| TypeScript standard-library authority | no declaration source supplies `Promise` or `Math` members to the resolver workspace | general declaration-backed lookup selected by P6-A; no target-name branch | missing | resolver impact spans 6 files / 5 modules | `E0-P0A-VERIFY1`, `E0-P0A-SRC1..E0-P0A-SRC3` | P6-B after P6-A |
| Project/package declaration lookup | current product has no project/package `.d.ts`, ambient-module, augmentation, or `node_modules` authority and current campaign evidence does not require it | preserve unavailable behavior; no production/test owner | correct | zero production owners | `E6-P6A-SRC1`, `E6-P6A-DECISION1` | P6-C1 closes preserve-only after P6-B; activation requires plan refresh |
| External target representation | no dedicated external node exists; Ladybug skips unknown labels and endpoint pairs, processes consume all CALLS, context/impact omit external provenance, and rename can search/edit graph targets | referenced-only deterministic `ExternalSymbol`; unavailable outcomes stay gaps; affected readers adapt without repository ownership pollution | missing | exact C2 owners: scope label, materializer/emitter, Ladybug CSV/schema, processes, MCP context/impact/rename | `E6-P6A-CONSUMER1`, `E6-P6A-DECISION1` | P6-C2 active after P6-B and preserve-only C1 closure |
| Structured resolution outcomes | resolver emits resolved references or generic unresolved diagnostics; no accepted external capability result exists | one structured target/reason/stage/proof per affected source site | missing | 50 related files at resolver owner | `E0-P0A-SRC2`, `E0-P0A-FD2` | P6-C3 after lookup/representation |
| Graph-health classification | `classifyDiagnostic` infers from target text using Go-oriented builtin/stdlib/external qualifier tables | mechanical projection from resolver outcome | wrong | 29 related files; exact symbol affects 2 files | `E0-P0A-SRC4`, `E0-P0A-FD3`, `E0-P0A-IMPACT4` | P6-D after P6-C3 |
| Bounded target sites | TypeScript oracle resolves all three; Anvien records gaps and later labels them as in-repository/analyzer issues | correct external or accepted external-capability outcome for all three; no in-repository analyzer gap | wrong | 3 exact sites | `E0-P0A-VERIFY1`, `E0-P0A-TARGET1` | terminal P6-D acceptance |
| Affected persistence/readers | graph JSON is generic; Ladybug requires explicit table/pairs; processes must filter external endpoints; context/impact must project provenance; rename must reject; health/ResolutionGap later consume outcome; HTTP/Web/file-detail are preserve/validate-only | edit only the named semantic consumers after fresh slice impact; preserve generic transports unless invalidated | partial | Ladybug `23/22`; processes `14`; MCP `28/41/9`; health `31/30/41`; preserve candidates `256/22/46` related files | `E6-P6A-CONSUMER1`, `E6-P6A-IMPACT1`, `E6-P6A-DECISION1` | P6-C2/C3/D implement exact active rows |
| Target boundary | target analyzed in place; source/worktree are not implementation surfaces | preserve target source; regenerate only normal target index during P6-D | correct | accepted target graph: 114,125 relationships | `E0-P0A-BOUNDARY1` | preserve until P6-D pre/post evidence |

## Status Refresh Log

| Refresh | Date | Repo Basis | Changed Scope | Status Changes | Evidence | Next Phase Update |
|---------|------|------------|---------------|----------------|----------|-------------------|
| R0 | 2026-08-10 | HEAD `238aec06d28286acc0fca0f4e6a69f9eb4ff6a49`; graph file-detail `stale=false`, analyzed `2026-08-09T19:19:54Z` | Child 06 source/report/ledger reset | removed preselected authority, source-category, public-option, and adapter assumptions; P6-A becomes mandatory design slice; P6-C1/C2 conditional on evidence | `E0-P0A-GRAPH1`, `E0-P0A-SRC1..E0-P0A-SRC4`, `E0-P0A-FD1..E0-P0A-FD3`, `E0-P0A-IMPACT1..E0-P0A-IMPACT4` | open P6-A only after Child 05 handoff; no production authority mechanism selected yet |
| R1 | 2026-08-22 | HEAD and exact Child 05 closure commit `ec765debff335540c77d409ebb2c9f45e4a0a77d`; parent/Pn-B `fcc44334c0f75b3b19046dc8f9f4de40eb459fa9`; Pn-A `b68e738d64eebea65a045afbf0b12d94dd43cbf4` | predecessor closure and Child 06 opening boundary | Child 05 dependency `blocked -> correct`; P6-A `locked -> open`; all production/test/target slices remain locked | `E6-P6A-DEPEND1` | execute only source/config/runtime/consumer/oracle inventory and P6-A decision documents; no production/test/target action |
| R2 | 2026-08-22 | same HEAD; fresh graph `1,985` scanned / `743` parsed / `0` failed / `116,467` nodes / `160,591` relationships / `stale=false`; immutable decision report SHA-256 `77D5E9AC8D76D98C76D1816C8D6E69265D4AFB30367E3DA50DF3EAA3445D7BA2` | P6-A source/config/package/consumer/oracle decision only | authority/config/consumer design `blocked -> decision-recorded`; P6-C1 `conditional -> preserve-only`; P6-C2 `conditional -> active`; P6-A review/commit and every production/test/target slice remain locked; `03:30 +07` report deadline missed; contract `18+/0-` scope deviation disclosed and frozen | `E6-P6A-IMPACT1`, `E6-P6A-SRC1`, `E6-P6A-ORACLE1`, `E6-P6A-CONSUMER1`, `E6-P6A-DECISION1`, `E6-P6A-BUILD1` | route immutable candidate to Supervisor; Main decides late-handoff/deviation acceptance; no production action |
| R3 | 2026-08-22 | initial Supervisor REJECT `5679FB24...6290`; contract restored byte-for-byte to HEAD blob `2020b479...`; reject-only Supervisor PASS `39CE249E...8384`; fresh Main planner graph `116,521 / 160,645` | sole sealed-boundary blocker repair and P6-A acceptance/commit transition | contract `18+/0- -> zero diff`; unchanged four-ledger/report technical candidate `REJECT -> PASS`; deadline miss remains disclosed/non-blocking; P6-A `decision-recorded -> Supervisor PASS`; P6-B remains production-locked until exact commit | `E6-P6A-REVIEW1`, `E6-P6A-COMMIT1` | commit exact seven-path P6-A decision manifest immediately, then hand the resulting commit anchor directly to P6-B |

## Phase Touch Map

| Unit / File / Surface | Plan-Relevant Relationship File | Relationship to Target | Plan Item | Touch Mode | Evidence | Constraint |
|-----------------------|---------------------------------|------------------------|-----------|------------|----------|------------|
| repository workspace | `internal/resolution/indexes.go` | builds all current resolution indexes from repository facts | P6-A/P6-B | inspect-only in P6-A; later minimum edit only after decision/impact | `E0-P0A-FD1`, `E0-P0A-IMPACT1` | CRITICAL; preserve unrelated languages/lookup behavior |
| call/type/member resolution | `internal/resolution/resolve.go` | consumes workspace and emits resolved/gap evidence | P6-A/P6-C3 | inspect-only until exact finalizer impact | `E0-P0A-FD2`, `E0-P0A-IMPACT2`, `E0-P0A-IMPACT3` | no target-name branch or stage overwrite |
| health classification | `internal/graphhealth/diagnostics.go` | currently guesses classification from target text | P6-A/P6-D | inspect in P6-A; edit same-invariant projection in P6-D | `E0-P0A-FD3`, `E0-P0A-IMPACT4` | health cannot invent resolver truth |
| offline catalog/config owners | new `anvien-web/scripts/generate-tsstdlib-catalog.mjs`, new `internal/tsstdlib/catalog.v1.json`, `catalog.go`, `profile.go`, plus minimum analyze/resolution wiring | exact authority/config inputs selected by P6-A | P6-B | locked until P6-A PASS/commit; then fresh impact before edit | `E6-P6A-SRC1`, `E6-P6A-DECISION1` | no runtime Node/network/install/script; no name map |
| Child 05 result | predecessor four-file set | repository terminal/unresolved input | P6-A/P6-C3 | dependency / inspect-only | `E0-P0A-DEPEND1` | no duplicate module/export traversal |
| affected readers | Ladybug CSV/schema, processes, MCP context/impact/rename, graph-health/ResolutionGap; graph JSON/HTTP/Web/file-detail validate/preserve-only | external target/outcome preservation | P6-C2/P6-C3/P6-D | edit only named active semantic rows after fresh impact | `E6-P6A-CONSUMER1`, `E6-P6A-IMPACT1` | no broad all-reader/UI expansion |
| `E:\cheapapp.org` source | target `.anvien` output | real integration subject | P6-D | source preserve; normal analyze/read only | `E0-P0A-BOUNDARY1` | no copy, fixture, report, or source edit in target |

## Detailed Findings

### Missing declaration authority precedes graph-health

Current state:

- `buildWorkspace` accepts repository ScopeIR and indexes its definitions, scopes, bindings, imports, members, and types.
- No current input supplies TypeScript standard-library declarations.
- `resolveTypeAnnotation` exempts a small primitive list, then queries the repository workspace.
- Member calls require receiver/type/member owners in the same workspace.
- `classifyDiagnostic` later infers categories from target strings and Go-focused tables.

Required state:

```text
source/config -> accepted declaration authority -> resolver target or explicit capability result
resolver outcome -> graph/persistence -> graph-health projection
```

Classification: declaration authority and structured outcomes are missing; graph-health classification is wrong for this invariant.

Allowed next action: P6-A chooses the authority and required target/outcome facts from source/config/runtime/consumer evidence; later slices implement only that accepted decision.

Forbidden next action: changing only graph-health labels or hardcoding the three target names.

### Broader declaration scope is not yet proved

Current state:

- The originating report proposes standard-library, project, package, ambient, augmentation, and intrinsic sources.
- The accepted bounded investigation proves the target workspace lacks TypeScript ambient/lib authority.
- It explicitly leaves broader declaration semantics and policy unresolved.

Required state:

```text
P6-A evidence decides which sources/config cases this campaign owns.
P6-B implements standard-library authority.
P6-C1/P6-C2 activate only for proved lookup/representation needs.
```

Classification after P6-A: standard-library authority is missing implementation; project/package scope is preserve-only; referenced external representation is required and missing implementation.

Allowed next action: independent Supervisor reviews the recorded decision. Only after PASS and commit may P6-B implement the exact embedded-catalog authority.

Forbidden next action: importing a prewritten file/source/transport matrix into the implementation plan.

## Next Phase Status Decisions

| Plan Item | Actual Status Finding | Required Status / Next-Action Update |
|-----------|-----------------------|--------------------------------------|
| P6-A | authority/config/outcome/owner/consumer decision has final reject-only Supervisor PASS; contract deviation is removed; commit pending | commit exact seven-path decision manifest, then open only P6-B |
| P6-B | standard-library authority is missing; exact embedded-catalog mechanism and root config subset are selected | implement only after P6-A PASS/commit; run fresh impact, full build, real runtime, determinism/size/performance gates |
| P6-C1 | current evidence does not require project/package declaration lookup | close preserve-only after P6-B; no production/test owner |
| P6-C2 | referenced external targets require dedicated representation and named reader adapters | active after P6-B/C1; implement referenced-only `ExternalSymbol` and affected parity |
| P6-C3 | current generic gap is insufficient; minimum status/reason/provenance contract is selected | implement one final versioned outcome per site after C2; refresh shared-struct plan only if Note carriage is not lossless |
| P6-D | health guesses from text and target is `0/3` | map outcomes mechanically and prove the three exact sites |

## Implementation Gate

- [x] Target scope is limited to Child 06 responsibilities.
- [x] Each current unit has a status and exact P0 evidence IDs.
- [x] Current target files have file-detail related-file counts.
- [x] HIGH/CRITICAL symbol impacts and blast-radius counts are recorded.
- [x] Problem origin, bounded verification, and proposed design are distinguished.
- [x] Missing, wrong, and blocked behaviors have one owning slice.
- [x] Target/scanner and predecessor boundaries are explicit.
- [x] R0 records the current repo/graph basis.
- [x] Child 05 handoff is accepted at exact closure commit `ec765debff335540c77d409ebb2c9f45e4a0a77d` and reflected in refresh row R1.
- [x] P6-A records the selected authority/config/consumer decision and updates later slice steps before production code.
- [x] P6-A immutable report records the missed `03:30 +07` deadline instead of claiming timely delivery.
- [x] Independent Supervisor accepts the P6-A decision after exact contract-boundary repair; Main is authorized to create the isolated decision commit before P6-B.

## Final P0 Decision

Choose one:

- [ ] P0 actual-status incomplete. Implementation is blocked.
- [ ] P0 complete. Next phase can proceed unchanged.
- [x] P0 complete. Next phase status, next action, or work steps must be updated before implementation.
- [ ] P0 complete. Target scope is preserve-only.
- [ ] P0 complete. Implementation is blocked by missing predecessor evidence.

Decision note:

P0 proves the missing workspace declaration authority and downstream text inference with current source and blast-radius evidence. Child 05 is closed at `ec765debff335540c77d409ebb2c9f45e4a0a77d`; P6-A has reject-only Supervisor PASS and closes through the immediate exact decision commit. Only after that commit may P6-B open; all production, later slices, and target access remain locked until their own gates.
