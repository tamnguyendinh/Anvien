# Versioned Persistence and Identity v2 Cutover Actual Status

Title: Versioned Persistence and Identity v2 Cutover
Date: 2026-07-28
Status: P0 Complete / authoring Supervisor pending / implementation blocked
Companion plan: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-02-versioned-persistence-and-v2-cutover/2026-07-28-02-versioned-persistence-and-v2-cutover-plan.md`
Companion evidence: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-02-versioned-persistence-and-v2-cutover/2026-07-28-02-versioned-persistence-and-v2-cutover-evidence.md`
Companion benchmark: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-02-versioned-persistence-and-v2-cutover/2026-07-28-02-versioned-persistence-and-v2-cutover-benchmark.md`
Roadmap: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-26-anvien-graph-identity-resolution-v2-roadmap.md`
Source phase: legacy `P2`

## Purpose

This file records the real current state before implementation.

Implementation must not start until the target scope has a completed status row, evidence IDs, and a downstream plan decision.

This file does not replace `evidence.md`. It classifies current state from evidence.

Use exact evidence IDs from `evidence.md`, such as `E0-P0A-SOURCE1`, not broad section IDs.

## Freshness / Refresh Rules

This actual-status file is a living current-state record, not a one-time P0 snapshot.

Update it after each completed local slice, before the next slice/phase if repo reality changed, whenever classifications change, and after every qualified upstream handoff. Append refresh rows; do not delete history. Update only stale next-action/work-step assumptions while preserving the accepted goal, scope, acceptance, and order.

## Scope

Target scope:

- Versioned persistence, reader compatibility, atomic generation, and identity-v2 cutover.
- 42 local implementation slices mapped exactly from legacy P2.
- This child's independent evidence, benchmark, current-status, closure, and roadmap handoff.

Out of scope:

- No binding/export/barrel/ambient semantic implementation.
- No reduction, relocation, or duplication of `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/index-reader-matrix.md`.
- No cutover before all 41 predecessor slices and child-01 handoff pass.
- Multi-plan authoring itself does not authorize production or target execution.

## Relationship / Impact Evidence

For each target file, prefer:

```text
anvien file-detail <path> --repo <repo> --json
```

Inherited counts below remain valid because no production path changed after the source P0. Refresh file-detail plus upstream impact immediately before the exact implementation slice edits a file.

