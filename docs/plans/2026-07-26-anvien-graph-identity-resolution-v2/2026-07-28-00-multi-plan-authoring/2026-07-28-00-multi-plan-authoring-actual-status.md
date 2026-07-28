# Anvien Graph Identity Resolution v2 Multi-Plan Authoring Actual Status

Title: Anvien Graph Identity Resolution v2 Multi-Plan Authoring
Date: 2026-07-28
Status: P0 Complete / execution not authorized
Companion plan: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/2026-07-28-00-multi-plan-authoring/2026-07-28-00-multi-plan-authoring-plan.md`
Companion evidence: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/2026-07-28-00-multi-plan-authoring/2026-07-28-00-multi-plan-authoring-evidence.md`
Companion benchmark: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/2026-07-28-00-multi-plan-authoring/2026-07-28-00-multi-plan-authoring-benchmark.md`

## Purpose

This file records the real current state before implementation.

Implementation must not start until the target scope has a completed status row, evidence IDs, and a downstream plan decision.

This file does not replace `evidence.md`. It classifies current state from evidence.

Use exact evidence IDs from `evidence.md`, such as `E0-P0A-SRC1`, not broad section IDs such as `E0` or `E1`.

## Freshness / Refresh Rules

This actual-status file is a living current-state record, not a one-time P0 snapshot.

P0 records the baseline before implementation. After implementation begins, keep the Current Status Matrix updated so the next agent can trust it as the latest repo reality.

Update this file:

- after each completed implementation slice;
- before starting the next phase if repo state changed;
- whenever evidence changes a current-state classification;
- whenever the next phase's status assumptions, next action, or work steps need updating because reality differs from the previous status.

When refreshing status:

- update only the rows affected by the completed work or new evidence;
- use explicit transitions such as `missing -> correct`, `partial -> correct`, `fake-or-stub -> removed`, or `unbound -> bound-correct`;
- append a Status Refresh Log row instead of deleting history;
- keep detailed proof in `evidence.md`; store only classifications, evidence IDs, touch mode, and plan consequences here.

## Scope

Target scope:

- The legacy five-artifact planning set under `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/`.
- The exact P1-P7 phase and 98-slice inventory that a later authorized execution will transform.
- The future campaign roadmap and seven complete implementation child plan sets.
- The ownership, traceability, dependency, ledger, P0, and Pn contracts required for the multi-plan campaign.
- The conditional authority cutover from the legacy giant plan to the verified roadmap/child set.

Out of scope:

- Production source, tests, runtime configuration, generated graph/index output, build, Docker, browser, and Playwright behavior.
- Any implementation of the graph-identity/resolution fixes.
- Any write, copy, move, staging artifact, report, fixture, or temporary material in `E:\cheapapp.org`.
- Any plan or temporary directory at the Anvien repository root.
- Re-auditing or redesigning the legacy plan's technical decisions.

## Relationship / Impact Evidence

For each target file, prefer:

```text
anvien file-detail <path> --repo <repo> --json
```

Record how many files the target is related to before deciding touch mode. A file with many relationships may still be editable, but the plan must narrow the exact phase, touch mode, and validation needed.

| Unit / File / Surface | File Detail Evidence | Related File Count | Relationship Summary | Impact Note |
|-----------------------|----------------------|--------------------|----------------------|-------------|
| Legacy giant plan | `E0-P0A-FD1`, `E0-P0A-QUERY1` | N/A — docs file is not indexed | No graph relationship inventory exists for this Markdown file; disk hash, links, structure, and Git diff are the valid boundary | low code impact; high planning-authority risk |
| Legacy evidence/benchmark/actual-status/matrix | `E0-P0A-SRC2` | N/A — documentation artifacts | Companion and auxiliary ownership is established from disk inventory and explicit plan references | low code impact; medium ledger-ownership risk |
| Future roadmap | `E0-P0A-SRC2` | 0 existing files | Missing output; no relationship can exist before creation | low code impact; high authority-index risk |
| Future seven child sets | `E0-P0A-SRC2`, `E0-P0A-TOTAL1` | 0 existing child files | Missing outputs; expected 7 folders and 28 standard files | low code impact; high completeness risk |
| `E:\cheapapp.org` boundary | `E0-P0A-BOUNDARY1`, `E0-P0A-SCOPE1` | out of scope | Separate indexed repository; no authoring operation may target it | critical scope boundary; do-not-touch |

## Status Rules

| Status | Meaning | Allowed next action |
|--------|---------|---------------------|
| `correct` | Already behaves as required. | Preserve. Add evidence or tests only if needed. |
| `partial` | Some required behavior exists, but gaps remain. | Change only the missing parts. Preserve correct parts. |
| `wrong` | Current behavior, source, or contract is incorrect. | Replace with required behavior. Record the exact reason. |
| `missing` | Required behavior, source, or contract does not exist. | Implement the missing piece only. |
| `unbound` | Surface exists but is not wired to the real source, flow, or contract. | Bind to the real source only. Preserve approved surface. |
| `fake-or-stub` | Prototype, demo, mock, fallback, or placeholder data is being used as real behavior. | Remove fake behavior or replace it with an approved truthful state. |
| `blocked` | Source, authority, contract, or required evidence is unclear. | Stop. Do not implement until resolved. |

## Current Status Matrix

| Unit | Current State | Required State | Status | Relationship Count | Evidence | Next Plan Decision |
|------|---------------|----------------|--------|--------------------|----------|--------------------|
| Legacy planning authority | Giant plan exists, P0 is complete, and implementation is not authorized | Remain active until the verified multi-plan set passes Supervisor | correct | N/A — docs not indexed | `E0-P0A-AUTH1`, `E0-P0A-FD1` | preserve through P3-A; metadata/pointer edit only in P3-B after PASS |
| Legacy plan sizing | One 5,467-line plan owns seven implementation phases and closure | Seven complete implementation children coordinated by one roadmap | wrong | N/A — docs not indexed | `E0-P0A-SRC1`, `E0-P0A-STRUCT1` | split through P1/P2 without rewriting source semantics |
| Phase-to-child boundary | Seven source implementation phases are measurable | One source phase per numbered child, P1->01 through P7->07 | correct | 7 phases | `E0-P0A-STRUCT1`, `E0-P0A-DECISION1` | preserve exact boundary in P1-A/P1-B |
| Source slice inventory | All source IDs and per-phase counts are known | 98 source slices mapped exactly once in original order | correct | 98 slices | `E0-P0A-SLICE1`, `E0-P0A-SLICE2`, `E0-P0A-SLICE3`, `E0-P0A-SLICE4`, `E0-P0A-SLICE5`, `E0-P0A-SLICE6`, `E0-P0A-SLICE7`, `E0-P0A-TOTAL1` | freeze in P1-A; validate again in P3-A |
| Legacy P8 closure | Three closure items exist only in the giant plan | Each of seven children has tailored Pn-A/Pn-B/Pn-C; no child 08 | partial | 3 source closure items / 0 child closure sets | `E0-P0A-TOTAL1`, `E0-P0A-DECISION1` | distribute during P2-A through P2-G |
| Campaign roadmap | No roadmap exists | One roadmap owns order, status, dependencies, handoffs, and active authority index | missing | 0 files | `E0-P0A-SRC2`, `E0-P0A-DECISION2` | create in P1-B only |
| Child 01 plan set | Does not exist | Four-file complete child for legacy P1 with P0/local P1/Pn | missing | 0 files | `E0-P0A-SRC2`, `E0-P0A-SLICE1` | create in P2-A |
| Child 02 plan set | Does not exist | Four-file complete child for legacy P2 with P0/local P1/Pn | missing | 0 files | `E0-P0A-SRC2`, `E0-P0A-SLICE2` | create in P2-B |
| Child 03 plan set | Does not exist | Four-file complete child for legacy P3 with P0/local P1/Pn | missing | 0 files | `E0-P0A-SRC2`, `E0-P0A-SLICE3` | create in P2-C |
| Child 04 plan set | Does not exist | Four-file complete child for legacy P4 with P0/local P1/Pn | missing | 0 files | `E0-P0A-SRC2`, `E0-P0A-SLICE4` | create in P2-D |
| Child 05 plan set | Does not exist | Four-file complete child for legacy P5 with P0/local P1/Pn | missing | 0 files | `E0-P0A-SRC2`, `E0-P0A-SLICE5` | create in P2-E |
| Child 06 plan set | Does not exist | Four-file complete child for legacy P6 with P0/local P1/Pn | missing | 0 files | `E0-P0A-SRC2`, `E0-P0A-SLICE6` | create in P2-F |
| Child 07 plan set | Does not exist | Four-file complete child for legacy P7 with P0/local P1/Pn | missing | 0 files | `E0-P0A-SRC2`, `E0-P0A-SLICE7` | create in P2-G |
| Child-independent ledgers | Evidence, benchmark, and current status are centralized in the legacy set | Each child owns a complete, phase-scoped ledger set and qualified cross-plan references | partial | 3 legacy ledgers / 0 child ledgers | `E0-P0A-SRC2`, `E0-P0A-DECISION2` | split ownership in each P2 slice; validate in P3-A |
| Reader matrix ownership | One matrix exists at campaign root but no child exists to own it | Child 02 is sole mutation owner; other children link inspect-only | unbound | 1 file | `E0-P0A-SRC2`, `E0-P0A-OWNER1` | bind ownership in P2-B; do not duplicate/move by default |
| Multi-plan authority | Candidate set does not exist | Roadmap becomes sole active campaign index after PASS | missing | 0 candidate files | `E0-P0A-SRC2`, `E0-P0A-AUTH1` | create/verify P1-P3; cut over only in P3-B |
| Fake/stub planning output | No roadmap or child placeholder is being treated as implemented | No placeholder or draft may be treated as active authority | correct | 0 fake child outputs | `E0-P0A-SRC2` | preserve; structural check must reject placeholder tokens |
| Target boundary | Target is a separate repository and not an authoring location | No target write/copy/move/read-as-source for this split | correct | out of scope | `E0-P0A-BOUNDARY1`, `E0-P0A-SCOPE1` | do-not-touch in every phase |
| Execution authorization | User authorized writing this guide plan, not executing the split | Explicit later user instruction before P1-A | blocked | N/A | `E0-P0A-USER1` | stop after plan creation and review |

## Status Refresh Log

| Refresh | Date | Repo Basis | Changed Scope | Status Changes | Evidence | Next Phase Update |
|---------|------|------------|----------------|----------------|----------|-------------------|
| R0 | 2026-07-28 | `master` at `68811c1643b604573e70551c7d4becb46e6ebbd8`; worktree clean before this authoring plan set; fresh Anvien index at the same commit | legacy five-artifact source set, phase/slice inventory, sample structure, future roadmap/child gap, target boundary | initial classification; P0 complete; execution blocked pending explicit instruction | `E0-P0A-GIT1`, `E0-P0A-GRAPH1`, `E0-P0A-STATUS1`, `E0-P0A-SRC1`, `E0-P0A-SRC2`, `E0-P0A-TOTAL1`, `E0-P0A-DECISION1` | keep P1-A blocked until explicit execution authority; do not create roadmap or children merely because this plan exists |

## Phase Touch Map

Use this map to prevent accidental edits. A related file is not automatically editable.

`Plan-Relevant Relationship File` lists only a relationship file that can directly affect or be affected by the planned phase or slice. Do not copy the full `file-detail` relationship inventory into this map. Include only files whose relationship can affect the phase/slice decision, touch mode, or validation.

| Unit / File / Surface | Plan-Relevant Relationship File | Relationship to Target | Plan Item | Touch Mode | Evidence | Constraint |
|-----------------------|---------------------------------|------------------------|-----------|------------|----------|------------|
| This authoring plan set | four files in `2026-07-28-00-multi-plan-authoring/` | control, evidence, metrics, and status for the split | P0-Pn | edit | `E0-P0A-TEMPLATE1` | each file retains its own primary ledger/control responsibility |
| Legacy giant plan | `2026-07-26-anvien-graph-identity-resolution-v2-plan.md` | source authority and transformation input | P1-A-P3-A | preserve-only | `E0-P0A-SRC1`, `E0-P0A-AUTH1` | no body edit before/after cutover |
| Legacy giant plan metadata/pointer | same legacy plan | authority marker | P3-B | edit | `E0-P0A-AUTH1` | edit only after unconditional Supervisor PASS |
| Legacy evidence ledger | `2026-07-26-anvien-graph-identity-resolution-v2-evidence.md` | source evidence/provenance | P1-P3 | inspect-only | `E0-P0A-SRC2` | do not rewrite historical evidence |
| Legacy benchmark ledger | `2026-07-26-anvien-graph-identity-resolution-v2-benchmark.md` | source metrics/provenance | P1-P3 | inspect-only | `E0-P0A-SRC2` | phase-specific metrics move to child ownership without deleting history |
| Legacy actual-status ledger | `2026-07-26-anvien-graph-identity-resolution-v2-actual-status.md` | source current-state baseline | P1-P3 | inspect-only | `E0-P0A-SRC2` | children receive refreshed, scoped status rather than a blind copy |
| Reader matrix | `index-reader-matrix.md` | P2 reader/cutover auxiliary contract | P2-B | inspect-only during authoring; future child-02 mutation owner | `E0-P0A-OWNER1` | one file, one owner; no duplicate |
| Campaign roadmap | source plan and seven future children | coordination and active-authority index | P1-B/P2/P3 | edit | `E0-P0A-DECISION2` | no implementation slice bodies |
| Child 01 standard set | legacy P1 and scoped ledgers | identity plan authority | P2-A | edit | `E0-P0A-SLICE1` | exactly four standard files and 11 slices |
| Child 02 standard set | legacy P2, child-01 handoff, matrix | persistence/cutover plan authority | P2-B | edit | `E0-P0A-SLICE2`, `E0-P0A-OWNER1` | exactly four standard files and 42 slices |
| Child 03 standard set | legacy P3 and child-02 handoff | binding-pattern plan authority | P2-C | edit | `E0-P0A-SLICE3` | exactly four standard files and 17 slices |
| Child 04 standard set | legacy P4 and child-03 handoff | export-semantics plan authority | P2-D | edit | `E0-P0A-SLICE4` | exactly four standard files and 15 slices |
| Child 05 standard set | legacy P5 and child-04 handoff | module/re-export plan authority | P2-E | edit | `E0-P0A-SLICE5` | exactly four standard files and 4 slices |
| Child 06 standard set | legacy P6 and child-05 handoff | ambient/external diagnostics plan authority | P2-F | edit | `E0-P0A-SLICE6` | exactly four standard files and 6 slices |
| Child 07 standard set | legacy P7 and children 01-06 handoffs | campaign acceptance plan authority | P2-G | edit | `E0-P0A-SLICE7` | exactly four standard files and 3 slices; target remains do-not-touch during authoring |
| Planner templates | `.agents/skills/planner/templates/*.template.md` | structure authority | P1-P3 | inspect-only | `E0-P0A-TEMPLATE1` | do not edit templates for this campaign |
| Restaurant Manager sample | `G:\Restaurant_manager\DOCS\plans\2026-07-02-prototype-ux-overhaul` | structural precedent | P0/P1 | inspect-only | `E0-P0A-SAMPLE1`, `E0-P0A-SAMPLE2` | sample child count is not copied |
| Anvien production/tests/runtime/graph | repo source and generated surfaces | forbidden non-doc scope | all | do-not-touch | `E0-P0A-BOUNDARY1` | any such diff blocks the slice |
| `E:\cheapapp.org` | separate target repository | forbidden authoring location | all | do-not-touch | `E0-P0A-BOUNDARY1`, `E0-P0A-SCOPE1` | no write, copy, move, staging, report, fixture, or temp artifact |

## Detailed Findings

### Legacy Giant Plan

Current state:

The active plan is a complete but oversized 5,467-line document. It already contains seven ordered implementation phases, 98 granular slices, and a closure phase. Its P0 and technical content are source authority; the problem is execution and ledger scale, not absence of detailed slices.

Required state:

```text
Preserve the legacy plan as source/history while transforming each P1-P7 phase into one complete child plan. Do not change source semantics or delete the legacy artifact.
```

Evidence:

- `E0-P0A-SRC1`: exact path, SHA-256, byte size, and line counts.
- `E0-P0A-STRUCT1`: exact implementation/closure heading positions.
- `E0-P0A-TOTAL1`: exact total of 98 implementation slices.

Relationship and impact:

- Related file count: N/A; `file-detail` does not index this docs plan.
- Relationship summary: four companion/auxiliary source artifacts on disk plus future roadmap/children.
- Impact note: low code impact, high planning-authority risk.

Classification:

`wrong` for executable campaign size; `correct` as preserved source authority until cutover.

Allowed next action:

Freeze its hash, read it as inspect-only input, and split its phase-owned content through P1/P2 after explicit execution authority.

Forbidden next action:

Do not edit technical body, delete it, mark it superseded early, or implement from partial child output.

### Multi-Plan Campaign Outputs

Current state:

No roadmap or numbered implementation child plan set exists. The newly created `00-multi-plan-authoring` set only controls future conversion and is not an implementation child.

Required state:

```text
One roadmap plus seven numbered, four-file child plan sets. Each child contains P0, one local implementation phase holding its complete legacy phase slice set, and Pn-A/Pn-B/Pn-C.
```

Evidence:

- `E0-P0A-SRC2`: current campaign-root inventory has no roadmap/child folders.
- `E0-P0A-DECISION1`: seven-child result from the user's phase-to-child rule.
- `E0-P0A-SAMPLE2`: verified complete child lifecycle precedent.

Relationship and impact:

- Related file count: 0 existing roadmap/child files; 29 future campaign output files.
- Relationship summary: roadmap coordinates seven children; each child owns one source phase and four independent standard files.
- Impact note: high completeness and authority risk, no runtime impact.

Classification:

`missing`.

Allowed next action:

Create outputs only through the ordered P1/P2 authoring slices after explicit execution authority.

Forbidden next action:

Do not treat this guide plan as the roadmap, create an arbitrary child count, omit child P0/Pn, or execute a child before accepted cutover.

### Ledger And Auxiliary Ownership

Current state:

Legacy evidence, benchmark, and actual status cover the entire giant plan. `index-reader-matrix.md` exists once at campaign root but has no child owner because children do not yet exist.

Required state:

```text
Each child owns phase-scoped plan/evidence/benchmark/actual-status files. Cross-child evidence is slug-qualified. Child 02 is the sole mutation owner of the single reader matrix.
```

Evidence:

- `E0-P0A-SRC2`: current companion and matrix inventory.
- `E0-P0A-OWNER1`: matrix ownership decision derived from legacy P2 responsibility.

Relationship and impact:

- Related file count: four legacy companion/auxiliary files plus 21 future child ledger files.
- Relationship summary: source ledgers are inspect-only history; child ledgers become execution-local truth.
- Impact note: medium evidence-collision and stale-status risk.

Classification:

`partial` for ledgers and `unbound` for matrix ownership.

Allowed next action:

Author independent, populated child ledgers and bind matrix mutation ownership in P2-B.

Forbidden next action:

Do not duplicate mutable evidence, use unqualified cross-child IDs, or allow multiple matrix owners.

### Target Boundary

Current state:

`E:\cheapapp.org` is a separate indexed repository used by the eventual implementation campaign as an in-place validation subject. It has no role as storage or source material for this docs conversion.

Required state:

```text
All authoring artifacts remain in E:\Anvien. E:\cheapapp.org receives no write, copy, move, report, fixture, or temporary artifact during multi-plan authoring.
```

Evidence:

- `E0-P0A-REPOS1`: Anvien and target are separately registered repositories.
- `E0-P0A-BOUNDARY1`: repository/user boundary.
- `E0-P0A-SCOPE1`: P0 authoring inspection stayed outside target content.

Relationship and impact:

- Related file count: out of scope.
- Relationship summary: future validation subject only, never an authoring location.
- Impact note: critical do-not-touch boundary.

Classification:

`correct`.

Allowed next action:

Preserve the boundary text in roadmap/children; do not access target during authoring.

Forbidden next action:

Do not write, copy, move, stage, inspect as split source, or create any artifact in the target.

## Next Phase Status Decisions

| Plan Item | Actual Status Finding | Required Status / Next-Action Update |
|-----------|-----------------------|--------------------------------------|
| P1-A | Source hash, phase order, and 98-slice inventory are known; docs file-detail is unavailable | keep P1-A contract; require explicit execution authority and hash recheck; use disk/Markdown/Git evidence |
| P1-B | Roadmap is missing and seven-child inventory is known | keep P1-B; create exactly one coordination roadmap after P1-A |
| P2-A | Child 01 is missing; legacy P1 has 11 slices | keep; author one complete four-file child with 11 mappings |
| P2-B | Child 02 is missing; legacy P2 has 42 slices; matrix owner is unbound | keep; author complete child and bind sole matrix ownership |
| P2-C | Child 03 is missing; legacy P3 has 17 slices | keep; require child-02 handoff |
| P2-D | Child 04 is missing; legacy P4 has 15 slices | keep; require child-03 handoff |
| P2-E | Child 05 is missing; legacy P5 has four slices and a semantic vector | keep; preserve vector with owning child |
| P2-F | Child 06 is missing; legacy P6 has six slices and a status matrix | keep; preserve status contract with owning child |
| P2-G | Child 07 is missing; legacy P7 has three slices and depends on P1-P6 outcomes | keep; require six qualified handoffs and preserve target boundary |
| P3-A | No candidate exists yet | keep blocked until 1 roadmap, 7 children, 28 files, and 98 mappings exist |
| P3-B | Legacy is correctly active and no accepted replacement exists | keep legacy active; cut over only after unconditional Supervisor PASS |

## Implementation Gate

- [x] Target scope is listed in Current Status Matrix.
- [x] Each target unit has a status.
- [x] Each status has evidence IDs.
- [x] Each target file has relationship count evidence from `file-detail` when applicable; the docs exclusion and N/A result are recorded explicitly.
- [x] Phase Touch Map lists plan-relevant relationship files that can affect the current phase/slice.
- [x] Phase Touch Map defines touch mode for every plan-relevant relationship unit that may be affected.
- [x] Correct parts are marked preserve-only.
- [x] Partial, missing, wrong, unbound, and fake-or-stub parts have exact next actions; no fake-or-stub output was observed.
- [x] Blockers are recorded, including lack of execution authority.
- [x] Next phase status assumptions, next action, and work steps have been updated from this status file when needed.
- [x] Status Refresh Log has an R0 baseline row.
- [x] Implementation has not started, so no post-slice status transition is due.
- [x] No refreshed status has authorized P1; the next action remains blocked pending explicit user instruction.

## Final P0 Decision

Choose one:

- [ ] P0 actual-status incomplete. Implementation is blocked.
- [ ] P0 complete. Next phase can proceed unchanged.
- [ ] P0 complete. Next phase status, next action, or work steps must be updated before implementation.
- [ ] P0 complete. Target scope is preserve-only.
- [x] P0 complete. Implementation is blocked by missing authority or evidence.

Decision note:

P0 is complete for writing this guide plan. The source structure, counts, ownership boundaries, target prohibition, and future outputs are known. The current user request authorizes creation of this four-file authoring plan set only; it does not authorize P1-A or any roadmap/child-plan creation. A later explicit execution instruction must open P1-A, and the legacy source hash must be rechecked at that time.
