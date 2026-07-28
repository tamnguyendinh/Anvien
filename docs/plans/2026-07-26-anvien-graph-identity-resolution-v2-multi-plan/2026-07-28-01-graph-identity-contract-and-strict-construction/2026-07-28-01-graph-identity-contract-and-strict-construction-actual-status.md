# Graph Identity Contract and Strict Graph Construction Actual Status

Title: Graph Identity Contract and Strict Graph Construction
Date: 2026-07-28
Status: P0 Complete / implementation blocked
Companion plan: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-plan.md`
Companion evidence: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-evidence.md`
Companion benchmark: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-benchmark.md`
Roadmap: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-26-anvien-graph-identity-resolution-v2-roadmap.md`
Source phase: legacy `P1`

## Purpose

This file records the real current state before implementation.

Implementation must not start until the target scope has a completed status row, evidence IDs, and a downstream plan decision.

This file does not replace `evidence.md`. It classifies current state from evidence.

Use exact evidence IDs from `evidence.md`, such as `E0-P0A-SOURCE1`, not broad section IDs.

## Freshness / Refresh Rules

This actual-status file is a living current-state record, not a one-time P0 snapshot.

Update it after each completed local slice, before the next slice/phase if repo reality changed, whenever classifications change, and after campaign-authority input changes. Child 01 has no upstream implementation-child handoff. Append refresh rows; do not delete history. Update only stale next-action/work-step assumptions while preserving the accepted goal, scope, acceptance, and order.

## Scope

Target scope:

- Declaration/Symbol identity and strict graph construction.
- 11 local implementation slices mapped exactly from legacy P1.
- This child's independent evidence, benchmark, current-status, closure, and roadmap handoff.

Out of scope:

- No active-v2 cutover or reader migration; child 02 owns that work.
- No TypeScript binding/export/barrel/ambient implementation.
- No target source copy or target v2 artifact.
- Multi-plan authoring itself does not authorize production or target execution.

## Relationship / Impact Evidence

For each target file, prefer:

```text
anvien file-detail <path> --repo <repo> --json
```

Inherited counts below remain valid because no production path changed after the source P0. Refresh file-detail plus upstream impact immediately before the exact implementation slice edits a file.

| Unit / File / Surface | File Detail Evidence | Related File Count | Relationship Summary | Impact Note |
|-----------------------|----------------------|--------------------|----------------------|-------------|
| `internal/graph/types.go` | `E0-P0A-FD1` | 238 | central node/relationship storage and indexing | high; Graph.AddNode/Graph.init CRITICAL |
| `internal/scopeir/facts.go` | `E0-P0A-FD2` | 231 | mixed ScopeIR fact contracts | medium; dedicated owner extraction required |
| `internal/scopeir/range.go` | `E0-P0A-FD3` | 227 | shared position/range contract | medium; all providers depend on encoding |
| `internal/scopeir/definition_index.go` | `E0-P0A-FD4` | 225 | definition identity index | medium; currently hides duplicates |
| `internal/resolution/indexes.go` | `E0-P0A-FD5` | 46 | identity construction and resolution indexes | high; several CRITICAL symbols |
| `internal/resolution/emit.go` | `E0-P0A-FD6` | 42 | graph projection | high identity parity scope |
| `internal/lbugload/csv.go` | `E0-P0A-FD7` | 19 | persistence projection boundary | high duplicate/closure risk |
| This child plan set | `E0-P0A-PLAN1`, `E0-P0A-FDPLAN1` | 4 docs / plan has 1 related file | plan/evidence/benchmark/status; roadmap imports child plan | low code impact; high execution-authority importance |
| `E:\cheapapp.org` | `E0-P0A-BOUNDARY1` | validation-only when explicitly opened | separate target repository | critical do-not-contaminate boundary |

## Status Rules

| Status | Meaning | Allowed next action |
|--------|---------|---------------------|
| `correct` | Already behaves as required. | Preserve. Add evidence/tests only if needed. |
| `partial` | Some required behavior exists, but gaps remain. | Change only missing parts. Preserve correct parts. |
| `wrong` | Current behavior/source/contract conflicts with requirement. | Replace only through the mapped slice. |
| `missing` | Required behavior/source/contract does not exist. | Implement the missing piece only. |
| `unbound` | Surface exists but is not wired to the real source/flow/contract. | Bind to the real source only. |
| `fake-or-stub` | Prototype/mock/fallback is presented as real. | Remove or replace with an approved truthful state. |
| `blocked` | Authority, dependency, or evidence is unavailable. | Stop until resolved. |

