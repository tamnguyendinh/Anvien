# Anvien Graph Identity and TypeScript Resolution Correctness v2 Roadmap

Date: 2026-07-28
Status: candidate / child 01 committed `ce82a341` / child 02 committed `a1c66865` / child 03 committed `35a0611c` / child 04 committed `2de220bb` / child 05 exact-copy Supervisor PASS / commit pending / children 06-07 not authored / legacy plan remains active
Source plan: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/2026-07-26-anvien-graph-identity-resolution-v2-plan.md`
Plan-set authoring plan: `docs/plans/2026-07-28-00-multi-plan-authoring/2026-07-28-00-multi-plan-authoring-plan.md`
Multi-plan root: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/`
Exact crosswalk: `docs/plans/2026-07-28-00-multi-plan-authoring/2026-07-28-00-multi-plan-authoring-evidence.md#e1-p1a-map1---exact-source-to-child-crosswalk`

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
- Every source phase and slice ID is preserved unchanged in its child. Conversion is mechanical copy/paste; phase IDs, slice IDs, order, wording, scope, gates, acceptance, and evidence IDs are not remapped or rewritten.
- Preserve source order, goals, scope, pre-flight fields, work-step order, gates, acceptance, evidence targets, actual-status transitions, and commit boundaries.
- Cross-child references always include the child slug plus the exact evidence ID. Bare local IDs are not valid cross-child evidence.
- After each implementation slice, update its child checklist/ledgers, run required validation/Supervisor/detect-changes, commit, and only then open the next slice when authorized.
- Every child plan carries its own successor-freshness rule. Before a non-terminal child may close or hand off, it must refresh the next child plan's `actual-status.md` from the latest accepted repo/runtime/evidence state, append the refresh-log row, update affected next actions/work steps, and record qualified `E2-PNC-NEXTSTATUS1` proof. Missing or stale successor status blocks closure. Child 07 records that no successor exists and refreshes roadmap/campaign closure status instead.
- Before a child implementation plan opens, refresh its P0 from current repo reality and the accepted previous-child handoff.
- A later validation child cannot repair implementation. A failed acceptance gate reopens the responsible upstream child/slice.
- Every file owns one primary planning or implementation responsibility. A file may link to multiple modules/files only when all links serve that responsibility.
- The legacy plan root, split-authoring plan root, and this multi-plan root are independent siblings directly under `docs/plans/`; no authoring plan or child plan may be nested inside the legacy root.

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
| 01 | [2026-07-28-01-graph-identity-contract-and-strict-construction](2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-plan.md) | Identity contract, lossless declaration/source-site identity, strict graph construction, shadow-v2 proof | P1 | 11 | candidate / mechanical copy conversion committed `ce82a341` / P0 complete / implementation not authorized | No predecessor; closure must update the planned child-02 actual-status from latest evidence before handoff |
| 02 | [2026-07-28-02-versioned-persistence-and-v2-cutover](2026-07-28-02-versioned-persistence-and-v2-cutover/2026-07-28-02-versioned-persistence-and-v2-cutover-plan.md) | Compatibility manifest, opaque readers, S0-S11 parity, atomic generations, v2 cutover | P2 | 42 | exact-copy candidate / committed `a1c66865` / implementation not authorized | Requires child-01 implementation handoff before implementation; source P2 IDs/content preserved; sole future reader-matrix mutation owner |
| 03 | [2026-07-28-03-typescript-binding-pattern-extraction](2026-07-28-03-typescript-binding-pattern-extraction/2026-07-28-03-typescript-binding-pattern-extraction-plan.md) | Recursive binding facts, declaration contexts, graph and adapter projection | P3 | 17 | exact-copy candidate / committed `35a0611c` / implementation not authorized | Requires `2026-07-28-02-versioned-persistence-and-v2-cutover::E2-PNC-HANDOFF1` (source `P8-C`, local `Pn-C`), including identity-v2 plus S0-S11/reader-matrix inspect-only baseline |
| 04 | [2026-07-28-04-typescript-export-semantics](2026-07-28-04-typescript-export-semantics/2026-07-28-04-typescript-export-semantics-plan.md) | ExportFact semantics, export syntax extraction, graph and adapter projection | P4 | 15 | exact-copy candidate / committed `2de220bb` / implementation not authorized | Requires `2026-07-28-03-typescript-binding-pattern-extraction::E2-PNC-HANDOFF1` (source `P8-C`, local `Pn-C`); consumes `2026-07-28-01-graph-identity-contract-and-strict-construction::E1-P1A-CONTRACT1` (source/local `P1-A`) and `2026-07-28-02-versioned-persistence-and-v2-cutover::E2-PNC-HANDOFF1` inspect-only |
| 05 | [2026-07-28-05-module-export-and-reexport-resolution](2026-07-28-05-module-export-and-reexport-resolution/2026-07-28-05-module-export-and-reexport-resolution-plan.md) | Module export tables, aliases, cycles, ambiguity, terminal barrel resolution | P5 | 4 | exact-copy candidate / Supervisor PASS / commit pending / implementation not authorized | Requires `2026-07-28-04-typescript-export-semantics::E2-PNC-HANDOFF1` (source `P8-C`, local `Pn-C`) |
| 06 | `2026-07-28-06-ambient-external-resolution-and-diagnostics` | Declaration universe, external authorization/materialization, immutable outcomes, diagnostics | P6 | 6 | not authored | Requires `2026-07-28-05-module-export-and-reexport-resolution::E2-PNC-HANDOFF1` (source `P8-C`, local `Pn-C`); consumes `2026-07-28-02-versioned-persistence-and-v2-cutover::E2-PNC-HANDOFF1` inspect-only |
| 07 | `2026-07-28-07-cross-surface-acceptance-and-target-validation` | Full determinism/parity/target/performance acceptance; no implementation repair | P7 | 3 | not authored | Requires `2026-07-28-01-graph-identity-contract-and-strict-construction::E2-PNC-HANDOFF1`, `2026-07-28-02-versioned-persistence-and-v2-cutover::E2-PNC-HANDOFF1`, `2026-07-28-03-typescript-binding-pattern-extraction::E2-PNC-HANDOFF1`, `2026-07-28-04-typescript-export-semantics::E2-PNC-HANDOFF1`, `2026-07-28-05-module-export-and-reexport-resolution::E2-PNC-HANDOFF1`, and `2026-07-28-06-ambient-external-resolution-and-diagnostics::E2-PNC-HANDOFF1` (each source `P8-C`, local `Pn-C`) |

