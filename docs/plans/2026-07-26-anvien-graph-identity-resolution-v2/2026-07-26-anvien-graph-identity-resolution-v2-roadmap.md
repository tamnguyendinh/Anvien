# Anvien Graph Identity and TypeScript Resolution Correctness v2 Roadmap

Date: 2026-07-28
Status: candidate / child plans not yet authored / legacy plan remains active
Source plan: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/2026-07-26-anvien-graph-identity-resolution-v2-plan.md`
Plan-set authoring plan: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/2026-07-28-00-multi-plan-authoring/2026-07-28-00-multi-plan-authoring-plan.md`
Exact crosswalk: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/2026-07-28-00-multi-plan-authoring/2026-07-28-00-multi-plan-authoring-evidence.md#e1-p1a-map1---exact-source-to-child-crosswalk`

## Goal

Coordinate the graph-identity and TypeScript-resolution remediation as seven complete implementation child plans, one child for each accepted legacy implementation phase P1-P7. Keep technical order, all 98 implementation slices, independent P0/evidence/benchmark/closure ledgers, explicit handoffs, and one active authority without turning this roadmap into another implementation plan.

## Authority State

- Active authority now: the legacy source plan at SHA-256 `365E17A7F7CD539426568A2874A1AD1231D120C5391BFE3D8B875D5F760A45FB`.
- Candidate authority: this roadmap plus seven numbered child plan sets after all are authored and verified.
- Partial child output is draft-only and does not authorize production implementation.
- Authority cutover occurs only after the authoring plan's P3-A deterministic checks and P3-B unconditional Supervisor PASS.
- During cutover, preserve the legacy technical body and companion ledgers. Change only its status/pointer to `superseded / reference-only`.
- After cutover, this roadmap is the sole campaign-order/status index and each child plan is the sole execution authority for its mapped legacy phase.
- Multi-plan authoring does not authorize production implementation. Opening an implementation child still requires explicit owner direction under that child's gate.

## Campaign Invariants

- The child count is seven because the source has seven implementation phases. The sample campaign's child count is not a formula.
- Legacy P1-P7 map one-to-one to children 01-07.
- Legacy P8 is closure, not implementation. Its A/B/C roles become every child's Pn-A/Pn-B/Pn-C; no child 08 exists.
- Every child is a complete four-file standard plan set with populated P0, one local implementation phase P1, all mapped slices, and tailored Pn closure.
- Every source slice maps exactly once. Child-local IDs remap `P<source>-<suffix>` to `P1-<suffix>` and every slice records `Source slice: P<source>-<suffix>`.
- Preserve source order, goals, scope, pre-flight fields, work-step order, gates, acceptance, evidence targets, actual-status transitions, and commit boundaries.
- Cross-child references always include the child slug plus the exact evidence ID. Bare local IDs are not valid cross-child evidence.
- After each implementation slice, update its child checklist/ledgers, run required validation/Supervisor/detect-changes, commit, and only then open the next slice when authorized.
- Before a child implementation plan opens, refresh its P0 from current repo reality and the accepted previous-child handoff.
- A later validation child cannot repair implementation. A failed acceptance gate reopens the responsible upstream child/slice.
- Every file owns one primary planning or implementation responsibility. A file may link to multiple modules/files only when all links serve that responsibility.

## Frozen Source Snapshot

| Source phase | Source title | Source slices | Child |
|--------------|--------------|--------------:|-------|
| P1 | Graph identity contract and strict graph construction | 11 | 01 |
| P2 | Versioned persistence, opaque consumers, atomic generation, and v2 cutover | 42 | 02 |
| P3 | TypeScript binding-pattern extraction | 17 | 03 |
| P4 | First-class TypeScript export semantics | 15 | 04 |
| P5 | Module export tables and barrel/re-export resolution | 4 | 05 |
| P6 | Ambient/external declaration universe and truthful diagnostics | 6 | 06 |
| P7 | Cross-surface acceptance, target validation, and performance | 3 | 07 |
| **Total** | **Implementation only** | **98** | **7 children** |

Legacy P8 contains three closure roles. It is excluded from the 98-row implementation crosswalk and distributed into all seven child plan closures.