## Current Status Matrix

| Unit | Current State | Required State | Status | Relationship Count | Evidence | Next Plan Decision |
|------|---------------|----------------|--------|--------------------|----------|--------------------|
| Declaration/Symbol identity and strict graph construction | Graph IDs omit scope/range, duplicate nodes can overwrite, and occurrence/source-site conservation is not guaranteed. | Versioned, deterministic, collision-safe Declaration/Symbol/Relationship identity with strict fail-closed mutation and lossless shadow-v2 proof. | wrong | 19-238 | `E0-P0A-STATUS1`, `E0-P0A-SOURCE1` | implement only through ordered local P1 after unblock |
| Child plan structure | four standard files and 11 mapped slices authored | complete independent P0/P1/Pn plan set | correct | 4 docs / 11 slices | `E0-P0A-PLAN1`, `E0-P0A-SOURCE1` | preserve and keep ledgers current |
| Upstream dependency | No upstream implementation child applies to child 01 | no upstream handoff required | correct | N/A | `E0-P0A-DEPENDENCY1` | preserve; local P1-A owns architecture ratification |
| Implementation authorization | multi-plan authoring is authorized; campaign cutover and production implementation are not | campaign authority cutover plus explicit owner direction for this child and each slice | blocked | N/A | `E0-P0A-DEPENDENCY1` | wait for campaign cutover and owner direction; no upstream handoff applies |
| Target boundary | target remains separate and untouched by authoring | in-place validation only in explicitly mapped slice | correct | validation-only | `E0-P0A-BOUNDARY1` | preserve |
| Scanner omission | exact eight omissions remain | zero additional omissions; repair out of scope | wrong | 8 omitted / 0 new | `E0-P0A-SCANNER1` | quarantine; do not fix here |

## Status Refresh Log

| Refresh | Date | Repo Basis | Changed Scope | Status Changes | Evidence | Next Phase Update |
|---------|------|------------|----------------|----------------|----------|-------------------|
| R0 | 2026-07-28 | `master` at roadmap commit `c444e8c4`; no production diff since `1932359b` | child 01 plan set and inherited P1 baseline | initial child classification; P0 complete; implementation blocked by campaign authority/owner authorization | `E0-P0A-PLAN1`, `E0-P0A-SOURCE1`, `E0-P0A-STATUS1`, `E0-P0A-DEPENDENCY1`, `E0-P0A-BOUNDARY1` | refresh after campaign cutover and explicit owner direction before local P1; no upstream handoff applies |
| R1 | 2026-07-28 | same HEAD after child authoring; fresh Anvien graph includes roadmap and child docs | child-plan document relationship | plan relationship `not indexed -> correct`: one roadmap inbound link, low risk, zero unresolved | `E0-P0A-FDPLAN1` | preserve candidate; no implementation gate changed |

## Phase Touch Map

| Unit / File / Surface | Plan-Relevant Relationship File | Relationship to Target | Plan Item | Touch Mode | Evidence | Constraint |
|-----------------------|---------------------------------|------------------------|-----------|------------|----------|------------|
| `internal/graph/types.go` | legacy P1 slice owner set | production source / consumer | P1 | inspect-only until exact slice opens; then scoped edit | `E0-P0A-FD1` | refresh file-detail + upstream impact before edit; HIGH/CRITICAL warns scope but does not ban work |
| `internal/scopeir/facts.go` | legacy P1 slice owner set | production source / consumer | P1 | inspect-only until exact slice opens; then scoped edit | `E0-P0A-FD2` | refresh file-detail + upstream impact before edit; HIGH/CRITICAL warns scope but does not ban work |
| `internal/scopeir/range.go` | legacy P1 slice owner set | production source / consumer | P1 | inspect-only until exact slice opens; then scoped edit | `E0-P0A-FD3` | refresh file-detail + upstream impact before edit; HIGH/CRITICAL warns scope but does not ban work |
| `internal/scopeir/definition_index.go` | legacy P1 slice owner set | production source / consumer | P1 | inspect-only until exact slice opens; then scoped edit | `E0-P0A-FD4` | refresh file-detail + upstream impact before edit; HIGH/CRITICAL warns scope but does not ban work |
| `internal/resolution/indexes.go` | legacy P1 slice owner set | production source / consumer | P1 | inspect-only until exact slice opens; then scoped edit | `E0-P0A-FD5` | refresh file-detail + upstream impact before edit; HIGH/CRITICAL warns scope but does not ban work |
| `internal/resolution/emit.go` | legacy P1 slice owner set | production source / consumer | P1 | inspect-only until exact slice opens; then scoped edit | `E0-P0A-FD6` | refresh file-detail + upstream impact before edit; HIGH/CRITICAL warns scope but does not ban work |
| `internal/lbugload/csv.go` | legacy P1 slice owner set | production source / consumer | P1 | inspect-only until exact slice opens; then scoped edit | `E0-P0A-FD7` | refresh file-detail + upstream impact before edit; HIGH/CRITICAL warns scope but does not ban work |
| This child standard set | roadmap and legacy source phase | plan authority/ledgers | P0-Pn | edit docs | `E0-P0A-PLAN1` | one primary responsibility per file |
| `E:\cheapapp.org` | mapped target-validation slice only | validation subject | P1 | do-not-touch until explicitly opened; then validate-only/normal .anvien output | `E0-P0A-BOUNDARY1` | no source/report/probe/fixture/temp write |