## Standard File Inventory

Every row below represents four required files inside its plan folder. Planned paths remain code-form until the corresponding P2 authoring slice creates the four-file candidate; afterward the roadmap links the real files while the Status column remains the acceptance authority.

| Child | Plan | Evidence | Benchmark | Actual status |
|-------|------|----------|-----------|---------------|
| 01 | [plan](2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-plan.md) | [evidence](2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-evidence.md) | [benchmark](2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-benchmark.md) | [actual status](2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-actual-status.md) |
| 02 | [plan](2026-07-28-02-versioned-persistence-and-v2-cutover/2026-07-28-02-versioned-persistence-and-v2-cutover-plan.md) | [evidence](2026-07-28-02-versioned-persistence-and-v2-cutover/2026-07-28-02-versioned-persistence-and-v2-cutover-evidence.md) | [benchmark](2026-07-28-02-versioned-persistence-and-v2-cutover/2026-07-28-02-versioned-persistence-and-v2-cutover-benchmark.md) | [actual status](2026-07-28-02-versioned-persistence-and-v2-cutover/2026-07-28-02-versioned-persistence-and-v2-cutover-actual-status.md) |
| 03 | [plan](2026-07-28-03-typescript-binding-pattern-extraction/2026-07-28-03-typescript-binding-pattern-extraction-plan.md) | [evidence](2026-07-28-03-typescript-binding-pattern-extraction/2026-07-28-03-typescript-binding-pattern-extraction-evidence.md) | [benchmark](2026-07-28-03-typescript-binding-pattern-extraction/2026-07-28-03-typescript-binding-pattern-extraction-benchmark.md) | [actual status](2026-07-28-03-typescript-binding-pattern-extraction/2026-07-28-03-typescript-binding-pattern-extraction-actual-status.md) |
| 04 | [plan](2026-07-28-04-typescript-export-semantics/2026-07-28-04-typescript-export-semantics-plan.md) | [evidence](2026-07-28-04-typescript-export-semantics/2026-07-28-04-typescript-export-semantics-evidence.md) | [benchmark](2026-07-28-04-typescript-export-semantics/2026-07-28-04-typescript-export-semantics-benchmark.md) | [actual status](2026-07-28-04-typescript-export-semantics/2026-07-28-04-typescript-export-semantics-actual-status.md) |
| 05 | [plan](2026-07-28-05-module-export-and-reexport-resolution/2026-07-28-05-module-export-and-reexport-resolution-plan.md) | [evidence](2026-07-28-05-module-export-and-reexport-resolution/2026-07-28-05-module-export-and-reexport-resolution-evidence.md) | [benchmark](2026-07-28-05-module-export-and-reexport-resolution/2026-07-28-05-module-export-and-reexport-resolution-benchmark.md) | [actual status](2026-07-28-05-module-export-and-reexport-resolution/2026-07-28-05-module-export-and-reexport-resolution-actual-status.md) |
| 06 | `2026-07-28-06-ambient-external-resolution-and-diagnostics-plan.md` | `2026-07-28-06-ambient-external-resolution-and-diagnostics-evidence.md` | `2026-07-28-06-ambient-external-resolution-and-diagnostics-benchmark.md` | `2026-07-28-06-ambient-external-resolution-and-diagnostics-actual-status.md` |
| 07 | `2026-07-28-07-cross-surface-acceptance-and-target-validation-plan.md` | `2026-07-28-07-cross-surface-acceptance-and-target-validation-evidence.md` | `2026-07-28-07-cross-surface-acceptance-and-target-validation-benchmark.md` | `2026-07-28-07-cross-surface-acceptance-and-target-validation-actual-status.md` |

