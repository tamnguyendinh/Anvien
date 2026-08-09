# Anvien Graph Identity and TypeScript Resolution Correctness v2 Roadmap

Date: 2026-07-28
Status: active / Owner-authorized 2026-08-09 / P3-B authority cutover PASS / seven-child sequential implementation in progress / Child 01 open
Source plan: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/2026-07-26-anvien-graph-identity-resolution-v2-plan.md`
Plan-set authoring plan: `docs/plans/2026-07-28-00-multi-plan-authoring/2026-07-28-00-multi-plan-authoring-plan.md`
Multi-plan root: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/`
Exact crosswalk: `docs/plans/2026-07-28-00-multi-plan-authoring/2026-07-28-00-multi-plan-authoring-evidence.md#e1-p1a-map1---exact-source-to-child-crosswalk`

## Goal

Coordinate the graph-identity and TypeScript-resolution remediation as seven complete implementation child plans, one child for each accepted legacy implementation phase P1-P7. Keep technical order, all 98 implementation slices, independent P0/evidence/benchmark/closure ledgers, explicit handoffs, and one active authority without turning this roadmap into another implementation plan.

## Authority State

- Active authority now: this roadmap plus the seven numbered child plan sets. The Owner authorized execution on 2026-08-09 and P3-B Supervisor review passed in `reports/Supervisor/rp_supervisor_260809_212310_by_gpt-5-codex_multiplan_authority_cutover.md`.
- Legacy authority: the source plan at SHA-256 `365E17A7F7CD539426568A2874A1AD1231D120C5391BFE3D8B875D5F760A45FB` is preserved as `superseded / reference-only`.
- Partial child implementation remains non-authoritative; only accepted and committed slice evidence may advance the campaign.
- Authority cutover completed after the authoring plan's P3-A deterministic checks and P3-B unconditional Supervisor PASS.
- Current execution: Child 01 is the only open implementation child. Children 02-07 remain blocked by their exact predecessor handoffs.
- The cutover preserved the legacy technical body and companion ledgers and changed only its status/pointer to `superseded / reference-only`.
- This roadmap is the sole campaign-order/status index and each child plan is the sole execution authority for its mapped legacy phase.
- The 2026-08-09 Owner instruction authorizes production implementation under each child's existing gate; it does not bypass slice dependencies, evidence, review, or commit boundaries.

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
- Every file owns one primary planning or implementation responsibility. A file may link to multiple modules/files only when all links serve that responsibility. A preserve-only/link row may appear in more than one ordered job only when it names the same coordinator responsibility, carries no new semantic logic, and has one explicitly named write-owner; the link rows are not additional business owners and may not be edited concurrently. Any new responsibility requires a dedicated owner file and a changed ownership row before editing.
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
| 01 | [2026-07-28-01-graph-identity-contract-and-strict-construction](2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-plan.md) | Identity contract, lossless declaration/source-site identity, strict graph construction, shadow-v2 proof | P1 | 11 | active / P0 refreshed / P1-A accepted pending checkpoint commit | No predecessor; closure must update Child 02 actual-status from latest evidence before handoff |
| 02 | [2026-07-28-02-versioned-persistence-and-v2-cutover](2026-07-28-02-versioned-persistence-and-v2-cutover/2026-07-28-02-versioned-persistence-and-v2-cutover-plan.md) | Compatibility manifest, opaque readers, S0-S11 parity, atomic generations, v2 cutover | P2 | 42 | ready / dependency-blocked | Requires Child 01 qualified implementation handoff; sole reader-matrix mutation owner |
| 03 | [2026-07-28-03-typescript-binding-pattern-extraction](2026-07-28-03-typescript-binding-pattern-extraction/2026-07-28-03-typescript-binding-pattern-extraction-plan.md) | Recursive binding facts, declaration contexts, graph and adapter projection | P3 | 17 | ready / dependency-blocked | Requires `2026-07-28-02-versioned-persistence-and-v2-cutover::E2-PNC-HANDOFF1` |
| 04 | [2026-07-28-04-typescript-export-semantics](2026-07-28-04-typescript-export-semantics/2026-07-28-04-typescript-export-semantics-plan.md) | ExportFact semantics, export syntax extraction, graph and adapter projection | P4 | 15 | ready / dependency-blocked | Requires Child 03 handoff; consumes the Child 01 contract and Child 02 denominator inspect-only |
| 05 | [2026-07-28-05-module-export-and-reexport-resolution](2026-07-28-05-module-export-and-reexport-resolution/2026-07-28-05-module-export-and-reexport-resolution-plan.md) | Module export tables, aliases, cycles, ambiguity, terminal barrel resolution | P5 | 4 | ready / dependency-blocked | Requires Child 04 qualified handoff |
| 06 | [2026-07-28-06-ambient-external-resolution-and-diagnostics](2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-plan.md) | Declaration universe, external authorization/materialization, immutable outcomes, diagnostics | P6 | 6 | ready / dependency-blocked | Requires Child 05 qualified handoff; consumes Child 02 denominator inspect-only |
| 07 | [2026-07-28-07-cross-surface-acceptance-and-target-validation](2026-07-28-07-cross-surface-acceptance-and-target-validation/2026-07-28-07-cross-surface-acceptance-and-target-validation-plan.md) | Full determinism/parity/target/performance acceptance; no implementation repair | P7 | 3 | ready / six-handoff dependency-blocked | Requires all six qualified predecessor handoffs |