## Detailed Findings

### Declaration/Symbol identity and strict graph construction

Current state:

Graph IDs omit scope/range, duplicate nodes can overwrite, and occurrence/source-site conservation is not guaranteed.

Required state:

```text
Versioned, deterministic, collision-safe Declaration/Symbol/Relationship identity with strict fail-closed mutation and lossless shadow-v2 proof.
```

Evidence:

- `E0-P0A-SOURCE1`: exact source phase and slice mapping.
- `E0-P0A-STATUS1`: inherited production status remains current because production diff is empty.
- `E0-P0A-DEPENDENCY1`: campaign-authority and owner-authorization gates; no upstream implementation child applies.

Relationship and impact:

- Related file count: 19-238.
- Relationship summary: see the scoped relationship table and local slice boundaries.
- Impact note: HIGH/CRITICAL values are blast-radius warnings; they require narrow edits and strong validation, not an edit ban.

Classification:

`wrong` for product behavior; `blocked` for implementation start.

Allowed next action:

After campaign authority cutover, refresh P0, obtain explicit owner direction, and open local P1-A only. No upstream-child handoff is required.

Forbidden next action:

Do not implement from candidate status, skip/reorder slices, use unqualified cross-child evidence, mutate another child's owner, or contaminate the target.

## Next Phase Status Decisions

| Plan Item | Actual Status Finding | Required Status / Next-Action Update |
|-----------|-----------------------|--------------------------------------|
| P1-A | child structure is correct; no upstream handoff applies; product behavior is wrong; campaign cutover/owner authorization is blocked | keep goal/scope/order; refresh evidence and unblock only after campaign cutover plus explicit owner direction |
| P1-B..last | source slices are mapped but unopened | preserve order; update only stale assumptions after each prior slice |
| Pn | no implementation exists to close | remain pending; review/cleanup/close this child only |

## Implementation Gate

- [x] Target scope is listed in Current Status Matrix.
- [x] Each target unit has a status and exact local evidence IDs.
- [x] Relationship counts or truthful bounded inventories are recorded.
- [x] Phase Touch Map lists relevant files and touch modes.
- [x] Correct parts and target boundaries are preserve-only.
- [x] Wrong/missing/partial/unbound behavior has the exact ordered local P1 action.
- [x] Dependency and implementation-authorization blockers are recorded.
- [x] R0 baseline exists and later refresh obligations are explicit.
- [x] Upstream dependency is N/A for child 01; no qualified upstream handoff is required.
- [ ] Explicit owner direction authorizes this child implementation.
- [ ] Fresh Anvien analyze, file-detail, and impact evidence exists for the first edited production owner.

## Final P0 Decision

Choose one:

- [ ] P0 actual-status incomplete. Implementation is blocked.
- [ ] P0 complete. Next phase can proceed unchanged.
- [ ] P0 complete. Next phase status, next action, or work steps must be updated before implementation.
- [ ] P0 complete. Target scope is preserve-only.
- [x] P0 complete. Implementation is blocked by missing authority or evidence.

Decision note:

P0 is complete for candidate plan authoring. Local P1 remains blocked until campaign authority cuts over, this status is refreshed from current repo reality, and the owner explicitly authorizes implementation. No upstream implementation child applies. Multi-plan authoring is not production authorization.