| Unit / File / Surface | File Detail Evidence | Related File Count | Relationship Summary | Impact Note |
|-----------------------|----------------------|--------------------|----------------------|-------------|
| `internal/analyze/analyze.go` | `E0-P0A-FD1` | 182 | analyze/build/load/snapshot orchestration | high; publication CRITICAL |
| `internal/repo/types.go` | `E0-P0A-FD2` | 72 | persisted repository metadata | high; version/generation fields |
| `internal/graph/types.go` | `E0-P0A-FD3` | 238 | canonical graph records | high/CRITICAL identity storage |
| `internal/lbugload/csv.go` | `E0-P0A-FD4` | 19 | Ladybug schema/CSV projection | high duplicate/parity risk |
| `internal/httpapi/graph.go` | `E0-P0A-FD5` | 22 | HTTP graph contract | high public contract |
| `anvien-web/src/services/backend-client.ts` | `E0-P0A-FD6` | 24 | Web stream/version parsing | high public runtime |
| `internal/repo/registry.go / internal/group/sync.go / internal/group/storage.go` | `E0-P0A-FD7` | bounded | global registry and group publication | high generation scope |
| `internal/mcp/resource_cache.go / internal/filecontext/* / embedding stores` | `E0-P0A-FD8` | bounded | cache and embedding readers | high stale-reference scope |
| `native and fallback Cypher readers` | `E0-P0A-FD9` | bounded | alternate query backends | high parity scope |
| This child plan set | `E0-P0A-PLAN1`, `E0-P0A-FD11` | 4 docs / 1 inbound roadmap link per file | plan/evidence/benchmark/status | low code impact; high execution-authority importance |
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
| Versioned persistence, reader compatibility, atomic generation, and identity-v2 cutover | Index versions/generations are incomplete, readers are not uniformly guarded, IDs are parsed semantically, and publication is not atomic. | Fail-closed negotiated v2 readers, canonical S0-S11 parity, opaque references, immutable atomic generations, and one validated active-v2 cutover. | missing | 19-238 / bounded reader inventory | `E0-P0A-STATUS1`, `E0-P0A-SOURCE1` | implement only through ordered local P1 after unblock |
| Child plan structure | four standard files and 42 mapped slices authored | complete independent P0/P1/Pn plan set | correct | 4 docs / 42 slices | `E0-P0A-PLAN1`, `E0-P0A-SOURCE1` | preserve and keep ledgers current |
| Upstream authoring basis | Child 01 authoring is accepted/committed at `b760d156`; its path-preserving structural move and the three-root contract are committed at `55bf021f` | accepted/committed prior-child authoring basis for child-02 authoring | correct | N/A | `E0-P0A-DEPENDENCY1`, `E0-P0A-GRAPH1` | authoring review may proceed; this does not open local P1 |
| Upstream implementation handoff | `2026-07-28-01-graph-identity-contract-and-strict-construction::E2-PNC-HANDOFF1` does not yet exist | accepted/committed qualified implementation handoff binding identity authority, shadow proof, benchmark, review, and implementation commit | blocked | N/A | `E0-P0A-DEPENDENCY1` | do not open local P1 |
| Implementation authorization | multi-plan authoring is authorized; production implementation is not | explicit owner direction for this child and each slice | blocked | N/A | `E0-P0A-DEPENDENCY1` | wait for owner direction after dependencies |
| Successor actual-status handoff | Child plan Rules/Pn-C now require a latest-evidence refresh of child 03 actual-status, refresh log, and affected next actions/work steps, with a qualified reserved proof | Child 03 status is current before child 02 closure/handoff; missing or stale status blocks closure | correct contract | 1 successor / 1 reserved proof | `E0-P0A-NEXTSTATUSRULE1` | preserve through implementation; populate `E2-PNC-NEXTSTATUS1` before Pn-C may close |
| Reader-matrix ownership contract | The matrix exists once at the preserved legacy-root path; this child plan binds sole future mutation ownership while authoring leaves content byte-identical | exactly one matrix file and exactly one mutation owner, child 02 | correct | 1 file | `E0-P0A-MATRIX1` | preserve content during authoring; mutate only in local P1-A1 after implementation gates open |
| Target boundary | target remains separate and untouched by authoring | in-place validation only in explicitly mapped slice | correct | validation-only | `E0-P0A-BOUNDARY1` | preserve |
| Scanner omission | exact eight omissions remain | zero additional omissions; repair out of scope | wrong | 8 omitted / 0 new | `E0-P0A-SCANNER1` | quarantine; do not fix here |

## Status Refresh Log

| Refresh | Date | Repo Basis | Changed Scope | Status Changes | Evidence | Next Phase Update |
|---------|------|------------|----------------|----------------|----------|-------------------|
| R0 | 2026-07-28 | `master` at roadmap commit `c444e8c4`; no production diff since `1932359b` | child 02 plan set and inherited P2 baseline | initial child classification; P0 complete; implementation blocked by dependency/authority | `E0-P0A-PLAN1`, `E0-P0A-SOURCE1`, `E0-P0A-STATUS1`, `E0-P0A-DEPENDENCY1`, `E0-P0A-BOUNDARY1` | historical baseline; superseded for current decisions by R1/R2 |
| R1 | 2026-07-28 | `master` at committed three-root basis `55bf021f`; fresh Anvien index at `2026-07-28T07:34:44Z` | child-01 authoring/structural dependency split, matrix ownership contract, and child-02 P0 freshness | authoring basis `blocked -> correct`; implementation handoff remains `blocked`; matrix owner contract `unbound -> correct` | `E0-P0A-GRAPH1`, `E0-P0A-FD10`, `E0-P0A-GIT1`, `E0-P0A-DEPENDENCY1`, `E0-P0A-MATRIX1` | proceed only with P2-B authoring review; keep local P1 closed |
| R2 | 2026-07-28 | `master` at `55bf021f`; fresh post-link Anvien index at `2026-07-28T07:44:40Z` | child-02 candidate links, current relationship evidence, and authoring-ledger refresh | child plan relationship `0 -> 1` inbound roadmap link; candidate remains `partial / Supervisor pending`; implementation remains `blocked` | `E0-P0A-GRAPH2`, `E0-P0A-FD11`, `E0-P0A-GIT2` | obtain unconditional authoring Supervisor PASS before marking P2-B complete/committing; implementation handoff and owner authorization remain separate blockers |
| R3 | 2026-07-28 | `master` at `55bf021f`; user-required successor-freshness correction during P2-B review | child-02 Rules, Pn-C, closure evidence reservation, and child-03 handoff qualification | successor freshness `implicit/roadmap-only -> correct child-owned contract`; implementation remains `blocked` | `E0-P0A-NEXTSTATUSRULE1` | preserve the rule; at Pn-C refresh child 03 actual-status and populate qualified `E2-PNC-NEXTSTATUS1` before closure |

