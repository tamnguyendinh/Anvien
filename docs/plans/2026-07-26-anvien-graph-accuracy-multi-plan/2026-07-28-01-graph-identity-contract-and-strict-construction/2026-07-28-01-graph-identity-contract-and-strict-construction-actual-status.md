# Anvien Graph Identity Contract and Strict Construction Actual Status

Title: Anvien Graph Identity Contract and Strict Construction
Date: 2026-08-10
Status: P0 Complete / P1-A Open
Companion plan: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-plan.md`
Companion evidence: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-evidence.md`
Companion benchmark: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-benchmark.md`

## Purpose

This file records the current source-backed identity/range/occurrence/collision boundary before Child 01 production implementation. It separates the measured target finding from the DRAFT architecture proposal; P1-A is documentation-only, while production implementation remains closed until the ordered slice gates pass.

## Freshness / Refresh Rules

- The Anvien graph was refreshed with `anvien analyze --force --json` at the current implementation HEAD before file-detail and impact queries.
- All file-detail rows below report `stale=false` and `changedSinceAnalyze=false`.
- Target validation is not run in P0; target source, graph, and pre-existing worktree state are preserve-only.
- Future slices must refresh only rows affected by accepted evidence or a completed slice and append a status-refresh row.

## Scope

Target scope:

- declaration/symbol identity and lexical-owner/meaning inputs;
- full range and selection-range facts;
- declaration occurrence conservation through the production ScopeIR-to-graph path;
- the proven graph node collision/replacement boundary;
- deterministic projection and endpoint integrity;
- bounded `time`/`now` target oracle at P1-E.

Out of scope:

- Graph JSON/Ladybug persistence and readers;
- TypeScript binding-pattern extraction, export semantics, module/re-export, ambient/external, and graph-health semantics;
- relationship identity/merge policy unless a later Child 01 impact proves it is required;
- scanner repair, target-source edits, and target analyze in P0.

## Relationship / Impact Evidence

All counts are fresh at implementation HEAD `7042dda8bfc02ee42b40c3a1c0ede89138439481`. Related-file counts are file counts, not symbol counts.

| Unit / File / Surface | File Detail Evidence | Related File Count | Relationship Summary | Impact Note |
|-----------------------|----------------------|--------------------:|----------------------|-------------|
| `internal/resolution/indexes.go` | `E0-P0A-FD1` | 46 | production workspace/index and graph-ID construction | high file risk; `graphIDForDef`/`buildWorkspace` CRITICAL |
| `internal/scopeir/range.go` | `E0-P0A-FD2` | 227 | shared four-coordinate range shape | medium file risk; `Range` CRITICAL |
| `internal/scopeir/definition_index.go` | `E0-P0A-FD3` | 225 | standalone definition-index utility | medium file risk; `BuildDefinitionIndex` LOW, test-only |
| `internal/graph/types.go` | `E0-P0A-FD4` | 238 | graph storage, duplicate-ID mutation, index rebuild | high file risk; `Graph.AddNode` CRITICAL |
| `internal/resolution/emit.go` | `E0-P0A-FD5` | 42 | definition projection, endpoints, separate relationship merge | high file risk; `emitNode`/`emitDefinitionNodes` CRITICAL |
| `internal/scopeir/facts.go` | `E0-P0A-FD6` | 231 | `DefinitionFact`/`ScopeFact` shared contracts | medium file risk; `DefinitionFact` CRITICAL |
| `internal/scopeir/ir.go` | `E0-P0A-FD7` | 229 | owned copies and deterministic normalization | high file risk; inspect-only |
| `internal/scopeir/sort_keys.go` | `E0-P0A-FD8` | 226 | definition/range sort keys | high file risk; inspect-only |
| `internal/providers/tsjs/definitions.go` | `E0-P0A-FD9` | 16 | TS definition construct range, scope binding, visibility input | high file risk; exact functions conditional P1-B/P1-C |
| `internal/providers/tsjs/nodes.go` | `E0-P0A-FD10` | 20 | tree-sitter position conversion | high file risk; `nodeRange` MEDIUM |
| `internal/providers/tsjs/scopes.go` | `E0-P0A-FD11` | 17 | lexical scope candidates and containment | high file risk; inspect-only unless P1-B proves need |
| `internal/providers/tsjs/extract.go` | `E0-P0A-FD12` | 23 | provider orchestration into ScopeIR | high file risk; preserve orchestration |
| `internal/resolution/resolve.go` | `E0-P0A-FD13` | 50 | binding/result orchestration and emission order | high file risk; preserve unless a slice proves plumbing need |
| `internal/resolution/types.go` | `E0-P0A-FD14` | 31 | resolution result/metrics contracts | high file risk; inspect-only |
| `internal/analyze/analyze.go` | `E0-P0A-FD15` | 182 | normal analyze phase orchestration | high file risk; preserve; no P0 edit |

Exact symbol impacts (`--direction upstream --include-tests`) are recorded in `E0-P0A-IMPACT1..E0-P0A-IMPACT2`. CRITICAL/HIGH is a blast-radius warning, not a prohibition.

## Status Rules

| Status | Meaning | Allowed next action |
|--------|---------|---------------------|
| `correct` | Current behavior satisfies the scoped requirement or a preserve-only sibling contract. | Preserve and validate. |
| `partial` | A required fact exists but is incomplete or not projected. | Change only the proven gap after the slice gate. |
| `wrong` | Current behavior contradicts a measured requirement. | Correct only at the evidence-backed owner. |
| `missing` | Required behavior or fact does not exist. | Implement only after owner evidence. |
| `unbound` | A fact exists but is not connected to its real flow or contract. | Bind only at the proven boundary. |
| `fake-or-stub` | Placeholder or fallback behavior is presented as real. | Remove or replace truthfully. |
| `blocked` | Required authority or acceptance evidence is not yet closed. | Do not implement. |

## Current Status Matrix

| Unit | Current State | Required State | Status | Relationship Count | Evidence | Next Plan Decision |
|------|---------------|----------------|--------|--------------------|----------|--------------------|
| problem/acceptance authority | bounded reports prove four ScopeIR occurrences become two graph nodes; report repair design remains DRAFT | use only measured `2/4 -> 4/4` finding and source-backed invariants | correct | N/A | `E0-P0A-REPORT1`, `E0-P0A-VERIFY1` | preserve; ratify mechanics only in P1-A |
| Anvien graph baseline | current self-graph is fresh at `7042dda8`; post-ledger refresh is `1,557/676/0`, `85,114` nodes, `123,982` relationships; pre-ledger baseline was `85,110/123,978` | all graph commands use this fresh index; the `+4/+4` document delta is not production behavior | correct | N/A | `E0-P0A-GRAPH1` | preserve and refresh after implementation |
| declaration contract (`DefinitionFact`) | full range, label, owner ID, and optional visibility exist; selection range, lexical-scope field, and explicit meaning lane do not | source-backed full/selection position plus identity inputs without inventing a topology | partial | 231 | `E0-P0A-SOURCE6`, `E0-P0A-FD6`, `E0-P0A-IMPACT1` | P1-B contract and exact provider/IR owner only |
| range semantics | four coordinates are present; TSJS uses one-based lines/byte-like columns while position lookup includes the end coordinate; encoding/interval contract is undocumented | explicit coordinate/encoding/interval semantics and a provider-supported selection range | partial | 227 | `E0-P0A-SOURCE2`, `E0-P0A-FD2`, `E0-P0A-FD10`, `E0-P0A-IMPACT1` | P1-B; no guessed conversion |
| lexical owner / meaning | ScopeIR scopes and bindings retain lexical membership; graph identity/projection does not consume scope and only `Label` acts as a meaning-like lane | distinct lexical owners and provider-distinguished meanings remain distinguishable | partial | 46 identity / 17 TS scope | `E0-P0A-SOURCE1`, `E0-P0A-SOURCE7`, `E0-P0A-FD1`, `E0-P0A-FD11` | P1-B/P1-C; no name-only or range-only shortcut |
| graph identity | `graphIDForDef` omits provider occurrence ID, range, scope, and meaning; `time`/`now` pairs map to one ID each | injective deterministic identity for distinct source occurrences and supported semantic lanes | wrong | 46 | `E0-P0A-SOURCE1`, `E0-P0A-FD1`, `E0-P0A-IMPACT1` | P1-C exact symbol owner |
| production occurrence path | provider/ScopeIR facts are appended and `defsByFile` retains each occurrence; `defsByID` is a separate provider-ID map and `DefinitionsIndexed` is not an occurrence denominator | every required input occurrence reaches a distinct accepted graph fact | wrong | 46 / 42 | `E0-P0A-SOURCE3`, `E0-P0A-SOURCE5`, `E0-P0A-SOURCE8` | P1-C uses `defsByFile`/source denominator; preserve `BuildDefinitionIndex` |
| graph collision/mutation | duplicate `Graph.AddNode` IDs replace payloads; `Graph.init` does not reject pre-existing duplicate IDs | conflicting canonical facts fail clearly or follow an explicit proven enrichment rule | wrong | 238 | `E0-P0A-SOURCE4`, `E0-P0A-FD4`, `E0-P0A-IMPACT2` | P1-D classify all legitimate callers before editing |
| definition projection/endpoints | `emitDefinitionNodes` attempts one node/`DEFINES` edge per `defsByFile` entry, but emits only line properties and colliding endpoints; relationship merge is separate | corrected identity/range facts reach nodes and every affected endpoint exists | partial | 42 | `E0-P0A-SOURCE8`, `E0-P0A-FD5`, `E0-P0A-IMPACT2` | inspect in P1-C/P1-E; edit only if proven |
| standalone definition index | first-writer duplicate behavior is covered by legacy tests but has zero production callers | retain its existing test contract; do not mislabel it as runtime loss owner | correct | 225 | `E0-P0A-SOURCE3`, `E0-P0A-FD3`, `E0-P0A-IMPACT2` | preserve-only; do not touch in Child 01 |
| analyze/resolution orchestration | direct provider → ScopeIR → workspace → bound resolution → definition emission path is source-proven; no hidden alternate runtime found | use normal production path; change orchestration only with new evidence | correct | 182 / 50 | `E0-P0A-SOURCE5`, `E0-P0A-FD13`, `E0-P0A-FD15` | preserve-only; no P0 edit; inspect each slice |
| target boundary | target source and worktree contain pre-existing user state; oracle file hash and graph snapshot are known | preserve all 13 entries and source; validate `4/4` only in P1-E | correct | external repository | `E0-P0A-BOUNDARY1`, `E0-P0A-BOUNDARY2` | preserve-only; no P0 target analyze or writes |

## Status Refresh Log

| Refresh | Date | Repo Basis | Changed Scope | Status Changes | Evidence | Next Phase Update |
|---------|------|-----------|---------------|----------------|----------|-------------------|
| R0 | 2026-08-10 | correction-closed HEAD `7042dda8`; prior bounded reports | Child 01 identity/range candidates | broad retained candidates pending refresh | `E0-P0A-REPORT1`, `E0-P0A-VERIFY1` | refresh all candidate owners before implementation |
| R1 | 2026-08-10 | fresh Anvien analyze + source/impact/target audits at `7042dda8` | Child 01 P0 source boundary | `BuildDefinitionIndex` causal candidate -> preserve-only; graph identity/collision -> confirmed wrong; range/IR inputs -> partial; orchestration -> preserve-only | `E0-P0A-GRAPH1`, `E0-P0A-SOURCE1..E0-P0A-SOURCE9`, `E0-P0A-FD1..E0-P0A-FD15`, `E0-P0A-IMPACT1..E0-P0A-IMPACT2`, `E0-P0A-BOUNDARY2` | P1-A remains gated by Supervisor; P1-B/P1-C/P1-D work steps narrowed to exact owners |
| R2 | 2026-08-10 | independent Supervisor report at `7042dda8` | P0-A acceptance and next-slice opening | P0 source/impact/status boundary -> accepted; P1-A documentation contract slice -> open; production implementation remains closed until P1-A and later implementation gates | `E0-P0A-REVIEW1` | P1-A may reconcile the source-backed contract; P1-B/P1-C/P1-D/P1-E remain ordered and closed |

## Phase Touch Map

Related files are not automatically editable. P0 touch mode is inspect-only or preserve-only for every production surface.

| Unit / File / Surface | Plan-Relevant Relationship File | Relationship to Target | Plan Item | Touch Mode | Evidence | Constraint |
|-----------------------|---------------------------------|------------------------|-----------|------------|----------|------------|
| graph-accuracy contract | Child 01 plan and ledgers | authority | P1-A | edit documentation | `E0-P0A-REPORT1`, `E0-P0A-VERIFY1` | exclude DRAFT implementation mechanics |
| `DefinitionFact`/`ScopeFact` | `internal/scopeir/facts.go` | shared IR contract | P1-B/P1-C | inspect-only until contract gate | `E0-P0A-FD6`, `E0-P0A-IMPACT1` | no mandatory field/topology name before source proof |
| range shape/semantics | `internal/scopeir/range.go`, `internal/providers/tsjs/nodes.go` | shared/provider position input | P1-B | inspect-only | `E0-P0A-FD2`, `E0-P0A-FD10`, `E0-P0A-SOURCE2` | no guessed encoding or inclusive/exclusive conversion |
| TS definition constructor | `internal/providers/tsjs/definitions.go` | construct range/scope/visibility input | P1-B/P1-C | inspect-only; edit only exact function after impact refresh | `E0-P0A-FD9`, `E0-P0A-SOURCE6` | preserve binding-pattern branch for Child 03 |
| TS scope builder | `internal/providers/tsjs/scopes.go` | lexical owner source | P1-B/P1-C | inspect-only | `E0-P0A-FD11`, `E0-P0A-SOURCE7` | preserve containment semantics unless P1 proves gap |
| normalization/sort | `internal/scopeir/ir.go`, `internal/scopeir/sort_keys.go` | deterministic occurrence ordering | P1-C/P1-E | inspect-only | `E0-P0A-FD7`, `E0-P0A-FD8` | no deduplication by name or ID without contract |
| graph identity/index | `internal/resolution/indexes.go` | `DefinitionFact` -> `defRef.GraphID` | P1-C | inspect-only; exact `graphIDForDef`/`buildWorkspace` candidate | `E0-P0A-FD1`, `E0-P0A-IMPACT1` | CRITICAL warning; no broad resolver rewrite |
| graph mutation | `internal/graph/types.go` | duplicate canonical node handling | P1-D | inspect-only; exact `Graph.AddNode` candidate | `E0-P0A-FD4`, `E0-P0A-IMPACT2` | classify 11 production caller families and enrichment first |
| node projection | `internal/resolution/emit.go` | identity/range -> graph node/endpoints | P1-C/P1-E | inspect/validate; edit only if impact proves need | `E0-P0A-FD5`, `E0-P0A-SOURCE8` | `emitRelationship` is separate and preserve-only |
| standalone `DefinitionIndex` | `internal/scopeir/definition_index.go` | legacy test utility | P1-C/P1-D | preserve-only / validate-only | `E0-P0A-FD3`, `E0-P0A-IMPACT2` | do not claim it is production loss owner |
| resolution/analyze orchestration | `internal/resolution/resolve.go`, `internal/resolution/types.go`, `internal/analyze/analyze.go` | normal command flow | all P1 slices | preserve-only / inspect-only | `E0-P0A-FD13..E0-P0A-FD15`, `E0-P0A-SOURCE5` | no orchestration change absent new proof |
| target source/graph/worktree | `E:\cheapapp.org` | bounded oracle and external boundary | P1-E | preserve-only / validate-only | `E0-P0A-BOUNDARY1`, `E0-P0A-BOUNDARY2` | no source copy, manual graph edit, or P0 analyze |

## Detailed Findings

### Identity, lexical owner, and meaning

Current state:

The source-backed ScopeIR path retains distinct `DefinitionFact.ID` values and lexical scope bindings for the four bounded declarations, but `graphIDForDef` recomputes a name/file-oriented ID without scope, range, provider occurrence ID, or an explicit meaning lane. The graph therefore maps both `time` facts to one ID and both `now` facts to one ID.

Required state:

```text
source occurrence + lexical owner + provider-supported meaning
-> deterministic injective graph identity
-> no name-only/range-only collapse
```

Evidence: `E0-P0A-SOURCE1`, `E0-P0A-SOURCE6`, `E0-P0A-SOURCE7`, `E0-P0A-BOUNDARY2`.

Classification: `wrong` at graph identity; `partial` at the shared IR input boundary.

Allowed next action: P1-A ratifies only source-backed invariants, then P1-B/P1-C select exact owners after their own impact gates.

Forbidden next action: import the problem report's DRAFT two-tier topology, add only line/range to the old ID, use random/absolute-path identity, or fall back to a global name.

### Ranges and selection range

Current state:

`Range` has four coordinates, but the provider coordinate conversion and position-index boundary behavior are not declared as one contract, and `DefinitionFact` has no selection-range field. Definition graph nodes drop columns entirely.

Required state:

```text
provider position facts -> explicit coordinate/interval semantics
                       -> full declaration range + identifier selection range
                       -> stable identity/projection inputs