## Standard File Inventory

Every row below represents four required files inside its plan folder. Planned paths remain code-form until the corresponding P2 authoring slice creates the four-file candidate; afterward the roadmap links the real files while the Status column remains the acceptance authority.

| Child | Plan | Evidence | Benchmark | Actual status |
|-------|------|----------|-----------|---------------|
| 01 | [plan](2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-plan.md) | [evidence](2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-evidence.md) | [benchmark](2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-benchmark.md) | [actual status](2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-actual-status.md) |
| 02 | [plan](2026-07-28-02-versioned-persistence-and-v2-cutover/2026-07-28-02-versioned-persistence-and-v2-cutover-plan.md) | [evidence](2026-07-28-02-versioned-persistence-and-v2-cutover/2026-07-28-02-versioned-persistence-and-v2-cutover-evidence.md) | [benchmark](2026-07-28-02-versioned-persistence-and-v2-cutover/2026-07-28-02-versioned-persistence-and-v2-cutover-benchmark.md) | [actual status](2026-07-28-02-versioned-persistence-and-v2-cutover/2026-07-28-02-versioned-persistence-and-v2-cutover-actual-status.md) |
| 03 | [plan](2026-07-28-03-typescript-binding-pattern-extraction/2026-07-28-03-typescript-binding-pattern-extraction-plan.md) | [evidence](2026-07-28-03-typescript-binding-pattern-extraction/2026-07-28-03-typescript-binding-pattern-extraction-evidence.md) | [benchmark](2026-07-28-03-typescript-binding-pattern-extraction/2026-07-28-03-typescript-binding-pattern-extraction-benchmark.md) | [actual status](2026-07-28-03-typescript-binding-pattern-extraction/2026-07-28-03-typescript-binding-pattern-extraction-actual-status.md) |
| 04 | [plan](2026-07-28-04-typescript-export-semantics/2026-07-28-04-typescript-export-semantics-plan.md) | [evidence](2026-07-28-04-typescript-export-semantics/2026-07-28-04-typescript-export-semantics-evidence.md) | [benchmark](2026-07-28-04-typescript-export-semantics/2026-07-28-04-typescript-export-semantics-benchmark.md) | [actual status](2026-07-28-04-typescript-export-semantics/2026-07-28-04-typescript-export-semantics-actual-status.md) |
| 05 | [plan](2026-07-28-05-module-export-and-reexport-resolution/2026-07-28-05-module-export-and-reexport-resolution-plan.md) | [evidence](2026-07-28-05-module-export-and-reexport-resolution/2026-07-28-05-module-export-and-reexport-resolution-evidence.md) | [benchmark](2026-07-28-05-module-export-and-reexport-resolution/2026-07-28-05-module-export-and-reexport-resolution-benchmark.md) | [actual status](2026-07-28-05-module-export-and-reexport-resolution/2026-07-28-05-module-export-and-reexport-resolution-actual-status.md) |
| 06 | [plan](2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-plan.md) | [evidence](2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-evidence.md) | [benchmark](2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-benchmark.md) | [actual status](2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-actual-status.md) |
| 07 | [plan](2026-07-28-07-cross-surface-acceptance-and-target-validation/2026-07-28-07-cross-surface-acceptance-and-target-validation-plan.md) | [evidence](2026-07-28-07-cross-surface-acceptance-and-target-validation/2026-07-28-07-cross-surface-acceptance-and-target-validation-evidence.md) | [benchmark](2026-07-28-07-cross-surface-acceptance-and-target-validation/2026-07-28-07-cross-surface-acceptance-and-target-validation-benchmark.md) | [actual status](2026-07-28-07-cross-surface-acceptance-and-target-validation/2026-07-28-07-cross-surface-acceptance-and-target-validation-actual-status.md) |

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
- `ready / dependency-blocked`: Owner authorization exists, but the child still requires its exact predecessor handoff and P0 refresh.
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
| 2026-07-28 | Child 05 exact-copy authoring set committed as `b19256e6` | legacy remains active; Child 06 authoring opens; production implementation remains unauthorized | `2026-07-28-00-multi-plan-authoring::E2-P2E-COMMIT1` |
| 2026-07-28 | Child 06 four-file candidate created by exact P6/E6/B6/source-status copy with six unchanged source IDs and the authoritative diagnostic matrix | legacy remains active; Child 06 validation is active; Child 07 and implementation remain unauthorized | `2026-07-28-00-multi-plan-authoring::E2-P2F-FILES1`, `2026-07-28-00-multi-plan-authoring::E2-P2F-MAP1`, `2026-07-28-00-multi-plan-authoring::E2-P2F-VALID1` |
| 2026-07-28 | Child 06 exact-copy candidate passed independent checks and Supervisor review | legacy remains active; Child 06 commit is pending; Child 07 and implementation remain unauthorized | `2026-07-28-00-multi-plan-authoring::E2-P2F-CHECK1`, `2026-07-28-00-multi-plan-authoring::E2-P2F-SUP1` |
| 2026-07-28 | Child 06 exact-copy authoring set committed as `e0582469` | legacy remains active; Child 07 authoring opens; production implementation remains unauthorized | `2026-07-28-00-multi-plan-authoring::E2-P2F-COMMIT1` |
| 2026-07-28 | Child 07 four-file candidate created by exact P7/E7/B7/source-status copy with three unchanged source IDs and terminal closure rule | legacy remains active; Child 07 Supervisor review pending; implementation remains unauthorized | `2026-07-28-00-multi-plan-authoring::E2-P2G-FILES1`, `2026-07-28-00-multi-plan-authoring::E2-P2G-MAP1`, `2026-07-28-00-multi-plan-authoring::E2-P2G-CLOSURE1` |
| 2026-07-28 | Child 07 exact-copy candidate passed independent checks and Supervisor review | legacy remains active; Child 07 commit is pending; production implementation remains unauthorized | `2026-07-28-00-multi-plan-authoring::E2-P2G-SUP1` |
| 2026-07-28 | Child 07 exact-copy authoring set committed as `4f1c94e5`; terminal child refreshed roadmap/campaign status | legacy remains active; all seven child sets are durable; full-campaign closure audit opens; production implementation remains unauthorized | `2026-07-28-00-multi-plan-authoring::E2-P2G-COMMIT1` |
| 2026-07-28 | P3-A full seven-child closure audit passed: 7 folders, 28 standard files, 98/98 source IDs, 35/0 links, exact P/E/B blocks, lifecycle/ownership checks, and current graph evidence | legacy remains active; bounded authoring closure is accepted; P3-B authority cutover and production implementation remain not authorized | `2026-07-28-00-multi-plan-authoring::E3-P3A-STRUCT1`, `2026-07-28-00-multi-plan-authoring::E3-P3A-LINK1`, `2026-07-28-00-multi-plan-authoring::E3-P3A-OWNER1`, `2026-07-28-00-multi-plan-authoring::E3-P3A-MAP1`, `2026-07-28-00-multi-plan-authoring::E3-P3A-FIELDS1`, `2026-07-28-00-multi-plan-authoring::E3-P3A-GRAPH1`, `2026-07-28-00-multi-plan-authoring::E3-P3A-SUPERVISOR1` |
| 2026-08-09 | Owner authorized execution and P3-B Supervisor PASS activated the roadmap | roadmap is sole active authority; legacy reference-only; Child 01 open | `2026-07-28-00-multi-plan-authoring::E3-P3B-SUPERVISOR1`, `2026-07-28-00-multi-plan-authoring::E3-P3B-CUTOVER1` |