## Child Plan Inventory

| No. | Plan folder | Primary responsibility | Source | Slices | Status | Depends on / handoff |
|-----|-------------|------------------------|--------|-------:|--------|----------------------|
| 01 | `2026-07-28-01-graph-identity-contract-and-strict-construction` | Identity contract, lossless declaration/source-site identity, strict graph construction, shadow-v2 proof | P1 | 11 | not authored | Accepted authoring roadmap/P0 |
| 02 | `2026-07-28-02-versioned-persistence-and-v2-cutover` | Compatibility manifest, opaque readers, S0-S11 parity, atomic generations, v2 cutover | P2 | 42 | not authored | Child 01 accepted/committed; owns reader matrix |
| 03 | `2026-07-28-03-typescript-binding-pattern-extraction` | Recursive binding facts, declaration contexts, graph and adapter projection | P3 | 17 | not authored | Child 02 identity-v2 cutover handoff |
| 04 | `2026-07-28-04-typescript-export-semantics` | ExportFact semantics, export syntax extraction, graph and adapter projection | P4 | 15 | not authored | Child 03; child 01 decision authority; child 02 matrix inspect-only |
| 05 | `2026-07-28-05-module-export-and-reexport-resolution` | Module export tables, aliases, cycles, ambiguity, terminal barrel resolution | P5 | 4 | not authored | Child 04 ExportFact/re-export handoff |
| 06 | `2026-07-28-06-ambient-external-resolution-and-diagnostics` | Declaration universe, external authorization/materialization, immutable outcomes, diagnostics | P6 | 6 | not authored | Child 05 candidate handoff; child 02 matrix inspect-only |
| 07 | `2026-07-28-07-cross-surface-acceptance-and-target-validation` | Full determinism/parity/target/performance acceptance; no implementation repair | P7 | 3 | not authored | All children 01-06 accepted/committed |

## Standard File Inventory

Every row below represents four required files inside its plan folder. Planned paths remain code-form until the corresponding P2 authoring slice creates them; afterward the roadmap must link the real files.

| Child | Plan | Evidence | Benchmark | Actual status |
|-------|------|----------|-----------|---------------|
| 01 | `2026-07-28-01-graph-identity-contract-and-strict-construction-plan.md` | `2026-07-28-01-graph-identity-contract-and-strict-construction-evidence.md` | `2026-07-28-01-graph-identity-contract-and-strict-construction-benchmark.md` | `2026-07-28-01-graph-identity-contract-and-strict-construction-actual-status.md` |
| 02 | `2026-07-28-02-versioned-persistence-and-v2-cutover-plan.md` | `2026-07-28-02-versioned-persistence-and-v2-cutover-evidence.md` | `2026-07-28-02-versioned-persistence-and-v2-cutover-benchmark.md` | `2026-07-28-02-versioned-persistence-and-v2-cutover-actual-status.md` |
| 03 | `2026-07-28-03-typescript-binding-pattern-extraction-plan.md` | `2026-07-28-03-typescript-binding-pattern-extraction-evidence.md` | `2026-07-28-03-typescript-binding-pattern-extraction-benchmark.md` | `2026-07-28-03-typescript-binding-pattern-extraction-actual-status.md` |
| 04 | `2026-07-28-04-typescript-export-semantics-plan.md` | `2026-07-28-04-typescript-export-semantics-evidence.md` | `2026-07-28-04-typescript-export-semantics-benchmark.md` | `2026-07-28-04-typescript-export-semantics-actual-status.md` |
| 05 | `2026-07-28-05-module-export-and-reexport-resolution-plan.md` | `2026-07-28-05-module-export-and-reexport-resolution-evidence.md` | `2026-07-28-05-module-export-and-reexport-resolution-benchmark.md` | `2026-07-28-05-module-export-and-reexport-resolution-actual-status.md` |
| 06 | `2026-07-28-06-ambient-external-resolution-and-diagnostics-plan.md` | `2026-07-28-06-ambient-external-resolution-and-diagnostics-evidence.md` | `2026-07-28-06-ambient-external-resolution-and-diagnostics-benchmark.md` | `2026-07-28-06-ambient-external-resolution-and-diagnostics-actual-status.md` |
| 07 | `2026-07-28-07-cross-surface-acceptance-and-target-validation-plan.md` | `2026-07-28-07-cross-surface-acceptance-and-target-validation-evidence.md` | `2026-07-28-07-cross-surface-acceptance-and-target-validation-benchmark.md` | `2026-07-28-07-cross-surface-acceptance-and-target-validation-actual-status.md` |