Target inventory: 7 plan folders, 28 standard child files, 7 P0 sections, preserved source phases P1-P7, 98 source-ID-preserved implementation slices, and 7 Pn-A/Pn-B/Pn-C closure sets.

## Execution Order And Handoffs

1. `2026-07-28-01-graph-identity-contract-and-strict-construction` establishes identity, lossless occurrence, RelationshipID, strict mutation, and shadow-v2 contracts while preserving the active v1 path; closure emits `2026-07-28-01-graph-identity-contract-and-strict-construction::E2-PNC-HANDOFF1` (source `P8-C`, local `Pn-C`).
2. `2026-07-28-02-versioned-persistence-and-v2-cutover` consumes that exact handoff, owns `index-reader-matrix.md`, cuts over all readers and immutable generations, activates identity v2, and emits `2026-07-28-02-versioned-persistence-and-v2-cutover::E2-PNC-HANDOFF1` (source `P8-C`, local `Pn-C`).
3. `2026-07-28-03-typescript-binding-pattern-extraction` opens only after consuming `2026-07-28-02-versioned-persistence-and-v2-cutover::E2-PNC-HANDOFF1` (source `P8-C`, local `Pn-C`), including its S0-S11/reader-matrix inspect-only denominator. Its internal order includes variable, parameter, catch, and for-of/for-in binding contexts before projection; closure emits `2026-07-28-03-typescript-binding-pattern-extraction::E2-PNC-HANDOFF1` (source `P8-C`, local `Pn-C`).
4. `2026-07-28-04-typescript-export-semantics` consumes `2026-07-28-03-typescript-binding-pattern-extraction::E2-PNC-HANDOFF1` (source `P8-C`, local `Pn-C`), `2026-07-28-01-graph-identity-contract-and-strict-construction::E1-P1A-CONTRACT1` (source/local `P1-A`), and the inspect-only denominator in `2026-07-28-02-versioned-persistence-and-v2-cutover::E2-PNC-HANDOFF1` (source `P8-C`, local `Pn-C`). Export facts remain distinct from access visibility and terminal barrel selection; closure emits `2026-07-28-04-typescript-export-semantics::E2-PNC-HANDOFF1` (source `P8-C`, local `Pn-C`).
5. `2026-07-28-05-module-export-and-reexport-resolution` consumes `2026-07-28-04-typescript-export-semantics::E2-PNC-HANDOFF1` (source `P8-C`, local `Pn-C`), owns the P5 semantic-vector manifest and terminal module/barrel traversal, and emits `2026-07-28-05-module-export-and-reexport-resolution::E2-PNC-HANDOFF1` (source `P8-C`, local `Pn-C`).
6. `2026-07-28-06-ambient-external-resolution-and-diagnostics` consumes immutable external candidates from `2026-07-28-05-module-export-and-reexport-resolution::E2-PNC-HANDOFF1` (source `P8-C`, local `Pn-C`) plus the inspect-only denominator from `2026-07-28-02-versioned-persistence-and-v2-cutover::E2-PNC-HANDOFF1` (source `P8-C`, local `Pn-C`), owns declaration authorization/materialization and the authoritative diagnostic status matrix, may not redo package-export resolution, and emits `2026-07-28-06-ambient-external-resolution-and-diagnostics::E2-PNC-HANDOFF1` (source `P8-C`, local `Pn-C`).
7. `2026-07-28-07-cross-surface-acceptance-and-target-validation` opens only after all six exact `::E2-PNC-HANDOFF1` records named in its inventory row (each source `P8-C`, local `Pn-C`) are accepted/committed. It validates determinism, S0-S11 parity, the real target boundary, and performance; failures return to the owning qualified child/slice.