```

Evidence: `E0-P0A-SOURCE2`, `E0-P0A-SOURCE6`, `E0-P0A-FD2`, `E0-P0A-FD10`.

Classification: `partial`.

Allowed next action: P1-B must establish the actual provider encoding and nearest owner before editing a shared contract.

Forbidden next action: silently convert byte columns to UTF-16, assume half-open/inclusive semantics, or make range alone the identity key.

### Occurrence conservation and collision

Current state:

The production `defsByFile` slice preserves each incoming definition occurrence, while `Graph.AddNode` last-writer replacement removes a distinct payload when graph IDs collide. `BuildDefinitionIndex` also skips duplicate provider IDs but is not called by production analyze and remains a legacy test contract.

Required state:

```text
input occurrence denominator -> accepted graph occurrences
                         100% with zero unexplained drops
```

Evidence: `E0-P0A-SOURCE3`, `E0-P0A-SOURCE4`, `E0-P0A-SOURCE5`, `E0-P0A-SOURCE8`, `E0-P0A-IMPACT1..E0-P0A-IMPACT2`.

Classification: `wrong` at graph mutation; `partial` at the identity/index boundary; standalone definition index `correct/preserve-only` relative to its existing tests.

Allowed next action: P1-C records the source/`defsByFile` denominator and P1-D classifies legitimate `AddNode` enrichment/reinsertion before changing canonical mutation behavior.

Forbidden next action: blanket-change every `AddNode` caller, merge relationship deduplication into node identity, or use output node counts alone as conservation proof.

### Projection and endpoint integrity

Current state:

`emitDefinitionNodes` emits a definition node and `DEFINES` edge per `defsByFile` entry, but it does not project columns, selection range, provider ID, or lexical scope. A colliding endpoint can still exist while pointing at the wrong surviving payload, so endpoint existence alone is insufficient.

Required state:

```text
corrected identity/range facts -> graph node properties
                             -> existing endpoints with no silent loss