Target inventory: 7 plan folders, 28 standard child files, 7 P0 sections, 7 local P1 phases, 98 mapped implementation slices, and 7 Pn-A/Pn-B/Pn-C closure sets.

## Execution Order And Handoffs

1. Child 01 establishes identity, lossless occurrence, RelationshipID, strict mutation, and shadow-v2 contracts while preserving the active v1 path.
2. Child 02 consumes child 01's accepted authority/shadow evidence, owns `index-reader-matrix.md`, cuts over all readers and immutable generations, and activates identity v2.
3. Child 03 opens only after child 02 identity-v2 cutover. Its internal order includes variable, parameter, catch, and for-of/for-in binding contexts before projection.
4. Child 04 consumes child 03 completion, child 01's qualified architecture decision, and child 02's S0-S11 denominator inspect-only. Export facts remain distinct from access visibility and terminal barrel selection.
5. Child 05 consumes child 04 ExportFact/re-export syntax and owns the P5 semantic-vector manifest and terminal module/barrel traversal.
6. Child 06 consumes immutable external candidates from child 05, owns declaration authorization/materialization and the authoritative diagnostic status matrix, and may not redo package-export resolution.
7. Child 07 opens after all children 01-06 are accepted and committed. It validates determinism, S0-S11 parity, the real target boundary, and performance; failures return to the owning child.

Each child Pn-C must update this roadmap and refresh the next child's actual-status before handing off. A child closure is a campaign checkpoint, not campaign completion.

## Cross-Child Evidence Contract

- A handoff reference has the form `{child-slug, exact evidence ID, source slice, local slice}`.
- Child-local evidence namespaces may repeat across children because the child slug is the outer namespace.
- Child 02 owns the single reader matrix and its source-derived inventory. Children 03, 04, 06, and 07 consume qualified child-02 evidence inspect-only.
- Child 05 owns the semantic vector; child 06 owns the resolution-status matrix; child 07 owns final cross-surface acceptance evidence.
- Global source provenance remains available from the legacy plan/ledgers and the authoring crosswalk; it does not replace fresh child P0 or implementation evidence.

## Known Source-Migration Hazards

- Child 01 keeps the strict source chain A -> B -> C0 -> C0A -> C0B -> C -> D -> D1 -> D2 -> D3 -> E. Shadow-v2 must not mutate the active v1 path.
- Child 02 remains serial across all 42 slices even when an individual legacy gate names only a subset of earlier work. The reader matrix is a source seed; it is not runtime completion evidence.
- Legacy P3-C omits immediately preceding P3-B2A from its written gate even though ordered execution and consumed binding facts require it. Child 03 must explicitly include local P1-B2A before P1-C and record this source normalization.
- Legacy P3-C1 has an ordered-list title different from its executable checklist heading. Child 03 uses the checklist title `Project binding JSON/Ladybug persistence adapters` and records the ordered-list alias as provenance.
- Legacy P4-A references `P1-A authority`; after remap this is a qualified dependency on child 01/local P1-A, not child 04's own local P1-A.
- Children 03 and 04 preserve the source adapter order S0-S2, S3, S4, S7, S5, S8, S6, S9, S10, S11. Do not reorder numerically or merge independently committed surfaces.
- Child 05 retains the full P5 semantic-vector manifest even though it sits outside the four slice bodies.
- Child 06 retains the full P6 authoritative status matrix even though it sits after its slice bodies.
- Child 07 owns validation only; it never patches a failure and never takes mutation ownership of the child-02 reader matrix.
- Legacy evidence section headings under-list some expanded nested slices while their tables and traceability rows contain the real evidence inventory. Child authoring must use exact slice/evidence rows, not the narrow heading text.
- Every child Pn is tailored to that child. Pn-C hands off to the next child/roadmap and cannot claim whole-campaign completion.