Each non-terminal child Pn-C must update this roadmap and refresh the next child's `actual-status.md` from the latest accepted evidence before handing off; the update must include a refresh-log row, affected next-action/work-step changes, and qualified `E2-PNC-NEXTSTATUS1` proof. Child 07 records the terminal no-successor case and refreshes roadmap/campaign closure status. A stale or missing successor update blocks closure, and a child closure is a campaign checkpoint rather than campaign completion.

## Cross-Child Evidence Contract

- A handoff reference has the form `{child-slug, exact evidence ID, source slice, local slice}`.
- Child-local evidence namespaces may repeat across children because the child slug is the outer namespace.
- `2026-07-28-02-versioned-persistence-and-v2-cutover` owns the single reader matrix and its source-derived inventory. `2026-07-28-03-typescript-binding-pattern-extraction`, `2026-07-28-04-typescript-export-semantics`, `2026-07-28-06-ambient-external-resolution-and-diagnostics`, and `2026-07-28-07-cross-surface-acceptance-and-target-validation` consume `2026-07-28-02-versioned-persistence-and-v2-cutover::E2-PNC-HANDOFF1` (source `P8-C`, local `Pn-C`) inspect-only.
- `2026-07-28-05-module-export-and-reexport-resolution` owns the semantic vector and emits it through `2026-07-28-05-module-export-and-reexport-resolution::E2-PNC-HANDOFF1` (source `P8-C`, local `Pn-C`); `2026-07-28-06-ambient-external-resolution-and-diagnostics` owns the resolution-status matrix and emits it through `2026-07-28-06-ambient-external-resolution-and-diagnostics::E2-PNC-HANDOFF1` (source `P8-C`, local `Pn-C`); `2026-07-28-07-cross-surface-acceptance-and-target-validation` owns final cross-surface acceptance evidence.
- Global source provenance remains available from the legacy plan/ledgers and `2026-07-28-00-multi-plan-authoring::E1-P1A-MAP1`; it does not replace fresh child P0 or implementation evidence.

## Known Source-Migration Hazards