Historical-status note: the earlier history rows describing the rejected/remapped Child 01 and pre-reset Child 02 candidate are preserved as audit history. Current Child 01 authority begins at `E2-P2A-REBUILD1`/`E2-P2A-SUP2`/`E2-P2A-COMMIT2`; those historical candidate hashes must not be read as current artifact hashes.

## Semantic Remediation Overlay (mandatory)

The seven source phase blocks and all 98 source slices remain byte-for-byte historical provenance. This overlay is a stronger execution contract added after `reports/Supervisor/rp_supervisor_260728_222405_by_gpt-5-codex_multiplan_semantic_adequacy.md`. If a copied gate is weaker, incomplete, or contradictory, this overlay controls; the copied source text is not rewritten.

- A child closure is a local checkpoint. Child 01-06 may close after its own local implementation, validation, Supervisor, ledger, detect-changes, and commit gates plus the qualified successor refresh. Pending Child 07/P7 performance rows block campaign/release closure only; they do not deadlock an earlier local child. Child 07 owns the campaign/release gate.
- Every non-terminal `Pn-C` must create both qualified evidence IDs `E2-PNC-NEXTSTATUS1` and `E2-PNC-HANDOFF1`. `NEXTSTATUS1` proves the successor `actual-status.md` was refreshed from the latest accepted evidence, a refresh-log row was appended, and affected next-action/work-step rows were updated. `HANDOFF1` proves the predecessor evidence set, local commit, owner decision, and successor opening conditions; its opening condition consumes this child’s own qualified `NEXTSTATUS1` after the successor refresh, never a future `NEXTSTATUS1` owned by the successor. Child 07 uses the same IDs for the terminal roadmap/campaign refresh and terminal handoff record; it must state `no successor` explicitly.
- Every implementation slice Acceptance is conjunctive with its complete evidence-ledger row. A narrow `REVIEW1` line cannot close a slice while any required `IMPACT`, `SRC`, `BUILD`, behavior `TEST`/oracle, `DETECT`, or `COMMIT` item is pending. Any `TBD`, wildcard, or unqualified owner/evidence reference fails the gate.
- Before editing, every implementation job must publish a machine-readable ownership row: `File` (exact path), `unique responsibility`, `allowed links`, and `prohibited contents`. The row covers production, test, generated, and fixture files; one file may link to many modules only in service of its one responsibility. Wildcard paths, `TBD`, catch-all owners, or mixed responsibilities fail the job gate.
- The roadmap-level manifest/handshake is authoritative: the nine semantic-correction fields `graphSchemaVersion`, `identitySchemaVersion`, `scopeIRSchemaVersion`, `graphGeneration`, `analyzerVersion`, `columnEncoding`/`positionEncoding`, `sourceFingerprint`, and `configFingerprint` are required in the persisted manifest and are represented in the reader handshake. The complete wire inventory is 15 persisted manifest fields (those nine plus six protocol/provenance fields) and 10 request fields. The request key `supportedScopeIrVersions[]` is the fixed wire spelling whose values are compared to the case-sensitive manifest field `scopeIRSchemaVersion`; no alternate spelling or omitted field is accepted. A missing field is a contract failure, not an optional compatibility omission.
- P4 owns syntactic `ExportFact`, `ModuleRequestFact`, and `ImportBindingFact` production and direct-export counts only. P5 owns module/export-table traversal and derived `resolvedExportEntryCount`/`publicApiSymbolCount`; P4 must not claim barrel reachability or package public API. The three counts have separate owners and ledger metrics.
- P6 declaration-universe capability is explicit: `exact|structural|degraded`, with `confidence` and `completeness` on every outcome and one immutable `DeclarationCapabilityDescriptor` in semantic graph metadata per `graphGeneration`. The descriptor is generation/config/catalog-bound and is separate from Child 02's fixed 15-field compatibility manifest. The universe covers repository declarations, project-owned `.d.ts`, installed package declarations, stdlib, intrinsics, ambient modules, and global augmentations, with parser/merge acceptance. External symbols are excluded by default from context, impact, rename, process, and group traversal; only an explicit `include_external` option opts in.
- P7 is a semantic oracle, not only parity: it must assert external isolation on context/impact/rename/process/groups, preserve syntactic `IMPORTS` path-resolution counts while resolving exports, accept a barrel with zero physical declarations only when its export surface is proven, and require `Promise`/`Math` outcomes to be `resolved_external` or an explicit external-capability failure—never `resolved_intrinsic`.