## Shared Artifact Ownership

| Artifact / contract | Sole mutation owner | Other consumers |
|---------------------|---------------------|-----------------|
| `index-reader-matrix.md` | Child 02 | Children 03/04/06 inspect; child 07 validates |
| Identity/ownership decision | Child 01 | Children 02-07 consume qualified handoff |
| P5 semantic vector manifest | Child 05 | Children 06-07 consume qualified result |
| P6 authoritative resolution-status matrix | Child 06 | Child 07 validates |
| Campaign order/status/authority | This roadmap | All child Pn closures update status/handoff only |
| Legacy history | Legacy plan and ledgers | Read-only provenance after cutover |

## Target And Repository Boundary

- Production and plan authoring occur only in `E:\Anvien`.
- `E:\cheapapp.org` is never copied into Anvien.
- Future target acceptance analyzes `E:\cheapapp.org` in place only when the responsible child explicitly authorizes it.
- Operational target graph/index/staging remains under `E:\cheapapp.org\.anvien`.
- Plans, reports, probes, fixtures, QA evidence, and Supervisor artifacts remain in `E:\Anvien`.
- No report, probe, fixture, or temporary investigation artifact may be written into the target.
- Preserve the target's pre-existing worktree. Never reset, clean, repair, or silently revert it.
- Supported analyze side effects on ignored generated guidance files must be measured and recorded, not manually reverted.
- The exact eight scanner omissions remain quarantined and out of scope; require zero additional omissions.
- Authoring this campaign does not authorize target analysis or target writes.

## Roadmap Status Protocol

Use these states:

- `not authored`: child standard files do not exist.
- `candidate / P0 complete`: child set exists and is structurally complete but campaign authority has not cut over.
- `ready / implementation not authorized`: campaign cutover passed; child is executable only after explicit owner direction and dependency/P0 refresh.
- `active`: the child is the currently authorized implementation plan.
- `complete / committed <hash>`: child acceptance and closure passed with a known commit.
- `blocked <evidence>`: exact blocker and qualified evidence are recorded.

Status history:

| Date | Event | Authority result | Evidence |
|------|-------|------------------|----------|
| 2026-07-28 | P1-A froze legacy hash and 98-row crosswalk | legacy remains active | authoring `E1-P1A-SNAPSHOT1`, `E1-P1A-MAP1`, `E1-P1A-MAP2` |
| 2026-07-28 | P1-B created candidate roadmap and seven-child inventory | legacy remains active; children not authored | authoring `E1-P1B-ROADMAP1`, `E1-P1B-INVENTORY1`, `E1-P1B-LINK1` |

## Candidate Acceptance Gate

This roadmap may become active only when:

- all seven child folders and all 28 standard files exist;
- every child has populated P0, one local P1, all mapped source slices, and tailored Pn-A/Pn-B/Pn-C;
- deterministic validation proves 98 source slices, 98 unique destinations, zero missing, zero duplicate, zero extra, and preserved order/required fields;
- all companion and cross-plan links resolve;
- every mutable artifact has one owner and `index-reader-matrix.md` belongs only to child 02;
- no placeholder, stale authority statement, unqualified cross-child evidence ID, dead draft, or repository-root artifact remains;
- Git scope contains only approved planning/report artifacts;
- production source, tests, runtime output, generated graph output, and `E:\cheapapp.org` remain unchanged;
- Supervisor gives an unconditional PASS against frozen candidate hashes.

## Campaign Completion Definition

Multi-plan authoring is complete when the candidate gate passes, the roadmap becomes active, and the legacy plan is preserved as reference-only.

The remediation campaign itself is complete only after:

- children 01-06 are implemented, validated, Supervisor-reviewed, evidenced, and committed in order;
- each next child P0 is refreshed from the prior accepted handoff before it opens;
- child 07 completes final determinism, parity, target, runtime, and performance acceptance;
- every child Pn cleanup/closure passes;
- this roadmap records all final child commits and no unresolved blocker remains.
