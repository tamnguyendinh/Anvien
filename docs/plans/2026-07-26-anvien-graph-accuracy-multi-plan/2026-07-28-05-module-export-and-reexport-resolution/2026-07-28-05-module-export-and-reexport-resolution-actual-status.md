# Anvien Module Export And Re-Export Resolution Actual Status

Title: Anvien Module Export And Re-Export Resolution
Date: 2026-07-28
Status: P0 Complete / Dependency Blocked
Companion plan: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-05-module-export-and-reexport-resolution/2026-07-28-05-module-export-and-reexport-resolution-plan.md`
Companion evidence: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-05-module-export-and-reexport-resolution/2026-07-28-05-module-export-and-reexport-resolution-evidence.md`
Companion benchmark: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-05-module-export-and-reexport-resolution/2026-07-28-05-module-export-and-reexport-resolution-benchmark.md`
Roadmap: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-26-anvien-graph-accuracy-roadmap.md`
Predecessor: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-04-typescript-export-semantics/2026-07-28-04-typescript-export-semantics-plan.md`
Successor: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-plan.md`

## Purpose

This file records the current module-request, path lookup, export-surface, re-export traversal, terminal call-binding, and target-boundary state before Child 05 production work.

Implementation does not start until every editable owner has fresh evidence and the Child 04 export-fact handoff is accepted. Detailed proof belongs in the evidence ledger.

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

The graph used for this P0 inspection reported current commit parity and `stale=false` on every listed file-detail result.

| Unit / File / Surface | File Detail Evidence | Related File Count | Relationship Summary | Impact Note |
|-----------------------|----------------------|-------------------:|----------------------|-------------|
| `internal/resolution/indexes.go` | `E0-P0A-FD1` | 46 | workspace definitions, module lookup, import binding, physical-definition lookup | high file risk; three Child 05 symbols are CRITICAL |
| `internal/providers/tsjs/imports.go` | `E0-P0A-FD2` | 15 | current TS import/re-export syntax to `ImportFact` | high file risk; inspect until Child 04 handoff |
| `internal/resolution/import_resolution.go` | `E0-P0A-FD3` | 33 | existing language/path strategies and wildcard synthesis | high file risk; preserve unaffected languages |
| `internal/resolution/resolve.go` | `E0-P0A-FD4` | 50 | call/access/type resolution and global-name branch | high file risk; `resolveCall` CRITICAL |

| Symbol | Impact Evidence | Risk | Impacted Symbols | Affected Files | Modules | Processes | Linked Flows / Tests |
|--------|-----------------|------|-----------------:|---------------:|--------:|----------:|----------------------:|
| `resolveImportedDef` | `E0-P0A-IMPACT1` | CRITICAL | 4 | 3 | 1 | 16 | 29 / 23 from containing file |
| `resolveImports` | `E0-P0A-IMPACT2` | CRITICAL | 6 | 5 | 3 | 18 | 29 / 23 from containing file |
| `buildWorkspace` | `E0-P0A-IMPACT3` | CRITICAL | 8 | 6 | 5 | 28 | 29 / 23 from containing file |
| `resolveCall` | `E0-P0A-IMPACT4` | CRITICAL | 6 | 4 | 4 | 35 | 21 / 26 from containing file |

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
| Child 04 fact handoff | current TS provider emits import/re-export-shaped `ImportFact`, but Child 04 has not yet implemented and accepted its export-fact handoff | immutable accepted module/import/export facts with name and meaning needed by P5 | blocked | 15 related files at current syntax owner | `E0-P0A-SRC1`, `E0-P0A-FD2` | P5-A waits for Child 04 closure |
| Bounded target module path | accepted investigation resolves the two source imports to the barrel file | preserve the source-written module/file result and its accounting | correct | 33 related files at path owner | `E0-P0A-VERIFY1`, `E0-P0A-SRC2`, `E0-P0A-FD3` | P5-A records and preserves the working path result |
| General module-request/path input contract | current logic has relative/index candidates and multiple language strategies, but no Child 05 manifest separates the exact inputs needed before export lookup | explicit current input/result contract; only proved gaps changed | partial | 33 related files | `E0-P0A-SRC1`, `E0-P0A-SRC2`, `E0-P0A-FD3` | P5-A inventory before code; no assumed rewrite |
| Module export surface | current workspace indexes definitions by file/name and has no first-class export table | deterministic table derived only from Child 04 facts | missing | 46 related files at workspace owner | `E0-P0A-SRC3`, `E0-P0A-FD1` | P5-B after P5-A |
| Re-export traversal | `resolveImportedDef` searches physical definitions in the resolved file and does not follow the barrel binding | terminal traversal with alias/star/cycle/ambiguity/meaning proof | wrong | 3 affected files for the exact symbol | `E0-P0A-VERIFY1`, `E0-P0A-SRC3`, `E0-P0A-IMPACT1` | P5-C after table acceptance |
| Explicit-import global-name-rescue boundary | missing scoped/import binding can reach current global-name call lookup; low-confidence matches become gaps, but the explicit import failure is not represented at its own export boundary | no repository-global same-name rescue; explicit export failure retained | wrong | 4 affected files for `resolveCall` | `E0-P0A-SRC4`, `E0-P0A-IMPACT4` | P5-C owns the no-global-rescue proof |
| Terminal call/proof emission | the two accepted target calls lack terminal `CALLS` relationships | both sites bind to the expected terminal Symbols with complete proof | unbound | resolver/emission impact must be refreshed in P5-D | `E0-P0A-VERIFY1` | P5-D after P5-C |
| Affected persistence/readers | no current Child 05 evidence identifies which corrected proof fields require adapters | only consumers named by Child 02 inventory and fresh P5-D impact are editable | blocked | unknown until Child 02/P5-D evidence | `E0-P0A-SCOPE1` | do not assume a fixed reader denominator |
| Target boundary | accepted target is analyzed in place and its source is not a fixture or edit surface | preserve source/worktree; regenerate only normal target index during P5-D validation | correct | accepted target graph: 114,125 relationships | `E0-P0A-BOUNDARY1` | preserve until P5-D pre/post evidence |