- `2026-07-28-01-graph-identity-contract-and-strict-construction` keeps the strict source chain A -> B -> C0 -> C0A -> C0B -> C -> D -> D1 -> D2 -> D3 -> E. Shadow-v2 must not mutate the active v1 path.
- `2026-07-28-02-versioned-persistence-and-v2-cutover` remains serial across all 42 slices even when an individual legacy gate names only a subset of earlier work. The reader matrix is a source seed; it is not runtime completion evidence.
- Legacy P3-C omits immediately preceding P3-B2A from its written gate. Conversion must copy that source block unchanged; authoring must not normalize it. Any technical correction requires a separate authorized plan edit after the mechanical split.
- Legacy P3-C1 has an ordered-list title different from its executable checklist heading. `2026-07-28-03-typescript-binding-pattern-extraction` uses the checklist title `Project binding JSON/Ladybug persistence adapters` and records the ordered-list alias as provenance.
- Legacy P4-A references `P1-A authority`. Preserve that source text unchanged and add any required qualified child-01 dependency outside the copied phase block; do not rewrite the source reference during conversion.
- `2026-07-28-03-typescript-binding-pattern-extraction` and `2026-07-28-04-typescript-export-semantics` preserve the source adapter order S0-S2, S3, S4, S7, S5, S8, S6, S9, S10, S11. Do not reorder numerically or merge independently committed surfaces.
- `2026-07-28-05-module-export-and-reexport-resolution` retains the full P5 semantic-vector manifest and publishes it through `2026-07-28-05-module-export-and-reexport-resolution::E2-PNC-HANDOFF1` (source `P8-C`, local `Pn-C`) even though the manifest sits outside the four slice bodies.
- `2026-07-28-06-ambient-external-resolution-and-diagnostics` retains the full P6 authoritative status matrix and publishes it through `2026-07-28-06-ambient-external-resolution-and-diagnostics::E2-PNC-HANDOFF1` (source `P8-C`, local `Pn-C`) even though the matrix sits after its slice bodies.
- `2026-07-28-07-cross-surface-acceptance-and-target-validation` owns validation only; it never patches a failure and never takes mutation ownership of the reader matrix received inspect-only through `2026-07-28-02-versioned-persistence-and-v2-cutover::E2-PNC-HANDOFF1` (source `P8-C`, local `Pn-C`).
- Legacy evidence section headings under-list some expanded nested slices while their tables and traceability rows contain the real evidence inventory. Child authoring must use exact slice/evidence rows, not the narrow heading text.
- Every child Pn is tailored to that child. Pn-C hands off to the next child/roadmap and cannot claim whole-campaign completion.

## Shared Artifact Ownership