### Semantic Remediation Evidence Matrix (mandatory)

The overlay is executable only when each owning child binds the invariant to a named gate, a separate benchmark row, and a separate pending evidence declaration. Aggregate rows are not sufficient proof and may not be used to close a child or the campaign.

| Invariant | Owning child | Required gate/evidence binding | Minimum measurable result |
|---|---|---|---|
| for-of/for-in binding dependency | Child 03 | `P3-C` overlay gate + `E2-PNC-BINDING1` + `E3-P3C-B2A-GATE1` | `P3-B2A` accepted/committed before projection; 6 named `.map()` sites and `ResolutionGap=0` |
| manifest/handshake completeness | Child 02 (later children inspect-only) | `E2-P2A-MANIFEST1`, `E2-P2A-HANDSHAKE1`, `E2-P2A-METADATA1` | every persisted manifest and handshake carries all nine required metadata fields; zero body-open on mismatch |
| P4 fact production and direct count | Child 04 | `E2-PNC-EXPORT1A..1D`, `E4-P4-FACT1` | one `ModuleRequestFact`/`ImportBindingFact` per syntax site, one syntactic `ExportFact`, exact direct count, no derived P5 fields |
| P5 derived export state | Child 05 | `E2-PNC-MODULE1A..1E`, `E5-P5-DERIVED1` | separate `resolvedExportEntryCount` and `publicApiSymbolCount` owners and metrics; no P4 mutation |
| declaration capability/source coverage | Child 06 | `E2-PNC-AMBIENT1A..1H`, `E2-PNC-AMBIENT1K`, `E6-P6-SOURCES1` | explicit capability/confidence/completeness for all seven source categories, parser/merge matrix complete |
| generation-level declaration capability metadata | Child 06 + Child 07 oracle | `E2-PNC-AMBIENT1S`, `E2-PNC-FINALORACLE1K` | exactly one descriptor per graphGeneration; no outcome is stronger than the descriptor; S0-S11 differences and mixed-generation refs are zero |
| external downstream isolation | Child 06 + Child 07 oracle | `E2-PNC-AMBIENT1I`, `E2-PNC-AMBIENT1L..1P`, `E2-PNC-FINALORACLE1A..1F` | five surfaces each prove default exclusion, typed-option propagation, and explicit `include_external` opt-in |
| zero-physical barrel/path accounting | Child 05 + Child 07 oracle | `E2-PNC-MODULE1B..1E`, `E2-PNC-FINALORACLE1G..1H` | `physicalDeclarationCount=0`, `resolvedExportEntryCount>0`, and pre/post path plus `IMPORTS` counts unchanged |
| Promise/Math outcomes | Child 06 + Child 07 oracle | `E2-PNC-AMBIENT1J`, `E2-PNC-AMBIENT1Q..1R`, `E2-PNC-FINALORACLE1I`, `E2-PNC-FINALORACLE1J` | each real target site is `resolved_external` 3/3; only named negative fixtures may use explicit external-capability failure; zero `resolved_intrinsic` |