## Status Refresh Log

| Refresh | Date | Repo Basis | Changed Scope | Status Changes | Evidence | Next Phase Update |
|---------|------|------------|---------------|----------------|----------|-------------------|
| R0 | 2026-08-10 | HEAD `238aec06d28286acc0fca0f4e6a69f9eb4ff6a49`; graph file-detail `stale=false`, analyzed `2026-08-09T19:19:54Z` | Child 05 source/report/ledger reset | removed campaign-wide assumptions; path behavior classified separately from export lookup; P5-A blocked only by Child 04 handoff | `E0-P0A-GRAPH1`, `E0-P0A-SRC1..E0-P0A-SRC4`, `E0-P0A-FD1..E0-P0A-FD4`, `E0-P0A-IMPACT1..E0-P0A-IMPACT4` | open P5-A only after predecessor refresh; inventory before code |

## Phase Touch Map

| Unit / File / Surface | Plan-Relevant Relationship File | Relationship to Target | Plan Item | Touch Mode | Evidence | Constraint |
|-----------------------|---------------------------------|------------------------|-----------|------------|----------|------------|
| current import/export syntax | `internal/providers/tsjs/imports.go` | producer of current `ImportFact`; Child 04 owner | P5-A/P5-B | inspect-only | `E0-P0A-FD2`, `E0-P0A-SRC1` | consume accepted Child 04 output; no syntax reimplementation |
| current workspace/import binding | `internal/resolution/indexes.go` | physical-definition lookup and import binding | P5-A/P5-B/P5-C | inspect-only until slice impact; then minimum edit | `E0-P0A-FD1`, `E0-P0A-IMPACT1..E0-P0A-IMPACT3` | CRITICAL; new semantics need one dedicated responsibility |
| current path strategies | `internal/resolution/import_resolution.go` | module/file resolution and other-language strategies | P5-A | inspect first; preserve unaffected branches | `E0-P0A-FD3`, `E0-P0A-SRC2` | no broad package/path rewrite without evidence |
| terminal call resolution | `internal/resolution/resolve.go` | consumes scope/import bindings and contains global lookup | P5-C/P5-D | inspect-only until fresh impact | `E0-P0A-FD4`, `E0-P0A-IMPACT4` | explicit import failure cannot use global rescue |
| Child 04 facts | predecessor four-file set | required immutable input | P5-A/P5-B | dependency / inspect-only | `E0-P0A-DEPEND1` | no implementation before accepted handoff |
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

Allowed next action: P5-A inventories the real input boundary and preserves the path behavior; P5-B/P5-C add the export boundary at evidence-selected owners.

Forbidden next action: treating every physical definition as exported or rebuilding all path/package handling before evidence requires it.

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
| P5-A | bounded module path works; exact fact/meaning contract depends on Child 04 | wait for predecessor, then inventory/preserve before any edit |
| P5-B | no export table exists | select exact owner from fresh impact and build only from accepted facts |
| P5-C | physical-definition lookup and explicit-import global-name rescue are wrong | add traversal/proof and no-global-rescue behavior after table acceptance |
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
- [ ] Child 04 export-fact handoff is accepted and reflected in a new refresh row.
- [ ] P5-A editable owners and absolute path/`IMPORTS` counts are refreshed immediately before implementation.

## Final P0 Decision

Choose one:

- [ ] P0 actual-status incomplete. Implementation is blocked.
- [ ] P0 complete. Next phase can proceed unchanged.
- [ ] P0 complete. Next phase status, next action, or work steps must be updated before implementation.
- [ ] P0 complete. Target scope is preserve-only.
- [x] P0 complete. Implementation is blocked by missing predecessor evidence.

Decision note:

P0 establishes the current Child 05 boundary with source, file-detail, and impact evidence. P5-A opens only after the accepted Child 04 fact handoff is recorded and this file is refreshed. Package/path changes require their own source and impact evidence.