| Artifact / contract | Sole mutation owner | Other consumers |
|---------------------|---------------------|-----------------|
| `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/index-reader-matrix.md` | `2026-07-28-02-versioned-persistence-and-v2-cutover` | Children 03/04/06 inspect and child 07 validates through `2026-07-28-02-versioned-persistence-and-v2-cutover::E2-PNC-HANDOFF1` (source `P8-C`, local `Pn-C`) |
| Identity/ownership decision | `2026-07-28-01-graph-identity-contract-and-strict-construction` | Children 02-07 consume `2026-07-28-01-graph-identity-contract-and-strict-construction::E2-PNC-HANDOFF1` (source `P8-C`, local `Pn-C`); exact decision authority is `2026-07-28-01-graph-identity-contract-and-strict-construction::E1-P1A-CONTRACT1` (source/local `P1-A`) |
| P5 semantic vector manifest | `2026-07-28-05-module-export-and-reexport-resolution` | Children 06-07 consume `2026-07-28-05-module-export-and-reexport-resolution::E2-PNC-HANDOFF1` (source `P8-C`, local `Pn-C`) |
| P6 authoritative resolution-status matrix | `2026-07-28-06-ambient-external-resolution-and-diagnostics` | Child 07 validates through `2026-07-28-06-ambient-external-resolution-and-diagnostics::E2-PNC-HANDOFF1` (source `P8-C`, local `Pn-C`) |
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
| 2026-07-28 | P1-A froze legacy hash and 98-row crosswalk | legacy remains active | `2026-07-28-00-multi-plan-authoring::E1-P1A-SNAPSHOT1`, `2026-07-28-00-multi-plan-authoring::E1-P1A-MAP1`, `2026-07-28-00-multi-plan-authoring::E1-P1A-MAP2` |
| 2026-07-28 | P1-B created candidate roadmap and seven-child inventory | legacy remains active; children not authored | `2026-07-28-00-multi-plan-authoring::E1-P1B-ROADMAP1`, `2026-07-28-00-multi-plan-authoring::E1-P1B-INVENTORY1`, `2026-07-28-00-multi-plan-authoring::E1-P1B-LINK1` |
| 2026-07-28 | P2-A authored child 01 with 11 mapped slices and 78 exact implementation evidence IDs | legacy remains active; child 01 is candidate and awaits authoring Supervisor | `2026-07-28-00-multi-plan-authoring::E2-P2A-FILES1`, `2026-07-28-00-multi-plan-authoring::E2-P2A-STRUCT1`, `2026-07-28-00-multi-plan-authoring::E2-P2A-MAP1` |
| 2026-07-28 | P2-A Supervisor PASS accepted child 01 after closing the false-upstream wording blocker | legacy remains active; child 01 authoring accepted; implementation remains unauthorized | `2026-07-28-00-multi-plan-authoring::E2-P2A-SUP1` |
| 2026-07-28 | User corrected the plan hierarchy and P2-A1 moved the authoring/roadmap/child artifacts toward three sibling roots | legacy remains active; P2-B paused until structural Supervisor PASS/commit | `2026-07-28-00-multi-plan-authoring::E2-P2A1-USER1`, pending `2026-07-28-00-multi-plan-authoring::E2-P2A1-SUP1` |
| 2026-07-28 | P2-A1 red-team resubmission and Supervisor PASS accepted the three-root correction | legacy remains active; P2-A1 awaited its isolated commit; child 02 remained excluded from that commit | `2026-07-28-00-multi-plan-authoring::E2-P2A1-REDTEAM1`, `2026-07-28-00-multi-plan-authoring::E2-P2A1-SUP1` |
| 2026-07-28 | P2-A1 structural correction committed as `55bf021f` | legacy remains active; child-02 authoring may open from the committed three-root basis | `2026-07-28-00-multi-plan-authoring::E2-P2A1-COMMIT1` |
| 2026-07-28 | P2-B authored the four-file child-02 candidate with 42 mapped slices and sole reader-matrix mutation ownership | legacy remains active; child 02 is linked as a candidate and awaits authoring Supervisor; implementation remains unauthorized | `2026-07-28-00-multi-plan-authoring::E2-P2B-FILES1`, `2026-07-28-00-multi-plan-authoring::E2-P2B-STRUCT1`, `2026-07-28-00-multi-plan-authoring::E2-P2B-MAP1`, `2026-07-28-00-multi-plan-authoring::E2-P2B-MATRIX1` |
| 2026-07-28 | User required every child to refresh the next child actual-status from latest evidence before closure; authored children and campaign contracts now carry the rule, Pn-C gate, and `NEXTSTATUS1` proof | legacy remains active; corrected P2-B candidate awaits successor-rule red-team and authoring Supervisor; implementation remains unauthorized | `2026-07-28-00-multi-plan-authoring::E2-P2B-USER2`, `2026-07-28-00-multi-plan-authoring::E2-P2B-HANDOFFRULE1`, `2026-07-28-00-multi-plan-authoring::E2-P2B-VALIDATION3` |
| 2026-07-28 | P2-B independent Supervisor PASS accepted child 02 and the successor-freshness correction | legacy remains active; child 02 authoring is accepted/commit pending; P2-C and implementation remain unauthorized until their gates open | `2026-07-28-00-multi-plan-authoring::E2-P2B-SUP1` |
| 2026-07-28 | Owner rejected the rewritten/remapped child output; child 01 was deleted and rebuilt by exact source-block copy/paste, while child 02 and later artifacts were removed/reset | legacy remains active; only rebuilt child 01 is accepted for the bounded conversion scope; child 02-07 are not authored and implementation remains unauthorized | `2026-07-28-00-multi-plan-authoring::E2-P2A-REBUILD1`, `2026-07-28-00-multi-plan-authoring::E2-P2A-SUP2`, `2026-07-28-00-multi-plan-authoring::E2-P2B-RESET1` |
| 2026-07-28 | Rebuilt Child 01, later-child reset, current roadmap/authoring tracking, and both Supervisor reports committed as `ce82a341` | legacy remains active; Child 01 authoring milestone is durable; Child 02 stays blocked/not authored and implementation remains unauthorized | `2026-07-28-00-multi-plan-authoring::E2-P2A-COMMIT2` |
| 2026-07-28 | Corrected the 98-row crosswalk to preserve source IDs and authored Child 02 by exact P2/E2/B2/source-status block copy | legacy remains active; Child 02 is a review-pending authoring candidate; Child 03 and implementation remain unauthorized | `2026-07-28-00-multi-plan-authoring::E1-P1A-MAP3`, `2026-07-28-00-multi-plan-authoring::E2-P2B-RFILES1`, `2026-07-28-00-multi-plan-authoring::E2-P2B-RVALID1` |
| 2026-07-28 | Child 02 exact-copy candidate passed bounded red-team closure and Supervisor review | legacy remains active; Child 02 commit is pending; Child 03 and implementation remain unauthorized | `2026-07-28-00-multi-plan-authoring::E2-P2B-RREDTEAM1`, `2026-07-28-00-multi-plan-authoring::E2-P2B-RSUP1` |
| 2026-07-28 | Child 02 exact-copy authoring set committed as `a1c66865` | legacy remains active; Child 03 authoring opens; production implementation remains unauthorized | `2026-07-28-00-multi-plan-authoring::E2-P2B-RCOMMIT1` |
| 2026-07-28 | Child 03 four-file candidate created by exact P3/E3/B3/source-status copy with 17 unchanged source IDs | legacy remains active; Child 03 validation is active; Child 04 and implementation remain unauthorized | `2026-07-28-00-multi-plan-authoring::E2-P2C-FILES1`, `2026-07-28-00-multi-plan-authoring::E2-P2C-MAP1`, `2026-07-28-00-multi-plan-authoring::E2-P2C-VALID1` |
| 2026-07-28 | Child 03 exact-copy candidate passed bounded red-team and independent Supervisor review | legacy remains active; Child 03 commit is pending; Child 04 and implementation remain unauthorized | `2026-07-28-00-multi-plan-authoring::E2-P2C-REDTEAM1`, `2026-07-28-00-multi-plan-authoring::E2-P2C-SUP1` |
| 2026-07-28 | Child 03 exact-copy authoring set committed as `35a0611c` | legacy remains active; Child 04 authoring opens; production implementation remains unauthorized | `2026-07-28-00-multi-plan-authoring::E2-P2C-COMMIT1` |
| 2026-07-28 | Child 04 four-file candidate created by exact P4/E4/B4/source-status copy with 15 unchanged source IDs | legacy remains active; Child 04 validation is active; Child 05 and implementation remain unauthorized | `2026-07-28-00-multi-plan-authoring::E2-P2D-FILES1`, `2026-07-28-00-multi-plan-authoring::E2-P2D-MAP1`, `2026-07-28-00-multi-plan-authoring::E2-P2D-VALID1` |
| 2026-07-28 | Child 04 exact-copy candidate passed independent checks and Supervisor review | legacy remains active; Child 04 commit is pending; Child 05 and implementation remain unauthorized | `2026-07-28-00-multi-plan-authoring::E2-P2D-CHECK1`, `2026-07-28-00-multi-plan-authoring::E2-P2D-SUP1` |
| 2026-07-28 | Child 04 exact-copy authoring set committed as `2de220bb` | legacy remains active; Child 05 authoring opens; production implementation remains unauthorized | `2026-07-28-00-multi-plan-authoring::E2-P2D-COMMIT1` |
| 2026-07-28 | Child 05 four-file candidate created by exact P5/E5/B5/source-status copy with four unchanged source IDs and semantic-vector contract material | legacy remains active; Child 05 validation is active; Child 06 and implementation remain unauthorized | `2026-07-28-00-multi-plan-authoring::E2-P2E-FILES1`, `2026-07-28-00-multi-plan-authoring::E2-P2E-MAP1`, `2026-07-28-00-multi-plan-authoring::E2-P2E-VALID1` |
| 2026-07-28 | Child 05 exact-copy candidate passed independent checks and Supervisor review | legacy remains active; Child 05 commit is pending; Child 06 and implementation remain unauthorized | `2026-07-28-00-multi-plan-authoring::E2-P2E-CHECK1`, `2026-07-28-00-multi-plan-authoring::E2-P2E-SUP1` |

## Candidate Acceptance Gate

This roadmap may become active only when:

- all seven child folders and all 28 standard files exist;
- every child has populated P0, its preserved source phase and source IDs, and tailored Pn-A/Pn-B/Pn-C;
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
