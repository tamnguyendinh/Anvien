# Anvien Ambient And External Resolution Actual Status

Title: Anvien Ambient And External Resolution
Date: 2026-07-28
Status: P0 Complete / Dependency Blocked
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
| Child 05 handoff | repository module/export outcomes are not yet implemented/accepted | immutable repository result and proof available before external lookup | blocked | predecessor scope | `E0-P0A-DEPEND1` | P6-A waits for Child 05 closure |
| Resolver workspace | `buildWorkspace` indexes only supplied repository ScopeIR definitions/scopes/bindings | declaration authority available through a source-backed P6-A design | missing | 46 related files | `E0-P0A-SRC1`, `E0-P0A-FD1`, `E0-P0A-IMPACT1` | P6-A traces; P6-B implements accepted authority |
| Supported TypeScript config inputs | current P0 source proves no dedicated declaration-universe profile; exact existing config inputs and product commitments are not yet inventoried | explicit supported inputs and unavailable behavior selected by P6-A | blocked | owner inventory pending P6-A | `E0-P0A-SRC1`, `E0-P0A-SCOPE1` | decide before P6-B |
| TypeScript standard-library authority | no declaration source supplies `Promise` or `Math` members to the resolver workspace | general declaration-backed lookup selected by P6-A; no target-name branch | missing | resolver impact spans 6 files / 5 modules | `E0-P0A-VERIFY1`, `E0-P0A-SRC1..E0-P0A-SRC3` | P6-B after P6-A |
| Project/package declaration lookup | originating report proposes broader sources, but bounded C5 evidence does not prove the exact required project/package scope | implement only P6-A-evidenced cases or preserve without code | blocked | unknown until P6-A | `E0-P0A-ORIGIN1`, `E0-P0A-SCOPE1` | conditional P6-C1 |
| External target representation | current graph uses internal definition targets and generic unresolved diagnostics; actual external consumer needs are not inventoried | minimum truthful representation chosen from actual consumers | blocked | affected-reader inventory pending | `E0-P0A-SRC2`, `E0-P0A-SCOPE2` | P6-A inventory; P6-C2 if required |
| Structured resolution outcomes | resolver emits resolved references or generic unresolved diagnostics; no accepted external capability result exists | one structured target/reason/stage/proof per affected source site | missing | 50 related files at resolver owner | `E0-P0A-SRC2`, `E0-P0A-FD2` | P6-C3 after lookup/representation |
| Graph-health classification | `classifyDiagnostic` infers from target text using Go-oriented builtin/stdlib/external qualifier tables | mechanical projection from resolver outcome | wrong | 29 related files; exact symbol affects 2 files | `E0-P0A-SRC4`, `E0-P0A-FD3`, `E0-P0A-IMPACT4` | P6-D after P6-C3 |
| Bounded target sites | TypeScript oracle resolves all three; Anvien records gaps and later labels them as in-repository/analyzer issues | correct external or accepted external-capability outcome for all three; no in-repository analyzer gap | wrong | 3 exact sites | `E0-P0A-VERIFY1`, `E0-P0A-TARGET1` | terminal P6-D acceptance |
| Affected persistence/readers | no current Child 06 evidence establishes which new target/outcome fields need adapters | edit only Child 02/fresh-impact affected consumers | blocked | unknown until P6-A/P6-C2 | `E0-P0A-SCOPE2` | no fixed reader or transport matrix |
| Target boundary | target analyzed in place; source/worktree are not implementation surfaces | preserve target source; regenerate only normal target index during P6-D | correct | accepted target graph: 114,125 relationships | `E0-P0A-BOUNDARY1` | preserve until P6-D pre/post evidence |

## Status Refresh Log

