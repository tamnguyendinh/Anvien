# Anvien Graph Identity Contract and Strict Construction Actual Status

Title: Anvien Graph Identity Contract and Strict Construction

Date: 2026-07-28

Status: Draft / P0 Incomplete / Implementation Blocked

Companion plan: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-plan.md`

Companion evidence: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-evidence.md`

Companion benchmark: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-benchmark.md`

## Purpose

This file classifies the identity/range/occurrence/collision state before Child 01 production work. It does not convert the problem report's DRAFT design into implementation authority.

Implementation remains blocked until the candidate source owners have fresh file-detail and impact evidence and Supervisor accepts the P0 boundary.

## Freshness / Refresh Rules

- Refresh the graph before every graph-based status update.
- Replace retained candidate counts with fresh P0 file counts before a file becomes editable.
- Refresh only rows affected by accepted evidence or a completed slice.
- Append a Status Refresh Log row; do not preserve obsolete planning histories as active state.
- Keep command details in the evidence ledger and measured values in the benchmark ledger.
- Update the next slice's assumptions and work steps when a classification changes.

## Scope

Target scope:

- declaration and symbol identity;
- lexical owner and meaning inputs;
- range and selection-range inputs;
- declaration occurrence conservation;
- the proven silent collision/replacement/skip boundary;
- deterministic and integrity-valid identity construction;
- bounded `time`/`now` target oracle.

Out of scope:

- Graph JSON/Ladybug persistence and readers;
- binding, export, module/re-export, ambient/external, and graph-health semantics;
- relationship redesign or broad producer rewrite without new impact evidence;
- scanner repair and target-source changes.

## Relationship / Impact Evidence

These are candidate source surfaces from the latest production-source audit. They remain inspect-only until P0 refreshes their file-detail and symbol impact at the implementation HEAD.

| Unit / File / Surface | File Detail Evidence | Related File Count | Relationship Summary | Impact Note |
|-----------------------|----------------------|-------------------:|----------------------|-------------|
| `internal/resolution/indexes.go` | `E0-P0A-FD1` | 46 files | current graph identity construction and resolver indexes | CRITICAL candidate; exact identity symbol impact pending |
| `internal/scopeir/range.go` | `E0-P0A-FD2` | 227 files | shared range input | HIGH shared-contract candidate |
| `internal/scopeir/definition_index.go` | `E0-P0A-FD3` | 225 files | current definition occurrence indexing | HIGH occurrence-loss candidate |
| `internal/graph/types.go` | `E0-P0A-FD4` | 238 files | current node storage and duplicate-ID behavior | CRITICAL collision candidate |
| `internal/resolution/emit.go` | `E0-P0A-FD5` | 42 files | identity-to-graph projection and endpoints | HIGH validation candidate; edit only if impact proves need |

## Status Rules

| Status | Meaning | Allowed next action |
|--------|---------|---------------------|
| `correct` | Current behavior satisfies the Child requirement. | Preserve and validate. |
| `partial` | Some required fact exists, but acceptance is incomplete. | Change only the proven gap. |
| `wrong` | Current behavior conflicts with the measured requirement. | Correct in the evidence-backed owner. |
| `missing` | Required behavior or fact is absent. | Implement only after owner evidence. |
| `unbound` | A fact exists but is not connected to the real construction flow. | Bind only at the proven boundary. |
| `fake-or-stub` | Placeholder behavior is presented as real. | Remove or replace truthfully. |
| `blocked` | Required authority or current evidence is absent. | Do not implement. |

## Current Status Matrix

| Unit | Current State | Required State | Status | Relationship Count | Evidence | Next Plan Decision |
|------|---------------|----------------|--------|-------------------:|----------|--------------------|
| problem/acceptance authority | bounded report proves four ScopeIR occurrences become two graph nodes; repair design remains DRAFT | use measured finding and `4/4` oracle without importing unproven design | correct | 2 collision groups in bounded oracle | `E0-P0A-REPORT1`, `E0-P0A-VERIFY1` | preserve in P1-A |
| declaration/symbol identity | current identity omits sufficient lexical/source distinction for the bounded same-name cases | deterministic distinct declaration/symbol identities for different lexical owners/meanings | wrong | 46 candidate related files | `E0-P0A-SOURCE1`, `E0-P0A-FD1`; fresh symbol impact pending | P1-C after P1-B |
| range and selection range | current shared range facts do not yet prove the complete position contract needed by the Child | source-backed range and selection-range inputs with explicit semantics | partial | 227 candidate related files | `E0-P0A-SOURCE2`, `E0-P0A-FD2`; fresh impact pending | P1-B |
| occurrence conservation | a duplicate definition identity may be skipped before graph output | every relevant declaration occurrence remains represented | wrong | 225 candidate related files | `E0-P0A-SOURCE3`, `E0-P0A-FD3`; fresh impact pending | P1-C |
| graph collision behavior | duplicate node IDs can replace an earlier payload | no distinct canonical occurrence is silently replaced, skipped, or collapsed | wrong | 238 candidate related files | `E0-P0A-SOURCE4`, `E0-P0A-FD4`; fresh impact pending | P1-D |
| endpoint/integration projection | projection consumes the current identities; whether it requires code change is unproven | affected endpoints exist and identity facts reach normal graph output | blocked pending impact | 42 candidate related files | `E0-P0A-FD5`; current source/impact refresh pending | inspect in P1-C/P1-E; edit only if proven |
| analyze implementation | current runtime constructs the baseline graph; its phase details have not been accepted as an invariant | use or change only what the owning slice evidence requires | partial baseline | P0 source-flow inventory pending | `E0-P0A-SOURCE5` pending | inspect in every slice; never freeze by plan |
| target boundary | target source is an external validation input and normal output stays in the target repository | preserve target source/worktree and validate `4/4` in place | correct | external repository boundary | `E0-P0A-BOUNDARY1` | validate in P1-E |

## Status Refresh Log

| Refresh | Date | Repo Basis | Changed Scope | Status Changes | Evidence | Next Phase Update |
|---------|------|------------|---------------|----------------|----------|-------------------|
| R0 | 2026-08-10 | documentation correction worktree; production-source candidates retained from the last bounded audit | Child 01 only | old broad campaign state removed; identity defect remains wrong; implementation owner refresh and Supervisor pending | `E0-P0A-REPORT1`, `E0-P0A-VERIFY1`, `E0-P0A-FD1..E0-P0A-FD5` | keep P1-A through P1-E blocked until current P0 evidence closes |

## Phase Touch Map

`Plan-Relevant Relationship File` lists only relationships that can affect a Child 01 decision. Candidate files remain inspect-only until their active-slice impact gate passes.

| Unit / File / Surface | Plan-Relevant Relationship File | Relationship to Target | Plan Item | Touch Mode | Evidence | Constraint |
|-----------------------|---------------------------------|------------------------|-----------|------------|----------|------------|
| graph-accuracy contract | Child 01 plan and ledgers | authority | P1-A | edit documentation | `E0-P0A-REPORT1`, `E0-P0A-VERIFY1` | no production design without source evidence |
| range input candidate | `internal/scopeir/range.go` | shared source position | P1-B | inspect-only until fresh impact; then evidence-scoped edit | `E0-P0A-FD2` | do not guess encoding or persistence shape |
| identity candidate | `internal/resolution/indexes.go` | declaration-to-graph identity | P1-C | inspect-only until fresh impact; then evidence-scoped edit | `E0-P0A-FD1` | no name-only or range-only shortcut |
| occurrence candidate | `internal/scopeir/definition_index.go` | duplicate occurrence handling | P1-C | inspect-only until fresh impact; then evidence-scoped edit | `E0-P0A-FD3` | conserve every required occurrence |
| collision candidate | `internal/graph/types.go` | duplicate node behavior | P1-D | inspect-only until fresh impact; then evidence-scoped edit | `E0-P0A-FD4` | preserve legitimate proven enrichment |
| projection candidate | `internal/resolution/emit.go` | affected endpoint projection | P1-C/P1-E | inspect/validate; edit only if impact proves need | `E0-P0A-FD5` | no broad producer rewrite |
| `E:\cheapapp.org` | normal target graph output | bounded source oracle | P1-E | preserve source; validate only | `E0-P0A-BOUNDARY1` | no target fixture/report/debug writes |

## Detailed Findings

### Same-name identity and occurrence loss

Current state:

The bounded investigation observed four source declarations in ScopeIR and two surviving graph nodes. Current source evidence identifies insufficient graph identity and duplicate-ID replacement as the first graph-side divergence.

Required state:

```text
four distinct source declarations
-> four conserved declaration occurrences
-> lexical-owner/meaning-correct symbols
-> four represented target identities
```

Evidence:

- `E0-P0A-REPORT1`: original report and bounded `2/4` baseline.
- `E0-P0A-VERIFY1`: causal synthesis and bounded Supervisor verification.
- `E0-P0A-SOURCE1`, `E0-P0A-SOURCE3`, `E0-P0A-SOURCE4`: current identity, occurrence, and collision facts; refresh pending.

Relationship and impact:

- Related file counts: 46 identity-candidate files, 225 occurrence-candidate files, and 238 collision-candidate files.
- Impact note: HIGH/CRITICAL candidate blast radius; editable symbols are not authorized until fresh P0 impact.

Classification: `wrong`, with implementation blocked by current owner evidence.

Allowed next action: complete P0, ratify P1-A, then execute P1-B through P1-E in order.

Forbidden next action: treat a range-only patch, broad producer rewrite, or report DRAFT topology as accepted implementation.

### Range and analyze boundaries

Current state:

The range owner supplies partial source positions, but the exact full/selection range contract and the need for pipeline changes are not yet established at the implementation HEAD.

Required state:

```text
provider facts -> explicit position semantics -> identity inputs
normal analyze -> either integrity-valid corrected identities or a clear failure
```

Classification: range `partial`; analyze implementation `partial baseline`; both require P0/P1 evidence.

## Next Phase Status Decisions

| Plan Item | Actual Status Finding | Required Status / Next-Action Update |
|-----------|-----------------------|--------------------------------------|
| P1-A | measured invariant is known; current source-owner and Supervisor gates remain open | keep documentation slice blocked until P0 closes |
| P1-B | range/selection semantics are partial | select exact owner only from fresh source/file-detail/impact |
| P1-C | identity and occurrence conservation are wrong | implement only after P1-B and current impact; update plan before any extra producer |
| P1-D | silent replacement is wrong | classify legitimate same-ID operations before editing the proven owner |
| P1-E | integration outcome is not yet measured | validation only; failures return to P1-B/P1-C/P1-D |

## Implementation Gate

- [ ] Target scope is listed in Current Status Matrix.
- [ ] Each target unit has a current implementation-HEAD status.
- [ ] Each status has exact current evidence IDs.
- [ ] Each candidate target file has refreshed file-detail evidence and a file count.
- [ ] Exact editable symbols/files have fresh upstream impact evidence.
- [ ] Phase Touch Map reflects the final evidence-backed owner set.
- [x] Correct target and report-authority boundaries are preserve-only.
- [ ] Partial/wrong/blocked rows have current exact next actions.
- [ ] Next-slice assumptions and work steps have been refreshed from current evidence.
- [x] Status Refresh Log has an R0 correction row.
- [ ] P0 Supervisor review passes.

## Final P0 Decision

- [x] P0 actual-status incomplete. Implementation is blocked.
- [ ] P0 complete. Next phase can proceed unchanged.
- [ ] P0 complete. Next phase status, next action, or work steps must be updated before implementation.
- [ ] P0 complete. Target scope is preserve-only.
- [ ] P0 complete. Implementation is blocked by missing authority or evidence.

Decision note:

The bounded identity defect and target oracle are established, but current file-detail/impact/source-flow evidence and P0 Supervisor acceptance are still pending. No production slice is open.