```

Evidence: `E0-P0A-SOURCE8`, `E0-P0A-FD5`, `E0-P0A-IMPACT2`.

Classification: `partial`.

Allowed next action: inspect in P1-C/P1-E; edit `emit.go` only if fresh impact proves projection or conflict plumbing is required.

Forbidden next action: treat `emitRelationship`'s semantic merge as declaration conservation or broaden Child 01 into persistence/readers.

### Orchestration and target boundary

Current state:

The normal production flow is source-proven and has no hidden definition-index stage. The target has a dirty pre-existing worktree and a pinned source/graph baseline; no target analyze was run in P0.

Classification: orchestration `correct/preserve-only`; target boundary `correct/preserve-only`.

Allowed next action: use the normal runtime only after P1-B through P1-D gates; run the target oracle in P1-E with before/after preservation evidence.

Forbidden next action: copy target source into Anvien, manually edit target graph, or treat pre-existing target files as this campaign's artifacts.

## Next Phase Status Decisions

| Plan Item | Actual Status Finding | Required Status / Next-Action Update |
|-----------|-----------------------|--------------------------------------|
| P1-A | measured invariant is accepted; source owners are now explicit; DRAFT architecture remains non-authoritative | open the documentation-only contract reconciliation; ratify source-backed rules and ownership, not a predetermined topology |
| P1-B | range/selection inputs are partial and shared `Range` has CRITICAL blast radius | narrow to `DefinitionFact`/range/provider position owners; define encoding and interval semantics before any shared edit; preserve binding-pattern behavior for Child 03 |
| P1-C | graph identity is wrong; `defsByFile` is the production occurrence denominator; `BuildDefinitionIndex` is not causal | edit only `graphIDForDef`/`buildWorkspace` and any projection owner proven necessary; do not edit the standalone index or broad providers without a plan refresh |
| P1-D | `Graph.AddNode` is the proven silent replacement owner with 11 production caller families | classify canonical insertion, enrichment, and reinsertion operations; change only the conflict owner and required plumbing |
| P1-E | target boundary is captured but target analyze is intentionally not run in P0 | retain the exact 13-entry pre-state and hashes; validate `time`/`now` `4/4` only after implementation slices are accepted |

## Implementation Gate

- [x] Target scope is listed in the Current Status Matrix.
- [x] Each target unit has a current implementation-HEAD status.
- [x] Each status has exact evidence IDs.
- [x] Each candidate target file has fresh file-detail evidence and a file count.
- [x] Phase Touch Map lists plan-relevant relationship files and touch modes.
- [x] Correct and preserve-only surfaces are explicit.
- [x] Partial/wrong surfaces have exact next actions and forbidden fallbacks.
- [x] Target/pre-existing boundary is recorded.
- [x] Next-phase assumptions, next actions, and work steps are refreshed from P0 evidence.
- [x] Status Refresh Log has R0 and fresh R1 rows.
- [x] P0 Supervisor review passes.

## Final P0 Decision

- [ ] P0 actual-status incomplete. Implementation is blocked.
- [ ] P0 complete. Next phase can proceed unchanged.
- [x] P0 complete. Next phase status, next action, or work steps must be updated before implementation.
- [ ] P0 complete. Target scope is preserve-only.
- [ ] P0 complete. Implementation is blocked by missing authority or evidence.

Decision note:

The source, fresh graph, file-detail, impact, source-flow, and target-boundary evidence closes the P0 classification. Supervisor report `rp_supervisor_260810_095511_by_gpt-5-codex_child01_p0a.md` is PASS. P1-A is now the only open next slice; it is documentation-only and must ratify source-backed contract rules before P1-B production work opens. No production or test edit is authorized outside the ordered slice gates.