| Refresh | Date | Repo Basis | Changed Scope | Status Changes | Evidence | Next Phase Update |
|---------|------|------------|---------------|----------------|----------|-------------------|
| R0 | 2026-08-10 | HEAD `238aec06d28286acc0fca0f4e6a69f9eb4ff6a49`; graph file-detail `stale=false`, analyzed `2026-08-09T19:19:54Z` | Child 06 source/report/ledger reset | removed preselected authority, source-category, public-option, and adapter assumptions; P6-A becomes mandatory design slice; P6-C1/C2 conditional on evidence | `E0-P0A-GRAPH1`, `E0-P0A-SRC1..E0-P0A-SRC4`, `E0-P0A-FD1..E0-P0A-FD3`, `E0-P0A-IMPACT1..E0-P0A-IMPACT4` | open P6-A only after Child 05 handoff; no production authority mechanism selected yet |

## Phase Touch Map

| Unit / File / Surface | Plan-Relevant Relationship File | Relationship to Target | Plan Item | Touch Mode | Evidence | Constraint |
|-----------------------|---------------------------------|------------------------|-----------|------------|----------|------------|
| repository workspace | `internal/resolution/indexes.go` | builds all current resolution indexes from repository facts | P6-A/P6-B | inspect-only in P6-A; later minimum edit only after decision/impact | `E0-P0A-FD1`, `E0-P0A-IMPACT1` | CRITICAL; preserve unrelated languages/lookup behavior |
| call/type/member resolution | `internal/resolution/resolve.go` | consumes workspace and emits resolved/gap evidence | P6-A/P6-C3 | inspect-only until exact finalizer impact | `E0-P0A-FD2`, `E0-P0A-IMPACT2`, `E0-P0A-IMPACT3` | no target-name branch or stage overwrite |
| health classification | `internal/graphhealth/diagnostics.go` | currently guesses classification from target text | P6-A/P6-D | inspect in P6-A; edit same-invariant projection in P6-D | `E0-P0A-FD3`, `E0-P0A-IMPACT4` | health cannot invent resolver truth |
| TypeScript config/provider owners | exact files pending source inventory | inputs to authority decision | P6-A/P6-B | inspect-only until P6-A owner map | `E0-P0A-SCOPE1` | no file-name-based assumption |
| Child 05 result | predecessor four-file set | repository terminal/unresolved input | P6-A/P6-C3 | dependency / inspect-only | `E0-P0A-DEPEND1` | no duplicate module/export traversal |
| Child 02 affected readers | current reader-impact inventory | target/outcome preservation | P6-A/P6-C2/P6-C3/P6-D | inspect; edit only named affected rows | `E0-P0A-SCOPE2` | no broad all-reader or transport expansion |
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

Classification: standard-library authority is missing; project/package scope and external representation are blocked pending design/consumer evidence.

Allowed next action: measure real inputs and consumers and update later slice steps before code.

Forbidden next action: importing a prewritten file/source/transport matrix into the implementation plan.

## Next Phase Status Decisions

| Plan Item | Actual Status Finding | Required Status / Next-Action Update |
|-----------|-----------------------|--------------------------------------|
| P6-A | cause is proved but authority/config/consumer design is not | inspect and decide; production source remains preserve-only |
| P6-B | standard-library authority is missing | implement only the accepted P6-A mechanism and behavior |
| P6-C1 | project/package scope is unproved | activate or close preserve-only from P6-A evidence |
| P6-C2 | external representation/consumers are unproved | select minimum representation after consumer inventory |
| P6-C3 | current generic gap is insufficient | define only statuses/fields required by accepted P5/P6 cases |
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
- [ ] Child 05 handoff is accepted and reflected in a new refresh row.
- [ ] P6-A records the selected authority/config/consumer decision and updates later slice steps before production code.

## Final P0 Decision

Choose one:

- [ ] P0 actual-status incomplete. Implementation is blocked.
- [ ] P0 complete. Next phase can proceed unchanged.
- [ ] P0 complete. Next phase status, next action, or work steps must be updated before implementation.
- [ ] P0 complete. Target scope is preserve-only.
- [x] P0 complete. Implementation is blocked by missing predecessor evidence.

Decision note:

P0 proves the missing workspace declaration authority and downstream text inference with current source and blast-radius evidence. P6-A opens after Child 05 closure and must decide the implementation contract before any production authority mechanism is written.