## Phase Touch Map

| Unit / File / Surface | Plan-Relevant Relationship File | Relationship to Target | Plan Item | Touch Mode | Evidence | Constraint |
|-----------------------|---------------------------------|------------------------|-----------|------------|----------|------------|
| `internal/analyze/analyze.go` | legacy P2 slice owner set | production source / consumer | P1 | inspect-only until exact slice opens; then scoped edit | `E0-P0A-FD1` | refresh file-detail + upstream impact before edit; HIGH/CRITICAL warns scope but does not ban work |
| `internal/repo/types.go` | legacy P2 slice owner set | production source / consumer | P1 | inspect-only until exact slice opens; then scoped edit | `E0-P0A-FD2` | refresh file-detail + upstream impact before edit; HIGH/CRITICAL warns scope but does not ban work |
| `internal/graph/types.go` | legacy P2 slice owner set | production source / consumer | P1 | inspect-only until exact slice opens; then scoped edit | `E0-P0A-FD3` | refresh file-detail + upstream impact before edit; HIGH/CRITICAL warns scope but does not ban work |
| `internal/lbugload/csv.go` | legacy P2 slice owner set | production source / consumer | P1 | inspect-only until exact slice opens; then scoped edit | `E0-P0A-FD4` | refresh file-detail + upstream impact before edit; HIGH/CRITICAL warns scope but does not ban work |
| `internal/httpapi/graph.go` | legacy P2 slice owner set | production source / consumer | P1 | inspect-only until exact slice opens; then scoped edit | `E0-P0A-FD5` | refresh file-detail + upstream impact before edit; HIGH/CRITICAL warns scope but does not ban work |
| `anvien-web/src/services/backend-client.ts` | legacy P2 slice owner set | production source / consumer | P1 | inspect-only until exact slice opens; then scoped edit | `E0-P0A-FD6` | refresh file-detail + upstream impact before edit; HIGH/CRITICAL warns scope but does not ban work |
| `internal/repo/registry.go / internal/group/sync.go / internal/group/storage.go` | legacy P2 slice owner set | production source / consumer | P1 | inspect-only until exact slice opens; then scoped edit | `E0-P0A-FD7` | refresh file-detail + upstream impact before edit; HIGH/CRITICAL warns scope but does not ban work |
| `internal/mcp/resource_cache.go / internal/filecontext/* / embedding stores` | legacy P2 slice owner set | production source / consumer | P1 | inspect-only until exact slice opens; then scoped edit | `E0-P0A-FD8` | refresh file-detail + upstream impact before edit; HIGH/CRITICAL warns scope but does not ban work |
| `native and fallback Cypher readers` | legacy P2 slice owner set | production source / consumer | P1 | inspect-only until exact slice opens; then scoped edit | `E0-P0A-FD9` | refresh file-detail + upstream impact before edit; HIGH/CRITICAL warns scope but does not ban work |
| `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/index-reader-matrix.md` | reader inventory and guard ownership | auxiliary source contract | P1-A1 | edit | `E0-P0A-MATRIX1` | sole mutation owner; source-rescan must prove complete |
| This child standard set | roadmap and legacy source phase | plan authority/ledgers | P0-Pn | edit docs | `E0-P0A-PLAN1` | one primary responsibility per file |
| `E:\cheapapp.org` | mapped target-validation slice only | validation subject | P1 | do-not-touch until explicitly opened; then validate-only/normal .anvien output | `E0-P0A-BOUNDARY1` | no source/report/probe/fixture/temp write |