Every correction evidence declaration remains `pending` until implementation and validation actually produce it. This matrix closes a contract/traceability gap; it does not claim any implementation slice is complete.

## Candidate Acceptance Gate

This roadmap may become active only when:

- all seven child folders and all 28 standard files exist;
- every child has populated P0, its preserved source phase and source IDs, and tailored Pn-A/Pn-B/Pn-C;
- deterministic validation proves 98 source slices, 98 unique destinations, zero missing, zero duplicate, zero extra, and preserved order/required fields;
- all companion and cross-plan links resolve;
- every mutable artifact has one owner and `index-reader-matrix.md` belongs only to child 02;
- every child contains the mandatory semantic-remediation overlay, qualified `E2-PNC-NEXTSTATUS1`/`E2-PNC-HANDOFF1` declarations, local-versus-campaign closure rule, full evidence-row binding, and job-granular ownership-table gate;
- P3-C is gated by P3-B2A; binding, manifest, fact ownership, module, ambient, capability, external-isolation, and Promise/Math semantic gates are measurable in the owning child ledgers;
- every semantic-correction row above has a child-local evidence declaration, a benchmark row with explicit numerator/denominator or before/after values, and a qualified closure binding; a single aggregate `REVIEW1`, parity result, or Supervisor note cannot substitute for those rows;
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