## Detailed Findings

### Versioned persistence, reader compatibility, atomic generation, and identity-v2 cutover

Current state:

Index versions/generations are incomplete, readers are not uniformly guarded, IDs are parsed semantically, and publication is not atomic.

Required state:

```text
Fail-closed negotiated v2 readers, canonical S0-S11 parity, opaque references, immutable atomic generations, and one validated active-v2 cutover.
```

Evidence:

- `E0-P0A-SOURCE1`: exact source phase and slice mapping.
- `E0-P0A-STATUS1`: inherited production status remains current because production diff is empty.
- `E0-P0A-DEPENDENCY1`: upstream and authorization gate.

Relationship and impact:

- Related file count: 19-238 / bounded reader inventory.
- Relationship summary: see the scoped relationship table and local slice boundaries.
- Impact note: HIGH/CRITICAL values are blast-radius warnings; they require narrow edits and strong validation, not an edit ban.

Classification:

`missing` for product behavior; `blocked` for implementation start.

Allowed next action:

Treat the committed child-01 authoring set and structural move as sufficient only for P2-B authoring review. Before local P1-A opens, refresh P0 again after `2026-07-28-01-graph-identity-contract-and-strict-construction::E2-PNC-HANDOFF1` exists, verify that it binds the required accepted implementation evidence, then obtain explicit owner direction. Matrix mutation remains owned only by this child and begins no earlier than local P1-A1.

Forbidden next action:

Do not implement from candidate status, skip/reorder slices, use unqualified cross-child evidence, mutate another child's owner, or contaminate the target.

## Next Phase Status Decisions

| Plan Item | Actual Status Finding | Required Status / Next-Action Update |
|-----------|-----------------------|--------------------------------------|
| P1-A | child structure and committed authoring basis are correct; product behavior is missing; qualified implementation handoff and authorization remain blocked | keep goal/scope/order; refresh evidence and unblock only after `2026-07-28-01-graph-identity-contract-and-strict-construction::E2-PNC-HANDOFF1` plus explicit owner direction |
| P1-B..last | source slices are mapped but unopened | preserve order; update only stale assumptions after each prior slice |
| Pn | no implementation exists to close; successor-freshness contract is now explicit | remain pending; review/cleanup/close this child only after refreshing child 03 actual-status and recording qualified `E2-PNC-NEXTSTATUS1` |

## Implementation Gate

- [x] Target scope is listed in Current Status Matrix.
- [x] Each target unit has a status and exact local evidence IDs.
- [x] Relationship counts or truthful bounded inventories are recorded.
- [x] Phase Touch Map lists relevant files and touch modes.
- [x] Correct parts and target boundaries are preserve-only.
- [x] Wrong/missing/partial/unbound behavior has the exact ordered local P1 action.
- [x] Dependency and implementation-authorization blockers are recorded.
- [x] R0 baseline exists and later refresh obligations are explicit.
- [ ] Upstream dependency is accepted/committed and consumed with qualified evidence.
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

P0 is current for candidate plan authoring at committed basis `55bf021f`: child-01 authoring and the three-root move are committed, the child-02 four-file set is linked as a Supervisor-pending candidate, the single reader-matrix mutation owner is bound to this child while matrix content remains unchanged, and Pn-C now blocks on a latest-evidence child-03 actual-status refresh proven by qualified `E2-PNC-NEXTSTATUS1`. Local P1 remains blocked because `2026-07-28-01-graph-identity-contract-and-strict-construction::E2-PNC-HANDOFF1` does not yet exist and the Owner has not authorized implementation. P2-B authoring acceptance cannot be treated as production authorization.
